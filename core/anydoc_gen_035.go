package core

import (
	"math/bits"
)

func (m *Module) fn1527(v0, v1 int32) {
	var v2, v3, v4, v5, v6, v7, v8, v9 int32
	t0 := m.g0
	v2 = t0 - i32(48)
	m.g0 = v2
	t1 := int32(load32(m.memory[int64(uint32(v0))+32:]))
	v3 = t1
	t2 := int32(load32(m.memory[int64(uint32(v0))+28:]))
	t3 := v2
	v0 = t2
	store32(m.memory[int64(uint32(t3))+32:], uint32(v0))
	store32(m.memory[int64(uint32(v2))+36:], uint32(v0+v3*i32(44)))
l1:
	{
		{
			t4 := m.fn904(v2 + i32(32))
			v0 = t4
			if v0 == 0 {
				m.g0 = v2 + i32(48)
				return
			}
			t5 := int32(load32(m.memory[int64(uint32(v0))+36:]))
			v3 = t5
			if v3 == 0 {
				goto l1
			}
			t6 := int32(load32(m.memory[int64(uint32(v0))+40:]))
			t7 := m.fn1337(v3+i32(8), t6, i32(1072544), i32(60))
			if t7 != 0 {
				goto l1
			}
			{
				t8 := int32(load32(m.memory[int64(uint32(v0))+4:]))
				v3 = t8
				t9 := int32(load32(m.memory[int64(uint32(v0))+8:]))
				t10 := v3
				v4 = t9
				t11 := m.fn15(t10, v4, i32(1083095), i32(2))
				if t11 != 0 {
					t14 := int32(load32(m.memory[uint32(v0+i32(28)):]))
					t15 := int32(load32(m.memory[uint32(v0+i32(32)):]))
					t16 := m.fn886(t14, t15, i32(1072544), i32(60), i32(1086211), i32(4))
					v3 = t16
					if v3 != 0 {
						{
							{
								v4 = v3 + i32(28)
								t17 := int32(load32(m.memory[uint32(v4):]))
								v7 = t17
								t18 := v7
								v8 = v3 + i32(32)
								t19 := int32(load32(m.memory[uint32(v8):]))
								v9 = t19
								t20 := m.fn886(t18, v9, i32(1072544), i32(60), i32(1074899), i32(6))
								v3 = t20
								if v3 != 0 {
									goto l6
								}
								v5 = i32(0)
								goto l7
							}
						l6:
							t21 := int32(load32(m.memory[uint32(v3+i32(16)):]))
							t22 := int32(load32(m.memory[uint32(v3+i32(20)):]))
							m.fn1046(v2+i32(24), t21, t22, i32(1072544), i32(60), i32(1073156), i32(3))
							{
								t23 := int32(load32(m.memory[int64(uint32(v2))+24:]))
								v3 = t23
								if v3 != 0 {
									goto l8
								}
								v5 = i32(1)
								goto l7
							}
						l8:
							t24 := int32(load32(m.memory[int64(uint32(v2))+28:]))
							t25 := m.fn15(v3, t24, i32(1074878), i32(7))
							v5 = t25 ^ i32(1)
						}
					l7:
						v6 = i32(1)
						{
							t26 := m.fn886(v7, v9, i32(1072544), i32(60), i32(1074885), i32(8))
							v3 = t26
							if v3 == 0 {
								goto l9
							}
							t27 := int32(load32(m.memory[uint32(v3+i32(16)):]))
							t28 := int32(load32(m.memory[uint32(v3+i32(20)):]))
							m.fn1046(v2+i32(16), t27, t28, i32(1072544), i32(60), i32(1073156), i32(3))
							t29 := int32(load32(m.memory[int64(uint32(v2))+16:]))
							v3 = t29
							if v3 == 0 {
								goto l9
							}
							t30 := int32(load32(m.memory[int64(uint32(v2))+20:]))
							m.fn197(v2+i32(40), v3, t30)
							t31 := int32(m.memory[int64(uint32(v2))+40])
							if t31 != 0 {
								goto l9
							}
							t32 := int32(load32(m.memory[int64(uint32(v2))+44:]))
							v3 = t32
							p33 := i32(1000)
							if uint32(v3) < uint32(i32(1000)) {
								p33 = v3
							}
							p34 := i32(1)
							if v3 != 0 {
								p34 = p33
							}
							v6 = p34
						}
					l9:
						t35 := int32(load32(m.memory[uint32(v4):]))
						t36 := int32(load32(m.memory[uint32(v8):]))
						t37 := m.fn886(t35, t36, i32(1072544), i32(60), i32(1074893), i32(6))
						v3 = t37
						if v3 == 0 {
							goto l5
						}
						t38 := int32(load32(m.memory[uint32(v3+i32(16)):]))
						t39 := int32(load32(m.memory[uint32(v3+i32(20)):]))
						m.fn1046(v2+i32(8), t38, t39, i32(1072544), i32(60), i32(1073156), i32(3))
						{
							t40 := int32(load32(m.memory[int64(uint32(v2))+8:]))
							v3 = t40
							if v3 == 0 {
								goto l10
							}
							t41 := int32(load32(m.memory[int64(uint32(v2))+12:]))
							t42 := m.fn15(v3, t41, i32(1074878), i32(7))
							if t42 != 0 {
								goto l5
							}
						}
					l10:
						t43 := int32(load32(m.memory[int64(uint32(v1))+8:]))
						v3 = t43
						if v3 == 0 {
							goto l5
						}
						t44 := int32(load32(m.memory[int64(uint32(v1))+4:]))
						v3 = t44 + v3*i32(28)
						v4 = v3 + i32(-28)
						if v4 == 0 {
							goto l5
						}
						t45 := int32(m.memory[uint32(v3+i32(-4))])
						if t45 != 0 {
							goto l5
						}
						v3 = v3 + i32(-12)
						t46 := int32(load32(m.memory[uint32(v3):]))
						store32(m.memory[uint32(v3):], uint32(t46+v6))
						m.fn1549(v4, v0)
						goto l1
					}
					v5 = i32(0)
					v6 = i32(1)
					goto l5
				}
				t12 := m.fn15(v3, v4, i32(1077645), i32(3))
				if t12 != 0 {
					t47 := int32(load32(m.memory[uint32(v0+i32(28)):]))
					t48 := int32(load32(m.memory[uint32(v0+i32(32)):]))
					t49 := m.fn886(t47, t48, i32(1072544), i32(60), i32(1077680), i32(10))
					v0 = t49
					if v0 == 0 {
						goto l1
					}
					m.fn1527(v0, v1)
					goto l1
				}
				t13 := m.fn15(v3, v4, i32(1077671), i32(9))
				if t13 == 0 {
					goto l1
				}
				m.fn1527(v0, v1)
				goto l1
			}
		}
	l5:
		{
			t50 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			v4 = t50
			t51 := int32(load32(m.memory[uint32(v1):]))
			if v4 != t51 {
				goto l11
			}
			m.fn1143(v1)
		}
	l11:
		t52 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		v3 = t52 + v4*i32(28)
		m.memory[int64(uint32(v3))+24] = byte(v5)
		store32(m.memory[int64(uint32(v3))+20:], uint32(i32(1)))
		store32(m.memory[int64(uint32(v3))+16:], uint32(v6))
		store32(m.memory[int64(uint32(v3))+12:], uint32(v0))
		store32(m.memory[int64(uint32(v3))+8:], uint32(i32(0)))
		store64(m.memory[uint32(v3):], uint64(i64(0x400000000)))
		store32(m.memory[int64(uint32(v1))+8:], uint32(v4+i32(1)))
		goto l1
	}
}
func (m *Module) fn1528(v0, v1, v2, v3 int32) int32 {
	if uint32(v2) < uint32(v1) {
		return v0 + v2*i32(28)
	}
	m.fn158(v2, v1, v3)
	panic("unreachable")
}
func (m *Module) fn1529(v0, v1 int32) {
	var v2 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		v2 = t0
		t1 := int32(load32(m.memory[uint32(v0):]))
		if v2 != t1 {
			goto l0
		}
		m.fn272(v0)
	}
