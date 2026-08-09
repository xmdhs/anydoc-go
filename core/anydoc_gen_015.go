package core

import (
	"math/bits"
)

func (m *Module) fn627(v0, v1 int32) {
	var v2, v3, v4, v5, v6, v7, v8, v9 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	t1 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	v3 = t1
	t2 := int32(load32(m.memory[int64(uint32(v1))+16:]))
	t3 := v3
	v4 = t2
	v5 = t3 + v4
	t4 := int32(load32(m.memory[uint32(v1):]))
	v6 = t4
	t5 := int32(load32(m.memory[int64(uint32(v1))+12:]))
	v7 = t5
l2:
	{
		store32(m.memory[int64(uint32(v2))+12:], uint32(v5))
		t6 := v2
		t7 := v3
		v8 = v7
		store32(m.memory[int64(uint32(t6))+8:], uint32(t7+v8))
		m.fn374(v2, v2+i32(8))
		{
			t8 := int32(load32(m.memory[uint32(v2):]))
			if t8&i32(1) != 0 {
				goto l0
			}
			v7 = i32(0)
			goto l1
		}
	l0:
		t9 := int32(load32(m.memory[int64(uint32(v2))+4:]))
		v9 = t9
		t10 := int32(load32(m.memory[int64(uint32(v2))+8:]))
		t11 := int32(load32(m.memory[int64(uint32(v2))+12:]))
		t12 := v1
		v7 = t10 - t11 + v4
		store32(m.memory[int64(uint32(t12))+12:], uint32(v7))
		if v9 == v6 {
			goto l2
		}
	}
	store32(m.memory[int64(uint32(v0))+8:], uint32(v7))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v8))
	v7 = i32(1)
