package core

import (
	"encoding/binary"
	"math"
	"math/bits"
	"runtime"
	"unsafe"
)

func (m *Module) fn942(v0 int32) {
	m.fn943(v0)
	panic("unreachable")
}
func (m *Module) fn943(v0 int32) {
	var v1, v2, v3 int32
	t0 := m.g0
	v1 = t0 - i32(16)
	m.g0 = v1
	{
		t1 := int32(load32(m.memory[uint32(v0):]))
		v2 = t1
		t2 := int32(load32(m.memory[int64(uint32(v2))+4:]))
		v3 = t2
		if v3&i32(1) == 0 {
			store32(m.memory[uint32(v1):], uint32(i32(-1)))
			store32(m.memory[int64(uint32(v1))+12:], uint32(v0))
			t8 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			t9 := v1
			v0 = t8
			t10 := int32(m.memory[int64(uint32(v0))+8])
			t11 := int32(m.memory[int64(uint32(v0))+9])
			m.fn936(t9, i32(1273780), t10, t11)
			panic("unreachable")
		}
		t3 := int32(load32(m.memory[uint32(v2):]))
		v2 = t3
		store32(m.memory[int64(uint32(v1))+4:], uint32(int32(uint32(v3)>>1)))
		store32(m.memory[uint32(v1):], uint32(v2))
		t4 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		t5 := v1
		v0 = t4
		t6 := int32(m.memory[int64(uint32(v0))+8])
		t7 := int32(m.memory[int64(uint32(v0))+9])
		m.fn936(t5, i32(1273752), t6, t7)
		panic("unreachable")
	}
}
func (m *Module) fn944(v0 int32) {
	var v1, v2, v3 int32
	{
		t0 := int32(load32(m.memory[uint32(v0):]))
		v1 = t0
		if v1 < i32(1) {
			return
		}
		t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		v2 = t1
		t2 := int32(load32(m.memory[uint32(v2+i32(-4)):]))
		v0 = t2
		v3 = v0 & i32(-8)
		t3 := v3
		v0 = v0 & i32(3)
		p4 := i32(8)
		if v0 != 0 {
			p4 = i32(4)
		}
		if uint32(t3) < uint32(p4+v1) {
			m.fn3(i32(1273840), i32(46), i32(1273888))
			panic("unreachable")
		}
		if v0 == 0 {
			goto l2
		}
		if uint32(v3) > uint32(v1+i32(39)) {
			m.fn3(i32(1273904), i32(46), i32(1273952))
			panic("unreachable")
		}
	l2:
		m.fn1(v2)
	}
}
func (m *Module) fn945(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	if t0 == i32(-1) {
		t7 := int32(load32(m.memory[uint32(v1):]))
		t8 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t9 := int32(load32(m.memory[int64(uint32(v0))+12:]))
		t10 := int32(load32(m.memory[uint32(t9):]))
		v0 = t10
		t11 := int32(load32(m.memory[uint32(v0):]))
		t12 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		t13 := m.fn45(t7, t8, t11, t12)
		return t13
	}
	t1 := int32(load32(m.memory[uint32(v1):]))
	t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t3 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	t4 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	t5 := int32(load32(m.memory[int64(uint32(t4))+12:]))
	t6 := m.t0[uint(t5)].(func(int32, int32, int32) int32)(t1, t2, t3)
	return t6
}
func (m *Module) fn946(v0, v1 int32) {
	var v2, v3 int32
	var v4 int64
	t0 := m.g0
	v2 = t0 - i32(32)
	m.g0 = v2
	{
		t1 := int32(load32(m.memory[uint32(v1):]))
		if t1 != i32(-1) {
			goto l0
		}
		t2 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		v3 = t2
		store32(m.memory[int64(uint32(v2))+24:], uint32(i32(0)))
		store64(m.memory[int64(uint32(v2))+16:], uint64(i64(0x100000000)))
		t3 := int32(load32(m.memory[uint32(v3):]))
		t4 := v2 + i32(16)
		v3 = t3
		t5 := int32(load32(m.memory[uint32(v3):]))
		t6 := int32(load32(m.memory[int64(uint32(v3))+4:]))
		_ = m.fn45(t4, i32(1273728), t5, t6)
		t8 := int32(load32(m.memory[int64(uint32(v2))+24:]))
		t9 := v2
		v3 = t8
		store32(m.memory[int64(uint32(t9))+8:], uint32(v3))
		t10 := int64(load64(m.memory[int64(uint32(v2))+16:]))
		t11 := v2
		v4 = t10
		store64(m.memory[uint32(t11):], uint64(v4))
		store32(m.memory[int64(uint32(v1))+8:], uint32(v3))
		store64(m.memory[uint32(v1):], uint64(v4))
	}
l0:
	t12 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	v3 = t12
	store32(m.memory[int64(uint32(v1))+8:], uint32(i32(0)))
	t13 := int64(load64(m.memory[uint32(v1):]))
	v4 = t13
	store64(m.memory[uint32(v1):], uint64(i64(0x100000000)))
	store32(m.memory[int64(uint32(v2))+24:], uint32(v3))
	store64(m.memory[int64(uint32(v2))+16:], uint64(v4))
	{
		t14 := m.fn5(i32(12))
		v1 = t14
		if v1 != 0 {
			goto l1
		}
		m.fn24(i32(4), i32(12))
		panic("unreachable")
	}
l1:
	t15 := int32(load32(m.memory[int64(uint32(v2))+24:]))
	store32(m.memory[int64(uint32(v1))+8:], uint32(t15))
	t16 := int64(load64(m.memory[int64(uint32(v2))+16:]))
	store64(m.memory[uint32(v1):], uint64(t16))
	store32(m.memory[int64(uint32(v0))+4:], uint32(i32(1274672)))
	store32(m.memory[uint32(v0):], uint32(v1))
	m.g0 = v2 + i32(32)
}
func (m *Module) fn947(v0, v1 int32) {
	var v2, v3 int32
	var v4 int64
	t0 := m.g0
	v2 = t0 - i32(32)
	m.g0 = v2
	{
		t1 := int32(load32(m.memory[uint32(v1):]))
		if t1 != i32(-1) {
			goto l0
		}
		t2 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		v3 = t2
		store32(m.memory[int64(uint32(v2))+28:], uint32(i32(0)))
		store64(m.memory[int64(uint32(v2))+20:], uint64(i64(0x100000000)))
		t3 := int32(load32(m.memory[uint32(v3):]))
		t4 := v2 + i32(20)
		v3 = t3
		t5 := int32(load32(m.memory[uint32(v3):]))
		t6 := int32(load32(m.memory[int64(uint32(v3))+4:]))
		_ = m.fn45(t4, i32(1273728), t5, t6)
		t8 := int32(load32(m.memory[int64(uint32(v2))+28:]))
		t9 := v2
		v3 = t8
		store32(m.memory[int64(uint32(t9))+16:], uint32(v3))
		t10 := int64(load64(m.memory[int64(uint32(v2))+20:]))
		t11 := v2
		v4 = t10
		store64(m.memory[int64(uint32(t11))+8:], uint64(v4))
		store32(m.memory[int64(uint32(v1))+8:], uint32(v3))
		store64(m.memory[uint32(v1):], uint64(v4))
	}
l0:
	store32(m.memory[int64(uint32(v0))+4:], uint32(i32(1274672)))
	store32(m.memory[uint32(v0):], uint32(v1))
	m.g0 = v2 + i32(32)
}
func (m *Module) fn948(v0, v1, v2 int32) {
	var v3 int32
	t0 := m.g0
	v3 = t0 - i32(16)
	m.g0 = v3
	v1 = v2 + v1
	if uint32(v1) >= uint32(v2) {
		goto l0
	}
	m.fn10(i32(0), i32(0))
	panic("unreachable")
l0:
	t1 := int32(load32(m.memory[uint32(v0):]))
	t2 := v3 + i32(4)
	v2 = t1
	t3 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t4 := v2
	t5 := v1
	v2 = v2 << 1
	p6 := v2
	if uint32(v1) > uint32(v2) {
		p6 = t5
	}
	v2 = p6
	p7 := i32(8)
	if uint32(v2) > uint32(i32(8)) {
		p7 = v2
	}
	v2 = p7
	m.fn949(t2, t4, t3, v2)
	{
		t8 := int32(load32(m.memory[int64(uint32(v3))+4:]))
		if t8 != i32(1) {
			goto l1
		}
		t9 := int32(load32(m.memory[int64(uint32(v3))+8:]))
		t10 := int32(load32(m.memory[int64(uint32(v3))+12:]))
		m.fn10(t9, t10)
		panic("unreachable")
	}
l1:
	t11 := int32(load32(m.memory[int64(uint32(v3))+8:]))
	v1 = t11
	store32(m.memory[uint32(v0):], uint32(v2))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
	m.g0 = v3 + i32(16)
}
func (m *Module) fn949(v0, v1, v2, v3 int32) {
	var v4 int32
	v4 = i32(0)
	if v3 >= i32(0) {
		goto l0
	}
	v1 = i32(1)
	v2 = i32(4)
	goto l1
l0:
	{
		{
			if v1 == 0 {
				goto l2
			}
			t0 := m.fn22(v2, v1, i32(1), v3)
			v4 = t0
			goto l3
		}
	l2:
		t1 := m.fn5(v3)
		v4 = t1
	}
l3:
	if v4 != 0 {
		goto l4
	}
	v1 = i32(1)
	store32(m.memory[int64(uint32(v0))+4:], uint32(i32(1)))
	goto l5
l4:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v4))
	v1 = i32(0)
