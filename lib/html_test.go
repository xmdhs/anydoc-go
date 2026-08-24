// HTML → GFM 分支的单测（纯 Go，不走 wasm）。
package anydoc

import (
	"os"
	"strings"
	"testing"
)

func TestHTMLConvertString(t *testing.T) {
	conv := NewConverter()
	cases := []struct {
		name string
		html string
		want []string
	}{
		{"heading", "<h1>Title</h1><p>hello</p>", []string{"# Title", "hello"}},
		{"table", "<table><tr><th>a</th><th>b</th></tr><tr><td>1</td><td>2</td></tr></table>", []string{"| a", "| 1"}},
		{"strikethrough", "<p><del>gone</del></p>", []string{"~~gone~~"}},
		{"list", "<ul><li>one</li><li>two</li></ul>", []string{"- one", "- two"}},
		{"link", `<p><a href="https://example.com">more</a></p>`, []string{"[more](https://example.com)"}},
		{"image", `<p><img src="https://example.com/x.png" alt="pic" /></p>`, []string{"![pic](https://example.com/x.png)"}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			md, err := conv.ConvertBytes([]byte(c.html), "html")
			if err != nil {
				t.Fatalf("convert html: %v", err)
			}
			for _, w := range c.want {
				if !strings.Contains(md, w) {
					t.Errorf("missing %q in %q", w, md)
				}
			}
		})
	}
}

func TestHTMLSniffViaDocExtension(t *testing.T) {
	// 伪装成 .doc 的 HTML（Word 导出）虽后缀为 .doc，但内容为 HTML 时应走 HTML 分支。
	html := `<html><head><meta http-equiv="Content-Type" content="text/html; charset=utf-8"></head><body><p>伪装 doc 的 HTML</p><table><tr><th>k</th><th>v</th></tr><tr><td>a</td><td>b</td></tr></table></body></html>`
	conv := NewConverter()
	md, err := conv.ConvertBytes([]byte(html), "doc")
	if err != nil {
		t.Fatalf("convert fake doc html: %v", err)
	}
	if !strings.Contains(md, "伪装") || !strings.Contains(md, "| k") {
		t.Errorf("unexpected md: %q", md)
	}
	// 同样本，以文件路径兜底（-f 为空时按内容嗅探，也应走 HTML 分支）
	dir := t.TempDir()
	path := dir + "/sample.doc"
	if err := os.WriteFile(path, []byte(html), 0o644); err != nil {
		t.Fatal(err)
	}
	md2, err := conv.ConvertFile(path, "")
	if err != nil {
		t.Fatalf("ConvertFile fake doc html: %v", err)
	}
	if !strings.Contains(md2, "伪装") {
		t.Errorf("ConvertFile sniff failed: %q", md2)
	}
	// ExtractAssets / ConvertWithAssets 对 HTML 应返回空资产、无错误
	if assets, err := conv.ExtractAssets([]byte(html), "html"); err != nil || len(assets) != 0 {
		t.Fatalf("ExtractAssets html want 0, got %d err %v", len(assets), err)
	}
	res, err := conv.ConvertWithAssets([]byte(html), "html")
	if err != nil {
		t.Fatalf("ConvertWithAssets html: %v", err)
	}
	if !strings.Contains(res.Markdown, "伪装") || len(res.Assets) != 0 {
		t.Errorf("ConvertWithAssets html unexpected: md %q assets %d", res.Markdown, len(res.Assets))
	}
}

func TestHTMLAutoDetect(t *testing.T) {
	html := `<!doctype html><html><body><h2>Auto</h2><p>detect</p></body></html>`
	conv := NewConverter()
	md, err := conv.ConvertBytes([]byte(html), "")
	if err != nil {
		t.Fatalf("auto-detect html: %v", err)
	}
	if !strings.Contains(md, "## Auto") {
		t.Errorf("auto-detect miss: %q", md)
	}
}

func TestIsHTMLData(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"<html>hi</html>", true},
		{"<!doctype html><html>", true},
		{"\xEF\xBB\xBF  <html xmlns:v=\"urn:schemas-microsoft-com:vml\">", true},
		{"<meta http-equiv=\"Content-Type\" content=\"text/html; charset=gb2312\"><word>", true},
		{"PK\x03\x04", false},
		{"hello world", false},
		{"", false},
	}
	for _, c := range cases {
		if got := isHTMLData([]byte(c.in)); got != c.want {
			t.Errorf("isHTMLData(%q)=%v want %v", c.in, got, c.want)
		}
	}
}