l1:
	store32(m.memory[uint32(v0):], uint32(v7))
	m.g0 = v2 + i32(16)
}
func (m *Module) fn628(v0, v1, v2 int32) {
	var v3 int32
	t0 := m.g0
	v3 = t0 - i32(32)
	m.g0 = v3
	store32(m.memory[int64(uint32(v3))+16:], uint32(i32(0)))
	store32(m.memory[int64(uint32(v3))+8:], uint32(v1))
	store32(m.memory[uint32(v3):], uint32(v1))
	store32(m.memory[int64(uint32(v3))+4:], uint32(v2))
	store32(m.memory[int64(uint32(v3))+12:], uint32(v1+v2))
l0:
	{
		m.fn629(v3+i32(20), v3)
		t1 := int32(load32(m.memory[int64(uint32(v3))+20:]))
		switch t1 {
		case 1:
			t2 := int32(load32(m.memory[int64(uint32(v3))+28:]))
			v2 = t2
			goto l3
		case 2:
			goto l2
		default:
			goto l0
		}
	}
l2:
	v2 = i32(0)
l3:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v2))
	store32(m.memory[uint32(v0):], uint32(v1))
	m.g0 = v3 + i32(32)
}
func (m *Module) fn629(v0, v1 int32) {
	var v2, v3, v4, v5, v6, v7 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	t1 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	v3 = t1
	t2 := int32(load32(m.memory[int64(uint32(v1))+12:]))
	v4 = t2
	m.fn574(v2+i32(8), v1+i32(8))
	{
		{
			t3 := int32(load32(m.memory[int64(uint32(v2))+12:]))
			v5 = t3
			if v5 != i32(-1) {
				goto l0
			}
			v1 = i32(2)
			goto l1
		}
	l0:
		t4 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		v6 = t4
		t5 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		v1 = t5
		t6 := int32(load32(m.memory[int64(uint32(v2))+8:]))
		t7 := v0
		v7 = t6
		store32(m.memory[int64(uint32(t7))+4:], uint32(v7))
		store32(m.memory[int64(uint32(v0))+8:], uint32(v7+(v6-v1+(v4-v3))))
		t8 := m.fn630(v5)
		v1 = t8 ^ i32(1)
	}
l1:
	store32(m.memory[uint32(v0):], uint32(v1))
	m.g0 = v2 + i32(16)
}
func (m *Module) fn630(v0 int32) int32 {
	var v1, v2 int32
	{
		v1 = v0 + i32(-9)
		if uint32(v1) > uint32(i32(23)) {
			goto l0
		}
		v2 = i32(1)
		if i32_shr_u(i32(8388639), v1)&i32(1) != 0 {
			goto l1
		}
	l0:
		v2 = i32(0)
		if uint32(v0) < uint32(i32(133)) {
			goto l1
		}
		v1 = int32(uint32(v0) >> 8)
		if v1 == 0 {
			t2 := int32(m.memory[int64(uint32(v0&i32(255)))+1148316])
			v2 = t2
			goto l1
		}
		{
			if v1 == i32(48) {
				var p1 int32
				if v0 == i32(12288) {
					p1 = 1
				}
				v2 = p1
				goto l1
			}
			if v1 == i32(32) {
				goto l4
			}
			if v1 != i32(22) {
				goto l1
			}
			var p0 int32
			if v0 == i32(5760) {
				p0 = 1
			}
			v2 = p0
			goto l1
		}
	l4:
		t3 := int32(m.memory[int64(uint32(v0&i32(255)))+1148316])
		v2 = int32(uint32(t3&i32(2)) >> 1)
	}
l1:
	return v2 & i32(1)
}
func (m *Module) fn631(v0, v1, v2, v3 int32) {
	var v4 int32
	t0 := m.g0
	v4 = t0 - i32(48)
	m.g0 = v4
	m.fn514(v4+i32(8), v3, v1, v2)
	m.fn627(v4+i32(36), v4+i32(8))
	t1 := int32(load32(m.memory[int64(uint32(v4))+40:]))
	t2 := int32(load32(m.memory[int64(uint32(v4))+36:]))
	t4 := v0
	t5 := v2
	p3 := v2
	if t2 != 0 {
		p3 = t1
	}
	v3 = p3
	store32(m.memory[int64(uint32(t4))+4:], uint32(t5-v3))
	store32(m.memory[uint32(v0):], uint32(v1+v3))
	m.g0 = v4 + i32(48)
}
func (m *Module) fn632(v0, v1, v2, v3, v4 int32) {
	var v5, v6, v7, v8, v9, v10, v11 int32
	t0 := m.g0
	v5 = t0 - i32(48)
	m.g0 = v5
	{
		if v2 == 0 {
			store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
			store64(m.memory[uint32(v0):], uint64(i64(0x100000000)))
			goto l13
		}
		{
			t1 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			v6 = t1
			t2 := v6
			t3 := v4
			v2 = v2<<3 + i32(-8)
			v7 = t2 + t3*int32(uint32(v2)>>3)
			if uint32(v7) < uint32(v6) {
				goto l1
			}
			t4 := int32(load32(m.memory[uint32(v1):]))
			v8 = t4
			v9 = v2
			v1 = v1 + i32(8)
			v10 = v1
		l3:
			{
				if v9 == 0 {
					m.fn59(v5+i32(8), v7, i32(1), i32(1))
					store32(m.memory[int64(uint32(v5))+28:], uint32(i32(0)))
					t6 := int64(load64(m.memory[int64(uint32(v5))+8:]))
					store64(m.memory[int64(uint32(v5))+20:], uint64(t6))
					m.fn634(v5+i32(20), v8, v8+v6)
					t7 := int32(load32(m.memory[int64(uint32(v5))+28:]))
					t8 := v7
					v10 = t7
					v9 = t8 - v10
					t9 := int32(load32(m.memory[int64(uint32(v5))+24:]))
					v10 = t9 + v10
					switch v4 + i32(-1) {
					default:
					l9:
						{
							if v2 == 0 {
								goto l8
							}
							t10 := int32(load32(m.memory[uint32(v1):]))
							v4 = t10
							t11 := int32(load32(m.memory[int64(uint32(v1))+4:]))
							v11 = t11
							m.fn635(v5+i32(32), v10, v9, i32(1))
							t12 := int32(load32(m.memory[int64(uint32(v5))+44:]))
							v9 = t12
							t13 := int32(load32(m.memory[int64(uint32(v5))+40:]))
							v10 = t13
							t14 := int32(load32(m.memory[int64(uint32(v5))+32:]))
							t15 := int32(load32(m.memory[int64(uint32(v5))+36:]))
							m.fn310(t14, t15, v3, i32(1), i32(1300940))
							m.fn635(v5+i32(32), v10, v9, v11)
							t16 := int32(load32(m.memory[int64(uint32(v5))+44:]))
							v9 = t16
							t17 := int32(load32(m.memory[int64(uint32(v5))+40:]))
							v10 = t17
							t18 := int32(load32(m.memory[int64(uint32(v5))+32:]))
							t19 := int32(load32(m.memory[int64(uint32(v5))+36:]))
							m.fn310(t18, t19, v4, v11, i32(1300940))
							v2 = v2 + i32(-8)
							v1 = v1 + i32(8)
							goto l9
						}
					case 1:
					l10:
						{
							if v2 == 0 {
								goto l8
							}
							t20 := int32(load32(m.memory[uint32(v1):]))
							v4 = t20
							t21 := int32(load32(m.memory[int64(uint32(v1))+4:]))
							v11 = t21
							m.fn635(v5+i32(32), v10, v9, i32(2))
							t22 := int32(load32(m.memory[int64(uint32(v5))+44:]))
							v9 = t22
							t23 := int32(load32(m.memory[int64(uint32(v5))+40:]))
							v10 = t23
							t24 := int32(load32(m.memory[int64(uint32(v5))+32:]))
							t25 := int32(load32(m.memory[int64(uint32(v5))+36:]))
							m.fn310(t24, t25, v3, i32(2), i32(1300940))
							m.fn635(v5+i32(32), v10, v9, v11)
							t26 := int32(load32(m.memory[int64(uint32(v5))+44:]))
							v9 = t26
							t27 := int32(load32(m.memory[int64(uint32(v5))+40:]))
							v10 = t27
							t28 := int32(load32(m.memory[int64(uint32(v5))+32:]))
							t29 := int32(load32(m.memory[int64(uint32(v5))+36:]))
							m.fn310(t28, t29, v4, v11, i32(1300940))
							v2 = v2 + i32(-8)
							v1 = v1 + i32(8)
							goto l10
						}
					case 2:
					l11:
						{
							if v2 == 0 {
								goto l8
							}
							t30 := int32(load32(m.memory[uint32(v1):]))
							v4 = t30
							t31 := int32(load32(m.memory[int64(uint32(v1))+4:]))
							v11 = t31
							m.fn635(v5+i32(32), v10, v9, i32(3))
							t32 := int32(load32(m.memory[int64(uint32(v5))+44:]))
							v9 = t32
							t33 := int32(load32(m.memory[int64(uint32(v5))+40:]))
							v10 = t33
							t34 := int32(load32(m.memory[int64(uint32(v5))+32:]))
							t35 := int32(load32(m.memory[int64(uint32(v5))+36:]))
							m.fn310(t34, t35, v3, i32(3), i32(1300940))
							m.fn635(v5+i32(32), v10, v9, v11)
							t36 := int32(load32(m.memory[int64(uint32(v5))+44:]))
							v9 = t36
							t37 := int32(load32(m.memory[int64(uint32(v5))+40:]))
							v10 = t37
							t38 := int32(load32(m.memory[int64(uint32(v5))+32:]))
							t39 := int32(load32(m.memory[int64(uint32(v5))+36:]))
							m.fn310(t38, t39, v4, v11, i32(1300940))
							v2 = v2 + i32(-8)
							v1 = v1 + i32(8)
							goto l11
						}
					case 3:
					l12:
						{
							if v2 == 0 {
								goto l8
							}
							t40 := int32(load32(m.memory[uint32(v1):]))
							v4 = t40
							t41 := int32(load32(m.memory[int64(uint32(v1))+4:]))
							v11 = t41
							m.fn635(v5+i32(32), v10, v9, i32(4))
							t42 := int32(load32(m.memory[int64(uint32(v5))+44:]))
							v9 = t42
							t43 := int32(load32(m.memory[int64(uint32(v5))+40:]))
							v10 = t43
							t44 := int32(load32(m.memory[int64(uint32(v5))+32:]))
							t45 := int32(load32(m.memory[int64(uint32(v5))+36:]))
							m.fn310(t44, t45, v3, i32(4), i32(1300940))
							m.fn635(v5+i32(32), v10, v9, v11)
							t46 := int32(load32(m.memory[int64(uint32(v5))+44:]))
							v9 = t46
							t47 := int32(load32(m.memory[int64(uint32(v5))+40:]))
							v10 = t47
							t48 := int32(load32(m.memory[int64(uint32(v5))+32:]))
							t49 := int32(load32(m.memory[int64(uint32(v5))+36:]))
							m.fn310(t48, t49, v4, v11, i32(1300940))
							v2 = v2 + i32(-8)
							v1 = v1 + i32(8)
							goto l12
						}
					}
				}
				v9 = v9 + i32(-8)
				t5 := int32(load32(m.memory[int64(uint32(v10))+4:]))
				v11 = t5
				v10 = v10 + i32(8)
				v7 = v11 + v7
				if uint32(v7) >= uint32(v11) {
					goto l3
				}
			}
		}
	l1:
		m.fn633(i32(1300956), i32(53), i32(1301012))
		panic("unreachable")
	l8:
		t50 := int64(load64(m.memory[int64(uint32(v5))+20:]))
		store64(m.memory[uint32(v0):], uint64(t50))
		store32(m.memory[int64(uint32(v0))+8:], uint32(v7-v9))
	}
l13:
	m.g0 = v5 + i32(48)
}
func (m *Module) fn633(v0, v1, v2 int32) {
	var v3 int32
	t0 := m.g0
	v3 = t0 - i32(16)
	m.g0 = v3
	store32(m.memory[int64(uint32(v3))+4:], uint32(v1))
	store32(m.memory[uint32(v3):], uint32(v0))
	store64(m.memory[int64(uint32(v3))+8:], uint64(int64(uint32(i32(41)))<<32|int64(uint32(v3))))
	m.fn91(i32(1052692), v3+i32(8), v2)
	panic("unreachable")
}
func (m *Module) fn634(v0, v1, v2 int32) {
	var v3, v4, v5, v6, v7, v8, v9, v10 int32
	t0 := m.g0
	v3 = t0 - i32(32)
	m.g0 = v3
	{
		v4 = v2 - v1
		t1 := int32(load32(m.memory[uint32(v0):]))
		t2 := v4
		v5 = t1
		t3 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		t4 := v5
		v6 = t3
		if uint32(t2) <= uint32(t4-v6) {
			goto l0
		}
		{
			v7 = v4 + v6
			if uint32(v7) >= uint32(v4) {
				goto l1
			}
			v0 = i32(0)
			goto l2
		l1:
			v8 = i32(0)
			v9 = v3 + i32(20)
			{
				t5 := v7
				v10 = v5 << 1
				p6 := v10
				if uint32(v7) > uint32(v10) {
					p6 = t5
				}
				v7 = p6
				p7 := i32(8)
				if uint32(v7) > uint32(i32(8)) {
					p7 = v7
				}
				v7 = p7
				if v7 < i32(0) {
					goto l3
				}
				{
					if v5 != 0 {
						goto l4
					}
					v5 = i32(0)
					v10 = v3 + i32(28)
					goto l5
				l4:
					t8 := int32(load32(m.memory[int64(uint32(v0))+4:]))
					v8 = t8
					store32(m.memory[int64(uint32(v3))+28:], uint32(i32(1)))
					v10 = v3 + i32(24)
				}
			l5:
				store32(m.memory[uint32(v10):], uint32(v5))
				{
					{
						t9 := int32(load32(m.memory[int64(uint32(v3))+28:]))
						if t9 == 0 {
							goto l6
						}
						{
							t10 := int32(load32(m.memory[int64(uint32(v3))+24:]))
							v5 = t10
							if v5 != 0 {
								t12 := m.fn89(v8, v5, i32(1), v7)
								v5 = t12
								goto l8
							}
							m.fn940(v3+i32(8), v7)
							t11 := int32(load32(m.memory[int64(uint32(v3))+8:]))
							v5 = t11
							goto l8
						}
					}
				l6:
					m.fn940(v3, v7)
					t13 := int32(load32(m.memory[uint32(v3):]))
					v5 = t13
				}
			l8:
				if v5 != 0 {
					goto l9
				}
				store32(m.memory[int64(uint32(v3))+20:], uint32(i32(1)))
				v9 = v3 + i32(16)
				v8 = v7
			}
		l3:
			store32(m.memory[uint32(v9):], uint32(v8))
			t14 := int32(load32(m.memory[int64(uint32(v3))+16:]))
			v1 = t14
			t15 := int32(load32(m.memory[int64(uint32(v3))+20:]))
			v0 = t15
		}
	l2:
		m.fn2(v0, v1)
		panic("unreachable")
	l9:
		store32(m.memory[uint32(v0):], uint32(v7))
		store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
	}
l0:
	{
		if v2 == v1 {
			goto l10
		}
		if v4 == 0 {
			goto l10
		}
		t16 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		memory_copy(m.memory, uint32(t16+v6), uint32(v1), uint32(v4))
	}
l10:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v6+v4))
	m.g0 = v3 + i32(32)
}
func (m *Module) fn635(v0, v1, v2, v3 int32) {
	if uint32(v2) >= uint32(v3) {
		goto l0
	}
	m.fn91(i32(1301172), i32(19), i32(1300940))
	panic("unreachable")
l0:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
	store32(m.memory[uint32(v0):], uint32(v1))
	store32(m.memory[int64(uint32(v0))+12:], uint32(v2-v3))
	store32(m.memory[int64(uint32(v0))+8:], uint32(v1+v3))
}
func (m *Module) fn636(v0, v1, v2 int32) int32 {
	var v3 int32
	t0 := m.g0
	v3 = t0 - i32(32)
	m.g0 = v3
	m.fn637(v3+i32(20), v1, v2)
	{
		{
			t1 := int32(load32(m.memory[int64(uint32(v3))+20:]))
			v2 = t1
			if v2 == i32(-1) {
				goto l0
			}
			t2 := int32(load32(m.memory[int64(uint32(v3))+24:]))
			t3 := v3 + i32(8)
			t4 := v0
			v1 = t2
			t5 := int32(load32(m.memory[int64(uint32(v3))+28:]))
			m.fn638(t3, t4, v1, t5)
			t6 := int32(load32(m.memory[int64(uint32(v3))+8:]))
			v0 = t6
			m.fn639(v2, v1)
			var p7 int32
			if v0 != i32(0) {
				p7 = 1
			}
			v2 = p7
			goto l1
		}
	l0:
		t8 := int32(m.memory[int64(uint32(v3))+24])
		t9 := int32(load32(m.memory[int64(uint32(v3))+28:]))
		m.fn119(t8, t9)
		v2 = i32(0)
	}
l1:
	m.g0 = v3 + i32(32)
	return v2
}
func (m *Module) fn637(v0, v1, v2 int32) {
	var v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14, v15 int32
	t0 := m.g0
	v3 = t0 - i32(32)
	m.g0 = v3
	store32(m.memory[int64(uint32(v3))+16:], uint32(i32(0)))
	store64(m.memory[int64(uint32(v3))+8:], uint64(i64(0x400000000)))
	v4 = i32(4)
	t1 := int32(m.memory[uint32(v1)])
	var p2 int32
	if t1&i32(255) == i32(47) {
		p2 = 1
	}
	v5 = p2
	v6 = i32(0)
	v7 = i32(1)
	{
	l36:
		{
			if v5 != 0 {
				v9 = v7
			l14:
				v10 = v1
				if v9&i32(255) == i32(1) {
					if v2 == 0 {
						m.fn151(i32(1), i32(0), i32(0), i32(1285356))
						panic("unreachable")
					}
					v1 = v10 + i32(1)
					v2 = v2 + i32(-1)
					v12 = i32(6)
					v7 = i32(2)
					goto l22
				}
				if v2 == 0 {
					goto l3
				}
				v11 = i32(0)
			l6:
				{
					v8 = i32(1)
					{
						t3 := int32(m.memory[uint32(v10+v11)])
						if t3 != i32(47) {
							goto l4
						}
						v9 = i32(1)
						goto l5
					}
				l4:
					t4 := v2
					v11 = v11 + i32(1)
					if t4 != v11 {
						goto l6
					}
				}
				v9 = i32(0)
				v11 = v2
			l5:
				v12 = i32(255)
				switch v11 {
				case 0:
					goto l7
				default:
					goto l10
				case 2:
					t5 := int32(m.memory[uint32(v10)])
					if t5 != i32(46) {
						goto l10
					}
					t6 := int32(m.memory[int64(uint32(v10))+1])
					if t6&i32(255) != i32(46) {
						goto l10
					}
					v12 = i32(8)
					goto l11
				case 1:
					t7 := int32(m.memory[uint32(v10)])
					if t7 == i32(46) {
						goto l7
					}
				}
			l10:
				v12 = i32(9)
			l11:
				v8 = i32(0)
			l7:
				{
					t8 := v2
					v13 = v9 + v11
					if uint32(t8) >= uint32(v13) {
						v1 = v10 + v13
						v2 = v2 - v13
						v9 = i32(2)
						v14 = v10
						if v8 != 0 {
							goto l14
						}
						v14 = v10
						goto l15
					}
					v8 = v2
					goto l13
				}
			}
			v8 = v2
			goto l17
		l17:
			v15 = v7
		l31:
			{
				v10 = v1
				if v15&i32(255) != i32(1) {
					goto l16
				}
				v7 = i32(2)
				v1 = v10
				switch v8 {
				case 0:
					goto l17
				default:
					v1 = v10
					t9 := int32(m.memory[uint32(v10)])
					if t9 != i32(46) {
						goto l17
					}
					v1 = v10
					t10 := int32(m.memory[int64(uint32(v10))+1])
					if t10 == i32(47) {
						goto l20
					}
					goto l17
				case 1:
					v8 = i32(1)
					v1 = v10
					t11 := int32(m.memory[uint32(v10)])
					if t11 != i32(46) {
						goto l17
					}
				}
			l20:
				if v2 == 0 {
					m.fn151(i32(1), i32(0), i32(0), i32(1285340))
					panic("unreachable")
				}
				v1 = v10 + i32(1)
				v2 = v2 + i32(-1)
				v12 = i32(7)
				v7 = i32(2)
				goto l22
			l16:
				if v8 == 0 {
					goto l3
				}
				v11 = i32(0)
			l25:
				{
					v9 = i32(1)
					{
						t12 := int32(m.memory[uint32(v10+v11)])
						if t12 != i32(47) {
							goto l23
						}
						v1 = i32(1)
						goto l24
					}
				l23:
					t13 := v8
					v11 = v11 + i32(1)
					if t13 != v11 {
						goto l25
					}
				}
				v1 = i32(0)
				v11 = v8
			l24:
				v12 = i32(255)
				switch v11 {
				case 0:
					goto l26
				default:
					goto l29
				case 1:
					t14 := int32(m.memory[uint32(v10)])
					if t14 == i32(46) {
						goto l26
					}
					goto l29
				case 2:
					t15 := int32(m.memory[uint32(v10)])
					if t15 != i32(46) {
						goto l29
					}
					t16 := int32(m.memory[int64(uint32(v10))+1])
					if t16&i32(255) != i32(46) {
						goto l29
					}
					v12 = i32(8)
					goto l30
				}
			l29:
				v12 = i32(9)
			l30:
				v9 = i32(0)
			l26:
				t17 := v8
				v13 = v1 + v11
				if uint32(t17) < uint32(v13) {
					goto l13
				}
				v1 = v10 + v13
				v15 = i32(2)
				v14 = v10
				v2 = v8 - v13
				v8 = v2
				if v9 != 0 {
					goto l31
				}
			}
			v14 = v10
		l15:
			if v12 != i32(255) {
				goto l22
			}
		l3:
			t18 := int32(load32(m.memory[int64(uint32(v3))+16:]))
			store32(m.memory[int64(uint32(v0))+8:], uint32(t18))
			t19 := int64(load64(m.memory[int64(uint32(v3))+8:]))
			store64(m.memory[uint32(v0):], uint64(t19))
			goto l32
		}
	l22:
		{
			v10 = v12 + i32(-5)
			p20 := v10
			if uint32(v10) > uint32(v12) {
				p20 = i32(0)
			}
			switch p20 {
			case 2:
				goto l36
			case 1:
				v6 = i32(0)
				store32(m.memory[int64(uint32(v3))+16:], uint32(i32(0)))
				goto l36
			default:
				m.fn1622(v0+i32(4), i32(1101900), i32(35))
				goto l39
			case 3:
				if v6 == 0 {
					m.fn1622(v0+i32(4), i32(1101935), i32(34))
					goto l39
				}
				t21 := v3
				v6 = v6 + i32(-1)
				store32(m.memory[int64(uint32(t21))+16:], uint32(v6))
				goto l36
			case 4:
				m.fn12(v3+i32(20), v14, v11)
				{
					t22 := int32(load32(m.memory[int64(uint32(v3))+20:]))
					if t22 != 0 {
						goto l41
					}
					t23 := int32(load32(m.memory[int64(uint32(v3))+28:]))
					v10 = t23
					t24 := int32(load32(m.memory[int64(uint32(v3))+24:]))
					v8 = t24
					{
						t25 := int32(load32(m.memory[int64(uint32(v3))+8:]))
						if v6 != t25 {
							goto l42
						}
						m.fn1523(v3 + i32(8))
						t26 := int32(load32(m.memory[int64(uint32(v3))+12:]))
						v4 = t26
					}
				l42:
					v9 = v4 + v6<<3
					store32(m.memory[int64(uint32(v9))+4:], uint32(v10))
					store32(m.memory[uint32(v9):], uint32(v8))
					t27 := v3
					v6 = v6 + i32(1)
					store32(m.memory[int64(uint32(t27))+16:], uint32(v6))
					goto l36
				}
			l41:
			}
		}
		m.fn1622(v0+i32(4), i32(1101969), i32(14))
	l39:
		store32(m.memory[uint32(v0):], uint32(i32(-1)))
		t28 := int32(load32(m.memory[int64(uint32(v3))+8:]))
		t29 := int32(load32(m.memory[int64(uint32(v3))+12:]))
		m.fn608(t28, t29, i32(4), i32(8))
	}
l32:
	m.g0 = v3 + i32(32)
	return
l13:
	m.fn151(v13, v8, v8, i32(1285372))
	panic("unreachable")
}
func (m *Module) fn638(v0, v1, v2, v3 int32) {
	var v4, v5, v6, v7, v8, v9, v10 int32
	t0 := m.g0
	v4 = t0 - i32(16)
	m.g0 = v4
	m.fn640(v4+i32(8), v1)
	v5 = v2 + v3<<3
	t1 := int32(load32(m.memory[int64(uint32(v4))+8:]))
	v1 = t1
	t2 := int32(load32(m.memory[uint32(v1+i32(84)):]))
	v6 = t2
	t3 := int32(load32(m.memory[uint32(v1+i32(80)):]))
	v7 = t3
	v8 = i32(0)
	t4 := int32(load32(m.memory[int64(uint32(v4))+12:]))
	v9 = t4
l8:
	if v2 != v5 {
		t5 := m.fn590(v7, v6, v8)
		v1 = t5 + i32(48)
		v10 = v2 + i32(8)
	l7:
		{
			t6 := int32(load32(m.memory[uint32(v1):]))
			v3 = t6
			if v3 != i32(-1) {
				t7 := m.fn590(v7, v6, v3)
				v1 = t7
				{
					t8 := int32(load32(m.memory[uint32(v2):]))
					t9 := int32(load32(m.memory[int64(uint32(v2))+4:]))
					t10 := int32(load32(m.memory[int64(uint32(v1))+64:]))
					t11 := int32(load32(m.memory[int64(uint32(v1))+68:]))
					t12 := m.fn592(t8, t9, t10, t11)
					switch t12 & i32(255) {
					case 0:
						v2 = v10
						v8 = v3
						goto l8
					default:
						v3 = i32(40)
						goto l6
					case 1:
						v3 = i32(44)
					}
				}
			l6:
				v1 = v1 + v3
				goto l7
			}
			v1 = i32(0)
			goto l1
		}
	}
	v1 = i32(1)
	goto l1
l1:
	m.fn641(v9)
	store32(m.memory[int64(uint32(v0))+4:], uint32(v8))
	store32(m.memory[uint32(v0):], uint32(v1))
	m.g0 = v4 + i32(16)
}
func (m *Module) fn639(v0, v1 int32) {
	m.fn608(v0, v1, i32(4), i32(8))
}
func (m *Module) fn640(v0, v1 int32) {
	var v2 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	m.fn642(v2+i32(4), v1+i32(8))
	t1 := int64(load64(m.memory[int64(uint32(v2))+8:]))
	store64(m.memory[uint32(v0):], uint64(t1))
	m.g0 = v2 + i32(16)
}
func (m *Module) fn641(v0 int32) {
	var v1 int32
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := v0
	v1 = t0
	store32(m.memory[uint32(t1):], uint32(v1+i32(-1)))
	if v1 > i32(0) {
		return
	}
	m.fn91(i32(1285284), i32(77), i32(0x139ccc))
	panic("unreachable")
}
func (m *Module) fn642(v0, v1 int32) {
	var v2, v3 int32
	{
		t0 := int32(load32(m.memory[uint32(v1):]))
		v2 = t0
		if v2 <= i32(-1) {
			panic("unreachable")
		}
		v3 = v2 + i32(1)
		if v3 < v2 {
			m.fn633(i32(1284596), i32(28), i32(1284624))
			panic("unreachable")
		}
		store32(m.memory[int64(uint32(v0))+8:], uint32(v1))
		store32(m.memory[uint32(v1):], uint32(v3))
		store32(m.memory[uint32(v0):], uint32(i32(0)))
		store32(m.memory[int64(uint32(v0))+4:], uint32(v1+i32(8)))
		return
	}
}
func (m *Module) fn643(v0, v1, v2, v3 int32) int32 {
	t1 := v0
	t2 := v2
	p0 := v3
	if uint32(v1) < uint32(v3) {
		p0 = v1
	}
	t3 := m.fn1851(t1, t2, p0)
	v2 = t3
	p4 := v1 - v3
	if v2 != 0 {
		p4 = v2
	}
	v3 = p4
	var p5 int32
	if v3 > i32(0) {
		p5 = 1
	}
	var p6 int32
	if v3 < i32(0) {
		p6 = 1
	}
	return p5 - p6
}
func (m *Module) fn644(v0, v1 int32, v2 int64, v3 int32) int32 {
	var v4 int32
	var v5 int64
	var v6, v7, v8 int32
	var v9 int64
	var v10, v11 int32
	v4 = v1 & int32(v2)
	v5 = int64(uint64(v2)>>25) & i64(127) * i64(72340172838076673)
	t0 := int32(load32(m.memory[int64(uint32(v3))+8:]))
	v6 = t0
	t1 := int32(load32(m.memory[int64(uint32(v3))+4:]))
	v7 = t1
	v8 = i32(0)
	var _ int32
l4:
	{
		t3 := int64(load64(m.memory[uint32(v0+v4):]))
		v9 = t3
		v2 = v9 ^ v5
		v2 = (v2 ^ i64(-1)) & (v2 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
		{
		l2:
			{
				var p4 int32
				if v2 == 0 {
					p4 = 1
				}
				v3 = p4
				if v3 != 0 {
					goto l0
				}
				t5 := v7
				t6 := v6
				t7 := v0
				v10 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v2))))>>3) + v4) & v1
				v11 = t7 + (i32(0)-v10)*i32(24)
				t8 := int32(load32(m.memory[uint32(v11+i32(-20)):]))
				t9 := int32(load32(m.memory[uint32(v11+i32(-16)):]))
				t10 := m.fn544(t5, t6, t8, t9)
				if t10 != 0 {
					goto l1
				}
				v2 = (v2 + i64(-1)) & v2
				goto l2
			}
		l0:
			if v9&(v9<<1)&i64(-0x7f7f7f7f7f7f7f80) == 0 {
				t12 := v4
				v8 = v8 + i32(8)
				v4 = (t12 + v8) & v1
				goto l4
			}
		l1:
			p11 := v0 + (i32(0)-v10)*i32(24)
			if v3 != 0 {
				p11 = i32(0)
			}
			return p11
		}
	}
}
func (m *Module) fn645(v0, v1 int32, v2 int64, v3 int32) int32 {
	var v4 int32
	var v5 int64
	var v6, v7, v8 int32
	var v9 int64
	var v10, v11 int32
	v4 = v1 & int32(v2)
	v5 = int64(uint64(v2)>>25) & i64(127) * i64(72340172838076673)
	t0 := int32(load32(m.memory[int64(uint32(v3))+8:]))
	v6 = t0
	t1 := int32(load32(m.memory[int64(uint32(v3))+4:]))
	v7 = t1
	v8 = i32(0)
	var _ int32
l4:
	{
		t3 := int64(load64(m.memory[uint32(v0+v4):]))
		v9 = t3
		v2 = v9 ^ v5
		v2 = (v2 ^ i64(-1)) & (v2 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
		{
		l2:
			{
				var p4 int32
				if v2 == 0 {
					p4 = 1
				}
				v3 = p4
				if v3 != 0 {
					goto l0
				}
				t5 := v7
				t6 := v6
				t7 := v0
				v10 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v2))))>>3) + v4) & v1
				v11 = t7 + (i32(0)-v10)*i32(24)
				t8 := int32(load32(m.memory[uint32(v11+i32(-20)):]))
				t9 := int32(load32(m.memory[uint32(v11+i32(-16)):]))
				t10 := m.fn544(t5, t6, t8, t9)
				if t10 != 0 {
					goto l1
				}
				v2 = (v2 + i64(-1)) & v2
				goto l2
			}
		l0:
			if v9&(v9<<1)&i64(-0x7f7f7f7f7f7f7f80) == 0 {
				t12 := v4
				v8 = v8 + i32(8)
				v4 = (t12 + v8) & v1
				goto l4
			}
		l1:
			p11 := v0 + (i32(0)-v10)*i32(24)
			if v3 != 0 {
				p11 = i32(0)
			}
			return p11
		}
	}
}
func (m *Module) fn646(v0, v1 int32, v2 int64, v3 int32) int32 {
	var v4 int32
	var v5 int64
	var v6, v7, v8 int32
	var v9 int64
	var v10, v11 int32
	v4 = v1 & int32(v2)
	v5 = int64(uint64(v2)>>25) & i64(127) * i64(72340172838076673)
	t0 := int32(load32(m.memory[int64(uint32(v3))+8:]))
	v6 = t0
	t1 := int32(load32(m.memory[int64(uint32(v3))+4:]))
	v7 = t1
	v8 = i32(0)
	var _ int32
l4:
	{
		t3 := int64(load64(m.memory[uint32(v0+v4):]))
		v9 = t3
		v2 = v9 ^ v5
		v2 = (v2 ^ i64(-1)) & (v2 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
		{
		l2:
			{
				var p4 int32
				if v2 == 0 {
					p4 = 1
				}
				v3 = p4
				if v3 != 0 {
					goto l0
				}
				t5 := v7
				t6 := v6
				t7 := v0
				v10 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v2))))>>3) + v4) & v1
				v11 = t7 + (i32(0)-v10)*i32(36)
				t8 := int32(load32(m.memory[uint32(v11+i32(-32)):]))
				t9 := int32(load32(m.memory[uint32(v11+i32(-28)):]))
				t10 := m.fn544(t5, t6, t8, t9)
				if t10 != 0 {
					goto l1
				}
				v2 = (v2 + i64(-1)) & v2
				goto l2
			}
		l0:
			if v9&(v9<<1)&i64(-0x7f7f7f7f7f7f7f80) == 0 {
				t12 := v4
				v8 = v8 + i32(8)
				v4 = (t12 + v8) & v1
				goto l4
			}
		l1:
			p11 := v0 + (i32(0)-v10)*i32(36)
			if v3 != 0 {
				p11 = i32(0)
			}
			return p11
		}
	}
}
func (m *Module) fn647(v0, v1 int32, v2 int64, v3 int32) int32 {
	var v4 int32
	var v5 int64
	var v6, v7, v8 int32
	var v9 int64
	var v10, v11 int32
	v4 = v1 & int32(v2)
	v5 = int64(uint64(v2)>>25) & i64(127) * i64(72340172838076673)
	t0 := int32(load32(m.memory[int64(uint32(v3))+8:]))
	v6 = t0
	t1 := int32(load32(m.memory[int64(uint32(v3))+4:]))
	v7 = t1
	v8 = i32(0)
	var _ int32
l4:
	{
		t3 := int64(load64(m.memory[uint32(v0+v4):]))
		v9 = t3
		v2 = v9 ^ v5
		v2 = (v2 ^ i64(-1)) & (v2 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
		{
		l2:
			{
				var p4 int32
				if v2 == 0 {
					p4 = 1
				}
				v3 = p4
				if v3 != 0 {
					goto l0
				}
				t5 := v7
				t6 := v6
				t7 := v0
				v10 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v2))))>>3) + v4) & v1
				v11 = t7 + (i32(0)-v10)*i32(680)
				t8 := int32(load32(m.memory[uint32(v11+i32(-676)):]))
				t9 := int32(load32(m.memory[uint32(v11+i32(-672)):]))
				t10 := m.fn544(t5, t6, t8, t9)
				if t10 != 0 {
					goto l1
				}
				v2 = (v2 + i64(-1)) & v2
				goto l2
			}
		l0:
			if v9&(v9<<1)&i64(-0x7f7f7f7f7f7f7f80) == 0 {
				t12 := v4
				v8 = v8 + i32(8)
				v4 = (t12 + v8) & v1
				goto l4
			}
		l1:
			p11 := v0 + (i32(0)-v10)*i32(680)
			if v3 != 0 {
				p11 = i32(0)
			}
			return p11
		}
	}
}
func (m *Module) fn648(v0, v1 int32, v2 int64, v3, v4 int32) int32 {
	var v5 int32
	var v6 int64
	var v7 int32
	var v8 int64
	var v9, v10, v11 int32
	v5 = v1 & int32(v2)
	v6 = int64(uint64(v2)>>25) & i64(127) * i64(72340172838076673)
	v7 = i32(0)
	var _ int32
l4:
	{
		t1 := int64(load64(m.memory[uint32(v0+v5):]))
		v8 = t1
		v2 = v8 ^ v6
		v2 = (v2 ^ i64(-1)) & (v2 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
		{
		l2:
			{
				var p2 int32
				if v2 == 0 {
					p2 = 1
				}
				v9 = p2
				if v9 != 0 {
					goto l0
				}
				t3 := v3
				t4 := v4
				t5 := v0
				v10 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v2))))>>3) + v5) & v1
				v11 = t5 + (i32(0)-v10)*i32(28)
				t6 := int32(load32(m.memory[uint32(v11+i32(-24)):]))
				t7 := int32(load32(m.memory[uint32(v11+i32(-20)):]))
				t8 := m.fn15(t3, t4, t6, t7)
				if t8 != 0 {
					goto l1
				}
				v2 = (v2 + i64(-1)) & v2
				goto l2
			}
		l0:
			if v8&(v8<<1)&i64(-0x7f7f7f7f7f7f7f80) == 0 {
				t10 := v5
				v7 = v7 + i32(8)
				v5 = (t10 + v7) & v1
				goto l4
			}
		l1:
			p9 := v0 + (i32(0)-v10)*i32(28)
			if v9 != 0 {
				p9 = i32(0)
			}
			return p9
		}
	}
}
func (m *Module) fn649(v0, v1 int32) int32 {
	var v2, v3 int32
	var v4 int64
	var v5 int32
	var v6 int64
	var v7 int32
	var v8 int64
	var v9 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
		if t0 == 0 {
			goto l0
		}
		t1 := int64(load64(m.memory[int64(uint32(v0))+16:]))
		t2 := int64(load64(m.memory[int64(uint32(v0))+24:]))
		t3 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		v2 = t3
		t4 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		t5 := v2
		v3 = t4
		t6 := m.fn540(t1, t2, t5, v3)
		v4 = t6
		t7 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		v5 = t7
		v1 = v5 & int32(v4)
		v6 = int64(uint64(v4)>>25) & i64(127) * i64(72340172838076673)
		t8 := int32(load32(m.memory[uint32(v0):]))
		v0 = t8
		v7 = i32(0)
	l3:
		{
			t9 := int64(load64(m.memory[uint32(v0+v1):]))
			v8 = t9
			v4 = v8 ^ v6
			v4 = (v4 ^ i64(-1)) & (v4 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
		l4:
			{
				if v4 == 0 {
					if !(v8&(v8<<1)&i64(-0x7f7f7f7f7f7f7f80) == 0) {
						goto l0
					}
					t15 := v1
					v7 = v7 + i32(8)
					v1 = (t15 + v7) & v5
					goto l3
				}
				t10 := v2
				t11 := v3
				v9 = v0 + (i32(0)-(int32(uint32(int64(bits.TrailingZeros64(uint64(v4))))>>3)+v1)&v5)*i32(12)
				t12 := int32(load32(m.memory[uint32(v9+i32(-8)):]))
				t13 := int32(load32(m.memory[uint32(v9+i32(-4)):]))
				t14 := m.fn544(t10, t11, t12, t13)
				if t14 == 0 {
					v4 = (v4 + i64(-1)) & v4
					goto l4
				}
				return i32(1)
			}
		}
	}
l0:
	return i32(0)
}
func (m *Module) fn650(v0, v1 int32) int32 {
	var v2 int64
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
		if t0 != 0 {
			t1 := int64(load64(m.memory[int64(uint32(v0))+16:]))
			t2 := int64(load64(m.memory[int64(uint32(v0))+24:]))
			t3 := int32(load32(m.memory[uint32(v1):]))
			t4 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t5 := m.fn651(t1, t2, t3, t4)
			v2 = t5
			t6 := int32(load32(m.memory[uint32(v0):]))
			t7 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t8 := m.fn652(t6, t7, v2, v1)
			var p9 int32
			if t8 != i32(0) {
				p9 = 1
			}
			return p9
		}
		return i32(0)
	}
}
func (m *Module) fn651(v0, v1 int64, v2, v3 int32) int64 {
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
	m.fn172(v4, v2)
	m.fn172(v4, v3)
	t1 := m.fn174(v4)
	v1 = t1
	m.g0 = v4 + i32(64)
	return v1
}
func (m *Module) fn652(v0, v1 int32, v2 int64, v3 int32) int32 {
	var v4 int32
	var v5 int64
	var v6, v7, v8 int32
	var v9 int64
	var v10, v11 int32
	v4 = v1 & int32(v2)
	v5 = int64(uint64(v2)>>25) & i64(127) * i64(72340172838076673)
	t0 := int32(load32(m.memory[int64(uint32(v3))+4:]))
	v6 = t0
	t1 := int32(load32(m.memory[uint32(v3):]))
	v7 = t1
	v8 = i32(0)
	var _ int32
l5:
	{
		t3 := int64(load64(m.memory[uint32(v0+v4):]))
		v9 = t3
		v2 = v9 ^ v5
		v2 = (v2 ^ i64(-1)) & (v2 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
		{
		l3:
			{
				var p4 int32
				if v2 == 0 {
					p4 = 1
				}
				v3 = p4
				if v3 != 0 {
					goto l0
				}
				{
					t5 := v7
					t6 := v0
					v10 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v2))))>>3) + v4) & v1
					v11 = t6 - v10<<4
					t7 := int32(load32(m.memory[uint32(v11+i32(-16)):]))
					if t5 != t7 {
						goto l1
					}
					t8 := int32(load32(m.memory[uint32(v11+i32(-12)):]))
					if v6 == t8 {
						goto l2
					}
				}
			l1:
				v2 = (v2 + i64(-1)) & v2
				goto l3
			}
		l0:
			if v9&(v9<<1)&i64(-0x7f7f7f7f7f7f7f80) == 0 {
				t10 := v4
				v8 = v8 + i32(8)
				v4 = (t10 + v8) & v1
				goto l5
			}
		l2:
			p9 := v0 - v10<<4
			if v3 != 0 {
				p9 = i32(0)
			}
			return p9
		}
	}
}
func (m *Module) fn653(v0, v1, v2 int32) {
	var v3 int64
	var v4, v5, v6, v7, v8 int32
	t0 := int64(load64(m.memory[int64(uint32(v1))+16:]))
	t1 := int64(load64(m.memory[int64(uint32(v1))+24:]))
	t2 := int32(load32(m.memory[uint32(v2):]))
	t3 := int32(load32(m.memory[int64(uint32(v2))+4:]))
	t4 := m.fn651(t0, t1, t2, t3)
	v3 = t4
	{
		{
			t5 := int32(load32(m.memory[uint32(v1):]))
			v4 = t5
			t6 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t7 := v4
			v5 = t6
			t8 := m.fn652(t7, v5, v3, v2)
			v2 = t8
			if v2 != 0 {
				goto l0
			}
			v1 = i32(0)
			goto l1
		}
	l0:
		v6 = i32(128)
		{
			t9 := v4
			v7 = (v4 - v2) >> 4
			v8 = t9 + v7
			t10 := int64(load64(m.memory[uint32(v8):]))
			v3 = t10
			t11 := int32(uint32(int64(bits.TrailingZeros64(uint64(v3&(v3<<1)&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
			v4 = v4 + (v7+i32(-8))&v5
			t12 := int64(load64(m.memory[uint32(v4):]))
			v3 = t12
			if uint32(t11+int32(uint32(int64(bits.LeadingZeros64(uint64(v3&(v3<<1)&i64(-0x7f7f7f7f7f7f7f80)))))>>3)) > uint32(i32(7)) {
				goto l2
			}
			t13 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			store32(m.memory[int64(uint32(v1))+8:], uint32(t13+i32(1)))
			v6 = i32(255)
		}
	l2:
		m.memory[uint32(v8)] = byte(v6)
		m.memory[uint32(v4+i32(8))] = byte(v6)
		t14 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		store32(m.memory[int64(uint32(v1))+12:], uint32(t14+i32(-1)))
		t15 := int64(load64(m.memory[uint32(v2+i32(-8)):]))
		store64(m.memory[int64(uint32(v0))+4:], uint64(t15))
		v1 = i32(1)
	}
l1:
	store32(m.memory[uint32(v0):], uint32(v1))
}
func (m *Module) fn654(v0, v1 int32) int32 {
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
				t19 := m.fn655(t17, t18, v10)
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
	m.fn241(v0, v2+i32(16), i32(99), i32(20))
l7:
	v5 = i32(-1)
l2:
	m.g0 = v2 + i32(64)
	return v5
}
func (m *Module) fn655(v0, v1, v2 int32) int64 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(t0):]))
	v0 = t1
	t2 := int64(load64(m.memory[uint32(v0):]))
	t3 := int64(load64(m.memory[uint32(v0+i32(8)):]))
	t4 := int32(load32(m.memory[uint32(v1):]))
	v0 = t4 + (i32(0)-v2)*i32(20)
	t5 := int32(load32(m.memory[uint32(v0+i32(-16)):]))
	t6 := int32(load32(m.memory[uint32(v0+i32(-12)):]))
	t7 := m.fn524(t2, t3, t5, t6)
	return t7
}
func (m *Module) fn656(v0, v1 int32) int32 {
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
		m.fn237(t5, i32(416), p6)
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
		store64(m.memory[int64(uint32(v2))+24:], uint64(i64(0x8000001a0)))
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
				t19 := m.fn657(t17, t18, v10)
				v11 = t19
				t20 := m.fn26(t15, t16, v11)
				v12 = t20
				t21 := t14 + v12
				v13 = int32(uint32(int32(v11)) >> 25)
				m.memory[uint32(t21)] = byte(v13)
				m.memory[uint32(v6+v5&(v12+i32(-8))+i32(8))] = byte(v13)
				t22 := int32(load32(m.memory[uint32(v0):]))
				memory_copy(m.memory, uint32(v6+(v12^i32(-1))*i32(416)), uint32(t22+(v10^i32(-1))*i32(416)), uint32(i32(416)))
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
	m.fn241(v0, v2+i32(16), i32(100), i32(416))
l7:
	v5 = i32(-1)
l2:
	m.g0 = v2 + i32(64)
	return v5
}
func (m *Module) fn657(v0, v1, v2 int32) int64 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(t0):]))
	v0 = t1
	t2 := int64(load64(m.memory[uint32(v0):]))
	t3 := int64(load64(m.memory[uint32(v0+i32(8)):]))
	t4 := int32(load32(m.memory[uint32(v1):]))
	v0 = t4 + (i32(0)-v2)*i32(416)
	t5 := int32(load32(m.memory[uint32(v0+i32(-412)):]))
	t6 := int32(load32(m.memory[uint32(v0+i32(-408)):]))
	t7 := m.fn540(t2, t3, t5, t6)
	return t7
}
func (m *Module) fn658(v0, v1 int32) int32 {
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
				t19 := m.fn659(t17, t18, v10)
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
	m.fn241(v0, v2+i32(16), i32(101), i32(20))
l7:
	v5 = i32(-1)
l2:
	m.g0 = v2 + i32(64)
	return v5
}
func (m *Module) fn659(v0, v1, v2 int32) int64 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(t0):]))
	v0 = t1
	t2 := int64(load64(m.memory[uint32(v0):]))
	t3 := int64(load64(m.memory[uint32(v0+i32(8)):]))
	t4 := int32(load32(m.memory[uint32(v1):]))
	v0 = t4 + (i32(0)-v2)*i32(20)
	t5 := int32(load32(m.memory[uint32(v0+i32(-16)):]))
	t6 := int32(load32(m.memory[uint32(v0+i32(-12)):]))
	t7 := m.fn540(t2, t3, t5, t6)
	return t7
}
func (m *Module) fn660(v0, v1 int32) int32 {
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
				t19 := m.fn661(t17, t18, v10)
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
	m.fn241(v0, v2+i32(16), i32(102), i32(24))
l7:
	v5 = i32(-1)
l2:
	m.g0 = v2 + i32(64)
	return v5
}
func (m *Module) fn661(v0, v1, v2 int32) int64 {
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
func (m *Module) fn662(v0, v1 int32) int32 {
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
				t19 := m.fn663(t17, t18, v10)
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
	m.fn241(v0, v2+i32(16), i32(103), i32(24))
l7:
	v5 = i32(-1)
l2:
	m.g0 = v2 + i32(64)
	return v5
}
func (m *Module) fn663(v0, v1, v2 int32) int64 {
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
func (m *Module) fn664(v0, v1 int32) int32 {
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
	m.fn241(v0, v2+i32(16), i32(104), i32(16))
l7:
	v5 = i32(-1)
l2:
	m.g0 = v2 + i32(64)
	return v5
}
func (m *Module) fn665(v0, v1, v2 int32) int64 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(t0):]))
	v0 = t1
	t2 := int64(load64(m.memory[uint32(v0):]))
	t3 := int64(load64(m.memory[uint32(v0+i32(8)):]))
	t4 := int32(load32(m.memory[uint32(v1):]))
	v2 = t4 - v2<<4
	t5 := int32(load32(m.memory[uint32(v2+i32(-12)):]))
	t6 := int32(load32(m.memory[uint32(v2+i32(-8)):]))
	t7 := m.fn540(t2, t3, t5, t6)
	return t7
}
func (m *Module) fn666(v0, v1 int32) int32 {
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
				t19 := m.fn667(t17, t18, v10)
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
	m.fn241(v0, v2+i32(16), i32(105), i32(16))
l7:
	v5 = i32(-1)
l2:
	m.g0 = v2 + i32(64)
	return v5
}
func (m *Module) fn667(v0, v1, v2 int32) int64 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(t0):]))
	v0 = t1
	t2 := int64(load64(m.memory[uint32(v0):]))
	t3 := int64(load64(m.memory[uint32(v0+i32(8)):]))
	t4 := int32(load32(m.memory[uint32(v1):]))
	v2 = t4 - v2<<4
	t5 := int32(load32(m.memory[uint32(v2+i32(-12)):]))
	t6 := int32(load32(m.memory[uint32(v2+i32(-8)):]))
	t7 := m.fn540(t2, t3, t5, t6)
	return t7
}
func (m *Module) fn668(v0, v1 int32) int32 {
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
				t19 := m.fn669(t17, t18, v10)
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
	m.fn241(v0, v2+i32(16), i32(106), i32(36))
l7:
	v5 = i32(-1)
l2:
	m.g0 = v2 + i32(64)
	return v5
}
func (m *Module) fn669(v0, v1, v2 int32) int64 {
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
func (m *Module) fn670(v0, v1 int32) int32 {
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
		m.fn237(t5, i32(680), p6)
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
		store64(m.memory[int64(uint32(v2))+24:], uint64(i64(0x8000002a8)))
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
				t19 := m.fn671(t17, t18, v10)
				v11 = t19
				t20 := m.fn26(t15, t16, v11)
				v12 = t20
				t21 := t14 + v12
				v13 = int32(uint32(int32(v11)) >> 25)
				m.memory[uint32(t21)] = byte(v13)
				m.memory[uint32(v6+v5&(v12+i32(-8))+i32(8))] = byte(v13)
				t22 := int32(load32(m.memory[uint32(v0):]))
				memory_copy(m.memory, uint32(v6+(v12^i32(-1))*i32(680)), uint32(t22+(v10^i32(-1))*i32(680)), uint32(i32(680)))
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
	m.fn241(v0, v2+i32(16), i32(107), i32(680))
l7:
	v5 = i32(-1)
l2:
	m.g0 = v2 + i32(64)
	return v5
}
func (m *Module) fn671(v0, v1, v2 int32) int64 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(t0):]))
	v0 = t1
	t2 := int64(load64(m.memory[uint32(v0):]))
	t3 := int64(load64(m.memory[uint32(v0+i32(8)):]))
	t4 := int32(load32(m.memory[uint32(v1):]))
	v0 = t4 + (i32(0)-v2)*i32(680)
	t5 := int32(load32(m.memory[uint32(v0+i32(-676)):]))
	t6 := int32(load32(m.memory[uint32(v0+i32(-672)):]))
	t7 := m.fn540(t2, t3, t5, t6)
	return t7
}