l5:
	v2 = i32(8)
	v4 = v3
l1:
	store32(m.memory[uint32(v0+v2):], uint32(v4))
	store32(m.memory[uint32(v0):], uint32(v1))
}
func (m *Module) fn950(v0, v1 int32) {
	t0 := int64(load64(m.memory[int64(uint32(i32(0)))+1273832:]))
	store64(m.memory[int64(uint32(v0))+8:], uint64(t0))
	t1 := int64(load64(m.memory[int64(uint32(i32(0)))+1273824:]))
	store64(m.memory[uint32(v0):], uint64(t1))
}
func (m *Module) fn951(v0, v1, v2 int32) int32 {
	var v3 int32
	{
		{
			t0 := int32(load32(m.memory[uint32(v0):]))
			t1 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			t2 := v2
			v3 = t1
			if uint32(t2) <= uint32(t0-v3) {
				goto l0
			}
			m.fn948(v0, v3, v2)
			t3 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			v3 = t3
			goto l1
		}
	l0:
		if v2 == 0 {
			goto l2
		}
	l1:
		if v2 == 0 {
			goto l2
		}
		t4 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		memory_copy(m.memory, uint32(t4+v3), uint32(v1), uint32(v2))
	}
l2:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v3+v2))
	return i32(0)
}
func (m *Module) fn952(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7 int32
	t0 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	v2 = t0
	{
		if uint32(v1) >= uint32(i32(128)) {
			goto l0
		}
		v3 = i32(1)
		goto l1
	l0:
		if uint32(v1) >= uint32(i32(2048)) {
			goto l2
		}
		v3 = i32(2)
		goto l1
	l2:
		p1 := i32(4)
		if uint32(v1) < uint32(i32(65536)) {
			p1 = i32(3)
		}
		v3 = p1
	}
l1:
	{
		t2 := int32(load32(m.memory[uint32(v0):]))
		if uint32(v3) <= uint32(t2-v2) {
			goto l3
		}
		m.fn948(v0, v2, v3)
	}
l3:
	t3 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	v4 = t3 + v2
	if uint32(v1) < uint32(i32(128)) {
		goto l4
	}
	v5 = v1&i32(63) | i32(-128)
	v6 = int32(uint32(v1) >> 6)
	if uint32(v1) >= uint32(i32(2048)) {
		v7 = int32(uint32(v1) >> 12)
		v6 = v6&i32(63) | i32(-128)
		if uint32(v1) > uint32(i32(0xffff)) {
			m.memory[int64(uint32(v4))+3] = byte(v5)
			m.memory[int64(uint32(v4))+2] = byte(v6)
			m.memory[int64(uint32(v4))+1] = byte(v7&i32(63) | i32(-128))
			m.memory[uint32(v4)] = byte(int32(uint32(v1)>>18) | i32(-16))
			goto l6
		}
		m.memory[int64(uint32(v4))+2] = byte(v5)
		m.memory[int64(uint32(v4))+1] = byte(v6)
		m.memory[uint32(v4)] = byte(v7 | i32(224))
		goto l6
	}
	m.memory[int64(uint32(v4))+1] = byte(v5)
	m.memory[uint32(v4)] = byte(v6 | i32(192))
	goto l6
l4:
	m.memory[uint32(v4)] = byte(v1)
l6:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v3+v2))
	return i32(0)
}
func (m *Module) fn953(v0, v1, v2 int32) int32 {
	t0 := m.fn45(v0, i32(1273728), v1, v2)
	return t0
}
func (m *Module) fn954(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v1):]))
	t1 := int32(load32(m.memory[uint32(v0):]))
	t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t3 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	t4 := int32(load32(m.memory[int64(uint32(t3))+12:]))
	t5 := m.t0[uint(t4)].(func(int32, int32, int32) int32)(t0, t1, t2)
	return t5
}
func (m *Module) fn955(v0, v1 int32) {
	var v2, v3 int32
	t0 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	v2 = t0
	t1 := int32(load32(m.memory[uint32(v1):]))
	v3 = t1
	{
		t2 := m.fn5(i32(8))
		v1 = t2
		if v1 != 0 {
			goto l0
		}
		m.fn24(i32(4), i32(8))
		panic("unreachable")
	}
l0:
	store32(m.memory[int64(uint32(v1))+4:], uint32(v2))
	store32(m.memory[uint32(v1):], uint32(v3))
	store32(m.memory[int64(uint32(v0))+4:], uint32(i32(1274084)))
	store32(m.memory[uint32(v0):], uint32(v1))
}
func (m *Module) fn956(v0, v1 int32) {
	store32(m.memory[int64(uint32(v0))+4:], uint32(i32(1274084)))
	store32(m.memory[uint32(v0):], uint32(v1))
}
func (m *Module) fn957(v0, v1 int32) {
	t0 := int64(load64(m.memory[uint32(v1):]))
	store64(m.memory[uint32(v0):], uint64(t0))
}
func (m *Module) fn958(v0, v1 int32) {
	var v2, v3, v4, v5 int32
	v2 = i32(31)
	{
		if uint32(v1) >= uint32(i32(0x1000000)) {
			goto l0
		}
		t0 := v1
		v2 = int32(bits.LeadingZeros32(uint32(int32(uint32(v1) >> 8))))
		v2 = i32_shr_u(t0, i32(38)-v2)&i32(1) | v2<<1 ^ i32(62)
	}
l0:
	store64(m.memory[int64(uint32(v0))+16:], uint64(i64(0)))
	store32(m.memory[int64(uint32(v0))+28:], uint32(v2))
	v3 = v2<<2 + i32(1293404)
	{
		t1 := int32(load32(m.memory[int64(uint32(i32(0)))+1293816:]))
		v4 = i32_shl(i32(1), v2)
		if t1&v4 != 0 {
			goto l1
		}
		store32(m.memory[uint32(v3):], uint32(v0))
		store32(m.memory[int64(uint32(v0))+24:], uint32(v3))
		store32(m.memory[int64(uint32(v0))+12:], uint32(v0))
		store32(m.memory[int64(uint32(v0))+8:], uint32(v0))
		t2 := int32(load32(m.memory[int64(uint32(i32(0)))+1293816:]))
		store32(m.memory[int64(uint32(i32(0)))+1293816:], uint32(t2|v4))
		return
	}
l1:
	{
		{
			{
				t3 := int32(load32(m.memory[uint32(v3):]))
				v4 = t3
				t4 := int32(load32(m.memory[int64(uint32(v4))+4:]))
				if t4&i32(-8) != v1 {
					goto l2
				}
				v2 = v4
				goto l3
			}
		l2:
			t6 := v1
			p5 := i32(25) - int32(uint32(v2)>>1)
			if v2 == i32(31) {
				p5 = i32(0)
			}
			v3 = i32_shl(t6, p5)
		l5:
			{
				v5 = v4 + int32(uint32(v3)>>29)&i32(4)
				t7 := int32(load32(m.memory[int64(uint32(v5))+16:]))
				v2 = t7
				if v2 == 0 {
					goto l4
				}
				v3 = v3 << 1
				v4 = v2
				t8 := int32(load32(m.memory[int64(uint32(v2))+4:]))
				if t8&i32(-8) != v1 {
					goto l5
				}
			}
		}
	l3:
		t9 := int32(load32(m.memory[int64(uint32(v2))+8:]))
		v3 = t9
		store32(m.memory[int64(uint32(v3))+12:], uint32(v0))
		store32(m.memory[int64(uint32(v2))+8:], uint32(v0))
		store32(m.memory[int64(uint32(v0))+24:], uint32(i32(0)))
		store32(m.memory[int64(uint32(v0))+12:], uint32(v2))
		store32(m.memory[int64(uint32(v0))+8:], uint32(v3))
		return
	}
l4:
	store32(m.memory[uint32(v5+i32(16)):], uint32(v0))
	store32(m.memory[int64(uint32(v0))+24:], uint32(v4))
	store32(m.memory[int64(uint32(v0))+12:], uint32(v0))
	store32(m.memory[int64(uint32(v0))+8:], uint32(v0))
}
func (m *Module) fn959(v0, v1 int32) {
	m.fn941(v0, v1)
	panic("unreachable")
}
func (m *Module) fn960(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t1 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	t2 := m.fn56(v1, t0, t1)
	return t2
}
func (m *Module) fn961(v0, v1 int32) int32 {
	var v2, v3 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	{
		{
			t1 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			v3 = t1
			if v3&i32(0x2000000) != 0 {
				t3 := int32(load32(m.memory[uint32(v0):]))
				v3 = t3
				v0 = i32(9)
			l3:
				{
					t4 := int32(m.memory[int64(uint32(v3&i32(15)))+1098816])
					m.memory[uint32(v2+i32(8)+v0+i32(-2))] = byte(t4)
					v0 = v0 + i32(-1)
					v3 = int32(uint32(v3) >> 4)
					if v3 != 0 {
						goto l3
					}
				}
				t5 := m.fn306(v1, i32(1), i32(1122550), i32(2), v2+i32(8)+v0+i32(-1), i32(9)-v0)
				v0 = t5
				goto l2
			}
			if v3&i32(0x4000000) != 0 {
				goto l1
			}
			t2 := m.fn513(v0, v1)
			v0 = t2
			goto l2
		}
	l1:
		t6 := int32(load32(m.memory[uint32(v0):]))
		v3 = t6
		v0 = i32(9)
	l4:
		{
			t7 := int32(m.memory[int64(uint32(v3&i32(15)))+1122552])
			m.memory[uint32(v2+i32(8)+v0+i32(-2))] = byte(t7)
			v0 = v0 + i32(-1)
			v3 = int32(uint32(v3) >> 4)
			if v3 != 0 {
				goto l4
			}
		}
		t8 := m.fn306(v1, i32(1), i32(1122550), i32(2), v2+i32(8)+v0+i32(-1), i32(9)-v0)
		v0 = t8
	}
l2:
	m.g0 = v2 + i32(16)
	return v0
}
func (m *Module) fn962(v0, v1, v2 int32) int32 {
	t0 := v1
	v0 = v0 & i32(255) << 2
	t1 := int32(load32(m.memory[int64(uint32(v0))+1291968:]))
	t2 := int32(load32(m.memory[int64(uint32(v0))+1291800:]))
	t3 := int32(load32(m.memory[int64(uint32(v2))+12:]))
	t4 := m.t0[uint(t3)].(func(int32, int32, int32) int32)(t0, t1, t2)
	return t4
}
func (m *Module) fn963(v0, v1 int32) {
	var v2 int32
	{
		t0 := m.fn5(i32(12))
		v2 = t0
		if v2 == 0 {
			m.fn24(i32(4), i32(12))
			panic("unreachable")
		}
		t1 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		store32(m.memory[int64(uint32(v2))+8:], uint32(t1))
		t2 := int64(load64(m.memory[uint32(v1):]))
		store64(m.memory[uint32(v2):], uint64(t2))
		t3 := m.fn5(i32(12))
		v1 = t3
		if v1 == 0 {
			m.fn24(i32(4), i32(12))
			panic("unreachable")
		}
		m.memory[int64(uint32(v1))+8] = byte(i32(21))
		store32(m.memory[int64(uint32(v1))+4:], uint32(i32(1275336)))
		store32(m.memory[uint32(v1):], uint32(v2))
		store64(m.memory[uint32(v0):], uint64(int64(uint32(v1))<<32|i64(3)))
		return
	}
}
func (m *Module) fn964(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	v3 = i32(1)
	{
		t1 := int32(load32(m.memory[uint32(v1):]))
		v4 = t1
		t2 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t3 := v4
		v5 = t2
		t4 := int32(load32(m.memory[int64(uint32(v5))+12:]))
		v6 = t4
		t5 := m.t0[uint(v6)].(func(int32, int32, int32) int32)(t3, i32(1276588), i32(15))
		if t5 != 0 {
			goto l0
		}
		{
			{
				t6 := int32(m.memory[int64(uint32(v1))+10])
				if t6&i32(128) != 0 {
					goto l1
				}
				v3 = i32(1)
				t7 := m.t0[uint(v6)].(func(int32, int32, int32) int32)(v4, i32(1099043), i32(1))
				if t7 != 0 {
					goto l0
				}
				t8 := int32(m.memory[uint32(v0)])
				t9 := v4
				v1 = t8 << 2
				t10 := int32(load32(m.memory[int64(uint32(v1))+1292160:]))
				t11 := int32(load32(m.memory[int64(uint32(v1))+1292136:]))
				t12 := m.t0[uint(v6)].(func(int32, int32, int32) int32)(t9, t10, t11)
				if t12 == 0 {
					goto l2
				}
				goto l0
			}
		l1:
			t13 := m.t0[uint(v6)].(func(int32, int32, int32) int32)(v4, i32(1099044), i32(2))
			if t13 != 0 {
				goto l0
			}
			store32(m.memory[int64(uint32(v2))+4:], uint32(v5))
			store32(m.memory[uint32(v2):], uint32(v4))
			v3 = i32(1)
			m.memory[int64(uint32(v2))+15] = byte(i32(1))
			t14 := int32(m.memory[uint32(v0)])
			v1 = t14 << 2
			t15 := int32(load32(m.memory[int64(uint32(v1))+1292184:]))
			v0 = t15
			t16 := int32(load32(m.memory[int64(uint32(v1))+1292208:]))
			v1 = t16
			store32(m.memory[int64(uint32(v2))+8:], uint32(v2+i32(15)))
			t17 := m.fn342(v2, v1, v0)
			if t17 != 0 {
				goto l0
			}
			t18 := m.fn342(v2, i32(1099041), i32(2))
			if t18 != 0 {
				goto l0
			}
		}
	l2:
		t19 := m.t0[uint(v6)].(func(int32, int32, int32) int32)(v4, i32(1272328), i32(1))
		v3 = t19
	}
l0:
	m.g0 = v2 + i32(16)
	return v3
}
func (m *Module) fn965(v0, v1, v2 int32) {
	var v3, v4, v5 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		v3 = t0
		if uint32(v3) > uint32(v2) {
			store32(m.memory[int64(uint32(v1))+12:], uint32(v3+(v2^i32(-1))))
			t1 := int32(load32(m.memory[uint32(v1):]))
			t2 := v1
			v4 = t1
			v3 = v4 + v2
			t3 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t4 := v3 + i32(1)
			v5 = t3
			t5 := v5
			t6 := v2
			v4 = v5 - v4
			p7 := t5
			if uint32(t6) < uint32(v4) {
				p7 = t4
			}
			store32(m.memory[uint32(t2):], uint32(p7))
			if uint32(v2) >= uint32(v4) {
				goto l2
			}
			t8 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			t9 := v1
			v2 = t8 + v2
			store32(m.memory[int64(uint32(t9))+8:], uint32(v2+i32(1)))
			goto l3
		}
		if v3 != 0 {
			t10 := int32(load32(m.memory[uint32(v1):]))
			t11 := v1
			v2 = t10
			t12 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t13 := v2 + v3
			v4 = t12
			t14 := v4
			v5 = v3 + i32(-1)
			t15 := v5
			v2 = v4 - v2
			p16 := t14
			if uint32(t15) < uint32(v2) {
				p16 = t13
			}
			store32(m.memory[uint32(t11):], uint32(p16))
			{
				if uint32(v5) >= uint32(v2) {
					goto l4
				}
				t17 := int32(load32(m.memory[int64(uint32(v1))+8:]))
				store32(m.memory[int64(uint32(v1))+8:], uint32(t17+v3))
			}
		l4:
			v3 = i32(0)
			store32(m.memory[int64(uint32(v1))+12:], uint32(i32(0)))
			goto l3
		}
		goto l2
	}
