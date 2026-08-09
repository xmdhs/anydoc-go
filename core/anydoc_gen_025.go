package core

import (
	"math"
	"math/bits"
)

func (m *Module) fn1077(v0, v1, v2 int32) {
	var v3 int32
	var v4 int64
	t0 := m.g0
	v3 = t0 - i32(32)
	m.g0 = v3
	{
		t1 := int32(load32(m.memory[uint32(v2):]))
		if t1 == i32(-1) {
			t6 := int32(load32(m.memory[int64(uint32(v2))+4:]))
			t7 := int32(load32(m.memory[int64(uint32(v2))+8:]))
			m.fn198(v0, t6, t7, v1)
			goto l2
		}
		t2 := int32(load32(m.memory[int64(uint32(v2))+4:]))
		t3 := int32(load32(m.memory[int64(uint32(v2))+8:]))
		m.fn198(v3+i32(8), t2, t3, v1)
		t4 := int64(load64(m.memory[int64(uint32(v3))+12:]))
		v4 = t4
		t5 := int32(load32(m.memory[int64(uint32(v3))+8:]))
		v2 = t5
		if v2 != i32(-2) {
			goto l1
		}
		store32(m.memory[uint32(v0):], uint32(i32(-2)))
		store64(m.memory[int64(uint32(v0))+4:], uint64(v4))
		goto l2
	}
l1:
	store64(m.memory[int64(uint32(v3))+24:], uint64(v4))
	store32(m.memory[int64(uint32(v3))+20:], uint32(v2))
	m.fn490(v0, v3+i32(20))
l2:
	m.g0 = v3 + i32(32)
}
func (m *Module) fn1078(v0, v1, v2, v3 int32) {
	var v4 int32
	t0 := m.g0
	v4 = t0 - i32(48)
	m.g0 = v4
	{
		{
			{
				p1 := i32(0)
				if v1 != 0 {
					p1 = v0 + v1*i32(44) + i32(-44)
				}
				v1 = p1
				p2 := v2
				if v1 != 0 {
					p2 = v1
				}
				v1 = p2
				t3 := int32(load32(m.memory[int64(uint32(v1))+32:]))
				v0 = t3
				if v0 == 0 {
					goto l0
				}
				t4 := int32(load32(m.memory[int64(uint32(v1))+28:]))
				v2 = t4 + v0*i32(44)
				v0 = v2 + i32(-44)
				if v0 == 0 {
					goto l0
				}
				t5 := int32(load32(m.memory[uint32(v0):]))
				if t5 == i32(-1) {
					goto l1
				}
			}
		l0:
			store32(m.memory[int64(uint32(v4))+4:], uint32(i32(-1)))
			t6 := int32(load32(m.memory[int64(uint32(v3))+8:]))
			store32(m.memory[int64(uint32(v4))+16:], uint32(t6))
			t7 := int64(load64(m.memory[uint32(v3):]))
			store64(m.memory[int64(uint32(v4))+8:], uint64(t7))
			m.fn1092(v1+i32(24), v4+i32(4))
			goto l2
		}
	l1:
		t8 := int32(load32(m.memory[int64(uint32(v3))+4:]))
		t9 := v2 + i32(-40)
		v1 = t8
		t10 := int32(load32(m.memory[int64(uint32(v3))+8:]))
		m.fn75(t9, v1, t10)
		t11 := int32(load32(m.memory[uint32(v3):]))
		m.fn16(t11, v1)
	}
l2:
	m.g0 = v4 + i32(48)
}
func (m *Module) fn1079(v0 int32) {
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
	m.fn1042(v3)
	v3 = v3 + i32(44)
	goto l1
l0:
	t2 := int32(load32(m.memory[uint32(v0):]))
	m.fn136(t2, v2, i32(4), i32(44))
}
func (m *Module) fn1080(v0, v1, v2, v3 int32) {
	var v4 int32
	t0 := m.g0
	v4 = t0 - i32(16)
	m.g0 = v4
	if v0 != 0 {
		goto l0
	}
	v0 = i32(0)
	v3 = v4 + i32(12)
	goto l1
l0:
	store32(m.memory[int64(uint32(v4))+12:], uint32(v2))
	v0 = v0 * v3
	v3 = v4 + i32(8)
l1:
	store32(m.memory[uint32(v3):], uint32(v0))
	{
		t1 := int32(load32(m.memory[int64(uint32(v4))+12:]))
		v0 = t1
		if v0 == 0 {
			goto l2
		}
		t2 := int32(load32(m.memory[int64(uint32(v4))+8:]))
		v3 = t2
		if v3 == 0 {
			goto l2
		}
		m.fn10(v1, v3, v0)
	}
l2:
	m.g0 = v4 + i32(16)
}
func (m *Module) fn1081(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := m.fn1631(t0, v1)
	return t1
}
func (m *Module) fn1082(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := m.fn1095(t0, v1)
	return t1
}
func (m *Module) fn1083(v0, v1, v2, v3 int32) {
	var v4 int32
	var v5 int64
	var v6, v7 int32
	var v8 int64
	var v9, v10 int32
	var v11 int64
	var v12, v13 int32
	t0 := m.g0
	v4 = t0 - i32(80)
	m.g0 = v4
	{
		{
			t1 := int32(load32(m.memory[int64(uint32(v1))+12:]))
			if t1 == 0 {
				goto l0
			}
			t2 := int64(load64(m.memory[int64(uint32(v1))+16:]))
			t3 := int64(load64(m.memory[int64(uint32(v1))+24:]))
			t4 := m.fn314(t2, t3, v2, v3)
			v5 = t4
			t5 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			v6 = t5
			v7 = v6 & int32(v5)
			v8 = int64(uint64(v5)>>25) & i64(127) * i64(72340172838076673)
			t6 := int32(load32(m.memory[uint32(v1):]))
			v9 = t6
			v10 = i32(0)
		l6:
			{
				t7 := int64(load64(m.memory[uint32(v9+v7):]))
				v11 = t7
				v5 = v11 ^ v8
				v5 = (v5 ^ i64(-1)) & (v5 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
				{
				l3:
					if v5 == 0 {
						if !(v11&(v11<<1)&i64(-0x7f7f7f7f7f7f7f80) == 0) {
							goto l0
						}
						t17 := v7
						v10 = v10 + i32(8)
						v7 = (t17 + v10) & v6
						goto l6
					}
					{
						t8 := v2
						t9 := v3
						v12 = v9 + (i32(0)-(int32(uint32(int64(bits.TrailingZeros64(uint64(v5))))>>3)+v7)&v6)*i32(20)
						t10 := int32(load32(m.memory[uint32(v12+i32(-16)):]))
						t11 := int32(load32(m.memory[uint32(v12+i32(-12)):]))
						t12 := m.fn123(t8, t9, t10, t11)
						if t12 != 0 {
							t13 := int32(load32(m.memory[uint32(v12+i32(-4)):]))
							v7 = t13
							t14 := int32(load32(m.memory[uint32(v12+i32(-8)):]))
							v12 = t14
							t15 := int32(load32(m.memory[uint32(v12):]))
							t16 := v12
							v1 = t15 + i32(1)
							store32(m.memory[uint32(t16):], uint32(v1))
							if v1 != 0 {
								goto l4
							}
							goto l5
						}
						v5 = (v5 + i64(-1)) & v5
						goto l3
					}
				}
			}
		}
	l0:
		m.fn92(v4+i32(44), v2, v3)
		m.fn490(v4+i32(20), v4+i32(44))
		t18 := int32(load32(m.memory[int64(uint32(v4))+24:]))
		t19 := int32(load32(m.memory[int64(uint32(v4))+28:]))
		m.fn877(v4+i32(44), t18, t19)
		m.fn878(v4+i32(32), v4+i32(44), v4+i32(20))
		t20 := int32(load32(m.memory[int64(uint32(v4))+40:]))
		v7 = t20
		if v7 <= i32(-1) {
			m.fn97(i32(1291936), i32(43), v4+i32(79), i32(1079968), i32(1079984))
			panic("unreachable")
		}
		t21 := int32(load32(m.memory[int64(uint32(v4))+36:]))
		v9 = t21
		t22 := m.fn96(v7)
		v6 = t22
		t23 := m.fn96(v7)
		m.fn247(v4+i32(8), i32(4), t23)
		t24 := int32(load32(m.memory[int64(uint32(v4))+8:]))
		v12 = t24
		if v12 == 0 {
			m.fn85(i32(4), v6)
			panic("unreachable")
		}
		store64(m.memory[uint32(v12):], uint64(i64(0x100000001)))
		if v7 == 0 {
			goto l9
		}
		memory_copy(m.memory, uint32(v12+i32(8)), uint32(v9), uint32(v7))
	l9:
		t25 := int32(load32(m.memory[int64(uint32(v4))+32:]))
		m.fn16(t25, v9)
		m.fn51(v4+i32(44), v2, v3)
		t26 := int32(load32(m.memory[uint32(v12):]))
		t27 := v12
		v3 = t26 + i32(1)
		store32(m.memory[uint32(t27):], uint32(v3))
		if v3 == 0 {
			goto l5
		}
		t28 := int64(load64(m.memory[int64(uint32(v1))+16:]))
		t29 := int64(load64(m.memory[int64(uint32(v1))+24:]))
		t30 := int32(load32(m.memory[int64(uint32(v4))+48:]))
		t31 := int32(load32(m.memory[int64(uint32(v4))+52:]))
		t32 := m.fn524(t28, t29, t30, t31)
		v5 = t32
		store32(m.memory[int64(uint32(v4))+72:], uint32(v4+i32(44)))
		{
			t33 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			if t33 != 0 {
				goto l10
			}
			_ = m.fn654(v1, v1+i32(16))
		}
	l10:
		store32(m.memory[int64(uint32(v4))+32:], uint32(v4+i32(72)))
		store32(m.memory[int64(uint32(v4))+36:], uint32(v1))
		t35 := int32(load32(m.memory[uint32(v1):]))
		t36 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		m.fn69(v4, t35, t36, v5, v4+i32(32), i32(157))
		t37 := int32(load32(m.memory[uint32(v1):]))
		v3 = t37
		t38 := int32(load32(m.memory[int64(uint32(v4))+4:]))
		v2 = t38
		{
			{
				t39 := int32(load32(m.memory[uint32(v4):]))
				if t39 != i32(1) {
					goto l11
				}
				v9 = v3 + v2
				t40 := int32(m.memory[uint32(v9)])
				v6 = t40
				t41 := int32(load32(m.memory[int64(uint32(v4))+52:]))
				v10 = t41
				t42 := int64(load64(m.memory[int64(uint32(v4))+44:]))
				v8 = t42
				t43 := v9
				v13 = int32(uint32(int32(v5)) >> 25)
				m.memory[uint32(t43)] = byte(v13)
				t44 := int32(load32(m.memory[int64(uint32(v1))+4:]))
				m.memory[uint32(v3+t44&(v2+i32(-8))+i32(8))] = byte(v13)
				t45 := int32(load32(m.memory[int64(uint32(v1))+12:]))
				store32(m.memory[int64(uint32(v1))+12:], uint32(t45+i32(1)))
				t46 := int32(load32(m.memory[int64(uint32(v1))+8:]))
				store32(m.memory[int64(uint32(v1))+8:], uint32(t46-v6&i32(1)))
				v1 = i32(0)
				v3 = v3 + (i32(0)-v2)*i32(20)
				v2 = v3 + i32(-20)
				store64(m.memory[uint32(v2):], uint64(v8))
				store32(m.memory[int64(uint32(v2))+8:], uint32(v10))
				store32(m.memory[uint32(v3+i32(-4)):], uint32(v7))
				store32(m.memory[uint32(v3+i32(-8)):], uint32(v12))
				goto l12
			}
		l11:
			v1 = v3 + (i32(0)-v2)*i32(20)
			v2 = v1 + i32(-4)
			t47 := int32(load32(m.memory[uint32(v2):]))
			v3 = t47
			store32(m.memory[uint32(v2):], uint32(v7))
			v2 = v1 + i32(-8)
			t48 := int32(load32(m.memory[uint32(v2):]))
			v1 = t48
			store32(m.memory[uint32(v2):], uint32(v12))
			t49 := int32(load32(m.memory[int64(uint32(v4))+44:]))
			t50 := int32(load32(m.memory[int64(uint32(v4))+48:]))
			m.fn16(t49, t50)
		}
	l12:
		store32(m.memory[int64(uint32(v4))+36:], uint32(v3))
		store32(m.memory[int64(uint32(v4))+32:], uint32(v1))
		m.fn1051(v4 + i32(32))
	}
l4:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v7))
	store32(m.memory[uint32(v0):], uint32(v12))
	m.g0 = v4 + i32(80)
	return
