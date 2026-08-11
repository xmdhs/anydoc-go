// anydoc converts documents (docx, pptx, xlsx, csv, pdf, ...) to
// GitHub-Flavored Markdown. The conversion core lives in the
// github.com/xmdhs/anydoc-go/lib
// package (封装 wasm2go 生成的 core 包)；本文件只是命令行入口。
//
// Usage:
//
//	anydoc [-o|FILE|-s] [-f format] [--imgs] <file|glob>...
//
// 默认把每个输入转换成同目录同名 <name>.md；-o 指定单个输出文件
// （"-" 也作 stdout 别名）；-s 输出到标准输出；参数支持 glob 通配符；
// --imgs 把内嵌资产落盘到 imgs/ 并重写引用（见 convertFileWithAssets）。
package main

import (
	"flag"
	"fmt"
	"net/url"
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
		imgs   = flag.Bool("imgs", false, "write embedded assets into imgs/ next to the Markdown and rewrite references")
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
		var md string
		var err error
		if *imgs {
			// 顺带取内嵌资产落盘并重写引用（一次合并解析，md 与资产同源）。
			// 默认输出（<input>.md 同目录）与 -o 时资产目录跟随 md 位置，引用
			// 相对 md 始终成立；-s（stdout）没有输出位置基准，资产落在输入
			// 目录的 imgs/。
			dir := filepath.Join(filepath.Dir(file), "imgs")
			if *output != "" && *output != "-" {
				dir = filepath.Join(filepath.Dir(*output), "imgs")
			}
			md, err = convertFileWithAssets(conv, file, *format, dir)
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

// convertFileWithAssets 把 file 转成 Markdown，并把内嵌资产写入 dir、重写
// md 里的 asset:// 引用。lib.ConvertFileWithAssets 一次合并解析（core 常驻
// anydoc_convert 导出），md 与资产同源于一个 Document，asset://<id> 与
// Assets[id].ID 一致由结构保证。
func convertFileWithAssets(conv *anydoc.Converter, file, format, dir string) (string, error) {
	res, err := conv.ConvertFileWithAssets(file, format)
	if err != nil {
		return "", err
	}
	return writeAssetRefs(res.Markdown, res.Assets, dir, fileStem(file))
}

// writeAssetRefs 把 assets 逐个写入 dir（文件名 <stem>-<ID>.<ext>），并把 md
// 里的 `asset://<ID>` 占位替换为相对引用。
//
// 引  用固定用 `/` 分隔（跨平台，且 filepath.Join 在 Windows 会产生 `\`
// 导致无法解析）；文件名按 URL 路径段做百分号编码转成 %XX，保证
// `![alt](imgs/xxx.png)` 可被解析。各输入导出到自己的 imgs/，无跨输入覆盖；
// 重复转换幂等（覆盖自己）。占位缺失时报错（而非静默保留坏链接）。
func writeAssetRefs(md string, assets []anydoc.Asset, dir, stem string) (string, error) {
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return md, err
	}
	refs := make(map[string]string, len(assets))
	for _, a := range assets {
		name := fmt.Sprintf("%s-%d.%s", stem, a.ID, assetExt(a.MediaType))
		if err := os.WriteFile(filepath.Join(dir, name), a.Bytes, 0o644); err != nil {
			return md, err
		}
		refs[fmt.Sprintf("asset://%d", a.ID)] = "imgs/" + url.PathEscape(name)
	}
	return rewriteAssetRefs(md, refs)
}

// 取文件没有扩展名的基础名（<stem>），用于资产命名 <stem>-<ID>.<ext>。
func fileStem(file string) string {
	return strings.TrimSuffix(filepath.Base(file), filepath.Ext(file))
}

// rewriteAssetRefs 把 md 图片链接里的 `asset://<id>` 占位替换为 refs 中对应引用。
//
// 只有"真实 markdown 图片链接"里恰好形如 `asset://<连续数字>` 的 dest 才当作
// 占位处理，其余一律按字面保留，替代旧的全文字符串扫描（`strings.Index(…, "asset://")`）：
//   - 正文、代码、URL 里出现的 `asset://123` 不再是占位，不会被误改写，也不会
//     （在 ID 缺失时）误报错——只有出现在 `![alt](…)` 的 dest 位置才被识别；
//   - 整段解析图片链接结构（含 alt 转义与括号配对），占位 ID 取 dest 里一整段
//     连续十进制数字，且必须紧贴 `)` 结束，故 `asset://3` 不会误吞 `asset://31` 前缀；
//   - 真正的占位 ID 在 refs 中缺失时**返回错误**而非静默保留——未导出资产却
//     留有占位引用意味着产物里有坏链接，应显式失败而非产出污染 markdown。
func rewriteAssetRefs(md string, refs map[string]string) (string, error) {
	var sb strings.Builder
	rest := md
	for {
		idx := strings.Index(rest, "![")
		if idx < 0 {
			sb.WriteString(rest)
			return sb.String(), nil
		}
		sb.WriteString(rest[:idx])
		destStart, destEnd, ok := parseImageLink(rest[idx:])
		if !ok { // 不是完整图片链接：`![` 按字面保留，继续扫后面的
			sb.WriteString("![")
			rest = rest[idx+len("!["):]
			continue
		}
		// 逐字节复制图片链接前缀（`![alt](`），保持原文其余不动，仅替换 dest。
		sb.WriteString(rest[idx : idx+destStart])
		dest := rest[idx+destStart : idx+destEnd]
		if idStr, isAsset := assetPlaceholderID(dest); isAsset {
			ref, ok := refs["asset://"+idStr]
			if !ok {
				return "", fmt.Errorf("asset placeholder asset://%s not among extracted assets", idStr)
			}
			sb.WriteString(ref)
		} else { // 普通图片 dest（真实路径等）：原样保留
			sb.WriteString(dest)
		}
		sb.WriteString(")")
		rest = rest[idx+destEnd+1:]
	}
}

// parseImageLink 解析以 `![` 开头的图片链接 `![alt](dest)`，返回 dest 在 s 中的
// 起止下标（不含包裹括号）。非图片链接（alt 未闭合、`]` 后非 `(`、dest 未闭合）
// 返回 ok=false，调用方应把 `![` 按字面保留。
func parseImageLink(s string) (destStart, destEnd int, ok bool) {
	i := len("![")
	depth := 0
	for i < len(s) {
		switch c := s[i]; {
		case c == '\\': // alt 内转义：跳过下一个字符
			i += 2
		case c == '[':
			depth++
			i++
		case c == ']':
			if depth == 0 {
				goto labelDone
			}
			depth--
			i++
		default:
			i++
		}
	}
	return 0, 0, false
labelDone:
	if i+1 >= len(s) || s[i+1] != '(' {
		return 0, 0, false
	}
	destStart = i + 2
	i = destStart
	depth = 0
	for i < len(s) {
		switch c := s[i]; {
		case c == '\\': // dest 内转义：跳过下一个字符
			i += 2
		case c == '(':
			depth++
			i++
		case c == ')':
			if depth == 0 {
				return destStart, i, true
			}
			depth--
			i++
		default:
			i++
		}
	}
	return 0, 0, false // dest 未闭合：不是图片链接
}

// assetPlaceholderID 若 dest 恰好是 `asset://<连续数字>` 则返回其数字 ID；
// 否则返回 isAsset=false，示意该 dest 不是资产占位。
func assetPlaceholderID(dest string) (string, bool) {
	const prefix = "asset://"
	if len(dest) <= len(prefix) || dest[:len(prefix)] != prefix {
		return "", false
	}
	digits := dest[len(prefix):]
	for _, r := range digits {
		if r < '0' || r > '9' {
			return "", false
		}
	}
	return digits, true
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
