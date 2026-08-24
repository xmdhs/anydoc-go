# anydoc-go

把常见文档（docx / pptx / xlsx / csv / odt / rtf / epub / html，以及伪装成 `.doc` 的 Word HTML）转成 GitHub-Flavored Markdown。基于 [firecrawl/anydoc]（Rust）与 [JohannesKaufmann/html-to-markdown]（Go），以纯 Go 方式分发，无需本地工具链。

## 安装与构建

```bash
go build -o bin/anydoc .                 # 纯 Go 构建（日常改 lib/main.go 足够）
./build.sh fetch && ./build.sh build     # 完整管线：拉第三方 → wasm → 纯 Go 生成物（core/）
```

## 用法（CLI）

```bash
bin/anydoc report.docx                   # → report.docx.md（同目录）
bin/anydoc "docs/*.docx" "*.csv"         # glob 批量（模式含 * 时加引号）
bin/anydoc -s report.docx                # 输出到 stdout（可多个输入）
bin/anydoc -o out.md report.docx         # 指定单个输出文件
bin/anydoc -o - report.docx              # 与 -s 等价
bin/anydoc -f docx odd.bin               # 按扩展名强制定格式
bin/anydoc --imgs "docs/*.docx"          # 导出内嵌图片到 imgs/
```

**格式识别**：`-f` 显式 > 内容自动检测 > 路径扩展名兜底。`html`/`htm` 或内容形如 `<!doctype html` / `<html` 时走 HTML 分支（Word 导出的伪装 `.doc` 自动识别，`gb2312/GBK` 自动转 UTF-8）。

**图片导出**（docx/pptx 等）：默认 `![alt](asset://<id>)` 占位；`--imgs` 将资产写入 `imgs/<stem>-<id>.<ext>` 并重写为相对路径。`lib/` 对应 `ExtractAssets` / `ConvertWithAssets`（同一次解析，`asset://` 与 `Assets[id]` 一一对应）。

## 用法（Go 库）

```go
import lib "github.com/xmdhs/anydoc-go/lib"

conv := lib.NewConverter()                      // 每个实例非线程安全；并发请每 goroutine 一个
md, err := conv.ConvertFile("report.docx", "")  // 自动检测（html 伪装 doc 也可）
md, err = conv.ConvertBytes(data, "csv")        // 无签名格式需显式扩展名
_ = conv.ConvertFileTo("a.docx", "", "a.md")    // 便捷写出

if cerr, ok := err.(*lib.ConvertError); ok {    // Code: unsupported / malformed / encrypted / ...
    // 处理
}

// 合并解析（一次得到 Markdown + 资产，CLI --imgs 即用此路径）
res, _ := conv.ConvertWithAssets(data, "")
_ = res.Markdown
_ = res.Assets // []Asset{ID, MediaType, Bytes}
```

## 验证

```bash
./build.sh test   # go test ./...（含 26 个伪装 HTML 样本的回归）
```

[firecrawl/anydoc]: https://github.com/firecrawl/anydoc
[JohannesKaufmann/html-to-markdown]: https://github.com/JohannesKaufmann/html-to-markdown
