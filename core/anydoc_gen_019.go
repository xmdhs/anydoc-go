package core

import (
	"math/bits"
)

func (m *Module) fn807(v0, v1 int32) {
	var v2, v3, v4, v5, v6, v7, v8 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	v3 = i32(0)
	{
		t1 := int32(m.memory[int64(uint32(v1))+29])
		if t1 != 0 {
			goto l3
		}
		v4 = v1 + i32(16)
		t2 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		v5 = t2
	l1:
		{
			t3 := int32(load32(m.memory[int64(uint32(v1))+16:]))
			v6 = t3
			t4 := int32(load32(m.memory[int64(uint32(v1))+20:]))
			v7 = t4
			m.fn572(v2+i32(8), v4)
			t5 := int32(load32(m.memory[int64(uint32(v2))+12:]))
			v8 = t5
			if v8 == i32(96) {
				goto l1
			}
		}
		{
			if v8 == i32(-1) {
				t12 := int32(m.memory[int64(uint32(v1))+29])
				if t12 != 0 {
					goto l3
				}
				m.memory[int64(uint32(v1))+29] = byte(i32(1))
				t13 := int32(load32(m.memory[int64(uint32(v1))+4:]))
				v6 = t13
				t14 := int32(load32(m.memory[uint32(v1):]))
				v8 = t14
				{
					t15 := int32(m.memory[int64(uint32(v1))+28])
					if t15 != 0 {
						goto l4
					}
					if v6 == v8 {
						goto l3
					}
				}
			l4:
				v4 = v6 - v8
				t16 := int32(load32(m.memory[int64(uint32(v1))+8:]))
				v3 = t16 + v8
				goto l3
			}
			t6 := int32(load32(m.memory[uint32(v1):]))
			v4 = t6
			t7 := int32(load32(m.memory[int64(uint32(v2))+8:]))
			t8 := v1
			t9 := v7 - v6
			v8 = t7
			t10 := int32(load32(m.memory[int64(uint32(v1))+16:]))
			t11 := int32(load32(m.memory[int64(uint32(v1))+20:]))
			store32(m.memory[uint32(t8):], uint32(t9+v8+t10-t11))
			v3 = v5 + v4
			v4 = v8 - v4
			goto l3
		}
	}
l3:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v4))
	store32(m.memory[uint32(v0):], uint32(v3))
	m.g0 = v2 + i32(16)
}
func (m *Module) fn808(v0, v1, v2 int32) {
	var v3, v4, v5 int32
	t0 := m.g0
	v3 = t0 - i32(32)
	m.g0 = v3
	{
		if v2 != 0 {
			goto l0
		}
		store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
		store64(m.memory[uint32(v0):], uint64(i64(0x100000000)))
		goto l1
	l0:
		m.fn59(v3+i32(8), v2, i32(1), i32(1))
		store32(m.memory[int64(uint32(v3))+28:], uint32(i32(0)))
		t1 := int64(load64(m.memory[int64(uint32(v3))+8:]))
		store64(m.memory[int64(uint32(v3))+20:], uint64(t1))
		m.fn634(v3+i32(20), v1, v1+i32(1))
		t2 := int32(load32(m.memory[int64(uint32(v3))+28:]))
		v1 = t2
		t3 := int32(load32(m.memory[int64(uint32(v3))+24:]))
		v4 = t3
		v5 = v2
	l6:
		v5 = int32(uint32(v5) >> 1)
		if v5 != 0 {
			if v1 == 0 {
				goto l5
			}
			memory_copy(m.memory, uint32(v4+v1), uint32(v4), uint32(v1))
		l5:
			v1 = v1 << 1
			goto l6
		}
		store32(m.memory[int64(uint32(v3))+28:], uint32(v1))
		if v2 != v1 {
			goto l3
		}
		goto l4
	l3:
		v5 = v2 - v1
		if v5 == 0 {
			goto l7
		}
		memory_copy(m.memory, uint32(v4+v1), uint32(v4), uint32(v5))
	l7:
		store32(m.memory[int64(uint32(v3))+28:], uint32(v2))
	l4:
		t4 := int32(load32(m.memory[int64(uint32(v3))+28:]))
		store32(m.memory[int64(uint32(v0))+8:], uint32(t4))
		t5 := int64(load64(m.memory[int64(uint32(v3))+20:]))
		store64(m.memory[uint32(v0):], uint64(t5))
	}
l1:
	m.g0 = v3 + i32(32)
}
func (m *Module) fn809(v0 int32, v1 int64) {
	var v2 int32
	t0 := m.g0
	v2 = t0 - i32(32)
	m.g0 = v2
	m.fn813(v2, v1, v2+i32(12), i32(20))
	t1 := int32(load32(m.memory[uint32(v2):]))
	t2 := int32(load32(m.memory[int64(uint32(v2))+4:]))
	m.fn51(v0, t1, t2)
	m.g0 = v2 + i32(32)
}
func (m *Module) fn810(v0 int32, v1 int64) {
	var v2, v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14 int32
	t0 := m.g0
	v2 = t0 - i32(32)
	m.g0 = v2
	if v1 != i64(0) {
		store32(m.memory[int64(uint32(v2))+12:], uint32(i32(0)))
		store64(m.memory[int64(uint32(v2))+4:], uint64(i64(0x100000000)))
	l10:
		{
			if v1 != i64(0) {
				t15 := v2 + i32(4)
				v1 = v1 + i64(-1)
				t16 := int64(uint64(v1) / uint64(i64(26)))
				t17 := v1
				v1 = t16
				m.fn145(t15, int32(t17-v1*i64(26))+i32(97))
				goto l10
			}
			t1 := int32(load32(m.memory[int64(uint32(v2))+8:]))
			t2 := v2 + i32(16)
			v3 = t1
			t3 := int32(load32(m.memory[int64(uint32(v2))+12:]))
			t4 := v3
			v4 = t3
			v5 = int32(uint32(v4) >> 1)
			m.fn814(t2, t4, v5, v5, i32(1301108))
			t5 := int32(load32(m.memory[int64(uint32(v2))+20:]))
			v6 = t5
			t6 := int32(load32(m.memory[int64(uint32(v2))+16:]))
			v7 = t6
			m.fn814(v2+i32(16), v3+v4-v5, v5, v5, i32(1301124))
			t7 := int32(load32(m.memory[int64(uint32(v2))+20:]))
			v8 = t7
			t8 := int32(load32(m.memory[int64(uint32(v2))+16:]))
			v9 = t8
			v10 = i32(0)
			v11 = v5 + i32(-1)
			v5 = v11
		l6:
			if v5 == i32(-1) {
				t12 := int32(load32(m.memory[int64(uint32(v2))+4:]))
				v5 = t12
				m.fn12(v2+i32(16), v3, v4)
				{
					{
						t13 := int32(load32(m.memory[int64(uint32(v2))+16:]))
						if t13 != 0 {
							goto l7
						}
						v1 = int64(uint32(v4))
						goto l8
					}
				l7:
					if v5 != i32(-1) {
						store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
						store64(m.memory[uint32(v0):], uint64(i64(0x100000000)))
						m.fn16(v5, v3)
						goto l1
					}
					t14 := int64(load64(m.memory[int64(uint32(v2))+20:]))
					v1 = t14
					v5 = v3
					v3 = v4
				}
			l8:
				store32(m.memory[int64(uint32(v0))+8:], uint32(v1))
				store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
				store32(m.memory[uint32(v0):], uint32(v5))
				goto l1
			}
			if v6 == v10 {
				m.fn158(v6, v6, i32(1301140))
				panic("unreachable")
			}
			{
				if uint32(v11) >= uint32(v8) {
					m.fn158(v5, v8, i32(1301156))
					panic("unreachable")
				}
				v12 = v7 + v10
				t9 := int32(m.memory[uint32(v12)])
				v13 = t9
				t10 := v12
				v14 = v9 + v5
				t11 := int32(m.memory[uint32(v14)])
				m.memory[uint32(t10)] = byte(t11)
				m.memory[uint32(v14)] = byte(v13)
				v5 = v5 + i32(-1)
				v10 = v10 + i32(1)
				goto l6
			}
		}
	}
	m.fn51(v0, i32(1108008), i32(1))
	goto l1
l1:
	m.g0 = v2 + i32(32)
}
func (m *Module) fn811(v0, v1, v2 int32) {
	var v3 int32
	t0 := m.g0
	v3 = t0 - i32(16)
	m.g0 = v3
	m.fn51(v3+i32(4), v1, v2)
	t1 := int32(load32(m.memory[int64(uint32(v3))+12:]))
	t2 := v0
	v2 = t1
	store32(m.memory[int64(uint32(t2))+8:], uint32(v2))
	t3 := int64(load64(m.memory[int64(uint32(v3))+4:]))
	store64(m.memory[uint32(v0):], uint64(t3))
	t4 := int32(load32(m.memory[int64(uint32(v3))+8:]))
	v0 = t4
l1:
	{
		if v2 == 0 {
			goto l0
		}
		t5 := int32(m.memory[uint32(v0)])
		t6 := v0
		v1 = t5
		p7 := i32(0)
		if uint32((v1+i32(-97))&i32(255)) < uint32(i32(26)) {
			p7 = i32(32)
		}
		m.memory[uint32(t6)] = byte(p7 ^ v1)
		v2 = v2 + i32(-1)
		v0 = v0 + i32(1)
		goto l1
	}
l0:
	m.g0 = v3 + i32(16)
}
func (m *Module) fn812(v0 int32, v1 int64) {
	var v2, v3, v4, v5, v6, v7 int32
	var v8 int64
	t0 := m.g0
	v2 = t0 - i32(240)
	m.g0 = v2
	{
		if uint64(v1+i64(-4000)) > uint64(i64(-4000)) {
			goto l0
		}
		m.fn809(v0, v1)
		goto l1
	l0:
		v3 = i32(0)
		store32(m.memory[int64(uint32(v2))+20:], uint32(i32(0)))
		store64(m.memory[int64(uint32(v2))+12:], uint64(i64(0x100000000)))
		v4 = v2 + i32(32)
		memory_copy(m.memory, uint32(v4), uint32(i32(1080280)), uint32(i32(208)))
	l3:
		{
			if v3 == i32(13) {
				goto l2
			}
			v5 = v4 + v3<<4
			t1 := int32(load32(m.memory[int64(uint32(v5))+8:]))
			v6 = t1
			if v6 == 0 {
				goto l2
			}
			v3 = v3 + i32(1)
			t2 := int32(load32(m.memory[int64(uint32(v5))+12:]))
			v7 = t2
			t3 := int64(load64(m.memory[uint32(v5):]))
			v8 = t3
		l4:
			if uint64(v1) < uint64(v8) {
				goto l3
			}
			m.fn75(v2+i32(12), v6, v7)
			v1 = v1 - v8
			goto l4
		}
	l2:
		t4 := int32(load32(m.memory[int64(uint32(v2))+20:]))
		store32(m.memory[int64(uint32(v0))+8:], uint32(t4))
		t5 := int64(load64(m.memory[int64(uint32(v2))+12:]))
		store64(m.memory[uint32(v0):], uint64(t5))
	}
l1:
	m.g0 = v2 + i32(240)
}
func (m *Module) fn813(v0 int32, v1 int64, v2, v3 int32) {
	var v4 int64
	var v5, v6, v7 int32
	var v8 int64
	var v9, v10 int32
	v4 = v1
	v5 = v3
	if uint64(v1) < uint64(i64(1000)) {
		goto l0
	}
	v6 = v2 + i32(-4)
	v5 = v3
	v4 = v1
l1:
	{
		v7 = v6 + v5
		t0 := v7
		v8 = v4
		t1 := int64(uint64(v8) / uint64(i64(10000)))
		t2 := v8
		v4 = t1
		v9 = int32(t2 - v4*i64(10000))
		t3 := int32(uint32(v9&i32(0xffff)) / uint32(i32(100)))
		v10 = t3
		t4 := int32(load16(m.memory[int64(uint32(v10<<1))+1109319:]))
		store16(m.memory[uint32(t0):], uint16(t4))
		t5 := int32(load16(m.memory[int64(uint32((v9-v10*i32(100))&i32(0xffff)<<1))+1109319:]))
		store16(m.memory[uint32(v7+i32(2)):], uint16(t5))
		v5 = v5 + i32(-4)
		if uint64(v8) > uint64(i64(9999999)) {
			goto l1
		}
	}
l0:
	{
		if uint64(v4) <= uint64(i64(9)) {
			goto l2
		}
		t6 := v2
		v5 = v5 + i32(-2)
		t7 := t6 + v5
		v7 = int32(v4)
		t8 := int32(uint32(v7&i32(0xffff)) / uint32(i32(100)))
		t9 := v7
		v7 = t8
		t10 := int32(load16(m.memory[int64(uint32((t9-v7*i32(100))&i32(0xffff)<<1))+1109319:]))
		store16(m.memory[uint32(t7):], uint16(t10))
		v4 = int64(uint32(v7))
	}
l2:
	{
		if v1 == 0 {
			goto l3
		}
		if v4 == 0 {
			goto l4
		}
	l3:
		t11 := v2
		v5 = v5 + i32(-1)
		t12 := int32(m.memory[int64(uint32(int32(v4)<<1))+1109320])
		m.memory[uint32(t11+v5)] = byte(t12)
	}
l4:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v3-v5))
	store32(m.memory[uint32(v0):], uint32(v2+v5))
}
func (m *Module) fn814(v0, v1, v2, v3, v4 int32) {
	if uint32(v2) >= uint32(v3) {
		goto l0
	}
	m.fn91(i32(1301172), i32(19), v4)
	panic("unreachable")
l0:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
	store32(m.memory[uint32(v0):], uint32(v1))
	store32(m.memory[int64(uint32(v0))+12:], uint32(v2-v3))
	store32(m.memory[int64(uint32(v0))+8:], uint32(v1+v3))
}
func (m *Module) fn815(v0, v1, v2, v3, v4 int32) {
	if uint32(v2) >= uint32(v3) {
		goto l0
	}
	m.fn151(v3, v2, v2, v4)
	panic("unreachable")
l0:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v2-v3))
	store32(m.memory[uint32(v0):], uint32(v1+v3<<2))
}
func (m *Module) fn816(v0, v1, v2 int32) int32 {
	var v3, v4 int32
	v1 = v1 << 2
	var _ int32
l2:
	v3 = i32(1)
	{
		if v1 == 0 {
			goto l0
		}
		t1 := int32(load32(m.memory[uint32(v0):]))
		v4 = t1
		if v4 == i32(10) {
			goto l0
		}
		if v4 == v2 {
			goto l1
		}
		if v4 == i32(32) {
			goto l1
		}
		if v4 == i32(9) {
			goto l1
		}
		v3 = i32(0)
	}
l0:
	return v3
l1:
	v0 = v0 + i32(4)
	v1 = v1 + i32(-4)
	goto l2
}
func (m *Module) fn817(v0 int32) int32 {
	var v1 int32
	v1 = i32(1)
	{
		if uint32(v0&i32(2097119)+i32(-65)) < uint32(i32(26)) {
			goto l0
		}
		if uint32(v0+i32(-48)) < uint32(i32(10)) {
			goto l0
		}
		if uint32(v0) >= uint32(i32(170)) {
			goto l1
		}
		return i32(0)
	l1:
		t0 := m.fn819(v0)
		if t0 != 0 {
			goto l0
		}
		v1 = i32(0)
		if uint32(v0) < uint32(i32(178)) {
			goto l0
		}
		t1 := m.fn820(v0)
		v1 = t1
	}
l0:
	return v1
}
func (m *Module) fn818(v0 int32) int32 {
	if v0 != i32(-1) {
		t0 := m.fn817(v0)
		return t0
	}
	return i32(0)
}
func (m *Module) fn819(v0 int32) int32 {
	var v1, v2, v3, v4, v5 int32
	v1 = i32(0)
	p0 := i32(25)
	if uint32(v0) < uint32(i32(92729)) {
		p0 = i32(0)
	}
	v2 = p0
	t1 := v2
	v2 = v2 + i32(13)
	t2 := int32(load32(m.memory[int64(uint32(v2<<2))+1111464:]))
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
	t7 := int32(load32(m.memory[int64(uint32(v3<<2))+1111464:]))
	p8 := v3
	if uint32(t7<<11) > uint32(v2) {
		p8 = t6
	}
	v3 = p8
	t9 := v3
	v3 = v3 + i32(3)
	t10 := int32(load32(m.memory[int64(uint32(v3<<2))+1111464:]))
	p11 := v3
	if uint32(t10<<11) > uint32(v2) {
		p11 = t9
	}
	v3 = p11
	t12 := v3
	v3 = v3 + i32(2)
	t13 := int32(load32(m.memory[int64(uint32(v3<<2))+1111464:]))
	p14 := v3
	if uint32(t13<<11) > uint32(v2) {
		p14 = t12
	}
	v3 = p14
	t15 := v3
	v3 = v3 + i32(1)
	t16 := int32(load32(m.memory[int64(uint32(v3<<2))+1111464:]))
	p17 := v3
	if uint32(t16<<11) > uint32(v2) {
		p17 = t15
	}
	v3 = p17
	t18 := int32(load32(m.memory[int64(uint32(v3<<2))+1111464:]))
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
	v5 = v2 + i32(1111464)
	t21 := int32(load32(m.memory[int64(uint32(v2))+1111464:]))
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
		t24 := int32(m.memory[uint32(v2+i32(1104040))])
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
func (m *Module) fn820(v0 int32) int32 {
	var v1, v2, v3, v4, v5 int32
	v1 = i32(0)
	p0 := i32(21)
	if uint32(v0) < uint32(i32(70736)) {
		p0 = i32(0)
	}
	v2 = p0
	t1 := v2
	v2 = v2 + i32(11)
	t2 := int32(load32(m.memory[int64(uint32(v2<<2))+1115120:]))
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
	t7 := int32(load32(m.memory[int64(uint32(v3<<2))+1115120:]))
	p8 := v3
	if uint32(t7<<11) > uint32(v2) {
		p8 = t6
	}
	v3 = p8
	t9 := v3
	v3 = v3 + i32(3)
	t10 := int32(load32(m.memory[int64(uint32(v3<<2))+1115120:]))
	p11 := v3
	if uint32(t10<<11) > uint32(v2) {
		p11 = t9
	}
	v3 = p11
	t12 := v3
	v3 = v3 + i32(1)
	t13 := int32(load32(m.memory[int64(uint32(v3<<2))+1115120:]))
	p14 := v3
	if uint32(t13<<11) > uint32(v2) {
		p14 = t12
	}
	v3 = p14
	t15 := v3
	v3 = v3 + i32(1)
	t16 := int32(load32(m.memory[int64(uint32(v3<<2))+1115120:]))
	p17 := v3
	if uint32(t16<<11) > uint32(v2) {
		p17 = t15
	}
	v3 = p17
	t18 := int32(load32(m.memory[int64(uint32(v3<<2))+1115120:]))
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
	v5 = v2 + i32(1115120)
	t21 := int32(load32(m.memory[int64(uint32(v2))+1115120:]))
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
		t24 := int32(m.memory[uint32(v2+i32(1107245))])
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
func (m *Module) fn821(v0, v1, v2, v3 int32) {
	var v4, v5, v6, v7, v8, v9, v10, v11, v12, v13 int32
	t0 := m.g0
	v4 = t0 - i32(64)
	m.g0 = v4
	v5 = i32(0)
	store32(m.memory[int64(uint32(v4))+40:], uint32(i32(0)))
	store64(m.memory[int64(uint32(v4))+32:], uint64(i64(0x400000000)))
	v6 = v2 * i32(28)
	v7 = v3 + i32(32)
l18:
	{
		if v6 == v5 {
			t3 := int32(load32(m.memory[int64(uint32(v4))+40:]))
			store32(m.memory[int64(uint32(v0))+8:], uint32(t3))
			t4 := int64(load64(m.memory[int64(uint32(v4))+32:]))
			store64(m.memory[uint32(v0):], uint64(t4))
			m.g0 = v4 + i32(64)
			return
		}
		v2 = v1 + v5
		t1 := int32(load32(m.memory[uint32(v2):]))
		v8 = t1
		p2 := i32(1)
		if uint32(v8) > uint32(i32(2)) {
			p2 = v8 + i32(-3)
		}
		switch p2 {
		case 1:
			t17 := int32(load32(m.memory[uint32(v2+i32(12)):]))
			if t17 == 0 {
				t19 := int32(load32(m.memory[uint32(v2+i32(20)):]))
				v8 = t19
				t20 := int32(load32(m.memory[uint32(v2+i32(24)):]))
				t21 := v8
				v2 = t20
				t22 := m.fn23(t21, v2)
				if t22 != 0 {
					goto l7
				}
				m.fn821(v4+i32(48), v8, v2, v3)
				t23 := int32(load32(m.memory[int64(uint32(v4))+52:]))
				v11 = t23
				t24 := int32(load32(m.memory[int64(uint32(v4))+48:]))
				v9 = t24
				{
					{
						t25 := int32(load32(m.memory[int64(uint32(v4))+56:]))
						v2 = t25
						t26 := int32(load32(m.memory[int64(uint32(v4))+32:]))
						t27 := int32(load32(m.memory[int64(uint32(v4))+40:]))
						t28 := v2
						v8 = t27
						if uint32(t28) <= uint32(t26-v8) {
							goto l12
						}
						m.fn62(v4+i32(32), v8, v2, i32(4), i32(16))
						t29 := int32(load32(m.memory[int64(uint32(v4))+40:]))
						v8 = t29
						goto l13
					}
				l12:
					if v2 == 0 {
						goto l14
					}
				l13:
					v10 = v2 << 4
					if v10 == 0 {
						goto l14
					}
					t30 := int32(load32(m.memory[int64(uint32(v4))+36:]))
					memory_copy(m.memory, uint32(t30+v8<<4), uint32(v11), uint32(v10))
				}
			l14:
				store32(m.memory[int64(uint32(v4))+40:], uint32(v8+v2))
				m.fn419(v9, v11)
				goto l7
			}
			store32(m.memory[int64(uint32(v4))+60:], uint32(v2))
			store32(m.memory[int64(uint32(v4))+48:], uint32(i32(-0x80000000)))
			t18 := int64(load64(m.memory[uint32(v2+i32(20)):]))
			store64(m.memory[int64(uint32(v4))+52:], uint64(t18))
			m.fn832(v4+i32(32), v4+i32(48))
			goto l7
		case 2:
			store32(m.memory[int64(uint32(v4))+48:], uint32(i32(-0x7fffffff)))
			store32(m.memory[int64(uint32(v4))+60:], uint32(v2+i32(16)))
			t5 := int64(load64(m.memory[uint32(v2+i32(8)):]))
			store64(m.memory[int64(uint32(v4))+52:], uint64(t5))
			m.fn832(v4+i32(32), v4+i32(48))
			goto l7
		case 3:
			t31 := int32(load32(m.memory[uint32(v2+i32(8)):]))
			t32 := v4 + i32(24)
			t33 := v7
			v8 = t31
			t34 := int32(load32(m.memory[uint32(v2+i32(12)):]))
			t35 := v8
			v2 = t34
			m.fn829(t32, t33, t35, v2)
			t36 := int32(load32(m.memory[int64(uint32(v4))+24:]))
			if t36 == 0 {
				goto l7
			}
			store32(m.memory[int64(uint32(v4))+56:], uint32(v2))
			store32(m.memory[int64(uint32(v4))+52:], uint32(v8))
			store32(m.memory[int64(uint32(v4))+48:], uint32(i32(-0x7ffffffe)))
			m.fn832(v4+i32(32), v4+i32(48))
			goto l7
		case 4:
			store32(m.memory[int64(uint32(v4))+48:], uint32(i32(-0x7ffffffd)))
			t6 := int64(load64(m.memory[uint32(v2+i32(8)):]))
			store64(m.memory[int64(uint32(v4))+52:], uint64(t6))
			m.fn832(v4+i32(32), v4+i32(48))
			goto l7
		case 5:
			store32(m.memory[int64(uint32(v4))+48:], uint32(i32(-0x7ffffffc)))
			m.fn832(v4+i32(32), v4+i32(48))
			goto l7
		default:
			t7 := int32(load32(m.memory[uint32(v2+i32(12)):]))
			v8 = t7
			if v8 == 0 {
				goto l7
			}
			t8 := int32(load32(m.memory[uint32(v2+i32(8)):]))
			t9 := v4 + i32(16)
			v9 = t8
			m.fn46(t9, v9, v8)
			{
				{
					t10 := int32(load32(m.memory[int64(uint32(v4))+20:]))
					if t10 != 0 {
						goto l8
					}
					v2 = i32(0)
					goto l9
				}
			l8:
				t11 := int32(load32(m.memory[uint32(v2+i32(16)):]))
				v2 = t11
			}
		l9:
			store32(m.memory[int64(uint32(v4))+44:], uint32(v2))
			t12 := int32(load32(m.memory[int64(uint32(v4))+36:]))
			v10 = t12
			t13 := int32(load32(m.memory[int64(uint32(v4))+40:]))
			v11 = t13
			if v11 == 0 {
				goto l10
			}
			v12 = v10 + v11<<4
			v13 = v12 + i32(-16)
			if v13 == 0 {
				goto l10
			}
			t14 := int32(load32(m.memory[uint32(v13):]))
			if t14 <= i32(-0x7ffffffc) {
				goto l10
			}
			t15 := m.fn823(v12+i32(-4), v4+i32(44))
			if t15 == 0 {
				goto l10
			}
			t16 := m.fn833(v13)
			m.fn75(t16, v9, v8)
			goto l7
		}
	}
l10:
	{
		t37 := m.fn822(v4+i32(44), i32(1287584))
		if t37 == 0 {
			goto l15
		}
		if v2&i32(0x1000000) != 0 {
			goto l15
		}
		if uint32(v11) < uint32(i32(2)) {
			goto l15
		}
		t38 := v10
		t39 := v11
		v12 = v11 + i32(-1)
		t40 := m.fn834(t38, t39, v12, i32(1084924))
		v13 = t40
		t41 := int32(load32(m.memory[uint32(v13):]))
		if t41 < i32(-0x7ffffffb) {
			goto l15
		}
		t42 := m.fn823(v13+i32(12), i32(1287584))
		if t42 == 0 {
			goto l15
		}
		t43 := int32(load32(m.memory[uint32(v13+i32(4)):]))
		t44 := int32(load32(m.memory[uint32(v13+i32(8)):]))
		m.fn46(v4+i32(8), t43, t44)
		t45 := int32(load32(m.memory[int64(uint32(v4))+12:]))
		if t45 != 0 {
			goto l15
		}
		t46 := m.fn834(v10, v11, v11+i32(-2), i32(1084940))
		v11 = t46
		t47 := int32(load32(m.memory[uint32(v11):]))
		if t47 < i32(-0x7ffffffb) {
			goto l15
		}
		t48 := m.fn823(v11+i32(12), v4+i32(44))
		if t48 == 0 {
			goto l15
		}
		store32(m.memory[int64(uint32(v4))+40:], uint32(v12))
		{
			v2 = v10 + v12<<4
			t49 := int32(load32(m.memory[uint32(v2):]))
			v11 = t49
			if v11 == i32(-2) {
				goto l16
			}
			if v11 <= i32(-0x7ffffffc) {
				goto l16
			}
			v10 = v2 + i32(-16)
			if v10 == 0 {
				goto l17
			}
			t50 := int32(load32(m.memory[uint32(v10):]))
			if t50 <= i32(-0x7ffffffc) {
				goto l17
			}
			t51 := int32(load32(m.memory[int64(uint32(v2))+8:]))
			v13 = t51
			t52 := int32(load32(m.memory[int64(uint32(v2))+4:]))
			v2 = t52
			t53 := m.fn833(v10)
			v10 = t53
			m.fn75(v10, v2, v13)
			m.fn75(v10, v9, v8)
			m.fn134(v11, v2)
			goto l7
		}
	l16:
		m.fn256(i32(1286542), i32(40), i32(1084972))
		panic("unreachable")
	l17:
		m.fn256(i32(1286542), i32(40), i32(1084956))
		panic("unreachable")
	}
l15:
	store32(m.memory[int64(uint32(v4))+60:], uint32(v2))
	store32(m.memory[int64(uint32(v4))+56:], uint32(v8))
	store32(m.memory[int64(uint32(v4))+52:], uint32(v9))
	store32(m.memory[int64(uint32(v4))+48:], uint32(i32(-1)))
	m.fn832(v4+i32(32), v4+i32(48))
l7:
	v5 = v5 + i32(28)
	goto l18
}
func (m *Module) fn822(v0, v1 int32) int32 {
	t0 := m.fn823(v0, v1)
	return t0 ^ i32(1)
}
func (m *Module) fn823(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(v1):]))
	var p2 int32
	if t0 == t1 {
		p2 = 1
	}
	return p2
}
func (m *Module) fn824(v0, v1, v2 int32) {
	var v3, v4, v5, v6, v7, v8, v9 int32
	v3 = v1 + v2
	v4 = i32(0)
	v5 = v1
l7:
	{
		v6 = v4
		if v5 != v3 {
			goto l0
		}
		v6 = v2
		goto l1
	l0:
		{
			{
				t0 := int32(int8(m.memory[uint32(v5)]))
				v7 = t0
				if v7 <= i32(-1) {
					goto l2
				}
				v8 = v5 + i32(1)
				v7 = v7 & i32(255)
				goto l3
			}
		l2:
			t1 := int32(m.memory[int64(uint32(v5))+1])
			v8 = t1 & i32(63)
			v4 = v7 & i32(31)
			if uint32(v7) > uint32(i32(-33)) {
				goto l4
			}
			v7 = v4<<6 | v8
			v8 = v5 + i32(2)
			goto l3
		l4:
			t2 := int32(m.memory[int64(uint32(v5))+2])
			v8 = v8<<6 | t2&i32(63)
			if uint32(v7) >= uint32(i32(-16)) {
				goto l5
			}
			v7 = v8 | v4<<12
			v8 = v5 + i32(3)
			goto l3
		l5:
			t3 := int32(m.memory[int64(uint32(v5))+3])
			v7 = v8<<6 | t3&i32(63) | v4<<18&i32(0x1c0000)
			v8 = v5 + i32(4)
		}
	l3:
		v4 = v6 - v5 + v8
		v9 = v7 + i32(-9)
		if uint32(v9) > uint32(i32(23)) {
			goto l6
		}
		v5 = v8
		if i32_shl(i32(1), v9)&i32(8388639) != 0 {
			goto l7
		}
	l6:
		if uint32(v7) < uint32(i32(133)) {
			goto l1
		}
		v5 = int32(uint32(v7) >> 8)
		if v5 == 0 {
			v5 = v8
			t4 := int32(m.memory[int64(uint32(v7&i32(255)))+1148316])
			if t4&i32(1) != 0 {
				goto l7
			}
			goto l1
		}
		if v5 == i32(48) {
			v5 = v8
			if v7 == i32(12288) {
				goto l7
			}
			goto l1
		}
		if v5 == i32(32) {
			goto l10
		}
		if v5 != i32(22) {
			goto l1
		}
		v5 = v8
		if v7 == i32(5760) {
			goto l7
		}
		goto l1
	l10:
		v5 = v8
		t5 := int32(m.memory[int64(uint32(v7&i32(255)))+1148316])
		if t5&i32(2) != 0 {
			goto l7
		}
	}
l1:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v2-v6))
	store32(m.memory[uint32(v0):], uint32(v1+v6))
}
func (m *Module) fn825(v0, v1, v2, v3, v4 int32) {
	if v3 == 0 {
		goto l0
	}
	{
		if uint32(v3) < uint32(v2) {
			goto l1
		}
		if v3 != v2 {
			goto l2
		}
		goto l0
	l1:
		t0 := int32(int8(m.memory[uint32(v1+v3)]))
		if t0 > i32(-65) {
			goto l0
		}
	}
l2:
	m.fn556(v1, v2, i32(0), v3, v4)
	panic("unreachable")
l0:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
	store32(m.memory[uint32(v0):], uint32(v1))
}
func (m *Module) fn826(v0, v1, v2, v3, v4 int32) {
	var v5, v6 int32
	t0 := m.g0
	v5 = t0 - i32(16)
	m.g0 = v5
	m.fn786(v5+i32(8), v1, v2, v3)
	{
		t1 := int32(load32(m.memory[int64(uint32(v5))+8:]))
		v6 = t1
		if v6 == 0 {
			m.fn556(v2, v3, v1, v3, v4)
			panic("unreachable")
		}
		t2 := int32(load32(m.memory[int64(uint32(v5))+12:]))
		v3 = t2
		store32(m.memory[uint32(v0):], uint32(v6))
		store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
		m.g0 = v5 + i32(16)
		return
	}
}
func (m *Module) fn827(v0, v1, v2 int32) int32 {
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
			t6 := m.fn648(t4, t5, v3, v1, v2)
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
func (m *Module) fn828(v0, v1, v2 int32) {
	var v3, v4 int32
	t0 := m.g0
	v3 = t0 - i32(32)
	m.g0 = v3
	m.fn140(v3+i32(8), v2)
	store32(m.memory[int64(uint32(v3))+24:], uint32(v1))
	store32(m.memory[int64(uint32(v3))+28:], uint32(v1+v2))
	{
	l10:
		{
			{
				t1 := m.fn48(v3 + i32(24))
				v2 = t1
				switch v2 + i32(-60) {
				case 0:
					m.fn75(v3+i32(8), i32(1084679), i32(3))
					goto l10
				case 1:
					goto l1
				case 2:
					m.fn75(v3+i32(8), i32(1084682), i32(3))
					goto l10
				default:
					if v2 == i32(124) {
						m.fn75(v3+i32(8), i32(1084685), i32(3))
						goto l10
					}
					if v2 != i32(-1) {
						goto l1
					}
					t2 := int32(load32(m.memory[int64(uint32(v3))+12:]))
					t3 := v3
					v2 = t2
					t4 := int32(load32(m.memory[int64(uint32(v3))+16:]))
					store32(m.memory[int64(uint32(t3))+28:], uint32(v2+t4))
					store32(m.memory[int64(uint32(v3))+24:], uint32(v2))
				l7:
					{
						t5 := m.fn48(v3 + i32(24))
						v2 = t5
						if v2 == i32(-1) {
							goto l5
						}
						t6 := m.fn630(v2)
						v1 = t6
						if v2&i32(0x1ffffe) == i32(40) {
							goto l6
						}
						if v1 == 0 {
							goto l7
						}
					l6:
					}
					store32(m.memory[int64(uint32(v3))+28:], uint32(i32(25)))
					store32(m.memory[int64(uint32(v3))+24:], uint32(v3+i32(8)))
					m.fn73(v0, i32(1068653), v3+i32(24))
					t7 := int32(load32(m.memory[int64(uint32(v3))+8:]))
					t8 := int32(load32(m.memory[int64(uint32(v3))+12:]))
					m.fn16(t7, t8)
					goto l8
				}
			}
		l1:
			if uint32(v2) < uint32(i32(32)) {
				goto l9
			}
			if uint32(v2+i32(-127)) < uint32(i32(33)) {
				goto l9
			}
			m.fn74(v3+i32(8), v2)
			goto l10
		l9:
			store32(m.memory[int64(uint32(v3))+20:], uint32(i32(0)))
			m.fn522(v3, v2, v3+i32(20))
			t9 := int32(load32(m.memory[int64(uint32(v3))+4:]))
			v2 = t9
			t10 := int32(load32(m.memory[uint32(v3):]))
			v1 = t10
		l11:
			{
				if v2 == 0 {
					goto l10
				}
				t11 := int32(m.memory[uint32(v1)])
				v4 = t11
				m.fn74(v3+i32(8), i32(37))
				t12 := int32(m.memory[int64(uint32(int32(uint32(v4)>>4)))+1131672])
				m.fn74(v3+i32(8), t12)
				t13 := int32(m.memory[int64(uint32(v4&i32(15)))+1131672])
				m.fn74(v3+i32(8), t13)
				v2 = v2 + i32(-1)
				v1 = v1 + i32(1)
				goto l11
			}
		}
	l5:
		t14 := int32(load32(m.memory[int64(uint32(v3))+16:]))
		store32(m.memory[int64(uint32(v0))+8:], uint32(t14))
		t15 := int64(load64(m.memory[int64(uint32(v3))+8:]))
		store64(m.memory[uint32(v0):], uint64(t15))
	}
l8:
	m.g0 = v3 + i32(32)
}
func (m *Module) fn829(v0, v1, v2, v3 int32) {
	var v4 int32
	v4 = i32(0)
	{
		{
			t0 := m.fn827(v1, v2, v3)
			v2 = t0
			if v2 != 0 {
				goto l0
			}
			goto l1
		}
	l0:
		t1 := int32(m.memory[int64(uint32(v2))+12])
		if t1 != i32(1) {
			goto l1
		}
		t2 := int32(load32(m.memory[int64(uint32(v2))+8:]))
		v3 = t2
		t3 := int32(load32(m.memory[int64(uint32(v2))+4:]))
		v4 = t3
	}
l1:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
	store32(m.memory[uint32(v0):], uint32(v4))
}
func (m *Module) fn830(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := m.fn72(t0, v1)
	return t1
}
func (m *Module) fn831(v0, v1, v2 int32) {
	var v3, v4 int32
	t0 := m.g0
	v3 = t0 - i32(16)
	m.g0 = v3
	t1 := v0
	v4 = v2 - v1
	t2 := int32(uint32(v4) >> 2)
	var p3 int32
	if v4&i32(3) != i32(0) {
		p3 = 1
	}
	m.fn47(t1, t2+p3)
	store32(m.memory[int64(uint32(v3))+12:], uint32(v2))
	store32(m.memory[int64(uint32(v3))+8:], uint32(v1))
l1:
	{
		m.fn577(v3, v3+i32(8))
		t4 := int32(load32(m.memory[uint32(v3):]))
		if t4 != i32(1) {
			goto l0
		}
		t5 := int32(load32(m.memory[int64(uint32(v3))+4:]))
		m.fn74(v0, t5)
		goto l1
	}
l0:
	m.g0 = v3 + i32(16)
}
func (m *Module) fn832(v0, v1 int32) {
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
func (m *Module) fn833(v0 int32) int32 {
	var v1, v2 int32
	t0 := m.g0
	v1 = t0 - i32(16)
	m.g0 = v1
	{
		t1 := int32(load32(m.memory[uint32(v0):]))
		if t1 != i32(-1) {
			goto l0
		}
		t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		t3 := v1 + i32(4)
		v2 = t2
		t4 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		m.fn51(t3, v2, t4)
		m.fn134(i32(-1), v2)
		t5 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		store32(m.memory[int64(uint32(v0))+8:], uint32(t5))
		t6 := int64(load64(m.memory[int64(uint32(v1))+4:]))
		store64(m.memory[uint32(v0):], uint64(t6))
		t7 := int32(load32(m.memory[uint32(v0):]))
		if t7 != i32(-1) {
			goto l0
		}
		m.fn256(i32(1286542), i32(40), i32(1077904))
		panic("unreachable")
	}
l0:
	m.g0 = v1 + i32(16)
	return v0
}
func (m *Module) fn834(v0, v1, v2, v3 int32) int32 {
	if uint32(v2) < uint32(v1) {
		return v0 + v2<<4
	}
	m.fn158(v2, v1, v3)
	panic("unreachable")
}
func (m *Module) fn835(v0, v1 int32) {
	var v2, v3, v4, v5 int32
	t0 := m.g0
	v2 = t0 - i32(128)
	m.g0 = v2
	{
		t1 := int32(load32(m.memory[uint32(v1):]))
		v3 = t1
		switch v3 >> 31 & (v3 + i32(-0x7fffffff)) {
		case 2:
			t2 := int32(load32(m.memory[int64(uint32(v1))+24:]))
			v4 = t2
			t3 := int32(load32(m.memory[int64(uint32(v1))+20:]))
			v1 = t3
			v3 = i32(0)
			store32(m.memory[int64(uint32(v2))+36:], uint32(i32(0)))
			store32(m.memory[int64(uint32(v2))+28:], uint32(i32(0)))
			store32(m.memory[int64(uint32(v2))+20:], uint32(v1))
			store32(m.memory[int64(uint32(v2))+24:], uint32(v1+v4*i32(28)))
			m.fn836(v2+i32(68), v2+i32(20))
			{
				t4 := int32(load32(m.memory[int64(uint32(v2))+68:]))
				if t4 == i32(-1) {
					goto l7
				}
				m.fn837(v2+i32(80), v2+i32(20))
				t5 := int32(load32(m.memory[int64(uint32(v2))+80:]))
				t6 := v2
				v1 = t5 + i32(1)
				p7 := i32(-1)
				if v1 != 0 {
					p7 = v1
				}
				v1 = p7
				p8 := i32(4)
				if uint32(v1) > uint32(i32(4)) {
					p8 = v1
				}
				m.fn59(t6, p8, i32(4), i32(12))
				t9 := int32(load32(m.memory[uint32(v2):]))
				v1 = t9
				t10 := int32(load32(m.memory[int64(uint32(v2))+4:]))
				v5 = t10
				t11 := int32(load32(m.memory[int64(uint32(v2))+76:]))
				store32(m.memory[int64(uint32(v5))+8:], uint32(t11))
				t12 := int64(load64(m.memory[int64(uint32(v2))+68:]))
				store64(m.memory[uint32(v5):], uint64(t12))
				store32(m.memory[int64(uint32(v2))+52:], uint32(i32(1)))
				store32(m.memory[int64(uint32(v2))+48:], uint32(v5))
				store32(m.memory[int64(uint32(v2))+44:], uint32(v1))
				t13 := int64(load64(m.memory[int64(uint32(v2))+36:]))
				store64(m.memory[int64(uint32(v2))+96:], uint64(t13))
				t14 := int64(load64(m.memory[int64(uint32(v2))+28:]))
				store64(m.memory[int64(uint32(v2))+88:], uint64(t14))
				t15 := int64(load64(m.memory[int64(uint32(v2))+20:]))
				store64(m.memory[int64(uint32(v2))+80:], uint64(t15))
				v1 = i32(12)
				v3 = i32(1)
			l10:
				{
					m.fn836(v2+i32(104), v2+i32(80))
					t16 := int32(load32(m.memory[int64(uint32(v2))+104:]))
					if t16 == i32(-1) {
						t25 := int32(load32(m.memory[int64(uint32(v2))+52:]))
						v3 = t25
						t26 := int32(load32(m.memory[int64(uint32(v2))+48:]))
						v1 = t26
						goto l11
					}
					{
						t17 := int32(load32(m.memory[int64(uint32(v2))+44:]))
						if v3 != t17 {
							goto l9
						}
						m.fn837(v2+i32(116), v2+i32(80))
						t18 := int32(load32(m.memory[int64(uint32(v2))+116:]))
						t19 := v2 + i32(44)
						v4 = t18 + i32(1)
						p20 := i32(-1)
						if v4 != 0 {
							p20 = v4
						}
						m.fn60(t19, p20)
						t21 := int32(load32(m.memory[int64(uint32(v2))+48:]))
						v5 = t21
					}
				l9:
					v4 = v5 + v1
					t22 := int32(load32(m.memory[int64(uint32(v2))+112:]))
					store32(m.memory[int64(uint32(v4))+8:], uint32(t22))
					t23 := int64(load64(m.memory[int64(uint32(v2))+104:]))
					store64(m.memory[uint32(v4):], uint64(t23))
					t24 := v2
					v3 = v3 + i32(1)
					store32(m.memory[int64(uint32(t24))+52:], uint32(v3))
					v1 = v1 + i32(12)
					goto l10
				}
			}
		l7:
			store32(m.memory[int64(uint32(v2))+52:], uint32(i32(0)))
			store64(m.memory[int64(uint32(v2))+44:], uint64(i64(0x400000000)))
			v1 = i32(4)
		l11:
			m.fn77(v0, v1, v3, i32(1097368), i32(1))
			m.fn78(v2 + i32(44))
			goto l12
		case 3:
			t27 := int32(load32(m.memory[int64(uint32(v1))+12:]))
			v4 = t27
			t28 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			v1 = t28
			v3 = i32(0)
			store32(m.memory[int64(uint32(v2))+36:], uint32(i32(0)))
			store32(m.memory[int64(uint32(v2))+28:], uint32(i32(0)))
			store32(m.memory[int64(uint32(v2))+20:], uint32(v1))
			store32(m.memory[int64(uint32(v2))+24:], uint32(v1+v4*i32(12)))
			m.fn838(v2+i32(68), v2+i32(20))
			{
				t29 := int32(load32(m.memory[int64(uint32(v2))+68:]))
				if t29 == i32(-1) {
					goto l13
				}
				m.fn839(v2+i32(80), v2+i32(20))
				t30 := int32(load32(m.memory[int64(uint32(v2))+80:]))
				t31 := v2 + i32(8)
				v1 = t30 + i32(1)
				p32 := i32(-1)
				if v1 != 0 {
					p32 = v1
				}
				v1 = p32
				p33 := i32(4)
				if uint32(v1) > uint32(i32(4)) {
					p33 = v1
				}
				m.fn59(t31, p33, i32(4), i32(12))
				t34 := int32(load32(m.memory[int64(uint32(v2))+8:]))
				v1 = t34
				t35 := int32(load32(m.memory[int64(uint32(v2))+12:]))
				v5 = t35
				t36 := int32(load32(m.memory[int64(uint32(v2))+76:]))
				store32(m.memory[int64(uint32(v5))+8:], uint32(t36))
				t37 := int64(load64(m.memory[int64(uint32(v2))+68:]))
				store64(m.memory[uint32(v5):], uint64(t37))
				store32(m.memory[int64(uint32(v2))+64:], uint32(i32(1)))
				store32(m.memory[int64(uint32(v2))+60:], uint32(v5))
				store32(m.memory[int64(uint32(v2))+56:], uint32(v1))
				t38 := int64(load64(m.memory[int64(uint32(v2))+36:]))
				store64(m.memory[int64(uint32(v2))+96:], uint64(t38))
				t39 := int64(load64(m.memory[int64(uint32(v2))+28:]))
				store64(m.memory[int64(uint32(v2))+88:], uint64(t39))
				t40 := int64(load64(m.memory[int64(uint32(v2))+20:]))
				store64(m.memory[int64(uint32(v2))+80:], uint64(t40))
				v1 = i32(12)
				v3 = i32(1)
			l16:
				{
					m.fn838(v2+i32(104), v2+i32(80))
					t41 := int32(load32(m.memory[int64(uint32(v2))+104:]))
					if t41 == i32(-1) {
						t50 := int32(load32(m.memory[int64(uint32(v2))+64:]))
						v3 = t50
						t51 := int32(load32(m.memory[int64(uint32(v2))+60:]))
						v1 = t51
						goto l17
					}
					{
						t42 := int32(load32(m.memory[int64(uint32(v2))+56:]))
						if v3 != t42 {
							goto l15
						}
						m.fn839(v2+i32(116), v2+i32(80))
						t43 := int32(load32(m.memory[int64(uint32(v2))+116:]))
						t44 := v2 + i32(56)
						v4 = t43 + i32(1)
						p45 := i32(-1)
						if v4 != 0 {
							p45 = v4
						}
						m.fn60(t44, p45)
						t46 := int32(load32(m.memory[int64(uint32(v2))+60:]))
						v5 = t46
					}
				l15:
					v4 = v5 + v1
					t47 := int32(load32(m.memory[int64(uint32(v2))+112:]))
					store32(m.memory[int64(uint32(v4))+8:], uint32(t47))
					t48 := int64(load64(m.memory[int64(uint32(v2))+104:]))
					store64(m.memory[uint32(v4):], uint64(t48))
					t49 := v2
					v3 = v3 + i32(1)
					store32(m.memory[int64(uint32(t49))+64:], uint32(v3))
					v1 = v1 + i32(12)
					goto l16
				}
			}
		l13:
			store32(m.memory[int64(uint32(v2))+64:], uint32(i32(0)))
			store64(m.memory[int64(uint32(v2))+56:], uint64(i64(0x400000000)))
			v1 = i32(4)
		l17:
			m.fn77(v0, v1, v3, i32(1097368), i32(1))
			m.fn78(v2 + i32(56))
			goto l12
		case 4:
			t52 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			t53 := v2 + i32(80)
			v3 = t52
			t54 := int32(load32(m.memory[int64(uint32(v1))+12:]))
			m.fn840(t53, v3, v3+t54<<5)
			t55 := int32(load32(m.memory[int64(uint32(v2))+84:]))
			t56 := int32(load32(m.memory[int64(uint32(v2))+88:]))
			m.fn77(v0, t55, t56, i32(1097368), i32(1))
			m.fn78(v2 + i32(80))
			goto l12
		case 5:
			t57 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			t58 := int32(load32(m.memory[int64(uint32(v1))+12:]))
			m.fn31(v0, t57, t58)
			goto l12
		case 6:
			store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
			store64(m.memory[uint32(v0):], uint64(i64(0x100000000)))
			goto l12
		case 1:
			v1 = v1 + i32(4)
			fallthrough
		default:
			t59 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t60 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			m.fn45(v0, t59, t60)
		}
	}
l12:
	m.g0 = v2 + i32(128)
}
func (m *Module) fn836(v0, v1 int32) {
	var v2, v3, v4, v5 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	v3 = v1 + i32(8)
l3:
	{
		m.fn842(v2+i32(4), v3)
		{
			t1 := int32(load32(m.memory[int64(uint32(v2))+4:]))
			if t1 == i32(-1) {
				goto l0
			}
			t2 := int32(load32(m.memory[int64(uint32(v2))+12:]))
			store32(m.memory[int64(uint32(v0))+8:], uint32(t2))
			t3 := int64(load64(m.memory[int64(uint32(v2))+4:]))
			store64(m.memory[uint32(v0):], uint64(t3))
			goto l1
		}
	l0:
		t4 := int32(load32(m.memory[int64(uint32(v2))+8:]))
		m.fn134(i32(-1), t4)
		{
			t5 := int32(load32(m.memory[uint32(v1):]))
			v4 = t5
			if v4 == 0 {
				goto l2
			}
			t6 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			if v4 == t6 {
				goto l2
			}
			store32(m.memory[uint32(v1):], uint32(v4+i32(28)))
			t7 := int32(load32(m.memory[int64(uint32(v4))+4:]))
			t8 := v1
			v5 = t7
			store32(m.memory[int64(uint32(t8))+8:], uint32(v5))
			t9 := int32(load32(m.memory[int64(uint32(v4))+8:]))
			store32(m.memory[int64(uint32(v1))+12:], uint32(v5+t9<<5))
			goto l3
		}
	l2:
	}
	m.fn842(v0, v1+i32(16))
l1:
	m.g0 = v2 + i32(16)
}
func (m *Module) fn837(v0, v1 int32) {
	var v2, v3, v4 int32
	v2 = i32(0)
	t0 := int32(load32(m.memory[int64(uint32(v1))+20:]))
	t1 := int32(load32(m.memory[int64(uint32(v1))+16:]))
	v3 = t1
	p2 := i32(0)
	if v3 != 0 {
		p2 = int32(uint32(t0-v3) >> 5)
	}
	t3 := int32(load32(m.memory[int64(uint32(v1))+12:]))
	t4 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	v3 = t4
	p5 := i32(0)
	if v3 != 0 {
		p5 = int32(uint32(t3-v3) >> 5)
	}
	v3 = p2 + p5
	{
		t6 := int32(load32(m.memory[uint32(v1):]))
		v4 = t6
		if v4 == 0 {
			goto l0
		}
		t7 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		if t7 != v4 {
			goto l1
		}
	}
l0:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v3))
	v2 = i32(1)
l1:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v2))
	store32(m.memory[uint32(v0):], uint32(v3))
}
func (m *Module) fn838(v0, v1 int32) {
	var v2, v3, v4, v5 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	v3 = v1 + i32(8)
l3:
	{
		m.fn841(v2+i32(4), v3)
		{
			t1 := int32(load32(m.memory[int64(uint32(v2))+4:]))
			if t1 == i32(-1) {
				goto l0
			}
			t2 := int32(load32(m.memory[int64(uint32(v2))+12:]))
			store32(m.memory[int64(uint32(v0))+8:], uint32(t2))
			t3 := int64(load64(m.memory[int64(uint32(v2))+4:]))
			store64(m.memory[uint32(v0):], uint64(t3))
			goto l1
		}
	l0:
		t4 := int32(load32(m.memory[int64(uint32(v2))+8:]))
		m.fn134(i32(-1), t4)
		{
			t5 := int32(load32(m.memory[uint32(v1):]))
			v4 = t5
			if v4 == 0 {
				goto l2
			}
			t6 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			if v4 == t6 {
				goto l2
			}
			store32(m.memory[uint32(v1):], uint32(v4+i32(12)))
			t7 := int32(load32(m.memory[int64(uint32(v4))+4:]))
			t8 := v1
			v5 = t7
			store32(m.memory[int64(uint32(t8))+8:], uint32(v5))
			t9 := int32(load32(m.memory[int64(uint32(v4))+8:]))
			store32(m.memory[int64(uint32(v1))+12:], uint32(v5+t9*i32(20)))
			goto l3
		}
	l2:
	}
	m.fn841(v0, v1+i32(16))
l1:
	m.g0 = v2 + i32(16)
}
func (m *Module) fn839(v0, v1 int32) {
	var v2, v3, v4, v5 int32
	v2 = i32(0)
	v3 = i32(0)
	{
		t0 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		v4 = t0
		if v4 == 0 {
			goto l0
		}
		t1 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		t2 := int32(uint32(t1-v4) / uint32(i32(20)))
		v3 = t2
	}
l0:
	{
		t3 := int32(load32(m.memory[int64(uint32(v1))+16:]))
		v4 = t3
		if v4 == 0 {
			goto l1
		}
		t4 := int32(load32(m.memory[int64(uint32(v1))+20:]))
		t5 := int32(uint32(t4-v4) / uint32(i32(20)))
		v2 = t5
	}
l1:
	{
		t6 := int32(load32(m.memory[uint32(v1):]))
		v5 = t6
		if v5 == 0 {
			goto l2
		}
		v4 = i32(0)
		t7 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		if t7 != v5 {
			goto l3
		}
	}
l2:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v2+v3))
	v4 = i32(1)