l0:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v2+i32(1)))
	t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	v0 = t2 + v2*i32(12)
	t3 := int64(load64(m.memory[uint32(v1):]))
	store64(m.memory[uint32(v0):], uint64(t3))
	t4 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	store32(m.memory[int64(uint32(v0))+8:], uint32(t4))
}
func (m *Module) fn1530(v0, v1 int32) {
	var v2 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	{
		if v1 == 0 {
			goto l0
		}
		m.fn39(v2+i32(4), i32(12), i32(8), v1+i32(1))
		t1 := int32(load32(m.memory[int64(uint32(v2))+12:]))
		t2 := int32(load32(m.memory[int64(uint32(v2))+4:]))
		t3 := int32(load32(m.memory[int64(uint32(v2))+8:]))
		m.fn40(v0-t1, t2, t3)
	}
l0:
	m.g0 = v2 + i32(16)
}
func (m *Module) fn1531(v0, v1 int32) int32 {
	t0 := m.fn1361(v0, v1, i32(1076204), i32(11))
	return t0
}
func (m *Module) fn1532(v0 int32) {
	var v1, v2, v3, v4, v5, v6, v7 int32
	t0 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	v1 = t0
	t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	v2 = t1
	v3 = i32(0)
l3:
	{
		if v3 == v1 {
			goto l0
		}
		v4 = v2 + v3*i32(12)
		t2 := int32(load32(m.memory[uint32(v4+i32(8)):]))
		v5 = t2
		v6 = v4 + i32(4)
		t3 := int32(load32(m.memory[uint32(v6):]))
		v7 = t3
	l2:
		{
			if v5 == 0 {
				t6 := int32(load32(m.memory[uint32(v4):]))
				t7 := int32(load32(m.memory[uint32(v6):]))
				m.fn136(t6, t7, i32(4), i32(28))
				v3 = v3 + i32(1)
				goto l3
			}
			t4 := int32(load32(m.memory[uint32(v7):]))
			t5 := int32(load32(m.memory[uint32(v7+i32(4)):]))
			m.fn44(t4, t5)
			v5 = v5 + i32(-1)
			v7 = v7 + i32(28)
			goto l2
		}
	}
l0:
	t8 := int32(load32(m.memory[uint32(v0):]))
	m.fn136(t8, v2, i32(4), i32(12))
}
func (m *Module) fn1533(v0, v1, v2, v3 int32) {
	var v4 int64
	{
		{
			t0 := int32(load32(m.memory[int64(uint32(v1))+12:]))
			if t0 != 0 {
				goto l0
			}
			v1 = i32(0)
			goto l1
		}
	l0:
		t1 := int64(load64(m.memory[int64(uint32(v1))+16:]))
		t2 := int64(load64(m.memory[int64(uint32(v1))+24:]))
		t3 := m.fn29(t1, t2, v2, v3)
		v4 = t3
		t4 := int32(load32(m.memory[uint32(v1):]))
		t5 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t6 := m.fn1320(t4, t5, v4, v2, v3)
		v3 = t6
		p7 := i32(0)
		if v3 != 0 {
			p7 = v3 + i32(-20)
		}
		v1 = p7
		v3 = v3 + i32(-12)
	}
l1:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
	store32(m.memory[uint32(v0):], uint32(v1))
}
func (m *Module) fn1534(v0, v1, v2 int32) int32 {
	t0 := m.fn1192(v0, v1, v2)
	v2 = t0
	if v2 == 0 {
		m.fn633(i32(1087080), i32(22), i32(1070828))
		panic("unreachable")
	}
	return v2
}
func (m *Module) fn1535(v0, v1, v2, v3, v4 int32) {
	var v5, v6, v7, v8, v9 int32
	var v10 int64
	var v11 int32
	var v12, v13 int64
	var v14 int32
	var v15 int64
	var v16, v17, v18, v19 int32
	var v20 int64
	var v21 int32
	t0 := m.g0
	v5 = t0 - i32(160)
	m.g0 = v5
	v6 = i32(0)
	{
		{
			if v1 == 0 {
				goto l4
			}
			t1 := int32(load32(m.memory[uint32(v1+i32(28)):]))
			t2 := int32(load32(m.memory[uint32(v1+i32(32)):]))
			t3 := m.fn886(t1, t2, i32(1072544), i32(60), i32(1073738), i32(5))
			v1 = t3
			if v1 == 0 {
				goto l4
			}
			v7 = i32(0)
			{
				t4 := int32(load32(m.memory[uint32(v1+i32(28)):]))
				v8 = t4
				t5 := int32(load32(m.memory[uint32(v1+i32(32)):]))
				t6 := v8
				v9 = t5
				t7 := m.fn886(t6, v9, i32(1072544), i32(60), i32(1072649), i32(5))
				v6 = t7
				if v6 == 0 {
					goto l2
				}
				t8 := int32(load32(m.memory[uint32(v6+i32(16)):]))
				t9 := int32(load32(m.memory[uint32(v6+i32(20)):]))
				m.fn1046(v5+i32(72), t8, t9, i32(1072544), i32(60), i32(1073156), i32(3))
				t10 := int32(load32(m.memory[int64(uint32(v5))+72:]))
				v6 = t10
				if v6 == 0 {
					goto l2
				}
				t11 := int32(load32(m.memory[int64(uint32(v5))+76:]))
				m.fn1190(v5+i32(120), v6, t11)
				t12 := int32(m.memory[int64(uint32(v5))+120])
				v7 = t12 ^ i32(1)
				t13 := int64(load64(m.memory[int64(uint32(v5))+128:]))
				v10 = t13
				t14 := int32(load32(m.memory[uint32(v1+i32(32)):]))
				v9 = t14
				t15 := int32(load32(m.memory[uint32(v1+i32(28)):]))
				v8 = t15
				goto l2
			}
		l2:
			{
				t16 := m.fn886(v8, v9, i32(1072544), i32(60), i32(1072629), i32(4))
				v1 = t16
				if v1 == 0 {
					goto l3
				}
				t17 := int32(load32(m.memory[uint32(v1+i32(16)):]))
				t18 := int32(load32(m.memory[uint32(v1+i32(20)):]))
				m.fn1046(v5+i32(64), t17, t18, i32(1072544), i32(60), i32(1073156), i32(3))
				t19 := int32(load32(m.memory[int64(uint32(v5))+64:]))
				v1 = t19
				if v1 == 0 {
					goto l3
				}
				t20 := int32(load32(m.memory[int64(uint32(v5))+68:]))
				m.fn197(v5+i32(120), v1, t20)
				t21 := int32(m.memory[int64(uint32(v5))+120])
				v6 = t21 ^ i32(1)
				t22 := int32(load32(m.memory[int64(uint32(v5))+124:]))
				v11 = t22
				if v7&i32(1) == 0 {
					goto l4
				}
				goto l5
			}
		l3:
			v6 = i32(0)
			if v7&i32(1) != 0 {
				goto l5
			}
			goto l4
		}
	l4:
		if v2 == 0 {
			goto l6
		}
		t23 := int32(load32(m.memory[int64(uint32(v4))+36:]))
		v8 = t23
		m.fn27(v5 + i32(120))
		m.fn1533(v5+i32(56), v8, v2, v3)
		{
			{
				t24 := int32(load32(m.memory[int64(uint32(v5))+56:]))
				v1 = t24
				if v1 != 0 {
					goto l7
				}
				v1 = i32(0)
				goto l15
			}
		l7:
			t25 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			v7 = t25
			t26 := int32(load32(m.memory[uint32(v1):]))
			v1 = t26
		}
	l15:
		if v1 == 0 {
			goto l9
		}
		store32(m.memory[int64(uint32(v5))+156:], uint32(v7))
		store32(m.memory[int64(uint32(v5))+152:], uint32(v1))
		{
			t27 := m.fn1521(v5+i32(120), v1, v7)
			if t27 != 0 {
				t29 := m.fn1534(v8, v1, v7)
				v1 = t29
				t30 := int32(load32(m.memory[int64(uint32(v1))+8:]))
				v9 = t30
				t31 := int32(load32(m.memory[int64(uint32(v1))+4:]))
				v7 = t31
				{
					t32 := int32(load32(m.memory[uint32(v1):]))
					v1 = t32
					t33 := int32(load32(m.memory[uint32(v1+i32(28)):]))
					t34 := int32(load32(m.memory[uint32(v1+i32(32)):]))
					t35 := m.fn886(t33, t34, i32(1072544), i32(60), i32(1073735), i32(3))
					v1 = t35
					if v1 == 0 {
						goto l12
					}
					t36 := int32(load32(m.memory[uint32(v1+i32(28)):]))
					t37 := int32(load32(m.memory[uint32(v1+i32(32)):]))
					t38 := m.fn886(t36, t37, i32(1072544), i32(60), i32(1073738), i32(5))
					v1 = t38
					if v1 == 0 {
						goto l12
					}
					t39 := int32(load32(m.memory[uint32(v1+i32(28)):]))
					t40 := int32(load32(m.memory[uint32(v1+i32(32)):]))
					t41 := m.fn886(t39, t40, i32(1072544), i32(60), i32(1072649), i32(5))
					v1 = t41
					if v1 == 0 {
						goto l12
					}
					t42 := int32(load32(m.memory[uint32(v1+i32(16)):]))
					t43 := int32(load32(m.memory[uint32(v1+i32(20)):]))
					m.fn1046(v5+i32(48), t42, t43, i32(1072544), i32(60), i32(1073156), i32(3))
					t44 := int32(load32(m.memory[int64(uint32(v5))+48:]))
					v1 = t44
					if v1 == 0 {
						goto l12
					}
					t45 := int32(load32(m.memory[int64(uint32(v5))+52:]))
					m.fn1190(v5+i32(104), v1, t45)
					t46 := int32(m.memory[int64(uint32(v5))+104])
					if t46 != 0 {
						goto l12
					}
					t47 := int64(load64(m.memory[int64(uint32(v5))+112:]))
					store64(m.memory[int64(uint32(v5))+96:], uint64(t47))
					store64(m.memory[int64(uint32(v5))+88:], uint64(i64(1)))
					goto l13
				}
			l12:
				v1 = i32(0)
				if v7 != 0 {
					m.fn1533(v5+i32(40), v8, v7, v9)
					t48 := int32(load32(m.memory[int64(uint32(v5))+40:]))
					v9 = t48
					if v9 == 0 {
						goto l15
					}
					t49 := int32(load32(m.memory[int64(uint32(v9))+4:]))
					v7 = t49
					t50 := int32(load32(m.memory[uint32(v9):]))
					v1 = t50
					goto l15
				}
				goto l15
			}
			store32(m.memory[int64(uint32(v5))+108:], uint32(i32(71)))
			store32(m.memory[int64(uint32(v5))+104:], uint32(v5+i32(152)))
			m.fn73(v5+i32(80), i32(1049807), v5+i32(104))
			store32(m.memory[int64(uint32(v5))+92:], uint32(i32(-1)))
			t28 := int32(load32(m.memory[int64(uint32(v5))+80:]))
			v1 = t28
			goto l11
		}
	l9:
		store64(m.memory[int64(uint32(v5))+88:], uint64(i64(0)))
	l13:
		v1 = i32(-1)
	l11:
		t51 := int32(load32(m.memory[int64(uint32(v5))+120:]))
		t52 := int32(load32(m.memory[int64(uint32(v5))+124:]))
		m.fn56(t51, t52)
		{
			if v1 == i32(-1) {
				goto l16
			}
			t53 := int32(load32(m.memory[int64(uint32(v5))+84:]))
			v7 = t53
			t54 := int64(load64(m.memory[int64(uint32(v5))+88:]))
			v10 = t54
			t55 := int64(load64(m.memory[int64(uint32(v5))+96:]))
			v12 = t55
			store32(m.memory[int64(uint32(v0))+28:], uint32(i32(-3)))
			store64(m.memory[int64(uint32(v0))+16:], uint64(v12))
			store64(m.memory[int64(uint32(v0))+8:], uint64(v10))
			store32(m.memory[int64(uint32(v0))+4:], uint32(v7))
			store32(m.memory[uint32(v0):], uint32(v1))
			goto l17
		}
	l16:
		t56 := int64(load64(m.memory[int64(uint32(v5))+88:]))
		if t56 != i64(1) {
			goto l6
		}
		t57 := int64(load64(m.memory[int64(uint32(v5))+96:]))
		v10 = t57
	}
l5:
	if v10 != i64(0) {
		t58 := int32(load32(m.memory[int64(uint32(v4))+40:]))
		v1 = t58
		t59 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		if t59 == 0 {
			goto l19
		}
		t60 := int64(load64(m.memory[int64(uint32(v1))+16:]))
		t61 := int64(load64(m.memory[uint32(v1+i32(24)):]))
		t62 := m.fn741(t60, t61, v10)
		v12 = t62
		t63 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		v9 = t63
		v7 = v9 & int32(v12)
		v13 = int64(uint64(v12)>>25) & i64(127) * i64(72340172838076673)
		t64 := int32(load32(m.memory[uint32(v1):]))
		v1 = t64
		v14 = i32(0)
	l65:
		{
			t65 := int64(load64(m.memory[uint32(v1+v7):]))
			v15 = t65
			v12 = v15 ^ v13
			v12 = (v12 ^ i64(-1)) & (v12 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
			{
			l22:
				{
					var p66 int32
					if v12 == 0 {
						p66 = 1
					}
					v8 = p66
					if v8 != 0 {
						goto l20
					}
					t67 := v10
					t68 := v1
					v16 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v12))))>>3) + v7) & v9
					t69 := int64(load64(m.memory[uint32(t68+(i32(0)-v16)*i32(480)+i32(-480)):]))
					if t67 == t69 {
						goto l21
					}
					v12 = (v12 + i64(-1)) & v12
					goto l22
				}
			l20:
				if v15&(v15<<1)&i64(-0x7f7f7f7f7f7f7f80) == 0 {
					t161 := v7
					v14 = v14 + i32(8)
					v7 = (t161 + v14) & v9
					goto l65
				}
			l21:
				if v8 != 0 {
					goto l19
				}
				p70 := v1 + (i32(0)-v16)*i32(480)
				if v8 != 0 {
					p70 = i32(0)
				}
				v8 = p70
				{
					if v6&i32(1) != 0 {
						goto l24
					}
					if v2 != 0 {
						goto l25
					}
					v11 = i32(0)
					goto l24
				l25:
					t71 := int32(load32(m.memory[int64(uint32(v4))+36:]))
					v6 = t71
					m.fn27(v5 + i32(120))
					m.fn1533(v5+i32(32), v6, v2, v3)
					{
						{
							t72 := int32(load32(m.memory[int64(uint32(v5))+32:]))
							v1 = t72
							if v1 != 0 {
								goto l26
							}
							v1 = i32(0)
							goto l35
						}
					l26:
						t73 := int32(load32(m.memory[int64(uint32(v1))+4:]))
						v7 = t73
						t74 := int32(load32(m.memory[uint32(v1):]))
						v1 = t74
					}
				l35:
					if v1 != 0 {
						store32(m.memory[int64(uint32(v5))+156:], uint32(v7))
						store32(m.memory[int64(uint32(v5))+152:], uint32(v1))
						{
							t75 := m.fn1521(v5+i32(120), v1, v7)
							if t75 != 0 {
								t79 := m.fn1534(v6, v1, v7)
								v1 = t79
								t80 := int32(load32(m.memory[int64(uint32(v1))+8:]))
								v11 = t80
								t81 := int32(load32(m.memory[int64(uint32(v1))+4:]))
								v3 = t81
								t82 := int32(load32(m.memory[uint32(v1):]))
								t83 := v5 + i32(24)
								v1 = t82
								t84 := int32(load32(m.memory[uint32(v1+i32(16)):]))
								t85 := int32(load32(m.memory[uint32(v1+i32(20)):]))
								m.fn1046(t83, t84, t85, i32(1072544), i32(60), i32(1073766), i32(7))
								{
									t86 := int32(load32(m.memory[int64(uint32(v5))+24:]))
									v9 = t86
									if v9 == 0 {
										goto l31
									}
									t87 := int32(load32(m.memory[int64(uint32(v5))+28:]))
									v2 = t87
									v7 = i32(0)
									v1 = i32(-112)
								l33:
									if v1 == i32(-4) {
										goto l31
									}
									m.fn855(v5+i32(16), v8+v1)
									{
										t88 := int32(load32(m.memory[int64(uint32(v5))+16:]))
										t89 := int32(load32(m.memory[int64(uint32(v5))+20:]))
										t90 := m.fn848(t88, t89, v9, v2)
										if t90 == 0 {
											v1 = v1 + i32(12)
											v7 = v7 + i32(1)
											goto l33
										}
										v1 = i32(-1)
										v9 = i32(1)
										goto l29
									}
								}
							l31:
								v1 = i32(0)
								if v3 != 0 {
									m.fn1533(v5+i32(8), v6, v3, v11)
									t91 := int32(load32(m.memory[int64(uint32(v5))+8:]))
									v9 = t91
									if v9 == 0 {
										goto l35
									}
									t92 := int32(load32(m.memory[int64(uint32(v9))+4:]))
									v7 = t92
									t93 := int32(load32(m.memory[uint32(v9):]))
									v1 = t93
									goto l35
								}
								goto l35
							}
							store32(m.memory[int64(uint32(v5))+108:], uint32(i32(71)))
							store32(m.memory[int64(uint32(v5))+104:], uint32(v5+i32(152)))
							m.fn73(v5+i32(80), i32(1049807), v5+i32(104))
							store32(m.memory[int64(uint32(v5))+92:], uint32(i32(-1)))
							t76 := int32(load32(m.memory[int64(uint32(v5))+88:]))
							v7 = t76
							t77 := int32(load32(m.memory[int64(uint32(v5))+84:]))
							v9 = t77
							t78 := int32(load32(m.memory[int64(uint32(v5))+80:]))
							v1 = t78
							goto l29
						}
					}
					v9 = i32(0)
					v1 = i32(-1)
					goto l29
				l29:
					t94 := int32(load32(m.memory[int64(uint32(v5))+120:]))
					t95 := int32(load32(m.memory[int64(uint32(v5))+124:]))
					m.fn56(t94, t95)
					{
						if v1 == i32(-1) {
							goto l36
						}
						t96 := int32(load32(m.memory[int64(uint32(v5))+100:]))
						store32(m.memory[int64(uint32(v0))+20:], uint32(t96))
						t97 := int64(load64(m.memory[int64(uint32(v5))+92:]))
						store64(m.memory[int64(uint32(v0))+12:], uint64(t97))
						store32(m.memory[int64(uint32(v0))+28:], uint32(i32(-3)))
						store32(m.memory[int64(uint32(v0))+8:], uint32(v7))
						store32(m.memory[int64(uint32(v0))+4:], uint32(v9))
						store32(m.memory[uint32(v0):], uint32(v1))
						goto l17
					}
				l36:
					p98 := i32(0)
					if v9&i32(1) != 0 {
						p98 = v7
					}
					v11 = p98
				}
			l24:
				{
					v3 = v8 + i32(-472)
					t100 := v3
					p99 := i32(8)
					if uint32(v11) < uint32(i32(8)) {
						p99 = v11
					}
					v16 = p99
					v17 = t100 + v16*i32(40)
					t101 := int32(m.memory[int64(uint32(v17))+32])
					v18 = t101
					if v18 != 0 {
						if v18 != i32(255) {
							t102 := int32(load32(m.memory[int64(uint32(v4))+44:]))
							v14 = t102
							t103 := int32(load32(m.memory[uint32(v14):]))
							if t103 != 0 {
								m.fn1326(i32(1086216))
								panic("unreachable")
							}
							store32(m.memory[uint32(v14):], uint32(i32(-1)))
							t104 := int64(load64(m.memory[int64(uint32(v14))+24:]))
							t105 := int64(load64(m.memory[int64(uint32(v14))+32:]))
							t106 := m.fn741(t104, t105, v10)
							v12 = t106
							t107 := int32(load32(m.memory[int64(uint32(v14))+12:]))
							v9 = t107
							t108 := v9
							v2 = int32(v12)
							v8 = t108 & v2
							v15 = int64(uint64(v12)>>25) & i64(127) * i64(72340172838076673)
							v6 = v14 + i32(24)
							v19 = v14 + i32(8)
							t109 := int32(load32(m.memory[int64(uint32(v14))+8:]))
							v7 = t109
							v4 = i32(0)
						l64:
							{
								t110 := int64(load64(m.memory[uint32(v7+v8):]))
								v20 = t110
								v13 = v20 ^ v15
								v13 = (v13 ^ i64(-1)) & (v13 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
								{
									{
									l43:
										{
											if v13 == 0 {
												goto l41
											}
											v1 = v7 + (i32(0)-(int32(uint32(int64(bits.TrailingZeros64(uint64(v13))))>>3)+v8)&v9)*i32(104)
											t111 := int64(load64(m.memory[uint32(v1+i32(-104)):]))
											if t111 == v10 {
												goto l42
											}
											v13 = (v13 + i64(-1)) & v13
											goto l43
										}
									l41:
										if v20&(v20<<1)&i64(-0x7f7f7f7f7f7f7f80) == 0 {
											t160 := v8
											v4 = v4 + i32(8)
											v8 = (t160 + v4) & v9
											goto l64
										}
										{
											t112 := int32(load32(m.memory[int64(uint32(v14))+16:]))
											if t112 != 0 {
												goto l45
											}
											_ = m.fn739(v19, v6)
											t114 := int32(load32(m.memory[int64(uint32(v14))+12:]))
											v9 = t114
											t115 := int32(load32(m.memory[int64(uint32(v14))+8:]))
											v7 = t115
										}
									l45:
										t116 := m.fn26(v7, v9, v12)
										t117 := v7
										v1 = t116
										v8 = t117 + v1
										t118 := int32(m.memory[uint32(v8)])
										v4 = t118
										t119 := v8
										v2 = int32(uint32(v2) >> 25)
										m.memory[uint32(t119)] = byte(v2)
										m.memory[uint32(v7+v9&(v1+i32(-8))+i32(8))] = byte(v2)
										t120 := int32(load32(m.memory[int64(uint32(v14))+16:]))
										store32(m.memory[int64(uint32(v14))+16:], uint32(t120-v4&i32(1)))
										v1 = v7 + (i32(0)-v1)*i32(104)
										store64(m.memory[uint32(v1+i32(-104)):], uint64(v10))
										memory_zero(m.memory, uint32(v1+i32(-96)), uint32(i32(90)))
										t121 := int32(load32(m.memory[int64(uint32(v14))+20:]))
										store32(m.memory[int64(uint32(v14))+20:], uint32(t121+i32(1)))
									}
								l42:
									v21 = v1 + i32(-96)
									{
										{
											v19 = v1 + i32(-24)
											v7 = v19 + v16
											t122 := int32(m.memory[uint32(v7)])
											if t122 != i32(1) {
												goto l46
											}
											t123 := int32(m.memory[uint32(v1+v16+i32(-15))])
											if t123 != 0 {
												goto l46
											}
											v7 = v21 + v16<<3
											t124 := int64(load64(m.memory[uint32(v7):]))
											t125 := v7
											v12 = t124 + i64(1)
											p126 := v12
											if v12 == 0 {
												p126 = i64(-1)
											}
											v12 = p126
											store64(m.memory[uint32(t125):], uint64(v12))
											goto l47
										}
									l46:
										m.memory[uint32(v7)] = byte(i32(1))
										t127 := int64(load64(m.memory[int64(uint32(v17))+24:]))
										t128 := v21 + v16<<3
										v12 = t127
										store64(m.memory[uint32(t128):], uint64(v12))
										m.memory[uint32(v1+v16+i32(-15))] = byte(i32(0))
									}
								l47:
									v6 = v1 + i32(-15)
									v8 = v16 + i32(1)
									v1 = i32(0)
									v2 = i32(0)
								l52:
									{
										if v8 != 0 {
											goto l48
										}
										v7 = v2
										if v1 != i32(360) {
											goto l49
										}
										goto l50
									l48:
										t129 := int32(uint32((i32(360)-v1)&i32(0xffff)) / uint32(i32(40)))
										if uint32(v8) >= uint32(t129) {
											goto l50
										}
										v7 = v8 + v2
										v1 = v8*i32(40) + v1
									}
								l49:
									v2 = v7 + i32(1)
									v9 = v1 + i32(40)
									{
										v4 = v3 + v1
										t130 := int32(load32(m.memory[uint32(v4):]))
										if t130 == 0 {
											goto l51
										}
										v8 = i32(0)
										v1 = v9
										t131 := int32(load32(m.memory[int64(uint32(v4))+4:]))
										if uint32(v16) >= uint32(t131) {
											goto l52
										}
									}
								l51:
									if uint32(v7) > uint32(i32(8)) {
										goto l53
									}
									m.memory[uint32(v6+v7)] = byte(i32(1))
									v8 = i32(0)
									v1 = v9
									goto l52
								l53:
									m.fn158(v7, i32(9), i32(1077944))
									panic("unreachable")
								l50:
									v1 = i32(-1)
									{
										t132 := int32(m.memory[int64(uint32(v17))+32])
										v4 = t132
										if v4 == i32(255) {
											goto l61
										}
										t133 := int32(load32(m.memory[int64(uint32(v17))+16:]))
										v7 = t133
										if v7 == 0 {
											goto l61
										}
										store32(m.memory[int64(uint32(v5))+88:], uint32(i32(0)))
										store64(m.memory[int64(uint32(v5))+80:], uint64(i64(0x100000000)))
										v7 = v7 * i32(12)
										t134 := int32(m.memory[int64(uint32(v17))+20])
										v2 = t134
										t135 := int32(load32(m.memory[int64(uint32(v17))+12:]))
										v1 = t135
									l63:
										{
											{
												if v7 == 0 {
													m.fn800(v5+i32(120), v4, v12)
													t143 := int32(load32(m.memory[int64(uint32(v5))+84:]))
													v8 = t143
													t144 := int32(load32(m.memory[int64(uint32(v5))+88:]))
													t145 := int32(load32(m.memory[int64(uint32(v5))+124:]))
													t146 := v8
													v1 = t145
													t147 := int32(load32(m.memory[int64(uint32(v5))+128:]))
													t148 := m.fn191(t146, t144, v1, t147)
													v7 = t148
													t149 := int32(load32(m.memory[int64(uint32(v5))+120:]))
													m.fn16(t149, v1)
													t150 := int32(load32(m.memory[int64(uint32(v5))+80:]))
													v1 = t150
													{
														if v7 != 0 {
															m.fn16(v1, v8)
															v1 = i32(-1)
															goto l61
														}
														t151 := int64(load64(m.memory[int64(uint32(v5))+84:]))
														v13 = t151
														goto l61
													}
												}
												t136 := int32(load32(m.memory[uint32(v1):]))
												if t136 != i32(-1) {
													t152 := int32(load32(m.memory[uint32(v1+i32(4)):]))
													t153 := int32(load32(m.memory[uint32(v1+i32(8)):]))
													m.fn75(v5+i32(80), t152, t153)
													goto l62
												}
												t137 := int32(m.memory[uint32(v1+i32(4))])
												v8 = t137
												p138 := i32(8)
												if uint32(v8) < uint32(i32(8)) {
													p138 = v8
												}
												v8 = p138
												v9 = i32(1)
												{
													if v2&i32(1) != 0 {
														goto l57
													}
													t139 := int32(m.memory[int64(uint32(v3+v8*i32(40)))+32])
													v9 = t139
													p140 := v9
													if v9 == i32(255) {
														p140 = i32(1)
													}
													v9 = p140
												}
											l57:
												t141 := int32(m.memory[uint32(v19+v8)])
												if t141 != i32(1) {
													goto l58
												}
												t142 := int32(m.memory[uint32(v6+v8)])
												if t142 != 0 {
													goto l58
												}
												v8 = v21 + v8<<3
												goto l59
											}
										l58:
											v8 = v3 + v8*i32(40) + i32(24)
										l59:
											t154 := int64(load64(m.memory[uint32(v8):]))
											m.fn804(v5+i32(120), v9, t154)
											t155 := int32(load32(m.memory[int64(uint32(v5))+124:]))
											t156 := v5 + i32(80)
											v8 = t155
											t157 := int32(load32(m.memory[int64(uint32(v5))+128:]))
											m.fn75(t156, v8, t157)
											t158 := int32(load32(m.memory[int64(uint32(v5))+120:]))
											m.fn16(t158, v8)
										}
									l62:
										v1 = v1 + i32(12)
										v7 = v7 + i32(-12)
										goto l63
									}
								l61:
									t159 := int32(load32(m.memory[uint32(v14):]))
									store32(m.memory[uint32(v14):], uint32(t159+i32(1)))
									goto l38
								}
							}
						}
						store32(m.memory[int64(uint32(v0))+28:], uint32(i32(-2)))
						goto l17
					}
					v1 = i32(-1)
					v12 = i64(0)
					goto l38
				}
			l38:
				store64(m.memory[int64(uint32(v0))+32:], uint64(v13))
				store32(m.memory[int64(uint32(v0))+28:], uint32(v1))
				store32(m.memory[int64(uint32(v0))+24:], uint32(v11))
				store64(m.memory[int64(uint32(v0))+16:], uint64(v12))
				m.memory[int64(uint32(v0))+8] = byte(v18)
				store64(m.memory[uint32(v0):], uint64(v10))
				goto l17
			}
		}
	}
	store32(m.memory[int64(uint32(v0))+28:], uint32(i32(-2)))
	goto l17
l6:
	store32(m.memory[int64(uint32(v0))+28:], uint32(i32(-2)))
	goto l17
l19:
	store32(m.memory[int64(uint32(v0))+28:], uint32(i32(-2)))
