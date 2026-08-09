package core

import (
	"math/bits"
)

func (m *Module) fn582(v0, v1, v2 int32) {
	m.fn621(v0, v1, v2, i64(0))
}
func (m *Module) fn583(v0, v1 int32) {
	var v2 int32
	var v3 int64
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	store32(m.memory[int64(uint32(v2))+4:], uint32(i32(0)))
	m.fn595(v2+i32(8), v1, v2+i32(4), i32(4))
	{
		{
			t1 := int32(m.memory[int64(uint32(v2))+8])
			if t1 == i32(255) {
				goto l0
			}
			t2 := int64(load64(m.memory[int64(uint32(v2))+8:]))
			v3 = t2
			if v3&i64(255) == i64(255) {
				goto l0
			}
			store64(m.memory[uint32(v0):], uint64(v3))
			goto l1
		}
	l0:
		t3 := int32(load32(m.memory[int64(uint32(v2))+4:]))
		store32(m.memory[int64(uint32(v0))+4:], uint32(t3))
		m.memory[uint32(v0)] = byte(i32(255))
	}
l1:
	m.g0 = v2 + i32(16)
}
func (m *Module) fn584(v0, v1 int32) {
	var v2 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		v2 = t0
		t1 := int32(load32(m.memory[uint32(v0):]))
		if v2 != t1 {
			goto l0
		}
		m.fn618(v0)
	}