l3:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v4))
	store32(m.memory[uint32(v0):], uint32(i32(0)))
}
func (m *Module) fn840(v0, v1, v2 int32) {
	var v3, v4, v5 int32
	t0 := m.g0
	v3 = t0 - i32(32)
	m.g0 = v3
	t1 := v3
	v4 = int32(uint32(v2-v1) >> 5)
	m.fn59(t1, v4, i32(4), i32(12))
	store32(m.memory[int64(uint32(v3))+16:], uint32(i32(0)))
	t2 := int64(load64(m.memory[uint32(v3):]))
	store64(m.memory[int64(uint32(v3))+8:], uint64(t2))
	m.fn60(v3+i32(8), v4)
	t3 := int32(load32(m.memory[int64(uint32(v3))+16:]))
	v5 = t3
	{
		if v2 == v1 {
			goto l0
		}
		v2 = v5 + v4
		t4 := int32(load32(m.memory[int64(uint32(v3))+12:]))
		v5 = t4 + v5*i32(12)
	l1:
		{
			m.fn835(v3+i32(20), v1)
			t5 := int32(load32(m.memory[int64(uint32(v3))+28:]))
			store32(m.memory[int64(uint32(v5))+8:], uint32(t5))
			t6 := int64(load64(m.memory[int64(uint32(v3))+20:]))
			store64(m.memory[uint32(v5):], uint64(t6))
			v5 = v5 + i32(12)
			v1 = v1 + i32(32)
			v4 = v4 + i32(-1)
			if v4 != 0 {
				goto l1
			}
		}
		v5 = v2
	}
l0:
	t7 := int64(load64(m.memory[int64(uint32(v3))+8:]))
	store64(m.memory[uint32(v0):], uint64(t7))
	store32(m.memory[int64(uint32(v0))+8:], uint32(v5))
	m.g0 = v3 + i32(32)
}
func (m *Module) fn841(v0, v1 int32) {
	var v2, v3, v4, v5, v6 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	{
		t1 := int32(load32(m.memory[uint32(v1):]))
		v3 = t1
		if v3 == 0 {
			store32(m.memory[uint32(v0):], uint32(i32(-1)))
			goto l4
		}
		t2 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		v4 = t2
	l3:
		{
			if v3 == v4 {
				goto l1
			}
			t3 := v1
			v5 = v3 + i32(20)
			store32(m.memory[uint32(t3):], uint32(v5))
			{
				t4 := int32(load32(m.memory[uint32(v3):]))
				if t4 != i32(-1) {
					t5 := int32(load32(m.memory[uint32(v3+i32(4)):]))
					t6 := v2 + i32(4)
					v6 = t5
					t7 := int32(load32(m.memory[uint32(v3+i32(8)):]))
					m.fn840(t6, v6, v6+t7<<5)
					t8 := int32(load32(m.memory[int64(uint32(v2))+8:]))
					t9 := int32(load32(m.memory[int64(uint32(v2))+12:]))
					m.fn77(v0, t8, t9, i32(1097368), i32(1))
					m.fn78(v2 + i32(4))
					v3 = v5
					t10 := int32(load32(m.memory[uint32(v0):]))
					if t10 == i32(-1) {
						goto l3
					}
					goto l4
				}
				store32(m.memory[uint32(v0):], uint32(i32(-1)))
				v3 = v5
				goto l3
			}
		}
	}
l1:
	store32(m.memory[uint32(v1):], uint32(i32(0)))
	store32(m.memory[uint32(v0):], uint32(i32(-1)))
l4:
	m.g0 = v2 + i32(16)
}
func (m *Module) fn842(v0, v1 int32) {
	var v2 int32
	{
		{
			t0 := int32(load32(m.memory[uint32(v1):]))
			v2 = t0
			if v2 == 0 {
				store32(m.memory[uint32(v0):], uint32(i32(-1)))
				return
			}
			t1 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			if v2 != t1 {
				goto l1
			}
			store32(m.memory[uint32(v0):], uint32(i32(-1)))
			goto l2
		}
	l1:
		store32(m.memory[uint32(v1):], uint32(v2+i32(32)))
		m.fn835(v0, v2)
		t2 := int32(load32(m.memory[uint32(v0):]))
		if t2 != i32(-1) {
			return
		}
	}
l2:
	store32(m.memory[uint32(v1):], uint32(i32(0)))
	return
}
func (m *Module) fn843(v0, v1 int32) {
	var v2, v3, v4, v5, v6, v7, v8, v9 int32
	var v10 int64
	t0 := m.g0
	v2 = t0 - i32(64)
	m.g0 = v2
	v3 = v1 + i32(8)
	t1 := int32(load32(m.memory[int64(uint32(v1))+28:]))
	v4 = t1
	t2 := int32(load32(m.memory[int64(uint32(v1))+24:]))
	v5 = t2
	t3 := int32(load32(m.memory[uint32(v1):]))
	v6 = t3
	t4 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	v7 = t4
	{
	l1:
		{
			v8 = v6
			if v8 == v7 {
				goto l0
			}
			t5 := v1
			v6 = v8 + i32(44)
			store32(m.memory[uint32(t5):], uint32(v6))
			t6 := int32(load32(m.memory[uint32(v8):]))
			if t6 == i32(-1) {
				goto l1
			}
			t7 := m.fn844(v3, v8)
			if t7 == 0 {
				goto l1
			}
			t8 := int32(load32(m.memory[uint32(v8+i32(16)):]))
			t9 := int32(load32(m.memory[uint32(v8+i32(20)):]))
			m.fn845(v2, t8, t9, i32(1073159), i32(67), i32(1073226), i32(2))
			t10 := int32(load32(m.memory[int64(uint32(v2))+4:]))
			v8 = t10
			t11 := int32(load32(m.memory[uint32(v2):]))
			v9 = t11
			if v9 == 0 {
				goto l1
			}
			t12 := m.fn846(v5, v9, v8)
			v8 = t12
			if v8 == 0 {
				goto l1
			}
			t13 := int32(m.memory[int64(uint32(v8))+24])
			if t13 != 0 {
				goto l1
			}
			t14 := int32(load32(m.memory[int64(uint32(v4))+4:]))
			t15 := int32(load32(m.memory[int64(uint32(v4))+8:]))
			t16 := int32(load32(m.memory[int64(uint32(v8))+4:]))
			t17 := int32(load32(m.memory[int64(uint32(v8))+8:]))
			m.fn774(v2+i32(36), t14, t15, t16, t17)
			m.fn780(v2+i32(12), v2+i32(36))
			t18 := int32(load32(m.memory[int64(uint32(v2))+12:]))
			v8 = t18
			if v8 == i32(-1) {
				goto l1
			}
		}
		t19 := int64(load64(m.memory[int64(uint32(v2))+16:]))
		v10 = t19
		t20 := int32(load32(m.memory[int64(uint32(v2))+24:]))
		t21 := int32(load32(m.memory[int64(uint32(v2))+28:]))
		m.fn134(t20, t21)
		store64(m.memory[int64(uint32(v0))+4:], uint64(v10))
		store32(m.memory[uint32(v0):], uint32(v8))
		goto l2
	}
l0:
	store32(m.memory[uint32(v0):], uint32(i32(-1)))
l2:
	m.g0 = v2 + i32(64)
}
func (m *Module) fn844(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t2 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	t3 := int32(load32(m.memory[int64(uint32(v0))+12:]))
	t4 := m.fn847(v1, t0, t1, t2, t3)
	return t4
}
func (m *Module) fn845(v0, v1, v2, v3, v4, v5, v6 int32) {
	var v7 int32
	v2 = v2 << 5
	{
	l4:
		if v2 != 0 {
			goto l0
		}
		v1 = i32(0)
		goto l1
	l0:
		{
			t0 := int32(load32(m.memory[uint32(v1+i32(4)):]))
			t1 := int32(load32(m.memory[uint32(v1+i32(8)):]))
			t2 := m.fn773(t0, t1, v5, v6)
			if t2 == 0 {
				goto l2
			}
			t3 := int32(load32(m.memory[uint32(v1+i32(24)):]))
			v7 = t3
			if v7 == 0 {
				goto l2
			}
			t4 := int32(load32(m.memory[uint32(v1+i32(28)):]))
			t5 := m.fn15(v7+i32(8), t4, v3, v4)
			if t5 != 0 {
				goto l3
			}
		}
	l2:
		v1 = v1 + i32(32)
		v2 = v2 + i32(-32)
		goto l4
	l3:
		t6 := int32(load32(m.memory[int64(uint32(v1))+20:]))
		v2 = t6
		t7 := int32(load32(m.memory[int64(uint32(v1))+16:]))
		v1 = t7
	}
l1:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v2))
	store32(m.memory[uint32(v0):], uint32(v1))
}
func (m *Module) fn846(v0, v1, v2 int32) int32 {
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
						v11 = t11 + (i32(0)-v10)*i32(40)
						t12 := int32(load32(m.memory[uint32(v11+i32(-36)):]))
						t13 := int32(load32(m.memory[uint32(v11+i32(-32)):]))
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
					p15 := v0 + (i32(0)-v10)*i32(40)
					if v9 != 0 {
						p15 = i32(0)
					}
					p16 := p15 + i32(-28)
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
func (m *Module) fn847(v0, v1, v2, v3, v4 int32) int32 {
	var v5 int32
	v5 = i32(0)
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		t1 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		t2 := m.fn773(t0, t1, v3, v4)
		if t2 == 0 {
			goto l0
		}
		t3 := int32(load32(m.memory[int64(uint32(v0))+36:]))
		v5 = t3
		p4 := i32(0)
		if v5 != 0 {
			p4 = v5 + i32(8)
		}
		t5 := int32(load32(m.memory[int64(uint32(v0))+40:]))
		t6 := m.fn848(p4, t5, v1, v2)
		v5 = t6
	}
l0:
	return v5
}
func (m *Module) fn848(v0, v1, v2, v3 int32) int32 {
	var v4 int32
	var p0 int32
	if v0|v2 == 0 {
		p0 = 1
	}
	v4 = p0
	{
		if v0 == 0 {
			goto l0
		}
		if v2 == 0 {
			goto l0
		}
		t1 := m.fn15(v0, v1, v2, v3)
		v4 = t1
	}
l0:
	return v4
}
func (m *Module) fn849(v0, v1, v2 int32) {
	var v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13 int32
	var v14, v15, v16 int64
	var v17, v18 int32
	t0 := m.g0
	v3 = t0 - i32(288)
	m.g0 = v3
	{
		{
			{
				t1 := int32(load32(m.memory[uint32(v1):]))
				v4 = t1
				switch v4 >> 31 & (v4 + i32(-0x7fffffff)) {
				case 3:
					goto l3
				case 4:
					t85 := int32(load32(m.memory[int64(uint32(v1))+8:]))
					t86 := int32(load32(m.memory[int64(uint32(v1))+12:]))
					m.fn65(v3+i32(120), t85, t86, v2)
					{
						t87 := int32(load32(m.memory[int64(uint32(v3))+128:]))
						v8 = t87
						if v8 == 0 {
							store32(m.memory[uint32(v0):], uint32(i32(-1)))
							t105 := int32(load32(m.memory[int64(uint32(v3))+120:]))
							t106 := int32(load32(m.memory[int64(uint32(v3))+124:]))
							m.fn16(t105, t106)
							goto l35
						}
						t88 := int32(load32(m.memory[int64(uint32(v3))+124:]))
						t89 := v3 + i32(160)
						v4 = t88
						m.fn70(t89, v4, v8)
						m.fn854(v3+i32(216), v3+i32(160))
						{
							t90 := int32(load32(m.memory[int64(uint32(v3))+216:]))
							if t90 == i32(-1) {
								goto l41
							}
							m.fn59(v3+i32(80), i32(4), i32(4), i32(12))
							t91 := int32(load32(m.memory[int64(uint32(v3))+80:]))
							v4 = t91
							t92 := int32(load32(m.memory[int64(uint32(v3))+84:]))
							v11 = t92
							t93 := int32(load32(m.memory[int64(uint32(v3))+224:]))
							store32(m.memory[int64(uint32(v11))+8:], uint32(t93))
							t94 := int64(load64(m.memory[int64(uint32(v3))+216:]))
							store64(m.memory[uint32(v11):], uint64(t94))
							store32(m.memory[int64(uint32(v3))+212:], uint32(i32(1)))
							store32(m.memory[int64(uint32(v3))+208:], uint32(v11))
							store32(m.memory[int64(uint32(v3))+204:], uint32(v4))
							memory_copy(m.memory, uint32(v3+i32(232)), uint32(v3+i32(160)), uint32(i32(40)))
							v4 = i32(12)
							v8 = i32(1)
						l44:
							{
								m.fn854(v3+i32(272), v3+i32(232))
								t95 := int32(load32(m.memory[int64(uint32(v3))+272:]))
								if t95 == i32(-1) {
									t101 := int32(load32(m.memory[int64(uint32(v3))+124:]))
									v4 = t101
									t102 := int32(load32(m.memory[int64(uint32(v3))+212:]))
									v8 = t102
									t103 := int32(load32(m.memory[int64(uint32(v3))+208:]))
									v9 = t103
									goto l45
								}
								{
									t96 := int32(load32(m.memory[int64(uint32(v3))+204:]))
									if v8 != t96 {
										goto l43
									}
									m.fn60(v3+i32(204), i32(1))
									t97 := int32(load32(m.memory[int64(uint32(v3))+208:]))
									v11 = t97
								}
							l43:
								v9 = v11 + v4
								t98 := int32(load32(m.memory[int64(uint32(v3))+280:]))
								store32(m.memory[int64(uint32(v9))+8:], uint32(t98))
								t99 := int64(load64(m.memory[int64(uint32(v3))+272:]))
								store64(m.memory[uint32(v9):], uint64(t99))
								t100 := v3
								v8 = v8 + i32(1)
								store32(m.memory[int64(uint32(t100))+212:], uint32(v8))
								v4 = v4 + i32(12)
								goto l44
							}
						}
					l41:
						v8 = i32(0)
						store32(m.memory[int64(uint32(v3))+212:], uint32(i32(0)))
						store64(m.memory[int64(uint32(v3))+204:], uint64(i64(0x400000000)))
						v9 = i32(4)
					l45:
						m.fn77(v0, v9, v8, i32(1108166), i32(1))
						m.fn78(v3 + i32(204))
						t104 := int32(load32(m.memory[int64(uint32(v3))+120:]))
						m.fn16(t104, v4)
						goto l35
					}
				case 5:
					t125 := int32(load32(m.memory[int64(uint32(v1))+8:]))
					t126 := v3 + i32(160)
					v8 = t125
					t127 := int32(load32(m.memory[int64(uint32(v1))+12:]))
					t128 := v8
					v9 = t127
					m.fn805(t126, t128, v9, i32(3))
					m.fn855(v3+i32(96), v1+i32(16))
					t129 := int32(load32(m.memory[int64(uint32(v3))+100:]))
					t130 := int32(load32(m.memory[int64(uint32(v3))+96:]))
					t131 := v3
					v4 = t130
					p132 := i32(0)
					if v4 != 0 {
						p132 = t129
					}
					store32(m.memory[int64(uint32(t131))+220:], uint32(p132))
					t134 := v3
					p133 := i32(1)
					if v4 != 0 {
						p133 = v4
					}
					store32(m.memory[int64(uint32(t134))+216:], uint32(p133))
					m.fn856(v3+i32(88), v8, v9, i32(10))
					t135 := int64(load64(m.memory[int64(uint32(v3))+88:]))
					store64(m.memory[int64(uint32(v3))+272:], uint64(t135))
					store32(m.memory[int64(uint32(v3))+252:], uint32(i32(1)))
					store32(m.memory[int64(uint32(v3))+244:], uint32(i32(1)))
					store32(m.memory[int64(uint32(v3))+236:], uint32(i32(25)))
					store32(m.memory[int64(uint32(v3))+248:], uint32(v3+i32(272)))
					store32(m.memory[int64(uint32(v3))+240:], uint32(v3+i32(216)))
					store32(m.memory[int64(uint32(v3))+232:], uint32(v3+i32(160)))
					m.fn73(v0, i32(1080548), v3+i32(232))
					t136 := int32(load32(m.memory[int64(uint32(v3))+160:]))
					t137 := int32(load32(m.memory[int64(uint32(v3))+164:]))
					m.fn16(t136, t137)
					goto l35
				case 6:
					m.fn51(v0, i32(1080559), i32(3))
					goto l35
				default:
					t109 := int32(load32(m.memory[int64(uint32(v1))+4:]))
					t110 := int32(load32(m.memory[int64(uint32(v1))+8:]))
					m.fn798(v3+i32(272), t109, t110, i32(1), v2)
					t111 := int32(load32(m.memory[int64(uint32(v3))+276:]))
					t112 := v3
					v4 = t111
					t113 := int32(load32(m.memory[int64(uint32(v3))+280:]))
					m.fn46(t112, v4, t113)
					t114 := int32(load32(m.memory[int64(uint32(v3))+4:]))
					t115 := v3
					v8 = t114
					store32(m.memory[int64(uint32(t115))+220:], uint32(v8))
					t116 := int32(load32(m.memory[uint32(v3):]))
					store32(m.memory[int64(uint32(v3))+216:], uint32(t116))
					{
						if v8 == 0 {
							store32(m.memory[uint32(v0):], uint32(i32(-1)))
							t124 := int32(load32(m.memory[int64(uint32(v3))+272:]))
							m.fn16(t124, v4)
							goto l35
						}
						t117 := int32(m.memory[int64(uint32(v1))+24])
						t118 := v3 + i32(160)
						v8 = t117
						p119 := i32(6)
						if uint32(v8) < uint32(i32(6)) {
							p119 = v8
						}
						p120 := i32(1)
						if v8 != 0 {
							p120 = p119
						}
						m.fn808(t118, i32(1072920), p120)
						store32(m.memory[int64(uint32(v3))+244:], uint32(i32(1)))
						store32(m.memory[int64(uint32(v3))+236:], uint32(i32(25)))
						store32(m.memory[int64(uint32(v3))+240:], uint32(v3+i32(216)))
						store32(m.memory[int64(uint32(v3))+232:], uint32(v3+i32(160)))
						m.fn73(v0, i32(1052689), v3+i32(232))
						t121 := int32(load32(m.memory[int64(uint32(v3))+160:]))
						t122 := int32(load32(m.memory[int64(uint32(v3))+164:]))
						m.fn16(t121, t122)
						t123 := int32(load32(m.memory[int64(uint32(v3))+272:]))
						m.fn16(t123, v4)
						goto l35
					}
				case 1:
					v5 = i32(0)
					t2 := int32(load32(m.memory[int64(uint32(v1))+8:]))
					t3 := int32(load32(m.memory[int64(uint32(v1))+12:]))
					m.fn798(v3+i32(216), t2, t3, i32(0), v2)
					t4 := int32(load32(m.memory[int64(uint32(v3))+220:]))
					t5 := int32(load32(m.memory[int64(uint32(v3))+224:]))
					m.fn70(v3+i32(160), t4, t5)
					m.fn850(v3+i32(40), v3+i32(160))
					{
						{
							t6 := int32(load32(m.memory[int64(uint32(v3))+40:]))
							v4 = t6
							if v4 != 0 {
								goto l7
							}
							v6 = i32(4)
							v4 = i32(0)
							v7 = i32(0)
							goto l8
						}
					l7:
						t7 := int32(load32(m.memory[int64(uint32(v3))+44:]))
						v8 = t7
						m.fn59(v3+i32(32), i32(4), i32(4), i32(8))
						t8 := int32(load32(m.memory[int64(uint32(v3))+32:]))
						v9 = t8
						t9 := int32(load32(m.memory[int64(uint32(v3))+36:]))
						v10 = t9
						store32(m.memory[int64(uint32(v10))+4:], uint32(v8))
						store32(m.memory[uint32(v10):], uint32(v4))
						store32(m.memory[int64(uint32(v3))+280:], uint32(i32(1)))
						store32(m.memory[int64(uint32(v3))+276:], uint32(v10))
						store32(m.memory[int64(uint32(v3))+272:], uint32(v9))
						memory_copy(m.memory, uint32(v3+i32(232)), uint32(v3+i32(160)), uint32(i32(40)))
						v8 = i32(12)
						v4 = i32(1)
					l11:
						{
							m.fn850(v3+i32(24), v3+i32(232))
							t10 := int32(load32(m.memory[int64(uint32(v3))+24:]))
							v9 = t10
							if v9 == 0 {
								goto l9
							}
							t11 := int32(load32(m.memory[int64(uint32(v3))+28:]))
							v11 = t11
							{
								t12 := int32(load32(m.memory[int64(uint32(v3))+272:]))
								if v4 != t12 {
									goto l10
								}
								m.fn797(v3 + i32(272))
								t13 := int32(load32(m.memory[int64(uint32(v3))+276:]))
								v10 = t13
							}
						l10:
							v2 = v10 + v8
							store32(m.memory[uint32(v2):], uint32(v11))
							store32(m.memory[uint32(v2+i32(-4)):], uint32(v9))
							t14 := v3
							v4 = v4 + i32(1)
							store32(m.memory[int64(uint32(t14))+280:], uint32(v4))
							v8 = v8 + i32(8)
							goto l11
						}
					l9:
						t15 := int32(load32(m.memory[int64(uint32(v3))+276:]))
						v6 = t15
						t16 := int32(load32(m.memory[int64(uint32(v3))+272:]))
						v7 = t16
					}
				l8:
					t17 := v6
					v9 = v4 << 3
					v11 = t17 + v9
					v8 = v6
				l13:
					{
						if v9 == 0 {
							goto l12
						}
						t18 := int32(load32(m.memory[int64(uint32(v8))+4:]))
						if t18 != 0 {
							goto l12
						}
						v9 = v9 + i32(-8)
						v5 = v5 + i32(1)
						v8 = v8 + i32(8)
						goto l13
					}
				l12:
					v12 = i32(0)
					v8 = i32(0) - v4<<3
					v2 = v4
					{
					l16:
						{
							v10 = v2
							v1 = i32(1)
							if v8 != 0 {
								goto l14
							}
							v8 = i32(0)
							goto l15
						l14:
							v8 = v8 + i32(8)
							v2 = v10 + i32(-1)
							v13 = v11 + i32(-4)
							v11 = v11 + i32(-8)
							t19 := int32(load32(m.memory[uint32(v13):]))
							if t19 == 0 {
								goto l16
							}
						}
						v8 = i32(0)
						if v9 == 0 {
							goto l15
						}
						if uint32(v10) < uint32(v5) {
							m.fn151(v5, v10, v4, i32(1080564))
							panic("unreachable")
						}
						m.fn632(v3+i32(232), v6+v5<<3, v10-v5, i32(1108166), i32(1))
						{
							t20 := int32(load32(m.memory[int64(uint32(v3))+236:]))
							v1 = t20
							t21 := int32(load32(m.memory[int64(uint32(v3))+240:]))
							t22 := v1
							v8 = t21
							t23 := m.fn851(t22, v8)
							if t23 == 0 {
								goto l18
							}
							store32(m.memory[int64(uint32(v3))+160:], uint32(v1))
							store32(m.memory[int64(uint32(v3))+164:], uint32(v1+v8))
							m.fn577(v3+i32(16), v3+i32(160))
							{
								t24 := int32(load32(m.memory[int64(uint32(v3))+16:]))
								if t24 != i32(1) {
									goto l19
								}
								{
									{
										t25 := int32(load32(m.memory[int64(uint32(v3))+20:]))
										v4 = t25
										if uint32(v4) >= uint32(i32(128)) {
											goto l20
										}
										v4 = i32(-1)
										goto l21
									}
								l20:
									if uint32(v4) >= uint32(i32(2048)) {
										goto l22
									}
									v4 = i32(-2)
									goto l21
								l22:
									p26 := i32(-4)
									if uint32(v4) < uint32(i32(65536)) {
										p26 = i32(-3)
									}
									v4 = p26
								}
							l21:
								t27 := v3
								v8 = v4 + v8
								store32(m.memory[int64(uint32(t27))+240:], uint32(v8))
							}
						l19:
							m.fn628(v3+i32(8), v1, v8)
							t28 := int32(load32(m.memory[int64(uint32(v3))+12:]))
							m.fn852(v3+i32(232), t28, i32(1080580))
							t29 := int32(load32(m.memory[int64(uint32(v3))+240:]))
							v8 = t29
							t30 := int32(load32(m.memory[int64(uint32(v3))+236:]))
							v1 = t30
						}
					l18:
						t31 := int32(load32(m.memory[int64(uint32(v3))+232:]))
						v12 = t31
					}
				l15:
					m.fn639(v7, v6)
					if v8 == 0 {
						goto l23
					}
					store32(m.memory[int64(uint32(v0))+8:], uint32(v8))
					store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
					store32(m.memory[uint32(v0):], uint32(v12))
					goto l24
				case 2:
					t32 := int32(load32(m.memory[int64(uint32(v1))+24:]))
					v8 = t32
					if v8 == 0 {
						store32(m.memory[uint32(v0):], uint32(i32(-1)))
						goto l35
					}
					store32(m.memory[int64(uint32(v3))+116:], uint32(i32(0)))
					store64(m.memory[int64(uint32(v3))+108:], uint64(i64(0x400000000)))
					t33 := int32(load32(m.memory[int64(uint32(v1))+20:]))
					v4 = t33
					v6 = v4 + v8*i32(28)
					t34 := int64(load64(m.memory[int64(uint32(v1))+8:]))
					v14 = t34
					t35 := int32(m.memory[int64(uint32(v1))+28])
					v1 = t35
					v15 = i64(0)
					v10 = i32(0)
				l39:
					{
						if v4 == v6 {
							t55 := int32(load32(m.memory[int64(uint32(v3))+112:]))
							t56 := int32(load32(m.memory[int64(uint32(v3))+116:]))
							t57 := v0
							v4 = v10 & i32(1)
							p58 := i32(1108166)
							if v4 != 0 {
								p58 = i32(1080488)
							}
							p59 := i32(1)
							if v4 != 0 {
								p59 = i32(2)
							}
							m.fn77(t57, t55, t56, p58, p59)
							m.fn78(v3 + i32(108))
							goto l35
						}
						{
							t36 := int32(load32(m.memory[int64(uint32(v4))+12:]))
							if t36 == i32(-1) {
								goto l27
							}
							t37 := int32(load32(m.memory[int64(uint32(v4))+16:]))
							t38 := int32(load32(m.memory[int64(uint32(v4))+20:]))
							m.fn799(v3+i32(232), t37, t38, i32(0))
							store32(m.memory[int64(uint32(v3))+276:], uint32(i32(25)))
							store32(m.memory[int64(uint32(v3))+272:], uint32(v3+i32(232)))
							m.fn73(v3+i32(160), i32(1070102), v3+i32(272))
							t39 := int32(load32(m.memory[int64(uint32(v3))+232:]))
							t40 := int32(load32(m.memory[int64(uint32(v3))+236:]))
							m.fn16(t39, t40)
							t41 := int64(load64(m.memory[int64(uint32(v3))+160:]))
							store64(m.memory[int64(uint32(v3))+120:], uint64(t41))
							t42 := int32(load32(m.memory[int64(uint32(v3))+168:]))
							store32(m.memory[int64(uint32(v3))+128:], uint32(t42))
							goto l28
						}
					l27:
						switch v1 {
						default:
							t43 := v3 + i32(232)
							t44 := v1
							v16 = v14 + v15
							p45 := v16
							if uint64(v16) < uint64(v14) {
								p45 = i64(-1)
							}
							m.fn800(t43, t44, p45)
							store32(m.memory[int64(uint32(v3))+276:], uint32(i32(25)))
							store32(m.memory[int64(uint32(v3))+272:], uint32(v3+i32(232)))
							m.fn73(v3+i32(160), i32(1070102), v3+i32(272))
							t46 := int32(load32(m.memory[int64(uint32(v3))+232:]))
							t47 := int32(load32(m.memory[int64(uint32(v3))+236:]))
							m.fn16(t46, t47)
							t48 := int64(load64(m.memory[int64(uint32(v3))+160:]))
							store64(m.memory[int64(uint32(v3))+120:], uint64(t48))
							t49 := int32(load32(m.memory[int64(uint32(v3))+168:]))
							store32(m.memory[int64(uint32(v3))+128:], uint32(t49))
							goto l28
						case 0:
							m.fn51(v3+i32(120), i32(1080490), i32(2))
							goto l28
						case 1:
							t50 := v3
							v16 = v14 + v15
							p51 := v16
							if uint64(v16) < uint64(v14) {
								p51 = i64(-1)
							}
							store64(m.memory[int64(uint32(t50))+272:], uint64(p51))
							store32(m.memory[int64(uint32(v3))+164:], uint32(i32(28)))
							store32(m.memory[int64(uint32(v3))+160:], uint32(v3+i32(272)))
							m.fn73(v3+i32(232), i32(1070097), v3+i32(160))
							t52 := int64(load64(m.memory[int64(uint32(v3))+232:]))
							store64(m.memory[int64(uint32(v3))+120:], uint64(t52))
							t53 := int32(load32(m.memory[int64(uint32(v3))+240:]))
							store32(m.memory[int64(uint32(v3))+128:], uint32(t53))
						}
					l28:
						v9 = i32(4)
						v8 = i32(1080492)
						{
							t54 := int32(m.memory[int64(uint32(v4))+24])
							switch t54 {
							case 0:
								goto l32
							default:
								goto l33
							case 2:
								v9 = i32(0)
								v8 = i32(1)
								goto l32
							}
						}
					l33:
						v8 = i32(1080496)
					l32:
						v15 = v15 + i64(1)
						v5 = v4 + i32(28)
						store32(m.memory[int64(uint32(v3))+136:], uint32(v9))
						store32(m.memory[int64(uint32(v3))+132:], uint32(v8))
						t60 := int32(load32(m.memory[int64(uint32(v4))+4:]))
						t61 := int32(load32(m.memory[int64(uint32(v4))+8:]))
						m.fn65(v3+i32(216), t60, t61, v2)
						t62 := int32(load32(m.memory[int64(uint32(v4))+8:]))
						v8 = t62
						t63 := int32(load32(m.memory[int64(uint32(v3))+124:]))
						t64 := v3 + i32(272)
						v4 = t63
						t65 := int32(load32(m.memory[int64(uint32(v3))+128:]))
						t66 := m.fn853(v4, v4+t65)
						m.fn808(t64, i32(1097368), t66)
						t67 := int32(load32(m.memory[int64(uint32(v3))+220:]))
						t68 := v3 + i32(160)
						v13 = t67
						t69 := int32(load32(m.memory[int64(uint32(v3))+224:]))
						m.fn70(t68, v13, t69)
						m.fn71(v3+i32(56), v3+i32(160))
						t70 := int32(load32(m.memory[int64(uint32(v3))+60:]))
						t71 := int32(load32(m.memory[int64(uint32(v3))+56:]))
						t72 := v3
						v4 = t71
						p73 := i32(0)
						if v4 != 0 {
							p73 = t70
						}
						store32(m.memory[int64(uint32(t72))+144:], uint32(p73))
						t75 := v3
						p74 := i32(1)
						if v4 != 0 {
							p74 = v4
						}
						store32(m.memory[int64(uint32(t75))+140:], uint32(p74))
						store32(m.memory[int64(uint32(v3))+252:], uint32(i32(1)))
						store32(m.memory[int64(uint32(v3))+244:], uint32(i32(1)))
						store32(m.memory[int64(uint32(v3))+236:], uint32(i32(25)))
						store32(m.memory[int64(uint32(v3))+248:], uint32(v3+i32(140)))
						store32(m.memory[int64(uint32(v3))+240:], uint32(v3+i32(132)))
						store32(m.memory[int64(uint32(v3))+232:], uint32(v3+i32(120)))
						m.fn73(v3+i32(148), i32(0x10004d), v3+i32(232))
						memory_copy(m.memory, uint32(v3+i32(232)), uint32(v3+i32(160)), uint32(i32(40)))
						var p76 int32
						if uint32(v8) > uint32(i32(1)) {
							p76 = 1
						}
						v10 = p76 | v10
						t77 := int32(load32(m.memory[int64(uint32(v3))+280:]))
						v11 = t77
						t78 := int32(load32(m.memory[int64(uint32(v3))+276:]))
						v9 = t78
					l38:
						m.fn71(v3+i32(48), v3+i32(232))
						{
							t79 := int32(load32(m.memory[int64(uint32(v3))+48:]))
							v8 = t79
							if v8 == 0 {
								m.fn33(v3+i32(108), v3+i32(148))
								t81 := int32(load32(m.memory[int64(uint32(v3))+272:]))
								m.fn16(t81, v9)
								t82 := int32(load32(m.memory[int64(uint32(v3))+216:]))
								m.fn16(t82, v13)
								t83 := int32(load32(m.memory[int64(uint32(v3))+120:]))
								t84 := int32(load32(m.memory[int64(uint32(v3))+124:]))
								m.fn16(t83, t84)
								v4 = v5
								goto l39
							}
							t80 := int32(load32(m.memory[int64(uint32(v3))+52:]))
							v4 = t80
							m.fn74(v3+i32(148), i32(10))
							if v4 != 0 {
								m.fn75(v3+i32(148), v9, v11)
								m.fn75(v3+i32(148), v8, v4)
								goto l38
							}
							v10 = i32(1)
							goto l38
						}
					}
				}
			}
		l23:
			store32(m.memory[uint32(v0):], uint32(i32(-1)))
			m.fn16(v12, v1)
		l24:
			t107 := int32(load32(m.memory[int64(uint32(v3))+216:]))
			t108 := int32(load32(m.memory[int64(uint32(v3))+220:]))
			m.fn16(t107, t108)
			goto l35
		}
	l3:
		{
			t138 := int32(m.memory[int64(uint32(v1))+20])
			if t138 != 0 {
				t140 := int32(load32(m.memory[int64(uint32(v1))+12:]))
				v12 = t140
				if v12 != i32(1) {
					goto l48
				}
				{
					t141 := int32(load32(m.memory[int64(uint32(v1))+8:]))
					v4 = t141
					t142 := int32(load32(m.memory[uint32(v4+i32(8)):]))
					if t142 != i32(1) {
						goto l49
					}
					t143 := int32(load32(m.memory[uint32(v4+i32(4)):]))
					t144 := m.fn412(t143, i32(1), i32(1079844))
					t145 := int32(load32(m.memory[uint32(t144):]))
					if t145 != i32(-1) {
						t146 := m.fn857(v4, i32(1), i32(0), i32(1080500))
						v4 = t146
						t147 := int32(load32(m.memory[uint32(v4+i32(4)):]))
						t148 := int32(load32(m.memory[uint32(v4+i32(8)):]))
						t149 := m.fn412(t147, t148, i32(1080516))
						v4 = t149
						t150 := int32(load32(m.memory[uint32(v4):]))
						if t150 == i32(-1) {
							m.fn256(i32(1286542), i32(40), i32(1080532))
							panic("unreachable")
						}
						t151 := int32(load32(m.memory[int64(uint32(v4))+4:]))
						t152 := int32(load32(m.memory[int64(uint32(v4))+8:]))
						m.fn65(v3+i32(232), t151, t152, v2)
						{
							t153 := int32(load32(m.memory[int64(uint32(v3))+240:]))
							if t153 == 0 {
								store32(m.memory[uint32(v0):], uint32(i32(-1)))
								t156 := int32(load32(m.memory[int64(uint32(v3))+232:]))
								t157 := int32(load32(m.memory[int64(uint32(v3))+236:]))
								m.fn16(t156, t157)
								goto l35
							}
							t154 := int32(load32(m.memory[int64(uint32(v3))+240:]))
							store32(m.memory[int64(uint32(v0))+8:], uint32(t154))
							t155 := int64(load64(m.memory[int64(uint32(v3))+232:]))
							store64(m.memory[uint32(v0):], uint64(t155))
							goto l35
						}
					}
				}
			l49:
				v12 = i32(1)
				goto l51
			}
			t139 := int32(load32(m.memory[int64(uint32(v1))+12:]))
			v12 = t139
			goto l48
		}
	l48:
		if v12 != 0 {
			goto l51
		}
		store32(m.memory[uint32(v0):], uint32(i32(-1)))
		goto l35
	l51:
		t158 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		v7 = t158
		t159 := int32(load32(m.memory[int64(uint32(v7))+8:]))
		v10 = t159
		{
			if v12 == i32(1) {
				goto l54
			}
			v4 = v7 + i32(20)
			t160 := int32(uint32(v12*i32(12)+i32(-12)) / uint32(i32(12)))
			v8 = t160
		l55:
			{
				t161 := int32(load32(m.memory[uint32(v4):]))
				t162 := v10
				v9 = t161
				p163 := v9
				if uint32(v10) > uint32(v9) {
					p163 = t162
				}
				v10 = p163
				v4 = v4 + i32(12)
				v8 = v8 + i32(-1)
				if v8 != 0 {
					goto l55
				}
			}
		}
	l54:
		m.fn59(v3+i32(72), v12, i32(4), i32(12))
		store32(m.memory[int64(uint32(v3))+280:], uint32(i32(0)))
		t164 := int32(load32(m.memory[int64(uint32(v3))+76:]))
		t165 := v3
		v17 = t164
		store32(m.memory[int64(uint32(t165))+276:], uint32(v17))
		t166 := int32(load32(m.memory[int64(uint32(v3))+72:]))
		t167 := v3
		v4 = t166
		store32(m.memory[int64(uint32(t167))+272:], uint32(v4))
		v5 = i32(0)
		{
			if uint32(v12) <= uint32(v4) {
				goto l56
			}
			m.fn62(v3+i32(272), i32(0), v12, i32(4), i32(12))
			t168 := int32(load32(m.memory[int64(uint32(v3))+280:]))
			v5 = t168
			t169 := int32(load32(m.memory[int64(uint32(v3))+276:]))
			v17 = t169
		}
	l56:
		v18 = i32(0) - v10
		v13 = i32(0)
	l64:
		{
			v4 = v7 + v13*i32(12)
			t170 := int32(load32(m.memory[int64(uint32(v4))+4:]))
			v8 = t170
			t171 := int32(load32(m.memory[int64(uint32(v4))+8:]))
			t172 := v3 + i32(64)
			v9 = t171
			m.fn59(t172, v9, i32(4), i32(16))
			store32(m.memory[int64(uint32(v3))+168:], uint32(i32(0)))
			t173 := int64(load64(m.memory[int64(uint32(v3))+64:]))
			store64(m.memory[int64(uint32(v3))+160:], uint64(t173))
			m.fn751(v3+i32(160), v9)
			t174 := int32(load32(m.memory[int64(uint32(v3))+168:]))
			v4 = t174
			{
				if v9 == 0 {
					goto l57
				}
				v6 = v9 + v4
				t175 := int32(load32(m.memory[int64(uint32(v3))+164:]))
				v4 = t175 + v4<<4
			l60:
				{
					{
						t176 := int32(load32(m.memory[uint32(v8):]))
						if t176 != i32(-1) {
							goto l58
						}
						store32(m.memory[int64(uint32(v3))+240:], uint32(i32(0)))
						store64(m.memory[int64(uint32(v3))+232:], uint64(i64(0x100000000)))
						v11 = i32(1)
						goto l59
					}
				l58:
					m.fn794(v3+i32(232), v8, v2)
					v11 = i32(0)
				l59:
					t177 := int64(load64(m.memory[int64(uint32(v3))+232:]))
					store64(m.memory[uint32(v4):], uint64(t177))
					m.memory[int64(uint32(v3))+244] = byte(v11)
					t178 := int64(load64(m.memory[int64(uint32(v3))+240:]))
					store64(m.memory[int64(uint32(v4))+8:], uint64(t178))
					v8 = v8 + i32(20)
					v4 = v4 + i32(16)
					v9 = v9 + i32(-1)
					if v9 != 0 {
						goto l60
					}
				}
				v4 = v6
			}
		l57:
			t179 := int64(load64(m.memory[int64(uint32(v3))+160:]))
			store64(m.memory[int64(uint32(v3))+232:], uint64(t179))
			store32(m.memory[int64(uint32(v3))+240:], uint32(v4))
			{
				if uint32(v10) > uint32(v4) {
					goto l61
				}
				m.fn750(v3+i32(232), v10)
				goto l62
			l61:
				m.fn751(v3+i32(232), v10-v4)
				v8 = v18 + v4
				t180 := int32(load32(m.memory[int64(uint32(v3))+240:]))
				t181 := v10
				v9 = t180
				v11 = t181 + v9 - v4
				t182 := int32(load32(m.memory[int64(uint32(v3))+236:]))
				v4 = t182 + v9<<4
			l63:
				store64(m.memory[uint32(v4):], uint64(i64(0x100000000)))
				m.memory[uint32(v4+i32(12))] = byte(i32(0))
				store32(m.memory[uint32(v4+i32(8)):], uint32(i32(0)))
				v4 = v4 + i32(16)
				v8 = v8 + i32(1)
				if v8 != 0 {
					goto l63
				}
				store32(m.memory[int64(uint32(v3))+240:], uint32(v11))
			}
		l62:
			t183 := int32(load32(m.memory[int64(uint32(v3))+240:]))
			t184 := v3
			v4 = t183
			store32(m.memory[int64(uint32(t184))+168:], uint32(v4))
			t185 := int64(load64(m.memory[int64(uint32(v3))+232:]))
			t186 := v3
			v15 = t185
			store64(m.memory[int64(uint32(t186))+160:], uint64(v15))
			v8 = v17 + v5*i32(12)
			store32(m.memory[int64(uint32(v8))+8:], uint32(v4))
			store64(m.memory[uint32(v8):], uint64(v15))
			v5 = v5 + i32(1)
			v13 = v13 + i32(1)
			if v13 != v12 {
				goto l64
			}
		}
		t187 := int64(load64(m.memory[int64(uint32(v3))+272:]))
		store64(m.memory[int64(uint32(v3))+216:], uint64(t187))
		store32(m.memory[int64(uint32(v3))+224:], uint32(v5))
		t188 := int32(load32(m.memory[int64(uint32(v3))+220:]))
		v2 = t188
	l75:
		{
			{
				if uint32(v5) < uint32(i32(2)) {
					goto l65
				}
				v4 = v2 + v5*i32(12)
				if v4 == i32(12) {
					goto l66
				}
				t189 := int32(load32(m.memory[uint32(v4+i32(-4)):]))
				v8 = t189 << 4
				t190 := int32(load32(m.memory[uint32(v4+i32(-8)):]))
				v4 = t190
			l68:
				{
					if v8 == 0 {
						t202 := v3
						t203 := v2
						v5 = v5 + i32(-1)
						v4 = t203 + v5*i32(12)
						t204 := int32(load32(m.memory[int64(uint32(v4))+8:]))
						store32(m.memory[int64(uint32(t202))+240:], uint32(t204))
						t205 := int64(load64(m.memory[uint32(v4):]))
						t206 := v3
						v15 = t205
						store64(m.memory[int64(uint32(t206))+232:], uint64(v15))
						if int32(v15) == i32(-1) {
							goto l75
						}
						m.fn859(v3 + i32(232))
						goto l75
					}
					t191 := int32(load32(m.memory[int64(uint32(v4))+8:]))
					if t191 != 0 {
						goto l65
					}
					v8 = v8 + i32(-16)
					t192 := int32(m.memory[int64(uint32(v4))+12])
					v9 = t192
					v4 = v4 + i32(16)
					if v9&i32(1) == 0 {
						goto l68
					}
				}
			}
		l65:
			store32(m.memory[int64(uint32(v3))+224:], uint32(v5))
			if v5 != 0 {
				goto l69
			}
			goto l70
		l66:
			store32(m.memory[int64(uint32(v3))+224:], uint32(v5))
		l69:
			t193 := int32(load32(m.memory[uint32(v2+i32(4)):]))
			t194 := int32(load32(m.memory[uint32(v2+i32(8)):]))
			t195 := m.fn858(t193, t194)
			v4 = t195
			{
				if v5 == i32(1) {
					goto l71
				}
				v8 = v2 + i32(20)
				t196 := int32(uint32(v5*i32(12)+i32(-12)) / uint32(i32(12)))
				v9 = t196
			l72:
				{
					t197 := int32(load32(m.memory[uint32(v8+i32(-4)):]))
					t198 := int32(load32(m.memory[uint32(v8):]))
					t199 := m.fn858(t197, t198)
					t200 := v4
					v11 = t199
					p201 := v11
					if uint32(v4) > uint32(v11) {
						p201 = t200
					}
					v4 = p201
					v8 = v8 + i32(12)
					v9 = v9 + i32(-1)
					if v9 != 0 {
						goto l72
					}
				}
			}
		l71:
			if v4 == 0 {
				goto l70
			}
			v11 = v2 + i32(12)
			v9 = v5 * i32(12)
			v8 = i32(0)
		l74:
			if v9 == v8 {
				store32(m.memory[int64(uint32(v3))+280:], uint32(i32(0)))
				store64(m.memory[int64(uint32(v3))+272:], uint64(i64(0x100000000)))
				{
					{
						t207 := int32(load32(m.memory[int64(uint32(v1))+16:]))
						if t207 != 0 {
							goto l76
						}
						store32(m.memory[int64(uint32(v3))+240:], uint32(i32(0)))
						store64(m.memory[int64(uint32(v3))+232:], uint64(i64(0x100000000)))
						m.fn189(v3+i32(160), v3+i32(232), v4)
						t208 := int32(load32(m.memory[int64(uint32(v3))+168:]))
						v9 = t208
						t209 := int32(load32(m.memory[int64(uint32(v3))+164:]))
						v8 = t209
						goto l77
					}
				l76:
					t210 := int32(load32(m.memory[uint32(v2+i32(8)):]))
					v8 = t210
					t211 := int32(load32(m.memory[uint32(v2+i32(4)):]))
					v10 = t211
					t212 := int32(load32(m.memory[uint32(v2):]))
					v6 = t212
					v9 = v5*i32(12) + i32(-12)
					if v9 == 0 {
						goto l78
					}
					memory_copy(m.memory, uint32(v2), uint32(v11), uint32(v9))
				l78:
					t213 := v3
					v5 = v5 + i32(-1)
					store32(m.memory[int64(uint32(t213))+224:], uint32(v5))
					if v6 == i32(-1) {
						m.fn98(v5, i32(1084656))
						panic("unreachable")
					}
					v1 = v8 << 4
					v9 = i32(0)
					v8 = i32(0)
				l81:
					{
						v11 = v10 + v9
						if v1 == v8 {
							goto l80
						}
						v13 = v10 + v8
						t214 := int64(load64(m.memory[uint32(v13):]))
						v15 = t214
						t215 := int32(load32(m.memory[int64(uint32(v13))+8:]))
						store32(m.memory[int64(uint32(v11))+8:], uint32(t215))
						store64(m.memory[uint32(v11):], uint64(v15))
						v9 = v9 + i32(12)
						v8 = v8 + i32(16)
						goto l81
					}
				l80:
					m.fn419(i32(0), i32(4))
					v13 = v6 << 4
					t216 := int32(uint32(v13) / uint32(i32(12)))
					v9 = t216
					v8 = v10
					{
						if v6 == 0 {
							goto l82
						}
						v8 = v10
						t217 := v13
						v1 = v9 * i32(12)
						if t217 == v1 {
							goto l82
						}
						t218 := m.fn392(v10, v13, v1)
						v8 = t218
						if v8 == 0 {
							m.fn85(i32(4), v1)
							panic("unreachable")
						}
					}
				l82:
					store32(m.memory[int64(uint32(v3))+164:], uint32(v8))
					store32(m.memory[int64(uint32(v3))+160:], uint32(v9))
					t219 := int32(uint32(v11-v10) / uint32(i32(12)))
					t220 := v3
					v9 = t219
					store32(m.memory[int64(uint32(t220))+168:], uint32(v9))
					m.fn419(i32(0), i32(4))
				}
			l77:
				m.fn51(v3+i32(232), i32(1072492), i32(1))
				v9 = v9 * i32(12)
			l85:
				{
					if v9 == 0 {
						t223 := int32(load32(m.memory[int64(uint32(v3))+232:]))
						v8 = t223
						t224 := int32(load32(m.memory[int64(uint32(v3))+236:]))
						t225 := v3 + i32(272)
						v9 = t224
						t226 := int32(load32(m.memory[int64(uint32(v3))+240:]))
						m.fn75(t225, v9, t226)
						m.fn16(v8, v9)
						m.fn74(v3+i32(272), i32(10))
						m.fn51(v3+i32(232), i32(1072492), i32(1))
					l87:
						if v4 == 0 {
							t227 := int32(load32(m.memory[int64(uint32(v3))+232:]))
							v4 = t227
							t228 := int32(load32(m.memory[int64(uint32(v3))+236:]))
							t229 := v3 + i32(272)
							v8 = t228
							t230 := int32(load32(m.memory[int64(uint32(v3))+240:]))
							m.fn75(t229, v8, t230)
							m.fn16(v4, v8)
							v10 = v2 + v5*i32(12)
						l91:
							{
								if v2 == v10 {
									t239 := int32(load32(m.memory[int64(uint32(v3))+280:]))
									store32(m.memory[int64(uint32(v0))+8:], uint32(t239))
									t240 := int64(load64(m.memory[int64(uint32(v3))+272:]))
									store64(m.memory[uint32(v0):], uint64(t240))
									m.fn78(v3 + i32(160))
									m.fn860(v3 + i32(216))
									goto l35
								}
								m.fn74(v3+i32(272), i32(10))
								t231 := int32(load32(m.memory[int64(uint32(v2))+4:]))
								v4 = t231
								t232 := int32(load32(m.memory[int64(uint32(v2))+8:]))
								v8 = t232
								m.fn51(v3+i32(232), i32(1072492), i32(1))
								v8 = v8 << 4
							l90:
								{
									if v8 == 0 {
										t235 := int32(load32(m.memory[int64(uint32(v3))+232:]))
										v4 = t235
										t236 := int32(load32(m.memory[int64(uint32(v3))+236:]))
										t237 := v3 + i32(272)
										v8 = t236
										t238 := int32(load32(m.memory[int64(uint32(v3))+240:]))
										m.fn75(t237, v8, t238)
										m.fn16(v4, v8)
										v2 = v2 + i32(12)
										goto l91
									}
									t233 := int32(load32(m.memory[int64(uint32(v4))+8:]))
									v9 = t233
									t234 := int32(load32(m.memory[int64(uint32(v4))+4:]))
									v11 = t234
									m.fn74(v3+i32(232), i32(32))
									m.fn75(v3+i32(232), v11, v9)
									m.fn75(v3+i32(232), i32(1072493), i32(2))
									v8 = v8 + i32(-16)
									v4 = v4 + i32(16)
									goto l90
								}
							}
						}
						m.fn74(v3+i32(232), i32(32))
						m.fn75(v3+i32(232), i32(1080559), i32(3))
						m.fn75(v3+i32(232), i32(1072493), i32(2))
						v4 = v4 + i32(-1)
						goto l87
					}
					t221 := int32(load32(m.memory[int64(uint32(v8))+8:]))
					v11 = t221
					t222 := int32(load32(m.memory[int64(uint32(v8))+4:]))
					v10 = t222
					m.fn74(v3+i32(232), i32(32))
					m.fn75(v3+i32(232), v10, v11)
					m.fn75(v3+i32(232), i32(1072493), i32(2))
					v9 = v9 + i32(-12)
					v8 = v8 + i32(12)
					goto l85
				}
			}
			m.fn750(v2+v8, v4)
			v8 = v8 + i32(12)
			goto l74
		}
	l70:
		store32(m.memory[uint32(v0):], uint32(i32(-1)))
		m.fn860(v3 + i32(216))
	}
l35:
	m.g0 = v3 + i32(288)
}
func (m *Module) fn850(v0, v1 int32) {
	var v2, v3, v4 int32
	t0 := m.g0
	v2 = t0 - i32(48)
	m.g0 = v2
	m.fn71(v2+i32(40), v1)
	{
		{
			t1 := int32(load32(m.memory[int64(uint32(v2))+40:]))
			v1 = t1
			if v1 != 0 {
				goto l0
			}
			v1 = i32(0)
			goto l1
		}
	l0:
		t2 := int32(load32(m.memory[int64(uint32(v2))+44:]))
		m.fn824(v2+i32(32), v1, t2)
		{
			t3 := int32(load32(m.memory[int64(uint32(v2))+32:]))
			v1 = t3
			t4 := int32(load32(m.memory[int64(uint32(v2))+36:]))
			t5 := v1
			v3 = t4
			t6 := m.fn851(t5, v3)
			if t6 != 0 {
				goto l2
			}
			m.fn628(v2+i32(24), v1, v3)
			t7 := int32(load32(m.memory[int64(uint32(v2))+28:]))
			v3 = t7
			t8 := int32(load32(m.memory[int64(uint32(v2))+24:]))
			v1 = t8
		}
	l2:
		m.fn856(v2+i32(16), v1, v3, i32(92))
		t9 := int32(load32(m.memory[int64(uint32(v2))+16:]))
		t10 := int32(load32(m.memory[int64(uint32(v2))+20:]))
		m.fn46(v2+i32(8), t9, t10)
		t11 := int32(load32(m.memory[int64(uint32(v2))+12:]))
		t12 := v1
		v4 = t11
		p13 := i32(1)
		if v4 != 0 {
			p13 = t12
		}
		v1 = p13
		p14 := i32(0)
		if v4 != 0 {
			p14 = v3
		}
		v3 = p14
	}
l1:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
	store32(m.memory[uint32(v0):], uint32(v1))
	m.g0 = v2 + i32(48)
}
func (m *Module) fn851(v0, v1 int32) int32 {
	var v2, v3 int32
	t0 := m.g0
	v2 = t0 - i32(32)
	m.g0 = v2
	v3 = i32(0)
	m.memory[int64(uint32(v2))+28] = byte(i32(0))
	store32(m.memory[int64(uint32(v2))+20:], uint32(v0))
	store32(m.memory[int64(uint32(v2))+24:], uint32(v0+v1))
l1:
	{
		m.fn577(v2+i32(8), v2+i32(20))
		t1 := int32(load32(m.memory[int64(uint32(v2))+8:]))
		if t1 != i32(1) {
			goto l0
		}
		t2 := int32(load32(m.memory[int64(uint32(v2))+12:]))
		if t2 != i32(92) {
			goto l0
		}
		v3 = v3 + i32(1)
		goto l1
	}
l0:
	m.g0 = v2 + i32(32)
	return v3 & i32(1)
}
