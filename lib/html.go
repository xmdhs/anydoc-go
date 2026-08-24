// html.go 提供 HTML → GFM Markdown 的纯 Go 分支。
//
// 选用 JohannesKaufmann/html-to-markdown (v2) 并启用 GFM 相关插件：
//   - base + commonmark  必需基础
//   - table             GFM 表格（对齐、合并单元格等）
//   - strikethrough     GFM 删除线 ~~
//
// Word 导出的 “假 .doc”（实为 HTML，charset 常为 gb2312/GBK）在此分支
// 完成解码与转换，不走 wasm/anydoc。检测策略：
//   显式 -f html/htm 优先；其次内容嗅探（去除 BOM/空白后以 <!doctype html
//   或 <html 开头，或首字符 '<' 且前 2KB 内含 <html/<head/<body）。
//
// 编码：Word HTML 的 <meta charset=gb2312> 与实际字节一致，
// golang.org/x/net/html/charset.NewReader 会据此把 GBK 转为 UTF-8。
package anydoc

import (
	"bytes"
	"io"
	"strings"
	"sync"

	"github.com/JohannesKaufmann/html-to-markdown/v2/converter"
	"github.com/JohannesKaufmann/html-to-markdown/v2/plugin/base"
	"github.com/JohannesKaufmann/html-to-markdown/v2/plugin/commonmark"
	"github.com/JohannesKaufmann/html-to-markdown/v2/plugin/strikethrough"
	"github.com/JohannesKaufmann/html-to-markdown/v2/plugin/table"

	"golang.org/x/net/html/charset"
)

var (
	htmlConv     *converter.Converter
	htmlConvOnce sync.Once
)

func getHTMLConverter() *converter.Converter {
	htmlConvOnce.Do(func() {
		htmlConv = converter.NewConverter(
			converter.WithPlugins(
				base.NewBasePlugin(),
				commonmark.NewCommonmarkPlugin(),
				// 表格：GFM 规范；Word 表往往无 <th>，提升首行作表头
				table.NewTablePlugin(
					table.WithHeaderPromotion(true),
				),
				strikethrough.NewStrikethroughPlugin(),
			),
		)
	})
	return htmlConv
}

// isHTMLFormat 判断扩展名是否显式指定为 HTML。
func isHTMLFormat(ext string) bool {
	ext = strings.ToLower(strings.TrimPrefix(strings.TrimSpace(ext), "."))
	switch ext {
	case "html", "htm", "xhtml", "htmlx":
		return true
	default:
		return false
	}
}

// isHTMLData 嗅探字节流是否像 HTML 文档。
//
// 规则保守：先去掉 UTF-8 BOM 与空白，再看前缀是否为 <!doctype html 或 <html，
// 或首字符为 '<' 且前 2KB 内含 <html/<head/<body（足够覆盖 Word 导出的
// <html xmlns:v=...> 形态）。纯文本/二进制不会以 '<' 开头，故误判极低。
func isHTMLData(data []byte) bool {
	if len(data) == 0 {
		return false
	}
	// 去 BOM + 空白
	s := data
	if bytes.HasPrefix(s, []byte{0xEF, 0xBB, 0xBF}) {
		s = s[3:]
	}
	s = bytes.TrimLeft(s, " \t\r\n\x0c")
	if len(s) == 0 || s[0] != '<' {
		return false
	}
	n := len(s)
	if n > 2048 {
		n = 2048
	}
	prefix := strings.ToLower(string(s[:n]))
	prefix = strings.TrimLeft(prefix, " \t\r\n")
	if strings.HasPrefix(prefix, "<!doctype html") || strings.HasPrefix(prefix, "<html") {
		return true
	}
	if strings.Contains(prefix, "<html") || strings.Contains(prefix, "<head") || strings.Contains(prefix, "<body") {
		return true
	}
	// Word 导出常含 xmlns:o / xmlns:w 但无 <!doctype>，已在上式覆盖；
	// 额外放宽：含 "<meta" 且含 "word" 的也可视为 HTML（极少误伤）
	if strings.Contains(prefix, "<meta") && strings.Contains(prefix, "word") {
		return true
	}
	return false
}

// decodeHTMLBytes 把 HTML 原始字节按其声明的 charset 解为 UTF-8 字符串。
//
// Word 导出的 HTML 往往在 <meta http-equiv=Content-Type content="text/html; charset=gb2312">
// 声明 GBK，若直接按 UTF-8 解会乱码。charset.NewReader 会嗅探前 1KB 的 <meta>
// 并用 transform 转码。
func decodeHTMLBytes(data []byte) (string, error) {
	if len(data) == 0 {
		return "", nil
	}
	r, err := charset.NewReader(bytes.NewReader(data), "")
	if err != nil {
		return "", err
	}
	decoded, err := io.ReadAll(r)
	if err != nil {
		return "", err
	}
	return string(decoded), nil
}

// htmlToMarkdown 执行纯 Go 的 HTML→GFM 转换。
func htmlToMarkdown(data []byte) (string, error) {
	htmlStr, err := decodeHTMLBytes(data)
	if err != nil {
		return "", err
	}
	conv := getHTMLConverter()
	md, err := conv.ConvertString(htmlStr)
	if err != nil {
		return "", err
	}
	return md, nil
}
