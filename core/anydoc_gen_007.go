package core

func (m *Module) fn267(v0, v1 int32) int32 {
	var v2, v3, v4 int32
	t0 := int32(load32(m.memory[uint32(v0):]))
	v2 = t0
	v0 = i32(1)
	{
		t1 := int32(load32(m.memory[uint32(v1):]))
		v3 = t1
		t2 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t3 := v3
		v1 = t2
		t4 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		v4 = t4
		t5 := m.t0[uint(v4)].(func(int32, int32, int32) int32)(t3, i32(1284186), i32(22))
		if t5 != 0 {
			goto l0
		}
		t6 := m.fn282(v3, v1, v2)
		if t6 != 0 {
			goto l0
		}
		t7 := m.t0[uint(v4)].(func(int32, int32, int32) int32)(v3, i32(1283984), i32(2))
		v0 = t7
	}
l0:
	return v0
}
func (m *Module) fn268(v0, v1 int32) int32 {
	var v2 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	t1 := int32(load32(m.memory[uint32(v0):]))
	store32(m.memory[int64(uint32(v2))+12:], uint32(t1))
	t2 := m.fn283(v1, i32(1087216), i32(9), i32(1087225), i32(7), v2+i32(12), i32(50))
	v0 = t2
	m.g0 = v2 + i32(16)
	return v0
}
func (m *Module) fn269(v0, v1 int32) int32 {
	var v2, v3, v4 int32
	t0 := int32(load32(m.memory[uint32(v0):]))
	v2 = t0
	v0 = i32(1)
	{
		t1 := int32(load32(m.memory[uint32(v1):]))
		v3 = t1
		t2 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t3 := v3
		v1 = t2
		t4 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		v4 = t4
		t5 := m.t0[uint(v4)].(func(int32, int32, int32) int32)(t3, i32(1284208), i32(19))
		if t5 != 0 {
			goto l0
		}
		t6 := m.fn282(v3, v1, v2)
		if t6 != 0 {
			goto l0
		}
		t7 := m.t0[uint(v4)].(func(int32, int32, int32) int32)(v3, i32(1283984), i32(2))
		v0 = t7
	}
l0:
	return v0
}
func (m *Module) fn270(v0, v1 int32) int32 {
	var v2, v3, v4 int32
	t0 := int32(load32(m.memory[uint32(v0):]))
	v2 = t0
	v0 = i32(1)
	{
		t1 := int32(load32(m.memory[uint32(v1):]))
		v3 = t1
		t2 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t3 := v3
		v1 = t2
		t4 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		v4 = t4
		t5 := m.t0[uint(v4)].(func(int32, int32, int32) int32)(t3, i32(1284260), i32(20))
		if t5 != 0 {
			goto l0
		}
		t6 := m.fn282(v3, v1, v2)
		if t6 != 0 {
			goto l0
		}
		t7 := m.t0[uint(v4)].(func(int32, int32, int32) int32)(v3, i32(1283984), i32(2))
		v0 = t7
	}
l0:
	return v0
}
func (m *Module) fn271(v0 int32) {
	var v1 int32
	t0 := m.g0
	v1 = t0 - i32(16)
	m.g0 = v1
	t1 := int32(load32(m.memory[uint32(v0):]))
	m.fn273(v1+i32(8), v0, t1, i32(1), i32(8), i32(24))
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
func (m *Module) fn272(v0 int32) {
	var v1 int32
	t0 := m.g0
	v1 = t0 - i32(16)
	m.g0 = v1
	t1 := int32(load32(m.memory[uint32(v0):]))
	m.fn273(v1+i32(8), v0, t1, i32(1), i32(4), i32(12))
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
func (m *Module) fn273(v0, v1, v2, v3, v4, v5 int32) {
	var v6 int32
	t0 := m.g0
	v6 = t0 - i32(16)
	m.g0 = v6
	{
		v2 = v3 + v2
		if uint32(v2) >= uint32(v3) {
			goto l0
		}
		v3 = i32(0)
		goto l3
	l0:
		t1 := int32(load32(m.memory[uint32(v1):]))
		t2 := v6 + i32(4)
		v3 = t1
		t3 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t4 := v3
		t5 := v2
		v3 = v3 << 1
		p6 := v3
		if uint32(v2) > uint32(v3) {
			p6 = t5
		}
		v3 = p6
		t8 := v3
		p7 := i32(4)
		if v5 == i32(1) {
			p7 = i32(8)
		}
		v2 = p7
		p9 := v2
		if uint32(v3) > uint32(v2) {
			p9 = t8
		}
		v3 = p9
		m.fn1719(t2, t4, t3, v3, v4, v5)
		{
			t10 := int32(load32(m.memory[int64(uint32(v6))+4:]))
			if t10 != i32(1) {
				goto l2
			}
			t11 := int32(load32(m.memory[int64(uint32(v6))+12:]))
			v2 = t11
			t12 := int32(load32(m.memory[int64(uint32(v6))+8:]))
			v3 = t12
			goto l3
		}
	l2:
		t13 := int32(load32(m.memory[int64(uint32(v6))+8:]))
		v2 = t13
		store32(m.memory[uint32(v1):], uint32(v3))
		store32(m.memory[int64(uint32(v1))+4:], uint32(v2))
		v3 = i32(-1)
	}
l3:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v2))
	store32(m.memory[uint32(v0):], uint32(v3))
	m.g0 = v6 + i32(16)
}
func (m *Module) fn274(v0, v1 int32) {
	t0 := int32(load32(m.memory[int64(uint32(v1))+12:]))
	m.fn1077(v0, t0, v1)
}
func (m *Module) fn275(v0, v1 int32) int32 {
	var v2, v3, v4 int32
	v2 = i32(0)
	{
		switch v1 + i32(-2) {
		default:
			goto l3
		case 0:
			v1 = i32(116)
			v3 = i32(1)
			{
				t0 := int32(m.memory[uint32(v0)])
				v4 = t0
				if v4 != i32(108) {
					if v4 != i32(103) {
						goto l3
					}
					v2 = i32(1282677)
					goto l5
				}
				v2 = i32(1282676)
				goto l5
			}
		case 1:
			t1 := int32(m.memory[uint32(v0)])
			if t1 != i32(97) {
				goto l3
			}
			t2 := int32(m.memory[int64(uint32(v0))+1])
			if t2 != i32(109) {
				goto l3
			}
			v2 = i32(1282678)
			v1 = i32(112)
			v3 = i32(2)
			goto l5
		case 2:
			{
				{
					t3 := int32(m.memory[uint32(v0)])
					v1 = t3
					if v1 == i32(113) {
						goto l6
					}
					if v1 != i32(97) {
						goto l3
					}
					t4 := int32(m.memory[int64(uint32(v0))+1])
					if t4 != i32(112) {
						goto l3
					}
					t5 := int32(m.memory[int64(uint32(v0))+2])
					if t5 != i32(111) {
						goto l3
					}
					v2 = i32(1282679)
					v1 = i32(115)
					goto l7
				}
			l6:
				t6 := int32(m.memory[int64(uint32(v0))+1])
				if t6 != i32(117) {
					goto l3
				}
				t7 := int32(m.memory[int64(uint32(v0))+2])
				if t7 != i32(111) {
					goto l3
				}
				v2 = i32(1282665)
				v1 = i32(116)
			}
		l7:
			v3 = i32(3)
		}
	l5:
		t8 := int32(m.memory[uint32(v0+v3)])
		p9 := i32(0)
		if t8 == v1 {
			p9 = v2
		}
		v2 = p9
	}
l3:
	return v2
}
func (m *Module) fn276(v0, v1, v2 int32) {
	var v3, v4 int32
	var v5 int64
	t0 := m.g0
	v3 = t0 - i32(16)
	m.g0 = v3
	m.fn13(v3, v1, v2, i32(120))
	{
		{
			t1 := int32(load32(m.memory[uint32(v3):]))
			v4 = t1
			if v4 == 0 {
				m.fn1753(v3+i32(8), v1, v2, i32(10))
				t5 := int32(m.memory[int64(uint32(v3))+8])
				if t5 == i32(255) {
					goto l1
				}
				t6 := int64(load64(m.memory[int64(uint32(v3))+8:]))
				v5 = t6
				if v5&i64(255) != i64(255) {
					store64(m.memory[uint32(v0):], uint64(v5))
					goto l5
				}
				v2 = int32(int64(uint64(v5) >> 32))
				goto l3
			}
			t2 := int32(load32(m.memory[int64(uint32(v3))+4:]))
			m.fn1753(v3+i32(8), v4, t2, i32(16))
			t3 := int32(m.memory[int64(uint32(v3))+8])
			if t3 == i32(255) {
				goto l1
			}
			t4 := int64(load64(m.memory[int64(uint32(v3))+8:]))
			v5 = t4
			if v5&i64(255) != i64(255) {
				store64(m.memory[uint32(v0):], uint64(v5))
				goto l5
			}
			v2 = int32(int64(uint64(v5) >> 32))
			goto l3
		}
	l1:
		t7 := int32(load32(m.memory[int64(uint32(v3))+12:]))
		v2 = t7
	}
l3:
	if v2 != 0 {
		goto l6
	}
	store32(m.memory[int64(uint32(v0))+4:], uint32(i32(0)))
	m.memory[uint32(v0)] = byte(i32(3))
	goto l5
l6:
	if uint32(v2^i32(55296)+i32(-1114112)) < uint32(i32(-1112064)) {
		goto l7
	}
	m.memory[uint32(v0)] = byte(i32(255))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v2))
	goto l5
l7:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v2))
	m.memory[uint32(v0)] = byte(i32(2))
