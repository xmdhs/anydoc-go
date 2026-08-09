package core

func (m *Module) fn447(v0, v1, v2 int32) {
	var v3 int32
	{
		t0 := int32(load32(m.memory[uint32(v1):]))
		v3 = t0
		if v3 == 0 {
			goto l0
		}
		t1 := int32(load16(m.memory[int64(uint32(v1))+884:]))
		store32(m.memory[int64(uint32(v0))+8:], uint32(t1))
		store32(m.memory[int64(uint32(v0))+4:], uint32(v2+i32(1)))
	}
l0:
	store32(m.memory[uint32(v0):], uint32(v3))
	t3 := v1
	p2 := i32(888)
	if v2 != 0 {
		p2 = i32(936)
	}
	m.fn40(t3, i32(4), p2)
}
func (m *Module) fn448(v0 int32) {
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
	m.fn136(t4, v2, i32(4), i32(20))
}
func (m *Module) fn449(v0, v1 int32) {
	m.fn1301(v0, v1, i32(4), i32(4))
}
func (m *Module) fn450(v0, v1, v2, v3, v4, v5, v6 int32) {
	var v7, v8, v9 int32
	var v10 int64
	t0 := m.g0
	v7 = t0 - i32(32)
	m.g0 = v7
	v8 = i32(0)
	store32(m.memory[int64(uint32(v7))+8:], uint32(i32(0)))
	store64(m.memory[uint32(v7):], uint64(i64(0x100000000)))
	{
	l4:
		{
			if v2 == i32(-2) {
				goto l0
			}
			m.fn481(v7+i32(12), v1, v2, v5)
			{
				t1 := int32(m.memory[int64(uint32(v7))+12])
				v9 = t1
				if v9 == i32(255) {
					goto l1
				}
				t2 := int32(m.memory[int64(uint32(v7))+15])
				m.memory[int64(uint32(v0))+3] = byte(t2)
				t3 := int32(load16(m.memory[int64(uint32(v7))+13:]))
				store16(m.memory[int64(uint32(v0))+1:], uint16(t3))
				t4 := int64(load64(m.memory[int64(uint32(v7))+16:]))
				v10 = t4
				t5 := int64(load64(m.memory[int64(uint32(v7))+24:]))
				store64(m.memory[int64(uint32(v0))+12:], uint64(t5))
				store64(m.memory[int64(uint32(v0))+4:], uint64(v10))
				m.memory[uint32(v0)] = byte(v9)
				t6 := int32(load32(m.memory[uint32(v7):]))
				t7 := int32(load32(m.memory[int64(uint32(v7))+4:]))
				m.fn16(t6, t7)
				goto l2
			}
		l1:
			t8 := int32(load32(m.memory[int64(uint32(v7))+16:]))
			t9 := v7
			v8 = v6 - v8
			t10 := int32(load32(m.memory[int64(uint32(v7))+20:]))
			t11 := v8
			v9 = t10
			p12 := v9
			if uint32(v8) < uint32(v9) {
				p12 = t11
			}
			m.fn147(t9, t8, p12)
			t13 := int32(load32(m.memory[int64(uint32(v7))+8:]))
			v8 = t13
			if uint32(v8) >= uint32(v6) {
				goto l0
			}
			{
				if uint32(v2) >= uint32(v4) {
					goto l3
				}
				t14 := int32(load32(m.memory[uint32(v3+v2<<2):]))
				v2 = t14
				goto l4
			}
		l3:
		}
		m.fn158(v2, v4, i32(1070708))
		panic("unreachable")
	l0:
		t15 := int32(load32(m.memory[int64(uint32(v7))+8:]))
		store32(m.memory[int64(uint32(v0))+12:], uint32(t15))
		t16 := int64(load64(m.memory[uint32(v7):]))
		store64(m.memory[int64(uint32(v0))+4:], uint64(t16))
		m.memory[uint32(v0)] = byte(i32(255))
	}
l2:
	m.g0 = v7 + i32(32)
}
func (m *Module) fn451(v0 int32) {
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
			m.fn367(v0)
			return
		case 1:
			t2 := int32(m.memory[int64(uint32(v0))+4])
			t3 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			m.fn119(t2, t3)
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
func (m *Module) fn452(v0 int32) {
	var v1 int32
	t0 := m.g0
	v1 = t0 - i32(16)
	m.g0 = v1
	t1 := int32(load32(m.memory[uint32(v0):]))
	m.fn1844(v1+i32(8), v0, t1, i32(1), i32(2), i32(2))
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
func (m *Module) fn453(v0, v1, v2 int32) {
	var v3 int32
	{
		t0 := int32(load32(m.memory[uint32(v1):]))
		v3 = t0
		if v3 == 0 {
			goto l0
		}
		t1 := int32(load16(m.memory[int64(uint32(v1))+4:]))
		store32(m.memory[int64(uint32(v0))+8:], uint32(t1))
		store32(m.memory[int64(uint32(v0))+4:], uint32(v2+i32(1)))
	}
l0:
	store32(m.memory[uint32(v0):], uint32(v3))
	t3 := v1
	p2 := i32(44)
	if v2 != 0 {
		p2 = i32(92)
	}
	m.fn40(t3, i32(4), p2)
}
func (m *Module) fn454(v0, v1 int32) int32 {
	var v2, v3 int32
	t0 := m.g0
	v2 = t0 - i32(32)
	m.g0 = v2
	{
		t1 := int32(load32(m.memory[uint32(v0):]))
		v0 = t1
		t2 := int32(m.memory[uint32(v0)])
		switch t2 {
		default:
			t3 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			store32(m.memory[int64(uint32(v2))+4:], uint32(t3))
			t4 := int32(load32(m.memory[uint32(v1):]))
			t5 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t6 := int32(load32(m.memory[int64(uint32(t5))+12:]))
			t7 := m.t0[uint(t6)].(func(int32, int32, int32) int32)(t4, i32(1285220), i32(2))
			v0 = t7
			m.memory[int64(uint32(v2))+13] = byte(i32(0))
			m.memory[int64(uint32(v2))+12] = byte(v0)
			store32(m.memory[int64(uint32(v2))+8:], uint32(v1))
			t8 := m.fn1647(v2+i32(8), i32(1285222), i32(4), v2+i32(4), i32(78))
			v1 = t8
			m.memory[int64(uint32(v2))+19] = byte(i32(41))
			t9 := m.fn1647(v1, i32(1285226), i32(4), v2+i32(19), i32(79))
			v0 = t9
			t10 := m.fn4(i32(20))
			v1 = t10
			if v1 == 0 {
				m.fn2(i32(1), i32(20))
				panic("unreachable")
			}
			t11 := int32(load32(m.memory[int64(uint32(i32(0)))+1284656:]))
			store32(m.memory[int64(uint32(v1))+16:], uint32(t11))
			t12 := int64(load64(m.memory[int64(uint32(i32(0)))+1284648:]))
			store64(m.memory[int64(uint32(v1))+8:], uint64(t12))
			t13 := int64(load64(m.memory[int64(uint32(i32(0)))+1284640:]))
			store64(m.memory[uint32(v1):], uint64(t13))
			store32(m.memory[int64(uint32(v2))+28:], uint32(i32(20)))
			store32(m.memory[int64(uint32(v2))+24:], uint32(v1))
			store32(m.memory[int64(uint32(v2))+20:], uint32(i32(20)))
			t14 := m.fn1647(v0, i32(1285230), i32(7), v2+i32(20), i32(80))
			t15 := m.fn1648(t14)
			v1 = t15
			t16 := int32(load32(m.memory[int64(uint32(v2))+20:]))
			v0 = t16
			if v0 == 0 {
				goto l5
			}
			t17 := int32(load32(m.memory[int64(uint32(v2))+24:]))
			m.fn10(t17, v0, i32(1))
			goto l5
		case 1:
			t18 := int32(m.memory[int64(uint32(v0))+1])
			m.memory[int64(uint32(v2))+8] = byte(t18)
			t19 := int32(load32(m.memory[uint32(v1):]))
			t20 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t21 := int32(load32(m.memory[int64(uint32(t20))+12:]))
			t22 := m.t0[uint(t21)].(func(int32, int32, int32) int32)(t19, i32(1285237), i32(4))
			m.memory[int64(uint32(v2))+28] = byte(t22)
			store32(m.memory[int64(uint32(v2))+24:], uint32(v1))
			m.memory[int64(uint32(v2))+29] = byte(i32(0))
			store32(m.memory[int64(uint32(v2))+20:], uint32(i32(0)))
			t23 := m.fn1590(v2+i32(20), v2+i32(8), i32(79))
			t24 := m.fn1591(t23)
			v1 = t24
			goto l5
		case 2:
			t25 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			v0 = t25
			t26 := int32(load32(m.memory[uint32(v1):]))
			t27 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t28 := int32(load32(m.memory[int64(uint32(t27))+12:]))
			t29 := m.t0[uint(t28)].(func(int32, int32, int32) int32)(t26, i32(1285241), i32(5))
			v3 = t29
			m.memory[int64(uint32(v2))+25] = byte(i32(0))
			m.memory[int64(uint32(v2))+24] = byte(v3)
			store32(m.memory[int64(uint32(v2))+20:], uint32(v1))
			t30 := m.fn1647(v2+i32(20), i32(1285226), i32(4), v0+i32(8), i32(79))
			t31 := m.fn1647(t30, i32(1285230), i32(7), v0, i32(81))
			t32 := m.fn1648(t31)
			v1 = t32
			goto l5
		case 3:
			t33 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t34 := v2
			v0 = t33
			store32(m.memory[int64(uint32(t34))+20:], uint32(v0))
			t35 := m.fn459(v1, i32(1285246), i32(6), i32(1285226), i32(4), v0+i32(8), i32(79), i32(1285252), i32(5), v2+i32(20), i32(82))
			v1 = t35
		}
	}
l5:
	m.g0 = v2 + i32(32)
	return v1
}
func (m *Module) fn455(v0, v1 int32) int32 {
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
			store32(m.memory[int64(uint32(v2))+12:], uint32(v0+i32(4)))
			t3 := m.fn264(v1, i32(1100477), i32(2), v2+i32(12), i32(68))
			v0 = t3
			goto l7
		case 1:
			store32(m.memory[int64(uint32(v2))+12:], uint32(v0+i32(1)))
			t4 := m.fn459(v1, i32(1086816), i32(3), i32(1086819), i32(3), v0+i32(12), i32(73), i32(1086822), i32(9), v2+i32(12), i32(83))
			v0 = t4
			goto l7
		case 2:
			t5 := int32(load32(m.memory[uint32(v1):]))
			t6 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t7 := int32(load32(m.memory[int64(uint32(t6))+12:]))
			t8 := m.t0[uint(t7)].(func(int32, int32, int32) int32)(t5, i32(1086831), i32(12))
			v0 = t8
			goto l7
		case 3:
			store32(m.memory[int64(uint32(v2))+12:], uint32(v0+i32(4)))
			t9 := m.fn264(v1, i32(1086843), i32(14), v2+i32(12), i32(76))
			v0 = t9
			goto l7
		case 4:
			store32(m.memory[int64(uint32(v2))+12:], uint32(v0+i32(2)))
			t10 := m.fn462(v1, i32(1086857), i32(7), i32(1073713), i32(4), v0+i32(4), i32(71), i32(1086672), i32(8), v0+i32(12), i32(71), i32(1086680), i32(5), v2+i32(12), i32(77))
			v0 = t10
			goto l7
		case 5:
			store32(m.memory[int64(uint32(v2))+12:], uint32(v0+i32(2)))
			t11 := m.fn264(v1, i32(1086864), i32(16), v2+i32(12), i32(77))
			v0 = t11
			goto l7
		case 6:
			store32(m.memory[int64(uint32(v2))+12:], uint32(v0+i32(4)))
			t12 := m.fn264(v1, i32(1086880), i32(15), v2+i32(12), i32(84))
			v0 = t12
		}
	}
l7:
	m.g0 = v2 + i32(16)
	return v0
}
func (m *Module) fn456(v0, v1 int32) int32 {
	var v2, v3 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	{
		t1 := int32(load32(m.memory[uint32(v0):]))
		v0 = t1
		t2 := int32(m.memory[uint32(v0)])
		v3 = t2
		p3 := i32(0)
		if uint32(v3) > uint32(i32(6)) {
			p3 = v3 + i32(-6)
		}
		switch p3 {
		default:
			store32(m.memory[int64(uint32(v2))+12:], uint32(v0))
			t4 := m.fn264(v1, i32(1086625), i32(3), v2+i32(12), i32(69))
			v1 = t4
			goto l6
		case 1:
			store32(m.memory[int64(uint32(v2))+12:], uint32(v0+i32(4)))
			t5 := m.fn264(v1, i32(1100477), i32(2), v2+i32(12), i32(68))
			v1 = t5
			goto l6
		case 2:
			store32(m.memory[int64(uint32(v2))+12:], uint32(v0+i32(4)))
			t6 := m.fn264(v1, i32(1086628), i32(14), v2+i32(12), i32(76))
			v1 = t6
			goto l6
		case 3:
			store32(m.memory[int64(uint32(v2))+12:], uint32(v0+i32(12)))
			t7 := m.fn459(v1, i32(1086642), i32(7), i32(1086649), i32(3), v0+i32(4), i32(71), i32(1073156), i32(3), v2+i32(12), i32(77))
			v1 = t7
			goto l6
		case 4:
			t8 := int32(load32(m.memory[uint32(v1):]))
			t9 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t10 := int32(load32(m.memory[int64(uint32(t9))+12:]))
			t11 := m.t0[uint(t10)].(func(int32, int32, int32) int32)(t8, i32(1086652), i32(5))
			v1 = t11
			goto l6
		case 5:
			store32(m.memory[int64(uint32(v2))+12:], uint32(v0+i32(4)))
			t12 := m.fn459(v1, i32(1086657), i32(15), i32(1086672), i32(8), v0+i32(2), i32(85), i32(1086680), i32(5), v2+i32(12), i32(77))
			v1 = t12
		}
	}
l6:
	m.g0 = v2 + i32(16)
	return v1
}
func (m *Module) fn457(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t2 := int32(load32(m.memory[uint32(v1):]))
	t3 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	t4 := m.fn107(t0, t1, t2, t3)
	return t4
}
func (m *Module) fn458(v0, v1 int32) int32 {
	var v2 int32
	t0 := int32(load32(m.memory[uint32(v0):]))
	v0 = t0
	{
		t1 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		v2 = t1
		if v2&i32(0x2000000) != 0 {
			t4 := m.fn378(v0, v1)
			return t4
		}
		{
			if v2&i32(0x4000000) != 0 {
				t3 := m.fn1669(v0, v1)
				return t3
			}
			t2 := m.fn596(v0, v1)
			return t2
		}
	}
}
func (m *Module) fn459(v0, v1, v2, v3, v4, v5, v6, v7, v8, v9, v10 int32) int32 {
	var v11 int32
	t0 := m.g0
	v11 = t0 - i32(16)
	m.g0 = v11
	t1 := int32(load32(m.memory[uint32(v0):]))
	t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t3 := int32(load32(m.memory[int64(uint32(t2))+12:]))
	t4 := m.t0[uint(t3)].(func(int32, int32, int32) int32)(t1, v1, v2)
	v2 = t4
	m.memory[int64(uint32(v11))+13] = byte(i32(0))
	m.memory[int64(uint32(v11))+12] = byte(v2)
	store32(m.memory[int64(uint32(v11))+8:], uint32(v0))
	t5 := m.fn1647(v11+i32(8), v3, v4, v5, v6)
	t6 := m.fn1647(t5, v7, v8, v9, v10)
	v10 = t6
	t7 := int32(m.memory[int64(uint32(v11))+13])
	v2 = t7
	t8 := int32(m.memory[int64(uint32(v11))+12])
	t9 := v2
	v1 = t8
	v0 = t9 | v1
	{
		if v2 != i32(1) {
			goto l0
		}
		if v1&i32(1) != 0 {
			goto l0
		}
		{
			t10 := int32(load32(m.memory[uint32(v10):]))
			v0 = t10
			t11 := int32(m.memory[int64(uint32(v0))+10])
			if t11&i32(128) != 0 {
				goto l1
			}
			t12 := int32(load32(m.memory[uint32(v0):]))
			t13 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t14 := int32(load32(m.memory[int64(uint32(t13))+12:]))
			t15 := m.t0[uint(t14)].(func(int32, int32, int32) int32)(t12, i32(1283984), i32(2))
			v0 = t15
			goto l0
		}
	l1:
		t16 := int32(load32(m.memory[uint32(v0):]))
		t17 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		t18 := int32(load32(m.memory[int64(uint32(t17))+12:]))
		t19 := m.t0[uint(t18)].(func(int32, int32, int32) int32)(t16, i32(1108167), i32(1))
		v0 = t19
	}
l0:
	m.g0 = v11 + i32(16)
	return v0 & i32(1)
}
func (m *Module) fn460(v0, v1 int32) int32 {
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
func (m *Module) fn461(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := m.fn457(t0, v1)
	return t1
}
func (m *Module) fn462(v0, v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14 int32) int32 {
	var v15 int32
	t0 := m.g0
	v15 = t0 - i32(16)
	m.g0 = v15
	t1 := int32(load32(m.memory[uint32(v0):]))
	t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t3 := int32(load32(m.memory[int64(uint32(t2))+12:]))
	t4 := m.t0[uint(t3)].(func(int32, int32, int32) int32)(t1, v1, v2)
	v2 = t4
	m.memory[int64(uint32(v15))+13] = byte(i32(0))
	m.memory[int64(uint32(v15))+12] = byte(v2)
	store32(m.memory[int64(uint32(v15))+8:], uint32(v0))
	t5 := m.fn1647(v15+i32(8), v3, v4, v5, v6)
	t6 := m.fn1647(t5, v7, v8, v9, v10)
	t7 := m.fn1647(t6, v11, v12, v13, v14)
	v14 = t7
	t8 := int32(m.memory[int64(uint32(v15))+13])
	v2 = t8
	t9 := int32(m.memory[int64(uint32(v15))+12])
	t10 := v2
	v1 = t9
	v0 = t10 | v1
	{
		if v2 != i32(1) {
			goto l0
		}
		if v1&i32(1) != 0 {
			goto l0
		}
		{
			t11 := int32(load32(m.memory[uint32(v14):]))
			v0 = t11
			t12 := int32(m.memory[int64(uint32(v0))+10])
			if t12&i32(128) != 0 {
				goto l1
			}
			t13 := int32(load32(m.memory[uint32(v0):]))
			t14 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t15 := int32(load32(m.memory[int64(uint32(t14))+12:]))
			t16 := m.t0[uint(t15)].(func(int32, int32, int32) int32)(t13, i32(1283984), i32(2))
			v0 = t16
			goto l0
		}
	l1:
		t17 := int32(load32(m.memory[uint32(v0):]))
		t18 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		t19 := int32(load32(m.memory[int64(uint32(t18))+12:]))
		t20 := m.t0[uint(t19)].(func(int32, int32, int32) int32)(t17, i32(1108167), i32(1))
		v0 = t20
	}
l0:
	m.g0 = v15 + i32(16)
	return v0 & i32(1)
}
func (m *Module) fn463(v0, v1 int32) int32 {
	var v2 int32
	t0 := int32(load32(m.memory[uint32(v0):]))
	v0 = t0
	{
		t1 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		v2 = t1
		if v2&i32(0x2000000) != 0 {
			t5 := int32(load32(m.memory[uint32(v0):]))
			t6 := m.fn467(t5, v1)
			return t6
		}
		{
			if v2&i32(0x4000000) != 0 {
				t3 := int32(load32(m.memory[uint32(v0):]))
				t4 := m.fn466(t3, v1)
				return t4
			}
			t2 := m.fn72(v0, v1)
			return t2
		}
	}
}
func (m *Module) fn464(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	v0 = t0
	t1 := int32(load32(m.memory[uint32(v0+i32(4)):]))
	t2 := int32(load32(m.memory[uint32(v0+i32(8)):]))
	t3 := int32(load32(m.memory[uint32(v1):]))
	t4 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	t5 := m.fn107(t1, t2, t3, t4)
	return t5
}
func (m *Module) fn465(v0, v1 int32) int32 {
	var v2 int32
	t0 := int32(load32(m.memory[uint32(v0):]))
	v0 = t0
	{
		t1 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		v2 = t1
		if v2&i32(0x2000000) != 0 {
			t4 := int32(load16(m.memory[uint32(v0):]))
			t5 := m.fn470(t4, v1)
			return t5
		}
		{
			if v2&i32(0x4000000) != 0 {
				t3 := m.fn469(v0, v1)
				return t3
			}
			t2 := m.fn340(v0, v1)
			return t2
		}
	}
}
func (m *Module) fn466(v0, v1 int32) int32 {
	var v2, v3 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	v3 = i32(9)
l0:
	{
		t1 := int32(m.memory[int64(uint32(v0&i32(15)))+1131672])
		m.memory[uint32(v2+i32(8)+v3+i32(-2))] = byte(t1)
		v3 = v3 + i32(-1)
		v0 = int32(uint32(v0) >> 4)
		if v0 != 0 {
			goto l0
		}
	}
	t2 := m.fn1638(v1, i32(1), i32(1131670), i32(2), v2+i32(8)+v3+i32(-1), i32(9)-v3)
	v3 = t2
	m.g0 = v2 + i32(16)
	return v3
}
func (m *Module) fn467(v0, v1 int32) int32 {
	var v2, v3 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	v3 = i32(9)
l0:
	{
		t1 := int32(m.memory[int64(uint32(v0&i32(15)))+1107936])
		m.memory[uint32(v2+i32(8)+v3+i32(-2))] = byte(t1)
		v3 = v3 + i32(-1)
		v0 = int32(uint32(v0) >> 4)
		if v0 != 0 {
			goto l0
		}
	}
	t2 := m.fn1638(v1, i32(1), i32(1131670), i32(2), v2+i32(8)+v3+i32(-1), i32(9)-v3)
	v3 = t2
	m.g0 = v2 + i32(16)
	return v3
}
func (m *Module) fn468(v0, v1 int32) int32 {
	var v2 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		v2 = t0
		if v2&i32(0x2000000) != 0 {
			t3 := int32(load16(m.memory[uint32(v0):]))
			t4 := m.fn470(t3, v1)
			return t4
		}
		{
			if v2&i32(0x4000000) != 0 {
				t2 := m.fn469(v0, v1)
				return t2
			}
			t1 := m.fn340(v0, v1)
			return t1
		}
	}
}
func (m *Module) fn469(v0, v1 int32) int32 {
	var v2, v3 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	t1 := int32(load16(m.memory[uint32(v0):]))
	v3 = t1
	v0 = i32(5)
l0:
	{
		t2 := int32(m.memory[uint32(v3&i32(15)+i32(1131672))])
		m.memory[uint32(v2+i32(12)+v0+i32(-2))] = byte(t2)
		v0 = v0 + i32(-1)
		v3 = int32(uint32(v3)>>4) & i32(0xfff)
		if v3 != 0 {
			goto l0
		}
	}
	t3 := m.fn1638(v1, i32(1), i32(1131670), i32(2), v2+i32(12)+v0+i32(-1), i32(5)-v0)
	v0 = t3
	m.g0 = v2 + i32(16)
	return v0
}
func (m *Module) fn470(v0, v1 int32) int32 {
	var v2, v3 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	v3 = i32(5)
l0:
	{
		t1 := int32(m.memory[uint32(v0&i32(15)+i32(1107936))])
		m.memory[uint32(v2+i32(12)+v3+i32(-2))] = byte(t1)
		v3 = v3 + i32(-1)
		v0 = int32(uint32(v0)>>4) & i32(0xfff)
		if v0 != 0 {
			goto l0
		}
	}
	t2 := m.fn1638(v1, i32(1), i32(1131670), i32(2), v2+i32(12)+v3+i32(-1), i32(5)-v3)
	v3 = t2
	m.g0 = v2 + i32(16)
	return v3
}
func (m *Module) fn471(v0, v1 int32) int32 {
	var v2 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	{
		{
			t1 := int32(load32(m.memory[uint32(v0):]))
			v0 = t1
			t2 := int32(m.memory[uint32(v0)])
			if t2 != i32(1) {
				goto l0
			}
			store32(m.memory[int64(uint32(v2))+12:], uint32(v0+i32(1)))
			t3 := m.fn264(v1, i32(1087236), i32(4), v2+i32(12), i32(86))
			v1 = t3
			goto l1
		}
	l0:
		t4 := int32(load32(m.memory[uint32(v1):]))
		t5 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t6 := int32(load32(m.memory[int64(uint32(t5))+12:]))
		t7 := m.t0[uint(t6)].(func(int32, int32, int32) int32)(t4, i32(1087232), i32(4))
		v1 = t7
	}
l1:
	m.g0 = v2 + i32(16)
	return v1
}
func (m *Module) fn472(v0, v1 int32) int32 {
	var v2 int32
	t0 := int32(load32(m.memory[uint32(v0):]))
	v0 = t0
	{
		t1 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		v2 = t1
		if v2&i32(0x2000000) != 0 {
			t4 := int32(load32(m.memory[uint32(v0):]))
			t5 := m.fn467(t4, v1)
			return t5
		}
		{
			if v2&i32(0x4000000) != 0 {
				t3 := m.fn581(v0, v1)
				return t3
			}
			t2 := m.fn72(v0, v1)
			return t2
		}
	}
}
func (m *Module) fn473(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := m.fn474(t0, v1)
	return t1
}
func (m *Module) fn474(v0, v1 int32) int32 {
	t0 := m.fn475(v0, i32(8), v1)
	return t0
}
func (m *Module) fn475(v0, v1, v2 int32) int32 {
	var v3, v4 int32
	t0 := m.g0
	v3 = t0 - i32(16)
	m.g0 = v3
	t1 := int32(load32(m.memory[uint32(v2):]))
	t2 := int32(load32(m.memory[int64(uint32(v2))+4:]))
	t3 := int32(load32(m.memory[int64(uint32(t2))+12:]))
	t4 := m.t0[uint(t3)].(func(int32, int32, int32) int32)(t1, i32(1109520), i32(1))
	v4 = t4
	m.memory[int64(uint32(v3))+9] = byte(i32(0))
	m.memory[int64(uint32(v3))+8] = byte(v4)
	store32(m.memory[int64(uint32(v3))+4:], uint32(v2))
	if v1 == 0 {
		goto l0
	}
l1:
	store32(m.memory[int64(uint32(v3))+12:], uint32(v0))
	_ = m.fn1651(v3+i32(4), v3+i32(12), i32(72))
	v0 = v0 + i32(1)
	v1 = v1 + i32(-1)
	if v1 != 0 {
		goto l1
	}
l0:
	t6 := m.fn1652(v3 + i32(4))
	v0 = t6
	m.g0 = v3 + i32(16)
	return v0
}
func (m *Module) fn476(v0, v1, v2, v3 int32) {
	var v4 int32
	{
		t0 := v1
		v4 = v2 + i32(1)
		if uint32(t0) <= uint32(v4) {
			goto l0
		}
		v1 = (v1 + (v2 ^ i32(-1))) * i32(68)
		if v1 == 0 {
			goto l0
		}
		memory_copy(m.memory, uint32(v0+v4*i32(68)), uint32(v0+v2*i32(68)), uint32(v1))
	}
l0:
	memory_copy(m.memory, uint32(v0+v2*i32(68)), uint32(v3), uint32(i32(68)))
}
func (m *Module) fn477(v0, v1, v2, v3, v4 int32) {
	var v5, v6, v7, v8 int32
	t0 := m.g0
	v5 = t0 - i32(16)
	m.g0 = v5
	{
		{
			t1 := int32(load32(m.memory[uint32(v1):]))
			v6 = t1
			if v6 != 0 {
				goto l0
			}
			v6 = i32(0)
			v7 = v5 + i32(12)
			goto l1
		}
	l0:
		store32(m.memory[int64(uint32(v5))+12:], uint32(v3))
		v6 = v6 * v4
		t2 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		v8 = t2
		v7 = v5 + i32(8)
	}
l1:
	store32(m.memory[uint32(v7):], uint32(v6))
	{
		t3 := int32(load32(m.memory[int64(uint32(v5))+12:]))
		v6 = t3
		if v6 == 0 {
			goto l2
		}
		t4 := int32(load32(m.memory[int64(uint32(v5))+8:]))
		v7 = t4
		{
			if v2 != 0 {
				goto l3
			}
			m.fn40(v8, v6, v7)
			goto l4
		l3:
			t5 := v8
			t6 := v7
			t7 := v6
			v4 = v4 * v2
			t8 := m.fn89(t5, t6, t7, v4)
			v3 = t8
			if v3 == 0 {
				goto l5
			}
		}
	l4:
		store32(m.memory[uint32(v1):], uint32(v2))
		store32(m.memory[int64(uint32(v1))+4:], uint32(v3))
	}
l2:
	v6 = i32(-1)
l5:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v4))
	store32(m.memory[uint32(v0):], uint32(v6))
	m.g0 = v5 + i32(16)
}
func (m *Module) fn478(v0, v1, v2, v3 int32) {
	var v4 int32
	{
		t0 := v1
		v4 = v2 + i32(1)
		if uint32(t0) <= uint32(v4) {
			goto l0
		}
		v1 = (v1 + (v2 ^ i32(-1))) << 1
		if v1 == 0 {
			goto l0
		}
		memory_copy(m.memory, uint32(v0+v4<<1), uint32(v0+v2<<1), uint32(v1))
	}
l0:
	store16(m.memory[uint32(v0+v2<<1):], uint16(v3))
}
func (m *Module) fn479(v0, v1, v2, v3 int32) {
	if v1 != v3 {
		m.fn256(i32(1072679), i32(40), i32(1072720))
		panic("unreachable")
	}
	v1 = v1 << 1
	if v1 == 0 {
		return
	}
	memory_copy(m.memory, uint32(v2), uint32(v0), uint32(v1))
}
func (m *Module) fn480(v0, v1, v2, v3 int32) {
	if v1 != v3 {
		m.fn256(i32(1072679), i32(40), i32(1072720))
		panic("unreachable")
	}
	if v1 == 0 {
		return
	}
	memory_copy(m.memory, uint32(v2), uint32(v0), uint32(v1))
}
func (m *Module) fn481(v0, v1, v2, v3 int32) {
	var v4, v5, v6, v7, v8, v9 int32
	var v10 int64
	t0 := m.g0
	v4 = t0 - i32(32)
	m.g0 = v4
	{
		{
			t1 := int32(load32(m.memory[int64(uint32(v1))+16:]))
			if uint32(t1) < uint32(v2) {
				store32(m.memory[int64(uint32(v0))+4:], uint32(v2))
				m.memory[uint32(v0)] = byte(i32(6))
				goto l7
			}
			{
				t2 := int32(load32(m.memory[int64(uint32(v1))+12:]))
				v5 = t2
				v6 = v5 * v2
				v5 = v6 + v5
				t3 := int32(load32(m.memory[int64(uint32(v1))+8:]))
				t4 := v5
				v7 = t3
				if uint32(t4) > uint32(v7) {
					m.fn482(v1, v5)
					t6 := int32(load32(m.memory[int64(uint32(v1))+8:]))
					v9 = t6
					t7 := int32(load32(m.memory[int64(uint32(v1))+4:]))
					v8 = t7
				l8:
					{
						v1 = v7
						if uint32(v1) < uint32(v5) {
							goto l3
						}
						v7 = v9
						goto l2
					l3:
						m.fn364(v4+i32(8), v1, v5, v8, v9, i32(1070676))
						t8 := int32(load32(m.memory[int64(uint32(v4))+8:]))
						t9 := int32(load32(m.memory[int64(uint32(v4))+12:]))
						m.fn301(v4+i32(24), v3, t8, t9)
						{
							t10 := int32(m.memory[int64(uint32(v4))+24])
							if t10 != i32(255) {
								t12 := int64(load64(m.memory[int64(uint32(v4))+24:]))
								v10 = t12
								if v10&i64(255) != i64(255) {
									store64(m.memory[int64(uint32(v0))+4:], uint64(v10))
									m.memory[uint32(v0)] = byte(i32(0))
									goto l7
								}
								v2 = int32(int64(uint64(v10) >> 32))
								goto l5
							}
							t11 := int32(load32(m.memory[int64(uint32(v4))+28:]))
							v2 = t11
							goto l5
						}
					l5:
						v7 = v2 + v1
						if v2 != 0 {
							goto l8
						}
					}
					m.fn483(v4, v8, v9, v6, v1, i32(1070692))
					t13 := int64(load64(m.memory[uint32(v4):]))
					store64(m.memory[int64(uint32(v0))+4:], uint64(t13))
					m.memory[uint32(v0)] = byte(i32(255))
					goto l7
				}
				t5 := int32(load32(m.memory[int64(uint32(v1))+4:]))
				v8 = t5
				goto l2
			}
		}
	l2:
		m.fn483(v4+i32(16), v8, v7, v6, v5, i32(1070660))
		t14 := int64(load64(m.memory[int64(uint32(v4))+16:]))
		store64(m.memory[int64(uint32(v0))+4:], uint64(t14))
		m.memory[uint32(v0)] = byte(i32(255))
	}
l7:
	m.g0 = v4 + i32(32)
}
func (m *Module) fn482(v0, v1 int32) {
	var v2, v3, v4 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		t1 := v1
		v2 = t0
		if uint32(t1) <= uint32(v2) {
			goto l0
		}
		v3 = v2
		{
			v4 = v1 - v2
			t2 := int32(load32(m.memory[uint32(v0):]))
			if uint32(v4) <= uint32(t2-v2) {
				goto l1
			}
			m.fn1684(v0, v2, v4, i32(1), i32(1))
			t3 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			v3 = t3
		}
	l1:
		v1 = v2 ^ i32(-1) + v1
		t4 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		v2 = t4 + v3
	l3:
		m.memory[uint32(v2)] = byte(i32(0))
		if v1 == 0 {
			goto l2
		}
		v1 = v1 + i32(-1)
		v2 = v2 + i32(1)
		goto l3
	l2:
		v1 = v3 + v4
	}