l0:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v2+i32(1)))
	t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	store32(m.memory[uint32(t2+v2<<2):], uint32(v1))
}
func (m *Module) fn585(v0, v1 int32) int32 {
	var v2 int32
	var v3 int64
	var v4 int32
	var v5 int64
	var v6 int32
	var v7 int64
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
		if t0 == 0 {
			goto l0
		}
		t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		v2 = t1
		t2 := m.fn619(v1)
		t3 := v2
		v3 = t2
		v4 = t3 & int32(v3)
		v5 = int64(uint64(v3)>>25) & i64(127) * i64(72340172838076673)
		t4 := int32(load32(m.memory[uint32(v0):]))
		v0 = t4
		v6 = i32(0)
	l3:
		{
			t5 := int64(load64(m.memory[uint32(v0+v4):]))
			v7 = t5
			v3 = v7 ^ v5
			v3 = (v3 ^ i64(-1)) & (v3 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
		l4:
			{
				if v3 == 0 {
					if !(v7&(v7<<1)&i64(-0x7f7f7f7f7f7f7f80) == 0) {
						goto l0
					}
					t7 := v4
					v6 = v6 + i32(8)
					v4 = (t7 + v6) & v2
					goto l3
				}
				t6 := int32(load32(m.memory[uint32(v0-(int32(uint32(int64(bits.TrailingZeros64(uint64(v3))))>>3)+v4)&v2<<2+i32(-4)):]))
				if v1 != t6 {
					v3 = (v3 + i64(-1)) & v3
					goto l4
				}
				return i32(1)
			}
		}
	}
l0:
	return i32(0)
}
func (m *Module) fn586(v0, v1 int32) {
	var v2 int32
	var v3 int64
	var v4, v5, v6, v7, v8 int32
	t0 := m.g0
	v2 = t0 - i32(32)
	m.g0 = v2
	store32(m.memory[int64(uint32(v2))+16:], uint32(v1))
	t1 := m.fn619(v1)
	v3 = t1
	store32(m.memory[int64(uint32(v2))+20:], uint32(v2+i32(16)))
	{
		t2 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		if t2 != 0 {
			goto l0
		}
		_ = m.fn620(v0, v0+i32(16))
	}
l0:
	store32(m.memory[int64(uint32(v2))+24:], uint32(v2+i32(20)))
	store32(m.memory[int64(uint32(v2))+28:], uint32(v0))
	t4 := int32(load32(m.memory[uint32(v0):]))
	t5 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	m.fn69(v2+i32(8), t4, t5, v3, v2+i32(24), i32(4))
	{
		t6 := int32(load32(m.memory[int64(uint32(v2))+8:]))
		if t6 != i32(1) {
			goto l1
		}
		t7 := int32(load32(m.memory[uint32(v0):]))
		v4 = t7
		t8 := int32(load32(m.memory[int64(uint32(v2))+12:]))
		t9 := v4
		v5 = t8
		v6 = t9 + v5
		t10 := int32(m.memory[uint32(v6)])
		v7 = t10
		t11 := v6
		v8 = int32(uint32(int32(v3)) >> 25)
		m.memory[uint32(t11)] = byte(v8)
		t12 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		m.memory[uint32(v4+t12&(v5+i32(-8))+i32(8))] = byte(v8)
		t13 := int32(load32(m.memory[int64(uint32(v0))+12:]))
		store32(m.memory[int64(uint32(v0))+12:], uint32(t13+i32(1)))
		t14 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		store32(m.memory[int64(uint32(v0))+8:], uint32(t14-v7&i32(1)))
		store32(m.memory[uint32(v4-v5<<2+i32(-4)):], uint32(v1))
	}
l1:
	m.g0 = v2 + i32(32)
}
func (m *Module) fn587(v0, v1 int32) {
	var v2 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	{
		if v1 == 0 {
			goto l0
		}
		m.fn39(v2+i32(4), i32(4), i32(8), v1+i32(1))
		t1 := int32(load32(m.memory[int64(uint32(v2))+12:]))
		t2 := int32(load32(m.memory[int64(uint32(v2))+4:]))
		t3 := int32(load32(m.memory[int64(uint32(v2))+8:]))
		m.fn40(v0-t1, t2, t3)
	}
l0:
	m.g0 = v2 + i32(16)
}
func (m *Module) fn588(v0 int32) {
	t0 := int32(load32(m.memory[int64(uint32(v0))+24:]))
	t1 := int32(load32(m.memory[int64(uint32(v0))+28:]))
	m.fn449(t0, t1)
	t2 := int32(load32(m.memory[int64(uint32(v0))+36:]))
	t3 := int32(load32(m.memory[int64(uint32(v0))+40:]))
	m.fn449(t2, t3)
	t4 := int32(load32(m.memory[int64(uint32(v0))+48:]))
	t5 := int32(load32(m.memory[int64(uint32(v0))+52:]))
	m.fn449(t4, t5)
	t6 := int32(load32(m.memory[int64(uint32(v0))+60:]))
	t7 := int32(load32(m.memory[int64(uint32(v0))+64:]))
	m.fn449(t6, t7)
}
func (m *Module) fn589(v0, v1 int32) int32 {
	t0 := m.fn590(v0, v1, i32(0))
	return t0
}
func (m *Module) fn590(v0, v1, v2 int32) int32 {
	if uint32(v2) < uint32(v1) {
		return v0 + v2*i32(80)
	}
	m.fn158(v2, v1, i32(1075616))
	panic("unreachable")
}
func (m *Module) fn591(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v1):]))
	t1 := int32(m.memory[uint32(v0)])
	v0 = t1 << 2
	t2 := int32(load32(m.memory[int64(uint32(v0))+1301476:]))
	t3 := int32(load32(m.memory[int64(uint32(v0))+1301460:]))
	t4 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	t5 := int32(load32(m.memory[int64(uint32(t4))+12:]))
	t6 := m.t0[uint(t5)].(func(int32, int32, int32) int32)(t0, t2, t3)
	return t6
}
func (m *Module) fn592(v0, v1, v2, v3 int32) int32 {
	var v4, v5, v6 int32
	t0 := m.g0
	v4 = t0 - i32(32)
	m.g0 = v4
	{
		{
			t1 := m.fn1630(v0, v1)
			if t1 == 0 {
				goto l0
			}
			t2 := m.fn1630(v2, v3)
			if t2 == 0 {
				goto l0
			}
			{
				if v1 == v3 {
					v5 = i32(0)
				l3:
					{
						if v1 == 0 {
							goto l2
						}
						t5 := int32(m.memory[uint32(v2)])
						v3 = t5
						t6 := int32(m.memory[uint32(v0)])
						v6 = t6
						v1 = v1 + i32(-1)
						v2 = v2 + i32(1)
						v0 = v0 + i32(1)
						t8 := v6
						p7 := i32(0)
						if uint32((v6+i32(-97))&i32(255)) < uint32(i32(26)) {
							p7 = i32(32)
						}
						v6 = t8 ^ p7
						t10 := v6 & i32(255)
						t11 := v3
						p9 := i32(0)
						if uint32((v3+i32(-97))&i32(255)) < uint32(i32(26)) {
							p9 = i32(32)
						}
						v3 = t11 ^ p9
						if t10 == v3&i32(255) {
							goto l3
						}
					}
					v0 = v6 & i32(255)
					t12 := v0
					v2 = v3 & i32(255)
					var p13 int32
					if uint32(t12) > uint32(v2) {
						p13 = 1
					}
					var p14 int32
					if uint32(v0) < uint32(v2) {
						p14 = 1
					}
					v5 = p13 - p14
					goto l2
				}
				var p3 int32
				if uint32(v1) > uint32(v3) {
					p3 = 1
				}
				var p4 int32
				if uint32(v1) < uint32(v3) {
					p4 = 1
				}
				v5 = p3 - p4
				goto l2
			}
		}
	l0:
		store16(m.memory[int64(uint32(v4))+20:], uint16(i32(0)))
		store32(m.memory[int64(uint32(v4))+12:], uint32(v0))
		t15 := v4
		v6 = v0 + v1
		store32(m.memory[int64(uint32(t15))+16:], uint32(v6))
		t16 := m.fn606(v4 + i32(12))
		v1 = t16
		store16(m.memory[int64(uint32(v4))+20:], uint16(i32(0)))
		store32(m.memory[int64(uint32(v4))+12:], uint32(v2))
		t17 := v4
		v5 = v2 + v3
		store32(m.memory[int64(uint32(t17))+16:], uint32(v5))
		{
			t18 := m.fn606(v4 + i32(12))
			t19 := v1
			v3 = t18
			if t19 == v3 {
				goto l4
			}
			var p20 int32
			if uint32(v1) > uint32(v3) {
				p20 = 1
			}
			var p21 int32
			if uint32(v1) < uint32(v3) {
				p21 = 1
			}
			v5 = p20 - p21
			goto l2
		}
	l4:
		store32(m.memory[int64(uint32(v4))+16:], uint32(v5))
		store32(m.memory[int64(uint32(v4))+12:], uint32(v2))
		store32(m.memory[int64(uint32(v4))+28:], uint32(v6))
		store32(m.memory[int64(uint32(v4))+24:], uint32(v0))
		{
		l7:
			{
				t22 := m.fn48(v4 + i32(24))
				v0 = t22
				if v0 == i32(-1) {
					goto l5
				}
				t23 := m.fn1625(v0)
				v0 = t23
				{
					t24 := m.fn1628(v4 + i32(12))
					v2 = t24
					if v2 != i32(-1) {
						goto l6
					}
					v5 = i32(1)
					goto l2
				}
			l6:
				if v0 == v2 {
					goto l7
				}
			}
			var p25 int32
			if uint32(v0) > uint32(v2) {
				p25 = 1
			}
			var p26 int32
			if uint32(v0) < uint32(v2) {
				p26 = 1
			}
			v5 = p25 - p26
			goto l2
		}
	l5:
		t27 := m.fn1628(v4 + i32(12))
		p28 := i32(0)
		if t27 != i32(-1) {
			p28 = i32(-1)
		}
		v5 = p28
	}
l2:
	m.g0 = v4 + i32(32)
	return v5
}
func (m *Module) fn593(v0, v1, v2 int32) {
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
	m.memory[int64(uint32(v0))+4] = byte(v2)
	store32(m.memory[uint32(v0):], uint32(v1))
}
func (m *Module) fn594(v0, v1 int32) {
	var v2 int32
	var v3 int64
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	store16(m.memory[int64(uint32(v2))+6:], uint16(i32(0)))
	m.fn595(v2+i32(8), v1, v2+i32(6), i32(2))
	{
		{
			t1 := int32(m.memory[int64(uint32(v2))+8])
			if t1 == i32(255) {
				goto l0
			}
			t2 := int64(load64(m.memory[int64(uint32(v2))+8:]))
			v3 = t2
			if v3&i64(255) == i64(255) {
				goto l0
			}
			store64(m.memory[uint32(v0):], uint64(v3))
			goto l1
		}
	l0:
		t3 := int32(load16(m.memory[int64(uint32(v2))+6:]))
		store16(m.memory[int64(uint32(v0))+2:], uint16(t3))
		m.memory[uint32(v0)] = byte(i32(255))
	}
l1:
	m.g0 = v2 + i32(16)
}
func (m *Module) fn595(v0, v1, v2, v3 int32) {
	var v4, v5 int32
	t0 := m.g0
	v4 = t0 - i32(16)
	m.g0 = v4
l6:
	{
		if v3 == 0 {
			m.memory[uint32(v0)] = byte(i32(255))
			goto l3
		}
		m.fn623(v4+i32(8), v1, v2, v3)
		{
			t1 := int32(m.memory[int64(uint32(v4))+8])
			v5 = t1
			if v5 == i32(255) {
				t4 := int32(load32(m.memory[int64(uint32(v4))+12:]))
				v5 = t4
				if v5 == 0 {
					goto l4
				}
				if uint32(v3) < uint32(v5) {
					m.fn151(v5, v3, v3, i32(1072408))
					panic("unreachable")
				}
				v2 = v2 + v5
				v3 = v3 - v5
				goto l6
			}
			t2 := m.fn313(v4 + i32(8))
			if t2 != 0 {
				t6 := int32(load32(m.memory[int64(uint32(v4))+12:]))
				m.fn119(v5, t6)
				goto l6
			}
			t3 := int64(load64(m.memory[int64(uint32(v4))+8:]))
			store64(m.memory[uint32(v0):], uint64(t3))
			goto l3
		}
	l4:
		t5 := int64(load64(m.memory[int64(uint32(i32(0)))+1287056:]))
		store64(m.memory[uint32(v0):], uint64(t5))
	}
l3:
	m.g0 = v4 + i32(16)
}
func (m *Module) fn596(v0, v1 int32) int32 {
	var v2, v3, v4 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	v3 = i32(3)
	t1 := int32(m.memory[uint32(v0)])
	v0 = t1
	v4 = v0
	{
		if uint32(v0) < uint32(i32(10)) {
			goto l0
		}
		v3 = i32(1)
		t2 := int32(uint32(v0) / uint32(i32(100)))
		t3 := v2
		t4 := v0
		v4 = t2
		t5 := int32(load16(m.memory[int64(uint32((t4-v4*i32(100))&i32(255)<<1))+1109319:]))
		store16(m.memory[int64(uint32(t3))+14:], uint16(t5))
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
		t6 := v2 + i32(13)
		v3 = v3 + i32(-1)
		t7 := int32(m.memory[int64(uint32(v4<<1))+1109320])
		m.memory[uint32(t6+v3)] = byte(t7)
	}
l2:
	t8 := m.fn1638(v1, i32(1), i32(1), i32(0), v2+i32(13)+v3, i32(3)-v3)
	v3 = t8
	m.g0 = v2 + i32(16)
	return v3
}
func (m *Module) fn597(v0, v1 int32) {
	var v2, v3 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	{
		{
			t1 := int32(load32(m.memory[int64(uint32(v1))+12:]))
			v3 = t1
			if v3 != 0 {
				goto l0
			}
			v1 = i32(0)
			goto l1
		}
	l0:
		store32(m.memory[int64(uint32(v1))+12:], uint32(v3+i32(-1)))
		m.fn1624(v2+i32(8), v1)
		t2 := int32(load16(m.memory[int64(uint32(v2))+10:]))
		v3 = t2
		t3 := int32(load16(m.memory[int64(uint32(v2))+8:]))
		v1 = t3
	}
l1:
	store16(m.memory[int64(uint32(v0))+2:], uint16(v3))
	store16(m.memory[uint32(v0):], uint16(v1))
	m.g0 = v2 + i32(16)
}
func (m *Module) fn598(v0, v1 int32) {
	var v2, v3, v4, v5 int32
	{
		{
			t0 := int32(load32(m.memory[int64(uint32(v1))+12:]))
			v2 = t0
			if v2 != 0 {
				goto l0
			}
			v1 = i32(0)
			v2 = i32(0)
			goto l1
		}
	l0:
		t1 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t2 := int32(load32(m.memory[uint32(v1):]))
		t3 := v2
		v3 = t1 - t2
		t4 := int32(uint32(v3) / uint32(i32(3)))
		v4 = t4
		t5 := int32(load16(m.memory[int64(uint32(v1))+8:]))
		t6 := v4
		var p7 int32
		if t5 != i32(0) {
			p7 = 1
		}
		v5 = p7
		t8 := t6 + v5
		var p9 int32
		if v3-v4*i32(3) != i32(0) {
			p9 = 1
		}
		v1 = t8 + p9
		p10 := v1
		if uint32(v2) < uint32(v1) {
			p10 = t3
		}
		v1 = p10
		v3 = v3 + v5
		p11 := v2
		if uint32(v3) < uint32(v2) {
			p11 = v3
		}
		v2 = p11
	}
l1:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v2))
	store32(m.memory[int64(uint32(v0))+4:], uint32(i32(1)))
	store32(m.memory[uint32(v0):], uint32(v1))
}
func (m *Module) fn599(v0, v1, v2, v3 int32) {
	var v4 int32
	t0 := m.g0
	v4 = t0 - i32(16)
	m.g0 = v4
	m.fn1342(v4+i32(4), v1, i32(0), v2, v3)
	t1 := int32(load32(m.memory[int64(uint32(v4))+8:]))
	v3 = t1
	{
		t2 := int32(load32(m.memory[int64(uint32(v4))+4:]))
		if t2 != i32(1) {
			goto l0
		}
		t3 := int32(load32(m.memory[int64(uint32(v4))+12:]))
		m.fn2(v3, t3)
		panic("unreachable")
	}
l0:
	t4 := int32(load32(m.memory[int64(uint32(v4))+12:]))
	store32(m.memory[int64(uint32(v0))+4:], uint32(t4))
	store32(m.memory[uint32(v0):], uint32(v3))
	m.g0 = v4 + i32(16)
}
func (m *Module) fn600(v0, v1, v2, v3, v4 int32) {
	var v5 int32
	t0 := m.g0
	v5 = t0 - i32(16)
	m.g0 = v5
	m.fn1629(v5+i32(8), v0, v1, v2, v3, v4)
	{
		t1 := int32(load32(m.memory[int64(uint32(v5))+8:]))
		v4 = t1
		if v4 == i32(-1) {
			goto l0
		}
		t2 := int32(load32(m.memory[int64(uint32(v5))+12:]))
		m.fn2(v4, t2)
		panic("unreachable")
	}
l0:
	m.g0 = v5 + i32(16)
}
func (m *Module) fn601(v0, v1, v2, v3, v4 int32) {
	var v5, v6, v7, v8, v9, v10, v11, v12, v13 int32
	var v14 int64
	var v15, v16 int32
	v5 = i32(1)
	v6 = i32(1)
	v7 = i32(0)
	v8 = i32(1)
	v9 = i32(0)
l4:
	v10 = v9 + v7
	if uint32(v10) >= uint32(v4) {
		m.fn158(v10, v4, i32(1109192))
		panic("unreachable")
	}
	{
		t0 := int32(m.memory[uint32(v3+v5)])
		v5 = t0 & i32(255)
		t1 := int32(m.memory[uint32(v3+v10)])
		t2 := v5
		v10 = t1
		if uint32(t2) < uint32(v10) {
			v6 = v6 + v7 + i32(1)
			v8 = v6 - v9
			v7 = i32(0)
			goto l3
		}
		if v5 == v10 {
			v5 = v7 + i32(1)
			t3 := v5
			var p4 int32
			if v5 == v8 {
				p4 = 1
			}
			v10 = p4
			p5 := t3
			if v10 != 0 {
				p5 = i32(0)
			}
			v7 = p5
			p6 := i32(0)
			if v10 != 0 {
				p6 = v5
			}
			v6 = p6 + v6
			goto l3
		}
		v8 = i32(1)
		v7 = i32(0)
		v9 = v6
		v6 = v6 + i32(1)
		goto l3
	}
l3:
	v5 = v6 + v7
	if uint32(v5) < uint32(v4) {
		goto l4
	}
	v5 = i32(1)
	v6 = i32(1)
	v7 = i32(0)
	v11 = i32(1)
	v12 = i32(0)
l9:
	{
		v10 = v12 + v7
		if uint32(v10) >= uint32(v4) {
			m.fn158(v10, v4, i32(1109192))
			panic("unreachable")
		}
		t7 := int32(m.memory[uint32(v3+v5)])
		v5 = t7 & i32(255)
		t8 := int32(m.memory[uint32(v3+v10)])
		t9 := v5
		v10 = t8
		if uint32(t9) > uint32(v10) {
			goto l6
		}
		if v5 == v10 {
			v5 = v7 + i32(1)
			t10 := v5
			var p11 int32
			if v5 == v11 {
				p11 = 1
			}
			v10 = p11
			p12 := t10
			if v10 != 0 {
				p12 = i32(0)
			}
			v7 = p12
			p13 := i32(0)
			if v10 != 0 {
				p13 = v5
			}
			v6 = p13 + v6
			goto l8
		}
		v11 = i32(1)
		v7 = i32(0)
		v12 = v6
		v6 = v6 + i32(1)
		goto l8
	}
l6:
	v6 = v6 + v7 + i32(1)
	v11 = v6 - v12
	v7 = i32(0)
l8:
	v5 = v6 + v7
	if uint32(v5) < uint32(v4) {
		goto l9
	}
	{
		t14 := v4
		t15 := v9
		t16 := v12
		var p17 int32
		if uint32(v9) > uint32(v12) {
			p17 = 1
		}
		v7 = p17
		p18 := t16
		if v7 != 0 {
			p18 = t15
		}
		v13 = p18
		if uint32(t14) < uint32(v13) {
			m.fn151(i32(0), v13, v4, i32(1109256))
			panic("unreachable")
		}
		p19 := v11
		if v7 != 0 {
			p19 = v8
		}
		v6 = p19
		v7 = v6 + v13
		if uint32(v7) < uint32(v6) {
			goto l11
		}
		if uint32(v7) > uint32(v4) {
			goto l11
		}
		{
			{
				t20 := m.fn1851(v3, v3+v6, v13)
				if t20 == 0 {
					goto l12
				}
				v14 = i64(0)
				v7 = v3
				v6 = v4
			l13:
				{
					t21 := int64(m.memory[uint32(v7)])
					v14 = i64_shl(i64(1), t21) | v14
					v7 = v7 + i32(1)
					v6 = v6 + i32(-1)
					if v6 != 0 {
						goto l13
					}
				}
				v7 = v4 - v13
				p22 := v13
				if uint32(v7) > uint32(v13) {
					p22 = v7
				}
				v6 = p22 + i32(1)
				v5 = i32(-1)
				v10 = v13
				v7 = i32(-1)
				goto l14
			}
		l12:
			v12 = v4 + i32(-1)
			v9 = i32(1)
			v7 = i32(0)
			v10 = i32(1)
			v11 = i32(0)
		l21:
			v5 = v10
			v15 = v5 + v7
			if uint32(v15) >= uint32(v4) {
				goto l15
			}
			v10 = v4 - v7 + (v5 ^ i32(-1))
			if uint32(v10) >= uint32(v4) {
				m.fn158(v10, v4, i32(1109208))
				panic("unreachable")
			}
			v8 = v12 - (v7 + v11)
			if uint32(v8) >= uint32(v4) {
				m.fn158(v8, v4, i32(1109224))
				panic("unreachable")
			}
			{
				{
					t23 := int32(m.memory[uint32(v3+v10)])
					v10 = t23 & i32(255)
					t24 := int32(m.memory[uint32(v3+v8)])
					t25 := v10
					v8 = t24
					if uint32(t25) < uint32(v8) {
						v10 = v15 + i32(1)
						v9 = v10 - v11
						v7 = i32(0)
						goto l20
					}
					if v10 == v8 {
						goto l19
					}
					v10 = v5 + i32(1)
					v7 = i32(0)
					v9 = i32(1)
					v11 = v5
					goto l20
				}
			l19:
				v10 = v7 + i32(1)
				t26 := v10
				var p27 int32
				if v10 == v9 {
					p27 = 1
				}
				v8 = p27
				p28 := t26
				if v8 != 0 {
					p28 = i32(0)
				}
				v7 = p28
				p29 := i32(0)
				if v8 != 0 {
					p29 = v10
				}
				v10 = p29 + v5
			}
		l20:
			if v9 != v6 {
				goto l21
			}
		l15:
			v9 = i32(1)
			v7 = i32(0)
			v10 = i32(1)
			v15 = i32(0)
		l28:
			v5 = v10
			v16 = v5 + v7
			if uint32(v16) >= uint32(v4) {
				goto l22
			}
			v10 = v4 - v7 + (v5 ^ i32(-1))
			if uint32(v10) >= uint32(v4) {
				m.fn158(v10, v4, i32(1109208))
				panic("unreachable")
			}
			v8 = v12 - (v7 + v15)
			if uint32(v8) >= uint32(v4) {
				m.fn158(v8, v4, i32(1109224))
				panic("unreachable")
			}
			{
				{
					t30 := int32(m.memory[uint32(v3+v10)])
					v10 = t30 & i32(255)
					t31 := int32(m.memory[uint32(v3+v8)])
					t32 := v10
					v8 = t31
					if uint32(t32) > uint32(v8) {
						v10 = v16 + i32(1)
						v9 = v10 - v15
						v7 = i32(0)
						goto l27
					}
					if v10 == v8 {
						goto l26
					}
					v10 = v5 + i32(1)
					v7 = i32(0)
					v9 = i32(1)
					v15 = v5
					goto l27
				}
			l26:
				v10 = v7 + i32(1)
				t33 := v10
				var p34 int32
				if v10 == v9 {
					p34 = 1
				}
				v8 = p34
				p35 := t33
				if v8 != 0 {
					p35 = i32(0)
				}
				v7 = p35
				p36 := i32(0)
				if v8 != 0 {
					p36 = v10
				}
				v10 = p36 + v5
			}
		l27:
			if v9 != v6 {
				goto l28
			}
		l22:
			t38 := v4
			p37 := v11
			if uint32(v15) > uint32(v11) {
				p37 = v15
			}
			v10 = t38 - p37
			v14 = i64(0)
			if v6 != 0 {
				goto l29
			}
			v6 = i32(0)
			v5 = i32(0)
			goto l30
		l29:
			v5 = i32(0)
			v7 = i32(0)
		l31:
			{
				t39 := int64(m.memory[uint32(v3+v7)])
				v14 = i64_shl(i64(1), t39) | v14
				t40 := v6
				v7 = v7 + i32(1)
				if t40 != v7 {
					goto l31
				}
			}
		l30:
			v7 = v4
		}
	l14:
		store32(m.memory[int64(uint32(v0))+60:], uint32(v4))
		store32(m.memory[int64(uint32(v0))+56:], uint32(v3))
		store32(m.memory[int64(uint32(v0))+52:], uint32(v2))
		store32(m.memory[int64(uint32(v0))+48:], uint32(v1))
		store32(m.memory[int64(uint32(v0))+40:], uint32(v7))
		store32(m.memory[int64(uint32(v0))+36:], uint32(v5))
		store32(m.memory[int64(uint32(v0))+32:], uint32(v2))
		store32(m.memory[int64(uint32(v0))+28:], uint32(i32(0)))
		store32(m.memory[int64(uint32(v0))+24:], uint32(v6))
		store32(m.memory[int64(uint32(v0))+20:], uint32(v10))
		store32(m.memory[int64(uint32(v0))+16:], uint32(v13))
		store64(m.memory[int64(uint32(v0))+8:], uint64(v14))
		store32(m.memory[uint32(v0):], uint32(i32(1)))
		return
	}
l11:
	m.fn151(v6, v7, v4, i32(1109240))
	panic("unreachable")
}
func (m *Module) fn602(v0, v1, v2, v3, v4, v5, v6 int32) {
	var v7, v8, v9, v10, v11 int32
	var v12 int64
	var v13, v14, v15, v16, v17, v18, v19, v20 int32
	v7 = v5 + i32(-1)
	t0 := int32(load32(m.memory[int64(uint32(v1))+16:]))
	t1 := v5
	v8 = t0
	v9 = t1 - v8
	t2 := int32(load32(m.memory[int64(uint32(v1))+28:]))
	v10 = t2
	t3 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	v11 = t3
	t4 := int64(load64(m.memory[uint32(v1):]))
	v12 = t4
	t5 := int32(load32(m.memory[int64(uint32(v1))+20:]))
	v13 = t5
l15:
	{
		p6 := v10
		if v6 != 0 {
			p6 = i32(0)
		}
		v14 = p6
		t8 := v11
		p7 := v11
		if uint32(v10) > uint32(v11) {
			p7 = v10
		}
		p9 := p7
		if v6 != 0 {
			p9 = t8
		}
		v15 = p9
		p10 := v5
		if uint32(v15) > uint32(v5) {
			p10 = v15
		}
		v16 = p10
	l10:
		v10 = v7 + v13
		if uint32(v10) < uint32(v3) {
			goto l0
		}
		store32(m.memory[int64(uint32(v1))+20:], uint32(v3))
		v10 = i32(0)
		goto l1
	l0:
		{
			t11 := int64(m.memory[uint32(v2+v10)])
			if i64_shr_u(v12, t11)&i64(1) == 0 {
				goto l2
			}
			v17 = v2 + v13
			v10 = v15
		l13:
			{
				if v16 != v10 {
					goto l3
				}
				v10 = v11
			l9:
				{
					if uint32(v14) < uint32(v10) {
						goto l4
					}
					t12 := v1
					v10 = v13 + v5
					store32(m.memory[int64(uint32(t12))+20:], uint32(v10))
					if v6 != 0 {
						goto l5
					}
					store32(m.memory[int64(uint32(v1))+28:], uint32(i32(0)))
				l5:
					store32(m.memory[int64(uint32(v0))+8:], uint32(v10))
					store32(m.memory[int64(uint32(v0))+4:], uint32(v13))
					v10 = i32(1)
					goto l1
				}
			l4:
				v10 = v10 + i32(-1)
				if uint32(v10) >= uint32(v5) {
					m.fn158(v10, v5, i32(1100784))
					panic("unreachable")
				}
				{
					v18 = v10 + v13
					if uint32(v18) >= uint32(v3) {
						goto l7
					}
					t13 := int32(m.memory[uint32(v4+v10)])
					t14 := int32(m.memory[uint32(v2+v18)])
					if t13 != t14 {
						t15 := v1
						v13 = v8 + v13
						store32(m.memory[int64(uint32(t15))+20:], uint32(v13))
						if v6 != 0 {
							goto l10
						}
						v10 = v9
						goto l11
					}
					goto l9
				}
			l7:
				m.fn158(v18, v3, i32(1100800))
				panic("unreachable")
			l3:
				v19 = v13 + v10
				if uint32(v19) >= uint32(v3) {
					t18 := v3
					v10 = v15 + v13
					p19 := v10
					if uint32(v3) > uint32(v10) {
						p19 = t18
					}
					m.fn158(p19, v3, i32(1100816))
					panic("unreachable")
				}
				v18 = v17 + v10
				v20 = v4 + v10
				v10 = v10 + i32(1)
				t16 := int32(m.memory[uint32(v20)])
				t17 := int32(m.memory[uint32(v18)])
				if t16 == t17 {
					goto l13
				}
			}
			v13 = v19 - v11 + i32(1)
			goto l14
		}
	l2:
		v13 = v13 + v5
	l14:
		store32(m.memory[int64(uint32(v1))+20:], uint32(v13))
		if v6 != 0 {
			goto l10
		}
		v10 = i32(0)
	l11:
		store32(m.memory[int64(uint32(v1))+28:], uint32(v10))
		goto l15
	l1:
	}
	store32(m.memory[uint32(v0):], uint32(v10))
}
func (m *Module) fn603(v0, v1 int32) int32 {
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
		t7 := m.fn110(v1, v2+i32(12), v0)
		v0 = t7
	}
