package core

import (
	"math"
	"math/bits"
)

func (m *Module) fn807(v0, v1 int32) {
	var v2, v3 int32
	v2 = i32(8)
	{
		t0 := int32(m.memory[uint32(v1)])
		switch t0 {
		default:
			t1 := int64(load64(m.memory[int64(uint32(v1))+8:]))
			store64(m.memory[int64(uint32(v0))+8:], uint64(t1))
			m.memory[uint32(v0)] = byte(i32(0))
			return
		case 1:
			t2 := math.Float64frombits(load64(m.memory[int64(uint32(v1))+8:]))
			store64(m.memory[int64(uint32(v0))+8:], math.Float64bits(t2))
			m.memory[uint32(v0)] = byte(i32(1))
			return
		case 2:
			t3 := int32(load32(m.memory[int64(uint32(v1))+12:]))
			store32(m.memory[int64(uint32(v0))+12:], uint32(t3))
			t4 := int64(load64(m.memory[int64(uint32(v1))+4:]))
			store64(m.memory[int64(uint32(v0))+4:], uint64(t4))
			m.memory[uint32(v0)] = byte(i32(2))
			return
		case 3:
			t5 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			v2 = t5
			if v2 <= i32(-1) {
				m.fn9()
				panic("unreachable")
			}
			if v2 != 0 {
				t6 := int32(load32(m.memory[int64(uint32(v1))+4:]))
				v3 = t6
				t7 := m.fn5(v2)
				v1 = t7
				if v1 != 0 {
					goto l13
				}
				m.fn10(i32(1), v2)
				panic("unreachable")
			}
			v1 = i32(1)
			goto l12
		case 4:
			t8 := int32(m.memory[int64(uint32(v1))+1])
			m.memory[int64(uint32(v0))+1] = byte(t8)
			m.memory[uint32(v0)] = byte(i32(3))
			return
		case 5:
			t9 := int64(load64(m.memory[int64(uint32(v1))+16:]))
			store64(m.memory[int64(uint32(v0))+16:], uint64(t9))
			t10 := int64(load64(m.memory[int64(uint32(v1))+8:]))
			store64(m.memory[int64(uint32(v0))+8:], uint64(t10))
			m.memory[uint32(v0)] = byte(i32(4))
			return
		case 6:
			t11 := int32(load32(m.memory[int64(uint32(v1))+12:]))
			store32(m.memory[int64(uint32(v0))+12:], uint32(t11))
			t12 := int64(load64(m.memory[int64(uint32(v1))+4:]))
			store64(m.memory[int64(uint32(v0))+4:], uint64(t12))
			m.memory[uint32(v0)] = byte(i32(5))
			return
		case 7:
			t13 := int32(load32(m.memory[int64(uint32(v1))+12:]))
			store32(m.memory[int64(uint32(v0))+12:], uint32(t13))
			t14 := int64(load64(m.memory[int64(uint32(v1))+4:]))
			store64(m.memory[int64(uint32(v0))+4:], uint64(t14))
			m.memory[uint32(v0)] = byte(i32(6))
			return
		case 8:
			t15 := int32(m.memory[int64(uint32(v1))+1])
			m.memory[int64(uint32(v0))+1] = byte(t15)
			v2 = i32(7)
			fallthrough
		case 9:
			m.memory[uint32(v0)] = byte(v2)
			return
		}
	}
l13:
	if v2 == 0 {
		goto l12
	}
	memory_copy(m.memory, uint32(v1), uint32(v3), uint32(v2))
l12:
	store32(m.memory[int64(uint32(v0))+12:], uint32(v2))
	store32(m.memory[int64(uint32(v0))+8:], uint32(v1))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v2))
	m.memory[uint32(v0)] = byte(i32(2))
}
func (m *Module) fn808(v0 int32) {
	var v1 int32
	var v2 int64
	t0 := m.g0
	v1 = t0 - i32(32)
	m.g0 = v1
	store32(m.memory[int64(uint32(v1))+8:], uint32(v0))
	store32(m.memory[int64(uint32(v1))+12:], uint32(i32(4)))
	t1 := v1
	v2 = int64(uint32(i32(2))) << 32
	store64(m.memory[int64(uint32(t1))+24:], uint64(v2|int64(uint32(v1+i32(12)))))
	store64(m.memory[int64(uint32(v1))+16:], uint64(v2|int64(uint32(v1+i32(8)))))
	m.fn28(i32(1066357), v1+i32(16), i32(1089976))
	panic("unreachable")
}
func (m *Module) fn809(v0, v1 int32) {
	var v2, v3, v4, v5, v6, v7, v8, v9, v10 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	{
		{
			if uint32(v0) < uint32(i32(26)) {
				goto l0
			}
			v3 = i32(0)
			store32(m.memory[int64(uint32(v2))+12:], uint32(i32(0)))
			store64(m.memory[int64(uint32(v2))+4:], uint64(i64(0x100000000)))
			v4 = i32(1)
		l2:
			{
				t1 := int32(uint32(v0) / uint32(i32(26)))
				t2 := v0
				v5 = t1
				v6 = t2 - v5*i32(26)
				{
					t3 := int32(load32(m.memory[int64(uint32(v2))+4:]))
					if v3 != t3 {
						goto l1
					}
					m.fn647(v2+i32(4), v3, i32(1), i32(1), i32(1))
					t4 := int32(load32(m.memory[int64(uint32(v2))+8:]))
					v4 = t4
				}
			l1:
				m.memory[uint32(v4+v3)] = byte(v6 + i32(65))
				t5 := v2
				v3 = v3 + i32(1)
				store32(m.memory[int64(uint32(t5))+12:], uint32(v3))
				var p6 int32
				if uint32(v0) > uint32(i32(675)) {
					p6 = 1
				}
				v6 = p6
				v0 = v5
				if v6 != 0 {
					goto l2
				}
			}
			t7 := int32(load32(m.memory[int64(uint32(v2))+8:]))
			v7 = t7
			{
				t8 := int32(uint32(v3) >> 2)
				var p9 int32
				if v3&i32(3) != i32(0) {
					p9 = 1
				}
				v0 = t8 + p9
				t10 := int32(load32(m.memory[uint32(v1):]))
				t11 := int32(load32(m.memory[int64(uint32(v1))+8:]))
				t12 := v0
				v5 = t11
				if uint32(t12) <= uint32(t10-v5) {
					goto l3
				}
				m.fn647(v1, v5, v0, i32(1), i32(1))
			}
		l3:
			v5 = v7 + v3
		l17:
			{
				{
					{
						v6 = v5 + i32(-1)
						t13 := int32(int8(m.memory[uint32(v6)]))
						v0 = t13
						if v0 > i32(-1) {
							t15 := int32(load32(m.memory[int64(uint32(v1))+8:]))
							v3 = t15
							v8 = i32(1)
							v5 = v6
							v6 = i32(1)
							goto l7
						}
						v4 = v5 + i32(-2)
						t14 := int32(m.memory[uint32(v4)])
						v3 = t14
						v6 = int32(int8(v3))
						if v6 < i32(-64) {
							goto l5
						}
						v5 = v3 & i32(31)
						goto l6
					}
				l5:
					{
						{
							v4 = v5 + i32(-3)
							t16 := int32(m.memory[uint32(v4)])
							v3 = t16
							v8 = int32(int8(v3))
							if v8 <= i32(-65) {
								goto l8
							}
							v3 = v3 & i32(15)
							goto l9
						}
					l8:
						v4 = v5 + i32(-4)
						t17 := int32(m.memory[uint32(v4)])
						v3 = t17&i32(7)<<6 | v8&i32(63)
					}
				l9:
					v5 = v3<<6 | v6&i32(63)
				l6:
					v0 = v5<<6 | v0&i32(63)
					t18 := int32(load32(m.memory[int64(uint32(v1))+8:]))
					v3 = t18
					v8 = i32(1)
					if uint32(v5) >= uint32(i32(2)) {
						goto l10
					}
					v5 = v4
					v6 = i32(1)
					goto l7
				l10:
					v6 = i32(2)
					v8 = i32(0)
					{
						if uint32(v5) < uint32(i32(32)) {
							goto l11
						}
						p19 := i32(4)
						if uint32(v5) < uint32(i32(1024)) {
							p19 = i32(3)
						}
						v6 = p19
					}
				l11:
					v5 = v4
				}
			l7:
				{
					t20 := int32(load32(m.memory[uint32(v1):]))
					if uint32(v6) <= uint32(t20-v3) {
						goto l12
					}
					m.fn647(v1, v3, v6, i32(1), i32(1))
				}
			l12:
				t21 := int32(load32(m.memory[int64(uint32(v1))+4:]))
				v4 = t21 + v3
				if v8 != 0 {
					goto l13
				}
				v8 = v0&i32(63) | i32(-128)
				v9 = int32(uint32(v0) >> 6)
				if uint32(v0) >= uint32(i32(2048)) {
					v10 = int32(uint32(v0) >> 12)
					v9 = v9&i32(63) | i32(-128)
					if uint32(v0) > uint32(i32(0xffff)) {
						m.memory[int64(uint32(v4))+3] = byte(v8)
						m.memory[int64(uint32(v4))+2] = byte(v9)
						m.memory[int64(uint32(v4))+1] = byte(v10&i32(63) | i32(-128))
						m.memory[uint32(v4)] = byte(int32(uint32(v0)>>18) | i32(-16))
						goto l15
					}
					m.memory[int64(uint32(v4))+2] = byte(v8)
					m.memory[int64(uint32(v4))+1] = byte(v9)
					m.memory[uint32(v4)] = byte(v10 | i32(224))
					goto l15
				}
				m.memory[int64(uint32(v4))+1] = byte(v8)
				m.memory[uint32(v4)] = byte(v9 | i32(192))
				goto l15
			l13:
				m.memory[uint32(v4)] = byte(v0)
			l15:
				store32(m.memory[int64(uint32(v1))+8:], uint32(v6+v3))
				if v7 != v5 {
					goto l17
				}
			}
			t22 := int32(load32(m.memory[int64(uint32(v2))+4:]))
			v0 = t22
			if v0 == 0 {
				goto l18
			}
			{
				t23 := int32(load32(m.memory[uint32(v7+i32(-4)):]))
				v1 = t23
				v3 = v1 & i32(-8)
				t24 := v3
				v1 = v1 & i32(3)
				p25 := i32(8)
				if v1 != 0 {
					p25 = i32(4)
				}
				if uint32(t24) < uint32(p25+v0) {
					m.fn3(i32(1273840), i32(46), i32(1273888))
					panic("unreachable")
				}
				if v1 == 0 {
					goto l20
				}
				if uint32(v3) > uint32(v0+i32(39)) {
					m.fn3(i32(1273904), i32(46), i32(1273952))
					panic("unreachable")
				}
			l20:
				m.fn1(v7)
				goto l18
			}
		}
	l0:
		{
			t26 := int32(load32(m.memory[uint32(v1):]))
			t27 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			v3 = t27
			if t26 != v3 {
				goto l22
			}
			m.fn647(v1, v3, i32(1), i32(1), i32(1))
		}
	l22:
		store32(m.memory[int64(uint32(v1))+8:], uint32(v3+i32(1)))
		t28 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		m.memory[uint32(t28+v3)] = byte(v0 + i32(65))
	}
l18:
	m.g0 = v2 + i32(16)
}
func (m *Module) fn810(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v1):]))
	t1 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	t2 := int32(load32(m.memory[int64(uint32(t1))+12:]))
	t3 := m.t0[uint(t2)].(func(int32, int32, int32) int32)(t0, i32(1274653), i32(5))
	return t3
}
func (m *Module) fn811(v0, v1, v2 int32) int32 {
	var v3 int32
	{
		{
			t0 := int32(load32(m.memory[uint32(v0):]))
			t1 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			t2 := v2
			v3 = t1
			if uint32(t2) <= uint32(t0-v3) {
				goto l0
			}
			m.fn647(v0, v3, v2, i32(1), i32(1))
			t3 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			v3 = t3
			goto l1
		}
	l0:
		if v2 == 0 {
			goto l2
		}
	l1:
		if v2 == 0 {
			goto l2
		}
		t4 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		memory_copy(m.memory, uint32(t4+v3), uint32(v1), uint32(v2))
	}
l2:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v3+v2))
	return i32(0)
}
func (m *Module) fn812(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7 int32
	t0 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	v2 = t0
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
		p1 := i32(4)
		if uint32(v1) < uint32(i32(65536)) {
			p1 = i32(3)
		}
		v3 = p1
	}
l1:
	{
		t2 := int32(load32(m.memory[uint32(v0):]))
		if uint32(v3) <= uint32(t2-v2) {
			goto l3
		}
		m.fn647(v0, v2, v3, i32(1), i32(1))
	}