l17:
	m.g0 = v5 + i32(160)
}
func (m *Module) fn1536(v0, v1, v2, v3 int32) {
	var v4, v5, v6, v7, v8, v9, v10 int32
	t0 := m.g0
	v4 = t0 - i32(96)
	m.g0 = v4
	m.fn27(v4 + i32(48))
	m.fn1533(v4+i32(16), v1, v2, v3)
	{
		{
			t1 := int32(load32(m.memory[int64(uint32(v4))+16:]))
			v3 = t1
			if v3 != 0 {
				goto l0
			}
			v3 = i32(0)
			goto l1
		}
	l0:
		t2 := int32(load32(m.memory[int64(uint32(v3))+4:]))
		v2 = t2
		t3 := int32(load32(m.memory[uint32(v3):]))
		v3 = t3
	}
l1:
	v5 = i32(0)
l8:
	if v3 == 0 {
		goto l2
	}
	store32(m.memory[int64(uint32(v4))+84:], uint32(v2))
	store32(m.memory[int64(uint32(v4))+80:], uint32(v3))
	{
		t4 := m.fn1521(v4+i32(48), v3, v2)
		if t4 != 0 {
			t6 := m.fn1534(v1, v3, v2)
			v3 = t6
			t7 := int32(load32(m.memory[int64(uint32(v3))+8:]))
			v6 = t7
			t8 := int32(load32(m.memory[int64(uint32(v3))+4:]))
			v2 = t8
			{
				t9 := int32(load32(m.memory[uint32(v3):]))
				v3 = t9
				t10 := int32(load32(m.memory[uint32(v3+i32(28)):]))
				t11 := int32(load32(m.memory[uint32(v3+i32(32)):]))
				t12 := m.fn886(t10, t11, i32(1072544), i32(60), i32(1073717), i32(3))
				v3 = t12
				if v3 == 0 {
					goto l5
				}
				v7 = i32(1)
				t13 := int32(load32(m.memory[uint32(v3+i32(28)):]))
				v8 = t13
				t14 := int32(load32(m.memory[uint32(v3+i32(32)):]))
				t15 := v8
				v3 = t14
				t16 := m.fn1318(t15, v3, i32(1073720), i32(1))
				v9 = t16 & i32(255)
				t17 := m.fn1318(v8, v3, i32(1073721), i32(1))
				v10 = t17 & i32(255)
				{
					t18 := m.fn1318(v8, v3, i32(1073722), i32(6))
					if t18&i32(255) == i32(1) {
						goto l6
					}
					t19 := m.fn1318(v8, v3, i32(1073728), i32(7))
					var p20 int32
					if t19&i32(255) == i32(1) {
						p20 = 1
					}
					v7 = p20
				}
			l6:
				p21 := i32(0)
				if int32(uint32(v5&i32(65536))>>16)^v7 != 0 {
					p21 = i32(65536)
				}
				p22 := i32(0)
				if v10 == i32(1) {
					p22 = i32(256)
				}
				var p23 int32
				if v9 == i32(1) {
					p23 = 1
				}
				v5 = p21 | (p22 | p23 ^ v5&i32(257))
			}
		l5:
			v3 = i32(0)
			if v2 != 0 {
				m.fn1533(v4+i32(8), v1, v2, v6)
				t24 := int32(load32(m.memory[int64(uint32(v4))+8:]))
				v8 = t24
				if v8 == 0 {
					goto l8
				}
				t25 := int32(load32(m.memory[int64(uint32(v8))+4:]))
				v2 = t25
				t26 := int32(load32(m.memory[uint32(v8):]))
				v3 = t26
				goto l8
			}
			goto l8
		}
		store32(m.memory[int64(uint32(v4))+92:], uint32(i32(71)))
		store32(m.memory[int64(uint32(v4))+88:], uint32(v4+i32(80)))
		m.fn73(v4+i32(24), i32(1049807), v4+i32(88))
		store32(m.memory[int64(uint32(v4))+36:], uint32(i32(-1)))
		t5 := int32(load32(m.memory[int64(uint32(v4))+24:]))
		v3 = t5
		goto l4
	}
l2:
	m.memory[int64(uint32(v4))+28] = byte(i32(0))
	v3 = i32(-1)
l4:
	t27 := int32(load32(m.memory[int64(uint32(v4))+48:]))
	t28 := int32(load32(m.memory[int64(uint32(v4))+52:]))
	m.fn56(t27, t28)
	{
		if v3 == i32(-1) {
			goto l9
		}
		t29 := int64(load64(m.memory[int64(uint32(v4))+29:]))
		store64(m.memory[int64(uint32(v0))+5:], uint64(t29))
		t30 := int64(load64(m.memory[int64(uint32(v4))+37:]))
		store64(m.memory[int64(uint32(v0))+13:], uint64(t30))
		t31 := int32(load32(m.memory[int64(uint32(v4))+44:]))
		store32(m.memory[int64(uint32(v0))+20:], uint32(t31))
		t32 := int32(m.memory[int64(uint32(v4))+28])
		m.memory[int64(uint32(v0))+4] = byte(t32)
		goto l10
	}
l9:
	store16(m.memory[int64(uint32(v0))+4:], uint16(v5))
	m.memory[uint32(v0+i32(6))] = byte(int32(uint32(v5) >> 16))
l10:
	store32(m.memory[uint32(v0):], uint32(v3))
	m.g0 = v4 + i32(96)
}
func fn1537(v0, v1 int32) int32 {
	v0 = v1 ^ v0
	t1 := v0&i32(1) | v1&i32(0x1000000)
	p0 := i32(0)
	if int32(uint32(v0&i32(65536))>>16) != 0 {
		p0 = i32(65536)
	}
	t3 := t1 | p0
	p2 := i32(0)
	if int32(uint32(v0&i32(256))>>8) != 0 {
		p2 = i32(256)
	}
	return t3 | p2
}
func (m *Module) fn1538(v0, v1, v2 int32) {
	var v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14 int32
	t0 := m.g0
	v3 = t0 - i32(192)
	m.g0 = v3
	t1 := int32(load32(m.memory[int64(uint32(v2))+32:]))
	v4 = t1
	t2 := int32(load32(m.memory[int64(uint32(v2))+28:]))
	t3 := v3
	v2 = t2
	store32(m.memory[int64(uint32(t3))+40:], uint32(v2))
	store32(m.memory[int64(uint32(v3))+44:], uint32(v2+v4*i32(44)))
	v5 = v3 + i32(76) + i32(7)
	v6 = v3 + i32(76) + i32(4)
	v7 = v3 + i32(144) + i32(7)
	v8 = v3 + i32(144) + i32(4)
	v9 = v3 + i32(132)
	{
	l2:
		{
			{
				{
					t4 := m.fn904(v3 + i32(40))
					v2 = t4
					if v2 == 0 {
						store32(m.memory[uint32(v0):], uint32(i32(-1)))
						goto l10
					}
					t5 := m.fn847(v2, i32(1073848), i32(59), i32(1077491), i32(16))
					if t5 != 0 {
						t119 := int32(load32(m.memory[uint32(v2+i32(28)):]))
						t120 := int32(load32(m.memory[uint32(v2+i32(32)):]))
						t121 := m.fn1531(t119, t120)
						v2 = t121
						if v2 == 0 {
							goto l2
						}
						m.fn1538(v3+i32(144), v1, v2)
						t122 := int32(load32(m.memory[int64(uint32(v3))+144:]))
						v2 = t122
						if v2 == i32(-1) {
							goto l2
						}
						t123 := int32(load32(m.memory[int64(uint32(v3))+164:]))
						store32(m.memory[int64(uint32(v0))+20:], uint32(t123))
						t124 := int64(load64(m.memory[int64(uint32(v3))+156:]))
						store64(m.memory[int64(uint32(v0))+12:], uint64(t124))
						t125 := int64(load64(m.memory[int64(uint32(v3))+148:]))
						store64(m.memory[int64(uint32(v0))+4:], uint64(t125))
						store32(m.memory[uint32(v0):], uint32(v2))
						goto l10
					}
					t6 := int32(load32(m.memory[int64(uint32(v2))+36:]))
					v4 = t6
					if v4 == 0 {
						goto l2
					}
					t7 := int32(load32(m.memory[int64(uint32(v2))+40:]))
					t8 := m.fn1337(v4+i32(8), t7, i32(1072544), i32(60))
					if t8 != 0 {
						goto l2
					}
					t9 := int32(load32(m.memory[int64(uint32(v2))+4:]))
					v4 = t9
					t10 := int32(load32(m.memory[int64(uint32(v2))+8:]))
					t11 := v4
					v10 = t10
					t12 := m.fn15(t11, v10, i32(1073735), i32(3))
					if t12 != 0 {
						goto l2
					}
					t13 := m.fn15(v4, v10, i32(1072195), i32(1))
					if t13 != 0 {
						{
							{
								t20 := int32(load32(m.memory[uint32(v2+i32(28)):]))
								t21 := int32(load32(m.memory[uint32(v2+i32(32)):]))
								t22 := m.fn886(t20, t21, i32(1072544), i32(60), i32(1073717), i32(3))
								v4 = t22
								if v4 == 0 {
									goto l11
								}
								v10 = i32(0)
								{
									v11 = v4 + i32(28)
									t23 := int32(load32(m.memory[uint32(v11):]))
									v12 = t23
									t24 := v12
									v13 = v4 + i32(32)
									t25 := int32(load32(m.memory[uint32(v13):]))
									v4 = t25
									t26 := m.fn886(t24, v4, i32(1072544), i32(60), i32(1077702), i32(6))
									v14 = t26
									if v14 == 0 {
										goto l12
									}
									t27 := int32(load32(m.memory[uint32(v14+i32(16)):]))
									t28 := int32(load32(m.memory[uint32(v14+i32(20)):]))
									m.fn1046(v3, t27, t28, i32(1072544), i32(60), i32(1073156), i32(3))
									t29 := int32(load32(m.memory[uint32(v3):]))
									v14 = t29
									if v14 == 0 {
										goto l12
									}
									t30 := int32(load32(m.memory[int64(uint32(v1))+40:]))
									t31 := int32(load32(m.memory[int64(uint32(t30))+36:]))
									t32 := int32(load32(m.memory[int64(uint32(v3))+4:]))
									m.fn1536(v3+i32(144), t31, v14, t32)
									t33 := int32(load16(m.memory[int64(uint32(v3))+148:]))
									t34 := int32(m.memory[uint32(v3+i32(144)+i32(6))])
									v10 = t33 | t34<<16
									{
										t35 := int32(load32(m.memory[int64(uint32(v3))+144:]))
										v4 = t35
										if v4 == i32(-1) {
											goto l13
										}
										t36 := int32(m.memory[int64(uint32(v7))+16])
										m.memory[int64(uint32(v5))+16] = byte(t36)
										t37 := int64(load64(m.memory[int64(uint32(v7))+8:]))
										store64(m.memory[int64(uint32(v5))+8:], uint64(t37))
										t38 := int64(load64(m.memory[uint32(v7):]))
										store64(m.memory[uint32(v5):], uint64(t38))
										store16(m.memory[int64(uint32(v3))+80:], uint16(v10))
										m.memory[uint32(v3+i32(82))] = byte(int32(uint32(v10) >> 16))
										goto l14
									}
								l13:
									t39 := int32(load32(m.memory[uint32(v13):]))
									v4 = t39
									t40 := int32(load32(m.memory[uint32(v11):]))
									v12 = t40
								}
							l12:
								t41 := int32(load32(m.memory[int64(uint32(v1))+36:]))
								t42 := fn1537(v10, t41)
								v10 = t42
								t43 := m.fn1187(v12, v4)
								t44 := fn1319(t43, v10)
								v4 = t44
								goto l15
							}
						l11:
							t45 := int32(load32(m.memory[int64(uint32(v1))+36:]))
							v4 = t45
						}
					l15:
						m.fn1543(v3+i32(76), v1, v2, v4)
						t46 := int32(load32(m.memory[int64(uint32(v3))+76:]))
						v4 = t46
						if v4 == i32(-1) {
							goto l2
						}
						goto l14
					}
					t14 := m.fn15(v4, v10, i32(1077614), i32(9))
					if t14 != 0 {
						t47 := int32(load32(m.memory[uint32(v2+i32(16)):]))
						t48 := v3 + i32(16)
						v10 = t47
						t49 := int32(load32(m.memory[uint32(v2+i32(20)):]))
						t50 := v10
						v12 = t49
						m.fn845(t48, t50, v12, i32(1073159), i32(67), i32(1073226), i32(2))
						{
							{
								t51 := int32(load32(m.memory[int64(uint32(v3))+16:]))
								v4 = t51
								if v4 == 0 {
									goto l16
								}
								t52 := int32(load32(m.memory[int64(uint32(v1))+40:]))
								t53 := int32(load32(m.memory[int64(uint32(v3))+20:]))
								t54 := m.fn846(t52, v4, t53)
								v4 = t54
								if v4 != 0 {
									goto l17
								}
							}
						l16:
							m.fn1046(v3+i32(8), v10, v12, i32(1072544), i32(60), i32(1077608), i32(6))
							{
								t55 := int32(load32(m.memory[int64(uint32(v3))+8:]))
								v4 = t55
								if v4 == 0 {
									store32(m.memory[int64(uint32(v3))+48:], uint32(i32(-1)))
									goto l19
								}
								t56 := int32(load32(m.memory[int64(uint32(v3))+12:]))
								m.fn51(v8, v4, t56)
								store32(m.memory[int64(uint32(v3))+144:], uint32(i32(2)))
								t57 := int64(load64(m.memory[int64(uint32(v3))+152:]))
								store64(m.memory[int64(uint32(v3))+56:], uint64(t57))
								t58 := int64(load64(m.memory[int64(uint32(v3))+144:]))
								store64(m.memory[int64(uint32(v3))+48:], uint64(t58))
								goto l19
							}
						l17:
							t59 := int32(m.memory[int64(uint32(v4))+24])
							t60 := int32(load32(m.memory[int64(uint32(v4))+4:]))
							t61 := int32(load32(m.memory[int64(uint32(v4))+8:]))
							m.fn1494(v3+i32(48), t59, t60, t61)
						}
					l19:
						store32(m.memory[int64(uint32(v3))+108:], uint32(i32(0)))
						store64(m.memory[int64(uint32(v3))+100:], uint64(i64(0x400000000)))
						store64(m.memory[int64(uint32(v3))+92:], uint64(i64(4)))
						store64(m.memory[int64(uint32(v3))+84:], uint64(i64(0)))
						store64(m.memory[int64(uint32(v3))+76:], uint64(i64(0x400000000)))
						t62 := int64(load64(m.memory[int64(uint32(v1))+36:]))
						store64(m.memory[int64(uint32(v3))+112:], uint64(t62))
						m.fn1538(v3+i32(144), v3+i32(76), v2)
						t63 := int32(load32(m.memory[int64(uint32(v3))+144:]))
						v2 = t63
						if v2 == i32(-1) {
							memory_copy(m.memory, uint32(v3+i32(144)), uint32(v3+i32(76)), uint32(i32(44)))
							m.fn1540(v3+i32(64), v3+i32(144))
							m.fn1545(v3+i32(120), v3+i32(64))
							t105 := int32(load32(m.memory[int64(uint32(v3))+120:]))
							v4 = t105
							t106 := int32(load32(m.memory[int64(uint32(v3))+124:]))
							v2 = t106
							t107 := int32(load32(m.memory[int64(uint32(v3))+128:]))
							v10 = t107
							t108 := int32(load32(m.memory[int64(uint32(v9))+8:]))
							store32(m.memory[int64(uint32(v3))+72:], uint32(t108))
							t109 := int64(load64(m.memory[uint32(v9):]))
							store64(m.memory[int64(uint32(v3))+64:], uint64(t109))
							{
								t110 := int32(load32(m.memory[int64(uint32(v3))+48:]))
								if t110 == i32(-1) {
									store32(m.memory[int64(uint32(v3))+128:], uint32(v4))
									store32(m.memory[int64(uint32(v3))+120:], uint32(v2))
									t113 := v3
									t114 := v2
									v4 = v10 * i32(28)
									v14 = t114 + v4
									store32(m.memory[int64(uint32(t113))+132:], uint32(v14))
									v10 = v2 + i32(28)
								l26:
									{
										if v4 == 0 {
											goto l24
										}
										t115 := int32(load32(m.memory[uint32(v2):]))
										v12 = t115
										if v12 != i32(-1) {
											t116 := int64(load64(m.memory[int64(uint32(v2))+4:]))
											store64(m.memory[uint32(v8):], uint64(t116))
											t117 := int64(load64(m.memory[int64(uint32(v2))+12:]))
											store64(m.memory[int64(uint32(v8))+8:], uint64(t117))
											t118 := int64(load64(m.memory[int64(uint32(v2))+20:]))
											store64(m.memory[int64(uint32(v8))+16:], uint64(t118))
											store32(m.memory[int64(uint32(v3))+144:], uint32(v12))
											m.fn1544(v1, v3+i32(144))
											v4 = v4 + i32(-28)
											v10 = v10 + i32(28)
											v2 = v2 + i32(28)
											goto l26
										}
										v14 = v10
									}
								l24:
									store32(m.memory[int64(uint32(v3))+124:], uint32(v14))
									m.fn900(v3 + i32(120))
									goto l23
								}
								t111 := int64(load64(m.memory[int64(uint32(v3))+56:]))
								store64(m.memory[int64(uint32(v3))+152:], uint64(t111))
								t112 := int64(load64(m.memory[int64(uint32(v3))+48:]))
								store64(m.memory[int64(uint32(v3))+144:], uint64(t112))
								store32(m.memory[int64(uint32(v3))+168:], uint32(v10))
								store32(m.memory[int64(uint32(v3))+164:], uint32(v2))
								store32(m.memory[int64(uint32(v3))+160:], uint32(v4))
								m.fn1544(v1, v3+i32(144))
								goto l23
							}
						}
						t64 := int32(load32(m.memory[int64(uint32(v3))+164:]))
						store32(m.memory[int64(uint32(v0))+20:], uint32(t64))
						t65 := int64(load64(m.memory[int64(uint32(v3))+156:]))
						store64(m.memory[int64(uint32(v0))+12:], uint64(t65))
						t66 := int64(load64(m.memory[int64(uint32(v3))+148:]))
						store64(m.memory[int64(uint32(v0))+4:], uint64(t66))
						store32(m.memory[uint32(v0):], uint32(v2))
						m.fn1539(v3 + i32(76))
						m.fn1419(v3 + i32(48))
						goto l10
					}
					t15 := m.fn15(v4, v10, i32(1077623), i32(9))
					if t15 != 0 {
						t83 := int32(load32(m.memory[uint32(v2+i32(16)):]))
						t84 := int32(load32(m.memory[uint32(v2+i32(20)):]))
						m.fn1046(v3+i32(24), t83, t84, i32(1072544), i32(60), i32(1077697), i32(5))
						t85 := int32(load32(m.memory[int64(uint32(v3))+24:]))
						t86 := v3 + i32(64)
						v4 = t85
						p87 := i32(1)
						if v4 != 0 {
							p87 = v4
						}
						t88 := int32(load32(m.memory[int64(uint32(v3))+28:]))
						p89 := i32(0)
						if v4 != 0 {
							p89 = t88
						}
						m.fn51(t86, p87, p89)
						store32(m.memory[int64(uint32(v3))+108:], uint32(i32(0)))
						store64(m.memory[int64(uint32(v3))+100:], uint64(i64(0x400000000)))
						store64(m.memory[int64(uint32(v3))+92:], uint64(i64(4)))
						store64(m.memory[int64(uint32(v3))+84:], uint64(i64(0)))
						store64(m.memory[int64(uint32(v3))+76:], uint64(i64(0x400000000)))
						t90 := int64(load64(m.memory[int64(uint32(v1))+36:]))
						store64(m.memory[int64(uint32(v3))+112:], uint64(t90))
						m.fn1538(v3+i32(144), v3+i32(76), v2)
						{
							t91 := int32(load32(m.memory[int64(uint32(v3))+144:]))
							v2 = t91
							if v2 == i32(-1) {
								memory_copy(m.memory, uint32(v3+i32(144)), uint32(v3+i32(76)), uint32(i32(44)))
								m.fn1540(v3+i32(48), v3+i32(144))
								m.fn1545(v3+i32(120), v3+i32(48))
								t97 := int32(load32(m.memory[int64(uint32(v3))+128:]))
								store32(m.memory[int64(uint32(v3))+56:], uint32(t97))
								t98 := int64(load64(m.memory[int64(uint32(v3))+120:]))
								store64(m.memory[int64(uint32(v3))+48:], uint64(t98))
								t99 := int32(load32(m.memory[int64(uint32(v9))+8:]))
								store32(m.memory[int64(uint32(v3))+152:], uint32(t99))
								t100 := int64(load64(m.memory[uint32(v9):]))
								store64(m.memory[int64(uint32(v3))+144:], uint64(t100))
								t101 := int32(load32(m.memory[int64(uint32(v3))+68:]))
								t102 := v1
								v2 = t101
								t103 := int32(load32(m.memory[int64(uint32(v3))+72:]))
								m.fn1546(t102, v2, t103, v3+i32(48))
								m.fn1547(v1, v3+i32(144))
								t104 := int32(load32(m.memory[int64(uint32(v3))+64:]))
								m.fn16(t104, v2)
								goto l2
							}
							t92 := int32(load32(m.memory[int64(uint32(v3))+164:]))
							store32(m.memory[int64(uint32(v0))+20:], uint32(t92))
							t93 := int64(load64(m.memory[int64(uint32(v3))+156:]))
							store64(m.memory[int64(uint32(v0))+12:], uint64(t93))
							t94 := int64(load64(m.memory[int64(uint32(v3))+148:]))
							store64(m.memory[int64(uint32(v0))+4:], uint64(t94))
							store32(m.memory[uint32(v0):], uint32(v2))
							m.fn1539(v3 + i32(76))
							t95 := int32(load32(m.memory[int64(uint32(v3))+64:]))
							t96 := int32(load32(m.memory[int64(uint32(v3))+68:]))
							m.fn16(t95, t96)
							goto l10
						}
					}
					t16 := m.fn15(v4, v10, i32(1077632), i32(13))
					if t16 != 0 {
						t78 := int32(load32(m.memory[uint32(v2+i32(16)):]))
						t79 := int32(load32(m.memory[uint32(v2+i32(20)):]))
						m.fn1046(v3+i32(32), t78, t79, i32(1072544), i32(60), i32(1073713), i32(4))
						t80 := int32(load32(m.memory[int64(uint32(v3))+36:]))
						v2 = t80
						t81 := int32(load32(m.memory[int64(uint32(v3))+32:]))
						v4 = t81
						if v4 == 0 {
							goto l2
						}
						t82 := m.fn1337(v4, v2, i32(1077690), i32(7))
						if t82 == 0 {
							goto l2
						}
						m.fn51(v8, v4, v2)
						store32(m.memory[int64(uint32(v3))+144:], uint32(i32(6)))
						m.fn1544(v1, v3+i32(144))
						goto l2
					}
					t17 := m.fn15(v4, v10, i32(1077645), i32(3))
					if t17 != 0 {
						t71 := int32(load32(m.memory[uint32(v2+i32(28)):]))
						t72 := int32(load32(m.memory[uint32(v2+i32(32)):]))
						t73 := m.fn886(t71, t72, i32(1072544), i32(60), i32(1077680), i32(10))
						v2 = t73
						if v2 == 0 {
							goto l2
						}
						m.fn1538(v3+i32(144), v1, v2)
						t74 := int32(load32(m.memory[int64(uint32(v3))+144:]))
						v2 = t74
						if v2 == i32(-1) {
							goto l2
						}
						t75 := int32(load32(m.memory[int64(uint32(v3))+164:]))
						store32(m.memory[int64(uint32(v0))+20:], uint32(t75))
						t76 := int64(load64(m.memory[int64(uint32(v3))+156:]))
						store64(m.memory[int64(uint32(v0))+12:], uint64(t76))
						t77 := int64(load64(m.memory[int64(uint32(v3))+148:]))
						store64(m.memory[int64(uint32(v0))+4:], uint64(t77))
						store32(m.memory[uint32(v0):], uint32(v2))
						goto l10
					}
					t18 := m.fn15(v4, v10, i32(1077648), i32(8))
					if t18 != 0 {
						goto l8
					}
					t19 := m.fn15(v4, v10, i32(1077656), i32(3))
					if t19 == 0 {
						t67 := m.fn15(v4, v10, i32(1077659), i32(3))
						if t67 != 0 {
							goto l8
						}
						t68 := m.fn15(v4, v10, i32(1077662), i32(3))
						if t68 != 0 {
							goto l8
						}
						t69 := m.fn15(v4, v10, i32(1077665), i32(6))
						if t69 != 0 {
							goto l8
						}
						t70 := m.fn15(v4, v10, i32(1077671), i32(9))
						if t70 == 0 {
							goto l2
						}
						goto l8
					}
					goto l8
				}
			l23:
				m.fn1547(v1, v3+i32(64))
				goto l2
			l14:
				t126 := int32(load32(m.memory[int64(uint32(v6))+16:]))
				store32(m.memory[int64(uint32(v0))+20:], uint32(t126))
				t127 := int64(load64(m.memory[int64(uint32(v6))+8:]))
				store64(m.memory[int64(uint32(v0))+12:], uint64(t127))
				t128 := int64(load64(m.memory[uint32(v6):]))
				store64(m.memory[int64(uint32(v0))+4:], uint64(t128))
				store32(m.memory[uint32(v0):], uint32(v4))
				goto l10
			}
		l8:
			m.fn1538(v3+i32(144), v1, v2)
			t129 := int32(load32(m.memory[int64(uint32(v3))+144:]))
			v2 = t129
			if v2 == i32(-1) {
				goto l2
			}
		}
		t130 := int32(load32(m.memory[int64(uint32(v3))+164:]))
		store32(m.memory[int64(uint32(v0))+20:], uint32(t130))
		t131 := int64(load64(m.memory[int64(uint32(v3))+156:]))
		store64(m.memory[int64(uint32(v0))+12:], uint64(t131))
		t132 := int64(load64(m.memory[int64(uint32(v3))+148:]))
		store64(m.memory[int64(uint32(v0))+4:], uint64(t132))
		store32(m.memory[uint32(v0):], uint32(v2))
	}
l10:
	m.g0 = v3 + i32(192)
}
func (m *Module) fn1539(v0 int32) {
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
	m.fn1548(v3)
	v3 = v3 + i32(16)
	goto l1
l0:
	t2 := int32(load32(m.memory[uint32(v0):]))
	m.fn419(t2, v2)
	m.fn894(v0 + i32(12))
	m.fn1444(v0 + i32(24))
}
func (m *Module) fn1540(v0, v1 int32) {
	var v2, v3, v4, v5, v6, v7, v8, v9, v10, v11 int32
	t0 := m.g0
	v2 = t0 - i32(48)
	m.g0 = v2
	v3 = v1 + i32(24)
	v4 = v2 + i32(24)
l3:
	{
		t1 := int32(load32(m.memory[int64(uint32(v1))+32:]))
		v5 = t1
		if v5 == 0 {
			goto l0
		}
		t2 := v1
		v5 = v5 + i32(-1)
		store32(m.memory[int64(uint32(t2))+32:], uint32(v5))
		t3 := int32(load32(m.memory[int64(uint32(v1))+28:]))
		v6 = t3 + v5*i32(28)
		t4 := int32(load32(m.memory[uint32(v6):]))
		v7 = t4
		if v7 == i32(-1) {
			goto l0
		}
		t5 := int32(load32(m.memory[int64(uint32(v6))+4:]))
		v8 = t5
		t6 := int32(load32(m.memory[int64(uint32(v6))+20:]))
		v9 = t6
		t7 := int32(load32(m.memory[int64(uint32(v6))+16:]))
		v5 = t7
		t8 := int32(load32(m.memory[int64(uint32(v6))+12:]))
		store32(m.memory[int64(uint32(v2))+12:], uint32(t8))
		store32(m.memory[int64(uint32(v2))+4:], uint32(v5))
		t9 := v2
		t10 := v5
		v6 = v9 * i32(28)
		v10 = t10 + v6
		store32(m.memory[int64(uint32(t9))+16:], uint32(v10))
		v9 = v5 + i32(28)
	l4:
		{
			if v6 == 0 {
				goto l1
			}
			t11 := int32(load32(m.memory[uint32(v5):]))
			v11 = t11
			if v11 != i32(-1) {
				t12 := int64(load64(m.memory[int64(uint32(v5))+4:]))
				store64(m.memory[uint32(v4):], uint64(t12))
				t13 := int64(load64(m.memory[int64(uint32(v5))+12:]))
				store64(m.memory[int64(uint32(v4))+8:], uint64(t13))
				t14 := int64(load64(m.memory[int64(uint32(v5))+20:]))
				store64(m.memory[int64(uint32(v4))+16:], uint64(t14))
				store32(m.memory[int64(uint32(v2))+20:], uint32(v11))
				m.fn1544(v1, v2+i32(20))
				v6 = v6 + i32(-28)
				v9 = v9 + i32(28)
				v5 = v5 + i32(28)
				goto l4
			}
			v10 = v9
		}
	l1:
		store32(m.memory[int64(uint32(v2))+8:], uint32(v10))
		m.fn900(v2 + i32(4))
		m.fn16(v7, v8)
		goto l3
	}
l0:
	v5 = v1 + i32(12)
	{
		{
			t15 := int32(load32(m.memory[int64(uint32(v1))+20:]))
			if t15 == 0 {
				goto l5
			}
			store32(m.memory[int64(uint32(v2))+20:], uint32(i32(0)))
			t16 := int32(load32(m.memory[int64(uint32(v5))+8:]))
			store32(m.memory[int64(uint32(v2))+32:], uint32(t16))
			t17 := int64(load64(m.memory[uint32(v5):]))
			store64(m.memory[int64(uint32(v2))+24:], uint64(t17))
			m.fn832(v1, v2+i32(20))
			t18 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			store32(m.memory[int64(uint32(v0))+8:], uint32(t18))
			t19 := int64(load64(m.memory[uint32(v1):]))
			store64(m.memory[uint32(v0):], uint64(t19))
			goto l6
		}
	l5:
		t20 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		store32(m.memory[int64(uint32(v0))+8:], uint32(t20))
		t21 := int64(load64(m.memory[uint32(v1):]))
		store64(m.memory[uint32(v0):], uint64(t21))
		m.fn894(v5)
	}
l6:
	m.fn1444(v3)
	m.g0 = v2 + i32(48)
}
func (m *Module) fn1541(v0, v1 int32) {
	var v2, v3, v4, v5, v6 int32
	t0 := m.g0
	v2 = t0 - i32(64)
	m.g0 = v2
	store32(m.memory[int64(uint32(v2))+12:], uint32(i32(0)))
	store64(m.memory[int64(uint32(v2))+4:], uint64(i64(0x800000000)))
	t1 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	v3 = t1
	t2 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	v4 = t2
	t3 := int32(load32(m.memory[uint32(v1):]))
	store32(m.memory[int64(uint32(v2))+24:], uint32(t3))
	store32(m.memory[int64(uint32(v2))+16:], uint32(v4))
	t4 := v2
	t5 := v4
	v1 = v3 << 4
	v5 = t5 + v1
	store32(m.memory[int64(uint32(t4))+28:], uint32(v5))
	v6 = v2 + i32(32) | i32(4)
l5:
	{
		{
			if v1 == 0 {
				goto l0
			}
			v3 = v4 + i32(4)
			{
				t6 := int32(load32(m.memory[uint32(v4):]))
				switch t6 {
				case 0:
					goto l1
				default:
					t9 := int32(load32(m.memory[int64(uint32(v3))+8:]))
					store32(m.memory[int64(uint32(v2))+40:], uint32(t9))
					t10 := int64(load64(m.memory[uint32(v3):]))
					store64(m.memory[int64(uint32(v2))+32:], uint64(t10))
					m.fn1271(v2+i32(4), v2+i32(32))
					goto l4
				case 2:
					v5 = v4 + i32(16)
				}
			}
		l0:
			store32(m.memory[int64(uint32(v2))+20:], uint32(v5))
			m.fn1542(v2 + i32(16))
			t7 := int32(load32(m.memory[int64(uint32(v2))+12:]))
			store32(m.memory[int64(uint32(v0))+8:], uint32(t7))
			t8 := int64(load64(m.memory[int64(uint32(v2))+4:]))
			store64(m.memory[uint32(v0):], uint64(t8))
			m.g0 = v2 + i32(64)
			return
		}
	l1:
		t11 := int32(load32(m.memory[int64(uint32(v3))+8:]))
		store32(m.memory[int64(uint32(v6))+8:], uint32(t11))
		t12 := int64(load64(m.memory[uint32(v3):]))
		store64(m.memory[uint32(v6):], uint64(t12))
		store32(m.memory[int64(uint32(v2))+32:], uint32(i32(-0x80000000)))
		m.fn338(v2+i32(4), v2+i32(32))
	}
l4:
	v4 = v4 + i32(16)
	v1 = v1 + i32(-16)
	goto l5
}
func (m *Module) fn1542(v0 int32) {
	var v1, v2 int32
	t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
	t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	v1 = t1
	v2 = int32(uint32(t0-v1) >> 4)
l1:
	if v2 == 0 {
		goto l0
	}
	v2 = v2 + i32(-1)
	m.fn1548(v1)
	v1 = v1 + i32(16)
	goto l1
l0:
	t2 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	t3 := int32(load32(m.memory[uint32(v0):]))
	m.fn419(t2, t3)
}
func (m *Module) fn1543(v0, v1, v2, v3 int32) {
	var v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14 int32
	var v15 int64
	var v16, v17, v18, v19 int32
	var v20 int64
	var v21, v22, v23 int32
	t0 := m.g0
	v4 = t0 - i32(432)
	m.g0 = v4
	t1 := int32(load32(m.memory[int64(uint32(v2))+32:]))
	v5 = t1
	t2 := int32(load32(m.memory[int64(uint32(v2))+28:]))
	t3 := v4
	v2 = t2
	store32(m.memory[int64(uint32(t3))+152:], uint32(v2))
	store32(m.memory[int64(uint32(v4))+156:], uint32(v2+v5*i32(44)))
	v6 = v1 + i32(24)
	v7 = v4 + i32(316) + i32(8)
	v8 = v4 + i32(316) + i32(4)
	v9 = v4 + i32(316) + i32(16)
	v10 = v4 + i32(344) + i32(4)
	v11 = v4 + i32(224) + i32(4)
	v12 = v4 + i32(388) + i32(4)
l2:
	{
		{
			t4 := m.fn904(v4 + i32(152))
			v2 = t4
			if v2 == 0 {
				store32(m.memory[uint32(v0):], uint32(i32(-1)))
				goto l6
			}
			t5 := m.fn847(v2, i32(1073848), i32(59), i32(1077491), i32(16))
			if t5 != 0 {
				t72 := int32(load32(m.memory[uint32(v2+i32(28)):]))
				t73 := int32(load32(m.memory[uint32(v2+i32(32)):]))
				t74 := m.fn1531(t72, t73)
				v2 = t74
				if v2 == 0 {
					goto l2
				}
				m.fn1543(v4+i32(316), v1, v2, v3)
				t75 := int32(load32(m.memory[int64(uint32(v4))+316:]))
				v2 = t75
				if v2 == i32(-1) {
					goto l2
				}
				t76 := int32(load32(m.memory[int64(uint32(v4))+336:]))
				store32(m.memory[int64(uint32(v0))+20:], uint32(t76))
				t77 := int64(load64(m.memory[int64(uint32(v4))+328:]))
				store64(m.memory[int64(uint32(v0))+12:], uint64(t77))
				t78 := int64(load64(m.memory[int64(uint32(v4))+320:]))
				store64(m.memory[int64(uint32(v0))+4:], uint64(t78))
				store32(m.memory[uint32(v0):], uint32(v2))
				goto l6
			}
			t6 := int32(load32(m.memory[int64(uint32(v2))+36:]))
			v5 = t6
			if v5 == 0 {
				goto l2
			}
			t7 := int32(load32(m.memory[int64(uint32(v2))+40:]))
			t8 := m.fn15(v5+i32(8), t7, i32(1072544), i32(60))
			if t8 == 0 {
				goto l2
			}
			t9 := int32(load32(m.memory[int64(uint32(v2))+4:]))
			v5 = t9
			t10 := int32(load32(m.memory[int64(uint32(v2))+8:]))
			t11 := v5
			v13 = t10
			t12 := m.fn15(t11, v13, i32(1072196), i32(1))
			if t12 != 0 {
				t15 := int32(load32(m.memory[uint32(v2+i32(16)):]))
				t16 := int32(load32(m.memory[uint32(v2+i32(20)):]))
				m.fn845(v4+i32(24), t15, t16, i32(1282584), i32(36), i32(1077603), i32(5))
				t17 := int32(load32(m.memory[int64(uint32(v4))+24:]))
				t18 := int32(load32(m.memory[int64(uint32(v4))+28:]))
				t19 := m.fn848(t17, t18, i32(1072206), i32(8))
				v5 = t19
				t20 := int32(load32(m.memory[uint32(v2+i32(28)):]))
				t21 := int32(load32(m.memory[uint32(v2+i32(32)):]))
				m.fn864(v4+i32(296), t20, t21)
				t22 := int32(load32(m.memory[int64(uint32(v4))+304:]))
				v2 = t22
				t23 := int32(load32(m.memory[int64(uint32(v4))+300:]))
				v14 = t23
				if v5 == 0 {
					store32(m.memory[int64(uint32(v4))+332:], uint32(i32(0)))
					store32(m.memory[int64(uint32(v4))+324:], uint32(v14))
					store32(m.memory[int64(uint32(v4))+316:], uint32(v14))
					store32(m.memory[int64(uint32(v4))+320:], uint32(v2))
					t67 := v4
					v16 = v14 + v2
					store32(m.memory[int64(uint32(t67))+328:], uint32(v16))
					v17 = v14
				l18:
					{
						m.fn572(v4+i32(16), v7)
						t68 := int32(load32(m.memory[int64(uint32(v4))+16:]))
						v5 = t68
						t69 := int32(load32(m.memory[int64(uint32(v4))+20:]))
						v13 = t69
						v2 = v13 + i32(-9)
						if uint32(v2) > uint32(i32(23)) {
							goto l17
						}
						if i32_shl(i32(1), v2)&i32(8388627) == 0 {
							goto l17
						}
						t70 := int32(load32(m.memory[int64(uint32(v4))+324:]))
						v17 = t70
						t71 := int32(load32(m.memory[int64(uint32(v4))+328:]))
						v16 = t71
						goto l18
					}
				}
				v13 = v14
				goto l8
			}
			{
				t13 := m.fn15(v5, v13, i32(1077507), i32(3))
				if t13 != 0 {
					goto l4
				}
				t14 := m.fn15(v5, v13, i32(1077510), i32(4))
				if t14 == 0 {
					t24 := m.fn15(v5, v13, i32(1077123), i32(2))
					if t24 != 0 {
						store32(m.memory[int64(uint32(v4))+316:], uint32(i32(8)))
						m.fn1544(v1, v4+i32(316))
						goto l2
					}
					t25 := m.fn15(v5, v13, i32(1077514), i32(2))
					if t25 != 0 {
						store32(m.memory[int64(uint32(v4))+316:], uint32(i32(8)))
						m.fn1544(v1, v4+i32(316))
						goto l2
					}
					{
						t26 := m.fn15(v5, v13, i32(1077516), i32(17))
						if t26 != 0 {
							t63 := int32(load32(m.memory[uint32(v2+i32(16)):]))
							t64 := int32(load32(m.memory[uint32(v2+i32(20)):]))
							m.fn1046(v4+i32(32), t63, t64, i32(1072544), i32(60), i32(1073226), i32(2))
							t65 := int32(load32(m.memory[int64(uint32(v4))+36:]))
							v2 = t65
							t66 := int32(load32(m.memory[int64(uint32(v4))+32:]))
							v5 = t66
							if v5 == 0 {
								goto l2
							}
							store32(m.memory[int64(uint32(v4))+296:], uint32(v5))
							store32(m.memory[int64(uint32(v4))+300:], uint32(v2))
							store32(m.memory[int64(uint32(v4))+228:], uint32(i32(1)))
							store32(m.memory[int64(uint32(v4))+224:], uint32(v4+i32(296)))
							m.fn73(v8, i32(1048663), v4+i32(224))
							store32(m.memory[int64(uint32(v4))+316:], uint32(i32(7)))
							m.fn1544(v1, v4+i32(316))
							goto l2
						}
						{
							t27 := m.fn15(v5, v13, i32(1077533), i32(16))
							if t27 != 0 {
								t59 := int32(load32(m.memory[uint32(v2+i32(16)):]))
								t60 := int32(load32(m.memory[uint32(v2+i32(20)):]))
								m.fn1046(v4+i32(40), t59, t60, i32(1072544), i32(60), i32(1073226), i32(2))
								t61 := int32(load32(m.memory[int64(uint32(v4))+44:]))
								v2 = t61
								t62 := int32(load32(m.memory[int64(uint32(v4))+40:]))
								v5 = t62
								if v5 == 0 {
									goto l2
								}
								store32(m.memory[int64(uint32(v4))+296:], uint32(v5))
								store32(m.memory[int64(uint32(v4))+300:], uint32(v2))
								store32(m.memory[int64(uint32(v4))+228:], uint32(i32(1)))
								store32(m.memory[int64(uint32(v4))+224:], uint32(v4+i32(296)))
								m.fn73(v8, i32(1048668), v4+i32(224))
								store32(m.memory[int64(uint32(v4))+316:], uint32(i32(7)))
								m.fn1544(v1, v4+i32(316))
								goto l2
							}
							t28 := m.fn15(v5, v13, i32(1077549), i32(7))
							if t28 != 0 {
								goto l13
							}
							t29 := m.fn15(v5, v13, i32(1077556), i32(4))
							if t29 != 0 {
								goto l13
							}
							t30 := m.fn15(v5, v13, i32(1077452), i32(6))
							if t30 != 0 {
								goto l13
							}
							{
								t31 := m.fn15(v5, v13, i32(1077560), i32(7))
								if t31 != 0 {
									t41 := int32(load32(m.memory[uint32(v2+i32(16)):]))
									t42 := int32(load32(m.memory[uint32(v2+i32(20)):]))
									m.fn1046(v4+i32(144), t41, t42, i32(1072544), i32(60), i32(1077576), i32(11))
									t43 := int32(load32(m.memory[int64(uint32(v4))+148:]))
									v2 = t43
									t44 := int32(load32(m.memory[int64(uint32(v4))+144:]))
									v5 = t44
									if v5 == 0 {
										goto l2
									}
									t45 := m.fn15(v5, v2, i32(1077587), i32(5))
									if t45 != 0 {
										m.memory[int64(uint32(v4))+340] = byte(i32(0))
										store32(m.memory[int64(uint32(v4))+324:], uint32(i32(0)))
										store64(m.memory[int64(uint32(v4))+316:], uint64(i64(0x100000000)))
										store32(m.memory[int64(uint32(v4))+336:], uint32(i32(0)))
										store64(m.memory[int64(uint32(v4))+328:], uint64(i64(0x400000000)))
										m.fn1230(v6, v4+i32(316))
										goto l2
									}
									{
										t46 := m.fn15(v5, v2, i32(1077592), i32(8))
										if t46 != 0 {
											t57 := int32(load32(m.memory[int64(uint32(v1))+28:]))
											t58 := int32(load32(m.memory[int64(uint32(v1))+32:]))
											v2 = t58
											v13 = t57 + v2*i32(28)
											v5 = v13 + i32(-28)
											if v2 == 0 {
												goto l2
											}
											if v5 == 0 {
												goto l2
											}
											m.memory[uint32(v13+i32(-4))] = byte(i32(1))
											goto l2
										}
										t47 := m.fn15(v5, v2, i32(1077600), i32(3))
										if t47 == 0 {
											goto l2
										}
										t48 := int32(load32(m.memory[int64(uint32(v1))+32:]))
										v2 = t48
										if v2 == 0 {
											goto l2
										}
										t49 := v1
										v2 = v2 + i32(-1)
										store32(m.memory[int64(uint32(t49))+32:], uint32(v2))
										t50 := int32(load32(m.memory[int64(uint32(v1))+28:]))
										t51 := v4
										v2 = t50 + v2*i32(28)
										t52 := int64(load64(m.memory[int64(uint32(v2))+12:]))
										store64(m.memory[int64(uint32(t51))+160:], uint64(t52))
										t53 := int32(load32(m.memory[int64(uint32(v2))+20:]))
										store32(m.memory[int64(uint32(v4))+168:], uint32(t53))
										t54 := int64(load64(m.memory[int64(uint32(v2))+4:]))
										v15 = t54
										t55 := int32(load32(m.memory[int64(uint32(v2))+4:]))
										v5 = t55
										t56 := int32(load32(m.memory[uint32(v2):]))
										v2 = t56
										if v2 == i32(-1) {
											goto l2
										}
										m.fn1546(v1, int32(v15), int32(int64(uint64(v15)>>32)), v4+i32(160))
										m.fn16(v2, v5)
										goto l2
									}
								}
								t32 := m.fn15(v5, v13, i32(1077567), i32(9))
								if t32 == 0 {
									goto l2
								}
								t33 := int32(load32(m.memory[int64(uint32(v1))+28:]))
								t34 := int32(load32(m.memory[int64(uint32(v1))+32:]))
								v13 = t34
								v5 = t33 + v13*i32(28) + i32(-28)
								if v13 == 0 {
									goto l2
								}
								if v5 == 0 {
									goto l2
								}
								t35 := int32(load32(m.memory[uint32(v2+i32(28)):]))
								t36 := int32(load32(m.memory[uint32(v2+i32(32)):]))
								m.fn864(v4+i32(316), t35, t36)
								t37 := int32(load32(m.memory[int64(uint32(v4))+320:]))
								t38 := v5
								v2 = t37
								t39 := int32(load32(m.memory[int64(uint32(v4))+324:]))
								m.fn75(t38, v2, t39)
								t40 := int32(load32(m.memory[int64(uint32(v4))+316:]))
								m.fn16(t40, v2)
								goto l2
							}
						}
					}
				}
			}
		l4:
			m.fn51(v8, i32(1097368), i32(1))
			store32(m.memory[int64(uint32(v4))+316:], uint32(i32(3)))
			store32(m.memory[int64(uint32(v4))+332:], uint32(i32(0)))
			m.fn1544(v1, v4+i32(316))
			goto l2
		}
	l17:
		{
			if v13 == i32(-1) {
				goto l19
			}
			t79 := int32(load32(m.memory[int64(uint32(v4))+324:]))
			t80 := int32(load32(m.memory[int64(uint32(v4))+328:]))
			v18 = v16 - v17 + v5 + t79 - t80
			goto l22
		}
	l19:
		v18 = i32(0)
		v5 = i32(0)
	l22:
		{
			t81 := int32(load32(m.memory[int64(uint32(v4))+324:]))
			v16 = t81
			t82 := int32(load32(m.memory[int64(uint32(v4))+328:]))
			v17 = t82
			m.fn574(v4+i32(8), v7)
			t83 := int32(load32(m.memory[int64(uint32(v4))+8:]))
			v19 = t83
			t84 := int32(load32(m.memory[int64(uint32(v4))+12:]))
			v13 = t84
			v2 = v13 + i32(-9)
			if uint32(v2) > uint32(i32(23)) {
				goto l21
			}
			if i32_shl(i32(1), v2)&i32(8388627) != 0 {
				goto l22
			}
		}
	l21:
		{
			if v13 == i32(-1) {
				goto l23
			}
			t85 := int32(load32(m.memory[int64(uint32(v4))+324:]))
			t86 := int32(load32(m.memory[int64(uint32(v4))+328:]))
			v18 = v17 - v16 + v19 + t85 - t86
		}
	l23:
		v13 = v14 + v5
		v2 = v18 - v5
	l8:
		m.fn865(v4+i32(224), v13, v2)
		{
			{
				t87 := int32(load32(m.memory[int64(uint32(v4))+232:]))
				if t87 == 0 {
					goto l24
				}
				t88 := int32(load32(m.memory[int64(uint32(v4))+232:]))
				store32(m.memory[int64(uint32(v8))+8:], uint32(t88))
				t89 := int64(load64(m.memory[int64(uint32(v4))+224:]))
				store64(m.memory[uint32(v8):], uint64(t89))
				store32(m.memory[int64(uint32(v4))+332:], uint32(v3))
				store32(m.memory[int64(uint32(v4))+316:], uint32(i32(3)))
				m.fn1544(v1, v4+i32(316))
				goto l25
			}
		l24:
			t90 := int32(load32(m.memory[int64(uint32(v4))+224:]))
			t91 := int32(load32(m.memory[int64(uint32(v4))+228:]))
			m.fn16(t90, t91)
		}
	l25:
		t92 := int32(load32(m.memory[int64(uint32(v4))+296:]))
		m.fn16(t92, v14)
		goto l2
	}
l13:
	store32(m.memory[int64(uint32(v4))+184:], uint32(i32(0)))
	store64(m.memory[int64(uint32(v4))+176:], uint64(i64(0x400000000)))
	m.fn1550(v2, v4+i32(176))
	{
		{
			{
				{
					{
						{
							{
								{
									{
										{
											{
												{
													t93 := int32(load32(m.memory[int64(uint32(v4))+184:]))
													v5 = t93
													if v5 == 0 {
														{
															{
																v5 = v2 + i32(28)
																t97 := int32(load32(m.memory[uint32(v5):]))
																v2 = v2 + i32(32)
																t98 := int32(load32(m.memory[uint32(v2):]))
																t99 := m.fn1097(t97, t98, i32(1073501), i32(70), i32(1077419), i32(5))
																v13 = t99
																if v13 != 0 {
																	goto l28
																}
																v13 = i32(0)
																goto l29
															}
														l28:
															t100 := int32(load32(m.memory[uint32(v13+i32(16)):]))
															t101 := int32(load32(m.memory[uint32(v13+i32(20)):]))
															m.fn1046(v4+i32(136), t100, t101, i32(1073501), i32(70), i32(1073571), i32(5))
															t102 := int32(load32(m.memory[int64(uint32(v4))+140:]))
															v16 = t102
															t103 := int32(load32(m.memory[int64(uint32(v4))+136:]))
															v13 = t103
														}
													l29:
														m.fn1358(v4+i32(316), v13, v16)
														t104 := int32(load32(m.memory[int64(uint32(v4))+324:]))
														v14 = t104
														t105 := int32(load32(m.memory[int64(uint32(v4))+320:]))
														t106 := int32(load32(m.memory[int64(uint32(v4))+316:]))
														v17 = t106
														var p107 int32
														if v17 == i32(-1) {
															p107 = 1
														}
														v16 = p107
														p108 := t105
														if v16 != 0 {
															p108 = i32(1)
														}
														v13 = p108
														p109 := v17
														if v16 != 0 {
															p109 = i32(0)
														}
														v19 = p109
														{
															t110 := int32(load32(m.memory[uint32(v5):]))
															t111 := int32(load32(m.memory[uint32(v2):]))
															t112 := m.fn1097(t110, t111, i32(1073932), i32(54), i32(1077424), i32(5))
															v17 = t112
															if v17 == 0 {
																goto l30
															}
															t113 := int32(load32(m.memory[uint32(v17+i32(16)):]))
															t114 := int32(load32(m.memory[uint32(v17+i32(20)):]))
															m.fn845(v4+i32(128), t113, t114, i32(1073159), i32(67), i32(1073226), i32(2))
															t115 := int32(load32(m.memory[int64(uint32(v4))+128:]))
															v17 = t115
															if v17 != 0 {
																t142 := int32(load32(m.memory[int64(uint32(v1))+40:]))
																t143 := int32(load32(m.memory[int64(uint32(v4))+132:]))
																m.fn1551(v4+i32(316), t142, v17, t143)
																t144 := int32(load32(m.memory[int64(uint32(v4))+336:]))
																v16 = t144
																t145 := int32(load32(m.memory[int64(uint32(v4))+332:]))
																v17 = t145
																t146 := int32(load32(m.memory[int64(uint32(v4))+320:]))
																v2 = t146
																{
																	t147 := int32(load32(m.memory[int64(uint32(v4))+316:]))
																	v5 = t147
																	if v5 == i32(-1) {
																		if v2 == i32(-1) {
																			v2 = i32(0)
																			store32(m.memory[int64(uint32(v4))+236:], uint32(i32(0)))
																			store64(m.memory[int64(uint32(v4))+228:], uint64(i64(0x800000000)))
																			t156 := int64(load64(m.memory[int64(uint32(v4))+232:]))
																			v15 = t156
																			goto l48
																		}
																		t149 := int32(load32(m.memory[int64(uint32(v4))+324:]))
																		v14 = t149
																		store32(m.memory[int64(uint32(v4))+320:], uint32(v16))
																		store32(m.memory[int64(uint32(v4))+316:], uint32(v17))
																		m.fn1053(v4+i32(388), v17+i32(8), v16)
																		t150 := int32(load32(m.memory[int64(uint32(v4))+388:]))
																		if t150 != i32(-1) {
																			t212 := int32(load32(m.memory[int64(uint32(v4))+416:]))
																			t213 := int32(load32(m.memory[int64(uint32(v4))+420:]))
																			m.fn1362(v11, t212, t213)
																			m.fn1042(v4 + i32(388))
																			goto l47
																		}
																		{
																			t151 := int32(load32(m.memory[int64(uint32(v4))+392:]))
																			if t151 != i32(-0x7ffffffd) {
																				store32(m.memory[int64(uint32(v4))+236:], uint32(i32(0)))
																				store64(m.memory[int64(uint32(v4))+228:], uint64(i64(0x800000000)))
																				m.fn785(v12)
																				goto l47
																			}
																			t152 := int64(load64(m.memory[int64(uint32(v12))+16:]))
																			store64(m.memory[int64(uint32(v4))+240:], uint64(t152))
																			t153 := int64(load64(m.memory[int64(uint32(v12))+8:]))
																			store64(m.memory[int64(uint32(v4))+232:], uint64(t153))
																			t154 := int64(load64(m.memory[uint32(v12):]))
																			t155 := v4
																			v15 = t154
																			store64(m.memory[int64(uint32(t155))+224:], uint64(v15))
																			v5 = int32(v15)
																			goto l46
																		}
																	}
																	t148 := int64(load64(m.memory[int64(uint32(v4))+324:]))
																	v15 = t148
																	store32(m.memory[int64(uint32(v4))+244:], uint32(v16))
																	store32(m.memory[int64(uint32(v4))+240:], uint32(v17))
																	store64(m.memory[int64(uint32(v4))+232:], uint64(v15))
																	store32(m.memory[int64(uint32(v4))+228:], uint32(v2))
																	store32(m.memory[int64(uint32(v4))+224:], uint32(v5))
																	goto l42
																}
															}
														}
													l30:
														{
															t116 := int32(load32(m.memory[uint32(v5):]))
															t117 := int32(load32(m.memory[uint32(v2):]))
															t118 := m.fn1097(t116, t117, i32(1073986), i32(56), i32(1077429), i32(6))
															v17 = t118
															if v17 == 0 {
																goto l32
															}
															t119 := int32(load32(m.memory[uint32(v17+i32(16)):]))
															t120 := int32(load32(m.memory[uint32(v17+i32(20)):]))
															m.fn845(v4+i32(120), t119, t120, i32(1073159), i32(67), i32(1077435), i32(2))
															t121 := int32(load32(m.memory[int64(uint32(v4))+120:]))
															v17 = t121
															if v17 != 0 {
																t157 := int32(load32(m.memory[int64(uint32(v1))+40:]))
																t158 := int32(load32(m.memory[int64(uint32(v4))+124:]))
																m.fn1551(v4+i32(316), t157, v17, t158)
																t159 := int32(load32(m.memory[int64(uint32(v4))+336:]))
																v16 = t159
																t160 := int32(load32(m.memory[int64(uint32(v4))+332:]))
																v17 = t160
																t161 := int32(load32(m.memory[int64(uint32(v4))+320:]))
																v2 = t161
																{
																	t162 := int32(load32(m.memory[int64(uint32(v4))+316:]))
																	v5 = t162
																	if v5 == i32(-1) {
																		if v2 == i32(-1) {
																			v2 = i32(0)
																			store32(m.memory[int64(uint32(v4))+236:], uint32(i32(0)))
																			store64(m.memory[int64(uint32(v4))+228:], uint64(i64(0x800000000)))
																			t171 := int64(load64(m.memory[int64(uint32(v4))+232:]))
																			v15 = t171
																			goto l56
																		}
																		t164 := int32(load32(m.memory[int64(uint32(v4))+324:]))
																		v14 = t164
																		store32(m.memory[int64(uint32(v4))+320:], uint32(v16))
																		store32(m.memory[int64(uint32(v4))+316:], uint32(v17))
																		m.fn1053(v4+i32(344), v17+i32(8), v16)
																		t165 := int32(load32(m.memory[int64(uint32(v4))+344:]))
																		if t165 != i32(-1) {
																			t214 := int32(load32(m.memory[int64(uint32(v4))+372:]))
																			t215 := int32(load32(m.memory[int64(uint32(v4))+376:]))
																			m.fn1363(v11, t214, t215)
																			m.fn1042(v4 + i32(344))
																			goto l55
																		}
																		{
																			t166 := int32(load32(m.memory[int64(uint32(v4))+348:]))
																			if t166 != i32(-0x7ffffffd) {
																				store32(m.memory[int64(uint32(v4))+236:], uint32(i32(0)))
																				store64(m.memory[int64(uint32(v4))+228:], uint64(i64(0x800000000)))
																				m.fn785(v10)
																				goto l55
																			}
																			t167 := int64(load64(m.memory[int64(uint32(v10))+16:]))
																			store64(m.memory[int64(uint32(v4))+240:], uint64(t167))
																			t168 := int64(load64(m.memory[int64(uint32(v10))+8:]))
																			store64(m.memory[int64(uint32(v4))+232:], uint64(t168))
																			t169 := int64(load64(m.memory[uint32(v10):]))
																			t170 := v4
																			v15 = t169
																			store64(m.memory[int64(uint32(t170))+224:], uint64(v15))
																			v5 = int32(v15)
																			goto l54
																		}
																	}
																	t163 := int64(load64(m.memory[int64(uint32(v4))+324:]))
																	v15 = t163
																	store32(m.memory[int64(uint32(v4))+244:], uint32(v16))
																	store32(m.memory[int64(uint32(v4))+240:], uint32(v17))
																	store64(m.memory[int64(uint32(v4))+232:], uint64(v15))
																	store32(m.memory[int64(uint32(v4))+228:], uint32(v2))
																	store32(m.memory[int64(uint32(v4))+224:], uint32(v5))
																	goto l50
																}
															}
														}
													l32:
														p122 := v14
														if v16 != 0 {
															p122 = i32(0)
														}
														v16 = p122
														t123 := int32(load32(m.memory[uint32(v5):]))
														t124 := int32(load32(m.memory[uint32(v2):]))
														t125 := m.fn1097(t123, t124, i32(1076032), i32(39), i32(1077437), i32(9))
														v17 = t125
														if v17 == 0 {
															{
																t172 := int32(load32(m.memory[uint32(v5):]))
																t173 := int32(load32(m.memory[uint32(v2):]))
																t174 := m.fn1097(t172, t173, i32(1074411), i32(53), i32(1077487), i32(4))
																v17 = t174
																if v17 == 0 {
																	goto l57
																}
																t175 := int32(load32(m.memory[uint32(v17+i32(16)):]))
																t176 := v4 + i32(80)
																v14 = t175
																t177 := int32(load32(m.memory[uint32(v17+i32(20)):]))
																t178 := v14
																v18 = t177
																m.fn845(t176, t178, v18, i32(1073159), i32(67), i32(1073614), i32(5))
																{
																	t179 := int32(load32(m.memory[int64(uint32(v4))+80:]))
																	v17 = t179
																	if v17 == 0 {
																		m.fn845(v4+i32(72), v14, v18, i32(1073159), i32(67), i32(1073228), i32(4))
																		t181 := int32(load32(m.memory[int64(uint32(v4))+72:]))
																		v17 = t181
																		if v17 == 0 {
																			goto l57
																		}
																		t182 := int32(load32(m.memory[int64(uint32(v4))+76:]))
																		v5 = t182
																		goto l59
																	}
																	t180 := int32(load32(m.memory[int64(uint32(v4))+84:]))
																	v5 = t180
																	goto l59
																}
															}
														l57:
															t183 := int32(load32(m.memory[uint32(v5):]))
															t184 := int32(load32(m.memory[uint32(v2):]))
															t185 := m.fn1097(t183, t184, i32(1073576), i32(29), i32(1073605), i32(9))
															v2 = t185
															if v2 == 0 {
																goto l60
															}
															t186 := int32(load32(m.memory[uint32(v2+i32(16)):]))
															t187 := int32(load32(m.memory[uint32(v2+i32(20)):]))
															m.fn845(v4+i32(64), t186, t187, i32(1073159), i32(67), i32(1073226), i32(2))
															t188 := int32(load32(m.memory[int64(uint32(v4))+64:]))
															v17 = t188
															if v17 == 0 {
																goto l60
															}
															t189 := int32(load32(m.memory[int64(uint32(v4))+68:]))
															v5 = t189
															goto l59
														}
														t126 := v4 + i32(112)
														v5 = v17 + i32(16)
														t127 := int32(load32(m.memory[uint32(v5):]))
														v17 = v17 + i32(20)
														t128 := int32(load32(m.memory[uint32(v17):]))
														m.fn1046(t126, t127, t128, i32(1076032), i32(39), i32(1077446), i32(6))
														t129 := int32(load32(m.memory[int64(uint32(v4))+112:]))
														t130 := v4 + i32(260)
														v2 = t129
														p131 := i32(1077452)
														if v2 != 0 {
															p131 = v2
														}
														t132 := int32(load32(m.memory[int64(uint32(v4))+116:]))
														p133 := i32(6)
														if v2 != 0 {
															p133 = t132
														}
														m.fn51(t130, p131, p133)
														m.fn46(v4+i32(104), v13, v16)
														t134 := int32(load32(m.memory[int64(uint32(v4))+108:]))
														if t134 == 0 {
															goto l35
														}
														m.fn31(v4+i32(272), v13, v16)
														goto l36
													}
													v2 = i32(0)
													store32(m.memory[int64(uint32(v4))+196:], uint32(i32(0)))
													store64(m.memory[int64(uint32(v4))+188:], uint64(i64(0x800000000)))
													v16 = v5 << 2
													t94 := int32(load32(m.memory[int64(uint32(v1))+40:]))
													v17 = t94
													t95 := int32(load32(m.memory[int64(uint32(v4))+180:]))
													v13 = t95
													t96 := int32(load32(m.memory[int64(uint32(v4))+176:]))
													v19 = t96
													{
													l39:
														{
															if v16 == v2 {
																m.fn1308(v13, v19)
																m.fn1547(v1, v4+i32(188))
																goto l2
															}
															t135 := int32(load32(m.memory[uint32(v13+v2):]))
															m.fn1314(v4+i32(316), t135, v17)
															{
																t136 := int32(load32(m.memory[int64(uint32(v4))+316:]))
																v5 = t136
																if v5 != i32(-1) {
																	goto l38
																}
																t137 := int64(load64(m.memory[int64(uint32(v4))+324:]))
																store64(m.memory[int64(uint32(v4))+204:], uint64(t137))
																t138 := int32(load32(m.memory[int64(uint32(v4))+320:]))
																store32(m.memory[int64(uint32(v4))+200:], uint32(t138))
																v2 = v2 + i32(4)
																m.fn1271(v4+i32(188), v4+i32(200))
																goto l39
															}
														l38:
														}
														t139 := int64(load64(m.memory[int64(uint32(v4))+324:]))
														v20 = t139
														t140 := int32(load32(m.memory[int64(uint32(v4))+320:]))
														v21 = t140
														t141 := int64(load64(m.memory[int64(uint32(v4))+332:]))
														v15 = t141
														m.fn1308(v13, v19)
														v22 = int32(int64(uint64(v15) >> 32))
														m.fn969(v4 + i32(188))
														v23 = int32(v15)
														goto l40
													}
												}
											l35:
												store32(m.memory[int64(uint32(v4))+320:], uint32(i32(25)))
												store32(m.memory[int64(uint32(v4))+316:], uint32(v4+i32(260)))
												m.fn73(v4+i32(272), i32(1051400), v4+i32(316))
											l36:
												t190 := int32(load32(m.memory[uint32(v5):]))
												t191 := int32(load32(m.memory[uint32(v17):]))
												m.fn845(v4+i32(96), t190, t191, i32(1073159), i32(67), i32(1073226), i32(2))
												t192 := int32(load32(m.memory[int64(uint32(v4))+96:]))
												v2 = t192
												if v2 == 0 {
													goto l61
												}
												t193 := int32(load32(m.memory[int64(uint32(v1))+40:]))
												t194 := v4 + i32(316)
												v14 = t193
												t195 := int32(load32(m.memory[int64(uint32(v4))+100:]))
												m.fn1551(t194, v14, v2, t195)
												t196 := int32(load32(m.memory[int64(uint32(v4))+336:]))
												v16 = t196
												t197 := int32(load32(m.memory[int64(uint32(v4))+332:]))
												v17 = t197
												t198 := int64(load64(m.memory[int64(uint32(v4))+324:]))
												v15 = t198
												t199 := int32(load32(m.memory[int64(uint32(v4))+320:]))
												v2 = t199
												{
													t200 := int32(load32(m.memory[int64(uint32(v4))+316:]))
													v5 = t200
													if v5 == i32(-1) {
														if v2 == i32(-1) {
															goto l61
														}
														store64(m.memory[int64(uint32(v4))+300:], uint64(v15))
														store32(m.memory[int64(uint32(v4))+296:], uint32(v2))
														store32(m.memory[int64(uint32(v4))+312:], uint32(v16))
														store32(m.memory[int64(uint32(v4))+308:], uint32(v17))
														m.fn51(v4+i32(224), i32(1077458), i32(29))
														t201 := int32(load32(m.memory[int64(uint32(v14))+48:]))
														m.fn1182(v4+i32(88), t201, i32(1076292))
														t202 := int32(load32(m.memory[int64(uint32(v4))+92:]))
														v2 = t202
														t203 := int32(load32(m.memory[int64(uint32(v4))+88:]))
														m.fn1296(v4+i32(316), t203, v4+i32(224), v4+i32(296), v17+i32(8), v16)
														t204 := int32(load32(m.memory[uint32(v2):]))
														store32(m.memory[uint32(v2):], uint32(t204+i32(1)))
														t205 := int32(load32(m.memory[int64(uint32(v4))+320:]))
														v2 = t205
														{
															t206 := int32(load32(m.memory[int64(uint32(v4))+316:]))
															v5 = t206
															if v5 == i32(-1) {
																store32(m.memory[int64(uint32(v4))+288:], uint32(v2))
																store32(m.memory[int64(uint32(v4))+284:], uint32(i32(-0x80000000)))
																m.fn754(v4 + i32(308))
																goto l65
															}
															t207 := int64(load64(m.memory[int64(uint32(v4))+332:]))
															v15 = t207
															v22 = int32(int64(uint64(v15) >> 32))
															t208 := int64(load64(m.memory[int64(uint32(v4))+324:]))
															v20 = t208
															m.fn754(v4 + i32(308))
															v23 = int32(v15)
															v21 = v2
															goto l63
														}
													}
													v21 = v2
													v20 = v15
													v23 = v17
													v22 = v16
													goto l63
												}
											}
										l60:
											m.fn46(v4+i32(48), v13, v16)
											{
												t209 := int32(load32(m.memory[int64(uint32(v4))+52:]))
												if t209 == 0 {
													goto l66
												}
												store32(m.memory[int64(uint32(v4))+332:], uint32(i32(-0x7fffffff)))
												store32(m.memory[int64(uint32(v4))+328:], uint32(v16))
												store32(m.memory[int64(uint32(v4))+324:], uint32(v13))
												store32(m.memory[int64(uint32(v4))+320:], uint32(v19))
												store32(m.memory[int64(uint32(v4))+316:], uint32(i32(5)))
												m.fn1544(v1, v4+i32(316))
												goto l67
											}
										l66:
											m.fn16(v19, v13)
										l67:
											t210 := int32(load32(m.memory[int64(uint32(v4))+176:]))
											t211 := int32(load32(m.memory[int64(uint32(v4))+180:]))
											m.fn44(t210, t211)
											goto l2
										}
									l61:
										store32(m.memory[int64(uint32(v4))+284:], uint32(i32(-1)))
									l65:
										t216 := int32(load32(m.memory[int64(uint32(v4))+280:]))
										store32(m.memory[int64(uint32(v8))+8:], uint32(t216))
										t217 := int64(load64(m.memory[int64(uint32(v4))+272:]))
										store64(m.memory[uint32(v8):], uint64(t217))
										store32(m.memory[int64(uint32(v4))+224:], uint32(i32(-0x7fffffff)))
										m.fn1360(v9, v4+i32(284), v4+i32(224))
										store32(m.memory[int64(uint32(v4))+316:], uint32(i32(5)))
										m.fn1544(v1, v4+i32(316))
										v5 = i32(-1)
										goto l68
									}
								l63:
									t218 := int32(load32(m.memory[int64(uint32(v4))+272:]))
									t219 := int32(load32(m.memory[int64(uint32(v4))+276:]))
									m.fn16(t218, t219)
								}
							l68:
								t220 := int32(load32(m.memory[int64(uint32(v4))+260:]))
								t221 := int32(load32(m.memory[int64(uint32(v4))+264:]))
								m.fn16(t220, t221)
								goto l69
							}
						l59:
							t222 := int32(load32(m.memory[int64(uint32(v1))+40:]))
							t223 := v4 + i32(316)
							v2 = t222
							t224 := int32(load32(m.memory[int64(uint32(v2))+32:]))
							t225 := int32(load32(m.memory[int64(uint32(v2))+56:]))
							t226 := int32(load32(m.memory[int64(uint32(v2))+60:]))
							t227 := int32(load32(m.memory[int64(uint32(v2))+48:]))
							m.fn1359(t223, t224, v2, t225, t226, t227, v17, v5)
							t228 := int64(load64(m.memory[int64(uint32(v4))+324:]))
							v15 = t228
							t229 := int32(load32(m.memory[int64(uint32(v4))+320:]))
							v2 = t229
							{
								t230 := int32(load32(m.memory[int64(uint32(v4))+316:]))
								v5 = t230
								if v5 == i32(-1) {
									{
										if v2 == i32(-1) {
											goto l71
										}
										store64(m.memory[int64(uint32(v4))+336:], uint64(v15))
										store32(m.memory[int64(uint32(v4))+332:], uint32(v2))
										store32(m.memory[int64(uint32(v4))+328:], uint32(v16))
										store32(m.memory[int64(uint32(v4))+324:], uint32(v13))
										store32(m.memory[int64(uint32(v4))+320:], uint32(v19))
										store32(m.memory[int64(uint32(v4))+316:], uint32(i32(5)))
										m.fn1544(v1, v4+i32(316))
										goto l72
									l71:
										m.fn46(v4+i32(56), v13, v16)
										t232 := int32(load32(m.memory[int64(uint32(v4))+60:]))
										if t232 == 0 {
											goto l73
										}
										store32(m.memory[int64(uint32(v4))+332:], uint32(i32(-0x7fffffff)))
										store32(m.memory[int64(uint32(v4))+328:], uint32(v16))
										store32(m.memory[int64(uint32(v4))+324:], uint32(v13))
										store32(m.memory[int64(uint32(v4))+320:], uint32(v19))
										store32(m.memory[int64(uint32(v4))+316:], uint32(i32(5)))
										m.fn1544(v1, v4+i32(316))
									}
								l72:
									v5 = i32(-1)
									goto l74
								}
								t231 := int64(load64(m.memory[int64(uint32(v4))+332:]))
								v20 = t231
								v22 = int32(int64(uint64(v20) >> 32))
								v23 = int32(v20)
								v21 = v2
								v20 = v15
								goto l69
							}
						}
					l55:
						v5 = i32(-1)
					l54:
						m.fn754(v4 + i32(316))
						m.fn16(v2, v14)
						t233 := int64(load64(m.memory[int64(uint32(v4))+232:]))
						v15 = t233
						t234 := int32(load32(m.memory[int64(uint32(v4))+228:]))
						v2 = t234
						if v5 == i32(-1) {
							goto l56
						}
					}
				l50:
					t235 := int64(load64(m.memory[int64(uint32(v4))+240:]))
					v20 = t235
					v22 = int32(int64(uint64(v20) >> 32))
					v23 = int32(v20)
					v21 = v2
					v20 = v15
					goto l69
				}
			l56:
				store64(m.memory[int64(uint32(v4))+252:], uint64(v15))
				store32(m.memory[int64(uint32(v4))+248:], uint32(v2))
				m.fn1547(v1, v4+i32(248))
				goto l73
			l47:
				v5 = i32(-1)
			l46:
				m.fn754(v4 + i32(316))
				m.fn16(v2, v14)
				t236 := int64(load64(m.memory[int64(uint32(v4))+232:]))
				v15 = t236
				t237 := int32(load32(m.memory[int64(uint32(v4))+228:]))
				v2 = t237
				if v5 == i32(-1) {
					goto l48
				}
			}
		l42:
			t238 := int64(load64(m.memory[int64(uint32(v4))+240:]))
			v20 = t238
			v22 = int32(int64(uint64(v20) >> 32))
			v23 = int32(v20)
			v21 = v2
			v20 = v15
			goto l69
		}
	l48:
		store64(m.memory[int64(uint32(v4))+216:], uint64(v15))
		store32(m.memory[int64(uint32(v4))+212:], uint32(v2))
		m.fn1547(v1, v4+i32(212))
	l73:
		v5 = i32(-1)
	l69:
		m.fn16(v19, v13)
	l74:
		t239 := int32(load32(m.memory[int64(uint32(v4))+176:]))
		t240 := int32(load32(m.memory[int64(uint32(v4))+180:]))
		m.fn44(t239, t240)
		if v5 == i32(-1) {
			goto l2
		}
	}
l40:
	store64(m.memory[int64(uint32(v0))+8:], uint64(v20))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v21))
	store32(m.memory[uint32(v0):], uint32(v5))
	store64(m.memory[int64(uint32(v0))+16:], uint64(int64(uint32(v22))<<32|int64(uint32(v23))))
