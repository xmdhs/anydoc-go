package core

import (
	"math/bits"
)

func (m *Module) fn222(v0, v1 int32) {
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
func (m *Module) fn223(v0 int32) {
	var v1 int32
	t0 := m.g0
	v1 = t0 - i32(16)
	m.g0 = v1
	t1 := int32(load32(m.memory[uint32(v0):]))
	m.fn273(v1+i32(8), v0, t1, i32(1), i32(4), i32(16))
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
func (m *Module) fn224(v0 int32) {
	m.fn185(v0)
	m.fn78(v0 + i32(28))
}
func (m *Module) fn225(v0, v1 int32) {
	{
		t0 := int32(load32(m.memory[uint32(v1):]))
		if t0 == i32(-1) {
			goto l0
		}
		t1 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t2 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		m.fn31(v0, t1, t2)
		return
	}
l0:
	store32(m.memory[uint32(v0):], uint32(i32(-1)))
}
func (m *Module) fn226(v0 int32) {
	var v1, v2, v3, v4 int32
	var v5 int64
	var v6, v7 int32
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
		{
			t3 := int32(load32(m.memory[int64(uint32(v0))+12:]))
			v4 = t3
			if v4 == 0 {
				goto l1
			}
			v0 = v3 + i32(8)
			t4 := int64(load64(m.memory[uint32(v3):]))
			v5 = (t4 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
			v6 = v3
		l4:
			if v4 == 0 {
				goto l1
			}
		l3:
			{
				if v5 != i64(0) {
					v7 = v6 - int32(int64(bits.TrailingZeros64(uint64(v5))))<<1&i32(240)
					t6 := int32(load32(m.memory[uint32(v7+i32(-16)):]))
					t7 := int32(load32(m.memory[uint32(v7+i32(-12)):]))
					m.fn134(t6, t7)
					v4 = v4 + i32(-1)
					v5 = (v5 + i64(-1)) & v5
					goto l4
				}
				v6 = v6 + i32(-128)
				t5 := int64(load64(m.memory[uint32(v0):]))
				v5 = (t5 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
				v0 = v0 + i32(8)
				goto l3
			}
		}
	l1:
		m.fn39(v1+i32(4), i32(16), i32(8), v2+i32(1))
		t8 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		t9 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t10 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		m.fn40(v3-t8, t9, t10)
	}
l0:
	m.g0 = v1 + i32(16)
}
func (m *Module) fn227(v0 int32) {
	m.fn228(v0)
	m.fn229(v0 + i32(232))
}
func (m *Module) fn228(v0 int32) {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	m.fn128(t0, t1)
	m.fn124(v0 + i32(24))
}
func (m *Module) fn229(v0 int32) {
	t0 := int32(load32(m.memory[int64(uint32(v0))+32:]))
	t1 := int32(load32(m.memory[int64(uint32(v0))+36:]))
	m.fn16(t0, t1)
	t2 := int32(load32(m.memory[int64(uint32(v0))+44:]))
	t3 := int32(load32(m.memory[int64(uint32(v0))+48:]))
	m.fn188(t2, t3)
}
func (m *Module) fn230(v0 int32) {
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
	m.fn136(t4, v2, i32(4), i32(16))
}
func (m *Module) fn231(v0 int32) {
	var v1, v2, v3, v4, v5, v6, v7 int32
	t0 := m.g0
	v1 = t0 - i32(16)
	m.g0 = v1
	v2 = i32(0)
	t1 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	t2 := int32(load32(m.memory[uint32(v0):]))
	v3 = t2
	p3 := i32(0)
	if v3 != 0 {
		p3 = t1
	}
	v4 = p3
	var p4 int32
	if v3 != i32(0) {
		p4 = 1
	}
	v5 = p4
	t5 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	v6 = t5
	v0 = v3
l15:
	if v4 != 0 {
		if v5&i32(1) == 0 {
			m.fn153(i32(1080172))
			panic("unreachable")
		}
		if v2 != 0 {
			goto l6
		}
	l8:
		{
			if v6 == 0 {
				v6 = i32(0)
				v0 = i32(0)
				v2 = v3
				goto l6
			}
			v6 = v6 + i32(-1)
			t7 := int32(load32(m.memory[int64(uint32(v3))+756:]))
			v3 = t7
			goto l8
		}
	l6:
		v4 = v4 + i32(-1)
	l11:
		{
			t8 := int32(load16(m.memory[int64(uint32(v2))+754:]))
			if uint32(v6) < uint32(t8) {
				if v0 != 0 {
					v3 = v2 + v6<<2 + i32(760)
				l14:
					{
						t12 := int32(load32(m.memory[uint32(v3):]))
						v7 = t12
						v3 = v7 + i32(756)
						v0 = v0 + i32(-1)
						if v0 != 0 {
							goto l14
						}
					}
					v0 = i32(0)
					goto l13
				}
				v0 = v6 + i32(1)
				v7 = v2
				goto l13
			}
			m.fn232(v1+i32(4), v2, v0)
			t9 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			v2 = t9
			if v2 == 0 {
				m.fn153(i32(1071172))
				panic("unreachable")
			}
			t10 := int32(load32(m.memory[int64(uint32(v1))+12:]))
			v6 = t10
			t11 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			v0 = t11
			goto l11
		}
	l13:
		v3 = v2 + v6*i32(12)
		t13 := int32(load32(m.memory[int64(uint32(v3))+620:]))
		t14 := int32(load32(m.memory[uint32(v3+i32(624)):]))
		m.fn16(t13, t14)
		m.fn224(v2 + v6*i32(56))
		v5 = i32(1)
		v3 = i32(0)
		v6 = v0
		v0 = i32(0)
		v2 = v7
		goto l15
	}
	if v5&i32(1) == 0 {
		goto l1
	}
	if v2 != 0 {
		goto l16
	}
l4:
	if v6 != 0 {
		v6 = v6 + i32(-1)
		t6 := int32(load32(m.memory[int64(uint32(v3))+756:]))
		v3 = t6
		goto l4
	}
	v2 = v3
	v3 = i32(0)
	goto l16
l16:
	{
		m.fn232(v1+i32(4), v2, v3)
		t15 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		v2 = t15
		if v2 == 0 {
			goto l1
		}
		t16 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		v3 = t16
		goto l16
	}
l1:
	m.g0 = v1 + i32(16)
}
func (m *Module) fn232(v0, v1, v2 int32) {
	var v3 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v1))+616:]))
		v3 = t0
		if v3 == 0 {
			goto l0
		}
		t1 := int32(load16(m.memory[int64(uint32(v1))+752:]))
		store32(m.memory[int64(uint32(v0))+8:], uint32(t1))
		store32(m.memory[int64(uint32(v0))+4:], uint32(v2+i32(1)))
	}
