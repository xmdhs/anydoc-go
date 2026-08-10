// convertFileWithAssets 唯一实现：core 常驻合并导出 AnydocConvert（CI 重建
// core 时由 wasm 的 anydoc_convert 导出生成），一次 wasm 解析同时拿到
// Markdown 与资产，无任何 build tag 门控。
//
// lib.ConvertFileWithAssets 的实现见 lib/convert_with_assets.go。
package main

import (
	anydoc "github.com/xmdhs/anydoc-go/lib"
)

// convertFileWithAssets 把 file 转成 Markdown，并把内嵌资产写入 dir、重写
// md 里的 asset:// 引用。一次 ConvertFileWithAssets 合并解析，md 与资产
// 同源于一个 Document，asset://<id> 与 Assets[id].ID 一致由结构保证。
func convertFileWithAssets(conv *anydoc.Converter, file, format, dir string) (string, error) {
	res, err := conv.ConvertFileWithAssets(file, format)
	if err != nil {
		return "", err
	}
	return writeAssetRefs(res.Markdown, res.Assets, dir, fileStem(file))
}
