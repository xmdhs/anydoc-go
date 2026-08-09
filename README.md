# anydoc-go: anydoc via wasm2go + a Go CLI

anydoc（Rust 文档→Markdown 库）的 C-ABI wasm 用 [wasm2go](https://github.com/ncruces/wasm2go)
翻译成纯标准库 Go 源码，再由一个 Go CLI 把任意文档
（docx/pptx/xlsx/ods/odp/epub/doc/odt/rtf/csv）转成 GitHub-Flavored Markdown。

独立仓库：第三方代码只放在 `third-party/`（不提交），其余全部自足。

```
cabi/            Rust cdylib：C ABI 导出 anydoc 核心库（避开 wasm-bindgen JS 依赖）
core/            wasm2go 生成物（package core，build.sh cli 生成，随仓库提交）
third-party/     build.sh fetch 拉取的第三方仓库：anydoc、wasm2go（git 工作树）
tools/           split_gen.py（wasm2go 产物按函数拆分）等辅助脚本
bin/             wasm2go 翻译器 + 最终 anydoc 二进制
target/          cargo 编译目录（.gitignore 排除）
testdata/        单元测试样本（fixtures/ 复制自 anydoc + expected/ 快照期望）
main.go          唯一手写 Go 入口（package main）
build.sh         构建管线：fetch / wasm / cli / test / build
```

## 架构

```
main.go + cabi/src/lib.rs
   ├── anydoc_alloc(size) -> ptr        （8 字节对齐分配，Rust 全局分配器）
   ├── anydoc_free(ptr, size)
   └── anydoc_to_markdown(input,len, fmt,fmt_len, out_len,out_code) -> ptr
        fmt==NULL/len==0 自动检测；CSV 无签名必须显式 -f
        out_code: 0=成功 / 1 unsupported / 2 malformed / 3 encrypted /
                  4 resourceLimit / 5 missingPart / 6 io / 7 other
        success → 返回 Markdown UTF-8；失败 → 错误消息；anydoc_free 释放
        │ cargo build --target wasm32-unknown-unknown
        ▼
anydoc_cabi.wasm（2MB，无 wasm-bindgen exports 之外零依赖）
        │ wasm2go -pkg main -embed
        ▼
anydoc_gen_*.go（~8MB，45 函数/文件）+ anydoc.wasm.dat
        │ go build -N -l（内存受限机器）
        ▼
bin/anydoc
```

- **wasm-bindgen 残留**：依赖树里 web-time 等仍引用 wasm-bindgen 的
  `__wbindgen_*` imports（现在只剩 throw/describe 两个 stub 接口），
  main.go 的 `hostEnv` 提供惰性实现；转换路径从不调用它们。
  若重新生成后接口变了，以 `anydoc_gen_001.go` 里的 `type X__wbindgen*` 为准。
- **PDF 被裁剪**（重要）：pdf-inspector 编进 wasm 后产生 **17 万行单函数**，
  Go 源码达 36MB，无内存的 2 核机器上无法编译。因此在 anydoc 页加了
  `pdf` feature 开关（构造 Cargo.toml + src 三处 cfg gate），cabi 用
  `default-features = false` 构建。需要 PDF 时打开 cabi 依赖的 default-features
  并用大内存机器重建即可。

## 构建

本机约束：/tmp 仅 1.6G、home 只读、2 核且 swap 常满。工具链与缓存全部放
父目录（.rustup/.cargo/.go-cache/.zig-cache）或本目录（target/），见 build.sh。

```bash
./build.sh fetch    # 首次 & 更新：拉取 third-party/anydoc、wasm2go（git clone/fetch）
                    #   URL 可 ANPYOC_REPO / WASM2GO_REPO 覆盖
./build.sh wasm     # cargo → anydoc_cabi.wasm
./build.sh cli      # wasm2go → 拆分 → go build
./build.sh test     # 单元测试（见「验证」）
./build.sh build    # wasm + cli 全流程
```

生成物编译参数说明：36MB 单文件在本机 codegen 无法完成（内存），拆分后
`go build -p 2 -gcflags=all="-N -l"`（~2 分钟）。大内存多核机器上可去掉
`-N -l`、调大 `-p`，甚至不拆（拆分是纯并行度问题，不影响产物）。

## 用法

```bash
bin/anydoc report.docx            # → 同目录 report.docx.md（默认）
bin/anydoc "docs/*.docx" "*.csv"   # glob 批量（模式含 * 时用引号包裹）
bin/anydoc -s report.docx          # 输出到 stdout（可多个输入）
bin/anydoc -o out.md report.docx   # 指定单个输出文件
bin/anydoc -o - a.docx             # 与 -s 等价
bin/anydoc -f docx weird.bin       # 强制定制解析
```

格式识别优先级：`-f` 显式 > 文件内容自动检测 > **路径扩展名兜底**（CSV 等
无签名格式直接转换，无需再传 `-f`）。字节流 API（`ConvertBytes`）没有路径
概念，仍需显式格式。

## Go 库

`lib/`（包名 `anydoc`）封装转换核心。**Converter 非线程安全**——同一实例
同一时间只能被一个 goroutine 使用（wasm 模块是共享线性内存/栈指针的单例）；
并发转换用每 goroutine 独立实例（内存各自独立，天然并行）：

```go
import (
    "fmt"
    "log"

    "anydoc-go/lib"
)

conv := lib.NewConverter() // 创建成本约几十毫秒，长期复用

md, err := conv.ConvertFile("report.docx", "") // 自动检测格式
md, err = conv.ConvertBytes(data, "csv")        // 无签名格式需显式扩展名
err = conv.ConvertFileTo("a.docx", "", "a.md")  // 或直接写出 md

if cerr, ok := err.(*lib.ConvertError); ok {
    log.Fatalf("%s: %s", cerr.Code, cerr.Msg) // Code: unsupported/malformed/...
}

// 并发：每 goroutine 一个实例
md, err = func() (string, error) {
    c := lib.NewConverter()
    defer ... // 无资源释放，忽略也可
    return c.ConvertFile("report.docx", "")
}()
```

导入路径随 module 名（当前为本地 `anydoc-go`，发布时改为真实仓库路径）。

## 验证

单元测试（`./build.sh test`，标准 go test）：

- **TestConvertMatchesExpected**：6 种格式（docx/csv/doc/pptx/epub/ods）用
  third-party/anydoc 的样本（testdata/fixtures/）转换，与官方快照期望
  （testdata/expected/，`tools/gen_expected.py` 从 anydoc 的 insta 快照生成）
  逐字节比对；
- **TestReuseModuleIsStateless**：同一模块连续转换结果一致；
- **TestErrors**：encrypted / malformed / 未识别格式 / CSV 需要 -f /
  PDF 未编译（开 pdf feature 后需更新该用例）。

新增格式样本：`cp third-party/anydoc/tests/fixtures/<fmt>/*.xxx testdata/fixtures/<fmt>/`
并跑 `python3 tools/gen_expected.py third-party/anydoc/tests/snapshots testdata/expected`。

## 已知限制

- PDF 转换未编入（见上）；panic = "abort"（wasm 内部崩溃直接终止进程）；
- 运行期线性内存 3MB 起、可 grow 至 64MB（maxMem），单次 docx 转换约 85ms；
- main.go 的 `hostEnv` stub 以生成接口为准，不要手工注释（编译即校验）。