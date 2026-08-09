package core

import (
	"math/bits"
)

func (m *Module) fn897(v0 int32) {
	t0 := int32(load32(m.memory[uint32(v0):]))
	if t0 == i32(-2) {
		return
	}
	m.fn898(v0)
}
func (m *Module) fn898(v0 int32) {
	t0 := int32(load32(m.memory[uint32(v0):]))
	if t0 == i32(-1) {
		return
	}
	m.fn893(v0)
}
func (m *Module) fn899(v0 int32) {
	m.fn897(v0)
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+28:]))
		if t0 == 0 {
			return
		}
		m.fn900(v0 + i32(28))
	}
}
func (m *Module) fn900(v0 int32) {
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
	m.fn893(v1)
	v1 = v1 + i32(28)
	goto l1
l0:
	t3 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	t4 := int32(load32(m.memory[uint32(v0):]))
	m.fn82(t3, t4)
}
func (m *Module) fn901(v0, v1 int32) {
	var v2, v3 int32
	t0 := int32(load32(m.memory[int64(uint32(v1))+28:]))
	v2 = t0
	{
		{
			t1 := int32(load32(m.memory[uint32(v1):]))
			v3 = t1
			if v3 == i32(-2) {
				v3 = i32(0)
				if v2 != 0 {
					t3 := int32(load32(m.memory[int64(uint32(v1))+40:]))
					t4 := int32(load32(m.memory[int64(uint32(v1))+32:]))
					t5 := int32(uint32(t3-t4) / uint32(i32(28)))
					v1 = t5
					goto l2
				}
				v1 = i32(0)
				goto l2
			}
			var p2 int32
			if v3 != i32(-1) {
				p2 = 1
			}
			v3 = p2
			if v2 != 0 {
				goto l1
			}
			v1 = v3
			goto l2
		}
	l1:
		t6 := int32(load32(m.memory[int64(uint32(v1))+40:]))
		t7 := int32(load32(m.memory[int64(uint32(v1))+32:]))
		t8 := int32(uint32(t6-t7) / uint32(i32(28)))
		v1 = t8 + v3
	}
l2:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v1))
	store32(m.memory[int64(uint32(v0))+4:], uint32(i32(1)))
	store32(m.memory[uint32(v0):], uint32(v3))
}
func (m *Module) fn902(v0, v1 int32) {
	var v2 int32
	{
		t0 := int32(load32(m.memory[uint32(v0):]))
		t1 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		t2 := v1
		v2 = t1
		if uint32(t2) <= uint32(t0-v2) {
			return
		}
		m.fn62(v0, v2, v1, i32(4), i32(28))
	}
}
func (m *Module) fn903(v0, v1, v2 int32) {
	var v3 int32
	t0 := m.g0
	v3 = t0 - i32(16)
	m.g0 = v3
	store32(m.memory[int64(uint32(v3))+12:], uint32(v2))
	store32(m.memory[int64(uint32(v3))+8:], uint32(v1))
l2:
	{
		t1 := m.fn904(v3 + i32(8))
		v1 = t1
		if v1 == 0 {
			goto l0
		}
		{
			t2 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			v2 = t2
			t3 := int32(load32(m.memory[uint32(v0):]))
			if v2 != t3 {
				goto l1
			}
			m.fn905(v0)
		}
	l1:
		store32(m.memory[int64(uint32(v0))+8:], uint32(v2+i32(1)))
		t4 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		store32(m.memory[uint32(t4+v2<<2):], uint32(v1))
		goto l2
	}
l0:
	m.g0 = v3 + i32(16)
}
func (m *Module) fn904(v0 int32) int32 {
	var v1, v2, v3, v4 int32
	t0 := int32(load32(m.memory[uint32(v0):]))
	v1 = t0
	t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	v2 = t1
l1:
	{
		if v1 != v2 {
			goto l0
		}
		return i32(0)
	l0:
		t2 := v0
		v3 = v1 + i32(44)
		store32(m.memory[uint32(t2):], uint32(v3))
		t3 := int32(load32(m.memory[uint32(v1):]))
		v4 = t3
		v1 = v3
		if v4 == i32(-1) {
			goto l1
		}
	}
	return v3 + i32(-44)
}
func (m *Module) fn905(v0 int32) {
	var v1 int32
	{
		t0 := int32(load32(m.memory[uint32(v0):]))
		t1 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		v1 = t1
		if t0 != v1 {
			return
		}
		m.fn62(v0, v1, i32(1), i32(4), i32(4))
	}
}
func (m *Module) fn906(v0 int32) int32 {
	var v1, v2, v3 int32
	t0 := int32(load32(m.memory[uint32(v0):]))
	v1 = t0
	t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	v2 = t1
l1:
	{
		v3 = v1
		if v3 != v2 {
			goto l0
		}
		return i32(0)
	l0:
		t2 := v0
		v1 = v3 + i32(44)
		store32(m.memory[uint32(t2):], uint32(v1))
		t3 := int32(load32(m.memory[uint32(v3):]))
		if t3 == i32(-1) {
			goto l1
		}
		t4 := int32(load32(m.memory[uint32(v3+i32(4)):]))
		t5 := int32(load32(m.memory[uint32(v3+i32(8)):]))
		t6 := m.fn773(t4, t5, i32(1073486), i32(2))
		if t6 == 0 {
			goto l1
		}
	}
	return v3
}
func (m *Module) fn907(v0 int32) int32 {
	var v1, v2, v3 int32
	t0 := int32(load32(m.memory[uint32(v0):]))
	v1 = t0
	t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	v2 = t1
l1:
	{
		v3 = v1
		if v3 != v2 {
			goto l0
		}
		return i32(0)
	l0:
		t2 := v0
		v1 = v3 + i32(44)
		store32(m.memory[uint32(t2):], uint32(v1))
		t3 := int32(load32(m.memory[uint32(v3):]))
		if t3 == i32(-1) {
			goto l1
		}
		t4 := m.fn847(v3, i32(1074726), i32(47), i32(1074842), i32(5))
		if t4 == 0 {
			goto l1
		}
	}
	return v3
}
func (m *Module) fn908(v0, v1 int32) {
	var v2, v3, v4, v5, v6 int32
	var v7 int64
	var v8, v9 int32
	var v10 int64
	var v11 int32
	var v12 int64
	var v13, v14 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	v3 = i32(0)
	{
	l2:
		{
			{
				t1 := m.fn866(v1)
				v4 = t1
				if v4 != 0 {
					goto l0
				}
				goto l1
			}
		l0:
			t2 := int32(load32(m.memory[uint32(v4):]))
			if t2 == i32(-1) {
				goto l2
			}
			t3 := int32(load32(m.memory[uint32(v4+i32(4)):]))
			t4 := int32(load32(m.memory[uint32(v4+i32(8)):]))
			t5 := int32(load32(m.memory[int64(uint32(v1))+12:]))
			t6 := int32(load32(m.memory[int64(uint32(v1))+16:]))
			t7 := m.fn773(t3, t4, t5, t6)
			if t7 == 0 {
				goto l2
			}
			t8 := int32(load32(m.memory[uint32(v4+i32(16)):]))
			t9 := int32(load32(m.memory[uint32(v4+i32(20)):]))
			m.fn909(v2+i32(8), t8, t9, i32(1074298), i32(5))
			t10 := int32(load32(m.memory[int64(uint32(v2))+12:]))
			v5 = t10
			t11 := int32(load32(m.memory[int64(uint32(v2))+8:]))
			v6 = t11
			if v6 == 0 {
				goto l2
			}
			t12 := int32(load32(m.memory[int64(uint32(v1))+20:]))
			v4 = t12
			t13 := int32(load32(m.memory[int64(uint32(v4))+12:]))
			if t13 == 0 {
				goto l2
			}
			t14 := int64(load64(m.memory[int64(uint32(v4))+16:]))
			t15 := int64(load64(m.memory[uint32(v4+i32(24)):]))
			t16 := m.fn29(t14, t15, v6, v5)
			v7 = t16
			t17 := int32(load32(m.memory[int64(uint32(v4))+4:]))
			v8 = t17
			v9 = v8 & int32(v7)
			v10 = int64(uint64(v7)>>25) & i64(127) * i64(72340172838076673)
			t18 := int32(load32(m.memory[uint32(v4):]))
			v4 = t18
			v11 = i32(0)
		l6:
			{
				t19 := int64(load64(m.memory[uint32(v4+v9):]))
				v12 = t19
				v7 = v12 ^ v10
				v7 = (v7 ^ i64(-1)) & (v7 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
			l5:
				{
					if v7 == 0 {
						if !(v12&(v12<<1)&i64(-0x7f7f7f7f7f7f7f80) == 0) {
							goto l2
						}
						t26 := v9
						v11 = v11 + i32(8)
						v9 = (t26 + v11) & v8
						goto l6
					}
					t20 := v6
					t21 := v5
					t22 := v4
					v13 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v7))))>>3) + v9) & v8
					v14 = t22 + (i32(0)-v13)*i32(36)
					t23 := int32(load32(m.memory[uint32(v14+i32(-32)):]))
					t24 := int32(load32(m.memory[uint32(v14+i32(-28)):]))
					t25 := m.fn15(t20, t21, t23, t24)
					if t25 != 0 {
						goto l4
					}
					v7 = (v7 + i64(-1)) & v7
					goto l5
				}
			l4:
			}
		}
		v1 = v4 + (i32(0)-v13)*i32(36)
		t27 := int32(load32(m.memory[uint32(v1+i32(-16)):]))
		v4 = t27
		t28 := int32(load32(m.memory[uint32(v1+i32(-20)):]))
		v3 = t28
	}
