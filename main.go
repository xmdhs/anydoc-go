// anydoc converts documents (docx, pptx, xlsx, csv, pdf, ...) to
// GitHub-Flavored Markdown. The conversion core lives in the
// github.com/xmdhs/anydoc-go/lib
// package (封装 wasm2go 生成的 core 包)；本文件只是命令行入口。
//
// Usage:
//
//	anydoc [-o|FILE|-s] [-f format] <file|glob>...
//
// 默认把每个输入转换成同目录同名 <name>.md；-o 指定单个输出文件
// （"-" 也作 stdout 别名）；-s 输出到标准输出；参数支持 glob 通配符。
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	anydoc "github.com/xmdhs/anydoc-go/lib" // 包的内部名也是 anydoc；别名保证可读
)

func main() {
	var (
		output = flag.String("o", "", "write Markdown to `file` (single input only; `-` = stdout)")
		stdout = flag.Bool("s", false, "print Markdown to stdout (works with multiple inputs)")
		format = flag.String("f", "", "force input `format` by extension (docx, csv, pdf, ...); default auto-detect")
		imgs   = flag.Bool("imgs", false, "write embedded assets into <input dir>/imgs")
	)
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: anydoc [options] <file|glob>...\n\n")
		fmt.Fprintf(os.Stderr, "Convert documents to GitHub-Flavored Markdown.\n")
		fmt.Fprintf(os.Stderr, "By default each input becomes a same-named <input>.md next to it.\n\nOptions:\n")
		flag.PrintDefaults()
	}
	flag.Parse()
	if flag.NArg() < 1 {
		flag.Usage()
		os.Exit(2)
	}

	// Glob 展开参数；无匹配的按字面路径处理（文件不存在时由转换报错）。
	var files []string
	for _, arg := range flag.Args() {
		matches, err := filepath.Glob(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "anydoc: bad pattern %q: %v\n", arg, err)
			os.Exit(2)
		}
		if len(matches) == 0 {
			files = append(files, arg)
		} else {
			files = append(files, matches...)
		}
	}

	if *output != "" && *output != "-" && len(files) > 1 {
		fmt.Fprintln(os.Stderr, "anydoc: -o 只适用单个输入；多个文件请用 -s（stdout）或默认生成同目录 .md")
		os.Exit(2)
	}

	conv := anydoc.NewConverter()

	failed := false
	for _, file := range files {
		// --imgs 时顺带取内嵌资产落盘并重写引用（convertFileWithAssets 按
		// anydoc_convert tag 可能是两次解析，也可能是一次合并解析）。
		var md string
		var err error
		if *imgs {
			md, err = convertFileWithAssets(conv, file, *format, filepath.Join(filepath.Dir(file), "imgs"))
		} else {
			md, err = conv.ConvertFile(file, *format)
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "anydoc: %s: %v\n", file, err)
			failed = true
			continue
		}
		switch {
		case *stdout || *output == "-":
			fmt.Fprint(os.Stdout, md)
			if len(md) > 0 && md[len(md)-1] != '\n' {
				fmt.Fprintln(os.Stdout)
			}
		case *output != "":
			if err := os.WriteFile(*output, []byte(md), 0o644); err != nil {
				fmt.Fprintf(os.Stderr, "anydoc: %v\n", err)
				os.Exit(1)
			}
		default: // <input>.md 同目录
			if err := os.WriteFile(file+".md", []byte(md), 0o644); err != nil {
				fmt.Fprintf(os.Stderr, "anydoc: %v\n", err)
				os.Exit(1)
			}
		}
	}
	if failed {
		os.Exit(1)
	}
}

// convertFileWithAssets 与资产落盘/引用实现分文件，按 anydoc_convert build
// tag 切换（见 asset_common.go / assets_fallback.go / assets_merge.go）。

// rewriteAssetRefs 把 md 里的 `asset://<id>` 占位替换为 refs 中对应引用。
//
// 顺序消费 + 显式报错，替代旧的全文正则扫描（regexp `asset://\d+`）：
//   - 不打正则、不做子串 ReplaceAll，因此 `asset://3` 不会误吞 `asset://31`
//     （按"asset://" 后一整段连续数字取 ID）；
//   - `asset://<id>` 后必须是连续十进制数字，否则按字面保留；
//   - 占位 ID 在 refs 中缺失时**返回错误**而非静默保留——未导出资产却留有
//     占位引用意味着产物里有坏链接，应当显式失败而非产出污染 markdown。
func rewriteAssetRefs(md string, refs map[string]string) (string, error) {
	var sb strings.Builder
	rest := md
	for {
		idx := strings.Index(rest, "asset://")
		if idx < 0 {
			sb.WriteString(rest)
			return sb.String(), nil
		}
		sb.WriteString(rest[:idx])
		// rest 从首个非 "asset://" 前缀的字符开始，读一整段数字作为 ID。
		rest = rest[idx+len("asset://"):]
		j := 0
		for j < len(rest) && rest[j] >= '0' && rest[j] <= '9' {
			j++
		}
		if j == 0 { // "asset://" 后不是数字：按字面保留
			sb.WriteString("asset://")
			continue
		}
		idStr := rest[:j]
		ref, ok := refs["asset://"+idStr]
		if !ok {
			return "", fmt.Errorf("asset placeholder asset://%s not among extracted assets", idStr)
		}
		sb.WriteString(ref)
		rest = rest[j:]
	}
}

// assetExt 从 media type 推导文件扩展名：image/* 取子类型字母数字，
// 其余一律 bin（与 anydoc 官方 example convert.rs 的规则一致）。
func assetExt(mediaType string) string {
	kind, subtype, ok := strings.Cut(mediaType, "/")
	if !ok || kind != "image" {
		return "bin"
	}
	var b strings.Builder
	for _, r := range subtype {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
			b.WriteRune(r)
		}
	}
	if b.Len() == 0 {
		return "bin"
	}
	return b.String()
}