l5:
	m.g0 = v3 + i32(16)
}
func (m *Module) fn277(v0, v1 int32) {
	if v0 == i32(-1) {
		return
	}
	m.fn1091(v0, v1)
}
func (m *Module) fn278(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t1 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	t2 := m.fn110(v1, t0, t1)
	return t2
}
func (m *Module) fn279(v0, v1 int32) {
	var v2, v3, v4 int32
	if uint32(v0) < uint32(i32(128)) {
		goto l0
	}
	v2 = v0&i32(63) | i32(-128)
	v3 = int32(uint32(v0) >> 6)
	if uint32(v0) >= uint32(i32(2048)) {
		v4 = int32(uint32(v0) >> 12)
		v3 = v3&i32(63) | i32(-128)
		if uint32(v0) > uint32(i32(0xffff)) {
			m.memory[int64(uint32(v1))+3] = byte(v2)
			m.memory[int64(uint32(v1))+2] = byte(v3)
			m.memory[int64(uint32(v1))+1] = byte(v4&i32(63) | i32(-128))
			m.memory[uint32(v1)] = byte(int32(uint32(v0)>>18) | i32(-16))
			return
		}
		m.memory[int64(uint32(v1))+2] = byte(v2)
		m.memory[int64(uint32(v1))+1] = byte(v3)
		m.memory[uint32(v1)] = byte(v4 | i32(224))
		return
	}
	m.memory[int64(uint32(v1))+1] = byte(v2)
	m.memory[uint32(v1)] = byte(v3 | i32(192))
	return
l0:
	m.memory[uint32(v1)] = byte(v0)
}
func (m *Module) fn280(v0, v1 int32) {
	var v2, v3, v4 int32
	v2 = i32(0)
	{
		{
			t0 := int32(load32(m.memory[uint32(v1):]))
			v3 = t0
			t1 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			if v3 != t1 {
				goto l0
			}
			goto l1
		}
	l0:
		t2 := int32(m.memory[uint32(v3)])
		v4 = t2 + i32(-48)
		if uint32(v4&i32(255)) > uint32(i32(9)) {
			goto l1
		}
		v2 = i32(1)
		store32(m.memory[uint32(v1):], uint32(v3+i32(1)))
	}
l1:
	m.memory[int64(uint32(v0))+1] = byte(v4)
	m.memory[uint32(v0)] = byte(v2)
}
func (m *Module) fn281(v0 int32) {
	var v1, v2 int32
	t0 := int32(load32(m.memory[int64(uint32(v0))+768:]))
	v1 = t0
	v2 = v1 + i32(-1)
	var p1 int32
	if uint32(v1) < uint32(i32(769)) {
		p1 = 1
	}
	v1 = p1
l3:
	{
		if v2 == i32(-1) {
			return
		}
		if v1 == 0 {
			m.fn158(v2, i32(768), i32(1087976))
			panic("unreachable")
		}
		t2 := int32(m.memory[uint32(v0+v2)])
		if t2 == 0 {
			store32(m.memory[int64(uint32(v0))+768:], uint32(v2))
			v2 = v2 + i32(-1)
			goto l3
		}
	}
}
func (m *Module) fn282(v0, v1, v2 int32) int32 {
	var v3, v4 int32
	t0 := int32(load32(m.memory[int64(uint32(v1))+12:]))
	v3 = t0
	{
		{
			{
				t1 := int32(load32(m.memory[uint32(v2):]))
				if t1 == i32(-1) {
					goto l0
				}
				v4 = i32(1)
				t2 := m.t0[uint(v3)].(func(int32, int32, int32) int32)(v0, i32(1282658), i32(6))
				if t2 != 0 {
					goto l1
				}
				t3 := int32(load32(m.memory[int64(uint32(v2))+4:]))
				t4 := int32(load32(m.memory[int64(uint32(v2))+8:]))
				t5 := m.fn1780(v0, v1, t3, t4)
				if t5 == 0 {
					goto l2
				}
				goto l1
			}
		l0:
			v4 = i32(1)
			t6 := m.t0[uint(v3)].(func(int32, int32, int32) int32)(v0, i32(1282649), i32(9))
			if t6 != 0 {
				goto l1
			}
			t7 := int32(load32(m.memory[int64(uint32(v2))+4:]))
			t8 := int32(load32(m.memory[int64(uint32(v2))+8:]))
			t9 := m.fn1780(v0, v1, t7, t8)
			if t9 != 0 {
				goto l1
			}
		}
	l2:
		t10 := m.t0[uint(v3)].(func(int32, int32, int32) int32)(v0, i32(1282664), i32(1))
		v4 = t10
	}
l1:
	return v4
}
func (m *Module) fn283(v0, v1, v2, v3, v4, v5, v6 int32) int32 {
	var v7 int32
	t0 := m.g0
	v7 = t0 - i32(16)
	m.g0 = v7
	t1 := int32(load32(m.memory[uint32(v0):]))
	t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t3 := int32(load32(m.memory[int64(uint32(t2))+12:]))
	t4 := m.t0[uint(t3)].(func(int32, int32, int32) int32)(t1, v1, v2)
	v2 = t4
	m.memory[int64(uint32(v7))+13] = byte(i32(0))
	m.memory[int64(uint32(v7))+12] = byte(v2)
	store32(m.memory[int64(uint32(v7))+8:], uint32(v0))
	t5 := m.fn1647(v7+i32(8), v3, v4, v5, v6)
	v6 = t5
	t6 := int32(m.memory[int64(uint32(v7))+13])
	v2 = t6
	t7 := int32(m.memory[int64(uint32(v7))+12])
	t8 := v2
	v1 = t7
	v0 = t8 | v1
	{
		if v2 != i32(1) {
			goto l0
		}
		if v1&i32(1) != 0 {
			goto l0
		}
		{
			t9 := int32(load32(m.memory[uint32(v6):]))
			v0 = t9
			t10 := int32(m.memory[int64(uint32(v0))+10])
			if t10&i32(128) != 0 {
				goto l1
			}
			t11 := int32(load32(m.memory[uint32(v0):]))
			t12 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t13 := int32(load32(m.memory[int64(uint32(t12))+12:]))
			t14 := m.t0[uint(t13)].(func(int32, int32, int32) int32)(t11, i32(1283984), i32(2))
			v0 = t14
			goto l0
		}
	l1:
		t15 := int32(load32(m.memory[uint32(v0):]))
		t16 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		t17 := int32(load32(m.memory[int64(uint32(t16))+12:]))
		t18 := m.t0[uint(t17)].(func(int32, int32, int32) int32)(t15, i32(1108167), i32(1))
		v0 = t18
	}
l0:
	m.g0 = v7 + i32(16)
	return v0 & i32(1)
}
func (m *Module) fn284(v0, v1, v2, v3 int32) int32 {
	t0 := m.fn100(v0, v1, v2, v3)
	return t0
}
func (m *Module) fn285(v0, v1, v2 int32) {
	var v3, v4 int32
	var v5 int64
	var v6 int32
	t0 := int32(load32(m.memory[int64(uint32(v0))+56:]))
	store32(m.memory[int64(uint32(v0))+56:], uint32(t0+v2))
	{
		{
			t1 := int32(load32(m.memory[int64(uint32(v0))+60:]))
			v3 = t1
			if v3 != 0 {
				t2 := v1
				t3 := v2
				v4 = i32(8) - v3
				p4 := v2
				if uint32(v4) < uint32(v2) {
					p4 = v4
				}
				t5 := m.fn287(t2, t3, i32(0), p4)
				v5 = t5
				t6 := int64(load64(m.memory[int64(uint32(v0))+48:]))
				t7 := v0
				v5 = t6 | i64_shl(v5, int64(uint32(v3<<3)))
				store64(m.memory[int64(uint32(t7))+48:], uint64(v5))
				{
					if uint32(v2) < uint32(v4) {
						v3 = v3 + v2
						goto l3
					}
					t8 := int64(load64(m.memory[int64(uint32(v0))+24:]))
					store64(m.memory[int64(uint32(v0))+24:], uint64(t8^v5))
					m.fn286(v0)
					store32(m.memory[int64(uint32(v0))+60:], uint32(i32(0)))
					t9 := int64(load64(m.memory[uint32(v0):]))
					t10 := int64(load64(m.memory[int64(uint32(v0))+48:]))
					store64(m.memory[uint32(v0):], uint64(t9^t10))
					goto l1
				}
			}
			v4 = i32(0)
			goto l1
		}
	l1:
		v6 = v2 - v4
		v3 = v6 & i32(-8)
	l5:
		{
			if uint32(v4) >= uint32(v3) {
				goto l4
			}
			t11 := int64(load64(m.memory[int64(uint32(v0))+24:]))
			t12 := int64(load64(m.memory[uint32(v1+v4):]))
			t13 := v0
			v5 = t12
			store64(m.memory[int64(uint32(t13))+24:], uint64(t11^v5))
			m.fn286(v0)
			t14 := int64(load64(m.memory[uint32(v0):]))
			store64(m.memory[uint32(v0):], uint64(v5^t14))
			v4 = v4 + i32(8)
			goto l5
		}
	l4:
		t15 := v0
		t16 := v1
		t17 := v2
		t18 := v4
		v3 = v6 & i32(7)
		t19 := m.fn287(t16, t17, t18, v3)
		store64(m.memory[int64(uint32(t15))+48:], uint64(t19))
	}
l3:
	store32(m.memory[int64(uint32(v0))+60:], uint32(v3))
}
func (m *Module) fn286(v0 int32) {
	var v1, v2, v3, v4, v5 int64
	t0 := int64(load64(m.memory[int64(uint32(v0))+24:]))
	t1 := v0
	v1 = t0
	t2 := int64(load64(m.memory[int64(uint32(v0))+8:]))
	t3 := i64_rotl(v1, i64(16))
	v1 = v1 + t2
	v2 = t3 ^ v1
	t4 := int64(load64(m.memory[int64(uint32(v0))+16:]))
	t5 := v2
	v3 = t4
	t6 := int64(load64(m.memory[uint32(v0):]))
	v4 = v3 + t6
	v5 = t5 + i64_rotl(v4, i64(32))
	store64(m.memory[uint32(t1):], uint64(v5))
	store64(m.memory[int64(uint32(v0))+24:], uint64(i64_rotl(v2, i64(21))^v5))
	t7 := v0
	t8 := v1
	v2 = i64_rotl(v3, i64(13)) ^ v4
	v1 = t8 + v2
	store64(m.memory[int64(uint32(t7))+16:], uint64(v1^i64_rotl(v2, i64(17))))
	store64(m.memory[int64(uint32(v0))+8:], uint64(i64_rotl(v1, i64(32))))
}
func (m *Module) fn287(v0, v1, v2, v3 int32) int64 {
	var v4 int32
	var v5 int64
	v4 = i32(4)
	{
		if uint32(v3) >= uint32(i32(4)) {
			goto l0
		}
		v5 = i64(0)
		v4 = i32(0)
		goto l1
	l0:
		t0 := int64(load32(m.memory[uint32(v0+v2):]))
		v5 = t0
	}
l1:
	{
		if uint32(v4|i32(1)) >= uint32(v3) {
			goto l2
		}
		t1 := int64(load16(m.memory[uint32(v0+v2+v4):]))
		v5 = i64_shl(t1, int64(uint32(v4<<3))) | v5
		v4 = v4 | i32(2)
	}
l2:
	{
		if uint32(v4) >= uint32(v3) {
			goto l3
		}
		t2 := int64(m.memory[uint32(v0+(v4+v2))])
		v5 = i64_shl(t2, int64(uint32(v4<<3))) | v5
	}
l3:
	return v5
}
func (m *Module) fn288(v0, v1 int32) {
	var v2 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		v2 = t0
		t1 := int32(load32(m.memory[uint32(v0):]))
		if v2 != t1 {
			goto l0
		}
		m.fn289(v0)
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
func (m *Module) fn289(v0 int32) {
	var v1 int32
	t0 := m.g0
	v1 = t0 - i32(16)
	m.g0 = v1
	t1 := int32(load32(m.memory[uint32(v0):]))
	m.fn273(v1+i32(8), v0, t1, i32(1), i32(4), i32(24))
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
func (m *Module) fn290(v0, v1, v2 int32) {
	var v3, v4, v5, v6 int32
	t0 := m.g0
	v3 = t0 - i32(32)
	m.g0 = v3
	t1 := int32(load32(m.memory[uint32(v1):]))
	v4 = t1
	{
		t2 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		v5 = t2
		t3 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		t4 := v5
		v6 = t3
		if uint32(t4) < uint32(v6) {
			goto l0
		}
		v5 = i32(0)
		store32(m.memory[int64(uint32(v3))+16:], uint32(i32(0)))
		store32(m.memory[int64(uint32(v3))+8:], uint32(v4))
		t5 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		store32(m.memory[int64(uint32(v3))+12:], uint32(t5))
		t6 := int32(m.memory[int64(uint32(v1))+16])
		m.memory[int64(uint32(v3))+20] = byte(t6)
		m.fn295(v3+i32(24), v2, v3+i32(8))
		store32(m.memory[int64(uint32(v1))+8:], uint32(i32(0)))
		t7 := int32(load32(m.memory[int64(uint32(v3))+16:]))
		t8 := v1
		v6 = t7
		store32(m.memory[int64(uint32(t8))+12:], uint32(v6))
		t9 := int32(m.memory[int64(uint32(v3))+20])
		m.memory[int64(uint32(v1))+16] = byte(t9)
		t10 := int32(m.memory[int64(uint32(v3))+24])
		if t10 == i32(255) {
			goto l0
		}
		t11 := int64(load64(m.memory[int64(uint32(v3))+24:]))
		store64(m.memory[int64(uint32(v0))+4:], uint64(t11))
		store32(m.memory[uint32(v0):], uint32(i32(1)))
		goto l1
	}
l0:
	store32(m.memory[uint32(v0):], uint32(i32(0)))
	store32(m.memory[int64(uint32(v0))+8:], uint32(v6-v5))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v4+v5))
l1:
	m.g0 = v3 + i32(32)
}
func (m *Module) fn291(v0, v1 int32) {
	var v2 int32
	{
		t0 := int32(load32(m.memory[uint32(v1):]))
		v2 = t0
		if v2 == 0 {
			goto l0
		}
		m.t0[uint(v2)].(func(int32))(v0)
	}
l0:
	{
		t1 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		v2 = t1
		if v2 == 0 {
			return
		}
		t2 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		m.fn40(v0, t2, v2)
	}
}
func (m *Module) fn292(v0 int32) {
	m.fn10(v0, i32(12), i32(4))
}
func (m *Module) fn293(v0 int64) int32 {
	var v1 int32
	t0 := m.fn113(i32(4), i32(16))
	v1 = t0
	store64(m.memory[int64(uint32(v1))+8:], uint64(v0))
	store64(m.memory[uint32(v1):], uint64(i64(0x100000001)))
	return v1
}
func (m *Module) fn294(v0, v1, v2, v3 int32) {
	var v4, v5, v6, v7 int32
	v4 = i32(0)
	if uint32(v2) < uint32(v3) {
		goto l0
	}
	goto l1
l0:
	v5 = v3 - v2
	if uint32(v5) > uint32(i32(3)) {
		t1 := int32(load32(m.memory[uint32(v2):]))
		v6 = v1 & i32(255) * i32(16843009)
		v7 = t1 ^ v6
		if (i32(16843008)-v7|v7)&i32(-2139062144) == i32(-2139062144) {
			v2 = v2&i32(-4) + i32(4)
			if uint32(v5) > uint32(i32(8)) {
				v7 = v3 + i32(-8)
			l10:
				{
					if uint32(v2) > uint32(v7) {
						goto l9
					}
					t4 := int32(load32(m.memory[uint32(v2):]))
					v5 = t4 ^ v6
					if (i32(16843008)-v5|v5)&i32(-2139062144) != i32(-2139062144) {
						goto l9
					}
					t5 := int32(load32(m.memory[uint32(v2+i32(4)):]))
					v5 = t5 ^ v6
					if (i32(16843008)-v5|v5)&i32(-2139062144) != i32(-2139062144) {
						goto l9
					}
					v2 = v2 + i32(8)
					goto l10
				}
			l9:
				v1 = v1 & i32(255)
			l11:
				{
					if uint32(v2) >= uint32(v3) {
						goto l1
					}
					t6 := int32(m.memory[uint32(v2)])
					if v1 == t6 {
						goto l3
					}
					v2 = v2 + i32(1)
					goto l11
				}
			}
			v1 = v1 & i32(255)
		l8:
			{
				if uint32(v2) >= uint32(v3) {
					goto l1
				}
				t3 := int32(m.memory[uint32(v2)])
				if v1 == t3 {
					goto l3
				}
				v2 = v2 + i32(1)
				goto l8
			}
		}
		v1 = v1 & i32(255)
	l6:
		{
			if uint32(v2) >= uint32(v3) {
				goto l1
			}
			t2 := int32(m.memory[uint32(v2)])
			if v1 == t2 {
				goto l3
			}
			v2 = v2 + i32(1)
			goto l6
		}
	}
	v1 = v1 & i32(255)
l4:
	{
		if uint32(v2) >= uint32(v3) {
			goto l1
		}
		t0 := int32(m.memory[uint32(v2)])
		if v1 == t0 {
			goto l3
		}
		v2 = v2 + i32(1)
		goto l4
	}
l3:
	v4 = i32(1)
l1:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v2))
	store32(m.memory[uint32(v0):], uint32(v4))
}
func (m *Module) fn295(v0, v1, v2 int32) {
	var v3, v4 int32
	t0 := m.g0
	v3 = t0 - i32(16)
	m.g0 = v3
	m.fn296(v3, v2)
	t1 := int32(load32(m.memory[uint32(v3):]))
	t2 := int32(load32(m.memory[int64(uint32(v3))+4:]))
	m.fn297(v3+i32(8), v1+i32(176), t1, t2)
	{
		{
			t3 := int32(m.memory[int64(uint32(v3))+8])
			if t3 == i32(255) {
				goto l0
			}
			t4 := int64(load64(m.memory[int64(uint32(v3))+8:]))
			store64(m.memory[uint32(v0):], uint64(t4))
			goto l1
		}
	l0:
		t5 := int32(load32(m.memory[int64(uint32(v3))+12:]))
		v1 = t5
		t6 := int32(load32(m.memory[int64(uint32(v2))+4:]))
		t7 := int32(load32(m.memory[int64(uint32(v2))+8:]))
		t8 := v1
		v4 = t7
		t9 := int32(m.memory[int64(uint32(v2))+12])
		p10 := i32(0)
		if t9 != 0 {
			p10 = t6 - v4
		}
		if uint32(t8) > uint32(p10) {
			m.fn256(i32(1072353), i32(36), i32(1072392))
			panic("unreachable")
		}
		m.memory[uint32(v0)] = byte(i32(255))
		store32(m.memory[int64(uint32(v2))+8:], uint32(v4+v1))
	}
l1:
	m.g0 = v3 + i32(16)
}
func (m *Module) fn296(v0, v1 int32) {
	var v2, v3 int32
	t0 := int32(load32(m.memory[uint32(v1):]))
	t1 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	v2 = t1
	v3 = t0 + v2
	t2 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	v2 = t2 - v2
	{
		t3 := int32(m.memory[int64(uint32(v1))+12])
		if t3 == i32(1) {
			goto l0
		}
		if v2 == 0 {
			goto l1
		}
		memory_zero(m.memory, uint32(v3), uint32(v2))
	l1:
		m.memory[int64(uint32(v1))+12] = byte(i32(1))
	}
l0:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v2))
	store32(m.memory[uint32(v0):], uint32(v3))
}
func (m *Module) fn297(v0, v1, v2, v3 int32) {
	t0 := int32(load32(m.memory[uint32(v1):]))
	switch t0 {
	default:
		m.fn125(v0, i32(1079732), i32(37))
		return
	case 1:
		m.fn298(v0, v1+i32(8), v2, v3)
		return
	case 2:
		t1 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		m.fn299(v0, t1, v2, v3)
		return
	case 3:
		t2 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		m.fn300(v0, t2, v2, v3)
	}
}
func (m *Module) fn298(v0, v1, v2, v3 int32) {
	var v4 int32
	var v5, v6 int64
	t0 := m.g0
	v4 = t0 - i32(16)
	m.g0 = v4
	{
		{
			t1 := int64(load64(m.memory[int64(uint32(v1))+8:]))
			v5 = t1
			if v5 != i64(0) {
				goto l0
			}
			m.memory[uint32(v0)] = byte(i32(255))
			store32(m.memory[int64(uint32(v0))+4:], uint32(i32(0)))
			goto l1
		}
	l0:
		t2 := int32(load32(m.memory[int64(uint32(v1))+16:]))
		t3 := v4 + i32(8)
		t4 := v2
		t5 := v5
		v6 = int64(uint32(v3))
		p6 := v6
		if uint64(v5) < uint64(v6) {
			p6 = t5
		}
		m.fn301(t3, t2, t4, int32(p6))
		{
			t7 := int32(m.memory[int64(uint32(v4))+8])
			if t7 == i32(255) {
				goto l2
			}
			t8 := int64(load64(m.memory[int64(uint32(v4))+8:]))
			store64(m.memory[uint32(v0):], uint64(t8))
			goto l1
		}
	l2:
		{
			t9 := int32(load32(m.memory[int64(uint32(v4))+12:]))
			t10 := v5
			v3 = t9
			v6 = int64(uint32(v3))
			if uint64(t10) >= uint64(v6) {
				goto l3
			}
			m.fn91(i32(1087384), i32(69), i32(1087420))
			panic("unreachable")
		}
	l3:
		store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
		m.memory[uint32(v0)] = byte(i32(255))
		store64(m.memory[int64(uint32(v1))+8:], uint64(v5-v6))
	}
l1:
	m.g0 = v4 + i32(16)
}
func (m *Module) fn299(v0, v1, v2, v3 int32) {
	var v4, v5 int32
	var v6 int64
	t0 := m.g0
	v4 = t0 - i32(16)
	m.g0 = v4
	m.fn302(v4+i32(8), v1, v2, v3)
	{
		{
			{
				t1 := int32(m.memory[int64(uint32(v4))+8])
				if t1 != i32(255) {
					goto l0
				}
				t2 := int32(load32(m.memory[int64(uint32(v4))+12:]))
				v5 = t2
				goto l1
			}
		l0:
			t3 := int64(load64(m.memory[int64(uint32(v4))+8:]))
			v6 = t3
			if v6&i64(255) != i64(255) {
				store64(m.memory[uint32(v0):], uint64(v6))
				goto l5
			}
			v5 = int32(int64(uint64(v6) >> 32))
		}
	l1:
		t4 := int32(m.memory[int64(uint32(v1))+68])
		if t4 == 0 {
			goto l3
		}
		if v3 == 0 {
			goto l4
		}
		if v5 != 0 {
			goto l4
		}
		t5 := int32(load32(m.memory[int64(uint32(v1))+64:]))
		t6 := int32(load32(m.memory[int64(uint32(v1))+56:]))
		if t5 == t6 {
			goto l4
		}
		m.fn303(v0, i32(21), i32(1087288), i32(16))
		goto l5
	}
l4:
	if uint32(v5) > uint32(v3) {
		m.fn151(i32(0), v5, v3, i32(1087304))
		panic("unreachable")
	}
	m.fn304(v1+i32(48), v2, v5)
	goto l3
l3:
	m.memory[uint32(v0)] = byte(i32(255))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
l5:
	m.g0 = v4 + i32(16)
}
func (m *Module) fn300(v0, v1, v2, v3 int32) {
	var v4, v5, v6 int32
	var v7, v8 int64
	var v9, v10, v11, v12, v13 int32
	t0 := m.g0
	v4 = t0 - i32(32)
	m.g0 = v4
	{
		{
			{
				t1 := int64(load64(m.memory[int64(uint32(v1))+24:]))
				if t1 == i64(2) {
					v9 = v1 + i32(56)
					{
						t23 := int32(load32(m.memory[int64(uint32(v1))+40:]))
						t24 := int32(load32(m.memory[int64(uint32(v1))+44:]))
						if t23 != t24 {
							goto l7
						}
						t25 := int32(load32(m.memory[int64(uint32(v1))+36:]))
						if uint32(v3) >= uint32(t25) {
							store64(m.memory[int64(uint32(v1))+40:], uint64(i64(0)))
							m.fn302(v4, v9, v2, v3)
							goto l2
						}
					}
				l7:
					m.fn305(v4+i32(16), v1+i32(32), v9)
					{
						t26 := int32(load32(m.memory[int64(uint32(v4))+16:]))
						if t26 == 0 {
							t28 := int64(load64(m.memory[int64(uint32(v4))+20:]))
							store64(m.memory[int64(uint32(v4))+8:], uint64(t28))
							m.fn307(v4+i32(16), v4+i32(8), v2, v3)
							t29 := int32(m.memory[int64(uint32(v4))+16])
							if t29 == i32(255) {
								t31 := int32(load32(m.memory[int64(uint32(v1))+44:]))
								t32 := v1
								v9 = t31
								t33 := int32(load32(m.memory[int64(uint32(v1))+40:]))
								t34 := int32(load32(m.memory[int64(uint32(v4))+20:]))
								t35 := v9
								v11 = t34
								v12 = t33 + v11
								p36 := v12
								if uint32(v9) < uint32(v12) {
									p36 = t35
								}
								store32(m.memory[int64(uint32(t32))+40:], uint32(p36))
								goto l5
							}
							t30 := int64(load64(m.memory[int64(uint32(v4))+16:]))
							store64(m.memory[uint32(v4):], uint64(t30))
							goto l2
						}
						t27 := int64(load64(m.memory[int64(uint32(v4))+20:]))
						store64(m.memory[uint32(v4):], uint64(t27))
						goto l2
					}
				}
				v5 = v1 + i32(24)
				v6 = v1 + i32(72)
			l6:
				m.fn305(v4+i32(16), v1, v5)
				{
					t2 := int32(load32(m.memory[int64(uint32(v4))+16:]))
					if t2 != i32(1) {
						t4 := int64(load64(m.memory[int64(uint32(v1))+80:]))
						v7 = t4
						t5 := int64(load64(m.memory[int64(uint32(v1))+72:]))
						v8 = t5
						t6 := int32(load32(m.memory[int64(uint32(v4))+20:]))
						t7 := int32(load32(m.memory[int64(uint32(v4))+24:]))
						t8 := v4 + i32(16)
						t9 := v6
						v9 = t7
						t11 := v9
						t12 := v2
						t13 := v3
						p10 := i32(4)
						if v9 != 0 {
							p10 = i32(0)
						}
						m.fn306(t8, t9, t6, t11, t12, t13, p10)
						t14 := int32(m.memory[int64(uint32(v4))+20])
						v10 = t14
						t15 := int32(load32(m.memory[int64(uint32(v4))+16:]))
						v11 = t15
						t16 := int32(load32(m.memory[int64(uint32(v1))+12:]))
						t17 := v1
						v12 = t16
						t18 := int32(load32(m.memory[int64(uint32(v1))+8:]))
						t19 := int64(load64(m.memory[int64(uint32(v1))+72:]))
						t20 := v12
						v13 = t18 + int32(t19-v8)
						p21 := v13
						if uint32(v12) < uint32(v13) {
							p21 = t20
						}
						store32(m.memory[int64(uint32(t17))+8:], uint32(p21))
						if v11 == i32(2) {
							t22 := int64(load64(m.memory[int64(uint32(v1))+80:]))
							v11 = int32(t22 - v7)
							switch v10 {
							case 2:
								goto l5
							default:
								if v9 == 0 {
									goto l5
								}
								if v3 == 0 {
									goto l5
								}
								if v11 == 0 {
									goto l6
								}
								goto l5
							}
						}
						m.fn303(v4, i32(20), i32(1071542), i32(22))
						goto l2
					}
					t3 := int64(load64(m.memory[int64(uint32(v4))+20:]))
					store64(m.memory[uint32(v4):], uint64(t3))
					goto l2
				}
			}
		l2:
			{
				t37 := int32(m.memory[uint32(v4)])
				if t37 != i32(255) {
					goto l11
				}
				t38 := int32(load32(m.memory[int64(uint32(v4))+4:]))
				v11 = t38
				goto l5
			}
		l11:
			t39 := int64(load64(m.memory[uint32(v4):]))
			v8 = t39
			if v8&i64(255) != i64(255) {
				store64(m.memory[uint32(v0):], uint64(v8))
				goto l15
			}
			v11 = int32(int64(uint64(v8) >> 32))
		}
	l5:
		t40 := int32(m.memory[int64(uint32(v1))+180])
		if t40 == 0 {
			goto l13
		}
		if v3 == 0 {
			goto l14
		}
		if v11 != 0 {
			goto l14
		}
		t41 := int32(load32(m.memory[int64(uint32(v1))+176:]))
		t42 := int32(load32(m.memory[int64(uint32(v1))+168:]))
		if t41 == t42 {
			goto l14
		}
		m.fn303(v0, i32(21), i32(1087288), i32(16))
		goto l15
	}
l14:
	if uint32(v11) > uint32(v3) {
		m.fn151(i32(0), v11, v3, i32(1087304))
		panic("unreachable")
	}
	m.fn304(v1+i32(160), v2, v11)
	goto l13
l13:
	m.memory[uint32(v0)] = byte(i32(255))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v11))
