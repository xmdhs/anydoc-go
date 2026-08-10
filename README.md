# anydoc-go（goccy 分支）: anydoc via goccy/wasm2go + a Go CLI

anydoc（Rust 文档→Markdown 库）的 C-ABI wasm 用 [goccy/wasm2go]（AOT 编译器：
生成 amd64/arm64 汇编 + 纯 Go 回退）编译成 Go 源码，再由一个 Go CLI 把任意
文档（docx/pptx/xlsx/ods/odp/epub/doc/odt/rtf/csv；pdf 被 stub 剔除，见
「架构」）转成 GitHub-Flavored Markdown。

> 与 main 分支（[ncruces/wasm2go] 纯 Go 翻译）的区别：goccy 把 wasm 真正编译
> 成宿主汇编，转换快约 1.6-2×，二进制约 5/8 体积；局限是生成物约 64MB、
> 一次生成耗时长（≈6-7 分钟），且 goccy 至今没有 tag 只跟 main。详细对比见
> 「工具链对比」。

独立仓库：第三方代码只放在 `third-party/`（不提交），其余全部自足。

```
cabi/            Rust cdylib：C ABI 导出 anydoc 核心库（避开 wasm-bindgen JS 依赖）
core/            goccy-wasm2go 生成产物（core.go + base/ + p0/ + p1/ + data.bin，
                 随仓库提交；amd64/arm64 双架构 asm + 纯 Go 回退）
third-party/     build.sh fetch 拉取的第三方仓库：anydoc、goccy-wasm2go（git 工作树）
patches/         本地 crate 替换：web-time（去 wasm-bindgen 依赖）、pdf-inspector stub
tools/           辅助脚本（split_gen.py 已不再使用——goccy 自带多包拆分）
bin/             goccy-wasm2go 翻译器 + 最终 anydoc 二进制
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
        │ goccy-wasm2go -pkg core -import anydoc-go/core -out-dir core
        ▼
core/core.go + base/ + p0/ + p1/ + data.bin（≈64MB，984 函数）
        │ go build -p 2 -trimpath -ldflags="-s -w"
        ▼
bin/anydoc（≈8MB static ELF）
```

- **SIMD 显式关闭**：`wasm_step` 用 `RUSTFLAGS="-C target-feature=-simd128"`。
  文档转换是标量/字符串/XML/zlib 主导，实测 SIMD 开启反而慢 3-4 倍（见
  「工具链对比」），goccy 虽支持 SIMD wasm 但本负载用不上。
- **无 wasm-bindgen 残留**：pdf 剔除（见下）把依赖树里的 web-time/getrandom
  一并带走了，`core.New()` 无宿主接口参数，wasm 是纯 anydoc 导出。
- **PDF 被 stub 剔除**（重要）：anydoc 新版把 pdf-inspector 移入普通依赖
  （非 optional feature），`default-features = false` 裁不掉。cabi 用
  `[patch.crates-io]` 把它替换为本地 stub（patches/pdf-inspector，85 行），
  依赖闭包 129 个 crate 全部消失：wasm 7.2→1.7MB、生成物 37→64MB 中
  pdf 相关代码为零。pdf 输入报 unsupported（"no extractable text"）。
  stub 接口即 anydoc 用到的 `PdfError`/`process_pdf_mem`/`PdfProcessResult`
  （map_error 穷尽匹配，接口漂移会编译失败提醒）。恢复 pdf：删掉
  cabi/Cargo.toml 里的 patch 行并重建即可。

## 构建

C 工具链：本机用 zig（`.zig/bin/cc` 是 `zig cc` 的 wrapper，作为 rustc
的宿主链接器——本机没有 gcc/clang）。有系统 cc 的环境直接用系统工具链，
无需 zig。goccy 翻译器内部会跑一次 `go build` 抓取 asm，需要
`GOCACHE`/`GOPATH` 可写（build.sh 已默认指到项目内缓存）。

```bash
./build.sh fetch    # 首次 & 更新：拉取 third-party/anydoc、third-party/goccy-wasm2go
                    #   anydoc 默认取语义排序的最新 tag；goccy 跟 main（无 tag）
./build.sh wasm     # cargo → anydoc_cabi.wasm（-simd128）
./build.sh cli      # goccy-wasm2go → core/ → go build（生成约 6-7 分钟，仅 wasm 变更后重跑）
./build.sh test     # 单元测试（见「验证」）
./build.sh build    # wasm + cli 全流程（默认）
```

