package core

import (
	"math"
)

func (m *Module) fn1392(v0, v1 int32) int32 {
	m.fn1487(v0, v1)
	t0 := m.fn1488(v0, v1)
	t1 := int32(load32(m.memory[int64(uint32(t0))+8:]))
	var p2 int32
	if t1 != i32(0) {
		p2 = 1
	}
	return p2
}
func (m *Module) fn1393(v0, v1, v2, v3 int32) {
	var v4, v5 int32
	var v6, v7, v8 int64
	t0 := m.g0
	v4 = t0 - i32(64)
	m.g0 = v4
	t1 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	t2 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	m.fn1489(v4+i32(32), t1, t2, v2)
	v2 = v4 + i32(32) | i32(4)
	{
		{
			t3 := int32(load32(m.memory[int64(uint32(v4))+32:]))
			v5 = t3
			if v5 != i32(-2) {
				goto l0
			}
			t4 := int64(load64(m.memory[int64(uint32(v2))+16:]))
			t5 := v4
			v6 = t4
			store64(m.memory[int64(uint32(t5))+16:], uint64(v6))
			t6 := int64(load64(m.memory[int64(uint32(v2))+8:]))
			t7 := v4
			v7 = t6
			store64(m.memory[int64(uint32(t7))+8:], uint64(v7))
			t8 := int64(load64(m.memory[uint32(v2):]))
			t9 := v4
			v8 = t8
			store64(m.memory[uint32(t9):], uint64(v8))
			store64(m.memory[int64(uint32(v0))+16:], uint64(v6))
			store64(m.memory[int64(uint32(v0))+8:], uint64(v7))
			store64(m.memory[uint32(v0):], uint64(v8))
			goto l1
		}
	l0:
		t10 := int32(load32(m.memory[int64(uint32(v2))+24:]))
		store32(m.memory[int64(uint32(v4))+24:], uint32(t10))
		t11 := int64(load64(m.memory[int64(uint32(v2))+16:]))
		store64(m.memory[int64(uint32(v4))+16:], uint64(t11))
		t12 := int64(load64(m.memory[int64(uint32(v2))+8:]))
		store64(m.memory[int64(uint32(v4))+8:], uint64(t12))
		t13 := int64(load64(m.memory[uint32(v2):]))
		store64(m.memory[uint32(v4):], uint64(t13))
		{
			if v5 == i32(-1) {
				goto l2
			}
			store32(m.memory[int64(uint32(v4))+32:], uint32(v5))
			t14 := int64(load64(m.memory[uint32(v4):]))
			store64(m.memory[int64(uint32(v4))+36:], uint64(t14))
			t15 := int64(load64(m.memory[int64(uint32(v4))+8:]))
			store64(m.memory[int64(uint32(v4))+44:], uint64(t15))
			t16 := int64(load64(m.memory[int64(uint32(v4))+16:]))
			store64(m.memory[int64(uint32(v4))+52:], uint64(t16))
			t17 := int32(load32(m.memory[int64(uint32(v4))+24:]))
			store32(m.memory[int64(uint32(v4))+60:], uint32(t17))
			m.fn1487(v1, v3)
			t18 := m.fn1488(v1, v3)
			m.fn338(t18, v4+i32(32))
		}
	l2:
		store32(m.memory[uint32(v0):], uint32(i32(-1)))
	}
l1:
	m.g0 = v4 + i32(64)
}
func (m *Module) fn1394(v0, v1 int32) {
	var v2, v3, v4 int32
	var v5, v6, v7 int64
	t0 := m.g0
	v2 = t0 - i32(64)
	m.g0 = v2
	t1 := int32(load32(m.memory[int64(uint32(v1))+336:]))
	t2 := int32(load32(m.memory[int64(uint32(v1))+340:]))
	m.fn1489(v2+i32(32), t1, t2, i32(1))
	v3 = v2 + i32(32) | i32(4)
	{
		{
			t3 := int32(load32(m.memory[int64(uint32(v2))+32:]))
			v4 = t3
			if v4 != i32(-2) {
				goto l0
			}
			t4 := int64(load64(m.memory[int64(uint32(v3))+16:]))
			t5 := v2
			v5 = t4
			store64(m.memory[int64(uint32(t5))+16:], uint64(v5))
			t6 := int64(load64(m.memory[int64(uint32(v3))+8:]))
			t7 := v2
			v6 = t6
			store64(m.memory[int64(uint32(t7))+8:], uint64(v6))
			t8 := int64(load64(m.memory[uint32(v3):]))
			t9 := v2
			v7 = t8
			store64(m.memory[uint32(t9):], uint64(v7))
			store64(m.memory[int64(uint32(v0))+16:], uint64(v5))
			store64(m.memory[int64(uint32(v0))+8:], uint64(v6))
			store64(m.memory[uint32(v0):], uint64(v7))
			goto l1
		}
	l0:
		t10 := int32(load32(m.memory[int64(uint32(v3))+24:]))
		store32(m.memory[int64(uint32(v2))+24:], uint32(t10))
		t11 := int64(load64(m.memory[int64(uint32(v3))+16:]))
		store64(m.memory[int64(uint32(v2))+16:], uint64(t11))
		t12 := int64(load64(m.memory[int64(uint32(v3))+8:]))
		store64(m.memory[int64(uint32(v2))+8:], uint64(t12))
		t13 := int64(load64(m.memory[uint32(v3):]))
		store64(m.memory[uint32(v2):], uint64(t13))
		{
			if v4 == i32(-1) {
				goto l2
			}
			store32(m.memory[int64(uint32(v2))+32:], uint32(v4))
			t14 := int64(load64(m.memory[uint32(v2):]))
			store64(m.memory[int64(uint32(v2))+36:], uint64(t14))
			t15 := int64(load64(m.memory[int64(uint32(v2))+8:]))
			store64(m.memory[int64(uint32(v2))+44:], uint64(t15))
			t16 := int64(load64(m.memory[int64(uint32(v2))+16:]))
			store64(m.memory[int64(uint32(v2))+52:], uint64(t16))
			t17 := int32(load32(m.memory[int64(uint32(v2))+24:]))
			store32(m.memory[int64(uint32(v2))+60:], uint32(t17))
			m.fn1396(v1)
			m.fn338(v1+i32(392), v2+i32(32))
		}
	l2:
		store32(m.memory[uint32(v0):], uint32(i32(-1)))
	}
l1:
	m.g0 = v2 + i32(64)
}
func (m *Module) fn1395(v0 int32) {
	t0 := int32(load32(m.memory[int64(uint32(v0))+368:]))
	t1 := int32(load32(m.memory[int64(uint32(v0))+372:]))
	m.fn1397(t0, t1)
	m.fn1398(v0)
	t2 := int32(load32(m.memory[int64(uint32(v0))+296:]))
	t3 := int32(load32(m.memory[int64(uint32(v0))+300:]))
	m.fn16(t2, t3)
	m.fn894(v0 + i32(380))
	m.fn969(v0 + i32(392))
	m.fn1302(v0 + i32(404))
	m.fn1332(v0 + i32(176))
	t4 := int32(load32(m.memory[int64(uint32(v0))+96:]))
	t5 := int32(load32(m.memory[int64(uint32(v0))+100:]))
	m.fn1399(t4, t5)
	m.fn1400(v0 + i32(332))
	m.fn448(v0 + i32(192))
	t6 := int32(load32(m.memory[int64(uint32(v0))+204:]))
	t7 := int32(load32(m.memory[int64(uint32(v0))+208:]))
	m.fn911(t6, t7)
	m.fn1229(v0 + i32(216))
	t8 := int32(load32(m.memory[int64(uint32(v0))+228:]))
	t9 := int32(load32(m.memory[int64(uint32(v0))+232:]))
	m.fn16(t8, t9)
	t10 := int32(load32(m.memory[int64(uint32(v0))+240:]))
	t11 := int32(load32(m.memory[int64(uint32(v0))+244:]))
	m.fn134(t10, t11)
	m.fn1388(v0 + i32(252))
	m.fn1274(v0 + i32(128))
}
func (m *Module) fn1396(v0 int32) {
	var v1 int32
	t0 := v0 + i32(176)
	v1 = v0 + i32(392)
	m.fn1333(t0, v1)
	m.fn1351(v1, v0+i32(404))
}
func (m *Module) fn1397(v0, v1 int32) {
	m.fn136(v0, v1, i32(4), i32(44))
}
func (m *Module) fn1398(v0 int32) {
	var v1, v2, v3 int32
	t0 := m.g0
	v1 = t0 - i32(16)
	m.g0 = v1
	{
		t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		v2 = t1
		if v2 == 0 {
			goto l0
		}
		t2 := int32(load32(m.memory[uint32(v0):]))
		v3 = t2
		m.fn39(v1+i32(4), i32(8), i32(8), v2+i32(1))
		t3 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		t4 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t5 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		m.fn40(v3-t3, t4, t5)
	}
l0:
	{
		t6 := int32(load32(m.memory[int64(uint32(v0))+36:]))
		v2 = t6
		if v2 == 0 {
			goto l1
		}
		t7 := int32(load32(m.memory[int64(uint32(v0))+32:]))
		v3 = t7
		m.fn39(v1+i32(4), i32(12), i32(8), v2+i32(1))
		t8 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		t9 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t10 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		m.fn40(v3-t8, t9, t10)
	}
l1:
	m.fn1381(v0 + i32(64))
	m.g0 = v1 + i32(16)
}
func (m *Module) fn1399(v0, v1 int32) {
	var v2 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	{
		if v1 == 0 {
			goto l0
		}
		m.fn39(v2+i32(4), i32(96), i32(8), v1+i32(1))
		t1 := int32(load32(m.memory[int64(uint32(v2))+12:]))
		t2 := int32(load32(m.memory[int64(uint32(v2))+4:]))
		t3 := int32(load32(m.memory[int64(uint32(v2))+8:]))
		m.fn40(v0-t1, t2, t3)
	}
l0:
	m.g0 = v2 + i32(16)
}
func (m *Module) fn1400(v0 int32) {
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
		v4 = v2 + v3<<6
		t2 := int32(load32(m.memory[uint32(v4+i32(40)):]))
		v5 = t2
		v6 = v4 + i32(36)
		t3 := int32(load32(m.memory[uint32(v6):]))
		v7 = t3
	l2:
		if v5 == 0 {
			t4 := int32(load32(m.memory[int64(uint32(v4))+32:]))
			t5 := int32(load32(m.memory[uint32(v6):]))
			m.fn419(t4, t5)
			t6 := int32(load32(m.memory[int64(uint32(v4))+44:]))
			t7 := int32(load32(m.memory[uint32(v4+i32(48)):]))
			m.fn1490(t6, t7)
			m.fn968(v4)
			v3 = v3 + i32(1)
			goto l3
		}
		v5 = v5 + i32(-1)
		m.fn968(v7)
		v7 = v7 + i32(16)
		goto l2
	}
l0:
	t8 := int32(load32(m.memory[uint32(v0):]))
	m.fn136(t8, v2, i32(8), i32(64))
	m.fn976(v0 + i32(12))
	t9 := int32(load32(m.memory[int64(uint32(v0))+32:]))
	v5 = t9
	t10 := int32(load32(m.memory[int64(uint32(v0))+28:]))
	v4 = t10
	v7 = v4
l5:
	if v5 == 0 {
		goto l4
	}
	v5 = v5 + i32(-1)
	m.fn1332(v7)
	v7 = v7 + i32(16)
	goto l5
