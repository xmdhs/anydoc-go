// anydoc_convert tag 构建变体：core 已由 CI 重建出合并导出 AnydocConvert 时
// 启用。一次 wasm 解析同时拿到 Markdown 与资产，避免第二次解析。
//
// lib.ConvertFileWithAssets 的实现见 lib/convert_with_assets.go，依赖 core 的
// AnydocConvert 符号——该符号仅在 wasm2go 转译含 anydoc_convert 导出的 wasm
// 时生成，故本文件整体由 build tag 隔离，保证无 tag 的默认构建不引用它。
//go:build anydoc_convert

package main

import (
	anydoc "github.com/xmdhs/anydoc-go/lib"
)

// convertFileWithAssets 把 file 转成 Markdown，并把内嵌资产写入 dir、重写
// md 里的 asset:// 引用。tag 变体：一次 ConvertFileWithAssets 合并解析，
// md 与资产同源于一个 Document，asset://<id> 与 Assets[id].ID 一致由
// 结构保证。
func convertFileWithAssets(conv *anydoc.Converter, file, format, dir string) (string, error) {
	res, err := conv.ConvertFileWithAssets(file, format)
	if err != nil {
		return "", err
	}
	return writeAssetRefs(res.Markdown, res.Assets, dir, fileStem(file))
}
