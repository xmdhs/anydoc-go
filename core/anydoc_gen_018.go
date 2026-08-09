package core

import (
	"math"
	"math/bits"
)

func (m *Module) fn762(v0, v1, v2, v3, v4, v5 int32) {
	t0 := int32(load32(m.memory[uint32(v1):]))
	if t0 != 0 {
		m.fn349(i32(1073116))
		panic("unreachable")
	}
	store32(m.memory[uint32(v1):], uint32(i32(-1)))
	m.fn439(v0, v1+i32(8), v2, v3, v4, v5)
	t1 := int32(load32(m.memory[uint32(v1):]))
	store32(m.memory[uint32(v1):], uint32(t1+i32(1)))
}
func (m *Module) fn763(v0, v1, v2 int32) {
	var v3, v4 int32
	v3 = i32(0)
	{
		{
			t0 := m.fn306(v1, v2, i32(1070929), i32(29), i32(1070958), i32(9))
			v2 = t0
			if v2 != 0 {
				goto l0
			}
			goto l1
		}
	l0:
		{
			t1 := int32(load32(m.memory[int64(uint32(v2))+20:]))
			v1 = t1
			if v1 != 0 {
				goto l2
			}
			goto l1
		}
	l2:
		v1 = v1 << 5
		t2 := int32(load32(m.memory[int64(uint32(v2))+16:]))
		v2 = t2
	l5:
		{
			t3 := int32(load32(m.memory[uint32(v2+i32(8)):]))
			if t3 != i32(2) {
				goto l3
			}
			t4 := int32(load32(m.memory[uint32(v2+i32(4)):]))
			t5 := int32(load16(m.memory[uint32(t4):]))
			if t5 != i32(25705) {
				goto l3
			}
			t6 := int32(load32(m.memory[uint32(v2+i32(24)):]))
			v4 = t6
			if v4 == 0 {
				goto l3
			}
			t7 := int32(load32(m.memory[uint32(v2+i32(28)):]))
			if t7 != i32(67) {
				goto l3
			}
			t8 := m.fn973(v4+i32(8), i32(1070500), i32(67))
			if t8 == 0 {
				goto l4
			}
		}
	l3:
		v2 = v2 + i32(32)
		v1 = v1 + i32(-32)
		if v1 != 0 {
			goto l5
		}
		goto l1
	l4:
		t9 := int32(load32(m.memory[int64(uint32(v2))+20:]))
		v1 = t9
		t10 := int32(load32(m.memory[int64(uint32(v2))+16:]))
		v3 = t10
	}
l1:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
	store32(m.memory[uint32(v0):], uint32(v3))
}
func (m *Module) fn764(v0, v1 int32) {
	var v2, v3, v4 int32
	var v5 int64
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	{
		{
			t1 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			if t1 != 0 {
				t7 := int32(load32(m.memory[int64(uint32(v0))+20:]))
				if t7 != 0 {
					t9 := int32(load32(m.memory[int64(uint32(v0))+20:]))
					v3 = t9
					store32(m.memory[int64(uint32(v0))+20:], uint32(i32(0)))
					t10 := int64(load64(m.memory[int64(uint32(v0))+12:]))
					v5 = t10
					store64(m.memory[int64(uint32(v0))+12:], uint64(i64(0x400000000)))
					store32(m.memory[int64(uint32(v2))+8:], uint32(v3))
					store64(m.memory[uint32(v2):], uint64(v5))
					{
						t11 := int32(load32(m.memory[int64(uint32(v0))+8:]))
						v4 = t11
						t12 := int32(load32(m.memory[uint32(v0):]))
						if v4 != t12 {
							goto l7
						}
						m.fn311(v0)
					}
				l7:
					t13 := v0
					v3 = v4 + i32(1)
					store32(m.memory[int64(uint32(t13))+8:], uint32(v3))
					t14 := int32(load32(m.memory[int64(uint32(v0))+4:]))
					v4 = t14 + v4<<4
					store32(m.memory[uint32(v4):], uint32(i32(0)))
					t15 := int64(load64(m.memory[uint32(v2):]))
					store64(m.memory[int64(uint32(v4))+4:], uint64(t15))
					t16 := int32(load32(m.memory[int64(uint32(v2))+8:]))
					store32(m.memory[int64(uint32(v4))+12:], uint32(t16))
					goto l6
				}
				t8 := int32(load32(m.memory[int64(uint32(v0))+8:]))
				v3 = t8
				goto l6
			}
			t2 := int32(load32(m.memory[uint32(v1):]))
			v0 = t2
			if v0 == 0 {
				goto l1
			}
			t3 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			v3 = t3
			t4 := int32(load32(m.memory[uint32(v3+i32(-4)):]))
			v1 = t4
			v4 = v1 & i32(-8)
			t5 := v4
			v1 = v1 & i32(3)
			p6 := i32(8)
			if v1 != 0 {
				p6 = i32(4)
			}
			v0 = v0 << 5
			if uint32(t5) < uint32(p6|v0) {
				m.fn3(i32(1274224), i32(46), i32(1274272))
				panic("unreachable")
			}
			if v1 == 0 {
				goto l3
			}
			if uint32(v4) > uint32(v0+i32(39)) {
				m.fn3(i32(1274288), i32(46), i32(1274336))
				panic("unreachable")
			}
		l3:
			m.fn1(v3)
			goto l1
		}
	l6:
		{
			t17 := int32(load32(m.memory[uint32(v0):]))
			if v3 != t17 {
				goto l8
			}
			m.fn311(v0)
		}
	l8:
		store32(m.memory[int64(uint32(v0))+8:], uint32(v3+i32(1)))
		t18 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		v0 = t18 + v3<<4
		store32(m.memory[uint32(v0):], uint32(i32(1)))
		t19 := int64(load64(m.memory[uint32(v1):]))
		store64(m.memory[int64(uint32(v0))+4:], uint64(t19))
		t20 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		store32(m.memory[int64(uint32(v0))+12:], uint32(t20))
	}
l1:
	m.g0 = v2 + i32(16)
}
func (m *Module) fn765(v0, v1, v2, v3, v4 int32) {
	var v5, v6, v7, v8 int32
	if v1 == 0 {
		return
	}
	v5 = v0 + v1<<5
l13:
	v6 = v0
	{
		t0 := int32(load32(m.memory[uint32(v0):]))
		v1 = t0
		switch v1 >> 31 & (v1 + i32(-0x7fffffff)) {
		case 5, 6:
			goto l6
		case 2:
			t1 := int32(load32(m.memory[int64(uint32(v0))+24:]))
			v1 = t1
			if v1 == 0 {
				goto l6
			}
			v6 = v1 * i32(28)
			t2 := int32(load32(m.memory[int64(uint32(v0))+20:]))
			v1 = t2 + i32(8)
		l7:
			{
				t3 := int32(load32(m.memory[uint32(v1+i32(-4)):]))
				t4 := int32(load32(m.memory[uint32(v1):]))
				m.fn765(t3, t4, v2, v3, v4)
				v1 = v1 + i32(28)
				v6 = v6 + i32(-28)
				if v6 != 0 {
					goto l7
				}
				goto l6
			}
		case 3:
			t5 := int32(load32(m.memory[int64(uint32(v0))+12:]))
			v1 = t5
			if v1 == 0 {
				goto l6
			}
			t6 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			v7 = t6
			v8 = v7 + v1*i32(12)
			goto l12
		case 4:
			t7 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			t8 := int32(load32(m.memory[int64(uint32(v0))+12:]))
			m.fn765(t7, t8, v2, v3, v4)
			goto l6
		case 1:
			v6 = v0 + i32(4)
			fallthrough
		default:
			t9 := int32(load32(m.memory[int64(uint32(v6))+4:]))
			t10 := int32(load32(m.memory[int64(uint32(v6))+8:]))
			m.fn778(t9, t10, v2, v3, v4)
			goto l6
		}
	}
l12:
	{
		t11 := int32(load32(m.memory[int64(uint32(v7))+8:]))
		v1 = t11
		if v1 == 0 {
			goto l9
		}
		v6 = v1 * i32(20)
		t12 := int32(load32(m.memory[int64(uint32(v7))+4:]))
		v1 = t12
	l11:
		{
			t13 := int32(load32(m.memory[uint32(v1):]))
			if t13 == i32(-1) {
				goto l10
			}
			t14 := int32(load32(m.memory[uint32(v1+i32(4)):]))
			t15 := int32(load32(m.memory[uint32(v1+i32(8)):]))
			m.fn765(t14, t15, v2, v3, v4)
		}
	l10:
		v1 = v1 + i32(20)
		v6 = v6 + i32(-20)
		if v6 != 0 {
			goto l11
		}
	}
l9:
	v7 = v7 + i32(12)
	if v7 != v8 {
		goto l12
	}
l6:
	v0 = v0 + i32(32)
	if v0 != v5 {
		goto l13
	}
}
func (m *Module) fn766(v0, v1, v2 int32) {
	var v3, v4 int32
	var v5 int64
	var v6, v7 int32
	var v8, v9 int64
	var v10, v11, v12 int32
	var v13 int64
	var v14, v15 int32
	t0 := int64(load64(m.memory[int64(uint32(v0))+16:]))
	t1 := int64(load64(m.memory[int64(uint32(v0))+24:]))
	t2 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	v3 = t2
	t3 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	t4 := v3
	v4 = t3
	t5 := m.fn64(t0, t1, t4, v4)
	v5 = t5
	{
		t6 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		if t6 != 0 {
			goto l0
		}
		_ = m.fn77(v0, i32(1), v0+i32(16))
	}
l0:
	t8 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	v6 = t8
	v7 = v6 & int32(v5)
	v8 = int64(uint64(v5) >> 25)
	v9 = v8 & i64(127) * i64(72340172838076673)
	t9 := int32(load32(m.memory[uint32(v0):]))
	v10 = t9
	v11 = i32(0)
	v12 = i32(0)
l14:
	{
		t10 := int64(load64(m.memory[uint32(v10+v7):]))
		v13 = t10
		v5 = v13 ^ v9
		v5 = (v5 ^ i64(-1)) & (v5 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
		if v5 == 0 {
			goto l1
		}
	l4:
		{
			t11 := v4
			v14 = v10 - (int32(uint32(int64(bits.TrailingZeros64(uint64(v5))))>>3)+v7)&v6<<4
			t12 := int32(load32(m.memory[uint32(v14+i32(-8)):]))
			if t11 != t12 {
				goto l2
			}
			t13 := int32(load32(m.memory[uint32(v14+i32(-12)):]))
			t14 := m.fn973(v3, t13, v4)
			if t14 == 0 {
				store32(m.memory[uint32(v14+i32(-4)):], uint32(v2))
				{
					t24 := int32(load32(m.memory[uint32(v1):]))
					v0 = t24
					if v0 == 0 {
						return
					}
					t25 := int32(load32(m.memory[uint32(v3+i32(-4)):]))
					v1 = t25
					v10 = v1 & i32(-8)
					t26 := v10
					v1 = v1 & i32(3)
					p27 := i32(8)
					if v1 != 0 {
						p27 = i32(4)
					}
					if uint32(t26) < uint32(p27+v0) {
						m.fn3(i32(1274224), i32(46), i32(1274272))
						panic("unreachable")
					}
					if v1 == 0 {
						goto l12
					}
					if uint32(v10) > uint32(v0+i32(39)) {
						m.fn3(i32(1274288), i32(46), i32(1274336))
						panic("unreachable")
					}
				l12:
					m.fn1(v3)
				}
				return
			}
		}
	l2:
		v5 = (v5 + i64(-1)) & v5
		if !(v5 == 0) {
			goto l4
		}
	}
l1:
	v5 = v13 & i64(-0x7f7f7f7f7f7f7f80)
	if v11 == i32(1) {
		goto l5
	}
	if v5 == 0 {
		v11 = i32(0)
		goto l8
	}
	v15 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v5))))>>3) + v7) & v6
l5:
	if v5&(v13<<1) != i64(0) {
		{
			t15 := int32(int8(m.memory[uint32(v10+v15)]))
			v7 = t15
			if v7 < i32(0) {
				goto l9
			}
			t16 := int64(load64(m.memory[uint32(v10):]))
			t17 := v10
			v15 = int32(uint32(int64(bits.TrailingZeros64(uint64(t16&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
			t18 := int32(m.memory[uint32(t17+v15)])
			v7 = t18
		}
	l9:
		t19 := v10 + v15
		v3 = int32(v8) & i32(127)
		m.memory[uint32(t19)] = byte(v3)
		m.memory[uint32(v10+(v15+i32(-8))&v6+i32(8))] = byte(v3)
		t20 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		store32(m.memory[int64(uint32(v0))+8:], uint32(t20-v7&i32(1)))
		t21 := int32(load32(m.memory[int64(uint32(v0))+12:]))
		store32(m.memory[int64(uint32(v0))+12:], uint32(t21+i32(1)))
		v0 = v10 - v15<<4
		store32(m.memory[uint32(v0+i32(-4)):], uint32(v2))
		v0 = v0 + i32(-16)
		t22 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		store32(m.memory[int64(uint32(v0))+8:], uint32(t22))
		t23 := int64(load64(m.memory[uint32(v1):]))
		store64(m.memory[uint32(v0):], uint64(t23))
		return
	}
	v11 = i32(1)
	goto l8
l8:
	v12 = v12 + i32(8)
	v7 = (v12 + v7) & v6
	goto l14
}
func (m *Module) fn767(v0, v1, v2 int32) {
	var v3 int32
	if v1 == 0 {
		return
	}
	v1 = v1 * i32(28)
l3:
	{
		t0 := int32(load32(m.memory[uint32(v0):]))
		v3 = t0
		if uint32(v3) >= uint32(i32(3)) {
			goto l1
		}
		{
			if v3 != i32(2) {
				goto l2
			}
			t1 := int32(load32(m.memory[uint32(v0+i32(8)):]))
			t2 := int32(load32(m.memory[uint32(v0+i32(12)):]))
			_ = m.fn740(v2, t1, t2)
		}
	l2:
		t4 := int32(load32(m.memory[uint32(v0+i32(20)):]))
		t5 := int32(load32(m.memory[uint32(v0+i32(24)):]))
		m.fn767(t4, t5, v2)
	}
l1:
	v0 = v0 + i32(28)
	v1 = v1 + i32(-28)
	if v1 != 0 {
		goto l3
	}
}
func (m *Module) fn768(v0 int32) int32 {
	var v1, v2, v3, v4, v5 int32
	v1 = i32(0)
	p0 := i32(25)
	if uint32(v0) < uint32(i32(92729)) {
		p0 = i32(0)
	}
	v2 = p0
	t1 := v2
	v2 = v2 + i32(13)
	t2 := int32(load32(m.memory[int64(uint32(v2<<2))+1102768:]))
	t3 := v2
	t4 := t2 << 11
	v2 = v0 << 11
	p5 := t3
	if uint32(t4) > uint32(v2) {
		p5 = t1
	}
	v3 = p5
	t6 := v3
	v3 = v3 + i32(6)
	t7 := int32(load32(m.memory[int64(uint32(v3<<2))+1102768:]))
	p8 := v3
	if uint32(t7<<11) > uint32(v2) {
		p8 = t6
	}
	v3 = p8
	t9 := v3
	v3 = v3 + i32(3)
	t10 := int32(load32(m.memory[int64(uint32(v3<<2))+1102768:]))
	p11 := v3
	if uint32(t10<<11) > uint32(v2) {
		p11 = t9
	}
	v3 = p11
	t12 := v3
	v3 = v3 + i32(2)
	t13 := int32(load32(m.memory[int64(uint32(v3<<2))+1102768:]))
	p14 := v3
	if uint32(t13<<11) > uint32(v2) {
		p14 = t12
	}
	v3 = p14
	t15 := v3
	v3 = v3 + i32(1)
	t16 := int32(load32(m.memory[int64(uint32(v3<<2))+1102768:]))
	p17 := v3
	if uint32(t16<<11) > uint32(v2) {
		p17 = t15
	}
	v3 = p17
	t18 := int32(load32(m.memory[int64(uint32(v3<<2))+1102768:]))
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
	v5 = v2 + i32(1102768)
	t21 := int32(load32(m.memory[int64(uint32(v2))+1102768:]))
	v2 = int32(uint32(t21) >> 21)
	v4 = i32(1519)
	{
		{
			if uint32(v3) > uint32(i32(49)) {
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
		t24 := int32(m.memory[uint32(v2+i32(1095340))])
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
func (m *Module) fn769(v0 int32) int32 {
	var v1, v2, v3, v4, v5 int32
	v1 = i32(0)
	p0 := i32(21)
	if uint32(v0) < uint32(i32(70736)) {
		p0 = i32(0)
	}
	v2 = p0
	t1 := v2
	v2 = v2 + i32(11)
	t2 := int32(load32(m.memory[int64(uint32(v2<<2))+1106424:]))
	t3 := v2
	t4 := t2 << 11
	v2 = v0 << 11
	p5 := t3
	if uint32(t4) > uint32(v2) {
		p5 = t1
	}
	v3 = p5
	t6 := v3
	v3 = v3 + i32(5)
	t7 := int32(load32(m.memory[int64(uint32(v3<<2))+1106424:]))
	p8 := v3
	if uint32(t7<<11) > uint32(v2) {
		p8 = t6
	}
	v3 = p8
	t9 := v3
	v3 = v3 + i32(3)
	t10 := int32(load32(m.memory[int64(uint32(v3<<2))+1106424:]))
	p11 := v3
	if uint32(t10<<11) > uint32(v2) {
		p11 = t9
	}
	v3 = p11
	t12 := v3
	v3 = v3 + i32(1)
	t13 := int32(load32(m.memory[int64(uint32(v3<<2))+1106424:]))
	p14 := v3
	if uint32(t13<<11) > uint32(v2) {
		p14 = t12
	}
	v3 = p14
	t15 := v3
	v3 = v3 + i32(1)
	t16 := int32(load32(m.memory[int64(uint32(v3<<2))+1106424:]))
	p17 := v3
	if uint32(t16<<11) > uint32(v2) {
		p17 = t15
	}
	v3 = p17
	t18 := int32(load32(m.memory[int64(uint32(v3<<2))+1106424:]))
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
	v5 = v2 + i32(1106424)
	t21 := int32(load32(m.memory[int64(uint32(v2))+1106424:]))
	v2 = int32(uint32(t21) >> 21)
	v4 = i32(291)
	{
		{
			if uint32(v3) > uint32(i32(41)) {
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
		t24 := int32(m.memory[uint32(v2+i32(1098545))])
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
func (m *Module) fn770(v0, v1, v2 int32) {
	var v3, v4, v5, v6, v7, v8 int32
	var v9 int64
	var v10 int32
	var v11 int64
	var v12 int32
	var v13 int64
	var v14 int32
	t0 := m.g0
	v3 = t0 - i32(64)
	m.g0 = v3
	{
		{
			t1 := int32(load32(m.memory[int64(uint32(v2))+8:]))
			v4 = t1
			if v4 != 0 {
				goto l0
			}
			v5 = i32(1)
			goto l1
		}
	l0:
		t2 := int32(load32(m.memory[int64(uint32(v2))+4:]))
		v6 = t2
		t3 := m.fn7(v4)
		v5 = t3
		if v5 == 0 {
			m.fn12(i32(1), v4)
			panic("unreachable")
		}
		if v4 == 0 {
			goto l1
		}
		memory_copy(m.memory, uint32(v5), uint32(v6), uint32(v4))
	}
l1:
	store32(m.memory[int64(uint32(v3))+32:], uint32(v4))
	store32(m.memory[int64(uint32(v3))+28:], uint32(v5))
	store32(m.memory[int64(uint32(v3))+24:], uint32(v4))
	{
		{
			{
				t4 := m.fn442(v1, v3+i32(24))
				if t4 == 0 {
					{
						{
							t20 := int32(load32(m.memory[int64(uint32(v2))+8:]))
							v4 = t20
							if v4 != 0 {
								goto l10
							}
							v5 = i32(1)
							goto l11
						}
					l10:
						t21 := int32(load32(m.memory[int64(uint32(v2))+4:]))
						v6 = t21
						t22 := m.fn7(v4)
						v5 = t22
						if v5 == 0 {
							m.fn12(i32(1), v4)
							panic("unreachable")
						}
						if v4 == 0 {
							goto l11
						}
						memory_copy(m.memory, uint32(v5), uint32(v6), uint32(v4))
					}
				l11:
					store32(m.memory[int64(uint32(v3))+20:], uint32(v4))
					store32(m.memory[int64(uint32(v3))+16:], uint32(v5))
					store32(m.memory[int64(uint32(v3))+12:], uint32(v4))
					m.fn779(v3+i32(24), v1+i32(32), v3+i32(12))
					{
						t23 := int32(load32(m.memory[int64(uint32(v3))+32:]))
						if t23 == i32(-1) {
							goto l13
						}
						t24 := int64(load64(m.memory[int64(uint32(v3))+24:]))
						v9 = t24
						t25 := v3
						v4 = v3 + i32(32)
						t26 := int32(load32(m.memory[int64(uint32(v4))+8:]))
						store32(m.memory[int64(uint32(t25))+56:], uint32(t26))
						t27 := int64(load64(m.memory[uint32(v4):]))
						store64(m.memory[int64(uint32(v3))+48:], uint64(t27))
						{
							t28 := int32(load32(m.memory[int64(uint32(v3))+44:]))
							v4 = t28
							t29 := int32(load32(m.memory[uint32(v4):]))
							v5 = t29
							t30 := int32(load32(m.memory[int64(uint32(v4))+4:]))
							t31 := v5
							v8 = t30
							t32 := v8
							v1 = int32(v9)
							v6 = t32 & v1
							t33 := int64(load64(m.memory[uint32(t31+v6):]))
							v9 = t33 & i64(-0x7f7f7f7f7f7f7f80)
							if v9 != i64(0) {
								goto l14
							}
							v14 = i32(8)
						l15:
							{
								v6 = v6 + v14
								v14 = v14 + i32(8)
								t34 := v5
								v6 = v6 & v8
								t35 := int64(load64(m.memory[uint32(t34+v6):]))
								v9 = t35 & i64(-0x7f7f7f7f7f7f7f80)
								if v9 == 0 {
									goto l15
								}
							}
						}
					l14:
						{
							t36 := v5
							v6 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v9))))>>3) + v6) & v8
							t37 := int32(int8(m.memory[uint32(t36+v6)]))
							v14 = t37
							if v14 < i32(0) {
								goto l16
							}
							t38 := int64(load64(m.memory[uint32(v5):]))
							t39 := v5
							v6 = int32(uint32(int64(bits.TrailingZeros64(uint64(t38&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
							t40 := int32(m.memory[uint32(t39+v6)])
							v14 = t40
						}
					l16:
						t41 := v5 + v6
						v1 = int32(uint32(v1) >> 25)
						m.memory[uint32(t41)] = byte(v1)
						m.memory[uint32(v5+(v6+i32(-8))&v8+i32(8))] = byte(v1)
						t42 := int32(load32(m.memory[int64(uint32(v4))+8:]))
						store32(m.memory[int64(uint32(v4))+8:], uint32(t42-v14&i32(1)))
						t43 := int32(load32(m.memory[int64(uint32(v4))+12:]))
						store32(m.memory[int64(uint32(v4))+12:], uint32(t43+i32(1)))
						v4 = v5 - v6<<4
						v5 = v4 + i32(-16)
						t44 := int64(load64(m.memory[int64(uint32(v3))+48:]))
						store64(m.memory[uint32(v5):], uint64(t44))
						t45 := int32(load32(m.memory[int64(uint32(v3))+56:]))
						store32(m.memory[int64(uint32(v5))+8:], uint32(t45))
						store32(m.memory[uint32(v4+i32(-4)):], uint32(i32(1)))
					}
				l13:
					t46 := int32(load32(m.memory[int64(uint32(v2))+8:]))
					store32(m.memory[int64(uint32(v0))+8:], uint32(t46))
					t47 := int64(load64(m.memory[uint32(v2):]))
					store64(m.memory[uint32(v0):], uint64(t47))
					goto l17
				}
				v4 = i32(1)
				t5 := int32(load32(m.memory[int64(uint32(v1))+44:]))
				if t5 == 0 {
					goto l4
				}
				t6 := int64(load64(m.memory[int64(uint32(v1))+48:]))
				t7 := int64(load64(m.memory[int64(uint32(v1))+56:]))
				t8 := int32(load32(m.memory[int64(uint32(v2))+4:]))
				v7 = t8
				t9 := int32(load32(m.memory[int64(uint32(v2))+8:]))
				t10 := v7
				v8 = t9
				t11 := m.fn64(t6, t7, t10, v8)
				v9 = t11
				t12 := int32(load32(m.memory[int64(uint32(v1))+36:]))
				v10 = t12
				v5 = v10 & int32(v9)
				v11 = int64(uint64(v9)>>25) & i64(127) * i64(72340172838076673)
				t13 := int32(load32(m.memory[int64(uint32(v1))+32:]))
				v6 = t13
				v12 = i32(0)
			l9:
				{
					{
						t14 := int64(load64(m.memory[uint32(v6+v5):]))
						v13 = t14
						v9 = v13 ^ v11
						v9 = (v9 ^ i64(-1)) & (v9 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
						if v9 == 0 {
							goto l5
						}
					l8:
						{
							t15 := v8
							v14 = v6 - (int32(uint32(int64(bits.TrailingZeros64(uint64(v9))))>>3)+v5)&v10<<4
							t16 := int32(load32(m.memory[uint32(v14+i32(-8)):]))
							if t15 != t16 {
								goto l6
							}
							t17 := int32(load32(m.memory[uint32(v14+i32(-12)):]))
							t18 := m.fn973(v7, t17, v8)
							if t18 == 0 {
								goto l7
							}
						}
					l6:
						v9 = (v9 + i64(-1)) & v9
						if !(v9 == 0) {
							goto l8
						}
					}
				l5:
					if !(v13&(v13<<1)&i64(-0x7f7f7f7f7f7f7f80) == 0) {
						goto l4
					}
					t19 := v5
					v12 = v12 + i32(8)
					v5 = (t19 + v12) & v10
					goto l9
				}
			}
		l7:
			t48 := int32(load32(m.memory[uint32(v14+i32(-4)):]))
			v4 = t48
		}
	l4:
		v10 = v1 + i32(32)
		store32(m.memory[int64(uint32(v3))+8:], uint32(v4))
		v9 = int64(uint32(i32(3)))<<32 | int64(uint32(v3+i32(8)))
		v11 = int64(uint32(i32(17)))<<32 | int64(uint32(v2))
	l22:
		{
			store64(m.memory[int64(uint32(v3))+32:], uint64(v9))
			store64(m.memory[int64(uint32(v3))+24:], uint64(v11))
			m.fn13(v3+i32(48), i32(1048811), v3+i32(24))
			t49 := int32(load32(m.memory[int64(uint32(v3))+48:]))
			v6 = t49
			t50 := int32(load32(m.memory[int64(uint32(v3))+52:]))
			v8 = t50
			t51 := int32(load32(m.memory[int64(uint32(v3))+8:]))
			v14 = t51
			if v14 == i32(-1) {
				store32(m.memory[uint32(v0):], uint32(i32(-1)))
				if v6 == 0 {
					goto l26
				}
				m.fn17(v8, v6, i32(1))
			l26:
				t58 := int32(load32(m.memory[uint32(v2):]))
				v4 = t58
				if v4 == 0 {
					goto l17
				}
				t59 := int32(load32(m.memory[int64(uint32(v2))+4:]))
				m.fn17(t59, v4, i32(1))
				goto l17
			}
			t52 := int32(load32(m.memory[int64(uint32(v3))+56:]))
			v4 = t52
			v5 = i32(1)
			store32(m.memory[int64(uint32(v3))+8:], uint32(v14+i32(1)))
			{
				if v4 == 0 {
					goto l19
				}
				t53 := m.fn7(v4)
				v5 = t53
				if v5 == 0 {
					m.fn12(i32(1), v4)
					panic("unreachable")
				}
				if v4 == 0 {
					goto l19
				}
				memory_copy(m.memory, uint32(v5), uint32(v8), uint32(v4))
			}
		l19:
			store32(m.memory[int64(uint32(v3))+32:], uint32(v4))
			store32(m.memory[int64(uint32(v3))+28:], uint32(v5))
			store32(m.memory[int64(uint32(v3))+24:], uint32(v4))
			t54 := m.fn442(v1, v3+i32(24))
			if t54 == 0 {
				goto l21
			}
			if v6 == 0 {
				goto l22
			}
			t55 := int32(load32(m.memory[uint32(v8+i32(-4)):]))
			v4 = t55
			v5 = v4 & i32(-8)
			t56 := v5
			v4 = v4 & i32(3)
			p57 := i32(8)
			if v4 != 0 {
				p57 = i32(4)
			}
			if uint32(t56) < uint32(p57+v6) {
				m.fn3(i32(1274224), i32(46), i32(1274272))
				panic("unreachable")
			}
			if v4 == 0 {
				goto l24
			}
			if uint32(v5) > uint32(v6+i32(39)) {
				m.fn3(i32(1274288), i32(46), i32(1274336))
				panic("unreachable")
			}
		l24:
			m.fn1(v8)
			goto l22
		}
	l21:
		t60 := int32(load32(m.memory[int64(uint32(v2))+8:]))
		store32(m.memory[int64(uint32(v3))+32:], uint32(t60))
		t61 := int64(load64(m.memory[uint32(v2):]))
		store64(m.memory[int64(uint32(v3))+24:], uint64(t61))
		t62 := int32(load32(m.memory[int64(uint32(v3))+8:]))
		m.fn766(v10, v3+i32(24), t62)
		{
			if v4 != 0 {
				goto l27
			}
			v5 = i32(1)
			goto l28
		l27:
			t63 := m.fn7(v4)
			v5 = t63
			if v5 == 0 {
				m.fn12(i32(1), v4)
				panic("unreachable")
			}
			if v4 == 0 {
				goto l28
			}
			memory_copy(m.memory, uint32(v5), uint32(v8), uint32(v4))
		}
	l28:
		store32(m.memory[int64(uint32(v3))+20:], uint32(v4))
		store32(m.memory[int64(uint32(v3))+16:], uint32(v5))
		store32(m.memory[int64(uint32(v3))+12:], uint32(v4))
		m.fn779(v3+i32(24), v10, v3+i32(12))
		{
			t64 := int32(load32(m.memory[int64(uint32(v3))+32:]))
			if t64 == i32(-1) {
				goto l30
			}
			t65 := int64(load64(m.memory[int64(uint32(v3))+24:]))
			v9 = t65
			t66 := v3
			v5 = v3 + i32(32)
			t67 := int32(load32(m.memory[int64(uint32(v5))+8:]))
			store32(m.memory[int64(uint32(t66))+56:], uint32(t67))
			t68 := int64(load64(m.memory[uint32(v5):]))
			store64(m.memory[int64(uint32(v3))+48:], uint64(t68))
			{
				t69 := int32(load32(m.memory[int64(uint32(v3))+44:]))
				v5 = t69
				t70 := int32(load32(m.memory[uint32(v5):]))
				v14 = t70
				t71 := int32(load32(m.memory[int64(uint32(v5))+4:]))
				t72 := v14
				v2 = t71
				t73 := v2
				v7 = int32(v9)
				v1 = t73 & v7
				t74 := int64(load64(m.memory[uint32(t72+v1):]))
				v9 = t74 & i64(-0x7f7f7f7f7f7f7f80)
				if v9 != i64(0) {
					goto l31
				}
				v10 = i32(8)
			l32:
				{
					v1 = v1 + v10
					v10 = v10 + i32(8)
					t75 := v14
					v1 = v1 & v2
					t76 := int64(load64(m.memory[uint32(t75+v1):]))
					v9 = t76 & i64(-0x7f7f7f7f7f7f7f80)
					if v9 == 0 {
						goto l32
					}
				}
			}
		l31:
			{
				t77 := v14
				v1 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v9))))>>3) + v1) & v2
				t78 := int32(int8(m.memory[uint32(t77+v1)]))
				v10 = t78
				if v10 < i32(0) {
					goto l33
				}
				t79 := int64(load64(m.memory[uint32(v14):]))
				t80 := v14
				v1 = int32(uint32(int64(bits.TrailingZeros64(uint64(t79&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
				t81 := int32(m.memory[uint32(t80+v1)])
				v10 = t81
			}
		l33:
			t82 := v14 + v1
			v7 = int32(uint32(v7) >> 25)
			m.memory[uint32(t82)] = byte(v7)
			m.memory[uint32(v14+(v1+i32(-8))&v2+i32(8))] = byte(v7)
			t83 := int32(load32(m.memory[int64(uint32(v5))+8:]))
			store32(m.memory[int64(uint32(v5))+8:], uint32(t83-v10&i32(1)))
			t84 := int32(load32(m.memory[int64(uint32(v5))+12:]))
			store32(m.memory[int64(uint32(v5))+12:], uint32(t84+i32(1)))
			v5 = v14 - v1<<4
			v14 = v5 + i32(-16)
			t85 := int64(load64(m.memory[int64(uint32(v3))+48:]))
			store64(m.memory[uint32(v14):], uint64(t85))
			t86 := int32(load32(m.memory[int64(uint32(v3))+56:]))
			store32(m.memory[int64(uint32(v14))+8:], uint32(t86))
			store32(m.memory[uint32(v5+i32(-4)):], uint32(i32(1)))
		}
	l30:
		store32(m.memory[int64(uint32(v0))+8:], uint32(v4))
		store32(m.memory[int64(uint32(v0))+4:], uint32(v8))
		store32(m.memory[uint32(v0):], uint32(v6))
	}
l17:
	m.g0 = v3 + i32(64)
}
func (m *Module) fn771(v0, v1, v2, v3 int32) {
	var v4 int32
	var v5 int64
	var v6, v7, v8 int32
	var v9 int64
	var v10, v11 int32
	var v12 int64
	if v3 <= i32(-1) {
		m.fn11()
		panic("unreachable")
	}
	if v3 != 0 {
		t0 := m.fn7(v3)
		v4 = t0
		if v4 != 0 {
			if v3 == 0 {
				goto l2
			}
			memory_copy(m.memory, uint32(v4), uint32(v2), uint32(v3))
			goto l2
		}
		m.fn12(i32(1), v3)
		panic("unreachable")
	}
	v4 = i32(1)
	goto l2
l2:
	t1 := int64(load64(m.memory[int64(uint32(v0))+16:]))
	t2 := int64(load64(m.memory[int64(uint32(v0))+24:]))
	t3 := m.fn64(t1, t2, v4, v3)
	v5 = t3
	t4 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	v6 = t4
	t5 := v6
	v7 = int32(v5)
	v8 = t5 & v7
	v9 = int64(uint64(v5)>>25) & i64(127) * i64(72340172838076673)
	t6 := int32(load32(m.memory[uint32(v0):]))
	v10 = t6
	v11 = i32(0)
l9:
	{
		{
			t7 := int64(load64(m.memory[uint32(v10+v8):]))
			v12 = t7
			v5 = v12 ^ v9
			v5 = (v5 ^ i64(-1)) & (v5 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
			if v5 == 0 {
				goto l4
			}
		l7:
			{
				v2 = v10 + (i32(0)-(int32(uint32(int64(bits.TrailingZeros64(uint64(v5))))>>3)+v8)&v6)*i32(28)
				t8 := int32(load32(m.memory[uint32(v2+i32(-20)):]))
				if t8 != v3 {
					goto l5
				}
				t9 := int32(load32(m.memory[uint32(v2+i32(-24)):]))
				t10 := m.fn973(t9, v4, v3)
				if t10 == 0 {
					v8 = i32(-1)
					{
						if v3 == 0 {
							goto l10
						}
						t12 := int32(load32(m.memory[uint32(v4+i32(-4)):]))
						v10 = t12
						v6 = v10 & i32(-8)
						t13 := v6
						v10 = v10 & i32(3)
						p14 := i32(8)
						if v10 != 0 {
							p14 = i32(4)
						}
						if uint32(t13) < uint32(p14+v3) {
							m.fn3(i32(1274224), i32(46), i32(1274272))
							panic("unreachable")
						}
						if v10 == 0 {
							goto l12
						}
						if uint32(v6) > uint32(v3+i32(39)) {
							m.fn3(i32(1274288), i32(46), i32(1274336))
							panic("unreachable")
						}
					l12:
						m.fn1(v4)
					}
				l10:
					goto l14
				}
			}
		l5:
			v5 = (v5 + i64(-1)) & v5
			if !(v5 == 0) {
				goto l7
			}
		}
	l4:
		if !(v12&(v12<<1)&i64(-0x7f7f7f7f7f7f7f80) == 0) {
			{
				t15 := int32(load32(m.memory[int64(uint32(v0))+8:]))
				if t15 != 0 {
					goto l15
				}
				_ = m.fn74(v0, v0+i32(16))
			}
		l15:
			v5 = int64(uint32(v3))<<32 | int64(uint32(v4))
			v2 = v7
			v8 = v3
			goto l14
		}
		t11 := v8
		v11 = v11 + i32(8)
		v8 = (t11 + v11) & v6
		goto l9
	}
l14:
	{
		{
			t17 := int32(load32(m.memory[uint32(v1+i32(8)):]))
			v3 = t17
			if v3 == 0 {
				goto l16
			}
			t18 := int32(load32(m.memory[uint32(v1+i32(4)):]))
			v4 = t18
			t19 := m.fn7(v3)
			v10 = t19
			if v10 == 0 {
				m.fn12(i32(1), v3)
				panic("unreachable")
			}
			if v3 == 0 {
				goto l18
			}
			memory_copy(m.memory, uint32(v10), uint32(v4), uint32(v3))
		l18:
			if v8 != i32(-1) {
				goto l19
			}
			t20 := int32(load32(m.memory[uint32(v10+i32(-4)):]))
			v0 = t20
			v4 = v0 & i32(-8)
			t21 := v4
			v0 = v0 & i32(3)
			p22 := i32(8)
			if v0 != 0 {
				p22 = i32(4)
			}
			if uint32(t21) < uint32(p22+v3) {
				m.fn3(i32(1274224), i32(46), i32(1274272))
				panic("unreachable")
			}
			if v0 == 0 {
				goto l21
			}
			if uint32(v4) > uint32(v3+i32(39)) {
				m.fn3(i32(1274288), i32(46), i32(1274336))
				panic("unreachable")
			}
		l21:
			m.fn1(v10)
			return
		}
	l16:
		v10 = i32(1)
		if v8 == i32(-1) {
			return
		}
	l19:
		{
			t23 := int32(load32(m.memory[uint32(v0):]))
			v4 = t23
			t24 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t25 := v4
			v6 = t24
			v1 = v6 & v2
			t26 := int64(load64(m.memory[uint32(t25+v1):]))
			v9 = t26 & i64(-0x7f7f7f7f7f7f7f80)
			if v9 != i64(0) {
				goto l24
			}
			v7 = i32(8)
		l25:
			{
				v1 = v1 + v7
				v7 = v7 + i32(8)
				t27 := v4
				v1 = v1 & v6
				t28 := int64(load64(m.memory[uint32(t27+v1):]))
				v9 = t28 & i64(-0x7f7f7f7f7f7f7f80)
				if v9 == 0 {
					goto l25
				}
			}
		}
	l24:
		{
			t29 := v4
			v1 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v9))))>>3) + v1) & v6
			t30 := int32(int8(m.memory[uint32(t29+v1)]))
			v7 = t30
			if v7 < i32(0) {
				goto l26
			}
			t31 := int64(load64(m.memory[uint32(v4):]))
			t32 := v4
			v1 = int32(uint32(int64(bits.TrailingZeros64(uint64(t31&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
			t33 := int32(m.memory[uint32(t32+v1)])
			v7 = t33
		}
	l26:
		t34 := v4 + v1
		v2 = int32(uint32(v2) >> 25)
		m.memory[uint32(t34)] = byte(v2)
		m.memory[uint32(v4+(v1+i32(-8))&v6+i32(8))] = byte(v2)
		t35 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		store32(m.memory[int64(uint32(v0))+8:], uint32(t35-v7&i32(1)))
		t36 := int32(load32(m.memory[int64(uint32(v0))+12:]))
		store32(m.memory[int64(uint32(v0))+12:], uint32(t36+i32(1)))
		v0 = v4 + (i32(0)-v1)*i32(28)
		m.memory[uint32(v0+i32(-4))] = byte(i32(0))
		store32(m.memory[uint32(v0+i32(-8)):], uint32(v3))
		store32(m.memory[uint32(v0+i32(-12)):], uint32(v10))
		store32(m.memory[uint32(v0+i32(-16)):], uint32(v3))
		store64(m.memory[uint32(v0+i32(-24)):], uint64(v5))
		store32(m.memory[uint32(v0+i32(-28)):], uint32(v8))
	}
}
func (m *Module) fn772(v0, v1, v2 int32) {
	var v3, v4, v5 int32
	{
		if v1 == 0 {
			return
		}
		v1 = v1 * i32(28)
		t0 := int32(load32(m.memory[int64(uint32(v2))+4:]))
		v3 = t0
		t1 := int32(load32(m.memory[uint32(v2):]))
		v4 = t1
	l4:
		{
			t2 := int32(load32(m.memory[uint32(v0):]))
			v5 = t2
			p3 := i32(1)
			if uint32(v5) > uint32(i32(2)) {
				p3 = v5 + i32(-3)
			}
			switch p3 + i32(-1) {
			default:
				goto l2
			case 0:
				t4 := int32(load32(m.memory[uint32(v0+i32(20)):]))
				t5 := int32(load32(m.memory[uint32(v0+i32(24)):]))
				m.fn772(t4, t5, v2)
				goto l2
			case 2:
				t6 := int32(load32(m.memory[uint32(v0+i32(8)):]))
				t7 := int32(load32(m.memory[uint32(v0+i32(12)):]))
				m.fn771(v4, v3, t6, t7)
			}
		}
	l2:
		v0 = v0 + i32(28)
		v1 = v1 + i32(-28)
		if v1 != 0 {
			goto l4
		}
	}
}
func (m *Module) fn773(v0, v1, v2 int32) {
	var v3, v4, v5, v6, v7, v8, v9, v10 int32
	var v11 int64
	var v12 int32
	var v13 int64
	var v14, v15 int32
	var v16 int64
	var v17, v18, v19, v20, v21, v22, v23, v24 int32
	var v25 int64
	t0 := m.g0
	v3 = t0 - i32(48)
	m.g0 = v3
	{
		if v1 == 0 {
			goto l0
		}
		v4 = v0 + v1*i32(28)
		t1 := int32(load32(m.memory[int64(uint32(v2))+8:]))
		v5 = t1
		t2 := int32(load32(m.memory[uint32(v2):]))
		v6 = t2
		t3 := int32(load32(m.memory[int64(uint32(v2))+4:]))
		v7 = t3
		v8 = v7 + i32(16)
	l76:
		{
			t4 := int32(load32(m.memory[uint32(v0):]))
			v1 = t4
			p5 := i32(1)
			if uint32(v1) > uint32(i32(2)) {
				p5 = v1 + i32(-3)
			}
			switch p5 + i32(-1) {
			default:
				goto l2
			case 0:
				t6 := int32(load32(m.memory[int64(uint32(v0))+20:]))
				t7 := int32(load32(m.memory[int64(uint32(v0))+24:]))
				m.fn773(t6, t7, v2)
				goto l2
			case 2:
				t8 := int32(load32(m.memory[int64(uint32(v6))+12:]))
				if t8 == 0 {
					goto l2
				}
				t9 := int64(load64(m.memory[int64(uint32(v6))+16:]))
				t10 := int64(load64(m.memory[int64(uint32(v6))+24:]))
				t11 := int32(load32(m.memory[int64(uint32(v0))+8:]))
				v9 = t11
				t12 := int32(load32(m.memory[int64(uint32(v0))+12:]))
				t13 := v9
				v10 = t12
				t14 := m.fn249(t9, t10, t13, v10)
				v11 = t14
				t15 := int32(load32(m.memory[int64(uint32(v6))+4:]))
				v12 = t15
				v1 = v12 & int32(v11)
				v13 = int64(uint64(v11)>>25) & i64(127) * i64(72340172838076673)
				t16 := int32(load32(m.memory[uint32(v6):]))
				v14 = t16
				v15 = i32(0)
			l8:
				{
					{
						t17 := int64(load64(m.memory[uint32(v14+v1):]))
						v16 = t17
						v11 = v16 ^ v13
						v11 = (v11 ^ i64(-1)) & (v11 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
						if v11 == 0 {
							goto l4
						}
					l7:
						{
							t18 := v10
							v17 = v14 - (int32(uint32(int64(bits.TrailingZeros64(uint64(v11))))>>3)+v1)&v12<<3
							t19 := int32(load32(m.memory[uint32(v17+i32(-4)):]))
							if t18 != t19 {
								goto l5
							}
							t20 := int32(load32(m.memory[uint32(v17+i32(-8)):]))
							t21 := m.fn973(v9, t20, v10)
							if t21 == 0 {
								{
									t23 := int32(load32(m.memory[int64(uint32(v7))+12:]))
									if t23 == 0 {
										goto l9
									}
									t24 := int64(load64(m.memory[int64(uint32(v7))+16:]))
									t25 := int64(load64(m.memory[int64(uint32(v7))+24:]))
									t26 := m.fn249(t24, t25, v9, v10)
									v11 = t26
									t27 := int32(load32(m.memory[int64(uint32(v7))+4:]))
									v12 = t27
									v1 = v12 & int32(v11)
									v13 = int64(uint64(v11)>>25) & i64(127) * i64(72340172838076673)
									t28 := int32(load32(m.memory[uint32(v7):]))
									v14 = t28
									v15 = i32(0)
								l13:
									{
										{
											t29 := int64(load64(m.memory[uint32(v14+v1):]))
											v16 = t29
											v11 = v16 ^ v13
											v11 = (v11 ^ i64(-1)) & (v11 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
											if v11 == 0 {
												goto l10
											}
										l12:
											{
												t30 := v10
												v17 = v14 + (i32(0)-(int32(uint32(int64(bits.TrailingZeros64(uint64(v11))))>>3)+v1)&v12)*i32(28)
												t31 := int32(load32(m.memory[uint32(v17+i32(-20)):]))
												if t30 != t31 {
													goto l11
												}
												t32 := int32(load32(m.memory[uint32(v17+i32(-24)):]))
												t33 := m.fn973(v9, t32, v10)
												if t33 == 0 {
													goto l2
												}
											}
										l11:
											v11 = (v11 + i64(-1)) & v11
											if !(v11 == 0) {
												goto l12
											}
										}
									l10:
										if !(v16&(v16<<1)&i64(-0x7f7f7f7f7f7f7f80) == 0) {
											goto l9
										}
										t34 := v1
										v15 = v15 + i32(8)
										v1 = (t34 + v15) & v12
										goto l13
									}
								}
							l9:
								{
									if v10 <= i32(-1) {
										goto l14
									}
									if v10 != 0 {
										t35 := m.fn7(v10)
										v18 = t35
										if v18 == 0 {
											m.fn12(i32(1), v10)
											panic("unreachable")
										}
										v12 = i32(0)
										store32(m.memory[int64(uint32(v3))+44:], uint32(i32(0)))
										store32(m.memory[int64(uint32(v3))+40:], uint32(v18))
										store32(m.memory[int64(uint32(v3))+36:], uint32(v10))
										v17 = v9 + v10
										v15 = i32(0)
										v1 = v9
									l31:
										{
											{
												if v15&i32(1) != 0 {
												l29:
													{
														{
															{
																t41 := int32(int8(m.memory[uint32(v1)]))
																v14 = t41
																if v14 <= i32(-1) {
																	goto l25
																}
																v1 = v1 + i32(1)
																v14 = v14 & i32(255)
																goto l26
															}
														l25:
															t42 := int32(m.memory[int64(uint32(v1))+1])
															v19 = t42 & i32(63)
															v15 = v14 & i32(31)
															if uint32(v14) > uint32(i32(-33)) {
																goto l27
															}
															v14 = v15<<6 | v19
															v1 = v1 + i32(2)
															goto l26
														l27:
															t43 := int32(m.memory[int64(uint32(v1))+2])
															v19 = v19<<6 | t43&i32(63)
															if uint32(v14) >= uint32(i32(-16)) {
																goto l28
															}
															v14 = v19 | v15<<12
															v1 = v1 + i32(3)
															goto l26
														l28:
															t44 := int32(m.memory[int64(uint32(v1))+3])
															v14 = v19<<6 | t44&i32(63) | v15<<18&i32(0x1c0000)
															v1 = v1 + i32(4)
														}
													l26:
														p45 := v14
														if uint32(v14+i32(-65)) < uint32(i32(26)) {
															p45 = v14 | i32(32)
														}
														v14 = p45
														if v14 == i32(95) {
															goto l24
														}
														if uint32(v14+i32(-97)) < uint32(i32(26)) {
															goto l24
														}
														if uint32(v14+i32(-48)) < uint32(i32(10)) {
															goto l24
														}
														if v1 == v17 {
															goto l16
														}
														goto l29
													}
												}
												{
													t36 := int32(int8(m.memory[uint32(v1)]))
													v14 = t36
													if v14 > i32(-1) {
														goto l19
													}
													t37 := int32(m.memory[int64(uint32(v1))+1])
													v15 = t37 & i32(63)
													v19 = v14 & i32(31)
													if uint32(v14) >= uint32(i32(-32)) {
														t38 := int32(m.memory[int64(uint32(v1))+2])
														v15 = v15<<6 | t38&i32(63)
														if uint32(v14) >= uint32(i32(-16)) {
															t39 := int32(m.memory[int64(uint32(v1))+3])
															v14 = v15<<6 | t39&i32(63) | v19<<18&i32(0x1c0000)
															v1 = v1 + i32(4)
															goto l21
														}
														v14 = v15 | v19<<12
														v1 = v1 + i32(3)
														goto l21
													}
													v14 = v19<<6 | v15
													v1 = v1 + i32(2)
													goto l21
												}
											l19:
												v1 = v1 + i32(1)
												v14 = v14 & i32(255)
											l21:
												p40 := v14
												if uint32(v14+i32(-65)) < uint32(i32(26)) {
													p40 = v14 | i32(32)
												}
												v15 = p40
												if uint32(v15+i32(-97)) < uint32(i32(26)) {
													goto l23
												}
												if uint32(v15+i32(-48)) < uint32(i32(10)) {
													goto l23
												}
												v14 = i32(45)
												if v15 == i32(45) {
													goto l23
												}
												if v15 == i32(95) {
													goto l23
												}
												goto l24
											}
										l23:
											v14 = v15
										l24:
											{
												t46 := int32(load32(m.memory[int64(uint32(v3))+36:]))
												if t46 != v12 {
													goto l30
												}
												m.fn196(v3+i32(36), v12, i32(1), i32(1), i32(1))
												t47 := int32(load32(m.memory[int64(uint32(v3))+40:]))
												v18 = t47
											}
										l30:
											;
											var p48 int32
											if v14 == i32(45) {
												p48 = 1
											}
											v15 = p48
											m.memory[uint32(v18+v12)] = byte(v14)
											t49 := v3
											v12 = v12 + i32(1)
											store32(m.memory[int64(uint32(t49))+44:], uint32(v12))
											if v1 == v17 {
												goto l16
											}
											goto l31
										}
									}
									store64(m.memory[int64(uint32(v3))+36:], uint64(i64(0x100000000)))
									v12 = i32(0)
									goto l16
								l16:
									t50 := int32(load32(m.memory[int64(uint32(v3))+40:]))
									v17 = t50
									v19 = v17 + v12
									v1 = i32(0)
								l38:
									v15 = v1
									if v15 != v12 {
										goto l32
									}
									v1 = v12
									v20 = i32(0)
									v15 = i32(0)
									goto l33
								l32:
									{
										{
											v1 = v17 + v15
											t51 := int32(int8(m.memory[uint32(v1)]))
											v14 = t51
											if v14 <= i32(-1) {
												goto l34
											}
											v1 = v1 + i32(1)
											v14 = v14 & i32(255)
											goto l35
										}
									l34:
										t52 := int32(m.memory[int64(uint32(v1))+1])
										v18 = t52 & i32(63)
										v20 = v14 & i32(31)
										if uint32(v14) > uint32(i32(-33)) {
											goto l36
										}
										v14 = v20<<6 | v18
										v1 = v1 + i32(2)
										goto l35
									l36:
										t53 := int32(m.memory[int64(uint32(v1))+2])
										v18 = v18<<6 | t53&i32(63)
										if uint32(v14) >= uint32(i32(-16)) {
											goto l37
										}
										v14 = v18 | v20<<12
										v1 = v1 + i32(3)
										goto l35
									l37:
										t54 := int32(m.memory[int64(uint32(v1))+3])
										v14 = v18<<6 | t54&i32(63) | v20<<18&i32(0x1c0000)
										v1 = v1 + i32(4)
									}
								l35:
									v1 = v1 - v19 + v12
									v20 = v1
									if v14 == i32(45) {
										goto l38
									}
								l33:
									v21 = v1 - (v17 + v1)
								l46:
									{
										t55 := v1
										v19 = v12
										if t55 != v19 {
											goto l39
										}
										v19 = v20
										goto l40
									}
								l39:
									{
										v18 = v17 + v19
										v12 = v18 + i32(-1)
										t56 := int32(int8(m.memory[uint32(v12)]))
										v14 = t56
										if v14 > i32(-1) {
											goto l41
										}
										{
											v12 = v18 + i32(-2)
											t57 := int32(m.memory[uint32(v12)])
											v22 = t57
											v23 = int32(int8(v22))
											if v23 < i32(-64) {
												goto l42
											}
											v18 = v22 & i32(31)
											goto l43
										}
									l42:
										{
											{
												v12 = v18 + i32(-3)
												t58 := int32(m.memory[uint32(v12)])
												v22 = t58
												v24 = int32(int8(v22))
												if v24 < i32(-64) {
													goto l44
												}
												v18 = v22 & i32(15)
												goto l45
											}
										l44:
											v12 = v18 + i32(-4)
											t59 := int32(m.memory[uint32(v12)])
											v18 = t59&i32(7)<<6 | v24&i32(63)
										}
									l45:
										v18 = v18<<6 | v23&i32(63)
									l43:
										v14 = v18<<6 | v14&i32(63)
									}
								l41:
									v12 = v21 + v12
									if v14 == i32(45) {
										goto l46
									}
								l40:
									{
										if v19 != v15 {
											v1 = v19 - v15
											if v1 <= i32(-1) {
												goto l14
											}
											t61 := m.fn7(v1)
											v14 = t61
											if v14 != 0 {
												if v1 == 0 {
													goto l50
												}
												memory_copy(m.memory, uint32(v14), uint32(v17+v15), uint32(v1))
												goto l50
											}
											m.fn12(i32(1), v1)
											panic("unreachable")
										}
										t60 := m.fn7(i32(6))
										v14 = t60
										if v14 != 0 {
											t62 := int32(load16(m.memory[int64(uint32(i32(0)))+1073957:]))
											store16(m.memory[int64(uint32(v14))+4:], uint16(t62))
											t63 := int32(load32(m.memory[int64(uint32(i32(0)))+1073953:]))
											store32(m.memory[uint32(v14):], uint32(t63))
											v1 = i32(6)
											goto l50
										}
										m.fn12(i32(1), i32(6))
										panic("unreachable")
									}
								}
							l14:
								m.fn11()
								panic("unreachable")
							l50:
								store32(m.memory[int64(uint32(v3))+32:], uint32(v1))
								store32(m.memory[int64(uint32(v3))+28:], uint32(v14))
								store32(m.memory[int64(uint32(v3))+24:], uint32(v1))
								{
									{
										t64 := int32(load32(m.memory[int64(uint32(v3))+36:]))
										v1 = t64
										if v1 == 0 {
											goto l51
										}
										t65 := int32(load32(m.memory[uint32(v17+i32(-4)):]))
										v14 = t65
										v12 = v14 & i32(-8)
										t66 := v12
										v14 = v14 & i32(3)
										p67 := i32(8)
										if v14 != 0 {
											p67 = i32(4)
										}
										if uint32(t66) < uint32(p67+v1) {
											m.fn3(i32(1274224), i32(46), i32(1274272))
											panic("unreachable")
										}
										if v14 == 0 {
											goto l53
										}
										if uint32(v12) > uint32(v1+i32(39)) {
											m.fn3(i32(1274288), i32(46), i32(1274336))
											panic("unreachable")
										}
									l53:
										m.fn1(v17)
									}
								l51:
									m.fn770(v3+i32(12), v5, v3+i32(24))
									t68 := int32(load32(m.memory[int64(uint32(v3))+12:]))
									if t68 == i32(-1) {
										goto l2
									}
									if v10 != 0 {
										t69 := m.fn7(v10)
										v1 = t69
										if v1 != 0 {
											if v10 == 0 {
												goto l56
											}
											memory_copy(m.memory, uint32(v1), uint32(v9), uint32(v10))
											goto l56
										}
										m.fn12(i32(1), v10)
										panic("unreachable")
									}
									v1 = i32(1)
									goto l56
								}
							l56:
								t70 := int64(load64(m.memory[int64(uint32(v7))+16:]))
								t71 := int64(load64(m.memory[int64(uint32(v7))+24:]))
								t72 := m.fn64(t70, t71, v1, v10)
								v11 = t72
								{
									t73 := int32(load32(m.memory[int64(uint32(v7))+8:]))
									if t73 != 0 {
										goto l58
									}
									_ = m.fn74(v7, v8)
								}
							l58:
								t75 := int32(load32(m.memory[int64(uint32(v7))+4:]))
								v19 = t75
								v17 = v19 & int32(v11)
								v25 = int64(uint64(v11) >> 25)
								v13 = v25 & i64(127) * i64(72340172838076673)
								t76 := int32(load32(m.memory[uint32(v7):]))
								v14 = t76
								v18 = i32(0)
								v20 = i32(0)
							l75:
								{
									t77 := int64(load64(m.memory[uint32(v14+v17):]))
									v16 = t77
									v11 = v16 ^ v13
									v11 = (v11 ^ i64(-1)) & (v11 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
									if v11 == 0 {
										goto l59
									}
								l62:
									{
										t78 := v10
										v12 = v14 + (i32(0)-(int32(uint32(int64(bits.TrailingZeros64(uint64(v11))))>>3)+v17)&v19)*i32(28)
										t79 := int32(load32(m.memory[uint32(v12+i32(-20)):]))
										if t78 != t79 {
											goto l60
										}
										t80 := int32(load32(m.memory[uint32(v12+i32(-24)):]))
										t81 := m.fn973(v1, t80, v10)
										if t81 == 0 {
											v14 = v12 + i32(-16)
											t91 := int32(load32(m.memory[int64(uint32(v3))+20:]))
											store32(m.memory[int64(uint32(v14))+8:], uint32(t91))
											m.memory[uint32(v12+i32(-4))] = byte(i32(1))
											t92 := int32(load32(m.memory[uint32(v12+i32(-12)):]))
											v17 = t92
											t93 := int32(load32(m.memory[uint32(v14):]))
											v12 = t93
											t94 := int64(load64(m.memory[int64(uint32(v3))+12:]))
											store64(m.memory[uint32(v14):], uint64(t94))
											{
												if v10 == 0 {
													goto l68
												}
												t95 := int32(load32(m.memory[uint32(v1+i32(-4)):]))
												v14 = t95
												v15 = v14 & i32(-8)
												t96 := v15
												v14 = v14 & i32(3)
												p97 := i32(8)
												if v14 != 0 {
													p97 = i32(4)
												}
												if uint32(t96) < uint32(p97+v10) {
													m.fn3(i32(1274224), i32(46), i32(1274272))
													panic("unreachable")
												}
												if v14 == 0 {
													goto l70
												}
												if uint32(v15) > uint32(v10+i32(39)) {
													m.fn3(i32(1274288), i32(46), i32(1274336))
													panic("unreachable")
												}
											l70:
												m.fn1(v1)
											}
										l68:
											if uint32(v12+i32(-1)) > uint32(i32(-3)) {
												goto l2
											}
											t98 := int32(load32(m.memory[uint32(v17+i32(-4)):]))
											v1 = t98
											v14 = v1 & i32(-8)
											t99 := v14
											v1 = v1 & i32(3)
											p100 := i32(8)
											if v1 != 0 {
												p100 = i32(4)
											}
											if uint32(t99) < uint32(p100+v12) {
												m.fn3(i32(1274224), i32(46), i32(1274272))
												panic("unreachable")
											}
											if v1 == 0 {
												goto l73
											}
											if uint32(v14) > uint32(v12+i32(39)) {
												m.fn3(i32(1274288), i32(46), i32(1274336))
												panic("unreachable")
											}
										l73:
											m.fn1(v17)
											goto l2
										}
									}
								l60:
									v11 = (v11 + i64(-1)) & v11
									if !(v11 == 0) {
										goto l62
									}
								}
							l59:
								v11 = v16 & i64(-0x7f7f7f7f7f7f7f80)
								if v18 == i32(1) {
									goto l63
								}
								if v11 == 0 {
									goto l64
								}
								v15 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v11))))>>3) + v17) & v19
							l63:
								if v11&(v16<<1) != i64(0) {
									{
										t82 := int32(int8(m.memory[uint32(v14+v15)]))
										v12 = t82
										if v12 < i32(0) {
											goto l67
										}
										t83 := int64(load64(m.memory[uint32(v14):]))
										t84 := v14
										v15 = int32(uint32(int64(bits.TrailingZeros64(uint64(t83&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
										t85 := int32(m.memory[uint32(t84+v15)])
										v12 = t85
									}
								l67:
									t86 := v14 + v15
									v17 = int32(v25) & i32(127)
									m.memory[uint32(t86)] = byte(v17)
									m.memory[uint32(v14+(v15+i32(-8))&v19+i32(8))] = byte(v17)
									t87 := int32(load32(m.memory[int64(uint32(v7))+8:]))
									store32(m.memory[int64(uint32(v7))+8:], uint32(t87-v12&i32(1)))
									t88 := int32(load32(m.memory[int64(uint32(v7))+12:]))
									store32(m.memory[int64(uint32(v7))+12:], uint32(t88+i32(1)))
									v14 = v14 + (i32(0)-v15)*i32(28)
									store32(m.memory[uint32(v14+i32(-28)):], uint32(v10))
									store32(m.memory[uint32(v14+i32(-24)):], uint32(v1))
									v1 = v14 + i32(-16)
									t89 := int32(load32(m.memory[int64(uint32(v3))+20:]))
									store32(m.memory[int64(uint32(v1))+8:], uint32(t89))
									t90 := int64(load64(m.memory[int64(uint32(v3))+12:]))
									store64(m.memory[uint32(v1):], uint64(t90))
									store32(m.memory[uint32(v14+i32(-20)):], uint32(v10))
									m.memory[uint32(v14+i32(-4))] = byte(i32(1))
									goto l2
								}
								v18 = i32(1)
								goto l66
							l64:
								v18 = i32(0)
							l66:
								v20 = v20 + i32(8)
								v17 = (v20 + v17) & v19
								goto l75
							}
						}
					l5:
						v11 = (v11 + i64(-1)) & v11
						if !(v11 == 0) {
							goto l7
						}
					}
				l4:
					if !(v16&(v16<<1)&i64(-0x7f7f7f7f7f7f7f80) == 0) {
						goto l2
					}
					t22 := v1
					v15 = v15 + i32(8)
					v1 = (t22 + v15) & v12
					goto l8
				}
			}
		}
	l2:
		v0 = v0 + i32(28)
		if v0 != v4 {
			goto l76
		}
	}
l0:
	m.g0 = v3 + i32(48)
}
func (m *Module) fn774(v0, v1, v2 int32) {
	var v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14 int32
	var v15, v16, v17, v18, v19, v20, v21, v22 int64
	var v23, v24, v25, v26, v27 int32
	t0 := m.g0
	v3 = t0 - i32(208)
	m.g0 = v3
	{
		{
			{
				t1 := int32(load32(m.memory[uint32(v1):]))
				v4 = t1
				switch v4 >> 31 & (v4 + i32(-0x7fffffff)) {
				case 3:
					{
						{
							t291 := int32(m.memory[int64(uint32(v1))+20])
							if t291 != 0 {
								v23 = i32(1)
								t293 := int32(load32(m.memory[int64(uint32(v1))+12:]))
								v13 = t293
								if v13 != i32(1) {
									goto l184
								}
								v24 = i32(12)
								{
									t294 := int32(load32(m.memory[int64(uint32(v1))+8:]))
									v25 = t294
									t295 := int32(load32(m.memory[int64(uint32(v25))+8:]))
									v7 = t295
									if v7 == i32(1) {
										t296 := int32(load32(m.memory[int64(uint32(v25))+4:]))
										v4 = t296
										t297 := int32(load32(m.memory[uint32(v4):]))
										if t297 != i32(-1) {
											t298 := int32(load32(m.memory[int64(uint32(v4))+4:]))
											v6 = t298
											t299 := int32(load32(m.memory[int64(uint32(v4))+8:]))
											v4 = t299
											store32(m.memory[int64(uint32(v3))+160:], uint32(v2))
											store32(m.memory[int64(uint32(v3))+156:], uint32(v6+v4<<5))
											store32(m.memory[int64(uint32(v3))+152:], uint32(v6))
											m.fn786(v3+i32(64), v3+i32(152))
											t300 := int32(load32(m.memory[int64(uint32(v3))+68:]))
											t301 := v3 + i32(152)
											v7 = t300
											t302 := int32(load32(m.memory[int64(uint32(v3))+72:]))
											t303 := v7
											v6 = t302
											m.fn202(t301, t303, v6, i32(1076056), i32(2))
											t304 := int32(load32(m.memory[int64(uint32(v3))+160:]))
											v11 = t304
											t305 := int32(load32(m.memory[int64(uint32(v3))+156:]))
											v5 = t305
											t306 := int32(load32(m.memory[int64(uint32(v3))+152:]))
											v12 = t306
											if v6 == 0 {
												goto l188
											}
											v4 = v7
										l193:
											{
												t307 := int32(load32(m.memory[uint32(v4):]))
												v8 = t307
												if v8 == 0 {
													goto l189
												}
												t308 := int32(load32(m.memory[uint32(v4+i32(4)):]))
												v10 = t308
												t309 := int32(load32(m.memory[uint32(v10+i32(-4)):]))
												v9 = t309
												v2 = v9 & i32(-8)
												t310 := v2
												v9 = v9 & i32(3)
												p311 := i32(8)
												if v9 != 0 {
													p311 = i32(4)
												}
												if uint32(t310) < uint32(p311+v8) {
													m.fn3(i32(1274224), i32(46), i32(1274272))
													panic("unreachable")
												}
												if v9 == 0 {
													goto l191
												}
												if uint32(v2) > uint32(v8+i32(39)) {
													m.fn3(i32(1274288), i32(46), i32(1274336))
													panic("unreachable")
												}
											l191:
												m.fn1(v10)
											}
										l189:
											v4 = v4 + i32(12)
											v6 = v6 + i32(-1)
											if v6 != 0 {
												goto l193
											}
										l188:
											{
												t312 := int32(load32(m.memory[int64(uint32(v3))+64:]))
												v4 = t312
												if v4 == 0 {
													goto l194
												}
												m.fn17(v7, v4*i32(12), i32(4))
											}
										l194:
											if v11 != 0 {
												store32(m.memory[int64(uint32(v0))+8:], uint32(v11))
												store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
												store32(m.memory[uint32(v0):], uint32(v12))
												goto l45
											}
											store32(m.memory[uint32(v0):], uint32(i32(-1)))
											if v12 == 0 {
												goto l45
											}
											m.fn17(v5, v12, i32(1))
											goto l45
										}
										v23 = i32(1)
										v13 = i32(1)
										v7 = i32(1)
										goto l186
									}
									v13 = i32(1)
									goto l186
								}
							}
							t292 := int32(load32(m.memory[int64(uint32(v1))+12:]))
							v13 = t292
							goto l184
						}
					l184:
						if v13 != 0 {
							goto l196
						}
						store32(m.memory[uint32(v0):], uint32(i32(-1)))
						goto l45
					l196:
						v24 = i32(12)
						v23 = i32(1)
						t313 := int32(load32(m.memory[int64(uint32(v1))+8:]))
						v25 = t313
						t314 := int32(load32(m.memory[uint32(v25+i32(8)):]))
						v7 = t314
						if v13 != i32(1) {
							goto l197
						}
						v13 = i32(1)
						goto l186
					l197:
						v24 = v13 * i32(12)
						t315 := int32(uint32(v24+i32(-12)) / uint32(i32(12)))
						v9 = t315
						v6 = v9 & i32(3)
						v8 = i32(0)
						if uint32(v9+i32(-1)) < uint32(i32(3)) {
							goto l198
						}
						v4 = v25 + i32(56)
						v11 = v9 & i32(0x1ffffffc)
						v8 = i32(0)
					l199:
						{
							t316 := int32(load32(m.memory[uint32(v4+i32(-36)):]))
							t317 := v7
							v9 = t316
							p318 := v9
							if uint32(v7) > uint32(v9) {
								p318 = t317
							}
							v9 = p318
							t319 := int32(load32(m.memory[uint32(v4+i32(-24)):]))
							t320 := v9
							v10 = t319
							p321 := v10
							if uint32(v9) > uint32(v10) {
								p321 = t320
							}
							v9 = p321
							t322 := int32(load32(m.memory[uint32(v4+i32(-12)):]))
							t323 := v9
							v10 = t322
							p324 := v10
							if uint32(v9) > uint32(v10) {
								p324 = t323
							}
							v9 = p324
							t325 := int32(load32(m.memory[uint32(v4):]))
							t326 := v9
							v10 = t325
							p327 := v10
							if uint32(v9) > uint32(v10) {
								p327 = t326
							}
							v7 = p327
							v4 = v4 + i32(48)
							t328 := v11
							v8 = v8 + i32(4)
							if t328 != v8 {
								goto l199
							}
						}
						if v6 == 0 {
							goto l200
						}
					l198:
						v4 = v8*i32(12) + v25 + i32(20)
					l201:
						{
							t329 := int32(load32(m.memory[uint32(v4):]))
							t330 := v7
							v8 = t329
							p331 := v8
							if uint32(v7) > uint32(v8) {
								p331 = t330
							}
							v7 = p331
							v4 = v4 + i32(12)
							v6 = v6 + i32(-1)
							if v6 != 0 {
								goto l201
							}
						}
					l200:
						v23 = i32(0)
					}
				l186:
					{
						{
							{
								{
									{
										t332 := m.fn7(v24)
										v14 = t332
										if v14 == 0 {
											m.fn12(i32(4), v24)
											panic("unreachable")
										}
										v12 = i32(0)
									l220:
										{
											v11 = i32(4)
											{
												t333 := v25
												v5 = v12 * i32(12)
												v4 = t333 + v5
												t334 := int32(load32(m.memory[uint32(v4+i32(8)):]))
												v10 = t334
												if v10 == 0 {
													goto l203
												}
												t335 := int32(load32(m.memory[uint32(v4+i32(4)):]))
												v6 = t335
												v9 = v10 << 4
												t336 := m.fn7(v9)
												v11 = t336
												v4 = v11
												v8 = v10
												if v11 == 0 {
													m.fn12(i32(4), v9)
													panic("unreachable")
												}
											l207:
												{
													{
														t337 := int32(load32(m.memory[uint32(v6):]))
														if t337 != i32(-1) {
															goto l205
														}
														store32(m.memory[int64(uint32(v3))+160:], uint32(i32(0)))
														store64(m.memory[int64(uint32(v3))+152:], uint64(i64(0x100000000)))
														v9 = i32(1)
														goto l206
													}
												l205:
													m.fn787(v3+i32(152), v6, v2)
													v9 = i32(0)
												l206:
													t338 := int64(load64(m.memory[int64(uint32(v3))+152:]))
													store64(m.memory[uint32(v4):], uint64(t338))
													m.memory[int64(uint32(v3))+164] = byte(v9)
													t339 := int64(load64(m.memory[int64(uint32(v3))+160:]))
													store64(m.memory[int64(uint32(v4))+8:], uint64(t339))
													v6 = v6 + i32(20)
													v4 = v4 + i32(16)
													v8 = v8 + i32(-1)
													if v8 != 0 {
														goto l207
													}
												}
											}
										l203:
											store32(m.memory[int64(uint32(v3))+160:], uint32(v10))
											store32(m.memory[int64(uint32(v3))+156:], uint32(v11))
											store32(m.memory[int64(uint32(v3))+152:], uint32(v10))
											{
												if uint32(v7) > uint32(v10) {
													goto l208
												}
												store32(m.memory[int64(uint32(v3))+160:], uint32(v7))
												if v10 == v7 {
													goto l209
												}
												v6 = v10 - v7
												v4 = v11 + v7<<4
											l214:
												{
													t340 := int32(load32(m.memory[uint32(v4):]))
													v8 = t340
													if v8 == 0 {
														goto l210
													}
													t341 := int32(load32(m.memory[uint32(v4+i32(4)):]))
													v10 = t341
													t342 := int32(load32(m.memory[uint32(v10+i32(-4)):]))
													v9 = t342
													v11 = v9 & i32(-8)
													t343 := v11
													v9 = v9 & i32(3)
													p344 := i32(8)
													if v9 != 0 {
														p344 = i32(4)
													}
													if uint32(t343) < uint32(p344+v8) {
														m.fn3(i32(1274224), i32(46), i32(1274272))
														panic("unreachable")
													}
													if v9 == 0 {
														goto l212
													}
													if uint32(v11) > uint32(v8+i32(39)) {
														m.fn3(i32(1274288), i32(46), i32(1274336))
														panic("unreachable")
													}
												l212:
													m.fn1(v10)
												}
											l210:
												v4 = v4 + i32(16)
												v6 = v6 + i32(-1)
												if v6 != 0 {
													goto l214
												}
												goto l209
											l208:
												t345 := v3 + i32(152)
												t346 := v10
												v8 = v7 - v10
												m.fn196(t345, t346, v8, i32(4), i32(16))
												t347 := int32(load32(m.memory[int64(uint32(v3))+156:]))
												v26 = t347
												t348 := int32(load32(m.memory[int64(uint32(v3))+160:]))
												v6 = t348
												v4 = v8 & i32(3)
												if v4 != 0 {
													goto l215
												}
												v9 = v6
												goto l216
											l215:
												v9 = v6 + v4
												v11 = v4 << 4
												v8 = v7 - v10 - v4
												v27 = v26 + v6<<4
												v4 = i32(0)
											l217:
												{
													v6 = v27 + v4
													store64(m.memory[uint32(v6):], uint64(i64(0x100000000)))
													m.memory[uint32(v6+i32(12))] = byte(i32(0))
													store32(m.memory[uint32(v6+i32(8)):], uint32(i32(0)))
													t349 := v11
													v4 = v4 + i32(16)
													if t349 != v4 {
														goto l217
													}
												}
											l216:
												if uint32(v10-v7) > uint32(i32(-4)) {
													goto l218
												}
												v4 = v26 + v9<<4
											l219:
												store64(m.memory[uint32(v4):], uint64(i64(0x100000000)))
												m.memory[uint32(v4+i32(60))] = byte(i32(0))
												store32(m.memory[uint32(v4+i32(56)):], uint32(i32(0)))
												store64(m.memory[uint32(v4+i32(48)):], uint64(i64(0x100000000)))
												m.memory[uint32(v4+i32(44))] = byte(i32(0))
												store32(m.memory[uint32(v4+i32(40)):], uint32(i32(0)))
												store64(m.memory[uint32(v4+i32(32)):], uint64(i64(0x100000000)))
												m.memory[uint32(v4+i32(28))] = byte(i32(0))
												store32(m.memory[uint32(v4+i32(24)):], uint32(i32(0)))
												store64(m.memory[uint32(v4+i32(16)):], uint64(i64(0x100000000)))
												m.memory[uint32(v4+i32(12))] = byte(i32(0))
												store32(m.memory[uint32(v4+i32(8)):], uint32(i32(0)))
												v4 = v4 + i32(64)
												v9 = v9 + i32(4)
												v8 = v8 + i32(-4)
												if v8 != 0 {
													goto l219
												}
											l218:
												store32(m.memory[int64(uint32(v3))+160:], uint32(v9))
											}
										l209:
											v4 = v14 + v5
											t350 := int32(load32(m.memory[int64(uint32(v3))+160:]))
											store32(m.memory[int64(uint32(v4))+8:], uint32(t350))
											t351 := int64(load64(m.memory[int64(uint32(v3))+152:]))
											store64(m.memory[uint32(v4):], uint64(t351))
											v12 = v12 + i32(1)
											if v12 != v13 {
												goto l220
											}
										}
										if v23 != 0 {
											v25 = i32(12)
											v12 = v14 + i32(12)
											v13 = i32(1)
											goto l234
										}
									l233:
										{
											t352 := v14
											v11 = v13
											v4 = t352 + v11*i32(12)
											t353 := int32(load32(m.memory[uint32(v4+i32(-4)):]))
											v9 = t353
											v6 = v9 << 4
											t354 := int32(load32(m.memory[uint32(v4+i32(-8)):]))
											v7 = t354
											v4 = v7
										l224:
											if v6 == 0 {
												{
													t357 := v14
													v13 = v11 + i32(-1)
													t358 := int32(load32(m.memory[uint32(t357+v13*i32(12)):]))
													v12 = t358
													if v12 == i32(-1) {
														goto l226
													}
													if v9 == 0 {
														goto l227
													}
													v4 = v7
												l232:
													{
														t359 := int32(load32(m.memory[uint32(v4):]))
														v6 = t359
														if v6 == 0 {
															goto l228
														}
														t360 := int32(load32(m.memory[uint32(v4+i32(4)):]))
														v10 = t360
														t361 := int32(load32(m.memory[uint32(v10+i32(-4)):]))
														v8 = t361
														v2 = v8 & i32(-8)
														t362 := v2
														v8 = v8 & i32(3)
														p363 := i32(8)
														if v8 != 0 {
															p363 = i32(4)
														}
														if uint32(t362) < uint32(p363+v6) {
															m.fn3(i32(1274224), i32(46), i32(1274272))
															panic("unreachable")
														}
														if v8 == 0 {
															goto l230
														}
														if uint32(v2) > uint32(v6+i32(39)) {
															m.fn3(i32(1274288), i32(46), i32(1274336))
															panic("unreachable")
														}
													l230:
														m.fn1(v10)
													}
												l228:
													v4 = v4 + i32(16)
													v9 = v9 + i32(-1)
													if v9 != 0 {
														goto l232
													}
												l227:
													if v12 == 0 {
														goto l226
													}
													m.fn17(v7, v12<<4, i32(4))
												}
											l226:
												if uint32(v11) > uint32(i32(2)) {
													goto l233
												}
												goto l225
											}
											{
												t355 := int32(load32(m.memory[uint32(v4+i32(8)):]))
												if t355 != 0 {
													goto l223
												}
												v6 = v6 + i32(-16)
												v8 = v4 + i32(12)
												v4 = v4 + i32(16)
												t356 := int32(m.memory[uint32(v8)])
												if t356&i32(1) == 0 {
													goto l224
												}
											}
										l223:
											v13 = v11
											goto l225
										}
									}
								l225:
									{
										if v13 == 0 {
											store32(m.memory[uint32(v0):], uint32(i32(-1)))
											goto l236
										}
										t364 := v14
										v25 = v13 * i32(12)
										v12 = t364 + v25
										goto l234
									}
								l234:
									t365 := int32(load32(m.memory[uint32(v14+i32(4)):]))
									t366 := int32(load32(m.memory[uint32(v14+i32(8)):]))
									v8 = t366
									v6 = v8 << 4
									v4 = t365 + v6
									v5 = v14 + i32(12)
									v6 = i32(0) - v6
								l239:
									{
										if v6 != 0 {
											goto l237
										}
										v11 = i32(0)
										goto l238
									l237:
										v11 = v8
										t367 := int32(load32(m.memory[uint32(v4+i32(-8)):]))
										if t367 != 0 {
											goto l238
										}
										v6 = v6 + i32(16)
										v8 = v8 + i32(-1)
										v9 = v4 + i32(-4)
										v4 = v4 + i32(-16)
										t368 := int32(m.memory[uint32(v9)])
										if t368&i32(1) == 0 {
											goto l239
										}
									}
								l238:
									{
										if v13 == i32(1) {
											goto l240
										}
										t369 := int32(uint32(v25+i32(-12)) / uint32(i32(12)))
										v7 = t369
										v2 = i32(0)
									l244:
										{
											v4 = v5 + v2*i32(12)
											t370 := int32(load32(m.memory[uint32(v4+i32(4)):]))
											t371 := int32(load32(m.memory[uint32(v4+i32(8)):]))
											v8 = t371
											v6 = v8 << 4
											v4 = t370 + v6
											v6 = i32(0) - v6
											{
											l243:
												v9 = v8
												if v6 == 0 {
													goto l241
												}
												{
													t372 := int32(load32(m.memory[uint32(v4+i32(-8)):]))
													if t372 != 0 {
														goto l242
													}
													v6 = v6 + i32(16)
													v8 = v9 + i32(-1)
													v10 = v4 + i32(-4)
													v4 = v4 + i32(-16)
													t373 := int32(m.memory[uint32(v10)])
													if t373&i32(1) == 0 {
														goto l243
													}
												}
											l242:
												p374 := v9
												if uint32(v11) > uint32(v9) {
													p374 = v11
												}
												v11 = p374
											}
										l241:
											v2 = v2 + i32(1)
											if v2 != v7 {
												goto l244
											}
										}
									}
								l240:
									{
										if v11 == 0 {
											goto l245
										}
										v27 = v11 << 4
										v7 = v14
									l252:
										{
											t375 := int32(load32(m.memory[int64(uint32(v7))+8:]))
											v4 = t375
											if uint32(v4) < uint32(v11) {
												goto l246
											}
											store32(m.memory[int64(uint32(v7))+8:], uint32(v11))
											if v4 == v11 {
												goto l246
											}
											v6 = v4 - v11
											t376 := int32(load32(m.memory[int64(uint32(v7))+4:]))
											v4 = t376 + v27
										l251:
											{
												t377 := int32(load32(m.memory[uint32(v4):]))
												v8 = t377
												if v8 == 0 {
													goto l247
												}
												t378 := int32(load32(m.memory[uint32(v4+i32(4)):]))
												v10 = t378
												t379 := int32(load32(m.memory[uint32(v10+i32(-4)):]))
												v9 = t379
												v2 = v9 & i32(-8)
												t380 := v2
												v9 = v9 & i32(3)
												p381 := i32(8)
												if v9 != 0 {
													p381 = i32(4)
												}
												if uint32(t380) < uint32(p381+v8) {
													m.fn3(i32(1274224), i32(46), i32(1274272))
													panic("unreachable")
												}
												if v9 == 0 {
													goto l249
												}
												if uint32(v2) > uint32(v8+i32(39)) {
													m.fn3(i32(1274288), i32(46), i32(1274336))
													panic("unreachable")
												}
											l249:
												m.fn1(v10)
											}
										l247:
											v4 = v4 + i32(16)
											v6 = v6 + i32(-1)
											if v6 != 0 {
												goto l251
											}
										}
									l246:
										v7 = v7 + i32(12)
										if v7 != v12 {
											goto l252
										}
										store32(m.memory[int64(uint32(v3))+200:], uint32(i32(0)))
										store64(m.memory[int64(uint32(v3))+192:], uint64(i64(0x100000000)))
										t382 := int32(load32(m.memory[int64(uint32(v1))+16:]))
										if t382 != 0 {
											goto l253
										}
										store32(m.memory[int64(uint32(v3))+160:], uint32(i32(0)))
										store64(m.memory[int64(uint32(v3))+152:], uint64(i64(0x100000000)))
										m.fn647(v3+i32(64), v3+i32(152), v11)
										t383 := int32(load32(m.memory[int64(uint32(v3))+72:]))
										v4 = t383
										t384 := int32(load32(m.memory[int64(uint32(v3))+68:]))
										v6 = t384
										goto l254
									}
								l245:
									store32(m.memory[uint32(v0):], uint32(i32(-1)))
									v11 = i32(0)
								l265:
									{
										v7 = v14 + v11*i32(12)
										t385 := int32(load32(m.memory[int64(uint32(v7))+4:]))
										v12 = t385
										{
											t386 := int32(load32(m.memory[int64(uint32(v7))+8:]))
											v6 = t386
											if v6 == 0 {
												goto l255
											}
											v4 = v12
										l260:
											{
												t387 := int32(load32(m.memory[uint32(v4):]))
												v8 = t387
												if v8 == 0 {
													goto l256
												}
												t388 := int32(load32(m.memory[uint32(v4+i32(4)):]))
												v10 = t388
												t389 := int32(load32(m.memory[uint32(v10+i32(-4)):]))
												v9 = t389
												v2 = v9 & i32(-8)
												t390 := v2
												v9 = v9 & i32(3)
												p391 := i32(8)
												if v9 != 0 {
													p391 = i32(4)
												}
												if uint32(t390) < uint32(p391+v8) {
													m.fn3(i32(1274224), i32(46), i32(1274272))
													panic("unreachable")
												}
												if v9 == 0 {
													goto l258
												}
												if uint32(v2) > uint32(v8+i32(39)) {
													m.fn3(i32(1274288), i32(46), i32(1274336))
													panic("unreachable")
												}
											l258:
												m.fn1(v10)
											}
										l256:
											v4 = v4 + i32(16)
											v6 = v6 + i32(-1)
											if v6 != 0 {
												goto l260
											}
										}
									l255:
										{
											t392 := int32(load32(m.memory[uint32(v7):]))
											v4 = t392
											if v4 == 0 {
												goto l261
											}
											t393 := int32(load32(m.memory[uint32(v12+i32(-4)):]))
											v6 = t393
											v8 = v6 & i32(-8)
											t394 := v8
											v6 = v6 & i32(3)
											p395 := i32(8)
											if v6 != 0 {
												p395 = i32(4)
											}
											v4 = v4 << 4
											if uint32(t394) < uint32(p395|v4) {
												m.fn3(i32(1274224), i32(46), i32(1274272))
												panic("unreachable")
											}
											if v6 == 0 {
												goto l263
											}
											if uint32(v8) > uint32(v4+i32(39)) {
												m.fn3(i32(1274288), i32(46), i32(1274336))
												panic("unreachable")
											}
										l263:
											m.fn1(v12)
										}
									l261:
										v11 = v11 + i32(1)
										if v11 != v13 {
											goto l265
										}
									}
								}
							l236:
								t396 := int32(load32(m.memory[uint32(v14+i32(-4)):]))
								v4 = t396
								v6 = v4 & i32(-8)
								t397 := v6
								v4 = v4 & i32(3)
								p398 := i32(8)
								if v4 != 0 {
									p398 = i32(4)
								}
								if uint32(t397) < uint32(p398+v24) {
									m.fn3(i32(1274224), i32(46), i32(1274272))
									panic("unreachable")
								}
								if v4 == 0 {
									goto l267
								}
								if uint32(v6) > uint32(v24+i32(39)) {
									m.fn3(i32(1274288), i32(46), i32(1274336))
									panic("unreachable")
								}
							l267:
								m.fn1(v14)
								goto l45
							}
						l253:
							t399 := int32(load32(m.memory[uint32(v14+i32(8)):]))
							v6 = t399
							t400 := int32(load32(m.memory[uint32(v14+i32(4)):]))
							v9 = t400
							t401 := int32(load32(m.memory[uint32(v14):]))
							v10 = t401
							v4 = v25 + i32(-12)
							if v4 == 0 {
								goto l269
							}
							memory_copy(m.memory, uint32(v14), uint32(v5), uint32(v4))
						l269:
							v13 = v13 + i32(-1)
							if v10 == i32(-1) {
								m.fn42(v13)
								panic("unreachable")
							}
							v7 = v10 << 4
							t402 := int32(uint32(v7) / uint32(i32(12)))
							v2 = t402
							v4 = v9
							if v6 == 0 {
								goto l271
							}
							v5 = v6 << 4
							v12 = v5 + i32(-16)
							if v12&i32(112) != i32(112) {
								goto l272
							}
							v4 = v9
							v6 = v9
							goto l273
						l272:
							v8 = i32(0) - (int32(uint32(v12)>>4)+i32(1))&i32(7)
							v4 = v9
							v6 = v9
						l274:
							{
								t403 := int64(load64(m.memory[uint32(v6):]))
								v15 = t403
								t404 := int32(load32(m.memory[int64(uint32(v6))+8:]))
								store32(m.memory[int64(uint32(v4))+8:], uint32(t404))
								store64(m.memory[uint32(v4):], uint64(v15))
								v4 = v4 + i32(12)
								v6 = v6 + i32(16)
								v8 = v8 + i32(1)
								if v8 != 0 {
									goto l274
								}
							}
						l273:
							if uint32(v12) < uint32(i32(112)) {
								goto l271
							}
							v8 = v9 + v5
						l275:
							{
								t405 := int64(load64(m.memory[uint32(v6):]))
								v15 = t405
								t406 := int32(load32(m.memory[int64(uint32(v6))+8:]))
								store32(m.memory[int64(uint32(v4))+8:], uint32(t406))
								store64(m.memory[uint32(v4):], uint64(v15))
								t407 := int64(load64(m.memory[uint32(v6+i32(16)):]))
								v15 = t407
								t408 := int32(load32(m.memory[uint32(v6+i32(24)):]))
								store32(m.memory[uint32(v4+i32(20)):], uint32(t408))
								store64(m.memory[uint32(v4+i32(12)):], uint64(v15))
								t409 := int64(load64(m.memory[uint32(v6+i32(32)):]))
								v15 = t409
								t410 := int32(load32(m.memory[uint32(v6+i32(40)):]))
								store32(m.memory[uint32(v4+i32(32)):], uint32(t410))
								store64(m.memory[uint32(v4+i32(24)):], uint64(v15))
								t411 := int64(load64(m.memory[uint32(v6+i32(48)):]))
								v15 = t411
								t412 := int32(load32(m.memory[uint32(v6+i32(56)):]))
								store32(m.memory[uint32(v4+i32(44)):], uint32(t412))
								store64(m.memory[uint32(v4+i32(36)):], uint64(v15))
								t413 := int64(load64(m.memory[uint32(v6+i32(64)):]))
								v15 = t413
								t414 := int32(load32(m.memory[uint32(v6+i32(72)):]))
								store32(m.memory[uint32(v4+i32(56)):], uint32(t414))
								store64(m.memory[uint32(v4+i32(48)):], uint64(v15))
								t415 := int64(load64(m.memory[uint32(v6+i32(80)):]))
								v15 = t415
								t416 := int32(load32(m.memory[uint32(v6+i32(88)):]))
								store32(m.memory[uint32(v4+i32(68)):], uint32(t416))
								store64(m.memory[uint32(v4+i32(60)):], uint64(v15))
								t417 := int64(load64(m.memory[uint32(v6+i32(96)):]))
								v15 = t417
								t418 := int32(load32(m.memory[uint32(v6+i32(104)):]))
								store32(m.memory[uint32(v4+i32(80)):], uint32(t418))
								store64(m.memory[uint32(v4+i32(72)):], uint64(v15))
								t419 := int64(load64(m.memory[uint32(v6+i32(112)):]))
								v15 = t419
								t420 := int32(load32(m.memory[uint32(v6+i32(120)):]))
								store32(m.memory[uint32(v4+i32(92)):], uint32(t420))
								store64(m.memory[uint32(v4+i32(84)):], uint64(v15))
								v4 = v4 + i32(96)
								v6 = v6 + i32(128)
								if v6 != v8 {
									goto l275
								}
							}
						l271:
							{
								if v10 != 0 {
									goto l276
								}
								v6 = v9
								goto l277
							l276:
								v6 = v9
								t421 := v7
								v8 = v2 * i32(12)
								if t421 == v8 {
									goto l277
								}
								if v7 != 0 {
									goto l278
								}
								v6 = i32(4)
								goto l277
							l278:
								t422 := m.fn21(v9, v7, i32(4), v8)
								v6 = t422
								if v6 == 0 {
									m.fn23(i32(4), v8)
									panic("unreachable")
								}
							}
						l277:
							store32(m.memory[int64(uint32(v3))+68:], uint32(v6))
							store32(m.memory[int64(uint32(v3))+64:], uint32(v2))
							t423 := int32(uint32(v4-v9) / uint32(i32(12)))
							t424 := v3
							v4 = t423
							store32(m.memory[int64(uint32(t424))+72:], uint32(v4))
						}
					l254:
						v8 = i32(1)
						t425 := m.fn7(i32(1))
						v9 = t425
						if v9 == 0 {
							m.fn12(i32(1), i32(1))
							panic("unreachable")
						}
						m.memory[uint32(v9)] = byte(i32(124))
						store32(m.memory[int64(uint32(v3))+160:], uint32(i32(1)))
						store32(m.memory[int64(uint32(v3))+156:], uint32(v9))
						store32(m.memory[int64(uint32(v3))+152:], uint32(i32(1)))
						{
							if v4 != 0 {
								goto l281
							}
							v10 = i32(0)
							v4 = i32(0)
							v6 = i32(1)
							goto l282
						l281:
							v10 = v4 * i32(12)
							v8 = v6 + i32(8)
							v6 = i32(1)
						l288:
							{
								t426 := int32(load32(m.memory[uint32(v8+i32(-4)):]))
								v2 = t426
								t427 := int32(load32(m.memory[uint32(v8):]))
								v4 = t427
								{
									t428 := int32(load32(m.memory[int64(uint32(v3))+152:]))
									if t428 != v6 {
										goto l283
									}
									m.fn196(v3+i32(152), v6, i32(1), i32(1), i32(1))
									t429 := int32(load32(m.memory[int64(uint32(v3))+156:]))
									v9 = t429
								}
							l283:
								m.memory[uint32(v9+v6)] = byte(i32(32))
								t430 := v3
								v6 = v6 + i32(1)
								store32(m.memory[int64(uint32(t430))+160:], uint32(v6))
								{
									{
										t431 := int32(load32(m.memory[int64(uint32(v3))+152:]))
										t432 := v4
										v9 = t431
										if uint32(t432) <= uint32(v9-v6) {
											goto l284
										}
										m.fn196(v3+i32(152), v6, v4, i32(1), i32(1))
										t433 := int32(load32(m.memory[int64(uint32(v3))+152:]))
										v9 = t433
										t434 := int32(load32(m.memory[int64(uint32(v3))+160:]))
										v6 = t434
										goto l285
									}
								l284:
									if v4 == 0 {
										goto l286
									}
								l285:
									if v4 == 0 {
										goto l286
									}
									t435 := int32(load32(m.memory[int64(uint32(v3))+156:]))
									memory_copy(m.memory, uint32(t435+v6), uint32(v2), uint32(v4))
								}
							l286:
								t436 := v3
								v4 = v6 + v4
								store32(m.memory[int64(uint32(t436))+160:], uint32(v4))
								{
									if uint32(v9-v4) > uint32(i32(1)) {
										goto l287
									}
									m.fn196(v3+i32(152), v4, i32(2), i32(1), i32(1))
									t437 := int32(load32(m.memory[int64(uint32(v3))+160:]))
									v4 = t437
								}
							l287:
								t438 := int32(load32(m.memory[int64(uint32(v3))+156:]))
								v9 = t438
								store16(m.memory[uint32(v9+v4):], uint16(i32(31776)))
								t439 := v3
								v6 = v4 + i32(2)
								store32(m.memory[int64(uint32(t439))+160:], uint32(v6))
								v8 = v8 + i32(12)
								v10 = v10 + i32(-12)
								if v10 != 0 {
									goto l288
								}
							}
							t440 := int32(load32(m.memory[int64(uint32(v3))+192:]))
							v10 = t440
							t441 := int32(load32(m.memory[int64(uint32(v3))+200:]))
							v4 = t441
							t442 := int32(load32(m.memory[int64(uint32(v3))+152:]))
							v8 = t442
						}
					l282:
						{
							if uint32(v6) <= uint32(v10-v4) {
								goto l289
							}
							m.fn196(v3+i32(192), v4, v6, i32(1), i32(1))
							t443 := int32(load32(m.memory[int64(uint32(v3))+200:]))
							v4 = t443
						}
					l289:
						t444 := int32(load32(m.memory[int64(uint32(v3))+196:]))
						v10 = t444
						if v6 == 0 {
							goto l290
						}
						memory_copy(m.memory, uint32(v10+v4), uint32(v9), uint32(v6))
					l290:
						t445 := v3
						v4 = v4 + v6
						store32(m.memory[int64(uint32(t445))+200:], uint32(v4))
						{
							if v8 == 0 {
								goto l291
							}
							t446 := int32(load32(m.memory[uint32(v9+i32(-4)):]))
							v6 = t446
							v2 = v6 & i32(-8)
							t447 := v2
							v6 = v6 & i32(3)
							p448 := i32(8)
							if v6 != 0 {
								p448 = i32(4)
							}
							if uint32(t447) < uint32(p448+v8) {
								m.fn3(i32(1274224), i32(46), i32(1274272))
								panic("unreachable")
							}
							if v6 == 0 {
								goto l293
							}
							if uint32(v2) > uint32(v8+i32(39)) {
								m.fn3(i32(1274288), i32(46), i32(1274336))
								panic("unreachable")
							}
						l293:
							m.fn1(v9)
						}
					l291:
						{
							t449 := int32(load32(m.memory[int64(uint32(v3))+192:]))
							if t449 != v4 {
								goto l295
							}
							m.fn196(v3+i32(192), v4, i32(1), i32(1), i32(1))
							t450 := int32(load32(m.memory[int64(uint32(v3))+196:]))
							v10 = t450
						}
					l295:
						m.memory[uint32(v10+v4)] = byte(i32(10))
						store32(m.memory[int64(uint32(v3))+200:], uint32(v4+i32(1)))
						t451 := m.fn7(i32(1))
						v6 = t451
						if v6 == 0 {
							m.fn12(i32(1), i32(1))
							panic("unreachable")
						}
						m.memory[uint32(v6)] = byte(i32(124))
						store32(m.memory[int64(uint32(v3))+160:], uint32(i32(1)))
						store32(m.memory[int64(uint32(v3))+156:], uint32(v6))
						store32(m.memory[int64(uint32(v3))+152:], uint32(i32(1)))
						v4 = i32(1)
					l300:
						{
							{
								t452 := int32(load32(m.memory[int64(uint32(v3))+152:]))
								if t452 != v4 {
									goto l297
								}
								m.fn196(v3+i32(152), v4, i32(1), i32(1), i32(1))
								t453 := int32(load32(m.memory[int64(uint32(v3))+156:]))
								v6 = t453
							}
						l297:
							m.memory[uint32(v6+v4)] = byte(i32(32))
							t454 := v3
							v4 = v4 + i32(1)
							store32(m.memory[int64(uint32(t454))+160:], uint32(v4))
							{
								t455 := int32(load32(m.memory[int64(uint32(v3))+152:]))
								v8 = t455
								if uint32(v8-v4) > uint32(i32(2)) {
									goto l298
								}
								m.fn196(v3+i32(152), v4, i32(3), i32(1), i32(1))
								t456 := int32(load32(m.memory[int64(uint32(v3))+152:]))
								v8 = t456
								t457 := int32(load32(m.memory[int64(uint32(v3))+160:]))
								v4 = t457
							}
						l298:
							t458 := int32(load32(m.memory[int64(uint32(v3))+156:]))
							v6 = t458
							v9 = v6 + v4
							store16(m.memory[uint32(v9):], uint16(i32(11565)))
							m.memory[int64(uint32(v9))+2] = byte(i32(45))
							t459 := v3
							v4 = v4 + i32(3)
							store32(m.memory[int64(uint32(t459))+160:], uint32(v4))
							{
								if uint32(v8-v4) > uint32(i32(1)) {
									goto l299
								}
								m.fn196(v3+i32(152), v4, i32(2), i32(1), i32(1))
								t460 := int32(load32(m.memory[int64(uint32(v3))+156:]))
								v6 = t460
								t461 := int32(load32(m.memory[int64(uint32(v3))+160:]))
								v4 = t461
							}
						l299:
							store16(m.memory[uint32(v6+v4):], uint16(i32(31776)))
							t462 := v3
							v4 = v4 + i32(2)
							store32(m.memory[int64(uint32(t462))+160:], uint32(v4))
							v11 = v11 + i32(-1)
							if v11 != 0 {
								goto l300
							}
						}
						t463 := int32(load32(m.memory[int64(uint32(v3))+152:]))
						v8 = t463
						{
							t464 := int32(load32(m.memory[int64(uint32(v3))+192:]))
							t465 := int32(load32(m.memory[int64(uint32(v3))+200:]))
							t466 := v4
							v9 = t465
							if uint32(t466) <= uint32(t464-v9) {
								goto l301
							}
							m.fn196(v3+i32(192), v9, v4, i32(1), i32(1))
							t467 := int32(load32(m.memory[int64(uint32(v3))+200:]))
							v9 = t467
						}
					l301:
						t468 := int32(load32(m.memory[int64(uint32(v3))+196:]))
						v10 = t468
						if v4 == 0 {
							goto l302
						}
						memory_copy(m.memory, uint32(v10+v9), uint32(v6), uint32(v4))
					l302:
						t469 := v3
						v4 = v9 + v4
						store32(m.memory[int64(uint32(t469))+200:], uint32(v4))
						{
							if v8 == 0 {
								goto l303
							}
							t470 := int32(load32(m.memory[uint32(v6+i32(-4)):]))
							v9 = t470
							v2 = v9 & i32(-8)
							t471 := v2
							v9 = v9 & i32(3)
							p472 := i32(8)
							if v9 != 0 {
								p472 = i32(4)
							}
							if uint32(t471) < uint32(p472+v8) {
								m.fn3(i32(1274224), i32(46), i32(1274272))
								panic("unreachable")
							}
							if v9 == 0 {
								goto l305
							}
							if uint32(v2) > uint32(v8+i32(39)) {
								m.fn3(i32(1274288), i32(46), i32(1274336))
								panic("unreachable")
							}
						l305:
							m.fn1(v6)
						}
					l303:
						if v13 == 0 {
							goto l307
						}
						v7 = v14 + v13*i32(12)
						v11 = v14
					l324:
						{
							{
								t473 := int32(load32(m.memory[int64(uint32(v3))+192:]))
								if t473 != v4 {
									goto l308
								}
								m.fn196(v3+i32(192), v4, i32(1), i32(1), i32(1))
								t474 := int32(load32(m.memory[int64(uint32(v3))+196:]))
								v10 = t474
							}
						l308:
							m.memory[uint32(v10+v4)] = byte(i32(10))
							t475 := v3
							v4 = v4 + i32(1)
							store32(m.memory[int64(uint32(t475))+200:], uint32(v4))
							t476 := int32(load32(m.memory[int64(uint32(v11))+8:]))
							v6 = t476
							t477 := int32(load32(m.memory[int64(uint32(v11))+4:]))
							v2 = t477
							t478 := m.fn7(i32(1))
							v9 = t478
							if v9 == 0 {
								m.fn12(i32(1), i32(1))
								panic("unreachable")
							}
							m.memory[uint32(v9)] = byte(i32(124))
							v8 = i32(1)
							store32(m.memory[int64(uint32(v3))+160:], uint32(i32(1)))
							store32(m.memory[int64(uint32(v3))+156:], uint32(v9))
							store32(m.memory[int64(uint32(v3))+152:], uint32(i32(1)))
							{
								if v6 != 0 {
									goto l310
								}
								v6 = i32(1)
								goto l311
							l310:
								v10 = v6 << 4
								v8 = v2 + i32(8)
								v6 = i32(1)
							l317:
								{
									t479 := int32(load32(m.memory[uint32(v8+i32(-4)):]))
									v2 = t479
									t480 := int32(load32(m.memory[uint32(v8):]))
									v4 = t480
									{
										t481 := int32(load32(m.memory[int64(uint32(v3))+152:]))
										if t481 != v6 {
											goto l312
										}
										m.fn196(v3+i32(152), v6, i32(1), i32(1), i32(1))
										t482 := int32(load32(m.memory[int64(uint32(v3))+156:]))
										v9 = t482
									}
								l312:
									m.memory[uint32(v9+v6)] = byte(i32(32))
									t483 := v3
									v6 = v6 + i32(1)
									store32(m.memory[int64(uint32(t483))+160:], uint32(v6))
									{
										{
											t484 := int32(load32(m.memory[int64(uint32(v3))+152:]))
											t485 := v4
											v9 = t484
											if uint32(t485) <= uint32(v9-v6) {
												goto l313
											}
											m.fn196(v3+i32(152), v6, v4, i32(1), i32(1))
											t486 := int32(load32(m.memory[int64(uint32(v3))+152:]))
											v9 = t486
											t487 := int32(load32(m.memory[int64(uint32(v3))+160:]))
											v6 = t487
											goto l314
										}
									l313:
										if v4 == 0 {
											goto l315
										}
									l314:
										if v4 == 0 {
											goto l315
										}
										t488 := int32(load32(m.memory[int64(uint32(v3))+156:]))
										memory_copy(m.memory, uint32(t488+v6), uint32(v2), uint32(v4))
									}
								l315:
									t489 := v3
									v4 = v6 + v4
									store32(m.memory[int64(uint32(t489))+160:], uint32(v4))
									{
										if uint32(v9-v4) > uint32(i32(1)) {
											goto l316
										}
										m.fn196(v3+i32(152), v4, i32(2), i32(1), i32(1))
										t490 := int32(load32(m.memory[int64(uint32(v3))+160:]))
										v4 = t490
									}
								l316:
									t491 := int32(load32(m.memory[int64(uint32(v3))+156:]))
									v9 = t491
									store16(m.memory[uint32(v9+v4):], uint16(i32(31776)))
									t492 := v3
									v6 = v4 + i32(2)
									store32(m.memory[int64(uint32(t492))+160:], uint32(v6))
									v8 = v8 + i32(16)
									v10 = v10 + i32(-16)
									if v10 != 0 {
										goto l317
									}
								}
								t493 := int32(load32(m.memory[int64(uint32(v3))+200:]))
								v4 = t493
								t494 := int32(load32(m.memory[int64(uint32(v3))+152:]))
								v8 = t494
							}
						l311:
							{
								t495 := int32(load32(m.memory[int64(uint32(v3))+192:]))
								if uint32(v6) <= uint32(t495-v4) {
									goto l318
								}
								m.fn196(v3+i32(192), v4, v6, i32(1), i32(1))
								t496 := int32(load32(m.memory[int64(uint32(v3))+200:]))
								v4 = t496
							}
						l318:
							t497 := int32(load32(m.memory[int64(uint32(v3))+196:]))
							v10 = t497
							if v6 == 0 {
								goto l319
							}
							memory_copy(m.memory, uint32(v10+v4), uint32(v9), uint32(v6))
						l319:
							t498 := v3
							v4 = v4 + v6
							store32(m.memory[int64(uint32(t498))+200:], uint32(v4))
							{
								if v8 == 0 {
									goto l320
								}
								t499 := int32(load32(m.memory[uint32(v9+i32(-4)):]))
								v6 = t499
								v2 = v6 & i32(-8)
								t500 := v2
								v6 = v6 & i32(3)
								p501 := i32(8)
								if v6 != 0 {
									p501 = i32(4)
								}
								if uint32(t500) < uint32(p501+v8) {
									m.fn3(i32(1274224), i32(46), i32(1274272))
									panic("unreachable")
								}
								if v6 == 0 {
									goto l322
								}
								if uint32(v2) > uint32(v8+i32(39)) {
									m.fn3(i32(1274288), i32(46), i32(1274336))
									panic("unreachable")
								}
							l322:
								m.fn1(v9)
							}
						l320:
							v11 = v11 + i32(12)
							if v11 != v7 {
								goto l324
							}
						}
					l307:
						t502 := int32(load32(m.memory[int64(uint32(v3))+200:]))
						store32(m.memory[int64(uint32(v0))+8:], uint32(t502))
						t503 := int64(load64(m.memory[int64(uint32(v3))+192:]))
						store64(m.memory[uint32(v0):], uint64(t503))
						t504 := int32(load32(m.memory[int64(uint32(v3))+68:]))
						v11 = t504
						{
							t505 := int32(load32(m.memory[int64(uint32(v3))+72:]))
							v6 = t505
							if v6 == 0 {
								goto l325
							}
							v4 = v11
						l330:
							{
								t506 := int32(load32(m.memory[uint32(v4):]))
								v8 = t506
								if v8 == 0 {
									goto l326
								}
								t507 := int32(load32(m.memory[uint32(v4+i32(4)):]))
								v10 = t507
								t508 := int32(load32(m.memory[uint32(v10+i32(-4)):]))
								v9 = t508
								v2 = v9 & i32(-8)
								t509 := v2
								v9 = v9 & i32(3)
								p510 := i32(8)
								if v9 != 0 {
									p510 = i32(4)
								}
								if uint32(t509) < uint32(p510+v8) {
									m.fn3(i32(1274224), i32(46), i32(1274272))
									panic("unreachable")
								}
								if v9 == 0 {
									goto l328
								}
								if uint32(v2) > uint32(v8+i32(39)) {
									m.fn3(i32(1274288), i32(46), i32(1274336))
									panic("unreachable")
								}
							l328:
								m.fn1(v10)
							}
						l326:
							v4 = v4 + i32(12)
							v6 = v6 + i32(-1)
							if v6 != 0 {
								goto l330
							}
						}
					l325:
						{
							t511 := int32(load32(m.memory[int64(uint32(v3))+64:]))
							v4 = t511
							if v4 == 0 {
								goto l331
							}
							t512 := int32(load32(m.memory[uint32(v11+i32(-4)):]))
							v6 = t512
							v8 = v6 & i32(-8)
							t513 := v8
							v6 = v6 & i32(3)
							p514 := i32(8)
							if v6 != 0 {
								p514 = i32(4)
							}
							v4 = v4 * i32(12)
							if uint32(t513) < uint32(p514+v4) {
								m.fn3(i32(1274224), i32(46), i32(1274272))
								panic("unreachable")
							}
							if v6 == 0 {
								goto l333
							}
							if uint32(v8) > uint32(v4+i32(39)) {
								m.fn3(i32(1274288), i32(46), i32(1274336))
								panic("unreachable")
							}
						l333:
							m.fn1(v11)
						}
					l331:
						if v13 == 0 {
							goto l335
						}
						v11 = i32(0)
					l346:
						{
							v7 = v14 + v11*i32(12)
							t515 := int32(load32(m.memory[int64(uint32(v7))+4:]))
							v12 = t515
							{
								t516 := int32(load32(m.memory[int64(uint32(v7))+8:]))
								v6 = t516
								if v6 == 0 {
									goto l336
								}
								v4 = v12
							l341:
								{
									t517 := int32(load32(m.memory[uint32(v4):]))
									v8 = t517
									if v8 == 0 {
										goto l337
									}
									t518 := int32(load32(m.memory[uint32(v4+i32(4)):]))
									v10 = t518
									t519 := int32(load32(m.memory[uint32(v10+i32(-4)):]))
									v9 = t519
									v2 = v9 & i32(-8)
									t520 := v2
									v9 = v9 & i32(3)
									p521 := i32(8)
									if v9 != 0 {
										p521 = i32(4)
									}
									if uint32(t520) < uint32(p521+v8) {
										m.fn3(i32(1274224), i32(46), i32(1274272))
										panic("unreachable")
									}
									if v9 == 0 {
										goto l339
									}
									if uint32(v2) > uint32(v8+i32(39)) {
										m.fn3(i32(1274288), i32(46), i32(1274336))
										panic("unreachable")
									}
								l339:
									m.fn1(v10)
								}
							l337:
								v4 = v4 + i32(16)
								v6 = v6 + i32(-1)
								if v6 != 0 {
									goto l341
								}
							}
						l336:
							{
								t522 := int32(load32(m.memory[uint32(v7):]))
								v4 = t522
								if v4 == 0 {
									goto l342
								}
								t523 := int32(load32(m.memory[uint32(v12+i32(-4)):]))
								v6 = t523
								v8 = v6 & i32(-8)
								t524 := v8
								v6 = v6 & i32(3)
								p525 := i32(8)
								if v6 != 0 {
									p525 = i32(4)
								}
								v4 = v4 << 4
								if uint32(t524) < uint32(p525|v4) {
									m.fn3(i32(1274224), i32(46), i32(1274272))
									panic("unreachable")
								}
								if v6 == 0 {
									goto l344
								}
								if uint32(v8) > uint32(v4+i32(39)) {
									m.fn3(i32(1274288), i32(46), i32(1274336))
									panic("unreachable")
								}
							l344:
								m.fn1(v12)
							}
						l342:
							v11 = v11 + i32(1)
							if v11 != v13 {
								goto l346
							}
						}
					l335:
						t526 := int32(load32(m.memory[uint32(v14+i32(-4)):]))
						v4 = t526
						v6 = v4 & i32(-8)
						t527 := v6
						v4 = v4 & i32(3)
						p528 := i32(8)
						if v4 != 0 {
							p528 = i32(4)
						}
						if uint32(t527) < uint32(p528+v24) {
							m.fn3(i32(1274224), i32(46), i32(1274272))
							panic("unreachable")
						}
						if v4 == 0 {
							goto l348
						}
						if uint32(v6) > uint32(v24+i32(39)) {
							m.fn3(i32(1274288), i32(46), i32(1274336))
							panic("unreachable")
						}
					l348:
						m.fn1(v14)
						goto l45
					}
				case 4:
					t233 := int32(load32(m.memory[int64(uint32(v1))+8:]))
					v4 = t233
					t234 := int32(load32(m.memory[int64(uint32(v1))+12:]))
					t235 := v4
					v6 = t234 << 5
					v8 = t235 + v6
					{
						{
						l147:
							{
								if v6 != 0 {
									goto l145
								}
								v11 = i32(4)
								v6 = i32(0)
								v1 = i32(0)
								goto l146
							l145:
								m.fn774(v3+i32(64), v4, v2)
								v6 = v6 + i32(-32)
								v4 = v4 + i32(32)
								t236 := int32(load32(m.memory[int64(uint32(v3))+64:]))
								if t236 == i32(-1) {
									goto l147
								}
							}
							t237 := m.fn7(i32(48))
							v9 = t237
							if v9 == 0 {
								m.fn12(i32(4), i32(48))
								panic("unreachable")
							}
							t238 := int32(load32(m.memory[int64(uint32(v3))+72:]))
							store32(m.memory[int64(uint32(v9))+8:], uint32(t238))
							t239 := int64(load64(m.memory[int64(uint32(v3))+64:]))
							store64(m.memory[uint32(v9):], uint64(t239))
							store32(m.memory[int64(uint32(v3))+200:], uint32(i32(1)))
							store32(m.memory[int64(uint32(v3))+196:], uint32(v9))
							store32(m.memory[int64(uint32(v3))+192:], uint32(i32(4)))
							v6 = i32(1)
						l150:
							{
								if v4 == v8 {
									goto l149
								}
								m.fn774(v3+i32(152), v4, v2)
								v4 = v4 + i32(32)
								t240 := int32(load32(m.memory[int64(uint32(v3))+152:]))
								if t240 == i32(-1) {
									goto l150
								}
								{
									t241 := int32(load32(m.memory[int64(uint32(v3))+192:]))
									if v6 != t241 {
										goto l151
									}
									m.fn196(v3+i32(192), v6, i32(1), i32(4), i32(12))
									t242 := int32(load32(m.memory[int64(uint32(v3))+196:]))
									v9 = t242
								}
							l151:
								v10 = v9 + v6*i32(12)
								t243 := int32(load32(m.memory[int64(uint32(v3))+160:]))
								store32(m.memory[int64(uint32(v10))+8:], uint32(t243))
								t244 := int64(load64(m.memory[int64(uint32(v3))+152:]))
								store64(m.memory[uint32(v10):], uint64(t244))
								t245 := v3
								v6 = v6 + i32(1)
								store32(m.memory[int64(uint32(t245))+200:], uint32(v6))
								goto l150
							}
						l149:
							t246 := int32(load32(m.memory[int64(uint32(v3))+196:]))
							v11 = t246
							t247 := int32(load32(m.memory[int64(uint32(v3))+192:]))
							v1 = t247
						}
					l146:
						m.fn202(v3+i32(152), v11, v6, i32(1076056), i32(2))
						t248 := int32(load32(m.memory[int64(uint32(v3))+160:]))
						v7 = t248
						t249 := int32(load32(m.memory[int64(uint32(v3))+156:]))
						v5 = t249
						t250 := int32(load32(m.memory[int64(uint32(v3))+152:]))
						v12 = t250
						if v6 == 0 {
							goto l152
						}
						v4 = v11
					l157:
						{
							t251 := int32(load32(m.memory[uint32(v4):]))
							v8 = t251
							if v8 == 0 {
								goto l153
							}
							t252 := int32(load32(m.memory[uint32(v4+i32(4)):]))
							v10 = t252
							t253 := int32(load32(m.memory[uint32(v10+i32(-4)):]))
							v9 = t253
							v2 = v9 & i32(-8)
							t254 := v2
							v9 = v9 & i32(3)
							p255 := i32(8)
							if v9 != 0 {
								p255 = i32(4)
							}
							if uint32(t254) < uint32(p255+v8) {
								m.fn3(i32(1274224), i32(46), i32(1274272))
								panic("unreachable")
							}
							if v9 == 0 {
								goto l155
							}
							if uint32(v2) > uint32(v8+i32(39)) {
								m.fn3(i32(1274288), i32(46), i32(1274336))
								panic("unreachable")
							}
						l155:
							m.fn1(v10)
						}
					l153:
						v4 = v4 + i32(12)
						v6 = v6 + i32(-1)
						if v6 != 0 {
							goto l157
						}
					l152:
						{
							if v1 == 0 {
								goto l158
							}
							t256 := int32(load32(m.memory[uint32(v11+i32(-4)):]))
							v4 = t256
							v6 = v4 & i32(-8)
							t257 := v6
							v4 = v4 & i32(3)
							p258 := i32(8)
							if v4 != 0 {
								p258 = i32(4)
							}
							v8 = v1 * i32(12)
							if uint32(t257) < uint32(p258+v8) {
								m.fn3(i32(1274224), i32(46), i32(1274272))
								panic("unreachable")
							}
							if v4 == 0 {
								goto l160
							}
							if uint32(v6) > uint32(v8+i32(39)) {
								m.fn3(i32(1274288), i32(46), i32(1274336))
								panic("unreachable")
							}
						l160:
							m.fn1(v11)
						}
					l158:
						{
							if v7 != 0 {
								store16(m.memory[int64(uint32(v3))+100:], uint16(i32(0)))
								store32(m.memory[int64(uint32(v3))+96:], uint32(v7))
								store32(m.memory[int64(uint32(v3))+92:], uint32(i32(0)))
								m.memory[int64(uint32(v3))+88] = byte(i32(1))
								store32(m.memory[int64(uint32(v3))+84:], uint32(i32(10)))
								store32(m.memory[int64(uint32(v3))+80:], uint32(v7))
								store32(m.memory[int64(uint32(v3))+76:], uint32(i32(0)))
								store32(m.memory[int64(uint32(v3))+72:], uint32(v7))
								store32(m.memory[int64(uint32(v3))+68:], uint32(v5))
								store32(m.memory[int64(uint32(v3))+64:], uint32(i32(10)))
								m.fn785(v3+i32(140), v3+i32(64))
								{
									{
										t262 := int32(load32(m.memory[int64(uint32(v3))+140:]))
										if t262 != i32(-1) {
											goto l166
										}
										store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
										store64(m.memory[uint32(v0):], uint64(i64(0x100000000)))
										goto l167
									}
								l166:
									t263 := m.fn7(i32(48))
									v9 = t263
									if v9 == 0 {
										m.fn12(i32(4), i32(48))
										panic("unreachable")
									}
									t264 := int32(load32(m.memory[int64(uint32(v3))+148:]))
									store32(m.memory[int64(uint32(v9))+8:], uint32(t264))
									t265 := int64(load64(m.memory[int64(uint32(v3))+140:]))
									store64(m.memory[uint32(v9):], uint64(t265))
									store32(m.memory[int64(uint32(v3))+136:], uint32(i32(1)))
									store32(m.memory[int64(uint32(v3))+132:], uint32(v9))
									store32(m.memory[int64(uint32(v3))+128:], uint32(i32(4)))
									t266 := int64(load64(m.memory[int64(uint32(v3))+96:]))
									store64(m.memory[int64(uint32(v3))+184:], uint64(t266))
									t267 := int64(load64(m.memory[int64(uint32(v3))+88:]))
									store64(m.memory[int64(uint32(v3))+176:], uint64(t267))
									t268 := int64(load64(m.memory[int64(uint32(v3))+80:]))
									store64(m.memory[int64(uint32(v3))+168:], uint64(t268))
									t269 := int64(load64(m.memory[int64(uint32(v3))+72:]))
									store64(m.memory[int64(uint32(v3))+160:], uint64(t269))
									t270 := int64(load64(m.memory[int64(uint32(v3))+64:]))
									store64(m.memory[int64(uint32(v3))+152:], uint64(t270))
									v6 = i32(12)
									v4 = i32(1)
								l171:
									{
										m.fn785(v3+i32(192), v3+i32(152))
										t271 := int32(load32(m.memory[int64(uint32(v3))+192:]))
										if t271 == i32(-1) {
											goto l169
										}
										{
											t272 := int32(load32(m.memory[int64(uint32(v3))+128:]))
											if v4 != t272 {
												goto l170
											}
											m.fn196(v3+i32(128), v4, i32(1), i32(4), i32(12))
											t273 := int32(load32(m.memory[int64(uint32(v3))+132:]))
											v9 = t273
										}
									l170:
										v8 = v9 + v6
										t274 := int32(load32(m.memory[int64(uint32(v3))+200:]))
										store32(m.memory[int64(uint32(v8))+8:], uint32(t274))
										t275 := int64(load64(m.memory[int64(uint32(v3))+192:]))
										store64(m.memory[uint32(v8):], uint64(t275))
										t276 := v3
										v4 = v4 + i32(1)
										store32(m.memory[int64(uint32(t276))+136:], uint32(v4))
										v6 = v6 + i32(12)
										goto l171
									}
								l169:
									t277 := int32(load32(m.memory[int64(uint32(v3))+128:]))
									v7 = t277
									t278 := int32(load32(m.memory[int64(uint32(v3))+132:]))
									t279 := v0
									v11 = t278
									m.fn202(t279, v11, v4, i32(1099470), i32(1))
									v6 = v11
								l176:
									{
										t280 := int32(load32(m.memory[uint32(v6):]))
										v8 = t280
										if v8 == 0 {
											goto l172
										}
										t281 := int32(load32(m.memory[uint32(v6+i32(4)):]))
										v10 = t281
										t282 := int32(load32(m.memory[uint32(v10+i32(-4)):]))
										v9 = t282
										v2 = v9 & i32(-8)
										t283 := v2
										v9 = v9 & i32(3)
										p284 := i32(8)
										if v9 != 0 {
											p284 = i32(4)
										}
										if uint32(t283) < uint32(p284+v8) {
											m.fn3(i32(1274224), i32(46), i32(1274272))
											panic("unreachable")
										}
										if v9 == 0 {
											goto l174
										}
										if uint32(v2) > uint32(v8+i32(39)) {
											m.fn3(i32(1274288), i32(46), i32(1274336))
											panic("unreachable")
										}
									l174:
										m.fn1(v10)
									}
								l172:
									v6 = v6 + i32(12)
									v4 = v4 + i32(-1)
									if v4 != 0 {
										goto l176
									}
									if v7 == 0 {
										goto l167
									}
									t285 := int32(load32(m.memory[uint32(v11+i32(-4)):]))
									v4 = t285
									v6 = v4 & i32(-8)
									t286 := v6
									v4 = v4 & i32(3)
									p287 := i32(8)
									if v4 != 0 {
										p287 = i32(4)
									}
									v8 = v7 * i32(12)
									if uint32(t286) < uint32(p287+v8) {
										m.fn3(i32(1274224), i32(46), i32(1274272))
										panic("unreachable")
									}
									if v4 == 0 {
										goto l178
									}
									if uint32(v6) > uint32(v8+i32(39)) {
										m.fn3(i32(1274288), i32(46), i32(1274336))
										panic("unreachable")
									}
								l178:
									m.fn1(v11)
								}
							l167:
								if v12 == 0 {
									goto l45
								}
								t288 := int32(load32(m.memory[uint32(v5+i32(-4)):]))
								v4 = t288
								v6 = v4 & i32(-8)
								t289 := v6
								v4 = v4 & i32(3)
								p290 := i32(8)
								if v4 != 0 {
									p290 = i32(4)
								}
								if uint32(t289) < uint32(p290+v12) {
									m.fn3(i32(1274224), i32(46), i32(1274272))
									panic("unreachable")
								}
								if v4 == 0 {
									goto l181
								}
								if uint32(v6) > uint32(v12+i32(39)) {
									m.fn3(i32(1274288), i32(46), i32(1274336))
									panic("unreachable")
								}
							l181:
								m.fn1(v5)
								goto l45
							}
							store32(m.memory[uint32(v0):], uint32(i32(-1)))
							if v12 == 0 {
								goto l45
							}
							t259 := int32(load32(m.memory[uint32(v5+i32(-4)):]))
							v4 = t259
							v6 = v4 & i32(-8)
							t260 := v6
							v4 = v4 & i32(3)
							p261 := i32(8)
							if v4 != 0 {
								p261 = i32(4)
							}
							if uint32(t260) < uint32(p261+v12) {
								m.fn3(i32(1274224), i32(46), i32(1274272))
								panic("unreachable")
							}
							if v4 == 0 {
								goto l164
							}
							if uint32(v6) > uint32(v12+i32(39)) {
								m.fn3(i32(1274288), i32(46), i32(1274336))
								panic("unreachable")
							}
						l164:
							m.fn1(v5)
							goto l45
						}
					}
				case 5:
					t179 := int32(load32(m.memory[int64(uint32(v1))+8:]))
					t180 := v3 + i32(64)
					v9 = t179
					t181 := int32(load32(m.memory[int64(uint32(v1))+12:]))
					t182 := v9
					v6 = t181
					m.fn784(t180, t182, v6, i32(3))
					t183 := int32(load32(m.memory[int64(uint32(v1))+24:]))
					t184 := int32(load32(m.memory[int64(uint32(v1))+16:]))
					t185 := v3
					var p186 int32
					if t184 == i32(-1) {
						p186 = 1
					}
					v4 = p186
					p187 := t183
					if v4 != 0 {
						p187 = i32(0)
					}
					store32(m.memory[int64(uint32(t185))+144:], uint32(p187))
					t188 := int32(load32(m.memory[int64(uint32(v1))+20:]))
					t190 := v3
					p189 := t188
					if v4 != 0 {
						p189 = i32(1)
					}
					store32(m.memory[int64(uint32(t190))+140:], uint32(p189))
				l124:
					v8 = v6
					if v8 != 0 {
						goto l117
					}
					v8 = i32(0)
					goto l118
				l117:
					{
						v10 = v9 + v8
						v6 = v10 + i32(-1)
						t191 := int32(int8(m.memory[uint32(v6)]))
						v4 = t191
						if v4 > i32(-1) {
							goto l119
						}
						{
							v6 = v10 + i32(-2)
							t192 := int32(m.memory[uint32(v6)])
							v2 = t192
							v11 = int32(int8(v2))
							if v11 < i32(-64) {
								goto l120
							}
							v10 = v2 & i32(31)
							goto l121
						}
					l120:
						{
							{
								v6 = v10 + i32(-3)
								t193 := int32(m.memory[uint32(v6)])
								v2 = t193
								v7 = int32(int8(v2))
								if v7 < i32(-64) {
									goto l122
								}
								v10 = v2 & i32(15)
								goto l123
							}
						l122:
							v6 = v10 + i32(-4)
							t194 := int32(m.memory[uint32(v6)])
							v10 = t194&i32(7)<<6 | v7&i32(63)
						}
					l123:
						v10 = v10<<6 | v11&i32(63)
					l121:
						v4 = v10<<6 | v4&i32(63)
					}
				l119:
					v6 = v6 - v9
					if v4 == i32(10) {
						goto l124
					}
				l118:
					store32(m.memory[int64(uint32(v3))+196:], uint32(v8))
					store32(m.memory[int64(uint32(v3))+192:], uint32(v9))
					t195 := v3
					v15 = int64(uint32(i32(1))) << 32
					store64(m.memory[int64(uint32(t195))+168:], uint64(v15|int64(uint32(v3+i32(192)))))
					store64(m.memory[int64(uint32(v3))+160:], uint64(v15|int64(uint32(v3+i32(140)))))
					store64(m.memory[int64(uint32(v3))+152:], uint64(int64(uint32(i32(17)))<<32|int64(uint32(v3+i32(64)))))
					m.fn13(v0, i32(1076058), v3+i32(152))
					t196 := int32(load32(m.memory[int64(uint32(v3))+64:]))
					v4 = t196
					if v4 == 0 {
						goto l45
					}
					{
						t197 := int32(load32(m.memory[int64(uint32(v3))+68:]))
						v8 = t197
						t198 := int32(load32(m.memory[uint32(v8+i32(-4)):]))
						v6 = t198
						v9 = v6 & i32(-8)
						t199 := v9
						v6 = v6 & i32(3)
						p200 := i32(8)
						if v6 != 0 {
							p200 = i32(4)
						}
						if uint32(t199) < uint32(p200+v4) {
							m.fn3(i32(1274224), i32(46), i32(1274272))
							panic("unreachable")
						}
						if v6 == 0 {
							goto l126
						}
						if uint32(v9) > uint32(v4+i32(39)) {
							m.fn3(i32(1274288), i32(46), i32(1274336))
							panic("unreachable")
						}
					l126:
						m.fn1(v8)
						goto l45
					}
				case 6:
					t178 := m.fn7(i32(3))
					v4 = t178
					if v4 == 0 {
						m.fn12(i32(1), i32(3))
						panic("unreachable")
					}
					m.memory[int64(uint32(v4))+2] = byte(i32(45))
					store16(m.memory[uint32(v4):], uint16(i32(11565)))
					store32(m.memory[int64(uint32(v0))+8:], uint32(i32(3)))
					store32(m.memory[int64(uint32(v0))+4:], uint32(v4))
					store32(m.memory[uint32(v0):], uint32(i32(3)))
					goto l45
				default:
					v8 = i32(1)
					t206 := int32(load32(m.memory[int64(uint32(v1))+4:]))
					t207 := int32(load32(m.memory[int64(uint32(v1))+8:]))
					m.fn780(v3+i32(192), t206, t207, i32(1), i32(0), v2)
					t208 := int32(load32(m.memory[int64(uint32(v3))+196:]))
					t209 := v3
					v10 = t208
					t210 := int32(load32(m.memory[int64(uint32(v3))+200:]))
					m.fn143(t209, v10, t210)
					t211 := int32(load32(m.memory[int64(uint32(v3))+4:]))
					t212 := v3
					v4 = t211
					store32(m.memory[int64(uint32(t212))+144:], uint32(v4))
					t213 := int32(load32(m.memory[uint32(v3):]))
					store32(m.memory[int64(uint32(v3))+140:], uint32(t213))
					if v4 == 0 {
						store32(m.memory[uint32(v0):], uint32(i32(-1)))
						t229 := int32(load32(m.memory[int64(uint32(v3))+192:]))
						v4 = t229
						if v4 == 0 {
							goto l45
						}
						t230 := int32(load32(m.memory[uint32(v10+i32(-4)):]))
						v6 = t230
						v8 = v6 & i32(-8)
						t231 := v8
						v6 = v6 & i32(3)
						p232 := i32(8)
						if v6 != 0 {
							p232 = i32(4)
						}
						if uint32(t231) < uint32(p232+v4) {
							m.fn3(i32(1274224), i32(46), i32(1274272))
							panic("unreachable")
						}
						if v6 == 0 {
							goto l142
						}
						if uint32(v8) <= uint32(v4+i32(39)) {
							goto l142
						}
						m.fn3(i32(1274288), i32(46), i32(1274336))
						panic("unreachable")
					}
					{
						{
							t214 := int32(m.memory[int64(uint32(v1))+24])
							v4 = t214
							p215 := i32(6)
							if uint32(v4) < uint32(i32(6)) {
								p215 = v4
							}
							p216 := i32(1)
							if v4 != 0 {
								p216 = p215
							}
							v9 = p216
							if v9 == 0 {
								goto l132
							}
							t217 := m.fn7(v9)
							v8 = t217
							if v8 == 0 {
								m.fn12(i32(1), v9)
								panic("unreachable")
							}
							m.memory[uint32(v8)] = byte(i32(35))
							v4 = i32(1)
							v6 = int32(uint32(v9) >> 1)
							if v6 == 0 {
								goto l134
							}
							v4 = i32(1)
						l136:
							if v4 == 0 {
								goto l135
							}
							memory_copy(m.memory, uint32(v8+v4), uint32(v8), uint32(v4))
						l135:
							v4 = v4 << 1
							v6 = int32(uint32(v6) >> 1)
							if v6 != 0 {
								goto l136
							}
						l134:
							if v9 == v4 {
								goto l132
							}
							v6 = v9 - v4
							if v6 == 0 {
								goto l132
							}
							memory_copy(m.memory, uint32(v8+v4), uint32(v8), uint32(v6))
						}
					l132:
						store32(m.memory[int64(uint32(v3))+72:], uint32(v9))
						store32(m.memory[int64(uint32(v3))+68:], uint32(v8))
						store32(m.memory[int64(uint32(v3))+64:], uint32(v9))
						store64(m.memory[int64(uint32(v3))+160:], uint64(int64(uint32(i32(1)))<<32|int64(uint32(v3+i32(140)))))
						store64(m.memory[int64(uint32(v3))+152:], uint64(int64(uint32(i32(17)))<<32|int64(uint32(v3+i32(64)))))
						m.fn13(v3+i32(32), i32(1052642), v3+i32(152))
						{
							t218 := int32(load32(m.memory[int64(uint32(v3))+64:]))
							v4 = t218
							if v4 == 0 {
								goto l137
							}
							t219 := int32(load32(m.memory[int64(uint32(v3))+68:]))
							v8 = t219
							t220 := int32(load32(m.memory[uint32(v8+i32(-4)):]))
							v6 = t220
							v9 = v6 & i32(-8)
							t221 := v9
							v6 = v6 & i32(3)
							p222 := i32(8)
							if v6 != 0 {
								p222 = i32(4)
							}
							if uint32(t221) < uint32(p222+v4) {
								m.fn3(i32(1274224), i32(46), i32(1274272))
								panic("unreachable")
							}
							if v6 == 0 {
								goto l139
							}
							if uint32(v9) > uint32(v4+i32(39)) {
								m.fn3(i32(1274288), i32(46), i32(1274336))
								panic("unreachable")
							}
						l139:
							m.fn1(v8)
						}
					l137:
						t223 := int32(load32(m.memory[int64(uint32(v3))+40:]))
						store32(m.memory[int64(uint32(v0))+8:], uint32(t223))
						t224 := int64(load64(m.memory[int64(uint32(v3))+32:]))
						store64(m.memory[uint32(v0):], uint64(t224))
						t225 := int32(load32(m.memory[int64(uint32(v3))+192:]))
						v4 = t225
						if v4 == 0 {
							goto l45
						}
						t226 := int32(load32(m.memory[uint32(v10+i32(-4)):]))
						v6 = t226
						v8 = v6 & i32(-8)
						t227 := v8
						v6 = v6 & i32(3)
						p228 := i32(8)
						if v6 != 0 {
							p228 = i32(4)
						}
						if uint32(t227) < uint32(p228+v4) {
							m.fn3(i32(1274224), i32(46), i32(1274272))
							panic("unreachable")
						}
						if v6 == 0 {
							goto l142
						}
						if uint32(v8) > uint32(v4+i32(39)) {
							m.fn3(i32(1274288), i32(46), i32(1274336))
							panic("unreachable")
						}
						goto l142
					}
				l142:
					m.fn1(v10)
					goto l45
				case 1:
					v5 = i32(0)
					t2 := int32(load32(m.memory[int64(uint32(v1))+8:]))
					t3 := int32(load32(m.memory[int64(uint32(v1))+12:]))
					m.fn780(v3+i32(140), t2, t3, i32(0), i32(0), v2)
					t4 := int32(load32(m.memory[int64(uint32(v3))+144:]))
					v6 = t4
					t5 := int32(load32(m.memory[int64(uint32(v3))+148:]))
					v4 = t5
					store16(m.memory[int64(uint32(v3))+100:], uint16(i32(0)))
					store32(m.memory[int64(uint32(v3))+96:], uint32(v4))
					store32(m.memory[int64(uint32(v3))+92:], uint32(i32(0)))
					m.memory[int64(uint32(v3))+88] = byte(i32(1))
					store32(m.memory[int64(uint32(v3))+84:], uint32(i32(10)))
					store32(m.memory[int64(uint32(v3))+80:], uint32(v4))
					store32(m.memory[int64(uint32(v3))+76:], uint32(i32(0)))
					store32(m.memory[int64(uint32(v3))+72:], uint32(v4))
					store32(m.memory[int64(uint32(v3))+68:], uint32(v6))
					store32(m.memory[int64(uint32(v3))+64:], uint32(i32(10)))
					m.fn781(v3+i32(24), v3+i32(64))
					{
						{
							t6 := int32(load32(m.memory[int64(uint32(v3))+24:]))
							v4 = t6
							if v4 != 0 {
								goto l7
							}
							v6 = i32(4)
							v7 = i32(4)
							v8 = i32(0)
							v1 = i32(0)
							goto l8
						}
					l7:
						t7 := int32(load32(m.memory[int64(uint32(v3))+28:]))
						v6 = t7
						t8 := m.fn7(i32(32))
						v2 = t8
						if v2 == 0 {
							m.fn12(i32(4), i32(32))
							panic("unreachable")
						}
						store32(m.memory[uint32(v2):], uint32(v4))
						store32(m.memory[int64(uint32(v2))+4:], uint32(v6))
						store32(m.memory[int64(uint32(v3))+200:], uint32(i32(1)))
						store32(m.memory[int64(uint32(v3))+196:], uint32(v2))
						store32(m.memory[int64(uint32(v3))+192:], uint32(i32(4)))
						t9 := int64(load64(m.memory[int64(uint32(v3))+96:]))
						store64(m.memory[int64(uint32(v3))+184:], uint64(t9))
						t10 := int64(load64(m.memory[int64(uint32(v3))+88:]))
						store64(m.memory[int64(uint32(v3))+176:], uint64(t10))
						t11 := int64(load64(m.memory[int64(uint32(v3))+80:]))
						store64(m.memory[int64(uint32(v3))+168:], uint64(t11))
						t12 := int64(load64(m.memory[int64(uint32(v3))+72:]))
						store64(m.memory[int64(uint32(v3))+160:], uint64(t12))
						t13 := int64(load64(m.memory[int64(uint32(v3))+64:]))
						store64(m.memory[int64(uint32(v3))+152:], uint64(t13))
						v4 = i32(8)
						v8 = i32(1)
					l12:
						{
							m.fn781(v3+i32(16), v3+i32(152))
							t14 := int32(load32(m.memory[int64(uint32(v3))+16:]))
							v6 = t14
							if v6 == 0 {
								goto l10
							}
							t15 := int32(load32(m.memory[int64(uint32(v3))+20:]))
							v9 = t15
							{
								t16 := int32(load32(m.memory[int64(uint32(v3))+192:]))
								if v8 != t16 {
									goto l11
								}
								m.fn196(v3+i32(192), v8, i32(1), i32(4), i32(8))
								t17 := int32(load32(m.memory[int64(uint32(v3))+196:]))
								v2 = t17
							}
						l11:
							v10 = v2 + v4
							store32(m.memory[uint32(v10):], uint32(v6))
							store32(m.memory[uint32(v10+i32(4)):], uint32(v9))
							t18 := v3
							v8 = v8 + i32(1)
							store32(m.memory[int64(uint32(t18))+200:], uint32(v8))
							v4 = v4 + i32(8)
							goto l12
						}
					l10:
						t19 := int32(load32(m.memory[int64(uint32(v3))+196:]))
						v7 = t19
						v9 = v7 + i32(4)
						v6 = v7 + v8<<3
						t20 := int32(load32(m.memory[int64(uint32(v3))+192:]))
						v5 = t20
						v2 = i32(0)
						v1 = i32(1)
					l13:
						{
							t21 := int32(load32(m.memory[uint32(v9):]))
							if t21 != 0 {
								goto l8
							}
							v9 = v9 + i32(8)
							v2 = v2 + i32(1)
							v4 = v4 + i32(-8)
							if v4 != 0 {
								goto l13
							}
						}
						v1 = i32(0)
					}
				l8:
					v4 = v8
					{
					l16:
						{
							v9 = v4
							v11 = i32(1)
							v12 = i32(0)
							if v6 != v7 {
								goto l14
							}
							v10 = i32(0)
							goto l15
						l14:
							v4 = v9 + i32(-1)
							v10 = v6 + i32(-4)
							v6 = v6 + i32(-8)
							t22 := int32(load32(m.memory[uint32(v10):]))
							if t22 == 0 {
								goto l16
							}
						}
						if v1 != 0 {
							goto l17
						}
						v10 = i32(0)
						goto l15
					l17:
						if uint32(v9) < uint32(v2) {
							m.fn120(v2, v9, v8, i32(1076072))
							panic("unreachable")
						}
						m.fn782(v3+i32(152), v7+v2<<3, v9-v2, i32(1099470), i32(1))
						t23 := int32(load32(m.memory[int64(uint32(v3))+156:]))
						v11 = t23
						t24 := int32(load32(m.memory[int64(uint32(v3))+152:]))
						v10 = t24
						t25 := int32(load32(m.memory[int64(uint32(v3))+160:]))
						v13 = t25
						if v13 == 0 {
							goto l15
						}
						v9 = i32(0)
						v14 = v11 + v13
						v8 = v14
					l25:
						{
							v4 = v8 + i32(-1)
							t26 := int32(int8(m.memory[uint32(v4)]))
							v6 = t26
							if v6 > i32(-1) {
								goto l19
							}
							{
								v4 = v8 + i32(-2)
								t27 := int32(m.memory[uint32(v4)])
								v2 = t27
								v12 = int32(int8(v2))
								if v12 < i32(-64) {
									goto l20
								}
								v8 = v2 & i32(31)
								goto l21
							}
						l20:
							{
								{
									v4 = v8 + i32(-3)
									t28 := int32(m.memory[uint32(v4)])
									v2 = t28
									v1 = int32(int8(v2))
									if v1 < i32(-64) {
										goto l22
									}
									v8 = v2 & i32(15)
									goto l23
								}
							l22:
								v4 = v8 + i32(-4)
								t29 := int32(m.memory[uint32(v4)])
								v8 = t29&i32(7)<<6 | v1&i32(63)
							}
						l23:
							v8 = v8<<6 | v12&i32(63)
						l21:
							v6 = v8<<6 | v6&i32(63)
						}
					l19:
						if v6 != i32(92) {
							goto l24
						}
						v9 = v9 + i32(1)
						v8 = v4
						if v11 != v4 {
							goto l25
						}
					l24:
						if v9&i32(1) != 0 {
							goto l26
						}
						v12 = v13
						goto l15
					l26:
						v4 = i32(-1)
						{
							t30 := int32(int8(m.memory[uint32(v14+i32(-1))]))
							if t30 > i32(-1) {
								goto l27
							}
							{
								t31 := int32(m.memory[uint32(v14+i32(-2))])
								v6 = t31
								v8 = int32(int8(v6))
								if v8 <= i32(-65) {
									goto l28
								}
								v6 = v6 & i32(31)
								goto l29
							}
						l28:
							{
								{
									t32 := int32(m.memory[uint32(v14+i32(-3))])
									v6 = t32
									v9 = int32(int8(v6))
									if v9 <= i32(-65) {
										goto l30
									}
									v6 = v6 & i32(15)
									goto l31
								}
							l30:
								t33 := int32(m.memory[uint32(v14+i32(-4))])
								v6 = t33&i32(7)<<6 | v9&i32(63)
							}
						l31:
							v6 = v6<<6 | v8&i32(63)
						l29:
							if uint32(v6) < uint32(i32(2)) {
								goto l27
							}
							v4 = i32(-2)
							if uint32(v6) < uint32(i32(32)) {
								goto l27
							}
							p34 := i32(-4)
							if uint32(v6) < uint32(i32(1024)) {
								p34 = i32(-3)
							}
							v4 = p34
						}
					l27:
						t35 := v3 + i32(8)
						t36 := v11
						v4 = v4 + v13
						m.fn693(t35, t36, v4)
						{
							t37 := int32(load32(m.memory[int64(uint32(v3))+12:]))
							v12 = t37
							if uint32(v12) <= uint32(v4) {
								goto l32
							}
							v12 = v4
							goto l15
						}
					l32:
						if v12 != 0 {
							goto l33
						}
						v12 = i32(0)
						goto l15
					l33:
						if uint32(v12) >= uint32(v4) {
							goto l15
						}
						t38 := int32(int8(m.memory[uint32(v11+v12)]))
						if t38 <= i32(-65) {
							m.fn3(i32(1080817), i32(48), i32(1076088))
							panic("unreachable")
						}
					}
				l15:
					{
						if v5 == 0 {
							goto l35
						}
						t39 := int32(load32(m.memory[uint32(v7+i32(-4)):]))
						v4 = t39
						v6 = v4 & i32(-8)
						t40 := v6
						v4 = v4 & i32(3)
						p41 := i32(8)
						if v4 != 0 {
							p41 = i32(4)
						}
						v8 = v5 << 3
						if uint32(t40) < uint32(p41+v8) {
							m.fn3(i32(1274224), i32(46), i32(1274272))
							panic("unreachable")
						}
						if v4 == 0 {
							goto l37
						}
						if uint32(v6) > uint32(v8+i32(39)) {
							m.fn3(i32(1274288), i32(46), i32(1274336))
							panic("unreachable")
						}
					l37:
						m.fn1(v7)
					}
				l35:
					if v12 != 0 {
						goto l39
					}
					store32(m.memory[uint32(v0):], uint32(i32(-1)))
					if v10 == 0 {
						goto l40
					}
					t42 := int32(load32(m.memory[uint32(v11+i32(-4)):]))
					v4 = t42
					v6 = v4 & i32(-8)
					t43 := v6
					v4 = v4 & i32(3)
					p44 := i32(8)
					if v4 != 0 {
						p44 = i32(4)
					}
					if uint32(t43) < uint32(p44+v10) {
						m.fn3(i32(1274224), i32(46), i32(1274272))
						panic("unreachable")
					}
					if v4 == 0 {
						goto l42
					}
					if uint32(v6) > uint32(v10+i32(39)) {
						m.fn3(i32(1274288), i32(46), i32(1274336))
						panic("unreachable")
					}
				l42:
					m.fn1(v11)
					goto l40
				case 2:
					t45 := int32(load32(m.memory[int64(uint32(v1))+24:]))
					v4 = t45
					if v4 != 0 {
						store32(m.memory[int64(uint32(v3))+52:], uint32(i32(0)))
						store64(m.memory[int64(uint32(v3))+44:], uint64(i64(0x400000000)))
						t46 := int32(load32(m.memory[int64(uint32(v1))+20:]))
						v5 = t46
						v13 = v5 + v4*i32(28)
						v15 = int64(uint32(i32(1))) << 32
						v16 = v15 | int64(uint32(v3+i32(108)))
						v17 = v15 | int64(uint32(v3+i32(56)))
						v15 = int64(uint32(i32(17))) << 32
						v18 = v15 | int64(uint32(v3+i32(128)))
						v19 = v15 | int64(uint32(v3+i32(152)))
						v20 = int64(uint32(i32(10)))<<32 | int64(uint32(v3+i32(192)))
						t47 := int64(load64(m.memory[int64(uint32(v1))+8:]))
						v21 = t47
						t48 := int32(m.memory[int64(uint32(v1))+28])
						v1 = t48
						v15 = i64(0)
						v10 = i32(0)
					l115:
						{
							{
								t49 := int32(load32(m.memory[int64(uint32(v5))+12:]))
								if t49 == i32(-1) {
									goto l46
								}
								t50 := int32(load32(m.memory[int64(uint32(v5))+16:]))
								t51 := int32(load32(m.memory[int64(uint32(v5))+20:]))
								m.fn783(v3+i32(152), t50, t51, i32(0))
								store64(m.memory[int64(uint32(v3))+192:], uint64(v19))
								m.fn13(v3+i32(64), i32(1067924), v3+i32(192))
								{
									t52 := int32(load32(m.memory[int64(uint32(v3))+152:]))
									v4 = t52
									if v4 == 0 {
										goto l47
									}
									t53 := int32(load32(m.memory[int64(uint32(v3))+156:]))
									v8 = t53
									t54 := int32(load32(m.memory[uint32(v8+i32(-4)):]))
									v6 = t54
									v9 = v6 & i32(-8)
									t55 := v9
									v6 = v6 & i32(3)
									p56 := i32(8)
									if v6 != 0 {
										p56 = i32(4)
									}
									if uint32(t55) < uint32(p56+v4) {
										m.fn3(i32(1274224), i32(46), i32(1274272))
										panic("unreachable")
									}
									if v6 == 0 {
										goto l49
									}
									if uint32(v9) > uint32(v4+i32(39)) {
										m.fn3(i32(1274288), i32(46), i32(1274336))
										panic("unreachable")
									}
								l49:
									m.fn1(v8)
								}
							l47:
								t57 := int64(load64(m.memory[int64(uint32(v3))+64:]))
								store64(m.memory[int64(uint32(v3))+128:], uint64(t57))
								t58 := int32(load32(m.memory[int64(uint32(v3))+72:]))
								store32(m.memory[int64(uint32(v3))+136:], uint32(t58))
								goto l51
							}
						l46:
							switch v1 {
							default:
								t59 := v3 + i32(152)
								t60 := v1
								v22 = v21 + v15
								p61 := v22
								if uint64(v22) < uint64(v21) {
									p61 = i64(-1)
								}
								m.fn301(t59, t60, p61)
								store64(m.memory[int64(uint32(v3))+192:], uint64(v19))
								m.fn13(v3+i32(64), i32(1067924), v3+i32(192))
								{
									t62 := int32(load32(m.memory[int64(uint32(v3))+152:]))
									v4 = t62
									if v4 == 0 {
										goto l55
									}
									t63 := int32(load32(m.memory[int64(uint32(v3))+156:]))
									v8 = t63
									t64 := int32(load32(m.memory[uint32(v8+i32(-4)):]))
									v6 = t64
									v9 = v6 & i32(-8)
									t65 := v9
									v6 = v6 & i32(3)
									p66 := i32(8)
									if v6 != 0 {
										p66 = i32(4)
									}
									if uint32(t65) < uint32(p66+v4) {
										m.fn3(i32(1274224), i32(46), i32(1274272))
										panic("unreachable")
									}
									if v6 == 0 {
										goto l57
									}
									if uint32(v9) > uint32(v4+i32(39)) {
										m.fn3(i32(1274288), i32(46), i32(1274336))
										panic("unreachable")
									}
								l57:
									m.fn1(v8)
								}
							l55:
								t67 := int64(load64(m.memory[int64(uint32(v3))+64:]))
								store64(m.memory[int64(uint32(v3))+128:], uint64(t67))
								t68 := int32(load32(m.memory[int64(uint32(v3))+72:]))
								store32(m.memory[int64(uint32(v3))+136:], uint32(t68))
								goto l51
							case 1:
								t69 := v3
								v22 = v21 + v15
								p70 := v22
								if uint64(v22) < uint64(v21) {
									p70 = i64(-1)
								}
								store64(m.memory[int64(uint32(t69))+192:], uint64(p70))
								store64(m.memory[int64(uint32(v3))+64:], uint64(v20))
								m.fn13(v3+i32(152), i32(1067919), v3+i32(64))
								t71 := int64(load64(m.memory[int64(uint32(v3))+152:]))
								store64(m.memory[int64(uint32(v3))+128:], uint64(t71))
								t72 := int32(load32(m.memory[int64(uint32(v3))+160:]))
								store32(m.memory[int64(uint32(v3))+136:], uint32(t72))
								goto l51
							case 0:
								t73 := m.fn7(i32(2))
								v4 = t73
								if v4 == 0 {
									m.fn12(i32(1), i32(2))
									panic("unreachable")
								}
								store16(m.memory[uint32(v4):], uint16(i32(8237)))
								store32(m.memory[int64(uint32(v3))+136:], uint32(i32(2)))
								store32(m.memory[int64(uint32(v3))+132:], uint32(v4))
								store32(m.memory[int64(uint32(v3))+128:], uint32(i32(2)))
							}
						l51:
							v6 = i32(4)
							v4 = i32(1076048)
							{
								t74 := int32(m.memory[int64(uint32(v5))+24])
								switch t74 {
								case 0:
									goto l60
								default:
									goto l61
								case 2:
									v6 = i32(0)
									v4 = i32(1)
									goto l60
								}
							}
						l61:
							v4 = i32(1076052)
						l60:
							store32(m.memory[int64(uint32(v3))+60:], uint32(v6))
							store32(m.memory[int64(uint32(v3))+56:], uint32(v4))
							t75 := int32(load32(m.memory[int64(uint32(v5))+4:]))
							t76 := int32(load32(m.memory[int64(uint32(v5))+8:]))
							t77 := v3 + i32(140)
							v12 = t76
							m.fn777(t77, t75, v12, v2)
							t78 := int32(load32(m.memory[int64(uint32(v3))+132:]))
							v11 = t78
							{
								t79 := int32(load32(m.memory[int64(uint32(v3))+136:]))
								v6 = t79
								if uint32(v6) < uint32(i32(16)) {
									goto l63
								}
								t80 := m.fn574(v11, v6)
								v4 = t80
								goto l64
							}
						l63:
							if v6 != 0 {
								goto l65
							}
							v4 = i32(0)
							goto l64
						l65:
							v8 = v6 & i32(3)
							v9 = i32(0)
							v4 = i32(0)
							if uint32(v6) < uint32(i32(4)) {
								goto l66
							}
							v7 = v6 & i32(12)
							v9 = i32(0)
							v4 = i32(0)
						l67:
							{
								t81 := v4
								v6 = v11 + v9
								t82 := int32(int8(m.memory[uint32(v6)]))
								var p83 int32
								if t82 > i32(-65) {
									p83 = 1
								}
								t84 := int32(int8(m.memory[uint32(v6+i32(1))]))
								t85 := t81 + p83
								var p86 int32
								if t84 > i32(-65) {
									p86 = 1
								}
								t87 := int32(int8(m.memory[uint32(v6+i32(2))]))
								t88 := t85 + p86
								var p89 int32
								if t87 > i32(-65) {
									p89 = 1
								}
								t90 := int32(int8(m.memory[uint32(v6+i32(3))]))
								t91 := t88 + p89
								var p92 int32
								if t90 > i32(-65) {
									p92 = 1
								}
								v4 = t91 + p92
								t93 := v7
								v9 = v9 + i32(4)
								if t93 != v9 {
									goto l67
								}
							}
							if v8 == 0 {
								goto l64
							}
						l66:
							v6 = v11 + v9
						l68:
							{
								t94 := int32(int8(m.memory[uint32(v6)]))
								t95 := v4
								var p96 int32
								if t94 > i32(-65) {
									p96 = 1
								}
								v4 = t95 + p96
								v6 = v6 + i32(1)
								v8 = v8 + i32(-1)
								if v8 != 0 {
									goto l68
								}
							}
						l64:
							{
								{
									if v4 != 0 {
										goto l69
									}
									v7 = i32(1)
									goto l70
								l69:
									if v4 <= i32(-1) {
										m.fn11()
										panic("unreachable")
									}
									t97 := m.fn7(v4)
									v7 = t97
									if v7 == 0 {
										m.fn12(i32(1), v4)
										panic("unreachable")
									}
									m.memory[uint32(v7)] = byte(i32(32))
									v6 = i32(1)
									v8 = int32(uint32(v4) >> 1)
									if v8 == 0 {
										goto l73
									}
									v6 = i32(1)
								l75:
									if v6 == 0 {
										goto l74
									}
									memory_copy(m.memory, uint32(v7+v6), uint32(v7), uint32(v6))
								l74:
									v6 = v6 << 1
									v8 = int32(uint32(v8) >> 1)
									if v8 != 0 {
										goto l75
									}
								l73:
									if v4 == v6 {
										goto l70
									}
									v8 = v4 - v6
									if v8 == 0 {
										goto l70
									}
									memory_copy(m.memory, uint32(v7+v6), uint32(v7), uint32(v8))
								}
							l70:
								;
								var p98 int32
								if uint32(v12) > uint32(i32(1)) {
									p98 = 1
								}
								v9 = p98
								store16(m.memory[int64(uint32(v3))+100:], uint16(i32(0)))
								t99 := int32(load32(m.memory[int64(uint32(v3))+148:]))
								t100 := v3
								v6 = t99
								store32(m.memory[int64(uint32(t100))+96:], uint32(v6))
								store32(m.memory[int64(uint32(v3))+92:], uint32(i32(0)))
								store32(m.memory[int64(uint32(v3))+84:], uint32(i32(10)))
								store32(m.memory[int64(uint32(v3))+80:], uint32(v6))
								store32(m.memory[int64(uint32(v3))+76:], uint32(i32(0)))
								store32(m.memory[int64(uint32(v3))+72:], uint32(v6))
								t101 := int32(load32(m.memory[int64(uint32(v3))+144:]))
								t102 := v3
								v6 = t101
								store32(m.memory[int64(uint32(t102))+68:], uint32(v6))
								store32(m.memory[int64(uint32(v3))+64:], uint32(i32(10)))
								m.memory[int64(uint32(v3))+88] = byte(i32(1))
								m.fn198(v3+i32(152), v3+i32(64))
								{
									{
										{
											t103 := int32(load32(m.memory[int64(uint32(v3))+152:]))
											if t103 != i32(1) {
												goto l76
											}
											t104 := int32(load32(m.memory[int64(uint32(v3))+92:]))
											v11 = t104
											t105 := int32(load32(m.memory[int64(uint32(v3))+160:]))
											t106 := v3
											v12 = t105
											store32(m.memory[int64(uint32(t106))+92:], uint32(v12))
											v8 = v6 + v11
											v6 = v12 - v11
											goto l77
										}
									l76:
										t107 := int32(m.memory[int64(uint32(v3))+101])
										if t107 != 0 {
											goto l78
										}
										m.memory[int64(uint32(v3))+101] = byte(i32(1))
										{
											{
												t108 := int32(m.memory[int64(uint32(v3))+100])
												if t108 != i32(1) {
													goto l79
												}
												t109 := int32(load32(m.memory[int64(uint32(v3))+96:]))
												v11 = t109
												t110 := int32(load32(m.memory[int64(uint32(v3))+92:]))
												v6 = t110
												goto l80
											}
										l79:
											t111 := int32(load32(m.memory[int64(uint32(v3))+96:]))
											v11 = t111
											t112 := int32(load32(m.memory[int64(uint32(v3))+92:]))
											t113 := v11
											v6 = t112
											if t113 == v6 {
												goto l78
											}
										}
									l80:
										t114 := int32(load32(m.memory[int64(uint32(v3))+68:]))
										v8 = t114 + v6
										v6 = v11 - v6
									}
								l77:
									if v6 == 0 {
										goto l81
									}
									t115 := v8
									v11 = v6 + i32(-1)
									t116 := int32(m.memory[uint32(t115+v11)])
									if t116 != i32(10) {
										goto l81
									}
									v6 = v6 + i32(-2)
									{
										if v11 != 0 {
											goto l82
										}
										v12 = i32(0)
										goto l83
									l82:
										t117 := int32(m.memory[uint32(v8+v6)])
										p118 := i32(0)
										if t117&i32(255) == i32(13) {
											p118 = v8
										}
										v12 = p118
									}
								l83:
									p119 := v11
									if v12 != 0 {
										p119 = v6
									}
									v6 = p119
									p120 := v8
									if v12 != 0 {
										p120 = v12
									}
									v8 = p120
									goto l81
								}
							l78:
								v6 = i32(0)
								v8 = i32(1)
							l81:
								v11 = v9 | v10
								store32(m.memory[int64(uint32(v3))+112:], uint32(v6))
								store32(m.memory[int64(uint32(v3))+108:], uint32(v8))
								store64(m.memory[int64(uint32(v3))+168:], uint64(v16))
								store64(m.memory[int64(uint32(v3))+160:], uint64(v17))
								store64(m.memory[int64(uint32(v3))+152:], uint64(v18))
								m.fn13(v3+i32(116), i32(0x100069), v3+i32(152))
								t121 := int64(load64(m.memory[int64(uint32(v3))+96:]))
								store64(m.memory[int64(uint32(v3))+184:], uint64(t121))
								t122 := int64(load64(m.memory[int64(uint32(v3))+88:]))
								store64(m.memory[int64(uint32(v3))+176:], uint64(t122))
								t123 := int64(load64(m.memory[int64(uint32(v3))+80:]))
								store64(m.memory[int64(uint32(v3))+168:], uint64(t123))
								t124 := int64(load64(m.memory[int64(uint32(v3))+72:]))
								store64(m.memory[int64(uint32(v3))+160:], uint64(t124))
								t125 := int64(load64(m.memory[int64(uint32(v3))+64:]))
								store64(m.memory[int64(uint32(v3))+152:], uint64(t125))
								t126 := int32(m.memory[int64(uint32(v3))+189])
								if t126 != 0 {
									goto l84
								}
							l99:
								{
									t127 := int32(load32(m.memory[int64(uint32(v3))+156:]))
									v6 = t127
									m.fn198(v3+i32(192), v3+i32(152))
									{
										{
											t128 := int32(load32(m.memory[int64(uint32(v3))+192:]))
											if t128 != i32(1) {
												goto l85
											}
											t129 := int32(load32(m.memory[int64(uint32(v3))+180:]))
											v8 = t129
											t130 := int32(load32(m.memory[int64(uint32(v3))+200:]))
											t131 := v3
											v10 = t130
											store32(m.memory[int64(uint32(t131))+180:], uint32(v10))
											v9 = v6 + v8
											v6 = v10 - v8
											goto l86
										}
									l85:
										t132 := int32(m.memory[int64(uint32(v3))+189])
										if t132 != 0 {
											goto l84
										}
										m.memory[int64(uint32(v3))+189] = byte(i32(1))
										{
											{
												t133 := int32(m.memory[int64(uint32(v3))+188])
												if t133 != i32(1) {
													goto l87
												}
												t134 := int32(load32(m.memory[int64(uint32(v3))+184:]))
												v8 = t134
												t135 := int32(load32(m.memory[int64(uint32(v3))+180:]))
												v6 = t135
												goto l88
											}
										l87:
											t136 := int32(load32(m.memory[int64(uint32(v3))+184:]))
											v8 = t136
											t137 := int32(load32(m.memory[int64(uint32(v3))+180:]))
											t138 := v8
											v6 = t137
											if t138 == v6 {
												goto l84
											}
										}
									l88:
										t139 := int32(load32(m.memory[int64(uint32(v3))+156:]))
										v9 = t139 + v6
										v6 = v8 - v6
									}
								l86:
									{
										if v6 == 0 {
											goto l89
										}
										t140 := v9
										v8 = v6 + i32(-1)
										t141 := int32(m.memory[uint32(t140+v8)])
										if t141 != i32(10) {
											goto l89
										}
										v6 = v6 + i32(-2)
										{
											if v8 != 0 {
												goto l90
											}
											v10 = i32(0)
											goto l91
										l90:
											t142 := int32(m.memory[uint32(v9+v6)])
											p143 := i32(0)
											if t142&i32(255) == i32(13) {
												p143 = v9
											}
											v10 = p143
										}
									l91:
										p144 := v8
										if v10 != 0 {
											p144 = v6
										}
										v6 = p144
										p145 := v9
										if v10 != 0 {
											p145 = v10
										}
										v9 = p145
									}
								l89:
									{
										t146 := int32(load32(m.memory[int64(uint32(v3))+116:]))
										t147 := int32(load32(m.memory[int64(uint32(v3))+124:]))
										v8 = t147
										if t146 != v8 {
											goto l92
										}
										m.fn196(v3+i32(116), v8, i32(1), i32(1), i32(1))
									}
								l92:
									t148 := int32(load32(m.memory[int64(uint32(v3))+120:]))
									v12 = t148
									m.memory[uint32(v12+v8)] = byte(i32(10))
									v10 = i32(1)
									t149 := v3
									v8 = v8 + i32(1)
									store32(m.memory[int64(uint32(t149))+124:], uint32(v8))
									{
										if v6 == 0 {
											goto l93
										}
										{
											t150 := int32(load32(m.memory[int64(uint32(v3))+116:]))
											t151 := v4
											v10 = t150
											if uint32(t151) <= uint32(v10-v8) {
												goto l94
											}
											m.fn196(v3+i32(116), v8, v4, i32(1), i32(1))
											t152 := int32(load32(m.memory[int64(uint32(v3))+116:]))
											v10 = t152
											t153 := int32(load32(m.memory[int64(uint32(v3))+120:]))
											v12 = t153
											t154 := int32(load32(m.memory[int64(uint32(v3))+124:]))
											v8 = t154
											goto l95
										}
									l94:
										if v4 == 0 {
											goto l96
										}
									l95:
										if v4 == 0 {
											goto l96
										}
										memory_copy(m.memory, uint32(v12+v8), uint32(v7), uint32(v4))
									l96:
										t155 := v3
										v8 = v8 + v4
										store32(m.memory[int64(uint32(t155))+124:], uint32(v8))
										{
											if uint32(v6) <= uint32(v10-v8) {
												goto l97
											}
											m.fn196(v3+i32(116), v8, v6, i32(1), i32(1))
											t156 := int32(load32(m.memory[int64(uint32(v3))+120:]))
											v12 = t156
											t157 := int32(load32(m.memory[int64(uint32(v3))+124:]))
											v8 = t157
										}
									l97:
										if v6 == 0 {
											goto l98
										}
										memory_copy(m.memory, uint32(v12+v8), uint32(v9), uint32(v6))
									l98:
										store32(m.memory[int64(uint32(v3))+124:], uint32(v8+v6))
										v10 = v11
									}
								l93:
									v11 = v10
									t158 := int32(m.memory[int64(uint32(v3))+189])
									if t158 == 0 {
										goto l99
									}
									goto l100
								}
							}
						l84:
							v10 = v11
						l100:
							{
								t159 := int32(load32(m.memory[int64(uint32(v3))+52:]))
								v6 = t159
								t160 := int32(load32(m.memory[int64(uint32(v3))+44:]))
								if v6 != t160 {
									goto l101
								}
								m.fn201(v3 + i32(44))
							}
						l101:
							t161 := int32(load32(m.memory[int64(uint32(v3))+48:]))
							v11 = t161
							v8 = v11 + v6*i32(12)
							t162 := int64(load64(m.memory[int64(uint32(v3))+116:]))
							store64(m.memory[uint32(v8):], uint64(t162))
							t163 := int32(load32(m.memory[int64(uint32(v3))+124:]))
							store32(m.memory[int64(uint32(v8))+8:], uint32(t163))
							t164 := v3
							v6 = v6 + i32(1)
							store32(m.memory[int64(uint32(t164))+52:], uint32(v6))
							{
								if v4 == 0 {
									goto l102
								}
								t165 := int32(load32(m.memory[uint32(v7+i32(-4)):]))
								v8 = t165
								v9 = v8 & i32(-8)
								t166 := v9
								v8 = v8 & i32(3)
								p167 := i32(8)
								if v8 != 0 {
									p167 = i32(4)
								}
								if uint32(t166) < uint32(p167+v4) {
									goto l103
								}
								if v8 == 0 {
									goto l104
								}
								if uint32(v9) > uint32(v4+i32(39)) {
									m.fn3(i32(1274288), i32(46), i32(1274336))
									panic("unreachable")
								}
							l104:
								m.fn1(v7)
							}
						l102:
							{
								t168 := int32(load32(m.memory[int64(uint32(v3))+140:]))
								v4 = t168
								if v4 == 0 {
									goto l106
								}
								t169 := int32(load32(m.memory[int64(uint32(v3))+144:]))
								v9 = t169
								t170 := int32(load32(m.memory[uint32(v9+i32(-4)):]))
								v8 = t170
								v7 = v8 & i32(-8)
								t171 := v7
								v8 = v8 & i32(3)
								p172 := i32(8)
								if v8 != 0 {
									p172 = i32(4)
								}
								if uint32(t171) < uint32(p172+v4) {
									m.fn3(i32(1274224), i32(46), i32(1274272))
									panic("unreachable")
								}
								if v8 == 0 {
									goto l108
								}
								if uint32(v7) > uint32(v4+i32(39)) {
									m.fn3(i32(1274288), i32(46), i32(1274336))
									panic("unreachable")
								}
							l108:
								m.fn1(v9)
							}
						l106:
							{
								t173 := int32(load32(m.memory[int64(uint32(v3))+128:]))
								v4 = t173
								if v4 == 0 {
									goto l110
								}
								t174 := int32(load32(m.memory[int64(uint32(v3))+132:]))
								v9 = t174
								t175 := int32(load32(m.memory[uint32(v9+i32(-4)):]))
								v8 = t175
								v7 = v8 & i32(-8)
								t176 := v7
								v8 = v8 & i32(3)
								p177 := i32(8)
								if v8 != 0 {
									p177 = i32(4)
								}
								if uint32(t176) < uint32(p177+v4) {
									m.fn3(i32(1274224), i32(46), i32(1274272))
									panic("unreachable")
								}
								if v8 == 0 {
									goto l112
								}
								if uint32(v7) > uint32(v4+i32(39)) {
									m.fn3(i32(1274288), i32(46), i32(1274336))
									panic("unreachable")
								}
							l112:
								m.fn1(v9)
							}
						l110:
							v15 = v15 + i64(1)
							v5 = v5 + i32(28)
							if v5 == v13 {
								goto l114
							}
							goto l115
						l103:
						}
						m.fn3(i32(1274224), i32(46), i32(1274272))
						panic("unreachable")
					}
					store32(m.memory[uint32(v0):], uint32(i32(-1)))
					goto l45
				}
			}
		l39:
			store32(m.memory[int64(uint32(v0))+8:], uint32(v12))
			store32(m.memory[int64(uint32(v0))+4:], uint32(v11))
			store32(m.memory[uint32(v0):], uint32(v10))
		l40:
			t201 := int32(load32(m.memory[int64(uint32(v3))+140:]))
			v4 = t201
			if v4 == 0 {
				goto l45
			}
			{
				t202 := int32(load32(m.memory[int64(uint32(v3))+144:]))
				v8 = t202
				t203 := int32(load32(m.memory[uint32(v8+i32(-4)):]))
				v6 = t203
				v9 = v6 & i32(-8)
				t204 := v9
				v6 = v6 & i32(3)
				p205 := i32(8)
				if v6 != 0 {
					p205 = i32(4)
				}
				if uint32(t204) < uint32(p205+v4) {
					m.fn3(i32(1274224), i32(46), i32(1274272))
					panic("unreachable")
				}
				if v6 == 0 {
					goto l129
				}
				if uint32(v9) > uint32(v4+i32(39)) {
					m.fn3(i32(1274288), i32(46), i32(1274336))
					panic("unreachable")
				}
			l129:
				m.fn1(v8)
				goto l45
			}
		}
	l114:
		t529 := v0
		t530 := v11
		t531 := v6
		v4 = v10 & i32(1)
		p532 := i32(1099470)
		if v4 != 0 {
			p532 = i32(1076056)
		}
		p533 := i32(1)
		if v4 != 0 {
			p533 = i32(2)
		}
		m.fn202(t529, t530, t531, p532, p533)
		if v6 == 0 {
			goto l350
		}
		v4 = v11
	l355:
		{
			t534 := int32(load32(m.memory[uint32(v4):]))
			v8 = t534
			if v8 == 0 {
				goto l351
			}
			t535 := int32(load32(m.memory[uint32(v4+i32(4)):]))
			v10 = t535
			t536 := int32(load32(m.memory[uint32(v10+i32(-4)):]))
			v9 = t536
			v2 = v9 & i32(-8)
			t537 := v2
			v9 = v9 & i32(3)
			p538 := i32(8)
			if v9 != 0 {
				p538 = i32(4)
			}
			if uint32(t537) < uint32(p538+v8) {
				m.fn3(i32(1274224), i32(46), i32(1274272))
				panic("unreachable")
			}
			if v9 == 0 {
				goto l353
			}
			if uint32(v2) > uint32(v8+i32(39)) {
				m.fn3(i32(1274288), i32(46), i32(1274336))
				panic("unreachable")
			}
		l353:
			m.fn1(v10)
		}
	l351:
		v4 = v4 + i32(12)
		v6 = v6 + i32(-1)
		if v6 != 0 {
			goto l355
		}
	l350:
		t539 := int32(load32(m.memory[int64(uint32(v3))+44:]))
		v4 = t539
		if v4 == 0 {
			goto l45
		}
		m.fn17(v11, v4*i32(12), i32(4))
	}
l45:
	m.g0 = v3 + i32(208)
}
func (m *Module) fn775(v0, v1 int32) {
	var v2, v3, v4, v5 int32
	var v6, v7 int64
	var v8, v9, v10 int32
	var v11 int64
	var v12 int32
	var v13 int64
	var v14, v15 int32
	var v16 int64
	var v17 int32
	v2 = i32(0)
	{
		{
			t0 := int32(load32(m.memory[uint32(v1):]))
			v3 = t0
			t1 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t2 := v3
			v4 = t1
			if t2 != v4 {
				goto l0
			}
			goto l1
		}
	l0:
		{
			t3 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			v5 = t3
			t4 := int32(load32(m.memory[int64(uint32(v5))+12:]))
			if t4 != 0 {
				goto l2
			}
		l3:
			v3 = v3 + i32(28)
			if v3 != v4 {
				goto l3
			}
			store32(m.memory[uint32(v1):], uint32(v3))
			goto l1
		}
	l2:
		t5 := int64(load64(m.memory[int64(uint32(v5))+24:]))
		v6 = t5
		t6 := int64(load64(m.memory[int64(uint32(v5))+16:]))
		v7 = t6
	l10:
		{
			t7 := v1
			v8 = v3 + i32(28)
			store32(m.memory[uint32(t7):], uint32(v8))
			t8 := int32(load32(m.memory[uint32(v3+i32(4)):]))
			t9 := v7
			t10 := v6
			v9 = t8
			t11 := int32(load32(m.memory[uint32(v3+i32(8)):]))
			t12 := v9
			v10 = t11
			t13 := m.fn64(t9, t10, t12, v10)
			v11 = t13
			t14 := int32(load32(m.memory[int64(uint32(v5))+4:]))
			v12 = t14
			v2 = v12 & int32(v11)
			v13 = int64(uint64(v11)>>25) & i64(127) * i64(72340172838076673)
			t15 := int32(load32(m.memory[uint32(v5):]))
			v14 = t15
			v15 = i32(0)
		l9:
			{
				{
					t16 := int64(load64(m.memory[uint32(v14+v2):]))
					v16 = t16
					v11 = v16 ^ v13
					v11 = (v11 ^ i64(-1)) & (v11 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
					if v11 == 0 {
						goto l4
					}
				l7:
					{
						t17 := v10
						v17 = v14 - (int32(uint32(int64(bits.TrailingZeros64(uint64(v11))))>>3)+v2)&v12<<4
						t18 := int32(load32(m.memory[uint32(v17+i32(-8)):]))
						if t17 != t18 {
							goto l5
						}
						t19 := int32(load32(m.memory[uint32(v17+i32(-12)):]))
						t20 := m.fn973(v9, t19, v10)
						if t20 == 0 {
							t22 := int32(load32(m.memory[uint32(v17+i32(-4)):]))
							v14 = t22
							v2 = v3
							goto l1
						}
					}
				l5:
					v11 = (v11 + i64(-1)) & v11
					if !(v11 == 0) {
						goto l7
					}
				}
			l4:
				if !(v16&(v16<<1)&i64(-0x7f7f7f7f7f7f7f80) == 0) {
					goto l8
				}
				t21 := v2
				v15 = v15 + i32(8)
				v2 = (t21 + v15) & v12
				goto l9
			}
		l8:
			v2 = i32(0)
			v3 = v8
			if v8 != v4 {
				goto l10
			}
		}
	}
l1:
	store32(m.memory[uint32(v0):], uint32(v2))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v14))
}
func (m *Module) fn776(v0, v1 int32) {
	var v2, v3, v4, v5, v6, v7, v8 int32
	v2 = v0 + i32(8)
	v3 = v0 + v1<<3
	v4 = i32(0)
	v1 = v0
l4:
	v5 = v2
	{
		t0 := int32(load32(m.memory[uint32(v1+i32(12)):]))
		v6 = t0
		t1 := int32(load32(m.memory[uint32(v1+i32(4)):]))
		if uint32(v6) >= uint32(t1) {
			goto l0
		}
		t2 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		v7 = t2
		v1 = v4
	l2:
		{
			v2 = v0 + v1
			v8 = v2 + i32(8)
			t3 := int64(load64(m.memory[uint32(v2):]))
			store64(m.memory[uint32(v8):], uint64(t3))
			if v1 == 0 {
				goto l1
			}
			v1 = v1 + i32(-8)
			t4 := int32(load32(m.memory[uint32(v2+i32(-4)):]))
			if uint32(v6) < uint32(t4) {
				goto l2
			}
		}
		v1 = v0 + v1 + i32(8)
		goto l3
	l1:
		v1 = v0
	l3:
		store32(m.memory[uint32(v1):], uint32(v7))
		store32(m.memory[uint32(v8+i32(-4)):], uint32(v6))
	}
l0:
	v4 = v4 + i32(8)
	v1 = v5
	v2 = v5 + i32(8)
	if v2 != v3 {
		goto l4
	}
}
func (m *Module) fn777(v0, v1, v2, v3 int32) {
	var v4, v5, v6, v7, v8 int32
	t0 := m.g0
	v4 = t0 - i32(48)
	m.g0 = v4
	t1 := v1
	v2 = v2 << 5
	v5 = t1 + v2
	{
	l1:
		{
			if v2 == 0 {
				m.fn202(v0, i32(4), i32(0), i32(1076056), i32(2))
				goto l11
			}
			m.fn774(v4+i32(24), v1, v3)
			v2 = v2 + i32(-32)
			v1 = v1 + i32(32)
			t2 := int32(load32(m.memory[int64(uint32(v4))+24:]))
			if t2 == i32(-1) {
				goto l1
			}
		}
		t3 := m.fn7(i32(48))
		v6 = t3
		if v6 == 0 {
			m.fn12(i32(4), i32(48))
			panic("unreachable")
		}
		t4 := int32(load32(m.memory[int64(uint32(v4))+32:]))
		store32(m.memory[int64(uint32(v6))+8:], uint32(t4))
		t5 := int64(load64(m.memory[int64(uint32(v4))+24:]))
		store64(m.memory[uint32(v6):], uint64(t5))
		store32(m.memory[int64(uint32(v4))+20:], uint32(i32(1)))
		store32(m.memory[int64(uint32(v4))+16:], uint32(v6))
		store32(m.memory[int64(uint32(v4))+12:], uint32(i32(4)))
		v2 = i32(1)
	l4:
		{
			if v1 == v5 {
				t12 := int32(load32(m.memory[int64(uint32(v4))+12:]))
				v8 = t12
				t13 := int32(load32(m.memory[int64(uint32(v4))+16:]))
				t14 := v0
				v7 = t13
				m.fn202(t14, v7, v2, i32(1076056), i32(2))
				v1 = v7
			l10:
				{
					t15 := int32(load32(m.memory[uint32(v1):]))
					v3 = t15
					if v3 == 0 {
						goto l6
					}
					t16 := int32(load32(m.memory[uint32(v1+i32(4)):]))
					v6 = t16
					t17 := int32(load32(m.memory[uint32(v6+i32(-4)):]))
					v5 = t17
					v0 = v5 & i32(-8)
					t18 := v0
					v5 = v5 & i32(3)
					p19 := i32(8)
					if v5 != 0 {
						p19 = i32(4)
					}
					if uint32(t18) < uint32(p19+v3) {
						m.fn3(i32(1274224), i32(46), i32(1274272))
						panic("unreachable")
					}
					if v5 == 0 {
						goto l8
					}
					if uint32(v0) > uint32(v3+i32(39)) {
						m.fn3(i32(1274288), i32(46), i32(1274336))
						panic("unreachable")
					}
				l8:
					m.fn1(v6)
				}
			l6:
				v1 = v1 + i32(12)
				v2 = v2 + i32(-1)
				if v2 != 0 {
					goto l10
				}
				if v8 == 0 {
					goto l11
				}
				t20 := int32(load32(m.memory[uint32(v7+i32(-4)):]))
				v1 = t20
				v2 = v1 & i32(-8)
				t21 := v2
				v1 = v1 & i32(3)
				p22 := i32(8)
				if v1 != 0 {
					p22 = i32(4)
				}
				v3 = v8 * i32(12)
				if uint32(t21) < uint32(p22+v3) {
					m.fn3(i32(1274224), i32(46), i32(1274272))
					panic("unreachable")
				}
				if v1 == 0 {
					goto l13
				}
				if uint32(v2) > uint32(v3+i32(39)) {
					m.fn3(i32(1274288), i32(46), i32(1274336))
					panic("unreachable")
				}
			l13:
				m.fn1(v7)
				goto l11
			}
			m.fn774(v4+i32(36), v1, v3)
			v1 = v1 + i32(32)
			t6 := int32(load32(m.memory[int64(uint32(v4))+36:]))
			if t6 == i32(-1) {
				goto l4
			}
			{
				t7 := int32(load32(m.memory[int64(uint32(v4))+12:]))
				if v2 != t7 {
					goto l5
				}
				m.fn196(v4+i32(12), v2, i32(1), i32(4), i32(12))
				t8 := int32(load32(m.memory[int64(uint32(v4))+16:]))
				v6 = t8
			}
		l5:
			v7 = v6 + v2*i32(12)
			t9 := int32(load32(m.memory[int64(uint32(v4))+44:]))
			store32(m.memory[int64(uint32(v7))+8:], uint32(t9))
			t10 := int64(load64(m.memory[int64(uint32(v4))+36:]))
			store64(m.memory[uint32(v7):], uint64(t10))
			t11 := v4
			v2 = v2 + i32(1)
			store32(m.memory[int64(uint32(t11))+20:], uint32(v2))
			goto l4
		}
	}
l11:
	m.g0 = v4 + i32(48)
}
func (m *Module) fn778(v0, v1, v2, v3, v4 int32) {
	var v5, v6, v7, v8 int32
	var v9, v10 int64
	var v11 int32
	var v12 int64
	var v13 int32
	var v14 int64
	var v15 int32
	var v16 int64
	var v17 int32
	t0 := m.g0
	v5 = t0 - i32(16)
	m.g0 = v5
	{
		if v1 == 0 {
			goto l0
		}
		v1 = v1 * i32(28)
		{
			t1 := int32(load32(m.memory[int64(uint32(v2))+12:]))
			if t1 != 0 {
				goto l1
			}
		l3:
			{
				t2 := int32(load32(m.memory[uint32(v0):]))
				if uint32(t2) > uint32(i32(2)) {
					goto l2
				}
				t3 := int32(load32(m.memory[uint32(v0+i32(20)):]))
				t4 := int32(load32(m.memory[uint32(v0+i32(24)):]))
				m.fn778(t3, t4, v2, v3, v4)
			}
		l2:
			v0 = v0 + i32(28)
			v1 = v1 + i32(-28)
			if v1 != 0 {
				goto l3
			}
			goto l0
		}
	l1:
		v6 = v0 + v1
		t5 := int32(load32(m.memory[uint32(v2):]))
		v7 = t5
		t6 := int32(load32(m.memory[int64(uint32(v2))+4:]))
		v8 = t6
		t7 := int64(load64(m.memory[int64(uint32(v2))+24:]))
		v9 = t7
		t8 := int64(load64(m.memory[int64(uint32(v2))+16:]))
		v10 = t8
	l19:
		{
			t9 := int32(load32(m.memory[uint32(v0):]))
			v1 = t9
			p10 := i32(1)
			if uint32(v1) > uint32(i32(2)) {
				p10 = v1 + i32(-3)
			}
			switch p10 + i32(-1) {
			default:
				goto l5
			case 0:
				t11 := int32(load32(m.memory[int64(uint32(v0))+20:]))
				t12 := int32(load32(m.memory[int64(uint32(v0))+24:]))
				m.fn778(t11, t12, v2, v3, v4)
				goto l5
			case 3:
				t13 := int32(load32(m.memory[uint32(v0+i32(8)):]))
				t14 := v8
				t15 := v10
				t16 := v9
				v11 = t13
				t17 := int32(load32(m.memory[uint32(v0+i32(12)):]))
				t18 := v11
				v1 = t17
				t19 := m.fn249(t15, t16, t18, v1)
				v12 = t19
				v13 = t14 & int32(v12)
				v14 = int64(uint64(v12)>>25) & i64(127) * i64(72340172838076673)
				v15 = i32(0)
			l11:
				{
					{
						t20 := int64(load64(m.memory[uint32(v7+v13):]))
						v16 = t20
						v12 = v16 ^ v14
						v12 = (v12 ^ i64(-1)) & (v12 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
						if v12 == 0 {
							goto l7
						}
					l10:
						{
							t21 := v1
							v17 = v7 + (i32(0)-(int32(uint32(int64(bits.TrailingZeros64(uint64(v12))))>>3)+v13)&v8)*i32(12)
							t22 := int32(load32(m.memory[uint32(v17+i32(-8)):]))
							if t21 != t22 {
								goto l8
							}
							t23 := int32(load32(m.memory[uint32(v17+i32(-12)):]))
							t24 := m.fn973(v11, t23, v1)
							if t24 == 0 {
								{
									if v1 != 0 {
										goto l12
									}
									v13 = i32(1)
									goto l13
								l12:
									t26 := m.fn7(v1)
									v13 = t26
									if v13 == 0 {
										m.fn12(i32(1), v1)
										panic("unreachable")
									}
									if v1 == 0 {
										goto l13
									}
									memory_copy(m.memory, uint32(v13), uint32(v11), uint32(v1))
								}
							l13:
								store32(m.memory[int64(uint32(v5))+12:], uint32(v1))
								store32(m.memory[int64(uint32(v5))+8:], uint32(v13))
								store32(m.memory[int64(uint32(v5))+4:], uint32(v1))
								t27 := m.fn442(v4, v5+i32(4))
								if t27 != 0 {
									goto l5
								}
								{
									if v1 != 0 {
										goto l15
									}
									v15 = i32(1)
									goto l16
								l15:
									t28 := m.fn7(v1)
									v15 = t28
									if v15 == 0 {
										m.fn12(i32(1), v1)
										panic("unreachable")
									}
									if v1 == 0 {
										goto l16
									}
									memory_copy(m.memory, uint32(v15), uint32(v11), uint32(v1))
								}
							l16:
								{
									t29 := int32(load32(m.memory[int64(uint32(v3))+8:]))
									v13 = t29
									t30 := int32(load32(m.memory[uint32(v3):]))
									if v13 != t30 {
										goto l18
									}
									m.fn201(v3)
								}
							l18:
								t31 := int32(load32(m.memory[int64(uint32(v3))+4:]))
								v11 = t31 + v13*i32(12)
								store32(m.memory[int64(uint32(v11))+8:], uint32(v1))
								store32(m.memory[int64(uint32(v11))+4:], uint32(v15))
								store32(m.memory[uint32(v11):], uint32(v1))
								store32(m.memory[int64(uint32(v3))+8:], uint32(v13+i32(1)))
								t32 := int32(load32(m.memory[uint32(v17+i32(-4)):]))
								v1 = t32
								t33 := int32(load32(m.memory[int64(uint32(v1))+16:]))
								t34 := int32(load32(m.memory[int64(uint32(v1))+20:]))
								m.fn765(t33, t34, v2, v3, v4)
								goto l5
							}
						}
					l8:
						v12 = (v12 + i64(-1)) & v12
						if !(v12 == 0) {
							goto l10
						}
					}
				l7:
					if !(v16&(v16<<1)&i64(-0x7f7f7f7f7f7f7f80) == 0) {
						goto l5
					}
					t25 := v13
					v15 = v15 + i32(8)
					v13 = (t25 + v15) & v8
					goto l11
				}
			}
		}
	l5:
		v0 = v0 + i32(28)
		if v0 != v6 {
			goto l19
		}
	}
l0:
	m.g0 = v5 + i32(16)
}
func (m *Module) fn779(v0, v1, v2 int32) {
	var v3, v4 int32
	var v5 int64
	var v6, v7 int32
	var v8 int64
	var v9, v10 int32
	var v11, v12 int64
	var v13 int32
	t0 := int64(load64(m.memory[int64(uint32(v1))+16:]))
	t1 := int64(load64(m.memory[int64(uint32(v1))+24:]))
	t2 := int32(load32(m.memory[int64(uint32(v2))+4:]))
	v3 = t2
	t3 := int32(load32(m.memory[int64(uint32(v2))+8:]))
	t4 := v3
	v4 = t3
	t5 := m.fn64(t0, t1, t4, v4)
	v5 = t5
	t6 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	v6 = t6
	v7 = v6 & int32(v5)
	v8 = int64(uint64(v5)>>25) & i64(127) * i64(72340172838076673)
	t7 := int32(load32(m.memory[uint32(v1):]))
	v9 = t7
	v10 = i32(0)
	{
	l5:
		{
			{
				t8 := int64(load64(m.memory[uint32(v9+v7):]))
				v11 = t8
				v12 = v11 ^ v8
				v12 = (v12 ^ i64(-1)) & (v12 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
				if v12 == 0 {
					goto l0
				}
			l3:
				{
					v13 = v9 - (int32(uint32(int64(bits.TrailingZeros64(uint64(v12))))>>3)+v7)&v6<<4
					t9 := int32(load32(m.memory[uint32(v13+i32(-8)):]))
					if t9 != v4 {
						goto l1
					}
					t10 := int32(load32(m.memory[uint32(v13+i32(-12)):]))
					t11 := m.fn973(t10, v3, v4)
					if t11 == 0 {
						store32(m.memory[int64(uint32(v0))+8:], uint32(i32(-1)))
						store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
						store32(m.memory[uint32(v0):], uint32(v13))
						t13 := int32(load32(m.memory[uint32(v2):]))
						v1 = t13
						if v1 == 0 {
							return
						}
						t14 := int32(load32(m.memory[uint32(v3+i32(-4)):]))
						v0 = t14
						v2 = v0 & i32(-8)
						t15 := v2
						v0 = v0 & i32(3)
						p16 := i32(8)
						if v0 != 0 {
							p16 = i32(4)
						}
						if uint32(t15) < uint32(p16+v1) {
							m.fn3(i32(1274224), i32(46), i32(1274272))
							panic("unreachable")
						}
						if v0 == 0 {
							goto l8
						}
						if uint32(v2) > uint32(v1+i32(39)) {
							m.fn3(i32(1274288), i32(46), i32(1274336))
							panic("unreachable")
						}
					l8:
						m.fn1(v3)
						return
					}
				}
			l1:
				v12 = (v12 + i64(-1)) & v12
				if !(v12 == 0) {
					goto l3
				}
			}
		l0:
			if !(v11&(v11<<1)&i64(-0x7f7f7f7f7f7f7f80) == 0) {
				goto l4
			}
			t12 := v7
			v10 = v10 + i32(8)
			v7 = (t12 + v10) & v6
			goto l5
		}
	l4:
		{
			t17 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			if t17 != 0 {
				goto l10
			}
			_ = m.fn77(v1, i32(1), v1+i32(16))
		}
	l10:
		store32(m.memory[int64(uint32(v0))+20:], uint32(v1))
		store64(m.memory[uint32(v0):], uint64(v5))
		t19 := int32(load32(m.memory[int64(uint32(v2))+8:]))
		store32(m.memory[int64(uint32(v0))+16:], uint32(t19))
		t20 := int64(load64(m.memory[uint32(v2):]))
		store64(m.memory[int64(uint32(v0))+8:], uint64(t20))
	}
}
func (m *Module) fn780(v0, v1, v2, v3, v4, v5 int32) {
	var v6, v7, v8, v9, v10 int32
	var v11, v12, v13, v14, v15, v16, v17 int64
	var v18, v19 int32
	var v20, v21 int64
	var v22, v23, v24 int32
	var v25, v26 int64
	var v27, v28, v29, v30, v31, v32 int32
	var v33 int64
	var v34, v35, v36, v37, v38, v39, v40, v41, v42, v43 int32
	t0 := m.g0
	v6 = t0 - i32(160)
	m.g0 = v6
	m.fn788(v6+i32(68), v1, v2, v5)
	store32(m.memory[int64(uint32(v6))+88:], uint32(i32(0)))
	store64(m.memory[int64(uint32(v6))+80:], uint64(i64(0x100000000)))
	t1 := int32(load32(m.memory[int64(uint32(v6))+72:]))
	v2 = t1
	{
		{
			t2 := int32(load32(m.memory[int64(uint32(v6))+76:]))
			v1 = t2
			if v1 != 0 {
				goto l0
			}
			t3 := int32(load32(m.memory[int64(uint32(v6))+88:]))
			store32(m.memory[int64(uint32(v0))+8:], uint32(t3))
			t4 := int64(load64(m.memory[int64(uint32(v6))+80:]))
			store64(m.memory[uint32(v0):], uint64(t4))
			goto l1
		}
	l0:
		v7 = v2 + v1<<4
		p5 := i32(256)
		if v4 != 0 {
			p5 = i32(16777472)
		}
		v8 = p5
		p6 := i32(0)
		if v4 != 0 {
			p6 = i32(0x1000000)
		}
		v9 = p6
		v10 = v5 + i32(32)
		v11 = int64(uint32(i32(17))) << 32
		t7 := v11
		v12 = int64(uint32(v6 + i32(148)))
		v13 = t7 | v12
		v14 = v11 | int64(uint32(v6+i32(92)))
		v15 = int64(uint32(i32(1)))<<32 | v12
		v16 = v11 | int64(uint32(v6+i32(116)))
		v17 = int64(uint32(i32(34)))<<32 | v12
		t8 := int32(load32(m.memory[int64(uint32(v5))+32:]))
		v18 = t8
		t9 := int32(load32(m.memory[int64(uint32(v5))+36:]))
		v19 = t9
		t10 := int64(load64(m.memory[int64(uint32(v5))+56:]))
		v20 = t10
		t11 := int64(load64(m.memory[int64(uint32(v5))+48:]))
		v21 = t11
		t12 := int32(load32(m.memory[int64(uint32(v5))+44:]))
		v22 = t12
		t13 := int32(load32(m.memory[uint32(v5):]))
		v23 = t13
		t14 := int32(load32(m.memory[int64(uint32(v5))+4:]))
		v24 = t14
		t15 := int64(load64(m.memory[int64(uint32(v5))+24:]))
		v25 = t15
		t16 := int64(load64(m.memory[int64(uint32(v5))+16:]))
		v26 = t16
		t17 := int32(load32(m.memory[int64(uint32(v5))+12:]))
		v27 = t17
		v4 = i32(0)
	l161:
		v4 = v4 + i32(1)
		{
			{
				{
					{
						{
							{
								{
									t18 := int32(load32(m.memory[uint32(v2):]))
									v1 = t18
									p19 := i32(0)
									if v1 < i32(-0x7ffffffb) {
										p19 = v1 + i32(-0x7fffffff)
									}
									switch p19 {
									case 2:
										t107 := int32(load32(m.memory[int64(uint32(v2))+8:]))
										v28 = t107
										t108 := int32(load32(m.memory[int64(uint32(v2))+4:]))
										v30 = t108
										{
											t109 := int32(load32(m.memory[int64(uint32(v2))+12:]))
											v1 = t109
											t110 := int32(load32(m.memory[uint32(v1):]))
											if t110 < i32(0) {
												m.fn143(v6+i32(48), v30, v28)
												t126 := int32(load32(m.memory[int64(uint32(v6))+52:]))
												if t126 == 0 {
													goto l45
												}
												m.fn143(v6+i32(40), v30, v28)
												t127 := int32(load32(m.memory[int64(uint32(v6))+40:]))
												t128 := int32(load32(m.memory[int64(uint32(v6))+44:]))
												m.fn790(v6+i32(128), t127, t128, v3, v9)
												t129 := int32(load32(m.memory[int64(uint32(v6))+132:]))
												v28 = t129
												{
													{
														t130 := int32(load32(m.memory[int64(uint32(v6))+136:]))
														v1 = t130
														t131 := int32(load32(m.memory[int64(uint32(v6))+80:]))
														t132 := int32(load32(m.memory[int64(uint32(v6))+88:]))
														t133 := v1
														v30 = t132
														if uint32(t133) <= uint32(t131-v30) {
															goto l65
														}
														m.fn196(v6+i32(80), v30, v1, i32(1), i32(1))
														t134 := int32(load32(m.memory[int64(uint32(v6))+88:]))
														v30 = t134
														goto l66
													}
												l65:
													if v1 == 0 {
														goto l67
													}
												l66:
													if v1 == 0 {
														goto l67
													}
													t135 := int32(load32(m.memory[int64(uint32(v6))+84:]))
													memory_copy(m.memory, uint32(t135+v30), uint32(v28), uint32(v1))
												}
											l67:
												store32(m.memory[int64(uint32(v6))+88:], uint32(v30+v1))
												t136 := int32(load32(m.memory[int64(uint32(v6))+128:]))
												v1 = t136
												if v1 == 0 {
													goto l45
												}
												t137 := int32(load32(m.memory[uint32(v28+i32(-4)):]))
												v30 = t137
												v29 = v30 & i32(-8)
												t138 := v29
												v30 = v30 & i32(3)
												p139 := i32(8)
												if v30 != 0 {
													p139 = i32(4)
												}
												if uint32(t138) < uint32(p139+v1) {
													m.fn3(i32(1274224), i32(46), i32(1274272))
													panic("unreachable")
												}
												if v30 == 0 {
													goto l69
												}
												if uint32(v29) > uint32(v1+i32(39)) {
													m.fn3(i32(1274288), i32(46), i32(1274336))
													panic("unreachable")
												}
											l69:
												m.fn1(v28)
												goto l45
											}
											m.fn143(v6+i32(32), v30, v28)
											t111 := int32(load32(m.memory[int64(uint32(v6))+32:]))
											t112 := int32(load32(m.memory[int64(uint32(v6))+36:]))
											m.fn790(v6+i32(116), t111, t112, v3, i32(0x1000000))
											t113 := int32(load32(m.memory[int64(uint32(v1))+4:]))
											t114 := int32(load32(m.memory[int64(uint32(v1))+8:]))
											m.fn791(v6+i32(148), t113, t114)
											store64(m.memory[int64(uint32(v6))+136:], uint64(v13))
											store64(m.memory[int64(uint32(v6))+128:], uint64(v16))
											_ = m.fn45(v6+i32(80), i32(1078840), i32(1066703), v6+i32(128))
											{
												t116 := int32(load32(m.memory[int64(uint32(v6))+148:]))
												v1 = t116
												if v1 == 0 {
													goto l58
												}
												t117 := int32(load32(m.memory[int64(uint32(v6))+152:]))
												v30 = t117
												t118 := int32(load32(m.memory[uint32(v30+i32(-4)):]))
												v28 = t118
												v29 = v28 & i32(-8)
												t119 := v29
												v28 = v28 & i32(3)
												p120 := i32(8)
												if v28 != 0 {
													p120 = i32(4)
												}
												if uint32(t119) < uint32(p120+v1) {
													m.fn3(i32(1274224), i32(46), i32(1274272))
													panic("unreachable")
												}
												if v28 == 0 {
													goto l60
												}
												if uint32(v29) > uint32(v1+i32(39)) {
													m.fn3(i32(1274288), i32(46), i32(1274336))
													panic("unreachable")
												}
											l60:
												m.fn1(v30)
											}
										l58:
											t121 := int32(load32(m.memory[int64(uint32(v6))+116:]))
											v1 = t121
											if v1 == 0 {
												goto l45
											}
											t122 := int32(load32(m.memory[int64(uint32(v6))+120:]))
											v30 = t122
											t123 := int32(load32(m.memory[uint32(v30+i32(-4)):]))
											v28 = t123
											v29 = v28 & i32(-8)
											t124 := v29
											v28 = v28 & i32(3)
											p125 := i32(8)
											if v28 != 0 {
												p125 = i32(4)
											}
											if uint32(t124) < uint32(p125+v1) {
												m.fn3(i32(1274224), i32(46), i32(1274272))
												panic("unreachable")
											}
											if v28 == 0 {
												goto l63
											}
											if uint32(v29) > uint32(v1+i32(39)) {
												m.fn3(i32(1274288), i32(46), i32(1274336))
												panic("unreachable")
											}
										l63:
											m.fn1(v30)
											goto l45
										}
									case 4:
										if v27 == 0 {
											goto l45
										}
										t93 := int32(load32(m.memory[int64(uint32(v2))+4:]))
										t94 := v24
										t95 := v26
										t96 := v25
										v29 = t93
										t97 := int32(load32(m.memory[int64(uint32(v2))+8:]))
										t98 := v29
										v28 = t97
										t99 := m.fn249(t95, t96, t98, v28)
										v11 = t99
										v1 = t94 & int32(v11)
										v12 = int64(uint64(v11)>>25) & i64(127) * i64(72340172838076673)
										v31 = i32(0)
									l56:
										{
											{
												t100 := int64(load64(m.memory[uint32(v23+v1):]))
												v33 = t100
												v11 = v33 ^ v12
												v11 = (v11 ^ i64(-1)) & (v11 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
												if v11 == 0 {
													goto l52
												}
											l55:
												{
													t101 := v28
													v30 = v23 - (int32(uint32(int64(bits.TrailingZeros64(uint64(v11))))>>3)+v1)&v24<<4
													t102 := int32(load32(m.memory[uint32(v30+i32(-8)):]))
													if t101 != t102 {
														goto l53
													}
													t103 := int32(load32(m.memory[uint32(v30+i32(-12)):]))
													t104 := m.fn973(v29, t103, v28)
													if t104 == 0 {
														store32(m.memory[int64(uint32(v6))+148:], uint32(v30+i32(-4)))
														store64(m.memory[int64(uint32(v6))+128:], uint64(v17))
														_ = m.fn45(v6+i32(80), i32(1078840), i32(1066411), v6+i32(128))
														goto l45
													}
												}
											l53:
												v11 = (v11 + i64(-1)) & v11
												if !(v11 == 0) {
													goto l55
												}
											}
										l52:
											if !(v33&(v33<<1)&i64(-0x7f7f7f7f7f7f7f80) == 0) {
												goto l45
											}
											t105 := v1
											v31 = v31 + i32(8)
											v1 = (t105 + v31) & v24
											goto l56
										}
									case 5:
										switch v3 & i32(255) {
										default:
											{
												t83 := int32(load32(m.memory[int64(uint32(v6))+80:]))
												t84 := int32(load32(m.memory[int64(uint32(v6))+88:]))
												v1 = t84
												if uint32(t83-v1) > uint32(i32(1)) {
													goto l49
												}
												m.fn196(v6+i32(80), v1, i32(2), i32(1), i32(1))
												t85 := int32(load32(m.memory[int64(uint32(v6))+88:]))
												v1 = t85
											}
										l49:
											t86 := int32(load32(m.memory[int64(uint32(v6))+84:]))
											store16(m.memory[uint32(t86+v1):], uint16(i32(2652)))
											store32(m.memory[int64(uint32(v6))+88:], uint32(v1+i32(2)))
											goto l45
										case 1:
											{
												t87 := int32(load32(m.memory[int64(uint32(v6))+80:]))
												t88 := int32(load32(m.memory[int64(uint32(v6))+88:]))
												v1 = t88
												if t87 != v1 {
													goto l50
												}
												m.fn196(v6+i32(80), v1, i32(1), i32(1), i32(1))
											}
										l50:
											t89 := int32(load32(m.memory[int64(uint32(v6))+84:]))
											m.memory[uint32(t89+v1)] = byte(i32(32))
											store32(m.memory[int64(uint32(v6))+88:], uint32(v1+i32(1)))
											goto l45
										case 2:
											{
												t90 := int32(load32(m.memory[int64(uint32(v6))+80:]))
												t91 := int32(load32(m.memory[int64(uint32(v6))+88:]))
												v1 = t91
												if t90 != v1 {
													goto l51
												}
												m.fn196(v6+i32(80), v1, i32(1), i32(1), i32(1))
											}
										l51:
											t92 := int32(load32(m.memory[int64(uint32(v6))+84:]))
											m.memory[uint32(t92+v1)] = byte(i32(10))
											store32(m.memory[int64(uint32(v6))+88:], uint32(v1+i32(1)))
											goto l45
										}
									default:
										v28 = i32(0)
										t20 := int32(load32(m.memory[int64(uint32(v6))+76:]))
										if uint32(v4) >= uint32(t20) {
											goto l8
										}
										t21 := int32(load32(m.memory[int64(uint32(v6))+72:]))
										v29 = t21 + v4<<4
										t22 := int32(load32(m.memory[uint32(v29):]))
										v1 = t22
										t23 := v1 + i32(-0x7fffffff)
										var p24 int32
										if v1 < i32(-0x7ffffffb) {
											p24 = 1
										}
										v30 = p24
										p25 := i32(0)
										if v30 != 0 {
											p25 = t23
										}
										v1 = p25
										if uint32(v1) > uint32(i32(4)) {
											goto l9
										}
										if i32_shl(i32(1), v1)&i32(22) == 0 {
											goto l9
										}
										goto l10
									case 1:
										t26 := int32(load32(m.memory[int64(uint32(v2))+12:]))
										v28 = t26
										t27 := int32(load32(m.memory[int64(uint32(v2))+4:]))
										t28 := v6 + i32(92)
										v29 = t27
										t29 := int32(load32(m.memory[int64(uint32(v2))+8:]))
										t30 := v29
										v31 = t29
										m.fn780(t28, t30, v31, v3, i32(1), v5)
										t31 := int32(load32(m.memory[uint32(v28+i32(12)):]))
										v1 = t31
										t32 := int32(load32(m.memory[uint32(v28+i32(8)):]))
										v30 = t32
										{
											t33 := int32(load32(m.memory[uint32(v28):]))
											if t33 != i32(2) {
												{
													if v1 != 0 {
														goto l18
													}
													v28 = i32(1)
													goto l19
												l18:
													t42 := m.fn7(v1)
													v28 = t42
													if v28 == 0 {
														m.fn12(i32(1), v1)
														panic("unreachable")
													}
													if v1 == 0 {
														goto l19
													}
													memory_copy(m.memory, uint32(v28), uint32(v30), uint32(v1))
												}
											l19:
												store32(m.memory[int64(uint32(v6))+112:], uint32(v1))
												store32(m.memory[int64(uint32(v6))+108:], uint32(v28))
												store32(m.memory[int64(uint32(v6))+104:], uint32(v1))
												t43 := int32(load32(m.memory[int64(uint32(v6))+96:]))
												t44 := int32(load32(m.memory[int64(uint32(v6))+100:]))
												m.fn143(v6+i32(24), t43, t44)
												t45 := int32(load32(m.memory[int64(uint32(v6))+28:]))
												if t45 != 0 {
													goto l21
												}
												store32(m.memory[int64(uint32(v6))+136:], uint32(i32(0)))
												store64(m.memory[int64(uint32(v6))+128:], uint64(i64(0x100000000)))
												{
													{
														if v1 != 0 {
															goto l22
														}
														t46 := m.fn7(i32(8))
														v30 = t46
														if v30 == 0 {
															m.fn12(i32(1), i32(8))
															panic("unreachable")
														}
														store32(m.memory[int64(uint32(v6))+124:], uint32(i32(0)))
														store32(m.memory[int64(uint32(v6))+120:], uint32(v30))
														store32(m.memory[int64(uint32(v6))+116:], uint32(i32(8)))
														goto l24
													}
												l22:
													v35 = v28 + v1
													t47 := v6 + i32(128)
													t48 := int32(uint32(v1) >> 2)
													var p49 int32
													if v1&i32(3) != i32(0) {
														p49 = 1
													}
													m.fn196(t47, i32(0), t48+p49, i32(1), i32(1))
													v30 = v28
												l36:
													{
														{
															{
																t50 := int32(int8(m.memory[uint32(v30)]))
																v29 = t50
																if v29 <= i32(-1) {
																	goto l25
																}
																v30 = v30 + i32(1)
																v29 = v29 & i32(255)
																goto l26
															}
														l25:
															t51 := int32(m.memory[int64(uint32(v30))+1])
															v31 = t51 & i32(63)
															v34 = v29 & i32(31)
															if uint32(v29) > uint32(i32(-33)) {
																goto l27
															}
															v29 = v34<<6 | v31
															v30 = v30 + i32(2)
															goto l26
														l27:
															t52 := int32(m.memory[int64(uint32(v30))+2])
															v31 = v31<<6 | t52&i32(63)
															if uint32(v29) >= uint32(i32(-16)) {
																goto l28
															}
															v29 = v31 | v34<<12
															v30 = v30 + i32(3)
															goto l26
														l28:
															t53 := int32(m.memory[int64(uint32(v30))+3])
															v29 = v31<<6 | t53&i32(63) | v34<<18&i32(0x1c0000)
															v30 = v30 + i32(4)
														}
													l26:
														t54 := int32(load32(m.memory[int64(uint32(v6))+136:]))
														v31 = t54
														{
															{
																p55 := v29
																if uint32(v29+i32(-127)) < uint32(i32(33)) {
																	p55 = i32(32)
																}
																p56 := p55
																if uint32(v29) < uint32(i32(32)) {
																	p56 = i32(32)
																}
																v29 = p56
																var p57 int32
																if uint32(v29) < uint32(i32(128)) {
																	p57 = 1
																}
																v36 = p57
																if v36 == 0 {
																	goto l29
																}
																v34 = i32(1)
																goto l30
															}
														l29:
															v34 = i32(2)
															if uint32(v29) < uint32(i32(2048)) {
																goto l30
															}
															p58 := i32(4)
															if uint32(v29) < uint32(i32(65536)) {
																p58 = i32(3)
															}
															v34 = p58
														}
													l30:
														{
															t59 := int32(load32(m.memory[int64(uint32(v6))+128:]))
															if uint32(v34) <= uint32(t59-v31) {
																goto l31
															}
															m.fn196(v6+i32(128), v31, v34, i32(1), i32(1))
														}
													l31:
														t60 := int32(load32(m.memory[int64(uint32(v6))+132:]))
														v37 = t60
														v32 = v37 + v31
														if v36 != 0 {
															goto l32
														}
														v36 = v29&i32(63) | i32(-128)
														v38 = int32(uint32(v29) >> 6)
														if uint32(v29) >= uint32(i32(2048)) {
															v39 = int32(uint32(v29) >> 12)
															v38 = v38&i32(63) | i32(-128)
															if uint32(v29) > uint32(i32(0xffff)) {
																m.memory[int64(uint32(v32))+3] = byte(v36)
																m.memory[int64(uint32(v32))+2] = byte(v38)
																m.memory[int64(uint32(v32))+1] = byte(v39&i32(63) | i32(-128))
																m.memory[uint32(v32)] = byte(int32(uint32(v29)>>18) | i32(-16))
																goto l34
															}
															m.memory[int64(uint32(v32))+2] = byte(v36)
															m.memory[int64(uint32(v32))+1] = byte(v38)
															m.memory[uint32(v32)] = byte(v39 | i32(224))
															goto l34
														}
														m.memory[int64(uint32(v32))+1] = byte(v36)
														m.memory[uint32(v32)] = byte(v38 | i32(192))
														goto l34
													l32:
														m.memory[uint32(v32)] = byte(v29)
													l34:
														t61 := v6
														v29 = v34 + v31
														store32(m.memory[int64(uint32(t61))+136:], uint32(v29))
														if v30 != v35 {
															goto l36
														}
													}
													t62 := int32(load32(m.memory[int64(uint32(v6))+128:]))
													v30 = t62
													m.fn790(v6+i32(116), v37, v29, v3, i32(0x1010000))
													if v30 == 0 {
														goto l24
													}
													m.fn17(v37, v30, i32(1))
												}
											l24:
												m.fn791(v6+i32(148), v28, v1)
												store64(m.memory[int64(uint32(v6))+136:], uint64(v13))
												store64(m.memory[int64(uint32(v6))+128:], uint64(v16))
												_ = m.fn45(v6+i32(80), i32(1078840), i32(1066714), v6+i32(128))
												{
													{
														t64 := int32(load32(m.memory[int64(uint32(v6))+148:]))
														v30 = t64
														if v30 == 0 {
															goto l37
														}
														t65 := int32(load32(m.memory[int64(uint32(v6))+152:]))
														v31 = t65
														t66 := int32(load32(m.memory[uint32(v31+i32(-4)):]))
														v29 = t66
														v34 = v29 & i32(-8)
														t67 := v34
														v29 = v29 & i32(3)
														p68 := i32(8)
														if v29 != 0 {
															p68 = i32(4)
														}
														if uint32(t67) < uint32(p68+v30) {
															m.fn3(i32(1274224), i32(46), i32(1274272))
															panic("unreachable")
														}
														if v29 == 0 {
															goto l39
														}
														if uint32(v34) > uint32(v30+i32(39)) {
															m.fn3(i32(1274288), i32(46), i32(1274336))
															panic("unreachable")
														}
													l39:
														m.fn1(v31)
													}
												l37:
													t69 := int32(load32(m.memory[int64(uint32(v6))+116:]))
													v30 = t69
													if v30 == 0 {
														goto l41
													}
													t70 := int32(load32(m.memory[int64(uint32(v6))+120:]))
													v31 = t70
													t71 := int32(load32(m.memory[uint32(v31+i32(-4)):]))
													v29 = t71
													v34 = v29 & i32(-8)
													t72 := v34
													v29 = v29 & i32(3)
													p73 := i32(8)
													if v29 != 0 {
														p73 = i32(4)
													}
													if uint32(t72) < uint32(p73+v30) {
														m.fn3(i32(1274224), i32(46), i32(1274272))
														panic("unreachable")
													}
													if v29 == 0 {
														goto l43
													}
													if uint32(v34) > uint32(v30+i32(39)) {
														m.fn3(i32(1274288), i32(46), i32(1274336))
														panic("unreachable")
													}
												l43:
													m.fn1(v31)
													goto l41
												}
											}
											if v22 == 0 {
												goto l12
											}
											t34 := m.fn249(v21, v20, v30, v1)
											t35 := v19
											v11 = t34
											v28 = t35 & int32(v11)
											v12 = int64(uint64(v11)>>25) & i64(127) * i64(72340172838076673)
											v32 = i32(0)
										l17:
											{
												{
													t36 := int64(load64(m.memory[uint32(v18+v28):]))
													v33 = t36
													v11 = v33 ^ v12
													v11 = (v11 ^ i64(-1)) & (v11 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
													if v11 == 0 {
														goto l13
													}
												l16:
													{
														t37 := v1
														v34 = v18 + (i32(0)-(int32(uint32(int64(bits.TrailingZeros64(uint64(v11))))>>3)+v28)&v19)*i32(28)
														t38 := int32(load32(m.memory[uint32(v34+i32(-20)):]))
														if t37 != t38 {
															goto l14
														}
														t39 := int32(load32(m.memory[uint32(v34+i32(-24)):]))
														t40 := m.fn973(v30, t39, v1)
														if t40 == 0 {
															t151 := int64(load64(m.memory[uint32(v34+i32(-12)):]))
															store64(m.memory[int64(uint32(v6))+148:], uint64(t151))
															store64(m.memory[int64(uint32(v6))+128:], uint64(v15))
															m.fn13(v6+i32(104), i32(1048821), v6+i32(128))
															t152 := int32(load32(m.memory[int64(uint32(v6))+96:]))
															t153 := int32(load32(m.memory[int64(uint32(v6))+100:]))
															m.fn143(v6+i32(16), t152, t153)
															{
																t154 := int32(load32(m.memory[int64(uint32(v6))+20:]))
																if t154 == 0 {
																	t157 := int32(load32(m.memory[int64(uint32(v6))+104:]))
																	v1 = t157
																	if v1 == 0 {
																		goto l74
																	}
																	{
																		t158 := int32(load32(m.memory[int64(uint32(v6))+108:]))
																		v30 = t158
																		t159 := int32(load32(m.memory[uint32(v30+i32(-4)):]))
																		v28 = t159
																		v29 = v28 & i32(-8)
																		t160 := v29
																		v28 = v28 & i32(3)
																		p161 := i32(8)
																		if v28 != 0 {
																			p161 = i32(4)
																		}
																		if uint32(t160) < uint32(p161+v1) {
																			m.fn3(i32(1274224), i32(46), i32(1274272))
																			panic("unreachable")
																		}
																		if v28 == 0 {
																			goto l80
																		}
																		if uint32(v29) > uint32(v1+i32(39)) {
																			m.fn3(i32(1274288), i32(46), i32(1274336))
																			panic("unreachable")
																		}
																	l80:
																		m.fn1(v30)
																		goto l74
																	}
																}
																t155 := int32(load32(m.memory[int64(uint32(v6))+112:]))
																v1 = t155
																t156 := int32(load32(m.memory[int64(uint32(v6))+108:]))
																v28 = t156
																goto l21
															}
														}
													}
												l14:
													v11 = (v11 + i64(-1)) & v11
													if !(v11 == 0) {
														goto l16
													}
												}
											l13:
												if !(v33&(v33<<1)&i64(-0x7f7f7f7f7f7f7f80) == 0) {
													goto l12
												}
												t41 := v28
												v32 = v32 + i32(8)
												v28 = (t41 + v32) & v19
												goto l17
											}
										}
									case 3:
										t74 := int32(load32(m.memory[int64(uint32(v2))+4:]))
										t75 := int32(load32(m.memory[int64(uint32(v2))+8:]))
										m.fn792(v6+i32(56), v10, t74, t75)
										t76 := int32(load32(m.memory[int64(uint32(v6))+56:]))
										v1 = t76
										if v1 == 0 {
											goto l45
										}
										t77 := int32(load32(m.memory[int64(uint32(v6))+60:]))
										v28 = t77
										store32(m.memory[int64(uint32(v6))+148:], uint32(v1))
										store32(m.memory[int64(uint32(v6))+152:], uint32(v28))
										store64(m.memory[int64(uint32(v6))+128:], uint64(v15))
										_ = m.fn45(v6+i32(80), i32(1078840), i32(1066458), v6+i32(128))
										goto l45
									}
								}
							l9:
								if v30 != 0 {
									goto l8
								}
								t79 := int32(m.memory[int64(uint32(v29))+12])
								if t79 != 0 {
									goto l10
								}
								t80 := int32(m.memory[int64(uint32(v29))+13])
								if t80 != 0 {
									goto l10
								}
								t81 := int32(m.memory[int64(uint32(v29))+14])
								if t81 != 0 {
									goto l10
								}
								t82 := int32(m.memory[int64(uint32(v29))+15])
								if t82 == i32(1) {
									goto l10
								}
								goto l8
							}
						l12:
							m.fn780(v6+i32(128), v29, v31, v3, i32(0), v5)
							t140 := int32(load32(m.memory[int64(uint32(v6))+132:]))
							v28 = t140
							{
								{
									t141 := int32(load32(m.memory[int64(uint32(v6))+136:]))
									v1 = t141
									t142 := int32(load32(m.memory[int64(uint32(v6))+80:]))
									t143 := int32(load32(m.memory[int64(uint32(v6))+88:]))
									t144 := v1
									v30 = t143
									if uint32(t144) <= uint32(t142-v30) {
										goto l71
									}
									m.fn196(v6+i32(80), v30, v1, i32(1), i32(1))
									t145 := int32(load32(m.memory[int64(uint32(v6))+88:]))
									v30 = t145
									goto l72
								}
							l71:
								if v1 == 0 {
									goto l73
								}
							l72:
								if v1 == 0 {
									goto l73
								}
								t146 := int32(load32(m.memory[int64(uint32(v6))+84:]))
								memory_copy(m.memory, uint32(t146+v30), uint32(v28), uint32(v1))
							}
						l73:
							store32(m.memory[int64(uint32(v6))+88:], uint32(v30+v1))
							t147 := int32(load32(m.memory[int64(uint32(v6))+128:]))
							v1 = t147
							if v1 == 0 {
								goto l74
							}
							{
								t148 := int32(load32(m.memory[uint32(v28+i32(-4)):]))
								v30 = t148
								v29 = v30 & i32(-8)
								t149 := v29
								v30 = v30 & i32(3)
								p150 := i32(8)
								if v30 != 0 {
									p150 = i32(4)
								}
								if uint32(t149) < uint32(p150+v1) {
									m.fn3(i32(1274224), i32(46), i32(1274272))
									panic("unreachable")
								}
								if v30 == 0 {
									goto l76
								}
								if uint32(v29) > uint32(v1+i32(39)) {
									m.fn3(i32(1274288), i32(46), i32(1274336))
									panic("unreachable")
								}
							l76:
								m.fn1(v28)
								goto l74
							}
						}
					l74:
						t162 := int32(load32(m.memory[int64(uint32(v6))+92:]))
						v1 = t162
						if v1 == 0 {
							goto l45
						}
						{
							t163 := int32(load32(m.memory[int64(uint32(v6))+96:]))
							v30 = t163
							t164 := int32(load32(m.memory[uint32(v30+i32(-4)):]))
							v28 = t164
							v29 = v28 & i32(-8)
							t165 := v29
							v28 = v28 & i32(3)
							p166 := i32(8)
							if v28 != 0 {
								p166 = i32(4)
							}
							if uint32(t165) < uint32(p166+v1) {
								m.fn3(i32(1274224), i32(46), i32(1274272))
								panic("unreachable")
							}
							if v28 == 0 {
								goto l83
							}
							if uint32(v29) > uint32(v1+i32(39)) {
								m.fn3(i32(1274288), i32(46), i32(1274336))
								panic("unreachable")
							}
						l83:
							m.fn1(v30)
							goto l45
						}
					}
				l21:
					m.fn791(v6+i32(148), v28, v1)
					store64(m.memory[int64(uint32(v6))+136:], uint64(v13))
					store64(m.memory[int64(uint32(v6))+128:], uint64(v14))
					_ = m.fn45(v6+i32(80), i32(1078840), i32(1066714), v6+i32(128))
					{
						t168 := int32(load32(m.memory[int64(uint32(v6))+148:]))
						v1 = t168
						if v1 == 0 {
							goto l85
						}
						t169 := int32(load32(m.memory[int64(uint32(v6))+152:]))
						v29 = t169
						t170 := int32(load32(m.memory[uint32(v29+i32(-4)):]))
						v30 = t170
						v31 = v30 & i32(-8)
						t171 := v31
						v30 = v30 & i32(3)
						p172 := i32(8)
						if v30 != 0 {
							p172 = i32(4)
						}
						if uint32(t171) < uint32(p172+v1) {
							m.fn3(i32(1274224), i32(46), i32(1274272))
							panic("unreachable")
						}
						if v30 == 0 {
							goto l87
						}
						if uint32(v31) > uint32(v1+i32(39)) {
							m.fn3(i32(1274288), i32(46), i32(1274336))
							panic("unreachable")
						}
					l87:
						m.fn1(v29)
					}
				l85:
					t173 := int32(load32(m.memory[int64(uint32(v6))+104:]))
					v1 = t173
				}
			l41:
				{
					if v1 == 0 {
						goto l89
					}
					t174 := int32(load32(m.memory[uint32(v28+i32(-4)):]))
					v30 = t174
					v29 = v30 & i32(-8)
					t175 := v29
					v30 = v30 & i32(3)
					p176 := i32(8)
					if v30 != 0 {
						p176 = i32(4)
					}
					if uint32(t175) < uint32(p176+v1) {
						m.fn3(i32(1274224), i32(46), i32(1274272))
						panic("unreachable")
					}
					if v30 == 0 {
						goto l91
					}
					if uint32(v29) > uint32(v1+i32(39)) {
						m.fn3(i32(1274288), i32(46), i32(1274336))
						panic("unreachable")
					}
				l91:
					m.fn1(v28)
				}
			l89:
				t177 := int32(load32(m.memory[int64(uint32(v6))+92:]))
				v1 = t177
				if v1 == 0 {
					goto l45
				}
				t178 := int32(load32(m.memory[int64(uint32(v6))+96:]))
				v30 = t178
				t179 := int32(load32(m.memory[uint32(v30+i32(-4)):]))
				v28 = t179
				v29 = v28 & i32(-8)
				t180 := v29
				v28 = v28 & i32(3)
				p181 := i32(8)
				if v28 != 0 {
					p181 = i32(4)
				}
				if uint32(t180) < uint32(p181+v1) {
					m.fn3(i32(1274224), i32(46), i32(1274272))
					panic("unreachable")
				}
				if v28 == 0 {
					goto l94
				}
				if uint32(v29) > uint32(v1+i32(39)) {
					m.fn3(i32(1274288), i32(46), i32(1274336))
					panic("unreachable")
				}
			l94:
				m.fn1(v30)
				goto l45
			}
		l10:
			v28 = i32(65536)
		l8:
			t182 := int32(load32(m.memory[int64(uint32(v2))+8:]))
			v1 = t182
			t183 := int32(load32(m.memory[int64(uint32(v2))+4:]))
			v30 = t183
			{
				{
					{
						{
							t184 := int32(load32(m.memory[int64(uint32(v2))+12:]))
							v34 = t184
							if v34&i32(16843009) == 0 {
								t187 := int32(load32(m.memory[int64(uint32(v6))+88:]))
								v29 = t187
								if v29 != 0 {
									goto l102
								}
								v31 = i32(1)
								goto l103
							}
							m.fn209(v6+i32(8), v30, v1)
							t185 := int32(load32(m.memory[int64(uint32(v6))+12:]))
							v31 = t185
							m.fn693(v6, v30, v1)
							v29 = v1 - v31
							t186 := int32(load32(m.memory[int64(uint32(v6))+4:]))
							v28 = t186
							if v1 == v31 {
								if v28 == 0 {
									goto l104
								}
								goto l100
							}
							if uint32(v29) < uint32(v1) {
								t188 := int32(int8(m.memory[uint32(v30+v29)]))
								if t188 <= i32(-65) {
									goto l99
								}
								if uint32(v28) >= uint32(v29) {
									goto l100
								}
								goto l101
							}
							if v31 != 0 {
								goto l99
							}
							if uint32(v28) >= uint32(v29) {
								goto l100
							}
							goto l101
						}
					l102:
						t189 := int32(load32(m.memory[int64(uint32(v6))+84:]))
						t190 := int32(m.memory[uint32(t189+v29+i32(-1))])
						var p191 int32
						if t190 == i32(10) {
							p191 = 1
						}
						v31 = p191
					}
				l103:
					m.fn790(v6+i32(128), v30, v1, v3, v28|v31|v9)
					t192 := int32(load32(m.memory[int64(uint32(v6))+132:]))
					v28 = t192
					{
						{
							t193 := int32(load32(m.memory[int64(uint32(v6))+136:]))
							v1 = t193
							t194 := int32(load32(m.memory[int64(uint32(v6))+80:]))
							if uint32(v1) <= uint32(t194-v29) {
								goto l105
							}
							m.fn196(v6+i32(80), v29, v1, i32(1), i32(1))
							t195 := int32(load32(m.memory[int64(uint32(v6))+88:]))
							v29 = t195
							goto l106
						}
					l105:
						if v1 == 0 {
							goto l107
						}
					l106:
						if v1 == 0 {
							goto l107
						}
						t196 := int32(load32(m.memory[int64(uint32(v6))+84:]))
						memory_copy(m.memory, uint32(t196+v29), uint32(v28), uint32(v1))
					}
				l107:
					store32(m.memory[int64(uint32(v6))+88:], uint32(v29+v1))
					t197 := int32(load32(m.memory[int64(uint32(v6))+128:]))
					v1 = t197
					if v1 == 0 {
						goto l45
					}
					t198 := int32(load32(m.memory[uint32(v28+i32(-4)):]))
					v30 = t198
					v29 = v30 & i32(-8)
					t199 := v29
					v30 = v30 & i32(3)
					p200 := i32(8)
					if v30 != 0 {
						p200 = i32(4)
					}
					if uint32(t199) < uint32(p200+v1) {
						m.fn3(i32(1274224), i32(46), i32(1274272))
						panic("unreachable")
					}
					if v30 == 0 {
						goto l109
					}
					if uint32(v29) > uint32(v1+i32(39)) {
						m.fn3(i32(1274288), i32(46), i32(1274336))
						panic("unreachable")
					}
				l109:
					m.fn1(v28)
					goto l45
				}
			l99:
				m.fn37(v30, v1, i32(0), v29, i32(1078876))
				panic("unreachable")
			l100:
				{
					if uint32(v1) > uint32(v28) {
						goto l111
					}
					if v1 != v28 {
						goto l101
					}
					goto l112
				l111:
					t201 := int32(int8(m.memory[uint32(v30+v28)]))
					if t201 <= i32(-65) {
						goto l101
					}
				}
			l112:
				if v1 == v31 {
					goto l113
				}
				{
					t202 := int32(load32(m.memory[int64(uint32(v6))+80:]))
					t203 := int32(load32(m.memory[int64(uint32(v6))+88:]))
					t204 := v29
					v31 = t203
					if uint32(t204) <= uint32(t202-v31) {
						goto l114
					}
					m.fn196(v6+i32(80), v31, v29, i32(1), i32(1))
					t205 := int32(load32(m.memory[int64(uint32(v6))+88:]))
					v31 = t205
				}
			l114:
				{
					if v29 == 0 {
						goto l115
					}
					t206 := int32(load32(m.memory[int64(uint32(v6))+84:]))
					memory_copy(m.memory, uint32(t206+v31), uint32(v30), uint32(v29))
				}
			l115:
				store32(m.memory[int64(uint32(v6))+88:], uint32(v31+v29))
			l113:
				if v28 == v29 {
					goto l104
				}
				v40 = v30 + v29
				v41 = v28 - v29
				{
					if v34&i32(0x1000000) != 0 {
						m.fn793(v40, v41, v6+i32(80))
						goto l104
					}
					v32 = i32(0)
					store32(m.memory[int64(uint32(v6))+156:], uint32(i32(0)))
					store64(m.memory[int64(uint32(v6))+148:], uint64(i64(0x100000000)))
					v31 = i32(1)
					if v34&i32(65536) != 0 {
						goto l117
					}
					goto l118
				l117:
					m.fn196(v6+i32(148), i32(0), i32(2), i32(1), i32(1))
					t207 := int32(load32(m.memory[int64(uint32(v6))+152:]))
					v31 = t207
					t208 := int32(load32(m.memory[int64(uint32(v6))+156:]))
					t209 := v31
					v29 = t208
					store16(m.memory[uint32(t209+v29):], uint16(i32(32382)))
					t210 := v6
					v32 = v29 + i32(2)
					store32(m.memory[int64(uint32(t210))+156:], uint32(v32))
				}
			l118:
				v29 = v34 & i32(256)
				{
					if v34&i32(1) == 0 {
						goto l119
					}
					{
						t211 := int32(load32(m.memory[int64(uint32(v6))+148:]))
						if uint32(t211-v32) > uint32(i32(1)) {
							goto l120
						}
						m.fn196(v6+i32(148), v32, i32(2), i32(1), i32(1))
						t212 := int32(load32(m.memory[int64(uint32(v6))+152:]))
						v31 = t212
						t213 := int32(load32(m.memory[int64(uint32(v6))+156:]))
						v32 = t213
					}
				l120:
					store16(m.memory[uint32(v31+v32):], uint16(i32(10794)))
					t214 := v6
					v32 = v32 + i32(2)
					store32(m.memory[int64(uint32(t214))+156:], uint32(v32))
				}
			l119:
				{
					{
						if v29 != 0 {
							{
								t216 := int32(load32(m.memory[int64(uint32(v6))+148:]))
								if t216 != v32 {
									goto l124
								}
								m.fn196(v6+i32(148), v32, i32(1), i32(1), i32(1))
							}
						l124:
							t217 := int32(load32(m.memory[int64(uint32(v6))+152:]))
							v38 = t217
							m.memory[uint32(v38+v32)] = byte(i32(42))
							t218 := v6
							v32 = v32 + i32(1)
							store32(m.memory[int64(uint32(t218))+156:], uint32(v32))
							store32(m.memory[int64(uint32(v6))+136:], uint32(i32(0)))
							store64(m.memory[int64(uint32(v6))+128:], uint64(i64(0x100000000)))
							goto l123
						}
						t215 := int32(load32(m.memory[int64(uint32(v6))+152:]))
						v38 = t215
						store32(m.memory[int64(uint32(v6))+136:], uint32(i32(0)))
						store64(m.memory[int64(uint32(v6))+128:], uint64(i64(0x100000000)))
						if v32 == 0 {
							v35 = i32(1)
							v32 = i32(0)
							t219 := int32(load32(m.memory[int64(uint32(v6))+80:]))
							v37 = t219
							t220 := int32(load32(m.memory[int64(uint32(v6))+88:]))
							v31 = t220
							v36 = i32(0)
							v29 = i32(0)
							goto l125
						}
						goto l123
					}
				l123:
					v34 = v38 + v32
					t221 := v6 + i32(128)
					t222 := int32(uint32(v32) >> 2)
					var p223 int32
					if v32&i32(3) != i32(0) {
						p223 = 1
					}
					m.fn196(t221, i32(0), t222+p223, i32(1), i32(1))
				l139:
					{
						{
							{
								v36 = v34 + i32(-1)
								t224 := int32(int8(m.memory[uint32(v36)]))
								v29 = t224
								if v29 > i32(-1) {
									v39 = i32(1)
									t226 := int32(load32(m.memory[int64(uint32(v6))+136:]))
									v31 = t226
									v34 = v36
									v36 = i32(1)
									goto l129
								}
								v37 = v34 + i32(-2)
								t225 := int32(m.memory[uint32(v37)])
								v31 = t225
								v36 = int32(int8(v31))
								if v36 < i32(-64) {
									goto l127
								}
								v34 = v31 & i32(31)
								goto l128
							}
						l127:
							{
								{
									v37 = v34 + i32(-3)
									t227 := int32(m.memory[uint32(v37)])
									v31 = t227
									v35 = int32(int8(v31))
									if v35 <= i32(-65) {
										goto l130
									}
									v31 = v31 & i32(15)
									goto l131
								}
							l130:
								v37 = v34 + i32(-4)
								t228 := int32(m.memory[uint32(v37)])
								v31 = t228&i32(7)<<6 | v35&i32(63)
							}
						l131:
							v34 = v31<<6 | v36&i32(63)
						l128:
							v29 = v34<<6 | v29&i32(63)
							v39 = i32(1)
							t229 := int32(load32(m.memory[int64(uint32(v6))+136:]))
							v31 = t229
							if uint32(v34) >= uint32(i32(2)) {
								goto l132
							}
							v34 = v37
							v36 = i32(1)
							goto l129
						l132:
							v36 = i32(2)
							v39 = i32(0)
							{
								if uint32(v34) < uint32(i32(32)) {
									goto l133
								}
								p230 := i32(4)
								if uint32(v34) < uint32(i32(1024)) {
									p230 = i32(3)
								}
								v36 = p230
							}
						l133:
							v34 = v37
						}
					l129:
						{
							t231 := int32(load32(m.memory[int64(uint32(v6))+128:]))
							if uint32(v36) <= uint32(t231-v31) {
								goto l134
							}
							m.fn196(v6+i32(128), v31, v36, i32(1), i32(1))
						}
					l134:
						t232 := int32(load32(m.memory[int64(uint32(v6))+132:]))
						v35 = t232
						v37 = v35 + v31
						if v39 != 0 {
							goto l135
						}
						v39 = v29&i32(63) | i32(-128)
						v42 = int32(uint32(v29) >> 6)
						if uint32(v29) >= uint32(i32(2048)) {
							v43 = int32(uint32(v29) >> 12)
							v42 = v42&i32(63) | i32(-128)
							if uint32(v29) > uint32(i32(0xffff)) {
								m.memory[int64(uint32(v37))+3] = byte(v39)
								m.memory[int64(uint32(v37))+2] = byte(v42)
								m.memory[int64(uint32(v37))+1] = byte(v43&i32(63) | i32(-128))
								m.memory[uint32(v37)] = byte(int32(uint32(v29)>>18) | i32(-16))
								goto l137
							}
							m.memory[int64(uint32(v37))+2] = byte(v39)
							m.memory[int64(uint32(v37))+1] = byte(v42)
							m.memory[uint32(v37)] = byte(v43 | i32(224))
							goto l137
						}
						m.memory[int64(uint32(v37))+1] = byte(v39)
						m.memory[uint32(v37)] = byte(v42 | i32(192))
						goto l137
					l135:
						m.memory[uint32(v37)] = byte(v29)
					l137:
						t233 := v6
						v29 = v36 + v31
						store32(m.memory[int64(uint32(t233))+136:], uint32(v29))
						if v38 != v34 {
							goto l139
						}
					}
					t234 := int32(load32(m.memory[int64(uint32(v6))+128:]))
					v36 = t234
					{
						t235 := int32(load32(m.memory[int64(uint32(v6))+80:]))
						t236 := v32
						v37 = t235
						t237 := int32(load32(m.memory[int64(uint32(v6))+88:]))
						t238 := v37
						v31 = t237
						if uint32(t236) <= uint32(t238-v31) {
							goto l140
						}
						m.fn196(v6+i32(80), v31, v32, i32(1), i32(1))
						t239 := int32(load32(m.memory[int64(uint32(v6))+80:]))
						v37 = t239
						t240 := int32(load32(m.memory[int64(uint32(v6))+88:]))
						v31 = t240
					}
				l140:
					if v32 == 0 {
						goto l125
					}
					t241 := int32(load32(m.memory[int64(uint32(v6))+84:]))
					memory_copy(m.memory, uint32(t241+v31), uint32(v38), uint32(v32))
				}
			l125:
				t242 := v6
				v34 = v31 + v32
				store32(m.memory[int64(uint32(t242))+88:], uint32(v34))
				m.fn790(v6+i32(128), v40, v41, v3, v8)
				t243 := int32(load32(m.memory[int64(uint32(v6))+132:]))
				v32 = t243
				{
					{
						t244 := int32(load32(m.memory[int64(uint32(v6))+136:]))
						v31 = t244
						if uint32(v31) <= uint32(v37-v34) {
							goto l141
						}
						m.fn196(v6+i32(80), v34, v31, i32(1), i32(1))
						t245 := int32(load32(m.memory[int64(uint32(v6))+88:]))
						v34 = t245
						goto l142
					}
				l141:
					if v31 == 0 {
						goto l143
					}
				l142:
					if v31 == 0 {
						goto l143
					}
					t246 := int32(load32(m.memory[int64(uint32(v6))+84:]))
					memory_copy(m.memory, uint32(t246+v34), uint32(v32), uint32(v31))
				}
			l143:
				t247 := v6
				v31 = v34 + v31
				store32(m.memory[int64(uint32(t247))+88:], uint32(v31))
				{
					{
						t248 := int32(load32(m.memory[int64(uint32(v6))+128:]))
						v34 = t248
						if v34 == 0 {
							goto l144
						}
						t249 := int32(load32(m.memory[uint32(v32+i32(-4)):]))
						v37 = t249
						v39 = v37 & i32(-8)
						t250 := v39
						v37 = v37 & i32(3)
						p251 := i32(8)
						if v37 != 0 {
							p251 = i32(4)
						}
						if uint32(t250) < uint32(p251+v34) {
							m.fn3(i32(1274224), i32(46), i32(1274272))
							panic("unreachable")
						}
						if v37 == 0 {
							goto l146
						}
						if uint32(v39) > uint32(v34+i32(39)) {
							m.fn3(i32(1274288), i32(46), i32(1274336))
							panic("unreachable")
						}
					l146:
						m.fn1(v32)
					}
				l144:
					{
						{
							t252 := int32(load32(m.memory[int64(uint32(v6))+80:]))
							if uint32(v29) <= uint32(t252-v31) {
								goto l148
							}
							m.fn196(v6+i32(80), v31, v29, i32(1), i32(1))
							t253 := int32(load32(m.memory[int64(uint32(v6))+88:]))
							v31 = t253
							goto l149
						}
					l148:
						if v29 == 0 {
							goto l150
						}
					l149:
						if v29 == 0 {
							goto l150
						}
						t254 := int32(load32(m.memory[int64(uint32(v6))+84:]))
						memory_copy(m.memory, uint32(t254+v31), uint32(v35), uint32(v29))
					}
				l150:
					store32(m.memory[int64(uint32(v6))+88:], uint32(v31+v29))
					{
						if v36 == 0 {
							goto l151
						}
						t255 := int32(load32(m.memory[uint32(v35+i32(-4)):]))
						v29 = t255
						v31 = v29 & i32(-8)
						t256 := v31
						v29 = v29 & i32(3)
						p257 := i32(8)
						if v29 != 0 {
							p257 = i32(4)
						}
						if uint32(t256) < uint32(p257+v36) {
							m.fn3(i32(1274224), i32(46), i32(1274272))
							panic("unreachable")
						}
						if v29 == 0 {
							goto l153
						}
						if uint32(v31) > uint32(v36+i32(39)) {
							m.fn3(i32(1274288), i32(46), i32(1274336))
							panic("unreachable")
						}
					l153:
						m.fn1(v35)
					}
				l151:
					t258 := int32(load32(m.memory[int64(uint32(v6))+148:]))
					v29 = t258
					if v29 == 0 {
						goto l104
					}
					t259 := int32(load32(m.memory[uint32(v38+i32(-4)):]))
					v31 = t259
					v34 = v31 & i32(-8)
					t260 := v34
					v31 = v31 & i32(3)
					p261 := i32(8)
					if v31 != 0 {
						p261 = i32(4)
					}
					if uint32(t260) < uint32(p261+v29) {
						m.fn3(i32(1274224), i32(46), i32(1274272))
						panic("unreachable")
					}
					if v31 == 0 {
						goto l156
					}
					if uint32(v34) > uint32(v29+i32(39)) {
						m.fn3(i32(1274288), i32(46), i32(1274336))
						panic("unreachable")
					}
				l156:
					m.fn1(v38)
					goto l104
				}
			}
		l104:
			if v1 == v28 {
				goto l45
			}
			{
				v1 = v1 - v28
				t262 := int32(load32(m.memory[int64(uint32(v6))+80:]))
				t263 := int32(load32(m.memory[int64(uint32(v6))+88:]))
				t264 := v1
				v29 = t263
				if uint32(t264) <= uint32(t262-v29) {
					goto l158
				}
				m.fn196(v6+i32(80), v29, v1, i32(1), i32(1))
				t265 := int32(load32(m.memory[int64(uint32(v6))+88:]))
				v29 = t265
			}
		l158:
			{
				if v1 == 0 {
					goto l159
				}
				t266 := int32(load32(m.memory[int64(uint32(v6))+84:]))
				memory_copy(m.memory, uint32(t266+v29), uint32(v30+v28), uint32(v1))
			}
		l159:
			store32(m.memory[int64(uint32(v6))+88:], uint32(v29+v1))
		}
	l45:
		v2 = v2 + i32(16)
		if v2 == v7 {
			goto l160
		}
		goto l161
	l101:
		m.fn37(v30, v1, v29, v28, i32(1078892))
		panic("unreachable")
	l160:
		t267 := int64(load64(m.memory[int64(uint32(v6))+80:]))
		store64(m.memory[uint32(v0):], uint64(t267))
		t268 := int32(load32(m.memory[int64(uint32(v6))+88:]))
		store32(m.memory[int64(uint32(v0))+8:], uint32(t268))
		t269 := int32(load32(m.memory[int64(uint32(v6))+72:]))
		v2 = t269
		t270 := int32(load32(m.memory[int64(uint32(v6))+76:]))
		v1 = t270
		if v1 == 0 {
			goto l1
		}
		v4 = v2
	l166:
		{
			t271 := int32(load32(m.memory[uint32(v4):]))
			v7 = t271
			if v7 < i32(-0x7ffffffb) {
				goto l162
			}
			if uint32(v7+i32(1)) < uint32(i32(2)) {
				goto l162
			}
			t272 := int32(load32(m.memory[uint32(v4+i32(4)):]))
			v5 = t272
			t273 := int32(load32(m.memory[uint32(v5+i32(-4)):]))
			v28 = t273
			v3 = v28 & i32(-8)
			t274 := v3
			v28 = v28 & i32(3)
			p275 := i32(8)
			if v28 != 0 {
				p275 = i32(4)
			}
			if uint32(t274) < uint32(p275+v7) {
				m.fn3(i32(1274224), i32(46), i32(1274272))
				panic("unreachable")
			}
			if v28 == 0 {
				goto l164
			}
			if uint32(v3) > uint32(v7+i32(39)) {
				m.fn3(i32(1274288), i32(46), i32(1274336))
				panic("unreachable")
			}
		l164:
			m.fn1(v5)
		}
	l162:
		v4 = v4 + i32(16)
		v1 = v1 + i32(-1)
		if v1 != 0 {
			goto l166
		}
	}
l1:
	{
		t276 := int32(load32(m.memory[int64(uint32(v6))+68:]))
		v4 = t276
		if v4 == 0 {
			goto l167
		}
		t277 := int32(load32(m.memory[uint32(v2+i32(-4)):]))
		v1 = t277
		v7 = v1 & i32(-8)
		t278 := v7
		v1 = v1 & i32(3)
		p279 := i32(8)
		if v1 != 0 {
			p279 = i32(4)
		}
		v4 = v4 << 4
		if uint32(t278) < uint32(p279|v4) {
			m.fn3(i32(1274224), i32(46), i32(1274272))
			panic("unreachable")
		}
		if v1 == 0 {
			goto l169
		}
		if uint32(v7) > uint32(v4+i32(39)) {
			m.fn3(i32(1274288), i32(46), i32(1274336))
			panic("unreachable")
		}
	l169:
		m.fn1(v2)
	}
l167:
	m.g0 = v6 + i32(160)
}
func (m *Module) fn781(v0, v1 int32) {
	var v2, v3, v4, v5, v6, v7, v8, v9, v10 int32
	t0 := m.g0
	v2 = t0 - i32(48)
	m.g0 = v2
	v3 = i32(0)
	{
		{
			t1 := int32(m.memory[int64(uint32(v1))+37])
			if t1 == 0 {
				goto l0
			}
			goto l1
		}
	l0:
		t2 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		v4 = t2
		m.fn198(v2+i32(36), v1)
		{
			{
				t3 := int32(load32(m.memory[int64(uint32(v2))+36:]))
				if t3 != i32(1) {
					goto l2
				}
				t4 := int32(load32(m.memory[int64(uint32(v1))+28:]))
				v5 = t4
				t5 := int32(load32(m.memory[int64(uint32(v2))+44:]))
				t6 := v1
				v6 = t5
				store32(m.memory[int64(uint32(t6))+28:], uint32(v6))
				v1 = v4 + v5
				v4 = v6 - v5
				goto l3
			}
		l2:
			t7 := int32(m.memory[int64(uint32(v1))+37])
			if t7 != 0 {
				goto l4
			}
			m.memory[int64(uint32(v1))+37] = byte(i32(1))
			{
				{
					t8 := int32(m.memory[int64(uint32(v1))+36])
					if t8 != i32(1) {
						goto l5
					}
					t9 := int32(load32(m.memory[int64(uint32(v1))+32:]))
					v4 = t9
					t10 := int32(load32(m.memory[int64(uint32(v1))+28:]))
					v5 = t10
					goto l6
				}
			l5:
				t11 := int32(load32(m.memory[int64(uint32(v1))+32:]))
				v4 = t11
				t12 := int32(load32(m.memory[int64(uint32(v1))+28:]))
				t13 := v4
				v5 = t12
				if t13 == v5 {
					goto l4
				}
			}
		l6:
			v4 = v4 - v5
			t14 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			v1 = t14 + v5
		}
	l3:
		{
			if v4 == 0 {
				goto l7
			}
			t15 := v1
			v5 = v4 + i32(-1)
			t16 := int32(m.memory[uint32(t15+v5)])
			if t16 != i32(10) {
				goto l7
			}
			v4 = v4 + i32(-2)
			{
				if v5 != 0 {
					goto l8
				}
				v6 = i32(0)
				goto l9
			l8:
				t17 := int32(m.memory[uint32(v1+v4)])
				p18 := i32(0)
				if t17&i32(255) == i32(13) {
					p18 = v1
				}
				v6 = p18
			}
		l9:
			p19 := v5
			if v6 != 0 {
				p19 = v4
			}
			v4 = p19
			p20 := v1
			if v6 != 0 {
				p20 = v6
			}
			v1 = p20
		}
	l7:
		if v1 != 0 {
			m.fn209(v2+i32(24), v1, v4)
			t21 := int32(load32(m.memory[int64(uint32(v2))+24:]))
			v1 = t21
			{
				{
					t22 := int32(load32(m.memory[int64(uint32(v2))+28:]))
					v7 = t22
					if v7 != 0 {
						goto l11
					}
					v7 = i32(0)
					goto l12
				}
			l11:
				v5 = v1 + v7
				v6 = i32(0)
			l19:
				{
					v3 = v5 + i32(-1)
					t23 := int32(int8(m.memory[uint32(v3)]))
					v4 = t23
					if v4 > i32(-1) {
						goto l13
					}
					{
						v3 = v5 + i32(-2)
						t24 := int32(m.memory[uint32(v3)])
						v8 = t24
						v9 = int32(int8(v8))
						if v9 < i32(-64) {
							goto l14
						}
						v5 = v8 & i32(31)
						goto l15
					}
				l14:
					{
						{
							v3 = v5 + i32(-3)
							t25 := int32(m.memory[uint32(v3)])
							v8 = t25
							v10 = int32(int8(v8))
							if v10 < i32(-64) {
								goto l16
							}
							v5 = v8 & i32(15)
							goto l17
						}
					l16:
						v3 = v5 + i32(-4)
						t26 := int32(m.memory[uint32(v3)])
						v5 = t26&i32(7)<<6 | v10&i32(63)
					}
				l17:
					v5 = v5<<6 | v9&i32(63)
				l15:
					v4 = v5<<6 | v4&i32(63)
				}
			l13:
				if v4 != i32(92) {
					goto l18
				}
				v6 = v6 + i32(1)
				v5 = v3
				if v1 != v3 {
					goto l19
				}
			l18:
				if v6&i32(1) != 0 {
					goto l12
				}
				m.fn693(v2+i32(16), v1, v7)
				t27 := int32(load32(m.memory[int64(uint32(v2))+20:]))
				v7 = t27
				t28 := int32(load32(m.memory[int64(uint32(v2))+16:]))
				v1 = t28
			}
		l12:
			v4 = v7
		l27:
			v5 = v4
			if v5 != 0 {
				goto l20
			}
			v5 = i32(0)
			goto l21
		l20:
			{
				v6 = v1 + v5
				v4 = v6 + i32(-1)
				t29 := int32(int8(m.memory[uint32(v4)]))
				v3 = t29
				if v3 > i32(-1) {
					goto l22
				}
				{
					v4 = v6 + i32(-2)
					t30 := int32(m.memory[uint32(v4)])
					v8 = t30
					v9 = int32(int8(v8))
					if v9 < i32(-64) {
						goto l23
					}
					v6 = v8 & i32(31)
					goto l24
				}
			l23:
				{
					{
						v4 = v6 + i32(-3)
						t31 := int32(m.memory[uint32(v4)])
						v8 = t31
						v10 = int32(int8(v8))
						if v10 < i32(-64) {
							goto l25
						}
						v6 = v8 & i32(15)
						goto l26
					}
				l25:
					v4 = v6 + i32(-4)
					t32 := int32(m.memory[uint32(v4)])
					v6 = t32&i32(7)<<6 | v10&i32(63)
				}
			l26:
				v6 = v6<<6 | v9&i32(63)
			l24:
				v3 = v6<<6 | v3&i32(63)
			}
		l22:
			v4 = v4 - v1
			if v3 == i32(92) {
				goto l27
			}
		l21:
			m.fn143(v2+i32(8), v1, v5)
			t33 := int32(load32(m.memory[int64(uint32(v2))+12:]))
			t34 := v1
			v4 = t33
			p35 := i32(1)
			if v4 != 0 {
				p35 = t34
			}
			v3 = p35
			p36 := i32(0)
			if v4 != 0 {
				p36 = v7
			}
			v1 = p36
			goto l1
		}
		goto l1
	l4:
	}
l1:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
	store32(m.memory[uint32(v0):], uint32(v3))
	m.g0 = v2 + i32(48)
}
func (m *Module) fn782(v0, v1, v2, v3, v4 int32) {
	var v5, v6, v7, v8, v9, v10, v11, v12 int32
	t0 := m.g0
	v5 = t0 - i32(16)
	m.g0 = v5
	{
		if v2 == 0 {
			store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
			store64(m.memory[uint32(v0):], uint64(i64(0x100000000)))
			goto l32
		}
		{
			t1 := int32(load32(m.memory[uint32(v1+i32(4)):]))
			v6 = t1
			t2 := v6
			v7 = v2 << 3
			v8 = v7 + i32(-8)
			v9 = t2 + int32(uint32(v8)>>3)*v4
			if uint32(v9) < uint32(v6) {
				goto l1
			}
			v10 = v1 + v7
			t3 := int32(load32(m.memory[uint32(v1):]))
			v11 = t3
			v12 = v1 + i32(8)
			v1 = v12
		l3:
			{
				if v8 == 0 {
					if v9 <= i32(-1) {
						m.fn11()
						panic("unreachable")
					}
					{
						if v9 != 0 {
							goto l5
						}
						v8 = i32(1)
						goto l6
					l5:
						t5 := m.fn7(v9)
						v8 = t5
						if v8 == 0 {
							m.fn12(i32(1), v9)
							panic("unreachable")
						}
					}
				l6:
					v1 = i32(0)
					store32(m.memory[int64(uint32(v5))+12:], uint32(i32(0)))
					store32(m.memory[int64(uint32(v5))+8:], uint32(v8))
					store32(m.memory[int64(uint32(v5))+4:], uint32(v9))
					{
						if uint32(v6) <= uint32(v9) {
							goto l8
						}
						m.fn196(v5+i32(4), i32(0), v6, i32(1), i32(1))
						t6 := int32(load32(m.memory[int64(uint32(v5))+8:]))
						v8 = t6
						t7 := int32(load32(m.memory[int64(uint32(v5))+12:]))
						v1 = t7
						goto l9
					}
				l8:
					if v6 == 0 {
						goto l10
					}
				l9:
					if v6 == 0 {
						goto l10
					}
					memory_copy(m.memory, uint32(v8+v1), uint32(v11), uint32(v6))
				l10:
					t8 := v9
					v7 = v1 + v6
					v1 = t8 - v7
					v8 = v8 + v7
					switch v4 + i32(-1) {
					case 2:
						if v2 == i32(1) {
							goto l15
						}
					l19:
						{
							if uint32(v1) <= uint32(i32(2)) {
								m.fn27(i32(1272168), i32(19), i32(1069620))
								panic("unreachable")
							}
							t9 := int32(load32(m.memory[uint32(v12+i32(4)):]))
							v7 = t9
							t10 := int32(load32(m.memory[uint32(v12):]))
							v2 = t10
							t11 := int32(m.memory[int64(uint32(v3))+2])
							m.memory[int64(uint32(v8))+2] = byte(t11)
							t12 := int32(load16(m.memory[uint32(v3):]))
							store16(m.memory[uint32(v8):], uint16(t12))
							v1 = v1 + i32(-3)
							if uint32(v1) < uint32(v7) {
								m.fn27(i32(1272168), i32(19), i32(1069620))
								panic("unreachable")
							}
							v8 = v8 + i32(3)
							if v7 == 0 {
								goto l18
							}
							memory_copy(m.memory, uint32(v8), uint32(v2), uint32(v7))
						l18:
							v1 = v1 - v7
							v8 = v8 + v7
							v12 = v12 + i32(8)
							if v12 != v10 {
								goto l19
							}
							goto l15
						}
					case 1:
						if v2 == i32(1) {
							goto l15
						}
					l23:
						{
							if uint32(v1) <= uint32(i32(1)) {
								m.fn27(i32(1272168), i32(19), i32(1069620))
								panic("unreachable")
							}
							t13 := int32(load32(m.memory[uint32(v12+i32(4)):]))
							v7 = t13
							t14 := int32(load32(m.memory[uint32(v12):]))
							v2 = t14
							t15 := int32(load16(m.memory[uint32(v3):]))
							store16(m.memory[uint32(v8):], uint16(t15))
							v1 = v1 + i32(-2)
							if uint32(v1) < uint32(v7) {
								m.fn27(i32(1272168), i32(19), i32(1069620))
								panic("unreachable")
							}
							v8 = v8 + i32(2)
							if v7 == 0 {
								goto l22
							}
							memory_copy(m.memory, uint32(v8), uint32(v2), uint32(v7))
						l22:
							v1 = v1 - v7
							v8 = v8 + v7
							v12 = v12 + i32(8)
							if v12 != v10 {
								goto l23
							}
							goto l15
						}
					default:
						if v2 == i32(1) {
							goto l15
						}
					l27:
						{
							if v1 == 0 {
								m.fn27(i32(1272168), i32(19), i32(1069620))
								panic("unreachable")
							}
							t16 := int32(load32(m.memory[uint32(v12+i32(4)):]))
							v7 = t16
							t17 := int32(load32(m.memory[uint32(v12):]))
							v2 = t17
							t18 := int32(m.memory[uint32(v3)])
							m.memory[uint32(v8)] = byte(t18)
							v1 = v1 + i32(-1)
							if uint32(v1) < uint32(v7) {
								m.fn27(i32(1272168), i32(19), i32(1069620))
								panic("unreachable")
							}
							v8 = v8 + i32(1)
							if v7 == 0 {
								goto l26
							}
							memory_copy(m.memory, uint32(v8), uint32(v2), uint32(v7))
						l26:
							v1 = v1 - v7
							v8 = v8 + v7
							v12 = v12 + i32(8)
							if v12 != v10 {
								goto l27
							}
							goto l15
						}
					case 3:
						if v2 == i32(1) {
							goto l15
						}
					l31:
						{
							if uint32(v1) <= uint32(i32(3)) {
								m.fn27(i32(1272168), i32(19), i32(1069620))
								panic("unreachable")
							}
							t19 := int32(load32(m.memory[uint32(v12+i32(4)):]))
							v7 = t19
							t20 := int32(load32(m.memory[uint32(v12):]))
							v2 = t20
							t21 := int32(load32(m.memory[uint32(v3):]))
							store32(m.memory[uint32(v8):], uint32(t21))
							v1 = v1 + i32(-4)
							if uint32(v1) < uint32(v7) {
								m.fn27(i32(1272168), i32(19), i32(1069620))
								panic("unreachable")
							}
							v8 = v8 + i32(4)
							if v7 == 0 {
								goto l30
							}
							memory_copy(m.memory, uint32(v8), uint32(v2), uint32(v7))
						l30:
							v1 = v1 - v7
							v8 = v8 + v7
							v12 = v12 + i32(8)
							if v12 == v10 {
								goto l15
							}
							goto l31
						}
					}
				}
				v7 = v1 + i32(4)
				v8 = v8 + i32(-8)
				v1 = v1 + i32(8)
				t4 := int32(load32(m.memory[uint32(v7):]))
				v7 = t4
				v9 = v7 + v9
				if uint32(v9) >= uint32(v7) {
					goto l3
				}
			}
		}
	l1:
		m.fn139(i32(1069636), i32(53), i32(1069692))
		panic("unreachable")
	l15:
		t22 := int64(load64(m.memory[int64(uint32(v5))+4:]))
		store64(m.memory[uint32(v0):], uint64(t22))
		store32(m.memory[int64(uint32(v0))+8:], uint32(v9-v1))
	}
l32:
	m.g0 = v5 + i32(16)
}
func (m *Module) fn783(v0, v1, v2, v3 int32) {
	var v4, v5, v6, v7, v8, v9, v10, v11, v12 int32
	t0 := m.g0
	v4 = t0 - i32(16)
	m.g0 = v4
	v5 = i32(0)
	store32(m.memory[int64(uint32(v4))+12:], uint32(i32(0)))
	store64(m.memory[int64(uint32(v4))+4:], uint64(i64(0x100000000)))
	v6 = i32(1)
	v7 = i32(0)
	{
		if v2 == 0 {
			goto l0
		}
		v8 = v1 + v2
		t1 := v4 + i32(4)
		t2 := int32(uint32(v2) >> 2)
		var p3 int32
		if v2&i32(3) != i32(0) {
			p3 = 1
		}
		m.fn196(t1, i32(0), t2+p3, i32(1), i32(1))
	l13:
		{
			{
				{
					t4 := int32(int8(m.memory[uint32(v1)]))
					v6 = t4
					if v6 <= i32(-1) {
						goto l1
					}
					v1 = v1 + i32(1)
					v6 = v6 & i32(255)
					goto l2
				}
			l1:
				t5 := int32(m.memory[int64(uint32(v1))+1])
				v5 = t5 & i32(63)
				v7 = v6 & i32(31)
				if uint32(v6) > uint32(i32(-33)) {
					goto l3
				}
				v6 = v7<<6 | v5
				v1 = v1 + i32(2)
				goto l2
			l3:
				t6 := int32(m.memory[int64(uint32(v1))+2])
				v5 = v5<<6 | t6&i32(63)
				if uint32(v6) >= uint32(i32(-16)) {
					goto l4
				}
				v6 = v5 | v7<<12
				v1 = v1 + i32(3)
				goto l2
			l4:
				t7 := int32(m.memory[int64(uint32(v1))+3])
				v6 = v5<<6 | t7&i32(63) | v7<<18&i32(0x1c0000)
				v1 = v1 + i32(4)
			}
		l2:
			{
				{
					p8 := v6
					if uint32(v6+i32(-127)) < uint32(i32(33)) {
						p8 = i32(32)
					}
					p9 := p8
					if uint32(v6) < uint32(i32(32)) {
						p9 = i32(32)
					}
					v5 = p9
					var p10 int32
					if uint32(v5) < uint32(i32(128)) {
						p10 = 1
					}
					v9 = p10
					if v9 == 0 {
						goto l5
					}
					v2 = i32(1)
					goto l6
				}
			l5:
				if uint32(v5) >= uint32(i32(2048)) {
					goto l7
				}
				v2 = i32(2)
				goto l6
			l7:
				p11 := i32(4)
				if uint32(v5) < uint32(i32(65536)) {
					p11 = i32(3)
				}
				v2 = p11
			}
		l6:
			{
				t12 := int32(load32(m.memory[int64(uint32(v4))+4:]))
				t13 := int32(load32(m.memory[int64(uint32(v4))+12:]))
				t14 := v2
				v7 = t13
				if uint32(t14) <= uint32(t12-v7) {
					goto l8
				}
				m.fn196(v4+i32(4), v7, v2, i32(1), i32(1))
			}
		l8:
			t15 := int32(load32(m.memory[int64(uint32(v4))+8:]))
			v6 = t15
			v10 = v6 + v7
			if v9 != 0 {
				goto l9
			}
			v9 = v5&i32(63) | i32(-128)
			v11 = int32(uint32(v5) >> 6)
			if uint32(v5) >= uint32(i32(2048)) {
				v12 = int32(uint32(v5) >> 12)
				v11 = v11&i32(63) | i32(-128)
				if uint32(v5) > uint32(i32(0xffff)) {
					m.memory[int64(uint32(v10))+3] = byte(v9)
					m.memory[int64(uint32(v10))+2] = byte(v11)
					m.memory[int64(uint32(v10))+1] = byte(v12&i32(63) | i32(-128))
					m.memory[uint32(v10)] = byte(int32(uint32(v5)>>18) | i32(-16))
					goto l11
				}
				m.memory[int64(uint32(v10))+2] = byte(v9)
				m.memory[int64(uint32(v10))+1] = byte(v11)
				m.memory[uint32(v10)] = byte(v12 | i32(224))
				goto l11
			}
			m.memory[int64(uint32(v10))+1] = byte(v9)
			m.memory[uint32(v10)] = byte(v11 | i32(192))
			goto l11
		l9:
			m.memory[uint32(v10)] = byte(v5)
		l11:
			t16 := v4
			v7 = v2 + v7
			store32(m.memory[int64(uint32(t16))+12:], uint32(v7))
			if v1 != v8 {
				goto l13
			}
		}
		t17 := int32(load32(m.memory[int64(uint32(v4))+4:]))
		v5 = t17
	}
l0:
	t18 := v0
	t19 := v6
	t20 := v7
	t21 := v3
	var p22 int32
	if v3&i32(255) == 0 {
		p22 = 1
	}
	m.fn790(t18, t19, t20, t21, p22|i32(65536))
	{
		if v5 == 0 {
			goto l14
		}
		t23 := int32(load32(m.memory[uint32(v6+i32(-4)):]))
		v3 = t23
		v7 = v3 & i32(-8)
		t24 := v7
		v3 = v3 & i32(3)
		p25 := i32(8)
		if v3 != 0 {
			p25 = i32(4)
		}
		if uint32(t24) < uint32(p25+v5) {
			m.fn3(i32(1274224), i32(46), i32(1274272))
			panic("unreachable")
		}
		if v3 == 0 {
			goto l16
		}
		if uint32(v7) > uint32(v5+i32(39)) {
			m.fn3(i32(1274288), i32(46), i32(1274336))
			panic("unreachable")
		}
	l16:
		m.fn1(v6)
	}
l14:
	m.g0 = v4 + i32(16)
}
func (m *Module) fn784(v0, v1, v2, v3 int32) {
	var v4, v5, v6, v7, v8, v9, v10 int32
	v4 = v1 + v2
	v5 = i32(0)
	{
	l5:
		v6 = v5
		v7 = v1
		if v7 == v4 {
			goto l0
		}
		{
			{
				t0 := int32(int8(m.memory[uint32(v7)]))
				v8 = t0
				if v8 <= i32(-1) {
					goto l1
				}
				v1 = v7 + i32(1)
				v8 = v8 & i32(255)
				goto l2
			}
		l1:
			t1 := int32(m.memory[int64(uint32(v7))+1])
			v1 = t1 & i32(63)
			v5 = v8 & i32(31)
			if uint32(v8) > uint32(i32(-33)) {
				goto l3
			}
			v8 = v5<<6 | v1
			v1 = v7 + i32(2)
			goto l2
		l3:
			t2 := int32(m.memory[int64(uint32(v7))+2])
			v1 = v1<<6 | t2&i32(63)
			if uint32(v8) >= uint32(i32(-16)) {
				goto l4
			}
			v8 = v1 | v5<<12
			v1 = v7 + i32(3)
			goto l2
		l4:
			t3 := int32(m.memory[int64(uint32(v7))+3])
			v8 = v1<<6 | t3&i32(63) | v5<<18&i32(0x1c0000)
			v1 = v7 + i32(4)
		}
	l2:
		v5 = v1 - v7 + v6
		if v8 == i32(96) {
			goto l5
		}
		v7 = v5
	l11:
		{
			v9 = v7
			v7 = v1
			if v7 == v4 {
				goto l6
			}
			{
				{
					t4 := int32(int8(m.memory[uint32(v7)]))
					v8 = t4
					if v8 <= i32(-1) {
						goto l7
					}
					v1 = v7 + i32(1)
					v8 = v8 & i32(255)
					goto l8
				}
			l7:
				t5 := int32(m.memory[int64(uint32(v7))+1])
				v1 = t5 & i32(63)
				v10 = v8 & i32(31)
				if uint32(v8) > uint32(i32(-33)) {
					goto l9
				}
				v8 = v10<<6 | v1
				v1 = v7 + i32(2)
				goto l8
			l9:
				t6 := int32(m.memory[int64(uint32(v7))+2])
				v1 = v1<<6 | t6&i32(63)
				if uint32(v8) >= uint32(i32(-16)) {
					goto l10
				}
				v8 = v1 | v10<<12
				v1 = v7 + i32(3)
				goto l8
			l10:
				t7 := int32(m.memory[int64(uint32(v7))+3])
				v8 = v1<<6 | t7&i32(63) | v10<<18&i32(0x1c0000)
				v1 = v7 + i32(4)
			}
		l8:
			v7 = v1 - v7 + v9
			if v8 == i32(96) {
				goto l11
			}
			t8 := v6
			v8 = v9 - v5
			p9 := v8
			if uint32(v6) > uint32(v8) {
				p9 = t8
			}
			v6 = p9
			v5 = v7
			goto l11
		}
	l6:
		t10 := v6
		v1 = v2 - v5
		p11 := v1
		if uint32(v6) > uint32(v1) {
			p11 = t10
		}
		v2 = p11
	}
l0:
	{
		t12 := v3
		v1 = v2 + i32(1)
		p13 := v1
		if uint32(v3) > uint32(v1) {
			p13 = t12
		}
		v5 = p13
		if v5 <= i32(-1) {
			m.fn11()
			panic("unreachable")
		}
		t14 := m.fn7(v5)
		v8 = t14
		if v8 == 0 {
			m.fn12(i32(1), v5)
			panic("unreachable")
		}
		m.memory[uint32(v8)] = byte(i32(96))
		v1 = i32(1)
		v7 = int32(uint32(v5) >> 1)
		if v7 == 0 {
			goto l14
		}
		v1 = i32(1)
	l16:
		if v1 == 0 {
			goto l15
		}
		memory_copy(m.memory, uint32(v8+v1), uint32(v8), uint32(v1))
	l15:
		v1 = v1 << 1
		v7 = int32(uint32(v7) >> 1)
		if v7 != 0 {
			goto l16
		}
	l14:
		if v5 == v1 {
			goto l17
		}
		v7 = v5 - v1
		if v7 == 0 {
			goto l17
		}
		memory_copy(m.memory, uint32(v8+v1), uint32(v8), uint32(v7))
	l17:
		store32(m.memory[int64(uint32(v0))+8:], uint32(v5))
		store32(m.memory[int64(uint32(v0))+4:], uint32(v8))
		store32(m.memory[uint32(v0):], uint32(v5))
		return
	}
}
func (m *Module) fn785(v0, v1 int32) {
	var v2, v3, v4, v5 int32
	t0 := m.g0
	v2 = t0 - i32(32)
	m.g0 = v2
	{
		{
			t1 := int32(m.memory[int64(uint32(v1))+37])
			if t1 != 0 {
				goto l0
			}
			t2 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			v3 = t2
			m.fn198(v2+i32(4), v1)
			{
				{
					t3 := int32(load32(m.memory[int64(uint32(v2))+4:]))
					if t3 != i32(1) {
						goto l1
					}
					t4 := int32(load32(m.memory[int64(uint32(v1))+28:]))
					v4 = t4
					t5 := int32(load32(m.memory[int64(uint32(v2))+12:]))
					t6 := v1
					v5 = t5
					store32(m.memory[int64(uint32(t6))+28:], uint32(v5))
					v1 = v3 + v4
					v3 = v5 - v4
					goto l2
				}
			l1:
				t7 := int32(m.memory[int64(uint32(v1))+37])
				if t7 != 0 {
					goto l0
				}
				m.memory[int64(uint32(v1))+37] = byte(i32(1))
				{
					{
						t8 := int32(m.memory[int64(uint32(v1))+36])
						if t8 != i32(1) {
							goto l3
						}
						t9 := int32(load32(m.memory[int64(uint32(v1))+32:]))
						v3 = t9
						t10 := int32(load32(m.memory[int64(uint32(v1))+28:]))
						v4 = t10
						goto l4
					}
				l3:
					t11 := int32(load32(m.memory[int64(uint32(v1))+32:]))
					v3 = t11
					t12 := int32(load32(m.memory[int64(uint32(v1))+28:]))
					t13 := v3
					v4 = t12
					if t13 == v4 {
						goto l0
					}
				}
			l4:
				v3 = v3 - v4
				t14 := int32(load32(m.memory[int64(uint32(v1))+4:]))
				v1 = t14 + v4
			}
		l2:
			{
				if v3 == 0 {
					goto l5
				}
				t15 := v1
				v4 = v3 + i32(-1)
				t16 := int32(m.memory[uint32(t15+v4)])
				if t16 != i32(10) {
					goto l5
				}
				v3 = v3 + i32(-2)
				{
					if v4 != 0 {
						goto l6
					}
					v5 = i32(0)
					goto l7
				l6:
					t17 := int32(m.memory[uint32(v1+v3)])
					p18 := i32(0)
					if t17&i32(255) == i32(13) {
						p18 = v1
					}
					v5 = p18
				}
			l7:
				p19 := v4
				if v5 != 0 {
					p19 = v3
				}
				v3 = p19
				p20 := v1
				if v5 != 0 {
					p20 = v5
				}
				v1 = p20
			}
		l5:
			if v1 == 0 {
				goto l0
			}
			store32(m.memory[int64(uint32(v2))+20:], uint32(v3))
			store32(m.memory[int64(uint32(v2))+16:], uint32(v1))
			if v3 != 0 {
				store64(m.memory[int64(uint32(v2))+24:], uint64(int64(uint32(i32(1)))<<32|int64(uint32(v2+i32(16)))))
				m.fn13(v2+i32(4), i32(1051129), v2+i32(24))
				goto l10
			}
			t21 := m.fn7(i32(1))
			v1 = t21
			if v1 == 0 {
				m.fn12(i32(1), i32(1))
				panic("unreachable")
			}
			m.memory[uint32(v1)] = byte(i32(62))
			store32(m.memory[int64(uint32(v2))+12:], uint32(i32(1)))
			store32(m.memory[int64(uint32(v2))+8:], uint32(v1))
			store32(m.memory[int64(uint32(v2))+4:], uint32(i32(1)))
			goto l10
		}
	l0:
		store32(m.memory[uint32(v0):], uint32(i32(-1)))
		goto l11
	l10:
		t22 := int32(load32(m.memory[int64(uint32(v2))+12:]))
		store32(m.memory[int64(uint32(v0))+8:], uint32(t22))
		t23 := int64(load64(m.memory[int64(uint32(v2))+4:]))
		store64(m.memory[uint32(v0):], uint64(t23))
	}
l11:
	m.g0 = v2 + i32(32)
}
func (m *Module) fn786(v0, v1 int32) {
	var v2, v3, v4, v5, v6, v7 int32
	t0 := m.g0
	v2 = t0 - i32(48)
	m.g0 = v2
	t1 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	v3 = t1
	t2 := int32(load32(m.memory[uint32(v1):]))
	v4 = t2
	t3 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	v5 = t3
	{
	l1:
		{
			if v4 == v5 {
				goto l0
			}
			t4 := v1
			v6 = v4 + i32(32)
			store32(m.memory[uint32(t4):], uint32(v6))
			m.fn774(v2+i32(24), v4, v3)
			v4 = v6
			t5 := int32(load32(m.memory[int64(uint32(v2))+24:]))
			if t5 == i32(-1) {
				goto l1
			}
		}
		t6 := m.fn7(i32(48))
		v4 = t6
		if v4 == 0 {
			m.fn12(i32(4), i32(48))
			panic("unreachable")
		}
		t7 := int32(load32(m.memory[int64(uint32(v2))+32:]))
		store32(m.memory[int64(uint32(v4))+8:], uint32(t7))
		t8 := int64(load64(m.memory[int64(uint32(v2))+24:]))
		store64(m.memory[uint32(v4):], uint64(t8))
		store32(m.memory[int64(uint32(v2))+20:], uint32(i32(1)))
		store32(m.memory[int64(uint32(v2))+16:], uint32(v4))
		store32(m.memory[int64(uint32(v2))+12:], uint32(i32(4)))
		v1 = i32(1)
	l4:
		{
			if v6 == v5 {
				t15 := int32(load32(m.memory[int64(uint32(v2))+20:]))
				store32(m.memory[int64(uint32(v0))+8:], uint32(t15))
				t16 := int64(load64(m.memory[int64(uint32(v2))+12:]))
				store64(m.memory[uint32(v0):], uint64(t16))
				goto l6
			}
			m.fn774(v2+i32(36), v6, v3)
			v6 = v6 + i32(32)
			t9 := int32(load32(m.memory[int64(uint32(v2))+36:]))
			if t9 == i32(-1) {
				goto l4
			}
			{
				t10 := int32(load32(m.memory[int64(uint32(v2))+12:]))
				if v1 != t10 {
					goto l5
				}
				m.fn196(v2+i32(12), v1, i32(1), i32(4), i32(12))
				t11 := int32(load32(m.memory[int64(uint32(v2))+16:]))
				v4 = t11
			}
		l5:
			v7 = v4 + v1*i32(12)
			t12 := int32(load32(m.memory[int64(uint32(v2))+44:]))
			store32(m.memory[int64(uint32(v7))+8:], uint32(t12))
			t13 := int64(load64(m.memory[int64(uint32(v2))+36:]))
			store64(m.memory[uint32(v7):], uint64(t13))
			t14 := v2
			v1 = v1 + i32(1)
			store32(m.memory[int64(uint32(t14))+20:], uint32(v1))
			goto l4
		}
	}
l0:
	store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
	store64(m.memory[uint32(v0):], uint64(i64(0x400000000)))
l6:
	m.g0 = v2 + i32(48)
}
func (m *Module) fn787(v0, v1, v2 int32) {
	var v3, v4, v5, v6, v7, v8, v9 int32
	t0 := m.g0
	v3 = t0 - i32(128)
	m.g0 = v3
	store32(m.memory[int64(uint32(v3))+32:], uint32(i32(0)))
	store64(m.memory[int64(uint32(v3))+24:], uint64(i64(0x400000000)))
	v4 = i32(0)
	v5 = i32(4)
	{
		t1 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		v6 = t1
		if v6 == 0 {
			goto l0
		}
		v4 = v6 << 5
		t2 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		v1 = t2
	l1:
		m.fn794(v1, v2, v3+i32(24))
		v1 = v1 + i32(32)
		v4 = v4 + i32(-32)
		if v4 != 0 {
			goto l1
		}
		t3 := int32(load32(m.memory[int64(uint32(v3))+32:]))
		v4 = t3
		t4 := int32(load32(m.memory[int64(uint32(v3))+28:]))
		v5 = t4
	}
l0:
	m.fn202(v3+i32(88), v5, v4, i32(1078784), i32(4))
	store16(m.memory[int64(uint32(v3))+72:], uint16(i32(0)))
	store32(m.memory[int64(uint32(v3))+64:], uint32(i32(0)))
	m.memory[int64(uint32(v3))+60] = byte(i32(1))
	store32(m.memory[int64(uint32(v3))+56:], uint32(i32(10)))
	store32(m.memory[int64(uint32(v3))+48:], uint32(i32(0)))
	store32(m.memory[int64(uint32(v3))+36:], uint32(i32(10)))
	t5 := int32(load32(m.memory[int64(uint32(v3))+96:]))
	t6 := v3
	v1 = t5
	store32(m.memory[int64(uint32(t6))+68:], uint32(v1))
	store32(m.memory[int64(uint32(v3))+52:], uint32(v1))
	store32(m.memory[int64(uint32(v3))+44:], uint32(v1))
	t7 := int32(load32(m.memory[int64(uint32(v3))+92:]))
	t8 := v3
	v7 = t7
	store32(m.memory[int64(uint32(t8))+40:], uint32(v7))
	t9 := int32(load32(m.memory[int64(uint32(v3))+88:]))
	v8 = t9
	m.fn795(v3+i32(16), v3+i32(36))
	{
		{
			{
				t10 := int32(load32(m.memory[int64(uint32(v3))+16:]))
				v1 = t10
				if v1 != 0 {
					goto l2
				}
				m.fn782(v0, i32(4), i32(0), i32(1078784), i32(4))
				goto l3
			}
		l2:
			t11 := int32(load32(m.memory[int64(uint32(v3))+20:]))
			v4 = t11
			t12 := m.fn7(i32(32))
			v9 = t12
			if v9 == 0 {
				m.fn12(i32(4), i32(32))
				panic("unreachable")
			}
			store32(m.memory[uint32(v9):], uint32(v1))
			store32(m.memory[int64(uint32(v9))+4:], uint32(v4))
			store32(m.memory[int64(uint32(v3))+84:], uint32(i32(1)))
			store32(m.memory[int64(uint32(v3))+80:], uint32(v9))
			store32(m.memory[int64(uint32(v3))+76:], uint32(i32(4)))
			t13 := int64(load64(m.memory[int64(uint32(v3))+68:]))
			store64(m.memory[int64(uint32(v3))+120:], uint64(t13))
			t14 := int64(load64(m.memory[int64(uint32(v3))+60:]))
			store64(m.memory[int64(uint32(v3))+112:], uint64(t14))
			t15 := int64(load64(m.memory[int64(uint32(v3))+52:]))
			store64(m.memory[int64(uint32(v3))+104:], uint64(t15))
			t16 := int64(load64(m.memory[int64(uint32(v3))+44:]))
			store64(m.memory[int64(uint32(v3))+96:], uint64(t16))
			t17 := int64(load64(m.memory[int64(uint32(v3))+36:]))
			store64(m.memory[int64(uint32(v3))+88:], uint64(t17))
			v4 = i32(12)
			v1 = i32(1)
		l7:
			{
				m.fn795(v3+i32(8), v3+i32(88))
				t18 := int32(load32(m.memory[int64(uint32(v3))+8:]))
				v2 = t18
				if v2 == 0 {
					goto l5
				}
				t19 := int32(load32(m.memory[int64(uint32(v3))+12:]))
				v5 = t19
				{
					t20 := int32(load32(m.memory[int64(uint32(v3))+76:]))
					if v1 != t20 {
						goto l6
					}
					m.fn196(v3+i32(76), v1, i32(1), i32(4), i32(8))
					t21 := int32(load32(m.memory[int64(uint32(v3))+80:]))
					v9 = t21
				}
			l6:
				v6 = v9 + v4
				store32(m.memory[uint32(v6):], uint32(v5))
				store32(m.memory[uint32(v6+i32(-4)):], uint32(v2))
				t22 := v3
				v1 = v1 + i32(1)
				store32(m.memory[int64(uint32(t22))+84:], uint32(v1))
				v4 = v4 + i32(8)
				goto l7
			}
		l5:
			t23 := int32(load32(m.memory[int64(uint32(v3))+76:]))
			v4 = t23
			t24 := int32(load32(m.memory[int64(uint32(v3))+80:]))
			t25 := v0
			v2 = t24
			m.fn782(t25, v2, v1, i32(1078784), i32(4))
			if v4 == 0 {
				goto l3
			}
			t26 := int32(load32(m.memory[uint32(v2+i32(-4)):]))
			v1 = t26
			v5 = v1 & i32(-8)
			t27 := v5
			v1 = v1 & i32(3)
			p28 := i32(8)
			if v1 != 0 {
				p28 = i32(4)
			}
			v4 = v4 << 3
			if uint32(t27) < uint32(p28+v4) {
				m.fn3(i32(1274224), i32(46), i32(1274272))
				panic("unreachable")
			}
			if v1 == 0 {
				goto l9
			}
			if uint32(v5) > uint32(v4+i32(39)) {
				m.fn3(i32(1274288), i32(46), i32(1274336))
				panic("unreachable")
			}
		l9:
			m.fn1(v2)
		}
	l3:
		{
			if v8 == 0 {
				goto l11
			}
			t29 := int32(load32(m.memory[uint32(v7+i32(-4)):]))
			v1 = t29
			v4 = v1 & i32(-8)
			t30 := v4
			v1 = v1 & i32(3)
			p31 := i32(8)
			if v1 != 0 {
				p31 = i32(4)
			}
			if uint32(t30) < uint32(p31+v8) {
				m.fn3(i32(1274224), i32(46), i32(1274272))
				panic("unreachable")
			}
			if v1 == 0 {
				goto l13
			}
			if uint32(v4) > uint32(v8+i32(39)) {
				m.fn3(i32(1274288), i32(46), i32(1274336))
				panic("unreachable")
			}
		l13:
			m.fn1(v7)
		}
	l11:
		t32 := int32(load32(m.memory[int64(uint32(v3))+28:]))
		v8 = t32
		{
			t33 := int32(load32(m.memory[int64(uint32(v3))+32:]))
			v4 = t33
			if v4 == 0 {
				goto l15
			}
			v1 = v8
		l20:
			{
				t34 := int32(load32(m.memory[uint32(v1):]))
				v2 = t34
				if v2 == 0 {
					goto l16
				}
				t35 := int32(load32(m.memory[uint32(v1+i32(4)):]))
				v6 = t35
				t36 := int32(load32(m.memory[uint32(v6+i32(-4)):]))
				v5 = t36
				v9 = v5 & i32(-8)
				t37 := v9
				v5 = v5 & i32(3)
				p38 := i32(8)
				if v5 != 0 {
					p38 = i32(4)
				}
				if uint32(t37) < uint32(p38+v2) {
					m.fn3(i32(1274224), i32(46), i32(1274272))
					panic("unreachable")
				}
				if v5 == 0 {
					goto l18
				}
				if uint32(v9) > uint32(v2+i32(39)) {
					m.fn3(i32(1274288), i32(46), i32(1274336))
					panic("unreachable")
				}
			l18:
				m.fn1(v6)
			}
		l16:
			v1 = v1 + i32(12)
			v4 = v4 + i32(-1)
			if v4 != 0 {
				goto l20
			}
		}
	l15:
		{
			t39 := int32(load32(m.memory[int64(uint32(v3))+24:]))
			v1 = t39
			if v1 == 0 {
				goto l21
			}
			t40 := int32(load32(m.memory[uint32(v8+i32(-4)):]))
			v4 = t40
			v2 = v4 & i32(-8)
			t41 := v2
			v4 = v4 & i32(3)
			p42 := i32(8)
			if v4 != 0 {
				p42 = i32(4)
			}
			v1 = v1 * i32(12)
			if uint32(t41) < uint32(p42+v1) {
				m.fn3(i32(1274224), i32(46), i32(1274272))
				panic("unreachable")
			}
			if v4 == 0 {
				goto l23
			}
			if uint32(v2) > uint32(v1+i32(39)) {
				m.fn3(i32(1274288), i32(46), i32(1274336))
				panic("unreachable")
			}
		l23:
			m.fn1(v8)
		}
	l21:
		m.g0 = v3 + i32(128)
		return
	}
}
func (m *Module) fn788(v0, v1, v2, v3 int32) {
	var v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14, v15, v16, v17 int32
	t0 := m.g0
	v4 = t0 - i32(48)
	m.g0 = v4
	v5 = i32(0)
	store32(m.memory[int64(uint32(v4))+32:], uint32(i32(0)))
	store64(m.memory[int64(uint32(v4))+24:], uint64(i64(0x400000000)))
	if v2 == 0 {
		goto l0
	}
	v6 = v1 + v2*i32(28)
	v7 = v3 + i32(32)
	v8 = i32(4)
l35:
	{
		{
			{
				{
					{
						{
							t1 := int32(load32(m.memory[uint32(v1):]))
							v2 = t1
							p2 := i32(1)
							if uint32(v2) > uint32(i32(2)) {
								p2 = v2 + i32(-3)
							}
							switch p2 {
							case 1:
								t31 := int32(load32(m.memory[int64(uint32(v1))+24:]))
								v13 = t31
								t32 := int32(load32(m.memory[int64(uint32(v1))+20:]))
								v9 = t32
								{
									t33 := int32(load32(m.memory[int64(uint32(v1))+12:]))
									if t33 != 0 {
										{
											t46 := int32(load32(m.memory[int64(uint32(v4))+24:]))
											if v5 != t46 {
												goto l28
											}
											m.fn311(v4 + i32(24))
										}
									l28:
										t47 := int32(load32(m.memory[int64(uint32(v4))+28:]))
										v8 = t47
										v2 = v8 + v5<<4
										store32(m.memory[int64(uint32(v2))+12:], uint32(v1))
										store32(m.memory[int64(uint32(v2))+8:], uint32(v13))
										store32(m.memory[int64(uint32(v2))+4:], uint32(v9))
										store32(m.memory[uint32(v2):], uint32(i32(-0x80000000)))
										goto l29
									}
									v14 = v13 * i32(28)
									v2 = i32(0)
								l21:
									{
										if v14 == v2 {
											goto l7
										}
										t34 := v9
										v2 = v2 + i32(28)
										t35 := m.fn305(t34 + v2 + i32(-28))
										if t35 != 0 {
											goto l21
										}
									}
									m.fn788(v4+i32(36), v9, v13, v3)
									t36 := int32(load32(m.memory[int64(uint32(v4))+36:]))
									v9 = t36
									t37 := int32(load32(m.memory[int64(uint32(v4))+40:]))
									v14 = t37
									{
										{
											t38 := int32(load32(m.memory[int64(uint32(v4))+44:]))
											v2 = t38
											t39 := int32(load32(m.memory[int64(uint32(v4))+24:]))
											if uint32(v2) <= uint32(t39-v5) {
												goto l22
											}
											m.fn196(v4+i32(24), v5, v2, i32(4), i32(16))
											t40 := int32(load32(m.memory[int64(uint32(v4))+32:]))
											v5 = t40
											goto l23
										}
									l22:
										if v2 == 0 {
											goto l24
										}
									l23:
										t41 := int32(load32(m.memory[int64(uint32(v4))+28:]))
										v8 = t41
										v13 = v2 << 4
										if v13 == 0 {
											goto l24
										}
										memory_copy(m.memory, uint32(v8+v5<<4), uint32(v14), uint32(v13))
									}
								l24:
									t42 := v4
									v5 = v5 + v2
									store32(m.memory[int64(uint32(t42))+32:], uint32(v5))
									if v9 == 0 {
										goto l7
									}
									t43 := int32(load32(m.memory[uint32(v14+i32(-4)):]))
									v2 = t43
									v13 = v2 & i32(-8)
									t44 := v13
									v2 = v2 & i32(3)
									p45 := i32(8)
									if v2 != 0 {
										p45 = i32(4)
									}
									v9 = v9 << 4
									if uint32(t44) < uint32(p45|v9) {
										m.fn3(i32(1274224), i32(46), i32(1274272))
										panic("unreachable")
									}
									if v2 == 0 {
										goto l26
									}
									if uint32(v13) > uint32(v9+i32(39)) {
										m.fn3(i32(1274288), i32(46), i32(1274336))
										panic("unreachable")
									}
								l26:
									m.fn1(v14)
									goto l7
								}
							case 2:
								t48 := int32(load32(m.memory[int64(uint32(v1))+12:]))
								v9 = t48
								t49 := int32(load32(m.memory[int64(uint32(v1))+8:]))
								v14 = t49
								{
									t50 := int32(load32(m.memory[int64(uint32(v4))+24:]))
									if v5 != t50 {
										goto l30
									}
									m.fn311(v4 + i32(24))
								}
							l30:
								t51 := int32(load32(m.memory[int64(uint32(v4))+28:]))
								v8 = t51
								v2 = v8 + v5<<4
								store32(m.memory[int64(uint32(v2))+12:], uint32(v1+i32(16)))
								store32(m.memory[int64(uint32(v2))+8:], uint32(v9))
								store32(m.memory[int64(uint32(v2))+4:], uint32(v14))
								store32(m.memory[uint32(v2):], uint32(i32(-0x7fffffff)))
								goto l29
							case 4:
								t54 := int32(load32(m.memory[int64(uint32(v1))+12:]))
								v9 = t54
								t55 := int32(load32(m.memory[int64(uint32(v1))+8:]))
								v14 = t55
								{
									t56 := int32(load32(m.memory[int64(uint32(v4))+24:]))
									if v5 != t56 {
										goto l32
									}
									m.fn311(v4 + i32(24))
								}
							l32:
								t57 := int32(load32(m.memory[int64(uint32(v4))+28:]))
								v8 = t57
								v2 = v8 + v5<<4
								store32(m.memory[int64(uint32(v2))+8:], uint32(v9))
								store32(m.memory[int64(uint32(v2))+4:], uint32(v14))
								store32(m.memory[uint32(v2):], uint32(i32(-0x7ffffffd)))
								goto l29
							case 5:
								{
									t58 := int32(load32(m.memory[int64(uint32(v4))+24:]))
									if v5 != t58 {
										goto l33
									}
									m.fn311(v4 + i32(24))
									t59 := int32(load32(m.memory[int64(uint32(v4))+28:]))
									v8 = t59
								}
							l33:
								store32(m.memory[uint32(v8+v5<<4):], uint32(i32(-0x7ffffffc)))
								goto l29
							default:
								t3 := int32(load32(m.memory[int64(uint32(v1))+12:]))
								v9 = t3
								if v9 == 0 {
									goto l7
								}
								t4 := int32(load32(m.memory[int64(uint32(v1))+8:]))
								t5 := v4 + i32(8)
								v10 = t4
								m.fn143(t5, v10, v9)
								t6 := int32(load32(m.memory[int64(uint32(v4))+12:]))
								if t6 != 0 {
									goto l8
								}
								v11 = i32(0)
								v12 = i32(0)
								v13 = i32(0)
								v14 = i32(0)
								goto l9
							case 3:
								t7 := int32(load32(m.memory[int64(uint32(v1))+8:]))
								t8 := v4 + i32(16)
								t9 := v7
								v9 = t7
								t10 := int32(load32(m.memory[int64(uint32(v1))+12:]))
								t11 := v9
								v14 = t10
								m.fn792(t8, t9, t11, v14)
								t12 := int32(load32(m.memory[int64(uint32(v4))+16:]))
								if t12 != 0 {
									{
										t52 := int32(load32(m.memory[int64(uint32(v4))+24:]))
										if v5 != t52 {
											goto l31
										}
										m.fn311(v4 + i32(24))
									}
								l31:
									t53 := int32(load32(m.memory[int64(uint32(v4))+28:]))
									v8 = t53
									v2 = v8 + v5<<4
									store32(m.memory[int64(uint32(v2))+8:], uint32(v14))
									store32(m.memory[int64(uint32(v2))+4:], uint32(v9))
									store32(m.memory[uint32(v2):], uint32(i32(-0x7ffffffe)))
									goto l29
								}
								goto l7
							}
						}
					l8:
						t13 := int32(m.memory[int64(uint32(v1))+19])
						v11 = t13
						t14 := int32(m.memory[int64(uint32(v1))+18])
						v12 = t14
						t15 := int32(m.memory[int64(uint32(v1))+17])
						v13 = t15
						t16 := int32(m.memory[int64(uint32(v1))+16])
						v14 = t16
					}
				l9:
					if v5 == 0 {
						goto l11
					}
					t17 := int32(load32(m.memory[int64(uint32(v4))+28:]))
					v8 = t17
					v2 = v8 + v5<<4
					v15 = v2 + i32(-16)
					t18 := int32(load32(m.memory[uint32(v15):]))
					v16 = t18
					if v16 <= i32(-0x7ffffffc) {
						goto l12
					}
					t19 := int32(m.memory[uint32(v2+i32(-4))])
					if t19 != v14&i32(255) {
						goto l12
					}
					t20 := int32(m.memory[uint32(v2+i32(-3))])
					if t20 != v13&i32(255) {
						goto l12
					}
					t21 := int32(m.memory[uint32(v2+i32(-2))])
					if t21 != v12&i32(255) {
						goto l12
					}
					t22 := int32(m.memory[uint32(v2+i32(-1))])
					if t22 != v11&i32(255) {
						goto l12
					}
					{
						if v16 != i32(-1) {
							goto l13
						}
						v14 = v2 + i32(-12)
						{
							t23 := int32(load32(m.memory[uint32(v2+i32(-8)):]))
							v16 = t23
							if v16 != 0 {
								goto l14
							}
							store32(m.memory[uint32(v14):], uint32(i32(1)))
							v16 = i32(0)
							store32(m.memory[uint32(v15):], uint32(i32(0)))
							goto l13
						}
					l14:
						t24 := int32(load32(m.memory[uint32(v14):]))
						v12 = t24
						t25 := m.fn7(v16)
						v13 = t25
						if v13 == 0 {
							m.fn12(i32(1), v16)
							panic("unreachable")
						}
						if v16 == 0 {
							goto l16
						}
						memory_copy(m.memory, uint32(v13), uint32(v12), uint32(v16))
					l16:
						store32(m.memory[uint32(v14):], uint32(v13))
						store32(m.memory[uint32(v15):], uint32(v16))
						if v16 == i32(-1) {
							m.fn3(i32(1274396), i32(40), i32(1074144))
							panic("unreachable")
						}
					}
				l13:
					{
						t26 := v9
						t27 := v16
						v13 = v2 + i32(-8)
						t28 := int32(load32(m.memory[uint32(v13):]))
						v14 = t28
						if uint32(t26) <= uint32(t27-v14) {
							goto l18
						}
						m.fn196(v15, v14, v9, i32(1), i32(1))
						t29 := int32(load32(m.memory[uint32(v13):]))
						v14 = t29
					}
				l18:
					{
						if v9 == 0 {
							goto l19
						}
						t30 := int32(load32(m.memory[uint32(v2+i32(-12)):]))
						memory_copy(m.memory, uint32(t30+v14), uint32(v10), uint32(v9))
					}
				l19:
					store32(m.memory[uint32(v13):], uint32(v14+v9))
					goto l7
				}
			l12:
				if (v14|v13|v12)&i32(1) == 0 {
					goto l11
				}
				if v11&i32(1) != 0 {
					goto l11
				}
				if v5 == i32(1) {
					goto l11
				}
				if v16 < i32(-0x7ffffffb) {
					goto l11
				}
				t60 := int32(m.memory[uint32(v2+i32(-4))])
				if t60 != 0 {
					goto l11
				}
				t61 := int32(m.memory[uint32(v2+i32(-3))])
				if t61 != 0 {
					goto l11
				}
				t62 := int32(m.memory[uint32(v2+i32(-2))])
				if t62 != 0 {
					goto l11
				}
				t63 := int32(m.memory[uint32(v2+i32(-1))])
				if t63 != 0 {
					goto l11
				}
				t64 := int32(load32(m.memory[uint32(v2+i32(-12)):]))
				t65 := v4
				v15 = t64
				t66 := int32(load32(m.memory[uint32(v2+i32(-8)):]))
				t67 := v15
				v17 = t66
				m.fn143(t65, t67, v17)
				t68 := int32(load32(m.memory[int64(uint32(v4))+4:]))
				if t68 != 0 {
					goto l11
				}
				t69 := int32(load32(m.memory[uint32(v2+i32(-32)):]))
				if t69 < i32(-0x7ffffffb) {
					goto l11
				}
				t70 := int32(m.memory[uint32(v2+i32(-20))])
				if t70 != v14&i32(255) {
					goto l11
				}
				t71 := int32(m.memory[uint32(v2+i32(-19))])
				if t71 != v13&i32(255) {
					goto l11
				}
				t72 := int32(m.memory[uint32(v2+i32(-18))])
				if t72 != v12&i32(255) {
					goto l11
				}
				t73 := int32(m.memory[uint32(v2+i32(-17))])
				if t73 != 0 {
					goto l11
				}
				t74 := v4
				v5 = v5 + i32(-1)
				store32(m.memory[int64(uint32(t74))+32:], uint32(v5))
				t75 := m.fn797(v8 + v5<<4 + i32(-16))
				v2 = t75
				m.fn622(v2, v15, v17)
				m.fn622(v2, v10, v9)
				if v16 < i32(1) {
					goto l7
				}
				m.fn17(v15, v16, i32(1))
				goto l7
			}
		l11:
			{
				t76 := int32(load32(m.memory[int64(uint32(v4))+24:]))
				if v5 != t76 {
					goto l34
				}
				m.fn311(v4 + i32(24))
			}
		l34:
			t77 := int32(load32(m.memory[int64(uint32(v4))+28:]))
			v8 = t77
			v2 = v8 + v5<<4
			m.memory[int64(uint32(v2))+15] = byte(v11)
			m.memory[int64(uint32(v2))+14] = byte(v12)
			m.memory[int64(uint32(v2))+13] = byte(v13)
			m.memory[int64(uint32(v2))+12] = byte(v14)
			store32(m.memory[int64(uint32(v2))+8:], uint32(v9))
			store32(m.memory[int64(uint32(v2))+4:], uint32(v10))
			store32(m.memory[uint32(v2):], uint32(i32(-1)))
		}
	l29:
		t78 := v4
		v5 = v5 + i32(1)
		store32(m.memory[int64(uint32(t78))+32:], uint32(v5))
	}
l7:
	v1 = v1 + i32(28)
	if v1 != v6 {
		goto l35
	}
l0:
	t79 := int32(load32(m.memory[int64(uint32(v4))+32:]))
	store32(m.memory[int64(uint32(v0))+8:], uint32(t79))
	t80 := int64(load64(m.memory[int64(uint32(v4))+24:]))
	store64(m.memory[uint32(v0):], uint64(t80))
	m.g0 = v4 + i32(48)
}
func (m *Module) fn789(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := m.fn43(t0, v1)
	return t1
}
func (m *Module) fn790(v0, v1, v2, v3, v4 int32) {
	var v5, v6, v7, v8, v9, v10, v11, v12, v13, v14, v15, v16, v17, v18, v19, v20, v21, v22, v23, v24, v25, v26, v27, v28, v29, v30, v31, v32, v33, v34, v35, v36 int32
	t0 := m.g0
	v5 = t0 - i32(16)
	m.g0 = v5
	{
		{
			{
				if v2 == 0 {
					v1 = i32(8)
					v10 = i32(0)
					v11 = i32(4)
					v12 = i32(0)
					v17 = i32(0)
					v16 = i32(0)
					v15 = i32(0)
					v14 = i32(0)
					v18 = i32(0)
					goto l31
				}
				v6 = v1 + v2
				{
					{
						t1 := int32(int8(m.memory[uint32(v1)]))
						v7 = t1
						if v7 <= i32(-1) {
							goto l1
						}
						v1 = v1 + i32(1)
						v8 = v7 & i32(255)
						goto l2
					}
				l1:
					t2 := int32(m.memory[int64(uint32(v1))+1])
					v8 = t2 & i32(63)
					v9 = v7 & i32(31)
					if uint32(v7) > uint32(i32(-33)) {
						goto l3
					}
					v8 = v9<<6 | v8
					v1 = v1 + i32(2)
					goto l2
				l3:
					t3 := int32(m.memory[int64(uint32(v1))+2])
					v8 = v8<<6 | t3&i32(63)
					if uint32(v7) >= uint32(i32(-16)) {
						goto l4
					}
					v8 = v8 | v9<<12
					v1 = v1 + i32(3)
					goto l2
				l4:
					t4 := int32(m.memory[int64(uint32(v1))+3])
					v8 = v8<<6 | t4&i32(63) | v9<<18&i32(0x1c0000)
					v1 = v1 + i32(4)
				}
			l2:
				v7 = v6 - v1
				t5 := int32(uint32(v7) >> 2)
				var p6 int32
				if v7&i32(3) != i32(0) {
					p6 = 1
				}
				v7 = t5 + p6
				if uint32(v7) > uint32(i32(0x3ffffffe)) {
					goto l5
				}
				p7 := i32(3)
				if uint32(v7) > uint32(i32(3)) {
					p7 = v7
				}
				v10 = p7 + i32(1)
				v7 = v10 << 2
				if uint32(v7) >= uint32(i32(0x7ffffffd)) {
					goto l5
				}
				{
					if v7 != 0 {
						goto l6
					}
					v11 = i32(4)
					v10 = i32(0)
					goto l7
				l6:
					t8 := m.fn7(v7)
					v11 = t8
					if v11 == 0 {
						m.fn12(i32(4), v7)
						panic("unreachable")
					}
				}
			l7:
				store32(m.memory[uint32(v11):], uint32(v8))
				v12 = i32(1)
				store32(m.memory[int64(uint32(v5))+12:], uint32(i32(1)))
				store32(m.memory[int64(uint32(v5))+8:], uint32(v11))
				store32(m.memory[int64(uint32(v5))+4:], uint32(v10))
				{
					if v1 == v6 {
						v14 = i32(0)
						v15 = i32(0)
						v16 = i32(0)
						v17 = i32(0)
						v7 = v11
						v1 = i32(0)
						v18 = i32(0)
						goto l30
					}
					v7 = i32(4)
					v8 = i32(0)
				l15:
					{
						{
							{
								t9 := int32(int8(m.memory[uint32(v1)]))
								v9 = t9
								if v9 <= i32(-1) {
									goto l10
								}
								v1 = v1 + i32(1)
								v13 = v9 & i32(255)
								goto l11
							}
						l10:
							t10 := int32(m.memory[int64(uint32(v1))+1])
							v13 = t10 & i32(63)
							v12 = v9 & i32(31)
							if uint32(v9) > uint32(i32(-33)) {
								goto l12
							}
							v13 = v12<<6 | v13
							v1 = v1 + i32(2)
							goto l11
						l12:
							t11 := int32(m.memory[int64(uint32(v1))+2])
							v13 = v13<<6 | t11&i32(63)
							if uint32(v9) >= uint32(i32(-16)) {
								goto l13
							}
							v13 = v13 | v12<<12
							v1 = v1 + i32(3)
							goto l11
						l13:
							t12 := int32(m.memory[int64(uint32(v1))+3])
							v13 = v13<<6 | t12&i32(63) | v12<<18&i32(0x1c0000)
							v1 = v1 + i32(4)
						}
					l11:
						{
							v9 = v8 + i32(1)
							t13 := int32(load32(m.memory[int64(uint32(v5))+4:]))
							if v9 != t13 {
								goto l14
							}
							t14 := v5 + i32(4)
							t15 := v9
							v11 = v6 - v1
							t16 := int32(uint32(v11) >> 2)
							var p17 int32
							if v11&i32(3) != i32(0) {
								p17 = 1
							}
							m.fn196(t14, t15, t16+p17+i32(1), i32(4), i32(4))
							t18 := int32(load32(m.memory[int64(uint32(v5))+8:]))
							v11 = t18
						}
					l14:
						store32(m.memory[uint32(v11+v7):], uint32(v13))
						t19 := v5
						v13 = v8 + i32(2)
						store32(m.memory[int64(uint32(t19))+12:], uint32(v13))
						v7 = v7 + i32(4)
						v8 = v9
						if v1 != v6 {
							goto l15
						}
					}
					v6 = v13 & i32(1)
					v12 = (v9&i32(0x3fffffff) + i32(1)) & i32(0x7ffffffe)
					t20 := int32(load32(m.memory[int64(uint32(v5))+4:]))
					v10 = t20
					v14 = i32(0)
					v15 = i32(0)
					v16 = i32(0)
					v17 = i32(0)
					t21 := int32(load32(m.memory[int64(uint32(v5))+8:]))
					v11 = t21
					v8 = v11
					v1 = i32(0)
					v18 = i32(0)
				l29:
					{
						{
							v7 = v8
							t22 := int32(load32(m.memory[uint32(v7):]))
							v8 = t22
							switch v8 + i32(-93) {
							case 1, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32:
								goto l17
							default:
								if v8 != i32(42) {
									goto l17
								}
								v18 = i32(1)
								v19 = v1
								goto l17
							case 2:
								v14 = i32(1)
								v20 = v1
								goto l17
							case 33:
								v15 = i32(1)
								v21 = v1
								goto l17
							case 3:
								v16 = i32(1)
								v22 = v1
								goto l17
							case 0:
								v17 = i32(1)
								v23 = v1
							}
						}
					l17:
						v8 = v1 + i32(1)
						{
							t23 := int32(load32(m.memory[uint32(v7+i32(4)):]))
							v9 = t23
							switch v9 + i32(-93) {
							case 1, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32:
								goto l23
							case 2:
								v14 = i32(1)
								v20 = v8
								goto l23
							case 33:
								v15 = i32(1)
								v21 = v8
								goto l23
							case 3:
								v16 = i32(1)
								v22 = v8
								goto l23
							case 0:
								v17 = i32(1)
								v23 = v8
								goto l23
							default:
								if v9 != i32(42) {
									goto l23
								}
								v18 = i32(1)
								v19 = v8
							}
						}
					l23:
						v8 = v7 + i32(8)
						t24 := v12
						v1 = v1 + i32(2)
						if t24 == v1 {
							goto l28
						}
						goto l29
					}
				}
			}
		l28:
			if v6 != 0 {
				goto l32
			}
			v12 = v13
			goto l33
		l32:
			v7 = v7 + i32(8)
			v12 = v13
		l30:
			{
				t25 := int32(load32(m.memory[uint32(v7):]))
				v7 = t25
				switch v7 + i32(-93) {
				case 1, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32:
					goto l33
				default:
					if v7 != i32(42) {
						goto l33
					}
					v18 = i32(1)
					v19 = v1
					goto l33
				case 0:
					v17 = i32(1)
					v23 = v1
					goto l33
				case 3:
					v16 = i32(1)
					v22 = v1
					goto l33
				case 33:
					v15 = i32(1)
					v21 = v1
					goto l33
				case 2:
					v14 = i32(1)
					v20 = v1
				}
			}
		l33:
			v1 = v2 + i32(8)
			if v1 <= i32(-1) {
				goto l5
			}
			if v1 != 0 {
				goto l31
			}
			store32(m.memory[int64(uint32(v5))+12:], uint32(i32(0)))
			store64(m.memory[int64(uint32(v5))+4:], uint64(i64(0x100000000)))
			v24 = i32(1)
			goto l39
		l31:
			{
				t26 := m.fn7(v1)
				v24 = t26
				if v24 != 0 {
					goto l40
				}
				m.fn12(i32(1), v1)
				panic("unreachable")
			}
		l40:
			store32(m.memory[int64(uint32(v5))+12:], uint32(i32(0)))
			store32(m.memory[int64(uint32(v5))+8:], uint32(v24))
			store32(m.memory[int64(uint32(v5))+4:], uint32(v1))
			if v2 == 0 {
				goto l41
			}
		l39:
			v25 = v4 & i32(256)
			v26 = v4 & i32(65536)
			v27 = int32(uint32(v26) >> 16)
			v28 = v3 & i32(255)
			var p27 int32
			if v28 != i32(0) {
				p27 = 1
			}
			v29 = p27
			v3 = v29 | (v4 ^ i32(1))
			v30 = v12 + i32(-1)
			v31 = v11 + i32(4)
			v32 = v4 & i32(65792)
			v33 = v4 & i32(0x1010000)
			var p28 int32
			if uint32(v4) > uint32(i32(0xffffff)) {
				p28 = 1
			}
			v34 = p28
			v1 = i32(0)
			v9 = i32(0)
		l149:
			v4 = v3
			{
				{
					t29 := v11
					v13 = v9
					v35 = v13 << 2
					v8 = t29 + v35
					t30 := int32(load32(m.memory[uint32(v8):]))
					v7 = t30
					if v7 == i32(32) {
						goto l42
					}
					{
						if v7 != i32(10) {
							goto l43
						}
						{
							t31 := int32(load32(m.memory[int64(uint32(v5))+4:]))
							if t31 != v1 {
								goto l44
							}
							m.fn196(v5+i32(4), v1, i32(1), i32(1), i32(1))
							t32 := int32(load32(m.memory[int64(uint32(v5))+8:]))
							v24 = t32
						}
					l44:
						m.memory[uint32(v24+v1)] = byte(i32(10))
						t33 := v5
						v1 = v1 + i32(1)
						store32(m.memory[int64(uint32(t33))+12:], uint32(v1))
						v9 = v13 + i32(1)
						v3 = v29 & v4
						goto l45
					}
				l43:
					v3 = v4
					if uint32(v7+i32(-9)) < uint32(i32(5)) {
						goto l42
					}
					if uint32(v7) < uint32(i32(133)) {
						goto l46
					}
					v9 = int32(uint32(v7) >> 8)
					switch v9 + i32(-22) {
					case 1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25:
						goto l46
					case 0:
						v3 = v4
						if v7 != i32(5760) {
							goto l46
						}
						goto l42
					case 26:
						v3 = v4
						if v7 != i32(12288) {
							goto l46
						}
						goto l42
					case 10:
						v3 = v4
						t34 := int32(m.memory[int64(uint32(v7&i32(255)))+1139588])
						if t34&i32(2) == 0 {
							goto l46
						}
						goto l42
					default:
						if v9 != 0 {
							goto l46
						}
						v3 = v4
						t35 := int32(m.memory[int64(uint32(v7&i32(255)))+1139588])
						if t35&i32(1) != 0 {
							goto l42
						}
					}
				l46:
					v3 = i32(1)
				}
			l42:
				v24 = i32(-1)
				v36 = v27
				{
					{
						{
							{
								{
									{
										{
											{
												{
													{
														v9 = v13 + i32(1)
														if uint32(v9) >= uint32(v12) {
															goto l51
														}
														{
															t36 := int32(load32(m.memory[uint32(v11+v9<<2):]))
															v24 = t36
															if uint32(v24+i32(-9)) < uint32(i32(5)) {
																goto l52
															}
															if v24 != i32(32) {
																goto l53
															}
														}
													l52:
														v2 = i32(1)
														v36 = i32(0)
														v6 = i32(1)
														switch v7 + i32(-33) {
														case 0:
															goto l54
														case 2:
															goto l56
														case 5:
															goto l57
														case 9:
															v2 = i32(1)
															var p46 int32
															if v25 == 0 {
																p46 = 1
															}
															if p46&v4 == 0 {
																goto l65
															}
															v6 = i32(1)
															goto l54
														case 10:
															v2 = i32(1)
															if v4&i32(1) == 0 {
																goto l65
															}
															v6 = i32(1)
															goto l54
														case 12:
															v2 = i32(1)
															if v4&i32(1) == 0 {
																goto l65
															}
															v6 = i32(1)
															goto l54
														case 27:
															goto l61
														case 28:
															goto l62
														case 29:
															goto l63
														case 58:
															goto l64
														case 59:
															goto l65
														case 60:
															goto l66
														case 62:
															goto l67
														case 63:
															goto l68
														case 91:
															goto l69
														case 93:
															v2 = i32(1)
															if v25 != 0 {
																goto l65
															}
															v6 = i32(1)
															goto l54
														default:
															goto l55
														}
													l53:
														if uint32(v24) < uint32(i32(133)) {
															v36 = i32(1)
															v2 = i32(1)
															v6 = i32(1)
															switch v7 + i32(-33) {
															case 0, 10:
																goto l54
															case 2:
																goto l56
															case 5:
																goto l57
															case 9:
																goto l83
															case 12:
																v2 = i32(1)
																if v4&i32(1) == 0 {
																	goto l95
																}
																v6 = i32(1)
																goto l54
															case 27:
																goto l61
															case 28:
																goto l62
															case 29:
																goto l63
															case 58:
																goto l64
															case 59:
																goto l65
															case 60:
																goto l66
															case 62:
																goto l67
															case 63:
																goto l68
															case 91:
																goto l69
															case 93:
																goto l85
															default:
																goto l55
															}
														}
														v6 = i32(0)
														v2 = int32(uint32(v24) >> 8)
														switch v2 + i32(-22) {
														case 1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25:
															goto l73
														case 0:
															var p37 int32
															if v24 == i32(5760) {
																p37 = 1
															}
															v6 = p37
															goto l73
														case 26:
															var p38 int32
															if v24 == i32(12288) {
																p38 = 1
															}
															v6 = p38
															goto l73
														default:
															if v2 != 0 {
																goto l73
															}
															t39 := int32(m.memory[int64(uint32(v24&i32(255)))+1139588])
															v6 = t39
															goto l73
														case 10:
															t40 := int32(m.memory[int64(uint32(v24&i32(255)))+1139588])
															v6 = int32(uint32(t40&i32(2)) >> 1)
														}
													l73:
														v36 = v6 ^ i32(1)
													l51:
														switch v7 + i32(-33) {
														case 0:
															v2 = i32(1)
															if v24 == i32(-1) {
																if v26 != 0 {
																	goto l65
																}
																v6 = i32(1)
																goto l54
															}
															v6 = i32(1)
															goto l54
														case 2:
															goto l56
														case 5:
															goto l57
														case 9:
															var p49 int32
															if v25 == 0 {
																p49 = 1
															}
															if p49&v4 == 0 {
																goto l65
															}
															v2 = i32(1)
															if v36&i32(1) != 0 {
																goto l97
															}
															v6 = i32(1)
															goto l54
														case 10:
															v2 = i32(1)
															if (v4|v36)&i32(1) == 0 {
																goto l65
															}
															v6 = i32(1)
															goto l54
														case 12:
															v2 = i32(1)
															if v4&i32(1) == 0 {
																goto l94
															}
															v6 = i32(1)
															goto l54
														case 27:
															if v24 != i32(-1) {
																goto l61
															}
															v2 = i32(1)
															v6 = i32(1)
															goto l54
														case 28:
															goto l62
														case 29:
															goto l63
														case 58:
															goto l64
														case 59:
															goto l65
														case 60:
															goto l66
														case 62:
															goto l67
														case 63:
															goto l68
														case 91:
															goto l69
														case 93:
															if v25 != 0 {
																goto l65
															}
															v2 = i32(1)
															if v36&i32(1) != 0 {
																goto l99
															}
															v6 = i32(1)
															goto l54
														default:
															goto l55
														}
													l55:
														;
														var p41 int32
														if uint32(v7+i32(-58)) < uint32(i32(-10)) {
															p41 = 1
														}
														if (p41|v4)&i32(1) != 0 {
															goto l86
														}
														if uint32(v12) <= uint32(v13) {
															goto l86
														}
														v2 = i32(0)
														v6 = v8
														v4 = v13
													l88:
														{
															t42 := int32(load32(m.memory[uint32(v6):]))
															v24 = t42
															if uint32(v24+i32(-48)) > uint32(i32(9)) {
																switch v24 + i32(-41) {
																default:
																	goto l86
																case 0, 5:
																	{
																		v35 = v4 + i32(1)
																		if uint32(v35) >= uint32(v12) {
																			goto l115
																		}
																		t67 := int32(load32(m.memory[uint32(v11+v35<<2):]))
																		v24 = t67
																		if uint32(v24+i32(-9)) < uint32(i32(5)) {
																			goto l115
																		}
																		if v24 == i32(32) {
																			goto l115
																		}
																		if uint32(v24) < uint32(i32(133)) {
																			goto l86
																		}
																		t68 := m.fn796(v24)
																		if t68 == 0 {
																			goto l86
																		}
																	}
																l115:
																	{
																		if uint32(v4) < uint32(v13) {
																			m.fn120(v13, v4, v12, i32(1078824))
																			panic("unreachable")
																		}
																		{
																			{
																				v9 = v4 - v13
																				t69 := int32(load32(m.memory[int64(uint32(v5))+4:]))
																				t70 := v9
																				v7 = t69
																				if uint32(t70) <= uint32(v7-v1) {
																					goto l117
																				}
																				m.fn196(v5+i32(4), v1, v9, i32(1), i32(1))
																				t71 := int32(load32(m.memory[int64(uint32(v5))+12:]))
																				v1 = t71
																				goto l128
																			}
																		l117:
																			if v4 == v13 {
																				goto l119
																			}
																		l128:
																			{
																				{
																					{
																						t72 := int32(load32(m.memory[uint32(v8):]))
																						v7 = t72
																						var p73 int32
																						if uint32(v7) < uint32(i32(128)) {
																							p73 = 1
																						}
																						v4 = p73
																						if v4 == 0 {
																							goto l120
																						}
																						v9 = i32(1)
																						goto l121
																					}
																				l120:
																					if uint32(v7) >= uint32(i32(2048)) {
																						goto l122
																					}
																					v9 = i32(2)
																					goto l121
																				l122:
																					p74 := i32(4)
																					if uint32(v7) < uint32(i32(65536)) {
																						p74 = i32(3)
																					}
																					v9 = p74
																				}
																			l121:
																				{
																					t75 := int32(load32(m.memory[int64(uint32(v5))+4:]))
																					if uint32(v9) <= uint32(t75-v1) {
																						goto l123
																					}
																					m.fn196(v5+i32(4), v1, v9, i32(1), i32(1))
																				}
																			l123:
																				t76 := int32(load32(m.memory[int64(uint32(v5))+8:]))
																				v13 = t76 + v1
																				if v4 != 0 {
																					goto l124
																				}
																				v24 = v7&i32(63) | i32(-128)
																				v4 = int32(uint32(v7) >> 6)
																				if uint32(v7) >= uint32(i32(2048)) {
																					v36 = int32(uint32(v7) >> 12)
																					v4 = v4&i32(63) | i32(-128)
																					if uint32(v7) > uint32(i32(0xffff)) {
																						m.memory[int64(uint32(v13))+3] = byte(v24)
																						m.memory[int64(uint32(v13))+2] = byte(v4)
																						m.memory[int64(uint32(v13))+1] = byte(v36&i32(63) | i32(-128))
																						m.memory[uint32(v13)] = byte(int32(uint32(v7)>>18) | i32(-16))
																						goto l126
																					}
																					m.memory[int64(uint32(v13))+2] = byte(v24)
																					m.memory[int64(uint32(v13))+1] = byte(v4)
																					m.memory[uint32(v13)] = byte(v36 | i32(224))
																					goto l126
																				}
																				m.memory[int64(uint32(v13))+1] = byte(v24)
																				m.memory[uint32(v13)] = byte(v4 | i32(192))
																				goto l126
																			l124:
																				m.memory[uint32(v13)] = byte(v7)
																			l126:
																				t77 := v5
																				v1 = v9 + v1
																				store32(m.memory[int64(uint32(t77))+12:], uint32(v1))
																				v8 = v8 + i32(4)
																				v2 = v2 + i32(-1)
																				if v2 != 0 {
																					goto l128
																				}
																			}
																			t78 := int32(load32(m.memory[int64(uint32(v5))+4:]))
																			v7 = t78
																		}
																	l119:
																		if v7 != v1 {
																			goto l129
																		}
																		m.fn196(v5+i32(4), v7, i32(1), i32(1), i32(1))
																	l129:
																		t79 := int32(load32(m.memory[int64(uint32(v5))+8:]))
																		v24 = t79
																		m.memory[uint32(v24+v1)] = byte(i32(92))
																		v8 = i32(1)
																		t80 := v5
																		v7 = v1 + i32(1)
																		store32(m.memory[int64(uint32(t80))+12:], uint32(v7))
																		{
																			t81 := int32(load32(m.memory[uint32(v6):]))
																			v1 = t81
																			var p82 int32
																			if uint32(v1) < uint32(i32(128)) {
																				p82 = 1
																			}
																			v13 = p82
																			if v13 != 0 {
																				goto l130
																			}
																			v8 = i32(2)
																			if uint32(v1) < uint32(i32(2048)) {
																				goto l130
																			}
																			p83 := i32(4)
																			if uint32(v1) < uint32(i32(65536)) {
																				p83 = i32(3)
																			}
																			v8 = p83
																		}
																	l130:
																		{
																			t84 := int32(load32(m.memory[int64(uint32(v5))+4:]))
																			if uint32(v8) <= uint32(t84-v7) {
																				goto l131
																			}
																			m.fn196(v5+i32(4), v7, v8, i32(1), i32(1))
																			t85 := int32(load32(m.memory[int64(uint32(v5))+8:]))
																			v24 = t85
																		}
																	l131:
																		v9 = v24 + v7
																		if v13 != 0 {
																			m.memory[uint32(v9)] = byte(v1)
																			goto l134
																		}
																		v6 = v1&i32(63) | i32(-128)
																		v13 = int32(uint32(v1) >> 6)
																		if uint32(v1) >= uint32(i32(2048)) {
																			v2 = int32(uint32(v1) >> 12)
																			v13 = v13&i32(63) | i32(-128)
																			if uint32(v1) > uint32(i32(0xffff)) {
																				m.memory[int64(uint32(v9))+3] = byte(v6)
																				m.memory[int64(uint32(v9))+2] = byte(v13)
																				m.memory[int64(uint32(v9))+1] = byte(v2&i32(63) | i32(-128))
																				m.memory[uint32(v9)] = byte(int32(uint32(v1)>>18) | i32(-16))
																				goto l134
																			}
																			m.memory[int64(uint32(v9))+2] = byte(v6)
																			m.memory[int64(uint32(v9))+1] = byte(v13)
																			m.memory[uint32(v9)] = byte(v2 | i32(224))
																			goto l134
																		}
																		m.memory[int64(uint32(v9))+1] = byte(v6)
																		m.memory[uint32(v9)] = byte(v13 | i32(192))
																		goto l134
																	}
																l134:
																	t86 := v5
																	v1 = v8 + v7
																	store32(m.memory[int64(uint32(t86))+12:], uint32(v1))
																	v9 = v35
																	goto l45
																}
															}
															v2 = v2 + i32(1)
															v6 = v6 + i32(4)
															t43 := v12
															v4 = v4 + i32(1)
															if t43 != v4 {
																goto l88
															}
															goto l86
														}
													}
												l66:
													v2 = i32(1)
													if v34 != 0 {
														goto l65
													}
													v6 = i32(1)
													goto l54
												l68:
													if v32 != 0 {
														goto l65
													}
													v2 = i32(1)
													t44 := v16
													var p45 int32
													if uint32(v22) > uint32(v13) {
														p45 = 1
													}
													if t44&p45 != 0 {
														goto l65
													}
													v6 = i32(1)
													goto l54
												}
											l67:
												if v13 != 0 {
													v6 = i32(1)
													t53 := int32(load32(m.memory[uint32(v8+i32(-4)):]))
													v8 = t53
													if uint32(v8&i32(2097119)+i32(-65)) < uint32(i32(26)) {
														goto l90
													}
													if uint32(v8+i32(-48)) <= uint32(i32(9)) {
														goto l90
													}
													if uint32(v8) >= uint32(i32(170)) {
														t54 := m.fn768(v8)
														if t54 != 0 {
															goto l90
														}
														v6 = i32(0)
														if uint32(v8) < uint32(i32(178)) {
															goto l90
														}
														t55 := m.fn769(v8)
														v6 = t55
														goto l90
													}
													v6 = i32(0)
													goto l90
												}
												v6 = i32(0)
												goto l90
											l64:
												if v33 != 0 {
													goto l65
												}
												v2 = i32(1)
												t47 := v17
												var p48 int32
												if uint32(v23) > uint32(v13) {
													p48 = 1
												}
												if t47&p48 != 0 {
													goto l65
												}
												v6 = i32(1)
												goto l54
											}
										l69:
											v2 = i32(1)
											if v28 == i32(2) {
												goto l65
											}
											v6 = i32(1)
											goto l54
										l57:
											v2 = i32(1)
											if uint32(v12-v13) > uint32(i32(1)) {
												t58 := int32(load32(m.memory[int64(uint32(v8))+4:]))
												if t58 == i32(35) {
													goto l101
												}
												v2 = v30 - v13
												v8 = v31 + v35
												v13 = i32(0)
											l103:
												{
													{
														t59 := int32(load32(m.memory[uint32(v8):]))
														v6 = t59
														if uint32(v6+i32(-48)) < uint32(i32(10)) {
															goto l102
														}
														if uint32(v6&i32(2097119)+i32(-65)) <= uint32(i32(25)) {
															goto l102
														}
														if v13 == 0 {
															goto l86
														}
														if v6 != i32(59) {
															goto l86
														}
														goto l101
													}
												l102:
													v8 = v8 + i32(4)
													t60 := v2
													v13 = v13 + i32(1)
													if t60 != v13 {
														goto l103
													}
													goto l86
												}
											}
											v6 = i32(1)
											goto l54
										l56:
											v2 = i32(1)
											if v4&i32(1) == 0 {
											l104:
												{
													t61 := v12
													v6 = v13
													if t61 == v6 {
														goto l65
													}
													v13 = v6 + i32(1)
													t62 := int32(load32(m.memory[uint32(v8):]))
													v2 = t62
													v8 = v8 + i32(4)
													if v2 == i32(35) {
														goto l104
													}
												}
												if uint32(v6) >= uint32(v12) {
													goto l65
												}
												if uint32(v2+i32(-9)) < uint32(i32(5)) {
													goto l65
												}
												if v2 == i32(32) {
													goto l65
												}
												if uint32(v2) < uint32(i32(133)) {
													goto l86
												}
												v8 = int32(uint32(v2) >> 8)
												switch v8 + i32(-22) {
												case 1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25:
													goto l86
												default:
													if v8 != 0 {
														goto l86
													}
													t63 := int32(m.memory[int64(uint32(v2&i32(255)))+1139588])
													if t63&i32(1) != 0 {
														goto l65
													}
													goto l86
												case 0:
													if v2 == i32(5760) {
														goto l65
													}
													goto l86
												case 26:
													if v2 == i32(12288) {
														goto l65
													}
													goto l86
												case 10:
													t64 := int32(m.memory[int64(uint32(v2&i32(255)))+1139588])
													if t64&i32(2) != 0 {
														goto l65
													}
													goto l86
												}
											}
											v6 = i32(1)
											goto l54
										l63:
											v2 = i32(1)
											if v4&i32(1) == 0 {
												goto l65
											}
											v6 = i32(1)
											goto l54
										l62:
											v2 = i32(1)
											if v4&i32(1) == 0 {
												v13 = (v12 - v13) << 2
											l113:
												{
													t66 := int32(load32(m.memory[uint32(v8):]))
													v6 = t66
													switch v6 + i32(-9) {
													case 1:
														goto l65
													case 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22:
														goto l86
													default:
														if v6 != i32(61) {
															goto l86
														}
														fallthrough
													case 0, 23:
														v8 = v8 + i32(4)
														v13 = v13 + i32(-4)
														if v13 == 0 {
															goto l65
														}
														goto l113
													}
												}
											}
											v6 = i32(1)
											goto l54
										l83:
											;
											var p50 int32
											if v25 == 0 {
												p50 = 1
											}
											if p50&v4 == 0 {
												goto l65
											}
										}
									l97:
										if v26 != 0 {
											goto l65
										}
										v2 = i32(1)
										t51 := v18
										var p52 int32
										if uint32(v19) > uint32(v13) {
											p52 = 1
										}
										if t51&p52 != 0 {
											goto l65
										}
										v6 = i32(1)
										goto l54
									}
								l85:
									if v25 != 0 {
										goto l65
									}
								l99:
									if v26 != 0 {
										goto l65
									}
									v2 = i32(1)
									t56 := v15
									var p57 int32
									if uint32(v21) > uint32(v13) {
										p57 = 1
									}
									if t56&p57 != 0 {
										goto l65
									}
									v6 = i32(1)
									goto l54
								}
							l61:
								if uint32(v24&i32(-33)+i32(-65)) < uint32(i32(26)) {
									goto l65
								}
								v2 = i32(1)
								v6 = i32(1)
								switch v24 + i32(-47) {
								case 0, 16:
									goto l65
								case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15:
									goto l54
								default:
									if v24 == i32(33) {
										goto l65
									}
									v6 = i32(1)
									goto l54
								}
							l94:
								if v36&i32(1) == 0 {
									goto l65
								}
							l95:
								v13 = (v12 - v13) << 2
							l110:
								{
									t65 := int32(load32(m.memory[uint32(v8):]))
									switch t65 + i32(-9) {
									case 1:
										goto l65
									default:
										goto l86
									case 0, 23, 36:
										v8 = v8 + i32(4)
										v13 = v13 + i32(-4)
										if v13 != 0 {
											goto l110
										}
										goto l65
									}
								}
							l101:
								{
									t87 := int32(load32(m.memory[int64(uint32(v5))+4:]))
									if uint32(t87-v1) > uint32(i32(4)) {
										goto l136
									}
									m.fn196(v5+i32(4), v1, i32(5), i32(1), i32(1))
									t88 := int32(load32(m.memory[int64(uint32(v5))+12:]))
									v1 = t88
								}
							l136:
								t89 := int32(load32(m.memory[int64(uint32(v5))+8:]))
								v24 = t89
								v7 = v24 + v1
								t90 := int32(load32(m.memory[int64(uint32(i32(0)))+1078816:]))
								store32(m.memory[uint32(v7):], uint32(t90))
								t91 := int32(m.memory[int64(uint32(i32(0)))+1078820])
								m.memory[int64(uint32(v7))+4] = byte(t91)
								t92 := v5
								v1 = v1 + i32(5)
								store32(m.memory[int64(uint32(t92))+12:], uint32(v1))
								goto l45
							}
						l90:
							if v24 == i32(-1) {
								if v25 != 0 {
									goto l65
								}
								v2 = i32(1)
								if v36&i32(1) != 0 {
									goto l141
								}
								v6 = i32(1)
								goto l54
							}
							v8 = i32(1)
							{
								if uint32(v24&i32(2097119)+i32(-65)) < uint32(i32(26)) {
									goto l138
								}
								if uint32(v24+i32(-48)) < uint32(i32(10)) {
									goto l138
								}
								if uint32(v24) >= uint32(i32(170)) {
									goto l139
								}
								v8 = i32(0)
								goto l138
							l139:
								t93 := m.fn768(v24)
								if t93 != 0 {
									goto l138
								}
								v8 = i32(0)
								if uint32(v24) < uint32(i32(178)) {
									goto l138
								}
								t94 := m.fn769(v24)
								v8 = t94
							}
						l138:
							if v25 != 0 {
								goto l65
							}
							v2 = i32(1)
							if v6&v8 == 0 {
								if (v36^i32(1))&i32(1) == 0 {
									goto l141
								}
								v6 = i32(1)
								goto l54
							}
							v6 = i32(1)
							goto l54
						l141:
							if v26 != 0 {
								goto l65
							}
							v2 = i32(1)
							t95 := v14
							var p96 int32
							if uint32(v20) > uint32(v13) {
								p96 = 1
							}
							if t95&p96 != 0 {
								goto l65
							}
							v6 = i32(1)
							goto l54
						}
					l65:
						{
							t97 := int32(load32(m.memory[int64(uint32(v5))+4:]))
							if t97 != v1 {
								goto l142
							}
							m.fn196(v5+i32(4), v1, i32(1), i32(1), i32(1))
						}
					l142:
						t98 := int32(load32(m.memory[int64(uint32(v5))+8:]))
						m.memory[uint32(t98+v1)] = byte(i32(92))
						t99 := v5
						v1 = v1 + i32(1)
						store32(m.memory[int64(uint32(t99))+12:], uint32(v1))
					}
				l86:
					v2 = i32(1)
					if uint32(v7) >= uint32(i32(128)) {
						goto l143
					}
					v6 = i32(1)
					goto l54
				l143:
					v6 = i32(2)
					v2 = i32(0)
					if uint32(v7) < uint32(i32(2048)) {
						goto l54
					}
					p100 := i32(4)
					if uint32(v7) < uint32(i32(65536)) {
						p100 = i32(3)
					}
					v6 = p100
				}
			l54:
				{
					t101 := int32(load32(m.memory[int64(uint32(v5))+4:]))
					if uint32(v6) <= uint32(t101-v1) {
						goto l144
					}
					m.fn196(v5+i32(4), v1, v6, i32(1), i32(1))
				}
			l144:
				t102 := int32(load32(m.memory[int64(uint32(v5))+8:]))
				v24 = t102
				v8 = v24 + v1
				if v2 != 0 {
					goto l145
				}
				v13 = v7&i32(63) | i32(-128)
				v2 = int32(uint32(v7) >> 6)
				if uint32(v7) >= uint32(i32(2048)) {
					v4 = int32(uint32(v7) >> 12)
					v2 = v2&i32(63) | i32(-128)
					if uint32(v7) > uint32(i32(0xffff)) {
						m.memory[int64(uint32(v8))+3] = byte(v13)
						m.memory[int64(uint32(v8))+2] = byte(v2)
						m.memory[int64(uint32(v8))+1] = byte(v4&i32(63) | i32(-128))
						m.memory[uint32(v8)] = byte(int32(uint32(v7)>>18) | i32(-16))
						goto l147
					}
					m.memory[int64(uint32(v8))+2] = byte(v13)
					m.memory[int64(uint32(v8))+1] = byte(v2)
					m.memory[uint32(v8)] = byte(v4 | i32(224))
					goto l147
				}
				m.memory[int64(uint32(v8))+1] = byte(v13)
				m.memory[uint32(v8)] = byte(v2 | i32(192))
				goto l147
			l145:
				m.memory[uint32(v8)] = byte(v7)
			l147:
				t103 := v5
				v1 = v6 + v1
				store32(m.memory[int64(uint32(t103))+12:], uint32(v1))
			}
		l45:
			if uint32(v9) < uint32(v12) {
				goto l149
			}
		}
	l41:
		t104 := int32(load32(m.memory[int64(uint32(v5))+12:]))
		store32(m.memory[int64(uint32(v0))+8:], uint32(t104))
		t105 := int64(load64(m.memory[int64(uint32(v5))+4:]))
		store64(m.memory[uint32(v0):], uint64(t105))
		{
			if v10 == 0 {
				goto l150
			}
			t106 := int32(load32(m.memory[uint32(v11+i32(-4)):]))
			v1 = t106
			v7 = v1 & i32(-8)
			t107 := v7
			v1 = v1 & i32(3)
			p108 := i32(8)
			if v1 != 0 {
				p108 = i32(4)
			}
			v8 = v10 << 2
			if uint32(t107) < uint32(p108+v8) {
				m.fn3(i32(1274224), i32(46), i32(1274272))
				panic("unreachable")
			}
			if v1 == 0 {
				goto l152
			}
			if uint32(v7) > uint32(v8+i32(39)) {
				m.fn3(i32(1274288), i32(46), i32(1274336))
				panic("unreachable")
			}
		l152:
			m.fn1(v11)
		}
	l150:
		m.g0 = v5 + i32(16)
		return
	}
l5:
	m.fn11()
	panic("unreachable")
}
func (m *Module) fn791(v0, v1, v2 int32) {
	var v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13 int32
	t0 := m.g0
	v3 = t0 - i32(32)
	m.g0 = v3
	{
		{
			if v2 <= i32(-1) {
				m.fn11()
				panic("unreachable")
			}
			if v2 != 0 {
				t1 := m.fn7(v2)
				v4 = t1
				if v4 == 0 {
					m.fn12(i32(1), v2)
					panic("unreachable")
				}
				v5 = v3 + i32(24) | i32(2)
				v6 = v3 + i32(24) | i32(1)
				store32(m.memory[int64(uint32(v3))+20:], uint32(i32(0)))
				store32(m.memory[int64(uint32(v3))+16:], uint32(v4))
				store32(m.memory[int64(uint32(v3))+12:], uint32(v2))
				v7 = v1 + v2
				v2 = i32(0)
			l33:
				{
					{
						t2 := int32(int8(m.memory[uint32(v1)]))
						v8 = t2
						if v8 <= i32(-1) {
							goto l4
						}
						v1 = v1 + i32(1)
						v8 = v8 & i32(255)
						goto l5
					}
				l4:
					t3 := int32(m.memory[int64(uint32(v1))+1])
					v9 = t3 & i32(63)
					v10 = v8 & i32(31)
					if uint32(v8) > uint32(i32(-33)) {
						goto l6
					}
					v8 = v10<<6 | v9
					v1 = v1 + i32(2)
					goto l5
				l6:
					t4 := int32(m.memory[int64(uint32(v1))+2])
					v9 = v9<<6 | t4&i32(63)
					if uint32(v8) >= uint32(i32(-16)) {
						goto l7
					}
					v8 = v9 | v10<<12
					v1 = v1 + i32(3)
					goto l5
				l7:
					t5 := int32(m.memory[int64(uint32(v1))+3])
					v8 = v9<<6 | t5&i32(63) | v10<<18&i32(0x1c0000)
					v1 = v1 + i32(4)
				}
			l5:
				{
					{
						{
							switch v8 + i32(-60) {
							default:
								if v8 == i32(124) {
									goto l12
								}
								fallthrough
							case 1:
								if uint32(v8) < uint32(i32(32)) {
									goto l13
								}
								if uint32(v8+i32(-127)) < uint32(i32(33)) {
									goto l13
								}
								var p6 int32
								if uint32(v8) < uint32(i32(128)) {
									p6 = 1
								}
								v11 = p6
								if v11 == 0 {
									goto l14
								}
								v10 = i32(1)
								goto l15
							case 0:
								{
									t7 := int32(load32(m.memory[int64(uint32(v3))+12:]))
									if uint32(t7-v2) > uint32(i32(2)) {
										goto l16
									}
									m.fn196(v3+i32(12), v2, i32(3), i32(1), i32(1))
									t8 := int32(load32(m.memory[int64(uint32(v3))+20:]))
									v2 = t8
								}
							l16:
								t9 := int32(load32(m.memory[int64(uint32(v3))+16:]))
								v4 = t9
								v8 = v4 + v2
								t10 := int32(load16(m.memory[int64(uint32(i32(0)))+1078807:]))
								store16(m.memory[uint32(v8):], uint16(t10))
								t11 := int32(m.memory[int64(uint32(i32(0)))+1078809])
								m.memory[int64(uint32(v8))+2] = byte(t11)
								goto l17
							case 2:
								{
									t12 := int32(load32(m.memory[int64(uint32(v3))+12:]))
									if uint32(t12-v2) > uint32(i32(2)) {
										goto l18
									}
									m.fn196(v3+i32(12), v2, i32(3), i32(1), i32(1))
									t13 := int32(load32(m.memory[int64(uint32(v3))+20:]))
									v2 = t13
								}
							l18:
								t14 := int32(load32(m.memory[int64(uint32(v3))+16:]))
								v4 = t14
								v8 = v4 + v2
								t15 := int32(load16(m.memory[int64(uint32(i32(0)))+1078810:]))
								store16(m.memory[uint32(v8):], uint16(t15))
								t16 := int32(m.memory[int64(uint32(i32(0)))+1078812])
								m.memory[int64(uint32(v8))+2] = byte(t16)
								goto l17
							}
						l12:
							{
								t17 := int32(load32(m.memory[int64(uint32(v3))+12:]))
								if uint32(t17-v2) > uint32(i32(2)) {
									goto l19
								}
								m.fn196(v3+i32(12), v2, i32(3), i32(1), i32(1))
								t18 := int32(load32(m.memory[int64(uint32(v3))+16:]))
								v4 = t18
								t19 := int32(load32(m.memory[int64(uint32(v3))+20:]))
								v2 = t19
							}
						l19:
							v8 = v4 + v2
							t20 := int32(m.memory[int64(uint32(i32(0)))+1078815])
							m.memory[int64(uint32(v8))+2] = byte(t20)
							t21 := int32(load16(m.memory[int64(uint32(i32(0)))+1078813:]))
							store16(m.memory[uint32(v8):], uint16(t21))
						}
					l17:
						v2 = v2 + i32(3)
						goto l20
					l14:
						if uint32(v8) >= uint32(i32(2048)) {
							goto l21
						}
						v10 = i32(2)
						goto l15
					l21:
						p22 := i32(4)
						if uint32(v8) < uint32(i32(65536)) {
							p22 = i32(3)
						}
						v10 = p22
					}
				l15:
					{
						t23 := int32(load32(m.memory[int64(uint32(v3))+12:]))
						if uint32(v10) <= uint32(t23-v2) {
							goto l22
						}
						m.fn196(v3+i32(12), v2, v10, i32(1), i32(1))
					}
				l22:
					t24 := int32(load32(m.memory[int64(uint32(v3))+16:]))
					v4 = t24
					v9 = v4 + v2
					if v11 != 0 {
						goto l23
					}
					v11 = v8&i32(63) | i32(-128)
					v12 = int32(uint32(v8) >> 6)
					if uint32(v8) >= uint32(i32(2048)) {
						v13 = int32(uint32(v8) >> 12)
						v12 = v12&i32(63) | i32(-128)
						if uint32(v8) > uint32(i32(0xffff)) {
							m.memory[int64(uint32(v9))+3] = byte(v11)
							m.memory[int64(uint32(v9))+2] = byte(v12)
							m.memory[int64(uint32(v9))+1] = byte(v13&i32(63) | i32(-128))
							m.memory[uint32(v9)] = byte(int32(uint32(v8)>>18) | i32(-16))
							v2 = v10 + v2
							goto l20
						}
						m.memory[int64(uint32(v9))+2] = byte(v11)
						m.memory[int64(uint32(v9))+1] = byte(v12)
						m.memory[uint32(v9)] = byte(v13 | i32(224))
						v2 = v10 + v2
						goto l20
					}
					m.memory[int64(uint32(v9))+1] = byte(v11)
					m.memory[uint32(v9)] = byte(v12 | i32(192))
					v2 = v10 + v2
					goto l20
				}
			l13:
				store32(m.memory[int64(uint32(v3))+24:], uint32(i32(0)))
				v13 = v6
				if uint32(v8) < uint32(i32(128)) {
					goto l26
				}
				m.memory[int64(uint32(v3))+25] = byte(v8&i32(63) | i32(128))
				v8 = i32(194)
				v13 = v5
			l26:
				m.memory[int64(uint32(v3))+24] = byte(v8)
				v8 = v3 + i32(24)
			l31:
				{
					t25 := int32(m.memory[uint32(v8)])
					v9 = t25
					{
						t26 := int32(load32(m.memory[int64(uint32(v3))+12:]))
						t27 := v2
						v10 = t26
						if t27 != v10 {
							goto l27
						}
						m.fn196(v3+i32(12), v2, i32(1), i32(1), i32(1))
						t28 := int32(load32(m.memory[int64(uint32(v3))+12:]))
						v10 = t28
					}
				l27:
					t29 := int32(load32(m.memory[int64(uint32(v3))+16:]))
					v4 = t29
					m.memory[uint32(v4+v2)] = byte(i32(37))
					t30 := v3
					v11 = v2 + i32(1)
					store32(m.memory[int64(uint32(t30))+20:], uint32(v11))
					t31 := int32(m.memory[int64(uint32(int32(uint32(v9)>>4)))+1122976])
					v12 = t31
					{
						if v11 != v10 {
							goto l28
						}
						m.fn196(v3+i32(12), v10, i32(1), i32(1), i32(1))
						t32 := int32(load32(m.memory[int64(uint32(v3))+12:]))
						v10 = t32
						t33 := int32(load32(m.memory[int64(uint32(v3))+16:]))
						v4 = t33
					}
				l28:
					m.memory[uint32(v4+v2+i32(1))] = byte(v12)
					t34 := v3
					v11 = v2 + i32(2)
					store32(m.memory[int64(uint32(t34))+20:], uint32(v11))
					t35 := int32(m.memory[int64(uint32(v9&i32(15)))+1122976])
					v9 = t35
					{
						if v11 != v10 {
							goto l29
						}
						m.fn196(v3+i32(12), v10, i32(1), i32(1), i32(1))
						t36 := int32(load32(m.memory[int64(uint32(v3))+16:]))
						v4 = t36
					}
				l29:
					m.memory[uint32(v4+v2+i32(2))] = byte(v9)
					t37 := v3
					v2 = v2 + i32(3)
					store32(m.memory[int64(uint32(t37))+20:], uint32(v2))
					v8 = v8 + i32(1)
					if v8 == v13 {
						goto l30
					}
					goto l31
				}
			l23:
				m.memory[uint32(v9)] = byte(v8)
				v2 = v10 + v2
			l20:
				store32(m.memory[int64(uint32(v3))+20:], uint32(v2))
			l30:
				if v1 == v7 {
					goto l32
				}
				goto l33
			}
			store32(m.memory[int64(uint32(v3))+20:], uint32(i32(0)))
			store64(m.memory[int64(uint32(v3))+12:], uint64(i64(0x100000000)))
			goto l2
		l32:
			t38 := int32(load32(m.memory[int64(uint32(v3))+16:]))
			v4 = t38
			v10 = v4 + v2
		l50:
			{
				{
					t39 := int32(int8(m.memory[uint32(v4)]))
					v2 = t39
					if v2 <= i32(-1) {
						goto l34
					}
					v4 = v4 + i32(1)
					v2 = v2 & i32(255)
					goto l35
				}
			l34:
				t40 := int32(m.memory[int64(uint32(v4))+1])
				v8 = t40 & i32(63)
				v9 = v2 & i32(31)
				if uint32(v2) > uint32(i32(-33)) {
					goto l36
				}
				v2 = v9<<6 | v8
				v4 = v4 + i32(2)
				goto l35
			l36:
				t41 := int32(m.memory[int64(uint32(v4))+2])
				v8 = v8<<6 | t41&i32(63)
				if uint32(v2) >= uint32(i32(-16)) {
					goto l37
				}
				v2 = v8 | v9<<12
				v4 = v4 + i32(3)
				goto l35
			l37:
				t42 := int32(m.memory[int64(uint32(v4))+3])
				v2 = v8<<6 | t42&i32(63) | v9<<18&i32(0x1c0000)
				v4 = v4 + i32(4)
			}
		l35:
			{
				v8 = v2 + i32(-9)
				if uint32(v8) > uint32(i32(23)) {
					goto l38
				}
				if i32_shl(i32(1), v8)&i32(8388639) != 0 {
					goto l39
				}
			l38:
				if uint32(v2) < uint32(i32(133)) {
					if v2&i32(254) != i32(40) {
						goto l42
					}
					goto l39
				}
				v8 = int32(uint32(v2) >> 8)
				switch v8 + i32(-22) {
				case 0:
					goto l41
				case 1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25:
					goto l42
				case 26:
					if v2 == i32(12288) {
						goto l39
					}
					goto l42
				case 10:
					t43 := int32(m.memory[int64(uint32(v2&i32(255)))+1139588])
					if t43&i32(2) == 0 {
						goto l42
					}
					goto l39
				default:
					if v8 != 0 {
						goto l42
					}
					t44 := int32(m.memory[int64(uint32(v2&i32(255)))+1139588])
					if t44&i32(1) == 0 {
						goto l42
					}
					goto l39
				}
			l41:
				if v2 != i32(5760) {
					goto l42
				}
			l39:
				store64(m.memory[int64(uint32(v3))+24:], uint64(int64(uint32(i32(17)))<<32|int64(uint32(v3+i32(12)))))
				m.fn13(v0, i32(1066475), v3+i32(24))
				t45 := int32(load32(m.memory[int64(uint32(v3))+12:]))
				v2 = t45
				if v2 == 0 {
					goto l46
				}
				{
					t46 := int32(load32(m.memory[int64(uint32(v3))+16:]))
					v8 = t46
					t47 := int32(load32(m.memory[uint32(v8+i32(-4)):]))
					v4 = t47
					v10 = v4 & i32(-8)
					t48 := v10
					v4 = v4 & i32(3)
					p49 := i32(8)
					if v4 != 0 {
						p49 = i32(4)
					}
					if uint32(t48) < uint32(p49+v2) {
						m.fn3(i32(1274224), i32(46), i32(1274272))
						panic("unreachable")
					}
					if v4 == 0 {
						goto l48
					}
					if uint32(v10) > uint32(v2+i32(39)) {
						m.fn3(i32(1274288), i32(46), i32(1274336))
						panic("unreachable")
					}
				l48:
					m.fn1(v8)
					goto l46
				}
			}
		l42:
			if v4 != v10 {
				goto l50
			}
		}
	l2:
		t50 := int32(load32(m.memory[int64(uint32(v3))+20:]))
		store32(m.memory[int64(uint32(v0))+8:], uint32(t50))
		t51 := int64(load64(m.memory[int64(uint32(v3))+12:]))
		store64(m.memory[uint32(v0):], uint64(t51))
	}
l46:
	m.g0 = v3 + i32(32)
}
func (m *Module) fn792(v0, v1, v2, v3 int32) {
	var v4 int64
	var v5, v6 int32
	var v7 int64
	var v8, v9 int32
	var v10 int64
	var v11 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		if t0 == 0 {
			goto l0
		}
		t1 := int64(load64(m.memory[int64(uint32(v1))+16:]))
		t2 := int64(load64(m.memory[int64(uint32(v1))+24:]))
		t3 := m.fn249(t1, t2, v2, v3)
		v4 = t3
		t4 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		v5 = t4
		v6 = v5 & int32(v4)
		v7 = int64(uint64(v4)>>25) & i64(127) * i64(72340172838076673)
		t5 := int32(load32(m.memory[uint32(v1):]))
		v8 = t5
		v9 = i32(0)
	l6:
		{
			{
				t6 := int64(load64(m.memory[uint32(v8+v6):]))
				v10 = t6
				v4 = v10 ^ v7
				v4 = (v4 ^ i64(-1)) & (v4 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
				if v4 == 0 {
					goto l1
				}
			l4:
				{
					t7 := v3
					v1 = v8 + (i32(0)-(int32(uint32(int64(bits.TrailingZeros64(uint64(v4))))>>3)+v6)&v5)*i32(28)
					t8 := int32(load32(m.memory[uint32(v1+i32(-20)):]))
					if t7 != t8 {
						goto l2
					}
					t9 := int32(load32(m.memory[uint32(v1+i32(-24)):]))
					t10 := m.fn973(v2, t9, v3)
					if t10 == 0 {
						t12 := int32(m.memory[uint32(v1+i32(-4))])
						if t12 != i32(1) {
							goto l0
						}
						t13 := int32(load32(m.memory[uint32(v1+i32(-8)):]))
						v11 = t13
						t14 := int32(load32(m.memory[uint32(v1+i32(-12)):]))
						v1 = t14
						goto l5
					}
				}
			l2:
				v4 = (v4 + i64(-1)) & v4
				if !(v4 == 0) {
					goto l4
				}
			}
		l1:
			v1 = i32(0)
			if !(v10&(v10<<1)&i64(-0x7f7f7f7f7f7f7f80) == 0) {
				goto l5
			}
			t11 := v6
			v9 = v9 + i32(8)
			v6 = (t11 + v9) & v5
			goto l6
		}
	}
l0:
	v1 = i32(0)
l5:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v11))
	store32(m.memory[uint32(v0):], uint32(v1))
}
func (m *Module) fn793(v0, v1, v2 int32) {
	var v3, v4, v5, v6, v7, v8, v9, v10 int32
	var v11 int64
	t0 := m.g0
	v3 = t0 - i32(64)
	m.g0 = v3
	{
		if v1 <= i32(-1) {
			m.fn11()
			panic("unreachable")
		}
		t1 := m.fn7(v1)
		v4 = t1
		if v4 == 0 {
			m.fn12(i32(1), v1)
			panic("unreachable")
		}
		v5 = v1 & i32(3)
		v6 = i32(0)
		if uint32(v1) < uint32(i32(4)) {
			goto l2
		}
		v7 = v1 & i32(0x7ffffffc)
		v6 = i32(0)
	l3:
		{
			v8 = v4 + v6
			t2 := v8
			v9 = v0 + v6
			t3 := int32(m.memory[uint32(v9)])
			v10 = t3
			p4 := v10
			if v10 == i32(10) {
				p4 = i32(32)
			}
			m.memory[uint32(t2)] = byte(p4)
			t5 := int32(m.memory[uint32(v9+i32(1))])
			t6 := v8 + i32(1)
			v10 = t5
			p7 := v10
			if v10 == i32(10) {
				p7 = i32(32)
			}
			m.memory[uint32(t6)] = byte(p7)
			t8 := int32(m.memory[uint32(v9+i32(2))])
			t9 := v8 + i32(2)
			v10 = t8
			p10 := v10
			if v10 == i32(10) {
				p10 = i32(32)
			}
			m.memory[uint32(t9)] = byte(p10)
			t11 := int32(m.memory[uint32(v9+i32(3))])
			t12 := v8 + i32(3)
			v8 = t11
			p13 := v8
			if v8 == i32(10) {
				p13 = i32(32)
			}
			m.memory[uint32(t12)] = byte(p13)
			t14 := v7
			v6 = v6 + i32(4)
			if t14 != v6 {
				goto l3
			}
		}
		if v5 == 0 {
			goto l4
		}
	l2:
		v8 = v0 + v6
		v6 = v4 + v6
	l5:
		{
			t15 := int32(m.memory[uint32(v8)])
			t16 := v6
			v9 = t15
			p17 := v9
			if v9 == i32(10) {
				p17 = i32(32)
			}
			m.memory[uint32(t16)] = byte(p17)
			v8 = v8 + i32(1)
			v6 = v6 + i32(1)
			v5 = v5 + i32(-1)
			if v5 != 0 {
				goto l5
			}
		}
	l4:
		store32(m.memory[int64(uint32(v3))+16:], uint32(v1))
		store32(m.memory[int64(uint32(v3))+12:], uint32(v4))
		store32(m.memory[int64(uint32(v3))+8:], uint32(v1))
		v8 = i32(1)
		m.fn784(v3+i32(20), v4, v1, i32(1))
		v6 = i32(1089817)
		{
			t18 := int32(m.memory[uint32(v4)])
			if t18 == i32(96) {
				goto l6
			}
			t19 := int32(m.memory[uint32(v4+v1+i32(-1))])
			var p20 int32
			if t19 == i32(96) {
				p20 = 1
			}
			v8 = p20
			p21 := i32(1)
			if v8 != 0 {
				p21 = i32(1089817)
			}
			v6 = p21
		}
	l6:
		store32(m.memory[int64(uint32(v3))+36:], uint32(v8))
		store32(m.memory[int64(uint32(v3))+32:], uint32(v6))
		t22 := v3
		v11 = int64(uint32(i32(17))) << 32
		store64(m.memory[int64(uint32(t22))+56:], uint64(v11|int64(uint32(v3+i32(8)))))
		store64(m.memory[int64(uint32(v3))+48:], uint64(int64(uint32(i32(1)))<<32|int64(uint32(v3+i32(32)))))
		store64(m.memory[int64(uint32(v3))+40:], uint64(v11|int64(uint32(v3+i32(20)))))
		_ = m.fn45(v2, i32(1078840), i32(1078864), v3+i32(40))
		{
			t24 := int32(load32(m.memory[int64(uint32(v3))+20:]))
			v8 = t24
			if v8 == 0 {
				goto l7
			}
			t25 := int32(load32(m.memory[int64(uint32(v3))+24:]))
			v9 = t25
			t26 := int32(load32(m.memory[uint32(v9+i32(-4)):]))
			v6 = t26
			v5 = v6 & i32(-8)
			t27 := v5
			v6 = v6 & i32(3)
			p28 := i32(8)
			if v6 != 0 {
				p28 = i32(4)
			}
			if uint32(t27) < uint32(p28+v8) {
				m.fn3(i32(1274224), i32(46), i32(1274272))
				panic("unreachable")
			}
			if v6 == 0 {
				goto l9
			}
			if uint32(v5) > uint32(v8+i32(39)) {
				m.fn3(i32(1274288), i32(46), i32(1274336))
				panic("unreachable")
			}
		l9:
			m.fn1(v9)
		}
	l7:
		{
			t29 := int32(load32(m.memory[int64(uint32(v3))+8:]))
			v8 = t29
			if v8 == 0 {
				goto l11
			}
			t30 := int32(load32(m.memory[int64(uint32(v3))+12:]))
			v9 = t30
			t31 := int32(load32(m.memory[uint32(v9+i32(-4)):]))
			v6 = t31
			v5 = v6 & i32(-8)
			t32 := v5
			v6 = v6 & i32(3)
			p33 := i32(8)
			if v6 != 0 {
				p33 = i32(4)
			}
			if uint32(t32) < uint32(p33+v8) {
				m.fn3(i32(1274224), i32(46), i32(1274272))
				panic("unreachable")
			}
			if v6 == 0 {
				goto l13
			}
			if uint32(v5) > uint32(v8+i32(39)) {
				m.fn3(i32(1274288), i32(46), i32(1274336))
				panic("unreachable")
			}
		l13:
			m.fn1(v9)
		}
	l11:
		m.g0 = v3 + i32(64)
		return
	}
}
func (m *Module) fn794(v0, v1, v2 int32) {
	var v3, v4, v5, v6 int32
	var v7, v8, v9, v10, v11 int64
	var v12, v13, v14, v15 int32
	var v16 int64
	var v17, v18, v19 int32
	t0 := m.g0
	v3 = t0 - i32(160)
	m.g0 = v3
	{
		t1 := int32(load32(m.memory[uint32(v0):]))
		v4 = t1
		switch v4 >> 31 & (v4 + i32(-0x7fffffff)) {
		case 6:
			goto l6
		case 2:
			t2 := int32(load32(m.memory[int64(uint32(v0))+24:]))
			v4 = t2
			if v4 == 0 {
				goto l6
			}
			t3 := int32(load32(m.memory[int64(uint32(v0))+20:]))
			v5 = t3
			v6 = v5 + v4*i32(28)
			v7 = int64(uint32(i32(17))) << 32
			v8 = v7 | int64(uint32(v3+i32(136)))
			v9 = v7 | int64(uint32(v3+i32(80)))
			v10 = v7 | int64(uint32(v3+i32(120)))
			t4 := int64(load64(m.memory[int64(uint32(v0))+8:]))
			v11 = t4
			v7 = i64(0)
			t5 := int32(m.memory[int64(uint32(v0))+28])
			v12 = t5
			v13 = v12 & i32(255)
		l41:
			store32(m.memory[int64(uint32(v3))+76:], uint32(i32(0)))
			store64(m.memory[int64(uint32(v3))+68:], uint64(i64(0x400000000)))
			{
				t6 := int32(load32(m.memory[int64(uint32(v5))+8:]))
				v0 = t6
				if v0 == 0 {
					goto l7
				}
				v4 = v0 << 5
				t7 := int32(load32(m.memory[int64(uint32(v5))+4:]))
				v0 = t7
			l8:
				m.fn794(v0, v1, v3+i32(68))
				v0 = v0 + i32(32)
				v4 = v4 + i32(-32)
				if v4 != 0 {
					goto l8
				}
			}
		l7:
			{
				{
					t8 := int32(load32(m.memory[int64(uint32(v5))+12:]))
					if t8 == i32(-1) {
						goto l9
					}
					t9 := int32(load32(m.memory[int64(uint32(v5))+16:]))
					t10 := int32(load32(m.memory[int64(uint32(v5))+20:]))
					m.fn783(v3+i32(120), t9, t10, i32(2))
					store64(m.memory[int64(uint32(v3))+96:], uint64(v10))
					m.fn13(v3+i32(136), i32(1067927), v3+i32(96))
					{
						t11 := int32(load32(m.memory[int64(uint32(v3))+120:]))
						v0 = t11
						if v0 == 0 {
							goto l10
						}
						t12 := int32(load32(m.memory[int64(uint32(v3))+124:]))
						v14 = t12
						t13 := int32(load32(m.memory[uint32(v14+i32(-4)):]))
						v4 = t13
						v15 = v4 & i32(-8)
						t14 := v15
						v4 = v4 & i32(3)
						p15 := i32(8)
						if v4 != 0 {
							p15 = i32(4)
						}
						if uint32(t14) < uint32(p15+v0) {
							m.fn3(i32(1274224), i32(46), i32(1274272))
							panic("unreachable")
						}
						if v4 == 0 {
							goto l12
						}
						if uint32(v15) > uint32(v0+i32(39)) {
							m.fn3(i32(1274288), i32(46), i32(1274336))
							panic("unreachable")
						}
					l12:
						m.fn1(v14)
					}
				l10:
					t16 := int64(load64(m.memory[int64(uint32(v3))+136:]))
					store64(m.memory[int64(uint32(v3))+80:], uint64(t16))
					t17 := int32(load32(m.memory[int64(uint32(v3))+144:]))
					store32(m.memory[int64(uint32(v3))+88:], uint32(t17))
					goto l14
				}
			l9:
				{
					if v13 != 0 {
						t19 := v3 + i32(120)
						t20 := v12
						v16 = v11 + v7
						p21 := v16
						if uint64(v16) < uint64(v11) {
							p21 = i64(-1)
						}
						m.fn301(t19, t20, p21)
						store64(m.memory[int64(uint32(v3))+96:], uint64(v10))
						m.fn13(v3+i32(136), i32(1067927), v3+i32(96))
						{
							t22 := int32(load32(m.memory[int64(uint32(v3))+120:]))
							v0 = t22
							if v0 == 0 {
								goto l17
							}
							t23 := int32(load32(m.memory[int64(uint32(v3))+124:]))
							v14 = t23
							t24 := int32(load32(m.memory[uint32(v14+i32(-4)):]))
							v4 = t24
							v15 = v4 & i32(-8)
							t25 := v15
							v4 = v4 & i32(3)
							p26 := i32(8)
							if v4 != 0 {
								p26 = i32(4)
							}
							if uint32(t25) < uint32(p26+v0) {
								m.fn3(i32(1274224), i32(46), i32(1274272))
								panic("unreachable")
							}
							if v4 == 0 {
								goto l19
							}
							if uint32(v15) > uint32(v0+i32(39)) {
								m.fn3(i32(1274288), i32(46), i32(1274336))
								panic("unreachable")
							}
						l19:
							m.fn1(v14)
						}
					l17:
						t27 := int64(load64(m.memory[int64(uint32(v3))+136:]))
						store64(m.memory[int64(uint32(v3))+80:], uint64(t27))
						t28 := int32(load32(m.memory[int64(uint32(v3))+144:]))
						store32(m.memory[int64(uint32(v3))+88:], uint32(t28))
						goto l14
					}
					t18 := m.fn7(i32(4))
					v0 = t18
					if v0 != 0 {
						goto l16
					}
					m.fn12(i32(1), i32(4))
					panic("unreachable")
				}
			l16:
				store32(m.memory[uint32(v0):], uint32(i32(547520738)))
				store32(m.memory[int64(uint32(v3))+88:], uint32(i32(4)))
				store32(m.memory[int64(uint32(v3))+84:], uint32(v0))
				store32(m.memory[int64(uint32(v3))+80:], uint32(i32(4)))
			l14:
				{
					t29 := int32(load32(m.memory[int64(uint32(v3))+76:]))
					v4 = t29
					if v4 == 0 {
						goto l21
					}
					t30 := int32(load32(m.memory[int64(uint32(v3))+72:]))
					m.fn202(v3+i32(120), t30, v4, i32(1089817), i32(1))
					t31 := int32(load32(m.memory[int64(uint32(v3))+128:]))
					store32(m.memory[int64(uint32(v3))+144:], uint32(t31))
					t32 := int64(load64(m.memory[int64(uint32(v3))+120:]))
					store64(m.memory[int64(uint32(v3))+136:], uint64(t32))
					store64(m.memory[int64(uint32(v3))+128:], uint64(v8))
					store64(m.memory[int64(uint32(v3))+120:], uint64(v9))
					m.fn13(v3+i32(108), i32(0x10006a), v3+i32(120))
					{
						t33 := int32(load32(m.memory[int64(uint32(v3))+136:]))
						v0 = t33
						if v0 == 0 {
							goto l22
						}
						t34 := int32(load32(m.memory[int64(uint32(v3))+140:]))
						v15 = t34
						t35 := int32(load32(m.memory[uint32(v15+i32(-4)):]))
						v14 = t35
						v17 = v14 & i32(-8)
						t36 := v17
						v14 = v14 & i32(3)
						p37 := i32(8)
						if v14 != 0 {
							p37 = i32(4)
						}
						if uint32(t36) < uint32(p37+v0) {
							m.fn3(i32(1274224), i32(46), i32(1274272))
							panic("unreachable")
						}
						if v14 == 0 {
							goto l24
						}
						if uint32(v17) > uint32(v0+i32(39)) {
							m.fn3(i32(1274288), i32(46), i32(1274336))
							panic("unreachable")
						}
					l24:
						m.fn1(v15)
					}
				l22:
					{
						t38 := int32(load32(m.memory[int64(uint32(v2))+8:]))
						v0 = t38
						t39 := int32(load32(m.memory[uint32(v2):]))
						if v0 != t39 {
							goto l26
						}
						m.fn201(v2)
					}
				l26:
					store32(m.memory[int64(uint32(v2))+8:], uint32(v0+i32(1)))
					t40 := int32(load32(m.memory[int64(uint32(v2))+4:]))
					v0 = t40 + v0*i32(12)
					t41 := int64(load64(m.memory[int64(uint32(v3))+108:]))
					store64(m.memory[uint32(v0):], uint64(t41))
					t42 := int32(load32(m.memory[int64(uint32(v3))+116:]))
					store32(m.memory[int64(uint32(v0))+8:], uint32(t42))
				}
			l21:
				{
					t43 := int32(load32(m.memory[int64(uint32(v3))+80:]))
					v0 = t43
					if v0 == 0 {
						goto l27
					}
					t44 := int32(load32(m.memory[int64(uint32(v3))+84:]))
					v15 = t44
					t45 := int32(load32(m.memory[uint32(v15+i32(-4)):]))
					v14 = t45
					v17 = v14 & i32(-8)
					t46 := v17
					v14 = v14 & i32(3)
					p47 := i32(8)
					if v14 != 0 {
						p47 = i32(4)
					}
					if uint32(t46) < uint32(p47+v0) {
						m.fn3(i32(1274224), i32(46), i32(1274272))
						panic("unreachable")
					}
					if v14 == 0 {
						goto l29
					}
					if uint32(v17) > uint32(v0+i32(39)) {
						m.fn3(i32(1274288), i32(46), i32(1274336))
						panic("unreachable")
					}
				l29:
					m.fn1(v15)
				}
			l27:
				t48 := int32(load32(m.memory[int64(uint32(v3))+72:]))
				v18 = t48
				if v4 == 0 {
					goto l31
				}
				v0 = v18
			l36:
				{
					t49 := int32(load32(m.memory[uint32(v0):]))
					v14 = t49
					if v14 == 0 {
						goto l32
					}
					t50 := int32(load32(m.memory[uint32(v0+i32(4)):]))
					v17 = t50
					t51 := int32(load32(m.memory[uint32(v17+i32(-4)):]))
					v15 = t51
					v19 = v15 & i32(-8)
					t52 := v19
					v15 = v15 & i32(3)
					p53 := i32(8)
					if v15 != 0 {
						p53 = i32(4)
					}
					if uint32(t52) < uint32(p53+v14) {
						m.fn3(i32(1274224), i32(46), i32(1274272))
						panic("unreachable")
					}
					if v15 == 0 {
						goto l34
					}
					if uint32(v19) > uint32(v14+i32(39)) {
						m.fn3(i32(1274288), i32(46), i32(1274336))
						panic("unreachable")
					}
				l34:
					m.fn1(v17)
				}
			l32:
				v0 = v0 + i32(12)
				v4 = v4 + i32(-1)
				if v4 != 0 {
					goto l36
				}
			l31:
				{
					t54 := int32(load32(m.memory[int64(uint32(v3))+68:]))
					v0 = t54
					if v0 == 0 {
						goto l37
					}
					t55 := int32(load32(m.memory[uint32(v18+i32(-4)):]))
					v4 = t55
					v14 = v4 & i32(-8)
					t56 := v14
					v4 = v4 & i32(3)
					p57 := i32(8)
					if v4 != 0 {
						p57 = i32(4)
					}
					v0 = v0 * i32(12)
					if uint32(t56) < uint32(p57+v0) {
						m.fn3(i32(1274224), i32(46), i32(1274272))
						panic("unreachable")
					}
					if v4 == 0 {
						goto l39
					}
					if uint32(v14) > uint32(v0+i32(39)) {
						m.fn3(i32(1274288), i32(46), i32(1274336))
						panic("unreachable")
					}
				l39:
					m.fn1(v18)
				}
			l37:
				v7 = v7 + i64(1)
				v5 = v5 + i32(28)
				if v5 != v6 {
					goto l41
				}
				goto l6
			}
		case 1:
			t58 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			t59 := int32(load32(m.memory[int64(uint32(v0))+12:]))
			m.fn780(v3+i32(56), t58, t59, i32(2), i32(0), v1)
			t60 := int32(load32(m.memory[int64(uint32(v3))+60:]))
			t61 := v3 + i32(24)
			v0 = t60
			t62 := int32(load32(m.memory[int64(uint32(v3))+64:]))
			m.fn143(t61, v0, t62)
			{
				t63 := int32(load32(m.memory[int64(uint32(v3))+28:]))
				if t63 == 0 {
					t69 := int32(load32(m.memory[int64(uint32(v3))+56:]))
					v4 = t69
					if v4 == 0 {
						goto l6
					}
					{
						t70 := int32(load32(m.memory[uint32(v0+i32(-4)):]))
						v14 = t70
						v15 = v14 & i32(-8)
						t71 := v15
						v14 = v14 & i32(3)
						p72 := i32(8)
						if v14 != 0 {
							p72 = i32(4)
						}
						if uint32(t71) < uint32(p72+v4) {
							m.fn3(i32(1274224), i32(46), i32(1274272))
							panic("unreachable")
						}
						if v14 == 0 {
							goto l45
						}
						if uint32(v15) > uint32(v4+i32(39)) {
							m.fn3(i32(1274288), i32(46), i32(1274336))
							panic("unreachable")
						}
					l45:
						m.fn1(v0)
						goto l6
					}
				}
				{
					t64 := int32(load32(m.memory[int64(uint32(v2))+8:]))
					v0 = t64
					t65 := int32(load32(m.memory[uint32(v2):]))
					if v0 != t65 {
						goto l43
					}
					m.fn201(v2)
				}
			l43:
				store32(m.memory[int64(uint32(v2))+8:], uint32(v0+i32(1)))
				t66 := int32(load32(m.memory[int64(uint32(v2))+4:]))
				v0 = t66 + v0*i32(12)
				t67 := int64(load64(m.memory[int64(uint32(v3))+56:]))
				store64(m.memory[uint32(v0):], uint64(t67))
				t68 := int32(load32(m.memory[int64(uint32(v3))+64:]))
				store32(m.memory[int64(uint32(v0))+8:], uint32(t68))
				goto l6
			}
		default:
			t73 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t74 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			m.fn780(v3+i32(120), t73, t74, i32(2), i32(0), v1)
			t75 := int32(load32(m.memory[int64(uint32(v3))+124:]))
			t76 := v3 + i32(16)
			v0 = t75
			t77 := int32(load32(m.memory[int64(uint32(v3))+128:]))
			t78 := v0
			v4 = t77
			m.fn143(t76, t78, v4)
			{
				t79 := int32(load32(m.memory[int64(uint32(v3))+20:]))
				if t79 == 0 {
					goto l47
				}
				m.fn143(v3+i32(8), v0, v4)
				t80 := int64(load64(m.memory[int64(uint32(v3))+8:]))
				store64(m.memory[int64(uint32(v3))+80:], uint64(t80))
				store64(m.memory[int64(uint32(v3))+136:], uint64(int64(uint32(i32(1)))<<32|int64(uint32(v3+i32(80)))))
				m.fn13(v3+i32(44), i32(1066491), v3+i32(136))
				{
					t81 := int32(load32(m.memory[int64(uint32(v2))+8:]))
					v4 = t81
					t82 := int32(load32(m.memory[uint32(v2):]))
					if v4 != t82 {
						goto l48
					}
					m.fn201(v2)
				}
			l48:
				store32(m.memory[int64(uint32(v2))+8:], uint32(v4+i32(1)))
				t83 := int32(load32(m.memory[int64(uint32(v2))+4:]))
				v4 = t83 + v4*i32(12)
				t84 := int64(load64(m.memory[int64(uint32(v3))+44:]))
				store64(m.memory[uint32(v4):], uint64(t84))
				t85 := int32(load32(m.memory[int64(uint32(v3))+52:]))
				store32(m.memory[int64(uint32(v4))+8:], uint32(t85))
			}
		l47:
			t86 := int32(load32(m.memory[int64(uint32(v3))+120:]))
			v4 = t86
			if v4 == 0 {
				goto l6
			}
			{
				t87 := int32(load32(m.memory[uint32(v0+i32(-4)):]))
				v14 = t87
				v15 = v14 & i32(-8)
				t88 := v15
				v14 = v14 & i32(3)
				p89 := i32(8)
				if v14 != 0 {
					p89 = i32(4)
				}
				if uint32(t88) < uint32(p89+v4) {
					m.fn3(i32(1274224), i32(46), i32(1274272))
					panic("unreachable")
				}
				if v14 == 0 {
					goto l50
				}
				if uint32(v15) > uint32(v4+i32(39)) {
					m.fn3(i32(1274288), i32(46), i32(1274336))
					panic("unreachable")
				}
			l50:
				m.fn1(v0)
				goto l6
			}
		case 5:
			t90 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			t91 := int32(load32(m.memory[int64(uint32(v0))+12:]))
			m.fn143(v3+i32(32), t90, t91)
			t92 := int32(load32(m.memory[int64(uint32(v3))+36:]))
			v0 = t92
			if v0 == 0 {
				goto l6
			}
			t93 := int32(load32(m.memory[int64(uint32(v3))+32:]))
			v4 = t93
			store32(m.memory[int64(uint32(v3))+144:], uint32(i32(0)))
			store64(m.memory[int64(uint32(v3))+136:], uint64(i64(0x100000000)))
			m.fn793(v4, v0, v3+i32(136))
			t94 := int32(load32(m.memory[int64(uint32(v3))+144:]))
			store32(m.memory[int64(uint32(v3))+128:], uint32(t94))
			t95 := int64(load64(m.memory[int64(uint32(v3))+136:]))
			store64(m.memory[int64(uint32(v3))+120:], uint64(t95))
			{
				t96 := int32(load32(m.memory[int64(uint32(v2))+8:]))
				v0 = t96
				t97 := int32(load32(m.memory[uint32(v2):]))
				if v0 != t97 {
					goto l52
				}
				m.fn201(v2)
			}
		l52:
			store32(m.memory[int64(uint32(v2))+8:], uint32(v0+i32(1)))
			t98 := int32(load32(m.memory[int64(uint32(v2))+4:]))
			v0 = t98 + v0*i32(12)
			t99 := int64(load64(m.memory[int64(uint32(v3))+120:]))
			store64(m.memory[uint32(v0):], uint64(t99))
			t100 := int32(load32(m.memory[int64(uint32(v3))+128:]))
			store32(m.memory[int64(uint32(v0))+8:], uint32(t100))
			goto l6
		case 4:
			t101 := int32(load32(m.memory[int64(uint32(v0))+12:]))
			v4 = t101
			if v4 == 0 {
				goto l6
			}
			v4 = v4 << 5
			t102 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			v0 = t102
		l53:
			m.fn794(v0, v1, v2)
			v0 = v0 + i32(32)
			v4 = v4 + i32(-32)
			if v4 != 0 {
				goto l53
			}
			goto l6
		case 3:
			t103 := int32(load32(m.memory[int64(uint32(v0))+12:]))
			v4 = t103
			if v4 == 0 {
				goto l6
			}
			t104 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			v5 = t104
			v13 = v5 + v4*i32(12)
		l74:
			{
				t105 := int32(load32(m.memory[int64(uint32(v5))+8:]))
				v15 = t105
				if uint32(v15) >= uint32(i32(0xaaaaaab)) {
					m.fn11()
					panic("unreachable")
				}
				v14 = v15 * i32(12)
				{
					if v15 != 0 {
						goto l55
					}
					v6 = i32(0)
					v18 = i32(4)
					goto l56
				l55:
					t106 := int32(load32(m.memory[int64(uint32(v5))+4:]))
					v4 = t106
					t107 := m.fn7(v14)
					v18 = t107
					v0 = v18
					v17 = v15
					if v18 == 0 {
						m.fn12(i32(4), v14)
						panic("unreachable")
					}
				l60:
					{
						{
							t108 := int32(load32(m.memory[uint32(v4):]))
							if t108 != i32(-1) {
								goto l58
							}
							store32(m.memory[int64(uint32(v3))+156:], uint32(i32(0)))
							store64(m.memory[int64(uint32(v3))+148:], uint64(i64(0x100000000)))
							goto l59
						}
					l58:
						m.fn787(v3+i32(148), v4, v1)
					l59:
						t109 := int32(load32(m.memory[int64(uint32(v3))+156:]))
						store32(m.memory[int64(uint32(v0))+8:], uint32(t109))
						t110 := int64(load64(m.memory[int64(uint32(v3))+148:]))
						store64(m.memory[uint32(v0):], uint64(t110))
						v4 = v4 + i32(20)
						v0 = v0 + i32(12)
						v17 = v17 + i32(-1)
						if v17 != 0 {
							goto l60
						}
					}
					v6 = v15
				}
			l56:
				v5 = v5 + i32(12)
				v0 = v18
				{
				l62:
					{
						if v14 == 0 {
							goto l61
						}
						v14 = v14 + i32(-12)
						v4 = v0 + i32(8)
						v0 = v0 + i32(12)
						t111 := int32(load32(m.memory[uint32(v4):]))
						if t111 == 0 {
							goto l62
						}
					}
					m.fn202(v3+i32(120), v18, v15, i32(1078804), i32(3))
					{
						t112 := int32(load32(m.memory[int64(uint32(v2))+8:]))
						v0 = t112
						t113 := int32(load32(m.memory[uint32(v2):]))
						if v0 != t113 {
							goto l63
						}
						m.fn201(v2)
					}
				l63:
					store32(m.memory[int64(uint32(v2))+8:], uint32(v0+i32(1)))
					t114 := int32(load32(m.memory[int64(uint32(v2))+4:]))
					v0 = t114 + v0*i32(12)
					t115 := int64(load64(m.memory[int64(uint32(v3))+120:]))
					store64(m.memory[uint32(v0):], uint64(t115))
					t116 := int32(load32(m.memory[int64(uint32(v3))+128:]))
					store32(m.memory[int64(uint32(v0))+8:], uint32(t116))
				}
			l61:
				if v15 == 0 {
					goto l64
				}
				v0 = v18
			l69:
				{
					t117 := int32(load32(m.memory[uint32(v0):]))
					v4 = t117
					if v4 == 0 {
						goto l65
					}
					t118 := int32(load32(m.memory[uint32(v0+i32(4)):]))
					v17 = t118
					t119 := int32(load32(m.memory[uint32(v17+i32(-4)):]))
					v14 = t119
					v19 = v14 & i32(-8)
					t120 := v19
					v14 = v14 & i32(3)
					p121 := i32(8)
					if v14 != 0 {
						p121 = i32(4)
					}
					if uint32(t120) < uint32(p121+v4) {
						m.fn3(i32(1274224), i32(46), i32(1274272))
						panic("unreachable")
					}
					if v14 == 0 {
						goto l67
					}
					if uint32(v19) > uint32(v4+i32(39)) {
						m.fn3(i32(1274288), i32(46), i32(1274336))
						panic("unreachable")
					}
				l67:
					m.fn1(v17)
				}
			l65:
				v0 = v0 + i32(12)
				v15 = v15 + i32(-1)
				if v15 != 0 {
					goto l69
				}
			l64:
				{
					if v6 == 0 {
						goto l70
					}
					t122 := int32(load32(m.memory[uint32(v18+i32(-4)):]))
					v0 = t122
					v4 = v0 & i32(-8)
					t123 := v4
					v0 = v0 & i32(3)
					p124 := i32(8)
					if v0 != 0 {
						p124 = i32(4)
					}
					v14 = v6 * i32(12)
					if uint32(t123) < uint32(p124+v14) {
						m.fn3(i32(1274224), i32(46), i32(1274272))
						panic("unreachable")
					}
					if v0 == 0 {
						goto l72
					}
					if uint32(v4) > uint32(v14+i32(39)) {
						m.fn3(i32(1274288), i32(46), i32(1274336))
						panic("unreachable")
					}
				l72:
					m.fn1(v18)
				}
			l70:
				if v5 != v13 {
					goto l74
				}
				goto l6
			}
		}
	}
l6:
	m.g0 = v3 + i32(160)
}
func (m *Module) fn795(v0, v1 int32) {
	var v2, v3, v4, v5, v6, v7 int32
	t0 := m.g0
	v2 = t0 - i32(32)
	m.g0 = v2
	v3 = i32(0)
	{
	l11:
		{
			{
				t1 := int32(m.memory[int64(uint32(v1))+37])
				if t1 == 0 {
					goto l0
				}
				goto l1
			}
		l0:
			t2 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			v4 = t2
			m.fn198(v2+i32(20), v1)
			{
				{
					t3 := int32(load32(m.memory[int64(uint32(v2))+20:]))
					if t3 != i32(1) {
						goto l2
					}
					t4 := int32(load32(m.memory[int64(uint32(v1))+28:]))
					v5 = t4
					t5 := int32(load32(m.memory[int64(uint32(v2))+28:]))
					t6 := v1
					v6 = t5
					store32(m.memory[int64(uint32(t6))+28:], uint32(v6))
					v4 = v4 + v5
					v5 = v6 - v5
					goto l3
				}
			l2:
				t7 := int32(m.memory[int64(uint32(v1))+37])
				if t7 != 0 {
					goto l1
				}
				m.memory[int64(uint32(v1))+37] = byte(i32(1))
				{
					{
						t8 := int32(m.memory[int64(uint32(v1))+36])
						if t8 != i32(1) {
							goto l5
						}
						t9 := int32(load32(m.memory[int64(uint32(v1))+32:]))
						v5 = t9
						t10 := int32(load32(m.memory[int64(uint32(v1))+28:]))
						v4 = t10
						goto l6
					}
				l5:
					t11 := int32(load32(m.memory[int64(uint32(v1))+32:]))
					v5 = t11
					t12 := int32(load32(m.memory[int64(uint32(v1))+28:]))
					t13 := v5
					v4 = t12
					if t13 == v4 {
						goto l1
					}
				}
			l6:
				v5 = v5 - v4
				t14 := int32(load32(m.memory[int64(uint32(v1))+4:]))
				v4 = t14 + v4
			}
		l3:
			{
				if v5 == 0 {
					goto l7
				}
				t15 := v4
				v6 = v5 + i32(-1)
				t16 := int32(m.memory[uint32(t15+v6)])
				if t16 != i32(10) {
					goto l7
				}
				v5 = v5 + i32(-2)
				{
					if v6 != 0 {
						goto l8
					}
					v7 = i32(0)
					goto l9
				l8:
					t17 := int32(m.memory[uint32(v4+v5)])
					p18 := i32(0)
					if t17&i32(255) == i32(13) {
						p18 = v4
					}
					v7 = p18
				}
			l9:
				p19 := v6
				if v7 != 0 {
					p19 = v5
				}
				v5 = p19
				p20 := v4
				if v7 != 0 {
					p20 = v7
				}
				v4 = p20
			}
		l7:
			if v4 != 0 {
				goto l10
			}
			goto l1
		l10:
			m.fn143(v2+i32(8), v4, v5)
			t21 := int32(load32(m.memory[int64(uint32(v2))+12:]))
			if t21 == 0 {
				goto l11
			}
		}
		m.fn143(v2, v4, v5)
		t22 := int32(load32(m.memory[int64(uint32(v2))+4:]))
		v1 = t22
		t23 := int32(load32(m.memory[uint32(v2):]))
		v3 = t23
		goto l1
	}
l1:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
	store32(m.memory[uint32(v0):], uint32(v3))
	m.g0 = v2 + i32(32)
}
func (m *Module) fn796(v0 int32) int32 {
	var v1, v2 int32
	v1 = i32(0)
	v2 = int32(uint32(v0) >> 8)
	switch v2 + i32(-22) {
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25:
		goto l1
	case 0:
		var p0 int32
		if v0 == i32(5760) {
			p0 = 1
		}
		v1 = p0
		goto l1
	case 26:
		var p1 int32
		if v0 == i32(12288) {
			p1 = 1
		}
		v1 = p1
		goto l1
	default:
		if v2 != 0 {
			goto l1
		}
		t2 := int32(m.memory[int64(uint32(v0&i32(255)))+1139588])
		v1 = t2
		goto l1
	case 10:
		t3 := int32(m.memory[int64(uint32(v0&i32(255)))+1139588])
		v1 = int32(uint32(t3&i32(2)) >> 1)
	}
l1:
	return v1 & i32(1)
}
func (m *Module) fn797(v0 int32) int32 {
	var v1, v2, v3 int32
	{
		t0 := int32(load32(m.memory[uint32(v0):]))
		if t0 != i32(-1) {
			goto l0
		}
		{
			t1 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			v1 = t1
			if v1 != 0 {
				goto l1
			}
			store64(m.memory[uint32(v0):], uint64(i64(0x100000000)))
			return v0
		}
	l1:
		t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		v2 = t2
		t3 := m.fn7(v1)
		v3 = t3
		if v3 == 0 {
			m.fn12(i32(1), v1)
			panic("unreachable")
		}
		if v1 == 0 {
			goto l3
		}
		memory_copy(m.memory, uint32(v3), uint32(v2), uint32(v1))
	l3:
		store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
		store32(m.memory[uint32(v0):], uint32(v1))
		if v1 == i32(-1) {
			m.fn3(i32(1274396), i32(40), i32(1074144))
			panic("unreachable")
		}
	}
l0:
	return v0
}
func (m *Module) fn798(v0, v1 int32) int32 {
	var v2, v3 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	{
		{
			t1 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			v3 = t1
			if v3&i32(0x2000000) != 0 {
				t3 := int32(load16(m.memory[uint32(v0):]))
				v3 = t3
				v0 = i32(5)
			l3:
				{
					t4 := int32(m.memory[uint32(v3&i32(15)+i32(1099240))])
					m.memory[uint32(v2+i32(8)+v0+i32(-2))] = byte(t4)
					v0 = v0 + i32(-1)
					v3 = int32(uint32(v3)>>4) & i32(0xfff)
					if v3 != 0 {
						goto l3
					}
				}
				t5 := m.fn679(v1, i32(1), i32(1122974), i32(2), v2+i32(8)+v0+i32(-1), i32(5)-v0)
				v0 = t5
				goto l2
			}
			if v3&i32(0x4000000) != 0 {
				goto l1
			}
			t2 := m.fn165(v0, v1)
			v0 = t2
			goto l2
		}
	l1:
		t6 := int32(load16(m.memory[uint32(v0):]))
		v3 = t6
		v0 = i32(5)
	l4:
		{
			t7 := int32(m.memory[uint32(v3&i32(15)+i32(1122976))])
			m.memory[uint32(v2+i32(12)+v0+i32(-2))] = byte(t7)
			v0 = v0 + i32(-1)
			v3 = int32(uint32(v3)>>4) & i32(0xfff)
			if v3 != 0 {
				goto l4
			}
		}
		t8 := m.fn679(v1, i32(1), i32(1122974), i32(2), v2+i32(12)+v0+i32(-1), i32(5)-v0)
		v0 = t8
	}
l2:
	m.g0 = v2 + i32(16)
	return v0
}
func (m *Module) fn799(v0 int32) {
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
func (m *Module) fn800(v0, v1, v2, v3, v4, v5 int32) {
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
			t0 := m.fn21(v2, v5*v1, v4, v3)
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
		t1 := m.fn19(v3, v4)
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
func (m *Module) fn801(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	v3 = i32(1)
	{
		t1 := int32(load32(m.memory[uint32(v1):]))
		v4 = t1
		t2 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t3 := v4
		v5 = t2
		t4 := int32(load32(m.memory[int64(uint32(v5))+12:]))
		v6 = t4
		t5 := m.t0[uint(v6)].(func(int32, int32, int32) int32)(t3, i32(1276972), i32(15))
		if t5 != 0 {
			goto l0
		}
		{
			{
				t6 := int32(m.memory[int64(uint32(v1))+10])
				if t6&i32(128) != 0 {
					goto l1
				}
				v3 = i32(1)
				t7 := m.t0[uint(v6)].(func(int32, int32, int32) int32)(v4, i32(1099467), i32(1))
				if t7 != 0 {
					goto l0
				}
				t8 := int32(m.memory[uint32(v0)])
				t9 := v4
				v1 = t8 << 2
				t10 := int32(load32(m.memory[int64(uint32(v1))+1290744:]))
				t11 := int32(load32(m.memory[int64(uint32(v1))+1290720:]))
				t12 := m.t0[uint(v6)].(func(int32, int32, int32) int32)(t9, t10, t11)
				if t12 == 0 {
					goto l2
				}
				goto l0
			}
		l1:
			t13 := m.t0[uint(v6)].(func(int32, int32, int32) int32)(v4, i32(1099468), i32(2))
			if t13 != 0 {
				goto l0
			}
			store32(m.memory[int64(uint32(v2))+4:], uint32(v5))
			store32(m.memory[uint32(v2):], uint32(v4))
			v3 = i32(1)
			m.memory[int64(uint32(v2))+15] = byte(i32(1))
			t14 := int32(m.memory[uint32(v0)])
			v1 = t14 << 2
			t15 := int32(load32(m.memory[int64(uint32(v1))+1290768:]))
			v0 = t15
			t16 := int32(load32(m.memory[int64(uint32(v1))+1290792:]))
			v1 = t16
			store32(m.memory[int64(uint32(v2))+8:], uint32(v2+i32(15)))
			t17 := m.fn336(v2, v1, v0)
			if t17 != 0 {
				goto l0
			}
			t18 := m.fn336(v2, i32(1099465), i32(2))
			if t18 != 0 {
				goto l0
			}
		}
	l2:
		t19 := m.t0[uint(v6)].(func(int32, int32, int32) int32)(v4, i32(1272712), i32(1))
		v3 = t19
	}
l0:
	m.g0 = v2 + i32(16)
	return v3
}
func (m *Module) fn802(v0, v1 int32) {
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
				m.fn11()
				panic("unreachable")
			}
			if v2 != 0 {
				t6 := int32(load32(m.memory[int64(uint32(v1))+4:]))
				v3 = t6
				t7 := m.fn7(v2)
				v1 = t7
				if v1 != 0 {
					goto l13
				}
				m.fn12(i32(1), v2)
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
func (m *Module) fn803(v0 int32) {
	var v1 int32
	var v2 int64
	t0 := m.g0
	v1 = t0 - i32(32)
	m.g0 = v1
	store32(m.memory[int64(uint32(v1))+8:], uint32(v0))
	store32(m.memory[int64(uint32(v1))+12:], uint32(i32(4)))
	t1 := v1
	v2 = int64(uint32(i32(3))) << 32
	store64(m.memory[int64(uint32(t1))+24:], uint64(v2|int64(uint32(v1+i32(12)))))
	store64(m.memory[int64(uint32(v1))+16:], uint64(v2|int64(uint32(v1+i32(8)))))
	m.fn27(i32(1066791), v1+i32(16), i32(1090396))
	panic("unreachable")
}
func (m *Module) fn804(v0, v1 int32) {
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
					m.fn640(v2+i32(4), v3, i32(1), i32(1), i32(1))
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
				m.fn640(v1, v5, v0, i32(1), i32(1))
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
					m.fn640(v1, v3, v6, i32(1), i32(1))
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
					m.fn3(i32(1274224), i32(46), i32(1274272))
					panic("unreachable")
				}
				if v1 == 0 {
					goto l20
				}
				if uint32(v3) > uint32(v0+i32(39)) {
					m.fn3(i32(1274288), i32(46), i32(1274336))
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
			m.fn640(v1, v3, i32(1), i32(1), i32(1))
		}
	l22:
		store32(m.memory[int64(uint32(v1))+8:], uint32(v3+i32(1)))
		t28 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		m.memory[uint32(t28+v3)] = byte(v0 + i32(65))
	}
l18:
	m.g0 = v2 + i32(16)
}
func (m *Module) fn805(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v1):]))
	t1 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	t2 := int32(load32(m.memory[int64(uint32(t1))+12:]))
	t3 := m.t0[uint(t2)].(func(int32, int32, int32) int32)(t0, i32(1275037), i32(5))
	return t3
}
func (m *Module) fn806(v0, v1, v2 int32) int32 {
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
			m.fn640(v0, v3, v2, i32(1), i32(1))
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