l4:
	t11 := int32(load32(m.memory[int64(uint32(v0))+24:]))
	m.fn136(t11, v4, i32(4), i32(16))
}
func (m *Module) fn1401(v0, v1 int32) {
	var v2, v3, v4, v5, v6, v7, v8, v9, v10, v11 int32
	var v12, v13, v14 int64
	t0 := m.g0
	v2 = t0 - i32(48)
	m.g0 = v2
	t1 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	t2 := v2
	v3 = t1
	t3 := int32(load32(m.memory[uint32(v1):]))
	t4 := v3
	v4 = t3
	t5 := int32(uint32(t4-v4) / uint32(i32(28)))
	v5 = t5
	m.fn59(t2, v5, i32(4), i32(28))
	store32(m.memory[int64(uint32(v2))+20:], uint32(i32(0)))
	t6 := int64(load64(m.memory[uint32(v2):]))
	store64(m.memory[int64(uint32(v2))+12:], uint64(t6))
	m.fn892(v2+i32(12), v5)
	v6 = v4 + i32(28)
	t7 := int32(load32(m.memory[int64(uint32(v2))+16:]))
	t8 := int32(load32(m.memory[int64(uint32(v2))+20:]))
	v7 = t8
	v5 = t7 + v7*i32(28)
	t9 := int32(load32(m.memory[int64(uint32(v1))+16:]))
	v8 = t9
	t10 := int32(load32(m.memory[int64(uint32(v1))+12:]))
	v9 = t10
	t11 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	v10 = t11
l6:
	{
		{
			if v4 == v3 {
				goto l0
			}
			v1 = v4 + i32(28)
			t12 := int32(load32(m.memory[uint32(v4):]))
			v11 = t12
			if v11 != i32(-1) {
				t17 := int64(load64(m.memory[uint32(v4+i32(20)):]))
				v12 = t17
				t18 := int64(load64(m.memory[uint32(v4+i32(12)):]))
				v13 = t18
				t19 := int64(load64(m.memory[uint32(v4+i32(4)):]))
				v14 = t19
				store32(m.memory[uint32(v5):], uint32(v11))
				store64(m.memory[uint32(v5+i32(4)):], uint64(v14))
				store64(m.memory[uint32(v5+i32(12)):], uint64(v13))
				store64(m.memory[uint32(v5+i32(20)):], uint64(v12))
				v5 = v5 + i32(28)
				v6 = v6 + i32(28)
				v7 = v7 + i32(1)
				v4 = v1
				goto l6
			}
			store32(m.memory[int64(uint32(v2))+20:], uint32(v7))
			if v3 == v1 {
				goto l2
			}
			t13 := int32(uint32(v3-v6) / uint32(i32(28)))
			v4 = t13
		l3:
			if v4 == 0 {
				goto l2
			}
			v4 = v4 + i32(-1)
			m.fn893(v6)
			v6 = v6 + i32(28)
			goto l3
		}
	l0:
		store32(m.memory[int64(uint32(v2))+20:], uint32(v7))
	l2:
		if v8 == 0 {
			goto l4
		}
		t14 := int32(load32(m.memory[int64(uint32(v10))+8:]))
		t15 := v9
		v4 = t14
		if t15 == v4 {
			goto l5
		}
		v5 = v8 * i32(28)
		if v5 == 0 {
			goto l5
		}
		t16 := int32(load32(m.memory[int64(uint32(v10))+4:]))
		v6 = t16
		memory_copy(m.memory, uint32(v6+v4*i32(28)), uint32(v6+v9*i32(28)), uint32(v5))
		goto l5
	}
l5:
	store32(m.memory[int64(uint32(v10))+8:], uint32(v8+v4))
l4:
	t20 := int32(load32(m.memory[int64(uint32(v2))+20:]))
	store32(m.memory[int64(uint32(v0))+8:], uint32(t20))
	t21 := int64(load64(m.memory[int64(uint32(v2))+12:]))
	store64(m.memory[uint32(v0):], uint64(t21))
	m.g0 = v2 + i32(48)
}
func (m *Module) fn1402(v0 int32) {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	m.fn16(t0, t1)
	t2 := int32(load32(m.memory[int64(uint32(v0))+12:]))
	t3 := int32(load32(m.memory[int64(uint32(v0))+16:]))
	m.fn1390(t2, t3)
}
func (m *Module) fn1403(v0, v1, v2, v3 int32) {
	var v4, v5, v6, v7, v8, v9, v10 int32
	var v11, v12 int64
	var v13, v14, v15, v16, v17, v18 int32
	t0 := m.g0
	v4 = t0 - i32(144)
	m.g0 = v4
	store32(m.memory[int64(uint32(v4))+116:], uint32(i32(0)))
	store64(m.memory[int64(uint32(v4))+108:], uint64(i64(0x400000000)))
	store32(m.memory[int64(uint32(v4))+40:], uint32(i32(-2)))
	store32(m.memory[int64(uint32(v4))+44:], uint32(v1))
	store32(m.memory[int64(uint32(v4))+48:], uint32(v1+v2))
l4:
	{
		t1 := m.fn870(v4 + i32(40))
		v1 = t1
		if v1 == 0 {
			goto l0
		}
		{
			t2 := int32(load32(m.memory[uint32(v1):]))
			v1 = t2
			t3 := m.fn630(v1)
			if t3 != 0 {
				_ = m.fn869(v4 + i32(40))
				goto l4
			}
			if v1 == i32(34) {
				_ = m.fn869(v4 + i32(40))
				store32(m.memory[int64(uint32(v4))+128:], uint32(i32(0)))
				store64(m.memory[int64(uint32(v4))+120:], uint64(i64(0x100000000)))
			l9:
				{
					{
						t12 := m.fn869(v4 + i32(40))
						v1 = t12
						if v1 == i32(92) {
							goto l5
						}
						if v1 == i32(-1) {
							goto l6
						}
						if v1 == i32(34) {
							goto l6
						}
						goto l7
					}
				l5:
					t13 := m.fn869(v4 + i32(40))
					v1 = t13
					if v1 == i32(34) {
						goto l7
					}
					if v1 == i32(92) {
						goto l7
					}
					if v1 != i32(-1) {
						goto l8
					}
				}
			l6:
				m.fn1321(v4+i32(108), v4+i32(120))
				goto l4
			l8:
				m.fn74(v4+i32(120), i32(92))
			l7:
				m.fn74(v4+i32(120), v1)
				goto l9
			}
			if v1 != i32(92) {
				store32(m.memory[int64(uint32(v4))+140:], uint32(i32(0)))
				store64(m.memory[int64(uint32(v4))+132:], uint64(i64(0x100000000)))
			l11:
				{
					t14 := m.fn870(v4 + i32(40))
					v1 = t14
					if v1 == 0 {
						goto l10
					}
					t15 := int32(load32(m.memory[uint32(v1):]))
					v1 = t15
					t16 := m.fn630(v1)
					if t16 != 0 {
						goto l10
					}
					m.fn74(v4+i32(132), v1)
					_ = m.fn869(v4 + i32(40))
					goto l11
				}
			l10:
				m.fn1321(v4+i32(108), v4+i32(132))
				goto l4
			}
			_ = m.fn869(v4 + i32(40))
			t5 := m.fn870(v4 + i32(40))
			v1 = t5
			if v1 == 0 {
				goto l4
			}
			t6 := int32(load32(m.memory[uint32(v1):]))
			v1 = t6
			_ = m.fn869(v4 + i32(40))
			store32(m.memory[int64(uint32(v4))+56:], uint32(i32(-1)))
			t9 := v4
			p8 := v1
			if uint32(v1+i32(-65)) < uint32(i32(26)) {
				p8 = v1 | i32(32)
			}
			store32(m.memory[int64(uint32(t9))+60:], uint32(p8))
			m.fn1321(v4+i32(108), v4+i32(56))
			goto l4
		}
	}
l0:
	t18 := int32(load32(m.memory[int64(uint32(v4))+112:]))
	v5 = t18
	t19 := int32(load32(m.memory[int64(uint32(v4))+116:]))
	t20 := v5
	v2 = t19
	v6 = t20 + v2*i32(12)
	t21 := int32(load32(m.memory[int64(uint32(v4))+108:]))
	v7 = t21
	v1 = v5
	{
		if v2 == 0 {
			goto l12
		}
		v1 = v5 + i32(12)
		t22 := int32(load32(m.memory[uint32(v5):]))
		v2 = t22
		if uint32(v2) < uint32(i32(-2)) {
			goto l13
		}
	}
l12:
	store32(m.memory[int64(uint32(v4))+40:], uint32(i32(-1)))
	goto l14
l13:
	{
		t23 := int32(load32(m.memory[int64(uint32(v5))+4:]))
		v8 = t23
		t24 := int32(load32(m.memory[int64(uint32(v5))+8:]))
		t25 := m.fn1032(v8, t24, i32(1092084), i32(9))
		if t25 == 0 {
			goto l15
		}
		m.fn16(v2, v8)
		v9 = i32(-1)
		v10 = i32(-1)
		v8 = v1
	l26:
		v2 = i32(-3)
	l21:
		{
			{
				if v2 == i32(-3) {
					goto l16
				}
				v11 = v12
				goto l17
			l16:
				if v8 == v6 {
					goto l18
				}
				t26 := int64(load64(m.memory[int64(uint32(v8))+4:]))
				v11 = t26
				t27 := int32(load32(m.memory[uint32(v8):]))
				v2 = t27
				v1 = v8 + i32(12)
				v8 = v1
			}
		l17:
			if v2 == i32(-2) {
				goto l18
			}
			v13 = int32(v11)
			if v2 != i32(-1) {
				if v9 != i32(-1) {
					goto l23
				}
				t28 := v4 + i32(32)
				t29 := v13
				v9 = int32(int64(uint64(v11) >> 32))
				m.fn46(t28, t29, v9)
				{
					t30 := int32(load32(m.memory[int64(uint32(v4))+36:]))
					if t30 != 0 {
						m.fn46(v4+i32(24), v13, v9)
						t31 := int32(load32(m.memory[int64(uint32(v4))+24:]))
						t32 := int32(load32(m.memory[int64(uint32(v4))+28:]))
						m.fn51(v4+i32(56), t31, t32)
						m.fn134(i32(-1), v14)
						t33 := int32(load32(m.memory[int64(uint32(v4))+64:]))
						v15 = t33
						t34 := int32(load32(m.memory[int64(uint32(v4))+60:]))
						v14 = t34
						t35 := int32(load32(m.memory[int64(uint32(v4))+56:]))
						v9 = t35
						goto l23
					}
					v9 = i32(-1)
					goto l23
				}
			}
			v2 = i32(-3)
			switch v13 + i32(-108) {
			case 0, 3:
				goto l20
			case 1, 2:
				goto l21
			default:
				goto l22
			}
		l22:
			if v13 != i32(116) {
				goto l21
			}
		l20:
			v2 = i32(-2)
			if v8 == v6 {
				goto l21
			}
			t36 := int32(load32(m.memory[uint32(v8):]))
			v2 = t36
			t37 := int64(load64(m.memory[int64(uint32(v8))+4:]))
			v12 = t37
			v16 = int32(v12)
			v1 = v8 + i32(12)
			v8 = v1
			if uint32(v2) > uint32(i32(-3)) {
				goto l21
			}
			{
				if v13 != i32(108) {
					goto l25
				}
				t38 := v4 + i32(16)
				t39 := v16
				v8 = int32(int64(uint64(v12) >> 32))
				m.fn46(t38, t39, v8)
				t40 := int32(load32(m.memory[int64(uint32(v4))+20:]))
				if t40 == 0 {
					goto l25
				}
				m.fn46(v4+i32(8), v16, v8)
				t41 := int32(load32(m.memory[int64(uint32(v4))+8:]))
				t42 := int32(load32(m.memory[int64(uint32(v4))+12:]))
				m.fn51(v4+i32(56), t41, t42)
				m.fn134(v10, v17)
				t43 := int32(load32(m.memory[int64(uint32(v4))+64:]))
				v18 = t43
				t44 := int32(load32(m.memory[int64(uint32(v4))+60:]))
				v17 = t44
				t45 := int32(load32(m.memory[int64(uint32(v4))+56:]))
				v10 = t45
				m.fn16(v2, v16)
				v8 = v1
				goto l26
			}
		l25:
			m.fn16(v2, v16)
			v2 = i32(-3)
			v8 = v1
			goto l21
		}
	l18:
		{
			if v9 == i32(-1) {
				if v10 == i32(-1) {
					store32(m.memory[int64(uint32(v4))+40:], uint32(i32(-1)))
					goto l14
				}
				store32(m.memory[int64(uint32(v4))+52:], uint32(v18))
				store32(m.memory[int64(uint32(v4))+48:], uint32(v17))
				store32(m.memory[int64(uint32(v4))+44:], uint32(v10))
				store32(m.memory[int64(uint32(v4))+40:], uint32(i32(2)))
				goto l14
			}
			if v10 == i32(-1) {
				store32(m.memory[int64(uint32(v4))+64:], uint32(v15))
				store32(m.memory[int64(uint32(v4))+60:], uint32(v14))
				store32(m.memory[int64(uint32(v4))+56:], uint32(v9))
				m.fn1454(v4+i32(40), v4+i32(56))
				goto l14
			}
			store32(m.memory[int64(uint32(v4))+92:], uint32(v15))
			store32(m.memory[int64(uint32(v4))+88:], uint32(v14))
			store32(m.memory[int64(uint32(v4))+84:], uint32(v9))
			store32(m.memory[int64(uint32(v4))+104:], uint32(v18))
			store32(m.memory[int64(uint32(v4))+100:], uint32(v17))
			store32(m.memory[int64(uint32(v4))+96:], uint32(v10))
			store32(m.memory[int64(uint32(v4))+68:], uint32(i32(25)))
			store32(m.memory[int64(uint32(v4))+60:], uint32(i32(25)))
			store32(m.memory[int64(uint32(v4))+64:], uint32(v4+i32(96)))
			store32(m.memory[int64(uint32(v4))+56:], uint32(v4+i32(84)))
			m.fn73(v4+i32(108), i32(0x1000d8), v4+i32(56))
			m.fn1454(v4+i32(40), v4+i32(108))
			t46 := int32(load32(m.memory[int64(uint32(v4))+96:]))
			t47 := int32(load32(m.memory[int64(uint32(v4))+100:]))
			m.fn16(t46, t47)
			t48 := int32(load32(m.memory[int64(uint32(v4))+84:]))
			t49 := int32(load32(m.memory[int64(uint32(v4))+88:]))
			m.fn16(t48, t49)
			goto l14
		}
	l23:
		m.fn16(v2, v13)
		goto l26
	}
l15:
	store32(m.memory[int64(uint32(v4))+40:], uint32(i32(-1)))
	m.fn16(v2, v8)
l14:
	t50 := int32(uint32(v6-v1) / uint32(i32(12)))
	v2 = t50
l31:
	{
		if v2 == 0 {
			goto l30
		}
		t51 := int32(load32(m.memory[uint32(v1):]))
		t52 := int32(load32(m.memory[uint32(v1+i32(4)):]))
		m.fn1390(t51, t52)
		v2 = v2 + i32(-1)
		v1 = v1 + i32(12)
		goto l31
	}
l30:
	m.fn136(v7, v5, i32(4), i32(12))
	{
		{
			t53 := int32(load32(m.memory[int64(uint32(v4))+40:]))
			if t53 != i32(-1) {
				goto l32
			}
			t54 := int32(load32(m.memory[int64(uint32(v3))+8:]))
			store32(m.memory[int64(uint32(v0))+8:], uint32(t54))
			t55 := int64(load64(m.memory[uint32(v3):]))
			store64(m.memory[uint32(v0):], uint64(t55))
			goto l33
		}
	l32:
		{
			t56 := int32(load32(m.memory[int64(uint32(v3))+4:]))
			t57 := int32(load32(m.memory[int64(uint32(v3))+8:]))
			t58 := m.fn23(t56, t57)
			if t58 != 0 {
				goto l34
			}
			t59 := int64(load64(m.memory[int64(uint32(v4))+48:]))
			store64(m.memory[int64(uint32(v4))+64:], uint64(t59))
			t60 := int64(load64(m.memory[int64(uint32(v4))+40:]))
			store64(m.memory[int64(uint32(v4))+56:], uint64(t60))
			t61 := m.fn113(i32(4), i32(28))
			v1 = t61
			t62 := int64(load64(m.memory[int64(uint32(v4))+56:]))
			store64(m.memory[uint32(v1):], uint64(t62))
			t63 := int64(load64(m.memory[int64(uint32(v4))+64:]))
			store64(m.memory[int64(uint32(v1))+8:], uint64(t63))
			t64 := int64(load64(m.memory[uint32(v3):]))
			store64(m.memory[int64(uint32(v1))+16:], uint64(t64))
			t65 := int32(load32(m.memory[int64(uint32(v3))+8:]))
			store32(m.memory[int64(uint32(v1))+24:], uint32(t65))
			store32(m.memory[int64(uint32(v0))+8:], uint32(i32(1)))
			store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
			store32(m.memory[uint32(v0):], uint32(i32(1)))
			goto l33
		}
	l34:
		t66 := int32(load32(m.memory[int64(uint32(v3))+8:]))
		store32(m.memory[int64(uint32(v0))+8:], uint32(t66))
		t67 := int64(load64(m.memory[uint32(v3):]))
		store64(m.memory[uint32(v0):], uint64(t67))
		t68 := int32(load32(m.memory[int64(uint32(v4))+44:]))
		t69 := int32(load32(m.memory[int64(uint32(v4))+48:]))
		m.fn16(t68, t69)
	}
l33:
	m.g0 = v4 + i32(144)
}
func (m *Module) fn1404(v0, v1, v2 int32) {
	var v3, v4, v5, v6 int32
	t0 := m.g0
	v3 = t0 - i32(144)
	m.g0 = v3
	store16(m.memory[int64(uint32(v3))+64:], uint16(i32(514)))
	store64(m.memory[int64(uint32(v3))+56:], uint64(i64(144680345676153346)))
	m.fn513(v3+i32(68), v1, v2, i32(59))
l1:
	{
		m.fn515(v3+i32(48), v3+i32(68))
		{
			t1 := int32(load32(m.memory[int64(uint32(v3))+48:]))
			v2 = t1
			if v2 == 0 {
				t58 := int32(load16(m.memory[int64(uint32(v3))+64:]))
				store16(m.memory[int64(uint32(v0))+8:], uint16(t58))
				t59 := int64(load64(m.memory[int64(uint32(v3))+56:]))
				store64(m.memory[uint32(v0):], uint64(t59))
				m.g0 = v3 + i32(144)
				return
			}
			t2 := int32(load32(m.memory[int64(uint32(v3))+52:]))
			m.fn775(v3+i32(108), v2, t2, i32(58))
			t3 := int32(load32(m.memory[int64(uint32(v3))+108:]))
			v2 = t3
			if v2 == 0 {
				goto l1
			}
			t4 := int32(load32(m.memory[int64(uint32(v3))+120:]))
			v1 = t4
			t5 := int32(load32(m.memory[int64(uint32(v3))+116:]))
			v4 = t5
			t6 := int32(load32(m.memory[int64(uint32(v3))+112:]))
			m.fn46(v3+i32(40), v2, t6)
			t7 := int32(load32(m.memory[int64(uint32(v3))+40:]))
			t8 := int32(load32(m.memory[int64(uint32(v3))+44:]))
			m.fn14(v3+i32(124), t7, t8)
			m.fn46(v3+i32(32), v4, v1)
			t9 := int32(load32(m.memory[int64(uint32(v3))+32:]))
			t10 := int32(load32(m.memory[int64(uint32(v3))+36:]))
			m.fn14(v3+i32(108), t9, t10)
			t11 := int32(load32(m.memory[int64(uint32(v3))+112:]))
			t12 := v3 + i32(24)
			v2 = t11
			t13 := int32(load32(m.memory[int64(uint32(v3))+116:]))
			t14 := v2
			v1 = t13
			m.fn1056(t12, t14, v1, i32(33))
			v4 = v3 + i32(56)
			{
				t15 := int32(load32(m.memory[int64(uint32(v3))+24:]))
				if t15&i32(1) == 0 {
					goto l2
				}
				t16 := int32(load32(m.memory[int64(uint32(v3))+28:]))
				t17 := v3 + i32(16)
				v4 = t16
				m.fn826(t17, v4+i32(1), v2, v1, i32(1081020))
				t18 := int32(load32(m.memory[int64(uint32(v3))+16:]))
				t19 := int32(load32(m.memory[int64(uint32(v3))+20:]))
				m.fn46(v3+i32(8), t18, t19)
				t20 := int32(load32(m.memory[int64(uint32(v3))+8:]))
				t21 := int32(load32(m.memory[int64(uint32(v3))+12:]))
				t22 := m.fn15(t20, t21, i32(1081036), i32(9))
				v2 = t22
				m.fn852(v3+i32(108), v4, i32(1081048))
				t23 := int32(load32(m.memory[int64(uint32(v3))+112:]))
				t24 := int32(load32(m.memory[int64(uint32(v3))+116:]))
				m.fn628(v3, t23, t24)
				t25 := int32(load32(m.memory[int64(uint32(v3))+4:]))
				m.fn852(v3+i32(108), t25, i32(1081064))
				t27 := v3 + i32(56)
				p26 := i32(0)
				if v2 != 0 {
					p26 = i32(5)
				}
				v4 = t27 | p26
			}
		l2:
			{
				t28 := int32(load32(m.memory[int64(uint32(v3))+128:]))
				v2 = t28
				t29 := int32(load32(m.memory[int64(uint32(v3))+132:]))
				t30 := v2
				v5 = t29
				t31 := m.fn15(t30, v5, i32(1081080), i32(11))
				if t31 != 0 {
					v5 = i32(1)
					{
						t38 := int32(load32(m.memory[int64(uint32(v3))+112:]))
						v1 = t38
						t39 := int32(load32(m.memory[int64(uint32(v3))+116:]))
						t40 := v1
						v6 = t39
						t41 := m.fn773(t40, v6, i32(1074847), i32(4))
						if t41 != 0 {
							goto l7
						}
						t42 := m.fn773(v1, v6, i32(1081155), i32(6))
						if t42 != 0 {
							goto l7
						}
						m.fn1071(v3+i32(136), v1, v6)
						t43 := int32(m.memory[int64(uint32(v3))+136])
						t44 := int32(load32(m.memory[int64(uint32(v3))+140:]))
						t45 := t43 ^ i32(-1)
						var p46 int32
						if uint32(t44) > uint32(i32(599)) {
							p46 = 1
						}
						v5 = t45 & p46
					}
				l7:
					m.memory[uint32(v4)] = byte(v5)
					goto l6
				}
				t32 := m.fn15(v2, v5, i32(1081091), i32(10))
				if t32 == 0 {
					{
						{
							t47 := m.fn15(v2, v5, i32(1081101), i32(15))
							if t47 == 0 {
								goto l8
							}
							t48 := int32(load32(m.memory[int64(uint32(v3))+112:]))
							v1 = t48
							goto l9
						}
					l8:
						t49 := int32(load32(m.memory[int64(uint32(v3))+112:]))
						v1 = t49
						t50 := m.fn15(v2, v5, i32(1081116), i32(20))
						if t50 == 0 {
							t55 := m.fn15(v2, v5, i32(1081136), i32(7))
							if t55 == 0 {
								goto l6
							}
							t56 := int32(load32(m.memory[int64(uint32(v3))+116:]))
							t57 := m.fn773(v1, t56, i32(1074851), i32(4))
							m.memory[int64(uint32(v4))+4] = byte(t57)
							goto l6
						}
					}
				l9:
					t51 := int32(load32(m.memory[int64(uint32(v3))+116:]))
					t52 := v1
					v5 = t51
					t53 := m.fn789(i32(1081143), i32(12), t52, v5)
					if t53 != 0 {
						m.memory[int64(uint32(v4))+2] = byte(i32(1))
						goto l6
					}
					t54 := m.fn773(v1, v5, i32(1074851), i32(4))
					if t54 == 0 {
						goto l6
					}
					m.memory[int64(uint32(v4))+2] = byte(i32(0))
					goto l6
				}
				v5 = i32(1)
				{
					t33 := int32(load32(m.memory[int64(uint32(v3))+112:]))
					v1 = t33
					t34 := int32(load32(m.memory[int64(uint32(v3))+116:]))
					t35 := v1
					v6 = t34
					t36 := m.fn773(t35, v6, i32(1074855), i32(6))
					if t36 != 0 {
						goto l5
					}
					t37 := m.fn773(v1, v6, i32(1074861), i32(7))
					v5 = t37
				}
			l5:
				m.memory[int64(uint32(v4))+1] = byte(v5)
				goto l6
			}
		}
	l6:
		t60 := int32(load32(m.memory[int64(uint32(v3))+108:]))
		m.fn16(t60, v1)
		t61 := int32(load32(m.memory[int64(uint32(v3))+124:]))
		m.fn16(t61, v2)
		goto l1
	}
}
func (m *Module) fn1405(v0 int32) int32 {
	var v1 int32
	v1 = i32(0)
	{
		t0 := int32(m.memory[uint32(v0)])
		if t0 != i32(2) {
			goto l0
		}
		t1 := int32(m.memory[int64(uint32(v0))+1])
		if t1&i32(255) != i32(2) {
			goto l0
		}
		t2 := int32(m.memory[int64(uint32(v0))+2])
		if t2&i32(255) != i32(2) {
			goto l0
		}
		t3 := int32(m.memory[int64(uint32(v0))+3])
		if t3&i32(255) != i32(2) {
			goto l0
		}
		t4 := int32(m.memory[int64(uint32(v0))+4])
		var p5 int32
		if t4 == i32(2) {
			p5 = 1
		}
		v1 = p5
	}
l0:
	return v1
}
func (m *Module) fn1406(v0, v1, v2 int32) {
	var v3, v4, v5 int32
	t0 := m.g0
	v3 = t0 - i32(32)
	m.g0 = v3
	m.fn140(v3+i32(12), v2)
	store32(m.memory[int64(uint32(v3))+24:], uint32(v1))
	store32(m.memory[int64(uint32(v3))+28:], uint32(v1+v2))
	v2 = i32(0)
l1:
	{
		t1 := m.fn48(v3 + i32(24))
		v1 = t1
		if v1 == i32(-1) {
			goto l0
		}
		t2 := m.fn630(v1)
		v4 = t2
		v5 = v4 & v2
		v2 = i32(1)
		if v5 != 0 {
			goto l1
		}
		t4 := v3 + i32(12)
		p3 := v1
		if v4 != 0 {
			p3 = i32(32)
		}
		m.fn74(t4, p3)
		v2 = v4
		goto l1
	}
l0:
	t5 := int32(load32(m.memory[int64(uint32(v3))+20:]))
	store32(m.memory[int64(uint32(v0))+8:], uint32(t5))
	t6 := int64(load64(m.memory[int64(uint32(v3))+12:]))
	store64(m.memory[uint32(v0):], uint64(t6))
	m.g0 = v3 + i32(32)
}
func (m *Module) fn1407(v0, v1, v2 int32) int32 {
	var v3, v4, v5, v6 int32
	t0 := m.g0
	v3 = t0 - i32(32)
	m.g0 = v3
l6:
	v4 = v1 * i32(28)
l7:
	if v4 != 0 {
		v5 = i32(0)
		{
			v6 = v0 + v4
			t1 := int32(load32(m.memory[uint32(v6+i32(-28)):]))
			v1 = t1
			p2 := i32(1)
			if uint32(v1) > uint32(i32(2)) {
				p2 = v1 + i32(-3)
			}
			switch p2 {
			case 2, 4:
				goto l1
			case 3:
				goto l4
			case 5:
				v5 = i32(1)
				goto l1
			default:
				t3 := int32(load32(m.memory[uint32(v6+i32(-16)):]))
				v1 = t3
				if v1 == 0 {
					goto l4
				}
				t4 := int32(load32(m.memory[uint32(v6+i32(-20)):]))
				v4 = t4
				store32(m.memory[int64(uint32(v3))+28:], uint32(i32(0)))
				store32(m.memory[int64(uint32(v3))+20:], uint32(v4))
				store32(m.memory[int64(uint32(v3))+12:], uint32(v4))
				store32(m.memory[int64(uint32(v3))+16:], uint32(v1))
				store32(m.memory[int64(uint32(v3))+24:], uint32(v4+v1))
				m.fn629(v3, v3+i32(12))
				t5 := int32(load32(m.memory[uint32(v3):]))
				t6 := int32(load32(m.memory[int64(uint32(v3))+8:]))
				var p7 int32
				if t5 == 0 {
					p7 = 1
				}
				var p8 int32
				if v1 == t6 {
					p8 = 1
				}
				v5 = p7 & p8
				goto l1
			case 1:
				t9 := int32(load32(m.memory[uint32(v6+i32(-8)):]))
				v5 = t9
				t10 := int32(load32(m.memory[uint32(v6+i32(-4)):]))
				t11 := v5
				v1 = t10
				t12 := m.fn23(t11, v1)
				if t12 != 0 {
					goto l4
				}
				v2 = i32(0)
				v0 = v5
				goto l6
			}
		}
	l4:
		v4 = v4 + i32(-28)
		goto l7
	}
	v5 = v2
	goto l1
l1:
	m.g0 = v3 + i32(32)
	return v5 & i32(1)
}
func (m *Module) fn1408(v0, v1 int32) {
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
func (m *Module) fn1409(v0 int32) {
	var v1, v2 int32
	var v3 int64
	t0 := m.g0
	v1 = t0 - i32(48)
	m.g0 = v1
	{
		t1 := int32(load32(m.memory[int64(uint32(v0))+20:]))
		if t1 == 0 {
			goto l0
		}
		t2 := int32(load32(m.memory[int64(uint32(v0))+20:]))
		v2 = t2
		store32(m.memory[int64(uint32(v0))+20:], uint32(i32(0)))
		t3 := int64(load64(m.memory[int64(uint32(v0))+12:]))
		v3 = t3
		store64(m.memory[int64(uint32(v0))+12:], uint64(i64(0x400000000)))
		store64(m.memory[uint32(v1):], uint64(v3))
		store32(m.memory[int64(uint32(v1))+8:], uint32(v2))
		{
			t4 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t5 := m.fn1421(t4, v2)
			if t5 != 0 {
				goto l1
			}
			m.fn894(v1)
			goto l0
		}
	l1:
		t6 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		store32(m.memory[int64(uint32(v1))+28:], uint32(t6))
		t7 := int64(load64(m.memory[uint32(v1):]))
		store64(m.memory[int64(uint32(v1))+20:], uint64(t7))
		store32(m.memory[int64(uint32(v1))+16:], uint32(i32(-0x80000000)))
		m.fn338(v0, v1+i32(16))
	}
l0:
	m.memory[int64(uint32(v0))+36] = byte(i32(1))
	m.g0 = v1 + i32(48)
}
func (m *Module) fn1410(v0, v1 int32) {
	var v2 int32
	var v3 int64
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	m.fn59(v2+i32(8), v1, i32(4), i32(28))
	t1 := int64(load64(m.memory[int64(uint32(v2))+8:]))
	v3 = t1
	store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
	store64(m.memory[uint32(v0):], uint64(v3))
	m.g0 = v2 + i32(16)
}
func (m *Module) fn1411(v0, v1, v2, v3 int32) {
	m.fn1469(v0, v1, v2, v3, i32(1))
}
func (m *Module) fn1412(v0, v1 int32) {
	var v2 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		v2 = t0
		t1 := int32(load32(m.memory[uint32(v0):]))
		if v2 != t1 {
			goto l0
		}
		m.fn1143(v0)
	}
l0:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v2+i32(1)))
	t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	v0 = t2 + v2*i32(28)
	t3 := int64(load64(m.memory[uint32(v1):]))
	store64(m.memory[uint32(v0):], uint64(t3))
	t4 := int64(load64(m.memory[int64(uint32(v1))+8:]))
	store64(m.memory[int64(uint32(v0))+8:], uint64(t4))
	t5 := int64(load64(m.memory[int64(uint32(v1))+16:]))
	store64(m.memory[int64(uint32(v0))+16:], uint64(t5))
	t6 := int32(load32(m.memory[int64(uint32(v1))+24:]))
	store32(m.memory[int64(uint32(v0))+24:], uint32(t6))
}
func (m *Module) fn1413(v0 int32, v1 int64) {
	var v2 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		v2 = t0
		t1 := int32(load32(m.memory[uint32(v0):]))
		if v2 != t1 {
			goto l0
		}
		m.fn1154(v0)
	}
