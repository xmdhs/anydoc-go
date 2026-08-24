// Package anydoc 把 anydoc 文档转换核心（goccy/wasm2go 生成的 core 包）封装
// 成一个简单、并发安全的 Go API。
//
// 用法：
//
//	conv := anydoc.NewConverter()
//	md, err := conv.ConvertFile("report.docx", "")   // 自动检测格式
//	md, err = conv.ConvertBytes(data, "csv")         // 或显式指定扩展名
//
// format 传空字符串表示从内容自动检测；CSV 等无签名格式必须显式命名。
// 转换错误是 *ConvertError，带有稳定的 Code（"unsupported"、"malformed"、
// "encrypted"、"resourceLimit"、"missingPart"、"io"、"other"）。
package anydoc

import (
	"encoding/binary"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	core "github.com/xmdhs/anydoc-go/core"
	corebase "github.com/xmdhs/anydoc-go/core/base"
)

// ConvertError 是一次转换失败的结构化错误。Code 是稳定机器可读的标识，
// 其值与 anydoc 官方绑定发布的 error.code 一致；Msg 是给人读的文本。
type ConvertError struct {
	Code string
	Msg  string
}

func (e *ConvertError) Error() string { return e.Msg }

// 错误码（与 cabi/src/lib.rs 的 ERR_* 对应）。
func codeName(code uint32) string {
	switch code {
	case 0:
		return "ok"
	case 1:
		return "unsupported"
	case 2:
		return "malformed"
	case 3:
		return "encrypted"
	case 4:
		return "resourceLimit"
	case 5:
		return "missingPart"
	case 6:
		return "io"
	default:
		return "other"
	}
}

func codeError(code uint32, msg string) *ConvertError {
	return &ConvertError{Code: codeName(code), Msg: msg}
}

// Converter 拥有一个 goccy 生成的 wasm 模块（线性内存 + 全局栈指针），可
// 反复转换任意文档。
//
// 非线程安全：同一实例同一时间只能被一个 goroutine 使用；并发转换请
// 为每个 goroutine 各自 NewConverter（各实例内存独立，天然并行）。
type Converter struct {
	mdl *corebase.Module
}

// NewConverter 创建并初始化转换器（初始化线性内存、载入内嵌数据）。
// 创建成本约 1ms 级，长期复用一个实例。
func NewConverter() *Converter {
	return &Converter{mdl: core.New()}
}

// 错误消息词表，供下面辅助函数使用。
var errOutOfMemory = errors.New("out of memory")

// ConvertBytes 把内存中的文档字节流转为 Markdown。非线程安全。
// format 为扩展名（"docx"、"csv"、…，可带点），空串自动从内容检测；
// 无签名格式（CSV）必须显式给出。HTML（含伪装成 .doc 的 Word 导出 HTML）
// 走纯 Go 的 html-to-markdown 分支（GFM：表格、删除线）。
func (c *Converter) ConvertBytes(data []byte, format string) (string, error) {
	ext := strings.TrimPrefix(format, ".")
	if isHTMLFormat(ext) {
		return htmlToMarkdown(data)
	}
	if ext == "" && isHTMLData(data) {
		return htmlToMarkdown(data)
	}
	// 显式 doc 但内容实为 HTML 时，视为 HTML（兼容 /mnt/hdd/docs 样本）。
	if strings.EqualFold(ext, "doc") && isHTMLData(data) {
		return htmlToMarkdown(data)
	}
	return c.convert(data, ext)
}