l2:
	v3 = i32(0)
l3:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
	store32(m.memory[uint32(v0):], uint32(v2))
}
func (m *Module) fn966(v0, v1 int32) {
	var v2 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	store32(m.memory[int64(uint32(v2))+12:], uint32(v1))
	store32(m.memory[int64(uint32(v2))+8:], uint32(v0))
	m.fn842(i32(1), v2+i32(8), i32(1279468), v2+i32(12), i32(1279468), i32(0), v2, i32(1275592))
	panic("unreachable")
}
func (m *Module) fn967(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v1):]))
	t1 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	t2 := int32(load32(m.memory[int64(uint32(t1))+12:]))
	t3 := m.t0[uint(t2)].(func(int32, int32, int32) int32)(t0, i32(1290048), i32(11))
	return t3
}
func (m *Module) fn968(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v1):]))
	t1 := int32(load32(m.memory[uint32(v0):]))
	t2 := int32(load32(m.memory[uint32(t1):]))
	v0 = t2 << 2
	t3 := int32(load32(m.memory[uint32(v0+i32(1293316)):]))
	t4 := int32(load32(m.memory[uint32(v0+i32(1293280)):]))
	t5 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	t6 := int32(load32(m.memory[int64(uint32(t5))+12:]))
	t7 := m.t0[uint(t6)].(func(int32, int32, int32) int32)(t0, t3, t4)
	return t7
}
func (m *Module) fn969(v0, v1 int32) int32 {
	var v2, v3 int32
	var v4 int64
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	t1 := int32(load32(m.memory[uint32(v0):]))
	v3 = t1
	{
		t2 := int64(load64(m.memory[int64(uint32(v1))+8:]))
		v4 = t2
		v0 = int32(v4)
		if v0&i32(0x800000) == 0 {
			goto l0
		}
		if v0&i32(0x8000000) == 0 {
			goto l1
		}
		v0 = v0 | i32(0x1000000)
		goto l0
	l1:
		store16(m.memory[int64(uint32(v1))+12:], uint16(i32(10)))
		v0 = v0 | i32(0x9000000)
	}
l0:
	store32(m.memory[int64(uint32(v1))+8:], uint32(v0|i32(0x800000)))
	v0 = i32(9)
l2:
	{
		t3 := int32(m.memory[int64(uint32(v3&i32(15)))+1098816])
		m.memory[uint32(v2+i32(8)+v0+i32(-2))] = byte(t3)
		v0 = v0 + i32(-1)
		v3 = int32(uint32(v3) >> 4)
		if v3 != 0 {
			goto l2
		}
	}
	t4 := m.fn306(v1, i32(1), i32(1122550), i32(2), v2+i32(8)+v0+i32(-1), i32(9)-v0)
	v0 = t4
	store64(m.memory[int64(uint32(v1))+8:], uint64(v4))
	m.g0 = v2 + i32(16)
	return v0
}
func (m *Module) fn970(v0, v1, v2, v3 int32) {
	var v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14, v15, v16, v17, v18, v19, v20, v21, v22, v23, v24, v25, v26, v27, v28, v29, v30, v31, v32 int32
	t0 := m.g0
	v4 = t0 - i32(128)
	m.g0 = v4
	{
		if uint32(v1) < uint32(i32(33)) {
			goto l0
		}
	l19:
		{
			if v3 != 0 {
				goto l1
			}
			m.fn971(v0, v1)
			goto l2
		l1:
			t1 := v0
			v5 = int32(uint32(v1) >> 3)
			v6 = t1 + v5*i32(28)
			v7 = v0 + v5<<4
			{
				{
					if uint32(v1) < uint32(i32(64)) {
						goto l3
					}
					t2 := m.fn972(v0, v7, v6, v5)
					v5 = t2
					goto l4
				}
			l3:
				t3 := int32(load32(m.memory[uint32(v0):]))
				t4 := v0
				t5 := v6
				t6 := v7
				v5 = t3
				t7 := int32(load32(m.memory[uint32(v7):]))
				t8 := v5
				v8 = t7
				var p9 int32
				if uint32(t8) < uint32(v8) {
					p9 = 1
				}
				v9 = p9
				t10 := int32(load32(m.memory[uint32(v6):]))
				t11 := v9
				t12 := v8
				v10 = t10
				var p13 int32
				if uint32(t12) < uint32(v10) {
					p13 = 1
				}
				p14 := t6
				if t11^p13 != 0 {
					p14 = t5
				}
				t15 := v9
				var p16 int32
				if uint32(v5) < uint32(v10) {
					p16 = 1
				}
				p17 := p14
				if t15^p16 != 0 {
					p17 = t4
				}
				v5 = p17
			}
		l4:
			v3 = v3 + i32(-1)
			v5 = v5 - v0
			{
				{
					if v2 != 0 {
						t20 := int32(load32(m.memory[uint32(v0):]))
						v7 = t20
						t21 := int32(load32(m.memory[uint32(v2):]))
						v8 = v0 + v5
						t22 := int32(load32(m.memory[uint32(v8):]))
						v6 = t22
						if uint32(t21) < uint32(v6) {
							goto l6
						}
						store32(m.memory[uint32(v0):], uint32(v6))
						store32(m.memory[uint32(v8):], uint32(v7))
						v8 = v0 + i32(4)
						t23 := int32(load32(m.memory[uint32(v0):]))
						v9 = t23
						t24 := int32(load32(m.memory[int64(uint32(v0))+4:]))
						v2 = t24
						v6 = i32(0)
						{
							v5 = v0 + i32(8)
							t25 := v5
							v11 = v0 + v1<<2
							v12 = v11 + i32(-4)
							if uint32(t25) < uint32(v12) {
								goto l7
							}
							v7 = v8
							goto l8
						}
					l7:
						v6 = i32(0)
					l9:
						{
							t26 := v5 + i32(-4)
							v7 = v8 + v6<<2
							t27 := int32(load32(m.memory[uint32(v7):]))
							store32(m.memory[uint32(t26):], uint32(t27))
							t28 := int32(load32(m.memory[uint32(v5):]))
							t29 := v7
							v10 = t28
							store32(m.memory[uint32(t29):], uint32(v10))
							t30 := v5
							t31 := v8
							t32 := v6
							var p33 int32
							if uint32(v9) >= uint32(v10) {
								p33 = 1
							}
							v6 = t32 + p33
							v7 = t31 + v6<<2
							t34 := int32(load32(m.memory[uint32(v7):]))
							store32(m.memory[uint32(t30):], uint32(t34))
							t35 := int32(load32(m.memory[uint32(v5+i32(4)):]))
							t36 := v7
							v10 = t35
							store32(m.memory[uint32(t36):], uint32(v10))
							t37 := v6
							var p38 int32
							if uint32(v9) >= uint32(v10) {
								p38 = 1
							}
							v6 = t37 + p38
							v5 = v5 + i32(8)
							if uint32(v5) < uint32(v12) {
								goto l9
							}
						}
						v7 = v5 + i32(-4)
					l8:
						if v5 == v11 {
							goto l10
						}
					l11:
						{
							t39 := v7
							v10 = v8 + v6<<2
							t40 := int32(load32(m.memory[uint32(v10):]))
							store32(m.memory[uint32(t39):], uint32(t40))
							t41 := v10
							v7 = v5
							t42 := int32(load32(m.memory[uint32(v7):]))
							v5 = t42
							store32(m.memory[uint32(t41):], uint32(v5))
							t43 := v6
							var p44 int32
							if uint32(v9) >= uint32(v5) {
								p44 = 1
							}
							v6 = t43 + p44
							v5 = v7 + i32(4)
							if v5 != v11 {
								goto l11
							}
						}
						v7 = v5 + i32(-4)
					l10:
						t45 := v7
						v5 = v8 + v6<<2
						t46 := int32(load32(m.memory[uint32(v5):]))
						store32(m.memory[uint32(t45):], uint32(t46))
						store32(m.memory[uint32(v5):], uint32(v2))
						t47 := v6
						var p48 int32
						if uint32(v9) >= uint32(v2) {
							p48 = 1
						}
						v5 = t47 + p48
						if uint32(v5) >= uint32(v1) {
							goto l12
						}
						t49 := int32(load32(m.memory[uint32(v0):]))
						v6 = t49
						t50 := v0
						v7 = v0 + v5<<2
						t51 := int32(load32(m.memory[uint32(v7):]))
						store32(m.memory[uint32(t50):], uint32(t51))
						store32(m.memory[uint32(v7):], uint32(v6))
						t52 := v1
						v5 = v5 + i32(1)
						v1 = t52 - v5
						v0 = v0 + v5<<2
						v2 = i32(0)
						goto l13
					}
					t18 := int32(load32(m.memory[uint32(v0+v5):]))
					v6 = t18
					t19 := int32(load32(m.memory[uint32(v0):]))
					v7 = t19
					goto l6
				}
			l6:
				store32(m.memory[uint32(v0):], uint32(v6))
				store32(m.memory[uint32(v0+v5):], uint32(v7))
				v8 = v0 + i32(4)
				t53 := int32(load32(m.memory[uint32(v0):]))
				v9 = t53
				t54 := int32(load32(m.memory[int64(uint32(v0))+4:]))
				v13 = t54
				v6 = i32(0)
				{
					v5 = v0 + i32(8)
					t55 := v5
					v11 = v0 + v1<<2
					v12 = v11 + i32(-4)
					if uint32(t55) < uint32(v12) {
						goto l14
					}
					v7 = v8
					goto l15
				}
			l14:
				v6 = i32(0)
			l16:
				{
					t56 := v5 + i32(-4)
					v7 = v8 + v6<<2
					t57 := int32(load32(m.memory[uint32(v7):]))
					store32(m.memory[uint32(t56):], uint32(t57))
					t58 := int32(load32(m.memory[uint32(v5):]))
					t59 := v7
					v10 = t58
					store32(m.memory[uint32(t59):], uint32(v10))
					t60 := v5
					t61 := v8
					t62 := v6
					var p63 int32
					if uint32(v10) < uint32(v9) {
						p63 = 1
					}
					v6 = t62 + p63
					v7 = t61 + v6<<2
					t64 := int32(load32(m.memory[uint32(v7):]))
					store32(m.memory[uint32(t60):], uint32(t64))
					t65 := int32(load32(m.memory[uint32(v5+i32(4)):]))
					t66 := v7
					v10 = t65
					store32(m.memory[uint32(t66):], uint32(v10))
					t67 := v6
					var p68 int32
					if uint32(v10) < uint32(v9) {
						p68 = 1
					}
					v6 = t67 + p68
					v5 = v5 + i32(8)
					if uint32(v5) < uint32(v12) {
						goto l16
					}
				}
				v7 = v5 + i32(-4)
			l15:
				if v5 == v11 {
					goto l17
				}
			l18:
				{
					t69 := v7
					v10 = v8 + v6<<2
					t70 := int32(load32(m.memory[uint32(v10):]))
					store32(m.memory[uint32(t69):], uint32(t70))
					t71 := v10
					v7 = v5
					t72 := int32(load32(m.memory[uint32(v7):]))
					v5 = t72
					store32(m.memory[uint32(t71):], uint32(v5))
					t73 := v6
					var p74 int32
					if uint32(v5) < uint32(v9) {
						p74 = 1
					}
					v6 = t73 + p74
					v5 = v7 + i32(4)
					if v5 != v11 {
						goto l18
					}
				}
				v7 = v5 + i32(-4)
			l17:
				t75 := v7
				v5 = v8 + v6<<2
				t76 := int32(load32(m.memory[uint32(v5):]))
				store32(m.memory[uint32(t75):], uint32(t76))
				store32(m.memory[uint32(v5):], uint32(v13))
				t77 := v6
				var p78 int32
				if uint32(v13) < uint32(v9) {
					p78 = 1
				}
				v5 = t77 + p78
				if uint32(v5) >= uint32(v1) {
					goto l12
				}
				t79 := int32(load32(m.memory[uint32(v0):]))
				v7 = t79
				t80 := v0
				v6 = v0 + v5<<2
				t81 := int32(load32(m.memory[uint32(v6):]))
				store32(m.memory[uint32(t80):], uint32(t81))
				store32(m.memory[uint32(v6):], uint32(v7))
				m.fn970(v0, v5, v2, v3)
				v1 = v1 + (v5 ^ i32(-1))
				v0 = v6 + i32(4)
				v2 = v6
			}
		l13:
			if uint32(v1) >= uint32(i32(33)) {
				goto l19
			}
		}
	l0:
		if uint32(v1) < uint32(i32(2)) {
			goto l2
		}
		t82 := v1
		v13 = int32(uint32(v1) >> 1)
		t83 := v13
		var p84 int32
		if uint32(v1) < uint32(i32(18)) {
			p84 = 1
		}
		v14 = p84
		p85 := t83
		if v14 != 0 {
			p85 = t82
		}
		v5 = p85
		v15 = v1 - v13
		v3 = v0 + v13<<2
		v8 = v0
	l28:
		{
			{
				{
					if uint32(v5) > uint32(i32(12)) {
						goto l20
					}
					v6 = i32(1)
					if uint32(v5) <= uint32(i32(8)) {
						goto l21
					}
					t86 := int32(load32(m.memory[int64(uint32(v8))+32:]))
					t87 := v8
					v6 = t86
					t88 := int32(load32(m.memory[int64(uint32(v8))+16:]))
					t89 := v6
					v7 = t88
					p90 := v7
					if uint32(v6) > uint32(v7) {
						p90 = t89
					}
					v9 = p90
					t91 := int32(load32(m.memory[int64(uint32(v8))+12:]))
					t92 := v9
					v10 = t91
					t93 := int32(load32(m.memory[uint32(v8):]))
					t94 := v10
					v11 = t93
					p95 := v11
					if uint32(v10) > uint32(v11) {
						p95 = t94
					}
					v12 = p95
					p96 := v12
					if uint32(v9) > uint32(v12) {
						p96 = t92
					}
					v2 = p96
					t97 := int32(load32(m.memory[int64(uint32(v8))+28:]))
					t98 := v2
					v16 = t97
					t99 := int32(load32(m.memory[int64(uint32(v8))+4:]))
					t100 := v16
					v17 = t99
					p101 := v17
					if uint32(v16) > uint32(v17) {
						p101 = t100
					}
					v18 = p101
					t103 := v18
					p102 := v11
					if uint32(v10) < uint32(v11) {
						p102 = v10
					}
					v10 = p102
					p104 := v10
					if uint32(v18) > uint32(v10) {
						p104 = t103
					}
					v11 = p104
					p105 := v11
					if uint32(v2) > uint32(v11) {
						p105 = t98
					}
					v19 = p105
					t106 := int32(load32(m.memory[int64(uint32(v8))+24:]))
					t107 := v19
					v20 = t106
					t108 := int32(load32(m.memory[int64(uint32(v8))+20:]))
					t109 := v20
					v21 = t108
					t110 := int32(load32(m.memory[int64(uint32(v8))+8:]))
					t111 := v21
					v22 = t110
					p112 := v22
					if uint32(v21) > uint32(v22) {
						p112 = t111
					}
					v23 = p112
					p113 := v23
					if uint32(v20) > uint32(v23) {
						p113 = t109
					}
					v24 = p113
					t115 := v24
					p114 := v12
					if uint32(v9) < uint32(v12) {
						p114 = v9
					}
					v9 = p114
					t117 := v9
					p116 := v17
					if uint32(v16) < uint32(v17) {
						p116 = v16
					}
					v12 = p116
					p118 := v12
					if uint32(v9) > uint32(v12) {
						p118 = t117
					}
					v16 = p118
					p119 := v16
					if uint32(v24) > uint32(v16) {
						p119 = t115
					}
					v17 = p119
					p120 := v17
					if uint32(v19) > uint32(v17) {
						p120 = t107
					}
					store32(m.memory[int64(uint32(t87))+32:], uint32(p120))
					t122 := v8
					p121 := v23
					if uint32(v20) < uint32(v23) {
						p121 = v20
					}
					v20 = p121
					t124 := v20
					p123 := v7
					if uint32(v6) < uint32(v7) {
						p123 = v6
					}
					v6 = p123
					t126 := v6
					p125 := v22
					if uint32(v21) < uint32(v22) {
						p125 = v21
					}
					v7 = p125
					p127 := v7
					if uint32(v6) > uint32(v7) {
						p127 = t126
					}
					v21 = p127
					p128 := v21
					if uint32(v20) < uint32(v21) {
						p128 = t124
					}
					v22 = p128
					t130 := v22
					p129 := v12
					if uint32(v9) < uint32(v12) {
						p129 = v9
					}
					v9 = p129
					p131 := v9
					if uint32(v22) < uint32(v9) {
						p131 = t130
					}
					v12 = p131
					t133 := v12
					p132 := v7
					if uint32(v6) < uint32(v7) {
						p132 = v6
					}
					v6 = p132
					t135 := v6
					p134 := v10
					if uint32(v18) < uint32(v10) {
						p134 = v18
					}
					v7 = p134
					p136 := v7
					if uint32(v6) < uint32(v7) {
						p136 = t135
					}
					v10 = p136
					p137 := v10
					if uint32(v12) < uint32(v10) {
						p137 = t133
					}
					store32(m.memory[uint32(t122):], uint32(p137))
					t139 := v8
					p138 := v11
					if uint32(v2) < uint32(v11) {
						p138 = v2
					}
					v11 = p138
					t141 := v11
					p140 := v21
					if uint32(v20) > uint32(v21) {
						p140 = v20
					}
					v2 = p140
					p142 := v2
					if uint32(v11) > uint32(v2) {
						p142 = t141
					}
					v18 = p142
					t144 := v18
					p143 := v17
					if uint32(v19) < uint32(v17) {
						p143 = v19
					}
					v17 = p143
					p145 := v17
					if uint32(v18) > uint32(v17) {
						p145 = t144
					}
					store32(m.memory[int64(uint32(t139))+28:], uint32(p145))
					t147 := v8
					p146 := v17
					if uint32(v18) < uint32(v17) {
						p146 = v18
					}
					v17 = p146
					t149 := v17
					p148 := v2
					if uint32(v11) < uint32(v2) {
						p148 = v11
					}
					v11 = p148
					t151 := v11
					p150 := v16
					if uint32(v24) < uint32(v16) {
						p150 = v24
					}
					v2 = p150
					p152 := v2
					if uint32(v11) > uint32(v2) {
						p152 = t151
					}
					v16 = p152
					t154 := v16
					p153 := v9
					if uint32(v22) > uint32(v9) {
						p153 = v22
					}
					v9 = p153
					t156 := v9
					p155 := v7
					if uint32(v6) > uint32(v7) {
						p155 = v6
					}
					v6 = p155
					p157 := v6
					if uint32(v9) > uint32(v6) {
						p157 = t156
					}
					v7 = p157
					p158 := v7
					if uint32(v16) > uint32(v7) {
						p158 = t154
					}
					v18 = p158
					p159 := v18
					if uint32(v17) > uint32(v18) {
						p159 = t149
					}
					store32(m.memory[int64(uint32(t147))+24:], uint32(p159))
					t161 := v8
					p160 := v18
					if uint32(v17) < uint32(v18) {
						p160 = v17
					}
					store32(m.memory[int64(uint32(t161))+20:], uint32(p160))
					t163 := v8
					p162 := v7
					if uint32(v16) < uint32(v7) {
						p162 = v16
					}
					v7 = p162
					t165 := v7
					p164 := v2
					if uint32(v11) < uint32(v2) {
						p164 = v11
					}
					v11 = p164
					t167 := v11
					p166 := v6
					if uint32(v9) < uint32(v6) {
						p166 = v9
					}
					v6 = p166
					p168 := v6
					if uint32(v11) > uint32(v6) {
						p168 = t167
					}
					v9 = p168
					p169 := v9
					if uint32(v7) > uint32(v9) {
						p169 = t165
					}
					store32(m.memory[int64(uint32(t163))+16:], uint32(p169))
					t171 := v8
					p170 := v9
					if uint32(v7) < uint32(v9) {
						p170 = v7
					}
					store32(m.memory[int64(uint32(t171))+12:], uint32(p170))
					t173 := v8
					p172 := v6
					if uint32(v11) < uint32(v6) {
						p172 = v11
					}
					v6 = p172
					t175 := v6
					p174 := v10
					if uint32(v12) > uint32(v10) {
						p174 = v12
					}
					v7 = p174
					p176 := v7
					if uint32(v6) > uint32(v7) {
						p176 = t175
					}
					store32(m.memory[int64(uint32(t173))+8:], uint32(p176))
					t178 := v8
					p177 := v7
					if uint32(v6) < uint32(v7) {
						p177 = v6
					}
					store32(m.memory[int64(uint32(t178))+4:], uint32(p177))
					v6 = i32(9)
					goto l21
				}
			l20:
				t179 := int32(load32(m.memory[int64(uint32(v8))+48:]))
				t180 := v8
				v6 = t179
				t181 := int32(load32(m.memory[uint32(v8):]))
				t182 := v6
				v7 = t181
				p183 := v7
				if uint32(v6) > uint32(v7) {
					p183 = t182
				}
				v9 = p183
				t184 := int32(load32(m.memory[int64(uint32(v8))+44:]))
				t185 := v9
				v10 = t184
				t186 := int32(load32(m.memory[int64(uint32(v8))+20:]))
				t187 := v10
				v11 = t186
				p188 := v11
				if uint32(v10) > uint32(v11) {
					p188 = t187
				}
				v12 = p188
				t189 := int32(load32(m.memory[int64(uint32(v8))+16:]))
				t190 := v12
				v2 = t189
				p191 := v2
				if uint32(v12) > uint32(v2) {
					p191 = t190
				}
				v16 = p191
				p192 := v16
				if uint32(v9) > uint32(v16) {
					p192 = t185
				}
				v17 = p192
				t193 := int32(load32(m.memory[int64(uint32(v8))+40:]))
				t194 := v17
				v18 = t193
				t195 := int32(load32(m.memory[int64(uint32(v8))+4:]))
				t196 := v18
				v19 = t195
				p197 := v19
				if uint32(v18) > uint32(v19) {
					p197 = t196
				}
				v20 = p197
				t198 := int32(load32(m.memory[int64(uint32(v8))+32:]))
				t199 := v20
				v21 = t198
				t200 := int32(load32(m.memory[int64(uint32(v8))+24:]))
				t201 := v21
				v22 = t200
				p202 := v22
				if uint32(v21) > uint32(v22) {
					p202 = t201
				}
				v23 = p202
				p203 := v23
				if uint32(v20) > uint32(v23) {
					p203 = t199
				}
				v24 = p203
				t204 := int32(load32(m.memory[int64(uint32(v8))+36:]))
				t205 := v24
				v25 = t204
				t206 := int32(load32(m.memory[int64(uint32(v8))+8:]))
				t207 := v25
				v26 = t206
				p208 := v26
				if uint32(v25) > uint32(v26) {
					p208 = t207
				}
				v27 = p208
				t209 := int32(load32(m.memory[int64(uint32(v8))+28:]))
				t210 := v27
				v28 = t209
				t211 := int32(load32(m.memory[int64(uint32(v8))+12:]))
				t212 := v28
				v29 = t211
				p213 := v29
				if uint32(v28) > uint32(v29) {
					p213 = t212
				}
				v30 = p213
				p214 := v30
				if uint32(v27) > uint32(v30) {
					p214 = t210
				}
				v31 = p214
				p215 := v31
				if uint32(v24) > uint32(v31) {
					p215 = t205
				}
				v32 = p215
				p216 := v32
				if uint32(v17) > uint32(v32) {
					p216 = t194
				}
				store32(m.memory[int64(uint32(t180))+48:], uint32(p216))
				t218 := v8
				p217 := v16
				if uint32(v9) < uint32(v16) {
					p217 = v9
				}
				v9 = p217
				t220 := v9
				p219 := v23
				if uint32(v20) < uint32(v23) {
					p219 = v20
				}
				v16 = p219
				t222 := v16
				p221 := v30
				if uint32(v27) < uint32(v30) {
					p221 = v27
				}
				v20 = p221
				p223 := v20
				if uint32(v16) > uint32(v20) {
					p223 = t222
				}
				v23 = p223
				p224 := v23
				if uint32(v9) > uint32(v23) {
					p224 = t220
				}
				v27 = p224
				t226 := v27
				p225 := v22
				if uint32(v21) < uint32(v22) {
					p225 = v21
				}
				v21 = p225
				t228 := v21
				p227 := v19
				if uint32(v18) < uint32(v19) {
					p227 = v18
				}
				v18 = p227
				p229 := v18
				if uint32(v21) > uint32(v18) {
					p229 = t228
				}
				v19 = p229
				t231 := v19
				p230 := v29
				if uint32(v28) < uint32(v29) {
					p230 = v28
				}
				v22 = p230
				t233 := v22
				p232 := v26
				if uint32(v25) < uint32(v26) {
					p232 = v25
				}
				v25 = p232
				p234 := v25
				if uint32(v22) > uint32(v25) {
					p234 = t233
				}
				v26 = p234
				p235 := v26
				if uint32(v19) > uint32(v26) {
					p235 = t231
				}
				v28 = p235
				t237 := v28
				p236 := v2
				if uint32(v12) < uint32(v2) {
					p236 = v12
				}
				v12 = p236
				t239 := v12
				p238 := v7
				if uint32(v6) < uint32(v7) {
					p238 = v6
				}
				v6 = p238
				p240 := v6
				if uint32(v12) > uint32(v6) {
					p240 = t239
				}
				v7 = p240
				p241 := v7
				if uint32(v28) > uint32(v7) {
					p241 = t237
				}
				v2 = p241
				p242 := v2
				if uint32(v27) > uint32(v2) {
					p242 = t226
				}
				v29 = p242
				t244 := v29
				p243 := v32
				if uint32(v17) < uint32(v32) {
					p243 = v17
				}
				v17 = p243
				t246 := v17
				p245 := v31
				if uint32(v24) < uint32(v31) {
					p245 = v24
				}
				v24 = p245
				t248 := v24
				p247 := v11
				if uint32(v10) < uint32(v11) {
					p247 = v10
				}
				v10 = p247
				p249 := v10
				if uint32(v24) > uint32(v10) {
					p249 = t248
				}
				v11 = p249
				p250 := v11
				if uint32(v17) > uint32(v11) {
					p250 = t246
				}
				v30 = p250
				p251 := v30
				if uint32(v29) > uint32(v30) {
					p251 = t244
				}
				store32(m.memory[int64(uint32(t218))+44:], uint32(p251))
				t253 := v8
				p252 := v25
				if uint32(v22) < uint32(v25) {
					p252 = v22
				}
				v22 = p252
				t255 := v22
				p254 := v18
				if uint32(v21) < uint32(v18) {
					p254 = v21
				}
				v18 = p254
				p256 := v18
				if uint32(v22) < uint32(v18) {
					p256 = t255
				}
				v21 = p256
				t258 := v21
				p257 := v10
				if uint32(v24) < uint32(v10) {
					p257 = v24
				}
				v10 = p257
				t260 := v10
				p259 := v6
				if uint32(v12) < uint32(v6) {
					p259 = v12
				}
				v6 = p259
				p261 := v6
				if uint32(v10) < uint32(v6) {
					p261 = t260
				}
				v12 = p261
				p262 := v12
				if uint32(v21) < uint32(v12) {
					p262 = t258
				}
				store32(m.memory[uint32(t253):], uint32(p262))
				t264 := v8
				p263 := v30
				if uint32(v29) < uint32(v30) {
					p263 = v29
				}
				v24 = p263
				t266 := v24
				p265 := v11
				if uint32(v17) < uint32(v11) {
					p265 = v17
				}
				v11 = p265
				t268 := v11
				p267 := v2
				if uint32(v27) < uint32(v2) {
					p267 = v27
				}
				v2 = p267
				p269 := v2
				if uint32(v11) > uint32(v2) {
					p269 = t268
				}
				v17 = p269
				p270 := v17
				if uint32(v24) > uint32(v17) {
					p270 = t266
				}
				store32(m.memory[int64(uint32(t264))+40:], uint32(p270))
				t272 := v8
				p271 := v20
				if uint32(v16) < uint32(v20) {
					p271 = v16
				}
				v16 = p271
				t274 := v16
				p273 := v7
				if uint32(v28) < uint32(v7) {
					p273 = v28
				}
				v7 = p273
				p275 := v7
				if uint32(v16) < uint32(v7) {
					p275 = t274
				}
				v20 = p275
				t277 := v20
				p276 := v6
				if uint32(v10) > uint32(v6) {
					p276 = v10
				}
				v6 = p276
				t279 := v6
				p278 := v18
				if uint32(v22) > uint32(v18) {
					p278 = v22
				}
				v10 = p278
				p280 := v10
				if uint32(v6) < uint32(v10) {
					p280 = t279
				}
				v18 = p280
				p281 := v18
				if uint32(v20) < uint32(v18) {
					p281 = t277
				}
				v22 = p281
				t283 := v22
				p282 := v23
				if uint32(v9) < uint32(v23) {
					p282 = v9
				}
				v9 = p282
				t285 := v9
				p284 := v26
				if uint32(v19) < uint32(v26) {
					p284 = v19
				}
				v19 = p284
				p286 := v19
				if uint32(v9) < uint32(v19) {
					p286 = t285
				}
				v23 = p286
				t288 := v23
				p287 := v12
				if uint32(v21) > uint32(v12) {
					p287 = v21
				}
				v12 = p287
				p289 := v12
				if uint32(v23) < uint32(v12) {
					p289 = t288
				}
				v21 = p289
				p290 := v21
				if uint32(v22) < uint32(v21) {
					p290 = t283
				}
				store32(m.memory[int64(uint32(t272))+4:], uint32(p290))
				t292 := v8
				p291 := v17
				if uint32(v24) < uint32(v17) {
					p291 = v24
				}
				v17 = p291
				t294 := v17
				p293 := v19
				if uint32(v9) > uint32(v19) {
					p293 = v9
				}
				v9 = p293
				t296 := v9
				p295 := v7
				if uint32(v16) > uint32(v7) {
					p295 = v16
				}
				v7 = p295
				p297 := v7
				if uint32(v9) > uint32(v7) {
					p297 = t296
				}
				v16 = p297
				t299 := v16
				p298 := v2
				if uint32(v11) < uint32(v2) {
					p298 = v11
				}
				v11 = p298
				t301 := v11
				p300 := v10
				if uint32(v6) > uint32(v10) {
					p300 = v6
				}
				v6 = p300
				p302 := v6
				if uint32(v11) > uint32(v6) {
					p302 = t301
				}
				v10 = p302
				p303 := v10
				if uint32(v16) > uint32(v10) {
					p303 = t299
				}
				v2 = p303
				p304 := v2
				if uint32(v17) > uint32(v2) {
					p304 = t294
				}
				store32(m.memory[int64(uint32(t292))+36:], uint32(p304))
				t306 := v8
				p305 := v2
				if uint32(v17) < uint32(v2) {
					p305 = v17
				}
				store32(m.memory[int64(uint32(t306))+32:], uint32(p305))
				t308 := v8
				p307 := v7
				if uint32(v9) < uint32(v7) {
					p307 = v9
				}
				v7 = p307
				t310 := v7
				p309 := v6
				if uint32(v11) < uint32(v6) {
					p309 = v11
				}
				v6 = p309
				p311 := v6
				if uint32(v7) > uint32(v6) {
					p311 = t310
				}
				v9 = p311
				t313 := v9
				p312 := v10
				if uint32(v16) < uint32(v10) {
					p312 = v16
				}
				v10 = p312
				p314 := v10
				if uint32(v9) > uint32(v10) {
					p314 = t313
				}
				store32(m.memory[int64(uint32(t308))+28:], uint32(p314))
				t316 := v8
				p315 := v18
				if uint32(v20) > uint32(v18) {
					p315 = v20
				}
				v11 = p315
				t318 := v11
				p317 := v12
				if uint32(v23) > uint32(v12) {
					p317 = v23
				}
				v12 = p317
				p319 := v12
				if uint32(v11) < uint32(v12) {
					p319 = t318
				}
				v2 = p319
				t321 := v2
				p320 := v21
				if uint32(v22) > uint32(v21) {
					p320 = v22
				}
				v16 = p320
				p322 := v16
				if uint32(v2) < uint32(v16) {
					p322 = t321
				}
				store32(m.memory[int64(uint32(t316))+8:], uint32(p322))
				t324 := v8
				p323 := v10
				if uint32(v9) < uint32(v10) {
					p323 = v9
				}
				v9 = p323
				t326 := v9
				p325 := v6
				if uint32(v7) < uint32(v6) {
					p325 = v7
				}
				v6 = p325
				t328 := v6
				p327 := v12
				if uint32(v11) > uint32(v12) {
					p327 = v11
				}
				v7 = p327
				p329 := v7
				if uint32(v6) > uint32(v7) {
					p329 = t328
				}
				v10 = p329
				p330 := v10
				if uint32(v9) > uint32(v10) {
					p330 = t326
				}
				store32(m.memory[int64(uint32(t324))+24:], uint32(p330))
				t332 := v8
				p331 := v10
				if uint32(v9) < uint32(v10) {
					p331 = v9
				}
				store32(m.memory[int64(uint32(t332))+20:], uint32(p331))
				t334 := v8
				p333 := v7
				if uint32(v6) < uint32(v7) {
					p333 = v6
				}
				v6 = p333
				t336 := v6
				p335 := v16
				if uint32(v2) > uint32(v16) {
					p335 = v2
				}
				v7 = p335
				p337 := v7
				if uint32(v6) > uint32(v7) {
					p337 = t336
				}
				store32(m.memory[int64(uint32(t334))+16:], uint32(p337))
				t339 := v8
				p338 := v7
				if uint32(v6) < uint32(v7) {
					p338 = v6
				}
				store32(m.memory[int64(uint32(t339))+12:], uint32(p338))
				v6 = i32(13)
			}
		l21:
			if uint32(v6) > uint32(v5) {
				goto l12
			}
			{
				if v6 == v5 {
					goto l22
				}
				v12 = v8 + v5<<2
				t340 := v8
				v11 = v6 << 2
				v10 = t340 + v11
			l27:
				{
					t341 := int32(load32(m.memory[uint32(v10):]))
					v9 = t341
					t342 := int32(load32(m.memory[uint32(v10+i32(-4)):]))
					t343 := v9
					v6 = t342
					if uint32(t343) >= uint32(v6) {
						goto l23
					}
					v5 = v11
				l26:
					{
						v7 = v8 + v5
						store32(m.memory[uint32(v7):], uint32(v6))
						if v5 != i32(4) {
							goto l24
						}
						v5 = v8
						goto l25
					l24:
						v5 = v5 + i32(-4)
						t344 := int32(load32(m.memory[uint32(v7+i32(-8)):]))
						t345 := v9
						v6 = t344
						if uint32(t345) < uint32(v6) {
							goto l26
						}
					}
					v5 = v8 + v5
				l25:
					store32(m.memory[uint32(v5):], uint32(v9))
				}
			l23:
				v11 = v11 + i32(4)
				v10 = v10 + i32(4)
				if v10 != v12 {
					goto l27
				}
			}
		l22:
			if v14 != 0 {
				goto l2
			}
			var p346 int32
			if v8 == v0 {
				p346 = 1
			}
			v6 = p346
			v5 = v15
			v8 = v3
			if v6 != 0 {
				goto l28
			}
		}
		v8 = v3 + i32(-4)
		t347 := v0
		v5 = v1<<2 + i32(-4)
		v9 = t347 + v5
		v11 = v4 + v5
		v10 = v4
		v7 = v0
	l29:
		{
			t348 := int32(load32(m.memory[uint32(v3):]))
			t349 := v10
			v12 = t348
			t350 := int32(load32(m.memory[uint32(v7):]))
			t351 := v12
			v2 = t350
			t352 := v2
			var p353 int32
			if uint32(v12) < uint32(v2) {
				p353 = 1
			}
			v16 = p353
			p354 := t352
			if v16 != 0 {
				p354 = t351
			}
			store32(m.memory[uint32(t349):], uint32(p354))
			t355 := int32(load32(m.memory[uint32(v9):]))
			t356 := v11
			v5 = t355
			t357 := int32(load32(m.memory[uint32(v8):]))
			t358 := v5
			v6 = t357
			p359 := v6
			if uint32(v5) > uint32(v6) {
				p359 = t358
			}
			store32(m.memory[uint32(t356):], uint32(p359))
			v11 = v11 + i32(-4)
			v10 = v10 + i32(4)
			t361 := v8
			p360 := i32(0)
			if uint32(v5) < uint32(v6) {
				p360 = i32(-4)
			}
			v8 = t361 + p360
			t363 := v9
			p362 := i32(0)
			if uint32(v5) >= uint32(v6) {
				p362 = i32(-4)
			}
			v9 = t363 + p362
			t364 := v7
			var p365 int32
			if uint32(v12) >= uint32(v2) {
				p365 = 1
			}
			v7 = t364 + p365<<2
			v3 = v3 + v16<<2
			v13 = v13 + i32(-1)
			if v13 != 0 {
				goto l29
			}
		}
		v5 = v8 + i32(4)
		{
			if v1&i32(1) == 0 {
				goto l30
			}
			t366 := v10
			t367 := v7
			t368 := v3
			var p369 int32
			if uint32(v7) < uint32(v5) {
				p369 = 1
			}
			v6 = p369
			p370 := t368
			if v6 != 0 {
				p370 = t367
			}
			t371 := int32(load32(m.memory[uint32(p370):]))
			store32(m.memory[uint32(t366):], uint32(t371))
			t372 := v3
			var p373 int32
			if uint32(v7) >= uint32(v5) {
				p373 = 1
			}
			v3 = t372 + p373<<2
			v7 = v7 + v6<<2
		}
	l30:
		if v7 != v5 {
			goto l31
		}
		if v3 != v9+i32(4) {
			goto l31
		}
		v5 = v1 << 2
		if v5 == 0 {
			goto l2
		}
		memory_copy(m.memory, uint32(v0), uint32(v4), uint32(v5))
		goto l2
	}
l12:
	panic("unreachable")
l31:
	m.fn122()
	panic("unreachable")
l2:
	m.g0 = v4 + i32(128)
}
func (m *Module) fn971(v0, v1 int32) {
	var v2, v3, v4, v5, v6, v7, v8 int32
	v2 = int32(uint32(v1)>>1) + v1
l6:
	{
		v2 = v2 + i32(-1)
		if uint32(v2) < uint32(v1) {
			goto l0
		}
		v3 = v2 - v1
		goto l1
	l0:
		t0 := int32(load32(m.memory[uint32(v0):]))
		v4 = t0
		t1 := v0
		v3 = v0 + v2<<2
		t2 := int32(load32(m.memory[uint32(v3):]))
		store32(m.memory[uint32(t1):], uint32(t2))
		store32(m.memory[uint32(v3):], uint32(v4))
		v3 = i32(0)
	}
l1:
	{
		v5 = v3 << 1
		v4 = v5 | i32(1)
		t4 := v4
		p3 := v2
		if uint32(v1) < uint32(v2) {
			p3 = v1
		}
		v6 = p3
		if uint32(t4) >= uint32(v6) {
			goto l2
		}
	l5:
		{
			{
				v5 = v5 + i32(2)
				if uint32(v5) < uint32(v6) {
					goto l3
				}
				goto l4
			l3:
				t5 := int32(load32(m.memory[uint32(v0+v4<<2):]))
				t6 := int32(load32(m.memory[uint32(v0+v5<<2):]))
				t7 := v4
				var p8 int32
				if uint32(t5) < uint32(t6) {
					p8 = 1
				}
				v4 = t7 + p8
			}
		l4:
			v3 = v0 + v3<<2
			t9 := int32(load32(m.memory[uint32(v3):]))
			v5 = t9
			t10 := v5
			v7 = v0 + v4<<2
			t11 := int32(load32(m.memory[uint32(v7):]))
			v8 = t11
			if uint32(t10) >= uint32(v8) {
				goto l2
			}
			store32(m.memory[uint32(v7):], uint32(v5))
			store32(m.memory[uint32(v3):], uint32(v8))
			v3 = v4
			v5 = v4 << 1
			v4 = v5 | i32(1)
			if uint32(v4) < uint32(v6) {
				goto l5
			}
		}
	}
l2:
	if v2 != 0 {
		goto l6
	}
}
func (m *Module) fn972(v0, v1, v2, v3 int32) int32 {
	var v4, v5, v6 int32
	{
		if uint32(v3) < uint32(i32(8)) {
			goto l0
		}
		t0 := v0
		t1 := v0
		v3 = int32(uint32(v3) >> 3)
		v4 = v3 << 4
		t2 := t1 + v4
		t3 := v0
		v5 = v3 * i32(28)
		t4 := m.fn972(t0, t2, t3+v5, v3)
		v0 = t4
		t5 := m.fn972(v1, v1+v4, v1+v5, v3)
		v1 = t5
		t6 := m.fn972(v2, v2+v4, v2+v5, v3)
		v2 = t6
	}
l0:
	t7 := int32(load32(m.memory[uint32(v0):]))
	t8 := v0
	t9 := v2
	t10 := v1
	v3 = t7
	t11 := int32(load32(m.memory[uint32(v1):]))
	t12 := v3
	v4 = t11
	var p13 int32
	if uint32(t12) < uint32(v4) {
		p13 = 1
	}
	v5 = p13
	t14 := int32(load32(m.memory[uint32(v2):]))
	t15 := v5
	t16 := v4
	v6 = t14
	var p17 int32
	if uint32(t16) < uint32(v6) {
		p17 = 1
	}
	p18 := t10
	if t15^p17 != 0 {
		p18 = t9
	}
	t19 := v5
	var p20 int32
	if uint32(v3) < uint32(v6) {
		p20 = 1
	}
	p21 := p18
	if t19^p20 != 0 {
		p21 = t8
	}
	return p21
}
func (m *Module) fn973(v0, v1, v2, v3, v4, v5 int32) {
	var v6, v7 int32
	var v8 int64
	v6 = i32(1)
	v7 = i32(4)
	v8 = int64(uint32(v5)) * int64(uint32(v3))
	if int32(int64(uint64(v8)>>32)) == 0 {
		goto l0
	}
	v3 = i32(0)
	goto l1
l0:
	v3 = int32(v8)
	if uint32(v3) <= uint32(i32(-0x80000000)-v4) {
		goto l2
	}
	v3 = i32(0)
	goto l1
l2:
	{
		{
			if v1 == 0 {
				goto l3
			}
			t0 := m.fn22(v2, v5*v1, v4, v3)
			v7 = t0
			goto l4
		}
	l3:
		if v3 != 0 {
			goto l5
		}
		v7 = v4
		goto l6
	l5:
		t1 := m.fn20(v3, v4)
		v7 = t1
	}
l4:
	if v7 != 0 {
		goto l6
	}
	store32(m.memory[int64(uint32(v0))+4:], uint32(v4))
	goto l7
l6:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v7))
	v6 = i32(0)
