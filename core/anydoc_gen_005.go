package core

import (
	"math"
	"math/bits"
)

func (m *Module) fn177(v0, v1 int32) {
	var v2, v3, v4 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	m.fn261(v2+i32(8), v1)
	{
		t1 := int32(load32(m.memory[int64(uint32(v2))+8:]))
		v3 = t1
		if v3 == 0 {
			goto l0
		}
		t2 := int32(load32(m.memory[int64(uint32(v2))+12:]))
		store32(m.memory[int64(uint32(v0))+8:], uint32(t2))
		t3 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		t4 := v0
		v4 = t3
		store32(m.memory[uint32(t4):], uint32(v4))
		store32(m.memory[int64(uint32(v1))+12:], uint32(v4+i32(1)))
	}
l0:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
	m.g0 = v2 + i32(16)
}
func (m *Module) fn178(v0, v1, v2 int32) {
	var v3, v4, v5, v6, v7 int32
	t0 := m.g0
	v3 = t0 - i32(48)
	m.g0 = v3
	m.fn59(v3, v2, i32(8), i32(24))
	store32(m.memory[int64(uint32(v3))+20:], uint32(i32(0)))
	t1 := int64(load64(m.memory[uint32(v3):]))
	store64(m.memory[int64(uint32(v3))+12:], uint64(t1))
	m.fn262(v3+i32(12), v2)
	p2 := i32(1)
	if uint32(v2) > uint32(i32(1)) {
		p2 = v2
	}
	v4 = p2
	v5 = v4 + i32(-1)
	t3 := int32(load32(m.memory[int64(uint32(v3))+16:]))
	t4 := int32(load32(m.memory[int64(uint32(v3))+20:]))
	v6 = t4
	v7 = t3 + v6*i32(24)
l3:
	{
		if v5 != 0 {
			m.fn219(v3+i32(24), v1)
			t9 := int64(load64(m.memory[int64(uint32(v3))+40:]))
			store64(m.memory[int64(uint32(v7))+16:], uint64(t9))
			t10 := int64(load64(m.memory[int64(uint32(v3))+32:]))
			store64(m.memory[int64(uint32(v7))+8:], uint64(t10))
			t11 := int64(load64(m.memory[int64(uint32(v3))+24:]))
			store64(m.memory[uint32(v7):], uint64(t11))
			v5 = v5 + i32(-1)
			v7 = v7 + i32(24)
			goto l3
		}
		v5 = v6 + v4
		{
			if v2 != 0 {
				goto l1
			}
			m.fn182(v1)
			v5 = v5 + i32(-1)
			goto l2
		l1:
			t5 := int64(load64(m.memory[int64(uint32(v1))+16:]))
			store64(m.memory[int64(uint32(v7))+16:], uint64(t5))
			t6 := int64(load64(m.memory[int64(uint32(v1))+8:]))
			store64(m.memory[int64(uint32(v7))+8:], uint64(t6))
			t7 := int64(load64(m.memory[uint32(v1):]))
			store64(m.memory[uint32(v7):], uint64(t7))
		}
	l2:
		t8 := int64(load64(m.memory[int64(uint32(v3))+12:]))
		store64(m.memory[uint32(v0):], uint64(t8))
		store32(m.memory[int64(uint32(v0))+8:], uint32(v5))
		m.g0 = v3 + i32(48)
		return
	}
}
func (m *Module) fn179(v0, v1 int32) {
	var v2, v3, v4, v5, v6, v7 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	{
		t1 := int32(load32(m.memory[int64(uint32(v1))+16:]))
		v3 = t1
		if v3 == 0 {
			goto l0
		}
		store32(m.memory[int64(uint32(v1))+16:], uint32(v3+i32(-1)))
		{
			{
				t2 := int32(load32(m.memory[int64(uint32(v1))+12:]))
				v4 = t2
				if v4 != 0 {
					goto l1
				}
				m.fn261(v2+i32(8), v1)
				t3 := int32(load32(m.memory[int64(uint32(v2))+12:]))
				v5 = t3
				t4 := int32(load32(m.memory[int64(uint32(v2))+8:]))
				v3 = t4
				goto l2
			}
		l1:
			v6 = i32(0)
			store32(m.memory[int64(uint32(v1))+12:], uint32(i32(0)))
			t5 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			v5 = t5
			v3 = i32(0)
			{
				t6 := int32(load32(m.memory[int64(uint32(v1))+4:]))
				v7 = t6
				if uint32(v7) < uint32(v4) {
					goto l3
				}
				v6 = i32(0)
				v3 = i32(0)
				t7 := v5
				v7 = v7 - v4
				if uint32(t7) > uint32(v7) {
					goto l3
				}
				t8 := int32(load32(m.memory[uint32(v1):]))
				t9 := v1
				v3 = t8 + v4<<2
				store32(m.memory[uint32(t9):], uint32(v3+i32(4)))
				v6 = v7 + i32(-1)
			}
		l3:
			store32(m.memory[int64(uint32(v1))+4:], uint32(v6))
		}
	l2:
		if v3 == 0 {
			goto l0
		}
		{
			t10 := int32(load32(m.memory[int64(uint32(v1))+32:]))
			v4 = t10
			if v4 == 0 {
				goto l4
			}
			store32(m.memory[int64(uint32(v1))+32:], uint32(v4+i32(-1)))
			{
				{
					t11 := int32(load32(m.memory[int64(uint32(v1))+28:]))
					v4 = t11
					if v4 != 0 {
						goto l5
					}
					t12 := int32(load32(m.memory[int64(uint32(v1))+20:]))
					v4 = t12
					t13 := int32(load32(m.memory[int64(uint32(v1))+24:]))
					if v4 != t13 {
						goto l6
					}
					goto l4
				}
			l5:
				store32(m.memory[int64(uint32(v1))+28:], uint32(i32(0)))
				t14 := int32(load32(m.memory[int64(uint32(v1))+24:]))
				t15 := v4
				v6 = t14
				t16 := int32(load32(m.memory[int64(uint32(v1))+20:]))
				t17 := v6
				v7 = t16
				if uint32(t15) >= uint32(int32(uint32(t17-v7)>>2)) {
					goto l7
				}
				v4 = v7 + v4<<2
			}
		l6:
			store32(m.memory[int64(uint32(v0))+8:], uint32(v4))
			store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
			store32(m.memory[uint32(v0):], uint32(v3))
			store32(m.memory[int64(uint32(v1))+20:], uint32(v4+i32(4)))
			goto l8
		l7:
			store32(m.memory[int64(uint32(v1))+20:], uint32(v6))
		}
	l4:
		store32(m.memory[uint32(v0):], uint32(i32(0)))
		goto l8
	}
l0:
	store32(m.memory[uint32(v0):], uint32(i32(0)))
l8:
	m.g0 = v2 + i32(16)
}
func (m *Module) fn180(v0, v1, v2, v3, v4, v5 int32) {
	if uint32(v4) < uint32(v3) {
		goto l0
	}
	if uint32(v4) <= uint32(v2) {
		goto l1
	}
l0:
	m.fn151(v3, v4, v2, v5)
	panic("unreachable")
l1:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v4-v3))
	store32(m.memory[uint32(v0):], uint32(v1+v3*i32(24)))
}
func (m *Module) fn181(v0, v1 int32) int32 {
	var v2, v3 int32
	{
		t0 := int32(m.memory[uint32(v0)])
		v2 = t0
		t1 := int32(m.memory[uint32(v1)])
		if v2 != t1 {
			goto l0
		}
		v3 = i32(1)
		switch v2 {
		case 8:
			goto l9
		case 7:
			t2 := int32(m.memory[int64(uint32(v0))+1])
			t3 := int32(m.memory[int64(uint32(v1))+1])
			var p4 int32
			if t2 == t3 {
				p4 = 1
			}
			return p4
		case 6:
			t5 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			t6 := int32(load32(m.memory[int64(uint32(v0))+12:]))
			t7 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			t8 := int32(load32(m.memory[int64(uint32(v1))+12:]))
			t9 := m.fn191(t5, t6, t7, t8)
			return t9
		case 5:
			t10 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			t11 := int32(load32(m.memory[int64(uint32(v0))+12:]))
			t12 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			t13 := int32(load32(m.memory[int64(uint32(v1))+12:]))
			t14 := m.fn191(t10, t11, t12, t13)
			return t14
		case 4:
			t15 := math.Float64frombits(load64(m.memory[int64(uint32(v0))+8:]))
			t16 := math.Float64frombits(load64(m.memory[int64(uint32(v1))+8:]))
			if t15 != t16 {
				goto l0
			}
			v3 = i32(0)
			t17 := int32(m.memory[int64(uint32(v0))+17])
			t18 := int32(m.memory[int64(uint32(v1))+17])
			if t17 != t18 {
				goto l9
			}
			t19 := int32(m.memory[int64(uint32(v0))+16])
			t20 := int32(m.memory[int64(uint32(v1))+16])
			var p21 int32
			if t19 == t20 {
				p21 = 1
			}
			return p21
		case 3:
			t22 := int32(m.memory[int64(uint32(v0))+1])
			t23 := int32(m.memory[int64(uint32(v1))+1])
			var p24 int32
			if t22 == t23 {
				p24 = 1
			}
			return p24
		case 2:
			t25 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			t26 := int32(load32(m.memory[int64(uint32(v0))+12:]))
			t27 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			t28 := int32(load32(m.memory[int64(uint32(v1))+12:]))
			t29 := m.fn191(t25, t26, t27, t28)
			return t29
		case 1:
			t30 := math.Float64frombits(load64(m.memory[int64(uint32(v0))+8:]))
			t31 := math.Float64frombits(load64(m.memory[int64(uint32(v1))+8:]))
			var p32 int32
			if t30 == t31 {
				p32 = 1
			}
			return p32
		default:
			t33 := int64(load64(m.memory[int64(uint32(v0))+8:]))
			t34 := int64(load64(m.memory[int64(uint32(v1))+8:]))
			var p35 int32
			if t33 == t34 {
				p35 = 1
			}
			return p35
		}
	}
l0:
	v3 = i32(0)
l9:
	return v3
}
func (m *Module) fn182(v0 int32) {
	var v1 int32
	{
		t0 := int32(m.memory[uint32(v0)])
		v1 = t0
		if uint32(v1) > uint32(i32(6)) {
			return
		}
		if i32_shl(i32(1), v1)&i32(100) == 0 {
			return
		}
		t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		t2 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		m.fn16(t1, t2)
	}
}
func (m *Module) fn183(v0, v1, v2, v3, v4 int32) {
	if uint32(v3) < uint32(v1) {
		m.fn151(v1, v3, v3, v4)
		panic("unreachable")
	}
	store32(m.memory[int64(uint32(v0))+4:], uint32(v3-v1))
	store32(m.memory[uint32(v0):], uint32(v2+v1*i32(24)))
}
func (m *Module) fn184(v0, v1, v2 int32) {
	var v3, v4, v5 int32
	t0 := m.g0
	v3 = t0 - i32(32)
	m.g0 = v3
	m.fn262(v0, v2)
	t1 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	v4 = t1
	{
		if v2 == 0 {
			goto l0
		}
		v5 = v4 + v2
		t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		v4 = t2 + v4*i32(24)
	l1:
		{
			m.fn219(v3+i32(8), v1)
			t3 := int64(load64(m.memory[int64(uint32(v3))+24:]))
			store64(m.memory[int64(uint32(v4))+16:], uint64(t3))
			t4 := int64(load64(m.memory[int64(uint32(v3))+16:]))
			store64(m.memory[int64(uint32(v4))+8:], uint64(t4))
			t5 := int64(load64(m.memory[int64(uint32(v3))+8:]))
			store64(m.memory[uint32(v4):], uint64(t5))
			v4 = v4 + i32(24)
			v1 = v1 + i32(24)
			v2 = v2 + i32(-1)
			if v2 != 0 {
				goto l1
			}
		}
		v4 = v5
	}
l0:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v4))
	m.g0 = v3 + i32(32)
}
func (m *Module) fn185(v0 int32) {
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
	v3 = v3 + i32(24)
	goto l1
l0:
	t2 := int32(load32(m.memory[uint32(v0):]))
	m.fn136(t2, v2, i32(8), i32(24))
}
func (m *Module) fn186(v0, v1 int32) int32 {
	t0 := m.fn181(v0, v1)
	return t0 ^ i32(1)
}
func (m *Module) fn187(v0 int32) int32 {
	var v1, v2 int32
	t0 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t1 := int32(load32(m.memory[uint32(v0):]))
	v1 = t1
	v2 = int32(uint32(t0-v1) >> 2)
	t2 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	t3 := v2
	v0 = t2
	p4 := v0
	if uint32(v2) < uint32(v0) {
		p4 = t3
	}
	v0 = p4
	v2 = i32(0)
l1:
	{
		if v0 == 0 {
			return v2
		}
		v0 = v0 + i32(-1)
		t5 := int32(load32(m.memory[uint32(v1):]))
		v2 = t5 + v2
		v1 = v1 + i32(4)
		goto l1
	}
}
func (m *Module) fn188(v0, v1 int32) {
	m.fn608(v0, v1, i32(4), i32(4))
}
func (m *Module) fn189(v0, v1, v2 int32) {
	var v3, v4, v5, v6, v7, v8, v9 int32
	t0 := m.g0
	v3 = t0 - i32(32)
	m.g0 = v3
	m.fn59(v3, v2, i32(4), i32(12))
	store32(m.memory[int64(uint32(v3))+16:], uint32(i32(0)))
	t1 := int64(load64(m.memory[uint32(v3):]))
	store64(m.memory[int64(uint32(v3))+8:], uint64(t1))
	m.fn60(v3+i32(8), v2)
	p2 := i32(1)
	if uint32(v2) > uint32(i32(1)) {
		p2 = v2
	}
	v4 = p2
	v5 = v4 + i32(-1)
	t3 := int32(load32(m.memory[int64(uint32(v3))+12:]))
	t4 := int32(load32(m.memory[int64(uint32(v3))+16:]))
	v6 = t4
	v7 = t3 + v6*i32(12)
	t5 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	v8 = t5
	t6 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	v9 = t6
l3:
	{
		if v5 != 0 {
			m.fn31(v3+i32(20), v9, v8)
			t11 := int32(load32(m.memory[int64(uint32(v3))+28:]))
			store32(m.memory[int64(uint32(v7))+8:], uint32(t11))
			t12 := int64(load64(m.memory[int64(uint32(v3))+20:]))
			store64(m.memory[uint32(v7):], uint64(t12))
			v5 = v5 + i32(-1)
			v7 = v7 + i32(12)
			goto l3
		}
		v5 = v6 + v4
		{
			{
				if v2 != 0 {
					goto l1
				}
				t7 := int32(load32(m.memory[uint32(v1):]))
				m.fn16(t7, v9)
				v5 = v5 + i32(-1)
				goto l2
			}
		l1:
			t8 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			store32(m.memory[int64(uint32(v7))+8:], uint32(t8))
			t9 := int64(load64(m.memory[uint32(v1):]))
			store64(m.memory[uint32(v7):], uint64(t9))
		}
	l2:
		t10 := int64(load64(m.memory[int64(uint32(v3))+8:]))
		store64(m.memory[uint32(v0):], uint64(t10))
		store32(m.memory[int64(uint32(v0))+8:], uint32(v5))
		m.g0 = v3 + i32(32)
		return
	}
}
func (m *Module) fn190(v0, v1, v2, v3, v4, v5 int32) {
	if uint32(v4) < uint32(v3) {
		goto l0
	}
	if uint32(v4) <= uint32(v2) {
		goto l1
	}
l0:
	m.fn151(v3, v4, v2, v5)
	panic("unreachable")
l1:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v4-v3))
	store32(m.memory[uint32(v0):], uint32(v1+v3*i32(12)))
}
func (m *Module) fn191(v0, v1, v2, v3 int32) int32 {
	var v4 int32
	v4 = i32(0)
	{
		if v1 != v3 {
			goto l0
		}
		t0 := m.fn235(v0, v2, v1)
		v4 = t0
	}
l0:
	return v4
}
func (m *Module) fn192(v0, v1, v2, v3, v4 int32) {
	if uint32(v3) < uint32(v1) {
		m.fn151(v1, v3, v3, v4)
		panic("unreachable")
	}
	store32(m.memory[int64(uint32(v0))+4:], uint32(v3-v1))
	store32(m.memory[uint32(v0):], uint32(v2+v1*i32(12)))
}
func (m *Module) fn193(v0, v1, v2 int32) {
	var v3, v4, v5 int32
	t0 := m.g0
	v3 = t0 - i32(16)
	m.g0 = v3
	m.fn60(v0, v2)
	t1 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	v4 = t1
	{
		if v2 == 0 {
			goto l0
		}
		v5 = v4 + v2
		v1 = v1 + i32(8)
		t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		v4 = t2 + v4*i32(12)
	l1:
		{
			t3 := int32(load32(m.memory[uint32(v1+i32(-4)):]))
			t4 := int32(load32(m.memory[uint32(v1):]))
			m.fn31(v3+i32(4), t3, t4)
			t5 := int32(load32(m.memory[int64(uint32(v3))+12:]))
			store32(m.memory[int64(uint32(v4))+8:], uint32(t5))
			t6 := int64(load64(m.memory[int64(uint32(v3))+4:]))
			store64(m.memory[uint32(v4):], uint64(t6))
			v1 = v1 + i32(12)
			v4 = v4 + i32(12)
			v2 = v2 + i32(-1)
			if v2 != 0 {
				goto l1
			}
		}
		v4 = v5
	}
l0:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v4))
	m.g0 = v3 + i32(16)
}
func (m *Module) fn194(v0, v1, v2, v3, v4 int32) {
	var v5, v6 int32
	t0 := int32(load32(m.memory[uint32(v1):]))
	v5 = t0
	{
		t1 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		v6 = t1
		if uint32(v6) >= uint32(v3) {
			goto l0
		}
		v6 = v6 + i32(1)
		t2 := int32(m.memory[int64(uint32(v1))+8])
		t4 := v6
		p3 := v5
		if t2 != 0 {
			p3 = v6
		}
		v5 = p3
		if uint32(t4) >= uint32(v5) {
			goto l1
		}
	}
l0:
	m.fn151(v5, v6, v3, v4)
	panic("unreachable")
l1:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v6-v5))
	store32(m.memory[uint32(v0):], uint32(v2+v5*i32(12)))
}
func (m *Module) fn195(v0, v1, v2, v3 int32) int32 {
	t0 := m.fn191(v0, v1, v2, v3)
	return t0 ^ i32(1)
}
func (m *Module) fn196(v0, v1, v2, v3 int32) {
	var v4, v5, v6, v7, v8, v9 int32
	var v10, v11 int64
	t0 := m.g0
	v4 = t0 - i32(176)
	m.g0 = v4
	m.fn198(v4+i32(100), v2, v3, v1)
	{
		{
			t1 := int32(load32(m.memory[int64(uint32(v4))+100:]))
			v5 = t1
			if v5 != i32(-2) {
				goto l0
			}
			t2 := int64(load64(m.memory[int64(uint32(v4))+104:]))
			store64(m.memory[int64(uint32(v0))+4:], uint64(t2))
			store32(m.memory[uint32(v0):], uint32(i32(-0x7ffffff4)))
			goto l1
		}
	l0:
		t3 := int32(load32(m.memory[int64(uint32(v4))+108:]))
		v2 = t3
		t4 := int32(load32(m.memory[int64(uint32(v4))+104:]))
		v3 = t4
		store32(m.memory[int64(uint32(v4))+132:], uint32(i32(-1)))
		store32(m.memory[int64(uint32(v4))+152:], uint32(v3+v2))
		store32(m.memory[int64(uint32(v4))+148:], uint32(v3))
		store32(m.memory[int64(uint32(v4))+144:], uint32(v3))
		store16(m.memory[int64(uint32(v4))+156:], uint16(i32(15142)))
		v6 = i32(0)
		{
			{
			l4:
				m.fn1581(v4+i32(64), v4+i32(144))
				{
					{
						t5 := int32(load32(m.memory[int64(uint32(v4))+64:]))
						if t5 != i32(1) {
							t32 := int32(load32(m.memory[int64(uint32(v4))+132:]))
							if t32 == i32(-1) {
								goto l15
							}
							t33 := int32(load32(m.memory[int64(uint32(v4))+140:]))
							store32(m.memory[int64(uint32(v4))+168:], uint32(t33))
							t34 := int64(load64(m.memory[int64(uint32(v4))+132:]))
							store64(m.memory[int64(uint32(v4))+160:], uint64(t34))
							m.fn1584(v4+i32(56), v3, v2, v6)
							{
								t35 := int32(load32(m.memory[int64(uint32(v4))+56:]))
								v1 = t35
								if v1 == 0 {
									goto l16
								}
								t36 := int32(load32(m.memory[int64(uint32(v4))+60:]))
								m.fn1583(v4+i32(160), v1, t36)
							}
						l16:
							t37 := int32(load32(m.memory[int64(uint32(v4))+168:]))
							store32(m.memory[int64(uint32(v4))+112:], uint32(t37))
							t38 := int64(load64(m.memory[int64(uint32(v4))+160:]))
							store64(m.memory[int64(uint32(v4))+104:], uint64(t38))
							store32(m.memory[int64(uint32(v4))+100:], uint32(i32(-1)))
							goto l17
						}
						t6 := int32(load32(m.memory[int64(uint32(v4))+68:]))
						v1 = t6
						if uint32(v1) >= uint32(v2) {
							m.fn158(v1, v2, i32(1281616))
							panic("unreachable")
						}
						t7 := int32(m.memory[uint32(v3+v1)])
						if t7 != i32(38) {
							goto l4
						}
						m.fn1581(v4+i32(48), v4+i32(144))
						{
							t8 := int32(load32(m.memory[int64(uint32(v4))+48:]))
							if t8&i32(1) == 0 {
								goto l5
							}
							t9 := int32(load32(m.memory[int64(uint32(v4))+52:]))
							v7 = t9
							if uint32(v7) >= uint32(v2) {
								m.fn158(v7, v2, i32(1281472))
								panic("unreachable")
							}
							t10 := int32(m.memory[uint32(v3+v7)])
							if t10 == i32(59) {
								{
									t11 := int32(load32(m.memory[int64(uint32(v4))+132:]))
									if t11 != i32(-1) {
										goto l9
									}
									m.fn1064(v4+i32(40), v2)
									t12 := int32(load32(m.memory[int64(uint32(v4))+40:]))
									v8 = t12
									t13 := int32(load32(m.memory[int64(uint32(v4))+44:]))
									v9 = t13
									t14 := int32(load32(m.memory[int64(uint32(v4))+136:]))
									m.fn277(i32(-1), t14)
									store32(m.memory[int64(uint32(v4))+140:], uint32(i32(0)))
									store32(m.memory[int64(uint32(v4))+136:], uint32(v9))
									store32(m.memory[int64(uint32(v4))+132:], uint32(v8))
									if v8 == i32(-1) {
										m.fn633(i32(1281488), i32(11), i32(1281500))
										panic("unreachable")
									}
								}
							l9:
								m.fn1582(v4+i32(32), v3, v2, v6, v1, i32(1281516))
								t15 := int32(load32(m.memory[int64(uint32(v4))+32:]))
								t16 := int32(load32(m.memory[int64(uint32(v4))+36:]))
								m.fn1583(v4+i32(132), t15, t16)
								t17 := v4 + i32(24)
								t18 := v3
								t19 := v2
								v9 = v1 + i32(1)
								m.fn1582(t17, t18, t19, v9, v7, i32(1281532))
								t20 := int32(load32(m.memory[int64(uint32(v4))+24:]))
								t21 := v4 + i32(16)
								v6 = t20
								t22 := int32(load32(m.memory[int64(uint32(v4))+28:]))
								t23 := v6
								v8 = t22
								m.fn13(t21, t23, v8, i32(35))
								{
									t24 := int32(load32(m.memory[int64(uint32(v4))+16:]))
									v1 = t24
									if v1 == 0 {
										t31 := m.fn275(v6, v8)
										v1 = t31
										if v1 == 0 {
											m.fn884(v4+i32(100), v6, v8)
											store32(m.memory[int64(uint32(v4))+116:], uint32(v7))
											store32(m.memory[int64(uint32(v4))+112:], uint32(v9))
											t39 := int32(load32(m.memory[int64(uint32(v4))+100:]))
											v1 = t39
											t40 := int32(load32(m.memory[int64(uint32(v4))+132:]))
											t41 := int32(load32(m.memory[int64(uint32(v4))+136:]))
											m.fn277(t40, t41)
											if v1 != i32(-1) {
												goto l18
											}
											goto l17
										}
										m.fn1583(v4+i32(132), v1, i32(1))
										goto l13
									}
									t25 := int32(load32(m.memory[int64(uint32(v4))+20:]))
									m.fn276(v4+i32(160), v1, t25)
									{
										t26 := int32(m.memory[int64(uint32(v4))+160])
										if t26 == i32(255) {
											t28 := int32(load32(m.memory[int64(uint32(v4))+164:]))
											v1 = t28
											store32(m.memory[int64(uint32(v4))+160:], uint32(i32(0)))
											m.fn522(v4+i32(8), v1, v4+i32(160))
											t29 := int32(load32(m.memory[int64(uint32(v4))+8:]))
											t30 := int32(load32(m.memory[int64(uint32(v4))+12:]))
											m.fn1583(v4+i32(132), t29, t30)
											goto l13
										}
										t27 := int64(load64(m.memory[int64(uint32(v4))+160:]))
										store64(m.memory[int64(uint32(v4))+104:], uint64(t27))
										store32(m.memory[int64(uint32(v4))+100:], uint32(i32(-0x7fffffff)))
										goto l8
									}
								}
							}
						}
					l5:
						store32(m.memory[int64(uint32(v4))+108:], uint32(v2))
						store32(m.memory[int64(uint32(v4))+104:], uint32(v1))
						store32(m.memory[int64(uint32(v4))+100:], uint32(i32(-0x80000000)))
						goto l8
					}
				l15:
					store32(m.memory[int64(uint32(v4))+112:], uint32(v2))
					store32(m.memory[int64(uint32(v4))+108:], uint32(v3))
					store64(m.memory[int64(uint32(v4))+100:], uint64(i64(-1)))
				l17:
					t42 := int64(load64(m.memory[int64(uint32(v4))+104:]))
					store64(m.memory[int64(uint32(v4))+120:], uint64(t42))
					t43 := int32(load32(m.memory[int64(uint32(v4))+112:]))
					store32(m.memory[int64(uint32(v4))+128:], uint32(t43))
					m.fn490(v0+i32(4), v4+i32(120))
					store32(m.memory[uint32(v0):], uint32(i32(-1)))
					goto l19
				}
			l13:
				v6 = v7 + i32(1)
				goto l4
			l8:
				t44 := int32(load32(m.memory[int64(uint32(v4))+132:]))
				t45 := int32(load32(m.memory[int64(uint32(v4))+136:]))
				m.fn277(t44, t45)
			}
		l18:
			t46 := int32(load32(m.memory[int64(uint32(v4))+108:]))
			t47 := v4
			v1 = t46
			store32(m.memory[int64(uint32(t47))+96:], uint32(v1))
			t48 := int64(load64(m.memory[int64(uint32(v4))+100:]))
			t49 := v4
			v10 = t48
			store64(m.memory[int64(uint32(t49))+72:], uint64(v10))
			store32(m.memory[int64(uint32(v4))+80:], uint32(v1))
			t50 := int64(load64(m.memory[int64(uint32(v4))+112:]))
			v11 = t50
			store32(m.memory[int64(uint32(v0))+12:], uint32(v1))
			store64(m.memory[int64(uint32(v0))+4:], uint64(v10))
			store64(m.memory[int64(uint32(v0))+16:], uint64(v11))
			store32(m.memory[uint32(v0):], uint32(i32(-0x7ffffff3)))
		}
	l19:
		m.fn1390(v5, v3)
	}
l1:
	m.g0 = v4 + i32(176)
}
func (m *Module) fn197(v0, v1, v2 int32) {
	var v3, v4, v5 int32
	var v6 int64
	t0 := m.g0
	v3 = t0 - i32(16)
	m.g0 = v3
	{
		switch v2 {
		case 0:
			m.memory[int64(uint32(v0))+1] = byte(i32(0))
			v2 = i32(1)
			goto l3
		case 1:
			t1 := int32(m.memory[uint32(v1)])
			v4 = t1
			switch v4 + i32(-43) {
			case 0, 2:
				goto l4
			default:
				goto l5
			}
		default:
			t2 := int32(m.memory[uint32(v1)])
			v4 = t2
		}
	l5:
		t3 := v1
		var p4 int32
		if v4&i32(255) == i32(43) {
			p4 = 1
		}
		v4 = p4
		v1 = t3 + v4
		v2 = v2 - v4
		if uint32(v2) < uint32(i32(9)) {
			v4 = i32(0)
		l10:
			{
				if v2 == 0 {
					goto l7
				}
				t8 := int32(m.memory[uint32(v1)])
				m.fn199(v3, t8, i32(10))
				t9 := int32(load32(m.memory[uint32(v3):]))
				if t9 != i32(1) {
					goto l4
				}
				v1 = v1 + i32(1)
				v2 = v2 + i32(-1)
				t10 := int32(load32(m.memory[int64(uint32(v3))+4:]))
				v4 = t10 + v4*i32(10)
				goto l10
			}
		}
		v4 = i32(0)
	l9:
		{
			if v2 == 0 {
				goto l7
			}
			t5 := int32(m.memory[uint32(v1)])
			m.fn199(v3+i32(8), t5, i32(10))
			t6 := int32(load32(m.memory[int64(uint32(v3))+8:]))
			v5 = t6
			v6 = int64(uint32(v4)) * i64(10)
			if int32(int64(uint64(v6)>>32)) != 0 {
				v2 = i32(1)
				if v5&i32(1) == 0 {
					goto l4
				}
				m.memory[int64(uint32(v0))+1] = byte(i32(2))
				goto l3
			}
			if v5&i32(1) == 0 {
				goto l4
			}
			v1 = v1 + i32(1)
			v2 = v2 + i32(-1)
			t7 := int32(load32(m.memory[int64(uint32(v3))+12:]))
			v5 = t7
			v4 = v5 + int32(v6)
			if uint32(v4) >= uint32(v5) {
				goto l9
			}
		}
		m.memory[int64(uint32(v0))+1] = byte(i32(2))
		v2 = i32(1)
		goto l3
	}
l4:
	v2 = i32(1)
	m.memory[int64(uint32(v0))+1] = byte(i32(1))
	goto l3
l7:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v4))
	v2 = i32(0)