// ConvertFile 读取文件并转换为 Markdown。非线程安全。
//
// format 为空时先按文件内容自动检测；检测不出来的无签名格式（CSV）会
// 用路径扩展名兜底（与 anydoc 核心库 to_markdown(path) 的行为一致）；
// format 非空时直接按该扩展名解析。HTML（含伪装成 .doc 的 Word 导出 HTML）
// 与 .html/.htm 显式格式走纯 Go 的 html-to-markdown 分支（GFM：表格、删除线）。
func (c *Converter) ConvertFile(path, format string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	ext := strings.TrimPrefix(format, ".")
	if ext != "" {
		if isHTMLFormat(ext) {
			return htmlToMarkdown(data)
		}
		if strings.EqualFold(ext, "doc") && isHTMLData(data) {
			return htmlToMarkdown(data)
		}
		return c.convert(data, ext)
	}
	// 未指定格式：先看内容是否像 HTML（含假 .doc），再走原 wasm 路径。
	if isHTMLData(data) {
		return htmlToMarkdown(data)
	}
	md, err := c.convert(data, "")
	if err != nil {
		// 内容检测不出（如 CSV / 伪装 HTML 未被上分支捕获），回退到文件扩展名。
		if fileExt := strings.TrimPrefix(filepath.Ext(path), "."); fileExt != "" {
			if isHTMLFormat(fileExt) {
				if md2, err2 := htmlToMarkdown(data); err2 == nil {
					return md2, nil
				}
			}
			if strings.EqualFold(fileExt, "doc") && isHTMLData(data) {
				if md2, err2 := htmlToMarkdown(data); err2 == nil {
					return md2, nil
				}
			}
			if md2, err2 := c.convert(data, fileExt); err2 == nil {
				return md2, nil
			}
		}
	}
	return md, err
}

// Asset 是文档内嵌的二进制资产（图片、嵌入对象），对应渲染占位符
// `![alt](asset://<ID>)`；ID 即 Document::assets 的下标。
type Asset struct {
	ID        uint32
	MediaType string
	Bytes     []byte
}

// ExtractAssets 解析文档并返回内嵌资产（字节原样保留）。与 ConvertBytes
// 同一输入协议；markdown 里的图片以 `![alt](asset://<ID>)` 占位，用返回
// 的 ID/MediaType/Bytes 落盘后可自行替换引用。非线程安全（同 Converter）。
// HTML 分支无内嵌资产概念（图片为外链），返回空切片。
func (c *Converter) ExtractAssets(data []byte, format string) ([]Asset, error) {
	ext := strings.TrimPrefix(strings.TrimSpace(format), ".")
	if isHTMLFormat(ext) || (ext == "" && isHTMLData(data)) || (strings.EqualFold(ext, "doc") && isHTMLData(data)) {
		return []Asset{}, nil
	}
	mdl := c.mdl
	if mdl == nil {
		return nil, errors.New("converter not initialized")
	}

	// 输入分配进线性内存（goccy 的 Memory 是 slice 值：alloc 可能触发
	// memory.grow 重赋 m.Memory，因此每次用到前都要重新获取）。
	in := core.AnydocAlloc(mdl, int32(len(data)))
	if in == 0 && len(data) > 0 {
		return nil, errOutOfMemory
	}
	defer core.AnydocFree(mdl, in, int32(len(data)))

	// 结果槽：out_len(4) + out_code(4)，8 字节对齐。
	slots := core.AnydocAlloc(mdl, 8)
	if slots == 0 {
		return nil, errOutOfMemory
	}
	defer core.AnydocFree(mdl, slots, 8)

	mem := core.Memory(mdl)
	if len(data) > 0 {
		copy(mem[in:in+int32(len(data))], data)
	}

	// 可选格式：扩展名字节（NULL = 自动检测）。
	var fptr, flen int32
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

	buf := core.AnydocAssets(mdl, in, int32(len(data)), fptr, flen, slots, slots+4)
	mem = core.Memory(mdl) // 调用内部可能 grow，取最新线性内存再读结果
	outLen := binary.LittleEndian.Uint32(mem[slots : slots+4])
	code := binary.LittleEndian.Uint32(mem[slots+4 : slots+8])

	result := ""
	if outLen > 0 && buf != 0 {
		result = string(mem[buf : buf+int32(outLen)])
	}
	core.AnydocFree(mdl, buf, int32(outLen))

	if code != 0 {
		if result == "" {
			result = "conversion failed"
		}
		return nil, codeError(code, result)
	}
	return parseAssets(result)
}

