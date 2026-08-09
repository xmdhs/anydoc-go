package core

import (
	"math/bits"
)

func (m *Module) fn1437(v0, v1, v2 int32) {
	var v3 int32
	t0 := m.g0
	v3 = t0 - i32(16)
	m.g0 = v3
	store32(m.memory[int64(uint32(v3))+12:], uint32(v2))
	{
		t1 := m.fn822(v3+i32(12), v0+i32(36))
		if t1 == 0 {
			goto l0
		}
		m.fn1441(v0)
		store32(m.memory[int64(uint32(v0))+36:], uint32(v2))
	}
l0:
	m.fn74(v0+i32(24), v1)
	m.g0 = v3 + i32(16)
}
func (m *Module) fn1438(v0, v1, v2 int32) int32 {
	var v3, v4, v5, v6 int32
	if v1 == 0 {
		goto l0
	}
	v3 = i32(0)
	v4 = v1
l2:
	{
		if uint32(v4) < uint32(i32(2)) {
			t3 := int32(load32(m.memory[int64(uint32(v0+v3*i32(72)))+64:]))
			t4 := v3
			var p5 int32
			if uint32(t3) <= uint32(v2) {
				p5 = 1
			}
			v4 = t4 + p5
			if v4 == 0 {
				goto l0
			}
			{
				v4 = v4 + i32(-1)
				if uint32(v4) >= uint32(v1) {
					m.fn158(v4, v1, i32(1079932))
					panic("unreachable")
				}
				v4 = v0 + v4*i32(72)
				t6 := int32(load32(m.memory[int64(uint32(v4))+68:]))
				p7 := i32(0)
				if uint32(v2) < uint32(t6) {
					p7 = v4
				}
				return p7
			}
		}
		t0 := v3
		v5 = int32(uint32(v4) >> 1)
		v6 = v5 + v3
		t1 := int32(load32(m.memory[int64(uint32(v0+v6*i32(72)))+64:]))
		p2 := v6
		if uint32(t1) > uint32(v2) {
			p2 = t0
		}
		v3 = p2
		v4 = v4 - v5
		goto l2
	}
l0:
	return i32(0)
}
func (m *Module) fn1439(v0, v1 int32) int32 {
	t0 := m.fn1209(v0+i32(56), v1)
	v1 = t0
	p1 := v0
	if v1 != 0 {
		p1 = v1
	}
	return p1
}
func (m *Module) fn1440(v0, v1, v2 int32) {
	var v3, v4 int32
	v3 = i32(0)
	{
		t0 := int32(load32(m.memory[int64(uint32(v1))+372:]))
		if uint32(v2) >= uint32(t0) {
			goto l1
		}
		t1 := int32(load32(m.memory[int64(uint32(v1))+368:]))
		v4 = t1 + v2<<3
		t2 := int32(load32(m.memory[uint32(v4):]))
		if t2 != i32(1) {
			goto l1
		}
		t3 := int32(load32(m.memory[int64(uint32(v4))+4:]))
		v4 = t3
		t4 := int32(load32(m.memory[int64(uint32(v1))+360:]))
		if uint32(v4) >= uint32(t4) {
			goto l1
		}
		t5 := int32(load32(m.memory[int64(uint32(v1))+356:]))
		v1 = t5 + v4*i32(12)
		t6 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		v2 = t6
		t7 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		v3 = t7
		goto l1
	}
l1:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v2))
	store32(m.memory[uint32(v0):], uint32(v3))
}
func (m *Module) fn1441(v0 int32) {
	var v1 int32
	var v2, v3 int64
	var v4, v5 int32
	t0 := m.g0
	v1 = t0 - i32(80)
	m.g0 = v1
	{
		t1 := int32(load32(m.memory[int64(uint32(v0))+32:]))
		if t1 == 0 {
			goto l0
		}
		t2 := int64(load64(m.memory[int64(uint32(v0))+32:]))
		v2 = t2
		store32(m.memory[int64(uint32(v0))+32:], uint32(i32(0)))
		t3 := int64(load64(m.memory[int64(uint32(v0))+24:]))
		v3 = t3
		store64(m.memory[int64(uint32(v0))+24:], uint64(i64(0x100000000)))
		store32(m.memory[int64(uint32(v1))+8:], uint32(i32(3)))
		store64(m.memory[int64(uint32(v1))+20:], uint64(v2))
		store64(m.memory[int64(uint32(v1))+12:], uint64(v3))
		{
			t4 := int32(load32(m.memory[int64(uint32(v0))+20:]))
			v4 = t4
			if v4 == 0 {
				goto l1
			}
			t5 := int32(load32(m.memory[int64(uint32(v0))+16:]))
			v4 = t5 + v4*i32(28)
			v5 = v4 + i32(-28)
			if v5 == 0 {
				goto l1
			}
			{
				t6 := int32(m.memory[uint32(v4+i32(-4))])
				if t6 != 0 {
					m.fn1340(v4+i32(-16), v1+i32(8))
					goto l0
				}
				t7 := int32(load32(m.memory[int64(uint32(v1))+32:]))
				store32(m.memory[int64(uint32(v1))+72:], uint32(t7))
				t8 := int64(load64(m.memory[int64(uint32(v1))+24:]))
				store64(m.memory[int64(uint32(v1))+64:], uint64(t8))
				t9 := int64(load64(m.memory[int64(uint32(v1))+16:]))
				store64(m.memory[int64(uint32(v1))+56:], uint64(t9))
				t10 := int64(load64(m.memory[int64(uint32(v1))+8:]))
				store64(m.memory[int64(uint32(v1))+48:], uint64(t10))
				m.fn45(v1+i32(36), v1+i32(48), i32(1))
				t11 := int32(load32(m.memory[int64(uint32(v1))+40:]))
				t12 := v5
				v0 = t11
				t13 := int32(load32(m.memory[int64(uint32(v1))+44:]))
				m.fn75(t12, v0, t13)
				t14 := int32(load32(m.memory[int64(uint32(v1))+36:]))
				m.fn16(t14, v0)
				m.fn893(v1 + i32(8))
				goto l0
			}
		}
	l1:
		m.fn1340(v0, v1+i32(8))
	}
l0:
	m.g0 = v1 + i32(80)
}
func (m *Module) fn1442(v0 int32) {
	var v1, v2, v3, v4, v5, v6, v7, v8, v9 int32
	t0 := m.g0
	v1 = t0 - i32(64)
	m.g0 = v1
	m.fn1441(v0)
	{
		t1 := int32(load32(m.memory[int64(uint32(v0))+20:]))
		v2 = t1
		if v2 == 0 {
			goto l0
		}
		t2 := v0
		v2 = v2 + i32(-1)
		store32(m.memory[int64(uint32(t2))+20:], uint32(v2))
		t3 := int32(load32(m.memory[int64(uint32(v0))+16:]))
		v2 = t3 + v2*i32(28)
		t4 := int32(load32(m.memory[uint32(v2):]))
		v3 = t4
		if v3 == i32(-1) {
			goto l0
		}
		t5 := int32(load32(m.memory[int64(uint32(v2))+8:]))
		v4 = t5
		t6 := int32(load32(m.memory[int64(uint32(v2))+4:]))
		v5 = t6
		t7 := int32(load32(m.memory[int64(uint32(v2))+20:]))
		store32(m.memory[int64(uint32(v1))+24:], uint32(t7))
		t8 := int64(load64(m.memory[int64(uint32(v2))+12:]))
		store64(m.memory[int64(uint32(v1))+16:], uint64(t8))
		m.fn1403(v1+i32(4), v5, v4, v1+i32(16))
		t9 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		v6 = t9
		t10 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		t11 := v1
		v2 = t10
		t12 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		t13 := v2
		v4 = t12 * i32(28)
		v7 = t13 + v4
		store32(m.memory[int64(uint32(t11))+28:], uint32(v7))
		store32(m.memory[int64(uint32(v1))+24:], uint32(v6))
		store32(m.memory[int64(uint32(v1))+16:], uint32(v2))
		v8 = v2 + i32(28)
		v6 = v1 + i32(40)
	l6:
		{
			if v4 == 0 {
				goto l1
			}
			t14 := int32(load32(m.memory[uint32(v2):]))
			v9 = t14
			if v9 != i32(-1) {
				t15 := int64(load64(m.memory[int64(uint32(v2))+4:]))
				store64(m.memory[uint32(v6):], uint64(t15))
				t16 := int64(load64(m.memory[int64(uint32(v2))+12:]))
				store64(m.memory[int64(uint32(v6))+8:], uint64(t16))
				t17 := int64(load64(m.memory[int64(uint32(v2))+20:]))
				store64(m.memory[int64(uint32(v6))+16:], uint64(t17))
				store32(m.memory[int64(uint32(v1))+36:], uint32(v9))
				{
					t18 := int32(load32(m.memory[int64(uint32(v0))+20:]))
					v9 = t18
					if v9 == 0 {
						goto l3
					}
					t19 := int32(load32(m.memory[int64(uint32(v0))+16:]))
					v9 = t19 + v9*i32(28)
					if v9+i32(-28) == 0 {
						goto l3
					}
					{
						t20 := int32(m.memory[uint32(v9+i32(-4))])
						if t20 != 0 {
							m.fn1340(v9+i32(-16), v1+i32(36))
							goto l5
						}
						m.fn893(v1 + i32(36))
						goto l5
					}
				}
			l3:
				m.fn1340(v0, v1+i32(36))
			l5:
				v2 = v2 + i32(28)
				v4 = v4 + i32(-28)
				v8 = v8 + i32(28)
				goto l6
			}
			v7 = v8
		}
	l1:
		store32(m.memory[int64(uint32(v1))+20:], uint32(v7))
		m.fn900(v1 + i32(16))
		m.fn16(v3, v5)
		goto l0
	}
l0:
	m.g0 = v1 + i32(64)
}
func (m *Module) fn1443(v0, v1, v2 int32) {
	var v3 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		v3 = t0
		t1 := int32(load32(m.memory[uint32(v0):]))
		if v3 != t1 {
			goto l0
		}
		m.fn625(v0)
	}
l0:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v3+i32(1)))
	t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	v0 = t2 + v3<<3
	store32(m.memory[int64(uint32(v0))+4:], uint32(v2))
	store32(m.memory[uint32(v0):], uint32(v1))
}
func (m *Module) fn1444(v0 int32) {
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
		m.fn894(v3 + i32(12))
		v1 = v1 + i32(-1)
		v3 = v3 + i32(28)
		goto l1
	}