l5:
	panic("unreachable")
}
func (m *Module) fn1084(v0, v1, v2, v3 int32) {
	var v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14 int32
	t0 := m.g0
	v4 = t0 - i32(48)
	m.g0 = v4
	{
		{
			{
				t1 := int32(load32(m.memory[uint32(v1):]))
				switch t1 {
				default:
					goto l0
				case 1:
					t2 := int32(load32(m.memory[int64(uint32(v1))+4:]))
					v5 = t2
					goto l4
				case 2:
					t3 := int32(load32(m.memory[int64(uint32(v1))+4:]))
					v6 = t3
					goto l5
				case 3:
					t4 := int32(load32(m.memory[int64(uint32(v1))+4:]))
					t5 := v4 + i32(24)
					v6 = t4
					m.fn148(t5, v6, v2, v3, i32(1282464))
					v7 = i32(0) - v6
					t6 := int32(load32(m.memory[int64(uint32(v4))+28:]))
					v6 = t6 + i32(-1)
					t7 := int32(load32(m.memory[int64(uint32(v4))+24:]))
					v8 = t7
				l11:
					if v6 == i32(-1) {
						goto l0
					}
					{
						t8 := int32(m.memory[uint32(v8)])
						v9 = t8
						if uint32(v9+i32(-9)) < uint32(i32(2)) {
							goto l6
						}
						switch v9 + i32(-32) {
						case 0:
							goto l6
						case 1:
							goto l7
						default:
							if v9 == i32(13) {
								goto l6
							}
							if v9 != i32(39) {
								goto l7
							}
							fallthrough
						case 2:
							v5 = i32(0)
						l10:
							{
								if v6 == 0 {
									goto l0
								}
								v10 = v8 + v5
								v6 = v6 + i32(-1)
								v5 = v5 + i32(1)
								t9 := int32(m.memory[uint32(v10+i32(1))])
								if t9 != v9 {
									goto l10
								}
							}
							v5 = v5 - v7
							goto l4
						}
					l7:
						v6 = i32(0) - v7
						goto l5
					}
				l6:
					v8 = v8 + i32(1)
					v7 = v7 + i32(-1)
					v6 = v6 + i32(-1)
					goto l11
				}
			}
		l5:
			m.fn148(v4+i32(16), v6, v2, v3, i32(1282448))
			v5 = v6 + i32(-1)
			t10 := int32(load32(m.memory[int64(uint32(v4))+20:]))
			v6 = t10
			t11 := int32(load32(m.memory[int64(uint32(v4))+16:]))
			v8 = t11
		l12:
			{
				if v6 == 0 {
					goto l0
				}
				t12 := int32(m.memory[uint32(v8)])
				v7 = t12
				v6 = v6 + i32(-1)
				v5 = v5 + i32(1)
				v8 = v8 + i32(1)
				v7 = v7 + i32(-9)
				if uint32(v7) > uint32(i32(23)) {
					goto l12
				}
				if i32_shl(i32(1), v7)&i32(8388627) == 0 {
					goto l12
				}
			}
		}
	l4:
		m.fn148(v4+i32(8), v5, v2, v3, i32(1282496))
		t13 := int32(load32(m.memory[int64(uint32(v4))+8:]))
		v9 = t13
		t14 := int32(load32(m.memory[int64(uint32(v4))+12:]))
		t15 := v9
		v7 = t14
		v11 = t15 + v7
		v6 = i32(0)
	l15:
		{
			if v7 == v6 {
				goto l13
			}
			t16 := int32(m.memory[uint32(v9+v6)])
			v8 = t16 + i32(-9)
			if uint32(v8) > uint32(i32(23)) {
				goto l14
			}
			if i32_shl(i32(1), v8)&i32(8388627) == 0 {
				goto l14
			}
			v6 = v6 + i32(1)
			goto l15
		}
	}
l0:
	store32(m.memory[uint32(v0):], uint32(i32(-2)))
	goto l16
l14:
	v12 = v5 + v6
	v13 = v7 + i32(-1)
	v14 = v9 + i32(1)
l24:
	{
		if v13 == v6 {
			store32(m.memory[uint32(v1):], uint32(i32(0)))
			m.fn1772(v0, v1, v2, v3, v12, v3, v3)
			goto l16
		}
		v8 = v14 + v6
		v6 = v6 + i32(1)
		t17 := int32(m.memory[uint32(v8)])
		v8 = t17
		v10 = v8 + i32(-9)
		if uint32(v10) > uint32(i32(23)) {
			goto l18
		}
		if i32_shl(i32(1), v10)&i32(8388627) == 0 {
			goto l18
		}
		v8 = v14 + v6
		v14 = v5 + v6
		v9 = v14 + i32(1)
		v6 = v6 ^ i32(-1) + v7
	l21:
		{
			if v6 == 0 {
				store32(m.memory[uint32(v1):], uint32(i32(0)))
				m.fn1772(v0, v1, v2, v3, v12, v14, v3)
				goto l16
			}
			t18 := int32(m.memory[uint32(v8)])
			v5 = t18
			v7 = v5 + i32(-9)
			if uint32(v7) > uint32(i32(23)) {
				goto l20
			}
			if i32_shl(i32(1), v7)&i32(8388627) == 0 {
				goto l20
			}
			v8 = v8 + i32(1)
			v9 = v9 + i32(1)
			v6 = v6 + i32(-1)
			goto l21
		}
	}
l20:
	if v5 != i32(61) {
		store32(m.memory[int64(uint32(v1))+4:], uint32(v9))
		store32(m.memory[uint32(v1):], uint32(i32(1)))
		m.fn1772(v0, v1, v2, v3, v12, v14, v9)
		goto l16
	}
	v10 = v8 + i32(1)
	v5 = v9 + i32(1)
	goto l23
l18:
	if v8 != i32(61) {
		goto l24
	}
	v10 = v9 + v6 + i32(1)
	v14 = v5 + v6
	v5 = v14 + i32(1)
	v9 = v14
l23:
	m.fn1773(v4+i32(36), v1, v2, v3, v12, v14)
	{
		t19 := int32(m.memory[int64(uint32(v4))+36])
		if t19 == i32(255) {
			t22 := int32(load32(m.memory[int64(uint32(v4))+44:]))
			v2 = t22
			t23 := int32(load32(m.memory[int64(uint32(v4))+40:]))
			v14 = t23
			v6 = i32(0)
		l37:
			{
				{
					v8 = v10 + v6
					if v8 == v11 {
						store32(m.memory[int64(uint32(v0))+8:], uint32(v3))
						m.memory[int64(uint32(v0))+4] = byte(i32(1))
						store32(m.memory[uint32(v0):], uint32(i32(-1)))
						store32(m.memory[uint32(v1):], uint32(i32(0)))
						goto l16
					}
					t24 := int32(m.memory[uint32(v8)])
					v7 = t24
					if uint32(v7+i32(-9)) < uint32(i32(2)) {
						goto l27
					}
					v9 = v5 + v6
					switch v7 + i32(-32) {
					case 0:
						goto l27
					case 1:
						goto l28
					default:
						if v7 == i32(13) {
							goto l27
						}
						if v7 != i32(39) {
							goto l28
						}
						fallthrough
					case 2:
						v10 = v9 + i32(1)
						v6 = i32(0)
					l32:
						{
							v5 = v8 + v6 + i32(1)
							if v5 == v11 {
								store32(m.memory[int64(uint32(v0))+8:], uint32(v3))
								m.memory[int64(uint32(v0))+5] = byte(v7)
								m.memory[int64(uint32(v0))+4] = byte(i32(3))
								store32(m.memory[uint32(v0):], uint32(i32(-1)))
								store32(m.memory[uint32(v1):], uint32(i32(0)))
								goto l16
							}
							v6 = v6 + i32(1)
							t25 := int32(m.memory[uint32(v5)])
							if t25 != v7 {
								goto l32
							}
						}
						v6 = v9 + v6
						v8 = v6 + i32(1)
						if v7 != i32(34) {
							store32(m.memory[int64(uint32(v1))+4:], uint32(v8))
							store32(m.memory[uint32(v1):], uint32(i32(1)))
							store32(m.memory[int64(uint32(v0))+16:], uint32(v6))
							store32(m.memory[int64(uint32(v0))+12:], uint32(v10))
							store32(m.memory[int64(uint32(v0))+8:], uint32(v2))
							store32(m.memory[int64(uint32(v0))+4:], uint32(v14))
							store32(m.memory[uint32(v0):], uint32(i32(1)))
							goto l16
						}
						store32(m.memory[int64(uint32(v1))+4:], uint32(v8))
						store32(m.memory[uint32(v1):], uint32(i32(1)))
						store32(m.memory[int64(uint32(v0))+16:], uint32(v6))
						store32(m.memory[int64(uint32(v0))+12:], uint32(v10))
						store32(m.memory[int64(uint32(v0))+8:], uint32(v2))
						store32(m.memory[int64(uint32(v0))+4:], uint32(v14))
						store32(m.memory[uint32(v0):], uint32(i32(0)))
						goto l16
					}
				}
			l28:
				t26 := int32(m.memory[int64(uint32(v1))+36])
				if t26 != i32(1) {
					store32(m.memory[int64(uint32(v1))+4:], uint32(v9))
					store32(m.memory[uint32(v1):], uint32(i32(2)))
					store32(m.memory[int64(uint32(v0))+8:], uint32(v9))
					m.memory[int64(uint32(v0))+4] = byte(i32(2))
					store32(m.memory[uint32(v0):], uint32(i32(-1)))
					goto l16
				}
				v6 = i32(0)
			l36:
				{
					v7 = v8 + v6 + i32(1)
					if v7 == v11 {
						goto l35
					}
					v6 = v6 + i32(1)
					t27 := int32(m.memory[uint32(v7)])
					v7 = t27 + i32(-9)
					if uint32(v7) > uint32(i32(23)) {
						goto l36
					}
					if i32_shl(i32(1), v7)&i32(8388627) == 0 {
						goto l36
					}
				}
				v3 = v9 + v6
			l35:
				store32(m.memory[int64(uint32(v1))+4:], uint32(v3))
				store32(m.memory[uint32(v1):], uint32(i32(1)))
				store32(m.memory[int64(uint32(v0))+16:], uint32(v3))
				store32(m.memory[int64(uint32(v0))+12:], uint32(v9))
				store32(m.memory[int64(uint32(v0))+8:], uint32(v2))
				store32(m.memory[int64(uint32(v0))+4:], uint32(v14))
				store32(m.memory[uint32(v0):], uint32(i32(2)))
				goto l16
			}
		l27:
			v6 = v6 + i32(1)
			goto l37
		}
		t20 := int32(load32(m.memory[int64(uint32(v4))+44:]))
		store32(m.memory[int64(uint32(v0))+12:], uint32(t20))
		t21 := int64(load64(m.memory[int64(uint32(v4))+36:]))
		store64(m.memory[int64(uint32(v0))+4:], uint64(t21))
		store32(m.memory[int64(uint32(v1))+4:], uint32(v9))
		store32(m.memory[uint32(v1):], uint32(i32(3)))
		store32(m.memory[uint32(v0):], uint32(i32(-1)))
		goto l16
	}
l13:
	store32(m.memory[uint32(v0):], uint32(i32(-2)))
	store32(m.memory[uint32(v1):], uint32(i32(0)))
l16:
	m.g0 = v4 + i32(48)
}
func (m *Module) fn1085(v0, v1, v2, v3, v4 int32) {
	if uint32(v4) < uint32(v3) {
		goto l0
	}
	if uint32(v4) <= uint32(v2) {
		goto l1
	}
l0:
	m.fn151(v3, v4, v2, i32(1281696))
	panic("unreachable")
l1:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v4-v3))
	store32(m.memory[uint32(v0):], uint32(v1+v3))
}
func (m *Module) fn1086(v0, v1 int32) {
	m.fn1080(v0, v1, i32(4), i32(8))
}
func (m *Module) fn1087(v0, v1 int32) {
	var v2 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	{
		if v1 == 0 {
			goto l0
		}
		m.fn1616(v2+i32(4), i32(8), i32(8), v1+i32(1))
		t1 := int32(load32(m.memory[int64(uint32(v2))+12:]))
		t2 := int32(load32(m.memory[int64(uint32(v2))+4:]))
		t3 := int32(load32(m.memory[int64(uint32(v2))+8:]))
		m.fn40(v0-t1, t2, t3)
	}
l0:
	m.g0 = v2 + i32(16)
}
func (m *Module) fn1088(v0, v1 int32) {
	var v2, v3, v4, v5, v6 int32
	t0 := int32(load32(m.memory[uint32(v1):]))
	v2 = t0
	t1 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	v3 = t1
	v4 = i32(0)
	v5 = i32(0)
l4:
	v6 = v2 + v5
	if v6 != v3 {
		goto l0
	}
	goto l1
l0:
	store32(m.memory[uint32(v1):], uint32(v6+i32(1)))
	{
		t2 := int32(m.memory[uint32(v6)])
		v6 = t2 + i32(-9)
		if uint32(v6) > uint32(i32(29)) {
			goto l2
		}
		if i32_shl(i32(1), v6)&i32(0x20000013) != 0 {
			goto l3
		}
	}
l2:
	v5 = v5 + i32(1)
	goto l4
l3:
	v4 = i32(1)
l1:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
	store32(m.memory[uint32(v0):], uint32(v4))
}
func (m *Module) fn1089(v0, v1, v2, v3, v4, v5, v6, v7 int32) {
	var v8, v9, v10, v11 int32
	t0 := m.g0
	v8 = t0 - i32(64)
	m.g0 = v8
	{
		if v7 != 0 {
			goto l0
		}
		store32(m.memory[uint32(v0):], uint32(i32(-0x7ffffffe)))
		goto l1
	l0:
		m.fn1582(v8+i32(24), v3, v4, v5, v6, i32(1281548))
		t1 := int32(load32(m.memory[int64(uint32(v8))+24:]))
		t2 := int32(load32(m.memory[int64(uint32(v8))+28:]))
		m.fn1583(v1, t1, t2)
		if uint32(v6) < uint32(v4) {
			goto l2
		}
		m.fn158(v6, v4, i32(1281564))
		panic("unreachable")
	l2:
		{
			t3 := int32(m.memory[uint32(v3+v6)])
			v7 = t3
			if v7 == i32(9) {
				m.fn1752(v1, i32(32))
				store32(m.memory[uint32(v0):], uint32(i32(-1)))
				store32(m.memory[int64(uint32(v0))+4:], uint32(v6+i32(1)))
				goto l1
			}
			{
				if v7 == i32(38) {
					v11 = v6 + i32(1)
					t8 := int32(load32(m.memory[uint32(v2):]))
					v10 = t8
					t9 := int32(load32(m.memory[int64(uint32(v2))+4:]))
					v9 = t9
					v7 = i32(0)
				l9:
					{
						v5 = v10 + v7
						if v5 == v9 {
							store32(m.memory[int64(uint32(v0))+8:], uint32(v4))
							store32(m.memory[int64(uint32(v0))+4:], uint32(v6))
							store32(m.memory[uint32(v0):], uint32(i32(-0x80000000)))
							goto l1
						}
						store32(m.memory[uint32(v2):], uint32(v5+i32(1)))
						t10 := int32(m.memory[uint32(v5)])
						if t10 == i32(59) {
							t11 := v8 + i32(16)
							t12 := v3
							t13 := v4
							t14 := v11
							v10 = v7 + v11
							m.fn1582(t11, t12, t13, t14, v10, i32(1281580))
							t15 := int32(load32(m.memory[int64(uint32(v8))+16:]))
							t16 := v8 + i32(8)
							v6 = t15
							t17 := int32(load32(m.memory[int64(uint32(v8))+20:]))
							t18 := v6
							v7 = t17
							m.fn13(t16, t18, v7, i32(35))
							{
								t19 := int32(load32(m.memory[int64(uint32(v8))+8:]))
								v5 = t19
								if v5 == 0 {
									{
										if v7 != i32(3) {
											goto l12
										}
										t23 := m.fn1851(v6, i32(1281596), i32(3))
										if t23 == 0 {
											m.fn1752(v1, i32(38))
											goto l15
										}
									}
								l12:
									t24 := m.fn275(v6, v7)
									v5 = t24
									if v5 == 0 {
										m.fn884(v0, v6, v7)
										store32(m.memory[int64(uint32(v0))+16:], uint32(v10))
										store32(m.memory[int64(uint32(v0))+12:], uint32(v11))
										goto l1
									}
									store32(m.memory[int64(uint32(v8))+60:], uint32(v5+i32(1)))
									store32(m.memory[int64(uint32(v8))+56:], uint32(v5))
									m.fn1090(v8+i32(32), v1, v8+i32(56), v5, i32(1), i32(0), i32(0))
									t25 := int32(load32(m.memory[int64(uint32(v8))+32:]))
									v7 = t25
									if v7 == i32(-1) {
										goto l15
									}
									t26 := int64(load64(m.memory[int64(uint32(v8))+44:]))
									store64(m.memory[int64(uint32(v0))+12:], uint64(t26))
									t27 := int64(load64(m.memory[int64(uint32(v8))+36:]))
									store64(m.memory[int64(uint32(v0))+4:], uint64(t27))
									store32(m.memory[uint32(v0):], uint32(v7))
									goto l1
								}
								t20 := int32(load32(m.memory[int64(uint32(v8))+12:]))
								m.fn276(v8+i32(32), v5, t20)
								t21 := int32(m.memory[int64(uint32(v8))+32])
								if t21 == i32(255) {
									t28 := int32(load32(m.memory[int64(uint32(v8))+36:]))
									v7 = t28
									store32(m.memory[int64(uint32(v8))+32:], uint32(i32(0)))
									m.fn522(v8, v7, v8+i32(32))
									t29 := int32(load32(m.memory[uint32(v8):]))
									t30 := int32(load32(m.memory[int64(uint32(v8))+4:]))
									m.fn1583(v1, t29, t30)
									goto l15
								}
								t22 := int64(load64(m.memory[int64(uint32(v8))+32:]))
								store64(m.memory[int64(uint32(v0))+4:], uint64(t22))
								store32(m.memory[uint32(v0):], uint32(i32(-0x7fffffff)))
								goto l1
							}
						}
						v7 = v7 + i32(1)
						goto l9
					}
				}
				t4 := m.fn1751(v1, v3, v4, v6, i32(32))
				v9 = t4
				v6 = v9 + (v6 ^ i32(-1))
				t5 := int32(load32(m.memory[int64(uint32(v2))+4:]))
				v10 = t5
				t6 := int32(load32(m.memory[uint32(v2):]))
				v5 = t6
				v7 = i32(0)
			l6:
				{
					if v6 == v7 {
						store32(m.memory[uint32(v0):], uint32(i32(-1)))
						store32(m.memory[int64(uint32(v0))+4:], uint32(v9))
						goto l1
					}
					v7 = v7 + i32(1)
					if v5 == v10 {
						goto l6
					}
					t7 := v2
					v5 = v5 + i32(1)
					store32(m.memory[uint32(t7):], uint32(v5))
					goto l6
				}
			}
		}
	l15:
		store32(m.memory[uint32(v0):], uint32(i32(-1)))
		store32(m.memory[int64(uint32(v0))+4:], uint32(v10+i32(1)))
	}
l1:
	m.g0 = v8 + i32(64)
}
func (m *Module) fn1090(v0, v1, v2, v3, v4, v5, v6 int32) {
	var v7, v8 int32
	t0 := m.g0
	v7 = t0 - i32(48)
	m.g0 = v7
	{
	l1:
		{
			m.fn1088(v7+i32(16), v2)
			t1 := int32(load32(m.memory[int64(uint32(v7))+16:]))
			if t1 != i32(1) {
				goto l0
			}
			t2 := int32(load32(m.memory[int64(uint32(v7))+20:]))
			m.fn1089(v7+i32(28), v1, v2, v3, v4, v5, t2+v5, v6)
			t3 := int32(load32(m.memory[int64(uint32(v7))+32:]))
			v5 = t3
			t4 := int32(load32(m.memory[int64(uint32(v7))+28:]))
			v8 = t4
			if v8 == i32(-1) {
				goto l1
			}
		}
		t5 := int32(load32(m.memory[int64(uint32(v7))+44:]))
		store32(m.memory[int64(uint32(v0))+16:], uint32(t5))
		t6 := int64(load64(m.memory[int64(uint32(v7))+36:]))
		store64(m.memory[int64(uint32(v0))+8:], uint64(t6))
		store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
		store32(m.memory[uint32(v0):], uint32(v8))
		goto l2
	}
l0:
	m.fn1584(v7+i32(8), v3, v4, v5)
	{
		t7 := int32(load32(m.memory[int64(uint32(v7))+8:]))
		v5 = t7
		if v5 == 0 {
			goto l3
		}
		t8 := int32(load32(m.memory[int64(uint32(v7))+12:]))
		m.fn1583(v1, v5, t8)
	}
l3:
	store32(m.memory[uint32(v0):], uint32(i32(-1)))
l2:
	m.g0 = v7 + i32(48)
}
func (m *Module) fn1091(v0, v1 int32) {
	m.fn1080(v0, v1, i32(1), i32(1))
}
func (m *Module) fn1092(v0, v1 int32) {
	var v2 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		v2 = t0
		t1 := int32(load32(m.memory[uint32(v0):]))
		if v2 != t1 {
			goto l0
		}
		m.fn1075(v0)
	}
