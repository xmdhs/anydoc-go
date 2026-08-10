// 默认构建变体（无 anydoc_convert tag）：当 core 尚无合并导出时的实现。
//
// 这里用 `{ConvertFile（拿 md）+ ExtractAssets（取资产）}` 两次 wasm 解析，
// 逐字节行为与合并变体一致；一旦 CI 用 anydoc_convert tag 重建出
// core.AnydocConvert，改用 assets_merge.go 的一次解析路径。
//go:build !anydoc_convert

package main

import (
	"os"
	"path/filepath"
	"strings"

	anydoc "github.com/xmdhs/anydoc-go/lib"
)

// convertFileWithAssets 把 file 转成 Markdown，并把内嵌资产写入 dir、重写
// md 里的 asset:// 引用。无 tag 变体：{ConvertFile + ExtractAssets} 两次
// 解析（各自带格式兜底）。
func convertFileWithAssets(conv *anydoc.Converter, file, format, dir string) (string, error) {
	md, err := conv.ConvertFile(file, format) // 自带格式兜底
	if err != nil {
		return "", err
	}
	data, err := os.ReadFile(file)
	if err != nil {
		return md, err
	}
	// 与 ConvertFile 相同的兜底：显式 -f → 内容自动检测 → 路径扩展名。
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
	return writeAssetRefs(md, assets, dir, fileStem(file))
}