l3:
	m.memory[uint32(v0)] = byte(v2)
	m.g0 = v3 + i32(16)
}
func (m *Module) fn198(v0, v1, v2, v3 int32) {
	var v4, v5, v6, v7 int32
	t0 := m.g0
	v4 = t0 - i32(112)
	m.g0 = v4
	{
		if v3 == i32(1148960) {
			v7 = i32(-1)
			t20 := m.fn1692(v1, v2)
			if t20 == v2 {
				goto l10
			}
			goto l11
		}
		{
			if v3 == i32(1153092) {
				goto l1
			}
			if v3 == i32(1153580) {
				goto l1
			}
			if v3 != i32(1153064) {
				{
					{
						if v3 == i32(1153400) {
							goto l5
						}
						t9 := m.fn1693(v1, v2)
						v6 = t9
						goto l6
					}
				l5:
					t10 := m.fn1700(v1, v2)
					v6 = t10
				}
			l6:
				{
					if v2 == v6 {
						v7 = i32(-1)
						goto l10
					}
					m.fn1695(v4+i32(48), v3)
					m.fn1696(v4+i32(16), v4+i32(48), v2-v6)
					t11 := int32(load32(m.memory[int64(uint32(v4))+16:]))
					if t11 != i32(1) {
						goto l8
					}
					t12 := int32(load32(m.memory[int64(uint32(v4))+20:]))
					v7 = t12
					v5 = v7 + v6
					if uint32(v5) < uint32(v7) {
						goto l8
					}
					m.fn1699(v4+i32(36), v5)
					t13 := int32(load32(m.memory[int64(uint32(v4))+40:]))
					v5 = t13
					if v6 == 0 {
						goto l9
					}
					memory_copy(m.memory, uint32(v5), uint32(v1), uint32(v6))
				l9:
					m.fn148(v4+i32(8), v6, v1, v2, i32(1154928))
					t14 := int64(load64(m.memory[int64(uint32(v4))+48:]))
					store64(m.memory[int64(uint32(v4))+80:], uint64(t14))
					t15 := int64(load64(m.memory[int64(uint32(v4))+56:]))
					store64(m.memory[int64(uint32(v4))+88:], uint64(t15))
					t16 := int64(load64(m.memory[int64(uint32(v4))+64:]))
					store64(m.memory[int64(uint32(v4))+96:], uint64(t16))
					t17 := int32(load32(m.memory[int64(uint32(v4))+72:]))
					store32(m.memory[int64(uint32(v4))+104:], uint32(t17))
					t18 := int32(load32(m.memory[int64(uint32(v4))+12:]))
					v2 = t18
					t19 := int32(load32(m.memory[int64(uint32(v4))+8:]))
					v1 = t19
					goto l4
				}
			}
		l1:
			m.fn1695(v4+i32(48), v3)
			m.fn1696(v4+i32(24), v4+i32(48), v2)
			t1 := int32(load32(m.memory[int64(uint32(v4))+24:]))
			if t1&i32(1) == 0 {
				m.fn153(i32(1154896))
				panic("unreachable")
			}
			t2 := int32(load32(m.memory[int64(uint32(v4))+28:]))
			m.fn1699(v4+i32(36), t2)
			t3 := int32(load32(m.memory[int64(uint32(v4))+72:]))
			store32(m.memory[int64(uint32(v4))+104:], uint32(t3))
			t4 := int64(load64(m.memory[int64(uint32(v4))+64:]))
			store64(m.memory[int64(uint32(v4))+96:], uint64(t4))
			t5 := int64(load64(m.memory[int64(uint32(v4))+56:]))
			store64(m.memory[int64(uint32(v4))+88:], uint64(t5))
			t6 := int64(load64(m.memory[int64(uint32(v4))+48:]))
			store64(m.memory[int64(uint32(v4))+80:], uint64(t6))
			t7 := int32(load32(m.memory[int64(uint32(v4))+40:]))
			v5 = t7
			t8 := int32(load32(m.memory[int64(uint32(v4))+44:]))
			v6 = t8
			goto l4
		}
	l8:
		m.fn153(i32(1154912))
		panic("unreachable")
	l4:
		t21 := int32(load32(m.memory[int64(uint32(v4))+104:]))
		store32(m.memory[int64(uint32(v4))+72:], uint32(t21))
		t22 := int64(load64(m.memory[int64(uint32(v4))+96:]))
		store64(m.memory[int64(uint32(v4))+64:], uint64(t22))
		t23 := int64(load64(m.memory[int64(uint32(v4))+88:]))
		store64(m.memory[int64(uint32(v4))+56:], uint64(t23))
		t24 := int64(load64(m.memory[int64(uint32(v4))+80:]))
		store64(m.memory[int64(uint32(v4))+48:], uint64(t24))
		t25 := int32(load32(m.memory[int64(uint32(v4))+36:]))
		t26 := v4
		t27 := v6
		t28 := v5
		v7 = t25
		m.fn212(t26, t27, t28, v7, i32(1155312))
		t29 := int32(load32(m.memory[uint32(v4):]))
		t30 := int32(load32(m.memory[int64(uint32(v4))+4:]))
		m.fn1701(v4+i32(80), v4+i32(48), v1, v2, t29, t30)
		{
			t31 := int32(m.memory[int64(uint32(v4))+84])
			switch t31 {
			case 1:
				m.fn256(i32(1286542), i32(40), i32(1154944))
				panic("unreachable")
			case 2:
				m.fn1091(v7, v5)
				goto l11
			default:
				t32 := int32(load32(m.memory[int64(uint32(v4))+88:]))
				v2 = t32 + v6
				v1 = v5
			}
		}
	}
l10:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v2))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
	store32(m.memory[uint32(v0):], uint32(v7))
	goto l15
