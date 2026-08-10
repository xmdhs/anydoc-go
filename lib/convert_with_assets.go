// anydoc_convert build tag：一次性合并转换（Markdown + 内嵌资产）的唯一实现。
//
// 依赖 core 的 AnydocConvert 导出——它只在 wasm2go 转译含 anydoc_convert
// 导出的 wasm 时生成（见 cabi/src/lib.rs），由 CI 重建 core 后才有。因此本
// 文件整体受 anydoc_convert tag 门控：无 tag 的默认构建不引用该符号，仓库
// 始终可编译；CI 重建出 core 后以 -tags anydoc_convert 构建即启用本路径。
//
// core.AnydocConvert 的 Go 签名以当时生成的 core 为准（wasm2go 把每个
// 指针/usize 参数编码为 int32、返回值 void）；若与生成产物不一致，启用
// tag 的集成测试会编译失败——这正是预期的校验点。
//go:build anydoc_convert

package anydoc

import (
	"encoding/binary"
	"errors"
	"os"
	"path/filepath"
	"strings"

	core "github.com/xmdhs/anydoc-go/core"
)

// ConvertResult 是一次合并转换的结果：Markdown 与内嵌资产。它们源自同一个
// Document，因此 Markdown 里的 asset://<id> 与 Assets[id].ID 一一对应。
type ConvertResult struct {
	Markdown string
	Assets   []Asset
}

// ConvertWithAssets 把字节流转为 Markdown 并返回内嵌资产（一次 wasm 调用，
// 单次解析）。format 为扩展名，空串自动从内容检测；无签名格式（CSV）须显式。
// 非线程安全（同 Converter）。
func (c *Converter) ConvertWithAssets(data []byte, format string) (*ConvertResult, error) {
	return c.convertWithAssets(data, strings.TrimPrefix(format, "."))
}

// ConvertFileWithAssets 读取文件并合并转换；format 空时先内容自动检测，
// 检测不出（如 CSV）回退到路径扩展名（与 ConvertFile 的兜底一致）。
func (c *Converter) ConvertFileWithAssets(path, format string) (*ConvertResult, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	if format == "" {
		res, err := c.convertWithAssets(data, "")
		if err != nil {
			if ext := strings.TrimPrefix(filepath.Ext(path), "."); ext != "" {
				if res2, err2 := c.convertWithAssets(data, ext); err2 == nil {
					return res2, nil
				}
			}
		}
		return res, err
	}
	return c.convertWithAssets(data, strings.TrimPrefix(format, "."))
}

// errConvertNotBuilt 用于 AnydocConvert 缺失时的防御（正常不会走到）。
var errConvertNotBuilt = errors.New("anydoc_convert is not available in this build")

// convertWithAssets 调用 anydoc_convert，一次解析同时读取 Markdown 长度与
// 资产流长度（双 out 槽 + out_code）。
func (c *Converter) convertWithAssets(data []byte, format string) (*ConvertResult, error) {
	mdl := c.mdl
	if mdl == nil {
		return nil, errors.New("converter not initialized")
	}

	// 输入分配进线性内存（alloc 可能触发 memory.grow 重赋 m.Memory，因此
	// 每次用前都重新获取）。
	in := core.AnydocAlloc(mdl, int32(len(data)))
	if in == 0 && len(data) > 0 {
		return nil, errOutOfMemory
	}
	defer core.AnydocFree(mdl, in, int32(len(data)))

	// 结果槽（8 字节对齐）：md_ptr(4) md_len(4) assets_ptr(4) assets_len(4)
	// code(4) = 20B，取 24 对齐。
	slots := core.AnydocAlloc(mdl, 24)
	if slots == 0 {
		return nil, errOutOfMemory
	}
	defer core.AnydocFree(mdl, slots, 24)

	mem := core.Memory(mdl)
	if len(data) > 0 {
		copy(mem[in:in+int32(len(data))], data)
	}

	// 可选格式扩展名（NULL = 自动检测）。
	fptr, flen := int32(0), int32(0)
	if format != "" {
		fbuf := core.AnydocAlloc(mdl, int32(len(format)))
		if fbuf == 0 {
			return nil, errOutOfMemory
		}
		defer core.AnydocFree(mdl, fbuf, int32(len(format)))
		mem = core.Memory(mdl) // 上面的 alloc 可能 grow
		copy(mem[fbuf:fbuf+int32(len(format))], format)
		fptr, flen = fbuf, int32(len(format))
	}

	// 一次调用填全部槽；out 参数是各槽在线性内存里的偏移。
	core.AnydocConvert(mdl, in, int32(len(data)), fptr, flen,
		slots, slots+4, slots+8, slots+12, slots+16)
	mem = core.Memory(mdl) // 调用内部可能 grow，取最新线性内存再读结果

	mdPtr := binary.LittleEndian.Uint32(mem[slots : slots+4])
	mdLen := binary.LittleEndian.Uint32(mem[slots+4 : slots+8])
	assetsPtr := binary.LittleEndian.Uint32(mem[slots+8 : slots+12])
	assetsLen := binary.LittleEndian.Uint32(mem[slots+12 : slots+16])
	code := binary.LittleEndian.Uint32(mem[slots+16 : slots+20])

	var res ConvertResult
	if mdLen > 0 {
		md := string(mem[int32(mdPtr) : int32(mdPtr)+int32(mdLen)])
		defer core.AnydocFree(mdl, int32(mdPtr), int32(mdLen))
		if code != 0 {
			return nil, codeError(code, md) // md 槽此时放错误消息
		}
		res.Markdown = md
	}
	if assetsLen > 0 {
		assets := string(mem[int32(assetsPtr) : int32(assetsPtr)+int32(assetsLen)])
		defer core.AnydocFree(mdl, int32(assetsPtr), int32(assetsLen))
		parsed, err := parseAssets(assets)
		if err != nil {
			return nil, err
		}
		res.Assets = parsed
	}
	if code != 0 {
		return nil, codeError(code, "conversion failed")
	}
	return &res, nil
}