l3:
	t3 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	v4 = t3 + v2
	if uint32(v1) < uint32(i32(128)) {
		goto l4
	}
	v5 = v1&i32(63) | i32(-128)
	v6 = int32(uint32(v1) >> 6)
	if uint32(v1) >= uint32(i32(2048)) {
		v7 = int32(uint32(v1) >> 12)
		v6 = v6&i32(63) | i32(-128)
		if uint32(v1) > uint32(i32(0xffff)) {
			m.memory[int64(uint32(v4))+3] = byte(v5)
			m.memory[int64(uint32(v4))+2] = byte(v6)
			m.memory[int64(uint32(v4))+1] = byte(v7&i32(63) | i32(-128))
			m.memory[uint32(v4)] = byte(int32(uint32(v1)>>18) | i32(-16))
			goto l6
		}
		m.memory[int64(uint32(v4))+2] = byte(v5)
		m.memory[int64(uint32(v4))+1] = byte(v6)
		m.memory[uint32(v4)] = byte(v7 | i32(224))
		goto l6
	}
	m.memory[int64(uint32(v4))+1] = byte(v5)
	m.memory[uint32(v4)] = byte(v6 | i32(192))
	goto l6
l4:
	m.memory[uint32(v4)] = byte(v1)
l6:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v3+v2))
	return i32(0)
}
func (m *Module) fn813(v0, v1, v2 int32) int32 {
	t0 := m.fn45(v0, i32(1080776), v1, v2)
	return t0
}
func (m *Module) fn814(v0 int32) {
	var v1, v2, v3, v4 int32
	{
		t0 := int32(load32(m.memory[uint32(v0):]))
		v1 = t0
		p1 := i32(2)
		if uint32(v1) > uint32(i32(-0x7ffffff2)) {
			p1 = v1 + i32(0x7ffffff1)
		}
		switch p1 {
		default:
			return
		case 0:
			t2 := int32(m.memory[int64(uint32(v0))+4])
			if t2 != i32(3) {
				return
			}
			t3 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			v0 = t3
			t4 := int32(load32(m.memory[uint32(v0):]))
			v1 = t4
			{
				t5 := int32(load32(m.memory[uint32(v0+i32(4)):]))
				v2 = t5
				t6 := int32(load32(m.memory[uint32(v2):]))
				v3 = t6
				if v3 == 0 {
					goto l8
				}
				m.t0[uint(v3)].(func(int32))(v1)
			}
		l8:
			{
				t7 := int32(load32(m.memory[int64(uint32(v2))+4:]))
				v2 = t7
				if v2 == 0 {
					goto l9
				}
				t8 := int32(load32(m.memory[uint32(v1+i32(-4)):]))
				v3 = t8
				v4 = v3 & i32(-8)
				t9 := v4
				v3 = v3 & i32(3)
				p10 := i32(8)
				if v3 != 0 {
					p10 = i32(4)
				}
				if uint32(t9) < uint32(p10+v2) {
					m.fn3(i32(1273840), i32(46), i32(1273888))
					panic("unreachable")
				}
				if v3 == 0 {
					goto l11
				}
				if uint32(v4) > uint32(v2+i32(39)) {
					m.fn3(i32(1273904), i32(46), i32(1273952))
					panic("unreachable")
				}
			l11:
				m.fn1(v1)
			}
		l9:
			t11 := int32(load32(m.memory[uint32(v0+i32(-4)):]))
			v1 = t11
			v2 = v1 & i32(-8)
			t12 := v2
			v1 = v1 & i32(3)
			p13 := i32(20)
			if v1 != 0 {
				p13 = i32(16)
			}
			if uint32(t12) < uint32(p13) {
				m.fn3(i32(1273840), i32(46), i32(1273888))
				panic("unreachable")
			}
			if v1 == 0 {
				goto l14
			}
			if uint32(v2) >= uint32(i32(52)) {
				m.fn3(i32(1273904), i32(46), i32(1273952))
				panic("unreachable")
			}
		l14:
			m.fn1(v0)
			return
		case 1:
			m.fn607(v0 + i32(4))
			return
		case 2:
			m.fn815(v0)
			return
		case 4:
			m.fn604(v0 + i32(4))
			return
		case 6:
			t14 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			v1 = t14
			if v1 == 0 {
				return
			}
			t15 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			v2 = t15
			t16 := int32(load32(m.memory[uint32(v2+i32(-4)):]))
			v0 = t16
			v3 = v0 & i32(-8)
			t17 := v3
			v0 = v0 & i32(3)
			p18 := i32(8)
			if v0 != 0 {
				p18 = i32(4)
			}
			if uint32(t17) < uint32(p18+v1) {
				m.fn3(i32(1273840), i32(46), i32(1273888))
				panic("unreachable")
			}
			if v0 == 0 {
				goto l17
			}
			if uint32(v3) > uint32(v1+i32(39)) {
				m.fn3(i32(1273904), i32(46), i32(1273952))
				panic("unreachable")
			}
		l17:
			m.fn1(v2)
			return
		case 15:
			t19 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			v1 = t19
			if v1 == 0 {
				return
			}
			t20 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			v2 = t20
			t21 := int32(load32(m.memory[uint32(v2+i32(-4)):]))
			v0 = t21
			v3 = v0 & i32(-8)
			t22 := v3
			v0 = v0 & i32(3)
			p23 := i32(8)
			if v0 != 0 {
				p23 = i32(4)
			}
			if uint32(t22) < uint32(p23+v1) {
				m.fn3(i32(1273840), i32(46), i32(1273888))
				panic("unreachable")
			}
			if v0 == 0 {
				goto l20
			}
			if uint32(v3) > uint32(v1+i32(39)) {
				m.fn3(i32(1273904), i32(46), i32(1273952))
				panic("unreachable")
			}
		l20:
			m.fn1(v2)
			return
		case 17:
			t24 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			v1 = t24
			if v1 == 0 {
				return
			}
			t25 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			v2 = t25
			t26 := int32(load32(m.memory[uint32(v2+i32(-4)):]))
			v0 = t26
			v3 = v0 & i32(-8)
			t27 := v3
			v0 = v0 & i32(3)
			p28 := i32(8)
			if v0 != 0 {
				p28 = i32(4)
			}
			if uint32(t27) < uint32(p28+v1) {
				m.fn3(i32(1273840), i32(46), i32(1273888))
				panic("unreachable")
			}
			if v0 == 0 {
				goto l23
			}
			if uint32(v3) > uint32(v1+i32(39)) {
				m.fn3(i32(1273904), i32(46), i32(1273952))
				panic("unreachable")
			}
		l23:
			m.fn1(v2)
		}
	}
}
func (m *Module) fn815(v0 int32) {
	var v1, v2, v3 int32
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
			switch t2 {
			default:
				return
			case 0, 1, 2, 3, 4:
				t3 := int32(load32(m.memory[int64(uint32(v0))+8:]))
				v1 = t3
				if v1 == 0 {
					return
				}
				t4 := int32(load32(m.memory[int64(uint32(v0))+12:]))
				v2 = t4
				t5 := int32(load32(m.memory[uint32(v2+i32(-4)):]))
				v0 = t5
				v3 = v0 & i32(-8)
				t6 := v3
				v0 = v0 & i32(3)
				p7 := i32(8)
				if v0 != 0 {
					p7 = i32(4)
				}
				if uint32(t6) < uint32(p7+v1) {
					m.fn3(i32(1273840), i32(46), i32(1273888))
					panic("unreachable")
				}
				if v0 == 0 {
					goto l7
				}
				if uint32(v3) > uint32(v1+i32(39)) {
					m.fn3(i32(1273904), i32(46), i32(1273952))
					panic("unreachable")
				}
			l7:
				m.fn1(v2)
				return
			}
		case 0:
			t8 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			v1 = t8
			t9 := int32(load32(m.memory[uint32(v1):]))
			t10 := v1
			v1 = t9
			store32(m.memory[uint32(t10):], uint32(v1+i32(-1)))
			if v1 != i32(1) {
				return
			}
			t11 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			m.fn238(t11)
			return
		case 2:
			p12 := i32(5)
			if v1 < i32(0) {
				p12 = v1 ^ i32(-0x80000000)
			}
			switch p12 {
			default:
				return
			case 0:
				t13 := int32(load32(m.memory[int64(uint32(v0))+4:]))
				v1 = t13
				if v1 <= i32(0) {
					return
				}
				v2 = i32(8)
				goto l13
			case 3:
				t14 := int32(load32(m.memory[int64(uint32(v0))+4:]))
				v1 = t14
				if v1 == 0 {
					return
				}
				v2 = i32(8)
				goto l13
			case 4:
				t15 := int32(load32(m.memory[int64(uint32(v0))+4:]))
				v1 = t15
				if v1 == 0 {
					return
				}
				v2 = i32(8)
				goto l13
			case 5:
				{
					if v1 == 0 {
						goto l14
					}
					t16 := int32(load32(m.memory[int64(uint32(v0))+4:]))
					m.fn18(t16, v1, i32(1))
				}
			l14:
				t17 := int32(load32(m.memory[int64(uint32(v0))+12:]))
				v1 = t17
				if v1 == 0 {
					return
				}
				v2 = i32(16)
				goto l13
			}
		case 5:
			t18 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			v1 = t18
			if v1 < i32(1) {
				return
			}
			t19 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			v2 = t19
			t20 := int32(load32(m.memory[uint32(v2+i32(-4)):]))
			v0 = t20
			v3 = v0 & i32(-8)
			t21 := v3
			v0 = v0 & i32(3)
			p22 := i32(8)
			if v0 != 0 {
				p22 = i32(4)
			}
			if uint32(t21) < uint32(p22+v1) {
				m.fn3(i32(1273840), i32(46), i32(1273888))
				panic("unreachable")
			}
			if v0 == 0 {
				goto l16
			}
			if uint32(v3) > uint32(v1+i32(39)) {
				m.fn3(i32(1273904), i32(46), i32(1273952))
				panic("unreachable")
			}
		l16:
			m.fn1(v2)
			return
		}
	}
