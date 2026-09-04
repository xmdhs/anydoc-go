//! C-ABI wasm exports wrapping anydoc's conversion core.
//!
//! The module deliberately avoids wasm-bindgen and any JS dependency: every
//! export is a bare `extern "C"` function taking pointers into the linear
//! memory, so the wasm can be translated with wasm2go and driven by a plain
//! Go host.
//!
//! Buffer protocol: allocations use 8-byte-aligned layouts, and every
//! allocated pointer must be released with `anydoc_free(ptr, size)` where
//! `size` is the length the buffer was allocated with (0 for an empty
//! buffer never comes out of the allocator).

use std::alloc::{alloc, dealloc, Layout};
use std::ptr;
use std::slice;

use anydoc::{ConvertError, Format};

/// Stable error code reported through the `out_code` argument of
/// [`anydoc_to_markdown`]; 0 means the result is Markdown.
pub const ERR_OK: u32 = 0;
pub const ERR_UNSUPPORTED: u32 = 1;
pub const ERR_MALFORMED: u32 = 2;
pub const ERR_ENCRYPTED: u32 = 3;
pub const ERR_RESOURCE_LIMIT: u32 = 4;
pub const ERR_MISSING_PART: u32 = 5;
pub const ERR_IO: u32 = 6;
pub const ERR_OTHER: u32 = 7;
pub const ERR_NEEDS_OCR: u32 = 8;

/// Map a stable error to its code (mirrors `ConvertError::code`).
fn code_of(err: &ConvertError) -> u32 {
    match err.code() {
        "unsupported" => ERR_UNSUPPORTED,
        "malformed" => ERR_MALFORMED,
        "encrypted" => ERR_ENCRYPTED,
        "resourceLimit" => ERR_RESOURCE_LIMIT,
        "missingPart" => ERR_MISSING_PART,
        "io" => ERR_IO,
        "needsOcr" => ERR_NEEDS_OCR,
        _ => ERR_OTHER,
    }
}

/// Write a failure buffer and its code; NULL when out of memory.
fn fail(err_msg: &str, code: u32, out_len: *mut usize, out_code: *mut u32) -> *mut u8 {
    let ptr = store_bytes(err_msg.as_bytes(), out_len);
    // SAFETY: caller-provided out slots.
    unsafe { *out_code = code }
    ptr
}

/// Parse the caller-supplied format extension (`fmt_len == 0` means
/// detect-from-content); `Err((message, code))` on invalid input.
fn resolve_format(fmt: *const u8, fmt_len: usize) -> Result<Option<Format>, (String, u32)> {
    if fmt_len == 0 {
        return Ok(None);
    }
    if fmt.is_null() {
        return Err(("invalid format pointer".to_string(), ERR_UNSUPPORTED));
    }
    // SAFETY: fmt_len bytes of valid memory.
    let fmt_bytes = unsafe { slice::from_raw_parts(fmt, fmt_len) };
    let Some(ext) = std::str::from_utf8(fmt_bytes).ok() else {
        return Err(("format is not valid UTF-8".to_string(), ERR_UNSUPPORTED));
    };
    let ext = ext.strip_prefix('.').unwrap_or(ext);
    let Some(format) = Format::from_extension(ext) else {
        return Err((format!("unrecognized format: {ext}"), ERR_UNSUPPORTED));
    };
    Ok(Some(format))
}

/// Append a little-endian u32 (the serialized asset stream is little-endian
/// on every wasm target).
fn push_u32(buf: &mut Vec<u8>, v: u32) {
    buf.extend_from_slice(&v.to_le_bytes());
}

/// Allocate `size` bytes in linear memory (8-byte aligned).
///
/// Returns the pointer the caller must eventually pass to [`anydoc_free`]
/// with the same `size`, or NULL when `size` is zero or allocation failed.
#[unsafe(no_mangle)]
pub extern "C" fn anydoc_alloc(size: usize) -> *mut u8 {
    let Ok(layout) = Layout::from_size_align(size, 8) else {
        return ptr::null_mut();
    };
    // SAFETY: layout is our own 8-byte-aligned allocation.
    unsafe { alloc(layout) }
}

/// Release a buffer allocated by [`anydoc_alloc`], or returned by
/// [`anydoc_to_markdown`] (pass its reported length as `size`).
#[unsafe(no_mangle)]
pub extern "C" fn anydoc_free(ptr: *mut u8, size: usize) {
    if ptr.is_null() || size == 0 {
        return;
    }
    let Ok(layout) = Layout::from_size_align(size, 8) else {
        return;
    };
    // SAFETY: ptr came from our alloc with the same 8-byte layout.
    unsafe { dealloc(ptr, layout) }
}