l0:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v2+i32(1)))
	t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	store64(m.memory[uint32(t2+v2<<3):], uint64(v1))
}
func (m *Module) fn1414(v0 int32) {
	t0 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	if t0 == i32(-1) {
		return
	}
	m.fn971(v0 + i32(8))
}
func (m *Module) fn1415(v0, v1 int32) {
	m.fn136(v0, v1, i32(8), i32(8))
}
func (m *Module) fn1416(v0, v1 int32) int32 {
	var v2 int32
	v2 = i32(1)
	{
		t0 := m.fn15(v0, v1, i32(1080916), i32(3))
		if t0 != 0 {
			goto l0
		}
		t1 := m.fn15(v0, v1, i32(1080919), i32(7))
		if t1 != 0 {
			goto l0
		}
		t2 := m.fn15(v0, v1, i32(1080926), i32(7))
		if t2 != 0 {
			goto l0
		}
		t3 := m.fn15(v0, v1, i32(1080933), i32(5))
		if t3 != 0 {
			goto l0
		}
		t4 := m.fn15(v0, v1, i32(1080938), i32(4))
		if t4 != 0 {
			goto l0
		}
		t5 := m.fn15(v0, v1, i32(1080942), i32(3))
		if t5 != 0 {
			goto l0
		}
		t6 := m.fn15(v0, v1, i32(1078275), i32(6))
		if t6 != 0 {
			goto l0
		}
		t7 := m.fn15(v0, v1, i32(1078281), i32(6))
		if t7 != 0 {
			goto l0
		}
		t8 := m.fn15(v0, v1, i32(1080945), i32(6))
		if t8 != 0 {
			goto l0
		}
		t9 := m.fn15(v0, v1, i32(1080951), i32(10))
		if t9 != 0 {
			goto l0
		}
		t10 := m.fn15(v0, v1, i32(1080961), i32(6))
		if t10 != 0 {
			goto l0
		}
		t11 := m.fn15(v0, v1, i32(1080967), i32(7))
		if t11 != 0 {
			goto l0
		}
		t12 := m.fn15(v0, v1, i32(1080974), i32(7))
		if t12 != 0 {
			goto l0
		}
		t13 := m.fn15(v0, v1, i32(1073486), i32(2))
		if t13 != 0 {
			goto l0
		}
		t14 := m.fn15(v0, v1, i32(1080981), i32(2))
		if t14 != 0 {
			goto l0
		}
		t15 := m.fn15(v0, v1, i32(1080983), i32(2))
		if t15 != 0 {
			goto l0
		}
		t16 := m.fn15(v0, v1, i32(1080985), i32(2))
		if t16 != 0 {
			goto l0
		}
		t17 := m.fn15(v0, v1, i32(1073232), i32(4))
		v2 = t17
	}
l0:
	return v2
}
func (m *Module) fn1417(v0, v1 int32) {
	var v2, v3, v4, v5, v6 int32
	t0 := m.g0
	v2 = t0 - i32(48)
	m.g0 = v2
	t1 := int32(load32(m.memory[int64(uint32(v1))+16:]))
	t2 := v2 + i32(8)
	v3 = t1
	t3 := int32(load32(m.memory[int64(uint32(v1))+20:]))
	t4 := v3
	v4 = t3
	m.fn909(t2, t4, v4, i32(1073226), i32(2))
	{
		t5 := int32(load32(m.memory[int64(uint32(v2))+8:]))
		v5 = t5
		if v5 == 0 {
			goto l0
		}
		t6 := int32(load32(m.memory[int64(uint32(v2))+12:]))
		v6 = t6
		if v6 == 0 {
			goto l0
		}
		t7 := int32(load32(m.memory[int64(uint32(v0))+28:]))
		t8 := int32(load32(m.memory[int64(uint32(v0))+32:]))
		t9 := int32(load32(m.memory[int64(uint32(t8))+20:]))
		m.t0[uint(t9)].(func(int32, int32, int32, int32))(v2+i32(24), t7, v5, v6)
		store32(m.memory[int64(uint32(v2))+20:], uint32(i32(6)))
		m.fn1340(v0+i32(12), v2+i32(20))
	}
l0:
	{
		t10 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t11 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		t12 := m.fn773(t10, t11, i32(1077050), i32(1))
		if t12 == 0 {
			goto l1
		}
		m.fn909(v2, v3, v4, i32(1073713), i32(4))
		t13 := int32(load32(m.memory[uint32(v2):]))
		v1 = t13
		if v1 == 0 {
			goto l1
		}
		t14 := int32(load32(m.memory[int64(uint32(v2))+4:]))
		v3 = t14
		if v3 == 0 {
			goto l1
		}
		t15 := int32(load32(m.memory[int64(uint32(v0))+28:]))
		t16 := int32(load32(m.memory[int64(uint32(v0))+32:]))
		t17 := int32(load32(m.memory[int64(uint32(t16))+20:]))
		m.t0[uint(t17)].(func(int32, int32, int32, int32))(v2+i32(24), t15, v1, v3)
		store32(m.memory[int64(uint32(v2))+20:], uint32(i32(6)))
		m.fn1340(v0+i32(12), v2+i32(20))
	}
l1:
	m.g0 = v2 + i32(48)
}
func (m *Module) fn1418(v0, v1, v2, v3, v4 int32) {
	var v5 int32
	var v6 int64
	var v7, v8, v9, v10, v11 int32
	t0 := m.g0
	v5 = t0 - i32(144)
	m.g0 = v5
	m.fn1469(v5+i32(104), v1, v2, v3, v4)
	t1 := int64(load64(m.memory[int64(uint32(v5))+108:]))
	store64(m.memory[int64(uint32(v5))+32:], uint64(t1))
	t2 := int32(load32(m.memory[int64(uint32(v5))+116:]))
	store32(m.memory[int64(uint32(v5))+40:], uint32(t2))
	{
		t3 := int32(load32(m.memory[int64(uint32(v5))+104:]))
		v4 = t3
		if v4 == i32(-1) {
			t7 := int64(load64(m.memory[int64(uint32(v5))+32:]))
			store64(m.memory[int64(uint32(v5))+8:], uint64(t7))
			t8 := int32(load32(m.memory[int64(uint32(v5))+40:]))
			t9 := v5
			v7 = t8
			store32(m.memory[int64(uint32(t9))+16:], uint32(v7))
			t10 := int32(load32(m.memory[int64(uint32(v5))+12:]))
			v4 = t10
			{
				if v7 != i32(1) {
					goto l2
				}
				t11 := int32(load32(m.memory[uint32(v4):]))
				if t11 != i32(-0x80000000) {
					goto l2
				}
				store32(m.memory[uint32(v0):], uint32(i32(-1)))
				t12 := int32(load32(m.memory[int64(uint32(v4))+12:]))
				store32(m.memory[int64(uint32(v0))+12:], uint32(t12))
				t13 := int64(load64(m.memory[int64(uint32(v4))+4:]))
				store64(m.memory[int64(uint32(v0))+4:], uint64(t13))
				store32(m.memory[int64(uint32(v4))+12:], uint32(i32(0)))
				store64(m.memory[int64(uint32(v4))+4:], uint64(i64(0x400000000)))
				m.fn969(v5 + i32(8))
				goto l1
			}
		l2:
			store32(m.memory[int64(uint32(v5))+28:], uint32(i32(0)))
			store64(m.memory[int64(uint32(v5))+20:], uint64(i64(0x400000000)))
			t14 := int32(load32(m.memory[int64(uint32(v5))+8:]))
			store32(m.memory[int64(uint32(v5))+40:], uint32(t14))
			store32(m.memory[int64(uint32(v5))+32:], uint32(v4))
			t15 := v5
			t16 := v4
			v1 = v7 << 5
			v8 = t16 + v1
			store32(m.memory[int64(uint32(t15))+44:], uint32(v8))
			v9 = v4 + i32(32)
			v10 = v5 + i32(104) + i32(4)
			v3 = v5 + i32(56) | i32(4)
			v11 = i32(0)
		l10:
			{
				{
					if v1 == 0 {
						goto l3
					}
					t17 := int32(load32(m.memory[uint32(v4):]))
					v2 = t17
					if v2 != i32(-1) {
						t20 := int32(load32(m.memory[int64(uint32(v4))+28:]))
						store32(m.memory[int64(uint32(v3))+24:], uint32(t20))
						t21 := int64(load64(m.memory[int64(uint32(v4))+20:]))
						store64(m.memory[int64(uint32(v3))+16:], uint64(t21))
						t22 := int64(load64(m.memory[int64(uint32(v4))+12:]))
						store64(m.memory[int64(uint32(v3))+8:], uint64(t22))
						t23 := int64(load64(m.memory[int64(uint32(v4))+4:]))
						store64(m.memory[uint32(v3):], uint64(t23))
						store32(m.memory[int64(uint32(v5))+56:], uint32(v2))
						if v11 == 0 {
							goto l5
						}
						store32(m.memory[int64(uint32(v5))+104:], uint32(i32(8)))
						m.fn1340(v5+i32(20), v5+i32(104))
					l5:
						switch v2 >> 31 & (v2 + i32(-0x7fffffff)) {
						default:
							m.fn835(v5+i32(132), v5+i32(56))
							t24 := int32(load32(m.memory[int64(uint32(v5))+136:]))
							t25 := v10
							v2 = t24
							t26 := int32(load32(m.memory[int64(uint32(v5))+140:]))
							m.fn1406(t25, v2, t26)
							store32(m.memory[int64(uint32(v5))+104:], uint32(i32(3)))
							store32(m.memory[int64(uint32(v5))+120:], uint32(i32(0)))
							m.fn1340(v5+i32(20), v5+i32(104))
							t27 := int32(load32(m.memory[int64(uint32(v5))+132:]))
							m.fn16(t27, v2)
							m.fn970(v5 + i32(56))
							goto l9
						case 1:
							t28 := int32(load32(m.memory[int64(uint32(v3))+8:]))
							store32(m.memory[int64(uint32(v5))+96:], uint32(t28))
							t29 := int64(load64(m.memory[uint32(v3):]))
							store64(m.memory[int64(uint32(v5))+88:], uint64(t29))
							m.fn1341(v5+i32(20), v5+i32(88))
							goto l9
						case 0:
							t30 := int32(load32(m.memory[int64(uint32(v5))+64:]))
							store32(m.memory[int64(uint32(v5))+96:], uint32(t30))
							t31 := int64(load64(m.memory[int64(uint32(v5))+56:]))
							store64(m.memory[int64(uint32(v5))+88:], uint64(t31))
							m.fn1341(v5+i32(20), v5+i32(88))
							t32 := int32(load32(m.memory[int64(uint32(v5))+68:]))
							t33 := int32(load32(m.memory[int64(uint32(v5))+72:]))
							m.fn134(t32, t33)
						}
					l9:
						v4 = v4 + i32(32)
						v11 = v11 + i32(1)
						v1 = v1 + i32(-32)
						v9 = v9 + i32(32)
						goto l10
					}
					v7 = v11
					v8 = v9
				}
			l3:
				store32(m.memory[int64(uint32(v5))+48:], uint32(v7))
				store32(m.memory[int64(uint32(v5))+36:], uint32(v8))
				m.fn1339(v5 + i32(32))
				t18 := int32(load32(m.memory[int64(uint32(v5))+28:]))
				store32(m.memory[int64(uint32(v0))+12:], uint32(t18))
				t19 := int64(load64(m.memory[int64(uint32(v5))+20:]))
				store64(m.memory[int64(uint32(v0))+4:], uint64(t19))
				store32(m.memory[uint32(v0):], uint32(i32(-1)))
				goto l1
			}
		}
		t4 := int64(load64(m.memory[int64(uint32(v5))+120:]))
		v6 = t4
		t5 := int32(load32(m.memory[int64(uint32(v5))+40:]))
		store32(m.memory[int64(uint32(v0))+12:], uint32(t5))
		t6 := int64(load64(m.memory[int64(uint32(v5))+32:]))
		store64(m.memory[int64(uint32(v0))+4:], uint64(t6))
		store64(m.memory[int64(uint32(v0))+16:], uint64(v6))
		store32(m.memory[uint32(v0):], uint32(v4))
		goto l1
	}
l1:
	m.g0 = v5 + i32(144)
}
func (m *Module) fn1419(v0 int32) {
	t0 := int32(load32(m.memory[uint32(v0):]))
	if t0 == i32(-1) {
		return
	}
	t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t2 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	m.fn16(t1, t2)
}
func (m *Module) fn1420(v0, v1, v2, v3 int32) {
	m.fn1418(v0, v1, v2, v3, i32(1))
}
func (m *Module) fn1421(v0, v1 int32) int32 {
	var v2 int32
	{
		t0 := m.fn23(v0, v1)
		if t0 == 0 {
			goto l0
		}
		v1 = v1 * i32(28)
	l2:
		{
			if v1 != 0 {
				goto l1
			}
			return i32(0)
		l1:
			v1 = v1 + i32(-28)
			t1 := int32(load32(m.memory[uint32(v0):]))
			v2 = t1
			v0 = v0 + i32(28)
			if v2 != i32(6) {
				goto l2
			}
		}
	}
l0:
	return i32(1)
}
func (m *Module) fn1422(v0, v1 int32) {
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
func (m *Module) fn1423(v0, v1, v2 int32) {
	var v3 int32
	t0 := m.g0
	v3 = t0 - i32(16)
	m.g0 = v3
	m.fn788(v3+i32(8), v1, v2, i32(10))
	t1 := int32(m.memory[int64(uint32(v3))+9])
	v2 = t1
	t2 := int32(m.memory[int64(uint32(v3))+8])
	m.memory[uint32(v0)] = byte(t2 & i32(1))
	m.memory[int64(uint32(v0))+1] = byte(v2)
	m.g0 = v3 + i32(16)
}
func (m *Module) fn1424(v0, v1 int32) {
	var v2, v3, v4 int32
	v2 = i32(0)
	{
		t0 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		v3 = t0
		if v3 == 0 {
			goto l1
		}
		store32(m.memory[int64(uint32(v1))+12:], uint32(v3+i32(-1)))
		t1 := int32(load32(m.memory[uint32(v1):]))
		v4 = t1
		t2 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		if v4 == t2 {
			goto l1
		}
		store32(m.memory[uint32(v1):], uint32(v4+i32(12)))
		t3 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		t4 := v1
		v3 = t3
		store32(m.memory[int64(uint32(t4))+8:], uint32(v3+i32(1)))
		v2 = v4
		goto l1
	}
l1:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v2))
	store32(m.memory[uint32(v0):], uint32(v3))
}
func (m *Module) fn1425(v0, v1, v2 int32) {
	var v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14, v15, v16, v17, v18, v19, v20, v21, v22, v23, v24, v25 int32
	t0 := m.g0
	v3 = t0 - i32(32)
	m.g0 = v3
	m.fn46(v3, v1, v2)
	{
		t1 := int32(load32(m.memory[int64(uint32(v3))+4:]))
		v4 = t1
		if v4 <= i32(-1) {
			m.fn86()
			panic("unreachable")
		}
		{
			{
				if v4 != 0 {
					goto l1
				}
				v5 = i32(0)
				v6 = i32(1)
				goto l2
			l1:
				t2 := int32(load32(m.memory[uint32(v3):]))
				v7 = t2
				t3 := m.fn4(v4)
				v6 = t3
				if v6 == 0 {
					m.fn2(i32(1), v4)
					panic("unreachable")
				}
				v1 = v6
				v5 = i32(0)
				v8 = v4
				v2 = v7
				if uint32(v4) < uint32(i32(16)) {
					goto l4
				}
				v5 = v4 & i32(0x7ffffff0)
				v9 = i32(0)
				v8 = v4
			l6:
				{
					v1 = v6 + v9
					{
						v2 = v7 + v9
						t4 := int32(int8(m.memory[uint32(v2+i32(1))]))
						v10 = t4
						t5 := int32(int8(m.memory[uint32(v2)]))
						t6 := int32(uint32((v10^i32(-1))&i32(128)) >> 7)
						v11 = t5
						t7 := int32(int8(m.memory[uint32(v2+i32(2))]))
						t8 := t6 + int32(uint32((v11^i32(-1))&i32(128))>>7)
						v12 = t7
						t9 := int32(int8(m.memory[uint32(v2+i32(3))]))
						t10 := t8 + int32(uint32((v12^i32(-1))&i32(128))>>7)
						v13 = t9
						t11 := int32(int8(m.memory[uint32(v2+i32(4))]))
						t12 := t10 + int32(uint32((v13^i32(-1))&i32(128))>>7)
						v14 = t11
						t13 := int32(int8(m.memory[uint32(v2+i32(5))]))
						t14 := t12 + int32(uint32((v14^i32(-1))&i32(128))>>7)
						v15 = t13
						t15 := int32(int8(m.memory[uint32(v2+i32(6))]))
						t16 := t14 + int32(uint32((v15^i32(-1))&i32(128))>>7)
						v16 = t15
						t17 := int32(int8(m.memory[uint32(v2+i32(7))]))
						t18 := t16 + int32(uint32((v16^i32(-1))&i32(128))>>7)
						v17 = t17
						t19 := int32(int8(m.memory[uint32(v2+i32(8))]))
						t20 := t18 + int32(uint32((v17^i32(-1))&i32(128))>>7)
						v18 = t19
						t21 := int32(int8(m.memory[uint32(v2+i32(9))]))
						t22 := t20 + int32(uint32((v18^i32(-1))&i32(128))>>7)
						v19 = t21
						t23 := int32(int8(m.memory[uint32(v2+i32(10))]))
						t24 := t22 + int32(uint32((v19^i32(-1))&i32(128))>>7)
						v20 = t23
						t25 := int32(int8(m.memory[uint32(v2+i32(11))]))
						t26 := t24 + int32(uint32((v20^i32(-1))&i32(128))>>7)
						v21 = t25
						t27 := int32(int8(m.memory[uint32(v2+i32(12))]))
						t28 := t26 + int32(uint32((v21^i32(-1))&i32(128))>>7)
						v22 = t27
						t29 := int32(int8(m.memory[uint32(v2+i32(13))]))
						t30 := t28 + int32(uint32((v22^i32(-1))&i32(128))>>7)
						v23 = t29
						t31 := int32(int8(m.memory[uint32(v2+i32(14))]))
						t32 := t30 + int32(uint32((v23^i32(-1))&i32(128))>>7)
						v24 = t31
						t33 := int32(int8(m.memory[uint32(v2+i32(15))]))
						t34 := t32 + int32(uint32((v24^i32(-1))&i32(128))>>7)
						v25 = t33
						if (t34+int32(uint32((v25^i32(-1))&i32(128))>>7))&i32(255) == i32(16) {
							goto l5
						}
						v5 = v9
						goto l4
					}
				l5:
					t36 := v1 + i32(15)
					p35 := i32(0)
					if uint32((v25+i32(-65))&i32(255)) < uint32(i32(26)) {
						p35 = i32(32)
					}
					m.memory[uint32(t36)] = byte(p35 | v25)
					t38 := v1 + i32(14)
					p37 := i32(0)
					if uint32((v24+i32(-65))&i32(255)) < uint32(i32(26)) {
						p37 = i32(32)
					}
					m.memory[uint32(t38)] = byte(p37 | v24)
					t40 := v1 + i32(13)
					p39 := i32(0)
					if uint32((v23+i32(-65))&i32(255)) < uint32(i32(26)) {
						p39 = i32(32)
					}
					m.memory[uint32(t40)] = byte(p39 | v23)
					t42 := v1 + i32(12)
					p41 := i32(0)
					if uint32((v22+i32(-65))&i32(255)) < uint32(i32(26)) {
						p41 = i32(32)
					}
					m.memory[uint32(t42)] = byte(p41 | v22)
					t44 := v1 + i32(11)
					p43 := i32(0)
					if uint32((v21+i32(-65))&i32(255)) < uint32(i32(26)) {
						p43 = i32(32)
					}
					m.memory[uint32(t44)] = byte(p43 | v21)
					t46 := v1 + i32(10)
					p45 := i32(0)
					if uint32((v20+i32(-65))&i32(255)) < uint32(i32(26)) {
						p45 = i32(32)
					}
					m.memory[uint32(t46)] = byte(p45 | v20)
					t48 := v1 + i32(9)
					p47 := i32(0)
					if uint32((v19+i32(-65))&i32(255)) < uint32(i32(26)) {
						p47 = i32(32)
					}
					m.memory[uint32(t48)] = byte(p47 | v19)
					t50 := v1 + i32(8)
					p49 := i32(0)
					if uint32((v18+i32(-65))&i32(255)) < uint32(i32(26)) {
						p49 = i32(32)
					}
					m.memory[uint32(t50)] = byte(p49 | v18)
					t52 := v1 + i32(7)
					p51 := i32(0)
					if uint32((v17+i32(-65))&i32(255)) < uint32(i32(26)) {
						p51 = i32(32)
					}
					m.memory[uint32(t52)] = byte(p51 | v17)
					t54 := v1 + i32(6)
					p53 := i32(0)
					if uint32((v16+i32(-65))&i32(255)) < uint32(i32(26)) {
						p53 = i32(32)
					}
					m.memory[uint32(t54)] = byte(p53 | v16)
					t56 := v1 + i32(5)
					p55 := i32(0)
					if uint32((v15+i32(-65))&i32(255)) < uint32(i32(26)) {
						p55 = i32(32)
					}
					m.memory[uint32(t56)] = byte(p55 | v15)
					t58 := v1 + i32(4)
					p57 := i32(0)
					if uint32((v14+i32(-65))&i32(255)) < uint32(i32(26)) {
						p57 = i32(32)
					}
					m.memory[uint32(t58)] = byte(p57 | v14)
					t60 := v1 + i32(3)
					p59 := i32(0)
					if uint32((v13+i32(-65))&i32(255)) < uint32(i32(26)) {
						p59 = i32(32)
					}
					m.memory[uint32(t60)] = byte(p59 | v13)
					t62 := v1 + i32(2)
					p61 := i32(0)
					if uint32((v12+i32(-65))&i32(255)) < uint32(i32(26)) {
						p61 = i32(32)
					}
					m.memory[uint32(t62)] = byte(p61 | v12)
					t64 := v1 + i32(1)
					p63 := i32(0)
					if uint32((v10+i32(-65))&i32(255)) < uint32(i32(26)) {
						p63 = i32(32)
					}
					m.memory[uint32(t64)] = byte(p63 | v10)
					t66 := v1
					p65 := i32(0)
					if uint32((v11+i32(-65))&i32(255)) < uint32(i32(26)) {
						p65 = i32(32)
					}
					m.memory[uint32(t66)] = byte(p65 | v11)
					v9 = v9 + i32(16)
					v8 = v8 + i32(-16)
					if uint32(v8) > uint32(i32(15)) {
						goto l6
					}
				}
				if v8 == 0 {
					goto l2
				}
				v1 = v6 + v9
				v2 = v7 + v9
			l4:
				v10 = v8 + v5
			l8:
				{
					t67 := int32(int8(m.memory[uint32(v2)]))
					v9 = t67
					if v9 < i32(0) {
						goto l7
					}
					t69 := v1
					p68 := i32(0)
					if uint32((v9+i32(-65))&i32(255)) < uint32(i32(26)) {
						p68 = i32(32)
					}
					m.memory[uint32(t69)] = byte(p68 | v9)
					v1 = v1 + i32(1)
					v2 = v2 + i32(1)
					v5 = v5 + i32(1)
					v8 = v8 + i32(-1)
					if v8 != 0 {
						goto l8
					}
				}
				v5 = v10
			}
		l2:
			store32(m.memory[int64(uint32(v3))+16:], uint32(v5))
			store32(m.memory[int64(uint32(v3))+12:], uint32(v6))
			store32(m.memory[int64(uint32(v3))+8:], uint32(v4))
			goto l9
		l7:
			store32(m.memory[int64(uint32(v3))+16:], uint32(v5))
			store32(m.memory[int64(uint32(v3))+12:], uint32(v6))
			v12 = v2 + v8
			store32(m.memory[int64(uint32(v3))+8:], uint32(v4))
			v17 = v7 + v4
			v10 = i32(0)
			v9 = v5
		l98:
			{
				{
					t70 := int32(int8(m.memory[uint32(v2)]))
					v1 = t70
					if v1 > i32(-1) {
						goto l10
					}
					t71 := int32(m.memory[int64(uint32(v2))+1])
					v8 = t71 & i32(63)
					v11 = v1 & i32(31)
					{
						if uint32(v1) > uint32(i32(-33)) {
							goto l11
						}
						v1 = v11<<6 | v8
						v11 = v2 + i32(2)
						goto l12
					l11:
						t72 := int32(m.memory[int64(uint32(v2))+2])
						v8 = v8<<6 | t72&i32(63)
						if uint32(v1) >= uint32(i32(-16)) {
							goto l13
						}
						v1 = v8 | v11<<12
						v11 = v2 + i32(3)
						goto l12
					l13:
						t73 := int32(m.memory[int64(uint32(v2))+3])
						v1 = v8<<6 | t73&i32(63) | v11<<18&i32(0x1c0000)
						v11 = v2 + i32(4)
					}
				l12:
					v14 = v10 - v2 + v11
					if v1 == i32(931) {
						v16 = i32(131)
						v15 = v10 + v5
						if v15 == 0 {
							goto l16
						}
						{
							if uint32(v15) < uint32(v4) {
								goto l17
							}
							if v15 == v4 {
								goto l18
							}
							goto l19
						l17:
							t74 := int32(int8(m.memory[uint32(v7+v15)]))
							if t74 < i32(-64) {
								goto l19
							}
						}
					l18:
						v2 = v7 + v15
						{
						l34:
							{
								v8 = v2 + i32(-1)
								t75 := int32(int8(m.memory[uint32(v8)]))
								v1 = t75
								if v1 > i32(-1) {
									goto l20
								}
								{
									v6 = v2 + i32(-2)
									t76 := int32(m.memory[uint32(v6)])
									v8 = t76
									v10 = int32(int8(v8))
									if v10 < i32(-64) {
										goto l21
									}
									v8 = v8 & i32(31)
									v2 = v6
									goto l22
								}
							l21:
								{
									{
										v6 = v2 + i32(-3)
										t77 := int32(m.memory[uint32(v6)])
										v8 = t77
										v13 = int32(int8(v8))
										if v13 <= i32(-65) {
											goto l23
										}
										v8 = v8 & i32(15)
										v2 = v6
										goto l24
									}
								l23:
									v2 = v2 + i32(-4)
									t78 := int32(m.memory[uint32(v2)])
									v8 = t78&i32(7)<<6 | v13&i32(63)
								}
							l24:
								v8 = v8<<6 | v10&i32(63)
							l22:
								v1 = v8<<6 | v1&i32(63)
								if uint32(v8) >= uint32(i32(2)) {
									if uint32(v1) <= uint32(i32(167)) {
										goto l28
									}
									t79 := m.fn1464(v1)
									if t79 != 0 {
										goto l29
									}
									goto l28
								}
								v8 = v2
							}
						l20:
							v2 = v1 + i32(-39)
							if uint32(v2) <= uint32(i32(19)) {
								if i32_shl(i32(1), v2)&i32(524417) == 0 {
									goto l27
								}
								v2 = v8
								goto l29
							}
							goto l27
						l27:
							v2 = v8
							switch v1 + i32(-94) {
							case 0, 2:
								goto l29
							default:
								goto l28
							}
						l28:
							{
								if uint32(v1&i32(2097119)+i32(-65)) < uint32(i32(26)) {
									goto l30
								}
								if uint32(v1) < uint32(i32(170)) {
									goto l16
								}
								t80 := m.fn1465(v1)
								if t80 != 0 {
									goto l30
								}
								t81 := m.fn1466(v1)
								if t81 != 0 {
									goto l30
								}
								if uint32(v1) < uint32(i32(453)) {
									goto l16
								}
								t82 := m.fn1467(v1)
								if t82 == 0 {
									goto l16
								}
							}
						l30:
							v2 = v15 + i32(2)
							if v2 == 0 {
								goto l31
							}
							if uint32(v2) < uint32(v4) {
								goto l32
							}
							if v2 == v4 {
								goto l31
							}
							goto l33
						l29:
							if v7 != v2 {
								goto l34
							}
							goto l16
						l32:
							t83 := int32(int8(m.memory[uint32(v7+v2)]))
							if t83 < i32(-64) {
								goto l33
							}
						}
					l31:
						v16 = i32(130)
						if v2 == v4 {
							goto l16
						}
						v2 = v7 + v2
					l44:
						{
							t84 := int32(int8(m.memory[uint32(v2)]))
							v1 = t84
							if v1 <= i32(-1) {
								t85 := int32(m.memory[int64(uint32(v2))+1])
								v8 = t85 & i32(63)
								v6 = v1 & i32(31)
								{
									if uint32(v1) > uint32(i32(-33)) {
										goto l37
									}
									v1 = v6<<6 | v8
									v2 = v2 + i32(2)
									goto l38
								l37:
									t86 := int32(m.memory[int64(uint32(v2))+2])
									v8 = v8<<6 | t86&i32(63)
									if uint32(v1) >= uint32(i32(-16)) {
										goto l39
									}
									v1 = v8 | v6<<12
									v2 = v2 + i32(3)
									goto l38
								l39:
									t87 := int32(m.memory[int64(uint32(v2))+3])
									v1 = v8<<6 | t87&i32(63) | v6<<18&i32(0x1c0000)
									v2 = v2 + i32(4)
								}
							l38:
								if uint32(v1) < uint32(i32(128)) {
									goto l36
								}
								if uint32(v1) <= uint32(i32(167)) {
									goto l40
								}
								t88 := m.fn1464(v1)
								if t88 != 0 {
									goto l41
								}
								goto l40
							}
							v2 = v2 + i32(1)
							v1 = v1 & i32(255)
							goto l36
						}
					l36:
						v8 = v1 + i32(-39)
						if uint32(v8) > uint32(i32(19)) {
							goto l42
						}
						if i32_shl(i32(1), v8)&i32(524417) != 0 {
							goto l41
						}
					l42:
						switch v1 + i32(-94) {
						case 0, 2:
							goto l41
						default:
							goto l40
						}
					l40:
						{
							if uint32(v1&i32(2097119)+i32(-65)) < uint32(i32(26)) {
								goto l43
							}
							if uint32(v1) < uint32(i32(170)) {
								goto l16
							}
							t89 := m.fn1465(v1)
							if t89 != 0 {
								goto l43
							}
							t90 := m.fn1466(v1)
							if t90 != 0 {
								goto l43
							}
							if uint32(v1) < uint32(i32(453)) {
								goto l16
							}
							t91 := m.fn1467(v1)
							if t91 == 0 {
								goto l16
							}
						}
					l43:
						v16 = i32(131)
						goto l16
					l41:
						if v2 != v17 {
							goto l44
						}
					l16:
						{
							t92 := int32(load32(m.memory[int64(uint32(v3))+8:]))
							if uint32(t92-v9) > uint32(i32(1)) {
								goto l45
							}
							m.fn87(v3+i32(8), v9, i32(2))
						}
					l45:
						t93 := int32(load32(m.memory[int64(uint32(v3))+12:]))
						v6 = t93
						v2 = v6 + v9
						m.memory[int64(uint32(v2))+1] = byte(v16)
						m.memory[uint32(v2)] = byte(i32(207))
						t94 := v3
						v9 = v9 + i32(2)
						store32(m.memory[int64(uint32(t94))+16:], uint32(v9))
						v10 = v14
						v2 = v11
						goto l46
					}
					v10 = v14
					v2 = v11
					goto l15
				}
			l10:
				v1 = v1 & i32(255)
				t95 := v10 - v2
				v2 = v2 + i32(1)
				v10 = t95 + v2
			}
		l15:
			m.fn49(v3+i32(20), v1)
			{
				{
					t96 := int32(load32(m.memory[int64(uint32(v3))+24:]))
					v1 = t96
					if v1 != 0 {
						t102 := int32(load32(m.memory[int64(uint32(v3))+20:]))
						v8 = t102
						t103 := int32(load32(m.memory[int64(uint32(v3))+28:]))
						v11 = t103
						if v11 == 0 {
							{
								{
									var p118 int32
									if uint32(v8) < uint32(i32(128)) {
										p118 = 1
									}
									v13 = p118
									if v13 == 0 {
										goto l82
									}
									v11 = i32(1)
									goto l83
								}
							l82:
								if uint32(v8) >= uint32(i32(2048)) {
									goto l84
								}
								v11 = i32(2)
								goto l83
							l84:
								p119 := i32(4)
								if uint32(v8) < uint32(i32(65536)) {
									p119 = i32(3)
								}
								v11 = p119
							}
						l83:
							{
								t120 := int32(load32(m.memory[int64(uint32(v3))+8:]))
								if uint32(v11) <= uint32(t120-v9) {
									goto l85
								}
								m.fn87(v3+i32(8), v9, v11)
								t121 := int32(load32(m.memory[int64(uint32(v3))+12:]))
								v6 = t121
							}
						l85:
							v6 = v6 + v9
							if v13 != 0 {
								goto l86
							}
							v13 = v8&i32(63) | i32(-128)
							v14 = int32(uint32(v8) >> 6)
							if uint32(v8) >= uint32(i32(2048)) {
								v15 = int32(uint32(v8) >> 12)
								v14 = v14&i32(63) | i32(-128)
								if uint32(v8) > uint32(i32(0xffff)) {
									m.memory[int64(uint32(v6))+3] = byte(v13)
									m.memory[int64(uint32(v6))+2] = byte(v14)
									m.memory[int64(uint32(v6))+1] = byte(v15&i32(63) | i32(-128))
									m.memory[uint32(v6)] = byte(int32(uint32(v8)>>18) | i32(-16))
									goto l88
								}
								m.memory[int64(uint32(v6))+2] = byte(v13)
								m.memory[int64(uint32(v6))+1] = byte(v14)
								m.memory[uint32(v6)] = byte(v15 | i32(224))
								goto l88
							}
							m.memory[int64(uint32(v6))+1] = byte(v13)
							m.memory[uint32(v6)] = byte(v14 | i32(192))
							goto l88
						l86:
							m.memory[uint32(v6)] = byte(v8)
						l88:
							t122 := v3
							v8 = v11 + v9
							store32(m.memory[int64(uint32(t122))+16:], uint32(v8))
							{
								{
									var p123 int32
									if uint32(v1) < uint32(i32(128)) {
										p123 = 1
									}
									v13 = p123
									if v13 == 0 {
										goto l90
									}
									v9 = i32(1)
									goto l91
								}
							l90:
								if uint32(v1) >= uint32(i32(2048)) {
									goto l92
								}
								v9 = i32(2)
								goto l91
							l92:
								p124 := i32(4)
								if uint32(v1) < uint32(i32(65536)) {
									p124 = i32(3)
								}
								v9 = p124
							}
						l91:
							{
								t125 := int32(load32(m.memory[int64(uint32(v3))+8:]))
								if uint32(v9) <= uint32(t125-v8) {
									goto l93
								}
								m.fn87(v3+i32(8), v8, v9)
							}
						l93:
							t126 := int32(load32(m.memory[int64(uint32(v3))+12:]))
							v6 = t126
							v11 = v6 + v8
							if v13 != 0 {
								goto l94
							}
							v13 = v1&i32(63) | i32(-128)
							v14 = int32(uint32(v1) >> 6)
							if uint32(v1) >= uint32(i32(2048)) {
								v15 = int32(uint32(v1) >> 12)
								v14 = v14&i32(63) | i32(-128)
								if uint32(v1) > uint32(i32(0xffff)) {
									m.memory[int64(uint32(v11))+3] = byte(v13)
									m.memory[int64(uint32(v11))+2] = byte(v14)
									m.memory[int64(uint32(v11))+1] = byte(v15&i32(63) | i32(-128))
									m.memory[uint32(v11)] = byte(int32(uint32(v1)>>18) | i32(-16))
									goto l96
								}
								m.memory[int64(uint32(v11))+2] = byte(v13)
								m.memory[int64(uint32(v11))+1] = byte(v14)
								m.memory[uint32(v11)] = byte(v15 | i32(224))
								goto l96
							}
							m.memory[int64(uint32(v11))+1] = byte(v13)
							m.memory[uint32(v11)] = byte(v14 | i32(192))
							goto l96
						l94:
							m.memory[uint32(v11)] = byte(v1)
						l96:
							v9 = v9 + v8
							goto l81
						}
						{
							{
								var p104 int32
								if uint32(v8) < uint32(i32(128)) {
									p104 = 1
								}
								v14 = p104
								if v14 == 0 {
									goto l57
								}
								v13 = i32(1)
								goto l58
							}
						l57:
							if uint32(v8) >= uint32(i32(2048)) {
								goto l59
							}
							v13 = i32(2)
							goto l58
						l59:
							p105 := i32(4)
							if uint32(v8) < uint32(i32(65536)) {
								p105 = i32(3)
							}
							v13 = p105
						}
					l58:
						{
							t106 := int32(load32(m.memory[int64(uint32(v3))+8:]))
							if uint32(v13) <= uint32(t106-v9) {
								goto l60
							}
							m.fn87(v3+i32(8), v9, v13)
							t107 := int32(load32(m.memory[int64(uint32(v3))+12:]))
							v6 = t107
						}
					l60:
						v6 = v6 + v9
						if v14 != 0 {
							goto l61
						}
						v14 = v8&i32(63) | i32(-128)
						v15 = int32(uint32(v8) >> 6)
						if uint32(v8) >= uint32(i32(2048)) {
							v16 = int32(uint32(v8) >> 12)
							v15 = v15&i32(63) | i32(-128)
							if uint32(v8) > uint32(i32(0xffff)) {
								m.memory[int64(uint32(v6))+3] = byte(v14)
								m.memory[int64(uint32(v6))+2] = byte(v15)
								m.memory[int64(uint32(v6))+1] = byte(v16&i32(63) | i32(-128))
								m.memory[uint32(v6)] = byte(int32(uint32(v8)>>18) | i32(-16))
								goto l63
							}
							m.memory[int64(uint32(v6))+2] = byte(v14)
							m.memory[int64(uint32(v6))+1] = byte(v15)
							m.memory[uint32(v6)] = byte(v16 | i32(224))
							goto l63
						}
						m.memory[int64(uint32(v6))+1] = byte(v14)
						m.memory[uint32(v6)] = byte(v15 | i32(192))
						goto l63
					}
					{
						{
							t97 := int32(load32(m.memory[int64(uint32(v3))+20:]))
							v1 = t97
							var p98 int32
							if uint32(v1) < uint32(i32(128)) {
								p98 = 1
							}
							v13 = p98
							if v13 == 0 {
								goto l48
							}
							v8 = i32(1)
							goto l49
						}
					l48:
						if uint32(v1) >= uint32(i32(2048)) {
							goto l50
						}
						v8 = i32(2)
						goto l49
					l50:
						p99 := i32(4)
						if uint32(v1) < uint32(i32(65536)) {
							p99 = i32(3)
						}
						v8 = p99
					}
				l49:
					{
						t100 := int32(load32(m.memory[int64(uint32(v3))+8:]))
						if uint32(v8) <= uint32(t100-v9) {
							goto l51
						}
						m.fn87(v3+i32(8), v9, v8)
						t101 := int32(load32(m.memory[int64(uint32(v3))+12:]))
						v6 = t101
					}
				l51:
					v11 = v6 + v9
					if v13 != 0 {
						m.memory[uint32(v11)] = byte(v1)
						goto l54
					}
					v13 = v1&i32(63) | i32(-128)
					v14 = int32(uint32(v1) >> 6)
					if uint32(v1) >= uint32(i32(2048)) {
						v15 = int32(uint32(v1) >> 12)
						v14 = v14&i32(63) | i32(-128)
						if uint32(v1) > uint32(i32(0xffff)) {
							m.memory[int64(uint32(v11))+3] = byte(v13)
							m.memory[int64(uint32(v11))+2] = byte(v14)
							m.memory[int64(uint32(v11))+1] = byte(v15&i32(63) | i32(-128))
							m.memory[uint32(v11)] = byte(int32(uint32(v1)>>18) | i32(-16))
							goto l54
						}
						m.memory[int64(uint32(v11))+2] = byte(v13)
						m.memory[int64(uint32(v11))+1] = byte(v14)
						m.memory[uint32(v11)] = byte(v15 | i32(224))
						goto l54
					}
					m.memory[int64(uint32(v11))+1] = byte(v13)
					m.memory[uint32(v11)] = byte(v14 | i32(192))
					goto l54
				}
			l61:
				m.memory[uint32(v6)] = byte(v8)
			l63:
				t108 := v3
				v8 = v13 + v9
				store32(m.memory[int64(uint32(t108))+16:], uint32(v8))
				{
					{
						var p109 int32
						if uint32(v1) < uint32(i32(128)) {
							p109 = 1
						}
						v14 = p109
						if v14 == 0 {
							goto l65
						}
						v9 = i32(1)
						goto l66
					}
				l65:
					if uint32(v1) >= uint32(i32(2048)) {
						goto l67
					}
					v9 = i32(2)
					goto l66
				l67:
					p110 := i32(4)
					if uint32(v1) < uint32(i32(65536)) {
						p110 = i32(3)
					}
					v9 = p110
				}
			l66:
				{
					t111 := int32(load32(m.memory[int64(uint32(v3))+8:]))
					if uint32(v9) <= uint32(t111-v8) {
						goto l68
					}
					m.fn87(v3+i32(8), v8, v9)
				}
			l68:
				t112 := int32(load32(m.memory[int64(uint32(v3))+12:]))
				v6 = t112
				v13 = v6 + v8
				if v14 != 0 {
					goto l69
				}
				v14 = v1&i32(63) | i32(-128)
				v15 = int32(uint32(v1) >> 6)
				if uint32(v1) >= uint32(i32(2048)) {
					v16 = int32(uint32(v1) >> 12)
					v15 = v15&i32(63) | i32(-128)
					if uint32(v1) > uint32(i32(0xffff)) {
						m.memory[int64(uint32(v13))+3] = byte(v14)
						m.memory[int64(uint32(v13))+2] = byte(v15)
						m.memory[int64(uint32(v13))+1] = byte(v16&i32(63) | i32(-128))
						m.memory[uint32(v13)] = byte(int32(uint32(v1)>>18) | i32(-16))
						goto l71
					}
					m.memory[int64(uint32(v13))+2] = byte(v14)
					m.memory[int64(uint32(v13))+1] = byte(v15)
					m.memory[uint32(v13)] = byte(v16 | i32(224))
					goto l71
				}
				m.memory[int64(uint32(v13))+1] = byte(v14)
				m.memory[uint32(v13)] = byte(v15 | i32(192))
				goto l71
			l69:
				m.memory[uint32(v13)] = byte(v1)
			l71:
				t113 := v3
				v1 = v9 + v8
				store32(m.memory[int64(uint32(t113))+16:], uint32(v1))
				{
					{
						var p114 int32
						if uint32(v11) < uint32(i32(128)) {
							p114 = 1
						}
						v13 = p114
						if v13 == 0 {
							goto l73
						}
						v8 = i32(1)
						goto l74
					}
				l73:
					if uint32(v11) >= uint32(i32(2048)) {
						goto l75
					}
					v8 = i32(2)
					goto l74
				l75:
					p115 := i32(4)
					if uint32(v11) < uint32(i32(65536)) {
						p115 = i32(3)
					}
					v8 = p115
				}
			l74:
				{
					t116 := int32(load32(m.memory[int64(uint32(v3))+8:]))
					if uint32(v8) <= uint32(t116-v1) {
						goto l76
					}
					m.fn87(v3+i32(8), v1, v8)
					t117 := int32(load32(m.memory[int64(uint32(v3))+12:]))
					v6 = t117
				}
			l76:
				v9 = v6 + v1
				if v13 != 0 {
					goto l77
				}
				v13 = v11&i32(63) | i32(-128)
				v14 = int32(uint32(v11) >> 6)
				if uint32(v11) >= uint32(i32(2048)) {
					v15 = int32(uint32(v11) >> 12)
					v14 = v14&i32(63) | i32(-128)
					if uint32(v11) > uint32(i32(0xffff)) {
						m.memory[int64(uint32(v9))+3] = byte(v13)
						m.memory[int64(uint32(v9))+2] = byte(v14)
						m.memory[int64(uint32(v9))+1] = byte(v15&i32(63) | i32(-128))
						m.memory[uint32(v9)] = byte(int32(uint32(v11)>>18) | i32(-16))
						goto l79
					}
					m.memory[int64(uint32(v9))+2] = byte(v13)
					m.memory[int64(uint32(v9))+1] = byte(v14)
					m.memory[uint32(v9)] = byte(v15 | i32(224))
					goto l79
				}
				m.memory[int64(uint32(v9))+1] = byte(v13)
				m.memory[uint32(v9)] = byte(v14 | i32(192))
				goto l79
			l77:
				m.memory[uint32(v9)] = byte(v11)
			l79:
				v9 = v8 + v1
				goto l81
			}
		l54:
			v9 = v8 + v9
		l81:
			store32(m.memory[int64(uint32(v3))+16:], uint32(v9))
		l46:
			if v2 != v12 {
				goto l98
			}
		l9:
			t127 := int32(load32(m.memory[int64(uint32(v3))+16:]))
			store32(m.memory[int64(uint32(v0))+8:], uint32(t127))
			t128 := int64(load64(m.memory[int64(uint32(v3))+8:]))
			store64(m.memory[uint32(v0):], uint64(t128))
			m.g0 = v3 + i32(32)
			return
		}
	}
l33:
	m.fn556(v7, v4, v2, v4, i32(1070228))
	panic("unreachable")
l19:
	m.fn556(v7, v4, i32(0), v15, i32(1070212))
	panic("unreachable")
}
func (m *Module) fn1426(v0 int32) {
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
	m.fn78(v3)
	v3 = v3 + i32(12)
	goto l1
l0:
	t2 := int32(load32(m.memory[uint32(v0):]))
	m.fn136(t2, v2, i32(4), i32(12))
}
func (m *Module) fn1427(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7, v8, v9 int32
	t0 := m.g0
	v2 = t0 - i32(96)
	m.g0 = v2
	m.fn46(v2+i32(40), v0, v1)
	{
		{
			t1 := int32(load32(m.memory[int64(uint32(v2))+44:]))
			v0 = t1
			if v0 != 0 {
				goto l0
			}
			v1 = i32(255)
			goto l1
		}
	l0:
		t2 := int32(load32(m.memory[int64(uint32(v2))+40:]))
		v3 = t2
		store32(m.memory[int64(uint32(v2))+48:], uint32(i32(0)))
		m.fn522(v2+i32(32), i32(37), v2+i32(48))
		t3 := int32(load32(m.memory[int64(uint32(v2))+32:]))
		t4 := int32(load32(m.memory[int64(uint32(v2))+36:]))
		m.fn626(v2+i32(24), t3, t4, v3, v0)
		t5 := int32(load32(m.memory[int64(uint32(v2))+28:]))
		v4 = t5
		t6 := int32(load32(m.memory[int64(uint32(v2))+24:]))
		v1 = t6
		store32(m.memory[int64(uint32(v2))+56:], uint32(i32(0)))
		store64(m.memory[int64(uint32(v2))+48:], uint64(i64(0x100000000)))
		m.fn47(v2+i32(48), i32(0))
		t8 := v2
		p7 := v3
		if v1 != 0 {
			p7 = v1
		}
		v5 = p7
		store32(m.memory[int64(uint32(t8))+88:], uint32(v5))
		t10 := v2
		t11 := v5
		p9 := v0
		if v1 != 0 {
			p9 = v4
		}
		store32(m.memory[int64(uint32(t10))+92:], uint32(t11+p9))
	l2:
		{
			t12 := m.fn48(v2 + i32(88))
			v1 = t12
			if v1 == i32(32) {
				goto l2
			}
			if v1 == i32(44) {
				goto l2
			}
			if v1 == i32(95) {
				goto l2
			}
			if v1 == i32(160) {
				goto l2
			}
			if v1 == i32(-1) {
				goto l3
			}
			m.fn74(v2+i32(48), v1)
			goto l2
		l3:
		}
		t13 := int32(load32(m.memory[int64(uint32(v2))+48:]))
		v5 = t13
		t14 := int32(load32(m.memory[int64(uint32(v2))+52:]))
		t15 := v2
		v4 = t14
		t16 := int32(load32(m.memory[int64(uint32(v2))+56:]))
		t17 := v4
		v6 = t16
		store32(m.memory[int64(uint32(t15))+52:], uint32(t17+v6))
		store32(m.memory[int64(uint32(v2))+48:], uint32(v4))
	l5:
		{
			t18 := m.fn48(v2 + i32(48))
			v1 = t18
			if v1 == i32(-1) {
				goto l4
			}
			if uint32(v1+i32(-58)) < uint32(i32(-10)) {
				goto l5
			}
		}
	l4:
		{
			if v1 == i32(-1) {
				goto l6
			}
			m.fn217(v2+i32(48), v4, v6)
			t19 := int32(m.memory[int64(uint32(v2))+48])
			v1 = t19
			m.fn16(v5, v4)
			if v1 != 0 {
				goto l7
			}
			v1 = i32(0)
			goto l1
		}
	l6:
		m.fn16(v5, v4)
	l7:
		m.fn1425(v2+i32(48), v3, v0)
		{
			t20 := int32(load32(m.memory[int64(uint32(v2))+52:]))
			v1 = t20
			t21 := int32(load32(m.memory[int64(uint32(v2))+56:]))
			t22 := v1
			v4 = t21
			t23 := m.fn15(t22, v4, i32(1071691), i32(4))
			if t23 != 0 {
				goto l8
			}
			t24 := m.fn15(v1, v4, i32(1081456), i32(5))
			if t24 != 0 {
				goto l8
			}
			t25 := m.fn15(v1, v4, i32(1081461), i32(3))
			if t25 != 0 {
				goto l8
			}
			t26 := m.fn15(v1, v4, i32(1081464), i32(2))
			if t26 != 0 {
				goto l8
			}
			t27 := int32(load32(m.memory[int64(uint32(v2))+48:]))
			m.fn16(t27, v1)
			store32(m.memory[int64(uint32(v2))+72:], uint32(i32(0)))
			store32(m.memory[int64(uint32(v2))+64:], uint32(v3))
			store32(m.memory[int64(uint32(v2))+56:], uint32(v3))
			store64(m.memory[int64(uint32(v2))+48:], uint64(i64(0x2000000054)))
			store32(m.memory[int64(uint32(v2))+60:], uint32(v0))
			t28 := v2
			v1 = v3 + v0
			store32(m.memory[int64(uint32(t28))+68:], uint32(v1))
			v7 = v2 + i32(64)
			v4 = v3
			{
			l10:
				{
					v8 = v1
					v9 = v4
					m.fn572(v2+i32(16), v7)
					t29 := int32(load32(m.memory[int64(uint32(v2))+20:]))
					v5 = t29
					if v5 == i32(-1) {
						goto l9
					}
					t30 := int32(load32(m.memory[int64(uint32(v2))+16:]))
					v6 = t30
					t31 := int32(load32(m.memory[int64(uint32(v2))+64:]))
					v4 = t31
					t32 := int32(load32(m.memory[int64(uint32(v2))+68:]))
					v1 = t32
					t33 := m.fn576(v5, v2+i32(48), i32(2))
					if t33 == 0 {
						goto l10
					}
				}
				t34 := v3
				v4 = v6 + v8 - (v9 + v1) + v4
				v1 = t34 + v4
				v5 = v0 - v4
				v4 = v3
				goto l11
			}
		l9:
			v4 = i32(0)
		l11:
			t36 := v2 + i32(8)
			p35 := i32(1)
			if v4 != 0 {
				p35 = v1
			}
			p37 := i32(0)
			if v4 != 0 {
				p37 = v5
			}
			m.fn856(t36, p35, p37, i32(90))
			t38 := int32(load32(m.memory[int64(uint32(v2))+8:]))
			t39 := int32(load32(m.memory[int64(uint32(v2))+12:]))
			m.fn513(v2+i32(48), t38, t39, i32(46))
			m.fn515(v2, v2+i32(48))
			t40 := int32(load32(m.memory[int64(uint32(v2))+4:]))
			t41 := int32(load32(m.memory[uint32(v2):]))
			v7 = t41
			p42 := i32(0)
			if v7 != 0 {
				p42 = t40
			}
			v5 = p42
			v1 = i32(3)
			{
				{
					p43 := v3
					if v4 != 0 {
						p43 = v4
					}
					v3 = p43
					t45 := v3
					p44 := v0
					if v4 != 0 {
						p44 = v6
					}
					v0 = p44
					t46 := m.fn1462(t45, v0, i32(1081380), i32(3))
					if t46 != i32(3) {
						goto l12
					}
					if v5 == 0 {
						goto l13
					}
					p47 := i32(1)
					if v7 != 0 {
						p47 = v7
					}
					t48 := m.fn1463(p47, v5)
					if t48 != 0 {
						goto l13
					}
					goto l1
				}
			l12:
				if v5 != 0 {
					goto l1
				}
				t49 := m.fn1463(v3, v0)
				if t49 == 0 {
					goto l1
				}
			}
		l13:
			v1 = i32(2)
			goto l1
		}
	l8:
		t50 := int32(load32(m.memory[int64(uint32(v2))+48:]))
		m.fn16(t50, v1)
		v1 = i32(1)
	}
l1:
	m.g0 = v2 + i32(96)
	return v1
}
func (m *Module) fn1428(v0, v1 int32) int32 {
	var v2, v3 int32
	var v4 float64
	t0 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	v2 = t0
	v3 = v2 & i32(0x200000)
	t1 := math.Float64frombits(load64(m.memory[uint32(v0):]))
	v4 = t1
	{
		if v2&i32(0x10000000) != 0 {
			t6 := int32(load16(m.memory[int64(uint32(v1))+14:]))
			t7 := v1
			t8 := v4
			var p9 int32
			if v3 != i32(0) {
				p9 = 1
			}
			t10 := m.fn1677(t7, t8, p9, t6)
			return t10
		}
		t2 := v1
		t3 := v4
		var p4 int32
		if v3 != i32(0) {
			p4 = 1
		}
		t5 := m.fn1675(t2, t3, p4)
		return t5
	}
}
func (m *Module) fn1429(v0, v1 int32) {
	var v2, v3, v4, v5, v6, v7, v8, v9, v10 int32
	t0 := int32(load32(m.memory[uint32(v1):]))
	v2 = t0
	t1 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	v3 = t1
	t2 := int32(load32(m.memory[int64(uint32(v1))+44:]))
	v4 = t2
	t3 := int32(load16(m.memory[int64(uint32(v1))+36:]))
	v5 = t3
	t4 := int32(load16(m.memory[int64(uint32(v1))+38:]))
	v6 = t4
	t5 := int32(m.memory[int64(uint32(v1))+40])
	v7 = t5
	t6 := int32(m.memory[int64(uint32(v1))+41])
	v8 = t6
	t7 := int32(m.memory[int64(uint32(v1))+42])
	v9 = t7
	t8 := int32(m.memory[int64(uint32(v1))+43])
	v10 = t8
	m.fn1449(v0+i32(8), v1+i32(8))
	m.memory[int64(uint32(v0))+43] = byte(v10)
	m.memory[int64(uint32(v0))+42] = byte(v9)
	m.memory[int64(uint32(v0))+41] = byte(v8)
	m.memory[int64(uint32(v0))+40] = byte(v7)
	store16(m.memory[int64(uint32(v0))+38:], uint16(v6))
	store16(m.memory[int64(uint32(v0))+36:], uint16(v5))
	store32(m.memory[int64(uint32(v0))+44:], uint32(v4))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
	store32(m.memory[uint32(v0):], uint32(v2))
}
func (m *Module) fn1430(v0, v1, v2 int32) int32 {
	var v3 int32
	v3 = i32(1)
	switch int32(uint32(v0&i32(57344)) >> 13) {
	case 2, 4, 5:
		return i32(2)
	case 3:
		return i32(4)
	case 7:
		v3 = i32(3)
		fallthrough
	default:
		return v3
	case 6:
		{
			if v0&i32(0xffff) == i32(54792) {
				goto l5
			}
			if v2 == 0 {
				goto l6
			}
			t0 := int32(m.memory[uint32(v1)])
			return t0 + i32(1)
		}
	l5:
		if uint32(v2) >= uint32(i32(2)) {
			t1 := int32(load16(m.memory[uint32(v1):]))
			return t1 + i32(1)
		}
	l6:
		return i32(0)
	}
}
func (m *Module) fn1431(v0, v1, v2 int32) int32 {
	var v3 int32
	v3 = i32(2)
	if v1 == 0 {
		goto l0
	}
	{
		t0 := int32(m.memory[uint32(v0)])
		v1 = t0
		switch v1 + i32(-128) {
		default:
			if uint32(v1) >= uint32(i32(2)) {
				goto l0
			}
			return v1
		case 0:
			return v2
		case 1:
			v3 = v2 ^ i32(1)
		}
	}
l0:
	return v3
}
func (m *Module) fn1432(v0, v1, v2 int32) int32 {
	var v3 int32
	t0 := m.g0
	v3 = t0 - i32(16)
	m.g0 = v3
	m.fn520(v3+i32(8), v2, v0, v1)
	t1 := int32(load32(m.memory[int64(uint32(v3))+8:]))
	v1 = t1
	m.g0 = v3 + i32(16)
	var p2 int32
	if v1 == i32(1) {
		p2 = 1
	}
	return p2
}
func (m *Module) fn1433(v0, v1 int32) {
	m.fn1441(v1)
l1:
	{
		t0 := int32(load32(m.memory[int64(uint32(v1))+20:]))
		if t0 == 0 {
			goto l0
		}
		m.fn1442(v1)
		goto l1
	}
l0:
	t1 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	store32(m.memory[int64(uint32(v0))+8:], uint32(t1))
	t2 := int64(load64(m.memory[uint32(v1):]))
	store64(m.memory[uint32(v0):], uint64(t2))
	m.fn1444(v1 + i32(12))
	t3 := int32(load32(m.memory[int64(uint32(v1))+24:]))
	t4 := int32(load32(m.memory[int64(uint32(v1))+28:]))
	m.fn16(t3, t4)
}
func (m *Module) fn1434(v0, v1, v2, v3, v4 int32) {
	var v5, v6 int32
	var v7 int64
	var v8, v9, v10, v11, v12, v13, v14, v15, v16, v17, v18, v19 int32
	var v20, v21, v22 int64
	var v23 int32
	var v24 int64
	var v25 int32
	var v26 int64
	t0 := m.g0
	v5 = t0 - i32(128)
	m.g0 = v5
	{
		t1 := int32(load32(m.memory[int64(uint32(v4))+8:]))
		if t1 == 0 {
			goto l0
		}
		t2 := int32(load32(m.memory[int64(uint32(v4))+8:]))
		v6 = t2
		store32(m.memory[int64(uint32(v4))+8:], uint32(i32(0)))
		t3 := int64(load64(m.memory[uint32(v4):]))
		v7 = t3
		store64(m.memory[uint32(v4):], uint64(i64(0x800000000)))
		store32(m.memory[int64(uint32(v5))+24:], uint32(v6))
		store64(m.memory[int64(uint32(v5))+16:], uint64(v7))
		m.fn1169(v3, v5+i32(16))
	}
l0:
	{
		t4 := int32(load32(m.memory[int64(uint32(v3))+8:]))
		if t4 == 0 {
			goto l1
		}
		t5 := int32(load32(m.memory[int64(uint32(v3))+8:]))
		v4 = t5
		store32(m.memory[int64(uint32(v3))+8:], uint32(i32(0)))
		t6 := int64(load64(m.memory[uint32(v3):]))
		v7 = t6
		store64(m.memory[uint32(v3):], uint64(i64(0x400000000)))
		store32(m.memory[int64(uint32(v5))+28:], uint32(i32(-1)))
		store32(m.memory[int64(uint32(v5))+24:], uint32(v4))
		store64(m.memory[int64(uint32(v5))+16:], uint64(v7))
		m.fn1450(v2, v5+i32(16))
	}
l1:
	{
		t7 := int32(load32(m.memory[int64(uint32(v2))+8:]))
		v3 = t7
		if v3 == 0 {
			goto l2
		}
		t8 := int32(load32(m.memory[int64(uint32(v2))+4:]))
		v8 = t8
		store64(m.memory[int64(uint32(v2))+4:], uint64(i64(4)))
		t9 := int32(load32(m.memory[uint32(v2):]))
		v9 = t9
		store32(m.memory[uint32(v2):], uint32(i32(0)))
		v10 = v8 + v3*i32(40)
		v11 = v5 + i32(16) + i32(25)
		v12 = v5 + i32(16) + i32(4)
		v13 = v8
		v14 = v8
	l11:
		{
			if v13 == v10 {
				m.fn84(i32(0), i32(4))
				v2 = v8
				{
					v3 = v9 * i32(40)
					if v3&i32(8) == 0 {
						goto l12
					}
					t41 := v8
					t42 := v3
					v4 = v3 & i32(-16)
					t43 := m.fn392(t41, t42, v4)
					v2 = t43
					if v2 == 0 {
						m.fn85(i32(4), v4)
						panic("unreachable")
					}
				}
			l12:
				store32(m.memory[int64(uint32(v5))+64:], uint32(v2))
				store32(m.memory[int64(uint32(v5))+60:], uint32(int32(uint32(v3)>>4)))
				store32(m.memory[int64(uint32(v5))+68:], uint32(int32(uint32(v14-v8)>>4)))
				m.fn84(i32(0), i32(4))
				m.fn1452(v5+i32(16), v5+i32(60))
				v3 = v5 + i32(16) | i32(4)
				{
					t44 := int32(load32(m.memory[int64(uint32(v5))+16:]))
					v2 = t44
					if v2 != i32(-2) {
						t51 := int32(load32(m.memory[int64(uint32(v3))+24:]))
						store32(m.memory[int64(uint32(v5))+96:], uint32(t51))
						t52 := int64(load64(m.memory[int64(uint32(v3))+16:]))
						store64(m.memory[int64(uint32(v5))+88:], uint64(t52))
						t53 := int64(load64(m.memory[int64(uint32(v3))+8:]))
						store64(m.memory[int64(uint32(v5))+80:], uint64(t53))
						t54 := int64(load64(m.memory[uint32(v3):]))
						store64(m.memory[int64(uint32(v5))+72:], uint64(t54))
						{
							if v2 == i32(-1) {
								goto l16
							}
							store32(m.memory[int64(uint32(v5))+16:], uint32(v2))
							t55 := int64(load64(m.memory[int64(uint32(v5))+72:]))
							store64(m.memory[int64(uint32(v5))+20:], uint64(t55))
							t56 := int64(load64(m.memory[int64(uint32(v5))+80:]))
							store64(m.memory[int64(uint32(v5))+28:], uint64(t56))
							t57 := int64(load64(m.memory[int64(uint32(v5))+88:]))
							store64(m.memory[int64(uint32(v5))+36:], uint64(t57))
							t58 := int32(load32(m.memory[int64(uint32(v5))+96:]))
							store32(m.memory[int64(uint32(v5))+44:], uint32(t58))
							m.fn338(v1, v5+i32(16))
						}
					l16:
						store32(m.memory[uint32(v0):], uint32(i32(-1)))
						goto l15
					}
					t45 := int64(load64(m.memory[int64(uint32(v3))+16:]))
					t46 := v5
					v7 = t45
					store64(m.memory[int64(uint32(t46))+88:], uint64(v7))
					t47 := int64(load64(m.memory[int64(uint32(v3))+8:]))
					t48 := v5
					v20 = t47
					store64(m.memory[int64(uint32(t48))+80:], uint64(v20))
					t49 := int64(load64(m.memory[uint32(v3):]))
					t50 := v5
					v24 = t49
					store64(m.memory[int64(uint32(t50))+72:], uint64(v24))
					store64(m.memory[int64(uint32(v0))+16:], uint64(v7))
					store64(m.memory[int64(uint32(v0))+8:], uint64(v20))
					store64(m.memory[uint32(v0):], uint64(v24))
					goto l15
				}
			}
			{
				t10 := int32(load32(m.memory[int64(uint32(v13))+12:]))
				v2 = t10
				if v2 == i32(-1) {
					goto l4
				}
				t11 := int32(m.memory[int64(uint32(v13))+36])
				v15 = t11
				t12 := int32(load32(m.memory[int64(uint32(v13))+32:]))
				store32(m.memory[int64(uint32(v12))+16:], uint32(t12))
				t13 := int64(load64(m.memory[int64(uint32(v13))+24:]))
				store64(m.memory[int64(uint32(v12))+8:], uint64(t13))
				t14 := int64(load64(m.memory[int64(uint32(v13))+16:]))
				store64(m.memory[uint32(v12):], uint64(t14))
				t15 := int32(load16(m.memory[int64(uint32(v13))+37:]))
				store16(m.memory[uint32(v11):], uint16(t15))
				t16 := int32(m.memory[int64(uint32(v13))+39])
				m.memory[int64(uint32(v11))+2] = byte(t16)
				goto l5
			}
		l4:
			v15 = i32(0)
			store32(m.memory[int64(uint32(v5))+36:], uint32(i32(0)))
			store64(m.memory[int64(uint32(v5))+28:], uint64(i64(0x100000000)))
			store64(m.memory[int64(uint32(v5))+20:], uint64(i64(2)))
			v2 = i32(0)
		l5:
			t17 := int32(load32(m.memory[int64(uint32(v13))+8:]))
			v3 = t17
			t18 := int32(load32(m.memory[int64(uint32(v13))+4:]))
			v16 = t18
			t19 := int32(load32(m.memory[uint32(v13):]))
			v17 = t19
			store32(m.memory[int64(uint32(v5))+16:], uint32(v2))
			m.memory[int64(uint32(v5))+40] = byte(v15)
			m.fn59(v5+i32(8), v3, i32(8), i32(32))
			v2 = i32(0)
			store32(m.memory[int64(uint32(v5))+124:], uint32(i32(0)))
			t20 := int32(load32(m.memory[int64(uint32(v5))+12:]))
			t21 := v5
			v4 = t20
			store32(m.memory[int64(uint32(t21))+120:], uint32(v4))
			t22 := int32(load32(m.memory[int64(uint32(v5))+8:]))
			t23 := v5
			v6 = t22
			store32(m.memory[int64(uint32(t23))+116:], uint32(v6))
			v18 = i32(0)
			{
				if uint32(v3) <= uint32(v6) {
					goto l6
				}
				m.fn62(v5+i32(116), i32(0), v3, i32(8), i32(32))
				t24 := int32(load32(m.memory[int64(uint32(v5))+124:]))
				v18 = t24
				t25 := int32(load32(m.memory[int64(uint32(v5))+120:]))
				v4 = t25
			}
		l6:
			v13 = v13 + i32(40)
			v19 = v3 * i32(12)
			t26 := int32(load32(m.memory[int64(uint32(v5))+20:]))
			v6 = t26 + i32(2)
			v3 = v4 + v18<<5
			v7 = i64(0)
			v20 = i64(1000)
			t27 := int64(load32(m.memory[int64(uint32(v5))+24:]))
			v21 = t27
			t28 := int64(load32(m.memory[int64(uint32(v5))+36:]))
			v22 = t28
			t29 := int32(load32(m.memory[int64(uint32(v5))+32:]))
			v23 = t29
		l10:
			{
				if v19 == v2 {
					m.fn911(v17, v16)
					store32(m.memory[int64(uint32(v5))+112:], uint32(v18))
					t38 := int64(load64(m.memory[int64(uint32(v5))+116:]))
					store64(m.memory[int64(uint32(v5))+104:], uint64(t38))
					m.fn767(v5 + i32(16))
					t39 := int32(load32(m.memory[int64(uint32(v5))+112:]))
					store32(m.memory[int64(uint32(v14))+8:], uint32(t39))
					t40 := int64(load64(m.memory[int64(uint32(v5))+104:]))
					store64(m.memory[uint32(v14):], uint64(t40))
					m.memory[int64(uint32(v14))+12] = byte(v15 & i32(1))
					v14 = v14 + i32(16)
					goto l11
				}
				v4 = i32(2)
				{
					if uint64(v7) >= uint64(v22) {
						goto l8
					}
					t30 := int32(load32(m.memory[uint32(v23):]))
					v4 = t30
				}
			l8:
				{
					v7 = v7 + i64(1)
					if uint64(v7) >= uint64(v21) {
						goto l9
					}
					t31 := int64(int16(load16(m.memory[uint32(v6):])))
					v24 = t31
				}
			l9:
				v25 = v16 + v2
				t32 := int64(load64(m.memory[uint32(v25):]))
				v26 = t32
				t33 := int32(load32(m.memory[int64(uint32(v25))+8:]))
				store32(m.memory[int64(uint32(v3))+8:], uint32(t33))
				store64(m.memory[uint32(v3):], uint64(v26))
				t35 := v3 + i32(16)
				p34 := v20
				if uint64(v7) < uint64(v21) {
					p34 = v24
				}
				store64(m.memory[uint32(t35):], uint64(p34))
				t37 := v3 + i32(24)
				p36 := v4
				if v4&i32(255) == i32(2) {
					p36 = i32(0)
				}
				v4 = p36
				m.memory[uint32(t37)] = byte(v4 & i32(1))
				m.memory[uint32(v3+i32(27))] = byte(int32(uint32(v4)>>24) & i32(1))
				m.memory[uint32(v3+i32(26))] = byte(int32(uint32(v4)>>16) & i32(1))
				m.memory[uint32(v3+i32(25))] = byte(int32(uint32(v4)>>8) & i32(1))
				v2 = v2 + i32(12)
				v3 = v3 + i32(32)
				v6 = v6 + i32(2)
				v20 = v20 + i64(1000)
				v23 = v23 + i32(4)
				v18 = v18 + i32(1)
				goto l10
			}
		}
	}
l2:
	store32(m.memory[uint32(v0):], uint32(i32(-1)))
l15:
	m.g0 = v5 + i32(128)
}
func (m *Module) fn1435(v0, v1 int32) {
	var v2 int32
	m.fn1441(v0)
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+20:]))
		v2 = t0
		if v2 == 0 {
			goto l0
		}
		t1 := int32(load32(m.memory[int64(uint32(v0))+16:]))
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
	m.fn1340(v0, v1)
}
func (m *Module) fn1436(v0, v1, v2 int32) int32 {
	var v3, v4, v5, v6, v7, v8, v9, v10 int32
	t0 := m.g0
	v3 = t0 - i32(16)
	m.g0 = v3
	v4 = i32(0)
	v5 = i32(0)
	{
		t1 := int32(load32(m.memory[uint32(v0+i32(344)):]))
		t2 := int32(load32(m.memory[uint32(v0+i32(348)):]))
		t3 := m.fn1438(t1, t2, v1)
		v6 = t3
		if v6 == 0 {
			goto l0
		}
		t4 := int32(load16(m.memory[int64(uint32(v6))+60:]))
		v5 = t4
	}
l0:
	{
		{
			t5 := int32(load32(m.memory[uint32(v0+i32(332)):]))
			t6 := int32(load32(m.memory[uint32(v0+i32(336)):]))
			t7 := m.fn1438(t5, t6, v1)
			v1 = t7
			if v1 != 0 {
				goto l1
			}
			goto l2
		}
	l1:
		t8 := int32(load32(m.memory[int64(uint32(v1))+56:]))
		v7 = t8
		t9 := int32(load32(m.memory[int64(uint32(v1))+52:]))
		v4 = t9
	}
l2:
	v6 = i32(0)
	p10 := i32(0)
	if v4 != 0 {
		p10 = v7
	}
	v1 = p10
	p11 := i32(1)
	if v4 != 0 {
		p11 = v4
	}
	v7 = p11
	v8 = i32(0)
	{
	l7:
		{
			t12 := v1
			v4 = v6 + i32(2)
			if uint32(t12) < uint32(v4) {
				goto l3
			}
			if uint32(v6) >= uint32(v1) {
				m.fn158(v6, v1, i32(1072496))
				panic("unreachable")
			}
			v9 = v6 + i32(1)
			if uint32(v9) >= uint32(v1) {
				m.fn158(v9, v1, i32(1072512))
				panic("unreachable")
			}
			t13 := int32(m.memory[uint32(v7+v6)])
			v6 = t13
			t14 := int32(m.memory[uint32(v7+v9)])
			v9 = t14
			m.fn148(v3+i32(8), v4, v7, v1, i32(1072528))
			v9 = v6 | v9<<8
			t15 := int32(load32(m.memory[int64(uint32(v3))+8:]))
			t16 := int32(load32(m.memory[int64(uint32(v3))+12:]))
			t17 := m.fn1430(v9, t15, t16)
			v6 = t17
			if uint32(v6) > uint32(v1-v4) {
				goto l3
			}
			{
				if v9&i32(0xffff) != i32(18992) {
					goto l6
				}
				v8 = i32(0)
				if uint32(v6) < uint32(i32(2)) {
					goto l6
				}
				t18 := int32(load16(m.memory[uint32(v7+v4):]))
				v10 = t18
				v8 = i32(1)
			}
		l6:
			v6 = v6 + v4
			goto l7
		}
	l3:
		t20 := v7
		t21 := v1
		t22 := v0
		p19 := v5
		if v8&i32(1) != 0 {
			p19 = v10
		}
		t23 := m.fn1439(t22, p19)
		t24 := int32(load32(m.memory[int64(uint32(t23))+48:]))
		v6 = t24
		t25 := m.fn1215(t20, t21, v6, v6)
		v1 = t25
		{
			t26 := int32(load32(m.memory[int64(uint32(v0))+132:]))
			if uint32(v2) >= uint32(t26) {
				goto l8
			}
			t27 := int32(load32(m.memory[int64(uint32(v0))+128:]))
			t28 := int32(load32(m.memory[uint32(t27+v2<<2):]))
			m.fn1440(v3, v0, t28)
			t29 := int32(load32(m.memory[uint32(v3):]))
			v4 = t29
			if v4 == 0 {
				goto l8
			}
			t30 := int32(load32(m.memory[int64(uint32(v3))+4:]))
			t31 := m.fn1215(v4, t30, v1, v6)
			v1 = t31
		}
	l8:
		m.g0 = v3 + i32(16)
		return v1
	}
}