l13:
	{
		t23 := int32(load32(m.memory[uint32(v0+v2):]))
		v2 = t23
		t24 := int32(load32(m.memory[uint32(v2+i32(-4)):]))
		v0 = t24
		v3 = v0 & i32(-8)
		t25 := v3
		v0 = v0 & i32(3)
		p26 := i32(8)
		if v0 != 0 {
			p26 = i32(4)
		}
		if uint32(t25) < uint32(p26+v1) {
			m.fn3(i32(1273840), i32(46), i32(1273888))
			panic("unreachable")
		}
		if v0 == 0 {
			goto l19
		}
		if uint32(v3) > uint32(v1+i32(39)) {
			m.fn3(i32(1273904), i32(46), i32(1273952))
			panic("unreachable")
		}
	l19:
		m.fn1(v2)
		return
	}
}
func (m *Module) fn816(v0, v1, v2 int32) {
	var v3 int32
	{
		{
			switch v2 + i32(-4) {
			case 3:
				t0 := int32(load32(m.memory[uint32(v1):]))
				t1 := t0 ^ i32(1447642147)
				v3 = v1 + i32(3)
				t2 := int32(load32(m.memory[uint32(v3):]))
				if t1|(t2^i32(556805974)) != 0 {
					goto l5
				}
				store32(m.memory[uint32(v0):], uint32(i32(-1)))
				m.memory[int64(uint32(v0))+4] = byte(i32(0))
				return
			case 0:
				t3 := int32(load32(m.memory[uint32(v1):]))
				if t3 != i32(1093619235) {
					goto l6
				}
				store32(m.memory[uint32(v0):], uint32(i32(-1)))
				m.memory[int64(uint32(v0))+4] = byte(i32(1))
				return
			case 2:
				t4 := int32(load32(m.memory[uint32(v1):]))
				t5 := t4 ^ i32(1296125475)
				v3 = v1 + i32(4)
				t6 := int32(load16(m.memory[uint32(v3):]))
				if t5|(t6^i32(16197)) != 0 {
					t7 := int32(load32(m.memory[uint32(v1):]))
					t8 := int32(load16(m.memory[uint32(v3):]))
					if t7^i32(0x4c554e23)|(t8^i32(8524)) != 0 {
						goto l6
					}
					store32(m.memory[uint32(v0):], uint32(i32(-1)))
					m.memory[int64(uint32(v0))+4] = byte(i32(3))
					return
				}
				store32(m.memory[uint32(v0):], uint32(i32(-1)))
				m.memory[int64(uint32(v0))+4] = byte(i32(2))
				return
			case 1:
				t9 := int32(load32(m.memory[uint32(v1):]))
				t10 := t9 ^ i32(0x4d554e23)
				v3 = v1 + i32(4)
				t11 := int32(m.memory[uint32(v3)])
				if t10|(t11^i32(33)) != 0 {
					t12 := int32(load32(m.memory[uint32(v1):]))
					t13 := int32(m.memory[uint32(v3)])
					if t12^i32(0x46455223)|(t13^i32(33)) != 0 {
						goto l6
					}
					store32(m.memory[uint32(v0):], uint32(i32(-1)))
					m.memory[int64(uint32(v0))+4] = byte(i32(5))
					return
				}
				store32(m.memory[uint32(v0):], uint32(i32(-1)))
				m.memory[int64(uint32(v0))+4] = byte(i32(4))
				return
			default:
				if v2 <= i32(-1) {
					m.fn9()
					panic("unreachable")
				}
				if v2 != 0 {
					goto l6
				}
				v3 = i32(1)
				goto l10
			}
		l5:
			t14 := int32(load32(m.memory[uint32(v1):]))
			t15 := int32(load32(m.memory[uint32(v3):]))
			if t14^i32(1279350307)|(t15^i32(558191948)) == 0 {
				store32(m.memory[uint32(v0):], uint32(i32(-1)))
				m.memory[int64(uint32(v0))+4] = byte(i32(6))
				return
			}
		}
	l6:
		t16 := m.fn5(v2)
		v3 = t16
		if v3 == 0 {
			m.fn10(i32(1), v2)
			panic("unreachable")
		}
		if v2 == 0 {
			goto l10
		}
		memory_copy(m.memory, uint32(v3), uint32(v1), uint32(v2))
	}
l10:
	store32(m.memory[int64(uint32(v0))+12:], uint32(v2))
	store32(m.memory[int64(uint32(v0))+8:], uint32(v3))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v2))
	store32(m.memory[uint32(v0):], uint32(i32(-0x7fffffdb)))
}
func (m *Module) fn817(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6 int32
	var v7 int64
	t0 := m.g0
	v2 = t0 - i32(32)
	m.g0 = v2
	t1 := int32(load32(m.memory[uint32(v0):]))
	v3 = t1
	{
		{
			t2 := int32(load32(m.memory[uint32(v1):]))
			v4 = t2
			t3 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t4 := v4
			v5 = t3
			t5 := int32(load32(m.memory[int64(uint32(v5))+12:]))
			v6 = t5
			t6 := m.t0[uint(v6)].(func(int32, int32, int32) int32)(t4, i32(1), i32(0))
			if t6 == 0 {
				goto l0
			}
			v0 = i32(1)
			goto l1
		}
	l0:
		{
			{
				t7 := int32(m.memory[int64(uint32(v1))+10])
				if t7&i32(128) != 0 {
					goto l2
				}
				v0 = i32(1)
				t8 := m.t0[uint(v6)].(func(int32, int32, int32) int32)(v4, i32(1099043), i32(1))
				if t8 != 0 {
					goto l1
				}
				t9 := m.fn688(v3, v1)
				if t9 != 0 {
					goto l1
				}
				t10 := int32(load32(m.memory[int64(uint32(v1))+8:]))
				v0 = t10
				goto l3
			}
		l2:
			{
				t11 := m.t0[uint(v6)].(func(int32, int32, int32) int32)(v4, i32(1099044), i32(2))
				if t11 == 0 {
					goto l4
				}
				v0 = i32(1)
				goto l1
			}
		l4:
			v0 = i32(1)
			m.memory[int64(uint32(v2))+15] = byte(i32(1))
			store32(m.memory[int64(uint32(v2))+4:], uint32(v5))
			store32(m.memory[uint32(v2):], uint32(v4))
			store32(m.memory[int64(uint32(v2))+20:], uint32(i32(1099920)))
			t12 := int64(load64(m.memory[int64(uint32(v1))+8:]))
			t13 := v2
			v7 = t12
			store64(m.memory[int64(uint32(t13))+24:], uint64(v7))
			store32(m.memory[int64(uint32(v2))+8:], uint32(v2+i32(15)))
			store32(m.memory[int64(uint32(v2))+16:], uint32(v2))
			t14 := m.fn688(v3, v2+i32(16))
			if t14 != 0 {
				goto l1
			}
			t15 := int32(load32(m.memory[int64(uint32(v2))+16:]))
			t16 := int32(load32(m.memory[int64(uint32(v2))+20:]))
			t17 := int32(load32(m.memory[int64(uint32(t16))+12:]))
			t18 := m.t0[uint(t17)].(func(int32, int32, int32) int32)(t15, i32(1099041), i32(2))
			if t18 != 0 {
				goto l1
			}
			v0 = int32(v7)
		}
	l3:
		v4 = v3 + i32(4)
		{
			if v0&i32(0x800000) != 0 {
				goto l5
			}
			{
				t19 := int32(load32(m.memory[uint32(v1):]))
				t20 := int32(load32(m.memory[int64(uint32(v1))+4:]))
				t21 := int32(load32(m.memory[int64(uint32(t20))+12:]))
				t22 := m.t0[uint(t21)].(func(int32, int32, int32) int32)(t19, i32(1099034), i32(2))
				if t22 == 0 {
					v0 = i32(1)
					t23 := m.fn688(v4, v1)
					if t23 == 0 {
						goto l7
					}
					goto l1
				}
				v0 = i32(1)
				goto l1
			}
		l5:
			t24 := int64(load64(m.memory[uint32(v1):]))
			v7 = t24
			v0 = i32(1)
			m.memory[int64(uint32(v2))+15] = byte(i32(1))
			store64(m.memory[uint32(v2):], uint64(v7))
			store32(m.memory[int64(uint32(v2))+20:], uint32(i32(1099920)))
			t25 := int64(load64(m.memory[int64(uint32(v1))+8:]))
			store64(m.memory[int64(uint32(v2))+24:], uint64(t25))
			store32(m.memory[int64(uint32(v2))+8:], uint32(v2+i32(15)))
			store32(m.memory[int64(uint32(v2))+16:], uint32(v2))
			t26 := m.fn688(v4, v2+i32(16))
			if t26 != 0 {
				goto l1
			}
			t27 := int32(load32(m.memory[int64(uint32(v2))+16:]))
			t28 := int32(load32(m.memory[int64(uint32(v2))+20:]))
			t29 := int32(load32(m.memory[int64(uint32(t28))+12:]))
			t30 := m.t0[uint(t29)].(func(int32, int32, int32) int32)(t27, i32(1099041), i32(2))
			if t30 != 0 {
				goto l1
			}
		}
	l7:
		t31 := int32(load32(m.memory[uint32(v1):]))
		t32 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t33 := int32(load32(m.memory[int64(uint32(t32))+12:]))
		t34 := m.t0[uint(t33)].(func(int32, int32, int32) int32)(t31, i32(1272328), i32(1))
		v0 = t34
	}
l1:
	m.g0 = v2 + i32(32)
	return v0
}
func (m *Module) fn818(v0, v1 int32) int32 {
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
			t5 := m.t0[uint(v1)].(func(int32, int32, int32) int32)(v2, i32(1081278), i32(7))
			return t5
		case 1:
			t6 := m.t0[uint(v1)].(func(int32, int32, int32) int32)(v2, i32(1081303), i32(4))
			return t6
		case 2:
			t7 := m.t0[uint(v1)].(func(int32, int32, int32) int32)(v2, i32(1081292), i32(6))
			return t7
		case 3:
			t8 := m.t0[uint(v1)].(func(int32, int32, int32) int32)(v2, i32(1081272), i32(6))
			return t8
		case 4:
			t9 := m.t0[uint(v1)].(func(int32, int32, int32) int32)(v2, i32(1081298), i32(5))
			return t9
		case 5:
			t10 := m.t0[uint(v1)].(func(int32, int32, int32) int32)(v2, i32(1080928), i32(5))
			return t10
		case 6:
			t11 := m.t0[uint(v1)].(func(int32, int32, int32) int32)(v2, i32(1081285), i32(7))
			return t11
		case 7:
			t12 := m.t0[uint(v1)].(func(int32, int32, int32) int32)(v2, i32(1091074), i32(6))
			return t12
		}
	}
}
func (m *Module) fn819(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6 int32
	t0 := m.g0
	v2 = t0 - i32(32)
	m.g0 = v2
	v3 = i32(1)
	{
		t1 := int32(load32(m.memory[uint32(v0):]))
		v4 = t1
		t2 := int32(load32(m.memory[uint32(v4):]))
		v0 = t2 ^ i32(-0x80000000)
		p3 := i32(1)
		if uint32(v0) < uint32(i32(6)) {
			p3 = v0
		}
		switch p3 {
		default:
			t4 := int32(load32(m.memory[uint32(v1):]))
			v0 = t4
			t5 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t6 := v0
			v5 = t5
			t7 := int32(load32(m.memory[int64(uint32(v5))+12:]))
			v6 = t7
			t8 := m.t0[uint(v6)].(func(int32, int32, int32) int32)(t6, i32(1091460), i32(2))
			if t8 != 0 {
				goto l6
			}
			v4 = v4 + i32(4)
			{
				{
					t9 := int32(m.memory[int64(uint32(v1))+10])
					if t9&i32(128) != 0 {
						goto l7
					}
					v3 = i32(1)
					t10 := m.t0[uint(v6)].(func(int32, int32, int32) int32)(v0, i32(1099043), i32(1))
					if t10 != 0 {
						goto l6
					}
					t11 := m.fn345(v4, v1)
					if t11 != 0 {
						goto l6
					}
					t12 := int32(load32(m.memory[uint32(v1):]))
					v0 = t12
					t13 := int32(load32(m.memory[int64(uint32(v1))+4:]))
					t14 := int32(load32(m.memory[int64(uint32(t13))+12:]))
					v6 = t14
					goto l8
				}
			l7:
				t15 := m.t0[uint(v6)].(func(int32, int32, int32) int32)(v0, i32(1099044), i32(2))
				if t15 != 0 {
					goto l6
				}
				v3 = i32(1)
				m.memory[int64(uint32(v2))+15] = byte(i32(1))
				store32(m.memory[int64(uint32(v2))+4:], uint32(v5))
				store32(m.memory[uint32(v2):], uint32(v0))
				store32(m.memory[int64(uint32(v2))+20:], uint32(i32(1099920)))
				t16 := int64(load64(m.memory[int64(uint32(v1))+8:]))
				store64(m.memory[int64(uint32(v2))+24:], uint64(t16))
				store32(m.memory[int64(uint32(v2))+8:], uint32(v2+i32(15)))
				store32(m.memory[int64(uint32(v2))+16:], uint32(v2))
				t17 := m.fn345(v4, v2+i32(16))
				if t17 != 0 {
					goto l6
				}
				t18 := int32(load32(m.memory[int64(uint32(v2))+16:]))
				t19 := int32(load32(m.memory[int64(uint32(v2))+20:]))
				t20 := int32(load32(m.memory[int64(uint32(t19))+12:]))
				t21 := m.t0[uint(t20)].(func(int32, int32, int32) int32)(t18, i32(1099041), i32(2))
				if t21 != 0 {
					goto l6
				}
			}
		l8:
			t22 := m.t0[uint(v6)].(func(int32, int32, int32) int32)(v0, i32(1272328), i32(1))
			v3 = t22
			goto l6
		case 1:
			t23 := int32(load32(m.memory[uint32(v1):]))
			v0 = t23
			t24 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t25 := v0
			v5 = t24
			t26 := int32(load32(m.memory[int64(uint32(v5))+12:]))
			v6 = t26
			t27 := m.t0[uint(v6)].(func(int32, int32, int32) int32)(t25, i32(1091462), i32(14))
			if t27 != 0 {
				goto l6
			}
			{
				{
					t28 := int32(m.memory[int64(uint32(v1))+10])
					if t28&i32(128) != 0 {
						goto l9
					}
					v3 = i32(1)
					t29 := m.t0[uint(v6)].(func(int32, int32, int32) int32)(v0, i32(1099043), i32(1))
					if t29 != 0 {
						goto l6
					}
					t30 := int32(load32(m.memory[int64(uint32(v4))+4:]))
					t31 := int32(load32(m.memory[int64(uint32(v4))+8:]))
					t32 := m.fn52(t30, t31, v0, v5)
					if t32 == 0 {
						goto l10
					}
					goto l6
				}
			l9:
				t33 := m.t0[uint(v6)].(func(int32, int32, int32) int32)(v0, i32(1099044), i32(2))
				if t33 != 0 {
					goto l6
				}
				store32(m.memory[int64(uint32(v2))+20:], uint32(v5))
				store32(m.memory[int64(uint32(v2))+16:], uint32(v0))
				v3 = i32(1)
				m.memory[uint32(v2)] = byte(i32(1))
				store32(m.memory[int64(uint32(v2))+24:], uint32(v2))
				t34 := int32(load32(m.memory[int64(uint32(v4))+4:]))
				t35 := int32(load32(m.memory[int64(uint32(v4))+8:]))
				t36 := m.fn52(t34, t35, v2+i32(16), i32(1099920))
				if t36 != 0 {
					goto l6
				}
				t37 := m.fn342(v2+i32(16), i32(1099041), i32(2))
				if t37 != 0 {
					goto l6
				}
			}
		l10:
			t38 := m.t0[uint(v6)].(func(int32, int32, int32) int32)(v0, i32(1272328), i32(1))
			v3 = t38
			goto l6
		case 2:
			t39 := int32(load32(m.memory[uint32(v1):]))
			v0 = t39
			t40 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t41 := v0
			v5 = t40
			t42 := int32(load32(m.memory[int64(uint32(v5))+12:]))
			v6 = t42
			t43 := m.t0[uint(v6)].(func(int32, int32, int32) int32)(t41, i32(1091476), i32(18))
			if t43 != 0 {
				goto l6
			}
			{
				{
					t44 := int32(m.memory[int64(uint32(v1))+10])
					if t44&i32(128) != 0 {
						goto l11
					}
					v3 = i32(1)
					t45 := m.t0[uint(v6)].(func(int32, int32, int32) int32)(v0, i32(1099043), i32(1))
					if t45 != 0 {
						goto l6
					}
					t46 := int32(load32(m.memory[int64(uint32(v4))+4:]))
					t47 := int32(load32(m.memory[uint32(v4+i32(8)):]))
					t48 := m.fn52(t46, t47, v0, v5)
					if t48 == 0 {
						goto l12
					}
					goto l6
				}
			l11:
				t49 := m.t0[uint(v6)].(func(int32, int32, int32) int32)(v0, i32(1099044), i32(2))
				if t49 != 0 {
					goto l6
				}
				store32(m.memory[int64(uint32(v2))+20:], uint32(v5))
				store32(m.memory[int64(uint32(v2))+16:], uint32(v0))
				v3 = i32(1)
				m.memory[uint32(v2)] = byte(i32(1))
				store32(m.memory[int64(uint32(v2))+24:], uint32(v2))
				t50 := int32(load32(m.memory[int64(uint32(v4))+4:]))
				t51 := int32(load32(m.memory[uint32(v4+i32(8)):]))
				t52 := m.fn52(t50, t51, v2+i32(16), i32(1099920))
				if t52 != 0 {
					goto l6
				}
				t53 := m.fn342(v2+i32(16), i32(1099041), i32(2))
				if t53 != 0 {
					goto l6
				}
			}
		l12:
			t54 := m.t0[uint(v6)].(func(int32, int32, int32) int32)(v0, i32(1272328), i32(1))
			v3 = t54
			goto l6
		case 3:
			t55 := int32(load32(m.memory[uint32(v1):]))
			t56 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t57 := int32(load32(m.memory[int64(uint32(t56))+12:]))
			t58 := m.t0[uint(t57)].(func(int32, int32, int32) int32)(t55, i32(1091494), i32(12))
			v3 = t58
			goto l6
		case 4:
			t59 := int32(load32(m.memory[uint32(v1):]))
			t60 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t61 := int32(load32(m.memory[int64(uint32(t60))+12:]))
			t62 := m.t0[uint(t61)].(func(int32, int32, int32) int32)(t59, i32(1091506), i32(15))
			v3 = t62
			goto l6
		case 5:
			t63 := int32(load32(m.memory[uint32(v1):]))
			v0 = t63
			t64 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t65 := v0
			v5 = t64
			t66 := int32(load32(m.memory[int64(uint32(v5))+12:]))
			v6 = t66
			t67 := m.t0[uint(v6)].(func(int32, int32, int32) int32)(t65, i32(1091521), i32(29))
			if t67 != 0 {
				goto l6
			}
			v4 = v4 + i32(4)
			{
				{
					t68 := int32(m.memory[int64(uint32(v1))+10])
					if t68&i32(128) != 0 {
						goto l13
					}
					v3 = i32(1)
					t69 := m.t0[uint(v6)].(func(int32, int32, int32) int32)(v0, i32(1099043), i32(1))
					if t69 != 0 {
						goto l6
					}
					t70 := m.fn804(v4, v1)
					if t70 != 0 {
						goto l6
					}
					t71 := int32(load32(m.memory[uint32(v1):]))
					v0 = t71
					t72 := int32(load32(m.memory[int64(uint32(v1))+4:]))
					t73 := int32(load32(m.memory[int64(uint32(t72))+12:]))
					v6 = t73
					goto l14
				}
			l13:
				t74 := m.t0[uint(v6)].(func(int32, int32, int32) int32)(v0, i32(1099044), i32(2))
				if t74 != 0 {
					goto l6
				}
				v3 = i32(1)
				m.memory[int64(uint32(v2))+15] = byte(i32(1))
				store32(m.memory[int64(uint32(v2))+4:], uint32(v5))
				store32(m.memory[uint32(v2):], uint32(v0))
				store32(m.memory[int64(uint32(v2))+20:], uint32(i32(1099920)))
				t75 := int64(load64(m.memory[int64(uint32(v1))+8:]))
				store64(m.memory[int64(uint32(v2))+24:], uint64(t75))
				store32(m.memory[int64(uint32(v2))+8:], uint32(v2+i32(15)))
				store32(m.memory[int64(uint32(v2))+16:], uint32(v2))
				t76 := m.fn804(v4, v2+i32(16))
				if t76 != 0 {
					goto l6
				}
				t77 := int32(load32(m.memory[int64(uint32(v2))+16:]))
				t78 := int32(load32(m.memory[int64(uint32(v2))+20:]))
				t79 := int32(load32(m.memory[int64(uint32(t78))+12:]))
				t80 := m.t0[uint(t79)].(func(int32, int32, int32) int32)(t77, i32(1099041), i32(2))
				if t80 != 0 {
					goto l6
				}
			}
		l14:
			t81 := m.t0[uint(v6)].(func(int32, int32, int32) int32)(v0, i32(1272328), i32(1))
			v3 = t81
		}
	}
l6:
	m.g0 = v2 + i32(32)
	return v3
}
func (m *Module) fn820(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := m.fn231(t0, v1)
	return t1
}
func (m *Module) fn821(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(m.memory[uint32(t0)])
	t2 := v1
	v0 = t1 << 2
	t3 := int32(load32(m.memory[int64(uint32(v0))+1290696:]))
	t4 := int32(load32(m.memory[int64(uint32(v0))+1290672:]))
	t5 := m.fn56(t2, t3, t4)
	return t5
}
func (m *Module) fn822(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(m.memory[uint32(t0)])
	if t1 != i32(1) {
		t3 := m.fn56(v1, i32(1122568), i32(36))
		return t3
	}
	t2 := m.fn56(v1, i32(1122604), i32(21))
	return t2
}
func (m *Module) fn823(v0, v1 int32) int32 {
	t0 := m.fn56(v1, i32(1122296), i32(41))
	return t0
}
func (m *Module) fn824(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6 int32
	var v7 int64
	t0 := m.g0
	v2 = t0 - i32(48)
	m.g0 = v2
	t1 := int32(load32(m.memory[uint32(v0):]))
	v0 = t1
	t2 := int32(load32(m.memory[uint32(v0+i32(8)):]))
	v3 = t2
	t3 := int32(load32(m.memory[uint32(v0+i32(4)):]))
	v4 = t3
	t4 := int32(load32(m.memory[uint32(v1):]))
	t5 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	t6 := int32(load32(m.memory[int64(uint32(t5))+12:]))
	t7 := m.t0[uint(t6)].(func(int32, int32, int32) int32)(t4, i32(1100400), i32(1))
	v0 = t7
	if v3 != 0 {
		goto l0
	}
	v5 = v0
	goto l1
l0:
	store32(m.memory[int64(uint32(v2))+12:], uint32(v4))
	v5 = i32(1)
	{
		if v0 != 0 {
			goto l2
		}
		{
			t8 := int32(m.memory[int64(uint32(v1))+10])
			if t8&i32(128) == 0 {
				goto l3
			}
			v5 = i32(1)
			t9 := int32(load32(m.memory[uint32(v1):]))
			v0 = t9
			t10 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t11 := v0
			v6 = t10
			t12 := int32(load32(m.memory[int64(uint32(v6))+12:]))
			t13 := m.t0[uint(t12)].(func(int32, int32, int32) int32)(t11, i32(1099046), i32(1))
			if t13 != 0 {
				goto l2
			}
			v5 = i32(1)
			m.memory[int64(uint32(v2))+31] = byte(i32(1))
			store32(m.memory[int64(uint32(v2))+20:], uint32(v6))
			store32(m.memory[int64(uint32(v2))+16:], uint32(v0))
			store32(m.memory[int64(uint32(v2))+36:], uint32(i32(1099920)))
			t14 := int64(load64(m.memory[int64(uint32(v1))+8:]))
			store64(m.memory[int64(uint32(v2))+40:], uint64(t14))
			store32(m.memory[int64(uint32(v2))+24:], uint32(v2+i32(31)))
			store32(m.memory[int64(uint32(v2))+32:], uint32(v2+i32(16)))
			t15 := m.fn293(v2+i32(12), v2+i32(32))
			if t15 != 0 {
				goto l2
			}
			t16 := int32(load32(m.memory[int64(uint32(v2))+32:]))
			t17 := int32(load32(m.memory[int64(uint32(v2))+36:]))
			t18 := int32(load32(m.memory[int64(uint32(t17))+12:]))
			t19 := m.t0[uint(t18)].(func(int32, int32, int32) int32)(t16, i32(1099041), i32(2))
			v5 = t19
			goto l2
		}
	l3:
		t20 := m.fn293(v2+i32(12), v1)
		v5 = t20
	}
l2:
	if v3 == i32(1) {
		goto l1
	}
	v0 = v4 + i32(1)
	v3 = v3 + i32(-1)
l7:
	store32(m.memory[int64(uint32(v2))+12:], uint32(v0))
	v4 = v5 & i32(1)
	v5 = i32(1)
	{
		if v4 != 0 {
			goto l4
		}
		{
			t21 := int32(m.memory[int64(uint32(v1))+10])
			if t21&i32(128) == 0 {
				v5 = i32(1)
				t25 := int32(load32(m.memory[uint32(v1):]))
				t26 := int32(load32(m.memory[int64(uint32(v1))+4:]))
				t27 := int32(load32(m.memory[int64(uint32(t26))+12:]))
				t28 := m.t0[uint(t27)].(func(int32, int32, int32) int32)(t25, i32(1099034), i32(2))
				if t28 != 0 {
					goto l4
				}
				t29 := m.fn293(v2+i32(12), v1)
				v5 = t29
				goto l4
			}
			t22 := int64(load64(m.memory[uint32(v1):]))
			v7 = t22
			m.memory[int64(uint32(v2))+31] = byte(i32(1))
			store64(m.memory[int64(uint32(v2))+16:], uint64(v7))
			store32(m.memory[int64(uint32(v2))+36:], uint32(i32(1099920)))
			t23 := int64(load64(m.memory[int64(uint32(v1))+8:]))
			store64(m.memory[int64(uint32(v2))+40:], uint64(t23))
			store32(m.memory[int64(uint32(v2))+24:], uint32(v2+i32(31)))
			store32(m.memory[int64(uint32(v2))+32:], uint32(v2+i32(16)))
			t24 := m.fn293(v2+i32(12), v2+i32(32))
			if t24 == 0 {
				goto l6
			}
			v5 = i32(1)
			goto l4
		}
	l6:
		t30 := int32(load32(m.memory[int64(uint32(v2))+32:]))
		t31 := int32(load32(m.memory[int64(uint32(v2))+36:]))
		t32 := int32(load32(m.memory[int64(uint32(t31))+12:]))
		t33 := m.t0[uint(t32)].(func(int32, int32, int32) int32)(t30, i32(1099041), i32(2))
		v5 = t33
	}
l4:
	v0 = v0 + i32(1)
	v3 = v3 + i32(-1)
	if v3 != 0 {
		goto l7
	}
l1:
	v0 = i32(1)
	{
		if v5 != 0 {
			goto l8
		}
		t34 := int32(load32(m.memory[uint32(v1):]))
		t35 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t36 := int32(load32(m.memory[int64(uint32(t35))+12:]))
		t37 := m.t0[uint(t36)].(func(int32, int32, int32) int32)(t34, i32(1099049), i32(1))
		v0 = t37
	}
l8:
	m.g0 = v2 + i32(48)
	return v0
}
func (m *Module) fn825(v0, v1 int32) int32 {
	var v2, v3 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	t1 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	v3 = t1
	t2 := int32(load32(m.memory[uint32(v1):]))
	v1 = t2
	{
		{
			t3 := int32(load32(m.memory[uint32(v0):]))
			v0 = t3
			t4 := int32(m.memory[int64(uint32(v0))+4])
			if t4 != i32(2) {
				goto l0
			}
			t5 := int32(load32(m.memory[uint32(v0):]))
			t6 := int64(load64(m.memory[int64(uint32(t5))+12:]))
			store64(m.memory[uint32(v2):], uint64(t6))
			store64(m.memory[int64(uint32(v2))+8:], uint64(int64(uint32(i32(1)))<<32|int64(uint32(v2))))
			t7 := m.fn45(v1, v3, i32(1050614), v2+i32(8))
			v1 = t7
			goto l1
		}
	l0:
		store32(m.memory[uint32(v2):], uint32(v0))
		store64(m.memory[int64(uint32(v2))+8:], uint64(int64(uint32(i32(31)))<<32|int64(uint32(v2))))
		t8 := m.fn45(v1, v3, i32(1052380), v2+i32(8))
		v1 = t8
	}
l1:
	m.g0 = v2 + i32(16)
	return v1
}
func (m *Module) fn826(v0, v1 int32) int32 {
	var v2 int32
	t0 := m.g0
	v2 = t0 - i32(32)
	m.g0 = v2
	{
		{
			t1 := int32(load32(m.memory[uint32(v0):]))
			v0 = t1
			t2 := int32(m.memory[int64(uint32(v0))+4])
			if t2 != i32(1) {
				goto l0
			}
			t3 := int32(m.memory[int64(uint32(v0))+5])
			m.memory[int64(uint32(v2))+15] = byte(t3)
			store64(m.memory[int64(uint32(v2))+24:], uint64(int64(uint32(i32(2)))<<32|int64(uint32(v0))))
			store64(m.memory[int64(uint32(v2))+16:], uint64(int64(uint32(i32(19)))<<32|int64(uint32(v2+i32(15)))))
			t4 := int32(load32(m.memory[uint32(v1):]))
			t5 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t6 := m.fn45(t4, t5, i32(1049276), v2+i32(16))
			v1 = t6
			goto l1
		}
	l0:
		store64(m.memory[int64(uint32(v2))+16:], uint64(int64(uint32(i32(2)))<<32|int64(uint32(v0))))
		t7 := int32(load32(m.memory[uint32(v1):]))
		t8 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t9 := m.fn45(t7, t8, i32(1049325), v2+i32(16))
		v1 = t9
	}
l1:
	m.g0 = v2 + i32(32)
	return v1
}
func (m *Module) fn827(v0, v1 int32) int32 {
	var v2, v3 int32
	var v4 int64
	var v5, v6 int32
	t0 := m.g0
	v2 = t0 - i32(48)
	m.g0 = v2
	t1 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	v3 = t1
	t2 := int32(load32(m.memory[uint32(v1):]))
	v1 = t2
	{
		{
			{
				t3 := int32(load32(m.memory[uint32(v0):]))
				v0 = t3
				t4 := int32(m.memory[uint32(v0)])
				switch t4 {
				default:
					store32(m.memory[int64(uint32(v2))+44:], uint32(v0+i32(4)))
					store64(m.memory[int64(uint32(v2))+16:], uint64(int64(uint32(i32(8)))<<32|int64(uint32(v2+i32(44)))))
					t5 := m.fn45(v1, v3, i32(1051734), v2+i32(16))
					v1 = t5
					goto l7
				case 3:
					store32(m.memory[int64(uint32(v2))+44:], uint32(v0+i32(4)))
					store64(m.memory[int64(uint32(v2))+16:], uint64(int64(uint32(i32(6)))<<32|int64(uint32(v2+i32(44)))))
					t6 := m.fn45(v1, v3, i32(1064620), v2+i32(16))
					v1 = t6
					goto l7
				case 4:
					store32(m.memory[int64(uint32(v2))+8:], uint32(v0+i32(4)))
					store32(m.memory[int64(uint32(v2))+12:], uint32(v0+i32(12)))
					store32(m.memory[int64(uint32(v2))+44:], uint32(v0+i32(2)))
					store64(m.memory[int64(uint32(v2))+32:], uint64(int64(uint32(i32(61)))<<32|int64(uint32(v2+i32(44)))))
					t7 := v2
					v4 = int64(uint32(i32(7))) << 32
					store64(m.memory[int64(uint32(t7))+24:], uint64(v4|int64(uint32(v2+i32(12)))))
					store64(m.memory[int64(uint32(v2))+16:], uint64(v4|int64(uint32(v2+i32(8)))))
					t8 := m.fn45(v1, v3, i32(1050942), v2+i32(16))
					v1 = t8
					goto l7
				case 5:
					store32(m.memory[int64(uint32(v2))+44:], uint32(v0+i32(2)))
					store64(m.memory[int64(uint32(v2))+16:], uint64(int64(uint32(i32(61)))<<32|int64(uint32(v2+i32(44)))))
					t9 := m.fn45(v1, v3, i32(1065487), v2+i32(16))
					v1 = t9
					goto l7
				case 6:
					store32(m.memory[int64(uint32(v2))+44:], uint32(v0+i32(4)))
					store64(m.memory[int64(uint32(v2))+16:], uint64(int64(uint32(i32(35)))<<32|int64(uint32(v2+i32(44)))))
					t10 := m.fn45(v1, v3, i32(1065175), v2+i32(16))
					v1 = t10
					goto l7
				case 1:
					t11 := v2
					v5 = v0 + i32(12)
					store32(m.memory[int64(uint32(t11))+12:], uint32(v5))
					t12 := int32(load32(m.memory[int64(uint32(v3))+12:]))
					t13 := v1
					v6 = t12
					t14 := m.t0[uint(v6)].(func(int32, int32, int32) int32)(t13, i32(1091550), i32(13))
					if t14 != 0 {
						goto l8
					}
					{
						t15 := int32(load32(m.memory[uint32(v5):]))
						v5 = t15
						if uint32(v5) < uint32(i32(512)) {
							store64(m.memory[int64(uint32(v2))+16:], uint64(int64(uint32(i32(35)))<<32|int64(uint32(v2+i32(12)))))
							t17 := m.fn45(v1, v3, i32(1064691), v2+i32(16))
							if t17 != 0 {
								goto l8
							}
							t18 := int32(m.memory[int64(uint32(v0))+1])
							if t18 != i32(1) {
								goto l10
							}
							t19 := m.t0[uint(v6)].(func(int32, int32, int32) int32)(v1, i32(1099034), i32(2))
							if t19 == 0 {
								goto l11
							}
							goto l8
						}
						t16 := int32(m.memory[int64(uint32(v0))+1])
						if t16 == 0 {
							goto l10
						}
						goto l11
					}
				case 2:
					t20 := int32(load32(m.memory[int64(uint32(v3))+12:]))
					t21 := m.t0[uint(t20)].(func(int32, int32, int32) int32)(v1, i32(1091617), i32(20))
					v1 = t21
					goto l7
				}
			}
		l11:
			t22 := m.t0[uint(v6)].(func(int32, int32, int32) int32)(v1, i32(1091563), i32(10))
			if t22 != 0 {
				goto l8
			}
			{
				if v5 == 0 {
					goto l12
				}
				p23 := i32(8)
				if uint32(v5) < uint32(i32(8)) {
					p23 = v5
				}
				v5 = p23
				v0 = v0 + i32(2)
				v4 = int64(uint32(i32(81)))<<32 | int64(uint32(v2+i32(44)))
			l13:
				{
					store32(m.memory[int64(uint32(v2))+44:], uint32(v0))
					store64(m.memory[int64(uint32(v2))+16:], uint64(v4))
					t24 := m.fn45(v1, v3, i32(1091573), v2+i32(16))
					if t24 != 0 {
						goto l8
					}
					v0 = v0 + i32(1)
					v5 = v5 + i32(-1)
					if v5 != 0 {
						goto l13
					}
				}
			}
		l12:
			t25 := m.t0[uint(v6)].(func(int32, int32, int32) int32)(v1, i32(1091584), i32(33))
			if t25 != 0 {
				goto l8
			}
		}
	l10:
		t26 := m.t0[uint(v6)].(func(int32, int32, int32) int32)(v1, i32(1272328), i32(1))
		v1 = t26
		goto l7
	}
l8:
	v1 = i32(1)
l7:
	m.g0 = v2 + i32(48)
	return v1
}
func (m *Module) fn828(v0, v1 int32) int32 {
	var v2, v3 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	t1 := int32(load32(m.memory[uint32(v0):]))
	t2 := int32(load16(m.memory[uint32(t1):]))
	v3 = t2
	v0 = i32(5)
l0:
	{
		t3 := int32(m.memory[uint32(v3&i32(15)+i32(1122552))])
		m.memory[uint32(v2+i32(12)+v0+i32(-2))] = byte(t3)
		v0 = v0 + i32(-1)
		v3 = int32(uint32(v3)>>4) & i32(0xfff)
		if v3 != 0 {
			goto l0
		}
	}
	t4 := m.fn306(v1, i32(1), i32(1122550), i32(2), v2+i32(12)+v0+i32(-1), i32(5)-v0)
	v0 = t4
	m.g0 = v2 + i32(16)
	return v0
}
func (m *Module) fn829(v0, v1 int32) int32 {
	var v2, v3 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	t1 := int32(load32(m.memory[uint32(v0):]))
	t2 := int32(m.memory[uint32(t1)])
	v3 = t2
	v0 = i32(3)
l0:
	{
		t3 := int32(m.memory[uint32(v3&i32(15)+i32(1122552))])
		m.memory[uint32(v2+i32(14)+v0+i32(-2))] = byte(t3)
		v0 = v0 + i32(-1)
		v3 = int32(uint32(v3)>>4) & i32(15)
		if v3 != 0 {
			goto l0
		}
	}
	t4 := m.fn306(v1, i32(1), i32(1122550), i32(2), v2+i32(14)+v0+i32(-1), i32(3)-v0)
	v0 = t4
	m.g0 = v2 + i32(16)
	return v0
}
func (m *Module) fn830(v0, v1 int32) int32 {
	var v2, v3 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	t1 := int32(load32(m.memory[uint32(v0):]))
	t2 := int32(load32(m.memory[uint32(t1):]))
	v3 = t2
	v0 = i32(9)
l0:
	{
		t3 := int32(m.memory[int64(uint32(v3&i32(15)))+1122552])
		m.memory[uint32(v2+i32(8)+v0+i32(-2))] = byte(t3)
		v0 = v0 + i32(-1)
		v3 = int32(uint32(v3) >> 4)
		if v3 != 0 {
			goto l0
		}
	}
	t4 := m.fn306(v1, i32(1), i32(1122550), i32(2), v2+i32(8)+v0+i32(-1), i32(9)-v0)
	v0 = t4
	m.g0 = v2 + i32(16)
	return v0
}
func (m *Module) fn831(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := m.fn166(t0, v1)
	return t1
}
func (m *Module) fn832(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := m.fn193(t0, v1)
	return t1
}
func (m *Module) fn833(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(v1):]))
	t2 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	t3 := m.fn932(t0, t1, t2)
	return t3
}
func (m *Module) fn834(v0, v1 int32) int32 {
	var v2, v3, v4 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	v3 = i32(3)
	t1 := int32(load32(m.memory[uint32(v0):]))
	t2 := int32(m.memory[uint32(t1)])
	v0 = t2
	v4 = v0
	{
		if uint32(v0) < uint32(i32(10)) {
			goto l0
		}
		v3 = i32(1)
		t3 := int32(uint32(v0) / uint32(i32(100)))
		t4 := v2
		t5 := v0
		v4 = t3
		t6 := int32(load16(m.memory[int64(uint32((t5-v4*i32(100))&i32(255)<<1))+1100199:]))
		store16(m.memory[int64(uint32(t4))+14:], uint16(t6))
	}
l0:
	{
		if v0 == 0 {
			goto l1
		}
		if v4 == 0 {
			goto l2
		}
	l1:
		t7 := v2 + i32(13)
		v3 = v3 + i32(-1)
		t8 := int32(m.memory[int64(uint32(v4<<1))+1100200])
		m.memory[uint32(t7+v3)] = byte(t8)
	}
l2:
	t9 := m.fn306(v1, i32(1), i32(1), i32(0), v2+i32(13)+v3, i32(3)-v3)
	v3 = t9
	m.g0 = v2 + i32(16)
	return v3
}
func (m *Module) fn835(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	v0 = t0
	t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t2 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	t3 := int32(load32(m.memory[uint32(v1):]))
	t4 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	t5 := m.fn52(t1, t2, t3, t4)
	return t5
}
func (m *Module) fn836(v0 int32) {
	var v1, v2, v3 int32
	{
		t0 := int32(load32(m.memory[uint32(v0):]))
		v1 = t0
		if v1 == 0 {
			return
		}
		t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		v2 = t1
		t2 := int32(load32(m.memory[uint32(v2+i32(-4)):]))
		v0 = t2
		v3 = v0 & i32(-8)
		t3 := v3
		v0 = v0 & i32(3)
		p4 := i32(8)
		if v0 != 0 {
			p4 = i32(4)
		}
		if uint32(t3) < uint32(p4+v1) {
			m.fn3(i32(1273840), i32(46), i32(1273888))
			panic("unreachable")
		}
		if v0 == 0 {
			goto l2
		}
		if uint32(v3) > uint32(v1+i32(39)) {
			m.fn3(i32(1273904), i32(46), i32(1273952))
			panic("unreachable")
		}
	l2:
		m.fn1(v2)
	}
}
func (m *Module) fn837(v0, v1 int32, v2, v3 int64) int32 {
	var v4, v5, v6, v7, v8, v9, v10, v11 int32
	var v12, v13, v14, v15, v16, v17, v18 int64
	var v19, v20 int32
	var v21 int64
	var v22, v23 int32
	var v24 int64
	var v25 int32
	{
		{
			{
				t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
				v4 = t0
				v1 = v4 + v1
				if uint32(v1) < uint32(v4) {
					m.fn28(i32(1271248), i32(57), i32(1271276))
					panic("unreachable")
				}
				{
					t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
					t2 := v1
					v5 = t1
					t3 := v5
					v6 = v5 + i32(1)
					v7 = int32(uint32(v6) >> 3)
					p4 := v7 * i32(7)
					if uint32(v5) < uint32(i32(8)) {
						p4 = t3
					}
					v8 = p4
					if uint32(t2) <= uint32(int32(uint32(v8)>>1)) {
						{
							if v6 == 0 {
								goto l6
							}
							t7 := int32(load32(m.memory[uint32(v0):]))
							v9 = t7
							v1 = i32(0)
							{
								{
									t8 := v7
									var p9 int32
									if v6&i32(7) != i32(0) {
										p9 = 1
									}
									v7 = t8 + p9
									if v7 == i32(1) {
										goto l7
									}
									v10 = v7 & i32(1)
									v11 = v7 & i32(0x3ffffffe)
									v1 = i32(0)
								l8:
									{
										v7 = v9 + v1
										t10 := int64(load64(m.memory[uint32(v7):]))
										t11 := v7
										v12 = t10
										store64(m.memory[uint32(t11):], uint64(int64(uint64(v12^i64(-1))>>7)&i64(72340172838076673)+(v12|i64(0x7f7f7f7f7f7f7f7f))))
										v7 = v7 + i32(8)
										t12 := int64(load64(m.memory[uint32(v7):]))
										t13 := v7
										v12 = t12
										store64(m.memory[uint32(t13):], uint64(int64(uint64(v12^i64(-1))>>7)&i64(72340172838076673)+(v12|i64(0x7f7f7f7f7f7f7f7f))))
										v1 = v1 + i32(16)
										v11 = v11 + i32(-2)
										if v11 != 0 {
											goto l8
										}
									}
									if v10 == 0 {
										goto l9
									}
								}
							l7:
								v1 = v9 + v1
								t14 := int64(load64(m.memory[uint32(v1):]))
								t15 := v1
								v12 = t14
								store64(m.memory[uint32(t15):], uint64(int64(uint64(v12^i64(-1))>>7)&i64(72340172838076673)+(v12|i64(0x7f7f7f7f7f7f7f7f))))
							}
						l9:
							{
								if uint32(v6) < uint32(i32(8)) {
									goto l10
								}
								t16 := int64(load64(m.memory[uint32(v9):]))
								store64(m.memory[uint32(v9+v6):], uint64(t16))
								goto l11
							}
						l10:
							if v6 == 0 {
								goto l11
							}
							memory_copy(m.memory, uint32(v9+i32(8)), uint32(v9), uint32(v6))
						l11:
							v12 = v3 ^ i64(7237128888997146477)
							v13 = v12 + (v2 ^ i64(8317987319222330741))
							v14 = i64_rotl(v13, i64(32))
							v15 = i64_rotl(v12, i64(13)) ^ v13
							v16 = i64_rotl(v15, i64(17))
							v17 = v3 ^ i64(8098989879002948979)
							v18 = v2 ^ i64(0x6c7967656e657261)
							v7 = i32(0)
						l19:
							{
								t17 := v9
								v1 = v7
								v11 = t17 + v1
								t18 := int32(m.memory[uint32(v11)])
								if t18 != i32(128) {
									goto l12
								}
								v19 = v9 - v1<<3 + i32(-8)
								v20 = v9 + (v1^i32(-1))<<3
								{
								l18:
									{
										t19 := int64(load32(m.memory[uint32(v19):]))
										t20 := v5
										v12 = t19
										v3 = v12 ^ v17
										t21 := i64_rotl(v3, i64(16))
										v3 = v3 + v18
										v2 = t21 ^ v3
										v13 = v2 + v14
										t22 := v13 ^ (v12 | i64(0x400000000000000))
										v12 = v15 + v3
										v3 = v12 ^ v16
										v21 = t22 + v3
										v3 = v21 ^ i64_rotl(v3, i64(13))
										t23 := v3
										t24 := i64_rotl(v12, i64(32)) ^ i64(255)
										v12 = i64_rotl(v2, i64(21)) ^ v13
										v2 = t24 + v12
										v13 = t23 + v2
										v3 = v13 ^ i64_rotl(v3, i64(17))
										t25 := i64_rotl(v3, i64(13))
										t26 := v3
										v12 = v2 ^ i64_rotl(v12, i64(16))
										v2 = v12 + i64_rotl(v21, i64(32))
										v3 = t26 + v2
										v21 = t25 ^ v3
										t27 := i64_rotl(v21, i64(17))
										t28 := v21
										v12 = i64_rotl(v12, i64(21)) ^ v2
										v2 = v12 + i64_rotl(v13, i64(32))
										v13 = t28 + v2
										v21 = t27 ^ v13
										t29 := i64_rotl(v21, i64(13))
										t30 := v21
										v12 = i64_rotl(v12, i64(16)) ^ v2
										v3 = v12 + i64_rotl(v3, i64(32))
										v2 = t29 ^ (t30 + v3)
										t31 := i64_rotl(v2, i64(17))
										v12 = i64_rotl(v12, i64(21)) ^ v3
										t32 := i64_rotl(v12, i64(16))
										v12 = v12 + i64_rotl(v13, i64(32))
										t33 := t31 ^ i64_rotl(t32^v12, i64(21))
										v12 = v2 + v12
										v22 = int32(t33 ^ int64(uint64(v12)>>32) ^ v12)
										v7 = t20 & v22
										v10 = v7
										{
											t34 := int64(load64(m.memory[uint32(v9+v7):]))
											v12 = t34 & i64(-0x7f7f7f7f7f7f7f80)
											if v12 != i64(0) {
												goto l13
											}
											v6 = i32(8)
											v10 = v7
										l14:
											{
												v10 = v10 + v6
												v6 = v6 + i32(8)
												t35 := v9
												v10 = v10 & v5
												t36 := int64(load64(m.memory[uint32(t35+v10):]))
												v12 = t36 & i64(-0x7f7f7f7f7f7f7f80)
												if v12 == 0 {
													goto l14
												}
											}
										}
									l13:
										{
											t37 := v9
											v10 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v12))))>>3) + v10) & v5
											t38 := int32(int8(m.memory[uint32(t37+v10)]))
											if t38 < i32(0) {
												goto l15
											}
											t39 := int64(load64(m.memory[uint32(v9):]))
											v10 = int32(uint32(int64(bits.TrailingZeros64(uint64(t39&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
										}
									l15:
										{
											if uint32((v10-v7^(v1-v7))&v5) < uint32(i32(8)) {
												goto l16
											}
											v7 = v9 + v10
											t40 := int32(m.memory[uint32(v7)])
											v6 = t40
											t41 := v7
											v22 = int32(uint32(v22) >> 25)
											m.memory[uint32(t41)] = byte(v22)
											m.memory[uint32(v9+(v10+i32(-8))&v5+i32(8))] = byte(v22)
											v7 = v9 - v10<<3 + i32(-8)
											if v6 == i32(255) {
												goto l17
											}
											t42 := int32(load32(m.memory[uint32(v7):]))
											v10 = t42
											t43 := int32(load32(m.memory[uint32(v20):]))
											store32(m.memory[uint32(v7):], uint32(t43))
											store32(m.memory[uint32(v20):], uint32(v10))
											t44 := int32(load32(m.memory[int64(uint32(v20))+4:]))
											v10 = t44
											t45 := int32(load32(m.memory[int64(uint32(v7))+4:]))
											store32(m.memory[int64(uint32(v20))+4:], uint32(t45))
											store32(m.memory[int64(uint32(v7))+4:], uint32(v10))
											goto l18
										}
									l16:
									}
									t46 := v11
									v7 = int32(uint32(v22) >> 25)
									m.memory[uint32(t46)] = byte(v7)
									m.memory[uint32(v9+(v1+i32(-8))&v5+i32(8))] = byte(v7)
									goto l12
								}
							l17:
								m.memory[uint32(v11)] = byte(i32(255))
								m.memory[uint32(v9+(v1+i32(-8))&v5+i32(8))] = byte(i32(255))
								t47 := int64(load64(m.memory[uint32(v20):]))
								store64(m.memory[uint32(v7):], uint64(t47))
							}
						l12:
							v7 = v1 + i32(1)
							if v1 != v5 {
								goto l19
							}
						}
					l6:
						store32(m.memory[int64(uint32(v0))+8:], uint32(v8-v4))
						goto l20
					}
					v9 = v8 + i32(1)
					p5 := v1
					if uint32(v9) > uint32(v1) {
						p5 = v9
					}
					v1 = p5
					if uint32(v1) < uint32(i32(15)) {
						goto l2
					}
					{
						if uint32(v1) > uint32(i32(0x1fffffff)) {
							m.fn28(i32(1271248), i32(57), i32(1271276))
							panic("unreachable")
						}
						t6 := int32(uint32(v1<<3) / uint32(i32(7)))
						v1 = i32_shr_u(i32(-1), int32(bits.LeadingZeros32(uint32(t6+i32(-1)))))
						if uint32(v1) > uint32(i32(0x1ffffffe)) {
							goto l4
						}
						v1 = v1 + i32(1)
						goto l5
					}
				}
			}
		l2:
			p48 := v1&i32(8) + i32(8)
			if uint32(v1) < uint32(i32(4)) {
				p48 = i32(4)
			}
			v1 = p48
		}
	l5:
		v9 = v1 + i32(8)
		t49 := v9
		v11 = v1 << 3
		v7 = t49 + v11
		if uint32(v7) < uint32(v9) {
			goto l4
		}
		if uint32(v7) > uint32(i32(0x7ffffff8)) {
			goto l4
		}
		{
			t50 := m.fn5(v7)
			v10 = t50
			if v10 != 0 {
				v7 = v10 + v11
				if v9 == 0 {
					goto l22
				}
				memory_fill(m.memory, uint32(v7), i32(255), uint32(v9))
			l22:
				v10 = v1 + i32(-1)
				p51 := int32(uint32(v1)>>3) * i32(7)
				if uint32(v1) < uint32(i32(9)) {
					p51 = v10
				}
				v23 = p51
				t52 := int32(load32(m.memory[uint32(v0):]))
				v8 = t52
				{
					if v4 == 0 {
						goto l23
					}
					v12 = v3 ^ i64(7237128888997146477)
					v13 = v12 + (v2 ^ i64(8317987319222330741))
					v15 = i64_rotl(v13, i64(32))
					v16 = i64_rotl(v12, i64(13)) ^ v13
					v17 = i64_rotl(v16, i64(17))
					v18 = v3 ^ i64(8098989879002948979)
					v24 = v2 ^ i64(0x6c7967656e657261)
					t53 := int64(load64(m.memory[uint32(v8):]))
					v12 = (t53 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
					v9 = v8
					v1 = i32(0)
					v22 = v4
				l29:
					{
						if v12 != i64(0) {
							goto l24
						}
					l25:
						{
							v1 = v1 + i32(8)
							v9 = v9 + i32(8)
							t54 := int64(load64(m.memory[uint32(v9):]))
							v12 = t54 & i64(-0x7f7f7f7f7f7f7f80)
							if v12 == i64(-0x7f7f7f7f7f7f7f80) {
								goto l25
							}
						}
						v12 = v12 ^ i64(-0x7f7f7f7f7f7f7f80)
					l24:
						{
							t55 := v7
							t56 := v10
							v19 = v8 - (int32(uint32(int64(bits.TrailingZeros64(uint64(v12))))>>3)+v1)<<3 + i32(-8)
							t57 := int64(load32(m.memory[uint32(v19):]))
							v3 = t57
							v2 = v3 ^ v18
							t58 := i64_rotl(v2, i64(16))
							v2 = v2 + v24
							v13 = t58 ^ v2
							v21 = v13 + v15
							t59 := v21 ^ (v3 | i64(0x400000000000000))
							v3 = v16 + v2
							v2 = v3 ^ v17
							v14 = t59 + v2
							v2 = v14 ^ i64_rotl(v2, i64(13))
							t60 := v2
							t61 := i64_rotl(v3, i64(32)) ^ i64(255)
							v3 = i64_rotl(v13, i64(21)) ^ v21
							v13 = t61 + v3
							v21 = t60 + v13
							v2 = v21 ^ i64_rotl(v2, i64(17))
							t62 := i64_rotl(v2, i64(13))
							t63 := v2
							v3 = v13 ^ i64_rotl(v3, i64(16))
							v13 = v3 + i64_rotl(v14, i64(32))
							v2 = t63 + v13
							v14 = t62 ^ v2
							t64 := i64_rotl(v14, i64(17))
							t65 := v14
							v3 = i64_rotl(v3, i64(21)) ^ v13
							v13 = v3 + i64_rotl(v21, i64(32))
							v21 = t65 + v13
							v14 = t64 ^ v21
							t66 := i64_rotl(v14, i64(13))
							t67 := v14
							v3 = i64_rotl(v3, i64(16)) ^ v13
							v2 = v3 + i64_rotl(v2, i64(32))
							v13 = t66 ^ (t67 + v2)
							t68 := i64_rotl(v13, i64(17))
							v3 = i64_rotl(v3, i64(21)) ^ v2
							t69 := i64_rotl(v3, i64(16))
							v3 = v3 + i64_rotl(v21, i64(32))
							t70 := t68 ^ i64_rotl(t69^v3, i64(21))
							v3 = v13 + v3
							v20 = int32(t70 ^ int64(uint64(v3)>>32) ^ v3)
							v11 = t56 & v20
							t71 := int64(load64(m.memory[uint32(t55+v11):]))
							v3 = t71 & i64(-0x7f7f7f7f7f7f7f80)
							if v3 != i64(0) {
								goto l26
							}
							v25 = i32(8)
						l27:
							{
								v11 = v11 + v25
								v25 = v25 + i32(8)
								t72 := v7
								v11 = v11 & v10
								t73 := int64(load64(m.memory[uint32(t72+v11):]))
								v3 = t73 & i64(-0x7f7f7f7f7f7f7f80)
								if v3 == 0 {
									goto l27
								}
							}
						}
					l26:
						v2 = v12 + i64(-1)
						{
							t74 := v7
							v11 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v3))))>>3) + v11) & v10
							t75 := int32(int8(m.memory[uint32(t74+v11)]))
							if t75 < i32(0) {
								goto l28
							}
							t76 := int64(load64(m.memory[uint32(v7):]))
							v11 = int32(uint32(int64(bits.TrailingZeros64(uint64(t76&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
						}
					l28:
						v12 = v2 & v12
						t77 := v7 + v11
						v20 = int32(uint32(v20) >> 25)
						m.memory[uint32(t77)] = byte(v20)
						m.memory[uint32(v7+(v11+i32(-8))&v10+i32(8))] = byte(v20)
						t78 := int64(load64(m.memory[uint32(v19):]))
						store64(m.memory[uint32(v7-v11<<3+i32(-8)):], uint64(t78))
						v22 = v22 + i32(-1)
						if v22 != 0 {
							goto l29
						}
					}
				}
			l23:
				store32(m.memory[int64(uint32(v0))+4:], uint32(v10))
				store32(m.memory[uint32(v0):], uint32(v7))
				store32(m.memory[int64(uint32(v0))+8:], uint32(v23-v4))
				if v5 == 0 {
					goto l20
				}
				t79 := v5
				v9 = v6 << 3
				v1 = t79 + v9 + i32(9)
				if v1 == 0 {
					goto l20
				}
				v5 = v8 - v9
				t80 := int32(load32(m.memory[uint32(v5+i32(-4)):]))
				v9 = t80
				v7 = v9 & i32(-8)
				t81 := v7
				v9 = v9 & i32(3)
				p82 := i32(8)
				if v9 != 0 {
					p82 = i32(4)
				}
				if uint32(t81) < uint32(p82+v1) {
					m.fn3(i32(1273840), i32(46), i32(1273888))
					panic("unreachable")
				}
				if v9 == 0 {
					goto l31
				}
				if uint32(v7) > uint32(v1+i32(39)) {
					m.fn3(i32(1273904), i32(46), i32(1273952))
					panic("unreachable")
				}
			l31:
				m.fn1(v5)
				return i32(-1)
			}
			m.fn24(i32(8), v7)
			panic("unreachable")
		}
	}