l1:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v4))
	store32(m.memory[uint32(v0):], uint32(v3))
	m.g0 = v2 + i32(16)
}
func (m *Module) fn909(v0, v1, v2, v3, v4 int32) {
	var v5 int32
	v5 = v2 << 5
	v1 = v1 + i32(8)
	{
	l2:
		{
			v2 = v1
			if v5 != 0 {
				goto l0
			}
			v2 = i32(0)
			goto l1
		l0:
			v5 = v5 + i32(-32)
			v1 = v2 + i32(32)
			t0 := int32(load32(m.memory[uint32(v2+i32(-4)):]))
			t1 := int32(load32(m.memory[uint32(v2):]))
			t2 := m.fn773(t0, t1, v3, v4)
			if t2 == 0 {
				goto l2
			}
		}
		v2 = v2 + i32(-8)
		t3 := int32(load32(m.memory[int64(uint32(v2))+20:]))
		v5 = t3
		t4 := int32(load32(m.memory[int64(uint32(v2))+16:]))
		v2 = t4
	}
l1:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
	store32(m.memory[uint32(v0):], uint32(v2))
}
func (m *Module) fn910(v0, v1 int32) {
	var v2, v3, v4, v5, v6, v7, v8, v9 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	t1 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	v3 = t1
	t2 := int32(load32(m.memory[uint32(v1):]))
	v4 = t2
	t3 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	v5 = t3
l3:
	{
		if v4 != v5 {
			goto l0
		}
		v6 = i32(0)
		goto l1
	l0:
		t4 := v1
		v7 = v4 + i32(12)
		store32(m.memory[uint32(t4):], uint32(v7))
		v6 = i32(0)
		{
			t5 := int32(load32(m.memory[uint32(v3):]))
			v8 = t5
			t6 := int32(load32(m.memory[uint32(v4+i32(8)):]))
			if uint32(v8) >= uint32(t6) {
				goto l2
			}
			t7 := int32(load32(m.memory[uint32(v4+i32(4)):]))
			t8 := v2 + i32(8)
			v4 = t7 + v8*i32(12)
			t9 := int32(load32(m.memory[int64(uint32(v4))+4:]))
			t10 := int32(load32(m.memory[int64(uint32(v4))+8:]))
			m.fn46(t8, t9, t10)
			t11 := int32(load32(m.memory[int64(uint32(v2))+8:]))
			t12 := int32(load32(m.memory[int64(uint32(v2))+12:]))
			v9 = t12
			p13 := i32(0)
			if v9 != 0 {
				p13 = t11
			}
			v6 = p13
		}
	l2:
		v4 = v7
		if v6 == 0 {
			goto l3
		}
	}
l1:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v9))
	store32(m.memory[uint32(v0):], uint32(v6))
	m.g0 = v2 + i32(16)
}
func (m *Module) fn911(v0, v1 int32) {
	m.fn136(v0, v1, i32(4), i32(12))
}
func (m *Module) fn912(v0, v1 int32) int32 {
	var v2, v3 int32
	var v4 int64
	var v5, v6 int32
	var v7 int64
	var v8 int32
	var v9 int64
	var v10, v11 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
		if t0 != 0 {
			t1 := int64(load64(m.memory[int64(uint32(v0))+16:]))
			t2 := int64(load64(m.memory[int64(uint32(v0))+24:]))
			t3 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			v2 = t3
			t4 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			t5 := v2
			v3 = t4
			t6 := m.fn540(t1, t2, t5, v3)
			v4 = t6
			t7 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			v5 = t7
			v6 = v5 & int32(v4)
			v7 = int64(uint64(v4)>>25) & i64(127) * i64(72340172838076673)
			t8 := int32(load32(m.memory[uint32(v0):]))
			v0 = t8
			v8 = i32(0)
			var _ int32
		l5:
			{
				t10 := int64(load64(m.memory[uint32(v0+v6):]))
				v9 = t10
				v4 = v9 ^ v7
				v4 = (v4 ^ i64(-1)) & (v4 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
				{
				l3:
					{
						var p11 int32
						if v4 == 0 {
							p11 = 1
						}
						v1 = p11
						if v1 != 0 {
							goto l1
						}
						t12 := v2
						t13 := v3
						t14 := v0
						v10 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v4))))>>3) + v6) & v5
						v11 = t14 - v10<<4
						t15 := int32(load32(m.memory[uint32(v11+i32(-12)):]))
						t16 := int32(load32(m.memory[uint32(v11+i32(-8)):]))
						t17 := m.fn544(t12, t13, t15, t16)
						if t17 != 0 {
							goto l2
						}
						v4 = (v4 + i64(-1)) & v4
						goto l3
					}
				l1:
					if v9&(v9<<1)&i64(-0x7f7f7f7f7f7f7f80) == 0 {
						t20 := v6
						v8 = v8 + i32(8)
						v6 = (t20 + v8) & v5
						goto l5
					}
				l2:
					p18 := v0 - v10<<4
					if v1 != 0 {
						p18 = i32(0)
					}
					p19 := p18 + i32(-4)
					if v1 != 0 {
						p19 = i32(0)
					}
					return p19
				}
			}
		}
		return i32(0)
	}
}
func (m *Module) fn913(v0, v1 int32) {
	var v2, v3, v4, v5, v6, v7, v8, v9, v10 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	t1 := int32(load32(m.memory[int64(uint32(v1))+12:]))
	v3 = t1
	t2 := int32(load32(m.memory[int64(uint32(v1))+16:]))
	v4 = t2
	t3 := int32(load32(m.memory[uint32(v1):]))
	v5 = t3
	t4 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	v6 = t4
	t5 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	v7 = t5
	v8 = v7 + i32(4)
	{
	l2:
		{
			v9 = v5
			if v9 == v6 {
				goto l0
			}
			t6 := v1
			v5 = v9 + i32(36)
			store32(m.memory[uint32(t6):], uint32(v5))
			m.fn855(v2+i32(8), v9)
			{
				t7 := int32(load32(m.memory[int64(uint32(v2))+8:]))
				v10 = t7
				if v10 == 0 {
					goto l1
				}
				t8 := int32(load32(m.memory[int64(uint32(v2))+12:]))
				t9 := int32(load32(m.memory[uint32(v7):]))
				t10 := int32(load32(m.memory[uint32(v8):]))
				t11 := m.fn15(v10, t8, t9, t10)
				if t11 == 0 {
					goto l2
				}
			}
		l1:
			m.fn855(v2, v9+i32(12))
			{
				t12 := int32(load32(m.memory[uint32(v2):]))
				v10 = t12
				if v10 == 0 {
					goto l3
				}
				t13 := int32(load32(m.memory[int64(uint32(v2))+4:]))
				t14 := m.fn914(v10, t13, v3, v4)
				if t14 == 0 {
					goto l2
				}
			}
		l3:
		}
		t15 := int32(m.memory[int64(uint32(v9))+32])
		m.memory[int64(uint32(v0))+8] = byte(t15)
		t16 := int64(load64(m.memory[int64(uint32(v9))+24:]))
		store64(m.memory[uint32(v0):], uint64(t16))
		goto l4
	}
l0:
	m.memory[int64(uint32(v0))+4] = byte(i32(255))
l4:
	m.g0 = v2 + i32(16)
}
func (m *Module) fn914(v0, v1, v2, v3 int32) int32 {
	v3 = v3 << 3
l2:
	if v3 == 0 {
		goto l0
	}
	{
		t0 := int32(load32(m.memory[uint32(v2+i32(4)):]))
		if t0 != v1 {
			goto l1
		}
		t1 := int32(load32(m.memory[uint32(v2):]))
		t2 := m.fn1851(t1, v0, v1)
		if t2 == 0 {
			goto l0
		}
	}
l1:
	v2 = v2 + i32(8)
	v3 = v3 + i32(-8)
	goto l2
l0:
	;
	var p3 int32
	if v3 != i32(0) {
		p3 = 1
	}
	return p3
}
func (m *Module) fn915(v0, v1 int32) {
	var v2, v3 int32
	v2 = i32(0)
	{
		{
			t0 := int32(m.memory[int64(uint32(v1))+8])
			if t0 != 0 {
				goto l0
			}
			t1 := int32(load32(m.memory[uint32(v1):]))
			v3 = t1
			t2 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			if v3 != t2 {
				goto l1
			}
		}
	l0:
		goto l2
	l1:
		v2 = i32(1)
		store32(m.memory[uint32(v1):], uint32(v3+i32(1)))
		t3 := int32(m.memory[uint32(v3)])
		v3 = t3
		if v3 != 0 {
			goto l2
		}
		m.memory[int64(uint32(v1))+8] = byte(i32(1))
		v2 = i32(0)
		v3 = i32(0)
	}
l2:
	m.memory[int64(uint32(v0))+1] = byte(v3)
	m.memory[uint32(v0)] = byte(v2)
}
func (m *Module) fn916(v0, v1 int32) {
	var v2, v3, v4 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	v3 = i32(0)
	{
		t1 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		v4 = t1
		if v4 == 0 {
			goto l2
		}
		store32(m.memory[int64(uint32(v1))+8:], uint32(v4+i32(-1)))
		m.fn917(v2+i32(8), v1)
		t2 := int32(load32(m.memory[int64(uint32(v2))+8:]))
		v1 = t2
		if v1 == i32(2) {
			goto l2
		}
		t3 := int32(load32(m.memory[int64(uint32(v2))+12:]))
		v4 = t3
		if v1 != i32(1) {
			if v4 == 0 {
				goto l2
			}
			t4 := int32(load32(m.memory[int64(uint32(v4))+44:]))
			v1 = t4
			m.fn919(v4)
			v3 = i32(1)
			goto l2
		}
		m.fn918(v4)
		goto l2
	}
l2:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
	store32(m.memory[uint32(v0):], uint32(v3))
	m.g0 = v2 + i32(16)
}
func (m *Module) fn917(v0, v1 int32) {
	var v2, v3, v4, v5, v6 int32
	var v7, v8, v9 int64
	var v10, v11, v12, v13, v14, v15, v16, v17, v18, v19, v20, v21 int32
	t0 := m.g0
	v2 = t0 - i32(192)
	m.g0 = v2
	t1 := int32(load32(m.memory[uint32(v1):]))
	t2 := v2
	v3 = t1
	t3 := int64(load64(m.memory[int64(uint32(v3))+48:]))
	store64(m.memory[int64(uint32(t2))+56:], uint64(t3))
	t4 := int64(load64(m.memory[int64(uint32(v3))+40:]))
	store64(m.memory[int64(uint32(v2))+48:], uint64(t4))
	t5 := int64(load64(m.memory[int64(uint32(v3))+32:]))
	store64(m.memory[int64(uint32(v2))+40:], uint64(t5))
	v4 = v1 + i32(4)
	{
		{
			t6 := int32(m.memory[int64(uint32(v3))+59])
			if t6 != 0 {
				goto l0
			}
			t7 := int32(m.memory[int64(uint32(v3))+56])
			if t7&i32(1) != 0 {
				goto l0
			}
			t8 := int32(m.memory[int64(uint32(v3))+58])
			if t8&i32(1) != 0 {
				goto l0
			}
			t9 := int32(load32(m.memory[int64(uint32(v3))+16:]))
			if t9 != i32(2) {
				m.memory[int64(uint32(v3))+58] = byte(i32(1))
				t56 := int32(load32(m.memory[int64(uint32(v3))+28:]))
				t57 := m.fn929(t56)
				v1 = t57
				t58 := int32(load32(m.memory[uint32(v4):]))
				m.fn919(t58)
				store32(m.memory[uint32(v4):], uint32(v1))
				{
					t59 := int32(m.memory[int64(uint32(v3))+61])
					if uint32(t59) < uint32(i32(2)) {
						goto l13
					}
					m.fn930(v4)
					t60 := int32(load32(m.memory[uint32(v4):]))
					v1 = t60
				}
			l13:
				v17 = i32(0)
				t61 := int32(load32(m.memory[int64(uint32(v1))+44:]))
				var p62 int32
				if t61 != i32(0) {
					p62 = 1
				}
				v5 = p62
				goto l14
			}
		}
	l0:
		v5 = i32(0)
		t10 := int32(load32(m.memory[uint32(v4):]))
		v1 = t10
		store32(m.memory[int64(uint32(v1))+44:], uint32(i32(0)))
		v6 = v3 + i32(32)
		t11 := int64(load64(m.memory[int64(uint32(v6))+16:]))
		v7 = t11
		t12 := int64(load64(m.memory[int64(uint32(v6))+8:]))
		v8 = t12
		t13 := int64(load64(m.memory[uint32(v6):]))
		v9 = t13
		store64(m.memory[uint32(v1):], uint64(i64(1)))
		store64(m.memory[int64(uint32(v1))+8:], uint64(v9))
		store64(m.memory[int64(uint32(v1))+16:], uint64(v8))
		store64(m.memory[int64(uint32(v1))+24:], uint64(v7))
		t14 := int32(m.memory[int64(uint32(v3))+60])
		if t14 != 0 {
			goto l2
		}
		v10 = v3 + i32(64)
		v1 = i32(0)
		t15 := int32(load32(m.memory[uint32(v4):]))
		v11 = t15
		v12 = v11 + i32(56)
		v13 = v11 + i32(52)
		v14 = v11 + i32(36)
		v15 = v11 + i32(40)
		v16 = i32(0)
	l5:
		{
			m.fn920(v2+i32(116), v10)
			{
				t16 := int32(load32(m.memory[int64(uint32(v2))+116:]))
				if t16 == 0 {
					goto l3
				}
				m.memory[int64(uint32(v3))+60] = byte(i32(2))
				t17 := int64(load64(m.memory[int64(uint32(v2))+120:]))
				t18 := m.fn921(t17)
				v1 = t18
				goto l4
			}
		l3:
			t19 := int32(load32(m.memory[int64(uint32(v3))+92:]))
			v17 = t19
			t20 := int32(load32(m.memory[int64(uint32(v2))+124:]))
			v5 = t20
			t21 := int32(load32(m.memory[int64(uint32(v2))+120:]))
			v18 = t21
			t22 := int32(load32(m.memory[uint32(v13):]))
			t23 := int32(load32(m.memory[uint32(v12):]))
			m.fn922(v2+i32(32), t22, t23, v16)
			t24 := int32(load32(m.memory[int64(uint32(v2))+36:]))
			v19 = t24
			t25 := int32(load32(m.memory[int64(uint32(v2))+32:]))
			v20 = t25
			t26 := int32(load32(m.memory[uint32(v14):]))
			t27 := int32(load32(m.memory[uint32(v15):]))
			m.fn923(v2+i32(24), t26, t27, v1)
			t28 := int32(load32(m.memory[int64(uint32(v2))+24:]))
			t29 := int32(load32(m.memory[int64(uint32(v2))+28:]))
			m.fn924(v2+i32(128), v17, v18, v5, v20, v19, t28, t29)
			t30 := int32(m.memory[int64(uint32(v2))+136])
			v5 = t30
			t31 := int32(load32(m.memory[int64(uint32(v2))+132:]))
			v18 = t31
			t32 := int32(load32(m.memory[int64(uint32(v2))+140:]))
			v19 = t32
			t33 := int64(load64(m.memory[int64(uint32(v3))+32:]))
			t34 := int32(load32(m.memory[int64(uint32(v2))+128:]))
			t35 := v3
			v20 = t34
			store64(m.memory[int64(uint32(t35))+32:], uint64(t33+int64(uint32(v20))))
			t36 := int32(load32(m.memory[int64(uint32(v3))+76:]))
			t37 := v3
			v21 = t36
			t38 := int32(load32(m.memory[int64(uint32(v3))+72:]))
			t39 := v21
			v20 = v20 + t38
			p40 := v20
			if uint32(v21) < uint32(v20) {
				p40 = t39
			}
			store32(m.memory[int64(uint32(t37))+72:], uint32(p40))
			t41 := int64(load64(m.memory[uint32(v17):]))
			_ = m.fn925(v6, t41)
			v1 = v19 + v1
			v16 = v18 + v16
			switch v5 {
			case 4:
				goto l9
			default:
				goto l5
			case 1:
				m.fn926(v11)
				goto l5
			case 2:
				m.fn927(v11)
				goto l5
			case 3:
			}
		}
		t43 := int32(load32(m.memory[uint32(v4):]))
		store32(m.memory[int64(uint32(t43))+44:], uint32(v1))
		t44 := int64(load64(m.memory[int64(uint32(v3))+48:]))
		v7 = t44
		if v7 == i64(-1) {
			m.fn153(i32(1078240))
			panic("unreachable")
		}
		store64(m.memory[int64(uint32(v3))+48:], uint64(v7+i64(1)))
		v5 = i32(1)
		t45 := int32(m.memory[int64(uint32(v3))+57])
		if t45 != 0 {
			goto l2
		}
		t46 := int64(load64(m.memory[uint32(v3):]))
		if t46 != i64(1) {
			store64(m.memory[uint32(v3):], uint64(i64(1)))
			t63 := int32(load32(m.memory[uint32(v4):]))
			t64 := int64(load32(m.memory[int64(uint32(t63))+44:]))
			store64(m.memory[int64(uint32(v3))+8:], uint64(t64))
			goto l2
		}
		t47 := int64(load64(m.memory[int64(uint32(v3))+8:]))
		v7 = t47
		t48 := int32(load32(m.memory[uint32(v4):]))
		t49 := v7
		v1 = t48
		t50 := int64(load32(m.memory[int64(uint32(v1))+44:]))
		v8 = t50
		if t49 == v8 {
			goto l2
		}
		v9 = i64(0)
		{
			t51 := int64(load64(m.memory[uint32(v1):]))
			if t51 != i64(1) {
				goto l12
			}
			t52 := int64(load64(m.memory[int64(uint32(v1))+24:]))
			store64(m.memory[int64(uint32(v2))+160:], uint64(t52))
			t53 := int64(load64(m.memory[int64(uint32(v1))+16:]))
			store64(m.memory[int64(uint32(v2))+152:], uint64(t53))
			t54 := int64(load64(m.memory[int64(uint32(v1))+8:]))
			store64(m.memory[int64(uint32(v2))+144:], uint64(t54))
			v9 = i64(1)
		}
	l12:
		store64(m.memory[int64(uint32(v2))+176:], uint64(v8))
		store64(m.memory[int64(uint32(v2))+168:], uint64(v7))
		store64(m.memory[int64(uint32(v2))+136:], uint64(v9))
		store64(m.memory[int64(uint32(v2))+128:], uint64(i64(4)))
		t55 := m.fn928(v2 + i32(128))
		v1 = t55
		goto l4
	}
l9:
	m.memory[int64(uint32(v3))+60] = byte(i32(1))
	v5 = i32(0)
l2:
	m.memory[int64(uint32(v3))+58] = byte(i32(1))
	{
		t65 := int32(m.memory[int64(uint32(v3))+59])
		if t65 != 0 {
			goto l15
		}
		t66 := int32(load32(m.memory[int64(uint32(v3))+16:]))
		if t66 != i32(2) {
			goto l15
		}
		t67 := int32(load32(m.memory[uint32(v4):]))
		t68 := m.fn929(t67)
		t69 := v2
		v16 = t68
		t70 := m.fn929(v16)
		v1 = t70
		store32(m.memory[int64(uint32(t69))+116:], uint32(v1))
		m.fn931(v2+i32(128), v2+i32(116))
		v17 = i32(1)
		{
			t71 := int32(load32(m.memory[int64(uint32(v2))+128:]))
			if t71 != i32(1) {
				goto l16
			}
			t72 := int32(load32(m.memory[int64(uint32(v2))+136:]))
			v18 = t72
			t73 := int32(load32(m.memory[int64(uint32(v2))+132:]))
			v19 = t73
			m.fn919(v1)
			goto l17
		}
	l16:
		v17 = i32(0)
		v19 = v1
	l17:
		v1 = v3 + i32(16)
		store32(m.memory[int64(uint32(v2))+136:], uint32(v18))
		store32(m.memory[int64(uint32(v2))+132:], uint32(v19))
		store32(m.memory[int64(uint32(v2))+128:], uint32(v17))
		store32(m.memory[int64(uint32(v2))+116:], uint32(v16))
		{
			t74 := int32(m.memory[int64(uint32(v3))+61])
			if t74&i32(1) == 0 {
				goto l18
			}
			if v17 != 0 {
				goto l19
			}
			m.fn932(v2 + i32(132))
		l19:
			m.fn930(v2 + i32(116))
			t75 := int32(load32(m.memory[int64(uint32(v2))+116:]))
			v16 = t75
		}
	l18:
		m.fn933(v1)
		store32(m.memory[int64(uint32(v3))+28:], uint32(v16))
		t76 := int32(load32(m.memory[int64(uint32(v2))+136:]))
		store32(m.memory[int64(uint32(v1))+8:], uint32(t76))
		t77 := int64(load64(m.memory[int64(uint32(v2))+128:]))
		store64(m.memory[uint32(v1):], uint64(t77))
		t78 := int32(m.memory[int64(uint32(v3))+56])
		if t78 != i32(1) {
			goto l15
		}
		v5 = i32(0)
		t79 := int32(load32(m.memory[uint32(v4):]))
		v1 = t79
		store32(m.memory[int64(uint32(v1))+44:], uint32(i32(0)))
		t80 := int64(load64(m.memory[int64(uint32(v6))+16:]))
		v7 = t80
		t81 := int64(load64(m.memory[int64(uint32(v6))+8:]))
		v8 = t81
		t82 := int64(load64(m.memory[uint32(v6):]))
		v9 = t82
		store64(m.memory[uint32(v1):], uint64(i64(1)))
		store64(m.memory[int64(uint32(v1))+8:], uint64(v9))
		store64(m.memory[int64(uint32(v1))+16:], uint64(v8))
		store64(m.memory[int64(uint32(v1))+24:], uint64(v7))
		v17 = i32(0)
		{
			{
				t83 := int32(m.memory[int64(uint32(v3))+60])
				if t83 != 0 {
					goto l20
				}
				v10 = v3 + i32(64)
				v1 = i32(0)
				t84 := int32(load32(m.memory[uint32(v4):]))
				v11 = t84
				v12 = v11 + i32(56)
				v13 = v11 + i32(52)
				v14 = v11 + i32(36)
				v15 = v11 + i32(40)
				v16 = i32(0)
				{
				l23:
					{
						m.fn920(v2+i32(116), v10)
						{
							t85 := int32(load32(m.memory[int64(uint32(v2))+116:]))
							if t85 == 0 {
								goto l21
							}
							m.memory[int64(uint32(v3))+60] = byte(i32(2))
							t86 := int64(load64(m.memory[int64(uint32(v2))+120:]))
							t87 := m.fn921(t86)
							v1 = t87
							goto l22
						}
					l21:
						t88 := int32(load32(m.memory[int64(uint32(v3))+92:]))
						v17 = t88
						t89 := int32(load32(m.memory[int64(uint32(v2))+124:]))
						v5 = t89
						t90 := int32(load32(m.memory[int64(uint32(v2))+120:]))
						v18 = t90
						t91 := int32(load32(m.memory[uint32(v13):]))
						t92 := int32(load32(m.memory[uint32(v12):]))
						m.fn922(v2+i32(16), t91, t92, v16)
						t93 := int32(load32(m.memory[int64(uint32(v2))+20:]))
						v19 = t93
						t94 := int32(load32(m.memory[int64(uint32(v2))+16:]))
						v20 = t94
						t95 := int32(load32(m.memory[uint32(v14):]))
						t96 := int32(load32(m.memory[uint32(v15):]))
						m.fn923(v2+i32(8), t95, t96, v1)
						t97 := int32(load32(m.memory[int64(uint32(v2))+8:]))
						t98 := int32(load32(m.memory[int64(uint32(v2))+12:]))
						m.fn924(v2+i32(128), v17, v18, v5, v20, v19, t97, t98)
						t99 := int32(m.memory[int64(uint32(v2))+136])
						v5 = t99
						t100 := int32(load32(m.memory[int64(uint32(v2))+132:]))
						v18 = t100
						t101 := int32(load32(m.memory[int64(uint32(v2))+140:]))
						v19 = t101
						t102 := int64(load64(m.memory[int64(uint32(v3))+32:]))
						t103 := int32(load32(m.memory[int64(uint32(v2))+128:]))
						t104 := v3
						v20 = t103
						store64(m.memory[int64(uint32(t104))+32:], uint64(t102+int64(uint32(v20))))
						t105 := int32(load32(m.memory[int64(uint32(v3))+76:]))
						t106 := v3
						v21 = t105
						t107 := int32(load32(m.memory[int64(uint32(v3))+72:]))
						t108 := v21
						v20 = v20 + t107
						p109 := v20
						if uint32(v21) < uint32(v20) {
							p109 = t108
						}
						store32(m.memory[int64(uint32(t106))+72:], uint32(p109))
						t110 := int64(load64(m.memory[uint32(v17):]))
						_ = m.fn925(v6, t110)
						v1 = v19 + v1
						v16 = v18 + v16
						switch v5 {
						case 4:
							goto l27
						default:
							goto l23
						case 1:
							m.fn926(v11)
							goto l23
						case 2:
							m.fn927(v11)
							goto l23
						case 3:
						}
					}
					t112 := int32(load32(m.memory[uint32(v4):]))
					store32(m.memory[int64(uint32(t112))+44:], uint32(v1))
					t113 := int64(load64(m.memory[int64(uint32(v3))+48:]))
					v7 = t113
					if v7 == i64(-1) {
						m.fn153(i32(1078240))
						panic("unreachable")
					}
					store64(m.memory[int64(uint32(v3))+48:], uint64(v7+i64(1)))
					v17 = i32(0)
					v5 = i32(1)
					{
						t114 := int32(m.memory[int64(uint32(v3))+57])
						if t114 == 0 {
							goto l29
						}
						goto l20
					}
				l29:
					t115 := int64(load64(m.memory[uint32(v3):]))
					if t115 != i64(1) {
						store64(m.memory[uint32(v3):], uint64(i64(1)))
						t125 := int32(load32(m.memory[uint32(v4):]))
						t126 := int64(load32(m.memory[int64(uint32(t125))+44:]))
						store64(m.memory[int64(uint32(v3))+8:], uint64(t126))
						goto l20
					}
					t116 := int64(load64(m.memory[int64(uint32(v3))+8:]))
					v7 = t116
					t117 := int32(load32(m.memory[uint32(v4):]))
					t118 := v7
					v16 = t117
					t119 := int64(load32(m.memory[int64(uint32(v16))+44:]))
					v8 = t119
					if t118 == v8 {
						goto l20
					}
					v9 = i64(0)
					{
						t120 := int64(load64(m.memory[uint32(v16):]))
						if t120 != i64(1) {
							goto l31
						}
						t121 := int64(load64(m.memory[int64(uint32(v16))+24:]))
						store64(m.memory[int64(uint32(v2))+160:], uint64(t121))
						t122 := int64(load64(m.memory[int64(uint32(v16))+16:]))
						store64(m.memory[int64(uint32(v2))+152:], uint64(t122))
						t123 := int64(load64(m.memory[int64(uint32(v16))+8:]))
						store64(m.memory[int64(uint32(v2))+144:], uint64(t123))
						v9 = i64(1)
					}
				l31:
					store64(m.memory[int64(uint32(v2))+176:], uint64(v8))
					store64(m.memory[int64(uint32(v2))+168:], uint64(v7))
					store64(m.memory[int64(uint32(v2))+136:], uint64(v9))
					store64(m.memory[int64(uint32(v2))+128:], uint64(i64(4)))
					t124 := m.fn928(v2 + i32(128))
					v1 = t124
				}
			l22:
				v17 = i32(1)
				goto l20
			l27:
				m.memory[int64(uint32(v3))+60] = byte(i32(1))
				v5 = i32(0)
				v17 = i32(0)
			}
		l20:
			t127 := int32(m.memory[int64(uint32(v3))+61])
			if uint32(t127) < uint32(i32(2)) {
				goto l14
			}
			m.fn930(v4)
			goto l14
		}
	}
l15:
	v17 = i32(0)
	{
		t128 := int32(m.memory[int64(uint32(v3))+61])
		if uint32(t128) <= uint32(i32(1)) {
			goto l32
		}
		m.fn930(v4)
	}
l32:
	goto l14
l4:
	v17 = i32(1)
l14:
	m.fn931(v2+i32(116), v4)
	v16 = i32(1)
	{
		{
			t129 := int32(load32(m.memory[int64(uint32(v2))+116:]))
			if t129 == i32(1) {
				goto l33
			}
			v16 = v17
			goto l34
		}
	l33:
		t130 := int32(load32(m.memory[uint32(v4):]))
		store32(m.memory[int64(uint32(t130))+44:], uint32(i32(0)))
		if v17 != 0 {
			goto l34
		}
		t131 := int64(load64(m.memory[int64(uint32(v2))+120:]))
		v7 = t131
		t132 := int64(load64(m.memory[int64(uint32(v2))+56:]))
		store64(m.memory[int64(uint32(v2))+160:], uint64(t132))
		t133 := int64(load64(m.memory[int64(uint32(v2))+48:]))
		store64(m.memory[int64(uint32(v2))+152:], uint64(t133))
		t134 := int64(load64(m.memory[int64(uint32(v2))+40:]))
		store64(m.memory[int64(uint32(v2))+144:], uint64(t134))
		store64(m.memory[int64(uint32(v2))+168:], uint64(v7))
		store64(m.memory[int64(uint32(v2))+136:], uint64(i64(1)))
		store64(m.memory[int64(uint32(v2))+128:], uint64(i64(3)))
		t135 := m.fn928(v2 + i32(128))
		v1 = t135
	}
l34:
	{
		t136 := int32(m.memory[int64(uint32(v3))+61])
		if uint32(t136) < uint32(i32(2)) {
			goto l35
		}
		m.fn932(v4)
	}
l35:
	v3 = i32(1)
	{
		if v16 != 0 {
			goto l36
		}
		if v5&i32(1) != 0 {
			goto l37
		}
		v3 = i32(2)
		goto l36
	l37:
		t137 := m.fn934()
		v1 = t137
		v7 = i64(0)
		{
			t138 := int32(load32(m.memory[uint32(v4):]))
			v3 = t138
			t139 := int64(load64(m.memory[uint32(v3):]))
			if t139 != i64(1) {
				goto l38
			}
			t140 := int64(load64(m.memory[int64(uint32(v3))+24:]))
			store64(m.memory[int64(uint32(v2))+144:], uint64(t140))
			t141 := int64(load64(m.memory[int64(uint32(v3))+16:]))
			store64(m.memory[int64(uint32(v2))+136:], uint64(t141))
			t142 := int64(load64(m.memory[int64(uint32(v3))+8:]))
			store64(m.memory[int64(uint32(v2))+128:], uint64(t142))
			v7 = i64(1)
		}
	l38:
		store64(m.memory[uint32(v1):], uint64(v7))
		t143 := int64(load64(m.memory[int64(uint32(v2))+128:]))
		store64(m.memory[int64(uint32(v1))+8:], uint64(t143))
		t144 := int64(load64(m.memory[int64(uint32(v2))+136:]))
		store64(m.memory[int64(uint32(v1))+16:], uint64(t144))
		t145 := int64(load64(m.memory[int64(uint32(v2))+144:]))
		store64(m.memory[int64(uint32(v1))+24:], uint64(t145))
		t146 := v2 + i32(128)
		v16 = v3 + i32(32)
		m.fn935(t146, v16)
		t147 := int32(load32(m.memory[int64(uint32(v1))+32:]))
		t148 := int32(load32(m.memory[uint32(v1+i32(36)):]))
		m.fn188(t147, t148)
		t149 := int64(load64(m.memory[int64(uint32(v2))+136:]))
		store64(m.memory[int64(uint32(v1))+40:], uint64(t149))
		t150 := int64(load64(m.memory[int64(uint32(v2))+128:]))
		store64(m.memory[int64(uint32(v1))+32:], uint64(t150))
		t151 := m.fn936(v16)
		v16 = t151
		t152 := int32(load32(m.memory[uint32(v3+i32(52)):]))
		t153 := int32(load32(m.memory[uint32(v3+i32(56)):]))
		m.fn518(v2, t152, t153, v16, i32(1079700))
		t154 := int32(load32(m.memory[uint32(v2):]))
		t155 := int32(load32(m.memory[int64(uint32(v2))+4:]))
		m.fn51(v2+i32(128), t154, t155)
		t156 := int32(load32(m.memory[int64(uint32(v1))+48:]))
		t157 := int32(load32(m.memory[uint32(v1+i32(52)):]))
		m.fn16(t156, t157)
		t158 := int32(load32(m.memory[int64(uint32(v2))+136:]))
		store32(m.memory[int64(uint32(v1))+56:], uint32(t158))
		t159 := int64(load64(m.memory[int64(uint32(v2))+128:]))
		store64(m.memory[int64(uint32(v1))+48:], uint64(t159))
		v3 = i32(0)
	}
l36:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
	store32(m.memory[uint32(v0):], uint32(v3))
	m.g0 = v2 + i32(192)
}
func (m *Module) fn918(v0 int32) {
	var v1 int64
	{
		t0 := int64(load64(m.memory[uint32(v0):]))
		v1 = t0
		p1 := i32(5)
		if uint64(v1) > uint64(i64(1)) {
			p1 = int32(v1) + i32(-2)
		}
		switch p1 {
		case 1, 2, 3:
			goto l1
		default:
			t2 := int32(m.memory[int64(uint32(v0))+48])
			if uint32(t2) > uint32(i32(1)) {
				goto l1
			}
			t3 := int32(load32(m.memory[int64(uint32(v0))+52:]))
			t4 := int32(load32(m.memory[uint32(v0+i32(56)):]))
			m.fn16(t3, t4)
			goto l1
		case 0:
			t5 := int32(m.memory[int64(uint32(v0))+8])
			t6 := int32(load32(m.memory[uint32(v0+i32(12)):]))
			m.fn119(t5, t6)
			goto l1
		case 4:
			t7 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			t8 := int32(load32(m.memory[uint32(v0+i32(12)):]))
			m.fn16(t7, t8)
		}
	}
l1:
	m.fn10(v0, i32(64), i32(8))
}
func (m *Module) fn919(v0 int32) {
	t0 := int32(load32(m.memory[int64(uint32(v0))+48:]))
	t1 := int32(load32(m.memory[uint32(v0+i32(52)):]))
	m.fn16(t0, t1)
	t2 := int32(load32(m.memory[int64(uint32(v0))+32:]))
	t3 := int32(load32(m.memory[uint32(v0+i32(36)):]))
	m.fn188(t2, t3)
	m.fn10(v0, i32(64), i32(8))
}
func (m *Module) fn920(v0, v1 int32) {
	var v2, v3, v4, v5, v6 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	t1 := int32(load32(m.memory[uint32(v1):]))
	v3 = t1
	{
		t2 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		v4 = t2
		t3 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		t4 := v4
		v5 = t3
		if uint32(t4) < uint32(v5) {
			goto l0
		}
		t5 := int32(load32(m.memory[int64(uint32(v1))+20:]))
		t6 := int32(load32(m.memory[int64(uint32(v1))+24:]))
		t7 := v2
		v5 = t6
		t8 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t9 := v5
		t10 := v5
		v4 = t8
		p11 := v4
		if uint32(v5) < uint32(v4) {
			p11 = t10
		}
		m.fn309(t7, t5, t9, p11, i32(1287064))
		t12 := int32(load32(m.memory[int64(uint32(v2))+4:]))
		t13 := v4
		v5 = t12
		if uint32(t13) < uint32(v5) {
			m.fn256(i32(1286600), i32(46), i32(1286648))
			panic("unreachable")
		}
		t14 := int32(load32(m.memory[int64(uint32(v2))+12:]))
		v4 = t14
		t15 := int32(load32(m.memory[int64(uint32(v2))+8:]))
		v6 = t15
		t16 := int32(load32(m.memory[uint32(v2):]))
		m.fn310(v3, v5, t16, v5, i32(1286664))
		store32(m.memory[int64(uint32(v1))+24:], uint32(v4))
		store32(m.memory[int64(uint32(v1))+20:], uint32(v6))
		store32(m.memory[int64(uint32(v1))+12:], uint32(v5))
		v4 = i32(0)
		store32(m.memory[int64(uint32(v1))+8:], uint32(i32(0)))
	}
l0:
	store32(m.memory[uint32(v0):], uint32(i32(0)))
	store32(m.memory[int64(uint32(v0))+8:], uint32(v5-v4))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v3+v4))
	m.g0 = v2 + i32(16)
}
func (m *Module) fn921(v0 int64) int32 {
	var v1, v2 int32
	t0 := m.g0
	v1 = t0 - i32(64)
	m.g0 = v1
	store64(m.memory[uint32(v1):], uint64(i64(2)))
	store64(m.memory[int64(uint32(v1))+8:], uint64(v0))
	t1 := m.fn928(v1)
	v2 = t1
	m.g0 = v1 + i32(64)
	return v2
}
func (m *Module) fn922(v0, v1, v2, v3 int32) {
	var v4 int32
	t0 := m.g0
	v4 = t0 - i32(16)
	m.g0 = v4
	m.fn212(v4+i32(8), v3, v1, v2, i32(1077872))
	t1 := int32(load32(m.memory[int64(uint32(v4))+12:]))
	v2 = t1
	t2 := int32(load32(m.memory[int64(uint32(v4))+8:]))
	store32(m.memory[uint32(v0):], uint32(t2))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v2))
	m.g0 = v4 + i32(16)
}
func (m *Module) fn923(v0, v1, v2, v3 int32) {
	if uint32(v2) >= uint32(v3) {
		goto l0
	}
	m.fn151(v3, v2, v2, i32(1077888))
	panic("unreachable")
l0:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v2-v3))
	store32(m.memory[uint32(v0):], uint32(v1+v3<<2))
}
func (m *Module) fn924(v0, v1, v2, v3, v4, v5, v6, v7 int32) {
	var v8, v9, v10, v11, v12, v13, v14, v15, v16, v17 int32
	var v18 int64
	var v19, v20, v21, v22, v23, v24, v25, v26 int32
	v8 = i32(0)
	{
		if uint32(v3) < uint32(i32(3)) {
			goto l0
		}
		t0 := int32(m.memory[int64(uint32(v1))+429])
		if t0&i32(1) != 0 {
			goto l0
		}
		t1 := int32(load16(m.memory[uint32(v2):]))
		t2 := int32(m.memory[uint32(v2+i32(2))])
		if t1|t2<<16 != i32(0xbfbbef) {
			goto l0
		}
		v8 = i32(3)
		v2 = v2 + i32(3)
		v3 = v3 + i32(-3)
	}
l0:
	{
		{
			t3 := int32(m.memory[int64(uint32(v1))+428])
			if t3 != 0 {
				{
					if v3 == 0 {
						{
							t48 := int32(m.memory[int64(uint32(v1))+419])
							v10 = t48
							if uint32(v10) <= uint32(i32(9)) {
								goto l45
							}
							switch v10 + i32(-200) {
							case 2:
								goto l47
							default:
								goto l46
							}
						}
					l45:
						if i32_shl(i32(1), v10)&i32(190) != 0 {
							goto l46
						}
					l47:
						m.memory[int64(uint32(v1))+419] = byte(i32(202))
						v12 = i32(4)
						goto l48
					l46:
						if v7 != 0 {
							m.memory[int64(uint32(v1))+419] = byte(i32(8))
							t49 := int32(load32(m.memory[int64(uint32(v1))+8:]))
							v10 = t49
							v11 = i32(0)
							store32(m.memory[int64(uint32(v1))+8:], uint32(i32(0)))
							store32(m.memory[uint32(v6):], uint32(v10))
							goto l50
						}
						v12 = i32(2)
					l48:
						v9 = i32(0)
						goto l6
					}
					v9 = i32(0)
					if v5 == 0 {
						goto l4
					}
					if v7 == 0 {
						goto l7
					}
					t30 := int32(m.memory[int64(uint32(v1))+427])
					v17 = t30
					t31 := int32(m.memory[int64(uint32(v1))+426])
					v16 = v17 & t31
					t32 := int32(m.memory[int64(uint32(v1))+424])
					v23 = t32
					t33 := int32(m.memory[int64(uint32(v1))+421])
					v21 = t33
					t34 := int32(m.memory[int64(uint32(v1))+420])
					v13 = t34
					t35 := int32(m.memory[int64(uint32(v1))+417])
					v20 = t35
					t36 := int32(m.memory[int64(uint32(v1))+418])
					v14 = t36
					t37 := int32(m.memory[int64(uint32(v1))+419])
					v12 = t37
					t38 := int32(load32(m.memory[int64(uint32(v1))+8:]))
					v11 = t38
					t39 := int32(m.memory[int64(uint32(v1))+425])
					v24 = t39 & i32(255)
					t40 := int32(m.memory[int64(uint32(v1))+422])
					v25 = t40 & i32(1)
					t41 := int32(m.memory[int64(uint32(v1))+423])
					v26 = t41 & i32(255)
					v9 = i32(0)
					v10 = i32(0)
				l31:
					if uint32(v10) >= uint32(v3) {
						goto l17
					}
					if uint32(v11) >= uint32(v5) {
						goto l17
					}
					if uint32(v9) < uint32(v7) {
						t42 := int32(m.memory[uint32(v2+v10)])
						v15 = t42
						v19 = i32(3)
						v22 = v12 & i32(255)
						switch v22 {
						case 4:
							goto l24
						case 5:
							t46 := v16
							var p47 int32
							if v14&i32(255) == v15&i32(255) {
								p47 = 1
							}
							if t46&p47 != 0 {
								goto l24
							}
							if v20&i32(255) == v15 {
								goto l39
							}
							if v13&i32(1) == 0 {
								v12 = i32(200)
								v19 = i32(2)
								switch v15 + i32(-10) {
								case 0, 3:
									goto l31
								default:
									goto l24
								}
							}
							v19 = i32(2)
							v12 = i32(200)
							if v15 == v21&i32(255) {
								goto l31
							}
							goto l24
						case 6:
							p43 := i32(6)
							if v15 == i32(10) {
								p43 = i32(0)
							}
							v19 = p43
							goto l36
						default:
							v19 = i32(202)
							v12 = i32(201)
							switch v22 + i32(-200) {
							case 1:
								p44 := i32(9)
								if v13&i32(255) != 0 {
									p44 = i32(8)
								}
								p45 := i32(8)
								if v15 == i32(13) {
									p45 = p44
								}
								v19 = p45
								goto l43
							case 2:
								goto l33
							default:
								goto l31
							}
						case 0:
							if v13&i32(1) == 0 {
								v19 = i32(0)
								switch v15 + i32(-10) {
								case 0, 3:
									goto l36
								default:
									goto l35
								}
							}
							if v15 != v21&i32(255) {
								goto l35
							}
							v19 = i32(0)
							goto l36
						case 1:
							if v17&i32(1) == 0 {
								goto l37
							}
							if v14&i32(255) != v15 {
								goto l37
							}
							v19 = i32(3)
							goto l36
						case 2:
							if v20&i32(255) != v15 {
								if v13&i32(1) == 0 {
									v12 = i32(200)
									v19 = i32(2)
									switch v15 + i32(-10) {
									case 0, 3:
										goto l31
									default:
										goto l24
									}
								}
								v19 = i32(2)
								v12 = i32(200)
								if v15 == v21&i32(255) {
									goto l31
								}
								goto l24
							}
							goto l39
						case 3:
							if v17&i32(1) == 0 {
								goto l24
							}
							if v14&i32(255) != v15 {
								if v25 == 0 {
									goto l24
								}
								if v26 != v15 {
									goto l24
								}
								v19 = i32(4)
								goto l36
							}
							v19 = i32(5)
							goto l36
						case 7:
							v12 = i32(1)
							goto l31
						case 8:
							v12 = i32(0)
							goto l31
						case 9:
							v12 = i32(0)
							v19 = i32(0)
							if v15 != i32(10) {
								goto l31
							}
							goto l36
						}
					l35:
						v12 = i32(1)
						if v23&i32(1) == 0 {
							goto l31
						}
						if v24 != v15 {
							goto l31
						}
						v19 = i32(6)
						goto l36
					l37:
						if v20&i32(255) == v15 {
							goto l39
						}
						if v13&i32(1) == 0 {
							v12 = i32(200)
							v19 = i32(2)
							switch v15 + i32(-10) {
							case 0, 3:
								goto l31
							default:
								goto l24
							}
						}
						v19 = i32(2)
						v12 = i32(200)
						if v15 == v21&i32(255) {
							goto l31
						}
						goto l24
					l24:
						m.memory[uint32(v4+v11)] = byte(v15)
						v11 = v11 + i32(1)
					l36:
						v10 = v10 + i32(1)
						v12 = v19
						goto l31
					l39:
						v19 = i32(7)
					l43:
						v10 = v10 + i32(1)
					l33:
						store32(m.memory[uint32(v6+v9<<2):], uint32(v11))
						v9 = v9 + i32(1)
						v12 = i32(7)
						if v19 == i32(7) {
							goto l31
						}
						goto l19
					}
				l17:
					v19 = v12
					goto l19
				}
			l19:
				{
					v15 = v19 & i32(255)
					if uint32(v15+i32(-8)) < uint32(i32(2)) {
						goto l51
					}
					v12 = i32(4)
					v2 = v11
					if v15 == i32(202) {
						goto l52
					}
					var p50 int32
					if uint32(v10) < uint32(v3) {
						p50 = 1
					}
					v12 = p50
					t52 := v12
					t53 := v12
					p51 := i32(2)
					if uint32(v9) < uint32(v7) {
						p51 = i32(0)
					}
					p54 := p51
					if uint32(v10) >= uint32(v3) {
						p54 = t53
					}
					p55 := p54
					if uint32(v11) >= uint32(v5) {
						p55 = t52
					}
					v12 = p55
					v2 = v11
					goto l52
				}
			l51:
				v2 = i32(0)
				v12 = i32(3)
			l52:
				store32(m.memory[int64(uint32(v1))+8:], uint32(v2))
				m.memory[int64(uint32(v1))+419] = byte(v19)
				goto l53
			}
			if v3 == 0 {
				v9 = i32(0)
				t4 := int32(load32(m.memory[int64(uint32(v1))+268:]))
				t5 := int32(m.memory[int64(uint32(v1))+416])
				v10 = t5
				t6 := int32(m.memory[int64(uint32(v1))+345])
				t7 := v10
				v11 = t6
				p8 := i32(8)
				if uint32(t7) >= uint32(v11) {
					p8 = i32(0)
				}
				p9 := i32(0)
				if v10 != 0 {
					p9 = p8
				}
				t10 := m.fn748(t4, p9)
				v10 = t10
				v12 = v10 & i32(255)
				if uint32(v12) >= uint32(v11) {
					goto l5
				}
				m.memory[int64(uint32(v1))+416] = byte(v10)
				v9 = i32(0)
				p11 := i32(4)
				if v12 != 0 {
					p11 = i32(0)
				}
				v12 = p11
				goto l6
			}
			v9 = i32(0)
			if v5 != 0 {
				if v7 == 0 {
					goto l7
				}
				v13 = v1 + i32(346)
				v14 = v1 + i32(272)
				v15 = v1 + i32(12)
				t12 := int32(load32(m.memory[int64(uint32(v1))+8:]))
				v16 = t12
				t13 := int32(m.memory[int64(uint32(v1))+344])
				v17 = t13
				t14 := int64(load64(m.memory[uint32(v1):]))
				v18 = t14
				t15 := int32(m.memory[int64(uint32(v1))+416])
				v19 = t15
				v11 = i32(0)
				t16 := int32(m.memory[int64(uint32(v1))+342])
				v20 = t16 & i32(255)
				t17 := int32(m.memory[int64(uint32(v1))+343])
				v21 = t17 & i32(255)
				v10 = i32(0)
				v9 = i32(0)
			l14:
				if uint32(v10) >= uint32(v3) {
					goto l8
				}
				if uint32(v11) >= uint32(v5) {
					goto l8
				}
				if uint32(v9) >= uint32(v7) {
					goto l8
				}
				{
					t18 := int32(m.memory[uint32(v2+v10)])
					t19 := v15
					v22 = t18
					t20 := int32(m.memory[uint32(t19+v22)])
					v12 = t20 + v19&i32(255)
					if uint32(v12) >= uint32(i32(70)) {
						m.fn158(v12, i32(70), i32(1148616))
						panic("unreachable")
					}
					t21 := v1
					t22 := v18
					var p23 int32
					if v22 == i32(10) {
						p23 = 1
					}
					v18 = t22 + int64(uint32(p23))
					store64(m.memory[uint32(t21):], uint64(v18))
					t24 := int32(m.memory[uint32(v14+v12)])
					v19 = t24
					t25 := int32(m.memory[uint32(v13+v12)])
					if t25 != 0 {
						goto l10
					}
					goto l11
				}
			l10:
				m.memory[uint32(v4+v11)] = byte(v22)
				v11 = v11 + i32(1)
			l11:
				v10 = v10 + i32(1)
				{
					v12 = v19 & i32(255)
					t26 := v12
					v22 = v17 & i32(255)
					if uint32(t26) < uint32(v22) {
						goto l12
					}
					store32(m.memory[uint32(v6+v9<<2):], uint32(v16+v11))
					v9 = v9 + i32(1)
					if uint32(v12) > uint32(v22) {
						goto l8
					}
				}
			l12:
				if v12 == v20 {
					goto l15
				}
				if v12 != v21 {
					goto l14
				}
			l15:
				{
					if uint32(v10) >= uint32(v3) {
						goto l14
					}
					if uint32(v11) >= uint32(v5) {
						goto l14
					}
					t27 := int32(m.memory[uint32(v2+v10)])
					t28 := v15
					v12 = t27
					t29 := int32(m.memory[uint32(t28+v12)])
					if t29 != 0 {
						goto l14
					}
					m.memory[uint32(v4+v11)] = byte(v12)
					v11 = v11 + i32(1)
					v10 = v10 + i32(1)
					goto l15
				}
			}
			goto l4
		}
	l5:
		if v7 == 0 {
			goto l7
		}
		m.memory[int64(uint32(v1))+416] = byte(v10)
		t56 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		v10 = t56
		v11 = i32(0)
		store32(m.memory[int64(uint32(v1))+8:], uint32(i32(0)))
		store32(m.memory[uint32(v6):], uint32(v10))
	}
l50:
	v12 = i32(3)
	v9 = i32(1)
	goto l54
l8:
	{
		t57 := int32(m.memory[int64(uint32(v1))+345])
		if uint32(v19&i32(255)) >= uint32(t57) {
			store32(m.memory[int64(uint32(v1))+8:], uint32(i32(0)))
			m.memory[int64(uint32(v1))+416] = byte(v19)
			v12 = i32(3)
			goto l53
		}
		m.memory[int64(uint32(v1))+416] = byte(v19)
		store32(m.memory[int64(uint32(v1))+8:], uint32(v16+v11))
		var p58 int32
		if uint32(v10) < uint32(v3) {
			p58 = 1
		}
		v12 = p58
		t60 := v12
		t61 := v12
		p59 := i32(2)
		if uint32(v9) < uint32(v7) {
			p59 = i32(0)
		}
		p62 := p59
		if uint32(v10) >= uint32(v3) {
			p62 = t61
		}
		p63 := p62
		if uint32(v11) >= uint32(v5) {
			p63 = t60
		}
		v12 = p63
		goto l53
	}
l7:
	v12 = i32(2)
	goto l6
l4:
	v12 = i32(1)
l6:
	v11 = i32(0)
l54:
	v10 = i32(0)
l53:
	m.memory[int64(uint32(v0))+8] = byte(v12)
	m.memory[int64(uint32(v1))+429] = byte(i32(1))
	store32(m.memory[int64(uint32(v0))+12:], uint32(v9))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v11))
	store32(m.memory[uint32(v0):], uint32(v10+v8))
}
func (m *Module) fn925(v0 int32, v1 int64) int32 {
	if v1 != i64(0) {
		store64(m.memory[int64(uint32(v0))+8:], uint64(v1))
		return v0
	}
	m.fn256(i32(1079769), i32(26), i32(1079796))
	panic("unreachable")
}
func (m *Module) fn926(v0 int32) {
	t0 := int32(load32(m.memory[int64(uint32(v0))+56:]))
	t1 := v0 + i32(48)
	v0 = t0 << 1
	p2 := i32(4)
	if uint32(v0) > uint32(i32(4)) {
		p2 = v0
	}
	m.fn482(t1, p2)
}
func (m *Module) fn927(v0 int32) {
	t0 := int32(load32(m.memory[int64(uint32(v0))+40:]))
	t1 := v0 + i32(32)
	v0 = t0 << 1
	p2 := i32(4)
	if uint32(v0) > uint32(i32(4)) {
		p2 = v0
	}
	m.fn939(t1, p2)
}
func (m *Module) fn928(v0 int32) int32 {
	var v1 int32
	t0 := m.fn1679(i32(64))
	v1 = t0
	memory_copy(m.memory, uint32(v1), uint32(v0), uint32(i32(64)))
	return v1
}
func (m *Module) fn929(v0 int32) int32 {
	var v1, v2 int32
	var v3 int64
	var v4, v5, v6 int32
	t0 := m.g0
	v1 = t0 - i32(112)
	m.g0 = v1
	m.fn247(v1+i32(16), i32(8), i32(64))
	{
		t1 := int32(load32(m.memory[int64(uint32(v1))+16:]))
		v2 = t1
		if v2 == 0 {
			m.fn85(i32(8), i32(64))
			panic("unreachable")
		}
		v3 = i64(0)
		{
			t2 := int64(load64(m.memory[uint32(v0):]))
			if t2 != i64(1) {
				goto l1
			}
			t3 := int64(load64(m.memory[int64(uint32(v0))+24:]))
			store64(m.memory[int64(uint32(v1))+104:], uint64(t3))
			t4 := int64(load64(m.memory[int64(uint32(v0))+16:]))
			store64(m.memory[int64(uint32(v1))+96:], uint64(t4))
			t5 := int64(load64(m.memory[int64(uint32(v0))+8:]))
			store64(m.memory[int64(uint32(v1))+88:], uint64(t5))
			v3 = i64(1)
		}
	l1:
		{
			t6 := int32(load32(m.memory[uint32(v0+i32(56)):]))
			v4 = t6
			if v4 == 0 {
				goto l2
			}
			t7 := int32(load32(m.memory[uint32(v0+i32(52)):]))
			v5 = t7
			m.fn940(v1+i32(8), v4)
			t8 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			v6 = t8
			if v6 == 0 {
				m.fn2(i32(1), v4)
				panic("unreachable")
			}
			store32(m.memory[int64(uint32(v1))+80:], uint32(i32(0)))
			store32(m.memory[int64(uint32(v1))+76:], uint32(v6))
			store32(m.memory[int64(uint32(v1))+72:], uint32(v4))
			if v4 == 0 {
				goto l4
			}
			memory_copy(m.memory, uint32(v6), uint32(v5), uint32(v4))
		l4:
			store32(m.memory[int64(uint32(v1))+80:], uint32(v4))
			goto l5
		}
	l2:
		store32(m.memory[int64(uint32(v1))+80:], uint32(i32(0)))
		store64(m.memory[int64(uint32(v1))+72:], uint64(i64(0x100000000)))
	l5:
		m.fn935(v1+i32(24)+i32(32), v0+i32(32))
		store64(m.memory[int64(uint32(v1))+24:], uint64(v3))
		t9 := int64(load64(m.memory[int64(uint32(v1))+88:]))
		store64(m.memory[int64(uint32(v1))+32:], uint64(t9))
		t10 := int64(load64(m.memory[int64(uint32(v1))+96:]))
		store64(m.memory[int64(uint32(v1))+40:], uint64(t10))
		t11 := int64(load64(m.memory[int64(uint32(v1))+104:]))
		store64(m.memory[int64(uint32(v1))+48:], uint64(t11))
		memory_copy(m.memory, uint32(v2), uint32(v1+i32(24)), uint32(i32(64)))
		m.g0 = v1 + i32(112)
		return v2
	}
}
func (m *Module) fn930(v0 int32) {
	var v1, v2, v3 int32
	var v4 int64
	var v5, v6, v7, v8 int32
	t0 := m.g0
	v1 = t0 - i32(32)
	m.g0 = v1
	{
		t1 := int32(load32(m.memory[uint32(v0):]))
		v2 = t1
		t2 := int32(load32(m.memory[int64(uint32(v2))+44:]))
		if t2 == 0 {
			goto l0
		}
		t3 := m.fn1682(v2)
		t4 := int32(load32(m.memory[int64(uint32(v2))+44:]))
		t5 := m.fn1686(t3, t4)
		v3 = t5
		v4 = i64(0)
		{
			t6 := int64(load64(m.memory[uint32(v2):]))
			if t6 != i64(1) {
				goto l1
			}
			t7 := int64(load64(m.memory[int64(uint32(v2))+24:]))
			store64(m.memory[int64(uint32(v1))+24:], uint64(t7))
			t8 := int64(load64(m.memory[int64(uint32(v2))+16:]))
			store64(m.memory[int64(uint32(v1))+16:], uint64(t8))
			t9 := int64(load64(m.memory[int64(uint32(v2))+8:]))
			store64(m.memory[int64(uint32(v1))+8:], uint64(t9))
			v4 = i64(1)
		}
	l1:
		store64(m.memory[uint32(v3):], uint64(v4))
		t10 := int64(load64(m.memory[int64(uint32(v1))+8:]))
		store64(m.memory[int64(uint32(v3))+8:], uint64(t10))
		t11 := int64(load64(m.memory[int64(uint32(v1))+16:]))
		store64(m.memory[int64(uint32(v3))+16:], uint64(t11))
		t12 := int64(load64(m.memory[int64(uint32(v1))+24:]))
		store64(m.memory[int64(uint32(v3))+24:], uint64(t12))
		m.fn1681(v1+i32(8), v0)
	l7:
		m.fn1683(v1, v1+i32(8))
		{
			t13 := int32(load32(m.memory[uint32(v1):]))
			v5 = t13
			if v5 == 0 {
				m.fn1687(v2)
				store32(m.memory[uint32(v0):], uint32(v3))
				goto l0
			}
			t14 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			v6 = t14
			v7 = v5 + i32(-1)
		l5:
			if v6 != 0 {
				t15 := int32(m.memory[uint32(v7+v6)])
				v8 = t15 + i32(-9)
				if uint32(v8) > uint32(i32(23)) {
					goto l4
				}
				if i32_shl(i32(1), v8)&i32(8388635) == 0 {
					goto l4
				}
				v6 = v6 + i32(-1)
				goto l5
			}
			v6 = i32(0)
			goto l4
		}
	l4:
		v7 = v5 + v6
	l9:
		if v6 != 0 {
			{
				t16 := int32(m.memory[uint32(v5)])
				v8 = t16 + i32(-9)
				if uint32(v8) > uint32(i32(23)) {
					goto l8
				}
				if i32_shl(i32(1), v8)&i32(8388635) == 0 {
					goto l8
				}
				v5 = v5 + i32(1)
				v6 = v6 + i32(-1)
				goto l9
			}
		l8:
			m.fn1688(v3, v5, v6)
			goto l7
		}
		m.fn1688(v3, v7, i32(0))
		goto l7
	}
l0:
	m.g0 = v1 + i32(32)
}
func (m *Module) fn931(v0, v1 int32) {
	var v2, v3, v4, v5 int32
	t0 := m.g0
	v2 = t0 - i32(64)
	m.g0 = v2
	t1 := int32(load32(m.memory[uint32(v1):]))
	v3 = t1
	t2 := m.fn936(v3 + i32(32))
	v4 = t2
	t3 := int32(load32(m.memory[uint32(v3+i32(52)):]))
	t4 := int32(load32(m.memory[uint32(v3+i32(56)):]))
	m.fn518(v2+i32(16), t3, t4, v4, i32(1079716))
	t5 := int32(load32(m.memory[int64(uint32(v2))+16:]))
	v5 = t5
	{
		{
			t6 := int32(load32(m.memory[int64(uint32(v2))+20:]))
			v3 = t6
			if uint32(v3) > uint32(i32(3)) {
				goto l0
			}
			v5 = v5 + i32(-1)
		l2:
			{
				if v3 == 0 {
					goto l1
				}
				v4 = v5 + v3
				v3 = v3 + i32(-1)
				t7 := int32(int8(m.memory[uint32(v4)]))
				if t7 > i32(-1) {
					goto l2
				}
				goto l3
			}
		}
	l0:
		t8 := int32(load32(m.memory[uint32(v5):]))
		if t8&i32(-2139062144) != 0 {
			goto l3
		}
		v4 = (v5 + i32(3)) & i32(-4)
		p9 := v4 - v5
		if v4 == v5 {
			p9 = i32(4)
		}
		v4 = p9
		v3 = v3 + i32(-4)
	l5:
		{
			if uint32(v4) >= uint32(v3) {
				goto l4
			}
			t10 := int32(load32(m.memory[uint32(v5+v4):]))
			if t10&i32(-2139062144) != 0 {
				goto l3
			}
			v4 = v4 + i32(4)
			goto l5
		}
	l4:
		t11 := int32(load32(m.memory[uint32(v5+v3):]))
		if t11&i32(-2139062144) == 0 {
			goto l1
		}
	}
l3:
	m.fn938(v2+i32(28), v1)
	store32(m.memory[int64(uint32(v2))+48:], uint32(i32(0)))
	{
	l7:
		{
			m.fn890(v2+i32(8), v2+i32(28))
			t12 := int32(load32(m.memory[int64(uint32(v2))+8:]))
			v3 = t12
			if v3 == 0 {
				store32(m.memory[uint32(v0):], uint32(i32(0)))
				goto l8
			}
			t13 := int32(load32(m.memory[int64(uint32(v2))+12:]))
			v4 = t13
			t14 := int32(load32(m.memory[int64(uint32(v2))+48:]))
			t15 := v2
			v5 = t14
			store32(m.memory[int64(uint32(t15))+48:], uint32(v5+i32(1)))
			m.fn12(v2+i32(52), v3, v4)
			t16 := int32(load32(m.memory[int64(uint32(v2))+52:]))
			if t16 == 0 {
				goto l7
			}
		}
		t17 := int64(load64(m.memory[int64(uint32(v2))+56:]))
		store32(m.memory[int64(uint32(v0))+8:], uint32(t17))
		store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
		store32(m.memory[uint32(v0):], uint32(i32(1)))
		goto l8
	}
l1:
	store32(m.memory[uint32(v0):], uint32(i32(0)))
l8:
	m.g0 = v2 + i32(64)
}
func (m *Module) fn932(v0 int32) {
	var v1, v2, v3 int32
	var v4 int64
	var v5 int32
	t0 := m.g0
	v1 = t0 - i32(48)
	m.g0 = v1
	{
		t1 := int32(load32(m.memory[uint32(v0):]))
		v2 = t1
		t2 := int32(load32(m.memory[int64(uint32(v2))+44:]))
		if t2 == 0 {
			goto l0
		}
		t3 := m.fn1682(v2)
		t4 := int32(load32(m.memory[int64(uint32(v2))+44:]))
		t5 := m.fn1686(t3, t4)
		v3 = t5
		v4 = i64(0)
		{
			t6 := int64(load64(m.memory[uint32(v2):]))
			if t6 != i64(1) {
				goto l1
			}
			t7 := int64(load64(m.memory[int64(uint32(v2))+24:]))
			store64(m.memory[int64(uint32(v1))+40:], uint64(t7))
			t8 := int64(load64(m.memory[int64(uint32(v2))+16:]))
			store64(m.memory[int64(uint32(v1))+32:], uint64(t8))
			t9 := int64(load64(m.memory[int64(uint32(v2))+8:]))
			store64(m.memory[int64(uint32(v1))+24:], uint64(t9))
			v4 = i64(1)
		}
	l1:
		store64(m.memory[uint32(v3):], uint64(v4))
		t10 := int64(load64(m.memory[int64(uint32(v1))+24:]))
		store64(m.memory[int64(uint32(v3))+8:], uint64(t10))
		t11 := int64(load64(m.memory[int64(uint32(v1))+32:]))
		store64(m.memory[int64(uint32(v3))+16:], uint64(t11))
		t12 := int64(load64(m.memory[int64(uint32(v1))+40:]))
		store64(m.memory[int64(uint32(v3))+24:], uint64(t12))
		m.fn1681(v1+i32(24), v0)
	l3:
		{
			m.fn1683(v1+i32(16), v1+i32(24))
			t13 := int32(load32(m.memory[int64(uint32(v1))+16:]))
			v5 = t13
			if v5 == 0 {
				goto l2
			}
			t14 := int32(load32(m.memory[int64(uint32(v1))+20:]))
			m.fn46(v1+i32(8), v5, t14)
			t15 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			t16 := int32(load32(m.memory[int64(uint32(v1))+12:]))
			m.fn1688(v3, t15, t16)
			goto l3
		}
	l2:
		m.fn1687(v2)
		store32(m.memory[uint32(v0):], uint32(v3))
	}
l0:
	m.g0 = v1 + i32(48)
}
func (m *Module) fn933(v0 int32) {
	var v1 int32
	{
		t0 := int32(load32(m.memory[uint32(v0):]))
		v1 = t0
		if v1 == i32(2) {
			return
		}
		t1 := int32(load32(m.memory[int64(uint32(v0))+12:]))
		m.fn919(t1)
		if v1 != 0 {
			return
		}
		t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		m.fn919(t2)
	}
}
func (m *Module) fn934() int32 {
	var v0, v1 int32
	t0 := m.g0
	v0 = t0 - i32(80)
	m.g0 = v0
	m.fn321(v0+i32(8), i32(0))
	m.fn937(v0+i32(64), i32(0))
	store32(m.memory[int64(uint32(v0))+76:], uint32(i32(0)))
	t1 := int64(load64(m.memory[int64(uint32(v0))+64:]))
	store64(m.memory[int64(uint32(v0))+48:], uint64(t1))
	t2 := int64(load64(m.memory[int64(uint32(v0))+72:]))
	store64(m.memory[int64(uint32(v0))+56:], uint64(t2))
	t3 := m.fn113(i32(8), i32(64))
	v1 = t3
	store64(m.memory[uint32(v1):], uint64(i64(0)))
	memory_copy(m.memory, uint32(v1+i32(8)), uint32(v0+i32(24)), uint32(i32(40)))
	t4 := int32(load32(m.memory[int64(uint32(v0))+16:]))
	store32(m.memory[int64(uint32(v1))+56:], uint32(t4))
	t5 := int64(load64(m.memory[int64(uint32(v0))+8:]))
	store64(m.memory[int64(uint32(v1))+48:], uint64(t5))
	m.g0 = v0 + i32(80)
	return v1
}
func (m *Module) fn935(v0, v1 int32) {
	var v2, v3, v4, v5, v6 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	t1 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	v3 = t1
	t2 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	t3 := v2 + i32(8)
	v4 = t2
	m.fn59(t3, v4, i32(4), i32(4))
	t4 := int32(load32(m.memory[int64(uint32(v2))+8:]))
	v5 = t4
	t5 := int32(load32(m.memory[int64(uint32(v2))+12:]))
	v6 = t5
	store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v6))
	store32(m.memory[uint32(v0):], uint32(v5))
	if v4 == 0 {
		goto l0
	}
	v5 = v4 << 2
	if v5 == 0 {
		goto l1
	}
	memory_copy(m.memory, uint32(v6), uint32(v3), uint32(v5))
