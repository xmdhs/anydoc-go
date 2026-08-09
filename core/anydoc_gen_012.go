package core

import (
	"math/bits"
)

func (m *Module) fn492(v0, v1, v2 int32) {
	if v2 == 0 {
		m.fn494(i32(1300840))
		panic("unreachable")
	}
	store32(m.memory[int64(uint32(v0))+4:], uint32(i32(1)))
	t0 := int32(uint32(v1) / uint32(v2))
	t1 := v0
	v2 = t0
	store32(m.memory[int64(uint32(t1))+8:], uint32(v2))
	store32(m.memory[uint32(v0):], uint32(v2))
}
func (m *Module) fn493(v0, v1 int32) {
	var v2 int32
	{
		t0 := int32(load32(m.memory[uint32(v0):]))
		t1 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		t2 := v1
		v2 = t1
		if uint32(t2) <= uint32(t0-v2) {
			return
		}
		m.fn1842(v0, v2, v1, i32(4), i32(4))
	}
}
func (m *Module) fn494(v0 int32) {
	m.fn91(i32(1111304), i32(51), v0)
	panic("unreachable")
}
func (m *Module) fn495(v0 int32) {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	m.fn16(t0, t1)
}
func (m *Module) fn496(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t1 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	t2 := int32(load32(m.memory[uint32(v1):]))
	t3 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	t4 := m.fn107(t0, t1, t2, t3)
	return t4
}
func (m *Module) fn497(v0, v1 int32) {
	var v2, v3 int32
	var v4, v5 int64
	var v6, v7, v8 int32
	t0 := m.g0
	v2 = t0 - i32(80)
	m.g0 = v2
	t1 := int32(load32(m.memory[int64(uint32(v1))+64:]))
	v3 = t1
	m.fn22(v2+i32(8), i32(3))
	t2 := int64(load64(m.memory[int64(uint32(v2))+8:]))
	v4 = t2
	t3 := int64(load64(m.memory[int64(uint32(v2))+16:]))
	v5 = t3
	m.fn237(v2+i32(64), i32(24), v3)
	store64(m.memory[int64(uint32(v2))+32:], uint64(v5))
	store64(m.memory[int64(uint32(v2))+24:], uint64(v4))
	t4 := int64(load64(m.memory[int64(uint32(v2))+72:]))
	store64(m.memory[int64(uint32(v2))+16:], uint64(t4))
	t5 := int64(load64(m.memory[int64(uint32(v2))+64:]))
	store64(m.memory[int64(uint32(v2))+8:], uint64(t5))
	t6 := int32(load32(m.memory[int64(uint32(v1))+48:]))
	v3 = t6 * i32(192)
	t7 := int32(load32(m.memory[int64(uint32(v1))+44:]))
	v1 = t7
l1:
	{
		if v3 == 0 {
			goto l0
		}
		t8 := int32(load32(m.memory[int64(uint32(v1))+40:]))
		t9 := v2 + i32(64)
		v6 = t8
		t10 := int32(load32(m.memory[int64(uint32(v1))+44:]))
		t11 := v6
		v7 = t10
		m.fn498(t9, t11, v7, i32(92), i32(47))
		t12 := int32(load32(m.memory[int64(uint32(v2))+68:]))
		t13 := v2 + i32(40)
		v8 = t12
		t14 := int32(load32(m.memory[int64(uint32(v2))+72:]))
		m.fn14(t13, v8, t14)
		t15 := int32(load32(m.memory[int64(uint32(v2))+64:]))
		m.fn16(t15, v8)
		m.fn51(v2+i32(64), v6, v7)
		m.fn499(v2+i32(52), v2+i32(8), v2+i32(40), v2+i32(64))
		t16 := int32(load32(m.memory[int64(uint32(v2))+52:]))
		t17 := int32(load32(m.memory[int64(uint32(v2))+56:]))
		m.fn134(t16, t17)
		v3 = v3 + i32(-192)
		v1 = v1 + i32(192)
		goto l1
	}
l0:
	t18 := int64(load64(m.memory[int64(uint32(v2))+32:]))
	store64(m.memory[int64(uint32(v0))+24:], uint64(t18))
	t19 := int64(load64(m.memory[int64(uint32(v2))+24:]))
	store64(m.memory[int64(uint32(v0))+16:], uint64(t19))
	t20 := int64(load64(m.memory[int64(uint32(v2))+16:]))
	store64(m.memory[int64(uint32(v0))+8:], uint64(t20))
	t21 := int64(load64(m.memory[int64(uint32(v2))+8:]))
	store64(m.memory[uint32(v0):], uint64(t21))
	m.g0 = v2 + i32(80)
}
func (m *Module) fn498(v0, v1, v2, v3, v4 int32) {
	var v5, v6, v7, v8 int32
	t0 := m.g0
	v5 = t0 - i32(32)
	m.g0 = v5
	m.fn59(v5+i32(8), v2, i32(1), i32(1))
	store32(m.memory[int64(uint32(v5))+28:], uint32(i32(0)))
	t1 := int64(load64(m.memory[int64(uint32(v5))+8:]))
	store64(m.memory[int64(uint32(v5))+20:], uint64(t1))
	m.fn47(v5+i32(20), v2)
	t2 := int32(load32(m.memory[int64(uint32(v5))+28:]))
	v6 = t2
	{
		if v2 == 0 {
			goto l0
		}
		t3 := int32(load32(m.memory[int64(uint32(v5))+24:]))
		v7 = t3
		v8 = v3 & i32(255)
	l1:
		{
			t4 := int32(m.memory[uint32(v1)])
			t5 := v7 + v6
			t6 := v4
			v3 = t4
			p7 := v3
			if v3 == v8 {
				p7 = t6
			}
			m.memory[uint32(t5)] = byte(p7)
			v1 = v1 + i32(1)
			v6 = v6 + i32(1)
			v2 = v2 + i32(-1)
			if v2 != 0 {
				goto l1
			}
		}
	}
l0:
	t8 := int64(load64(m.memory[int64(uint32(v5))+20:]))
	store64(m.memory[uint32(v0):], uint64(t8))
	store32(m.memory[int64(uint32(v0))+8:], uint32(v6))
	m.g0 = v5 + i32(32)
}
func (m *Module) fn499(v0, v1, v2, v3 int32) {
	var v4 int32
	var v5 int64
	var v6, v7, v8, v9, v10 int32
	var v11 int64
	t0 := m.g0
	v4 = t0 - i32(48)
	m.g0 = v4
	t1 := int64(load64(m.memory[int64(uint32(v1))+16:]))
	t2 := int64(load64(m.memory[int64(uint32(v1))+24:]))
	t3 := int32(load32(m.memory[int64(uint32(v2))+4:]))
	t4 := int32(load32(m.memory[int64(uint32(v2))+8:]))
	t5 := m.fn540(t1, t2, t3, t4)
	v5 = t5
	store32(m.memory[int64(uint32(v4))+44:], uint32(v2))
	m.fn541(v1, i32(1), v1+i32(16))
	store32(m.memory[int64(uint32(v4))+20:], uint32(v1))
	store32(m.memory[int64(uint32(v4))+16:], uint32(v4+i32(44)))
	t6 := int32(load32(m.memory[uint32(v1):]))
	t7 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	m.fn69(v4+i32(8), t6, t7, v5, v4+i32(16), i32(87))
	t8 := int32(load32(m.memory[uint32(v1):]))
	v6 = t8
	t9 := int32(load32(m.memory[int64(uint32(v4))+12:]))
	v7 = t9
	{
		{
			t10 := int32(load32(m.memory[int64(uint32(v4))+8:]))
			if t10 != i32(1) {
				goto l0
			}
			v8 = v6 + v7
			t11 := int32(m.memory[uint32(v8)])
			v9 = t11
			t12 := int32(load32(m.memory[int64(uint32(v2))+8:]))
			v10 = t12
			t13 := int64(load64(m.memory[uint32(v2):]))
			v11 = t13
			t14 := v8
			v2 = int32(uint32(int32(v5)) >> 25)
			m.memory[uint32(t14)] = byte(v2)
			t15 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			m.memory[uint32(v6+t15&(v7+i32(-8))+i32(8))] = byte(v2)
			t16 := int32(load32(m.memory[int64(uint32(v1))+12:]))
			store32(m.memory[int64(uint32(v1))+12:], uint32(t16+i32(1)))
			t17 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			store32(m.memory[int64(uint32(v1))+8:], uint32(t17-v9&i32(1)))
			v1 = v6 + (i32(0)-v7)*i32(24) + i32(-24)
			store64(m.memory[uint32(v1):], uint64(v11))
			store32(m.memory[int64(uint32(v4))+24:], uint32(v10))
			t18 := int64(load64(m.memory[uint32(v3):]))
			store64(m.memory[int64(uint32(v4))+28:], uint64(t18))
			t19 := int64(load64(m.memory[int64(uint32(v4))+24:]))
			store64(m.memory[int64(uint32(v1))+8:], uint64(t19))
			t20 := int32(load32(m.memory[int64(uint32(v3))+8:]))
			store32(m.memory[int64(uint32(v4))+36:], uint32(t20))
			t21 := int64(load64(m.memory[int64(uint32(v4))+32:]))
			store64(m.memory[int64(uint32(v1))+16:], uint64(t21))
			store32(m.memory[uint32(v0):], uint32(i32(-1)))
			goto l1
		}
	l0:
		v1 = v6 + (i32(0)-v7)*i32(24) + i32(-12)
		t22 := int64(load64(m.memory[uint32(v1):]))
		v5 = t22
		t23 := int64(load64(m.memory[uint32(v3):]))
		store64(m.memory[uint32(v1):], uint64(t23))
		t24 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		v6 = t24
		t25 := int32(load32(m.memory[int64(uint32(v3))+8:]))
		store32(m.memory[int64(uint32(v1))+8:], uint32(t25))
		store32(m.memory[int64(uint32(v0))+8:], uint32(v6))
		store64(m.memory[uint32(v0):], uint64(v5))
		t26 := int32(load32(m.memory[uint32(v2):]))
		t27 := int32(load32(m.memory[int64(uint32(v2))+4:]))
		m.fn16(t26, t27)
	}
l1:
	m.g0 = v4 + i32(48)
}
func (m *Module) fn500(v0 int32) {
	var v1, v2, v3, v4 int32
	var v5 int64
	var v6, v7 int32
	t0 := m.g0
	v1 = t0 - i32(16)
	m.g0 = v1
	{
		t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		v2 = t1
		if v2 == 0 {
			goto l0
		}
		t2 := int32(load32(m.memory[uint32(v0):]))
		v3 = t2
		{
			t3 := int32(load32(m.memory[int64(uint32(v0))+12:]))
			v4 = t3
			if v4 == 0 {
				goto l1
			}
			v0 = v3 + i32(8)
			t4 := int64(load64(m.memory[uint32(v3):]))
			v5 = (t4 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
			v6 = v3
		l4:
			if v4 == 0 {
				goto l1
			}
		l3:
			{
				if v5 != i64(0) {
					v7 = v6 + (i32(0)-int32(uint32(int64(bits.TrailingZeros64(uint64(v5))))>>3))*i32(24)
					t6 := int32(load32(m.memory[uint32(v7+i32(-24)):]))
					t7 := int32(load32(m.memory[uint32(v7+i32(-20)):]))
					m.fn16(t6, t7)
					t8 := int32(load32(m.memory[uint32(v7+i32(-12)):]))
					t9 := int32(load32(m.memory[uint32(v7+i32(-8)):]))
					m.fn16(t8, t9)
					v4 = v4 + i32(-1)
					v5 = (v5 + i64(-1)) & v5
					goto l4
				}
				v6 = v6 + i32(-192)
				t5 := int64(load64(m.memory[uint32(v0):]))
				v5 = (t5 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
				v0 = v0 + i32(8)
				goto l3
			}
		}
	l1:
		m.fn39(v1+i32(4), i32(24), i32(8), v2+i32(1))
		t10 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		t11 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t12 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		m.fn40(v3-t10, t11, t12)
	}
l0:
	m.g0 = v1 + i32(16)
}
func (m *Module) fn501(v0 int32) {
	m.fn132(v0 + i32(32))
	m.fn78(v0 + i32(88))
	m.fn168(v0 + i32(100))
	m.fn78(v0 + i32(112))
	t0 := int32(load32(m.memory[int64(uint32(v0))+124:]))
	t1 := int32(load32(m.memory[int64(uint32(v0))+128:]))
	m.fn16(t0, t1)
	m.fn445(v0 + i32(8))
	m.fn502(v0 + i32(56))
}
func (m *Module) fn502(v0 int32) {
	var v1, v2, v3, v4 int32
	var v5 int64
	var v6 int32
	t0 := m.g0
	v1 = t0 - i32(16)
	m.g0 = v1
	{
		t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		v2 = t1
		if v2 == 0 {
			goto l0
		}
		t2 := int32(load32(m.memory[uint32(v0):]))
		v3 = t2
		{
			t3 := int32(load32(m.memory[int64(uint32(v0))+12:]))
			v4 = t3
			if v4 == 0 {
				goto l1
			}
			v0 = v3 + i32(8)
			t4 := int64(load64(m.memory[uint32(v3):]))
			v5 = (t4 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
			v6 = v3
		l4:
			if v4 == 0 {
				goto l1
			}
		l3:
			{
				if v5 != i64(0) {
					m.fn169(v6 + (i32(0)-int32(uint32(int64(bits.TrailingZeros64(uint64(v5))))>>3))*i32(24) + i32(-24))
					v4 = v4 + i32(-1)
					v5 = (v5 + i64(-1)) & v5
					goto l4
				}
				v6 = v6 + i32(-192)
				t5 := int64(load64(m.memory[uint32(v0):]))
				v5 = (t5 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
				v0 = v0 + i32(8)
				goto l3
			}
		}
	l1:
		m.fn39(v1+i32(4), i32(24), i32(8), v2+i32(1))
		t6 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		t7 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t8 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		m.fn40(v3-t6, t7, t8)
	}
l0:
	m.g0 = v1 + i32(16)
}
func (m *Module) fn503(v0, v1, v2, v3, v4 int32) {
	var v5 int32
	t0 := m.g0
	v5 = t0 - i32(224)
	m.g0 = v5
	m.fn504(v5+i32(8), v4, v2, v3)
	t1 := int32(load32(m.memory[int64(uint32(v5))+8:]))
	t2 := int32(load32(m.memory[int64(uint32(v5))+12:]))
	m.fn114(v5+i32(16), v1, t1, t2)
	{
		t3 := int64(load64(m.memory[int64(uint32(v5))+16:]))
		if t3 != i64(-1) {
			goto l0
		}
		v4 = v0 + i32(4)
		v1 = v5 + i32(24)
		{
			t4 := int32(load32(m.memory[int64(uint32(v5))+24:]))
			if t4 == i32(-0x7ffffffd) {
				m.fn51(v4, v2, v3)
				store64(m.memory[int64(uint32(v0))+24:], uint64(i64(-1)))
				store32(m.memory[uint32(v0):], uint32(i32(-0x7fffffeb)))
				m.fn116(v1)
				goto l2
			}
			store64(m.memory[int64(uint32(v0))+24:], uint64(i64(-1)))
			store32(m.memory[uint32(v0):], uint32(i32(-0x7ffffff0)))
			t5 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			store32(m.memory[int64(uint32(v4))+8:], uint32(t5))
			t6 := int64(load64(m.memory[uint32(v1):]))
			store64(m.memory[uint32(v4):], uint64(t6))
			goto l2
		}
	}
l0:
	m.fn139(v0, v5+i32(16))
	m.memory[int64(uint32(v0))+232] = byte(i32(0))
l2:
	m.g0 = v5 + i32(224)
}
func (m *Module) fn504(v0, v1, v2, v3 int32) {
	var v4, v5, v6, v7, v8, v9 int32
	var v10 int64
	var v11 int32
	var v12 int64
	var v13 int32
	var v14 int64
	t0 := m.g0
	v4 = t0 - i32(16)
	m.g0 = v4
	m.fn377(v4+i32(4), v2, v3)
	t1 := int32(load32(m.memory[int64(uint32(v4))+12:]))
	v5 = t1
	t2 := int32(load32(m.memory[int64(uint32(v4))+8:]))
	v6 = t2
	t3 := int32(load32(m.memory[int64(uint32(v4))+4:]))
	v7 = t3
	v8 = i32(0)
l1:
	{
		if v5 == v8 {
			goto l0
		}
		v9 = v6 + v8
		t4 := int32(m.memory[uint32(v9)])
		t5 := v9
		v9 = t4
		p6 := i32(0)
		if uint32((v9+i32(-65))&i32(255)) < uint32(i32(26)) {
			p6 = i32(32)
		}
		m.memory[uint32(t5)] = byte(p6 | v9)
		v8 = v8 + i32(1)
		goto l1
	}
l0:
	{
		t7 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		if t7 == 0 {
			goto l2
		}
		t8 := int64(load64(m.memory[int64(uint32(v1))+16:]))
		t9 := int64(load64(m.memory[int64(uint32(v1))+24:]))
		t10 := m.fn540(t8, t9, v6, v5)
		v10 = t10
		t11 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		v11 = t11
		v9 = v11 & int32(v10)
		v12 = int64(uint64(v10)>>25) & i64(127) * i64(72340172838076673)
		t12 := int32(load32(m.memory[uint32(v1):]))
		v1 = t12
		v13 = i32(0)
	l6:
		{
			t13 := int64(load64(m.memory[uint32(v1+v9):]))
			v14 = t13
			v10 = v14 ^ v12
			v10 = (v10 ^ i64(-1)) & (v10 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
		l5:
			{
				if v10 == 0 {
					if !(v14&(v14<<1)&i64(-0x7f7f7f7f7f7f7f80) == 0) {
						goto l2
					}
					t19 := v9
					v13 = v13 + i32(8)
					v9 = (t19 + v13) & v11
					goto l6
				}
				t14 := v6
				t15 := v5
				v8 = v1 + (i32(0)-(int32(uint32(int64(bits.TrailingZeros64(uint64(v10))))>>3)+v9)&v11)*i32(24)
				t16 := int32(load32(m.memory[uint32(v8+i32(-20)):]))
				t17 := int32(load32(m.memory[uint32(v8+i32(-16)):]))
				t18 := m.fn545(t14, t15, t16, t17)
				if t18 != 0 {
					goto l4
				}
				v10 = (v10 + i64(-1)) & v10
				goto l5
			}
		l4:
		}
		t20 := int32(load32(m.memory[uint32(v8+i32(-4)):]))
		v9 = t20
		t21 := int32(load32(m.memory[uint32(v8+i32(-8)):]))
		v8 = t21
		goto l7
	}
l2:
	v8 = i32(0)
l7:
	m.fn16(v7, v6)
	t23 := v0
	p22 := v3
	if v8 != 0 {
		p22 = v9
	}
	store32(m.memory[int64(uint32(t23))+4:], uint32(p22))
	t25 := v0
	p24 := v2
	if v8 != 0 {
		p24 = v8
	}
	store32(m.memory[uint32(t25):], uint32(p24))
	m.g0 = v4 + i32(16)
}
func (m *Module) fn505(v0, v1 int32) {
	var v2, v3 int32
	var v4 int64
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	m.fn506(v2+i32(8), v1)
	{
		{
			t1 := int32(m.memory[int64(uint32(v2))+8])
			if t1 != i32(255) {
				goto l0
			}
			t2 := int32(m.memory[int64(uint32(v2))+9])
			v3 = t2
			goto l1
		}
	l0:
		t3 := int64(load64(m.memory[int64(uint32(v2))+8:]))
		v4 = t3
		if v4&i64(255) != i64(255) {
			store64(m.memory[uint32(v0):], uint64(v4))
			goto l8
		}
		v3 = int32(int64(uint64(v4) >> 8))
	}
l1:
	if int32(int8(v3)) > i32(-1) {
		v1 = v3 & i32(255)
		goto l7
	}
	m.fn506(v2+i32(8), v1)
	{
		t4 := int32(m.memory[int64(uint32(v2))+8])
		if t4 != i32(255) {
			t6 := int64(load64(m.memory[int64(uint32(v2))+8:]))
			v4 = t6
			if v4&i64(255) != i64(255) {
				store64(m.memory[uint32(v0):], uint64(v4))
				goto l8
			}
			v1 = int32(int64(uint64(v4) >> 8))
			goto l5
		}
		t5 := int32(m.memory[int64(uint32(v2))+9])
		v1 = t5
		goto l5
	}
l5:
	v1 = v1&i32(127)<<7 | v3&i32(127)
l7:
	m.memory[uint32(v0)] = byte(i32(255))
	store16(m.memory[int64(uint32(v0))+2:], uint16(v1))
l8:
	m.g0 = v2 + i32(16)
}
func (m *Module) fn506(v0, v1 int32) {
	var v2, v3 int32
	var v4 int64
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	t1 := v2 + i32(8)
	t2 := v1
	v3 = v1 + i32(232)
	m.fn508(t1, t2, v3, i32(1))
	{
		{
			t3 := int32(m.memory[int64(uint32(v2))+8])
			if t3 == i32(255) {
				goto l0
			}
			t4 := int64(load64(m.memory[int64(uint32(v2))+8:]))
			v4 = t4
			if v4&i64(255) == i64(255) {
				goto l0
			}
			store64(m.memory[uint32(v0):], uint64(v4))
			goto l1
		}
	l0:
		m.memory[uint32(v0)] = byte(i32(255))
		t5 := int32(m.memory[uint32(v3)])
		m.memory[int64(uint32(v0))+1] = byte(t5)
	}
l1:
	m.g0 = v2 + i32(16)
}
func (m *Module) fn507(v0, v1, v2 int32) {
	var v3, v4 int32
	var v5 int64
	var v6, v7 int32
	t0 := m.g0
	v3 = t0 - i32(16)
	m.g0 = v3
	m.fn506(v3+i32(8), v1)
	{
		{
			t1 := int32(m.memory[int64(uint32(v3))+8])
			if t1 != i32(255) {
				goto l0
			}
			t2 := int32(m.memory[int64(uint32(v3))+9])
			v4 = t2
			goto l1
		}
	l0:
		t3 := int64(load64(m.memory[int64(uint32(v3))+8:]))
		v5 = t3
		if v5&i64(255) != i64(255) {
			goto l2
		}
		v4 = int32(int64(uint64(v5) >> 8))
	}
l1:
	v6 = v4 & i32(127)
	v7 = i32(1)
l11:
	{
		{
			if uint32(v7) > uint32(i32(3)) {
				goto l3
			}
			if int32(int8(v4)) < i32(0) {
				m.fn506(v3+i32(8), v1)
				{
					t11 := int32(m.memory[int64(uint32(v3))+8])
					if t11 != i32(255) {
						t13 := int64(load64(m.memory[int64(uint32(v3))+8:]))
						v5 = t13
						if v5&i64(255) != i64(255) {
							store64(m.memory[uint32(v0):], uint64(v5))
							goto l7
						}
						v4 = int32(int64(uint64(v5) >> 8))
						goto l9
					}
					t12 := int32(m.memory[int64(uint32(v3))+9])
					v4 = t12
					goto l9
				}
			}
		l3:
			{
				t4 := int32(load32(m.memory[int64(uint32(v2))+8:]))
				v7 = t4
				if uint32(v7) >= uint32(v6) {
					goto l5
				}
				m.fn482(v2, v6)
				t5 := int32(load32(m.memory[int64(uint32(v2))+8:]))
				v7 = t5
			}
		l5:
			t6 := int32(load32(m.memory[int64(uint32(v2))+4:]))
			m.fn363(v3, v6, t6, v7, i32(1078224))
			t7 := int32(load32(m.memory[uint32(v3):]))
			t8 := int32(load32(m.memory[int64(uint32(v3))+4:]))
			m.fn508(v3+i32(8), v1, t7, t8)
			t9 := int32(m.memory[int64(uint32(v3))+8])
			if t9 == i32(255) {
				goto l6
			}
			t10 := int64(load64(m.memory[int64(uint32(v3))+8:]))
			v5 = t10
			if v5&i64(255) == i64(255) {
				goto l6
			}
			store64(m.memory[uint32(v0):], uint64(v5))
			goto l7
		}
	l6:
		m.memory[uint32(v0)] = byte(i32(255))
		store32(m.memory[int64(uint32(v0))+4:], uint32(v6))
		goto l7
	l9:
		v6 = i32_shl(v4&i32(127), v7*i32(7)) + v6
		t14 := v7
		var p15 int32
		if uint32(v7) < uint32(i32(4)) {
			p15 = 1
		}
		v7 = t14 + p15
		goto l11
	}
l2:
	store64(m.memory[uint32(v0):], uint64(v5))
l7:
	m.g0 = v3 + i32(16)
}
func (m *Module) fn508(v0, v1, v2, v3 int32) {
	var v4, v5, v6, v7 int32
	var v8 int64
	var v9, v10 int32
	t0 := m.g0
	v4 = t0 - i32(32)
	m.g0 = v4
	{
		{
			t1 := int32(load32(m.memory[int64(uint32(v1))+12:]))
			t2 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			t3 := v3
			v5 = t2
			if uint32(t3) <= uint32(t1-v5) {
				goto l0
			}
			v6 = v1 + i32(24)
			v7 = v1 + i32(200)
		l14:
			{
				if v3 == 0 {
					m.memory[uint32(v0)] = byte(i32(255))
					goto l8
				}
				{
					t4 := int32(load32(m.memory[int64(uint32(v1))+8:]))
					t5 := int32(load32(m.memory[int64(uint32(v1))+12:]))
					if t4 != t5 {
						goto l2
					}
					t6 := int32(load32(m.memory[int64(uint32(v1))+4:]))
					if uint32(v3) >= uint32(t6) {
						store64(m.memory[int64(uint32(v1))+8:], uint64(i64(0)))
						m.fn297(v4, v7, v2, v3)
						t12 := int32(m.memory[uint32(v4)])
						v5 = t12
						goto l7
					}
				}
			l2:
				m.fn290(v4+i32(16), v1, v6)
				{
					{
						t7 := int32(load32(m.memory[int64(uint32(v4))+16:]))
						if t7 == 0 {
							goto l4
						}
						t8 := int64(load64(m.memory[int64(uint32(v4))+20:]))
						v8 = t8
						goto l5
					}
				l4:
					t9 := int64(load64(m.memory[int64(uint32(v4))+20:]))
					store64(m.memory[int64(uint32(v4))+8:], uint64(t9))
					m.fn307(v4+i32(16), v4+i32(8), v2, v3)
					t10 := int32(m.memory[int64(uint32(v4))+16])
					if t10 == i32(255) {
						t13 := int32(load32(m.memory[int64(uint32(v1))+12:]))
						t14 := v1
						v9 = t13
						t15 := int32(load32(m.memory[int64(uint32(v1))+8:]))
						t16 := int32(load32(m.memory[int64(uint32(v4))+20:]))
						t17 := v9
						v5 = t16
						v10 = t15 + v5
						p18 := v10
						if uint32(v9) < uint32(v10) {
							p18 = t17
						}
						store32(m.memory[int64(uint32(t14))+8:], uint32(p18))
						goto l9
					}
					t11 := int64(load64(m.memory[int64(uint32(v4))+16:]))
					v8 = t11
				}
			l5:
				store64(m.memory[uint32(v4):], uint64(v8))
				v5 = int32(v8)
				goto l7
			l7:
				{
					if v5&i32(255) == i32(255) {
						goto l10
					}
					t19 := m.fn313(v4)
					if t19 != 0 {
						t23 := int32(load32(m.memory[int64(uint32(v4))+4:]))
						m.fn119(v5, t23)
						goto l14
					}
					t20 := int64(load64(m.memory[uint32(v4):]))
					store64(m.memory[uint32(v0):], uint64(t20))
					goto l8
				}
			l10:
				t21 := int32(load32(m.memory[int64(uint32(v4))+4:]))
				v5 = t21
			}
		l9:
			if v5 == 0 {
				t22 := int64(load64(m.memory[int64(uint32(i32(0)))+1287056:]))
				store64(m.memory[uint32(v0):], uint64(t22))
				goto l8
			}
			if uint32(v3) < uint32(v5) {
				m.fn151(v5, v3, v3, i32(1072408))
				panic("unreachable")
			}
			v2 = v2 + v5
			v3 = v3 - v5
			goto l14
		}
	l0:
		t24 := int32(load32(m.memory[uint32(v1):]))
		m.fn310(v2, v3, t24+v5, v3, i32(1074932))
		m.memory[uint32(v0)] = byte(i32(255))
		store32(m.memory[int64(uint32(v1))+8:], uint32(v5+v3))
	}
l8:
	m.g0 = v4 + i32(32)
}
func (m *Module) fn509(v0, v1, v2, v3 int32) int32 {
	if uint32(v2) < uint32(v1) {
		return v0 + v2
	}
	m.fn158(v2, v1, v3)
	panic("unreachable")
}
func (m *Module) fn510(v0, v1, v2, v3 int32) {
	var v4, v5 int32
	t0 := m.g0
	v4 = t0 - i32(32)
	m.g0 = v4
	v5 = i32(3)
	{
		{
			{
				t1 := m.fn159(v2, v3, i32(1282764), i32(3))
				if t1 == 0 {
					goto l0
				}
				v1 = i32(0x139000)
				goto l1
			}
		l0:
			v5 = i32(2)
			{
				t2 := m.fn159(v2, v3, i32(1282762), i32(2))
				if t2 == 0 {
					goto l2
				}
				v1 = i32(1282400)
				goto l1
			}
		l2:
			v5 = i32(2)
			t3 := m.fn159(v2, v3, i32(1282760), i32(2))
			if t3 == 0 {
				goto l3
			}
			v1 = i32(1282404)
		}
	l1:
		m.fn148(v4+i32(8), v5, v2, v3, i32(1087640))
		t4 := int32(load32(m.memory[uint32(v1):]))
		v1 = t4
		t5 := int32(load32(m.memory[int64(uint32(v4))+12:]))
		v3 = t5
		t6 := int32(load32(m.memory[int64(uint32(v4))+8:]))
		v2 = t6
	}
l3:
	m.fn511(v4+i32(16), v1, v2, v3)
	t7 := int32(load32(m.memory[int64(uint32(v4))+24:]))
	store32(m.memory[int64(uint32(v0))+8:], uint32(t7))
	t8 := int64(load64(m.memory[int64(uint32(v4))+16:]))
	store64(m.memory[uint32(v0):], uint64(t8))
	t9 := int32(m.memory[int64(uint32(v4))+28])
	m.memory[int64(uint32(v0))+16] = byte(t9)
	store32(m.memory[int64(uint32(v0))+12:], uint32(v1))
	m.g0 = v4 + i32(32)
}
func (m *Module) fn511(v0, v1, v2, v3 int32) {
	var v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14, v15 int32
	t0 := m.g0
	v4 = t0 - i32(176)
	m.g0 = v4
	{
		{
			{
				if v1 == i32(1153092) {
					goto l0
				}
				if v1 == i32(1153580) {
					goto l0
				}
				if v1 != i32(1153064) {
					if v1 == i32(1148960) {
						goto l6
					}
					{
						if v1 == i32(1153400) {
							t14 := m.fn1700(v2, v3)
							v8 = t14
							goto l8
						}
						t13 := m.fn1693(v2, v3)
						v8 = t13
						goto l8
					}
				}
			l0:
				m.fn1695(v4+i32(112), v1)
				m.fn1696(v4+i32(88), v4+i32(112), v3)
				t1 := int32(load32(m.memory[int64(uint32(v4))+88:]))
				t2 := int32(load32(m.memory[int64(uint32(v4))+92:]))
				m.fn1697(v4+i32(80), t1, t2)
				t3 := int32(load32(m.memory[int64(uint32(v4))+84:]))
				v5 = t3
				t4 := int32(load32(m.memory[int64(uint32(v4))+80:]))
				v6 = t4
				m.fn1698(v4+i32(72), v4+i32(112), v3)
				t5 := int32(load32(m.memory[int64(uint32(v4))+76:]))
				v1 = t5
				t6 := int32(load32(m.memory[int64(uint32(v4))+72:]))
				v7 = t6
				{
					if v6 != i32(1) {
						goto l2
					}
					p7 := v5
					if uint32(v1) < uint32(v5) {
						p7 = v1
					}
					p8 := v5
					if v7&i32(1) != 0 {
						p8 = p7
					}
					v1 = p8
					goto l3
				}
			l2:
				if v7&i32(1) == 0 {
					m.fn153(i32(1154832))
					panic("unreachable")
				}
			l3:
				m.fn1699(v4+i32(100), v1)
				t9 := int32(load32(m.memory[int64(uint32(v4))+136:]))
				store32(m.memory[int64(uint32(v4))+168:], uint32(t9))
				t10 := int64(load64(m.memory[int64(uint32(v4))+128:]))
				store64(m.memory[int64(uint32(v4))+160:], uint64(t10))
				t11 := int64(load64(m.memory[int64(uint32(v4))+120:]))
				store64(m.memory[int64(uint32(v4))+152:], uint64(t11))
				t12 := int64(load64(m.memory[int64(uint32(v4))+112:]))
				store64(m.memory[int64(uint32(v4))+144:], uint64(t12))
				v8 = i32(0)
				goto l5
			}
		l6:
			t15 := m.fn1692(v2, v3)
			v8 = t15
		}
	l8:
		{
			if v3 == v8 {
				m.memory[int64(uint32(v0))+12] = byte(i32(0))
				store32(m.memory[int64(uint32(v0))+8:], uint32(v3))
				store32(m.memory[int64(uint32(v0))+4:], uint32(v2))
				store32(m.memory[uint32(v0):], uint32(i32(-1)))
				goto l14
			}
			m.fn1695(v4+i32(112), v1)
			t16 := v4 + i32(64)
			t17 := v4 + i32(112)
			v1 = v3 - v8
			m.fn1696(t16, t17, v1)
			t18 := int32(load32(m.memory[int64(uint32(v4))+64:]))
			t19 := int32(load32(m.memory[int64(uint32(v4))+68:]))
			t20 := v4 + i32(56)
			v5 = t19
			v6 = v5 + v8
			var p21 int32
			if uint32(v6) >= uint32(v5) {
				p21 = 1
			}
			m.fn1697(t20, t18&p21, v6)
			t22 := int32(load32(m.memory[int64(uint32(v4))+60:]))
			v5 = t22
			t23 := int32(load32(m.memory[int64(uint32(v4))+56:]))
			v6 = t23
			m.fn1698(v4+i32(48), v4+i32(112), v1)
			t24 := int32(load32(m.memory[int64(uint32(v4))+48:]))
			t25 := int32(load32(m.memory[int64(uint32(v4))+52:]))
			v7 = t25
			v1 = v7 + v8
			var p26 int32
			if uint32(v1) >= uint32(v7) {
				p26 = 1
			}
			v7 = t24 & p26
			{
				if v6 != i32(1) {
					goto l10
				}
				p27 := v5
				if uint32(v1) < uint32(v5) {
					p27 = v1
				}
				p28 := v5
				if v7 != 0 {
					p28 = p27
				}
				v1 = p28
				goto l11
			}
		l10:
			if v7 == 0 {
				m.fn153(i32(1154848))
				panic("unreachable")
			}
		l11:
			m.fn1699(v4+i32(100), v1)
			store32(m.memory[int64(uint32(v4))+108:], uint32(v8))
			{
				if v8 == 0 {
					goto l13
				}
				t29 := int32(load32(m.memory[int64(uint32(v4))+104:]))
				memory_copy(m.memory, uint32(t29), uint32(v2), uint32(v8))
			}
		l13:
			t30 := int32(load32(m.memory[int64(uint32(v4))+136:]))
			store32(m.memory[int64(uint32(v4))+168:], uint32(t30))
			t31 := int64(load64(m.memory[int64(uint32(v4))+128:]))
			store64(m.memory[int64(uint32(v4))+160:], uint64(t31))
			t32 := int64(load64(m.memory[int64(uint32(v4))+120:]))
			store64(m.memory[int64(uint32(v4))+152:], uint64(t32))
			t33 := int64(load64(m.memory[int64(uint32(v4))+112:]))
			store64(m.memory[int64(uint32(v4))+144:], uint64(t33))
			goto l5
		}
	l5:
		t34 := int32(load32(m.memory[int64(uint32(v4))+168:]))
		store32(m.memory[int64(uint32(v4))+136:], uint32(t34))
		t35 := int64(load64(m.memory[int64(uint32(v4))+160:]))
		store64(m.memory[int64(uint32(v4))+128:], uint64(t35))
		t36 := int64(load64(m.memory[int64(uint32(v4))+152:]))
		store64(m.memory[int64(uint32(v4))+120:], uint64(t36))
		t37 := int64(load64(m.memory[int64(uint32(v4))+144:]))
		store64(m.memory[int64(uint32(v4))+112:], uint64(t37))
		v9 = i32(0)
	l22:
		{
			m.fn148(v4+i32(40), v8, v2, v3, i32(1154864))
			t38 := int32(load32(m.memory[int64(uint32(v4))+44:]))
			v10 = t38
			t39 := int32(load32(m.memory[int64(uint32(v4))+40:]))
			v11 = t39
			t40 := int32(load32(m.memory[int64(uint32(v4))+108:]))
			t41 := v4 + i32(32)
			v12 = t40
			t42 := int32(load32(m.memory[int64(uint32(v4))+104:]))
			t43 := int32(load32(m.memory[int64(uint32(v4))+100:]))
			m.fn212(t41, v12, t42, t43, i32(1155220))
			t44 := int32(load32(m.memory[int64(uint32(v4))+36:]))
			v5 = t44
			t45 := int32(load32(m.memory[int64(uint32(v4))+32:]))
			v6 = t45
			v1 = i32(0)
			v7 = i32(0)
			v13 = i32(0)
		l19:
			{
				m.fn148(v4+i32(24), v7, v11, v10, i32(1155140))
				t46 := int32(load32(m.memory[int64(uint32(v4))+28:]))
				v14 = t46
				t47 := int32(load32(m.memory[int64(uint32(v4))+24:]))
				v15 = t47
				m.fn212(v4+i32(16), v1, v6, v5, i32(1155156))
				t48 := int32(load32(m.memory[int64(uint32(v4))+16:]))
				t49 := int32(load32(m.memory[int64(uint32(v4))+20:]))
				m.fn1701(v4+i32(144), v4+i32(112), v15, v14, t48, t49)
				t50 := int32(load32(m.memory[int64(uint32(v4))+152:]))
				v1 = t50 + v1
				t51 := int32(load32(m.memory[int64(uint32(v4))+144:]))
				v7 = t51 + v7
				t52 := int32(m.memory[int64(uint32(v4))+148])
				v14 = t52
				if v14 != i32(2) {
					goto l15
				}
				if uint32(v1) >= uint32(v5) {
					m.fn158(v1, v5, i32(1155172))
					panic("unreachable")
				}
				m.memory[uint32(v6+v1)] = byte(i32(239))
				v14 = v1 + i32(1)
				if uint32(v14) >= uint32(v5) {
					m.fn158(v14, v5, i32(1155188))
					panic("unreachable")
				}
				m.memory[uint32(v6+v14)] = byte(i32(191))
				v14 = v1 + i32(2)
				if uint32(v14) >= uint32(v5) {
					goto l18
				}
				m.memory[uint32(v6+v14)] = byte(i32(189))
				v1 = v1 + i32(3)
				v13 = i32(1)
				goto l19
			l18:
			}
			m.fn158(v14, v5, i32(1155204))
			panic("unreachable")
		l15:
			store32(m.memory[int64(uint32(v4))+108:], uint32(v1+v12))
			v9 = v9 | v13
			{
				if v14&i32(1) == 0 {
					goto l20
				}
				t53 := v4 + i32(8)
				t54 := v4 + i32(112)
				t55 := v3
				v8 = v7 + v8
				m.fn1698(t53, t54, t55-v8)
				t56 := int32(load32(m.memory[int64(uint32(v4))+8:]))
				if t56&i32(1) == 0 {
					m.fn153(i32(1154880))
					panic("unreachable")
				}
				t57 := int32(load32(m.memory[int64(uint32(v4))+12:]))
				m.fn1702(v4+i32(100), t57)
				goto l22
			}
		l20:
		}
		t58 := int32(load32(m.memory[int64(uint32(v4))+108:]))
		store32(m.memory[int64(uint32(v0))+8:], uint32(t58))
		t59 := int64(load64(m.memory[int64(uint32(v4))+100:]))
		store64(m.memory[uint32(v0):], uint64(t59))
		m.memory[int64(uint32(v0))+12] = byte(v9 & i32(1))
	}
l14:
	m.g0 = v4 + i32(176)
}
func (m *Module) fn512(v0, v1, v2 int32) int32 {
	var v3 int64
	var v4, v5 int32
	var v6 int64
	var v7 int32
	var v8 int64
	var v9, v10, v11 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
		if t0 != 0 {
			t1 := int64(load64(m.memory[int64(uint32(v0))+16:]))
			t2 := int64(load64(m.memory[int64(uint32(v0))+24:]))
			t3 := m.fn314(t1, t2, v1, v2)
			v3 = t3
			t4 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			v4 = t4
			v5 = v4 & int32(v3)
			v6 = int64(uint64(v3)>>25) & i64(127) * i64(72340172838076673)
			t5 := int32(load32(m.memory[uint32(v0):]))
			v0 = t5
			v7 = i32(0)
			var _ int32
		l5:
			{
				t7 := int64(load64(m.memory[uint32(v0+v5):]))
				v8 = t7
				v3 = v8 ^ v6
				v3 = (v3 ^ i64(-1)) & (v3 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
				{
				l3:
					{
						var p8 int32
						if v3 == 0 {
							p8 = 1
						}
						v9 = p8
						if v9 != 0 {
							goto l1
						}
						t9 := v1
						t10 := v2
						t11 := v0
						v10 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v3))))>>3) + v5) & v4
						v11 = t11 + (i32(0)-v10)*i32(24)
						t12 := int32(load32(m.memory[uint32(v11+i32(-20)):]))
						t13 := int32(load32(m.memory[uint32(v11+i32(-16)):]))
						t14 := m.fn123(t9, t10, t12, t13)
						if t14 != 0 {
							goto l2
						}
						v3 = (v3 + i64(-1)) & v3
						goto l3
					}
				l1:
					if v8&(v8<<1)&i64(-0x7f7f7f7f7f7f7f80) == 0 {
						t17 := v5
						v7 = v7 + i32(8)
						v5 = (t17 + v7) & v4
						goto l5
					}
				l2:
					p15 := v0 + (i32(0)-v10)*i32(24)
					if v9 != 0 {
						p15 = i32(0)
					}
					p16 := p15 + i32(-12)
					if v9 != 0 {
						p16 = i32(0)
					}
					return p16
				}
			}
		}
		return i32(0)
	}
}
func (m *Module) fn513(v0, v1, v2, v3 int32) {
	m.fn514(v0, v3, v1, v2)
	store16(m.memory[int64(uint32(v0))+36:], uint16(i32(1)))
	store32(m.memory[int64(uint32(v0))+32:], uint32(v2))
	store32(m.memory[int64(uint32(v0))+28:], uint32(i32(0)))
}
func (m *Module) fn514(v0, v1, v2, v3 int32) {
	var v4, v5 int32
	t0 := m.g0
	v4 = t0 - i32(16)
	m.g0 = v4
	store32(m.memory[int64(uint32(v4))+12:], uint32(i32(0)))
	m.fn522(v4, v1, v4+i32(12))
	t1 := int32(load32(m.memory[int64(uint32(v4))+4:]))
	v5 = t1
	store32(m.memory[int64(uint32(v0))+16:], uint32(v3))
	store32(m.memory[int64(uint32(v0))+12:], uint32(i32(0)))
	store32(m.memory[int64(uint32(v0))+8:], uint32(v3))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v2))
	m.memory[int64(uint32(v0))+24] = byte(v5)
	store32(m.memory[uint32(v0):], uint32(v1))
	t2 := int32(load32(m.memory[int64(uint32(v4))+12:]))
	store32(m.memory[int64(uint32(v0))+20:], uint32(t2))
	m.g0 = v4 + i32(16)
}
func (m *Module) fn515(v0, v1 int32) {
	var v2, v3, v4 int32
	t0 := m.g0
	v2 = t0 - i32(32)
	m.g0 = v2
	{
		{
			t1 := int32(m.memory[int64(uint32(v1))+37])
			if t1 == 0 {
				goto l0
			}
			v1 = i32(0)
			goto l1
		}
	l0:
		t2 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		v3 = t2
		m.fn516(v2+i32(20), v1)
		{
			t3 := int32(load32(m.memory[int64(uint32(v2))+20:]))
			if t3 != i32(1) {
				goto l2
			}
			t4 := int32(load32(m.memory[int64(uint32(v1))+28:]))
			v4 = t4
			t5 := int32(load32(m.memory[int64(uint32(v2))+28:]))
			store32(m.memory[int64(uint32(v1))+28:], uint32(t5))
			v1 = v3 + v4
			t6 := int32(load32(m.memory[int64(uint32(v2))+24:]))
			v3 = t6 - v4
			goto l1
		}
	l2:
		m.fn517(v2+i32(8), v1)
		t7 := int32(load32(m.memory[int64(uint32(v2))+12:]))
		v3 = t7
		t8 := int32(load32(m.memory[int64(uint32(v2))+8:]))
		v1 = t8
	}
l1:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
	store32(m.memory[uint32(v0):], uint32(v1))
	m.g0 = v2 + i32(32)
}
func (m *Module) fn516(v0, v1 int32) {
	var v2, v3, v4, v5, v6, v7, v8, v9, v10, v11, v12 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	v3 = i32(0)
	{
		t1 := int32(load32(m.memory[int64(uint32(v1))+16:]))
		v4 = t1
		t2 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		t3 := v4
		v5 = t2
		if uint32(t3) > uint32(v5) {
			goto l0
		}
		v6 = v1 + i32(20)
		t4 := int32(m.memory[int64(uint32(v1))+24])
		t5 := v6
		v7 = t4
		v8 = t5 + v7 + i32(-1)
		t6 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		v9 = t6
		t7 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		v10 = t7
		var p8 int32
		if uint32(v7) < uint32(i32(5)) {
			p8 = 1
		}
		v11 = p8
	l2:
		{
			if uint32(v4) < uint32(v10) {
				goto l0
			}
			t9 := int32(m.memory[uint32(v8)])
			m.fn520(v2+i32(8), t9, v9+v10, v4-v10)
			t10 := int32(load32(m.memory[int64(uint32(v2))+8:]))
			if t10 != i32(1) {
				store32(m.memory[int64(uint32(v1))+12:], uint32(v4))
				goto l0
			}
			t11 := int32(load32(m.memory[int64(uint32(v2))+12:]))
			t12 := v1
			v10 = v10 + t11 + i32(1)
			store32(m.memory[int64(uint32(t12))+12:], uint32(v10))
			if uint32(v10) < uint32(v7) {
				goto l2
			}
			v12 = v10 - v7
			if uint32(v10) > uint32(v5) {
				goto l2
			}
			if v11 == 0 {
				m.fn151(i32(0), v7, i32(4), i32(1087368))
				panic("unreachable")
			}
			t13 := m.fn1851(v9+v12, v6, v7)
			if t13 != 0 {
				goto l2
			}
		}
		store32(m.memory[int64(uint32(v0))+8:], uint32(v10))
		store32(m.memory[int64(uint32(v0))+4:], uint32(v12))
		v3 = i32(1)
		goto l0
	}
l0:
	store32(m.memory[uint32(v0):], uint32(v3))
	m.g0 = v2 + i32(16)
}
func (m *Module) fn517(v0, v1 int32) {
	var v2, v3, v4, v5 int32
	v2 = i32(0)
	{
		{
			t0 := int32(m.memory[int64(uint32(v1))+37])
			if t0 == 0 {
				goto l0
			}
			goto l1
		}
	l0:
		m.memory[int64(uint32(v1))+37] = byte(i32(1))
		{
			{
				t1 := int32(m.memory[int64(uint32(v1))+36])
				if t1 != i32(1) {
					goto l2
				}
				t2 := int32(load32(m.memory[int64(uint32(v1))+32:]))
				v3 = t2
				t3 := int32(load32(m.memory[int64(uint32(v1))+28:]))
				v4 = t3
				goto l3
			}
		l2:
			t4 := int32(load32(m.memory[int64(uint32(v1))+32:]))
			v3 = t4
			t5 := int32(load32(m.memory[int64(uint32(v1))+28:]))
			t6 := v3
			v4 = t5
			if t6 == v4 {
				goto l1
			}
		}
	l3:
		v5 = v3 - v4
		t7 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		v2 = t7 + v4
	}
l1:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
	store32(m.memory[uint32(v0):], uint32(v2))
}
func (m *Module) fn518(v0, v1, v2, v3, v4 int32) {
	if uint32(v3) <= uint32(v2) {
		goto l0
	}
	m.fn151(i32(0), v3, v2, v4)
	panic("unreachable")
l0:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
	store32(m.memory[uint32(v0):], uint32(v1))
}
func (m *Module) fn519(v0, v1 int32) {
	var v2, v3, v4 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	{
		{
			t1 := int32(load32(m.memory[int64(uint32(v1))+16:]))
			v3 = t1
			if v3 != 0 {
				goto l0
			}
			v1 = i32(0)
			v3 = i32(0)
			goto l1
		}
	l0:
		t2 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t3 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		m.fn488(v2+i32(4), t2, t3)
		t4 := int32(load32(m.memory[int64(uint32(v2))+4:]))
		t5 := v3
		v1 = t4
		p6 := v1
		if uint32(v3) < uint32(v1) {
			p6 = t5
		}
		v1 = p6
		t7 := int32(load32(m.memory[int64(uint32(v2))+12:]))
		v4 = t7
		p8 := v3
		if uint32(v4) < uint32(v3) {
			p8 = v4
		}
		t9 := int32(load32(m.memory[int64(uint32(v2))+8:]))
		p10 := v3
		if t9 != 0 {
			p10 = p8
		}
		v3 = p10
	}
l1:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v3))
	store32(m.memory[int64(uint32(v0))+4:], uint32(i32(1)))
	store32(m.memory[uint32(v0):], uint32(v1))
	m.g0 = v2 + i32(16)
}
func (m *Module) fn520(v0, v1, v2, v3 int32) {
	var v4, v5, v6 int32
	t0 := m.g0
	v4 = t0 - i32(16)
	m.g0 = v4
	{
		if uint32(v3) > uint32(i32(7)) {
			goto l0
		}
		v5 = i32(0)
		v6 = v1 & i32(255)
		v1 = i32(0)
	l4:
		if v3 != v1 {
			t1 := int32(m.memory[uint32(v2+v1)])
			if t1 != v6 {
				v1 = v1 + i32(1)
				goto l4
			}
			v5 = i32(1)
			goto l2
		}
		v1 = v3
		goto l2
	l0:
		m.fn521(v4+i32(8), v1, v2, v3)
		t2 := int32(load32(m.memory[int64(uint32(v4))+12:]))
		v1 = t2
		t3 := int32(load32(m.memory[int64(uint32(v4))+8:]))
		v5 = t3
	}
l2:
	store32(m.memory[uint32(v0):], uint32(v5))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
	m.g0 = v4 + i32(16)
}
func (m *Module) fn521(v0, v1, v2, v3 int32) {
	var v4, v5, v6, v7, v8 int32
	{
		v4 = (v2 + i32(3)) & i32(-4)
		if v4 == v2 {
			goto l0
		}
		v4 = v4 - v2
		v5 = i32(0)
		v6 = v1 & i32(255)
		v7 = i32(1)
	l2:
		{
			t0 := int32(m.memory[uint32(v2+v5)])
			if t0 == v6 {
				goto l1
			}
			t1 := v4
			v5 = v5 + i32(1)
			if t1 != v5 {
				goto l2
			}
		}
		t2 := v4
		v8 = v3 + i32(-8)
		if uint32(t2) > uint32(v8) {
			goto l3
		}
		goto l4
	}
l0:
	v8 = v3 + i32(-8)
	v4 = i32(0)
l4:
	v5 = v1 & i32(255) * i32(16843009)
l5:
	{
		v6 = v2 + v4
		t3 := int32(load32(m.memory[uint32(v6):]))
		v7 = t3 ^ v5
		t4 := int32(load32(m.memory[uint32(v6+i32(4)):]))
		t5 := i32(16843008) - v7 | v7
		v6 = t4 ^ v5
		if t5&(i32(16843008)-v6|v6)&i32(-2139062144) != i32(-2139062144) {
			goto l3
		}
		v4 = v4 + i32(8)
		if uint32(v4) <= uint32(v8) {
			goto l5
		}
	}
l3:
	if v3 == v4 {
		goto l6
	}
	v7 = v3 - v4
	v2 = v2 + v4
	v5 = i32(0)
	v6 = v1 & i32(255)
l8:
	{
		t6 := int32(m.memory[uint32(v2+v5)])
		if t6 == v6 {
			v5 = v5 + v4
			v7 = i32(1)
			goto l1
		}
		t7 := v7
		v5 = v5 + i32(1)
		if t7 == v5 {
			goto l6
		}
		goto l8
	}
l6:
	v7 = i32(0)
l1:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
	store32(m.memory[uint32(v0):], uint32(v7))
}
func (m *Module) fn522(v0, v1, v2 int32) {
	var v3 int32
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
		p0 := i32(4)
		if uint32(v1) < uint32(i32(65536)) {
			p0 = i32(3)
		}
		v3 = p0
	}
