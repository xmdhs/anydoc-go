//! Patch 版 web-time：直接用 [`std::time`] 替代。
//!
//! 上游 web-time 在 wasm32-unknown-unknown 上依赖 js-sys/wasm-bindgen
//! （浏览器时钟），会把整个 wasm-bindgen 编译进来（生成上千个 describe
//! 导出桩）。anydoc-go/cabi 只做纯字节转换，不需要浏览器时钟：
//! [`std::time`] 在 wasm32-unknown-unknown 上提供完整的算术/比较能力
//! （时间源为固定 mock，仅影响写入时间戳的语义，而转换路径不写）。
//! API 面与 web-time 1.1.0 非 wasm 分支一致。

pub use std::time::*;