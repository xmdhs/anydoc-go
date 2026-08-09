// anydoc converts documents (docx, pptx, xlsx, csv, pdf, ...) to
// GitHub-Flavored Markdown. The conversion core lives in the anydoc-go/lib
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

	anydoc "anydoc-go/lib" // 包的内部名也是 anydoc；别名保证可读
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
		md, err := conv.ConvertFile(file, *format)
		if err != nil {
			fmt.Fprintf(os.Stderr, "anydoc: %s: %v\n", file, err)
			failed = true
			continue
		}
		// 图片导出：--imgs 时导出到输入文件同目录的 imgs/（引用相对该目录）；
		// 否则不导出，占位原样保留。
		if *imgs {
			md, err = extractAssets(conv, file, *format, md, filepath.Join(filepath.Dir(file), "imgs"))
			if err != nil {
				fmt.Fprintf(os.Stderr, "anydoc: %s: %v\n", file, err)
				failed = true
				continue
			}
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

// extractAssets 把文档内嵌资产写入 dir（文件名 <stem>-<ID>.<ext>），并把
// markdown 里的 asset://<ID> 占位替换为相对引用（ref 为固定 imgs/ 目录，
// 与输入文件同级的 imgs/）。
// 与 ConvertFile 相同的格式兜底（显式 -f → 内容自动检测 → 路径扩展名）。
// 不同输入各自导到自己的 imgs/，无跨输入覆盖。
func extractAssets(conv *anydoc.Converter, file, format, md, dir string) (string, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return md, err
	}
	formats := []string{format}
	if format == "" {
		if ext := strings.TrimPrefix(filepath.Ext(file), "."); ext != "" {
			formats = append(formats, ext)
		}
	}
	var assets []anydoc.Asset
	for _, f := range formats {
		if assets, err = conv.ExtractAssets(data, f); err == nil {
			break
		}
	}
	if err != nil {
		return md, err
	}
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return md, err
	}
	// 各输入导出到自己的 imgs/，无跨输入覆盖；重复转换幂等（覆盖自己）。
	stem := strings.TrimSuffix(filepath.Base(file), filepath.Ext(file))
	for _, a := range assets {
		name := fmt.Sprintf("%s-%d.%s", stem, a.ID, assetExt(a.MediaType))
		if err := os.WriteFile(filepath.Join(dir, name), a.Bytes, 0o644); err != nil {
			return md, err
		}
		md = strings.ReplaceAll(md, fmt.Sprintf("asset://%d", a.ID), filepath.Join("imgs", name))
	}
	return md, nil
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