l1:
	m.g0 = v2 + i32(16)
	return v0
}
func (m *Module) fn604(v0, v1, v2 int32) {
	if v2&i32(1) == 0 {
		goto l0
	}
	m.fn1623(v0, v1, int32(uint32(v2)>>1))
	return
l0:
	m.fn6(v0, v1, v2)
}
func (m *Module) fn605(v0, v1 int32) {
	var v2 int32
	t0 := m.fn1620()
	v2 = t0
	t1 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	store32(m.memory[int64(uint32(v2))+8:], uint32(t1))
	t2 := int64(load64(m.memory[uint32(v1):]))
	store64(m.memory[uint32(v2):], uint64(t2))
	m.fn343(v0, i32(20), v2, i32(1102000))
}
func (m *Module) fn606(v0 int32) int32 {
	var v1, v2 int32
	t0 := m.g0
	v1 = t0 - i32(16)
	m.g0 = v1
	v2 = i32(-1)
l0:
	{
		m.fn1624(v1+i32(8), v0)
		v2 = v2 + i32(1)
		t1 := int32(load16(m.memory[int64(uint32(v1))+8:]))
		if t1&i32(1) != 0 {
			goto l0
		}
	}
	m.g0 = v1 + i32(16)
	return v2
}
func (m *Module) fn607(v0, v1, v2, v3 int32) int32 {
	var v4 int32
	v4 = i32(1)
	{
		if v1 != v3 {
			goto l0
		}
		t0 := m.fn1851(v0, v2, v1)
		var p1 int32
		if t0 != i32(0) {
			p1 = 1
		}
		v4 = p1
	}
l0:
	return v4
}
func (m *Module) fn608(v0, v1, v2, v3 int32) {
	var v4 int32
	t0 := m.g0
	v4 = t0 - i32(16)
	m.g0 = v4
	if v0 != 0 {
		goto l0
	}
	v0 = i32(0)
	v3 = v4 + i32(12)
	goto l1
l0:
	store32(m.memory[int64(uint32(v4))+12:], uint32(v2))
	v0 = v0 * v3
	v3 = v4 + i32(8)
l1:
	store32(m.memory[uint32(v3):], uint32(v0))
	{
		t1 := int32(load32(m.memory[int64(uint32(v4))+12:]))
		v0 = t1
		if v0 == 0 {
			goto l2
		}
		t2 := int32(load32(m.memory[int64(uint32(v4))+8:]))
		v3 = t2
		if v3 == 0 {
			goto l2
		}
		m.fn10(v1, v3, v0)
	}
l2:
	m.g0 = v4 + i32(16)
}
func (m *Module) fn609(v0, v1 int32) {
	var v2 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	m.fn610(v2, v1)
	v1 = i32(1)
	{
		{
			t1 := int32(load32(m.memory[uint32(v2):]))
			if t1 != i32(1) {
				goto l0
			}
			t2 := int64(load64(m.memory[int64(uint32(v2))+4:]))
			store64(m.memory[int64(uint32(v0))+4:], uint64(t2))
			goto l1
		}
	l0:
		t3 := int64(load64(m.memory[int64(uint32(v2))+8:]))
		store64(m.memory[int64(uint32(v0))+8:], uint64(t3))
		v1 = i32(0)
	}
l1:
	store32(m.memory[uint32(v0):], uint32(v1))
	m.g0 = v2 + i32(16)
}
func (m *Module) fn610(v0, v1 int32) {
	var v2 int32
	var v3 int64
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	store64(m.memory[uint32(v2):], uint64(i64(0)))
	m.fn595(v2+i32(8), v1, v2, i32(8))
	{
		{
			t1 := int32(m.memory[int64(uint32(v2))+8])
			if t1 == i32(255) {
				goto l0
			}
			t2 := int64(load64(m.memory[int64(uint32(v2))+8:]))
			v3 = t2
			if v3&i64(255) == i64(255) {
				goto l0
			}
			store64(m.memory[int64(uint32(v0))+4:], uint64(v3))
			v1 = i32(1)
			goto l1
		}
	l0:
		t3 := int64(load64(m.memory[uint32(v2):]))
		store64(m.memory[int64(uint32(v0))+8:], uint64(t3))
		v1 = i32(0)
	}
l1:
	store32(m.memory[uint32(v0):], uint32(v1))
	m.g0 = v2 + i32(16)
}
func (m *Module) fn611(v0 int32) {
	var v1 int32
	t0 := m.g0
	v1 = t0 - i32(16)
	m.g0 = v1
	t1 := int32(load32(m.memory[uint32(v0):]))
	m.fn273(v1+i32(8), v0, t1, i32(1), i32(8), i32(80))
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
func (m *Module) fn612(v0, v1, v2, v3 int32) {
	var v4 int32
	t0 := m.g0
	v4 = t0 - i32(48)
	m.g0 = v4
	store32(m.memory[int64(uint32(v4))+12:], uint32(v3))
	{
		if uint32(v3) >= uint32(v2) {
			goto l0
		}
		t1 := m.fn622(v1, v2, v3, i32(1075484))
		t2 := int32(load32(m.memory[uint32(t1):]))
		t3 := v4
		v3 = t2
		store32(m.memory[int64(uint32(t3))+44:], uint32(v3))
		if v3 == i32(-2) {
			goto l1
		}
		if uint32(v3) >= uint32(v2) {
			store32(m.memory[int64(uint32(v4))+20:], uint32(i32(5)))
			store32(m.memory[int64(uint32(v4))+16:], uint32(v4+i32(44)))
			m.fn73(v4+i32(28), i32(1068224), v4+i32(16))
			m.fn580(v0, i32(21), v4+i32(28))
			goto l3
		}
	l1:
		m.memory[uint32(v0)] = byte(i32(255))
		store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
		goto l3
	}
l0:
	store32(m.memory[int64(uint32(v4))+44:], uint32(v2))
	store32(m.memory[int64(uint32(v4))+40:], uint32(i32(5)))
	store32(m.memory[int64(uint32(v4))+32:], uint32(i32(5)))
	store32(m.memory[int64(uint32(v4))+36:], uint32(v4+i32(44)))
	store32(m.memory[int64(uint32(v4))+28:], uint32(v4+i32(12)))
	m.fn73(v4+i32(16), i32(1066915), v4+i32(28))
	m.fn580(v0, i32(21), v4+i32(16))
l3:
	m.g0 = v4 + i32(48)
}
func (m *Module) fn613(v0 int32) {
	m.fn588(v0)
	m.fn617(v0 + i32(76))
}
func (m *Module) fn614(v0, v1, v2, v3 int32) {
	var v4, v5 int32
	var v6 int64
	t0 := m.g0
	v4 = t0 - i32(48)
	m.g0 = v4
	store32(m.memory[int64(uint32(v4))+16:], uint32(i32(0)))
	store64(m.memory[int64(uint32(v4))+8:], uint64(i64(0x400000000)))
	v5 = v2
	{
	l5:
		{
			{
				if v5 != i32(-2) {
					goto l0
				}
				t1 := int32(load32(m.memory[int64(uint32(v4))+16:]))
				store32(m.memory[int64(uint32(v0))+16:], uint32(t1))
				t2 := int64(load64(m.memory[int64(uint32(v4))+8:]))
				store64(m.memory[int64(uint32(v0))+8:], uint64(t2))
				m.memory[int64(uint32(v0))+24] = byte(v3)
				store32(m.memory[int64(uint32(v0))+20:], uint32(v1))
				store64(m.memory[uint32(v0):], uint64(i64(0)))
				goto l1
			}
		l0:
			m.fn584(v4+i32(8), v5)
			t3 := int32(load32(m.memory[int64(uint32(v1))+52:]))
			t4 := int32(load32(m.memory[int64(uint32(v1))+56:]))
			m.fn612(v4+i32(24), t3, t4, v5)
			{
				{
					t5 := int32(m.memory[int64(uint32(v4))+24])
					if t5 != i32(255) {
						goto l2
					}
					t6 := int32(load32(m.memory[int64(uint32(v4))+28:]))
					v5 = t6
					goto l3
				}
			l2:
				t7 := int64(load64(m.memory[int64(uint32(v4))+24:]))
				v6 = t7
				if v6&i64(255) != i64(255) {
					goto l4
				}
				v5 = int32(int64(uint64(v6) >> 32))
			}
		l3:
			store32(m.memory[int64(uint32(v4))+20:], uint32(v5))
			if v5 != v2 {
				goto l5
			}
		}
		store32(m.memory[int64(uint32(v4))+44:], uint32(i32(5)))
		store32(m.memory[int64(uint32(v4))+40:], uint32(v4+i32(20)))
		m.fn73(v4+i32(24), i32(0x100990), v4+i32(40))
		m.fn580(v0, i32(21), v4+i32(24))
		goto l6
	l4:
		store64(m.memory[uint32(v0):], uint64(v6))
	l6:
		store32(m.memory[int64(uint32(v0))+8:], uint32(i32(-1)))
		t8 := int32(load32(m.memory[int64(uint32(v4))+8:]))
		t9 := int32(load32(m.memory[int64(uint32(v4))+12:]))
		m.fn449(t8, t9)
	}
l1:
	m.g0 = v4 + i32(48)
}
func (m *Module) fn615(v0, v1, v2, v3 int32) {
	var v4 int32
	var v5 int64
	var v6, v7 int32
	var v8 int64
	var v9, v10, v11 int32
	var v12, v13 int64
	var v14 int32
	t0 := m.g0
	v4 = t0 - i32(48)
	m.g0 = v4
	t1 := int64(load64(m.memory[uint32(v1):]))
	v5 = t1
	t2 := int32(load32(m.memory[int64(uint32(v1))+12:]))
	v6 = t2
	t3 := int32(load32(m.memory[int64(uint32(v1))+16:]))
	v7 = t3
	v8 = int64(uint32(v7))
	t4 := int32(load32(m.memory[int64(uint32(v1))+20:]))
	v9 = t4
	v10 = v9 + i32(20)
l10:
	{
		{
			{
				if v3 == 0 {
					m.memory[uint32(v0)] = byte(i32(255))
					goto l7
				}
				t5 := int32(m.memory[uint32(v10)])
				t6 := v8
				v11 = t5
				p7 := i64(9)
				if v11 != 0 {
					p7 = i64(12)
				}
				v12 = p7
				v13 = i64_shl(t6, v12)
				if v13 == v5 {
					goto l1
				}
				t8 := m.fn622(v6, v7, int32(i64_shr_u(v5, v12)), i32(1086688))
				t9 := int32(load32(m.memory[uint32(t8):]))
				t11 := v4 + i32(32)
				t12 := v9
				t13 := v5
				p10 := i64(511)
				if v11 != 0 {
					p10 = i64(0xfff)
				}
				m.fn621(t11, t12, t9, t13&p10)
				t14 := int64(load64(m.memory[int64(uint32(v4))+36:]))
				v12 = t14
				{
					t15 := int32(load32(m.memory[int64(uint32(v4))+32:]))
					v11 = t15
					if v11 != 0 {
						store64(m.memory[int64(uint32(v4))+24:], uint64(v12))
						store32(m.memory[int64(uint32(v4))+20:], uint32(v11))
						t16 := v4
						v12 = v13 - v5
						t17 := v12
						v13 = int64(uint32(v3))
						p18 := v13
						if uint64(v12) < uint64(v13) {
							p18 = t17
						}
						m.fn364(t16, i32(0), int32(p18), v2, v3, i32(1086704))
						t19 := int32(load32(m.memory[uint32(v4):]))
						t20 := int32(load32(m.memory[int64(uint32(v4))+4:]))
						m.fn623(v4+i32(32), v4+i32(20), t19, t20)
						{
							{
								t21 := int32(m.memory[int64(uint32(v4))+32])
								if t21 != i32(255) {
									goto l4
								}
								t22 := int32(load32(m.memory[int64(uint32(v4))+36:]))
								v11 = t22
								goto l5
							}
						l4:
							t23 := int64(load64(m.memory[int64(uint32(v4))+32:]))
							v12 = t23
							v11 = int32(int64(uint64(v12) >> 32))
							if v12&i64(255) != i64(255) {
								goto l3
							}
						}
					l5:
						t24 := v1
						v5 = v5 + int64(uint32(v11))
						store64(m.memory[uint32(t24):], uint64(v5))
						goto l6
					}
					v11 = int32(int64(uint64(v12) >> 32))
					goto l3
				}
			}
		l3:
			store64(m.memory[int64(uint32(v4))+8:], uint64(v12))
			v14 = int32(v12)
			if v14&i32(255) == i32(255) {
				goto l6
			}
			t25 := m.fn313(v4 + i32(8))
			if t25 != 0 {
				m.fn119(v14, v11)
				goto l10
			}
			store64(m.memory[uint32(v0):], uint64(v12))
			goto l7
		}
	l6:
		if v11 == 0 {
			goto l1
		}
		if uint32(v3) < uint32(v11) {
			m.fn151(v11, v3, v3, i32(1072408))
			panic("unreachable")
		}
		v2 = v2 + v11
		v3 = v3 - v11
		goto l10
	l1:
		t26 := int64(load64(m.memory[int64(uint32(i32(0)))+1287056:]))
		store64(m.memory[uint32(v0):], uint64(t26))
	}
l7:
	m.g0 = v4 + i32(48)
}
func (m *Module) fn616(v0 int32) {
	m.fn613(v0)
	t0 := int32(load32(m.memory[int64(uint32(v0))+92:]))
	t1 := int32(load32(m.memory[int64(uint32(v0))+96:]))
	m.fn449(t0, t1)
	t2 := int32(load32(m.memory[int64(uint32(v0))+104:]))
	t3 := int32(load32(m.memory[int64(uint32(v0))+108:]))
	m.fn449(t2, t3)
}
func (m *Module) fn617(v0 int32) {
	var v1, v2, v3 int32
	t0 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	v1 = t0
	v2 = v1 + i32(64)
	t1 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	v3 = t1
l1:
	{
		if v3 == 0 {
			goto l0
		}
		t2 := int32(load32(m.memory[uint32(v2+i32(-4)):]))
		t3 := int32(load32(m.memory[uint32(v2):]))
		m.fn16(t2, t3)
		v3 = v3 + i32(-1)
		v2 = v2 + i32(80)
		goto l1
	}
l0:
	t4 := int32(load32(m.memory[uint32(v0):]))
	m.fn136(t4, v1, i32(8), i32(80))
}
func (m *Module) fn618(v0 int32) {
	var v1 int32
	t0 := m.g0
	v1 = t0 - i32(16)
	m.g0 = v1
	t1 := int32(load32(m.memory[uint32(v0):]))
	m.fn273(v1+i32(8), v0, t1, i32(1), i32(4), i32(4))
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
func (m *Module) fn619(v0 int32) int64 {
	var v1 int32
	var v2 int64
	t0 := m.g0
	v1 = t0 - i32(16)
	store32(m.memory[int64(uint32(v1))+12:], uint32(v0))
	v2 = i64(-0x340d631b7bdddcdb)
	v0 = i32(0)
l1:
	{
		if v0 == i32(4) {
			return v2
		}
		t1 := int64(m.memory[uint32(v1+i32(12)+v0)])
		v2 = (v2 ^ t1) * i64(0x100000001b3)
		v0 = v0 + i32(1)
		goto l1
	}
}
func (m *Module) fn620(v0, v1 int32) int32 {
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
		m.fn237(t5, i32(4), p6)
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
		store64(m.memory[int64(uint32(v2))+24:], uint64(i64(0x800000004)))
		store32(m.memory[int64(uint32(v2))+20:], uint32(v0+i32(16)))
		t11 := int32(load32(m.memory[uint32(v0):]))
		v4 = t11
		t12 := int64(load64(m.memory[uint32(v4):]))
		v8 = (t12 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
		v9 = v2 + i32(32)
		v1 = i32(0)
	l6:
		if v3 == 0 {
			t24 := int32(load32(m.memory[int64(uint32(v0))+12:]))
			t25 := v2
			v1 = t24
			store32(m.memory[int64(uint32(t25))+44:], uint32(v1))
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
				t19 := m.fn624(t17, t18, v10)
				v11 = t19
				t20 := m.fn26(t15, t16, v11)
				v12 = t20
				t21 := t14 + v12
				v13 = int32(uint32(int32(v11)) >> 25)
				m.memory[uint32(t21)] = byte(v13)
				m.memory[uint32(v6+v5&(v12+i32(-8))+i32(8))] = byte(v13)
				t22 := int32(load32(m.memory[uint32(v0):]))
				t23 := int32(load32(m.memory[uint32(t22+(v10^i32(-1))<<2):]))
				store32(m.memory[uint32(v6+(v12^i32(-1))<<2):], uint32(t23))
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
	m.fn241(v0, v2+i32(16), i32(98), i32(4))
l7:
	v5 = i32(-1)
l2:
	m.g0 = v2 + i32(64)
	return v5
}
func (m *Module) fn621(v0, v1, v2 int32, v3 int64) {
	var v4, v5 int32
	t0 := m.g0
	v4 = t0 - i32(32)
	m.g0 = v4
	store32(m.memory[uint32(v4):], uint32(v2))
	{
		{
			t1 := int32(load32(m.memory[int64(uint32(v1))+16:]))
			if uint32(v2) < uint32(t1) {
				goto l0
			}
			store32(m.memory[int64(uint32(v4))+28:], uint32(i32(5)))
			store32(m.memory[int64(uint32(v4))+24:], uint32(v1+i32(16)))
			store32(m.memory[int64(uint32(v4))+20:], uint32(i32(5)))
			store32(m.memory[int64(uint32(v4))+16:], uint32(v4))
			m.fn73(v4+i32(4), i32(1048993), v4+i32(16))
			m.fn580(v0+i32(4), i32(21), v4+i32(4))
			v1 = i32(0)
			goto l1
		}
	l0:
		store32(m.memory[int64(uint32(v0))+8:], uint32(v3))
		t2 := int32(m.memory[int64(uint32(v1))+20])
		t3 := v0
		v5 = t2
		p4 := i32(512)
		if v5 != 0 {
			p4 = i32(4096)
		}
		store32(m.memory[int64(uint32(t3))+4:], uint32(p4))
		t6 := v1
		t7 := int64(uint32(v2 + i32(1)))
		p5 := i64(9)
		if v5 != 0 {
			p5 = i64(12)
		}
		store64(m.memory[int64(uint32(t6))+8:], uint64(i64_shl(t7, p5)+v3))
	}
l1:
	store32(m.memory[uint32(v0):], uint32(v1))
	m.g0 = v4 + i32(32)
}
func (m *Module) fn622(v0, v1, v2, v3 int32) int32 {
	if uint32(v2) < uint32(v1) {
		return v0 + v2<<2
	}
	m.fn158(v2, v1, v3)
	panic("unreachable")
}
func (m *Module) fn623(v0, v1, v2, v3 int32) {
	var v4, v5, v6, v7 int32
	var v8 int64
	t0 := m.g0
	v4 = t0 - i32(16)
	m.g0 = v4
	{
		{
			t1 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t2 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			v5 = t2
			v6 = t1 - v5
			p3 := v3
			if uint32(v6) < uint32(v3) {
				p3 = v6
			}
			v6 = p3
			if v6 != 0 {
				goto l0
			}
			m.memory[uint32(v0)] = byte(i32(255))
			store32(m.memory[int64(uint32(v0))+4:], uint32(i32(0)))
			goto l1
		}
	l0:
		t4 := int32(load32(m.memory[uint32(v1):]))
		v7 = t4
		m.fn364(v4, i32(0), v6, v2, v3, i32(1086720))
		t5 := int32(load32(m.memory[uint32(v4):]))
		t6 := int32(load32(m.memory[int64(uint32(v4))+4:]))
		m.fn301(v4+i32(8), v7, t5, t6)
		{
			{
				t7 := int32(m.memory[int64(uint32(v4))+8])
				if t7 != i32(255) {
					goto l2
				}
				t8 := int32(load32(m.memory[int64(uint32(v4))+12:]))
				v3 = t8
				goto l3
			}
		l2:
			t9 := int64(load64(m.memory[int64(uint32(v4))+8:]))
			v8 = t9
			if v8&i64(255) != i64(255) {
				goto l4
			}
			v3 = int32(int64(uint64(v8) >> 32))
		}
	l3:
		store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
		m.memory[uint32(v0)] = byte(i32(255))
		store32(m.memory[int64(uint32(v1))+8:], uint32(v3+v5))
		goto l1
	l4:
		store64(m.memory[uint32(v0):], uint64(v8))
	}
l1:
	m.g0 = v4 + i32(16)
}
func (m *Module) fn624(v0, v1, v2 int32) int64 {
	t0 := int32(load32(m.memory[uint32(v1):]))
	t1 := int32(load32(m.memory[uint32(t0-v2<<2+i32(-4)):]))
	t2 := m.fn619(t1)
	return t2
}
func (m *Module) fn625(v0 int32) {
	var v1 int32
	t0 := m.g0
	v1 = t0 - i32(16)
	m.g0 = v1
	t1 := int32(load32(m.memory[uint32(v0):]))
	m.fn273(v1+i32(8), v0, t1, i32(1), i32(4), i32(8))
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
func (m *Module) fn626(v0, v1, v2, v3, v4 int32) {
	t0 := m.fn156(v3, v4, v1, v2)
	v1 = t0
	store32(m.memory[int64(uint32(v0))+4:], uint32(v4-v2))
	t2 := v0
	p1 := i32(0)
	if v1 != 0 {
		p1 = v3
	}
	store32(m.memory[uint32(t2):], uint32(p1))
}
