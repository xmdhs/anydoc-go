package core

func (m *Module) fn852(v0, v1, v2 int32) {
	var v3 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		t1 := v1
		v3 = t0
		if uint32(t1) > uint32(v3) {
			return
		}
		{
			if v1 == 0 {
				goto l1
			}
			if uint32(v1) >= uint32(v3) {
				goto l1
			}
			t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t3 := int32(int8(m.memory[uint32(t2+v1)]))
			if t3 > i32(-65) {
				goto l1
			}
			m.fn256(i32(1087836), i32(48), v2)
			panic("unreachable")
		}
	l1:
		store32(m.memory[int64(uint32(v0))+8:], uint32(v1))
	}
}
func (m *Module) fn853(v0, v1 int32) int32 {
	var v2, v3 int32
	{
		v2 = v1 - v0
		if uint32(v2) < uint32(i32(16)) {
			v3 = i32(0)
			if v1 == v0 {
				goto l1
			}
		l2:
			{
				t1 := int32(int8(m.memory[uint32(v0)]))
				t2 := v3
				var p3 int32
				if t1 > i32(-65) {
					p3 = 1
				}
				v3 = t2 + p3
				v0 = v0 + i32(1)
				v2 = v2 + i32(-1)
				if v2 != 0 {
					goto l2
				}
			}
		l1:
			return v3
		}
		t0 := m.fn861(v0, v2)
		return t0
	}
}
func (m *Module) fn854(v0, v1 int32) {
	var v2, v3 int32
	t0 := m.g0
	v2 = t0 - i32(48)
	m.g0 = v2
	m.fn71(v2+i32(8), v1)
	{
		t1 := int32(load32(m.memory[int64(uint32(v2))+8:]))
		v1 = t1
		if v1 == 0 {
			goto l0
		}
		t2 := int32(load32(m.memory[int64(uint32(v2))+12:]))
		t3 := v2
		v3 = t2
		store32(m.memory[int64(uint32(t3))+36:], uint32(v3))
		store32(m.memory[int64(uint32(v2))+32:], uint32(v1))
		if v3 == 0 {
			goto l1
		}
		store32(m.memory[int64(uint32(v2))+44:], uint32(i32(1)))
		store32(m.memory[int64(uint32(v2))+40:], uint32(v2+i32(32)))
		m.fn73(v2+i32(20), i32(1051101), v2+i32(40))
		goto l2
	l1:
		m.fn51(v2+i32(20), i32(1282677), i32(1))
	l2:
		t4 := int32(load32(m.memory[int64(uint32(v2))+28:]))
		store32(m.memory[int64(uint32(v0))+8:], uint32(t4))
		t5 := int64(load64(m.memory[int64(uint32(v2))+20:]))
		store64(m.memory[uint32(v0):], uint64(t5))
		goto l3
	}
l0:
	store32(m.memory[uint32(v0):], uint32(i32(-1)))
l3:
	m.g0 = v2 + i32(48)
}
func (m *Module) fn855(v0, v1 int32) {
	t0 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	store32(m.memory[int64(uint32(v0))+4:], uint32(t0))
	t1 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	t2 := int32(load32(m.memory[uint32(v1):]))
	t4 := v0
	p3 := t1
	if t2 == i32(-1) {
		p3 = i32(0)
	}
	store32(m.memory[uint32(t4):], uint32(p3))
}
func (m *Module) fn856(v0, v1, v2, v3 int32) {
	var v4, v5, v6, v7, v8, v9 int32
l7:
	v4 = v2
	if v4 != 0 {
		goto l0
	}
	v4 = i32(0)
	goto l1
l0:
	{
		v5 = v1 + v4
		v2 = v5 + i32(-1)
		t0 := int32(int8(m.memory[uint32(v2)]))
		v6 = t0
		if v6 > i32(-1) {
			goto l2
		}
		{
			v2 = v5 + i32(-2)
			t1 := int32(m.memory[uint32(v2)])
			v7 = t1
			v8 = int32(int8(v7))
			if v8 < i32(-64) {
				goto l3
			}
			v5 = v7 & i32(31)
			goto l4
		}
	l3:
		{
			{
				v2 = v5 + i32(-3)
				t2 := int32(m.memory[uint32(v2)])
				v7 = t2
				v9 = int32(int8(v7))
				if v9 < i32(-64) {
					goto l5
				}
				v5 = v7 & i32(15)
				goto l6
			}
		l5:
			v2 = v5 + i32(-4)
			t3 := int32(m.memory[uint32(v2)])
			v5 = t3&i32(7)<<6 | v9&i32(63)
		}
	l6:
		v5 = v5<<6 | v8&i32(63)
	l4:
		v6 = v5<<6 | v6&i32(63)
	}
l2:
	v2 = v2 - v1
	if v6 == v3 {
		goto l7
	}
l1:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v4))
	store32(m.memory[uint32(v0):], uint32(v1))
}
func (m *Module) fn857(v0, v1, v2, v3 int32) int32 {
	if uint32(v2) < uint32(v1) {
		return v0 + v2*i32(12)
	}
	m.fn158(v2, v1, v3)
	panic("unreachable")
}
func (m *Module) fn858(v0, v1 int32) int32 {
	var v2, v3, v4 int32
	v2 = v1 << 4
	v3 = i32(0) - v2
	v0 = v0 + v2
l2:
	if v3 != 0 {
		v4 = v1
		{
			t0 := int32(load32(m.memory[uint32(v0+i32(-8)):]))
			if t0 != 0 {
				goto l1
			}
			v3 = v3 + i32(16)
			v1 = v1 + i32(-1)
			v2 = v0 + i32(-4)
			v0 = v0 + i32(-16)
			t1 := int32(m.memory[uint32(v2)])
			if t1&i32(1) == 0 {
				goto l2
			}
		}
	l1:
		return v4
	}
	return i32(0)
}
func (m *Module) fn859(v0 int32) {
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
		v1 = v1 + i32(-1)
		v3 = v3 + i32(16)
		goto l1
	}
