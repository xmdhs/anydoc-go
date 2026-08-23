//! getrandom 本地 patch：去掉 wasm-unknown-unknown 上的 wasm-bindgen。
//!
//! 上游 getrandom 0.4.3 在 wasm32/64-unknown-unknown 上要求 feature `wasm_js`
//! 并通过 `wasm_js` backend 绑定 `globalThis.crypto.getRandomValues`（js-sys +
//! wasm-bindgen，大量 `__wbindgen_*` 导入）。本 patch 保持其他 target 原样，
//! 仅对 wasm-unknown-family 做确定性填充：全零（lopdf 仅为 trailer ID 去重，
//! 全零仍满足唯一性；若下游依赖真随机会在非 wasm 端走系统熵）。

#![no_std]

#[cfg(feature = "std")]
extern crate std;

mod error {
    use core::fmt;

    pub type RawOsError = i32;
    type NonZeroRawOsError = core::num::NonZeroI32;

    #[derive(Copy, Clone, Eq, PartialEq)]
    pub struct Error(NonZeroRawOsError);

    impl Error {
        pub const UNSUPPORTED: Error = Self::new_internal(0);
        pub const UNEXPECTED: Error = Self::new_internal(2);
        const INTERNAL_START: RawOsError = 1 << 16;
        const CUSTOM_START: RawOsError = 1 << 17;

        pub const fn new_custom(n: u16) -> Error {
            let code = Self::CUSTOM_START + (n as RawOsError);
            Error(unsafe { NonZeroRawOsError::new_unchecked(code) })
        }

        pub(crate) const fn new_internal(n: u16) -> Error {
            let code = Self::INTERNAL_START + (n as RawOsError);
            Error(unsafe { NonZeroRawOsError::new_unchecked(code) })
        }

        pub fn raw_os_error(self) -> Option<RawOsError> {
            let code = self.0.get();
            if code >= 0 {
                None
            } else {
                code.checked_neg()
            }
        }
    }

    impl core::error::Error for Error {}

    impl fmt::Debug for Error {
        fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
            f.debug_struct("Error")
                .field("code", &self.0.get())
                .finish()
        }
    }

    impl fmt::Display for Error {
        fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
            write!(f, "getrandom error {}", self.0.get())
        }
    }
}

pub use error::{Error, RawOsError};

use core::mem::MaybeUninit;

#[cfg(feature = "sys_rng")]
mod sys_rng {
    use crate::Error;
    use rand_core::{TryCryptoRng, TryRng};

    #[derive(Clone, Copy, Debug, Default)]
    pub struct SysRng;

    impl TryRng for SysRng {
        type Error = Error;

        #[inline]
        fn try_next_u32(&mut self) -> Result<u32, Error> {
            crate::u32()
        }

        #[inline]
        fn try_next_u64(&mut self) -> Result<u64, Error> {
            crate::u64()
        }

        #[inline]
        fn try_fill_bytes(&mut self, dest: &mut [u8]) -> Result<(), Error> {
            crate::fill(dest)
        }
    }

    impl TryCryptoRng for SysRng {}
}

#[cfg(feature = "sys_rng")]
pub use sys_rng::SysRng;

#[cfg(feature = "sys_rng")]
pub use rand_core;

#[cfg(all(target_family = "wasm", any(target_os = "unknown", target_os = "none")))]
mod wasm_stub {
    use super::{Error, MaybeUninit};

    pub fn fill_inner(dest: &mut [MaybeUninit<u8>]) -> Result<(), Error> {
        // 确定性填充：全零（不引 js-sys/wasm-bindgen）。
        for b in dest.iter_mut() {
            b.write(0);
        }
        Ok(())
    }

    pub fn inner_u32() -> Result<u32, Error> {
        Ok(0)
    }

    pub fn inner_u64() -> Result<u64, Error> {
        Ok(0)
    }
}

#[cfg(all(target_family = "wasm", any(target_os = "unknown", target_os = "none")))]
pub use wasm_stub::{fill_inner, inner_u32, inner_u64};

#[cfg(not(all(target_family = "wasm", any(target_os = "unknown", target_os = "none"))))]
mod not_wasm {
    use super::{Error, MaybeUninit};

    pub fn fill_inner(_dest: &mut [MaybeUninit<u8>]) -> Result<(), Error> {
        Err(Error::UNSUPPORTED)
    }

    pub fn inner_u32() -> Result<u32, Error> {
        Err(Error::UNSUPPORTED)
    }

    pub fn inner_u64() -> Result<u64, Error> {
        Err(Error::UNSUPPORTED)
    }
}

#[cfg(not(all(target_family = "wasm", any(target_os = "unknown", target_os = "none"))))]
pub use not_wasm::{fill_inner, inner_u32, inner_u64};

/// 与上游同签名薄封装。
#[inline]
pub fn fill(dest: &mut [u8]) -> Result<(), Error> {
    fill_uninit(unsafe {
        core::slice::from_raw_parts_mut(dest.as_mut_ptr() as *mut MaybeUninit<u8>, dest.len())
    })?;
    Ok(())
}

#[inline]
pub fn fill_uninit(dest: &mut [MaybeUninit<u8>]) -> Result<&mut [u8], Error> {
    if !dest.is_empty() {
        fill_inner(dest)?;
    }
    Ok(unsafe { core::slice::from_raw_parts_mut(dest.as_mut_ptr() as *mut u8, dest.len()) })
}

#[inline]
pub fn u32() -> Result<u32, Error> {
    inner_u32()
}

#[inline]
pub fn u64() -> Result<u64, Error> {
    inner_u64()
}
