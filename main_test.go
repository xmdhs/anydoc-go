// 单元测试：用 anydoc 官方样本（testdata/fixtures）与快照期望
// （testdata/expected，tools/gen_expected.py 生成）逐字节对比；
// 错误路径断言稳定错误消息。
package main

import (
	"os"
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
		// 当前构建无 pdf feature：应报 unsupported。
		// 若将来打开 pdf 重建，本用例需更新为期望成功。
		input := readFile(t, "testdata/fixtures/pdf/text.pdf")
		_, err := conv.ConvertBytes(input, "")
		if err == nil || !strings.Contains(err.Error(), "unsupported") {
			t.Errorf("want unsupported error for pdf build, got %v", err)
		}
	})
}