l7:
	v7 = i32(8)
l1:
	store32(m.memory[uint32(v0+v7):], uint32(v3))
	store32(m.memory[uint32(v0):], uint32(v6))
}
func (m *Module) fn974(v0, v1, v2 int32) int32 {
	var v3, v4, v5 int32
	v3 = i32(0)
	if v2 == 0 {
		goto l0
	}
l2:
	{
		t0 := int32(m.memory[uint32(v0)])
		v4 = t0
		t1 := int32(m.memory[uint32(v1)])
		t2 := v4
		v5 = t1
		if t2 != v5 {
			goto l1
		}
		v0 = v0 + i32(1)
		v1 = v1 + i32(1)
		v2 = v2 + i32(-1)
		if v2 == 0 {
			goto l0
		}
		goto l2
	}
l1:
	v3 = v4 - v5
l0:
	return v3
}
func (m *Module) fn975(v0 int32) int32 {
	var v1, v2 int32
	{
		t0 := int32(m.memory[uint32(v0)])
		if t0 != 0 {
			goto l0
		}
		v0 = i32(0)
		goto l1
	}
l0:
	v1 = v0 + i32(1)
	v0 = i32(0)
l2:
	{
		v2 = v0
		v0 = v2 + i32(1)
		t1 := int32(m.memory[uint32(v1+v2)])
		if t1 != 0 {
			goto l2
		}
	}
l1:
	return v0
}
func (m *Module) fn976(v0 int32, v1, v2, v3, v4 int64) {
	var v5, v6, v7, v8, v9, v10 int64
	t0 := v0
	v5 = v3 & i64(0xffffffff)
	t1 := v5
	v6 = v1 & i64(0xffffffff)
	v7 = t1 * v6
	t2 := v7
	v8 = int64(uint64(v3) >> 32)
	v6 = v8 * v6
	t3 := v6
	t4 := v5
	v9 = int64(uint64(v1) >> 32)
	v5 = t3 + t4*v9
	v10 = t2 + v5<<32
	store64(m.memory[uint32(t0):], uint64(v10))
	t5 := v0
	t6 := v8 * v9
	var p7 int32
	if uint64(v5) < uint64(v6) {
		p7 = 1
	}
	t8 := t6 + (int64(uint32(p7))<<32 | int64(uint64(v5)>>32))
	var p9 int32
	if uint64(v10) < uint64(v7) {
		p9 = 1
	}
	store64(m.memory[int64(uint32(t5))+8:], uint64(t8+int64(uint32(p9))+(v4*v1+v3*v2)))
}
func fn977(v0 float64) float64 {
	var v1 int64
	var v2 int32
	{
		v0 = float64(v0 + math.Copysign(float64(0.49999999999999994), v0))
		v1 = int64(math.Float64bits(v0))
		v2 = int32(int64(uint64(v1)>>52)) & i32(2047)
		if uint32(v2) > uint32(i32(1074)) {
			goto l0
		}
		p0 := i64_shr_s(i64(-0x10000000000000), int64(uint32(v2+i32(-1023))))
		if uint32(v2) < uint32(i32(1023)) {
			p0 = i64(-0x8000000000000000)
		}
		v0 = math.Float64frombits(uint64(p0 & v1))
	}
l0:
	return v0
}
func (m *Module) Xmemory() Memory {
	return (*wasmMemory)(&m.memory)
}