l11:
	m.memory[int64(uint32(v0))+8] = byte(i32(2))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
	store32(m.memory[uint32(v0):], uint32(i32(-2)))
l15:
	m.g0 = v4 + i32(112)
}
func (m *Module) fn199(v0, v1, v2 int32) {
	var v3 int32
	t0 := v0
	t1 := (v1+i32(-65))&i32(-33) + i32(10)
	v3 = v1 + i32(-48)
	p2 := v3
	if uint32(v2) > uint32(i32(10)) {
		p2 = t1
	}
	p3 := v3
	if uint32(v1) > uint32(i32(57)) {
		p3 = p2
	}
	v1 = p3
	store32(m.memory[int64(uint32(t0))+4:], uint32(v1))
	t4 := v0
	var p5 int32
	if uint32(v1) < uint32(v2) {
		p5 = 1
	}
	store32(m.memory[uint32(t4):], uint32(p5))
}
func (m *Module) fn200(v0 int32) {
	t0 := int32(load32(m.memory[uint32(v0):]))
	if uint32(t0) > uint32(i32(9)) {
		return
	}
	t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t2 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	m.fn134(t1, t2)
}
func (m *Module) fn201(v0, v1 int32) {
	t0 := int32(load32(m.memory[int64(uint32(v1))+12:]))
	m.fn571(v0, t0, v1)
}
func (m *Module) fn202(v0, v1, v2 int32) {
	var v3 int32
	var v4 int64
	var v5, v6, v7, v8, v9, v10 int32
	t0 := m.g0
	v3 = t0 - i32(64)
	m.g0 = v3
	m.fn274(v3+i32(48), v1)
	t1 := int64(load64(m.memory[int64(uint32(v3))+52:]))
	v4 = t1
	{
		t2 := int32(load32(m.memory[int64(uint32(v3))+48:]))
		v5 = t2
		if v5 != i32(-2) {
			goto l0
		}
		store64(m.memory[int64(uint32(v0))+4:], uint64(v4))
		store32(m.memory[uint32(v0):], uint32(i32(-0x7ffffff4)))
		goto l1
	}
l0:
	store32(m.memory[int64(uint32(v3))+20:], uint32(v5))
	store64(m.memory[int64(uint32(v3))+24:], uint64(v4))
	{
		v6 = int32(v4)
		t3 := m.fn275(v6, int32(int64(uint64(v4)>>32)))
		v7 = t3
		if v7 == 0 {
			m.fn274(v3+i32(48), v1)
			t4 := int64(load64(m.memory[int64(uint32(v3))+52:]))
			v4 = t4
			{
				{
					{
						t5 := int32(load32(m.memory[int64(uint32(v3))+48:]))
						v8 = t5
						if v8 != i32(-2) {
							goto l4
						}
						store64(m.memory[int64(uint32(v3))+32:], uint64(v4))
						v1 = int32(v4)
						v7 = i32(-0x7ffffff4)
						goto l5
					}
				l4:
					t6 := int32(load32(m.memory[int64(uint32(v3))+52:]))
					v9 = t6
					m.fn13(v3+i32(8), int32(v4), int32(int64(uint64(v4)>>32)), i32(35))
					t7 := int32(load32(m.memory[int64(uint32(v3))+8:]))
					v1 = t7
					if v1 == 0 {
						m.fn277(v8, v9)
						store32(m.memory[int64(uint32(v3))+52:], uint32(i32(49)))
						store32(m.memory[int64(uint32(v3))+48:], uint32(v3+i32(20)))
						m.fn73(v0+i32(4), i32(1068659), v3+i32(48))
						store64(m.memory[int64(uint32(v0))+16:], uint64(i64(0)))
						store32(m.memory[uint32(v0):], uint32(i32(-0x7ffffff3)))
						t13 := int32(load32(m.memory[int64(uint32(v3))+20:]))
						t14 := int32(load32(m.memory[int64(uint32(v3))+24:]))
						m.fn134(t13, t14)
						goto l1
					}
					t8 := int32(load32(m.memory[int64(uint32(v3))+12:]))
					m.fn276(v3+i32(48), v1, t8)
					{
						{
							t9 := int32(m.memory[int64(uint32(v3))+48])
							v10 = t9
							if v10 == i32(255) {
								goto l7
							}
							t10 := int64(load64(m.memory[int64(uint32(v3))+48:]))
							store64(m.memory[int64(uint32(v3))+36:], uint64(t10))
							v7 = i32(-0x7ffffff3)
							v1 = i32(-0x7fffffff)
							goto l8
						}
					l7:
						t11 := int32(load32(m.memory[int64(uint32(v3))+52:]))
						v1 = t11
						v7 = i32(-1)
					}
				l8:
					m.fn277(v8, v9)
					if v10 == i32(255) {
						m.fn74(v2, v1)
						store32(m.memory[uint32(v0):], uint32(i32(-1)))
						goto l3
					}
				}
			l5:
				t12 := int64(load64(m.memory[int64(uint32(v3))+36:]))
				store64(m.memory[int64(uint32(v0))+8:], uint64(t12))
				store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
				store32(m.memory[uint32(v0):], uint32(v7))
				goto l3
			}
		}
		m.fn75(v2, v7, i32(1))
		store32(m.memory[uint32(v0):], uint32(i32(-1)))
		goto l3
	}
l3:
	m.fn134(v5, v6)
l1:
	m.g0 = v3 + i32(64)
}
func (m *Module) fn203(v0, v1 int32) {
	var v2 int32
	var v3 int64
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	t1 := int64(load64(m.memory[uint32(v1):]))
	v3 = t1
l1:
	{
		m.fn280(v2+i32(8), v0)
		t2 := int32(m.memory[int64(uint32(v2))+8])
		if t2 == 0 {
			goto l0
		}
		t3 := int32(m.memory[int64(uint32(v2))+9])
		t4 := v1
		v3 = v3*i64(10) + int64(uint32(t3))&i64(255)
		store64(m.memory[uint32(t4):], uint64(v3))
		goto l1
	}
l0:
	m.g0 = v2 + i32(16)
}
func fn204(v0 int64) int64 {
	v0 = v0 + i64(-3472328296227680304)
	v0 = v0*i64(10) + int64(uint64(v0)>>8)
	return int64(uint64(int64(uint64(v0)>>16)&i64(0xff000000ff)*i64(0x271000000001)+v0&i64(0xff000000ff)*i64(0xf424000000064)) >> 32)
}
func (m *Module) fn205(v0, v1 int32) {
	var v2 int32
	var v3 int64
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	t1 := int64(load64(m.memory[uint32(v1):]))
	v3 = t1
l1:
	{
		m.fn280(v2+i32(8), v0)
		t2 := int32(m.memory[int64(uint32(v2))+8])
		if t2 == 0 {
			goto l0
		}
		if v3 > i64(0xffff) {
			goto l1
		}
		t3 := int32(m.memory[int64(uint32(v2))+9])
		t4 := v1
		v3 = v3*i64(10) + int64(uint32(t3))&i64(255)
		store64(m.memory[uint32(t4):], uint64(v3))
		goto l1
	}
l0:
	m.g0 = v2 + i32(16)
}
func (m *Module) fn206(v0, v1, v2 int32) int32 {
	var v3 int32
	v3 = i32(0)
	{
		if uint32(v1) < uint32(i32(3)) {
			goto l0
		}
		v1 = i32(0)
		v3 = i32(0)
	l2:
		{
			if v3 == i32(3) {
				goto l1
			}
			t0 := int32(m.memory[uint32(v2+v3)])
			t1 := int32(m.memory[uint32(v0+v3)])
			v1 = t0 ^ t1 | v1
			v3 = v3 + i32(1)
			goto l2
		}
	l1:
		;
		var p2 int32
		if v1&i32(223) == 0 {
			p2 = 1
		}
		v3 = p2
	}
l0:
	return v3
}
func (m *Module) fn207(v0, v1, v2, v3 int32) {
	var v4 int32
	t0 := m.g0
	v4 = t0 - i32(16)
	m.g0 = v4
	m.fn148(v4+i32(8), v3, v1, v2, i32(1100696))
	t1 := int32(load32(m.memory[int64(uint32(v4))+12:]))
	v2 = t1
	t2 := int32(load32(m.memory[int64(uint32(v4))+8:]))
	store32(m.memory[uint32(v0):], uint32(t2))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v2))
	m.g0 = v4 + i32(16)
}
func (m *Module) fn208(v0, v1 int32) int32 {
	var v2 int32
	v2 = i32(3)
	{
		if uint32(v1) < uint32(i32(8)) {
			goto l0
		}
		v0 = v0 + i32(3)
		v1 = i32(0)
		v2 = i32(0)
	l2:
		{
			if v2 == i32(5) {
				goto l1
			}
			t0 := int32(m.memory[uint32(v2+i32(1280576))])
			t1 := int32(m.memory[uint32(v0+v2)])
			v1 = t0 ^ t1 | v1
			v2 = v2 + i32(1)
			goto l2
		}
	l1:
		p2 := i32(8)
		if v1&i32(223) != 0 {
			p2 = i32(3)
		}
		v2 = p2
	}
l0:
	return v2
}
func (m *Module) fn209(v0, v1 int32) {
	var v2 int32
	var v3 int64
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	t1 := int64(load64(m.memory[uint32(v1):]))
	v3 = t1
l2:
	{
		if uint64(v3) > uint64(i64(999999999999999999)) {
			goto l0
		}
		m.fn280(v2+i32(8), v0)
		t2 := int32(m.memory[int64(uint32(v2))+8])
		if t2 != 0 {
			t3 := int32(m.memory[int64(uint32(v2))+9])
			t4 := v1
			v3 = v3*i64(10) + int64(uint32(t3))&i64(255)
			store64(m.memory[uint32(t4):], uint64(v3))
			goto l2
		}
	}
l0:
	m.g0 = v2 + i32(16)
}
func (m *Module) fn210(v0 int32, v1, v2 int64) {
	var v3, v4, v5 int32
	var v6, v7, v8, v9, v10 int64
	var v11 int32
	t0 := m.g0
	v3 = t0 - i32(32)
	m.g0 = v3
	store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
	store64(m.memory[uint32(v0):], uint64(i64(0)))
	if v1 < i64(-342) {
		goto l0
	}
	if v2 == 0 {
		goto l0
	}
	v4 = i32(2047)
	{
		if v1 > i64(308) {
			goto l1
		}
		t1 := v3 + i32(16)
		v5 = int32(v1)
		v4 = v5 << 4
		t2 := int64(load64(m.memory[uint32(v4+i32(1123784)):]))
		t3 := v2
		v6 = int64(bits.LeadingZeros64(uint64(v2)))
		v7 = i64_shl(t3, v6)
		m.fn1853(t1, t2, i64(0), v7, i64(0))
		t4 := int64(load64(m.memory[int64(uint32(v3))+16:]))
		v2 = t4
		{
			t5 := int64(load64(m.memory[int64(uint32(v3))+24:]))
			v8 = t5
			if v8&i64(511) != i64(511) {
				goto l2
			}
			t6 := int64(load64(m.memory[uint32(v4+i32(1118312)+i32(5480)):]))
			m.fn1853(v3, t6, i64(0), v7, i64(0))
			t7 := int64(load64(m.memory[int64(uint32(v3))+8:]))
			v7 = t7
			v2 = v7 + v2
			var p8 int32
			if uint64(v2) < uint64(v7) {
				p8 = 1
			}
			v8 = int64(uint32(p8)) + v8
		}
	l2:
		if v2 != i64(-1) {
			goto l3
		}
		v4 = i32(-1)
		if uint64(v1+i64(27)) > uint64(i64(82)) {
			goto l1
		}
	l3:
		t9 := v8
		v9 = int64(uint64(v8) >> 63)
		v10 = v9 + i64(9)
		v7 = i64_shr_u(t9, v10)
		{
			{
				v5 = v5*i32(217706)>>16 - int32(v6) + int32(v9) + i32(63)
				if v5 < i32(-1022) {
					goto l4
				}
				v4 = i32(2047)
				p10 := v7
				if i64_shl(v7, v10) == v8 {
					p10 = v7 & i64(0xfffffffffffffc)
				}
				p11 := v7
				if v7&i64(3) == i64(1) {
					p11 = p10
				}
				p12 := v7
				if uint64(v2) < uint64(i64(2)) {
					p12 = p11
				}
				p13 := v7
				if uint64(v1+i64(4)) < uint64(i64(28)) {
					p13 = p12
				}
				v1 = p13
				v1 = v1&i64(1) + v1
				var p14 int32
				if uint64(v1) > uint64(i64(0x3fffffffffffff)) {
					p14 = 1
				}
				v11 = p14
				p15 := i32(1023)
				if v11 != 0 {
					p15 = i32(1024)
				}
				v5 = p15 + v5
				if uint32(v5) > uint32(i32(2046)) {
					goto l1
				}
				p16 := int64(uint64(v1)>>1) & i64(0x7fefffffffffffff)
				if v11 != 0 {
					p16 = i64(0)
				}
				v1 = p16
				v4 = v5
				goto l5
			}
		l4:
			if uint32(v5) < uint32(i32(-1085)) {
				goto l0
			}
			v1 = i64_shr_u(v7, int64(uint32(i32(2)-v5)))
			v1 = v1&i64(1) + v1
			var p17 int32
			if uint64(v1) > uint64(i64(0x1fffffffffffff)) {
				p17 = 1
			}
			v4 = p17
			v1 = int64(uint64(v1) >> 1)
		}
	l5:
		store64(m.memory[uint32(v0):], uint64(v1))
	}
l1:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v4))
l0:
	m.g0 = v3 + i32(32)
}
func (m *Module) fn211(v0, v1, v2 int32) {
	var v3 int32
	t0 := m.g0
	v3 = t0 - i32(16)
	m.g0 = v3
l2:
	{
		if v2 == 0 {
			goto l0
		}
		t1 := int32(m.memory[uint32(v1)])
		if t1 == i32(48) {
			m.fn207(v3+i32(8), v1, v2, i32(1))
			t2 := int32(load32(m.memory[int64(uint32(v3))+12:]))
			v2 = t2
			t3 := int32(load32(m.memory[int64(uint32(v3))+8:]))
			v1 = t3
			goto l2
		}
	}
l0:
	store32(m.memory[uint32(v0):], uint32(v1))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v2))
	m.g0 = v3 + i32(16)
}
func (m *Module) fn212(v0, v1, v2, v3, v4 int32) {
	if uint32(v3) < uint32(v1) {
		m.fn151(v1, v3, v3, v4)
		panic("unreachable")
	}
	store32(m.memory[int64(uint32(v0))+4:], uint32(v3-v1))
	store32(m.memory[uint32(v0):], uint32(v2+v1))
}
func (m *Module) fn213(v0, v1 int32) {
	var v2, v3, v4, v5, v6 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	t1 := int32(load32(m.memory[uint32(v1):]))
	v3 = t1
	t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	v4 = t2
	t3 := int32(load32(m.memory[uint32(v0):]))
	v5 = t3
l2:
	{
		if v4 == 0 {
			goto l0
		}
		t4 := int32(m.memory[uint32(v5)])
		v6 = (t4 + i32(-48)) & i32(255)
		if uint32(v6) >= uint32(i32(10)) {
			goto l0
		}
		{
			if v3 > i32(0xffff) {
				goto l1
			}
			t5 := v1
			v3 = v3*i32(10) + v6
			store32(m.memory[uint32(t5):], uint32(v3))
		}
	l1:
		m.fn207(v2+i32(8), v5, v4, i32(1))
		t6 := int32(load32(m.memory[int64(uint32(v2))+8:]))
		v5 = t6
		t7 := int32(load32(m.memory[int64(uint32(v2))+12:]))
		t8 := v0
		v4 = t7
		store32(m.memory[int64(uint32(t8))+4:], uint32(v4))
		store32(m.memory[uint32(v0):], uint32(v5))
		goto l2
	}
l0:
	m.g0 = v2 + i32(16)
}
func (m *Module) fn214(v0, v1 int32) {
	var v2 int32
	var v3 int64
	var v4 int32
	var v5 int64
	var v6, v7, v8 int32
	var v9, v10 int64
	v2 = v0 + i32(768)
	v3 = int64(uint32(v1 & i32(63)))
	t0 := int32(load32(m.memory[int64(uint32(v0))+768:]))
	v4 = t0
	v1 = i32(0)
	v5 = i64(0)
	{
	l5:
		if i64_shr_u(v5, v3) != i64(0) {
			goto l0
		}
		if v4 != v1 {
			goto l1
		}
		if v5 == 0 {
			return
		}
		v1 = v4
	l3:
		if i64_shr_u(v5, v3) != i64(0) {
			goto l0
		}
		v1 = v1 + i32(1)
		v5 = v5 * i64(10)
		goto l3
	l1:
		{
			if v1 == i32(768) {
				goto l4
			}
			t1 := int64(m.memory[uint32(v0+v1)])
			v5 = v5*i64(10) + t1
			v1 = v1 + i32(1)
			goto l5
		}
	l4:
		m.fn158(i32(768), i32(768), i32(1087944))
		panic("unreachable")
	l0:
		t2 := int32(load32(m.memory[int64(uint32(v0))+772:]))
		t3 := v0
		v6 = t2 - v1 + i32(1)
		store32(m.memory[int64(uint32(t3))+772:], uint32(v6))
		{
			if v6 < i32(-2047) {
				goto l6
			}
			v2 = i32(0)
			v6 = i32(768) - v1
			p4 := v6
			if uint32(v6) > uint32(i32(768)) {
				p4 = i32(0)
			}
			v7 = p4
			v8 = v0 + v1
			v9 = i64_shl(i64(-1), v3) ^ i64(-1)
		l12:
			v6 = v1 + v2
			if uint32(v6) < uint32(v4) {
				if v7 == v2 {
					m.fn158(v6, i32(768), i32(1087960))
					panic("unreachable")
				}
				t5 := int64(m.memory[uint32(v8+v2)])
				v10 = t5
				m.memory[uint32(v0+v2)] = byte(i64_shr_u(v5, v3))
				v5 = v10 + v5&v9*i64(10)
				v2 = v2 + i32(1)
				t6 := int32(load32(m.memory[int64(uint32(v0))+768:]))
				v4 = t6
				goto l12
			}
		l10:
			if v5 != i64(0) {
				v10 = v5 & v9 * i64(10)
				v1 = int32(i64_shr_u(v5, v3))
				if uint32(v2) < uint32(i32(768)) {
					m.memory[uint32(v0+v2)] = byte(v1)
					v2 = v2 + i32(1)
					v5 = v10
					goto l10
				}
				v5 = v10
				if v1&i32(255) == 0 {
					goto l10
				}
				m.memory[int64(uint32(v0))+777] = byte(i32(1))
				v5 = v10
				goto l10
			}
			store32(m.memory[int64(uint32(v0))+768:], uint32(v2))
			m.fn281(v0)
			return
		}
	l6:
		store16(m.memory[int64(uint32(v2))+8:], uint16(i32(0)))
		store64(m.memory[uint32(v2):], uint64(i64(0)))
	}
}
func (m *Module) fn215(v0, v1 int32) {
	var v2, v3, v4, v5, v6, v7, v8, v9, v10 int32
	var v11, v12, v13 int64
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	{
		t1 := int32(load32(m.memory[int64(uint32(v0))+768:]))
		v3 = t1
		if v3 == 0 {
			goto l0
		}
		t2 := v2 + i32(8)
		v4 = v1 & i32(63)
		v1 = v4 << 1
		t3 := int32(load16(m.memory[int64(uint32(v1))+1116840:]))
		v5 = t3
		v6 = v5 & i32(2047)
		m.fn148(t2, v6, i32(1116970), i32(1308), i32(1088168))
		t4 := int32(load16(m.memory[int64(uint32(v1))+1116842:]))
		v7 = t4&i32(2047) - v6
		v8 = int32(uint32(v5) >> 11)
		v1 = i32(0)
		t5 := int32(load32(m.memory[int64(uint32(v2))+12:]))
		v9 = t5
		t6 := int32(load32(m.memory[int64(uint32(v2))+8:]))
		v10 = t6
		{
		l4:
			{
				if v7 == v1 {
					goto l1
				}
				if v9 == v1 {
					goto l1
				}
				if v3 == v1 {
					v8 = v8 + i32(-1)
					goto l1
				}
				if v1 == i32(768) {
					m.fn158(i32(768), i32(768), i32(1088184))
					panic("unreachable")
				}
				v5 = v10 + v1
				v6 = v0 + v1
				v1 = v1 + i32(1)
				t7 := int32(m.memory[uint32(v6)])
				v6 = t7
				t8 := int32(m.memory[uint32(v5)])
				t9 := v6
				v5 = t8
				if t9 == v5&i32(255) {
					goto l4
				}
			}
			t10 := v8
			var p11 int32
			if uint32(v6) < uint32(v5&i32(255)) {
				p11 = 1
			}
			v8 = t10 - p11
			goto l1
		}
	l1:
		v1 = v3 + i32(-1)
		v5 = v1 + v8
		v11 = int64(uint32(v4))
		v12 = i64(0)
		var p12 int32
		if uint32(v3) < uint32(i32(769)) {
			p12 = 1
		}
		v6 = p12
	l13:
		if v1 != i32(-1) {
			if v6 == 0 {
				m.fn158(v1, i32(768), i32(1087928))
				panic("unreachable")
			}
			t18 := int64(m.memory[uint32(v0+v1)])
			v12 = i64_shl(t18, v11) + v12
			t19 := int64(uint64(v12) / uint64(i64(10)))
			t20 := v12
			v12 = t19
			v13 = t20 + v12*i64(-10)
			if uint32(v5) < uint32(i32(768)) {
				goto l11
			}
			if v13 == 0 {
				goto l12
			}
			m.memory[int64(uint32(v0))+777] = byte(i32(1))
			goto l12
		l11:
			m.memory[uint32(v0+v5)] = byte(v13)
		l12:
			v1 = v1 + i32(-1)
			v5 = v5 + i32(-1)
			goto l13
		}
		v1 = v8 + i32(-1)
	l9:
		{
			if v12 != i64(0) {
				t17 := int64(uint64(v12) / uint64(i64(10)))
				v13 = t17
				v12 = v13*i64(-10) + v12
				if uint32(v1) < uint32(i32(768)) {
					goto l7
				}
				if v12 == 0 {
					goto l8
				}
				m.memory[int64(uint32(v0))+777] = byte(i32(1))
				goto l8
			l7:
				m.memory[uint32(v0+v1)] = byte(v12)
			l8:
				v1 = v1 + i32(-1)
				v12 = v13
				goto l9
			}
			t13 := int32(load32(m.memory[int64(uint32(v0))+772:]))
			store32(m.memory[int64(uint32(v0))+772:], uint32(t13+v8))
			t14 := int32(load32(m.memory[int64(uint32(v0))+768:]))
			t15 := v0
			v1 = t14 + v8
			p16 := i32(768)
			if uint32(v1) < uint32(i32(768)) {
				p16 = v1
			}
			store32(m.memory[int64(uint32(t15))+768:], uint32(p16))
			m.fn281(v0)
			goto l0
		}
	}
l0:
	m.g0 = v2 + i32(16)
}
func (m *Module) fn216(v0 int32) int64 {
	var v1 int64
	var v2, v3, v4, v5 int32
	v1 = i64(0)
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+768:]))
		v2 = t0
		if v2 == 0 {
			goto l0
		}
		t1 := int32(load32(m.memory[int64(uint32(v0))+772:]))
		v3 = t1
		if v3 < i32(0) {
			goto l0
		}
		v1 = i64(-1)
		if uint32(v3) > uint32(i32(18)) {
			goto l0
		}
		v1 = i64(0)
		v4 = i32(0)
	l3:
		if v3 == v4 {
			goto l1
		}
		v1 = v1 * i64(10)
		{
			if uint32(v4) >= uint32(v2) {
				goto l2
			}
			t2 := int64(m.memory[uint32(v0+v4)])
			v1 = v1 + t2
		}
	l2:
		v4 = v4 + i32(1)
		goto l3
	l1:
		if uint32(v3) >= uint32(v2) {
			goto l0
		}
		v5 = v0 + v3
		t3 := int32(m.memory[uint32(v5)])
		v4 = t3
		{
			if v3+i32(1) != v2 {
				goto l4
			}
			if v4&i32(255) == i32(5) {
				goto l5
			}
		l4:
			if uint32(v4&i32(255)) > uint32(i32(4)) {
				goto l6
			}
			goto l0
		l5:
			t4 := int32(m.memory[int64(uint32(v0))+777])
			if t4 != 0 {
				goto l6
			}
			if v3 == 0 {
				goto l0
			}
			t5 := int32(m.memory[uint32(v5+i32(-1))])
			if t5&i32(1) == 0 {
				goto l0
			}
		}
	l6:
		v1 = v1 + i64(1)
	}
