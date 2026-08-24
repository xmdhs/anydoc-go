// 单元测试：用 anydoc 官方样本（testdata/fixtures）与快照期望
// （testdata/expected，tools/gen_expected.py 生成）逐字节对比；
// 错误路径断言稳定错误消息。
package main

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"fmt"
	"hash/crc32"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"testing"

	anydoc "github.com/xmdhs/anydoc-go/lib"
)

func readFile(t *testing.T, path string) []byte {
	t.Helper()
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return b
}

func TestConvertMatchesExpected(t *testing.T) {
	conv := anydoc.NewConverter()
	cases := []struct {
		name, format, expected string
	}{
		{"docx/text.docx", "", "docx/text.md"},
		{"csv/sheet.csv", "csv", "csv/sheet.md"},
		{"doc/text.doc", "", "doc/text.md"},
		{"pptx/pres.pptx", "", "pptx/pres.md"},
		{"epub/book.epub", "", "epub/book.md"},
		{"ods/sheet.ods", "", "ods/sheet.md"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			input := readFile(t, "testdata/fixtures/"+c.name)
			got, err := conv.ConvertBytes(input, c.format)
			if err != nil {
				t.Fatalf("convert: %v", err)
			}
			want := readFile(t, "testdata/expected/"+c.expected)
			if got != string(want) {
				t.Errorf("output mismatch: got %d bytes, want %d bytes", len(got), len(want))
			}
		})
	}
}

func TestExtractAssetsCLI(t *testing.T) {
	conv := anydoc.NewConverter()
	dir := t.TempDir()
	md, err := convertFileWithAssets(conv, "testdata/fixtures/docx/in-image.docx", "", dir)
	if err != nil {
		t.Fatalf("convertFileWithAssets: %v", err)
	}
	// 名称为 <stem>-<id>.png；各输入独立 imgs/ 目录，无需防覆盖。
	entries, err := os.ReadDir(dir)
	if err != nil {
		t.Fatalf("read dir: %v", err)
	}
	if len(entries) != 1 || !strings.HasPrefix(entries[0].Name(), "in-image-") ||
		!strings.HasSuffix(entries[0].Name(), ".png") {
		t.Fatalf("unexpected assets: %v", entries)
	}
	name := entries[0].Name()
	want := "![a red circle](imgs/" + name + ")"
	if !strings.Contains(md, want) {
		t.Errorf("markdown missing rewritten asset link: %q", md)
	}
	b, err := os.ReadFile(filepath.Join(dir, name))
	if err != nil {
		t.Fatalf("asset not written: %v", err)
	}
	if len(b) != 70 {
		t.Errorf("asset bytes: want 70, got %d", len(b))
	}
	if ext := assetExt("image/png"); ext != "png" {
		t.Errorf("assetExt(image/png) = %q", ext)
	}
	if ext := assetExt("application/octet-stream"); ext != "bin" {
		t.Errorf("assetExt(octet) = %q", ext)
	}
}

func TestConverterIsStateless(t *testing.T) {
	// 同一转换器连续转换两次，结果必须一致（alloc/free 不留状态）。
	conv := anydoc.NewConverter()
	input := readFile(t, "testdata/fixtures/docx/text.docx")
	first, err := conv.ConvertBytes(input, "")
	if err != nil {
		t.Fatalf("first convert: %v", err)
	}
	second, err := conv.ConvertBytes(input, "")
	if err != nil {
		t.Fatalf("second convert: %v", err)
	}
	if first != second {
		t.Error("two conversions of the same input differ")
	}
}

