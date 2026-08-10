package core

import (
	base "github.com/xmdhs/anydoc-go/core/base"
	"sync"
	"sync/atomic"
	"unsafe"
	_ "github.com/xmdhs/anydoc-go/core/p1"
	_ "embed"
)

func New() *base.Module {
	m := &base.Module{}
	m.Memory = make([]byte, 1310720, 1638400)
	m.MemMu = &sync.Mutex{}
	m.MemSize = &atomic.Uint64{}
	m.Threads = &base.ThreadPool{}
	m.MemSize.Store(1310720)
	m.M = unsafe.Pointer(unsafe.SliceData(m.Memory))
	m.MaxMem = 4294967296
	m.T0 = make([]any, 147)
	m.G0 = int32(1048576)
	InitElemSeg_0_0(m)
	InitElemSeg_1_0(m)
	m.DataEnd = 1293520
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
	m.T0 = make([]any, 147)
	m.G0 = int32(1048576)
	InitElemSeg_0_0(m)
	InitElemSeg_1_0(m)
	m.DataEnd = 1293520
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
	m.T0 = make([]any, 147)
	m.G0 = int32(1048576)
	InitElemSeg_0_0(m)
	InitElemSeg_1_0(m)
	m.DataEnd = 1293520
	base.RestoreGlobals(m, globals)
	return m
}
func initData_0(m *base.Module) {
	copy(m.Memory[1048576:], wasm2goData_data_bin[0:244944])
}
func AnydocAlloc(m *base.Module, l0 int32) int32 {
	return Fn16(m, l0)
}
func AnydocAssets(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	return Fn17(m, l0, l1, l2, l3, l4, l5)
}
func AnydocFree(m *base.Module, l0 int32, l1 int32) {
	Fn22(m, l0, l1)
}
func AnydocToMarkdown(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	return Fn23(m, l0, l1, l2, l3, l4, l5)
}
func Memory(m *base.Module) []byte {
	return m.Memory
}
//go:embed data.bin
var wasm2goData_data_bin []byte