l0:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v1))
}
func (m *Module) fn483(v0, v1, v2, v3, v4, v5 int32) {
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
	store32(m.memory[uint32(v0):], uint32(v1+v3))
}
func (m *Module) fn484(v0, v1, v2 int32) {
	var v3, v4, v5, v6 int32
	var v7 int64
	var v8, v9, v10, v11, v12, v13, v14, v15, v16, v17, v18, v19, v20, v21, v22, v23 int32
	t0 := m.g0
	v3 = t0 - i32(4208)
	m.g0 = v3
	v4 = i32(0)
	memory_zero(m.memory, uint32(v3+i32(96)), uint32(i32(512)))
	{
	l4:
		{
			t1 := v3 + i32(24)
			v5 = v4
			m.fn212(t1, v5, v3+i32(96), i32(512), i32(1070592))
			t2 := int32(load32(m.memory[int64(uint32(v3))+24:]))
			t3 := int32(load32(m.memory[int64(uint32(v3))+28:]))
			m.fn301(v3+i32(608), v1, t2, t3)
			{
				t4 := int32(m.memory[int64(uint32(v3))+608])
				if t4 != i32(255) {
					t6 := int64(load64(m.memory[int64(uint32(v3))+608:]))
					v7 = t6
					if v7&i64(255) != i64(255) {
						t7 := int32(load32(m.memory[int64(uint32(v3))+612:]))
						v8 = t7
						t8 := int32(load32(m.memory[int64(uint32(v3))+608:]))
						v9 = t8
						goto l3
					}
					v6 = int32(int64(uint64(v7) >> 32))
					goto l1
				}
				t5 := int32(load32(m.memory[int64(uint32(v3))+612:]))
				v6 = t5
				goto l1
			}
		l1:
			v4 = v6 + v5
			if v6 != 0 {
				goto l4
			}
		}
		t9 := int64(load64(m.memory[int64(uint32(v3))+96:]))
		v7 = t9
		{
			{
				if uint32(v5) < uint32(i32(512)) {
					goto l5
				}
				if v7 != i64(-0x1ee54e5e1fee3030) {
					goto l5
				}
				v10 = i32(512)
				t10 := int32(load16(m.memory[int64(uint32(v3))+122:]))
				v11 = t10
				{
					t11 := int32(load16(m.memory[int64(uint32(v3))+126:]))
					v6 = t11
					switch v6 + i32(-9) {
					default:
						v12 = v6<<16 | i32(4)
						v9 = i32(1070636)
						v5 = i32(1070648)
						v6 = i32(12)
						v8 = i32(12)
						goto l9
					case 3:
						m.fn117(v3+i32(48), v1, v3+i32(608), i32(3584))
						{
							t12 := int32(m.memory[int64(uint32(v3))+48])
							if t12 == i32(255) {
								goto l10
							}
							t13 := int64(m.memory[int64(uint32(v3))+48])
							if t13 == i64(255) {
								goto l10
							}
							t14 := int32(load32(m.memory[int64(uint32(v3))+52:]))
							v8 = t14
							t15 := int32(load32(m.memory[int64(uint32(v3))+48:]))
							v9 = t15
							goto l3
						}
					l10:
						v10 = i32(4096)
						fallthrough
					case 0:
						t16 := int32(load16(m.memory[int64(uint32(v3))+128:]))
						v5 = t16
						if v5 == i32(6) {
							t21 := int32(load32(m.memory[int64(uint32(v3))+164:]))
							v6 = t21
							t22 := int32(load32(m.memory[int64(uint32(v3))+160:]))
							v8 = t22
							t23 := int32(load32(m.memory[int64(uint32(v3))+156:]))
							v5 = t23
							t24 := int32(load32(m.memory[int64(uint32(v3))+144:]))
							v12 = t24
							t25 := int32(load32(m.memory[int64(uint32(v3))+140:]))
							v9 = t25
							t26 := int32(load32(m.memory[int64(uint32(v3))+158:]))
							m.fn485(v3+i32(48), t26)
							store64(m.memory[int64(uint32(v3))+620:], uint64(i64(0x400000000)))
							store32(m.memory[int64(uint32(v3))+616:], uint32(v3+i32(608)))
							store32(m.memory[int64(uint32(v3))+612:], uint32(i32(436)))
							store32(m.memory[int64(uint32(v3))+608:], uint32(v3+i32(172)))
							m.fn486(v3+i32(48), v3+i32(608))
							t27 := int32(load32(m.memory[int64(uint32(v3))+48:]))
							v13 = t27
							if v13 == i32(-1) {
								goto l9
							}
							t28 := int64(load64(m.memory[int64(uint32(v3))+52:]))
							v7 = t28
							store32(m.memory[int64(uint32(v3))+36:], uint32(v13))
							store64(m.memory[int64(uint32(v3))+40:], uint64(v7))
							m.fn140(v3+i32(48), i32(1024))
							t29 := int32(uint32(v2) / uint32(v10))
							t30 := v3
							v14 = t29
							store32(m.memory[int64(uint32(t30))+64:], uint32(v14))
							store32(m.memory[int64(uint32(v3))+60:], uint32(v10))
							v4 = int32(int64(uint64(v7) >> 32))
							v2 = int32(v7)
						l49:
							{
								{
									if uint32(v6) > uint32(i32(-7)) {
										m.fn485(v3+i32(68), v9)
										v6 = v4 << 2
										v4 = v2
									l47:
										{
											{
												{
													if v6 == 0 {
														m.fn487(v2, v13)
														t39 := int32(load32(m.memory[int64(uint32(v3))+72:]))
														t40 := v3 + i32(608)
														t41 := v3 + i32(48)
														t42 := v12
														v15 = t39
														t43 := int32(load32(m.memory[int64(uint32(v3))+76:]))
														t44 := v15
														v16 = t43
														m.fn450(t40, t41, t42, t44, v16, v1, i32(-1))
														{
															t45 := int32(m.memory[int64(uint32(v3))+608])
															v6 = t45
															if v6 == i32(255) {
																t50 := int32(load32(m.memory[int64(uint32(v3))+616:]))
																v17 = t50
																t51 := int32(load32(m.memory[int64(uint32(v3))+612:]))
																v18 = t51
																t52 := int32(load32(m.memory[int64(uint32(v3))+620:]))
																t53 := v3 + i32(608)
																v13 = t52
																m.fn488(t53, v13, i32(128))
																{
																	t54 := int32(load32(m.memory[int64(uint32(v3))+612:]))
																	if t54 != i32(1) {
																		m.fn91(i32(1087526), i32(35), i32(1100680))
																		panic("unreachable")
																	}
																	t55 := int32(load32(m.memory[int64(uint32(v3))+616:]))
																	m.fn59(v3+i32(16), t55, i32(4), i32(20))
																	v2 = i32(0)
																	store32(m.memory[int64(uint32(v3))+4204:], uint32(i32(0)))
																	t56 := int32(load32(m.memory[int64(uint32(v3))+20:]))
																	t57 := v3
																	v19 = t56
																	store32(m.memory[int64(uint32(t57))+4200:], uint32(v19))
																	t58 := int32(load32(m.memory[int64(uint32(v3))+16:]))
																	t59 := v3
																	v6 = t58
																	store32(m.memory[int64(uint32(t59))+4196:], uint32(v6))
																	m.fn488(v3+i32(608), v13, i32(128))
																	{
																		t60 := int32(load32(m.memory[int64(uint32(v3))+612:]))
																		if t60 != i32(1) {
																			m.fn91(i32(1087526), i32(35), i32(1087544))
																			panic("unreachable")
																		}
																		{
																			t61 := int32(load32(m.memory[int64(uint32(v3))+616:]))
																			v4 = t61
																			if uint32(v4) <= uint32(v6) {
																				goto l22
																			}
																			m.fn62(v3+i32(4196), i32(0), v4, i32(4), i32(20))
																			t62 := int32(load32(m.memory[int64(uint32(v3))+4204:]))
																			v2 = t62
																			t63 := int32(load32(m.memory[int64(uint32(v3))+4200:]))
																			v19 = t63
																		}
																	l22:
																		;
																		var p64 int32
																		if v10 == i32(512) {
																			p64 = 1
																		}
																		v20 = p64
																		v21 = v17
																	l37:
																		{
																			if v13 == 0 {
																				t85 := int64(load64(m.memory[int64(uint32(v3))+4196:]))
																				store64(m.memory[int64(uint32(v3))+80:], uint64(t85))
																				store32(m.memory[int64(uint32(v3))+88:], uint32(v2))
																				{
																					if v2 == 0 {
																						goto l38
																					}
																					if v11&i32(0xffff) == i32(3) {
																						goto l39
																					}
																					t86 := int32(load32(m.memory[int64(uint32(v3))+84:]))
																					t87 := m.fn412(t86, v2, i32(1070528))
																					t88 := int32(load32(m.memory[int64(uint32(t87))+12:]))
																					if t88 != i32(-2) {
																						goto l39
																					}
																				}
																			l38:
																				store32(m.memory[uint32(v0):], uint32(i32(-1)))
																				m.memory[int64(uint32(v0))+4] = byte(i32(2))
																				goto l40
																			l39:
																				if v8 != 0 {
																					t89 := int32(load32(m.memory[int64(uint32(v3))+84:]))
																					t90 := m.fn412(t89, v2, i32(1070544))
																					t91 := v3 + i32(608)
																					t92 := v3 + i32(48)
																					v6 = t90
																					t93 := int32(load32(m.memory[int64(uint32(v6))+12:]))
																					t94 := m.fn412(v6, v2, i32(1070560))
																					t95 := int32(load32(m.memory[int64(uint32(t94))+16:]))
																					m.fn450(t91, t92, t93, v15, v16, v1, t95)
																					{
																						t96 := int32(m.memory[int64(uint32(v3))+608])
																						v6 = t96
																						if v6 == i32(255) {
																							t101 := int32(load32(m.memory[int64(uint32(v3))+620:]))
																							v9 = t101
																							t102 := int32(load32(m.memory[int64(uint32(v3))+616:]))
																							v6 = t102
																							t103 := int32(load32(m.memory[int64(uint32(v3))+612:]))
																							v4 = t103
																							m.fn450(v3+i32(608), v3+i32(48), v5, v15, v16, v1, v8*v10)
																							{
																								t104 := int32(m.memory[int64(uint32(v3))+608])
																								v5 = t104
																								if v5 == i32(255) {
																									t109 := int32(load32(m.memory[int64(uint32(v3))+612:]))
																									v8 = t109
																									t110 := int32(load32(m.memory[int64(uint32(v3))+616:]))
																									t111 := v3 + i32(608)
																									v13 = t110
																									t112 := int32(load32(m.memory[int64(uint32(v3))+620:]))
																									m.fn491(t111, v13, t112)
																									t113 := int32(load32(m.memory[int64(uint32(v3))+612:]))
																									t114 := int32(load32(m.memory[int64(uint32(v3))+624:]))
																									m.fn492(v3+i32(96), t113, t114)
																									t115 := int32(load32(m.memory[int64(uint32(v3))+100:]))
																									if t115 == 0 {
																										m.fn91(i32(1087526), i32(35), i32(1100680))
																										panic("unreachable")
																									}
																									t116 := int32(load32(m.memory[int64(uint32(v3))+104:]))
																									m.fn59(v3+i32(8), t116, i32(4), i32(4))
																									store32(m.memory[int64(uint32(v3))+4204:], uint32(i32(0)))
																									t117 := int64(load64(m.memory[int64(uint32(v3))+8:]))
																									store64(m.memory[int64(uint32(v3))+4196:], uint64(t117))
																									m.fn486(v3+i32(4196), v3+i32(608))
																									t118 := int32(load32(m.memory[int64(uint32(v3))+4204:]))
																									v12 = t118
																									t119 := int32(load32(m.memory[int64(uint32(v3))+4200:]))
																									v5 = t119
																									t120 := int32(load32(m.memory[int64(uint32(v3))+4196:]))
																									v1 = t120
																									m.fn16(v8, v13)
																									goto l42
																								}
																								t105 := int32(m.memory[int64(uint32(v3))+611])
																								m.memory[int64(uint32(v0))+7] = byte(t105)
																								t106 := int32(load16(m.memory[int64(uint32(v3))+609:]))
																								store16(m.memory[int64(uint32(v0))+5:], uint16(t106))
																								t107 := int64(load64(m.memory[int64(uint32(v3))+620:]))
																								v7 = t107
																								t108 := int64(load64(m.memory[int64(uint32(v3))+612:]))
																								store64(m.memory[int64(uint32(v0))+8:], uint64(t108))
																								store64(m.memory[int64(uint32(v0))+16:], uint64(v7))
																								m.memory[int64(uint32(v0))+4] = byte(v5)
																								store32(m.memory[uint32(v0):], uint32(i32(-1)))
																								m.fn16(v4, v6)
																								goto l40
																							}
																						}
																						t97 := int32(m.memory[int64(uint32(v3))+611])
																						m.memory[int64(uint32(v0))+7] = byte(t97)
																						t98 := int32(load16(m.memory[int64(uint32(v3))+609:]))
																						store16(m.memory[int64(uint32(v0))+5:], uint16(t98))
																						t99 := int64(load64(m.memory[int64(uint32(v3))+620:]))
																						v7 = t99
																						t100 := int64(load64(m.memory[int64(uint32(v3))+612:]))
																						store64(m.memory[int64(uint32(v0))+8:], uint64(t100))
																						store64(m.memory[int64(uint32(v0))+16:], uint64(v7))
																						m.memory[int64(uint32(v0))+4] = byte(v6)
																						store32(m.memory[uint32(v0):], uint32(i32(-1)))
																						goto l40
																					}
																				}
																				v6 = i32(1)
																				v5 = i32(4)
																				v1 = i32(0)
																				v12 = i32(0)
																				v4 = i32(0)
																				v9 = i32(0)
																				goto l42
																			}
																			t66 := v3 + i32(608)
																			t67 := v21
																			t68 := v13
																			p65 := i32(128)
																			if uint32(v13) < uint32(i32(128)) {
																				p65 = v13
																			}
																			m.fn309(t66, t67, t68, p65, i32(1100308))
																			t69 := int32(load32(m.memory[int64(uint32(v3))+612:]))
																			v22 = t69
																			if uint32(v22) <= uint32(i32(63)) {
																				m.fn151(i32(0), i32(64), v22, i32(1088008))
																				panic("unreachable")
																			}
																			t70 := int32(load32(m.memory[int64(uint32(v3))+620:]))
																			v13 = t70
																			t71 := int32(load32(m.memory[int64(uint32(v3))+616:]))
																			v21 = t71
																			t72 := int32(load32(m.memory[int64(uint32(v3))+608:]))
																			t73 := v3 + i32(608)
																			v23 = t72
																			m.fn489(t73, i32(1153092), v23, i32(64))
																			m.fn490(v3+i32(96), v3+i32(608))
																			v6 = i32(0)
																			t74 := int32(load32(m.memory[int64(uint32(v3))+104:]))
																			v4 = t74
																			t75 := int32(load32(m.memory[int64(uint32(v3))+100:]))
																			v9 = t75
																		l27:
																			if v4 == v6 {
																				goto l25
																			}
																			{
																				v12 = v9 + v6
																				t76 := int32(m.memory[uint32(v12)])
																				if t76 == 0 {
																					goto l26
																				}
																				v6 = v6 + i32(1)
																				goto l27
																			}
																		l26:
																			{
																				if v6 != 0 {
																					goto l28
																				}
																				v6 = i32(0)
																				goto l29
																			l28:
																				t77 := int32(int8(m.memory[uint32(v12)]))
																				if t77 <= i32(-65) {
																					m.fn256(i32(1087836), i32(48), i32(1088024))
																					panic("unreachable")
																				}
																			}
																		l29:
																			store32(m.memory[int64(uint32(v3))+104:], uint32(v6))
																		l25:
																			if uint32(v22) <= uint32(i32(119)) {
																				m.fn151(i32(116), i32(120), v22, i32(1088040))
																				panic("unreachable")
																			}
																			t78 := int32(load32(m.memory[int64(uint32(v23))+116:]))
																			v4 = t78
																			{
																				{
																					if v20 != 0 {
																						goto l32
																					}
																					if uint32(v22) <= uint32(i32(127)) {
																						m.fn151(i32(120), i32(128), v22, i32(1088056))
																						panic("unreachable")
																					}
																					t79 := int64(load64(m.memory[int64(uint32(v23))+120:]))
																					v7 = t79
																					if uint64(v7) >= uint64(i64(0x100000000)) {
																						m.memory[int64(uint32(v3))+608] = byte(i32(2))
																						m.fn97(i32(1291936), i32(43), v3+i32(608), i32(1087760), i32(1088072))
																						panic("unreachable")
																					}
																					v9 = int32(v7)
																					goto l35
																				}
																			l32:
																				if uint32(v22) <= uint32(i32(123)) {
																					m.fn151(i32(120), i32(124), v22, i32(1088088))
																					panic("unreachable")
																				}
																				t80 := int32(load32(m.memory[int64(uint32(v23))+120:]))
																				v9 = t80
																			}
																		l35:
																			t81 := int32(load32(m.memory[int64(uint32(v3))+104:]))
																			t82 := v3
																			v12 = t81
																			store32(m.memory[int64(uint32(t82))+616:], uint32(v12))
																			t83 := int64(load64(m.memory[int64(uint32(v3))+96:]))
																			t84 := v3
																			v7 = t83
																			store64(m.memory[int64(uint32(t84))+608:], uint64(v7))
																			v6 = v19 + v2*i32(20)
																			store32(m.memory[int64(uint32(v6))+8:], uint32(v12))
																			store64(m.memory[uint32(v6):], uint64(v7))
																			store32(m.memory[int64(uint32(v6))+16:], uint32(v9))
																			store32(m.memory[int64(uint32(v6))+12:], uint32(v4))
																			v2 = v2 + i32(1)
																			goto l37
																		}
																	}
																}
															}
															t46 := int32(m.memory[int64(uint32(v3))+611])
															m.memory[int64(uint32(v0))+7] = byte(t46)
															t47 := int32(load16(m.memory[int64(uint32(v3))+609:]))
															store16(m.memory[int64(uint32(v0))+5:], uint16(t47))
															t48 := int64(load64(m.memory[int64(uint32(v3))+620:]))
															v7 = t48
															t49 := int64(load64(m.memory[int64(uint32(v3))+612:]))
															store64(m.memory[int64(uint32(v0))+8:], uint64(t49))
															store64(m.memory[int64(uint32(v0))+16:], uint64(v7))
															m.memory[int64(uint32(v0))+4] = byte(v6)
															store32(m.memory[uint32(v0):], uint32(i32(-1)))
															goto l19
														}
													}
													t38 := int32(load32(m.memory[uint32(v4):]))
													v9 = t38
													if uint32(v9) < uint32(i32(-4)) {
														m.fn481(v3+i32(608), v3+i32(48), v9, v1)
														{
															t121 := int32(m.memory[int64(uint32(v3))+608])
															v9 = t121
															if v9 == i32(255) {
																t127 := int32(load32(m.memory[int64(uint32(v3))+612:]))
																t128 := int32(load32(m.memory[int64(uint32(v3))+616:]))
																m.fn491(v3+i32(96), t127, t128)
																m.fn486(v3+i32(68), v3+i32(96))
																goto l17
															}
															t122 := int32(m.memory[int64(uint32(v3))+611])
															m.memory[int64(uint32(v0))+7] = byte(t122)
															t123 := int32(load16(m.memory[int64(uint32(v3))+609:]))
															store16(m.memory[int64(uint32(v0))+5:], uint16(t123))
															t124 := int64(load64(m.memory[int64(uint32(v3))+612:]))
															v7 = t124
															t125 := int64(load64(m.memory[int64(uint32(v3))+620:]))
															store64(m.memory[int64(uint32(v0))+16:], uint64(t125))
															store64(m.memory[int64(uint32(v0))+8:], uint64(v7))
															m.memory[int64(uint32(v0))+4] = byte(v9)
															store32(m.memory[uint32(v0):], uint32(i32(-1)))
															m.fn487(v2, v13)
															t126 := int32(load32(m.memory[int64(uint32(v3))+72:]))
															v15 = t126
															goto l19
														}
													}
													goto l17
												}
											l40:
												m.fn448(v3 + i32(80))
												m.fn16(v18, v17)
											l19:
												t129 := int32(load32(m.memory[int64(uint32(v3))+68:]))
												m.fn449(t129, v15)
												t130 := int32(load32(m.memory[int64(uint32(v3))+48:]))
												t131 := int32(load32(m.memory[int64(uint32(v3))+52:]))
												m.fn16(t130, t131)
												goto l14
											}
										l42:
											t132 := int32(load32(m.memory[int64(uint32(v3))+88:]))
											store32(m.memory[int64(uint32(v3))+616:], uint32(t132))
											t133 := int64(load64(m.memory[int64(uint32(v3))+80:]))
											store64(m.memory[int64(uint32(v3))+608:], uint64(t133))
											t134 := int64(load64(m.memory[int64(uint32(v3))+48:]))
											store64(m.memory[int64(uint32(v3))+620:], uint64(t134))
											t135 := int64(load64(m.memory[int64(uint32(v3))+56:]))
											store64(m.memory[int64(uint32(v3))+628:], uint64(t135))
											t136 := int32(load32(m.memory[int64(uint32(v3))+64:]))
											store32(m.memory[int64(uint32(v3))+636:], uint32(t136))
											t137 := int64(load64(m.memory[int64(uint32(v3))+68:]))
											store64(m.memory[int64(uint32(v3))+640:], uint64(t137))
											t138 := int32(load32(m.memory[int64(uint32(v3))+76:]))
											store32(m.memory[int64(uint32(v3))+648:], uint32(t138))
											memory_copy(m.memory, uint32(v0), uint32(v3+i32(608)), uint32(i32(44)))
											store32(m.memory[int64(uint32(v0))+72:], uint32(v12))
											store32(m.memory[int64(uint32(v0))+68:], uint32(v5))
											store32(m.memory[int64(uint32(v0))+64:], uint32(v1))
											store32(m.memory[int64(uint32(v0))+60:], uint32(v14<<3))
											store32(m.memory[int64(uint32(v0))+56:], uint32(i32(64)))
											store32(m.memory[int64(uint32(v0))+52:], uint32(v9))
											store32(m.memory[int64(uint32(v0))+48:], uint32(v6))
											store32(m.memory[int64(uint32(v0))+44:], uint32(v4))
											m.fn16(v18, v17)
											goto l14
										}
									l17:
										v4 = v4 + i32(4)
										v6 = v6 + i32(-4)
										goto l47
									}
									m.fn481(v3+i32(608), v3+i32(48), v6, v1)
									t31 := int32(m.memory[int64(uint32(v3))+608])
									v6 = t31
									if v6 == i32(255) {
										goto l13
									}
									t32 := int32(m.memory[int64(uint32(v3))+611])
									m.memory[int64(uint32(v0))+7] = byte(t32)
									t33 := int32(load16(m.memory[int64(uint32(v3))+609:]))
									store16(m.memory[int64(uint32(v0))+5:], uint16(t33))
									t34 := int64(load64(m.memory[int64(uint32(v3))+612:]))
									v7 = t34
									t35 := int64(load64(m.memory[int64(uint32(v3))+620:]))
									store64(m.memory[int64(uint32(v0))+16:], uint64(t35))
									store64(m.memory[int64(uint32(v0))+8:], uint64(v7))
									m.memory[int64(uint32(v0))+4] = byte(v6)
									store32(m.memory[uint32(v0):], uint32(i32(-1)))
									t36 := int32(load32(m.memory[int64(uint32(v3))+48:]))
									t37 := int32(load32(m.memory[int64(uint32(v3))+52:]))
									m.fn16(t36, t37)
									m.fn449(v13, v2)
									goto l14
								}
							l13:
								t139 := int32(load32(m.memory[int64(uint32(v3))+612:]))
								t140 := int32(load32(m.memory[int64(uint32(v3))+616:]))
								m.fn491(v3+i32(96), t139, t140)
								m.fn486(v3+i32(36), v3+i32(96))
								{
									t141 := int32(load32(m.memory[int64(uint32(v3))+44:]))
									v6 = t141
									if v6 == 0 {
										goto l48
									}
									t142 := v3
									v4 = v6 + i32(-1)
									store32(m.memory[int64(uint32(t142))+44:], uint32(v4))
									t143 := int32(load32(m.memory[int64(uint32(v3))+40:]))
									v2 = t143
									t144 := int32(load32(m.memory[uint32(v2+v4<<2):]))
									v6 = t144
									t145 := int32(load32(m.memory[int64(uint32(v3))+36:]))
									v13 = t145
									goto l49
								}
							l48:
							}
							m.fn153(i32(1070576))
							panic("unreachable")
						}
						v8 = i32(16)
						v6 = i32(4)
						v12 = v5<<16 | i32(4)
						v9 = i32(1070616)
						v5 = i32(1070632)
						goto l9
					}
				}
			}
		l5:
			p17 := i32(257)
			if v7 == i64(-0x1ee54e5e1fee3030) {
				p17 = i32(1)
			}
			t18 := int32(load16(m.memory[int64(uint32(v3))+96:]))
			v12 = p17 | t18<<16
			t19 := int32(load16(m.memory[int64(uint32(v3))+102:]))
			v8 = t19
			t20 := int32(load32(m.memory[int64(uint32(v3))+98:]))
			v9 = t20
			goto l9
		}
	}
l3:
	v12 = i32(0)
l9:
	store32(m.memory[int64(uint32(v0))+20:], uint32(v6))
	store32(m.memory[int64(uint32(v0))+16:], uint32(v5))
	store32(m.memory[int64(uint32(v0))+12:], uint32(v8))
	store32(m.memory[int64(uint32(v0))+8:], uint32(v9))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v12))
	store32(m.memory[uint32(v0):], uint32(i32(-1)))