func TestErrors(t *testing.T) {
	conv := anydoc.NewConverter()
	t.Run("encrypted", func(t *testing.T) {
		input := readFile(t, "testdata/fixtures/malformed/encrypted--errors.odt")
		_, err := conv.ConvertBytes(input, "")
		if err == nil || !strings.Contains(err.Error(), "encrypted") {
			t.Errorf("want encrypted error, got %v", err)
		}
	})
	t.Run("malformed", func(t *testing.T) {
		// 截断的 docx 连签名都无法识别，自动检测会先报 unsupported；
		// 显式指定 docx 后进入解析器才得到 malformed。
		input := readFile(t, "testdata/fixtures/malformed/truncated--errors.docx")
		_, err := conv.ConvertBytes(input, "docx")
		if err == nil || !strings.Contains(err.Error(), "malformed") {
			t.Errorf("want malformed error, got %v", err)
		}
	})
	t.Run("unrecognized", func(t *testing.T) {
		_, err := conv.ConvertBytes([]byte("not a document at all"), "")
		if err == nil || !strings.Contains(err.Error(), "unsupported") {
			t.Errorf("want unsupported error, got %v", err)
		}
	})
	t.Run("csvNeedsFormat", func(t *testing.T) {
		// CSV 无签名：不传格式应报 unsupported，传了则以固定格式成功。
		input := readFile(t, "testdata/fixtures/csv/sheet.csv")
		if _, err := conv.ConvertBytes(input, ""); err == nil {
			t.Error("csv without format should fail")
		}
		if _, err := conv.ConvertBytes(input, "csv"); err != nil {
			t.Errorf("csv with format should succeed: %v", err)
		}
	})
	t.Run("pdfUncompiled", func(t *testing.T) {
		// 已启用真实 pdf-inspector（patches/getrandom 去 wasm-bindgen，
		// 无需 stub）：testdata/fixtures/pdf/text.pdf 应可转换出内容。
		input := readFile(t, "testdata/fixtures/pdf/text.pdf")
		md, err := conv.ConvertBytes(input, "")
		if err != nil {
			t.Fatalf("pdf should convert, got %v", err)
		}
		if !strings.Contains(md, "Fixture Document") {
			t.Errorf("pdf markdown missing expected content: %q", md[:min(len(md), 80)])
		}
	})
}

func TestRewriteAssetRefs(t *testing.T) {
	// 画像：asset://3 与 asset://31 并存。旧实现 strings.ReplaceAll 先替换
	// asset://3 会误吞 asset://31 前缀，残留 `1` 产生 imgs/doc-3.png1。
	refs := map[string]string{
		"asset://3":  "imgs/doc-3.png",
		"asset://31": "imgs/doc-31.png",
	}
	md := "![alt-3](asset://3) and ![alt-31](asset://31)"
	got, err := rewriteAssetRefs(md, refs)
	if err != nil {
		t.Fatalf("rewriteAssetRefs: %v", err)
	}
	want := "![alt-3](imgs/doc-3.png) and ![alt-31](imgs/doc-31.png)"
	if got != want {
		t.Errorf("rewriteAssetRefs:\n got %q\nwant %q", got, want)
	}
}

func TestRewriteAssetRefsMissingID(t *testing.T) {
	// 未在映射中的占位必须报错，而不是静默保留（否则产物里是坏链接）。
	md := "![x](asset://99)"
	if _, err := rewriteAssetRefs(md, map[string]string{"asset://1": "imgs/a.png"}); err == nil {
		t.Error("missing asset id should error, not silently keep placeholder")
	}
}

func TestWriteAssetRefsOrderIndependent(t *testing.T) {
	// 资产按乱序给出（模拟多图、非按渲染顺序），引用必须仍按 asset://<ID>
	// 精确配对（顺序消费依赖的是"占位在 md 里的出现"与 ID，而非资产列表
	// 顺序）。这同时锁住「asset://<id> 与 Assets[id].ID 一一对应」的不变量。
	md := "![b](asset://1) ![a](asset://0)"
	assets := []anydoc.Asset{
		{ID: 1, MediaType: "image/png", Bytes: []byte("png-1")},
		{ID: 0, MediaType: "image/png", Bytes: []byte("png-0")},
	}
	dir := t.TempDir()
	out, err := writeAssetRefs(md, assets, dir, "doc")
	if err != nil {
		t.Fatalf("writeAssetRefs: %v", err)
	}
	for _, want := range []string{"imgs/doc-0.png", "imgs/doc-1.png"} {
		if !strings.Contains(out, want) {
			t.Errorf("output missing %q: %q", want, out)
		}
	}
	// 字节按各自 ID 落盘正确。
	for _, a := range assets {
		b, err := os.ReadFile(filepath.Join(dir, fmt.Sprintf("doc-%d.png", a.ID)))
		if err != nil {
			t.Fatalf("asset %d not written: %v", a.ID, err)
		}
		if string(b) != string(a.Bytes) {
			t.Errorf("asset %d bytes mismatch: got %d bytes", a.ID, len(b))
		}
	}
}

func TestRewriteAssetRefsNonNumericPrefix(t *testing.T) {
	// "asset://" 后不是数字时不当作占位，原样保留、不误吞正文。
	md := "text asset://abc and asset://"
	got, err := rewriteAssetRefs(md, map[string]string{})
	if err != nil {
		t.Fatalf("non-numeric placeholder should not error: %v", err)
	}
	if got != md {
		t.Errorf("non-numeric prefix should stay intact: got %q", got)
	}
}

