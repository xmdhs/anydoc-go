package core

import (
	"math"
	"math/bits"
)

func (m *Module) fn1482(v0, v1, v2 int32, v3 int64) int64 {
	var v4, v5 int32
	t0 := m.g0
	v4 = t0 - i32(16)
	m.g0 = v4
	m.fn1483(v4, v0, v1)
	{
		t1 := m.fn1484(v4)
		v1 = t1
		v0 = v1 + i32(72)
		t3 := v0
		p2 := i32(8)
		if uint32(v2) < uint32(i32(8)) {
			p2 = v2
		}
		v2 = p2
		v5 = t3 + v2
		t4 := int32(m.memory[uint32(v5)])
		if t4 != i32(1) {
			goto l0
		}
		t5 := int64(load64(m.memory[uint32(v1+v2<<3):]))
		v3 = t5 + i64(1)
		p6 := v3
		if v3 == 0 {
			p6 = i64(-1)
		}
		v3 = p6
	}
l0:
	m.memory[uint32(v5)] = byte(i32(1))
	store64(m.memory[uint32(v1+v2<<3):], uint64(v3))
	store32(m.memory[int64(uint32(v4))+8:], uint32(v2+i32(1)))
	store32(m.memory[int64(uint32(v4))+4:], uint32(v1+i32(81)))
	store32(m.memory[uint32(v4):], uint32(v0))
l2:
	{
		t7 := m.fn1370(v4)
		v1 = t7
		if v1 == 0 {
			m.g0 = v4 + i32(16)
			return v3
		}
		m.memory[uint32(v1)] = byte(i32(0))
		goto l2
	}
}
func (m *Module) fn1483(v0, v1, v2 int32) {
	var v3 int64
	var v4, v5, v6 int32
	var v7 int64
	var v8, v9, v10 int32
	var v11, v12 int64
	var v13 int32
	t0 := int64(load64(m.memory[int64(uint32(v1))+16:]))
	t1 := int64(load64(m.memory[int64(uint32(v1))+24:]))
	t2 := m.fn66(t0, t1, v2)
	v3 = t2
	t3 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	v4 = t3
	t4 := v4
	v5 = int32(v3)
	v6 = t4 & v5
	v7 = int64(uint64(v3)>>25) & i64(127) * i64(72340172838076673)
	v8 = v1 + i32(16)
	t5 := int32(load32(m.memory[uint32(v1):]))
	v9 = t5
	v10 = i32(0)
l6:
	{
		t6 := int64(load64(m.memory[uint32(v9+v6):]))
		v11 = t6
		v12 = v11 ^ v7
		v12 = (v12 ^ i64(-1)) & (v12 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
	l2:
		{
			if v12 == 0 {
				if v11&(v11<<1)&i64(-0x7f7f7f7f7f7f7f80) == 0 {
					t10 := v6
					v10 = v10 + i32(8)
					v6 = (t10 + v10) & v4
					goto l6
				}
				{
					t8 := int32(load32(m.memory[int64(uint32(v1))+8:]))
					if t8 != 0 {
						goto l4
					}
					_ = m.fn721(v1, v8)
				}
			l4:
				store32(m.memory[int64(uint32(v0))+8:], uint32(v2))
				store32(m.memory[uint32(v0):], uint32(v5))
				store32(m.memory[int64(uint32(v0))+4:], uint32(int64(uint64(v3)>>32)))
				goto l5
			}
			v13 = v9 + (i32(0)-(int32(uint32(int64(bits.TrailingZeros64(uint64(v12))))>>3)+v6)&v4)*i32(96)
			t7 := int32(load32(m.memory[uint32(v13+i32(-96)):]))
			if t7 == v2 {
				goto l1
			}
			v12 = (v12 + i64(-1)) & v12
			goto l2
		}
	l1:
		store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
		store32(m.memory[uint32(v0):], uint32(v13))
		v1 = i32(0)
	l5:
		store32(m.memory[int64(uint32(v0))+12:], uint32(v1))
		return
	}
}
func (m *Module) fn1484(v0 int32) int32 {
	var v1, v2, v3 int32
	var v4 int64
	var v5, v6, v7, v8, v9 int32
	t0 := m.g0
	v1 = t0 - i32(96)
	m.g0 = v1
	{
		{
			t1 := int32(load32(m.memory[int64(uint32(v0))+12:]))
			v2 = t1
			if v2 == 0 {
				goto l0
			}
			t2 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			v3 = t2
			t3 := int64(load64(m.memory[uint32(v0):]))
			v4 = t3
			memory_zero(m.memory, uint32(v1+i32(12)), uint32(i32(81)))
			t4 := int32(load32(m.memory[uint32(v2):]))
			v0 = t4
			t5 := int32(load32(m.memory[int64(uint32(v2))+4:]))
			t6 := v0
			t7 := v0
			v5 = t5
			t8 := m.fn26(t7, v5, v4)
			v6 = t8
			v7 = t6 + v6
			t9 := int32(m.memory[uint32(v7)])
			v8 = t9
			t10 := v7
			v9 = int32(uint32(int32(v4)) >> 25)
			m.memory[uint32(t10)] = byte(v9)
			m.memory[uint32(v0+v5&(v6+i32(-8))+i32(8))] = byte(v9)
			t11 := int32(load32(m.memory[int64(uint32(v2))+8:]))
			store32(m.memory[int64(uint32(v2))+8:], uint32(t11-v8&i32(1)))
			v0 = v0 + (i32(0)-v6)*i32(96)
			store32(m.memory[uint32(v0+i32(-96)):], uint32(v3))
			memory_copy(m.memory, uint32(v0+i32(-92)), uint32(v1+i32(8)), uint32(i32(85)))
			t12 := int32(load32(m.memory[int64(uint32(v2))+12:]))
			store32(m.memory[int64(uint32(v2))+12:], uint32(t12+i32(1)))
			goto l1
		}
	l0:
		t13 := int32(load32(m.memory[uint32(v0):]))
		v0 = t13
	}
l1:
	m.g0 = v1 + i32(96)
	return v0 + i32(-88)
}
func (m *Module) fn1485(v0, v1 int32) {
	var v2, v3, v4, v5 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	v3 = v0 + i32(12)
l4:
	{
		t1 := int32(load32(m.memory[int64(uint32(v0))+20:]))
		if uint32(t1) < uint32(v1) {
			store32(m.memory[int64(uint32(v2))+12:], uint32(i32(0)))
			store64(m.memory[int64(uint32(v2))+4:], uint64(i64(0x800000000)))
			m.fn1169(v3, v2+i32(4))
			goto l4
		}
		t2 := int32(load32(m.memory[int64(uint32(v0))+32:]))
		v3 = t2
		p3 := v1
		if uint32(v3) > uint32(v1) {
			p3 = v3
		}
		v4 = p3
		v5 = v0 + i32(24)
		v1 = v3 << 4
	l3:
		{
			if v4 == v3 {
				goto l1
			}
			{
				t4 := int32(load32(m.memory[uint32(v5):]))
				if v3 != t4 {
					goto l2
				}
				m.fn223(v5)
			}
		l2:
			t5 := v0
			v3 = v3 + i32(1)
			store32(m.memory[int64(uint32(t5))+32:], uint32(v3))
			t6 := int32(load32(m.memory[int64(uint32(v0))+28:]))
			store32(m.memory[uint32(t6+v1):], uint32(i32(0)))
			v1 = v1 + i32(16)
			goto l3
		}
	}
l1:
	m.g0 = v2 + i32(16)
}
func (m *Module) fn1486(v0, v1, v2, v3 int32) int32 {
	if uint32(v2) < uint32(v1) {
		return v0 + v2<<6
	}
	m.fn158(v2, v1, v3)
	panic("unreachable")
}
func (m *Module) fn1487(v0, v1 int32) {
	var v2 int32
	m.fn1485(v0, v1)
	t0 := int32(load32(m.memory[int64(uint32(v0))+16:]))
	t1 := int32(load32(m.memory[int64(uint32(v0))+20:]))
	v1 = v1 + i32(-1)
	t2 := m.fn857(t0, t1, v1, i32(1075896))
	v2 = t2
	t3 := int32(load32(m.memory[int64(uint32(v0))+28:]))
	t4 := int32(load32(m.memory[int64(uint32(v0))+32:]))
	t5 := m.fn1491(t3, t4, v1, i32(1075912))
	m.fn1333(t5, v2)
}
func (m *Module) fn1488(v0, v1 int32) int32 {
	m.fn1485(v0, v1)
	t0 := int32(load32(m.memory[int64(uint32(v0))+16:]))
	t1 := int32(load32(m.memory[int64(uint32(v0))+20:]))
	t2 := m.fn857(t0, t1, v1+i32(-1), i32(1075880))
	return t2
}
func (m *Module) fn1489(v0, v1, v2, v3 int32) {
	var v4, v5 int32
	t0 := m.g0
	v4 = t0 - i32(80)
	m.g0 = v4
	{
		if uint32(v2) < uint32(v3) {
			goto l0
		}
		v3 = v3 + i32(-1)
		if uint32(v3) >= uint32(v2) {
			m.fn158(v3, v2, i32(1075848))
			panic("unreachable")
		}
		t1 := int32(load32(m.memory[int64(uint32(v1+v3<<6))+40:]))
		if t1 == 0 {
			goto l0
		}
		t2 := m.fn1486(v1, v2, v3, i32(1075864))
		v3 = t2
		v2 = i32(0)
		store32(m.memory[int64(uint32(v4))+75:], uint32(i32(0)))
		store64(m.memory[int64(uint32(v4))+67:], uint64(i64(0)))
		memory_copy(m.memory, uint32(v4), uint32(v3), uint32(i32(64)))
		m.memory[int64(uint32(v3))+12] = byte(i32(0))
		store32(m.memory[int64(uint32(v3))+8:], uint32(i32(0)))
		store64(m.memory[uint32(v3):], uint64(i64(0x800000000)))
		store32(m.memory[int64(uint32(v3))+32:], uint32(i32(0)))
		store64(m.memory[int64(uint32(v3))+36:], uint64(i64(4)))
		t3 := int64(load64(m.memory[int64(uint32(v4))+71:]))
		store64(m.memory[int64(uint32(v3))+20:], uint64(t3))
		t4 := int64(load64(m.memory[int64(uint32(v4))+64:]))
		store64(m.memory[int64(uint32(v3))+13:], uint64(t4))
		store64(m.memory[int64(uint32(v3))+52:], uint64(i64(0)))
		store64(m.memory[int64(uint32(v3))+44:], uint64(i64(0x800000000)))
		m.memory[int64(uint32(v3))+60] = byte(i32(0))
		t5 := int32(load32(m.memory[int64(uint32(v4))+40:]))
		v3 = t5 << 4
		t6 := int32(load32(m.memory[int64(uint32(v4))+36:]))
		v1 = t6
		t7 := int32(load32(m.memory[int64(uint32(v4))+32:]))
		v5 = t7
	l3:
		if v3 == v2 {
			m.fn419(i32(0), i32(4))
			store32(m.memory[int64(uint32(v4))+64:], uint32(v5))
			store32(m.memory[int64(uint32(v4))+68:], uint32(v1))
			store32(m.memory[int64(uint32(v4))+72:], uint32(int32(uint32(v1+v2-v1)>>4)))
			m.fn419(i32(0), i32(4))
			m.fn1452(v0, v4+i32(64))
			t8 := int32(load32(m.memory[int64(uint32(v4))+44:]))
			t9 := int32(load32(m.memory[int64(uint32(v4))+48:]))
			m.fn1490(t8, t9)
			m.fn968(v4)
			goto l4
		}
		v2 = v2 + i32(16)
		goto l3
	}
l0:
	store32(m.memory[uint32(v0):], uint32(i32(-1)))
	goto l4
l4:
	m.g0 = v4 + i32(80)
}
func (m *Module) fn1490(v0, v1 int32) {
	m.fn136(v0, v1, i32(8), i32(16))
}
func (m *Module) fn1491(v0, v1, v2, v3 int32) int32 {
	if uint32(v2) < uint32(v1) {
		return v0 + v2<<4
	}
	m.fn158(v2, v1, v3)
	panic("unreachable")
}
func (m *Module) fn1492(v0, v1, v2, v3, v4, v5 int32) int32 {
	var v6, v7, v8, v9, v10 int32
	t0 := m.g0
	v6 = t0 - i32(16)
	m.g0 = v6
	{
		if v4 == 0 {
			goto l0
		}
		v7 = v1 * i32(240)
		v8 = v0 + i32(-240)
	l1:
		{
			if v7 == 0 {
				goto l0
			}
			m.fn855(v6+i32(8), v8+i32(456))
			v7 = v7 + i32(-240)
			v8 = v8 + i32(240)
			t1 := int32(load32(m.memory[int64(uint32(v6))+8:]))
			t2 := int32(load32(m.memory[int64(uint32(v6))+12:]))
			t3 := m.fn848(t1, t2, v4, v5)
			if t3 == 0 {
				goto l1
			}
			goto l2
		}
	l0:
		v7 = v1 * i32(240)
		t4 := m.fn1364(v2, v3)
		v4 = t4
		v8 = v0
	l8:
		if v7 != 0 {
			{
				t8 := int32(load32(m.memory[uint32(v8+i32(232)):]))
				t9 := v4
				v5 = t8
				t10 := int32(load32(m.memory[uint32(v8+i32(236)):]))
				t11 := v5
				v9 = t10
				t12 := m.fn1364(t11, v9)
				v10 = t12
				if t9&v10 != i32(1) {
					goto l6
				}
				t13 := m.fn773(v5, v9, v2, v3)
				if t13 == 0 {
					goto l7
				}
				goto l2
			}
		l6:
			if v4 == v10 {
				goto l2
			}
		l7:
			v8 = v8 + i32(240)
			v7 = v7 + i32(-240)
			goto l8
		}
		v7 = v1 * i32(240)
		v8 = v0 + i32(236)
	l5:
		{
			if v7 != 0 {
				goto l4
			}
			v8 = i32(0)
			goto l2
		l4:
			v7 = v7 + i32(-240)
			v5 = v8 + i32(-4)
			t5 := int32(load32(m.memory[uint32(v8):]))
			v9 = t5
			v10 = v8 + i32(240)
			v8 = v10
			t6 := int32(load32(m.memory[uint32(v5):]))
			t7 := m.fn1364(t6, v9)
			if v4 != t7 {
				goto l5
			}
		}
		v8 = v10 + i32(-476)
		goto l2
	}
l2:
	m.g0 = v6 + i32(16)
	return v8
}
func (m *Module) fn1493(v0, v1 int32) int32 {
	var v2, v3, v4 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	t1 := m.fn1495(v0, v1, i32(1073720))
	v3 = t1
	t2 := m.fn1495(v0, v1, i32(1073721))
	v4 = t2
	m.fn1046(v2+i32(8), v0, v1, i32(1074411), i32(53), i32(1073722), i32(6))
	v1 = v3 & i32(255)
	v0 = v4 & i32(255)
	{
		{
			t3 := int32(load32(m.memory[int64(uint32(v2))+8:]))
			v3 = t3
			if v3 != 0 {
				goto l0
			}
			v3 = i32(0x20000)
			goto l1
		}
	l0:
		{
			t4 := int32(load32(m.memory[int64(uint32(v2))+12:]))
			t5 := v3
			v4 = t4
			t6 := m.fn15(t5, v4, i32(1074913), i32(9))
			if t6 == 0 {
				goto l2
			}
			v3 = i32(65536)
			goto l1
		}
	l2:
		t7 := m.fn15(v3, v4, i32(1074922), i32(9))
		p8 := i32(0)
		if t7 != 0 {
			p8 = i32(65536)
		}
		v3 = p8
	}
l1:
	m.g0 = v2 + i32(16)
	return v3 | v0<<8 | v1 | i32(0x2000000)
}
func (m *Module) fn1494(v0, v1, v2, v3 int32) {
	var v4 int32
	t0 := m.g0
	v4 = t0 - i32(16)
	m.g0 = v4
	if v1 != 0 {
		goto l0
	}
	m.fn51(v0+i32(4), v2, v3)
	store32(m.memory[uint32(v0):], uint32(i32(1)))
	goto l1
l0:
	m.fn51(v4+i32(4), v2, v3)
	m.fn1454(v0, v4+i32(4))
l1:
	m.g0 = v4 + i32(16)
}
func (m *Module) fn1495(v0, v1, v2 int32) int32 {
	var v3, v4 int32
	t0 := m.g0
	v3 = t0 - i32(16)
	m.g0 = v3
	v4 = i32(1)
	m.fn1046(v3+i32(8), v0, v1, i32(1074411), i32(53), v2, i32(1))
	{
		{
			t1 := int32(load32(m.memory[int64(uint32(v3))+8:]))
			v2 = t1
			if v2 != 0 {
				goto l0
			}
			v4 = i32(2)
			goto l1
		}
	l0:
		t2 := int32(load32(m.memory[int64(uint32(v3))+12:]))
		t3 := v2
		v1 = t2
		t4 := m.fn15(t3, v1, i32(1073318), i32(1))
		if t4 != 0 {
			goto l1
		}
		t5 := m.fn15(v2, v1, i32(1071691), i32(4))
		if t5 != 0 {
			goto l1
		}
		t6 := m.fn15(v2, v1, i32(1073319), i32(2))
		v4 = t6
	}
l1:
	m.g0 = v3 + i32(16)
	return v4
}
func (m *Module) fn1496(v0, v1, v2, v3, v4, v5, v6 int32) {
	var v7 int32
	t0 := m.g0
	v7 = t0 - i32(64)
	m.g0 = v7
	{
		t1 := m.fn846(v2, v5, v6)
		v6 = t1
		if v6 == 0 {
			store64(m.memory[uint32(v0):], uint64(i64(-1)))
			goto l4
		}
		t2 := int32(m.memory[int64(uint32(v6))+24])
		if t2 != 0 {
			goto l1
		}
		t3 := int32(load32(m.memory[int64(uint32(v6))+4:]))
		t4 := int32(load32(m.memory[int64(uint32(v6))+8:]))
		m.fn774(v7+i32(36), v3, v4, t3, t4)
		v6 = v7 + i32(40)
		{
			t5 := int32(load32(m.memory[int64(uint32(v7))+36:]))
			if t5 != 0 {
				store64(m.memory[uint32(v0):], uint64(i64(-1)))
				m.fn785(v6)
				goto l4
			}
			t6 := int64(load64(m.memory[int64(uint32(v6))+16:]))
			store64(m.memory[int64(uint32(v7))+24:], uint64(t6))
			t7 := int64(load64(m.memory[int64(uint32(v6))+8:]))
			store64(m.memory[int64(uint32(v7))+16:], uint64(t7))
			t8 := int64(load64(m.memory[uint32(v6):]))
			store64(m.memory[int64(uint32(v7))+8:], uint64(t8))
			m.fn1182(v7, v1, i32(1084140))
			t9 := int32(load32(m.memory[int64(uint32(v7))+4:]))
			v6 = t9
			t10 := int32(load32(m.memory[uint32(v7):]))
			t11 := int32(load32(m.memory[int64(uint32(v7))+12:]))
			t12 := v7 + i32(36)
			v3 = t11
			t13 := int32(load32(m.memory[int64(uint32(v7))+16:]))
			m.fn1035(t12, t10, v3, t13)
			t14 := int32(load32(m.memory[int64(uint32(v7))+44:]))
			v2 = t14
			t15 := int32(load32(m.memory[int64(uint32(v7))+40:]))
			v5 = t15
			{
				t16 := int32(load32(m.memory[int64(uint32(v7))+36:]))
				v4 = t16
				if v4 == i32(-1) {
					v4 = v0 + i32(4)
					{
						{
							if v5 == 0 {
								goto l5
							}
							t20 := int32(load32(m.memory[int64(uint32(v7))+16:]))
							store32(m.memory[int64(uint32(v4))+8:], uint32(t20))
							t21 := int64(load64(m.memory[int64(uint32(v7))+8:]))
							store64(m.memory[uint32(v4):], uint64(t21))
							store32(m.memory[int64(uint32(v0))+20:], uint32(v2))
							store32(m.memory[int64(uint32(v0))+16:], uint32(v5))
							t22 := int32(load32(m.memory[uint32(v6):]))
							store32(m.memory[uint32(v6):], uint32(t22+i32(1)))
							goto l6
						}
					l5:
						store32(m.memory[uint32(v4):], uint32(i32(-1)))
						t23 := int32(load32(m.memory[uint32(v6):]))
						store32(m.memory[uint32(v6):], uint32(t23+i32(1)))
						t24 := int32(load32(m.memory[int64(uint32(v7))+8:]))
						m.fn16(t24, v3)
					}
				l6:
					store32(m.memory[uint32(v0):], uint32(i32(-1)))
					t25 := int32(load32(m.memory[int64(uint32(v7))+20:]))
					t26 := int32(load32(m.memory[int64(uint32(v7))+24:]))
					m.fn134(t25, t26)
					goto l4
				}
				t17 := int32(load32(m.memory[int64(uint32(v7))+56:]))
				store32(m.memory[int64(uint32(v0))+20:], uint32(t17))
				t18 := int64(load64(m.memory[int64(uint32(v7))+48:]))
				store64(m.memory[int64(uint32(v0))+12:], uint64(t18))
				store32(m.memory[int64(uint32(v0))+8:], uint32(v2))
				store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
				store32(m.memory[uint32(v0):], uint32(v4))
				t19 := int32(load32(m.memory[uint32(v6):]))
				store32(m.memory[uint32(v6):], uint32(t19+i32(1)))
				m.fn784(v7 + i32(8))
				goto l4
			}
		}
	}
l1:
	store64(m.memory[uint32(v0):], uint64(i64(-1)))
l4:
	m.g0 = v7 + i32(64)
}
func (m *Module) fn1497(v0, v1, v2 int32) {
	var v3, v4 int32
	t0 := m.g0
	v3 = t0 - i32(64)
	m.g0 = v3
	store32(m.memory[int64(uint32(v3))+20:], uint32(i32(0)))
	store64(m.memory[int64(uint32(v3))+12:], uint64(i64(0x400000000)))
	m.fn868(v3+i32(24), v1, v2)
	store32(m.memory[int64(uint32(v3))+48:], uint32(i32(1)))
	store32(m.memory[int64(uint32(v3))+44:], uint32(i32(1077161)))
	store32(m.memory[int64(uint32(v3))+40:], uint32(i32(53)))
	store32(m.memory[int64(uint32(v3))+36:], uint32(i32(1074411)))
l2:
	{
		t1 := m.fn863(v3 + i32(24))
		v4 = t1
		if v4 == 0 {
			t8 := int32(load32(m.memory[int64(uint32(v3))+24:]))
			t9 := int32(load32(m.memory[int64(uint32(v3))+28:]))
			m.fn44(t8, t9)
			{
				t10 := int32(load32(m.memory[int64(uint32(v3))+20:]))
				v4 = t10
				if v4 == 0 {
					goto l3
				}
				t11 := int32(load32(m.memory[int64(uint32(v3))+16:]))
				m.fn77(v0, t11, v4, i32(1097368), i32(1))
				goto l4
			}
		l3:
			m.fn864(v0, v1, v2)
		l4:
			m.fn78(v3 + i32(12))
			m.g0 = v3 + i32(64)
			return
		}
		t2 := int32(load32(m.memory[uint32(v4+i32(28)):]))
		t3 := int32(load32(m.memory[uint32(v4+i32(32)):]))
		m.fn864(v3+i32(52), t2, t3)
		t4 := int32(load32(m.memory[int64(uint32(v3))+56:]))
		t5 := v3
		v4 = t4
		t6 := int32(load32(m.memory[int64(uint32(v3))+60:]))
		m.fn46(t5, v4, t6)
		t7 := int32(load32(m.memory[int64(uint32(v3))+4:]))
		if t7 == 0 {
			t12 := int32(load32(m.memory[int64(uint32(v3))+52:]))
			m.fn16(t12, v4)
			goto l2
		}
		m.fn33(v3+i32(12), v3+i32(52))
		goto l2
	}
}
func (m *Module) fn1498(v0 int32) {
	t0 := int32(load32(m.memory[uint32(v0):]))
	if t0 == i32(-1) {
		return
	}
	m.fn1499(v0)
}
func (m *Module) fn1499(v0 int32) {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	m.fn136(t0, t1, i32(4), i32(8))
	t2 := int32(load32(m.memory[int64(uint32(v0))+12:]))
	t3 := int32(load32(m.memory[int64(uint32(v0))+16:]))
	m.fn136(t2, t3, i32(4), i32(8))
}
func (m *Module) fn1500(v0, v1, v2 int32) int32 {
	var v3, v4 int32
	v3 = i32(255)
	{
		t0 := v1
		v2 = v2 & i32(0xffff)
		if uint32(t0) <= uint32(v2) {
			goto l0
		}
		v3 = v0 + v2*i32(3)
		t1 := int32(load16(m.memory[int64(uint32(v3))+1:]))
		v4 = t1
		t2 := int32(m.memory[uint32(v3)])
		v3 = t2
	}
l0:
	t3 := v4 << 8
	var p4 int32
	if v3&i32(255) == i32(255) {
		p4 = 1
	}
	v4 = p4
	p5 := t3
	if v4 != 0 {
		p5 = i32(131584)
	}
	p6 := v3
	if v4 != 0 {
		p6 = i32(2)
	}
	return p5 | p6&i32(255)
}
func (m *Module) fn1501(v0, v1 int32) {
	var v2 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		v2 = t0
		t1 := int32(load32(m.memory[uint32(v0):]))
		if v2 != t1 {
			goto l0
		}
		m.fn223(v0)
	}
l0:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v2+i32(1)))
	t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	v0 = t2 + v2<<4
	t3 := int64(load64(m.memory[uint32(v1):]))
	store64(m.memory[uint32(v0):], uint64(t3))
	t4 := int64(load64(m.memory[int64(uint32(v1))+8:]))
	store64(m.memory[int64(uint32(v0))+8:], uint64(t4))
}
func (m *Module) fn1502(v0, v1 int32) {
	var v2 int32
	v2 = v0 + i32(16)
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+16:]))
		if t0 == i32(-1) {
			goto l0
		}
		t1 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t2 := v2
		v0 = t1
		t3 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		m.fn75(t2, v0, t3)
		t4 := int32(load32(m.memory[uint32(v1):]))
		m.fn16(t4, v0)
		return
	}