l1:
	m.fn279(v1, v2)
	store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
	store32(m.memory[uint32(v0):], uint32(v2))
}
func (m *Module) fn523(v0, v1, v2, v3 int32) {
	var v4 int32
	var v5 int64
	var v6, v7, v8, v9, v10 int32
	var v11 int64
	t0 := m.g0
	v4 = t0 - i32(48)
	m.g0 = v4
	t1 := int64(load64(m.memory[int64(uint32(v1))+16:]))
	t2 := int64(load64(m.memory[int64(uint32(v1))+24:]))
	t3 := int32(load32(m.memory[int64(uint32(v2))+4:]))
	t4 := int32(load32(m.memory[int64(uint32(v2))+8:]))
	t5 := m.fn524(t1, t2, t3, t4)
	v5 = t5
	store32(m.memory[int64(uint32(v4))+44:], uint32(v2))
	{
		t6 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		if t6 != 0 {
			goto l0
		}
		_ = m.fn525(v1, v1+i32(16))
	}
l0:
	store32(m.memory[int64(uint32(v4))+16:], uint32(v4+i32(44)))
	store32(m.memory[int64(uint32(v4))+20:], uint32(v1))
	t8 := int32(load32(m.memory[uint32(v1):]))
	t9 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	m.fn69(v4+i32(8), t8, t9, v5, v4+i32(16), i32(88))
	t10 := int32(load32(m.memory[uint32(v1):]))
	v6 = t10
	t11 := int32(load32(m.memory[int64(uint32(v4))+12:]))
	v7 = t11
	{
		{
			t12 := int32(load32(m.memory[int64(uint32(v4))+8:]))
			if t12 != i32(1) {
				goto l1
			}
			v8 = v6 + v7
			t13 := int32(m.memory[uint32(v8)])
			v9 = t13
			t14 := int32(load32(m.memory[int64(uint32(v2))+8:]))
			v10 = t14
			t15 := int64(load64(m.memory[uint32(v2):]))
			v11 = t15
			t16 := v8
			v2 = int32(uint32(int32(v5)) >> 25)
			m.memory[uint32(t16)] = byte(v2)
			t17 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			m.memory[uint32(v6+t17&(v7+i32(-8))+i32(8))] = byte(v2)
			t18 := int32(load32(m.memory[int64(uint32(v1))+12:]))
			store32(m.memory[int64(uint32(v1))+12:], uint32(t18+i32(1)))
			t19 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			store32(m.memory[int64(uint32(v1))+8:], uint32(t19-v9&i32(1)))
			v1 = v6 + (i32(0)-v7)*i32(24) + i32(-24)
			store64(m.memory[uint32(v1):], uint64(v11))
			store32(m.memory[int64(uint32(v4))+24:], uint32(v10))
			t20 := int64(load64(m.memory[uint32(v3):]))
			store64(m.memory[int64(uint32(v4))+28:], uint64(t20))
			t21 := int64(load64(m.memory[int64(uint32(v4))+24:]))
			store64(m.memory[int64(uint32(v1))+8:], uint64(t21))
			t22 := int32(load32(m.memory[int64(uint32(v3))+8:]))
			store32(m.memory[int64(uint32(v4))+36:], uint32(t22))
			t23 := int64(load64(m.memory[int64(uint32(v4))+32:]))
			store64(m.memory[int64(uint32(v1))+16:], uint64(t23))
			store32(m.memory[uint32(v0):], uint32(i32(-1)))
			goto l2
		}
	l1:
		v1 = v6 + (i32(0)-v7)*i32(24) + i32(-12)
		t24 := int64(load64(m.memory[uint32(v1):]))
		v5 = t24
		t25 := int64(load64(m.memory[uint32(v3):]))
		store64(m.memory[uint32(v1):], uint64(t25))
		t26 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		v6 = t26
		t27 := int32(load32(m.memory[int64(uint32(v3))+8:]))
		store32(m.memory[int64(uint32(v1))+8:], uint32(t27))
		store32(m.memory[int64(uint32(v0))+8:], uint32(v6))
		store64(m.memory[uint32(v0):], uint64(v5))
		t28 := int32(load32(m.memory[uint32(v2):]))
		t29 := int32(load32(m.memory[int64(uint32(v2))+4:]))
		m.fn16(t28, t29)
	}
l2:
	m.g0 = v4 + i32(48)
}
func (m *Module) fn524(v0, v1 int64, v2, v3 int32) int64 {
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
	m.fn172(v4, v3)
	m.fn285(v4, v2, v3)
	t1 := m.fn174(v4)
	v1 = t1
	m.g0 = v4 + i32(64)
	return v1
}
func (m *Module) fn525(v0, v1 int32) int32 {
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
				t19 := m.fn527(t17, t18, v10)
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
	m.fn241(v0, v2+i32(16), i32(89), i32(24))
l7:
	v5 = i32(-1)
l2:
	m.g0 = v2 + i32(64)
	return v5
}
func (m *Module) fn526(v0, v1 int32) int32 {
	var v2 int32
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(t0):]))
	v2 = t1
	t2 := int32(load32(m.memory[uint32(v2+i32(4)):]))
	t3 := int32(load32(m.memory[uint32(v2+i32(8)):]))
	t4 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t5 := int32(load32(m.memory[uint32(t4):]))
	v0 = t5 + (i32(0)-v1)*i32(24)
	t6 := int32(load32(m.memory[uint32(v0+i32(-20)):]))
	t7 := int32(load32(m.memory[uint32(v0+i32(-16)):]))
	t8 := m.fn191(t2, t3, t6, t7)
	return t8
}
func (m *Module) fn527(v0, v1, v2 int32) int64 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(t0):]))
	v0 = t1
	t2 := int64(load64(m.memory[uint32(v0):]))
	t3 := int64(load64(m.memory[uint32(v0+i32(8)):]))
	t4 := int32(load32(m.memory[uint32(v1):]))
	v0 = t4 + (i32(0)-v2)*i32(24)
	t5 := int32(load32(m.memory[uint32(v0+i32(-20)):]))
	t6 := int32(load32(m.memory[uint32(v0+i32(-16)):]))
	t7 := m.fn524(t2, t3, t5, t6)
	return t7
}
func (m *Module) fn528(v0, v1, v2, v3, v4, v5 int32) {
	var v6, v7, v8, v9 int32
	var v10 int64
	t0 := m.g0
	v6 = t0 - i32(16)
	m.g0 = v6
	v7 = v4 * i32(6)
	v8 = v3 + i32(-6)
	v9 = v2 & i32(0xffff)
l7:
	{
		m.fn505(v6+i32(8), v1)
		{
			{
				t1 := int32(m.memory[int64(uint32(v6))+8])
				if t1 != i32(255) {
					goto l0
				}
				t2 := int32(load16(m.memory[int64(uint32(v6))+10:]))
				v4 = t2
				goto l1
			}
		l0:
			t3 := int64(load64(m.memory[int64(uint32(v6))+8:]))
			v10 = t3
			if v10&i64(255) != i64(255) {
				store64(m.memory[int64(uint32(v0))+4:], uint64(v10))
				goto l13
			}
			v4 = int32(int64(uint64(v10) >> 16))
		}
	l1:
		m.fn507(v6+i32(8), v1, v5)
		{
			{
				t4 := int32(m.memory[int64(uint32(v6))+8])
				if t4 != i32(255) {
					goto l3
				}
				t5 := int32(load32(m.memory[int64(uint32(v6))+12:]))
				v3 = t5
				goto l4
			}
		l3:
			t6 := int64(load64(m.memory[int64(uint32(v6))+8:]))
			v10 = t6
			if v10&i64(255) != i64(255) {
				store64(m.memory[int64(uint32(v0))+4:], uint64(v10))
				goto l13
			}
			v3 = int32(int64(uint64(v10) >> 32))
		}
	l4:
		v2 = v4 & i32(0xffff)
		if v2 == v9 {
			store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
			v4 = i32(-1)
			goto l15
		}
		v4 = v7
		v3 = v8
	l8:
		{
			if v4 == 0 {
				goto l7
			}
			v4 = v4 + i32(-6)
			v3 = v3 + i32(6)
			t7 := int32(load16(m.memory[uint32(v3):]))
			if t7 != v2 {
				goto l8
			}
		}
		t8 := int32(load16(m.memory[int64(uint32(v3))+4:]))
		v2 = t8
		t9 := int32(load16(m.memory[int64(uint32(v3))+2:]))
		if t9 != i32(1) {
			goto l7
		}
	l14:
		{
			m.fn505(v6+i32(8), v1)
			{
				{
					t10 := int32(m.memory[int64(uint32(v6))+8])
					if t10 != i32(255) {
						goto l9
					}
					t11 := int32(load16(m.memory[int64(uint32(v6))+10:]))
					v4 = t11
					goto l10
				}
			l9:
				t12 := int64(load64(m.memory[int64(uint32(v6))+8:]))
				v10 = t12
				if v10&i64(255) != i64(255) {
					goto l11
				}
				v4 = int32(int64(uint64(v10) >> 16))
			}
		l10:
			{
				if v4&i32(0xffff) != v2&i32(0xffff) {
					goto l12
				}
				m.fn507(v6+i32(8), v1, v5)
				t13 := int32(m.memory[int64(uint32(v6))+8])
				if t13 == i32(255) {
					goto l7
				}
				t14 := int64(load64(m.memory[int64(uint32(v6))+8:]))
				v10 = t14
				if v10&i64(255) == i64(255) {
					goto l7
				}
				store64(m.memory[int64(uint32(v0))+4:], uint64(v10))
				goto l13
			}
		l12:
			m.fn507(v6+i32(8), v1, v5)
			t15 := int32(m.memory[int64(uint32(v6))+8])
			if t15 == i32(255) {
				goto l14
			}
			t16 := int64(load64(m.memory[int64(uint32(v6))+8:]))
			v10 = t16
			if v10&i64(255) == i64(255) {
				goto l14
			}
		}
	}
	store64(m.memory[int64(uint32(v0))+4:], uint64(v10))
	goto l13
l11:
	store64(m.memory[int64(uint32(v0))+4:], uint64(v10))
l13:
	v4 = i32(-0x7ffffff1)
l15:
	store32(m.memory[uint32(v0):], uint32(v4))
	m.g0 = v6 + i32(16)
}
func (m *Module) fn529(v0, v1 int64, v2 int32) int64 {
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
	m.fn530(v2, v3)
	t1 := m.fn174(v3)
	v1 = t1
	m.g0 = v3 + i32(64)
	return v1
}
func (m *Module) fn530(v0, v1 int32) {
	var v2 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	store16(m.memory[int64(uint32(v2))+14:], uint16(v0))
	m.fn285(v1, v2+i32(14), i32(2))
	m.g0 = v2 + i32(16)
}
func (m *Module) fn531(v0, v1 int32) {
	var v2 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		v2 = t0
		t1 := int32(load32(m.memory[uint32(v0):]))
		if v2 != t1 {
			goto l0
		}
		m.fn532(v0)
	}
