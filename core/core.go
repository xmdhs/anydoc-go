package core

import (
	base "github.com/xmdhs/anydoc-go/core/base"
	"sync"
	"sync/atomic"
	"unsafe"
	_ "github.com/xmdhs/anydoc-go/core/p5"
	_ "embed"
)

func New() *base.Module {
	m := &base.Module{}
	m.Memory = make([]byte, 3145728, 3932160)
	m.MemMu = &sync.Mutex{}
	m.MemSize = &atomic.Uint64{}
	m.Threads = &base.ThreadPool{}
	m.MemSize.Store(3145728)
	m.M = unsafe.Pointer(unsafe.SliceData(m.Memory))
	m.MaxMem = 4294967296
	m.T0 = make([]any, 590)
	m.G0 = int32(1048576)
	InitElemSeg_1_0(m)
	InitElemSeg_2_0(m)
	InitElemSeg_3_0(m)
	InitElemSeg_4_0(m)
	InitElemSeg_5_0(m)
	InitElemSeg_5_1(m)
	m.DataEnd = 3124300
	initData_0(m)
	return m
}
func NewWithMemory(memory []byte, memSize uint64) *base.Module {
	m := &base.Module{}
	m.Memory = memory
	m.MemMu = &sync.Mutex{}
	m.MemSize = &atomic.Uint64{}
	m.Threads = &base.ThreadPool{}
	if memSize > 4294836224 {
		panic("wasm2go: memory size exceeds the implementation limit (4294836224 bytes)")
	}
	m.MemSize.Store(memSize)
	m.M = unsafe.Pointer(unsafe.SliceData(m.Memory))
	m.MaxMem = uint64(len(memory))
	m.T0 = make([]any, 590)
	m.G0 = int32(1048576)
	InitElemSeg_1_0(m)
	InitElemSeg_2_0(m)
	InitElemSeg_3_0(m)
	InitElemSeg_4_0(m)
	InitElemSeg_5_0(m)
	InitElemSeg_5_1(m)
	m.DataEnd = 3124300
	return m
}
func NewFromSnapshot(memory []byte, memSize uint64, globals []uint64) *base.Module {
	m := &base.Module{}
	m.Memory = memory
	m.MemMu = &sync.Mutex{}
	m.MemSize = &atomic.Uint64{}
	m.Threads = &base.ThreadPool{}
	if memSize > 4294836224 {
		panic("wasm2go: memory size exceeds the implementation limit (4294836224 bytes)")
	}
	m.MemSize.Store(memSize)
	m.M = unsafe.Pointer(unsafe.SliceData(m.Memory))
	m.MaxMem = uint64(len(memory))
	m.T0 = make([]any, 590)
	m.G0 = int32(1048576)
	InitElemSeg_1_0(m)
	InitElemSeg_2_0(m)
	InitElemSeg_3_0(m)
	InitElemSeg_4_0(m)
	InitElemSeg_5_0(m)
	InitElemSeg_5_1(m)
	m.DataEnd = 3124300
	base.RestoreGlobals(m, globals)
	return m
}
func initData_0(m *base.Module) {
	copy(m.Memory[1048576:], wasm2goData_data_bin[0:2075724])
}
func AnydocAlloc(m *base.Module, l0 int32) int32 {
	return Fn17(m, l0)
}
func AnydocAssets(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	return Fn18(m, l0, l1, l2, l3, l4, l5)
}
func AnydocConvert(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) {
	Fn23(m, l0, l1, l2, l3, l4, l5, l6, l7, l8)
}
func AnydocFree(m *base.Module, l0 int32, l1 int32) {
	Fn25(m, l0, l1)
}
func AnydocToMarkdown(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	return Fn26(m, l0, l1, l2, l3, l4, l5)
}
func Memory(m *base.Module) []byte {
	return m.Memory
}
//go:embed data.bin
var wasm2goData_data_bin []byte
