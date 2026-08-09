// lib 包的测试：基本转换 + 并发安全。
package anydoc

import (
	"errors"
	"os"
	"strings"
	"sync"
	"testing"
)

func readFixture(t *testing.T, path string) []byte {
	t.Helper()
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return b
}

// Word 系格式（docx：ZIP 容器；doc：OLE/CFB 容器，两条不同解析路径）
// 都要与官方快照逐字节一致。
func TestConvertWordFormats(t *testing.T) {
	cases := []struct{ fixture, expected string }{
		{"docx/text.docx", "docx/text.md"},
		{"doc/text.doc", "doc/text.md"},
	}
	for _, c := range cases {
		t.Run(c.fixture, func(t *testing.T) {
			conv := NewConverter()
			got, err := conv.ConvertBytes(readFixture(t, "../testdata/fixtures/"+c.fixture), "")
			if err != nil {
				t.Fatalf("convert: %v", err)
			}
			want := readFixture(t, "../testdata/expected/"+c.expected)
			if got != string(want) {
				t.Errorf("output mismatch: got %d bytes, want %d bytes", len(got), len(want))
			}
		})
	}
}

// 多实例并行：Converter 非线程安全，并发转换的正确姿势是每 goroutine
// 一个独立实例（模块内存各自独立，天然并行）。
func TestConverterConcurrent(t *testing.T) {
	input := readFixture(t, "../testdata/fixtures/docx/text.docx")
	want, err := NewConverter().ConvertBytes(input, "")
	if err != nil {
		t.Fatal(err)
	}

	var wg sync.WaitGroup
	errs := make([]error, 8)
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			conv := NewConverter() // 每 goroutine 独立实例
			got, err := conv.ConvertBytes(input, "")
			if err != nil {
				errs[i] = err
				return
			}
			if got != want {
				errs[i] = os.ErrInvalid // 哨兵表达结果不一致
			}
		}(i)
	}
	wg.Wait()
	for i, err := range errs {
		if err != nil {
			t.Fatalf("goroutine %d: %v", i, err)
		}
	}
}

// ConvertFile 的格式兜底：无签名格式（CSV）自动检测失败后回退扩展名。
func TestConvertFileFallsBackToExtension(t *testing.T) {
	conv := NewConverter()
	md, err := conv.ConvertFile("../testdata/fixtures/csv/sheet.csv", "")
	if err != nil {
		t.Fatalf("convert file: %v", err)
	}
	if !strings.Contains(md, "| Kind |") {
		t.Errorf("unexpected csv output: %q", md[:min(len(md), 60)])
	}
	// 显式 -f 语义不变：错误格式仍要报错。
	if _, err := conv.ConvertFile("../testdata/fixtures/csv/sheet.csv", "pptx"); err == nil {
		t.Error("wrong explicit format should fail")
	}
}

// 错误路径返回 *ConvertError，其 Code 是与官方 anydoc 绑定一致的稳定标识。
func TestLibErrorCode(t *testing.T) {
	conv := NewConverter()
	_, err := conv.ConvertBytes([]byte("garbage"), "")
	if err == nil {
		t.Fatal("want error")
	}
	var cerr *ConvertError
	if !errors.As(err, &cerr) {
		t.Fatalf("want *ConvertError, got %T", err)
	}
	if cerr.Code != "unsupported" {
		t.Errorf("want code unsupported, got %q", cerr.Code)
	}
}