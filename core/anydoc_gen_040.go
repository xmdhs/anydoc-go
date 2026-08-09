package core

import (
	"math/bits"
)

func (m *Module) fn1752(v0, v1 int32) {
	var v2 int32
	t0 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	v2 = t0
	m.fn1702(v0, i32(1))
	t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	m.fn279(v1, v2+t1)
	store32(m.memory[int64(uint32(v0))+8:], uint32(v2+i32(1)))
}
func (m *Module) fn1753(v0, v1, v2, v3 int32) {
	var v4 int32
	var v5 int64
	var v6 int32
	var v7 int64
	var v8 int32
	var v9 int64
	t0 := m.g0
	v4 = t0 - i32(16)
	m.g0 = v4
	if v2 != 0 {
		goto l0
	}
	v5 = i64(1)
	goto l1
l0:
	{
		t1 := int32(m.memory[uint32(v1)])
		v6 = t1
		switch v6 + i32(-43) {
		case 0, 2:
			goto l2
		default:
			v5 = i64(257)
			if v2 != i32(1) {
				goto l4
			}
			if v6 == i32(43) {
				goto l1
			}
		l4:
			t2 := v1
			var p3 int32
			if v6 == i32(43) {
				p3 = 1
			}
			v6 = p3
			v1 = t2 + v6
			v2 = v2 - v6
			if uint32(v2) < uint32(i32(9)) {
				v6 = i32(0)
			l9:
				{
					if v2 == 0 {
						goto l6
					}
					t7 := int32(m.memory[uint32(v1)])
					m.fn1754(v4, t7, v3)
					t8 := int32(load32(m.memory[uint32(v4):]))
					if t8 != i32(1) {
						goto l1
					}
					v1 = v1 + i32(1)
					v2 = v2 + i32(-1)
					t9 := int32(load32(m.memory[int64(uint32(v4))+4:]))
					v6 = t9 + v6*v3
					goto l9
				}
			}
			v6 = i32(0)
			v7 = int64(uint32(v3))
		l8:
			{
				if v2 == 0 {
					goto l6
				}
				t4 := int32(m.memory[uint32(v1)])
				m.fn1754(v4+i32(8), t4, v3)
				t5 := int32(load32(m.memory[int64(uint32(v4))+8:]))
				v8 = t5
				v9 = int64(uint32(v6)) * v7
				if int32(int64(uint64(v9)>>32)) != 0 {
					goto l7
				}
				if v8&i32(1) == 0 {
					goto l1
				}
				v1 = v1 + i32(1)
				v2 = v2 + i32(-1)
				t6 := int32(load32(m.memory[int64(uint32(v4))+12:]))
				v8 = t6
				v6 = v8 + int32(v9)
				if uint32(v6) >= uint32(v8) {
					goto l8
				}
			}
			v5 = i64(513)
			goto l1
		l7:
			p10 := i64(257)
			if v8&i32(1) != 0 {
				p10 = i64(513)
			}
			v5 = p10
		}
	}
l1:
	store64(m.memory[uint32(v0):], uint64(v5))
	goto l10
l6:
	m.memory[uint32(v0)] = byte(i32(255))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v6))
	goto l10
l2:
	m.memory[uint32(v0)] = byte(i32(0))
