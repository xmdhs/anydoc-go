package core

import (
	"math/bits"
)

func (m *Module) fn762(v0 int32) {
	var v1, v2 int32
	v1 = i32(0)
l1:
	{
		if v1 == i32(108) {
			return
		}
		v2 = v0 + v1
		t0 := int32(load32(m.memory[uint32(v2):]))
		t1 := int32(load32(m.memory[uint32(v2+i32(4)):]))
		m.fn134(t0, t1)
		v1 = v1 + i32(12)
		goto l1
	}
}
func (m *Module) fn763(v0 int32) {
	var v1, v2, v3, v4 int32
	t0 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	v1 = t0
	t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	v2 = t1
	v3 = v2
l2:
	if v1 == 0 {
		goto l0
	}
	{
		t2 := int32(load32(m.memory[uint32(v3):]))
		v4 = t2
		if v4 == i32(-1) {
			goto l1
		}
		t3 := int32(load32(m.memory[uint32(v3+i32(4)):]))
		m.fn16(v4, t3)
	}
l1:
	v1 = v1 + i32(-1)
	v3 = v3 + i32(12)
	goto l2
l0:
	t4 := int32(load32(m.memory[uint32(v0):]))
	m.fn136(t4, v2, i32(4), i32(12))
}
func (m *Module) fn764(v0 int32) {
	var v1 int32
	v1 = i32(0)
l1:
	if v1 == i32(288) {
		return
	}
	m.fn763(v0 + v1)
	v1 = v1 + i32(32)
	goto l1
}
func (m *Module) fn765(v0, v1 int32) {
	m.fn136(v0, v1, i32(1), i32(3))
}
func (m *Module) fn766(v0 int32) {
	t0 := int32(load32(m.memory[uint32(v0):]))
	if t0 == i32(-1) {
		return
	}
	m.fn767(v0)
}
func (m *Module) fn767(v0 int32) {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	m.fn768(t0, t1)
	t2 := int32(load32(m.memory[int64(uint32(v0))+12:]))
	t3 := int32(load32(m.memory[int64(uint32(v0))+16:]))
	m.fn136(t2, t3, i32(1), i32(4))
}
func (m *Module) fn768(v0, v1 int32) {
	m.fn136(v0, v1, i32(2), i32(2))
}
func (m *Module) fn769(v0 int32) int32 {
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
		v1 = v1 + i32(-288)
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
	return v1 + (i32(0)-int32(uint32(int64(bits.TrailingZeros64(uint64(v3))))>>3))*i32(36)
}
func (m *Module) fn770(v0 int32) {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	m.fn16(t0, t1)
	t2 := int32(load32(m.memory[int64(uint32(v0))+12:]))
	t3 := int32(load32(m.memory[int64(uint32(v0))+16:]))
	m.fn16(t2, t3)
}
func (m *Module) fn771(v0 int32) {
	m.fn761(v0)
	m.fn762(v0 + i32(360))
}
func (m *Module) fn772(v0, v1, v2 int32) int32 {
	var v3 int32
	v3 = i32(0)
	{
		t0 := int32(load32(m.memory[uint32(v2+i32(16)):]))
		t1 := int32(load32(m.memory[uint32(v2+i32(20)):]))
		t2 := m.fn773(t0, t1, v0, v1)
		if t2 == 0 {
			goto l0
		}
		t3 := int32(m.memory[int64(uint32(v2))+24])
		var p4 int32
		if t3 == 0 {
			p4 = 1
		}
		v3 = p4
	}
l0:
	return v3
}
func (m *Module) fn773(v0, v1, v2, v3 int32) int32 {
	var v4 int32
	v4 = i32(0)
	{
		if v1 != v3 {
			goto l0
		}
		t0 := m.fn1851(v0, v2, v1)
		var p1 int32
		if t0 == 0 {
			p1 = 1
		}
		v4 = p1
	}
l0:
	return v4
}
func (m *Module) fn774(v0, v1, v2, v3, v4 int32) {
	var v5, v6, v7 int32
	var v8 int64
	var v9, v10, v11 int32
	t0 := m.g0
	v5 = t0 - i32(112)
	m.g0 = v5
	m.fn775(v5+i32(60), v3, v4, i32(35))
	{
		{
			t1 := int32(load32(m.memory[int64(uint32(v5))+60:]))
			v6 = t1
			if v6 != 0 {
				goto l0
			}
			v7 = i32(-1)
			goto l1
		}
	l0:
		t2 := int32(load32(m.memory[int64(uint32(v5))+64:]))
		v4 = t2
		t3 := int32(load32(m.memory[int64(uint32(v5))+68:]))
		t4 := int32(load32(m.memory[int64(uint32(v5))+72:]))
		m.fn776(v5+i32(36), t3, t4)
		t5 := int64(load64(m.memory[int64(uint32(v5))+40:]))
		v8 = t5
		t6 := int32(load32(m.memory[int64(uint32(v5))+36:]))
		v7 = t6
		v3 = v6
	}
l1:
	m.fn775(v5+i32(60), v3, v4, i32(63))
	v9 = int32(int64(uint64(v8) >> 32))
	v10 = int32(v8)
	{
		{
			t7 := int32(load32(m.memory[int64(uint32(v5))+64:]))
			t8 := int32(load32(m.memory[int64(uint32(v5))+60:]))
			t9 := v4
			v6 = t8
			p10 := t9
			if v6 != 0 {
				p10 = t7
			}
			v11 = p10
			if v11 == 0 {
				m.fn51(v5+i32(60), v1, v2)
				store32(m.memory[uint32(v0):], uint32(i32(0)))
				t23 := int64(load64(m.memory[int64(uint32(v5))+60:]))
				store64(m.memory[int64(uint32(v0))+4:], uint64(t23))
				store32(m.memory[int64(uint32(v5))+72:], uint32(v7))
				t24 := int64(load64(m.memory[int64(uint32(v5))+68:]))
				store64(m.memory[int64(uint32(v0))+12:], uint64(t24))
				store32(m.memory[int64(uint32(v5))+80:], uint32(v9))
				store32(m.memory[int64(uint32(v5))+76:], uint32(v10))
				t25 := int64(load64(m.memory[int64(uint32(v5))+76:]))
				store64(m.memory[int64(uint32(v0))+20:], uint64(t25))
				goto l6
			}
			store32(m.memory[int64(uint32(v5))+32:], uint32(i32(0)))
			store64(m.memory[int64(uint32(v5))+24:], uint64(i64(0x400000000)))
			p11 := v3
			if v6 != 0 {
				p11 = v6
			}
			v6 = p11
			t12 := m.fn777(v6, v11, i32(47))
			if t12 != 0 {
				goto l3
			}
			m.fn778(v5+i32(36), v1, v2, i32(47))
			t13 := int32(load32(m.memory[int64(uint32(v5))+36:]))
			v4 = t13
			if v4 == 0 {
				goto l3
			}
			t14 := int32(load32(m.memory[int64(uint32(v5))+40:]))
			m.fn513(v5+i32(60), v4, t14, i32(47))
			v2 = i32(0)
			v1 = i32(4)
		l4:
			{
				m.fn515(v5+i32(16), v5+i32(60))
				t15 := int32(load32(m.memory[int64(uint32(v5))+16:]))
				v4 = t15
				if v4 == 0 {
					goto l3
				}
				t16 := int32(load32(m.memory[int64(uint32(v5))+20:]))
				v3 = t16
				if v3 == 0 {
					goto l4
				}
				m.fn51(v5+i32(100), v4, v3)
				t17 := int32(load32(m.memory[int64(uint32(v5))+100:]))
				if t17 == i32(-1) {
					goto l3
				}
				{
					t18 := int32(load32(m.memory[int64(uint32(v5))+24:]))
					if v2 != t18 {
						goto l5
					}
					m.fn60(v5+i32(24), i32(1))
					t19 := int32(load32(m.memory[int64(uint32(v5))+28:]))
					v1 = t19
				}
			l5:
				v4 = v1 + v2*i32(12)
				t20 := int32(load32(m.memory[int64(uint32(v5))+108:]))
				store32(m.memory[int64(uint32(v4))+8:], uint32(t20))
				t21 := int64(load64(m.memory[int64(uint32(v5))+100:]))
				store64(m.memory[uint32(v4):], uint64(t21))
				t22 := v5
				v2 = v2 + i32(1)
				store32(m.memory[int64(uint32(t22))+32:], uint32(v2))
				goto l4
			}
		}
	l3:
		m.fn513(v5+i32(60), v6, v11, i32(47))
	l8:
		m.fn515(v5+i32(8), v5+i32(60))
		{
			t26 := int32(load32(m.memory[int64(uint32(v5))+8:]))
			v4 = t26
			if v4 == 0 {
				t39 := int32(load32(m.memory[int64(uint32(v5))+28:]))
				t40 := int32(load32(m.memory[int64(uint32(v5))+32:]))
				m.fn77(v5+i32(60), t39, t40, i32(1101983), i32(1))
				store32(m.memory[uint32(v0):], uint32(i32(0)))
				t41 := int64(load64(m.memory[int64(uint32(v5))+60:]))
				store64(m.memory[int64(uint32(v0))+4:], uint64(t41))
				store32(m.memory[int64(uint32(v5))+72:], uint32(v7))
				t42 := int64(load64(m.memory[int64(uint32(v5))+68:]))
				store64(m.memory[int64(uint32(v0))+12:], uint64(t42))
				store32(m.memory[int64(uint32(v5))+80:], uint32(v9))
				store32(m.memory[int64(uint32(v5))+76:], uint32(v10))
				t43 := int64(load64(m.memory[int64(uint32(v5))+76:]))
				store64(m.memory[int64(uint32(v0))+20:], uint64(t43))
				m.fn78(v5 + i32(24))
				goto l6
			}
			t27 := int32(load32(m.memory[int64(uint32(v5))+12:]))
			t28 := v4
			v3 = t27
			t29 := m.fn15(t28, v3, i32(1), i32(0))
			if t29 != 0 {
				goto l8
			}
			t30 := m.fn15(v4, v3, i32(1109519), i32(1))
			if t30 != 0 {
				goto l8
			}
			t31 := m.fn15(v4, v3, i32(1284184), i32(2))
			if t31 != 0 {
				{
					{
						t44 := int32(load32(m.memory[int64(uint32(v5))+32:]))
						v4 = t44
						if v4 != 0 {
							goto l12
						}
						v4 = i32(-1)
						goto l13
					}
				l12:
					t45 := v5
					v4 = v4 + i32(-1)
					store32(m.memory[int64(uint32(t45))+32:], uint32(v4))
					t46 := int32(load32(m.memory[int64(uint32(v5))+28:]))
					v4 = t46 + v4*i32(12)
					t47 := int32(load32(m.memory[int64(uint32(v4))+4:]))
					v6 = t47
					t48 := int32(load32(m.memory[uint32(v4):]))
					v4 = t48
				}
			l13:
				m.fn134(v4, v6)
				goto l8
			}
			store32(m.memory[int64(uint32(v5))+56:], uint32(v3))
			store32(m.memory[int64(uint32(v5))+52:], uint32(v4))
			m.fn776(v5+i32(36), v4, v3)
			t32 := int32(load32(m.memory[int64(uint32(v5))+40:]))
			v4 = t32
			t33 := int32(load32(m.memory[int64(uint32(v5))+44:]))
			t34 := v4
			v3 = t33
			t35 := m.fn779(t34, v3, i32(47))
			if t35 != 0 {
				goto l10
			}
			t36 := m.fn779(v4, v3, i32(92))
			if t36 != 0 {
				goto l10
			}
			t37 := m.fn773(v4, v3, i32(1109519), i32(1))
			if t37 != 0 {
				goto l11
			}
			t38 := m.fn773(v4, v3, i32(1284184), i32(2))
			if t38 != 0 {
				goto l11
			}
			m.fn33(v5+i32(24), v5+i32(36))
			goto l8
		}
	l11:
		store32(m.memory[int64(uint32(v5))+104:], uint32(i32(71)))
		store32(m.memory[int64(uint32(v5))+100:], uint32(v5+i32(52)))
		m.fn73(v0+i32(4), i32(1049627), v5+i32(100))
		goto l14
	l10:
		store32(m.memory[int64(uint32(v5))+104:], uint32(i32(71)))
		store32(m.memory[int64(uint32(v5))+100:], uint32(v5+i32(52)))
		m.fn73(v0+i32(4), i32(1049569), v5+i32(100))
	l14:
		store32(m.memory[uint32(v0):], uint32(i32(1)))
		store32(m.memory[int64(uint32(v0))+16:], uint32(i32(-1)))
		t49 := int32(load32(m.memory[int64(uint32(v5))+36:]))
		m.fn16(t49, v4)
		m.fn78(v5 + i32(24))
		m.fn134(v7, v10)
	}
l6:
	m.g0 = v5 + i32(112)
}
func (m *Module) fn775(v0, v1, v2, v3 int32) {
	var v4, v5 int32
	t0 := m.g0
	v4 = t0 - i32(48)
	m.g0 = v4
	m.fn514(v4+i32(20), v3, v1, v2)
	m.fn516(v4+i32(8), v4+i32(20))
	{
		t1 := int32(load32(m.memory[int64(uint32(v4))+8:]))
		if t1 != i32(1) {
			goto l0
		}
		t2 := int32(load32(m.memory[int64(uint32(v4))+12:]))
		v3 = t2
		t3 := int32(load32(m.memory[int64(uint32(v4))+16:]))
		t4 := v0
		t5 := v2
		v5 = t3
		store32(m.memory[int64(uint32(t4))+12:], uint32(t5-v5))
		store32(m.memory[int64(uint32(v0))+8:], uint32(v1+v5))
		store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
		goto l1
	}
l0:
	v1 = i32(0)
l1:
	store32(m.memory[uint32(v0):], uint32(v1))
	m.g0 = v4 + i32(48)
}
func (m *Module) fn776(v0, v1, v2 int32) {
	var v3, v4, v5, v6, v7, v8, v9, v10 int32
	t0 := m.g0
	v3 = t0 - i32(48)
	m.g0 = v3
	m.fn140(v3+i32(24), v2)
	v4 = i32(0)
l4:
	if uint32(v4) >= uint32(v2) {
		goto l0
	}
	v5 = i32(1)
	{
		t1 := int32(m.memory[uint32(v1+v4)])
		v6 = t1
		if v6 != i32(37) {
			goto l1
		}
		m.fn786(v3+i32(16), v4, v1, v2)
		v6 = i32(37)
		t2 := int32(load32(m.memory[int64(uint32(v3))+16:]))
		v7 = t2
		if v7 == 0 {
			goto l1
		}
		v5 = i32(1)
		t3 := int32(load32(m.memory[int64(uint32(v3))+20:]))
		m.fn787(v3+i32(8), i32(1), i32(3), v7, t3)
		t4 := int32(load32(m.memory[int64(uint32(v3))+8:]))
		v8 = t4
		if v8 == 0 {
			goto l1
		}
		t5 := int32(load32(m.memory[int64(uint32(v3))+12:]))
		v9 = t5
		v7 = i32(0)
	l3:
		{
			if v9 == v7 {
				goto l2
			}
			v10 = v8 + v7
			v5 = i32(1)
			v7 = v7 + i32(1)
			t6 := int32(m.memory[uint32(v10)])
			v10 = t6
			if uint32((v10+i32(-58))&i32(255)) > uint32(i32(245)) {
				goto l3
			}
			if uint32((v10&i32(-33)+i32(-71))&i32(255)) >= uint32(i32(250)) {
				goto l3
			}
			goto l1
		}
	l2:
		m.fn788(v3, v8, v9, i32(16))
		t7 := int32(m.memory[uint32(v3)])
		v7 = t7
		p8 := i32(3)
		if v7 != 0 {
			p8 = i32(1)
		}
		v5 = p8
		t9 := int32(m.memory[int64(uint32(v3))+1])
		p10 := t9
		if v7 != 0 {
			p10 = i32(37)
		}
		v6 = p10
	}
l1:
	m.fn145(v3+i32(24), v6)
	v4 = v4 + v5
	goto l4
l0:
	t11 := int32(load32(m.memory[int64(uint32(v3))+28:]))
	t12 := v3 + i32(36)
	v7 = t11
	t13 := int32(load32(m.memory[int64(uint32(v3))+32:]))
	m.fn92(t12, v7, t13)
	m.fn490(v0, v3+i32(36))
	t14 := int32(load32(m.memory[int64(uint32(v3))+24:]))
	m.fn16(t14, v7)
	m.g0 = v3 + i32(48)
}
func (m *Module) fn777(v0, v1, v2 int32) int32 {
	var v3 int32
	t0 := m.g0
	v3 = t0 - i32(16)
	m.g0 = v3
	store32(m.memory[int64(uint32(v3))+12:], uint32(i32(0)))
	m.memory[int64(uint32(v3))+12] = byte(v2)
	t1 := m.fn159(v0, v1, v3+i32(12), i32(1))
	v2 = t1
	m.g0 = v3 + i32(16)
	return v2
}
func (m *Module) fn778(v0, v1, v2, v3 int32) {
	var v4, v5 int32
	t0 := m.g0
	v4 = t0 - i32(48)
	m.g0 = v4
	m.fn514(v4+i32(20), v3, v1, v2)
	m.fn554(v4+i32(8), v4+i32(20))
	{
		t1 := int32(load32(m.memory[int64(uint32(v4))+8:]))
		if t1 != i32(1) {
			goto l0
		}
		t2 := int32(load32(m.memory[int64(uint32(v4))+12:]))
		v3 = t2
		t3 := int32(load32(m.memory[int64(uint32(v4))+16:]))
		t4 := v0
		t5 := v2
		v5 = t3
		store32(m.memory[int64(uint32(t4))+12:], uint32(t5-v5))
		store32(m.memory[int64(uint32(v0))+8:], uint32(v1+v5))
		store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
		goto l1
	}
l0:
	v1 = i32(0)
l1:
	store32(m.memory[uint32(v0):], uint32(v1))
	m.g0 = v4 + i32(48)
}
func (m *Module) fn779(v0, v1, v2 int32) int32 {
	var v3 int32
	t0 := m.g0
	v3 = t0 - i32(32)
	m.g0 = v3
	{
		{
			if uint32(v2) < uint32(i32(128)) {
				goto l0
			}
			store32(m.memory[int64(uint32(v3))+28:], uint32(i32(0)))
			m.fn522(v3+i32(16), v2, v3+i32(28))
			t1 := int32(load32(m.memory[int64(uint32(v3))+16:]))
			t2 := int32(load32(m.memory[int64(uint32(v3))+20:]))
			t3 := m.fn789(t1, t2, v0, v1)
			v2 = t3
			goto l1
		}
	l0:
		m.fn520(v3+i32(8), v2, v0, v1)
		t4 := int32(load32(m.memory[int64(uint32(v3))+8:]))
		var p5 int32
		if t4 == i32(1) {
			p5 = 1
		}
		v2 = p5
	}
l1:
	m.g0 = v3 + i32(32)
	return v2
}
func (m *Module) fn780(v0, v1 int32) {
	{
		t0 := int32(load32(m.memory[uint32(v1):]))
		if t0 != 0 {
			goto l0
		}
		t1 := int64(load64(m.memory[int64(uint32(v1))+20:]))
		store64(m.memory[int64(uint32(v0))+16:], uint64(t1))
		t2 := int64(load64(m.memory[int64(uint32(v1))+12:]))
		store64(m.memory[int64(uint32(v0))+8:], uint64(t2))
		t3 := int64(load64(m.memory[int64(uint32(v1))+4:]))
		store64(m.memory[uint32(v0):], uint64(t3))
		return
	}
l0:
	store32(m.memory[uint32(v0):], uint32(i32(-1)))
	m.fn781(v1)
}
func (m *Module) fn781(v0 int32) {
	var v1 int32
	v1 = v0 + i32(4)
	{
		t0 := int32(load32(m.memory[uint32(v0):]))
		if t0 != 0 {
			goto l0
		}
		m.fn784(v1)
		return
	}
l0:
	m.fn785(v1)
}
func (m *Module) fn782(v0, v1 int32) int32 {
	var v2 int32
	var v3 int64
	var v4, v5, v6, v7, v8, v9 int32
	var v10 int64
	t0 := m.g0
	v2 = t0 - i32(32)
	m.g0 = v2
	t1 := int64(load64(m.memory[int64(uint32(v0))+16:]))
	t2 := int64(load64(m.memory[int64(uint32(v0))+24:]))
	t3 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	t4 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	t5 := m.fn540(t1, t2, t3, t4)
	v3 = t5
	store32(m.memory[int64(uint32(v2))+20:], uint32(v1))
	m.fn684(v0, i32(1), v0+i32(16))
	store32(m.memory[int64(uint32(v2))+28:], uint32(v0))
	store32(m.memory[int64(uint32(v2))+24:], uint32(v2+i32(20)))
	t6 := int32(load32(m.memory[uint32(v0):]))
	t7 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	m.fn69(v2, t6, t7, v3, v2+i32(24), i32(140))
	{
		{
			t8 := int32(load32(m.memory[uint32(v2):]))
			v4 = t8
			if v4 != i32(1) {
				goto l0
			}
			t9 := int32(load32(m.memory[uint32(v0):]))
			v5 = t9
			t10 := int32(load32(m.memory[int64(uint32(v2))+4:]))
			t11 := v5
			v6 = t10
			v7 = t11 + v6
			t12 := int32(m.memory[uint32(v7)])
			v8 = t12
			t13 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			v9 = t13
			t14 := int64(load64(m.memory[uint32(v1):]))
			v10 = t14
			t15 := v7
			v1 = int32(uint32(int32(v3)) >> 25)
			m.memory[uint32(t15)] = byte(v1)
			t16 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			m.memory[uint32(v5+t16&(v6+i32(-8))+i32(8))] = byte(v1)
			t17 := int32(load32(m.memory[int64(uint32(v0))+12:]))
			store32(m.memory[int64(uint32(v0))+12:], uint32(t17+i32(1)))
			t18 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			store32(m.memory[int64(uint32(v0))+8:], uint32(t18-v8&i32(1)))
			v0 = v5 + (i32(0)-v6)*i32(12) + i32(-12)
			store64(m.memory[uint32(v0):], uint64(v10))
			store32(m.memory[int64(uint32(v0))+8:], uint32(v9))
			goto l1
		}
	l0:
		t19 := int32(load32(m.memory[uint32(v1):]))
		t20 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		m.fn16(t19, t20)
	}
l1:
	m.g0 = v2 + i32(32)
	return (v4 ^ i32(-1)) & i32(1)
}
func (m *Module) fn783(v0, v1 int32) int32 {
	var v2 int32
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(t0):]))
	v2 = t1
	t2 := int32(load32(m.memory[uint32(v2+i32(4)):]))
	t3 := int32(load32(m.memory[uint32(v2+i32(8)):]))
	t4 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t5 := int32(load32(m.memory[uint32(t4):]))
	v0 = t5 + (i32(0)-v1)*i32(12)
	t6 := int32(load32(m.memory[uint32(v0+i32(-8)):]))
	t7 := int32(load32(m.memory[uint32(v0+i32(-4)):]))
	t8 := m.fn544(t2, t3, t6, t7)
	return t8
}
func (m *Module) fn784(v0 int32) {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	m.fn16(t0, t1)
	t2 := int32(load32(m.memory[int64(uint32(v0))+12:]))
	t3 := int32(load32(m.memory[int64(uint32(v0))+16:]))
	m.fn134(t2, t3)
}
func (m *Module) fn785(v0 int32) {
	var v1 int32
	{
		t0 := int32(load32(m.memory[uint32(v0):]))
		v1 = t0
		p1 := i32(1)
		if v1 < i32(0) {
			p1 = v1 ^ i32(-0x80000000)
		}
		switch p1 {
		default:
			t2 := int32(m.memory[int64(uint32(v0))+4])
			t3 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			m.fn119(t2, t3)
			return
		case 0:
			t4 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t5 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			m.fn16(t4, t5)
			return
		case 1:
			t6 := int32(load32(m.memory[int64(uint32(v0))+12:]))
			t7 := int32(load32(m.memory[int64(uint32(v0))+16:]))
			m.fn134(t6, t7)
			t8 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			m.fn16(v1, t8)
			fallthrough
		case 2:
			return
		case 3:
			t9 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t10 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			m.fn16(t9, t10)
			return
		case 4:
			t11 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t12 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			m.fn16(t11, t12)
		}
	}
}
func (m *Module) fn786(v0, v1, v2, v3 int32) {
	if v1 == 0 {
		goto l0
	}
	{
		if uint32(v3) > uint32(v1) {
			goto l1
		}
		if v3 != v1 {
			goto l2
		}
		goto l0
	l1:
		t0 := int32(int8(m.memory[uint32(v2+v1)]))
		if t0 > i32(-65) {
			goto l0
		}
	}
l2:
	v2 = i32(0)
	goto l3
l0:
	v2 = v2 + v1
	v1 = v3 - v1
l3:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
	store32(m.memory[uint32(v0):], uint32(v2))
}
func (m *Module) fn787(v0, v1, v2, v3, v4 int32) {
	var v5, v6 int32
	v5 = i32(0)
	if uint32(v2) < uint32(v1) {
		goto l5
	}
	{
		if v1 == 0 {
			goto l1
		}
		if uint32(v1) < uint32(v4) {
			goto l2
		}
		if v1 != v4 {
			goto l5
		}
		goto l1
	l2:
		t0 := int32(int8(m.memory[uint32(v3+v1)]))
		if t0 <= i32(-65) {
			goto l5
		}
	}
l1:
	{
		if v2 == 0 {
			goto l3
		}
		if uint32(v2) < uint32(v4) {
			goto l4
		}
		if v2 != v4 {
			goto l5
		}
		goto l3
	l4:
		t1 := int32(int8(m.memory[uint32(v3+v2)]))
		if t1 <= i32(-65) {
			goto l5
		}
	}
l3:
	v5 = v3 + v1
	v6 = v2 - v1
	goto l5
l5:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v6))
	store32(m.memory[uint32(v0):], uint32(v5))
}
func (m *Module) fn788(v0, v1, v2, v3 int32) {
	var v4, v5, v6, v7 int32
	t0 := m.g0
	v4 = t0 - i32(16)
	m.g0 = v4
	v5 = i32(1)
	v6 = i32(0)
	{
		switch v2 {
		case 0:
			goto l0
		case 1:
			v6 = i32(1)
			v5 = i32(1)
			t1 := int32(m.memory[uint32(v1)])
			v7 = t1
			switch v7 + i32(-43) {
			case 0, 2:
				goto l0
			default:
				goto l3
			}
		default:
			t2 := int32(m.memory[uint32(v1)])
			v7 = t2
		}
	l3:
		t3 := v1
		var p4 int32
		if v7&i32(255) == i32(43) {
			p4 = 1
		}
		v6 = p4
		v1 = t3 + v6
		v2 = v2 - v6
		if uint32(v2) < uint32(i32(3)) {
			v5 = i32(0)
			v7 = i32(0)
		l11:
			if v2 != 0 {
				t9 := int32(m.memory[uint32(v1)])
				m.fn199(v4, t9, v3)
				v6 = i32(1)
				{
					t10 := int32(load32(m.memory[uint32(v4):]))
					if t10 == i32(1) {
						v1 = v1 + i32(1)
						v2 = v2 + i32(-1)
						t11 := int32(load32(m.memory[int64(uint32(v4))+4:]))
						v7 = v7*v3 + t11
						goto l11
					}
					v5 = i32(1)
					goto l0
				}
			}
			v6 = v7
			goto l0
		}
		v6 = i32(0)
	l8:
		{
			if v2 != 0 {
				goto l5
			}
			v5 = i32(0)
			goto l0
		l5:
			t5 := int32(m.memory[uint32(v1)])
			m.fn199(v4+i32(8), t5, v3)
			t6 := int32(load32(m.memory[int64(uint32(v4))+8:]))
			v7 = t6
			v6 = v6 & i32(255) * (v3 & i32(255))
			if int32(uint32(v6)>>8) != 0 {
				v5 = i32(1)
				p7 := i32(1)
				if v7&i32(1) != 0 {
					p7 = i32(2)
				}
				v6 = p7
				goto l0
			}
			v5 = i32(1)
			if v7&i32(1) != 0 {
				goto l7
			}
			v6 = i32(1)
			goto l0
		l7:
			v1 = v1 + i32(1)
			v2 = v2 + i32(-1)
			t8 := int32(load32(m.memory[int64(uint32(v4))+12:]))
			v6 = v6&i32(255) + t8&i32(255)
			if v6&i32(255) == v6 {
				goto l8
			}
		}
		v6 = i32(2)
		goto l0
	}
l0:
	m.memory[int64(uint32(v0))+1] = byte(v6)
	m.memory[uint32(v0)] = byte(v5)
	m.g0 = v4 + i32(16)
}
func (m *Module) fn789(v0, v1, v2, v3 int32) int32 {
	var v4, v5 int32
	t0 := m.g0
	v4 = t0 - i32(96)
	m.g0 = v4
	{
		{
			if uint32(v1) >= uint32(v3) {
				v5 = i32(0)
				if v1 == v3 {
					goto l3
				}
				goto l2
			}
			if v1 == i32(1) {
				t2 := int32(m.memory[uint32(v0)])
				m.fn520(v4+i32(8), t2, v2, v3)
				t3 := int32(load32(m.memory[int64(uint32(v4))+8:]))
				var p4 int32
				if t3 == i32(1) {
					p4 = 1
				}
				v5 = p4
				goto l2
			}
			m.fn601(v4+i32(32), v2, v3, v0, v1)
			m.fn790(v4+i32(20), v4+i32(32))
			t1 := int32(load32(m.memory[int64(uint32(v4))+20:]))
			v5 = t1
			goto l2
		}
	l3:
		t5 := m.fn1851(v0, v2, v1)
		var p6 int32
		if t5 == 0 {
			p6 = 1
		}
		v5 = p6
	}
l2:
	m.g0 = v4 + i32(96)
	return v5
}
func (m *Module) fn790(v0, v1 int32) {
	var v2, v3, v4, v5, v6, v7, v8, v9, v10 int32
	t0 := m.g0
	v2 = t0 - i32(32)
	m.g0 = v2
	{
		t1 := int32(load32(m.memory[uint32(v1):]))
		if t1 != 0 {
			v6 = v1 + i32(8)
			t13 := int32(load32(m.memory[int64(uint32(v1))+60:]))
			v8 = t13
			t14 := int32(load32(m.memory[int64(uint32(v1))+56:]))
			v10 = t14
			t15 := int32(load32(m.memory[int64(uint32(v1))+52:]))
			v9 = t15
			t16 := int32(load32(m.memory[int64(uint32(v1))+48:]))
			v3 = t16
			t17 := int32(load32(m.memory[int64(uint32(v1))+36:]))
			if t17 == i32(-1) {
				goto l9
			}
			m.fn602(v0, v6, v3, v9, v10, v8, i32(0))
			goto l8
		}
		t2 := int32(m.memory[int64(uint32(v1))+12])
		v3 = t2
		t3 := int32(load32(m.memory[int64(uint32(v1))+52:]))
		v4 = t3
		t4 := int32(load32(m.memory[int64(uint32(v1))+48:]))
		v5 = t4
		t5 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		v6 = t5
		t6 := int32(m.memory[int64(uint32(v1))+14])
		v7 = t6 & i32(1)
	l6:
		{
			v8 = i32(0)
			if v7 != 0 {
				goto l1
			}
			m.memory[int64(uint32(v1))+12] = byte((v3 ^ i32(-1)) & i32(1))
			m.fn786(v2+i32(16), v6, v5, v4)
			t7 := int32(load32(m.memory[int64(uint32(v2))+16:]))
			v9 = t7
			if v9 == 0 {
				m.fn556(v5, v4, v6, v4, i32(1102044))
				panic("unreachable")
			}
			t8 := int32(load32(m.memory[int64(uint32(v2))+20:]))
			v10 = t8
			store32(m.memory[int64(uint32(v2))+24:], uint32(v9))
			store32(m.memory[int64(uint32(v2))+28:], uint32(v9+v10))
			m.fn374(v2+i32(8), v2+i32(24))
			{
				t9 := int32(load32(m.memory[int64(uint32(v2))+8:]))
				if t9&i32(1) == 0 {
					goto l3
				}
				if v3&i32(1) != 0 {
					goto l4
				}
				v3 = i32(1)
				v9 = i32(1)
				{
					t10 := int32(load32(m.memory[int64(uint32(v2))+12:]))
					v10 = t10
					if uint32(v10) < uint32(i32(128)) {
						goto l5
					}
					v9 = i32(2)
					if uint32(v10) < uint32(i32(2048)) {
						goto l5
					}
					p11 := i32(4)
					if uint32(v10) < uint32(i32(65536)) {
						p11 = i32(3)
					}
					v9 = p11
				}
			l5:
				t12 := v1
				v6 = v9 + v6
				store32(m.memory[int64(uint32(t12))+4:], uint32(v6))
				goto l6
			}
		l3:
		}
		if v3&i32(1) == 0 {
			goto l7
		}
	l4:
		store32(m.memory[int64(uint32(v0))+8:], uint32(v6))
		store32(m.memory[int64(uint32(v0))+4:], uint32(v6))
		v8 = i32(1)
		goto l1
	l7:
		m.memory[int64(uint32(v1))+14] = byte(i32(1))
	l1:
		store32(m.memory[uint32(v0):], uint32(v8))
		goto l8
	}
l9:
	m.fn602(v0, v6, v3, v9, v10, v8, i32(1))
l8:
	m.g0 = v2 + i32(32)
}
func (m *Module) fn791(v0, v1, v2, v3 int32) int32 {
	var v4, v5, v6, v7, v8, v9, v10 int32
	v4 = i32(0)
	if v0 == 0 {
		goto l0
	}
l5:
	{
		v5 = v0 + i32(4)
		t0 := int32(load16(m.memory[int64(uint32(v0))+886:]))
		v6 = t0
		v7 = v6 * i32(12)
		v8 = i32(-1)
	l4:
		{
			if v7 != 0 {
				goto l1
			}
			v8 = v6
			goto l2
		l1:
			t1 := int32(load32(m.memory[int64(uint32(v5))+8:]))
			v9 = t1
			t2 := int32(load32(m.memory[int64(uint32(v5))+4:]))
			v10 = t2
			v7 = v7 + i32(-12)
			v8 = v8 + i32(1)
			v5 = v5 + i32(12)
			{
				t3 := m.fn643(v2, v3, v10, v9)
				switch t3 & i32(255) {
				case 1:
					goto l4
				default:
					goto l2
				case 0:
				}
			}
		}
		v4 = v0 + v8*i32(68) + i32(136)
		goto l0
	l2:
		if v1 == 0 {
			goto l0
		}
		v1 = v1 + i32(-1)
		t4 := int32(load32(m.memory[int64(uint32(v0+v8<<2))+888:]))
		v0 = t4
		goto l5
	}
l0:
	return v4
}
func (m *Module) fn792(v0, v1, v2 int32) {
	var v3, v4, v5, v6, v7, v8, v9, v10 int32
	t0 := m.g0
	v3 = t0 - i32(48)
	m.g0 = v3
	t1 := int32(uint32(v2-v1) / uint32(i32(20)))
	t2 := v3
	v4 = t1
	m.fn59(t2, v4, i32(4), i32(12))
	store32(m.memory[int64(uint32(v3))+20:], uint32(i32(0)))
	t3 := int64(load64(m.memory[uint32(v3):]))
	store64(m.memory[int64(uint32(v3))+12:], uint64(t3))
	m.fn60(v3+i32(12), v4)
	t4 := int32(load32(m.memory[int64(uint32(v3))+20:]))
	v5 = t4
	{
		if v2 == v1 {
			goto l0
		}
		t5 := int32(load32(m.memory[int64(uint32(v3))+16:]))
		v6 = t5
		v7 = i32(0)
	l7:
		{
			{
				v2 = v1 + v7*i32(20)
				t6 := int32(load32(m.memory[uint32(v2):]))
				if t6 != i32(-1) {
					goto l1
				}
				v8 = i32(0)
				v9 = i32(1)
				v10 = i32(0)
				goto l2
			}
		l1:
			store32(m.memory[int64(uint32(v3))+32:], uint32(i32(0)))
			store64(m.memory[int64(uint32(v3))+24:], uint64(i64(0x100000000)))
			t7 := int32(load32(m.memory[int64(uint32(v2))+8:]))
			v8 = t7 << 5
			t8 := int32(load32(m.memory[int64(uint32(v2))+4:]))
			v2 = t8
		l6:
			{
				if v8 == 0 {
					t17 := int32(load32(m.memory[int64(uint32(v3))+32:]))
					v8 = t17
					t18 := int32(load32(m.memory[int64(uint32(v3))+28:]))
					v9 = t18
					t19 := int32(load32(m.memory[int64(uint32(v3))+24:]))
					v10 = t19
					goto l2
				}
				t9 := int32(load32(m.memory[uint32(v2):]))
				if t9 != i32(-0x80000000) {
					goto l4
				}
				{
					t10 := int32(load32(m.memory[int64(uint32(v3))+32:]))
					if t10 == 0 {
						goto l5
					}
					m.fn74(v3+i32(24), i32(10))
				}
			l5:
				t11 := int32(load32(m.memory[int64(uint32(v2))+8:]))
				t12 := int32(load32(m.memory[int64(uint32(v2))+12:]))
				m.fn45(v3+i32(36), t11, t12)
				t13 := int32(load32(m.memory[int64(uint32(v3))+40:]))
				t14 := v3 + i32(24)
				v9 = t13
				t15 := int32(load32(m.memory[int64(uint32(v3))+44:]))
				m.fn75(t14, v9, t15)
				t16 := int32(load32(m.memory[int64(uint32(v3))+36:]))
				m.fn16(t16, v9)
				v8 = v8 + i32(-32)
				v2 = v2 + i32(32)
				goto l6
			}
		l4:
			t20 := int32(load32(m.memory[int64(uint32(v3))+24:]))
			t21 := int32(load32(m.memory[int64(uint32(v3))+28:]))
			m.fn16(t20, t21)
			v9 = i32(1)
			v8 = i32(0)
			v10 = i32(0)
		}
	l2:
		v2 = v6 + v5*i32(12)
		store32(m.memory[int64(uint32(v2))+8:], uint32(v8))
		store32(m.memory[int64(uint32(v2))+4:], uint32(v9))
		store32(m.memory[uint32(v2):], uint32(v10))
		v5 = v5 + i32(1)
		v7 = v7 + i32(1)
		if v7 != v4 {
			goto l7
		}
	}
l0:
	t22 := int64(load64(m.memory[int64(uint32(v3))+12:]))
	store64(m.memory[uint32(v0):], uint64(t22))
	store32(m.memory[int64(uint32(v0))+8:], uint32(v5))
	m.g0 = v3 + i32(48)
}
func (m *Module) fn793(v0, v1, v2 int32) {
	var v3 int32
	v1 = v1 * i32(28)
l6:
	{
		if v1 == 0 {
			return
		}
		t0 := int32(load32(m.memory[uint32(v0):]))
		v3 = t0
		p1 := i32(1)
		if uint32(v3) > uint32(i32(2)) {
			p1 = v3 + i32(-3)
		}
		switch p1 {
		case 1:
			t4 := int32(load32(m.memory[uint32(v0+i32(20)):]))
			t5 := int32(load32(m.memory[uint32(v0+i32(24)):]))
			m.fn793(t4, t5, v2)
			goto l4
		case 2:
			t6 := int32(load32(m.memory[uint32(v0+i32(8)):]))
			t7 := int32(load32(m.memory[uint32(v0+i32(12)):]))
			m.fn75(v2, t6, t7)
			goto l4
		case 3, 4:
			goto l4
		case 5:
			goto l5
		default:
			t2 := int32(load32(m.memory[uint32(v0+i32(8)):]))
			t3 := int32(load32(m.memory[uint32(v0+i32(12)):]))
			m.fn75(v2, t2, t3)
			goto l4
		}
	}
l5:
	m.fn74(v2, i32(10))
l4:
	v0 = v0 + i32(28)
	v1 = v1 + i32(-28)
	goto l6
}
func (m *Module) fn794(v0, v1, v2 int32) {
	var v3, v4, v5, v6, v7 int32
	t0 := m.g0
	v3 = t0 - i32(144)
	m.g0 = v3
	store32(m.memory[int64(uint32(v3))+36:], uint32(i32(0)))
	store64(m.memory[int64(uint32(v3))+28:], uint64(i64(0x400000000)))
	t1 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	v4 = t1 << 5
	t2 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	v1 = t2
l1:
	if v4 == 0 {
		goto l0
	}
	m.fn795(v1, v2, v3+i32(28))
	v4 = v4 + i32(-32)
	v1 = v1 + i32(32)
	goto l1
l0:
	t3 := int32(load32(m.memory[int64(uint32(v3))+32:]))
	t4 := int32(load32(m.memory[int64(uint32(v3))+36:]))
	m.fn77(v3+i32(40), t3, t4, i32(1084652), i32(4))
	t5 := int32(load32(m.memory[int64(uint32(v3))+44:]))
	t6 := v3 + i32(52)
	v2 = t5
	t7 := int32(load32(m.memory[int64(uint32(v3))+48:]))
	m.fn70(t6, v2, t7)
	m.fn796(v3+i32(16), v3+i32(52))
	{
		{
			t8 := int32(load32(m.memory[int64(uint32(v3))+16:]))
			v4 = t8
			if v4 != 0 {
				goto l2
			}
			v4 = i32(0)
			v1 = i32(4)
			v5 = i32(0)
			goto l3
		}
	l2:
		t9 := int32(load32(m.memory[int64(uint32(v3))+20:]))
		v1 = t9
		m.fn59(v3+i32(8), i32(4), i32(4), i32(8))
		t10 := int32(load32(m.memory[int64(uint32(v3))+8:]))
		v2 = t10
		t11 := int32(load32(m.memory[int64(uint32(v3))+12:]))
		v6 = t11
		store32(m.memory[int64(uint32(v6))+4:], uint32(v1))
		store32(m.memory[uint32(v6):], uint32(v4))
		store32(m.memory[int64(uint32(v3))+100:], uint32(i32(1)))
		store32(m.memory[int64(uint32(v3))+96:], uint32(v6))
		store32(m.memory[int64(uint32(v3))+92:], uint32(v2))
		memory_copy(m.memory, uint32(v3+i32(104)), uint32(v3+i32(52)), uint32(i32(40)))
		v1 = i32(12)
		v4 = i32(1)
	l6:
		{
			m.fn796(v3, v3+i32(104))
			t12 := int32(load32(m.memory[uint32(v3):]))
			v2 = t12
			if v2 == 0 {
				goto l4
			}
			t13 := int32(load32(m.memory[int64(uint32(v3))+4:]))
			v5 = t13
			{
				t14 := int32(load32(m.memory[int64(uint32(v3))+92:]))
				if v4 != t14 {
					goto l5
				}
				m.fn797(v3 + i32(92))
				t15 := int32(load32(m.memory[int64(uint32(v3))+96:]))
				v6 = t15
			}
		l5:
			v7 = v6 + v1
			store32(m.memory[uint32(v7):], uint32(v5))
			store32(m.memory[uint32(v7+i32(-4)):], uint32(v2))
			t16 := v3
			v4 = v4 + i32(1)
			store32(m.memory[int64(uint32(t16))+100:], uint32(v4))
			v1 = v1 + i32(8)
			goto l6
		}
	l4:
		t17 := int32(load32(m.memory[int64(uint32(v3))+44:]))
		v2 = t17
		t18 := int32(load32(m.memory[int64(uint32(v3))+96:]))
		v1 = t18
		t19 := int32(load32(m.memory[int64(uint32(v3))+92:]))
		v5 = t19
	}
l3:
	m.fn632(v0, v1, v4, i32(1084652), i32(4))
	m.fn639(v5, v1)
	t20 := int32(load32(m.memory[int64(uint32(v3))+40:]))
	m.fn16(t20, v2)
	m.fn78(v3 + i32(28))
	m.g0 = v3 + i32(144)
}
func (m *Module) fn795(v0, v1, v2 int32) {
	var v3, v4, v5, v6 int32
	var v7, v8 int64
	var v9, v10 int32
	var v11 int64
	var v12 int32
	t0 := m.g0
	v3 = t0 - i32(112)
	m.g0 = v3
	{
		{
			t1 := int32(load32(m.memory[uint32(v0):]))
			v4 = t1
			switch v4 >> 31 & (v4 + i32(-0x7fffffff)) {
			case 6:
				goto l6
			case 5:
				t2 := int32(load32(m.memory[int64(uint32(v0))+8:]))
				t3 := int32(load32(m.memory[int64(uint32(v0))+12:]))
				m.fn46(v3+i32(32), t2, t3)
				t4 := int32(load32(m.memory[int64(uint32(v3))+36:]))
				v4 = t4
				if v4 != 0 {
					goto l7
				}
				goto l6
			default:
				t5 := int32(load32(m.memory[int64(uint32(v0))+4:]))
				t6 := int32(load32(m.memory[int64(uint32(v0))+8:]))
				m.fn798(v3+i32(100), t5, t6, i32(2), v1)
				t7 := int32(load32(m.memory[int64(uint32(v3))+104:]))
				t8 := v3 + i32(8)
				v4 = t7
				t9 := int32(load32(m.memory[int64(uint32(v3))+108:]))
				t10 := v4
				v0 = t9
				m.fn46(t8, t10, v0)
				{
					t11 := int32(load32(m.memory[int64(uint32(v3))+12:]))
					if t11 == 0 {
						goto l8
					}
					m.fn46(v3, v4, v0)
					t12 := int64(load64(m.memory[uint32(v3):]))
					store64(m.memory[int64(uint32(v3))+56:], uint64(t12))
					store32(m.memory[int64(uint32(v3))+92:], uint32(i32(1)))
					store32(m.memory[int64(uint32(v3))+88:], uint32(v3+i32(56)))
					m.fn73(v3+i32(72), i32(1068669), v3+i32(88))
					m.fn33(v2, v3+i32(72))
				}
			l8:
				t13 := int32(load32(m.memory[int64(uint32(v3))+100:]))
				m.fn16(t13, v4)
				goto l6
			case 1:
				t14 := int32(load32(m.memory[int64(uint32(v0))+8:]))
				t15 := int32(load32(m.memory[int64(uint32(v0))+12:]))
				m.fn798(v3+i32(72), t14, t15, i32(2), v1)
				t16 := int32(load32(m.memory[int64(uint32(v3))+76:]))
				t17 := v3 + i32(16)
				v4 = t16
				t18 := int32(load32(m.memory[int64(uint32(v3))+80:]))
				m.fn46(t17, v4, t18)
				{
					t19 := int32(load32(m.memory[int64(uint32(v3))+20:]))
					if t19 == 0 {
						t20 := int32(load32(m.memory[int64(uint32(v3))+72:]))
						m.fn16(t20, v4)
						goto l6
					}
					m.fn33(v2, v3+i32(72))
					goto l6
				}
			case 2:
				t21 := int32(load32(m.memory[int64(uint32(v0))+20:]))
				v5 = t21
				t22 := int32(load32(m.memory[int64(uint32(v0))+24:]))
				v6 = v5 + t22*i32(28)
				t23 := int64(load64(m.memory[int64(uint32(v0))+8:]))
				v7 = t23
				v8 = i64(0)
				t24 := int32(m.memory[int64(uint32(v0))+28])
				v9 = t24
				v10 = v9 & i32(255)
			l16:
				{
					if v5 == v6 {
						goto l6
					}
					store32(m.memory[int64(uint32(v3))+52:], uint32(i32(0)))
					store64(m.memory[int64(uint32(v3))+44:], uint64(i64(0x400000000)))
					t25 := int32(load32(m.memory[int64(uint32(v5))+8:]))
					v4 = t25 << 5
					t26 := int32(load32(m.memory[int64(uint32(v5))+4:]))
					v0 = t26
				l11:
					if v4 == 0 {
						{
							{
								t27 := int32(load32(m.memory[int64(uint32(v5))+12:]))
								if t27 == i32(-1) {
									goto l12
								}
								t28 := int32(load32(m.memory[int64(uint32(v5))+16:]))
								t29 := int32(load32(m.memory[int64(uint32(v5))+20:]))
								m.fn799(v3+i32(72), t28, t29, i32(2))
								store32(m.memory[int64(uint32(v3))+92:], uint32(i32(25)))
								store32(m.memory[int64(uint32(v3))+88:], uint32(v3+i32(72)))
								m.fn73(v3+i32(100), i32(1070105), v3+i32(88))
								t30 := int32(load32(m.memory[int64(uint32(v3))+72:]))
								t31 := int32(load32(m.memory[int64(uint32(v3))+76:]))
								m.fn16(t30, t31)
								t32 := int64(load64(m.memory[int64(uint32(v3))+100:]))
								store64(m.memory[int64(uint32(v3))+56:], uint64(t32))
								t33 := int32(load32(m.memory[int64(uint32(v3))+108:]))
								store32(m.memory[int64(uint32(v3))+64:], uint32(t33))
								goto l13
							}
						l12:
							if v10 != 0 {
								goto l14
							}
							m.fn51(v3+i32(56), i32(1084672), i32(4))
							goto l13
						l14:
							t34 := v3 + i32(72)
							t35 := v9
							v11 = v7 + v8
							p36 := v11
							if uint64(v11) < uint64(v7) {
								p36 = i64(-1)
							}
							m.fn800(t34, t35, p36)
							store32(m.memory[int64(uint32(v3))+92:], uint32(i32(25)))
							store32(m.memory[int64(uint32(v3))+88:], uint32(v3+i32(72)))
							m.fn73(v3+i32(100), i32(1070105), v3+i32(88))
							t37 := int32(load32(m.memory[int64(uint32(v3))+72:]))
							t38 := int32(load32(m.memory[int64(uint32(v3))+76:]))
							m.fn16(t37, t38)
							t39 := int64(load64(m.memory[int64(uint32(v3))+100:]))
							store64(m.memory[int64(uint32(v3))+56:], uint64(t39))
							t40 := int32(load32(m.memory[int64(uint32(v3))+108:]))
							store32(m.memory[int64(uint32(v3))+64:], uint32(t40))
						}
					l13:
						{
							t41 := int32(load32(m.memory[int64(uint32(v3))+52:]))
							v4 = t41
							if v4 == 0 {
								goto l15
							}
							t42 := int32(load32(m.memory[int64(uint32(v3))+48:]))
							m.fn77(v3+i32(100), t42, v4, i32(1097368), i32(1))
							store32(m.memory[int64(uint32(v3))+84:], uint32(i32(25)))
							store32(m.memory[int64(uint32(v3))+76:], uint32(i32(25)))
							store32(m.memory[int64(uint32(v3))+80:], uint32(v3+i32(100)))
							store32(m.memory[int64(uint32(v3))+72:], uint32(v3+i32(56)))
							m.fn73(v3+i32(88), i32(0x10004e), v3+i32(72))
							t43 := int32(load32(m.memory[int64(uint32(v3))+100:]))
							t44 := int32(load32(m.memory[int64(uint32(v3))+104:]))
							m.fn16(t43, t44)
							m.fn33(v2, v3+i32(88))
						}
					l15:
						v8 = v8 + i64(1)
						v5 = v5 + i32(28)
						t45 := int32(load32(m.memory[int64(uint32(v3))+56:]))
						t46 := int32(load32(m.memory[int64(uint32(v3))+60:]))
						m.fn16(t45, t46)
						m.fn78(v3 + i32(44))
						goto l16
					}
					m.fn795(v0, v1, v3+i32(44))
					v4 = v4 + i32(-32)
					v0 = v0 + i32(32)
					goto l11
				}
			case 3:
				t47 := int32(load32(m.memory[int64(uint32(v0))+8:]))
				v10 = t47
				t48 := int32(load32(m.memory[int64(uint32(v0))+12:]))
				v12 = v10 + t48*i32(12)
			l23:
				{
					if v10 == v12 {
						goto l6
					}
					t49 := int32(load32(m.memory[int64(uint32(v10))+4:]))
					v0 = t49
					t50 := int32(load32(m.memory[int64(uint32(v10))+8:]))
					t51 := v3 + i32(24)
					v5 = t50
					m.fn59(t51, v5, i32(4), i32(12))
					store32(m.memory[int64(uint32(v3))+108:], uint32(i32(0)))
					t52 := int64(load64(m.memory[int64(uint32(v3))+24:]))
					store64(m.memory[int64(uint32(v3))+100:], uint64(t52))
					m.fn60(v3+i32(100), v5)
					t53 := int32(load32(m.memory[int64(uint32(v3))+108:]))
					v6 = t53
					{
						if v5 == 0 {
							goto l17
						}
						v9 = v5 + v6
						t54 := int32(load32(m.memory[int64(uint32(v3))+104:]))
						v4 = t54 + v6*i32(12)
					l20:
						{
							{
								t55 := int32(load32(m.memory[uint32(v0):]))
								if t55 != i32(-1) {
									goto l18
								}
								store32(m.memory[int64(uint32(v3))+80:], uint32(i32(0)))
								store64(m.memory[int64(uint32(v3))+72:], uint64(i64(0x100000000)))
								goto l19
							}
						l18:
							m.fn794(v3+i32(72), v0, v1)
						l19:
							t56 := int32(load32(m.memory[int64(uint32(v3))+80:]))
							store32(m.memory[int64(uint32(v4))+8:], uint32(t56))
							t57 := int64(load64(m.memory[int64(uint32(v3))+72:]))
							store64(m.memory[uint32(v4):], uint64(t57))
							v0 = v0 + i32(20)
							v4 = v4 + i32(12)
							v5 = v5 + i32(-1)
							if v5 != 0 {
								goto l20
							}
						}
						v6 = v9
					}
				l17:
					v10 = v10 + i32(12)
					t58 := int64(load64(m.memory[int64(uint32(v3))+100:]))
					store64(m.memory[int64(uint32(v3))+88:], uint64(t58))
					store32(m.memory[int64(uint32(v3))+96:], uint32(v6))
					v4 = v6 * i32(12)
					t59 := int32(load32(m.memory[int64(uint32(v3))+92:]))
					v9 = t59
					v0 = v9
				l22:
					{
						if v4 == 0 {
							goto l21
						}
						v4 = v4 + i32(-12)
						t60 := int32(load32(m.memory[int64(uint32(v0))+8:]))
						v5 = t60
						v0 = v0 + i32(12)
						if v5 == 0 {
							goto l22
						}
					}
					m.fn77(v3+i32(72), v9, v6, i32(1084676), i32(3))
					m.fn33(v2, v3+i32(72))
				l21:
					m.fn78(v3 + i32(88))
					goto l23
				}
			case 4:
				t61 := int32(load32(m.memory[int64(uint32(v0))+12:]))
				v4 = t61 << 5
				t62 := int32(load32(m.memory[int64(uint32(v0))+8:]))
				v0 = t62
			l24:
				if v4 == 0 {
					goto l6
				}
				m.fn795(v0, v1, v2)
				v4 = v4 + i32(-32)
				v0 = v0 + i32(32)
				goto l24
			}
		}
	l7:
		t63 := int32(load32(m.memory[int64(uint32(v3))+32:]))
		v0 = t63
		store32(m.memory[int64(uint32(v3))+108:], uint32(i32(0)))
		store64(m.memory[int64(uint32(v3))+100:], uint64(i64(0x100000000)))
		m.fn801(v0, v4, v3+i32(100))
		t64 := int32(load32(m.memory[int64(uint32(v3))+108:]))
		store32(m.memory[int64(uint32(v3))+80:], uint32(t64))
		t65 := int64(load64(m.memory[int64(uint32(v3))+100:]))
		store64(m.memory[int64(uint32(v3))+72:], uint64(t65))
		m.fn33(v2, v3+i32(72))
	}
l6:
	m.g0 = v3 + i32(112)
}
func (m *Module) fn796(v0, v1 int32) {
	var v2, v3, v4 int32
	t0 := m.g0
	v2 = t0 - i32(32)
	m.g0 = v2
	{
	l2:
		{
			m.fn71(v2+i32(24), v1)
			{
				t1 := int32(load32(m.memory[int64(uint32(v2))+24:]))
				v3 = t1
				if v3 != 0 {
					goto l0
				}
				v3 = i32(0)
				goto l1
			}
		l0:
			t2 := int32(load32(m.memory[int64(uint32(v2))+28:]))
			t3 := v2 + i32(16)
			t4 := v3
			v4 = t2
			m.fn46(t3, t4, v4)
			t5 := int32(load32(m.memory[int64(uint32(v2))+20:]))
			if t5 == 0 {
				goto l2
			}
		}
		m.fn46(v2+i32(8), v3, v4)
		t6 := int32(load32(m.memory[int64(uint32(v2))+12:]))
		v1 = t6
		t7 := int32(load32(m.memory[int64(uint32(v2))+8:]))
		v3 = t7
	}
l1:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
	store32(m.memory[uint32(v0):], uint32(v3))
	m.g0 = v2 + i32(32)
}
func (m *Module) fn797(v0 int32) {
	var v1 int32
	{
		t0 := int32(load32(m.memory[uint32(v0):]))
		t1 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		v1 = t1
		if t0 != v1 {
			return
		}
		m.fn62(v0, v1, i32(1), i32(4), i32(8))
	}
}
func (m *Module) fn798(v0, v1, v2, v3, v4 int32) {
	m.fn802(v0, v1, v2, v3, i32(0), v4)
}
func (m *Module) fn799(v0, v1, v2, v3 int32) {
	var v4 int32
	t0 := m.g0
	v4 = t0 - i32(32)
	m.g0 = v4
	store64(m.memory[int64(uint32(v4))+12:], uint64(i64(0x100000000)))
	store32(m.memory[int64(uint32(v4))+20:], uint32(i32(0)))
	t1 := v4 + i32(12)
	t2 := int32(uint32(v2) >> 2)
	var p3 int32
	if v2&i32(3) != i32(0) {
		p3 = 1
	}
	m.fn47(t1, t2+p3)
	store32(m.memory[int64(uint32(v4))+24:], uint32(v1))
	store32(m.memory[int64(uint32(v4))+28:], uint32(v1+v2))
l1:
	{
		t4 := m.fn48(v4 + i32(24))
		v2 = t4
		if v2 == i32(-1) {
			goto l0
		}
		t6 := v4 + i32(12)
		p5 := v2
		if uint32(v2+i32(-127)) < uint32(i32(33)) {
			p5 = i32(32)
		}
		p7 := p5
		if uint32(v2) < uint32(i32(32)) {
			p7 = i32(32)
		}
		m.fn74(t6, p7)
		goto l1
	}
l0:
	t8 := int32(load32(m.memory[int64(uint32(v4))+12:]))
	v2 = t8
	t9 := int32(load32(m.memory[int64(uint32(v4))+16:]))
	t10 := v0
	v1 = t9
	t11 := int32(load32(m.memory[int64(uint32(v4))+20:]))
	t12 := v1
	t13 := v3
	var p14 int32
	if v3&i32(255) == 0 {
		p14 = 1
	}
	m.fn803(t10, t12, t11, t13, p14|i32(65536))
	m.fn16(v2, v1)
	m.g0 = v4 + i32(32)
}
func (m *Module) fn800(v0, v1 int32, v2 int64) {
	var v3 int32
	t0 := m.g0
	v3 = t0 - i32(32)
	m.g0 = v3
	{
		if v1&i32(255) != 0 {
			goto l0
		}
		m.fn51(v0, i32(1108000), i32(1))
		goto l1
	l0:
		m.fn804(v3+i32(12), v1, v2)
		store32(m.memory[int64(uint32(v3))+28:], uint32(i32(25)))
		store32(m.memory[int64(uint32(v3))+24:], uint32(v3+i32(12)))
		m.fn73(v0, i32(1068665), v3+i32(24))
		t1 := int32(load32(m.memory[int64(uint32(v3))+12:]))
		t2 := int32(load32(m.memory[int64(uint32(v3))+16:]))
		m.fn16(t1, t2)
	}
l1:
	m.g0 = v3 + i32(32)
}
func (m *Module) fn801(v0, v1, v2 int32) {
	var v3, v4, v5 int32
	t0 := m.g0
	v3 = t0 - i32(64)
	m.g0 = v3
	m.fn498(v3+i32(8), v0, v1, i32(10), i32(32))
	v1 = i32(1)
	t1 := int32(load32(m.memory[int64(uint32(v3))+12:]))
	t2 := v3 + i32(20)
	v0 = t1
	t3 := int32(load32(m.memory[int64(uint32(v3))+16:]))
	t4 := v0
	v4 = t3
	m.fn805(t2, t4, v4, i32(1))
	v5 = i32(1097368)
	{
		t5 := m.fn777(v0, v4, i32(96))
		if t5 != 0 {
			goto l0
		}
		t6 := m.fn806(v0, v4, i32(96))
		v1 = t6
		p7 := i32(1)
		if v1 != 0 {
			p7 = i32(1097368)
		}
		v5 = p7
	}
l0:
	store32(m.memory[int64(uint32(v3))+36:], uint32(v1))
	store32(m.memory[int64(uint32(v3))+32:], uint32(v5))
	store32(m.memory[int64(uint32(v3))+60:], uint32(i32(25)))
	store32(m.memory[int64(uint32(v3))+52:], uint32(i32(1)))
	store32(m.memory[int64(uint32(v3))+44:], uint32(i32(25)))
	store32(m.memory[int64(uint32(v3))+56:], uint32(v3+i32(8)))
	store32(m.memory[int64(uint32(v3))+48:], uint32(v3+i32(32)))
	store32(m.memory[int64(uint32(v3))+40:], uint32(v3+i32(20)))
	_ = m.fn404(v2, i32(1084857), v3+i32(40))
	t9 := int32(load32(m.memory[int64(uint32(v3))+20:]))
	t10 := int32(load32(m.memory[int64(uint32(v3))+24:]))
	m.fn16(t9, t10)
	t11 := int32(load32(m.memory[int64(uint32(v3))+8:]))
	t12 := int32(load32(m.memory[int64(uint32(v3))+12:]))
	m.fn16(t11, t12)
	m.g0 = v3 + i32(64)
}
func (m *Module) fn802(v0, v1, v2, v3, v4, v5 int32) {
	var v6, v7, v8, v9, v10, v11 int32
	var v12, v13 int64
	var v14, v15, v16, v17, v18, v19, v20, v21, v22, v23, v24, v25 int32
	var v26, v27, v28 int64
	var v29 int32
	t0 := m.g0
	v6 = t0 - i32(176)
	m.g0 = v6
	m.fn821(v6+i32(88), v1, v2, v5)
	store32(m.memory[int64(uint32(v6))+108:], uint32(i32(0)))
	store64(m.memory[int64(uint32(v6))+100:], uint64(i64(0x100000000)))
	p1 := i32(256)
	if v4 != 0 {
		p1 = i32(16777472)
	}
	v7 = p1
	p2 := i32(0)
	if v4 != 0 {
		p2 = i32(0x1000000)
	}
	v8 = p2
	v9 = v5 + i32(32)
	t3 := int32(load32(m.memory[int64(uint32(v6))+92:]))
	v10 = t3
	t4 := int32(load32(m.memory[int64(uint32(v6))+96:]))
	t5 := v10
	v2 = t4
	v11 = t5 + v2<<4
	t6 := int64(load64(m.memory[int64(uint32(v5))+24:]))
	v12 = t6
	t7 := int64(load64(m.memory[int64(uint32(v5))+16:]))
	v13 = t7
	t8 := int32(load32(m.memory[uint32(v5):]))
	v14 = t8
	t9 := int32(load32(m.memory[int64(uint32(v5))+4:]))
	v15 = t9
	t10 := int32(load32(m.memory[int64(uint32(v5))+12:]))
	v16 = t10
	v17 = v3 & i32(255)
	v18 = v10
	v1 = i32(0)
l29:
	{
		{
			{
				{
					{
						{
							{
								v4 = v18
								if v4 == v11 {
									t13 := int32(load32(m.memory[int64(uint32(v6))+108:]))
									store32(m.memory[int64(uint32(v0))+8:], uint32(t13))
									t14 := int64(load64(m.memory[int64(uint32(v6))+100:]))
									store64(m.memory[uint32(v0):], uint64(t14))
									v4 = v10
								l9:
									if v2 == 0 {
										t17 := int32(load32(m.memory[int64(uint32(v6))+88:]))
										m.fn419(t17, v10)
										m.g0 = v6 + i32(176)
										return
									}
									{
										t15 := int32(load32(m.memory[uint32(v4):]))
										v1 = t15
										if v1 < i32(-0x7ffffffb) {
											goto l8
										}
										t16 := int32(load32(m.memory[uint32(v4+i32(4)):]))
										m.fn134(v1, t16)
									}
								l8:
									v2 = v2 + i32(-1)
									v4 = v4 + i32(16)
									goto l9
								}
								v1 = v1 + i32(1)
								v18 = v4 + i32(16)
								t11 := int32(load32(m.memory[uint32(v4):]))
								v19 = t11
								p12 := i32(0)
								if v19 < i32(-0x7ffffffb) {
									p12 = v19 + i32(-0x7fffffff)
								}
								switch p12 {
								case 1:
									t40 := int32(load32(m.memory[int64(uint32(v4))+12:]))
									v19 = t40
									t41 := int32(load32(m.memory[int64(uint32(v4))+4:]))
									t42 := v6 + i32(112)
									v20 = t41
									t43 := int32(load32(m.memory[int64(uint32(v4))+8:]))
									t44 := v20
									v21 = t43
									m.fn802(t42, t44, v21, v3, i32(1), v5)
									t45 := int32(load32(m.memory[uint32(v19+i32(12)):]))
									v4 = t45
									t46 := int32(load32(m.memory[uint32(v19+i32(8)):]))
									v22 = t46
									{
										{
											t47 := int32(load32(m.memory[uint32(v19):]))
											v19 = t47
											if v19 != i32(2) {
												goto l19
											}
											t48 := m.fn827(v9, v22, v4)
											v4 = t48
											if v4 == 0 {
												m.fn802(v6+i32(148), v20, v21, v3, i32(0), v5)
												t59 := int32(load32(m.memory[int64(uint32(v6))+152:]))
												t60 := v6 + i32(100)
												v4 = t59
												t61 := int32(load32(m.memory[int64(uint32(v6))+156:]))
												m.fn75(t60, v4, t61)
												t62 := int32(load32(m.memory[int64(uint32(v6))+148:]))
												m.fn16(t62, v4)
												goto l24
											}
											t49 := int64(load64(m.memory[int64(uint32(v4))+4:]))
											store64(m.memory[int64(uint32(v6))+164:], uint64(t49))
											store32(m.memory[int64(uint32(v6))+152:], uint32(i32(1)))
											store32(m.memory[int64(uint32(v6))+148:], uint32(v6+i32(164)))
											m.fn73(v6+i32(124), i32(0x1000d9), v6+i32(148))
											goto l21
										}
									l19:
										m.fn31(v6+i32(124), v22, v4)
									l21:
										t50 := int32(load32(m.memory[int64(uint32(v6))+116:]))
										t51 := int32(load32(m.memory[int64(uint32(v6))+120:]))
										m.fn46(v6+i32(48), t50, t51)
										t52 := int32(load32(m.memory[int64(uint32(v6))+52:]))
										if t52 == 0 {
											if v19 == i32(2) {
												goto l25
											}
											t63 := int32(load32(m.memory[int64(uint32(v6))+128:]))
											v19 = t63
											t64 := int32(load32(m.memory[int64(uint32(v6))+132:]))
											v22 = t64
											store64(m.memory[int64(uint32(v6))+148:], uint64(i64(0x100000000)))
											store32(m.memory[int64(uint32(v6))+156:], uint32(i32(0)))
											t65 := v6 + i32(148)
											t66 := int32(uint32(v22) >> 2)
											var p67 int32
											if v22&i32(3) != i32(0) {
												p67 = 1
											}
											m.fn47(t65, t66+p67)
											store32(m.memory[int64(uint32(v6))+168:], uint32(v19+v22))
											store32(m.memory[int64(uint32(v6))+164:], uint32(v19))
										l27:
											{
												t68 := m.fn48(v6 + i32(164))
												v4 = t68
												if v4 == i32(-1) {
													t72 := int32(load32(m.memory[int64(uint32(v6))+148:]))
													v4 = t72
													t73 := int32(load32(m.memory[int64(uint32(v6))+152:]))
													t74 := v6 + i32(136)
													v20 = t73
													t75 := int32(load32(m.memory[int64(uint32(v6))+156:]))
													m.fn803(t74, v20, t75, v3, i32(0x1010000))
													m.fn16(v4, v20)
													m.fn828(v6+i32(164), v19, v22)
													store32(m.memory[int64(uint32(v6))+160:], uint32(i32(25)))
													store32(m.memory[int64(uint32(v6))+152:], uint32(i32(25)))
													store32(m.memory[int64(uint32(v6))+156:], uint32(v6+i32(164)))
													store32(m.memory[int64(uint32(v6))+148:], uint32(v6+i32(136)))
													_ = m.fn404(v6+i32(100), i32(1068892), v6+i32(148))
													t77 := int32(load32(m.memory[int64(uint32(v6))+164:]))
													t78 := int32(load32(m.memory[int64(uint32(v6))+168:]))
													m.fn16(t77, t78)
													t79 := int32(load32(m.memory[int64(uint32(v6))+136:]))
													t80 := int32(load32(m.memory[int64(uint32(v6))+140:]))
													m.fn16(t79, t80)
													goto l23
												}
												t70 := v6 + i32(148)
												p69 := v4
												if uint32(v4+i32(-127)) < uint32(i32(33)) {
													p69 = i32(32)
												}
												p71 := p69
												if uint32(v4) < uint32(i32(32)) {
													p71 = i32(32)
												}
												m.fn74(t70, p71)
												goto l27
											}
										}
										t53 := int32(load32(m.memory[int64(uint32(v6))+128:]))
										t54 := v6 + i32(164)
										v19 = t53
										t55 := int32(load32(m.memory[int64(uint32(v6))+132:]))
										m.fn828(t54, v19, t55)
										store32(m.memory[int64(uint32(v6))+160:], uint32(i32(25)))
										store32(m.memory[int64(uint32(v6))+152:], uint32(i32(25)))
										store32(m.memory[int64(uint32(v6))+156:], uint32(v6+i32(164)))
										store32(m.memory[int64(uint32(v6))+148:], uint32(v6+i32(112)))
										_ = m.fn404(v6+i32(100), i32(1068892), v6+i32(148))
										t57 := int32(load32(m.memory[int64(uint32(v6))+164:]))
										t58 := int32(load32(m.memory[int64(uint32(v6))+168:]))
										m.fn16(t57, t58)
										goto l23
									}
								case 2:
									t81 := int32(load32(m.memory[int64(uint32(v4))+8:]))
									v19 = t81
									t82 := int32(load32(m.memory[int64(uint32(v4))+4:]))
									v22 = t82
									{
										t83 := int32(load32(m.memory[int64(uint32(v4))+12:]))
										v4 = t83
										t84 := int32(load32(m.memory[uint32(v4):]))
										if t84 < i32(0) {
											m.fn46(v6+i32(72), v22, v19)
											t94 := int32(load32(m.memory[int64(uint32(v6))+76:]))
											if t94 == 0 {
												goto l29
											}
											m.fn46(v6+i32(64), v22, v19)
											t95 := int32(load32(m.memory[int64(uint32(v6))+64:]))
											t96 := int32(load32(m.memory[int64(uint32(v6))+68:]))
											m.fn803(v6+i32(148), t95, t96, v3, v8)
											t97 := int32(load32(m.memory[int64(uint32(v6))+152:]))
											t98 := v6 + i32(100)
											v4 = t97
											t99 := int32(load32(m.memory[int64(uint32(v6))+156:]))
											m.fn75(t98, v4, t99)
											t100 := int32(load32(m.memory[int64(uint32(v6))+148:]))
											m.fn16(t100, v4)
											goto l29
										}
										m.fn46(v6+i32(56), v22, v19)
										t85 := int32(load32(m.memory[int64(uint32(v6))+56:]))
										t86 := int32(load32(m.memory[int64(uint32(v6))+60:]))
										m.fn803(v6+i32(136), t85, t86, v3, i32(0x1000000))
										t87 := int32(load32(m.memory[int64(uint32(v4))+4:]))
										t88 := int32(load32(m.memory[int64(uint32(v4))+8:]))
										m.fn828(v6+i32(164), t87, t88)
										store32(m.memory[int64(uint32(v6))+160:], uint32(i32(25)))
										store32(m.memory[int64(uint32(v6))+152:], uint32(i32(25)))
										store32(m.memory[int64(uint32(v6))+156:], uint32(v6+i32(164)))
										store32(m.memory[int64(uint32(v6))+148:], uint32(v6+i32(136)))
										_ = m.fn404(v6+i32(100), i32(1068881), v6+i32(148))
										t90 := int32(load32(m.memory[int64(uint32(v6))+164:]))
										t91 := int32(load32(m.memory[int64(uint32(v6))+168:]))
										m.fn16(t90, t91)
										t92 := int32(load32(m.memory[int64(uint32(v6))+136:]))
										t93 := int32(load32(m.memory[int64(uint32(v6))+140:]))
										m.fn16(t92, t93)
										goto l29
									}
								case 3:
									t122 := int32(load32(m.memory[int64(uint32(v4))+4:]))
									t123 := int32(load32(m.memory[int64(uint32(v4))+8:]))
									m.fn829(v6+i32(80), v9, t122, t123)
									t124 := int32(load32(m.memory[int64(uint32(v6))+84:]))
									v4 = t124
									t125 := int32(load32(m.memory[int64(uint32(v6))+80:]))
									v19 = t125
									if v19 == 0 {
										goto l29
									}
									store32(m.memory[int64(uint32(v6))+164:], uint32(v19))
									store32(m.memory[int64(uint32(v6))+168:], uint32(v4))
									store32(m.memory[int64(uint32(v6))+152:], uint32(i32(1)))
									store32(m.memory[int64(uint32(v6))+148:], uint32(v6+i32(164)))
									_ = m.fn404(v6+i32(100), i32(1068636), v6+i32(148))
									goto l29
								case 4:
									if v16 == 0 {
										goto l29
									}
									t101 := int32(load32(m.memory[int64(uint32(v4))+4:]))
									t102 := v15
									t103 := v13
									t104 := v12
									v22 = t101
									t105 := int32(load32(m.memory[int64(uint32(v4))+8:]))
									t106 := v22
									v20 = t105
									t107 := m.fn29(t103, t104, t106, v20)
									v26 = t107
									v4 = t102 & int32(v26)
									v27 = int64(uint64(v26)>>25) & i64(127) * i64(72340172838076673)
									v21 = i32(0)
								l33:
									{
										t108 := int64(load64(m.memory[uint32(v14+v4):]))
										v28 = t108
										v26 = v28 ^ v27
										v26 = (v26 ^ i64(-1)) & (v26 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
									l32:
										{
											if v26 == 0 {
												if !(v28&(v28<<1)&i64(-0x7f7f7f7f7f7f7f80) == 0) {
													goto l29
												}
												t114 := v4
												v21 = v21 + i32(8)
												v4 = (t114 + v21) & v15
												goto l33
											}
											t109 := v22
											t110 := v20
											v19 = v14 - (int32(uint32(int64(bits.TrailingZeros64(uint64(v26))))>>3)+v4)&v15<<4
											t111 := int32(load32(m.memory[uint32(v19+i32(-12)):]))
											t112 := int32(load32(m.memory[uint32(v19+i32(-8)):]))
											t113 := m.fn15(t109, t110, t111, t112)
											if t113 != 0 {
												store32(m.memory[int64(uint32(v6))+164:], uint32(v19+i32(-4)))
												store32(m.memory[int64(uint32(v6))+152:], uint32(i32(141)))
												store32(m.memory[int64(uint32(v6))+148:], uint32(v6+i32(164)))
												_ = m.fn404(v6+i32(100), i32(1068589), v6+i32(148))
												goto l29
											}
											v26 = (v26 + i64(-1)) & v26
											goto l32
										}
									}
								case 5:
									switch v17 {
									case 1:
										m.fn74(v6+i32(100), i32(32))
										goto l29
									case 2:
										m.fn74(v6+i32(100), i32(10))
										goto l29
									default:
										m.fn75(v6+i32(100), i32(1084920), i32(2))
										goto l29
									}
								default:
									v20 = i32(0)
									{
										if uint32(v1) >= uint32(v2) {
											goto l10
										}
										{
											v21 = v10 + v1<<4
											t18 := int32(load32(m.memory[uint32(v21):]))
											v19 = t18
											t19 := v19 + i32(-0x7fffffff)
											var p20 int32
											if v19 < i32(-0x7ffffffb) {
												p20 = 1
											}
											v22 = p20
											p21 := i32(0)
											if v22 != 0 {
												p21 = t19
											}
											v19 = p21
											if uint32(v19) > uint32(i32(4)) {
												goto l11
											}
											v20 = i32(65536)
											if i32_shl(i32(1), v19)&i32(22) != 0 {
												goto l10
											}
										}
									l11:
										v20 = i32(0)
										if v22 != 0 {
											goto l10
										}
										t22 := m.fn822(v21+i32(12), i32(1287584))
										p23 := i32(0)
										if t22 != 0 {
											p23 = i32(65536)
										}
										v20 = p23
									}
								l10:
									t24 := int32(load32(m.memory[uint32(v4+i32(8)):]))
									v19 = t24
									t25 := int32(load32(m.memory[uint32(v4+i32(4)):]))
									v22 = t25
									t26 := int32(load32(m.memory[int64(uint32(v4))+12:]))
									t27 := v6
									v4 = t26
									store32(m.memory[int64(uint32(t27))+136:], uint32(v4))
									t28 := m.fn823(v6+i32(136), i32(1287584))
									if t28 != 0 {
										t115 := int32(load32(m.memory[int64(uint32(v6))+108:]))
										v4 = t115
										if v4 != 0 {
											goto l37
										}
										v4 = i32(1)
										goto l38
									}
									m.fn824(v6+i32(40), v22, v19)
									t29 := int32(load32(m.memory[int64(uint32(v6))+44:]))
									v21 = t29
									m.fn628(v6+i32(32), v22, v19)
									t30 := int32(load32(m.memory[int64(uint32(v6))+36:]))
									v20 = t30
									t31 := v6 + i32(24)
									t32 := v22
									t33 := v19
									v23 = v19 - v21
									m.fn825(t31, t32, t33, v23, i32(1084868))
									t34 := int32(load32(m.memory[int64(uint32(v6))+28:]))
									v21 = t34
									t35 := int32(load32(m.memory[int64(uint32(v6))+24:]))
									v24 = t35
									m.fn787(v6+i32(16), v23, v20, v22, v19)
									t36 := int32(load32(m.memory[int64(uint32(v6))+16:]))
									v25 = t36
									if v25 == 0 {
										m.fn556(v22, v19, v23, v20, i32(1084884))
										panic("unreachable")
									}
									t37 := int32(load32(m.memory[int64(uint32(v6))+20:]))
									v23 = t37
									m.fn826(v6+i32(8), v20, v22, v19, i32(1084900))
									t38 := int32(load32(m.memory[int64(uint32(v6))+12:]))
									v19 = t38
									t39 := int32(load32(m.memory[int64(uint32(v6))+8:]))
									v22 = t39
									if v21 == 0 {
										goto l14
									}
									m.fn75(v6+i32(100), v24, v21)
								l14:
									if v23 == 0 {
										goto l15
									}
									if v4&i32(0x1000000) != 0 {
										m.fn801(v25, v23, v6+i32(100))
										goto l15
									}
									store32(m.memory[int64(uint32(v6))+172:], uint32(i32(0)))
									store64(m.memory[int64(uint32(v6))+164:], uint64(i64(0x100000000)))
									if v4&i32(65536) != 0 {
										m.fn75(v6+i32(164), i32(1084916), i32(2))
										goto l18
									}
									goto l18
								}
							}
						l37:
							t116 := int32(load32(m.memory[int64(uint32(v6))+104:]))
							t117 := m.fn806(t116, v4, i32(10))
							v4 = t117
						}
					l38:
						m.fn803(v6+i32(148), v22, v19, v3, v20|v4|v8)
						t118 := int32(load32(m.memory[int64(uint32(v6))+152:]))
						t119 := v6 + i32(100)
						v4 = t118
						t120 := int32(load32(m.memory[int64(uint32(v6))+156:]))
						m.fn75(t119, v4, t120)
						t121 := int32(load32(m.memory[int64(uint32(v6))+148:]))
						m.fn16(t121, v4)
						goto l29
					}
				l25:
					t128 := int32(load32(m.memory[int64(uint32(v6))+124:]))
					t129 := int32(load32(m.memory[int64(uint32(v6))+128:]))
					m.fn16(t128, t129)
				}
			l24:
				t130 := int32(load32(m.memory[int64(uint32(v6))+112:]))
				t131 := int32(load32(m.memory[int64(uint32(v6))+116:]))
				m.fn16(t130, t131)
				goto l29
			}
		l23:
			t132 := int32(load32(m.memory[int64(uint32(v6))+124:]))
			m.fn16(t132, v19)
			t133 := int32(load32(m.memory[int64(uint32(v6))+112:]))
			t134 := int32(load32(m.memory[int64(uint32(v6))+116:]))
			m.fn16(t133, t134)
			goto l29
		}
	l18:
		if v4&i32(1) == 0 {
			goto l39
		}
		m.fn75(v6+i32(164), i32(1084918), i32(2))
	l39:
		if v4&i32(256) == 0 {
			goto l40
		}
		m.fn74(v6+i32(164), i32(42))
	l40:
		t135 := int32(load32(m.memory[int64(uint32(v6))+172:]))
		v20 = t135
		t136 := int32(load32(m.memory[int64(uint32(v6))+168:]))
		v4 = t136
		store32(m.memory[int64(uint32(v6))+156:], uint32(i32(0)))
		store64(m.memory[int64(uint32(v6))+148:], uint64(i64(0x100000000)))
		m.fn831(v6+i32(148), v4, v4+v20)
		t137 := int32(load32(m.memory[int64(uint32(v6))+148:]))
		v24 = t137
		t138 := int32(load32(m.memory[int64(uint32(v6))+156:]))
		v29 = t138
		t139 := int32(load32(m.memory[int64(uint32(v6))+152:]))
		v21 = t139
		m.fn75(v6+i32(100), v4, v20)
		m.fn803(v6+i32(148), v25, v23, v3, v7)
		t140 := int32(load32(m.memory[int64(uint32(v6))+152:]))
		t141 := v6 + i32(100)
		v20 = t140
		t142 := int32(load32(m.memory[int64(uint32(v6))+156:]))
		m.fn75(t141, v20, t142)
		t143 := int32(load32(m.memory[int64(uint32(v6))+148:]))
		m.fn16(t143, v20)
		m.fn75(v6+i32(100), v21, v29)
		m.fn16(v24, v21)
		t144 := int32(load32(m.memory[int64(uint32(v6))+164:]))
		m.fn16(t144, v4)
	}
l15:
	if v19 == 0 {
		goto l29
	}
	m.fn75(v6+i32(100), v22, v19)
	goto l29
}
func (m *Module) fn803(v0, v1, v2, v3, v4 int32) {
	var v5, v6, v7, v8, v9, v10, v11, v12, v13, v14, v15, v16, v17, v18, v19, v20, v21, v22, v23, v24, v25, v26, v27, v28, v29, v30 int32
	t0 := m.g0
	v5 = t0 - i32(96)
	m.g0 = v5
	store32(m.memory[int64(uint32(v5))+88:], uint32(v1))
	store32(m.memory[int64(uint32(v5))+92:], uint32(v1+v2))
	{
		{
			t1 := m.fn48(v5 + i32(88))
			v1 = t1
			if v1 != i32(-1) {
				goto l0
			}
			v6 = i32(4)
			v1 = i32(0)
			v7 = i32(0)
			goto l1
		}
	l0:
		t2 := int32(load32(m.memory[int64(uint32(v5))+92:]))
		t3 := v5 + i32(24)
		v8 = t2
		t4 := int32(load32(m.memory[int64(uint32(v5))+88:]))
		t5 := v8
		v9 = t4
		v10 = t5 - v9
		t6 := int32(uint32(v10) >> 2)
		var p7 int32
		if v10&i32(3) != i32(0) {
			p7 = 1
		}
		v10 = t6 + p7
		p8 := i32(3)
		if uint32(v10) > uint32(i32(3)) {
			p8 = v10
		}
		m.fn59(t3, p8+i32(1), i32(4), i32(4))
		t9 := int32(load32(m.memory[int64(uint32(v5))+24:]))
		v11 = t9
		t10 := int32(load32(m.memory[int64(uint32(v5))+28:]))
		v10 = t10
		store32(m.memory[uint32(v10):], uint32(v1))
		store32(m.memory[int64(uint32(v5))+44:], uint32(i32(1)))
		store32(m.memory[int64(uint32(v5))+40:], uint32(v10))
		store32(m.memory[int64(uint32(v5))+36:], uint32(v11))
		store32(m.memory[int64(uint32(v5))+80:], uint32(v8))
		store32(m.memory[int64(uint32(v5))+76:], uint32(v9))
		v8 = i32(4)
		v1 = i32(1)
	l4:
		{
			t11 := m.fn48(v5 + i32(76))
			v9 = t11
			if v9 == i32(-1) {
				goto l2
			}
			{
				t12 := int32(load32(m.memory[int64(uint32(v5))+36:]))
				if v1 != t12 {
					goto l3
				}
				t13 := int32(load32(m.memory[int64(uint32(v5))+80:]))
				t14 := int32(load32(m.memory[int64(uint32(v5))+76:]))
				t15 := v5 + i32(36)
				t16 := v1
				v10 = t13 - t14
				t17 := int32(uint32(v10) >> 2)
				var p18 int32
				if v10&i32(3) != i32(0) {
					p18 = 1
				}
				m.fn62(t15, t16, t17+p18+i32(1), i32(4), i32(4))
				t19 := int32(load32(m.memory[int64(uint32(v5))+40:]))
				v10 = t19
			}
		l3:
			store32(m.memory[uint32(v10+v8):], uint32(v9))
			t20 := v5
			v1 = v1 + i32(1)
			store32(m.memory[int64(uint32(t20))+44:], uint32(v1))
			v8 = v8 + i32(4)
			goto l4
		}
	l2:
		t21 := int32(load32(m.memory[int64(uint32(v5))+40:]))
		v6 = t21
		t22 := int32(load32(m.memory[int64(uint32(v5))+36:]))
		v7 = t22
	}
l1:
	v8 = i32(0)
l6:
	if v8 == i32(40) {
		v11 = v6 + v1<<2
		v9 = i32(0)
		t23 := int32(load32(m.memory[int64(uint32(v5))+44:]))
		v12 = t23
		t24 := int32(load32(m.memory[int64(uint32(v5))+52:]))
		v13 = t24
		t25 := int32(load32(m.memory[int64(uint32(v5))+60:]))
		v14 = t25
		t26 := int32(load32(m.memory[int64(uint32(v5))+68:]))
		v15 = t26
		t27 := int32(load32(m.memory[int64(uint32(v5))+36:]))
		v16 = t27
		v10 = v6
	l9:
		v8 = v10
		if v8 == v11 {
			m.fn140(v5+i32(76), v2+i32(8))
			v10 = i32(0)
			v22 = v4 & i32(65536)
			t29 := v22
			v23 = v4 & i32(256)
			var p30 int32
			if t29|v23 != i32(0) {
				p30 = 1
			}
			v24 = p30
			v25 = v3 & i32(255)
			var p31 int32
			if v25 != i32(0) {
				p31 = 1
			}
			v26 = p31
			v2 = v26 | (v4 ^ i32(1))
			v27 = v4 & i32(0x1010000)
			v28 = int32(uint32(v22) >> 16)
			var p32 int32
			if uint32(v4) > uint32(i32(0xffffff)) {
				p32 = 1
			}
			v29 = p32
		l16:
			v11 = v2
			v8 = v10
			if uint32(v8) >= uint32(v1) {
				t52 := int32(load32(m.memory[int64(uint32(v5))+84:]))
				store32(m.memory[int64(uint32(v0))+8:], uint32(t52))
				t53 := int64(load64(m.memory[int64(uint32(v5))+76:]))
				store64(m.memory[uint32(v0):], uint64(t53))
				m.fn44(v7, v6)
				m.g0 = v5 + i32(96)
				return
			}
			{
				t33 := m.fn622(v6, v1, v8, i32(1084688))
				t34 := int32(load32(m.memory[uint32(t33):]))
				v9 = t34
				if v9 != i32(10) {
					t35 := m.fn630(v9)
					v2 = t35 ^ i32(1)
					v3 = i32(-1)
					v4 = v28
					{
						v10 = v8 + i32(1)
						var p36 int32
						if uint32(v10) >= uint32(v1) {
							p36 = 1
						}
						v30 = p36
						if v30 != 0 {
							goto l17
						}
						t37 := int32(load32(m.memory[uint32(v6+v10<<2):]))
						v3 = t37
						t38 := m.fn630(v3)
						v4 = t38 ^ i32(1)
					}
				l17:
					v2 = v2 | v11
					switch v9 + i32(-33) {
					case 0:
						if uint32(v10) < uint32(v1) {
							goto l38
						}
						if v22 != 0 {
							goto l27
						}
						goto l38
					case 1, 3, 4, 6, 7, 8, 11:
						goto l19
					case 2:
						if v11&i32(1) != 0 {
							goto l38
						}
					l47:
						{
							t60 := v1
							v11 = v8
							if t60 == v11 {
								goto l27
							}
							v8 = v11 + i32(1)
							t61 := m.fn622(v6, v1, v11, i32(1074604))
							t62 := int32(load32(m.memory[uint32(t61):]))
							if t62 == i32(35) {
								goto l47
							}
						}
						if uint32(v11) >= uint32(v1) {
							goto l27
						}
						t63 := int32(load32(m.memory[uint32(v6+v11<<2):]))
						t64 := m.fn630(t63)
						if t64 != 0 {
							goto l27
						}
						goto l38
					case 5:
						m.fn815(v5, v6, v1, v8, i32(1084720))
						t77 := int32(load32(m.memory[uint32(v5):]))
						v8 = t77
						{
							{
								t78 := int32(load32(m.memory[int64(uint32(v5))+4:]))
								v4 = t78
								if uint32(v4) <= uint32(i32(1)) {
									goto l50
								}
								t79 := int32(load32(m.memory[int64(uint32(v8))+4:]))
								if t79 == i32(35) {
									goto l51
								}
							}
						l50:
							v11 = v8 + i32(4)
							v8 = i32(0)
							t80 := v4
							var p81 int32
							if v4 != i32(0) {
								p81 = 1
							}
							v3 = t80 - p81
						l53:
							if v3 == v8 {
								goto l38
							}
							{
								t82 := int32(load32(m.memory[uint32(v11):]))
								v4 = t82
								if uint32(v4+i32(-48)) < uint32(i32(10)) {
									goto l52
								}
								if uint32(v4&i32(2097119)+i32(-65)) <= uint32(i32(25)) {
									goto l52
								}
								if v8 == 0 {
									goto l38
								}
								if v4 == i32(59) {
									goto l51
								}
								goto l38
							}
						l52:
							v11 = v11 + i32(4)
							v8 = v8 + i32(1)
							goto l53
						}
					l51:
						m.fn75(v5+i32(76), i32(1084736), i32(5))
						goto l16
					case 9:
						var p74 int32
						if v23 == 0 {
							p74 = 1
						}
						if p74&v11 == 0 {
							goto l27
						}
						if v4 == 0 {
							goto l38
						}
						if v22 != 0 {
							goto l27
						}
						t75 := v16
						var p76 int32
						if uint32(v20) > uint32(v8) {
							p76 = 1
						}
						if t75&p76 != 0 {
							goto l27
						}
						goto l38
					case 10:
						if (v11|v4)&i32(1) == 0 {
							goto l27
						}
						goto l38
					case 12:
						if v11&i32(1) != 0 {
							goto l38
						}
						if v4 == 0 {
							goto l27
						}
						m.fn815(v5+i32(8), v6, v1, v8, i32(1084744))
						t57 := int32(load32(m.memory[int64(uint32(v5))+8:]))
						t58 := int32(load32(m.memory[int64(uint32(v5))+12:]))
						t59 := m.fn816(t57, t58, i32(45))
						if t59 == 0 {
							goto l38
						}
						goto l27
					default:
						switch v9 + i32(-91) {
						case 0:
							if v27 != 0 {
								goto l27
							}
							t83 := v15
							var p84 int32
							if uint32(v19) > uint32(v8) {
								p84 = 1
							}
							if t83&p84 != 0 {
								goto l27
							}
							goto l38
						case 1:
							goto l27
						case 2:
							goto l28
						case 3:
							goto l19
						case 4:
							{
								if v8 == 0 {
									goto l48
								}
								t67 := m.fn622(v6, v1, v8+i32(-1), i32(1084704))
								t68 := int32(load32(m.memory[uint32(t67):]))
								t69 := m.fn817(t68)
								v11 = t69
								t70 := m.fn818(v3)
								v3 = t70
								if v23 != 0 {
									goto l27
								}
								if v11&v3 != 0 {
									goto l38
								}
								if v4^i32(1) == 0 {
									goto l49
								}
								goto l38
							}
						l48:
							_ = m.fn818(v3)
							if v23 != 0 {
								goto l27
							}
							if v4 == 0 {
								goto l38
							}
						l49:
							if v22 != 0 {
								goto l27
							}
							t72 := v12
							var p73 int32
							if uint32(v17) > uint32(v8) {
								p73 = 1
							}
							if t72&p73 != 0 {
								goto l27
							}
							goto l38
						case 5:
							t85 := v24
							t86 := v14
							var p87 int32
							if uint32(v18) > uint32(v8) {
								p87 = 1
							}
							if t85|t86&p87 != 0 {
								goto l27
							}
							goto l38
						default:
							switch v9 + i32(-60) {
							case 0:
								if v30 != 0 {
									goto l38
								}
								if uint32(v3&i32(2097119)+i32(-65)) < uint32(i32(26)) {
									goto l27
								}
								v8 = v3 + i32(-33)
								if uint32(v8) > uint32(i32(30)) {
									goto l38
								}
								if i32_shl(i32(1), v8)&i32(0x40004001) != 0 {
									goto l27
								}
								goto l38
							case 1:
								if v11&i32(1) != 0 {
									goto l38
								}
								m.fn815(v5+i32(16), v6, v1, v8, i32(1084760))
								t54 := int32(load32(m.memory[int64(uint32(v5))+16:]))
								t55 := int32(load32(m.memory[int64(uint32(v5))+20:]))
								t56 := m.fn816(t54, t55, i32(61))
								if t56 != 0 {
									goto l27
								}
								goto l38
							case 2:
								if v11&i32(1) == 0 {
									goto l27
								}
								goto l38
							default:
								switch v9 + i32(-124) {
								case 0:
									if v25 == i32(2) {
										goto l27
									}
									goto l38
								case 2:
									if v23 != 0 {
										goto l27
									}
									if v4 == 0 {
										goto l38
									}
									if v22 != 0 {
										goto l27
									}
									t65 := v13
									var p66 int32
									if uint32(v21) > uint32(v8) {
										p66 = 1
									}
									if t65&p66 != 0 {
										goto l27
									}
									goto l38
								default:
									goto l19
								}
							}
						}
					}
				l19:
					;
					var p39 int32
					if uint32(v9+i32(-58)) < uint32(i32(-10)) {
						p39 = 1
					}
					if (p39|v11)&i32(1) != 0 {
						goto l38
					}
					p40 := v1
					if uint32(v8) > uint32(v1) {
						p40 = v8
					}
					v3 = p40
					v4 = i32(0)
					v11 = v8
				l46:
					{
						{
							if v3 == v11 {
								goto l39
							}
							t41 := m.fn622(v6, v1, v11, i32(1084776))
							t42 := int32(load32(m.memory[uint32(t41):]))
							if uint32(t42+i32(-48)) < uint32(i32(10)) {
								v4 = v4 + i32(1)
								v11 = v11 + i32(1)
								goto l46
							}
						}
					l39:
						if uint32(v11) >= uint32(v1) {
							goto l38
						}
						{
							t43 := m.fn622(v6, v1, v11, i32(1084792))
							t44 := int32(load32(m.memory[uint32(t43):]))
							if t44 == i32(46) {
								goto l41
							}
							t45 := m.fn622(v6, v1, v11, i32(1084808))
							t46 := int32(load32(m.memory[uint32(t45):]))
							if t46 != i32(41) {
								goto l38
							}
						}
					l41:
						{
							v3 = v11 + i32(1)
							if uint32(v3) >= uint32(v1) {
								goto l42
							}
							t47 := int32(load32(m.memory[uint32(v6+v3<<2):]))
							t48 := m.fn630(t47)
							if t48 == 0 {
								goto l38
							}
						}
					l42:
						if uint32(v11) < uint32(v8) {
							m.fn151(v8, v11, v1, i32(1084824))
							panic("unreachable")
						}
						m.fn47(v5+i32(76), v11-v8)
						if v11 == v8 {
							goto l44
						}
						v8 = v6 + v8<<2
					l45:
						{
							t49 := int32(load32(m.memory[uint32(v8):]))
							m.fn74(v5+i32(76), t49)
							v8 = v8 + i32(4)
							v4 = v4 + i32(-1)
							if v4 != 0 {
								goto l45
							}
						}
					l44:
						m.fn74(v5+i32(76), i32(92))
						t50 := m.fn622(v6, v1, v11, i32(1084840))
						t51 := int32(load32(m.memory[uint32(t50):]))
						m.fn74(v5+i32(76), t51)
						v10 = v3
						goto l16
					}
				}
				m.fn74(v5+i32(76), i32(10))
				v10 = v8 + i32(1)
				v2 = v26 & v11
				goto l16
			}
		l28:
			if v29 == 0 {
				goto l38
			}
		l27:
			m.fn74(v5+i32(76), i32(92))
		l38:
			m.fn74(v5+i32(76), v9)
			goto l16
		}
		v9 = v9 + i32(1)
		v10 = v8 + i32(4)
		{
			t28 := int32(load32(m.memory[uint32(v8):]))
			v8 = t28
			switch v8 + i32(-93) {
			case 1:
				goto l9
			case 2:
				v17 = v9 + i32(-1)
				v12 = i32(1)
				goto l9
			case 3:
				v18 = v9 + i32(-1)
				v14 = i32(1)
				goto l9
			case 0:
				v19 = v9 + i32(-1)
				v15 = i32(1)
				goto l9
			default:
				if v8 != i32(42) {
					if v8 != i32(126) {
						goto l9
					}
					v21 = v9 + i32(-1)
					v13 = i32(1)
					goto l9
				}
				v20 = v9 + i32(-1)
				v16 = i32(1)
				goto l9
			}
		}
	}
	store32(m.memory[uint32(v5+i32(36)+v8):], uint32(i32(0)))
	v8 = v8 + i32(8)
	goto l6
}
func (m *Module) fn804(v0, v1 int32, v2 int64) {
	var v3 int32
	t0 := m.g0
	v3 = t0 - i32(16)
	m.g0 = v3
	switch v1 & i32(255) {
	default:
		m.fn51(v0, i32(1108000), i32(1))
		goto l6
	case 1:
		m.fn809(v0, v2)
		goto l6
	case 2:
		m.fn810(v0, v2)
		goto l6
	case 3:
		m.fn810(v3+i32(4), v2)
		t1 := int32(load32(m.memory[int64(uint32(v3))+8:]))
		t2 := v0
		v1 = t1
		t3 := int32(load32(m.memory[int64(uint32(v3))+12:]))
		m.fn811(t2, v1, t3)
		t4 := int32(load32(m.memory[int64(uint32(v3))+4:]))
		m.fn16(t4, v1)
		goto l6
	case 4:
		m.fn812(v0, v2)
		goto l6
	case 5:
		m.fn812(v3+i32(4), v2)
		t5 := int32(load32(m.memory[int64(uint32(v3))+8:]))
		t6 := v0
		v1 = t5
		t7 := int32(load32(m.memory[int64(uint32(v3))+12:]))
		m.fn811(t6, v1, t7)
		t8 := int32(load32(m.memory[int64(uint32(v3))+4:]))
		m.fn16(t8, v1)
	}
l6:
	m.g0 = v3 + i32(16)
}
func (m *Module) fn805(v0, v1, v2, v3 int32) {
	var v4 int32
	t0 := m.g0
	v4 = t0 - i32(80)
	m.g0 = v4
	store16(m.memory[int64(uint32(v4))+44:], uint16(i32(1)))
	store32(m.memory[int64(uint32(v4))+40:], uint32(i32(0)))
	store32(m.memory[int64(uint32(v4))+32:], uint32(v1))
	store32(m.memory[int64(uint32(v4))+28:], uint32(v2))
	store32(m.memory[int64(uint32(v4))+24:], uint32(v1))
	store32(m.memory[int64(uint32(v4))+20:], uint32(v2))
	store32(m.memory[int64(uint32(v4))+16:], uint32(i32(0)))
	store32(m.memory[int64(uint32(v4))+36:], uint32(v1+v2))
	m.fn807(v4+i32(8), v4+i32(16))
	{
		t1 := int32(load32(m.memory[int64(uint32(v4))+8:]))
		if t1 == 0 {
			goto l0
		}
		t2 := int32(load32(m.memory[int64(uint32(v4))+12:]))
		v1 = t2
		t3 := int64(load64(m.memory[int64(uint32(v4))+40:]))
		store64(m.memory[int64(uint32(v4))+72:], uint64(t3))
		t4 := int64(load64(m.memory[int64(uint32(v4))+32:]))
		store64(m.memory[int64(uint32(v4))+64:], uint64(t4))
		t5 := int64(load64(m.memory[int64(uint32(v4))+24:]))
		store64(m.memory[int64(uint32(v4))+56:], uint64(t5))
		t6 := int64(load64(m.memory[int64(uint32(v4))+16:]))
		store64(m.memory[int64(uint32(v4))+48:], uint64(t6))
	l2:
		{
			m.fn807(v4, v4+i32(48))
			t7 := int32(load32(m.memory[uint32(v4):]))
			if t7 == 0 {
				goto l1
			}
			t8 := int32(load32(m.memory[int64(uint32(v4))+4:]))
			t9 := v1
			v2 = t8
			p10 := v2
			if uint32(v1) > uint32(v2) {
				p10 = t9
			}
			v1 = p10
			goto l2
		}
	l1:
		t11 := v3
		v1 = v1 + i32(1)
		p12 := v1
		if uint32(v3) > uint32(v1) {
			p12 = t11
		}
		v3 = p12
	}
l0:
	m.fn808(v0, i32(1084856), v3)
	m.g0 = v4 + i32(80)
}
func (m *Module) fn806(v0, v1, v2 int32) int32 {
	var v3 int32
	t0 := m.g0
	v3 = t0 - i32(16)
	m.g0 = v3
	store32(m.memory[int64(uint32(v3))+12:], uint32(i32(0)))
	m.memory[int64(uint32(v3))+12] = byte(v2)
	t1 := m.fn156(v0, v1, v3+i32(12), i32(1))
	v2 = t1
	m.g0 = v3 + i32(16)
	return v2
}
