package core

import (
	"math/bits"
)

func (m *Module) fn1842(v0, v1 int32) {
	store32(m.memory[uint32(v0):], uint32(i32(0)))
}
func (m *Module) fn1843(v0, v1 int32) {
	t0 := int64(load64(m.memory[int64(uint32(i32(0)))+1274380:]))
	store64(m.memory[int64(uint32(v0))+8:], uint64(t0))
	t1 := int64(load64(m.memory[int64(uint32(i32(0)))+1274372:]))
	store64(m.memory[uint32(v0):], uint64(t1))
}
func (m *Module) fn1844(v0, v1 int32) int32 {
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
							t12 := int32(m.memory[int64(uint32(v3&i32(15)))+1099352])
							m.memory[uint32(v2+i32(6)+v0+i32(-2))] = byte(t12)
							v0 = v0 + i32(-1)
							v3 = int32(uint32(v3) >> 4)
							if v3 != 0 {
								goto l6
							}
						}
						t13 := m.fn681(v1, i32(1), i32(1123086), i32(2), v2+i32(6)+v0+i32(-1), i32(9)-v0)
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
							t9 := int32(load16(m.memory[int64(uint32(v8<<1))+1100735:]))
							store16(m.memory[uint32(t5):], uint16(t9))
							t10 := int32(load16(m.memory[int64(uint32((v7-v8*i32(100))&i32(0xffff)<<1))+1100735:]))
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
				t18 := int32(load16(m.memory[int64(uint32((t17-v4*i32(100))&i32(0xffff)<<1))+1100735:]))
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
				t20 := int32(m.memory[int64(uint32(v4<<1))+1100736])
				m.memory[uint32(t19+v3)] = byte(t20)
			}
		l9:
			t21 := m.fn681(v1, int32(uint32(v0^i32(-1))>>31), i32(1), i32(0), v2+i32(6)+v3, i32(10)-v3)
			v0 = t21
			goto l7
		}
	l1:
		v3 = i32(9)
	l10:
		{
			t22 := int32(m.memory[int64(uint32(v0&i32(15)))+1123088])
			m.memory[uint32(v2+i32(6)+v3+i32(-2))] = byte(t22)
			v3 = v3 + i32(-1)
			v0 = int32(uint32(v0) >> 4)
			if v0 != 0 {
				goto l10
			}
		}
		t23 := m.fn681(v1, i32(1), i32(1123086), i32(2), v2+i32(6)+v3+i32(-1), i32(9)-v3)
		v0 = t23
	}