l1:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v4))
l0:
	t6 := int32(load32(m.memory[int64(uint32(v1))+12:]))
	store32(m.memory[int64(uint32(v0))+12:], uint32(t6))
	m.g0 = v2 + i32(16)
}
func (m *Module) fn936(v0 int32) int32 {
	var v1, v2 int32
	t0 := m.g0
	v1 = t0 - i32(16)
	m.g0 = v1
	m.fn891(v1+i32(8), v0)
	v0 = i32(0)
	{
		t1 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		v2 = t1
		if v2 == 0 {
			goto l0
		}
		t2 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		v2 = t2 + v2<<2 + i32(-4)
		if v2 == 0 {
			goto l0
		}
		t3 := int32(load32(m.memory[uint32(v2):]))
		v0 = t3
	}
l0:
	m.g0 = v1 + i32(16)
	return v0
}
func (m *Module) fn937(v0, v1 int32) {
	var v2, v3, v4 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	m.fn1685(v2+i32(4), v1, i32(4), i32(4))
	t1 := int32(load32(m.memory[int64(uint32(v2))+8:]))
	v3 = t1
	{
		t2 := int32(load32(m.memory[int64(uint32(v2))+4:]))
		if t2 != i32(1) {
			goto l0
		}
		t3 := int32(load32(m.memory[int64(uint32(v2))+12:]))
		m.fn2(v3, t3)
		panic("unreachable")
	}
l0:
	t4 := int32(load32(m.memory[int64(uint32(v2))+12:]))
	v4 = t4
	store32(m.memory[int64(uint32(v0))+8:], uint32(v1))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v4))
	store32(m.memory[uint32(v0):], uint32(v3))
	m.g0 = v2 + i32(16)
}
func (m *Module) fn938(v0, v1 int32) {
	var v2, v3, v4 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	t1 := int32(load32(m.memory[uint32(v1):]))
	v3 = t1
	t2 := m.fn936(v3 + i32(32))
	v4 = t2
	t3 := int32(load32(m.memory[uint32(v3+i32(52)):]))
	t4 := int32(load32(m.memory[uint32(v3+i32(56)):]))
	m.fn518(v2+i32(8), t3, t4, v4, i32(1148268))
	t5 := int32(load32(m.memory[int64(uint32(v2))+12:]))
	v4 = t5
	t6 := int32(load32(m.memory[int64(uint32(v3))+44:]))
	store32(m.memory[int64(uint32(v0))+16:], uint32(t6))
	store64(m.memory[int64(uint32(v0))+8:], uint64(i64(0)))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v4))
	store32(m.memory[uint32(v0):], uint32(v1))
	m.g0 = v2 + i32(16)
}
func (m *Module) fn939(v0, v1 int32) {
	var v2, v3, v4 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		t1 := v1
		v2 = t0
		if uint32(t1) <= uint32(v2) {
			goto l0
		}
		t2 := v0
		v3 = v1 - v2
		m.fn1233(t2, v3)
		v1 = v3 + i32(-1)
		t3 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		t4 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		v4 = t4
		v2 = t3 + v4<<2
	l2:
		store32(m.memory[uint32(v2):], uint32(i32(0)))
		if v1 == 0 {
			goto l1
		}
		v1 = v1 + i32(-1)
		v2 = v2 + i32(4)
		goto l2
	l1:
		v1 = v4 + v3
	}