func TestRewriteAssetRefsKeepsPlainImage(t *testing.T) {
	// 普通图片链接（dest 非 asset://）原样保留，不被误改写。
	md := "![photo](images/photo.png) and ![alt-3](asset://3)"
	refs := map[string]string{"asset://3": "imgs/doc-3.png"}
	got, err := rewriteAssetRefs(md, refs)
	if err != nil {
		t.Fatalf("rewriteAssetRefs: %v", err)
	}
	want := "![photo](images/photo.png) and ![alt-3](imgs/doc-3.png)"
	if got != want {
		t.Errorf("rewriteAssetRefs:\n got %q\nwant %q", got, want)
	}
}

func TestRewriteAssetRefsLiteralPrefixUntouched(t *testing.T) {
	// 正文/非图片上下文里的 `asset://<数字>` 不是占位：保留原文且不报错
	// （旧实现会把它当占位，缺失即整份转换失败或误改写）。
	md := "prose mentions asset://99 and asset://3 here"
	got, err := rewriteAssetRefs(md, map[string]string{"asset://3": "imgs/doc-3.png"})
	if err != nil {
		t.Fatalf("literal asset:// text should not error: %v", err)
	}
	if got != md {
		t.Errorf("literal asset:// should stay intact: got %q", got)
	}
}

func TestRewriteAssetRefsEscapedAlt(t *testing.T) {
	// alt 含未转义括号、转义 `\]`、以及内嵌子图片链接时，仍能正确配对
	// 外层 `](dest)` 而只重写 dest 里的 asset 占位。
	md := `![a \[literal\] (paren) recursed![?](img)](asset://3) tail`
	refs := map[string]string{"asset://3": "imgs/doc-3.png"}
	got, err := rewriteAssetRefs(md, refs)
	if err != nil {
		t.Fatalf("rewriteAssetRefs: %v", err)
	}
	want := `![a \[literal\] (paren) recursed![?](img)](imgs/doc-3.png) tail`
	if got != want {
		t.Errorf("rewriteAssetRefs:\n got %q\nwant %q", got, want)
	}
}

