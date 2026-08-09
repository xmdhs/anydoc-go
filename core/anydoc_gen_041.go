package core

import (
	"math/bits"
)

func (m *Module) fn1797(v0, v1 int32) {
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
		_ = m.fn100(t4, i32(1284356), t5, t6)
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
		t14 := m.fn4(i32(12))
		v1 = t14
		if v1 != 0 {
			goto l1
		}
		m.fn85(i32(4), i32(12))
		panic("unreachable")
	}
l1:
	t15 := int32(load32(m.memory[int64(uint32(v2))+24:]))
	store32(m.memory[int64(uint32(v1))+8:], uint32(t15))
	t16 := int64(load64(m.memory[int64(uint32(v2))+16:]))
	store64(m.memory[uint32(v1):], uint64(t16))
	store32(m.memory[int64(uint32(v0))+4:], uint32(i32(1285268)))
	store32(m.memory[uint32(v0):], uint32(v1))
	m.g0 = v2 + i32(32)
}
func (m *Module) fn1798(v0, v1 int32) {
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
		_ = m.fn100(t4, i32(1284356), t5, t6)
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
	store32(m.memory[int64(uint32(v0))+4:], uint32(i32(1285268)))
	store32(m.memory[uint32(v0):], uint32(v1))
	m.g0 = v2 + i32(32)
}
func (m *Module) fn1799(v0, v1, v2 int32) {
	var v3 int32
	t0 := m.g0
	v3 = t0 - i32(16)
	m.g0 = v3
	v1 = v2 + v1
	if uint32(v1) >= uint32(v2) {
		goto l0
	}
	m.fn2(i32(0), i32(0))
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
	m.fn1800(t2, t4, t3, v2)
	{
		t8 := int32(load32(m.memory[int64(uint32(v3))+4:]))
		if t8 != i32(1) {
			goto l1
		}
		t9 := int32(load32(m.memory[int64(uint32(v3))+8:]))
		t10 := int32(load32(m.memory[int64(uint32(v3))+12:]))
		m.fn2(t9, t10)
		panic("unreachable")
	}
l1:
	t11 := int32(load32(m.memory[int64(uint32(v3))+8:]))
	v1 = t11
	store32(m.memory[uint32(v0):], uint32(v2))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
	m.g0 = v3 + i32(16)
}
func (m *Module) fn1800(v0, v1, v2, v3 int32) {
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
			t0 := m.fn89(v2, v1, i32(1), v3)
			v4 = t0
			goto l3
		}
	l2:
		t1 := m.fn4(v3)
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
func (m *Module) fn1801(v0, v1 int32) {
	t0 := int64(load64(m.memory[int64(uint32(i32(0)))+1284460:]))
	store64(m.memory[int64(uint32(v0))+8:], uint64(t0))
	t1 := int64(load64(m.memory[int64(uint32(i32(0)))+1284452:]))
	store64(m.memory[uint32(v0):], uint64(t1))
}
func (m *Module) fn1802(v0, v1, v2 int32) int32 {
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
			m.fn1799(v0, v3, v2)
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
func (m *Module) fn1803(v0, v1 int32) int32 {
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
		m.fn1799(v0, v2, v3)
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
func (m *Module) fn1804(v0, v1, v2 int32) int32 {
	t0 := m.fn100(v0, i32(1284356), v1, v2)
	return t0
}
func (m *Module) fn1805(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v1):]))
	t1 := int32(load32(m.memory[uint32(v0):]))
	t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t3 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	t4 := int32(load32(m.memory[int64(uint32(t3))+12:]))
	t5 := m.t0[uint(t4)].(func(int32, int32, int32) int32)(t0, t1, t2)
	return t5
}
func (m *Module) fn1806(v0, v1 int32) {
	var v2, v3 int32
	t0 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	v2 = t0
	t1 := int32(load32(m.memory[uint32(v1):]))
	v3 = t1
	{
		t2 := m.fn4(i32(8))
		v1 = t2
		if v1 != 0 {
			goto l0
		}
		m.fn85(i32(4), i32(8))
		panic("unreachable")
	}
l0:
	store32(m.memory[int64(uint32(v1))+4:], uint32(v2))
	store32(m.memory[uint32(v1):], uint32(v3))
	store32(m.memory[int64(uint32(v0))+4:], uint32(i32(1284672)))
	store32(m.memory[uint32(v0):], uint32(v1))
}
func (m *Module) fn1807(v0, v1 int32) {
	store32(m.memory[int64(uint32(v0))+4:], uint32(i32(1284672)))
	store32(m.memory[uint32(v0):], uint32(v1))
}
func (m *Module) fn1808(v0, v1 int32) {
	t0 := int64(load64(m.memory[uint32(v1):]))
	store64(m.memory[uint32(v0):], uint64(t0))
}
func (m *Module) fn1809(v0, v1 int32) {
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
	v3 = v2<<2 + i32(1303156)
	{
		t1 := int32(load32(m.memory[int64(uint32(i32(0)))+1303568:]))
		v4 = i32_shl(i32(1), v2)
		if t1&v4 != 0 {
			goto l1
		}
		store32(m.memory[uint32(v3):], uint32(v0))
		store32(m.memory[int64(uint32(v0))+24:], uint32(v3))
		store32(m.memory[int64(uint32(v0))+12:], uint32(v0))
		store32(m.memory[int64(uint32(v0))+8:], uint32(v0))
		t2 := int32(load32(m.memory[int64(uint32(i32(0)))+1303568:]))
		store32(m.memory[int64(uint32(i32(0)))+1303568:], uint32(t2|v4))
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
func (m *Module) fn1810(v0, v1 int32) {
	m.fn1792(v0, v1)
	panic("unreachable")
}
func (m *Module) fn1811(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t1 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	t2 := m.fn110(v1, t0, t1)
	return t2
}
func (m *Module) fn1812(v0 int32) {
	var v1 int32
	t0 := m.g0
	v1 = t0 - i32(16)
	m.g0 = v1
	store64(m.memory[uint32(v1):], uint64(int64(uint32(i32(202)))<<32|int64(uint32(v1+i32(15)))))
	m.fn91(i32(0x100caa), v1, v0)
	panic("unreachable")
}
func (m *Module) fn1813(v0, v1 int32) int32 {
	var v2, v3 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	t1 := int32(load32(m.memory[uint32(v1):]))
	t2 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	t3 := int32(load32(m.memory[int64(uint32(t2))+12:]))
	t4 := m.t0[uint(t3)].(func(int32, int32, int32) int32)(t1, i32(1285257), i32(11))
	v3 = t4
	m.memory[int64(uint32(v2))+13] = byte(i32(0))
	m.memory[int64(uint32(v2))+12] = byte(v3)
	store32(m.memory[int64(uint32(v2))+8:], uint32(v1))
	t5 := m.fn1648(v2 + i32(8))
	v1 = t5
	m.g0 = v2 + i32(16)
	return v1
}
func (m *Module) fn1814(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	v0 = t0
	t1 := int32(load32(m.memory[uint32(v0):]))
	t2 := int32(load32(m.memory[uint32(v0+i32(4)):]))
	t3 := int32(load32(m.memory[int64(uint32(t2))+12:]))
	t4 := m.t0[uint(t3)].(func(int32, int32) int32)(t1, v1)
	return t4
}
func (m *Module) fn1815(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t2 := int32(load32(m.memory[uint32(v1):]))
	t3 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	t4 := m.fn107(t0, t1, t2, t3)
	return t4
}
func (m *Module) fn1816(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t1 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	t2 := int32(load32(m.memory[uint32(v1):]))
	t3 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	t4 := m.fn107(t0, t1, t2, t3)
	return t4
}
func (m *Module) fn1817(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v1):]))
	t1 := int32(m.memory[uint32(v0)])
	v0 = t1 << 2
	t2 := int32(load32(m.memory[int64(uint32(v0))+1302324:]))
	t3 := int32(load32(m.memory[int64(uint32(v0))+1302156:]))
	t4 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	t5 := int32(load32(m.memory[int64(uint32(t4))+12:]))
	t6 := m.t0[uint(t5)].(func(int32, int32, int32) int32)(t0, t2, t3)
	return t6
}
func (m *Module) fn1818(v0, v1 int32) int32 {
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
					t4 := int32(m.memory[int64(uint32(v3&i32(15)))+1107936])
					m.memory[uint32(v2+i32(8)+v0+i32(-2))] = byte(t4)
					v0 = v0 + i32(-1)
					v3 = int32(uint32(v3) >> 4)
					if v3 != 0 {
						goto l3
					}
				}
				t5 := m.fn1638(v1, i32(1), i32(1131670), i32(2), v2+i32(8)+v0+i32(-1), i32(9)-v0)
				v0 = t5
				goto l2
			}
			if v3&i32(0x4000000) != 0 {
				goto l1
			}
			t2 := m.fn1256(v0, v1)
			v0 = t2
			goto l2
		}
	l1:
		t6 := int32(load32(m.memory[uint32(v0):]))
		v3 = t6
		v0 = i32(9)
	l4:
		{
			t7 := int32(m.memory[int64(uint32(v3&i32(15)))+1131672])
			m.memory[uint32(v2+i32(8)+v0+i32(-2))] = byte(t7)
			v0 = v0 + i32(-1)
			v3 = int32(uint32(v3) >> 4)
			if v3 != 0 {
				goto l4
			}
		}
		t8 := m.fn1638(v1, i32(1), i32(1131670), i32(2), v2+i32(8)+v0+i32(-1), i32(9)-v0)
		v0 = t8
	}