l0:
	t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	memory_copy(m.memory, uint32(t2+v2*i32(44)), uint32(v1), uint32(i32(44)))
	store32(m.memory[int64(uint32(v0))+8:], uint32(v2+i32(1)))
}
func (m *Module) fn1093(v0, v1 int32) int32 {
	var v2 int32
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(t0):]))
	v2 = t1
	t2 := int32(load32(m.memory[uint32(v2+i32(4)):]))
	t3 := int32(load32(m.memory[uint32(v2+i32(8)):]))
	t4 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t5 := int32(load32(m.memory[uint32(t4):]))
	v0 = t5 + (i32(0)-v1)*i32(20)
	t6 := int32(load32(m.memory[uint32(v0+i32(-16)):]))
	t7 := int32(load32(m.memory[uint32(v0+i32(-12)):]))
	t8 := m.fn191(t2, t3, t6, t7)
	return t8
}
func (m *Module) fn1094(v0, v1 int32) {
l1:
	if v1 == 0 {
		return
	}
	m.memory[uint32(v0)] = byte(i32(0))
	v1 = v1 + i32(-1)
	v0 = v0 + i32(1)
	goto l1
}
func (m *Module) fn1095(v0, v1 int32) int32 {
	var v2 int32
	t0 := m.g0
	v2 = t0 - i32(32)
	m.g0 = v2
	{
		t1 := int32(m.memory[uint32(v0)])
		switch t1 {
		default:
			t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			store32(m.memory[uint32(v2):], uint32(t2))
			t3 := m.fn4(i32(20))
			v0 = t3
			if v0 == 0 {
				m.fn2(i32(1), i32(20))
				panic("unreachable")
			}
			t4 := int32(load32(m.memory[int64(uint32(i32(0)))+1284656:]))
			store32(m.memory[int64(uint32(v0))+16:], uint32(t4))
			t5 := int64(load64(m.memory[int64(uint32(i32(0)))+1284648:]))
			store64(m.memory[int64(uint32(v0))+8:], uint64(t5))
			t6 := int64(load64(m.memory[int64(uint32(i32(0)))+1284640:]))
			store64(m.memory[uint32(v0):], uint64(t6))
			store32(m.memory[int64(uint32(v2))+12:], uint32(i32(20)))
			store32(m.memory[int64(uint32(v2))+8:], uint32(v0))
			store32(m.memory[int64(uint32(v2))+4:], uint32(i32(20)))
			store64(m.memory[int64(uint32(v2))+24:], uint64(int64(uint32(i32(29)))<<32|int64(uint32(v2))))
			store64(m.memory[int64(uint32(v2))+16:], uint64(int64(uint32(i32(158)))<<32|int64(uint32(v2+i32(4)))))
			t7 := int32(load32(m.memory[uint32(v1):]))
			t8 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t9 := m.fn100(t7, t8, i32(1069318), v2+i32(16))
			v0 = t9
			t10 := int32(load32(m.memory[int64(uint32(v2))+4:]))
			v1 = t10
			if v1 == 0 {
				goto l5
			}
			t11 := int32(load32(m.memory[int64(uint32(v2))+8:]))
			m.fn10(t11, v1, i32(1))
			goto l5
		case 1:
			t12 := int32(load32(m.memory[uint32(v1):]))
			t13 := int32(m.memory[int64(uint32(v0))+1])
			v0 = t13 << 2
			t14 := int32(load32(m.memory[int64(uint32(v0))+1301988:]))
			t15 := int32(load32(m.memory[int64(uint32(v0))+1301820:]))
			t16 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t17 := int32(load32(m.memory[int64(uint32(t16))+12:]))
			t18 := m.t0[uint(t17)].(func(int32, int32, int32) int32)(t12, t14, t15)
			v0 = t18
			goto l5
		case 2:
			t19 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t20 := v1
			v0 = t19
			t21 := int32(load32(m.memory[uint32(v0):]))
			t22 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t23 := m.fn110(t20, t21, t22)
			v0 = t23
			goto l5
		case 3:
			t24 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			v0 = t24
			t25 := int32(load32(m.memory[uint32(v0):]))
			t26 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t27 := int32(load32(m.memory[int64(uint32(t26))+16:]))
			t28 := m.t0[uint(t27)].(func(int32, int32) int32)(t25, v1)
			v0 = t28
		}
	}
l5:
	m.g0 = v2 + i32(32)
	return v0
}
func (m *Module) fn1096(v0, v1 int32) int32 {
	var v2 int32
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(t0):]))
	v2 = t1
	t2 := int32(load32(m.memory[uint32(v2+i32(4)):]))
	t3 := int32(load32(m.memory[uint32(v2+i32(8)):]))
	t4 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t5 := int32(load32(m.memory[uint32(t4):]))
	v0 = t5 + (i32(0)-v1)*i32(20)
	t6 := int32(load32(m.memory[uint32(v0+i32(-16)):]))
	t7 := int32(load32(m.memory[uint32(v0+i32(-12)):]))
	t8 := m.fn545(t2, t3, t6, t7)
	return t8
}
func (m *Module) fn1097(v0, v1, v2, v3, v4, v5 int32) int32 {
	var v6 int32
	t0 := m.g0
	v6 = t0 - i32(16)
	m.g0 = v6
	m.fn868(v6+i32(4), v0, v1)
l2:
	{
		{
			t1 := m.fn866(v6 + i32(4))
			v1 = t1
			if v1 != 0 {
				goto l0
			}
			v1 = i32(0)
			goto l1
		}
	l0:
		t2 := int32(load32(m.memory[uint32(v1):]))
		if t2 == i32(-1) {
			goto l2
		}
		t3 := m.fn847(v1, v2, v3, v4, v5)
		if t3 == 0 {
			goto l2
		}
	}
l1:
	t4 := int32(load32(m.memory[int64(uint32(v6))+4:]))
	t5 := int32(load32(m.memory[int64(uint32(v6))+8:]))
	m.fn44(t4, t5)
	m.g0 = v6 + i32(16)
	return v1
}
func (m *Module) fn1098(v0, v1 int32) {
	var v2, v3, v4 int32
	var v5 int64
	var v6 int32
	var v7 float64
	t0 := m.g0
	v2 = t0 - i32(112)
	m.g0 = v2
	v3 = v1 + i32(288)
	{
		{
		l6:
			{
				m.fn505(v2+i32(80), v1)
				{
					{
						t1 := int32(m.memory[int64(uint32(v2))+80])
						if t1 != i32(255) {
							goto l0
						}
						t2 := int32(load16(m.memory[int64(uint32(v2))+82:]))
						v4 = t2
						goto l1
					}
				l0:
					t3 := int64(load64(m.memory[int64(uint32(v2))+80:]))
					v5 = t3
					if v5&i64(255) != i64(255) {
						store32(m.memory[int64(uint32(v0))+4:], uint32(i32(-0x7ffffff1)))
						store64(m.memory[int64(uint32(v0))+8:], uint64(v5))
						m.memory[uint32(v0)] = byte(i32(254))
						goto l4
					}
					v4 = int32(int64(uint64(v5) >> 16))
				}
			l1:
				store16(m.memory[int64(uint32(v1))+304:], uint16(v4))
				m.fn507(v2+i32(80), v1, v3)
				{
					t4 := int32(m.memory[int64(uint32(v2))+80])
					if t4 == i32(255) {
						goto l3
					}
					t5 := int64(load64(m.memory[int64(uint32(v2))+80:]))
					v5 = t5
					if v5&i64(255) == i64(255) {
						goto l3
					}
					store32(m.memory[int64(uint32(v0))+4:], uint32(i32(-0x7ffffff1)))
					store64(m.memory[int64(uint32(v0))+8:], uint64(v5))
					m.memory[uint32(v0)] = byte(i32(254))
					goto l4
				}
			l3:
				{
					t6 := int32(load16(m.memory[int64(uint32(v1))+304:]))
					v4 = t6
					switch v4 {
					case 0:
						goto l5
					case 1:
						goto l6
					case 6, 8:
						t63 := int32(load32(m.memory[int64(uint32(v1))+292:]))
						t64 := int32(load32(m.memory[int64(uint32(v1))+296:]))
						m.fn148(v2+i32(24), i32(8), t63, t64, i32(1075436))
						t65 := int32(load32(m.memory[int64(uint32(v2))+24:]))
						t66 := int32(load32(m.memory[int64(uint32(v2))+28:]))
						m.fn946(v2+i32(80), t65, t66, v2+i32(108))
						t67 := int64(load64(m.memory[int64(uint32(v2))+84:]))
						store64(m.memory[int64(uint32(v2))+64:], uint64(t67))
						t68 := int32(load32(m.memory[int64(uint32(v2))+92:]))
						store32(m.memory[int64(uint32(v2))+72:], uint32(t68))
						{
							t69 := int32(load32(m.memory[int64(uint32(v2))+80:]))
							v4 = t69
							if v4 == i32(-1) {
								m.fn490(v2+i32(40)|i32(4), v2+i32(64))
								m.memory[int64(uint32(v2))+40] = byte(i32(2))
								goto l16
							}
							t70 := int64(load64(m.memory[int64(uint32(v2))+96:]))
							v5 = t70
							t71 := int32(load32(m.memory[int64(uint32(v2))+72:]))
							store32(m.memory[int64(uint32(v0))+16:], uint32(t71))
							t72 := int64(load64(m.memory[int64(uint32(v2))+64:]))
							store64(m.memory[int64(uint32(v0))+8:], uint64(t72))
							store64(m.memory[int64(uint32(v0))+20:], uint64(v5))
							store32(m.memory[int64(uint32(v0))+4:], uint32(v4))
							m.memory[uint32(v0)] = byte(i32(254))
							goto l4
						}
					case 7:
						t73 := int32(load32(m.memory[int64(uint32(v1))+292:]))
						t74 := int32(load32(m.memory[int64(uint32(v1))+296:]))
						m.fn483(v2+i32(32), t73, t74, i32(8), i32(12), i32(1075452))
						t75 := int32(load32(m.memory[int64(uint32(v2))+32:]))
						t76 := int32(load32(m.memory[int64(uint32(v2))+36:]))
						t77 := m.fn371(t75, t76)
						v4 = t77
						t78 := int32(load32(m.memory[int64(uint32(v1))+252:]))
						t79 := v4
						v3 = t78
						if uint32(t79) < uint32(v3) {
							goto l31
						}
						m.fn158(v4, v3, i32(1075468))
						panic("unreachable")
					case 2:
						t7 := int32(load32(m.memory[int64(uint32(v1))+292:]))
						t8 := int32(load32(m.memory[int64(uint32(v1))+296:]))
						t9 := m.fn509(t7, t8, i32(8), i32(1075292))
						t10 := int32(m.memory[uint32(t9)])
						v4 = t10
						t11 := int32(load32(m.memory[int64(uint32(v1))+292:]))
						t12 := int32(load32(m.memory[int64(uint32(v1))+296:]))
						t13 := m.fn509(t11, t12, i32(8), i32(1075308))
						v3 = t13
						t14 := int32(load32(m.memory[int64(uint32(v1))+296:]))
						v6 = t14
						if uint32(v6) <= uint32(i32(8)) {
							m.fn158(i32(8), v6, i32(1075324))
							panic("unreachable")
						}
						v6 = v4 & i32(1)
						t15 := int32(m.memory[uint32(v3)])
						v4 = t15
						t16 := int32(load32(m.memory[int64(uint32(v1))+292:]))
						v3 = t16
						t17 := int32(m.memory[int64(uint32(v3))+8])
						m.memory[int64(uint32(v3))+8] = byte(t17 & i32(252))
						if v4&i32(2) != 0 {
							t52 := int32(load32(m.memory[int64(uint32(v1))+292:]))
							t53 := int32(load32(m.memory[int64(uint32(v1))+296:]))
							m.fn483(v2+i32(8), t52, t53, i32(8), i32(12), i32(1075372))
							t54 := int32(load32(m.memory[int64(uint32(v2))+12:]))
							v4 = t54
							if uint32(v4) <= uint32(i32(3)) {
								m.fn151(i32(0), i32(4), v4, i32(1099732))
								panic("unreachable")
							}
							t55 := int32(load32(m.memory[int64(uint32(v2))+8:]))
							t56 := int32(load32(m.memory[uint32(t55):]))
							v4 = t56 >> 2
							if v6 != 0 {
								t57 := int32(load32(m.memory[int64(uint32(v1))+240:]))
								t58 := int32(load32(m.memory[int64(uint32(v1))+244:]))
								t59 := int32(load32(m.memory[int64(uint32(v1))+292:]))
								t60 := int32(load32(m.memory[int64(uint32(v1))+296:]))
								t61 := m.fn1099(t57, t58, t59, t60)
								t62 := int32(m.memory[int64(uint32(v1))+306])
								m.fn1100(v2+i32(40), float64(float64(v4)/float64(100)), t61, t62)
								goto l16
							}
							m.memory[int64(uint32(v2))+40] = byte(i32(0))
							store64(m.memory[int64(uint32(v2))+48:], uint64(int64(v4)))
							goto l16
						}
						store64(m.memory[int64(uint32(v2))+80:], uint64(i64(0)))
						t18 := int32(load32(m.memory[int64(uint32(v1))+292:]))
						t19 := int32(load32(m.memory[int64(uint32(v1))+296:]))
						m.fn483(v2, t18, t19, i32(8), i32(12), i32(1075340))
						t20 := int32(load32(m.memory[uint32(v2):]))
						t21 := int32(load32(m.memory[int64(uint32(v2))+4:]))
						m.fn310(v2+i32(80)|i32(4), i32(4), t20, t21, i32(1075356))
						t22 := math.Float64frombits(load64(m.memory[int64(uint32(v2))+80:]))
						t23 := v2 + i32(40)
						v7 = t22
						p24 := v7
						if v6 != 0 {
							p24 = float64(v7 / float64(100))
						}
						t25 := int32(load32(m.memory[int64(uint32(v1))+240:]))
						t26 := int32(load32(m.memory[int64(uint32(v1))+244:]))
						t27 := int32(load32(m.memory[int64(uint32(v1))+292:]))
						t28 := int32(load32(m.memory[int64(uint32(v1))+296:]))
						t29 := m.fn1099(t25, t26, t27, t28)
						t30 := int32(m.memory[int64(uint32(v1))+306])
						m.fn1100(t23, p24, t29, t30)
						goto l16
					case 3:
						t31 := int32(load32(m.memory[int64(uint32(v1))+292:]))
						t32 := int32(load32(m.memory[int64(uint32(v1))+296:]))
						t33 := m.fn509(t31, t32, i32(8), i32(1075388))
						t34 := int32(m.memory[uint32(t33)])
						v4 = t34
						switch v4 + i32(-42) {
						case 0:
							v4 = i32(1)
							goto l21
						case 1:
							goto l18
						default:
							if v4 != 0 {
								if v4 == i32(7) {
									v4 = i32(0)
									goto l21
								}
								if v4 == i32(15) {
									v4 = i32(6)
									goto l21
								}
								if v4 == i32(23) {
									v4 = i32(5)
									goto l21
								}
								if v4 == i32(29) {
									v4 = i32(2)
									goto l21
								}
								if v4 == i32(36) {
									v4 = i32(4)
									goto l21
								}
								m.memory[int64(uint32(v0))+8] = byte(v4)
								store32(m.memory[int64(uint32(v0))+4:], uint32(i32(-0x7fffffe4)))
								m.memory[uint32(v0)] = byte(i32(254))
								goto l4
							}
							v4 = i32(3)
							goto l21
						}
					case 4, 10:
						t35 := int32(load32(m.memory[int64(uint32(v1))+292:]))
						t36 := int32(load32(m.memory[int64(uint32(v1))+296:]))
						t37 := m.fn509(t35, t36, i32(8), i32(1075404))
						t38 := int32(m.memory[uint32(t37)])
						v4 = t38
						m.memory[int64(uint32(v2))+40] = byte(i32(4))
						t39 := v2
						var p40 int32
						if v4 != i32(0) {
							p40 = 1
						}
						m.memory[int64(uint32(t39))+41] = byte(p40)
						goto l16
					case 5, 9:
						t41 := int32(load32(m.memory[int64(uint32(v1))+292:]))
						t42 := int32(load32(m.memory[int64(uint32(v1))+296:]))
						m.fn483(v2+i32(16), t41, t42, i32(8), i32(16), i32(1075420))
						t43 := int32(load32(m.memory[int64(uint32(v2))+20:]))
						v4 = t43
						if uint32(v4) <= uint32(i32(7)) {
							m.fn151(i32(0), i32(8), v4, i32(1099700))
							panic("unreachable")
						}
						t44 := int32(load32(m.memory[int64(uint32(v2))+16:]))
						t45 := math.Float64frombits(load64(m.memory[uint32(t44):]))
						t46 := int32(load32(m.memory[int64(uint32(v1))+240:]))
						t47 := int32(load32(m.memory[int64(uint32(v1))+244:]))
						t48 := int32(load32(m.memory[int64(uint32(v1))+292:]))
						t49 := int32(load32(m.memory[int64(uint32(v1))+296:]))
						t50 := m.fn1099(t46, t47, t48, t49)
						t51 := int32(m.memory[int64(uint32(v1))+306])
						m.fn1100(v2+i32(40), t45, t50, t51)
						goto l16
					default:
						if v4 != i32(146) {
							goto l6
						}
						m.memory[uint32(v0)] = byte(i32(255))
						goto l4
					}
				}
			l18:
				v4 = i32(7)
			l21:
				m.memory[int64(uint32(v2))+40] = byte(i32(8))
				m.memory[int64(uint32(v2))+41] = byte(v4)
				goto l16
			l5:
				t80 := int32(load32(m.memory[int64(uint32(v1))+292:]))
				t81 := int32(load32(m.memory[int64(uint32(v1))+296:]))
				t82 := m.fn371(t80, t81)
				t83 := v1
				v4 = t82
				store32(m.memory[int64(uint32(t83))+300:], uint32(v4))
				if uint32(v4) <= uint32(i32(0x100000)) {
					goto l6
				}
			}
			m.memory[uint32(v0)] = byte(i32(255))
			goto l4
		l31:
			m.memory[int64(uint32(v2))+40] = byte(i32(3))
			t84 := int32(load32(m.memory[int64(uint32(v1))+248:]))
			t85 := int64(load64(m.memory[int64(uint32(t84+v4*i32(12)))+4:]))
			store64(m.memory[int64(uint32(v2))+44:], uint64(t85))
		}
	l16:
		t86 := int32(load32(m.memory[int64(uint32(v1))+292:]))
		t87 := int32(load32(m.memory[int64(uint32(v1))+296:]))
		t88 := m.fn371(t86, t87)
		store32(m.memory[int64(uint32(v0))+28:], uint32(t88))
		t89 := int64(load64(m.memory[int64(uint32(v2))+40:]))
		store64(m.memory[uint32(v0):], uint64(t89))
		t90 := int64(load64(m.memory[int64(uint32(v2))+48:]))
		store64(m.memory[int64(uint32(v0))+8:], uint64(t90))
		t91 := int64(load64(m.memory[int64(uint32(v2))+56:]))
		store64(m.memory[int64(uint32(v0))+16:], uint64(t91))
		t92 := int32(load32(m.memory[int64(uint32(v1))+300:]))
		store32(m.memory[int64(uint32(v0))+24:], uint32(t92))
	}
l4:
	m.g0 = v2 + i32(112)
}
func (m *Module) fn1099(v0, v1, v2, v3 int32) int32 {
	if uint32(v3) < uint32(i32(5)) {
		m.fn158(i32(4), v3, i32(1098380))
		panic("unreachable")
	}
	if v3 == i32(5) {
		m.fn158(i32(5), i32(5), i32(1098396))
		panic("unreachable")
	}
	if uint32(v3) > uint32(i32(6)) {
		t0 := int32(m.memory[int64(uint32(v2))+5])
		t1 := int32(m.memory[int64(uint32(v2))+4])
		t2 := int32(m.memory[int64(uint32(v2))+6])
		t3 := v0
		v3 = t0<<8 | t1 | t2<<16
		p4 := i32(0)
		if uint32(v3) < uint32(v1) {
			p4 = t3 + v3
		}
		return p4
	}
	m.fn158(i32(6), i32(6), i32(1098412))
	panic("unreachable")
}
func (m *Module) fn1100(v0 int32, v1 float64, v2, v3 int32) {
	{
		if v2 == 0 {
			goto l0
		}
		t0 := int32(m.memory[uint32(v2)])
		switch t0 {
		case 1:
			m.memory[int64(uint32(v0))+17] = byte(v3)
			m.memory[int64(uint32(v0))+16] = byte(i32(0))
			goto l4
		case 2:
			goto l2
		default:
			goto l0
		}
	}
l0:
	store64(m.memory[int64(uint32(v0))+8:], math.Float64bits(v1))
	v2 = i32(1)
	goto l3
l2:
	m.memory[int64(uint32(v0))+17] = byte(v3)
	m.memory[int64(uint32(v0))+16] = byte(i32(1))
l4:
	store64(m.memory[int64(uint32(v0))+8:], math.Float64bits(v1))
	v2 = i32(5)
l3:
	m.memory[uint32(v0)] = byte(v2)
}
func (m *Module) fn1101(v0, v1, v2 int32) int32 {
	var v3 int32
	v3 = i32(0)
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		if t0 != v2 {
			goto l0
		}
		t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		t2 := m.fn1851(t1, v1, v2)
		var p3 int32
		if t2 == 0 {
			p3 = 1
		}
		v3 = p3
	}