func TestURLPathEscape(t *testing.T) {
	cases := []struct{ in, want string }{
		{"doc-3.png", "doc-3.png"},               // 仅 unreserved 保持原样
		{"my report-1.png", "my%20report-1.png"}, // 空格 → %20
		{"图-2.png", "%E5%9B%BE-2.png"},           // 非 ASCII → %XX
		{"keep._~-x.png", "keep._~-x.png"},       // 运算符等子集保持
	}
	for _, c := range cases {
		if got := url.PathEscape(c.in); got != c.want {
			t.Errorf("url.PathEscape(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}

func TestPdfImageAssets(t *testing.T) {
	conv := anydoc.NewConverter()
	cases := []struct {
		name       string
		wantAssets int
		wantMedia  map[string]bool
		wantDims   [][2]int  // per-asset width×height, 0 means skip
		checkFirst [][3]byte // per-asset first pixel RGB (non-empty means check)
	}{
		{"pdf/with-text-image.pdf", 1, map[string]bool{"image/jpeg": true}, nil, nil},
		{"pdf/scanned-image.pdf", 1, map[string]bool{"image/jpeg": true}, nil, nil},
		{"pdf/multi-image.pdf", 2, map[string]bool{"image/jpeg": true, "image/png": true}, nil, nil},
		{"pdf/ccitt-g4.pdf", 1, map[string]bool{"image/png": true}, [][2]int{{415, 314}}, nil},
		{"pdf/jbig2-globals.pdf", 1, map[string]bool{"image/png": true}, [][2]int{{1747, 2554}}, nil},
		{"pdf/indexed-rgb.pdf", 1, map[string]bool{"image/png": true}, [][2]int{{717, 717}}, nil},
		{"pdf/cmyk-flate.pdf", 1, map[string]bool{"image/png": true}, [][2]int{{2, 2}}, [][3]byte{{0x00, 0xFF, 0xFF}}}, // cyan
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			data := readFile(t, "testdata/fixtures/"+c.name)
			res, err := conv.ConvertWithAssets(data, "")
			if err != nil {
				t.Fatalf("ConvertWithAssets: %v", err)
			}
			if len(res.Assets) != c.wantAssets {
				t.Fatalf("want %d assets, got %d", c.wantAssets, len(res.Assets))
			}
			for _, a := range res.Assets {
				if !c.wantMedia[a.MediaType] {
					t.Errorf("unexpected media_type %q", a.MediaType)
				}
				if len(a.Bytes) == 0 {
					t.Errorf("asset %d empty bytes", a.ID)
				}
				// JPEG magic
				if a.MediaType == "image/jpeg" && (len(a.Bytes) < 2 || a.Bytes[0] != 0xFF || a.Bytes[1] != 0xD8) {
					t.Errorf("jpeg asset %d missing SOI FF D8", a.ID)
				}
				if a.MediaType == "image/png" {
					if len(a.Bytes) < 4 || a.Bytes[0] != 0x89 {
						t.Errorf("png asset %d missing PNG sig", a.ID)
					}
					// Pixel correctness: decode PNG IHDR + IDAT, check dims and first pixel
					if len(c.wantDims) > int(a.ID) && c.wantDims[a.ID][0] != 0 {
						w, h, ct, raw := decodePngForTest(a.Bytes)
						wantW, wantH := c.wantDims[a.ID][0], c.wantDims[a.ID][1]
						if w != wantW || h != wantH {
							t.Errorf("png asset %d dims %dx%d, want %dx%d", a.ID, w, h, wantW, wantH)
						}
						// Grayscale / RGB share first pixel checks
						_ = raw
						_ = ct
					}
					if len(c.checkFirst) > int(a.ID) && c.checkFirst[a.ID] != [3]byte{} {
						_, _, _, raw := decodePngForTest(a.Bytes)
						want := c.checkFirst[a.ID]
						if len(raw) < 3 || raw[0] != want[0] || raw[1] != want[1] || raw[2] != want[2] {
							t.Errorf("png asset %d first pixel got %02x %02x %02x, want %02x %02x %02x", a.ID, raw[0], raw[1], raw[2], want[0], want[1], want[2])
						}
					}
				}
			}
			// Markdown must contain asset:// placeholders
			for _, a := range res.Assets {
				placeholder := fmt.Sprintf("asset://%d", a.ID)
				if !strings.Contains(res.Markdown, placeholder) {
					t.Errorf("markdown missing %q, got: %q", placeholder, res.Markdown[:testMin(300, len(res.Markdown))])
				}
			}
			// Also via writeAssetRefs:落盘并重写
			dir := t.TempDir()
			out, err := writeAssetRefs(res.Markdown, res.Assets, dir, strings.TrimSuffix(filepath.Base(c.name), filepath.Ext(c.name)))
			if err != nil {
				t.Fatalf("writeAssetRefs: %v", err)
			}
			if !strings.Contains(out, "imgs/") {
				t.Errorf("rewritten markdown missing imgs/: %q", out[:testMin(200, len(out))])
			}
			entries, _ := os.ReadDir(dir)
			if len(entries) != c.wantAssets {
				t.Errorf("want %d files on disk, got %d", c.wantAssets, len(entries))
			}
		})
	}
}

// decodePngForTest 解码本项目 encode_png 生成的极简 PNG（IHDR+IDAT(Zlib)+IEND，filter 0），返回宽高、colorType 与原始像素。
func decodePngForTest(png []byte) (w, h int, colorType byte, raw []byte) {
	if len(png) < 8 || string(png[:8]) != "\x89PNG\r\n\x1a\n" {
		return 0, 0, 0, nil
	}
	pos := 8
	var idat []byte
	for pos+12 <= len(png) {
		length := int(binary.BigEndian.Uint32(png[pos : pos+4]))
		typ := string(png[pos+4 : pos+8])
		data := png[pos+8 : pos+8+length]
		wantCrc := binary.BigEndian.Uint32(png[pos+8+length : pos+12+length])
		chunkHdr := png[pos+4 : pos+8+length]
		if crc32.ChecksumIEEE(chunkHdr) != wantCrc {
			break
		}
		if typ == "IHDR" && length >= 13 {
			w = int(binary.BigEndian.Uint32(data[0:4]))
			h = int(binary.BigEndian.Uint32(data[4:8]))
			colorType = data[9]
		} else if typ == "IDAT" {
			idat = append(idat, data...)
		}
		pos += 12 + length
	}
	if len(idat) == 0 {
		return w, h, colorType, nil
	}
	zr, err := zlib.NewReader(bytes.NewReader(idat))
	if err != nil {
		return w, h, colorType, nil
	}
	r := io.NopCloser(zr)
	defer func() { _ = r.Close() }()
	inflated, err := io.ReadAll(r)
	if err != nil {
		return w, h, colorType, nil
	}
	// Each row: 1 filter byte (0) + channels*w bytes. Our encoder always uses filter 0.
	channels := 1
	switch colorType {
	case 0:
		channels = 1
	case 2:
		channels = 3
	case 6:
		channels = 4
	}
	raw = make([]byte, 0, w*h*channels)
	for y := 0; y < h; y++ {
		off := y * (w*channels + 1)
		if off+w*channels+1 > len(inflated) {
			break
		}
		raw = append(raw, inflated[off+1:off+1+w*channels]...)
	}
	return w, h, colorType, raw
}

func testMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