/// Convert a document to GitHub-Flavored Markdown.
///
/// Arguments:
/// - `input`, `input_len`: the document bytes.
/// - `fmt`, `fmt_len`: a bare extension naming a format to force ("docx",
///   "csv", "pdf", ...). `fmt_len == 0` (and any pointer) detects the format
///   from the content instead; signature-less formats (CSV) must name the
///   extension.
/// - `out_len`, `out_code`: out parameters receiving the result length and
///   the [`ERR_*`] code (0 = success).
///
/// Returns a freshly allocated buffer with the Markdown (when `out_code` is
/// 0) or the error message (otherwise); release it with
/// `anydoc_free(ptr, *out_len)`.
#[unsafe(no_mangle)]
pub extern "C" fn anydoc_to_markdown(
    input: *const u8,
    input_len: usize,
    fmt: *const u8,
    fmt_len: usize,
    out_len: *mut usize,
    out_code: *mut u32,
) -> *mut u8 {
    // SAFETY: caller-provided out slots; we always initialize them first.
    unsafe {
        *out_code = ERR_OK;
        *out_len = 0;
    }

    if input_len > 0 && input.is_null() {
        return fail("invalid input pointer", ERR_OTHER, out_len, out_code);
    }
    // SAFETY: input_len bytes of valid memory when input_len > 0 (or NULL).
    let input_bytes = unsafe { slice::from_raw_parts(input, input_len) };

    let format = match resolve_format(fmt, fmt_len) {
        Ok(format) => format,
        Err((msg, code)) => return fail(&msg, code, out_len, out_code),
    };

    let markdown = match anydoc::to_markdown_bytes(input_bytes, format) {
        Ok(markdown) => markdown,
        Err(err) => {
            let code = code_of(&err);
            return fail(&format!("{err}"), code, out_len, out_code);
        }
    };

    let ptr = store_bytes(markdown.as_bytes(), out_len);
    if ptr.is_null() && !markdown.is_empty() {
        // SAFETY: out slot.
        unsafe { *out_code = ERR_RESOURCE_LIMIT }
        return ptr::null_mut();
    }
    // SAFETY: out_code is still ERR_OK from initialization; length was set
    // by store_bytes.
    ptr
}

/// Extract embedded assets (images, object payloads) of a document.
///
/// Same input/fmt protocol as [`anydoc_to_markdown`]. `out_code == 0` on
/// success even when the document has no assets (the stream then just holds a
/// zero count). The returned buffer is the serialized asset stream:
///
/// ```text
/// u32 count
///   per asset (little-endian):
///     u32 id            index into the document's assets (matches the
///                       `asset://<id>` placeholder rendered into Markdown)
///     u32 type_len      media type bytes
///     bytes             e.g. "image/png"
///     u32 bytes_len     payload bytes
///     bytes
/// ```
///
/// Release the buffer with `anydoc_free(ptr, *out_len)`. PDF 等未编入格式
/// 与 to_markdown 同样报 unsupported。
#[unsafe(no_mangle)]
pub extern "C" fn anydoc_assets(
    input: *const u8,
    input_len: usize,
    fmt: *const u8,
    fmt_len: usize,
    out_len: *mut usize,
    out_code: *mut u32,
) -> *mut u8 {
    // SAFETY: caller-provided out slots; we always initialize them first.
    unsafe {
        *out_code = ERR_OK;
        *out_len = 0;
    }

    if input_len > 0 && input.is_null() {
        return fail("invalid input pointer", ERR_OTHER, out_len, out_code);
    }
    // SAFETY: input_len bytes of valid memory when input_len > 0 (or NULL).
    let input_bytes = unsafe { slice::from_raw_parts(input, input_len) };

    let format = match resolve_format(fmt, fmt_len) {
        Ok(format) => format,
        Err((msg, code)) => return fail(&msg, code, out_len, out_code),
    };

    let document = match anydoc::to_document(input_bytes, format) {
        Ok(document) => document,
        Err(err) => return fail(&format!("{err}"), code_of(&err), out_len, out_code),
    };

    store_bytes(&serialize_assets(&document), out_len)
}