l20:
	return i32(-1)
l4:
	m.fn28(i32(1271248), i32(57), i32(1271276))
	panic("unreachable")
}
func (m *Module) fn838(v0, v1, v2, v3, v4, v5 int32) {
	var v6, v7 int32
	var v8 int64
	v6 = i32(1)
	v7 = i32(4)
	v8 = int64(uint32(v5)) * int64(uint32(v3))
	if int32(int64(uint64(v8)>>32)) == 0 {
		goto l0
	}
	v3 = i32(0)
	goto l1
l0:
	v3 = int32(v8)
	if uint32(v3) <= uint32(i32(-0x80000000)-v4) {
		goto l2
	}
	v3 = i32(0)
	goto l1
l2:
	{
		{
			if v1 == 0 {
				goto l3
			}
			t0 := m.fn22(v2, v5*v1, v4, v3)
			v7 = t0
			goto l4
		}
	l3:
		if v3 != 0 {
			goto l5
		}
		v7 = v4
		goto l6
	l5:
		t1 := m.fn20(v3, v4)
		v7 = t1
	}
l4:
	if v7 != 0 {
		goto l6
	}
	store32(m.memory[int64(uint32(v0))+4:], uint32(v4))
	goto l7
l6:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v7))
	v6 = i32(0)
l7:
	v7 = i32(8)
