package core

import (
	"math/bits"
)

func (m *Module) fn42(v0, v1 int32) {
	var v2 int32
	{
		t0 := int32(load32(m.memory[uint32(v0):]))
		v2 = t0
		switch v2 >> 31 & (v2 + i32(-0x7fffffff)) {
		case 1:
			v0 = v0 + i32(4)
			fallthrough
		case 0:
			t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t2 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			m.fn1553(t1, t2, v1)
			fallthrough
		default:
		}
	}
}
func (m *Module) fn43(v0, v1, v2 int32) {
	var v3, v4 int32
	{
		v3 = int32(uint32(v2-v1) >> 5)
		t0 := int32(load32(m.memory[uint32(v0):]))
		t1 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		t2 := v3
		v4 = t1
		if uint32(t2) <= uint32(t0-v4) {
			goto l0
		}
		m.fn62(v0, v4, v3, i32(4), i32(4))
		t3 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		v4 = t3
	}
l0:
	v3 = v4 + v3
	t4 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	v4 = t4 + v4<<2
l2:
	{
		if v2 == v1 {
			goto l1
		}
		t5 := v4
		v2 = v2 + i32(-32)
		store32(m.memory[uint32(t5):], uint32(v2))
		v4 = v4 + i32(4)
		goto l2
	}
l1:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v3))
}
func (m *Module) fn44(v0, v1 int32) {
	m.fn136(v0, v1, i32(4), i32(4))
}
func (m *Module) fn45(v0, v1, v2 int32) {
	var v3 int32
	t0 := m.g0
	v3 = t0 - i32(16)
	m.g0 = v3
	store32(m.memory[int64(uint32(v3))+12:], uint32(i32(0)))
	store64(m.memory[int64(uint32(v3))+4:], uint64(i64(0x100000000)))
	m.fn793(v1, v2, v3+i32(4))
	t1 := int32(load32(m.memory[int64(uint32(v3))+12:]))
	store32(m.memory[int64(uint32(v0))+8:], uint32(t1))
	t2 := int64(load64(m.memory[int64(uint32(v3))+4:]))
	store64(m.memory[uint32(v0):], uint64(t2))
	m.g0 = v3 + i32(16)
}
func (m *Module) fn46(v0, v1, v2 int32) {
	var v3, v4, v5, v6, v7, v8, v9, v10, v11 int32
	v3 = v1 + v2
	v4 = i32(0)
	v5 = v1
l6:
	{
		v6 = v4
		v2 = v5
		if v2 != v3 {
			goto l0
		}
		v4 = v6
		v5 = v3
		v7 = i32(0)
		v6 = i32(0)
		goto l13
	l0:
		{
			{
				t0 := int32(int8(m.memory[uint32(v2)]))
				v8 = t0
				if v8 <= i32(-1) {
					goto l2
				}
				v5 = v2 + i32(1)
				v8 = v8 & i32(255)
				goto l3
			}
		l2:
			t1 := int32(m.memory[int64(uint32(v2))+1])
			v5 = t1 & i32(63)
			v4 = v8 & i32(31)
			if uint32(v8) > uint32(i32(-33)) {
				goto l4
			}
			v8 = v4<<6 | v5
			v5 = v2 + i32(2)
			goto l3
		l4:
			t2 := int32(m.memory[int64(uint32(v2))+2])
			v5 = v5<<6 | t2&i32(63)
			if uint32(v8) >= uint32(i32(-16)) {
				goto l5
			}
			v8 = v5 | v4<<12
			v5 = v2 + i32(3)
			goto l3
		l5:
			t3 := int32(m.memory[int64(uint32(v2))+3])
			v8 = v5<<6 | t3&i32(63) | v4<<18&i32(0x1c0000)
			v5 = v2 + i32(4)
		}
	l3:
		v4 = v6 - v2 + v5
		v7 = v4
		t4 := m.fn630(v8)
		if t4 != 0 {
			goto l6
		}
	}
l13:
	{
		t5 := v5
		v8 = v3
		if t5 == v8 {
			goto l7
		}
		{
			v3 = v8 + i32(-1)
			t6 := int32(int8(m.memory[uint32(v3)]))
			v2 = t6
			if v2 > i32(-1) {
				goto l8
			}
			{
				v3 = v8 + i32(-2)
				t7 := int32(m.memory[uint32(v3)])
				v9 = t7
				v10 = int32(int8(v9))
				if v10 < i32(-64) {
					goto l9
				}
				v9 = v9 & i32(31)
				goto l10
			}
		l9:
			{
				{
					v3 = v8 + i32(-3)
					t8 := int32(m.memory[uint32(v3)])
					v9 = t8
					v11 = int32(int8(v9))
					if v11 < i32(-64) {
						goto l11
					}
					v9 = v9 & i32(15)
					goto l12
				}
			l11:
				v3 = v8 + i32(-4)
				t9 := int32(m.memory[uint32(v3)])
				v9 = t9&i32(7)<<6 | v11&i32(63)
			}
		l12:
			v9 = v9<<6 | v10&i32(63)
		l10:
			v2 = v9<<6 | v2&i32(63)
		}
	l8:
		t10 := m.fn630(v2)
		if t10 != 0 {
			goto l13
		}
	}
	v7 = v8 - v5 + v4
l7:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v7-v6))
	store32(m.memory[uint32(v0):], uint32(v1+v6))
}
func (m *Module) fn47(v0, v1 int32) {
	var v2 int32
	{
		t0 := int32(load32(m.memory[uint32(v0):]))
		t1 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		t2 := v1
		v2 = t1
		if uint32(t2) <= uint32(t0-v2) {
			return
		}
		m.fn62(v0, v2, v1, i32(1), i32(1))
	}
}
func (m *Module) fn48(v0 int32) int32 {
	var v1, v2 int32
	t0 := m.g0
	v1 = t0 - i32(16)
	m.g0 = v1
	m.fn374(v1+i32(8), v0)
	t1 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	v0 = t1
	t2 := int32(load32(m.memory[int64(uint32(v1))+12:]))
	v2 = t2
	m.g0 = v1 + i32(16)
	p3 := i32(-1)
	if v0&i32(1) != 0 {
		p3 = v2
	}
	return p3
}
func (m *Module) fn49(v0, v1 int32) {
	var v2 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	{
		{
			if uint32(v1) < uint32(i32(192)) {
				goto l0
			}
			m.fn1627(v2+i32(4), v1, i32(1112780))
			store64(m.memory[int64(uint32(v0))+4:], uint64(i64(0)))
			store32(m.memory[uint32(v0):], uint32(v1))
			t1 := int32(load32(m.memory[int64(uint32(v2))+4:]))
			if t1 == i32(-1) {
				goto l1
			}
			t2 := int32(load32(m.memory[int64(uint32(v2))+12:]))
			store32(m.memory[int64(uint32(v0))+8:], uint32(t2))
			t3 := int64(load64(m.memory[int64(uint32(v2))+4:]))
			store64(m.memory[uint32(v0):], uint64(t3))
			goto l1
		}
	l0:
		store64(m.memory[int64(uint32(v0))+4:], uint64(i64(0)))
		t5 := v0
		p4 := v1
		if uint32(v1+i32(-65)) < uint32(i32(26)) {
			p4 = v1 | i32(32)
		}
		store32(m.memory[uint32(t5):], uint32(p4))
	}
l1:
	m.g0 = v2 + i32(16)
}
func (m *Module) fn50(v0, v1 int32) {
	var v2, v3, v4, v5 int32
	t0 := int32(load32(m.memory[uint32(v0):]))
	v2 = t0
	v3 = v2<<2 + v0 + i32(8)
	t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	v4 = t1
l5:
	{
		if v4 == v2 {
			return
		}
		v5 = i32(45)
		t2 := int32(load32(m.memory[uint32(v3):]))
		v0 = t2
		if v0 == i32(32) {
			goto l1
		}
		if v0 == i32(45) {
			goto l1
		}
		t3 := m.fn817(v0)
		if t3 != 0 {
			goto l2
		}
		if v0&i32(0x1ffff0) == i32(65056) {
			goto l2
		}
		if v0&i32(2097088) == i32(7616) {
			goto l2
		}
		if v0&i32(0x1ffffe) == i32(2402) {
			goto l2
		}
		if v0&i32(0x1ffffc) == i32(2304) {
			goto l2
		}
		if uint32(v0+i32(-768)) < uint32(i32(112)) {
			goto l2
		}
		if uint32(v0+i32(-1155)) < uint32(i32(7)) {
			goto l2
		}
		if uint32(v0+i32(-2362)) < uint32(i32(22)) {
			goto l2
		}
		if uint32(v0+i32(-2385)) < uint32(i32(7)) {
			goto l2
		}
		if uint32(v0+i32(-6832)) < uint32(i32(80)) {
			goto l2
		}
		if uint32(v0+i32(-8400)) < uint32(i32(48)) {
			goto l2
		}
		v5 = v0 + i32(-8255)
		if uint32(v5) > uint32(i32(21)) {
			goto l3
		}
		if i32_shl(i32(1), v5)&i32(0x200003) != 0 {
			goto l2
		}
		goto l3
	}
l3:
	if uint32(v0+i32(-65075)) < uint32(i32(2)) {
		goto l2
	}
	if v0 == i32(95) {
		goto l2
	}
	if v0 == i32(65343) {
		goto l2
	}
	v5 = v0
	if uint32(v0+i32(-65101)) < uint32(i32(3)) {
		goto l1
	}
	goto l4
l2:
	v5 = v0
l1:
	m.fn74(v1, v5)
l4:
	v2 = v2 + i32(1)
	v3 = v3 + i32(4)
	goto l5
}
func (m *Module) fn51(v0, v1, v2 int32) {
	var v3, v4, v5 int32
	t0 := m.g0
	v3 = t0 - i32(16)
	m.g0 = v3
	m.fn59(v3+i32(8), v2, i32(1), i32(1))
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
func (m *Module) fn52(v0, v1, v2 int32) {
	var v3, v4, v5, v6, v7, v8 int32
	t0 := m.g0
	v3 = t0 - i32(48)
	m.g0 = v3
	t1 := int32(load32(m.memory[int64(uint32(v2))+4:]))
	t2 := v3 + i32(8)
	v4 = t1
	t3 := int32(load32(m.memory[int64(uint32(v2))+8:]))
	t4 := v4
	v5 = t3
	m.fn31(t2, t4, v5)
	v6 = v1 + i32(32)
	{
		{
			t5 := m.fn32(v1, v3+i32(8))
			if t5 != 0 {
				goto l0
			}
			v4 = i32(1)
			{
				t6 := m.fn912(v6, v2)
				v5 = t6
				if v5 == 0 {
					goto l1
				}
				t7 := int32(load32(m.memory[uint32(v5):]))
				v4 = t7
			}
		l1:
			store32(m.memory[int64(uint32(v3))+4:], uint32(v4))
			v4 = v4 + i32(1)
		l5:
			{
				store32(m.memory[int64(uint32(v3))+20:], uint32(i32(5)))
				store32(m.memory[int64(uint32(v3))+12:], uint32(i32(25)))
				store32(m.memory[int64(uint32(v3))+8:], uint32(v2))
				store32(m.memory[int64(uint32(v3))+16:], uint32(v3+i32(4)))
				m.fn73(v3+i32(36), i32(0x1000cf), v3+i32(8))
				t8 := int32(load32(m.memory[int64(uint32(v3))+36:]))
				v7 = t8
				t9 := int32(load32(m.memory[int64(uint32(v3))+40:]))
				v5 = t9
				{
					if v4 != 0 {
						goto l2
					}
					store32(m.memory[uint32(v0):], uint32(i32(-1)))
					m.fn16(v7, v5)
					t10 := int32(load32(m.memory[uint32(v2):]))
					t11 := int32(load32(m.memory[int64(uint32(v2))+4:]))
					m.fn16(t10, t11)
					goto l3
				}
			l2:
				t12 := int32(load32(m.memory[int64(uint32(v3))+44:]))
				v8 = t12
				store32(m.memory[int64(uint32(v3))+4:], uint32(v4))
				m.fn31(v3+i32(8), v5, v8)
				{
					t13 := m.fn32(v1, v3+i32(8))
					if t13 != 0 {
						goto l4
					}
					m.fn16(v7, v5)
					v4 = v4 + i32(1)
					goto l5
				}
			l4:
			}
			t14 := int32(load32(m.memory[int64(uint32(v2))+8:]))
			store32(m.memory[int64(uint32(v3))+16:], uint32(t14))
			t15 := int64(load64(m.memory[uint32(v2):]))
			store64(m.memory[int64(uint32(v3))+8:], uint64(t15))
			m.fn36(v6, v3+i32(8), v4)
			m.fn31(v3+i32(36), v5, v8)
			m.fn1554(v3+i32(8), v6, v3+i32(36))
			m.fn1555(v3 + i32(8))
			store32(m.memory[int64(uint32(v0))+8:], uint32(v8))
			store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
			store32(m.memory[uint32(v0):], uint32(v7))
			goto l3
		}
	l0:
		m.fn31(v3+i32(36), v4, v5)
		m.fn1554(v3+i32(8), v6, v3+i32(36))
		m.fn1555(v3 + i32(8))
		t16 := int32(load32(m.memory[int64(uint32(v2))+8:]))
		store32(m.memory[int64(uint32(v0))+8:], uint32(t16))
		t17 := int64(load64(m.memory[uint32(v2):]))
		store64(m.memory[uint32(v0):], uint64(t17))
	}
l3:
	m.g0 = v3 + i32(48)
}
func (m *Module) fn53(v0, v1, v2, v3 int32) {
	var v4, v5, v6 int32
	var v7 int64
	var v8, v9 int32
	var v10 int64
	var v11, v12 int32
	var v13, v14 int64
	var v15 int32
	t0 := m.g0
	v4 = t0 - i32(32)
	m.g0 = v4
	m.fn51(v4+i32(16), v2, v3)
	t1 := int64(load64(m.memory[int64(uint32(v0))+16:]))
	t2 := int64(load64(m.memory[int64(uint32(v0))+24:]))
	t3 := int32(load32(m.memory[int64(uint32(v4))+20:]))
	v5 = t3
	t4 := int32(load32(m.memory[int64(uint32(v4))+24:]))
	t5 := v5
	v6 = t4
	t6 := m.fn540(t1, t2, t5, v6)
	v7 = t6
	t7 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	v8 = t7
	t8 := v8
	v9 = int32(v7)
	v3 = t8 & v9
	v10 = int64(uint64(v7)>>25) & i64(127) * i64(72340172838076673)
	v11 = v0 + i32(16)
	t9 := int32(load32(m.memory[uint32(v0):]))
	v2 = t9
	v12 = i32(0)
l7:
	{
		t10 := int64(load64(m.memory[uint32(v2+v3):]))
		v13 = t10
		v14 = v13 ^ v10
		v14 = (v14 ^ i64(-1)) & (v14 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
		{
			{
				{
				l2:
					if v14 == 0 {
						goto l0
					}
					{
						v15 = v2 + (i32(0)-(int32(uint32(int64(bits.TrailingZeros64(uint64(v14))))>>3)+v3)&v8)*i32(28)
						t11 := int32(load32(m.memory[uint32(v15+i32(-24)):]))
						t12 := int32(load32(m.memory[uint32(v15+i32(-20)):]))
						t13 := m.fn191(t11, t12, v5, v6)
						if t13 != 0 {
							t14 := int32(load32(m.memory[int64(uint32(v4))+16:]))
							m.fn16(t14, v5)
							v3 = i32(-1)
							v2 = v0
							goto l3
						}
						v14 = (v14 + i64(-1)) & v14
						goto l2
					}
				}
			l0:
				if v13&(v13<<1)&i64(-0x7f7f7f7f7f7f7f80) == 0 {
					t36 := v3
					v12 = v12 + i32(8)
					v3 = (t36 + v12) & v8
					goto l7
				}
				m.fn676(v0, v11)
				v2 = int32(int64(uint64(v7) >> 32))
				t15 := int64(load64(m.memory[int64(uint32(v4))+20:]))
				v7 = t15
				t16 := int32(load32(m.memory[int64(uint32(v4))+16:]))
				v3 = t16
				v15 = v9
			}
		l3:
			t17 := int32(load32(m.memory[uint32(v1+i32(4)):]))
			t18 := int32(load32(m.memory[uint32(v1+i32(8)):]))
			m.fn31(v4, t17, t18)
			m.memory[int64(uint32(v4))+12] = byte(i32(0))
			{
				{
					if v3 == i32(-1) {
						goto l5
					}
					t19 := int64(load64(m.memory[int64(uint32(v4))+8:]))
					store64(m.memory[int64(uint32(v4))+24:], uint64(t19))
					t20 := int64(load64(m.memory[uint32(v4):]))
					store64(m.memory[int64(uint32(v4))+16:], uint64(t20))
					t21 := int32(load32(m.memory[uint32(v0):]))
					v1 = t21
					t22 := int32(load32(m.memory[int64(uint32(v0))+4:]))
					t23 := v1
					t24 := v1
					v5 = t22
					t25 := v5
					t26 := int64(uint32(v2)) << 32
					v14 = int64(uint32(v15))
					t27 := m.fn26(t24, t25, t26|v14)
					v2 = t27
					v15 = t23 + v2
					t28 := int32(m.memory[uint32(v15)])
					v8 = t28
					t29 := v15
					v6 = int32(int64(uint64(v14) >> 25))
					m.memory[uint32(t29)] = byte(v6)
					m.memory[uint32(v1+v5&(v2+i32(-8))+i32(8))] = byte(v6)
					t30 := int32(load32(m.memory[int64(uint32(v0))+12:]))
					store32(m.memory[int64(uint32(v0))+12:], uint32(t30+i32(1)))
					t31 := int32(load32(m.memory[int64(uint32(v0))+8:]))
					store32(m.memory[int64(uint32(v0))+8:], uint32(t31-v8&i32(1)))
					v0 = v1 + (i32(0)-v2)*i32(28)
					store64(m.memory[uint32(v0+i32(-24)):], uint64(v7))
					store32(m.memory[uint32(v0+i32(-28)):], uint32(v3))
					v0 = v0 + i32(-16)
					t32 := int64(load64(m.memory[int64(uint32(v4))+16:]))
					store64(m.memory[uint32(v0):], uint64(t32))
					t33 := int64(load64(m.memory[int64(uint32(v4))+24:]))
					store64(m.memory[int64(uint32(v0))+8:], uint64(t33))
					goto l6
				}
			l5:
				t34 := int32(load32(m.memory[uint32(v4):]))
				t35 := int32(load32(m.memory[int64(uint32(v4))+4:]))
				m.fn16(t34, t35)
			}
		l6:
			m.g0 = v4 + i32(32)
			return
		}
	}
}
func (m *Module) fn54(v0, v1, v2 int32) {
	var v3, v4, v5 int32
	v1 = v1 * i32(28)
	t0 := int32(load32(m.memory[int64(uint32(v2))+4:]))
	v3 = t0
	t1 := int32(load32(m.memory[uint32(v2):]))
	v4 = t1
l4:
	{
		{
			if v1 == 0 {
				return
			}
			t2 := int32(load32(m.memory[uint32(v0):]))
			v5 = t2
			p3 := i32(1)
			if uint32(v5) > uint32(i32(2)) {
				p3 = v5 + i32(-3)
			}
			switch p3 + i32(-1) {
			case 0:
				t4 := int32(load32(m.memory[uint32(v0+i32(20)):]))
				t5 := int32(load32(m.memory[uint32(v0+i32(24)):]))
				m.fn54(t4, t5, v2)
				goto l2
			case 2:
				goto l3
			default:
				goto l2
			}
		}
	l3:
		t6 := int32(load32(m.memory[uint32(v0+i32(8)):]))
		t7 := int32(load32(m.memory[uint32(v0+i32(12)):]))
		m.fn53(v4, v3, t6, t7)
	}
l2:
	v0 = v0 + i32(28)
	v1 = v1 + i32(-28)
	goto l4
}
func (m *Module) fn55(v0, v1 int32) {
	var v2 int32
	{
		t0 := int32(load32(m.memory[uint32(v0):]))
		v2 = t0
		switch v2 >> 31 & (v2 + i32(-0x7fffffff)) {
		case 1:
			v0 = v0 + i32(4)
			fallthrough
		case 0:
			t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t2 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			m.fn1556(t1, t2, v1)
			fallthrough
		default:
		}
	}
}
func (m *Module) fn56(v0, v1 int32) {
	var v2 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	{
		if v1 == 0 {
			goto l0
		}
		m.fn39(v2+i32(4), i32(8), i32(8), v1+i32(1))
		t1 := int32(load32(m.memory[int64(uint32(v2))+12:]))
		t2 := int32(load32(m.memory[int64(uint32(v2))+4:]))
		t3 := int32(load32(m.memory[int64(uint32(v2))+8:]))
		m.fn40(v0-t1, t2, t3)
	}
l0:
	m.g0 = v2 + i32(16)
}
func (m *Module) fn57(v0 int32) {
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
					m.fn16(t6, t7)
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
func (m *Module) fn58(v0, v1 int32) {
	var v2, v3, v4, v5 int32
	t0 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	v2 = t0
	t1 := int32(load32(m.memory[uint32(v1):]))
	v3 = t1
	t2 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	v4 = t2
l1:
	{
		if v3 == v4 {
			goto l0
		}
		t3 := v1
		v5 = v3 + i32(32)
		store32(m.memory[uint32(t3):], uint32(v5))
		m.fn849(v0, v3, v2)
		v3 = v5
		t4 := int32(load32(m.memory[uint32(v0):]))
		if t4 == i32(-1) {
			goto l1
		}
		return
	}
l0:
	store32(m.memory[uint32(v0):], uint32(i32(-1)))
}
func (m *Module) fn59(v0, v1, v2, v3 int32) {
	var v4 int32
	t0 := m.g0
	v4 = t0 - i32(16)
	m.g0 = v4
	m.fn1(v4+i32(4), v1, i32(0), v2, v3)
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
func (m *Module) fn60(v0, v1 int32) {
	var v2 int32
	{
		t0 := int32(load32(m.memory[uint32(v0):]))
		t1 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		t2 := v1
		v2 = t1
		if uint32(t2) <= uint32(t0-v2) {
			return
		}
		m.fn1842(v0, v2, v1, i32(4), i32(12))
	}
}
func (m *Module) fn61(v0, v1 int32) {
	var v2, v3, v4, v5, v6 int32
	t0 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	v2 = t0
	t1 := int32(load32(m.memory[uint32(v1):]))
	v3 = t1
	t2 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	v4 = t2
	{
	l2:
		{
			v5 = v3
			if v5 != v4 {
				goto l0
			}
			v5 = i32(0)
			goto l1
		l0:
			t3 := v1
			v3 = v5 + i32(28)
			store32(m.memory[uint32(t3):], uint32(v3))
			t4 := m.fn912(v2, v5)
			v6 = t4
			if v6 == 0 {
				goto l2
			}
		}
		t5 := int32(load32(m.memory[uint32(v6):]))
		v3 = t5
	}
l1:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
	store32(m.memory[uint32(v0):], uint32(v5))
}
func (m *Module) fn62(v0, v1, v2, v3, v4 int32) {
	var v5 int32
	t0 := m.g0
	v5 = t0 - i32(16)
	m.g0 = v5
	m.fn273(v5+i32(8), v0, v1, v2, v3, v4)
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
func (m *Module) fn63(v0, v1 int32) {
	var v2, v3, v4 int32
	t0 := m.g0
	v2 = t0 - i32(4112)
	m.g0 = v2
	{
		{
			p1 := i32(1000000)
			if uint32(v1) < uint32(i32(1000000)) {
				p1 = v1
			}
			v3 = p1
			t2 := v3
			v4 = v1 - int32(uint32(v1)>>1)
			p3 := v4
			if uint32(v3) > uint32(v4) {
				p3 = t2
			}
			v3 = p3
			if uint32(v3) < uint32(i32(513)) {
				goto l0
			}
			m.fn59(v2+i32(8), v3, i32(4), i32(8))
			t4 := int32(load32(m.memory[int64(uint32(v2))+12:]))
			t5 := v0
			t6 := v1
			v3 = t4
			t7 := int32(load32(m.memory[int64(uint32(v2))+8:]))
			t8 := v3
			v4 = t7
			t9 := v4
			var p10 int32
			if uint32(v1) < uint32(i32(65)) {
				p10 = 1
			}
			m.fn989(t5, t6, t8, t9, p10)
			m.fn76(v4, v3)
			goto l1
		}
	l0:
		t11 := v0
		t12 := v1
		t13 := v2 + i32(16)
		var p14 int32
		if uint32(v1) < uint32(i32(65)) {
			p14 = 1
		}
		m.fn989(t11, t12, t13, i32(512), p14)
	}
l1:
	m.g0 = v2 + i32(4112)
}
func (m *Module) fn64(v0, v1 int32) {
	var v2, v3, v4 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		v2 = t0
		t1 := int32(load32(m.memory[uint32(v1+i32(-4)):]))
		if uint32(v2) >= uint32(t1) {
			return
		}
		v3 = v1 + i32(-12)
		t2 := int32(load32(m.memory[uint32(v1):]))
		v4 = t2
	l2:
		{
			v1 = v3
			t3 := v1 + i32(12)
			v3 = v1 + i32(4)
			t4 := int64(load64(m.memory[uint32(v3):]))
			store64(m.memory[uint32(t3):], uint64(t4))
			if v3 == v0 {
				goto l1
			}
			v3 = v1 + i32(-8)
			t5 := int32(load32(m.memory[uint32(v1):]))
			if uint32(v2) < uint32(t5) {
				goto l2
			}
		}
		v3 = v1 + i32(12)
		goto l3
	l1:
		v3 = v1 + i32(12)
	l3:
		store32(m.memory[uint32(v1+i32(4)):], uint32(v4))
		store32(m.memory[uint32(v3+i32(-4)):], uint32(v2))
	}
}
func (m *Module) fn65(v0, v1, v2, v3 int32) {
	var v4, v5 int32
	t0 := m.g0
	v4 = t0 - i32(80)
	m.g0 = v4
	store32(m.memory[int64(uint32(v4))+28:], uint32(v3))
	store32(m.memory[int64(uint32(v4))+20:], uint32(v1))
	store32(m.memory[int64(uint32(v4))+24:], uint32(v1+v2<<5))
	m.fn58(v4+i32(44), v4+i32(20))
	{
		t1 := int32(load32(m.memory[int64(uint32(v4))+44:]))
		if t1 == i32(-1) {
			goto l0
		}
		m.fn59(v4+i32(8), i32(4), i32(4), i32(12))
		t2 := int32(load32(m.memory[int64(uint32(v4))+8:]))
		v1 = t2
		t3 := int32(load32(m.memory[int64(uint32(v4))+12:]))
		v5 = t3
		t4 := int32(load32(m.memory[int64(uint32(v4))+52:]))
		store32(m.memory[int64(uint32(v5))+8:], uint32(t4))
		t5 := int64(load64(m.memory[int64(uint32(v4))+44:]))
		store64(m.memory[uint32(v5):], uint64(t5))
		store32(m.memory[int64(uint32(v4))+40:], uint32(i32(1)))
		store32(m.memory[int64(uint32(v4))+36:], uint32(v5))
		store32(m.memory[int64(uint32(v4))+32:], uint32(v1))
		t6 := int32(load32(m.memory[int64(uint32(v4))+28:]))
		store32(m.memory[int64(uint32(v4))+64:], uint32(t6))
		t7 := int64(load64(m.memory[int64(uint32(v4))+20:]))
		store64(m.memory[int64(uint32(v4))+56:], uint64(t7))
		v1 = i32(12)
		v3 = i32(1)
	l3:
		{
			m.fn58(v4+i32(68), v4+i32(56))
			t8 := int32(load32(m.memory[int64(uint32(v4))+68:]))
			if t8 == i32(-1) {
				t14 := int32(load32(m.memory[int64(uint32(v4))+40:]))
				v1 = t14
				t15 := int32(load32(m.memory[int64(uint32(v4))+36:]))
				v3 = t15
				goto l4
			}
			{
				t9 := int32(load32(m.memory[int64(uint32(v4))+32:]))
				if v3 != t9 {
					goto l2
				}
				m.fn60(v4+i32(32), i32(1))
				t10 := int32(load32(m.memory[int64(uint32(v4))+36:]))
				v5 = t10
			}
		l2:
			v2 = v5 + v1
			t11 := int32(load32(m.memory[int64(uint32(v4))+76:]))
			store32(m.memory[int64(uint32(v2))+8:], uint32(t11))
			t12 := int64(load64(m.memory[int64(uint32(v4))+68:]))
			store64(m.memory[uint32(v2):], uint64(t12))
			t13 := v4
			v3 = v3 + i32(1)
			store32(m.memory[int64(uint32(t13))+40:], uint32(v3))
			v1 = v1 + i32(12)
			goto l3
		}
	}
l0:
	v1 = i32(0)
	store32(m.memory[int64(uint32(v4))+40:], uint32(i32(0)))
	store64(m.memory[int64(uint32(v4))+32:], uint64(i64(0x400000000)))
	v3 = i32(4)
l4:
	m.fn77(v0, v3, v1, i32(1080488), i32(2))
	m.fn78(v4 + i32(32))
	m.g0 = v4 + i32(80)
}
func (m *Module) fn66(v0, v1 int64, v2 int32) int64 {
	var v3 int32
	t0 := m.g0
	v3 = t0 - i32(64)
	m.g0 = v3
	store64(m.memory[int64(uint32(v3))+48:], uint64(i64(0)))
	store64(m.memory[int64(uint32(v3))+56:], uint64(i64(0)))
	store64(m.memory[int64(uint32(v3))+40:], uint64(v1))
	store64(m.memory[int64(uint32(v3))+24:], uint64(v1^i64(8387220255154660723)))
	store64(m.memory[int64(uint32(v3))+16:], uint64(v1^i64(7237128888997146477)))
	store64(m.memory[int64(uint32(v3))+32:], uint64(v0))
	store64(m.memory[int64(uint32(v3))+8:], uint64(v0^i64(0x6c7967656e657261)))
	store64(m.memory[uint32(v3):], uint64(v0^i64(8317987319222330741)))
	m.fn172(v3, v2)
	t1 := m.fn174(v3)
	v1 = t1
	m.g0 = v3 + i32(64)
	return v1
}
func (m *Module) fn67(v0, v1 int32) int32 {
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
				t19 := m.fn714(t17, t18, v10)
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
	m.fn241(v0, v2+i32(16), i32(39), i32(4))
l7:
	v5 = i32(-1)
l2:
	m.g0 = v2 + i32(64)
	return v5
}
func (m *Module) fn68(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(t0):]))
	t2 := int32(load32(m.memory[uint32(t1):]))
	t3 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t4 := int32(load32(m.memory[uint32(t3):]))
	t5 := int32(load32(m.memory[uint32(t4-v1<<2+i32(-4)):]))
	var p6 int32
	if t2 == t5 {
		p6 = 1
	}
	return p6
}
func (m *Module) fn69(v0, v1, v2 int32, v3 int64, v4, v5 int32) {
	var v6 int32
	var v7 int64
	var v8, v9 int32
	var v10 int64
	var v11, v12, v13 int32
	v6 = v2 & int32(v3)
	v7 = int64(uint64(v3)>>25) & i64(127) * i64(72340172838076673)
	v8 = i32(0)
	v9 = i32(0)
	{
	l7:
		{
			t0 := int64(load64(m.memory[uint32(v1+v6):]))
			v10 = t0
			v3 = v10 ^ v7
			v3 = (v3 ^ i64(-1)) & (v3 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
		l2:
			{
				if v3 == 0 {
					goto l0
				}
				v11 = i32(0)
				t1 := v4
				v12 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v3))))>>3) + v6) & v2
				t2 := m.t0[uint(v5)].(func(int32, int32) int32)(t1, v12)
				if t2 != 0 {
					goto l1
				}
				v3 = (v3 + i64(-1)) & v3
				goto l2
			}
		l0:
			v3 = v10 & i64(-0x7f7f7f7f7f7f7f80)
			{
				if v8 == i32(1) {
					goto l3
				}
				if !(v3 == 0) {
					goto l4
				}
				v8 = i32(0)
				goto l5
			l4:
				v13 = (v6 + int32(uint32(int64(bits.TrailingZeros64(uint64(v3))))>>3)) & v2
			l3:
				if v3&(v10<<1) != i64(0) {
					goto l6
				}
				v8 = i32(1)
			l5:
				t3 := v6
				v9 = v9 + i32(8)
				v6 = (t3 + v9) & v2
				goto l7
			}
		l6:
		}
		v11 = i32(1)
		{
			t4 := int32(int8(m.memory[uint32(v1+v13)]))
			if t4 >= i32(0) {
				goto l8
			}
			v12 = v13
			goto l1
		}
	l8:
		t5 := int64(load64(m.memory[uint32(v1):]))
		v12 = int32(uint32(int64(bits.TrailingZeros64(uint64(t5&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
	}
l1:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v12))
	store32(m.memory[uint32(v0):], uint32(v11))
}
func (m *Module) fn70(v0, v1, v2 int32) {
	m.fn514(v0, i32(10), v1, v2)
	store16(m.memory[int64(uint32(v0))+36:], uint16(i32(0)))
	store32(m.memory[int64(uint32(v0))+32:], uint32(v2))
	store32(m.memory[int64(uint32(v0))+28:], uint32(i32(0)))
}
func (m *Module) fn71(v0, v1 int32) {
	var v2, v3, v4, v5 int32
	t0 := m.g0
	v2 = t0 - i32(64)
	m.g0 = v2
	v3 = i32(0)
	{
		{
			{
				t1 := int32(m.memory[int64(uint32(v1))+37])
				if t1 != 0 {
					goto l0
				}
				t2 := int32(load32(m.memory[int64(uint32(v1))+4:]))
				v4 = t2
				m.fn516(v2+i32(52), v1)
				{
					t3 := int32(load32(m.memory[int64(uint32(v2))+52:]))
					if t3 != i32(1) {
						goto l1
					}
					t4 := int32(load32(m.memory[int64(uint32(v1))+28:]))
					v3 = t4
					t5 := int32(load32(m.memory[int64(uint32(v2))+60:]))
					t6 := v1
					v5 = t5
					store32(m.memory[int64(uint32(t6))+28:], uint32(v5))
					v1 = v5 - v3
					v3 = v4 + v3
					goto l2
				}
			l1:
				m.fn517(v2+i32(40), v1)
				t7 := int32(load32(m.memory[int64(uint32(v2))+40:]))
				v4 = t7
				if v4 != 0 {
					goto l3
				}
			}
		l0:
			goto l4
		l3:
			t8 := int32(load32(m.memory[int64(uint32(v2))+44:]))
			v1 = t8
			v3 = v4
		}
	l2:
		store32(m.memory[int64(uint32(v2))+52:], uint32(i32(0)))
		m.fn522(v2+i32(32), i32(10), v2+i32(52))
		t9 := int32(load32(m.memory[int64(uint32(v2))+32:]))
		t10 := int32(load32(m.memory[int64(uint32(v2))+36:]))
		m.fn626(v2+i32(24), t9, t10, v3, v1)
		t11 := int32(load32(m.memory[int64(uint32(v2))+24:]))
		v4 = t11
		if v4 == 0 {
			goto l4
		}
		t12 := int32(load32(m.memory[int64(uint32(v2))+28:]))
		v1 = t12
		store32(m.memory[int64(uint32(v2))+52:], uint32(i32(0)))
		m.fn522(v2+i32(16), i32(13), v2+i32(52))
		t13 := int32(load32(m.memory[int64(uint32(v2))+16:]))
		t14 := int32(load32(m.memory[int64(uint32(v2))+20:]))
		m.fn626(v2+i32(8), t13, t14, v4, v1)
		t15 := int32(load32(m.memory[int64(uint32(v2))+12:]))
		t16 := int32(load32(m.memory[int64(uint32(v2))+8:]))
		t17 := v1
		v3 = t16
		p18 := t17
		if v3 != 0 {
			p18 = t15
		}
		v1 = p18
		p19 := v4
		if v3 != 0 {
			p19 = v3
		}
		v3 = p19
	}
l4:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
	store32(m.memory[uint32(v0):], uint32(v3))
	m.g0 = v2 + i32(64)
}
func (m *Module) fn72(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7, v8 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	v3 = i32(10)
	t1 := int32(load32(m.memory[uint32(v0):]))
	v4 = t1
	v5 = v4
	if uint32(v4) < uint32(i32(1000)) {
		goto l0
	}
	v3 = i32(10)
	v5 = v4
l1:
	{
		v6 = v2 + i32(6) + v3
		t2 := v6 + i32(-4)
		v0 = v5
		t3 := int32(uint32(v0) / uint32(i32(10000)))
		t4 := v0
		v5 = t3
		v7 = t4 - v5*i32(10000)
		t5 := int32(uint32(v7&i32(0xffff)) / uint32(i32(100)))
		v8 = t5
		t6 := int32(load16(m.memory[int64(uint32(v8<<1))+1109319:]))
		store16(m.memory[uint32(t2):], uint16(t6))
		t7 := int32(load16(m.memory[int64(uint32((v7-v8*i32(100))&i32(0xffff)<<1))+1109319:]))
		store16(m.memory[uint32(v6+i32(-2)):], uint16(t7))
		v3 = v3 + i32(-4)
		if uint32(v0) > uint32(i32(9999999)) {
			goto l1
		}
	}
l0:
	{
		if uint32(v5) > uint32(i32(9)) {
			goto l2
		}
		v0 = v5
		goto l3
	l2:
		t8 := v2 + i32(6)
		v3 = v3 + i32(-2)
		t9 := int32(uint32(v5&i32(0xffff)) / uint32(i32(100)))
		t10 := t8 + v3
		t11 := v5
		v0 = t9
		t12 := int32(load16(m.memory[int64(uint32((t11-v0*i32(100))&i32(0xffff)<<1))+1109319:]))
		store16(m.memory[uint32(t10):], uint16(t12))
	}
l3:
	{
		if v4 == 0 {
			goto l4
		}
		if v0 == 0 {
			goto l5
		}
	l4:
		t13 := v2 + i32(6)
		v3 = v3 + i32(-1)
		t14 := int32(m.memory[int64(uint32(v0<<1))+1109320])
		m.memory[uint32(t13+v3)] = byte(t14)
	}
l5:
	t15 := m.fn1638(v1, i32(1), i32(1), i32(0), v2+i32(6)+v3, i32(10)-v3)
	v3 = t15
	m.g0 = v2 + i32(16)
	return v3
}
func (m *Module) fn73(v0, v1, v2 int32) {
	if v2&i32(1) == 0 {
		goto l0
	}
	m.fn51(v0, v1, int32(uint32(v2)>>1))
	return
l0:
	m.fn6(v0, v1, v2)
}
func (m *Module) fn74(v0, v1 int32) {
	var v2, v3 int32
	t0 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	v2 = t0
	{
		if uint32(v1) >= uint32(i32(128)) {
			goto l0
		}
		v3 = i32(1)
		goto l1
	l0:
		if uint32(v1) >= uint32(i32(2048)) {
			goto l2
		}
		v3 = i32(2)
		goto l1
	l2:
		p1 := i32(4)
		if uint32(v1) < uint32(i32(65536)) {
			p1 = i32(3)
		}
		v3 = p1
	}
l1:
	m.fn47(v0, v3)
	t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	m.fn279(v1, t2+v2)
	store32(m.memory[int64(uint32(v0))+8:], uint32(v3+v2))
}
func (m *Module) fn75(v0, v1, v2 int32) {
	var v3 int32
	m.fn47(v0, v2)
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
func (m *Module) fn76(v0, v1 int32) {
	m.fn136(v0, v1, i32(4), i32(8))
}
func (m *Module) fn77(v0, v1, v2, v3, v4 int32) {
	var v5, v6, v7, v8, v9, v10, v11 int32
	t0 := m.g0
	v5 = t0 - i32(32)
	m.g0 = v5
	{
		if v2 == 0 {
			store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
			store64(m.memory[uint32(v0):], uint64(i64(0x100000000)))
			goto l17
		}
		{
			t1 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			v6 = t1
			t2 := v6
			t3 := v4
			v7 = v2 * i32(12)
			v2 = v7 + i32(-12)
			t4 := int32(uint32(v2) / uint32(i32(12)))
			v8 = t2 + t3*t4
			if uint32(v8) < uint32(v6) {
				goto l1
			}
			t5 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			v9 = t5
			v1 = v1 + i32(12)
			v10 = v1
		l3:
			{
				if v2 == 0 {
					m.fn1299(v5+i32(16), v8, i32(0))
					t7 := int32(load32(m.memory[int64(uint32(v5))+20:]))
					v2 = t7
					t8 := int32(load32(m.memory[int64(uint32(v5))+16:]))
					if t8 == i32(1) {
						t54 := int32(load32(m.memory[int64(uint32(v5))+24:]))
						m.fn2(v2, t54)
						panic("unreachable")
					}
					store32(m.memory[int64(uint32(v5))+12:], uint32(i32(0)))
					t9 := int32(load32(m.memory[int64(uint32(v5))+24:]))
					t10 := v5
					v11 = t9
					store32(m.memory[int64(uint32(t10))+8:], uint32(v11))
					store32(m.memory[int64(uint32(v5))+4:], uint32(v2))
					{
						if uint32(v6) <= uint32(v2) {
							goto l5
						}
						m.fn1842(v5+i32(4), i32(0), v6, i32(1), i32(1))
						t11 := int32(load32(m.memory[int64(uint32(v5))+8:]))
						v11 = t11
						t12 := int32(load32(m.memory[int64(uint32(v5))+12:]))
						v2 = t12
						goto l6
					}
				l5:
					v2 = i32(0)
					if v6 == 0 {
						goto l7
					}
				l6:
					if v6 == 0 {
						goto l7
					}
					memory_copy(m.memory, uint32(v11+v2), uint32(v9), uint32(v6))
				l7:
					t13 := v5
					v10 = v2 + v6
					store32(m.memory[int64(uint32(t13))+12:], uint32(v10))
					v2 = v8 - v10
					v10 = v11 + v10
					switch v4 + i32(-1) {
					default:
						v11 = v7 + i32(-12)
					l13:
						{
							if v11 == 0 {
								goto l12
							}
							t14 := int32(load32(m.memory[int64(uint32(v1))+4:]))
							v4 = t14
							t15 := int32(load32(m.memory[int64(uint32(v1))+8:]))
							v6 = t15
							m.fn635(v5+i32(16), v10, v2, i32(1))
							t16 := int32(load32(m.memory[int64(uint32(v5))+28:]))
							v2 = t16
							t17 := int32(load32(m.memory[int64(uint32(v5))+24:]))
							v10 = t17
							t18 := int32(load32(m.memory[int64(uint32(v5))+16:]))
							t19 := int32(load32(m.memory[int64(uint32(v5))+20:]))
							m.fn1843(t18, t19, v3, i32(1))
							m.fn635(v5+i32(16), v10, v2, v6)
							t20 := int32(load32(m.memory[int64(uint32(v5))+28:]))
							v2 = t20
							t21 := int32(load32(m.memory[int64(uint32(v5))+24:]))
							v10 = t21
							t22 := int32(load32(m.memory[int64(uint32(v5))+16:]))
							t23 := int32(load32(m.memory[int64(uint32(v5))+20:]))
							m.fn1843(t22, t23, v4, v6)
							v11 = v11 + i32(-12)
							v1 = v1 + i32(12)
							goto l13
						}
					case 1:
						v11 = v7 + i32(-12)
					l14:
						{
							if v11 == 0 {
								goto l12
							}
							t24 := int32(load32(m.memory[int64(uint32(v1))+4:]))
							v4 = t24
							t25 := int32(load32(m.memory[int64(uint32(v1))+8:]))
							v6 = t25
							m.fn635(v5+i32(16), v10, v2, i32(2))
							t26 := int32(load32(m.memory[int64(uint32(v5))+28:]))
							v2 = t26
							t27 := int32(load32(m.memory[int64(uint32(v5))+24:]))
							v10 = t27
							t28 := int32(load32(m.memory[int64(uint32(v5))+16:]))
							t29 := int32(load32(m.memory[int64(uint32(v5))+20:]))
							m.fn1843(t28, t29, v3, i32(2))
							m.fn635(v5+i32(16), v10, v2, v6)
							t30 := int32(load32(m.memory[int64(uint32(v5))+28:]))
							v2 = t30
							t31 := int32(load32(m.memory[int64(uint32(v5))+24:]))
							v10 = t31
							t32 := int32(load32(m.memory[int64(uint32(v5))+16:]))
							t33 := int32(load32(m.memory[int64(uint32(v5))+20:]))
							m.fn1843(t32, t33, v4, v6)
							v11 = v11 + i32(-12)
							v1 = v1 + i32(12)
							goto l14
						}
					case 2:
						v11 = v7 + i32(-12)
					l15:
						{
							if v11 == 0 {
								goto l12
							}
							t34 := int32(load32(m.memory[int64(uint32(v1))+4:]))
							v4 = t34
							t35 := int32(load32(m.memory[int64(uint32(v1))+8:]))
							v6 = t35
							m.fn635(v5+i32(16), v10, v2, i32(3))
							t36 := int32(load32(m.memory[int64(uint32(v5))+28:]))
							v2 = t36
							t37 := int32(load32(m.memory[int64(uint32(v5))+24:]))
							v10 = t37
							t38 := int32(load32(m.memory[int64(uint32(v5))+16:]))
							t39 := int32(load32(m.memory[int64(uint32(v5))+20:]))
							m.fn1843(t38, t39, v3, i32(3))
							m.fn635(v5+i32(16), v10, v2, v6)
							t40 := int32(load32(m.memory[int64(uint32(v5))+28:]))
							v2 = t40
							t41 := int32(load32(m.memory[int64(uint32(v5))+24:]))
							v10 = t41
							t42 := int32(load32(m.memory[int64(uint32(v5))+16:]))
							t43 := int32(load32(m.memory[int64(uint32(v5))+20:]))
							m.fn1843(t42, t43, v4, v6)
							v11 = v11 + i32(-12)
							v1 = v1 + i32(12)
							goto l15
						}
					case 3:
						v11 = v7 + i32(-12)
					l16:
						{
							if v11 == 0 {
								goto l12
							}
							t44 := int32(load32(m.memory[int64(uint32(v1))+4:]))
							v4 = t44
							t45 := int32(load32(m.memory[int64(uint32(v1))+8:]))
							v6 = t45
							m.fn635(v5+i32(16), v10, v2, i32(4))
							t46 := int32(load32(m.memory[int64(uint32(v5))+28:]))
							v2 = t46
							t47 := int32(load32(m.memory[int64(uint32(v5))+24:]))
							v10 = t47
							t48 := int32(load32(m.memory[int64(uint32(v5))+16:]))
							t49 := int32(load32(m.memory[int64(uint32(v5))+20:]))
							m.fn1843(t48, t49, v3, i32(4))
							m.fn635(v5+i32(16), v10, v2, v6)
							t50 := int32(load32(m.memory[int64(uint32(v5))+28:]))
							v2 = t50
							t51 := int32(load32(m.memory[int64(uint32(v5))+24:]))
							v10 = t51
							t52 := int32(load32(m.memory[int64(uint32(v5))+16:]))
							t53 := int32(load32(m.memory[int64(uint32(v5))+20:]))
							m.fn1843(t52, t53, v4, v6)
							v11 = v11 + i32(-12)
							v1 = v1 + i32(12)
							goto l16
						}
					}
				}
				v2 = v2 + i32(-12)
				t6 := int32(load32(m.memory[int64(uint32(v10))+8:]))
				v11 = t6
				v10 = v10 + i32(12)
				v8 = v11 + v8
				if uint32(v8) >= uint32(v11) {
					goto l3
				}
			}
		}
	l1:
		m.fn633(i32(1300956), i32(53), i32(1301012))
		panic("unreachable")
	l12:
		t55 := int64(load64(m.memory[int64(uint32(v5))+4:]))
		store64(m.memory[uint32(v0):], uint64(t55))
		store32(m.memory[int64(uint32(v0))+8:], uint32(v8-v2))
	}
l17:
	m.g0 = v5 + i32(32)
}
func (m *Module) fn78(v0 int32) {
	var v1 int32
	t0 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	v1 = t0
	t1 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	m.fn245(v1, t1)
	t2 := int32(load32(m.memory[uint32(v0):]))
	m.fn37(t2, v1)
}
func (m *Module) fn79(v0 int32) {
	var v1 int32
	t0 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	v1 = t0
	t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	v0 = t1
l1:
	if v1 == 0 {
		return
	}
	v1 = v1 + i32(-1)
	m.fn970(v0)
	v0 = v0 + i32(32)
	goto l1
}
func (m *Module) fn80(v0, v1 int32) {
	m.fn136(v0, v1, i32(8), i32(32))
}
func (m *Module) fn81(v0, v1 int32) {
l1:
	{
		if v1 == 0 {
			return
		}
		t0 := int32(load32(m.memory[uint32(v0):]))
		t1 := int32(load32(m.memory[uint32(v0+i32(4)):]))
		m.fn16(t0, t1)
		m.fn969(v0 + i32(12))
		v1 = v1 + i32(-1)
		v0 = v0 + i32(28)
		goto l1
	}
}
func (m *Module) fn82(v0, v1 int32) {
	m.fn136(v0, v1, i32(4), i32(28))
}
func (m *Module) fn83(v0, v1 int32) {
l1:
	{
		if v1 == 0 {
			return
		}
		t0 := int32(load32(m.memory[uint32(v0):]))
		t1 := int32(load32(m.memory[uint32(v0+i32(4)):]))
		m.fn16(t0, t1)
		t2 := int32(load32(m.memory[uint32(v0+i32(12)):]))
		t3 := int32(load32(m.memory[uint32(v0+i32(16)):]))
		m.fn16(t2, t3)
		t4 := int32(load32(m.memory[uint32(v0+i32(24)):]))
		t5 := int32(load32(m.memory[uint32(v0+i32(28)):]))
		m.fn16(t4, t5)
		v1 = v1 + i32(-1)
		v0 = v0 + i32(40)
		goto l1
	}
}
func (m *Module) fn84(v0, v1 int32) {
	m.fn136(v0, v1, i32(4), i32(40))
}
func (m *Module) fn85(v0, v1 int32) {
	m.fn90(v1, v0)
	panic("unreachable")
}
func (m *Module) fn86() {
	m.fn91(i32(1087526), i32(35), i32(1070260))
	panic("unreachable")
}