l7:
	m.g0 = v2 + i32(16)
	return v0
}
func (m *Module) fn1845(v0, v1 int32) {
	m.memory[int64(uint32(i32(0)))+1294464] = byte(i32(1))
	panic("unreachable")
}
func (m *Module) fn1846(v0 int32) {
	m.fn1847(v0)
	panic("unreachable")
}
func (m *Module) fn1847(v0 int32) {
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
			m.fn1840(t9, i32(1274344), t10, t11)
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
		m.fn1840(t5, i32(1274316), t6, t7)
		panic("unreachable")
	}
}
func (m *Module) fn1848(v0 int32) {
	var v1, v2, v3 int32
	{
		t0 := int32(load32(m.memory[uint32(v0):]))
		v1 = t0
		if v1 < i32(1) {
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
			m.fn7(i32(1274404), i32(46), i32(1274452))
			panic("unreachable")
		}
		if v0 == 0 {
			goto l2
		}
		if uint32(v3) > uint32(v1+i32(39)) {
			m.fn7(i32(1274468), i32(46), i32(1274516))
			panic("unreachable")
		}
	l2:
		m.fn5(v2)
	}
}
func (m *Module) fn1849(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	if t0 == i32(-1) {
		t7 := int32(load32(m.memory[uint32(v1):]))
		t8 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t9 := int32(load32(m.memory[int64(uint32(v0))+12:]))
		t10 := int32(load32(m.memory[uint32(t9):]))
		v0 = t10
		t11 := int32(load32(m.memory[uint32(v0):]))
		t12 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		t13 := m.fn46(t7, t8, t11, t12)
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
func (m *Module) fn1850(v0, v1 int32) {
	var v2, v3 int32
	var v4 int64
	t0 := m.g0
	v2 = t0 - i32(32)
	m.g0 = v2
	{
		t1 := int32(load32(m.memory[uint32(v1):]))
		if t1 != i32(-1) {
			goto l0
		}
		t2 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		v3 = t2
		store32(m.memory[int64(uint32(v2))+24:], uint32(i32(0)))
		store64(m.memory[int64(uint32(v2))+16:], uint64(i64(0x100000000)))
		t3 := int32(load32(m.memory[uint32(v3):]))
		t4 := v2 + i32(16)
		v3 = t3
		t5 := int32(load32(m.memory[uint32(v3):]))
		t6 := int32(load32(m.memory[int64(uint32(v3))+4:]))
		_ = m.fn46(t4, i32(1274292), t5, t6)
		t8 := int32(load32(m.memory[int64(uint32(v2))+24:]))
		t9 := v2
		v3 = t8
		store32(m.memory[int64(uint32(t9))+8:], uint32(v3))
		t10 := int64(load64(m.memory[int64(uint32(v2))+16:]))
		t11 := v2
		v4 = t10
		store64(m.memory[uint32(t11):], uint64(v4))
		store32(m.memory[int64(uint32(v1))+8:], uint32(v3))
		store64(m.memory[uint32(v1):], uint64(v4))
	}
l0:
	t12 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	v3 = t12
	store32(m.memory[int64(uint32(v1))+8:], uint32(i32(0)))
	t13 := int64(load64(m.memory[uint32(v1):]))
	v4 = t13
	store64(m.memory[uint32(v1):], uint64(i64(0x100000000)))
	store32(m.memory[int64(uint32(v2))+24:], uint32(v3))
	store64(m.memory[int64(uint32(v2))+16:], uint64(v4))
	{
		t14 := m.fn11(i32(12))
		v1 = t14
		if v1 != 0 {
			goto l1
		}
		m.fn23(i32(4), i32(12))
		panic("unreachable")
	}
l1:
	t15 := int32(load32(m.memory[int64(uint32(v2))+24:]))
	store32(m.memory[int64(uint32(v1))+8:], uint32(t15))
	t16 := int64(load64(m.memory[int64(uint32(v2))+16:]))
	store64(m.memory[uint32(v1):], uint64(t16))
	store32(m.memory[int64(uint32(v0))+4:], uint32(i32(1275236)))
	store32(m.memory[uint32(v0):], uint32(v1))
	m.g0 = v2 + i32(32)
}
func (m *Module) fn1851(v0, v1 int32) {
	var v2, v3 int32
	var v4 int64
	t0 := m.g0
	v2 = t0 - i32(32)
	m.g0 = v2
	{
		t1 := int32(load32(m.memory[uint32(v1):]))
		if t1 != i32(-1) {
			goto l0
		}
		t2 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		v3 = t2
		store32(m.memory[int64(uint32(v2))+28:], uint32(i32(0)))
		store64(m.memory[int64(uint32(v2))+20:], uint64(i64(0x100000000)))
		t3 := int32(load32(m.memory[uint32(v3):]))
		t4 := v2 + i32(20)
		v3 = t3
		t5 := int32(load32(m.memory[uint32(v3):]))
		t6 := int32(load32(m.memory[int64(uint32(v3))+4:]))
		_ = m.fn46(t4, i32(1274292), t5, t6)
		t8 := int32(load32(m.memory[int64(uint32(v2))+28:]))
		t9 := v2
		v3 = t8
		store32(m.memory[int64(uint32(t9))+16:], uint32(v3))
		t10 := int64(load64(m.memory[int64(uint32(v2))+20:]))
		t11 := v2
		v4 = t10
		store64(m.memory[int64(uint32(t11))+8:], uint64(v4))
		store32(m.memory[int64(uint32(v1))+8:], uint32(v3))
		store64(m.memory[uint32(v1):], uint64(v4))
	}
l0:
	store32(m.memory[int64(uint32(v0))+4:], uint32(i32(1275236)))
	store32(m.memory[uint32(v0):], uint32(v1))
	m.g0 = v2 + i32(32)
}
func (m *Module) fn1852(v0, v1, v2 int32) {
	var v3 int32
	t0 := m.g0
	v3 = t0 - i32(16)
	m.g0 = v3
	v1 = v2 + v1
	if uint32(v1) >= uint32(v2) {
		goto l0
	}
	m.fn16(i32(0), i32(0))
	panic("unreachable")
l0:
	t1 := int32(load32(m.memory[uint32(v0):]))
	t2 := v3 + i32(4)
	v2 = t1
	t3 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t4 := v2
	t5 := v1
	v2 = v2 << 1
	p6 := v2
	if uint32(v1) > uint32(v2) {
		p6 = t5
	}
	v2 = p6
	p7 := i32(8)
	if uint32(v2) > uint32(i32(8)) {
		p7 = v2
	}
	v2 = p7
	m.fn1853(t2, t4, t3, v2)
	{
		t8 := int32(load32(m.memory[int64(uint32(v3))+4:]))
		if t8 != i32(1) {
			goto l1
		}
		t9 := int32(load32(m.memory[int64(uint32(v3))+8:]))
		t10 := int32(load32(m.memory[int64(uint32(v3))+12:]))
		m.fn16(t9, t10)
		panic("unreachable")
	}
l1:
	t11 := int32(load32(m.memory[int64(uint32(v3))+8:]))
	v1 = t11
	store32(m.memory[uint32(v0):], uint32(v2))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
	m.g0 = v3 + i32(16)
}
func (m *Module) fn1853(v0, v1, v2, v3 int32) {
	var v4 int32
	v4 = i32(0)
	if v3 >= i32(0) {
		goto l0
	}
	v1 = i32(1)
	v2 = i32(4)
	goto l1
l0:
	{
		{
			if v1 == 0 {
				goto l2
			}
			t0 := m.fn26(v2, v1, i32(1), v3)
			v4 = t0
			goto l3
		}
	l2:
		t1 := m.fn11(v3)
		v4 = t1
	}
l3:
	if v4 != 0 {
		goto l4
	}
	v1 = i32(1)
	store32(m.memory[int64(uint32(v0))+4:], uint32(i32(1)))
	goto l5
l4:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v4))
	v1 = i32(0)
l5:
	v2 = i32(8)
	v4 = v3
l1:
	store32(m.memory[uint32(v0+v2):], uint32(v4))
	store32(m.memory[uint32(v0):], uint32(v1))
}
func (m *Module) fn1854(v0, v1 int32) {
	t0 := int64(load64(m.memory[int64(uint32(i32(0)))+1274396:]))
	store64(m.memory[int64(uint32(v0))+8:], uint64(t0))
	t1 := int64(load64(m.memory[int64(uint32(i32(0)))+1274388:]))
	store64(m.memory[uint32(v0):], uint64(t1))
}
func (m *Module) fn1855(v0, v1, v2 int32) int32 {
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
			m.fn1852(v0, v3, v2)
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
func (m *Module) fn1856(v0, v1 int32) int32 {
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
		m.fn1852(v0, v2, v3)
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
func (m *Module) fn1857(v0, v1, v2 int32) int32 {
	t0 := m.fn46(v0, i32(1274292), v1, v2)
	return t0
}
func (m *Module) fn1858(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v1):]))
	t1 := int32(load32(m.memory[uint32(v0):]))
	t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t3 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	t4 := int32(load32(m.memory[int64(uint32(t3))+12:]))
	t5 := m.t0[uint(t4)].(func(int32, int32, int32) int32)(t0, t1, t2)
	return t5
}
func (m *Module) fn1859(v0, v1 int32) {
	var v2, v3 int32
	t0 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	v2 = t0
	t1 := int32(load32(m.memory[uint32(v1):]))
	v3 = t1
	{
		t2 := m.fn11(i32(8))
		v1 = t2
		if v1 != 0 {
			goto l0
		}
		m.fn23(i32(4), i32(8))
		panic("unreachable")
	}
l0:
	store32(m.memory[int64(uint32(v1))+4:], uint32(v2))
	store32(m.memory[uint32(v1):], uint32(v3))
	store32(m.memory[int64(uint32(v0))+4:], uint32(i32(1274648)))
	store32(m.memory[uint32(v0):], uint32(v1))
}
func (m *Module) fn1860(v0, v1 int32) {
	store32(m.memory[int64(uint32(v0))+4:], uint32(i32(1274648)))
	store32(m.memory[uint32(v0):], uint32(v1))
}
func (m *Module) fn1861(v0, v1 int32) {
	t0 := int64(load64(m.memory[uint32(v1):]))
	store64(m.memory[uint32(v0):], uint64(t0))
}
func (m *Module) fn1862(v0, v1 int32) {
	var v2, v3, v4, v5 int32
	v2 = i32(31)
	{
		if uint32(v1) >= uint32(i32(0x1000000)) {
			goto l0
		}
		t0 := v1
		v2 = int32(bits.LeadingZeros32(uint32(int32(uint32(v1) >> 8))))
		v2 = i32_shr_u(t0, i32(38)-v2)&i32(1) | v2<<1 ^ i32(62)
	}
l0:
	store64(m.memory[int64(uint32(v0))+16:], uint64(i64(0)))
	store32(m.memory[int64(uint32(v0))+28:], uint32(v2))
	v3 = v2<<2 + i32(1294012)
	{
		t1 := int32(load32(m.memory[int64(uint32(i32(0)))+1294424:]))
		v4 = i32_shl(i32(1), v2)
		if t1&v4 != 0 {
			goto l1
		}
		store32(m.memory[uint32(v3):], uint32(v0))
		store32(m.memory[int64(uint32(v0))+24:], uint32(v3))
		store32(m.memory[int64(uint32(v0))+12:], uint32(v0))
		store32(m.memory[int64(uint32(v0))+8:], uint32(v0))
		t2 := int32(load32(m.memory[int64(uint32(i32(0)))+1294424:]))
		store32(m.memory[int64(uint32(i32(0)))+1294424:], uint32(t2|v4))
		return
	}
l1:
	{
		{
			{
				t3 := int32(load32(m.memory[uint32(v3):]))
				v4 = t3
				t4 := int32(load32(m.memory[int64(uint32(v4))+4:]))
				if t4&i32(-8) != v1 {
					goto l2
				}
				v2 = v4
				goto l3
			}
		l2:
			t6 := v1
			p5 := i32(25) - int32(uint32(v2)>>1)
			if v2 == i32(31) {
				p5 = i32(0)
			}
			v3 = i32_shl(t6, p5)
		l5:
			{
				v5 = v4 + int32(uint32(v3)>>29)&i32(4)
				t7 := int32(load32(m.memory[int64(uint32(v5))+16:]))
				v2 = t7
				if v2 == 0 {
					goto l4
				}
				v3 = v3 << 1
				v4 = v2
				t8 := int32(load32(m.memory[int64(uint32(v2))+4:]))
				if t8&i32(-8) != v1 {
					goto l5
				}
			}
		}
	l3:
		t9 := int32(load32(m.memory[int64(uint32(v2))+8:]))
		v3 = t9
		store32(m.memory[int64(uint32(v3))+12:], uint32(v0))
		store32(m.memory[int64(uint32(v2))+8:], uint32(v0))
		store32(m.memory[int64(uint32(v0))+24:], uint32(i32(0)))
		store32(m.memory[int64(uint32(v0))+12:], uint32(v2))
		store32(m.memory[int64(uint32(v0))+8:], uint32(v3))
		return
	}
l4:
	store32(m.memory[uint32(v5+i32(16)):], uint32(v0))
	store32(m.memory[int64(uint32(v0))+24:], uint32(v4))
	store32(m.memory[int64(uint32(v0))+12:], uint32(v0))
	store32(m.memory[int64(uint32(v0))+8:], uint32(v0))
}
func (m *Module) fn1863(v0, v1 int32) {
	m.fn1845(v0, v1)
	panic("unreachable")
}
func (m *Module) fn1864(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t1 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	t2 := m.fn9(v1, t0, t1)
	return t2
}
func (m *Module) fn1865(v0, v1 int32) int32 {
	var v2, v3 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	{
		{
			t1 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			v3 = t1
			if v3&i32(0x2000000) != 0 {
				t3 := int32(load32(m.memory[uint32(v0):]))
				v3 = t3
				v0 = i32(9)
			l3:
				{
					t4 := int32(m.memory[int64(uint32(v3&i32(15)))+1099352])
					m.memory[uint32(v2+i32(8)+v0+i32(-2))] = byte(t4)
					v0 = v0 + i32(-1)
					v3 = int32(uint32(v3) >> 4)
					if v3 != 0 {
						goto l3
					}
				}
				t5 := m.fn681(v1, i32(1), i32(1123086), i32(2), v2+i32(8)+v0+i32(-1), i32(9)-v0)
				v0 = t5
				goto l2
			}
			if v3&i32(0x4000000) != 0 {
				goto l1
			}
			t2 := m.fn508(v0, v1)
			v0 = t2
			goto l2
		}
	l1:
		t6 := int32(load32(m.memory[uint32(v0):]))
		v3 = t6
		v0 = i32(9)
	l4:
		{
			t7 := int32(m.memory[int64(uint32(v3&i32(15)))+1123088])
			m.memory[uint32(v2+i32(8)+v0+i32(-2))] = byte(t7)
			v0 = v0 + i32(-1)
			v3 = int32(uint32(v3) >> 4)
			if v3 != 0 {
				goto l4
			}
		}
		t8 := m.fn681(v1, i32(1), i32(1123086), i32(2), v2+i32(8)+v0+i32(-1), i32(9)-v0)
		v0 = t8
	}
l2:
	m.g0 = v2 + i32(16)
	return v0
}
func (m *Module) fn1866(v0, v1, v2 int32) int32 {
	t0 := v1
	v0 = v0 & i32(255) << 2
	t1 := int32(load32(m.memory[int64(uint32(v0))+1292568:]))
	t2 := int32(load32(m.memory[int64(uint32(v0))+1292400:]))
	t3 := int32(load32(m.memory[int64(uint32(v2))+12:]))
	t4 := m.t0[uint(t3)].(func(int32, int32, int32) int32)(t0, t1, t2)
	return t4
}
func (m *Module) fn1867() {
	m.fn1(i32(1271804), i32(50))
}
func (m *Module) fn1868(v0, v1, v2, v3 int32) {
	var v4 int32
	v4 = i32(1)
	if uint32(v3) <= uint32(i32(0x1fffffff)) {
		goto l0
	}
	v3 = i32(0)
	v1 = i32(4)
	goto l1
l0:
	v3 = v3 << 2
	{
		{
			if v1 == 0 {
				goto l2
			}
			t0 := m.fn26(v2, v1<<2, i32(4), v3)
			v1 = t0
			goto l3
		}
	l2:
		t1 := m.fn11(v3)
		v1 = t1
	}
l3:
	if v1 != 0 {
		goto l4
	}
	store32(m.memory[int64(uint32(v0))+4:], uint32(i32(4)))
	goto l5
l4:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
	v4 = i32(0)
l5:
	v1 = i32(8)
l1:
	store32(m.memory[uint32(v0+v1):], uint32(v3))
	store32(m.memory[uint32(v0):], uint32(v4))
}
func (m *Module) X__externref_drop_slice(v0, v1 int32) {
	var v2, v3 int32
	if v1 == 0 {
		return
	}
	v2 = v1 << 2
l4:
	{
		t0 := int32(load32(m.memory[uint32(v0):]))
		v1 = t0
		if uint32(v1) < uint32(i32(1028)) {
			goto l1
		}
		m.fn2(v1)
		t1 := int32(load32(m.memory[int64(uint32(i32(0)))+1294468:]))
		if t1 != 0 {
			m.fn350(i32(1275372))
			panic("unreachable")
		}
		t2 := int32(load32(m.memory[int64(uint32(i32(0)))+1294484:]))
		t3 := v1
		v3 = t2
		if uint32(t3) < uint32(v3) {
			goto l3
		}
		v1 = v1 - v3
		t4 := int32(load32(m.memory[int64(uint32(i32(0)))+1294476:]))
		if uint32(v1) >= uint32(t4) {
			goto l3
		}
		t5 := int32(load32(m.memory[int64(uint32(i32(0)))+1293932:]))
		t6 := int32(load32(m.memory[int64(uint32(i32(0)))+1294480:]))
		store32(m.memory[uint32(t5+v1<<2):], uint32(t6))
		store32(m.memory[int64(uint32(i32(0)))+1294480:], uint32(v1))
		store32(m.memory[int64(uint32(i32(0)))+1294468:], uint32(i32(0)))
	}
l1:
	v0 = v0 + i32(4)
	v2 = v2 + i32(-4)
	if v2 != 0 {
		goto l4
	}
	return
l3:
	panic("unreachable")
}
func (m *Module) X__externref_table_alloc() int32 {
	var v0, v1, v2, v3, v4 int32
	t0 := m.g0
	v0 = t0 - i32(16)
	m.g0 = v0
	{
		t1 := int32(load32(m.memory[int64(uint32(i32(0)))+1294468:]))
		if t1 != 0 {
			m.fn350(i32(1275356))
			panic("unreachable")
		}
		store32(m.memory[int64(uint32(i32(0)))+1294468:], uint32(i32(-1)))
		{
			{
				{
					t2 := int32(load32(m.memory[int64(uint32(i32(0)))+1294480:]))
					v1 = t2
					t3 := int32(load32(m.memory[int64(uint32(i32(0)))+1294476:]))
					t4 := v1
					v2 = t3
					if t4 != v2 {
						if uint32(v1) >= uint32(v2) {
							goto l4
						}
						v3 = i32(0)
						t9 := int32(load32(m.memory[int64(uint32(i32(0)))+1293932:]))
						t10 := int32(load32(m.memory[uint32(t9+v1<<2):]))
						v2 = t10
						goto l5
					}
					t5 := int32(load32(m.memory[int64(uint32(i32(0)))+1294472:]))
					t6 := v1
					v2 = t5
					if t6 != v2 {
						if uint32(v1) >= uint32(v2) {
							goto l4
						}
						t11 := int32(load32(m.memory[int64(uint32(i32(0)))+1293932:]))
						v2 = t11
						goto l6
					}
					p7 := i32(128)
					if uint32(v1) > uint32(i32(128)) {
						p7 = v1
					}
					v3 = p7
					t8 := m.fn3(v3)
					v2 = t8
					if v2 != i32(-1) {
						goto l3
					}
					goto l4
				}
			l3:
				{
					t12 := int32(load32(m.memory[int64(uint32(i32(0)))+1294484:]))
					v4 = t12
					if v4 != 0 {
						goto l7
					}
					store32(m.memory[int64(uint32(i32(0)))+1294484:], uint32(v2))
					goto l8
				}
			l7:
				if v4+v1 != v2 {
					goto l4
				}
			l8:
				t13 := int32(load32(m.memory[int64(uint32(i32(0)))+1293932:]))
				t14 := v0 + i32(4)
				t15 := v1
				v3 = v3 + v1
				m.fn1868(t14, t15, t13, v3)
				t16 := int32(load32(m.memory[int64(uint32(v0))+4:]))
				if t16 == i32(1) {
					goto l4
				}
				t17 := int32(load32(m.memory[int64(uint32(v0))+8:]))
				v2 = t17
				store32(m.memory[int64(uint32(i32(0)))+1293932:], uint32(v2))
				store32(m.memory[int64(uint32(i32(0)))+1294472:], uint32(v3))
			}
		l6:
			t18 := v2 + v1<<2
			v2 = v1 + i32(1)
			store32(m.memory[uint32(t18):], uint32(v2))
			store32(m.memory[int64(uint32(i32(0)))+1294476:], uint32(v2))
			t19 := int32(load32(m.memory[int64(uint32(i32(0)))+1294468:]))
			v3 = t19 + i32(1)
		}
	l5:
		store32(m.memory[int64(uint32(i32(0)))+1294480:], uint32(v2))
		store32(m.memory[int64(uint32(i32(0)))+1294468:], uint32(v3))
		t20 := int32(load32(m.memory[int64(uint32(i32(0)))+1294484:]))
		v2 = t20
		m.g0 = v0 + i32(16)
		return v2 + v1
	}
l4:
	panic("unreachable")
}
func (m *Module) X__externref_table_dealloc(v0 int32) {
	var v1 int32
	{
		if uint32(v0) < uint32(i32(1028)) {
			return
		}
		m.fn2(v0)
		t0 := int32(load32(m.memory[int64(uint32(i32(0)))+1294468:]))
		if t0 != 0 {
			m.fn350(i32(1275372))
			panic("unreachable")
		}
		t1 := int32(load32(m.memory[int64(uint32(i32(0)))+1294484:]))
		t2 := v0
		v1 = t1
		if uint32(t2) < uint32(v1) {
			goto l2
		}
		v0 = v0 - v1
		t3 := int32(load32(m.memory[int64(uint32(i32(0)))+1294476:]))
		if uint32(v0) >= uint32(t3) {
			goto l2
		}
		t4 := int32(load32(m.memory[int64(uint32(i32(0)))+1293932:]))
		t5 := int32(load32(m.memory[int64(uint32(i32(0)))+1294480:]))
		store32(m.memory[uint32(t4+v0<<2):], uint32(t5))
		store32(m.memory[int64(uint32(i32(0)))+1294480:], uint32(v0))
		store32(m.memory[int64(uint32(i32(0)))+1294468:], uint32(i32(0)))
	}
	return
l2:
	panic("unreachable")
}
func (m *Module) X__wbindgen_describe___wbg_Error_408e67f47ca7b58b() {
	m.fn0(i32(15))
	m.fn0(i32(0))
	m.fn0(i32(1))
	m.fn0(i32(19))
	m.fn0(i32(18))
	m.fn0(i32(24))
	m.fn0(i32(24))
}
func (m *Module) X__wbindgen_describe___wbg_now_e7c6795a7f81e10f() {
	m.fn0(i32(15))
	m.fn0(i32(0))
	m.fn0(i32(1))
	m.fn0(i32(19))
	m.fn0(i32(24))
	m.fn0(i32(13))
	m.fn0(i32(13))
}
func (m *Module) X__wbindgen_describe___wbg_Symbol_4bf93dc8964e55f5() {
	m.fn0(i32(15))
	m.fn0(i32(0))
	m.fn0(i32(1))
	m.fn0(i32(31))
	m.fn0(i32(19))
	m.fn0(i32(18))
	m.fn0(i32(24))
	m.fn0(i32(24))
}
func (m *Module) X__wbindgen_describe___wbg_get_84cd043713c6afab() {
	m.fn0(i32(15))
	m.fn0(i32(0))
	m.fn0(i32(2))
	m.fn0(i32(19))
	m.fn0(i32(24))
	m.fn0(i32(19))
	m.fn0(i32(24))
	m.fn0(i32(24))
	m.fn0(i32(24))
}
func (m *Module) X__wbindgen_describe___wbg___wbindgen_bigint_get_as_i64_c4ecf48528083721() {
	m.fn0(i32(15))
	m.fn0(i32(0))
	m.fn0(i32(1))
	m.fn0(i32(19))
	m.fn0(i32(24))
	m.fn0(i32(31))
	m.fn0(i32(6))
	m.fn0(i32(31))
	m.fn0(i32(6))
}
func (m *Module) X__wbindgen_describe___wbg_performance_3fcf6e32a7e1ed0a() {
	m.fn0(i32(15))
	m.fn0(i32(0))
	m.fn0(i32(1))
	m.fn0(i32(19))
	m.fn0(i32(24))
	m.fn0(i32(24))
	m.fn0(i32(24))
}
func (m *Module) X__wbindgen_describe___wbg_get_configurable_24b052a12692eb0f() {
	m.fn0(i32(15))
	m.fn0(i32(0))
	m.fn0(i32(1))
	m.fn0(i32(19))
	m.fn0(i32(24))
	m.fn0(i32(31))
	m.fn0(i32(14))
	m.fn0(i32(31))
	m.fn0(i32(14))
}
func (m *Module) X__wbindgen_describe___wbg___wbindgen_copy_to_typed_array_c7f28e53671b41e8() {
	m.fn0(i32(15))
	m.fn0(i32(0))
	m.fn0(i32(2))
	m.fn0(i32(19))
	m.fn0(i32(22))
	m.fn0(i32(1))
	m.fn0(i32(19))
	m.fn0(i32(24))
	m.fn0(i32(33))
	m.fn0(i32(33))
}
func (m *Module) X__wbindgen_describe___wbg___wbindgen_debug_string_a57024b9c6e4a48b() {
	m.fn0(i32(15))
	m.fn0(i32(0))
	m.fn0(i32(1))
	m.fn0(i32(19))
	m.fn0(i32(24))
	m.fn0(i32(18))
	m.fn0(i32(18))
}
func (m *Module) X__wbindgen_describe___wbg___wbindgen_module_d70c256490b5f616() {
	m.fn0(i32(15))
	m.fn0(i32(0))
	m.fn0(i32(0))
	m.fn0(i32(24))
	m.fn0(i32(24))
}
func (m *Module) X__wbindgen_describe___wbg___wbindgen_lt_94fbb50645571f95() {
	m.fn0(i32(15))
	m.fn0(i32(0))
	m.fn0(i32(2))
	m.fn0(i32(19))
	m.fn0(i32(24))
	m.fn0(i32(19))
	m.fn0(i32(24))
	m.fn0(i32(14))
	m.fn0(i32(14))
}
func (m *Module) X__wbindgen_describe___wbg___wbindgen_reinit_eaa1836ea9a8a649() {
	m.fn0(i32(15))
	m.fn0(i32(0))
	m.fn0(i32(0))
	m.fn0(i32(33))
	m.fn0(i32(33))
}
func (m *Module) X__wbindgen_describe___wbg_isArray_291e8fbbc73f8b2e() {
	m.fn0(i32(15))
	m.fn0(i32(0))
	m.fn0(i32(1))
	m.fn0(i32(19))
	m.fn0(i32(24))
	m.fn0(i32(14))
	m.fn0(i32(14))
}
func (m *Module) X__wbindgen_describe___wbg_days_2e65b0c3f09c133d() {
	m.fn0(i32(15))
	m.fn0(i32(0))
	m.fn0(i32(1))
	m.fn0(i32(19))
	m.fn0(i32(24))
	m.fn0(i32(31))
	m.fn0(i32(13))
	m.fn0(i32(31))
	m.fn0(i32(13))
}
func (m *Module) X__wbindgen_describe___wbg___wbindgen_rethrow_fbd2dcd7d2b9ac5f() {
	m.fn0(i32(15))
	m.fn0(i32(0))
	m.fn0(i32(1))
	m.fn0(i32(24))
	m.fn0(i32(33))
	m.fn0(i32(33))
}
