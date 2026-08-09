package core

import (
	"math/bits"
)

func (m *Module) fn717(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7 int32
	var v8 int64
	var v9, v10 int32
	var v11 int64
	var v12, v13 int32
	t0 := m.g0
	v2 = t0 - i32(64)
	m.g0 = v2
	store32(m.memory[int64(uint32(v2))+12:], uint32(v1))
	t1 := int32(load32(m.memory[int64(uint32(v0))+12:]))
	v3 = t1
	store32(m.memory[int64(uint32(v2))+16:], uint32(v2+i32(12)))
	v1 = v3 + i32(1)
	if v1 == 0 {
		m.fn242()
		panic("unreachable")
	}
	{
		t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		t3 := v1
		v4 = t2
		p4 := int32(uint32(v4+i32(1))>>3) * i32(7)
		if uint32(v4) < uint32(i32(8)) {
			p4 = v4
		}
		v4 = p4
		if uint32(t3) <= uint32(int32(uint32(v4)>>1)) {
			goto l1
		}
		t5 := v2 + i32(48)
		v4 = v4 + i32(1)
		p6 := v1
		if uint32(v4) > uint32(v1) {
			p6 = v4
		}
		m.fn237(t5, i32(12), p6)
		t7 := int32(load32(m.memory[int64(uint32(v2))+52:]))
		v5 = t7
		t8 := int32(load32(m.memory[int64(uint32(v2))+48:]))
		v6 = t8
		if v6 == 0 {
			goto l2
		}
		t9 := int32(load32(m.memory[int64(uint32(v2))+56:]))
		v7 = t9
		t10 := int32(load32(m.memory[int64(uint32(v2))+60:]))
		store32(m.memory[int64(uint32(v2))+44:], uint32(t10))
		store32(m.memory[int64(uint32(v2))+40:], uint32(v7))
		store32(m.memory[int64(uint32(v2))+36:], uint32(v5))
		store32(m.memory[int64(uint32(v2))+32:], uint32(v6))
		store64(m.memory[int64(uint32(v2))+24:], uint64(i64(0x80000000c)))
		store32(m.memory[int64(uint32(v2))+20:], uint32(v0+i32(16)))
		t11 := int32(load32(m.memory[uint32(v0):]))
		v4 = t11
		t12 := int64(load64(m.memory[uint32(v4):]))
		v8 = (t12 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
		v9 = v2 + i32(20) + i32(12)
		v1 = i32(0)
	l6:
		if v3 == 0 {
			t26 := int32(load32(m.memory[int64(uint32(v0))+12:]))
			t27 := v2
			v1 = t26
			store32(m.memory[int64(uint32(t27))+44:], uint32(v1))
			store32(m.memory[int64(uint32(v2))+40:], uint32(v7-v1))
			m.fn239(v0, v9)
			m.fn240(v2 + i32(20))
			goto l7
		}
	l5:
		{
			if v8 != i64(0) {
				t14 := v6
				t15 := v6
				t16 := v5
				t17 := v2 + i32(16)
				t18 := v0
				v10 = int32(uint32(int64(bits.TrailingZeros64(uint64(v8))))>>3) + v1
				t19 := m.fn718(t17, t18, v10)
				v11 = t19
				t20 := m.fn26(t15, t16, v11)
				v12 = t20
				t21 := t14 + v12
				v13 = int32(uint32(int32(v11)) >> 25)
				m.memory[uint32(t21)] = byte(v13)
				m.memory[uint32(v6+v5&(v12+i32(-8))+i32(8))] = byte(v13)
				v12 = v6 + (v12^i32(-1))*i32(12)
				t22 := int32(load32(m.memory[uint32(v0):]))
				t23 := v12
				v10 = t22 + (v10^i32(-1))*i32(12)
				t24 := int32(load32(m.memory[int64(uint32(v10))+8:]))
				store32(m.memory[int64(uint32(t23))+8:], uint32(t24))
				t25 := int64(load64(m.memory[uint32(v10):]))
				store64(m.memory[uint32(v12):], uint64(t25))
				v3 = v3 + i32(-1)
				v8 = (v8 + i64(-1)) & v8
				goto l6
			}
			v1 = v1 + i32(8)
			v4 = v4 + i32(8)
			t13 := int64(load64(m.memory[uint32(v4):]))
			v8 = (t13 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
			goto l5
		}
	}
l1:
	m.fn241(v0, v2+i32(16), i32(126), i32(12))
l7:
	v5 = i32(-1)
l2:
	m.g0 = v2 + i32(64)
	return v5
}
func (m *Module) fn718(v0, v1, v2 int32) int64 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(t0):]))
	v0 = t1
	t2 := int64(load64(m.memory[uint32(v0):]))
	t3 := int64(load64(m.memory[uint32(v0+i32(8)):]))
	t4 := int32(load32(m.memory[uint32(v1):]))
	t5 := int32(load32(m.memory[uint32(t4+(i32(0)-v2)*i32(12)+i32(-12)):]))
	t6 := m.fn66(t2, t3, t5)
	return t6
}
func (m *Module) fn719(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7 int32
	var v8 int64
	var v9, v10 int32
	var v11 int64
	var v12, v13 int32
	t0 := m.g0
	v2 = t0 - i32(64)
	m.g0 = v2
	store32(m.memory[int64(uint32(v2))+12:], uint32(v1))
	t1 := int32(load32(m.memory[int64(uint32(v0))+12:]))
	v3 = t1
	store32(m.memory[int64(uint32(v2))+16:], uint32(v2+i32(12)))
	v1 = v3 + i32(1)
	if v1 == 0 {
		m.fn242()
		panic("unreachable")
	}
	{
		t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		t3 := v1
		v4 = t2
		p4 := int32(uint32(v4+i32(1))>>3) * i32(7)
		if uint32(v4) < uint32(i32(8)) {
			p4 = v4
		}
		v4 = p4
		if uint32(t3) <= uint32(int32(uint32(v4)>>1)) {
			goto l1
		}
		t5 := v2 + i32(48)
		v4 = v4 + i32(1)
		p6 := v1
		if uint32(v4) > uint32(v1) {
			p6 = v4
		}
		m.fn237(t5, i32(8), p6)
		t7 := int32(load32(m.memory[int64(uint32(v2))+52:]))
		v5 = t7
		t8 := int32(load32(m.memory[int64(uint32(v2))+48:]))
		v6 = t8
		if v6 == 0 {
			goto l2
		}
		t9 := int32(load32(m.memory[int64(uint32(v2))+56:]))
		v7 = t9
		t10 := int32(load32(m.memory[int64(uint32(v2))+60:]))
		store32(m.memory[int64(uint32(v2))+44:], uint32(t10))
		store32(m.memory[int64(uint32(v2))+40:], uint32(v7))
		store32(m.memory[int64(uint32(v2))+36:], uint32(v5))
		store32(m.memory[int64(uint32(v2))+32:], uint32(v6))
		store64(m.memory[int64(uint32(v2))+24:], uint64(i64(0x800000008)))
		store32(m.memory[int64(uint32(v2))+20:], uint32(v0+i32(16)))
		t11 := int32(load32(m.memory[uint32(v0):]))
		v4 = t11
		t12 := int64(load64(m.memory[uint32(v4):]))
		v8 = (t12 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
		v9 = v2 + i32(32)
		v1 = i32(0)
	l6:
		if v3 == 0 {
			t24 := int32(load32(m.memory[int64(uint32(v0))+12:]))
			t25 := v2
			v1 = t24
			store32(m.memory[int64(uint32(t25))+44:], uint32(v1))
			store32(m.memory[int64(uint32(v2))+40:], uint32(v7-v1))
			m.fn239(v0, v9)
			m.fn240(v2 + i32(20))
			goto l7
		}
	l5:
		{
			if v8 != i64(0) {
				t14 := v6
				t15 := v6
				t16 := v5
				t17 := v2 + i32(16)
				t18 := v0
				v10 = int32(uint32(int64(bits.TrailingZeros64(uint64(v8))))>>3) + v1
				t19 := m.fn720(t17, t18, v10)
				v11 = t19
				t20 := m.fn26(t15, t16, v11)
				v12 = t20
				t21 := t14 + v12
				v13 = int32(uint32(int32(v11)) >> 25)
				m.memory[uint32(t21)] = byte(v13)
				m.memory[uint32(v6+v5&(v12+i32(-8))+i32(8))] = byte(v13)
				t22 := int32(load32(m.memory[uint32(v0):]))
				t23 := int64(load64(m.memory[uint32(t22+(v10^i32(-1))<<3):]))
				store64(m.memory[uint32(v6+(v12^i32(-1))<<3):], uint64(t23))
				v3 = v3 + i32(-1)
				v8 = (v8 + i64(-1)) & v8
				goto l6
			}
			v1 = v1 + i32(8)
			v4 = v4 + i32(8)
			t13 := int64(load64(m.memory[uint32(v4):]))
			v8 = (t13 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
			goto l5
		}
	}
l1:
	m.fn241(v0, v2+i32(16), i32(127), i32(8))
l7:
	v5 = i32(-1)
l2:
	m.g0 = v2 + i32(64)
	return v5
}
func (m *Module) fn720(v0, v1, v2 int32) int64 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(t0):]))
	v0 = t1
	t2 := int64(load64(m.memory[uint32(v0):]))
	t3 := int64(load64(m.memory[uint32(v0+i32(8)):]))
	t4 := int32(load32(m.memory[uint32(v1):]))
	t5 := int32(load32(m.memory[uint32(t4-v2<<3+i32(-8)):]))
	t6 := m.fn66(t2, t3, t5)
	return t6
}
func (m *Module) fn721(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7 int32
	var v8 int64
	var v9, v10 int32
	var v11 int64
	var v12, v13 int32
	t0 := m.g0
	v2 = t0 - i32(64)
	m.g0 = v2
	store32(m.memory[int64(uint32(v2))+12:], uint32(v1))
	t1 := int32(load32(m.memory[int64(uint32(v0))+12:]))
	v3 = t1
	store32(m.memory[int64(uint32(v2))+16:], uint32(v2+i32(12)))
	v1 = v3 + i32(1)
	if v1 == 0 {
		m.fn242()
		panic("unreachable")
	}
	{
		t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		t3 := v1
		v4 = t2
		p4 := int32(uint32(v4+i32(1))>>3) * i32(7)
		if uint32(v4) < uint32(i32(8)) {
			p4 = v4
		}
		v4 = p4
		if uint32(t3) <= uint32(int32(uint32(v4)>>1)) {
			goto l1
		}
		t5 := v2 + i32(48)
		v4 = v4 + i32(1)
		p6 := v1
		if uint32(v4) > uint32(v1) {
			p6 = v4
		}
		m.fn237(t5, i32(96), p6)
		t7 := int32(load32(m.memory[int64(uint32(v2))+52:]))
		v5 = t7
		t8 := int32(load32(m.memory[int64(uint32(v2))+48:]))
		v6 = t8
		if v6 == 0 {
			goto l2
		}
		t9 := int32(load32(m.memory[int64(uint32(v2))+56:]))
		v7 = t9
		t10 := int32(load32(m.memory[int64(uint32(v2))+60:]))
		store32(m.memory[int64(uint32(v2))+44:], uint32(t10))
		store32(m.memory[int64(uint32(v2))+40:], uint32(v7))
		store32(m.memory[int64(uint32(v2))+36:], uint32(v5))
		store32(m.memory[int64(uint32(v2))+32:], uint32(v6))
		store64(m.memory[int64(uint32(v2))+24:], uint64(i64(0x800000060)))
		store32(m.memory[int64(uint32(v2))+20:], uint32(v0+i32(16)))
		t11 := int32(load32(m.memory[uint32(v0):]))
		v4 = t11
		t12 := int64(load64(m.memory[uint32(v4):]))
		v8 = (t12 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
		v9 = v2 + i32(32)
		v1 = i32(0)
	l6:
		if v3 == 0 {
			t23 := int32(load32(m.memory[int64(uint32(v0))+12:]))
			t24 := v2
			v1 = t23
			store32(m.memory[int64(uint32(t24))+44:], uint32(v1))
			store32(m.memory[int64(uint32(v2))+40:], uint32(v7-v1))
			m.fn239(v0, v9)
			m.fn240(v2 + i32(20))
			goto l7
		}
	l5:
		{
			if v8 != i64(0) {
				t14 := v6
				t15 := v6
				t16 := v5
				t17 := v2 + i32(16)
				t18 := v0
				v10 = int32(uint32(int64(bits.TrailingZeros64(uint64(v8))))>>3) + v1
				t19 := m.fn722(t17, t18, v10)
				v11 = t19
				t20 := m.fn26(t15, t16, v11)
				v12 = t20
				t21 := t14 + v12
				v13 = int32(uint32(int32(v11)) >> 25)
				m.memory[uint32(t21)] = byte(v13)
				m.memory[uint32(v6+v5&(v12+i32(-8))+i32(8))] = byte(v13)
				t22 := int32(load32(m.memory[uint32(v0):]))
				memory_copy(m.memory, uint32(v6+(v12^i32(-1))*i32(96)), uint32(t22+(v10^i32(-1))*i32(96)), uint32(i32(96)))
				v3 = v3 + i32(-1)
				v8 = (v8 + i64(-1)) & v8
				goto l6
			}
			v1 = v1 + i32(8)
			v4 = v4 + i32(8)
			t13 := int64(load64(m.memory[uint32(v4):]))
			v8 = (t13 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
			goto l5
		}
	}
l1:
	m.fn241(v0, v2+i32(16), i32(128), i32(96))
l7:
	v5 = i32(-1)
l2:
	m.g0 = v2 + i32(64)
	return v5
}
func (m *Module) fn722(v0, v1, v2 int32) int64 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(t0):]))
	v0 = t1
	t2 := int64(load64(m.memory[uint32(v0):]))
	t3 := int64(load64(m.memory[uint32(v0+i32(8)):]))
	t4 := int32(load32(m.memory[uint32(v1):]))
	t5 := int32(load32(m.memory[uint32(t4+(i32(0)-v2)*i32(96)+i32(-96)):]))
	t6 := m.fn66(t2, t3, t5)
	return t6
}
func (m *Module) fn723(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7 int32
	var v8 int64
	var v9, v10 int32
	var v11 int64
	var v12, v13 int32
	t0 := m.g0
	v2 = t0 - i32(64)
	m.g0 = v2
	store32(m.memory[int64(uint32(v2))+12:], uint32(v1))
	t1 := int32(load32(m.memory[int64(uint32(v0))+12:]))
	v3 = t1
	store32(m.memory[int64(uint32(v2))+16:], uint32(v2+i32(12)))
	v1 = v3 + i32(1)
	if v1 == 0 {
		m.fn242()
		panic("unreachable")
	}
	{
		t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		t3 := v1
		v4 = t2
		p4 := int32(uint32(v4+i32(1))>>3) * i32(7)
		if uint32(v4) < uint32(i32(8)) {
			p4 = v4
		}
		v4 = p4
		if uint32(t3) <= uint32(int32(uint32(v4)>>1)) {
			goto l1
		}
		t5 := v2 + i32(48)
		v4 = v4 + i32(1)
		p6 := v1
		if uint32(v4) > uint32(v1) {
			p6 = v4
		}
		m.fn237(t5, i32(20), p6)
		t7 := int32(load32(m.memory[int64(uint32(v2))+52:]))
		v5 = t7
		t8 := int32(load32(m.memory[int64(uint32(v2))+48:]))
		v6 = t8
		if v6 == 0 {
			goto l2
		}
		t9 := int32(load32(m.memory[int64(uint32(v2))+56:]))
		v7 = t9
		t10 := int32(load32(m.memory[int64(uint32(v2))+60:]))
		store32(m.memory[int64(uint32(v2))+44:], uint32(t10))
		store32(m.memory[int64(uint32(v2))+40:], uint32(v7))
		store32(m.memory[int64(uint32(v2))+36:], uint32(v5))
		store32(m.memory[int64(uint32(v2))+32:], uint32(v6))
		store64(m.memory[int64(uint32(v2))+24:], uint64(i64(0x800000014)))
		store32(m.memory[int64(uint32(v2))+20:], uint32(v0+i32(16)))
		t11 := int32(load32(m.memory[uint32(v0):]))
		v4 = t11
		t12 := int64(load64(m.memory[uint32(v4):]))
		v8 = (t12 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
		v9 = v2 + i32(32)
		v1 = i32(0)
	l6:
		if v3 == 0 {
			t27 := int32(load32(m.memory[int64(uint32(v0))+12:]))
			t28 := v2
			v1 = t27
			store32(m.memory[int64(uint32(t28))+44:], uint32(v1))
			store32(m.memory[int64(uint32(v2))+40:], uint32(v7-v1))
			m.fn239(v0, v9)
			m.fn240(v2 + i32(20))
			goto l7
		}
	l5:
		{
			if v8 != i64(0) {
				t14 := v6
				t15 := v6
				t16 := v5
				t17 := v2 + i32(16)
				t18 := v0
				v10 = int32(uint32(int64(bits.TrailingZeros64(uint64(v8))))>>3) + v1
				t19 := m.fn724(t17, t18, v10)
				v11 = t19
				t20 := m.fn26(t15, t16, v11)
				v12 = t20
				t21 := t14 + v12
				v13 = int32(uint32(int32(v11)) >> 25)
				m.memory[uint32(t21)] = byte(v13)
				m.memory[uint32(v6+v5&(v12+i32(-8))+i32(8))] = byte(v13)
				v12 = v6 + (v12^i32(-1))*i32(20)
				t22 := int32(load32(m.memory[uint32(v0):]))
				t23 := v12
				v10 = t22 + (v10^i32(-1))*i32(20)
				t24 := int32(load32(m.memory[int64(uint32(v10))+16:]))
				store32(m.memory[int64(uint32(t23))+16:], uint32(t24))
				t25 := int64(load64(m.memory[int64(uint32(v10))+8:]))
				store64(m.memory[int64(uint32(v12))+8:], uint64(t25))
				t26 := int64(load64(m.memory[uint32(v10):]))
				store64(m.memory[uint32(v12):], uint64(t26))
				v3 = v3 + i32(-1)
				v8 = (v8 + i64(-1)) & v8
				goto l6
			}
			v1 = v1 + i32(8)
			v4 = v4 + i32(8)
			t13 := int64(load64(m.memory[uint32(v4):]))
			v8 = (t13 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
			goto l5
		}
	}
l1:
	m.fn241(v0, v2+i32(16), i32(129), i32(20))
l7:
	v5 = i32(-1)
l2:
	m.g0 = v2 + i32(64)
	return v5
}
func (m *Module) fn724(v0, v1, v2 int32) int64 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(t0):]))
	v0 = t1
	t2 := int64(load64(m.memory[uint32(v0):]))
	t3 := int64(load64(m.memory[uint32(v0+i32(8)):]))
	t4 := int32(load32(m.memory[uint32(v1):]))
	t5 := int32(load32(m.memory[uint32(t4+(i32(0)-v2)*i32(20)+i32(-20)):]))
	t6 := m.fn66(t2, t3, t5)
	return t6
}
func (m *Module) fn725(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7 int32
	var v8 int64
	var v9, v10 int32
	var v11 int64
	var v12, v13 int32
	t0 := m.g0
	v2 = t0 - i32(64)
	m.g0 = v2
	store32(m.memory[int64(uint32(v2))+12:], uint32(v1))
	t1 := int32(load32(m.memory[int64(uint32(v0))+12:]))
	v3 = t1
	store32(m.memory[int64(uint32(v2))+16:], uint32(v2+i32(12)))
	v1 = v3 + i32(1)
	if v1 == 0 {
		m.fn242()
		panic("unreachable")
	}
	{
		t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		t3 := v1
		v4 = t2
		p4 := int32(uint32(v4+i32(1))>>3) * i32(7)
		if uint32(v4) < uint32(i32(8)) {
			p4 = v4
		}
		v4 = p4
		if uint32(t3) <= uint32(int32(uint32(v4)>>1)) {
			goto l1
		}
		t5 := v2 + i32(48)
		v4 = v4 + i32(1)
		p6 := v1
		if uint32(v4) > uint32(v1) {
			p6 = v4
		}
		m.fn237(t5, i32(4), p6)
		t7 := int32(load32(m.memory[int64(uint32(v2))+52:]))
		v5 = t7
		t8 := int32(load32(m.memory[int64(uint32(v2))+48:]))
		v6 = t8
		if v6 == 0 {
			goto l2
		}
		t9 := int32(load32(m.memory[int64(uint32(v2))+56:]))
		v7 = t9
		t10 := int32(load32(m.memory[int64(uint32(v2))+60:]))
		store32(m.memory[int64(uint32(v2))+44:], uint32(t10))
		store32(m.memory[int64(uint32(v2))+40:], uint32(v7))
		store32(m.memory[int64(uint32(v2))+36:], uint32(v5))
		store32(m.memory[int64(uint32(v2))+32:], uint32(v6))
		store64(m.memory[int64(uint32(v2))+24:], uint64(i64(0x800000004)))
		store32(m.memory[int64(uint32(v2))+20:], uint32(v0+i32(16)))
		t11 := int32(load32(m.memory[uint32(v0):]))
		v4 = t11
		t12 := int64(load64(m.memory[uint32(v4):]))
		v8 = (t12 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
		v9 = v2 + i32(32)
		v1 = i32(0)
	l6:
		if v3 == 0 {
			t24 := int32(load32(m.memory[int64(uint32(v0))+12:]))
			t25 := v2
			v1 = t24
			store32(m.memory[int64(uint32(t25))+44:], uint32(v1))
			store32(m.memory[int64(uint32(v2))+40:], uint32(v7-v1))
			m.fn239(v0, v9)
			m.fn240(v2 + i32(20))
			goto l7
		}
	l5:
		{
			if v8 != i64(0) {
				t14 := v6
				t15 := v6
				t16 := v5
				t17 := v2 + i32(16)
				t18 := v0
				v10 = int32(uint32(int64(bits.TrailingZeros64(uint64(v8))))>>3) + v1
				t19 := m.fn726(t17, t18, v10)
				v11 = t19
				t20 := m.fn26(t15, t16, v11)
				v12 = t20
				t21 := t14 + v12
				v13 = int32(uint32(int32(v11)) >> 25)
				m.memory[uint32(t21)] = byte(v13)
				m.memory[uint32(v6+v5&(v12+i32(-8))+i32(8))] = byte(v13)
				t22 := int32(load32(m.memory[uint32(v0):]))
				t23 := int32(load32(m.memory[uint32(t22+(v10^i32(-1))<<2):]))
				store32(m.memory[uint32(v6+(v12^i32(-1))<<2):], uint32(t23))
				v3 = v3 + i32(-1)
				v8 = (v8 + i64(-1)) & v8
				goto l6
			}
			v1 = v1 + i32(8)
			v4 = v4 + i32(8)
			t13 := int64(load64(m.memory[uint32(v4):]))
			v8 = (t13 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
			goto l5
		}
	}
l1:
	m.fn241(v0, v2+i32(16), i32(130), i32(4))
l7:
	v5 = i32(-1)
l2:
	m.g0 = v2 + i32(64)
	return v5
}
func (m *Module) fn726(v0, v1, v2 int32) int64 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(t0):]))
	v0 = t1
	t2 := int64(load64(m.memory[uint32(v0):]))
	t3 := int64(load64(m.memory[uint32(v0+i32(8)):]))
	t4 := int32(load32(m.memory[uint32(v1):]))
	t5 := int32(load32(m.memory[uint32(t4-v2<<2+i32(-4)):]))
	t6 := m.fn66(t2, t3, t5)
	return t6
}
func (m *Module) fn727(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7 int32
	var v8 int64
	var v9, v10 int32
	var v11 int64
	var v12, v13 int32
	t0 := m.g0
	v2 = t0 - i32(64)
	m.g0 = v2
	store32(m.memory[int64(uint32(v2))+12:], uint32(v1))
	t1 := int32(load32(m.memory[int64(uint32(v0))+12:]))
	v3 = t1
	store32(m.memory[int64(uint32(v2))+16:], uint32(v2+i32(12)))
	v1 = v3 + i32(1)
	if v1 == 0 {
		m.fn242()
		panic("unreachable")
	}
	{
		t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		t3 := v1
		v4 = t2
		p4 := int32(uint32(v4+i32(1))>>3) * i32(7)
		if uint32(v4) < uint32(i32(8)) {
			p4 = v4
		}
		v4 = p4
		if uint32(t3) <= uint32(int32(uint32(v4)>>1)) {
			goto l1
		}
		t5 := v2 + i32(48)
		v4 = v4 + i32(1)
		p6 := v1
		if uint32(v4) > uint32(v1) {
			p6 = v4
		}
		m.fn237(t5, i32(368), p6)
		t7 := int32(load32(m.memory[int64(uint32(v2))+52:]))
		v5 = t7
		t8 := int32(load32(m.memory[int64(uint32(v2))+48:]))
		v6 = t8
		if v6 == 0 {
			goto l2
		}
		t9 := int32(load32(m.memory[int64(uint32(v2))+56:]))
		v7 = t9
		t10 := int32(load32(m.memory[int64(uint32(v2))+60:]))
		store32(m.memory[int64(uint32(v2))+44:], uint32(t10))
		store32(m.memory[int64(uint32(v2))+40:], uint32(v7))
		store32(m.memory[int64(uint32(v2))+36:], uint32(v5))
		store32(m.memory[int64(uint32(v2))+32:], uint32(v6))
		store64(m.memory[int64(uint32(v2))+24:], uint64(i64(0x800000170)))
		store32(m.memory[int64(uint32(v2))+20:], uint32(v0+i32(16)))
		t11 := int32(load32(m.memory[uint32(v0):]))
		v4 = t11
		t12 := int64(load64(m.memory[uint32(v4):]))
		v8 = (t12 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
		v9 = v2 + i32(32)
		v1 = i32(0)
	l6:
		if v3 == 0 {
			t23 := int32(load32(m.memory[int64(uint32(v0))+12:]))
			t24 := v2
			v1 = t23
			store32(m.memory[int64(uint32(t24))+44:], uint32(v1))
			store32(m.memory[int64(uint32(v2))+40:], uint32(v7-v1))
			m.fn239(v0, v9)
			m.fn240(v2 + i32(20))
			goto l7
		}
	l5:
		{
			if v8 != i64(0) {
				t14 := v6
				t15 := v6
				t16 := v5
				t17 := v2 + i32(16)
				t18 := v0
				v10 = int32(uint32(int64(bits.TrailingZeros64(uint64(v8))))>>3) + v1
				t19 := m.fn728(t17, t18, v10)
				v11 = t19
				t20 := m.fn26(t15, t16, v11)
				v12 = t20
				t21 := t14 + v12
				v13 = int32(uint32(int32(v11)) >> 25)
				m.memory[uint32(t21)] = byte(v13)
				m.memory[uint32(v6+v5&(v12+i32(-8))+i32(8))] = byte(v13)
				t22 := int32(load32(m.memory[uint32(v0):]))
				memory_copy(m.memory, uint32(v6+(v12^i32(-1))*i32(368)), uint32(t22+(v10^i32(-1))*i32(368)), uint32(i32(368)))
				v3 = v3 + i32(-1)
				v8 = (v8 + i64(-1)) & v8
				goto l6
			}
			v1 = v1 + i32(8)
			v4 = v4 + i32(8)
			t13 := int64(load64(m.memory[uint32(v4):]))
			v8 = (t13 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
			goto l5
		}
	}
l1:
	m.fn241(v0, v2+i32(16), i32(131), i32(368))
l7:
	v5 = i32(-1)
l2:
	m.g0 = v2 + i32(64)
	return v5
}
func (m *Module) fn728(v0, v1, v2 int32) int64 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(t0):]))
	v0 = t1
	t2 := int64(load64(m.memory[uint32(v0):]))
	t3 := int64(load64(m.memory[uint32(v0+i32(8)):]))
	t4 := int32(load32(m.memory[uint32(v1):]))
	t5 := int32(load32(m.memory[uint32(t4+(i32(0)-v2)*i32(368)+i32(-368)):]))
	t6 := m.fn66(t2, t3, t5)
	return t6
}
func (m *Module) fn729(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7 int32
	var v8 int64
	var v9, v10 int32
	var v11 int64
	var v12, v13 int32
	t0 := m.g0
	v2 = t0 - i32(64)
	m.g0 = v2
	store32(m.memory[int64(uint32(v2))+12:], uint32(v1))
	t1 := int32(load32(m.memory[int64(uint32(v0))+12:]))
	v3 = t1
	store32(m.memory[int64(uint32(v2))+16:], uint32(v2+i32(12)))
	v1 = v3 + i32(1)
	if v1 == 0 {
		m.fn242()
		panic("unreachable")
	}
	{
		t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		t3 := v1
		v4 = t2
		p4 := int32(uint32(v4+i32(1))>>3) * i32(7)
		if uint32(v4) < uint32(i32(8)) {
			p4 = v4
		}
		v4 = p4
		if uint32(t3) <= uint32(int32(uint32(v4)>>1)) {
			goto l1
		}
		t5 := v2 + i32(48)
		v4 = v4 + i32(1)
		p6 := v1
		if uint32(v4) > uint32(v1) {
			p6 = v4
		}
		m.fn237(t5, i32(16), p6)
		t7 := int32(load32(m.memory[int64(uint32(v2))+52:]))
		v5 = t7
		t8 := int32(load32(m.memory[int64(uint32(v2))+48:]))
		v6 = t8
		if v6 == 0 {
			goto l2
		}
		t9 := int32(load32(m.memory[int64(uint32(v2))+56:]))
		v7 = t9
		t10 := int32(load32(m.memory[int64(uint32(v2))+60:]))
		store32(m.memory[int64(uint32(v2))+44:], uint32(t10))
		store32(m.memory[int64(uint32(v2))+40:], uint32(v7))
		store32(m.memory[int64(uint32(v2))+36:], uint32(v5))
		store32(m.memory[int64(uint32(v2))+32:], uint32(v6))
		store64(m.memory[int64(uint32(v2))+24:], uint64(i64(0x800000010)))
		store32(m.memory[int64(uint32(v2))+20:], uint32(v0+i32(16)))
		t11 := int32(load32(m.memory[uint32(v0):]))
		v4 = t11
		t12 := int64(load64(m.memory[uint32(v4):]))
		v8 = (t12 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
		v9 = v2 + i32(32)
		v1 = i32(0)
	l6:
		if v3 == 0 {
			t26 := int32(load32(m.memory[int64(uint32(v0))+12:]))
			t27 := v2
			v1 = t26
			store32(m.memory[int64(uint32(t27))+44:], uint32(v1))
			store32(m.memory[int64(uint32(v2))+40:], uint32(v7-v1))
			m.fn239(v0, v9)
			m.fn240(v2 + i32(20))
			goto l7
		}
	l5:
		{
			if v8 != i64(0) {
				t14 := v6
				t15 := v6
				t16 := v5
				t17 := v2 + i32(16)
				t18 := v0
				v10 = int32(uint32(int64(bits.TrailingZeros64(uint64(v8))))>>3) + v1
				t19 := m.fn730(t17, t18, v10)
				v11 = t19
				t20 := m.fn26(t15, t16, v11)
				v12 = t20
				t21 := t14 + v12
				v13 = int32(uint32(int32(v11)) >> 25)
				m.memory[uint32(t21)] = byte(v13)
				m.memory[uint32(v6+v5&(v12+i32(-8))+i32(8))] = byte(v13)
				v12 = v6 + (v12^i32(-1))<<4
				t22 := int32(load32(m.memory[uint32(v0):]))
				t23 := v12
				v10 = t22 + (v10^i32(-1))<<4
				t24 := int64(load64(m.memory[int64(uint32(v10))+8:]))
				store64(m.memory[int64(uint32(t23))+8:], uint64(t24))
				t25 := int64(load64(m.memory[uint32(v10):]))
				store64(m.memory[uint32(v12):], uint64(t25))
				v3 = v3 + i32(-1)
				v8 = (v8 + i64(-1)) & v8
				goto l6
			}
			v1 = v1 + i32(8)
			v4 = v4 + i32(8)
			t13 := int64(load64(m.memory[uint32(v4):]))
			v8 = (t13 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
			goto l5
		}
	}
l1:
	m.fn241(v0, v2+i32(16), i32(132), i32(16))
l7:
	v5 = i32(-1)
l2:
	m.g0 = v2 + i32(64)
	return v5
}
func (m *Module) fn730(v0, v1, v2 int32) int64 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(t0):]))
	v0 = t1
	t2 := int64(load64(m.memory[uint32(v0):]))
	t3 := int64(load64(m.memory[uint32(v0+i32(8)):]))
	t4 := int32(load32(m.memory[uint32(v1):]))
	t5 := int32(load16(m.memory[uint32(t4-v2<<4+i32(-16)):]))
	t6 := m.fn529(t2, t3, t5)
	return t6
}
func (m *Module) fn731(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7 int32
	var v8 int64
	var v9, v10 int32
	var v11 int64
	var v12, v13 int32
	t0 := m.g0
	v2 = t0 - i32(64)
	m.g0 = v2
	store32(m.memory[int64(uint32(v2))+12:], uint32(v1))
	t1 := int32(load32(m.memory[int64(uint32(v0))+12:]))
	v3 = t1
	store32(m.memory[int64(uint32(v2))+16:], uint32(v2+i32(12)))
	v1 = v3 + i32(1)
	if v1 == 0 {
		m.fn242()
		panic("unreachable")
	}
	{
		t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		t3 := v1
		v4 = t2
		p4 := int32(uint32(v4+i32(1))>>3) * i32(7)
		if uint32(v4) < uint32(i32(8)) {
			p4 = v4
		}
		v4 = p4
		if uint32(t3) <= uint32(int32(uint32(v4)>>1)) {
			goto l1
		}
		t5 := v2 + i32(48)
		v4 = v4 + i32(1)
		p6 := v1
		if uint32(v4) > uint32(v1) {
			p6 = v4
		}
		m.fn237(t5, i32(60), p6)
		t7 := int32(load32(m.memory[int64(uint32(v2))+52:]))
		v5 = t7
		t8 := int32(load32(m.memory[int64(uint32(v2))+48:]))
		v6 = t8
		if v6 == 0 {
			goto l2
		}
		t9 := int32(load32(m.memory[int64(uint32(v2))+56:]))
		v7 = t9
		t10 := int32(load32(m.memory[int64(uint32(v2))+60:]))
		store32(m.memory[int64(uint32(v2))+44:], uint32(t10))
		store32(m.memory[int64(uint32(v2))+40:], uint32(v7))
		store32(m.memory[int64(uint32(v2))+36:], uint32(v5))
		store32(m.memory[int64(uint32(v2))+32:], uint32(v6))
		store64(m.memory[int64(uint32(v2))+24:], uint64(i64(0x80000003c)))
		store32(m.memory[int64(uint32(v2))+20:], uint32(v0+i32(16)))
		t11 := int32(load32(m.memory[uint32(v0):]))
		v4 = t11
		t12 := int64(load64(m.memory[uint32(v4):]))
		v8 = (t12 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
		v9 = v2 + i32(32)
		v1 = i32(0)
	l6:
		if v3 == 0 {
			t23 := int32(load32(m.memory[int64(uint32(v0))+12:]))
			t24 := v2
			v1 = t23
			store32(m.memory[int64(uint32(t24))+44:], uint32(v1))
			store32(m.memory[int64(uint32(v2))+40:], uint32(v7-v1))
			m.fn239(v0, v9)
			m.fn240(v2 + i32(20))
			goto l7
		}
	l5:
		{
			if v8 != i64(0) {
				t14 := v6
				t15 := v6
				t16 := v5
				t17 := v2 + i32(16)
				t18 := v0
				v10 = int32(uint32(int64(bits.TrailingZeros64(uint64(v8))))>>3) + v1
				t19 := m.fn732(t17, t18, v10)
				v11 = t19
				t20 := m.fn26(t15, t16, v11)
				v12 = t20
				t21 := t14 + v12
				v13 = int32(uint32(int32(v11)) >> 25)
				m.memory[uint32(t21)] = byte(v13)
				m.memory[uint32(v6+v5&(v12+i32(-8))+i32(8))] = byte(v13)
				t22 := int32(load32(m.memory[uint32(v0):]))
				memory_copy(m.memory, uint32(v6+(v12^i32(-1))*i32(60)), uint32(t22+(v10^i32(-1))*i32(60)), uint32(i32(60)))
				v3 = v3 + i32(-1)
				v8 = (v8 + i64(-1)) & v8
				goto l6
			}
			v1 = v1 + i32(8)
			v4 = v4 + i32(8)
			t13 := int64(load64(m.memory[uint32(v4):]))
			v8 = (t13 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
			goto l5
		}
	}
l1:
	m.fn241(v0, v2+i32(16), i32(133), i32(60))
l7:
	v5 = i32(-1)
l2:
	m.g0 = v2 + i32(64)
	return v5
}
func (m *Module) fn732(v0, v1, v2 int32) int64 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(t0):]))
	v0 = t1
	t2 := int64(load64(m.memory[uint32(v0):]))
	t3 := int64(load64(m.memory[uint32(v0+i32(8)):]))
	t4 := int32(load32(m.memory[uint32(v1):]))
	t5 := int32(load16(m.memory[uint32(t4+(i32(0)-v2)*i32(60)+i32(-60)):]))
	t6 := m.fn529(t2, t3, t5)
	return t6
}
func (m *Module) fn733(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7 int32
	var v8 int64
	var v9, v10 int32
	var v11 int64
	var v12, v13 int32
	t0 := m.g0
	v2 = t0 - i32(64)
	m.g0 = v2
	store32(m.memory[int64(uint32(v2))+12:], uint32(v1))
	t1 := int32(load32(m.memory[int64(uint32(v0))+12:]))
	v3 = t1
	store32(m.memory[int64(uint32(v2))+16:], uint32(v2+i32(12)))
	v1 = v3 + i32(1)
	if v1 == 0 {
		m.fn242()
		panic("unreachable")
	}
	{
		t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		t3 := v1
		v4 = t2
		p4 := int32(uint32(v4+i32(1))>>3) * i32(7)
		if uint32(v4) < uint32(i32(8)) {
			p4 = v4
		}
		v4 = p4
		if uint32(t3) <= uint32(int32(uint32(v4)>>1)) {
			goto l1
		}
		t5 := v2 + i32(48)
		v4 = v4 + i32(1)
		p6 := v1
		if uint32(v4) > uint32(v1) {
			p6 = v4
		}
		m.fn237(t5, i32(36), p6)
		t7 := int32(load32(m.memory[int64(uint32(v2))+52:]))
		v5 = t7
		t8 := int32(load32(m.memory[int64(uint32(v2))+48:]))
		v6 = t8
		if v6 == 0 {
			goto l2
		}
		t9 := int32(load32(m.memory[int64(uint32(v2))+56:]))
		v7 = t9
		t10 := int32(load32(m.memory[int64(uint32(v2))+60:]))
		store32(m.memory[int64(uint32(v2))+44:], uint32(t10))
		store32(m.memory[int64(uint32(v2))+40:], uint32(v7))
		store32(m.memory[int64(uint32(v2))+36:], uint32(v5))
		store32(m.memory[int64(uint32(v2))+32:], uint32(v6))
		store64(m.memory[int64(uint32(v2))+24:], uint64(i64(0x800000024)))
		store32(m.memory[int64(uint32(v2))+20:], uint32(v0+i32(16)))
		t11 := int32(load32(m.memory[uint32(v0):]))
		v4 = t11
		t12 := int64(load64(m.memory[uint32(v4):]))
		v8 = (t12 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
		v9 = v2 + i32(32)
		v1 = i32(0)
	l6:
		if v3 == 0 {
			t23 := int32(load32(m.memory[int64(uint32(v0))+12:]))
			t24 := v2
			v1 = t23
			store32(m.memory[int64(uint32(t24))+44:], uint32(v1))
			store32(m.memory[int64(uint32(v2))+40:], uint32(v7-v1))
			m.fn239(v0, v9)
			m.fn240(v2 + i32(20))
			goto l7
		}
	l5:
		{
			if v8 != i64(0) {
				t14 := v6
				t15 := v6
				t16 := v5
				t17 := v2 + i32(16)
				t18 := v0
				v10 = int32(uint32(int64(bits.TrailingZeros64(uint64(v8))))>>3) + v1
				t19 := m.fn734(t17, t18, v10)
				v11 = t19
				t20 := m.fn26(t15, t16, v11)
				v12 = t20
				t21 := t14 + v12
				v13 = int32(uint32(int32(v11)) >> 25)
				m.memory[uint32(t21)] = byte(v13)
				m.memory[uint32(v6+v5&(v12+i32(-8))+i32(8))] = byte(v13)
				t22 := int32(load32(m.memory[uint32(v0):]))
				memory_copy(m.memory, uint32(v6+(v12^i32(-1))*i32(36)), uint32(t22+(v10^i32(-1))*i32(36)), uint32(i32(36)))
				v3 = v3 + i32(-1)
				v8 = (v8 + i64(-1)) & v8
				goto l6
			}
			v1 = v1 + i32(8)
			v4 = v4 + i32(8)
			t13 := int64(load64(m.memory[uint32(v4):]))
			v8 = (t13 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
			goto l5
		}
	}
l1:
	m.fn241(v0, v2+i32(16), i32(134), i32(36))
l7:
	v5 = i32(-1)
l2:
	m.g0 = v2 + i32(64)
	return v5
}
func (m *Module) fn734(v0, v1, v2 int32) int64 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(t0):]))
	v0 = t1
	t2 := int64(load64(m.memory[uint32(v0):]))
	t3 := int64(load64(m.memory[uint32(v0+i32(8)):]))
	t4 := int32(load32(m.memory[uint32(v1):]))
	t5 := int32(load16(m.memory[uint32(t4+(i32(0)-v2)*i32(36)+i32(-36)):]))
	t6 := m.fn529(t2, t3, t5)
	return t6
}
func (m *Module) fn735(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7 int32
	var v8 int64
	var v9, v10 int32
	var v11 int64
	var v12, v13 int32
	t0 := m.g0
	v2 = t0 - i32(64)
	m.g0 = v2
	store32(m.memory[int64(uint32(v2))+12:], uint32(v1))
	t1 := int32(load32(m.memory[int64(uint32(v0))+12:]))
	v3 = t1
	store32(m.memory[int64(uint32(v2))+16:], uint32(v2+i32(12)))
	v1 = v3 + i32(1)
	if v1 == 0 {
		m.fn242()
		panic("unreachable")
	}
	{
		t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		t3 := v1
		v4 = t2
		p4 := int32(uint32(v4+i32(1))>>3) * i32(7)
		if uint32(v4) < uint32(i32(8)) {
			p4 = v4
		}
		v4 = p4
		if uint32(t3) <= uint32(int32(uint32(v4)>>1)) {
			goto l1
		}
		t5 := v2 + i32(48)
		v4 = v4 + i32(1)
		p6 := v1
		if uint32(v4) > uint32(v1) {
			p6 = v4
		}
		m.fn237(t5, i32(520), p6)
		t7 := int32(load32(m.memory[int64(uint32(v2))+52:]))
		v5 = t7
		t8 := int32(load32(m.memory[int64(uint32(v2))+48:]))
		v6 = t8
		if v6 == 0 {
			goto l2
		}
		t9 := int32(load32(m.memory[int64(uint32(v2))+56:]))
		v7 = t9
		t10 := int32(load32(m.memory[int64(uint32(v2))+60:]))
		store32(m.memory[int64(uint32(v2))+44:], uint32(t10))
		store32(m.memory[int64(uint32(v2))+40:], uint32(v7))
		store32(m.memory[int64(uint32(v2))+36:], uint32(v5))
		store32(m.memory[int64(uint32(v2))+32:], uint32(v6))
		store64(m.memory[int64(uint32(v2))+24:], uint64(i64(0x800000208)))
		store32(m.memory[int64(uint32(v2))+20:], uint32(v0+i32(16)))
		t11 := int32(load32(m.memory[uint32(v0):]))
		v4 = t11
		t12 := int64(load64(m.memory[uint32(v4):]))
		v8 = (t12 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
		v9 = v2 + i32(32)
		v1 = i32(0)
	l6:
		if v3 == 0 {
			t23 := int32(load32(m.memory[int64(uint32(v0))+12:]))
			t24 := v2
			v1 = t23
			store32(m.memory[int64(uint32(t24))+44:], uint32(v1))
			store32(m.memory[int64(uint32(v2))+40:], uint32(v7-v1))
			m.fn239(v0, v9)
			m.fn240(v2 + i32(20))
			goto l7
		}
	l5:
		{
			if v8 != i64(0) {
				t14 := v6
				t15 := v6
				t16 := v5
				t17 := v2 + i32(16)
				t18 := v0
				v10 = int32(uint32(int64(bits.TrailingZeros64(uint64(v8))))>>3) + v1
				t19 := m.fn736(t17, t18, v10)
				v11 = t19
				t20 := m.fn26(t15, t16, v11)
				v12 = t20
				t21 := t14 + v12
				v13 = int32(uint32(int32(v11)) >> 25)
				m.memory[uint32(t21)] = byte(v13)
				m.memory[uint32(v6+v5&(v12+i32(-8))+i32(8))] = byte(v13)
				t22 := int32(load32(m.memory[uint32(v0):]))
				memory_copy(m.memory, uint32(v6+(v12^i32(-1))*i32(520)), uint32(t22+(v10^i32(-1))*i32(520)), uint32(i32(520)))
				v3 = v3 + i32(-1)
				v8 = (v8 + i64(-1)) & v8
				goto l6
			}
			v1 = v1 + i32(8)
			v4 = v4 + i32(8)
			t13 := int64(load64(m.memory[uint32(v4):]))
			v8 = (t13 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
			goto l5
		}
	}
l1:
	m.fn241(v0, v2+i32(16), i32(135), i32(520))
l7:
	v5 = i32(-1)
l2:
	m.g0 = v2 + i32(64)
	return v5
}
func (m *Module) fn736(v0, v1, v2 int32) int64 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(t0):]))
	v0 = t1
	t2 := int64(load64(m.memory[uint32(v0):]))
	t3 := int64(load64(m.memory[uint32(v0+i32(8)):]))
	t4 := int32(load32(m.memory[uint32(v1):]))
	t5 := int32(load16(m.memory[uint32(t4+(i32(0)-v2)*i32(520)+i32(-520)):]))
	t6 := m.fn529(t2, t3, t5)
	return t6
}
func (m *Module) fn737(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7 int32
	var v8 int64
	var v9, v10 int32
	var v11 int64
	var v12, v13 int32
	t0 := m.g0
	v2 = t0 - i32(64)
	m.g0 = v2
	store32(m.memory[int64(uint32(v2))+12:], uint32(v1))
	t1 := int32(load32(m.memory[int64(uint32(v0))+12:]))
	v3 = t1
	store32(m.memory[int64(uint32(v2))+16:], uint32(v2+i32(12)))
	v1 = v3 + i32(1)
	if v1 == 0 {
		m.fn242()
		panic("unreachable")
	}
	{
		t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		t3 := v1
		v4 = t2
		p4 := int32(uint32(v4+i32(1))>>3) * i32(7)
		if uint32(v4) < uint32(i32(8)) {
			p4 = v4
		}
		v4 = p4
		if uint32(t3) <= uint32(int32(uint32(v4)>>1)) {
			goto l1
		}
		t5 := v2 + i32(48)
		v4 = v4 + i32(1)
		p6 := v1
		if uint32(v4) > uint32(v1) {
			p6 = v4
		}
		m.fn237(t5, i32(2), p6)
		t7 := int32(load32(m.memory[int64(uint32(v2))+52:]))
		v5 = t7
		t8 := int32(load32(m.memory[int64(uint32(v2))+48:]))
		v6 = t8
		if v6 == 0 {
			goto l2
		}
		t9 := int32(load32(m.memory[int64(uint32(v2))+56:]))
		v7 = t9
		t10 := int32(load32(m.memory[int64(uint32(v2))+60:]))
		store32(m.memory[int64(uint32(v2))+44:], uint32(t10))
		store32(m.memory[int64(uint32(v2))+40:], uint32(v7))
		store32(m.memory[int64(uint32(v2))+36:], uint32(v5))
		store32(m.memory[int64(uint32(v2))+32:], uint32(v6))
		store64(m.memory[int64(uint32(v2))+24:], uint64(i64(0x800000002)))
		store32(m.memory[int64(uint32(v2))+20:], uint32(v0+i32(16)))
		t11 := int32(load32(m.memory[uint32(v0):]))
		v4 = t11
		t12 := int64(load64(m.memory[uint32(v4):]))
		v8 = (t12 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
		v9 = v2 + i32(32)
		v1 = i32(0)
	l6:
		if v3 == 0 {
			t24 := int32(load32(m.memory[int64(uint32(v0))+12:]))
			t25 := v2
			v1 = t24
			store32(m.memory[int64(uint32(t25))+44:], uint32(v1))
			store32(m.memory[int64(uint32(v2))+40:], uint32(v7-v1))
			m.fn239(v0, v9)
			m.fn240(v2 + i32(20))
			goto l7
		}
	l5:
		{
			if v8 != i64(0) {
				t14 := v6
				t15 := v6
				t16 := v5
				t17 := v2 + i32(16)
				t18 := v0
				v10 = int32(uint32(int64(bits.TrailingZeros64(uint64(v8))))>>3) + v1
				t19 := m.fn738(t17, t18, v10)
				v11 = t19
				t20 := m.fn26(t15, t16, v11)
				v12 = t20
				t21 := t14 + v12
				v13 = int32(uint32(int32(v11)) >> 25)
				m.memory[uint32(t21)] = byte(v13)
				m.memory[uint32(v6+v5&(v12+i32(-8))+i32(8))] = byte(v13)
				t22 := int32(load32(m.memory[uint32(v0):]))
				t23 := int32(load16(m.memory[uint32(t22+(v10^i32(-1))<<1):]))
				store16(m.memory[uint32(v6+(v12^i32(-1))<<1):], uint16(t23))
				v3 = v3 + i32(-1)
				v8 = (v8 + i64(-1)) & v8
				goto l6
			}
			v1 = v1 + i32(8)
			v4 = v4 + i32(8)
			t13 := int64(load64(m.memory[uint32(v4):]))
			v8 = (t13 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
			goto l5
		}
	}
l1:
	m.fn241(v0, v2+i32(16), i32(136), i32(2))
l7:
	v5 = i32(-1)
l2:
	m.g0 = v2 + i32(64)
	return v5
}
func (m *Module) fn738(v0, v1, v2 int32) int64 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(t0):]))
	v0 = t1
	t2 := int64(load64(m.memory[uint32(v0):]))
	t3 := int64(load64(m.memory[uint32(v0+i32(8)):]))
	t4 := int32(load32(m.memory[uint32(v1):]))
	t5 := int32(load16(m.memory[uint32(t4-v2<<1+i32(-2)):]))
	t6 := m.fn529(t2, t3, t5)
	return t6
}
func (m *Module) fn739(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7 int32
	var v8 int64
	var v9, v10 int32
	var v11 int64
	var v12, v13 int32
	t0 := m.g0
	v2 = t0 - i32(64)
	m.g0 = v2
	store32(m.memory[int64(uint32(v2))+12:], uint32(v1))
	t1 := int32(load32(m.memory[int64(uint32(v0))+12:]))
	v3 = t1
	store32(m.memory[int64(uint32(v2))+16:], uint32(v2+i32(12)))
	v1 = v3 + i32(1)
	if v1 == 0 {
		m.fn242()
		panic("unreachable")
	}
	{
		t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		t3 := v1
		v4 = t2
		p4 := int32(uint32(v4+i32(1))>>3) * i32(7)
		if uint32(v4) < uint32(i32(8)) {
			p4 = v4
		}
		v4 = p4
		if uint32(t3) <= uint32(int32(uint32(v4)>>1)) {
			goto l1
		}
		t5 := v2 + i32(48)
		v4 = v4 + i32(1)
		p6 := v1
		if uint32(v4) > uint32(v1) {
			p6 = v4
		}
		m.fn237(t5, i32(104), p6)
		t7 := int32(load32(m.memory[int64(uint32(v2))+52:]))
		v5 = t7
		t8 := int32(load32(m.memory[int64(uint32(v2))+48:]))
		v6 = t8
		if v6 == 0 {
			goto l2
		}
		t9 := int32(load32(m.memory[int64(uint32(v2))+56:]))
		v7 = t9
		t10 := int32(load32(m.memory[int64(uint32(v2))+60:]))
		store32(m.memory[int64(uint32(v2))+44:], uint32(t10))
		store32(m.memory[int64(uint32(v2))+40:], uint32(v7))
		store32(m.memory[int64(uint32(v2))+36:], uint32(v5))
		store32(m.memory[int64(uint32(v2))+32:], uint32(v6))
		store64(m.memory[int64(uint32(v2))+24:], uint64(i64(0x800000068)))
		store32(m.memory[int64(uint32(v2))+20:], uint32(v0+i32(16)))
		t11 := int32(load32(m.memory[uint32(v0):]))
		v4 = t11
		t12 := int64(load64(m.memory[uint32(v4):]))
		v8 = (t12 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
		v9 = v2 + i32(32)
		v1 = i32(0)
	l6:
		if v3 == 0 {
			t23 := int32(load32(m.memory[int64(uint32(v0))+12:]))
			t24 := v2
			v1 = t23
			store32(m.memory[int64(uint32(t24))+44:], uint32(v1))
			store32(m.memory[int64(uint32(v2))+40:], uint32(v7-v1))
			m.fn239(v0, v9)
			m.fn240(v2 + i32(20))
			goto l7
		}
	l5:
		{
			if v8 != i64(0) {
				t14 := v6
				t15 := v6
				t16 := v5
				t17 := v2 + i32(16)
				t18 := v0
				v10 = int32(uint32(int64(bits.TrailingZeros64(uint64(v8))))>>3) + v1
				t19 := m.fn740(t17, t18, v10)
				v11 = t19
				t20 := m.fn26(t15, t16, v11)
				v12 = t20
				t21 := t14 + v12
				v13 = int32(uint32(int32(v11)) >> 25)
				m.memory[uint32(t21)] = byte(v13)
				m.memory[uint32(v6+v5&(v12+i32(-8))+i32(8))] = byte(v13)
				t22 := int32(load32(m.memory[uint32(v0):]))
				memory_copy(m.memory, uint32(v6+(v12^i32(-1))*i32(104)), uint32(t22+(v10^i32(-1))*i32(104)), uint32(i32(104)))
				v3 = v3 + i32(-1)
				v8 = (v8 + i64(-1)) & v8
				goto l6
			}
			v1 = v1 + i32(8)
			v4 = v4 + i32(8)
			t13 := int64(load64(m.memory[uint32(v4):]))
			v8 = (t13 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
			goto l5
		}
	}
l1:
	m.fn241(v0, v2+i32(16), i32(137), i32(104))
l7:
	v5 = i32(-1)
l2:
	m.g0 = v2 + i32(64)
	return v5
}
func (m *Module) fn740(v0, v1, v2 int32) int64 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(t0):]))
	v0 = t1
	t2 := int64(load64(m.memory[uint32(v0):]))
	t3 := int64(load64(m.memory[uint32(v0+i32(8)):]))
	t4 := int32(load32(m.memory[uint32(v1):]))
	t5 := int64(load64(m.memory[uint32(t4+(i32(0)-v2)*i32(104)+i32(-104)):]))
	t6 := m.fn741(t2, t3, t5)
	return t6
}
func (m *Module) fn741(v0, v1, v2 int64) int64 {
	var v3 int32
	t0 := m.g0
	v3 = t0 - i32(80)
	m.g0 = v3
	store64(m.memory[int64(uint32(v3))+56:], uint64(i64(0)))
	store64(m.memory[int64(uint32(v3))+64:], uint64(i64(0)))
	store64(m.memory[int64(uint32(v3))+48:], uint64(v1))
	store64(m.memory[int64(uint32(v3))+32:], uint64(v1^i64(8387220255154660723)))
	store64(m.memory[int64(uint32(v3))+24:], uint64(v1^i64(7237128888997146477)))
	store64(m.memory[int64(uint32(v3))+40:], uint64(v0))
	store64(m.memory[int64(uint32(v3))+16:], uint64(v0^i64(0x6c7967656e657261)))
	store64(m.memory[int64(uint32(v3))+8:], uint64(v0^i64(8317987319222330741)))
	store64(m.memory[int64(uint32(v3))+72:], uint64(v2))
	m.fn285(v3+i32(8), v3+i32(72), i32(8))
	t1 := m.fn174(v3 + i32(8))
	v1 = t1
	m.g0 = v3 + i32(80)
	return v1
}
func (m *Module) fn742(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7 int32
	var v8 int64
	var v9, v10 int32
	var v11 int64
	var v12, v13 int32
	t0 := m.g0
	v2 = t0 - i32(64)
	m.g0 = v2
	store32(m.memory[int64(uint32(v2))+12:], uint32(v1))
	t1 := int32(load32(m.memory[int64(uint32(v0))+12:]))
	v3 = t1
	store32(m.memory[int64(uint32(v2))+16:], uint32(v2+i32(12)))
	v1 = v3 + i32(1)
	if v1 == 0 {
		m.fn242()
		panic("unreachable")
	}
	{
		t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		t3 := v1
		v4 = t2
		p4 := int32(uint32(v4+i32(1))>>3) * i32(7)
		if uint32(v4) < uint32(i32(8)) {
			p4 = v4
		}
		v4 = p4
		if uint32(t3) <= uint32(int32(uint32(v4)>>1)) {
			goto l1
		}
		t5 := v2 + i32(48)
		v4 = v4 + i32(1)
		p6 := v1
		if uint32(v4) > uint32(v1) {
			p6 = v4
		}
		m.fn237(t5, i32(480), p6)
		t7 := int32(load32(m.memory[int64(uint32(v2))+52:]))
		v5 = t7
		t8 := int32(load32(m.memory[int64(uint32(v2))+48:]))
		v6 = t8
		if v6 == 0 {
			goto l2
		}
		t9 := int32(load32(m.memory[int64(uint32(v2))+56:]))
		v7 = t9
		t10 := int32(load32(m.memory[int64(uint32(v2))+60:]))
		store32(m.memory[int64(uint32(v2))+44:], uint32(t10))
		store32(m.memory[int64(uint32(v2))+40:], uint32(v7))
		store32(m.memory[int64(uint32(v2))+36:], uint32(v5))
		store32(m.memory[int64(uint32(v2))+32:], uint32(v6))
		store64(m.memory[int64(uint32(v2))+24:], uint64(i64(0x8000001e0)))
		store32(m.memory[int64(uint32(v2))+20:], uint32(v0+i32(16)))
		t11 := int32(load32(m.memory[uint32(v0):]))
		v4 = t11
		t12 := int64(load64(m.memory[uint32(v4):]))
		v8 = (t12 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
		v9 = v2 + i32(32)
		v1 = i32(0)
	l6:
		if v3 == 0 {
			t23 := int32(load32(m.memory[int64(uint32(v0))+12:]))
			t24 := v2
			v1 = t23
			store32(m.memory[int64(uint32(t24))+44:], uint32(v1))
			store32(m.memory[int64(uint32(v2))+40:], uint32(v7-v1))
			m.fn239(v0, v9)
			m.fn240(v2 + i32(20))
			goto l7
		}
	l5:
		{
			if v8 != i64(0) {
				t14 := v6
				t15 := v6
				t16 := v5
				t17 := v2 + i32(16)
				t18 := v0
				v10 = int32(uint32(int64(bits.TrailingZeros64(uint64(v8))))>>3) + v1
				t19 := m.fn743(t17, t18, v10)
				v11 = t19
				t20 := m.fn26(t15, t16, v11)
				v12 = t20
				t21 := t14 + v12
				v13 = int32(uint32(int32(v11)) >> 25)
				m.memory[uint32(t21)] = byte(v13)
				m.memory[uint32(v6+v5&(v12+i32(-8))+i32(8))] = byte(v13)
				t22 := int32(load32(m.memory[uint32(v0):]))
				memory_copy(m.memory, uint32(v6+(v12^i32(-1))*i32(480)), uint32(t22+(v10^i32(-1))*i32(480)), uint32(i32(480)))
				v3 = v3 + i32(-1)
				v8 = (v8 + i64(-1)) & v8
				goto l6
			}
			v1 = v1 + i32(8)
			v4 = v4 + i32(8)
			t13 := int64(load64(m.memory[uint32(v4):]))
			v8 = (t13 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
			goto l5
		}
	}
l1:
	m.fn241(v0, v2+i32(16), i32(138), i32(480))
l7:
	v5 = i32(-1)
l2:
	m.g0 = v2 + i32(64)
	return v5
}
func (m *Module) fn743(v0, v1, v2 int32) int64 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(t0):]))
	v0 = t1
	t2 := int64(load64(m.memory[uint32(v0):]))
	t3 := int64(load64(m.memory[uint32(v0+i32(8)):]))
	t4 := int32(load32(m.memory[uint32(v1):]))
	t5 := int64(load64(m.memory[uint32(t4+(i32(0)-v2)*i32(480)+i32(-480)):]))
	t6 := m.fn741(t2, t3, t5)
	return t6
}
func (m *Module) fn744(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7 int32
	var v8 int64
	var v9, v10 int32
	var v11 int64
	var v12, v13 int32
	t0 := m.g0
	v2 = t0 - i32(64)
	m.g0 = v2
	store32(m.memory[int64(uint32(v2))+12:], uint32(v1))
	t1 := int32(load32(m.memory[int64(uint32(v0))+12:]))
	v3 = t1
	store32(m.memory[int64(uint32(v2))+16:], uint32(v2+i32(12)))
	v1 = v3 + i32(1)
	if v1 == 0 {
		m.fn242()
		panic("unreachable")
	}
	{
		t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		t3 := v1
		v4 = t2
		p4 := int32(uint32(v4+i32(1))>>3) * i32(7)
		if uint32(v4) < uint32(i32(8)) {
			p4 = v4
		}
		v4 = p4
		if uint32(t3) <= uint32(int32(uint32(v4)>>1)) {
			goto l1
		}
		t5 := v2 + i32(48)
		v4 = v4 + i32(1)
		p6 := v1
		if uint32(v4) > uint32(v1) {
			p6 = v4
		}
		m.fn237(t5, i32(24), p6)
		t7 := int32(load32(m.memory[int64(uint32(v2))+52:]))
		v5 = t7
		t8 := int32(load32(m.memory[int64(uint32(v2))+48:]))
		v6 = t8
		if v6 == 0 {
			goto l2
		}
		t9 := int32(load32(m.memory[int64(uint32(v2))+56:]))
		v7 = t9
		t10 := int32(load32(m.memory[int64(uint32(v2))+60:]))
		store32(m.memory[int64(uint32(v2))+44:], uint32(t10))
		store32(m.memory[int64(uint32(v2))+40:], uint32(v7))
		store32(m.memory[int64(uint32(v2))+36:], uint32(v5))
		store32(m.memory[int64(uint32(v2))+32:], uint32(v6))
		store64(m.memory[int64(uint32(v2))+24:], uint64(i64(0x800000018)))
		store32(m.memory[int64(uint32(v2))+20:], uint32(v0+i32(16)))
		t11 := int32(load32(m.memory[uint32(v0):]))
		v4 = t11
		t12 := int64(load64(m.memory[uint32(v4):]))
		v8 = (t12 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
		v9 = v2 + i32(32)
		v1 = i32(0)
	l6:
		if v3 == 0 {
			t27 := int32(load32(m.memory[int64(uint32(v0))+12:]))
			t28 := v2
			v1 = t27
			store32(m.memory[int64(uint32(t28))+44:], uint32(v1))
			store32(m.memory[int64(uint32(v2))+40:], uint32(v7-v1))
			m.fn239(v0, v9)
			m.fn240(v2 + i32(20))
			goto l7
		}
	l5:
		{
			if v8 != i64(0) {
				t14 := v6
				t15 := v6
				t16 := v5
				t17 := v2 + i32(16)
				t18 := v0
				v10 = int32(uint32(int64(bits.TrailingZeros64(uint64(v8))))>>3) + v1
				t19 := m.fn745(t17, t18, v10)
				v11 = t19
				t20 := m.fn26(t15, t16, v11)
				v12 = t20
				t21 := t14 + v12
				v13 = int32(uint32(int32(v11)) >> 25)
				m.memory[uint32(t21)] = byte(v13)
				m.memory[uint32(v6+v5&(v12+i32(-8))+i32(8))] = byte(v13)
				v12 = v6 + (v12^i32(-1))*i32(24)
				t22 := int32(load32(m.memory[uint32(v0):]))
				t23 := v12
				v10 = t22 + (v10^i32(-1))*i32(24)
				t24 := int64(load64(m.memory[int64(uint32(v10))+16:]))
				store64(m.memory[int64(uint32(t23))+16:], uint64(t24))
				t25 := int64(load64(m.memory[int64(uint32(v10))+8:]))
				store64(m.memory[int64(uint32(v12))+8:], uint64(t25))
				t26 := int64(load64(m.memory[uint32(v10):]))
				store64(m.memory[uint32(v12):], uint64(t26))
				v3 = v3 + i32(-1)
				v8 = (v8 + i64(-1)) & v8
				goto l6
			}
			v1 = v1 + i32(8)
			v4 = v4 + i32(8)
			t13 := int64(load64(m.memory[uint32(v4):]))
			v8 = (t13 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
			goto l5
		}
	}
l1:
	m.fn241(v0, v2+i32(16), i32(139), i32(24))
l7:
	v5 = i32(-1)
l2:
	m.g0 = v2 + i32(64)
	return v5
}
func (m *Module) fn745(v0, v1, v2 int32) int64 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(t0):]))
	v0 = t1
	t2 := int64(load64(m.memory[uint32(v0):]))
	t3 := int64(load64(m.memory[uint32(v0+i32(8)):]))
	t4 := int32(load32(m.memory[uint32(v1):]))
	t5 := int64(load64(m.memory[uint32(t4+(i32(0)-v2)*i32(24)+i32(-24)):]))
	t6 := m.fn741(t2, t3, t5)
	return t6
}
func (m *Module) fn746(v0, v1, v2, v3 int32) {
	var v4, v5, v6, v7, v8, v9, v10 int32
	var v11 int64
	var v12, v13, v14, v15, v16, v17, v18, v19, v20, v21, v22, v23, v24, v25, v26, v27, v28, v29, v30, v31, v32 int32
	t0 := m.g0
	v4 = t0 - i32(880)
	m.g0 = v4
	t1 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	v5 = t1
	v6 = v4 + i32(794)
	memory_zero(m.memory, uint32(v6), uint32(i32(70)))
	v7 = v4 + i32(460)
	memory_zero(m.memory, uint32(v7), uint32(i32(256)))
	v8 = v4 + i32(448) + i32(272)
	memory_copy(m.memory, uint32(v8), uint32(v5+i32(272)), uint32(i32(70)))
	store32(m.memory[int64(uint32(v4))+790:], uint32(i32(0)))
	store32(m.memory[int64(uint32(v4))+716:], uint32(i32(1)))
	t2 := int32(m.memory[int64(uint32(v5))+419])
	m.memory[int64(uint32(v4))+867] = byte(t2)
	t3 := int32(m.memory[int64(uint32(v5))+416])
	m.memory[int64(uint32(v4))+864] = byte(t3)
	t4 := int32(m.memory[int64(uint32(v5))+417])
	t5 := v4
	v9 = t4
	m.memory[int64(uint32(t5))+865] = byte(v9)
	t6 := int32(load32(m.memory[int64(uint32(v5))+8:]))
	v10 = t6
	t7 := int64(load64(m.memory[uint32(v5):]))
	v11 = t7
	t8 := int32(m.memory[int64(uint32(v5))+418])
	v12 = t8
	t9 := int32(load16(m.memory[int64(uint32(v5))+420:]))
	store16(m.memory[int64(uint32(v4))+868:], uint16(t9))
	t10 := int32(m.memory[int64(uint32(v5))+429])
	v13 = t10
	t11 := int32(m.memory[int64(uint32(v5))+424])
	v14 = t11
	t12 := int32(m.memory[int64(uint32(v5))+425])
	v15 = t12
	t13 := int32(load16(m.memory[int64(uint32(v5))+427:]))
	v16 = t13
	t14 := int32(m.memory[int64(uint32(v5))+422])
	v17 = t14
	t15 := int32(m.memory[int64(uint32(v5))+423])
	v18 = t15
	t16 := int32(m.memory[int64(uint32(v5))+426])
	m.memory[int64(uint32(v4))+874] = byte(t16)
	m.memory[int64(uint32(v4))+871] = byte(v18)
	m.memory[int64(uint32(v4))+870] = byte(v17)
	m.memory[int64(uint32(v4))+866] = byte(v12)
	store16(m.memory[int64(uint32(v4))+875:], uint16(v16))
	m.memory[int64(uint32(v4))+873] = byte(v15)
	m.memory[int64(uint32(v4))+872] = byte(v14)
	m.memory[int64(uint32(v4))+877] = byte(v13)
	store64(m.memory[int64(uint32(v4))+448:], uint64(v11))
	store32(m.memory[int64(uint32(v4))+456:], uint32(v10))
	m.fn747(v7, v9)
	{
		t17 := int32(m.memory[int64(uint32(v4))+875])
		if t17 == 0 {
			goto l0
		}
		t18 := int32(m.memory[int64(uint32(v4))+866])
		m.fn747(v7, t18)
		t19 := int32(m.memory[int64(uint32(v4))+870])
		if t19 != i32(1) {
			goto l0
		}
		t20 := int32(m.memory[int64(uint32(v4))+871])
		m.fn747(v7, t20)
	}
l0:
	{
		t21 := int32(m.memory[int64(uint32(v4))+872])
		if t21 == 0 {
			goto l1
		}
		t22 := int32(m.memory[int64(uint32(v4))+873])
		m.fn747(v7, t22)
	}
l1:
	{
		t23 := int32(m.memory[int64(uint32(v4))+868])
		if t23 != i32(1) {
			goto l2
		}
		t24 := int32(m.memory[int64(uint32(v4))+869])
		v5 = t24
		goto l3
	}
l2:
	m.fn747(v7, i32(13))
	v5 = i32(10)
l3:
	m.fn747(v7, v5)
	v5 = i32(0)
	t25 := int32(load32(m.memory[int64(uint32(v4))+716:]))
	v19 = t25
l6:
	{
		if v5 == i32(10) {
			goto l4
		}
		v20 = v5 + i32(1)
		t26 := int32(m.memory[int64(uint32(v5))+1148604])
		v21 = t26
		v12 = i32(0)
	l10:
		if v12 != i32(256) {
			t27 := int32(m.memory[int64(uint32(v4))+869])
			v5 = v12 & i32(255)
			var p28 int32
			if t27 == v5 {
				p28 = 1
			}
			v22 = p28
			t29 := v22
			var p30 int32
			if v12 == i32(13) {
				p30 = 1
			}
			v23 = p30
			t31 := v23
			var p32 int32
			if v12 == i32(10) {
				p32 = 1
			}
			v10 = p32
			t33 := int32(m.memory[int64(uint32(v4))+868])
			t34 := t31 | v10
			v13 = t33
			p35 := t34
			if v13 != 0 {
				p35 = t29
			}
			v15 = p35
			p36 := i32(2)
			if v15 != 0 {
				p36 = i32(-56)
			}
			t37 := int32(m.memory[int64(uint32(v4))+865])
			var p38 int32
			if t37&i32(255) == v5 {
				p38 = 1
			}
			v17 = p38
			p39 := p36
			if v17 != 0 {
				p39 = i32(7)
			}
			v14 = p39
			t40 := int32(m.memory[int64(uint32(v4))+875])
			t41 := v14
			v9 = t40
			t42 := int32(m.memory[int64(uint32(v4))+866])
			t43 := v9
			v24 = t42 & i32(255)
			var p44 int32
			if v24 == v5 {
				p44 = 1
			}
			v18 = t43 & p44
			p45 := t41
			if v18 != 0 {
				p45 = i32(3)
			}
			v16 = p45
			p46 := v15 ^ i32(1)
			if v17 != 0 {
				p46 = i32(2)
			}
			v15 = p46
			p47 := v15
			if v18 != 0 {
				p47 = i32(2)
			}
			v17 = p47
			t48 := int32(m.memory[int64(uint32(v4))+874])
			t49 := v14
			v25 = v18 & t48
			p50 := t49
			if v25 != 0 {
				p50 = i32(3)
			}
			v18 = p50
			p51 := v15
			if v25 != 0 {
				p51 = i32(1)
			}
			v25 = p51
			t52 := int32(m.memory[int64(uint32(v4))+870])
			t53 := int32(m.memory[int64(uint32(v4))+871])
			var p54 int32
			if t53&i32(255) == v5 {
				p54 = 1
			}
			v26 = t52 & p54
			p55 := i32(3)
			if v26 != 0 {
				p55 = i32(4)
			}
			var p56 int32
			if v24 != v5 {
				p56 = 1
			}
			v27 = p56
			p57 := i32(5)
			if v27 != 0 {
				p57 = p55
			}
			p58 := i32(3)
			if v9 != 0 {
				p58 = p57
			}
			v24 = p58
			p59 := i32(1)
			if v26 != 0 {
				p59 = i32(2)
			}
			v26 = v9 + i32(1)
			p60 := v26
			if v27 != 0 {
				p60 = p59
			}
			p61 := v26
			if v9 != 0 {
				p61 = p60
			}
			v26 = p61
			p62 := i32(9)
			if v13 != 0 {
				p62 = i32(8)
			}
			p63 := i32(8)
			if v23 != 0 {
				p63 = p62
			}
			v28 = p63
			t64 := int32(m.memory[int64(uint32(v4))+872])
			t65 := int32(m.memory[int64(uint32(v4))+873])
			var p66 int32
			if t65&i32(255) == v5 {
				p66 = 1
			}
			v5 = t64 & p66
			p67 := i32(1)
			if v5 != 0 {
				p67 = i32(6)
			}
			v29 = p67
			v9 = i32(0)
			p68 := i32(0)
			if v5 != 0 {
				p68 = i32(2)
			}
			v30 = p68
			p69 := i32(6)
			if v10 != 0 {
				p69 = i32(0)
			}
			v23 = p69
			p70 := i32(0)
			if v10 != 0 {
				p70 = i32(2)
			}
			v27 = p70
			v31 = v12 + i32(1)
			v32 = v12 + i32(-10)
			v5 = v21
		l22:
			{
				v10 = v5 & i32(255)
				if v10 == i32(202) {
					goto l7
				}
				if v9&i32(255) == 0 {
					switch v10 {
					default:
						v5 = i32(201)
						v9 = i32(0)
						switch v10 + i32(-200) {
						default:
							goto l22
						case 1:
							v9 = i32(2)
							v5 = v28
							goto l22
						}
					case 0:
						if v13 == 0 {
							v9 = i32(2)
							v5 = i32(0)
							switch v32 {
							case 0, 3:
								goto l22
							default:
								goto l25
							}
						}
						v9 = i32(2)
						v5 = i32(0)
						if v22 != 0 {
							goto l22
						}
						goto l25
					case 1:
						v5 = v16
						v9 = v17
						goto l22
					case 2:
						v5 = v14
						v9 = v15
						goto l22
					case 3:
						v5 = v24
						v9 = v26
						goto l22
					case 4:
						v9 = i32(1)
						v5 = i32(3)
						goto l22
					case 5:
						v5 = v18
						v9 = v25
						goto l22
					case 6:
						v9 = i32(2)
						v5 = v23
						goto l22
					case 7:
						v5 = i32(1)
						v9 = i32(0)
						goto l22
					case 8:
						v5 = i32(0)
						v9 = i32(0)
						goto l22
					case 9:
						v5 = i32(0)
						v9 = v27
						goto l22
					}
				}
			l7:
				t71 := m.fn748(v19, v21)
				v10 = t71
				t72 := m.fn748(v19, v5)
				v13 = t72
				t73 := int32(m.memory[uint32(v7+v12)])
				v5 = t73 + v10&i32(255)
				if uint32(v5) >= uint32(i32(70)) {
					m.fn158(v5, i32(70), i32(1148632))
					panic("unreachable")
				}
				m.memory[uint32(v8+v5)] = byte(v13)
				t74 := v6 + v5
				var p75 int32
				if v9&i32(255) == i32(1) {
					p75 = 1
				}
				m.memory[uint32(t74)] = byte(p75)
				v12 = v31
				goto l10
			}
		l25:
			v5 = v29
			v9 = v30
			goto l22
		}
		v5 = v20
		goto l6
	}
l4:
	t76 := m.fn748(v19, i32(0))
	m.memory[int64(uint32(v4))+864] = byte(t76)
	t77 := m.fn748(v19, i32(2))
	m.memory[int64(uint32(v4))+790] = byte(t77)
	t78 := m.fn748(v19, i32(3))
	m.memory[int64(uint32(v4))+791] = byte(t78)
	t79 := m.fn748(v19, i32(7))
	m.memory[int64(uint32(v4))+792] = byte(t79)
	t80 := m.fn748(v19, i32(8))
	m.memory[int64(uint32(v4))+793] = byte(t80)
	memory_copy(m.memory, uint32(v4+i32(16)), uint32(v4+i32(448)), uint32(i32(432)))
	t81 := m.fn113(i32(8), i32(432))
	v5 = t81
	memory_copy(m.memory, uint32(v5), uint32(v4+i32(16)), uint32(i32(432)))
	t82 := int32(load32(m.memory[uint32(v1):]))
	m.fn120(v4+i32(8), t82)
	t83 := int32(load32(m.memory[int64(uint32(v4))+8:]))
	v9 = t83
	t84 := int32(load32(m.memory[int64(uint32(v4))+12:]))
	v10 = t84
	m.memory[int64(uint32(v0))+80] = byte(i32(0))
	store64(m.memory[int64(uint32(v0))+72:], uint64(i64(0)))
	store32(m.memory[int64(uint32(v0))+92:], uint32(v5))
	store32(m.memory[int64(uint32(v0))+88:], uint32(v3))
	store32(m.memory[int64(uint32(v0))+84:], uint32(v2))
	store32(m.memory[int64(uint32(v0))+68:], uint32(v10))
	store32(m.memory[int64(uint32(v0))+64:], uint32(v9))
	t85 := int32(m.memory[int64(uint32(v1))+8])
	m.memory[int64(uint32(v0))+61] = byte(t85)
	m.memory[int64(uint32(v0))+60] = byte(i32(0))
	store16(m.memory[int64(uint32(v0))+58:], uint16(i32(0)))
	t86 := int32(m.memory[int64(uint32(v1))+9])
	m.memory[int64(uint32(v0))+57] = byte(t86)
	t87 := int32(m.memory[int64(uint32(v1))+10])
	m.memory[int64(uint32(v0))+56] = byte(t87)
	store64(m.memory[int64(uint32(v0))+48:], uint64(i64(0)))
	store64(m.memory[int64(uint32(v0))+40:], uint64(i64(1)))
	store64(m.memory[int64(uint32(v0))+32:], uint64(i64(0)))
	store32(m.memory[int64(uint32(v0))+16:], uint32(i32(2)))
	store64(m.memory[uint32(v0):], uint64(i64(0)))
	m.g0 = v4 + i32(880)
}
func (m *Module) fn747(v0, v1 int32) {
	var v2 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+256:]))
		v2 = t0
		if uint32(v2) < uint32(i32(257)) {
			goto l0
		}
		m.fn256(i32(1148664), i32(22), i32(1148688))
		panic("unreachable")
	}
