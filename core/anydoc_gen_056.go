package core

import (
	"encoding/binary"
	"math"
	"math/bits"
	_ "embed"
)

func (m *Module) X__wbindgen_describe___wbg_trimLeft_c2d7f626d16eb452() {
	m.X__wbindgen_describe___wbg_valueOf_e5a9f33c39b20482()
}
func (m *Module) X__wbindgen_describe___wbg_trimRight_94c8fd17ab9a9577() {
	m.X__wbindgen_describe___wbg_valueOf_e5a9f33c39b20482()
}
func (m *Module) X__wbindgen_describe___wbg_trimStart_ad580c600ff62e97() {
	m.X__wbindgen_describe___wbg_valueOf_e5a9f33c39b20482()
}
func (m *Module) X__wbindgen_describe___wbg_trim_ac1105accdb11925() {
	m.X__wbindgen_describe___wbg_valueOf_e5a9f33c39b20482()
}
func (m *Module) X__wbindgen_describe___wbg_unit_3c4c20c7a4909e14() {
	m.X__wbindgen_describe___wbg_unit_9bb3e63a284a5fdc()
}
func (m *Module) X__wbindgen_describe___wbg_validate_e1903c15fa10acf1() {
	m.X__wbindgen_describe___wbg_isArray_291e8fbbc73f8b2e()
}
func (m *Module) X__wbindgen_describe___wbg_value_1015d7b8bee3f61a() {
	m.X__wbindgen_describe___wbg_value_ec596f31ec0b1ea0()
}
func (m *Module) X__wbindgen_describe___wbg_value_2f96e40c0b6bdcd0() {
	m.X__wbindgen_describe___wbg_value_ec596f31ec0b1ea0()
}
func (m *Module) X__wbindgen_describe___wbg_value_89ecd7a686af6643() {
	m.X__wbindgen_describe___wbg_value_ec596f31ec0b1ea0()
}
func (m *Module) X__wbindgen_describe___wbg_value_aabe92197590c020() {
	m.X__wbindgen_describe___wbg_value_ec596f31ec0b1ea0()
}
func (m *Module) X__wbindgen_describe___wbg_values_2a12fb5a6064a244() {
	m.X__wbindgen_describe___wbg_entries_f9b99828af86744b()
}
func (m *Module) X__wbindgen_describe___wbg_values_bbaf1bcbcd5d1deb() {
	m.X__wbindgen_describe___wbg_values_d8b566d092c7d444()
}
func (m *Module) X__wbindgen_describe___wbg_values_ec22efea02f28344() {
	m.X__wbindgen_describe___wbg_entries_71601bcb707d5078()
}
func (m *Module) X__wbindgen_describe___wbg_weeks_e0ab147c34cf651d() {
	m.X__wbindgen_describe___wbg_days_2e65b0c3f09c133d()
}
func (m *Module) X__wbindgen_describe___wbg_years_6bd531eab5bbde56() {
	m.X__wbindgen_describe___wbg_days_2e65b0c3f09c133d()
}
func (m *Module) Xmemory() Memory {
	return (*wasmMemory)(&m.memory)
}

//go:nosplit
func i32(x int32) int32 { return x }

//go:nosplit
func i64(x int64) int64 { return x }

//go:nosplit
func i32_shl(x, y int32) int32 {
	return x << (y & 31)
}

//go:nosplit
func i32_shr_u(x, y int32) int32 {
	return int32(uint32(x) >> (y & 31))
}

//go:nosplit
func i64_shl(x, y int64) int64 {
	return x << (y & 63)
}

//go:nosplit
func i64_shr_s(x, y int64) int64 {
	return x >> (y & 63)
}

//go:nosplit
func i64_shr_u(x, y int64) int64 {
	return int64(uint64(x) >> (y & 63))
}

//go:nosplit
func i32_rotl(x, y int32) int32 {
	return int32(bits.RotateLeft32(uint32(x), int(y)))
}

//go:nosplit
func i32_rotr(x, y int32) int32 {
	return int32(bits.RotateLeft32(uint32(x), -int(y)))
}

//go:nosplit
func i64_rotl(x, y int64) int64 {
	return int64(bits.RotateLeft64(uint64(x), int(y)))
}

//go:nosplit
func i64_trunc_sat_f64_s(f float64) int64 {
	switch {
	case f < math.MinInt64:
		return math.MinInt64
	case f >= math.MaxInt64:
		return math.MaxInt64
	case f != f:
		return 0
	}
	return int64(f)
}

//go:nosplit
func i64_trunc_sat_f64_u(f float64) int64 {
	var i uint64
	switch {
	case f <= 0 || f != f:
		i = 0
	case f >= math.MaxUint64:
		i = math.MaxUint64
	default:
		i = uint64(f)
	}
	return int64(i)
}

//go:nosplit
func load16(b []byte) uint16 {
	return binary.LittleEndian.Uint16(b)
}

//go:nosplit
func store16(b []byte, v uint16) {
	binary.LittleEndian.PutUint16(b, v)
}

//go:nosplit
func load32(b []byte) uint32 {
	return binary.LittleEndian.Uint32(b)
}

//go:nosplit
func store32(b []byte, v uint32) {
	binary.LittleEndian.PutUint32(b, v)
}

//go:nosplit
func load64(b []byte) uint64 {
	return binary.LittleEndian.Uint64(b)
}

//go:nosplit
func store64(b []byte, v uint64) {
	binary.LittleEndian.PutUint64(b, v)
}

func memory_grow(mem *[]byte, delta, max int64) int64 {
	buf := *mem
	len := int64(len(buf))
	old := len >> 16
	if delta == 0 {
		return old
	}
	new := old + delta
	add := new<<16 - len
	max = min(max, int64(math.MaxInt)>>16)
	if new > max || new < old || add < 0 {
		return -1
	}
	*mem = append(buf, make([]byte, add)...)
	return old
}

func memory_init[T1, T2 int | uint32 | uint64](mem []byte, data string, dest T1, src, n T2) {
	x := uint64(dest)
	z := uint64(src)
	y := x + uint64(n)
	w := z + uint64(n)
	copy(mem[x:y], data[z:w])
}

func memory_copy[T uint32 | uint64](mem []byte, dest, src, n T) {
	x := uint64(dest)
	z := uint64(src)
	y := x + uint64(n)
	w := z + uint64(n)
	copy(mem[x:y], mem[z:w])
}

func memory_fill[T uint32 | uint64](mem []byte, dest T, val int32, n T) {
	x := uint64(dest)
	y := x + uint64(n)
	buf := mem[x:y]
	if len(buf) > 0 {
		buf[0] = byte(val)
		for i := 1; i < len(buf); {
			chunk := min(i, 8192)
			i += copy(buf[i:], buf[:chunk])
		}
	}
}

func memory_zero[T uint32 | uint64](mem []byte, dest, n T) {
	x := uint64(dest)
	y := x + uint64(n)
	clear(mem[x:y])
}

func table_init[T1, T2, T3 int | int32 | int64](tab, elems []any, dest T1, src T2, n T3) {
	x := uint64(dest)
	z := uint64(src)
	y := x + uint64(n)
	w := z + uint64(n)
	copy(tab[x:y], elems[z:w])
}
//go:embed anydoc.wasm.dat
var data string