l15:
	m.g0 = v4 + i32(32)
}
func (m *Module) fn301(v0, v1, v2, v3 int32) {
	var v4 int32
	t0 := m.g0
	v4 = t0 - i32(32)
	m.g0 = v4
	m.fn311(v4+i32(16), v1)
	m.fn307(v4+i32(8), v4+i32(24), v2, v3)
	{
		{
			t1 := int32(m.memory[int64(uint32(v4))+8])
			if t1 == i32(255) {
				goto l0
			}
			t2 := int64(load64(m.memory[int64(uint32(v4))+8:]))
			store64(m.memory[uint32(v0):], uint64(t2))
			goto l1
		}
	l0:
		t3 := int32(load32(m.memory[int64(uint32(v4))+12:]))
		t4 := v0
		v3 = t3
		store32(m.memory[int64(uint32(t4))+4:], uint32(v3))
		m.memory[uint32(v0)] = byte(i32(255))
		t5 := int64(load64(m.memory[int64(uint32(v1))+8:]))
		store64(m.memory[int64(uint32(v1))+8:], uint64(t5+int64(uint32(v3))))
	}
l1:
	m.g0 = v4 + i32(32)
}
func (m *Module) fn302(v0, v1, v2, v3 int32) {
	var v4, v5, v6 int32
	var v7 int64
	var v8 int32
	t0 := m.g0
	v4 = t0 - i32(16)
	m.g0 = v4
	v5 = v1 + i32(8)
	{
		t1 := int64(load64(m.memory[uint32(v1):]))
		if t1 != i64(1) {
			m.fn298(v0, v5, v2, v3)
			goto l6
		}
		m.fn298(v4+i32(8), v5, v2, v3)
		{
			{
				t2 := int32(m.memory[int64(uint32(v4))+8])
				if t2 != i32(255) {
					goto l1
				}
				t3 := int32(load32(m.memory[int64(uint32(v4))+12:]))
				v6 = t3
				goto l2
			}
		l1:
			t4 := int64(load64(m.memory[int64(uint32(v4))+8:]))
			v7 = t4
			if v7&i64(255) != i64(255) {
				goto l3
			}
			v6 = int32(int64(uint64(v7) >> 32))
		}
	l2:
		v8 = v1 + i32(32)
		v1 = i32(0)
	l5:
		{
			if v6 == v1 {
				goto l4
			}
			if v3 == v1 {
				goto l4
			}
			v5 = v2 + v1
			t5 := int32(m.memory[uint32(v5)])
			t6 := m.fn308(v8, t5)
			m.memory[uint32(v5)] = byte(t6)
			v1 = v1 + i32(1)
			goto l5
		}
	l4:
		m.memory[uint32(v0)] = byte(i32(255))
		store32(m.memory[int64(uint32(v0))+4:], uint32(v6))
		goto l6
	}
l3:
	store64(m.memory[uint32(v0):], uint64(v7))
l6:
	m.g0 = v4 + i32(16)
}
func (m *Module) fn303(v0, v1, v2, v3 int32) {
	var v4, v5, v6 int32
	t0 := m.g0
	v4 = t0 - i32(32)
	m.g0 = v4
	m.fn1(v4+i32(20), v3, i32(0), i32(1), i32(1))
	t1 := int32(load32(m.memory[int64(uint32(v4))+24:]))
	v5 = t1
	{
		t2 := int32(load32(m.memory[int64(uint32(v4))+20:]))
		if t2 == i32(1) {
			t5 := int32(load32(m.memory[int64(uint32(v4))+28:]))
			m.fn2(v5, t5)
			panic("unreachable")
		}
		t3 := int32(load32(m.memory[int64(uint32(v4))+28:]))
		v6 = t3
		if v3 == 0 {
			goto l1
		}
		memory_copy(m.memory, uint32(v6), uint32(v2), uint32(v3))
	l1:
		m.fn1680(v4+i32(8), i32(4), i32(12), i32(0))
		t4 := int32(load32(m.memory[int64(uint32(v4))+8:]))
		v2 = t4
		if v2 == 0 {
			m.fn85(i32(4), i32(12))
			panic("unreachable")
		}
		store32(m.memory[int64(uint32(v2))+8:], uint32(v3))
		store32(m.memory[int64(uint32(v2))+4:], uint32(v6))
		store32(m.memory[uint32(v2):], uint32(v5))
		m.fn343(v0, v1, v2, i32(1280600))
		m.g0 = v4 + i32(32)
		return
	}
}
func (m *Module) fn304(v0, v1, v2 int32) {
	var v3, v4, v5, v6, v7, v8, v9 int32
	t0 := int64(load64(m.memory[uint32(v0):]))
	store64(m.memory[uint32(v0):], uint64(t0+int64(uint32(v2))))
	v3 = v2 & i32(63)
	v4 = v1 + v2&i32(0x7fffffc0)
	t1 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	v5 = t1 ^ i32(-1)
l2:
	if uint32(v2) < uint32(i32(64)) {
		goto l6
	}
	v6 = int32(uint32(v2) >> 4)
	v7 = v1 + i32(64)
	v2 = v2 + i32(-64)
	v8 = i32(0)
l4:
	v9 = v1 + v8
	if v8 != i32(64) {
		goto l1
	}
	v1 = v7
	goto l2
l1:
	{
		if v6 == 0 {
			goto l3
		}
		t2 := int32(m.memory[uint32(v9+i32(14))])
		t3 := int32(load32(m.memory[int64(uint32(t2<<2))+1132860:]))
		t4 := int32(m.memory[uint32(v9+i32(15))])
		t5 := int32(load32(m.memory[int64(uint32(t4<<2))+1131836:]))
		t6 := int32(m.memory[uint32(v9+i32(13))])
		t7 := int32(load32(m.memory[int64(uint32(t6<<2))+1133884:]))
		t8 := int32(m.memory[uint32(v9+i32(12))])
		t9 := int32(load32(m.memory[int64(uint32(t8<<2))+1134908:]))
		t10 := int32(m.memory[uint32(v9+i32(11))])
		t11 := int32(load32(m.memory[int64(uint32(t10<<2))+1135932:]))
		t12 := int32(m.memory[uint32(v9+i32(10))])
		t13 := int32(load32(m.memory[int64(uint32(t12<<2))+1136956:]))
		t14 := int32(m.memory[uint32(v9+i32(9))])
		t15 := int32(load32(m.memory[int64(uint32(t14<<2))+1137980:]))
		t16 := int32(m.memory[uint32(v9+i32(8))])
		t17 := int32(load32(m.memory[int64(uint32(t16<<2))+1139004:]))
		t18 := int32(m.memory[uint32(v9+i32(7))])
		t19 := int32(load32(m.memory[int64(uint32(t18<<2))+1140028:]))
		t20 := int32(m.memory[uint32(v9+i32(6))])
		t21 := int32(load32(m.memory[int64(uint32(t20<<2))+1141052:]))
		t22 := int32(m.memory[uint32(v9+i32(5))])
		t23 := int32(load32(m.memory[int64(uint32(t22<<2))+1142076:]))
		t24 := int32(m.memory[uint32(v9+i32(4))])
		t25 := int32(load32(m.memory[int64(uint32(t24<<2))+1143100:]))
		t26 := int32(m.memory[uint32(v9+i32(3))])
		t27 := int32(load32(m.memory[int64(uint32((int32(uint32(v5)>>24)^t26)<<2))+1144124:]))
		t28 := int32(m.memory[uint32(v9+i32(2))])
		t29 := int32(load32(m.memory[int64(uint32((int32(uint32(v5)>>16)&i32(255)^t28)<<2))+1145148:]))
		t30 := int32(m.memory[uint32(v9+i32(1))])
		t31 := int32(load32(m.memory[int64(uint32((int32(uint32(v5)>>8)&i32(255)^t30)<<2))+1146172:]))
		t32 := int32(m.memory[uint32(v9)])
		t33 := int32(load32(m.memory[int64(uint32((v5&i32(255)^t32)<<2))+1147196:]))
		v5 = t3 ^ t5 ^ t7 ^ t9 ^ t11 ^ t13 ^ t15 ^ t17 ^ t19 ^ t21 ^ t23 ^ t25 ^ t27 ^ t29 ^ t31 ^ t33
		v8 = v8 + i32(16)
		v6 = v6 + i32(-1)
		goto l4
	}
l3:
	m.fn158(i32(15), v2, i32(1148220))
	panic("unreachable")
l6:
	{
		if v3 == 0 {
			goto l5
		}
		t34 := int32(m.memory[uint32(v4)])
		t35 := int32(load32(m.memory[int64(uint32((t34^v5)&i32(255)<<2))+1131836:]))
		v5 = t35 ^ int32(uint32(v5)>>8)
		v3 = v3 + i32(-1)
		v4 = v4 + i32(1)
		goto l6
	}
l5:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v5^i32(-1)))
}
func (m *Module) fn305(v0, v1, v2 int32) {
	var v3, v4, v5, v6, v7 int32
	var v8, v9 int64
	t0 := m.g0
	v3 = t0 - i32(16)
	m.g0 = v3
	t1 := int32(load32(m.memory[uint32(v1):]))
	v4 = t1
	{
		t2 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		v5 = t2
		t3 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		t4 := v5
		v6 = t3
		if uint32(t4) < uint32(v6) {
			goto l0
		}
		t5 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		v5 = t5
		{
			t6 := int32(m.memory[int64(uint32(v1))+16])
			if t6 != 0 {
				goto l1
			}
			if v5 == 0 {
				goto l1
			}
			memory_zero(m.memory, uint32(v4), uint32(v5))
		}
	l1:
		m.fn302(v3+i32(8), v2, v4, v5)
		v7 = i32(255)
		{
			{
				t7 := int32(m.memory[int64(uint32(v3))+8])
				if t7 == i32(255) {
					goto l2
				}
				t8 := int64(load64(m.memory[int64(uint32(v3))+8:]))
				v8 = t8
				v9 = v8 & i64(-256)
				v7 = int32(v8)
				v6 = i32(0)
				goto l3
			}
		l2:
			v9 = i64(0)
			t9 := int32(load32(m.memory[int64(uint32(v3))+12:]))
			v6 = t9
			if uint32(v6) > uint32(v5) {
				m.fn256(i32(1072353), i32(36), i32(1072392))
				panic("unreachable")
			}
		}
	l3:
		v2 = i32(1)
		m.memory[int64(uint32(v1))+16] = byte(i32(1))
		store32(m.memory[int64(uint32(v1))+12:], uint32(v6))
		v5 = i32(0)
		store32(m.memory[int64(uint32(v1))+8:], uint32(i32(0)))
		if v7&i32(255) != i32(255) {
			goto l5
		}
	}
l0:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v6-v5))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v4+v5))
	v2 = i32(0)
	goto l6
l5:
	store64(m.memory[int64(uint32(v0))+4:], uint64(v9|int64(uint32(v7))&i64(255)))
