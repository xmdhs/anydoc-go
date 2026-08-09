package core

import (
	"math/bits"
)

func (m *Module) fn42(v0 int32) {
	var v1 int32
	var v2 int64
	t0 := m.g0
	v1 = t0 - i32(32)
	m.g0 = v1
	store32(m.memory[int64(uint32(v1))+12:], uint32(v0))
	store32(m.memory[int64(uint32(v1))+8:], uint32(i32(0)))
	t1 := v1
	v2 = int64(uint32(i32(3))) << 32
	store64(m.memory[int64(uint32(t1))+24:], uint64(v2|int64(uint32(v1+i32(12)))))
	store64(m.memory[int64(uint32(v1))+16:], uint64(v2|int64(uint32(v1+i32(8)))))
	m.fn27(i32(1067030), v1+i32(16), i32(1078788))
	panic("unreachable")
}
func (m *Module) fn43(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7, v8 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	v3 = i32(10)
	t1 := int32(load32(m.memory[uint32(v0):]))
	v4 = t1
	v5 = v4
	if uint32(v4) < uint32(i32(1000)) {
		goto l0
	}
	v3 = i32(10)
	v5 = v4
l1:
	{
		v6 = v2 + i32(6) + v3
		t2 := v6 + i32(-4)
		v0 = v5
		t3 := int32(uint32(v0) / uint32(i32(10000)))
		t4 := v0
		v5 = t3
		v7 = t4 - v5*i32(10000)
		t5 := int32(uint32(v7&i32(0xffff)) / uint32(i32(100)))
		v8 = t5
		t6 := int32(load16(m.memory[int64(uint32(v8<<1))+1100623:]))
		store16(m.memory[uint32(t2):], uint16(t6))
		t7 := int32(load16(m.memory[int64(uint32((v7-v8*i32(100))&i32(0xffff)<<1))+1100623:]))
		store16(m.memory[uint32(v6+i32(-2)):], uint16(t7))
		v3 = v3 + i32(-4)
		if uint32(v0) > uint32(i32(9999999)) {
			goto l1
		}
	}
l0:
	{
		if uint32(v5) > uint32(i32(9)) {
			goto l2
		}
		v0 = v5
		goto l3
	l2:
		t8 := v2 + i32(6)
		v3 = v3 + i32(-2)
		t9 := int32(uint32(v5&i32(0xffff)) / uint32(i32(100)))
		t10 := t8 + v3
		t11 := v5
		v0 = t9
		t12 := int32(load16(m.memory[int64(uint32((t11-v0*i32(100))&i32(0xffff)<<1))+1100623:]))
		store16(m.memory[uint32(t10):], uint16(t12))
	}
l3:
	{
		if v4 == 0 {
			goto l4
		}
		if v0 == 0 {
			goto l5
		}
	l4:
		t13 := v2 + i32(6)
		v3 = v3 + i32(-1)
		t14 := int32(m.memory[int64(uint32(v0<<1))+1100624])
		m.memory[uint32(t13+v3)] = byte(t14)
	}
l5:
	t15 := m.fn679(v1, i32(1), i32(1), i32(0), v2+i32(6)+v3, i32(10)-v3)
	v3 = t15
	m.g0 = v2 + i32(16)
	return v3
}
func (m *Module) fn44(v0, v1, v2 int32) {
	var v3 int32
	var v4 int64
	t0 := m.g0
	v3 = t0 - i32(32)
	m.g0 = v3
	store32(m.memory[int64(uint32(v3))+12:], uint32(v1))
	store32(m.memory[int64(uint32(v3))+8:], uint32(v0))
	t1 := v3
	v4 = int64(uint32(i32(3))) << 32
	store64(m.memory[int64(uint32(t1))+24:], uint64(v4|int64(uint32(v3+i32(12)))))
	store64(m.memory[int64(uint32(v3))+16:], uint64(v4|int64(uint32(v3+i32(8)))))
	m.fn27(i32(1066979), v3+i32(16), v2)
	panic("unreachable")
}
func (m *Module) fn45(v0, v1, v2, v3 int32) int32 {
	var v4, v5, v6, v7, v8, v9, v10, v11 int32
	t0 := m.g0
	v4 = t0 - i32(16)
	m.g0 = v4
	{
		{
			if v3&i32(1) != 0 {
				t2 := int32(load32(m.memory[int64(uint32(v1))+12:]))
				t3 := m.t0[uint(t2)].(func(int32, int32, int32) int32)(v0, v2, int32(uint32(v3)>>1))
				v5 = t3
				goto l2
			}
			t1 := int32(m.memory[uint32(v2)])
			v5 = t1
			if v5 != 0 {
				goto l1
			}
			v5 = i32(0)
			goto l2
		}
	l1:
		t4 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		v6 = t4
		v7 = i32(0)
	l20:
		{
			v8 = v2 + i32(1)
			{
				if int32(int8(v5)) > i32(-1) {
					t8 := v0
					t9 := v8
					v5 = v5 & i32(255)
					t10 := m.t0[uint(v6)].(func(int32, int32, int32) int32)(t8, t9, v5)
					if t10 != 0 {
						v5 = i32(1)
						goto l2
					}
					v2 = v8 + v5
					goto l8
				}
				v9 = v5 & i32(255)
				if v9 == i32(128) {
					t11 := v0
					v5 = v2 + i32(3)
					t12 := int32(load16(m.memory[int64(uint32(v2))+1:]))
					t13 := v5
					v2 = t12
					t14 := m.t0[uint(v6)].(func(int32, int32, int32) int32)(t11, t13, v2)
					if t14 != 0 {
						v5 = i32(1)
						goto l2
					}
					v2 = v5 + v2
					goto l8
				}
				if v9 != i32(192) {
					goto l5
				}
				store32(m.memory[int64(uint32(v4))+4:], uint32(v1))
				store32(m.memory[uint32(v4):], uint32(v0))
				store64(m.memory[int64(uint32(v4))+8:], uint64(i64(0x60000020)))
				v5 = v3 + v7<<3
				t5 := int32(load32(m.memory[uint32(v5):]))
				t6 := int32(load32(m.memory[int64(uint32(v5))+4:]))
				t7 := m.t0[uint(t6)].(func(int32, int32) int32)(t5, v4)
				if t7 == 0 {
					v7 = v7 + i32(1)
					v2 = v8
					goto l8
				}
				v5 = i32(1)
				goto l2
			}
		l5:
			v10 = i32(0x60000020)
			{
				if v5&i32(1) == 0 {
					goto l10
				}
				v8 = v2 + i32(5)
				t15 := int32(load32(m.memory[int64(uint32(v2))+1:]))
				v10 = t15
			}
		l10:
			v9 = i32(0)
			{
				if v5&i32(2) != 0 {
					goto l11
				}
				v11 = i32(0)
				v2 = v8
				goto l12
			l11:
				v2 = v8 + i32(2)
				t16 := int32(load16(m.memory[uint32(v8):]))
				v11 = t16
			}
		l12:
			{
				if v5&i32(4) != 0 {
					goto l13
				}
				v8 = v2
				goto l14
			l13:
				v8 = v2 + i32(2)
				t17 := int32(load16(m.memory[uint32(v2):]))
				v9 = t17
			}
		l14:
			{
				if v5&i32(8) != 0 {
					goto l15
				}
				v2 = v8
				goto l16
			l15:
				v2 = v8 + i32(2)
				t18 := int32(load16(m.memory[uint32(v8):]))
				v7 = t18
			}
		l16:
			{
				if v5&i32(16) == 0 {
					goto l17
				}
				t19 := int32(load16(m.memory[int64(uint32(v3+v11&i32(0xffff)<<3))+4:]))
				v11 = t19
			}
		l17:
			{
				if v5&i32(32) == 0 {
					goto l18
				}
				t20 := int32(load16(m.memory[int64(uint32(v3+v9&i32(0xffff)<<3))+4:]))
				v9 = t20
			}
		l18:
			store16(m.memory[int64(uint32(v4))+14:], uint16(v9))
			store16(m.memory[int64(uint32(v4))+12:], uint16(v11))
			store32(m.memory[int64(uint32(v4))+8:], uint32(v10))
			store32(m.memory[int64(uint32(v4))+4:], uint32(v1))
			store32(m.memory[uint32(v4):], uint32(v0))
			{
				v5 = v3 + v7<<3
				t21 := int32(load32(m.memory[uint32(v5):]))
				t22 := int32(load32(m.memory[int64(uint32(v5))+4:]))
				t23 := m.t0[uint(t22)].(func(int32, int32) int32)(t21, v4)
				if t23 == 0 {
					goto l19
				}
				v5 = i32(1)
				goto l2
			}
		l19:
			v7 = v7 + i32(1)
		l8:
			t24 := int32(m.memory[uint32(v2)])
			v5 = t24
			if v5 != 0 {
				goto l20
			}
		}
		v5 = i32(0)
	}
l2:
	m.g0 = v4 + i32(16)
	return v5
}
func (m *Module) fn46(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v1):]))
	t1 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	t2 := int32(load32(m.memory[int64(uint32(t1))+12:]))
	t3 := m.t0[uint(t2)].(func(int32, int32, int32) int32)(t0, i32(1275037), i32(5))
	return t3
}
func (m *Module) fn47(v0 int32) {
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
			m.fn3(i32(1274224), i32(46), i32(1274272))
			panic("unreachable")
		}
		if v0 == 0 {
			goto l2
		}
		if uint32(v3) > uint32(v1+i32(39)) {
			m.fn3(i32(1274288), i32(46), i32(1274336))
			panic("unreachable")
		}
	l2:
		m.fn1(v2)
	}
}
func (m *Module) fn48(v0, v1, v2 int32) int32 {
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
			m.fn24(v0, v3, v2)
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
func (m *Module) fn49(v0, v1 int32) int32 {
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
		m.fn24(v0, v2, v3)
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
func (m *Module) fn50(v0, v1, v2 int32) int32 {
	t0 := m.fn45(v0, i32(1068076), v1, v2)
	return t0
}
func (m *Module) fn51(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t1 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	t2 := int32(load32(m.memory[uint32(v1):]))
	t3 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	t4 := m.fn52(t0, t1, t2, t3)
	return t4
}
func (m *Module) fn52(v0, v1, v2, v3 int32) int32 {
	var v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14 int32
	t0 := m.g0
	v4 = t0 - i32(16)
	m.g0 = v4
	v5 = i32(1)
	{
		t1 := int32(load32(m.memory[int64(uint32(v3))+16:]))
		t2 := v2
		v6 = t1
		t3 := m.t0[uint(v6)].(func(int32, int32) int32)(t2, i32(34))
		if t3 != 0 {
			goto l0
		}
		{
			if v1 != 0 {
				goto l1
			}
			v1 = i32(0)
			v7 = i32(0)
			goto l2
		l1:
			v8 = i32(0)
			v9 = i32(0)
			v10 = v1
			v11 = v0
		l25:
			v12 = v11 + v10
			v7 = i32(0)
		l4:
			{
				v13 = v11 + v7
				t4 := int32(m.memory[uint32(v13)])
				v14 = t4
				if uint32((v14+i32(-127))&i32(255)) < uint32(i32(161)) {
					goto l3
				}
				if v14 == i32(34) {
					goto l3
				}
				if v14 == i32(92) {
					goto l3
				}
				t5 := v10
				v7 = v7 + i32(1)
				if t5 != v7 {
					goto l4
				}
			}
			v9 = v9 + v10
			goto l5
		l3:
			{
				{
					t6 := int32(int8(m.memory[uint32(v13)]))
					v14 = t6
					if v14 <= i32(-1) {
						goto l6
					}
					v11 = v13 + i32(1)
					v14 = v14 & i32(255)
					goto l7
				}
			l6:
				t7 := int32(m.memory[int64(uint32(v13))+1])
				v10 = t7 & i32(63)
				v11 = v14 & i32(31)
				if uint32(v14) > uint32(i32(-33)) {
					goto l8
				}
				v14 = v11<<6 | v10
				v11 = v13 + i32(2)
				goto l7
			l8:
				t8 := int32(m.memory[int64(uint32(v13))+2])
				v10 = v10<<6 | t8&i32(63)
				if uint32(v14) >= uint32(i32(-16)) {
					goto l9
				}
				v14 = v10 | v11<<12
				v11 = v13 + i32(3)
				goto l7
			l9:
				t9 := int32(m.memory[int64(uint32(v13))+3])
				v14 = v10<<6 | t9&i32(63) | v11<<18&i32(0x1c0000)
				v11 = v13 + i32(4)
			}
		l7:
			v7 = v7 + v9
			m.fn852(v4, v14, i32(65537))
			{
				t10 := int32(m.memory[int64(uint32(v4))+13])
				v13 = t10
				t11 := int32(m.memory[int64(uint32(v4))+12])
				t12 := v13
				v10 = t11
				v9 = t12 - v10
				if v9&i32(255) == i32(1) {
					goto l10
				}
				{
					if uint32(v7) < uint32(v8) {
						goto l11
					}
					{
						if v8 == 0 {
							goto l12
						}
						if uint32(v8) < uint32(v1) {
							goto l13
						}
						if v8 != v1 {
							goto l11
						}
						goto l12
					l13:
						t13 := int32(int8(m.memory[uint32(v0+v8)]))
						if t13 <= i32(-65) {
							goto l11
						}
					}
				l12:
					{
						if v7 == 0 {
							goto l14
						}
						if uint32(v7) < uint32(v1) {
							goto l15
						}
						if v7 == v1 {
							goto l14
						}
						goto l11
					l15:
						t14 := int32(int8(m.memory[uint32(v0+v7)]))
						if t14 <= i32(-65) {
							goto l11
						}
					}
				l14:
					t15 := int32(load32(m.memory[int64(uint32(v3))+12:]))
					t16 := v2
					t17 := v0 + v8
					t18 := v7 - v8
					v8 = t15
					t19 := m.t0[uint(v8)].(func(int32, int32, int32) int32)(t16, t17, t18)
					if t19 == 0 {
						{
							{
								if uint32(v13) < uint32(i32(129)) {
									goto l18
								}
								t20 := int32(load32(m.memory[uint32(v4):]))
								t21 := m.t0[uint(v6)].(func(int32, int32) int32)(v2, t20)
								if t21 != 0 {
									goto l17
								}
								goto l19
							}
						l18:
							t22 := m.t0[uint(v8)].(func(int32, int32, int32) int32)(v2, v4+v10, v9)
							if t22 != 0 {
								goto l17
							}
						}
					l19:
						if uint32(v14) >= uint32(i32(128)) {
							if uint32(v14) >= uint32(i32(2048)) {
								p23 := i32(4)
								if uint32(v14) < uint32(i32(65536)) {
									p23 = i32(3)
								}
								v8 = p23 + v7
								goto l10
							}
							v8 = i32(2) + v7
							goto l10
						}
						v8 = i32(1) + v7
						goto l10
					}
					goto l17
				}
			l11:
				m.fn37(v0, v1, v8, v7, i32(1123052))
				panic("unreachable")
			l17:
				v5 = i32(1)
				goto l0
			}
		l10:
			{
				if uint32(v14) >= uint32(i32(128)) {
					goto l22
				}
				v14 = i32(1)
				goto l23
			l22:
				if uint32(v14) >= uint32(i32(2048)) {
					goto l24
				}
				v14 = i32(2)
				goto l23
			l24:
				p24 := i32(4)
				if uint32(v14) < uint32(i32(65536)) {
					p24 = i32(3)
				}
				v14 = p24
			}
		l23:
			v9 = v14 + v7
			v10 = v12 - v11
			if v10 != 0 {
				goto l25
			}
		l5:
			if uint32(v8) > uint32(v9) {
				goto l26
			}
			v7 = i32(0)
			{
				if v8 == 0 {
					goto l27
				}
				if uint32(v8) < uint32(v1) {
					goto l28
				}
				v7 = v1
				if v8 != v1 {
					goto l26
				}
				goto l27
			l28:
				v7 = v8
				t25 := int32(int8(m.memory[uint32(v0+v8)]))
				if t25 <= i32(-65) {
					goto l26
				}
			}
		l27:
			if v9 != 0 {
				goto l29
			}
			v1 = i32(0)
			goto l2
		l29:
			if uint32(v9) < uint32(v1) {
				goto l30
			}
			if v9 == v1 {
				goto l2
			}
			v8 = v7
			goto l26
		l30:
			{
				t26 := int32(int8(m.memory[uint32(v0+v9)]))
				if t26 > i32(-65) {
					goto l31
				}
				v8 = v7
				goto l26
			}
		l31:
			v1 = v9
		l2:
			t27 := int32(load32(m.memory[int64(uint32(v3))+12:]))
			t28 := m.t0[uint(t27)].(func(int32, int32, int32) int32)(v2, v0+v7, v1-v7)
			if t28 != 0 {
				goto l0
			}
			t29 := m.t0[uint(v6)].(func(int32, int32) int32)(v2, i32(34))
			v5 = t29
			goto l0
		}
	l26:
		m.fn37(v0, v1, v8, v9, i32(1123068))
		panic("unreachable")
	}
l0:
	m.g0 = v4 + i32(16)
	return v5
}
func (m *Module) fn53(v0, v1, v2 int32) {
	var v3 int32
	{
		if v2 != 0 {
			goto l0
		}
		v3 = i32(1)
		goto l1
	l0:
		t0 := m.fn7(v2)
		v3 = t0
		if v3 == 0 {
			m.fn12(i32(1), v2)
			panic("unreachable")
		}
		if v2 == 0 {
			goto l1
		}
		memory_copy(m.memory, uint32(v3), uint32(v1), uint32(v2))
	}
l1:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v2))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
	store32(m.memory[uint32(v0):], uint32(v2))
}
func (m *Module) fn54(v0, v1 int32) {
	var v2, v3, v4, v5 int32
	t0 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	v2 = t0
	{
		t1 := int32(load32(m.memory[uint32(v1):]))
		v3 = t1
		t2 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		t3 := v3
		v1 = t2
		if uint32(t3) <= uint32(v1) {
			goto l0
		}
		{
			if v1 != 0 {
				goto l1
			}
			t4 := int32(load32(m.memory[uint32(v2+i32(-4)):]))
			v4 = t4
			v5 = v4 & i32(-8)
			t5 := v5
			v4 = v4 & i32(3)
			p6 := i32(8)
			if v4 != 0 {
				p6 = i32(4)
			}
			if uint32(t5) < uint32(p6+v3) {
				m.fn3(i32(1274224), i32(46), i32(1274272))
				panic("unreachable")
			}
			if v4 == 0 {
				goto l3
			}
			if uint32(v5) > uint32(v3+i32(39)) {
				m.fn3(i32(1274288), i32(46), i32(1274336))
				panic("unreachable")
			}
		l3:
			m.fn1(v2)
			v2 = i32(1)
			goto l0
		}
	l1:
		t7 := m.fn21(v2, v3, i32(1), v1)
		v2 = t7
		if v2 == 0 {
			m.fn12(i32(1), v1)
			panic("unreachable")
		}
	}
l0:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
	store32(m.memory[uint32(v0):], uint32(v2))
}
func (m *Module) fn55(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t1 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	t2 := m.fn5(v1, t0, t1)
	return t2
}
func (m *Module) fn56(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7, v8, v9, v10 int32
	var v11, v12, v13 int64
	var v14, v15 int32
	var v16, v17 int64
	var v18 int32
	{
		{
			{
				t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
				v2 = t0
				v3 = v2 + i32(1)
				if v3 == 0 {
					m.fn27(i32(1271632), i32(57), i32(1271660))
					panic("unreachable")
				}
				{
					t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
					t2 := v3
					v4 = t1
					t3 := v4
					v5 = v4 + i32(1)
					v6 = int32(uint32(v5) >> 3)
					v7 = v6 * i32(7)
					p4 := v7
					if uint32(v4) < uint32(i32(8)) {
						p4 = t3
					}
					v8 = p4
					if uint32(t2) <= uint32(int32(uint32(v8)>>1)) {
						{
							if v5 != 0 {
								goto l6
							}
							v3 = i32(0)
							goto l7
						l6:
							t7 := int32(load32(m.memory[uint32(v0):]))
							v8 = t7
							v3 = i32(0)
							{
								{
									t8 := v6
									var p9 int32
									if v5&i32(7) != i32(0) {
										p9 = 1
									}
									v6 = t8 + p9
									if v6 == i32(1) {
										goto l8
									}
									v9 = v6 & i32(1)
									v10 = v6 & i32(0x3ffffffe)
									v3 = i32(0)
								l9:
									{
										v6 = v8 + v3
										t10 := int64(load64(m.memory[uint32(v6):]))
										t11 := v6
										v11 = t10
										store64(m.memory[uint32(t11):], uint64(int64(uint64(v11^i64(-1))>>7)&i64(72340172838076673)+(v11|i64(0x7f7f7f7f7f7f7f7f))))
										v6 = v6 + i32(8)
										t12 := int64(load64(m.memory[uint32(v6):]))
										t13 := v6
										v11 = t12
										store64(m.memory[uint32(t13):], uint64(int64(uint64(v11^i64(-1))>>7)&i64(72340172838076673)+(v11|i64(0x7f7f7f7f7f7f7f7f))))
										v3 = v3 + i32(16)
										v10 = v10 + i32(-2)
										if v10 != 0 {
											goto l9
										}
									}
									if v9 == 0 {
										goto l10
									}
								}
							l8:
								v3 = v8 + v3
								t14 := int64(load64(m.memory[uint32(v3):]))
								t15 := v3
								v11 = t14
								store64(m.memory[uint32(t15):], uint64(int64(uint64(v11^i64(-1))>>7)&i64(72340172838076673)+(v11|i64(0x7f7f7f7f7f7f7f7f))))
							}
						l10:
							{
								if uint32(v5) < uint32(i32(8)) {
									goto l11
								}
								t16 := int64(load64(m.memory[uint32(v8):]))
								store64(m.memory[uint32(v8+v5):], uint64(t16))
								goto l12
							}
						l11:
							if v5 == 0 {
								goto l12
							}
							memory_copy(m.memory, uint32(v8+i32(8)), uint32(v8), uint32(v5))
						l12:
							t17 := int64(load64(m.memory[int64(uint32(v1))+8:]))
							v12 = t17
							t18 := int64(load64(m.memory[uint32(v1):]))
							v13 = t18
							v6 = i32(0)
						l20:
							{
								t19 := v8
								v3 = v6
								v10 = t19 + v3
								t20 := int32(m.memory[uint32(v10)])
								if t20 != i32(128) {
									goto l13
								}
								v1 = v8 - v3<<4 + i32(-16)
								v14 = v8 + (v3^i32(-1))<<4
							l19:
								{
									t21 := m.fn57(v13, v12, v1)
									t22 := v4
									v9 = int32(t21)
									v6 = t22 & v9
									v5 = v6
									{
										t23 := int64(load64(m.memory[uint32(v8+v6):]))
										v11 = t23 & i64(-0x7f7f7f7f7f7f7f80)
										if v11 != i64(0) {
											goto l14
										}
										v15 = i32(8)
										v5 = v6
									l15:
										{
											v5 = v5 + v15
											v15 = v15 + i32(8)
											t24 := v8
											v5 = v5 & v4
											t25 := int64(load64(m.memory[uint32(t24+v5):]))
											v11 = t25 & i64(-0x7f7f7f7f7f7f7f80)
											if v11 == 0 {
												goto l15
											}
										}
									}
								l14:
									{
										t26 := v8
										v5 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v11))))>>3) + v5) & v4
										t27 := int32(int8(m.memory[uint32(t26+v5)]))
										if t27 < i32(0) {
											goto l16
										}
										t28 := int64(load64(m.memory[uint32(v8):]))
										v5 = int32(uint32(int64(bits.TrailingZeros64(uint64(t28&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
									}
								l16:
									{
										if uint32((v5-v6^(v3-v6))&v4) < uint32(i32(8)) {
											goto l17
										}
										v6 = v8 + v5
										t29 := int32(m.memory[uint32(v6)])
										v15 = t29
										t30 := v6
										v9 = int32(uint32(v9) >> 25)
										m.memory[uint32(t30)] = byte(v9)
										m.memory[uint32(v8+(v5+i32(-8))&v4+i32(8))] = byte(v9)
										v6 = v8 + (v5^i32(-1))<<4
										{
											if v15 != i32(255) {
												t33 := int64(load64(m.memory[uint32(v14):]))
												v11 = t33
												t34 := int64(load64(m.memory[uint32(v6):]))
												store64(m.memory[uint32(v14):], uint64(t34))
												store64(m.memory[uint32(v6):], uint64(v11))
												t35 := int64(load64(m.memory[int64(uint32(v14))+8:]))
												v11 = t35
												t36 := int64(load64(m.memory[int64(uint32(v6))+8:]))
												store64(m.memory[int64(uint32(v14))+8:], uint64(t36))
												store64(m.memory[int64(uint32(v6))+8:], uint64(v11))
												goto l19
											}
											m.memory[uint32(v10)] = byte(i32(255))
											m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(i32(255))
											t31 := int64(load64(m.memory[int64(uint32(v14))+8:]))
											store64(m.memory[int64(uint32(v6))+8:], uint64(t31))
											t32 := int64(load64(m.memory[uint32(v14):]))
											store64(m.memory[uint32(v6):], uint64(t32))
											goto l13
										}
									}
								l17:
								}
								t37 := v10
								v6 = int32(uint32(v9) >> 25)
								m.memory[uint32(t37)] = byte(v6)
								m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(v6)
							}
						l13:
							v6 = v3 + i32(1)
							if v3 != v4 {
								goto l20
							}
							p38 := v7
							if uint32(v4) < uint32(i32(8)) {
								p38 = v4
							}
							v3 = p38
						}
					l7:
						store32(m.memory[int64(uint32(v0))+8:], uint32(v3-v2))
						goto l21
					}
					v8 = v8 + i32(1)
					p5 := v3
					if uint32(v8) > uint32(v3) {
						p5 = v8
					}
					v3 = p5
					if uint32(v3) < uint32(i32(15)) {
						goto l2
					}
					{
						if uint32(v3) > uint32(i32(0x1fffffff)) {
							m.fn27(i32(1271632), i32(57), i32(1271660))
							panic("unreachable")
						}
						t6 := int32(uint32(v3<<3) / uint32(i32(7)))
						v3 = i32_shr_u(i32(-1), int32(bits.LeadingZeros32(uint32(t6+i32(-1)))))
						if uint32(v3) > uint32(i32(0xffffffe)) {
							goto l4
						}
						v3 = v3 + i32(1)
						goto l5
					}
				}
			}
		l2:
			p39 := v3&i32(8) + i32(8)
			if uint32(v3) < uint32(i32(4)) {
				p39 = i32(4)
			}
			v3 = p39
		}
	l5:
		v8 = v3 + i32(8)
		t40 := v8
		v10 = v3 << 4
		v6 = t40 + v10
		if uint32(v6) < uint32(v8) {
			goto l4
		}
		if uint32(v6) > uint32(i32(0x7ffffff8)) {
			goto l4
		}
		{
			t41 := m.fn7(v6)
			v5 = t41
			if v5 != 0 {
				v6 = v5 + v10
				if v8 == 0 {
					goto l23
				}
				memory_fill(m.memory, uint32(v6), i32(255), uint32(v8))
			l23:
				v5 = v3 + i32(-1)
				p42 := int32(uint32(v3)>>3) * i32(7)
				if uint32(v3) < uint32(i32(9)) {
					p42 = v5
				}
				v15 = p42
				t43 := int32(load32(m.memory[uint32(v0):]))
				v9 = t43
				{
					if v2 == 0 {
						goto l24
					}
					t44 := int64(load64(m.memory[uint32(v9):]))
					v11 = (t44 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
					t45 := int64(load64(m.memory[int64(uint32(v1))+8:]))
					v16 = t45
					t46 := int64(load64(m.memory[uint32(v1):]))
					v17 = t46
					v8 = v9
					v1 = v2
					v3 = i32(0)
				l30:
					{
						if v11 != i64(0) {
							goto l25
						}
					l26:
						{
							v3 = v3 + i32(8)
							v8 = v8 + i32(8)
							t47 := int64(load64(m.memory[uint32(v8):]))
							v11 = t47 & i64(-0x7f7f7f7f7f7f7f80)
							if v11 == i64(-0x7f7f7f7f7f7f7f80) {
								goto l26
							}
						}
						v11 = v11 ^ i64(-0x7f7f7f7f7f7f7f80)
					l25:
						{
							t48 := v6
							t49 := v5
							t50 := v17
							t51 := v16
							t52 := v9
							v14 = int32(uint32(int64(bits.TrailingZeros64(uint64(v11))))>>3) + v3
							t53 := m.fn57(t50, t51, t52-v14<<4+i32(-16))
							v7 = int32(t53)
							v10 = t49 & v7
							t54 := int64(load64(m.memory[uint32(t48+v10):]))
							v12 = t54 & i64(-0x7f7f7f7f7f7f7f80)
							if v12 != i64(0) {
								goto l27
							}
							v18 = i32(8)
						l28:
							{
								v10 = v10 + v18
								v18 = v18 + i32(8)
								t55 := v6
								v10 = v10 & v5
								t56 := int64(load64(m.memory[uint32(t55+v10):]))
								v12 = t56 & i64(-0x7f7f7f7f7f7f7f80)
								if v12 == 0 {
									goto l28
								}
							}
						}
					l27:
						v13 = v11 + i64(-1)
						{
							t57 := v6
							v10 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v12))))>>3) + v10) & v5
							t58 := int32(int8(m.memory[uint32(t57+v10)]))
							if t58 < i32(0) {
								goto l29
							}
							t59 := int64(load64(m.memory[uint32(v6):]))
							v10 = int32(uint32(int64(bits.TrailingZeros64(uint64(t59&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
						}
					l29:
						v11 = v13 & v11
						t60 := v6 + v10
						v7 = int32(uint32(v7) >> 25)
						m.memory[uint32(t60)] = byte(v7)
						m.memory[uint32(v6+(v10+i32(-8))&v5+i32(8))] = byte(v7)
						v10 = v6 + (v10^i32(-1))<<4
						t61 := v10
						v14 = v9 + (v14^i32(-1))<<4
						t62 := int64(load64(m.memory[int64(uint32(v14))+8:]))
						store64(m.memory[int64(uint32(t61))+8:], uint64(t62))
						t63 := int64(load64(m.memory[uint32(v14):]))
						store64(m.memory[uint32(v10):], uint64(t63))
						v1 = v1 + i32(-1)
						if v1 != 0 {
							goto l30
						}
					}
				}
			l24:
				store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
				store32(m.memory[uint32(v0):], uint32(v6))
				store32(m.memory[int64(uint32(v0))+8:], uint32(v15-v2))
				if v4 == 0 {
					goto l21
				}
				t64 := v4
				v8 = (v4<<4 + i32(23)) & i32(-16)
				v3 = t64 + v8 + i32(9)
				if v3 == 0 {
					goto l21
				}
				v4 = v9 - v8
				t65 := int32(load32(m.memory[uint32(v4+i32(-4)):]))
				v8 = t65
				v6 = v8 & i32(-8)
				t66 := v6
				v8 = v8 & i32(3)
				p67 := i32(8)
				if v8 != 0 {
					p67 = i32(4)
				}
				if uint32(t66) < uint32(p67+v3) {
					m.fn3(i32(1274224), i32(46), i32(1274272))
					panic("unreachable")
				}
				if v8 == 0 {
					goto l32
				}
				if uint32(v6) > uint32(v3+i32(39)) {
					m.fn3(i32(1274288), i32(46), i32(1274336))
					panic("unreachable")
				}
			l32:
				m.fn1(v4)
				return i32(-1)
			}
			m.fn23(i32(8), v6)
			panic("unreachable")
		}
	}
l21:
	return i32(-1)
l4:
	m.fn27(i32(1271632), i32(57), i32(1271660))
	panic("unreachable")
}
func (m *Module) fn57(v0, v1 int64, v2 int32) int64 {
	var v3, v4 int32
	var v5, v6, v7, v8 int64
	t0 := m.g0
	v3 = t0 - i32(80)
	m.g0 = v3
	store64(m.memory[int64(uint32(v3))+56:], uint64(i64(0)))
	store64(m.memory[int64(uint32(v3))+64:], uint64(i64(0)))
	store64(m.memory[int64(uint32(v3))+48:], uint64(v1))
	store64(m.memory[int64(uint32(v3))+32:], uint64(v1^i64(8387220255154660723)))
	store64(m.memory[int64(uint32(v3))+24:], uint64(v1^i64(7237128888997146477)))
	store64(m.memory[int64(uint32(v3))+40:], uint64(v0))
	store64(m.memory[int64(uint32(v3))+16:], uint64(v0^i64(0x6c7967656e657261)))
	store64(m.memory[int64(uint32(v3))+8:], uint64(v0^i64(8317987319222330741)))
	t1 := int32(load32(m.memory[uint32(v2):]))
	t2 := v3
	v4 = t1
	var p3 int32
	if v4 != i32(-1) {
		p3 = 1
	}
	store32(m.memory[int64(uint32(t2))+76:], uint32(p3))
	m.fn58(v3+i32(8), v3+i32(76), i32(4))
	{
		if v4 == i32(-1) {
			goto l0
		}
		t4 := int32(load32(m.memory[int64(uint32(v2))+4:]))
		t5 := int32(load32(m.memory[int64(uint32(v2))+8:]))
		m.fn58(v3+i32(8), t4, t5)
		m.memory[int64(uint32(v3))+76] = byte(i32(255))
		m.fn58(v3+i32(8), v3+i32(76), i32(1))
	}
l0:
	t6 := int64(load64(m.memory[int64(uint32(v3))+8:]))
	v0 = t6
	t7 := int64(load64(m.memory[int64(uint32(v3))+24:]))
	v1 = t7
	t8 := int64(load32(m.memory[int64(uint32(v3))+64:]))
	v5 = t8
	t9 := int64(load64(m.memory[int64(uint32(v3))+56:]))
	v6 = t9
	t10 := int64(load64(m.memory[int64(uint32(v3))+32:]))
	v7 = t10
	t11 := int64(load64(m.memory[int64(uint32(v3))+16:]))
	v8 = t11
	m.g0 = v3 + i32(80)
	t12 := v7
	v5 = v6 | v5<<56
	v6 = t12 ^ v5
	t13 := i64_rotl(v6, i64(16))
	v6 = v6 + v8
	v7 = t13 ^ v6
	t14 := i64_rotl(v7, i64(21))
	t15 := v7
	v0 = v1 + v0
	v7 = t15 + i64_rotl(v0, i64(32))
	v8 = t14 ^ v7
	t16 := i64_rotl(v8, i64(16))
	t17 := v8
	t18 := v6
	v1 = i64_rotl(v1, i64(13)) ^ v0
	v0 = t18 + v1
	v6 = t17 + (i64_rotl(v0, i64(32)) ^ i64(255))
	v8 = t16 ^ v6
	t19 := i64_rotl(v8, i64(21))
	t20 := v8
	t21 := v7 ^ v5
	v1 = v0 ^ i64_rotl(v1, i64(17))
	v0 = t21 + v1
	v5 = t20 + i64_rotl(v0, i64(32))
	v7 = t19 ^ v5
	t22 := i64_rotl(v7, i64(16))
	t23 := v7
	v1 = v0 ^ i64_rotl(v1, i64(13))
	v0 = v1 + v6
	v6 = t23 + i64_rotl(v0, i64(32))
	v7 = t22 ^ v6
	t24 := i64_rotl(v7, i64(21))
	t25 := v7
	v1 = v0 ^ i64_rotl(v1, i64(17))
	v0 = v1 + v5
	v5 = t25 + i64_rotl(v0, i64(32))
	v7 = t24 ^ v5
	t26 := i64_rotl(v7, i64(16))
	t27 := v7
	v1 = i64_rotl(v1, i64(13)) ^ v0
	v0 = v1 + v6
	v6 = t27 + i64_rotl(v0, i64(32))
	t28 := i64_rotl(t26^v6, i64(21))
	v1 = i64_rotl(v1, i64(17)) ^ v0
	v1 = i64_rotl(v1, i64(13)) ^ (v1 + v5)
	t29 := t28 ^ i64_rotl(v1, i64(17))
	v1 = v1 + v6
	return t29 ^ i64_rotl(v1, i64(32)) ^ v1
}
func (m *Module) fn58(v0, v1, v2 int32) {
	var v3, v4, v5, v6 int32
	var v7, v8, v9, v10, v11, v12 int64
	t0 := int32(load32(m.memory[int64(uint32(v0))+56:]))
	store32(m.memory[int64(uint32(v0))+56:], uint32(t0+v2))
	{
		t1 := int32(load32(m.memory[int64(uint32(v0))+60:]))
		v3 = t1
		if v3 != 0 {
			v5 = i32(4)
			{
				{
					v4 = i32(8) - v3
					p2 := v2
					if uint32(v4) < uint32(v2) {
						p2 = v4
					}
					v6 = p2
					if uint32(v6) >= uint32(i32(4)) {
						goto l2
					}
					v7 = i64(0)
					v5 = i32(0)
					goto l3
				}
			l2:
				t3 := int64(load32(m.memory[uint32(v1):]))
				v7 = t3
			}
		l3:
			{
				if uint32(v5|i32(1)) >= uint32(v6) {
					goto l4
				}
				t4 := int64(load16(m.memory[uint32(v1+v5):]))
				v7 = i64_shl(t4, int64(uint32(v5<<3))) | v7
				v5 = v5 | i32(2)
			}
		l4:
			{
				if uint32(v5) >= uint32(v6) {
					goto l5
				}
				t5 := int64(m.memory[uint32(v1+v5)])
				v7 = i64_shl(t5, int64(uint32(v5<<3))) | v7
			}
		l5:
			t6 := int64(load64(m.memory[int64(uint32(v0))+48:]))
			t7 := v0
			v7 = t6 | i64_shl(v7, int64(uint32(v3<<3)))
			store64(m.memory[int64(uint32(t7))+48:], uint64(v7))
			{
				if uint32(v2) < uint32(v4) {
					store32(m.memory[int64(uint32(v0))+60:], uint32(v3+v2))
					return
				}
				t8 := int64(load64(m.memory[int64(uint32(v0))+8:]))
				t9 := int64(load64(m.memory[int64(uint32(v0))+24:]))
				t10 := v0
				v8 = t9 ^ v7
				v9 = t8 + v8
				t11 := int64(load64(m.memory[int64(uint32(v0))+16:]))
				t12 := v9
				v10 = t11
				t13 := int64(load64(m.memory[uint32(v0):]))
				t14 := i64_rotl(v10, i64(13))
				v10 = v10 + t13
				v11 = t14 ^ v10
				v12 = t12 + v11
				store64(m.memory[int64(uint32(t10))+16:], uint64(v12^i64_rotl(v11, i64(17))))
				store64(m.memory[int64(uint32(v0))+8:], uint64(i64_rotl(v12, i64(32))))
				t15 := v0
				v8 = v9 ^ i64_rotl(v8, i64(16))
				t16 := i64_rotl(v8, i64(21))
				v8 = v8 + i64_rotl(v10, i64(32))
				store64(m.memory[int64(uint32(t15))+24:], uint64(t16^v8))
				store64(m.memory[uint32(v0):], uint64(v8^v7))
				goto l1
			}
		}
		v4 = i32(0)
		goto l1
	}
l1:
	v2 = v2 - v4
	v5 = v2 & i32(7)
	{
		t17 := v4
		v2 = v2 & i32(-8)
		if uint32(t17) >= uint32(v2) {
			goto l7
		}
		t18 := int64(load64(m.memory[int64(uint32(v0))+8:]))
		v8 = t18
		t19 := int64(load64(m.memory[int64(uint32(v0))+16:]))
		v7 = t19
		t20 := int64(load64(m.memory[int64(uint32(v0))+24:]))
		v9 = t20
		t21 := int64(load64(m.memory[uint32(v0):]))
		v10 = t21
	l8:
		{
			t22 := int64(load64(m.memory[uint32(v1+v4):]))
			t23 := v8
			t24 := v9
			v11 = t22
			v9 = t24 ^ v11
			v8 = t23 + v9
			t25 := v8
			t26 := i64_rotl(v7, i64(13))
			v10 = v7 + v10
			v7 = t26 ^ v10
			v12 = t25 + v7
			v7 = v12 ^ i64_rotl(v7, i64(17))
			v8 = v8 ^ i64_rotl(v9, i64(16))
			t27 := i64_rotl(v8, i64(21))
			v10 = v8 + i64_rotl(v10, i64(32))
			v9 = t27 ^ v10
			v8 = i64_rotl(v12, i64(32))
			v10 = v10 ^ v11
			v4 = v4 + i32(8)
			if uint32(v4) < uint32(v2) {
				goto l8
			}
		}
		store64(m.memory[int64(uint32(v0))+16:], uint64(v7))
		store64(m.memory[int64(uint32(v0))+24:], uint64(v9))
		store64(m.memory[int64(uint32(v0))+8:], uint64(v8))
		store64(m.memory[uint32(v0):], uint64(v10))
	}
l7:
	v2 = i32(4)
	{
		if uint32(v5) >= uint32(i32(4)) {
			goto l9
		}
		v7 = i64(0)
		v2 = i32(0)
		goto l10
	l9:
		t28 := int64(load32(m.memory[uint32(v1+v4):]))
		v7 = t28
	}
l10:
	{
		if uint32(v2|i32(1)) >= uint32(v5) {
			goto l11
		}
		t29 := int64(load16(m.memory[uint32(v1+v4+v2):]))
		v7 = i64_shl(t29, int64(uint32(v2<<3))) | v7
		v2 = v2 | i32(2)
	}
l11:
	{
		if uint32(v2) >= uint32(v5) {
			goto l12
		}
		t30 := int64(m.memory[uint32(v1+(v2+v4))])
		v7 = i64_shl(t30, int64(uint32(v2<<3))) | v7
	}
l12:
	store64(m.memory[int64(uint32(v0))+48:], uint64(v7))
	store32(m.memory[int64(uint32(v0))+60:], uint32(v5))
}
func (m *Module) fn59(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7, v8 int32
	var v9 int64
	var v10, v11 int32
	var v12, v13 int64
	var v14, v15 int32
	var v16 int64
	var v17 int32
	var v18 int64
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
		v2 = t0
		v3 = v2 + i32(1)
		if v3 == 0 {
			m.fn27(i32(1271632), i32(57), i32(1271660))
			panic("unreachable")
		}
		{
			t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t2 := v3
			v4 = t1
			t3 := v4
			v5 = v4 + i32(1)
			v6 = int32(uint32(v5) >> 3)
			v7 = v6 * i32(7)
			p4 := v7
			if uint32(v4) < uint32(i32(8)) {
				p4 = t3
			}
			v8 = p4
			if uint32(t2) <= uint32(int32(uint32(v8)>>1)) {
				goto l1
			}
			{
				{
					{
						v8 = v8 + i32(1)
						p5 := v3
						if uint32(v8) > uint32(v3) {
							p5 = v8
						}
						v3 = p5
						if uint32(v3) < uint32(i32(15)) {
							goto l2
						}
						if uint32(v3) > uint32(i32(0x1fffffff)) {
							m.fn27(i32(1271632), i32(57), i32(1271660))
							panic("unreachable")
						}
						t6 := int32(uint32(v3<<3) / uint32(i32(7)))
						v3 = i32_shr_u(i32(-1), int32(bits.LeadingZeros32(uint32(t6+i32(-1))))) + i32(1)
						goto l4
					}
				l2:
					p7 := v3&i32(8) + i32(8)
					if uint32(v3) < uint32(i32(4)) {
						p7 = i32(4)
					}
					v3 = p7
				}
			l4:
				v9 = int64(uint32(v3)) * i64(20)
				if int32(int64(uint64(v9)>>32)) != 0 {
					goto l5
				}
				v6 = int32(v9)
				if uint32(v6) > uint32(i32(-8)) {
					goto l5
				}
				v8 = v3 + i32(8)
				t8 := v8
				v10 = (v6 + i32(7)) & i32(-8)
				v6 = t8 + v10
				if uint32(v6) < uint32(v8) {
					goto l5
				}
				if uint32(v6) > uint32(i32(0x7ffffff8)) {
					goto l5
				}
				t9 := m.fn7(v6)
				v5 = t9
				if v5 != 0 {
					v6 = v5 + v10
					if v8 == 0 {
						goto l7
					}
					memory_fill(m.memory, uint32(v6), i32(255), uint32(v8))
				l7:
					v5 = v3 + i32(-1)
					p10 := int32(uint32(v3)>>3) * i32(7)
					if uint32(v3) < uint32(i32(9)) {
						p10 = v5
					}
					v7 = p10
					t11 := int32(load32(m.memory[uint32(v0):]))
					v11 = t11
					{
						if v2 == 0 {
							goto l8
						}
						t12 := int64(load64(m.memory[uint32(v11):]))
						v9 = (t12 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
						t13 := int64(load64(m.memory[int64(uint32(v1))+8:]))
						v12 = t13
						t14 := int64(load64(m.memory[uint32(v1):]))
						v13 = t14
						v8 = v11
						v14 = v2
						v3 = i32(0)
					l14:
						{
							if v9 != i64(0) {
								goto l9
							}
						l10:
							{
								v3 = v3 + i32(8)
								v8 = v8 + i32(8)
								t15 := int64(load64(m.memory[uint32(v8):]))
								v9 = t15 & i64(-0x7f7f7f7f7f7f7f80)
								if v9 == i64(-0x7f7f7f7f7f7f7f80) {
									goto l10
								}
							}
							v9 = v9 ^ i64(-0x7f7f7f7f7f7f7f80)
						l9:
							{
								t16 := v6
								t17 := v5
								t18 := v13
								t19 := v12
								t20 := v11
								v1 = int32(uint32(int64(bits.TrailingZeros64(uint64(v9))))>>3) + v3
								v10 = t20 + (i32(0)-v1)*i32(20)
								t21 := int32(load32(m.memory[uint32(v10+i32(-16)):]))
								t22 := int32(load32(m.memory[uint32(v10+i32(-12)):]))
								t23 := m.fn60(t18, t19, t21, t22)
								v15 = int32(t23)
								v10 = t17 & v15
								t24 := int64(load64(m.memory[uint32(t16+v10):]))
								v16 = t24 & i64(-0x7f7f7f7f7f7f7f80)
								if v16 != i64(0) {
									goto l11
								}
								v17 = i32(8)
							l12:
								{
									v10 = v10 + v17
									v17 = v17 + i32(8)
									t25 := v6
									v10 = v10 & v5
									t26 := int64(load64(m.memory[uint32(t25+v10):]))
									v16 = t26 & i64(-0x7f7f7f7f7f7f7f80)
									if v16 == 0 {
										goto l12
									}
								}
							}
						l11:
							v18 = v9 + i64(-1)
							{
								t27 := v6
								v10 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v16))))>>3) + v10) & v5
								t28 := int32(int8(m.memory[uint32(t27+v10)]))
								if t28 < i32(0) {
									goto l13
								}
								t29 := int64(load64(m.memory[uint32(v6):]))
								v10 = int32(uint32(int64(bits.TrailingZeros64(uint64(t29&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
							}
						l13:
							v9 = v18 & v9
							t30 := v6 + v10
							v15 = int32(uint32(v15) >> 25)
							m.memory[uint32(t30)] = byte(v15)
							m.memory[uint32(v6+(v10+i32(-8))&v5+i32(8))] = byte(v15)
							v10 = v6 + (v10^i32(-1))*i32(20)
							t31 := v10
							v1 = v11 + (v1^i32(-1))*i32(20)
							t32 := int32(load32(m.memory[int64(uint32(v1))+16:]))
							store32(m.memory[int64(uint32(t31))+16:], uint32(t32))
							t33 := int64(load64(m.memory[int64(uint32(v1))+8:]))
							store64(m.memory[int64(uint32(v10))+8:], uint64(t33))
							t34 := int64(load64(m.memory[uint32(v1):]))
							store64(m.memory[uint32(v10):], uint64(t34))
							v14 = v14 + i32(-1)
							if v14 != 0 {
								goto l14
							}
						}
					}
				l8:
					store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
					store32(m.memory[uint32(v0):], uint32(v6))
					store32(m.memory[int64(uint32(v0))+8:], uint32(v7-v2))
					if v4 == 0 {
						goto l15
					}
					t35 := v4
					v8 = (v4*i32(20) + i32(27)) & i32(-8)
					v3 = t35 + v8 + i32(9)
					if v3 == 0 {
						goto l15
					}
					v4 = v11 - v8
					t36 := int32(load32(m.memory[uint32(v4+i32(-4)):]))
					v8 = t36
					v6 = v8 & i32(-8)
					t37 := v6
					v8 = v8 & i32(3)
					p38 := i32(8)
					if v8 != 0 {
						p38 = i32(4)
					}
					if uint32(t37) < uint32(p38+v3) {
						m.fn3(i32(1274224), i32(46), i32(1274272))
						panic("unreachable")
					}
					if v8 == 0 {
						goto l17
					}
					if uint32(v6) > uint32(v3+i32(39)) {
						m.fn3(i32(1274288), i32(46), i32(1274336))
						panic("unreachable")
					}
				l17:
					m.fn1(v4)
					return i32(-1)
				}
				m.fn23(i32(8), v6)
				panic("unreachable")
			}
		}
	l1:
		{
			if v5 != 0 {
				goto l19
			}
			v3 = i32(0)
			goto l20
		l19:
			t39 := int32(load32(m.memory[uint32(v0):]))
			v8 = t39
			v3 = i32(0)
			{
				{
					t40 := v6
					var p41 int32
					if v5&i32(7) != i32(0) {
						p41 = 1
					}
					v6 = t40 + p41
					if v6 == i32(1) {
						goto l21
					}
					v11 = v6 & i32(1)
					v10 = v6 & i32(0x3ffffffe)
					v3 = i32(0)
				l22:
					{
						v6 = v8 + v3
						t42 := int64(load64(m.memory[uint32(v6):]))
						t43 := v6
						v9 = t42
						store64(m.memory[uint32(t43):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
						v6 = v6 + i32(8)
						t44 := int64(load64(m.memory[uint32(v6):]))
						t45 := v6
						v9 = t44
						store64(m.memory[uint32(t45):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
						v3 = v3 + i32(16)
						v10 = v10 + i32(-2)
						if v10 != 0 {
							goto l22
						}
					}
					if v11 == 0 {
						goto l23
					}
				}
			l21:
				v3 = v8 + v3
				t46 := int64(load64(m.memory[uint32(v3):]))
				t47 := v3
				v9 = t46
				store64(m.memory[uint32(t47):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
			}
		l23:
			{
				if uint32(v5) < uint32(i32(8)) {
					goto l24
				}
				t48 := int64(load64(m.memory[uint32(v8):]))
				store64(m.memory[uint32(v8+v5):], uint64(t48))
				goto l25
			}
		l24:
			if v5 == 0 {
				goto l25
			}
			memory_copy(m.memory, uint32(v8+i32(8)), uint32(v8), uint32(v5))
		l25:
			t49 := int64(load64(m.memory[int64(uint32(v1))+8:]))
			v16 = t49
			t50 := int64(load64(m.memory[uint32(v1):]))
			v18 = t50
			v6 = i32(0)
		l33:
			{
				t51 := v8
				v3 = v6
				v10 = t51 + v3
				t52 := int32(m.memory[uint32(v10)])
				if t52 != i32(128) {
					goto l26
				}
				v15 = v8 + (v3^i32(-1))*i32(20)
				v6 = v8 + (i32(0)-v3)*i32(20)
				v11 = v6 + i32(-12)
				v14 = v6 + i32(-16)
			l32:
				{
					t53 := int32(load32(m.memory[uint32(v14):]))
					t54 := int32(load32(m.memory[uint32(v11):]))
					t55 := m.fn60(v18, v16, t53, t54)
					t56 := v4
					v1 = int32(t55)
					v6 = t56 & v1
					v5 = v6
					{
						t57 := int64(load64(m.memory[uint32(v8+v6):]))
						v9 = t57 & i64(-0x7f7f7f7f7f7f7f80)
						if v9 != i64(0) {
							goto l27
						}
						v17 = i32(8)
						v5 = v6
					l28:
						{
							v5 = v5 + v17
							v17 = v17 + i32(8)
							t58 := v8
							v5 = v5 & v4
							t59 := int64(load64(m.memory[uint32(t58+v5):]))
							v9 = t59 & i64(-0x7f7f7f7f7f7f7f80)
							if v9 == 0 {
								goto l28
							}
						}
					}
				l27:
					{
						t60 := v8
						v5 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v9))))>>3) + v5) & v4
						t61 := int32(int8(m.memory[uint32(t60+v5)]))
						if t61 < i32(0) {
							goto l29
						}
						t62 := int64(load64(m.memory[uint32(v8):]))
						v5 = int32(uint32(int64(bits.TrailingZeros64(uint64(t62&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
					}
				l29:
					{
						if uint32((v5-v6^(v3-v6))&v4) < uint32(i32(8)) {
							goto l30
						}
						v6 = v8 + v5
						t63 := int32(m.memory[uint32(v6)])
						v17 = t63
						t64 := v6
						v1 = int32(uint32(v1) >> 25)
						m.memory[uint32(t64)] = byte(v1)
						m.memory[uint32(v8+(v5+i32(-8))&v4+i32(8))] = byte(v1)
						v6 = v8 + (v5^i32(-1))*i32(20)
						{
							if v17 != i32(255) {
								t68 := int32(load32(m.memory[uint32(v15):]))
								v5 = t68
								t69 := int32(load32(m.memory[uint32(v6):]))
								store32(m.memory[uint32(v15):], uint32(t69))
								store32(m.memory[uint32(v6):], uint32(v5))
								t70 := int32(load32(m.memory[int64(uint32(v6))+4:]))
								v5 = t70
								t71 := int32(load32(m.memory[int64(uint32(v15))+4:]))
								store32(m.memory[int64(uint32(v6))+4:], uint32(t71))
								store32(m.memory[int64(uint32(v15))+4:], uint32(v5))
								t72 := int32(load32(m.memory[int64(uint32(v15))+8:]))
								v5 = t72
								t73 := int32(load32(m.memory[int64(uint32(v6))+8:]))
								store32(m.memory[int64(uint32(v15))+8:], uint32(t73))
								store32(m.memory[int64(uint32(v6))+8:], uint32(v5))
								t74 := int32(load32(m.memory[int64(uint32(v6))+12:]))
								v5 = t74
								t75 := int32(load32(m.memory[int64(uint32(v15))+12:]))
								store32(m.memory[int64(uint32(v6))+12:], uint32(t75))
								store32(m.memory[int64(uint32(v15))+12:], uint32(v5))
								t76 := int32(load32(m.memory[int64(uint32(v15))+16:]))
								v5 = t76
								t77 := int32(load32(m.memory[int64(uint32(v6))+16:]))
								store32(m.memory[int64(uint32(v15))+16:], uint32(t77))
								store32(m.memory[int64(uint32(v6))+16:], uint32(v5))
								goto l32
							}
							m.memory[uint32(v10)] = byte(i32(255))
							m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(i32(255))
							t65 := int32(load32(m.memory[int64(uint32(v15))+16:]))
							store32(m.memory[int64(uint32(v6))+16:], uint32(t65))
							t66 := int64(load64(m.memory[int64(uint32(v15))+8:]))
							store64(m.memory[int64(uint32(v6))+8:], uint64(t66))
							t67 := int64(load64(m.memory[uint32(v15):]))
							store64(m.memory[uint32(v6):], uint64(t67))
							goto l26
						}
					}
				l30:
				}
				t78 := v10
				v6 = int32(uint32(v1) >> 25)
				m.memory[uint32(t78)] = byte(v6)
				m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(v6)
			}
		l26:
			v6 = v3 + i32(1)
			if v3 != v4 {
				goto l33
			}
			p79 := v7
			if uint32(v4) < uint32(i32(8)) {
				p79 = v4
			}
			v3 = p79
		}
	l20:
		store32(m.memory[int64(uint32(v0))+8:], uint32(v3-v2))
	l15:
		return i32(-1)
	}
l5:
	m.fn27(i32(1271632), i32(57), i32(1271660))
	panic("unreachable")
}
func (m *Module) fn60(v0, v1 int64, v2, v3 int32) int64 {
	var v4 int32
	var v5, v6, v7, v8 int64
	t0 := m.g0
	v4 = t0 - i32(80)
	m.g0 = v4
	store64(m.memory[int64(uint32(v4))+56:], uint64(i64(0)))
	store64(m.memory[int64(uint32(v4))+64:], uint64(i64(0)))
	store64(m.memory[int64(uint32(v4))+48:], uint64(v1))
	store64(m.memory[int64(uint32(v4))+32:], uint64(v1^i64(8387220255154660723)))
	store64(m.memory[int64(uint32(v4))+24:], uint64(v1^i64(7237128888997146477)))
	store64(m.memory[int64(uint32(v4))+40:], uint64(v0))
	store64(m.memory[int64(uint32(v4))+16:], uint64(v0^i64(0x6c7967656e657261)))
	store64(m.memory[int64(uint32(v4))+8:], uint64(v0^i64(8317987319222330741)))
	store32(m.memory[int64(uint32(v4))+76:], uint32(v3))
	m.fn58(v4+i32(8), v4+i32(76), i32(4))
	m.fn58(v4+i32(8), v2, v3)
	t1 := int64(load64(m.memory[int64(uint32(v4))+8:]))
	v0 = t1
	t2 := int64(load64(m.memory[int64(uint32(v4))+24:]))
	v1 = t2
	t3 := int64(load32(m.memory[int64(uint32(v4))+64:]))
	v5 = t3
	t4 := int64(load64(m.memory[int64(uint32(v4))+56:]))
	v6 = t4
	t5 := int64(load64(m.memory[int64(uint32(v4))+32:]))
	v7 = t5
	t6 := int64(load64(m.memory[int64(uint32(v4))+16:]))
	v8 = t6
	m.g0 = v4 + i32(80)
	t7 := v7
	v5 = v6 | v5<<56
	v6 = t7 ^ v5
	t8 := i64_rotl(v6, i64(16))
	v6 = v6 + v8
	v7 = t8 ^ v6
	t9 := i64_rotl(v7, i64(21))
	t10 := v7
	v0 = v1 + v0
	v7 = t10 + i64_rotl(v0, i64(32))
	v8 = t9 ^ v7
	t11 := i64_rotl(v8, i64(16))
	t12 := v8
	t13 := v6
	v1 = i64_rotl(v1, i64(13)) ^ v0
	v0 = t13 + v1
	v6 = t12 + (i64_rotl(v0, i64(32)) ^ i64(255))
	v8 = t11 ^ v6
	t14 := i64_rotl(v8, i64(21))
	t15 := v8
	t16 := v7 ^ v5
	v1 = v0 ^ i64_rotl(v1, i64(17))
	v0 = t16 + v1
	v5 = t15 + i64_rotl(v0, i64(32))
	v7 = t14 ^ v5
	t17 := i64_rotl(v7, i64(16))
	t18 := v7
	v1 = v0 ^ i64_rotl(v1, i64(13))
	v0 = v1 + v6
	v6 = t18 + i64_rotl(v0, i64(32))
	v7 = t17 ^ v6
	t19 := i64_rotl(v7, i64(21))
	t20 := v7
	v1 = v0 ^ i64_rotl(v1, i64(17))
	v0 = v1 + v5
	v5 = t20 + i64_rotl(v0, i64(32))
	v7 = t19 ^ v5
	t21 := i64_rotl(v7, i64(16))
	t22 := v7
	v1 = i64_rotl(v1, i64(13)) ^ v0
	v0 = v1 + v6
	v6 = t22 + i64_rotl(v0, i64(32))
	t23 := i64_rotl(t21^v6, i64(21))
	v1 = i64_rotl(v1, i64(17)) ^ v0
	v1 = i64_rotl(v1, i64(13)) ^ (v1 + v5)
	t24 := t23 ^ i64_rotl(v1, i64(17))
	v1 = v1 + v6
	return t24 ^ i64_rotl(v1, i64(32)) ^ v1
}
func (m *Module) fn61(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7, v8 int32
	var v9 int64
	var v10, v11 int32
	var v12, v13 int64
	var v14, v15 int32
	var v16 int64
	var v17 int32
	var v18 int64
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
		v2 = t0
		v3 = v2 + i32(1)
		if v3 == 0 {
			m.fn27(i32(1271632), i32(57), i32(1271660))
			panic("unreachable")
		}
		{
			t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t2 := v3
			v4 = t1
			t3 := v4
			v5 = v4 + i32(1)
			v6 = int32(uint32(v5) >> 3)
			v7 = v6 * i32(7)
			p4 := v7
			if uint32(v4) < uint32(i32(8)) {
				p4 = t3
			}
			v8 = p4
			if uint32(t2) <= uint32(int32(uint32(v8)>>1)) {
				goto l1
			}
			{
				{
					{
						v8 = v8 + i32(1)
						p5 := v3
						if uint32(v8) > uint32(v3) {
							p5 = v8
						}
						v3 = p5
						if uint32(v3) < uint32(i32(15)) {
							goto l2
						}
						if uint32(v3) > uint32(i32(0x1fffffff)) {
							m.fn27(i32(1271632), i32(57), i32(1271660))
							panic("unreachable")
						}
						t6 := int32(uint32(v3<<3) / uint32(i32(7)))
						v3 = i32_shr_u(i32(-1), int32(bits.LeadingZeros32(uint32(t6+i32(-1))))) + i32(1)
						goto l4
					}
				l2:
					p7 := v3&i32(8) + i32(8)
					if uint32(v3) < uint32(i32(4)) {
						p7 = i32(4)
					}
					v3 = p7
				}
			l4:
				v9 = int64(uint32(v3)) * i64(24)
				if int32(int64(uint64(v9)>>32)) != 0 {
					goto l5
				}
				v8 = v3 + i32(8)
				t8 := v8
				v10 = int32(v9)
				v6 = t8 + v10
				if uint32(v6) < uint32(v8) {
					goto l5
				}
				if uint32(v6) > uint32(i32(0x7ffffff8)) {
					goto l5
				}
				t9 := m.fn7(v6)
				v5 = t9
				if v5 != 0 {
					v6 = v5 + v10
					if v8 == 0 {
						goto l7
					}
					memory_fill(m.memory, uint32(v6), i32(255), uint32(v8))
				l7:
					v5 = v3 + i32(-1)
					p10 := int32(uint32(v3)>>3) * i32(7)
					if uint32(v3) < uint32(i32(9)) {
						p10 = v5
					}
					v7 = p10
					t11 := int32(load32(m.memory[uint32(v0):]))
					v11 = t11
					{
						if v2 == 0 {
							goto l8
						}
						t12 := int64(load64(m.memory[uint32(v11):]))
						v9 = (t12 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
						t13 := int64(load64(m.memory[int64(uint32(v1))+8:]))
						v12 = t13
						t14 := int64(load64(m.memory[uint32(v1):]))
						v13 = t14
						v8 = v11
						v14 = v2
						v3 = i32(0)
					l14:
						{
							if v9 != i64(0) {
								goto l9
							}
						l10:
							{
								v3 = v3 + i32(8)
								v8 = v8 + i32(8)
								t15 := int64(load64(m.memory[uint32(v8):]))
								v9 = t15 & i64(-0x7f7f7f7f7f7f7f80)
								if v9 == i64(-0x7f7f7f7f7f7f7f80) {
									goto l10
								}
							}
							v9 = v9 ^ i64(-0x7f7f7f7f7f7f7f80)
						l9:
							{
								t16 := v6
								t17 := v5
								t18 := v13
								t19 := v12
								t20 := v11
								v1 = int32(uint32(int64(bits.TrailingZeros64(uint64(v9))))>>3) + v3
								v10 = t20 + (i32(0)-v1)*i32(24)
								t21 := int32(load32(m.memory[uint32(v10+i32(-20)):]))
								t22 := int32(load32(m.memory[uint32(v10+i32(-16)):]))
								t23 := m.fn60(t18, t19, t21, t22)
								v15 = int32(t23)
								v10 = t17 & v15
								t24 := int64(load64(m.memory[uint32(t16+v10):]))
								v16 = t24 & i64(-0x7f7f7f7f7f7f7f80)
								if v16 != i64(0) {
									goto l11
								}
								v17 = i32(8)
							l12:
								{
									v10 = v10 + v17
									v17 = v17 + i32(8)
									t25 := v6
									v10 = v10 & v5
									t26 := int64(load64(m.memory[uint32(t25+v10):]))
									v16 = t26 & i64(-0x7f7f7f7f7f7f7f80)
									if v16 == 0 {
										goto l12
									}
								}
							}
						l11:
							v18 = v9 + i64(-1)
							{
								t27 := v6
								v10 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v16))))>>3) + v10) & v5
								t28 := int32(int8(m.memory[uint32(t27+v10)]))
								if t28 < i32(0) {
									goto l13
								}
								t29 := int64(load64(m.memory[uint32(v6):]))
								v10 = int32(uint32(int64(bits.TrailingZeros64(uint64(t29&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
							}
						l13:
							v9 = v18 & v9
							t30 := v6 + v10
							v15 = int32(uint32(v15) >> 25)
							m.memory[uint32(t30)] = byte(v15)
							m.memory[uint32(v6+(v10+i32(-8))&v5+i32(8))] = byte(v15)
							v10 = v6 + (v10^i32(-1))*i32(24)
							t31 := v10
							v1 = v11 + (v1^i32(-1))*i32(24)
							t32 := int64(load64(m.memory[int64(uint32(v1))+16:]))
							store64(m.memory[int64(uint32(t31))+16:], uint64(t32))
							t33 := int64(load64(m.memory[int64(uint32(v1))+8:]))
							store64(m.memory[int64(uint32(v10))+8:], uint64(t33))
							t34 := int64(load64(m.memory[uint32(v1):]))
							store64(m.memory[uint32(v10):], uint64(t34))
							v14 = v14 + i32(-1)
							if v14 != 0 {
								goto l14
							}
						}
					}
				l8:
					store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
					store32(m.memory[uint32(v0):], uint32(v6))
					store32(m.memory[int64(uint32(v0))+8:], uint32(v7-v2))
					if v4 == 0 {
						goto l15
					}
					t35 := v4
					v8 = (v4*i32(24) + i32(31)) & i32(-8)
					v3 = t35 + v8 + i32(9)
					if v3 == 0 {
						goto l15
					}
					v4 = v11 - v8
					t36 := int32(load32(m.memory[uint32(v4+i32(-4)):]))
					v8 = t36
					v6 = v8 & i32(-8)
					t37 := v6
					v8 = v8 & i32(3)
					p38 := i32(8)
					if v8 != 0 {
						p38 = i32(4)
					}
					if uint32(t37) < uint32(p38+v3) {
						m.fn3(i32(1274224), i32(46), i32(1274272))
						panic("unreachable")
					}
					if v8 == 0 {
						goto l17
					}
					if uint32(v6) > uint32(v3+i32(39)) {
						m.fn3(i32(1274288), i32(46), i32(1274336))
						panic("unreachable")
					}
				l17:
					m.fn1(v4)
					return i32(-1)
				}
				m.fn23(i32(8), v6)
				panic("unreachable")
			}
		}
	l1:
		{
			if v5 != 0 {
				goto l19
			}
			v3 = i32(0)
			goto l20
		l19:
			t39 := int32(load32(m.memory[uint32(v0):]))
			v8 = t39
			v3 = i32(0)
			{
				{
					t40 := v6
					var p41 int32
					if v5&i32(7) != i32(0) {
						p41 = 1
					}
					v6 = t40 + p41
					if v6 == i32(1) {
						goto l21
					}
					v11 = v6 & i32(1)
					v10 = v6 & i32(0x3ffffffe)
					v3 = i32(0)
				l22:
					{
						v6 = v8 + v3
						t42 := int64(load64(m.memory[uint32(v6):]))
						t43 := v6
						v9 = t42
						store64(m.memory[uint32(t43):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
						v6 = v6 + i32(8)
						t44 := int64(load64(m.memory[uint32(v6):]))
						t45 := v6
						v9 = t44
						store64(m.memory[uint32(t45):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
						v3 = v3 + i32(16)
						v10 = v10 + i32(-2)
						if v10 != 0 {
							goto l22
						}
					}
					if v11 == 0 {
						goto l23
					}
				}
			l21:
				v3 = v8 + v3
				t46 := int64(load64(m.memory[uint32(v3):]))
				t47 := v3
				v9 = t46
				store64(m.memory[uint32(t47):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
			}
		l23:
			{
				if uint32(v5) < uint32(i32(8)) {
					goto l24
				}
				t48 := int64(load64(m.memory[uint32(v8):]))
				store64(m.memory[uint32(v8+v5):], uint64(t48))
				goto l25
			}
		l24:
			if v5 == 0 {
				goto l25
			}
			memory_copy(m.memory, uint32(v8+i32(8)), uint32(v8), uint32(v5))
		l25:
			t49 := int64(load64(m.memory[int64(uint32(v1))+8:]))
			v16 = t49
			t50 := int64(load64(m.memory[uint32(v1):]))
			v18 = t50
			v6 = i32(0)
		l33:
			{
				t51 := v8
				v3 = v6
				v10 = t51 + v3
				t52 := int32(m.memory[uint32(v10)])
				if t52 != i32(128) {
					goto l26
				}
				v15 = v8 + (v3^i32(-1))*i32(24)
				v6 = v8 + (i32(0)-v3)*i32(24)
				v11 = v6 + i32(-16)
				v14 = v6 + i32(-20)
			l32:
				{
					t53 := int32(load32(m.memory[uint32(v14):]))
					t54 := int32(load32(m.memory[uint32(v11):]))
					t55 := m.fn60(v18, v16, t53, t54)
					t56 := v4
					v1 = int32(t55)
					v6 = t56 & v1
					v5 = v6
					{
						t57 := int64(load64(m.memory[uint32(v8+v6):]))
						v9 = t57 & i64(-0x7f7f7f7f7f7f7f80)
						if v9 != i64(0) {
							goto l27
						}
						v17 = i32(8)
						v5 = v6
					l28:
						{
							v5 = v5 + v17
							v17 = v17 + i32(8)
							t58 := v8
							v5 = v5 & v4
							t59 := int64(load64(m.memory[uint32(t58+v5):]))
							v9 = t59 & i64(-0x7f7f7f7f7f7f7f80)
							if v9 == 0 {
								goto l28
							}
						}
					}
				l27:
					{
						t60 := v8
						v5 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v9))))>>3) + v5) & v4
						t61 := int32(int8(m.memory[uint32(t60+v5)]))
						if t61 < i32(0) {
							goto l29
						}
						t62 := int64(load64(m.memory[uint32(v8):]))
						v5 = int32(uint32(int64(bits.TrailingZeros64(uint64(t62&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
					}
				l29:
					{
						if uint32((v5-v6^(v3-v6))&v4) < uint32(i32(8)) {
							goto l30
						}
						v6 = v8 + v5
						t63 := int32(m.memory[uint32(v6)])
						v17 = t63
						t64 := v6
						v1 = int32(uint32(v1) >> 25)
						m.memory[uint32(t64)] = byte(v1)
						m.memory[uint32(v8+(v5+i32(-8))&v4+i32(8))] = byte(v1)
						v6 = v8 + (v5^i32(-1))*i32(24)
						{
							if v17 != i32(255) {
								t68 := int64(load64(m.memory[uint32(v15):]))
								v9 = t68
								t69 := int64(load64(m.memory[uint32(v6):]))
								store64(m.memory[uint32(v15):], uint64(t69))
								store64(m.memory[uint32(v6):], uint64(v9))
								t70 := int64(load64(m.memory[int64(uint32(v6))+8:]))
								v9 = t70
								t71 := int64(load64(m.memory[int64(uint32(v15))+8:]))
								store64(m.memory[int64(uint32(v6))+8:], uint64(t71))
								store64(m.memory[int64(uint32(v15))+8:], uint64(v9))
								t72 := int32(load32(m.memory[int64(uint32(v15))+16:]))
								v5 = t72
								t73 := int32(load32(m.memory[int64(uint32(v6))+16:]))
								store32(m.memory[int64(uint32(v15))+16:], uint32(t73))
								t74 := int32(load32(m.memory[int64(uint32(v6))+20:]))
								v1 = t74
								t75 := int32(load32(m.memory[int64(uint32(v15))+20:]))
								store32(m.memory[int64(uint32(v6))+20:], uint32(t75))
								store32(m.memory[int64(uint32(v15))+20:], uint32(v1))
								store32(m.memory[int64(uint32(v6))+16:], uint32(v5))
								goto l32
							}
							m.memory[uint32(v10)] = byte(i32(255))
							m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(i32(255))
							t65 := int64(load64(m.memory[int64(uint32(v15))+16:]))
							store64(m.memory[int64(uint32(v6))+16:], uint64(t65))
							t66 := int64(load64(m.memory[int64(uint32(v15))+8:]))
							store64(m.memory[int64(uint32(v6))+8:], uint64(t66))
							t67 := int64(load64(m.memory[uint32(v15):]))
							store64(m.memory[uint32(v6):], uint64(t67))
							goto l26
						}
					}
				l30:
				}
				t76 := v10
				v6 = int32(uint32(v1) >> 25)
				m.memory[uint32(t76)] = byte(v6)
				m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(v6)
			}
		l26:
			v6 = v3 + i32(1)
			if v3 != v4 {
				goto l33
			}
			p77 := v7
			if uint32(v4) < uint32(i32(8)) {
				p77 = v4
			}
			v3 = p77
		}
	l20:
		store32(m.memory[int64(uint32(v0))+8:], uint32(v3-v2))
	l15:
		return i32(-1)
	}
l5:
	m.fn27(i32(1271632), i32(57), i32(1271660))
	panic("unreachable")
}
func (m *Module) fn62(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7, v8 int32
	var v9 int64
	var v10, v11, v12 int32
	var v13, v14 int64
	var v15 int32
	var v16 int64
	var v17 int32
	var v18 int64
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
		v2 = t0
		v3 = v2 + i32(1)
		if v3 == 0 {
			m.fn27(i32(1271632), i32(57), i32(1271660))
			panic("unreachable")
		}
		{
			t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t2 := v3
			v4 = t1
			t3 := v4
			v5 = v4 + i32(1)
			v6 = int32(uint32(v5) >> 3)
			v7 = v6 * i32(7)
			p4 := v7
			if uint32(v4) < uint32(i32(8)) {
				p4 = t3
			}
			v8 = p4
			if uint32(t2) <= uint32(int32(uint32(v8)>>1)) {
				goto l1
			}
			{
				{
					{
						v8 = v8 + i32(1)
						p5 := v3
						if uint32(v8) > uint32(v3) {
							p5 = v8
						}
						v3 = p5
						if uint32(v3) < uint32(i32(15)) {
							goto l2
						}
						if uint32(v3) > uint32(i32(0x1fffffff)) {
							m.fn27(i32(1271632), i32(57), i32(1271660))
							panic("unreachable")
						}
						t6 := int32(uint32(v3<<3) / uint32(i32(7)))
						v3 = i32_shr_u(i32(-1), int32(bits.LeadingZeros32(uint32(t6+i32(-1))))) + i32(1)
						goto l4
					}
				l2:
					p7 := v3&i32(8) + i32(8)
					if uint32(v3) < uint32(i32(4)) {
						p7 = i32(4)
					}
					v3 = p7
				}
			l4:
				v9 = int64(uint32(v3)) * i64(36)
				if int32(int64(uint64(v9)>>32)) != 0 {
					goto l5
				}
				v6 = int32(v9)
				if uint32(v6) > uint32(i32(-8)) {
					goto l5
				}
				v8 = v3 + i32(8)
				t8 := v8
				v10 = (v6 + i32(7)) & i32(-8)
				v6 = t8 + v10
				if uint32(v6) < uint32(v8) {
					goto l5
				}
				if uint32(v6) > uint32(i32(0x7ffffff8)) {
					goto l5
				}
				t9 := m.fn7(v6)
				v5 = t9
				if v5 != 0 {
					v6 = v5 + v10
					if v8 == 0 {
						goto l7
					}
					memory_fill(m.memory, uint32(v6), i32(255), uint32(v8))
				l7:
					v11 = v3 + i32(-1)
					p10 := int32(uint32(v3)>>3) * i32(7)
					if uint32(v3) < uint32(i32(9)) {
						p10 = v11
					}
					v7 = p10
					t11 := int32(load32(m.memory[uint32(v0):]))
					v12 = t11
					{
						if v2 == 0 {
							goto l8
						}
						t12 := int64(load64(m.memory[uint32(v12):]))
						v9 = (t12 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
						t13 := int64(load64(m.memory[int64(uint32(v1))+8:]))
						v13 = t13
						t14 := int64(load64(m.memory[uint32(v1):]))
						v14 = t14
						v8 = v12
						v1 = v2
						v3 = i32(0)
					l14:
						{
							if v9 != i64(0) {
								goto l9
							}
						l10:
							{
								v3 = v3 + i32(8)
								v8 = v8 + i32(8)
								t15 := int64(load64(m.memory[uint32(v8):]))
								v9 = t15 & i64(-0x7f7f7f7f7f7f7f80)
								if v9 == i64(-0x7f7f7f7f7f7f7f80) {
									goto l10
								}
							}
							v9 = v9 ^ i64(-0x7f7f7f7f7f7f7f80)
						l9:
							{
								t16 := v6
								t17 := v11
								t18 := v14
								t19 := v13
								t20 := v12
								v5 = int32(uint32(int64(bits.TrailingZeros64(uint64(v9))))>>3) + v3
								v10 = t20 + (i32(0)-v5)*i32(36)
								t21 := int32(load32(m.memory[uint32(v10+i32(-32)):]))
								t22 := int32(load32(m.memory[uint32(v10+i32(-28)):]))
								t23 := m.fn60(t18, t19, t21, t22)
								v15 = int32(t23)
								v10 = t17 & v15
								t24 := int64(load64(m.memory[uint32(t16+v10):]))
								v16 = t24 & i64(-0x7f7f7f7f7f7f7f80)
								if v16 != i64(0) {
									goto l11
								}
								v17 = i32(8)
							l12:
								{
									v10 = v10 + v17
									v17 = v17 + i32(8)
									t25 := v6
									v10 = v10 & v11
									t26 := int64(load64(m.memory[uint32(t25+v10):]))
									v16 = t26 & i64(-0x7f7f7f7f7f7f7f80)
									if v16 == 0 {
										goto l12
									}
								}
							}
						l11:
							v18 = v9 + i64(-1)
							{
								t27 := v6
								v10 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v16))))>>3) + v10) & v11
								t28 := int32(int8(m.memory[uint32(t27+v10)]))
								if t28 < i32(0) {
									goto l13
								}
								t29 := int64(load64(m.memory[uint32(v6):]))
								v10 = int32(uint32(int64(bits.TrailingZeros64(uint64(t29&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
							}
						l13:
							v9 = v18 & v9
							t30 := v6 + v10
							v15 = int32(uint32(v15) >> 25)
							m.memory[uint32(t30)] = byte(v15)
							m.memory[uint32(v6+(v10+i32(-8))&v11+i32(8))] = byte(v15)
							v10 = v6 + (v10^i32(-1))*i32(36)
							t31 := v10
							v5 = v12 + (v5^i32(-1))*i32(36)
							t32 := int32(load32(m.memory[int64(uint32(v5))+32:]))
							store32(m.memory[int64(uint32(t31))+32:], uint32(t32))
							t33 := int64(load64(m.memory[int64(uint32(v5))+24:]))
							store64(m.memory[int64(uint32(v10))+24:], uint64(t33))
							t34 := int64(load64(m.memory[int64(uint32(v5))+16:]))
							store64(m.memory[int64(uint32(v10))+16:], uint64(t34))
							t35 := int64(load64(m.memory[int64(uint32(v5))+8:]))
							store64(m.memory[int64(uint32(v10))+8:], uint64(t35))
							t36 := int64(load64(m.memory[uint32(v5):]))
							store64(m.memory[uint32(v10):], uint64(t36))
							v1 = v1 + i32(-1)
							if v1 != 0 {
								goto l14
							}
						}
					}
				l8:
					store32(m.memory[int64(uint32(v0))+4:], uint32(v11))
					store32(m.memory[uint32(v0):], uint32(v6))
					store32(m.memory[int64(uint32(v0))+8:], uint32(v7-v2))
					if v4 == 0 {
						goto l15
					}
					t37 := v4
					v8 = (v4*i32(36) + i32(43)) & i32(-8)
					v3 = t37 + v8 + i32(9)
					if v3 == 0 {
						goto l15
					}
					v4 = v12 - v8
					t38 := int32(load32(m.memory[uint32(v4+i32(-4)):]))
					v8 = t38
					v6 = v8 & i32(-8)
					t39 := v6
					v8 = v8 & i32(3)
					p40 := i32(8)
					if v8 != 0 {
						p40 = i32(4)
					}
					if uint32(t39) < uint32(p40+v3) {
						m.fn3(i32(1274224), i32(46), i32(1274272))
						panic("unreachable")
					}
					if v8 == 0 {
						goto l17
					}
					if uint32(v6) > uint32(v3+i32(39)) {
						m.fn3(i32(1274288), i32(46), i32(1274336))
						panic("unreachable")
					}
				l17:
					m.fn1(v4)
					return i32(-1)
				}
				m.fn23(i32(8), v6)
				panic("unreachable")
			}
		}
	l1:
		{
			if v5 != 0 {
				goto l19
			}
			v3 = i32(0)
			goto l20
		l19:
			t41 := int32(load32(m.memory[uint32(v0):]))
			v8 = t41
			v3 = i32(0)
			{
				{
					t42 := v6
					var p43 int32
					if v5&i32(7) != i32(0) {
						p43 = 1
					}
					v6 = t42 + p43
					if v6 == i32(1) {
						goto l21
					}
					v11 = v6 & i32(1)
					v10 = v6 & i32(0x3ffffffe)
					v3 = i32(0)
				l22:
					{
						v6 = v8 + v3
						t44 := int64(load64(m.memory[uint32(v6):]))
						t45 := v6
						v9 = t44
						store64(m.memory[uint32(t45):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
						v6 = v6 + i32(8)
						t46 := int64(load64(m.memory[uint32(v6):]))
						t47 := v6
						v9 = t46
						store64(m.memory[uint32(t47):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
						v3 = v3 + i32(16)
						v10 = v10 + i32(-2)
						if v10 != 0 {
							goto l22
						}
					}
					if v11 == 0 {
						goto l23
					}
				}
			l21:
				v3 = v8 + v3
				t48 := int64(load64(m.memory[uint32(v3):]))
				t49 := v3
				v9 = t48
				store64(m.memory[uint32(t49):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
			}
		l23:
			{
				if uint32(v5) < uint32(i32(8)) {
					goto l24
				}
				t50 := int64(load64(m.memory[uint32(v8):]))
				store64(m.memory[uint32(v8+v5):], uint64(t50))
				goto l25
			}
		l24:
			if v5 == 0 {
				goto l25
			}
			memory_copy(m.memory, uint32(v8+i32(8)), uint32(v8), uint32(v5))
		l25:
			t51 := int64(load64(m.memory[int64(uint32(v1))+8:]))
			v16 = t51
			t52 := int64(load64(m.memory[uint32(v1):]))
			v18 = t52
			v6 = i32(0)
		l33:
			{
				t53 := v8
				v3 = v6
				v10 = t53 + v3
				t54 := int32(m.memory[uint32(v10)])
				if t54 != i32(128) {
					goto l26
				}
				v15 = v8 + (v3^i32(-1))*i32(36)
				v6 = v8 + (i32(0)-v3)*i32(36)
				v12 = v6 + i32(-28)
				v1 = v6 + i32(-32)
			l32:
				{
					t55 := int32(load32(m.memory[uint32(v1):]))
					t56 := int32(load32(m.memory[uint32(v12):]))
					t57 := m.fn60(v18, v16, t55, t56)
					t58 := v4
					v11 = int32(t57)
					v6 = t58 & v11
					v5 = v6
					{
						t59 := int64(load64(m.memory[uint32(v8+v6):]))
						v9 = t59 & i64(-0x7f7f7f7f7f7f7f80)
						if v9 != i64(0) {
							goto l27
						}
						v17 = i32(8)
						v5 = v6
					l28:
						{
							v5 = v5 + v17
							v17 = v17 + i32(8)
							t60 := v8
							v5 = v5 & v4
							t61 := int64(load64(m.memory[uint32(t60+v5):]))
							v9 = t61 & i64(-0x7f7f7f7f7f7f7f80)
							if v9 == 0 {
								goto l28
							}
						}
					}
				l27:
					{
						t62 := v8
						v5 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v9))))>>3) + v5) & v4
						t63 := int32(int8(m.memory[uint32(t62+v5)]))
						if t63 < i32(0) {
							goto l29
						}
						t64 := int64(load64(m.memory[uint32(v8):]))
						v5 = int32(uint32(int64(bits.TrailingZeros64(uint64(t64&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
					}
				l29:
					{
						if uint32((v5-v6^(v3-v6))&v4) < uint32(i32(8)) {
							goto l30
						}
						v6 = v8 + v5
						t65 := int32(m.memory[uint32(v6)])
						v17 = t65
						t66 := v6
						v11 = int32(uint32(v11) >> 25)
						m.memory[uint32(t66)] = byte(v11)
						m.memory[uint32(v8+(v5+i32(-8))&v4+i32(8))] = byte(v11)
						v6 = v8 + (v5^i32(-1))*i32(36)
						{
							if v17 != i32(255) {
								t72 := int32(load32(m.memory[uint32(v15):]))
								v5 = t72
								t73 := int32(load32(m.memory[uint32(v6):]))
								store32(m.memory[uint32(v15):], uint32(t73))
								store32(m.memory[uint32(v6):], uint32(v5))
								t74 := int32(load32(m.memory[int64(uint32(v6))+4:]))
								v5 = t74
								t75 := int32(load32(m.memory[int64(uint32(v15))+4:]))
								store32(m.memory[int64(uint32(v6))+4:], uint32(t75))
								store32(m.memory[int64(uint32(v15))+4:], uint32(v5))
								t76 := int32(load32(m.memory[int64(uint32(v15))+8:]))
								v5 = t76
								t77 := int32(load32(m.memory[int64(uint32(v6))+8:]))
								store32(m.memory[int64(uint32(v15))+8:], uint32(t77))
								store32(m.memory[int64(uint32(v6))+8:], uint32(v5))
								t78 := int32(load32(m.memory[int64(uint32(v6))+12:]))
								v5 = t78
								t79 := int32(load32(m.memory[int64(uint32(v15))+12:]))
								store32(m.memory[int64(uint32(v6))+12:], uint32(t79))
								store32(m.memory[int64(uint32(v15))+12:], uint32(v5))
								t80 := int32(load32(m.memory[int64(uint32(v15))+16:]))
								v5 = t80
								t81 := int32(load32(m.memory[int64(uint32(v6))+16:]))
								store32(m.memory[int64(uint32(v15))+16:], uint32(t81))
								store32(m.memory[int64(uint32(v6))+16:], uint32(v5))
								t82 := int32(load32(m.memory[int64(uint32(v6))+20:]))
								v5 = t82
								t83 := int32(load32(m.memory[int64(uint32(v15))+20:]))
								store32(m.memory[int64(uint32(v6))+20:], uint32(t83))
								store32(m.memory[int64(uint32(v15))+20:], uint32(v5))
								t84 := int32(load32(m.memory[int64(uint32(v15))+24:]))
								v5 = t84
								t85 := int32(load32(m.memory[int64(uint32(v6))+24:]))
								store32(m.memory[int64(uint32(v15))+24:], uint32(t85))
								store32(m.memory[int64(uint32(v6))+24:], uint32(v5))
								t86 := int32(load32(m.memory[int64(uint32(v6))+28:]))
								v5 = t86
								t87 := int32(load32(m.memory[int64(uint32(v15))+28:]))
								store32(m.memory[int64(uint32(v6))+28:], uint32(t87))
								store32(m.memory[int64(uint32(v15))+28:], uint32(v5))
								t88 := int32(load32(m.memory[int64(uint32(v15))+32:]))
								v5 = t88
								t89 := int32(load32(m.memory[int64(uint32(v6))+32:]))
								store32(m.memory[int64(uint32(v15))+32:], uint32(t89))
								store32(m.memory[int64(uint32(v6))+32:], uint32(v5))
								goto l32
							}
							m.memory[uint32(v10)] = byte(i32(255))
							m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(i32(255))
							t67 := int32(load32(m.memory[int64(uint32(v15))+32:]))
							store32(m.memory[int64(uint32(v6))+32:], uint32(t67))
							t68 := int64(load64(m.memory[int64(uint32(v15))+24:]))
							store64(m.memory[int64(uint32(v6))+24:], uint64(t68))
							t69 := int64(load64(m.memory[int64(uint32(v15))+16:]))
							store64(m.memory[int64(uint32(v6))+16:], uint64(t69))
							t70 := int64(load64(m.memory[int64(uint32(v15))+8:]))
							store64(m.memory[int64(uint32(v6))+8:], uint64(t70))
							t71 := int64(load64(m.memory[uint32(v15):]))
							store64(m.memory[uint32(v6):], uint64(t71))
							goto l26
						}
					}
				l30:
				}
				t90 := v10
				v6 = int32(uint32(v11) >> 25)
				m.memory[uint32(t90)] = byte(v6)
				m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(v6)
			}
		l26:
			v6 = v3 + i32(1)
			if v3 != v4 {
				goto l33
			}
			p91 := v7
			if uint32(v4) < uint32(i32(8)) {
				p91 = v4
			}
			v3 = p91
		}
	l20:
		store32(m.memory[int64(uint32(v0))+8:], uint32(v3-v2))
	l15:
		return i32(-1)
	}
l5:
	m.fn27(i32(1271632), i32(57), i32(1271660))
	panic("unreachable")
}
func (m *Module) fn63(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7, v8 int32
	var v9 int64
	var v10, v11, v12 int32
	var v13, v14 int64
	var v15, v16 int32
	var v17, v18 int64
	var v19, v20, v21, v22 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
		v2 = t0
		v3 = v2 + i32(1)
		if v3 == 0 {
			m.fn27(i32(1271632), i32(57), i32(1271660))
			panic("unreachable")
		}
		{
			t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t2 := v3
			v4 = t1
			t3 := v4
			v5 = v4 + i32(1)
			v6 = int32(uint32(v5) >> 3)
			v7 = v6 * i32(7)
			p4 := v7
			if uint32(v4) < uint32(i32(8)) {
				p4 = t3
			}
			v8 = p4
			if uint32(t2) <= uint32(int32(uint32(v8)>>1)) {
				goto l1
			}
			{
				{
					{
						v8 = v8 + i32(1)
						p5 := v3
						if uint32(v8) > uint32(v3) {
							p5 = v8
						}
						v3 = p5
						if uint32(v3) < uint32(i32(15)) {
							goto l2
						}
						if uint32(v3) > uint32(i32(0x1fffffff)) {
							m.fn27(i32(1271632), i32(57), i32(1271660))
							panic("unreachable")
						}
						t6 := int32(uint32(v3<<3) / uint32(i32(7)))
						v3 = i32_shr_u(i32(-1), int32(bits.LeadingZeros32(uint32(t6+i32(-1))))) + i32(1)
						goto l4
					}
				l2:
					p7 := v3&i32(8) + i32(8)
					if uint32(v3) < uint32(i32(4)) {
						p7 = i32(4)
					}
					v3 = p7
				}
			l4:
				v9 = int64(uint32(v3)) * i64(416)
				if int32(int64(uint64(v9)>>32)) != 0 {
					goto l5
				}
				v8 = v3 + i32(8)
				t8 := v8
				v10 = int32(v9)
				v6 = t8 + v10
				if uint32(v6) < uint32(v8) {
					goto l5
				}
				if uint32(v6) > uint32(i32(0x7ffffff8)) {
					goto l5
				}
				t9 := m.fn7(v6)
				v5 = t9
				if v5 != 0 {
					v6 = v5 + v10
					if v8 == 0 {
						goto l7
					}
					memory_fill(m.memory, uint32(v6), i32(255), uint32(v8))
				l7:
					v5 = v3 + i32(-1)
					p10 := int32(uint32(v3)>>3) * i32(7)
					if uint32(v3) < uint32(i32(9)) {
						p10 = v5
					}
					v11 = p10
					t11 := int32(load32(m.memory[uint32(v0):]))
					v12 = t11
					{
						if v2 == 0 {
							goto l8
						}
						t12 := int64(load64(m.memory[uint32(v12):]))
						v9 = (t12 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
						t13 := int64(load64(m.memory[int64(uint32(v1))+8:]))
						v13 = t13
						t14 := int64(load64(m.memory[uint32(v1):]))
						v14 = t14
						v8 = v12
						v1 = v2
						v3 = i32(0)
					l14:
						{
							if v9 != i64(0) {
								goto l9
							}
						l10:
							{
								v3 = v3 + i32(8)
								v8 = v8 + i32(8)
								t15 := int64(load64(m.memory[uint32(v8):]))
								v9 = t15 & i64(-0x7f7f7f7f7f7f7f80)
								if v9 == i64(-0x7f7f7f7f7f7f7f80) {
									goto l10
								}
							}
							v9 = v9 ^ i64(-0x7f7f7f7f7f7f7f80)
						l9:
							{
								t16 := v6
								t17 := v5
								t18 := v14
								t19 := v13
								t20 := v12
								v15 = int32(uint32(int64(bits.TrailingZeros64(uint64(v9))))>>3) + v3
								v10 = t20 + (i32(0)-v15)*i32(416)
								t21 := int32(load32(m.memory[uint32(v10+i32(-412)):]))
								t22 := int32(load32(m.memory[uint32(v10+i32(-408)):]))
								t23 := m.fn64(t18, t19, t21, t22)
								v16 = int32(t23)
								v10 = t17 & v16
								t24 := int64(load64(m.memory[uint32(t16+v10):]))
								v17 = t24 & i64(-0x7f7f7f7f7f7f7f80)
								if v17 != i64(0) {
									goto l11
								}
								v7 = i32(8)
							l12:
								{
									v10 = v10 + v7
									v7 = v7 + i32(8)
									t25 := v6
									v10 = v10 & v5
									t26 := int64(load64(m.memory[uint32(t25+v10):]))
									v17 = t26 & i64(-0x7f7f7f7f7f7f7f80)
									if v17 == 0 {
										goto l12
									}
								}
							}
						l11:
							v18 = v9 + i64(-1)
							{
								t27 := v6
								v10 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v17))))>>3) + v10) & v5
								t28 := int32(int8(m.memory[uint32(t27+v10)]))
								if t28 < i32(0) {
									goto l13
								}
								t29 := int64(load64(m.memory[uint32(v6):]))
								v10 = int32(uint32(int64(bits.TrailingZeros64(uint64(t29&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
							}
						l13:
							v9 = v18 & v9
							t30 := v6 + v10
							v16 = int32(uint32(v16) >> 25)
							m.memory[uint32(t30)] = byte(v16)
							m.memory[uint32(v6+(v10+i32(-8))&v5+i32(8))] = byte(v16)
							memory_copy(m.memory, uint32(v6+(v10^i32(-1))*i32(416)), uint32(v12+(v15^i32(-1))*i32(416)), uint32(i32(416)))
							v1 = v1 + i32(-1)
							if v1 != 0 {
								goto l14
							}
						}
					}
				l8:
					store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
					store32(m.memory[uint32(v0):], uint32(v6))
					store32(m.memory[int64(uint32(v0))+8:], uint32(v11-v2))
					if v4 == 0 {
						goto l15
					}
					t31 := v4
					v8 = (v4*i32(416) + i32(423)) & i32(-32)
					v3 = t31 + v8 + i32(9)
					if v3 == 0 {
						goto l15
					}
					v4 = v12 - v8
					t32 := int32(load32(m.memory[uint32(v4+i32(-4)):]))
					v8 = t32
					v6 = v8 & i32(-8)
					t33 := v6
					v8 = v8 & i32(3)
					p34 := i32(8)
					if v8 != 0 {
						p34 = i32(4)
					}
					if uint32(t33) < uint32(p34+v3) {
						m.fn3(i32(1274224), i32(46), i32(1274272))
						panic("unreachable")
					}
					if v8 == 0 {
						goto l17
					}
					if uint32(v6) > uint32(v3+i32(39)) {
						m.fn3(i32(1274288), i32(46), i32(1274336))
						panic("unreachable")
					}
				l17:
					m.fn1(v4)
					return i32(-1)
				}
				m.fn23(i32(8), v6)
				panic("unreachable")
			}
		}
	l1:
		{
			if v5 != 0 {
				goto l19
			}
			v3 = i32(0)
			goto l20
		l19:
			t35 := int32(load32(m.memory[uint32(v0):]))
			v8 = t35
			v3 = i32(0)
			{
				{
					t36 := v6
					var p37 int32
					if v5&i32(7) != i32(0) {
						p37 = 1
					}
					v6 = t36 + p37
					if v6 == i32(1) {
						goto l21
					}
					v12 = v6 & i32(1)
					v10 = v6 & i32(0x3ffffffe)
					v3 = i32(0)
				l22:
					{
						v6 = v8 + v3
						t38 := int64(load64(m.memory[uint32(v6):]))
						t39 := v6
						v9 = t38
						store64(m.memory[uint32(t39):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
						v6 = v6 + i32(8)
						t40 := int64(load64(m.memory[uint32(v6):]))
						t41 := v6
						v9 = t40
						store64(m.memory[uint32(t41):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
						v3 = v3 + i32(16)
						v10 = v10 + i32(-2)
						if v10 != 0 {
							goto l22
						}
					}
					if v12 == 0 {
						goto l23
					}
				}
			l21:
				v3 = v8 + v3
				t42 := int64(load64(m.memory[uint32(v3):]))
				t43 := v3
				v9 = t42
				store64(m.memory[uint32(t43):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
			}
		l23:
			{
				if uint32(v5) < uint32(i32(8)) {
					goto l24
				}
				t44 := int64(load64(m.memory[uint32(v8):]))
				store64(m.memory[uint32(v8+v5):], uint64(t44))
				goto l25
			}
		l24:
			if v5 == 0 {
				goto l25
			}
			memory_copy(m.memory, uint32(v8+i32(8)), uint32(v8), uint32(v5))
		l25:
			t45 := int64(load64(m.memory[int64(uint32(v1))+8:]))
			v17 = t45
			t46 := int64(load64(m.memory[uint32(v1):]))
			v18 = t46
			v6 = v8
			v10 = i32(0)
		l34:
			{
				t47 := v8
				v3 = v10
				v5 = t47 + v3
				t48 := int32(m.memory[uint32(v5)])
				if t48 != i32(128) {
					goto l26
				}
				v11 = v8 + (v3^i32(-1))*i32(416)
				v10 = v8 + (i32(0)-v3)*i32(416)
				v15 = v10 + i32(-408)
				v16 = v10 + i32(-412)
				{
				l33:
					{
						t49 := int32(load32(m.memory[uint32(v16):]))
						t50 := int32(load32(m.memory[uint32(v15):]))
						t51 := m.fn64(v18, v17, t49, t50)
						t52 := v4
						v1 = int32(t51)
						v10 = t52 & v1
						v12 = v10
						{
							t53 := int64(load64(m.memory[uint32(v8+v10):]))
							v9 = t53 & i64(-0x7f7f7f7f7f7f7f80)
							if v9 != i64(0) {
								goto l27
							}
							v19 = i32(8)
							v12 = v10
						l28:
							{
								v12 = v12 + v19
								v19 = v19 + i32(8)
								t54 := v8
								v12 = v12 & v4
								t55 := int64(load64(m.memory[uint32(t54+v12):]))
								v9 = t55 & i64(-0x7f7f7f7f7f7f7f80)
								if v9 == 0 {
									goto l28
								}
							}
						}
					l27:
						{
							t56 := v8
							v12 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v9))))>>3) + v12) & v4
							t57 := int32(int8(m.memory[uint32(t56+v12)]))
							if t57 < i32(0) {
								goto l29
							}
							t58 := int64(load64(m.memory[uint32(v8):]))
							v12 = int32(uint32(int64(bits.TrailingZeros64(uint64(t58&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
						}
					l29:
						{
							if uint32((v12-v10^(v3-v10))&v4) < uint32(i32(8)) {
								goto l30
							}
							v10 = v8 + v12
							t59 := int32(m.memory[uint32(v10)])
							v19 = t59
							t60 := v10
							v1 = int32(uint32(v1) >> 25)
							m.memory[uint32(t60)] = byte(v1)
							m.memory[uint32(v8+(v12+i32(-8))&v4+i32(8))] = byte(v1)
							if v19 == i32(255) {
								goto l31
							}
							v1 = i32(-416)
							v20 = v8 + v12*i32(-416)
						l32:
							{
								v10 = v6 + v1
								t61 := int32(load32(m.memory[uint32(v10):]))
								v19 = t61
								t62 := v10
								v12 = v20 + v1
								t63 := int32(load32(m.memory[uint32(v12):]))
								store32(m.memory[uint32(t62):], uint32(t63))
								store32(m.memory[uint32(v12):], uint32(v19))
								v19 = v12 + i32(4)
								t64 := int32(load32(m.memory[uint32(v19):]))
								v21 = t64
								t65 := v19
								v22 = v10 + i32(4)
								t66 := int32(load32(m.memory[uint32(v22):]))
								store32(m.memory[uint32(t65):], uint32(t66))
								store32(m.memory[uint32(v22):], uint32(v21))
								v19 = v10 + i32(8)
								t67 := int32(load32(m.memory[uint32(v19):]))
								v21 = t67
								t68 := v19
								v22 = v12 + i32(8)
								t69 := int32(load32(m.memory[uint32(v22):]))
								store32(m.memory[uint32(t68):], uint32(t69))
								store32(m.memory[uint32(v22):], uint32(v21))
								v12 = v12 + i32(12)
								t70 := int32(load32(m.memory[uint32(v12):]))
								v19 = t70
								t71 := v12
								v10 = v10 + i32(12)
								t72 := int32(load32(m.memory[uint32(v10):]))
								store32(m.memory[uint32(t71):], uint32(t72))
								store32(m.memory[uint32(v10):], uint32(v19))
								v1 = v1 + i32(16)
								if v1 != 0 {
									goto l32
								}
								goto l33
							}
						}
					l30:
					}
					t73 := v5
					v10 = int32(uint32(v1) >> 25)
					m.memory[uint32(t73)] = byte(v10)
					m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(v10)
					goto l26
				}
			l31:
				m.memory[uint32(v5)] = byte(i32(255))
				m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(i32(255))
				memory_copy(m.memory, uint32(v8+(v12^i32(-1))*i32(416)), uint32(v11), uint32(i32(416)))
			}
		l26:
			v10 = v3 + i32(1)
			v6 = v6 + i32(-416)
			if v3 != v4 {
				goto l34
			}
			p74 := v7
			if uint32(v4) < uint32(i32(8)) {
				p74 = v4
			}
			v3 = p74
		}
	l20:
		store32(m.memory[int64(uint32(v0))+8:], uint32(v3-v2))
	l15:
		return i32(-1)
	l5:
		m.fn27(i32(1271632), i32(57), i32(1271660))
		panic("unreachable")
	}
}
func (m *Module) fn64(v0, v1 int64, v2, v3 int32) int64 {
	var v4 int32
	var v5, v6, v7, v8 int64
	t0 := m.g0
	v4 = t0 - i32(80)
	m.g0 = v4
	store64(m.memory[int64(uint32(v4))+56:], uint64(i64(0)))
	store64(m.memory[int64(uint32(v4))+64:], uint64(i64(0)))
	store64(m.memory[int64(uint32(v4))+48:], uint64(v1))
	store64(m.memory[int64(uint32(v4))+32:], uint64(v1^i64(8387220255154660723)))
	store64(m.memory[int64(uint32(v4))+24:], uint64(v1^i64(7237128888997146477)))
	store64(m.memory[int64(uint32(v4))+40:], uint64(v0))
	store64(m.memory[int64(uint32(v4))+16:], uint64(v0^i64(0x6c7967656e657261)))
	store64(m.memory[int64(uint32(v4))+8:], uint64(v0^i64(8317987319222330741)))
	m.fn58(v4+i32(8), v2, v3)
	m.memory[int64(uint32(v4))+79] = byte(i32(255))
	m.fn58(v4+i32(8), v4+i32(79), i32(1))
	t1 := int64(load64(m.memory[int64(uint32(v4))+8:]))
	v0 = t1
	t2 := int64(load64(m.memory[int64(uint32(v4))+24:]))
	v1 = t2
	t3 := int64(load32(m.memory[int64(uint32(v4))+64:]))
	v5 = t3
	t4 := int64(load64(m.memory[int64(uint32(v4))+56:]))
	v6 = t4
	t5 := int64(load64(m.memory[int64(uint32(v4))+32:]))
	v7 = t5
	t6 := int64(load64(m.memory[int64(uint32(v4))+16:]))
	v8 = t6
	m.g0 = v4 + i32(80)
	t7 := v7
	v5 = v6 | v5<<56
	v6 = t7 ^ v5
	t8 := i64_rotl(v6, i64(16))
	v6 = v6 + v8
	v7 = t8 ^ v6
	t9 := i64_rotl(v7, i64(21))
	t10 := v7
	v0 = v1 + v0
	v7 = t10 + i64_rotl(v0, i64(32))
	v8 = t9 ^ v7
	t11 := i64_rotl(v8, i64(16))
	t12 := v8
	t13 := v6
	v1 = i64_rotl(v1, i64(13)) ^ v0
	v0 = t13 + v1
	v6 = t12 + (i64_rotl(v0, i64(32)) ^ i64(255))
	v8 = t11 ^ v6
	t14 := i64_rotl(v8, i64(21))
	t15 := v8
	t16 := v7 ^ v5
	v1 = v0 ^ i64_rotl(v1, i64(17))
	v0 = t16 + v1
	v5 = t15 + i64_rotl(v0, i64(32))
	v7 = t14 ^ v5
	t17 := i64_rotl(v7, i64(16))
	t18 := v7
	v1 = v0 ^ i64_rotl(v1, i64(13))
	v0 = v1 + v6
	v6 = t18 + i64_rotl(v0, i64(32))
	v7 = t17 ^ v6
	t19 := i64_rotl(v7, i64(21))
	t20 := v7
	v1 = v0 ^ i64_rotl(v1, i64(17))
	v0 = v1 + v5
	v5 = t20 + i64_rotl(v0, i64(32))
	v7 = t19 ^ v5
	t21 := i64_rotl(v7, i64(16))
	t22 := v7
	v1 = i64_rotl(v1, i64(13)) ^ v0
	v0 = v1 + v6
	v6 = t22 + i64_rotl(v0, i64(32))
	t23 := i64_rotl(t21^v6, i64(21))
	v1 = i64_rotl(v1, i64(17)) ^ v0
	v1 = i64_rotl(v1, i64(13)) ^ (v1 + v5)
	t24 := t23 ^ i64_rotl(v1, i64(17))
	v1 = v1 + v6
	return t24 ^ i64_rotl(v1, i64(32)) ^ v1
}
func (m *Module) fn65(v0, v1, v2 int32) int32 {
	var v3, v4, v5, v6, v7, v8 int32
	var v9 int64
	var v10, v11 int32
	var v12, v13 int64
	var v14, v15 int32
	var v16 int64
	var v17 int32
	var v18 int64
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
		v3 = t0
		v1 = v3 + v1
		if uint32(v1) < uint32(v3) {
			m.fn27(i32(1271632), i32(57), i32(1271660))
			panic("unreachable")
		}
		{
			t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t2 := v1
			v4 = t1
			t3 := v4
			v5 = v4 + i32(1)
			v6 = int32(uint32(v5) >> 3)
			v7 = v6 * i32(7)
			p4 := v7
			if uint32(v4) < uint32(i32(8)) {
				p4 = t3
			}
			v8 = p4
			if uint32(t2) <= uint32(int32(uint32(v8)>>1)) {
				goto l1
			}
			{
				{
					{
						v8 = v8 + i32(1)
						p5 := v1
						if uint32(v8) > uint32(v1) {
							p5 = v8
						}
						v1 = p5
						if uint32(v1) < uint32(i32(15)) {
							goto l2
						}
						if uint32(v1) > uint32(i32(0x1fffffff)) {
							m.fn27(i32(1271632), i32(57), i32(1271660))
							panic("unreachable")
						}
						t6 := int32(uint32(v1<<3) / uint32(i32(7)))
						v1 = i32_shr_u(i32(-1), int32(bits.LeadingZeros32(uint32(t6+i32(-1))))) + i32(1)
						goto l4
					}
				l2:
					p7 := v1&i32(8) + i32(8)
					if uint32(v1) < uint32(i32(4)) {
						p7 = i32(4)
					}
					v1 = p7
				}
			l4:
				v9 = int64(uint32(v1)) * i64(24)
				if int32(int64(uint64(v9)>>32)) != 0 {
					goto l5
				}
				v8 = v1 + i32(8)
				t8 := v8
				v10 = int32(v9)
				v6 = t8 + v10
				if uint32(v6) < uint32(v8) {
					goto l5
				}
				if uint32(v6) > uint32(i32(0x7ffffff8)) {
					goto l5
				}
				t9 := m.fn7(v6)
				v5 = t9
				if v5 != 0 {
					v6 = v5 + v10
					if v8 == 0 {
						goto l7
					}
					memory_fill(m.memory, uint32(v6), i32(255), uint32(v8))
				l7:
					v5 = v1 + i32(-1)
					p10 := int32(uint32(v1)>>3) * i32(7)
					if uint32(v1) < uint32(i32(9)) {
						p10 = v5
					}
					v7 = p10
					t11 := int32(load32(m.memory[uint32(v0):]))
					v11 = t11
					{
						if v3 == 0 {
							goto l8
						}
						t12 := int64(load64(m.memory[uint32(v11):]))
						v9 = (t12 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
						t13 := int64(load64(m.memory[int64(uint32(v2))+8:]))
						v12 = t13
						t14 := int64(load64(m.memory[uint32(v2):]))
						v13 = t14
						v8 = v11
						v14 = v3
						v1 = i32(0)
					l14:
						{
							if v9 != i64(0) {
								goto l9
							}
						l10:
							{
								v1 = v1 + i32(8)
								v8 = v8 + i32(8)
								t15 := int64(load64(m.memory[uint32(v8):]))
								v9 = t15 & i64(-0x7f7f7f7f7f7f7f80)
								if v9 == i64(-0x7f7f7f7f7f7f7f80) {
									goto l10
								}
							}
							v9 = v9 ^ i64(-0x7f7f7f7f7f7f7f80)
						l9:
							{
								t16 := v6
								t17 := v5
								t18 := v13
								t19 := v12
								t20 := v11
								v2 = int32(uint32(int64(bits.TrailingZeros64(uint64(v9))))>>3) + v1
								v10 = t20 + (i32(0)-v2)*i32(24)
								t21 := int32(load32(m.memory[uint32(v10+i32(-20)):]))
								t22 := int32(load32(m.memory[uint32(v10+i32(-16)):]))
								t23 := m.fn64(t18, t19, t21, t22)
								v15 = int32(t23)
								v10 = t17 & v15
								t24 := int64(load64(m.memory[uint32(t16+v10):]))
								v16 = t24 & i64(-0x7f7f7f7f7f7f7f80)
								if v16 != i64(0) {
									goto l11
								}
								v17 = i32(8)
							l12:
								{
									v10 = v10 + v17
									v17 = v17 + i32(8)
									t25 := v6
									v10 = v10 & v5
									t26 := int64(load64(m.memory[uint32(t25+v10):]))
									v16 = t26 & i64(-0x7f7f7f7f7f7f7f80)
									if v16 == 0 {
										goto l12
									}
								}
							}
						l11:
							v18 = v9 + i64(-1)
							{
								t27 := v6
								v10 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v16))))>>3) + v10) & v5
								t28 := int32(int8(m.memory[uint32(t27+v10)]))
								if t28 < i32(0) {
									goto l13
								}
								t29 := int64(load64(m.memory[uint32(v6):]))
								v10 = int32(uint32(int64(bits.TrailingZeros64(uint64(t29&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
							}
						l13:
							v9 = v18 & v9
							t30 := v6 + v10
							v15 = int32(uint32(v15) >> 25)
							m.memory[uint32(t30)] = byte(v15)
							m.memory[uint32(v6+(v10+i32(-8))&v5+i32(8))] = byte(v15)
							v10 = v6 + (v10^i32(-1))*i32(24)
							t31 := v10
							v2 = v11 + (v2^i32(-1))*i32(24)
							t32 := int64(load64(m.memory[int64(uint32(v2))+16:]))
							store64(m.memory[int64(uint32(t31))+16:], uint64(t32))
							t33 := int64(load64(m.memory[int64(uint32(v2))+8:]))
							store64(m.memory[int64(uint32(v10))+8:], uint64(t33))
							t34 := int64(load64(m.memory[uint32(v2):]))
							store64(m.memory[uint32(v10):], uint64(t34))
							v14 = v14 + i32(-1)
							if v14 != 0 {
								goto l14
							}
						}
					}
				l8:
					store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
					store32(m.memory[uint32(v0):], uint32(v6))
					store32(m.memory[int64(uint32(v0))+8:], uint32(v7-v3))
					if v4 == 0 {
						goto l15
					}
					t35 := v4
					v8 = (v4*i32(24) + i32(31)) & i32(-8)
					v1 = t35 + v8 + i32(9)
					if v1 == 0 {
						goto l15
					}
					v4 = v11 - v8
					t36 := int32(load32(m.memory[uint32(v4+i32(-4)):]))
					v8 = t36
					v6 = v8 & i32(-8)
					t37 := v6
					v8 = v8 & i32(3)
					p38 := i32(8)
					if v8 != 0 {
						p38 = i32(4)
					}
					if uint32(t37) < uint32(p38+v1) {
						m.fn3(i32(1274224), i32(46), i32(1274272))
						panic("unreachable")
					}
					if v8 == 0 {
						goto l17
					}
					if uint32(v6) > uint32(v1+i32(39)) {
						m.fn3(i32(1274288), i32(46), i32(1274336))
						panic("unreachable")
					}
				l17:
					m.fn1(v4)
					return i32(-1)
				}
				m.fn23(i32(8), v6)
				panic("unreachable")
			}
		}
	l1:
		{
			if v5 != 0 {
				goto l19
			}
			v1 = i32(0)
			goto l20
		l19:
			t39 := int32(load32(m.memory[uint32(v0):]))
			v8 = t39
			v1 = i32(0)
			{
				{
					t40 := v6
					var p41 int32
					if v5&i32(7) != i32(0) {
						p41 = 1
					}
					v6 = t40 + p41
					if v6 == i32(1) {
						goto l21
					}
					v11 = v6 & i32(1)
					v10 = v6 & i32(0x3ffffffe)
					v1 = i32(0)
				l22:
					{
						v6 = v8 + v1
						t42 := int64(load64(m.memory[uint32(v6):]))
						t43 := v6
						v9 = t42
						store64(m.memory[uint32(t43):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
						v6 = v6 + i32(8)
						t44 := int64(load64(m.memory[uint32(v6):]))
						t45 := v6
						v9 = t44
						store64(m.memory[uint32(t45):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
						v1 = v1 + i32(16)
						v10 = v10 + i32(-2)
						if v10 != 0 {
							goto l22
						}
					}
					if v11 == 0 {
						goto l23
					}
				}
			l21:
				v1 = v8 + v1
				t46 := int64(load64(m.memory[uint32(v1):]))
				t47 := v1
				v9 = t46
				store64(m.memory[uint32(t47):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
			}
		l23:
			{
				if uint32(v5) < uint32(i32(8)) {
					goto l24
				}
				t48 := int64(load64(m.memory[uint32(v8):]))
				store64(m.memory[uint32(v8+v5):], uint64(t48))
				goto l25
			}
		l24:
			if v5 == 0 {
				goto l25
			}
			memory_copy(m.memory, uint32(v8+i32(8)), uint32(v8), uint32(v5))
		l25:
			t49 := int64(load64(m.memory[int64(uint32(v2))+8:]))
			v16 = t49
			t50 := int64(load64(m.memory[uint32(v2):]))
			v18 = t50
			v6 = i32(0)
		l33:
			{
				t51 := v8
				v1 = v6
				v10 = t51 + v1
				t52 := int32(m.memory[uint32(v10)])
				if t52 != i32(128) {
					goto l26
				}
				v15 = v8 + (v1^i32(-1))*i32(24)
				v6 = v8 + (i32(0)-v1)*i32(24)
				v11 = v6 + i32(-16)
				v14 = v6 + i32(-20)
			l32:
				{
					t53 := int32(load32(m.memory[uint32(v14):]))
					t54 := int32(load32(m.memory[uint32(v11):]))
					t55 := m.fn64(v18, v16, t53, t54)
					t56 := v4
					v2 = int32(t55)
					v6 = t56 & v2
					v5 = v6
					{
						t57 := int64(load64(m.memory[uint32(v8+v6):]))
						v9 = t57 & i64(-0x7f7f7f7f7f7f7f80)
						if v9 != i64(0) {
							goto l27
						}
						v17 = i32(8)
						v5 = v6
					l28:
						{
							v5 = v5 + v17
							v17 = v17 + i32(8)
							t58 := v8
							v5 = v5 & v4
							t59 := int64(load64(m.memory[uint32(t58+v5):]))
							v9 = t59 & i64(-0x7f7f7f7f7f7f7f80)
							if v9 == 0 {
								goto l28
							}
						}
					}
				l27:
					{
						t60 := v8
						v5 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v9))))>>3) + v5) & v4
						t61 := int32(int8(m.memory[uint32(t60+v5)]))
						if t61 < i32(0) {
							goto l29
						}
						t62 := int64(load64(m.memory[uint32(v8):]))
						v5 = int32(uint32(int64(bits.TrailingZeros64(uint64(t62&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
					}
				l29:
					{
						if uint32((v5-v6^(v1-v6))&v4) < uint32(i32(8)) {
							goto l30
						}
						v6 = v8 + v5
						t63 := int32(m.memory[uint32(v6)])
						v17 = t63
						t64 := v6
						v2 = int32(uint32(v2) >> 25)
						m.memory[uint32(t64)] = byte(v2)
						m.memory[uint32(v8+(v5+i32(-8))&v4+i32(8))] = byte(v2)
						v6 = v8 + (v5^i32(-1))*i32(24)
						{
							if v17 != i32(255) {
								t68 := int64(load64(m.memory[uint32(v15):]))
								v9 = t68
								t69 := int64(load64(m.memory[uint32(v6):]))
								store64(m.memory[uint32(v15):], uint64(t69))
								store64(m.memory[uint32(v6):], uint64(v9))
								t70 := int64(load64(m.memory[int64(uint32(v6))+8:]))
								v9 = t70
								t71 := int64(load64(m.memory[int64(uint32(v15))+8:]))
								store64(m.memory[int64(uint32(v6))+8:], uint64(t71))
								store64(m.memory[int64(uint32(v15))+8:], uint64(v9))
								t72 := int32(load32(m.memory[int64(uint32(v15))+16:]))
								v5 = t72
								t73 := int32(load32(m.memory[int64(uint32(v6))+16:]))
								store32(m.memory[int64(uint32(v15))+16:], uint32(t73))
								t74 := int32(load32(m.memory[int64(uint32(v6))+20:]))
								v2 = t74
								t75 := int32(load32(m.memory[int64(uint32(v15))+20:]))
								store32(m.memory[int64(uint32(v6))+20:], uint32(t75))
								store32(m.memory[int64(uint32(v15))+20:], uint32(v2))
								store32(m.memory[int64(uint32(v6))+16:], uint32(v5))
								goto l32
							}
							m.memory[uint32(v10)] = byte(i32(255))
							m.memory[uint32(v8+v4&(v1+i32(-8))+i32(8))] = byte(i32(255))
							t65 := int64(load64(m.memory[int64(uint32(v15))+16:]))
							store64(m.memory[int64(uint32(v6))+16:], uint64(t65))
							t66 := int64(load64(m.memory[int64(uint32(v15))+8:]))
							store64(m.memory[int64(uint32(v6))+8:], uint64(t66))
							t67 := int64(load64(m.memory[uint32(v15):]))
							store64(m.memory[uint32(v6):], uint64(t67))
							goto l26
						}
					}
				l30:
				}
				t76 := v10
				v6 = int32(uint32(v2) >> 25)
				m.memory[uint32(t76)] = byte(v6)
				m.memory[uint32(v8+v4&(v1+i32(-8))+i32(8))] = byte(v6)
			}
		l26:
			v6 = v1 + i32(1)
			if v1 != v4 {
				goto l33
			}
			p77 := v7
			if uint32(v4) < uint32(i32(8)) {
				p77 = v4
			}
			v1 = p77
		}
	l20:
		store32(m.memory[int64(uint32(v0))+8:], uint32(v1-v3))
	l15:
		return i32(-1)
	}
l5:
	m.fn27(i32(1271632), i32(57), i32(1271660))
	panic("unreachable")
}
func (m *Module) fn66(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7, v8 int32
	var v9 int64
	var v10, v11 int32
	var v12, v13 int64
	var v14, v15 int32
	var v16 int64
	var v17 int32
	var v18 int64
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
		v2 = t0
		v3 = v2 + i32(1)
		if v3 == 0 {
			m.fn27(i32(1271632), i32(57), i32(1271660))
			panic("unreachable")
		}
		{
			t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t2 := v3
			v4 = t1
			t3 := v4
			v5 = v4 + i32(1)
			v6 = int32(uint32(v5) >> 3)
			v7 = v6 * i32(7)
			p4 := v7
			if uint32(v4) < uint32(i32(8)) {
				p4 = t3
			}
			v8 = p4
			if uint32(t2) <= uint32(int32(uint32(v8)>>1)) {
				goto l1
			}
			{
				{
					{
						v8 = v8 + i32(1)
						p5 := v3
						if uint32(v8) > uint32(v3) {
							p5 = v8
						}
						v3 = p5
						if uint32(v3) < uint32(i32(15)) {
							goto l2
						}
						if uint32(v3) > uint32(i32(0x1fffffff)) {
							m.fn27(i32(1271632), i32(57), i32(1271660))
							panic("unreachable")
						}
						t6 := int32(uint32(v3<<3) / uint32(i32(7)))
						v3 = i32_shr_u(i32(-1), int32(bits.LeadingZeros32(uint32(t6+i32(-1))))) + i32(1)
						goto l4
					}
				l2:
					p7 := v3&i32(8) + i32(8)
					if uint32(v3) < uint32(i32(4)) {
						p7 = i32(4)
					}
					v3 = p7
				}
			l4:
				v9 = int64(uint32(v3)) * i64(20)
				if int32(int64(uint64(v9)>>32)) != 0 {
					goto l5
				}
				v6 = int32(v9)
				if uint32(v6) > uint32(i32(-8)) {
					goto l5
				}
				v8 = v3 + i32(8)
				t8 := v8
				v10 = (v6 + i32(7)) & i32(-8)
				v6 = t8 + v10
				if uint32(v6) < uint32(v8) {
					goto l5
				}
				if uint32(v6) > uint32(i32(0x7ffffff8)) {
					goto l5
				}
				t9 := m.fn7(v6)
				v5 = t9
				if v5 != 0 {
					v6 = v5 + v10
					if v8 == 0 {
						goto l7
					}
					memory_fill(m.memory, uint32(v6), i32(255), uint32(v8))
				l7:
					v5 = v3 + i32(-1)
					p10 := int32(uint32(v3)>>3) * i32(7)
					if uint32(v3) < uint32(i32(9)) {
						p10 = v5
					}
					v7 = p10
					t11 := int32(load32(m.memory[uint32(v0):]))
					v11 = t11
					{
						if v2 == 0 {
							goto l8
						}
						t12 := int64(load64(m.memory[uint32(v11):]))
						v9 = (t12 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
						t13 := int64(load64(m.memory[int64(uint32(v1))+8:]))
						v12 = t13
						t14 := int64(load64(m.memory[uint32(v1):]))
						v13 = t14
						v8 = v11
						v14 = v2
						v3 = i32(0)
					l14:
						{
							if v9 != i64(0) {
								goto l9
							}
						l10:
							{
								v3 = v3 + i32(8)
								v8 = v8 + i32(8)
								t15 := int64(load64(m.memory[uint32(v8):]))
								v9 = t15 & i64(-0x7f7f7f7f7f7f7f80)
								if v9 == i64(-0x7f7f7f7f7f7f7f80) {
									goto l10
								}
							}
							v9 = v9 ^ i64(-0x7f7f7f7f7f7f7f80)
						l9:
							{
								t16 := v6
								t17 := v5
								t18 := v13
								t19 := v12
								t20 := v11
								v1 = int32(uint32(int64(bits.TrailingZeros64(uint64(v9))))>>3) + v3
								v10 = t20 + (i32(0)-v1)*i32(20)
								t21 := int32(load32(m.memory[uint32(v10+i32(-16)):]))
								t22 := int32(load32(m.memory[uint32(v10+i32(-12)):]))
								t23 := m.fn64(t18, t19, t21, t22)
								v15 = int32(t23)
								v10 = t17 & v15
								t24 := int64(load64(m.memory[uint32(t16+v10):]))
								v16 = t24 & i64(-0x7f7f7f7f7f7f7f80)
								if v16 != i64(0) {
									goto l11
								}
								v17 = i32(8)
							l12:
								{
									v10 = v10 + v17
									v17 = v17 + i32(8)
									t25 := v6
									v10 = v10 & v5
									t26 := int64(load64(m.memory[uint32(t25+v10):]))
									v16 = t26 & i64(-0x7f7f7f7f7f7f7f80)
									if v16 == 0 {
										goto l12
									}
								}
							}
						l11:
							v18 = v9 + i64(-1)
							{
								t27 := v6
								v10 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v16))))>>3) + v10) & v5
								t28 := int32(int8(m.memory[uint32(t27+v10)]))
								if t28 < i32(0) {
									goto l13
								}
								t29 := int64(load64(m.memory[uint32(v6):]))
								v10 = int32(uint32(int64(bits.TrailingZeros64(uint64(t29&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
							}
						l13:
							v9 = v18 & v9
							t30 := v6 + v10
							v15 = int32(uint32(v15) >> 25)
							m.memory[uint32(t30)] = byte(v15)
							m.memory[uint32(v6+(v10+i32(-8))&v5+i32(8))] = byte(v15)
							v10 = v6 + (v10^i32(-1))*i32(20)
							t31 := v10
							v1 = v11 + (v1^i32(-1))*i32(20)
							t32 := int32(load32(m.memory[int64(uint32(v1))+16:]))
							store32(m.memory[int64(uint32(t31))+16:], uint32(t32))
							t33 := int64(load64(m.memory[int64(uint32(v1))+8:]))
							store64(m.memory[int64(uint32(v10))+8:], uint64(t33))
							t34 := int64(load64(m.memory[uint32(v1):]))
							store64(m.memory[uint32(v10):], uint64(t34))
							v14 = v14 + i32(-1)
							if v14 != 0 {
								goto l14
							}
						}
					}
				l8:
					store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
					store32(m.memory[uint32(v0):], uint32(v6))
					store32(m.memory[int64(uint32(v0))+8:], uint32(v7-v2))
					if v4 == 0 {
						goto l15
					}
					t35 := v4
					v8 = (v4*i32(20) + i32(27)) & i32(-8)
					v3 = t35 + v8 + i32(9)
					if v3 == 0 {
						goto l15
					}
					v4 = v11 - v8
					t36 := int32(load32(m.memory[uint32(v4+i32(-4)):]))
					v8 = t36
					v6 = v8 & i32(-8)
					t37 := v6
					v8 = v8 & i32(3)
					p38 := i32(8)
					if v8 != 0 {
						p38 = i32(4)
					}
					if uint32(t37) < uint32(p38+v3) {
						m.fn3(i32(1274224), i32(46), i32(1274272))
						panic("unreachable")
					}
					if v8 == 0 {
						goto l17
					}
					if uint32(v6) > uint32(v3+i32(39)) {
						m.fn3(i32(1274288), i32(46), i32(1274336))
						panic("unreachable")
					}
				l17:
					m.fn1(v4)
					return i32(-1)
				}
				m.fn23(i32(8), v6)
				panic("unreachable")
			}
		}
	l1:
		{
			if v5 != 0 {
				goto l19
			}
			v3 = i32(0)
			goto l20
		l19:
			t39 := int32(load32(m.memory[uint32(v0):]))
			v8 = t39
			v3 = i32(0)
			{
				{
					t40 := v6
					var p41 int32
					if v5&i32(7) != i32(0) {
						p41 = 1
					}
					v6 = t40 + p41
					if v6 == i32(1) {
						goto l21
					}
					v11 = v6 & i32(1)
					v10 = v6 & i32(0x3ffffffe)
					v3 = i32(0)
				l22:
					{
						v6 = v8 + v3
						t42 := int64(load64(m.memory[uint32(v6):]))
						t43 := v6
						v9 = t42
						store64(m.memory[uint32(t43):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
						v6 = v6 + i32(8)
						t44 := int64(load64(m.memory[uint32(v6):]))
						t45 := v6
						v9 = t44
						store64(m.memory[uint32(t45):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
						v3 = v3 + i32(16)
						v10 = v10 + i32(-2)
						if v10 != 0 {
							goto l22
						}
					}
					if v11 == 0 {
						goto l23
					}
				}
			l21:
				v3 = v8 + v3
				t46 := int64(load64(m.memory[uint32(v3):]))
				t47 := v3
				v9 = t46
				store64(m.memory[uint32(t47):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
			}
		l23:
			{
				if uint32(v5) < uint32(i32(8)) {
					goto l24
				}
				t48 := int64(load64(m.memory[uint32(v8):]))
				store64(m.memory[uint32(v8+v5):], uint64(t48))
				goto l25
			}
		l24:
			if v5 == 0 {
				goto l25
			}
			memory_copy(m.memory, uint32(v8+i32(8)), uint32(v8), uint32(v5))
		l25:
			t49 := int64(load64(m.memory[int64(uint32(v1))+8:]))
			v16 = t49
			t50 := int64(load64(m.memory[uint32(v1):]))
			v18 = t50
			v6 = i32(0)
		l33:
			{
				t51 := v8
				v3 = v6
				v10 = t51 + v3
				t52 := int32(m.memory[uint32(v10)])
				if t52 != i32(128) {
					goto l26
				}
				v15 = v8 + (v3^i32(-1))*i32(20)
				v6 = v8 + (i32(0)-v3)*i32(20)
				v11 = v6 + i32(-12)
				v14 = v6 + i32(-16)
			l32:
				{
					t53 := int32(load32(m.memory[uint32(v14):]))
					t54 := int32(load32(m.memory[uint32(v11):]))
					t55 := m.fn64(v18, v16, t53, t54)
					t56 := v4
					v1 = int32(t55)
					v6 = t56 & v1
					v5 = v6
					{
						t57 := int64(load64(m.memory[uint32(v8+v6):]))
						v9 = t57 & i64(-0x7f7f7f7f7f7f7f80)
						if v9 != i64(0) {
							goto l27
						}
						v17 = i32(8)
						v5 = v6
					l28:
						{
							v5 = v5 + v17
							v17 = v17 + i32(8)
							t58 := v8
							v5 = v5 & v4
							t59 := int64(load64(m.memory[uint32(t58+v5):]))
							v9 = t59 & i64(-0x7f7f7f7f7f7f7f80)
							if v9 == 0 {
								goto l28
							}
						}
					}
				l27:
					{
						t60 := v8
						v5 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v9))))>>3) + v5) & v4
						t61 := int32(int8(m.memory[uint32(t60+v5)]))
						if t61 < i32(0) {
							goto l29
						}
						t62 := int64(load64(m.memory[uint32(v8):]))
						v5 = int32(uint32(int64(bits.TrailingZeros64(uint64(t62&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
					}
				l29:
					{
						if uint32((v5-v6^(v3-v6))&v4) < uint32(i32(8)) {
							goto l30
						}
						v6 = v8 + v5
						t63 := int32(m.memory[uint32(v6)])
						v17 = t63
						t64 := v6
						v1 = int32(uint32(v1) >> 25)
						m.memory[uint32(t64)] = byte(v1)
						m.memory[uint32(v8+(v5+i32(-8))&v4+i32(8))] = byte(v1)
						v6 = v8 + (v5^i32(-1))*i32(20)
						{
							if v17 != i32(255) {
								t68 := int32(load32(m.memory[uint32(v15):]))
								v5 = t68
								t69 := int32(load32(m.memory[uint32(v6):]))
								store32(m.memory[uint32(v15):], uint32(t69))
								store32(m.memory[uint32(v6):], uint32(v5))
								t70 := int32(load32(m.memory[int64(uint32(v6))+4:]))
								v5 = t70
								t71 := int32(load32(m.memory[int64(uint32(v15))+4:]))
								store32(m.memory[int64(uint32(v6))+4:], uint32(t71))
								store32(m.memory[int64(uint32(v15))+4:], uint32(v5))
								t72 := int32(load32(m.memory[int64(uint32(v15))+8:]))
								v5 = t72
								t73 := int32(load32(m.memory[int64(uint32(v6))+8:]))
								store32(m.memory[int64(uint32(v15))+8:], uint32(t73))
								store32(m.memory[int64(uint32(v6))+8:], uint32(v5))
								t74 := int32(load32(m.memory[int64(uint32(v6))+12:]))
								v5 = t74
								t75 := int32(load32(m.memory[int64(uint32(v15))+12:]))
								store32(m.memory[int64(uint32(v6))+12:], uint32(t75))
								store32(m.memory[int64(uint32(v15))+12:], uint32(v5))
								t76 := int32(load32(m.memory[int64(uint32(v15))+16:]))
								v5 = t76
								t77 := int32(load32(m.memory[int64(uint32(v6))+16:]))
								store32(m.memory[int64(uint32(v15))+16:], uint32(t77))
								store32(m.memory[int64(uint32(v6))+16:], uint32(v5))
								goto l32
							}
							m.memory[uint32(v10)] = byte(i32(255))
							m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(i32(255))
							t65 := int32(load32(m.memory[int64(uint32(v15))+16:]))
							store32(m.memory[int64(uint32(v6))+16:], uint32(t65))
							t66 := int64(load64(m.memory[int64(uint32(v15))+8:]))
							store64(m.memory[int64(uint32(v6))+8:], uint64(t66))
							t67 := int64(load64(m.memory[uint32(v15):]))
							store64(m.memory[uint32(v6):], uint64(t67))
							goto l26
						}
					}
				l30:
				}
				t78 := v10
				v6 = int32(uint32(v1) >> 25)
				m.memory[uint32(t78)] = byte(v6)
				m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(v6)
			}
		l26:
			v6 = v3 + i32(1)
			if v3 != v4 {
				goto l33
			}
			p79 := v7
			if uint32(v4) < uint32(i32(8)) {
				p79 = v4
			}
			v3 = p79
		}
	l20:
		store32(m.memory[int64(uint32(v0))+8:], uint32(v3-v2))
	l15:
		return i32(-1)
	}
l5:
	m.fn27(i32(1271632), i32(57), i32(1271660))
	panic("unreachable")
}
func (m *Module) fn67(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7, v8 int32
	var v9 int64
	var v10, v11 int32
	var v12, v13 int64
	var v14, v15 int32
	var v16 int64
	var v17 int32
	var v18 int64
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
		v2 = t0
		v3 = v2 + i32(1)
		if v3 == 0 {
			m.fn27(i32(1271632), i32(57), i32(1271660))
			panic("unreachable")
		}
		{
			t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t2 := v3
			v4 = t1
			t3 := v4
			v5 = v4 + i32(1)
			v6 = int32(uint32(v5) >> 3)
			v7 = v6 * i32(7)
			p4 := v7
			if uint32(v4) < uint32(i32(8)) {
				p4 = t3
			}
			v8 = p4
			if uint32(t2) <= uint32(int32(uint32(v8)>>1)) {
				goto l1
			}
			{
				{
					{
						v8 = v8 + i32(1)
						p5 := v3
						if uint32(v8) > uint32(v3) {
							p5 = v8
						}
						v3 = p5
						if uint32(v3) < uint32(i32(15)) {
							goto l2
						}
						if uint32(v3) > uint32(i32(0x1fffffff)) {
							m.fn27(i32(1271632), i32(57), i32(1271660))
							panic("unreachable")
						}
						t6 := int32(uint32(v3<<3) / uint32(i32(7)))
						v3 = i32_shr_u(i32(-1), int32(bits.LeadingZeros32(uint32(t6+i32(-1))))) + i32(1)
						goto l4
					}
				l2:
					p7 := v3&i32(8) + i32(8)
					if uint32(v3) < uint32(i32(4)) {
						p7 = i32(4)
					}
					v3 = p7
				}
			l4:
				v9 = int64(uint32(v3)) * i64(24)
				if int32(int64(uint64(v9)>>32)) != 0 {
					goto l5
				}
				v8 = v3 + i32(8)
				t8 := v8
				v10 = int32(v9)
				v6 = t8 + v10
				if uint32(v6) < uint32(v8) {
					goto l5
				}
				if uint32(v6) > uint32(i32(0x7ffffff8)) {
					goto l5
				}
				t9 := m.fn7(v6)
				v5 = t9
				if v5 != 0 {
					v6 = v5 + v10
					if v8 == 0 {
						goto l7
					}
					memory_fill(m.memory, uint32(v6), i32(255), uint32(v8))
				l7:
					v5 = v3 + i32(-1)
					p10 := int32(uint32(v3)>>3) * i32(7)
					if uint32(v3) < uint32(i32(9)) {
						p10 = v5
					}
					v7 = p10
					t11 := int32(load32(m.memory[uint32(v0):]))
					v11 = t11
					{
						if v2 == 0 {
							goto l8
						}
						t12 := int64(load64(m.memory[uint32(v11):]))
						v9 = (t12 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
						t13 := int64(load64(m.memory[int64(uint32(v1))+8:]))
						v12 = t13
						t14 := int64(load64(m.memory[uint32(v1):]))
						v13 = t14
						v8 = v11
						v14 = v2
						v3 = i32(0)
					l14:
						{
							if v9 != i64(0) {
								goto l9
							}
						l10:
							{
								v3 = v3 + i32(8)
								v8 = v8 + i32(8)
								t15 := int64(load64(m.memory[uint32(v8):]))
								v9 = t15 & i64(-0x7f7f7f7f7f7f7f80)
								if v9 == i64(-0x7f7f7f7f7f7f7f80) {
									goto l10
								}
							}
							v9 = v9 ^ i64(-0x7f7f7f7f7f7f7f80)
						l9:
							{
								t16 := v6
								t17 := v5
								t18 := v13
								t19 := v12
								t20 := v11
								v1 = int32(uint32(int64(bits.TrailingZeros64(uint64(v9))))>>3) + v3
								v10 = t20 + (i32(0)-v1)*i32(24)
								t21 := int32(load32(m.memory[uint32(v10+i32(-20)):]))
								t22 := int32(load32(m.memory[uint32(v10+i32(-16)):]))
								t23 := m.fn64(t18, t19, t21, t22)
								v15 = int32(t23)
								v10 = t17 & v15
								t24 := int64(load64(m.memory[uint32(t16+v10):]))
								v16 = t24 & i64(-0x7f7f7f7f7f7f7f80)
								if v16 != i64(0) {
									goto l11
								}
								v17 = i32(8)
							l12:
								{
									v10 = v10 + v17
									v17 = v17 + i32(8)
									t25 := v6
									v10 = v10 & v5
									t26 := int64(load64(m.memory[uint32(t25+v10):]))
									v16 = t26 & i64(-0x7f7f7f7f7f7f7f80)
									if v16 == 0 {
										goto l12
									}
								}
							}
						l11:
							v18 = v9 + i64(-1)
							{
								t27 := v6
								v10 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v16))))>>3) + v10) & v5
								t28 := int32(int8(m.memory[uint32(t27+v10)]))
								if t28 < i32(0) {
									goto l13
								}
								t29 := int64(load64(m.memory[uint32(v6):]))
								v10 = int32(uint32(int64(bits.TrailingZeros64(uint64(t29&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
							}
						l13:
							v9 = v18 & v9
							t30 := v6 + v10
							v15 = int32(uint32(v15) >> 25)
							m.memory[uint32(t30)] = byte(v15)
							m.memory[uint32(v6+(v10+i32(-8))&v5+i32(8))] = byte(v15)
							v10 = v6 + (v10^i32(-1))*i32(24)
							t31 := v10
							v1 = v11 + (v1^i32(-1))*i32(24)
							t32 := int64(load64(m.memory[int64(uint32(v1))+16:]))
							store64(m.memory[int64(uint32(t31))+16:], uint64(t32))
							t33 := int64(load64(m.memory[int64(uint32(v1))+8:]))
							store64(m.memory[int64(uint32(v10))+8:], uint64(t33))
							t34 := int64(load64(m.memory[uint32(v1):]))
							store64(m.memory[uint32(v10):], uint64(t34))
							v14 = v14 + i32(-1)
							if v14 != 0 {
								goto l14
							}
						}
					}
				l8:
					store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
					store32(m.memory[uint32(v0):], uint32(v6))
					store32(m.memory[int64(uint32(v0))+8:], uint32(v7-v2))
					if v4 == 0 {
						goto l15
					}
					t35 := v4
					v8 = (v4*i32(24) + i32(31)) & i32(-8)
					v3 = t35 + v8 + i32(9)
					if v3 == 0 {
						goto l15
					}
					v4 = v11 - v8
					t36 := int32(load32(m.memory[uint32(v4+i32(-4)):]))
					v8 = t36
					v6 = v8 & i32(-8)
					t37 := v6
					v8 = v8 & i32(3)
					p38 := i32(8)
					if v8 != 0 {
						p38 = i32(4)
					}
					if uint32(t37) < uint32(p38+v3) {
						m.fn3(i32(1274224), i32(46), i32(1274272))
						panic("unreachable")
					}
					if v8 == 0 {
						goto l17
					}
					if uint32(v6) > uint32(v3+i32(39)) {
						m.fn3(i32(1274288), i32(46), i32(1274336))
						panic("unreachable")
					}
				l17:
					m.fn1(v4)
					return i32(-1)
				}
				m.fn23(i32(8), v6)
				panic("unreachable")
			}
		}
	l1:
		{
			if v5 != 0 {
				goto l19
			}
			v3 = i32(0)
			goto l20
		l19:
			t39 := int32(load32(m.memory[uint32(v0):]))
			v8 = t39
			v3 = i32(0)
			{
				{
					t40 := v6
					var p41 int32
					if v5&i32(7) != i32(0) {
						p41 = 1
					}
					v6 = t40 + p41
					if v6 == i32(1) {
						goto l21
					}
					v11 = v6 & i32(1)
					v10 = v6 & i32(0x3ffffffe)
					v3 = i32(0)
				l22:
					{
						v6 = v8 + v3
						t42 := int64(load64(m.memory[uint32(v6):]))
						t43 := v6
						v9 = t42
						store64(m.memory[uint32(t43):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
						v6 = v6 + i32(8)
						t44 := int64(load64(m.memory[uint32(v6):]))
						t45 := v6
						v9 = t44
						store64(m.memory[uint32(t45):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
						v3 = v3 + i32(16)
						v10 = v10 + i32(-2)
						if v10 != 0 {
							goto l22
						}
					}
					if v11 == 0 {
						goto l23
					}
				}
			l21:
				v3 = v8 + v3
				t46 := int64(load64(m.memory[uint32(v3):]))
				t47 := v3
				v9 = t46
				store64(m.memory[uint32(t47):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
			}
		l23:
			{
				if uint32(v5) < uint32(i32(8)) {
					goto l24
				}
				t48 := int64(load64(m.memory[uint32(v8):]))
				store64(m.memory[uint32(v8+v5):], uint64(t48))
				goto l25
			}
		l24:
			if v5 == 0 {
				goto l25
			}
			memory_copy(m.memory, uint32(v8+i32(8)), uint32(v8), uint32(v5))
		l25:
			t49 := int64(load64(m.memory[int64(uint32(v1))+8:]))
			v16 = t49
			t50 := int64(load64(m.memory[uint32(v1):]))
			v18 = t50
			v6 = i32(0)
		l33:
			{
				t51 := v8
				v3 = v6
				v10 = t51 + v3
				t52 := int32(m.memory[uint32(v10)])
				if t52 != i32(128) {
					goto l26
				}
				v15 = v8 + (v3^i32(-1))*i32(24)
				v6 = v8 + (i32(0)-v3)*i32(24)
				v11 = v6 + i32(-16)
				v14 = v6 + i32(-20)
			l32:
				{
					t53 := int32(load32(m.memory[uint32(v14):]))
					t54 := int32(load32(m.memory[uint32(v11):]))
					t55 := m.fn64(v18, v16, t53, t54)
					t56 := v4
					v1 = int32(t55)
					v6 = t56 & v1
					v5 = v6
					{
						t57 := int64(load64(m.memory[uint32(v8+v6):]))
						v9 = t57 & i64(-0x7f7f7f7f7f7f7f80)
						if v9 != i64(0) {
							goto l27
						}
						v17 = i32(8)
						v5 = v6
					l28:
						{
							v5 = v5 + v17
							v17 = v17 + i32(8)
							t58 := v8
							v5 = v5 & v4
							t59 := int64(load64(m.memory[uint32(t58+v5):]))
							v9 = t59 & i64(-0x7f7f7f7f7f7f7f80)
							if v9 == 0 {
								goto l28
							}
						}
					}
				l27:
					{
						t60 := v8
						v5 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v9))))>>3) + v5) & v4
						t61 := int32(int8(m.memory[uint32(t60+v5)]))
						if t61 < i32(0) {
							goto l29
						}
						t62 := int64(load64(m.memory[uint32(v8):]))
						v5 = int32(uint32(int64(bits.TrailingZeros64(uint64(t62&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
					}
				l29:
					{
						if uint32((v5-v6^(v3-v6))&v4) < uint32(i32(8)) {
							goto l30
						}
						v6 = v8 + v5
						t63 := int32(m.memory[uint32(v6)])
						v17 = t63
						t64 := v6
						v1 = int32(uint32(v1) >> 25)
						m.memory[uint32(t64)] = byte(v1)
						m.memory[uint32(v8+(v5+i32(-8))&v4+i32(8))] = byte(v1)
						v6 = v8 + (v5^i32(-1))*i32(24)
						{
							if v17 != i32(255) {
								t68 := int64(load64(m.memory[uint32(v15):]))
								v9 = t68
								t69 := int64(load64(m.memory[uint32(v6):]))
								store64(m.memory[uint32(v15):], uint64(t69))
								store64(m.memory[uint32(v6):], uint64(v9))
								t70 := int64(load64(m.memory[int64(uint32(v6))+8:]))
								v9 = t70
								t71 := int64(load64(m.memory[int64(uint32(v15))+8:]))
								store64(m.memory[int64(uint32(v6))+8:], uint64(t71))
								store64(m.memory[int64(uint32(v15))+8:], uint64(v9))
								t72 := int32(load32(m.memory[int64(uint32(v15))+16:]))
								v5 = t72
								t73 := int32(load32(m.memory[int64(uint32(v6))+16:]))
								store32(m.memory[int64(uint32(v15))+16:], uint32(t73))
								t74 := int32(load32(m.memory[int64(uint32(v6))+20:]))
								v1 = t74
								t75 := int32(load32(m.memory[int64(uint32(v15))+20:]))
								store32(m.memory[int64(uint32(v6))+20:], uint32(t75))
								store32(m.memory[int64(uint32(v15))+20:], uint32(v1))
								store32(m.memory[int64(uint32(v6))+16:], uint32(v5))
								goto l32
							}
							m.memory[uint32(v10)] = byte(i32(255))
							m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(i32(255))
							t65 := int64(load64(m.memory[int64(uint32(v15))+16:]))
							store64(m.memory[int64(uint32(v6))+16:], uint64(t65))
							t66 := int64(load64(m.memory[int64(uint32(v15))+8:]))
							store64(m.memory[int64(uint32(v6))+8:], uint64(t66))
							t67 := int64(load64(m.memory[uint32(v15):]))
							store64(m.memory[uint32(v6):], uint64(t67))
							goto l26
						}
					}
				l30:
				}
				t76 := v10
				v6 = int32(uint32(v1) >> 25)
				m.memory[uint32(t76)] = byte(v6)
				m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(v6)
			}
		l26:
			v6 = v3 + i32(1)
			if v3 != v4 {
				goto l33
			}
			p77 := v7
			if uint32(v4) < uint32(i32(8)) {
				p77 = v4
			}
			v3 = p77
		}
	l20:
		store32(m.memory[int64(uint32(v0))+8:], uint32(v3-v2))
	l15:
		return i32(-1)
	}
l5:
	m.fn27(i32(1271632), i32(57), i32(1271660))
	panic("unreachable")
}
func (m *Module) fn68(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7, v8 int32
	var v9 int64
	var v10, v11 int32
	var v12, v13 int64
	var v14, v15 int32
	var v16 int64
	var v17 int32
	var v18 int64
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
		v2 = t0
		v3 = v2 + i32(1)
		if v3 == 0 {
			m.fn27(i32(1271632), i32(57), i32(1271660))
			panic("unreachable")
		}
		{
			t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t2 := v3
			v4 = t1
			t3 := v4
			v5 = v4 + i32(1)
			v6 = int32(uint32(v5) >> 3)
			v7 = v6 * i32(7)
			p4 := v7
			if uint32(v4) < uint32(i32(8)) {
				p4 = t3
			}
			v8 = p4
			if uint32(t2) <= uint32(int32(uint32(v8)>>1)) {
				goto l1
			}
			{
				{
					{
						v8 = v8 + i32(1)
						p5 := v3
						if uint32(v8) > uint32(v3) {
							p5 = v8
						}
						v3 = p5
						if uint32(v3) < uint32(i32(15)) {
							goto l2
						}
						if uint32(v3) > uint32(i32(0x1fffffff)) {
							m.fn27(i32(1271632), i32(57), i32(1271660))
							panic("unreachable")
						}
						t6 := int32(uint32(v3<<3) / uint32(i32(7)))
						v3 = i32_shr_u(i32(-1), int32(bits.LeadingZeros32(uint32(t6+i32(-1))))) + i32(1)
						goto l4
					}
				l2:
					p7 := v3&i32(8) + i32(8)
					if uint32(v3) < uint32(i32(4)) {
						p7 = i32(4)
					}
					v3 = p7
				}
			l4:
				v9 = int64(uint32(v3)) * i64(24)
				if int32(int64(uint64(v9)>>32)) != 0 {
					goto l5
				}
				v8 = v3 + i32(8)
				t8 := v8
				v10 = int32(v9)
				v6 = t8 + v10
				if uint32(v6) < uint32(v8) {
					goto l5
				}
				if uint32(v6) > uint32(i32(0x7ffffff8)) {
					goto l5
				}
				t9 := m.fn7(v6)
				v5 = t9
				if v5 != 0 {
					v6 = v5 + v10
					if v8 == 0 {
						goto l7
					}
					memory_fill(m.memory, uint32(v6), i32(255), uint32(v8))
				l7:
					v5 = v3 + i32(-1)
					p10 := int32(uint32(v3)>>3) * i32(7)
					if uint32(v3) < uint32(i32(9)) {
						p10 = v5
					}
					v7 = p10
					t11 := int32(load32(m.memory[uint32(v0):]))
					v11 = t11
					{
						if v2 == 0 {
							goto l8
						}
						t12 := int64(load64(m.memory[uint32(v11):]))
						v9 = (t12 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
						t13 := int64(load64(m.memory[int64(uint32(v1))+8:]))
						v12 = t13
						t14 := int64(load64(m.memory[uint32(v1):]))
						v13 = t14
						v8 = v11
						v14 = v2
						v3 = i32(0)
					l14:
						{
							if v9 != i64(0) {
								goto l9
							}
						l10:
							{
								v3 = v3 + i32(8)
								v8 = v8 + i32(8)
								t15 := int64(load64(m.memory[uint32(v8):]))
								v9 = t15 & i64(-0x7f7f7f7f7f7f7f80)
								if v9 == i64(-0x7f7f7f7f7f7f7f80) {
									goto l10
								}
							}
							v9 = v9 ^ i64(-0x7f7f7f7f7f7f7f80)
						l9:
							{
								t16 := v6
								t17 := v5
								t18 := v13
								t19 := v12
								t20 := v11
								v1 = int32(uint32(int64(bits.TrailingZeros64(uint64(v9))))>>3) + v3
								v10 = t20 + (i32(0)-v1)*i32(24)
								t21 := int32(load32(m.memory[uint32(v10+i32(-20)):]))
								t22 := int32(load32(m.memory[uint32(v10+i32(-16)):]))
								t23 := m.fn64(t18, t19, t21, t22)
								v15 = int32(t23)
								v10 = t17 & v15
								t24 := int64(load64(m.memory[uint32(t16+v10):]))
								v16 = t24 & i64(-0x7f7f7f7f7f7f7f80)
								if v16 != i64(0) {
									goto l11
								}
								v17 = i32(8)
							l12:
								{
									v10 = v10 + v17
									v17 = v17 + i32(8)
									t25 := v6
									v10 = v10 & v5
									t26 := int64(load64(m.memory[uint32(t25+v10):]))
									v16 = t26 & i64(-0x7f7f7f7f7f7f7f80)
									if v16 == 0 {
										goto l12
									}
								}
							}
						l11:
							v18 = v9 + i64(-1)
							{
								t27 := v6
								v10 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v16))))>>3) + v10) & v5
								t28 := int32(int8(m.memory[uint32(t27+v10)]))
								if t28 < i32(0) {
									goto l13
								}
								t29 := int64(load64(m.memory[uint32(v6):]))
								v10 = int32(uint32(int64(bits.TrailingZeros64(uint64(t29&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
							}
						l13:
							v9 = v18 & v9
							t30 := v6 + v10
							v15 = int32(uint32(v15) >> 25)
							m.memory[uint32(t30)] = byte(v15)
							m.memory[uint32(v6+(v10+i32(-8))&v5+i32(8))] = byte(v15)
							v10 = v6 + (v10^i32(-1))*i32(24)
							t31 := v10
							v1 = v11 + (v1^i32(-1))*i32(24)
							t32 := int64(load64(m.memory[int64(uint32(v1))+16:]))
							store64(m.memory[int64(uint32(t31))+16:], uint64(t32))
							t33 := int64(load64(m.memory[int64(uint32(v1))+8:]))
							store64(m.memory[int64(uint32(v10))+8:], uint64(t33))
							t34 := int64(load64(m.memory[uint32(v1):]))
							store64(m.memory[uint32(v10):], uint64(t34))
							v14 = v14 + i32(-1)
							if v14 != 0 {
								goto l14
							}
						}
					}
				l8:
					store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
					store32(m.memory[uint32(v0):], uint32(v6))
					store32(m.memory[int64(uint32(v0))+8:], uint32(v7-v2))
					if v4 == 0 {
						goto l15
					}
					t35 := v4
					v8 = (v4*i32(24) + i32(31)) & i32(-8)
					v3 = t35 + v8 + i32(9)
					if v3 == 0 {
						goto l15
					}
					v4 = v11 - v8
					t36 := int32(load32(m.memory[uint32(v4+i32(-4)):]))
					v8 = t36
					v6 = v8 & i32(-8)
					t37 := v6
					v8 = v8 & i32(3)
					p38 := i32(8)
					if v8 != 0 {
						p38 = i32(4)
					}
					if uint32(t37) < uint32(p38+v3) {
						m.fn3(i32(1274224), i32(46), i32(1274272))
						panic("unreachable")
					}
					if v8 == 0 {
						goto l17
					}
					if uint32(v6) > uint32(v3+i32(39)) {
						m.fn3(i32(1274288), i32(46), i32(1274336))
						panic("unreachable")
					}
				l17:
					m.fn1(v4)
					return i32(-1)
				}
				m.fn23(i32(8), v6)
				panic("unreachable")
			}
		}
	l1:
		{
			if v5 != 0 {
				goto l19
			}
			v3 = i32(0)
			goto l20
		l19:
			t39 := int32(load32(m.memory[uint32(v0):]))
			v8 = t39
			v3 = i32(0)
			{
				{
					t40 := v6
					var p41 int32
					if v5&i32(7) != i32(0) {
						p41 = 1
					}
					v6 = t40 + p41
					if v6 == i32(1) {
						goto l21
					}
					v11 = v6 & i32(1)
					v10 = v6 & i32(0x3ffffffe)
					v3 = i32(0)
				l22:
					{
						v6 = v8 + v3
						t42 := int64(load64(m.memory[uint32(v6):]))
						t43 := v6
						v9 = t42
						store64(m.memory[uint32(t43):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
						v6 = v6 + i32(8)
						t44 := int64(load64(m.memory[uint32(v6):]))
						t45 := v6
						v9 = t44
						store64(m.memory[uint32(t45):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
						v3 = v3 + i32(16)
						v10 = v10 + i32(-2)
						if v10 != 0 {
							goto l22
						}
					}
					if v11 == 0 {
						goto l23
					}
				}
			l21:
				v3 = v8 + v3
				t46 := int64(load64(m.memory[uint32(v3):]))
				t47 := v3
				v9 = t46
				store64(m.memory[uint32(t47):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
			}
		l23:
			{
				if uint32(v5) < uint32(i32(8)) {
					goto l24
				}
				t48 := int64(load64(m.memory[uint32(v8):]))
				store64(m.memory[uint32(v8+v5):], uint64(t48))
				goto l25
			}
		l24:
			if v5 == 0 {
				goto l25
			}
			memory_copy(m.memory, uint32(v8+i32(8)), uint32(v8), uint32(v5))
		l25:
			t49 := int64(load64(m.memory[int64(uint32(v1))+8:]))
			v16 = t49
			t50 := int64(load64(m.memory[uint32(v1):]))
			v18 = t50
			v6 = i32(0)
		l33:
			{
				t51 := v8
				v3 = v6
				v10 = t51 + v3
				t52 := int32(m.memory[uint32(v10)])
				if t52 != i32(128) {
					goto l26
				}
				v15 = v8 + (v3^i32(-1))*i32(24)
				v6 = v8 + (i32(0)-v3)*i32(24)
				v11 = v6 + i32(-16)
				v14 = v6 + i32(-20)
			l32:
				{
					t53 := int32(load32(m.memory[uint32(v14):]))
					t54 := int32(load32(m.memory[uint32(v11):]))
					t55 := m.fn64(v18, v16, t53, t54)
					t56 := v4
					v1 = int32(t55)
					v6 = t56 & v1
					v5 = v6
					{
						t57 := int64(load64(m.memory[uint32(v8+v6):]))
						v9 = t57 & i64(-0x7f7f7f7f7f7f7f80)
						if v9 != i64(0) {
							goto l27
						}
						v17 = i32(8)
						v5 = v6
					l28:
						{
							v5 = v5 + v17
							v17 = v17 + i32(8)
							t58 := v8
							v5 = v5 & v4
							t59 := int64(load64(m.memory[uint32(t58+v5):]))
							v9 = t59 & i64(-0x7f7f7f7f7f7f7f80)
							if v9 == 0 {
								goto l28
							}
						}
					}
				l27:
					{
						t60 := v8
						v5 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v9))))>>3) + v5) & v4
						t61 := int32(int8(m.memory[uint32(t60+v5)]))
						if t61 < i32(0) {
							goto l29
						}
						t62 := int64(load64(m.memory[uint32(v8):]))
						v5 = int32(uint32(int64(bits.TrailingZeros64(uint64(t62&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
					}
				l29:
					{
						if uint32((v5-v6^(v3-v6))&v4) < uint32(i32(8)) {
							goto l30
						}
						v6 = v8 + v5
						t63 := int32(m.memory[uint32(v6)])
						v17 = t63
						t64 := v6
						v1 = int32(uint32(v1) >> 25)
						m.memory[uint32(t64)] = byte(v1)
						m.memory[uint32(v8+(v5+i32(-8))&v4+i32(8))] = byte(v1)
						v6 = v8 + (v5^i32(-1))*i32(24)
						{
							if v17 != i32(255) {
								t68 := int64(load64(m.memory[uint32(v15):]))
								v9 = t68
								t69 := int64(load64(m.memory[uint32(v6):]))
								store64(m.memory[uint32(v15):], uint64(t69))
								store64(m.memory[uint32(v6):], uint64(v9))
								t70 := int64(load64(m.memory[int64(uint32(v6))+8:]))
								v9 = t70
								t71 := int64(load64(m.memory[int64(uint32(v15))+8:]))
								store64(m.memory[int64(uint32(v6))+8:], uint64(t71))
								store64(m.memory[int64(uint32(v15))+8:], uint64(v9))
								t72 := int32(load32(m.memory[int64(uint32(v15))+16:]))
								v5 = t72
								t73 := int32(load32(m.memory[int64(uint32(v6))+16:]))
								store32(m.memory[int64(uint32(v15))+16:], uint32(t73))
								t74 := int32(load32(m.memory[int64(uint32(v6))+20:]))
								v1 = t74
								t75 := int32(load32(m.memory[int64(uint32(v15))+20:]))
								store32(m.memory[int64(uint32(v6))+20:], uint32(t75))
								store32(m.memory[int64(uint32(v15))+20:], uint32(v1))
								store32(m.memory[int64(uint32(v6))+16:], uint32(v5))
								goto l32
							}
							m.memory[uint32(v10)] = byte(i32(255))
							m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(i32(255))
							t65 := int64(load64(m.memory[int64(uint32(v15))+16:]))
							store64(m.memory[int64(uint32(v6))+16:], uint64(t65))
							t66 := int64(load64(m.memory[int64(uint32(v15))+8:]))
							store64(m.memory[int64(uint32(v6))+8:], uint64(t66))
							t67 := int64(load64(m.memory[uint32(v15):]))
							store64(m.memory[uint32(v6):], uint64(t67))
							goto l26
						}
					}
				l30:
				}
				t76 := v10
				v6 = int32(uint32(v1) >> 25)
				m.memory[uint32(t76)] = byte(v6)
				m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(v6)
			}
		l26:
			v6 = v3 + i32(1)
			if v3 != v4 {
				goto l33
			}
			p77 := v7
			if uint32(v4) < uint32(i32(8)) {
				p77 = v4
			}
			v3 = p77
		}
	l20:
		store32(m.memory[int64(uint32(v0))+8:], uint32(v3-v2))
	l15:
		return i32(-1)
	}
l5:
	m.fn27(i32(1271632), i32(57), i32(1271660))
	panic("unreachable")
}
func (m *Module) fn69(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7, v8, v9, v10 int32
	var v11, v12, v13 int64
	var v14, v15, v16 int32
	var v17, v18 int64
	{
		{
			{
				t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
				v2 = t0
				v3 = v2 + i32(1)
				if v3 == 0 {
					m.fn27(i32(1271632), i32(57), i32(1271660))
					panic("unreachable")
				}
				{
					t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
					t2 := v3
					v4 = t1
					t3 := v4
					v5 = v4 + i32(1)
					v6 = int32(uint32(v5) >> 3)
					v7 = v6 * i32(7)
					p4 := v7
					if uint32(v4) < uint32(i32(8)) {
						p4 = t3
					}
					v8 = p4
					if uint32(t2) <= uint32(int32(uint32(v8)>>1)) {
						{
							if v5 != 0 {
								goto l6
							}
							v3 = i32(0)
							goto l7
						l6:
							t7 := int32(load32(m.memory[uint32(v0):]))
							v8 = t7
							v3 = i32(0)
							{
								{
									t8 := v6
									var p9 int32
									if v5&i32(7) != i32(0) {
										p9 = 1
									}
									v6 = t8 + p9
									if v6 == i32(1) {
										goto l8
									}
									v9 = v6 & i32(1)
									v10 = v6 & i32(0x3ffffffe)
									v3 = i32(0)
								l9:
									{
										v6 = v8 + v3
										t10 := int64(load64(m.memory[uint32(v6):]))
										t11 := v6
										v11 = t10
										store64(m.memory[uint32(t11):], uint64(int64(uint64(v11^i64(-1))>>7)&i64(72340172838076673)+(v11|i64(0x7f7f7f7f7f7f7f7f))))
										v6 = v6 + i32(8)
										t12 := int64(load64(m.memory[uint32(v6):]))
										t13 := v6
										v11 = t12
										store64(m.memory[uint32(t13):], uint64(int64(uint64(v11^i64(-1))>>7)&i64(72340172838076673)+(v11|i64(0x7f7f7f7f7f7f7f7f))))
										v3 = v3 + i32(16)
										v10 = v10 + i32(-2)
										if v10 != 0 {
											goto l9
										}
									}
									if v9 == 0 {
										goto l10
									}
								}
							l8:
								v3 = v8 + v3
								t14 := int64(load64(m.memory[uint32(v3):]))
								t15 := v3
								v11 = t14
								store64(m.memory[uint32(t15):], uint64(int64(uint64(v11^i64(-1))>>7)&i64(72340172838076673)+(v11|i64(0x7f7f7f7f7f7f7f7f))))
							}
						l10:
							{
								if uint32(v5) < uint32(i32(8)) {
									goto l11
								}
								t16 := int64(load64(m.memory[uint32(v8):]))
								store64(m.memory[uint32(v8+v5):], uint64(t16))
								goto l12
							}
						l11:
							if v5 == 0 {
								goto l12
							}
							memory_copy(m.memory, uint32(v8+i32(8)), uint32(v8), uint32(v5))
						l12:
							t17 := int64(load64(m.memory[int64(uint32(v1))+8:]))
							v12 = t17
							t18 := int64(load64(m.memory[uint32(v1):]))
							v13 = t18
							v6 = i32(0)
						l20:
							{
								t19 := v8
								v3 = v6
								v10 = t19 + v3
								t20 := int32(m.memory[uint32(v10)])
								if t20 != i32(128) {
									goto l13
								}
								v6 = v8 - v3<<4
								v1 = v6 + i32(-8)
								v14 = v6 + i32(-12)
								v15 = v8 + (v3^i32(-1))<<4
							l19:
								{
									t21 := int32(load32(m.memory[uint32(v14):]))
									t22 := int32(load32(m.memory[uint32(v1):]))
									t23 := m.fn64(v13, v12, t21, t22)
									t24 := v4
									v9 = int32(t23)
									v6 = t24 & v9
									v5 = v6
									{
										t25 := int64(load64(m.memory[uint32(v8+v6):]))
										v11 = t25 & i64(-0x7f7f7f7f7f7f7f80)
										if v11 != i64(0) {
											goto l14
										}
										v16 = i32(8)
										v5 = v6
									l15:
										{
											v5 = v5 + v16
											v16 = v16 + i32(8)
											t26 := v8
											v5 = v5 & v4
											t27 := int64(load64(m.memory[uint32(t26+v5):]))
											v11 = t27 & i64(-0x7f7f7f7f7f7f7f80)
											if v11 == 0 {
												goto l15
											}
										}
									}
								l14:
									{
										t28 := v8
										v5 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v11))))>>3) + v5) & v4
										t29 := int32(int8(m.memory[uint32(t28+v5)]))
										if t29 < i32(0) {
											goto l16
										}
										t30 := int64(load64(m.memory[uint32(v8):]))
										v5 = int32(uint32(int64(bits.TrailingZeros64(uint64(t30&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
									}
								l16:
									{
										if uint32((v5-v6^(v3-v6))&v4) < uint32(i32(8)) {
											goto l17
										}
										v6 = v8 + v5
										t31 := int32(m.memory[uint32(v6)])
										v16 = t31
										t32 := v6
										v9 = int32(uint32(v9) >> 25)
										m.memory[uint32(t32)] = byte(v9)
										m.memory[uint32(v8+(v5+i32(-8))&v4+i32(8))] = byte(v9)
										v6 = v8 + (v5^i32(-1))<<4
										{
											if v16 != i32(255) {
												t35 := int64(load64(m.memory[uint32(v15):]))
												v11 = t35
												t36 := int64(load64(m.memory[uint32(v6):]))
												store64(m.memory[uint32(v15):], uint64(t36))
												store64(m.memory[uint32(v6):], uint64(v11))
												t37 := int64(load64(m.memory[int64(uint32(v15))+8:]))
												v11 = t37
												t38 := int64(load64(m.memory[int64(uint32(v6))+8:]))
												store64(m.memory[int64(uint32(v15))+8:], uint64(t38))
												store64(m.memory[int64(uint32(v6))+8:], uint64(v11))
												goto l19
											}
											m.memory[uint32(v10)] = byte(i32(255))
											m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(i32(255))
											t33 := int64(load64(m.memory[int64(uint32(v15))+8:]))
											store64(m.memory[int64(uint32(v6))+8:], uint64(t33))
											t34 := int64(load64(m.memory[uint32(v15):]))
											store64(m.memory[uint32(v6):], uint64(t34))
											goto l13
										}
									}
								l17:
								}
								t39 := v10
								v6 = int32(uint32(v9) >> 25)
								m.memory[uint32(t39)] = byte(v6)
								m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(v6)
							}
						l13:
							v6 = v3 + i32(1)
							if v3 != v4 {
								goto l20
							}
							p40 := v7
							if uint32(v4) < uint32(i32(8)) {
								p40 = v4
							}
							v3 = p40
						}
					l7:
						store32(m.memory[int64(uint32(v0))+8:], uint32(v3-v2))
						goto l21
					}
					v8 = v8 + i32(1)
					p5 := v3
					if uint32(v8) > uint32(v3) {
						p5 = v8
					}
					v3 = p5
					if uint32(v3) < uint32(i32(15)) {
						goto l2
					}
					{
						if uint32(v3) > uint32(i32(0x1fffffff)) {
							m.fn27(i32(1271632), i32(57), i32(1271660))
							panic("unreachable")
						}
						t6 := int32(uint32(v3<<3) / uint32(i32(7)))
						v3 = i32_shr_u(i32(-1), int32(bits.LeadingZeros32(uint32(t6+i32(-1)))))
						if uint32(v3) > uint32(i32(0xffffffe)) {
							goto l4
						}
						v3 = v3 + i32(1)
						goto l5
					}
				}
			}
		l2:
			p41 := v3&i32(8) + i32(8)
			if uint32(v3) < uint32(i32(4)) {
				p41 = i32(4)
			}
			v3 = p41
		}
	l5:
		v8 = v3 + i32(8)
		t42 := v8
		v10 = v3 << 4
		v6 = t42 + v10
		if uint32(v6) < uint32(v8) {
			goto l4
		}
		if uint32(v6) > uint32(i32(0x7ffffff8)) {
			goto l4
		}
		{
			t43 := m.fn7(v6)
			v5 = t43
			if v5 != 0 {
				v6 = v5 + v10
				if v8 == 0 {
					goto l23
				}
				memory_fill(m.memory, uint32(v6), i32(255), uint32(v8))
			l23:
				v5 = v3 + i32(-1)
				p44 := int32(uint32(v3)>>3) * i32(7)
				if uint32(v3) < uint32(i32(9)) {
					p44 = v5
				}
				v7 = p44
				t45 := int32(load32(m.memory[uint32(v0):]))
				v9 = t45
				{
					if v2 == 0 {
						goto l24
					}
					t46 := int64(load64(m.memory[uint32(v9):]))
					v11 = (t46 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
					t47 := int64(load64(m.memory[int64(uint32(v1))+8:]))
					v17 = t47
					t48 := int64(load64(m.memory[uint32(v1):]))
					v18 = t48
					v8 = v9
					v1 = v2
					v3 = i32(0)
				l30:
					{
						if v11 != i64(0) {
							goto l25
						}
					l26:
						{
							v3 = v3 + i32(8)
							v8 = v8 + i32(8)
							t49 := int64(load64(m.memory[uint32(v8):]))
							v11 = t49 & i64(-0x7f7f7f7f7f7f7f80)
							if v11 == i64(-0x7f7f7f7f7f7f7f80) {
								goto l26
							}
						}
						v11 = v11 ^ i64(-0x7f7f7f7f7f7f7f80)
					l25:
						{
							t50 := v6
							t51 := v5
							t52 := v18
							t53 := v17
							t54 := v9
							v14 = int32(uint32(int64(bits.TrailingZeros64(uint64(v11))))>>3) + v3
							v10 = t54 - v14<<4
							t55 := int32(load32(m.memory[uint32(v10+i32(-12)):]))
							t56 := int32(load32(m.memory[uint32(v10+i32(-8)):]))
							t57 := m.fn64(t52, t53, t55, t56)
							v15 = int32(t57)
							v10 = t51 & v15
							t58 := int64(load64(m.memory[uint32(t50+v10):]))
							v12 = t58 & i64(-0x7f7f7f7f7f7f7f80)
							if v12 != i64(0) {
								goto l27
							}
							v16 = i32(8)
						l28:
							{
								v10 = v10 + v16
								v16 = v16 + i32(8)
								t59 := v6
								v10 = v10 & v5
								t60 := int64(load64(m.memory[uint32(t59+v10):]))
								v12 = t60 & i64(-0x7f7f7f7f7f7f7f80)
								if v12 == 0 {
									goto l28
								}
							}
						}
					l27:
						v13 = v11 + i64(-1)
						{
							t61 := v6
							v10 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v12))))>>3) + v10) & v5
							t62 := int32(int8(m.memory[uint32(t61+v10)]))
							if t62 < i32(0) {
								goto l29
							}
							t63 := int64(load64(m.memory[uint32(v6):]))
							v10 = int32(uint32(int64(bits.TrailingZeros64(uint64(t63&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
						}
					l29:
						v11 = v13 & v11
						t64 := v6 + v10
						v15 = int32(uint32(v15) >> 25)
						m.memory[uint32(t64)] = byte(v15)
						m.memory[uint32(v6+(v10+i32(-8))&v5+i32(8))] = byte(v15)
						v10 = v6 + (v10^i32(-1))<<4
						t65 := v10
						v14 = v9 + (v14^i32(-1))<<4
						t66 := int64(load64(m.memory[int64(uint32(v14))+8:]))
						store64(m.memory[int64(uint32(t65))+8:], uint64(t66))
						t67 := int64(load64(m.memory[uint32(v14):]))
						store64(m.memory[uint32(v10):], uint64(t67))
						v1 = v1 + i32(-1)
						if v1 != 0 {
							goto l30
						}
					}
				}
			l24:
				store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
				store32(m.memory[uint32(v0):], uint32(v6))
				store32(m.memory[int64(uint32(v0))+8:], uint32(v7-v2))
				if v4 == 0 {
					goto l21
				}
				t68 := v4
				v8 = (v4<<4 + i32(23)) & i32(-16)
				v3 = t68 + v8 + i32(9)
				if v3 == 0 {
					goto l21
				}
				v4 = v9 - v8
				t69 := int32(load32(m.memory[uint32(v4+i32(-4)):]))
				v8 = t69
				v6 = v8 & i32(-8)
				t70 := v6
				v8 = v8 & i32(3)
				p71 := i32(8)
				if v8 != 0 {
					p71 = i32(4)
				}
				if uint32(t70) < uint32(p71+v3) {
					m.fn3(i32(1274224), i32(46), i32(1274272))
					panic("unreachable")
				}
				if v8 == 0 {
					goto l32
				}
				if uint32(v6) > uint32(v3+i32(39)) {
					m.fn3(i32(1274288), i32(46), i32(1274336))
					panic("unreachable")
				}
			l32:
				m.fn1(v4)
				return i32(-1)
			}
			m.fn23(i32(8), v6)
			panic("unreachable")
		}
	}
l21:
	return i32(-1)
l4:
	m.fn27(i32(1271632), i32(57), i32(1271660))
	panic("unreachable")
}
func (m *Module) fn70(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7, v8, v9, v10 int32
	var v11, v12, v13 int64
	var v14, v15, v16 int32
	var v17, v18 int64
	{
		{
			{
				t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
				v2 = t0
				v3 = v2 + i32(1)
				if v3 == 0 {
					m.fn27(i32(1271632), i32(57), i32(1271660))
					panic("unreachable")
				}
				{
					t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
					t2 := v3
					v4 = t1
					t3 := v4
					v5 = v4 + i32(1)
					v6 = int32(uint32(v5) >> 3)
					v7 = v6 * i32(7)
					p4 := v7
					if uint32(v4) < uint32(i32(8)) {
						p4 = t3
					}
					v8 = p4
					if uint32(t2) <= uint32(int32(uint32(v8)>>1)) {
						{
							if v5 != 0 {
								goto l6
							}
							v3 = i32(0)
							goto l7
						l6:
							t7 := int32(load32(m.memory[uint32(v0):]))
							v8 = t7
							v3 = i32(0)
							{
								{
									t8 := v6
									var p9 int32
									if v5&i32(7) != i32(0) {
										p9 = 1
									}
									v6 = t8 + p9
									if v6 == i32(1) {
										goto l8
									}
									v9 = v6 & i32(1)
									v10 = v6 & i32(0x3ffffffe)
									v3 = i32(0)
								l9:
									{
										v6 = v8 + v3
										t10 := int64(load64(m.memory[uint32(v6):]))
										t11 := v6
										v11 = t10
										store64(m.memory[uint32(t11):], uint64(int64(uint64(v11^i64(-1))>>7)&i64(72340172838076673)+(v11|i64(0x7f7f7f7f7f7f7f7f))))
										v6 = v6 + i32(8)
										t12 := int64(load64(m.memory[uint32(v6):]))
										t13 := v6
										v11 = t12
										store64(m.memory[uint32(t13):], uint64(int64(uint64(v11^i64(-1))>>7)&i64(72340172838076673)+(v11|i64(0x7f7f7f7f7f7f7f7f))))
										v3 = v3 + i32(16)
										v10 = v10 + i32(-2)
										if v10 != 0 {
											goto l9
										}
									}
									if v9 == 0 {
										goto l10
									}
								}
							l8:
								v3 = v8 + v3
								t14 := int64(load64(m.memory[uint32(v3):]))
								t15 := v3
								v11 = t14
								store64(m.memory[uint32(t15):], uint64(int64(uint64(v11^i64(-1))>>7)&i64(72340172838076673)+(v11|i64(0x7f7f7f7f7f7f7f7f))))
							}
						l10:
							{
								if uint32(v5) < uint32(i32(8)) {
									goto l11
								}
								t16 := int64(load64(m.memory[uint32(v8):]))
								store64(m.memory[uint32(v8+v5):], uint64(t16))
								goto l12
							}
						l11:
							if v5 == 0 {
								goto l12
							}
							memory_copy(m.memory, uint32(v8+i32(8)), uint32(v8), uint32(v5))
						l12:
							t17 := int64(load64(m.memory[int64(uint32(v1))+8:]))
							v12 = t17
							t18 := int64(load64(m.memory[uint32(v1):]))
							v13 = t18
							v6 = i32(0)
						l20:
							{
								t19 := v8
								v3 = v6
								v10 = t19 + v3
								t20 := int32(m.memory[uint32(v10)])
								if t20 != i32(128) {
									goto l13
								}
								v6 = v8 - v3<<4
								v1 = v6 + i32(-8)
								v14 = v6 + i32(-12)
								v15 = v8 + (v3^i32(-1))<<4
							l19:
								{
									t21 := int32(load32(m.memory[uint32(v14):]))
									t22 := int32(load32(m.memory[uint32(v1):]))
									t23 := m.fn64(v13, v12, t21, t22)
									t24 := v4
									v9 = int32(t23)
									v6 = t24 & v9
									v5 = v6
									{
										t25 := int64(load64(m.memory[uint32(v8+v6):]))
										v11 = t25 & i64(-0x7f7f7f7f7f7f7f80)
										if v11 != i64(0) {
											goto l14
										}
										v16 = i32(8)
										v5 = v6
									l15:
										{
											v5 = v5 + v16
											v16 = v16 + i32(8)
											t26 := v8
											v5 = v5 & v4
											t27 := int64(load64(m.memory[uint32(t26+v5):]))
											v11 = t27 & i64(-0x7f7f7f7f7f7f7f80)
											if v11 == 0 {
												goto l15
											}
										}
									}
								l14:
									{
										t28 := v8
										v5 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v11))))>>3) + v5) & v4
										t29 := int32(int8(m.memory[uint32(t28+v5)]))
										if t29 < i32(0) {
											goto l16
										}
										t30 := int64(load64(m.memory[uint32(v8):]))
										v5 = int32(uint32(int64(bits.TrailingZeros64(uint64(t30&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
									}
								l16:
									{
										if uint32((v5-v6^(v3-v6))&v4) < uint32(i32(8)) {
											goto l17
										}
										v6 = v8 + v5
										t31 := int32(m.memory[uint32(v6)])
										v16 = t31
										t32 := v6
										v9 = int32(uint32(v9) >> 25)
										m.memory[uint32(t32)] = byte(v9)
										m.memory[uint32(v8+(v5+i32(-8))&v4+i32(8))] = byte(v9)
										v6 = v8 + (v5^i32(-1))<<4
										{
											if v16 != i32(255) {
												t35 := int64(load64(m.memory[uint32(v15):]))
												v11 = t35
												t36 := int64(load64(m.memory[uint32(v6):]))
												store64(m.memory[uint32(v15):], uint64(t36))
												store64(m.memory[uint32(v6):], uint64(v11))
												t37 := int64(load64(m.memory[int64(uint32(v15))+8:]))
												v11 = t37
												t38 := int64(load64(m.memory[int64(uint32(v6))+8:]))
												store64(m.memory[int64(uint32(v15))+8:], uint64(t38))
												store64(m.memory[int64(uint32(v6))+8:], uint64(v11))
												goto l19
											}
											m.memory[uint32(v10)] = byte(i32(255))
											m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(i32(255))
											t33 := int64(load64(m.memory[int64(uint32(v15))+8:]))
											store64(m.memory[int64(uint32(v6))+8:], uint64(t33))
											t34 := int64(load64(m.memory[uint32(v15):]))
											store64(m.memory[uint32(v6):], uint64(t34))
											goto l13
										}
									}
								l17:
								}
								t39 := v10
								v6 = int32(uint32(v9) >> 25)
								m.memory[uint32(t39)] = byte(v6)
								m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(v6)
							}
						l13:
							v6 = v3 + i32(1)
							if v3 != v4 {
								goto l20
							}
							p40 := v7
							if uint32(v4) < uint32(i32(8)) {
								p40 = v4
							}
							v3 = p40
						}
					l7:
						store32(m.memory[int64(uint32(v0))+8:], uint32(v3-v2))
						goto l21
					}
					v8 = v8 + i32(1)
					p5 := v3
					if uint32(v8) > uint32(v3) {
						p5 = v8
					}
					v3 = p5
					if uint32(v3) < uint32(i32(15)) {
						goto l2
					}
					{
						if uint32(v3) > uint32(i32(0x1fffffff)) {
							m.fn27(i32(1271632), i32(57), i32(1271660))
							panic("unreachable")
						}
						t6 := int32(uint32(v3<<3) / uint32(i32(7)))
						v3 = i32_shr_u(i32(-1), int32(bits.LeadingZeros32(uint32(t6+i32(-1)))))
						if uint32(v3) > uint32(i32(0xffffffe)) {
							goto l4
						}
						v3 = v3 + i32(1)
						goto l5
					}
				}
			}
		l2:
			p41 := v3&i32(8) + i32(8)
			if uint32(v3) < uint32(i32(4)) {
				p41 = i32(4)
			}
			v3 = p41
		}
	l5:
		v8 = v3 + i32(8)
		t42 := v8
		v10 = v3 << 4
		v6 = t42 + v10
		if uint32(v6) < uint32(v8) {
			goto l4
		}
		if uint32(v6) > uint32(i32(0x7ffffff8)) {
			goto l4
		}
		{
			t43 := m.fn7(v6)
			v5 = t43
			if v5 != 0 {
				v6 = v5 + v10
				if v8 == 0 {
					goto l23
				}
				memory_fill(m.memory, uint32(v6), i32(255), uint32(v8))
			l23:
				v5 = v3 + i32(-1)
				p44 := int32(uint32(v3)>>3) * i32(7)
				if uint32(v3) < uint32(i32(9)) {
					p44 = v5
				}
				v7 = p44
				t45 := int32(load32(m.memory[uint32(v0):]))
				v9 = t45
				{
					if v2 == 0 {
						goto l24
					}
					t46 := int64(load64(m.memory[uint32(v9):]))
					v11 = (t46 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
					t47 := int64(load64(m.memory[int64(uint32(v1))+8:]))
					v17 = t47
					t48 := int64(load64(m.memory[uint32(v1):]))
					v18 = t48
					v8 = v9
					v1 = v2
					v3 = i32(0)
				l30:
					{
						if v11 != i64(0) {
							goto l25
						}
					l26:
						{
							v3 = v3 + i32(8)
							v8 = v8 + i32(8)
							t49 := int64(load64(m.memory[uint32(v8):]))
							v11 = t49 & i64(-0x7f7f7f7f7f7f7f80)
							if v11 == i64(-0x7f7f7f7f7f7f7f80) {
								goto l26
							}
						}
						v11 = v11 ^ i64(-0x7f7f7f7f7f7f7f80)
					l25:
						{
							t50 := v6
							t51 := v5
							t52 := v18
							t53 := v17
							t54 := v9
							v14 = int32(uint32(int64(bits.TrailingZeros64(uint64(v11))))>>3) + v3
							v10 = t54 - v14<<4
							t55 := int32(load32(m.memory[uint32(v10+i32(-12)):]))
							t56 := int32(load32(m.memory[uint32(v10+i32(-8)):]))
							t57 := m.fn64(t52, t53, t55, t56)
							v15 = int32(t57)
							v10 = t51 & v15
							t58 := int64(load64(m.memory[uint32(t50+v10):]))
							v12 = t58 & i64(-0x7f7f7f7f7f7f7f80)
							if v12 != i64(0) {
								goto l27
							}
							v16 = i32(8)
						l28:
							{
								v10 = v10 + v16
								v16 = v16 + i32(8)
								t59 := v6
								v10 = v10 & v5
								t60 := int64(load64(m.memory[uint32(t59+v10):]))
								v12 = t60 & i64(-0x7f7f7f7f7f7f7f80)
								if v12 == 0 {
									goto l28
								}
							}
						}
					l27:
						v13 = v11 + i64(-1)
						{
							t61 := v6
							v10 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v12))))>>3) + v10) & v5
							t62 := int32(int8(m.memory[uint32(t61+v10)]))
							if t62 < i32(0) {
								goto l29
							}
							t63 := int64(load64(m.memory[uint32(v6):]))
							v10 = int32(uint32(int64(bits.TrailingZeros64(uint64(t63&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
						}
					l29:
						v11 = v13 & v11
						t64 := v6 + v10
						v15 = int32(uint32(v15) >> 25)
						m.memory[uint32(t64)] = byte(v15)
						m.memory[uint32(v6+(v10+i32(-8))&v5+i32(8))] = byte(v15)
						v10 = v6 + (v10^i32(-1))<<4
						t65 := v10
						v14 = v9 + (v14^i32(-1))<<4
						t66 := int64(load64(m.memory[int64(uint32(v14))+8:]))
						store64(m.memory[int64(uint32(t65))+8:], uint64(t66))
						t67 := int64(load64(m.memory[uint32(v14):]))
						store64(m.memory[uint32(v10):], uint64(t67))
						v1 = v1 + i32(-1)
						if v1 != 0 {
							goto l30
						}
					}
				}
			l24:
				store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
				store32(m.memory[uint32(v0):], uint32(v6))
				store32(m.memory[int64(uint32(v0))+8:], uint32(v7-v2))
				if v4 == 0 {
					goto l21
				}
				t68 := v4
				v8 = (v4<<4 + i32(23)) & i32(-16)
				v3 = t68 + v8 + i32(9)
				if v3 == 0 {
					goto l21
				}
				v4 = v9 - v8
				t69 := int32(load32(m.memory[uint32(v4+i32(-4)):]))
				v8 = t69
				v6 = v8 & i32(-8)
				t70 := v6
				v8 = v8 & i32(3)
				p71 := i32(8)
				if v8 != 0 {
					p71 = i32(4)
				}
				if uint32(t70) < uint32(p71+v3) {
					m.fn3(i32(1274224), i32(46), i32(1274272))
					panic("unreachable")
				}
				if v8 == 0 {
					goto l32
				}
				if uint32(v6) > uint32(v3+i32(39)) {
					m.fn3(i32(1274288), i32(46), i32(1274336))
					panic("unreachable")
				}
			l32:
				m.fn1(v4)
				return i32(-1)
			}
			m.fn23(i32(8), v6)
			panic("unreachable")
		}
	}
l21:
	return i32(-1)
l4:
	m.fn27(i32(1271632), i32(57), i32(1271660))
	panic("unreachable")
}
func (m *Module) fn71(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7, v8 int32
	var v9 int64
	var v10, v11, v12 int32
	var v13, v14 int64
	var v15 int32
	var v16 int64
	var v17 int32
	var v18 int64
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
		v2 = t0
		v3 = v2 + i32(1)
		if v3 == 0 {
			m.fn27(i32(1271632), i32(57), i32(1271660))
			panic("unreachable")
		}
		{
			t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t2 := v3
			v4 = t1
			t3 := v4
			v5 = v4 + i32(1)
			v6 = int32(uint32(v5) >> 3)
			v7 = v6 * i32(7)
			p4 := v7
			if uint32(v4) < uint32(i32(8)) {
				p4 = t3
			}
			v8 = p4
			if uint32(t2) <= uint32(int32(uint32(v8)>>1)) {
				goto l1
			}
			{
				{
					{
						v8 = v8 + i32(1)
						p5 := v3
						if uint32(v8) > uint32(v3) {
							p5 = v8
						}
						v3 = p5
						if uint32(v3) < uint32(i32(15)) {
							goto l2
						}
						if uint32(v3) > uint32(i32(0x1fffffff)) {
							m.fn27(i32(1271632), i32(57), i32(1271660))
							panic("unreachable")
						}
						t6 := int32(uint32(v3<<3) / uint32(i32(7)))
						v3 = i32_shr_u(i32(-1), int32(bits.LeadingZeros32(uint32(t6+i32(-1))))) + i32(1)
						goto l4
					}
				l2:
					p7 := v3&i32(8) + i32(8)
					if uint32(v3) < uint32(i32(4)) {
						p7 = i32(4)
					}
					v3 = p7
				}
			l4:
				v9 = int64(uint32(v3)) * i64(36)
				if int32(int64(uint64(v9)>>32)) != 0 {
					goto l5
				}
				v6 = int32(v9)
				if uint32(v6) > uint32(i32(-8)) {
					goto l5
				}
				v8 = v3 + i32(8)
				t8 := v8
				v10 = (v6 + i32(7)) & i32(-8)
				v6 = t8 + v10
				if uint32(v6) < uint32(v8) {
					goto l5
				}
				if uint32(v6) > uint32(i32(0x7ffffff8)) {
					goto l5
				}
				t9 := m.fn7(v6)
				v5 = t9
				if v5 != 0 {
					v6 = v5 + v10
					if v8 == 0 {
						goto l7
					}
					memory_fill(m.memory, uint32(v6), i32(255), uint32(v8))
				l7:
					v11 = v3 + i32(-1)
					p10 := int32(uint32(v3)>>3) * i32(7)
					if uint32(v3) < uint32(i32(9)) {
						p10 = v11
					}
					v7 = p10
					t11 := int32(load32(m.memory[uint32(v0):]))
					v12 = t11
					{
						if v2 == 0 {
							goto l8
						}
						t12 := int64(load64(m.memory[uint32(v12):]))
						v9 = (t12 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
						t13 := int64(load64(m.memory[int64(uint32(v1))+8:]))
						v13 = t13
						t14 := int64(load64(m.memory[uint32(v1):]))
						v14 = t14
						v8 = v12
						v1 = v2
						v3 = i32(0)
					l14:
						{
							if v9 != i64(0) {
								goto l9
							}
						l10:
							{
								v3 = v3 + i32(8)
								v8 = v8 + i32(8)
								t15 := int64(load64(m.memory[uint32(v8):]))
								v9 = t15 & i64(-0x7f7f7f7f7f7f7f80)
								if v9 == i64(-0x7f7f7f7f7f7f7f80) {
									goto l10
								}
							}
							v9 = v9 ^ i64(-0x7f7f7f7f7f7f7f80)
						l9:
							{
								t16 := v6
								t17 := v11
								t18 := v14
								t19 := v13
								t20 := v12
								v5 = int32(uint32(int64(bits.TrailingZeros64(uint64(v9))))>>3) + v3
								v10 = t20 + (i32(0)-v5)*i32(36)
								t21 := int32(load32(m.memory[uint32(v10+i32(-32)):]))
								t22 := int32(load32(m.memory[uint32(v10+i32(-28)):]))
								t23 := m.fn64(t18, t19, t21, t22)
								v15 = int32(t23)
								v10 = t17 & v15
								t24 := int64(load64(m.memory[uint32(t16+v10):]))
								v16 = t24 & i64(-0x7f7f7f7f7f7f7f80)
								if v16 != i64(0) {
									goto l11
								}
								v17 = i32(8)
							l12:
								{
									v10 = v10 + v17
									v17 = v17 + i32(8)
									t25 := v6
									v10 = v10 & v11
									t26 := int64(load64(m.memory[uint32(t25+v10):]))
									v16 = t26 & i64(-0x7f7f7f7f7f7f7f80)
									if v16 == 0 {
										goto l12
									}
								}
							}
						l11:
							v18 = v9 + i64(-1)
							{
								t27 := v6
								v10 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v16))))>>3) + v10) & v11
								t28 := int32(int8(m.memory[uint32(t27+v10)]))
								if t28 < i32(0) {
									goto l13
								}
								t29 := int64(load64(m.memory[uint32(v6):]))
								v10 = int32(uint32(int64(bits.TrailingZeros64(uint64(t29&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
							}
						l13:
							v9 = v18 & v9
							t30 := v6 + v10
							v15 = int32(uint32(v15) >> 25)
							m.memory[uint32(t30)] = byte(v15)
							m.memory[uint32(v6+(v10+i32(-8))&v11+i32(8))] = byte(v15)
							v10 = v6 + (v10^i32(-1))*i32(36)
							t31 := v10
							v5 = v12 + (v5^i32(-1))*i32(36)
							t32 := int32(load32(m.memory[int64(uint32(v5))+32:]))
							store32(m.memory[int64(uint32(t31))+32:], uint32(t32))
							t33 := int64(load64(m.memory[int64(uint32(v5))+24:]))
							store64(m.memory[int64(uint32(v10))+24:], uint64(t33))
							t34 := int64(load64(m.memory[int64(uint32(v5))+16:]))
							store64(m.memory[int64(uint32(v10))+16:], uint64(t34))
							t35 := int64(load64(m.memory[int64(uint32(v5))+8:]))
							store64(m.memory[int64(uint32(v10))+8:], uint64(t35))
							t36 := int64(load64(m.memory[uint32(v5):]))
							store64(m.memory[uint32(v10):], uint64(t36))
							v1 = v1 + i32(-1)
							if v1 != 0 {
								goto l14
							}
						}
					}
				l8:
					store32(m.memory[int64(uint32(v0))+4:], uint32(v11))
					store32(m.memory[uint32(v0):], uint32(v6))
					store32(m.memory[int64(uint32(v0))+8:], uint32(v7-v2))
					if v4 == 0 {
						goto l15
					}
					t37 := v4
					v8 = (v4*i32(36) + i32(43)) & i32(-8)
					v3 = t37 + v8 + i32(9)
					if v3 == 0 {
						goto l15
					}
					v4 = v12 - v8
					t38 := int32(load32(m.memory[uint32(v4+i32(-4)):]))
					v8 = t38
					v6 = v8 & i32(-8)
					t39 := v6
					v8 = v8 & i32(3)
					p40 := i32(8)
					if v8 != 0 {
						p40 = i32(4)
					}
					if uint32(t39) < uint32(p40+v3) {
						m.fn3(i32(1274224), i32(46), i32(1274272))
						panic("unreachable")
					}
					if v8 == 0 {
						goto l17
					}
					if uint32(v6) > uint32(v3+i32(39)) {
						m.fn3(i32(1274288), i32(46), i32(1274336))
						panic("unreachable")
					}
				l17:
					m.fn1(v4)
					return i32(-1)
				}
				m.fn23(i32(8), v6)
				panic("unreachable")
			}
		}
	l1:
		{
			if v5 != 0 {
				goto l19
			}
			v3 = i32(0)
			goto l20
		l19:
			t41 := int32(load32(m.memory[uint32(v0):]))
			v8 = t41
			v3 = i32(0)
			{
				{
					t42 := v6
					var p43 int32
					if v5&i32(7) != i32(0) {
						p43 = 1
					}
					v6 = t42 + p43
					if v6 == i32(1) {
						goto l21
					}
					v11 = v6 & i32(1)
					v10 = v6 & i32(0x3ffffffe)
					v3 = i32(0)
				l22:
					{
						v6 = v8 + v3
						t44 := int64(load64(m.memory[uint32(v6):]))
						t45 := v6
						v9 = t44
						store64(m.memory[uint32(t45):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
						v6 = v6 + i32(8)
						t46 := int64(load64(m.memory[uint32(v6):]))
						t47 := v6
						v9 = t46
						store64(m.memory[uint32(t47):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
						v3 = v3 + i32(16)
						v10 = v10 + i32(-2)
						if v10 != 0 {
							goto l22
						}
					}
					if v11 == 0 {
						goto l23
					}
				}
			l21:
				v3 = v8 + v3
				t48 := int64(load64(m.memory[uint32(v3):]))
				t49 := v3
				v9 = t48
				store64(m.memory[uint32(t49):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
			}
		l23:
			{
				if uint32(v5) < uint32(i32(8)) {
					goto l24
				}
				t50 := int64(load64(m.memory[uint32(v8):]))
				store64(m.memory[uint32(v8+v5):], uint64(t50))
				goto l25
			}
		l24:
			if v5 == 0 {
				goto l25
			}
			memory_copy(m.memory, uint32(v8+i32(8)), uint32(v8), uint32(v5))
		l25:
			t51 := int64(load64(m.memory[int64(uint32(v1))+8:]))
			v16 = t51
			t52 := int64(load64(m.memory[uint32(v1):]))
			v18 = t52
			v6 = i32(0)
		l33:
			{
				t53 := v8
				v3 = v6
				v10 = t53 + v3
				t54 := int32(m.memory[uint32(v10)])
				if t54 != i32(128) {
					goto l26
				}
				v15 = v8 + (v3^i32(-1))*i32(36)
				v6 = v8 + (i32(0)-v3)*i32(36)
				v12 = v6 + i32(-28)
				v1 = v6 + i32(-32)
			l32:
				{
					t55 := int32(load32(m.memory[uint32(v1):]))
					t56 := int32(load32(m.memory[uint32(v12):]))
					t57 := m.fn64(v18, v16, t55, t56)
					t58 := v4
					v11 = int32(t57)
					v6 = t58 & v11
					v5 = v6
					{
						t59 := int64(load64(m.memory[uint32(v8+v6):]))
						v9 = t59 & i64(-0x7f7f7f7f7f7f7f80)
						if v9 != i64(0) {
							goto l27
						}
						v17 = i32(8)
						v5 = v6
					l28:
						{
							v5 = v5 + v17
							v17 = v17 + i32(8)
							t60 := v8
							v5 = v5 & v4
							t61 := int64(load64(m.memory[uint32(t60+v5):]))
							v9 = t61 & i64(-0x7f7f7f7f7f7f7f80)
							if v9 == 0 {
								goto l28
							}
						}
					}
				l27:
					{
						t62 := v8
						v5 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v9))))>>3) + v5) & v4
						t63 := int32(int8(m.memory[uint32(t62+v5)]))
						if t63 < i32(0) {
							goto l29
						}
						t64 := int64(load64(m.memory[uint32(v8):]))
						v5 = int32(uint32(int64(bits.TrailingZeros64(uint64(t64&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
					}
				l29:
					{
						if uint32((v5-v6^(v3-v6))&v4) < uint32(i32(8)) {
							goto l30
						}
						v6 = v8 + v5
						t65 := int32(m.memory[uint32(v6)])
						v17 = t65
						t66 := v6
						v11 = int32(uint32(v11) >> 25)
						m.memory[uint32(t66)] = byte(v11)
						m.memory[uint32(v8+(v5+i32(-8))&v4+i32(8))] = byte(v11)
						v6 = v8 + (v5^i32(-1))*i32(36)
						{
							if v17 != i32(255) {
								t72 := int32(load32(m.memory[uint32(v15):]))
								v5 = t72
								t73 := int32(load32(m.memory[uint32(v6):]))
								store32(m.memory[uint32(v15):], uint32(t73))
								store32(m.memory[uint32(v6):], uint32(v5))
								t74 := int32(load32(m.memory[int64(uint32(v6))+4:]))
								v5 = t74
								t75 := int32(load32(m.memory[int64(uint32(v15))+4:]))
								store32(m.memory[int64(uint32(v6))+4:], uint32(t75))
								store32(m.memory[int64(uint32(v15))+4:], uint32(v5))
								t76 := int32(load32(m.memory[int64(uint32(v15))+8:]))
								v5 = t76
								t77 := int32(load32(m.memory[int64(uint32(v6))+8:]))
								store32(m.memory[int64(uint32(v15))+8:], uint32(t77))
								store32(m.memory[int64(uint32(v6))+8:], uint32(v5))
								t78 := int32(load32(m.memory[int64(uint32(v6))+12:]))
								v5 = t78
								t79 := int32(load32(m.memory[int64(uint32(v15))+12:]))
								store32(m.memory[int64(uint32(v6))+12:], uint32(t79))
								store32(m.memory[int64(uint32(v15))+12:], uint32(v5))
								t80 := int32(load32(m.memory[int64(uint32(v15))+16:]))
								v5 = t80
								t81 := int32(load32(m.memory[int64(uint32(v6))+16:]))
								store32(m.memory[int64(uint32(v15))+16:], uint32(t81))
								store32(m.memory[int64(uint32(v6))+16:], uint32(v5))
								t82 := int32(load32(m.memory[int64(uint32(v6))+20:]))
								v5 = t82
								t83 := int32(load32(m.memory[int64(uint32(v15))+20:]))
								store32(m.memory[int64(uint32(v6))+20:], uint32(t83))
								store32(m.memory[int64(uint32(v15))+20:], uint32(v5))
								t84 := int32(load32(m.memory[int64(uint32(v15))+24:]))
								v5 = t84
								t85 := int32(load32(m.memory[int64(uint32(v6))+24:]))
								store32(m.memory[int64(uint32(v15))+24:], uint32(t85))
								store32(m.memory[int64(uint32(v6))+24:], uint32(v5))
								t86 := int32(load32(m.memory[int64(uint32(v6))+28:]))
								v5 = t86
								t87 := int32(load32(m.memory[int64(uint32(v15))+28:]))
								store32(m.memory[int64(uint32(v6))+28:], uint32(t87))
								store32(m.memory[int64(uint32(v15))+28:], uint32(v5))
								t88 := int32(load32(m.memory[int64(uint32(v15))+32:]))
								v5 = t88
								t89 := int32(load32(m.memory[int64(uint32(v6))+32:]))
								store32(m.memory[int64(uint32(v15))+32:], uint32(t89))
								store32(m.memory[int64(uint32(v6))+32:], uint32(v5))
								goto l32
							}
							m.memory[uint32(v10)] = byte(i32(255))
							m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(i32(255))
							t67 := int32(load32(m.memory[int64(uint32(v15))+32:]))
							store32(m.memory[int64(uint32(v6))+32:], uint32(t67))
							t68 := int64(load64(m.memory[int64(uint32(v15))+24:]))
							store64(m.memory[int64(uint32(v6))+24:], uint64(t68))
							t69 := int64(load64(m.memory[int64(uint32(v15))+16:]))
							store64(m.memory[int64(uint32(v6))+16:], uint64(t69))
							t70 := int64(load64(m.memory[int64(uint32(v15))+8:]))
							store64(m.memory[int64(uint32(v6))+8:], uint64(t70))
							t71 := int64(load64(m.memory[uint32(v15):]))
							store64(m.memory[uint32(v6):], uint64(t71))
							goto l26
						}
					}
				l30:
				}
				t90 := v10
				v6 = int32(uint32(v11) >> 25)
				m.memory[uint32(t90)] = byte(v6)
				m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(v6)
			}
		l26:
			v6 = v3 + i32(1)
			if v3 != v4 {
				goto l33
			}
			p91 := v7
			if uint32(v4) < uint32(i32(8)) {
				p91 = v4
			}
			v3 = p91
		}
	l20:
		store32(m.memory[int64(uint32(v0))+8:], uint32(v3-v2))
	l15:
		return i32(-1)
	}
l5:
	m.fn27(i32(1271632), i32(57), i32(1271660))
	panic("unreachable")
}
func (m *Module) fn72(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7, v8 int32
	var v9 int64
	var v10, v11, v12 int32
	var v13, v14 int64
	var v15, v16 int32
	var v17, v18 int64
	var v19, v20 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
		v2 = t0
		v3 = v2 + i32(1)
		if v3 == 0 {
			m.fn27(i32(1271632), i32(57), i32(1271660))
			panic("unreachable")
		}
		{
			t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t2 := v3
			v4 = t1
			t3 := v4
			v5 = v4 + i32(1)
			v6 = int32(uint32(v5) >> 3)
			v7 = v6 * i32(7)
			p4 := v7
			if uint32(v4) < uint32(i32(8)) {
				p4 = t3
			}
			v8 = p4
			if uint32(t2) <= uint32(int32(uint32(v8)>>1)) {
				goto l1
			}
			{
				{
					{
						v8 = v8 + i32(1)
						p5 := v3
						if uint32(v8) > uint32(v3) {
							p5 = v8
						}
						v3 = p5
						if uint32(v3) < uint32(i32(15)) {
							goto l2
						}
						if uint32(v3) > uint32(i32(0x1fffffff)) {
							m.fn27(i32(1271632), i32(57), i32(1271660))
							panic("unreachable")
						}
						t6 := int32(uint32(v3<<3) / uint32(i32(7)))
						v3 = i32_shr_u(i32(-1), int32(bits.LeadingZeros32(uint32(t6+i32(-1))))) + i32(1)
						goto l4
					}
				l2:
					p7 := v3&i32(8) + i32(8)
					if uint32(v3) < uint32(i32(4)) {
						p7 = i32(4)
					}
					v3 = p7
				}
			l4:
				v9 = int64(uint32(v3)) * i64(680)
				if int32(int64(uint64(v9)>>32)) != 0 {
					goto l5
				}
				v8 = v3 + i32(8)
				t8 := v8
				v10 = int32(v9)
				v6 = t8 + v10
				if uint32(v6) < uint32(v8) {
					goto l5
				}
				if uint32(v6) > uint32(i32(0x7ffffff8)) {
					goto l5
				}
				t9 := m.fn7(v6)
				v5 = t9
				if v5 != 0 {
					v6 = v5 + v10
					if v8 == 0 {
						goto l7
					}
					memory_fill(m.memory, uint32(v6), i32(255), uint32(v8))
				l7:
					v5 = v3 + i32(-1)
					p10 := int32(uint32(v3)>>3) * i32(7)
					if uint32(v3) < uint32(i32(9)) {
						p10 = v5
					}
					v11 = p10
					t11 := int32(load32(m.memory[uint32(v0):]))
					v12 = t11
					{
						if v2 == 0 {
							goto l8
						}
						t12 := int64(load64(m.memory[uint32(v12):]))
						v9 = (t12 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
						t13 := int64(load64(m.memory[int64(uint32(v1))+8:]))
						v13 = t13
						t14 := int64(load64(m.memory[uint32(v1):]))
						v14 = t14
						v8 = v12
						v1 = v2
						v3 = i32(0)
					l14:
						{
							if v9 != i64(0) {
								goto l9
							}
						l10:
							{
								v3 = v3 + i32(8)
								v8 = v8 + i32(8)
								t15 := int64(load64(m.memory[uint32(v8):]))
								v9 = t15 & i64(-0x7f7f7f7f7f7f7f80)
								if v9 == i64(-0x7f7f7f7f7f7f7f80) {
									goto l10
								}
							}
							v9 = v9 ^ i64(-0x7f7f7f7f7f7f7f80)
						l9:
							{
								t16 := v6
								t17 := v5
								t18 := v14
								t19 := v13
								t20 := v12
								v15 = int32(uint32(int64(bits.TrailingZeros64(uint64(v9))))>>3) + v3
								v10 = t20 + (i32(0)-v15)*i32(680)
								t21 := int32(load32(m.memory[uint32(v10+i32(-676)):]))
								t22 := int32(load32(m.memory[uint32(v10+i32(-672)):]))
								t23 := m.fn64(t18, t19, t21, t22)
								v16 = int32(t23)
								v10 = t17 & v16
								t24 := int64(load64(m.memory[uint32(t16+v10):]))
								v17 = t24 & i64(-0x7f7f7f7f7f7f7f80)
								if v17 != i64(0) {
									goto l11
								}
								v7 = i32(8)
							l12:
								{
									v10 = v10 + v7
									v7 = v7 + i32(8)
									t25 := v6
									v10 = v10 & v5
									t26 := int64(load64(m.memory[uint32(t25+v10):]))
									v17 = t26 & i64(-0x7f7f7f7f7f7f7f80)
									if v17 == 0 {
										goto l12
									}
								}
							}
						l11:
							v18 = v9 + i64(-1)
							{
								t27 := v6
								v10 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v17))))>>3) + v10) & v5
								t28 := int32(int8(m.memory[uint32(t27+v10)]))
								if t28 < i32(0) {
									goto l13
								}
								t29 := int64(load64(m.memory[uint32(v6):]))
								v10 = int32(uint32(int64(bits.TrailingZeros64(uint64(t29&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
							}
						l13:
							v9 = v18 & v9
							t30 := v6 + v10
							v16 = int32(uint32(v16) >> 25)
							m.memory[uint32(t30)] = byte(v16)
							m.memory[uint32(v6+(v10+i32(-8))&v5+i32(8))] = byte(v16)
							memory_copy(m.memory, uint32(v6+(v10^i32(-1))*i32(680)), uint32(v12+(v15^i32(-1))*i32(680)), uint32(i32(680)))
							v1 = v1 + i32(-1)
							if v1 != 0 {
								goto l14
							}
						}
					}
				l8:
					store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
					store32(m.memory[uint32(v0):], uint32(v6))
					store32(m.memory[int64(uint32(v0))+8:], uint32(v11-v2))
					if v4 == 0 {
						goto l15
					}
					t31 := v4
					v8 = (v4*i32(680) + i32(687)) & i32(-8)
					v3 = t31 + v8 + i32(9)
					if v3 == 0 {
						goto l15
					}
					v4 = v12 - v8
					t32 := int32(load32(m.memory[uint32(v4+i32(-4)):]))
					v8 = t32
					v6 = v8 & i32(-8)
					t33 := v6
					v8 = v8 & i32(3)
					p34 := i32(8)
					if v8 != 0 {
						p34 = i32(4)
					}
					if uint32(t33) < uint32(p34+v3) {
						m.fn3(i32(1274224), i32(46), i32(1274272))
						panic("unreachable")
					}
					if v8 == 0 {
						goto l17
					}
					if uint32(v6) > uint32(v3+i32(39)) {
						m.fn3(i32(1274288), i32(46), i32(1274336))
						panic("unreachable")
					}
				l17:
					m.fn1(v4)
					return i32(-1)
				}
				m.fn23(i32(8), v6)
				panic("unreachable")
			}
		}
	l1:
		{
			if v5 != 0 {
				goto l19
			}
			v3 = i32(0)
			goto l20
		l19:
			t35 := int32(load32(m.memory[uint32(v0):]))
			v8 = t35
			v3 = i32(0)
			{
				{
					t36 := v6
					var p37 int32
					if v5&i32(7) != i32(0) {
						p37 = 1
					}
					v6 = t36 + p37
					if v6 == i32(1) {
						goto l21
					}
					v12 = v6 & i32(1)
					v10 = v6 & i32(0x3ffffffe)
					v3 = i32(0)
				l22:
					{
						v6 = v8 + v3
						t38 := int64(load64(m.memory[uint32(v6):]))
						t39 := v6
						v9 = t38
						store64(m.memory[uint32(t39):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
						v6 = v6 + i32(8)
						t40 := int64(load64(m.memory[uint32(v6):]))
						t41 := v6
						v9 = t40
						store64(m.memory[uint32(t41):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
						v3 = v3 + i32(16)
						v10 = v10 + i32(-2)
						if v10 != 0 {
							goto l22
						}
					}
					if v12 == 0 {
						goto l23
					}
				}
			l21:
				v3 = v8 + v3
				t42 := int64(load64(m.memory[uint32(v3):]))
				t43 := v3
				v9 = t42
				store64(m.memory[uint32(t43):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
			}
		l23:
			{
				if uint32(v5) < uint32(i32(8)) {
					goto l24
				}
				t44 := int64(load64(m.memory[uint32(v8):]))
				store64(m.memory[uint32(v8+v5):], uint64(t44))
				goto l25
			}
		l24:
			if v5 == 0 {
				goto l25
			}
			memory_copy(m.memory, uint32(v8+i32(8)), uint32(v8), uint32(v5))
		l25:
			t45 := int64(load64(m.memory[int64(uint32(v1))+8:]))
			v17 = t45
			t46 := int64(load64(m.memory[uint32(v1):]))
			v18 = t46
			v6 = v8
			v10 = i32(0)
		l34:
			{
				t47 := v8
				v3 = v10
				v5 = t47 + v3
				t48 := int32(m.memory[uint32(v5)])
				if t48 != i32(128) {
					goto l26
				}
				v11 = v8 + (v3^i32(-1))*i32(680)
				v10 = v8 + (i32(0)-v3)*i32(680)
				v15 = v10 + i32(-672)
				v16 = v10 + i32(-676)
				{
				l33:
					{
						t49 := int32(load32(m.memory[uint32(v16):]))
						t50 := int32(load32(m.memory[uint32(v15):]))
						t51 := m.fn64(v18, v17, t49, t50)
						t52 := v4
						v1 = int32(t51)
						v10 = t52 & v1
						v12 = v10
						{
							t53 := int64(load64(m.memory[uint32(v8+v10):]))
							v9 = t53 & i64(-0x7f7f7f7f7f7f7f80)
							if v9 != i64(0) {
								goto l27
							}
							v19 = i32(8)
							v12 = v10
						l28:
							{
								v12 = v12 + v19
								v19 = v19 + i32(8)
								t54 := v8
								v12 = v12 & v4
								t55 := int64(load64(m.memory[uint32(t54+v12):]))
								v9 = t55 & i64(-0x7f7f7f7f7f7f7f80)
								if v9 == 0 {
									goto l28
								}
							}
						}
					l27:
						{
							t56 := v8
							v12 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v9))))>>3) + v12) & v4
							t57 := int32(int8(m.memory[uint32(t56+v12)]))
							if t57 < i32(0) {
								goto l29
							}
							t58 := int64(load64(m.memory[uint32(v8):]))
							v12 = int32(uint32(int64(bits.TrailingZeros64(uint64(t58&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
						}
					l29:
						{
							if uint32((v12-v10^(v3-v10))&v4) < uint32(i32(8)) {
								goto l30
							}
							v10 = v8 + v12
							t59 := int32(m.memory[uint32(v10)])
							v19 = t59
							t60 := v10
							v1 = int32(uint32(v1) >> 25)
							m.memory[uint32(t60)] = byte(v1)
							m.memory[uint32(v8+(v12+i32(-8))&v4+i32(8))] = byte(v1)
							if v19 == i32(255) {
								goto l31
							}
							v10 = i32(-680)
							v20 = v8 + v12*i32(-680)
						l32:
							{
								v12 = v20 + v10
								t61 := int32(load32(m.memory[uint32(v12):]))
								v19 = t61
								t62 := v12
								v1 = v6 + v10
								t63 := int32(load32(m.memory[uint32(v1):]))
								store32(m.memory[uint32(t62):], uint32(t63))
								store32(m.memory[uint32(v1):], uint32(v19))
								v1 = v1 + i32(4)
								t64 := int32(load32(m.memory[uint32(v1):]))
								v19 = t64
								t65 := v1
								v12 = v12 + i32(4)
								t66 := int32(load32(m.memory[uint32(v12):]))
								store32(m.memory[uint32(t65):], uint32(t66))
								store32(m.memory[uint32(v12):], uint32(v19))
								v10 = v10 + i32(8)
								if v10 != 0 {
									goto l32
								}
								goto l33
							}
						}
					l30:
					}
					t67 := v5
					v10 = int32(uint32(v1) >> 25)
					m.memory[uint32(t67)] = byte(v10)
					m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(v10)
					goto l26
				}
			l31:
				m.memory[uint32(v5)] = byte(i32(255))
				m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(i32(255))
				memory_copy(m.memory, uint32(v8+(v12^i32(-1))*i32(680)), uint32(v11), uint32(i32(680)))
			}
		l26:
			v10 = v3 + i32(1)
			v6 = v6 + i32(-680)
			if v3 != v4 {
				goto l34
			}
			p68 := v7
			if uint32(v4) < uint32(i32(8)) {
				p68 = v4
			}
			v3 = p68
		}
	l20:
		store32(m.memory[int64(uint32(v0))+8:], uint32(v3-v2))
	l15:
		return i32(-1)
	l5:
		m.fn27(i32(1271632), i32(57), i32(1271660))
		panic("unreachable")
	}
}
func (m *Module) fn73(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7, v8 int32
	var v9 int64
	var v10, v11, v12 int32
	var v13, v14 int64
	var v15 int32
	var v16 int64
	var v17 int32
	var v18 int64
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
		v2 = t0
		v3 = v2 + i32(1)
		if v3 == 0 {
			m.fn27(i32(1271632), i32(57), i32(1271660))
			panic("unreachable")
		}
		{
			t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t2 := v3
			v4 = t1
			t3 := v4
			v5 = v4 + i32(1)
			v6 = int32(uint32(v5) >> 3)
			v7 = v6 * i32(7)
			p4 := v7
			if uint32(v4) < uint32(i32(8)) {
				p4 = t3
			}
			v8 = p4
			if uint32(t2) <= uint32(int32(uint32(v8)>>1)) {
				goto l1
			}
			{
				{
					{
						v8 = v8 + i32(1)
						p5 := v3
						if uint32(v8) > uint32(v3) {
							p5 = v8
						}
						v3 = p5
						if uint32(v3) < uint32(i32(15)) {
							goto l2
						}
						if uint32(v3) > uint32(i32(0x1fffffff)) {
							m.fn27(i32(1271632), i32(57), i32(1271660))
							panic("unreachable")
						}
						t6 := int32(uint32(v3<<3) / uint32(i32(7)))
						v3 = i32_shr_u(i32(-1), int32(bits.LeadingZeros32(uint32(t6+i32(-1))))) + i32(1)
						goto l4
					}
				l2:
					p7 := v3&i32(8) + i32(8)
					if uint32(v3) < uint32(i32(4)) {
						p7 = i32(4)
					}
					v3 = p7
				}
			l4:
				v9 = int64(uint32(v3)) * i64(40)
				if int32(int64(uint64(v9)>>32)) != 0 {
					goto l5
				}
				v8 = v3 + i32(8)
				t8 := v8
				v10 = int32(v9)
				v6 = t8 + v10
				if uint32(v6) < uint32(v8) {
					goto l5
				}
				if uint32(v6) > uint32(i32(0x7ffffff8)) {
					goto l5
				}
				t9 := m.fn7(v6)
				v5 = t9
				if v5 != 0 {
					v6 = v5 + v10
					if v8 == 0 {
						goto l7
					}
					memory_fill(m.memory, uint32(v6), i32(255), uint32(v8))
				l7:
					v11 = v3 + i32(-1)
					p10 := int32(uint32(v3)>>3) * i32(7)
					if uint32(v3) < uint32(i32(9)) {
						p10 = v11
					}
					v7 = p10
					t11 := int32(load32(m.memory[uint32(v0):]))
					v12 = t11
					{
						if v2 == 0 {
							goto l8
						}
						t12 := int64(load64(m.memory[uint32(v12):]))
						v9 = (t12 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
						t13 := int64(load64(m.memory[int64(uint32(v1))+8:]))
						v13 = t13
						t14 := int64(load64(m.memory[uint32(v1):]))
						v14 = t14
						v8 = v12
						v1 = v2
						v3 = i32(0)
					l14:
						{
							if v9 != i64(0) {
								goto l9
							}
						l10:
							{
								v3 = v3 + i32(8)
								v8 = v8 + i32(8)
								t15 := int64(load64(m.memory[uint32(v8):]))
								v9 = t15 & i64(-0x7f7f7f7f7f7f7f80)
								if v9 == i64(-0x7f7f7f7f7f7f7f80) {
									goto l10
								}
							}
							v9 = v9 ^ i64(-0x7f7f7f7f7f7f7f80)
						l9:
							{
								t16 := v6
								t17 := v11
								t18 := v14
								t19 := v13
								t20 := v12
								v5 = int32(uint32(int64(bits.TrailingZeros64(uint64(v9))))>>3) + v3
								v10 = t20 + (i32(0)-v5)*i32(40)
								t21 := int32(load32(m.memory[uint32(v10+i32(-36)):]))
								t22 := int32(load32(m.memory[uint32(v10+i32(-32)):]))
								t23 := m.fn64(t18, t19, t21, t22)
								v15 = int32(t23)
								v10 = t17 & v15
								t24 := int64(load64(m.memory[uint32(t16+v10):]))
								v16 = t24 & i64(-0x7f7f7f7f7f7f7f80)
								if v16 != i64(0) {
									goto l11
								}
								v17 = i32(8)
							l12:
								{
									v10 = v10 + v17
									v17 = v17 + i32(8)
									t25 := v6
									v10 = v10 & v11
									t26 := int64(load64(m.memory[uint32(t25+v10):]))
									v16 = t26 & i64(-0x7f7f7f7f7f7f7f80)
									if v16 == 0 {
										goto l12
									}
								}
							}
						l11:
							v18 = v9 + i64(-1)
							{
								t27 := v6
								v10 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v16))))>>3) + v10) & v11
								t28 := int32(int8(m.memory[uint32(t27+v10)]))
								if t28 < i32(0) {
									goto l13
								}
								t29 := int64(load64(m.memory[uint32(v6):]))
								v10 = int32(uint32(int64(bits.TrailingZeros64(uint64(t29&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
							}
						l13:
							v9 = v18 & v9
							t30 := v6 + v10
							v15 = int32(uint32(v15) >> 25)
							m.memory[uint32(t30)] = byte(v15)
							m.memory[uint32(v6+(v10+i32(-8))&v11+i32(8))] = byte(v15)
							v10 = v6 + (v10^i32(-1))*i32(40)
							t31 := v10
							v5 = v12 + (v5^i32(-1))*i32(40)
							t32 := int64(load64(m.memory[int64(uint32(v5))+32:]))
							store64(m.memory[int64(uint32(t31))+32:], uint64(t32))
							t33 := int64(load64(m.memory[int64(uint32(v5))+24:]))
							store64(m.memory[int64(uint32(v10))+24:], uint64(t33))
							t34 := int64(load64(m.memory[int64(uint32(v5))+16:]))
							store64(m.memory[int64(uint32(v10))+16:], uint64(t34))
							t35 := int64(load64(m.memory[int64(uint32(v5))+8:]))
							store64(m.memory[int64(uint32(v10))+8:], uint64(t35))
							t36 := int64(load64(m.memory[uint32(v5):]))
							store64(m.memory[uint32(v10):], uint64(t36))
							v1 = v1 + i32(-1)
							if v1 != 0 {
								goto l14
							}
						}
					}
				l8:
					store32(m.memory[int64(uint32(v0))+4:], uint32(v11))
					store32(m.memory[uint32(v0):], uint32(v6))
					store32(m.memory[int64(uint32(v0))+8:], uint32(v7-v2))
					if v4 == 0 {
						goto l15
					}
					t37 := v4
					v8 = (v4*i32(40) + i32(47)) & i32(-8)
					v3 = t37 + v8 + i32(9)
					if v3 == 0 {
						goto l15
					}
					v4 = v12 - v8
					t38 := int32(load32(m.memory[uint32(v4+i32(-4)):]))
					v8 = t38
					v6 = v8 & i32(-8)
					t39 := v6
					v8 = v8 & i32(3)
					p40 := i32(8)
					if v8 != 0 {
						p40 = i32(4)
					}
					if uint32(t39) < uint32(p40+v3) {
						m.fn3(i32(1274224), i32(46), i32(1274272))
						panic("unreachable")
					}
					if v8 == 0 {
						goto l17
					}
					if uint32(v6) > uint32(v3+i32(39)) {
						m.fn3(i32(1274288), i32(46), i32(1274336))
						panic("unreachable")
					}
				l17:
					m.fn1(v4)
					return i32(-1)
				}
				m.fn23(i32(8), v6)
				panic("unreachable")
			}
		}
	l1:
		{
			if v5 != 0 {
				goto l19
			}
			v3 = i32(0)
			goto l20
		l19:
			t41 := int32(load32(m.memory[uint32(v0):]))
			v8 = t41
			v3 = i32(0)
			{
				{
					t42 := v6
					var p43 int32
					if v5&i32(7) != i32(0) {
						p43 = 1
					}
					v6 = t42 + p43
					if v6 == i32(1) {
						goto l21
					}
					v11 = v6 & i32(1)
					v10 = v6 & i32(0x3ffffffe)
					v3 = i32(0)
				l22:
					{
						v6 = v8 + v3
						t44 := int64(load64(m.memory[uint32(v6):]))
						t45 := v6
						v9 = t44
						store64(m.memory[uint32(t45):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
						v6 = v6 + i32(8)
						t46 := int64(load64(m.memory[uint32(v6):]))
						t47 := v6
						v9 = t46
						store64(m.memory[uint32(t47):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
						v3 = v3 + i32(16)
						v10 = v10 + i32(-2)
						if v10 != 0 {
							goto l22
						}
					}
					if v11 == 0 {
						goto l23
					}
				}
			l21:
				v3 = v8 + v3
				t48 := int64(load64(m.memory[uint32(v3):]))
				t49 := v3
				v9 = t48
				store64(m.memory[uint32(t49):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
			}
		l23:
			{
				if uint32(v5) < uint32(i32(8)) {
					goto l24
				}
				t50 := int64(load64(m.memory[uint32(v8):]))
				store64(m.memory[uint32(v8+v5):], uint64(t50))
				goto l25
			}
		l24:
			if v5 == 0 {
				goto l25
			}
			memory_copy(m.memory, uint32(v8+i32(8)), uint32(v8), uint32(v5))
		l25:
			t51 := int64(load64(m.memory[int64(uint32(v1))+8:]))
			v16 = t51
			t52 := int64(load64(m.memory[uint32(v1):]))
			v18 = t52
			v6 = i32(0)
		l33:
			{
				t53 := v8
				v3 = v6
				v10 = t53 + v3
				t54 := int32(m.memory[uint32(v10)])
				if t54 != i32(128) {
					goto l26
				}
				v15 = v8 + (v3^i32(-1))*i32(40)
				v6 = v8 + (i32(0)-v3)*i32(40)
				v12 = v6 + i32(-32)
				v1 = v6 + i32(-36)
			l32:
				{
					t55 := int32(load32(m.memory[uint32(v1):]))
					t56 := int32(load32(m.memory[uint32(v12):]))
					t57 := m.fn64(v18, v16, t55, t56)
					t58 := v4
					v11 = int32(t57)
					v6 = t58 & v11
					v5 = v6
					{
						t59 := int64(load64(m.memory[uint32(v8+v6):]))
						v9 = t59 & i64(-0x7f7f7f7f7f7f7f80)
						if v9 != i64(0) {
							goto l27
						}
						v17 = i32(8)
						v5 = v6
					l28:
						{
							v5 = v5 + v17
							v17 = v17 + i32(8)
							t60 := v8
							v5 = v5 & v4
							t61 := int64(load64(m.memory[uint32(t60+v5):]))
							v9 = t61 & i64(-0x7f7f7f7f7f7f7f80)
							if v9 == 0 {
								goto l28
							}
						}
					}
				l27:
					{
						t62 := v8
						v5 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v9))))>>3) + v5) & v4
						t63 := int32(int8(m.memory[uint32(t62+v5)]))
						if t63 < i32(0) {
							goto l29
						}
						t64 := int64(load64(m.memory[uint32(v8):]))
						v5 = int32(uint32(int64(bits.TrailingZeros64(uint64(t64&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
					}
				l29:
					{
						if uint32((v5-v6^(v3-v6))&v4) < uint32(i32(8)) {
							goto l30
						}
						v6 = v8 + v5
						t65 := int32(m.memory[uint32(v6)])
						v17 = t65
						t66 := v6
						v11 = int32(uint32(v11) >> 25)
						m.memory[uint32(t66)] = byte(v11)
						m.memory[uint32(v8+(v5+i32(-8))&v4+i32(8))] = byte(v11)
						v6 = v8 + (v5^i32(-1))*i32(40)
						{
							if v17 != i32(255) {
								t72 := int64(load64(m.memory[uint32(v15):]))
								v9 = t72
								t73 := int64(load64(m.memory[uint32(v6):]))
								store64(m.memory[uint32(v15):], uint64(t73))
								store64(m.memory[uint32(v6):], uint64(v9))
								t74 := int64(load64(m.memory[int64(uint32(v6))+8:]))
								v9 = t74
								t75 := int64(load64(m.memory[int64(uint32(v15))+8:]))
								store64(m.memory[int64(uint32(v6))+8:], uint64(t75))
								store64(m.memory[int64(uint32(v15))+8:], uint64(v9))
								t76 := int64(load64(m.memory[int64(uint32(v15))+16:]))
								v9 = t76
								t77 := int32(load32(m.memory[int64(uint32(v6))+20:]))
								store32(m.memory[int64(uint32(v15))+20:], uint32(t77))
								t78 := int32(load32(m.memory[int64(uint32(v6))+16:]))
								v5 = t78
								store64(m.memory[int64(uint32(v6))+16:], uint64(v9))
								store32(m.memory[int64(uint32(v15))+16:], uint32(v5))
								t79 := int32(load32(m.memory[int64(uint32(v6))+28:]))
								v5 = t79
								t80 := int32(load32(m.memory[int64(uint32(v15))+28:]))
								store32(m.memory[int64(uint32(v6))+28:], uint32(t80))
								store32(m.memory[int64(uint32(v15))+28:], uint32(v5))
								t81 := int32(load32(m.memory[int64(uint32(v15))+24:]))
								v5 = t81
								t82 := int32(load32(m.memory[int64(uint32(v6))+24:]))
								store32(m.memory[int64(uint32(v15))+24:], uint32(t82))
								store32(m.memory[int64(uint32(v6))+24:], uint32(v5))
								t83 := int32(load32(m.memory[int64(uint32(v15))+36:]))
								v5 = t83
								t84 := int32(load32(m.memory[int64(uint32(v6))+36:]))
								store32(m.memory[int64(uint32(v15))+36:], uint32(t84))
								t85 := int32(load32(m.memory[int64(uint32(v6))+32:]))
								v11 = t85
								t86 := int32(load32(m.memory[int64(uint32(v15))+32:]))
								store32(m.memory[int64(uint32(v6))+32:], uint32(t86))
								store32(m.memory[int64(uint32(v15))+32:], uint32(v11))
								store32(m.memory[int64(uint32(v6))+36:], uint32(v5))
								goto l32
							}
							m.memory[uint32(v10)] = byte(i32(255))
							m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(i32(255))
							t67 := int64(load64(m.memory[int64(uint32(v15))+32:]))
							store64(m.memory[int64(uint32(v6))+32:], uint64(t67))
							t68 := int64(load64(m.memory[int64(uint32(v15))+24:]))
							store64(m.memory[int64(uint32(v6))+24:], uint64(t68))
							t69 := int64(load64(m.memory[int64(uint32(v15))+16:]))
							store64(m.memory[int64(uint32(v6))+16:], uint64(t69))
							t70 := int64(load64(m.memory[int64(uint32(v15))+8:]))
							store64(m.memory[int64(uint32(v6))+8:], uint64(t70))
							t71 := int64(load64(m.memory[uint32(v15):]))
							store64(m.memory[uint32(v6):], uint64(t71))
							goto l26
						}
					}
				l30:
				}
				t87 := v10
				v6 = int32(uint32(v11) >> 25)
				m.memory[uint32(t87)] = byte(v6)
				m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(v6)
			}
		l26:
			v6 = v3 + i32(1)
			if v3 != v4 {
				goto l33
			}
			p88 := v7
			if uint32(v4) < uint32(i32(8)) {
				p88 = v4
			}
			v3 = p88
		}
	l20:
		store32(m.memory[int64(uint32(v0))+8:], uint32(v3-v2))
	l15:
		return i32(-1)
	}
l5:
	m.fn27(i32(1271632), i32(57), i32(1271660))
	panic("unreachable")
}
func (m *Module) fn74(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7, v8 int32
	var v9 int64
	var v10, v11, v12 int32
	var v13, v14 int64
	var v15 int32
	var v16 int64
	var v17 int32
	var v18 int64
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
		v2 = t0
		v3 = v2 + i32(1)
		if v3 == 0 {
			m.fn27(i32(1271632), i32(57), i32(1271660))
			panic("unreachable")
		}
		{
			t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t2 := v3
			v4 = t1
			t3 := v4
			v5 = v4 + i32(1)
			v6 = int32(uint32(v5) >> 3)
			v7 = v6 * i32(7)
			p4 := v7
			if uint32(v4) < uint32(i32(8)) {
				p4 = t3
			}
			v8 = p4
			if uint32(t2) <= uint32(int32(uint32(v8)>>1)) {
				goto l1
			}
			{
				{
					{
						v8 = v8 + i32(1)
						p5 := v3
						if uint32(v8) > uint32(v3) {
							p5 = v8
						}
						v3 = p5
						if uint32(v3) < uint32(i32(15)) {
							goto l2
						}
						if uint32(v3) > uint32(i32(0x1fffffff)) {
							m.fn27(i32(1271632), i32(57), i32(1271660))
							panic("unreachable")
						}
						t6 := int32(uint32(v3<<3) / uint32(i32(7)))
						v3 = i32_shr_u(i32(-1), int32(bits.LeadingZeros32(uint32(t6+i32(-1))))) + i32(1)
						goto l4
					}
				l2:
					p7 := v3&i32(8) + i32(8)
					if uint32(v3) < uint32(i32(4)) {
						p7 = i32(4)
					}
					v3 = p7
				}
			l4:
				v9 = int64(uint32(v3)) * i64(28)
				if int32(int64(uint64(v9)>>32)) != 0 {
					goto l5
				}
				v6 = int32(v9)
				if uint32(v6) > uint32(i32(-8)) {
					goto l5
				}
				v8 = v3 + i32(8)
				t8 := v8
				v10 = (v6 + i32(7)) & i32(-8)
				v6 = t8 + v10
				if uint32(v6) < uint32(v8) {
					goto l5
				}
				if uint32(v6) > uint32(i32(0x7ffffff8)) {
					goto l5
				}
				t9 := m.fn7(v6)
				v5 = t9
				if v5 != 0 {
					v6 = v5 + v10
					if v8 == 0 {
						goto l7
					}
					memory_fill(m.memory, uint32(v6), i32(255), uint32(v8))
				l7:
					v11 = v3 + i32(-1)
					p10 := int32(uint32(v3)>>3) * i32(7)
					if uint32(v3) < uint32(i32(9)) {
						p10 = v11
					}
					v7 = p10
					t11 := int32(load32(m.memory[uint32(v0):]))
					v12 = t11
					{
						if v2 == 0 {
							goto l8
						}
						t12 := int64(load64(m.memory[uint32(v12):]))
						v9 = (t12 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
						t13 := int64(load64(m.memory[int64(uint32(v1))+8:]))
						v13 = t13
						t14 := int64(load64(m.memory[uint32(v1):]))
						v14 = t14
						v8 = v12
						v1 = v2
						v3 = i32(0)
					l14:
						{
							if v9 != i64(0) {
								goto l9
							}
						l10:
							{
								v3 = v3 + i32(8)
								v8 = v8 + i32(8)
								t15 := int64(load64(m.memory[uint32(v8):]))
								v9 = t15 & i64(-0x7f7f7f7f7f7f7f80)
								if v9 == i64(-0x7f7f7f7f7f7f7f80) {
									goto l10
								}
							}
							v9 = v9 ^ i64(-0x7f7f7f7f7f7f7f80)
						l9:
							{
								t16 := v6
								t17 := v11
								t18 := v14
								t19 := v13
								t20 := v12
								v5 = int32(uint32(int64(bits.TrailingZeros64(uint64(v9))))>>3) + v3
								v10 = t20 + (i32(0)-v5)*i32(28)
								t21 := int32(load32(m.memory[uint32(v10+i32(-24)):]))
								t22 := int32(load32(m.memory[uint32(v10+i32(-20)):]))
								t23 := m.fn64(t18, t19, t21, t22)
								v15 = int32(t23)
								v10 = t17 & v15
								t24 := int64(load64(m.memory[uint32(t16+v10):]))
								v16 = t24 & i64(-0x7f7f7f7f7f7f7f80)
								if v16 != i64(0) {
									goto l11
								}
								v17 = i32(8)
							l12:
								{
									v10 = v10 + v17
									v17 = v17 + i32(8)
									t25 := v6
									v10 = v10 & v11
									t26 := int64(load64(m.memory[uint32(t25+v10):]))
									v16 = t26 & i64(-0x7f7f7f7f7f7f7f80)
									if v16 == 0 {
										goto l12
									}
								}
							}
						l11:
							v18 = v9 + i64(-1)
							{
								t27 := v6
								v10 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v16))))>>3) + v10) & v11
								t28 := int32(int8(m.memory[uint32(t27+v10)]))
								if t28 < i32(0) {
									goto l13
								}
								t29 := int64(load64(m.memory[uint32(v6):]))
								v10 = int32(uint32(int64(bits.TrailingZeros64(uint64(t29&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
							}
						l13:
							v9 = v18 & v9
							t30 := v6 + v10
							v15 = int32(uint32(v15) >> 25)
							m.memory[uint32(t30)] = byte(v15)
							m.memory[uint32(v6+(v10+i32(-8))&v11+i32(8))] = byte(v15)
							v10 = v6 + (v10^i32(-1))*i32(28)
							t31 := v10
							v5 = v12 + (v5^i32(-1))*i32(28)
							t32 := int32(load32(m.memory[int64(uint32(v5))+24:]))
							store32(m.memory[int64(uint32(t31))+24:], uint32(t32))
							t33 := int64(load64(m.memory[int64(uint32(v5))+16:]))
							store64(m.memory[int64(uint32(v10))+16:], uint64(t33))
							t34 := int64(load64(m.memory[int64(uint32(v5))+8:]))
							store64(m.memory[int64(uint32(v10))+8:], uint64(t34))
							t35 := int64(load64(m.memory[uint32(v5):]))
							store64(m.memory[uint32(v10):], uint64(t35))
							v1 = v1 + i32(-1)
							if v1 != 0 {
								goto l14
							}
						}
					}
				l8:
					store32(m.memory[int64(uint32(v0))+4:], uint32(v11))
					store32(m.memory[uint32(v0):], uint32(v6))
					store32(m.memory[int64(uint32(v0))+8:], uint32(v7-v2))
					if v4 == 0 {
						goto l15
					}
					t36 := v4
					v8 = (v4*i32(28) + i32(35)) & i32(-8)
					v3 = t36 + v8 + i32(9)
					if v3 == 0 {
						goto l15
					}
					v4 = v12 - v8
					t37 := int32(load32(m.memory[uint32(v4+i32(-4)):]))
					v8 = t37
					v6 = v8 & i32(-8)
					t38 := v6
					v8 = v8 & i32(3)
					p39 := i32(8)
					if v8 != 0 {
						p39 = i32(4)
					}
					if uint32(t38) < uint32(p39+v3) {
						m.fn3(i32(1274224), i32(46), i32(1274272))
						panic("unreachable")
					}
					if v8 == 0 {
						goto l17
					}
					if uint32(v6) > uint32(v3+i32(39)) {
						m.fn3(i32(1274288), i32(46), i32(1274336))
						panic("unreachable")
					}
				l17:
					m.fn1(v4)
					return i32(-1)
				}
				m.fn23(i32(8), v6)
				panic("unreachable")
			}
		}
	l1:
		{
			if v5 != 0 {
				goto l19
			}
			v3 = i32(0)
			goto l20
		l19:
			t40 := int32(load32(m.memory[uint32(v0):]))
			v8 = t40
			v3 = i32(0)
			{
				{
					t41 := v6
					var p42 int32
					if v5&i32(7) != i32(0) {
						p42 = 1
					}
					v6 = t41 + p42
					if v6 == i32(1) {
						goto l21
					}
					v11 = v6 & i32(1)
					v10 = v6 & i32(0x3ffffffe)
					v3 = i32(0)
				l22:
					{
						v6 = v8 + v3
						t43 := int64(load64(m.memory[uint32(v6):]))
						t44 := v6
						v9 = t43
						store64(m.memory[uint32(t44):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
						v6 = v6 + i32(8)
						t45 := int64(load64(m.memory[uint32(v6):]))
						t46 := v6
						v9 = t45
						store64(m.memory[uint32(t46):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
						v3 = v3 + i32(16)
						v10 = v10 + i32(-2)
						if v10 != 0 {
							goto l22
						}
					}
					if v11 == 0 {
						goto l23
					}
				}
			l21:
				v3 = v8 + v3
				t47 := int64(load64(m.memory[uint32(v3):]))
				t48 := v3
				v9 = t47
				store64(m.memory[uint32(t48):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
			}
		l23:
			{
				if uint32(v5) < uint32(i32(8)) {
					goto l24
				}
				t49 := int64(load64(m.memory[uint32(v8):]))
				store64(m.memory[uint32(v8+v5):], uint64(t49))
				goto l25
			}
		l24:
			if v5 == 0 {
				goto l25
			}
			memory_copy(m.memory, uint32(v8+i32(8)), uint32(v8), uint32(v5))
		l25:
			t50 := int64(load64(m.memory[int64(uint32(v1))+8:]))
			v16 = t50
			t51 := int64(load64(m.memory[uint32(v1):]))
			v18 = t51
			v6 = i32(0)
		l33:
			{
				t52 := v8
				v3 = v6
				v10 = t52 + v3
				t53 := int32(m.memory[uint32(v10)])
				if t53 != i32(128) {
					goto l26
				}
				v15 = v8 + (v3^i32(-1))*i32(28)
				v6 = v8 + (i32(0)-v3)*i32(28)
				v12 = v6 + i32(-20)
				v1 = v6 + i32(-24)
			l32:
				{
					t54 := int32(load32(m.memory[uint32(v1):]))
					t55 := int32(load32(m.memory[uint32(v12):]))
					t56 := m.fn64(v18, v16, t54, t55)
					t57 := v4
					v11 = int32(t56)
					v6 = t57 & v11
					v5 = v6
					{
						t58 := int64(load64(m.memory[uint32(v8+v6):]))
						v9 = t58 & i64(-0x7f7f7f7f7f7f7f80)
						if v9 != i64(0) {
							goto l27
						}
						v17 = i32(8)
						v5 = v6
					l28:
						{
							v5 = v5 + v17
							v17 = v17 + i32(8)
							t59 := v8
							v5 = v5 & v4
							t60 := int64(load64(m.memory[uint32(t59+v5):]))
							v9 = t60 & i64(-0x7f7f7f7f7f7f7f80)
							if v9 == 0 {
								goto l28
							}
						}
					}
				l27:
					{
						t61 := v8
						v5 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v9))))>>3) + v5) & v4
						t62 := int32(int8(m.memory[uint32(t61+v5)]))
						if t62 < i32(0) {
							goto l29
						}
						t63 := int64(load64(m.memory[uint32(v8):]))
						v5 = int32(uint32(int64(bits.TrailingZeros64(uint64(t63&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
					}
				l29:
					{
						if uint32((v5-v6^(v3-v6))&v4) < uint32(i32(8)) {
							goto l30
						}
						v6 = v8 + v5
						t64 := int32(m.memory[uint32(v6)])
						v17 = t64
						t65 := v6
						v11 = int32(uint32(v11) >> 25)
						m.memory[uint32(t65)] = byte(v11)
						m.memory[uint32(v8+(v5+i32(-8))&v4+i32(8))] = byte(v11)
						v6 = v8 + (v5^i32(-1))*i32(28)
						{
							if v17 != i32(255) {
								t70 := int32(load32(m.memory[uint32(v6):]))
								v5 = t70
								t71 := int32(load32(m.memory[uint32(v15):]))
								store32(m.memory[uint32(v6):], uint32(t71))
								store32(m.memory[uint32(v15):], uint32(v5))
								t72 := int32(load32(m.memory[int64(uint32(v15))+4:]))
								v5 = t72
								t73 := int32(load32(m.memory[int64(uint32(v6))+4:]))
								store32(m.memory[int64(uint32(v15))+4:], uint32(t73))
								store32(m.memory[int64(uint32(v6))+4:], uint32(v5))
								t74 := int32(load32(m.memory[int64(uint32(v6))+8:]))
								v5 = t74
								t75 := int32(load32(m.memory[int64(uint32(v15))+8:]))
								store32(m.memory[int64(uint32(v6))+8:], uint32(t75))
								store32(m.memory[int64(uint32(v15))+8:], uint32(v5))
								t76 := int32(load32(m.memory[int64(uint32(v15))+12:]))
								v5 = t76
								t77 := int32(load32(m.memory[int64(uint32(v6))+12:]))
								store32(m.memory[int64(uint32(v15))+12:], uint32(t77))
								store32(m.memory[int64(uint32(v6))+12:], uint32(v5))
								t78 := int32(load32(m.memory[int64(uint32(v6))+16:]))
								v5 = t78
								t79 := int32(load32(m.memory[int64(uint32(v15))+16:]))
								store32(m.memory[int64(uint32(v6))+16:], uint32(t79))
								store32(m.memory[int64(uint32(v15))+16:], uint32(v5))
								t80 := int32(load32(m.memory[int64(uint32(v15))+20:]))
								v5 = t80
								t81 := int32(load32(m.memory[int64(uint32(v6))+20:]))
								store32(m.memory[int64(uint32(v15))+20:], uint32(t81))
								store32(m.memory[int64(uint32(v6))+20:], uint32(v5))
								t82 := int32(load32(m.memory[int64(uint32(v6))+24:]))
								v5 = t82
								t83 := int32(load32(m.memory[int64(uint32(v15))+24:]))
								store32(m.memory[int64(uint32(v6))+24:], uint32(t83))
								store32(m.memory[int64(uint32(v15))+24:], uint32(v5))
								goto l32
							}
							m.memory[uint32(v10)] = byte(i32(255))
							m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(i32(255))
							t66 := int32(load32(m.memory[int64(uint32(v15))+24:]))
							store32(m.memory[int64(uint32(v6))+24:], uint32(t66))
							t67 := int64(load64(m.memory[int64(uint32(v15))+16:]))
							store64(m.memory[int64(uint32(v6))+16:], uint64(t67))
							t68 := int64(load64(m.memory[int64(uint32(v15))+8:]))
							store64(m.memory[int64(uint32(v6))+8:], uint64(t68))
							t69 := int64(load64(m.memory[uint32(v15):]))
							store64(m.memory[uint32(v6):], uint64(t69))
							goto l26
						}
					}
				l30:
				}
				t84 := v10
				v6 = int32(uint32(v11) >> 25)
				m.memory[uint32(t84)] = byte(v6)
				m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(v6)
			}
		l26:
			v6 = v3 + i32(1)
			if v3 != v4 {
				goto l33
			}
			p85 := v7
			if uint32(v4) < uint32(i32(8)) {
				p85 = v4
			}
			v3 = p85
		}
	l20:
		store32(m.memory[int64(uint32(v0))+8:], uint32(v3-v2))
	l15:
		return i32(-1)
	}
l5:
	m.fn27(i32(1271632), i32(57), i32(1271660))
	panic("unreachable")
}
func (m *Module) fn75(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7, v8 int32
	var v9 int64
	var v10, v11, v12 int32
	var v13, v14 int64
	var v15 int32
	var v16 int64
	var v17 int32
	var v18 int64
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
		v2 = t0
		v3 = v2 + i32(1)
		if v3 == 0 {
			m.fn27(i32(1271632), i32(57), i32(1271660))
			panic("unreachable")
		}
		{
			t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t2 := v3
			v4 = t1
			t3 := v4
			v5 = v4 + i32(1)
			v6 = int32(uint32(v5) >> 3)
			v7 = v6 * i32(7)
			p4 := v7
			if uint32(v4) < uint32(i32(8)) {
				p4 = t3
			}
			v8 = p4
			if uint32(t2) <= uint32(int32(uint32(v8)>>1)) {
				goto l1
			}
			{
				{
					{
						v8 = v8 + i32(1)
						p5 := v3
						if uint32(v8) > uint32(v3) {
							p5 = v8
						}
						v3 = p5
						if uint32(v3) < uint32(i32(15)) {
							goto l2
						}
						if uint32(v3) > uint32(i32(0x1fffffff)) {
							m.fn27(i32(1271632), i32(57), i32(1271660))
							panic("unreachable")
						}
						t6 := int32(uint32(v3<<3) / uint32(i32(7)))
						v3 = i32_shr_u(i32(-1), int32(bits.LeadingZeros32(uint32(t6+i32(-1))))) + i32(1)
						goto l4
					}
				l2:
					p7 := v3&i32(8) + i32(8)
					if uint32(v3) < uint32(i32(4)) {
						p7 = i32(4)
					}
					v3 = p7
				}
			l4:
				v9 = int64(uint32(v3)) * i64(36)
				if int32(int64(uint64(v9)>>32)) != 0 {
					goto l5
				}
				v6 = int32(v9)
				if uint32(v6) > uint32(i32(-8)) {
					goto l5
				}
				v8 = v3 + i32(8)
				t8 := v8
				v10 = (v6 + i32(7)) & i32(-8)
				v6 = t8 + v10
				if uint32(v6) < uint32(v8) {
					goto l5
				}
				if uint32(v6) > uint32(i32(0x7ffffff8)) {
					goto l5
				}
				t9 := m.fn7(v6)
				v5 = t9
				if v5 != 0 {
					v6 = v5 + v10
					if v8 == 0 {
						goto l7
					}
					memory_fill(m.memory, uint32(v6), i32(255), uint32(v8))
				l7:
					v11 = v3 + i32(-1)
					p10 := int32(uint32(v3)>>3) * i32(7)
					if uint32(v3) < uint32(i32(9)) {
						p10 = v11
					}
					v7 = p10
					t11 := int32(load32(m.memory[uint32(v0):]))
					v12 = t11
					{
						if v2 == 0 {
							goto l8
						}
						t12 := int64(load64(m.memory[uint32(v12):]))
						v9 = (t12 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
						t13 := int64(load64(m.memory[int64(uint32(v1))+8:]))
						v13 = t13
						t14 := int64(load64(m.memory[uint32(v1):]))
						v14 = t14
						v8 = v12
						v1 = v2
						v3 = i32(0)
					l14:
						{
							if v9 != i64(0) {
								goto l9
							}
						l10:
							{
								v3 = v3 + i32(8)
								v8 = v8 + i32(8)
								t15 := int64(load64(m.memory[uint32(v8):]))
								v9 = t15 & i64(-0x7f7f7f7f7f7f7f80)
								if v9 == i64(-0x7f7f7f7f7f7f7f80) {
									goto l10
								}
							}
							v9 = v9 ^ i64(-0x7f7f7f7f7f7f7f80)
						l9:
							{
								t16 := v6
								t17 := v11
								t18 := v14
								t19 := v13
								t20 := v12
								v5 = int32(uint32(int64(bits.TrailingZeros64(uint64(v9))))>>3) + v3
								v10 = t20 + (i32(0)-v5)*i32(36)
								t21 := int32(load32(m.memory[uint32(v10+i32(-32)):]))
								t22 := int32(load32(m.memory[uint32(v10+i32(-28)):]))
								t23 := m.fn64(t18, t19, t21, t22)
								v15 = int32(t23)
								v10 = t17 & v15
								t24 := int64(load64(m.memory[uint32(t16+v10):]))
								v16 = t24 & i64(-0x7f7f7f7f7f7f7f80)
								if v16 != i64(0) {
									goto l11
								}
								v17 = i32(8)
							l12:
								{
									v10 = v10 + v17
									v17 = v17 + i32(8)
									t25 := v6
									v10 = v10 & v11
									t26 := int64(load64(m.memory[uint32(t25+v10):]))
									v16 = t26 & i64(-0x7f7f7f7f7f7f7f80)
									if v16 == 0 {
										goto l12
									}
								}
							}
						l11:
							v18 = v9 + i64(-1)
							{
								t27 := v6
								v10 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v16))))>>3) + v10) & v11
								t28 := int32(int8(m.memory[uint32(t27+v10)]))
								if t28 < i32(0) {
									goto l13
								}
								t29 := int64(load64(m.memory[uint32(v6):]))
								v10 = int32(uint32(int64(bits.TrailingZeros64(uint64(t29&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
							}
						l13:
							v9 = v18 & v9
							t30 := v6 + v10
							v15 = int32(uint32(v15) >> 25)
							m.memory[uint32(t30)] = byte(v15)
							m.memory[uint32(v6+(v10+i32(-8))&v11+i32(8))] = byte(v15)
							v10 = v6 + (v10^i32(-1))*i32(36)
							t31 := v10
							v5 = v12 + (v5^i32(-1))*i32(36)
							t32 := int32(load32(m.memory[int64(uint32(v5))+32:]))
							store32(m.memory[int64(uint32(t31))+32:], uint32(t32))
							t33 := int64(load64(m.memory[int64(uint32(v5))+24:]))
							store64(m.memory[int64(uint32(v10))+24:], uint64(t33))
							t34 := int64(load64(m.memory[int64(uint32(v5))+16:]))
							store64(m.memory[int64(uint32(v10))+16:], uint64(t34))
							t35 := int64(load64(m.memory[int64(uint32(v5))+8:]))
							store64(m.memory[int64(uint32(v10))+8:], uint64(t35))
							t36 := int64(load64(m.memory[uint32(v5):]))
							store64(m.memory[uint32(v10):], uint64(t36))
							v1 = v1 + i32(-1)
							if v1 != 0 {
								goto l14
							}
						}
					}
				l8:
					store32(m.memory[int64(uint32(v0))+4:], uint32(v11))
					store32(m.memory[uint32(v0):], uint32(v6))
					store32(m.memory[int64(uint32(v0))+8:], uint32(v7-v2))
					if v4 == 0 {
						goto l15
					}
					t37 := v4
					v8 = (v4*i32(36) + i32(43)) & i32(-8)
					v3 = t37 + v8 + i32(9)
					if v3 == 0 {
						goto l15
					}
					v4 = v12 - v8
					t38 := int32(load32(m.memory[uint32(v4+i32(-4)):]))
					v8 = t38
					v6 = v8 & i32(-8)
					t39 := v6
					v8 = v8 & i32(3)
					p40 := i32(8)
					if v8 != 0 {
						p40 = i32(4)
					}
					if uint32(t39) < uint32(p40+v3) {
						m.fn3(i32(1274224), i32(46), i32(1274272))
						panic("unreachable")
					}
					if v8 == 0 {
						goto l17
					}
					if uint32(v6) > uint32(v3+i32(39)) {
						m.fn3(i32(1274288), i32(46), i32(1274336))
						panic("unreachable")
					}
				l17:
					m.fn1(v4)
					return i32(-1)
				}
				m.fn23(i32(8), v6)
				panic("unreachable")
			}
		}
	l1:
		{
			if v5 != 0 {
				goto l19
			}
			v3 = i32(0)
			goto l20
		l19:
			t41 := int32(load32(m.memory[uint32(v0):]))
			v8 = t41
			v3 = i32(0)
			{
				{
					t42 := v6
					var p43 int32
					if v5&i32(7) != i32(0) {
						p43 = 1
					}
					v6 = t42 + p43
					if v6 == i32(1) {
						goto l21
					}
					v11 = v6 & i32(1)
					v10 = v6 & i32(0x3ffffffe)
					v3 = i32(0)
				l22:
					{
						v6 = v8 + v3
						t44 := int64(load64(m.memory[uint32(v6):]))
						t45 := v6
						v9 = t44
						store64(m.memory[uint32(t45):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
						v6 = v6 + i32(8)
						t46 := int64(load64(m.memory[uint32(v6):]))
						t47 := v6
						v9 = t46
						store64(m.memory[uint32(t47):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
						v3 = v3 + i32(16)
						v10 = v10 + i32(-2)
						if v10 != 0 {
							goto l22
						}
					}
					if v11 == 0 {
						goto l23
					}
				}
			l21:
				v3 = v8 + v3
				t48 := int64(load64(m.memory[uint32(v3):]))
				t49 := v3
				v9 = t48
				store64(m.memory[uint32(t49):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
			}
		l23:
			{
				if uint32(v5) < uint32(i32(8)) {
					goto l24
				}
				t50 := int64(load64(m.memory[uint32(v8):]))
				store64(m.memory[uint32(v8+v5):], uint64(t50))
				goto l25
			}
		l24:
			if v5 == 0 {
				goto l25
			}
			memory_copy(m.memory, uint32(v8+i32(8)), uint32(v8), uint32(v5))
		l25:
			t51 := int64(load64(m.memory[int64(uint32(v1))+8:]))
			v16 = t51
			t52 := int64(load64(m.memory[uint32(v1):]))
			v18 = t52
			v6 = i32(0)
		l33:
			{
				t53 := v8
				v3 = v6
				v10 = t53 + v3
				t54 := int32(m.memory[uint32(v10)])
				if t54 != i32(128) {
					goto l26
				}
				v15 = v8 + (v3^i32(-1))*i32(36)
				v6 = v8 + (i32(0)-v3)*i32(36)
				v12 = v6 + i32(-28)
				v1 = v6 + i32(-32)
			l32:
				{
					t55 := int32(load32(m.memory[uint32(v1):]))
					t56 := int32(load32(m.memory[uint32(v12):]))
					t57 := m.fn64(v18, v16, t55, t56)
					t58 := v4
					v11 = int32(t57)
					v6 = t58 & v11
					v5 = v6
					{
						t59 := int64(load64(m.memory[uint32(v8+v6):]))
						v9 = t59 & i64(-0x7f7f7f7f7f7f7f80)
						if v9 != i64(0) {
							goto l27
						}
						v17 = i32(8)
						v5 = v6
					l28:
						{
							v5 = v5 + v17
							v17 = v17 + i32(8)
							t60 := v8
							v5 = v5 & v4
							t61 := int64(load64(m.memory[uint32(t60+v5):]))
							v9 = t61 & i64(-0x7f7f7f7f7f7f7f80)
							if v9 == 0 {
								goto l28
							}
						}
					}
				l27:
					{
						t62 := v8
						v5 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v9))))>>3) + v5) & v4
						t63 := int32(int8(m.memory[uint32(t62+v5)]))
						if t63 < i32(0) {
							goto l29
						}
						t64 := int64(load64(m.memory[uint32(v8):]))
						v5 = int32(uint32(int64(bits.TrailingZeros64(uint64(t64&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
					}
				l29:
					{
						if uint32((v5-v6^(v3-v6))&v4) < uint32(i32(8)) {
							goto l30
						}
						v6 = v8 + v5
						t65 := int32(m.memory[uint32(v6)])
						v17 = t65
						t66 := v6
						v11 = int32(uint32(v11) >> 25)
						m.memory[uint32(t66)] = byte(v11)
						m.memory[uint32(v8+(v5+i32(-8))&v4+i32(8))] = byte(v11)
						v6 = v8 + (v5^i32(-1))*i32(36)
						{
							if v17 != i32(255) {
								t72 := int32(load32(m.memory[uint32(v15):]))
								v5 = t72
								t73 := int32(load32(m.memory[uint32(v6):]))
								store32(m.memory[uint32(v15):], uint32(t73))
								store32(m.memory[uint32(v6):], uint32(v5))
								t74 := int32(load32(m.memory[int64(uint32(v6))+4:]))
								v5 = t74
								t75 := int32(load32(m.memory[int64(uint32(v15))+4:]))
								store32(m.memory[int64(uint32(v6))+4:], uint32(t75))
								store32(m.memory[int64(uint32(v15))+4:], uint32(v5))
								t76 := int32(load32(m.memory[int64(uint32(v15))+8:]))
								v5 = t76
								t77 := int32(load32(m.memory[int64(uint32(v6))+8:]))
								store32(m.memory[int64(uint32(v15))+8:], uint32(t77))
								store32(m.memory[int64(uint32(v6))+8:], uint32(v5))
								t78 := int32(load32(m.memory[int64(uint32(v6))+12:]))
								v5 = t78
								t79 := int32(load32(m.memory[int64(uint32(v15))+12:]))
								store32(m.memory[int64(uint32(v6))+12:], uint32(t79))
								store32(m.memory[int64(uint32(v15))+12:], uint32(v5))
								t80 := int32(load32(m.memory[int64(uint32(v15))+16:]))
								v5 = t80
								t81 := int32(load32(m.memory[int64(uint32(v6))+16:]))
								store32(m.memory[int64(uint32(v15))+16:], uint32(t81))
								store32(m.memory[int64(uint32(v6))+16:], uint32(v5))
								t82 := int32(load32(m.memory[int64(uint32(v6))+20:]))
								v5 = t82
								t83 := int32(load32(m.memory[int64(uint32(v15))+20:]))
								store32(m.memory[int64(uint32(v6))+20:], uint32(t83))
								store32(m.memory[int64(uint32(v15))+20:], uint32(v5))
								t84 := int32(load32(m.memory[int64(uint32(v15))+24:]))
								v5 = t84
								t85 := int32(load32(m.memory[int64(uint32(v6))+24:]))
								store32(m.memory[int64(uint32(v15))+24:], uint32(t85))
								store32(m.memory[int64(uint32(v6))+24:], uint32(v5))
								t86 := int32(load32(m.memory[int64(uint32(v6))+28:]))
								v5 = t86
								t87 := int32(load32(m.memory[int64(uint32(v15))+28:]))
								store32(m.memory[int64(uint32(v6))+28:], uint32(t87))
								store32(m.memory[int64(uint32(v15))+28:], uint32(v5))
								t88 := int32(load32(m.memory[int64(uint32(v15))+32:]))
								v5 = t88
								t89 := int32(load32(m.memory[int64(uint32(v6))+32:]))
								store32(m.memory[int64(uint32(v15))+32:], uint32(t89))
								store32(m.memory[int64(uint32(v6))+32:], uint32(v5))
								goto l32
							}
							m.memory[uint32(v10)] = byte(i32(255))
							m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(i32(255))
							t67 := int32(load32(m.memory[int64(uint32(v15))+32:]))
							store32(m.memory[int64(uint32(v6))+32:], uint32(t67))
							t68 := int64(load64(m.memory[int64(uint32(v15))+24:]))
							store64(m.memory[int64(uint32(v6))+24:], uint64(t68))
							t69 := int64(load64(m.memory[int64(uint32(v15))+16:]))
							store64(m.memory[int64(uint32(v6))+16:], uint64(t69))
							t70 := int64(load64(m.memory[int64(uint32(v15))+8:]))
							store64(m.memory[int64(uint32(v6))+8:], uint64(t70))
							t71 := int64(load64(m.memory[uint32(v15):]))
							store64(m.memory[uint32(v6):], uint64(t71))
							goto l26
						}
					}
				l30:
				}
				t90 := v10
				v6 = int32(uint32(v11) >> 25)
				m.memory[uint32(t90)] = byte(v6)
				m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(v6)
			}
		l26:
			v6 = v3 + i32(1)
			if v3 != v4 {
				goto l33
			}
			p91 := v7
			if uint32(v4) < uint32(i32(8)) {
				p91 = v4
			}
			v3 = p91
		}
	l20:
		store32(m.memory[int64(uint32(v0))+8:], uint32(v3-v2))
	l15:
		return i32(-1)
	}
l5:
	m.fn27(i32(1271632), i32(57), i32(1271660))
	panic("unreachable")
}
func (m *Module) fn76(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7, v8 int32
	var v9 int64
	var v10, v11, v12 int32
	var v13, v14 int64
	var v15 int32
	var v16 int64
	var v17 int32
	var v18 int64
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
		v2 = t0
		v3 = v2 + i32(1)
		if v3 == 0 {
			m.fn27(i32(1271632), i32(57), i32(1271660))
			panic("unreachable")
		}
		{
			t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t2 := v3
			v4 = t1
			t3 := v4
			v5 = v4 + i32(1)
			v6 = int32(uint32(v5) >> 3)
			v7 = v6 * i32(7)
			p4 := v7
			if uint32(v4) < uint32(i32(8)) {
				p4 = t3
			}
			v8 = p4
			if uint32(t2) <= uint32(int32(uint32(v8)>>1)) {
				goto l1
			}
			{
				{
					{
						v8 = v8 + i32(1)
						p5 := v3
						if uint32(v8) > uint32(v3) {
							p5 = v8
						}
						v3 = p5
						if uint32(v3) < uint32(i32(15)) {
							goto l2
						}
						if uint32(v3) > uint32(i32(0x1fffffff)) {
							m.fn27(i32(1271632), i32(57), i32(1271660))
							panic("unreachable")
						}
						t6 := int32(uint32(v3<<3) / uint32(i32(7)))
						v3 = i32_shr_u(i32(-1), int32(bits.LeadingZeros32(uint32(t6+i32(-1))))) + i32(1)
						goto l4
					}
				l2:
					p7 := v3&i32(8) + i32(8)
					if uint32(v3) < uint32(i32(4)) {
						p7 = i32(4)
					}
					v3 = p7
				}
			l4:
				v9 = int64(uint32(v3)) * i64(28)
				if int32(int64(uint64(v9)>>32)) != 0 {
					goto l5
				}
				v6 = int32(v9)
				if uint32(v6) > uint32(i32(-8)) {
					goto l5
				}
				v8 = v3 + i32(8)
				t8 := v8
				v10 = (v6 + i32(7)) & i32(-8)
				v6 = t8 + v10
				if uint32(v6) < uint32(v8) {
					goto l5
				}
				if uint32(v6) > uint32(i32(0x7ffffff8)) {
					goto l5
				}
				t9 := m.fn7(v6)
				v5 = t9
				if v5 != 0 {
					v6 = v5 + v10
					if v8 == 0 {
						goto l7
					}
					memory_fill(m.memory, uint32(v6), i32(255), uint32(v8))
				l7:
					v11 = v3 + i32(-1)
					p10 := int32(uint32(v3)>>3) * i32(7)
					if uint32(v3) < uint32(i32(9)) {
						p10 = v11
					}
					v7 = p10
					t11 := int32(load32(m.memory[uint32(v0):]))
					v12 = t11
					{
						if v2 == 0 {
							goto l8
						}
						t12 := int64(load64(m.memory[uint32(v12):]))
						v9 = (t12 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
						t13 := int64(load64(m.memory[int64(uint32(v1))+8:]))
						v13 = t13
						t14 := int64(load64(m.memory[uint32(v1):]))
						v14 = t14
						v8 = v12
						v1 = v2
						v3 = i32(0)
					l14:
						{
							if v9 != i64(0) {
								goto l9
							}
						l10:
							{
								v3 = v3 + i32(8)
								v8 = v8 + i32(8)
								t15 := int64(load64(m.memory[uint32(v8):]))
								v9 = t15 & i64(-0x7f7f7f7f7f7f7f80)
								if v9 == i64(-0x7f7f7f7f7f7f7f80) {
									goto l10
								}
							}
							v9 = v9 ^ i64(-0x7f7f7f7f7f7f7f80)
						l9:
							{
								t16 := v6
								t17 := v11
								t18 := v14
								t19 := v13
								t20 := v12
								v5 = int32(uint32(int64(bits.TrailingZeros64(uint64(v9))))>>3) + v3
								v10 = t20 + (i32(0)-v5)*i32(28)
								t21 := int32(load32(m.memory[uint32(v10+i32(-24)):]))
								t22 := int32(load32(m.memory[uint32(v10+i32(-20)):]))
								t23 := m.fn64(t18, t19, t21, t22)
								v15 = int32(t23)
								v10 = t17 & v15
								t24 := int64(load64(m.memory[uint32(t16+v10):]))
								v16 = t24 & i64(-0x7f7f7f7f7f7f7f80)
								if v16 != i64(0) {
									goto l11
								}
								v17 = i32(8)
							l12:
								{
									v10 = v10 + v17
									v17 = v17 + i32(8)
									t25 := v6
									v10 = v10 & v11
									t26 := int64(load64(m.memory[uint32(t25+v10):]))
									v16 = t26 & i64(-0x7f7f7f7f7f7f7f80)
									if v16 == 0 {
										goto l12
									}
								}
							}
						l11:
							v18 = v9 + i64(-1)
							{
								t27 := v6
								v10 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v16))))>>3) + v10) & v11
								t28 := int32(int8(m.memory[uint32(t27+v10)]))
								if t28 < i32(0) {
									goto l13
								}
								t29 := int64(load64(m.memory[uint32(v6):]))
								v10 = int32(uint32(int64(bits.TrailingZeros64(uint64(t29&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
							}
						l13:
							v9 = v18 & v9
							t30 := v6 + v10
							v15 = int32(uint32(v15) >> 25)
							m.memory[uint32(t30)] = byte(v15)
							m.memory[uint32(v6+(v10+i32(-8))&v11+i32(8))] = byte(v15)
							v10 = v6 + (v10^i32(-1))*i32(28)
							t31 := v10
							v5 = v12 + (v5^i32(-1))*i32(28)
							t32 := int32(load32(m.memory[int64(uint32(v5))+24:]))
							store32(m.memory[int64(uint32(t31))+24:], uint32(t32))
							t33 := int64(load64(m.memory[int64(uint32(v5))+16:]))
							store64(m.memory[int64(uint32(v10))+16:], uint64(t33))
							t34 := int64(load64(m.memory[int64(uint32(v5))+8:]))
							store64(m.memory[int64(uint32(v10))+8:], uint64(t34))
							t35 := int64(load64(m.memory[uint32(v5):]))
							store64(m.memory[uint32(v10):], uint64(t35))
							v1 = v1 + i32(-1)
							if v1 != 0 {
								goto l14
							}
						}
					}
				l8:
					store32(m.memory[int64(uint32(v0))+4:], uint32(v11))
					store32(m.memory[uint32(v0):], uint32(v6))
					store32(m.memory[int64(uint32(v0))+8:], uint32(v7-v2))
					if v4 == 0 {
						goto l15
					}
					t36 := v4
					v8 = (v4*i32(28) + i32(35)) & i32(-8)
					v3 = t36 + v8 + i32(9)
					if v3 == 0 {
						goto l15
					}
					v4 = v12 - v8
					t37 := int32(load32(m.memory[uint32(v4+i32(-4)):]))
					v8 = t37
					v6 = v8 & i32(-8)
					t38 := v6
					v8 = v8 & i32(3)
					p39 := i32(8)
					if v8 != 0 {
						p39 = i32(4)
					}
					if uint32(t38) < uint32(p39+v3) {
						m.fn3(i32(1274224), i32(46), i32(1274272))
						panic("unreachable")
					}
					if v8 == 0 {
						goto l17
					}
					if uint32(v6) > uint32(v3+i32(39)) {
						m.fn3(i32(1274288), i32(46), i32(1274336))
						panic("unreachable")
					}
				l17:
					m.fn1(v4)
					return i32(-1)
				}
				m.fn23(i32(8), v6)
				panic("unreachable")
			}
		}
	l1:
		{
			if v5 != 0 {
				goto l19
			}
			v3 = i32(0)
			goto l20
		l19:
			t40 := int32(load32(m.memory[uint32(v0):]))
			v8 = t40
			v3 = i32(0)
			{
				{
					t41 := v6
					var p42 int32
					if v5&i32(7) != i32(0) {
						p42 = 1
					}
					v6 = t41 + p42
					if v6 == i32(1) {
						goto l21
					}
					v11 = v6 & i32(1)
					v10 = v6 & i32(0x3ffffffe)
					v3 = i32(0)
				l22:
					{
						v6 = v8 + v3
						t43 := int64(load64(m.memory[uint32(v6):]))
						t44 := v6
						v9 = t43
						store64(m.memory[uint32(t44):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
						v6 = v6 + i32(8)
						t45 := int64(load64(m.memory[uint32(v6):]))
						t46 := v6
						v9 = t45
						store64(m.memory[uint32(t46):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
						v3 = v3 + i32(16)
						v10 = v10 + i32(-2)
						if v10 != 0 {
							goto l22
						}
					}
					if v11 == 0 {
						goto l23
					}
				}
			l21:
				v3 = v8 + v3
				t47 := int64(load64(m.memory[uint32(v3):]))
				t48 := v3
				v9 = t47
				store64(m.memory[uint32(t48):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
			}
		l23:
			{
				if uint32(v5) < uint32(i32(8)) {
					goto l24
				}
				t49 := int64(load64(m.memory[uint32(v8):]))
				store64(m.memory[uint32(v8+v5):], uint64(t49))
				goto l25
			}
		l24:
			if v5 == 0 {
				goto l25
			}
			memory_copy(m.memory, uint32(v8+i32(8)), uint32(v8), uint32(v5))
		l25:
			t50 := int64(load64(m.memory[int64(uint32(v1))+8:]))
			v16 = t50
			t51 := int64(load64(m.memory[uint32(v1):]))
			v18 = t51
			v6 = i32(0)
		l33:
			{
				t52 := v8
				v3 = v6
				v10 = t52 + v3
				t53 := int32(m.memory[uint32(v10)])
				if t53 != i32(128) {
					goto l26
				}
				v15 = v8 + (v3^i32(-1))*i32(28)
				v6 = v8 + (i32(0)-v3)*i32(28)
				v12 = v6 + i32(-20)
				v1 = v6 + i32(-24)
			l32:
				{
					t54 := int32(load32(m.memory[uint32(v1):]))
					t55 := int32(load32(m.memory[uint32(v12):]))
					t56 := m.fn64(v18, v16, t54, t55)
					t57 := v4
					v11 = int32(t56)
					v6 = t57 & v11
					v5 = v6
					{
						t58 := int64(load64(m.memory[uint32(v8+v6):]))
						v9 = t58 & i64(-0x7f7f7f7f7f7f7f80)
						if v9 != i64(0) {
							goto l27
						}
						v17 = i32(8)
						v5 = v6
					l28:
						{
							v5 = v5 + v17
							v17 = v17 + i32(8)
							t59 := v8
							v5 = v5 & v4
							t60 := int64(load64(m.memory[uint32(t59+v5):]))
							v9 = t60 & i64(-0x7f7f7f7f7f7f7f80)
							if v9 == 0 {
								goto l28
							}
						}
					}
				l27:
					{
						t61 := v8
						v5 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v9))))>>3) + v5) & v4
						t62 := int32(int8(m.memory[uint32(t61+v5)]))
						if t62 < i32(0) {
							goto l29
						}
						t63 := int64(load64(m.memory[uint32(v8):]))
						v5 = int32(uint32(int64(bits.TrailingZeros64(uint64(t63&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
					}
				l29:
					{
						if uint32((v5-v6^(v3-v6))&v4) < uint32(i32(8)) {
							goto l30
						}
						v6 = v8 + v5
						t64 := int32(m.memory[uint32(v6)])
						v17 = t64
						t65 := v6
						v11 = int32(uint32(v11) >> 25)
						m.memory[uint32(t65)] = byte(v11)
						m.memory[uint32(v8+(v5+i32(-8))&v4+i32(8))] = byte(v11)
						v6 = v8 + (v5^i32(-1))*i32(28)
						{
							if v17 != i32(255) {
								t70 := int32(load32(m.memory[uint32(v6):]))
								v5 = t70
								t71 := int32(load32(m.memory[uint32(v15):]))
								store32(m.memory[uint32(v6):], uint32(t71))
								store32(m.memory[uint32(v15):], uint32(v5))
								t72 := int32(load32(m.memory[int64(uint32(v15))+4:]))
								v5 = t72
								t73 := int32(load32(m.memory[int64(uint32(v6))+4:]))
								store32(m.memory[int64(uint32(v15))+4:], uint32(t73))
								store32(m.memory[int64(uint32(v6))+4:], uint32(v5))
								t74 := int32(load32(m.memory[int64(uint32(v6))+8:]))
								v5 = t74
								t75 := int32(load32(m.memory[int64(uint32(v15))+8:]))
								store32(m.memory[int64(uint32(v6))+8:], uint32(t75))
								store32(m.memory[int64(uint32(v15))+8:], uint32(v5))
								t76 := int32(load32(m.memory[int64(uint32(v15))+12:]))
								v5 = t76
								t77 := int32(load32(m.memory[int64(uint32(v6))+12:]))
								store32(m.memory[int64(uint32(v15))+12:], uint32(t77))
								store32(m.memory[int64(uint32(v6))+12:], uint32(v5))
								t78 := int32(load32(m.memory[int64(uint32(v6))+16:]))
								v5 = t78
								t79 := int32(load32(m.memory[int64(uint32(v15))+16:]))
								store32(m.memory[int64(uint32(v6))+16:], uint32(t79))
								store32(m.memory[int64(uint32(v15))+16:], uint32(v5))
								t80 := int32(load32(m.memory[int64(uint32(v15))+20:]))
								v5 = t80
								t81 := int32(load32(m.memory[int64(uint32(v6))+20:]))
								store32(m.memory[int64(uint32(v15))+20:], uint32(t81))
								store32(m.memory[int64(uint32(v6))+20:], uint32(v5))
								t82 := int32(load32(m.memory[int64(uint32(v6))+24:]))
								v5 = t82
								t83 := int32(load32(m.memory[int64(uint32(v15))+24:]))
								store32(m.memory[int64(uint32(v6))+24:], uint32(t83))
								store32(m.memory[int64(uint32(v15))+24:], uint32(v5))
								goto l32
							}
							m.memory[uint32(v10)] = byte(i32(255))
							m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(i32(255))
							t66 := int32(load32(m.memory[int64(uint32(v15))+24:]))
							store32(m.memory[int64(uint32(v6))+24:], uint32(t66))
							t67 := int64(load64(m.memory[int64(uint32(v15))+16:]))
							store64(m.memory[int64(uint32(v6))+16:], uint64(t67))
							t68 := int64(load64(m.memory[int64(uint32(v15))+8:]))
							store64(m.memory[int64(uint32(v6))+8:], uint64(t68))
							t69 := int64(load64(m.memory[uint32(v15):]))
							store64(m.memory[uint32(v6):], uint64(t69))
							goto l26
						}
					}
				l30:
				}
				t84 := v10
				v6 = int32(uint32(v11) >> 25)
				m.memory[uint32(t84)] = byte(v6)
				m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(v6)
			}
		l26:
			v6 = v3 + i32(1)
			if v3 != v4 {
				goto l33
			}
			p85 := v7
			if uint32(v4) < uint32(i32(8)) {
				p85 = v4
			}
			v3 = p85
		}
	l20:
		store32(m.memory[int64(uint32(v0))+8:], uint32(v3-v2))
	l15:
		return i32(-1)
	}
l5:
	m.fn27(i32(1271632), i32(57), i32(1271660))
	panic("unreachable")
}
func (m *Module) fn77(v0, v1, v2 int32) int32 {
	var v3, v4, v5, v6, v7, v8, v9, v10 int32
	var v11, v12, v13 int64
	var v14, v15, v16 int32
	var v17, v18 int64
	{
		{
			{
				t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
				v3 = t0
				v1 = v3 + v1
				if uint32(v1) < uint32(v3) {
					m.fn27(i32(1271632), i32(57), i32(1271660))
					panic("unreachable")
				}
				{
					t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
					t2 := v1
					v4 = t1
					t3 := v4
					v5 = v4 + i32(1)
					v6 = int32(uint32(v5) >> 3)
					v7 = v6 * i32(7)
					p4 := v7
					if uint32(v4) < uint32(i32(8)) {
						p4 = t3
					}
					v8 = p4
					if uint32(t2) <= uint32(int32(uint32(v8)>>1)) {
						{
							if v5 != 0 {
								goto l6
							}
							v1 = i32(0)
							goto l7
						l6:
							t7 := int32(load32(m.memory[uint32(v0):]))
							v8 = t7
							v1 = i32(0)
							{
								{
									t8 := v6
									var p9 int32
									if v5&i32(7) != i32(0) {
										p9 = 1
									}
									v6 = t8 + p9
									if v6 == i32(1) {
										goto l8
									}
									v9 = v6 & i32(1)
									v10 = v6 & i32(0x3ffffffe)
									v1 = i32(0)
								l9:
									{
										v6 = v8 + v1
										t10 := int64(load64(m.memory[uint32(v6):]))
										t11 := v6
										v11 = t10
										store64(m.memory[uint32(t11):], uint64(int64(uint64(v11^i64(-1))>>7)&i64(72340172838076673)+(v11|i64(0x7f7f7f7f7f7f7f7f))))
										v6 = v6 + i32(8)
										t12 := int64(load64(m.memory[uint32(v6):]))
										t13 := v6
										v11 = t12
										store64(m.memory[uint32(t13):], uint64(int64(uint64(v11^i64(-1))>>7)&i64(72340172838076673)+(v11|i64(0x7f7f7f7f7f7f7f7f))))
										v1 = v1 + i32(16)
										v10 = v10 + i32(-2)
										if v10 != 0 {
											goto l9
										}
									}
									if v9 == 0 {
										goto l10
									}
								}
							l8:
								v1 = v8 + v1
								t14 := int64(load64(m.memory[uint32(v1):]))
								t15 := v1
								v11 = t14
								store64(m.memory[uint32(t15):], uint64(int64(uint64(v11^i64(-1))>>7)&i64(72340172838076673)+(v11|i64(0x7f7f7f7f7f7f7f7f))))
							}
						l10:
							{
								if uint32(v5) < uint32(i32(8)) {
									goto l11
								}
								t16 := int64(load64(m.memory[uint32(v8):]))
								store64(m.memory[uint32(v8+v5):], uint64(t16))
								goto l12
							}
						l11:
							if v5 == 0 {
								goto l12
							}
							memory_copy(m.memory, uint32(v8+i32(8)), uint32(v8), uint32(v5))
						l12:
							t17 := int64(load64(m.memory[int64(uint32(v2))+8:]))
							v12 = t17
							t18 := int64(load64(m.memory[uint32(v2):]))
							v13 = t18
							v6 = i32(0)
						l20:
							{
								t19 := v8
								v1 = v6
								v10 = t19 + v1
								t20 := int32(m.memory[uint32(v10)])
								if t20 != i32(128) {
									goto l13
								}
								v6 = v8 - v1<<4
								v2 = v6 + i32(-8)
								v14 = v6 + i32(-12)
								v15 = v8 + (v1^i32(-1))<<4
							l19:
								{
									t21 := int32(load32(m.memory[uint32(v14):]))
									t22 := int32(load32(m.memory[uint32(v2):]))
									t23 := m.fn64(v13, v12, t21, t22)
									t24 := v4
									v9 = int32(t23)
									v6 = t24 & v9
									v5 = v6
									{
										t25 := int64(load64(m.memory[uint32(v8+v6):]))
										v11 = t25 & i64(-0x7f7f7f7f7f7f7f80)
										if v11 != i64(0) {
											goto l14
										}
										v16 = i32(8)
										v5 = v6
									l15:
										{
											v5 = v5 + v16
											v16 = v16 + i32(8)
											t26 := v8
											v5 = v5 & v4
											t27 := int64(load64(m.memory[uint32(t26+v5):]))
											v11 = t27 & i64(-0x7f7f7f7f7f7f7f80)
											if v11 == 0 {
												goto l15
											}
										}
									}
								l14:
									{
										t28 := v8
										v5 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v11))))>>3) + v5) & v4
										t29 := int32(int8(m.memory[uint32(t28+v5)]))
										if t29 < i32(0) {
											goto l16
										}
										t30 := int64(load64(m.memory[uint32(v8):]))
										v5 = int32(uint32(int64(bits.TrailingZeros64(uint64(t30&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
									}
								l16:
									{
										if uint32((v5-v6^(v1-v6))&v4) < uint32(i32(8)) {
											goto l17
										}
										v6 = v8 + v5
										t31 := int32(m.memory[uint32(v6)])
										v16 = t31
										t32 := v6
										v9 = int32(uint32(v9) >> 25)
										m.memory[uint32(t32)] = byte(v9)
										m.memory[uint32(v8+(v5+i32(-8))&v4+i32(8))] = byte(v9)
										v6 = v8 + (v5^i32(-1))<<4
										{
											if v16 != i32(255) {
												t35 := int64(load64(m.memory[uint32(v15):]))
												v11 = t35
												t36 := int64(load64(m.memory[uint32(v6):]))
												store64(m.memory[uint32(v15):], uint64(t36))
												store64(m.memory[uint32(v6):], uint64(v11))
												t37 := int64(load64(m.memory[int64(uint32(v15))+8:]))
												v11 = t37
												t38 := int64(load64(m.memory[int64(uint32(v6))+8:]))
												store64(m.memory[int64(uint32(v15))+8:], uint64(t38))
												store64(m.memory[int64(uint32(v6))+8:], uint64(v11))
												goto l19
											}
											m.memory[uint32(v10)] = byte(i32(255))
											m.memory[uint32(v8+v4&(v1+i32(-8))+i32(8))] = byte(i32(255))
											t33 := int64(load64(m.memory[int64(uint32(v15))+8:]))
											store64(m.memory[int64(uint32(v6))+8:], uint64(t33))
											t34 := int64(load64(m.memory[uint32(v15):]))
											store64(m.memory[uint32(v6):], uint64(t34))
											goto l13
										}
									}
								l17:
								}
								t39 := v10
								v6 = int32(uint32(v9) >> 25)
								m.memory[uint32(t39)] = byte(v6)
								m.memory[uint32(v8+v4&(v1+i32(-8))+i32(8))] = byte(v6)
							}
						l13:
							v6 = v1 + i32(1)
							if v1 != v4 {
								goto l20
							}
							p40 := v7
							if uint32(v4) < uint32(i32(8)) {
								p40 = v4
							}
							v1 = p40
						}
					l7:
						store32(m.memory[int64(uint32(v0))+8:], uint32(v1-v3))
						goto l21
					}
					v8 = v8 + i32(1)
					p5 := v1
					if uint32(v8) > uint32(v1) {
						p5 = v8
					}
					v1 = p5
					if uint32(v1) < uint32(i32(15)) {
						goto l2
					}
					{
						if uint32(v1) > uint32(i32(0x1fffffff)) {
							m.fn27(i32(1271632), i32(57), i32(1271660))
							panic("unreachable")
						}
						t6 := int32(uint32(v1<<3) / uint32(i32(7)))
						v1 = i32_shr_u(i32(-1), int32(bits.LeadingZeros32(uint32(t6+i32(-1)))))
						if uint32(v1) > uint32(i32(0xffffffe)) {
							goto l4
						}
						v1 = v1 + i32(1)
						goto l5
					}
				}
			}
		l2:
			p41 := v1&i32(8) + i32(8)
			if uint32(v1) < uint32(i32(4)) {
				p41 = i32(4)
			}
			v1 = p41
		}
	l5:
		v8 = v1 + i32(8)
		t42 := v8
		v10 = v1 << 4
		v6 = t42 + v10
		if uint32(v6) < uint32(v8) {
			goto l4
		}
		if uint32(v6) > uint32(i32(0x7ffffff8)) {
			goto l4
		}
		{
			t43 := m.fn7(v6)
			v5 = t43
			if v5 != 0 {
				v6 = v5 + v10
				if v8 == 0 {
					goto l23
				}
				memory_fill(m.memory, uint32(v6), i32(255), uint32(v8))
			l23:
				v5 = v1 + i32(-1)
				p44 := int32(uint32(v1)>>3) * i32(7)
				if uint32(v1) < uint32(i32(9)) {
					p44 = v5
				}
				v7 = p44
				t45 := int32(load32(m.memory[uint32(v0):]))
				v9 = t45
				{
					if v3 == 0 {
						goto l24
					}
					t46 := int64(load64(m.memory[uint32(v9):]))
					v11 = (t46 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
					t47 := int64(load64(m.memory[int64(uint32(v2))+8:]))
					v17 = t47
					t48 := int64(load64(m.memory[uint32(v2):]))
					v18 = t48
					v8 = v9
					v2 = v3
					v1 = i32(0)
				l30:
					{
						if v11 != i64(0) {
							goto l25
						}
					l26:
						{
							v1 = v1 + i32(8)
							v8 = v8 + i32(8)
							t49 := int64(load64(m.memory[uint32(v8):]))
							v11 = t49 & i64(-0x7f7f7f7f7f7f7f80)
							if v11 == i64(-0x7f7f7f7f7f7f7f80) {
								goto l26
							}
						}
						v11 = v11 ^ i64(-0x7f7f7f7f7f7f7f80)
					l25:
						{
							t50 := v6
							t51 := v5
							t52 := v18
							t53 := v17
							t54 := v9
							v14 = int32(uint32(int64(bits.TrailingZeros64(uint64(v11))))>>3) + v1
							v10 = t54 - v14<<4
							t55 := int32(load32(m.memory[uint32(v10+i32(-12)):]))
							t56 := int32(load32(m.memory[uint32(v10+i32(-8)):]))
							t57 := m.fn64(t52, t53, t55, t56)
							v15 = int32(t57)
							v10 = t51 & v15
							t58 := int64(load64(m.memory[uint32(t50+v10):]))
							v12 = t58 & i64(-0x7f7f7f7f7f7f7f80)
							if v12 != i64(0) {
								goto l27
							}
							v16 = i32(8)
						l28:
							{
								v10 = v10 + v16
								v16 = v16 + i32(8)
								t59 := v6
								v10 = v10 & v5
								t60 := int64(load64(m.memory[uint32(t59+v10):]))
								v12 = t60 & i64(-0x7f7f7f7f7f7f7f80)
								if v12 == 0 {
									goto l28
								}
							}
						}
					l27:
						v13 = v11 + i64(-1)
						{
							t61 := v6
							v10 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v12))))>>3) + v10) & v5
							t62 := int32(int8(m.memory[uint32(t61+v10)]))
							if t62 < i32(0) {
								goto l29
							}
							t63 := int64(load64(m.memory[uint32(v6):]))
							v10 = int32(uint32(int64(bits.TrailingZeros64(uint64(t63&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
						}
					l29:
						v11 = v13 & v11
						t64 := v6 + v10
						v15 = int32(uint32(v15) >> 25)
						m.memory[uint32(t64)] = byte(v15)
						m.memory[uint32(v6+(v10+i32(-8))&v5+i32(8))] = byte(v15)
						v10 = v6 + (v10^i32(-1))<<4
						t65 := v10
						v14 = v9 + (v14^i32(-1))<<4
						t66 := int64(load64(m.memory[int64(uint32(v14))+8:]))
						store64(m.memory[int64(uint32(t65))+8:], uint64(t66))
						t67 := int64(load64(m.memory[uint32(v14):]))
						store64(m.memory[uint32(v10):], uint64(t67))
						v2 = v2 + i32(-1)
						if v2 != 0 {
							goto l30
						}
					}
				}
			l24:
				store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
				store32(m.memory[uint32(v0):], uint32(v6))
				store32(m.memory[int64(uint32(v0))+8:], uint32(v7-v3))
				if v4 == 0 {
					goto l21
				}
				t68 := v4
				v8 = (v4<<4 + i32(23)) & i32(-16)
				v1 = t68 + v8 + i32(9)
				if v1 == 0 {
					goto l21
				}
				v4 = v9 - v8
				t69 := int32(load32(m.memory[uint32(v4+i32(-4)):]))
				v8 = t69
				v6 = v8 & i32(-8)
				t70 := v6
				v8 = v8 & i32(3)
				p71 := i32(8)
				if v8 != 0 {
					p71 = i32(4)
				}
				if uint32(t70) < uint32(p71+v1) {
					m.fn3(i32(1274224), i32(46), i32(1274272))
					panic("unreachable")
				}
				if v8 == 0 {
					goto l32
				}
				if uint32(v6) > uint32(v1+i32(39)) {
					m.fn3(i32(1274288), i32(46), i32(1274336))
					panic("unreachable")
				}
			l32:
				m.fn1(v4)
				return i32(-1)
			}
			m.fn23(i32(8), v6)
			panic("unreachable")
		}
	}
l21:
	return i32(-1)
l4:
	m.fn27(i32(1271632), i32(57), i32(1271660))
	panic("unreachable")
}
func (m *Module) fn78(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7, v8 int32
	var v9 int64
	var v10, v11 int32
	var v12, v13 int64
	var v14, v15 int32
	var v16 int64
	var v17 int32
	var v18 int64
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
		v2 = t0
		v3 = v2 + i32(1)
		if v3 == 0 {
			m.fn27(i32(1271632), i32(57), i32(1271660))
			panic("unreachable")
		}
		{
			t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t2 := v3
			v4 = t1
			t3 := v4
			v5 = v4 + i32(1)
			v6 = int32(uint32(v5) >> 3)
			v7 = v6 * i32(7)
			p4 := v7
			if uint32(v4) < uint32(i32(8)) {
				p4 = t3
			}
			v8 = p4
			if uint32(t2) <= uint32(int32(uint32(v8)>>1)) {
				goto l1
			}
			{
				{
					{
						v8 = v8 + i32(1)
						p5 := v3
						if uint32(v8) > uint32(v3) {
							p5 = v8
						}
						v3 = p5
						if uint32(v3) < uint32(i32(15)) {
							goto l2
						}
						if uint32(v3) > uint32(i32(0x1fffffff)) {
							m.fn27(i32(1271632), i32(57), i32(1271660))
							panic("unreachable")
						}
						t6 := int32(uint32(v3<<3) / uint32(i32(7)))
						v3 = i32_shr_u(i32(-1), int32(bits.LeadingZeros32(uint32(t6+i32(-1))))) + i32(1)
						goto l4
					}
				l2:
					p7 := v3&i32(8) + i32(8)
					if uint32(v3) < uint32(i32(4)) {
						p7 = i32(4)
					}
					v3 = p7
				}
			l4:
				v9 = int64(uint32(v3)) * i64(12)
				if int32(int64(uint64(v9)>>32)) != 0 {
					goto l5
				}
				v6 = int32(v9)
				if uint32(v6) > uint32(i32(-8)) {
					goto l5
				}
				v8 = v3 + i32(8)
				t8 := v8
				v10 = (v6 + i32(7)) & i32(-8)
				v6 = t8 + v10
				if uint32(v6) < uint32(v8) {
					goto l5
				}
				if uint32(v6) > uint32(i32(0x7ffffff8)) {
					goto l5
				}
				t9 := m.fn7(v6)
				v5 = t9
				if v5 != 0 {
					v6 = v5 + v10
					if v8 == 0 {
						goto l7
					}
					memory_fill(m.memory, uint32(v6), i32(255), uint32(v8))
				l7:
					v5 = v3 + i32(-1)
					p10 := int32(uint32(v3)>>3) * i32(7)
					if uint32(v3) < uint32(i32(9)) {
						p10 = v5
					}
					v7 = p10
					t11 := int32(load32(m.memory[uint32(v0):]))
					v11 = t11
					{
						if v2 == 0 {
							goto l8
						}
						t12 := int64(load64(m.memory[uint32(v11):]))
						v9 = (t12 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
						t13 := int64(load64(m.memory[int64(uint32(v1))+8:]))
						v12 = t13
						t14 := int64(load64(m.memory[uint32(v1):]))
						v13 = t14
						v8 = v11
						v1 = v2
						v3 = i32(0)
					l14:
						{
							if v9 != i64(0) {
								goto l9
							}
						l10:
							{
								v3 = v3 + i32(8)
								v8 = v8 + i32(8)
								t15 := int64(load64(m.memory[uint32(v8):]))
								v9 = t15 & i64(-0x7f7f7f7f7f7f7f80)
								if v9 == i64(-0x7f7f7f7f7f7f7f80) {
									goto l10
								}
							}
							v9 = v9 ^ i64(-0x7f7f7f7f7f7f7f80)
						l9:
							{
								t16 := v6
								t17 := v5
								t18 := v13
								t19 := v12
								t20 := v11
								v14 = int32(uint32(int64(bits.TrailingZeros64(uint64(v9))))>>3) + v3
								v10 = t20 + (i32(0)-v14)*i32(12)
								t21 := int32(load32(m.memory[uint32(v10+i32(-8)):]))
								t22 := int32(load32(m.memory[uint32(v10+i32(-4)):]))
								t23 := m.fn64(t18, t19, t21, t22)
								v15 = int32(t23)
								v10 = t17 & v15
								t24 := int64(load64(m.memory[uint32(t16+v10):]))
								v16 = t24 & i64(-0x7f7f7f7f7f7f7f80)
								if v16 != i64(0) {
									goto l11
								}
								v17 = i32(8)
							l12:
								{
									v10 = v10 + v17
									v17 = v17 + i32(8)
									t25 := v6
									v10 = v10 & v5
									t26 := int64(load64(m.memory[uint32(t25+v10):]))
									v16 = t26 & i64(-0x7f7f7f7f7f7f7f80)
									if v16 == 0 {
										goto l12
									}
								}
							}
						l11:
							v18 = v9 + i64(-1)
							{
								t27 := v6
								v10 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v16))))>>3) + v10) & v5
								t28 := int32(int8(m.memory[uint32(t27+v10)]))
								if t28 < i32(0) {
									goto l13
								}
								t29 := int64(load64(m.memory[uint32(v6):]))
								v10 = int32(uint32(int64(bits.TrailingZeros64(uint64(t29&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
							}
						l13:
							v9 = v18 & v9
							t30 := v6 + v10
							v15 = int32(uint32(v15) >> 25)
							m.memory[uint32(t30)] = byte(v15)
							m.memory[uint32(v6+(v10+i32(-8))&v5+i32(8))] = byte(v15)
							v10 = v6 + (v10^i32(-1))*i32(12)
							t31 := v10
							v14 = v11 + (v14^i32(-1))*i32(12)
							t32 := int32(load32(m.memory[int64(uint32(v14))+8:]))
							store32(m.memory[int64(uint32(t31))+8:], uint32(t32))
							t33 := int64(load64(m.memory[uint32(v14):]))
							store64(m.memory[uint32(v10):], uint64(t33))
							v1 = v1 + i32(-1)
							if v1 != 0 {
								goto l14
							}
						}
					}
				l8:
					store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
					store32(m.memory[uint32(v0):], uint32(v6))
					store32(m.memory[int64(uint32(v0))+8:], uint32(v7-v2))
					if v4 == 0 {
						goto l15
					}
					t34 := v4
					v8 = (v4*i32(12) + i32(19)) & i32(-8)
					v3 = t34 + v8 + i32(9)
					if v3 == 0 {
						goto l15
					}
					v4 = v11 - v8
					t35 := int32(load32(m.memory[uint32(v4+i32(-4)):]))
					v8 = t35
					v6 = v8 & i32(-8)
					t36 := v6
					v8 = v8 & i32(3)
					p37 := i32(8)
					if v8 != 0 {
						p37 = i32(4)
					}
					if uint32(t36) < uint32(p37+v3) {
						m.fn3(i32(1274224), i32(46), i32(1274272))
						panic("unreachable")
					}
					if v8 == 0 {
						goto l17
					}
					if uint32(v6) > uint32(v3+i32(39)) {
						m.fn3(i32(1274288), i32(46), i32(1274336))
						panic("unreachable")
					}
				l17:
					m.fn1(v4)
					return i32(-1)
				}
				m.fn23(i32(8), v6)
				panic("unreachable")
			}
		}
	l1:
		{
			if v5 != 0 {
				goto l19
			}
			v3 = i32(0)
			goto l20
		l19:
			t38 := int32(load32(m.memory[uint32(v0):]))
			v8 = t38
			v3 = i32(0)
			{
				{
					t39 := v6
					var p40 int32
					if v5&i32(7) != i32(0) {
						p40 = 1
					}
					v6 = t39 + p40
					if v6 == i32(1) {
						goto l21
					}
					v11 = v6 & i32(1)
					v10 = v6 & i32(0x3ffffffe)
					v3 = i32(0)
				l22:
					{
						v6 = v8 + v3
						t41 := int64(load64(m.memory[uint32(v6):]))
						t42 := v6
						v9 = t41
						store64(m.memory[uint32(t42):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
						v6 = v6 + i32(8)
						t43 := int64(load64(m.memory[uint32(v6):]))
						t44 := v6
						v9 = t43
						store64(m.memory[uint32(t44):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
						v3 = v3 + i32(16)
						v10 = v10 + i32(-2)
						if v10 != 0 {
							goto l22
						}
					}
					if v11 == 0 {
						goto l23
					}
				}
			l21:
				v3 = v8 + v3
				t45 := int64(load64(m.memory[uint32(v3):]))
				t46 := v3
				v9 = t45
				store64(m.memory[uint32(t46):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
			}
		l23:
			{
				if uint32(v5) < uint32(i32(8)) {
					goto l24
				}
				t47 := int64(load64(m.memory[uint32(v8):]))
				store64(m.memory[uint32(v8+v5):], uint64(t47))
				goto l25
			}
		l24:
			if v5 == 0 {
				goto l25
			}
			memory_copy(m.memory, uint32(v8+i32(8)), uint32(v8), uint32(v5))
		l25:
			t48 := int64(load64(m.memory[int64(uint32(v1))+8:]))
			v16 = t48
			t49 := int64(load64(m.memory[uint32(v1):]))
			v18 = t49
			v6 = i32(0)
		l33:
			{
				t50 := v8
				v3 = v6
				v10 = t50 + v3
				t51 := int32(m.memory[uint32(v10)])
				if t51 != i32(128) {
					goto l26
				}
				v15 = v8 + (v3^i32(-1))*i32(12)
				v6 = v8 + (i32(0)-v3)*i32(12)
				v1 = v6 + i32(-4)
				v14 = v6 + i32(-8)
			l32:
				{
					t52 := int32(load32(m.memory[uint32(v14):]))
					t53 := int32(load32(m.memory[uint32(v1):]))
					t54 := m.fn64(v18, v16, t52, t53)
					t55 := v4
					v11 = int32(t54)
					v6 = t55 & v11
					v5 = v6
					{
						t56 := int64(load64(m.memory[uint32(v8+v6):]))
						v9 = t56 & i64(-0x7f7f7f7f7f7f7f80)
						if v9 != i64(0) {
							goto l27
						}
						v17 = i32(8)
						v5 = v6
					l28:
						{
							v5 = v5 + v17
							v17 = v17 + i32(8)
							t57 := v8
							v5 = v5 & v4
							t58 := int64(load64(m.memory[uint32(t57+v5):]))
							v9 = t58 & i64(-0x7f7f7f7f7f7f7f80)
							if v9 == 0 {
								goto l28
							}
						}
					}
				l27:
					{
						t59 := v8
						v5 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v9))))>>3) + v5) & v4
						t60 := int32(int8(m.memory[uint32(t59+v5)]))
						if t60 < i32(0) {
							goto l29
						}
						t61 := int64(load64(m.memory[uint32(v8):]))
						v5 = int32(uint32(int64(bits.TrailingZeros64(uint64(t61&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
					}
				l29:
					{
						if uint32((v5-v6^(v3-v6))&v4) < uint32(i32(8)) {
							goto l30
						}
						v6 = v8 + v5
						t62 := int32(m.memory[uint32(v6)])
						v17 = t62
						t63 := v6
						v11 = int32(uint32(v11) >> 25)
						m.memory[uint32(t63)] = byte(v11)
						m.memory[uint32(v8+(v5+i32(-8))&v4+i32(8))] = byte(v11)
						v6 = v8 + (v5^i32(-1))*i32(12)
						{
							if v17 != i32(255) {
								t66 := int32(load32(m.memory[uint32(v15):]))
								v5 = t66
								t67 := int32(load32(m.memory[uint32(v6):]))
								store32(m.memory[uint32(v15):], uint32(t67))
								store32(m.memory[uint32(v6):], uint32(v5))
								t68 := int32(load32(m.memory[int64(uint32(v6))+4:]))
								v5 = t68
								t69 := int32(load32(m.memory[int64(uint32(v15))+4:]))
								store32(m.memory[int64(uint32(v6))+4:], uint32(t69))
								store32(m.memory[int64(uint32(v15))+4:], uint32(v5))
								t70 := int32(load32(m.memory[int64(uint32(v15))+8:]))
								v5 = t70
								t71 := int32(load32(m.memory[int64(uint32(v6))+8:]))
								store32(m.memory[int64(uint32(v15))+8:], uint32(t71))
								store32(m.memory[int64(uint32(v6))+8:], uint32(v5))
								goto l32
							}
							m.memory[uint32(v10)] = byte(i32(255))
							m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(i32(255))
							t64 := int32(load32(m.memory[int64(uint32(v15))+8:]))
							store32(m.memory[int64(uint32(v6))+8:], uint32(t64))
							t65 := int64(load64(m.memory[uint32(v15):]))
							store64(m.memory[uint32(v6):], uint64(t65))
							goto l26
						}
					}
				l30:
				}
				t72 := v10
				v6 = int32(uint32(v11) >> 25)
				m.memory[uint32(t72)] = byte(v6)
				m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(v6)
			}
		l26:
			v6 = v3 + i32(1)
			if v3 != v4 {
				goto l33
			}
			p73 := v7
			if uint32(v4) < uint32(i32(8)) {
				p73 = v4
			}
			v3 = p73
		}
	l20:
		store32(m.memory[int64(uint32(v0))+8:], uint32(v3-v2))
	l15:
		return i32(-1)
	}
l5:
	m.fn27(i32(1271632), i32(57), i32(1271660))
	panic("unreachable")
}
func (m *Module) fn79(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7, v8 int32
	var v9 int64
	var v10, v11 int32
	var v12, v13 int64
	var v14, v15 int32
	var v16 int64
	var v17 int32
	var v18 int64
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
		v2 = t0
		v3 = v2 + i32(1)
		if v3 == 0 {
			m.fn27(i32(1271632), i32(57), i32(1271660))
			panic("unreachable")
		}
		{
			t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t2 := v3
			v4 = t1
			t3 := v4
			v5 = v4 + i32(1)
			v6 = int32(uint32(v5) >> 3)
			v7 = v6 * i32(7)
			p4 := v7
			if uint32(v4) < uint32(i32(8)) {
				p4 = t3
			}
			v8 = p4
			if uint32(t2) <= uint32(int32(uint32(v8)>>1)) {
				goto l1
			}
			{
				{
					{
						v8 = v8 + i32(1)
						p5 := v3
						if uint32(v8) > uint32(v3) {
							p5 = v8
						}
						v3 = p5
						if uint32(v3) < uint32(i32(15)) {
							goto l2
						}
						if uint32(v3) > uint32(i32(0x1fffffff)) {
							m.fn27(i32(1271632), i32(57), i32(1271660))
							panic("unreachable")
						}
						t6 := int32(uint32(v3<<3) / uint32(i32(7)))
						v3 = i32_shr_u(i32(-1), int32(bits.LeadingZeros32(uint32(t6+i32(-1))))) + i32(1)
						goto l4
					}
				l2:
					p7 := v3&i32(8) + i32(8)
					if uint32(v3) < uint32(i32(4)) {
						p7 = i32(4)
					}
					v3 = p7
				}
			l4:
				v9 = int64(uint32(v3)) * i64(24)
				if int32(int64(uint64(v9)>>32)) != 0 {
					goto l5
				}
				v8 = v3 + i32(8)
				t8 := v8
				v10 = int32(v9)
				v6 = t8 + v10
				if uint32(v6) < uint32(v8) {
					goto l5
				}
				if uint32(v6) > uint32(i32(0x7ffffff8)) {
					goto l5
				}
				t9 := m.fn7(v6)
				v5 = t9
				if v5 != 0 {
					v6 = v5 + v10
					if v8 == 0 {
						goto l7
					}
					memory_fill(m.memory, uint32(v6), i32(255), uint32(v8))
				l7:
					v5 = v3 + i32(-1)
					p10 := int32(uint32(v3)>>3) * i32(7)
					if uint32(v3) < uint32(i32(9)) {
						p10 = v5
					}
					v7 = p10
					t11 := int32(load32(m.memory[uint32(v0):]))
					v11 = t11
					{
						if v2 == 0 {
							goto l8
						}
						t12 := int64(load64(m.memory[uint32(v11):]))
						v9 = (t12 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
						t13 := int64(load64(m.memory[int64(uint32(v1))+8:]))
						v12 = t13
						t14 := int64(load64(m.memory[uint32(v1):]))
						v13 = t14
						v8 = v11
						v14 = v2
						v3 = i32(0)
					l14:
						{
							if v9 != i64(0) {
								goto l9
							}
						l10:
							{
								v3 = v3 + i32(8)
								v8 = v8 + i32(8)
								t15 := int64(load64(m.memory[uint32(v8):]))
								v9 = t15 & i64(-0x7f7f7f7f7f7f7f80)
								if v9 == i64(-0x7f7f7f7f7f7f7f80) {
									goto l10
								}
							}
							v9 = v9 ^ i64(-0x7f7f7f7f7f7f7f80)
						l9:
							{
								t16 := v6
								t17 := v5
								t18 := v13
								t19 := v12
								t20 := v11
								v1 = int32(uint32(int64(bits.TrailingZeros64(uint64(v9))))>>3) + v3
								v10 = t20 + (i32(0)-v1)*i32(24)
								t21 := int32(load32(m.memory[uint32(v10+i32(-20)):]))
								t22 := int32(load32(m.memory[uint32(v10+i32(-16)):]))
								t23 := m.fn64(t18, t19, t21, t22)
								v15 = int32(t23)
								v10 = t17 & v15
								t24 := int64(load64(m.memory[uint32(t16+v10):]))
								v16 = t24 & i64(-0x7f7f7f7f7f7f7f80)
								if v16 != i64(0) {
									goto l11
								}
								v17 = i32(8)
							l12:
								{
									v10 = v10 + v17
									v17 = v17 + i32(8)
									t25 := v6
									v10 = v10 & v5
									t26 := int64(load64(m.memory[uint32(t25+v10):]))
									v16 = t26 & i64(-0x7f7f7f7f7f7f7f80)
									if v16 == 0 {
										goto l12
									}
								}
							}
						l11:
							v18 = v9 + i64(-1)
							{
								t27 := v6
								v10 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v16))))>>3) + v10) & v5
								t28 := int32(int8(m.memory[uint32(t27+v10)]))
								if t28 < i32(0) {
									goto l13
								}
								t29 := int64(load64(m.memory[uint32(v6):]))
								v10 = int32(uint32(int64(bits.TrailingZeros64(uint64(t29&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
							}
						l13:
							v9 = v18 & v9
							t30 := v6 + v10
							v15 = int32(uint32(v15) >> 25)
							m.memory[uint32(t30)] = byte(v15)
							m.memory[uint32(v6+(v10+i32(-8))&v5+i32(8))] = byte(v15)
							v10 = v6 + (v10^i32(-1))*i32(24)
							t31 := v10
							v1 = v11 + (v1^i32(-1))*i32(24)
							t32 := int64(load64(m.memory[int64(uint32(v1))+16:]))
							store64(m.memory[int64(uint32(t31))+16:], uint64(t32))
							t33 := int64(load64(m.memory[int64(uint32(v1))+8:]))
							store64(m.memory[int64(uint32(v10))+8:], uint64(t33))
							t34 := int64(load64(m.memory[uint32(v1):]))
							store64(m.memory[uint32(v10):], uint64(t34))
							v14 = v14 + i32(-1)
							if v14 != 0 {
								goto l14
							}
						}
					}
				l8:
					store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
					store32(m.memory[uint32(v0):], uint32(v6))
					store32(m.memory[int64(uint32(v0))+8:], uint32(v7-v2))
					if v4 == 0 {
						goto l15
					}
					t35 := v4
					v8 = (v4*i32(24) + i32(31)) & i32(-8)
					v3 = t35 + v8 + i32(9)
					if v3 == 0 {
						goto l15
					}
					v4 = v11 - v8
					t36 := int32(load32(m.memory[uint32(v4+i32(-4)):]))
					v8 = t36
					v6 = v8 & i32(-8)
					t37 := v6
					v8 = v8 & i32(3)
					p38 := i32(8)
					if v8 != 0 {
						p38 = i32(4)
					}
					if uint32(t37) < uint32(p38+v3) {
						m.fn3(i32(1274224), i32(46), i32(1274272))
						panic("unreachable")
					}
					if v8 == 0 {
						goto l17
					}
					if uint32(v6) > uint32(v3+i32(39)) {
						m.fn3(i32(1274288), i32(46), i32(1274336))
						panic("unreachable")
					}
				l17:
					m.fn1(v4)
					return i32(-1)
				}
				m.fn23(i32(8), v6)
				panic("unreachable")
			}
		}
	l1:
		{
			if v5 != 0 {
				goto l19
			}
			v3 = i32(0)
			goto l20
		l19:
			t39 := int32(load32(m.memory[uint32(v0):]))
			v8 = t39
			v3 = i32(0)
			{
				{
					t40 := v6
					var p41 int32
					if v5&i32(7) != i32(0) {
						p41 = 1
					}
					v6 = t40 + p41
					if v6 == i32(1) {
						goto l21
					}
					v11 = v6 & i32(1)
					v10 = v6 & i32(0x3ffffffe)
					v3 = i32(0)
				l22:
					{
						v6 = v8 + v3
						t42 := int64(load64(m.memory[uint32(v6):]))
						t43 := v6
						v9 = t42
						store64(m.memory[uint32(t43):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
						v6 = v6 + i32(8)
						t44 := int64(load64(m.memory[uint32(v6):]))
						t45 := v6
						v9 = t44
						store64(m.memory[uint32(t45):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
						v3 = v3 + i32(16)
						v10 = v10 + i32(-2)
						if v10 != 0 {
							goto l22
						}
					}
					if v11 == 0 {
						goto l23
					}
				}
			l21:
				v3 = v8 + v3
				t46 := int64(load64(m.memory[uint32(v3):]))
				t47 := v3
				v9 = t46
				store64(m.memory[uint32(t47):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
			}
		l23:
			{
				if uint32(v5) < uint32(i32(8)) {
					goto l24
				}
				t48 := int64(load64(m.memory[uint32(v8):]))
				store64(m.memory[uint32(v8+v5):], uint64(t48))
				goto l25
			}
		l24:
			if v5 == 0 {
				goto l25
			}
			memory_copy(m.memory, uint32(v8+i32(8)), uint32(v8), uint32(v5))
		l25:
			t49 := int64(load64(m.memory[int64(uint32(v1))+8:]))
			v16 = t49
			t50 := int64(load64(m.memory[uint32(v1):]))
			v18 = t50
			v6 = i32(0)
		l33:
			{
				t51 := v8
				v3 = v6
				v10 = t51 + v3
				t52 := int32(m.memory[uint32(v10)])
				if t52 != i32(128) {
					goto l26
				}
				v15 = v8 + (v3^i32(-1))*i32(24)
				v6 = v8 + (i32(0)-v3)*i32(24)
				v11 = v6 + i32(-16)
				v14 = v6 + i32(-20)
			l32:
				{
					t53 := int32(load32(m.memory[uint32(v14):]))
					t54 := int32(load32(m.memory[uint32(v11):]))
					t55 := m.fn64(v18, v16, t53, t54)
					t56 := v4
					v1 = int32(t55)
					v6 = t56 & v1
					v5 = v6
					{
						t57 := int64(load64(m.memory[uint32(v8+v6):]))
						v9 = t57 & i64(-0x7f7f7f7f7f7f7f80)
						if v9 != i64(0) {
							goto l27
						}
						v17 = i32(8)
						v5 = v6
					l28:
						{
							v5 = v5 + v17
							v17 = v17 + i32(8)
							t58 := v8
							v5 = v5 & v4
							t59 := int64(load64(m.memory[uint32(t58+v5):]))
							v9 = t59 & i64(-0x7f7f7f7f7f7f7f80)
							if v9 == 0 {
								goto l28
							}
						}
					}
				l27:
					{
						t60 := v8
						v5 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v9))))>>3) + v5) & v4
						t61 := int32(int8(m.memory[uint32(t60+v5)]))
						if t61 < i32(0) {
							goto l29
						}
						t62 := int64(load64(m.memory[uint32(v8):]))
						v5 = int32(uint32(int64(bits.TrailingZeros64(uint64(t62&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
					}
				l29:
					{
						if uint32((v5-v6^(v3-v6))&v4) < uint32(i32(8)) {
							goto l30
						}
						v6 = v8 + v5
						t63 := int32(m.memory[uint32(v6)])
						v17 = t63
						t64 := v6
						v1 = int32(uint32(v1) >> 25)
						m.memory[uint32(t64)] = byte(v1)
						m.memory[uint32(v8+(v5+i32(-8))&v4+i32(8))] = byte(v1)
						v6 = v8 + (v5^i32(-1))*i32(24)
						{
							if v17 != i32(255) {
								t68 := int64(load64(m.memory[uint32(v15):]))
								v9 = t68
								t69 := int64(load64(m.memory[uint32(v6):]))
								store64(m.memory[uint32(v15):], uint64(t69))
								store64(m.memory[uint32(v6):], uint64(v9))
								t70 := int64(load64(m.memory[int64(uint32(v6))+8:]))
								v9 = t70
								t71 := int64(load64(m.memory[int64(uint32(v15))+8:]))
								store64(m.memory[int64(uint32(v6))+8:], uint64(t71))
								store64(m.memory[int64(uint32(v15))+8:], uint64(v9))
								t72 := int32(load32(m.memory[int64(uint32(v15))+16:]))
								v5 = t72
								t73 := int32(load32(m.memory[int64(uint32(v6))+16:]))
								store32(m.memory[int64(uint32(v15))+16:], uint32(t73))
								t74 := int32(load32(m.memory[int64(uint32(v6))+20:]))
								v1 = t74
								t75 := int32(load32(m.memory[int64(uint32(v15))+20:]))
								store32(m.memory[int64(uint32(v6))+20:], uint32(t75))
								store32(m.memory[int64(uint32(v15))+20:], uint32(v1))
								store32(m.memory[int64(uint32(v6))+16:], uint32(v5))
								goto l32
							}
							m.memory[uint32(v10)] = byte(i32(255))
							m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(i32(255))
							t65 := int64(load64(m.memory[int64(uint32(v15))+16:]))
							store64(m.memory[int64(uint32(v6))+16:], uint64(t65))
							t66 := int64(load64(m.memory[int64(uint32(v15))+8:]))
							store64(m.memory[int64(uint32(v6))+8:], uint64(t66))
							t67 := int64(load64(m.memory[uint32(v15):]))
							store64(m.memory[uint32(v6):], uint64(t67))
							goto l26
						}
					}
				l30:
				}
				t76 := v10
				v6 = int32(uint32(v1) >> 25)
				m.memory[uint32(t76)] = byte(v6)
				m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(v6)
			}
		l26:
			v6 = v3 + i32(1)
			if v3 != v4 {
				goto l33
			}
			p77 := v7
			if uint32(v4) < uint32(i32(8)) {
				p77 = v4
			}
			v3 = p77
		}
	l20:
		store32(m.memory[int64(uint32(v0))+8:], uint32(v3-v2))
	l15:
		return i32(-1)
	}
l5:
	m.fn27(i32(1271632), i32(57), i32(1271660))
	panic("unreachable")
}
func (m *Module) fn80(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7, v8 int32
	var v9 int64
	var v10, v11, v12 int32
	var v13, v14 int64
	var v15, v16 int32
	var v17, v18 int64
	var v19, v20 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
		v2 = t0
		v3 = v2 + i32(1)
		if v3 == 0 {
			m.fn27(i32(1271632), i32(57), i32(1271660))
			panic("unreachable")
		}
		{
			t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t2 := v3
			v4 = t1
			t3 := v4
			v5 = v4 + i32(1)
			v6 = int32(uint32(v5) >> 3)
			v7 = v6 * i32(7)
			p4 := v7
			if uint32(v4) < uint32(i32(8)) {
				p4 = t3
			}
			v8 = p4
			if uint32(t2) <= uint32(int32(uint32(v8)>>1)) {
				goto l1
			}
			{
				{
					{
						v8 = v8 + i32(1)
						p5 := v3
						if uint32(v8) > uint32(v3) {
							p5 = v8
						}
						v3 = p5
						if uint32(v3) < uint32(i32(15)) {
							goto l2
						}
						if uint32(v3) > uint32(i32(0x1fffffff)) {
							m.fn27(i32(1271632), i32(57), i32(1271660))
							panic("unreachable")
						}
						t6 := int32(uint32(v3<<3) / uint32(i32(7)))
						v3 = i32_shr_u(i32(-1), int32(bits.LeadingZeros32(uint32(t6+i32(-1))))) + i32(1)
						goto l4
					}
				l2:
					p7 := v3&i32(8) + i32(8)
					if uint32(v3) < uint32(i32(4)) {
						p7 = i32(4)
					}
					v3 = p7
				}
			l4:
				v9 = int64(uint32(v3)) * i64(488)
				if int32(int64(uint64(v9)>>32)) != 0 {
					goto l5
				}
				v8 = v3 + i32(8)
				t8 := v8
				v10 = int32(v9)
				v6 = t8 + v10
				if uint32(v6) < uint32(v8) {
					goto l5
				}
				if uint32(v6) > uint32(i32(0x7ffffff8)) {
					goto l5
				}
				t9 := m.fn7(v6)
				v5 = t9
				if v5 != 0 {
					v6 = v5 + v10
					if v8 == 0 {
						goto l7
					}
					memory_fill(m.memory, uint32(v6), i32(255), uint32(v8))
				l7:
					v5 = v3 + i32(-1)
					p10 := int32(uint32(v3)>>3) * i32(7)
					if uint32(v3) < uint32(i32(9)) {
						p10 = v5
					}
					v11 = p10
					t11 := int32(load32(m.memory[uint32(v0):]))
					v12 = t11
					{
						if v2 == 0 {
							goto l8
						}
						t12 := int64(load64(m.memory[uint32(v12):]))
						v9 = (t12 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
						t13 := int64(load64(m.memory[int64(uint32(v1))+8:]))
						v13 = t13
						t14 := int64(load64(m.memory[uint32(v1):]))
						v14 = t14
						v8 = v12
						v1 = v2
						v3 = i32(0)
					l14:
						{
							if v9 != i64(0) {
								goto l9
							}
						l10:
							{
								v3 = v3 + i32(8)
								v8 = v8 + i32(8)
								t15 := int64(load64(m.memory[uint32(v8):]))
								v9 = t15 & i64(-0x7f7f7f7f7f7f7f80)
								if v9 == i64(-0x7f7f7f7f7f7f7f80) {
									goto l10
								}
							}
							v9 = v9 ^ i64(-0x7f7f7f7f7f7f7f80)
						l9:
							{
								t16 := v6
								t17 := v5
								t18 := v14
								t19 := v13
								t20 := v12
								v15 = int32(uint32(int64(bits.TrailingZeros64(uint64(v9))))>>3) + v3
								v10 = t20 + (i32(0)-v15)*i32(488)
								t21 := int32(load32(m.memory[uint32(v10+i32(-488)):]))
								t22 := int32(load32(m.memory[uint32(v10+i32(-484)):]))
								t23 := m.fn81(t18, t19, t21, t22)
								v16 = int32(t23)
								v10 = t17 & v16
								t24 := int64(load64(m.memory[uint32(t16+v10):]))
								v17 = t24 & i64(-0x7f7f7f7f7f7f7f80)
								if v17 != i64(0) {
									goto l11
								}
								v7 = i32(8)
							l12:
								{
									v10 = v10 + v7
									v7 = v7 + i32(8)
									t25 := v6
									v10 = v10 & v5
									t26 := int64(load64(m.memory[uint32(t25+v10):]))
									v17 = t26 & i64(-0x7f7f7f7f7f7f7f80)
									if v17 == 0 {
										goto l12
									}
								}
							}
						l11:
							v18 = v9 + i64(-1)
							{
								t27 := v6
								v10 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v17))))>>3) + v10) & v5
								t28 := int32(int8(m.memory[uint32(t27+v10)]))
								if t28 < i32(0) {
									goto l13
								}
								t29 := int64(load64(m.memory[uint32(v6):]))
								v10 = int32(uint32(int64(bits.TrailingZeros64(uint64(t29&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
							}
						l13:
							v9 = v18 & v9
							t30 := v6 + v10
							v16 = int32(uint32(v16) >> 25)
							m.memory[uint32(t30)] = byte(v16)
							m.memory[uint32(v6+(v10+i32(-8))&v5+i32(8))] = byte(v16)
							memory_copy(m.memory, uint32(v6+(v10^i32(-1))*i32(488)), uint32(v12+(v15^i32(-1))*i32(488)), uint32(i32(488)))
							v1 = v1 + i32(-1)
							if v1 != 0 {
								goto l14
							}
						}
					}
				l8:
					store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
					store32(m.memory[uint32(v0):], uint32(v6))
					store32(m.memory[int64(uint32(v0))+8:], uint32(v11-v2))
					if v4 == 0 {
						goto l15
					}
					t31 := v4
					v8 = (v4*i32(488) + i32(495)) & i32(-8)
					v3 = t31 + v8 + i32(9)
					if v3 == 0 {
						goto l15
					}
					v4 = v12 - v8
					t32 := int32(load32(m.memory[uint32(v4+i32(-4)):]))
					v8 = t32
					v6 = v8 & i32(-8)
					t33 := v6
					v8 = v8 & i32(3)
					p34 := i32(8)
					if v8 != 0 {
						p34 = i32(4)
					}
					if uint32(t33) < uint32(p34+v3) {
						m.fn3(i32(1274224), i32(46), i32(1274272))
						panic("unreachable")
					}
					if v8 == 0 {
						goto l17
					}
					if uint32(v6) > uint32(v3+i32(39)) {
						m.fn3(i32(1274288), i32(46), i32(1274336))
						panic("unreachable")
					}
				l17:
					m.fn1(v4)
					return i32(-1)
				}
				m.fn23(i32(8), v6)
				panic("unreachable")
			}
		}
	l1:
		{
			if v5 != 0 {
				goto l19
			}
			v3 = i32(0)
			goto l20
		l19:
			t35 := int32(load32(m.memory[uint32(v0):]))
			v8 = t35
			v3 = i32(0)
			{
				{
					t36 := v6
					var p37 int32
					if v5&i32(7) != i32(0) {
						p37 = 1
					}
					v6 = t36 + p37
					if v6 == i32(1) {
						goto l21
					}
					v12 = v6 & i32(1)
					v10 = v6 & i32(0x3ffffffe)
					v3 = i32(0)
				l22:
					{
						v6 = v8 + v3
						t38 := int64(load64(m.memory[uint32(v6):]))
						t39 := v6
						v9 = t38
						store64(m.memory[uint32(t39):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
						v6 = v6 + i32(8)
						t40 := int64(load64(m.memory[uint32(v6):]))
						t41 := v6
						v9 = t40
						store64(m.memory[uint32(t41):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
						v3 = v3 + i32(16)
						v10 = v10 + i32(-2)
						if v10 != 0 {
							goto l22
						}
					}
					if v12 == 0 {
						goto l23
					}
				}
			l21:
				v3 = v8 + v3
				t42 := int64(load64(m.memory[uint32(v3):]))
				t43 := v3
				v9 = t42
				store64(m.memory[uint32(t43):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
			}
		l23:
			{
				if uint32(v5) < uint32(i32(8)) {
					goto l24
				}
				t44 := int64(load64(m.memory[uint32(v8):]))
				store64(m.memory[uint32(v8+v5):], uint64(t44))
				goto l25
			}
		l24:
			if v5 == 0 {
				goto l25
			}
			memory_copy(m.memory, uint32(v8+i32(8)), uint32(v8), uint32(v5))
		l25:
			t45 := int64(load64(m.memory[int64(uint32(v1))+8:]))
			v17 = t45
			t46 := int64(load64(m.memory[uint32(v1):]))
			v18 = t46
			v6 = v8
			v10 = i32(0)
		l34:
			{
				t47 := v8
				v3 = v10
				v5 = t47 + v3
				t48 := int32(m.memory[uint32(v5)])
				if t48 != i32(128) {
					goto l26
				}
				v11 = v8 + (v3^i32(-1))*i32(488)
				v10 = v8 + (i32(0)-v3)*i32(488)
				v15 = v10 + i32(-484)
				v16 = v10 + i32(-488)
				{
				l33:
					{
						t49 := int32(load32(m.memory[uint32(v16):]))
						t50 := int32(load32(m.memory[uint32(v15):]))
						t51 := m.fn81(v18, v17, t49, t50)
						t52 := v4
						v1 = int32(t51)
						v10 = t52 & v1
						v12 = v10
						{
							t53 := int64(load64(m.memory[uint32(v8+v10):]))
							v9 = t53 & i64(-0x7f7f7f7f7f7f7f80)
							if v9 != i64(0) {
								goto l27
							}
							v19 = i32(8)
							v12 = v10
						l28:
							{
								v12 = v12 + v19
								v19 = v19 + i32(8)
								t54 := v8
								v12 = v12 & v4
								t55 := int64(load64(m.memory[uint32(t54+v12):]))
								v9 = t55 & i64(-0x7f7f7f7f7f7f7f80)
								if v9 == 0 {
									goto l28
								}
							}
						}
					l27:
						{
							t56 := v8
							v12 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v9))))>>3) + v12) & v4
							t57 := int32(int8(m.memory[uint32(t56+v12)]))
							if t57 < i32(0) {
								goto l29
							}
							t58 := int64(load64(m.memory[uint32(v8):]))
							v12 = int32(uint32(int64(bits.TrailingZeros64(uint64(t58&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
						}
					l29:
						{
							if uint32((v12-v10^(v3-v10))&v4) < uint32(i32(8)) {
								goto l30
							}
							v10 = v8 + v12
							t59 := int32(m.memory[uint32(v10)])
							v19 = t59
							t60 := v10
							v1 = int32(uint32(v1) >> 25)
							m.memory[uint32(t60)] = byte(v1)
							m.memory[uint32(v8+(v12+i32(-8))&v4+i32(8))] = byte(v1)
							if v19 == i32(255) {
								goto l31
							}
							v10 = i32(-488)
							v20 = v8 + v12*i32(-488)
						l32:
							{
								v12 = v20 + v10
								t61 := int32(load32(m.memory[uint32(v12):]))
								v19 = t61
								t62 := v12
								v1 = v6 + v10
								t63 := int32(load32(m.memory[uint32(v1):]))
								store32(m.memory[uint32(t62):], uint32(t63))
								store32(m.memory[uint32(v1):], uint32(v19))
								v1 = v1 + i32(4)
								t64 := int32(load32(m.memory[uint32(v1):]))
								v19 = t64
								t65 := v1
								v12 = v12 + i32(4)
								t66 := int32(load32(m.memory[uint32(v12):]))
								store32(m.memory[uint32(t65):], uint32(t66))
								store32(m.memory[uint32(v12):], uint32(v19))
								v10 = v10 + i32(8)
								if v10 != 0 {
									goto l32
								}
								goto l33
							}
						}
					l30:
					}
					t67 := v5
					v10 = int32(uint32(v1) >> 25)
					m.memory[uint32(t67)] = byte(v10)
					m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(v10)
					goto l26
				}
			l31:
				m.memory[uint32(v5)] = byte(i32(255))
				m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(i32(255))
				memory_copy(m.memory, uint32(v8+(v12^i32(-1))*i32(488)), uint32(v11), uint32(i32(488)))
			}
		l26:
			v10 = v3 + i32(1)
			v6 = v6 + i32(-488)
			if v3 != v4 {
				goto l34
			}
			p68 := v7
			if uint32(v4) < uint32(i32(8)) {
				p68 = v4
			}
			v3 = p68
		}
	l20:
		store32(m.memory[int64(uint32(v0))+8:], uint32(v3-v2))
	l15:
		return i32(-1)
	l5:
		m.fn27(i32(1271632), i32(57), i32(1271660))
		panic("unreachable")
	}
}
func (m *Module) fn81(v0, v1 int64, v2, v3 int32) int64 {
	var v4 int32
	var v5, v6, v7, v8 int64
	t0 := m.g0
	v4 = t0 - i32(80)
	m.g0 = v4
	store64(m.memory[int64(uint32(v4))+56:], uint64(i64(0)))
	store64(m.memory[int64(uint32(v4))+64:], uint64(i64(0)))
	store64(m.memory[int64(uint32(v4))+48:], uint64(v1))
	store64(m.memory[int64(uint32(v4))+32:], uint64(v1^i64(8387220255154660723)))
	store64(m.memory[int64(uint32(v4))+24:], uint64(v1^i64(7237128888997146477)))
	store64(m.memory[int64(uint32(v4))+40:], uint64(v0))
	store64(m.memory[int64(uint32(v4))+16:], uint64(v0^i64(0x6c7967656e657261)))
	store64(m.memory[int64(uint32(v4))+8:], uint64(v0^i64(8317987319222330741)))
	m.fn58(v4+i32(8), v2, v3)
	m.memory[int64(uint32(v4))+79] = byte(i32(255))
	m.fn58(v4+i32(8), v4+i32(79), i32(1))
	t1 := int64(load64(m.memory[int64(uint32(v4))+8:]))
	v0 = t1
	t2 := int64(load64(m.memory[int64(uint32(v4))+24:]))
	v1 = t2
	t3 := int64(load32(m.memory[int64(uint32(v4))+64:]))
	v5 = t3
	t4 := int64(load64(m.memory[int64(uint32(v4))+56:]))
	v6 = t4
	t5 := int64(load64(m.memory[int64(uint32(v4))+32:]))
	v7 = t5
	t6 := int64(load64(m.memory[int64(uint32(v4))+16:]))
	v8 = t6
	m.g0 = v4 + i32(80)
	t7 := v7
	v5 = v6 | v5<<56
	v6 = t7 ^ v5
	t8 := i64_rotl(v6, i64(16))
	v6 = v6 + v8
	v7 = t8 ^ v6
	t9 := i64_rotl(v7, i64(21))
	t10 := v7
	v0 = v1 + v0
	v7 = t10 + i64_rotl(v0, i64(32))
	v8 = t9 ^ v7
	t11 := i64_rotl(v8, i64(16))
	t12 := v8
	t13 := v6
	v1 = i64_rotl(v1, i64(13)) ^ v0
	v0 = t13 + v1
	v6 = t12 + (i64_rotl(v0, i64(32)) ^ i64(255))
	v8 = t11 ^ v6
	t14 := i64_rotl(v8, i64(21))
	t15 := v8
	t16 := v7 ^ v5
	v1 = v0 ^ i64_rotl(v1, i64(17))
	v0 = t16 + v1
	v5 = t15 + i64_rotl(v0, i64(32))
	v7 = t14 ^ v5
	t17 := i64_rotl(v7, i64(16))
	t18 := v7
	v1 = v0 ^ i64_rotl(v1, i64(13))
	v0 = v1 + v6
	v6 = t18 + i64_rotl(v0, i64(32))
	v7 = t17 ^ v6
	t19 := i64_rotl(v7, i64(21))
	t20 := v7
	v1 = v0 ^ i64_rotl(v1, i64(17))
	v0 = v1 + v5
	v5 = t20 + i64_rotl(v0, i64(32))
	v7 = t19 ^ v5
	t21 := i64_rotl(v7, i64(16))
	t22 := v7
	v1 = i64_rotl(v1, i64(13)) ^ v0
	v0 = v1 + v6
	v6 = t22 + i64_rotl(v0, i64(32))
	t23 := i64_rotl(t21^v6, i64(21))
	v1 = i64_rotl(v1, i64(17)) ^ v0
	v1 = i64_rotl(v1, i64(13)) ^ (v1 + v5)
	t24 := t23 ^ i64_rotl(v1, i64(17))
	v1 = v1 + v6
	return t24 ^ i64_rotl(v1, i64(32)) ^ v1
}
func (m *Module) fn82(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7, v8 int32
	var v9 int64
	var v10, v11 int32
	var v12, v13 int64
	var v14, v15 int32
	var v16 int64
	var v17 int32
	var v18 int64
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
		v2 = t0
		v3 = v2 + i32(1)
		if v3 == 0 {
			m.fn27(i32(1271632), i32(57), i32(1271660))
			panic("unreachable")
		}
		{
			t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t2 := v3
			v4 = t1
			t3 := v4
			v5 = v4 + i32(1)
			v6 = int32(uint32(v5) >> 3)
			v7 = v6 * i32(7)
			p4 := v7
			if uint32(v4) < uint32(i32(8)) {
				p4 = t3
			}
			v8 = p4
			if uint32(t2) <= uint32(int32(uint32(v8)>>1)) {
				goto l1
			}
			{
				{
					{
						v8 = v8 + i32(1)
						p5 := v3
						if uint32(v8) > uint32(v3) {
							p5 = v8
						}
						v3 = p5
						if uint32(v3) < uint32(i32(15)) {
							goto l2
						}
						if uint32(v3) > uint32(i32(0x1fffffff)) {
							m.fn27(i32(1271632), i32(57), i32(1271660))
							panic("unreachable")
						}
						t6 := int32(uint32(v3<<3) / uint32(i32(7)))
						v3 = i32_shr_u(i32(-1), int32(bits.LeadingZeros32(uint32(t6+i32(-1))))) + i32(1)
						goto l4
					}
				l2:
					p7 := v3&i32(8) + i32(8)
					if uint32(v3) < uint32(i32(4)) {
						p7 = i32(4)
					}
					v3 = p7
				}
			l4:
				v9 = int64(uint32(v3)) * i64(12)
				if int32(int64(uint64(v9)>>32)) != 0 {
					goto l5
				}
				v6 = int32(v9)
				if uint32(v6) > uint32(i32(-8)) {
					goto l5
				}
				v8 = v3 + i32(8)
				t8 := v8
				v10 = (v6 + i32(7)) & i32(-8)
				v6 = t8 + v10
				if uint32(v6) < uint32(v8) {
					goto l5
				}
				if uint32(v6) > uint32(i32(0x7ffffff8)) {
					goto l5
				}
				t9 := m.fn7(v6)
				v5 = t9
				if v5 != 0 {
					v6 = v5 + v10
					if v8 == 0 {
						goto l7
					}
					memory_fill(m.memory, uint32(v6), i32(255), uint32(v8))
				l7:
					v5 = v3 + i32(-1)
					p10 := int32(uint32(v3)>>3) * i32(7)
					if uint32(v3) < uint32(i32(9)) {
						p10 = v5
					}
					v7 = p10
					t11 := int32(load32(m.memory[uint32(v0):]))
					v11 = t11
					{
						if v2 == 0 {
							goto l8
						}
						t12 := int64(load64(m.memory[uint32(v11):]))
						v9 = (t12 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
						t13 := int64(load64(m.memory[int64(uint32(v1))+8:]))
						v12 = t13
						t14 := int64(load64(m.memory[uint32(v1):]))
						v13 = t14
						v8 = v11
						v1 = v2
						v3 = i32(0)
					l14:
						{
							if v9 != i64(0) {
								goto l9
							}
						l10:
							{
								v3 = v3 + i32(8)
								v8 = v8 + i32(8)
								t15 := int64(load64(m.memory[uint32(v8):]))
								v9 = t15 & i64(-0x7f7f7f7f7f7f7f80)
								if v9 == i64(-0x7f7f7f7f7f7f7f80) {
									goto l10
								}
							}
							v9 = v9 ^ i64(-0x7f7f7f7f7f7f7f80)
						l9:
							{
								t16 := v6
								t17 := v5
								t18 := v13
								t19 := v12
								t20 := v11
								v14 = int32(uint32(int64(bits.TrailingZeros64(uint64(v9))))>>3) + v3
								v10 = t20 + (i32(0)-v14)*i32(12)
								t21 := int32(load32(m.memory[uint32(v10+i32(-12)):]))
								t22 := int32(load32(m.memory[uint32(v10+i32(-8)):]))
								t23 := m.fn81(t18, t19, t21, t22)
								v15 = int32(t23)
								v10 = t17 & v15
								t24 := int64(load64(m.memory[uint32(t16+v10):]))
								v16 = t24 & i64(-0x7f7f7f7f7f7f7f80)
								if v16 != i64(0) {
									goto l11
								}
								v17 = i32(8)
							l12:
								{
									v10 = v10 + v17
									v17 = v17 + i32(8)
									t25 := v6
									v10 = v10 & v5
									t26 := int64(load64(m.memory[uint32(t25+v10):]))
									v16 = t26 & i64(-0x7f7f7f7f7f7f7f80)
									if v16 == 0 {
										goto l12
									}
								}
							}
						l11:
							v18 = v9 + i64(-1)
							{
								t27 := v6
								v10 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v16))))>>3) + v10) & v5
								t28 := int32(int8(m.memory[uint32(t27+v10)]))
								if t28 < i32(0) {
									goto l13
								}
								t29 := int64(load64(m.memory[uint32(v6):]))
								v10 = int32(uint32(int64(bits.TrailingZeros64(uint64(t29&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
							}
						l13:
							v9 = v18 & v9
							t30 := v6 + v10
							v15 = int32(uint32(v15) >> 25)
							m.memory[uint32(t30)] = byte(v15)
							m.memory[uint32(v6+(v10+i32(-8))&v5+i32(8))] = byte(v15)
							v10 = v6 + (v10^i32(-1))*i32(12)
							t31 := v10
							v14 = v11 + (v14^i32(-1))*i32(12)
							t32 := int32(load32(m.memory[int64(uint32(v14))+8:]))
							store32(m.memory[int64(uint32(t31))+8:], uint32(t32))
							t33 := int64(load64(m.memory[uint32(v14):]))
							store64(m.memory[uint32(v10):], uint64(t33))
							v1 = v1 + i32(-1)
							if v1 != 0 {
								goto l14
							}
						}
					}
				l8:
					store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
					store32(m.memory[uint32(v0):], uint32(v6))
					store32(m.memory[int64(uint32(v0))+8:], uint32(v7-v2))
					if v4 == 0 {
						goto l15
					}
					t34 := v4
					v8 = (v4*i32(12) + i32(19)) & i32(-8)
					v3 = t34 + v8 + i32(9)
					if v3 == 0 {
						goto l15
					}
					v4 = v11 - v8
					t35 := int32(load32(m.memory[uint32(v4+i32(-4)):]))
					v8 = t35
					v6 = v8 & i32(-8)
					t36 := v6
					v8 = v8 & i32(3)
					p37 := i32(8)
					if v8 != 0 {
						p37 = i32(4)
					}
					if uint32(t36) < uint32(p37+v3) {
						m.fn3(i32(1274224), i32(46), i32(1274272))
						panic("unreachable")
					}
					if v8 == 0 {
						goto l17
					}
					if uint32(v6) > uint32(v3+i32(39)) {
						m.fn3(i32(1274288), i32(46), i32(1274336))
						panic("unreachable")
					}
				l17:
					m.fn1(v4)
					return i32(-1)
				}
				m.fn23(i32(8), v6)
				panic("unreachable")
			}
		}
	l1:
		{
			if v5 != 0 {
				goto l19
			}
			v3 = i32(0)
			goto l20
		l19:
			t38 := int32(load32(m.memory[uint32(v0):]))
			v8 = t38
			v3 = i32(0)
			{
				{
					t39 := v6
					var p40 int32
					if v5&i32(7) != i32(0) {
						p40 = 1
					}
					v6 = t39 + p40
					if v6 == i32(1) {
						goto l21
					}
					v11 = v6 & i32(1)
					v10 = v6 & i32(0x3ffffffe)
					v3 = i32(0)
				l22:
					{
						v6 = v8 + v3
						t41 := int64(load64(m.memory[uint32(v6):]))
						t42 := v6
						v9 = t41
						store64(m.memory[uint32(t42):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
						v6 = v6 + i32(8)
						t43 := int64(load64(m.memory[uint32(v6):]))
						t44 := v6
						v9 = t43
						store64(m.memory[uint32(t44):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
						v3 = v3 + i32(16)
						v10 = v10 + i32(-2)
						if v10 != 0 {
							goto l22
						}
					}
					if v11 == 0 {
						goto l23
					}
				}
			l21:
				v3 = v8 + v3
				t45 := int64(load64(m.memory[uint32(v3):]))
				t46 := v3
				v9 = t45
				store64(m.memory[uint32(t46):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
			}
		l23:
			{
				if uint32(v5) < uint32(i32(8)) {
					goto l24
				}
				t47 := int64(load64(m.memory[uint32(v8):]))
				store64(m.memory[uint32(v8+v5):], uint64(t47))
				goto l25
			}
		l24:
			if v5 == 0 {
				goto l25
			}
			memory_copy(m.memory, uint32(v8+i32(8)), uint32(v8), uint32(v5))
		l25:
			t48 := int64(load64(m.memory[int64(uint32(v1))+8:]))
			v16 = t48
			t49 := int64(load64(m.memory[uint32(v1):]))
			v18 = t49
			v6 = i32(0)
		l33:
			{
				t50 := v8
				v3 = v6
				v10 = t50 + v3
				t51 := int32(m.memory[uint32(v10)])
				if t51 != i32(128) {
					goto l26
				}
				v15 = v8 + (v3^i32(-1))*i32(12)
				v6 = v8 + (i32(0)-v3)*i32(12)
				v1 = v6 + i32(-8)
				v14 = v6 + i32(-12)
			l32:
				{
					t52 := int32(load32(m.memory[uint32(v14):]))
					t53 := int32(load32(m.memory[uint32(v1):]))
					t54 := m.fn81(v18, v16, t52, t53)
					t55 := v4
					v11 = int32(t54)
					v6 = t55 & v11
					v5 = v6
					{
						t56 := int64(load64(m.memory[uint32(v8+v6):]))
						v9 = t56 & i64(-0x7f7f7f7f7f7f7f80)
						if v9 != i64(0) {
							goto l27
						}
						v17 = i32(8)
						v5 = v6
					l28:
						{
							v5 = v5 + v17
							v17 = v17 + i32(8)
							t57 := v8
							v5 = v5 & v4
							t58 := int64(load64(m.memory[uint32(t57+v5):]))
							v9 = t58 & i64(-0x7f7f7f7f7f7f7f80)
							if v9 == 0 {
								goto l28
							}
						}
					}
				l27:
					{
						t59 := v8
						v5 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v9))))>>3) + v5) & v4
						t60 := int32(int8(m.memory[uint32(t59+v5)]))
						if t60 < i32(0) {
							goto l29
						}
						t61 := int64(load64(m.memory[uint32(v8):]))
						v5 = int32(uint32(int64(bits.TrailingZeros64(uint64(t61&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
					}
				l29:
					{
						if uint32((v5-v6^(v3-v6))&v4) < uint32(i32(8)) {
							goto l30
						}
						v6 = v8 + v5
						t62 := int32(m.memory[uint32(v6)])
						v17 = t62
						t63 := v6
						v11 = int32(uint32(v11) >> 25)
						m.memory[uint32(t63)] = byte(v11)
						m.memory[uint32(v8+(v5+i32(-8))&v4+i32(8))] = byte(v11)
						v6 = v8 + (v5^i32(-1))*i32(12)
						{
							if v17 != i32(255) {
								t66 := int32(load32(m.memory[uint32(v15):]))
								v5 = t66
								t67 := int32(load32(m.memory[uint32(v6):]))
								store32(m.memory[uint32(v15):], uint32(t67))
								store32(m.memory[uint32(v6):], uint32(v5))
								t68 := int32(load32(m.memory[int64(uint32(v6))+4:]))
								v5 = t68
								t69 := int32(load32(m.memory[int64(uint32(v15))+4:]))
								store32(m.memory[int64(uint32(v6))+4:], uint32(t69))
								store32(m.memory[int64(uint32(v15))+4:], uint32(v5))
								t70 := int32(load32(m.memory[int64(uint32(v15))+8:]))
								v5 = t70
								t71 := int32(load32(m.memory[int64(uint32(v6))+8:]))
								store32(m.memory[int64(uint32(v15))+8:], uint32(t71))
								store32(m.memory[int64(uint32(v6))+8:], uint32(v5))
								goto l32
							}
							m.memory[uint32(v10)] = byte(i32(255))
							m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(i32(255))
							t64 := int32(load32(m.memory[int64(uint32(v15))+8:]))
							store32(m.memory[int64(uint32(v6))+8:], uint32(t64))
							t65 := int64(load64(m.memory[uint32(v15):]))
							store64(m.memory[uint32(v6):], uint64(t65))
							goto l26
						}
					}
				l30:
				}
				t72 := v10
				v6 = int32(uint32(v11) >> 25)
				m.memory[uint32(t72)] = byte(v6)
				m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(v6)
			}
		l26:
			v6 = v3 + i32(1)
			if v3 != v4 {
				goto l33
			}
			p73 := v7
			if uint32(v4) < uint32(i32(8)) {
				p73 = v4
			}
			v3 = p73
		}
	l20:
		store32(m.memory[int64(uint32(v0))+8:], uint32(v3-v2))
	l15:
		return i32(-1)
	}
l5:
	m.fn27(i32(1271632), i32(57), i32(1271660))
	panic("unreachable")
}
func (m *Module) fn83(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7, v8 int32
	var v9 int64
	var v10, v11 int32
	var v12, v13 int64
	var v14, v15 int32
	var v16 int64
	var v17 int32
	var v18 int64
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
		v2 = t0
		v3 = v2 + i32(1)
		if v3 == 0 {
			m.fn27(i32(1271632), i32(57), i32(1271660))
			panic("unreachable")
		}
		{
			t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t2 := v3
			v4 = t1
			t3 := v4
			v5 = v4 + i32(1)
			v6 = int32(uint32(v5) >> 3)
			v7 = v6 * i32(7)
			p4 := v7
			if uint32(v4) < uint32(i32(8)) {
				p4 = t3
			}
			v8 = p4
			if uint32(t2) <= uint32(int32(uint32(v8)>>1)) {
				goto l1
			}
			{
				{
					{
						v8 = v8 + i32(1)
						p5 := v3
						if uint32(v8) > uint32(v3) {
							p5 = v8
						}
						v3 = p5
						if uint32(v3) < uint32(i32(15)) {
							goto l2
						}
						if uint32(v3) > uint32(i32(0x1fffffff)) {
							m.fn27(i32(1271632), i32(57), i32(1271660))
							panic("unreachable")
						}
						t6 := int32(uint32(v3<<3) / uint32(i32(7)))
						v3 = i32_shr_u(i32(-1), int32(bits.LeadingZeros32(uint32(t6+i32(-1))))) + i32(1)
						goto l4
					}
				l2:
					p7 := v3&i32(8) + i32(8)
					if uint32(v3) < uint32(i32(4)) {
						p7 = i32(4)
					}
					v3 = p7
				}
			l4:
				v9 = int64(uint32(v3)) * i64(20)
				if int32(int64(uint64(v9)>>32)) != 0 {
					goto l5
				}
				v6 = int32(v9)
				if uint32(v6) > uint32(i32(-8)) {
					goto l5
				}
				v8 = v3 + i32(8)
				t8 := v8
				v10 = (v6 + i32(7)) & i32(-8)
				v6 = t8 + v10
				if uint32(v6) < uint32(v8) {
					goto l5
				}
				if uint32(v6) > uint32(i32(0x7ffffff8)) {
					goto l5
				}
				t9 := m.fn7(v6)
				v5 = t9
				if v5 != 0 {
					v6 = v5 + v10
					if v8 == 0 {
						goto l7
					}
					memory_fill(m.memory, uint32(v6), i32(255), uint32(v8))
				l7:
					v5 = v3 + i32(-1)
					p10 := int32(uint32(v3)>>3) * i32(7)
					if uint32(v3) < uint32(i32(9)) {
						p10 = v5
					}
					v7 = p10
					t11 := int32(load32(m.memory[uint32(v0):]))
					v11 = t11
					{
						if v2 == 0 {
							goto l8
						}
						t12 := int64(load64(m.memory[uint32(v11):]))
						v9 = (t12 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
						t13 := int64(load64(m.memory[int64(uint32(v1))+8:]))
						v12 = t13
						t14 := int64(load64(m.memory[uint32(v1):]))
						v13 = t14
						v8 = v11
						v14 = v2
						v3 = i32(0)
					l14:
						{
							if v9 != i64(0) {
								goto l9
							}
						l10:
							{
								v3 = v3 + i32(8)
								v8 = v8 + i32(8)
								t15 := int64(load64(m.memory[uint32(v8):]))
								v9 = t15 & i64(-0x7f7f7f7f7f7f7f80)
								if v9 == i64(-0x7f7f7f7f7f7f7f80) {
									goto l10
								}
							}
							v9 = v9 ^ i64(-0x7f7f7f7f7f7f7f80)
						l9:
							{
								t16 := v6
								t17 := v5
								t18 := v13
								t19 := v12
								t20 := v11
								v1 = int32(uint32(int64(bits.TrailingZeros64(uint64(v9))))>>3) + v3
								v10 = t20 + (i32(0)-v1)*i32(20)
								t21 := int32(load32(m.memory[uint32(v10+i32(-20)):]))
								t22 := int32(load32(m.memory[uint32(v10+i32(-16)):]))
								t23 := m.fn81(t18, t19, t21, t22)
								v15 = int32(t23)
								v10 = t17 & v15
								t24 := int64(load64(m.memory[uint32(t16+v10):]))
								v16 = t24 & i64(-0x7f7f7f7f7f7f7f80)
								if v16 != i64(0) {
									goto l11
								}
								v17 = i32(8)
							l12:
								{
									v10 = v10 + v17
									v17 = v17 + i32(8)
									t25 := v6
									v10 = v10 & v5
									t26 := int64(load64(m.memory[uint32(t25+v10):]))
									v16 = t26 & i64(-0x7f7f7f7f7f7f7f80)
									if v16 == 0 {
										goto l12
									}
								}
							}
						l11:
							v18 = v9 + i64(-1)
							{
								t27 := v6
								v10 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v16))))>>3) + v10) & v5
								t28 := int32(int8(m.memory[uint32(t27+v10)]))
								if t28 < i32(0) {
									goto l13
								}
								t29 := int64(load64(m.memory[uint32(v6):]))
								v10 = int32(uint32(int64(bits.TrailingZeros64(uint64(t29&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
							}
						l13:
							v9 = v18 & v9
							t30 := v6 + v10
							v15 = int32(uint32(v15) >> 25)
							m.memory[uint32(t30)] = byte(v15)
							m.memory[uint32(v6+(v10+i32(-8))&v5+i32(8))] = byte(v15)
							v10 = v6 + (v10^i32(-1))*i32(20)
							t31 := v10
							v1 = v11 + (v1^i32(-1))*i32(20)
							t32 := int32(load32(m.memory[int64(uint32(v1))+16:]))
							store32(m.memory[int64(uint32(t31))+16:], uint32(t32))
							t33 := int64(load64(m.memory[int64(uint32(v1))+8:]))
							store64(m.memory[int64(uint32(v10))+8:], uint64(t33))
							t34 := int64(load64(m.memory[uint32(v1):]))
							store64(m.memory[uint32(v10):], uint64(t34))
							v14 = v14 + i32(-1)
							if v14 != 0 {
								goto l14
							}
						}
					}
				l8:
					store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
					store32(m.memory[uint32(v0):], uint32(v6))
					store32(m.memory[int64(uint32(v0))+8:], uint32(v7-v2))
					if v4 == 0 {
						goto l15
					}
					t35 := v4
					v8 = (v4*i32(20) + i32(27)) & i32(-8)
					v3 = t35 + v8 + i32(9)
					if v3 == 0 {
						goto l15
					}
					v4 = v11 - v8
					t36 := int32(load32(m.memory[uint32(v4+i32(-4)):]))
					v8 = t36
					v6 = v8 & i32(-8)
					t37 := v6
					v8 = v8 & i32(3)
					p38 := i32(8)
					if v8 != 0 {
						p38 = i32(4)
					}
					if uint32(t37) < uint32(p38+v3) {
						m.fn3(i32(1274224), i32(46), i32(1274272))
						panic("unreachable")
					}
					if v8 == 0 {
						goto l17
					}
					if uint32(v6) > uint32(v3+i32(39)) {
						m.fn3(i32(1274288), i32(46), i32(1274336))
						panic("unreachable")
					}
				l17:
					m.fn1(v4)
					return i32(-1)
				}
				m.fn23(i32(8), v6)
				panic("unreachable")
			}
		}
	l1:
		{
			if v5 != 0 {
				goto l19
			}
			v3 = i32(0)
			goto l20
		l19:
			t39 := int32(load32(m.memory[uint32(v0):]))
			v8 = t39
			v3 = i32(0)
			{
				{
					t40 := v6
					var p41 int32
					if v5&i32(7) != i32(0) {
						p41 = 1
					}
					v6 = t40 + p41
					if v6 == i32(1) {
						goto l21
					}
					v11 = v6 & i32(1)
					v10 = v6 & i32(0x3ffffffe)
					v3 = i32(0)
				l22:
					{
						v6 = v8 + v3
						t42 := int64(load64(m.memory[uint32(v6):]))
						t43 := v6
						v9 = t42
						store64(m.memory[uint32(t43):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
						v6 = v6 + i32(8)
						t44 := int64(load64(m.memory[uint32(v6):]))
						t45 := v6
						v9 = t44
						store64(m.memory[uint32(t45):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
						v3 = v3 + i32(16)
						v10 = v10 + i32(-2)
						if v10 != 0 {
							goto l22
						}
					}
					if v11 == 0 {
						goto l23
					}
				}
			l21:
				v3 = v8 + v3
				t46 := int64(load64(m.memory[uint32(v3):]))
				t47 := v3
				v9 = t46
				store64(m.memory[uint32(t47):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
			}
		l23:
			{
				if uint32(v5) < uint32(i32(8)) {
					goto l24
				}
				t48 := int64(load64(m.memory[uint32(v8):]))
				store64(m.memory[uint32(v8+v5):], uint64(t48))
				goto l25
			}
		l24:
			if v5 == 0 {
				goto l25
			}
			memory_copy(m.memory, uint32(v8+i32(8)), uint32(v8), uint32(v5))
		l25:
			t49 := int64(load64(m.memory[int64(uint32(v1))+8:]))
			v16 = t49
			t50 := int64(load64(m.memory[uint32(v1):]))
			v18 = t50
			v6 = i32(0)
		l33:
			{
				t51 := v8
				v3 = v6
				v10 = t51 + v3
				t52 := int32(m.memory[uint32(v10)])
				if t52 != i32(128) {
					goto l26
				}
				v15 = v8 + (v3^i32(-1))*i32(20)
				v6 = v8 + (i32(0)-v3)*i32(20)
				v11 = v6 + i32(-16)
				v14 = v6 + i32(-20)
			l32:
				{
					t53 := int32(load32(m.memory[uint32(v14):]))
					t54 := int32(load32(m.memory[uint32(v11):]))
					t55 := m.fn81(v18, v16, t53, t54)
					t56 := v4
					v1 = int32(t55)
					v6 = t56 & v1
					v5 = v6
					{
						t57 := int64(load64(m.memory[uint32(v8+v6):]))
						v9 = t57 & i64(-0x7f7f7f7f7f7f7f80)
						if v9 != i64(0) {
							goto l27
						}
						v17 = i32(8)
						v5 = v6
					l28:
						{
							v5 = v5 + v17
							v17 = v17 + i32(8)
							t58 := v8
							v5 = v5 & v4
							t59 := int64(load64(m.memory[uint32(t58+v5):]))
							v9 = t59 & i64(-0x7f7f7f7f7f7f7f80)
							if v9 == 0 {
								goto l28
							}
						}
					}
				l27:
					{
						t60 := v8
						v5 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v9))))>>3) + v5) & v4
						t61 := int32(int8(m.memory[uint32(t60+v5)]))
						if t61 < i32(0) {
							goto l29
						}
						t62 := int64(load64(m.memory[uint32(v8):]))
						v5 = int32(uint32(int64(bits.TrailingZeros64(uint64(t62&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
					}
				l29:
					{
						if uint32((v5-v6^(v3-v6))&v4) < uint32(i32(8)) {
							goto l30
						}
						v6 = v8 + v5
						t63 := int32(m.memory[uint32(v6)])
						v17 = t63
						t64 := v6
						v1 = int32(uint32(v1) >> 25)
						m.memory[uint32(t64)] = byte(v1)
						m.memory[uint32(v8+(v5+i32(-8))&v4+i32(8))] = byte(v1)
						v6 = v8 + (v5^i32(-1))*i32(20)
						{
							if v17 != i32(255) {
								t68 := int32(load32(m.memory[uint32(v15):]))
								v5 = t68
								t69 := int32(load32(m.memory[uint32(v6):]))
								store32(m.memory[uint32(v15):], uint32(t69))
								store32(m.memory[uint32(v6):], uint32(v5))
								t70 := int32(load32(m.memory[int64(uint32(v6))+4:]))
								v5 = t70
								t71 := int32(load32(m.memory[int64(uint32(v15))+4:]))
								store32(m.memory[int64(uint32(v6))+4:], uint32(t71))
								store32(m.memory[int64(uint32(v15))+4:], uint32(v5))
								t72 := int32(load32(m.memory[int64(uint32(v15))+8:]))
								v5 = t72
								t73 := int32(load32(m.memory[int64(uint32(v6))+8:]))
								store32(m.memory[int64(uint32(v15))+8:], uint32(t73))
								store32(m.memory[int64(uint32(v6))+8:], uint32(v5))
								t74 := int32(load32(m.memory[int64(uint32(v6))+12:]))
								v5 = t74
								t75 := int32(load32(m.memory[int64(uint32(v15))+12:]))
								store32(m.memory[int64(uint32(v6))+12:], uint32(t75))
								store32(m.memory[int64(uint32(v15))+12:], uint32(v5))
								t76 := int32(load32(m.memory[int64(uint32(v15))+16:]))
								v5 = t76
								t77 := int32(load32(m.memory[int64(uint32(v6))+16:]))
								store32(m.memory[int64(uint32(v15))+16:], uint32(t77))
								store32(m.memory[int64(uint32(v6))+16:], uint32(v5))
								goto l32
							}
							m.memory[uint32(v10)] = byte(i32(255))
							m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(i32(255))
							t65 := int32(load32(m.memory[int64(uint32(v15))+16:]))
							store32(m.memory[int64(uint32(v6))+16:], uint32(t65))
							t66 := int64(load64(m.memory[int64(uint32(v15))+8:]))
							store64(m.memory[int64(uint32(v6))+8:], uint64(t66))
							t67 := int64(load64(m.memory[uint32(v15):]))
							store64(m.memory[uint32(v6):], uint64(t67))
							goto l26
						}
					}
				l30:
				}
				t78 := v10
				v6 = int32(uint32(v1) >> 25)
				m.memory[uint32(t78)] = byte(v6)
				m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(v6)
			}
		l26:
			v6 = v3 + i32(1)
			if v3 != v4 {
				goto l33
			}
			p79 := v7
			if uint32(v4) < uint32(i32(8)) {
				p79 = v4
			}
			v3 = p79
		}
	l20:
		store32(m.memory[int64(uint32(v0))+8:], uint32(v3-v2))
	l15:
		return i32(-1)
	}
l5:
	m.fn27(i32(1271632), i32(57), i32(1271660))
	panic("unreachable")
}
func (m *Module) fn84(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7, v8, v9, v10 int32
	var v11, v12, v13 int64
	var v14, v15, v16 int32
	var v17, v18 int64
	{
		{
			{
				t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
				v2 = t0
				v3 = v2 + i32(1)
				if v3 == 0 {
					m.fn27(i32(1271632), i32(57), i32(1271660))
					panic("unreachable")
				}
				{
					t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
					t2 := v3
					v4 = t1
					t3 := v4
					v5 = v4 + i32(1)
					v6 = int32(uint32(v5) >> 3)
					v7 = v6 * i32(7)
					p4 := v7
					if uint32(v4) < uint32(i32(8)) {
						p4 = t3
					}
					v8 = p4
					if uint32(t2) <= uint32(int32(uint32(v8)>>1)) {
						{
							if v5 != 0 {
								goto l6
							}
							v3 = i32(0)
							goto l7
						l6:
							t7 := int32(load32(m.memory[uint32(v0):]))
							v8 = t7
							v3 = i32(0)
							{
								{
									t8 := v6
									var p9 int32
									if v5&i32(7) != i32(0) {
										p9 = 1
									}
									v6 = t8 + p9
									if v6 == i32(1) {
										goto l8
									}
									v9 = v6 & i32(1)
									v10 = v6 & i32(0x3ffffffe)
									v3 = i32(0)
								l9:
									{
										v6 = v8 + v3
										t10 := int64(load64(m.memory[uint32(v6):]))
										t11 := v6
										v11 = t10
										store64(m.memory[uint32(t11):], uint64(int64(uint64(v11^i64(-1))>>7)&i64(72340172838076673)+(v11|i64(0x7f7f7f7f7f7f7f7f))))
										v6 = v6 + i32(8)
										t12 := int64(load64(m.memory[uint32(v6):]))
										t13 := v6
										v11 = t12
										store64(m.memory[uint32(t13):], uint64(int64(uint64(v11^i64(-1))>>7)&i64(72340172838076673)+(v11|i64(0x7f7f7f7f7f7f7f7f))))
										v3 = v3 + i32(16)
										v10 = v10 + i32(-2)
										if v10 != 0 {
											goto l9
										}
									}
									if v9 == 0 {
										goto l10
									}
								}
							l8:
								v3 = v8 + v3
								t14 := int64(load64(m.memory[uint32(v3):]))
								t15 := v3
								v11 = t14
								store64(m.memory[uint32(t15):], uint64(int64(uint64(v11^i64(-1))>>7)&i64(72340172838076673)+(v11|i64(0x7f7f7f7f7f7f7f7f))))
							}
						l10:
							{
								if uint32(v5) < uint32(i32(8)) {
									goto l11
								}
								t16 := int64(load64(m.memory[uint32(v8):]))
								store64(m.memory[uint32(v8+v5):], uint64(t16))
								goto l12
							}
						l11:
							if v5 == 0 {
								goto l12
							}
							memory_copy(m.memory, uint32(v8+i32(8)), uint32(v8), uint32(v5))
						l12:
							t17 := int64(load64(m.memory[int64(uint32(v1))+8:]))
							v12 = t17
							t18 := int64(load64(m.memory[uint32(v1):]))
							v13 = t18
							v6 = i32(0)
						l20:
							{
								t19 := v8
								v3 = v6
								v10 = t19 + v3
								t20 := int32(m.memory[uint32(v10)])
								if t20 != i32(128) {
									goto l13
								}
								v6 = v8 - v3<<3
								v1 = v6 + i32(-4)
								v14 = v6 + i32(-8)
								v15 = v8 + (v3^i32(-1))<<3
							l19:
								{
									t21 := int32(load32(m.memory[uint32(v14):]))
									t22 := int32(load32(m.memory[uint32(v1):]))
									t23 := m.fn81(v13, v12, t21, t22)
									t24 := v4
									v9 = int32(t23)
									v6 = t24 & v9
									v5 = v6
									{
										t25 := int64(load64(m.memory[uint32(v8+v6):]))
										v11 = t25 & i64(-0x7f7f7f7f7f7f7f80)
										if v11 != i64(0) {
											goto l14
										}
										v16 = i32(8)
										v5 = v6
									l15:
										{
											v5 = v5 + v16
											v16 = v16 + i32(8)
											t26 := v8
											v5 = v5 & v4
											t27 := int64(load64(m.memory[uint32(t26+v5):]))
											v11 = t27 & i64(-0x7f7f7f7f7f7f7f80)
											if v11 == 0 {
												goto l15
											}
										}
									}
								l14:
									{
										t28 := v8
										v5 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v11))))>>3) + v5) & v4
										t29 := int32(int8(m.memory[uint32(t28+v5)]))
										if t29 < i32(0) {
											goto l16
										}
										t30 := int64(load64(m.memory[uint32(v8):]))
										v5 = int32(uint32(int64(bits.TrailingZeros64(uint64(t30&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
									}
								l16:
									{
										if uint32((v5-v6^(v3-v6))&v4) < uint32(i32(8)) {
											goto l17
										}
										v6 = v8 + v5
										t31 := int32(m.memory[uint32(v6)])
										v16 = t31
										t32 := v6
										v9 = int32(uint32(v9) >> 25)
										m.memory[uint32(t32)] = byte(v9)
										m.memory[uint32(v8+(v5+i32(-8))&v4+i32(8))] = byte(v9)
										v6 = v8 + (v5^i32(-1))<<3
										{
											if v16 != i32(255) {
												t34 := int64(load64(m.memory[uint32(v15):]))
												v11 = t34
												t35 := int64(load64(m.memory[uint32(v6):]))
												store64(m.memory[uint32(v15):], uint64(t35))
												store64(m.memory[uint32(v6):], uint64(v11))
												goto l19
											}
											m.memory[uint32(v10)] = byte(i32(255))
											m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(i32(255))
											t33 := int64(load64(m.memory[uint32(v15):]))
											store64(m.memory[uint32(v6):], uint64(t33))
											goto l13
										}
									}
								l17:
								}
								t36 := v10
								v6 = int32(uint32(v9) >> 25)
								m.memory[uint32(t36)] = byte(v6)
								m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(v6)
							}
						l13:
							v6 = v3 + i32(1)
							if v3 != v4 {
								goto l20
							}
							p37 := v7
							if uint32(v4) < uint32(i32(8)) {
								p37 = v4
							}
							v3 = p37
						}
					l7:
						store32(m.memory[int64(uint32(v0))+8:], uint32(v3-v2))
						goto l21
					}
					v8 = v8 + i32(1)
					p5 := v3
					if uint32(v8) > uint32(v3) {
						p5 = v8
					}
					v3 = p5
					if uint32(v3) < uint32(i32(15)) {
						goto l2
					}
					{
						if uint32(v3) > uint32(i32(0x1fffffff)) {
							m.fn27(i32(1271632), i32(57), i32(1271660))
							panic("unreachable")
						}
						t6 := int32(uint32(v3<<3) / uint32(i32(7)))
						v3 = i32_shr_u(i32(-1), int32(bits.LeadingZeros32(uint32(t6+i32(-1)))))
						if uint32(v3) > uint32(i32(0x1ffffffe)) {
							goto l4
						}
						v3 = v3 + i32(1)
						goto l5
					}
				}
			}
		l2:
			p38 := v3&i32(8) + i32(8)
			if uint32(v3) < uint32(i32(4)) {
				p38 = i32(4)
			}
			v3 = p38
		}
	l5:
		v8 = v3 + i32(8)
		t39 := v8
		v10 = v3 << 3
		v6 = t39 + v10
		if uint32(v6) < uint32(v8) {
			goto l4
		}
		if uint32(v6) > uint32(i32(0x7ffffff8)) {
			goto l4
		}
		{
			t40 := m.fn7(v6)
			v5 = t40
			if v5 != 0 {
				v6 = v5 + v10
				if v8 == 0 {
					goto l23
				}
				memory_fill(m.memory, uint32(v6), i32(255), uint32(v8))
			l23:
				v5 = v3 + i32(-1)
				p41 := int32(uint32(v3)>>3) * i32(7)
				if uint32(v3) < uint32(i32(9)) {
					p41 = v5
				}
				v7 = p41
				t42 := int32(load32(m.memory[uint32(v0):]))
				v9 = t42
				{
					if v2 == 0 {
						goto l24
					}
					t43 := int64(load64(m.memory[uint32(v9):]))
					v11 = (t43 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
					t44 := int64(load64(m.memory[int64(uint32(v1))+8:]))
					v17 = t44
					t45 := int64(load64(m.memory[uint32(v1):]))
					v18 = t45
					v8 = v9
					v1 = v2
					v3 = i32(0)
				l30:
					{
						if v11 != i64(0) {
							goto l25
						}
					l26:
						{
							v3 = v3 + i32(8)
							v8 = v8 + i32(8)
							t46 := int64(load64(m.memory[uint32(v8):]))
							v11 = t46 & i64(-0x7f7f7f7f7f7f7f80)
							if v11 == i64(-0x7f7f7f7f7f7f7f80) {
								goto l26
							}
						}
						v11 = v11 ^ i64(-0x7f7f7f7f7f7f7f80)
					l25:
						{
							t47 := v6
							t48 := v5
							t49 := v18
							t50 := v17
							t51 := v9
							v14 = int32(uint32(int64(bits.TrailingZeros64(uint64(v11))))>>3) + v3
							v10 = t51 - v14<<3
							t52 := int32(load32(m.memory[uint32(v10+i32(-8)):]))
							t53 := int32(load32(m.memory[uint32(v10+i32(-4)):]))
							t54 := m.fn81(t49, t50, t52, t53)
							v15 = int32(t54)
							v10 = t48 & v15
							t55 := int64(load64(m.memory[uint32(t47+v10):]))
							v12 = t55 & i64(-0x7f7f7f7f7f7f7f80)
							if v12 != i64(0) {
								goto l27
							}
							v16 = i32(8)
						l28:
							{
								v10 = v10 + v16
								v16 = v16 + i32(8)
								t56 := v6
								v10 = v10 & v5
								t57 := int64(load64(m.memory[uint32(t56+v10):]))
								v12 = t57 & i64(-0x7f7f7f7f7f7f7f80)
								if v12 == 0 {
									goto l28
								}
							}
						}
					l27:
						v13 = v11 + i64(-1)
						{
							t58 := v6
							v10 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v12))))>>3) + v10) & v5
							t59 := int32(int8(m.memory[uint32(t58+v10)]))
							if t59 < i32(0) {
								goto l29
							}
							t60 := int64(load64(m.memory[uint32(v6):]))
							v10 = int32(uint32(int64(bits.TrailingZeros64(uint64(t60&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
						}
					l29:
						v11 = v13 & v11
						t61 := v6 + v10
						v15 = int32(uint32(v15) >> 25)
						m.memory[uint32(t61)] = byte(v15)
						m.memory[uint32(v6+(v10+i32(-8))&v5+i32(8))] = byte(v15)
						t62 := int64(load64(m.memory[uint32(v9+(v14^i32(-1))<<3):]))
						store64(m.memory[uint32(v6+(v10^i32(-1))<<3):], uint64(t62))
						v1 = v1 + i32(-1)
						if v1 != 0 {
							goto l30
						}
					}
				}
			l24:
				store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
				store32(m.memory[uint32(v0):], uint32(v6))
				store32(m.memory[int64(uint32(v0))+8:], uint32(v7-v2))
				if v4 == 0 {
					goto l21
				}
				t63 := v4
				v8 = (v4<<3 + i32(15)) & i32(-8)
				v3 = t63 + v8 + i32(9)
				if v3 == 0 {
					goto l21
				}
				v4 = v9 - v8
				t64 := int32(load32(m.memory[uint32(v4+i32(-4)):]))
				v8 = t64
				v6 = v8 & i32(-8)
				t65 := v6
				v8 = v8 & i32(3)
				p66 := i32(8)
				if v8 != 0 {
					p66 = i32(4)
				}
				if uint32(t65) < uint32(p66+v3) {
					m.fn3(i32(1274224), i32(46), i32(1274272))
					panic("unreachable")
				}
				if v8 == 0 {
					goto l32
				}
				if uint32(v6) > uint32(v3+i32(39)) {
					m.fn3(i32(1274288), i32(46), i32(1274336))
					panic("unreachable")
				}
			l32:
				m.fn1(v4)
				return i32(-1)
			}
			m.fn23(i32(8), v6)
			panic("unreachable")
		}
	}
l21:
	return i32(-1)
l4:
	m.fn27(i32(1271632), i32(57), i32(1271660))
	panic("unreachable")
}
func (m *Module) fn85(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7, v8 int32
	var v9 int64
	var v10, v11, v12 int32
	var v13, v14 int64
	var v15 int32
	var v16 int64
	var v17 int32
	var v18 int64
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
		v2 = t0
		v3 = v2 + i32(1)
		if v3 == 0 {
			m.fn27(i32(1271632), i32(57), i32(1271660))
			panic("unreachable")
		}
		{
			t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t2 := v3
			v4 = t1
			t3 := v4
			v5 = v4 + i32(1)
			v6 = int32(uint32(v5) >> 3)
			v7 = v6 * i32(7)
			p4 := v7
			if uint32(v4) < uint32(i32(8)) {
				p4 = t3
			}
			v8 = p4
			if uint32(t2) <= uint32(int32(uint32(v8)>>1)) {
				goto l1
			}
			{
				{
					{
						v8 = v8 + i32(1)
						p5 := v3
						if uint32(v8) > uint32(v3) {
							p5 = v8
						}
						v3 = p5
						if uint32(v3) < uint32(i32(15)) {
							goto l2
						}
						if uint32(v3) > uint32(i32(0x1fffffff)) {
							m.fn27(i32(1271632), i32(57), i32(1271660))
							panic("unreachable")
						}
						t6 := int32(uint32(v3<<3) / uint32(i32(7)))
						v3 = i32_shr_u(i32(-1), int32(bits.LeadingZeros32(uint32(t6+i32(-1))))) + i32(1)
						goto l4
					}
				l2:
					p7 := v3&i32(8) + i32(8)
					if uint32(v3) < uint32(i32(4)) {
						p7 = i32(4)
					}
					v3 = p7
				}
			l4:
				v9 = int64(uint32(v3)) * i64(24)
				if int32(int64(uint64(v9)>>32)) != 0 {
					goto l5
				}
				v8 = v3 + i32(8)
				t8 := v8
				v10 = int32(v9)
				v6 = t8 + v10
				if uint32(v6) < uint32(v8) {
					goto l5
				}
				if uint32(v6) > uint32(i32(0x7ffffff8)) {
					goto l5
				}
				t9 := m.fn7(v6)
				v5 = t9
				if v5 != 0 {
					v6 = v5 + v10
					if v8 == 0 {
						goto l7
					}
					memory_fill(m.memory, uint32(v6), i32(255), uint32(v8))
				l7:
					v5 = v3 + i32(-1)
					p10 := int32(uint32(v3)>>3) * i32(7)
					if uint32(v3) < uint32(i32(9)) {
						p10 = v5
					}
					v11 = p10
					t11 := int32(load32(m.memory[uint32(v0):]))
					v12 = t11
					{
						if v2 == 0 {
							goto l8
						}
						t12 := int64(load64(m.memory[uint32(v12):]))
						v9 = (t12 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
						t13 := int64(load64(m.memory[int64(uint32(v1))+8:]))
						v13 = t13
						t14 := int64(load64(m.memory[uint32(v1):]))
						v14 = t14
						v8 = v12
						v15 = v2
						v3 = i32(0)
					l14:
						{
							if v9 != i64(0) {
								goto l9
							}
						l10:
							{
								v3 = v3 + i32(8)
								v8 = v8 + i32(8)
								t15 := int64(load64(m.memory[uint32(v8):]))
								v9 = t15 & i64(-0x7f7f7f7f7f7f7f80)
								if v9 == i64(-0x7f7f7f7f7f7f7f80) {
									goto l10
								}
							}
							v9 = v9 ^ i64(-0x7f7f7f7f7f7f7f80)
						l9:
							{
								t16 := v6
								t17 := v5
								t18 := v14
								t19 := v13
								t20 := v12
								v1 = int32(uint32(int64(bits.TrailingZeros64(uint64(v9))))>>3) + v3
								t21 := m.fn86(t18, t19, t20+(i32(0)-v1)*i32(24)+i32(-24))
								v7 = int32(t21)
								v10 = t17 & v7
								t22 := int64(load64(m.memory[uint32(t16+v10):]))
								v16 = t22 & i64(-0x7f7f7f7f7f7f7f80)
								if v16 != i64(0) {
									goto l11
								}
								v17 = i32(8)
							l12:
								{
									v10 = v10 + v17
									v17 = v17 + i32(8)
									t23 := v6
									v10 = v10 & v5
									t24 := int64(load64(m.memory[uint32(t23+v10):]))
									v16 = t24 & i64(-0x7f7f7f7f7f7f7f80)
									if v16 == 0 {
										goto l12
									}
								}
							}
						l11:
							v18 = v9 + i64(-1)
							{
								t25 := v6
								v10 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v16))))>>3) + v10) & v5
								t26 := int32(int8(m.memory[uint32(t25+v10)]))
								if t26 < i32(0) {
									goto l13
								}
								t27 := int64(load64(m.memory[uint32(v6):]))
								v10 = int32(uint32(int64(bits.TrailingZeros64(uint64(t27&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
							}
						l13:
							v9 = v18 & v9
							t28 := v6 + v10
							v7 = int32(uint32(v7) >> 25)
							m.memory[uint32(t28)] = byte(v7)
							m.memory[uint32(v6+(v10+i32(-8))&v5+i32(8))] = byte(v7)
							v10 = v6 + (v10^i32(-1))*i32(24)
							t29 := v10
							v1 = v12 + (v1^i32(-1))*i32(24)
							t30 := int64(load64(m.memory[int64(uint32(v1))+16:]))
							store64(m.memory[int64(uint32(t29))+16:], uint64(t30))
							t31 := int64(load64(m.memory[int64(uint32(v1))+8:]))
							store64(m.memory[int64(uint32(v10))+8:], uint64(t31))
							t32 := int64(load64(m.memory[uint32(v1):]))
							store64(m.memory[uint32(v10):], uint64(t32))
							v15 = v15 + i32(-1)
							if v15 != 0 {
								goto l14
							}
						}
					}
				l8:
					store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
					store32(m.memory[uint32(v0):], uint32(v6))
					store32(m.memory[int64(uint32(v0))+8:], uint32(v11-v2))
					if v4 == 0 {
						goto l15
					}
					t33 := v4
					v8 = (v4*i32(24) + i32(31)) & i32(-8)
					v3 = t33 + v8 + i32(9)
					if v3 == 0 {
						goto l15
					}
					v4 = v12 - v8
					t34 := int32(load32(m.memory[uint32(v4+i32(-4)):]))
					v8 = t34
					v6 = v8 & i32(-8)
					t35 := v6
					v8 = v8 & i32(3)
					p36 := i32(8)
					if v8 != 0 {
						p36 = i32(4)
					}
					if uint32(t35) < uint32(p36+v3) {
						m.fn3(i32(1274224), i32(46), i32(1274272))
						panic("unreachable")
					}
					if v8 == 0 {
						goto l17
					}
					if uint32(v6) > uint32(v3+i32(39)) {
						m.fn3(i32(1274288), i32(46), i32(1274336))
						panic("unreachable")
					}
				l17:
					m.fn1(v4)
					return i32(-1)
				}
				m.fn23(i32(8), v6)
				panic("unreachable")
			}
		}
	l1:
		{
			if v5 != 0 {
				goto l19
			}
			v3 = i32(0)
			goto l20
		l19:
			t37 := int32(load32(m.memory[uint32(v0):]))
			v8 = t37
			v3 = i32(0)
			{
				{
					t38 := v6
					var p39 int32
					if v5&i32(7) != i32(0) {
						p39 = 1
					}
					v6 = t38 + p39
					if v6 == i32(1) {
						goto l21
					}
					v12 = v6 & i32(1)
					v10 = v6 & i32(0x3ffffffe)
					v3 = i32(0)
				l22:
					{
						v6 = v8 + v3
						t40 := int64(load64(m.memory[uint32(v6):]))
						t41 := v6
						v9 = t40
						store64(m.memory[uint32(t41):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
						v6 = v6 + i32(8)
						t42 := int64(load64(m.memory[uint32(v6):]))
						t43 := v6
						v9 = t42
						store64(m.memory[uint32(t43):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
						v3 = v3 + i32(16)
						v10 = v10 + i32(-2)
						if v10 != 0 {
							goto l22
						}
					}
					if v12 == 0 {
						goto l23
					}
				}
			l21:
				v3 = v8 + v3
				t44 := int64(load64(m.memory[uint32(v3):]))
				t45 := v3
				v9 = t44
				store64(m.memory[uint32(t45):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
			}
		l23:
			{
				if uint32(v5) < uint32(i32(8)) {
					goto l24
				}
				t46 := int64(load64(m.memory[uint32(v8):]))
				store64(m.memory[uint32(v8+v5):], uint64(t46))
				goto l25
			}
		l24:
			if v5 == 0 {
				goto l25
			}
			memory_copy(m.memory, uint32(v8+i32(8)), uint32(v8), uint32(v5))
		l25:
			t47 := int64(load64(m.memory[int64(uint32(v1))+8:]))
			v16 = t47
			t48 := int64(load64(m.memory[uint32(v1):]))
			v18 = t48
			v6 = i32(0)
		l33:
			{
				t49 := v8
				v3 = v6
				v10 = t49 + v3
				t50 := int32(m.memory[uint32(v10)])
				if t50 != i32(128) {
					goto l26
				}
				v15 = v8 + (v3^i32(-1))*i32(24)
				v12 = v8 + (i32(0)-v3)*i32(24) + i32(-24)
			l32:
				{
					t51 := m.fn86(v18, v16, v12)
					t52 := v4
					v1 = int32(t51)
					v6 = t52 & v1
					v5 = v6
					{
						t53 := int64(load64(m.memory[uint32(v8+v6):]))
						v9 = t53 & i64(-0x7f7f7f7f7f7f7f80)
						if v9 != i64(0) {
							goto l27
						}
						v11 = i32(8)
						v5 = v6
					l28:
						{
							v5 = v5 + v11
							v11 = v11 + i32(8)
							t54 := v8
							v5 = v5 & v4
							t55 := int64(load64(m.memory[uint32(t54+v5):]))
							v9 = t55 & i64(-0x7f7f7f7f7f7f7f80)
							if v9 == 0 {
								goto l28
							}
						}
					}
				l27:
					{
						t56 := v8
						v5 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v9))))>>3) + v5) & v4
						t57 := int32(int8(m.memory[uint32(t56+v5)]))
						if t57 < i32(0) {
							goto l29
						}
						t58 := int64(load64(m.memory[uint32(v8):]))
						v5 = int32(uint32(int64(bits.TrailingZeros64(uint64(t58&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
					}
				l29:
					{
						if uint32((v5-v6^(v3-v6))&v4) < uint32(i32(8)) {
							goto l30
						}
						v6 = v8 + v5
						t59 := int32(m.memory[uint32(v6)])
						v11 = t59
						t60 := v6
						v1 = int32(uint32(v1) >> 25)
						m.memory[uint32(t60)] = byte(v1)
						m.memory[uint32(v8+(v5+i32(-8))&v4+i32(8))] = byte(v1)
						v6 = v8 + (v5^i32(-1))*i32(24)
						{
							if v11 != i32(255) {
								t64 := int64(load64(m.memory[uint32(v15):]))
								v9 = t64
								t65 := int64(load64(m.memory[uint32(v6):]))
								store64(m.memory[uint32(v15):], uint64(t65))
								store64(m.memory[uint32(v6):], uint64(v9))
								t66 := int64(load64(m.memory[int64(uint32(v6))+8:]))
								v9 = t66
								t67 := int64(load64(m.memory[int64(uint32(v15))+8:]))
								store64(m.memory[int64(uint32(v6))+8:], uint64(t67))
								store64(m.memory[int64(uint32(v15))+8:], uint64(v9))
								t68 := int32(load32(m.memory[int64(uint32(v15))+16:]))
								v5 = t68
								t69 := int32(load32(m.memory[int64(uint32(v6))+16:]))
								store32(m.memory[int64(uint32(v15))+16:], uint32(t69))
								t70 := int32(load32(m.memory[int64(uint32(v6))+20:]))
								v1 = t70
								t71 := int32(load32(m.memory[int64(uint32(v15))+20:]))
								store32(m.memory[int64(uint32(v6))+20:], uint32(t71))
								store32(m.memory[int64(uint32(v15))+20:], uint32(v1))
								store32(m.memory[int64(uint32(v6))+16:], uint32(v5))
								goto l32
							}
							m.memory[uint32(v10)] = byte(i32(255))
							m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(i32(255))
							t61 := int64(load64(m.memory[int64(uint32(v15))+16:]))
							store64(m.memory[int64(uint32(v6))+16:], uint64(t61))
							t62 := int64(load64(m.memory[int64(uint32(v15))+8:]))
							store64(m.memory[int64(uint32(v6))+8:], uint64(t62))
							t63 := int64(load64(m.memory[uint32(v15):]))
							store64(m.memory[uint32(v6):], uint64(t63))
							goto l26
						}
					}
				l30:
				}
				t72 := v10
				v6 = int32(uint32(v1) >> 25)
				m.memory[uint32(t72)] = byte(v6)
				m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(v6)
			}
		l26:
			v6 = v3 + i32(1)
			if v3 != v4 {
				goto l33
			}
			p73 := v7
			if uint32(v4) < uint32(i32(8)) {
				p73 = v4
			}
			v3 = p73
		}
	l20:
		store32(m.memory[int64(uint32(v0))+8:], uint32(v3-v2))
	l15:
		return i32(-1)
	}
l5:
	m.fn27(i32(1271632), i32(57), i32(1271660))
	panic("unreachable")
}
func (m *Module) fn86(v0, v1 int64, v2 int32) int64 {
	var v3 int32
	var v4, v5, v6, v7 int64
	t0 := m.g0
	v3 = t0 - i32(80)
	m.g0 = v3
	store64(m.memory[int64(uint32(v3))+56:], uint64(i64(0)))
	store64(m.memory[int64(uint32(v3))+64:], uint64(i64(0)))
	store64(m.memory[int64(uint32(v3))+48:], uint64(v1))
	store64(m.memory[int64(uint32(v3))+32:], uint64(v1^i64(8387220255154660723)))
	store64(m.memory[int64(uint32(v3))+24:], uint64(v1^i64(7237128888997146477)))
	store64(m.memory[int64(uint32(v3))+40:], uint64(v0))
	store64(m.memory[int64(uint32(v3))+16:], uint64(v0^i64(0x6c7967656e657261)))
	store64(m.memory[int64(uint32(v3))+8:], uint64(v0^i64(8317987319222330741)))
	t1 := int32(load32(m.memory[int64(uint32(v2))+4:]))
	t2 := int32(load32(m.memory[int64(uint32(v2))+8:]))
	m.fn58(v3+i32(8), t1, t2)
	m.memory[int64(uint32(v3))+76] = byte(i32(255))
	m.fn58(v3+i32(8), v3+i32(76), i32(1))
	t3 := int32(load32(m.memory[int64(uint32(v2))+12:]))
	store32(m.memory[int64(uint32(v3))+76:], uint32(t3))
	m.fn58(v3+i32(8), v3+i32(76), i32(4))
	t4 := int64(load64(m.memory[int64(uint32(v3))+8:]))
	v0 = t4
	t5 := int64(load64(m.memory[int64(uint32(v3))+24:]))
	v1 = t5
	t6 := int64(load32(m.memory[int64(uint32(v3))+64:]))
	v4 = t6
	t7 := int64(load64(m.memory[int64(uint32(v3))+56:]))
	v5 = t7
	t8 := int64(load64(m.memory[int64(uint32(v3))+32:]))
	v6 = t8
	t9 := int64(load64(m.memory[int64(uint32(v3))+16:]))
	v7 = t9
	m.g0 = v3 + i32(80)
	t10 := v6
	v4 = v5 | v4<<56
	v5 = t10 ^ v4
	t11 := i64_rotl(v5, i64(16))
	v5 = v5 + v7
	v6 = t11 ^ v5
	t12 := i64_rotl(v6, i64(21))
	t13 := v6
	v0 = v1 + v0
	v6 = t13 + i64_rotl(v0, i64(32))
	v7 = t12 ^ v6
	t14 := i64_rotl(v7, i64(16))
	t15 := v7
	t16 := v5
	v1 = i64_rotl(v1, i64(13)) ^ v0
	v0 = t16 + v1
	v5 = t15 + (i64_rotl(v0, i64(32)) ^ i64(255))
	v7 = t14 ^ v5
	t17 := i64_rotl(v7, i64(21))
	t18 := v7
	t19 := v6 ^ v4
	v1 = v0 ^ i64_rotl(v1, i64(17))
	v0 = t19 + v1
	v4 = t18 + i64_rotl(v0, i64(32))
	v6 = t17 ^ v4
	t20 := i64_rotl(v6, i64(16))
	t21 := v6
	v1 = v0 ^ i64_rotl(v1, i64(13))
	v0 = v1 + v5
	v5 = t21 + i64_rotl(v0, i64(32))
	v6 = t20 ^ v5
	t22 := i64_rotl(v6, i64(21))
	t23 := v6
	v1 = v0 ^ i64_rotl(v1, i64(17))
	v0 = v1 + v4
	v4 = t23 + i64_rotl(v0, i64(32))
	v6 = t22 ^ v4
	t24 := i64_rotl(v6, i64(16))
	t25 := v6
	v1 = i64_rotl(v1, i64(13)) ^ v0
	v0 = v1 + v5
	v5 = t25 + i64_rotl(v0, i64(32))
	t26 := i64_rotl(t24^v5, i64(21))
	v1 = i64_rotl(v1, i64(17)) ^ v0
	v1 = i64_rotl(v1, i64(13)) ^ (v1 + v4)
	t27 := t26 ^ i64_rotl(v1, i64(17))
	v1 = v1 + v5
	return t27 ^ i64_rotl(v1, i64(32)) ^ v1
}