l0:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v1))
}
func (m *Module) fn940(v0, v1 int32) {
	var v2 int32
	t0 := m.fn4(v1)
	v2 = t0
	store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
	store32(m.memory[uint32(v0):], uint32(v2))
}
func (m *Module) fn941(v0, v1 int32) {
	var v2, v3, v4 int32
	var v5 int64
	var v6 int32
	{
		{
			t0 := int32(load32(m.memory[int64(uint32(v1))+24:]))
			v2 = t0
			if v2 != 0 {
				goto l0
			}
			v1 = i32(0)
			goto l1
		}
	l0:
		t1 := int32(load32(m.memory[int64(uint32(v1))+16:]))
		v3 = t1
		t2 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		v4 = t2
		t3 := int64(load64(m.memory[uint32(v1):]))
		v5 = t3
	l3:
		{
			if v5 != i64(0) {
				goto l2
			}
			t4 := v1
			v3 = v3 + i32(-64)
			store32(m.memory[int64(uint32(t4))+16:], uint32(v3))
			t5 := v1
			v6 = v4 + i32(8)
			store32(m.memory[int64(uint32(t5))+8:], uint32(v6))
			t6 := int64(load64(m.memory[uint32(v4):]))
			t7 := v1
			v5 = (t6 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
			store64(m.memory[uint32(t7):], uint64(v5))
			v4 = v6
			goto l3
		}
	l2:
		store32(m.memory[int64(uint32(v1))+24:], uint32(v2+i32(-1)))
		store64(m.memory[uint32(v1):], uint64((v5+i64(-1))&v5))
		t8 := int32(load32(m.memory[uint32(v3-int32(int64(bits.TrailingZeros64(uint64(v5))))&i32(120)+i32(-4)):]))
		v4 = t8
		v1 = i32(1)
	}
l1:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v4))
	store32(m.memory[uint32(v0):], uint32(v1))
}