l6:
	store32(m.memory[uint32(v0):], uint32(v2))
	m.g0 = v3 + i32(16)
}
func (m *Module) fn306(v0, v1, v2, v3, v4, v5, v6 int32) {
	var v7, v8, v9 int32
	var v10, v11 int64
	var v12, v13, v14, v15, v16, v17, v18, v19, v20, v21 int32
	var v22, v23, v24 int64
	var v25, v26, v27, v28, v29, v30, v31, v32, v33, v34, v35 int32
	t0 := m.g0
	v7 = t0 - i32(176)
	m.g0 = v7
	store32(m.memory[int64(uint32(v1))+48:], uint32(v5))
	store32(m.memory[int64(uint32(v1))+36:], uint32(v3))
	store32(m.memory[int64(uint32(v1))+44:], uint32(v4))
	store32(m.memory[int64(uint32(v1))+32:], uint32(v2))
	{
		t1 := int32(load32(m.memory[int64(uint32(v1))+60:]))
		v8 = t1
		t2 := int32(m.memory[uint32(v8)])
		v9 = t2
		if v9 != i32(12) {
			goto l0
		}
		v9 = i32(13)
		m.memory[uint32(v8)] = byte(i32(13))
	}
l0:
	t3 := int64(load64(m.memory[int64(uint32(v1))+24:]))
	v10 = t3
	t4 := int64(load64(m.memory[int64(uint32(v1))+16:]))
	v11 = t4
	store32(m.memory[int64(uint32(v8))+116:], uint32(v5))
	store32(m.memory[int64(uint32(v8))+112:], uint32(v3))
	store32(m.memory[int64(uint32(v8))+80:], uint32(i32(0)))
	store32(m.memory[int64(uint32(v8))+76:], uint32(v5))
	store32(m.memory[int64(uint32(v8))+72:], uint32(v4))
	store32(m.memory[int64(uint32(v8))+56:], uint32(v2))
	m.memory[int64(uint32(v8))+4] = byte(v6)
	store32(m.memory[int64(uint32(v8))+60:], uint32(v2+v3))
	v12 = v8 + i32(128)
	v13 = v8 + i32(65)
	v14 = v8 + i32(8)
	v15 = v8 + i32(164)
	v16 = v8 + i32(10820)
	v17 = v8 + i32(13828)
	v18 = v8 + i32(5492)
	v19 = v8 + i32(13188)
	v20 = v8 + i32(72)
l18:
	v3 = v9 & i32(255)
	v5 = i32(1)
	v21 = i32(29)
	v9 = i32(18)
	{
		{
			switch v3 {
			case 4:
				t241 := int32(load32(m.memory[int64(uint32(v8))+120:]))
				v21 = t241
				if v21&i32(1024) == 0 {
					v9 = i32(5)
					t250 := int32(load32(m.memory[int64(uint32(v8))+140:]))
					v5 = t250
					if v5 == 0 {
						goto l18
					}
					store32(m.memory[int64(uint32(v5))+16:], uint32(i32(0)))
					goto l18
				}
				t242 := int64(m.memory[int64(uint32(v8))+64])
				v22 = t242
				t243 := int32(load32(m.memory[int64(uint32(v8))+56:]))
				v9 = t243
				t244 := int32(load32(m.memory[int64(uint32(v8))+60:]))
				v3 = t244
				t245 := int64(load64(m.memory[int64(uint32(v8))+48:]))
				v23 = t245
			l122:
				if uint64(v22) > uint64(i64(15)) {
					t251 := v8
					v9 = int32(v23)
					store32(m.memory[int64(uint32(t251))+88:], uint32(v9))
					{
						t252 := int32(load32(m.memory[int64(uint32(v8))+140:]))
						v5 = t252
						if v5 == 0 {
							goto l123
						}
						store32(m.memory[int64(uint32(v5))+20:], uint32(v9))
						t253 := int32(load32(m.memory[int64(uint32(v8))+120:]))
						v21 = t253
					}
				l123:
					{
						if v21&i32(512) == 0 {
							goto l124
						}
						t254 := int32(m.memory[int64(uint32(v8))+3])
						if t254&i32(4) == 0 {
							goto l124
						}
						t255 := int64(load64(m.memory[int64(uint32(v8))+48:]))
						store16(m.memory[int64(uint32(v7))+160:], uint16(t255))
						t256 := int32(load32(m.memory[int64(uint32(v8))+124:]))
						t257 := m.fn1729(t256, v7+i32(160), i32(2))
						store32(m.memory[int64(uint32(v8))+124:], uint32(t257))
					}
				l124:
					m.memory[int64(uint32(v8))+64] = byte(i32(0))
					store64(m.memory[int64(uint32(v8))+48:], uint64(i64(0)))
					v9 = i32(5)
					goto l18
				}
				if v9 != v3 {
					t246 := v8
					v5 = v9 + i32(1)
					store32(m.memory[int64(uint32(t246))+56:], uint32(v5))
					t247 := v8
					v24 = v22 + i64(8)
					m.memory[int64(uint32(t247))+64] = byte(v24)
					t248 := int64(m.memory[uint32(v9)])
					t249 := v8
					v23 = i64_shl(t248, v22) | v23
					store64(m.memory[int64(uint32(t249))+48:], uint64(v23))
					v22 = v24
					v9 = v5
					goto l122
				}
				v5 = i32(0)
				v21 = i32(4)
				goto l30
			case 5:
				{
					t210 := int32(load32(m.memory[int64(uint32(v8))+120:]))
					v3 = t210
					if v3&i32(1024) == 0 {
						goto l113
					}
					{
						t211 := int32(load32(m.memory[int64(uint32(v8))+60:]))
						t212 := int32(load32(m.memory[int64(uint32(v8))+56:]))
						v25 = t212
						v5 = t211 - v25
						t213 := int32(load32(m.memory[int64(uint32(v8))+88:]))
						t214 := v5
						v9 = t213
						p215 := v9
						if uint32(v5) < uint32(v9) {
							p215 = t214
						}
						v5 = p215
						if v5 == 0 {
							goto l114
						}
						{
							t216 := int32(load32(m.memory[int64(uint32(v8))+140:]))
							v21 = t216
							if v21 == 0 {
								goto l115
							}
							t217 := int32(load32(m.memory[int64(uint32(v21))+16:]))
							v26 = t217
							if v26 == 0 {
								goto l115
							}
							{
								t218 := int32(load32(m.memory[int64(uint32(v21))+24:]))
								t219 := v5
								v3 = t218
								t220 := int32(load32(m.memory[int64(uint32(v21))+20:]))
								t221 := v3
								v9 = t220 - v9
								v21 = t221 - v9
								p222 := v21
								if uint32(v21) > uint32(v3) {
									p222 = i32(0)
								}
								v21 = p222
								p223 := v21
								if uint32(v5) < uint32(v21) {
									p223 = t219
								}
								v21 = p223
								if v21 == 0 {
									goto l116
								}
								t225 := v26
								p224 := v9
								if uint32(v3) < uint32(v9) {
									p224 = v3
								}
								memory_copy(m.memory, uint32(t225+p224), uint32(v25), uint32(v21))
							}
						l116:
							t226 := int32(load32(m.memory[int64(uint32(v8))+120:]))
							v3 = t226
						}
					l115:
						{
							if v3&i32(512) == 0 {
								goto l117
							}
							t227 := int32(m.memory[int64(uint32(v8))+3])
							if t227&i32(4) == 0 {
								goto l117
							}
							{
								t228 := int32(load32(m.memory[int64(uint32(v8))+60:]))
								t229 := int32(load32(m.memory[int64(uint32(v8))+56:]))
								t230 := v5
								v9 = t229
								v3 = t228 - v9
								if uint32(t230) > uint32(v3) {
									m.fn151(i32(0), v5, v3, i32(1294756))
									panic("unreachable")
								}
								t231 := int32(load32(m.memory[int64(uint32(v8))+124:]))
								t232 := m.fn1729(t231, v9, v5)
								store32(m.memory[int64(uint32(v8))+124:], uint32(t232))
								goto l117
							}
						}
					l117:
						t233 := int32(load32(m.memory[int64(uint32(v8))+112:]))
						store32(m.memory[int64(uint32(v8))+112:], uint32(t233-v5))
						t234 := int32(load32(m.memory[int64(uint32(v8))+88:]))
						t235 := v8
						v9 = t234 - v5
						store32(m.memory[int64(uint32(t235))+88:], uint32(v9))
						t236 := int32(load32(m.memory[int64(uint32(v8))+60:]))
						t237 := v8
						v3 = t236
						t238 := int32(load32(m.memory[int64(uint32(v8))+56:]))
						t239 := v3
						v5 = t238 + v5
						p240 := v5
						if uint32(v3) < uint32(v5) {
							p240 = t239
						}
						store32(m.memory[int64(uint32(t237))+56:], uint32(p240))
					}
				l114:
					if v9 == 0 {
						goto l113
					}
					v5 = i32(0)
					v21 = i32(5)
					goto l30
				}
			l113:
				store32(m.memory[int64(uint32(v8))+88:], uint32(i32(0)))
				v9 = i32(6)
				goto l18
			case 6:
				{
					{
						t183 := int32(m.memory[int64(uint32(v8))+121])
						if t183&i32(8) == 0 {
							t185 := int32(load32(m.memory[int64(uint32(v8))+140:]))
							v9 = t185
							if v9 == 0 {
								goto l104
							}
							store32(m.memory[int64(uint32(v9))+28:], uint32(i32(0)))
							goto l104
						}
						v21 = i32(6)
						t184 := int32(load32(m.memory[int64(uint32(v8))+112:]))
						if t184 != 0 {
							goto l103
						}
						goto l93
					}
				l103:
					t186 := int32(load32(m.memory[int64(uint32(v8))+60:]))
					v25 = t186
					t187 := int32(load32(m.memory[int64(uint32(v8))+56:]))
					t188 := v25
					v3 = t187
					v26 = t188 - v3
					v9 = i32(0)
					{
					l107:
						v5 = v3 + v9
						if v5 == v25 {
							goto l105
						}
						{
							t189 := int32(m.memory[uint32(v5)])
							if t189 == 0 {
								goto l106
							}
							v9 = v9 + i32(1)
							goto l107
						}
					l106:
						m.fn1728(v7+i32(64), v9, v3, v26, i32(1294772))
						t190 := int32(load32(m.memory[int64(uint32(v7))+68:]))
						v26 = t190
						t191 := int32(load32(m.memory[int64(uint32(v7))+64:]))
						v3 = t191
					}
				l105:
					{
						t192 := int32(load32(m.memory[int64(uint32(v8))+140:]))
						v9 = t192
						if v9 == 0 {
							goto l108
						}
						t193 := int32(load32(m.memory[int64(uint32(v9))+28:]))
						v25 = t193
						if v25 == 0 {
							goto l108
						}
						{
							t194 := int32(load32(m.memory[int64(uint32(v9))+32:]))
							v5 = t194
							t195 := int32(load32(m.memory[int64(uint32(v8))+88:]))
							t196 := v5
							v9 = t195
							if uint32(t196) < uint32(v9) {
								m.fn633(i32(1294788), i32(18), i32(1294808))
								panic("unreachable")
							}
							{
								v5 = v5 - v9
								p197 := v26
								if uint32(v5) < uint32(v26) {
									p197 = v5
								}
								v5 = p197
								if v5 == 0 {
									goto l110
								}
								memory_copy(m.memory, uint32(v25+v9), uint32(v3), uint32(v5))
							}
						l110:
							t198 := int32(load32(m.memory[int64(uint32(v8))+88:]))
							store32(m.memory[int64(uint32(v8))+88:], uint32(t198+v5))
							goto l108
						}
					}
				l108:
					{
						t199 := int32(m.memory[int64(uint32(v8))+121])
						if t199&i32(2) == 0 {
							goto l111
						}
						t200 := int32(m.memory[int64(uint32(v8))+3])
						if t200&i32(4) == 0 {
							goto l111
						}
						t201 := int32(load32(m.memory[int64(uint32(v8))+124:]))
						t202 := m.fn1729(t201, v3, v26)
						store32(m.memory[int64(uint32(v8))+124:], uint32(t202))
					}
				l111:
					v5 = i32(0)
					v9 = i32(0)
					{
						if v26 == 0 {
							goto l112
						}
						v9 = i32(0)
						v3 = v3 + v26 + i32(-1)
						if v3 == 0 {
							goto l112
						}
						t203 := int32(m.memory[uint32(v3)])
						var p204 int32
						if t203 == 0 {
							p204 = 1
						}
						v9 = p204
					}
				l112:
					t205 := int32(load32(m.memory[int64(uint32(v8))+60:]))
					t206 := v8
					v3 = t205
					t207 := int32(load32(m.memory[int64(uint32(v8))+56:]))
					t208 := v3
					v25 = t207 + v26
					p209 := v25
					if uint32(v3) < uint32(v25) {
						p209 = t208
					}
					v25 = p209
					store32(m.memory[int64(uint32(t206))+56:], uint32(v25))
					if v9 != 0 {
						goto l104
					}
					if v3 == v25 {
						goto l30
					}
				}
			l104:
				store32(m.memory[int64(uint32(v8))+88:], uint32(i32(0)))
				v9 = i32(7)
				goto l18
			case 7:
				{
					{
						t156 := int32(m.memory[int64(uint32(v8))+121])
						if t156&i32(16) == 0 {
							v9 = i32(8)
							t162 := int32(load32(m.memory[int64(uint32(v8))+140:]))
							v5 = t162
							if v5 == 0 {
								goto l18
							}
							store32(m.memory[int64(uint32(v5))+36:], uint32(i32(0)))
							goto l18
						}
						v21 = i32(7)
						t157 := int32(load32(m.memory[int64(uint32(v8))+112:]))
						if t157 == 0 {
							goto l93
						}
						t158 := int32(load32(m.memory[int64(uint32(v8))+60:]))
						v25 = t158
						t159 := int32(load32(m.memory[int64(uint32(v8))+56:]))
						t160 := v25
						v3 = t159
						v26 = t160 - v3
						v9 = i32(0)
					l96:
						{
							v5 = v3 + v9
							if v5 == v25 {
								goto l94
							}
							t161 := int32(m.memory[uint32(v5)])
							if t161 == 0 {
								goto l95
							}
							v9 = v9 + i32(1)
							goto l96
						}
					}
				l95:
					m.fn1728(v7+i32(72), v9, v3, v26, i32(1294824))
					t163 := int32(load32(m.memory[int64(uint32(v7))+76:]))
					v26 = t163
					t164 := int32(load32(m.memory[int64(uint32(v7))+72:]))
					v3 = t164
				}
			l94:
				{
					t165 := int32(load32(m.memory[int64(uint32(v8))+140:]))
					v9 = t165
					if v9 == 0 {
						goto l97
					}
					t166 := int32(load32(m.memory[int64(uint32(v9))+36:]))
					v25 = t166
					if v25 == 0 {
						goto l97
					}
					{
						t167 := int32(load32(m.memory[int64(uint32(v9))+40:]))
						v5 = t167
						t168 := int32(load32(m.memory[int64(uint32(v8))+88:]))
						t169 := v5
						v9 = t168
						if uint32(t169) < uint32(v9) {
							m.fn633(i32(1294840), i32(18), i32(1294860))
							panic("unreachable")
						}
						{
							v5 = v5 - v9
							p170 := v26
							if uint32(v5) < uint32(v26) {
								p170 = v5
							}
							v5 = p170
							if v5 == 0 {
								goto l99
							}
							memory_copy(m.memory, uint32(v25+v9), uint32(v3), uint32(v5))
						}
					l99:
						t171 := int32(load32(m.memory[int64(uint32(v8))+88:]))
						store32(m.memory[int64(uint32(v8))+88:], uint32(t171+v5))
						goto l97
					}
				}
			l97:
				{
					t172 := int32(m.memory[int64(uint32(v8))+121])
					if t172&i32(2) == 0 {
						goto l100
					}
					t173 := int32(m.memory[int64(uint32(v8))+3])
					if t173&i32(4) == 0 {
						goto l100
					}
					t174 := int32(load32(m.memory[int64(uint32(v8))+124:]))
					t175 := m.fn1729(t174, v3, v26)
					store32(m.memory[int64(uint32(v8))+124:], uint32(t175))
				}
			l100:
				v5 = i32(0)
				v25 = i32(0)
				{
					if v26 == 0 {
						goto l101
					}
					v25 = i32(0)
					v9 = v3 + v26 + i32(-1)
					if v9 == 0 {
						goto l101
					}
					t176 := int32(m.memory[uint32(v9)])
					var p177 int32
					if t176 == 0 {
						p177 = 1
					}
					v25 = p177
				}
			l101:
				t178 := int32(load32(m.memory[int64(uint32(v8))+60:]))
				t179 := v8
				v3 = t178
				t180 := int32(load32(m.memory[int64(uint32(v8))+56:]))
				t181 := v3
				v9 = t180 + v26
				p182 := v9
				if uint32(v3) < uint32(v9) {
					p182 = t181
				}
				v26 = p182
				store32(m.memory[int64(uint32(t179))+56:], uint32(v26))
				v9 = i32(8)
				if v25 != 0 {
					goto l18
				}
				if v3 != v26 {
					goto l18
				}
				goto l30
			case 8:
				{
					t141 := int32(load32(m.memory[int64(uint32(v8))+120:]))
					v21 = t141
					if v21&i32(512) == 0 {
						goto l85
					}
					t142 := int64(m.memory[int64(uint32(v8))+64])
					v22 = t142
					t143 := int32(load32(m.memory[int64(uint32(v8))+56:]))
					v9 = t143
					t144 := int32(load32(m.memory[int64(uint32(v8))+60:]))
					v3 = t144
					t145 := int64(load64(m.memory[int64(uint32(v8))+48:]))
					v23 = t145
				l88:
					if uint64(v22) > uint64(i64(15)) {
						{
							t150 := int32(m.memory[int64(uint32(v8))+3])
							if t150&i32(4) == 0 {
								goto l89
							}
							t151 := int32(load16(m.memory[int64(uint32(v8))+124:]))
							if t151 != int32(v23) {
								store32(m.memory[int64(uint32(v8))+136:], uint32(i32(20)))
								store32(m.memory[int64(uint32(v8))+132:], uint32(i32(1067520)))
								goto l58
							}
						}
					l89:
						m.memory[int64(uint32(v8))+64] = byte(i32(0))
						store64(m.memory[int64(uint32(v8))+48:], uint64(i64(0)))
						goto l85
					}
					if v9 != v3 {
						t146 := v8
						v5 = v9 + i32(1)
						store32(m.memory[int64(uint32(t146))+56:], uint32(v5))
						t147 := v8
						v24 = v22 + i64(8)
						m.memory[int64(uint32(t147))+64] = byte(v24)
						t148 := int64(m.memory[uint32(v9)])
						t149 := v8
						v23 = i64_shl(t148, v22) | v23
						store64(m.memory[int64(uint32(t149))+48:], uint64(v23))
						v22 = v24
						v9 = v5
						goto l88
					}
					v5 = i32(0)
					v21 = i32(8)
					goto l30
				}
			l85:
				{
					t152 := int32(load32(m.memory[int64(uint32(v8))+140:]))
					v9 = t152
					if v9 == 0 {
						goto l91
					}
					store32(m.memory[int64(uint32(v9))+44:], uint32(int32(uint32(v21)>>9)&i32(1)))
					t153 := int32(load32(m.memory[int64(uint32(v8))+140:]))
					store32(m.memory[int64(uint32(t153))+48:], uint32(i32(1)))
					t154 := int32(load32(m.memory[int64(uint32(v8))+120:]))
					v21 = t154
				}
			l91:
				v9 = i32(12)
				t155 := int32(m.memory[int64(uint32(v8))+3])
				if t155&i32(4) == 0 {
					goto l18
				}
				if v21 == 0 {
					goto l18
				}
				store64(m.memory[int64(uint32(v8))+124:], uint64(i64(0)))
				goto l18
			case 13:
				t125 := int32(m.memory[int64(uint32(v8))+64])
				v3 = t125
				{
					t126 := int32(m.memory[int64(uint32(v8))+1])
					v26 = t126
					if v26&i32(1) == 0 {
						v5 = v3 | i32(8)
						t128 := int32(load32(m.memory[int64(uint32(v8))+56:]))
						v9 = t128
						t129 := int32(load32(m.memory[int64(uint32(v8))+60:]))
						v25 = t129
						t130 := int64(load64(m.memory[int64(uint32(v8))+48:]))
						v22 = t130
					l75:
						if uint32(v3&i32(255)) > uint32(i32(2)) {
							m.memory[int64(uint32(v8))+64] = byte(v3 + i32(-1))
							t134 := v8
							v23 = int64(uint64(v22) >> 1)
							store64(m.memory[int64(uint32(t134))+48:], uint64(v23))
							m.memory[int64(uint32(v8))+1] = byte(int32(v22)&i32(1) | v26)
							switch int32(v23) & i32(3) {
							default:
								m.memory[int64(uint32(v8))+64] = byte(v3 + i32(-3))
								store64(m.memory[int64(uint32(v8))+48:], uint64(int64(uint64(v22)>>3)))
								v9 = i32(14)
								goto l18
							case 1:
								v5 = i32(0)
								m.memory[int64(uint32(v8))+160] = byte(i32(0))
								store32(m.memory[int64(uint32(v8))+156:], uint32(i32(5)))
								m.memory[int64(uint32(v8))+152] = byte(i32(0))
								store32(m.memory[int64(uint32(v8))+148:], uint32(i32(9)))
								m.memory[int64(uint32(v8))+64] = byte(v3 + i32(-3))
								store64(m.memory[int64(uint32(v8))+48:], uint64(int64(uint64(v22)>>3)))
								v9 = i32(17)
								v21 = i32(17)
								t135 := int32(m.memory[int64(uint32(v8))+4])
								if t135 != i32(6) {
									goto l18
								}
								goto l30
							case 2:
								m.memory[int64(uint32(v8))+64] = byte(v3 + i32(-3))
								store64(m.memory[int64(uint32(v8))+48:], uint64(int64(uint64(v22)>>3)))
								v9 = i32(24)
								goto l18
							case 3:
								store32(m.memory[int64(uint32(v8))+136:], uint32(i32(19)))
								store32(m.memory[int64(uint32(v8))+132:], uint32(i32(1067785)))
								v5 = i32(-3)
								m.memory[int64(uint32(v8))+64] = byte(v3 + i32(-3))
								store64(m.memory[int64(uint32(v8))+48:], uint64(int64(uint64(v22)>>3)))
								v21 = i32(30)
								goto l30
							}
						}
						if v9 != v25 {
							m.memory[int64(uint32(v8))+64] = byte(v5)
							t131 := v8
							v21 = v9 + i32(1)
							store32(m.memory[int64(uint32(t131))+56:], uint32(v21))
							t132 := int64(m.memory[uint32(v9)])
							t133 := v8
							v22 = i64_shl(t132, int64(uint32(v3))&i64(255)) | v22
							store64(m.memory[int64(uint32(t133))+48:], uint64(v22))
							v9 = v21
							v3 = v5
							goto l75
						}
						v5 = i32(0)
						v21 = i32(13)
						goto l30
					}
					m.memory[int64(uint32(v8))+64] = byte(v3 & i32(248))
					t127 := int64(load64(m.memory[int64(uint32(v8))+48:]))
					store64(m.memory[int64(uint32(v8))+48:], uint64(i64_shr_u(t127, int64(uint32(v3))&i64(7))))
					v9 = i32(16)
					goto l18
				}
			case 17:
				goto l18
			case 18:
				m.memory[uint32(v8)] = byte(i32(18))
				{
					t449 := int32(load32(m.memory[int64(uint32(v8))+60:]))
					v3 = t449
					t450 := int32(load32(m.memory[int64(uint32(v8))+56:]))
					t451 := v3
					v30 = t450
					if uint32(t451-v30) < uint32(i32(15)) {
						goto l208
					}
					t452 := int32(load32(m.memory[int64(uint32(v8))+76:]))
					t453 := int32(load32(m.memory[int64(uint32(v8))+80:]))
					if uint32(t452-t453) <= uint32(i32(259)) {
						goto l208
					}
					m.fn1738(v8)
					t454 := int32(m.memory[uint32(v8)])
					v9 = t454
					if v9 != i32(18) {
						goto l18
					}
					t455 := int32(load32(m.memory[int64(uint32(v8))+60:]))
					v3 = t455
					t456 := int32(load32(m.memory[int64(uint32(v8))+56:]))
					v30 = t456
				}
			l208:
				t457 := int32(load32(m.memory[int64(uint32(v20))+8:]))
				v9 = t457
				store32(m.memory[int64(uint32(v8))+80:], uint32(i32(0)))
				t458 := int64(load64(m.memory[uint32(v20):]))
				v22 = t458
				store64(m.memory[int64(uint32(v8))+72:], uint64(i64(1)))
				store32(m.memory[int64(uint32(v7))+168:], uint32(v9))
				store64(m.memory[int64(uint32(v7))+160:], uint64(v22))
				t459 := int32(load32(m.memory[uint32(v13):]))
				store32(m.memory[int64(uint32(v7))+152:], uint32(t459))
				t460 := int32(load32(m.memory[int64(uint32(v13))+3:]))
				store32(m.memory[int64(uint32(v7))+155:], uint32(t460))
				v29 = i32(1292484)
				v27 = i32(512)
				{
					t461 := int32(m.memory[int64(uint32(v8))+152])
					switch t461 {
					default:
						goto l209
					case 1:
						v27 = i32(1332)
						v29 = v15
						goto l209
					case 2:
						v27 = i32(1332)
						v29 = v18
						goto l209
					case 3:
						v27 = i32(592)
						v29 = v16
					}
				}
			l209:
				v26 = i32(1294548)
				v25 = i32(32)
				{
					t462 := int32(m.memory[int64(uint32(v8))+160])
					switch t462 {
					default:
						goto l213
					case 1:
						v25 = i32(1332)
						v26 = v15
						goto l213
					case 2:
						v25 = i32(1332)
						v26 = v18
						goto l213
					case 3:
						v25 = i32(592)
						v26 = v16
					}
				}
			l213:
				t463 := int32(m.memory[int64(uint32(v8))+64])
				v31 = t463
				t464 := int64(load64(m.memory[int64(uint32(v8))+48:]))
				v22 = t464
				{
				l242:
					v21 = i32(18)
				l255:
					{
						{
							switch v21&i32(255) + i32(-18) {
							case 1:
								t486 := int32(load32(m.memory[int64(uint32(v7))+168:]))
								t487 := int32(load32(m.memory[int64(uint32(v7))+164:]))
								if t486 == t487 {
									v21 = i32(19)
									m.memory[uint32(v8)] = byte(i32(19))
									t489 := int64(load64(m.memory[int64(uint32(v7))+160:]))
									store64(m.memory[uint32(v20):], uint64(t489))
									t490 := int32(load32(m.memory[int64(uint32(v7))+168:]))
									store32(m.memory[int64(uint32(v20))+8:], uint32(t490))
									m.memory[int64(uint32(v8))+64] = byte(v31)
									store32(m.memory[int64(uint32(v8))+60:], uint32(v3))
									store32(m.memory[int64(uint32(v8))+56:], uint32(v30))
									store64(m.memory[int64(uint32(v8))+48:], uint64(v22))
									t491 := int32(load32(m.memory[int64(uint32(v7))+152:]))
									store32(m.memory[uint32(v13):], uint32(t491))
									t492 := int32(load32(m.memory[int64(uint32(v7))+155:]))
									store32(m.memory[int64(uint32(v13))+3:], uint32(t492))
									goto l93
								}
								t488 := int32(load32(m.memory[int64(uint32(v8))+88:]))
								m.fn1737(v7+i32(160), t488)
								goto l242
							case 3:
								t502 := int64(load32(m.memory[int64(uint32(v8))+156:]))
								v23 = i64_shl(i64(-1), t502) ^ i64(-1)
								v9 = v31
								v5 = v30
								{
								l249:
									{
										t503 := v25
										v21 = int32(v22 & v23)
										if uint32(t503) <= uint32(v21) {
											m.fn158(v21, v25, i32(1294708))
											panic("unreachable")
										}
										{
											v28 = v26 + v21<<2
											t504 := int32(m.memory[int64(uint32(v28))+3])
											v21 = t504
											if uint32(v21) <= uint32(v9&i32(255)) {
												t509 := int32(load16(m.memory[uint32(v28):]))
												v32 = t509
												t510 := int32(m.memory[int64(uint32(v28))+2])
												v33 = t510
												if uint32(v33) <= uint32(i32(15)) {
													goto l247
												}
												v28 = v21
												goto l248
											}
											if v5 != v3 {
												t511 := int64(m.memory[uint32(v5)])
												v22 = i64_shl(t511, int64(uint32(v9))) | v22
												v9 = v9 + i32(8)
												v5 = v5 + i32(1)
												goto l249
											}
											v21 = i32(21)
											m.memory[uint32(v8)] = byte(i32(21))
											t505 := int64(load64(m.memory[int64(uint32(v7))+160:]))
											store64(m.memory[uint32(v20):], uint64(t505))
											t506 := int32(load32(m.memory[int64(uint32(v7))+168:]))
											store32(m.memory[int64(uint32(v20))+8:], uint32(t506))
											store32(m.memory[int64(uint32(v8))+60:], uint32(v3))
											store32(m.memory[int64(uint32(v8))+56:], uint32(v3))
											store64(m.memory[int64(uint32(v8))+48:], uint64(v22))
											t507 := int32(load32(m.memory[int64(uint32(v7))+152:]))
											store32(m.memory[uint32(v13):], uint32(t507))
											t508 := int32(load32(m.memory[int64(uint32(v7))+155:]))
											store32(m.memory[int64(uint32(v13))+3:], uint32(t508))
											m.memory[int64(uint32(v8))+64] = byte(v31 + v3<<3 - v30<<3)
											goto l93
										}
									l247:
									}
									v35 = v21 & i32(31)
									v23 = i64_shl(i64(-1), int64(uint32(v33+v21))) ^ i64(-1)
								l253:
									v28 = i32_shr_u(int32(v22&v23), v35) + v32
									if uint32(v28) >= uint32(v25) {
										m.fn158(v28, v25, i32(1294724))
										panic("unreachable")
									}
									{
										v34 = v26 + v28<<2
										t512 := int32(m.memory[int64(uint32(v34))+3])
										v28 = t512
										if uint32((v28+v21)&i32(255)) <= uint32(v9&i32(255)) {
											goto l251
										}
										{
											if v5 != v3 {
												t517 := int64(m.memory[uint32(v5)])
												v22 = i64_shl(t517, int64(uint32(v9))) | v22
												v9 = v9 + i32(8)
												v5 = v5 + i32(1)
												goto l253
											}
											v21 = i32(21)
											m.memory[uint32(v8)] = byte(i32(21))
											t513 := int64(load64(m.memory[int64(uint32(v7))+160:]))
											store64(m.memory[uint32(v20):], uint64(t513))
											t514 := int32(load32(m.memory[int64(uint32(v7))+168:]))
											store32(m.memory[int64(uint32(v20))+8:], uint32(t514))
											store32(m.memory[int64(uint32(v8))+60:], uint32(v3))
											store32(m.memory[int64(uint32(v8))+56:], uint32(v3))
											store64(m.memory[int64(uint32(v8))+48:], uint64(v22))
											t515 := int32(load32(m.memory[int64(uint32(v7))+152:]))
											store32(m.memory[uint32(v13):], uint32(t515))
											t516 := int32(load32(m.memory[int64(uint32(v7))+155:]))
											store32(m.memory[int64(uint32(v13))+3:], uint32(t516))
											m.memory[int64(uint32(v8))+64] = byte(v31 + v3<<3 - v30<<3)
											goto l93
										}
									}
								l251:
									t518 := int32(m.memory[int64(uint32(v34))+2])
									v33 = t518
									t519 := int32(load16(m.memory[uint32(v34):]))
									v32 = t519
									t520 := int32(load32(m.memory[int64(uint32(v8))+100:]))
									store32(m.memory[int64(uint32(v8))+100:], uint32(t520+v21))
									v9 = v9 - v21
									v22 = i64_shr_u(v22, int64(uint32(v21)))
								}
							l248:
								v31 = v9 - v28
								v22 = i64_shr_u(v22, int64(uint32(v28)))
								if v33&i32(64) != 0 {
									t521 := int32(load32(m.memory[int64(uint32(v7))+168:]))
									store32(m.memory[int64(uint32(v20))+8:], uint32(t521))
									t522 := int64(load64(m.memory[int64(uint32(v7))+160:]))
									store64(m.memory[uint32(v20):], uint64(t522))
									m.memory[int64(uint32(v8))+64] = byte(v31)
									store32(m.memory[int64(uint32(v8))+60:], uint32(v3))
									store32(m.memory[int64(uint32(v8))+56:], uint32(v5))
									store64(m.memory[int64(uint32(v8))+48:], uint64(v22))
									t523 := int32(load32(m.memory[int64(uint32(v7))+152:]))
									store32(m.memory[uint32(v13):], uint32(t523))
									t524 := int32(load32(m.memory[int64(uint32(v7))+155:]))
									store32(m.memory[int64(uint32(v13))+3:], uint32(t524))
									store32(m.memory[int64(uint32(v8))+136:], uint32(i32(22)))
									store32(m.memory[int64(uint32(v8))+132:], uint32(i32(1067932)))
									v21 = i32(30)
									goto l256
								}
								store32(m.memory[int64(uint32(v8))+96:], uint32(v33&i32(15)))
								store32(m.memory[int64(uint32(v8))+92:], uint32(v32&i32(0xffff)))
								v21 = i32(22)
								v30 = v5
								goto l255
							case 4:
								v21 = i32(23)
								v9 = v31
								v5 = v30
								t525 := int32(load32(m.memory[int64(uint32(v8))+96:]))
								v28 = t525
								if v28 == 0 {
									goto l255
								}
								{
								l259:
									if uint32(v28) <= uint32(v9&i32(255)) {
										t531 := int32(load32(m.memory[int64(uint32(v8))+100:]))
										store32(m.memory[int64(uint32(v8))+100:], uint32(t531+v28))
										t532 := int32(load32(m.memory[int64(uint32(v8))+92:]))
										t533 := v8
										t534 := v22
										v23 = int64(uint32(v28))
										store32(m.memory[int64(uint32(t533))+92:], uint32(t532+int32(t534&(i64_shl(i64(-1), v23)^i64(-1)))))
										v31 = v9 - v28
										v22 = i64_shr_u(v22, v23)
										v30 = v5
										goto l255
									}
									{
										if v5 == v3 {
											v21 = i32(22)
											m.memory[uint32(v8)] = byte(i32(22))
											t527 := int64(load64(m.memory[int64(uint32(v7))+160:]))
											store64(m.memory[uint32(v20):], uint64(t527))
											t528 := int32(load32(m.memory[int64(uint32(v7))+168:]))
											store32(m.memory[int64(uint32(v20))+8:], uint32(t528))
											store32(m.memory[int64(uint32(v8))+60:], uint32(v3))
											store32(m.memory[int64(uint32(v8))+56:], uint32(v3))
											store64(m.memory[int64(uint32(v8))+48:], uint64(v22))
											t529 := int32(load32(m.memory[int64(uint32(v7))+152:]))
											store32(m.memory[uint32(v13):], uint32(t529))
											t530 := int32(load32(m.memory[int64(uint32(v7))+155:]))
											store32(m.memory[int64(uint32(v13))+3:], uint32(t530))
											m.memory[int64(uint32(v8))+64] = byte(v31 + v3<<3 - v30<<3)
											goto l93
										}
										t526 := int64(m.memory[uint32(v5)])
										v22 = i64_shl(t526, int64(uint32(v9))) | v22
										v9 = v9 + i32(8)
										v5 = v5 + i32(1)
										goto l259
									}
								}
							case 5:
								goto l222
							default:
								{
									if uint32(v3-v30) < uint32(i32(15)) {
										goto l223
									}
									t465 := int32(load32(m.memory[int64(uint32(v7))+164:]))
									t466 := int32(load32(m.memory[int64(uint32(v7))+168:]))
									if uint32(t465-t466) > uint32(i32(259)) {
										m.memory[uint32(v8)] = byte(i32(18))
										t567 := int64(load64(m.memory[int64(uint32(v7))+160:]))
										store64(m.memory[uint32(v20):], uint64(t567))
										t568 := int32(load32(m.memory[int64(uint32(v7))+168:]))
										store32(m.memory[int64(uint32(v20))+8:], uint32(t568))
										m.memory[int64(uint32(v8))+64] = byte(v31)
										store32(m.memory[int64(uint32(v8))+60:], uint32(v3))
										store32(m.memory[int64(uint32(v8))+56:], uint32(v30))
										store64(m.memory[int64(uint32(v8))+48:], uint64(v22))
										t569 := int32(load32(m.memory[int64(uint32(v7))+152:]))
										store32(m.memory[uint32(v13):], uint32(t569))
										t570 := int32(load32(m.memory[int64(uint32(v7))+155:]))
										store32(m.memory[int64(uint32(v13))+3:], uint32(t570))
										m.fn1738(v8)
										t571 := int32(m.memory[uint32(v8)])
										v9 = t571
										goto l18
									}
								}
							l223:
								store32(m.memory[int64(uint32(v8))+100:], uint32(i32(0)))
								t467 := int64(load32(m.memory[int64(uint32(v8))+148:]))
								v23 = i64_shl(i64(-1), t467) ^ i64(-1)
								v9 = v31
								v5 = v30
								{
								l233:
									{
										t470 := v27
										v21 = int32(v22 & v23)
										if uint32(t470) <= uint32(v21) {
											m.fn158(v21, v27, i32(1294676))
											panic("unreachable")
										}
										v28 = v29 + v21<<2
										t471 := int32(m.memory[int64(uint32(v28))+3])
										v21 = t471
										if uint32(v21) <= uint32(v9&i32(255)) {
											t477 := int32(load16(m.memory[uint32(v28):]))
											v32 = t477
											{
												t478 := int32(m.memory[int64(uint32(v28))+2])
												v33 = t478
												if v33 != 0 {
													v28 = i32(0)
													if uint32(v33) <= uint32(i32(15)) {
														v34 = v21 & i32(15)
														v23 = i64_shl(i64(-1), int64(uint32(v33+v21))) ^ i64(-1)
													l240:
														{
															t479 := v27
															v28 = (i32_shr_u(int32(v22&v23)&i32(0xffff), v34) + v32) & i32(0xffff)
															if uint32(t479) <= uint32(v28) {
																m.fn158(v28, v27, i32(1294692))
																panic("unreachable")
															}
															v28 = v29 + v28<<2
															t480 := int32(m.memory[int64(uint32(v28))+3])
															v35 = t480
															if uint32((v35+v21)&i32(255)) <= uint32(v9&i32(255)) {
																v9 = v9 - v21
																v22 = i64_shr_u(v22, int64(uint32(v21)))
																t572 := int32(m.memory[int64(uint32(v28))+2])
																v33 = t572
																t573 := int32(load16(m.memory[uint32(v28):]))
																v32 = t573
																v28 = v21
																v30 = v5
																v21 = v35
																goto l235
															}
															{
																if v5 != v3 {
																	t485 := int64(m.memory[uint32(v5)])
																	v22 = i64_shl(t485, int64(uint32(v9))) | v22
																	v9 = v9 + i32(8)
																	v5 = v5 + i32(1)
																	goto l240
																}
																v21 = i32(18)
																m.memory[uint32(v8)] = byte(i32(18))
																t481 := int64(load64(m.memory[int64(uint32(v7))+160:]))
																store64(m.memory[uint32(v20):], uint64(t481))
																t482 := int32(load32(m.memory[int64(uint32(v7))+168:]))
																store32(m.memory[int64(uint32(v20))+8:], uint32(t482))
																store32(m.memory[int64(uint32(v8))+60:], uint32(v3))
																store32(m.memory[int64(uint32(v8))+56:], uint32(v3))
																store64(m.memory[int64(uint32(v8))+48:], uint64(v22))
																t483 := int32(load32(m.memory[int64(uint32(v7))+152:]))
																store32(m.memory[uint32(v13):], uint32(t483))
																t484 := int32(load32(m.memory[int64(uint32(v7))+155:]))
																store32(m.memory[int64(uint32(v13))+3:], uint32(t484))
																m.memory[int64(uint32(v8))+64] = byte(v31 + v3<<3 - v30<<3)
																goto l93
															}
														}
													}
													v30 = v5
													goto l235
												}
												v28 = i32(0)
												v30 = v5
												v33 = i32(0)
												goto l235
											}
										}
										{
											if v5 == v3 {
												goto l232
											}
											t472 := int64(m.memory[uint32(v5)])
											v22 = i64_shl(t472, int64(uint32(v9))) | v22
											v9 = v9 + i32(8)
											v5 = v5 + i32(1)
											goto l233
										}
									l232:
									}
									v21 = i32(18)
									m.memory[uint32(v8)] = byte(i32(18))
									t473 := int64(load64(m.memory[int64(uint32(v7))+160:]))
									store64(m.memory[uint32(v20):], uint64(t473))
									t474 := int32(load32(m.memory[int64(uint32(v7))+168:]))
									store32(m.memory[int64(uint32(v20))+8:], uint32(t474))
									store32(m.memory[int64(uint32(v8))+60:], uint32(v3))
									store32(m.memory[int64(uint32(v8))+56:], uint32(v3))
									store64(m.memory[int64(uint32(v8))+48:], uint64(v22))
									t475 := int32(load32(m.memory[int64(uint32(v7))+152:]))
									store32(m.memory[uint32(v13):], uint32(t475))
									t476 := int32(load32(m.memory[int64(uint32(v7))+155:]))
									store32(m.memory[int64(uint32(v13))+3:], uint32(t476))
									m.memory[int64(uint32(v8))+64] = byte(v31 + v3<<3 - v30<<3)
									goto l93
								}
							case 2:
								t468 := int32(load32(m.memory[int64(uint32(v8))+96:]))
								v21 = t468
								if v21 == 0 {
									t493 := int32(load32(m.memory[int64(uint32(v8))+88:]))
									v28 = t493
									goto l243
								}
								v9 = v31
								v5 = v30
							l229:
								{
									if uint32(v21) <= uint32(v9&i32(255)) {
										t498 := int32(load32(m.memory[int64(uint32(v8))+100:]))
										store32(m.memory[int64(uint32(v8))+100:], uint32(t498+v21))
										t499 := int32(load32(m.memory[int64(uint32(v8))+88:]))
										t500 := v8
										t501 := v22
										v23 = int64(uint32(v21))
										v28 = t499 + int32(t501&(i64_shl(i64(-1), v23)^i64(-1)))
										store32(m.memory[int64(uint32(t500))+88:], uint32(v28))
										v31 = v9 - v21
										v22 = i64_shr_u(v22, v23)
										v30 = v5
										goto l243
									}
									if v5 == v3 {
										v21 = i32(20)
										m.memory[uint32(v8)] = byte(i32(20))
										t494 := int64(load64(m.memory[int64(uint32(v7))+160:]))
										store64(m.memory[uint32(v20):], uint64(t494))
										t495 := int32(load32(m.memory[int64(uint32(v7))+168:]))
										store32(m.memory[int64(uint32(v20))+8:], uint32(t495))
										store32(m.memory[int64(uint32(v8))+60:], uint32(v3))
										store32(m.memory[int64(uint32(v8))+56:], uint32(v3))
										store64(m.memory[int64(uint32(v8))+48:], uint64(v22))
										t496 := int32(load32(m.memory[int64(uint32(v7))+152:]))
										store32(m.memory[uint32(v13):], uint32(t496))
										t497 := int32(load32(m.memory[int64(uint32(v7))+155:]))
										store32(m.memory[int64(uint32(v13))+3:], uint32(t497))
										m.memory[int64(uint32(v8))+64] = byte(v31 + v3<<3 - v30<<3)
										goto l93
									}
									t469 := int64(m.memory[uint32(v5)])
									v22 = i64_shl(t469, int64(uint32(v9))) | v22
									v9 = v9 + i32(8)
									v5 = v5 + i32(1)
									goto l229
								}
							}
						l222:
							t535 := int32(load32(m.memory[int64(uint32(v7))+164:]))
							v5 = t535
							t536 := int32(load32(m.memory[int64(uint32(v7))+168:]))
							t537 := v5
							v9 = t536
							if t537 == v9 {
								v21 = i32(23)
								m.memory[uint32(v8)] = byte(i32(23))
								t574 := int64(load64(m.memory[int64(uint32(v7))+160:]))
								store64(m.memory[uint32(v20):], uint64(t574))
								t575 := int32(load32(m.memory[int64(uint32(v7))+168:]))
								store32(m.memory[int64(uint32(v20))+8:], uint32(t575))
								m.memory[int64(uint32(v8))+64] = byte(v31)
								store32(m.memory[int64(uint32(v8))+60:], uint32(v3))
								store32(m.memory[int64(uint32(v8))+56:], uint32(v30))
								store64(m.memory[int64(uint32(v8))+48:], uint64(v22))
								t576 := int32(load32(m.memory[int64(uint32(v7))+152:]))
								store32(m.memory[uint32(v13):], uint32(t576))
								t577 := int32(load32(m.memory[int64(uint32(v7))+155:]))
								store32(m.memory[int64(uint32(v13))+3:], uint32(t577))
								goto l93
							}
							v5 = v5 - v9
							{
								t538 := int32(load32(m.memory[int64(uint32(v8))+92:]))
								v21 = t538
								if uint32(v21) > uint32(v9) {
									goto l261
								}
								t539 := int32(load32(m.memory[int64(uint32(v8))+88:]))
								t540 := v7 + i32(160)
								t541 := v21
								t542 := v5
								v9 = t539
								p543 := v9
								if uint32(v5) < uint32(v9) {
									p543 = t542
								}
								v9 = p543
								m.fn1734(t540, t541, v9)
								goto l262
							}
						l261:
							{
								v9 = v21 - v9
								t544 := int32(load32(m.memory[int64(uint32(v8))+16:]))
								if uint32(v9) > uint32(t544) {
									goto l263
								}
								t545 := int32(load32(m.memory[int64(uint32(v8))+20:]))
								v21 = t545
								t546 := int32(load32(m.memory[uint32(v8+i32(12)):]))
								t547 := m.fn1735(t546)
								t548 := v7 + i32(160)
								t549 := v14
								v28 = v9 - v21
								t550 := t547 - v28
								t551 := v21 - v9
								var p552 int32
								if uint32(v9) > uint32(v21) {
									p552 = 1
								}
								v21 = p552
								p553 := t551
								if v21 != 0 {
									p553 = t550
								}
								v32 = p553
								t554 := int32(load32(m.memory[int64(uint32(v8))+88:]))
								t555 := v32
								t556 := v32
								t557 := v5
								v34 = t554
								t559 := v34
								p558 := v9
								if v21 != 0 {
									p558 = v28
								}
								v9 = p558
								p560 := v9
								if uint32(v34) < uint32(v9) {
									p560 = t559
								}
								v9 = p560
								p561 := v9
								if uint32(v5) < uint32(v9) {
									p561 = t557
								}
								v9 = p561
								m.fn1736(t548, t549, t555, t556+v9)
								goto l262
							}
						l263:
							t562 := int32(m.memory[int64(uint32(v8))+1])
							if t562&i32(4) == 0 {
								m.fn91(i32(1287716), i32(85), i32(1294740))
								panic("unreachable")
							}
							t563 := int32(load32(m.memory[int64(uint32(v7))+168:]))
							store32(m.memory[int64(uint32(v20))+8:], uint32(t563))
							t564 := int64(load64(m.memory[int64(uint32(v7))+160:]))
							store64(m.memory[uint32(v20):], uint64(t564))
							m.memory[int64(uint32(v8))+64] = byte(v31)
							store32(m.memory[int64(uint32(v8))+60:], uint32(v3))
							store32(m.memory[int64(uint32(v8))+56:], uint32(v30))
							store64(m.memory[int64(uint32(v8))+48:], uint64(v22))
							t565 := int32(load32(m.memory[int64(uint32(v7))+152:]))
							store32(m.memory[uint32(v13):], uint32(t565))
							t566 := int32(load32(m.memory[int64(uint32(v7))+155:]))
							store32(m.memory[int64(uint32(v13))+3:], uint32(t566))
							v21 = i32(30)
							store32(m.memory[int64(uint32(v8))+136:], uint32(i32(30)))
							store32(m.memory[int64(uint32(v8))+132:], uint32(i32(1067490)))
						}
					l256:
						m.memory[uint32(v8)] = byte(v21)
						v5 = i32(-3)
						goto l30
					l262:
						t578 := int32(load32(m.memory[int64(uint32(v8))+88:]))
						t579 := v8
						v5 = t578
						store32(m.memory[int64(uint32(t579))+88:], uint32(v5-v9))
						p580 := i32(23)
						if v5 == v9 {
							p580 = i32(18)
						}
						v21 = p580
						goto l255
					}
				l243:
					store32(m.memory[int64(uint32(v8))+104:], uint32(v28))
					v21 = i32(21)
					goto l255
				l235:
					store32(m.memory[int64(uint32(v8))+88:], uint32(v32&i32(0xffff)))
					store32(m.memory[int64(uint32(v8))+100:], uint32(v28+v21&i32(255)))
					v31 = v9 - v21
					v22 = i64_shr_u(v22, int64(uint32(v21)))
					v21 = i32(19)
					if v33&i32(255) == 0 {
						goto l255
					}
					if v33&i32(32) != 0 {
						v9 = i32(12)
						m.memory[uint32(v8)] = byte(i32(12))
						store32(m.memory[int64(uint32(v8))+100:], uint32(i32(-1)))
						t585 := int64(load64(m.memory[int64(uint32(v7))+160:]))
						store64(m.memory[uint32(v20):], uint64(t585))
						t586 := int32(load32(m.memory[int64(uint32(v7))+168:]))
						store32(m.memory[int64(uint32(v20))+8:], uint32(t586))
						m.memory[int64(uint32(v8))+64] = byte(v31)
						store32(m.memory[int64(uint32(v8))+60:], uint32(v3))
						store32(m.memory[int64(uint32(v8))+56:], uint32(v30))
						store64(m.memory[int64(uint32(v8))+48:], uint64(v22))
						t587 := int32(load32(m.memory[int64(uint32(v7))+152:]))
						store32(m.memory[uint32(v13):], uint32(t587))
						t588 := int32(load32(m.memory[int64(uint32(v7))+155:]))
						store32(m.memory[int64(uint32(v13))+3:], uint32(t588))
						goto l18
					}
					if v33&i32(64) != 0 {
						v21 = i32(30)
						m.memory[uint32(v8)] = byte(i32(30))
						t581 := int64(load64(m.memory[int64(uint32(v7))+160:]))
						store64(m.memory[uint32(v20):], uint64(t581))
						t582 := int32(load32(m.memory[int64(uint32(v7))+168:]))
						store32(m.memory[int64(uint32(v20))+8:], uint32(t582))
						m.memory[int64(uint32(v8))+64] = byte(v31)
						store32(m.memory[int64(uint32(v8))+60:], uint32(v3))
						store32(m.memory[int64(uint32(v8))+56:], uint32(v30))
						store64(m.memory[int64(uint32(v8))+48:], uint64(v22))
						t583 := int32(load32(m.memory[int64(uint32(v7))+152:]))
						store32(m.memory[uint32(v13):], uint32(t583))
						t584 := int32(load32(m.memory[int64(uint32(v7))+155:]))
						store32(m.memory[int64(uint32(v13))+3:], uint32(t584))
						store32(m.memory[int64(uint32(v8))+136:], uint32(i32(28)))
						store32(m.memory[int64(uint32(v8))+132:], uint32(i32(1067904)))
						v5 = i32(-3)
						goto l30
					}
					store32(m.memory[int64(uint32(v8))+96:], uint32(v33&i32(15)))
					v21 = i32(20)
					goto l255
				}
			case 19:
				t446 := int32(load32(m.memory[int64(uint32(v8))+80:]))
				t447 := int32(load32(m.memory[int64(uint32(v8))+76:]))
				if t446 != t447 {
					t448 := int32(load32(m.memory[int64(uint32(v8))+88:]))
					m.fn1737(v20, t448)
					goto l18
				}
				v5 = i32(0)
				v21 = i32(19)
				goto l30
			case 20:
				{
					{
						t431 := int32(load32(m.memory[int64(uint32(v8))+96:]))
						v25 = t431
						if v25 != 0 {
							goto l202
						}
						t432 := int32(load32(m.memory[int64(uint32(v8))+88:]))
						v9 = t432
						goto l203
					}
				l202:
					t433 := int32(load32(m.memory[int64(uint32(v8))+56:]))
					v5 = t433
					t434 := int32(load32(m.memory[int64(uint32(v8))+60:]))
					v26 = t434
					t435 := int64(load64(m.memory[int64(uint32(v8))+48:]))
					v22 = t435
					t436 := int32(m.memory[int64(uint32(v8))+64])
					v9 = t436
				l206:
					if uint32(v25) <= uint32(v9&i32(255)) {
						goto l204
					}
					if v5 != v26 {
						t437 := v8
						v3 = v5 + i32(1)
						store32(m.memory[int64(uint32(t437))+56:], uint32(v3))
						t438 := v8
						v21 = v9 + i32(8)
						m.memory[int64(uint32(t438))+64] = byte(v21)
						t439 := int64(m.memory[uint32(v5)])
						t440 := v8
						v22 = i64_shl(t439, int64(uint32(v9))) | v22
						store64(m.memory[int64(uint32(t440))+48:], uint64(v22))
						v5 = v3
						v9 = v21
						goto l206
					}
					v5 = i32(0)
					v21 = i32(20)
					goto l30
				l204:
					m.memory[int64(uint32(v8))+64] = byte(v9 - v25)
					t441 := v8
					t442 := v22
					v23 = int64(uint32(v25))
					store64(m.memory[int64(uint32(t441))+48:], uint64(i64_shr_u(t442, v23)))
					t443 := int32(load32(m.memory[int64(uint32(v8))+100:]))
					store32(m.memory[int64(uint32(v8))+100:], uint32(t443+v25))
					t444 := int32(load32(m.memory[int64(uint32(v8))+88:]))
					t445 := v8
					v9 = t444 + int32(v22&(i64_shl(i64(-1), v23)^i64(-1)))
					store32(m.memory[int64(uint32(t445))+88:], uint32(v9))
				}
			l203:
				store32(m.memory[int64(uint32(v8))+104:], uint32(v9))
				v9 = i32(21)
				goto l18
			case 23:
			l201:
				{
					t401 := int32(load32(m.memory[int64(uint32(v8))+76:]))
					v3 = t401
					t402 := int32(load32(m.memory[int64(uint32(v8))+80:]))
					t403 := v3
					v5 = t402
					if t403 != v5 {
						v3 = v3 - v5
						{
							t404 := int32(load32(m.memory[int64(uint32(v8))+92:]))
							v21 = t404
							if uint32(v21) > uint32(v5) {
								v5 = v21 - v5
								t410 := int32(load32(m.memory[int64(uint32(v8))+16:]))
								if uint32(v5) > uint32(t410) {
									t428 := int32(m.memory[int64(uint32(v8))+1])
									if t428&i32(4) == 0 {
										m.fn91(i32(1287716), i32(85), i32(1294876))
										panic("unreachable")
									}
									v21 = i32(30)
									store32(m.memory[int64(uint32(v8))+136:], uint32(i32(30)))
									store32(m.memory[int64(uint32(v8))+132:], uint32(i32(1067490)))
									v5 = i32(-3)
									goto l30
								}
								t411 := int32(load32(m.memory[int64(uint32(v8))+20:]))
								v21 = t411
								t412 := int32(load32(m.memory[uint32(v8+i32(12)):]))
								t413 := m.fn1735(t412)
								t414 := v20
								t415 := v14
								v25 = v5 - v21
								t416 := t413 - v25
								t417 := v21 - v5
								var p418 int32
								if uint32(v5) > uint32(v21) {
									p418 = 1
								}
								v21 = p418
								p419 := t417
								if v21 != 0 {
									p419 = t416
								}
								v26 = p419
								t420 := int32(load32(m.memory[int64(uint32(v8))+88:]))
								t421 := v26
								t422 := v26
								t423 := v3
								v27 = t420
								t425 := v27
								p424 := v5
								if v21 != 0 {
									p424 = v25
								}
								v5 = p424
								p426 := v5
								if uint32(v27) < uint32(v5) {
									p426 = t425
								}
								v5 = p426
								p427 := v5
								if uint32(v3) < uint32(v5) {
									p427 = t423
								}
								v5 = p427
								m.fn1736(t414, t415, t421, t422+v5)
								goto l198
							}
							t405 := int32(load32(m.memory[int64(uint32(v8))+88:]))
							t406 := v20
							t407 := v21
							t408 := v3
							v5 = t405
							p409 := v5
							if uint32(v3) < uint32(v5) {
								p409 = t408
							}
							v5 = p409
							m.fn1734(t406, t407, v5)
							goto l198
						}
					l198:
						t429 := int32(load32(m.memory[int64(uint32(v8))+88:]))
						t430 := v8
						v3 = t429
						store32(m.memory[int64(uint32(t430))+88:], uint32(v3-v5))
						if v3 == v5 {
							goto l18
						}
						goto l201
					}
					v5 = i32(0)
					v21 = i32(23)
					goto l30
				}
			case 24:
				goto l25
			case 25:
				t370 := int32(load32(m.memory[int64(uint32(v8))+36:]))
				v27 = t370
				p371 := i32(19)
				if uint32(v27) > uint32(i32(19)) {
					p371 = v27
				}
				v29 = p371
				t372 := int32(load32(m.memory[int64(uint32(v8))+24:]))
				t373 := v27
				v9 = t372
				p374 := v9
				if uint32(v27) > uint32(v9) {
					p374 = t373
				}
				v26 = p374
				t375 := int32(load32(m.memory[int64(uint32(v8))+56:]))
				v9 = t375
				t376 := int64(load64(m.memory[int64(uint32(v8))+48:]))
				v22 = t376
				t377 := int32(m.memory[int64(uint32(v8))+64])
				v5 = t377
				t378 := int32(load32(m.memory[int64(uint32(v8))+60:]))
				v25 = t378
			l191:
				if v27 == v26 {
				l188:
					{
						if uint32(v26) >= uint32(i32(19)) {
							m.fn1730(v7+i32(160), i32(0), v19, i32(19), v15, i32(1332), i32(7), v17)
							{
								t384 := int32(load32(m.memory[int64(uint32(v7))+160:]))
								if t384 != 0 {
									store32(m.memory[int64(uint32(v8))+136:], uint32(i32(25)))
									store32(m.memory[int64(uint32(v8))+132:], uint32(i32(1052963)))
									goto l58
								}
								t385 := int32(load32(m.memory[int64(uint32(v7))+164:]))
								v9 = t385
								t386 := int32(load32(m.memory[int64(uint32(v7))+168:]))
								v5 = t386
								m.memory[int64(uint32(v8))+152] = byte(i32(1))
								store32(m.memory[int64(uint32(v8))+40:], uint32(v5))
								store32(m.memory[int64(uint32(v8))+148:], uint32(v9))
								store32(m.memory[int64(uint32(v8))+36:], uint32(i32(0)))
								v9 = i32(26)
								goto l18
							}
						}
						t382 := v8
						v9 = v26 + i32(1)
						store32(m.memory[int64(uint32(t382))+36:], uint32(v9))
						t383 := int32(m.memory[uint32(v26+i32(1294892))])
						store16(m.memory[uint32(v19+t383<<1):], uint16(i32(0)))
						v26 = v9
						goto l188
					}
				}
				v3 = v5 | i32(8)
			l186:
				if uint32(v5&i32(255)) > uint32(i32(2)) {
					if v27 == v29 {
						m.fn158(v29, i32(19), i32(1294912))
						panic("unreachable")
					}
					t387 := v8
					v5 = v5 + i32(-3)
					m.memory[int64(uint32(t387))+64] = byte(v5)
					t388 := v8
					v23 = int64(uint64(v22) >> 3)
					store64(m.memory[int64(uint32(t388))+48:], uint64(v23))
					t389 := v8
					v3 = v27 + i32(1)
					store32(m.memory[int64(uint32(t389))+36:], uint32(v3))
					t390 := int32(m.memory[int64(uint32(v27))+1294892])
					store16(m.memory[uint32(v19+t390<<1):], uint16(int32(v22)&i32(7)))
					v22 = v23
					v27 = v3
					goto l191
				}
				{
					if v9 == v25 {
						v5 = i32(0)
						v21 = i32(25)
						goto l30
					}
					m.memory[int64(uint32(v8))+64] = byte(v3)
					t379 := v8
					v21 = v9 + i32(1)
					store32(m.memory[int64(uint32(t379))+56:], uint32(v21))
					t380 := int64(m.memory[uint32(v9)])
					t381 := v8
					v22 = i64_shl(t380, int64(uint32(v5))&i64(255)) | v22
					store64(m.memory[int64(uint32(t381))+48:], uint64(v22))
					v9 = v21
					v5 = v3
					goto l186
				}
			case 26:
				t308 := int32(load32(m.memory[int64(uint32(v8))+36:]))
				v26 = t308
			l182:
				{
					t309 := int32(load32(m.memory[int64(uint32(v8))+32:]))
					t310 := int32(load32(m.memory[int64(uint32(v8))+28:]))
					t311 := v26
					v9 = t310
					v28 = t309 + v9
					if uint32(t311) >= uint32(v28) {
						t317 := int32(load16(m.memory[int64(uint32(v8))+13700:]))
						if t317 != 0 {
							if uint32(v9) >= uint32(i32(321)) {
								m.fn151(i32(0), v9, i32(320), i32(1294928))
								panic("unreachable")
							}
							m.fn1730(v7+i32(160), i32(1), v19, v9, v18, i32(1332), i32(10), v17)
							t318 := int32(load32(m.memory[int64(uint32(v7))+160:]))
							if t318 != 0 {
								store32(m.memory[int64(uint32(v8))+136:], uint32(i32(28)))
								store32(m.memory[int64(uint32(v8))+132:], uint32(i32(1052935)))
								goto l58
							}
							t319 := int32(load32(m.memory[int64(uint32(v7))+168:]))
							v9 = t319
							t320 := int32(load32(m.memory[int64(uint32(v7))+164:]))
							store32(m.memory[int64(uint32(v8))+148:], uint32(t320))
							m.memory[int64(uint32(v8))+152] = byte(i32(2))
							store32(m.memory[int64(uint32(v8))+40:], uint32(v9))
							t321 := int32(load32(m.memory[int64(uint32(v8))+28:]))
							v9 = t321
							if uint32(v9) >= uint32(i32(321)) {
								m.fn151(v9, i32(320), i32(320), i32(1294944))
								panic("unreachable")
							}
							t322 := int32(load32(m.memory[int64(uint32(v8))+32:]))
							v5 = t322
							t323 := v5
							v3 = i32(320) - v9
							if uint32(t323) > uint32(v3) {
								m.fn151(i32(0), v5, v3, i32(1294960))
								panic("unreachable")
							}
							m.fn1730(v7+i32(160), i32(2), v19+v9<<1, v5, v16, i32(592), i32(9), v17)
							{
								t324 := int32(load32(m.memory[int64(uint32(v7))+160:]))
								if t324 != 0 {
									store32(m.memory[int64(uint32(v8))+136:], uint32(i32(22)))
									store32(m.memory[int64(uint32(v8))+132:], uint32(i32(0x101155)))
									goto l58
								}
								t325 := int32(load32(m.memory[int64(uint32(v7))+168:]))
								v9 = t325
								t326 := int32(load32(m.memory[int64(uint32(v7))+164:]))
								v5 = t326
								m.memory[int64(uint32(v8))+160] = byte(i32(3))
								store32(m.memory[int64(uint32(v8))+156:], uint32(v5))
								t327 := int32(load32(m.memory[int64(uint32(v8))+40:]))
								store32(m.memory[int64(uint32(v8))+40:], uint32(v9+t327))
								v9 = i32(17)
								t328 := int32(m.memory[int64(uint32(v8))+4])
								if t328 != i32(6) {
									goto l18
								}
								v5 = i32(0)
								v21 = i32(17)
								goto l30
							}
						}
						store32(m.memory[int64(uint32(v8))+136:], uint32(i32(37)))
						store32(m.memory[int64(uint32(v8))+132:], uint32(i32(1067386)))
						goto l58
					}
					t312 := int64(load32(m.memory[int64(uint32(v8))+148:]))
					v23 = i64_shl(i64(-1), t312) ^ i64(-1)
					t313 := int32(load32(m.memory[int64(uint32(v8))+56:]))
					v5 = t313
					t314 := int32(m.memory[int64(uint32(v8))+64])
					v9 = t314
					t315 := int32(load32(m.memory[int64(uint32(v8))+60:]))
					v25 = t315
					t316 := int64(load64(m.memory[int64(uint32(v8))+48:]))
					v22 = t316
				l160:
					{
						m.fn1731(v7+i32(144), v8)
						t329 := int32(load32(m.memory[int64(uint32(v7))+148:]))
						v21 = t329
						t330 := v21
						v3 = int32(v22 & v23)
						if uint32(t330) <= uint32(v3) {
							m.fn158(v3, v21, i32(1292468))
							panic("unreachable")
						}
						{
							t331 := int32(load32(m.memory[int64(uint32(v7))+144:]))
							t332 := int32(load32(m.memory[uint32(t331+v3<<2):]))
							t333 := v9 & i32(255)
							v21 = t332
							v3 = int32(uint32(v21) >> 24)
							if uint32(t333) >= uint32(v3) {
								goto l158
							}
							if v5 == v25 {
								goto l159
							}
							t334 := v8
							v3 = v5 + i32(1)
							store32(m.memory[int64(uint32(t334))+56:], uint32(v3))
							t335 := v8
							v21 = v9 + i32(8)
							m.memory[int64(uint32(t335))+64] = byte(v21)
							t336 := int64(m.memory[uint32(v5)])
							t337 := v8
							v22 = i64_shl(t336, int64(uint32(v9))) | v22
							store64(m.memory[int64(uint32(t337))+48:], uint64(v22))
							v5 = v3
							v9 = v21
							goto l160
						}
					l158:
					}
					v27 = v21 & i32(0xffff)
					if uint32(v27) < uint32(i32(16)) {
						m.memory[int64(uint32(v8))+64] = byte(v9 - v3)
						store64(m.memory[int64(uint32(v8))+48:], uint64(i64_shr_u(v22, int64(uint32(v3)))))
						if uint32(v26) < uint32(i32(320)) {
							goto l181
						}
						m.fn158(v26, i32(320), i32(1295088))
						panic("unreachable")
					}
					switch v27 + i32(-16) {
					default:
						v29 = v3 + i32(7)
					l166:
						{
							if uint32(v29) <= uint32(v9&i32(255)) {
								m.memory[int64(uint32(v8))+64] = byte(v9 - v3 + i32(-7))
								t364 := v8
								v22 = i64_shr_u(v22, int64(uint32(v3)))
								store64(m.memory[int64(uint32(t364))+48:], uint64(int64(uint64(v22)>>7)))
								{
									v9 = int32(v22)&i32(127) + i32(11)
									if uint32(v9+v26) > uint32(v28) {
										store32(m.memory[int64(uint32(v8))+136:], uint32(i32(26)))
										store32(m.memory[int64(uint32(v8))+132:], uint32(i32(1053136)))
										goto l58
									}
									m.fn1732(v7+i32(104), v19, v26, i32(1295056))
									t365 := int32(load32(m.memory[int64(uint32(v7))+104:]))
									t366 := int32(load32(m.memory[int64(uint32(v7))+108:]))
									m.fn1733(v7+i32(96), v9, t365, t366, i32(1295072))
									{
										t367 := int32(load32(m.memory[int64(uint32(v7))+100:]))
										v5 = t367 << 1
										if v5 == 0 {
											goto l180
										}
										t368 := int32(load32(m.memory[int64(uint32(v7))+96:]))
										memory_zero(m.memory, uint32(t368), uint32(v5))
									}
								l180:
									t369 := int32(load32(m.memory[int64(uint32(v8))+36:]))
									v26 = t369 + v9
									goto l176
								}
							}
							if v5 == v25 {
								goto l159
							}
							t338 := v8
							v21 = v5 + i32(1)
							store32(m.memory[int64(uint32(t338))+56:], uint32(v21))
							t339 := v8
							v27 = v9 + i32(8)
							m.memory[int64(uint32(t339))+64] = byte(v27)
							t340 := int64(m.memory[uint32(v5)])
							t341 := v8
							v22 = i64_shl(t340, int64(uint32(v9))) | v22
							store64(m.memory[int64(uint32(t341))+48:], uint64(v22))
							v5 = v21
							v9 = v27
							goto l166
						}
					case 0:
						v29 = v3 + i32(2)
					l168:
						{
							if uint32(v29) <= uint32(v9&i32(255)) {
								t350 := v8
								v9 = v9 - v3
								m.memory[int64(uint32(t350))+64] = byte(v9)
								t351 := v8
								v22 = i64_shr_u(v22, int64(uint32(v3)))
								store64(m.memory[int64(uint32(t351))+48:], uint64(v22))
								if v26 != 0 {
									v5 = v26 + i32(-1)
									if uint32(v26) > uint32(i32(320)) {
										m.fn158(v5, i32(320), i32(1294976))
										panic("unreachable")
									}
									t352 := int32(load16(m.memory[uint32(v19+v5<<1):]))
									v3 = t352
									m.memory[int64(uint32(v8))+64] = byte(v9 + i32(-2))
									store64(m.memory[int64(uint32(v8))+48:], uint64(int64(uint64(v22)>>2)))
									{
										v21 = int32(v22)&i32(3) + i32(3)
										if uint32(v21+v26) > uint32(v28) {
											store32(m.memory[int64(uint32(v8))+136:], uint32(i32(26)))
											store32(m.memory[int64(uint32(v8))+132:], uint32(i32(1053136)))
											goto l58
										}
										m.fn1732(v7+i32(120), v19, v26, i32(1294992))
										t353 := int32(load32(m.memory[int64(uint32(v7))+120:]))
										t354 := int32(load32(m.memory[int64(uint32(v7))+124:]))
										m.fn1733(v7+i32(112), v21, t353, t354, i32(1295008))
										t355 := int32(load32(m.memory[int64(uint32(v7))+116:]))
										v9 = t355 << 1
										t356 := int32(load32(m.memory[int64(uint32(v7))+112:]))
										v5 = t356
									l175:
										if v9 == 0 {
											t357 := int32(load32(m.memory[int64(uint32(v8))+36:]))
											v26 = t357 + v21
											goto l176
										}
										store16(m.memory[uint32(v5):], uint16(v3))
										v9 = v9 + i32(-2)
										v5 = v5 + i32(2)
										goto l175
									}
								}
								store32(m.memory[int64(uint32(v8))+136:], uint32(i32(26)))
								store32(m.memory[int64(uint32(v8))+132:], uint32(i32(1053136)))
								goto l58
							}
							if v5 == v25 {
								goto l159
							}
							t342 := v8
							v21 = v5 + i32(1)
							store32(m.memory[int64(uint32(t342))+56:], uint32(v21))
							t343 := v8
							v27 = v9 + i32(8)
							m.memory[int64(uint32(t343))+64] = byte(v27)
							t344 := int64(m.memory[uint32(v5)])
							t345 := v8
							v22 = i64_shl(t344, int64(uint32(v9))) | v22
							store64(m.memory[int64(uint32(t345))+48:], uint64(v22))
							v5 = v21
							v9 = v27
							goto l168
						}
					case 1:
						v29 = v3 + i32(3)
					l170:
						{
							if uint32(v29) <= uint32(v9&i32(255)) {
								v5 = i32(-3)
								m.memory[int64(uint32(v8))+64] = byte(v9 - v3 + i32(-3))
								t358 := v8
								v22 = i64_shr_u(v22, int64(uint32(v3)))
								store64(m.memory[int64(uint32(t358))+48:], uint64(int64(uint64(v22)>>3)))
								{
									v9 = int32(v22)&i32(7) + i32(3)
									if uint32(v9+v26) > uint32(v28) {
										store32(m.memory[int64(uint32(v8))+136:], uint32(i32(26)))
										store32(m.memory[int64(uint32(v8))+132:], uint32(i32(1053136)))
										v21 = i32(30)
										goto l30
									}
									m.fn1732(v7+i32(136), v19, v26, i32(1295024))
									t359 := int32(load32(m.memory[int64(uint32(v7))+136:]))
									t360 := int32(load32(m.memory[int64(uint32(v7))+140:]))
									m.fn1733(v7+i32(128), v9, t359, t360, i32(1295040))
									{
										t361 := int32(load32(m.memory[int64(uint32(v7))+132:]))
										v5 = t361 << 1
										if v5 == 0 {
											goto l178
										}
										t362 := int32(load32(m.memory[int64(uint32(v7))+128:]))
										memory_zero(m.memory, uint32(t362), uint32(v5))
									}
								l178:
									t363 := int32(load32(m.memory[int64(uint32(v8))+36:]))
									v26 = t363 + v9
									goto l176
								}
							}
							if v5 == v25 {
								goto l159
							}
							t346 := v8
							v21 = v5 + i32(1)
							store32(m.memory[int64(uint32(t346))+56:], uint32(v21))
							t347 := v8
							v27 = v9 + i32(8)
							m.memory[int64(uint32(t347))+64] = byte(v27)
							t348 := int64(m.memory[uint32(v5)])
							t349 := v8
							v22 = i64_shl(t348, int64(uint32(v9))) | v22
							store64(m.memory[int64(uint32(t349))+48:], uint64(v22))
							v5 = v21
							v9 = v27
							goto l170
						}
					}
				}
			l181:
				store16(m.memory[uint32(v19+v26<<1):], uint16(v21))
				v26 = v26 + i32(1)
			l176:
				store32(m.memory[int64(uint32(v8))+36:], uint32(v26))
				goto l182
			l159:
				v5 = i32(0)
				v21 = i32(26)
				goto l30
			case 27:
				t299 := int64(m.memory[int64(uint32(v8))+64])
				v22 = t299
				t300 := int32(load32(m.memory[int64(uint32(v8))+56:]))
				v9 = t300
				t301 := int32(load32(m.memory[int64(uint32(v8))+60:]))
				v3 = t301
				t302 := int64(load64(m.memory[int64(uint32(v8))+48:]))
				v23 = t302
			l148:
				if uint64(v22) > uint64(i64(31)) {
					m.memory[int64(uint32(v8))+64] = byte(i32(0))
					store64(m.memory[int64(uint32(v8))+48:], uint64(i64(0)))
					t307 := v8
					v9 = int32(v23)
					store32(m.memory[int64(uint32(t307))+124:], uint32(i32_rotr(v9&i32(0xff00ff), i32(8))|i32_rotr(v9, i32(24))&i32(0xff00ff)))
					v9 = i32(28)
					goto l18
				}
				if v9 != v3 {
					t303 := v8
					v5 = v9 + i32(1)
					store32(m.memory[int64(uint32(t303))+56:], uint32(v5))
					t304 := v8
					v24 = v22 + i64(8)
					m.memory[int64(uint32(t304))+64] = byte(v24)
					t305 := int64(m.memory[uint32(v9)])
					t306 := v8
					v23 = i64_shl(t305, v22) | v23
					store64(m.memory[int64(uint32(t306))+48:], uint64(v23))
					v22 = v24
					v9 = v5
					goto l148
				}
				v5 = i32(0)
				v21 = i32(27)
				goto l30
			case 28:
				v5 = i32(2)
				{
					t298 := int32(m.memory[int64(uint32(v8))+1])
					if t298&i32(2) != 0 {
						store32(m.memory[int64(uint32(v8))+124:], uint32(i32(1)))
						v9 = i32(12)
						goto l18
					}
					v21 = i32(28)
					goto l30
				}
			case 29:
				goto l30
			case 30:
				store32(m.memory[int64(uint32(v8))+136:], uint32(i32(29)))
				store32(m.memory[int64(uint32(v8))+132:], uint32(i32(1067756)))
				goto l58
			default:
				v9 = i32(13)
				t279 := int32(m.memory[int64(uint32(v8))+3])
				v21 = t279
				if v21 == 0 {
					goto l18
				}
				t280 := int64(m.memory[int64(uint32(v8))+64])
				v22 = t280
				t281 := int32(load32(m.memory[int64(uint32(v8))+56:]))
				v9 = t281
				t282 := int32(load32(m.memory[int64(uint32(v8))+60:]))
				v3 = t282
				t283 := int64(load64(m.memory[int64(uint32(v8))+48:]))
				v23 = t283
			l135:
				if uint64(v22) > uint64(i64(15)) {
					if v21&i32(2) == 0 {
						goto l136
					}
					if v23 == i64(35615) {
						{
							t292 := int32(m.memory[int64(uint32(v8))+2])
							if t292 != 0 {
								goto l141
							}
							m.memory[int64(uint32(v8))+2] = byte(i32(15))
						}
					l141:
						store16(m.memory[int64(uint32(v7))+160:], uint16(i32(35615)))
						t293 := m.fn1729(i32(0), v7+i32(160), i32(2))
						store32(m.memory[int64(uint32(v8))+124:], uint32(t293))
						m.memory[int64(uint32(v8))+64] = byte(i32(0))
						store64(m.memory[int64(uint32(v8))+48:], uint64(i64(0)))
						v9 = i32(1)
						goto l18
					}
				l136:
					{
						t288 := int32(load32(m.memory[int64(uint32(v8))+140:]))
						v9 = t288
						if v9 == 0 {
							goto l138
						}
						store32(m.memory[int64(uint32(v9))+48:], uint32(i32(-1)))
						t289 := int32(m.memory[int64(uint32(v8))+3])
						v21 = t289
					}
				l138:
					{
						if v21&i32(1) == 0 {
							goto l139
						}
						t290 := int64(load64(m.memory[int64(uint32(v8))+48:]))
						v22 = t290
						t291 := int64(uint64(v22<<8&i64(0xff00)+int64(uint64(v22)>>8)) % uint64(i64(31)))
						if t291 == 0 {
							if v22&i64(15) != i64(8) {
								store32(m.memory[int64(uint32(v8))+136:], uint32(i32(27)))
								store32(m.memory[int64(uint32(v8))+132:], uint32(i32(1068085)))
								goto l58
							}
							t294 := v8
							v23 = int64(uint64(v22) >> 4)
							store64(m.memory[int64(uint32(t294))+48:], uint64(v23))
							t295 := int32(m.memory[int64(uint32(v8))+64])
							m.memory[int64(uint32(v8))+64] = byte(t295 + i32(-4))
							v3 = int32(v23) & i32(15)
							v9 = v3 + i32(8)
							{
								t296 := int32(m.memory[int64(uint32(v8))+2])
								v5 = t296
								if v5 != 0 {
									goto l143
								}
								m.memory[int64(uint32(v8))+2] = byte(v9)
								v5 = v9
							}
						l143:
							{
								if uint32(v3) > uint32(i32(7)) {
									goto l144
								}
								if uint32(v9) > uint32(v5) {
									goto l144
								}
								store64(m.memory[int64(uint32(v8))+120:], uint64(i64(0x100000000)))
								m.memory[int64(uint32(v8))+64] = byte(i32(0))
								store64(m.memory[int64(uint32(v8))+48:], uint64(i64(0)))
								store32(m.memory[int64(uint32(v8))+144:], uint32(i32_shl(i32(1), v9)))
								p297 := i32(27)
								if v22&i64(8192) == 0 {
									p297 = i32(12)
								}
								v9 = p297
								goto l18
							}
						l144:
							store32(m.memory[int64(uint32(v8))+136:], uint32(i32(20)))
							store32(m.memory[int64(uint32(v8))+132:], uint32(i32(1067587)))
							goto l58
						}
					}
				l139:
					store32(m.memory[int64(uint32(v8))+136:], uint32(i32(23)))
					store32(m.memory[int64(uint32(v8))+132:], uint32(i32(1067423)))
					goto l58
				}
				if v9 != v3 {
					t284 := v8
					v5 = v9 + i32(1)
					store32(m.memory[int64(uint32(t284))+56:], uint32(v5))
					t285 := v8
					v24 = v22 + i64(8)
					m.memory[int64(uint32(t285))+64] = byte(v24)
					t286 := int64(m.memory[uint32(v9)])
					t287 := v8
					v23 = i64_shl(t286, v22) | v23
					store64(m.memory[int64(uint32(t287))+48:], uint64(v23))
					v22 = v24
					v9 = v5
					goto l135
				}
				v21 = i32(0)
				goto l93
			case 1:
				t5 := int64(m.memory[int64(uint32(v8))+64])
				v22 = t5
				t6 := int32(load32(m.memory[int64(uint32(v8))+56:]))
				v9 = t6
				t7 := int32(load32(m.memory[int64(uint32(v8))+60:]))
				v3 = t7
				t8 := int64(load64(m.memory[int64(uint32(v8))+48:]))
				v23 = t8
			l34:
				if uint64(v22) > uint64(i64(15)) {
					t272 := v8
					v9 = int32(v23)
					store32(m.memory[int64(uint32(t272))+120:], uint32(v9))
					{
						if v9&i32(255) != i32(8) {
							store32(m.memory[int64(uint32(v8))+136:], uint32(i32(27)))
							store32(m.memory[int64(uint32(v8))+132:], uint32(i32(1068085)))
							goto l58
						}
						if v9&i32(57344) != 0 {
							store32(m.memory[int64(uint32(v8))+136:], uint32(i32(25)))
							store32(m.memory[int64(uint32(v8))+132:], uint32(i32(1052988)))
							goto l58
						}
						{
							t273 := int32(load32(m.memory[int64(uint32(v8))+140:]))
							v5 = t273
							if v5 == 0 {
								goto l131
							}
							store32(m.memory[uint32(v5):], uint32(int32(uint32(v9)>>8)&i32(1)))
							t274 := int32(load32(m.memory[int64(uint32(v8))+120:]))
							v9 = t274
						}
					l131:
						if v9&i32(512) == 0 {
							goto l132
						}
						t275 := int32(m.memory[int64(uint32(v8))+3])
						if t275&i32(4) == 0 {
							goto l132
						}
						t276 := int32(load32(m.memory[int64(uint32(v8))+124:]))
						v9 = t276
						t277 := int64(load64(m.memory[int64(uint32(v8))+48:]))
						store16(m.memory[int64(uint32(v7))+160:], uint16(t277))
						t278 := m.fn1729(v9, v7+i32(160), i32(2))
						store32(m.memory[int64(uint32(v8))+124:], uint32(t278))
						goto l132
					}
				l132:
					m.memory[int64(uint32(v8))+64] = byte(i32(0))
					store64(m.memory[int64(uint32(v8))+48:], uint64(i64(0)))
					v9 = i32(2)
					goto l18
				}
				if v9 != v3 {
					t9 := v8
					v5 = v9 + i32(1)
					store32(m.memory[int64(uint32(t9))+56:], uint32(v5))
					t10 := v8
					v24 = v22 + i64(8)
					m.memory[int64(uint32(t10))+64] = byte(v24)
					t11 := int64(m.memory[uint32(v9)])
					t12 := v8
					v23 = i64_shl(t11, v22) | v23
					store64(m.memory[int64(uint32(t12))+48:], uint64(v23))
					v22 = v24
					v9 = v5
					goto l34
				}
				v5 = i32(0)
				v21 = i32(1)
				goto l30
			case 2:
				t13 := int64(m.memory[int64(uint32(v8))+64])
				v22 = t13
				t14 := int32(load32(m.memory[int64(uint32(v8))+56:]))
				v9 = t14
				t15 := int32(load32(m.memory[int64(uint32(v8))+60:]))
				v3 = t15
				t16 := int64(load64(m.memory[int64(uint32(v8))+48:]))
				v23 = t16
			l37:
				if uint64(v22) > uint64(i64(31)) {
					{
						t266 := int32(load32(m.memory[int64(uint32(v8))+140:]))
						v9 = t266
						if v9 == 0 {
							goto l127
						}
						store32(m.memory[int64(uint32(v9))+4:], uint32(v23))
					}
				l127:
					{
						t267 := int32(m.memory[int64(uint32(v8))+121])
						if t267&i32(2) == 0 {
							goto l128
						}
						t268 := int32(m.memory[int64(uint32(v8))+3])
						if t268&i32(4) == 0 {
							goto l128
						}
						t269 := int64(load64(m.memory[int64(uint32(v8))+48:]))
						store32(m.memory[int64(uint32(v7))+160:], uint32(t269))
						t270 := int32(load32(m.memory[int64(uint32(v8))+124:]))
						t271 := m.fn1729(t270, v7+i32(160), i32(4))
						store32(m.memory[int64(uint32(v8))+124:], uint32(t271))
					}
				l128:
					m.memory[int64(uint32(v8))+64] = byte(i32(0))
					store64(m.memory[int64(uint32(v8))+48:], uint64(i64(0)))
					v9 = i32(3)
					goto l18
				}
				if v9 != v3 {
					t17 := v8
					v5 = v9 + i32(1)
					store32(m.memory[int64(uint32(t17))+56:], uint32(v5))
					t18 := v8
					v24 = v22 + i64(8)
					m.memory[int64(uint32(t18))+64] = byte(v24)
					t19 := int64(m.memory[uint32(v9)])
					t20 := v8
					v23 = i64_shl(t19, v22) | v23
					store64(m.memory[int64(uint32(t20))+48:], uint64(v23))
					v22 = v24
					v9 = v5
					goto l37
				}
				v5 = i32(0)
				v21 = i32(2)
				goto l30
			case 3:
				t21 := int64(m.memory[int64(uint32(v8))+64])
				v22 = t21
				t22 := int32(load32(m.memory[int64(uint32(v8))+56:]))
				v9 = t22
				t23 := int32(load32(m.memory[int64(uint32(v8))+60:]))
				v3 = t23
				t24 := int64(load64(m.memory[int64(uint32(v8))+48:]))
				v23 = t24
			l40:
				if uint64(v22) > uint64(i64(15)) {
					{
						t258 := int32(load32(m.memory[int64(uint32(v8))+140:]))
						v9 = t258
						if v9 == 0 {
							goto l125
						}
						store32(m.memory[int64(uint32(v9))+8:], uint32(int32(v23)&i32(255)))
						t259 := int32(load32(m.memory[int64(uint32(v8))+140:]))
						t260 := int64(load64(m.memory[int64(uint32(v8))+48:]))
						store32(m.memory[int64(uint32(t259))+12:], uint32(int64(uint64(t260)>>8)))
					}
				l125:
					{
						t261 := int32(m.memory[int64(uint32(v8))+121])
						if t261&i32(2) == 0 {
							goto l126
						}
						t262 := int32(m.memory[int64(uint32(v8))+3])
						if t262&i32(4) == 0 {
							goto l126
						}
						t263 := int64(load64(m.memory[int64(uint32(v8))+48:]))
						store16(m.memory[int64(uint32(v7))+160:], uint16(t263))
						t264 := int32(load32(m.memory[int64(uint32(v8))+124:]))
						t265 := m.fn1729(t264, v7+i32(160), i32(2))
						store32(m.memory[int64(uint32(v8))+124:], uint32(t265))
					}
				l126:
					m.memory[int64(uint32(v8))+64] = byte(i32(0))
					store64(m.memory[int64(uint32(v8))+48:], uint64(i64(0)))
					v9 = i32(4)
					goto l18
				}
				if v9 != v3 {
					t25 := v8
					v5 = v9 + i32(1)
					store32(m.memory[int64(uint32(t25))+56:], uint32(v5))
					t26 := v8
					v24 = v22 + i64(8)
					m.memory[int64(uint32(t26))+64] = byte(v24)
					t27 := int64(m.memory[uint32(v9)])
					t28 := v8
					v23 = i64_shl(t27, v22) | v23
					store64(m.memory[int64(uint32(t28))+48:], uint64(v23))
					v22 = v24
					v9 = v5
					goto l40
				}
				v5 = i32(0)
				v21 = i32(3)
				goto l30
			case 11:
				t29 := int32(m.memory[int64(uint32(v8))+3])
				v25 = t29
				if v25 == 0 {
					goto l30
				}
				t30 := int32(load32(m.memory[int64(uint32(v8))+120:]))
				if t30 == 0 {
					goto l30
				}
				t31 := int64(m.memory[int64(uint32(v8))+64])
				v22 = t31
				t32 := int32(load32(m.memory[int64(uint32(v8))+56:]))
				v9 = t32
				t33 := int32(load32(m.memory[int64(uint32(v8))+60:]))
				v3 = t33
				t34 := int64(load64(m.memory[int64(uint32(v8))+48:]))
				v23 = t34
			l82:
				if uint64(v22) > uint64(i64(31)) {
					{
						if v25&i32(4) == 0 {
							goto l83
						}
						t140 := int32(load32(m.memory[int64(uint32(v8))+84:]))
						if t140 != int32(v23) {
							store32(m.memory[int64(uint32(v8))+136:], uint32(i32(23)))
							store32(m.memory[int64(uint32(v8))+132:], uint32(i32(1067446)))
							goto l58
						}
					}
				l83:
					m.memory[int64(uint32(v8))+64] = byte(i32(0))
					store64(m.memory[int64(uint32(v8))+48:], uint64(i64(0)))
					v5 = i32(1)
					goto l30
				}
				{
					if v9 == v3 {
						v5 = i32(0)
						v21 = i32(11)
						goto l30
					}
					t136 := v8
					v5 = v9 + i32(1)
					store32(m.memory[int64(uint32(t136))+56:], uint32(v5))
					t137 := v8
					v24 = v22 + i64(8)
					m.memory[int64(uint32(t137))+64] = byte(v24)
					t138 := int64(m.memory[uint32(v9)])
					t139 := v8
					v23 = i64_shl(t138, v22) | v23
					store64(m.memory[int64(uint32(t139))+48:], uint64(v23))
					v22 = v24
					v9 = v5
					goto l82
				}
			case 12:
				v9 = i32(13)
				t35 := int32(m.memory[int64(uint32(v8))+4])
				if uint32(t35) <= uint32(i32(4)) {
					goto l18
				}
				v5 = i32(0)
				v21 = i32(12)
				goto l30
			case 14:
				t36 := int64(load64(m.memory[int64(uint32(v8))+48:]))
				t37 := int32(m.memory[int64(uint32(v8))+64])
				t38 := v8
				v9 = t37
				v23 = i64_shr_u(t36, int64(uint32(v9&i32(7))))
				store64(m.memory[int64(uint32(t38))+48:], uint64(v23))
				v22 = int64(uint32(v9)) & i64(248)
				t39 := int32(load32(m.memory[int64(uint32(v8))+56:]))
				v9 = t39
				t40 := int32(load32(m.memory[int64(uint32(v8))+60:]))
				v3 = t40
			l45:
				m.memory[int64(uint32(v8))+64] = byte(v22)
				if uint64(v22) > uint64(i64(31)) {
					goto l42
				}
				if v9 != v3 {
					t41 := v8
					v5 = v9 + i32(1)
					store32(m.memory[int64(uint32(t41))+56:], uint32(v5))
					t42 := int64(m.memory[uint32(v9)])
					t43 := v8
					v23 = i64_shl(t42, v22) | v23
					store64(m.memory[int64(uint32(t43))+48:], uint64(v23))
					v22 = v22 + i64(8)
					v9 = v5
					goto l45
				}
				v5 = i32(0)
				goto l44
			case 15:
				t44 := int32(load32(m.memory[int64(uint32(v8))+88:]))
				v5 = t44
			l47:
				v9 = i32(12)
				if v5 == 0 {
					goto l18
				}
				{
					t45 := int32(load32(m.memory[int64(uint32(v8))+60:]))
					t46 := int32(load32(m.memory[int64(uint32(v8))+56:]))
					v21 = t46
					v9 = t45 - v21
					t47 := int32(load32(m.memory[int64(uint32(v8))+76:]))
					t48 := v9
					v25 = t47
					t49 := int32(load32(m.memory[int64(uint32(v8))+80:]))
					t50 := v25
					v26 = t49
					v3 = t50 - v26
					p51 := v5
					if uint32(v3) < uint32(v5) {
						p51 = v3
					}
					v5 = p51
					p52 := v5
					if uint32(v9) < uint32(v5) {
						p52 = t48
					}
					v9 = p52
					if v9 != 0 {
						t53 := int32(load32(m.memory[int64(uint32(v8))+72:]))
						m.fn1723(v7+i32(88), v26, t53, v25, i32(1292312))
						t54 := int32(load32(m.memory[int64(uint32(v7))+88:]))
						t55 := int32(load32(m.memory[int64(uint32(v7))+92:]))
						m.fn1724(v7+i32(80), v9, t54, t55, i32(1292328))
						t56 := int32(load32(m.memory[int64(uint32(v7))+80:]))
						t57 := int32(load32(m.memory[int64(uint32(v7))+84:]))
						m.fn310(t56, t57, v21, v9, i32(1292344))
						t58 := int32(load32(m.memory[int64(uint32(v8))+80:]))
						store32(m.memory[int64(uint32(v8))+80:], uint32(t58+v9))
						t59 := int32(load32(m.memory[int64(uint32(v8))+88:]))
						t60 := v8
						v5 = t59 - v9
						store32(m.memory[int64(uint32(t60))+88:], uint32(v5))
						t61 := int32(load32(m.memory[int64(uint32(v8))+60:]))
						t62 := v8
						v3 = t61
						t63 := int32(load32(m.memory[int64(uint32(v8))+56:]))
						t64 := v3
						v9 = t63 + v9
						p65 := v9
						if uint32(v3) < uint32(v9) {
							p65 = t64
						}
						store32(m.memory[int64(uint32(t62))+56:], uint32(p65))
						goto l47
					}
					v5 = i32(0)
					v21 = i32(15)
					goto l30
				}
			case 16:
				v9 = i32(11)
				t66 := int32(m.memory[int64(uint32(v8))+3])
				v25 = t66
				if v25 == 0 {
					goto l18
				}
				t67 := int64(m.memory[int64(uint32(v8))+64])
				v22 = t67
				t68 := int32(load32(m.memory[int64(uint32(v8))+56:]))
				v5 = t68
				t69 := int32(load32(m.memory[int64(uint32(v8))+60:]))
				v21 = t69
				t70 := int64(load64(m.memory[int64(uint32(v8))+48:]))
				v23 = t70
			l50:
				if uint64(v22) > uint64(i64(31)) {
					t75 := int32(load32(m.memory[int64(uint32(v8))+84:]))
					t76 := int32(load32(m.memory[int64(uint32(v8))+80:]))
					t77 := v8
					v3 = t76
					store32(m.memory[int64(uint32(t77))+84:], uint32(t75+v3))
					t78 := int32(load32(m.memory[int64(uint32(v8))+120:]))
					v5 = t78
					{
						if v25&i32(4) == 0 {
							goto l51
						}
						{
							if v5 != 0 {
								goto l52
							}
							t79 := int32(load32(m.memory[int64(uint32(v8))+124:]))
							t80 := int32(load32(m.memory[int64(uint32(v8))+72:]))
							t81 := m.fn1725(t79, t80, v3)
							store32(m.memory[int64(uint32(v8))+124:], uint32(t81))
							goto l53
						}
					l52:
						t82 := int32(load32(m.memory[int64(uint32(v8))+72:]))
						m.fn1726(v12, t82, v3)
						t83 := int32(load32(m.memory[int64(uint32(v8))+128:]))
						store32(m.memory[int64(uint32(v8))+124:], uint32(t83))
						t84 := int64(load64(m.memory[int64(uint32(v8))+48:]))
						v23 = t84
						t85 := int32(load32(m.memory[int64(uint32(v8))+120:]))
						v5 = t85
					}
				l51:
					if v5 != 0 {
						goto l54
					}
				l53:
					v5 = int32(v23)
					v5 = i32_rotr(v5&i32(0xff00ff), i32(8)) | i32_rotr(v5, i32(24))&i32(0xff00ff)
					goto l55
				l54:
					v5 = int32(v23)
				l55:
					t86 := int32(load32(m.memory[int64(uint32(v8))+76:]))
					t87 := int32(load32(m.memory[int64(uint32(v8))+80:]))
					store32(m.memory[int64(uint32(v8))+116:], uint32(t86-t87))
					{
						t88 := int32(m.memory[int64(uint32(v8))+3])
						if t88&i32(4) == 0 {
							goto l56
						}
						t89 := int32(load32(m.memory[int64(uint32(v8))+124:]))
						if v5 != t89 {
							store32(m.memory[int64(uint32(v8))+136:], uint32(i32(21)))
							store32(m.memory[int64(uint32(v8))+132:], uint32(i32(1067469)))
							goto l58
						}
					}
				l56:
					m.memory[int64(uint32(v8))+64] = byte(i32(0))
					store64(m.memory[int64(uint32(v8))+48:], uint64(i64(0)))
					goto l18
				}
				if v5 != v21 {
					t71 := v8
					v3 = v5 + i32(1)
					store32(m.memory[int64(uint32(t71))+56:], uint32(v3))
					t72 := v8
					v24 = v22 + i64(8)
					m.memory[int64(uint32(t72))+64] = byte(v24)
					t73 := int64(m.memory[uint32(v5)])
					t74 := v8
					v23 = i64_shl(t73, v22) | v23
					store64(m.memory[int64(uint32(t74))+48:], uint64(v23))
					v22 = v24
					v5 = v3
					goto l50
				}
				v5 = i32(0)
				v21 = i32(16)
				goto l30
			case 21:
				t90 := int64(load64(m.memory[int64(uint32(v8))+48:]))
				v22 = t90
				{
				l64:
					{
						t91 := int64(load32(m.memory[int64(uint32(v8))+156:]))
						t92 := m.fn1727(v8, int32(v22&(i64_shl(i64(-1), t91)^i64(-1))))
						v5 = t92
						{
							t93 := int32(m.memory[int64(uint32(v8))+64])
							v9 = t93
							t94 := v9
							v3 = int32(uint32(v5) >> 24)
							if uint32(t94) >= uint32(v3) {
								t97 := int64(load64(m.memory[int64(uint32(v8))+48:]))
								v22 = t97
								v27 = int32(uint32(v5) >> 16)
								if uint32(v27&i32(255)) <= uint32(i32(15)) {
									goto l62
								}
								v21 = v3
								goto l63
							}
							t95 := int32(load32(m.memory[int64(uint32(v8))+56:]))
							v5 = t95
							t96 := int32(load32(m.memory[int64(uint32(v8))+60:]))
							if v5 != t96 {
								store32(m.memory[int64(uint32(v8))+56:], uint32(v5+i32(1)))
								m.memory[int64(uint32(v8))+64] = byte(v9 + i32(8))
								t98 := int64(load64(m.memory[int64(uint32(v8))+48:]))
								t99 := int64(m.memory[uint32(v5)])
								t100 := v8
								v22 = t98 | i64_shl(t99, int64(uint32(v9)))
								store64(m.memory[int64(uint32(t100))+48:], uint64(v22))
								goto l64
							}
							goto l61
						}
					}
				l62:
					v25 = v3 & i32(31)
					v26 = v5 & i32(0xffff)
					v23 = i64_shl(i64(-1), int64(uint32(v3+v27))) ^ i64(-1)
				l66:
					{
						t101 := m.fn1727(v8, i32_shr_u(int32(v22&v23), v25)+v26)
						v5 = t101
						v21 = int32(uint32(v5) >> 24)
						t102 := int32(m.memory[int64(uint32(v8))+64])
						t103 := (v21 + v3) & i32(255)
						v9 = t102
						if uint32(t103) <= uint32(v9) {
							goto l65
						}
						t104 := int32(load32(m.memory[int64(uint32(v8))+56:]))
						v5 = t104
						t105 := int32(load32(m.memory[int64(uint32(v8))+60:]))
						if v5 == t105 {
							goto l61
						}
						store32(m.memory[int64(uint32(v8))+56:], uint32(v5+i32(1)))
						m.memory[int64(uint32(v8))+64] = byte(v9 + i32(8))
						t106 := int64(load64(m.memory[int64(uint32(v8))+48:]))
						t107 := int64(m.memory[uint32(v5)])
						t108 := v8
						v22 = t106 | i64_shl(t107, int64(uint32(v9)))
						store64(m.memory[int64(uint32(t108))+48:], uint64(v22))
						goto l66
					}
				l65:
					t109 := int32(load32(m.memory[int64(uint32(v8))+100:]))
					store32(m.memory[int64(uint32(v8))+100:], uint32(t109+v3))
					v9 = v9 - v3
					v27 = int32(uint32(v5) >> 16)
					t110 := int64(load64(m.memory[int64(uint32(v8))+48:]))
					v22 = i64_shr_u(t110, int64(uint32(v3)))
				}
			l63:
				m.memory[int64(uint32(v8))+64] = byte(v9 - v21)
				store64(m.memory[int64(uint32(v8))+48:], uint64(i64_shr_u(v22, int64(uint32(v21)))))
				if v27&i32(64) != 0 {
					store32(m.memory[int64(uint32(v8))+136:], uint32(i32(22)))
					store32(m.memory[int64(uint32(v8))+132:], uint32(i32(1067932)))
					goto l58
				}
				store32(m.memory[int64(uint32(v8))+96:], uint32(v27&i32(15)))
				store32(m.memory[int64(uint32(v8))+92:], uint32(v5&i32(0xffff)))
				v9 = i32(22)
				goto l18
			case 22:
				v9 = i32(23)
				t111 := int32(load32(m.memory[int64(uint32(v8))+96:]))
				v26 = t111
				if v26 == 0 {
					goto l18
				}
				t112 := int32(load32(m.memory[int64(uint32(v8))+56:]))
				v3 = t112
				t113 := int32(load32(m.memory[int64(uint32(v8))+60:]))
				v27 = t113
				t114 := int64(load64(m.memory[int64(uint32(v8))+48:]))
				v22 = t114
				t115 := int32(m.memory[int64(uint32(v8))+64])
				v5 = t115
			l70:
				if uint32(v26) <= uint32(v5&i32(255)) {
					m.memory[int64(uint32(v8))+64] = byte(v5 - v26)
					t120 := v8
					t121 := v22
					v23 = int64(uint32(v26))
					store64(m.memory[int64(uint32(t120))+48:], uint64(i64_shr_u(t121, v23)))
					t122 := int32(load32(m.memory[int64(uint32(v8))+100:]))
					store32(m.memory[int64(uint32(v8))+100:], uint32(t122+v26))
					t123 := int32(load32(m.memory[int64(uint32(v8))+92:]))
					store32(m.memory[int64(uint32(v8))+92:], uint32(t123+int32(v22&(i64_shl(i64(-1), v23)^i64(-1)))))
					goto l18
				}
				if v3 != v27 {
					t116 := v8
					v21 = v3 + i32(1)
					store32(m.memory[int64(uint32(t116))+56:], uint32(v21))
					t117 := v8
					v25 = v5 + i32(8)
					m.memory[int64(uint32(t117))+64] = byte(v25)
					t118 := int64(m.memory[uint32(v3)])
					t119 := v8
					v22 = i64_shl(t118, int64(uint32(v5))) | v22
					store64(m.memory[int64(uint32(t119))+48:], uint64(v22))
					v3 = v21
					v5 = v25
					goto l70
				}
				v5 = i32(0)
				v21 = i32(22)
				goto l30
			case 9:
				v5 = i32(-2)
				v21 = i32(9)
				goto l30
			case 10:
				v5 = i32(-4)
				v21 = i32(10)
				goto l30
			}
		l61:
			v5 = i32(0)
			v21 = i32(21)
			goto l30
		l42:
			if (int64(uint64(v23)>>16)^v23)&i64(0xffff) != i64(0xffff) {
				store32(m.memory[int64(uint32(v8))+136:], uint32(i32(29)))
				store32(m.memory[int64(uint32(v8))+132:], uint32(i32(1066739)))
				goto l58
			}
			v5 = i32(0)
			m.memory[int64(uint32(v8))+64] = byte(i32(0))
			store64(m.memory[int64(uint32(v8))+48:], uint64(i64(0)))
			store32(m.memory[int64(uint32(v8))+88:], uint32(int32(v23)&i32(0xffff)))
			v9 = i32(15)
			t124 := int32(m.memory[int64(uint32(v8))+4])
			if t124 != i32(6) {
				goto l18
			}
		}
	l44:
		v21 = i32(14)
		goto l30
	l25:
		t391 := int32(m.memory[int64(uint32(v8))+64])
		v3 = t391
		v22 = int64(uint32(v3)) & i64(255)
		t392 := int32(load32(m.memory[int64(uint32(v8))+56:]))
		v9 = t392
		t393 := int32(load32(m.memory[int64(uint32(v8))+60:]))
		v21 = t393
		t394 := int64(load64(m.memory[int64(uint32(v8))+48:]))
		v24 = t394
	l194:
		if uint64(v22) > uint64(i64(13)) {
			goto l192
		}
		if v9 != v21 {
			t395 := v8
			v5 = v9 + i32(1)
			store32(m.memory[int64(uint32(t395))+56:], uint32(v5))
			t396 := v8
			v23 = v22 + i64(8)
			m.memory[int64(uint32(t396))+64] = byte(v23)
			t397 := int64(m.memory[uint32(v9)])
			t398 := v8
			v24 = i64_shl(t397, v22) | v24
			store64(m.memory[int64(uint32(t398))+48:], uint64(v24))
			v3 = int32(v23)
			v22 = v23
			v9 = v5
			goto l194
		}
		v5 = i32(0)
		v21 = i32(24)
		goto l30
	l192:
		m.memory[int64(uint32(v8))+64] = byte(v3 + i32(-14))
		store64(m.memory[int64(uint32(v8))+48:], uint64(int64(uint64(v24)>>14)))
		t399 := v8
		v9 = int32(v24)
		v5 = v9 & i32(31)
		store32(m.memory[int64(uint32(t399))+28:], uint32(v5+i32(257)))
		t400 := v8
		v3 = int32(uint32(v9)>>5) & i32(31)
		store32(m.memory[int64(uint32(t400))+32:], uint32(v3+i32(1)))
		store32(m.memory[int64(uint32(v8))+24:], uint32(int32(uint32(v9)>>10)&i32(15)+i32(4)))
		if uint32(v5) > uint32(i32(29)) {
			goto l195
		}
		if uint32(v3) > uint32(i32(29)) {
			goto l195
		}
		store32(m.memory[int64(uint32(v8))+36:], uint32(i32(0)))
		v9 = i32(25)
		goto l18
	l195:
		store32(m.memory[int64(uint32(v8))+136:], uint32(i32(36)))
		store32(m.memory[int64(uint32(v8))+132:], uint32(i32(1066645)))
	}
l58:
	v5 = i32(-3)
	v21 = i32(30)
	goto l30
l93:
	v5 = i32(0)
l30:
	m.memory[uint32(v8)] = byte(v21)
	t589 := int32(load32(m.memory[int64(uint32(v1))+32:]))
	v21 = t589
	t590 := int32(load32(m.memory[int64(uint32(v1))+60:]))
	t591 := v1
	v9 = t590
	t592 := int32(load32(m.memory[int64(uint32(v9))+56:]))
	v25 = t592
	store32(m.memory[int64(uint32(t591))+32:], uint32(v25))
	t593 := int32(load32(m.memory[int64(uint32(v8))+56:]))
	t594 := v1
	v20 = t593
	t595 := int32(load32(m.memory[int64(uint32(v1))+40:]))
	store32(m.memory[int64(uint32(t594))+40:], uint32(v20-v21+t595))
	t596 := int32(load32(m.memory[int64(uint32(v9))+60:]))
	store32(m.memory[int64(uint32(v1))+36:], uint32(t596-v25))
	t597 := int32(load32(m.memory[int64(uint32(v9))+76:]))
	t598 := v1
	v8 = t597
	t599 := int32(load32(m.memory[int64(uint32(v9))+80:]))
	t600 := v8
	v3 = t599
	store32(m.memory[int64(uint32(t598))+48:], uint32(t600-v3))
	t601 := int32(load32(m.memory[int64(uint32(v9))+84:]))
	t602 := int32(load32(m.memory[int64(uint32(v9))+116:]))
	t603 := v9
	v8 = v3 - v8 + t602
	v26 = t601 + v8
	store32(m.memory[int64(uint32(t603))+84:], uint32(v26))
	store32(m.memory[int64(uint32(v1))+52:], uint32(v26))
	t604 := int32(load32(m.memory[int64(uint32(v9))+72:]))
	t605 := v1
	v26 = v3 + t604
	store32(m.memory[int64(uint32(t605))+44:], uint32(v26))
	t606 := int32(load32(m.memory[int64(uint32(v9))+124:]))
	store32(m.memory[int64(uint32(v1))+80:], uint32(t606))
	{
		{
			t607 := int32(load32(m.memory[uint32(v9+i32(12)):]))
			t608 := m.fn1735(t607)
			if t608 != 0 {
				goto l267
			}
			if v8 == 0 {
				goto l268
			}
			{
				t609 := int32(m.memory[uint32(v9)])
				v3 = t609
				switch v3 + i32(-9) {
				case 0, 1:
					goto l268
				case 3, 4, 5, 6:
					goto l267
				default:
					if v3 != i32(30) {
						goto l267
					}
					goto l268
				case 2, 7:
					t610 := int32(m.memory[int64(uint32(v9))+4])
					if t610 == i32(4) {
						goto l268
					}
				}
			}
		}
	l267:
		{
			{
				{
					t611 := int32(load32(m.memory[int64(uint32(v9))+80:]))
					t612 := v8
					v3 = t611
					if uint32(t612) > uint32(v3) {
						m.fn151(i32(0), v8, v3, i32(1295168))
						panic("unreachable")
					}
					v15 = v9 + i32(128)
					t613 := int32(m.memory[int64(uint32(v9))+3])
					v14 = t613 & i32(4)
					t614 := int32(load32(m.memory[int64(uint32(v9))+120:]))
					v27 = t614
					t615 := int32(load32(m.memory[int64(uint32(v9))+72:]))
					v19 = t615
					{
						t616 := int32(load32(m.memory[uint32(v9+i32(12)):]))
						t617 := v8
						v3 = t616
						t618 := m.fn1735(v3)
						v13 = t618
						if uint32(t617) >= uint32(v13) {
							t639 := m.fn1735(v3)
							t640 := v7 + i32(160)
							t641 := v19
							t642 := v8
							v18 = v8 - t639
							p643 := v18
							if uint32(v18) > uint32(v8) {
								p643 = i32(0)
							}
							m.fn1739(t640, t641, t642, p643, i32(1292216))
							t644 := int32(load32(m.memory[int64(uint32(v7))+172:]))
							v19 = t644
							t645 := int32(load32(m.memory[int64(uint32(v7))+168:]))
							v18 = t645
							{
								if v14 == 0 {
									t652 := int32(load32(m.memory[int64(uint32(v9))+8:]))
									m.fn1740(v7+i32(48), v13, t652, v3, i32(1292232))
									t653 := int32(load32(m.memory[int64(uint32(v7))+48:]))
									t654 := int32(load32(m.memory[int64(uint32(v7))+52:]))
									m.fn1689(t653, t654, v18, v19, i32(1292248))
									goto l278
								}
								t646 := int32(load32(m.memory[int64(uint32(v7))+164:]))
								v14 = t646
								t647 := int32(load32(m.memory[int64(uint32(v7))+160:]))
								v16 = t647
								if v27 != 0 {
									m.fn1726(v15, v16, v14)
									t655 := int32(load32(m.memory[int64(uint32(v9))+8:]))
									m.fn1740(v7+i32(56), v13, t655, v3, i32(1292264))
									t656 := int32(load32(m.memory[int64(uint32(v7))+56:]))
									t657 := int32(load32(m.memory[int64(uint32(v7))+60:]))
									m.fn1741(v15, t656, t657, v18, v19)
									goto l278
								}
								t648 := int32(load32(m.memory[int64(uint32(v9))+124:]))
								t649 := m.fn1725(t648, v16, v14)
								t650 := int32(load32(m.memory[int64(uint32(v9))+8:]))
								t651 := m.fn1742(t649, t650, v3, v18, v19)
								store32(m.memory[int64(uint32(v9))+124:], uint32(t651))
								goto l278
							}
						}
						t619 := int32(load32(m.memory[int64(uint32(v9))+20:]))
						t620 := v7 + i32(160)
						t621 := v19
						t622 := v8
						t623 := v8
						t624 := v13
						v16 = t619
						v13 = t624 - v16
						p625 := v13
						if uint32(v8) < uint32(v13) {
							p625 = t623
						}
						v29 = p625
						m.fn1739(t620, t621, t622, v29, i32(1292088))
						t626 := int32(load32(m.memory[int64(uint32(v7))+172:]))
						v19 = t626
						t627 := int32(load32(m.memory[int64(uint32(v7))+168:]))
						v12 = t627
						t628 := int32(load32(m.memory[int64(uint32(v7))+164:]))
						v13 = t628
						t629 := int32(load32(m.memory[int64(uint32(v7))+160:]))
						v17 = t629
						t630 := int32(load32(m.memory[int64(uint32(v9))+8:]))
						v18 = t630
						{
							if v14 != 0 {
								m.fn212(v7+i32(40), v16, v18, v3, i32(1292152))
								t635 := int32(load32(m.memory[int64(uint32(v7))+40:]))
								t636 := int32(load32(m.memory[int64(uint32(v7))+44:]))
								m.fn1740(v7+i32(32), v13, t635, t636, i32(1292168))
								t637 := int32(load32(m.memory[int64(uint32(v7))+36:]))
								v30 = t637
								t638 := int32(load32(m.memory[int64(uint32(v7))+32:]))
								v28 = t638
								if v27 == 0 {
									goto l275
								}
								m.fn1741(v15, v28, v30, v17, v13)
								goto l274
							}
							m.fn212(v7+i32(24), v16, v18, v3, i32(1292104))
							t631 := int32(load32(m.memory[int64(uint32(v7))+24:]))
							t632 := int32(load32(m.memory[int64(uint32(v7))+28:]))
							m.fn1740(v7+i32(16), v13, t631, t632, i32(1292120))
							t633 := int32(load32(m.memory[int64(uint32(v7))+16:]))
							t634 := int32(load32(m.memory[int64(uint32(v7))+20:]))
							m.fn1689(t633, t634, v17, v13, i32(1292136))
							goto l274
						}
					}
				}
			l278:
				store32(m.memory[int64(uint32(v9))+20:], uint32(i32(0)))
				t658 := m.fn1735(v3)
				store32(m.memory[int64(uint32(v9))+16:], uint32(t658))
				goto l268
			}
		l275:
			t659 := int32(load32(m.memory[int64(uint32(v9))+124:]))
			t660 := m.fn1742(t659, v28, v30, v17, v13)
			store32(m.memory[int64(uint32(v9))+124:], uint32(t660))
		}
	l274:
		{
			{
				if v19 == 0 {
					t663 := v9
					v14 = v29 + v16
					store32(m.memory[int64(uint32(t663))+20:], uint32(v14))
					{
						t664 := m.fn1735(v3)
						if v14 != t664 {
							goto l283
						}
						store32(m.memory[int64(uint32(v9))+20:], uint32(i32(0)))
					}
				l283:
					t665 := int32(load32(m.memory[int64(uint32(v9))+16:]))
					v14 = t665
					t666 := m.fn1735(v3)
					if uint32(v14) >= uint32(t666) {
						goto l268
					}
					store32(m.memory[int64(uint32(v9))+16:], uint32(v14+v29))
					goto l268
				}
				m.fn1740(v7+i32(8), v19, v18, v3, i32(1292184))
				t661 := int32(load32(m.memory[int64(uint32(v7))+12:]))
				v13 = t661
				t662 := int32(load32(m.memory[int64(uint32(v7))+8:]))
				v18 = t662
				if v14 != 0 {
					if v27 == 0 {
						goto l282
					}
					m.fn1741(v15, v18, v13, v12, v19)
					goto l281
				}
				m.fn1689(v18, v13, v12, v19, i32(1292200))
				goto l281
			}
		l282:
			t667 := int32(load32(m.memory[int64(uint32(v9))+124:]))
			t668 := m.fn1742(t667, v18, v13, v12, v19)
			store32(m.memory[int64(uint32(v9))+124:], uint32(t668))
		}
	l281:
		store32(m.memory[int64(uint32(v9))+20:], uint32(v19))
		t669 := m.fn1735(v3)
		store32(m.memory[int64(uint32(v9))+16:], uint32(t669))
	}
l268:
	{
		t670 := int32(load32(m.memory[int64(uint32(v9))+132:]))
		v3 = t670
		if v3 == 0 {
			goto l284
		}
		{
			t671 := int32(load32(m.memory[int64(uint32(v9))+136:]))
			t672 := m.fn806(v3, t671, i32(0))
			if t672 != 0 {
				goto l285
			}
			m.fn256(i32(1295184), i32(37), i32(1295224))
			panic("unreachable")
		}
	l285:
		store32(m.memory[int64(uint32(v1))+56:], uint32(v3))
	}
l284:
	t673 := int32(m.memory[int64(uint32(v9))+1])
	v14 = t673
	t674 := int32(m.memory[int64(uint32(v9))+64])
	v19 = t674
	v3 = i32(0)
	{
		t675 := int32(m.memory[uint32(v9)])
		v9 = (t675 + i32(-12)) & i32(255)
		if uint32(v9) > uint32(i32(5)) {
			goto l286
		}
		t676 := int32(load32(m.memory[int64(uint32(v9<<2))+1301748:]))
		v3 = t676
	}
l286:
	t677 := int64(load64(m.memory[int64(uint32(v1))+16:]))
	t678 := v1
	v22 = t677 + int64(uint32(v25-v2))
	store64(m.memory[int64(uint32(t678))+16:], uint64(v22))
	t679 := int64(load64(m.memory[int64(uint32(v1))+24:]))
	t680 := v1
	v23 = t679 + int64(uint32(v26-v4))
	store64(m.memory[int64(uint32(t680))+24:], uint64(v23))
	store32(m.memory[int64(uint32(v1))+76:], uint32(v3|(v14<<6&i32(64)|v19)&i32(255)))
	{
		{
			p681 := i32(-5)
			if v5 != 0 {
				p681 = v5
			}
			v9 = p681
			t683 := v9
			p682 := v9
			if v8 != 0 {
				p682 = v5
			}
			p684 := v5
			if v20 == v21 {
				p684 = p682
			}
			p685 := p684
			if v6&i32(255) == i32(4) {
				p685 = t683
			}
			v8 = p685
			switch v8 + i32(5) {
			case 1, 2, 3:
				t691 := int64(load64(m.memory[uint32(v1):]))
				store64(m.memory[uint32(v1):], uint64(v22-v11+t691))
				t692 := int64(load64(m.memory[int64(uint32(v1))+8:]))
				store64(m.memory[int64(uint32(v1))+8:], uint64(v23-v10+t692))
				v5 = v8 + i32(-2)
				switch v5 {
				case 0:
					goto l294
				case 1:
					goto l293
				default:
					{
						{
							t693 := int32(load32(m.memory[int64(uint32(v1))+56:]))
							v8 = t693
							if v8 != 0 {
								goto l296
							}
							v1 = i32(0)
							goto l297
						}
					l296:
						t694 := m.fn1852(v8)
						m.fn12(v7+i32(160), v8, t694)
						t695 := int32(load32(m.memory[int64(uint32(v7))+164:]))
						t696 := int32(load32(m.memory[int64(uint32(v7))+160:]))
						p697 := t695
						if t696 != 0 {
							p697 = i32(0)
						}
						v1 = p697
						t698 := int32(load32(m.memory[int64(uint32(v7))+168:]))
						v8 = t698
					}
				l297:
					store32(m.memory[int64(uint32(v0))+8:], uint32(v8))
					store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
					store32(m.memory[uint32(v0):], uint32(i32(0)))
					goto l298
				}
			case 5:
				goto l290
			case 7:
				t688 := int64(load64(m.memory[uint32(v1):]))
				store64(m.memory[uint32(v1):], uint64(v22-v11+t688))
				t689 := int64(load64(m.memory[int64(uint32(v1))+8:]))
				store64(m.memory[int64(uint32(v1))+8:], uint64(v23-v10+t689))
				t690 := int32(load32(m.memory[int64(uint32(v1))+80:]))
				v9 = t690
				goto l294
			case 6:
				v8 = i32(2)
				goto l290
			case 4:
				m.fn91(i32(1292360), i32(147), i32(1292436))
				panic("unreachable")
			default:
				v8 = i32(1)
			}
		}
	l290:
		t686 := int64(load64(m.memory[uint32(v1):]))
		store64(m.memory[uint32(v1):], uint64(v22-v11+t686))
		t687 := int64(load64(m.memory[int64(uint32(v1))+8:]))
		store64(m.memory[int64(uint32(v1))+8:], uint64(v23-v10+t687))
		goto l293
	}
l294:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v9))
	store32(m.memory[uint32(v0):], uint32(i32(1)))
	goto l298
