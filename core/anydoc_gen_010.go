package core

import (
	"math"
)

func (m *Module) fn402(v0, v1 int32) {
	var v2 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		v2 = t0
		t1 := int32(load32(m.memory[uint32(v0):]))
		if v2 != t1 {
			goto l0
		}
		m.fn260(v0)
	}
l0:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v2+i32(1)))
	t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	store32(m.memory[uint32(t2+v2<<2):], uint32(v1))
}
func (m *Module) fn403(v0, v1 int32) {
	var v2, v3 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	if uint32(v0) < uint32(i32(26)) {
		goto l0
	}
	store32(m.memory[int64(uint32(v2))+12:], uint32(i32(0)))
	store64(m.memory[int64(uint32(v2))+4:], uint64(i64(0x100000000)))
l3:
	{
		if uint32(v0) > uint32(i32(25)) {
			t5 := int32(uint32(v0) / uint32(i32(26)))
			t6 := v2 + i32(4)
			t7 := v0
			v3 = t5
			m.fn74(t6, t7-v3*i32(26)+i32(65))
			v0 = v3
			goto l3
		}
		t1 := int32(load32(m.memory[int64(uint32(v2))+8:]))
		t2 := v1
		v0 = t1
		t3 := int32(load32(m.memory[int64(uint32(v2))+12:]))
		m.fn831(t2, v0, v0+t3)
		t4 := int32(load32(m.memory[int64(uint32(v2))+4:]))
		m.fn16(t4, v0)
		goto l2
	}
l0:
	m.fn74(v1, v0+i32(65))
l2:
	m.g0 = v2 + i32(16)
}
func (m *Module) fn404(v0, v1, v2 int32) int32 {
	t0 := m.fn100(v0, i32(1100568), v1, v2)
	return t0
}
func (m *Module) fn405(v0, v1 int32) int32 {
	var v2, v3 int32
	var v4 float64
	t0 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	v2 = t0
	v3 = v2 & i32(0x200000)
	t1 := math.Float64frombits(load64(m.memory[uint32(v0):]))
	v4 = t1
	{
		if v2&i32(0x10000000) != 0 {
			t6 := int32(load16(m.memory[int64(uint32(v1))+14:]))
			t7 := v1
			t8 := v4
			var p9 int32
			if v3 != i32(0) {
				p9 = 1
			}
			t10 := m.fn1676(t7, t8, p9, t6)
			return t10
		}
		t2 := v1
		t3 := v4
		var p4 int32
		if v3 != i32(0) {
			p4 = 1
		}
		t5 := m.fn1674(t2, t3, p4, i32(0))
		return t5
	}
}
func (m *Module) fn406(v0, v1, v2 int32) {
	if v1 == 0 {
		goto l0
	}
	m.memory[uint32(v0)] = byte(i32(255))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
	m.fn409(v2)
	return
l0:
	t0 := int64(load64(m.memory[int64(uint32(v2))+16:]))
	store64(m.memory[int64(uint32(v0))+16:], uint64(t0))
	t1 := int64(load64(m.memory[int64(uint32(v2))+8:]))
	store64(m.memory[int64(uint32(v0))+8:], uint64(t1))
	t2 := int64(load64(m.memory[uint32(v2):]))
	store64(m.memory[uint32(v0):], uint64(t2))
}
func (m *Module) fn407(v0, v1, v2, v3 int32) {
	var v4 int32
	t0 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	v4 = t0
	if v1 == 0 {
		goto l0
	}
	{
		if uint32(v4) > uint32(v1) {
			goto l1
		}
		if v4 != v1 {
			goto l2
		}
		goto l0
	l1:
		t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		t2 := int32(int8(m.memory[uint32(t1+v1)]))
		if t2 > i32(-65) {
			goto l0
		}
	}
l2:
	m.fn256(i32(1087792), i32(44), v3)
	panic("unreachable")
l0:
	m.fn47(v0, i32(1))
	t3 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	v3 = t3 + v1
	v1 = v4 - v1
	if v1 == 0 {
		goto l3
	}
	memory_copy(m.memory, uint32(v3+i32(1)), uint32(v3), uint32(v1))
l3:
	m.fn279(v2, v3)
	store32(m.memory[int64(uint32(v0))+8:], uint32(v4+i32(1)))
}
func (m *Module) fn408(v0, v1, v2, v3 int32) {
	var v4, v5, v6, v7 int32
	t0 := m.g0
	v4 = t0 - i32(16)
	m.g0 = v4
	t1 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	v5 = t1
	t2 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	v6 = t2
	if v2 == 0 {
		goto l0
	}
	{
		if uint32(v5) > uint32(v2) {
			goto l1
		}
		if v5 != v2 {
			goto l2
		}
		goto l0
	l1:
		t3 := int32(int8(m.memory[uint32(v6+v2)]))
		if t3 > i32(-65) {
			goto l0
		}
	}
l2:
	m.fn256(i32(1087884), i32(43), v3)
	panic("unreachable")
l0:
	{
		if uint32(v5) < uint32(v2) {
			m.fn99(v2, v5, v3)
			panic("unreachable")
		}
		t4 := v4 + i32(8)
		v5 = v5 - v2
		m.fn382(t4, v5, i32(1), i32(1))
		t5 := int32(load32(m.memory[int64(uint32(v4))+8:]))
		v7 = t5
		t6 := int32(load32(m.memory[int64(uint32(v4))+12:]))
		v3 = t6
		store32(m.memory[int64(uint32(v1))+8:], uint32(v2))
		if v5 == 0 {
			goto l4
		}
		memory_copy(m.memory, uint32(v3), uint32(v6+v2), uint32(v5))
	l4:
		store32(m.memory[int64(uint32(v0))+8:], uint32(v5))
		store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
		store32(m.memory[uint32(v0):], uint32(v7))
		m.g0 = v4 + i32(16)
		return
	}
}
func (m *Module) fn409(v0 int32) {
	var v1 int32
	{
		t0 := int32(m.memory[uint32(v0)])
		v1 = t0
		switch v1 {
		default:
			if v1 == i32(13) {
				goto l4
			}
			return
		case 0:
			t1 := int32(m.memory[int64(uint32(v0))+4])
			t2 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			m.fn1565(t1, t2)
			return
		case 1:
			m.fn1569(v0 + i32(4))
			return
		case 2:
			m.fn1567(v0 + i32(4))
			return
		}
	}
l4:
	t3 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t4 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	m.fn16(t3, t4)
}
func (m *Module) fn410(v0, v1, v2 int32) {
	if v1 == 0 {
		goto l0
	}
	m.memory[uint32(v0)] = byte(i32(255))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
	m.fn409(v2)
	return
l0:
	t0 := int64(load64(m.memory[int64(uint32(v2))+16:]))
	store64(m.memory[int64(uint32(v0))+16:], uint64(t0))
	t1 := int64(load64(m.memory[int64(uint32(v2))+8:]))
	store64(m.memory[int64(uint32(v0))+8:], uint64(t1))
	t2 := int64(load64(m.memory[uint32(v2):]))
	store64(m.memory[uint32(v0):], uint64(t2))
}
func (m *Module) fn411(v0, v1, v2, v3 int32) {
	var v4, v5, v6 int32
	t0 := m.g0
	v4 = t0 - i32(16)
	m.g0 = v4
	{
		t1 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		v5 = t1
		if uint32(v5) < uint32(v2) {
			m.fn99(v2, v5, v3)
			panic("unreachable")
		}
		t2 := v4 + i32(8)
		v5 = v5 - v2
		m.fn382(t2, v5, i32(4), i32(4))
		t3 := int32(load32(m.memory[int64(uint32(v4))+8:]))
		v3 = t3
		t4 := int32(load32(m.memory[int64(uint32(v4))+12:]))
		t5 := v0
		v6 = t4
		store32(m.memory[int64(uint32(t5))+4:], uint32(v6))
		store32(m.memory[uint32(v0):], uint32(v3))
		store32(m.memory[int64(uint32(v1))+8:], uint32(v2))
		store32(m.memory[int64(uint32(v0))+8:], uint32(v5))
		{
			v0 = v5 << 2
			if v0 == 0 {
				goto l1
			}
			t6 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			memory_copy(m.memory, uint32(v6), uint32(t6+v2<<2), uint32(v0))
		}
	l1:
		m.g0 = v4 + i32(16)
		return
	}
}
func (m *Module) fn412(v0, v1, v2 int32) int32 {
	if v1 != 0 {
		return v0
	}
	m.fn158(i32(0), i32(0), v2)
	panic("unreachable")
}
func (m *Module) fn413(v0, v1 int32) {
	m.fn1080(v0, v1, i32(4), i32(4))
}
func (m *Module) fn414(v0 int32) {
	var v1, v2, v3 int32
	t0 := m.g0
	v1 = t0 - i32(16)
	m.g0 = v1
	t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t2 := v1
	v2 = t1
	store32(m.memory[int64(uint32(t2))+8:], uint32(v2))
	t3 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	t4 := v1
	t5 := v2
	v3 = t3
	store32(m.memory[int64(uint32(t4))+12:], uint32(t5+v3))
	m.fn577(v1, v1+i32(8))
	{
		t6 := int32(load32(m.memory[uint32(v1):]))
		if t6 != i32(1) {
			goto l0
		}
		{
			{
				t7 := int32(load32(m.memory[int64(uint32(v1))+4:]))
				v2 = t7
				if uint32(v2) >= uint32(i32(128)) {
					goto l1
				}
				v2 = i32(-1)
				goto l2
			}
		l1:
			if uint32(v2) >= uint32(i32(2048)) {
				goto l3
			}
			v2 = i32(-2)
			goto l2
		l3:
			p8 := i32(-4)
			if uint32(v2) < uint32(i32(65536)) {
				p8 = i32(-3)
			}
			v2 = p8
		}
	l2:
		store32(m.memory[int64(uint32(v0))+8:], uint32(v2+v3))
	}
l0:
	m.g0 = v1 + i32(16)
}
func (m *Module) fn415(v0, v1, v2, v3, v4, v5 int32) {
	if uint32(v4) < uint32(v3) {
		goto l0
	}
	{
		if v3 == 0 {
			goto l1
		}
		if uint32(v3) < uint32(v2) {
			goto l2
		}
		if v3 != v2 {
			goto l0
		}
		goto l1
	l2:
		t0 := int32(int8(m.memory[uint32(v1+v3)]))
		if t0 <= i32(-65) {
			goto l0
		}
	}
l1:
	{
		if v4 == 0 {
			goto l3
		}
		if uint32(v4) < uint32(v2) {
			goto l4
		}
		if v4 == v2 {
			goto l3
		}
		goto l0
	l4:
		t1 := int32(int8(m.memory[uint32(v1+v4)]))
		if t1 <= i32(-65) {
			goto l0
		}
	}
l3:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v4-v3))
	store32(m.memory[uint32(v0):], uint32(v1+v3))
	return