l10:
	m.g0 = v4 + i32(16)
}
func (m *Module) fn1754(v0, v1, v2 int32) {
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
func (m *Module) fn1755(v0, v1, v2 int32) int32 {
	t0 := m.fn1851(v0, v1, v2)
	var p1 int32
	if t0 == 0 {
		p1 = 1
	}
	return p1
}
func (m *Module) fn1756(v0, v1 int32) int32 {
	var v2 int32
	v2 = i32(1)
	{
		t0 := int32(load32(m.memory[uint32(v0):]))
		v0 = t0
		t1 := m.fn1757(v0, v1)
		if t1 != 0 {
			goto l0
		}
		t2 := int32(load32(m.memory[uint32(v1):]))
		t3 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t4 := int32(load32(m.memory[int64(uint32(t3))+12:]))
		t5 := m.t0[uint(t4)].(func(int32, int32, int32) int32)(t2, i32(1284184), i32(2))
		if t5 != 0 {
			goto l0
		}
		t6 := m.fn1757(v0+i32(4), v1)
		v2 = t6
	}
l0:
	return v2
}
func (m *Module) fn1757(v0, v1 int32) int32 {
	var v2 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		v2 = t0
		if v2&i32(0x2000000) != 0 {
			t4 := int32(load32(m.memory[uint32(v0):]))
			t5 := m.fn467(t4, v1)
			return t5
		}
		{
			if v2&i32(0x4000000) != 0 {
				t2 := int32(load32(m.memory[uint32(v0):]))
				t3 := m.fn466(t2, v1)
				return t3
			}
			t1 := m.fn72(v0, v1)
			return t1
		}
	}
}
func (m *Module) fn1758(v0 int32) int32 {
	var v1, v2, v3, v4, v5, v6, v7, v8, v9, v10 int32
	var v11 int64
	var v12, v13 int32
	var v14 int64
	var v15, v16, v17 int32
	t0 := m.g0
	v1 = t0 - i32(48)
	m.g0 = v1
	{
		t1 := int32(load32(m.memory[int64(uint32(v0))+12:]))
		v2 = t1
		v3 = v2 + i32(1)
		if v3 == 0 {
			m.fn242()
			panic("unreachable")
		}
		{
			t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t3 := v3
			v4 = t2
			t4 := v4
			v5 = v4 + i32(1)
			v6 = int32(uint32(v5) >> 3)
			p5 := v6 * i32(7)
			if uint32(v4) < uint32(i32(8)) {
				p5 = t4
			}
			v7 = p5
			if uint32(t3) <= uint32(int32(uint32(v7)>>1)) {
				t29 := v6
				var p30 int32
				if v5&i32(7) != i32(0) {
					p30 = 1
				}
				v8 = t29 + p30
				t31 := int32(load32(m.memory[uint32(v0):]))
				v6 = t31
				v3 = v6
			l16:
				if v8 != 0 {
					t43 := int64(load64(m.memory[uint32(v3):]))
					t44 := v3
					v11 = t43
					store64(m.memory[uint32(t44):], uint64(int64(uint64(v11^i64(-1))>>7)&i64(72340172838076673)+(v11|i64(0x7f7f7f7f7f7f7f7f))))
					v3 = v3 + i32(8)
					v8 = v8 + i32(-1)
					goto l16
				}
				{
					if uint32(v5) < uint32(i32(8)) {
						goto l9
					}
					t32 := int64(load64(m.memory[uint32(v6):]))
					store64(m.memory[uint32(v6+v5):], uint64(t32))
					goto l10
				}
			l9:
				if v5 == 0 {
					goto l10
				}
				memory_copy(m.memory, uint32(v6+i32(8)), uint32(v6), uint32(v5))
			l10:
				v8 = i32(0)
			l12:
				{
					v3 = v8
					if v3 == v5 {
						store32(m.memory[int64(uint32(v0))+8:], uint32(v7-v2))
						goto l7
					}
					v8 = v3 + i32(1)
					v15 = v6 + v3
					t33 := int32(m.memory[uint32(v15)])
					if t33 != i32(128) {
						goto l12
					}
					v10 = v6 - v3<<3 + i32(-8)
					v12 = v6 + (v3^i32(-1))<<3
				l15:
					{
						t34 := int64(load64(m.memory[uint32(v10):]))
						t35 := v3
						t36 := v4
						v11 = t34
						v13 = int32(v11)
						v16 = t36 & v13
						t37 := m.fn26(v6, v4, v11)
						t38 := t35 - v16
						v9 = t37
						if uint32((t38^(v9-v16))&v4) < uint32(i32(8)) {
							t42 := v15
							v13 = int32(uint32(v13) >> 25)
							m.memory[uint32(t42)] = byte(v13)
							m.memory[uint32(v6+(v3+i32(-8))&v4+i32(8))] = byte(v13)
							goto l12
						}
						v16 = v6 + v9
						t39 := int32(m.memory[uint32(v16)])
						v17 = t39
						t40 := v16
						v13 = int32(uint32(v13) >> 25)
						m.memory[uint32(t40)] = byte(v13)
						m.memory[uint32(v6+(v9+i32(-8))&v4+i32(8))] = byte(v13)
						v13 = v6 - v9<<3 + i32(-8)
						{
							if v17 != i32(255) {
								m.fn1619(v12, v13, i32(2))
								goto l15
							}
							m.memory[uint32(v15)] = byte(i32(255))
							m.memory[uint32(v6+(v3+i32(-8))&v4+i32(8))] = byte(i32(255))
							t41 := int64(load64(m.memory[uint32(v12):]))
							store64(m.memory[uint32(v13):], uint64(t41))
							goto l12
						}
					}
				}
			}
			t6 := v1 + i32(32)
			v6 = v7 + i32(1)
			p7 := v3
			if uint32(v6) > uint32(v3) {
				p7 = v6
			}
			m.fn1759(t6, p7)
			t8 := int32(load32(m.memory[int64(uint32(v1))+36:]))
			v5 = t8
			t9 := int32(load32(m.memory[int64(uint32(v1))+32:]))
			v8 = t9
			if v8 == 0 {
				goto l2
			}
			t10 := int32(load32(m.memory[int64(uint32(v1))+40:]))
			v9 = t10
			t11 := int32(load32(m.memory[int64(uint32(v1))+44:]))
			store32(m.memory[int64(uint32(v1))+28:], uint32(t11))
			store32(m.memory[int64(uint32(v1))+24:], uint32(v9))
			store32(m.memory[int64(uint32(v1))+20:], uint32(v5))
			store32(m.memory[int64(uint32(v1))+16:], uint32(v8))
			store64(m.memory[int64(uint32(v1))+8:], uint64(i64(0x800000008)))
			store32(m.memory[int64(uint32(v1))+4:], uint32(v0+i32(16)))
			t12 := int32(load32(m.memory[uint32(v0):]))
			v10 = t12
			t13 := int64(load64(m.memory[uint32(v10):]))
			v11 = (t13 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
			v12 = v1 + i32(16)
			v3 = i32(0)
			v4 = v2
			v6 = v10
		l6:
			if v4 == 0 {
				store32(m.memory[int64(uint32(v1))+28:], uint32(v2))
				store32(m.memory[int64(uint32(v1))+24:], uint32(v9-v2))
				m.fn1619(v0, v12, i32(4))
				t22 := int32(load32(m.memory[int64(uint32(v1))+20:]))
				v3 = t22
				if v3 == 0 {
					goto l7
				}
				t23 := int32(load32(m.memory[int64(uint32(v1))+8:]))
				t24 := int32(load32(m.memory[int64(uint32(v1))+12:]))
				m.fn1616(v1+i32(32), t23, t24, v3+i32(1))
				t25 := int32(load32(m.memory[int64(uint32(v1))+16:]))
				t26 := int32(load32(m.memory[int64(uint32(v1))+40:]))
				t27 := int32(load32(m.memory[int64(uint32(v1))+32:]))
				t28 := int32(load32(m.memory[int64(uint32(v1))+36:]))
				m.fn40(t25-t26, t27, t28)
				goto l7
			}
		l5:
			{
				if v11 != i64(0) {
					t15 := v8
					t16 := v8
					t17 := v5
					v13 = v10 - (int32(uint32(int64(bits.TrailingZeros64(uint64(v11))))>>3)+v3)<<3 + i32(-8)
					t18 := int64(load64(m.memory[uint32(v13):]))
					v14 = t18
					t19 := m.fn26(t16, t17, v14)
					v15 = t19
					t20 := t15 + v15
					v16 = int32(uint32(int32(v14)) >> 25)
					m.memory[uint32(t20)] = byte(v16)
					m.memory[uint32(v8+(v15+i32(-8))&v5+i32(8))] = byte(v16)
					t21 := int64(load64(m.memory[uint32(v13):]))
					store64(m.memory[uint32(v8-v15<<3+i32(-8)):], uint64(t21))
					v4 = v4 + i32(-1)
					v11 = (v11 + i64(-1)) & v11
					goto l6
				}
				v3 = v3 + i32(8)
				v6 = v6 + i32(8)
				t14 := int64(load64(m.memory[uint32(v6):]))
				v11 = (t14 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
				goto l5
			}
		}
	}
l7:
	v5 = i32(-1)
l2:
	m.g0 = v1 + i32(48)
	return v5
}
func (m *Module) fn1759(v0, v1 int32) {
	var v2, v3, v4, v5, v6 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	{
		{
			{
				if uint32(v1) < uint32(i32(15)) {
					goto l0
				}
				if uint32(v1) > uint32(i32(0x1fffffff)) {
					m.fn242()
					panic("unreachable")
				}
				t1 := int32(uint32(v1<<3) / uint32(i32(7)))
				v1 = i32_shr_u(i32(-1), int32(bits.LeadingZeros32(uint32(t1+i32(-1))))) + i32(1)
				goto l2
			}
		l0:
			p2 := v1&i32(8) + i32(8)
			if uint32(v1) < uint32(i32(4)) {
				p2 = i32(4)
			}
			v1 = p2
		}
	l2:
		m.fn243(v2, i32(8), v1)
		t3 := int32(load32(m.memory[int64(uint32(v2))+8:]))
		v3 = t3
		t4 := int32(load32(m.memory[int64(uint32(v2))+4:]))
		v1 = t4
		{
			{
				t5 := int32(load32(m.memory[uint32(v2):]))
				v4 = t5
				if v4 != 0 {
					goto l3
				}
				store32(m.memory[int64(uint32(v0))+8:], uint32(v3))
				store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
				store32(m.memory[uint32(v0):], uint32(i32(0)))
				goto l4
			}
		l3:
			t6 := int32(load32(m.memory[int64(uint32(v2))+12:]))
			v5 = t6
			v6 = v1 + i32(9)
			if v6 == 0 {
				goto l5
			}
			memory_fill(m.memory, uint32(v4), i32(255), uint32(v6))
		l5:
			store32(m.memory[int64(uint32(v0))+12:], uint32(v5))
			store32(m.memory[int64(uint32(v0))+8:], uint32(v3))
			store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
			store32(m.memory[uint32(v0):], uint32(v4))
		}
	l4:
		m.g0 = v2 + i32(16)
		return
	}
}
func (m *Module) fn1760(v0, v1 int32) {
	var v2 int32
	{
		{
			t0 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			v2 = t0
			if v2 != 0 {
				goto l0
			}
			v1 = i32(0)
			goto l1
		}
	l0:
		t1 := v1
		v2 = v2 + i32(-1)
		store32(m.memory[int64(uint32(t1))+8:], uint32(v2))
		t2 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t3 := int32(load32(m.memory[uint32(t2+v2<<2):]))
		v2 = t3
		v1 = i32(1)
	}
l1:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v2))
	store32(m.memory[uint32(v0):], uint32(v1))
}
func (m *Module) fn1761(v0, v1, v2, v3, v4 int32) {
	var v5, v6, v7, v8, v9, v10 int32
	v5 = i32(0)
	if uint32(v3) < uint32(v4) {
		goto l0
	}
	goto l1
l0:
	if uint32(v4-v3) > uint32(i32(3)) {
		{
			t2 := int32(load32(m.memory[uint32(v3):]))
			v7 = t2
			t3 := v7
			v8 = v1 & i32(255) * i32(16843009)
			v6 = t3 ^ v8
			if (i32(16843008)-v6|v6)&i32(-2139062144) != i32(-2139062144) {
				goto l10
			}
			t4 := v7
			v9 = v2 & i32(255) * i32(16843009)
			v6 = t4 ^ v9
			if (i32(16843008)-v6|v6)&i32(-2139062144) != i32(-2139062144) {
				goto l10
			}
			v10 = v4 + i32(-4)
			v3 = v3&i32(-4) + i32(4)
		l9:
			{
				if uint32(v3) > uint32(v10) {
					goto l6
				}
				t5 := int32(load32(m.memory[uint32(v3):]))
				v7 = t5
				v6 = v7 ^ v8
				if (i32(16843008)-v6|v6)&i32(-2139062144) != i32(-2139062144) {
					goto l6
				}
				v6 = v7 ^ v9
				if (i32(16843008)-v6|v6)&i32(-2139062144) == i32(-2139062144) {
					v3 = v3 + i32(4)
					goto l9
				}
			}
		l6:
			v1 = v1 & i32(255)
		l8:
			{
				if uint32(v3) >= uint32(v4) {
					goto l1
				}
				t6 := int32(m.memory[uint32(v3)])
				t7 := v1
				v6 = t6
				if t7 == v6 {
					goto l3
				}
				if v2&i32(255) == v6 {
					goto l3
				}
				v3 = v3 + i32(1)
				goto l8
			}
		}
	l10:
		{
			if uint32(v3) >= uint32(v4) {
				goto l1
			}
			t8 := int32(m.memory[uint32(v3)])
			t9 := v1 & i32(255)
			v6 = t8
			if t9 == v6 {
				goto l3
			}
			if v2&i32(255) == v6 {
				goto l3
			}
			v3 = v3 + i32(1)
			goto l10
		}
	}
	v1 = v1 & i32(255)
l4:
	{
		if uint32(v3) >= uint32(v4) {
			goto l1
		}
		t0 := int32(m.memory[uint32(v3)])
		t1 := v1
		v6 = t0
		if t1 == v6 {
			goto l3
		}
		if v2&i32(255) == v6 {
			goto l3
		}
		v3 = v3 + i32(1)
		goto l4
	}
l3:
	v5 = i32(1)
l1:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
	store32(m.memory[uint32(v0):], uint32(v5))
}
func (m *Module) fn1762(v0, v1, v2, v3 int32) {
	var v4, v5, v6, v7, v8, v9, v10 int32
	t0 := m.g0
	v4 = t0 - i32(16)
	m.g0 = v4
	v5 = v2 + v3
	t1 := int32(m.memory[uint32(v1)])
	v6 = t1
	v7 = v2
l9:
	m.fn146(v4+i32(8), i32(62), i32(39), i32(34), v7, v5)
	{
		t2 := int32(load32(m.memory[int64(uint32(v4))+8:]))
		if t2 == i32(1) {
			t3 := int32(load32(m.memory[int64(uint32(v4))+12:]))
			v9 = t3
			v10 = v9 - v2
			if uint32(v10) < uint32(v3) {
				v7 = v9 + i32(1)
				t4 := int32(m.memory[uint32(v2+v10)])
				v9 = t4
				switch v6 & i32(255) {
				default:
					if v9 == i32(34) {
						v6 = i32(2)
						goto l8
					}
					v8 = i32(1)
					if v9 != i32(39) {
						v6 = i32(0)
						if v9 != i32(62) {
							goto l9
						}
						goto l1
					}
					v6 = i32(1)
					goto l8
				case 1:
					v6 = i32(1)
					if v9 != i32(39) {
						goto l9
					}
					goto l10
				case 2:
					v6 = i32(2)
					if v9 != i32(34) {
						goto l9
					}
					goto l10
				}
			l10:
				v6 = i32(0)
			l8:
				m.memory[uint32(v1)] = byte(v6)
				goto l9
			}
			m.fn158(v10, v3, i32(1283596))
			panic("unreachable")
		}
		v8 = i32(0)
		goto l1
	}
l1:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v10))
	store32(m.memory[uint32(v0):], uint32(v8))
	m.g0 = v4 + i32(16)
}
func (m *Module) fn1763(v0, v1, v2 int32) {
	var v3, v4 int32
	t0 := m.g0
	v3 = t0 - i32(32)
	m.g0 = v3
	m.memory[int64(uint32(v3))+28] = byte(i32(62))
	store32(m.memory[int64(uint32(v3))+24:], uint32(v1+v2))
	store32(m.memory[int64(uint32(v3))+20:], uint32(v1))
	store32(m.memory[int64(uint32(v3))+16:], uint32(v1))
l3:
	{
		m.fn155(v3+i32(8), v3+i32(16))
		{
			t1 := int32(load32(m.memory[int64(uint32(v3))+8:]))
			if t1 == i32(1) {
				goto l0
			}
			v1 = i32(0)
			goto l1
		}
	l0:
		t2 := int32(load32(m.memory[int64(uint32(v3))+12:]))
		v4 = t2
		if uint32(v4) > uint32(v2) {
			m.fn151(i32(0), v4, v2, i32(1282780))
			panic("unreachable")
		}
		t3 := m.fn1061(v1, v4, i32(1282444), i32(2))
		if t3 == 0 {
			goto l3
		}
	}
	v1 = i32(1)
	v4 = v4 + i32(1)
l1:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v4))
	store32(m.memory[uint32(v0):], uint32(v1))
	m.g0 = v3 + i32(32)
}
func (m *Module) fn1764(v0, v1, v2 int32) {
	var v3, v4, v5 int32
	t0 := m.g0
	v3 = t0 - i32(16)
	m.g0 = v3
	v4 = i32(1)
	v5 = i32(2)
	{
		t1 := m.fn159(v1, v2, i32(1282796), i32(2))
		if t1 != 0 {
			goto l0
		}
		m.fn1763(v3+i32(8), v1, v2)
		t2 := int32(load32(m.memory[int64(uint32(v3))+12:]))
		v5 = t2
		t3 := int32(load32(m.memory[int64(uint32(v3))+8:]))
		v4 = t3
	}
l0:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
	store32(m.memory[uint32(v0):], uint32(v4))
	m.g0 = v3 + i32(16)
}
func (m *Module) fn1765(v0, v1, v2, v3, v4 int32) {
	if uint32(v2) < uint32(v1) {
		goto l0
	}
	if uint32(v2) > uint32(i32(9)) {
		goto l0
	}
	store32(m.memory[int64(uint32(v0))+4:], uint32(v2-v1))
	store32(m.memory[uint32(v0):], uint32(v3+v1))
	return
l0:
	m.fn151(v1, v2, i32(9), v4)
	panic("unreachable")
}
func (m *Module) fn1766(v0, v1, v2, v3 int32) {
	var v4, v5, v6 int32
	t0 := m.g0
	v4 = t0 - i32(16)
	m.g0 = v4
	{
		if v3 == 0 {
			goto l0
		}
		{
			t1 := int32(m.memory[uint32(v2)])
			v5 = t1
			if v5 != i32(63) {
				if uint32(v3) < uint32(i32(3)) {
					goto l0
				}
				if v5 != i32(33) {
					goto l0
				}
				t2 := int32(m.memory[int64(uint32(v2))+2])
				v5 = t2
				{
					{
						t3 := int32(m.memory[int64(uint32(v2))+1])
						v6 = t3
						if v6 != i32(45) {
							goto l3
						}
						if v5 == i32(45) {
							v2 = i32(3)
							store16(m.memory[uint32(v1):], uint16(i32(3)))
							v3 = i32(1)
							goto l2
						}
					}
				l3:
					{
						if uint32(v3) > uint32(i32(7)) {
							switch v6 + i32(-65) {
							default:
								goto l8
							case 4:
								switch v5 + i32(-76) {
								default:
									goto l8
								case 0:
									t8 := int32(m.memory[int64(uint32(v2))+3])
									if t8 != i32(69) {
										goto l8
									}
									t9 := int32(m.memory[int64(uint32(v2))+4])
									if t9 != i32(77) {
										goto l8
									}
									t10 := int32(m.memory[int64(uint32(v2))+5])
									if t10 != i32(69) {
										goto l8
									}
									t11 := int32(m.memory[int64(uint32(v2))+6])
									if t11 != i32(78) {
										goto l8
									}
									t12 := int32(m.memory[int64(uint32(v2))+7])
									if t12 != i32(84) {
										goto l8
									}
									m.memory[uint32(v1)] = byte(i32(5))
									goto l12
								case 2:
									t13 := int32(m.memory[int64(uint32(v2))+3])
									if t13 != i32(84) {
										goto l8
									}
									t14 := int32(m.memory[int64(uint32(v2))+4])
									if t14 != i32(73) {
										goto l8
									}
									t15 := int32(m.memory[int64(uint32(v2))+5])
									if t15 != i32(84) {
										goto l8
									}
									t16 := int32(m.memory[int64(uint32(v2))+6])
									if t16 == i32(89) {
										goto l6
									}
									goto l8
								}
							case 0:
								if v5 != i32(84) {
									goto l8
								}
								t17 := int32(m.memory[int64(uint32(v2))+3])
								if t17&i32(255) != i32(84) {
									goto l8
								}
								t18 := int32(m.memory[int64(uint32(v2))+4])
								if t18 != i32(76) {
									goto l8
								}
								t19 := int32(m.memory[int64(uint32(v2))+5])
								if t19 != i32(73) {
									goto l8
								}
								t20 := int32(m.memory[int64(uint32(v2))+6])
								if t20 != i32(83) {
									goto l8
								}
								t21 := int32(m.memory[int64(uint32(v2))+7])
								if t21 != i32(84) {
									goto l8
								}
								store16(m.memory[uint32(v1):], uint16(i32(6)))
								goto l12
							}
						}
						if v3 != i32(7) {
							goto l0
						}
						if v6 != i32(69) {
							goto l0
						}
						if v5 != i32(78) {
							goto l0
						}
						t4 := int32(m.memory[int64(uint32(v2))+3])
						if t4 != i32(84) {
							goto l0
						}
						t5 := int32(m.memory[int64(uint32(v2))+4])
						if t5 != i32(73) {
							goto l0
						}
						t6 := int32(m.memory[int64(uint32(v2))+5])
						if t6 != i32(84) {
							goto l0
						}
						t7 := int32(m.memory[int64(uint32(v2))+6])
						if t7 != i32(89) {
							goto l0
						}
						goto l6
					}
				l8:
					if v3 == i32(8) {
						goto l0
					}
					if v6 != i32(78) {
						goto l0
					}
					if v5 != i32(79) {
						goto l0
					}
					t22 := int32(m.memory[int64(uint32(v2))+3])
					if t22 != i32(84) {
						goto l0
					}
					t23 := int32(m.memory[int64(uint32(v2))+4])
					if t23 != i32(65) {
						goto l0
					}
					t24 := int32(m.memory[int64(uint32(v2))+5])
					if t24 != i32(84) {
						goto l0
					}
					t25 := int32(m.memory[int64(uint32(v2))+6])
					if t25 != i32(73) {
						goto l0
					}
					t26 := int32(m.memory[int64(uint32(v2))+7])
					if t26 != i32(79) {
						goto l0
					}
					t27 := int32(m.memory[int64(uint32(v2))+8])
					if t27 != i32(78) {
						goto l0
					}
					store16(m.memory[uint32(v1):], uint16(i32(6)))
					v3 = i32(1)
					v2 = i32(9)
					goto l2
				}
			l12:
				v3 = i32(1)
				v2 = i32(8)
				goto l2
			l6:
				store16(m.memory[uint32(v1):], uint16(i32(6)))
				v3 = i32(1)
				v2 = i32(7)
				goto l2
			}
			store16(m.memory[uint32(v1):], uint16(i32(4)))
			v2 = i32(1)
			v3 = i32(1)
			goto l2
		}
	l0:
		m.fn881(v4+i32(8), i32(62), v2, v3)
		v3 = i32(1)
		{
			t28 := int32(load32(m.memory[int64(uint32(v4))+8:]))
			if t28&i32(1) != 0 {
				goto l13
			}
			v3 = i32(0)
			goto l2
		}
	l13:
		t29 := int32(load32(m.memory[int64(uint32(v4))+12:]))
		v2 = t29
		m.memory[uint32(v1)] = byte(i32(1))
		v2 = v2 + i32(1)
	}
l2:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v2))
	store32(m.memory[uint32(v0):], uint32(v3))
	m.g0 = v4 + i32(16)
}
func (m *Module) fn1767(v0, v1 int32) {
	var v2, v3, v4, v5, v6, v7, v8 int32
	t0 := m.g0
	v2 = t0 - i32(48)
	m.g0 = v2
	t1 := int32(load32(m.memory[int64(uint32(v1))+40:]))
	t2 := int32(load32(m.memory[int64(uint32(v1))+44:]))
	m.fn1084(v2+i32(28), v1, t1, t2)
	{
		t3 := int32(load32(m.memory[int64(uint32(v2))+28:]))
		v3 = t3
		switch v3 + i32(2) {
		case 0:
			store32(m.memory[uint32(v0):], uint32(i32(-3)))
			goto l3
		case 1:
			t4 := int32(load32(m.memory[int64(uint32(v2))+40:]))
			store32(m.memory[int64(uint32(v0))+12:], uint32(t4))
			t5 := int64(load64(m.memory[int64(uint32(v2))+32:]))
			store64(m.memory[int64(uint32(v0))+4:], uint64(t5))
			store32(m.memory[uint32(v0):], uint32(i32(-2)))
			goto l3
		default:
			t6 := int32(load32(m.memory[int64(uint32(v2))+44:]))
			v4 = t6
			t7 := int32(load32(m.memory[int64(uint32(v2))+40:]))
			v5 = t7
			t8 := int32(load32(m.memory[int64(uint32(v1))+40:]))
			t9 := v2 + i32(16)
			v6 = t8
			t10 := int32(load32(m.memory[int64(uint32(v1))+44:]))
			t11 := v6
			v7 = t10
			t12 := int32(load32(m.memory[int64(uint32(v2))+32:]))
			t13 := int32(load32(m.memory[int64(uint32(v2))+36:]))
			m.fn1085(t9, t11, v7, t12, t13)
			t14 := int32(load32(m.memory[int64(uint32(v2))+20:]))
			v1 = t14
			t15 := int32(load32(m.memory[int64(uint32(v2))+16:]))
			v8 = t15
			{
				if v3 != i32(3) {
					goto l4
				}
				v3 = i32(1)
				v4 = i32(0)
				goto l5
			l4:
				m.fn1085(v2+i32(8), v6, v7, v5, v4)
				t16 := int32(load32(m.memory[int64(uint32(v2))+8:]))
				v3 = t16
				t17 := int32(load32(m.memory[int64(uint32(v2))+12:]))
				v4 = t17
			}
		l5:
			store32(m.memory[int64(uint32(v0))+12:], uint32(v8))
			store32(m.memory[int64(uint32(v0))+8:], uint32(v4))
			store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
			store32(m.memory[uint32(v0):], uint32(i32(-1)))
			store32(m.memory[int64(uint32(v0))+16:], uint32(v1))
		}
	}
l3:
	m.g0 = v2 + i32(48)
}
func (m *Module) fn1768(v0 int32) {
	var v1 int32
	t0 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	t1 := int32(load32(m.memory[int64(uint32(v0))+12:]))
	m.fn1086(t0, t1)
	{
		t2 := int32(load32(m.memory[int64(uint32(v0))+20:]))
		v1 = t2
		if v1 == 0 {
			return
		}
		t3 := int32(load32(m.memory[int64(uint32(v0))+24:]))
		m.fn1087(v1, t3)
	}
}
func (m *Module) fn1769(v0, v1 int32) int32 {
	var v2, v3 int32
	v2 = i32(0)
l3:
	if v1 != v2 {
		{
			t0 := int32(m.memory[uint32(v0+v2)])
			v3 = t0 + i32(-9)
			if uint32(v3) > uint32(i32(23)) {
				goto l2
			}
			if i32_shl(i32(1), v3)&i32(8388627) != 0 {
				goto l1
			}
		}
	l2:
		v2 = v2 + i32(1)
		goto l3
	}
	v2 = v1
	goto l1
l1:
	return v2
}
func (m *Module) fn1770(v0, v1 int32) {
	var v2, v3 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v1))+16:]))
		v2 = t0
		t1 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		t2 := v2
		v3 = t1
		if uint32(t2) <= uint32(v3) {
			goto l0
		}
		m.fn151(i32(0), v2, v3, i32(1281712))
		panic("unreachable")
	}