l0:
	t4 := int32(load32(m.memory[uint32(v0):]))
	m.fn419(t4, v2)
}
func (m *Module) fn860(v0 int32) {
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
	m.fn859(v3)
	v3 = v3 + i32(12)
	goto l1
l0:
	t2 := int32(load32(m.memory[uint32(v0):]))
	m.fn136(t2, v2, i32(4), i32(12))
}
func (m *Module) fn861(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7, v8, v9 int32
	t0 := v1
	t1 := v0
	v2 = (v0 + i32(3)) & i32(-4)
	v3 = t1 - v2
	v4 = t0 + v3
	v5 = v4 & i32(3)
	v1 = i32(0)
	v6 = i32(0)
	if v0 == v2 {
		goto l0
	}
	v6 = i32(0)
l1:
	{
		t2 := int32(int8(m.memory[uint32(v0)]))
		t3 := v6
		var p4 int32
		if t2 > i32(-65) {
			p4 = 1
		}
		v6 = t3 + p4
		v0 = v0 + i32(1)
		v3 = v3 + i32(1)
		if v3 != 0 {
			goto l1
		}
	}
l0:
	if v5 == 0 {
		goto l2
	}
	v0 = v2 + v4&i32(0x7ffffffc)
	v1 = i32(0)
l3:
	{
		t5 := int32(int8(m.memory[uint32(v0)]))
		t6 := v1
		var p7 int32
		if t5 > i32(-65) {
			p7 = 1
		}
		v1 = t6 + p7
		v0 = v0 + i32(1)
		v5 = v5 + i32(-1)
		if v5 != 0 {
			goto l3
		}
	}
l2:
	v3 = int32(uint32(v4) >> 2)
	v7 = v1 + v6
l8:
	{
		v6 = v2
		if v3 == 0 {
			goto l4
		}
		p8 := i32(192)
		if uint32(v3) < uint32(i32(192)) {
			p8 = v3
		}
		v4 = p8
		v8 = v4 & i32(3)
		v9 = v4 << 2
		v0 = v9 & i32(1008)
		if v0 != 0 {
			goto l5
		}
		v1 = i32(0)
		goto l6
	l5:
		v2 = v6 + v0
		v1 = i32(0)
		v0 = v6
	l7:
		{
			t9 := int32(load32(m.memory[uint32(v0+i32(12)):]))
			v5 = t9
			t10 := int32(load32(m.memory[uint32(v0+i32(8)):]))
			t11 := (int32(uint32(v5^i32(-1))>>7) | int32(uint32(v5)>>6)) & i32(16843009)
			v5 = t10
			t12 := int32(load32(m.memory[uint32(v0+i32(4)):]))
			t13 := (int32(uint32(v5^i32(-1))>>7) | int32(uint32(v5)>>6)) & i32(16843009)
			v5 = t12
			t14 := int32(load32(m.memory[uint32(v0):]))
			t15 := (int32(uint32(v5^i32(-1))>>7) | int32(uint32(v5)>>6)) & i32(16843009)
			v5 = t14
			v1 = t11 + (t13 + (t15 + ((int32(uint32(v5^i32(-1))>>7)|int32(uint32(v5)>>6))&i32(16843009) + v1)))
			v0 = v0 + i32(16)
			if v0 != v2 {
				goto l7
			}
		}
	l6:
		v3 = v3 - v4
		v2 = v6 + v9
		v7 = int32(uint32((int32(uint32(v1)>>8)&i32(0xff00ff)+v1&i32(0xff00ff))*i32(65537))>>16) + v7
		if v8 == 0 {
			goto l8
		}
	}
	v5 = v8 << 2
	v0 = v6 + v4&i32(252)<<2
	v1 = i32(0)
l9:
	{
		t16 := int32(load32(m.memory[uint32(v0):]))
		v2 = t16
		v1 = (int32(uint32(v2^i32(-1))>>7)|int32(uint32(v2)>>6))&i32(16843009) + v1
		v0 = v0 + i32(4)
		v5 = v5 + i32(-4)
		if v5 != 0 {
			goto l9
		}
	}
	v7 = int32(uint32((int32(uint32(v1)>>8)&i32(0xff00ff)+v1&i32(0xff00ff))*i32(65537))>>16) + v7
l4:
	return v7
}
func (m *Module) fn862(v0, v1 int32) {
	var v2 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	{
		t1 := m.fn863(v1)
		v1 = t1
		if v1 == 0 {
			goto l0
		}
		t2 := int32(load32(m.memory[uint32(v1+i32(28)):]))
		t3 := int32(load32(m.memory[uint32(v1+i32(32)):]))
		m.fn864(v2+i32(4), t2, t3)
		t4 := int32(load32(m.memory[int64(uint32(v2))+8:]))
		t5 := v0
		v1 = t4
		t6 := int32(load32(m.memory[int64(uint32(v2))+12:]))
		m.fn865(t5, v1, t6)
		t7 := int32(load32(m.memory[int64(uint32(v2))+4:]))
		m.fn16(t7, v1)
		goto l1
	}
l0:
	store32(m.memory[uint32(v0):], uint32(i32(-1)))
l1:
	m.g0 = v2 + i32(16)
}
func (m *Module) fn863(v0 int32) int32 {
	var v1, v2 int32
	v1 = v0 + i32(12)
l1:
	{
		{
			t0 := m.fn866(v0)
			v2 = t0
			if v2 != 0 {
				goto l0
			}
			return i32(0)
		}
	l0:
		t1 := int32(load32(m.memory[uint32(v2):]))
		if t1 == i32(-1) {
			goto l1
		}
		t2 := m.fn867(v1, v2)
		if t2 == 0 {
			goto l1
		}
	}
	return v2
}
func (m *Module) fn864(v0, v1, v2 int32) {
	var v3 int32
	t0 := m.g0
	v3 = t0 - i32(32)
	m.g0 = v3
	store32(m.memory[int64(uint32(v3))+16:], uint32(i32(0)))
	store64(m.memory[int64(uint32(v3))+8:], uint64(i64(0x100000000)))
	m.fn868(v3+i32(20), v1, v2)
l1:
	{
		t1 := m.fn866(v3 + i32(20))
		v2 = t1
		if v2 == 0 {
			goto l0
		}
		t2 := int32(load32(m.memory[uint32(v2):]))
		if t2 != i32(-1) {
			goto l1
		}
		t3 := int32(load32(m.memory[int64(uint32(v2))+8:]))
		t4 := int32(load32(m.memory[int64(uint32(v2))+12:]))
		m.fn75(v3+i32(8), t3, t4)
		goto l1
	}
l0:
	t5 := int32(load32(m.memory[int64(uint32(v3))+20:]))
	t6 := int32(load32(m.memory[int64(uint32(v3))+24:]))
	m.fn44(t5, t6)
	t7 := int32(load32(m.memory[int64(uint32(v3))+16:]))
	store32(m.memory[int64(uint32(v0))+8:], uint32(t7))
	t8 := int64(load64(m.memory[int64(uint32(v3))+8:]))
	store64(m.memory[uint32(v0):], uint64(t8))
	m.g0 = v3 + i32(32)
}
func (m *Module) fn865(v0, v1, v2 int32) {
	var v3 int32
	t0 := m.g0
	v3 = t0 - i32(32)
	m.g0 = v3
	m.fn140(v3+i32(8), v2)
	store32(m.memory[int64(uint32(v3))+20:], uint32(i32(-2)))
	store32(m.memory[int64(uint32(v3))+24:], uint32(v1))
	store32(m.memory[int64(uint32(v3))+28:], uint32(v1+v2))
l5:
	{
		t1 := m.fn869(v3 + i32(20))
		v2 = t1
		switch v2 + i32(-9) {
		case 0:
			goto l0
		case 1:
			goto l1
		case 4:
			v2 = i32(32)
			t2 := m.fn870(v3 + i32(20))
			v1 = t2
			if v1 == 0 {
				goto l0
			}
			t3 := int32(load32(m.memory[uint32(v1):]))
			if t3 != i32(10) {
				goto l0
			}
			_ = m.fn869(v3 + i32(20))
			goto l0
		default:
			if v2 == i32(160) {
				goto l1
			}
			if v2 == i32(173) {
				goto l5
			}
			if v2 == i32(8203) {
				goto l5
			}
			if v2 == i32(65279) {
				goto l5
			}
			if v2 == i32(-1) {
				t5 := int32(load32(m.memory[int64(uint32(v3))+16:]))
				store32(m.memory[int64(uint32(v0))+8:], uint32(t5))
				t6 := int64(load64(m.memory[int64(uint32(v3))+8:]))
				store64(m.memory[uint32(v0):], uint64(t6))
				m.g0 = v3 + i32(32)
				return
			}
			fallthrough
		case 2, 3:
			v1 = v2 + i32(-127)
			if uint32(v2) < uint32(i32(32)) {
				goto l5
			}
			if uint32(v1) < uint32(i32(33)) {
				goto l5
			}
			goto l0
		}
	}
l1:
	v2 = i32(32)
	goto l0
l0:
	m.fn74(v3+i32(8), v2)
	goto l5
}
func (m *Module) fn866(v0 int32) int32 {
	var v1, v2, v3 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		v1 = t0
		if v1 != 0 {
			t1 := v0
			v1 = v1 + i32(-1)
			store32(m.memory[int64(uint32(t1))+8:], uint32(v1))
			{
				t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
				t3 := int32(load32(m.memory[uint32(t2+v1<<2):]))
				v2 = t3
				t4 := int32(load32(m.memory[uint32(v2):]))
				if t4 == i32(-1) {
					goto l1
				}
				t5 := int32(load32(m.memory[int64(uint32(v2))+28:]))
				t6 := v0
				v3 = t5
				t7 := int32(load32(m.memory[int64(uint32(v2))+32:]))
				m.fn871(t6, v3, v3+t7*i32(44))
				t8 := int32(load32(m.memory[int64(uint32(v0))+8:]))
				v3 = t8
				if uint32(v3) < uint32(v1) {
					m.fn151(v1, v3, v3, i32(1087352))
					panic("unreachable")
				}
				t9 := int32(load32(m.memory[int64(uint32(v0))+4:]))
				m.fn872(t9+v1<<2, v3-v1)
			}
		l1:
			return v2
		}
		return i32(0)
	}
}
func (m *Module) fn867(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t2 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	t3 := int32(load32(m.memory[int64(uint32(v0))+12:]))
	t4 := m.fn847(v1, t0, t1, t2, t3)
	return t4
}
func (m *Module) fn868(v0, v1, v2 int32) {
	var v3, v4 int32
	t0 := m.g0
	v3 = t0 - i32(32)
	m.g0 = v3
	m.fn59(v3+i32(8), v2, i32(4), i32(4))
	store32(m.memory[int64(uint32(v3))+28:], uint32(i32(0)))
	t1 := int64(load64(m.memory[int64(uint32(v3))+8:]))
	store64(m.memory[int64(uint32(v3))+20:], uint64(t1))
	m.fn871(v3+i32(20), v1, v1+v2*i32(44))
	t2 := int32(load32(m.memory[int64(uint32(v3))+20:]))
	v2 = t2
	t3 := int32(load32(m.memory[int64(uint32(v3))+24:]))
	v1 = t3
	t4 := int32(load32(m.memory[int64(uint32(v3))+28:]))
	t5 := v1
	v4 = t4
	m.fn872(t5, v4)
	store32(m.memory[int64(uint32(v0))+8:], uint32(v4))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
	store32(m.memory[uint32(v0):], uint32(v2))
	m.g0 = v3 + i32(32)
}
func (m *Module) fn869(v0 int32) int32 {
	var v1 int32
	t0 := int32(load32(m.memory[uint32(v0):]))
	v1 = t0
	store32(m.memory[uint32(v0):], uint32(i32(-2)))
	{
		if v1 != i32(-2) {
			goto l0
		}
		t1 := m.fn48(v0 + i32(4))
		v1 = t1
	}
l0:
	return v1
}
func (m *Module) fn870(v0 int32) int32 {
	var v1 int32
	{
		t0 := int32(load32(m.memory[uint32(v0):]))
		v1 = t0
		if v1 != i32(-2) {
			goto l0
		}
		t1 := m.fn48(v0 + i32(4))
		t2 := v0
		v1 = t1
		store32(m.memory[uint32(t2):], uint32(v1))
	}
l0:
	p3 := v0
	if v1 == i32(-1) {
		p3 = i32(0)
	}
	return p3
}
func (m *Module) fn871(v0, v1, v2 int32) {
	var v3, v4 int32
	{
		t0 := int32(uint32(v2-v1) / uint32(i32(44)))
		v3 = t0
		t1 := int32(load32(m.memory[uint32(v0):]))
		t2 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		t3 := v3
		v4 = t2
		if uint32(t3) <= uint32(t1-v4) {
			goto l0
		}
		m.fn62(v0, v4, v3, i32(4), i32(4))
		t4 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		v4 = t4
	}
l0:
	v3 = v4 + v3
	t5 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	v4 = t5 + v4<<2
l2:
	if v1 == v2 {
		goto l1
	}
	store32(m.memory[uint32(v4):], uint32(v1))
	v4 = v4 + i32(4)
	v1 = v1 + i32(44)
	goto l2
l1:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v3))
}
func (m *Module) fn872(v0, v1 int32) {
	var v2, v3, v4, v5, v6, v7, v8 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	t1 := v2
	t2 := v0
	v3 = int32(uint32(v1) >> 1)
	m.fn873(t1, t2, v3, v3, i32(1301108))
	t3 := int32(load32(m.memory[int64(uint32(v2))+4:]))
	v4 = t3
	t4 := int32(load32(m.memory[uint32(v2):]))
	v5 = t4
	t5 := v2
	t6 := v0 + v1<<2
	v1 = v3 << 2
	m.fn873(t5, t6-v1, v3, v3, i32(1301124))
	t7 := int32(load32(m.memory[uint32(v2):]))
	v0 = v1 + t7 + i32(-4)
	v1 = i32(0)
	t8 := int32(load32(m.memory[int64(uint32(v2))+4:]))
	t9 := v3 + i32(-1)
	v6 = t8
	var p10 int32
	if uint32(t9) < uint32(v6) {
		p10 = 1
	}
	v7 = p10
l3:
	v8 = v3 + v1
	if v8 == 0 {
		m.g0 = v2 + i32(16)
		return
	}
	if v4+v1 == 0 {
		m.fn158(v4, v4, i32(1301140))
		panic("unreachable")
	}
	{
		if v7 == 0 {
			m.fn158(v8+i32(-1), v6, i32(1301156))
			panic("unreachable")
		}
		t11 := int32(load32(m.memory[uint32(v5):]))
		v8 = t11
		t12 := int32(load32(m.memory[uint32(v0):]))
		store32(m.memory[uint32(v5):], uint32(t12))
		store32(m.memory[uint32(v0):], uint32(v8))
		v5 = v5 + i32(4)
		v0 = v0 + i32(-4)
		v1 = v1 + i32(-1)
		goto l3
	}
}
func (m *Module) fn873(v0, v1, v2, v3, v4 int32) {
	if uint32(v2) >= uint32(v3) {
		goto l0
	}
	m.fn91(i32(1301172), i32(19), v4)
	panic("unreachable")
l0:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
	store32(m.memory[uint32(v0):], uint32(v1))
	store32(m.memory[int64(uint32(v0))+12:], uint32(v2-v3))
	store32(m.memory[int64(uint32(v0))+8:], uint32(v1+v3<<2))
}
func (m *Module) fn874(v0, v1 int32) {
	var v2, v3, v4, v5 int32
	t0 := m.g0
	v2 = t0 - i32(96)
	m.g0 = v2
	m.fn875(v2+i32(8), v1+i32(4))
	{
		t1 := int32(load32(m.memory[int64(uint32(v2))+8:]))
		v3 = t1
		if v3 == 0 {
			goto l0
		}
		t2 := int32(load32(m.memory[int64(uint32(v2))+12:]))
		t3 := v2
		v4 = t2
		store32(m.memory[int64(uint32(t3))+36:], uint32(v4))
		store32(m.memory[int64(uint32(v2))+32:], uint32(v3))
		store32(m.memory[int64(uint32(v2))+56:], uint32(i32(1)))
		store32(m.memory[int64(uint32(v2))+52:], uint32(v2+i32(32)))
		m.fn73(v2+i32(40), i32(1052712), v2+i32(52))
		t4 := int32(load32(m.memory[uint32(v1):]))
		t5 := int32(load32(m.memory[int64(uint32(v2))+44:]))
		t6 := v2 + i32(52)
		t7 := t4 + i32(72)
		v1 = t5
		t8 := int32(load32(m.memory[int64(uint32(v2))+48:]))
		m.fn876(t6, t7, v1, t8)
		{
			{
				t9 := int32(load32(m.memory[int64(uint32(v2))+52:]))
				v5 = t9
				if v5 != i32(-0x7fffffff) {
					goto l1
				}
				t10 := int32(load32(m.memory[int64(uint32(v2))+56:]))
				t11 := v2 + i32(84)
				v3 = t10
				t12 := int32(load32(m.memory[int64(uint32(v2))+60:]))
				m.fn92(t11, v3, t12)
				m.fn490(v2+i32(72), v2+i32(84))
				t13 := int32(load32(m.memory[int64(uint32(v2))+76:]))
				t14 := int32(load32(m.memory[int64(uint32(v2))+80:]))
				m.fn877(v2+i32(84), t13, t14)
				m.fn878(v2+i32(20), v2+i32(84), v2+i32(72))
				goto l2
			}
		l1:
			m.fn51(v2+i32(20), v3, v4)
			t15 := int32(load32(m.memory[int64(uint32(v2))+56:]))
			v3 = t15
		}
	l2:
		m.fn879(v5, v3)
		t16 := int32(load32(m.memory[int64(uint32(v2))+40:]))
		m.fn16(t16, v1)
		t17 := int32(load32(m.memory[int64(uint32(v2))+28:]))
		store32(m.memory[int64(uint32(v0))+8:], uint32(t17))
		t18 := int64(load64(m.memory[int64(uint32(v2))+20:]))
		store64(m.memory[uint32(v0):], uint64(t18))
		goto l3
	}
l0:
	store32(m.memory[uint32(v0):], uint32(i32(-1)))
l3:
	m.g0 = v2 + i32(96)
}
func (m *Module) fn875(v0, v1 int32) {
	var v2, v3, v4, v5, v6, v7, v8, v9 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	v3 = v1 + i32(16)
l6:
	{
		v4 = i32(0)
		t1 := int32(m.memory[int64(uint32(v1))+29])
		if t1 != 0 {
			goto l5
		}
		t2 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		v5 = t2
		{
			{
			l2:
				{
					t3 := int32(load32(m.memory[int64(uint32(v1))+16:]))
					v6 = t3
					t4 := int32(load32(m.memory[int64(uint32(v1))+20:]))
					v7 = t4
					m.fn572(v2+i32(8), v3)
					t5 := int32(load32(m.memory[int64(uint32(v2))+12:]))
					v8 = t5
					if v8 == i32(-1) {
						goto l1
					}
					t6 := int32(load32(m.memory[int64(uint32(v2))+8:]))
					v9 = t6
					t7 := m.fn630(v8)
					if t7 == 0 {
						goto l2
					}
				}
				t8 := int32(load32(m.memory[uint32(v1):]))
				v8 = t8
				t9 := int32(load32(m.memory[int64(uint32(v1))+16:]))
				t10 := int32(load32(m.memory[int64(uint32(v1))+20:]))
				store32(m.memory[uint32(v1):], uint32(v7-v6+v9+t9-t10))
				v4 = v5 + v8
				v8 = v9 - v8
				goto l3
			}
		l1:
			t11 := int32(m.memory[int64(uint32(v1))+29])
			if t11 != 0 {
				goto l5
			}
			m.memory[int64(uint32(v1))+29] = byte(i32(1))
			t12 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			v7 = t12
			t13 := int32(load32(m.memory[uint32(v1):]))
			v6 = t13
			{
				t14 := int32(m.memory[int64(uint32(v1))+28])
				if t14 != 0 {
					goto l4
				}
				if v7 == v6 {
					goto l5
				}
			}
		l4:
			v8 = v7 - v6
			t15 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			v4 = t15 + v6
		}
	l3:
		if v8 == 0 {
			goto l6
		}
		goto l5
	}
l5:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v8))
	store32(m.memory[uint32(v0):], uint32(v4))
	m.g0 = v2 + i32(16)
}
func (m *Module) fn876(v0, v1, v2, v3 int32) {
	m.fn880(v0, v1, v2, v3, i32(1))
}
func (m *Module) fn877(v0, v1, v2 int32) {
	var v3 int32
	var v4 int64
	t0 := m.g0
	v3 = t0 - i32(48)
	m.g0 = v3
	m.fn565(v3+i32(8), v1, v2, i32(1084512), i32(27))
	{
		{
			t1 := int32(load32(m.memory[int64(uint32(v3))+8:]))
			v2 = t1
			if v2 != 0 {
				goto l0
			}
			store32(m.memory[uint32(v0):], uint32(i32(-1)))
			goto l1
		}
	l0:
		t2 := int32(load32(m.memory[int64(uint32(v3))+12:]))
		m.fn775(v3+i32(32), v2, t2, i32(47))
		{
			t3 := int32(load32(m.memory[int64(uint32(v3))+32:]))
			v2 = t3
			if v2 != 0 {
				goto l2
			}
			store32(m.memory[uint32(v0):], uint32(i32(-1)))
			goto l1
		}
	l2:
		t4 := int64(load64(m.memory[int64(uint32(v3))+40:]))
		v4 = t4
		t5 := int32(load32(m.memory[int64(uint32(v3))+36:]))
		store32(m.memory[int64(uint32(v3))+20:], uint32(t5))
		store32(m.memory[int64(uint32(v3))+16:], uint32(v2))
		store64(m.memory[int64(uint32(v3))+24:], uint64(v4))
		store32(m.memory[int64(uint32(v3))+44:], uint32(i32(1)))
		store32(m.memory[int64(uint32(v3))+36:], uint32(i32(1)))
		store32(m.memory[int64(uint32(v3))+40:], uint32(v3+i32(24)))
		store32(m.memory[int64(uint32(v3))+32:], uint32(v3+i32(16)))
		m.fn73(v0, i32(0x10006d), v3+i32(32))
	}
l1:
	m.g0 = v3 + i32(48)
}
func (m *Module) fn878(v0, v1, v2 int32) {
	{
		t0 := int32(load32(m.memory[uint32(v1):]))
		if t0 == i32(-1) {
			goto l0
		}
		t1 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		store32(m.memory[int64(uint32(v0))+8:], uint32(t1))
		t2 := int64(load64(m.memory[uint32(v1):]))
		store64(m.memory[uint32(v0):], uint64(t2))
		t3 := int32(load32(m.memory[uint32(v2):]))
		t4 := int32(load32(m.memory[int64(uint32(v2))+4:]))
		m.fn16(t3, t4)
		return
	}
l0:
	t5 := int32(load32(m.memory[int64(uint32(v2))+8:]))
	store32(m.memory[int64(uint32(v0))+8:], uint32(t5))
	t6 := int64(load64(m.memory[uint32(v2):]))
	store64(m.memory[uint32(v0):], uint64(t6))
}
func (m *Module) fn879(v0, v1 int32) {
	if v0 < i32(-0x7ffffffe) {
		return
	}
	m.fn16(v0, v1)
}
func (m *Module) fn880(v0, v1, v2, v3, v4 int32) {
	var v5, v6, v7, v8, v9, v10, v11, v12 int32
	t0 := m.g0
	v5 = t0 - i32(32)
	m.g0 = v5
	m.fn881(v5+i32(8), i32(58), v2, v3)
	{
		{
			{
				t1 := int32(load32(m.memory[int64(uint32(v5))+8:]))
				if t1 == i32(1) {
					goto l0
				}
				v6 = i32(0)
				goto l1
			}
		l0:
			t2 := int32(load32(m.memory[int64(uint32(v5))+12:]))
			t3 := v5
			v7 = t2
			m.fn148(t3, v7+i32(1), v2, v3, i32(1282368))
			if uint32(v7) > uint32(v3) {
				m.fn151(i32(0), v7, v3, i32(1282384))
				panic("unreachable")
			}
			t4 := int32(load32(m.memory[int64(uint32(v5))+4:]))
			v3 = t4
			v6 = v2
			t5 := int32(load32(m.memory[uint32(v5):]))
			v2 = t5
		}
	l1:
		t6 := int32(load32(m.memory[int64(uint32(v1))+20:]))
		v8 = t6
		v9 = v8 << 4
		t7 := int32(load32(m.memory[int64(uint32(v1))+16:]))
		v10 = t7
		{
			{
				if v6 == 0 {
					if v4 == 0 {
						goto l8
					}
					v4 = v10 + v9
					v9 = v8 << 4
					{
					l10:
						{
							if v9 == 0 {
								store32(m.memory[uint32(v0):], uint32(i32(-0x80000000)))
								goto l11
							}
							v9 = v9 + i32(-16)
							v8 = v4 + i32(-12)
							v4 = v4 + i32(-16)
							t20 := int32(load32(m.memory[uint32(v8):]))
							if t20 != 0 {
								goto l10
							}
						}
						t21 := int32(load32(m.memory[int64(uint32(v1))+4:]))
						t22 := int32(load32(m.memory[int64(uint32(v1))+8:]))
						m.fn883(v0, v10+v9, t21, t22)
						goto l11
					}
				}
				t8 := int32(load32(m.memory[int64(uint32(v1))+8:]))
				v11 = t8
				t9 := int32(load32(m.memory[int64(uint32(v1))+4:]))
				v1 = t9
			l7:
				if v9 == 0 {
					goto l4
				}
				{
					v4 = v10 + v9
					t10 := int32(load32(m.memory[uint32(v4+i32(-12)):]))
					v8 = t10
					if v8 == 0 {
						goto l5
					}
					t11 := v5 + i32(16)
					t12 := v1
					t13 := v11
					v12 = v4 + i32(-16)
					t14 := int32(load32(m.memory[uint32(v12):]))
					m.fn309(t11, t12, t13, t14, i32(1282512))
					t15 := int32(load32(m.memory[int64(uint32(v5))+24:]))
					t16 := int32(load32(m.memory[int64(uint32(v5))+28:]))
					m.fn309(v5+i32(16), t15, t16, v8, i32(1282528))
					t17 := int32(load32(m.memory[int64(uint32(v5))+16:]))
					t18 := int32(load32(m.memory[int64(uint32(v5))+20:]))
					t19 := m.fn882(t17, t18, v6, v7)
					if t19 != 0 {
						goto l6
					}
				}
			l5:
				v9 = v9 + i32(-16)
				goto l7
			}
		l6:
			t23 := int32(load32(m.memory[uint32(v4+i32(-8)):]))
			if t23 != 0 {
				m.fn883(v0, v12, v1, v11)
				goto l11
			}
		}
	l4:
		m.fn884(v0, v6, v7)
		goto l11
	}
l8:
	store32(m.memory[uint32(v0):], uint32(i32(-0x80000000)))
l11:
	store32(m.memory[int64(uint32(v0))+16:], uint32(v3))
	store32(m.memory[int64(uint32(v0))+12:], uint32(v2))
	m.g0 = v5 + i32(32)
}
func (m *Module) fn881(v0, v1, v2, v3 int32) {
	var v4 int32
	t0 := m.g0
	v4 = t0 - i32(16)
	m.g0 = v4
	m.fn294(v4+i32(8), v1, v2, v2+v3)
	v3 = i32(1)
	{
		{
			t1 := int32(load32(m.memory[int64(uint32(v4))+8:]))
			if t1 == i32(1) {
				goto l0
			}
			v3 = i32(0)
			goto l1
		}
	l0:
		t2 := int32(load32(m.memory[int64(uint32(v4))+12:]))
		v2 = t2 - v2
	}
l1:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v2))
	store32(m.memory[uint32(v0):], uint32(v3))
	m.g0 = v4 + i32(16)
}
func (m *Module) fn882(v0, v1, v2, v3 int32) int32 {
	var v4 int32
	v4 = i32(0)
	{
		if v1 != v3 {
			goto l0
		}
		t0 := m.fn1755(v0, v2, v1)
		v4 = t0
	}
l0:
	return v4
}
func (m *Module) fn883(v0, v1, v2, v3 int32) {
	var v4, v5 int32
	t0 := m.g0
	v4 = t0 - i32(16)
	m.g0 = v4
	{
		{
			t1 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			v5 = t1
			if v5 != 0 {
				goto l0
			}
			v1 = i32(-0x80000000)
			goto l1
		}
	l0:
		t2 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t3 := int32(load32(m.memory[uint32(v1):]))
		m.fn309(v4, v2, v3, t2+t3, i32(1282544))
		t4 := int32(load32(m.memory[int64(uint32(v4))+8:]))
		t5 := int32(load32(m.memory[int64(uint32(v4))+12:]))
		m.fn309(v4, t4, t5, v5, i32(1282560))
		t6 := int64(load64(m.memory[uint32(v4):]))
		store64(m.memory[int64(uint32(v0))+4:], uint64(t6))
		v1 = i32(-0x7fffffff)
	}
l1:
	store32(m.memory[uint32(v0):], uint32(v1))
	m.g0 = v4 + i32(16)
}
func (m *Module) fn884(v0, v1, v2 int32) {
	var v3, v4, v5 int32
	t0 := m.g0
	v3 = t0 - i32(16)
	m.g0 = v3
	m.fn1064(v3+i32(8), v2)
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
func (m *Module) fn885(v0, v1 int32) {
	var v2, v3, v4, v5 int32
	t0 := m.g0
	v2 = t0 - i32(32)
	m.g0 = v2
	v3 = v1 + i32(12)
	v4 = i32(-1)
	{
	l1:
		{
			t1 := m.fn866(v1)
			v5 = t1
			if v5 == 0 {
				goto l0
			}
			t2 := int32(load32(m.memory[uint32(v5):]))
			if t2 == i32(-1) {
				goto l1
			}
			t3 := m.fn867(v3, v5)
			if t3 == 0 {
				goto l1
			}
			t4 := int32(load32(m.memory[uint32(v5+i32(28)):]))
			t5 := int32(load32(m.memory[uint32(v5+i32(32)):]))
			t6 := m.fn886(t4, t5, i32(1073986), i32(56), i32(1072196), i32(1))
			v5 = t6
			if v5 == 0 {
				goto l1
			}
			t7 := int32(load32(m.memory[uint32(v5+i32(28)):]))
			t8 := int32(load32(m.memory[uint32(v5+i32(32)):]))
			m.fn864(v2+i32(20), t7, t8)
			t9 := int32(load32(m.memory[int64(uint32(v2))+24:]))
			t10 := v2 + i32(8)
			v5 = t9
			t11 := int32(load32(m.memory[int64(uint32(v2))+28:]))
			m.fn865(t10, v5, t11)
			t12 := int32(load32(m.memory[int64(uint32(v2))+20:]))
			m.fn16(t12, v5)
			t13 := int32(load32(m.memory[int64(uint32(v2))+12:]))
			t14 := v2
			v5 = t13
			t15 := int32(load32(m.memory[int64(uint32(v2))+16:]))
			m.fn46(t14, v5, t15)
			{
				t16 := int32(load32(m.memory[int64(uint32(v2))+4:]))
				if t16 != 0 {
					goto l2
				}
				t17 := int32(load32(m.memory[int64(uint32(v2))+8:]))
				m.fn16(t17, v5)
				goto l1
			}
		l2:
		}
		t18 := m.fn113(i32(8), i32(32))
		v5 = t18
		t19 := m.fn113(i32(4), i32(28))
		v1 = t19
		store32(m.memory[uint32(v1):], uint32(i32(3)))
		store32(m.memory[int64(uint32(v1))+16:], uint32(i32(0)))
		t20 := int64(load64(m.memory[int64(uint32(v2))+8:]))
		store64(m.memory[int64(uint32(v1))+4:], uint64(t20))
		t21 := int32(load32(m.memory[int64(uint32(v2))+16:]))
		store32(m.memory[int64(uint32(v1))+12:], uint32(t21))
		store64(m.memory[uint32(v5):], uint64(i64(0x180000000)))
		v4 = i32(1)
		store32(m.memory[int64(uint32(v5))+12:], uint32(i32(1)))
		store32(m.memory[int64(uint32(v5))+8:], uint32(v1))
		m.memory[int64(uint32(v0))+24] = byte(i32(2))
		store64(m.memory[int64(uint32(v0))+8:], uint64(i64(-0xffffffff)))
		store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
	}
l0:
	store32(m.memory[uint32(v0):], uint32(v4))
	m.g0 = v2 + i32(32)
}
func (m *Module) fn886(v0, v1, v2, v3, v4, v5 int32) int32 {
	v1 = v1 * i32(44)
l3:
	if v1 != 0 {
		{
			t0 := int32(load32(m.memory[uint32(v0):]))
			if t0 == i32(-1) {
				goto l2
			}
			t1 := m.fn847(v0, v2, v3, v4, v5)
			if t1 != 0 {
				goto l1
			}
		}
	l2:
		v0 = v0 + i32(44)
		v1 = v1 + i32(-44)
		goto l3
	}
	v0 = i32(0)
	goto l1
l1:
	return v0
}
func (m *Module) fn887(v0, v1 int32) {
	var v2 int32
	{
		t0 := int32(load32(m.memory[uint32(v0):]))
		t1 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		t2 := v1
		v2 = t1
		if uint32(t2) <= uint32(t0-v2) {
			return
		}
		m.fn62(v0, v2, v1, i32(4), i32(20))
	}
}
func (m *Module) fn888(v0, v1 int32) {
	var v2 int32
	t0 := m.fn113(i32(8), i32(32))
	v2 = t0
	store32(m.memory[uint32(v2):], uint32(i32(-0x80000000)))
	store32(m.memory[int64(uint32(v0))+16:], uint32(i32(1)))
	store64(m.memory[int64(uint32(v0))+8:], uint64(i64(0x100000001)))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v2))
	store32(m.memory[uint32(v0):], uint32(i32(1)))
	t1 := int64(load64(m.memory[uint32(v1):]))
	store64(m.memory[int64(uint32(v2))+4:], uint64(t1))
	t2 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	store32(m.memory[int64(uint32(v2))+12:], uint32(t2))
}
func (m *Module) fn889(v0, v1 int32) {
	var v2, v3, v4 int32
	t0 := m.g0
	v2 = t0 - i32(32)
	m.g0 = v2
	m.fn890(v2, v1)
	{
		t1 := int32(load32(m.memory[uint32(v2):]))
		v3 = t1
		if v3 == 0 {
			goto l0
		}
		t2 := int32(load32(m.memory[int64(uint32(v2))+4:]))
		v4 = t2
		t3 := m.fn113(i32(4), i32(28))
		v1 = t3
		m.fn865(v2+i32(20), v3, v4)
		store32(m.memory[uint32(v1):], uint32(i32(3)))
		store32(m.memory[int64(uint32(v1))+16:], uint32(i32(0)))
		t4 := int64(load64(m.memory[int64(uint32(v2))+20:]))
		store64(m.memory[int64(uint32(v1))+4:], uint64(t4))
		t5 := int32(load32(m.memory[int64(uint32(v2))+28:]))
		store32(m.memory[int64(uint32(v1))+12:], uint32(t5))
		store32(m.memory[int64(uint32(v2))+16:], uint32(i32(1)))
		store32(m.memory[int64(uint32(v2))+12:], uint32(v1))
		store32(m.memory[int64(uint32(v2))+8:], uint32(i32(1)))
		m.fn888(v0, v2+i32(8))
		goto l1
	}
l0:
	store32(m.memory[uint32(v0):], uint32(i32(-1)))
l1:
	m.g0 = v2 + i32(32)
}
func (m *Module) fn890(v0, v1 int32) {
	var v2, v3, v4, v5, v6 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	{
		{
			t1 := int32(load32(m.memory[int64(uint32(v1))+12:]))
			v3 = t1
			t2 := int32(load32(m.memory[int64(uint32(v1))+16:]))
			if v3 != t2 {
				goto l0
			}
			v1 = i32(0)
			goto l1
		}
	l0:
		t3 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		v4 = t3
		t4 := int32(load32(m.memory[uint32(v1):]))
		t5 := v2 + i32(8)
		v5 = t4
		t6 := int32(load32(m.memory[uint32(v5):]))
		m.fn891(t5, t6+i32(32))
		t7 := int32(load32(m.memory[int64(uint32(v2))+12:]))
		t8 := v3
		v6 = t7
		if uint32(t8) >= uint32(v6) {
			m.fn158(v3, v6, i32(1148572))
			panic("unreachable")
		}
		t9 := int32(load32(m.memory[int64(uint32(v2))+8:]))
		v6 = t9
		store32(m.memory[int64(uint32(v1))+12:], uint32(v3+i32(1)))
		t10 := int32(load32(m.memory[uint32(v6+v3<<2):]))
		t11 := v1
		v3 = t10
		store32(m.memory[int64(uint32(t11))+8:], uint32(v3))
		t12 := int32(load32(m.memory[uint32(v5):]))
		t13 := v2
		v1 = t12
		t14 := int32(load32(m.memory[uint32(v1+i32(52)):]))
		t15 := int32(load32(m.memory[uint32(v1+i32(56)):]))
		m.fn483(t13, t14, t15, v4, v3, i32(1148588))
		t16 := int32(load32(m.memory[int64(uint32(v2))+4:]))
		v3 = t16
		t17 := int32(load32(m.memory[uint32(v2):]))
		v1 = t17
	}
l1:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
	store32(m.memory[uint32(v0):], uint32(v1))
	m.g0 = v2 + i32(16)
}
func (m *Module) fn891(v0, v1 int32) {
	var v2, v3 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		v2 = t0
		t1 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		t2 := v2
		v3 = t1
		if uint32(t2) <= uint32(v3) {
			goto l0
		}
		m.fn151(i32(0), v2, v3, i32(1148300))
		panic("unreachable")
	}