l14:
	m.g0 = v3 + i32(4208)
}
func (m *Module) fn485(v0, v1 int32) {
	var v2 int32
	var v3 int64
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	m.fn59(v2+i32(8), v1, i32(4), i32(4))
	t1 := int64(load64(m.memory[int64(uint32(v2))+8:]))
	v3 = t1
	store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
	store64(m.memory[uint32(v0):], uint64(v3))
	m.g0 = v2 + i32(16)
}
func (m *Module) fn486(v0, v1 int32) {
	var v2, v3, v4, v5, v6, v7, v8 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	t1 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	t2 := v2 + i32(4)
	v3 = t1
	t3 := int32(load32(m.memory[int64(uint32(v1))+16:]))
	t4 := v3
	v4 = t3
	m.fn492(t2, t4, v4)
	{
		t5 := int32(load32(m.memory[int64(uint32(v2))+8:]))
		if t5 != i32(1) {
			m.fn91(i32(1087526), i32(35), i32(1087544))
			panic("unreachable")
		}
		t6 := int32(load32(m.memory[int64(uint32(v2))+12:]))
		m.fn493(v0, t6)
		t7 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		t8 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		v5 = t8
		v6 = t7 + v5<<2
		t9 := int32(load32(m.memory[uint32(v1):]))
		v1 = t9
		var p10 int32
		if v4 == i32(1) {
			p10 = 1
		}
		v7 = p10
		var p11 int32
		if v4 == i32(3) {
			p11 = 1
		}
		v8 = p11
	l6:
		{
			if uint32(v3) < uint32(v4) {
				store32(m.memory[int64(uint32(v0))+8:], uint32(v5))
				m.g0 = v2 + i32(16)
				return
			}
			if v4 == 0 {
				m.fn158(i32(0), i32(0), i32(1073784))
				panic("unreachable")
			}
			if v7 != 0 {
				m.fn158(i32(1), i32(1), i32(1073800))
				panic("unreachable")
			}
			if uint32(v4) <= uint32(i32(2)) {
				m.fn158(i32(2), i32(2), i32(1073816))
				panic("unreachable")
			}
			if v8 != 0 {
				m.fn158(i32(3), i32(3), i32(1073832))
				panic("unreachable")
			}
			v3 = v3 - v4
			t12 := int32(m.memory[uint32(v1+i32(1))])
			t13 := int32(m.memory[uint32(v1)])
			t14 := int32(m.memory[uint32(v1+i32(2))])
			t15 := int32(m.memory[uint32(v1+i32(3))])
			store32(m.memory[uint32(v6):], uint32(t12<<8|t13|t14<<16|t15<<24))
			v6 = v6 + i32(4)
			v5 = v5 + i32(1)
			v1 = v1 + v4
			goto l6
		}
	}
}
func (m *Module) fn487(v0, v1 int32) {
	m.fn449(v1, v0)
}
func (m *Module) fn488(v0, v1, v2 int32) {
	var v3 int32
	if v1 != 0 {
		if v2 == 0 {
			m.fn494(i32(1086504))
			panic("unreachable")
		}
		t0 := int32(uint32(v1) / uint32(v2))
		v3 = t0
		t1 := v3
		var p2 int32
		if v1-v3*v2 != i32(0) {
			p2 = 1
		}
		v1 = t1 + p2
		goto l1
	}
	v1 = i32(0)
	goto l1
l1:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v1))
	store32(m.memory[int64(uint32(v0))+4:], uint32(i32(1)))
	store32(m.memory[uint32(v0):], uint32(v1))
}
func (m *Module) fn489(v0, v1, v2, v3 int32) {
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
func (m *Module) fn490(v0, v1 int32) {
	{
		t0 := int32(load32(m.memory[uint32(v1):]))
		if t0 == i32(-1) {
			goto l0
		}
		t1 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		store32(m.memory[int64(uint32(v0))+8:], uint32(t1))
		t2 := int64(load64(m.memory[uint32(v1):]))
		store64(m.memory[uint32(v0):], uint64(t2))
		return
	}
l0:
	t3 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	t4 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	m.fn884(v0, t3, t4)
}
func (m *Module) fn491(v0, v1, v2 int32) {
	var v3, v4 int32
	t0 := m.g0
	v3 = t0 - i32(16)
	m.g0 = v3
	t1 := v3
	v4 = v2 & i32(3)
	store32(m.memory[int64(uint32(t1))+12:], uint32(v4))
	if v4 != 0 {
		m.fn1586(i32(0), v3+i32(12), i32(1287584), i32(0), v0, i32(1099684))
		panic("unreachable")
	}
	store64(m.memory[int64(uint32(v0))+12:], uint64(i64(0x400000000)))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v2))
	store32(m.memory[uint32(v0):], uint32(v1))
	store32(m.memory[int64(uint32(v0))+8:], uint32(v1+v2))
	m.g0 = v3 + i32(16)
}