l0:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v2))
	t3 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	store32(m.memory[uint32(v0):], uint32(t3))
}
func (m *Module) fn1771(v0, v1, v2 int32) {
	m.fn1583(v0, v1, v2-v1)
}
func (m *Module) fn1772(v0, v1, v2, v3, v4, v5, v6 int32) {
	var v7 int32
	var v8 int64
	t0 := m.g0
	v7 = t0 - i32(16)
	m.g0 = v7
	{
		{
			t1 := int32(m.memory[int64(uint32(v1))+36])
			if t1 == i32(1) {
				goto l0
			}
			v1 = i32(0)
			v5 = i32(-1)
			goto l3
		}
	l0:
		m.fn1773(v7+i32(4), v1, v2, v3, v4, v5)
		{
			t2 := int32(m.memory[int64(uint32(v7))+4])
			if t2 == i32(255) {
				goto l2
			}
			v5 = i32(-1)
			t3 := int64(load32(m.memory[int64(uint32(v7))+12:]))
			v8 = t3
			t4 := int32(load32(m.memory[int64(uint32(v7))+8:]))
			v6 = t4
			t5 := int32(load32(m.memory[int64(uint32(v7))+4:]))
			v1 = t5
			goto l3
		}
	l2:
		v5 = i32(3)
		t6 := int32(load32(m.memory[int64(uint32(v7))+12:]))
		v6 = t6
		t7 := int32(load32(m.memory[int64(uint32(v7))+8:]))
		v1 = t7
	}
l3:
	store64(m.memory[int64(uint32(v0))+12:], uint64(v8))
	store32(m.memory[int64(uint32(v0))+8:], uint32(v6))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
	store32(m.memory[uint32(v0):], uint32(v5))
	m.g0 = v7 + i32(16)
}
func (m *Module) fn1773(v0, v1, v2, v3, v4, v5 int32) {
	var v6, v7, v8, v9, v10, v11, v12, v13 int32
	var v14, v15 int64
	t0 := m.g0
	v6 = t0 - i32(32)
	m.g0 = v6
	{
		t1 := int32(m.memory[int64(uint32(v1))+37])
		if t1 == 0 {
			goto l0
		}
		v7 = v1 + i32(8)
		{
			t2 := int32(load32(m.memory[int64(uint32(v1))+16:]))
			v8 = t2
			if uint32(v8) > uint32(i32(31)) {
				v11 = v1 + i32(20)
				{
					t10 := int32(load32(m.memory[int64(uint32(v1))+20:]))
					if t10 != 0 {
						goto l7
					}
					m.fn1759(v6+i32(16), v8<<1)
					t11 := int64(load64(m.memory[int64(uint32(v6))+24:]))
					store64(m.memory[int64(uint32(v6))+8:], uint64(t11))
					t12 := int64(load64(m.memory[int64(uint32(v6))+16:]))
					store64(m.memory[uint32(v6):], uint64(t12))
					v13 = v8 << 3
					t13 := int32(load32(m.memory[int64(uint32(v1))+12:]))
					v8 = t13
				l10:
					{
						if v13 == 0 {
							goto l8
						}
						t14 := int32(load32(m.memory[int64(uint32(v8))+4:]))
						v12 = t14
						t15 := int32(load32(m.memory[uint32(v8):]))
						t16 := v12
						v9 = t15
						if uint32(t16) < uint32(v9) {
							goto l9
						}
						if uint32(v12) > uint32(v3) {
							goto l9
						}
						t17 := m.fn1774(v2+v9, v12-v9)
						_ = m.fn1775(v6, t17)
						v13 = v13 + i32(-8)
						v8 = v8 + i32(8)
						goto l10
					}
				l8:
					t19 := int64(load64(m.memory[int64(uint32(v6))+8:]))
					t20 := v6
					v14 = t19
					store64(m.memory[int64(uint32(t20))+24:], uint64(v14))
					t21 := int64(load64(m.memory[uint32(v6):]))
					t22 := v6
					v15 = t21
					store64(m.memory[int64(uint32(t22))+16:], uint64(v15))
					store64(m.memory[int64(uint32(v11))+8:], uint64(v14))
					store64(m.memory[uint32(v11):], uint64(v15))
				}
			l7:
				if uint32(v5) < uint32(v4) {
					goto l11
				}
				if uint32(v5) > uint32(v3) {
					goto l11
				}
				t23 := v11
				v9 = v2 + v4
				t24 := v9
				v10 = v5 - v4
				t25 := m.fn1774(t24, v10)
				t26 := m.fn1775(t23, t25)
				if t26 != 0 {
					goto l12
				}
				t27 := int32(load32(m.memory[int64(uint32(v1))+16:]))
				v8 = t27 << 3
				t28 := int32(load32(m.memory[int64(uint32(v1))+12:]))
				v1 = t28
			l14:
				{
					if v8 == 0 {
						goto l12
					}
					t29 := int32(load32(m.memory[int64(uint32(v1))+4:]))
					v12 = t29
					t30 := int32(load32(m.memory[uint32(v1):]))
					t31 := v12
					v13 = t30
					if uint32(t31) < uint32(v13) {
						goto l13
					}
					if uint32(v12) > uint32(v3) {
						goto l13
					}
					v8 = v8 + i32(-8)
					v1 = v1 + i32(8)
					t32 := m.fn882(v2+v13, v12-v13, v9, v10)
					if t32 == 0 {
						goto l14
					}
				}
				store32(m.memory[int64(uint32(v0))+8:], uint32(v13))
				store32(m.memory[int64(uint32(v0))+4:], uint32(v4))
				m.memory[uint32(v0)] = byte(i32(4))
				goto l6
			}
			var p3 int32
			if uint32(v5) < uint32(v4) {
				p3 = 1
			}
			var p4 int32
			if uint32(v5) > uint32(v3) {
				p4 = 1
			}
			v9 = p3 | p4
			v8 = v8 << 3
			v10 = v2 + v4
			v11 = v5 - v4
			t5 := int32(load32(m.memory[int64(uint32(v1))+12:]))
			v1 = t5
		l5:
			{
				if v8 == 0 {
					m.fn1776(v7, v4, v5)
					goto l0
				}
				t6 := int32(load32(m.memory[int64(uint32(v1))+4:]))
				v12 = t6
				t7 := int32(load32(m.memory[uint32(v1):]))
				t8 := v12
				v13 = t7
				if uint32(t8) < uint32(v13) {
					goto l3
				}
				if uint32(v12) > uint32(v3) {
					goto l3
				}
				if v9 != 0 {
					m.fn151(v4, v5, v3, i32(1281648))
					panic("unreachable")
				}
				v1 = v1 + i32(8)
				v8 = v8 + i32(-8)
				t9 := m.fn882(v2+v13, v12-v13, v10, v11)
				if t9 == 0 {
					goto l5
				}
			}
			store32(m.memory[int64(uint32(v0))+8:], uint32(v13))
			store32(m.memory[int64(uint32(v0))+4:], uint32(v4))
			m.memory[uint32(v0)] = byte(i32(4))
			goto l6
		}
	l12:
		m.fn1776(v7, v4, v5)
		store32(m.memory[int64(uint32(v0))+8:], uint32(v5))
		store32(m.memory[int64(uint32(v0))+4:], uint32(v4))
		m.memory[uint32(v0)] = byte(i32(255))
		goto l6
	l3:
		m.fn151(v13, v12, v3, i32(1281632))
		panic("unreachable")
	l9:
		m.fn151(v9, v12, v3, i32(1281664))
		panic("unreachable")
	l11:
		m.fn151(v4, v5, v3, i32(1282480))
		panic("unreachable")
	l13:
		m.fn151(v13, v12, v3, i32(1281680))
		panic("unreachable")
	}
l0:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v5))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v4))
	m.memory[uint32(v0)] = byte(i32(255))