// Compiler error if endianess is unknown.
var _ = map[bool]struct{}{big: {}, little: {}}

// go.dev/src/cmd/compile/internal/ssa/config.go
const (
	big = false ||
		runtime.GOARCH == "ppc64" || runtime.GOARCH == "s390x" ||
		runtime.GOARCH == "mips" || runtime.GOARCH == "mips64"

	little = false ||
		runtime.GOARCH == "386" || runtime.GOARCH == "amd64" ||
		runtime.GOARCH == "arm" || runtime.GOARCH == "arm64" ||
		runtime.GOARCH == "riscv64" || runtime.GOARCH == "wasm" ||
		runtime.GOARCH == "ppc64le" || runtime.GOARCH == "loong64" ||
		runtime.GOARCH == "mipsle" || runtime.GOARCH == "mips64le"

	unalignedOK = false ||
		runtime.GOARCH == "386" || runtime.GOARCH == "amd64" ||
		runtime.GOARCH == "arm64" || runtime.GOARCH == "loong64" ||
		runtime.GOARCH == "ppc64" || runtime.GOARCH == "ppc64le" ||
		runtime.GOARCH == "s390x" || runtime.GOARCH == "wasm"
)

//go:nosplit
func load16(b []byte) uint16 {
	if !unalignedOK {
		return binary.LittleEndian.Uint16(b)
	}
	v := *(*uint16)(unsafe.Pointer((*[2]byte)(b)))
	if big {
		return bits.ReverseBytes16(v)
	}
	return v
}

