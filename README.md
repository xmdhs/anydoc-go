# anydoc-go: anydoc via wasm2go + a Go CLI

anydoc（Rust 文档→Markdown 库）的 C-ABI wasm 用 [wasm2go](https://github.com/ncruces/wasm2go)
翻译成纯标准库 Go 源码，再由一个 Go CLI 把任意文档
（docx/pptx/xlsx/ods/odp/epub/doc/odt/rtf/csv；pdf 被 stub 剔除，
见「架构」）转成 GitHub-Flavored Markdown。

独立仓库：第三方代码只放在 `third-party/`（不提交），其余全部自足。

```
cabi/            Rust cdylib：C ABI 导出 anydoc 核心库（避开 wasm-bindgen JS 依赖）
core/            wasm2go 生成产物（package core，build.sh cli 生成，随仓库提交）
third-party/     build.sh fetch 拉取的第三方仓库：anydoc、wasm2go（git 工作树）
patches/         本地 crate 替换：web-time（去 wasm-bindgen 依赖）、pdf-inspector stub
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
        │ cargo build --release --target wasm32-unknown-unknown（zig cc 提供宿主链接器）
        ▼
anydoc_cabi.wasm（≈1.7MB；pdf-inspector 被本地 stub patch 剔除）
        │ wasm2go -unsafe -pkg core -embed
        ▼
anydoc_gen_*.go（≈9.7MB，1007 函数拆分 24 文件）+ anydoc.wasm.dat
        │ go build -p 2 -N -l
        ▼
bin/anydoc
```

- **无 wasm-bindgen 残留**：pdf 剔除（见下）把依赖树里的 web-time/getrandom
  一并带走了，`core.New()` 无宿主接口参数，wasm 是纯 anydoc 导出。
- **PDF 被 stub 剔除**（重要）：anydoc 新版把 pdf-inspector 移入普通依赖
  （非 optional feature），`default-features = false` 裁不掉。cabi 用
  `[patch.crates-io]` 把它替换为本地 stub（patches/pdf-inspector，85 行），
  依赖闭包 129 个 crate 全部消失：wasm 7.2→1.7MB、生成 Go 源码 37→9.7MB。
  pdf 输入报 unsupported（"no extractable text"）。stub 接口即 anydoc
  用到的 `PdfError`/`process_pdf_mem`/`PdfProcessResult`（map_error 穷尽
  匹配，接口漂移会编译失败提醒）。恢复 pdf：删掉 cabi/Cargo.toml 里的
  patch 行并重建即可。

## 构建

C 工具链用 zig：`.zig/bin/cc` 是 `zig cc` 的 wrapper，
作为 rustc 的宿主链接器（本机没有 gcc/clang）。

```bash
./build.sh fetch    # 首次 & 更新：拉取 third-party/anydoc、wasm2go（git clone/fetch）
                    #   anydoc 默认取语义排序的最新 tag（如 v0.1.7），
                    #   wasm2go 默认 main；FETCH_REF 可统一覆盖
./build.sh wasm     # cargo → anydoc_cabi.wasm
./build.sh cli      # wasm2go -unsafe → 拆分 → go build
./build.sh test     # 单元测试（见「验证」）
./build.sh build    # wasm + cli 全流程
```

### `-unsafe` 与性能

wasm2go 默认生成用 `encoding/binary` 边界检查解码的代码；`-unsafe` 启用
helpers 里的指针直读（load16/32/64），生成代码仍保持边界检查语义
（wasm2go 文档：加入 unsafe 后所有内存访问依旧有界检查），转换路径不涉及
其他 unsafe 危险行为（lib 的 Converter 仍非线程安全）。

本机实测（`go test -bench .`，`-gcflags=all="-N -l"`，中位数，各 6 次）：

| 输入 | 基线 | -unsafe | 加速 |
| --- | ---: | ---: | ---: |
| sheet.csv | 1.24 ms | 1.11 ms | 1.12x |
| book.epub | 3.25 ms | 2.61 ms | 1.25x |
| pres.pptx | 10.7 ms | 8.31 ms | 1.29x |
| sheet.ods | 5.33 ms | 4.12 ms | 1.29x |
| text.docx | 13.7 ms | 10.6 ms | 1.29x |
| text.doc | 2.75 ms | 2.07 ms | 1.33x |
| 2MB CSV | 1208 ms | 852 ms | 1.42x |

输入越大收益越大（内存访问占比高）：小样本约 1.1-1.3×，2MB CSV 约 1.4×。
表中数值为含 pdf 版生成物（pdf 剔除前）的实测；pdf 剔除不影响转换路径，
相对加速比结论不变。复测：`go test -bench . -benchtime 10x -count 3
-run 'Benchmark'`——注意 `-run '^$'` 会把子基准（BenchmarkConvert/xxx）
一并过滤掉；对比时两个版本用相同编译参数。

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

## CI（GitHub Actions）

`.github/workflows/build-anydoc.yml` 在 runner 上重放本机流程（与 build.sh
一致：zig C 工具链 → cargo wasm → wasm2go -unsafe → go build → go test →
CLI 冒烟），并上传 `bin/anydoc` artifact：

- `push` 到 main：构建 + 测试 + 归档产物；
- `workflow_dispatch` 且填写 version：额外创建同名 tag 的 GitHub Release
  （需仓库权限；版本随意填写如 `1.0.0`）。

产物可直接在 workflow 运行页的 Artifacts 下载，7 天保留。

## 验证

单元测试（`./build.sh test`，标准 go test）：

- **TestConvertMatchesExpected**：6 种格式（docx/csv/doc/pptx/epub/ods）用
  third-party/anydoc 的样本（testdata/fixtures/）转换，与官方快照期望
  （testdata/expected/，`tools/gen_expected.py` 从 anydoc 的 insta 快照生成）
  逐字节比对；
- **TestReuseModuleIsStateless**：同一模块连续转换结果一致；
- **TestErrors**：encrypted / malformed / 未识别格式 / CSV 需要 -f /
  PDF 报 unsupported（stub 剔除后不解析）。

新增格式样本：`cp third-party/anydoc/tests/fixtures/<fmt>/*.xxx testdata/fixtures/<fmt>/`
并跑 `python3 tools/gen_expected.py third-party/anydoc/tests/snapshots testdata/expected`。

## 已知限制

- panic = "abort"（wasm 内部崩溃直接终止进程）；
- 运行期线性内存 3MB 起、可 grow 至 64MB（maxMem），单次 docx 转换约 10ms；
- PDF 不解析（stub 剔除，报 unsupported），见「架构」；
- 生成物接口以 `core` 包为准（如 `core.New` 签名、`Xanydoc_*` 导出名），
  不要手工注释或改名（编译即校验）。