l6:
	m.g0 = v6 + i32(32)
}
func (m *Module) fn1774(v0, v1 int32) int64 {
	var v2, v3, v4 int32
	var v5, v6, v7, v8, v9 int64
	t0 := m.g0
	v2 = t0 - i32(96)
	m.g0 = v2
	store64(m.memory[int64(uint32(v2))+56:], uint64(i64(0)))
	store64(m.memory[int64(uint32(v2))+32:], uint64(i64(0)))
	store64(m.memory[int64(uint32(v2))+24:], uint64(i64(8387220255154660723)))
	store64(m.memory[int64(uint32(v2))+16:], uint64(i64(7237128888997146477)))
	store64(m.memory[int64(uint32(v2))+8:], uint64(i64(0x6c7967656e657261)))
	store64(m.memory[uint32(v2):], uint64(i64(8317987319222330741)))
	store64(m.memory[int64(uint32(v2))+40:], uint64(i64(0)))
	store64(m.memory[int64(uint32(v2))+48:], uint64(i64(0)))
	store32(m.memory[int64(uint32(v2))+56:], uint32(v1))
	v3 = v1 & i32(0x7ffffff8)
	v4 = i32(0)
l1:
	{
		if uint32(v4) >= uint32(v3) {
			v3 = i32(4)
			{
				v1 = v1 & i32(7)
				if uint32(v1) >= uint32(i32(4)) {
					goto l2
				}
				v5 = i64(0)
				v3 = i32(0)
				goto l3
			l2:
				t5 := int64(load32(m.memory[uint32(v0+v4):]))
				v5 = t5
			}
		l3:
			{
				if uint32(v3|i32(1)) >= uint32(v1) {
					goto l4
				}
				t6 := int64(load16(m.memory[uint32(v0+v4+v3):]))
				v5 = i64_shl(t6, int64(uint32(v3<<3))) | v5
				v3 = v3 | i32(2)
			}
		l4:
			{
				if uint32(v3) >= uint32(v1) {
					goto l5
				}
				t7 := int64(m.memory[uint32(v0+(v3+v4))])
				v5 = i64_shl(t7, int64(uint32(v3<<3))) | v5
			}
		l5:
			t8 := int64(load32(m.memory[int64(uint32(v2))+56:]))
			v6 = t8
			t9 := int64(load64(m.memory[int64(uint32(v2))+16:]))
			store64(m.memory[int64(uint32(v2))+80:], uint64(t9))
			t10 := int64(load64(m.memory[int64(uint32(v2))+8:]))
			store64(m.memory[int64(uint32(v2))+72:], uint64(t10))
			t11 := int64(load64(m.memory[uint32(v2):]))
			store64(m.memory[int64(uint32(v2))+64:], uint64(t11))
			t12 := v2
			v7 = v6<<56 | v5
			t13 := int64(load64(m.memory[int64(uint32(v2))+24:]))
			store64(m.memory[int64(uint32(t12))+88:], uint64(v7^t13))
			m.fn286(v2 + i32(64))
			t14 := int64(load64(m.memory[int64(uint32(v2))+80:]))
			v5 = t14
			t15 := int64(load64(m.memory[int64(uint32(v2))+64:]))
			v8 = t15
			t16 := int64(load64(m.memory[int64(uint32(v2))+72:]))
			v9 = t16
			t17 := int64(load64(m.memory[int64(uint32(v2))+88:]))
			v6 = t17
			m.g0 = v2 + i32(96)
			v9 = v6 + (v9 ^ i64(255))
			t18 := v9
			t19 := i64_rotl(v5, i64(13))
			v5 = v5 + (v8 ^ v7)
			v7 = t19 ^ v5
			v8 = t18 + v7
			v7 = v8 ^ i64_rotl(v7, i64(17))
			t20 := i64_rotl(v7, i64(13))
			v6 = i64_rotl(v6, i64(16)) ^ v9
			v5 = v6 + i64_rotl(v5, i64(32))
			v7 = v5 + v7
			v9 = t20 ^ v7
			t21 := i64_rotl(v9, i64(17))
			v5 = i64_rotl(v6, i64(21)) ^ v5
			v6 = v5 + i64_rotl(v8, i64(32))
			v8 = v6 + v9
			v9 = t21 ^ v8
			t22 := i64_rotl(v9, i64(13))
			v5 = i64_rotl(v5, i64(16)) ^ v6
			v6 = v5 + i64_rotl(v7, i64(32))
			v7 = t22 ^ (v6 + v9)
			t23 := i64_rotl(v7, i64(17))
			v5 = i64_rotl(v5, i64(21)) ^ v6
			v6 = v5 + i64_rotl(v8, i64(32))
			v7 = v6 + v7
			return t23 ^ i64_rotl(v7, i64(32)) ^ i64_rotl(i64_rotl(v5, i64(16))^v6, i64(21)) ^ v7
		}
		t1 := int64(load64(m.memory[int64(uint32(v2))+24:]))
		t2 := int64(load64(m.memory[uint32(v0+v4):]))
		t3 := v2
		v5 = t2
		store64(m.memory[int64(uint32(t3))+24:], uint64(t1^v5))
		m.fn286(v2)
		t4 := int64(load64(m.memory[uint32(v2):]))
		store64(m.memory[uint32(v2):], uint64(v5^t4))
		v4 = v4 + i32(8)
		goto l1
	}
}
func (m *Module) fn1775(v0 int32, v1 int64) int32 {
	var v2, v3 int32
	var v4, v5 int64
	var v6, v7, v8 int32
	var v9, v10 int64
	var v11, v12 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		if t0 != 0 {
			goto l0
		}
		_ = m.fn1758(v0)
	}