l0:
	t4 := int32(load32(m.memory[uint32(v0):]))
	m.fn136(t4, v2, i32(4), i32(28))
}
func (m *Module) fn1445(v0, v1, v2, v3 int32) {
	var v4, v5, v6 int32
	t0 := m.g0
	v4 = t0 - i32(32)
	m.g0 = v4
	v5 = i32(0)
	{
		t1 := int32(load32(m.memory[uint32(v0):]))
		v6 = t1
		switch v6 {
		case 2:
			v5 = i32(1)
			fallthrough
		case 1:
			if v5 == v1 {
				goto l3
			}
			fallthrough
		default:
			m.fn1333(v0, v3)
			m.fn1332(v0)
			store32(m.memory[int64(uint32(v0))+12:], uint32(i32(0)))
			t3 := v0
			p2 := i32(8)
			if v1 != 0 {
				p2 = i32(4)
			}
			store32(m.memory[int64(uint32(t3))+8:], uint32(p2))
			store32(m.memory[int64(uint32(v0))+4:], uint32(i32(0)))
			t5 := v0
			p4 := i32(1)
			if v1 != 0 {
				p4 = i32(2)
			}
			v6 = p4
			store32(m.memory[uint32(t5):], uint32(v6))
		}
	}
l3:
	{
		if v6 == i32(2) {
			t9 := int32(load32(m.memory[int64(uint32(v2))+4:]))
			t10 := int32(load32(m.memory[int64(uint32(v2))+8:]))
			m.fn45(v4, t9, t10)
			m.fn33(v0+i32(4), v4)
			goto l6
		}
		t6 := int32(load32(m.memory[int64(uint32(v2))+4:]))
		t7 := int32(load32(m.memory[int64(uint32(v2))+8:]))
		t8 := m.fn23(t6, t7)
		if t8 == 0 {
			store32(m.memory[uint32(v4):], uint32(i32(-0x80000000)))
			t11 := int32(load32(m.memory[int64(uint32(v2))+8:]))
			store32(m.memory[int64(uint32(v4))+12:], uint32(t11))
			t12 := int64(load64(m.memory[uint32(v2):]))
			store64(m.memory[int64(uint32(v4))+4:], uint64(t12))
			m.fn338(v0+i32(4), v4)
			goto l7
		}
		goto l6
	}
l6:
	m.fn894(v2)
l7:
	m.g0 = v4 + i32(32)
}
func (m *Module) fn1446(v0, v1 int32) int32 {
	var v2 int64
	var v3, v4 int32
	var v5 int64
	var v6, v7 int32
	var v8 int64
	var v9 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
		if t0 != 0 {
			t1 := int64(load64(m.memory[int64(uint32(v0))+16:]))
			t2 := int64(load64(m.memory[int64(uint32(v0))+24:]))
			t3 := m.fn529(t1, t2, v1)
			v2 = t3
			t4 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			v3 = t4
			v4 = v3 & int32(v2)
			v5 = int64(uint64(v2)>>25) & i64(127) * i64(72340172838076673)
			t5 := int32(load32(m.memory[uint32(v0):]))
			v0 = t5
			v6 = v1 & i32(0xffff)
			v7 = i32(0)
			var _ int32
		l5:
			{
				t7 := int64(load64(m.memory[uint32(v0+v4):]))
				v8 = t7
				v2 = v8 ^ v5
				v2 = (v2 ^ i64(-1)) & (v2 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
				{
				l3:
					{
						var p8 int32
						if v2 == 0 {
							p8 = 1
						}
						v1 = p8
						if v1 != 0 {
							goto l1
						}
						t9 := v6
						t10 := v0
						v9 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v2))))>>3) + v4) & v3
						t11 := int32(load16(m.memory[uint32(t10+(i32(0)-v9)*i32(520)+i32(-520)):]))
						if t9 == t11 {
							goto l2
						}
						v2 = (v2 + i64(-1)) & v2
						goto l3
					}
				l1:
					if v8&(v8<<1)&i64(-0x7f7f7f7f7f7f7f80) == 0 {
						t14 := v4
						v7 = v7 + i32(8)
						v4 = (t14 + v7) & v3
						goto l5
					}
				l2:
					p12 := v0 + (i32(0)-v9)*i32(520)
					if v1 != 0 {
						p12 = i32(0)
					}
					p13 := p12 + i32(-512)
					if v1 != 0 {
						p13 = i32(0)
					}
					return p13
				}
			}
		}
		return i32(0)
	}
}
func (m *Module) fn1447(v0, v1, v2, v3, v4 int32) {
	var v5, v6 int32
	var v7 int64
	var v8, v9, v10, v11, v12 int32
	var v13 int64
	var v14, v15 int32
	var v16, v17 int64
	t0 := m.g0
	v5 = t0 - i32(112)
	m.g0 = v5
	store16(m.memory[int64(uint32(v5))+96:], uint16(v2))
	t2 := v5
	p1 := i32(8)
	if uint32(v4) < uint32(i32(8)) {
		p1 = v4
	}
	v6 = p1
	store32(m.memory[int64(uint32(t2))+100:], uint32(v6))
	v4 = v1 + i32(32)
	t3 := int64(load64(m.memory[int64(uint32(v1))+48:]))
	t4 := int64(load64(m.memory[int64(uint32(v1))+56:]))
	t5 := m.fn703(t3, t4, v2, v6)
	v7 = t5
	store32(m.memory[int64(uint32(v5))+108:], uint32(v5+i32(96)))
	{
		t6 := int32(load32(m.memory[int64(uint32(v1))+40:]))
		if t6 != 0 {
			goto l0
		}
		_ = m.fn701(v4, v1+i32(48))
	}
l0:
	store32(m.memory[int64(uint32(v5))+12:], uint32(v4))
	store32(m.memory[int64(uint32(v5))+8:], uint32(v5+i32(108)))
	t8 := int32(load32(m.memory[int64(uint32(v1))+32:]))
	t9 := int32(load32(m.memory[int64(uint32(v1))+36:]))
	m.fn69(v5, t8, t9, v7, v5+i32(8), i32(189))
	{
		t10 := int32(load32(m.memory[uint32(v5):]))
		v8 = t10
		if v8 != i32(1) {
			goto l1
		}
		t11 := int32(load32(m.memory[int64(uint32(v1))+32:]))
		v4 = t11
		t12 := int32(load32(m.memory[int64(uint32(v5))+4:]))
		t13 := v4
		v9 = t12
		v10 = t13 + v9
		t14 := int32(m.memory[uint32(v10)])
		v11 = t14
		t15 := v10
		v12 = int32(uint32(int32(v7)) >> 25)
		m.memory[uint32(t15)] = byte(v12)
		t16 := int32(load32(m.memory[int64(uint32(v1))+36:]))
		m.memory[uint32(v4+t16&(v9+i32(-8))+i32(8))] = byte(v12)
		t17 := int32(load32(m.memory[int64(uint32(v1))+44:]))
		store32(m.memory[int64(uint32(v1))+44:], uint32(t17+i32(1)))
		t18 := int32(load32(m.memory[int64(uint32(v1))+40:]))
		store32(m.memory[int64(uint32(v1))+40:], uint32(t18-v11&i32(1)))
		v4 = v4 - v9<<3
		store16(m.memory[uint32(v4+i32(-8)):], uint16(v2))
		store32(m.memory[uint32(v4+i32(-4)):], uint32(v6))
	}
l1:
	t19 := int64(load64(m.memory[int64(uint32(v1))+16:]))
	t20 := int64(load64(m.memory[int64(uint32(v1))+24:]))
	t21 := int32(load32(m.memory[int64(uint32(v3))+504:]))
	v11 = t21
	t22 := m.fn66(t19, t20, v11)
	v7 = t22
	t23 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	v10 = t23
	t24 := v10
	v12 = int32(v7)
	v9 = t24 & v12
	v13 = int64(uint64(v7)>>25) & i64(127) * i64(72340172838076673)
	v14 = v1 + i32(16)
	t25 := int32(load32(m.memory[uint32(v1):]))
	v4 = t25
	v15 = i32(0)
l24:
	{
		t26 := int64(load64(m.memory[uint32(v4+v9):]))
		v16 = t26
		v17 = v16 ^ v13
		v17 = (v17 ^ i64(-1)) & (v17 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
		{
		l4:
			{
				if v17 == 0 {
					goto l2
				}
				v2 = v4 + (i32(0)-(int32(uint32(int64(bits.TrailingZeros64(uint64(v17))))>>3)+v9)&v10)*i32(96)
				t27 := int32(load32(m.memory[uint32(v2+i32(-96)):]))
				if t27 == v11 {
					goto l3
				}
				v17 = (v17 + i64(-1)) & v17
				goto l4
			}
		l2:
			if v16&(v16<<1)&i64(-0x7f7f7f7f7f7f7f80) == 0 {
				t76 := v9
				v15 = v15 + i32(8)
				v9 = (t76 + v15) & v10
				goto l24
			}
			{
				t28 := int32(load32(m.memory[int64(uint32(v1))+8:]))
				v9 = t28
				if v9 != 0 {
					goto l6
				}
				_ = m.fn721(v1, v14)
				t30 := int32(load32(m.memory[int64(uint32(v1))+8:]))
				v9 = t30
				t31 := int32(load32(m.memory[int64(uint32(v1))+4:]))
				v10 = t31
				t32 := int32(load32(m.memory[uint32(v1):]))
				v4 = t32
			}
		l6:
			memory_zero(m.memory, uint32(v5+i32(12)), uint32(i32(81)))
			t33 := m.fn26(v4, v10, v7)
			t34 := v4
			v2 = t33
			v15 = t34 + v2
			t35 := int32(m.memory[uint32(v15)])
			v14 = t35
			t36 := v15
			v12 = int32(uint32(v12) >> 25)
			m.memory[uint32(t36)] = byte(v12)
			m.memory[uint32(v4+v10&(v2+i32(-8))+i32(8))] = byte(v12)
			store32(m.memory[int64(uint32(v1))+8:], uint32(v9-v14&i32(1)))
			v2 = v4 + (i32(0)-v2)*i32(96)
			store32(m.memory[uint32(v2+i32(-96)):], uint32(v11))
			memory_copy(m.memory, uint32(v2+i32(-92)), uint32(v5+i32(8)), uint32(i32(85)))
			t37 := int32(load32(m.memory[int64(uint32(v1))+12:]))
			store32(m.memory[int64(uint32(v1))+12:], uint32(t37+i32(1)))
		}
	l3:
		v12 = v2 + i32(-88)
		{
			{
				if v8 != i32(1) {
					goto l7
				}
				v1 = v3 + v6<<4
				t38 := int32(load32(m.memory[uint32(v1):]))
				if t38 == 0 {
					goto l7
				}
				t39 := int64(load64(m.memory[int64(uint32(v1))+8:]))
				v7 = t39
				goto l8
			}
		l7:
			{
				t40 := int32(m.memory[uint32(v2+v6+i32(-16))])
				if t40 != 0 {
					goto l9
				}
				t41 := int64(load64(m.memory[int64(uint32(v3+v6*i32(40)))+168:]))
				v7 = t41
				goto l8
			}
		l9:
			t42 := int64(load64(m.memory[uint32(v12+v6<<3):]))
			v7 = t42 + i64(1)
			p43 := v7
			if v7 == 0 {
				p43 = i64(-1)
			}
			v7 = p43
		}
	l8:
		v8 = v2 + i32(-16)
		m.memory[uint32(v8+v6)] = byte(i32(1))
		store64(m.memory[uint32(v12+v6<<3):], uint64(v7))
		v4 = v6 + i32(1)
		v1 = i32(144)
		v15 = v3 + i32(144)
		v10 = i32(0)
	l14:
		{
			if v4 != 0 {
				goto l10
			}
			v2 = v10
			if v1 != i32(504) {
				goto l11
			}
			goto l12
		l10:
			t44 := int32(uint32((i32(504)-v1)&i32(0xffff)) / uint32(i32(40)))
			if uint32(v4) >= uint32(t44) {
				goto l12
			}
			v2 = v4 + v10
			v1 = v4*i32(40) + v1
		}
	l11:
		v10 = v2 + i32(1)
		v9 = v1 + i32(40)
		{
			v11 = v3 + v1
			t45 := int32(load32(m.memory[uint32(v11):]))
			if t45 == 0 {
				goto l13
			}
			v4 = i32(0)
			v1 = v9
			t46 := int32(load32(m.memory[int64(uint32(v11))+4:]))
			if uint32(v6) >= uint32(t46) {
				goto l14
			}
		}
	l13:
		if uint32(v2) > uint32(i32(8)) {
			goto l15
		}
		v4 = i32(0)
		m.memory[uint32(v8+v2)] = byte(i32(0))
		v1 = v9
		goto l14
	l15:
		m.fn158(v2, i32(9), i32(1077224))
		panic("unreachable")
	l12:
		v1 = i32(-1)
		{
			v4 = v15 + v6*i32(40)
			t47 := int32(m.memory[int64(uint32(v4))+32])
			v11 = t47
			if v11 == i32(255) {
				goto l22
			}
			t48 := int32(load32(m.memory[int64(uint32(v4))+16:]))
			v2 = t48
			if v2 == 0 {
				goto l22
			}
			store32(m.memory[int64(uint32(v5))+104:], uint32(i32(0)))
			store64(m.memory[int64(uint32(v5))+96:], uint64(i64(0x100000000)))
			v2 = v2 * i32(12)
			t49 := int32(m.memory[int64(uint32(v4))+20])
			v10 = t49
			t50 := int32(load32(m.memory[int64(uint32(v4))+12:]))
			v1 = t50
		l23:
			{
				{
					if v2 == 0 {
						m.fn800(v5+i32(8), v11, v7)
						t65 := int32(load32(m.memory[int64(uint32(v5))+100:]))
						v4 = t65
						t66 := int32(load32(m.memory[int64(uint32(v5))+104:]))
						t67 := int32(load32(m.memory[int64(uint32(v5))+12:]))
						t68 := v4
						v1 = t67
						t69 := int32(load32(m.memory[int64(uint32(v5))+16:]))
						t70 := m.fn191(t68, t66, v1, t69)
						v2 = t70
						t71 := int32(load32(m.memory[int64(uint32(v5))+8:]))
						m.fn16(t71, v1)
						t72 := int32(load32(m.memory[int64(uint32(v5))+96:]))
						v1 = t72
						{
							if v2 != 0 {
								m.fn16(v1, v4)
								v1 = i32(-1)
								goto l22
							}
							t73 := int64(load64(m.memory[int64(uint32(v5))+100:]))
							v17 = t73
							goto l22
						}
					}
					t51 := int32(load32(m.memory[uint32(v1):]))
					if t51 != i32(-1) {
						goto l18
					}
					t52 := int32(m.memory[uint32(v1+i32(4))])
					v4 = t52
					p53 := i32(8)
					if uint32(v4) < uint32(i32(8)) {
						p53 = v4
					}
					v4 = p53
					v9 = i32(1)
					{
						if v10&i32(1) != 0 {
							goto l19
						}
						t54 := int32(m.memory[int64(uint32(v3+v4*i32(40)))+176])
						v9 = t54
						p55 := v9
						if v9 == i32(255) {
							p55 = i32(1)
						}
						v9 = p55
					}
				l19:
					t56 := int32(m.memory[uint32(v8+v4)])
					t58 := v5 + i32(8)
					t59 := v9
					p57 := v3 + v4*i32(40) + i32(168)
					if t56 != 0 {
						p57 = v12 + v4<<3
					}
					t60 := int64(load64(m.memory[uint32(p57):]))
					m.fn804(t58, t59, t60)
					t61 := int32(load32(m.memory[int64(uint32(v5))+12:]))
					t62 := v5 + i32(96)
					v4 = t61
					t63 := int32(load32(m.memory[int64(uint32(v5))+16:]))
					m.fn75(t62, v4, t63)
					t64 := int32(load32(m.memory[int64(uint32(v5))+8:]))
					m.fn16(t64, v4)
					goto l20
				}
			l18:
				t74 := int32(load32(m.memory[uint32(v1+i32(4)):]))
				t75 := int32(load32(m.memory[uint32(v1+i32(8)):]))
				m.fn75(v5+i32(96), t74, t75)
			}
		l20:
			v1 = v1 + i32(12)
			v2 = v2 + i32(-12)
			goto l23
		}
	l22:
		store64(m.memory[int64(uint32(v0))+12:], uint64(v17))
		store32(m.memory[int64(uint32(v0))+8:], uint32(v1))
		store64(m.memory[uint32(v0):], uint64(v7))
		m.g0 = v5 + i32(112)
		return
	}
}
func (m *Module) fn1448(v0, v1, v2, v3, v4 int32) {
	var v5 int32
	t0 := m.g0
	v5 = t0 - i32(32)
	m.g0 = v5
	{
		t1 := m.fn1439(v0, v1)
		t2 := int32(m.memory[int64(uint32(t1))+54])
		v1 = t2
		if v1 == i32(2) {
			goto l0
		}
		m.fn1445(v4, v1&i32(1), v2, v3)
		goto l1
	}
l0:
	m.fn1333(v4, v3)
	{
		t3 := int32(load32(m.memory[int64(uint32(v2))+4:]))
		t4 := int32(load32(m.memory[int64(uint32(v2))+8:]))
		t5 := m.fn23(t3, t4)
		if t5 != 0 {
			goto l2
		}
		store32(m.memory[uint32(v5):], uint32(i32(-0x80000000)))
		t6 := int32(load32(m.memory[int64(uint32(v2))+8:]))
		store32(m.memory[int64(uint32(v5))+12:], uint32(t6))
		t7 := int64(load64(m.memory[uint32(v2):]))
		store64(m.memory[int64(uint32(v5))+4:], uint64(t7))
		m.fn338(v3, v5)
		goto l1
	}
l2:
	m.fn894(v2)
l1:
	m.g0 = v5 + i32(32)
}
func (m *Module) fn1449(v0, v1 int32) {
	var v2, v3, v4, v5, v6, v7, v8, v9, v10 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	{
		t1 := int32(load32(m.memory[uint32(v1):]))
		if t1 == i32(-1) {
			goto l0
		}
		t2 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		v3 = t2
		t3 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		t4 := v2 + i32(8)
		v4 = t3
		m.fn59(t4, v4, i32(2), i32(2))
		v5 = i32(0)
		t5 := int32(load32(m.memory[int64(uint32(v2))+12:]))
		v6 = t5
		t6 := int32(load32(m.memory[int64(uint32(v2))+8:]))
		v7 = t6
		v8 = i32(0)
		if v4 == 0 {
			goto l1
		}
		v8 = v4 << 1
		if v8 == 0 {
			goto l2
		}
		memory_copy(m.memory, uint32(v6), uint32(v3), uint32(v8))
	l2:
		v8 = v4
	l1:
		t7 := int32(load32(m.memory[int64(uint32(v1))+16:]))
		v9 = t7
		t8 := int32(load32(m.memory[int64(uint32(v1))+20:]))
		t9 := v2
		v4 = t8
		m.fn59(t9, v4, i32(1), i32(4))
		t10 := int32(load32(m.memory[int64(uint32(v2))+4:]))
		v3 = t10
		t11 := int32(load32(m.memory[uint32(v2):]))
		v10 = t11
		if v4 == 0 {
			goto l3
		}
		v5 = v4 << 2
		if v5 == 0 {
			goto l4
		}
		memory_copy(m.memory, uint32(v3), uint32(v9), uint32(v5))
	l4:
		v5 = v4
	l3:
		store32(m.memory[int64(uint32(v0))+20:], uint32(v5))
		store32(m.memory[int64(uint32(v0))+16:], uint32(v3))
		store32(m.memory[int64(uint32(v0))+12:], uint32(v10))
		store32(m.memory[int64(uint32(v0))+8:], uint32(v8))
		store32(m.memory[int64(uint32(v0))+4:], uint32(v6))
		store32(m.memory[uint32(v0):], uint32(v7))
		t12 := int32(m.memory[int64(uint32(v1))+24])
		m.memory[int64(uint32(v0))+24] = byte(t12)
		goto l5
	}
l0:
	store32(m.memory[uint32(v0):], uint32(i32(-1)))
l5:
	m.g0 = v2 + i32(16)
}
func (m *Module) fn1450(v0, v1 int32) {
	var v2 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		v2 = t0
		t1 := int32(load32(m.memory[uint32(v0):]))
		if v2 != t1 {
			goto l0
		}
		m.fn1144(v0)
	}
l0:
	t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	memory_copy(m.memory, uint32(t2+v2*i32(40)), uint32(v1), uint32(i32(40)))
	store32(m.memory[int64(uint32(v0))+8:], uint32(v2+i32(1)))
}
func (m *Module) fn1451(v0 int32) {
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
	m.fn975(v3)
	v3 = v3 + i32(40)
	goto l1
l0:
	t2 := int32(load32(m.memory[uint32(v0):]))
	m.fn84(t2, v2)
}
func (m *Module) fn1452(v0, v1 int32) {
	var v2, v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14, v15, v16 int32
	var v17, v18 int64
	var v19, v20, v21, v22, v23, v24, v25, v26 int32
	t0 := m.g0
	v2 = t0 - i32(288)
	m.g0 = v2
	t1 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	v3 = t1
	t2 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	t3 := v3
	v4 = t2
	v5 = v4 << 4
	v6 = t3 + v5
	v7 = i32(0)
	v8 = v3
l2:
	if v5 != 0 {
		t4 := int32(m.memory[int64(uint32(v8))+12])
		if t4 != i32(1) {
			goto l1
		}
		v5 = v5 + i32(-16)
		v7 = v7 + i32(1)
		v8 = v8 + i32(16)
		goto l2
	}
	v7 = v4
	goto l1
l1:
	t5 := int32(load32(m.memory[uint32(v1):]))
	v9 = t5
	v10 = v9 << 4
	t6 := int32(uint32(v10) / uint32(i32(12)))
	v11 = t6
	v12 = i32(0)
	v13 = v3
	v14 = v3
l57:
	{
		if v13 == v6 {
			m.fn419(i32(0), i32(4))
			v19 = v3
			{
				{
					if v9 == 0 {
						goto l6
					}
					v19 = v3
					t13 := v10
					v5 = v11 * i32(12)
					if t13 == v5 {
						goto l6
					}
					t14 := m.fn392(v3, v10, v5)
					v19 = t14
					if v19 == 0 {
						m.fn85(i32(4), v5)
						panic("unreachable")
					}
				}
			l6:
				m.fn419(i32(0), i32(4))
				store32(m.memory[int64(uint32(v2))+112:], uint32(i32(0)))
				store32(m.memory[int64(uint32(v2))+104:], uint32(i32(0)))
				store32(m.memory[int64(uint32(v2))+96:], uint32(v19))
				t15 := v2
				v20 = v19 + (v14 - v3)
				store32(m.memory[int64(uint32(t15))+100:], uint32(v20))
				m.fn942(v2+i32(232), v2+i32(96))
				{
					{
						t16 := int64(load64(m.memory[int64(uint32(v2))+232:]))
						if t16 != i64(1) {
							v1 = i32(8)
							v13 = i32(0)
							v8 = i32(0)
							goto l12
						}
						t17 := int64(load64(m.memory[int64(uint32(v2))+240:]))
						v17 = t17
						m.fn944(v2+i32(64), v2+i32(96))
						t18 := int32(load32(m.memory[int64(uint32(v2))+64:]))
						t19 := v2 + i32(8)
						v5 = t18 + i32(1)
						p20 := i32(-1)
						if v5 != 0 {
							p20 = v5
						}
						v5 = p20
						p21 := i32(4)
						if uint32(v5) > uint32(i32(4)) {
							p21 = v5
						}
						m.fn59(t19, p21, i32(8), i32(8))
						t22 := int32(load32(m.memory[int64(uint32(v2))+8:]))
						v5 = t22
						t23 := int32(load32(m.memory[int64(uint32(v2))+12:]))
						v1 = t23
						store64(m.memory[uint32(v1):], uint64(v17))
						store32(m.memory[int64(uint32(v2))+216:], uint32(i32(1)))
						store32(m.memory[int64(uint32(v2))+212:], uint32(v1))
						store32(m.memory[int64(uint32(v2))+208:], uint32(v5))
						t24 := int64(load64(m.memory[int64(uint32(v2))+112:]))
						store64(m.memory[int64(uint32(v2))+248:], uint64(t24))
						t25 := int64(load64(m.memory[int64(uint32(v2))+104:]))
						store64(m.memory[int64(uint32(v2))+240:], uint64(t25))
						t26 := int64(load64(m.memory[int64(uint32(v2))+96:]))
						store64(m.memory[int64(uint32(v2))+232:], uint64(t26))
						v5 = i32(8)
						v8 = i32(1)
					l11:
						{
							m.fn942(v2+i32(64), v2+i32(232))
							t27 := int64(load64(m.memory[int64(uint32(v2))+64:]))
							if t27 != i64(1) {
								goto l9
							}
							t28 := int64(load64(m.memory[int64(uint32(v2))+72:]))
							v17 = t28
							{
								t29 := int32(load32(m.memory[int64(uint32(v2))+208:]))
								if v8 != t29 {
									goto l10
								}
								m.fn944(v2+i32(184), v2+i32(232))
								t30 := int32(load32(m.memory[int64(uint32(v2))+184:]))
								t31 := v2 + i32(208)
								t32 := v8
								v1 = t30 + i32(1)
								p33 := i32(-1)
								if v1 != 0 {
									p33 = v1
								}
								m.fn62(t31, t32, p33, i32(8), i32(8))
								t34 := int32(load32(m.memory[int64(uint32(v2))+212:]))
								v1 = t34
							}
						l10:
							store64(m.memory[uint32(v1+v5):], uint64(v17))
							t35 := v2
							v8 = v8 + i32(1)
							store32(m.memory[int64(uint32(t35))+216:], uint32(v8))
							v5 = v5 + i32(8)
							goto l11
						}
					}
				l9:
					t36 := int32(load32(m.memory[int64(uint32(v2))+208:]))
					v13 = t36
					if uint32(v8) < uint32(i32(2)) {
						goto l12
					}
					if uint32(v8) < uint32(i32(21)) {
						goto l13
					}
					m.fn1007(v1, v8)
					goto l12
				l13:
					m.fn1011(v1, v8, i32(1))
				}
			l12:
				v5 = i32(0)
				store32(m.memory[int64(uint32(v2))+24:], uint32(i32(0)))
				store64(m.memory[int64(uint32(v2))+16:], uint64(i64(0x800000000)))
				v15 = v8 << 3
			l17:
				{
					if v15 == v5 {
						m.fn1415(v13, v1)
						v15 = i32(0)
						store32(m.memory[int64(uint32(v2))+36:], uint32(i32(0)))
						store64(m.memory[int64(uint32(v2))+28:], uint64(i64(0x400000000)))
						v21 = v19 + v4*i32(12)
						v13 = i32(4)
						v1 = v2 + i32(232) | i32(4)
						v22 = v2 + i32(232) + i32(32)
						t41 := int32(load32(m.memory[int64(uint32(v2))+24:]))
						v23 = t41
						t42 := int32(load32(m.memory[int64(uint32(v2))+20:]))
						v6 = t42
						v5 = v19
					l44:
						{
							{
								if v5 == v20 {
									goto l18
								}
								v24 = v5 + i32(12)
								t43 := int32(load32(m.memory[uint32(v5):]))
								v8 = t43
								if v8 != i32(-1) {
									v10 = i32(0)
									store32(m.memory[int64(uint32(v2))+48:], uint32(i32(0)))
									store64(m.memory[int64(uint32(v2))+40:], uint64(i64(0x400000000)))
									t92 := int64(load64(m.memory[int64(uint32(v5))+4:]))
									v17 = t92
									store32(m.memory[int64(uint32(v2))+272:], uint32(v8))
									t93 := v2
									v5 = int32(v17)
									store32(m.memory[int64(uint32(t93))+268:], uint32(v5))
									store32(m.memory[int64(uint32(v2))+264:], uint32(v5))
									t94 := v2
									v3 = v5 + int32(int64(uint64(v17)>>32))<<5
									store32(m.memory[int64(uint32(t94))+276:], uint32(v3))
									v13 = i32(-2)
									v26 = i32(4)
									v4 = i32(0)
								l42:
									{
										store32(m.memory[int64(uint32(v2))+232:], uint32(i32(-2)))
										{
											{
												if v13 == i32(-2) {
													goto l29
												}
												t95 := int32(m.memory[int64(uint32(v2))+256])
												v8 = t95
												t96 := int32(m.memory[int64(uint32(v2))+259])
												v9 = t96
												t97 := int32(m.memory[int64(uint32(v2))+258])
												v25 = t97
												t98 := int64(load64(m.memory[int64(uint32(v2))+248:]))
												v17 = t98
												t99 := int64(load64(m.memory[int64(uint32(v2))+236:]))
												v18 = t99
												goto l30
											}
										l29:
											if v5 == v3 {
												goto l31
											}
											t100 := v2
											v15 = v5 + i32(32)
											store32(m.memory[int64(uint32(t100))+268:], uint32(v15))
											t101 := int32(m.memory[int64(uint32(v5))+24])
											v8 = t101
											t102 := int32(m.memory[int64(uint32(v5))+27])
											v9 = t102
											t103 := int32(m.memory[int64(uint32(v5))+26])
											v25 = t103
											t104 := int64(load64(m.memory[int64(uint32(v5))+16:]))
											v17 = t104
											t105 := int64(load64(m.memory[int64(uint32(v5))+4:]))
											v18 = t105
											t106 := int32(load32(m.memory[uint32(v5):]))
											v13 = t106
											v5 = v15
										}
									l30:
										if v13 == i32(-1) {
											goto l31
										}
										store64(m.memory[int64(uint32(v2))+56:], uint64(v18))
										store32(m.memory[int64(uint32(v2))+52:], uint32(v13))
										v13 = i32(-2)
										if v8&i32(1) == 0 {
											goto l32
										}
									l36:
										{
											{
												if v5 != v3 {
													goto l33
												}
												v13 = i32(-1)
												goto l34
											l33:
												t107 := int32(load32(m.memory[uint32(v5):]))
												v13 = t107
												t108 := int32(load32(m.memory[int64(uint32(v5))+28:]))
												store32(m.memory[int64(uint32(v2))+120:], uint32(t108))
												t109 := int64(load64(m.memory[int64(uint32(v5))+20:]))
												store64(m.memory[int64(uint32(v2))+112:], uint64(t109))
												t110 := int64(load64(m.memory[int64(uint32(v5))+12:]))
												store64(m.memory[int64(uint32(v2))+104:], uint64(t110))
												t111 := int64(load64(m.memory[int64(uint32(v5))+4:]))
												store64(m.memory[int64(uint32(v2))+96:], uint64(t111))
												v5 = v5 + i32(32)
											}
										l34:
											t112 := int32(load32(m.memory[int64(uint32(v2))+120:]))
											store32(m.memory[int64(uint32(v1))+24:], uint32(t112))
											t113 := int64(load64(m.memory[int64(uint32(v2))+112:]))
											store64(m.memory[int64(uint32(v1))+16:], uint64(t113))
											t114 := int64(load64(m.memory[int64(uint32(v2))+104:]))
											store64(m.memory[int64(uint32(v1))+8:], uint64(t114))
											t115 := int64(load64(m.memory[int64(uint32(v2))+96:]))
											store64(m.memory[uint32(v1):], uint64(t115))
											{
												if v13 == i32(-1) {
													goto l35
												}
												t116 := int32(m.memory[int64(uint32(v2))+257])
												if t116&i32(1) == 0 {
													goto l35
												}
												t117 := int64(load64(m.memory[int64(uint32(v2))+248:]))
												v17 = t117
												t118 := int64(load64(m.memory[int64(uint32(v2))+236:]))
												store64(m.memory[int64(uint32(v2))+100:], uint64(t118))
												store32(m.memory[int64(uint32(v2))+96:], uint32(v13))
												m.fn1271(v2+i32(52), v2+i32(96))
												goto l36
											}
										l35:
										}
										store32(m.memory[int64(uint32(v2))+232:], uint32(v13))
										store32(m.memory[int64(uint32(v2))+268:], uint32(v5))
									l32:
										{
											if v23 != 0 {
												goto l37
											}
											v8 = i32(1)
											goto l38
										l37:
											v17 = v17 + i64(-10)
											v15 = i32(0)
											v8 = v23
										l40:
											{
												if uint32(v8) < uint32(i32(2)) {
													goto l39
												}
												v14 = int32(uint32(v8) >> 1)
												v16 = v14 + v15
												t119 := int64(load64(m.memory[uint32(v6+v16<<3):]))
												p120 := v15
												if t119 < v17 {
													p120 = v16
												}
												v15 = p120
												v8 = v8 - v14
												goto l40
											}
										l39:
											t121 := int64(load64(m.memory[uint32(v6+v15<<3):]))
											t122 := v15
											var p123 int32
											if t121 < v17 {
												p123 = 1
											}
											v8 = t122 + p123 + i32(1)
										}
									l38:
										v15 = v10 + i32(1)
										p124 := v8
										if uint32(v15) > uint32(v8) {
											p124 = v15
										}
										v15 = p124
										{
											t125 := int32(load32(m.memory[int64(uint32(v2))+40:]))
											if v4 != t125 {
												goto l41
											}
											m.fn1143(v2 + i32(40))
											t126 := int32(load32(m.memory[int64(uint32(v2))+44:]))
											v26 = t126
										}
									l41:
										t127 := int64(load64(m.memory[int64(uint32(v2))+52:]))
										v17 = t127
										v8 = v26 + v4*i32(28)
										t128 := int32(load32(m.memory[int64(uint32(v2))+60:]))
										store32(m.memory[int64(uint32(v8))+8:], uint32(t128))
										store64(m.memory[uint32(v8):], uint64(v17))
										m.memory[int64(uint32(v8))+25] = byte(v9)
										m.memory[int64(uint32(v8))+24] = byte(v25)
										store32(m.memory[int64(uint32(v8))+20:], uint32(i32(1)))
										store32(m.memory[int64(uint32(v8))+16:], uint32(v15))
										store32(m.memory[int64(uint32(v8))+12:], uint32(v10))
										t129 := v2
										v4 = v4 + i32(1)
										store32(m.memory[int64(uint32(t129))+48:], uint32(v4))
										v10 = v15
										goto l42
									}
								l31:
									{
										t130 := int32(load32(m.memory[int64(uint32(v2))+36:]))
										v5 = t130
										t131 := int32(load32(m.memory[int64(uint32(v2))+28:]))
										if v5 != t131 {
											goto l43
										}
										m.fn272(v2 + i32(28))
									}
								l43:
									t132 := int32(load32(m.memory[int64(uint32(v2))+32:]))
									v13 = t132
									v8 = v13 + v5*i32(12)
									t133 := int64(load64(m.memory[int64(uint32(v2))+40:]))
									store64(m.memory[uint32(v8):], uint64(t133))
									t134 := int32(load32(m.memory[int64(uint32(v2))+48:]))
									store32(m.memory[int64(uint32(v8))+8:], uint32(t134))
									t135 := v2
									v15 = v5 + i32(1)
									store32(m.memory[int64(uint32(t135))+36:], uint32(v15))
									m.fn974(v22)
									v5 = v24
									goto l44
								}
								v21 = v24
							}
						l18:
							t44 := int32(uint32(v19-v21+v12) / uint32(i32(12)))
							v5 = t44
						l21:
							if v5 == 0 {
								m.fn136(v11, v19, i32(4), i32(12))
								m.fn27(v2 + i32(64))
								v8 = i32(0)
							l27:
								{
									if v8 == v15 {
										m.fn1165(v2 + i32(96))
										t136 := int32(load32(m.memory[int64(uint32(v2))+28:]))
										store32(m.memory[int64(uint32(v2))+160:], uint32(t136))
										store32(m.memory[int64(uint32(v2))+152:], uint32(v13))
										t137 := v2
										v16 = v13 + v15*i32(12)
										store32(m.memory[int64(uint32(t137))+164:], uint32(v16))
										v8 = v2 + i32(236)
									l51:
										{
											if v13 == v16 {
												goto l45
											}
											v14 = v13 + i32(12)
											t138 := int32(load32(m.memory[uint32(v13):]))
											v5 = t138
											if v5 != i32(-1) {
												t145 := int64(load64(m.memory[int64(uint32(v13))+4:]))
												v17 = t145
												m.fn1166(v2 + i32(96))
												store32(m.memory[int64(uint32(v2))+176:], uint32(v5))
												t146 := v2
												v5 = int32(v17)
												store32(m.memory[int64(uint32(t146))+168:], uint32(v5))
												t147 := v2
												v13 = v5 + int32(int64(uint64(v17)>>32))*i32(28)
												store32(m.memory[int64(uint32(t147))+180:], uint32(v13))
											l56:
												{
													if v5 == v13 {
														goto l49
													}
													v1 = v5 + i32(28)
													t148 := int32(load32(m.memory[uint32(v5):]))
													v15 = t148
													if v15 != i32(-1) {
														t149 := int64(load64(m.memory[int64(uint32(v5))+20:]))
														store64(m.memory[int64(uint32(v8))+16:], uint64(t149))
														t150 := int64(load64(m.memory[int64(uint32(v5))+12:]))
														store64(m.memory[int64(uint32(v8))+8:], uint64(t150))
														t151 := int64(load64(m.memory[int64(uint32(v5))+4:]))
														store64(m.memory[uint32(v8):], uint64(t151))
														store32(m.memory[int64(uint32(v2))+232:], uint32(v15))
														t152 := int32(load32(m.memory[int64(uint32(v2))+248:]))
														t153 := int32(load32(m.memory[int64(uint32(v2))+244:]))
														v5 = t152 - t153
														{
															t154 := int32(m.memory[int64(uint32(v2))+257])
															if t154 == i32(1) {
																goto l55
															}
															t155 := int32(load32(m.memory[int64(uint32(v2))+240:]))
															store32(m.memory[int64(uint32(v2))+216:], uint32(t155))
															t156 := int64(load64(m.memory[int64(uint32(v2))+232:]))
															store64(m.memory[int64(uint32(v2))+208:], uint64(t156))
															t157 := int32(load32(m.memory[int64(uint32(v2))+252:]))
															t158 := v2
															v15 = t157
															p159 := i32(1)
															if uint32(v15) > uint32(i32(1)) {
																p159 = v15
															}
															store32(m.memory[int64(uint32(t158))+224:], uint32(p159))
															t161 := v2
															p160 := i32(1)
															if uint32(v5) > uint32(i32(1)) {
																p160 = v5
															}
															store32(m.memory[int64(uint32(t161))+220:], uint32(p160))
															m.fn1167(v2+i32(184), v2+i32(96), v2+i32(208))
															t162 := int32(load32(m.memory[int64(uint32(v2))+184:]))
															v5 = t162
															if v5 == i32(-1) {
																goto l53
															}
															t163 := int64(load64(m.memory[int64(uint32(v2))+188:]))
															store64(m.memory[int64(uint32(v0))+8:], uint64(t163))
															t164 := int64(load64(m.memory[int64(uint32(v2))+196:]))
															store64(m.memory[int64(uint32(v0))+16:], uint64(t164))
															t165 := int32(load32(m.memory[int64(uint32(v2))+204:]))
															store32(m.memory[int64(uint32(v0))+24:], uint32(t165))
															store32(m.memory[int64(uint32(v2))+172:], uint32(v1))
															store32(m.memory[int64(uint32(v2))+156:], uint32(v14))
															store32(m.memory[uint32(v0):], uint32(i32(-2)))
															store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
															m.fn1461(v2 + i32(168))
															m.fn1460(v2 + i32(152))
															m.fn1259(v2 + i32(96))
															goto l48
														}
													l55:
														if v5 == 0 {
															goto l54
														}
														v5 = v5 + i32(-1)
														_ = m.fn1260(v2 + i32(96))
														goto l55
													l54:
														m.fn969(v2 + i32(232))
													l53:
														v5 = v1
														goto l56
													}
													v13 = v1
												}
											l49:
												store32(m.memory[int64(uint32(v2))+172:], uint32(v13))
												m.fn1461(v2 + i32(168))
												v13 = v14
												goto l51
											}
											v13 = v14
										}
									l45:
										store32(m.memory[int64(uint32(v2))+156:], uint32(v13))
										m.fn1460(v2 + i32(152))
										memory_copy(m.memory, uint32(v2+i32(232)), uint32(v2+i32(96)), uint32(i32(56)))
										m.fn1168(v2+i32(184), v2+i32(232))
										{
											t139 := int32(load32(m.memory[int64(uint32(v2))+192:]))
											v5 = t139
											if v5 == 0 {
												store32(m.memory[uint32(v0):], uint32(i32(-1)))
												m.fn972(v2 + i32(184))
												goto l48
											}
											t140 := int32(load32(m.memory[int64(uint32(v2))+188:]))
											t141 := m.fn1234(t140, v5, v7)
											v5 = t141
											store32(m.memory[uint32(v0):], uint32(i32(-0x7ffffffe)))
											t142 := int64(load64(m.memory[int64(uint32(v2))+184:]))
											store64(m.memory[int64(uint32(v0))+4:], uint64(t142))
											t143 := int32(load32(m.memory[int64(uint32(v2))+200:]))
											store32(m.memory[int64(uint32(v0))+20:], uint32(t143))
											store32(m.memory[int64(uint32(v2))+196:], uint32(v5))
											t144 := int64(load64(m.memory[int64(uint32(v2))+192:]))
											store64(m.memory[int64(uint32(v0))+12:], uint64(t144))
											goto l48
										}
									l48:
										t167 := int32(load32(m.memory[int64(uint32(v2))+64:]))
										t168 := int32(load32(m.memory[int64(uint32(v2))+68:]))
										m.fn1174(t167, t168)
										t169 := int32(load32(m.memory[int64(uint32(v2))+16:]))
										t170 := int32(load32(m.memory[int64(uint32(v2))+20:]))
										m.fn1415(t169, t170)
										m.g0 = v2 + i32(288)
										return
									}
									m.fn27(v2 + i32(232))
									t45 := m.fn857(v13, v15, v8, i32(1080708))
									t46 := int32(load32(m.memory[int64(uint32(t45))+8:]))
									v6 = t46
									v5 = i32(0)
									t47 := int32(load32(m.memory[int64(uint32(v2))+68:]))
									v9 = t47
									t48 := int32(load32(m.memory[int64(uint32(v2))+64:]))
									v25 = t48
									t49 := int32(load32(m.memory[int64(uint32(v2))+76:]))
									v4 = t49
								l28:
									{
										if v6 == v5 {
											t88 := int64(load64(m.memory[int64(uint32(v2))+256:]))
											store64(m.memory[int64(uint32(v2))+88:], uint64(t88))
											t89 := int64(load64(m.memory[int64(uint32(v2))+248:]))
											store64(m.memory[int64(uint32(v2))+80:], uint64(t89))
											t90 := int64(load64(m.memory[int64(uint32(v2))+240:]))
											store64(m.memory[int64(uint32(v2))+72:], uint64(t90))
											t91 := int64(load64(m.memory[int64(uint32(v2))+232:]))
											store64(m.memory[int64(uint32(v2))+64:], uint64(t91))
											m.fn1174(v25, v9)
											v8 = v8 + i32(1)
											goto l27
										}
										t50 := m.fn857(v13, v15, v8, i32(1080724))
										v1 = t50
										t51 := int32(load32(m.memory[uint32(v1+i32(4)):]))
										t52 := int32(load32(m.memory[uint32(v1+i32(8)):]))
										t53 := m.fn1459(t51, t52, v5, i32(1080740))
										t54 := int32(load32(m.memory[int64(uint32(t53))+12:]))
										v1 = t54
										t55 := m.fn857(v13, v15, v8, i32(1080756))
										v14 = t55
										t56 := int32(load32(m.memory[uint32(v14+i32(4)):]))
										t57 := int32(load32(m.memory[uint32(v14+i32(8)):]))
										t58 := m.fn1459(t56, t57, v5, i32(1080772))
										v14 = t58
										store32(m.memory[int64(uint32(v2))+184:], uint32(v1))
										t59 := int32(load32(m.memory[int64(uint32(v14))+16:]))
										t60 := v2
										v16 = t59
										store32(m.memory[int64(uint32(t60))+188:], uint32(v16))
										{
											t61 := m.fn857(v13, v15, v8, i32(1080788))
											v14 = t61
											t62 := int32(load32(m.memory[uint32(v14+i32(4)):]))
											t63 := int32(load32(m.memory[uint32(v14+i32(8)):]))
											t64 := m.fn1459(t62, t63, v5, i32(1080804))
											t65 := int32(m.memory[int64(uint32(t64))+25])
											if t65 == 0 {
												goto l24
											}
											{
												if v4 == 0 {
													goto l25
												}
												t66 := int64(load64(m.memory[int64(uint32(v2))+80:]))
												t67 := int64(load64(m.memory[int64(uint32(v2))+88:]))
												t68 := m.fn651(t66, t67, v1, v16)
												t69 := m.fn652(v25, v9, t68, v2+i32(184))
												v14 = t69
												if v14 == 0 {
													goto l25
												}
												t70 := int32(load32(m.memory[uint32(v14+i32(-4)):]))
												v3 = t70
												t71 := int32(load32(m.memory[uint32(v14+i32(-8)):]))
												t72 := v13
												t73 := v15
												v14 = t71
												t74 := m.fn857(t72, t73, v14, i32(1080820))
												v10 = t74
												t75 := int32(load32(m.memory[uint32(v10+i32(4)):]))
												t76 := int32(load32(m.memory[uint32(v10+i32(8)):]))
												t77 := m.fn1459(t75, t76, v3, i32(1080836))
												v10 = t77
												t78 := int32(load32(m.memory[int64(uint32(v10))+20:]))
												store32(m.memory[int64(uint32(v10))+20:], uint32(t78+i32(1)))
												m.fn1118(v2+i32(96), v2+i32(232), v1, v16, v14, v3)
												goto l26
											}
										l25:
											t79 := m.fn857(v13, v15, v8, i32(1080852))
											v14 = t79
											t80 := int32(load32(m.memory[uint32(v14+i32(4)):]))
											t81 := int32(load32(m.memory[uint32(v14+i32(8)):]))
											t82 := m.fn1459(t80, t81, v5, i32(1080868))
											m.memory[int64(uint32(t82))+25] = byte(i32(0))
										}
									l24:
										t83 := m.fn857(v13, v15, v8, i32(1080884))
										v14 = t83
										t84 := int32(load32(m.memory[uint32(v14+i32(4)):]))
										t85 := int32(load32(m.memory[uint32(v14+i32(8)):]))
										t86 := m.fn1459(t84, t85, v5, i32(1080900))
										t87 := int32(m.memory[int64(uint32(t86))+24])
										if t87 == 0 {
											goto l26
										}
										m.fn1118(v2+i32(96), v2+i32(232), v1, v16, v8, v5)
										goto l26
									}
								l26:
									v5 = v5 + i32(1)
									goto l28
								}
							}
							v5 = v5 + i32(-1)
							m.fn968(v21)
							v21 = v21 + i32(12)
							goto l21
						}
					}
					t37 := int64(load64(m.memory[uint32(v1+v5):]))
					v17 = t37
					{
						t38 := int32(load32(m.memory[int64(uint32(v2))+24:]))
						v8 = t38
						if v8 == 0 {
							goto l15
						}
						t39 := int32(load32(m.memory[int64(uint32(v2))+20:]))
						v8 = t39 + v8<<3 + i32(-8)
						if v8 == 0 {
							goto l15
						}
						t40 := int64(load64(m.memory[uint32(v8):]))
						if v17-t40 <= i64(10) {
							goto l16
						}
					}
				l15:
					m.fn1413(v2+i32(16), v17)
				l16:
					v5 = v5 + i32(8)
					goto l17
				}
			}
		}
		t7 := int32(load32(m.memory[int64(uint32(v13))+8:]))
		v15 = t7 << 5
		t8 := int32(load32(m.memory[int64(uint32(v13))+4:]))
		v1 = t8
		t9 := int32(load32(m.memory[uint32(v13):]))
		v16 = t9
		store32(m.memory[int64(uint32(v2))+248:], uint32(v2+i32(96)))
		v17 = i64(-0x8000000000000000)
		v5 = i32(0)
	l5:
		{
			v8 = v1 + v5
			if v15 == v5 {
				store32(m.memory[int64(uint32(v2))+240:], uint32(i32(0)))
				store64(m.memory[int64(uint32(v2))+96:], uint64(v17))
				store32(m.memory[int64(uint32(v2))+232:], uint32(i32(8)))
				m.fn80(i32(0), i32(8))
				store32(m.memory[int64(uint32(v2))+244:], uint32(i32(8)))
				store32(m.memory[int64(uint32(v2))+236:], uint32(i32(8)))
				m.fn974(v2 + i32(232))
				store32(m.memory[int64(uint32(v14))+8:], uint32(int32(uint32(v8-v1)>>5)))
				store32(m.memory[int64(uint32(v14))+4:], uint32(v1))
				store32(m.memory[uint32(v14):], uint32(v16))
				v12 = v12 + i32(12)
				v14 = v14 + i32(12)
				v13 = v13 + i32(16)
				goto l57
			}
			v8 = v8 + i32(16)
			t10 := int64(load64(m.memory[uint32(v8):]))
			t11 := v8
			v18 = t10
			p12 := v17 + i64(1)
			if v18 > v17 {
				p12 = v18
			}
			v17 = p12
			store64(m.memory[uint32(t11):], uint64(v17))
			v5 = v5 + i32(32)
			goto l5
		}
	}
}
func (m *Module) fn1453(v0, v1 int32) {
	var v2, v3, v4, v5, v6, v7, v8, v9, v10 int32
	var v11, v12, v13 int64
	var v14 int32
	var v15 int64
	t0 := m.g0
	v2 = t0 - i32(272)
	m.g0 = v2
	{
		{
			t1 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			v3 = t1
			if v3 == 0 {
				store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
				store64(m.memory[uint32(v0):], uint64(i64(0x800000000)))
				m.fn1302(v1)
				goto l22
			}
			t2 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			v4 = t2
			t3 := int32(load32(m.memory[int64(uint32(v4))+24:]))
			v5 = t3
			{
				if v3 == i32(1) {
					goto l1
				}
				v6 = v4 + i32(80)
				t4 := int32(uint32(v3*i32(56)+i32(-56)) / uint32(i32(56)))
				v7 = t4
			l2:
				{
					t5 := int32(load32(m.memory[uint32(v6):]))
					t6 := v5
					v8 = t5
					p7 := v8
					if uint32(v5) < uint32(v8) {
						p7 = t6
					}
					v5 = p7
					v6 = v6 + i32(56)
					v7 = v7 + i32(-1)
					if v7 != 0 {
						goto l2
					}
				}
			}
		l1:
			store32(m.memory[int64(uint32(v2))+20:], uint32(i32(0)))
			store64(m.memory[int64(uint32(v2))+12:], uint64(i64(0x800000000)))
			store32(m.memory[int64(uint32(v2))+32:], uint32(i32(-1)))
			store32(m.memory[int64(uint32(v2))+128:], uint32(i32(-2)))
			store32(m.memory[int64(uint32(v2))+84:], uint32(v4+v3*i32(56)))
			store32(m.memory[int64(uint32(v2))+76:], uint32(v4))
			store32(m.memory[int64(uint32(v2))+72:], uint32(v4))
			t8 := int32(load32(m.memory[uint32(v1):]))
			store32(m.memory[int64(uint32(v2))+80:], uint32(t8))
			v8 = v2 + i32(160) + i32(44)
			v3 = v2 + i32(172)
			v9 = v2 + i32(216) + i32(28)
			v1 = v2 + i32(88)
			v4 = v2 + i32(132)
			v10 = v2 + i32(32)
		l14:
			{
				t9 := m.fn1456(v2 + i32(72))
				v6 = t9
				if v6 == 0 {
					goto l3
				}
				{
					t10 := int32(load32(m.memory[int64(uint32(v6))+24:]))
					if uint32(t10) <= uint32(v5) {
						t24 := int64(load64(m.memory[int64(uint32(v6))+16:]))
						v11 = t24
						t25 := int32(m.memory[int64(uint32(v6))+8])
						v7 = t25
						t26 := int64(load64(m.memory[uint32(v6):]))
						v12 = t26
						t27 := int32(load32(m.memory[int64(uint32(v2))+128:]))
						v6 = t27
						store32(m.memory[int64(uint32(v2))+128:], uint32(i32(-2)))
						{
							{
								{
									if v6 == i32(-2) {
										goto l17
									}
									memory_copy(m.memory, uint32(v2+i32(160)), uint32(v1), uint32(i32(40)))
									t28 := int64(load64(m.memory[int64(uint32(v2))+132:]))
									v13 = t28
									goto l18
								}
							l17:
								t29 := int32(load32(m.memory[int64(uint32(v2))+76:]))
								v6 = t29
								t30 := int32(load32(m.memory[int64(uint32(v2))+84:]))
								if v6 == t30 {
									goto l19
								}
								store32(m.memory[int64(uint32(v2))+76:], uint32(v6+i32(56)))
								memory_copy(m.memory, uint32(v2+i32(160)), uint32(v6), uint32(i32(40)))
								t31 := int64(load64(m.memory[int64(uint32(v6))+44:]))
								v13 = t31
								t32 := int32(load32(m.memory[int64(uint32(v6))+40:]))
								v6 = t32
							}
						l18:
							if v6 == i32(-1) {
								goto l19
							}
							memory_copy(m.memory, uint32(v2+i32(216)), uint32(v2+i32(160)), uint32(i32(40)))
							{
								t33 := int32(load32(m.memory[int64(uint32(v2))+32:]))
								if t33 == i32(-1) {
									goto l20
								}
								t34 := int64(load64(m.memory[int64(uint32(v2))+48:]))
								if t34 != v12 {
									goto l20
								}
								t35 := int32(m.memory[int64(uint32(v2))+56])
								t36 := t35 & i32(255)
								v14 = v7 & i32(255)
								if t36 != v14 {
									goto l20
								}
								if v14 == 0 {
									goto l21
								}
								t37 := int64(load64(m.memory[int64(uint32(v2))+64:]))
								v15 = t37
								if v15 == i64(-1) {
									goto l20
								}
								if v15+i64(1) == v11 {
									goto l21
								}
							}
						l20:
							m.fn1458(v2+i32(24), v2+i32(12))
							m.fn1457(v2 + i32(24))
							m.memory[int64(uint32(v2))+56] = byte(v7)
							store64(m.memory[int64(uint32(v2))+48:], uint64(v12))
							m.memory[int64(uint32(v2))+44] = byte(v7)
							store32(m.memory[int64(uint32(v2))+40:], uint32(i32(0)))
							store64(m.memory[int64(uint32(v2))+32:], uint64(i64(0x400000000)))
							store64(m.memory[int64(uint32(v2))+64:], uint64(v11))
							t39 := v2
							p38 := i64(1)
							if v7&i32(255) != 0 {
								p38 = v11
							}
							store64(m.memory[int64(uint32(t39))+24:], uint64(p38))
							goto l21
						}
					l19:
						m.fn153(i32(1081248))
						panic("unreachable")
					l21:
						t40 := int32(load32(m.memory[int64(uint32(v9))+8:]))
						store32(m.memory[int64(uint32(v3))+8:], uint32(t40))
						t41 := int64(load64(m.memory[uint32(v9):]))
						store64(m.memory[uint32(v3):], uint64(t41))
						m.memory[int64(uint32(v2))+184] = byte(i32(2))
						store64(m.memory[int64(uint32(v2))+164:], uint64(v13))
						store32(m.memory[int64(uint32(v2))+160:], uint32(v6))
						m.fn1412(v10, v2+i32(160))
						store64(m.memory[int64(uint32(v2))+64:], uint64(v11))
						goto l14
					}
					store32(m.memory[int64(uint32(v2))+156:], uint32(i32(0)))
					store64(m.memory[int64(uint32(v2))+148:], uint64(i64(0x800000000)))
					{
						{
							{
							l13:
								{
									{
										t11 := m.fn1456(v2 + i32(72))
										v6 = t11
										if v6 == 0 {
											goto l5
										}
										t12 := int32(load32(m.memory[int64(uint32(v6))+24:]))
										if uint32(t12) > uint32(v5) {
											t15 := int32(load32(m.memory[int64(uint32(v2))+128:]))
											v6 = t15
											store32(m.memory[int64(uint32(v2))+128:], uint32(i32(-2)))
											{
												if v6 == i32(-2) {
													goto l10
												}
												memory_copy(m.memory, uint32(v2+i32(216)), uint32(v1), uint32(i32(40)))
												v7 = v4
												goto l11
											l10:
												t16 := int32(load32(m.memory[int64(uint32(v2))+76:]))
												v6 = t16
												t17 := int32(load32(m.memory[int64(uint32(v2))+84:]))
												if v6 == t17 {
													goto l12
												}
												store32(m.memory[int64(uint32(v2))+76:], uint32(v6+i32(56)))
												memory_copy(m.memory, uint32(v2+i32(216)), uint32(v6), uint32(i32(40)))
												v7 = v6 + i32(44)
												t18 := int32(load32(m.memory[int64(uint32(v6))+40:]))
												v6 = t18
											}
										l11:
											if v6 == i32(-1) {
												goto l12
											}
											memory_copy(m.memory, uint32(v2+i32(160)), uint32(v2+i32(216)), uint32(i32(40)))
											t19 := int64(load64(m.memory[uint32(v7):]))
											store64(m.memory[uint32(v8):], uint64(t19))
											t20 := int32(load32(m.memory[int64(uint32(v7))+8:]))
											store32(m.memory[int64(uint32(v8))+8:], uint32(t20))
											store32(m.memory[int64(uint32(v2))+200:], uint32(v6))
											m.fn1369(v2+i32(148), v2+i32(160))
											goto l13
										}
									}
								l5:
									m.fn1453(v2+i32(260), v2+i32(148))
									t13 := int32(load32(m.memory[int64(uint32(v2))+268:]))
									if t13 == 0 {
										m.fn969(v2 + i32(260))
										goto l14
									}
									t14 := int32(load32(m.memory[int64(uint32(v2))+32:]))
									if t14 != i32(-1) {
										goto l8
									}
									m.fn1457(v2 + i32(24))
									store64(m.memory[int64(uint32(v2))+64:], uint64(i64(0)))
									m.memory[int64(uint32(v2))+56] = byte(i32(0))
									store64(m.memory[int64(uint32(v2))+48:], uint64(i64(-1)))
									m.memory[int64(uint32(v2))+44] = byte(i32(0))
									store32(m.memory[int64(uint32(v2))+40:], uint32(i32(0)))
									store64(m.memory[int64(uint32(v2))+32:], uint64(i64(0x400000000)))
									store64(m.memory[int64(uint32(v2))+24:], uint64(i64(1)))
									goto l9
								}
							l8:
								t21 := int32(load32(m.memory[int64(uint32(v2))+40:]))
								v6 = t21
								if v6 != 0 {
									goto l15
								}
							}
						l9:
							m.memory[int64(uint32(v2))+184] = byte(i32(2))
							store64(m.memory[int64(uint32(v2))+160:], uint64(i64(0x800000000)))
							store64(m.memory[int64(uint32(v2))+168:], uint64(i64(-0x100000000)))
							m.fn1412(v10, v2+i32(160))
							t22 := int32(load32(m.memory[int64(uint32(v2))+40:]))
							v6 = t22
						}
					l15:
						if v6 == 0 {
							goto l16
						}
						t23 := int32(load32(m.memory[int64(uint32(v2))+36:]))
						v6 = t23 + v6*i32(28) + i32(-28)
						if v6 == 0 {
							goto l16
						}
						m.fn1271(v6, v2+i32(260))
						goto l14
					}
				l12:
					m.fn153(i32(1081232))
					panic("unreachable")
				l16:
					m.fn153(i32(1081216))
					panic("unreachable")
				}
			}
		}
	l3:
		m.fn1458(v2+i32(24), v2+i32(12))
		t42 := int32(load32(m.memory[int64(uint32(v2))+20:]))
		store32(m.memory[int64(uint32(v0))+8:], uint32(t42))
		t43 := int64(load64(m.memory[int64(uint32(v2))+12:]))
		store64(m.memory[uint32(v0):], uint64(t43))
		t44 := int32(load32(m.memory[int64(uint32(v2))+84:]))
		t45 := int32(load32(m.memory[int64(uint32(v2))+76:]))
		v7 = t45
		t46 := int32(uint32(t44-v7) / uint32(i32(56)))
		v6 = t46
	l24:
		if v6 == 0 {
			goto l23
		}
		v6 = v6 + i32(-1)
		m.fn1353(v7)
		v7 = v7 + i32(56)
		goto l24
	l23:
		t47 := int32(load32(m.memory[int64(uint32(v2))+80:]))
		t48 := int32(load32(m.memory[int64(uint32(v2))+72:]))
		m.fn1354(t47, t48)
		{
			t49 := int32(load32(m.memory[int64(uint32(v2))+128:]))
			if uint32(t49) > uint32(i32(-3)) {
				goto l25
			}
			m.fn1353(v1)
		}
	l25:
		m.fn1457(v2 + i32(24))
	}
l22:
	m.g0 = v2 + i32(272)
}
func (m *Module) fn1454(v0, v1 int32) {
	var v2, v3, v4, v5, v6 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	t1 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	t2 := v2 + i32(8)
	v3 = t1
	t3 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	t4 := v3
	v4 = t3
	m.fn13(t2, t4, v4, i32(35))
	v5 = v0 + i32(4)
	{
		{
			t5 := int32(load32(m.memory[int64(uint32(v2))+8:]))
			v6 = t5
			if v6 == 0 {
				goto l0
			}
			t6 := int32(load32(m.memory[int64(uint32(v2))+12:]))
			m.fn51(v5, v6, t6)
			store32(m.memory[uint32(v0):], uint32(i32(2)))
			t7 := int32(load32(m.memory[uint32(v1):]))
			m.fn16(t7, v3)
			goto l1
		}
	l0:
		t8 := m.fn1455(v3, v4)
		v3 = t8
		t9 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		store32(m.memory[int64(uint32(v5))+8:], uint32(t9))
		t10 := int64(load64(m.memory[uint32(v1):]))
		store64(m.memory[uint32(v5):], uint64(t10))
		if v3 != 0 {
			goto l2
		}
		store32(m.memory[uint32(v0):], uint32(i32(1)))
		goto l1
	l2:
		store32(m.memory[uint32(v0):], uint32(i32(0)))
	}
l1:
	m.g0 = v2 + i32(16)
}
func (m *Module) fn1455(v0, v1 int32) int32 {
	var v2, v3, v4, v5 int32
	t0 := m.g0
	v2 = t0 - i32(32)
	m.g0 = v2
	m.fn1056(v2+i32(16), v0, v1, i32(58))
	v3 = i32(0)
	{
		t1 := int32(load32(m.memory[int64(uint32(v2))+16:]))
		if t1 != i32(1) {
			goto l0
		}
		t2 := int32(load32(m.memory[int64(uint32(v2))+20:]))
		m.fn825(v2+i32(8), v0, v1, t2, i32(1080692))
		t3 := int32(load32(m.memory[int64(uint32(v2))+8:]))
		t4 := v2
		v4 = t3
		store32(m.memory[int64(uint32(t4))+24:], uint32(v4))
		t5 := int32(load32(m.memory[int64(uint32(v2))+12:]))
		store32(m.memory[int64(uint32(v2))+28:], uint32(v4+t5))
		t6 := m.fn48(v2 + i32(24))
		if uint32(t6&i32(2097119)+i32(-65)) > uint32(i32(25)) {
			goto l0
		}
	l2:
		{
			t7 := m.fn48(v2 + i32(24))
			v4 = t7
			if v4 == i32(-1) {
				goto l1
			}
			v5 = v4&i32(2097119) + i32(-65)
			if uint32(v4+i32(-48)) < uint32(i32(10)) {
				goto l2
			}
			if uint32(v5) < uint32(i32(26)) {
				goto l2
			}
			v4 = v4 + i32(-43)
			if uint32(v4) > uint32(i32(3)) {
				goto l0
			}
			if v4 != i32(1) {
				goto l2
			}
			goto l0
		}
	l1:
		v3 = i32(1)
		if uint32(v1) < uint32(i32(3)) {
			goto l0
		}
		t8 := int32(m.memory[uint32(v0)])
		if uint32((t8&i32(223)+i32(-65))&i32(255)) > uint32(i32(25)) {
			goto l0
		}
		t9 := int32(m.memory[int64(uint32(v0))+1])
		if t9 != i32(58) {
			goto l0
		}
		t10 := int32(m.memory[int64(uint32(v0))+2])
		v4 = t10
		var p11 int32
		if v4 != i32(92) {
			p11 = 1
		}
		var p12 int32
		if v4 != i32(47) {
			p12 = 1
		}
		v3 = p11 & p12
	}
l0:
	m.g0 = v2 + i32(32)
	return v3
}
func (m *Module) fn1456(v0 int32) int32 {
	var v1, v2, v3 int32
	t0 := m.g0
	v1 = t0 - i32(64)
	v2 = v0 + i32(16)
	{
		t1 := int32(load32(m.memory[int64(uint32(v0))+56:]))
		v3 = t1
		if v3 != i32(-2) {
			goto l0
		}
		{
			{
				t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
				v3 = t2
				t3 := int32(load32(m.memory[int64(uint32(v0))+12:]))
				if v3 != t3 {
					goto l1
				}
				v3 = i32(-1)
				goto l2
			}
		l1:
			store32(m.memory[int64(uint32(v0))+4:], uint32(v3+i32(56)))
			memory_copy(m.memory, uint32(v1+i32(24)), uint32(v3), uint32(i32(40)))
			t4 := int64(load64(m.memory[int64(uint32(v3))+44:]))
			store64(m.memory[int64(uint32(v1))+8:], uint64(t4))
			t5 := int32(load32(m.memory[int64(uint32(v3))+52:]))
			store32(m.memory[int64(uint32(v1))+16:], uint32(t5))
			t6 := int32(load32(m.memory[int64(uint32(v3))+40:]))
			v3 = t6
		}
	l2:
		memory_copy(m.memory, uint32(v2), uint32(v1+i32(24)), uint32(i32(40)))
		store32(m.memory[int64(uint32(v0))+56:], uint32(v3))
		t7 := int64(load64(m.memory[int64(uint32(v1))+8:]))
		store64(m.memory[int64(uint32(v0))+60:], uint64(t7))
		t8 := int32(load32(m.memory[int64(uint32(v1))+16:]))
		store32(m.memory[int64(uint32(v0))+68:], uint32(t8))
	}
l0:
	p9 := v2
	if v3 == i32(-1) {
		p9 = i32(0)
	}
	return p9
}
func (m *Module) fn1457(v0 int32) {
	t0 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	if t0 == i32(-1) {
		return
	}
	m.fn971(v0 + i32(8))
}
func (m *Module) fn1458(v0, v1 int32) {
	var v2, v3 int32
	var v4 int64
	t0 := m.g0
	v2 = t0 - i32(64)
	m.g0 = v2
	t1 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	v3 = t1
	store32(m.memory[int64(uint32(v0))+8:], uint32(i32(-1)))
	{
		if v3 == i32(-1) {
			goto l0
		}
		t2 := int64(load64(m.memory[uint32(v0):]))
		v4 = t2
		t3 := int64(load64(m.memory[int64(uint32(v0))+12:]))
		store64(m.memory[int64(uint32(v2))+20:], uint64(t3))
		store32(m.memory[int64(uint32(v2))+16:], uint32(v3))
		store64(m.memory[int64(uint32(v2))+8:], uint64(v4))
		t4 := int32(load32(m.memory[int64(uint32(v0))+20:]))
		store32(m.memory[int64(uint32(v2))+28:], uint32(t4))
		{
			t5 := int32(load32(m.memory[int64(uint32(v2))+24:]))
			if t5 == 0 {
				goto l1
			}
			t6 := int64(load64(m.memory[int64(uint32(v2))+24:]))
			store64(m.memory[int64(uint32(v2))+56:], uint64(t6))
			t7 := int64(load64(m.memory[int64(uint32(v2))+16:]))
			store64(m.memory[int64(uint32(v2))+48:], uint64(t7))
			t8 := int64(load64(m.memory[int64(uint32(v2))+8:]))
			store64(m.memory[int64(uint32(v2))+40:], uint64(t8))
			store32(m.memory[int64(uint32(v2))+32:], uint32(i32(-0x7fffffff)))
			m.fn338(v1, v2+i32(32))
			goto l0
		}
	l1:
		m.fn971(v2 + i32(16))
	}
l0:
	m.g0 = v2 + i32(64)
}
func (m *Module) fn1459(v0, v1, v2, v3 int32) int32 {
	if uint32(v2) < uint32(v1) {
		return v0 + v2*i32(28)
	}
	m.fn158(v2, v1, v3)
	panic("unreachable")
}
func (m *Module) fn1460(v0 int32) {
	var v1, v2, v3, v4, v5, v6, v7 int32
	t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
	t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	v1 = t1
	t2 := int32(uint32(t0-v1) / uint32(i32(12)))
	v2 = t2
	v3 = i32(0)
l3:
	{
		if v3 == v2 {
			goto l0
		}
		v4 = v1 + v3*i32(12)
		t3 := int32(load32(m.memory[uint32(v4+i32(8)):]))
		v5 = t3
		v6 = v4 + i32(4)
		t4 := int32(load32(m.memory[uint32(v6):]))
		v7 = t4
	l2:
		if v5 == 0 {
			t5 := int32(load32(m.memory[uint32(v4):]))
			t6 := int32(load32(m.memory[uint32(v6):]))
			m.fn82(t5, t6)
			v3 = v3 + i32(1)
			goto l3
		}
		v5 = v5 + i32(-1)
		m.fn969(v7)
		v7 = v7 + i32(28)
		goto l2
	}
l0:
	t7 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	t8 := int32(load32(m.memory[uint32(v0):]))
	m.fn136(t7, t8, i32(4), i32(12))
}
func (m *Module) fn1461(v0 int32) {
	var v1, v2 int32
	t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
	t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	v1 = t1
	t2 := int32(uint32(t0-v1) / uint32(i32(28)))
	v2 = t2
l1:
	if v2 == 0 {
		goto l0
	}
	v2 = v2 + i32(-1)
	m.fn969(v1)
	v1 = v1 + i32(28)
	goto l1
l0:
	t3 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	t4 := int32(load32(m.memory[uint32(v0):]))
	m.fn82(t3, t4)
}
func (m *Module) fn1462(v0, v1, v2, v3 int32) int32 {
	var v4, v5, v6, v7, v8 int32
	t0 := m.g0
	v4 = t0 - i32(128)
	m.g0 = v4
	v3 = v3 << 2
l2:
	{
		if v3 != 0 {
			goto l0
		}
		v0 = i32(0)
		goto l1
	l0:
		v3 = v3 + i32(-4)
		t1 := int32(load32(m.memory[uint32(v2):]))
		v5 = t1
		v2 = v2 + i32(4)
		t2 := m.fn779(v0, v1, v5)
		if t2 == 0 {
			goto l2
		}
	}
	m.fn513(v4+i32(36), v0, v1, v5)
	m.fn515(v4+i32(24), v4+i32(36))
	{
		{
			t3 := int32(load32(m.memory[int64(uint32(v4))+24:]))
			v3 = t3
			if v3 != 0 {
				goto l3
			}
			v6 = i32(4)
			v2 = i32(0)
			v7 = i32(0)
			goto l4
		}
	l3:
		t4 := int32(load32(m.memory[int64(uint32(v4))+28:]))
		v2 = t4
		m.fn59(v4+i32(16), i32(4), i32(4), i32(8))
		t5 := int32(load32(m.memory[int64(uint32(v4))+16:]))
		v5 = t5
		t6 := int32(load32(m.memory[int64(uint32(v4))+20:]))
		v8 = t6
		store32(m.memory[int64(uint32(v8))+4:], uint32(v2))
		store32(m.memory[uint32(v8):], uint32(v3))
		store32(m.memory[int64(uint32(v4))+84:], uint32(i32(1)))
		store32(m.memory[int64(uint32(v4))+80:], uint32(v8))
		store32(m.memory[int64(uint32(v4))+76:], uint32(v5))
		memory_copy(m.memory, uint32(v4+i32(88)), uint32(v4+i32(36)), uint32(i32(40)))
		v3 = i32(12)
		v2 = i32(1)
	l7:
		{
			m.fn515(v4+i32(8), v4+i32(88))
			t7 := int32(load32(m.memory[int64(uint32(v4))+8:]))
			v5 = t7
			if v5 == 0 {
				goto l5
			}
			t8 := int32(load32(m.memory[int64(uint32(v4))+12:]))
			v1 = t8
			{
				t9 := int32(load32(m.memory[int64(uint32(v4))+76:]))
				if v2 != t9 {
					goto l6
				}
				m.fn797(v4 + i32(76))
				t10 := int32(load32(m.memory[int64(uint32(v4))+80:]))
				v8 = t10
			}
		l6:
			v0 = v8 + v3
			store32(m.memory[uint32(v0):], uint32(v1))
			store32(m.memory[uint32(v0+i32(-4)):], uint32(v5))
			t11 := v4
			v2 = v2 + i32(1)
			store32(m.memory[int64(uint32(t11))+84:], uint32(v2))
			v3 = v3 + i32(8)
			goto l7
		}
	l5:
		t12 := int32(load32(m.memory[int64(uint32(v4))+80:]))
		v6 = t12
		t13 := int32(load32(m.memory[int64(uint32(v4))+76:]))
		v7 = t13
	}
l4:
	v8 = v6 + v2<<3
	v5 = v6
l11:
	{
		if v5 != v8 {
			goto l8
		}
		v0 = v2
		goto l9
	l8:
		v0 = i32(0)
		t14 := int32(load32(m.memory[int64(uint32(v5))+4:]))
		v3 = t14
		t15 := m.fn1468(i32(1073920), v3)
		if t15 == 0 {
			goto l9
		}
		v1 = v5 + i32(8)
		t16 := int32(load32(m.memory[uint32(v5):]))
		t17 := v4
		v5 = t16
		store32(m.memory[int64(uint32(t17))+88:], uint32(v5))
		store32(m.memory[int64(uint32(v4))+92:], uint32(v5+v3))
	l10:
		{
			t18 := m.fn48(v4 + i32(88))
			v3 = t18
			if uint32(v3+i32(-48)) < uint32(i32(10)) {
				goto l10
			}
		}
		v5 = v1
		if v3 == i32(-1) {
			goto l11
		}
	}
l9:
	m.fn639(v7, v6)
l1:
	m.g0 = v4 + i32(128)
	return v0
}
func (m *Module) fn1463(v0, v1 int32) int32 {
	t0 := m.fn1462(v0, v1, i32(1073916), i32(1))
	var p1 int32
	if t0&i32(0xffffffe) == i32(2) {
		p1 = 1
	}
	return p1
}
func (m *Module) fn1464(v0 int32) int32 {
	var v1, v2, v3, v4, v5 int32
	v1 = i32(0)
	p0 := i32(18)
	if uint32(v0) < uint32(i32(73459)) {
		p0 = i32(0)
	}
	v2 = p0
	t1 := v2
	v2 = v2 | i32(9)
	t2 := int32(load32(m.memory[int64(uint32(v2<<2))+1114844:]))
	t3 := v2
	t4 := t2 << 11
	v2 = v0 << 11
	p5 := t3
	if uint32(t4) > uint32(v2) {
		p5 = t1
	}
	v3 = p5
	t6 := v3
	v3 = v3 | i32(4)
	t7 := int32(load32(m.memory[int64(uint32(v3<<2))+1114844:]))
	p8 := v3
	if uint32(t7<<11) > uint32(v2) {
		p8 = t6
	}
	v3 = p8
	t9 := v3
	v3 = v3 + i32(2)
	t10 := int32(load32(m.memory[int64(uint32(v3<<2))+1114844:]))
	p11 := v3
	if uint32(t10<<11) > uint32(v2) {
		p11 = t9
	}
	v3 = p11
	t12 := v3
	v3 = v3 + i32(1)
	t13 := int32(load32(m.memory[int64(uint32(v3<<2))+1114844:]))
	p14 := v3
	if uint32(t13<<11) > uint32(v2) {
		p14 = t12
	}
	v3 = p14
	t15 := v3
	v3 = v3 + i32(1)
	t16 := int32(load32(m.memory[int64(uint32(v3<<2))+1114844:]))
	p17 := v3
	if uint32(t16<<11) > uint32(v2) {
		p17 = t15
	}
	v3 = p17
	t18 := int32(load32(m.memory[int64(uint32(v3<<2))+1114844:]))
	v4 = t18 << 11
	var p19 int32
	if v4 == v2 {
		p19 = 1
	}
	var p20 int32
	if uint32(v4) < uint32(v2) {
		p20 = 1
	}
	v3 = p19 + p20 + v3
	v2 = v3 << 2
	v5 = v2 + i32(1114844)
	t21 := int32(load32(m.memory[int64(uint32(v2))+1114844:]))
	v2 = int32(uint32(t21) >> 21)
	v4 = i32(919)
	{
		{
			if uint32(v3) > uint32(i32(34)) {
				goto l0
			}
			t22 := int32(load32(m.memory[int64(uint32(v5))+4:]))
			v4 = int32(uint32(t22) >> 21)
			if v3 == 0 {
				goto l1
			}
		}
	l0:
		t23 := int32(load32(m.memory[uint32(v5+i32(-4)):]))
		v1 = t23 & i32(0x1fffff)
	}
l1:
	if v4+(v2^i32(-1)) == 0 {
		goto l2
	}
	v3 = v0 - v1
	v4 = v4 + i32(-1)
	v0 = i32(0)
l3:
	{
		t24 := int32(m.memory[uint32(v2+i32(1105559))])
		v0 = v0 + t24
		if uint32(v0) > uint32(v3) {
			goto l2
		}
		t25 := v4
		v2 = v2 + i32(1)
		if t25 != v2 {
			goto l3
		}
	}
l2:
	return v2 & i32(1)
}
func (m *Module) fn1465(v0 int32) int32 {
	var v1, v2 int32
	var v3, v4 int64
	v1 = i32(0)
	if uint32(v0) > uint32(i32(125951)) {
		goto l0
	}
	{
		t0 := int32(m.memory[int64(uint32(int32(uint32(v0)>>10)))+1107557])
		t1 := int32(m.memory[int64(uint32(t0<<4|int32(uint32(v0)>>6)&i32(15)))+1115808])
		v1 = t1
		if uint32(v1) < uint32(i32(57)) {
			t6 := int64(load64(m.memory[int64(uint32(v1<<3))+1115352:]))
			v3 = t6
			goto l4
		}
		v2 = v1 + i32(-57)
		if uint32(v1) >= uint32(i32(79)) {
			m.fn158(v2, i32(22), i32(1107984))
			panic("unreachable")
		}
		v1 = v2 << 1
		t2 := int32(m.memory[int64(uint32(v1))+1115304])
		t3 := int64(load64(m.memory[int64(uint32(t2<<3))+1115352:]))
		v2 = i32_shl(i32(1), v2)
		p4 := i64(-1)
		if v2&i32(2047998) != 0 {
			p4 = i64(0)
		}
		v3 = t3 ^ p4
		t5 := int32(m.memory[int64(uint32(v1))+1115305])
		v4 = int64(uint32(t5))
		if v2&i32(0x2cc001) == 0 {
			v3 = i64_shr_u(v3, v4)
			goto l4
		}
		v3 = i64_rotl(v3, v4&i64(255))
		goto l4
	}
l4:
	v1 = int32(i64_shr_u(v3, int64(uint32(v0))))
l0:
	return v1 & i32(1)
}
func (m *Module) fn1466(v0 int32) int32 {
	var v1, v2 int32
	var v3, v4 int64
	v1 = i32(0)
	if uint32(v0+i32(-192)) > uint32(i32(127807)) {
		goto l0
	}
	{
		t0 := int32(m.memory[int64(uint32(int32(uint32(v0)>>10)))+1107680])
		t1 := int32(m.memory[int64(uint32(t0<<4|int32(uint32(v0)>>6)&i32(15)))+1116536])
		v1 = t1
		if uint32(v1) < uint32(i32(44)) {
			t6 := int64(load64(m.memory[int64(uint32(v1<<3))+1116184:]))
			v3 = t6
			goto l4
		}
		v2 = v1 + i32(-44)
		if uint32(v1) >= uint32(i32(69)) {
			m.fn158(v2, i32(25), i32(1107984))
			panic("unreachable")
		}
		v1 = v2 << 1
		t2 := int32(m.memory[int64(uint32(v1))+1116128])
		t3 := int64(load64(m.memory[int64(uint32(t2<<3))+1116184:]))
		v2 = i32_shl(i32(1), v2)
		p4 := i64(-1)
		if v2&i32(33539069) != 0 {
			p4 = i64(0)
		}
		v3 = t3 ^ p4
		t5 := int32(m.memory[int64(uint32(v1))+1116129])
		v4 = int64(uint32(t5))
		if v2&i32(4258818) == 0 {
			v3 = i64_shr_u(v3, v4)
			goto l4
		}
		v3 = i64_rotl(v3, v4&i64(255))
		goto l4
	}
l4:
	v1 = int32(i64_shr_u(v3, int64(uint32(v0))))
l0:
	return v1 & i32(1)
}
func (m *Module) fn1467(v0 int32) int32 {
	var v1, v2, v3, v4 int32
	var p0 int32
	if uint32(v0) > uint32(i32(8071)) {
		p0 = 1
	}
	v1 = p0
	t2 := v1
	p1 := i32(1)
	if v1 != 0 {
		p1 = i32(2)
	}
	v1 = p1
	t3 := int32(load32(m.memory[int64(uint32(v1<<2))+1115292:]))
	t4 := v1
	t5 := t3 << 11
	v1 = v0 << 11
	p6 := t4
	if uint32(t5) > uint32(v1) {
		p6 = t2
	}
	v2 = p6
	t7 := int32(load32(m.memory[int64(uint32(v2<<2))+1115292:]))
	t8 := v2
	v2 = t7 << 11
	var p9 int32
	if uint32(v2) < uint32(v1) {
		p9 = 1
	}
	t10 := t8 + p9
	var p11 int32
	if v2 == v1 {
		p11 = 1
	}
	v2 = t10 + p11
	v1 = v2 << 2
	v3 = v1 + i32(1115292)
	v4 = i32(21)
	t12 := int32(load32(m.memory[int64(uint32(v1))+1115292:]))
	v1 = int32(uint32(t12) >> 21)
	{
		{
			if uint32(v2) > uint32(i32(1)) {
				goto l0
			}
			t13 := int32(load32(m.memory[int64(uint32(v3))+4:]))
			v4 = int32(uint32(t13) >> 21)
			if v2 != 0 {
				goto l0
			}
			v2 = i32(0)
			goto l1
		}
	l0:
		t14 := int32(load32(m.memory[uint32(v3+i32(-4)):]))
		v2 = t14 & i32(0x1fffff)
	}
l1:
	if v4+(v1^i32(-1)) == 0 {
		goto l2
	}
	v2 = v0 - v2
	v4 = v4 + i32(-1)
	v0 = i32(0)
l3:
	{
		t15 := int32(m.memory[uint32(v1+i32(1107536))])
		v0 = v0 + t15
		if uint32(v0) > uint32(v2) {
			goto l2
		}
		t16 := v4
		v1 = v1 + i32(1)
		if t16 != v1 {
			goto l3
		}
	}
l2:
	return v1 & i32(1)
}
func (m *Module) fn1468(v0, v1 int32) int32 {
	var v2 int32
	v2 = i32(0)
	{
		t0 := int32(load32(m.memory[uint32(v0):]))
		if uint32(t0) > uint32(v1) {
			goto l0
		}
		t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		v2 = t1
		{
			t2 := int32(m.memory[int64(uint32(v0))+8])
			if t2 != 0 {
				goto l1
			}
			var p3 int32
			if uint32(v1) <= uint32(v2) {
				p3 = 1
			}
			return p3
		}
	l1:
		;
		var p4 int32
		if uint32(v1) < uint32(v2) {
			p4 = 1
		}
		v2 = p4
	}
l0:
	return v2
}
func (m *Module) fn1469(v0, v1, v2, v3, v4 int32) {
	var v5 int32
	t0 := m.g0
	v5 = t0 - i32(80)
	m.g0 = v5
	m.memory[int64(uint32(v5))+36] = byte(v4)
	store64(m.memory[int64(uint32(v5))+16:], uint64(i64(4)))
	store64(m.memory[int64(uint32(v5))+8:], uint64(i64(0)))
	store64(m.memory[uint32(v5):], uint64(i64(0x800000000)))
	t1 := int32(load32(m.memory[int64(uint32(v1))+32:]))
	store32(m.memory[int64(uint32(v5))+32:], uint32(t1))
	t2 := int64(load64(m.memory[int64(uint32(v1))+24:]))
	store64(m.memory[int64(uint32(v5))+24:], uint64(t2))
	m.fn1267(v5+i32(40), v5, v2, v3)
	{
		t3 := int32(load32(m.memory[int64(uint32(v5))+40:]))
		v1 = t3
		if v1 == i32(-1) {
			goto l0
		}
		t4 := int32(load32(m.memory[int64(uint32(v5))+60:]))
		store32(m.memory[int64(uint32(v0))+20:], uint32(t4))
		t5 := int64(load64(m.memory[int64(uint32(v5))+52:]))
		store64(m.memory[int64(uint32(v0))+12:], uint64(t5))
		t6 := int64(load64(m.memory[int64(uint32(v5))+44:]))
		store64(m.memory[int64(uint32(v0))+4:], uint64(t6))
		m.fn1268(v5)
		goto l1
	}
l0:
	memory_copy(m.memory, uint32(v5+i32(40)), uint32(v5), uint32(i32(40)))
	m.fn1270(v0+i32(4), v5+i32(40))
l1:
	store32(m.memory[uint32(v0):], uint32(v1))
	m.g0 = v5 + i32(80)
}
func (m *Module) fn1470(v0, v1, v2 int32) {
	var v3, v4 int32
	m.fn892(v0, v2)
	t0 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	v3 = t0
	{
		if v2 == 0 {
			goto l0
		}
		v4 = v2 * i32(28)
		if v4 == 0 {
			goto l0
		}
		t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		memory_copy(m.memory, uint32(t1+v3*i32(28)), uint32(v1), uint32(v4))
	}
l0:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v3+v2))
}
func (m *Module) fn1471(v0, v1 int32) {
	if v0 == i32(-1) {
		return
	}
	m.fn895(v0, v1)
}
func (m *Module) fn1472(v0 int32) {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	m.fn16(t0, t1)
}
func (m *Module) fn1473(v0, v1, v2, v3 int32) {
	var v4, v5 int32
	t0 := m.g0
	v4 = t0 - i32(48)
	m.g0 = v4
	{
		if v3 == 0 {
			store32(m.memory[uint32(v0):], uint32(i32(-1)))
			goto l2
		}
		m.fn13(v4+i32(8), v2, v3, i32(35))
		t1 := int32(load32(m.memory[int64(uint32(v4))+8:]))
		v5 = t1
		if v5 == 0 {
			goto l1
		}
		t2 := int32(load32(m.memory[int64(uint32(v4))+12:]))
		m.fn776(v4+i32(20), v5, t2)
		t3 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t4 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		t5 := int32(load32(m.memory[int64(uint32(v4))+24:]))
		t6 := v0 + i32(4)
		v3 = t5
		t7 := int32(load32(m.memory[int64(uint32(v4))+28:]))
		m.fn1474(t6, t3, t4, v3, t7)
		store32(m.memory[uint32(v0):], uint32(i32(2)))
		t8 := int32(load32(m.memory[int64(uint32(v4))+20:]))
		m.fn16(t8, v3)
		goto l2
	}
l1:
	{
		t9 := m.fn1455(v2, v3)
		if t9 != 0 {
			goto l3
		}
		t10 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t11 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		m.fn774(v4+i32(20), t10, t11, v2, v3)
		{
			t12 := int32(load32(m.memory[int64(uint32(v4))+20:]))
			if t12 != 0 {
				m.fn51(v0+i32(4), v2, v3)
				store32(m.memory[uint32(v0):], uint32(i32(1)))
				m.fn781(v4 + i32(20))
				goto l2
			}
			t13 := int32(load32(m.memory[int64(uint32(v1))+20:]))
			v1 = v4 + i32(20) + i32(4)
			t14 := m.fn649(t13, v1)
			if t14 == 0 {
				m.fn51(v0+i32(4), v2, v3)
				store32(m.memory[uint32(v0):], uint32(i32(1)))
				m.fn784(v1)
				goto l2
			}
			t15 := int32(load32(m.memory[int64(uint32(v4))+28:]))
			t16 := int32(load32(m.memory[int64(uint32(v4))+32:]))
			t17 := int32(load32(m.memory[int64(uint32(v4))+40:]))
			t18 := int32(load32(m.memory[int64(uint32(v4))+36:]))
			t20 := v0 + i32(4)
			p19 := t17
			if t18 == i32(-1) {
				p19 = i32(0)
			}
			t21 := int32(load32(m.memory[int64(uint32(v4))+44:]))
			m.fn1474(t20, t15, t16, p19, t21)
			store32(m.memory[uint32(v0):], uint32(i32(2)))
			m.fn784(v1)
			goto l2
		}
	}
l3:
	m.fn51(v0+i32(4), v2, v3)
	store32(m.memory[uint32(v0):], uint32(i32(0)))
l2:
	m.g0 = v4 + i32(48)
}
func (m *Module) fn1474(v0, v1, v2, v3, v4 int32) {
	var v5 int32
	t0 := m.g0
	v5 = t0 - i32(32)
	m.g0 = v5
	store32(m.memory[int64(uint32(v5))+4:], uint32(v2))
	store32(m.memory[uint32(v5):], uint32(v1))
	if v3 == 0 {
		goto l0
	}
	if v4 != 0 {
		goto l1
	}
l0:
	m.fn51(v0, v1, v2)
	goto l2
l1:
	store32(m.memory[int64(uint32(v5))+12:], uint32(v4))
	store32(m.memory[int64(uint32(v5))+8:], uint32(v3))
	store32(m.memory[int64(uint32(v5))+28:], uint32(i32(1)))
	store32(m.memory[int64(uint32(v5))+20:], uint32(i32(1)))
	store32(m.memory[int64(uint32(v5))+24:], uint32(v5+i32(8)))
	store32(m.memory[int64(uint32(v5))+16:], uint32(v5))
	m.fn73(v0, i32(0x1000d8), v5+i32(16))
l2:
	m.g0 = v5 + i32(32)
}
func (m *Module) fn1475(v0, v1, v2, v3 int32) {
	var v4, v5, v6, v7, v8, v9, v10, v11 int32
	t0 := m.g0
	v4 = t0 - i32(80)
	m.g0 = v4
	{
		if v3 == 0 {
			store64(m.memory[uint32(v0):], uint64(i64(-1)))
			goto l8
		}
		t1 := m.fn1455(v2, v3)
		if t1 != 0 {
			m.fn51(v0+i32(4), v2, v3)
			store32(m.memory[uint32(v0):], uint32(i32(-1)))
			goto l8
		}
		t2 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t3 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		m.fn774(v4+i32(20), t2, t3, v2, v3)
		t4 := int32(load32(m.memory[int64(uint32(v4))+20:]))
		if t4 != 0 {
			m.fn781(v4 + i32(20))
			store64(m.memory[uint32(v0):], uint64(i64(-1)))
			goto l8
		}
		t5 := int32(load32(m.memory[int64(uint32(v4))+40:]))
		v5 = t5
		t6 := int32(load32(m.memory[int64(uint32(v4))+36:]))
		v6 = t6
		t7 := int32(load32(m.memory[int64(uint32(v4))+24:]))
		v7 = t7
		t8 := int32(load32(m.memory[int64(uint32(v4))+32:]))
		v8 = t8
		t9 := int32(load32(m.memory[int64(uint32(v4))+28:]))
		v3 = t9
		t10 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		m.fn1182(v4+i32(8), t10, i32(1086440))
		t11 := int32(load32(m.memory[int64(uint32(v4))+12:]))
		v2 = t11
		t12 := int32(load32(m.memory[int64(uint32(v4))+8:]))
		m.fn1035(v4+i32(20), t12, v3, v8)
		t13 := int32(load32(m.memory[int64(uint32(v4))+28:]))
		v9 = t13
		t14 := int32(load32(m.memory[int64(uint32(v4))+24:]))
		v10 = t14
		{
			{
				t15 := int32(load32(m.memory[int64(uint32(v4))+20:]))
				v11 = t15
				if v11 == i32(-1) {
					goto l3
				}
				t16 := int32(load32(m.memory[int64(uint32(v4))+40:]))
				store32(m.memory[int64(uint32(v0))+20:], uint32(t16))
				t17 := int64(load64(m.memory[int64(uint32(v4))+32:]))
				store64(m.memory[int64(uint32(v0))+12:], uint64(t17))
				store32(m.memory[int64(uint32(v0))+8:], uint32(v9))
				store32(m.memory[int64(uint32(v0))+4:], uint32(v10))
				store32(m.memory[uint32(v0):], uint32(v11))
				t18 := int32(load32(m.memory[uint32(v2):]))
				store32(m.memory[uint32(v2):], uint32(t18+i32(1)))
				goto l4
			}
		l3:
			{
				if v10 == 0 {
					goto l5
				}
				store32(m.memory[int64(uint32(v4))+52:], uint32(v9))
				store32(m.memory[int64(uint32(v4))+48:], uint32(v10))
				m.fn1476(v4+i32(56), v3, v8)
				t19 := int32(load32(m.memory[int64(uint32(v1))+16:]))
				m.fn1182(v4, t19, i32(1086456))
				t20 := int32(load32(m.memory[int64(uint32(v4))+4:]))
				v1 = t20
				t21 := int32(load32(m.memory[uint32(v4):]))
				v11 = t21
				store32(m.memory[int64(uint32(v4))+76:], uint32(v8))
				store32(m.memory[int64(uint32(v4))+72:], uint32(v3))
				store32(m.memory[int64(uint32(v4))+68:], uint32(v7))
				m.fn1296(v4+i32(20), v11, v4+i32(56), v4+i32(68), v10+i32(8), v9)
				t22 := int32(load32(m.memory[int64(uint32(v4))+24:]))
				v3 = t22
				t23 := int32(load32(m.memory[int64(uint32(v4))+20:]))
				v8 = t23
				if v8 == i32(-1) {
					store32(m.memory[int64(uint32(v0))+8:], uint32(v3))
					store64(m.memory[uint32(v0):], uint64(i64(-0x7fffffff00000001)))
					t29 := int32(load32(m.memory[uint32(v1):]))
					store32(m.memory[uint32(v1):], uint32(t29+i32(1)))
					m.fn754(v4 + i32(48))
					t30 := int32(load32(m.memory[uint32(v2):]))
					store32(m.memory[uint32(v2):], uint32(t30+i32(1)))
					goto l7
				}
				t24 := int64(load64(m.memory[int64(uint32(v4))+36:]))
				store64(m.memory[int64(uint32(v0))+16:], uint64(t24))
				t25 := int64(load64(m.memory[int64(uint32(v4))+28:]))
				store64(m.memory[int64(uint32(v0))+8:], uint64(t25))
				store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
				store32(m.memory[uint32(v0):], uint32(v8))
				t26 := int32(load32(m.memory[uint32(v1):]))
				store32(m.memory[uint32(v1):], uint32(t26+i32(1)))
				m.fn754(v4 + i32(48))
				t27 := int32(load32(m.memory[uint32(v2):]))
				store32(m.memory[uint32(v2):], uint32(t27+i32(1)))
				goto l7
			}
		l5:
			store64(m.memory[uint32(v0):], uint64(i64(-1)))
			t28 := int32(load32(m.memory[uint32(v2):]))
			store32(m.memory[uint32(v2):], uint32(t28+i32(1)))
		}
	l4:
		m.fn16(v7, v3)
		goto l7
	}
l7:
	m.fn134(v6, v5)
l8:
	m.g0 = v4 + i32(80)
}
func (m *Module) fn1476(v0, v1, v2 int32) {
	var v3, v4, v5, v6 int32
	t0 := m.g0
	v3 = t0 - i32(32)
	m.g0 = v3
	m.fn778(v3+i32(16), v1, v2, i32(46))
	{
		t1 := int32(load32(m.memory[int64(uint32(v3))+16:]))
		if t1 == 0 {
			goto l0
		}
		t2 := int32(load32(m.memory[int64(uint32(v3))+24:]))
		t3 := int32(load32(m.memory[int64(uint32(v3))+28:]))
		m.fn14(v3+i32(4), t2, t3)
		t4 := int32(load32(m.memory[int64(uint32(v3))+4:]))
		v4 = t4
		if v4 == i32(-1) {
			goto l0
		}
		t5 := int32(load32(m.memory[int64(uint32(v3))+12:]))
		v1 = t5
		t6 := int32(load32(m.memory[int64(uint32(v3))+8:]))
		v2 = t6
		goto l1
	}
l0:
	v4 = i32(0)
	v2 = i32(1)
	v1 = i32(0)
l1:
	{
		{
			t7 := m.fn15(v2, v1, i32(1079563), i32(3))
			if t7 == 0 {
				goto l2
			}
			v5 = i32(1079554)
			v6 = i32(9)
			goto l3
		}
	l2:
		v5 = i32(1079541)
		v6 = i32(10)
		t8 := m.fn15(v2, v1, i32(1079551), i32(3))
		if t8 != 0 {
			goto l3
		}
		t9 := m.fn15(v2, v1, i32(1081264), i32(4))
		if t9 != 0 {
			goto l3
		}
		{
			t10 := m.fn15(v2, v1, i32(1081268), i32(3))
			if t10 == 0 {
				goto l4
			}
			v5 = i32(1081354)
			v6 = i32(9)
			goto l3
		}
	l4:
		{
			t11 := m.fn15(v2, v1, i32(1081271), i32(3))
			if t11 == 0 {
				goto l5
			}
			v5 = i32(1081345)
			v6 = i32(9)
			goto l3
		}
	l5:
		v5 = i32(1081335)
		t12 := m.fn15(v2, v1, i32(1081274), i32(3))
		if t12 != 0 {
			goto l3
		}
		t13 := m.fn15(v2, v1, i32(1081277), i32(4))
		if t13 != 0 {
			goto l3
		}
		{
			t14 := m.fn15(v2, v1, i32(1081281), i32(3))
			if t14 == 0 {
				goto l6
			}
			v5 = i32(1081322)
			v6 = i32(13)
			goto l3
		}
	l6:
		{
			t15 := m.fn15(v2, v1, i32(1079538), i32(3))
			if t15 == 0 {
				goto l7
			}
			v5 = i32(1079529)
			v6 = i32(9)
			goto l3
		}
	l7:
		{
			t16 := m.fn15(v2, v1, i32(1079526), i32(3))
			if t16 == 0 {
				goto l8
			}
			v5 = i32(1079517)
			v6 = i32(9)
			goto l3
		}
	l8:
		t17 := m.fn15(v2, v1, i32(1081284), i32(4))
		v1 = t17
		p18 := i32(1081288)
		if v1 != 0 {
			p18 = i32(1081312)
		}
		v5 = p18
		p19 := i32(24)
		if v1 != 0 {
			p19 = i32(10)
		}
		v6 = p19
	}
l3:
	m.fn51(v0, v5, v6)
	m.fn16(v4, v2)
	m.g0 = v3 + i32(32)
}
func (m *Module) fn1477(v0, v1, v2, v3 int32) {
	t0 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	t1 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	m.fn1474(v0, t0, t1, v2, v3)
}
func (m *Module) fn1478(v0, v1, v2 int32) {
	var v3, v4 int32
	t0 := m.g0
	v3 = t0 - i32(32)
	m.g0 = v3
	{
		t1 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		v4 = t1
		if v4 == 0 {
			goto l0
		}
		t2 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		m.fn510(v3+i32(12), v0, t2, v4)
		t3 := int32(load32(m.memory[int64(uint32(v3))+20:]))
		store32(m.memory[int64(uint32(v3))+8:], uint32(t3))
		t4 := int64(load64(m.memory[int64(uint32(v3))+12:]))
		store64(m.memory[uint32(v3):], uint64(t4))
		m.fn490(v3+i32(12), v3)
		m.fn1321(v2, v3+i32(12))
		store32(m.memory[int64(uint32(v1))+8:], uint32(i32(0)))
	}
l0:
	m.g0 = v3 + i32(32)
}
func (m *Module) fn1479(v0, v1 int32) int32 {
	var v2 int64
	var v3, v4 int32
	var v5 int64
	var v6 int32
	var v7 int64
	var v8, v9 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
		if t0 != 0 {
			t1 := int64(load64(m.memory[int64(uint32(v0))+16:]))
			t2 := int64(load64(m.memory[int64(uint32(v0))+24:]))
			t3 := m.fn66(t1, t2, v1)
			v2 = t3
			t4 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			v3 = t4
			v4 = v3 & int32(v2)
			v5 = int64(uint64(v2)>>25) & i64(127) * i64(72340172838076673)
			t5 := int32(load32(m.memory[uint32(v0):]))
			v0 = t5
			v6 = i32(0)
			var _ int32
		l5:
			{
				t7 := int64(load64(m.memory[uint32(v0+v4):]))
				v7 = t7
				v2 = v7 ^ v5
				v2 = (v2 ^ i64(-1)) & (v2 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
				{
				l3:
					{
						var p8 int32
						if v2 == 0 {
							p8 = 1
						}
						v8 = p8
						if v8 != 0 {
							goto l1
						}
						t9 := v1
						t10 := v0
						v9 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v2))))>>3) + v4) & v3
						t11 := int32(load32(m.memory[uint32(t10+(i32(0)-v9)*i32(296)+i32(-296)):]))
						if t9 == t11 {
							goto l2
						}
						v2 = (v2 + i64(-1)) & v2
						goto l3
					}
				l1:
					if v7&(v7<<1)&i64(-0x7f7f7f7f7f7f7f80) == 0 {
						t14 := v4
						v6 = v6 + i32(8)
						v4 = (t14 + v6) & v3
						goto l5
					}
				l2:
					p12 := v0 + (i32(0)-v9)*i32(296)
					if v8 != 0 {
						p12 = i32(0)
					}
					p13 := p12 + i32(-288)
					if v8 != 0 {
						p13 = i32(0)
					}
					return p13
				}
			}
		}
		return i32(0)
	}
}
func (m *Module) fn1480(v0, v1, v2, v3 int32) {
	var v4, v5 int32
	t0 := m.g0
	v4 = t0 - i32(32)
	m.g0 = v4
	m.fn1485(v0, v1)
	t1 := int32(load32(m.memory[int64(uint32(v0))+16:]))
	t2 := int32(load32(m.memory[int64(uint32(v0))+20:]))
	v5 = v1 + i32(-1)
	t3 := m.fn857(t1, t2, v5, i32(1075928))
	v1 = t3
	t4 := int32(load32(m.memory[int64(uint32(v0))+28:]))
	t5 := int32(load32(m.memory[int64(uint32(v0))+32:]))
	t6 := m.fn1491(t4, t5, v5, i32(1075944))
	v0 = t6
	if v2&i32(255) == i32(2) {
		goto l0
	}
	m.fn1445(v0, v2&i32(1), v3, v1)
	goto l1
l0:
	m.fn1333(v0, v1)
	{
		t7 := int32(load32(m.memory[int64(uint32(v3))+4:]))
		t8 := int32(load32(m.memory[int64(uint32(v3))+8:]))
		t9 := m.fn23(t7, t8)
		if t9 != 0 {
			goto l2
		}
		store32(m.memory[uint32(v4):], uint32(i32(-0x80000000)))
		t10 := int32(load32(m.memory[int64(uint32(v3))+8:]))
		store32(m.memory[int64(uint32(v4))+12:], uint32(t10))
		t11 := int64(load64(m.memory[uint32(v3):]))
		store64(m.memory[int64(uint32(v4))+4:], uint64(t11))
		m.fn338(v1, v4)
		goto l1
	}
l2:
	m.fn894(v3)
l1:
	m.g0 = v4 + i32(32)
}
func (m *Module) fn1481(v0, v1, v2 int32) {
	var v3 int32
	t0 := m.g0
	v3 = t0 - i32(16)
	m.g0 = v3
	{
		if v1 != 0 {
			goto l0
		}
		v1 = i32(0)
		goto l1
	l0:
		m.fn46(v3+i32(8), v1, v2)
		t1 := int32(load32(m.memory[int64(uint32(v3))+12:]))
		v2 = t1
		t2 := int32(load32(m.memory[int64(uint32(v3))+8:]))
		v1 = t2
	}
l1:
	t4 := v0
	p3 := i32(0)
	if v2 != 0 {
		p3 = v1
	}
	p5 := i32(0)
	if v1 != 0 {
		p5 = p3
	}
	m.fn1041(t4, p5, v2)
	m.g0 = v3 + i32(16)
}