l0:
	return v1
}
func (m *Module) fn217(v0, v1, v2 int32) {
	var v3, v4, v5 int32
	var v6 int64
	var v7 int32
	var v8, v9 int64
	var v10, v11, v12, v13, v14, v15 int32
	var v16, v17 int64
	var v18 float64
	var v19, v20, v21 int32
	t0 := m.g0
	v3 = t0 - i32(1600)
	m.g0 = v3
	{
		if v2 != 0 {
			goto l0
		}
		m.memory[int64(uint32(v0))+1] = byte(i32(0))
		v4 = i32(1)
		goto l1
	l0:
		{
			{
				{
					t1 := int32(m.memory[uint32(v1)])
					v5 = t1
					switch v5 + i32(-43) {
					case 0, 2:
						v4 = i32(1)
						v2 = v2 + i32(-1)
						if v2 == 0 {
							goto l4
						}
						v1 = v1 + i32(1)
						fallthrough
					default:
						v6 = i64(0)
						v4 = v1
						v7 = v2
						{
							{
								{
									{
										{
											if uint32(v2) < uint32(i32(8)) {
												goto l9
											}
											v6 = i64(0)
											v4 = v1
											v7 = v2
										l6:
											{
												t2 := int64(load64(m.memory[uint32(v4):]))
												v8 = t2
												t3 := v8 + i64(5063812098665367110)
												v8 = v8 + i64(-3472328296227680304)
												if !((t3|v8)&i64(-0x7f7f7f7f7f7f7f80) == 0) {
													goto l9
												}
												v8 = v8*i64(10) + int64(uint64(v8)>>8)
												v6 = int64(uint64(int64(uint64(v8)>>16)&i64(0xff000000ff)*i64(0x271000000001)+v8&i64(0xff000000ff)*i64(0xf424000000064))>>32) + v6*i64(100000000)
												v4 = v4 + i32(8)
												v7 = v7 + i32(-8)
												if uint32(v7) > uint32(i32(7)) {
													goto l6
												}
											}
											if v7 != 0 {
												goto l9
											}
											v9 = i64(0)
											v10 = i32(1)
											goto l7
										l9:
											{
												t4 := int32(m.memory[uint32(v4)])
												v11 = t4
												v12 = v11 + i32(-48)
												if uint32(v12&i32(255)) > uint32(i32(9)) {
													goto l8
												}
												v6 = v6*i64(10) + int64(uint32(v12))&i64(255)
												v10 = i32(1)
												v4 = v4 + i32(1)
												v7 = v7 + i32(-1)
												if v7 != 0 {
													goto l9
												}
											}
											v9 = i64(0)
										l7:
											v7 = i32(0)
											v11 = v2
											v8 = i64(0)
											goto l10
										l8:
											v13 = v2 - v7
											if v11&i32(255) == i32(46) {
												goto l11
											}
											v8 = i64(0)
											v11 = i32(0)
											v12 = v7
											goto l12
										l11:
											v4 = v4 + i32(1)
											v10 = v7 + i32(-1)
											if v7 >= i32(9) {
												goto l13
											}
											v12 = v10
											goto l14
										l13:
											v12 = v10
										l16:
											{
												t5 := int64(load64(m.memory[uint32(v4):]))
												v8 = t5
												t6 := v8 + i64(5063812098665367110)
												v8 = v8 + i64(-3472328296227680304)
												if !((t6|v8)&i64(-0x7f7f7f7f7f7f7f80) == 0) {
													goto l15
												}
												v8 = v8*i64(10) + int64(uint64(v8)>>8)
												v6 = int64(uint64(int64(uint64(v8)>>16)&i64(0xff000000ff)*i64(0x271000000001)+v8&i64(0xff000000ff)*i64(0xf424000000064))>>32) + v6*i64(100000000)
												v4 = v4 + i32(8)
												v12 = v12 + i32(-8)
												if uint32(v12) > uint32(i32(7)) {
													goto l16
												}
											}
										l14:
											if v12 == 0 {
												goto l17
											}
										l15:
											v11 = v4
											v4 = v11 + v12
										l20:
											{
												t7 := int32(m.memory[uint32(v11)])
												v14 = t7 + i32(-48)
												if uint32(v14&i32(255)) <= uint32(i32(9)) {
													goto l18
												}
												v4 = v11
												goto l19
											}
										l18:
											v6 = v6*i64(10) + int64(uint32(v14))&i64(255)
											v11 = v11 + i32(1)
											v12 = v12 + i32(-1)
											if v12 != 0 {
												goto l20
											}
										l17:
											v12 = i32(0)
										l19:
											v11 = v10 - v12
											v8 = int64(i32(0) - v11)
										l12:
											v11 = v11 + v13
											if v11 == 0 {
												goto l21
											}
											v9 = i64(0)
											if v12 != 0 {
												goto l22
											}
											v10 = i32(1)
											goto l10
										l22:
											{
												t8 := int32(m.memory[uint32(v4)])
												if t8|i32(32) == i32(101) {
													goto l23
												}
												v10 = i32(0)
												goto l10
											}
										l23:
											v13 = v12 + i32(-1)
											if v13 == 0 {
												goto l21
											}
											v14 = v4 + i32(1)
											t9 := int32(m.memory[int64(uint32(v4))+1])
											v10 = t9
											v15 = v10
											switch v10 + i32(-43) {
											case 0, 2:
												v13 = v12 + i32(-2)
												if v13 == 0 {
													goto l21
												}
												v14 = v4 + i32(2)
												t10 := int32(m.memory[int64(uint32(v4))+2])
												v15 = t10
												fallthrough
											default:
												if uint32((v15+i32(-48))&i32(255)) > uint32(i32(9)) {
													goto l21
												}
												v16 = i64(0)
												v9 = i64(0)
											l27:
												{
													t11 := int32(m.memory[uint32(v14)])
													v4 = t11 + i32(-48)
													if uint32(v4&i32(255)) > uint32(i32(9)) {
														goto l26
													}
													v17 = v9*i64(10) + int64(uint32(v4))&i64(255)
													t12 := v17
													t13 := v9
													var p14 int32
													if v9 < i64(65536) {
														p14 = 1
													}
													v4 = p14
													p15 := t13
													if v4 != 0 {
														p15 = t12
													}
													v9 = p15
													p16 := v16
													if v4 != 0 {
														p16 = v17
													}
													v16 = p16
													v14 = v14 + i32(1)
													v13 = v13 + i32(-1)
													if v13 != 0 {
														goto l27
													}
												}
												v13 = i32(0)
											l26:
												p17 := v16
												if v10 == i32(45) {
													p17 = i64(0) - v16
												}
												v9 = p17
												v8 = v9 + v8
												var p18 int32
												if v13 == 0 {
													p18 = 1
												}
												v10 = p18
											}
										}
									l10:
										if v11 >= i32(20) {
											goto l28
										}
										v4 = i32(0)
										goto l29
									l28:
										v11 = v11 + i32(-19)
										v14 = v2
										v4 = v1
									l32:
										{
											t19 := int32(m.memory[uint32(v4)])
											v12 = t19
											switch v12 + i32(-46) {
											default:
												goto l31
											case 0, 2:
												t20 := v11
												v13 = v12 + i32(-47)
												p21 := v13
												if uint32(v13) > uint32(v12) {
													p21 = i32(0)
												}
												v11 = t20 - p21
												v4 = v4 + i32(1)
												v14 = v14 + i32(-1)
												if v14 != 0 {
													goto l32
												}
											}
										}
									l31:
										if v11 >= i32(1) {
											goto l33
										}
										v4 = i32(0)
										goto l29
									l33:
										v12 = i32(0) - v2
										v6 = i64(0)
										v4 = v1
									l36:
										{
											v11 = v12
											t22 := int32(m.memory[uint32(v4)])
											v14 = t22 + i32(-48)
											if uint32(v14&i32(255)) > uint32(i32(9)) {
												goto l34
											}
											v4 = v4 + i32(1)
											v12 = v11 + i32(1)
											v6 = v6*i64(10) + int64(uint32(v14))&i64(255)
											if uint64(v6) > uint64(i64(999999999999999999)) {
												goto l35
											}
											if v12 != 0 {
												goto l36
											}
										l35:
										}
										if uint64(v6) > uint64(i64(999999999999999999)) {
											goto l37
										}
										if v11 == i32(-1) {
											m.fn151(i32(1), i32(0), i32(0), i32(1118296))
											panic("unreachable")
										}
										v7 = i32(0) - v12
										goto l39
									l34:
										v7 = i32(0) - v11
									l39:
										v14 = v7 + i32(-1)
										if v14 != 0 {
											v4 = v4 + i32(1)
											v7 = v14
										l44:
											{
												t23 := int32(m.memory[uint32(v4)])
												v12 = t23 + i32(-48)
												if uint32(v12&i32(255)) <= uint32(i32(9)) {
													v11 = v7 + i32(-1)
													{
														v6 = v6*i64(10) + int64(uint32(v12))&i64(255)
														if uint64(v6) > uint64(i64(999999999999999999)) {
															goto l43
														}
														v4 = v4 + i32(1)
														var p24 int32
														if v7 != i32(1) {
															p24 = 1
														}
														v12 = p24
														v7 = v11
														if v12 != 0 {
															goto l44
														}
													}
												l43:
													v4 = v11 - v14
													goto l41
												}
												v4 = v7 - v14
												goto l41
											}
										}
										v4 = i32(0) - v14
										goto l41
									l37:
										v4 = i32(0) - (v7 + v12)
									l41:
										v8 = v9 + int64(v4)
										v4 = i32(1)
									l29:
										if v10 == 0 {
											goto l21
										}
										var p25 int32
										if uint64(v8+i64(-38)) < uint64(i64(-60)) {
											p25 = 1
										}
										var p26 int32
										if uint64(v6) > uint64(i64(0x20000000000000)) {
											p26 = 1
										}
										if p25|p26|v4 != 0 {
											goto l45
										}
										{
											if v8 > i64(22) {
												t28 := int64(load64(m.memory[uint32(int32(v8)<<3+i32(1107632)):]))
												m.fn1853(v3, v6, i64(0), t28, i64(0))
												t29 := int64(load64(m.memory[int64(uint32(v3))+8:]))
												if t29 != i64(0) {
													goto l45
												}
												t30 := int64(load64(m.memory[uint32(v3):]))
												v9 = t30
												if uint64(v9) > uint64(i64(0x20000000000000)) {
													goto l45
												}
												v18 = float64(float64(uint64(v9)) * float64(1e+22))
												goto l48
											}
											v4 = int32(v8)
											v18 = float64(uint64(v6))
											if v8 < i64(0) {
												goto l47
											}
											t27 := math.Float64frombits(load64(m.memory[int64(uint32(v4<<3))+1131160:]))
											v18 = float64(t27 * v18)
											goto l48
										}
									}
								l21:
									switch v2 + i32(-3) {
									default:
										goto l50
									case 5:
										t31 := int64(load64(m.memory[uint32(v1):]))
										if t31&i64(-2314885530818453537) != i64(0x5954494e49464e49) {
											goto l50
										}
										v18 = math.Float64frombits(0x7ff0000000000000)
										goto l52
									case 0:
										{
											t32 := int64(load16(m.memory[uint32(v1):]))
											t33 := int64(m.memory[int64(uint32(v1))+2])
											v6 = (t32 | t33<<16) & i64(14671839)
											if v6 != i64(4607561) {
												goto l53
											}
											v18 = math.Float64frombits(0x7ff0000000000000)
											goto l52
										}
									l53:
										if v6 != i64(5128526) {
											goto l50
										}
										v18 = math.Float64frombits(0x7ff8000000000000)
									}
								l52:
									t35 := v0
									p34 := v18
									if v5 == i32(45) {
										p34 = -v18
									}
									store64(m.memory[int64(uint32(t35))+8:], math.Float64bits(p34))
									v4 = i32(0)
									goto l1
								}
							l47:
								t36 := math.Float64frombits(load64(m.memory[uint32(i32(1131160)-v4<<3):]))
								v18 = float64(v18 / t36)
							}
						l48:
							t38 := v0
							p37 := v18
							if v5 == i32(45) {
								p37 = -v18
							}
							store64(m.memory[int64(uint32(t38))+8:], math.Float64bits(p37))
							v4 = i32(0)
							goto l1
						}
					l45:
						m.fn1672(v3+i32(16), v8, v6)
						{
							{
								t39 := int32(load32(m.memory[int64(uint32(v3))+24:]))
								t40 := v4
								v11 = t39
								var p41 int32
								if v11 > i32(-1) {
									p41 = 1
								}
								if t40&p41 != 0 {
									goto l54
								}
								if v11 < i32(0) {
									goto l55
								}
								t42 := int64(load64(m.memory[int64(uint32(v3))+16:]))
								v6 = t42
								goto l56
							}
						l54:
							m.fn1672(v3+i32(816), v8, v6+i64(1))
							t43 := int64(load64(m.memory[int64(uint32(v3))+16:]))
							t44 := int64(load64(m.memory[int64(uint32(v3))+816:]))
							v6 = t44
							if t43 != v6 {
								goto l55
							}
							t45 := int32(load32(m.memory[int64(uint32(v3))+824:]))
							if v11 == t45 {
								goto l56
							}
						}
					l55:
						v19 = v3 + i32(816)
						v7 = i32(0)
						memory_zero(m.memory, uint32(v3+i32(816)), uint32(i32(777)))
						v15 = v3 + i32(824)
						v4 = i32(0)
					l64:
						{
							{
								v12 = v1 + v4
								t46 := int32(m.memory[uint32(v12)])
								v11 = t46
								if v11 == i32(48) {
									goto l57
								}
								v14 = v2 + v7
								v13 = v11 + i32(-48)
								if uint32(v13&i32(255)) > uint32(i32(9)) {
									if v11 == i32(46) {
										v11 = v12 + i32(1)
										v10 = v14 + i32(-1)
										goto l69
									}
									v10 = i32(0)
									v15 = i32(0)
									goto l68
								}
								v10 = v1 + v4
								v12 = v4 ^ i32(-1) + v2
								v4 = i32(0)
							l62:
								if uint32(v4) > uint32(i32(767)) {
									goto l59
								}
								m.memory[uint32(v15+v4)] = byte(v13)
							l59:
								v11 = v10 + v4 + i32(1)
								v7 = v4 + i32(1)
								{
									if v12 == v4 {
										store32(m.memory[uint32(v19):], uint32(v7))
										v15 = i32(0)
										v13 = i32(0)
										goto l63
									}
									v14 = v14 + i32(-1)
									v4 = v7
									t47 := int32(m.memory[uint32(v11)])
									v11 = t47
									v13 = v11 + i32(-48)
									if uint32(v13&i32(255)) > uint32(i32(9)) {
										v12 = v10 + v7
										store32(m.memory[int64(uint32(v3))+816:], uint32(v7))
										v15 = i32(0)
										if v11&i32(255) == i32(46) {
											goto l66
										}
										v13 = v14
										v11 = v12
										goto l63
									}
									goto l62
								}
							}
						l57:
							v7 = v7 + i32(-1)
							t48 := v2
							v4 = v4 + i32(1)
							if t48 != v4 {
								goto l64
							}
						}
						v10 = i32(0)
						goto l65
					l66:
						v11 = v10 + v7 + i32(-1) + i32(2)
						v10 = v14 + i32(1) + i32(-2)
						v13 = v10
						if v7 != 0 {
							goto l70
						}
					l69:
						if v10 != 0 {
							goto l71
						}
						v10 = i32(0)
						v7 = i32(0)
						v13 = i32(0)
						goto l72
					l71:
						v14 = v12 + v14
						v4 = i32(0)
					l74:
						{
							v12 = v11 + v4
							t49 := int32(m.memory[uint32(v12)])
							if t49 != i32(48) {
								goto l73
							}
							t50 := v10
							v4 = v4 + i32(1)
							if t50 != v4 {
								goto l74
							}
						}
						v7 = i32(0)
						v13 = i32(0)
						v11 = v14
						goto l72
					l73:
						v13 = v10 - v4
						v7 = i32(0)
						v11 = v12
					l70:
						if uint32(v13) < uint32(i32(8)) {
							goto l75
						}
						v4 = v7 + i32(8)
					l81:
						{
							v7 = v4
							if uint32(v7) < uint32(i32(768)) {
								goto l76
							}
							v7 = v7 + i32(-8)
							goto l77
						l76:
							t51 := int64(load64(m.memory[uint32(v11):]))
							v6 = t51
							t52 := v6 + i64(5063812098665367110)
							v6 = v6 + i64(-3472328296227680304)
							if (t52|v6)&i64(-0x7f7f7f7f7f7f7f80) != i64(0) {
								goto l78
							}
							v4 = v7 + i32(-8)
							if uint32(v4) > uint32(i32(768)) {
								goto l79
							}
							store64(m.memory[uint32(v3+i32(816)+v7):], uint64(v6))
							v4 = v7 + i32(8)
							v11 = v11 + i32(8)
							v13 = v13 + i32(-8)
							if uint32(v13) <= uint32(i32(7)) {
								goto l80
							}
							goto l81
						l79:
						}
						m.fn151(v4, i32(768), i32(768), i32(1116808))
						panic("unreachable")
					l50:
						v4 = i32(1)
					}
				}
			l4:
				m.memory[int64(uint32(v0))+1] = byte(v4)
				goto l1
			l78:
				v7 = v7 + i32(-8)
			l77:
				store32(m.memory[int64(uint32(v3))+816:], uint32(v7))
				goto l82
			l80:
				store32(m.memory[int64(uint32(v3))+816:], uint32(v7))
			l75:
				if v13 != 0 {
					goto l82
				}
				v13 = i32(0)
				goto l72
			l82:
				{
					t53 := int32(m.memory[uint32(v11)])
					v14 = t53 + i32(-48)
					if uint32(v14&i32(255)) > uint32(i32(9)) {
						goto l83
					}
					v20 = v11 + i32(1)
					v15 = v13 + i32(-1)
					v21 = v7 + (v3 + i32(816)) + i32(8)
					v12 = i32(0)
				l87:
					{
						t54 := v7
						v4 = v12
						v19 = t54 + v4
						if uint32(v19) > uint32(i32(767)) {
							goto l84
						}
						m.memory[uint32(v21+v4)] = byte(v14)
					}
				l84:
					{
						if v15 == v4 {
							goto l85
						}
						v13 = v13 + i32(-1)
						v12 = v4 + i32(1)
						t55 := int32(m.memory[uint32(v20+v4)])
						v14 = t55 + i32(-48)
						if uint32(v14&i32(255)) > uint32(i32(9)) {
							goto l86
						}
						goto l87
					}
				l85:
					v13 = i32(0)
				l86:
					v11 = v11 + v4 + i32(1)
					v7 = v19 + i32(1)
				}
			l83:
				store32(m.memory[int64(uint32(v3))+816:], uint32(v7))
			l72:
				t56 := v3
				v15 = v13 - v10
				store32(m.memory[int64(uint32(t56))+820:], uint32(v15))
			}
		l63:
			if v7 != 0 {
				v4 = v2 - v13
				{
					if uint32(v2) < uint32(v13) {
						m.fn151(i32(0), v4, v2, i32(1116824))
						panic("unreachable")
					}
					v12 = i32(0)
					if v2 == v13 {
						goto l91
					}
					v14 = v1 + i32(-1)
					v12 = i32(0)
				l94:
					{
						t57 := int32(m.memory[uint32(v14+v4)])
						switch t57 + i32(-46) {
						default:
							goto l91
						case 2:
							v12 = v12 + i32(1)
							fallthrough
						case 0:
							v4 = v4 + i32(-1)
							if v4 != 0 {
								goto l94
							}
						}
					}
				l91:
					t58 := v3
					v15 = v15 + v7
					store32(m.memory[int64(uint32(t58))+820:], uint32(v15))
					t59 := v3
					v10 = v7 - v12
					store32(m.memory[int64(uint32(t59))+816:], uint32(v10))
					if uint32(v10) < uint32(i32(769)) {
						goto l89
					}
					v10 = i32(768)
					store32(m.memory[int64(uint32(v3))+816:], uint32(i32(768)))
					m.memory[int64(uint32(v3))+1592] = byte(i32(1))
					goto l89
				}
			}
			v10 = i32(0)
			goto l89
		l89:
			v12 = v11
			v14 = v13
		l68:
			{
				if v14 == 0 {
					goto l95
				}
				t60 := int32(m.memory[uint32(v12)])
				if t60|i32(32) != i32(101) {
					goto l95
				}
				{
					v11 = v14 + i32(-1)
					if v11 != 0 {
						goto l96
					}
					v4 = i32(0)
					goto l97
				l96:
					{
						v7 = v12 + i32(1)
						t61 := int32(m.memory[uint32(v7)])
						v2 = t61
						switch v2 + i32(-43) {
						case 0, 2:
							v11 = v14 + i32(-2)
							if v11 == 0 {
								goto l100
							}
							v7 = v12 + i32(2)
							fallthrough
						default:
							v12 = i32(0)
							v4 = i32(0)
						l102:
							{
								t62 := int32(m.memory[uint32(v7)])
								v14 = (t62 + i32(-48)) & i32(255)
								if uint32(v14) > uint32(i32(9)) {
									goto l101
								}
								v14 = v4*i32(10) + v14
								t63 := v14
								t64 := v4
								var p65 int32
								if v4 < i32(65536) {
									p65 = 1
								}
								v13 = p65
								p66 := t64
								if v13 != 0 {
									p66 = t63
								}
								v4 = p66
								p67 := v12
								if v13 != 0 {
									p67 = v14
								}
								v12 = p67
								v7 = v7 + i32(1)
								v11 = v11 + i32(-1)
								if v11 != 0 {
									goto l102
								}
								goto l101
							}
						}
					}
				l100:
					v12 = i32(0)
				l101:
					p68 := v12
					if v2 == i32(45) {
						p68 = i32(0) - v12
					}
					v4 = p68
				}
			l97:
				store32(m.memory[int64(uint32(v3))+820:], uint32(v15+v4))
			}
		l95:
			if uint32(v10) > uint32(i32(18)) {
				goto l103
			}
		l65:
			v4 = i32(19) - v10
			if v4 == 0 {
				goto l103
			}
			memory_zero(m.memory, uint32(v3+i32(816)+v10+i32(8)), uint32(v4))
		l103:
			memory_copy(m.memory, uint32(v3+i32(36)), uint32(v3+i32(816)), uint32(i32(780)))
			v6 = i64(0)
			v11 = i32(0)
			t69 := int32(load32(m.memory[int64(uint32(v3))+36:]))
			if t69 == 0 {
				goto l56
			}
			t70 := int32(load32(m.memory[int64(uint32(v3))+40:]))
			v4 = t70
			if v4 < i32(-324) {
				goto l56
			}
			v11 = i32(2047)
			if v4 > i32(309) {
				goto l56
			}
			if v4 >= i32(1) {
				v7 = i32(0)
			l108:
				v12 = i32(60)
				{
					if uint32(v4) >= uint32(i32(19)) {
						goto l106
					}
					t71 := int32(m.memory[int64(uint32(v4))+1108132])
					v12 = t71
				}
			l106:
				m.fn1654(v3+i32(36), v12)
				{
					t72 := int32(load32(m.memory[int64(uint32(v3))+40:]))
					v4 = t72
					if v4 <= i32(-2048) {
						v11 = i32(0)
						goto l56
					}
					v7 = v12 + v7
					if v4 < i32(1) {
						goto l105
					}
					goto l108
				}
			}
			v7 = i32(0)
			goto l105
		l105:
			v14 = v3 + i32(44)
		l113:
			{
				{
					if v4 != 0 {
						goto l109
					}
					t73 := int32(m.memory[uint32(v14)])
					v4 = t73
					if uint32(v4) > uint32(i32(4)) {
						goto l110
					}
					p74 := i32(1)
					if uint32(v4) < uint32(i32(2)) {
						p74 = i32(2)
					}
					v12 = p74
					goto l111
				}
			l109:
				v12 = i32(60)
				v4 = i32(0) - v4
				if uint32(v4) >= uint32(i32(19)) {
					goto l111
				}
				t75 := int32(m.memory[int64(uint32(v4))+1108132])
				v12 = t75
			}
		l111:
			m.fn1653(v3+i32(36), v12)
			{
				t76 := int32(load32(m.memory[int64(uint32(v3))+40:]))
				v4 = t76
				if v4 <= i32(2047) {
					goto l112
				}
				v11 = i32(2047)
				goto l56
			}
		l112:
			v7 = v7 - v12
			if v4 < i32(1) {
				goto l113
			}
		l110:
			v4 = v7 + i32(-1)
			if v4 > i32(-1023) {
				goto l114
			}
		l115:
			{
				t77 := v3 + i32(36)
				v7 = i32(-1022) - v4
				p78 := i32(60)
				if uint32(v7) < uint32(i32(60)) {
					p78 = v7
				}
				v7 = p78
				m.fn1654(t77, v7)
				v4 = v7 + v4
				if uint32(v4) < uint32(i32(-1022)) {
					goto l115
				}
			}
		l114:
			if v4+i32(1023) > i32(2046) {
				goto l56
			}
			m.fn1653(v3+i32(36), i32(53))
			{
				{
					t79 := int32(load32(m.memory[int64(uint32(v3))+36:]))
					v13 = t79
					if v13 == 0 {
						goto l116
					}
					t80 := int32(load32(m.memory[int64(uint32(v3))+40:]))
					v12 = t80
					if v12 < i32(0) {
						goto l116
					}
					if uint32(v12) > uint32(i32(18)) {
						goto l117
					}
					if v12 != 0 {
						v7 = i32(0)
						v8 = i64(0)
					l121:
						{
							v8 = v8 * i64(10)
							{
								if uint32(v7) >= uint32(v13) {
									goto l120
								}
								t81 := int64(m.memory[uint32(v14+v7)])
								v8 = v8 + t81
							}
						l120:
							t82 := v12
							v7 = v7 + i32(1)
							if t82 == v7 {
								goto l119
							}
							goto l121
						}
					}
					v8 = i64(0)
					goto l119
				}
			l116:
				v11 = v4 + i32(1022)
				goto l56
			l119:
				{
					if uint32(v12) >= uint32(v13) {
						goto l122
					}
					v14 = v14 + v12
					t83 := int32(m.memory[uint32(v14)])
					v7 = t83
					{
						if v12+i32(1) != v13 {
							goto l123
						}
						if v7&i32(255) == i32(5) {
							goto l124
						}
					l123:
						if uint32(v7&i32(255)) > uint32(i32(4)) {
							goto l125
						}
						goto l122
					l124:
						t84 := int32(m.memory[int64(uint32(v3))+812])
						if t84 != 0 {
							goto l125
						}
						if v12 == 0 {
							goto l122
						}
						t85 := int32(m.memory[uint32(v14+i32(-1))])
						if t85&i32(1) == 0 {
							goto l122
						}
					}
				l125:
					v8 = v8 + i64(1)
				}
			l122:
				if uint64(v8) < uint64(i64(0x20000000000000)) {
					goto l126
				}
			l117:
				m.fn1654(v3+i32(36), i32(1))
				t86 := m.fn1655(v3 + i32(36))
				v8 = t86
				if v4+i32(1024) > i32(2046) {
					goto l56
				}
				v4 = v4 + i32(1)
			}
		l126:
			v6 = v8 & i64(0xfffffffffffff)
			p87 := i32(1023)
			if uint64(v8) < uint64(i64(0x10000000000000)) {
				p87 = i32(1022)
			}
			v11 = p87 + v4
		}
	l56:
		t88 := v0
		v18 = math.Float64frombits(uint64(int64(uint32(v11))<<52 | v6))
		p89 := v18
		if v5 == i32(45) {
			p89 = -v18
		}
		store64(m.memory[int64(uint32(t88))+8:], math.Float64bits(p89))
		v4 = i32(0)
	}
l1:
	m.memory[uint32(v0)] = byte(v4)
	m.g0 = v3 + i32(1600)
}
func (m *Module) fn218(v0, v1 int32) {
	var v2 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		v2 = t0
		t1 := int32(load32(m.memory[uint32(v0):]))
		if v2 != t1 {
			goto l0
		}
		m.fn271(v0)
	}