l0:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v2))
	t3 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	store32(m.memory[uint32(v0):], uint32(t3))
}
func (m *Module) fn892(v0, v1 int32) {
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
func (m *Module) fn893(v0 int32) {
	var v1 int32
	{
		t0 := int32(load32(m.memory[uint32(v0):]))
		v1 = t0
		p1 := i32(1)
		if uint32(v1) > uint32(i32(2)) {
			p1 = v1 + i32(-3)
		}
		switch p1 {
		case 0:
			t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t3 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			m.fn16(t2, t3)
			return
		case 1:
			m.fn894(v0 + i32(16))
			t4 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t5 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			m.fn16(t4, t5)
			return
		case 2:
			t6 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t7 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			m.fn16(t6, t7)
			t8 := int32(load32(m.memory[int64(uint32(v0))+16:]))
			t9 := int32(load32(m.memory[int64(uint32(v0))+20:]))
			m.fn895(t8, t9)
			return
		case 3:
			t10 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t11 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			m.fn16(t10, t11)
			return
		case 4:
			t12 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t13 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			m.fn16(t12, t13)
			fallthrough
		default:
		}
	}
}
func (m *Module) fn894(v0 int32) {
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
	m.fn893(v3)
	v3 = v3 + i32(28)
	goto l1
l0:
	t2 := int32(load32(m.memory[uint32(v0):]))
	m.fn82(t2, v2)
}
func (m *Module) fn895(v0, v1 int32) {
	if v0 < i32(0) {
		return
	}
	m.fn16(v0, v1)
}
func (m *Module) fn896(v0, v1 int32) {
	var v2 int32
	var v3 int64
	var v4, v5, v6, v7 int32
	t0 := m.g0
	v2 = t0 - i32(64)
	m.g0 = v2
	{
		t1 := int32(load32(m.memory[uint32(v1):]))
		if t1 == i32(-2) {
			goto l0
		}
		t2 := int64(load64(m.memory[uint32(v1):]))
		v3 = t2
		store32(m.memory[uint32(v1):], uint32(i32(-1)))
		t3 := int32(load32(m.memory[int64(uint32(v1))+24:]))
		store32(m.memory[int64(uint32(v2))+24:], uint32(t3))
		t4 := int64(load64(m.memory[int64(uint32(v1))+16:]))
		store64(m.memory[int64(uint32(v2))+16:], uint64(t4))
		t5 := int64(load64(m.memory[int64(uint32(v1))+8:]))
		store64(m.memory[int64(uint32(v2))+8:], uint64(t5))
		store64(m.memory[uint32(v2):], uint64(v3))
		if int32(v3) != i32(-1) {
			t6 := int32(load32(m.memory[int64(uint32(v2))+24:]))
			store32(m.memory[int64(uint32(v0))+24:], uint32(t6))
			t7 := int64(load64(m.memory[int64(uint32(v2))+16:]))
			store64(m.memory[int64(uint32(v0))+16:], uint64(t7))
			t8 := int64(load64(m.memory[int64(uint32(v2))+8:]))
			store64(m.memory[int64(uint32(v0))+8:], uint64(t8))
			t9 := int64(load64(m.memory[uint32(v2):]))
			store64(m.memory[uint32(v0):], uint64(t9))
			goto l3
		}
		m.fn897(v1)
		store32(m.memory[uint32(v1):], uint32(i32(-2)))
		goto l2
	}
l0:
	store32(m.memory[uint32(v2):], uint32(i32(-1)))
l2:
	v4 = i32(-1)
	{
		t10 := int32(load32(m.memory[int64(uint32(v1))+28:]))
		if t10 == 0 {
			goto l4
		}
		t11 := int32(load32(m.memory[int64(uint32(v1))+32:]))
		v5 = t11
		t12 := int32(load32(m.memory[int64(uint32(v1))+40:]))
		v6 = t12
	l6:
		{
			if v5 == v6 {
				goto l4
			}
			t13 := v1
			v7 = v5 + i32(28)
			store32(m.memory[int64(uint32(t13))+32:], uint32(v7))
			t14 := int32(load32(m.memory[int64(uint32(v5))+24:]))
			store32(m.memory[int64(uint32(v2))+56:], uint32(t14))
			t15 := int64(load64(m.memory[int64(uint32(v5))+16:]))
			store64(m.memory[int64(uint32(v2))+48:], uint64(t15))
			t16 := int64(load64(m.memory[int64(uint32(v5))+8:]))
			store64(m.memory[int64(uint32(v2))+40:], uint64(t16))
			t17 := int64(load64(m.memory[uint32(v5):]))
			t18 := v2
			v3 = t17
			store64(m.memory[int64(uint32(t18))+32:], uint64(v3))
			if int32(v3) == i32(6) {
				goto l5
			}
			m.fn893(v2 + i32(32))
			v5 = v7
			goto l6
		l5:
		}
		t19 := int64(load64(m.memory[int64(uint32(v5))+20:]))
		store64(m.memory[int64(uint32(v0))+20:], uint64(t19))
		t20 := int64(load64(m.memory[int64(uint32(v5))+12:]))
		store64(m.memory[int64(uint32(v0))+12:], uint64(t20))
		t21 := int64(load64(m.memory[int64(uint32(v5))+4:]))
		store64(m.memory[int64(uint32(v0))+4:], uint64(t21))
		v4 = i32(6)
	}
l4:
	store32(m.memory[uint32(v0):], uint32(v4))
	m.fn898(v2)
l3:
	m.g0 = v2 + i32(64)
}