l0:
	return v3
}
func (m *Module) fn1102(v0, v1 int32) int32 {
	var v2 int32
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(t0):]))
	v2 = t1
	t2 := int32(load32(m.memory[uint32(v2+i32(4)):]))
	t3 := int32(load32(m.memory[uint32(v2+i32(8)):]))
	t4 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t5 := int32(load32(m.memory[uint32(t4):]))
	v0 = t5 + (i32(0)-v1)*i32(416)
	t6 := int32(load32(m.memory[uint32(v0+i32(-412)):]))
	t7 := int32(load32(m.memory[uint32(v0+i32(-408)):]))
	t8 := m.fn545(t2, t3, t6, t7)
	return t8
}
func (m *Module) fn1103(v0, v1 int32) int32 {
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
	t8 := m.fn545(t2, t3, t6, t7)
	return t8
}
func (m *Module) fn1104(v0, v1 int32) int32 {
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
	t8 := m.fn544(t2, t3, t6, t7)
	return t8
}
func (m *Module) fn1105(v0, v1, v2 int32) {
	var v3 int32
	var v4 int64
	var v5, v6, v7, v8, v9 int32
	var v10 int64
	t0 := m.g0
	v3 = t0 - i32(32)
	m.g0 = v3
	t1 := int64(load64(m.memory[int64(uint32(v0))+16:]))
	t2 := int64(load64(m.memory[int64(uint32(v0))+24:]))
	t3 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	t4 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	t5 := m.fn540(t1, t2, t3, t4)
	v4 = t5
	store32(m.memory[int64(uint32(v3))+20:], uint32(v1))
	{
		t6 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		if t6 != 0 {
			goto l0
		}
		_ = m.fn666(v0, v0+i32(16))
	}
l0:
	store32(m.memory[int64(uint32(v3))+24:], uint32(v3+i32(20)))
	store32(m.memory[int64(uint32(v3))+28:], uint32(v0))
	t8 := int32(load32(m.memory[uint32(v0):]))
	t9 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	m.fn69(v3, t8, t9, v4, v3+i32(24), i32(159))
	t10 := int32(load32(m.memory[uint32(v0):]))
	v5 = t10
	t11 := int32(load32(m.memory[int64(uint32(v3))+4:]))
	v6 = t11
	{
		{
			t12 := int32(load32(m.memory[uint32(v3):]))
			if t12 != i32(1) {
				goto l1
			}
			v7 = v5 + v6
			t13 := int32(m.memory[uint32(v7)])
			v8 = t13
			t14 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			v9 = t14
			t15 := int64(load64(m.memory[uint32(v1):]))
			v10 = t15
			t16 := v7
			v1 = int32(uint32(int32(v4)) >> 25)
			m.memory[uint32(t16)] = byte(v1)
			t17 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			m.memory[uint32(v5+t17&(v6+i32(-8))+i32(8))] = byte(v1)
			t18 := int32(load32(m.memory[int64(uint32(v0))+12:]))
			store32(m.memory[int64(uint32(v0))+12:], uint32(t18+i32(1)))
			t19 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			store32(m.memory[int64(uint32(v0))+8:], uint32(t19-v8&i32(1)))
			v0 = v5 - v6<<4
			v1 = v0 + i32(-16)
			store64(m.memory[uint32(v1):], uint64(v10))
			store32(m.memory[int64(uint32(v1))+8:], uint32(v9))
			store32(m.memory[uint32(v0+i32(-4)):], uint32(v2))
			goto l2
		}
	l1:
		store32(m.memory[uint32(v5-v6<<4+i32(-4)):], uint32(v2))
		t20 := int32(load32(m.memory[uint32(v1):]))
		t21 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		m.fn16(t20, t21)
	}
l2:
	m.g0 = v3 + i32(32)
}
func (m *Module) fn1106(v0, v1 int32) int32 {
	var v2 int32
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(t0):]))
	v2 = t1
	t2 := int32(load32(m.memory[uint32(v2+i32(4)):]))
	t3 := int32(load32(m.memory[uint32(v2+i32(8)):]))
	t4 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t5 := int32(load32(m.memory[uint32(t4):]))
	v0 = t5 - v1<<4
	t6 := int32(load32(m.memory[uint32(v0+i32(-12)):]))
	t7 := int32(load32(m.memory[uint32(v0+i32(-8)):]))
	t8 := m.fn544(t2, t3, t6, t7)
	return t8
}
func (m *Module) fn1107(v0, v1 int32) int32 {
	var v2 int32
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(t0):]))
	v2 = t1
	t2 := int32(load32(m.memory[uint32(v2+i32(4)):]))
	t3 := int32(load32(m.memory[uint32(v2+i32(8)):]))
	t4 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t5 := int32(load32(m.memory[uint32(t4):]))
	v0 = t5 + (i32(0)-v1)*i32(36)
	t6 := int32(load32(m.memory[uint32(v0+i32(-32)):]))
	t7 := int32(load32(m.memory[uint32(v0+i32(-28)):]))
	t8 := m.fn544(t2, t3, t6, t7)
	return t8
}
func (m *Module) fn1108(v0, v1 int32) int32 {
	var v2 int32
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(t0):]))
	v2 = t1
	t2 := int32(load32(m.memory[uint32(v2+i32(4)):]))
	t3 := int32(load32(m.memory[uint32(v2+i32(8)):]))
	t4 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t5 := int32(load32(m.memory[uint32(t4):]))
	v0 = t5 + (i32(0)-v1)*i32(680)
	t6 := int32(load32(m.memory[uint32(v0+i32(-676)):]))
	t7 := int32(load32(m.memory[uint32(v0+i32(-672)):]))
	t8 := m.fn544(t2, t3, t6, t7)
	return t8
}
func (m *Module) fn1109(v0, v1 int32) int32 {
	var v2 int32
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(t0):]))
	v2 = t1
	t2 := int32(load32(m.memory[uint32(v2+i32(4)):]))
	t3 := int32(load32(m.memory[uint32(v2+i32(8)):]))
	t4 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t5 := int32(load32(m.memory[uint32(t4):]))
	v0 = t5 + (i32(0)-v1)*i32(28)
	t6 := int32(load32(m.memory[uint32(v0+i32(-24)):]))
	t7 := int32(load32(m.memory[uint32(v0+i32(-20)):]))
	t8 := m.fn545(t2, t3, t6, t7)
	return t8
}
func (m *Module) fn1110(v0, v1 int32) int32 {
	var v2 int32
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(t0):]))
	v2 = t1
	t2 := int32(load32(m.memory[uint32(v2+i32(4)):]))
	t3 := int32(load32(m.memory[uint32(v2+i32(8)):]))
	t4 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t5 := int32(load32(m.memory[uint32(t4):]))
	v0 = t5 + (i32(0)-v1)*i32(36)
	t6 := int32(load32(m.memory[uint32(v0+i32(-32)):]))
	t7 := int32(load32(m.memory[uint32(v0+i32(-28)):]))
	t8 := m.fn545(t2, t3, t6, t7)
	return t8
}
func (m *Module) fn1111(v0, v1 int32) int32 {
	var v2 int32
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(t0):]))
	v2 = t1
	t2 := int32(load32(m.memory[uint32(v2+i32(4)):]))
	t3 := int32(load32(m.memory[uint32(v2+i32(8)):]))
	t4 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t5 := int32(load32(m.memory[uint32(t4):]))
	v0 = t5 + (i32(0)-v1)*i32(28)
	t6 := int32(load32(m.memory[uint32(v0+i32(-24)):]))
	t7 := int32(load32(m.memory[uint32(v0+i32(-20)):]))
	t8 := m.fn544(t2, t3, t6, t7)
	return t8
}
func (m *Module) fn1112(v0, v1 int32) int32 {
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
	t8 := m.fn545(t2, t3, t6, t7)
	return t8
}
func (m *Module) fn1113(v0, v1 int32) int32 {
	var v2 int32
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(t0):]))
	v2 = t1
	t2 := int32(load32(m.memory[uint32(v2):]))
	t3 := int32(load32(m.memory[uint32(v2+i32(4)):]))
	t4 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t5 := int32(load32(m.memory[uint32(t4):]))
	v0 = t5 + (i32(0)-v1)*i32(488)
	t6 := int32(load32(m.memory[uint32(v0+i32(-488)):]))
	t7 := int32(load32(m.memory[uint32(v0+i32(-484)):]))
	t8 := m.fn15(t2, t3, t6, t7)
	return t8
}
func (m *Module) fn1114(v0, v1 int32) int32 {
	var v2 int32
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(t0):]))
	v2 = t1
	t2 := int32(load32(m.memory[uint32(v2):]))
	t3 := int32(load32(m.memory[uint32(v2+i32(4)):]))
	t4 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t5 := int32(load32(m.memory[uint32(t4):]))
	v0 = t5 + (i32(0)-v1)*i32(20)
	t6 := int32(load32(m.memory[uint32(v0+i32(-20)):]))
	t7 := int32(load32(m.memory[uint32(v0+i32(-16)):]))
	t8 := m.fn15(t2, t3, t6, t7)
	return t8
}
func (m *Module) fn1115(v0, v1 int32) int32 {
	var v2 int32
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(t0):]))
	v2 = t1
	t2 := int32(load32(m.memory[uint32(v2):]))
	t3 := int32(load32(m.memory[uint32(v2+i32(4)):]))
	t4 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t5 := int32(load32(m.memory[uint32(t4):]))
	v0 = t5 - v1<<3
	t6 := int32(load32(m.memory[uint32(v0+i32(-8)):]))
	t7 := int32(load32(m.memory[uint32(v0+i32(-4)):]))
	t8 := m.fn15(t2, t3, t6, t7)
	return t8
}
func (m *Module) fn1116(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(t0):]))
	t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t3 := int32(load32(m.memory[uint32(t2):]))
	t4 := m.fn1117(t1, t3+(i32(0)-v1)*i32(24)+i32(-24))
	return t4
}
func (m *Module) fn1117(v0, v1 int32) int32 {
	var v2 int32
	v2 = i32(0)
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		t1 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		t2 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t3 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		t4 := m.fn191(t0, t1, t2, t3)
		if t4 == 0 {
			goto l0
		}
		t5 := int32(load32(m.memory[int64(uint32(v0))+12:]))
		t6 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		var p7 int32
		if t5 == t6 {
			p7 = 1
		}
		v2 = p7
	}