l0:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v2+i32(1)))
	t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	v0 = t2 + v2*i32(24)
	t3 := int64(load64(m.memory[uint32(v1):]))
	store64(m.memory[uint32(v0):], uint64(t3))
	t4 := int64(load64(m.memory[int64(uint32(v1))+8:]))
	store64(m.memory[int64(uint32(v0))+8:], uint64(t4))
	t5 := int64(load64(m.memory[int64(uint32(v1))+16:]))
	store64(m.memory[int64(uint32(v0))+16:], uint64(t5))
}
func (m *Module) fn219(v0, v1 int32) {
	t0 := int32(m.memory[uint32(v1)])
	switch t0 {
	default:
		t1 := int64(load64(m.memory[int64(uint32(v1))+16:]))
		store64(m.memory[int64(uint32(v0))+16:], uint64(t1))
		t2 := int64(load64(m.memory[int64(uint32(v1))+8:]))
		store64(m.memory[int64(uint32(v0))+8:], uint64(t2))
		t3 := int64(load64(m.memory[uint32(v1):]))
		store64(m.memory[uint32(v0):], uint64(t3))
		return
	case 1:
		t4 := int64(load64(m.memory[int64(uint32(v1))+16:]))
		store64(m.memory[int64(uint32(v0))+16:], uint64(t4))
		t5 := int64(load64(m.memory[int64(uint32(v1))+8:]))
		store64(m.memory[int64(uint32(v0))+8:], uint64(t5))
		t6 := int64(load64(m.memory[uint32(v1):]))
		store64(m.memory[uint32(v0):], uint64(t6))
		return
	case 2:
		t7 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		t8 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		m.fn31(v0+i32(4), t7, t8)
		m.memory[uint32(v0)] = byte(i32(2))
		return
	case 3:
		t9 := int64(load64(m.memory[int64(uint32(v1))+16:]))
		store64(m.memory[int64(uint32(v0))+16:], uint64(t9))
		t10 := int64(load64(m.memory[int64(uint32(v1))+8:]))
		store64(m.memory[int64(uint32(v0))+8:], uint64(t10))
		t11 := int64(load64(m.memory[uint32(v1):]))
		store64(m.memory[uint32(v0):], uint64(t11))
		return
	case 4:
		m.memory[uint32(v0)] = byte(i32(4))
		t12 := int64(load64(m.memory[int64(uint32(v1))+16:]))
		store64(m.memory[int64(uint32(v0))+16:], uint64(t12))
		t13 := int64(load64(m.memory[int64(uint32(v1))+8:]))
		store64(m.memory[int64(uint32(v0))+8:], uint64(t13))
		return
	case 5:
		t14 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		t15 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		m.fn31(v0+i32(4), t14, t15)
		m.memory[uint32(v0)] = byte(i32(5))
		return
	case 6:
		t16 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		t17 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		m.fn31(v0+i32(4), t16, t17)
		m.memory[uint32(v0)] = byte(i32(6))
		return
	case 7:
		m.memory[uint32(v0)] = byte(i32(7))
		t18 := int32(m.memory[int64(uint32(v1))+1])
		m.memory[int64(uint32(v0))+1] = byte(t18)
		return
	case 8:
		m.memory[uint32(v0)] = byte(i32(8))
	}
}
func (m *Module) fn220(v0, v1, v2, v3, v4 int32) {
	var v5, v6, v7, v8 int32
	var v9, v10 int64
	var v11, v12, v13 int32
	t0 := m.g0
	v5 = t0 - i32(48)
	m.g0 = v5
	t1 := int32(m.memory[int64(uint32(v1))+246])
	v6 = t1
	m.memory[int64(uint32(v1))+246] = byte(i32(0))
	v7 = v5 + i32(16)
	v8 = v5 + i32(12)
	t2 := int64(load64(m.memory[int64(uint32(v1))+248:]))
	v9 = t2
	v10 = v9
	v11 = i32(0)
l16:
	{
		store32(m.memory[int64(uint32(v4))+8:], uint32(i32(0)))
		m.fn141(v5+i32(8), v1, v4)
		{
			t3 := int32(load32(m.memory[int64(uint32(v5))+8:]))
			if t3 != i32(1) {
				goto l0
			}
			m.memory[int64(uint32(v1))+246] = byte(v6)
			t4 := int64(load64(m.memory[int64(uint32(v8))+16:]))
			store64(m.memory[int64(uint32(v0))+16:], uint64(t4))
			t5 := int64(load64(m.memory[int64(uint32(v8))+8:]))
			store64(m.memory[int64(uint32(v0))+8:], uint64(t5))
			t6 := int64(load64(m.memory[uint32(v8):]))
			store64(m.memory[uint32(v0):], uint64(t6))
			goto l1
		}
	l0:
		{
			t7 := int32(load32(m.memory[int64(uint32(v5))+12:]))
			v12 = t7
			switch v12 {
			case 0:
				m.fn164(v5, v7)
				t8 := int32(load32(m.memory[uint32(v5):]))
				t9 := int32(load32(m.memory[int64(uint32(v5))+4:]))
				t10 := m.fn123(t8, t9, v2, v3)
				v12 = t10
				t11 := int32(load32(m.memory[int64(uint32(v5))+16:]))
				t12 := int32(load32(m.memory[int64(uint32(v5))+20:]))
				m.fn134(t11, t12)
				p13 := v11
				if v12 != 0 {
					p13 = v11 + i32(1)
				}
				v11 = p13
				goto l5
			default:
				if v12 == i32(10) {
					m.memory[int64(uint32(v1))+246] = byte(v6)
					t14 := int32(load32(m.memory[int64(uint32(v1))+236:]))
					m.fn198(v5+i32(36), v2, v3, t14)
					v1 = v0 + i32(4)
					{
						t15 := int32(load32(m.memory[int64(uint32(v5))+36:]))
						v4 = t15
						if v4 != i32(-2) {
							goto l7
						}
						t16 := int64(load64(m.memory[int64(uint32(v5))+40:]))
						store64(m.memory[uint32(v1):], uint64(t16))
						v1 = i32(-0x7ffffff4)
						goto l8
					}
				l7:
					{
						if v4 == i32(-1) {
							goto l9
						}
						t17 := int32(load32(m.memory[int64(uint32(v5))+44:]))
						store32(m.memory[int64(uint32(v1))+8:], uint32(t17))
						t18 := int64(load64(m.memory[int64(uint32(v5))+36:]))
						store64(m.memory[uint32(v1):], uint64(t18))
						goto l10
					}
				l9:
					{
						{
							t19 := int32(load32(m.memory[int64(uint32(v5))+44:]))
							v1 = t19
							if v1 != 0 {
								goto l11
							}
							v4 = i32(1)
							goto l12
						}
					l11:
						t20 := int32(load32(m.memory[int64(uint32(v5))+40:]))
						v11 = t20
						t21 := m.fn4(v1)
						v4 = t21
						if v4 == 0 {
							m.fn2(i32(1), v1)
							panic("unreachable")
						}
						if v1 == 0 {
							goto l12
						}
						memory_copy(m.memory, uint32(v4), uint32(v11), uint32(v1))
					}
				l12:
					store32(m.memory[int64(uint32(v0))+12:], uint32(v1))
					store32(m.memory[int64(uint32(v0))+8:], uint32(v4))
					store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
				l10:
					v1 = i32(-0x7ffffffd)
				l8:
					store32(m.memory[uint32(v0):], uint32(v1))
					m.fn200(v8)
					goto l1
				}
				m.fn200(v8)
				goto l5
			case 1:
				t22 := int32(load32(m.memory[int64(uint32(v5))+16:]))
				v12 = t22
				t23 := int32(load32(m.memory[int64(uint32(v5))+20:]))
				v13 = t23
				t24 := int32(load32(m.memory[int64(uint32(v5))+24:]))
				t25 := m.fn123(v13, t24, v2, v3)
				if t25 != 0 {
					goto l14
				}
				t26 := int32(load32(m.memory[int64(uint32(v5))+20:]))
				m.fn134(v12, t26)
				goto l5
			}
		}
	l14:
		if v11 != 0 {
			goto l15
		}
		m.memory[int64(uint32(v1))+246] = byte(v6)
		m.fn134(v12, v13)
		store64(m.memory[int64(uint32(v0))+16:], uint64(v10))
		store64(m.memory[int64(uint32(v0))+8:], uint64(v9))
		store32(m.memory[uint32(v0):], uint32(i32(-1)))
	l1:
		m.g0 = v5 + i32(48)
		return
	l15:
		m.fn134(v12, v13)
		v11 = v11 + i32(-1)
	l5:
		t27 := int64(load64(m.memory[int64(uint32(v1))+248:]))
		v10 = t27
		goto l16
	}
}
func (m *Module) fn221(v0, v1 int32) int32 {
	var v2 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	{
		t1 := int32(load32(m.memory[uint32(v0):]))
		switch t1 {
		default:
			store32(m.memory[int64(uint32(v2))+12:], uint32(v0+i32(4)))
			t2 := m.fn264(v1, i32(1087240), i32(5), v2+i32(12), i32(50))
			v0 = t2
			goto l11
		case 1:
			store32(m.memory[int64(uint32(v2))+12:], uint32(v0+i32(4)))
			t3 := m.fn264(v1, i32(1087245), i32(3), v2+i32(12), i32(51))
			v0 = t3
			goto l11
		case 2:
			store32(m.memory[int64(uint32(v2))+12:], uint32(v0+i32(4)))
			t4 := m.fn264(v1, i32(1284280), i32(5), v2+i32(12), i32(50))
			v0 = t4
			goto l11
		case 3:
			store32(m.memory[int64(uint32(v2))+12:], uint32(v0+i32(4)))
			t5 := m.fn264(v1, i32(1087248), i32(4), v2+i32(12), i32(52))
			v0 = t5
			goto l11
		case 4:
			store32(m.memory[int64(uint32(v2))+12:], uint32(v0+i32(4)))
			t6 := m.fn264(v1, i32(1087252), i32(5), v2+i32(12), i32(53))
			v0 = t6
			goto l11
		case 5:
			store32(m.memory[int64(uint32(v2))+12:], uint32(v0+i32(4)))
			t7 := m.fn264(v1, i32(1087257), i32(7), v2+i32(12), i32(52))
			v0 = t7
			goto l11
		case 6:
			store32(m.memory[int64(uint32(v2))+12:], uint32(v0+i32(4)))
			t8 := m.fn264(v1, i32(1087264), i32(4), v2+i32(12), i32(54))
			v0 = t8
			goto l11
		case 7:
			store32(m.memory[int64(uint32(v2))+12:], uint32(v0+i32(4)))
			t9 := m.fn264(v1, i32(1089862), i32(2), v2+i32(12), i32(55))
			v0 = t9
			goto l11
		case 8:
			store32(m.memory[int64(uint32(v2))+12:], uint32(v0+i32(4)))
			t10 := m.fn264(v1, i32(1087268), i32(7), v2+i32(12), i32(52))
			v0 = t10
			goto l11
		case 9:
			store32(m.memory[int64(uint32(v2))+12:], uint32(v0+i32(4)))
			t11 := m.fn264(v1, i32(1087275), i32(10), v2+i32(12), i32(56))
			v0 = t11
			goto l11
		case 10:
			t12 := int32(load32(m.memory[uint32(v1):]))
			t13 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t14 := int32(load32(m.memory[int64(uint32(t13))+12:]))
			t15 := m.t0[uint(t14)].(func(int32, int32, int32) int32)(t12, i32(1087285), i32(3))
			v0 = t15
		}
	}
l11:
	m.g0 = v2 + i32(16)
	return v0
}