l0:
	m.fn1303(v2)
	m.memory[int64(uint32(v0))+52] = byte(i32(1))
	store32(m.memory[int64(uint32(v0))+28:], uint32(i32(-1)))
	t5 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	store32(m.memory[int64(uint32(v2))+8:], uint32(t5))
	t6 := int64(load64(m.memory[uint32(v1):]))
	store64(m.memory[uint32(v2):], uint64(t6))
}
func (m *Module) fn1503(v0, v1 int32, v2 int64) {
	var v3 int64
	t0 := int64(load64(m.memory[uint32(v1):]))
	t1 := v1
	v3 = t0
	v2 = v3 + v2
	p2 := v2
	if uint64(v2) < uint64(v3) {
		p2 = i64(-1)
	}
	v2 = p2
	store64(m.memory[uint32(t1):], uint64(v2))
	v1 = i32(-1)
	if uint64(v2) < uint64(i64(4000001)) {
		goto l0
	}
	m.fn51(v0+i32(4), i32(1075731), i32(49))
	store32(m.memory[int64(uint32(v0))+20:], uint32(i32(13)))
	store32(m.memory[int64(uint32(v0))+16:], uint32(i32(1075780)))
	v1 = i32(-0x7ffffffd)
l0:
	store32(m.memory[uint32(v0):], uint32(v1))
}
func (m *Module) fn1504(v0, v1 int32) {
	var v2 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		v2 = t0
		t1 := int32(load32(m.memory[uint32(v0):]))
		if v2 != t1 {
			goto l0
		}
		m.fn1149(v0)
	}