l0:
	return v2
}
func (m *Module) fn1118(v0, v1, v2, v3, v4, v5 int32) {
	var v6 int32
	var v7 int64
	var v8, v9, v10, v11, v12 int32
	t0 := m.g0
	v6 = t0 - i32(32)
	m.g0 = v6
	store32(m.memory[int64(uint32(v6))+16:], uint32(v3))
	store32(m.memory[int64(uint32(v6))+12:], uint32(v2))
	t1 := int64(load64(m.memory[int64(uint32(v1))+16:]))
	t2 := int64(load64(m.memory[int64(uint32(v1))+24:]))
	t3 := m.fn651(t1, t2, v2, v3)
	v7 = t3
	store32(m.memory[int64(uint32(v6))+20:], uint32(v6+i32(12)))
	{
		t4 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		if t4 != 0 {
			goto l0
		}
		_ = m.fn697(v1, v1+i32(16))
	}
l0:
	store32(m.memory[int64(uint32(v6))+24:], uint32(v6+i32(20)))
	store32(m.memory[int64(uint32(v6))+28:], uint32(v1))
	t6 := int32(load32(m.memory[uint32(v1):]))
	t7 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	m.fn69(v6, t6, t7, v7, v6+i32(24), i32(26))
	v8 = i32(1)
	t8 := int32(load32(m.memory[uint32(v1):]))
	v9 = t8
	t9 := int32(load32(m.memory[int64(uint32(v6))+4:]))
	v10 = t9
	{
		{
			t10 := int32(load32(m.memory[uint32(v6):]))
			if t10 != i32(1) {
				goto l1
			}
			v8 = v9 + v10
			t11 := int32(m.memory[uint32(v8)])
			v11 = t11
			t12 := v8
			v12 = int32(uint32(int32(v7)) >> 25)
			m.memory[uint32(t12)] = byte(v12)
			t13 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			m.memory[uint32(v9+t13&(v10+i32(-8))+i32(8))] = byte(v12)
			t14 := int32(load32(m.memory[int64(uint32(v1))+12:]))
			store32(m.memory[int64(uint32(v1))+12:], uint32(t14+i32(1)))
			t15 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			store32(m.memory[int64(uint32(v1))+8:], uint32(t15-v11&i32(1)))
			v1 = v9 - v10<<4
			store32(m.memory[uint32(v1+i32(-4)):], uint32(v5))
			store32(m.memory[uint32(v1+i32(-8)):], uint32(v4))
			store32(m.memory[uint32(v1+i32(-12)):], uint32(v3))
			store32(m.memory[uint32(v1+i32(-16)):], uint32(v2))
			v8 = i32(0)
			goto l2
		}
	l1:
		v1 = v9 - v10<<4
		v3 = v1 + i32(-8)
		t16 := int64(load64(m.memory[uint32(v3):]))
		v7 = t16
		store32(m.memory[uint32(v1+i32(-4)):], uint32(v5))
		store32(m.memory[uint32(v3):], uint32(v4))
		store64(m.memory[int64(uint32(v0))+4:], uint64(v7))
	}
l2:
	store32(m.memory[uint32(v0):], uint32(v8))
	m.g0 = v6 + i32(32)
}
func (m *Module) fn1119(v0, v1 int32) int32 {
	var v2, v3 int32
	v2 = i32(0)
	{
		t0 := int32(load32(m.memory[uint32(v0):]))
		t1 := int32(load32(m.memory[uint32(t0):]))
		v3 = t1
		t2 := int32(load32(m.memory[uint32(v3):]))
		t3 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		t4 := int32(load32(m.memory[uint32(t3):]))
		v0 = t4 - v1<<4
		t5 := int32(load32(m.memory[uint32(v0+i32(-16)):]))
		if t2 != t5 {
			goto l0
		}
		t6 := int32(load32(m.memory[int64(uint32(v3))+4:]))
		t7 := int32(load32(m.memory[uint32(v0+i32(-12)):]))
		var p8 int32
		if t6 == t7 {
			p8 = 1
		}
		v2 = p8
	}
l0:
	return v2
}
func (m *Module) fn1120(v0, v1 int32) int32 {
	var v2, v3 int32
	v2 = i32(0)
	{
		t0 := int32(load32(m.memory[uint32(v0):]))
		t1 := int32(load32(m.memory[uint32(t0):]))
		v3 = t1
		t2 := int32(load32(m.memory[uint32(v3):]))
		t3 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		t4 := int32(load32(m.memory[uint32(t3):]))
		v0 = t4 - v1<<3
		t5 := int32(load32(m.memory[uint32(v0+i32(-8)):]))
		if t2 != t5 {
			goto l0
		}
		t6 := int32(load32(m.memory[int64(uint32(v3))+4:]))
		t7 := int32(load32(m.memory[uint32(v0+i32(-4)):]))
		var p8 int32
		if t6 == t7 {
			p8 = 1
		}
		v2 = p8
	}
l0:
	return v2
}
func (m *Module) fn1121(v0, v1 int32) int32 {
	var v2, v3 int32
	v2 = i32(0)
	{
		t0 := int32(load32(m.memory[uint32(v0):]))
		t1 := int32(load32(m.memory[uint32(t0):]))
		v3 = t1
		t2 := int32(load16(m.memory[uint32(v3):]))
		t3 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		t4 := int32(load32(m.memory[uint32(t3):]))
		v0 = t4 - v1<<3
		t5 := int32(load16(m.memory[uint32(v0+i32(-8)):]))
		if t2 != t5 {
			goto l0
		}
		t6 := int32(load32(m.memory[int64(uint32(v3))+4:]))
		t7 := int32(load32(m.memory[uint32(v0+i32(-4)):]))
		var p8 int32
		if t6 == t7 {
			p8 = 1
		}
		v2 = p8
	}
l0:
	return v2
}