注意：`cli` 步骤的 goccy 翻译是本分支耗时大头（2 核机约 6-7 分钟，内部要
为 984 个函数跑 SSA + 抓 amd64/arm64 asm）；core/ 随仓库提交，日常改
main.go/lib 只跑 `go build` 即可，无需重新翻译。

## 工具链对比（why goccy）

同一份 wasm 分别在 ncruces（main 分支）与 goccy（本分支）下实测，本机
（AMD GX-215JJ，2 核）中位耗时（ms）：

| 输入 | ncruces | goccy（本分支） |
| --- | ---: | ---: |
| book.epub | 2.41 | 1.17 |
| pres.pptx | 6.17 | 3.36 |
| sheet.ods | 3.19 | 1.71 |
| text.docx | 7.08 | 4.65 |
| 2MB CSV | 623 | 321 |

结论：goccy 快约 1.3-2×（CSV/大文件约 1.9×），最终二进制 7.8MB vs 12.5MB。
代价：① 生成物大（64MB vs 9.7MB）且翻译时间长；② 无发布 tag；③ 做不了
ncruces 的 `-embed` 单文件（多包结构）；④ SIMD 开启反而慢——本分支已在
`wasm_step` 显式关闭。

## 用法

```bash
bin/anydoc report.docx              # → 同目录 report.docx.md（默认）
bin/anydoc "docs/*.docx" "*.csv"    # glob 批量（模式含 * 时用引号包裹）
bin/anydoc -s report.docx           # 输出到 stdout（可多个输入）
bin/anydoc -o out.md report.docx    # 指定单个输出文件
bin/anydoc -o - a.docx              # 与 -s 等价
bin/anydoc -f docx weird.bin          # 强制定制解析
bin/anydoc --imgs "docs/*.docx"        # 图片导出到每个输入旁的 imgs/
```

**内嵌图片**（docx/pptx 等容器内的图片字节不会出现在 markdown 里——本地
patch 渲染为 `![alt](asset://<id>)` 占位，`<id>` 即资产下标）：
- **默认不导出**：md 里保留 `asset://` 占位；
- `--imgs`：`<stem>-<id>.<ext>` 写入**输入文件同目录的 `imgs/`**，md
  引用替换为相对 `imgs/` 路径；各输入独立目录无跨输入覆盖，重复转换
  幂等；
- 扩展名由 media type 派生，其余一律 `bin`。
- 库 API 对应 `Converter.ExtractAssets`（返回 `ID/MediaType/Bytes`）。

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

conv := lib.NewConverter() // 创建成本约 1ms 级，长期复用

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
  third-party/anydoc 的样本（testdata/fixtures/）转换，与快照期望
  （testdata/expected/，`tools/gen_expected.py` 从 anydoc 的 insta 快照生成，
  **例外**：含内嵌图片的样例在本地 patch 下多出 `asset://` 占位，
  与官方快照不同，已手工更新 docx/doc/epub 三个文件）逐字节比对；
- **TestExtractAssets**：手工构造的 in-image.docx（1×1 PNG）→ markdown
  含 `asset://0` 占位、`ExtractAssets` 取回原始 PNG 字节；CLI 集成用例
  验证 `--imgs` 落盘 + 引用重写；
- **TestReuseModuleIsStateless**：同一模块连续转换结果一致；
- **TestErrors**：encrypted / malformed / 未识别格式 / CSV 需要 -f /
  PDF 报 unsupported（stub 剔除后不解析）。

新增格式样本：`cp third-party/anydoc/tests/fixtures/<fmt>/*.xxx testdata/fixtures/<fmt>/`
并跑 `python3 tools/gen_expected.py third-party/anydoc/tests/snapshots testdata/expected`。

## 已知限制

- panic = "abort"（wasm 内部崩溃直接终止进程）；
- 运行期线性内存 1.25MB 起、可 grow，单次 docx 转换约 4-5ms（goccy AOT）；
- PDF 不解析（stub 剔除，报 unsupported），见「架构」；
- 生成物接口以 `core` 包为准：导出为顶层 `core.New()`/`core.Memory(m)`/
  `core.Anydoc*` 函数，模块类型是 `core/base.Module`。lib 依赖这些名字，
  不要手工注释或改名（编译即校验）；
- `cli` 步骤（goccy 翻译）耗时长，core/ 已随仓库提交，日常构建无需重跑。

[goccy/wasm2go]: https://github.com/goccy/wasm2go
[ncruces/wasm2go]: https://github.com/ncruces/wasm2go