l1:
	store32(m.memory[uint32(v0+v7):], uint32(v3))
	store32(m.memory[uint32(v0):], uint32(v6))
}
func (m *Module) fn839(v0 int32) int32 {
	var v1, v2, v3, v4 int32
	var v5, v6, v7, v8, v9, v10 int64
	var v11, v12, v13 int32
	t0 := m.g0
	v1 = t0 - i32(16)
	m.g0 = v1
	{
		t1 := int32(m.memory[int64(uint32(i32(0)))+1293376])
		if t1 == i32(3) {
			goto l0
		}
		m.fn840()
	}
l0:
	v2 = i32(-1)
	{
		t2 := int32(load32(m.memory[int64(uint32(i32(0)))+1293356:]))
		if t2 == 0 {
			goto l1
		}
		v3 = i32(0)
		t3 := int32(load32(m.memory[int64(uint32(i32(0)))+1293348:]))
		v4 = t3
		t4 := int64(load64(m.memory[int64(uint32(i32(0)))+1293368:]))
		t5 := v4
		v5 = t4
		t6 := v5
		v6 = int64(uint32(v0))
		v7 = t6 ^ v6 ^ i64(8098989879002948979)
		t7 := int64(load64(m.memory[int64(uint32(i32(0)))+1293360:]))
		t8 := i64_rotl(v7, i64(16))
		t9 := v7
		v8 = t7
		v7 = t9 + (v8 ^ i64(0x6c7967656e657261))
		v9 = t8 ^ v7
		t10 := v9
		v5 = v5 ^ i64(7237128888997146477)
		v8 = v5 + (v8 ^ i64(8317987319222330741))
		v10 = t10 + i64_rotl(v8, i64(32))
		t11 := v10 ^ (v6 | i64(0x400000000000000))
		v5 = i64_rotl(v5, i64(13)) ^ v8
		v6 = v5 + v7
		v5 = v6 ^ i64_rotl(v5, i64(17))
		v7 = t11 + v5
		v5 = v7 ^ i64_rotl(v5, i64(13))
		t12 := v5
		t13 := i64_rotl(v6, i64(32)) ^ i64(255)
		v6 = i64_rotl(v9, i64(21)) ^ v10
		v8 = t13 + v6
		v9 = t12 + v8
		v5 = v9 ^ i64_rotl(v5, i64(17))
		t14 := i64_rotl(v5, i64(13))
		t15 := v5
		v6 = v8 ^ i64_rotl(v6, i64(16))
		v7 = v6 + i64_rotl(v7, i64(32))
		v5 = t15 + v7
		v8 = t14 ^ v5
		t16 := i64_rotl(v8, i64(17))
		t17 := v8
		v6 = i64_rotl(v6, i64(21)) ^ v7
		v7 = v6 + i64_rotl(v9, i64(32))
		v8 = t17 + v7
		v9 = t16 ^ v8
		t18 := i64_rotl(v9, i64(13))
		t19 := v9
		v6 = i64_rotl(v6, i64(16)) ^ v7
		v5 = v6 + i64_rotl(v5, i64(32))
		v7 = t18 ^ (t19 + v5)
		t20 := i64_rotl(v7, i64(17))
		v5 = i64_rotl(v6, i64(21)) ^ v5
		t21 := i64_rotl(v5, i64(16))
		v5 = v5 + i64_rotl(v8, i64(32))
		t22 := t20 ^ i64_rotl(t21^v5, i64(21))
		v5 = v7 + v5
		v5 = t22 ^ int64(uint64(v5)>>32) ^ v5
		v11 = t5 & int32(v5)
		v6 = int64(uint64(v5)>>25) & i64(127) * i64(72340172838076673)
		t23 := int32(load32(m.memory[int64(uint32(i32(0)))+1293344:]))
		v12 = t23
	l5:
		{
			{
				t24 := int64(load64(m.memory[uint32(v12+v11):]))
				v7 = t24
				v5 = v7 ^ v6
				v5 = (v5 ^ i64(-1)) & (v5 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
				if v5 == 0 {
					goto l2
				}
			l4:
				{
					t25 := v0
					v13 = v12 - (int32(uint32(int64(bits.TrailingZeros64(uint64(v5))))>>3)+v11)&v4<<3
					t26 := int32(load32(m.memory[uint32(v13+i32(-8)):]))
					if t25 == t26 {
						goto l3
					}
					v5 = (v5 + i64(-1)) & v5
					if !(v5 == 0) {
						goto l4
					}
				}
			}
		l2:
			if !(v7&(v7<<1)&i64(-0x7f7f7f7f7f7f7f80) == 0) {
				goto l1
			}
			t27 := v11
			v3 = v3 + i32(8)
			v11 = (t27 + v3) & v4
			goto l5
		}
	l3:
		t28 := int32(load32(m.memory[uint32(v13+i32(-4)):]))
		v2 = t28
	}
l1:
	{
		{
			if uint32(v0) < uint32(i32(181)) {
				goto l6
			}
			m.fn34(v1+i32(4), v0, i32(1105692))
			t29 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t30 := v0
			v11 = t29
			p31 := v11
			if v11 == i32(-1) {
				p31 = t30
			}
			v0 = p31
			goto l7
		}
	l6:
		p32 := v0
		if uint32(v0+i32(-97)) < uint32(i32(26)) {
			p32 = v0 & i32(95)
		}
		v0 = p32
	}
l7:
	m.g0 = v1 + i32(16)
	p33 := v2
	if v2 == i32(-1) {
		p33 = v0
	}
	return p33
}
func (m *Module) fn840() {
	var v0 int32
	var v1, v2 int64
	var v3, v4, v5 int32
	var v6, v7, v8, v9, v10, v11 int64
	var v12, v13, v14, v15, v16, v17, v18 int32
	t0 := m.g0
	v0 = t0 - i32(48)
	m.g0 = v0
	{
		t1 := int32(m.memory[int64(uint32(i32(0)))+1293376])
		switch t1 {
		case 3:
			goto l2
		default:
			m.memory[int64(uint32(i32(0)))+1293376] = byte(i32(2))
			{
				{
					t2 := int32(m.memory[int64(uint32(i32(0)))+1293880])
					if t2 == 0 {
						goto l3
					}
					t3 := int64(load64(m.memory[int64(uint32(i32(0)))+1293872:]))
					v1 = t3
					t4 := int64(load64(m.memory[int64(uint32(i32(0)))+1293864:]))
					v2 = t4
					goto l4
				}
			l3:
				m.fn194(v0 + i32(32))
				m.memory[int64(uint32(i32(0)))+1293880] = byte(i32(1))
				t5 := int64(load64(m.memory[int64(uint32(v0))+40:]))
				v1 = t5
				store64(m.memory[int64(uint32(i32(0)))+1293872:], uint64(v1))
				t6 := int64(load64(m.memory[int64(uint32(v0))+32:]))
				v2 = t6
			}
		l4:
			store64(m.memory[int64(uint32(i32(0)))+1293864:], uint64(v2+i64(1)))
			store64(m.memory[int64(uint32(v0))+24:], uint64(v1))
			store64(m.memory[int64(uint32(v0))+16:], uint64(v2))
			t7 := int64(load64(m.memory[int64(uint32(i32(0)))+1275664:]))
			store64(m.memory[uint32(v0):], uint64(t7))
			t8 := int64(load64(m.memory[int64(uint32(i32(0)))+1275672:]))
			store64(m.memory[int64(uint32(v0))+8:], uint64(t8))
			_ = m.fn837(v0, i32(129), v2, v1)
			v3 = i32(0)
		l15:
			{
				t10 := int64(load64(m.memory[int64(uint32(v0))+24:]))
				v1 = t10
				t11 := v1
				v4 = v3 << 3
				t12 := int32(load32(m.memory[int64(uint32(v4))+1091760:]))
				v5 = t12
				v2 = int64(uint32(v5))
				v6 = t11 ^ v2 ^ i64(8098989879002948979)
				t13 := int64(load64(m.memory[int64(uint32(v0))+16:]))
				t14 := i64_rotl(v6, i64(16))
				t15 := v6
				v7 = t13
				v6 = t15 + (v7 ^ i64(0x6c7967656e657261))
				v8 = t14 ^ v6
				t16 := v8
				v9 = v1 ^ i64(7237128888997146477)
				v10 = v9 + (v7 ^ i64(8317987319222330741))
				v11 = t16 + i64_rotl(v10, i64(32))
				t17 := v11 ^ (v2 | i64(0x400000000000000))
				v2 = i64_rotl(v9, i64(13)) ^ v10
				v6 = v2 + v6
				v2 = v6 ^ i64_rotl(v2, i64(17))
				v9 = t17 + v2
				v2 = v9 ^ i64_rotl(v2, i64(13))
				t18 := v2
				t19 := i64_rotl(v6, i64(32)) ^ i64(255)
				v6 = i64_rotl(v8, i64(21)) ^ v11
				v8 = t19 + v6
				v10 = t18 + v8
				v2 = v10 ^ i64_rotl(v2, i64(17))
				t20 := i64_rotl(v2, i64(13))
				t21 := v2
				v6 = v8 ^ i64_rotl(v6, i64(16))
				v8 = v6 + i64_rotl(v9, i64(32))
				v2 = t21 + v8
				v9 = t20 ^ v2
				t22 := i64_rotl(v9, i64(17))
				t23 := v9
				v6 = i64_rotl(v6, i64(21)) ^ v8
				v8 = v6 + i64_rotl(v10, i64(32))
				v9 = t23 + v8
				v10 = t22 ^ v9
				t24 := i64_rotl(v10, i64(13))
				t25 := v10
				v6 = i64_rotl(v6, i64(16)) ^ v8
				v2 = v6 + i64_rotl(v2, i64(32))
				v8 = t24 ^ (t25 + v2)
				t26 := i64_rotl(v8, i64(17))
				v2 = i64_rotl(v6, i64(21)) ^ v2
				t27 := i64_rotl(v2, i64(16))
				v2 = v2 + i64_rotl(v9, i64(32))
				t28 := t26 ^ i64_rotl(t27^v2, i64(21))
				v2 = v8 + v2
				v2 = t28 ^ int64(uint64(v2)>>32) ^ v2
				t29 := int32(load32(m.memory[uint32(v4+i32(1091764)):]))
				v12 = t29
				{
					t30 := int32(load32(m.memory[int64(uint32(v0))+8:]))
					if t30 != 0 {
						goto l5
					}
					_ = m.fn837(v0, i32(1), v7, v1)
				}
			l5:
				t32 := int32(load32(m.memory[int64(uint32(v0))+4:]))
				v13 = t32
				v14 = v13 & int32(v2)
				v6 = int64(uint64(v2) >> 25)
				v1 = v6 & i64(127) * i64(72340172838076673)
				v15 = i32(0)
				t33 := int32(load32(m.memory[uint32(v0):]))
				v4 = t33
				v16 = i32(0)
			l16:
				{
					{
						{
							t34 := int64(load64(m.memory[uint32(v4+v14):]))
							v7 = t34
							v2 = v7 ^ v1
							v2 = (v2 ^ i64(-1)) & (v2 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
							if v2 == 0 {
								goto l6
							}
						l9:
							{
								t35 := v5
								t36 := v4
								v17 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v2))))>>3) + v14) & v13
								t37 := int32(load32(m.memory[uint32(t36-v17<<3+i32(-8)):]))
								if t35 != t37 {
									goto l7
								}
								v14 = i32(0) - v17
								goto l8
							}
						l7:
							v2 = (v2 + i64(-1)) & v2
							if !(v2 == 0) {
								goto l9
							}
						}
					l6:
						v2 = v7 & i64(-0x7f7f7f7f7f7f7f80)
						if v15 == i32(1) {
							goto l10
						}
						if v2 == 0 {
							goto l11
						}
						v18 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v2))))>>3) + v14) & v13
					l10:
						if v2&(v7<<1) != i64(0) {
							goto l12
						}
						v15 = i32(1)
						goto l13
					l12:
						{
							t38 := int32(int8(m.memory[uint32(v4+v18)]))
							v14 = t38
							if v14 < i32(0) {
								goto l14
							}
							t39 := int64(load64(m.memory[uint32(v4):]))
							t40 := v4
							v18 = int32(uint32(int64(bits.TrailingZeros64(uint64(t39&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
							t41 := int32(m.memory[uint32(t40+v18)])
							v14 = t41
						}
					l14:
						t42 := v4 + v18
						v15 = int32(v6) & i32(127)
						m.memory[uint32(t42)] = byte(v15)
						m.memory[uint32(v4+(v18+i32(-8))&v13+i32(8))] = byte(v15)
						store32(m.memory[uint32(v4-v18<<3+i32(-8)):], uint32(v5))
						t43 := int32(load32(m.memory[int64(uint32(v0))+12:]))
						store32(m.memory[int64(uint32(v0))+12:], uint32(t43+i32(1)))
						t44 := int32(load32(m.memory[int64(uint32(v0))+8:]))
						store32(m.memory[int64(uint32(v0))+8:], uint32(t44-v14&i32(1)))
						v14 = i32(0) - v18
					}
				l8:
					store32(m.memory[uint32(v4+v14<<3+i32(-4)):], uint32(v12))
					v3 = v3 + i32(1)
					if v3 != i32(129) {
						goto l15
					}
					t45 := int64(load64(m.memory[int64(uint32(v0))+24:]))
					store64(m.memory[int64(uint32(i32(0)))+1293368:], uint64(t45))
					t46 := int64(load64(m.memory[int64(uint32(v0))+16:]))
					store64(m.memory[int64(uint32(i32(0)))+1293360:], uint64(t46))
					t47 := int64(load64(m.memory[int64(uint32(v0))+8:]))
					store64(m.memory[int64(uint32(i32(0)))+1293352:], uint64(t47))
					t48 := int64(load64(m.memory[uint32(v0):]))
					store64(m.memory[int64(uint32(i32(0)))+1293344:], uint64(t48))
					m.memory[int64(uint32(i32(0)))+1293376] = byte(i32(3))
					goto l2
				}
			l11:
				v15 = i32(0)
			l13:
				v16 = v16 + i32(8)
				v14 = (v16 + v14) & v13
				goto l16
			}
		case 2:
			m.fn28(i32(1091656), i32(113), i32(1091640))
			panic("unreachable")
		}
	}