l6:
	m.g0 = v4 + i32(432)
}
func (m *Module) fn1544(v0, v1 int32) {
	var v2 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+32:]))
		v2 = t0
		if v2 == 0 {
			goto l0
		}
		t1 := int32(load32(m.memory[int64(uint32(v0))+28:]))
		v2 = t1 + v2*i32(28)
		if v2+i32(-28) == 0 {
			goto l0
		}
		{
			t2 := int32(m.memory[uint32(v2+i32(-4))])
			if t2 != 0 {
				m.fn1340(v2+i32(-16), v1)
				return
			}
			m.fn893(v1)
			return
		}
	}
l0:
	m.fn1340(v0+i32(12), v1)
}
func (m *Module) fn1545(v0, v1 int32) {
	var v2, v3, v4, v5 int32
	t0 := m.g0
	v2 = t0 - i32(64)
	m.g0 = v2
	store32(m.memory[int64(uint32(v2))+16:], uint32(i32(0)))
	store64(m.memory[int64(uint32(v2))+8:], uint64(i64(0x400000000)))
	store32(m.memory[int64(uint32(v2))+28:], uint32(i32(0)))
	store64(m.memory[int64(uint32(v2))+20:], uint64(i64(0x800000000)))
	t1 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	v3 = t1
	t2 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	v4 = t2
	t3 := int32(load32(m.memory[uint32(v1):]))
	store32(m.memory[int64(uint32(v2))+40:], uint32(t3))
	store32(m.memory[int64(uint32(v2))+32:], uint32(v4))
	t4 := v2
	t5 := v4
	v1 = v3 << 4
	v5 = t5 + v1
	store32(m.memory[int64(uint32(t4))+44:], uint32(v5))
l5:
	{
		{
			if v1 == 0 {
				goto l0
			}
			v3 = v4 + i32(4)
			{
				t6 := int32(load32(m.memory[uint32(v4):]))
				switch t6 {
				case 0:
					goto l1
				default:
					t11 := int32(load32(m.memory[int64(uint32(v3))+8:]))
					store32(m.memory[int64(uint32(v2))+56:], uint32(t11))
					t12 := int64(load64(m.memory[uint32(v3):]))
					store64(m.memory[int64(uint32(v2))+48:], uint64(t12))
					m.fn1271(v2+i32(20), v2+i32(48))
					goto l4
				case 2:
					v5 = v4 + i32(16)
				}
			}
		l0:
			store32(m.memory[int64(uint32(v2))+36:], uint32(v5))
			m.fn1542(v2 + i32(32))
			t7 := int32(load32(m.memory[int64(uint32(v2))+16:]))
			store32(m.memory[int64(uint32(v0))+8:], uint32(t7))
			t8 := int64(load64(m.memory[int64(uint32(v2))+8:]))
			store64(m.memory[uint32(v0):], uint64(t8))
			t9 := int64(load64(m.memory[int64(uint32(v2))+20:]))
			store64(m.memory[int64(uint32(v0))+12:], uint64(t9))
			t10 := int32(load32(m.memory[int64(uint32(v2))+28:]))
			store32(m.memory[int64(uint32(v0))+20:], uint32(t10))
			m.g0 = v2 + i32(64)
			return
		}
	l1:
		t13 := int32(load32(m.memory[int64(uint32(v3))+8:]))
		store32(m.memory[int64(uint32(v2))+56:], uint32(t13))
		t14 := int64(load64(m.memory[uint32(v3):]))
		store64(m.memory[int64(uint32(v2))+48:], uint64(t14))
		m.fn1525(v2+i32(8), v2+i32(48))
		m.fn894(v2 + i32(48))
	}
l4:
	v4 = v4 + i32(16)
	v1 = v1 + i32(-16)
	goto l5
}
func (m *Module) fn1546(v0, v1, v2, v3 int32) {
	var v4, v5, v6, v7 int32
	t0 := m.g0
	v4 = t0 - i32(48)
	m.g0 = v4
	m.fn1403(v4+i32(20), v1, v2, v3)
	t1 := int32(load32(m.memory[int64(uint32(v4))+20:]))
	v1 = t1
	t2 := int32(load32(m.memory[int64(uint32(v4))+24:]))
	t3 := v4
	v3 = t2
	t4 := int32(load32(m.memory[int64(uint32(v4))+28:]))
	t5 := v3
	v2 = t4 * i32(28)
	v5 = t5 + v2
	store32(m.memory[int64(uint32(t3))+16:], uint32(v5))
	store32(m.memory[int64(uint32(v4))+12:], uint32(v1))
	store32(m.memory[int64(uint32(v4))+4:], uint32(v3))
	v6 = v3 + i32(28)
	v1 = v4 + i32(24)
l2:
	{
		if v2 == 0 {
			goto l0
		}
		t6 := int32(load32(m.memory[uint32(v3):]))
		v7 = t6
		if v7 != i32(-1) {
			t7 := int64(load64(m.memory[int64(uint32(v3))+4:]))
			store64(m.memory[uint32(v1):], uint64(t7))
			t8 := int64(load64(m.memory[int64(uint32(v3))+12:]))
			store64(m.memory[int64(uint32(v1))+8:], uint64(t8))
			t9 := int64(load64(m.memory[int64(uint32(v3))+20:]))
			store64(m.memory[int64(uint32(v1))+16:], uint64(t9))
			store32(m.memory[int64(uint32(v4))+20:], uint32(v7))
			m.fn1544(v0, v4+i32(20))
			v2 = v2 + i32(-28)
			v6 = v6 + i32(28)
			v3 = v3 + i32(28)
			goto l2
		}
		v5 = v6
	}
l0:
	store32(m.memory[int64(uint32(v4))+8:], uint32(v5))
	m.fn900(v4 + i32(4))
	m.g0 = v4 + i32(48)
}
func (m *Module) fn1547(v0, v1 int32) {
	var v2, v3 int32
	var v4 int64
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	{
		t1 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		if t1 == 0 {
			goto l0
		}
		{
			t2 := int32(load32(m.memory[int64(uint32(v0))+20:]))
			if t2 == 0 {
				goto l1
			}
			t3 := int32(load32(m.memory[int64(uint32(v0))+20:]))
			v3 = t3
			store32(m.memory[int64(uint32(v0))+20:], uint32(i32(0)))
			t4 := int64(load64(m.memory[int64(uint32(v0))+12:]))
			v4 = t4
			store64(m.memory[int64(uint32(v0))+12:], uint64(i64(0x400000000)))
			store32(m.memory[uint32(v2):], uint32(i32(0)))
			store32(m.memory[int64(uint32(v2))+12:], uint32(v3))
			store64(m.memory[int64(uint32(v2))+4:], uint64(v4))
			m.fn832(v0, v2)
		}
	l1:
		store32(m.memory[uint32(v2):], uint32(i32(1)))
		t5 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		store32(m.memory[int64(uint32(v2))+12:], uint32(t5))
		t6 := int64(load64(m.memory[uint32(v1):]))
		store64(m.memory[int64(uint32(v2))+4:], uint64(t6))
		m.fn832(v0, v2)
		goto l2
	}
l0:
	m.fn969(v1)
l2:
	m.g0 = v2 + i32(16)
}
func (m *Module) fn1548(v0 int32) {
	var v1 int32
	v1 = v0 + i32(4)
	{
		t0 := int32(load32(m.memory[uint32(v0):]))
		if t0 != 0 {
			goto l0
		}
		m.fn894(v1)
		return
	}
l0:
	m.fn969(v1)
}
func (m *Module) fn1549(v0, v1 int32) {
	var v2 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		v2 = t0
		t1 := int32(load32(m.memory[uint32(v0):]))
		if v2 != t1 {
			goto l0
		}
		m.fn618(v0)
	}
