package core

import (
	"math/bits"
)

func (m *Module) fn1572(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v1):]))
	t1 := int32(load32(m.memory[uint32(v0):]))
	t2 := int32(m.memory[uint32(t1)])
	v0 = t2 << 2
	t3 := int32(load32(m.memory[int64(uint32(v0))+1301796:]))
	t4 := int32(load32(m.memory[int64(uint32(v0))+1301772:]))
	t5 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	t6 := int32(load32(m.memory[int64(uint32(t5))+12:]))
	t7 := m.t0[uint(t6)].(func(int32, int32, int32) int32)(t0, t3, t4)
	return t7
}
func (m *Module) fn1573(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v1):]))
	t1 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	t2 := int32(load32(m.memory[int64(uint32(t1))+12:]))
	t3 := m.t0[uint(t2)].(func(int32, int32, int32) int32)(t0, i32(1285241), i32(5))
	return t3
}
func (m *Module) fn1574(v0, v1, v2 int32) int32 {
	m.fn75(v0, v1, v2)
	return i32(0)
}
func (m *Module) fn1575(v0, v1 int32) int32 {
	m.fn74(v0, v1)
	return i32(0)
}
func (m *Module) fn1576(v0, v1, v2 int32) {
	if v1 == 0 {
		goto l0
	}
	store32(m.memory[uint32(v0):], uint32(i32(-1)))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
	m.fn1577(v2)
	return
l0:
	t0 := int64(load64(m.memory[int64(uint32(v2))+16:]))
	store64(m.memory[int64(uint32(v0))+16:], uint64(t0))
	t1 := int64(load64(m.memory[int64(uint32(v2))+8:]))
	store64(m.memory[int64(uint32(v0))+8:], uint64(t1))
	t2 := int64(load64(m.memory[uint32(v2):]))
	store64(m.memory[uint32(v0):], uint64(t2))
}
func (m *Module) fn1577(v0 int32) {
	var v1 int32
	{
		{
			t0 := int32(load32(m.memory[uint32(v0):]))
			v1 = t0
			p1 := i32(2)
			if uint32(v1) > uint32(i32(-0x7ffffff2)) {
				p1 = v1 + i32(0x7ffffff1)
			}
			v1 = p1
			switch v1 {
			case 3, 5:
				return
			default:
				switch v1 + i32(-15) {
				case 0:
					t6 := int32(load32(m.memory[int64(uint32(v0))+4:]))
					t7 := int32(load32(m.memory[int64(uint32(v0))+8:]))
					m.fn16(t6, t7)
					return
				case 2:
					goto l8
				default:
					return
				}
			case 0:
				t2 := int32(m.memory[int64(uint32(v0))+4])
				t3 := int32(load32(m.memory[int64(uint32(v0))+8:]))
				m.fn1565(t2, t3)
				return
			case 1:
				m.fn1566(v0 + i32(4))
				return
			case 2:
				m.fn1568(v0)
				return
			case 4:
				m.fn1567(v0 + i32(4))
				return
			case 6:
				t4 := int32(load32(m.memory[int64(uint32(v0))+4:]))
				t5 := int32(load32(m.memory[int64(uint32(v0))+8:]))
				m.fn16(t4, t5)
				return
			}
		}
	l8:
		t8 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		t9 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		m.fn16(t8, t9)
	}
}
func (m *Module) fn1578(v0, v1, v2 int32) int32 {
	t0 := m.fn1851(v0, v1, v2)
	var p1 int32
	if t0 == 0 {
		p1 = 1
	}
	return p1
}
func (m *Module) fn1579(v0, v1, v2 int32) {
	var v3 int32
	t0 := m.g0
	v3 = t0 - i32(16)
	m.g0 = v3
l2:
	{
		if v2 == 0 {
			goto l0
		}
		t1 := int32(m.memory[uint32(v1)])
		if t1 == i32(48) {
			m.fn207(v3+i32(8), v1, v2, i32(1))
			t2 := int32(load32(m.memory[int64(uint32(v3))+12:]))
			v2 = t2
			t3 := int32(load32(m.memory[int64(uint32(v3))+8:]))
			v1 = t3
			goto l2
		}
	}
l0:
	store32(m.memory[uint32(v0):], uint32(v1))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v2))
	m.g0 = v3 + i32(16)
}
func (m *Module) fn1580(v0, v1, v2, v3 int32) int32 {
	if uint32(v2) < uint32(v1) {
		return v0 + v2<<3
	}
	m.fn158(v2, v1, v3)
	panic("unreachable")
}
func (m *Module) fn1581(v0, v1 int32) {
	var v2, v3, v4 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	t1 := int32(m.memory[int64(uint32(v1))+12])
	t2 := int32(m.memory[int64(uint32(v1))+13])
	t3 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	t4 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	m.fn1761(v2+i32(8), t1, t2, t3, t4)
	v3 = i32(1)
	{
		{
			t5 := int32(load32(m.memory[int64(uint32(v2))+8:]))
			if t5 == i32(1) {
				goto l0
			}
			v3 = i32(0)
			goto l1
		}
	l0:
		t6 := int32(load32(m.memory[int64(uint32(v2))+12:]))
		t7 := v1
		v4 = t6
		store32(m.memory[int64(uint32(t7))+4:], uint32(v4+i32(1)))
		t8 := int32(load32(m.memory[uint32(v1):]))
		v1 = v4 - t8
	}
l1:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
	store32(m.memory[uint32(v0):], uint32(v3))
	m.g0 = v2 + i32(16)
}
func (m *Module) fn1582(v0, v1, v2, v3, v4, v5 int32) {
	if uint32(v4) < uint32(v3) {
		goto l0
	}
	{
		if v3 == 0 {
			goto l1
		}
		if uint32(v3) < uint32(v2) {
			goto l2
		}
		if v3 != v2 {
			goto l0
		}
		goto l1
	l2:
		t0 := int32(int8(m.memory[uint32(v1+v3)]))
		if t0 <= i32(-65) {
			goto l0
		}
	}
l1:
	{
		if v4 == 0 {
			goto l3
		}
		if uint32(v4) < uint32(v2) {
			goto l4
		}
		if v4 == v2 {
			goto l3
		}
		goto l0
	l4:
		t1 := int32(int8(m.memory[uint32(v1+v4)]))
		if t1 <= i32(-65) {
			goto l0
		}
	}
l3:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v4-v3))
	store32(m.memory[uint32(v0):], uint32(v1+v3))
	return
