# anydoc-go 项目说明

anydoc（Rust 文档→Markdown 库）经 wasm2go 翻译成纯标准库 Go，并提供 Go 文件转换 CLI。
本文件汇总项目的架构、构建路径、本机约束与已验证事实，供后续维护直接引用。

## 一句话架构

```
cabi/src/lib.rs（C ABI 壳） ──cargo──▶ anydoc_cabi.wasm ──wasm2go──▶ .go + .dat ──go build──▶ bin/anydoc
```

依赖关系：cabi（Rust cdylib，path 依赖 anydoc 核心库，无 wasm-bindgen/Js 依赖） →
wasm2go（CLI 翻译器）→ split_gen.py（按函数拆分生成文件）→ `go build` → 单二进制。

## 目录

| 路径 | 内容 |
|---|---|
| `cabi/` | Rust cdylib：`anydoc_alloc` / `anydoc_free` / `anydoc_to_markdown` |
| `cli/` | Go CLI（main.go + split 生成文件 `anydoc_gen_*.go` + `anydoc.wasm.dat`） |
| `target/` | cargo 产物（`wasm32-unknown-unknown/release/anydoc_cabi.wasm`） |
| `tools/` | split_gen.py（把 wasm2go 单文件按顶层函数拆块）；re_split.py（换粒度用） |
| `bin/` | wasm2go 翻译器 + 最终 anydoc 二进制 |
| `build.sh` | 构建管线（wasm / cli / build 三个子命令） |

## C ABI 协议（cabi/src/lib.rs）

- `anydoc_alloc(size) -> ptr`：8 字节对齐全局分配器；size==0 返回 NULL。
- `anydoc_free(ptr, size)`：释放（size 必须与分配时一致）。
- `anydoc_to_markdown(input, input_len, fmt, fmt_len, out_len, out_code) -> ptr`：
  - `fmt_len==0` 自动检测格式（`Format::from_bytes`）；否则按扩展名（`-f csv` 等）。
  - 成功：`out_code=0`，返回 Markdown UTF-8 文本，长度在 `out_len`。
  - 失败：`out_code` = 1 unsupported / 2 malformed / 3 encrypted / 4 resourceLimit /
    5 missingPart / 6 io / 7 other，返回错误消息文本；用 `anydoc_free(ptr, *out_len)` 释放。
- 所有指针是 wasm 线性内存地址（wasm32 上是 32 位值），host 通过导出 memory
  （`m.Xmemory().Slice()` 拿 `*[]byte`）读写。

## 关键设计决策（不要再推翻）

1. **C ABI 壳，不走 wasm-bindgen**：wasm-bindgen 产物带 `__wbg_*` JS imports + 复杂
   ABI 编码，wasm2go 翻译后需要 Go 侧实现 JS 语义 shim。C ABI 后模块几乎零依赖。
2. **getRandomValues 残留**：依赖树中的 getrandom（wasm32 路径）拉进
   `__wbindgen_*` imports；`cli/main.go` 的 `hostEnv` 用 crypto/rand 实现
   `getRandomValues`，其余是 stub（`extrefEnv` 同）。接口名带 CRC 后缀，勿改。
3. **PDF feature 裁剪（重要）**：pdf-inspector 编进 wasm 后产生 **17 万行单个函数**，
   wasm 7.5MB → Go 源码 36MB，在内存不足时（本机 2 核 / 3.5G / swap 满）codegen
   卡死数小时。故给 anydoc 加了 `pdf` feature（`anydoc/Cargo.toml` `[features] pdf`，
   `src/formats/mod.rs` 与 `src/lib.rs` 的 `#[cfg(feature = "pdf")]` gate），
   cabi 以 `default-features = false` 构建 → wasm 2MB、Go 8MB。需要 PDF 时打开
   feature 并用大内存机器。
4. **编译参数**：生成文件必须拆成小块（45 函数/文件）再 `go build -p 2
   -gcflags=all="-N -l"`；更好地机器上可去掉 `-N -l`（默认优化、运行更快）并
   调大 `-p`，拆不拆均可（拆分是纯并行粒度问题，不影响产物）。

## 本机工具链与环境（重要）

- /tmp 只有 1.6G、home 只读**：所有工具链/缓存放 /mnt/hdd/frontend/git：
  `.rustup`、`.cargo`、`zig/`（zig 0.16 替代 cc/gcc/clang 的 shim 在 `tools/`）、
  `.go-cache`、`.gopath`、`.zig-cache`（环境变量见 build.sh）。
- 无 make / sudo：用 `./build.sh [wasm|cli|build]`。
- 网络：crates.io、goproxy.cn 可用；GitHub API 慢/不通（勿依赖其下载）。
- 构建命令示例：
  ```
  RUSTUP_HOME=/mnt/hdd/frontend/git/.rustup CARGO_HOME=.../../.cargo \
  CARGO_TARGET_DIR=.../target ZIG_GLOBAL_CACHE_DIR=.../.zig-cache \
  ZIG_LOCAL_CACHE_DIR=.../libzig-cache ./build.sh wasm
  ```
  每次改 cabi 或 anydoc 源码后：`./build.sh wasm && ./build.sh cli`。

## 验证方式（回归）

- 样本在 `anydoc/tests/fixtures/`（docx/csv/pptx/pdf/epub/ods/odt/ppt/doc/rtf/…），
  期望输出在 `anydoc/tests/snapshots/*.snap`（insta，正文在第二个 `---` 之后）。
- 快速对照：`bin/anydoc fixtures/docx/text.docx` 与 snapshots 正文逐字比对（已验证
  docx/csv/doc/rtf 一致）。
- 错误路径（pdf 无 feature → unsupported；encrypted → 3；缺文件）应退出码 1。
- 单文件 stdout / `-o` / 多文件（各输出 `<basename>.md` 到 cwd）三种模式都过。

## 已知限制 & 后续

- PDF 转换未编译（有 feature 时用大内存机器重建即可，编译时间会明显上升）。
- `panic = "abort"`：wasm 内部 panic 直接终止进程。
- 运行内存：线性内存初始 3MB + 转换分配（rust 分配器）可增长；测试 docx 约 85-120ms。
- build.sh 里 `-p 2 -N -l` 是本机妥协；换机器按上文调整。