l0:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v2+i32(1)))
	t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	m.memory[uint32(t2+v2)] = byte(v1)
}
func (m *Module) fn532(v0 int32) {
	var v1 int32
	t0 := m.g0
	v1 = t0 - i32(16)
	m.g0 = v1
	t1 := int32(load32(m.memory[uint32(v0):]))
	m.fn273(v1+i32(8), v0, t1, i32(1), i32(1), i32(1))
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
func (m *Module) fn533(v0 int32) {
	{
		t0 := int64(load64(m.memory[int64(uint32(v0))+24:]))
		if t0 == i64(-1) {
			goto l0
		}
		m.fn228(v0)
		return
	}
l0:
	m.fn534(v0)
}
func (m *Module) fn534(v0 int32) {
	var v1 int32
	{
		{
			t0 := int32(load32(m.memory[uint32(v0):]))
			v1 = t0
			p1 := i32(2)
			if uint32(v1) > uint32(i32(-0x7ffffff2)) {
				p1 = v1 + i32(0x7ffffff1)
			}
			v1 = p1
			switch v1 {
			case 3, 5:
				return
			default:
				switch v1 + i32(-15) {
				case 0:
					t6 := int32(load32(m.memory[int64(uint32(v0))+4:]))
					t7 := int32(load32(m.memory[int64(uint32(v0))+8:]))
					m.fn16(t6, t7)
					return
				case 2:
					goto l8
				default:
					return
				}
			case 0:
				t2 := int32(m.memory[int64(uint32(v0))+4])
				t3 := int32(load32(m.memory[int64(uint32(v0))+8:]))
				m.fn119(t2, t3)
				return
			case 1:
				m.fn116(v0 + i32(4))
				return
			case 2:
				m.fn535(v0)
				return
			case 4:
				m.fn451(v0 + i32(4))
				return
			case 6:
				t4 := int32(load32(m.memory[int64(uint32(v0))+4:]))
				t5 := int32(load32(m.memory[int64(uint32(v0))+8:]))
				m.fn16(t4, t5)
				return
			}
		}
	l8:
		t8 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		t9 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		m.fn16(t8, t9)
	}
}
func (m *Module) fn535(v0 int32) {
	var v1 int32
	{
		t0 := int32(load32(m.memory[uint32(v0):]))
		v1 = t0
		p1 := i32(2)
		if uint32(v1) > uint32(i32(-0x7ffffff9)) {
			p1 = v1 + i32(0x7ffffff8)
		}
		switch p1 {
		case 1, 3, 4:
			return
		default:
			t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			if uint32(t2) > uint32(i32(4)) {
				return
			}
			t3 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			t4 := int32(load32(m.memory[int64(uint32(v0))+12:]))
			m.fn16(t3, t4)
			return
		case 0:
			t5 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			v1 = t5
			t6 := int32(load32(m.memory[uint32(v1):]))
			t7 := v1
			v1 = t6
			store32(m.memory[uint32(t7):], uint32(v1+i32(-1)))
			if v1 != i32(1) {
				return
			}
			t8 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			m.fn536(t8)
			return
		case 2:
			p9 := i32(5)
			if v1 < i32(0) {
				p9 = v1 ^ i32(-0x80000000)
			}
			switch p9 {
			default:
				return
			case 0:
				t10 := int32(load32(m.memory[int64(uint32(v0))+4:]))
				t11 := int32(load32(m.memory[int64(uint32(v0))+8:]))
				m.fn134(t10, t11)
				return
			case 3:
				t12 := int32(load32(m.memory[int64(uint32(v0))+4:]))
				t13 := int32(load32(m.memory[int64(uint32(v0))+8:]))
				m.fn16(t12, t13)
				return
			case 4:
				t14 := int32(load32(m.memory[int64(uint32(v0))+4:]))
				t15 := int32(load32(m.memory[int64(uint32(v0))+8:]))
				m.fn16(t14, t15)
				return
			case 5:
				t16 := int32(load32(m.memory[int64(uint32(v0))+4:]))
				m.fn16(v1, t16)
				t17 := int32(load32(m.memory[int64(uint32(v0))+12:]))
				t18 := int32(load32(m.memory[int64(uint32(v0))+16:]))
				m.fn16(t17, t18)
				return
			}
		case 5:
			t19 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			v1 = t19
			if v1 < i32(0) {
				return
			}
			t20 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			m.fn16(v1, t20)
			return
		}
	}
}
func (m *Module) fn536(v0 int32) {
	var v1 int32
	t0 := int32(m.memory[int64(uint32(v0))+8])
	t1 := int32(load32(m.memory[uint32(v0+i32(12)):]))
	m.fn1565(t0, t1)
	{
		if v0 == i32(-1) {
			return
		}
		t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		t3 := v0
		v1 = t2
		store32(m.memory[int64(uint32(t3))+4:], uint32(v1+i32(-1)))
		if v1 != i32(1) {
			return
		}
		m.fn10(v0, i32(16), i32(4))
	}
}