l0:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v2+i32(1)))
	t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	store32(m.memory[uint32(t2+v2<<2):], uint32(v1))
}
func (m *Module) fn1550(v0, v1 int32) {
	var v2, v3 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	t1 := int32(load32(m.memory[int64(uint32(v0))+32:]))
	v3 = t1
	t2 := int32(load32(m.memory[int64(uint32(v0))+28:]))
	t3 := v2
	v0 = t2
	store32(m.memory[int64(uint32(t3))+8:], uint32(v0))
	store32(m.memory[int64(uint32(v2))+12:], uint32(v0+v3*i32(44)))
l1:
	{
		t4 := m.fn904(v2 + i32(8))
		v0 = t4
		if v0 == 0 {
			goto l0
		}
		t5 := m.fn847(v0, i32(1073848), i32(59), i32(1073907), i32(8))
		if t5 != 0 {
			goto l1
		}
		{
			t6 := m.fn847(v0, i32(1072544), i32(60), i32(1086232), i32(11))
			if t6 != 0 {
				m.fn1549(v1, v0)
				goto l1
			}
			m.fn1550(v0, v1)
			goto l1
		}
	}
l0:
	m.g0 = v2 + i32(16)
}
func (m *Module) fn1551(v0, v1, v2, v3 int32) {
	t0 := int32(load32(m.memory[int64(uint32(v1))+32:]))
	t1 := int32(load32(m.memory[int64(uint32(v1))+56:]))
	t2 := int32(load32(m.memory[int64(uint32(v1))+60:]))
	m.fn1496(v0, t0, v1, t1, t2, v2, v3)
}
func (m *Module) fn1552(v0, v1, v2, v3, v4 int32) {
	var v5, v6, v7 int32
	var v8, v9 int64
	var v10, v11, v12, v13 int32
	t0 := m.g0
	v5 = t0 - i32(16)
	m.g0 = v5
	v1 = v1 * i32(28)
	t1 := int32(load32(m.memory[int64(uint32(v2))+4:]))
	v6 = t1
	t2 := int32(load32(m.memory[uint32(v2):]))
	v7 = t2
	t3 := int64(load64(m.memory[int64(uint32(v2))+24:]))
	v8 = t3
	t4 := int64(load64(m.memory[int64(uint32(v2))+16:]))
	v9 = t4
	t5 := int32(load32(m.memory[int64(uint32(v2))+12:]))
	v10 = t5
l4:
	{
		{
			if v1 == 0 {
				m.g0 = v5 + i32(16)
				return
			}
			t6 := int32(load32(m.memory[uint32(v0):]))
			v11 = t6
			p7 := i32(1)
			if uint32(v11) > uint32(i32(2)) {
				p7 = v11 + i32(-3)
			}
			switch p7 + i32(-1) {
			case 0:
				t8 := int32(load32(m.memory[uint32(v0+i32(20)):]))
				t9 := int32(load32(m.memory[uint32(v0+i32(24)):]))
				m.fn1552(t8, t9, v2, v3, v4)
				goto l2
			case 3:
				goto l3
			default:
				goto l2
			}
		}
	l3:
		if v10 == 0 {
			goto l2
		}
		t10 := int32(load32(m.memory[uint32(v0+i32(8)):]))
		t11 := v7
		t12 := v6
		t13 := v9
		t14 := v8
		v11 = t10
		t15 := int32(load32(m.memory[uint32(v0+i32(12)):]))
		t16 := v11
		v12 = t15
		t17 := m.fn29(t13, t14, t16, v12)
		t18 := m.fn30(t11, t12, t17, v11, v12)
		v13 = t18
		if v13 == 0 {
			goto l2
		}
		m.fn31(v5+i32(4), v11, v12)
		t19 := m.fn32(v4, v5+i32(4))
		if t19 == 0 {
			goto l2
		}
		m.fn31(v5+i32(4), v11, v12)
		m.fn33(v3, v5+i32(4))
		t20 := int32(load32(m.memory[uint32(v13+i32(-4)):]))
		v11 = t20
		t21 := int32(load32(m.memory[int64(uint32(v11))+16:]))
		t22 := int32(load32(m.memory[int64(uint32(v11))+20:]))
		m.fn28(t21, t22, v2, v3, v4)
	}
l2:
	v0 = v0 + i32(28)
	v1 = v1 + i32(-28)
	goto l4
}
func (m *Module) fn1553(v0, v1, v2 int32) {
	var v3 int32
	v1 = v1 * i32(28)
l3:
	if v1 == 0 {
		return
	}
	{
		t0 := int32(load32(m.memory[uint32(v0):]))
		v3 = t0
		if uint32(v3) >= uint32(i32(3)) {
			goto l1
		}
		{
			if v3 != i32(2) {
				goto l2
			}
			t1 := int32(load32(m.memory[uint32(v0+i32(8)):]))
			t2 := int32(load32(m.memory[uint32(v0+i32(12)):]))
			_ = m.fn1521(v2, t1, t2)
		}
	l2:
		t4 := int32(load32(m.memory[uint32(v0+i32(20)):]))
		t5 := int32(load32(m.memory[uint32(v0+i32(24)):]))
		m.fn1553(t4, t5, v2)
	}
l1:
	v0 = v0 + i32(28)
	v1 = v1 + i32(-28)
	goto l3
}
func (m *Module) fn1554(v0, v1, v2 int32) {
	var v3, v4 int32
	var v5 int64
	var v6, v7, v8 int32
	var v9 int64
	var v10, v11, v12 int32
	var v13, v14 int64
	var v15 int32
	t0 := int64(load64(m.memory[int64(uint32(v1))+16:]))
	t1 := int64(load64(m.memory[int64(uint32(v1))+24:]))
	t2 := int32(load32(m.memory[int64(uint32(v2))+4:]))
	v3 = t2
	t3 := int32(load32(m.memory[int64(uint32(v2))+8:]))
	t4 := v3
	v4 = t3
	t5 := m.fn540(t0, t1, t4, v4)
	v5 = t5
	t6 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	v6 = t6
	t7 := v6
	v7 = int32(v5)
	v8 = t7 & v7
	v9 = int64(uint64(v5)>>25) & i64(127) * i64(72340172838076673)
	v10 = v1 + i32(16)
	t8 := int32(load32(m.memory[uint32(v1):]))
	v11 = t8
	v12 = i32(0)
l6:
	{
		t9 := int64(load64(m.memory[uint32(v11+v8):]))
		v13 = t9
		v14 = v13 ^ v9
		v14 = (v14 ^ i64(-1)) & (v14 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
		{
		l2:
			{
				if v14 == 0 {
					if v13&(v13<<1)&i64(-0x7f7f7f7f7f7f7f80) == 0 {
						t16 := v8
						v12 = v12 + i32(8)
						v8 = (t16 + v12) & v6
						goto l6
					}
					m.fn35(v1, i32(1), v10)
					v8 = int32(int64(uint64(v5) >> 32))
					{
						t13 := int32(load32(m.memory[uint32(v2):]))
						v11 = t13
						if v11 != i32(-1) {
							t14 := int64(load64(m.memory[int64(uint32(v2))+4:]))
							v5 = t14
							store32(m.memory[int64(uint32(v0))+20:], uint32(v1))
							store64(m.memory[int64(uint32(v0))+12:], uint64(v5))
							store32(m.memory[int64(uint32(v0))+8:], uint32(v11))
							store32(m.memory[int64(uint32(v0))+4:], uint32(v8))
							store32(m.memory[uint32(v0):], uint32(v7))
							return
						}
						v1 = v8
						v15 = v7
						goto l5
					}
				}
				v15 = v11 - (int32(uint32(int64(bits.TrailingZeros64(uint64(v14))))>>3)+v8)&v6<<4
				t10 := int32(load32(m.memory[uint32(v15+i32(-12)):]))
				t11 := int32(load32(m.memory[uint32(v15+i32(-8)):]))
				t12 := m.fn191(t10, t11, v3, v4)
				if t12 != 0 {
					goto l1
				}
				v14 = (v14 + i64(-1)) & v14
				goto l2
			}
		l1:
			t15 := int32(load32(m.memory[uint32(v2):]))
			m.fn16(t15, v3)
		}
	l5:
		store32(m.memory[int64(uint32(v0))+8:], uint32(i32(-1)))
		store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
		store32(m.memory[uint32(v0):], uint32(v15))
		return
	}
}
func (m *Module) fn1555(v0 int32) {
	var v1, v2, v3 int32
	var v4 int64
	var v5, v6, v7, v8 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		if t0 == i32(-1) {
			return
		}
		t1 := int32(load32(m.memory[int64(uint32(v0))+20:]))
		v1 = t1
		t2 := int32(load32(m.memory[uint32(v1):]))
		v2 = t2
		t3 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t4 := v2
		t5 := v2
		v3 = t3
		t6 := int64(load64(m.memory[uint32(v0):]))
		t7 := v3
		v4 = t6
		t8 := m.fn26(t5, t7, v4)
		v5 = t8
		v6 = t4 + v5
		t9 := int32(m.memory[uint32(v6)])
		v7 = t9
		t10 := v6
		v8 = int32(uint32(int32(v4)) >> 25)
		m.memory[uint32(t10)] = byte(v8)
		m.memory[uint32(v2+v3&(v5+i32(-8))+i32(8))] = byte(v8)
		t11 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		store32(m.memory[int64(uint32(v1))+8:], uint32(t11-v7&i32(1)))
		v2 = v2 - v5<<4
		v5 = v2 + i32(-16)
		t12 := v5
		v0 = v0 + i32(8)
		t13 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		store32(m.memory[int64(uint32(t12))+8:], uint32(t13))
		t14 := int64(load64(m.memory[uint32(v0):]))
		store64(m.memory[uint32(v5):], uint64(t14))
		store32(m.memory[uint32(v2+i32(-4)):], uint32(i32(1)))
		t15 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		store32(m.memory[int64(uint32(v1))+12:], uint32(t15+i32(1)))
	}
}
func (m *Module) fn1556(v0, v1, v2 int32) {
	var v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13 int32
	var v14 int64
	var v15 int32
	var v16 int64
	var v17 int32
	var v18 int64
	var v19, v20, v21, v22, v23, v24 int32
	t0 := m.g0
	v3 = t0 - i32(80)
	m.g0 = v3
	t1 := int32(load32(m.memory[int64(uint32(v2))+4:]))
	v4 = t1
	v5 = v4 + i32(16)
	v6 = v0 + v1*i32(28)
	v7 = v3 + i32(52)
	t2 := int32(load32(m.memory[int64(uint32(v2))+8:]))
	v8 = t2
	t3 := int32(load32(m.memory[uint32(v2):]))
	v9 = t3
	v10 = v4 + i32(4)
l2:
	{
		v1 = v0
		if v1 == v6 {
			m.g0 = v3 + i32(80)
			return
		}
		v0 = v1 + i32(28)
		t4 := int32(load32(m.memory[uint32(v1):]))
		v11 = t4
		p5 := i32(1)
		if uint32(v11) > uint32(i32(2)) {
			p5 = v11 + i32(-3)
		}
		switch p5 + i32(-1) {
		case 0:
			t6 := int32(load32(m.memory[int64(uint32(v1))+20:]))
			t7 := int32(load32(m.memory[int64(uint32(v1))+24:]))
			m.fn1556(t6, t7, v2)
			goto l2
		case 2:
			t8 := int32(load32(m.memory[int64(uint32(v1))+12:]))
			v12 = t8
			t9 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			v13 = t9
			t10 := int32(load32(m.memory[int64(uint32(v9))+12:]))
			if t10 == 0 {
				goto l2
			}
			t11 := int64(load64(m.memory[int64(uint32(v9))+16:]))
			t12 := int64(load64(m.memory[int64(uint32(v9))+24:]))
			t13 := m.fn29(t11, t12, v13, v12)
			v14 = t13
			t14 := int32(load32(m.memory[int64(uint32(v9))+4:]))
			v15 = t14
			v1 = v15 & int32(v14)
			v16 = int64(uint64(v14)>>25) & i64(127) * i64(72340172838076673)
			t15 := int32(load32(m.memory[uint32(v9):]))
			v11 = t15
			v17 = i32(0)
		l7:
			{
				t16 := int64(load64(m.memory[uint32(v11+v1):]))
				v18 = t16
				v14 = v18 ^ v16
				v14 = (v14 ^ i64(-1)) & (v14 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
			l6:
				{
					if v14 == 0 {
						if !(v18&(v18<<1)&i64(-0x7f7f7f7f7f7f7f80) == 0) {
							goto l2
						}
						t22 := v1
						v17 = v17 + i32(8)
						v1 = (t22 + v17) & v15
						goto l7
					}
					t17 := v13
					t18 := v12
					v19 = v11 - (int32(uint32(int64(bits.TrailingZeros64(uint64(v14))))>>3)+v1)&v15<<3
					t19 := int32(load32(m.memory[uint32(v19+i32(-8)):]))
					t20 := int32(load32(m.memory[uint32(v19+i32(-4)):]))
					t21 := m.fn15(t17, t18, t19, t20)
					if t21 != 0 {
						goto l5
					}
					v14 = (v14 + i64(-1)) & v14
					goto l6
				}
			l5:
			}
			{
				t23 := int32(load32(m.memory[int64(uint32(v4))+12:]))
				if t23 == 0 {
					goto l8
				}
				t24 := int64(load64(m.memory[int64(uint32(v4))+16:]))
				t25 := int64(load64(m.memory[int64(uint32(v4))+24:]))
				t26 := m.fn29(t24, t25, v13, v12)
				v14 = t26
				t27 := int32(load32(m.memory[uint32(v4):]))
				t28 := int32(load32(m.memory[uint32(v10):]))
				t29 := m.fn648(t27, t28, v14, v13, v12)
				if t29 != 0 {
					goto l2
				}
			}
		l8:
			m.fn140(v3+i32(28), v12)
			store32(m.memory[int64(uint32(v3))+40:], uint32(v13))
			store32(m.memory[int64(uint32(v3))+44:], uint32(v13+v12))
			v15 = i32(0)
		l12:
			{
				{
					t30 := m.fn48(v3 + i32(40))
					v1 = t30
					if v1 == i32(-1) {
						t33 := int32(load32(m.memory[int64(uint32(v3))+32:]))
						t34 := v3 + i32(40)
						v20 = t33
						t35 := int32(load32(m.memory[int64(uint32(v3))+36:]))
						m.fn514(t34, i32(45), v20, t35)
						m.fn627(v3+i32(68), v3+i32(40))
						t36 := int32(load32(m.memory[int64(uint32(v3))+76:]))
						t37 := int32(load32(m.memory[int64(uint32(v3))+68:]))
						v21 = t37
						p38 := i32(0)
						if v21 != 0 {
							p38 = t36
						}
						v22 = p38
						t39 := int32(load32(m.memory[int64(uint32(v3))+44:]))
						v15 = t39
						t40 := int32(load32(m.memory[int64(uint32(v3))+52:]))
						t41 := v15
						v17 = t40
						v19 = t41 + v17
						t42 := int32(load32(m.memory[int64(uint32(v3))+72:]))
						v23 = t42
						t43 := int32(load32(m.memory[int64(uint32(v3))+40:]))
						v24 = t43
						t44 := int32(load32(m.memory[int64(uint32(v3))+56:]))
						v1 = t44
					l14:
						{
							store32(m.memory[int64(uint32(v3))+68:], uint32(v19))
							t45 := v3
							t46 := v15
							v11 = v1
							store32(m.memory[int64(uint32(t45))+72:], uint32(t46+v11))
							m.fn577(v3+i32(8), v3+i32(68))
							t47 := int32(load32(m.memory[int64(uint32(v3))+8:]))
							if t47 != i32(1) {
								goto l13
							}
							t48 := int32(load32(m.memory[int64(uint32(v3))+72:]))
							t49 := int32(load32(m.memory[int64(uint32(v3))+68:]))
							v1 = t48 - t49 + v17
							t50 := int32(load32(m.memory[int64(uint32(v3))+12:]))
							if t50 == v24 {
								goto l14
							}
							goto l15
						}
					l13:
						v11 = v22
					l15:
						t52 := v11
						p51 := i32(0)
						if v21 != 0 {
							p51 = v23
						}
						v1 = p51
						if t52 == v1 {
							goto l16
						}
						m.fn51(v3+i32(40), v20+v1, v11-v1)
						goto l17
					}
					p31 := v1
					if uint32(v1+i32(-65)) < uint32(i32(26)) {
						p31 = v1 | i32(32)
					}
					v1 = p31
					if uint32(v1+i32(-97)) < uint32(i32(26)) {
						goto l10
					}
					if uint32(v1+i32(-48)) < uint32(i32(10)) {
						goto l10
					}
					var p32 int32
					if v1 == i32(45) {
						p32 = 1
					}
					v11 = p32
					if v11 != 0 {
						goto l11
					}
					if v1 == i32(95) {
						goto l11
					}
					if v15&i32(1) != 0 {
						goto l12
					}
					v1 = i32(45)
					goto l10
				}
			l11:
				if v15&v11 != 0 {
					goto l12
				}
			l10:
				m.fn74(v3+i32(28), v1)
				var p53 int32
				if v1 == i32(45) {
					p53 = 1
				}
				v15 = p53
				goto l12
			}
		l16:
			m.fn51(v3+i32(40), i32(1077608), i32(6))
		l17:
			t54 := int32(load32(m.memory[int64(uint32(v3))+28:]))
			m.fn16(t54, v20)
			m.fn52(v3+i32(16), v8, v3+i32(40))
			t55 := int32(load32(m.memory[int64(uint32(v3))+16:]))
			if t55 == i32(-1) {
				goto l2
			}
			m.fn51(v3+i32(68), v13, v12)
			t56 := int64(load64(m.memory[int64(uint32(v4))+16:]))
			t57 := int64(load64(m.memory[int64(uint32(v4))+24:]))
			t58 := int32(load32(m.memory[int64(uint32(v3))+72:]))
			t59 := int32(load32(m.memory[int64(uint32(v3))+76:]))
			t60 := m.fn540(t56, t57, t58, t59)
			v14 = t60
			store32(m.memory[int64(uint32(v3))+28:], uint32(v3+i32(68)))
			m.fn676(v4, v5)
			store32(m.memory[int64(uint32(v3))+44:], uint32(v4))
			t61 := int32(load32(m.memory[uint32(v10):]))
			v1 = t61
			store32(m.memory[int64(uint32(v3))+40:], uint32(v3+i32(28)))
			t62 := int32(load32(m.memory[uint32(v4):]))
			m.fn69(v3, t62, v1, v14, v3+i32(40), i32(194))
			t63 := int32(load32(m.memory[uint32(v4):]))
			v1 = t63
			t64 := int32(load32(m.memory[int64(uint32(v3))+4:]))
			v11 = t64
			{
				t65 := int32(load32(m.memory[uint32(v3):]))
				if t65 != i32(1) {
					v1 = v1 + (i32(0)-v11)*i32(28)
					v11 = v1 + i32(-16)
					t77 := int32(load32(m.memory[int64(uint32(v3))+24:]))
					store32(m.memory[int64(uint32(v11))+8:], uint32(t77))
					m.memory[uint32(v1+i32(-4))] = byte(i32(1))
					t78 := int32(load32(m.memory[uint32(v1+i32(-12)):]))
					v15 = t78
					t79 := int32(load32(m.memory[uint32(v11):]))
					v1 = t79
					t80 := int64(load64(m.memory[int64(uint32(v3))+16:]))
					store64(m.memory[uint32(v11):], uint64(t80))
					t81 := int32(load32(m.memory[int64(uint32(v3))+68:]))
					t82 := int32(load32(m.memory[int64(uint32(v3))+72:]))
					m.fn16(t81, t82)
					if v1 == i32(-1) {
						goto l2
					}
					m.fn16(v1, v15)
					goto l2
				}
				v15 = v1 + v11
				t66 := int32(m.memory[uint32(v15)])
				v19 = t66
				t67 := int32(load32(m.memory[int64(uint32(v3))+76:]))
				v17 = t67
				t68 := int64(load64(m.memory[int64(uint32(v3))+68:]))
				v16 = t68
				t69 := v15
				v24 = int32(uint32(int32(v14)) >> 25)
				m.memory[uint32(t69)] = byte(v24)
				t70 := int32(load32(m.memory[uint32(v10):]))
				m.memory[uint32(v1+t70&(v11+i32(-8))+i32(8))] = byte(v24)
				t71 := int64(load64(m.memory[int64(uint32(v3))+16:]))
				store64(m.memory[uint32(v7):], uint64(t71))
				t72 := int32(load32(m.memory[int64(uint32(v3))+24:]))
				store32(m.memory[int64(uint32(v7))+8:], uint32(t72))
				t73 := int32(load32(m.memory[int64(uint32(v4))+12:]))
				store32(m.memory[int64(uint32(v4))+12:], uint32(t73+i32(1)))
				t74 := int32(load32(m.memory[int64(uint32(v4))+8:]))
				store32(m.memory[int64(uint32(v4))+8:], uint32(t74-v19&i32(1)))
				v11 = v1 + (i32(0)-v11)*i32(28)
				v1 = v11 + i32(-28)
				store64(m.memory[uint32(v1):], uint64(v16))
				store32(m.memory[int64(uint32(v3))+48:], uint32(v17))
				t75 := int64(load64(m.memory[int64(uint32(v3))+48:]))
				store64(m.memory[int64(uint32(v1))+8:], uint64(t75))
				t76 := int64(load64(m.memory[int64(uint32(v3))+56:]))
				store64(m.memory[int64(uint32(v1))+16:], uint64(t76))
				m.memory[uint32(v11+i32(-4))] = byte(i32(1))
				goto l2
			}
		default:
			goto l2
		}
	}
}
func (m *Module) fn1557(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6 int32
	v2 = i32(0)
	{
		t1 := v1
		p0 := i32(16)
		if uint32(v0) > uint32(i32(16)) {
			p0 = v0
		}
		v0 = p0
		if uint32(t1) >= uint32(i32(-65587)-v0) {
			goto l0
		}
		t3 := v0
		p2 := (v1 + i32(11)) & i32(-8)
		if uint32(v1) < uint32(i32(11)) {
			p2 = i32(16)
		}
		v3 = p2
		t4 := m.fn4(t3 + v3 + i32(12))
		v1 = t4
		if v1 == 0 {
			goto l0
		}
		v2 = v1 + i32(-8)
		{
			v4 = v0 + i32(-1)
			if v4&v1 != 0 {
				goto l1
			}
			v0 = v2
			goto l2
		l1:
			v5 = v1 + i32(-4)
			t5 := int32(load32(m.memory[uint32(v5):]))
			v6 = t5
			t6 := v6 & i32(-8)
			v1 = (v4+v1)&(i32(0)-v0) + i32(-8)
			t8 := v1
			p7 := v0
			if uint32(v1-v2) > uint32(i32(16)) {
				p7 = i32(0)
			}
			v0 = t8 + p7
			v1 = v0 - v2
			v4 = t6 - v1
			{
				if v6&i32(3) == 0 {
					goto l3
				}
				t9 := int32(load32(m.memory[int64(uint32(v0))+4:]))
				store32(m.memory[int64(uint32(v0))+4:], uint32(v4|t9&i32(1)|i32(2)))
				v4 = v0 + v4
				t10 := int32(load32(m.memory[int64(uint32(v4))+4:]))
				store32(m.memory[int64(uint32(v4))+4:], uint32(t10|i32(1)))
				t11 := int32(load32(m.memory[uint32(v5):]))
				store32(m.memory[uint32(v5):], uint32(v1|t11&i32(1)|i32(2)))
				v4 = v2 + v1
				t12 := int32(load32(m.memory[int64(uint32(v4))+4:]))
				store32(m.memory[int64(uint32(v4))+4:], uint32(t12|i32(1)))
				m.fn1560(v2, v1)
				goto l2
			}
		l3:
			t13 := int32(load32(m.memory[uint32(v2):]))
			v2 = t13
			store32(m.memory[int64(uint32(v0))+4:], uint32(v4))
			store32(m.memory[uint32(v0):], uint32(v2+v1))
		}
	l2:
		{
			t14 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			v1 = t14
			if v1&i32(3) == 0 {
				goto l4
			}
			v2 = v1 & i32(-8)
			if uint32(v2) <= uint32(v3+i32(16)) {
				goto l4
			}
			store32(m.memory[int64(uint32(v0))+4:], uint32(v3|v1&i32(1)|i32(2)))
			v1 = v0 + v3
			t15 := v1
			v3 = v2 - v3
			store32(m.memory[int64(uint32(t15))+4:], uint32(v3|i32(3)))
			v2 = v0 + v2
			t16 := int32(load32(m.memory[int64(uint32(v2))+4:]))
			store32(m.memory[int64(uint32(v2))+4:], uint32(t16|i32(1)))
			m.fn1560(v1, v3)
		}
	l4:
		v2 = v0 + i32(8)
	}
l0:
	return v2
}
func (m *Module) fn1558(v0 int32) {
	var v1, v2, v3, v4 int32
	v1 = v0 + i32(-8)
	t0 := int32(load32(m.memory[uint32(v0+i32(-4)):]))
	t1 := v1
	v2 = t0
	v0 = v2 & i32(-8)
	v3 = t1 + v0
	{
		if v2&i32(1) != 0 {
			goto l0
		}
		if v2&i32(2) == 0 {
			return
		}
		t2 := int32(load32(m.memory[uint32(v1):]))
		v2 = t2
		v0 = v2 + v0
		{
			v1 = v1 - v2
			t3 := int32(load32(m.memory[int64(uint32(i32(0)))+1303580:]))
			if v1 != t3 {
				goto l2
			}
			t4 := int32(load32(m.memory[int64(uint32(v3))+4:]))
			if t4&i32(3) != i32(3) {
				goto l0
			}
			store32(m.memory[int64(uint32(i32(0)))+1303572:], uint32(v0))
			t5 := int32(load32(m.memory[int64(uint32(v3))+4:]))
			store32(m.memory[int64(uint32(v3))+4:], uint32(t5&i32(-2)))
			store32(m.memory[int64(uint32(v1))+4:], uint32(v0|i32(1)))
			store32(m.memory[uint32(v3):], uint32(v0))
			return
		}
	l2:
		m.fn1559(v1, v2)
	}
l0:
	{
		{
			{
				t6 := int32(load32(m.memory[int64(uint32(v3))+4:]))
				v2 = t6
				if v2&i32(2) != 0 {
					goto l3
				}
				t7 := int32(load32(m.memory[int64(uint32(i32(0)))+1303584:]))
				if v3 == t7 {
					store32(m.memory[int64(uint32(i32(0)))+1303584:], uint32(v1))
					t14 := int32(load32(m.memory[int64(uint32(i32(0)))+1303576:]))
					v0 = t14 + v0
					store32(m.memory[int64(uint32(i32(0)))+1303576:], uint32(v0))
					store32(m.memory[int64(uint32(v1))+4:], uint32(v0|i32(1)))
					{
						t15 := int32(load32(m.memory[int64(uint32(i32(0)))+1303580:]))
						if v1 != t15 {
							goto l10
						}
						store32(m.memory[int64(uint32(i32(0)))+1303572:], uint32(i32(0)))
						store32(m.memory[int64(uint32(i32(0)))+1303580:], uint32(i32(0)))
					}
				l10:
					t16 := int32(load32(m.memory[int64(uint32(i32(0)))+1303596:]))
					t17 := v0
					v2 = t16
					if uint32(t17) <= uint32(v2) {
						return
					}
					t18 := int32(load32(m.memory[int64(uint32(i32(0)))+1303584:]))
					v0 = t18
					if v0 == 0 {
						return
					}
					t19 := int32(load32(m.memory[int64(uint32(i32(0)))+1303576:]))
					v4 = t19
					if uint32(v4) < uint32(i32(41)) {
						goto l11
					}
					v1 = i32(1303284)
				l13:
					{
						{
							t20 := int32(load32(m.memory[uint32(v1):]))
							v3 = t20
							if uint32(v3) > uint32(v0) {
								goto l12
							}
							t21 := int32(load32(m.memory[int64(uint32(v1))+4:]))
							if uint32(v0) < uint32(v3+t21) {
								goto l11
							}
						}
					l12:
						t22 := int32(load32(m.memory[int64(uint32(v1))+8:]))
						v1 = t22
						goto l13
					}
				}
				t8 := int32(load32(m.memory[int64(uint32(i32(0)))+1303580:]))
				if v3 == t8 {
					store32(m.memory[int64(uint32(i32(0)))+1303580:], uint32(v1))
					t23 := int32(load32(m.memory[int64(uint32(i32(0)))+1303572:]))
					v0 = t23 + v0
					store32(m.memory[int64(uint32(i32(0)))+1303572:], uint32(v0))
					store32(m.memory[int64(uint32(v1))+4:], uint32(v0|i32(1)))
					store32(m.memory[uint32(v1+v0):], uint32(v0))
					return
				}
				t9 := v3
				v2 = v2 & i32(-8)
				m.fn1559(t9, v2)
				t10 := v1
				v0 = v2 + v0
				store32(m.memory[int64(uint32(t10))+4:], uint32(v0|i32(1)))
				store32(m.memory[uint32(v1+v0):], uint32(v0))
				t11 := int32(load32(m.memory[int64(uint32(i32(0)))+1303580:]))
				if v1 != t11 {
					goto l6
				}
				store32(m.memory[int64(uint32(i32(0)))+1303572:], uint32(v0))
				return
			}
		l3:
			store32(m.memory[int64(uint32(v3))+4:], uint32(v2&i32(-2)))
			store32(m.memory[int64(uint32(v1))+4:], uint32(v0|i32(1)))
			store32(m.memory[uint32(v1+v0):], uint32(v0))
		l6:
			if uint32(v0) < uint32(i32(256)) {
				{
					{
						t26 := int32(load32(m.memory[int64(uint32(i32(0)))+1303564:]))
						v3 = t26
						t27 := v3
						v2 = i32_shl(i32(1), int32(uint32(v0)>>3))
						if t27&v2 != 0 {
							goto l15
						}
						store32(m.memory[int64(uint32(i32(0)))+1303564:], uint32(v3|v2))
						v0 = v0&i32(248) + i32(1303300)
						v3 = v0
						goto l16
					}
				l15:
					v0 = v0 & i32(248)
					v3 = v0 + i32(1303300)
					t28 := int32(load32(m.memory[uint32(v0+i32(1303308)):]))
					v0 = t28
				}
			l16:
				store32(m.memory[int64(uint32(v3))+8:], uint32(v1))
				store32(m.memory[int64(uint32(v0))+12:], uint32(v1))
				store32(m.memory[int64(uint32(v1))+12:], uint32(v3))
				store32(m.memory[int64(uint32(v1))+8:], uint32(v0))
				return
			}
			m.fn1809(v1, v0)
			t12 := int32(load32(m.memory[int64(uint32(i32(0)))+1303604:]))
			v1 = t12 + i32(-1)
			store32(m.memory[int64(uint32(i32(0)))+1303604:], uint32(v1))
			if v1 != 0 {
				return
			}
			t13 := int32(load32(m.memory[int64(uint32(i32(0)))+1303292:]))
			v0 = t13
			if v0 != 0 {
				goto l8
			}
			v1 = i32(0xfff)
			goto l9
		}
	l8:
		v1 = i32(0)
	l14:
		{
			v1 = v1 + i32(1)
			t24 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			v0 = t24
			if v0 != 0 {
				goto l14
			}
		}
		p25 := i32(0xfff)
		if uint32(v1) > uint32(i32(0xfff)) {
			p25 = v1
		}
		v1 = p25
	}
l9:
	store32(m.memory[int64(uint32(i32(0)))+1303604:], uint32(v1))
	return
l11:
	{
		{
			t29 := int32(load32(m.memory[int64(uint32(i32(0)))+1303292:]))
			v0 = t29
			if v0 != 0 {
				goto l17
			}
			v1 = i32(0xfff)
			goto l18
		}
	l17:
		v1 = i32(0)
	l19:
		{
			v1 = v1 + i32(1)
			t30 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			v0 = t30
			if v0 != 0 {
				goto l19
			}
		}
		p31 := i32(0xfff)
		if uint32(v1) > uint32(i32(0xfff)) {
			p31 = v1
		}
		v1 = p31
	}
l18:
	store32(m.memory[int64(uint32(i32(0)))+1303604:], uint32(v1))
	if uint32(v4) <= uint32(v2) {
		return
	}
	store32(m.memory[int64(uint32(i32(0)))+1303596:], uint32(i32(-1)))
}
func (m *Module) fn1559(v0, v1 int32) {
	var v2, v3, v4, v5 int32
	t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
	v2 = t0
	{
		{
			if uint32(v1) < uint32(i32(256)) {
				t19 := int32(load32(m.memory[int64(uint32(v0))+8:]))
				t20 := v2
				v4 = t19
				if t20 == v4 {
					t21 := int32(load32(m.memory[int64(uint32(i32(0)))+1303564:]))
					store32(m.memory[int64(uint32(i32(0)))+1303564:], uint32(t21&i32_rotl(i32(-2), int32(uint32(v1)>>3))))
					return
				}
				store32(m.memory[int64(uint32(v4))+12:], uint32(v2))
				store32(m.memory[int64(uint32(v2))+8:], uint32(v4))
				return
			}
			t1 := int32(load32(m.memory[int64(uint32(v0))+24:]))
			v3 = t1
			{
				{
					if v2 != v0 {
						t6 := int32(load32(m.memory[int64(uint32(v0))+8:]))
						v1 = t6
						store32(m.memory[int64(uint32(v1))+12:], uint32(v2))
						store32(m.memory[int64(uint32(v2))+8:], uint32(v1))
						goto l3
					}
					t2 := int32(load32(m.memory[int64(uint32(v0))+20:]))
					t3 := v0
					v2 = t2
					p4 := i32(16)
					if v2 != 0 {
						p4 = i32(20)
					}
					t5 := int32(load32(m.memory[uint32(t3+p4):]))
					v1 = t5
					if v1 != 0 {
						goto l2
					}
					v2 = i32(0)
					goto l3
				}
			l2:
				p7 := v0 + i32(16)
				if v2 != 0 {
					p7 = v0 + i32(20)
				}
				v4 = p7
			l4:
				{
					v5 = v4
					v2 = v1
					t8 := int32(load32(m.memory[int64(uint32(v2))+20:]))
					t9 := v2 + i32(20)
					t10 := v2 + i32(16)
					v1 = t8
					p11 := t10
					if v1 != 0 {
						p11 = t9
					}
					v4 = p11
					t13 := v2
					p12 := i32(16)
					if v1 != 0 {
						p12 = i32(20)
					}
					t14 := int32(load32(m.memory[uint32(t13+p12):]))
					v1 = t14
					if v1 != 0 {
						goto l4
					}
				}
				store32(m.memory[uint32(v5):], uint32(i32(0)))
			}
		l3:
			if v3 == 0 {
				return
			}
			{
				t15 := int32(load32(m.memory[int64(uint32(v0))+28:]))
				t16 := v0
				v1 = t15<<2 + i32(1303156)
				t17 := int32(load32(m.memory[uint32(v1):]))
				if t16 == t17 {
					store32(m.memory[uint32(v1):], uint32(v2))
					if v2 == 0 {
						goto l9
					}
					goto l8
				}
				t18 := int32(load32(m.memory[int64(uint32(v3))+16:]))
				if t18 == v0 {
					store32(m.memory[int64(uint32(v3))+16:], uint32(v2))
					if v2 != 0 {
						goto l8
					}
					return
				}
				store32(m.memory[int64(uint32(v3))+20:], uint32(v2))
				if v2 != 0 {
					goto l8
				}
				return
			}
		}
	l8:
		store32(m.memory[int64(uint32(v2))+24:], uint32(v3))
		{
			t22 := int32(load32(m.memory[int64(uint32(v0))+16:]))
			v1 = t22
			if v1 == 0 {
				goto l11
			}
			store32(m.memory[int64(uint32(v2))+16:], uint32(v1))
			store32(m.memory[int64(uint32(v1))+24:], uint32(v2))
		}
	l11:
		t23 := int32(load32(m.memory[int64(uint32(v0))+20:]))
		v1 = t23
		if v1 == 0 {
			return
		}
		store32(m.memory[int64(uint32(v2))+20:], uint32(v1))
		store32(m.memory[int64(uint32(v1))+24:], uint32(v2))
		return
	}
	return
l9:
	t24 := int32(load32(m.memory[int64(uint32(i32(0)))+1303568:]))
	t25 := int32(load32(m.memory[int64(uint32(v0))+28:]))
	store32(m.memory[int64(uint32(i32(0)))+1303568:], uint32(t24&i32_rotl(i32(-2), t25)))
}
func (m *Module) fn1560(v0, v1 int32) {
	var v2, v3 int32
	v2 = v0 + v1
	{
		{
			t0 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			v3 = t0
			if v3&i32(1) != 0 {
				goto l0
			}
			if v3&i32(2) == 0 {
				return
			}
			t1 := int32(load32(m.memory[uint32(v0):]))
			v3 = t1
			v1 = v3 + v1
			{
				v0 = v0 - v3
				t2 := int32(load32(m.memory[int64(uint32(i32(0)))+1303580:]))
				if v0 != t2 {
					goto l2
				}
				t3 := int32(load32(m.memory[int64(uint32(v2))+4:]))
				if t3&i32(3) != i32(3) {
					goto l0
				}
				store32(m.memory[int64(uint32(i32(0)))+1303572:], uint32(v1))
				t4 := int32(load32(m.memory[int64(uint32(v2))+4:]))
				store32(m.memory[int64(uint32(v2))+4:], uint32(t4&i32(-2)))
				store32(m.memory[int64(uint32(v0))+4:], uint32(v1|i32(1)))
				store32(m.memory[uint32(v2):], uint32(v1))
				return
			}
		l2:
			m.fn1559(v0, v3)
		}
	l0:
		{
			t5 := int32(load32(m.memory[int64(uint32(v2))+4:]))
			v3 = t5
			if v3&i32(2) != 0 {
				goto l3
			}
			t6 := int32(load32(m.memory[int64(uint32(i32(0)))+1303584:]))
			if v2 == t6 {
				goto l4
			}
			t7 := int32(load32(m.memory[int64(uint32(i32(0)))+1303580:]))
			if v2 == t7 {
				goto l5
			}
			t8 := v2
			v3 = v3 & i32(-8)
			m.fn1559(t8, v3)
			t9 := v0
			v1 = v3 + v1
			store32(m.memory[int64(uint32(t9))+4:], uint32(v1|i32(1)))
			store32(m.memory[uint32(v0+v1):], uint32(v1))
			t10 := int32(load32(m.memory[int64(uint32(i32(0)))+1303580:]))
			if v0 != t10 {
				goto l6
			}
			store32(m.memory[int64(uint32(i32(0)))+1303572:], uint32(v1))
			return
		}
	l3:
		store32(m.memory[int64(uint32(v2))+4:], uint32(v3&i32(-2)))
		store32(m.memory[int64(uint32(v0))+4:], uint32(v1|i32(1)))
		store32(m.memory[uint32(v0+v1):], uint32(v1))
	l6:
		if uint32(v1) < uint32(i32(256)) {
			{
				{
					t11 := int32(load32(m.memory[int64(uint32(i32(0)))+1303564:]))
					v2 = t11
					t12 := v2
					v3 = i32_shl(i32(1), int32(uint32(v1)>>3))
					if t12&v3 != 0 {
						goto l8
					}
					store32(m.memory[int64(uint32(i32(0)))+1303564:], uint32(v2|v3))
					v1 = v1&i32(248) + i32(1303300)
					v2 = v1
					goto l9
				}
			l8:
				v1 = v1 & i32(248)
				v2 = v1 + i32(1303300)
				t13 := int32(load32(m.memory[uint32(v1+i32(1303308)):]))
				v1 = t13
			}
		l9:
			store32(m.memory[int64(uint32(v2))+8:], uint32(v0))
			store32(m.memory[int64(uint32(v1))+12:], uint32(v0))
			store32(m.memory[int64(uint32(v0))+12:], uint32(v2))
			store32(m.memory[int64(uint32(v0))+8:], uint32(v1))
			return
		}
		m.fn1809(v0, v1)
		return
	l4:
		store32(m.memory[int64(uint32(i32(0)))+1303584:], uint32(v0))
		t14 := int32(load32(m.memory[int64(uint32(i32(0)))+1303576:]))
		v1 = t14 + v1
		store32(m.memory[int64(uint32(i32(0)))+1303576:], uint32(v1))
		store32(m.memory[int64(uint32(v0))+4:], uint32(v1|i32(1)))
		t15 := int32(load32(m.memory[int64(uint32(i32(0)))+1303580:]))
		if v0 != t15 {
			return
		}
		store32(m.memory[int64(uint32(i32(0)))+1303572:], uint32(i32(0)))
		store32(m.memory[int64(uint32(i32(0)))+1303580:], uint32(i32(0)))
	}
	return
l5:
	store32(m.memory[int64(uint32(i32(0)))+1303580:], uint32(v0))
	t16 := int32(load32(m.memory[int64(uint32(i32(0)))+1303572:]))
	v1 = t16 + v1
	store32(m.memory[int64(uint32(i32(0)))+1303572:], uint32(v1))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v1|i32(1)))
	store32(m.memory[uint32(v0+v1):], uint32(v1))
}
func (m *Module) fn1561(v0, v1 int32) int32 {
	var v2 int32
	{
		t0 := m.fn4(v0)
		v2 = t0
		if v2 == 0 {
			goto l0
		}
		t1 := int32(m.memory[uint32(v2+i32(-4))])
		if t1&i32(3) == 0 {
			goto l0
		}
		if v0 == 0 {
			goto l0
		}
		memory_zero(m.memory, uint32(v2), uint32(v0))
	}
l0:
	return v2
}
func (m *Module) fn1562(v0, v1, v2, v3 int32) int32 {
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
func (m *Module) fn1563(v0, v1 int32) {
	var v2, v3, v4, v5, v6, v7 int32
	var v8 int64
	t0 := m.g0
	v2 = t0 - i32(32)
	m.g0 = v2
	v3 = i32(0)
	{
		t1 := int32(m.memory[int64(uint32(v1))+12])
		if t1 != 0 {
			goto l0
		}
		t2 := int32(load32(m.memory[uint32(v1):]))
		v4 = t2
		t3 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		v5 = t3
		t4 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		v6 = t4
		v3 = i32(0)
		{
		l3:
			if v5 == v3 {
				goto l1
			}
			{
				t5 := int32(m.memory[uint32(v6+v3)])
				if t5 == i32(58) {
					t6 := v1
					t7 := v5
					v7 = v3 + i32(1)
					store32(m.memory[int64(uint32(t6))+8:], uint32(t7-v7))
					store32(m.memory[int64(uint32(v1))+4:], uint32(v6+v7))
					goto l4
				}
				v3 = v3 + i32(1)
				goto l3
			}
		}
	l1:
		m.memory[int64(uint32(v1))+12] = byte(i32(1))
		v3 = v5
	l4:
		m.fn1140(v2+i32(8), v6, v3)
		t8 := int64(load64(m.memory[int64(uint32(v2))+12:]))
		v8 = t8
		{
			t9 := int32(load32(m.memory[int64(uint32(v2))+8:]))
			v5 = t9
			if v5 == i32(-1) {
				goto l5
			}
			v3 = v2 + i32(20)
			{
				t10 := int32(load32(m.memory[uint32(v4):]))
				if t10 == i32(-1) {
					goto l6
				}
				m.fn1564(v4)
			}
		l6:
			store64(m.memory[int64(uint32(v4))+4:], uint64(v8))
			store32(m.memory[uint32(v4):], uint32(v5))
			t11 := int64(load64(m.memory[uint32(v3):]))
			store64(m.memory[int64(uint32(v4))+12:], uint64(t11))
			t12 := int32(load32(m.memory[int64(uint32(v3))+8:]))
			store32(m.memory[int64(uint32(v4))+20:], uint32(t12))
			v3 = i32(0)
			goto l0
		}
	l5:
		store64(m.memory[int64(uint32(v0))+4:], uint64(v8))
		v3 = i32(1)
	}
l0:
	store32(m.memory[uint32(v0):], uint32(v3))
	m.g0 = v2 + i32(32)
}
func (m *Module) fn1564(v0 int32) {
	var v1 int32
	{
		t0 := int32(load32(m.memory[uint32(v0):]))
		v1 = t0
		p1 := i32(3)
		if uint32(v1) > uint32(i32(-0x7ffffff2)) {
			p1 = v1 + i32(0x7ffffff1)
		}
		switch p1 {
		default:
			t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t3 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			m.fn16(t2, t3)
			return
		case 0:
			t4 := int32(m.memory[int64(uint32(v0))+4])
			t5 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			m.fn1565(t4, t5)
			return
		case 1:
			m.fn1566(v0 + i32(4))
			return
		case 2:
			m.fn1567(v0 + i32(4))
			return
		case 3:
			m.fn1568(v0)
			fallthrough
		case 4, 5, 6, 7, 8, 9, 11, 12, 13, 14, 15, 16, 17, 18, 20, 23, 27:
			return
		case 10:
			t6 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t7 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			m.fn16(t6, t7)
			return
		case 19:
			t8 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t9 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			m.fn16(t8, t9)
			return
		case 21:
			t10 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t11 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			m.fn16(t10, t11)
			return
		case 22:
			t12 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t13 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			m.fn16(t12, t13)
			return
		case 24:
			t14 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t15 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			m.fn16(t14, t15)
			return
		case 25:
			t16 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t17 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			m.fn16(t16, t17)
			return
		case 26:
			t18 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t19 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			m.fn16(t18, t19)
		}
	}
}
func (m *Module) fn1565(v0, v1 int32) {
	var v2 int32
	if v0&i32(255) != i32(3) {
		return
	}
	{
		t0 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		v0 = t0
		t1 := int32(load32(m.memory[uint32(v0):]))
		v2 = t1
		if v2 == 0 {
			goto l1
		}
		t2 := int32(load32(m.memory[uint32(v1):]))
		m.t0[uint(v2)].(func(int32))(t2)
	}
l1:
	{
		t3 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		v2 = t3
		if v2 == 0 {
			goto l2
		}
		t4 := int32(load32(m.memory[uint32(v1):]))
		t5 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		m.fn10(t4, v2, t5)
	}
l2:
	m.fn10(v1, i32(12), i32(4))
}
func (m *Module) fn1566(v0 int32) {
	var v1, v2 int32
	{
		t0 := int32(load32(m.memory[uint32(v0):]))
		v1 = t0
		v2 = v1 ^ i32(-0x80000000)
		p1 := i32(1)
		if uint32(v2) < uint32(i32(6)) {
			p1 = v2
		}
		switch p1 {
		case 0:
			t2 := int32(m.memory[int64(uint32(v0))+4])
			t3 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			m.fn1565(t2, t3)
			return
		case 1:
			t4 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			m.fn1390(v1, t4)
			fallthrough
		default:
		}
	}
}
func (m *Module) fn1567(v0 int32) {
	var v1 int32
	{
		t0 := int32(m.memory[uint32(v0)])
		v1 = t0
		p1 := i32(0)
		if uint32(v1) > uint32(i32(6)) {
			p1 = v1 + i32(-6)
		}
		switch p1 {
		case 0:
			m.fn1569(v0)
			return
		case 1:
			t2 := int32(m.memory[int64(uint32(v0))+4])
			t3 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			m.fn1565(t2, t3)
			return
		case 2:
			t4 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t5 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			m.fn16(t4, t5)
			fallthrough
		default:
		}
	}
}
func (m *Module) fn1568(v0 int32) {
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
				v1 = t10
				if v1 == i32(-1) {
					return
				}
				t11 := int32(load32(m.memory[int64(uint32(v0))+8:]))
				m.fn16(v1, t11)
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
func (m *Module) fn1569(v0 int32) {
	t0 := int32(m.memory[uint32(v0)])
	switch t0 {
	case 0:
		t1 := int32(m.memory[int64(uint32(v0))+4])
		t2 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		m.fn1565(t1, t2)
		return
	case 3:
		t3 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		t4 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		m.fn16(t3, t4)
		fallthrough
	default:
	}
}
func (m *Module) fn1570(v0, v1, v2 int32) {
	var v3 int32
	var v4 int64
	t0 := m.g0
	v3 = t0 - i32(48)
	m.g0 = v3
	store32(m.memory[int64(uint32(v3))+16:], uint32(v1))
	store32(m.memory[int64(uint32(v3))+12:], uint32(v0))
	store32(m.memory[int64(uint32(v3))+20:], uint32(v2))
	t1 := v3
	v4 = int64(uint32(i32(5))) << 32
	store64(m.memory[int64(uint32(t1))+40:], uint64(v4|int64(uint32(v3+i32(20)))))
	store64(m.memory[int64(uint32(v3))+32:], uint64(int64(uint32(i32(60)))<<32|int64(uint32(v3+i32(12)))))
	store64(m.memory[int64(uint32(v3))+24:], uint64(v4|int64(uint32(v3+i32(16)))))
	m.fn91(i32(1099908), v3+i32(24), i32(1099892))
	panic("unreachable")
}
func (m *Module) fn1571(v0, v1 int32) int32 {
	var v2 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	store32(m.memory[int64(uint32(v2))+12:], uint32(v0))
	t1 := m.fn264(v1, i32(1286950), i32(15), v2+i32(12), i32(195))
	v0 = t1
	m.g0 = v2 + i32(16)
	return v0
}