l293:
	store32(m.memory[uint32(v0):], uint32(i32(2)))
	m.memory[int64(uint32(v0))+4] = byte(v8)
l298:
	m.g0 = v7 + i32(176)
}
func (m *Module) fn307(v0, v1, v2, v3 int32) {
	var v4, v5, v6, v7, v8 int32
	t0 := m.g0
	v4 = t0 - i32(16)
	m.g0 = v4
	t1 := int32(load32(m.memory[uint32(v1):]))
	t2 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	t3 := v4
	v5 = t2
	t5 := v5
	p4 := v3
	if uint32(v5) < uint32(v3) {
		p4 = v5
	}
	v3 = p4
	m.fn309(t3, t1, t5, v3, i32(1087032))
	t6 := int32(load32(m.memory[int64(uint32(v4))+12:]))
	v6 = t6
	t7 := int32(load32(m.memory[int64(uint32(v4))+8:]))
	v7 = t7
	t8 := int32(load32(m.memory[int64(uint32(v4))+4:]))
	v5 = t8
	t9 := int32(load32(m.memory[uint32(v4):]))
	v8 = t9
	{
		if v3 != i32(1) {
			m.fn310(v2, v3, v8, v5, i32(1087064))
			goto l2
		}
		if v5 == 0 {
			m.fn158(i32(0), i32(0), i32(1087048))
			panic("unreachable")
		}
		t10 := int32(m.memory[uint32(v8)])
		m.memory[uint32(v2)] = byte(t10)
		goto l2
	}
l2:
	store32(m.memory[int64(uint32(v1))+4:], uint32(v6))
	store32(m.memory[uint32(v1):], uint32(v7))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
	m.memory[uint32(v0)] = byte(i32(255))
	m.g0 = v4 + i32(16)
}
func (m *Module) fn308(v0, v1 int32) int32 {
	var v2 int32
	t0 := int32(load16(m.memory[int64(uint32(v0))+8:]))
	t1 := v0
	t2 := v1
	v2 = t0 | i32(2)
	v1 = t2 ^ int32(uint32((v2^i32(1))*v2&i32(0xff00))>>8)
	m.fn1833(t1, v1)
	return v1
}
func (m *Module) fn309(v0, v1, v2, v3, v4 int32) {
	if uint32(v2) >= uint32(v3) {
		goto l0
	}
	m.fn91(i32(1301172), i32(19), v4)
	panic("unreachable")
l0:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
	store32(m.memory[uint32(v0):], uint32(v1))
	store32(m.memory[int64(uint32(v0))+12:], uint32(v2-v3))
	store32(m.memory[int64(uint32(v0))+8:], uint32(v1+v3))
}
func (m *Module) fn310(v0, v1, v2, v3, v4 int32) {
	if v1 != v3 {
		m.fn1668(v1, v3, v4)
		panic("unreachable")
	}
	if v1 == 0 {
		return
	}
	memory_copy(m.memory, uint32(v0), uint32(v2), uint32(v1))
}
func (m *Module) fn311(v0, v1 int32) {
	var v2 int32
	var v3, v4 int64
	t0 := int32(load32(m.memory[uint32(v1):]))
	t1 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	t2 := v0
	v2 = t1
	t3 := int64(load64(m.memory[int64(uint32(v1))+8:]))
	t4 := v2
	v3 = t3
	t5 := v3
	v4 = int64(uint32(v2))
	p6 := v4
	if uint64(v3) < uint64(v4) {
		p6 = t5
	}
	m.fn309(t2, t0, t4, int32(p6), i32(1286696))
}