l0:
	t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	memory_copy(m.memory, uint32(t2+v2*i32(40)), uint32(v1), uint32(i32(40)))
	store32(m.memory[int64(uint32(v0))+8:], uint32(v2+i32(1)))
}
func (m *Module) fn1505(v0 int32) {
	var v1, v2, v3 int32
	t0 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	v1 = t0
	t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	v2 = t1
	v3 = v2
l2:
	if v1 == 0 {
		goto l0
	}
	{
		t2 := int32(load32(m.memory[uint32(v3):]))
		if t2 == 0 {
			goto l1
		}
		m.fn969(v3 + i32(12))
	}
l1:
	v1 = v1 + i32(-1)
	v3 = v3 + i32(40)
	goto l2
l0:
	t3 := int32(load32(m.memory[uint32(v0):]))
	m.fn136(t3, v2, i32(8), i32(40))
}
func (m *Module) fn1506(v0, v1, v2 int32) {
	var v3 int32
	var v4 int64
	var v5 int32
	t0 := m.g0
	v3 = t0 - i32(48)
	m.g0 = v3
	{
		t1 := int64(load64(m.memory[uint32(v2):]))
		v4 = t1
		if v4 == 0 {
			goto l0
		}
		m.fn1503(v3+i32(4), v1, v4)
		{
			t2 := int32(load32(m.memory[int64(uint32(v3))+4:]))
			v5 = t2
			if v5 == i32(-1) {
				goto l1
			}
			t3 := int32(load32(m.memory[int64(uint32(v3))+24:]))
			store32(m.memory[int64(uint32(v0))+20:], uint32(t3))
			t4 := int64(load64(m.memory[int64(uint32(v3))+16:]))
			store64(m.memory[int64(uint32(v0))+12:], uint64(t4))
			t5 := int64(load64(m.memory[int64(uint32(v3))+8:]))
			store64(m.memory[int64(uint32(v0))+4:], uint64(t5))
			goto l2
		}
	l1:
		v1 = v1 + i32(24)
	l5:
		if v4 == 0 {
			goto l3
		}
		store32(m.memory[int64(uint32(v3))+44:], uint32(i32(0)))
		store64(m.memory[int64(uint32(v3))+36:], uint64(i64(0)))
		store64(m.memory[int64(uint32(v3))+28:], uint64(i64(0x800000000)))
		m.fn1167(v3+i32(4), v1, v3+i32(28))
		{
			t6 := int32(load32(m.memory[int64(uint32(v3))+4:]))
			v5 = t6
			if v5 == i32(-1) {
				v4 = v4 + i64(-1)
				goto l5
			}
			t7 := int32(load32(m.memory[int64(uint32(v3))+24:]))
			store32(m.memory[int64(uint32(v0))+20:], uint32(t7))
			t8 := int64(load64(m.memory[int64(uint32(v3))+16:]))
			store64(m.memory[int64(uint32(v0))+12:], uint64(t8))
			t9 := int64(load64(m.memory[int64(uint32(v3))+8:]))
			store64(m.memory[int64(uint32(v0))+4:], uint64(t9))
			goto l2
		}
	l3:
		store64(m.memory[uint32(v2):], uint64(i64(0)))
	}
l0:
	v5 = i32(-1)
l2:
	store32(m.memory[uint32(v0):], uint32(v5))
	m.g0 = v3 + i32(48)
}
func (m *Module) fn1507(v0, v1 int32) {
	var v2, v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14 int32
	var v15 int64
	var v16 int32
	var v17 int64
	var v18, v19, v20, v21, v22, v23, v24, v25, v26, v27, v28, v29, v30, v31 int32
	var v32 int64
	t0 := m.g0
	v2 = t0 - i32(80)
	m.g0 = v2
	t1 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	v3 = t1
	t2 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	t3 := v2 + i32(32)
	v4 = t2
	m.fn59(t3, v4, i32(8), i32(32))
	v5 = v3 + v4<<5
	v6 = v2 + i32(40) + i32(12)
	v7 = i32(0)
	t4 := int32(load32(m.memory[int64(uint32(v2))+36:]))
	v8 = t4
	v9 = v2 + i32(75)
	t5 := int32(load32(m.memory[int64(uint32(v2))+32:]))
	v10 = t5
	v11 = v10
l16:
	if v11 == 0 {
		goto l0
	}
	if v3 == v5 {
		goto l0
	}
	v12 = v8 + v7<<5
	v1 = i32(-0x7ffffffb)
	{
		t6 := int32(load32(m.memory[uint32(v3):]))
		v13 = t6
		switch v13 >> 31 & (v13 + i32(-0x7fffffff)) {
		case 6:
			goto l7
		default:
			t7 := int32(m.memory[int64(uint32(v3))+24])
			v14 = t7
			m.fn225(v2+i32(64), v3+i32(12))
			m.fn1509(v2+i32(40), v3)
			t8 := int64(load32(m.memory[int64(uint32(v2))+64:]))
			t9 := int64(load32(m.memory[int64(uint32(v2))+48:]))
			v15 = t8<<32 | t9
			t10 := int32(load32(m.memory[int64(uint32(v2))+44:]))
			v16 = t10
			t11 := int32(load32(m.memory[int64(uint32(v2))+40:]))
			v1 = t11
			t12 := int64(load64(m.memory[int64(uint32(v2))+68:]))
			v17 = t12
			goto l7
		case 1:
			m.fn1509(v2+i32(40), v3+i32(4))
			t13 := int64(load64(m.memory[int64(uint32(v2))+44:]))
			v15 = t13
			t14 := int32(load32(m.memory[int64(uint32(v2))+40:]))
			v16 = t14
			v1 = i32(-0x80000000)
			goto l7
		case 2:
			t15 := int64(load64(m.memory[int64(uint32(v3))+8:]))
			v15 = t15
			t16 := int32(m.memory[int64(uint32(v3))+28])
			v18 = t16
			t17 := int32(load32(m.memory[int64(uint32(v3))+20:]))
			v13 = t17
			t18 := int32(load32(m.memory[int64(uint32(v3))+24:]))
			t19 := v2 + i32(8)
			v14 = t18
			m.fn59(t19, v14, i32(4), i32(28))
			t20 := int32(load32(m.memory[int64(uint32(v2))+8:]))
			v19 = t20
			t21 := v19
			v1 = v14 & i32(0x3fffffff)
			p22 := v1
			if uint32(v19) < uint32(v1) {
				p22 = t21
			}
			v20 = p22
			t23 := int32(load32(m.memory[int64(uint32(v2))+12:]))
			v21 = t23
			v1 = v21
		l9:
			{
				if v20 == 0 {
					v23 = int32(uint32(v14) >> 8)
					v17 = int64(uint32(v21))<<32 | int64(uint32(v19))
					v1 = i32(-0x7fffffff)
					goto l7
				}
				m.fn1507(v2+i32(64), v13)
				t24 := int32(m.memory[uint32(v13+i32(24))])
				v22 = t24
				m.fn225(v6, v13+i32(12))
				t25 := int32(load32(m.memory[int64(uint32(v2))+72:]))
				store32(m.memory[int64(uint32(v2))+48:], uint32(t25))
				t26 := int64(load64(m.memory[int64(uint32(v2))+64:]))
				t27 := v2
				v17 = t26
				store64(m.memory[int64(uint32(t27))+40:], uint64(v17))
				t28 := int64(load64(m.memory[int64(uint32(v2))+56:]))
				store64(m.memory[int64(uint32(v1))+16:], uint64(t28))
				t29 := int64(load64(m.memory[int64(uint32(v2))+48:]))
				store64(m.memory[int64(uint32(v1))+8:], uint64(t29))
				store64(m.memory[uint32(v1):], uint64(v17))
				m.memory[uint32(v1+i32(24))] = byte(v22)
				v20 = v20 + i32(-1)
				v1 = v1 + i32(28)
				v13 = v13 + i32(28)
				goto l9
			}
		case 3:
			t30 := int32(load32(m.memory[int64(uint32(v3))+8:]))
			v24 = t30
			t31 := int32(load32(m.memory[int64(uint32(v3))+12:]))
			t32 := v2 + i32(24)
			v25 = t31
			m.fn59(t32, v25, i32(4), i32(12))
			t33 := int32(load32(m.memory[int64(uint32(v2))+24:]))
			v16 = t33
			t34 := v16
			v1 = v25 & i32(0x3fffffff)
			p35 := v1
			if uint32(v16) < uint32(v1) {
				p35 = t34
			}
			v26 = p35
			v1 = i32(0)
			t36 := int32(load32(m.memory[int64(uint32(v2))+28:]))
			v27 = t36
		l15:
			{
				if v1 == v26 {
					v15 = int64(uint32(v25))<<32 | int64(uint32(v27))
					v1 = i32(-0x7ffffffe)
					t48 := int64(m.memory[int64(uint32(v3))+20])
					t49 := int64(load32(m.memory[int64(uint32(v3))+16:]))
					v17 = t48<<32 | (v32&i64(-0x10000000000) | t49)
					v32 = v17
					goto l7
				}
				t37 := int32(load32(m.memory[int64(uint32(v24))+4:]))
				v21 = t37
				t38 := int32(load32(m.memory[int64(uint32(v24))+8:]))
				t39 := v2 + i32(16)
				v28 = t38
				m.fn59(t39, v28, i32(4), i32(20))
				v24 = v24 + i32(12)
				v29 = v1 + i32(1)
				v19 = v28 * i32(20)
				v30 = v27 + v1*i32(12)
				v1 = i32(0)
				t40 := int32(load32(m.memory[int64(uint32(v2))+20:]))
				v22 = t40
				t41 := int32(load32(m.memory[int64(uint32(v2))+16:]))
				v31 = t41
				v20 = v31
			l14:
				{
					if v20 == 0 {
						goto l11
					}
					if v19 == v1 {
						goto l11
					}
					{
						v13 = v21 + v1
						t42 := int32(load32(m.memory[uint32(v13):]))
						if t42 != i32(-1) {
							goto l12
						}
						t43 := int32(load32(m.memory[int64(uint32(v13))+8:]))
						store32(m.memory[int64(uint32(v2))+48:], uint32(t43))
						t44 := int64(load64(m.memory[uint32(v13):]))
						store64(m.memory[int64(uint32(v2))+40:], uint64(t44))
						goto l13
					}
				l12:
					m.fn1507(v2+i32(40), v13)
				l13:
					t45 := int64(load64(m.memory[uint32(v13+i32(12)):]))
					v17 = t45
					v13 = v22 + v1
					t46 := int64(load64(m.memory[int64(uint32(v2))+40:]))
					store64(m.memory[uint32(v13):], uint64(t46))
					t47 := int32(load32(m.memory[int64(uint32(v2))+48:]))
					store32(m.memory[int64(uint32(v13))+8:], uint32(t47))
					store64(m.memory[uint32(v13+i32(12)):], uint64(v17))
					v20 = v20 + i32(-1)
					v1 = v1 + i32(20)
					goto l14
				}
			l11:
				store32(m.memory[int64(uint32(v30))+8:], uint32(v28))
				store32(m.memory[int64(uint32(v30))+4:], uint32(v22))
				store32(m.memory[uint32(v30):], uint32(v31))
				v1 = v29
				goto l15
			}
		case 4:
			m.fn1507(v2+i32(40), v3+i32(4))
			t50 := int64(load64(m.memory[int64(uint32(v2))+44:]))
			v15 = t50
			t51 := int32(load32(m.memory[int64(uint32(v2))+40:]))
			v16 = t51
			v1 = i32(-0x7ffffffd)
			goto l7
		case 5:
			m.fn225(v2+i32(64), v3+i32(16))
			t52 := int32(load32(m.memory[uint32(v3+i32(8)):]))
			t53 := int32(load32(m.memory[uint32(v3+i32(12)):]))
			m.fn31(v2+i32(40), t52, t53)
			t54 := int32(load16(m.memory[int64(uint32(v2))+73:]))
			t55 := int32(m.memory[uint32(v9)])
			v23 = t54 | t55<<16
			t56 := int64(load64(m.memory[int64(uint32(v2))+44:]))
			v15 = t56
			t57 := int32(load32(m.memory[int64(uint32(v2))+40:]))
			v16 = t57
			t58 := int32(m.memory[int64(uint32(v2))+72])
			v14 = t58
			t59 := int64(load64(m.memory[int64(uint32(v2))+64:]))
			v17 = t59
			v1 = i32(-0x7ffffffc)
		}
	}
l7:
	v11 = v11 + i32(-1)
	v7 = v7 + i32(1)
	v3 = v3 + i32(32)
	store16(m.memory[int64(uint32(v12))+25:], uint16(v23))
	m.memory[int64(uint32(v12))+28] = byte(v18)
	m.memory[int64(uint32(v12))+24] = byte(v14)
	store64(m.memory[int64(uint32(v12))+16:], uint64(v17))
	store64(m.memory[int64(uint32(v12))+8:], uint64(v15))
	store32(m.memory[int64(uint32(v12))+4:], uint32(v16))
	store32(m.memory[uint32(v12):], uint32(v1))
	m.memory[uint32(v12+i32(27))] = byte(int32(uint32(v23) >> 16))
	goto l16
l0:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v4))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v8))
	store32(m.memory[uint32(v0):], uint32(v10))
	m.g0 = v2 + i32(80)
}
func (m *Module) fn1508(v0 int32, v1 float64) {
	var v2 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	store64(m.memory[uint32(v2):], math.Float64bits(v1))
	store32(m.memory[int64(uint32(v2))+12:], uint32(i32(66)))
	store32(m.memory[int64(uint32(v2))+8:], uint32(v2))
	m.fn73(v0, i32(1052692), v2+i32(8))
	m.g0 = v2 + i32(16)
}
func (m *Module) fn1509(v0, v1 int32) {
	var v2, v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13 int32
	var v14 int64
	t0 := m.g0
	v2 = t0 - i32(48)
	m.g0 = v2
	t1 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	v3 = t1
	t2 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	t3 := v2
	v4 = t2
	m.fn59(t3, v4, i32(4), i32(28))
	v5 = v4 * i32(28)
	v6 = v2 + i32(20) + i32(4)
	v1 = i32(0)
	t4 := int32(load32(m.memory[int64(uint32(v2))+4:]))
	v7 = t4
	t5 := int32(load32(m.memory[uint32(v2):]))
	v8 = t5
	v9 = v8
l10:
	{
		if v9 == 0 {
			goto l0
		}
		if v5 == v1 {
			goto l0
		}
		v10 = i32(8)
		{
			v11 = v3 + v1
			t6 := int32(load32(m.memory[uint32(v11):]))
			v12 = t6
			p7 := i32(1)
			if uint32(v12) > uint32(i32(2)) {
				p7 = v12 + i32(-3)
			}
			switch p7 {
			case 5:
				goto l6
			default:
				t8 := int32(load32(m.memory[uint32(v11+i32(8)):]))
				t9 := int32(load32(m.memory[uint32(v11+i32(12)):]))
				m.fn31(v2+i32(8), t8, t9)
				t10 := int32(load32(m.memory[uint32(v11+i32(16)):]))
				v13 = t10
				v10 = i32(3)
				goto l6
			case 1:
				m.fn1509(v2+i32(36), v11+i32(16))
				t11 := int32(load32(m.memory[uint32(v11):]))
				v10 = t11
				t12 := int32(load32(m.memory[uint32(v11+i32(8)):]))
				t13 := int32(load32(m.memory[uint32(v11+i32(12)):]))
				m.fn31(v6, t12, t13)
				t14 := int64(load64(m.memory[uint32(v6):]))
				store64(m.memory[int64(uint32(v2))+8:], uint64(t14))
				t15 := int32(load32(m.memory[int64(uint32(v6))+8:]))
				store32(m.memory[int64(uint32(v2))+16:], uint32(t15))
				t16 := int64(load64(m.memory[int64(uint32(v2))+40:]))
				v14 = t16
				t17 := int32(load32(m.memory[int64(uint32(v2))+36:]))
				v13 = t17
				goto l6
			case 2:
				t18 := int32(load32(m.memory[uint32(v11+i32(8)):]))
				t19 := int32(load32(m.memory[uint32(v11+i32(12)):]))
				m.fn31(v2+i32(36), t18, t19)
				v13 = i32(-0x7fffffff)
				{
					t20 := int32(load32(m.memory[uint32(v11+i32(16)):]))
					v10 = t20
					switch v10 >> 31 & (v10 + i32(-0x7fffffff)) {
					case 2:
						goto l9
					default:
						t21 := int32(load32(m.memory[uint32(v11+i32(20)):]))
						t22 := int32(load32(m.memory[uint32(v11+i32(24)):]))
						m.fn31(v2+i32(20), t21, t22)
						t23 := int32(load32(m.memory[int64(uint32(v2))+20:]))
						v13 = t23
						goto l9
					case 1:
						t24 := int32(load32(m.memory[uint32(v11+i32(20)):]))
						store32(m.memory[int64(uint32(v2))+24:], uint32(t24))
						v13 = i32(-0x80000000)
					}
				}
			l9:
				t25 := int32(load32(m.memory[int64(uint32(v2))+44:]))
				store32(m.memory[int64(uint32(v2))+16:], uint32(t25))
				t26 := int64(load64(m.memory[int64(uint32(v2))+36:]))
				store64(m.memory[int64(uint32(v2))+8:], uint64(t26))
				t27 := int64(load64(m.memory[int64(uint32(v2))+24:]))
				v14 = t27
				v10 = i32(5)
				goto l6
			case 3:
				t28 := int32(load32(m.memory[uint32(v11+i32(8)):]))
				t29 := int32(load32(m.memory[uint32(v11+i32(12)):]))
				m.fn31(v2+i32(8), t28, t29)
				v10 = i32(6)
				goto l6
			case 4:
				t30 := int32(load32(m.memory[uint32(v11+i32(8)):]))
				t31 := int32(load32(m.memory[uint32(v11+i32(12)):]))
				m.fn31(v2+i32(8), t30, t31)
				v10 = i32(7)
			}
		}
	l6:
		v11 = v7 + v1
		store32(m.memory[uint32(v11):], uint32(v10))
		t32 := int64(load64(m.memory[int64(uint32(v2))+8:]))
		store64(m.memory[uint32(v11+i32(4)):], uint64(t32))
		t33 := int32(load32(m.memory[int64(uint32(v2))+16:]))
		store32(m.memory[uint32(v11+i32(12)):], uint32(t33))
		store64(m.memory[uint32(v11+i32(20)):], uint64(v14))
		store32(m.memory[uint32(v11+i32(16)):], uint32(v13))
		v9 = v9 + i32(-1)
		v1 = v1 + i32(28)
		goto l10
	}
l0:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v4))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v7))
	store32(m.memory[uint32(v0):], uint32(v8))
	m.g0 = v2 + i32(48)
}
func (m *Module) fn1510(v0, v1, v2 int32) {
	var v3, v4, v5 int32
	t0 := m.g0
	v3 = t0 - i32(48)
	m.g0 = v3
	t1 := int32(load32(m.memory[int64(uint32(v1))+16:]))
	t2 := int32(load32(m.memory[int64(uint32(v1))+20:]))
	t3 := int32(load32(m.memory[uint32(v2+i32(200)):]))
	m.fn1513(v3+i32(24), t1, t2, t3)
	t4 := int32(load32(m.memory[int64(uint32(v3))+28:]))
	v4 = t4
	{
		{
			t5 := int32(load32(m.memory[int64(uint32(v3))+24:]))
			v5 = t5
			if v5 == i32(-1) {
				goto l0
			}
			t6 := int64(load64(m.memory[int64(uint32(v3))+40:]))
			store64(m.memory[int64(uint32(v0))+20:], uint64(t6))
			t7 := int64(load64(m.memory[int64(uint32(v3))+32:]))
			store64(m.memory[int64(uint32(v0))+12:], uint64(t7))
			store32(m.memory[int64(uint32(v0))+8:], uint32(v4))
			store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
			store32(m.memory[uint32(v0):], uint32(i32(1)))
			goto l1
		}
	l0:
		store32(m.memory[int64(uint32(v3))+8:], uint32(i32(0)))
		store64(m.memory[uint32(v3):], uint64(i64(0x400000000)))
		store32(m.memory[int64(uint32(v3))+20:], uint32(i32(0)))
		store64(m.memory[int64(uint32(v3))+12:], uint64(i64(0x800000000)))
		m.fn1516(v3+i32(24), v1, v2, v4, v3, v3+i32(12))
		{
			t8 := int32(load32(m.memory[int64(uint32(v3))+24:]))
			v1 = t8
			if v1 == i32(-1) {
				goto l2
			}
			t9 := int32(load32(m.memory[int64(uint32(v3))+44:]))
			store32(m.memory[int64(uint32(v0))+24:], uint32(t9))
			t10 := int64(load64(m.memory[int64(uint32(v3))+36:]))
			store64(m.memory[int64(uint32(v0))+16:], uint64(t10))
			t11 := int64(load64(m.memory[int64(uint32(v3))+28:]))
			store64(m.memory[int64(uint32(v0))+8:], uint64(t11))
			store32(m.memory[uint32(v0):], uint32(i32(1)))
			store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
			m.fn969(v3 + i32(12))
			m.fn894(v3)
			goto l1
		}
	l2:
		v1 = v0 + i32(4)
		t12 := int64(load64(m.memory[uint32(v3):]))
		store64(m.memory[uint32(v1):], uint64(t12))
		store32(m.memory[uint32(v0):], uint32(i32(0)))
		t13 := int32(load32(m.memory[int64(uint32(v3))+8:]))
		store32(m.memory[int64(uint32(v3))+32:], uint32(t13))
		t14 := int64(load64(m.memory[int64(uint32(v3))+12:]))
		store64(m.memory[int64(uint32(v3))+36:], uint64(t14))
		t15 := int64(load64(m.memory[int64(uint32(v3))+32:]))
		store64(m.memory[int64(uint32(v1))+8:], uint64(t15))
		t16 := int32(load32(m.memory[int64(uint32(v3))+20:]))
		store32(m.memory[int64(uint32(v3))+44:], uint32(t16))
		t17 := int64(load64(m.memory[int64(uint32(v3))+40:]))
		store64(m.memory[int64(uint32(v1))+16:], uint64(t17))
	}
l1:
	m.g0 = v3 + i32(48)
}
func (m *Module) fn1511(v0, v1, v2, v3, v4, v5, v6, v7 int32) {
	var v8, v9, v10, v11, v12, v13, v14, v15 int32
	var v16 int64
	var v17 int32
	var v18 int64
	var v19 int32
	var v20 int64
	var v21, v22 int32
	var v23 int64
	var v24, v25, v26, v27, v28, v29 int32
	t0 := m.g0
	v8 = t0 - i32(304)
	m.g0 = v8
	t1 := int32(load32(m.memory[int64(uint32(v1))+16:]))
	t2 := v8 + i32(64)
	v9 = t1
	t3 := int32(load32(m.memory[int64(uint32(v1))+20:]))
	t4 := v9
	v10 = t3
	m.fn1046(t2, t4, v10, i32(1074680), i32(46), i32(1085052), i32(10))
	v11 = i32(0)
	v12 = i32(1)
	{
		t5 := int32(load32(m.memory[int64(uint32(v2))+200:]))
		t6 := int32(load32(m.memory[int64(uint32(v8))+64:]))
		t7 := t5 + i32(432)
		v13 = t6
		p8 := v4
		if v13 != 0 {
			p8 = v13
		}
		v4 = p8
		p9 := i32(1)
		if v4 != 0 {
			p9 = v4
		}
		v14 = p9
		t10 := int32(load32(m.memory[int64(uint32(v8))+68:]))
		t12 := v14
		p11 := v5
		if v13 != 0 {
			p11 = t10
		}
		v5 = p11
		p13 := i32(0)
		if v4 != 0 {
			p13 = v5
		}
		v13 = p13
		t14 := m.fn1517(t7, t12, v13)
		v15 = t14
		if v15 == 0 {
			goto l0
		}
		if uint32(v3) >= uint32(i32(10)) {
			goto l0
		}
		v15 = v15 + v3*i32(40)
		t15 := int64(load64(m.memory[uint32(v15):]))
		v16 = t15
		t16 := int32(m.memory[int64(uint32(v15))+36])
		v17 = t16
		t17 := int32(load32(m.memory[uint32(v15+i32(12)):]))
		t18 := int32(load32(m.memory[uint32(v15+i32(16)):]))
		m.fn31(v8+i32(256), t17, t18)
		m.fn225(v8+i32(248)+i32(20), v15+i32(20))
		m.memory[int64(uint32(v8))+284] = byte(v17)
		t19 := int32(load32(m.memory[int64(uint32(v15))+32:]))
		store32(m.memory[int64(uint32(v8))+280:], uint32(t19))
		t20 := int64(load64(m.memory[int64(uint32(v8))+260:]))
		store64(m.memory[int64(uint32(v8))+192:], uint64(t20))
		t21 := int64(load64(m.memory[int64(uint32(v8))+268:]))
		store64(m.memory[int64(uint32(v8))+200:], uint64(t21))
		t22 := int32(load32(m.memory[int64(uint32(v8))+284:]))
		store32(m.memory[int64(uint32(v8))+216:], uint32(t22))
		t23 := int64(load64(m.memory[int64(uint32(v8))+276:]))
		store64(m.memory[int64(uint32(v8))+208:], uint64(t23))
		t24 := int32(load32(m.memory[int64(uint32(v8))+256:]))
		v15 = t24
		if v15 == i32(-1) {
			goto l0
		}
		store32(m.memory[int64(uint32(v8))+256:], uint32(v15))
		store64(m.memory[int64(uint32(v8))+248:], uint64(v16))
		t25 := int64(load64(m.memory[int64(uint32(v8))+192:]))
		store64(m.memory[int64(uint32(v8))+260:], uint64(t25))
		t26 := int64(load64(m.memory[int64(uint32(v8))+200:]))
		store64(m.memory[int64(uint32(v8))+268:], uint64(t26))
		t27 := int64(load64(m.memory[int64(uint32(v8))+208:]))
		store64(m.memory[int64(uint32(v8))+276:], uint64(t27))
		t28 := int32(load32(m.memory[int64(uint32(v8))+216:]))
		t29 := v8
		v11 = t28
		store32(m.memory[int64(uint32(t29))+284:], uint32(v11))
		var p30 int32
		if v11&i32(255) == 0 {
			p30 = 1
		}
		v12 = p30
		goto l1
	}
l0:
	store32(m.memory[int64(uint32(v8))+280:], uint32(i32(1)))
	store64(m.memory[int64(uint32(v8))+264:], uint64(i64(-0x100000000)))
	store64(m.memory[int64(uint32(v8))+256:], uint64(i64(0x100000000)))
	v16 = i64(1)
	store64(m.memory[int64(uint32(v8))+248:], uint64(i64(1)))
	m.memory[int64(uint32(v8))+284] = byte(i32(0))
l1:
	m.fn51(v8+i32(76), v14, v13)
	store32(m.memory[int64(uint32(v8))+88:], uint32(v3))
	if v12 != 0 {
		goto l2
	}
	m.fn1046(v8+i32(56), v9, v10, i32(1074680), i32(46), i32(1085062), i32(13))
	{
		t31 := int32(load32(m.memory[int64(uint32(v8))+56:]))
		v14 = t31
		if v14 == 0 {
			m.fn1046(v8+i32(48), v9, v10, i32(1074680), i32(46), i32(1085092), i32(18))
			t48 := int32(load32(m.memory[int64(uint32(v8))+48:]))
			t49 := int32(load32(m.memory[int64(uint32(v8))+52:]))
			t50 := m.fn848(t48, t49, i32(1071691), i32(4))
			if t50 == 0 {
				goto l2
			}
			t51 := int32(load32(m.memory[int64(uint32(v2))+16:]))
			v13 = t51
			if uint32(v13) >= uint32(i32(0x7fffffff)) {
				m.fn1518(i32(1085112))
				panic("unreachable")
			}
			store32(m.memory[int64(uint32(v2))+16:], uint32(v13+i32(1)))
			t52 := int32(load32(m.memory[int64(uint32(v2))+36:]))
			if t52 == 0 {
				goto l11
			}
			t53 := int64(load64(m.memory[int64(uint32(v2))+40:]))
			t54 := int64(load64(m.memory[uint32(v2+i32(48)):]))
			t55 := m.fn696(t53, t54, v8+i32(76))
			v18 = t55
			t56 := int32(load32(m.memory[int64(uint32(v2))+28:]))
			v15 = t56
			v13 = v15 & int32(v18)
			v20 = int64(uint64(v18)>>25) & i64(127) * i64(72340172838076673)
			t57 := int32(load32(m.memory[int64(uint32(v2))+24:]))
			v14 = t57
			v21 = i32(0)
		l16:
			{
				t58 := int64(load64(m.memory[uint32(v14+v13):]))
				v23 = t58
				v18 = v23 ^ v20
				v18 = (v18 ^ i64(-1)) & (v18 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
			l14:
				{
					if v18 == 0 {
						if v23&(v23<<1)&i64(-0x7f7f7f7f7f7f7f80) == 0 {
							t62 := v13
							v21 = v21 + i32(8)
							v13 = (t62 + v21) & v15
							goto l16
						}
						t61 := int32(load32(m.memory[int64(uint32(v2))+16:]))
						v13 = t61 + i32(-1)
						goto l11
					}
					t59 := v8 + i32(76)
					v17 = v14 + (i32(0)-(int32(uint32(int64(bits.TrailingZeros64(uint64(v18))))>>3)+v13)&v15)*i32(24)
					t60 := m.fn1117(t59, v17+i32(-24))
					if t60 != 0 {
						goto l13
					}
					v18 = (v18 + i64(-1)) & v18
					goto l14
				}
			l13:
			}
			t63 := int64(load64(m.memory[uint32(v17+i32(-8)):]))
			v16 = t63
			t64 := int32(load32(m.memory[int64(uint32(v2))+16:]))
			store32(m.memory[int64(uint32(v2))+16:], uint32(t64+i32(-1)))
			goto l2
		}
		t32 := int32(load32(m.memory[int64(uint32(v2))+56:]))
		v13 = t32
		if uint32(v13) >= uint32(i32(0x7fffffff)) {
			m.fn1518(i32(1085076))
			panic("unreachable")
		}
		t33 := int32(load32(m.memory[int64(uint32(v8))+60:]))
		v15 = t33
		store32(m.memory[int64(uint32(v2))+56:], uint32(v13+i32(1)))
		{
			t34 := int32(load32(m.memory[int64(uint32(v2))+76:]))
			if t34 == 0 {
				goto l5
			}
			t35 := int64(load64(m.memory[int64(uint32(v2))+80:]))
			t36 := int64(load64(m.memory[uint32(v2+i32(88)):]))
			t37 := m.fn29(t35, t36, v14, v15)
			v18 = t37
			t38 := int32(load32(m.memory[int64(uint32(v2))+68:]))
			v19 = t38
			v17 = v19 & int32(v18)
			v20 = int64(uint64(v18)>>25) & i64(127) * i64(72340172838076673)
			t39 := int32(load32(m.memory[int64(uint32(v2))+64:]))
			v21 = t39
			v22 = i32(0)
		l9:
			{
				t40 := int64(load64(m.memory[uint32(v21+v17):]))
				v23 = t40
				v18 = v23 ^ v20
				v18 = (v18 ^ i64(-1)) & (v18 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
			l8:
				{
					if v18 == 0 {
						if !(v23&(v23<<1)&i64(-0x7f7f7f7f7f7f7f80) == 0) {
							goto l5
						}
						t46 := v17
						v22 = v22 + i32(8)
						v17 = (t46 + v22) & v19
						goto l9
					}
					t41 := v14
					t42 := v15
					v24 = v21 + (i32(0)-(int32(uint32(int64(bits.TrailingZeros64(uint64(v18))))>>3)+v17)&v19)*i32(24)
					t43 := int32(load32(m.memory[uint32(v24+i32(-20)):]))
					t44 := int32(load32(m.memory[uint32(v24+i32(-16)):]))
					t45 := m.fn15(t41, t42, t43, t44)
					if t45 != 0 {
						goto l7
					}
					v18 = (v18 + i64(-1)) & v18
					goto l8
				}
			l7:
			}
			t47 := int64(load64(m.memory[uint32(v24+i32(-8)):]))
			v16 = t47
		}
	l5:
		store32(m.memory[int64(uint32(v2))+56:], uint32(v13))
		goto l2
	}
l11:
	store32(m.memory[int64(uint32(v2))+16:], uint32(v13))
l2:
	store32(m.memory[int64(uint32(v8))+100:], uint32(i32(0)))
	store64(m.memory[int64(uint32(v8))+92:], uint64(i64(0x800000000)))
	store32(m.memory[int64(uint32(v8))+120:], uint32(i32(0)))
	store64(m.memory[int64(uint32(v8))+112:], uint64(i64(0x400000000)))
	store64(m.memory[int64(uint32(v8))+104:], uint64(v16))
	m.memory[int64(uint32(v8))+124] = byte(v11)
	t65 := int32(load32(m.memory[int64(uint32(v1))+32:]))
	v13 = t65
	t66 := int32(load32(m.memory[int64(uint32(v1))+28:]))
	t67 := v8
	v1 = t66
	store32(m.memory[int64(uint32(t67))+128:], uint32(v1))
	store32(m.memory[int64(uint32(v8))+132:], uint32(v1+v13*i32(44)))
	v11 = v7 + i32(1)
	v14 = v3 + i32(1)
	v15 = v7 << 3
	v1 = v8 + i32(192) + i32(4)
	v19 = v8 + i32(104) + i32(8)
	p68 := i32(9)
	if uint32(v3) < uint32(i32(9)) {
		p68 = v3
	}
	v25 = p68
	v26 = v25 * i32(40)
	v22 = i32(1)
	{
		{
		l19:
			{
				{
					{
						t69 := m.fn904(v8 + i32(128))
						v3 = t69
						if v3 == 0 {
							m.fn1519(v8+i32(104), v8+i32(92), v16)
							if v12 != 0 {
								goto l21
							}
							t72 := int32(load32(m.memory[int64(uint32(v8))+100:]))
							if t72 == 0 {
								goto l21
							}
							t73 := int32(load32(m.memory[int64(uint32(v2))+16:]))
							if t73 != 0 {
								m.fn1326(i32(1085128))
								panic("unreachable")
							}
							store32(m.memory[int64(uint32(v2))+16:], uint32(i32(-1)))
							t74 := int64(load64(m.memory[int64(uint32(v8))+84:]))
							store64(m.memory[int64(uint32(v8))+200:], uint64(t74))
							t75 := int64(load64(m.memory[int64(uint32(v8))+76:]))
							store64(m.memory[int64(uint32(v8))+192:], uint64(t75))
							v3 = v2 + i32(24)
							t76 := int64(load64(m.memory[int64(uint32(v2))+40:]))
							t77 := int64(load64(m.memory[int64(uint32(v2))+48:]))
							t78 := m.fn696(t76, t77, v8+i32(76))
							v18 = t78
							store32(m.memory[int64(uint32(v8))+236:], uint32(v8+i32(192)))
							{
								t79 := int32(load32(m.memory[int64(uint32(v2))+32:]))
								if t79 != 0 {
									goto l23
								}
								_ = m.fn694(v3, v2+i32(40))
							}
						l23:
							store32(m.memory[int64(uint32(v8))+164:], uint32(v3))
							store32(m.memory[int64(uint32(v8))+160:], uint32(v8+i32(236)))
							t81 := int32(load32(m.memory[int64(uint32(v2))+24:]))
							t82 := int32(load32(m.memory[int64(uint32(v2))+28:]))
							m.fn69(v8+i32(24), t81, t82, v18, v8+i32(160), i32(190))
							t83 := int32(load32(m.memory[int64(uint32(v2))+24:]))
							v3 = t83
							t84 := int32(load32(m.memory[int64(uint32(v8))+28:]))
							v1 = t84
							{
								{
									t85 := int32(load32(m.memory[int64(uint32(v8))+24:]))
									if t85 != i32(1) {
										goto l24
									}
									v13 = v3 + v1
									t86 := int32(m.memory[uint32(v13)])
									v4 = t86
									t87 := int64(load64(m.memory[int64(uint32(v8))+200:]))
									v20 = t87
									t88 := int64(load64(m.memory[int64(uint32(v8))+192:]))
									v23 = t88
									t89 := v13
									v5 = int32(uint32(int32(v18)) >> 25)
									m.memory[uint32(t89)] = byte(v5)
									t90 := int32(load32(m.memory[int64(uint32(v2))+28:]))
									m.memory[uint32(v3+t90&(v1+i32(-8))+i32(8))] = byte(v5)
									t91 := int32(load32(m.memory[int64(uint32(v2))+36:]))
									store32(m.memory[int64(uint32(v2))+36:], uint32(t91+i32(1)))
									t92 := int32(load32(m.memory[int64(uint32(v2))+32:]))
									store32(m.memory[int64(uint32(v2))+32:], uint32(t92-v4&i32(1)))
									v3 = v3 + (i32(0)-v1)*i32(24)
									v1 = v3 + i32(-24)
									store64(m.memory[uint32(v1):], uint64(v23))
									store64(m.memory[int64(uint32(v1))+8:], uint64(v20))
									store64(m.memory[uint32(v3+i32(-8)):], uint64(v16))
									goto l25
								}
							l24:
								store64(m.memory[uint32(v3+(i32(0)-v1)*i32(24)+i32(-8)):], uint64(v16))
								t93 := int32(load32(m.memory[int64(uint32(v8))+192:]))
								t94 := int32(load32(m.memory[int64(uint32(v8))+196:]))
								m.fn16(t93, t94)
							}
						l25:
							t95 := int32(load32(m.memory[int64(uint32(v2))+16:]))
							store32(m.memory[int64(uint32(v2))+16:], uint32(t95+i32(1)))
							m.fn845(v8+i32(16), v9, v10, i32(1282584), i32(36), i32(1073226), i32(2))
							{
								t96 := int32(load32(m.memory[int64(uint32(v8))+16:]))
								v3 = t96
								if v3 == 0 {
									goto l26
								}
								t97 := int32(load32(m.memory[int64(uint32(v2))+56:]))
								if t97 != 0 {
									m.fn1326(i32(1085144))
									panic("unreachable")
								}
								t98 := int32(load32(m.memory[int64(uint32(v8))+20:]))
								v1 = t98
								store32(m.memory[int64(uint32(v2))+56:], uint32(i32(-1)))
								m.fn51(v8+i32(160), v3, v1)
								v3 = v2 + i32(64)
								t99 := int64(load64(m.memory[int64(uint32(v2))+80:]))
								t100 := int64(load64(m.memory[int64(uint32(v2))+88:]))
								t101 := int32(load32(m.memory[int64(uint32(v8))+164:]))
								t102 := int32(load32(m.memory[int64(uint32(v8))+168:]))
								t103 := m.fn540(t99, t100, t101, t102)
								v18 = t103
								store32(m.memory[int64(uint32(v8))+236:], uint32(v8+i32(160)))
								{
									t104 := int32(load32(m.memory[int64(uint32(v2))+72:]))
									if t104 != 0 {
										goto l28
									}
									_ = m.fn685(v3, v2+i32(80))
								}
							l28:
								store32(m.memory[int64(uint32(v8))+196:], uint32(v3))
								store32(m.memory[int64(uint32(v8))+192:], uint32(v8+i32(236)))
								t106 := int32(load32(m.memory[int64(uint32(v2))+64:]))
								t107 := int32(load32(m.memory[int64(uint32(v2))+68:]))
								m.fn69(v8+i32(8), t106, t107, v18, v8+i32(192), i32(191))
								t108 := int32(load32(m.memory[int64(uint32(v2))+64:]))
								v3 = t108
								t109 := int32(load32(m.memory[int64(uint32(v8))+12:]))
								v1 = t109
								{
									{
										t110 := int32(load32(m.memory[int64(uint32(v8))+8:]))
										if t110 != i32(1) {
											goto l29
										}
										v13 = v3 + v1
										t111 := int32(m.memory[uint32(v13)])
										v4 = t111
										t112 := int32(load32(m.memory[int64(uint32(v8))+168:]))
										v5 = t112
										t113 := int64(load64(m.memory[int64(uint32(v8))+160:]))
										v20 = t113
										t114 := v13
										v11 = int32(uint32(int32(v18)) >> 25)
										m.memory[uint32(t114)] = byte(v11)
										t115 := int32(load32(m.memory[int64(uint32(v2))+68:]))
										m.memory[uint32(v3+t115&(v1+i32(-8))+i32(8))] = byte(v11)
										t116 := int32(load32(m.memory[int64(uint32(v2))+76:]))
										store32(m.memory[int64(uint32(v2))+76:], uint32(t116+i32(1)))
										t117 := int32(load32(m.memory[int64(uint32(v2))+72:]))
										store32(m.memory[int64(uint32(v2))+72:], uint32(t117-v4&i32(1)))
										v3 = v3 + (i32(0)-v1)*i32(24)
										v1 = v3 + i32(-24)
										store64(m.memory[uint32(v1):], uint64(v20))
										store32(m.memory[int64(uint32(v8))+200:], uint32(v5))
										t118 := int64(load64(m.memory[int64(uint32(v8))+200:]))
										store64(m.memory[int64(uint32(v1))+8:], uint64(t118))
										store64(m.memory[uint32(v3+i32(-8)):], uint64(v16))
										goto l30
									}
								l29:
									store64(m.memory[uint32(v3+(i32(0)-v1)*i32(24)+i32(-8)):], uint64(v16))
									t119 := int32(load32(m.memory[int64(uint32(v8))+160:]))
									t120 := int32(load32(m.memory[int64(uint32(v8))+164:]))
									m.fn16(t119, t120)
								}
							l30:
								t121 := int32(load32(m.memory[int64(uint32(v2))+56:]))
								store32(m.memory[int64(uint32(v2))+56:], uint32(t121+i32(1)))
							}
						l26:
							t122 := int32(load32(m.memory[int64(uint32(v8))+100:]))
							store32(m.memory[int64(uint32(v0))+12:], uint32(t122))
							t123 := int64(load64(m.memory[int64(uint32(v8))+92:]))
							store64(m.memory[int64(uint32(v0))+4:], uint64(t123))
							store32(m.memory[uint32(v0):], uint32(i32(-1)))
							m.fn971(v19)
							goto l31
						}
						t70 := m.fn847(v3, i32(1074680), i32(46), i32(1085160), i32(11))
						v24 = t70
						if v24 != 0 {
							goto l18
						}
						t71 := m.fn847(v3, i32(1074680), i32(46), i32(1085171), i32(9))
						if t71 == 0 {
							goto l19
						}
						if v12 == 0 {
							t124 := int32(load32(m.memory[uint32(v3+i32(16)):]))
							t125 := int32(load32(m.memory[uint32(v3+i32(20)):]))
							m.fn1046(v8+i32(40), t124, t125, i32(1074680), i32(46), i32(1085180), i32(11))
							t126 := int32(load32(m.memory[int64(uint32(v8))+40:]))
							t127 := int32(load32(m.memory[int64(uint32(v8))+44:]))
							m.fn1514(v8+i32(192), t126, t127)
							t128 := int64(load64(m.memory[int64(uint32(v8))+192:]))
							if t128 != i64(1) {
								goto l18
							}
							t129 := int64(load64(m.memory[int64(uint32(v8))+200:]))
							v18 = t129
							if v22&i32(1) != 0 {
								store64(m.memory[int64(uint32(v8))+104:], uint64(v18))
								goto l18
							}
							m.fn1519(v8+i32(104), v8+i32(92), v18)
							goto l18
						}
						goto l18
					}
				l21:
					t130 := int32(load32(m.memory[int64(uint32(v8))+100:]))
					store32(m.memory[int64(uint32(v0))+12:], uint32(t130))
					t131 := int64(load64(m.memory[int64(uint32(v8))+92:]))
					store64(m.memory[int64(uint32(v0))+4:], uint64(t131))
					store32(m.memory[uint32(v0):], uint32(i32(-1)))
					m.fn971(v19)
					t132 := int32(load32(m.memory[int64(uint32(v8))+76:]))
					t133 := int32(load32(m.memory[int64(uint32(v8))+80:]))
					m.fn16(t132, t133)
					goto l31
				}
			l18:
				t134 := int64(load32(m.memory[int64(uint32(v8))+120:]))
				v20 = t134
				t135 := int64(load64(m.memory[int64(uint32(v8))+104:]))
				v18 = t135
				m.fn59(v8+i32(32), v7, i32(8), i32(8))
				store32(m.memory[int64(uint32(v8))+144:], uint32(i32(0)))
				t136 := int32(load32(m.memory[int64(uint32(v8))+36:]))
				t137 := v8
				v13 = t136
				store32(m.memory[int64(uint32(t137))+140:], uint32(v13))
				t138 := int32(load32(m.memory[int64(uint32(v8))+32:]))
				t139 := v8
				v17 = t138
				store32(m.memory[int64(uint32(t139))+136:], uint32(v17))
				v20 = v18 + v20
				var p140 int32
				if uint64(v20) < uint64(v18) {
					p140 = 1
				}
				v21 = p140
				if v7 == 0 {
					goto l33
				}
				if v15 == 0 {
					goto l34
				}
				memory_copy(m.memory, uint32(v13), uint32(v6), uint32(v15))
			l34:
				store32(m.memory[int64(uint32(v8))+144:], uint32(v7))
			l33:
				p141 := v20
				if v21 != 0 {
					p141 = i64(-1)
				}
				v18 = p141
				{
					if v7 != v17 {
						goto l35
					}
					m.fn1154(v8 + i32(136))
					t142 := int32(load32(m.memory[int64(uint32(v8))+140:]))
					v13 = t142
				}
			l35:
				store64(m.memory[uint32(v13+v15):], uint64(v18))
				store32(m.memory[int64(uint32(v8))+156:], uint32(i32(0)))
				store64(m.memory[int64(uint32(v8))+148:], uint64(i64(0x800000000)))
				store32(m.memory[int64(uint32(v8))+160:], uint32(i32(0)))
				t143 := int32(load32(m.memory[int64(uint32(v3))+32:]))
				v17 = t143
				t144 := int32(load32(m.memory[int64(uint32(v3))+28:]))
				t145 := v8
				v3 = t144
				store32(m.memory[int64(uint32(t145))+236:], uint32(v3))
				store32(m.memory[int64(uint32(v8))+240:], uint32(v3+v17*i32(44)))
			l38:
				{
					{
						t146 := m.fn904(v8 + i32(236))
						v3 = t146
						if v3 == 0 {
							m.fn1333(v8+i32(160), v8+i32(148))
							{
								if v24 == 0 {
									v3 = i32(-1)
									{
										if v4 == 0 {
											goto l41
										}
										t153 := int32(load32(m.memory[int64(uint32(v2))+200:]))
										t154 := m.fn1517(t153+i32(432), v4, v5)
										v27 = t154
										if v27 == 0 {
											goto l41
										}
										v17 = v27 + v26
										t155 := int32(m.memory[int64(uint32(v17))+36])
										v28 = t155
										if v28 == 0 {
											goto l41
										}
										m.fn1515(v8+i32(192), v17, v25)
										v3 = i32(-1)
										{
											t156 := int32(load32(m.memory[int64(uint32(v8))+200:]))
											v17 = t156
											if v17 == 0 {
												goto l42
											}
											v3 = v13 + v11<<3 + i32(-8)
											p157 := i32(1079200)
											if v3 != 0 {
												p157 = v3
											}
											t158 := int64(load64(m.memory[uint32(p157):]))
											v18 = t158
											store32(m.memory[int64(uint32(v8))+232:], uint32(i32(0)))
											store64(m.memory[int64(uint32(v8))+224:], uint64(i64(0x100000000)))
											v17 = v17 * i32(12)
											t159 := int32(m.memory[int64(uint32(v8))+204])
											v29 = t159
											t160 := int32(load32(m.memory[int64(uint32(v8))+196:]))
											v3 = t160
										l48:
											{
												{
													if v17 == 0 {
														m.fn800(v8+i32(236), v28, v18)
														t173 := int32(load32(m.memory[int64(uint32(v8))+228:]))
														v21 = t173
														t174 := int32(load32(m.memory[int64(uint32(v8))+232:]))
														t175 := int32(load32(m.memory[int64(uint32(v8))+240:]))
														t176 := v21
														v3 = t175
														t177 := int32(load32(m.memory[int64(uint32(v8))+244:]))
														t178 := m.fn191(t176, t174, v3, t177)
														v17 = t178
														t179 := int32(load32(m.memory[int64(uint32(v8))+236:]))
														m.fn16(t179, v3)
														t180 := int32(load32(m.memory[int64(uint32(v8))+224:]))
														v3 = t180
														if v17 != 0 {
															goto l47
														}
														t181 := int64(load64(m.memory[int64(uint32(v8))+228:]))
														v23 = t181
														goto l42
													}
													t161 := int32(load32(m.memory[uint32(v3):]))
													if t161 != i32(-1) {
														goto l44
													}
													t162 := int32(m.memory[uint32(v3+i32(4))])
													v21 = t162
													p163 := i32(9)
													if uint32(v21) < uint32(i32(9)) {
														p163 = v21
													}
													v24 = p163
													v22 = i32(1)
													{
														if v29&i32(1) != 0 {
															goto l45
														}
														t164 := int32(m.memory[int64(uint32(v27+v24*i32(40)))+36])
														v22 = t164
													}
												l45:
													t166 := v8 + i32(236)
													t167 := v22
													p165 := v13 + v21<<3
													if uint32(v7) < uint32(v21) {
														p165 = v27 + v24*i32(40)
													}
													t168 := int64(load64(m.memory[uint32(p165):]))
													m.fn804(t166, t167, t168)
													t169 := int32(load32(m.memory[int64(uint32(v8))+240:]))
													t170 := v8 + i32(224)
													v21 = t169
													t171 := int32(load32(m.memory[int64(uint32(v8))+244:]))
													m.fn75(t170, v21, t171)
													t172 := int32(load32(m.memory[int64(uint32(v8))+236:]))
													m.fn16(t172, v21)
													goto l46
												}
											l44:
												t182 := int32(load32(m.memory[uint32(v3+i32(4)):]))
												t183 := int32(load32(m.memory[uint32(v3+i32(8)):]))
												m.fn75(v8+i32(224), t182, t183)
											}
										l46:
											v3 = v3 + i32(12)
											v17 = v17 + i32(-12)
											goto l48
										l47:
											m.fn16(v3, v21)
											v3 = i32(-1)
										}
									l42:
										m.fn763(v8 + i32(192))
									}
								l41:
									t184 := int32(load32(m.memory[int64(uint32(v8))+156:]))
									store32(m.memory[int64(uint32(v8))+200:], uint32(t184))
									t185 := int64(load64(m.memory[int64(uint32(v8))+148:]))
									store64(m.memory[int64(uint32(v8))+192:], uint64(t185))
									m.memory[int64(uint32(v8))+216] = byte(i32(2))
									store64(m.memory[int64(uint32(v8))+208:], uint64(v23))
									store32(m.memory[int64(uint32(v8))+204:], uint32(v3))
									m.fn1412(v19, v8+i32(192))
									t186 := int64(load64(m.memory[int64(uint32(v8))+104:]))
									v18 = t186
									t187 := int64(load32(m.memory[int64(uint32(v8))+120:]))
									v20 = t187
									m.fn1332(v8 + i32(160))
									t188 := int32(load32(m.memory[int64(uint32(v8))+136:]))
									m.fn1415(t188, v13)
									v20 = v18 + v20
									p189 := v20
									if uint64(v20) < uint64(v18) {
										p189 = i64(-1)
									}
									v16 = p189
									v22 = i32(0)
									goto l19
								}
								m.fn1519(v8+i32(104), v8+i32(92), v16)
								m.fn1271(v8+i32(92), v8+i32(148))
								m.fn1332(v8 + i32(160))
								t152 := int32(load32(m.memory[int64(uint32(v8))+136:]))
								m.fn1415(t152, v13)
								goto l19
							}
						}
						t147 := m.fn847(v3, i32(1074680), i32(46), i32(1081789), i32(4))
						if t147 != 0 {
							goto l37
						}
						m.fn1331(v8+i32(192), v3, v2, v8+i32(148), v8+i32(160))
						t148 := int32(load32(m.memory[int64(uint32(v8))+192:]))
						v3 = t148
						if v3 == i32(-1) {
							goto l38
						}
						t149 := int32(load32(m.memory[int64(uint32(v8))+212:]))
						store32(m.memory[int64(uint32(v0))+20:], uint32(t149))
						t150 := int64(load64(m.memory[int64(uint32(v8))+204:]))
						store64(m.memory[int64(uint32(v0))+12:], uint64(t150))
						t151 := int64(load64(m.memory[int64(uint32(v8))+196:]))
						store64(m.memory[int64(uint32(v0))+4:], uint64(t151))
						store32(m.memory[uint32(v0):], uint32(v3))
						goto l39
					}
				l37:
					m.fn1333(v8+i32(160), v8+i32(148))
					m.fn1511(v8+i32(192), v3, v2, v14, v4, v5, v13, v11)
					t190 := int64(load64(m.memory[uint32(v1):]))
					store64(m.memory[int64(uint32(v8))+176:], uint64(t190))
					t191 := int32(load32(m.memory[int64(uint32(v1))+8:]))
					store32(m.memory[int64(uint32(v8))+184:], uint32(t191))
					{
						t192 := int32(load32(m.memory[int64(uint32(v8))+192:]))
						v3 = t192
						if v3 != i32(-1) {
							goto l49
						}
						m.fn1271(v8+i32(148), v8+i32(176))
						goto l38
					}
				l49:
				}
			}
			t193 := int64(load64(m.memory[int64(uint32(v8))+208:]))
			v18 = t193
			t194 := int32(load32(m.memory[int64(uint32(v8))+184:]))
			store32(m.memory[int64(uint32(v0))+12:], uint32(t194))
			t195 := int64(load64(m.memory[int64(uint32(v8))+176:]))
			store64(m.memory[int64(uint32(v0))+4:], uint64(t195))
			store64(m.memory[int64(uint32(v0))+16:], uint64(v18))
			store32(m.memory[uint32(v0):], uint32(v3))
		}
	l39:
		m.fn1332(v8 + i32(160))
		m.fn969(v8 + i32(148))
		t196 := int32(load32(m.memory[int64(uint32(v8))+136:]))
		m.fn1415(t196, v13)
		m.fn971(v19)
		m.fn969(v8 + i32(92))
		t197 := int32(load32(m.memory[int64(uint32(v8))+76:]))
		t198 := int32(load32(m.memory[int64(uint32(v8))+80:]))
		m.fn16(t197, t198)
	}
l31:
	m.fn753(v8 + i32(248))
	m.g0 = v8 + i32(304)
}
func (m *Module) fn1512(v0, v1 int32, v2 int64, v3 int32) int32 {
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
				v11 = t7 + (i32(0)-v10)*i32(28)
				t8 := int32(load32(m.memory[uint32(v11+i32(-24)):]))
				t9 := int32(load32(m.memory[uint32(v11+i32(-20)):]))
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
			p11 := v0 + (i32(0)-v10)*i32(28)
			if v3 != 0 {
				p11 = i32(0)
			}
			return p11
		}
	}
}
func (m *Module) fn1513(v0, v1, v2, v3 int32) {
	var v4 int32
	t0 := m.g0
	v4 = t0 - i32(16)
	m.g0 = v4
	m.fn1046(v4+i32(8), v1, v2, i32(1074680), i32(46), i32(1085052), i32(10))
	t1 := int32(load32(m.memory[int64(uint32(v4))+8:]))
	t2 := v0
	t3 := v3
	v2 = t1
	p4 := i32(1)
	if v2 != 0 {
		p4 = v2
	}
	t5 := int32(load32(m.memory[int64(uint32(v4))+12:]))
	p6 := i32(0)
	if v2 != 0 {
		p6 = t5
	}
	m.fn1520(t2, t3, i32(1077240), i32(9), p4, p6)
	m.g0 = v4 + i32(16)
}
func (m *Module) fn1514(v0, v1, v2 int32) {
	var v3 int32
	var v4 int64
	t0 := m.g0
	v3 = t0 - i32(16)
	m.g0 = v3
	v4 = i64(0)
	{
		if v1 == 0 {
			goto l0
		}
		m.fn1190(v3, v1, v2)
		t1 := int32(m.memory[uint32(v3)])
		if t1 == i32(1) {
			goto l0
		}
		t2 := int64(load64(m.memory[int64(uint32(v3))+8:]))
		t3 := v0
		v4 = t2
		p4 := i64(0xffffffff)
		if uint64(v4) < uint64(i64(0xffffffff)) {
			p4 = v4
		}
		store64(m.memory[int64(uint32(t3))+8:], uint64(p4))
		v4 = i64(1)
	}
l0:
	store64(m.memory[uint32(v0):], uint64(v4))
	m.g0 = v3 + i32(16)
}
func (m *Module) fn1515(v0, v1, v2 int32) {
	var v3, v4, v5 int32
	t0 := m.g0
	v3 = t0 - i32(48)
	m.g0 = v3
	store32(m.memory[int64(uint32(v3))+16:], uint32(i32(0)))
	store64(m.memory[int64(uint32(v3))+8:], uint64(i64(0x400000000)))
	{
		t1 := int32(load32(m.memory[int64(uint32(v1))+16:]))
		v4 = t1
		if v4 == 0 {
			goto l0
		}
		t2 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		m.fn31(v3+i32(20), t2, v4)
		m.fn1321(v3+i32(8), v3+i32(20))
	}
l0:
	t3 := int32(load32(m.memory[int64(uint32(v1))+32:]))
	v4 = t3
	m.memory[int64(uint32(v3))+32] = byte(i32(0))
	store32(m.memory[int64(uint32(v3))+20:], uint32(i32(0)))
	store32(m.memory[int64(uint32(v3))+28:], uint32(v2))
	t4 := v3
	v2 = v2 + i32(1)
	t6 := v2
	p5 := v2
	if uint32(v4) < uint32(v2) {
		p5 = v4
	}
	p7 := i32(1)
	if v4 != 0 {
		p7 = p5
	}
	store32(m.memory[int64(uint32(t4))+24:], uint32(t6-p7))
	v5 = v3 + i32(24)
l5:
	m.fn254(v3, v5)
	{
		t8 := int32(load32(m.memory[uint32(v3):]))
		if t8 != i32(1) {
			{
				t12 := int32(load32(m.memory[int64(uint32(v1))+20:]))
				if t12 == i32(-1) {
					goto l4
				}
				t13 := int32(load32(m.memory[int64(uint32(v1))+28:]))
				v4 = t13
				if v4 == 0 {
					goto l4
				}
				t14 := int32(load32(m.memory[int64(uint32(v1))+24:]))
				m.fn31(v3+i32(20), t14, v4)
				m.fn1321(v3+i32(8), v3+i32(20))
			}
		l4:
			t15 := int32(load32(m.memory[int64(uint32(v3))+16:]))
			store32(m.memory[int64(uint32(v0))+8:], uint32(t15))
			t16 := int64(load64(m.memory[int64(uint32(v3))+8:]))
			store64(m.memory[uint32(v0):], uint64(t16))
			m.memory[int64(uint32(v0))+12] = byte(i32(0))
			m.g0 = v3 + i32(48)
			return
		}
		t9 := int32(load32(m.memory[int64(uint32(v3))+4:]))
		v4 = t9
		t10 := int32(load32(m.memory[int64(uint32(v3))+20:]))
		t11 := v3
		v2 = t10
		store32(m.memory[int64(uint32(t11))+20:], uint32(v2+i32(1)))
		if v2 != 0 {
			goto l2
		}
		goto l3
	}
l2:
	m.fn51(v3+i32(36), i32(1109519), i32(1))
	m.fn1321(v3+i32(8), v3+i32(36))
l3:
	store32(m.memory[int64(uint32(v3))+36:], uint32(i32(-1)))
	m.memory[int64(uint32(v3))+40] = byte(v4)
	m.fn1321(v3+i32(8), v3+i32(36))
	goto l5
}
func (m *Module) fn1516(v0, v1, v2, v3, v4, v5 int32) {
	var v6, v7, v8, v9, v10, v11, v12, v13, v14, v15, v16, v17, v18 int32
	t0 := m.g0
	v6 = t0 - i32(144)
	m.g0 = v6
	v7 = v2 + i32(4)
	t1 := m.fn1188(v3)
	v8 = t1
	t2 := int32(load32(m.memory[int64(uint32(v1))+32:]))
	v9 = t2 * i32(44)
	v10 = v6 + i32(104) + i32(4)
	v11 = v6 + i32(104) + i32(16)
	v12 = v6 + i32(80) + i32(4)
	t3 := int32(load32(m.memory[int64(uint32(v1))+28:]))
	v13 = t3
	v14 = i32(0)
l31:
	{
		{
			{
				if v9 == v14 {
					store32(m.memory[uint32(v0):], uint32(i32(-1)))
					goto l4
				}
				v1 = v13 + v14
				t4 := int32(load32(m.memory[uint32(v1):]))
				if t4 != i32(-1) {
					t14 := int32(load32(m.memory[uint32(v1+i32(36)):]))
					v15 = t14
					if v15 == 0 {
						goto l5
					}
					t15 := int32(load32(m.memory[uint32(v1+i32(40)):]))
					t16 := m.fn15(v15+i32(8), t15, i32(1074680), i32(46))
					if t16 == 0 {
						goto l5
					}
					{
						t17 := int32(load32(m.memory[uint32(v1+i32(4)):]))
						v15 = t17
						t18 := int32(load32(m.memory[uint32(v1+i32(8)):]))
						t19 := v15
						v16 = t18
						t20 := m.fn15(t19, v16, i32(1085219), i32(4))
						if t20 != 0 {
							t89 := int32(load32(m.memory[uint32(v1+i32(16)):]))
							t90 := int32(load32(m.memory[uint32(v1+i32(20)):]))
							m.fn1046(v6, t89, t90, i32(1074680), i32(46), i32(1085052), i32(10))
							v15 = v3
							t91 := int32(load32(m.memory[uint32(v6):]))
							v16 = t91
							if v16 == 0 {
								goto l25
							}
							t92 := int32(load32(m.memory[int64(uint32(v2))+200:]))
							t93 := int32(load32(m.memory[int64(uint32(v6))+4:]))
							m.fn1520(v6+i32(104), t92, i32(1081916), i32(4), v16, t93)
							t94 := int32(load32(m.memory[int64(uint32(v6))+108:]))
							v15 = t94
							{
								t95 := int32(load32(m.memory[int64(uint32(v6))+104:]))
								v16 = t95
								if v16 == i32(-1) {
									t98 := fn1373(v3, v15)
									v15 = t98
									goto l25
								}
								t96 := int64(load64(m.memory[int64(uint32(v6))+120:]))
								store64(m.memory[int64(uint32(v0))+16:], uint64(t96))
								t97 := int64(load64(m.memory[int64(uint32(v6))+112:]))
								store64(m.memory[int64(uint32(v0))+8:], uint64(t97))
								store32(m.memory[int64(uint32(v0))+4:], uint32(v15))
								store32(m.memory[uint32(v0):], uint32(v16))
								goto l4
							}
						}
						{
							t21 := m.fn15(v15, v16, i32(1077050), i32(1))
							if t21 != 0 {
								t68 := int32(load32(m.memory[uint32(v1+i32(16)):]))
								t69 := int32(load32(m.memory[uint32(v1+i32(20)):]))
								m.fn1046(v6+i32(24), t68, t69, i32(1085191), i32(28), i32(1073490), i32(4))
								t70 := int32(load32(m.memory[int64(uint32(v6))+28:]))
								v16 = t70
								t71 := int32(load32(m.memory[int64(uint32(v6))+24:]))
								v15 = t71
								store64(m.memory[int64(uint32(v6))+68:], uint64(i64(0x400000000)))
								store32(m.memory[int64(uint32(v6))+76:], uint32(i32(0)))
								m.fn1516(v6+i32(104), v1, v2, v3, v6+i32(68), v5)
								{
									t72 := int32(load32(m.memory[int64(uint32(v6))+104:]))
									v1 = t72
									if v1 == i32(-1) {
										p76 := i32(0)
										if v15 != 0 {
											p76 = v16
										}
										v1 = p76
										if v1 == 0 {
											goto l22
										}
										t78 := v6 + i32(16)
										p77 := i32(1)
										if v15 != 0 {
											p77 = v15
										}
										v16 = p77
										m.fn13(t78, v16, v1, i32(35))
										t79 := int32(load32(m.memory[int64(uint32(v6))+16:]))
										v15 = t79
										if v15 == 0 {
											t101 := m.fn1455(v16, v1)
											v15 = t101
											m.fn51(v12, v16, v1)
											v1 = v15 ^ i32(1)
											goto l24
										}
										t80 := int32(load32(m.memory[int64(uint32(v6))+20:]))
										t81 := v6 + i32(104)
										t82 := v15
										v16 = t80
										m.fn513(t81, t82, v16, i32(124))
										m.fn515(v6+i32(8), v6+i32(104))
										t83 := int32(load32(m.memory[int64(uint32(v6))+12:]))
										t84 := int32(load32(m.memory[int64(uint32(v6))+8:]))
										t85 := v16
										v1 = t84
										p86 := t85
										if v1 != 0 {
											p86 = t83
										}
										v16 = p86
										if v16 == 0 {
											goto l22
										}
										t88 := v12
										p87 := v15
										if v1 != 0 {
											p87 = v1
										}
										m.fn776(t88, p87, v16)
										v1 = i32(2)
										goto l24
									}
									t73 := int32(load32(m.memory[int64(uint32(v6))+124:]))
									store32(m.memory[int64(uint32(v0))+20:], uint32(t73))
									t74 := int64(load64(m.memory[int64(uint32(v6))+116:]))
									store64(m.memory[int64(uint32(v0))+12:], uint64(t74))
									t75 := int64(load64(m.memory[int64(uint32(v6))+108:]))
									store64(m.memory[int64(uint32(v0))+4:], uint64(t75))
									store32(m.memory[uint32(v0):], uint32(v1))
									m.fn894(v6 + i32(68))
									goto l4
								}
							}
							{
								t22 := m.fn15(v15, v16, i32(1079224), i32(1))
								if t22 != 0 {
									v15 = i32(1)
									t31 := int32(load32(m.memory[uint32(v1+i32(16)):]))
									t32 := int32(load32(m.memory[uint32(v1+i32(20)):]))
									m.fn1046(v6+i32(32), t31, t32, i32(1074680), i32(46), i32(1077932), i32(1))
									{
										t33 := int32(load32(m.memory[int64(uint32(v6))+32:]))
										v1 = t33
										if v1 == 0 {
											goto l13
										}
										t34 := int32(load32(m.memory[int64(uint32(v6))+36:]))
										m.fn197(v6+i32(104), v1, t34)
										t35 := int32(load32(m.memory[int64(uint32(v6))+108:]))
										v1 = t35
										p36 := i32(20)
										if uint32(v1) < uint32(i32(20)) {
											p36 = v1
										}
										t37 := int32(m.memory[int64(uint32(v6))+104])
										p38 := p36
										if t37 != 0 {
											p38 = i32(1)
										}
										v15 = p38
									}
								l13:
									m.fn808(v10, i32(1097368), v15)
									store32(m.memory[int64(uint32(v6))+104:], uint32(i32(3)))
									store32(m.memory[int64(uint32(v6))+120:], uint32(i32(0)))
									m.fn1340(v4, v6+i32(104))
									goto l3
								}
								t23 := m.fn15(v15, v16, i32(1077507), i32(3))
								if t23 != 0 {
									m.fn51(v10, i32(1097368), i32(1))
									store32(m.memory[int64(uint32(v6))+104:], uint32(i32(3)))
									store32(m.memory[int64(uint32(v6))+120:], uint32(i32(0)))
									m.fn1340(v4, v6+i32(104))
									goto l3
								}
								t24 := m.fn15(v15, v16, i32(1085223), i32(10))
								if t24 != 0 {
									store32(m.memory[int64(uint32(v6))+104:], uint32(i32(8)))
									m.fn1340(v4, v6+i32(104))
									goto l3
								}
								{
									t25 := m.fn15(v15, v16, i32(1085233), i32(8))
									if t25 != 0 {
										goto l11
									}
									t26 := m.fn15(v15, v16, i32(1085241), i32(14))
									if t26 == 0 {
										t39 := m.fn15(v15, v16, i32(1085255), i32(4))
										if t39 != 0 {
											t43 := int32(load32(m.memory[uint32(v2):]))
											v15 = t43
											if uint32(v15) >= uint32(i32(0x7fffffff)) {
												m.fn1518(i32(1085292))
												panic("unreachable")
											}
											store32(m.memory[uint32(v2):], uint32(v15))
											t44 := int32(load32(m.memory[int64(uint32(v2))+12:]))
											store32(m.memory[int64(uint32(v6))+100:], uint32(t44))
											t45 := v6 + i32(56)
											v15 = v1 + i32(16)
											t46 := int32(load32(m.memory[uint32(v15):]))
											v16 = v1 + i32(20)
											t47 := int32(load32(m.memory[uint32(v16):]))
											m.fn1046(t45, t46, t47, i32(1074680), i32(46), i32(1073226), i32(2))
											{
												t48 := int32(load32(m.memory[int64(uint32(v6))+56:]))
												v17 = t48
												if v17 == 0 {
													goto l16
												}
												t49 := int32(load32(m.memory[int64(uint32(v6))+60:]))
												m.fn51(v6+i32(104), v17, t49)
												t50 := int32(load32(m.memory[int64(uint32(v6))+104:]))
												if t50 == i32(-1) {
													goto l16
												}
												t51 := int32(load32(m.memory[int64(uint32(v6))+112:]))
												store32(m.memory[int64(uint32(v6))+88:], uint32(t51))
												t52 := int64(load64(m.memory[int64(uint32(v6))+104:]))
												store64(m.memory[int64(uint32(v6))+80:], uint64(t52))
												goto l17
											}
										l16:
											store32(m.memory[int64(uint32(v6))+72:], uint32(i32(5)))
											store32(m.memory[int64(uint32(v6))+68:], uint32(v6+i32(100)))
											m.fn73(v6+i32(80), i32(0x100051), v6+i32(68))
										l17:
											t53 := int32(load32(m.memory[uint32(v15):]))
											t54 := int32(load32(m.memory[uint32(v16):]))
											m.fn1046(v6+i32(48), t53, t54, i32(1074680), i32(46), i32(1085308), i32(10))
											v15 = i32(0)
											v16 = i32(0)
											{
												t55 := int32(load32(m.memory[int64(uint32(v6))+48:]))
												v17 = t55
												if v17 == 0 {
													goto l18
												}
												t56 := int32(load32(m.memory[int64(uint32(v6))+52:]))
												t57 := m.fn15(v17, t56, i32(1082611), i32(7))
												v16 = t57
											}
										l18:
											{
												t58 := int32(load32(m.memory[uint32(v1+i32(28)):]))
												t59 := int32(load32(m.memory[uint32(v1+i32(32)):]))
												t60 := m.fn886(t58, t59, i32(1074680), i32(46), i32(1085318), i32(9))
												v1 = t60
												if v1 != 0 {
													m.fn1306(v6+i32(104), v1, v2)
													t61 := int32(load32(m.memory[int64(uint32(v6))+116:]))
													v17 = t61
													t62 := int32(load32(m.memory[int64(uint32(v6))+112:]))
													v1 = t62
													t63 := int32(load32(m.memory[int64(uint32(v6))+108:]))
													v15 = t63
													t64 := int32(load32(m.memory[int64(uint32(v6))+104:]))
													v18 = t64
													if v18 == i32(-1) {
														goto l20
													}
													t65 := int64(load64(m.memory[int64(uint32(v6))+120:]))
													store64(m.memory[int64(uint32(v0))+16:], uint64(t65))
													store32(m.memory[int64(uint32(v0))+12:], uint32(v17))
													store32(m.memory[int64(uint32(v0))+8:], uint32(v1))
													store32(m.memory[int64(uint32(v0))+4:], uint32(v15))
													store32(m.memory[uint32(v0):], uint32(v18))
													t66 := int32(load32(m.memory[int64(uint32(v6))+80:]))
													t67 := int32(load32(m.memory[int64(uint32(v6))+84:]))
													m.fn16(t66, t67)
													goto l4
												}
												v1 = i32(8)
												v17 = i32(0)
												goto l20
											}
										}
										t40 := m.fn15(v15, v16, i32(1078484), i32(10))
										if t40 != 0 {
											goto l3
										}
										t41 := m.fn15(v15, v16, i32(1085259), i32(15))
										if t41 != 0 {
											goto l3
										}
										t42 := m.fn15(v15, v16, i32(1085274), i32(15))
										if t42 == 0 {
											goto l5
										}
										goto l3
									}
								}
							l11:
								t27 := int32(load32(m.memory[uint32(v1+i32(16)):]))
								t28 := int32(load32(m.memory[uint32(v1+i32(20)):]))
								m.fn1046(v6+i32(40), t27, t28, i32(1074680), i32(46), i32(1073713), i32(4))
								t29 := int32(load32(m.memory[int64(uint32(v6))+40:]))
								v1 = t29
								if v1 == 0 {
									goto l3
								}
								t30 := int32(load32(m.memory[int64(uint32(v6))+44:]))
								m.fn51(v10, v1, t30)
								store32(m.memory[int64(uint32(v6))+104:], uint32(i32(6)))
								m.fn1340(v4, v6+i32(104))
								goto l3
							}
						}
					}
				}
				t5 := int32(load32(m.memory[uint32(v1+i32(8)):]))
				t6 := int32(load32(m.memory[uint32(v1+i32(12)):]))
				m.fn865(v6+i32(104), t5, t6)
				t7 := int32(load32(m.memory[int64(uint32(v6))+108:]))
				t8 := v6 + i32(80)
				v1 = t7
				t9 := int32(load32(m.memory[int64(uint32(v6))+112:]))
				m.fn1406(t8, v1, t9)
				t10 := int32(load32(m.memory[int64(uint32(v6))+104:]))
				m.fn16(t10, v1)
				t11 := int32(load32(m.memory[int64(uint32(v6))+88:]))
				if t11 == 0 {
					t99 := int32(load32(m.memory[int64(uint32(v6))+80:]))
					t100 := int32(load32(m.memory[int64(uint32(v6))+84:]))
					m.fn16(t99, t100)
					goto l3
				}
				t12 := int32(load32(m.memory[int64(uint32(v6))+88:]))
				store32(m.memory[int64(uint32(v10))+8:], uint32(t12))
				t13 := int64(load64(m.memory[int64(uint32(v6))+80:]))
				store64(m.memory[uint32(v10):], uint64(t13))
				store32(m.memory[int64(uint32(v6))+120:], uint32(v8))
				store32(m.memory[int64(uint32(v6))+104:], uint32(i32(3)))
				m.fn1340(v4, v6+i32(104))
				goto l3
			}
		l22:
			m.fn1525(v4, v6+i32(68))
			goto l27
		l25:
			m.fn1516(v6+i32(104), v1, v2, v15, v4, v5)
			t102 := int32(load32(m.memory[int64(uint32(v6))+104:]))
			v1 = t102
			if v1 == i32(-1) {
				goto l3
			}
			t103 := int32(load32(m.memory[int64(uint32(v6))+124:]))
			store32(m.memory[int64(uint32(v0))+20:], uint32(t103))
			t104 := int64(load64(m.memory[int64(uint32(v6))+116:]))
			store64(m.memory[int64(uint32(v0))+12:], uint64(t104))
			t105 := int64(load64(m.memory[int64(uint32(v6))+108:]))
			store64(m.memory[int64(uint32(v0))+4:], uint64(t105))
			store32(m.memory[uint32(v0):], uint32(v1))
			goto l4
		}
	l24:
		store32(m.memory[int64(uint32(v6))+80:], uint32(v1))
		{
			t106 := int32(load32(m.memory[int64(uint32(v6))+72:]))
			t107 := int32(load32(m.memory[int64(uint32(v6))+76:]))
			t108 := m.fn23(t106, t107)
			if t108 != 0 {
				goto l28
			}
			t109 := int64(load64(m.memory[int64(uint32(v6))+68:]))
			store64(m.memory[uint32(v11):], uint64(t109))
			t110 := int32(load32(m.memory[int64(uint32(v6))+76:]))
			store32(m.memory[int64(uint32(v11))+8:], uint32(t110))
			t111 := int64(load64(m.memory[int64(uint32(v6))+88:]))
			store64(m.memory[int64(uint32(v6))+112:], uint64(t111))
			t112 := int64(load64(m.memory[int64(uint32(v6))+80:]))
			store64(m.memory[int64(uint32(v6))+104:], uint64(t112))
			m.fn1340(v4, v6+i32(104))
			goto l3
		}
	l28:
		m.fn1525(v4, v6+i32(68))
		t113 := int32(load32(m.memory[int64(uint32(v6))+84:]))
		t114 := int32(load32(m.memory[int64(uint32(v6))+88:]))
		m.fn16(t113, t114)
	}
l27:
	m.fn894(v6 + i32(68))
	goto l3
l20:
	{
		t115 := int32(load32(m.memory[uint32(v2):]))
		if t115 != 0 {
			m.fn1326(i32(1085328))
			panic("unreachable")
		}
		store32(m.memory[uint32(v2):], uint32(i32(-1)))
		t116 := int32(load32(m.memory[int64(uint32(v6))+84:]))
		t117 := int32(load32(m.memory[int64(uint32(v6))+88:]))
		m.fn31(v6+i32(104), t116, t117)
		store32(m.memory[int64(uint32(v6))+124:], uint32(v17))
		store32(m.memory[int64(uint32(v6))+120:], uint32(v1))
		store32(m.memory[int64(uint32(v6))+116:], uint32(v15))
		m.memory[int64(uint32(v6))+128] = byte(v16)
		m.fn1230(v7, v6+i32(104))
		t118 := int32(load32(m.memory[uint32(v2):]))
		store32(m.memory[uint32(v2):], uint32(t118+i32(1)))
		t119 := int32(load32(m.memory[int64(uint32(v6))+88:]))
		store32(m.memory[int64(uint32(v10))+8:], uint32(t119))
		t120 := int64(load64(m.memory[int64(uint32(v6))+80:]))
		store64(m.memory[uint32(v10):], uint64(t120))
		store32(m.memory[int64(uint32(v6))+104:], uint32(i32(7)))
		m.fn1340(v4, v6+i32(104))
		goto l3
	}
l5:
	{
		t121 := m.fn847(v1, i32(1074120), i32(49), i32(1081740), i32(5))
		if t121 != 0 {
			m.fn1338(v6+i32(104), v1, v2, v4, v5)
			t126 := int32(load32(m.memory[int64(uint32(v6))+104:]))
			v1 = t126
			if v1 == i32(-1) {
				goto l3
			}
			t127 := int32(load32(m.memory[int64(uint32(v6))+124:]))
			store32(m.memory[int64(uint32(v0))+20:], uint32(t127))
			t128 := int64(load64(m.memory[int64(uint32(v6))+116:]))
			store64(m.memory[int64(uint32(v0))+12:], uint64(t128))
			t129 := int64(load64(m.memory[int64(uint32(v6))+108:]))
			store64(m.memory[int64(uint32(v0))+4:], uint64(t129))
			store32(m.memory[uint32(v0):], uint32(v1))
			goto l4
		}
		m.fn1516(v6+i32(104), v1, v2, v3, v4, v5)
		t122 := int32(load32(m.memory[int64(uint32(v6))+104:]))
		v1 = t122
		if v1 == i32(-1) {
			goto l3
		}
		t123 := int32(load32(m.memory[int64(uint32(v6))+124:]))
		store32(m.memory[int64(uint32(v0))+20:], uint32(t123))
		t124 := int64(load64(m.memory[int64(uint32(v6))+116:]))
		store64(m.memory[int64(uint32(v0))+12:], uint64(t124))
		t125 := int64(load64(m.memory[int64(uint32(v6))+108:]))
		store64(m.memory[int64(uint32(v0))+4:], uint64(t125))
		store32(m.memory[uint32(v0):], uint32(v1))
		goto l4
	}
l3:
	v14 = v14 + i32(44)
	goto l31
l4:
	m.g0 = v6 + i32(144)
}
func (m *Module) fn1517(v0, v1, v2 int32) int32 {
	var v3 int64
	var v4, v5 int32
	var v6 int64
	var v7 int32
	var v8 int64
	var v9, v10, v11 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
		if t0 != 0 {
			t1 := int64(load64(m.memory[int64(uint32(v0))+16:]))
			t2 := int64(load64(m.memory[int64(uint32(v0))+24:]))
			t3 := m.fn29(t1, t2, v1, v2)
			v3 = t3
			t4 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			v4 = t4
			v5 = v4 & int32(v3)
			v6 = int64(uint64(v3)>>25) & i64(127) * i64(72340172838076673)
			t5 := int32(load32(m.memory[uint32(v0):]))
			v0 = t5
			v7 = i32(0)
			var _ int32
		l5:
			{
				t7 := int64(load64(m.memory[uint32(v0+v5):]))
				v8 = t7
				v3 = v8 ^ v6
				v3 = (v3 ^ i64(-1)) & (v3 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
				{
				l3:
					{
						var p8 int32
						if v3 == 0 {
							p8 = 1
						}
						v9 = p8
						if v9 != 0 {
							goto l1
						}
						t9 := v1
						t10 := v2
						t11 := v0
						v10 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v3))))>>3) + v5) & v4
						v11 = t11 + (i32(0)-v10)*i32(416)
						t12 := int32(load32(m.memory[uint32(v11+i32(-412)):]))
						t13 := int32(load32(m.memory[uint32(v11+i32(-408)):]))
						t14 := m.fn15(t9, t10, t12, t13)
						if t14 != 0 {
							goto l2
						}
						v3 = (v3 + i64(-1)) & v3
						goto l3
					}
				l1:
					if v8&(v8<<1)&i64(-0x7f7f7f7f7f7f7f80) == 0 {
						t17 := v5
						v7 = v7 + i32(8)
						v5 = (t17 + v7) & v4
						goto l5
					}
				l2:
					p15 := v0 + (i32(0)-v10)*i32(416)
					if v9 != 0 {
						p15 = i32(0)
					}
					p16 := p15 + i32(-400)
					if v9 != 0 {
						p16 = i32(0)
					}
					return p16
				}
			}
		}
		return i32(0)
	}
}
func (m *Module) fn1518(v0 int32) {
	var v1 int32
	t0 := m.g0
	v1 = t0 - i32(16)
	m.g0 = v1
	store64(m.memory[uint32(v1):], uint64(int64(uint32(i32(192)))<<32|int64(uint32(v1+i32(15)))))
	m.fn91(i32(1052692), v1, v0)
	panic("unreachable")
}
func (m *Module) fn1519(v0, v1 int32, v2 int64) {
	var v3 int32
	var v4, v5, v6 int64
	t0 := m.g0
	v3 = t0 - i32(64)
	m.g0 = v3
	t1 := int64(load64(m.memory[int64(uint32(v0))+16:]))
	v4 = t1
	store32(m.memory[int64(uint32(v0))+16:], uint32(i32(0)))
	t2 := int64(load64(m.memory[int64(uint32(v0))+8:]))
	v5 = t2
	store64(m.memory[int64(uint32(v0))+8:], uint64(i64(0x400000000)))
	t3 := int64(load64(m.memory[uint32(v0):]))
	v6 = t3
	store64(m.memory[uint32(v0):], uint64(v2))
	store64(m.memory[int64(uint32(v3))+24:], uint64(v4))
	store64(m.memory[int64(uint32(v3))+16:], uint64(v5))
	store64(m.memory[int64(uint32(v3))+8:], uint64(v6))
	{
		if int32(v4) == 0 {
			goto l0
		}
		t4 := int64(load64(m.memory[int64(uint32(v3))+24:]))
		store64(m.memory[int64(uint32(v3))+56:], uint64(t4))
		t5 := int64(load64(m.memory[int64(uint32(v3))+16:]))
		store64(m.memory[int64(uint32(v3))+48:], uint64(t5))
		t6 := int64(load64(m.memory[int64(uint32(v3))+8:]))
		store64(m.memory[int64(uint32(v3))+40:], uint64(t6))
		store32(m.memory[int64(uint32(v3))+32:], uint32(i32(-0x7fffffff)))
		m.fn338(v1, v3+i32(32))
		goto l1
	}
l0:
	m.fn971(v3 + i32(16))
l1:
	m.g0 = v3 + i32(64)
}
func (m *Module) fn1520(v0, v1, v2, v3, v4, v5 int32) {
	var v6, v7, v8 int32
	var v9 int64
	var v10 int32
	var v11 int64
	var v12, v13 int32
	var v14 int64
	var v15, v16, v17, v18 int32
	t0 := m.g0
	v6 = t0 - i32(80)
	m.g0 = v6
	m.fn1328(v6, v2, v3, v4, v5)
	{
		t1 := int32(load32(m.memory[int64(uint32(v1))+496:]))
		v5 = t1
		if uint32(v5) >= uint32(i32(0x7fffffff)) {
			m.fn1518(i32(1077296))
			panic("unreachable")
		}
		store32(m.memory[int64(uint32(v1))+496:], uint32(v5+i32(1)))
		t2 := int32(load32(m.memory[int64(uint32(v1))+516:]))
		if t2 == 0 {
			goto l1
		}
		t3 := int64(load64(m.memory[int64(uint32(v1))+520:]))
		t4 := int64(load64(m.memory[uint32(v1+i32(528)):]))
		t5 := int32(load32(m.memory[int64(uint32(v6))+4:]))
		v7 = t5
		t6 := int32(load32(m.memory[int64(uint32(v6))+8:]))
		t7 := v7
		v8 = t6
		t8 := m.fn540(t3, t4, t7, v8)
		v9 = t8
		t9 := int32(load32(m.memory[int64(uint32(v1))+508:]))
		v10 = t9
		v4 = v10 & int32(v9)
		v11 = int64(uint64(v9)>>25) & i64(127) * i64(72340172838076673)
		t10 := int32(load32(m.memory[int64(uint32(v1))+504:]))
		v12 = t10
		v13 = i32(0)
	l7:
		{
			t11 := int64(load64(m.memory[uint32(v12+v4):]))
			v14 = t11
			v9 = v14 ^ v11
			v9 = (v9 ^ i64(-1)) & (v9 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
		l4:
			{
				if v9 == 0 {
					if v14&(v14<<1)&i64(-0x7f7f7f7f7f7f7f80) == 0 {
						t18 := v4
						v13 = v13 + i32(8)
						v4 = (t18 + v13) & v10
						goto l7
					}
					goto l1
				}
				t12 := v7
				t13 := v8
				v15 = v12 - (int32(uint32(int64(bits.TrailingZeros64(uint64(v9))))>>3)+v4)&v10<<4
				t14 := int32(load32(m.memory[uint32(v15+i32(-12)):]))
				t15 := int32(load32(m.memory[uint32(v15+i32(-8)):]))
				t16 := m.fn544(t12, t13, t14, t15)
				if t16 != 0 {
					store32(m.memory[uint32(v0):], uint32(i32(-1)))
					t17 := int32(load32(m.memory[uint32(v15+i32(-4)):]))
					store32(m.memory[int64(uint32(v0))+4:], uint32(t17))
					store32(m.memory[int64(uint32(v1))+496:], uint32(v5))
					goto l6
				}
				v9 = (v9 + i64(-1)) & v9
				goto l4
			}
		}
	}
l1:
	store32(m.memory[int64(uint32(v1))+496:], uint32(v5))
	v12 = i32(0)
	store32(m.memory[int64(uint32(v6))+20:], uint32(i32(0)))
	store64(m.memory[int64(uint32(v6))+12:], uint64(i64(0x400000000)))
	m.fn27(v6 + i32(24))
	{
		t19 := int32(load32(m.memory[int64(uint32(v1))+412:]))
		if t19 == 0 {
			goto l8
		}
		t20 := int64(load64(m.memory[int64(uint32(v1))+416:]))
		t21 := int64(load64(m.memory[uint32(v1+i32(424)):]))
		t22 := int32(load32(m.memory[int64(uint32(v6))+4:]))
		t23 := int32(load32(m.memory[int64(uint32(v6))+8:]))
		t24 := m.fn540(t20, t21, t22, t23)
		v9 = t24
		t25 := int32(load32(m.memory[int64(uint32(v1))+400:]))
		t26 := int32(load32(m.memory[uint32(v1+i32(404)):]))
		t27 := m.fn1512(t25, t26, v9, v6)
		v5 = t27
		if v5 == 0 {
			goto l8
		}
		t28 := int32(load32(m.memory[uint32(v5+i32(-20)):]))
		v4 = t28
		t29 := int32(load32(m.memory[uint32(v5+i32(-24)):]))
		v5 = t29
		goto l9
	}
l8:
	v5 = i32(0)
l9:
	v13 = v1 + i32(504)
	v10 = v1 + i32(400)
	v16 = v1 + i32(404)
	v17 = v1 + i32(424)
	v8 = i32(4)
	v15 = i32(0)
l25:
	{
		if v5 == 0 {
			v4 = i32(255)
			{
				t34 := int32(load32(m.memory[int64(uint32(v1))+476:]))
				if t34 == 0 {
					goto l12
				}
				t35 := int64(load64(m.memory[int64(uint32(v1))+480:]))
				t36 := int64(load64(m.memory[uint32(v1+i32(488)):]))
				t37 := m.fn29(t35, t36, v2, v3)
				v9 = t37
				t38 := int32(load32(m.memory[int64(uint32(v1))+468:]))
				v8 = t38
				v5 = v8 & int32(v9)
				v11 = int64(uint64(v9)>>25) & i64(127) * i64(72340172838076673)
				t39 := int32(load32(m.memory[int64(uint32(v1))+464:]))
				v4 = t39
				v18 = i32(0)
			l17:
				{
					t40 := int64(load64(m.memory[uint32(v4+v5):]))
					v14 = t40
					v9 = v14 ^ v11
					v9 = (v9 ^ i64(-1)) & (v9 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
				l15:
					{
						if v9 == 0 {
							if v14&(v14<<1)&i64(-0x7f7f7f7f7f7f7f80) == 0 {
								t46 := v5
								v18 = v18 + i32(8)
								v5 = (t46 + v18) & v8
								goto l17
							}
							v4 = i32(255)
							goto l12
						}
						t41 := v2
						t42 := v3
						v7 = v4 - (int32(uint32(int64(bits.TrailingZeros64(uint64(v9))))>>3)+v5)&v8<<4
						t43 := int32(load32(m.memory[uint32(v7+i32(-12)):]))
						t44 := int32(load32(m.memory[uint32(v7+i32(-8)):]))
						t45 := m.fn15(t41, t42, t43, t44)
						if t45 != 0 {
							goto l14
						}
						v9 = (v9 + i64(-1)) & v9
						goto l15
					}
				l14:
				}
				t47 := int32(load32(m.memory[uint32(v7+i32(-4)):]))
				v4 = t47
			}
		l12:
			v5 = i32(0) - v12
			p48 := v4
			if v4&i32(255) == i32(255) {
				p48 = i32(33686018)
			}
			v4 = p48
			t49 := int32(load32(m.memory[int64(uint32(v6))+16:]))
			v18 = t49
			v12 = v18 + v15<<3
			t50 := int32(load32(m.memory[int64(uint32(v6))+12:]))
			v3 = t50
		l21:
			{
				if v5 == 0 {
					m.fn639(v3, v18)
					store32(m.memory[uint32(v0):], uint32(i32(-1)))
					store32(m.memory[int64(uint32(v0))+4:], uint32(v4))
					t80 := int32(load32(m.memory[int64(uint32(v6))+24:]))
					t81 := int32(load32(m.memory[int64(uint32(v6))+28:]))
					m.fn56(t80, t81)
					goto l6
				}
				v15 = v12 + i32(-4)
				t51 := v10
				v12 = v12 + i32(-8)
				t52 := int32(load32(m.memory[uint32(v12):]))
				v7 = t52
				t53 := int32(load32(m.memory[uint32(v15):]))
				t54 := v7
				v8 = t53
				t55 := m.fn1522(t51, t54, v8)
				v15 = t55
				if v15 == 0 {
					m.fn633(i32(1087080), i32(22), i32(1077312))
					panic("unreachable")
				}
				t56 := int32(load32(m.memory[uint32(v15):]))
				v15 = t56
				t57 := int32(load32(m.memory[uint32(v15+i32(28)):]))
				t58 := int32(load32(m.memory[uint32(v15+i32(32)):]))
				t59 := m.fn1327(t57, t58)
				v15 = t59
				t60 := int32(load32(m.memory[int64(uint32(v1))+496:]))
				if t60 != 0 {
					m.fn1326(i32(1077328))
					panic("unreachable")
				}
				t61 := fn1373(v4, v15)
				v4 = t61
				store32(m.memory[int64(uint32(v1))+496:], uint32(i32(-1)))
				m.fn51(v6+i32(68), v7, v8)
				m.fn1105(v13, v6+i32(68), v4)
				t62 := int32(load32(m.memory[int64(uint32(v1))+496:]))
				store32(m.memory[int64(uint32(v1))+496:], uint32(t62+i32(1)))
				v5 = v5 + i32(8)
				goto l21
			}
		}
		store32(m.memory[int64(uint32(v6))+64:], uint32(v4))
		store32(m.memory[int64(uint32(v6))+60:], uint32(v5))
		t30 := m.fn1521(v6+i32(24), v5, v4)
		if t30 != 0 {
			{
				t63 := int32(load32(m.memory[int64(uint32(v6))+12:]))
				if v15 != t63 {
					goto l22
				}
				m.fn1523(v6 + i32(12))
				t64 := int32(load32(m.memory[int64(uint32(v6))+16:]))
				v8 = t64
			}
		l22:
			v7 = v8 + v12
			store32(m.memory[uint32(v7):], uint32(v5))
			store32(m.memory[uint32(v7+i32(4)):], uint32(v4))
			t65 := v6
			v15 = v15 + i32(1)
			store32(m.memory[int64(uint32(t65))+20:], uint32(v15))
			t66 := m.fn1522(v10, v5, v4)
			v5 = t66
			if v5 == 0 {
				goto l23
			}
			t67 := int32(load32(m.memory[int64(uint32(v5))+4:]))
			if t67 == i32(-1) {
				goto l23
			}
			t68 := int32(load32(m.memory[int64(uint32(v1))+412:]))
			if t68 == 0 {
				goto l23
			}
			t69 := int64(load64(m.memory[int64(uint32(v1))+416:]))
			t70 := int64(load64(m.memory[uint32(v17):]))
			t71 := int32(load32(m.memory[int64(uint32(v5))+8:]))
			v7 = t71
			t72 := int32(load32(m.memory[int64(uint32(v5))+12:]))
			t73 := v7
			v18 = t72
			t74 := m.fn29(t69, t70, t73, v18)
			v9 = t74
			v5 = i32(0)
			t75 := int32(load32(m.memory[int64(uint32(v1))+400:]))
			t76 := int32(load32(m.memory[uint32(v16):]))
			t77 := m.fn1524(t75, t76, v9, v7, v18)
			v7 = t77
			if v7 == 0 {
				goto l24
			}
			t78 := int32(load32(m.memory[uint32(v7+i32(-20)):]))
			v4 = t78
			t79 := int32(load32(m.memory[uint32(v7+i32(-24)):]))
			v5 = t79
			goto l24
		}
		store32(m.memory[int64(uint32(v6))+72:], uint32(i32(71)))
		store32(m.memory[int64(uint32(v6))+68:], uint32(v6+i32(60)))
		m.fn73(v0, i32(1049807), v6+i32(68))
		store32(m.memory[int64(uint32(v0))+12:], uint32(i32(-1)))
		t31 := int32(load32(m.memory[int64(uint32(v6))+24:]))
		t32 := int32(load32(m.memory[int64(uint32(v6))+28:]))
		m.fn56(t31, t32)
		t33 := int32(load32(m.memory[int64(uint32(v6))+12:]))
		m.fn639(t33, v8)
		goto l6
	}
l23:
	v5 = i32(0)
l24:
	v12 = v12 + i32(8)
	goto l25
l6:
	t82 := int32(load32(m.memory[uint32(v6):]))
	t83 := int32(load32(m.memory[int64(uint32(v6))+4:]))
	m.fn16(t82, t83)
	m.g0 = v6 + i32(80)
}
func (m *Module) fn1521(v0, v1, v2 int32) int32 {
	var v3 int32
	var v4 int64
	var v5, v6, v7, v8, v9, v10 int32
	t0 := m.g0
	v3 = t0 - i32(32)
	m.g0 = v3
	store32(m.memory[int64(uint32(v3))+16:], uint32(v2))
	store32(m.memory[int64(uint32(v3))+12:], uint32(v1))
	t1 := int64(load64(m.memory[int64(uint32(v0))+16:]))
	t2 := int64(load64(m.memory[int64(uint32(v0))+24:]))
	t3 := m.fn24(t1, t2, v1, v2)
	v4 = t3
	store32(m.memory[int64(uint32(v3))+20:], uint32(v3+i32(12)))
	{
		t4 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		if t4 != 0 {
			goto l0
		}
		_ = m.fn692(v0, v0+i32(16))
	}
l0:
	store32(m.memory[int64(uint32(v3))+24:], uint32(v3+i32(20)))
	store32(m.memory[int64(uint32(v3))+28:], uint32(v0))
	t6 := int32(load32(m.memory[uint32(v0):]))
	t7 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	m.fn69(v3, t6, t7, v4, v3+i32(24), i32(193))
	{
		t8 := int32(load32(m.memory[uint32(v3):]))
		v5 = t8
		if v5 != i32(1) {
			goto l1
		}
		t9 := int32(load32(m.memory[uint32(v0):]))
		v6 = t9
		t10 := int32(load32(m.memory[int64(uint32(v3))+4:]))
		t11 := v6
		v7 = t10
		v8 = t11 + v7
		t12 := int32(m.memory[uint32(v8)])
		v9 = t12
		t13 := v8
		v10 = int32(uint32(int32(v4)) >> 25)
		m.memory[uint32(t13)] = byte(v10)
		t14 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		m.memory[uint32(v6+t14&(v7+i32(-8))+i32(8))] = byte(v10)
		t15 := int32(load32(m.memory[int64(uint32(v0))+12:]))
		store32(m.memory[int64(uint32(v0))+12:], uint32(t15+i32(1)))
		t16 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		store32(m.memory[int64(uint32(v0))+8:], uint32(t16-v9&i32(1)))
		v0 = v6 - v7<<3
		store32(m.memory[uint32(v0+i32(-4)):], uint32(v2))
		store32(m.memory[uint32(v0+i32(-8)):], uint32(v1))
	}
l1:
	m.g0 = v3 + i32(32)
	return v5
}
func (m *Module) fn1522(v0, v1, v2 int32) int32 {
	var v3 int64
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
		if t0 != 0 {
			t1 := int64(load64(m.memory[int64(uint32(v0))+16:]))
			t2 := int64(load64(m.memory[int64(uint32(v0))+24:]))
			t3 := m.fn29(t1, t2, v1, v2)
			v3 = t3
			t4 := int32(load32(m.memory[uint32(v0):]))
			t5 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t6 := m.fn1524(t4, t5, v3, v1, v2)
			v0 = t6
			p7 := i32(0)
			if v0 != 0 {
				p7 = v0 + i32(-16)
			}
			return p7
		}
		return i32(0)
	}
}
func (m *Module) fn1523(v0 int32) {
	var v1 int32
	t0 := m.g0
	v1 = t0 - i32(16)
	m.g0 = v1
	t1 := int32(load32(m.memory[uint32(v0):]))
	m.fn1629(v1+i32(8), v0, t1, i32(1), i32(4), i32(8))
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
func (m *Module) fn1524(v0, v1 int32, v2 int64, v3, v4 int32) int32 {
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
func (m *Module) fn1525(v0, v1 int32) {
	t0 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	t1 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	m.fn1470(v0, t0, t1)
	store32(m.memory[int64(uint32(v1))+8:], uint32(i32(0)))
}
func (m *Module) fn1526(v0, v1, v2 int32) int32 {
	var v3, v4 int32
	t0 := m.g0
	v3 = t0 - i32(16)
	m.g0 = v3
	v4 = i32(0)
	{
		if v0 == 0 {
			goto l0
		}
		t1 := int32(load32(m.memory[uint32(v0+i32(28)):]))
		t2 := int32(load32(m.memory[uint32(v0+i32(32)):]))
		t3 := m.fn886(t1, t2, i32(1072544), i32(60), v1, v2)
		v0 = t3
		if v0 == 0 {
			goto l0
		}
		t4 := int32(load32(m.memory[uint32(v0+i32(16)):]))
		t5 := int32(load32(m.memory[uint32(v0+i32(20)):]))
		m.fn1046(v3, t4, t5, i32(1072544), i32(60), i32(1073156), i32(3))
		t6 := int32(load32(m.memory[uint32(v3):]))
		v0 = t6
		if v0 == 0 {
			goto l0
		}
		t7 := int32(load32(m.memory[int64(uint32(v3))+4:]))
		m.fn197(v3+i32(8), v0, t7)
		t8 := int32(load32(m.memory[int64(uint32(v3))+12:]))
		v0 = t8
		p9 := i32(1000)
		if uint32(v0) < uint32(i32(1000)) {
			p9 = v0
		}
		t10 := int32(m.memory[int64(uint32(v3))+8])
		p11 := p9
		if t10 != 0 {
			p11 = i32(0)
		}
		v4 = p11
	}
l0:
	m.g0 = v3 + i32(16)
	return v4
}