//go:nosplit
func store16(b []byte, v uint16) {
	if !unalignedOK {
		binary.LittleEndian.PutUint16(b, v)
		return
	}
	if big {
		v = bits.ReverseBytes16(v)
	}
	*(*uint16)(unsafe.Pointer((*[2]byte)(b))) = v
}

//go:nosplit
func load32(b []byte) uint32 {
	if !unalignedOK {
		return binary.LittleEndian.Uint32(b)
	}
	v := *(*uint32)(unsafe.Pointer((*[4]byte)(b)))
	if big {
		return bits.ReverseBytes32(v)
	}
	return v
}

//go:nosplit
func store32(b []byte, v uint32) {
	if !unalignedOK {
		binary.LittleEndian.PutUint32(b, v)
		return
	}
	if big {
		v = bits.ReverseBytes32(v)
	}
	*(*uint32)(unsafe.Pointer((*[4]byte)(b))) = v
}

//go:nosplit
func load64(b []byte) uint64 {
	if !unalignedOK {
		return binary.LittleEndian.Uint64(b)
	}
	v := *(*uint64)(unsafe.Pointer((*[8]byte)(b)))
	if big {
		return bits.ReverseBytes64(v)
	}
	return v
}

//go:nosplit
func store64(b []byte, v uint64) {
	if !unalignedOK {
		binary.LittleEndian.PutUint64(b, v)
		return
	}
	if big {
		v = bits.ReverseBytes64(v)
	}
	*(*uint64)(unsafe.Pointer((*[8]byte)(b))) = v
}

//go:nosplit
func i32(x int32) int32 { return x }

//go:nosplit
func i64(x int64) int64 { return x }