l0:
	store32(m.memory[uint32(v0):], uint32(v3))
	t3 := v1
	p2 := i32(756)
	if v2 != 0 {
		p2 = i32(804)
	}
	m.fn40(t3, i32(4), p2)
}
func (m *Module) fn233(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(t0):]))
	t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t3 := int32(load32(m.memory[uint32(t2):]))
	t4 := m.fn234(t1, t3-v1<<4+i32(-16))
	return t4
}
func (m *Module) fn234(v0, v1 int32) int32 {
	var v2, v3, v4 int32
	t0 := int32(load32(m.memory[uint32(v1):]))
	v2 = t0
	t1 := int32(load32(m.memory[uint32(v0):]))
	t2 := v2
	v3 = t1
	var p3 int32
	if t2&v3 == i32(-1) {
		p3 = 1
	}
	v4 = p3
	{
		if v3 == i32(-1) {
			goto l0
		}
		if v2 == i32(-1) {
			goto l0
		}
		t4 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		t5 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		t6 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t7 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		t8 := m.fn191(t4, t5, t6, t7)
		v4 = t8
	}
l0:
	return v4
}
func (m *Module) fn235(v0, v1, v2 int32) int32 {
	t0 := m.fn1851(v0, v1, v2)
	var p1 int32
	if t0 == 0 {
		p1 = 1
	}
	return p1
}
func (m *Module) fn236(v0, v1 int32) int32 {
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
		m.fn237(t5, i32(16), p6)
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
		store64(m.memory[int64(uint32(v2))+24:], uint64(i64(0x800000010)))
		store32(m.memory[int64(uint32(v2))+20:], uint32(v0+i32(16)))
		t11 := int32(load32(m.memory[uint32(v0):]))
		v4 = t11
		t12 := int64(load64(m.memory[uint32(v4):]))
		v8 = (t12 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
		v9 = v2 + i32(32)
		v1 = i32(0)
	l6:
		if v3 == 0 {
			t26 := int32(load32(m.memory[int64(uint32(v0))+12:]))
			t27 := v2
			v1 = t26
			store32(m.memory[int64(uint32(t27))+44:], uint32(v1))
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
				t19 := m.fn238(t17, t18, v10)
				v11 = t19
				t20 := m.fn26(t15, t16, v11)
				v12 = t20
				t21 := t14 + v12
				v13 = int32(uint32(int32(v11)) >> 25)
				m.memory[uint32(t21)] = byte(v13)
				m.memory[uint32(v6+v5&(v12+i32(-8))+i32(8))] = byte(v13)
				v12 = v6 + (v12^i32(-1))<<4
				t22 := int32(load32(m.memory[uint32(v0):]))
				t23 := v12
				v10 = t22 + (v10^i32(-1))<<4
				t24 := int64(load64(m.memory[int64(uint32(v10))+8:]))
				store64(m.memory[int64(uint32(t23))+8:], uint64(t24))
				t25 := int64(load64(m.memory[uint32(v10):]))
				store64(m.memory[uint32(v12):], uint64(t25))
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
	m.fn241(v0, v2+i32(16), i32(57), i32(16))
l7:
	v5 = i32(-1)
l2:
	m.g0 = v2 + i32(64)
	return v5
}
func (m *Module) fn237(v0, v1, v2 int32) {
	var v3, v4, v5, v6 int32
	t0 := m.g0
	v3 = t0 - i32(16)
	m.g0 = v3
	{
		{
			if v2 != 0 {
				goto l0
			}
			t1 := int64(load64(m.memory[int64(uint32(i32(0)))+1286264:]))
			store64(m.memory[int64(uint32(v0))+8:], uint64(t1))
			t2 := int64(load64(m.memory[int64(uint32(i32(0)))+1286256:]))
			store64(m.memory[uint32(v0):], uint64(t2))
			goto l1
		}
	l0:
		{
			{
				if uint32(v2) < uint32(i32(15)) {
					goto l2
				}
				if uint32(v2) > uint32(i32(0x1fffffff)) {
					m.fn242()
					panic("unreachable")
				}
				t3 := int32(uint32(v2<<3) / uint32(i32(7)))
				v2 = i32_shr_u(i32(-1), int32(bits.LeadingZeros32(uint32(t3+i32(-1))))) + i32(1)
				goto l4
			}
		l2:
			p4 := v2&i32(8) + i32(8)
			if uint32(v2) < uint32(i32(4)) {
				p4 = i32(4)
			}
			v2 = p4
		}
	l4:
		m.fn243(v3, v1, v2)
		t5 := int32(load32(m.memory[int64(uint32(v3))+8:]))
		v1 = t5
		t6 := int32(load32(m.memory[int64(uint32(v3))+4:]))
		v2 = t6
		{
			t7 := int32(load32(m.memory[uint32(v3):]))
			v4 = t7
			if v4 != 0 {
				goto l5
			}
			store32(m.memory[int64(uint32(v0))+8:], uint32(v1))
			store32(m.memory[int64(uint32(v0))+4:], uint32(v2))
			store32(m.memory[uint32(v0):], uint32(i32(0)))
			goto l1
		}
	l5:
		t8 := int32(load32(m.memory[int64(uint32(v3))+12:]))
		v5 = t8
		v6 = v2 + i32(9)
		if v6 == 0 {
			goto l6
		}
		memory_fill(m.memory, uint32(v4), i32(255), uint32(v6))
	l6:
		store32(m.memory[int64(uint32(v0))+12:], uint32(v5))
		store32(m.memory[int64(uint32(v0))+8:], uint32(v1))
		store32(m.memory[int64(uint32(v0))+4:], uint32(v2))
		store32(m.memory[uint32(v0):], uint32(v4))
	}
l1:
	m.g0 = v3 + i32(16)
}
func (m *Module) fn238(v0, v1, v2 int32) int64 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(t0):]))
	v0 = t1
	t2 := int64(load64(m.memory[uint32(v0):]))
	t3 := int64(load64(m.memory[uint32(v0+i32(8)):]))
	t4 := int32(load32(m.memory[uint32(v1):]))
	t5 := m.fn171(t2, t3, t4-v2<<4+i32(-16))
	return t5
}
func (m *Module) fn239(v0, v1 int32) {
	m.fn244(v0, v1, i32(4))
}
func (m *Module) fn240(v0 int32) {
	var v1, v2 int32
	t0 := m.g0
	v1 = t0 - i32(16)
	m.g0 = v1
	{
		t1 := int32(load32(m.memory[int64(uint32(v0))+16:]))
		v2 = t1
		if v2 == 0 {
			goto l0
		}
		t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		t3 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		m.fn39(v1+i32(4), t2, t3, v2+i32(1))
		t4 := int32(load32(m.memory[int64(uint32(v0))+12:]))
		t5 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		t6 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t7 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		m.fn40(t4-t5, t6, t7)
	}
l0:
	m.g0 = v1 + i32(16)
}
func (m *Module) fn241(v0, v1, v2, v3 int32) {
	var v4, v5, v6, v7, v8, v9, v10, v11, v12, v13 int32
	var v14 int64
	var v15, v16, v17, v18, v19 int32
	t0 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	v4 = t0 + i32(1)
	t1 := int32(uint32(v4) >> 3)
	var p2 int32
	if v4&i32(7) != i32(0) {
		p2 = 1
	}
	v5 = t1 + p2
	t3 := int32(load32(m.memory[uint32(v0):]))
	v6 = t3
	v7 = v6
l10:
	if v5 != 0 {
		t32 := int64(load64(m.memory[uint32(v7):]))
		t33 := v7
		v14 = t32
		store64(m.memory[uint32(t33):], uint64(int64(uint64(v14^i64(-1))>>7)&i64(72340172838076673)+(v14|i64(0x7f7f7f7f7f7f7f7f))))
		v7 = v7 + i32(8)
		v5 = v5 + i32(-1)
		goto l10
	}
	{
		if uint32(v4) < uint32(i32(8)) {
			goto l1
		}
		t4 := int64(load64(m.memory[uint32(v6):]))
		store64(m.memory[uint32(v6+v4):], uint64(t4))
		goto l2
	}
l1:
	if v4 == 0 {
		goto l2
	}
	memory_copy(m.memory, uint32(v6+i32(8)), uint32(v6), uint32(v4))
l2:
	v8 = v3 & i32(1)
	v9 = v3 & i32(1020)
	v10 = v3 & i32(3)
	v11 = int32(uint32(v3) >> 2)
	v6 = i32(0)
l4:
	{
		v7 = v6
		if v7 == v4 {
			t26 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t27 := v0
			v7 = t26
			p28 := int32(uint32(v7+i32(1))>>3) * i32(7)
			if uint32(v7) < uint32(i32(8)) {
				p28 = v7
			}
			t29 := int32(load32(m.memory[int64(uint32(v0))+12:]))
			store32(m.memory[int64(uint32(t27))+8:], uint32(p28-t29))
			return
		}
		v6 = v7 + i32(1)
		t5 := int32(load32(m.memory[uint32(v0):]))
		v5 = t5
		t6 := int32(m.memory[uint32(v5+v7)])
		if t6 != i32(128) {
			goto l4
		}
		v12 = v5 + v3*(v7^i32(-1))
		v13 = v12 + v9
	l8:
		{
			t7 := m.t0[uint(v2)].(func(int32, int32, int32) int64)(v1, v0, v7)
			v14 = t7
			t8 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			v5 = t8
			t9 := v5
			t10 := v7
			t11 := v5
			v15 = int32(v14)
			v16 = t11 & v15
			t12 := int32(load32(m.memory[uint32(v0):]))
			t13 := t10 - v16
			v17 = t12
			t14 := m.fn26(v17, v5, v14)
			v18 = t14
			if uint32(t9&(t13^(v18-v16))) < uint32(i32(8)) {
				t30 := v17 + v7
				v17 = int32(uint32(v15) >> 25)
				m.memory[uint32(t30)] = byte(v17)
				t31 := int32(load32(m.memory[uint32(v0):]))
				m.memory[uint32(t31+v5&(v7+i32(-8))+i32(8))] = byte(v17)
				goto l4
			}
			v16 = v17 + v18
			t15 := int32(m.memory[uint32(v16)])
			v19 = t15
			t16 := v16
			v15 = int32(uint32(v15) >> 25)
			m.memory[uint32(t16)] = byte(v15)
			t17 := int32(load32(m.memory[uint32(v0):]))
			m.memory[uint32(t17+(v18+i32(-8))&v5+i32(8))] = byte(v15)
			v5 = v17 + v3*(v18^i32(-1))
			{
				if v19 != i32(255) {
					if v11 == 0 {
						goto l7
					}
					m.fn244(v12, v5, v11)
				l7:
					if v10 == 0 {
						goto l8
					}
					v17 = v5 + v9
					v5 = i32(0)
					{
						if v10 == i32(1) {
							goto l9
						}
						t21 := int32(load16(m.memory[uint32(v13):]))
						v5 = t21
						t22 := int32(load16(m.memory[uint32(v17):]))
						store16(m.memory[uint32(v13):], uint16(t22))
						store16(m.memory[uint32(v17):], uint16(v5))
						v5 = i32(2)
						if v8 == 0 {
							goto l8
						}
					}
				l9:
					v15 = v13 + v5
					t23 := int32(m.memory[uint32(v15)])
					v16 = t23
					t24 := v15
					v5 = v17 + v5
					t25 := int32(m.memory[uint32(v5)])
					m.memory[uint32(t24)] = byte(t25)
					m.memory[uint32(v5)] = byte(v16)
					goto l8
				}
				t18 := int32(load32(m.memory[int64(uint32(v0))+4:]))
				v17 = t18
				t19 := int32(load32(m.memory[uint32(v0):]))
				m.memory[uint32(t19+v7)] = byte(i32(255))
				t20 := int32(load32(m.memory[uint32(v0):]))
				m.memory[uint32(t20+v17&(v7+i32(-8))+i32(8))] = byte(i32(255))
				if v3 == 0 {
					goto l4
				}
				memory_copy(m.memory, uint32(v5), uint32(v12), uint32(v3))
				goto l4
			}
		}
	}
}
func (m *Module) fn242() {
	m.fn91(i32(1280660), i32(57), i32(1280688))
	panic("unreachable")
}
func (m *Module) fn243(v0, v1, v2 int32) {
	var v3, v4, v5, v6 int32
	t0 := m.g0
	v3 = t0 - i32(16)
	m.g0 = v3
	m.fn1616(v3+i32(4), v1, i32(8), v2)
	{
		t1 := int32(load32(m.memory[int64(uint32(v3))+4:]))
		v4 = t1
		if v4 == 0 {
			m.fn242()
			panic("unreachable")
		}
		t2 := int32(load32(m.memory[int64(uint32(v3))+12:]))
		v5 = t2
		{
			{
				t3 := int32(load32(m.memory[int64(uint32(v3))+8:]))
				v6 = t3
				if v6 != 0 {
					goto l1
				}
				v1 = v4
				goto l2
			}
		l1:
			t4 := m.fn248(v6, v4)
			v1 = t4
		}
	l2:
		if v1 == 0 {
			m.fn85(v4, v6)
			panic("unreachable")
		}
		store32(m.memory[int64(uint32(v0))+12:], uint32(i32(0)))
		t5 := v0
		v4 = v2 + i32(-1)
		store32(m.memory[int64(uint32(t5))+4:], uint32(v4))
		store32(m.memory[uint32(v0):], uint32(v1+v5))
		t7 := v0
		p6 := int32(uint32(v2)>>3) * i32(7)
		if uint32(v2) < uint32(i32(9)) {
			p6 = v4
		}
		store32(m.memory[int64(uint32(t7))+8:], uint32(p6))
		m.g0 = v3 + i32(16)
		return
	}
}
func (m *Module) fn244(v0, v1, v2 int32) {
	var v3 int32
l1:
	if v2 != 0 {
		t0 := int32(load32(m.memory[uint32(v0):]))
		v3 = t0
		t1 := int32(load32(m.memory[uint32(v1):]))
		store32(m.memory[uint32(v0):], uint32(t1))
		store32(m.memory[uint32(v1):], uint32(v3))
		v2 = v2 + i32(-1)
		v1 = v1 + i32(4)
		v0 = v0 + i32(4)
		goto l1
	}
}
func (m *Module) fn245(v0, v1 int32) {
l1:
	{
		if v1 == 0 {
			return
		}
		t0 := int32(load32(m.memory[uint32(v0):]))
		t1 := int32(load32(m.memory[uint32(v0+i32(4)):]))
		m.fn1301(t0, t1, i32(1), i32(1))
		v1 = v1 + i32(-1)
		v0 = v0 + i32(12)
		goto l1
	}
}
func (m *Module) fn246() int32 {
	var v0, v1 int32
	t0 := m.g0
	v0 = t0 - i32(16)
	m.g0 = v0
	m.fn247(v0+i32(8), i32(4), i32(756))
	{
		t1 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		v1 = t1
		if v1 != 0 {
			m.g0 = v0 + i32(16)
			return v1
		}
		m.fn85(i32(4), i32(756))
		panic("unreachable")
	}
}
func (m *Module) fn247(v0, v1, v2 int32) {
	t0 := m.fn248(v2, v1)
	v1 = t0
	store32(m.memory[int64(uint32(v0))+4:], uint32(v2))
	store32(m.memory[uint32(v0):], uint32(v1))
}
func (m *Module) fn248(v0, v1 int32) int32 {
	if uint32(v1) < uint32(i32(9)) {
		t1 := m.fn4(v0)
		return t1
	}
	t0 := m.fn1557(v1, v0)
	return t0
}
func (m *Module) fn249(v0, v1, v2, v3, v4 int32) {
	var v5, v6, v7, v8 int32
	t0 := int32(load16(m.memory[int64(uint32(v0))+754:]))
	t1 := v0 + i32(620)
	v5 = t0
	v6 = v5 + i32(1)
	m.fn250(t1, v6, v1, v2)
	m.fn251(v0, v6, v1, v3)
	v2 = v1 + i32(1)
	v7 = v0 + i32(756)
	{
		v3 = v5 + i32(2)
		t2 := v3
		v8 = v1 + i32(2)
		if uint32(t2) <= uint32(v8) {
			goto l0
		}
		v5 = (v5 - v1) << 2
		if v5 == 0 {
			goto l0
		}
		memory_copy(m.memory, uint32(v7+v8<<2), uint32(v7+v2<<2), uint32(v5))
	}
l0:
	store32(m.memory[uint32(v7+v2<<2):], uint32(v4))
	store16(m.memory[int64(uint32(v0))+754:], uint16(v6))
	p3 := v2
	if uint32(v3) > uint32(v2) {
		p3 = v3
	}
	v3 = p3
	v1 = v1<<2 + v0 + i32(760)
l2:
	{
		if v3 == v2 {
			return
		}
		t4 := int32(load32(m.memory[uint32(v1):]))
		v6 = t4
		store16(m.memory[int64(uint32(v6))+752:], uint16(v2))
		store32(m.memory[int64(uint32(v6))+616:], uint32(v0))
		v1 = v1 + i32(4)
		v2 = v2 + i32(1)
		goto l2
	}
}
func (m *Module) fn250(v0, v1, v2, v3 int32) {
	var v4 int32
	{
		t0 := v1
		v4 = v2 + i32(1)
		if uint32(t0) <= uint32(v4) {
			goto l0
		}
		v1 = (v1 + (v2 ^ i32(-1))) * i32(12)
		if v1 == 0 {
			goto l0
		}
		memory_copy(m.memory, uint32(v0+v4*i32(12)), uint32(v0+v2*i32(12)), uint32(v1))
	}
l0:
	v2 = v0 + v2*i32(12)
	t1 := int32(load32(m.memory[int64(uint32(v3))+8:]))
	store32(m.memory[int64(uint32(v2))+8:], uint32(t1))
	t2 := int64(load64(m.memory[uint32(v3):]))
	store64(m.memory[uint32(v2):], uint64(t2))
}
func (m *Module) fn251(v0, v1, v2, v3 int32) {
	var v4 int32
	{
		t0 := v1
		v4 = v2 + i32(1)
		if uint32(t0) <= uint32(v4) {
			goto l0
		}
		v1 = (v1 + (v2 ^ i32(-1))) * i32(56)
		if v1 == 0 {
			goto l0
		}
		memory_copy(m.memory, uint32(v0+v4*i32(56)), uint32(v0+v2*i32(56)), uint32(v1))
	}
l0:
	memory_copy(m.memory, uint32(v0+v2*i32(56)), uint32(v3), uint32(i32(56)))
}
func (m *Module) fn252() int32 {
	var v0, v1 int32
	t0 := m.g0
	v0 = t0 - i32(16)
	m.g0 = v0
	m.fn247(v0+i32(8), i32(4), i32(804))
	{
		t1 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		v1 = t1
		if v1 != 0 {
			m.g0 = v0 + i32(16)
			return v1
		}
		m.fn85(i32(4), i32(804))
		panic("unreachable")
	}
}
func (m *Module) fn253(v0, v1, v2 int32) {
	var v3, v4, v5, v6 int32
	t0 := m.g0
	v3 = t0 - i32(32)
	m.g0 = v3
	t1 := int32(load16(m.memory[int64(uint32(v1))+754:]))
	v4 = t1
	m.memory[int64(uint32(v3))+28] = byte(i32(0))
	store32(m.memory[int64(uint32(v3))+24:], uint32(v4))
	store32(m.memory[int64(uint32(v3))+20:], uint32(i32(0)))
	v5 = v1 + i32(756)
l1:
	{
		m.fn254(v3+i32(8), v3+i32(20))
		t2 := int32(load32(m.memory[int64(uint32(v3))+8:]))
		if t2 != i32(1) {
			goto l0
		}
		t3 := int32(load32(m.memory[int64(uint32(v3))+12:]))
		t4 := v5
		v4 = t3
		t5 := int32(load32(m.memory[uint32(t4+v4<<2):]))
		v6 = t5
		store16(m.memory[int64(uint32(v6))+752:], uint16(v4))
		store32(m.memory[int64(uint32(v6))+616:], uint32(v1))
		goto l1
	}
l0:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v2))
	store32(m.memory[uint32(v0):], uint32(v1))
	m.g0 = v3 + i32(32)
}
func (m *Module) fn254(v0, v1 int32) {
	var v2, v3, v4 int32
	v2 = i32(0)
	{
		{
			t0 := int32(m.memory[int64(uint32(v1))+8])
			if t0 == 0 {
				goto l0
			}
			goto l1
		}
	l0:
		t1 := int32(load32(m.memory[uint32(v1):]))
		v3 = t1
		t2 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t3 := v3
		v4 = t2
		if uint32(t3) > uint32(v4) {
			goto l1
		}
		if uint32(v3) < uint32(v4) {
			goto l2
		}
		v2 = i32(1)
		m.memory[int64(uint32(v1))+8] = byte(i32(1))
		goto l1
	l2:
		v2 = i32(1)
		store32(m.memory[uint32(v1):], uint32(v3+i32(1)))
	}
l1:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
	store32(m.memory[uint32(v0):], uint32(v2))
}
func (m *Module) fn255(v0, v1, v2, v3 int32) {
	if v1 != v3 {
		m.fn256(i32(1072679), i32(40), i32(1072720))
		panic("unreachable")
	}
	v1 = v1 * i32(12)
	if v1 == 0 {
		return
	}
	memory_copy(m.memory, uint32(v2), uint32(v0), uint32(v1))
}
func (m *Module) fn256(v0, v1, v2 int32) {
	m.fn91(v0, v1<<1|i32(1), v2)
	panic("unreachable")
}
func (m *Module) fn257(v0, v1, v2, v3 int32) {
	if v1 != v3 {
		m.fn256(i32(1072679), i32(40), i32(1072720))
		panic("unreachable")
	}
	v1 = v1 * i32(56)
	if v1 == 0 {
		return
	}
	memory_copy(m.memory, uint32(v2), uint32(v0), uint32(v1))
}
func (m *Module) fn258(v0, v1, v2, v3 int32) {
	var v4, v5, v6 int32
	t0 := int32(load32(m.memory[uint32(v1):]))
	v4 = t0
	t1 := int32(load16(m.memory[int64(uint32(v4))+754:]))
	t2 := v4 + i32(620)
	v5 = t1 + i32(1)
	t3 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	t4 := v5
	v6 = t3
	m.fn250(t2, t4, v6, v2)
	m.fn251(v4, v5, v6, v3)
	store16(m.memory[int64(uint32(v4))+754:], uint16(v5))
	store32(m.memory[int64(uint32(v0))+8:], uint32(v6))
	store32(m.memory[uint32(v0):], uint32(v4))
	t5 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	store32(m.memory[int64(uint32(v0))+4:], uint32(t5))
}
func (m *Module) fn259(v0, v1, v2, v3 int32) int32 {
	t1 := v0
	t2 := v2
	p0 := v3
	if uint32(v1) < uint32(v3) {
		p0 = v1
	}
	t3 := m.fn1851(t1, t2, p0)
	v2 = t3
	p4 := v1 - v3
	if v2 != 0 {
		p4 = v2
	}
	v3 = p4
	var p5 int32
	if v3 > i32(0) {
		p5 = 1
	}
	var p6 int32
	if v3 < i32(0) {
		p6 = 1
	}
	return p5 - p6
}
func (m *Module) fn260(v0 int32) {
	var v1 int32
	t0 := m.g0
	v1 = t0 - i32(16)
	m.g0 = v1
	t1 := int32(load32(m.memory[uint32(v0):]))
	m.fn1690(v1+i32(8), v0, t1, i32(1), i32(4), i32(4))
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
func (m *Module) fn261(v0, v1 int32) {
	var v2, v3, v4 int32
	v2 = i32(0)
	{
		t0 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		v3 = t0
		t1 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t2 := v3
		v4 = t1
		if uint32(t2) > uint32(v4) {
			goto l0
		}
		store32(m.memory[int64(uint32(v1))+4:], uint32(v4+i32(-1)))
		t3 := int32(load32(m.memory[uint32(v1):]))
		t4 := v1
		v2 = t3
		store32(m.memory[uint32(t4):], uint32(v2+i32(4)))
	}
l0:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
	store32(m.memory[uint32(v0):], uint32(v2))
}
func (m *Module) fn262(v0, v1 int32) {
	var v2 int32
	{
		t0 := int32(load32(m.memory[uint32(v0):]))
		t1 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		t2 := v1
		v2 = t1
		if uint32(t2) <= uint32(t0-v2) {
			return
		}
		m.fn62(v0, v2, v1, i32(8), i32(24))
	}
}
func (m *Module) fn263(v0, v1 int32) int32 {
	var v2, v3, v4 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	t1 := int32(load32(m.memory[uint32(v0):]))
	v3 = t1
	v0 = i32(1)
	{
		t2 := int32(load32(m.memory[uint32(v1):]))
		v4 = t2
		t3 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t4 := v4
		v1 = t3
		t5 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		t6 := m.t0[uint(t5)].(func(int32, int32, int32) int32)(t4, i32(1283949), i32(18))
		if t6 != 0 {
			goto l0
		}
		t7 := m.fn282(v4, v1, v3)
		if t7 != 0 {
			goto l0
		}
		store32(m.memory[int64(uint32(v2))+12:], uint32(i32(5)))
		store32(m.memory[int64(uint32(v2))+8:], uint32(v3+i32(16)))
		t8 := m.fn284(v4, v1, i32(1052694), v2+i32(8))
		v0 = t8
	}
l0:
	m.g0 = v2 + i32(16)
	return v0
}
func (m *Module) fn264(v0, v1, v2, v3, v4 int32) int32 {
	var v5, v6, v7, v8, v9 int32
	t0 := m.g0
	v5 = t0 - i32(32)
	m.g0 = v5
	v6 = i32(1)
	{
		t1 := int32(load32(m.memory[uint32(v0):]))
		v7 = t1
		t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		t3 := v7
		t4 := v1
		t5 := v2
		v8 = t2
		t6 := int32(load32(m.memory[int64(uint32(v8))+12:]))
		v9 = t6
		t7 := m.t0[uint(v9)].(func(int32, int32, int32) int32)(t3, t4, t5)
		if t7 != 0 {
			goto l0
		}
		{
			{
				t8 := int32(m.memory[int64(uint32(v0))+10])
				if t8&i32(128) != 0 {
					goto l1
				}
				v6 = i32(1)
				t9 := m.t0[uint(v9)].(func(int32, int32, int32) int32)(v7, i32(1108163), i32(1))
				if t9 != 0 {
					goto l0
				}
				t10 := m.t0[uint(v4)].(func(int32, int32) int32)(v3, v0)
				if t10 != 0 {
					goto l0
				}
				t11 := int32(load32(m.memory[uint32(v0):]))
				v7 = t11
				t12 := int32(load32(m.memory[int64(uint32(v0))+4:]))
				t13 := int32(load32(m.memory[int64(uint32(t12))+12:]))
				v9 = t13
				goto l2
			}
		l1:
			t14 := m.t0[uint(v9)].(func(int32, int32, int32) int32)(v7, i32(1108164), i32(2))
			if t14 != 0 {
				goto l0
			}
			v6 = i32(1)
			m.memory[int64(uint32(v5))+15] = byte(i32(1))
			store32(m.memory[int64(uint32(v5))+4:], uint32(v8))
			store32(m.memory[uint32(v5):], uint32(v7))
			store32(m.memory[int64(uint32(v5))+20:], uint32(i32(1109040)))
			t15 := int64(load64(m.memory[int64(uint32(v0))+8:]))
			store64(m.memory[int64(uint32(v5))+24:], uint64(t15))
			store32(m.memory[int64(uint32(v5))+8:], uint32(v5+i32(15)))
			store32(m.memory[int64(uint32(v5))+16:], uint32(v5))
			t16 := m.t0[uint(v4)].(func(int32, int32) int32)(v3, v5+i32(16))
			if t16 != 0 {
				goto l0
			}
			t17 := int32(load32(m.memory[int64(uint32(v5))+16:]))
			t18 := int32(load32(m.memory[int64(uint32(v5))+20:]))
			t19 := int32(load32(m.memory[int64(uint32(t18))+12:]))
			t20 := m.t0[uint(t19)].(func(int32, int32, int32) int32)(t17, i32(1108161), i32(2))
			if t20 != 0 {
				goto l0
			}
		}
	l2:
		t21 := m.t0[uint(v9)].(func(int32, int32, int32) int32)(v7, i32(1282664), i32(1))
		v6 = t21
	}
l0:
	m.g0 = v5 + i32(32)
	return v6
}
func (m *Module) fn265(v0, v1 int32) int32 {
	var v2, v3, v4 int32
	t0 := int32(load32(m.memory[uint32(v0):]))
	v2 = t0
	v0 = i32(1)
	{
		t1 := int32(load32(m.memory[uint32(v1):]))
		v3 = t1
		t2 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t3 := v3
		v1 = t2
		t4 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		v4 = t4
		t5 := m.t0[uint(v4)].(func(int32, int32, int32) int32)(t3, i32(1283967), i32(17))
		if t5 != 0 {
			goto l0
		}
		t6 := m.fn282(v3, v1, v2)
		if t6 != 0 {
			goto l0
		}
		t7 := m.t0[uint(v4)].(func(int32, int32, int32) int32)(v3, i32(1283984), i32(2))
		v0 = t7
	}
l0:
	return v0
}
func (m *Module) fn266(v0, v1 int32) int32 {
	var v2, v3, v4 int32
	t0 := int32(load32(m.memory[uint32(v0):]))
	v2 = t0
	v0 = i32(1)
	{
		t1 := int32(load32(m.memory[uint32(v1):]))
		v3 = t1
		t2 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t3 := v3
		v1 = t2
		t4 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		v4 = t4
		t5 := m.t0[uint(v4)].(func(int32, int32, int32) int32)(t3, i32(1283986), i32(21))
		if t5 != 0 {
			goto l0
		}
		t6 := m.fn282(v3, v1, v2)
		if t6 != 0 {
			goto l0
		}
		t7 := m.t0[uint(v4)].(func(int32, int32, int32) int32)(v3, i32(1283984), i32(2))
		v0 = t7
	}
l0:
	return v0
}