l0:
	store32(m.memory[int64(uint32(v0))+256:], uint32(v2+i32(1)))
	m.memory[uint32(v0+v1&i32(255))] = byte(v2)
}
func (m *Module) fn748(v0, v1 int32) int32 {
	v0 = v1 & i32(255) * (v0 & i32(255))
	if int32(uint32(v0)>>8) == 0 {
		return v0
	}
	m.fn153(i32(1148648))
	panic("unreachable")
}
func (m *Module) fn749(v0, v1, v2 int32) {
	var v3 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		t1 := v2
		v3 = t0
		if uint32(t1) <= uint32(v3) {
			goto l0
		}
		m.fn151(v2, v3, v3, i32(1300824))
		panic("unreachable")
	}
l0:
	store32(m.memory[int64(uint32(v1))+8:], uint32(v2))
	store32(m.memory[int64(uint32(v0))+16:], uint32(i32(0)))
	store32(m.memory[int64(uint32(v0))+12:], uint32(v3))
	store32(m.memory[int64(uint32(v0))+8:], uint32(v1))
	t2 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	t3 := v0
	v1 = t2
	store32(m.memory[int64(uint32(t3))+4:], uint32(v1+v3*i32(28)))
	store32(m.memory[uint32(v0):], uint32(v1+v2*i32(28)))
}
func (m *Module) fn750(v0, v1 int32) {
	var v2 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		v2 = t0
		if uint32(v2) < uint32(v1) {
			return
		}
		store32(m.memory[int64(uint32(v0))+8:], uint32(v1))
		v2 = v2 - v1
		t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		v1 = t1 + v1<<4
	l1:
		{
			if v2 == 0 {
				return
			}
			t2 := int32(load32(m.memory[uint32(v1):]))
			t3 := int32(load32(m.memory[uint32(v1+i32(4)):]))
			m.fn16(t2, t3)
			v2 = v2 + i32(-1)
			v1 = v1 + i32(16)
			goto l1
		}
	}
}
func (m *Module) fn751(v0, v1 int32) {
	var v2 int32
	{
		t0 := int32(load32(m.memory[uint32(v0):]))
		t1 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		t2 := v1
		v2 = t1
		if uint32(t2) <= uint32(t0-v2) {
			return
		}
		m.fn62(v0, v2, v1, i32(4), i32(16))
	}
}
func (m *Module) fn752(v0 int32) {
	var v1 int32
	v1 = i32(0)
l1:
	if v1 == i32(400) {
		return
	}
	m.fn753(v0 + v1)
	v1 = v1 + i32(40)
	goto l1
}
func (m *Module) fn753(v0 int32) {
	t0 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	t1 := int32(load32(m.memory[int64(uint32(v0))+12:]))
	m.fn16(t0, t1)
	t2 := int32(load32(m.memory[int64(uint32(v0))+20:]))
	t3 := int32(load32(m.memory[int64(uint32(v0))+24:]))
	m.fn134(t2, t3)
}
func (m *Module) fn754(v0 int32) {
	var v1 int32
	t0 := int32(load32(m.memory[uint32(v0):]))
	v1 = t0
	t1 := int32(load32(m.memory[uint32(v1):]))
	t2 := v1
	v1 = t1 + i32(-1)
	store32(m.memory[uint32(t2):], uint32(v1))
	if v1 != 0 {
		return
	}
	m.fn755(v0)
}
func (m *Module) fn755(v0 int32) {
	var v1, v2 int32
	{
		t0 := int32(load32(m.memory[uint32(v0):]))
		v1 = t0
		if v1 == i32(-1) {
			return
		}
		t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		v2 = t1
		t2 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t3 := v1
		v0 = t2 + i32(-1)
		store32(m.memory[int64(uint32(t3))+4:], uint32(v0))
		if v0 != 0 {
			return
		}
		m.fn40(v1, i32(4), (v2+i32(11))&i32(-4))
	}
}
func (m *Module) fn756(v0 int32) {
	m.fn757(v0)
	t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
	t1 := int32(load32(m.memory[int64(uint32(v0))+16:]))
	m.fn134(t0, t1)
}
func (m *Module) fn757(v0 int32) {
	var v1, v2, v3 int32
	t0 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	v1 = t0
	v2 = v1 + i32(232)
	t1 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	v3 = t1
l1:
	{
		if v3 == 0 {
			goto l0
		}
		t2 := int32(load32(m.memory[uint32(v2+i32(-4)):]))
		t3 := int32(load32(m.memory[uint32(v2):]))
		m.fn16(t2, t3)
		t4 := int32(load32(m.memory[uint32(v2+i32(-16)):]))
		t5 := int32(load32(m.memory[uint32(v2+i32(-12)):]))
		m.fn134(t4, t5)
		v3 = v3 + i32(-1)
		v2 = v2 + i32(240)
		goto l1
	}
l0:
	t6 := int32(load32(m.memory[uint32(v0):]))
	m.fn136(t6, v1, i32(8), i32(240))
}
func (m *Module) fn758(v0 int32) int32 {
	var v1, v2 int32
	var v3 int64
	var v4 int32
	t0 := int32(load32(m.memory[int64(uint32(v0))+16:]))
	v1 = t0
	t1 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	v2 = t1
	t2 := int64(load64(m.memory[uint32(v0):]))
	v3 = t2
	var _ int32
l1:
	if v3 == 0 {
		t4 := v0
		v1 = v1 + i32(-320)
		store32(m.memory[int64(uint32(t4))+16:], uint32(v1))
		t5 := v0
		v4 = v2 + i32(8)
		store32(m.memory[int64(uint32(t5))+8:], uint32(v4))
		t6 := int64(load64(m.memory[uint32(v2):]))
		t7 := v0
		v3 = (t6 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
		store64(m.memory[uint32(t7):], uint64(v3))
		v2 = v4
		goto l1
	}
	store64(m.memory[uint32(v0):], uint64((v3+i64(-1))&v3))
	return v1 + (i32(0)-int32(uint32(int64(bits.TrailingZeros64(uint64(v3))))>>3))*i32(40)
}
func (m *Module) fn759(v0 int32) {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	m.fn16(t0, t1)
	t2 := int32(load32(m.memory[int64(uint32(v0))+12:]))
	t3 := int32(load32(m.memory[int64(uint32(v0))+16:]))
	m.fn16(t2, t3)
}
func (m *Module) fn760(v0 int32) {
	m.fn761(v0)
	m.fn762(v0 + i32(360))
	t0 := int32(load32(m.memory[int64(uint32(v0))+468:]))
	t1 := int32(load32(m.memory[int64(uint32(v0))+472:]))
	m.fn134(t0, t1)
}
func (m *Module) fn761(v0 int32) {
	var v1 int32
	v1 = i32(8)
l1:
	if v1 == i32(368) {
		return
	}
	m.fn763(v0 + v1)
	v1 = v1 + i32(40)
	goto l1
}