l0:
	m.fn556(v1, v2, v3, v4, v5)
	panic("unreachable")
}
func (m *Module) fn416(v0, v1 int32) int32 {
	var v2 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	{
		t1 := int32(m.memory[uint32(v0)])
		switch t1 {
		default:
			store32(m.memory[int64(uint32(v2))+12:], uint32(v0+i32(4)))
			t2 := m.fn264(v1, i32(1100477), i32(2), v2+i32(12), i32(68))
			v0 = t2
			goto l15
		case 1:
			store32(m.memory[int64(uint32(v2))+12:], uint32(v0+i32(4)))
			t3 := m.fn264(v1, i32(1086625), i32(3), v2+i32(12), i32(69))
			v0 = t3
			goto l15
		case 2:
			store32(m.memory[int64(uint32(v2))+12:], uint32(v0+i32(4)))
			t4 := m.fn264(v1, i32(1086895), i32(3), v2+i32(12), i32(70))
			v0 = t4
			goto l15
		case 3:
			t5 := int32(load32(m.memory[uint32(v1):]))
			t6 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t7 := int32(load32(m.memory[int64(uint32(t6))+12:]))
			t8 := m.t0[uint(t7)].(func(int32, int32, int32) int32)(t5, i32(1086898), i32(8))
			v0 = t8
			goto l15
		case 4:
			store32(m.memory[int64(uint32(v2))+12:], uint32(v0+i32(1)))
			t9 := m.fn459(v1, i32(1086906), i32(12), i32(1086649), i32(3), v0+i32(4), i32(71), i32(1073156), i32(3), v2+i32(12), i32(72))
			v0 = t9
			goto l15
		case 5:
			t10 := int32(load32(m.memory[uint32(v1):]))
			t11 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t12 := int32(load32(m.memory[int64(uint32(t11))+12:]))
			t13 := m.t0[uint(t12)].(func(int32, int32, int32) int32)(t10, i32(1086918), i32(8))
			v0 = t13
			goto l15
		case 6:
			store32(m.memory[int64(uint32(v2))+12:], uint32(v0+i32(12)))
			t14 := m.fn462(v1, i32(1086926), i32(3), i32(1086672), i32(8), v0+i32(4), i32(73), i32(1086680), i32(5), v0+i32(8), i32(73), i32(1086649), i32(3), v2+i32(12), i32(74))
			v0 = t14
			goto l15
		case 7:
			t15 := int32(load32(m.memory[uint32(v1):]))
			t16 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t17 := int32(load32(m.memory[int64(uint32(t16))+12:]))
			t18 := m.t0[uint(t17)].(func(int32, int32, int32) int32)(t15, i32(1086929), i32(22))
			v0 = t18
			goto l15
		case 8:
			store32(m.memory[int64(uint32(v2))+12:], uint32(v0+i32(4)))
			t19 := m.fn264(v1, i32(1086951), i32(8), v2+i32(12), i32(74))
			v0 = t19
			goto l15
		case 9:
			store32(m.memory[int64(uint32(v2))+12:], uint32(v0+i32(4)))
			t20 := m.fn283(v1, i32(1086959), i32(14), i32(1086973), i32(10), v2+i32(12), i32(75))
			v0 = t20
			goto l15
		case 10:
			store32(m.memory[int64(uint32(v2))+12:], uint32(v0+i32(4)))
			t21 := m.fn264(v1, i32(1086983), i32(5), v2+i32(12), i32(75))
			v0 = t21
			goto l15
		case 11:
			store32(m.memory[int64(uint32(v2))+12:], uint32(v0+i32(1)))
			t22 := m.fn264(v1, i32(1086988), i32(4), v2+i32(12), i32(72))
			v0 = t22
			goto l15
		case 12:
			t23 := int32(load32(m.memory[uint32(v1):]))
			t24 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t25 := int32(load32(m.memory[int64(uint32(t24))+12:]))
			t26 := m.t0[uint(t25)].(func(int32, int32, int32) int32)(t23, i32(1086992), i32(5))
			v0 = t26
			goto l15
		case 13:
			store32(m.memory[int64(uint32(v2))+12:], uint32(v0+i32(4)))
			t27 := m.fn264(v1, i32(1086997), i32(17), v2+i32(12), i32(76))
			v0 = t27
			goto l15
		case 14:
			store32(m.memory[int64(uint32(v2))+12:], uint32(v0+i32(2)))
			t28 := m.fn283(v1, i32(1087014), i32(13), i32(1087027), i32(4), v2+i32(12), i32(77))
			v0 = t28
		}
	}
l15:
	m.g0 = v2 + i32(16)
	return v0
}
func (m *Module) fn417(v0 int32) {
	var v1 int32
	{
		t0 := int32(m.memory[uint32(v0)])
		v1 = t0
		switch v1 {
		default:
			if v1 == i32(13) {
				goto l4
			}
			return
		case 0:
			t1 := int32(m.memory[int64(uint32(v0))+4])
			t2 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			m.fn119(t1, t2)
			return
		case 1:
			m.fn367(v0 + i32(4))
			return
		case 2:
			m.fn451(v0 + i32(4))
			return
		}
	}
l4:
	t3 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t4 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	m.fn16(t3, t4)
}
func (m *Module) fn418(v0 int32) {
	var v1 int32
	t0 := m.g0
	v1 = t0 - i32(16)
	m.g0 = v1
	t1 := int32(load32(m.memory[uint32(v0):]))
	m.fn273(v1+i32(8), v0, t1, i32(1), i32(4), i32(20))
	{
		t2 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		v0 = t2
		if v0 == i32(-1) {
			goto l0
		}
		t3 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		m.fn2(v0, t3)
		panic("unreachable")
	}
l0:
	m.g0 = v1 + i32(16)
}
func (m *Module) fn419(v0, v1 int32) {
	m.fn136(v0, v1, i32(4), i32(16))
}
func (m *Module) fn420(v0 int32) {
	var v1, v2, v3 int32
	t0 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	v1 = t0
	t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	v2 = t1
	v3 = v2
l1:
	{
		if v1 == 0 {
			goto l0
		}
		t2 := int32(load32(m.memory[uint32(v3):]))
		t3 := int32(load32(m.memory[uint32(v3+i32(4)):]))
		m.fn16(t2, t3)
		v1 = v1 + i32(-1)
		v3 = v3 + i32(20)
		goto l1
	}
l0:
	t4 := int32(load32(m.memory[uint32(v0):]))
	m.fn426(t4, v2)
}
func (m *Module) fn421(v0 int32) {
	var v1, v2, v3 int32
	t0 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	v1 = t0
	t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	v2 = t1
	v3 = v2
l1:
	if v1 == 0 {
		goto l0
	}
	v1 = v1 + i32(-1)
	m.fn182(v3)
	v3 = v3 + i32(32)
	goto l1
l0:
	t2 := int32(load32(m.memory[uint32(v0):]))
	m.fn80(t2, v2)
}
func (m *Module) fn422(v0 int32) {
	var v1, v2 int32
	t0 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	v1 = t0
	v2 = v1 + i32(8)
	t1 := int32(load32(m.memory[int64(uint32(v0))+12:]))
	v1 = int32(uint32(t1-v1) >> 4)
l1:
	{
		if v1 == 0 {
			goto l0
		}
		t2 := int32(load32(m.memory[uint32(v2+i32(-4)):]))
		t3 := int32(load32(m.memory[uint32(v2):]))
		m.fn16(t2, t3)
		v1 = v1 + i32(-1)
		v2 = v2 + i32(16)
		goto l1
	}
l0:
	t4 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	t5 := int32(load32(m.memory[uint32(v0):]))
	m.fn419(t4, t5)
}
func (m *Module) fn423(v0 int32) {
	var v1, v2, v3, v4, v5, v6, v7 int32
	t0 := m.g0
	v1 = t0 - i32(16)
	m.g0 = v1
	v2 = i32(0)
	t1 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	t2 := int32(load32(m.memory[uint32(v0):]))
	v3 = t2
	p3 := i32(0)
	if v3 != 0 {
		p3 = t1
	}
	v4 = p3
	var p4 int32
	if v3 != i32(0) {
		p4 = 1
	}
	v5 = p4
	t5 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	v6 = t5
	v0 = v3
l15:
	if v4 != 0 {
		if v5&i32(1) == 0 {
			m.fn153(i32(1080172))
			panic("unreachable")
		}
		if v2 != 0 {
			goto l6
		}
	l8:
		{
			if v6 == 0 {
				v6 = i32(0)
				v0 = i32(0)
				v2 = v3
				goto l6
			}
			v6 = v6 + i32(-1)
			t7 := int32(load32(m.memory[int64(uint32(v3))+888:]))
			v3 = t7
			goto l8
		}
	l6:
		v4 = v4 + i32(-1)
	l11:
		{
			t8 := int32(load16(m.memory[int64(uint32(v2))+886:]))
			if uint32(v6) < uint32(t8) {
				if v0 != 0 {
					v3 = v2 + v6<<2 + i32(892)
				l14:
					{
						t12 := int32(load32(m.memory[uint32(v3):]))
						v7 = t12
						v3 = v7 + i32(888)
						v0 = v0 + i32(-1)
						if v0 != 0 {
							goto l14
						}
					}
					v0 = i32(0)
					goto l13
				}
				v0 = v6 + i32(1)
				v7 = v2
				goto l13
			}
			m.fn447(v1+i32(4), v2, v0)
			t9 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			v2 = t9
			if v2 == 0 {
				m.fn153(i32(1071172))
				panic("unreachable")
			}
			t10 := int32(load32(m.memory[int64(uint32(v1))+12:]))
			v6 = t10
			t11 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			v0 = t11
			goto l11
		}
	l13:
		v3 = v2 + v6*i32(12)
		t13 := int32(load32(m.memory[int64(uint32(v3))+4:]))
		t14 := int32(load32(m.memory[uint32(v3+i32(8)):]))
		m.fn16(t13, t14)
		m.fn427(v2 + v6*i32(68) + i32(136))
		v5 = i32(1)
		v3 = i32(0)
		v6 = v0
		v0 = i32(0)
		v2 = v7
		goto l15
	}
	if v5&i32(1) == 0 {
		goto l1
	}
	if v2 != 0 {
		goto l16
	}
l4:
	if v6 != 0 {
		v6 = v6 + i32(-1)
		t6 := int32(load32(m.memory[int64(uint32(v3))+888:]))
		v3 = t6
		goto l4
	}
	v2 = v3
	v3 = i32(0)
	goto l16
l16:
	{
		m.fn447(v1+i32(4), v2, v3)
		t15 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		v2 = t15
		if v2 == 0 {
			goto l1
		}
		t16 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		v3 = t16
		goto l16
	}
l1:
	m.g0 = v1 + i32(16)
}
func (m *Module) fn424(v0 int32) {
	var v1, v2, v3, v4, v5, v6 int32
	t0 := m.g0
	v1 = t0 - i32(16)
	m.g0 = v1
	v2 = i32(0)
	t1 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	t2 := int32(load32(m.memory[uint32(v0):]))
	v3 = t2
	p3 := i32(0)
	if v3 != 0 {
		p3 = t1
	}
	v4 = p3
	var p4 int32
	if v3 != i32(0) {
		p4 = 1
	}
	v5 = p4
	t5 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	v6 = t5
	v0 = v3
l13:
	if v4 != 0 {
		goto l0
	}
	if v5&i32(1) == 0 {
		goto l1
	}
	if v2 != 0 {
		goto l15
	}
	v2 = v3
l4:
	if v6 != 0 {
		v6 = v6 + i32(-1)
		t6 := int32(load32(m.memory[int64(uint32(v2))+44:]))
		v2 = t6
		goto l4
	}
	v3 = i32(0)
	goto l15
l0:
	if v5&i32(1) == 0 {
		m.fn153(i32(1080172))
		panic("unreachable")
	}
	if v2 != 0 {
		goto l6
	}
l8:
	{
		if v6 == 0 {
			v6 = i32(0)
			v0 = i32(0)
			v2 = v3
			goto l6
		}
		v6 = v6 + i32(-1)
		t7 := int32(load32(m.memory[int64(uint32(v3))+44:]))
		v3 = t7
		goto l8
	}
l6:
	v4 = v4 + i32(-1)
l11:
	{
		t8 := int32(load16(m.memory[int64(uint32(v2))+6:]))
		if uint32(v6) < uint32(t8) {
			if v0 != 0 {
				v6 = v2 + v6<<2 + i32(48)
			l14:
				{
					t12 := int32(load32(m.memory[uint32(v6):]))
					v2 = t12
					v6 = v2 + i32(44)
					v0 = v0 + i32(-1)
					if v0 != 0 {
						goto l14
					}
				}
				v5 = i32(1)
				v3 = i32(0)
				v6 = i32(0)
				v0 = i32(0)
				goto l13
			}
			v5 = i32(1)
			v6 = v6 + i32(1)
			v3 = i32(0)
			v0 = i32(0)
			goto l13
		}
		m.fn453(v1+i32(4), v2, v0)
		t9 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		v2 = t9
		if v2 == 0 {
			goto l10
		}
		t10 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		v6 = t10
		t11 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		v0 = t11
		goto l11
	}
l10:
	m.fn153(i32(1071172))
	panic("unreachable")
l15:
	{
		m.fn453(v1+i32(4), v2, v3)
		t13 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		v2 = t13
		if v2 == 0 {
			goto l1
		}
		t14 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		v3 = t14
		goto l15
	}
l1:
	m.g0 = v1 + i32(16)
}
func (m *Module) fn425(v0, v1, v2, v3, v4 int32) {
	var v5 int32
	t0 := m.g0
	v5 = t0 - i32(16)
	m.g0 = v5
	{
		t1 := int32(load32(m.memory[uint32(v1):]))
		if uint32(v2) <= uint32(t1) {
			goto l0
		}
		m.fn91(i32(1078043), i32(73), i32(1078080))
		panic("unreachable")
	}
l0:
	m.fn477(v5+i32(8), v1, v2, v3, v4)
	t2 := int32(load32(m.memory[int64(uint32(v5))+12:]))
	v1 = t2
	t3 := int32(load32(m.memory[int64(uint32(v5))+8:]))
	store32(m.memory[uint32(v0):], uint32(t3))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
	m.g0 = v5 + i32(16)
}
func (m *Module) fn426(v0, v1 int32) {
	m.fn136(v0, v1, i32(4), i32(20))
}
func (m *Module) fn427(v0 int32) {
	m.fn185(v0)
	m.fn78(v0 + i32(28))
	t0 := int32(load32(m.memory[int64(uint32(v0))+56:]))
	t1 := int32(load32(m.memory[int64(uint32(v0))+60:]))
	m.fn419(t0, t1)
}
func (m *Module) fn428() int32 {
	var v0, v1 int32
	t0 := m.g0
	v0 = t0 - i32(16)
	m.g0 = v0
	m.fn247(v0+i32(8), i32(4), i32(888))
	{
		t1 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		v1 = t1
		if v1 != 0 {
			m.g0 = v0 + i32(16)
			return v1
		}
		m.fn85(i32(4), i32(888))
		panic("unreachable")
	}
}
func (m *Module) fn429(v0, v1, v2, v3 int32) {
	if v1 != v3 {
		m.fn256(i32(1072679), i32(40), i32(1072720))
		panic("unreachable")
	}
	v1 = v1 * i32(68)
	if v1 == 0 {
		return
	}
	memory_copy(m.memory, uint32(v2), uint32(v0), uint32(v1))
}
func (m *Module) fn430(v0, v1, v2, v3 int32) {
	var v4, v5, v6 int32
	t0 := int32(load32(m.memory[uint32(v1):]))
	v4 = t0
	t1 := int32(load16(m.memory[int64(uint32(v4))+886:]))
	t2 := v4 + i32(4)
	v5 = t1 + i32(1)
	t3 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	t4 := v5
	v6 = t3
	m.fn250(t2, t4, v6, v2)
	m.fn476(v4+i32(136), v5, v6, v3)
	store16(m.memory[int64(uint32(v4))+886:], uint16(v5))
	store32(m.memory[int64(uint32(v0))+8:], uint32(v6))
	store32(m.memory[uint32(v0):], uint32(v4))
	t5 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	store32(m.memory[int64(uint32(v0))+4:], uint32(t5))
}
func (m *Module) fn431() int32 {
	var v0, v1 int32
	t0 := m.g0
	v0 = t0 - i32(16)
	m.g0 = v0
	m.fn247(v0+i32(8), i32(4), i32(936))
	{
		t1 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		v1 = t1
		if v1 != 0 {
			m.g0 = v0 + i32(16)
			return v1
		}
		m.fn85(i32(4), i32(936))
		panic("unreachable")
	}
}
func (m *Module) fn432(v0, v1, v2 int32) {
	var v3, v4, v5, v6 int32
	t0 := m.g0
	v3 = t0 - i32(32)
	m.g0 = v3
	t1 := int32(load16(m.memory[int64(uint32(v1))+886:]))
	v4 = t1
	m.memory[int64(uint32(v3))+28] = byte(i32(0))
	store32(m.memory[int64(uint32(v3))+24:], uint32(v4))
	store32(m.memory[int64(uint32(v3))+20:], uint32(i32(0)))
	v5 = v1 + i32(888)
l1:
	{
		m.fn254(v3+i32(8), v3+i32(20))
		t2 := int32(load32(m.memory[int64(uint32(v3))+8:]))
		if t2 != i32(1) {
			goto l0
		}
		t3 := int32(load32(m.memory[int64(uint32(v3))+12:]))
		t4 := v5
		v4 = t3
		t5 := int32(load32(m.memory[uint32(t4+v4<<2):]))
		v6 = t5
		store16(m.memory[int64(uint32(v6))+884:], uint16(v4))
		store32(m.memory[uint32(v6):], uint32(v1))
		goto l1
	}
l0:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v2))
	store32(m.memory[uint32(v0):], uint32(v1))
	m.g0 = v3 + i32(32)
}
func (m *Module) fn433(v0, v1, v2, v3, v4 int32) {
	var v5, v6, v7, v8 int32
	t0 := int32(load16(m.memory[int64(uint32(v0))+886:]))
	t1 := v0 + i32(4)
	v5 = t0
	v6 = v5 + i32(1)
	m.fn250(t1, v6, v1, v2)
	m.fn476(v0+i32(136), v6, v1, v3)
	v2 = v1 + i32(1)
	v7 = v0 + i32(888)
	{
		v3 = v5 + i32(2)
		t2 := v3
		v8 = v1 + i32(2)
		if uint32(t2) <= uint32(v8) {
			goto l0
		}
		v5 = (v5 - v1) << 2
		if v5 == 0 {
			goto l0
		}
		memory_copy(m.memory, uint32(v7+v8<<2), uint32(v7+v2<<2), uint32(v5))
	}
l0:
	store32(m.memory[uint32(v7+v2<<2):], uint32(v4))
	store16(m.memory[int64(uint32(v0))+886:], uint16(v6))
	p3 := v2
	if uint32(v3) > uint32(v2) {
		p3 = v3
	}
	v3 = p3
	v1 = v1<<2 + v0 + i32(892)
l2:
	{
		if v3 == v2 {
			return
		}
		t4 := int32(load32(m.memory[uint32(v1):]))
		v6 = t4
		store16(m.memory[int64(uint32(v6))+884:], uint16(v2))
		store32(m.memory[uint32(v6):], uint32(v0))
		v1 = v1 + i32(4)
		v2 = v2 + i32(1)
		goto l2
	}
}
func (m *Module) fn434(v0 int32) {
	var v1 int32
	t0 := m.g0
	v1 = t0 - i32(16)
	m.g0 = v1
	t1 := int32(load32(m.memory[uint32(v0):]))
	m.fn273(v1+i32(8), v0, t1, i32(1), i32(4), i32(32))
	{
		t2 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		v0 = t2
		if v0 == i32(-1) {
			goto l0
		}
		t3 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		m.fn2(v0, t3)
		panic("unreachable")
	}
l0:
	m.g0 = v1 + i32(16)
}
func (m *Module) fn435(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7, v8, v9, v10 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	store32(m.memory[int64(uint32(v2))+8:], uint32(v0))
	store32(m.memory[int64(uint32(v2))+12:], uint32(v0+v1))
	v1 = i32(0)
	v0 = i32(32)
	v3 = i32(0)
	v4 = i32(0)
	v5 = i32(0)
l19:
	v6 = i32(0)
l2:
	v7 = v4
	v8 = v1
	v9 = v0
	{
		t1 := m.fn48(v2 + i32(8))
		v0 = t1
		if v0 != i32(-1) {
			goto l0
		}
		v1 = i32(0)
		goto l1
	}
l0:
	v10 = v6 & i32(1)
	v1 = v8
	v4 = v7
	v6 = i32(0)
	if v10 != 0 {
		goto l2
	}
	v6 = i32(1)
	v1 = v8
	v4 = v7
	switch v0 + i32(-92) {
	case 0, 3:
		goto l2
	default:
		if v0 == i32(34) {
			v5 = v5 ^ i32(1)
			goto l10
		}
		v1 = v8
		v4 = v7
		if v0 == i32(42) {
			goto l2
		}
		fallthrough
	case 1, 2:
		v10 = v5 & i32(1)
		v1 = v8
		v4 = v7
		v5 = i32(1)
		v6 = i32(0)
		if v10 != 0 {
			goto l2
		}
		switch v0 + i32(-91) {
		case 0:
			v4 = v7 + i32(1)
			v5 = i32(0)
			goto l15
		case 1:
			goto l7
		case 2:
			v1 = v7 & i32(255)
			var p2 int32
			if v1 == i32(1) {
				p2 = 1
			}
			if p2&v3 == 0 {
				v5 = i32(0)
				t3 := v1
				var p4 int32
				if v1 != i32(0) {
					p4 = 1
				}
				v4 = t3 - p4
				goto l15
			}
			v1 = i32(2)
			goto l1
		default:
			goto l9
		}
	}
l9:
	v6 = i32(0)
	v1 = v8
	v4 = v7
	v5 = i32(1)
	if v0 == i32(34) {
		goto l2
	}
	v1 = i32(0)
	if v0 == i32(59) {
		goto l1
	}
l7:
	if v8&i32(1) != 0 {
		if v7&i32(255) != 0 {
			goto l12
		}
		v1 = i32(1)
		switch v0 + i32(-77) {
		case 0, 3:
			goto l1
		case 1, 2:
			goto l12
		default:
			switch v0 + i32(-109) {
			case 0, 3:
				goto l1
			case 1, 2:
				goto l12
			default:
				if v0 == i32(47) {
					goto l1
				}
				goto l12
			}
		}
	}
	if v7&i32(255) != 0 {
		goto l12
	}
	v1 = i32(1)
	v4 = i32(0)
	v5 = i32(0)
	v6 = i32(0)
	if v0 == i32(65) {
		goto l2
	}
	v5 = i32(0)
	v6 = i32(0)
	if v0 == i32(97) {
		goto l2
	}
	v4 = v0 + i32(-109)
	if uint32(v4) <= uint32(i32(12)) {
		if i32_shl(i32(1), v4)&i32(4161) != 0 {
			goto l1
		}
		goto l14
	}
	goto l14
l15:
	v1 = v8
	goto l19
l14:
	switch v0 + i32(-68) {
	case 0, 4:
		goto l1
	case 1, 2, 3:
		goto l12
	default:
		switch v0 + i32(-100) {
		case 0, 4:
			goto l1
		case 1, 2, 3:
			goto l12
		default:
			if v0 == i32(77) {
				goto l1
			}
			if v0 == i32(83) {
				goto l1
			}
			if v0 == i32(89) {
				goto l1
			}
		}
	}
l12:
	{
		if v3&i32(1) == 0 {
			goto l22
		}
		v5 = i32(0)
		v3 = i32(1)
		v1 = v8
		v4 = v7
		v6 = i32(0)
		p5 := v0
		if uint32(v0+i32(-65)) < uint32(i32(26)) {
			p5 = v0 | i32(32)
		}
		p6 := v9
		if uint32(v9+i32(-65)) < uint32(i32(26)) {
			p6 = v9 | i32(32)
		}
		if p5 == p6 {
			goto l2
		}
	}
l22:
	v5 = i32(0)
	v3 = i32(0)
	v1 = v8
	v4 = v7
	v6 = i32(0)
	if v9 != i32(91) {
		goto l2
	}
	v1 = v0 + i32(-72)
	if uint32(v1) > uint32(i32(11)) {
		goto l23
	}
	if i32_shl(i32(1), v1)&i32(2081) != 0 {
		goto l24
	}
l23:
	v3 = i32(0)
	v1 = v8
	v4 = v7
	v6 = i32(0)
	v10 = v0 + i32(-104)
	if uint32(v10) > uint32(i32(11)) {
		goto l2
	}
	v3 = i32(0)
	v1 = v8
	v4 = v7
	v6 = i32(0)
	if i32_shl(i32(1), v10)&i32(2081) == 0 {
		goto l2
	}
l24:
	v3 = i32(1)
l10:
	v1 = v8
	v4 = v7
	goto l19
l1:
	m.g0 = v2 + i32(16)
	return v1
}
func (m *Module) fn436(v0, v1 int32) {
	var v2, v3, v4, v5, v6, v7, v8, v9, v10, v11 int32
	t0 := m.fn442()
	v2 = t0
	store16(m.memory[int64(uint32(v2))+6:], uint16(i32(0)))
	store32(m.memory[uint32(v2):], uint32(i32(0)))
	t1 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	t2 := v2
	v3 = t1
	t3 := int32(load32(m.memory[uint32(v1):]))
	t4 := v3 ^ i32(-1)
	v4 = t3
	t5 := int32(load16(m.memory[int64(uint32(v4))+6:]))
	v5 = t5
	v6 = t4 + v5
	store16(m.memory[int64(uint32(t2))+6:], uint16(v6))
	if uint32(v6) < uint32(i32(12)) {
		goto l0
	}
	m.fn151(i32(0), v6, i32(11), i32(1079812))
	panic("unreachable")
l0:
	v7 = v4 + i32(8)
	t6 := int32(load16(m.memory[uint32(v7+v3<<1):]))
	v8 = t6
	v9 = v4 + i32(30)
	t7 := int32(m.memory[uint32(v9+v3)])
	v10 = t7
	t8 := v7
	v11 = v3 + i32(1)
	t9 := t8 + v11<<1
	v5 = v5 - v11
	m.fn479(t9, v5, v2+i32(8), v6)
	m.fn480(v9+v11, v5, v2+i32(30), v6)
	store16(m.memory[int64(uint32(v4))+6:], uint16(v3))
	m.memory[int64(uint32(v0))+18] = byte(v10)
	store16(m.memory[int64(uint32(v0))+16:], uint16(v8))
	store32(m.memory[uint32(v0):], uint32(v4))
	store32(m.memory[int64(uint32(v0))+12:], uint32(i32(0)))
	store32(m.memory[int64(uint32(v0))+8:], uint32(v2))
	t10 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	store32(m.memory[int64(uint32(v0))+4:], uint32(t10))
}
func (m *Module) fn437(v0, v1, v2, v3 int32) {
	var v4, v5, v6, v7, v8 int32
	t0 := int32(load32(m.memory[uint32(v1):]))
	v4 = t0
	t1 := int32(load16(m.memory[int64(uint32(v4))+6:]))
	t2 := v4 + i32(8)
	v5 = t1
	v6 = v5 + i32(1)
	t3 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	t4 := v6
	v7 = t3
	m.fn478(t2, t4, v7, v2)
	v2 = v4 + i32(30)
	v8 = v7 + i32(1)
	if uint32(v8) > uint32(v5) {
		goto l0
	}
	v5 = v5 - v7
	if v5 == 0 {
		goto l0
	}
	memory_copy(m.memory, uint32(v2+v8), uint32(v2+v7), uint32(v5))
l0:
	store16(m.memory[int64(uint32(v4))+6:], uint16(v6))
	store32(m.memory[int64(uint32(v0))+8:], uint32(v7))
	store32(m.memory[uint32(v0):], uint32(v4))
	m.memory[uint32(v2+v7)] = byte(v3)
	t5 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	store32(m.memory[int64(uint32(v0))+4:], uint32(t5))
}
func (m *Module) fn438() int32 {
	var v0, v1 int32
	t0 := m.g0
	v0 = t0 - i32(16)
	m.g0 = v0
	m.fn247(v0+i32(8), i32(4), i32(92))
	{
		t1 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		v1 = t1
		if v1 != 0 {
			m.g0 = v0 + i32(16)
			return v1
		}
		m.fn85(i32(4), i32(92))
		panic("unreachable")
	}
}
func (m *Module) fn439(v0, v1, v2 int32) {
	var v3, v4, v5, v6 int32
	t0 := m.g0
	v3 = t0 - i32(32)
	m.g0 = v3
	t1 := int32(load16(m.memory[int64(uint32(v1))+6:]))
	v4 = t1
	m.memory[int64(uint32(v3))+28] = byte(i32(0))
	store32(m.memory[int64(uint32(v3))+24:], uint32(v4))
	store32(m.memory[int64(uint32(v3))+20:], uint32(i32(0)))
	v5 = v1 + i32(44)
l1:
	{
		m.fn254(v3+i32(8), v3+i32(20))
		t2 := int32(load32(m.memory[int64(uint32(v3))+8:]))
		if t2 != i32(1) {
			goto l0
		}
		t3 := int32(load32(m.memory[int64(uint32(v3))+12:]))
		t4 := v5
		v4 = t3
		t5 := int32(load32(m.memory[uint32(t4+v4<<2):]))
		v6 = t5
		store16(m.memory[int64(uint32(v6))+4:], uint16(v4))
		store32(m.memory[uint32(v6):], uint32(v1))
		goto l1
	}
l0:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v2))
	store32(m.memory[uint32(v0):], uint32(v1))
	m.g0 = v3 + i32(32)
}
func (m *Module) fn440(v0, v1 int32) {
	var v2, v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13 int32
	var v14 int64
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	t1 := int32(load32(m.memory[uint32(v1):]))
	v3 = t1
	t2 := int32(load16(m.memory[int64(uint32(v3))+6:]))
	v4 = t2
	t3 := m.fn438()
	v5 = t3
	store16(m.memory[int64(uint32(v5))+6:], uint16(i32(0)))
	store32(m.memory[uint32(v5):], uint32(i32(0)))
	t4 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	t5 := v5
	v6 = t4
	t6 := int32(load16(m.memory[int64(uint32(v3))+6:]))
	t7 := v6 ^ i32(-1)
	v7 = t6
	v8 = t7 + v7
	store16(m.memory[int64(uint32(t5))+6:], uint16(v8))
	{
		if uint32(v8) >= uint32(i32(12)) {
			m.fn151(i32(0), v8, i32(11), i32(1079812))
			panic("unreachable")
		}
		v9 = v3 + i32(8)
		t8 := int32(load16(m.memory[uint32(v9+v6<<1):]))
		v10 = t8
		v11 = v3 + i32(30)
		t9 := int32(m.memory[uint32(v11+v6)])
		v12 = t9
		t10 := v9
		v13 = v6 + i32(1)
		t11 := t10 + v13<<1
		v7 = v7 - v13
		m.fn479(t11, v7, v5+i32(8), v8)
		m.fn480(v11+v13, v7, v5+i32(30), v8)
		store16(m.memory[int64(uint32(v3))+6:], uint16(v6))
		t12 := int32(load16(m.memory[int64(uint32(v5))+6:]))
		v13 = t12
		v8 = v13 + i32(1)
		{
			if uint32(v13) > uint32(i32(11)) {
				m.fn151(i32(0), v8, i32(12), i32(1070812))
				panic("unreachable")
			}
			if v4-v6 != v8 {
				m.fn256(i32(1072679), i32(40), i32(1072720))
				panic("unreachable")
			}
			v8 = v8 << 2
			if v8 == 0 {
				goto l3
			}
			memory_copy(m.memory, uint32(v5+i32(44)), uint32(v3+v6<<2+i32(48)), uint32(v8))
		l3:
			t13 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t14 := v2 + i32(8)
			t15 := v5
			v6 = t13
			m.fn439(t14, t15, v6)
			t16 := int64(load64(m.memory[int64(uint32(v2))+8:]))
			v14 = t16
			m.memory[int64(uint32(v0))+18] = byte(v12)
			store16(m.memory[int64(uint32(v0))+16:], uint16(v10))
			store32(m.memory[int64(uint32(v0))+4:], uint32(v6))
			store32(m.memory[uint32(v0):], uint32(v3))
			store64(m.memory[int64(uint32(v0))+8:], uint64(v14))
			m.g0 = v2 + i32(16)
			return
		}
	}
}
func (m *Module) fn441(v0, v1, v2, v3, v4 int32) {
	var v5, v6, v7, v8 int32
	t0 := int32(load16(m.memory[int64(uint32(v0))+6:]))
	t1 := v0 + i32(8)
	v5 = t0
	v6 = v5 + i32(1)
	m.fn478(t1, v6, v1, v2)
	v7 = v0 + i32(30)
	v2 = v1 + i32(1)
	if uint32(v2) > uint32(v5) {
		goto l0
	}
	v8 = v5 - v1
	if v8 == 0 {
		goto l0
	}
	memory_copy(m.memory, uint32(v7+v2), uint32(v7+v1), uint32(v8))
l0:
	m.memory[uint32(v7+v1)] = byte(v3)
	v3 = v0 + i32(44)
	{
		v7 = v5 + i32(2)
		t2 := v7
		v8 = v1 + i32(2)
		if uint32(t2) <= uint32(v8) {
			goto l1
		}
		v5 = (v5 - v1) << 2
		if v5 == 0 {
			goto l1
		}
		memory_copy(m.memory, uint32(v3+v8<<2), uint32(v3+v2<<2), uint32(v5))
	}
l1:
	store32(m.memory[uint32(v3+v2<<2):], uint32(v4))
	store16(m.memory[int64(uint32(v0))+6:], uint16(v6))
	p3 := v2
	if uint32(v7) > uint32(v2) {
		p3 = v7
	}
	v7 = p3
	v1 = v1<<2 + v0 + i32(48)
l3:
	{
		if v7 == v2 {
			return
		}
		t4 := int32(load32(m.memory[uint32(v1):]))
		v5 = t4
		store16(m.memory[int64(uint32(v5))+4:], uint16(v2))
		store32(m.memory[uint32(v5):], uint32(v0))
		v1 = v1 + i32(4)
		v2 = v2 + i32(1)
		goto l3
	}
}
func (m *Module) fn442() int32 {
	var v0, v1 int32
	t0 := m.g0
	v0 = t0 - i32(16)
	m.g0 = v0
	m.fn247(v0+i32(8), i32(4), i32(44))
	{
		t1 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		v1 = t1
		if v1 != 0 {
			m.g0 = v0 + i32(16)
			return v1
		}
		m.fn85(i32(4), i32(44))
		panic("unreachable")
	}
}
func (m *Module) fn443(v0 int32) {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	m.fn16(t0, t1)
	t2 := int32(load32(m.memory[int64(uint32(v0))+20:]))
	t3 := int32(load32(m.memory[int64(uint32(v0))+24:]))
	m.fn16(t2, t3)
}
func (m *Module) fn444(v0 int32) {
	m.fn423(v0 + i32(40))
	m.fn445(v0 + i32(16))
	m.fn446(v0 + i32(52))
	t0 := int32(load32(m.memory[int64(uint32(v0))+128:]))
	t1 := int32(load32(m.memory[int64(uint32(v0))+132:]))
	m.fn16(t0, t1)
}
func (m *Module) fn445(v0 int32) {
	m.fn230(v0)
	m.fn168(v0 + i32(12))
}
func (m *Module) fn446(v0 int32) {
	m.fn448(v0)
	t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
	t1 := int32(load32(m.memory[int64(uint32(v0))+16:]))
	m.fn16(t0, t1)
	t2 := int32(load32(m.memory[int64(uint32(v0))+32:]))
	t3 := int32(load32(m.memory[int64(uint32(v0))+36:]))
	m.fn449(t2, t3)
	t4 := int32(load32(m.memory[int64(uint32(v0))+44:]))
	t5 := int32(load32(m.memory[int64(uint32(v0))+48:]))
	m.fn16(t4, t5)
	t6 := int32(load32(m.memory[int64(uint32(v0))+64:]))
	t7 := int32(load32(m.memory[int64(uint32(v0))+68:]))
	m.fn449(t6, t7)
}