// parseAssets 解析 anydoc_assets 的序列化流
// （u32 数量 + 每项 u32 id / u32 media_len + bytes / u32 bytes_len + bytes）。
func parseAssets(stream string) ([]Asset, error) {
	b, pos := []byte(stream), 0
	read32 := func() (uint32, bool) {
		if pos+4 > len(b) {
			return 0, false
		}
		v := binary.LittleEndian.Uint32(b[pos : pos+4])
		pos += 4
		return v, true
	}
	readBytes := func() ([]byte, bool) {
		n, ok := read32()
		if !ok || pos+int(n) > len(b) {
			return nil, false
		}
		v := b[pos : pos+int(n)]
		pos += int(n)
		return v, true
	}
	count, ok := read32()
	if !ok {
		return nil, errors.New("assets stream truncated")
	}
	assets := make([]Asset, 0, count)
	for range count {
		id, ok := read32()
		if !ok {
			return nil, errors.New("assets stream truncated")
		}
		mt, ok := readBytes()
		if !ok {
			return nil, errors.New("assets stream truncated")
		}
		bytes_, ok := readBytes()
		if !ok {
			return nil, errors.New("assets stream truncated")
		}
		assets = append(assets, Asset{ID: id, MediaType: string(mt), Bytes: bytes_})
	}
	return assets, nil
}

// ConvertFileTo 是 ConvertFile + 写出 Markdown 的便捷封装；非线程安全。
func (c *Converter) ConvertFileTo(path, format, outPath string) error {
	md, err := c.ConvertFile(path, format)
	if err != nil {
		return err
	}
	return os.WriteFile(outPath, []byte(md), 0o644)
}

// convert 核心转换（未加锁；调用方遵守非线程安全约定）。
func (c *Converter) convert(data []byte, format string) (string, error) {
	mdl := c.mdl
	if mdl == nil {
		return "", errors.New("converter not initialized")
	}

	// 输入分配进线性内存（goccy 的 Memory 是 slice 值：alloc 可能触发
	// memory.grow 重赋 m.Memory，因此每次用到前都要重新获取）。
	in := core.AnydocAlloc(mdl, int32(len(data)))
	if in == 0 && len(data) > 0 {
		return "", errOutOfMemory
	}
	defer core.AnydocFree(mdl, in, int32(len(data)))

	// 结果槽：out_len(4) + out_code(4)，8 字节对齐。
	slots := core.AnydocAlloc(mdl, 8)
	if slots == 0 {
		return "", errOutOfMemory
	}
	defer core.AnydocFree(mdl, slots, 8)

	mem := core.Memory(mdl)
	if len(data) > 0 {
		copy(mem[in:in+int32(len(data))], data)
	}

	// 可选格式：扩展名字节（NULL = 自动检测）。
	var fptr, flen int32
	if format != "" {
		fbuf := core.AnydocAlloc(mdl, int32(len(format)))
		if fbuf == 0 {
			return "", errOutOfMemory
		}
		defer core.AnydocFree(mdl, fbuf, int32(len(format)))
		mem = core.Memory(mdl) // 上面的 alloc 可能 grow
		copy(mem[fbuf:fbuf+int32(len(format))], format)
		fptr, flen = fbuf, int32(len(format))
	}

	// 调用 wasm 转换；槽接收（长度, 错误码）。
	buf := core.AnydocToMarkdown(mdl, in, int32(len(data)), fptr, flen, slots, slots+4)
	mem = core.Memory(mdl) // 转换内部可能 grow，取最新线性内存再读结果
	outLen := binary.LittleEndian.Uint32(mem[slots : slots+4])
	code := binary.LittleEndian.Uint32(mem[slots+4 : slots+8])

	result := ""
	if outLen > 0 {
		if buf == 0 {
			return "", fmt.Errorf("conversion produced length %d but no buffer (code %d)", outLen, code)
		}
		result = string(mem[buf : buf+int32(outLen)])
	}
	core.AnydocFree(mdl, buf, int32(outLen))

	if code != 0 {
		if result == "" {
			return "", codeError(code, "conversion failed")
		}
		return "", codeError(code, result)
	}
	return result, nil
}