l0:
	t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	v2 = t2
	v3 = v2 & int32(v1)
	v4 = int64(uint64(v1) >> 25)
	v5 = v4 & i64(127) * i64(72340172838076673)
	t3 := int32(load32(m.memory[uint32(v0):]))
	v6 = t3
	v7 = i32(0)
	v8 = i32(0)
	{
	l8:
		{
			t4 := int64(load64(m.memory[uint32(v6+v3):]))
			v9 = t4
			v10 = v9 ^ v5
			v10 = (v10 ^ i64(-1)) & (v10 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
		l3:
			{
				var p5 int32
				if v10 == 0 {
					p5 = 1
				}
				v11 = p5
				if v11 != 0 {
					goto l1
				}
				t6 := int64(load64(m.memory[uint32(v6-(int32(uint32(int64(bits.TrailingZeros64(uint64(v10))))>>3)+v3)&v2<<3+i32(-8)):]))
				if v1 == t6 {
					goto l2
				}
				v10 = (v10 + i64(-1)) & v10
				goto l3
			}
		l1:
			v10 = v9 & i64(-0x7f7f7f7f7f7f7f80)
			{
				if v7 == i32(1) {
					goto l4
				}
				if !(v10 == 0) {
					goto l5
				}
				v7 = i32(0)
				goto l6
			l5:
				v12 = (v3 + int32(uint32(int64(bits.TrailingZeros64(uint64(v10))))>>3)) & v2
			l4:
				if v10&(v9<<1) != i64(0) {
					goto l7
				}
				v7 = i32(1)
			l6:
				t7 := v3
				v8 = v8 + i32(8)
				v3 = (t7 + v8) & v2
				goto l8
			}
		l7:
		}
		{
			t8 := int32(int8(m.memory[uint32(v6+v12)]))
			v3 = t8
			if v3 < i32(0) {
				goto l9
			}
			t9 := int64(load64(m.memory[uint32(v6):]))
			t10 := v6
			v12 = int32(uint32(int64(bits.TrailingZeros64(uint64(t9&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
			t11 := int32(m.memory[uint32(t10+v12)])
			v3 = t11
		}
	l9:
		t12 := v6 + v12
		v7 = int32(v4) & i32(127)
		m.memory[uint32(t12)] = byte(v7)
		m.memory[uint32(v6+(v12+i32(-8))&v2+i32(8))] = byte(v7)
		t13 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		store32(m.memory[int64(uint32(v0))+8:], uint32(t13-v3&i32(1)))
		t14 := int32(load32(m.memory[int64(uint32(v0))+12:]))
		store32(m.memory[int64(uint32(v0))+12:], uint32(t14+i32(1)))
		store64(m.memory[uint32(v6-v12<<3+i32(-8)):], uint64(v1))
	}
l2:
	return v11
}
func (m *Module) fn1776(v0, v1, v2 int32) {
	var v3 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		v3 = t0
		t1 := int32(load32(m.memory[uint32(v0):]))
		if v3 != t1 {
			goto l0
		}
		m.fn1777(v0)
	}
l0:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v3+i32(1)))
	t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	v0 = t2 + v3<<3
	store32(m.memory[int64(uint32(v0))+4:], uint32(v2))
	store32(m.memory[uint32(v0):], uint32(v1))
}
func (m *Module) fn1777(v0 int32) {
	var v1 int32
	t0 := m.g0
	v1 = t0 - i32(16)
	m.g0 = v1
	t1 := int32(load32(m.memory[uint32(v0):]))
	m.fn1690(v1+i32(8), v0, t1, i32(1), i32(4), i32(8))
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
func (m *Module) fn1778(v0 int32) {
	var v1 int32
	t0 := m.g0
	v1 = t0 - i32(16)
	m.g0 = v1
	t1 := int32(load32(m.memory[uint32(v0):]))
	m.fn1690(v1+i32(8), v0, t1, i32(1), i32(4), i32(16))
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
func (m *Module) fn1779(v0, v1, v2 int32) {
	var v3, v4, v5, v6 int32
	t0 := m.g0
	v3 = t0 - i32(80)
	m.g0 = v3
	m.fn881(v3+i32(56), i32(13), v1, v2)
	{
		t1 := int32(load32(m.memory[int64(uint32(v3))+56:]))
		if t1 != i32(1) {
			goto l0
		}
		t2 := int32(load32(m.memory[int64(uint32(v3))+60:]))
		v4 = t2
		m.fn1064(v3+i32(48), v2)
		store32(m.memory[int64(uint32(v3))+76:], uint32(i32(0)))
		t3 := int64(load64(m.memory[int64(uint32(v3))+48:]))
		store64(m.memory[int64(uint32(v3))+68:], uint64(t3))
		m.fn1582(v3+i32(40), v1, v2, i32(0), v4, i32(1282680))
		t4 := int32(load32(m.memory[int64(uint32(v3))+44:]))
		v5 = t4
		t5 := int32(load32(m.memory[int64(uint32(v3))+40:]))
		v6 = t5
	l2:
		{
			m.fn1583(v3+i32(68), v6, v5)
			t6 := m.fn1751(v3+i32(68), v1, v2, v4, i32(10))
			t7 := v3 + i32(32)
			v5 = t6
			m.fn148(t7, v5, v1, v2, i32(1282696))
			t8 := int32(load32(m.memory[int64(uint32(v3))+32:]))
			t9 := int32(load32(m.memory[int64(uint32(v3))+36:]))
			m.fn881(v3+i32(24), i32(13), t8, t9)
			t10 := int32(load32(m.memory[int64(uint32(v3))+24:]))
			if t10 != i32(1) {
				m.fn1584(v3+i32(16), v1, v2, v5)
				{
					t18 := int32(load32(m.memory[int64(uint32(v3))+16:]))
					v2 = t18
					if v2 == 0 {
						goto l3
					}
					t19 := int32(load32(m.memory[int64(uint32(v3))+20:]))
					m.fn1583(v3+i32(68), v2, t19)
				}
			l3:
				t20 := int32(load32(m.memory[int64(uint32(v3))+76:]))
				store32(m.memory[int64(uint32(v0))+8:], uint32(t20))
				t21 := int64(load64(m.memory[int64(uint32(v3))+68:]))
				store64(m.memory[uint32(v0):], uint64(t21))
				goto l4
			}
			t11 := int32(load32(m.memory[int64(uint32(v3))+28:]))
			t12 := v3 + i32(8)
			t13 := v1
			t14 := v2
			t15 := v5
			v4 = t11 + v5
			m.fn1582(t12, t13, t14, t15, v4, i32(1282712))
			t16 := int32(load32(m.memory[int64(uint32(v3))+12:]))
			v5 = t16
			t17 := int32(load32(m.memory[int64(uint32(v3))+8:]))
			v6 = t17
			goto l2
		}
	}
l0:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v2))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
	store32(m.memory[uint32(v0):], uint32(i32(-1)))
l4:
	m.g0 = v3 + i32(80)
}
func (m *Module) fn1780(v0, v1, v2, v3 int32) int32 {
	var v4, v5, v6 int32
	t0 := m.g0
	v4 = t0 - i32(16)
	m.g0 = v4
	v5 = i32(1)
	{
		t1 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		t2 := v0
		v6 = t1
		t3 := m.t0[uint(v6)].(func(int32, int32, int32) int32)(t2, i32(1282665), i32(1))
		if t3 != 0 {
			goto l0
		}
	l7:
		{
			{
				if v3 == 0 {
					t6 := m.t0[uint(v6)].(func(int32, int32, int32) int32)(v0, i32(1282665), i32(1))
					v5 = t6
					goto l0
				}
				store32(m.memory[uint32(v4):], uint32(v2))
				{
					t4 := int32(m.memory[uint32(v2)])
					v5 = t4
					if v5&i32(254) == i32(32) {
						goto l2
					}
					if uint32((v5+i32(-35))&i32(255)) >= uint32(i32(92)) {
						goto l3
					}
				}
			l2:
				store32(m.memory[int64(uint32(v4))+4:], uint32(v5))
				store32(m.memory[int64(uint32(v4))+12:], uint32(i32(97)))
				store32(m.memory[int64(uint32(v4))+8:], uint32(v4+i32(4)))
				t5 := m.fn284(v0, v1, i32(1052692), v4+i32(8))
				if t5 != 0 {
					goto l4
				}
				goto l5
			}
		l3:
			{
				if v5 != i32(34) {
					goto l6
				}
				t7 := m.t0[uint(v6)].(func(int32, int32, int32) int32)(v0, i32(1282666), i32(2))
				if t7 == 0 {
					goto l5
				}
				goto l4
			}
		l6:
			store32(m.memory[int64(uint32(v4))+12:], uint32(i32(170)))
			store32(m.memory[int64(uint32(v4))+8:], uint32(v4))
			t8 := m.fn284(v0, v1, i32(1282668), v4+i32(8))
			if t8 != 0 {
				goto l4
			}
		}
	l5:
		v2 = v2 + i32(1)
		v3 = v3 + i32(-1)
		goto l7
	l4:
		v5 = i32(1)
	}
l0:
	m.g0 = v4 + i32(16)
	return v5
}
func (m *Module) fn1781(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(t0):]))
	t2 := m.fn467(t1, v1)
	return t2
}
func (m *Module) fn1782(v0, v1 int32) int32 {
	var v2 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	{
		t1 := int32(load32(m.memory[uint32(v0):]))
		v0 = t1
		t2 := int32(m.memory[uint32(v0)])
		switch t2 {
		default:
			t3 := int32(load32(m.memory[uint32(v1):]))
			t4 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t5 := int32(load32(m.memory[int64(uint32(t4))+12:]))
			t6 := m.t0[uint(t5)].(func(int32, int32, int32) int32)(t3, i32(1283550), i32(22))
			v1 = t6
			goto l4
		case 1:
			t7 := int32(m.memory[int64(uint32(v0))+1])
			t8 := m.fn1597(t7, v1)
			v1 = t8
			goto l4
		case 2:
			store32(m.memory[int64(uint32(v2))+4:], uint32(v0+i32(4)))
			store32(m.memory[int64(uint32(v2))+12:], uint32(i32(141)))
			store32(m.memory[int64(uint32(v2))+8:], uint32(v2+i32(4)))
			t9 := int32(load32(m.memory[uint32(v1):]))
			t10 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t11 := m.fn284(t9, t10, i32(1052778), v2+i32(8))
			v1 = t11
			goto l4
		case 3:
			store32(m.memory[int64(uint32(v2))+4:], uint32(v0+i32(4)))
			store32(m.memory[int64(uint32(v2))+12:], uint32(i32(201)))
			store32(m.memory[int64(uint32(v2))+8:], uint32(v2+i32(4)))
			t12 := int32(load32(m.memory[uint32(v1):]))
			t13 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t14 := m.fn284(t12, t13, i32(1068596), v2+i32(8))
			v1 = t14
		}
	}
l4:
	m.g0 = v2 + i32(16)
	return v1
}
func (m *Module) fn1783(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := m.fn1678(t0, v1)
	return t1
}
func (m *Module) fn1784(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(t0):]))
	t2 := m.fn1095(t1+i32(8), v1)
	return t2
}
func (m *Module) fn1785(v0, v1 int32) int32 {
	var v2 int32
	t0 := int32(load32(m.memory[uint32(v1):]))
	v2 = t0
	t1 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	t2 := int32(load32(m.memory[int64(uint32(t1))+12:]))
	v1 = t2
	{
		t3 := int32(load32(m.memory[uint32(v0):]))
		t4 := int32(m.memory[uint32(t3)])
		switch t4 {
		default:
			t5 := m.t0[uint(v1)].(func(int32, int32, int32) int32)(v2, i32(1283053), i32(34))
			return t5
		case 1:
			t6 := m.t0[uint(v1)].(func(int32, int32, int32) int32)(v2, i32(1283087), i32(69))
			return t6
		case 2:
			t7 := m.t0[uint(v1)].(func(int32, int32, int32) int32)(v2, i32(1283156), i32(62))
			return t7
		case 3:
			t8 := m.t0[uint(v1)].(func(int32, int32, int32) int32)(v2, i32(1283218), i32(55))
			return t8
		case 4:
			t9 := m.t0[uint(v1)].(func(int32, int32, int32) int32)(v2, i32(1283273), i32(53))
			return t9
		case 5:
			t10 := m.t0[uint(v1)].(func(int32, int32, int32) int32)(v2, i32(1283326), i32(53))
			return t10
		case 6:
			t11 := m.t0[uint(v1)].(func(int32, int32, int32) int32)(v2, i32(1283379), i32(49))
			return t11
		case 7:
			t12 := m.t0[uint(v1)].(func(int32, int32, int32) int32)(v2, i32(1283428), i32(61))
			return t12
		case 8:
			t13 := m.t0[uint(v1)].(func(int32, int32, int32) int32)(v2, i32(1283489), i32(61))
			return t13
		}
	}
}
func (m *Module) fn1786(v0, v1 int32) int32 {
	var v2, v3, v4 int32
	t0 := m.g0
	v2 = t0 - i32(32)
	m.g0 = v2
	t1 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	v3 = t1
	t2 := int32(load32(m.memory[uint32(v1):]))
	v1 = t2
	{
		t3 := int32(load32(m.memory[uint32(v0):]))
		v4 = t3
		t4 := int32(load32(m.memory[uint32(v4):]))
		v0 = t4
		p5 := i32(5)
		if v0 < i32(0) {
			p5 = v0 ^ i32(-0x80000000)
		}
		switch p5 {
		case 1:
			t6 := int32(load32(m.memory[int64(uint32(v3))+12:]))
			t7 := m.t0[uint(t6)].(func(int32, int32, int32) int32)(v1, i32(1283667), i32(50))
			v1 = t7
			goto l8
		case 2:
			t8 := int32(load32(m.memory[int64(uint32(v3))+12:]))
			t9 := m.t0[uint(t8)].(func(int32, int32, int32) int32)(v1, i32(1283717), i32(67))
			v1 = t9
			goto l8
		case 3:
			store32(m.memory[int64(uint32(v2))+12:], uint32(v4+i32(4)))
			store32(m.memory[int64(uint32(v2))+20:], uint32(i32(36)))
			store32(m.memory[int64(uint32(v2))+16:], uint32(v2+i32(12)))
			t10 := m.fn284(v1, v3, i32(1052717), v2+i32(16))
			v1 = t10
			goto l8
		case 4:
			store32(m.memory[int64(uint32(v2))+12:], uint32(v4+i32(4)))
			store32(m.memory[int64(uint32(v2))+20:], uint32(i32(36)))
			store32(m.memory[int64(uint32(v2))+16:], uint32(v2+i32(12)))
			t11 := m.fn284(v1, v3, i32(1067540), v2+i32(16))
			v1 = t11
			goto l8
		case 5:
			store32(m.memory[int64(uint32(v2))+8:], uint32(v4))
			store32(m.memory[int64(uint32(v2))+12:], uint32(v4+i32(12)))
			store32(m.memory[int64(uint32(v2))+28:], uint32(i32(36)))
			store32(m.memory[int64(uint32(v2))+20:], uint32(i32(36)))
			store32(m.memory[int64(uint32(v2))+24:], uint32(v2+i32(12)))
			store32(m.memory[int64(uint32(v2))+16:], uint32(v2+i32(8)))
			t12 := m.fn284(v1, v3, i32(1068183), v2+i32(16))
			v1 = t12
			goto l8
		case 6:
			t13 := int32(load32(m.memory[int64(uint32(v3))+12:]))
			t14 := m.t0[uint(t13)].(func(int32, int32, int32) int32)(v1, i32(1283784), i32(44))
			v1 = t14
			goto l8
		case 7:
			t15 := int32(load32(m.memory[int64(uint32(v3))+12:]))
			t16 := m.t0[uint(t15)].(func(int32, int32, int32) int32)(v1, i32(1283828), i32(75))
			v1 = t16
			goto l8
		default:
			{
				t17 := int32(load32(m.memory[int64(uint32(v4))+4:]))
				if t17 == i32(-1) {
					goto l9
				}
				store32(m.memory[int64(uint32(v2))+12:], uint32(v4+i32(4)))
				store32(m.memory[int64(uint32(v2))+20:], uint32(i32(36)))
				store32(m.memory[int64(uint32(v2))+16:], uint32(v2+i32(12)))
				t18 := m.fn284(v1, v3, i32(1068508), v2+i32(16))
				v1 = t18
				goto l8
			}
		l9:
			t19 := int32(load32(m.memory[int64(uint32(v3))+12:]))
			t20 := m.t0[uint(t19)].(func(int32, int32, int32) int32)(v1, i32(1283612), i32(55))
			v1 = t20
		}
	}
l8:
	m.g0 = v2 + i32(32)
	return v1
}
func (m *Module) fn1787(v0, v1, v2, v3 int32) {
	var v4, v5 int32
	t0 := m.g0
	v4 = t0 - i32(16)
	m.g0 = v4
	t1 := int32(load32(m.memory[int64(uint32(i32(0)))+1303152:]))
	v5 = t1
	store32(m.memory[int64(uint32(i32(0)))+1303152:], uint32(v5+i32(1)))
	if v5 < i32(0) {
		goto l0
	}
	{
		t2 := int32(m.memory[int64(uint32(i32(0)))+1303144])
		if t2 != 0 {
			t5 := int32(load32(m.memory[int64(uint32(v1))+24:]))
			m.t0[uint(t5)].(func(int32, int32))(v4+i32(8), v0)
			panic("unreachable")
		}
		m.memory[int64(uint32(i32(0)))+1303144] = byte(i32(1))
		t3 := int32(load32(m.memory[int64(uint32(i32(0)))+1303140:]))
		store32(m.memory[int64(uint32(i32(0)))+1303140:], uint32(t3+i32(1)))
		t4 := int32(load32(m.memory[int64(uint32(i32(0)))+1303148:]))
		v5 = t4
		if v5 <= i32(-1) {
			goto l0
		}
		v1 = v5 + i32(1)
		if v1 >= v5 {
			goto l2
		}
		m.fn633(i32(1284596), i32(28), i32(1284624))
		panic("unreachable")
	}
l2:
	store32(m.memory[int64(uint32(i32(0)))+1303148:], uint32(v1+i32(-1)))
	if v1 <= i32(0) {
		m.fn91(i32(1285284), i32(77), i32(0x139ccc))
		panic("unreachable")
	}
	m.memory[int64(uint32(i32(0)))+1303144] = byte(i32(0))
	if v2 != 0 {
		fn1788()
		panic("unreachable")
	}
l0:
	panic("unreachable")
}
func fn1788() {
	panic("unreachable")
}
func (m *Module) fn1789(v0, v1 int32) {
	store32(m.memory[uint32(v0):], uint32(i32(0)))
}
func (m *Module) fn1790(v0, v1 int32) {
	t0 := int64(load64(m.memory[int64(uint32(i32(0)))+1284444:]))
	store64(m.memory[int64(uint32(v0))+8:], uint64(t0))
	t1 := int64(load64(m.memory[int64(uint32(i32(0)))+1284436:]))
	store64(m.memory[uint32(v0):], uint64(t1))
}
func (m *Module) fn1791(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7, v8 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	t1 := int32(load32(m.memory[uint32(v0):]))
	v0 = t1
	{
		{
			{
				{
					t2 := int32(load32(m.memory[int64(uint32(v1))+8:]))
					v3 = t2
					if v3&i32(0x2000000) != 0 {
						t11 := int32(load32(m.memory[uint32(v0):]))
						v3 = t11
						v0 = i32(9)
					l6:
						{
							t12 := int32(m.memory[int64(uint32(v3&i32(15)))+1107936])
							m.memory[uint32(v2+i32(6)+v0+i32(-2))] = byte(t12)
							v0 = v0 + i32(-1)
							v3 = int32(uint32(v3) >> 4)
							if v3 != 0 {
								goto l6
							}
						}
						t13 := m.fn1638(v1, i32(1), i32(1131670), i32(2), v2+i32(6)+v0+i32(-1), i32(9)-v0)
						v0 = t13
						goto l7
					}
					t3 := int32(load32(m.memory[uint32(v0):]))
					v0 = t3
					if v3&i32(0x4000000) != 0 {
						goto l1
					}
					v3 = i32(10)
					{
						t4 := v0
						v4 = v0 >> 31
						v5 = t4 ^ v4 - v4
						if uint32(v5) < uint32(i32(1000)) {
							goto l2
						}
						v3 = i32(10)
					l3:
						{
							v6 = v2 + i32(6) + v3
							t5 := v6 + i32(-4)
							v4 = v5
							t6 := int32(uint32(v4) / uint32(i32(10000)))
							t7 := v4
							v5 = t6
							v7 = t7 - v5*i32(10000)
							t8 := int32(uint32(v7&i32(0xffff)) / uint32(i32(100)))
							v8 = t8
							t9 := int32(load16(m.memory[int64(uint32(v8<<1))+1109319:]))
							store16(m.memory[uint32(t5):], uint16(t9))
							t10 := int32(load16(m.memory[int64(uint32((v7-v8*i32(100))&i32(0xffff)<<1))+1109319:]))
							store16(m.memory[uint32(v6+i32(-2)):], uint16(t10))
							v3 = v3 + i32(-4)
							if uint32(v4) > uint32(i32(9999999)) {
								goto l3
							}
						}
					}
				l2:
					if uint32(v5) > uint32(i32(9)) {
						goto l4
					}
					v4 = v5
					goto l5
				}
			l4:
				t14 := v2 + i32(6)
				v3 = v3 + i32(-2)
				t15 := int32(uint32(v5&i32(0xffff)) / uint32(i32(100)))
				t16 := t14 + v3
				t17 := v5
				v4 = t15
				t18 := int32(load16(m.memory[int64(uint32((t17-v4*i32(100))&i32(0xffff)<<1))+1109319:]))
				store16(m.memory[uint32(t16):], uint16(t18))
			}
		l5:
			{
				if v0 == 0 {
					goto l8
				}
				if v4 == 0 {
					goto l9
				}
			l8:
				t19 := v2 + i32(6)
				v3 = v3 + i32(-1)
				t20 := int32(m.memory[int64(uint32(v4<<1))+1109320])
				m.memory[uint32(t19+v3)] = byte(t20)
			}
		l9:
			t21 := m.fn1638(v1, int32(uint32(v0^i32(-1))>>31), i32(1), i32(0), v2+i32(6)+v3, i32(10)-v3)
			v0 = t21
			goto l7
		}
	l1:
		v3 = i32(9)
	l10:
		{
			t22 := int32(m.memory[int64(uint32(v0&i32(15)))+1131672])
			m.memory[uint32(v2+i32(6)+v3+i32(-2))] = byte(t22)
			v3 = v3 + i32(-1)
			v0 = int32(uint32(v0) >> 4)
			if v0 != 0 {
				goto l10
			}
		}
		t23 := m.fn1638(v1, i32(1), i32(1131670), i32(2), v2+i32(6)+v3+i32(-1), i32(9)-v3)
		v0 = t23
	}
l7:
	m.g0 = v2 + i32(16)
	return v0
}
func (m *Module) fn1792(v0, v1 int32) {
	m.memory[int64(uint32(i32(0)))+1303608] = byte(i32(1))
	panic("unreachable")
}
func (m *Module) fn1793(v0 int32) {
	m.fn1794(v0)
	panic("unreachable")
}
func (m *Module) fn1794(v0 int32) {
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
			m.fn1787(t9, i32(1284408), t10, t11)
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
		m.fn1787(t5, i32(1284380), t6, t7)
		panic("unreachable")
	}
}
func (m *Module) fn1795(v0 int32) {
	var v1 int32
	{
		t0 := int32(load32(m.memory[uint32(v0):]))
		v1 = t0
		if v1 < i32(1) {
			return
		}
		t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		m.fn10(t1, v1, i32(1))
	}
}
func (m *Module) fn1796(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	if t0 == i32(-1) {
		t7 := int32(load32(m.memory[uint32(v1):]))
		t8 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t9 := int32(load32(m.memory[int64(uint32(v0))+12:]))
		t10 := int32(load32(m.memory[uint32(t9):]))
		v0 = t10
		t11 := int32(load32(m.memory[uint32(v0):]))
		t12 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		t13 := m.fn100(t7, t8, t11, t12)
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
