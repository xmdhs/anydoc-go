// Package anydoc 把 anydoc 文档转换核心（wasm2go 生成的 core 包）封装成
// 一个简单、并发安全的 Go API。
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

	"anydoc-go/core"
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

// Converter 拥有一个 wasm 转换模块（线性内存 + 全局栈指针），可反复
// 转换任意文档。
//
// 非线程安全：同一实例同一时间只能被一个 goroutine 使用；并发转换请
// 为每个 goroutine 各自 NewConverter（各实例内存独立，天然并行）。
type Converter struct {
	mdl *core.Module
	env hostEnv
}

// NewConverter 创建并初始化转换器（加载 wasm 模块、填充线性内存）。
// 创建成本约数十毫秒，长期复用一个实例。
func NewConverter() *Converter {
	c := new(Converter)
	c.mdl = core.New(&c.env, &extrefEnv{})
	return c
}

// 错误消息词表，供下面辅助函数使用。
var errOutOfMemory = errors.New("out of memory")

// ConvertBytes 把内存中的文档字节流转为 Markdown。非线程安全。
// format 为扩展名（"docx"、"csv"、…，可带点），空串自动从内容检测；
// 无签名格式（CSV）必须显式给出。
func (c *Converter) ConvertBytes(data []byte, format string) (string, error) {
	return c.convert(data, strings.TrimPrefix(format, "."))
}

// ConvertFile 读取文件并转换为 Markdown。非线程安全。
//
// format 为空时先按文件内容自动检测；检测不出来的无签名格式（CSV）会
// 用路径扩展名兜底（与 anydoc 核心库 to_markdown(path) 的行为一致）；
// format 非空时直接按该扩展名解析。
func (c *Converter) ConvertFile(path, format string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	if format == "" {
		md, err := c.convert(data, "")
		if err != nil {
			// 内容检测不出（如 CSV），回退到文件扩展名。
			if ext := strings.TrimPrefix(filepath.Ext(path), "."); ext != "" {
				if md2, err2 := c.convert(data, ext); err2 == nil {
					return md2, nil
				}
			}
		}
		return md, err
	}
	return c.convert(data, strings.TrimPrefix(format, "."))
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
	mem := mdl.Xmemory().Slice()

	// 输入分配进线性内存。
	in := mdl.Xanydoc_alloc(int32(len(data)))
	if in == 0 && len(data) > 0 {
		return "", errOutOfMemory
	}
	defer mdl.Xanydoc_free(in, int32(len(data)))

	// 结果槽：out_len(4) + out_code(4)，8 字节对齐。
	slots := mdl.Xanydoc_alloc(8)
	if slots == 0 {
		return "", errOutOfMemory
	}
	defer mdl.Xanydoc_free(slots, 8)

	if len(data) > 0 {
		copy((*mem)[in:in+int32(len(data))], data)
	}

	// 可选格式：扩展名字节（NULL = 自动检测）。
	var fptr, flen int32
	if format != "" {
		fbuf := mdl.Xanydoc_alloc(int32(len(format)))
		if fbuf == 0 {
			return "", errOutOfMemory
		}
		defer mdl.Xanydoc_free(fbuf, int32(len(format)))
		copy((*mem)[fbuf:fbuf+int32(len(format))], format)
		fptr, flen = fbuf, int32(len(format))
	}

	// 调用 wasm 转换；槽接收（长度, 错误码）。
	buf := mdl.Xanydoc_to_markdown(in, int32(len(data)), fptr, flen, slots, slots+4)
	outLen := binary.LittleEndian.Uint32((*mem)[slots : slots+4])
	code := binary.LittleEndian.Uint32((*mem)[slots+4 : slots+8])

	result := ""
	if outLen > 0 {
		if buf == 0 {
			return "", fmt.Errorf("conversion produced length %d but no buffer (code %d)", outLen, code)
		}
		result = string((*mem)[buf : buf+int32(outLen)])
	}
	mdl.Xanydoc_free(buf, int32(outLen))

	if code != 0 {
		if result == "" {
			return "", codeError(code, "conversion failed")
		}
		return "", codeError(code, result)
	}
	return result, nil
}