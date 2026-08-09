package core

import (
	"math/bits"
)

func (m *Module) fn672(v0, v1 int32) int32 {
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
		m.fn237(t5, i32(40), p6)
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
		store64(m.memory[int64(uint32(v2))+24:], uint64(i64(0x800000028)))
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
				t19 := m.fn673(t17, t18, v10)
				v11 = t19
				t20 := m.fn26(t15, t16, v11)
				v12 = t20
				t21 := t14 + v12
				v13 = int32(uint32(int32(v11)) >> 25)
				m.memory[uint32(t21)] = byte(v13)
				m.memory[uint32(v6+v5&(v12+i32(-8))+i32(8))] = byte(v13)
				t22 := int32(load32(m.memory[uint32(v0):]))
				memory_copy(m.memory, uint32(v6+(v12^i32(-1))*i32(40)), uint32(t22+(v10^i32(-1))*i32(40)), uint32(i32(40)))
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
	m.fn241(v0, v2+i32(16), i32(108), i32(40))
l7:
	v5 = i32(-1)
l2:
	m.g0 = v2 + i32(64)
	return v5
}
func (m *Module) fn673(v0, v1, v2 int32) int64 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(t0):]))
	v0 = t1
	t2 := int64(load64(m.memory[uint32(v0):]))
	t3 := int64(load64(m.memory[uint32(v0+i32(8)):]))
	t4 := int32(load32(m.memory[uint32(v1):]))
	v0 = t4 + (i32(0)-v2)*i32(40)
	t5 := int32(load32(m.memory[uint32(v0+i32(-36)):]))
	t6 := int32(load32(m.memory[uint32(v0+i32(-32)):]))
	t7 := m.fn540(t2, t3, t5, t6)
	return t7
}
func (m *Module) fn674(v0, v1 int32) int32 {
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
		m.fn237(t5, i32(28), p6)
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
		store64(m.memory[int64(uint32(v2))+24:], uint64(i64(0x80000001c)))
		store32(m.memory[int64(uint32(v2))+20:], uint32(v0+i32(16)))
		t11 := int32(load32(m.memory[uint32(v0):]))
		v4 = t11
		t12 := int64(load64(m.memory[uint32(v4):]))
		v8 = (t12 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
		v9 = v2 + i32(32)
		v1 = i32(0)
	l6:
		if v3 == 0 {
			t28 := int32(load32(m.memory[int64(uint32(v0))+12:]))
			t29 := v2
			v1 = t28
			store32(m.memory[int64(uint32(t29))+44:], uint32(v1))
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
				t19 := m.fn675(t17, t18, v10)
				v11 = t19
				t20 := m.fn26(t15, t16, v11)
				v12 = t20
				t21 := t14 + v12
				v13 = int32(uint32(int32(v11)) >> 25)
				m.memory[uint32(t21)] = byte(v13)
				m.memory[uint32(v6+v5&(v12+i32(-8))+i32(8))] = byte(v13)
				v12 = v6 + (v12^i32(-1))*i32(28)
				t22 := int32(load32(m.memory[uint32(v0):]))
				t23 := v12
				v10 = t22 + (v10^i32(-1))*i32(28)
				t24 := int32(load32(m.memory[int64(uint32(v10))+24:]))
				store32(m.memory[int64(uint32(t23))+24:], uint32(t24))
				t25 := int64(load64(m.memory[int64(uint32(v10))+16:]))
				store64(m.memory[int64(uint32(v12))+16:], uint64(t25))
				t26 := int64(load64(m.memory[int64(uint32(v10))+8:]))
				store64(m.memory[int64(uint32(v12))+8:], uint64(t26))
				t27 := int64(load64(m.memory[uint32(v10):]))
				store64(m.memory[uint32(v12):], uint64(t27))
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
	m.fn241(v0, v2+i32(16), i32(109), i32(28))
l7:
	v5 = i32(-1)
l2:
	m.g0 = v2 + i32(64)
	return v5
}
func (m *Module) fn675(v0, v1, v2 int32) int64 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(t0):]))
	v0 = t1
	t2 := int64(load64(m.memory[uint32(v0):]))
	t3 := int64(load64(m.memory[uint32(v0+i32(8)):]))
	t4 := int32(load32(m.memory[uint32(v1):]))
	v0 = t4 + (i32(0)-v2)*i32(28)
	t5 := int32(load32(m.memory[uint32(v0+i32(-24)):]))
	t6 := int32(load32(m.memory[uint32(v0+i32(-20)):]))
	t7 := m.fn540(t2, t3, t5, t6)
	return t7
}
func (m *Module) fn676(v0, v1 int32) {
	t0 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	if t0 != 0 {
		return
	}
	_ = m.fn674(v0, v1)
}
func (m *Module) fn677(v0, v1 int32) int32 {
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
				t19 := m.fn678(t17, t18, v10)
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
	m.fn241(v0, v2+i32(16), i32(110), i32(36))
l7:
	v5 = i32(-1)
l2:
	m.g0 = v2 + i32(64)
	return v5
}
func (m *Module) fn678(v0, v1, v2 int32) int64 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(t0):]))
	v0 = t1
	t2 := int64(load64(m.memory[uint32(v0):]))
	t3 := int64(load64(m.memory[uint32(v0+i32(8)):]))
	t4 := int32(load32(m.memory[uint32(v1):]))
	v0 = t4 + (i32(0)-v2)*i32(36)
	t5 := int32(load32(m.memory[uint32(v0+i32(-32)):]))
	t6 := int32(load32(m.memory[uint32(v0+i32(-28)):]))
	t7 := m.fn540(t2, t3, t5, t6)
	return t7
}
func (m *Module) fn679(v0, v1 int32) int32 {
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
		m.fn237(t5, i32(28), p6)
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
		store64(m.memory[int64(uint32(v2))+24:], uint64(i64(0x80000001c)))
		store32(m.memory[int64(uint32(v2))+20:], uint32(v0+i32(16)))
		t11 := int32(load32(m.memory[uint32(v0):]))
		v4 = t11
		t12 := int64(load64(m.memory[uint32(v4):]))
		v8 = (t12 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
		v9 = v2 + i32(32)
		v1 = i32(0)
	l6:
		if v3 == 0 {
			t28 := int32(load32(m.memory[int64(uint32(v0))+12:]))
			t29 := v2
			v1 = t28
			store32(m.memory[int64(uint32(t29))+44:], uint32(v1))
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
				t19 := m.fn680(t17, t18, v10)
				v11 = t19
				t20 := m.fn26(t15, t16, v11)
				v12 = t20
				t21 := t14 + v12
				v13 = int32(uint32(int32(v11)) >> 25)
				m.memory[uint32(t21)] = byte(v13)
				m.memory[uint32(v6+v5&(v12+i32(-8))+i32(8))] = byte(v13)
				v12 = v6 + (v12^i32(-1))*i32(28)
				t22 := int32(load32(m.memory[uint32(v0):]))
				t23 := v12
				v10 = t22 + (v10^i32(-1))*i32(28)
				t24 := int32(load32(m.memory[int64(uint32(v10))+24:]))
				store32(m.memory[int64(uint32(t23))+24:], uint32(t24))
				t25 := int64(load64(m.memory[int64(uint32(v10))+16:]))
				store64(m.memory[int64(uint32(v12))+16:], uint64(t25))
				t26 := int64(load64(m.memory[int64(uint32(v10))+8:]))
				store64(m.memory[int64(uint32(v12))+8:], uint64(t26))
				t27 := int64(load64(m.memory[uint32(v10):]))
				store64(m.memory[uint32(v12):], uint64(t27))
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
	m.fn241(v0, v2+i32(16), i32(111), i32(28))
l7:
	v5 = i32(-1)
l2:
	m.g0 = v2 + i32(64)
	return v5
}
func (m *Module) fn680(v0, v1, v2 int32) int64 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(t0):]))
	v0 = t1
	t2 := int64(load64(m.memory[uint32(v0):]))
	t3 := int64(load64(m.memory[uint32(v0+i32(8)):]))
	t4 := int32(load32(m.memory[uint32(v1):]))
	v0 = t4 + (i32(0)-v2)*i32(28)
	t5 := int32(load32(m.memory[uint32(v0+i32(-24)):]))
	t6 := int32(load32(m.memory[uint32(v0+i32(-20)):]))
	t7 := m.fn540(t2, t3, t5, t6)
	return t7
}
func (m *Module) fn681(v0, v1, v2 int32) int32 {
	var v3, v4, v5, v6, v7 int32
	var v8 int64
	var v9, v10 int32
	var v11 int64
	var v12, v13 int32
	t0 := m.g0
	v3 = t0 - i32(64)
	m.g0 = v3
	store32(m.memory[int64(uint32(v3))+12:], uint32(v2))
	t1 := int32(load32(m.memory[int64(uint32(v0))+12:]))
	v4 = t1
	store32(m.memory[int64(uint32(v3))+16:], uint32(v3+i32(12)))
	v2 = v4 + v1
	if uint32(v2) < uint32(v4) {
		m.fn242()
		panic("unreachable")
	}
	{
		t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		t3 := v2
		v1 = t2
		p4 := int32(uint32(v1+i32(1))>>3) * i32(7)
		if uint32(v1) < uint32(i32(8)) {
			p4 = v1
		}
		v1 = p4
		if uint32(t3) <= uint32(int32(uint32(v1)>>1)) {
			goto l1
		}
		t5 := v3 + i32(48)
		v1 = v1 + i32(1)
		p6 := v2
		if uint32(v1) > uint32(v2) {
			p6 = v1
		}
		m.fn237(t5, i32(16), p6)
		t7 := int32(load32(m.memory[int64(uint32(v3))+52:]))
		v5 = t7
		t8 := int32(load32(m.memory[int64(uint32(v3))+48:]))
		v6 = t8
		if v6 == 0 {
			goto l2
		}
		t9 := int32(load32(m.memory[int64(uint32(v3))+56:]))
		v7 = t9
		t10 := int32(load32(m.memory[int64(uint32(v3))+60:]))
		store32(m.memory[int64(uint32(v3))+44:], uint32(t10))
		store32(m.memory[int64(uint32(v3))+40:], uint32(v7))
		store32(m.memory[int64(uint32(v3))+36:], uint32(v5))
		store32(m.memory[int64(uint32(v3))+32:], uint32(v6))
		store64(m.memory[int64(uint32(v3))+24:], uint64(i64(0x800000010)))
		store32(m.memory[int64(uint32(v3))+20:], uint32(v0+i32(16)))
		t11 := int32(load32(m.memory[uint32(v0):]))
		v1 = t11
		t12 := int64(load64(m.memory[uint32(v1):]))
		v8 = (t12 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
		v9 = v3 + i32(32)
		v2 = i32(0)
	l6:
		if v4 == 0 {
			t26 := int32(load32(m.memory[int64(uint32(v0))+12:]))
			t27 := v3
			v2 = t26
			store32(m.memory[int64(uint32(t27))+44:], uint32(v2))
			store32(m.memory[int64(uint32(v3))+40:], uint32(v7-v2))
			m.fn239(v0, v9)
			m.fn240(v3 + i32(20))
			goto l7
		}
	l5:
		{
			if v8 != i64(0) {
				t14 := v6
				t15 := v6
				t16 := v5
				t17 := v3 + i32(16)
				t18 := v0
				v10 = int32(uint32(int64(bits.TrailingZeros64(uint64(v8))))>>3) + v2
				t19 := m.fn665(t17, t18, v10)
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
				v4 = v4 + i32(-1)
				v8 = (v8 + i64(-1)) & v8
				goto l6
			}
			v2 = v2 + i32(8)
			v1 = v1 + i32(8)
			t13 := int64(load64(m.memory[uint32(v1):]))
			v8 = (t13 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
			goto l5
		}
	}
l1:
	m.fn241(v0, v3+i32(16), i32(104), i32(16))
l7:
	v5 = i32(-1)
l2:
	m.g0 = v3 + i32(64)
	return v5
}
func (m *Module) fn682(v0, v1, v2 int32) int32 {
	var v3, v4, v5, v6, v7 int32
	var v8 int64
	var v9, v10 int32
	var v11 int64
	var v12, v13 int32
	t0 := m.g0
	v3 = t0 - i32(64)
	m.g0 = v3
	store32(m.memory[int64(uint32(v3))+12:], uint32(v2))
	t1 := int32(load32(m.memory[int64(uint32(v0))+12:]))
	v4 = t1
	store32(m.memory[int64(uint32(v3))+16:], uint32(v3+i32(12)))
	v2 = v4 + v1
	if uint32(v2) < uint32(v4) {
		m.fn242()
		panic("unreachable")
	}
	{
		t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		t3 := v2
		v1 = t2
		p4 := int32(uint32(v1+i32(1))>>3) * i32(7)
		if uint32(v1) < uint32(i32(8)) {
			p4 = v1
		}
		v1 = p4
		if uint32(t3) <= uint32(int32(uint32(v1)>>1)) {
			goto l1
		}
		t5 := v3 + i32(48)
		v1 = v1 + i32(1)
		p6 := v2
		if uint32(v1) > uint32(v2) {
			p6 = v1
		}
		m.fn237(t5, i32(12), p6)
		t7 := int32(load32(m.memory[int64(uint32(v3))+52:]))
		v5 = t7
		t8 := int32(load32(m.memory[int64(uint32(v3))+48:]))
		v6 = t8
		if v6 == 0 {
			goto l2
		}
		t9 := int32(load32(m.memory[int64(uint32(v3))+56:]))
		v7 = t9
		t10 := int32(load32(m.memory[int64(uint32(v3))+60:]))
		store32(m.memory[int64(uint32(v3))+44:], uint32(t10))
		store32(m.memory[int64(uint32(v3))+40:], uint32(v7))
		store32(m.memory[int64(uint32(v3))+36:], uint32(v5))
		store32(m.memory[int64(uint32(v3))+32:], uint32(v6))
		store64(m.memory[int64(uint32(v3))+24:], uint64(i64(0x80000000c)))
		store32(m.memory[int64(uint32(v3))+20:], uint32(v0+i32(16)))
		t11 := int32(load32(m.memory[uint32(v0):]))
		v1 = t11
		t12 := int64(load64(m.memory[uint32(v1):]))
		v8 = (t12 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
		v9 = v3 + i32(20) + i32(12)
		v2 = i32(0)
	l6:
		if v4 == 0 {
			t26 := int32(load32(m.memory[int64(uint32(v0))+12:]))
			t27 := v3
			v2 = t26
			store32(m.memory[int64(uint32(t27))+44:], uint32(v2))
			store32(m.memory[int64(uint32(v3))+40:], uint32(v7-v2))
			m.fn239(v0, v9)
			m.fn240(v3 + i32(20))
			goto l7
		}
	l5:
		{
			if v8 != i64(0) {
				t14 := v6
				t15 := v6
				t16 := v5
				t17 := v3 + i32(16)
				t18 := v0
				v10 = int32(uint32(int64(bits.TrailingZeros64(uint64(v8))))>>3) + v2
				t19 := m.fn683(t17, t18, v10)
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
				v4 = v4 + i32(-1)
				v8 = (v8 + i64(-1)) & v8
				goto l6
			}
			v2 = v2 + i32(8)
			v1 = v1 + i32(8)
			t13 := int64(load64(m.memory[uint32(v1):]))
			v8 = (t13 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
			goto l5
		}
	}
l1:
	m.fn241(v0, v3+i32(16), i32(112), i32(12))
l7:
	v5 = i32(-1)
l2:
	m.g0 = v3 + i32(64)
	return v5
}
func (m *Module) fn683(v0, v1, v2 int32) int64 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(t0):]))
	v0 = t1
	t2 := int64(load64(m.memory[uint32(v0):]))
	t3 := int64(load64(m.memory[uint32(v0+i32(8)):]))
	t4 := int32(load32(m.memory[uint32(v1):]))
	v0 = t4 + (i32(0)-v2)*i32(12)
	t5 := int32(load32(m.memory[uint32(v0+i32(-8)):]))
	t6 := int32(load32(m.memory[uint32(v0+i32(-4)):]))
	t7 := m.fn540(t2, t3, t5, t6)
	return t7
}
func (m *Module) fn684(v0, v1, v2 int32) {
	t0 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	if uint32(v1) <= uint32(t0) {
		return
	}
	_ = m.fn682(v0, v1, v2)
}
func (m *Module) fn685(v0, v1 int32) int32 {
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
				t19 := m.fn686(t17, t18, v10)
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
	m.fn241(v0, v2+i32(16), i32(113), i32(24))
l7:
	v5 = i32(-1)
l2:
	m.g0 = v2 + i32(64)
	return v5
}
func (m *Module) fn686(v0, v1, v2 int32) int64 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(t0):]))
	v0 = t1
	t2 := int64(load64(m.memory[uint32(v0):]))
	t3 := int64(load64(m.memory[uint32(v0+i32(8)):]))
	t4 := int32(load32(m.memory[uint32(v1):]))
	v0 = t4 + (i32(0)-v2)*i32(24)
	t5 := int32(load32(m.memory[uint32(v0+i32(-20)):]))
	t6 := int32(load32(m.memory[uint32(v0+i32(-16)):]))
	t7 := m.fn540(t2, t3, t5, t6)
	return t7
}
func (m *Module) fn687(v0, v1 int32) int32 {
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
		m.fn237(t5, i32(488), p6)
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
		store64(m.memory[int64(uint32(v2))+24:], uint64(i64(0x8000001e8)))
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
				t19 := m.fn688(t17, t18, v10)
				v11 = t19
				t20 := m.fn26(t15, t16, v11)
				v12 = t20
				t21 := t14 + v12
				v13 = int32(uint32(int32(v11)) >> 25)
				m.memory[uint32(t21)] = byte(v13)
				m.memory[uint32(v6+v5&(v12+i32(-8))+i32(8))] = byte(v13)
				t22 := int32(load32(m.memory[uint32(v0):]))
				memory_copy(m.memory, uint32(v6+(v12^i32(-1))*i32(488)), uint32(t22+(v10^i32(-1))*i32(488)), uint32(i32(488)))
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
	m.fn241(v0, v2+i32(16), i32(114), i32(488))
l7:
	v5 = i32(-1)
l2:
	m.g0 = v2 + i32(64)
	return v5
}
func (m *Module) fn688(v0, v1, v2 int32) int64 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(t0):]))
	v0 = t1
	t2 := int64(load64(m.memory[uint32(v0):]))
	t3 := int64(load64(m.memory[uint32(v0+i32(8)):]))
	t4 := int32(load32(m.memory[uint32(v1):]))
	v0 = t4 + (i32(0)-v2)*i32(488)
	t5 := int32(load32(m.memory[uint32(v0+i32(-488)):]))
	t6 := int32(load32(m.memory[uint32(v0+i32(-484)):]))
	t7 := m.fn24(t2, t3, t5, t6)
	return t7
}
func (m *Module) fn689(v0, v1, v2 int32) int64 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(t0):]))
	v0 = t1
	t2 := int64(load64(m.memory[uint32(v0):]))
	t3 := int64(load64(m.memory[uint32(v0+i32(8)):]))
	t4 := int32(load32(m.memory[uint32(v1):]))
	v0 = t4 + (i32(0)-v2)*i32(12)
	t5 := int32(load32(m.memory[uint32(v0+i32(-12)):]))
	t6 := int32(load32(m.memory[uint32(v0+i32(-8)):]))
	t7 := m.fn24(t2, t3, t5, t6)
	return t7
}
func (m *Module) fn690(v0, v1 int32) int32 {
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
				t19 := m.fn691(t17, t18, v10)
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
	m.fn241(v0, v2+i32(16), i32(115), i32(20))
l7:
	v5 = i32(-1)
l2:
	m.g0 = v2 + i32(64)
	return v5
}
func (m *Module) fn691(v0, v1, v2 int32) int64 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(t0):]))
	v0 = t1
	t2 := int64(load64(m.memory[uint32(v0):]))
	t3 := int64(load64(m.memory[uint32(v0+i32(8)):]))
	t4 := int32(load32(m.memory[uint32(v1):]))
	v0 = t4 + (i32(0)-v2)*i32(20)
	t5 := int32(load32(m.memory[uint32(v0+i32(-20)):]))
	t6 := int32(load32(m.memory[uint32(v0+i32(-16)):]))
	t7 := m.fn24(t2, t3, t5, t6)
	return t7
}
func (m *Module) fn692(v0, v1 int32) int32 {
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
				t19 := m.fn693(t17, t18, v10)
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
	m.fn241(v0, v2+i32(16), i32(116), i32(8))
l7:
	v5 = i32(-1)
l2:
	m.g0 = v2 + i32(64)
	return v5
}
func (m *Module) fn693(v0, v1, v2 int32) int64 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(t0):]))
	v0 = t1
	t2 := int64(load64(m.memory[uint32(v0):]))
	t3 := int64(load64(m.memory[uint32(v0+i32(8)):]))
	t4 := int32(load32(m.memory[uint32(v1):]))
	v2 = t4 - v2<<3
	t5 := int32(load32(m.memory[uint32(v2+i32(-8)):]))
	t6 := int32(load32(m.memory[uint32(v2+i32(-4)):]))
	t7 := m.fn24(t2, t3, t5, t6)
	return t7
}
func (m *Module) fn694(v0, v1 int32) int32 {
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
				t19 := m.fn695(t17, t18, v10)
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
	m.fn241(v0, v2+i32(16), i32(117), i32(24))
l7:
	v5 = i32(-1)
l2:
	m.g0 = v2 + i32(64)
	return v5
}
func (m *Module) fn695(v0, v1, v2 int32) int64 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(t0):]))
	v0 = t1
	t2 := int64(load64(m.memory[uint32(v0):]))
	t3 := int64(load64(m.memory[uint32(v0+i32(8)):]))
	t4 := int32(load32(m.memory[uint32(v1):]))
	t5 := m.fn696(t2, t3, t4+(i32(0)-v2)*i32(24)+i32(-24))
	return t5
}
func (m *Module) fn696(v0, v1 int64, v2 int32) int64 {
	var v3 int32
	t0 := m.g0
	v3 = t0 - i32(64)
	m.g0 = v3
	store64(m.memory[int64(uint32(v3))+48:], uint64(i64(0)))
	store64(m.memory[int64(uint32(v3))+56:], uint64(i64(0)))
	store64(m.memory[int64(uint32(v3))+40:], uint64(v1))
	store64(m.memory[int64(uint32(v3))+24:], uint64(v1^i64(8387220255154660723)))
	store64(m.memory[int64(uint32(v3))+16:], uint64(v1^i64(7237128888997146477)))
	store64(m.memory[int64(uint32(v3))+32:], uint64(v0))
	store64(m.memory[int64(uint32(v3))+8:], uint64(v0^i64(0x6c7967656e657261)))
	store64(m.memory[uint32(v3):], uint64(v0^i64(8317987319222330741)))
	t1 := int32(load32(m.memory[int64(uint32(v2))+4:]))
	t2 := int32(load32(m.memory[int64(uint32(v2))+8:]))
	m.fn173(v3, t1, t2)
	t3 := int32(load32(m.memory[int64(uint32(v2))+12:]))
	m.fn172(v3, t3)
	t4 := m.fn174(v3)
	v1 = t4
	m.g0 = v3 + i32(64)
	return v1
}
func (m *Module) fn697(v0, v1 int32) int32 {
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
				t19 := m.fn698(t17, t18, v10)
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
	m.fn241(v0, v2+i32(16), i32(118), i32(16))
l7:
	v5 = i32(-1)
l2:
	m.g0 = v2 + i32(64)
	return v5
}
func (m *Module) fn698(v0, v1, v2 int32) int64 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(t0):]))
	v0 = t1
	t2 := int64(load64(m.memory[uint32(v0):]))
	t3 := int64(load64(m.memory[uint32(v0+i32(8)):]))
	t4 := int32(load32(m.memory[uint32(v1):]))
	v2 = t4 - v2<<4
	t5 := int32(load32(m.memory[uint32(v2+i32(-16)):]))
	t6 := int32(load32(m.memory[uint32(v2+i32(-12)):]))
	t7 := m.fn651(t2, t3, t5, t6)
	return t7
}
func (m *Module) fn699(v0, v1 int32) int32 {
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
				t19 := m.fn700(t17, t18, v10)
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
	m.fn241(v0, v2+i32(16), i32(119), i32(8))
l7:
	v5 = i32(-1)
l2:
	m.g0 = v2 + i32(64)
	return v5
}
func (m *Module) fn700(v0, v1, v2 int32) int64 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(t0):]))
	v0 = t1
	t2 := int64(load64(m.memory[uint32(v0):]))
	t3 := int64(load64(m.memory[uint32(v0+i32(8)):]))
	t4 := int32(load32(m.memory[uint32(v1):]))
	v2 = t4 - v2<<3
	t5 := int32(load32(m.memory[uint32(v2+i32(-8)):]))
	t6 := int32(load32(m.memory[uint32(v2+i32(-4)):]))
	t7 := m.fn651(t2, t3, t5, t6)
	return t7
}
func (m *Module) fn701(v0, v1 int32) int32 {
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
				t19 := m.fn702(t17, t18, v10)
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
	m.fn241(v0, v2+i32(16), i32(120), i32(8))
l7:
	v5 = i32(-1)
l2:
	m.g0 = v2 + i32(64)
	return v5
}
func (m *Module) fn702(v0, v1, v2 int32) int64 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(t0):]))
	v0 = t1
	t2 := int64(load64(m.memory[uint32(v0):]))
	t3 := int64(load64(m.memory[uint32(v0+i32(8)):]))
	t4 := int32(load32(m.memory[uint32(v1):]))
	v2 = t4 - v2<<3
	t5 := int32(load16(m.memory[uint32(v2+i32(-8)):]))
	t6 := int32(load32(m.memory[uint32(v2+i32(-4)):]))
	t7 := m.fn703(t2, t3, t5, t6)
	return t7
}
func (m *Module) fn703(v0, v1 int64, v2, v3 int32) int64 {
	var v4 int32
	t0 := m.g0
	v4 = t0 - i32(64)
	m.g0 = v4
	store64(m.memory[int64(uint32(v4))+48:], uint64(i64(0)))
	store64(m.memory[int64(uint32(v4))+56:], uint64(i64(0)))
	store64(m.memory[int64(uint32(v4))+40:], uint64(v1))
	store64(m.memory[int64(uint32(v4))+24:], uint64(v1^i64(8387220255154660723)))
	store64(m.memory[int64(uint32(v4))+16:], uint64(v1^i64(7237128888997146477)))
	store64(m.memory[int64(uint32(v4))+32:], uint64(v0))
	store64(m.memory[int64(uint32(v4))+8:], uint64(v0^i64(0x6c7967656e657261)))
	store64(m.memory[uint32(v4):], uint64(v0^i64(8317987319222330741)))
	m.fn530(v2, v4)
	m.fn172(v4, v3)
	t1 := m.fn174(v4)
	v1 = t1
	m.g0 = v4 + i32(64)
	return v1
}
func (m *Module) fn704(v0, v1 int32) int32 {
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
				t19 := m.fn705(t17, t18, v10)
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
	m.fn241(v0, v2+i32(16), i32(121), i32(16))
l7:
	v5 = i32(-1)
l2:
	m.g0 = v2 + i32(64)
	return v5
}
func (m *Module) fn705(v0, v1, v2 int32) int64 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(t0):]))
	v0 = t1
	t2 := int64(load64(m.memory[uint32(v0):]))
	t3 := int64(load64(m.memory[uint32(v0+i32(8)):]))
	t4 := int32(load32(m.memory[uint32(v1):]))
	t5 := int32(load32(m.memory[uint32(t4-v2<<4+i32(-16)):]))
	t6 := m.fn66(t2, t3, t5)
	return t6
}
func (m *Module) fn706(v0, v1 int32) int32 {
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
				t19 := m.fn707(t17, t18, v10)
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
	m.fn241(v0, v2+i32(16), i32(122), i32(16))
l7:
	v5 = i32(-1)
l2:
	m.g0 = v2 + i32(64)
	return v5
}
func (m *Module) fn707(v0, v1, v2 int32) int64 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(t0):]))
	v0 = t1
	t2 := int64(load64(m.memory[uint32(v0):]))
	t3 := int64(load64(m.memory[uint32(v0+i32(8)):]))
	t4 := int32(load32(m.memory[uint32(v1):]))
	t5 := int32(load32(m.memory[uint32(t4-v2<<4+i32(-16)):]))
	t6 := m.fn66(t2, t3, t5)
	return t6
}
func (m *Module) fn708(v0, v1, v2 int32) int32 {
	var v3, v4, v5, v6, v7 int32
	var v8 int64
	var v9, v10 int32
	var v11 int64
	var v12, v13 int32
	t0 := m.g0
	v3 = t0 - i32(64)
	m.g0 = v3
	store32(m.memory[int64(uint32(v3))+12:], uint32(v2))
	t1 := int32(load32(m.memory[int64(uint32(v0))+12:]))
	v4 = t1
	store32(m.memory[int64(uint32(v3))+16:], uint32(v3+i32(12)))
	v2 = v4 + v1
	if uint32(v2) < uint32(v4) {
		m.fn242()
		panic("unreachable")
	}
	{
		t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		t3 := v2
		v1 = t2
		p4 := int32(uint32(v1+i32(1))>>3) * i32(7)
		if uint32(v1) < uint32(i32(8)) {
			p4 = v1
		}
		v1 = p4
		if uint32(t3) <= uint32(int32(uint32(v1)>>1)) {
			goto l1
		}
		t5 := v3 + i32(48)
		v1 = v1 + i32(1)
		p6 := v2
		if uint32(v1) > uint32(v2) {
			p6 = v1
		}
		m.fn237(t5, i32(12), p6)
		t7 := int32(load32(m.memory[int64(uint32(v3))+52:]))
		v5 = t7
		t8 := int32(load32(m.memory[int64(uint32(v3))+48:]))
		v6 = t8
		if v6 == 0 {
			goto l2
		}
		t9 := int32(load32(m.memory[int64(uint32(v3))+56:]))
		v7 = t9
		t10 := int32(load32(m.memory[int64(uint32(v3))+60:]))
		store32(m.memory[int64(uint32(v3))+44:], uint32(t10))
		store32(m.memory[int64(uint32(v3))+40:], uint32(v7))
		store32(m.memory[int64(uint32(v3))+36:], uint32(v5))
		store32(m.memory[int64(uint32(v3))+32:], uint32(v6))
		store64(m.memory[int64(uint32(v3))+24:], uint64(i64(0x80000000c)))
		store32(m.memory[int64(uint32(v3))+20:], uint32(v0+i32(16)))
		t11 := int32(load32(m.memory[uint32(v0):]))
		v1 = t11
		t12 := int64(load64(m.memory[uint32(v1):]))
		v8 = (t12 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
		v9 = v3 + i32(20) + i32(12)
		v2 = i32(0)
	l6:
		if v4 == 0 {
			t26 := int32(load32(m.memory[int64(uint32(v0))+12:]))
			t27 := v3
			v2 = t26
			store32(m.memory[int64(uint32(t27))+44:], uint32(v2))
			store32(m.memory[int64(uint32(v3))+40:], uint32(v7-v2))
			m.fn239(v0, v9)
			m.fn240(v3 + i32(20))
			goto l7
		}
	l5:
		{
			if v8 != i64(0) {
				t14 := v6
				t15 := v6
				t16 := v5
				t17 := v3 + i32(16)
				t18 := v0
				v10 = int32(uint32(int64(bits.TrailingZeros64(uint64(v8))))>>3) + v2
				t19 := m.fn709(t17, t18, v10)
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
				v4 = v4 + i32(-1)
				v8 = (v8 + i64(-1)) & v8
				goto l6
			}
			v2 = v2 + i32(8)
			v1 = v1 + i32(8)
			t13 := int64(load64(m.memory[uint32(v1):]))
			v8 = (t13 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
			goto l5
		}
	}
l1:
	m.fn241(v0, v3+i32(16), i32(123), i32(12))
l7:
	v5 = i32(-1)
l2:
	m.g0 = v3 + i32(64)
	return v5
}
func (m *Module) fn709(v0, v1, v2 int32) int64 {
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
func (m *Module) fn710(v0, v1, v2 int32) {
	t0 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	if uint32(v1) <= uint32(t0) {
		return
	}
	_ = m.fn708(v0, v1, v2)
}
func (m *Module) fn711(v0, v1 int32) int32 {
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
				t19 := m.fn712(t17, t18, v10)
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
	m.fn241(v0, v2+i32(16), i32(124), i32(8))
l7:
	v5 = i32(-1)
l2:
	m.g0 = v2 + i32(64)
	return v5
}
func (m *Module) fn712(v0, v1, v2 int32) int64 {
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
func (m *Module) fn713(v0, v1 int32) {
	t0 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	if t0 != 0 {
		return
	}
	_ = m.fn711(v0, v1)
}
func (m *Module) fn714(v0, v1, v2 int32) int64 {
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
func (m *Module) fn715(v0, v1 int32) int32 {
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
		m.fn237(t5, i32(296), p6)
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
		store64(m.memory[int64(uint32(v2))+24:], uint64(i64(0x800000128)))
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
				t19 := m.fn716(t17, t18, v10)
				v11 = t19
				t20 := m.fn26(t15, t16, v11)
				v12 = t20
				t21 := t14 + v12
				v13 = int32(uint32(int32(v11)) >> 25)
				m.memory[uint32(t21)] = byte(v13)
				m.memory[uint32(v6+v5&(v12+i32(-8))+i32(8))] = byte(v13)
				t22 := int32(load32(m.memory[uint32(v0):]))
				memory_copy(m.memory, uint32(v6+(v12^i32(-1))*i32(296)), uint32(t22+(v10^i32(-1))*i32(296)), uint32(i32(296)))
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
	m.fn241(v0, v2+i32(16), i32(125), i32(296))
l7:
	v5 = i32(-1)
l2:
	m.g0 = v2 + i32(64)
	return v5
}
func (m *Module) fn716(v0, v1, v2 int32) int64 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(t0):]))
	v0 = t1
	t2 := int64(load64(m.memory[uint32(v0):]))
	t3 := int64(load64(m.memory[uint32(v0+i32(8)):]))
	t4 := int32(load32(m.memory[uint32(v1):]))
	t5 := int32(load32(m.memory[uint32(t4+(i32(0)-v2)*i32(296)+i32(-296)):]))
	t6 := m.fn66(t2, t3, t5)
	return t6
}