l2:
	m.g0 = v2 + i32(16)
	return v0
}
func (m *Module) fn1819(v0, v1, v2, v3 int32) {
	{
		if v2 == 0 {
			goto l0
		}
		{
			if v3 != 0 {
				goto l1
			}
			t0 := m.fn248(v2, v1)
			v1 = t0
			goto l0
		}
	l1:
		t1 := m.fn1561(v2, v1)
		v1 = t1
	}
l0:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v2))
	store32(m.memory[uint32(v0):], uint32(v1))
}
func (m *Module) fn1820(v0, v1, v2 int32) {
	var v3, v4, v5 int32
	t0 := m.g0
	v3 = t0 - i32(16)
	m.g0 = v3
	m.fn355(v3+i32(8), v2, i32(1), i32(1))
	t1 := int32(load32(m.memory[int64(uint32(v3))+8:]))
	v4 = t1
	t2 := int32(load32(m.memory[int64(uint32(v3))+12:]))
	v5 = t2
	store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
	store32(m.memory[uint32(v0):], uint32(v4))
	if v2 == 0 {
		goto l0
	}
	if v2 == 0 {
		goto l1
	}
	memory_copy(m.memory, uint32(v5), uint32(v1), uint32(v2))
l1:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v2))
l0:
	m.g0 = v3 + i32(16)
}
func (m *Module) fn1821(v0, v1 int32) {
	t0 := int64(load64(m.memory[int64(uint32(i32(0)))+1287332:]))
	store64(m.memory[int64(uint32(v0))+8:], uint64(t0))
	t1 := int64(load64(m.memory[int64(uint32(i32(0)))+1287324:]))
	store64(m.memory[uint32(v0):], uint64(t1))
}
func (m *Module) fn1822(v0, v1 int32) {
	store32(m.memory[int64(uint32(v0))+4:], uint32(i32(40)))
	store32(m.memory[uint32(v0):], uint32(i32(1287284)))
}
func fn1823(v0, v1, v2 int32) {
}
func (m *Module) fn1824(v0, v1 int32) {
	var v2 int32
	t0 := m.fn4(v1)
	v2 = t0
	store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
	store32(m.memory[uint32(v0):], uint32(v2))
}
func (m *Module) fn1825(v0, v1, v2 int32) {
	if v2 == 0 {
		return
	}
	m.fn10(v0, v2, v1)
}
func (m *Module) fn1826(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := m.fn278(t0, v1)
	return t1
}
func (m *Module) fn1827(v0, v1, v2, v3 int32) int64 {
	t0 := int32(load32(m.memory[uint32(v2-v3<<2+i32(-4)):]))
	v3 = t0
	if uint32(v3) < uint32(v1) {
		t1 := int64(load32(m.memory[int64(uint32(v0+v3*i32(192)))+184:]))
		return t1
	}
	m.fn158(v3, v1, i32(1286320))
	panic("unreachable")
}
func (m *Module) fn1828(v0, v1 int32, v2 int64) int32 {
	var v3, v4, v5 int32
	t0 := m.g0
	v3 = t0 - i32(16)
	m.g0 = v3
	v4 = v1 & int32(v2)
	v5 = i32(8)
l1:
	{
		t1 := int64(load64(m.memory[uint32(v0+v4):]))
		m.fn359(v3+i32(8), v1, t1, v4)
		t2 := int32(load32(m.memory[int64(uint32(v3))+8:]))
		if t2 != 0 {
			{
				t3 := int32(load32(m.memory[int64(uint32(v3))+12:]))
				t4 := v0
				v4 = t3
				t5 := int32(int8(m.memory[uint32(t4+v4)]))
				if t5 < i32(0) {
					goto l2
				}
				t6 := int64(load64(m.memory[uint32(v0):]))
				v4 = int32(uint32(int64(bits.TrailingZeros64(uint64(t6&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
			}
		l2:
			m.g0 = v3 + i32(16)
			return v4
		}
		v4 = (v4 + v5) & v1
		v5 = v5 + i32(8)
		goto l1
	}
}
func (m *Module) fn1829(v0, v1, v2, v3 int32) {
	var v4 int64
	var v5 int32
	v4 = int64(uint32(v1)) * int64(uint32(v3))
	if int32(int64(uint64(v4)>>32)) == 0 {
		goto l0
	}
	store32(m.memory[uint32(v0):], uint32(i32(0)))
	return
l0:
	{
		t0 := v2
		v1 = int32(v4)
		v5 = t0 + v1 + i32(-1)
		if uint32(v5) >= uint32(v1) {
			goto l1
		}
		store32(m.memory[uint32(v0):], uint32(i32(0)))
		return
	}
l1:
	{
		v3 = v3 + i32(8)
		t1 := v3
		v5 = v5 & (i32(0) - v2)
		v1 = t1 + v5
		if uint32(v1) >= uint32(v3) {
			goto l2
		}
		store32(m.memory[uint32(v0):], uint32(i32(0)))
		return
	}
l2:
	if uint32(v1) > uint32(i32(-0x80000000)-v2) {
		goto l3
	}
	store32(m.memory[int64(uint32(v0))+8:], uint32(v5))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
	store32(m.memory[uint32(v0):], uint32(v2))
	return
l3:
	store32(m.memory[uint32(v0):], uint32(i32(0)))
}
func (m *Module) fn1830(v0 int32) {
	var v1, v2 int32
	var v3, v4 int64
	t0 := m.g0
	v1 = t0 - i32(16)
	m.g0 = v1
	{
		{
			if v0 == 0 {
				goto l0
			}
			t1 := int32(load32(m.memory[uint32(v0):]))
			v2 = t1
			store64(m.memory[uint32(v0):], uint64(i64(0)))
			if v2&i32(1) == 0 {
				goto l0
			}
			t2 := int64(load64(m.memory[int64(uint32(v0))+16:]))
			v3 = t2
			t3 := int64(load64(m.memory[int64(uint32(v0))+8:]))
			v4 = t3
			goto l1
		}
	l0:
		m.fn1626(v1)
		t4 := int64(load64(m.memory[int64(uint32(v1))+8:]))
		v3 = t4
		t5 := int64(load64(m.memory[uint32(v1):]))
		v4 = t5
	}
l1:
	{
		t6 := int32(m.memory[int64(uint32(i32(0)))+1303632])
		if t6 != i32(2) {
			goto l2
		}
		m.fn91(i32(1286164), i32(125), i32(1286228))
		panic("unreachable")
	}
l2:
	m.memory[int64(uint32(i32(0)))+1303632] = byte(i32(1))
	store64(m.memory[int64(uint32(i32(0)))+1303624:], uint64(v3))
	store64(m.memory[int64(uint32(i32(0)))+1303616:], uint64(v4))
	m.g0 = v1 + i32(16)
}
func (m *Module) fn1831(v0, v1, v2 int32) {
	var v3, v4, v5, v6 int32
	t0 := m.g0
	v3 = t0 - i32(32)
	m.g0 = v3
	t1 := int32(load32(m.memory[int64(uint32(v2))+8:]))
	v4 = t1
	m.fn311(v3, v1)
	t2 := int32(load32(m.memory[int64(uint32(v3))+8:]))
	t3 := int32(load32(m.memory[int64(uint32(v3))+12:]))
	t4 := v3 + i32(16)
	v5 = t3
	t5 := int32(load32(m.memory[int64(uint32(v2))+4:]))
	t6 := v5
	t7 := v5
	v6 = t5 - v4
	p8 := v6
	if uint32(v5) < uint32(v6) {
		p8 = t7
	}
	m.fn309(t4, t2, t6, p8, i32(1287064))
	{
		t9 := int32(load32(m.memory[int64(uint32(v3))+20:]))
		t10 := v6
		v5 = t9
		if uint32(t10) >= uint32(v5) {
			goto l0
		}
		m.fn256(i32(1286600), i32(46), i32(1286648))
		panic("unreachable")
	}
l0:
	t11 := int32(load32(m.memory[uint32(v2):]))
	t12 := int32(load32(m.memory[int64(uint32(v3))+16:]))
	m.fn310(t11+v4, v5, t12, v5, i32(1286664))
	store32(m.memory[int64(uint32(v2))+8:], uint32(v5+v4))
	m.memory[uint32(v0)] = byte(i32(255))
	t13 := int64(load64(m.memory[int64(uint32(v1))+8:]))
	store64(m.memory[int64(uint32(v1))+8:], uint64(t13+int64(uint32(v5))))
	m.g0 = v3 + i32(32)
}
func (m *Module) fn1832(v0 int32) {
	var v1, v2 int32
	{
		t0 := int32(load32(m.memory[uint32(v0):]))
		v1 = t0
		if v1 == i32(-1) {
			return
		}
		t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		v2 = t1
		t2 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t3 := v1
		v0 = t2
		store32(m.memory[int64(uint32(t3))+4:], uint32(v0+i32(-1)))
		if v0 != i32(1) {
			return
		}
		m.fn1825(v1, i32(4), (v2+i32(11))&i32(-4))
	}
}
func (m *Module) fn1833(v0, v1 int32) {
	var v2 int32
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := v0
	t2 := v1
	v2 = t0
	t3 := int32(load32(m.memory[int64(uint32((t2^v2)&i32(255)<<2))+1295272:]))
	v1 = t3 ^ int32(uint32(v2)>>8)
	store32(m.memory[uint32(t1):], uint32(v1))
	t4 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t5 := v0
	v1 = (v1&i32(255)+t4)*i32(134775813) + i32(1)
	store32(m.memory[int64(uint32(t5))+4:], uint32(v1))
	t6 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	t7 := v0
	t8 := int32(uint32(v1) >> 24)
	v1 = t6
	t9 := int32(load32(m.memory[int64(uint32((t8^v1&i32(255))<<2))+1295272:]))
	store32(m.memory[int64(uint32(t7))+8:], uint32(t9^int32(uint32(v1)>>8)))
}
func (m *Module) fn1834(v0, v1 int32) {
	var v2 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	store32(m.memory[int64(uint32(v2))+12:], uint32(v1))
	store32(m.memory[int64(uint32(v2))+8:], uint32(v0))
	m.fn1632(i32(1), v2+i32(8), i32(1287792), v2+i32(12), i32(1287792), i32(0), v2, i32(1286116))
	panic("unreachable")
}
func (m *Module) fn1835(v0, v1 int32) int32 {
	var v2, v3 int32
	var v4 int64
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	{
		{
			t1 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			v3 = t1
			if v3&i32(0x2000000) != 0 {
				t3 := int64(load64(m.memory[uint32(v0):]))
				v4 = t3
				v0 = i32(17)
			l3:
				{
					t4 := int32(m.memory[int64(uint32(int32(v4)&i32(15)))+1107936])
					m.memory[uint32(v2+v0+i32(-2))] = byte(t4)
					v0 = v0 + i32(-1)
					v4 = int64(uint64(v4) >> 4)
					if v4 != i64(0) {
						goto l3
					}
				}
				t5 := m.fn1638(v1, i32(1), i32(1131670), i32(2), v2+v0+i32(-1), i32(17)-v0)
				v0 = t5
				goto l2
			}
			if v3&i32(0x4000000) != 0 {
				goto l1
			}
			t2 := m.fn579(v0, v1)
			v0 = t2
			goto l2
		}
	l1:
		t6 := int64(load64(m.memory[uint32(v0):]))
		v4 = t6
		v0 = i32(17)
	l4:
		{
			t7 := int32(m.memory[int64(uint32(int32(v4)&i32(15)))+1131672])
			m.memory[uint32(v2+v0+i32(-2))] = byte(t7)
			v0 = v0 + i32(-1)
			v4 = int64(uint64(v4) >> 4)
			if v4 != i64(0) {
				goto l4
			}
		}
		t8 := m.fn1638(v1, i32(1), i32(1131670), i32(2), v2+v0+i32(-1), i32(17)-v0)
		v0 = t8
	}
l2:
	m.g0 = v2 + i32(16)
	return v0
}
func (m *Module) fn1836(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v1):]))
	t1 := int32(load32(m.memory[uint32(v0):]))
	t2 := int32(load32(m.memory[uint32(t1):]))
	v0 = t2 << 2
	t3 := int32(load32(m.memory[uint32(v0+i32(1303064)):]))
	t4 := int32(load32(m.memory[uint32(v0+i32(1303028)):]))
	t5 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	t6 := int32(load32(m.memory[int64(uint32(t5))+12:]))
	t7 := m.t0[uint(t6)].(func(int32, int32, int32) int32)(t0, t3, t4)
	return t7
}
func (m *Module) fn1837(v0, v1 int32) {
	var v2, v3 int32
	v2 = i32(1294548)
	v3 = i32(32)
	{
		t0 := int32(m.memory[int64(uint32(v1))+160])
		switch t0 {
		default:
			goto l0
		case 1:
			v2 = v1 + i32(164)
			v3 = i32(1332)
			goto l0
		case 2:
			v2 = v1 + i32(5492)
			v3 = i32(1332)
			goto l0
		case 3:
			v2 = v1 + i32(10820)
			v3 = i32(592)
		}
	}
l0:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
	store32(m.memory[uint32(v0):], uint32(v2))
}
func (m *Module) fn1838(v0, v1, v2 int32) int32 {
	if v1 == 0 {
		goto l0
	}
l1:
	{
		t0 := int32(m.memory[uint32(v0)])
		t1 := int32(load32(m.memory[int64(uint32((t0^v2)&i32(255)<<2))+1295272:]))
		v2 = t1 ^ int32(uint32(v2)>>8)
		v0 = v0 + i32(1)
		v1 = v1 + i32(-1)
		if v1 != 0 {
			goto l1
		}
	}
l0:
	return v2
}
func (m *Module) fn1839(v0, v1, v2, v3 int32) int32 {
l1:
	{
		if v2 == 0 {
			t1 := int32(uint32(v0) % uint32(i32(65521)))
			t2 := int32(uint32(v3) % uint32(i32(65521)))
			return t1 | t2<<16
		}
		v2 = v2 + i32(-1)
		t0 := int32(m.memory[uint32(v1)])
		v0 = v0 + t0
		v3 = v0 + v3
		v1 = v1 + i32(1)
		goto l1
	}
}
func (m *Module) fn1840(v0, v1 int32) int32 {
	var v2, v3 int32
	var v4 int64
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	t1 := int32(load32(m.memory[uint32(v0):]))
	v3 = t1
	{
		t2 := int64(load64(m.memory[int64(uint32(v1))+8:]))
		v4 = t2
		v0 = int32(v4)
		if v0&i32(0x800000) == 0 {
			goto l0
		}
		if v0&i32(0x8000000) == 0 {
			goto l1
		}
		v0 = v0 | i32(0x1000000)
		goto l0
	l1:
		store16(m.memory[int64(uint32(v1))+12:], uint16(i32(10)))
		v0 = v0 | i32(0x9000000)
	}
l0:
	store32(m.memory[int64(uint32(v1))+8:], uint32(v0|i32(0x800000)))
	v0 = i32(9)
l2:
	{
		t3 := int32(m.memory[int64(uint32(v3&i32(15)))+1107936])
		m.memory[uint32(v2+i32(8)+v0+i32(-2))] = byte(t3)
		v0 = v0 + i32(-1)
		v3 = int32(uint32(v3) >> 4)
		if v3 != 0 {
			goto l2
		}
	}
	t4 := m.fn1638(v1, i32(1), i32(1131670), i32(2), v2+i32(8)+v0+i32(-1), i32(9)-v0)
	v0 = t4
	store64(m.memory[int64(uint32(v1))+8:], uint64(v4))
	m.g0 = v2 + i32(16)
	return v0
}
func (m *Module) fn1841(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v1):]))
	t1 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	t2 := int32(load32(m.memory[int64(uint32(t1))+12:]))
	t3 := m.t0[uint(t2)].(func(int32, int32, int32) int32)(t0, i32(1300856), i32(11))
	return t3
}