l2:
	m.g0 = v0 + i32(48)
}
func (m *Module) fn841(v0, v1 int32) int32 {
	var v2, v3, v4, v5 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	t1 := int32(load32(m.memory[uint32(v0):]))
	v0 = t1
	{
		{
			t2 := int32(m.memory[int64(uint32(v1))+11])
			if t2&i32(24) != 0 {
				goto l0
			}
			t3 := int32(load32(m.memory[uint32(v1):]))
			t4 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t5 := int32(load32(m.memory[int64(uint32(t4))+16:]))
			t6 := m.t0[uint(t5)].(func(int32, int32) int32)(t3, v0)
			v0 = t6
			goto l1
		}
	l0:
		store32(m.memory[int64(uint32(v2))+12:], uint32(i32(0)))
		if uint32(v0) < uint32(i32(128)) {
			goto l2
		}
		v3 = v0&i32(63) | i32(-128)
		v4 = int32(uint32(v0) >> 6)
		if uint32(v0) >= uint32(i32(2048)) {
			v5 = int32(uint32(v0) >> 12)
			v4 = v4&i32(63) | i32(-128)
			if uint32(v0) > uint32(i32(0xffff)) {
				m.memory[int64(uint32(v2))+15] = byte(v3)
				m.memory[int64(uint32(v2))+14] = byte(v4)
				m.memory[int64(uint32(v2))+13] = byte(v5&i32(63) | i32(-128))
				m.memory[int64(uint32(v2))+12] = byte(int32(uint32(v0)>>18) | i32(-16))
				v0 = i32(4)
				goto l4
			}
			m.memory[int64(uint32(v2))+14] = byte(v3)
			m.memory[int64(uint32(v2))+13] = byte(v4)
			m.memory[int64(uint32(v2))+12] = byte(v5 | i32(224))
			v0 = i32(3)
			goto l4
		}
		m.memory[int64(uint32(v2))+13] = byte(v3)
		m.memory[int64(uint32(v2))+12] = byte(v4 | i32(192))
		v0 = i32(2)
		goto l4
	l2:
		m.memory[int64(uint32(v2))+12] = byte(v0)
		v0 = i32(1)
	l4:
		t7 := m.fn56(v1, v2+i32(12), v0)
		v0 = t7
	}
l1:
	m.g0 = v2 + i32(16)
	return v0
}
func (m *Module) fn842(v0, v1, v2, v3, v4, v5, v6, v7 int32) {
	var v8 int32
	var v9 int64
	t0 := m.g0
	v8 = t0 - i32(64)
	m.g0 = v8
	store32(m.memory[int64(uint32(v8))+4:], uint32(v2))
	store32(m.memory[uint32(v8):], uint32(v1))
	store32(m.memory[int64(uint32(v8))+12:], uint32(v4))
	store32(m.memory[int64(uint32(v8))+8:], uint32(v3))
	store32(m.memory[int64(uint32(v8))+20:], uint32(i32(2)))
	t2 := v8
	p1 := i32(1100195)
	if v0&i32(1) != 0 {
		p1 = i32(1100197)
	}
	store32(m.memory[int64(uint32(t2))+16:], uint32(p1))
	{
		if v5 == 0 {
			t4 := v8
			v9 = int64(uint32(i32(9))) << 32
			store64(m.memory[int64(uint32(t4))+48:], uint64(v9|int64(uint32(v8+i32(8)))))
			store64(m.memory[int64(uint32(v8))+40:], uint64(v9|int64(uint32(v8))))
			store64(m.memory[int64(uint32(v8))+32:], uint64(int64(uint32(i32(10)))<<32|int64(uint32(v8+i32(16)))))
			m.fn28(i32(1051281), v8+i32(32), v7)
			panic("unreachable")
		}
		store32(m.memory[int64(uint32(v8))+28:], uint32(v6))
		store32(m.memory[int64(uint32(v8))+24:], uint32(v5))
		t3 := v8
		v9 = int64(uint32(i32(9))) << 32
		store64(m.memory[int64(uint32(t3))+56:], uint64(v9|int64(uint32(v8+i32(8)))))
		store64(m.memory[int64(uint32(v8))+48:], uint64(v9|int64(uint32(v8))))
		store64(m.memory[int64(uint32(v8))+40:], uint64(int64(uint32(i32(89)))<<32|int64(uint32(v8+i32(24)))))
		store64(m.memory[int64(uint32(v8))+32:], uint64(int64(uint32(i32(10)))<<32|int64(uint32(v8+i32(16)))))
		m.fn28(i32(1051336), v8+i32(32), v7)
		panic("unreachable")
	}
}
func (m *Module) fn843(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t2 := int32(load32(m.memory[int64(uint32(t1))+12:]))
	t3 := m.t0[uint(t2)].(func(int32, int32) int32)(t0, v1)
	return t3
}
func (m *Module) fn844(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v1):]))
	t1 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	t2 := int32(load32(m.memory[uint32(v0):]))
	t3 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t4 := m.fn45(t0, t1, t2, t3)
	return t4
}
func (m *Module) fn845(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t2 := m.fn56(v1, t0, t1)
	return t2
}
func (m *Module) fn846(v0 int32) {
	var v1 int32
	var v2 int64
	t0 := m.g0
	v1 = t0 - i32(16)
	m.g0 = v1
	t1 := int64(load64(m.memory[uint32(v0):]))
	v2 = t1
	store32(m.memory[int64(uint32(v1))+12:], uint32(v0))
	store64(m.memory[int64(uint32(v1))+4:], uint64(v2))
	m.fn942(v1 + i32(4))
	panic("unreachable")
}
func (m *Module) fn847(v0, v1 int32) int32 {
	var v2, v3 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	t1 := int32(load32(m.memory[uint32(v0):]))
	v0 = t1
	{
		{
			t2 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			v3 = t2
			if v3&i32(0x2000000) != 0 {
				t4 := int32(load32(m.memory[uint32(v0):]))
				v3 = t4
				v0 = i32(9)
			l3:
				{
					t5 := int32(m.memory[int64(uint32(v3&i32(15)))+1098816])
					m.memory[uint32(v2+i32(8)+v0+i32(-2))] = byte(t5)
					v0 = v0 + i32(-1)
					v3 = int32(uint32(v3) >> 4)
					if v3 != 0 {
						goto l3
					}
				}
				t6 := m.fn306(v1, i32(1), i32(1122550), i32(2), v2+i32(8)+v0+i32(-1), i32(9)-v0)
				v0 = t6
				goto l2
			}
			if v3&i32(0x4000000) != 0 {
				goto l1
			}
			t3 := m.fn14(v0, v1)
			v0 = t3
			goto l2
		}
	l1:
		t7 := int32(load32(m.memory[uint32(v0):]))
		v3 = t7
		v0 = i32(9)
	l4:
		{
			t8 := int32(m.memory[int64(uint32(v3&i32(15)))+1122552])
			m.memory[uint32(v2+i32(8)+v0+i32(-2))] = byte(t8)
			v0 = v0 + i32(-1)
			v3 = int32(uint32(v3) >> 4)
			if v3 != 0 {
				goto l4
			}
		}
		t9 := m.fn306(v1, i32(1), i32(1122550), i32(2), v2+i32(8)+v0+i32(-1), i32(9)-v0)
		v0 = t9
	}
l2:
	m.g0 = v2 + i32(16)
	return v0
}
func (m *Module) fn848(v0, v1 int32) int32 {
	var v2, v3 int32
	t0 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	v2 = t0
	t1 := int32(load32(m.memory[uint32(v0):]))
	v3 = t1
	{
		t2 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		v0 = t2
		t3 := int32(m.memory[uint32(v0)])
		if t3 == 0 {
			goto l0
		}
		t4 := int32(load32(m.memory[int64(uint32(v2))+12:]))
		t5 := m.t0[uint(t4)].(func(int32, int32, int32) int32)(v3, i32(1121988), i32(4))
		if t5 == 0 {
			goto l0
		}
		return i32(1)
	}
l0:
	t6 := v0
	var p7 int32
	if v1 == i32(10) {
		p7 = 1
	}
	m.memory[uint32(t6)] = byte(p7)
	t8 := int32(load32(m.memory[int64(uint32(v2))+16:]))
	t9 := m.t0[uint(t8)].(func(int32, int32) int32)(v3, v1)
	return t9
}
func (m *Module) fn849(v0, v1, v2 int32) int32 {
	t0 := m.fn45(v0, i32(1099920), v1, v2)
	return t0
}
func (m *Module) fn850(v0, v1, v2, v3, v4 int32) int32 {
	{
		if v2 == i32(-1) {
			goto l0
		}
		t0 := int32(load32(m.memory[int64(uint32(v1))+16:]))
		t1 := m.t0[uint(t0)].(func(int32, int32) int32)(v0, v2)
		if t1 == 0 {
			goto l0
		}
		return i32(1)
	}
l0:
	if v3 != 0 {
		t2 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		t3 := m.t0[uint(t2)].(func(int32, int32, int32) int32)(v0, v3, v4)
		return t3
	}
	return i32(0)
}
func (m *Module) fn851(v0, v1 int32) {
	var v2 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	store32(m.memory[int64(uint32(v2))+12:], uint32(v1))
	store32(m.memory[int64(uint32(v2))+8:], uint32(v0))
	m.fn842(i32(0), v2+i32(8), i32(1098848), v2+i32(12), i32(1098848), i32(0), v2, i32(1099800))
	panic("unreachable")
}
