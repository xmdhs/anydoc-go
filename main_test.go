// 单元测试：用 anydoc 官方样本（testdata/fixtures）与快照期望
// （testdata/expected，tools/gen_expected.py 生成）逐字节对比；
// 错误路径断言稳定错误消息。
package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	anydoc "anydoc-go/lib"
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
	md0, err := conv.ConvertBytes(readFile(t, "testdata/fixtures/docx/in-image.docx"), "")
	if err != nil {
		t.Fatalf("convert: %v", err)
	}
	md, err := extractAssets(conv, "testdata/fixtures/docx/in-image.docx", "", md0, dir)
	if err != nil {
		t.Fatalf("extractAssets: %v", err)
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
		// pdf-inspector 被本地 stub patch 剔除（cabi Cargo.toml 的
		// [patch.crates-io]）：应报 unsupported。
		input := readFile(t, "testdata/fixtures/pdf/text.pdf")
		_, err := conv.ConvertBytes(input, "")
		if err == nil || !strings.Contains(err.Error(), "unsupported") {
			t.Errorf("want unsupported error for pdf build, got %v", err)
		}
	})
}
