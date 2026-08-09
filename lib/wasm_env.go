package anydoc

import "anydoc-go/core"

// hostEnv 实现 core 包要求的 wasm-bindgen imports（panic 上报与类型描述
// 钩子）。转换路径从不调用它们，所以是惰性 stub；若 core 生成物更新后
// 接口变化，以此文件（和 core/anydoc_gen_*.go 的接口定义）为准。
type hostEnv struct {
	m *core.Module
}

func (h *hostEnv) Init(m any) { h.m = m.(*core.Module) }

func (h *hostEnv) X__wbg___wbindgen_throw_bb96b2010945f0bc(v0, v1 int32) {
	panic("anydoc wasm: unrecoverable wasm-bindgen throw")
}

func (h *hostEnv) X__wbindgen_describe(v0 int32) {}

// extrefEnv 实现 externref 表 shim；转换不经过它。
type extrefEnv struct{}

func (extrefEnv) X__wbindgen_externref_table_grow(v0 int32) int32 { return 0 }
func (extrefEnv) X__wbindgen_externref_table_set_null(v0 int32)   {}