# anydoc-go 项目说明

anydoc（Rust 文档→Markdown 库）经 wasm2go 翻译成纯标准库 Go（core/），lib/
封装成 Go 库 API，CLI（main.go）调用库。本文件是维护速查；用户文档见根 README.md。

## 架构

```
cabi/src/lib.rs（C ABI 壳，path 依赖 third-party/anydoc）
   ──cargo（wasm32-unknown-unknown）──▶ target/<…>/anydoc_cabi.wasm（2MB）
   ──wasm2go -pkg core -embed──▶ core/*.go + core/anydoc.wasm.dat（随仓库提交）
   ──lib/（Converter 封装，非线程安全）──▶ bin/anydoc（CLI 入口 main.go）
   build.sh 串联：fetch / wasm / cli / test / build
```

## 目录

| 路径 | 内容 |
|---|---|
| `cabi/`   | Rust cdylib：anydoc_alloc / anydoc_free / anydoc_to_markdown |
| `third-party/` | build.sh fetch 拉取的 anydoc、wasm2go（`.gitignore` 排除） |
| `core/`   | wasm2go 生成代码（package core，生成物但提交） |
| `lib/`    | 对外 Go 库 `package anydoc`：ConvertFile/ConvertBytes/ConvertFileTo + ConvertError |
| `main.go` | CLI（glob 批量、默认同目录 `<name>.md`、`-s` stdout、`-f` 强制格式） |
| `testdata/` | fixtures（anydoc 样本）+ expected（快照期望） |
| `tools/`  | split_gen.py（wasm2go 产物按函数拆分）、gen_expected.py |
| `bin/` `target/` | 构建产物（gitignore） |

## C ABI 协议（cabi/src/lib.rs）

- `anydoc_alloc(size) -> ptr`：8 字节对齐分配器；size==0 返回 NULL。
- `anydoc_free(ptr, size)`：释放（size 与分配一致）。
- `anydoc_to_markdown(input,len, fmt,fmt_len, out_len,out_code) -> ptr`：
  - `fmt_len==0` 自动检测（core Format::from_bytes）；否则按扩展名（-f）。
  - 成功：`out_code=0`，返回 Markdown UTF-8，长度在 `out_len`。
  - 失败：`out_code` 1 unsupported / 2 malformed / 3 encrypted / 4 resourceLimit /
    5 missingPart / 6 io / 7 other，内容为错误消息；`anydoc_free(ptr,*out_len)` 释放。
- 指针 = wasm 线性内存地址（i32）；host 经 `m.Xmemory().Slice()`（`*[]byte`）读写。
- Go 侧格式优先级：`-f` > 内容检测 > 路径扩展名兜底（CSV 等无签名格式依赖它）。

## 关键设计决策（不要再推翻）

1. **C ABI 壳，不走 wasm-bindgen**：避免 `__wbg_*` JS imports 与复杂 ABI。
2. **PDF feature 裁剪**：pdf-inspector 编入后出现 17 万行单函数，wasm 7.5MB→Go 36MB，
   本机（2 核 / swap 满）codegen 无法完成。anydoc 加了 `pdf` feature（third-party/anydoc：
   Cargo.toml `[features] pdf` + lib.rs/formats/mod.rs 三处 cfg gate），cabi
   `default-features = false` 构建。要 PDF：开 default-features + 大内存机器。
3. **生成代码编译**：单文件 36MB 不可编，必须 split_gen.py 拆（45 函数/文件）+
   `go build -p 2 -gcflags=all="-N -l"`。好机器可去掉 `-N -l`、调大 `-p`（拆分只是并行度）。
4. **wasm-bindgen 残留**：web-time 等拉进的 imports 只有 `__wbindgen_throw`/
   `__wbindgen_describe`（stub 在 lib/wasm_env.go）。接口名单以 core 内为准；勿手改。

## 本机工具链（重要）

- /tmp 1.6G、home 只读；工具链放父级：`.rustup` `  .cargo` `zig/` + `tools/`（cc 库 shim，
  实际是 `/mnt/hdd/frontend/git/tools`，不在本仓库）、`.go-cache`、`.gopath`。
- 无 make/sudo；`.git/config` 只读（git 要用 `-c user.name=xmdhs -c user.email=…` 与
  URL 参数 push）。
- 网络：crates.io / goproxy.cn 可用；GitHub API 慢（push 走 ssh 正常）。
- 改 cabi 或 anydoc 后：`./build.sh wasm && ./build.sh cli`。

## 验证

- `./build.sh test`（go test ./...）：6 格式×快照逐字节、复用无状态、错误路径
  （encrypted/malformed/unrecognized/csv/pd + 并发多实例）+ lib 扩展名兜底。
- 新增格式样本：`cp third-party/anydoc/tests/fixtures/<fmt>/* tests static/fixtures/<fmt>/`
  + `python3 tools/gen_expected.py third-party/anydoc/tests/snapshots testdata/expected`。

## 已知限制

- PDF 未编入（见上）；`panic = "abort"`；模块内存 3MB 起可 grow（maxMem 64MB）；
  单次 docx 约 85ms（本机）。CLI 默认写同目录 `.md`，`-s` 走 stdout。