l0:
	m.fn556(v1, v2, v3, v4, v5)
	panic("unreachable")
}
func (m *Module) fn1583(v0, v1, v2 int32) {
	var v3 int32
	m.fn1702(v0, v2)
	t0 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	v3 = t0
	{
		if v2 == 0 {
			goto l0
		}
		if v2 == 0 {
			goto l0
		}
		t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		memory_copy(m.memory, uint32(t1+v3), uint32(v1), uint32(v2))
	}
l0:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v3+v2))
}
func (m *Module) fn1584(v0, v1, v2, v3 int32) {
	{
		if v3 == 0 {
			goto l0
		}
		if uint32(v2) > uint32(v3) {
			goto l1
		}
		if v2 == v3 {
			goto l0
		}
		goto l2
	l1:
		t0 := int32(int8(m.memory[uint32(v1+v3)]))
		if t0 < i32(-64) {
			goto l2
		}
	}
l0:
	v1 = v1 + v3
	v3 = v2 - v3
	goto l3
l2:
	v1 = i32(0)
l3:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
	store32(m.memory[uint32(v0):], uint32(v1))
}
func (m *Module) fn1585(v0, v1 int32) {
	t1 := v0
	p0 := v1 + i32(-48)
	if uint32(v1) > uint32(i32(57)) {
		p0 = (v1+i32(-65))&i32(-33) + i32(10)
	}
	v1 = p0
	store32(m.memory[int64(uint32(t1))+4:], uint32(v1))
	t2 := v0
	var p3 int32
	if uint32(v1) < uint32(i32(16)) {
		p3 = 1
	}
	store32(m.memory[uint32(t2):], uint32(p3))
}
func (m *Module) fn1586(v0, v1, v2, v3, v4, v5 int32) {
	var v6 int32
	t0 := m.g0
	v6 = t0 - i32(16)
	m.g0 = v6
	store32(m.memory[int64(uint32(v6))+12:], uint32(v2))
	store32(m.memory[int64(uint32(v6))+8:], uint32(v1))
	m.fn1632(v0, v6+i32(8), i32(1107952), v6+i32(12), i32(1107952), v3, v4, v5)
	panic("unreachable")
}
func (m *Module) fn1587(v0, v1, v2 int32) {
	var v3, v4 int32
	var v5 int64
	v3 = i32(0)
	{
		if v2 == 0 {
			goto l3
		}
		t0 := int32(m.memory[uint32(v1)])
		v4 = t0
		if uint32((v4+i32(-48))&i32(255)) > uint32(i32(9)) {
			goto l3
		}
		v5 = int64(uint32(v4)) & i64(15)
		v4 = i32(1)
	l6:
		if v2 == v4 {
			store32(m.memory[uint32(v0):], uint32(i32(-1)))
			store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
			return
		}
		{
			t1 := int32(m.memory[uint32(v1+v4)])
			v3 = t1
			if uint32((v3+i32(-48))&i32(255)) <= uint32(i32(9)) {
				if uint64(v5) < uint64(i64(0x19999999)) {
					goto l4
				}
				if v5 != i64(0x19999999) {
					goto l5
				}
				if uint32(v3&i32(15)) > uint32(i32(5)) {
					goto l5
				}
			l4:
				v4 = v4 + i32(1)
				v5 = v5*i64(10) + int64(uint32(v3))&i64(15)
				goto l6
			l5:
				v5 = int64(uint32(v2))
				v3 = i32(2)
				v4 = v1
				goto l3
			}
			v3 = i32(3)
			goto l3
		}
	}
l3:
	store32(m.memory[int64(uint32(v0))+20:], uint32(v2))
	store32(m.memory[int64(uint32(v0))+16:], uint32(v1))
	store64(m.memory[int64(uint32(v0))+8:], uint64(v5))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v4))
	store32(m.memory[uint32(v0):], uint32(v3))
}
func (m *Module) fn1588(v0, v1, v2 int32) {
	var v3 int32
	t0 := m.g0
	v3 = t0 - i32(16)
	m.g0 = v3
	m.fn12(v3+i32(4), v1, v2)
	t1 := int32(load32(m.memory[int64(uint32(v3))+12:]))
	t2 := int32(load32(m.memory[int64(uint32(v3))+4:]))
	t3 := v0
	v2 = t2
	p4 := t1
	if v2 != 0 {
		p4 = i32(27)
	}
	store32(m.memory[int64(uint32(t3))+8:], uint32(p4))
	t5 := int32(load32(m.memory[int64(uint32(v3))+8:]))
	t7 := v0
	p6 := t5
	if v2 != 0 {
		p6 = i32(1087612)
	}
	store32(m.memory[int64(uint32(t7))+4:], uint32(p6))
	t9 := v0
	p8 := i32(-1)
	if v2 != 0 {
		p8 = i32(-0x7fffffdd)
	}
	store32(m.memory[uint32(t9):], uint32(p8))
	m.g0 = v3 + i32(16)
}
func (m *Module) fn1589(v0, v1 int32) int32 {
	var v2 int32
	t0 := m.g0
	v2 = t0 - i32(32)
	m.g0 = v2
	t1 := int32(load32(m.memory[uint32(v0):]))
	v0 = t1
	t2 := int32(load32(m.memory[uint32(v1):]))
	t3 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	t4 := int32(load32(m.memory[int64(uint32(t3))+12:]))
	t5 := m.t0[uint(t4)].(func(int32, int32, int32) int32)(t2, i32(1), i32(0))
	m.memory[int64(uint32(v2))+20] = byte(t5)
	store32(m.memory[int64(uint32(v2))+16:], uint32(v1))
	m.memory[int64(uint32(v2))+21] = byte(i32(1))
	store32(m.memory[int64(uint32(v2))+12:], uint32(i32(0)))
	store32(m.memory[int64(uint32(v2))+24:], uint32(v0))
	store32(m.memory[int64(uint32(v2))+28:], uint32(v0+i32(4)))
	t6 := m.fn1590(v2+i32(12), v2+i32(24), i32(84))
	t7 := m.fn1590(t6, v2+i32(28), i32(84))
	t8 := m.fn1591(t7)
	v1 = t8
	m.g0 = v2 + i32(32)
	return v1
}
func (m *Module) fn1590(v0, v1, v2 int32) int32 {
	var v3, v4, v5, v6 int32
	t0 := m.g0
	v3 = t0 - i32(32)
	m.g0 = v3
	t1 := int32(load32(m.memory[uint32(v0):]))
	v4 = t1
	v5 = i32(1)
	{
		t2 := int32(m.memory[int64(uint32(v0))+8])
		if t2 != 0 {
			goto l0
		}
		{
			t3 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			v6 = t3
			t4 := int32(m.memory[int64(uint32(v6))+10])
			if t4&i32(128) != 0 {
				goto l1
			}
			v5 = i32(1)
			t5 := int32(load32(m.memory[uint32(v6):]))
			p6 := i32(1108163)
			if v4 != 0 {
				p6 = i32(1108154)
			}
			p7 := i32(1)
			if v4 != 0 {
				p7 = i32(2)
			}
			t8 := int32(load32(m.memory[int64(uint32(v6))+4:]))
			t9 := int32(load32(m.memory[int64(uint32(t8))+12:]))
			t10 := m.t0[uint(t9)].(func(int32, int32, int32) int32)(t5, p6, p7)
			if t10 != 0 {
				goto l0
			}
			t11 := m.t0[uint(v2)].(func(int32, int32) int32)(v1, v6)
			v5 = t11
			goto l0
		}
	l1:
		{
			if v4 != 0 {
				goto l2
			}
			v5 = i32(1)
			t12 := int32(load32(m.memory[uint32(v6):]))
			t13 := int32(load32(m.memory[int64(uint32(v6))+4:]))
			t14 := int32(load32(m.memory[int64(uint32(t13))+12:]))
			t15 := m.t0[uint(t14)].(func(int32, int32, int32) int32)(t12, i32(1108164), i32(2))
			if t15 != 0 {
				goto l0
			}
		}
	l2:
		v5 = i32(1)
		m.memory[int64(uint32(v3))+15] = byte(i32(1))
		store32(m.memory[int64(uint32(v3))+20:], uint32(i32(1109040)))
		t16 := int64(load64(m.memory[uint32(v6):]))
		store64(m.memory[uint32(v3):], uint64(t16))
		t17 := int64(load64(m.memory[int64(uint32(v6))+8:]))
		store64(m.memory[int64(uint32(v3))+24:], uint64(t17))
		store32(m.memory[int64(uint32(v3))+8:], uint32(v3+i32(15)))
		store32(m.memory[int64(uint32(v3))+16:], uint32(v3))
		t18 := m.t0[uint(v2)].(func(int32, int32) int32)(v1, v3+i32(16))
		if t18 != 0 {
			goto l0
		}
		t19 := int32(load32(m.memory[int64(uint32(v3))+16:]))
		t20 := int32(load32(m.memory[int64(uint32(v3))+20:]))
		t21 := int32(load32(m.memory[int64(uint32(t20))+12:]))
		t22 := m.t0[uint(t21)].(func(int32, int32, int32) int32)(t19, i32(1108161), i32(2))
		v5 = t22
	}
l0:
	m.memory[int64(uint32(v0))+8] = byte(v5)
	store32(m.memory[uint32(v0):], uint32(v4+i32(1)))
	m.g0 = v3 + i32(32)
	return v0
}
func (m *Module) fn1591(v0 int32) int32 {
	var v1, v2, v3 int32
	t0 := int32(m.memory[int64(uint32(v0))+8])
	v1 = t0
	{
		{
			t1 := int32(load32(m.memory[uint32(v0):]))
			v2 = t1
			if v2 != 0 {
				goto l0
			}
			v3 = v1
			goto l1
		}
	l0:
		v3 = i32(1)
		{
			if v1&i32(1) != 0 {
				goto l2
			}
			t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			v1 = t2
			if v2 != i32(1) {
				goto l3
			}
			t3 := int32(m.memory[int64(uint32(v0))+9])
			if t3&i32(1) == 0 {
				goto l3
			}
			t4 := int32(m.memory[int64(uint32(v1))+10])
			if t4&i32(128) != 0 {
				goto l3
			}
			v3 = i32(1)
			t5 := int32(load32(m.memory[uint32(v1):]))
			t6 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t7 := int32(load32(m.memory[int64(uint32(t6))+12:]))
			t8 := m.t0[uint(t7)].(func(int32, int32, int32) int32)(t5, i32(1108168), i32(1))
			if t8 == 0 {
				goto l3
			}
		}
	l2:
		m.memory[int64(uint32(v0))+8] = byte(v3)
		goto l1
	l3:
		t9 := int32(load32(m.memory[uint32(v1):]))
		t10 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t11 := int32(load32(m.memory[int64(uint32(t10))+12:]))
		t12 := m.t0[uint(t11)].(func(int32, int32, int32) int32)(t9, i32(1282664), i32(1))
		t13 := v0
		v3 = t12
		m.memory[int64(uint32(t13))+8] = byte(v3)
	}
l1:
	return v3 & i32(1)
}
func (m *Module) fn1592(v0, v1 int32) int32 {
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
			t5 := m.t0[uint(v1)].(func(int32, int32, int32) int32)(v2, i32(1089054), i32(7))
			return t5
		case 1:
			t6 := m.t0[uint(v1)].(func(int32, int32, int32) int32)(v2, i32(1089079), i32(4))
			return t6
		case 2:
			t7 := m.t0[uint(v1)].(func(int32, int32, int32) int32)(v2, i32(1089068), i32(6))
			return t7
		case 3:
			t8 := m.t0[uint(v1)].(func(int32, int32, int32) int32)(v2, i32(1089048), i32(6))
			return t8
		case 4:
			t9 := m.t0[uint(v1)].(func(int32, int32, int32) int32)(v2, i32(1089074), i32(5))
			return t9
		case 5:
			t10 := m.t0[uint(v1)].(func(int32, int32, int32) int32)(v2, i32(1088624), i32(5))
			return t10
		case 6:
			t11 := m.t0[uint(v1)].(func(int32, int32, int32) int32)(v2, i32(1089061), i32(7))
			return t11
		case 7:
			t12 := m.t0[uint(v1)].(func(int32, int32, int32) int32)(v2, i32(1099979), i32(6))
			return t12
		}
	}
}
func (m *Module) fn1593(v0, v1 int32) int32 {
	var v2, v3 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	{
		t1 := int32(load32(m.memory[uint32(v0):]))
		v3 = t1
		t2 := int32(load32(m.memory[uint32(v3):]))
		v0 = t2 ^ i32(-0x80000000)
		p3 := i32(1)
		if uint32(v0) < uint32(i32(6)) {
			p3 = v0
		}
		switch p3 {
		default:
			store32(m.memory[int64(uint32(v2))+12:], uint32(v3+i32(4)))
			t4 := m.fn264(v1, i32(1100477), i32(2), v2+i32(12), i32(68))
			v1 = t4
			goto l6
		case 1:
			store32(m.memory[int64(uint32(v2))+12:], uint32(v3))
			t5 := m.fn264(v1, i32(1100479), i32(14), v2+i32(12), i32(196))
			v1 = t5
			goto l6
		case 2:
			store32(m.memory[int64(uint32(v2))+12:], uint32(v3+i32(4)))
			t6 := m.fn264(v1, i32(1100493), i32(18), v2+i32(12), i32(74))
			v1 = t6
			goto l6
		case 3:
			t7 := int32(load32(m.memory[uint32(v1):]))
			t8 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t9 := int32(load32(m.memory[int64(uint32(t8))+12:]))
			t10 := m.t0[uint(t9)].(func(int32, int32, int32) int32)(t7, i32(1100511), i32(12))
			v1 = t10
			goto l6
		case 4:
			t11 := int32(load32(m.memory[uint32(v1):]))
			t12 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t13 := int32(load32(m.memory[int64(uint32(t12))+12:]))
			t14 := m.t0[uint(t13)].(func(int32, int32, int32) int32)(t11, i32(1100523), i32(15))
			v1 = t14
			goto l6
		case 5:
			store32(m.memory[int64(uint32(v2))+12:], uint32(v3+i32(4)))
			t15 := m.fn264(v1, i32(1100538), i32(29), v2+i32(12), i32(77))
			v1 = t15
		}
	}
l6:
	m.g0 = v2 + i32(16)
	return v1
}
func (m *Module) fn1594(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	v0 = t0
	t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t2 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	t3 := int32(load32(m.memory[uint32(v1):]))
	t4 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	t5 := m.fn107(t1, t2, t3, t4)
	return t5
}
func (m *Module) fn1595(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := m.fn1069(t0, v1)
	return t1
}
func (m *Module) fn1596(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(m.memory[uint32(t0)])
	t2 := m.fn1597(t1, v1)
	return t2
}
func (m *Module) fn1597(v0, v1 int32) int32 {
	t0 := v1
	v0 = v0 & i32(255) << 2
	t1 := int32(load32(m.memory[int64(uint32(v0))+1301724:]))
	t2 := int32(load32(m.memory[int64(uint32(v0))+1301700:]))
	t3 := m.fn110(t0, t1, t2)
	return t3
}
func (m *Module) fn1598(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(m.memory[uint32(t0)])
	t2 := m.fn1599(t1, v1)
	return t2
}
func (m *Module) fn1599(v0, v1 int32) int32 {
	if v0&i32(1) == 0 {
		t1 := m.fn110(v1, i32(1131688), i32(36))
		return t1
	}
	t0 := m.fn110(v1, i32(0x1144cc), i32(21))
	return t0
}
func (m *Module) fn1600(v0, v1 int32) int32 {
	t0 := m.fn1601(v1)
	return t0
}
func (m *Module) fn1601(v0 int32) int32 {
	t0 := m.fn110(v0, i32(1131416), i32(41))
	return t0
}
func (m *Module) fn1602(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	v0 = t0
	t1 := int32(load32(m.memory[uint32(v0+i32(4)):]))
	t2 := int32(load32(m.memory[uint32(v0+i32(8)):]))
	t3 := m.fn475(t1, t2, v1)
	return t3
}
func (m *Module) fn1603(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(v1):]))
	t2 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	t3 := m.fn1604(t0, t1, t2)
	return t3
}
func (m *Module) fn1604(v0, v1, v2 int32) int32 {
	var v3 int32
	t0 := m.g0
	v3 = t0 - i32(16)
	m.g0 = v3
	{
		{
			t1 := int32(m.memory[int64(uint32(v0))+4])
			if t1 != i32(2) {
				goto l0
			}
			t2 := int32(load32(m.memory[uint32(v0):]))
			t3 := int64(load64(m.memory[int64(uint32(t2))+12:]))
			store64(m.memory[uint32(v3):], uint64(t3))
			store32(m.memory[int64(uint32(v3))+12:], uint32(i32(1)))
			store32(m.memory[int64(uint32(v3))+8:], uint32(v3))
			t4 := m.fn284(v1, v2, i32(1050619), v3+i32(8))
			v0 = t4
			goto l1
		}
	l0:
		store32(m.memory[uint32(v3):], uint32(v0))
		store32(m.memory[int64(uint32(v3))+12:], uint32(i32(197)))
		store32(m.memory[int64(uint32(v3))+8:], uint32(v3))
		t5 := m.fn284(v1, v2, i32(1052460), v3+i32(8))
		v0 = t5
	}
l1:
	m.g0 = v3 + i32(16)
	return v0
}
func (m *Module) fn1605(v0, v1 int32) int32 {
	var v2, v3, v4, v5 int32
	t0 := m.g0
	v2 = t0 - i32(48)
	m.g0 = v2
	t1 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	v3 = t1
	t2 := int32(load32(m.memory[uint32(v1):]))
	v1 = t2
	{
		t3 := int32(load32(m.memory[uint32(v0):]))
		v0 = t3
		t4 := int32(m.memory[uint32(v0)])
		switch t4 {
		default:
			store32(m.memory[int64(uint32(v2))+44:], uint32(v0+i32(4)))
			store32(m.memory[int64(uint32(v2))+24:], uint32(i32(24)))
			store32(m.memory[int64(uint32(v2))+20:], uint32(v2+i32(44)))
			t5 := m.fn284(v1, v3, i32(1051739), v2+i32(20))
			v1 = t5
			goto l7
		case 2:
			t6 := int32(load32(m.memory[int64(uint32(v3))+12:]))
			t7 := m.t0[uint(t6)].(func(int32, int32, int32) int32)(v1, i32(1100659), i32(20))
			v1 = t7
			goto l7
		case 3:
			store32(m.memory[int64(uint32(v2))+44:], uint32(v0+i32(4)))
			store32(m.memory[int64(uint32(v2))+24:], uint32(i32(36)))
			store32(m.memory[int64(uint32(v2))+20:], uint32(v2+i32(44)))
			t8 := m.fn284(v1, v3, i32(1067293), v2+i32(20))
			v1 = t8
			goto l7
		case 4:
			store32(m.memory[int64(uint32(v2))+12:], uint32(v0+i32(4)))
			store32(m.memory[int64(uint32(v2))+16:], uint32(v0+i32(12)))
			store32(m.memory[int64(uint32(v2))+44:], uint32(v0+i32(2)))
			store32(m.memory[int64(uint32(v2))+40:], uint32(i32(169)))
			store32(m.memory[int64(uint32(v2))+32:], uint32(i32(22)))
			store32(m.memory[int64(uint32(v2))+24:], uint32(i32(22)))
			store32(m.memory[int64(uint32(v2))+36:], uint32(v2+i32(44)))
			store32(m.memory[int64(uint32(v2))+28:], uint32(v2+i32(16)))
			store32(m.memory[int64(uint32(v2))+20:], uint32(v2+i32(12)))
			t9 := m.fn284(v1, v3, i32(1050947), v2+i32(20))
			v1 = t9
			goto l7
		case 5:
			store32(m.memory[int64(uint32(v2))+44:], uint32(v0+i32(2)))
			store32(m.memory[int64(uint32(v2))+24:], uint32(i32(169)))
			store32(m.memory[int64(uint32(v2))+20:], uint32(v2+i32(44)))
			t10 := m.fn284(v1, v3, i32(1068160), v2+i32(20))
			v1 = t10
			goto l7
		case 6:
			store32(m.memory[int64(uint32(v2))+44:], uint32(v0+i32(4)))
			store32(m.memory[int64(uint32(v2))+24:], uint32(i32(141)))
			store32(m.memory[int64(uint32(v2))+20:], uint32(v2+i32(44)))
			t11 := m.fn284(v1, v3, i32(1067848), v2+i32(20))
			v1 = t11
			goto l7
		case 1:
			t12 := v2
			v4 = v0 + i32(12)
			store32(m.memory[int64(uint32(t12))+16:], uint32(v4))
			{
				t13 := int32(load32(m.memory[int64(uint32(v3))+12:]))
				t14 := v1
				v5 = t13
				t15 := m.t0[uint(v5)].(func(int32, int32, int32) int32)(t14, i32(1100592), i32(13))
				if t15 != 0 {
					goto l8
				}
				{
					{
						t16 := int32(load32(m.memory[uint32(v4):]))
						v4 = t16
						if uint32(v4) < uint32(i32(512)) {
							goto l9
						}
						t17 := int32(m.memory[int64(uint32(v0))+1])
						if t17 != 0 {
							goto l10
						}
						goto l11
					}
				l9:
					store32(m.memory[int64(uint32(v2))+24:], uint32(i32(141)))
					store32(m.memory[int64(uint32(v2))+20:], uint32(v2+i32(16)))
					t18 := m.fn284(v1, v3, i32(1067364), v2+i32(20))
					if t18 != 0 {
						goto l8
					}
					t19 := int32(m.memory[int64(uint32(v0))+1])
					if t19 != i32(1) {
						goto l11
					}
					t20 := m.t0[uint(v5)].(func(int32, int32, int32) int32)(v1, i32(1108154), i32(2))
					if t20 != 0 {
						goto l8
					}
				}
			l10:
				t21 := m.t0[uint(v5)].(func(int32, int32, int32) int32)(v1, i32(1100605), i32(10))
				if t21 != 0 {
					goto l8
				}
				p22 := i32(8)
				if uint32(v4) < uint32(i32(8)) {
					p22 = v4
				}
				v4 = p22
				v0 = v0 + i32(2)
			l13:
				{
					if v4 == 0 {
						goto l12
					}
					store32(m.memory[int64(uint32(v2))+44:], uint32(v0))
					store32(m.memory[int64(uint32(v2))+24:], uint32(i32(72)))
					store32(m.memory[int64(uint32(v2))+20:], uint32(v2+i32(44)))
					t23 := m.fn284(v1, v3, i32(1100648), v2+i32(20))
					if t23 != 0 {
						goto l8
					}
					v4 = v4 + i32(-1)
					v0 = v0 + i32(1)
					goto l13
				}
			l12:
				t24 := m.t0[uint(v5)].(func(int32, int32, int32) int32)(v1, i32(1100615), i32(33))
				if t24 == 0 {
					goto l11
				}
			}
		l8:
			v1 = i32(1)
			goto l7
		l11:
			t25 := m.t0[uint(v5)].(func(int32, int32, int32) int32)(v1, i32(1282664), i32(1))
			v1 = t25
		}
	}
l7:
	m.g0 = v2 + i32(48)
	return v1
}
func (m *Module) fn1606(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := m.fn469(t0, v1)
	return t1
}
func (m *Module) fn1607(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := m.fn1669(t0, v1)
	return t1
}
func (m *Module) fn1608(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(t0):]))
	t2 := m.fn466(t1, v1)
	return t2
}
func (m *Module) fn1609(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := m.fn340(t0, v1)
	return t1
}
func (m *Module) fn1610(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := m.fn1049(t0, v1)
	return t1
}
func (m *Module) fn1611(v0, v1 int32) int32 {
	var v2, v3 int32
	t0 := m.g0
	v2 = t0 - i32(32)
	m.g0 = v2
	t1 := int32(load32(m.memory[uint32(v0):]))
	v0 = t1
	v3 = v0 + i32(4)
	{
		t2 := int32(m.memory[uint32(v0)])
		switch t2 {
		default:
			store32(m.memory[int64(uint32(v2))+12:], uint32(v3))
			store32(m.memory[int64(uint32(v2))+20:], uint32(i32(141)))
			store32(m.memory[int64(uint32(v2))+16:], uint32(v2+i32(12)))
			t3 := int32(load32(m.memory[uint32(v1):]))
			t4 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t5 := m.fn284(t3, t4, i32(1068015), v2+i32(16))
			v1 = t5
			goto l5
		case 1:
			store32(m.memory[int64(uint32(v2))+12:], uint32(v3))
			store32(m.memory[int64(uint32(v2))+20:], uint32(i32(141)))
			store32(m.memory[int64(uint32(v2))+16:], uint32(v2+i32(12)))
			t6 := int32(load32(m.memory[uint32(v1):]))
			t7 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t8 := m.fn284(t6, t7, i32(1067699), v2+i32(16))
			v1 = t8
			goto l5
		case 2:
			store32(m.memory[int64(uint32(v2))+12:], uint32(v3))
			store32(m.memory[int64(uint32(v2))+20:], uint32(i32(141)))
			store32(m.memory[int64(uint32(v2))+16:], uint32(v2+i32(12)))
			t9 := int32(load32(m.memory[uint32(v1):]))
			t10 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t11 := m.fn284(t9, t10, i32(1068315), v2+i32(16))
			v1 = t11
			goto l5
		case 3:
			store32(m.memory[int64(uint32(v2))+8:], uint32(v3))
			t12 := int32(m.memory[int64(uint32(v0))+1])
			store32(m.memory[int64(uint32(v2))+12:], uint32(t12))
			store32(m.memory[int64(uint32(v2))+28:], uint32(i32(97)))
			store32(m.memory[int64(uint32(v2))+20:], uint32(i32(141)))
			store32(m.memory[int64(uint32(v2))+24:], uint32(v2+i32(12)))
			store32(m.memory[int64(uint32(v2))+16:], uint32(v2+i32(8)))
			t13 := int32(load32(m.memory[uint32(v1):]))
			t14 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t15 := m.fn284(t13, t14, i32(1067639), v2+i32(16))
			v1 = t15
			goto l5
		case 4:
			store32(m.memory[int64(uint32(v2))+8:], uint32(v3))
			store32(m.memory[int64(uint32(v2))+12:], uint32(v0+i32(8)))
			store32(m.memory[int64(uint32(v2))+28:], uint32(i32(141)))
			store32(m.memory[int64(uint32(v2))+20:], uint32(i32(141)))
			store32(m.memory[int64(uint32(v2))+24:], uint32(v2+i32(12)))
			store32(m.memory[int64(uint32(v2))+16:], uint32(v2+i32(8)))
			t16 := int32(load32(m.memory[uint32(v1):]))
			t17 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t18 := m.fn284(t16, t17, i32(1050280), v2+i32(16))
			v1 = t18
		}
	}
l5:
	m.g0 = v2 + i32(32)
	return v1
}
func (m *Module) fn1612(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := m.fn596(t0, v1)
	return t1
}
func (m *Module) fn1613(v0, v1, v2 int32) {
	t0 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	if uint32(v1) <= uint32(t0) {
		return
	}
	_ = m.fn1614(v0, v1, v2)
}
func (m *Module) fn1614(v0, v1, v2 int32) int32 {
	var v3, v4, v5, v6, v7, v8, v9, v10, v11, v12 int32
	var v13 int64
	var v14 int32
	var v15 int64
	var v16, v17 int32
	t0 := m.g0
	v3 = t0 - i32(48)
	m.g0 = v3
	{
		t1 := int32(load32(m.memory[int64(uint32(v0))+12:]))
		v4 = t1
		v1 = v4 + v1
		if uint32(v1) < uint32(v4) {
			goto l0
		}
		{
			t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t3 := v1
			v5 = t2
			t4 := v5
			v6 = v5 + i32(1)
			v7 = int32(uint32(v6) >> 3)
			p5 := v7 * i32(7)
			if uint32(v5) < uint32(i32(8)) {
				p5 = t4
			}
			v8 = p5
			if uint32(t3) <= uint32(int32(uint32(v8)>>1)) {
				t41 := v7
				var p42 int32
				if v6&i32(7) != i32(0) {
					p42 = 1
				}
				v9 = t41 + p42
				t43 := int32(load32(m.memory[uint32(v0):]))
				v7 = t43
				v1 = v7
			l20:
				if v9 != 0 {
					t55 := int64(load64(m.memory[uint32(v1):]))
					t56 := v1
					v13 = t55
					store64(m.memory[uint32(t56):], uint64(int64(uint64(v13^i64(-1))>>7)&i64(72340172838076673)+(v13|i64(0x7f7f7f7f7f7f7f7f))))
					v1 = v1 + i32(8)
					v9 = v9 + i32(-1)
					goto l20
				}
				{
					if uint32(v6) < uint32(i32(8)) {
						goto l13
					}
					t44 := int64(load64(m.memory[uint32(v7):]))
					store64(m.memory[uint32(v7+v6):], uint64(t44))
					goto l14
				}
			l13:
				if v6 == 0 {
					goto l14
				}
				memory_copy(m.memory, uint32(v7+i32(8)), uint32(v7), uint32(v6))
			l14:
				v9 = i32(0)
			l16:
				{
					v1 = v9
					if v1 == v6 {
						store32(m.memory[int64(uint32(v0))+8:], uint32(v8-v4))
						goto l11
					}
					v9 = v1 + i32(1)
					v10 = v7 + v1
					t45 := int32(m.memory[uint32(v10)])
					if t45 != i32(128) {
						goto l16
					}
					v11 = v7 + (v1^i32(-1))<<3
				l19:
					{
						t46 := m.fn1618(v2, v7, v1)
						t47 := v1
						t48 := v5
						v13 = t46
						v12 = int32(v13)
						v14 = t48 & v12
						t49 := m.fn26(v7, v5, v13)
						t50 := t47 - v14
						v16 = t49
						if uint32((t50^(v16-v14))&v5) < uint32(i32(8)) {
							t54 := v10
							v12 = int32(uint32(v12) >> 25)
							m.memory[uint32(t54)] = byte(v12)
							m.memory[uint32(v7+(v1+i32(-8))&v5+i32(8))] = byte(v12)
							goto l16
						}
						v14 = v7 + v16
						t51 := int32(m.memory[uint32(v14)])
						v17 = t51
						t52 := v14
						v12 = int32(uint32(v12) >> 25)
						m.memory[uint32(t52)] = byte(v12)
						m.memory[uint32(v7+(v16+i32(-8))&v5+i32(8))] = byte(v12)
						v12 = v7 - v16<<3 + i32(-8)
						{
							if v17 != i32(255) {
								m.fn1619(v11, v12, i32(2))
								goto l19
							}
							m.memory[uint32(v10)] = byte(i32(255))
							m.memory[uint32(v7+(v1+i32(-8))&v5+i32(8))] = byte(i32(255))
							t53 := int64(load64(m.memory[uint32(v11):]))
							store64(m.memory[uint32(v12):], uint64(t53))
							goto l16
						}
					}
				}
			}
			{
				{
					v7 = v8 + i32(1)
					p6 := v1
					if uint32(v7) > uint32(v1) {
						p6 = v7
					}
					v1 = p6
					if uint32(v1) < uint32(i32(15)) {
						goto l2
					}
					if uint32(v1) > uint32(i32(0x1fffffff)) {
						goto l0
					}
					t7 := int32(uint32(v1<<3) / uint32(i32(7)))
					v7 = i32_shr_u(i32(-1), int32(bits.LeadingZeros32(uint32(t7+i32(-1))))) + i32(1)
					goto l3
				}
			l2:
				p8 := v1&i32(8) + i32(8)
				if uint32(v1) < uint32(i32(4)) {
					p8 = i32(4)
				}
				v7 = p8
			}
		l3:
			m.fn1616(v3+i32(8), i32(8), i32(8), v7)
			t9 := int32(load32(m.memory[int64(uint32(v3))+8:]))
			v9 = t9
			if v9 == 0 {
				goto l0
			}
			t10 := int32(load32(m.memory[int64(uint32(v3))+16:]))
			v1 = t10
			t11 := int32(load32(m.memory[int64(uint32(v3))+12:]))
			t12 := v3
			t13 := v9
			v5 = t11
			m.fn1617(t12, t13, v5)
			t14 := int32(load32(m.memory[uint32(v3):]))
			v10 = t14
			if v10 == 0 {
				m.fn85(v9, v5)
				panic("unreachable")
			}
			{
				t15 := int32(load32(m.memory[int64(uint32(v3))+4:]))
				v9 = t15
				if v9 == v5 {
					goto l5
				}
				t16 := int32(uint32(v9+i32(-8)) / uint32(i32(9)))
				t17 := v3 + i32(8)
				v7 = i32_shl(i32(1), int32(bits.LeadingZeros32(uint32(t16)))^i32(-1))
				m.fn1616(t17, i32(8), i32(8), v7)
				t18 := int32(load32(m.memory[int64(uint32(v3))+16:]))
				v1 = t18
			}
		l5:
			v5 = v0 + i32(16)
			v9 = v10 + v1
			v1 = v7 + i32(8)
			if v1 == 0 {
				goto l6
			}
			memory_fill(m.memory, uint32(v9), i32(255), uint32(v1))
		l6:
			v1 = i32(0)
			store32(m.memory[int64(uint32(v3))+32:], uint32(i32(0)))
			t19 := v3
			v6 = v7 + i32(-1)
			store32(m.memory[int64(uint32(t19))+24:], uint32(v6))
			store32(m.memory[int64(uint32(v3))+20:], uint32(v9))
			store64(m.memory[int64(uint32(v3))+12:], uint64(i64(0x800000008)))
			store32(m.memory[int64(uint32(v3))+8:], uint32(v5))
			t21 := v3
			p20 := int32(uint32(v7)>>3) * i32(7)
			if uint32(v7) < uint32(i32(9)) {
				p20 = v6
			}
			v11 = p20
			store32(m.memory[int64(uint32(t21))+28:], uint32(v11))
			t22 := int32(load32(m.memory[uint32(v0):]))
			v12 = t22
			t23 := int64(load64(m.memory[uint32(v12):]))
			v13 = (t23 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
			v8 = v3 + i32(20)
			v7 = v12
			v5 = v4
		l10:
			if v5 == 0 {
				store32(m.memory[int64(uint32(v3))+32:], uint32(v4))
				store32(m.memory[int64(uint32(v3))+28:], uint32(v11-v4))
				m.fn1619(v0, v8, i32(4))
				t34 := int32(load32(m.memory[int64(uint32(v3))+24:]))
				v1 = t34
				if v1 == 0 {
					goto l11
				}
				t35 := int32(load32(m.memory[int64(uint32(v3))+12:]))
				t36 := int32(load32(m.memory[int64(uint32(v3))+16:]))
				m.fn1616(v3+i32(36), t35, t36, v1+i32(1))
				t37 := int32(load32(m.memory[int64(uint32(v3))+40:]))
				v1 = t37
				if v1 == 0 {
					goto l11
				}
				t38 := int32(load32(m.memory[int64(uint32(v3))+20:]))
				t39 := int32(load32(m.memory[int64(uint32(v3))+44:]))
				t40 := int32(load32(m.memory[int64(uint32(v3))+36:]))
				m.fn10(t38-t39, v1, t40)
				goto l11
			}
		l9:
			{
				if v13 != i64(0) {
					t25 := v9
					t26 := v9
					t27 := v6
					t28 := v2
					t29 := v12
					v14 = int32(uint32(int64(bits.TrailingZeros64(uint64(v13))))>>3) + v1
					t30 := m.fn1618(t28, t29, v14)
					v15 = t30
					t31 := m.fn26(t26, t27, v15)
					v10 = t31
					t32 := t25 + v10
					v16 = int32(uint32(int32(v15)) >> 25)
					m.memory[uint32(t32)] = byte(v16)
					m.memory[uint32(v9+(v10+i32(-8))&v6+i32(8))] = byte(v16)
					t33 := int64(load64(m.memory[uint32(v12-v14<<3+i32(-8)):]))
					store64(m.memory[uint32(v9-v10<<3+i32(-8)):], uint64(t33))
					v5 = v5 + i32(-1)
					v13 = (v13 + i64(-1)) & v13
					goto l10
				}
				v1 = v1 + i32(8)
				v7 = v7 + i32(8)
				t24 := int64(load64(m.memory[uint32(v7):]))
				v13 = (t24 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
				goto l9
			}
		}
	}
l0:
	m.fn242()
	panic("unreachable")
l11:
	m.g0 = v3 + i32(48)
	return i32(-1)
}
func fn1615(v0, v1 int64, v2 int32) int64 {
	var v3, v4, v5, v6 int64
	t0 := v1
	v3 = int64(uint32(v2)) | i64(0x400000000000000)
	v4 = t0 ^ v3 ^ i64(8387220255154660723)
	t1 := i64_rotl(v4, i64(16))
	v4 = v4 + (v0 ^ i64(0x6c7967656e657261))
	v5 = t1 ^ v4
	t2 := i64_rotl(v5, i64(21))
	t3 := v5
	v1 = v1 ^ i64(7237128888997146477)
	v0 = v1 + (v0 ^ i64(8317987319222330741))
	v5 = t3 + i64_rotl(v0, i64(32))
	v6 = t2 ^ v5
	t4 := v6
	t5 := v4
	v0 = i64_rotl(v1, i64(13)) ^ v0
	v1 = t5 + v0
	v4 = t4 + (i64_rotl(v1, i64(32)) ^ i64(255))
	t6 := v4
	v0 = v1 ^ i64_rotl(v0, i64(17))
	t7 := i64_rotl(v0, i64(13))
	v0 = v0 + (v5 ^ v3)
	v1 = t7 ^ v0
	v3 = t6 + v1
	v1 = v3 ^ i64_rotl(v1, i64(17))
	t8 := i64_rotl(v1, i64(13))
	v4 = i64_rotl(v6, i64(16)) ^ v4
	v0 = v4 + i64_rotl(v0, i64(32))
	v1 = v0 + v1
	v5 = t8 ^ v1
	t9 := i64_rotl(v5, i64(17))
	v0 = i64_rotl(v4, i64(21)) ^ v0
	v3 = v0 + i64_rotl(v3, i64(32))
	v4 = v3 + v5
	v5 = t9 ^ v4
	t10 := i64_rotl(v5, i64(13))
	v0 = i64_rotl(v0, i64(16)) ^ v3
	v1 = v0 + i64_rotl(v1, i64(32))
	v3 = t10 ^ (v1 + v5)
	t11 := i64_rotl(v3, i64(17))
	v0 = i64_rotl(v0, i64(21)) ^ v1
	v1 = v0 + i64_rotl(v4, i64(32))
	v3 = v1 + v3
	return t11 ^ i64_rotl(v3, i64(32)) ^ i64_rotl(i64_rotl(v0, i64(16))^v1, i64(21)) ^ v3
}
func (m *Module) fn1616(v0, v1, v2, v3 int32) {
	var v4 int64
	var v5 int32
	{
		v4 = int64(uint32(v1)) * int64(uint32(v3))
		if int32(int64(uint64(v4)>>32)) != 0 {
			goto l0
		}
		t0 := v2
		v1 = int32(v4)
		v5 = t0 + v1 + i32(-1)
		if uint32(v5) < uint32(v1) {
			goto l0
		}
		v1 = v3 + i32(8)
		t1 := v1
		v5 = v5 & (i32(0) - v2)
		v3 = t1 + v5
		if uint32(v3) < uint32(v1) {
			goto l1
		}
		if uint32(v3) > uint32(i32(-0x80000000)-v2) {
			store32(m.memory[uint32(v0):], uint32(i32(0)))
			return
		}
		store32(m.memory[int64(uint32(v0))+8:], uint32(v5))
		store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
		store32(m.memory[uint32(v0):], uint32(v2))
		return
	}
l0:
	store32(m.memory[uint32(v0):], uint32(i32(0)))
	return
l1:
	store32(m.memory[uint32(v0):], uint32(i32(0)))
}
