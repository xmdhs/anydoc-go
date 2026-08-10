// 资产落盘与 markdown 引用重写的共享实现，两个 convertFileWithAssets 构建
// 变体（assets_fallback.go / assets_merge.go）复用。
package main

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	anydoc "github.com/xmdhs/anydoc-go/lib"
)

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
