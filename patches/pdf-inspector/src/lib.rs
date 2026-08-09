//! pdf-inspector 的本地 stub。
//!
//! 通过 cabi 的 `[patch.crates-io]` 替换真实实现，把 pdf 解析（以及它的
//! 129 个依赖 crate）从依赖图中剔除：wasm 变小、生成 Go 源码变小。
//! anydoc 只使用 `PdfError` / `process_pdf_mem` / `PdfProcessResult` 三处
//! 公开接口（见 third-party/anydoc/src/formats/pdf.rs）；接口漂移会在
//! anydoc 编译时暴露（map_error 是穷尽匹配），无需手工跟进。
//!
//! 行为：`process_pdf_mem` 返回 `markdown = None`，anydoc 据此报
//! Unsupported（"PDF has no extractable text … OCR is required"），与旧版
//! "pdf 未编入" 的语义一致（CLI 报 unsupported）。

use std::error;
use std::fmt;
use std::io;

/// 检测到的 PDF 类型。anydoc 只对它做 `{:?}` 打印，单变体即可满足编译。
#[derive(Debug)]
pub enum PdfType {
    /// PDF has extractable text (Tj/TJ operators found)
    TextBased,
}

/// 高层 PDF 处理结果。字段即 anydoc 访问的子集。
#[derive(Debug)]
pub struct PdfProcessResult {
    /// The detected PDF type.
    pub pdf_type: PdfType,
    /// Markdown output (populated in full processing mode, `None` otherwise).
    pub markdown: Option<String>,
    /// Page count.
    pub page_count: u32,
    /// 1-indexed page numbers that need OCR.
    pub pages_needing_ocr: Vec<u32>,
    /// `true` when broken font encodings are detected.
    pub has_encoding_issues: bool,
}

/// 与上游同形的错误枚举：anydoc 的 `map_error` 穷尽匹配这 5 个变体，
/// 不能增删变体（增了 anydoc 编译失败，删了语义失真）。
#[derive(Debug)]
pub enum PdfError {
    /// IO error.
    Io(io::Error),
    /// PDF parsing error.
    Parse(String),
    /// PDF is encrypted.
    Encrypted,
    /// Invalid PDF structure.
    InvalidStructure,
    /// Not a PDF.
    NotAPdf(String),
}

impl fmt::Display for PdfError {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        match self {
            Self::Io(e) => write!(f, "IO error: {e}"),
            Self::Parse(d) => write!(f, "PDF parsing error: {d}"),
            Self::Encrypted => f.write_str("PDF is encrypted"),
            Self::InvalidStructure => f.write_str("Invalid PDF structure"),
            Self::NotAPdf(d) => write!(f, "Not a PDF: {d}"),
        }
    }
}

impl error::Error for PdfError {}

impl From<io::Error> for PdfError {
    fn from(e: io::Error) -> Self {
        Self::Io(e)
    }
}

/// 未编入 pdf 解析：返回空结果让 anydoc 走 "no extractable text"
/// → Unsupported（转换路径报 unsupported，与剔除前旧版一致）。
pub fn process_pdf_mem(_buffer: &[u8]) -> Result<PdfProcessResult, PdfError> {
    Ok(PdfProcessResult {
        pdf_type: PdfType::TextBased,
        markdown: None,
        page_count: 0,
        pages_needing_ocr: vec![],
        has_encoding_issues: false,
    })
}