/// Serialize a document's assets into the stream format shared by
/// [`anydoc_assets`] and [`anydoc_convert`] (little-endian):
/// `u32 count`, then per asset `u32 id`, `u32 type_len` + bytes,
/// `u32 bytes_len` + bytes. `id` is the index into `Document::assets`, so it
/// matches the `asset://<id>` placeholder the renderer writes.
fn serialize_assets(document: &anydoc::model::Document) -> Vec<u8> {
    let mut buf: Vec<u8> = Vec::new();
    push_u32(&mut buf, document.assets.len() as u32);
    for asset in &document.assets {
        push_u32(&mut buf, asset.id.0 as u32);
        let media_type = asset.media_type.as_bytes();
        push_u32(&mut buf, media_type.len() as u32);
        buf.extend_from_slice(media_type);
        push_u32(&mut buf, asset.bytes.len() as u32);
        buf.extend_from_slice(&asset.bytes);
    }
    buf
}

/// Convert a document to Markdown and return its embedded assets in one
/// wasm call (single `to_document` parse), so a host can avoid a second
/// parse to fetch assets (`--imgs`). Same input/fmt protocol as
/// [`anydoc_to_markdown`]; `md` is the Markdown (or error message on
/// `out_code != 0`), `assets` is the serialized asset stream (see
/// [`serialize_assets`]) on success and NULL on failure. Release both
/// buffers with `anydoc_free(ptr, len)` using their reported lengths.
///
/// Requires anydoc's `document_to_markdown` to be part of its public API —
/// enabled by the local patch `patches/anydoc-render-pub.patch`.
#[unsafe(no_mangle)]
pub extern "C" fn anydoc_convert(
    input: *const u8,
    input_len: usize,
    fmt: *const u8,
    fmt_len: usize,
    out_md: *mut *mut u8,
    out_md_len: *mut usize,
    out_assets: *mut *mut u8,
    out_assets_len: *mut usize,
    out_code: *mut u32,
) {
    // SAFETY: caller-provided out slots; we always initialize them first.
    unsafe {
        *out_code = ERR_OK;
        *out_md_len = 0;
        *out_assets_len = 0;
        *out_md = ptr::null_mut();
        *out_assets = ptr::null_mut();
    }

    if input_len > 0 && input.is_null() {
        // SAFETY: out_md slot.
        unsafe { *out_md = fail("invalid input pointer", ERR_OTHER, out_md_len, out_code) }
        return;
    }
    // SAFETY: input_len bytes of valid memory when input_len > 0 (or NULL).
    let input_bytes = unsafe { slice::from_raw_parts(input, input_len) };

    let format = match resolve_format(fmt, fmt_len) {
        Ok(format) => format,
        Err((msg, code)) => {
            // SAFETY: out_md slot.
            unsafe { *out_md = fail(&msg, code, out_md_len, out_code) }
            return;
        }
    };

    let document = match anydoc::to_document(input_bytes, format) {
        Ok(document) => document,
        Err(err) => {
            // SAFETY: out_md slot.
            unsafe { *out_md = fail(&format!("{err}"), code_of(&err), out_md_len, out_code) }
            return;
        }
    };

    // Single parse now yields both outputs; `md` and `assets` share the same
    // `Document`, so `asset://<id>` in `md` and `Assets[id].ID` are identical.
    let md = anydoc::document_to_markdown(&document);
    let assets = serialize_assets(&document);

    // SAFETY: both out slots; store_bytes returns the freshly allocated
    // buffer (NULL when empty).
    let md_ptr = store_bytes(md.as_bytes(), out_md_len);
    let assets_ptr = store_bytes(&assets, out_assets_len);
    unsafe {
        *out_md = md_ptr;
        *out_assets = assets_ptr;
    }
    if md_ptr.is_null() && !md.is_empty() {
        // SAFETY: out_code slot.
        unsafe { *out_code = ERR_RESOURCE_LIMIT }
    }
}

/// Copy `bytes` into a fresh 8-byte-aligned allocation; `*out_len` receives
/// the length (0 when `bytes` is empty, and NULL is returned).
fn store_bytes(bytes: &[u8], out_len: *mut usize) -> *mut u8 {
    let n = bytes.len();
    let Ok(layout) = Layout::from_size_align(n, 8) else {
        // SAFETY: out slot write.
        unsafe { *out_len = 0 }
        return ptr::null_mut();
    };
    if n == 0 {
        // SAFETY: out slot write.
        unsafe { *out_len = 0 }
        return ptr::null_mut();
    }
    // SAFETY: layout is valid for n bytes with 8-byte alignment.
    let out = unsafe { alloc(layout) };
    // SAFETY: out has room for n bytes.
    if !out.is_null() {
        unsafe { ptr::copy_nonoverlapping(bytes.as_ptr(), out, n) };
    }
    // SAFETY: out slot write.
    unsafe { *out_len = n }
    out
}