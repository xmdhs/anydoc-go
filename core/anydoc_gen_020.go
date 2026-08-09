package core

import (
	"math"
	"math/bits"
)

func (m *Module) fn852(v0, v1, v2 int32) int32 {
	t0 := m.fn49(v0, i32(1099824), v1, v2)
	return t0
}
func (m *Module) fn853(v0, v1, v2, v3, v4 int32) int32 {
	{
		if v2 == i32(-1) {
			goto l0
		}
		t0 := int32(load32(m.memory[int64(uint32(v1))+16:]))
		t1 := m.t0[uint(t0)].(func(int32, int32) int32)(v0, v2)
		if t1 == 0 {
			goto l0
		}
		return i32(1)
	}
l0:
	if v3 != 0 {
		t2 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		t3 := m.t0[uint(t2)].(func(int32, int32, int32) int32)(v0, v3, v4)
		return t3
	}
	return i32(0)
}
func (m *Module) fn854(v0, v1 int32) {
	var v2 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	store32(m.memory[int64(uint32(v2))+12:], uint32(v1))
	store32(m.memory[int64(uint32(v2))+8:], uint32(v0))
	m.fn845(i32(0), v2+i32(8), i32(1098752), v2+i32(12), i32(1098752), i32(0), v2, i32(1099704))
	panic("unreachable")
}
func (m *Module) fn855(v0, v1 int32) int32 {
	var v2, v3 int32
	var v4 int64
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	t1 := int32(load32(m.memory[uint32(v0):]))
	v0 = t1
	{
		{
			t2 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			v3 = t2
			if v3&i32(0x2000000) != 0 {
				t4 := int64(load64(m.memory[uint32(v0):]))
				v4 = t4
				v0 = i32(17)
			l3:
				{
					t5 := int32(m.memory[int64(uint32(int32(v4)&i32(15)))+1098720])
					m.memory[uint32(v2+v0+i32(-2))] = byte(t5)
					v0 = v0 + i32(-1)
					v4 = int64(uint64(v4) >> 4)
					if v4 != i64(0) {
						goto l3
					}
				}
				t6 := m.fn683(v1, i32(1), i32(1122454), i32(2), v2+v0+i32(-1), i32(17)-v0)
				v0 = t6
				goto l2
			}
			if v3&i32(0x4000000) != 0 {
				goto l1
			}
			t3 := m.fn165(v0, v1)
			v0 = t3
			goto l2
		}
	l1:
		t7 := int64(load64(m.memory[uint32(v0):]))
		v4 = t7
		v0 = i32(17)
	l4:
		{
			t8 := int32(m.memory[int64(uint32(int32(v4)&i32(15)))+1122456])
			m.memory[uint32(v2+v0+i32(-2))] = byte(t8)
			v0 = v0 + i32(-1)
			v4 = int64(uint64(v4) >> 4)
			if v4 != i64(0) {
				goto l4
			}
		}
		t9 := m.fn683(v1, i32(1), i32(1122454), i32(2), v2+v0+i32(-1), i32(17)-v0)
		v0 = t9
	}
l2:
	m.g0 = v2 + i32(16)
	return v0
}
func (m *Module) fn856(v0, v1, v2 int32) {
	var v3, v4, v5, v6, v7, v8 int32
	t0 := m.g0
	v3 = t0 - i32(32)
	m.g0 = v3
	{
		switch v1 {
		case 34:
			goto l5
		case 39:
			if v2&i32(256) == 0 {
				goto l12
			}
			store64(m.memory[int64(uint32(v0))+2:], uint64(i64(0)))
			store16(m.memory[uint32(v0):], uint16(i32(10076)))
			goto l14
		default:
			if v1 == i32(92) {
				store64(m.memory[int64(uint32(v0))+2:], uint64(i64(0)))
				store16(m.memory[uint32(v0):], uint16(i32(23644)))
				goto l14
			}
			fallthrough
		case 1, 2, 3, 4, 5, 6, 7, 8, 11, 12, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 35, 36, 37, 38:
			if uint32(v1) < uint32(i32(768)) {
				goto l9
			}
			if v2&i32(1) != 0 {
				t1 := m.fn857(v1)
				if t1 == 0 {
					goto l13
				}
				m.memory[int64(uint32(v3))+14] = byte(i32(0))
				store16(m.memory[int64(uint32(v3))+12:], uint16(i32(0)))
				t2 := int32(m.memory[int64(uint32(int32(uint32(v1)>>20)))+1098720])
				m.memory[int64(uint32(v3))+15] = byte(t2)
				t3 := int32(m.memory[int64(uint32(int32(uint32(v1)>>4)&i32(15)))+1098720])
				m.memory[int64(uint32(v3))+19] = byte(t3)
				t4 := int32(m.memory[int64(uint32(int32(uint32(v1)>>8)&i32(15)))+1098720])
				m.memory[int64(uint32(v3))+18] = byte(t4)
				t5 := int32(m.memory[int64(uint32(int32(uint32(v1)>>12)&i32(15)))+1098720])
				m.memory[int64(uint32(v3))+17] = byte(t5)
				t6 := int32(m.memory[int64(uint32(int32(uint32(v1)>>16)&i32(15)))+1098720])
				m.memory[int64(uint32(v3))+16] = byte(t6)
				t7 := v3 + i32(12)
				v2 = int32(uint32(int32(bits.LeadingZeros32(uint32(v1|i32(1))))) >> 2)
				v4 = t7 + v2
				m.memory[uint32(v4)] = byte(i32(123))
				m.memory[uint32(v4+i32(-1))] = byte(i32(117))
				t8 := v3 + i32(12)
				v2 = v2 + i32(-2)
				m.memory[uint32(t8+v2)] = byte(i32(92))
				t9 := int64(load64(m.memory[int64(uint32(v3))+12:]))
				store64(m.memory[uint32(v0):], uint64(t9))
				m.memory[int64(uint32(v3))+21] = byte(i32(125))
				t10 := int32(m.memory[int64(uint32(v1&i32(15)))+1098720])
				m.memory[int64(uint32(v3))+20] = byte(t10)
				t11 := int32(load16(m.memory[int64(uint32(v3))+20:]))
				store16(m.memory[int64(uint32(v0))+8:], uint16(t11))
				goto l15
			}
		l9:
			if uint32(v1) < uint32(i32(32)) {
				goto l11
			}
			if uint32(v1) < uint32(i32(127)) {
				goto l12
			}
			goto l13
		case 0:
			store64(m.memory[int64(uint32(v0))+2:], uint64(i64(0)))
			store16(m.memory[uint32(v0):], uint16(i32(12380)))
			goto l14
		case 9:
			store64(m.memory[int64(uint32(v0))+2:], uint64(i64(0)))
			store16(m.memory[uint32(v0):], uint16(i32(29788)))
			goto l14
		case 13:
			store64(m.memory[int64(uint32(v0))+2:], uint64(i64(0)))
			store16(m.memory[uint32(v0):], uint16(i32(29276)))
			goto l14
		case 10:
			store64(m.memory[int64(uint32(v0))+2:], uint64(i64(0)))
			store16(m.memory[uint32(v0):], uint16(i32(28252)))
			goto l14
		}
	l5:
		if uint32(v2&i32(0xffffff)) < uint32(i32(65536)) {
			goto l12
		}
		store64(m.memory[int64(uint32(v0))+2:], uint64(i64(0)))
		store16(m.memory[uint32(v0):], uint16(i32(8796)))
	l14:
		v1 = i32(2)
		v2 = i32(0)
		goto l16
	l13:
		if uint32(v1) < uint32(i32(65536)) {
			v5 = i32(0)
			v6 = int32(uint32(v1)>>8) & i32(255)
			v2 = i32(0)
		l21:
			{
				v7 = v2 + i32(2)
				t12 := int32(m.memory[int64(uint32(v2))+1101405])
				t13 := v5
				v4 = t12
				v8 = t13 + v4
				{
					t14 := int32(m.memory[int64(uint32(v2))+1101404])
					v2 = t14
					if v2 == v6 {
						if uint32(v8) < uint32(v5) {
							goto l22
						}
						if uint32(v8) > uint32(i32(284)) {
							goto l22
						}
						if v4 == 0 {
							goto l23
						}
						v2 = v5 + i32(1101480)
						goto l25
					l22:
						m.fn124(v5, v8, i32(284), i32(1102056))
						panic("unreachable")
					l25:
						{
							t15 := int32(m.memory[uint32(v2)])
							if t15 == v1&i32(255) {
								goto l11
							}
							v2 = v2 + i32(1)
							v4 = v4 + i32(-1)
							if v4 != 0 {
								goto l25
							}
						}
					l23:
						v5 = v8
						v2 = v7
						if v7 != i32(76) {
							goto l21
						}
						goto l20
					}
					if uint32(v2) > uint32(v6) {
						goto l20
					}
					v5 = v8
					v2 = v7
					if v7 != i32(76) {
						goto l21
					}
					goto l20
				}
			}
		}
		if uint32(v1) < uint32(i32(0x20000)) {
			v5 = i32(0)
			v6 = int32(uint32(v1)>>8) & i32(255)
			v2 = i32(0)
		l28:
			{
				v7 = v2 + i32(2)
				t16 := int32(m.memory[int64(uint32(v2))+1100597])
				t17 := v5
				v4 = t16
				v8 = t17 + v4
				{
					t18 := int32(m.memory[int64(uint32(v2))+1100596])
					v2 = t18
					if v2 == v6 {
						goto l26
					}
					if uint32(v2) > uint32(v6) {
						goto l27
					}
					v5 = v8
					v2 = v7
					if v7 != i32(92) {
						goto l28
					}
					goto l27
				}
			l26:
				if uint32(v8) < uint32(v5) {
					goto l29
				}
				if uint32(v8) > uint32(i32(212)) {
					goto l29
				}
				if v4 == 0 {
					goto l30
				}
				v2 = v5 + i32(1100688)
				goto l32
			l29:
				m.fn124(v5, v8, i32(212), i32(1102056))
				panic("unreachable")
			l32:
				{
					t19 := int32(m.memory[uint32(v2)])
					if t19 == v1&i32(255) {
						goto l11
					}
					v2 = v2 + i32(1)
					v4 = v4 + i32(-1)
					if v4 != 0 {
						goto l32
					}
				}
			l30:
				v5 = v8
				v2 = v7
				if v7 != i32(92) {
					goto l28
				}
			}
		l27:
			v5 = v1 & i32(0xffff)
			v4 = i32(1)
			v2 = i32(0)
		l37:
			v7 = v2 + i32(1)
			{
				t20 := int32(int8(m.memory[int64(uint32(v2))+1100900]))
				v8 = t20
				if v8 < i32(0) {
					if v7 == i32(504) {
						m.fn222(i32(1102072))
						panic("unreachable")
					}
					t21 := int32(m.memory[uint32(v2+i32(1100901))])
					v8 = v8&i32(127)<<8 | t21
					v2 = v2 + i32(2)
					goto l34
				}
				v2 = v7
				goto l34
			}
		l34:
			v5 = v5 - v8
			if v5 < i32(0) {
				goto l36
			}
			v4 = v4 ^ i32(1)
			if v2 != i32(504) {
				goto l37
			}
			goto l36
		}
		v2 = v1 & i32(0x1ffffe)
		if v2 == i32(183982) {
			goto l11
		}
		if v1&i32(0x1fffe0) == i32(173792) {
			goto l11
		}
		if v2 == i32(178206) {
			goto l11
		}
		if uint32(v1+i32(-191472)) > uint32(i32(-16)) {
			goto l11
		}
		if uint32(v1+i32(-194560)) > uint32(i32(-2467)) {
			goto l11
		}
		if uint32(v1+i32(-0x30000)) > uint32(i32(-1507)) {
			goto l11
		}
		if uint32(v1+i32(-201552)) > uint32(i32(-6)) {
			goto l11
		}
		if uint32(v1+i32(-917760)) > uint32(i32(-707719)) {
			goto l11
		}
		if uint32(v1) < uint32(i32(918000)) {
			goto l12
		}
		goto l11
	l20:
		v4 = i32(1)
		v5 = v1
		v2 = i32(0)
	l41:
		v7 = v2 + i32(1)
		{
			t22 := int32(int8(m.memory[int64(uint32(v2))+1101764]))
			v8 = t22
			if v8 < i32(0) {
				if v7 == i32(292) {
					m.fn222(i32(1102072))
					panic("unreachable")
				}
				t23 := int32(m.memory[uint32(v2+i32(1101765))])
				v8 = v8&i32(127)<<8 | t23
				v2 = v2 + i32(2)
				goto l39
			}
			v2 = v7
			goto l39
		}
	l39:
		v5 = v5 - v8
		if v5 < i32(0) {
			goto l36
		}
		v4 = v4 ^ i32(1)
		if v2 != i32(292) {
			goto l41
		}
	l36:
		if v4&i32(1) == 0 {
			goto l11
		}
	l12:
		store32(m.memory[uint32(v0):], uint32(v1))
		v1 = i32(129)
		v2 = i32(128)
		goto l16
	l11:
		m.memory[int64(uint32(v3))+24] = byte(i32(0))
		store16(m.memory[int64(uint32(v3))+22:], uint16(i32(0)))
		t24 := int32(m.memory[int64(uint32(int32(uint32(v1)>>20)))+1098720])
		m.memory[int64(uint32(v3))+25] = byte(t24)
		t25 := int32(m.memory[int64(uint32(int32(uint32(v1)>>4)&i32(15)))+1098720])
		m.memory[int64(uint32(v3))+29] = byte(t25)
		t26 := int32(m.memory[int64(uint32(int32(uint32(v1)>>8)&i32(15)))+1098720])
		m.memory[int64(uint32(v3))+28] = byte(t26)
		t27 := int32(m.memory[int64(uint32(int32(uint32(v1)>>12)&i32(15)))+1098720])
		m.memory[int64(uint32(v3))+27] = byte(t27)
		t28 := int32(m.memory[int64(uint32(int32(uint32(v1)>>16)&i32(15)))+1098720])
		m.memory[int64(uint32(v3))+26] = byte(t28)
		t29 := v3 + i32(22)
		v2 = int32(uint32(int32(bits.LeadingZeros32(uint32(v1|i32(1))))) >> 2)
		v4 = t29 + v2
		m.memory[uint32(v4)] = byte(i32(123))
		m.memory[uint32(v4+i32(-1))] = byte(i32(117))
		t30 := v3 + i32(22)
		v2 = v2 + i32(-2)
		m.memory[uint32(t30+v2)] = byte(i32(92))
		t31 := int64(load64(m.memory[int64(uint32(v3))+22:]))
		store64(m.memory[uint32(v0):], uint64(t31))
		m.memory[int64(uint32(v3))+31] = byte(i32(125))
		t32 := int32(m.memory[int64(uint32(v1&i32(15)))+1098720])
		m.memory[int64(uint32(v3))+30] = byte(t32)
		t33 := int32(load16(m.memory[int64(uint32(v3))+30:]))
		store16(m.memory[int64(uint32(v0))+8:], uint16(t33))
	}
l15:
	v1 = i32(10)
l16:
	m.memory[int64(uint32(v0))+13] = byte(v1)
	m.memory[int64(uint32(v0))+12] = byte(v2)
	m.g0 = v3 + i32(32)
}
func (m *Module) fn857(v0 int32) int32 {
	var v1, v2, v3, v4, v5 int32
	v1 = i32(0)
	p0 := i32(16)
	if uint32(v0) < uint32(i32(69291)) {
		p0 = i32(0)
	}
	v2 = p0
	t1 := v2
	v2 = v2 | i32(8)
	t2 := int32(load32(m.memory[int64(uint32(v2<<2))+1105772:]))
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
	t7 := int32(load32(m.memory[int64(uint32(v3<<2))+1105772:]))
	p8 := v3
	if uint32(t7<<11) > uint32(v2) {
		p8 = t6
	}
	v3 = p8
	t9 := v3
	v3 = v3 | i32(2)
	t10 := int32(load32(m.memory[int64(uint32(v3<<2))+1105772:]))
	p11 := v3
	if uint32(t10<<11) > uint32(v2) {
		p11 = t9
	}
	v3 = p11
	t12 := v3
	v3 = v3 + i32(1)
	t13 := int32(load32(m.memory[int64(uint32(v3<<2))+1105772:]))
	p14 := v3
	if uint32(t13<<11) > uint32(v2) {
		p14 = t12
	}
	v3 = p14
	t15 := v3
	v3 = v3 + i32(1)
	t16 := int32(load32(m.memory[int64(uint32(v3<<2))+1105772:]))
	p17 := v3
	if uint32(t16<<11) > uint32(v2) {
		p17 = t15
	}
	v3 = p17
	t18 := int32(load32(m.memory[int64(uint32(v3<<2))+1105772:]))
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
	v5 = v2 + i32(1105772)
	t21 := int32(load32(m.memory[int64(uint32(v2))+1105772:]))
	v2 = int32(uint32(t21) >> 21)
	v4 = i32(767)
	{
		{
			if uint32(v3) > uint32(i32(31)) {
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
		t24 := int32(m.memory[uint32(v2+i32(1097258))])
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
func (m *Module) fn858(v0, v1, v2 int32) int32 {
	var v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14 int32
	var v15, v16 int64
	t0 := m.g0
	v3 = t0 - i32(160)
	m.g0 = v3
	v4 = i32(0)
	memory_zero(m.memory, uint32(v3), uint32(i32(160)))
	{
		t1 := int32(load32(m.memory[int64(uint32(v0))+160:]))
		v5 = t1
		if uint32(v5) < uint32(v2) {
			v6 = v2 + i32(1)
			v7 = v2 << 2
			v14 = v0 + v5<<2
			v11 = v0
			v10 = i32(0)
		l14:
			v9 = v3 + v4<<2
		l9:
			{
				v13 = v4
				v12 = v9
				if v11 == v14 {
					goto l2
				}
				v9 = v12 + i32(4)
				v4 = v13 + i32(1)
				t8 := int32(load32(m.memory[uint32(v11):]))
				v8 = t8
				v5 = v11 + i32(4)
				v11 = v5
				if v8 == 0 {
					goto l9
				}
			}
			v15 = int64(uint32(v8))
			v16 = i64(0)
			v8 = v7
			v11 = v13
			v9 = v1
		l11:
			{
				if uint32(v11) >= uint32(i32(40)) {
					m.fn36(v11, i32(40), i32(1099776))
					panic("unreachable")
				}
				t9 := int64(load32(m.memory[uint32(v12):]))
				t10 := int64(load32(m.memory[uint32(v9):]))
				t11 := v12
				v16 = v16 + t9 + t10*v15
				store32(m.memory[uint32(t11):], uint32(v16))
				v16 = int64(uint64(v16) >> 32)
				v12 = v12 + i32(4)
				v11 = v11 + i32(1)
				v9 = v9 + i32(4)
				v8 = v8 + i32(-4)
				if v8 != 0 {
					goto l11
				}
			}
			v12 = v2
			{
				if v16 == 0 {
					goto l12
				}
				v12 = v13 + v2
				if uint32(v12) >= uint32(i32(40)) {
					m.fn36(v12, i32(40), i32(1099776))
					panic("unreachable")
				}
				store32(m.memory[uint32(v3+v12<<2):], uint32(int32(v16)))
				v12 = v6
			l12:
				t12 := v10
				v12 = v12 + v13
				p13 := v12
				if uint32(v10) > uint32(v12) {
					p13 = t12
				}
				v10 = p13
				v11 = v5
				goto l14
			}
		}
		if uint32(v5) >= uint32(i32(41)) {
			m.fn124(i32(0), v5, i32(40), i32(1099776))
			panic("unreachable")
		}
		v6 = v5 + i32(1)
		v7 = v5 << 2
		v8 = v1 + v2<<2
		v9 = i32(0)
		v10 = i32(0)
	l8:
		v11 = v3 + v9<<2
	l3:
		{
			v4 = v9
			v12 = v11
			if v1 == v8 {
				goto l2
			}
			v11 = v12 + i32(4)
			v9 = v4 + i32(1)
			t2 := int32(load32(m.memory[uint32(v1):]))
			v13 = t2
			v14 = v1 + i32(4)
			v1 = v14
			if v13 == 0 {
				goto l3
			}
		}
		v15 = int64(uint32(v13))
		v16 = i64(0)
		v13 = v7
		v1 = v4
		v11 = v0
	l5:
		{
			if uint32(v1) >= uint32(i32(40)) {
				m.fn36(v1, i32(40), i32(1099776))
				panic("unreachable")
			}
			t3 := int64(load32(m.memory[uint32(v12):]))
			t4 := int64(load32(m.memory[uint32(v11):]))
			t5 := v12
			v16 = v16 + t3 + t4*v15
			store32(m.memory[uint32(t5):], uint32(v16))
			v16 = int64(uint64(v16) >> 32)
			v12 = v12 + i32(4)
			v1 = v1 + i32(1)
			v11 = v11 + i32(4)
			v13 = v13 + i32(-4)
			if v13 != 0 {
				goto l5
			}
		}
		v12 = v5
		{
			if v16 == 0 {
				goto l6
			}
			v12 = v4 + v5
			if uint32(v12) >= uint32(i32(40)) {
				m.fn36(v12, i32(40), i32(1099776))
				panic("unreachable")
			}
			store32(m.memory[uint32(v3+v12<<2):], uint32(int32(v16)))
			v12 = v6
		l6:
			t6 := v10
			v12 = v12 + v4
			p7 := v12
			if uint32(v10) > uint32(v12) {
				p7 = t6
			}
			v10 = p7
			v1 = v14
			goto l8
		}
	}
l2:
	memory_copy(m.memory, uint32(v0), uint32(v3), uint32(i32(160)))
	store32(m.memory[int64(uint32(v0))+160:], uint32(v10))
	m.g0 = v3 + i32(160)
	return v0
}
func (m *Module) fn859(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7, v8, v9, v10 int32
	if uint32(v1) >= uint32(i32(1280)) {
		m.fn7(i32(1099792), i32(29), i32(1099776))
		panic("unreachable")
	}
	v2 = int32(uint32(v1) >> 5)
	{
		{
			t0 := int32(load32(m.memory[int64(uint32(v0))+160:]))
			v3 = t0
			if v3 == 0 {
				goto l1
			}
			v4 = v3 + i32(-1)
			v5 = v3<<2 + v0 + i32(-4)
			v6 = (v3+v2)<<2 + v0 + i32(-4)
			var p1 int32
			if uint32(v3) < uint32(i32(41)) {
				p1 = 1
			}
			v3 = p1
		l4:
			{
				if v3 == 0 {
					m.fn36(v4, i32(40), i32(1099776))
					panic("unreachable")
				}
				v7 = v2 + v4
				if uint32(v7) >= uint32(i32(40)) {
					m.fn36(v7, i32(40), i32(1099776))
					panic("unreachable")
				}
				t2 := int32(load32(m.memory[uint32(v5):]))
				store32(m.memory[uint32(v6):], uint32(t2))
				v6 = v6 + i32(-4)
				v5 = v5 + i32(-4)
				v4 = v4 + i32(-1)
				if v4 != i32(-1) {
					goto l4
				}
			}
		}
	l1:
		v5 = v1 & i32(31)
		if v2 == 0 {
			goto l5
		}
		v4 = v2 << 2
		if v4 == 0 {
			goto l5
		}
		memory_zero(m.memory, uint32(v0), uint32(v4))
	l5:
		t3 := int32(load32(m.memory[int64(uint32(v0))+160:]))
		v7 = t3
		v4 = v7 + v2
		if v5 != 0 {
			v6 = v4 + i32(-1)
			if uint32(v6) > uint32(i32(39)) {
				m.fn36(v6, i32(40), i32(1099776))
				panic("unreachable")
			}
			v8 = v4
			t4 := int32(load32(m.memory[uint32(v0+v6<<2):]))
			v6 = i32(32) - v5
			v3 = i32_shr_u(t4, v6)
			if v3 == 0 {
				goto l8
			}
			if uint32(v4) > uint32(i32(39)) {
				m.fn36(v4, i32(40), i32(1099776))
				panic("unreachable")
			}
			store32(m.memory[uint32(v0+v4<<2):], uint32(v3))
			v8 = v4 + i32(1)
			goto l8
		}
		store32(m.memory[int64(uint32(v0))+160:], uint32(v4))
		return v0
	}
l8:
	v9 = v2 + i32(1)
	if uint32(v9) >= uint32(v4) {
		goto l10
	}
	{
		if v7&i32(1) == 0 {
			goto l11
		}
		v3 = v4
		goto l12
	l11:
		t5 := v0
		v3 = v4 + i32(-1)
		v1 = t5 + v3<<2
		t6 := int32(load32(m.memory[uint32(v0+v4<<2+i32(-8)):]))
		t7 := int32(load32(m.memory[uint32(v1):]))
		store32(m.memory[uint32(v1):], uint32(i32_shr_u(t6, v6)|i32_shl(t7, v5)))
	}
l12:
	if v7 == i32(2) {
		goto l10
	}
	v4 = v3<<2 + v0 + i32(-12)
l13:
	{
		v7 = v4 + i32(8)
		t8 := v7
		v1 = v4 + i32(4)
		t9 := int32(load32(m.memory[uint32(v1):]))
		v10 = t9
		t10 := int32(load32(m.memory[uint32(v7):]))
		store32(m.memory[uint32(t8):], uint32(i32_shr_u(v10, v6)|i32_shl(t10, v5)))
		t11 := int32(load32(m.memory[uint32(v4):]))
		store32(m.memory[uint32(v1):], uint32(i32_shr_u(t11, v6)|i32_shl(v10, v5)))
		v4 = v4 + i32(-8)
		t12 := v9
		v3 = v3 + i32(-2)
		if uint32(t12) < uint32(v3) {
			goto l13
		}
	}
l10:
	v4 = v0 + v2<<2
	t13 := int32(load32(m.memory[uint32(v4):]))
	store32(m.memory[uint32(v4):], uint32(i32_shl(t13, v5)))
	store32(m.memory[int64(uint32(v0))+160:], uint32(v8))
	return v0
}
func (m *Module) fn860(v0, v1 int32) {
	var v2, v3, v4, v5, v6, v7, v8, v9, v10 int32
	var v11, v12, v13, v14 int64
	{
		t0 := int32(load32(m.memory[uint32(v0):]))
		v2 = t0
		if v2 == 0 {
			return
		}
		v3 = v1 & i32(63)
		v1 = v3 << 1
		t1 := int32(load16(m.memory[int64(uint32(v1))+1107624:]))
		v4 = t1
		v5 = v4 & i32(2047)
		t2 := int32(load16(m.memory[int64(uint32(v1))+1107626:]))
		v6 = v5 - t2&i32(2047)
		v7 = i32(0) - v2
		v8 = v0 + i32(8)
		v4 = int32(uint32(v4) >> 11)
		v1 = i32(-1308)
		{
		l4:
			{
				if v6+v1 == i32(-1308) {
					goto l1
				}
				v9 = v5 + v1
				if v9 == 0 {
					goto l1
				}
				if v7+v1 == i32(-1308) {
					v4 = v4 + i32(-1)
					goto l1
				}
				if v1 == i32(-540) {
					m.fn36(i32(768), i32(768), i32(1109064))
					panic("unreachable")
				}
				v10 = v8 + v1
				v1 = v1 + i32(1)
				t3 := int32(m.memory[uint32(v10+i32(1308))])
				v10 = t3
				t4 := int32(m.memory[uint32(v9+i32(1109062))])
				t5 := v10
				v9 = t4
				if t5 == v9&i32(255) {
					goto l4
				}
			}
			t6 := v4
			var p7 int32
			if uint32(v10) < uint32(v9&i32(255)) {
				p7 = 1
			}
			v4 = t6 - p7
			goto l1
		}
	l1:
		v9 = v0 + i32(7)
		v10 = v9 + v4
		v11 = int64(uint32(v3))
		v12 = i64(0)
	l8:
		v1 = v2
		v2 = v1 + i32(-1)
		{
			if uint32(v1) >= uint32(i32(769)) {
				m.fn36(v2, i32(768), i32(1099864))
				panic("unreachable")
			}
			t8 := int64(m.memory[uint32(v9+v1)])
			v13 = i64_shl(t8, v11) + v12
			t9 := int64(uint64(v13) / uint64(i64(10)))
			t10 := v13
			v12 = t9
			v14 = t10 + v12*i64(-10)
			if uint32(v2+v4) < uint32(i32(768)) {
				goto l6
			}
			if v14 == 0 {
				goto l7
			}
			m.memory[int64(uint32(v0))+776] = byte(i32(1))
			goto l7
		}
	l6:
		m.memory[uint32(v10+v1)] = byte(v14)
	l7:
		if v2 != 0 {
			goto l8
		}
		if uint64(v13) < uint64(i64(10)) {
			goto l9
		}
		v2 = v4 + i32(7)
	l12:
		{
			v13 = v12
			t11 := int64(uint64(v13) / uint64(i64(10)))
			t12 := v13
			v12 = t11
			v14 = t12 + v12*i64(-10)
			if uint32(v2+i32(-8)) < uint32(i32(768)) {
				goto l10
			}
			if v14 == 0 {
				goto l11
			}
			m.memory[int64(uint32(v0))+776] = byte(i32(1))
			goto l11
		l10:
			m.memory[uint32(v0+v2)] = byte(v14)
		l11:
			v2 = v2 + i32(-1)
			if uint64(v13) >= uint64(i64(10)) {
				goto l12
			}
		}
	l9:
		t13 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		store32(m.memory[int64(uint32(v0))+4:], uint32(t13+v4))
		t14 := int32(load32(m.memory[uint32(v0):]))
		t15 := v0
		v1 = t14 + v4
		p16 := i32(768)
		if uint32(v1) < uint32(i32(768)) {
			p16 = v1
		}
		v2 = p16
		store32(m.memory[uint32(t15):], uint32(v2))
		if v1 == 0 {
			return
		}
	l15:
		v1 = v2 + i32(-1)
		{
			if uint32(v2) > uint32(i32(768)) {
				m.fn36(v1, i32(768), i32(1099848))
				panic("unreachable")
			}
			t17 := int32(m.memory[uint32(v0+v2+i32(7))])
			if t17 == 0 {
				goto l14
			}
			return
		}
	l14:
		store32(m.memory[uint32(v0):], uint32(v1))
		v2 = v1
		if v1 != 0 {
			goto l15
		}
	}
}
func (m *Module) fn861(v0, v1 int32) {
	var v2, v3, v4 int32
	var v5, v6 int64
	var v7, v8 int32
	var v9 int64
	var v10, v11 int32
	var v12 int64
	v2 = v0 + i32(8)
	t0 := int32(load32(m.memory[uint32(v0):]))
	v3 = t0
	v4 = i32(0) - v3
	v5 = int64(uint32(v1 & i32(63)))
	v1 = i32(-768)
	v6 = i64(0)
	{
	l3:
		{
			v7 = v4 + v1
			if v7 == i32(-768) {
				goto l0
			}
			if v1 == 0 {
				m.fn36(i32(768), i32(768), i32(1099880))
				panic("unreachable")
			}
			t1 := v6 * i64(10)
			v8 = v0 + v1
			t2 := int64(m.memory[uint32(v8+i32(776))])
			v6 = t1 + t2
			if i64_shr_u(v6, v5) != i64(0) {
				v8 = v1 + i32(769)
				goto l4
			}
			if v7 == i32(-769) {
				goto l0
			}
			v1 = v1 + i32(2)
			t3 := int64(m.memory[uint32(v8+i32(777))])
			v6 = v6*i64(10) + t3
			if i64_shr_u(v6, v5) == 0 {
				goto l3
			}
		}
		v8 = v1 + i32(768)
		goto l4
	l0:
		if v6 == 0 {
			return
		}
		if i64_shr_u(v6, v5) == i64(0) {
			goto l6
		}
		v8 = v3
		goto l4
	l6:
		v8 = v3
	l7:
		v8 = v8 + i32(1)
		v6 = v6 * i64(10)
		if i64_shr_u(v6, v5) == 0 {
			goto l7
		}
	l4:
		t4 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		t5 := v0
		v1 = t4 - v8 + i32(1)
		store32(m.memory[int64(uint32(t5))+4:], uint32(v1))
		{
			if v1 < i32(-2047) {
				goto l8
			}
			v9 = i64_shl(i64(-1), v5) ^ i64(-1)
			v1 = i32(0)
			{
				if uint32(v3) <= uint32(v8) {
					goto l9
				}
				v7 = i32(0)
				v1 = i32(768) - v8
				p6 := v1
				if uint32(v1) > uint32(i32(768)) {
					p6 = i32(0)
				}
				v4 = p6
				v10 = v8 - v3
				v11 = v2 + v8
				v1 = v3 - v8
			l11:
				{
					if v4 != v7 {
						goto l10
					}
					m.fn36(v8+v7, i32(768), i32(1099896))
					panic("unreachable")
				l10:
					t7 := int64(m.memory[uint32(v11+v7)])
					v12 = t7
					m.memory[uint32(v2+v7)] = byte(i64_shr_u(v6, v5))
					v6 = v12 + v6&v9*i64(10)
					t8 := v10
					v7 = v7 + i32(1)
					if t8+v7 != 0 {
						goto l11
					}
				}
			}
		l9:
			if v6 == 0 {
				goto l12
			}
		l15:
			v12 = v6
			v6 = v12 & v9 * i64(10)
			v7 = int32(i64_shr_u(v12, v5))
			if uint32(v1) < uint32(i32(768)) {
				goto l13
			}
			if v7&i32(255) == 0 {
				goto l14
			}
			m.memory[int64(uint32(v0))+776] = byte(i32(1))
			goto l14
		l13:
			m.memory[uint32(v2+v1)] = byte(v7)
			v1 = v1 + i32(1)
		l14:
			if !(v6 == 0) {
				goto l15
			}
		l12:
			;
			var p9 int32
			if uint32(v1) > uint32(i32(768)) {
				p9 = 1
			}
			v2 = p9
		l17:
			store32(m.memory[uint32(v0):], uint32(v1))
			if v1 == 0 {
				return
			}
			v7 = v1 + i32(-1)
			{
				if v2 != 0 {
					m.fn36(v7, i32(768), i32(1099848))
					panic("unreachable")
				}
				v8 = v0 + v1
				v1 = v7
				t10 := int32(m.memory[uint32(v8+i32(7))])
				if t10 == 0 {
					goto l17
				}
				return
			}
		}
	l8:
		m.memory[int64(uint32(v0))+776] = byte(i32(0))
		store64(m.memory[uint32(v0):], uint64(i64(0)))
	}
}
func (m *Module) fn862(v0 int32) int64 {
	var v1 int64
	var v2, v3, v4, v5, v6, v7 int32
	v1 = i64(0)
	{
		t0 := int32(load32(m.memory[uint32(v0):]))
		v2 = t0
		if v2 == 0 {
			goto l0
		}
		t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		v3 = t1
		if v3 < i32(0) {
			goto l0
		}
		v1 = i64(-1)
		if uint32(v3) > uint32(i32(18)) {
			goto l0
		}
		{
			if v3 != 0 {
				goto l1
			}
			v1 = i64(0)
			goto l2
		l1:
			if v3 != i32(1) {
				goto l3
			}
			v4 = i32(0)
			v1 = i64(0)
			goto l4
		l3:
			v5 = v3 & i32(1)
			v6 = v3 & i32(30)
			v7 = i32(0)
			v1 = i64(0)
		l7:
			v1 = v1 * i64(10)
			{
				v4 = v7
				if uint32(v4) >= uint32(v2) {
					goto l5
				}
				t2 := int64(m.memory[uint32(v0+v4+i32(8))])
				v1 = v1 + t2
			}
		l5:
			v1 = v1 * i64(10)
			{
				v7 = v4 + i32(1)
				if uint32(v7) >= uint32(v2) {
					goto l6
				}
				t3 := int64(m.memory[uint32(v0+v4+i32(9))])
				v1 = v1 + t3
			}
		l6:
			v7 = v7 + i32(1)
			if v7 != v6 {
				goto l7
			}
			if v5 == 0 {
				goto l2
			}
			v4 = v4 + i32(2)
		l4:
			v1 = v1 * i64(10)
			if uint32(v4) >= uint32(v2) {
				goto l2
			}
			t4 := int64(m.memory[uint32(v0+i32(8)+v4)])
			v1 = v1 + t4
		}
	l2:
		if uint32(v3) >= uint32(v2) {
			goto l0
		}
		v7 = v0 + v3
		t5 := int32(m.memory[int64(uint32(v7))+8])
		v4 = t5
		{
			if v3+i32(1) != v2 {
				goto l8
			}
			if v4&i32(255) == i32(5) {
				goto l9
			}
		l8:
			if uint32(v4&i32(255)) > uint32(i32(4)) {
				goto l10
			}
			goto l0
		l9:
			t6 := int32(m.memory[int64(uint32(v0))+776])
			if t6 != 0 {
				goto l10
			}
			if v3 == 0 {
				goto l0
			}
			t7 := int32(m.memory[uint32(v7+i32(8)+i32(-1))])
			if t7&i32(1) == 0 {
				goto l0
			}
		}
	l10:
		v1 = v1 + i64(1)
	}
l0:
	return v1
}
func (m *Module) fn863(v0, v1, v2, v3, v4 int32) {
	var v5, v6 int32
	var v7 int64
	t0 := m.g0
	v5 = t0 - i32(48)
	m.g0 = v5
	store32(m.memory[int64(uint32(v5))+4:], uint32(v3))
	store32(m.memory[uint32(v5):], uint32(v2))
	store32(m.memory[int64(uint32(v5))+8:], uint32(v1))
	{
		{
			if uint32(v2) > uint32(v1) {
				t5 := v5
				v7 = int64(uint32(i32(3))) << 32
				store64(m.memory[int64(uint32(t5))+32:], uint64(v7|int64(uint32(v5+i32(8)))))
				store64(m.memory[int64(uint32(v5))+24:], uint64(v7|int64(uint32(v5))))
				m.fn31(i32(1050379), v5+i32(24), v4)
				panic("unreachable")
			}
			if uint32(v3) > uint32(v1) {
				t6 := v5
				v7 = int64(uint32(i32(3))) << 32
				store64(m.memory[int64(uint32(t6))+32:], uint64(v7|int64(uint32(v5+i32(8)))))
				store64(m.memory[int64(uint32(v5))+24:], uint64(v7|int64(uint32(v5+i32(4)))))
				m.fn31(i32(1050440), v5+i32(24), v4)
				panic("unreachable")
			}
			if uint32(v2) > uint32(v3) {
				t7 := v5
				v7 = int64(uint32(i32(3))) << 32
				store64(m.memory[int64(uint32(t7))+32:], uint64(v7|int64(uint32(v5+i32(4)))))
				store64(m.memory[int64(uint32(v5))+24:], uint64(v7|int64(uint32(v5))))
				m.fn31(i32(1049796), v5+i32(24), v4)
				panic("unreachable")
			}
			if v2 == 0 {
				goto l3
			}
			if uint32(v2) >= uint32(v1) {
				goto l3
			}
			t1 := int32(int8(m.memory[uint32(v0+v2)]))
			if t1 > i32(-65) {
				goto l3
			}
			v6 = v2
		l5:
			{
				t2 := int32(int8(m.memory[uint32(v0+v6)]))
				if t2 > i32(-65) {
					goto l7
				}
				v6 = v6 + i32(-1)
				if v6 != 0 {
					goto l5
				}
			}
			v6 = i32(0)
		l7:
			{
				t3 := int32(int8(m.memory[uint32(v0+v2)]))
				if t3 > i32(-65) {
					goto l6
				}
				t4 := v1
				v2 = v2 + i32(1)
				if t4 != v2 {
					goto l7
				}
			}
			v2 = v1
			goto l6
		}
	l6:
		store32(m.memory[int64(uint32(v5))+12:], uint32(v6))
		store32(m.memory[int64(uint32(v5))+16:], uint32(v2))
		if uint32(v6) > uint32(v2) {
			goto l8
		}
		{
			if v6 == 0 {
				goto l9
			}
			if uint32(v6) < uint32(v1) {
				goto l10
			}
			if v6 == v1 {
				goto l9
			}
			goto l8
		l10:
			t8 := int32(int8(m.memory[uint32(v0+v6)]))
			if t8 < i32(-64) {
				goto l8
			}
		}
	l9:
		{
			if uint32(v2) < uint32(v1) {
				goto l11
			}
			if v2 != v1 {
				goto l8
			}
			goto l12
		l11:
			t9 := int32(int8(m.memory[uint32(v0+v2)]))
			if t9 <= i32(-65) {
				goto l8
			}
		}
	l12:
		if v6 == v2 {
			goto l13
		}
		{
			{
				v0 = v0 + v6
				t10 := int32(int8(m.memory[uint32(v0)]))
				v6 = t10
				if v6 <= i32(-1) {
					goto l14
				}
				v6 = v6 & i32(255)
				goto l15
			}
		l14:
			t11 := int32(m.memory[int64(uint32(v0))+1])
			v3 = t11 & i32(63)
			v2 = v6 & i32(31)
			if uint32(v6) > uint32(i32(-33)) {
				goto l16
			}
			v6 = v2<<6 | v3
			goto l15
		l16:
			t12 := int32(m.memory[int64(uint32(v0))+2])
			v3 = v3<<6 | t12&i32(63)
			if uint32(v6) >= uint32(i32(-16)) {
				goto l17
			}
			v6 = v3 | v2<<12
			goto l15
		l17:
			t13 := int32(m.memory[int64(uint32(v0))+3])
			v6 = v3<<6 | t13&i32(63) | v2<<18&i32(0x1c0000)
		}
	l15:
		store32(m.memory[int64(uint32(v5))+20:], uint32(v6))
		store64(m.memory[int64(uint32(v5))+40:], uint64(int64(uint32(i32(89)))<<32|int64(uint32(v5+i32(12)))))
		store64(m.memory[int64(uint32(v5))+32:], uint64(int64(uint32(i32(90)))<<32|int64(uint32(v5+i32(20)))))
		store64(m.memory[int64(uint32(v5))+24:], uint64(int64(uint32(i32(3)))<<32|int64(uint32(v5))))
		m.fn31(i32(1066016), v5+i32(24), v4)
		panic("unreachable")
	l8:
		m.fn41(v0, v1, v6, v2, v4)
		panic("unreachable")
	l3:
		if v3 == 0 {
			goto l18
		}
		if uint32(v3) >= uint32(v1) {
			goto l18
		}
		t14 := int32(int8(m.memory[uint32(v0+v3)]))
		if t14 > i32(-65) {
			goto l18
		}
		v6 = v3
	l20:
		{
			t15 := int32(int8(m.memory[uint32(v0+v6)]))
			if t15 > i32(-65) {
				goto l22
			}
			v6 = v6 + i32(-1)
			if v6 != 0 {
				goto l20
			}
		}
		v6 = i32(0)
	l22:
		{
			t16 := int32(int8(m.memory[uint32(v0+v3)]))
			if t16 > i32(-65) {
				goto l21
			}
			t17 := v1
			v3 = v3 + i32(1)
			if t17 != v3 {
				goto l22
			}
		}
		v3 = v1
	l21:
		store32(m.memory[int64(uint32(v5))+12:], uint32(v6))
		store32(m.memory[int64(uint32(v5))+16:], uint32(v3))
		if uint32(v6) > uint32(v3) {
			goto l23
		}
		{
			if v6 == 0 {
				goto l24
			}
			if uint32(v6) < uint32(v1) {
				goto l25
			}
			if v6 == v1 {
				goto l24
			}
			goto l23
		l25:
			t18 := int32(int8(m.memory[uint32(v0+v6)]))
			if t18 < i32(-64) {
				goto l23
			}
		}
	l24:
		{
			if uint32(v3) < uint32(v1) {
				goto l26
			}
			if v3 != v1 {
				goto l23
			}
			goto l27
		l26:
			t19 := int32(int8(m.memory[uint32(v0+v3)]))
			if t19 <= i32(-65) {
				goto l23
			}
		}
	l27:
		if v6 == v3 {
			goto l13
		}
		{
			{
				v0 = v0 + v6
				t20 := int32(int8(m.memory[uint32(v0)]))
				v6 = t20
				if v6 <= i32(-1) {
					goto l28
				}
				v6 = v6 & i32(255)
				goto l29
			}
		l28:
			t21 := int32(m.memory[int64(uint32(v0))+1])
			v3 = t21 & i32(63)
			v2 = v6 & i32(31)
			if uint32(v6) > uint32(i32(-33)) {
				goto l30
			}
			v6 = v2<<6 | v3
			goto l29
		l30:
			t22 := int32(m.memory[int64(uint32(v0))+2])
			v3 = v3<<6 | t22&i32(63)
			if uint32(v6) >= uint32(i32(-16)) {
				goto l31
			}
			v6 = v3 | v2<<12
			goto l29
		l31:
			t23 := int32(m.memory[int64(uint32(v0))+3])
			v6 = v3<<6 | t23&i32(63) | v2<<18&i32(0x1c0000)
		}
	l29:
		store32(m.memory[int64(uint32(v5))+20:], uint32(v6))
		store64(m.memory[int64(uint32(v5))+40:], uint64(int64(uint32(i32(89)))<<32|int64(uint32(v5+i32(12)))))
		store64(m.memory[int64(uint32(v5))+32:], uint64(int64(uint32(i32(90)))<<32|int64(uint32(v5+i32(20)))))
		store64(m.memory[int64(uint32(v5))+24:], uint64(int64(uint32(i32(3)))<<32|int64(uint32(v5+i32(4)))))
		m.fn31(i32(1066098), v5+i32(24), v4)
		panic("unreachable")
	}
l13:
	m.fn222(v4)
	panic("unreachable")
l23:
	m.fn41(v0, v1, v6, v3, v4)
	panic("unreachable")
l18:
	t24 := v5
	v7 = int64(uint32(i32(3))) << 32
	store64(m.memory[int64(uint32(t24))+32:], uint64(v7|int64(uint32(v5+i32(8)))))
	store64(m.memory[int64(uint32(v5))+24:], uint64(v7|int64(uint32(v5+i32(4)))))
	m.fn31(i32(1050440), v5+i32(24), v4)
	panic("unreachable")
}
func (m *Module) fn864(v0, v1 int32) int32 {
	var v2, v3, v4 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	{
		{
			{
				t1 := int32(load32(m.memory[int64(uint32(v1))+8:]))
				v3 = t1
				if v3&i32(0x2000000) != 0 {
					t3 := int32(load32(m.memory[uint32(v0):]))
					v4 = t3
					v3 = i32(9)
				l4:
					{
						t4 := int32(m.memory[int64(uint32(v4&i32(15)))+1098720])
						m.memory[uint32(v2+i32(8)+v3+i32(-2))] = byte(t4)
						v3 = v3 + i32(-1)
						v4 = int32(uint32(v4) >> 4)
						if v4 != 0 {
							goto l4
						}
					}
					v4 = i32(1)
					t5 := m.fn683(v1, i32(1), i32(1122454), i32(2), v2+i32(8)+v3+i32(-1), i32(9)-v3)
					if t5 == 0 {
						goto l2
					}
					goto l3
				}
				if v3&i32(0x4000000) != 0 {
					goto l1
				}
				t2 := m.fn47(v0, v1)
				if t2 == 0 {
					goto l2
				}
				v4 = i32(1)
				goto l3
			}
		l1:
			t6 := int32(load32(m.memory[uint32(v0):]))
			v4 = t6
			v3 = i32(9)
		l5:
			{
				t7 := int32(m.memory[int64(uint32(v4&i32(15)))+1122456])
				m.memory[uint32(v2+i32(8)+v3+i32(-2))] = byte(t7)
				v3 = v3 + i32(-1)
				v4 = int32(uint32(v4) >> 4)
				if v4 != 0 {
					goto l5
				}
			}
			v4 = i32(1)
			t8 := m.fn683(v1, i32(1), i32(1122454), i32(2), v2+i32(8)+v3+i32(-1), i32(9)-v3)
			if t8 != 0 {
				goto l3
			}
		}
	l2:
		{
			t9 := int32(load32(m.memory[uint32(v1):]))
			t10 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t11 := int32(load32(m.memory[int64(uint32(t10))+12:]))
			t12 := m.t0[uint(t11)].(func(int32, int32, int32) int32)(t9, i32(1273571), i32(2))
			if t12 == 0 {
				goto l6
			}
			v4 = i32(1)
			goto l3
		}
	l6:
		v3 = v0 + i32(4)
		{
			t13 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			v4 = t13
			if v4&i32(0x2000000) != 0 {
				t15 := int32(load32(m.memory[uint32(v3):]))
				v4 = t15
				v3 = i32(9)
			l9:
				{
					t16 := int32(m.memory[int64(uint32(v4&i32(15)))+1098720])
					m.memory[uint32(v2+i32(8)+v3+i32(-2))] = byte(t16)
					v3 = v3 + i32(-1)
					v4 = int32(uint32(v4) >> 4)
					if v4 != 0 {
						goto l9
					}
				}
				t17 := m.fn683(v1, i32(1), i32(1122454), i32(2), v2+i32(8)+v3+i32(-1), i32(9)-v3)
				v4 = t17
				goto l3
			}
			if v4&i32(0x4000000) != 0 {
				goto l8
			}
			t14 := m.fn47(v3, v1)
			v4 = t14
			goto l3
		}
	l8:
		t18 := int32(load32(m.memory[uint32(v3):]))
		v4 = t18
		v3 = i32(9)
	l10:
		{
			t19 := int32(m.memory[int64(uint32(v4&i32(15)))+1122456])
			m.memory[uint32(v2+i32(8)+v3+i32(-2))] = byte(t19)
			v3 = v3 + i32(-1)
			v4 = int32(uint32(v4) >> 4)
			if v4 != 0 {
				goto l10
			}
		}
		t20 := m.fn683(v1, i32(1), i32(1122454), i32(2), v2+i32(8)+v3+i32(-1), i32(9)-v3)
		v4 = t20
	}
l3:
	m.g0 = v2 + i32(16)
	return v4
}
func (m *Module) fn865(v0, v1 int32) int32 {
	var v2, v3, v4, v5 int32
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
		t4 := int32(load32(m.memory[int64(uint32(v5))+16:]))
		v1 = t4
		t5 := m.t0[uint(v1)].(func(int32, int32) int32)(t3, i32(39))
		if t5 != 0 {
			goto l0
		}
		t6 := int32(load32(m.memory[uint32(v0):]))
		m.fn856(v2, t6, i32(257))
		{
			t7 := int32(m.memory[int64(uint32(v2))+13])
			v3 = t7
			if uint32(v3) < uint32(i32(129)) {
				t10 := int32(m.memory[int64(uint32(v2))+12])
				t11 := v4
				t12 := v2
				v0 = t10
				t13 := int32(load32(m.memory[int64(uint32(v5))+12:]))
				t14 := m.t0[uint(t13)].(func(int32, int32, int32) int32)(t11, t12+v0, v3-v0)
				if t14 == 0 {
					goto l2
				}
				v3 = i32(1)
				goto l0
			}
			t8 := int32(load32(m.memory[uint32(v2):]))
			t9 := m.t0[uint(v1)].(func(int32, int32) int32)(v4, t8)
			if t9 == 0 {
				goto l2
			}
			v3 = i32(1)
			goto l0
		}
	l2:
		t15 := m.t0[uint(v1)].(func(int32, int32) int32)(v4, i32(39))
		v3 = t15
	}
l0:
	m.g0 = v2 + i32(16)
	return v3
}
func (m *Module) fn866(v0, v1 int32) int32 {
	t0 := m.fn9(v1, i32(1122596), i32(24))
	return t0
}
func (m *Module) fn867(v0, v1 int32) int32 {
	t0 := m.fn9(v1, i32(1122564), i32(32))
	return t0
}
func (m *Module) fn868(v0, v1, v2, v3, v4, v5 int32) {
	var v6 int32
	{
		if v2 == 0 {
			m.fn7(i32(1102113), i32(33), i32(1102148))
			panic("unreachable")
		}
		t0 := int32(m.memory[uint32(v1)])
		if uint32(t0) <= uint32(i32(48)) {
			m.fn7(i32(1102164), i32(31), i32(1102196))
			panic("unreachable")
		}
		store16(m.memory[uint32(v5):], uint16(i32(2)))
		{
			{
				v6 = int32(int16(v3))
				if v6 < i32(1) {
					store32(m.memory[int64(uint32(v5))+32:], uint32(v2))
					store32(m.memory[int64(uint32(v5))+28:], uint32(v1))
					store16(m.memory[int64(uint32(v5))+24:], uint16(i32(2)))
					store16(m.memory[int64(uint32(v5))+12:], uint16(i32(0)))
					store32(m.memory[int64(uint32(v5))+8:], uint32(i32(2)))
					store32(m.memory[int64(uint32(v5))+4:], uint32(i32(1098793)))
					t2 := v5
					v3 = i32(0) - v6
					store32(m.memory[int64(uint32(t2))+16:], uint32(v3))
					v1 = i32(3)
					if uint32(v4) <= uint32(v2) {
						goto l5
					}
					v2 = v4 - v2
					if uint32(v2) <= uint32(v3) {
						goto l5
					}
					v4 = v2 + v6
					goto l6
				}
				store32(m.memory[int64(uint32(v5))+4:], uint32(v1))
				t1 := v2
				v3 = v3 & i32(0xffff)
				if uint32(t1) > uint32(v3) {
					goto l3
				}
				store16(m.memory[int64(uint32(v5))+12:], uint16(i32(0)))
				store32(m.memory[int64(uint32(v5))+8:], uint32(v2))
				store32(m.memory[int64(uint32(v5))+16:], uint32(v3-v2))
				if v4 != 0 {
					store32(m.memory[int64(uint32(v5))+32:], uint32(i32(1)))
					store32(m.memory[int64(uint32(v5))+28:], uint32(i32(1100303)))
					store16(m.memory[int64(uint32(v5))+24:], uint16(i32(2)))
					goto l6
				}
				v1 = i32(2)
				goto l5
			}
		l3:
			store16(m.memory[int64(uint32(v5))+24:], uint16(i32(2)))
			store32(m.memory[int64(uint32(v5))+20:], uint32(i32(1)))
			store32(m.memory[int64(uint32(v5))+16:], uint32(i32(1100303)))
			store16(m.memory[int64(uint32(v5))+12:], uint16(i32(2)))
			store32(m.memory[int64(uint32(v5))+8:], uint32(v3))
			t3 := v5
			v2 = v2 - v3
			store32(m.memory[int64(uint32(t3))+32:], uint32(v2))
			store32(m.memory[int64(uint32(v5))+28:], uint32(v1+v3))
			if uint32(v4) > uint32(v2) {
				goto l7
			}
			v1 = i32(3)
			goto l5
		l7:
			v4 = v4 - v2
		}
	l6:
		store32(m.memory[int64(uint32(v5))+40:], uint32(v4))
		store16(m.memory[int64(uint32(v5))+36:], uint16(i32(0)))
		v1 = i32(4)
	l5:
		store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
		store32(m.memory[uint32(v0):], uint32(v5))
		return
	}
}
func (m *Module) fn869(v0 int32, v1 int64, v2, v3, v4, v5 int32) {
	var v6 int32
	var v7 int64
	var v8, v9, v10 int32
	var v11, v12 int64
	var v13, v14 int32
	var v15 int64
	var v16, v17, v18 int32
	var v19 int64
	t0 := m.g0
	v6 = t0 - i32(16)
	m.g0 = v6
	{
		{
			if v1 == i64(0) {
				m.fn7(i32(0x111a88), i32(28), i32(0x111aa4))
				panic("unreachable")
			}
			t1 := v2
			v7 = int64(bits.LeadingZeros64(uint64(v1)))
			v8 = t1 - int32(v7)
			v2 = (int32(int16(i32(-96)-v8))*i32(80) + i32(86960)) / i32(2126)
			if uint32(v2) > uint32(i32(80)) {
				m.fn36(v2, i32(81), i32(1120948))
				panic("unreachable")
			}
			t2 := v6
			v2 = v2 << 4
			t3 := int64(load64(m.memory[int64(uint32(v2))+1119608:]))
			m.fn1911(t2, t3, i64(0), i64_shl(v1, v7), i64(0))
			t4 := int64(load64(m.memory[uint32(v6):]))
			t5 := int64(load64(m.memory[int64(uint32(v6))+8:]))
			v1 = int64(uint64(t4)>>63) + t5
			t6 := int32(load16(m.memory[int64(uint32(v2))+1119616:]))
			t7 := v1
			v9 = i32(-64) - (v8 + t6)
			v7 = int64(uint32(v9))
			v10 = int32(i64_shr_u(t7, v7))
			t8 := int32(load16(m.memory[int64(uint32(v2))+1119618:]))
			v2 = t8
			{
				v11 = i64_shl(i64(1), v7)
				v12 = v11 + i64(-1)
				v7 = v12 & v1
				if !(v7 == 0) {
					goto l2
				}
				if uint32(v4) > uint32(i32(10)) {
					goto l3
				}
				t9 := int32(load32(m.memory[uint32(v4<<2+i32(1121644)):]))
				if uint32(t9) > uint32(v10) {
					goto l3
				}
			}
		l2:
			v13 = v9 & i32(63)
			if uint32(v10) < uint32(i32(10000)) {
				if uint32(v10) < uint32(i32(100)) {
					var p19 int32
					if uint32(v10) > uint32(i32(9)) {
						p19 = 1
					}
					v14 = p19
					p20 := i32(1)
					if v14 != 0 {
						p20 = i32(10)
					}
					v8 = p20
					goto l7
				}
				var p16 int32
				if uint32(v10) < uint32(i32(1000)) {
					p16 = 1
				}
				v8 = p16
				p17 := i32(3)
				if v8 != 0 {
					p17 = i32(2)
				}
				v14 = p17
				p18 := i32(1000)
				if v8 != 0 {
					p18 = i32(100)
				}
				v8 = p18
				goto l7
			}
			if uint32(v10) < uint32(i32(1000000)) {
				goto l5
			}
			{
				if uint32(v10) < uint32(i32(100000000)) {
					var p13 int32
					if uint32(v10) < uint32(i32(10000000)) {
						p13 = 1
					}
					v8 = p13
					p14 := i32(7)
					if v8 != 0 {
						p14 = i32(6)
					}
					v14 = p14
					p15 := i32(10000000)
					if v8 != 0 {
						p15 = i32(1000000)
					}
					v8 = p15
					goto l7
				}
				var p10 int32
				if uint32(v10) < uint32(i32(1000000000)) {
					p10 = 1
				}
				v8 = p10
				p11 := i32(9)
				if v8 != 0 {
					p11 = i32(8)
				}
				v14 = p11
				p12 := i32(1000000000)
				if v8 != 0 {
					p12 = i32(100000000)
				}
				v8 = p12
				goto l7
			}
		}
	l5:
		;
		var p21 int32
		if uint32(v10) < uint32(i32(100000)) {
			p21 = 1
		}
		v8 = p21
		p22 := i32(5)
		if v8 != 0 {
			p22 = i32(4)
		}
		v14 = p22
		p23 := i32(100000)
		if v8 != 0 {
			p23 = i32(10000)
		}
		v8 = p23
	}
l7:
	v15 = int64(uint32(v13))
	{
		v16 = int32(int16(v14 - v2 + i32(1)))
		t24 := v16
		v2 = int32(int16(v5))
		if t24 <= v2 {
			t29 := int64(uint64(v1) / uint64(i64(10)))
			m.fn870(v0, v3, v4, i32(0), v16, v5, t29, i64_shl(int64(uint32(v8)), v15), v11)
			goto l14
		}
		v17 = v9 & i32(0xffff)
		p25 := v4
		if uint32(v16-v2) < uint32(v4) {
			p25 = int32(int16(v16 - v5))
		}
		v18 = p25
		v13 = v18 + i32(-1)
		v2 = i32(0)
	l13:
		{
			t26 := int32(uint32(v10) / uint32(v8))
			v9 = t26
			if v4 == v2 {
				m.fn36(v4, v4, i32(1120980))
				panic("unreachable")
			}
			v10 = v10 - v9*v8
			m.memory[uint32(v3+v2)] = byte(v9 + i32(48))
			if v13 == v2 {
				m.fn870(v0, v3, v4, v18, v16, v5, i64_shl(int64(uint32(v10)), v15)+v7, i64_shl(int64(uint32(v8)), v15), v11)
				goto l14
			}
			if v14 == v2 {
				v2 = v2 + i32(1)
				v19 = int64(uint32((v17 + i32(-1)) & i32(63)))
				v1 = i64(1)
			l17:
				{
					if i64_shr_u(v1, v19) == 0 {
						goto l15
					}
					store32(m.memory[uint32(v0):], uint32(i32(0)))
					goto l14
				l15:
					if uint32(v2) >= uint32(v4) {
						m.fn36(v2, v4, i32(1120996))
						panic("unreachable")
					}
					t30 := v3 + v2
					v7 = v7 * i64(10)
					m.memory[uint32(t30)] = byte(int32(i64_shr_u(v7, v15)) + i32(48))
					v1 = v1 * i64(10)
					v7 = v7 & v12
					t31 := v18
					v2 = v2 + i32(1)
					if t31 != v2 {
						goto l17
					}
				}
				m.fn870(v0, v3, v4, v18, v16, v5, v7, v11, v1)
				goto l14
			}
			v2 = v2 + i32(1)
			var p27 int32
			if uint32(v8) < uint32(i32(10)) {
				p27 = 1
			}
			v9 = p27
			t28 := int32(uint32(v8) / uint32(i32(10)))
			v8 = t28
			if v9 == 0 {
				goto l13
			}
		}
		m.fn687(i32(1120964))
		panic("unreachable")
	}
l3:
	store32(m.memory[uint32(v0):], uint32(i32(0)))
l14:
	m.g0 = v6 + i32(16)
}
func (m *Module) fn870(v0, v1, v2, v3, v4, v5 int32, v6, v7, v8 int64) {
	var v9, v10, v11, v12 int32
	{
		if uint64(v7) <= uint64(v8) {
			store32(m.memory[uint32(v0):], uint32(i32(0)))
			return
		}
		if uint64(v7-v8) <= uint64(v8) {
			store32(m.memory[uint32(v0):], uint32(i32(0)))
			return
		}
		if uint64(v7-v6) <= uint64(v6) {
			goto l2
		}
		if uint64(v7-v6<<1) >= uint64(v8<<1) {
			if uint32(v3) <= uint32(v2) {
				goto l11
			}
			m.fn124(i32(0), v3, v2, i32(1121876))
			panic("unreachable")
		}
	l2:
		if uint64(v6) <= uint64(v8) {
			goto l4
		}
		t0 := v7
		v8 = v6 - v8
		if uint64(t0-v8) > uint64(v8) {
			goto l4
		}
		if uint32(v3) > uint32(v2) {
			m.fn124(i32(0), v3, v2, i32(0x111e44))
			panic("unreachable")
		}
		v9 = v1 + v3
		v10 = v3
	l7:
		{
			v11 = v10
			if v11 == 0 {
				goto l6
			}
			v10 = v11 + i32(-1)
			v12 = v10 + v1
			t1 := int32(m.memory[uint32(v12)])
			if t1 == i32(57) {
				goto l7
			}
		}
		t2 := int32(m.memory[uint32(v12)])
		m.memory[uint32(v12)] = byte(t2 + i32(1))
		v10 = v3 - v11
		if v10 == 0 {
			goto l8
		}
		memory_fill(m.memory, uint32(v1+v11), i32(48), uint32(v10))
		goto l8
	}
l6:
	if v3 != 0 {
		goto l9
	}
	v10 = i32(49)
	goto l10
l9:
	m.memory[uint32(v1)] = byte(i32(49))
	v10 = i32(48)
	v11 = v3 + i32(-1)
	if v11 == 0 {
		goto l10
	}
	memory_fill(m.memory, uint32(v1+i32(1)), i32(48), uint32(v11))
l10:
	v4 = int32(int16(v4 + i32(1)))
	if v4 <= int32(int16(v5)) {
		goto l8
	}
	if uint32(v3) >= uint32(v2) {
		goto l8
	}
	m.memory[uint32(v9)] = byte(v10)
	v3 = v3 + i32(1)
l8:
	if uint32(v3) <= uint32(v2) {
		goto l11
	}
	m.fn124(i32(0), v3, v2, i32(1121844))
	panic("unreachable")
l11:
	store16(m.memory[int64(uint32(v0))+8:], uint16(v4))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
	store32(m.memory[uint32(v0):], uint32(v1))
	return
l4:
	store32(m.memory[uint32(v0):], uint32(i32(0)))
}
func (m *Module) fn871(v0, v1, v2, v3, v4 int32) {
	var v5 int32
	var v6, v7, v8 int64
	var v9, v10, v11, v12, v13, v14, v15, v16, v17, v18, v19, v20, v21, v22, v23, v24, v25, v26, v27, v28, v29, v30, v31, v32, v33, v34, v35 int32
	t0 := m.g0
	v5 = t0 - i32(832)
	m.g0 = v5
	{
		t1 := int64(load64(m.memory[uint32(v1):]))
		v6 = t1
		if v6 == i64(0) {
			m.fn7(i32(0x111a88), i32(28), i32(1121360))
			panic("unreachable")
		}
		t2 := int64(load64(m.memory[int64(uint32(v1))+8:]))
		v7 = t2
		if v7 == i64(0) {
			m.fn7(i32(1121028), i32(29), i32(1121376))
			panic("unreachable")
		}
		t3 := int64(load64(m.memory[int64(uint32(v1))+16:]))
		v8 = t3
		if v8 == i64(0) {
			m.fn7(i32(1121076), i32(28), i32(1121392))
			panic("unreachable")
		}
		if uint64(v8) > uint64(v6^i64(-1)) {
			m.fn7(i32(1121288), i32(54), i32(1121504))
			panic("unreachable")
		}
		if uint64(v6) < uint64(v7) {
			m.fn7(i32(1121216), i32(55), i32(1121488))
			panic("unreachable")
		}
		t4 := int32(int16(load16(m.memory[int64(uint32(v1))+24:])))
		v1 = t4
		store64(m.memory[int64(uint32(v5))+8:], uint64(v6))
		t6 := v5
		p5 := i32(2)
		if uint64(v6) < uint64(i64(0x100000000)) {
			p5 = i32(1)
		}
		store32(m.memory[int64(uint32(t6))+168:], uint32(p5))
		memory_zero(m.memory, uint32(v5+i32(16)), uint32(i32(152)))
		memory_zero(m.memory, uint32(v5+i32(180)), uint32(i32(156)))
		store32(m.memory[int64(uint32(v5))+176:], uint32(i32(1)))
		store32(m.memory[int64(uint32(v5))+336:], uint32(i32(1)))
		v9 = int32(int64(uint64((int64(v1)-int64(bits.LeadingZeros64(uint64(v6+i64(-1)))))*i64(1292913986)+i64(82746495104)) >> 32))
		v10 = int32(int16(v9))
		if v1 < i32(0) {
			goto l5
		}
		_ = m.fn859(v5+i32(8), v1)
		goto l6
	l5:
		_ = m.fn859(v5+i32(176), int32(int16(i32(0)-v1)))
	l6:
		if v10 > i32(-1) {
			goto l7
		}
		_ = m.fn872(v5+i32(8), (i32(0)-v10)&i32(0xffff))
		goto l8
	l7:
		_ = m.fn872(v5+i32(176), v9&i32(0x7fff))
	l8:
		memory_copy(m.memory, uint32(v5+i32(668)), uint32(v5+i32(176)), uint32(i32(164)))
		v11 = v3
		if uint32(v3) < uint32(i32(10)) {
			goto l9
		}
		v12 = v5 + i32(668) + i32(-8)
		v11 = v3
	l15:
		{
			{
				t11 := int32(load32(m.memory[int64(uint32(v5))+828:]))
				v1 = t11
				if uint32(v1) > uint32(i32(40)) {
					m.fn124(i32(0), v1, i32(40), i32(1099776))
					panic("unreachable")
				}
				if v1 == 0 {
					goto l11
				}
				v1 = v1 << 2
				v9 = v1 + i32(-4)
				if v9 != 0 {
					goto l12
				}
				v1 = v5 + i32(668) + v1
				v6 = i64(0)
				goto l13
			}
		l12:
			v13 = int32(uint32(v9)>>2) + i32(1)
			v14 = v13 & i32(1)
			v9 = v12 + v1
			v13 = v13 & i32(0x7ffffffe)
			v6 = i64(0)
		l14:
			{
				v1 = v9
				v9 = v1 + i32(4)
				t12 := int64(load32(m.memory[uint32(v9):]))
				t13 := v9
				v6 = v6<<32 | t12
				t14 := int64(uint64(v6) / uint64(i64(1000000000)))
				v7 = t14
				store32(m.memory[uint32(t13):], uint32(v7))
				t15 := int64(load32(m.memory[uint32(v1):]))
				t16 := v1
				v6 = (v6-v7*i64(1000000000))<<32 | t15
				t17 := int64(uint64(v6) / uint64(i64(1000000000)))
				v7 = t17
				store32(m.memory[uint32(t16):], uint32(v7))
				v6 = v6 - v7*i64(1000000000)
				v9 = v1 + i32(-8)
				v13 = v13 + i32(-2)
				if v13 != 0 {
					goto l14
				}
			}
			if v14 == 0 {
				goto l11
			}
		l13:
			v1 = v1 + i32(-4)
			t18 := int64(load32(m.memory[uint32(v1):]))
			t19 := int64(uint64(v6<<32|t18) / uint64(i64(1000000000)))
			store32(m.memory[uint32(v1):], uint32(t19))
		}
	l11:
		v11 = v11 + i32(-9)
		if uint32(v11) > uint32(i32(9)) {
			goto l15
		}
	l9:
		t20 := int32(load32(m.memory[int64(uint32(v11<<2))+1121648:]))
		v9 = t20 << 1
		if v9 == 0 {
			m.fn7(i32(1099720), i32(27), i32(1099776))
			panic("unreachable")
		}
		{
			{
				{
					t21 := int32(load32(m.memory[int64(uint32(v5))+828:]))
					v1 = t21
					if uint32(v1) > uint32(i32(40)) {
						m.fn124(i32(0), v1, i32(40), i32(1099776))
						panic("unreachable")
					}
					if v1 != 0 {
						v6 = int64(uint32(v9))
						v1 = v1 << 2
						v9 = v1 + i32(-4)
						if v9 != 0 {
							goto l20
						}
						v1 = v5 + i32(668) + v1
						v7 = i64(0)
						goto l21
					}
					v1 = i32(0)
					goto l19
				}
			l20:
				v9 = int32(uint32(v9)>>2) + i32(1)
				v11 = v9 & i32(1)
				v9 = v9 & i32(0x7ffffffe)
				v13 = v1 + (v5 + i32(668)) + i32(-8)
				v7 = i64(0)
			l22:
				{
					v1 = v13
					v13 = v1 + i32(4)
					t22 := int64(load32(m.memory[uint32(v13):]))
					t23 := v13
					v7 = v7<<32 | t22
					t24 := int64(uint64(v7) / uint64(v6))
					v8 = t24
					store32(m.memory[uint32(t23):], uint32(v8))
					t25 := int64(load32(m.memory[uint32(v1):]))
					t26 := v1
					v7 = (v7-v8*v6)<<32 | t25
					t27 := int64(uint64(v7) / uint64(v6))
					v8 = t27
					store32(m.memory[uint32(t26):], uint32(v8))
					v7 = v7 - v8*v6
					v13 = v1 + i32(-8)
					v9 = v9 + i32(-2)
					if v9 != 0 {
						goto l22
					}
				}
				if v11 == 0 {
					goto l23
				}
			l21:
				v1 = v1 + i32(-4)
				t28 := int64(load32(m.memory[uint32(v1):]))
				t29 := int64(uint64(v7<<32|t28) / uint64(v6))
				store32(m.memory[uint32(v1):], uint32(t29))
			}
		l23:
			t30 := int32(load32(m.memory[int64(uint32(v5))+828:]))
			v1 = t30
		}
	l19:
		t31 := int32(load32(m.memory[int64(uint32(v5))+168:]))
		v15 = t31
		p32 := v1
		if uint32(v15) > uint32(v1) {
			p32 = v15
		}
		v16 = p32
		if uint32(v16) > uint32(i32(40)) {
			m.fn124(i32(0), v16, i32(40), i32(1099776))
			panic("unreachable")
		}
		if v16 != 0 {
			goto l25
		}
		v16 = i32(0)
		goto l26
	l25:
		v11 = i32(0)
		v14 = i32(0)
		{
			if v16 == i32(1) {
				goto l27
			}
			v17 = v16 & i32(1)
			v18 = v16 & i32(62)
			v11 = i32(0)
			v1 = v5 + i32(668)
			v9 = v5 + i32(8)
			v14 = i32(0)
		l28:
			{
				t33 := int32(load32(m.memory[uint32(v9):]))
				t34 := v1
				v12 = t33
				t35 := int32(load32(m.memory[uint32(v1):]))
				v13 = v12 + t35
				v19 = v13 + v11&i32(1)
				store32(m.memory[uint32(t34):], uint32(v19))
				v11 = v1 + i32(4)
				t36 := int32(load32(m.memory[uint32(v9+i32(4)):]))
				t37 := v11
				v20 = t36
				t38 := int32(load32(m.memory[uint32(v11):]))
				v11 = v20 + t38
				t39 := v11
				var p40 int32
				if uint32(v13) < uint32(v12) {
					p40 = 1
				}
				var p41 int32
				if uint32(v19) < uint32(v13) {
					p41 = 1
				}
				v13 = t39 + (p40 | p41)
				store32(m.memory[uint32(t37):], uint32(v13))
				var p42 int32
				if uint32(v11) < uint32(v20) {
					p42 = 1
				}
				var p43 int32
				if uint32(v13) < uint32(v11) {
					p43 = 1
				}
				v11 = p42 | p43
				v9 = v9 + i32(8)
				v1 = v1 + i32(8)
				t44 := v18
				v14 = v14 + i32(2)
				if t44 != v14 {
					goto l28
				}
			}
			if v17 == 0 {
				goto l29
			}
		l27:
			t45 := v5 + i32(668)
			v1 = v14 << 2
			v9 = t45 + v1
			t46 := int32(load32(m.memory[uint32(v5+i32(8)+v1):]))
			t47 := v9
			v13 = t46
			t48 := int32(load32(m.memory[uint32(v9):]))
			v1 = v13 + t48
			v9 = v1 + v11
			store32(m.memory[uint32(t47):], uint32(v9))
			if uint32(v1) < uint32(v13) {
				goto l30
			}
			if uint32(v9) < uint32(v1) {
				goto l30
			}
			goto l26
		}
	l29:
		if v11 == 0 {
			goto l26
		}
	l30:
		if v16 == i32(40) {
			m.fn36(i32(40), i32(40), i32(1099776))
			panic("unreachable")
		}
		store32(m.memory[uint32(v5+i32(668)+v16<<2):], uint32(i32(1)))
		v16 = v16 + i32(1)
	l26:
		store32(m.memory[int64(uint32(v5))+828:], uint32(v16))
		t49 := int32(load32(m.memory[int64(uint32(v5))+336:]))
		v21 = t49
		p50 := v16
		if uint32(v21) > uint32(v16) {
			p50 = v21
		}
		v1 = p50
		if uint32(v1) >= uint32(i32(41)) {
			m.fn124(i32(0), v1, i32(40), i32(1099776))
			panic("unreachable")
		}
		v1 = v1 << 2
		v9 = v5 + i32(668) + i32(-4)
	l34:
		{
			if v1 == 0 {
				goto l33
			}
			t51 := int32(load32(m.memory[uint32(v9+v1):]))
			v13 = t51
			t52 := v13
			v1 = v1 + i32(-4)
			t53 := int32(load32(m.memory[uint32(v1+(v5+i32(176))):]))
			v11 = t53
			if t52 == v11 {
				goto l34
			}
		}
		if uint32(v13) >= uint32(v11) {
			goto l33
		}
		if v15 != 0 {
			v14 = v15 << 2
			v1 = v14 + i32(-4)
			v9 = int32(uint32(v1)>>2) + i32(1)
			v11 = v9 & i32(3)
			if uint32(v1) >= uint32(i32(12)) {
				goto l37
			}
			v6 = i64(0)
			v1 = v5 + i32(8)
			goto l38
		l37:
			v9 = v9 & i32(0x7ffffffc)
			v6 = i64(0)
			v1 = v5 + i32(8)
		l39:
			{
				t54 := int64(load32(m.memory[uint32(v1):]))
				t55 := v1
				v6 = t54*i64(10) + v6
				store32(m.memory[uint32(t55):], uint32(v6))
				v13 = v1 + i32(4)
				t56 := int64(load32(m.memory[uint32(v13):]))
				t57 := v13
				v6 = t56*i64(10) + int64(uint64(v6)>>32)
				store32(m.memory[uint32(t57):], uint32(v6))
				v13 = v1 + i32(8)
				t58 := int64(load32(m.memory[uint32(v13):]))
				t59 := v13
				v6 = t58*i64(10) + int64(uint64(v6)>>32)
				store32(m.memory[uint32(t59):], uint32(v6))
				v13 = v1 + i32(12)
				t60 := int64(load32(m.memory[uint32(v13):]))
				t61 := v13
				v6 = t60*i64(10) + int64(uint64(v6)>>32)
				store32(m.memory[uint32(t61):], uint32(v6))
				v6 = int64(uint64(v6) >> 32)
				v1 = v1 + i32(16)
				v9 = v9 + i32(-4)
				if v9 != 0 {
					goto l39
				}
			}
			if v11 == 0 {
				goto l40
			}
		l38:
			v9 = v11 << 2
		l41:
			{
				t62 := int64(load32(m.memory[uint32(v1):]))
				t63 := v1
				v6 = t62*i64(10) + v6
				store32(m.memory[uint32(t63):], uint32(v6))
				v1 = v1 + i32(4)
				v6 = int64(uint64(v6) >> 32)
				v9 = v9 + i32(-4)
				if v9 != 0 {
					goto l41
				}
			}
		l40:
			if v6 == 0 {
				goto l42
			}
			if v15 == i32(40) {
				m.fn36(i32(40), i32(40), i32(1099776))
				panic("unreachable")
			}
			store32(m.memory[uint32(v5+i32(8)+v14):], uint32(int32(v6)))
			v15 = v15 + i32(1)
		l42:
			store32(m.memory[int64(uint32(v5))+168:], uint32(v15))
			goto l36
		}
		v15 = i32(0)
		store32(m.memory[int64(uint32(v5))+168:], uint32(i32(0)))
		goto l36
	l33:
		v10 = v10 + i32(1)
	l36:
		v22 = i32(0)
		v12 = i32(1)
		{
			v1 = int32(int16(v10))
			t64 := v1
			v9 = int32(int16(v4))
			var p65 int32
			if t64 < v9 {
				p65 = 1
			}
			v23 = p65
			if v23 != 0 {
				goto l44
			}
			p66 := v3
			if uint32(v1-v9) < uint32(v3) {
				p66 = int32(int16(v10 - v4))
			}
			v24 = p66
			if v24 != 0 {
				memory_copy(m.memory, uint32(v5+i32(340)), uint32(v5+i32(176)), uint32(i32(164)))
				t67 := m.fn859(v5+i32(340), i32(1))
				v25 = t67
				memory_copy(m.memory, uint32(v5+i32(504)), uint32(v5+i32(176)), uint32(i32(164)))
				t68 := m.fn859(v5+i32(504), i32(2))
				v26 = t68
				memory_copy(m.memory, uint32(v5+i32(668)), uint32(v5+i32(176)), uint32(i32(164)))
				v17 = v5 + i32(176) + i32(-4)
				v4 = v5 + i32(340) + i32(-4)
				v16 = v5 + i32(504) + i32(-4)
				v18 = v5 + i32(668) + i32(-4)
				t69 := m.fn859(v5+i32(668), i32(3))
				v27 = t69
				t70 := int32(load32(m.memory[int64(uint32(v25))+160:]))
				v28 = t70
				t71 := int32(load32(m.memory[int64(uint32(v26))+160:]))
				v29 = t71
				t72 := int32(load32(m.memory[int64(uint32(v27))+160:]))
				v30 = t72
				v31 = i32(0)
			l95:
				{
					v32 = v31
					if uint32(v15) >= uint32(i32(41)) {
						m.fn124(i32(0), v15, i32(40), i32(1099776))
						panic("unreachable")
					}
					v31 = v32 + i32(1)
					v13 = v15 << 2
					v1 = i32(0)
				l49:
					{
						if v13 == v1 {
							if uint32(v24) > uint32(v3) {
								m.fn124(v32, v24, v3, i32(0x111cc0))
								panic("unreachable")
							}
							if v24 == v32 {
								goto l97
							}
							v1 = v24 - v32
							if v1 == 0 {
								goto l97
							}
							memory_fill(m.memory, uint32(v2+v32), i32(48), uint32(v1))
						l97:
							store16(m.memory[int64(uint32(v0))+8:], uint16(v10))
							store32(m.memory[int64(uint32(v0))+4:], uint32(v24))
							goto l98
						}
						v9 = v5 + i32(8) + v1
						v1 = v1 + i32(4)
						t73 := int32(load32(m.memory[uint32(v9):]))
						if t73 == 0 {
							goto l49
						}
					}
					p74 := v15
					if uint32(v30) > uint32(v15) {
						p74 = v30
					}
					v33 = p74
					if uint32(v33) >= uint32(i32(41)) {
						m.fn124(i32(0), v33, i32(40), i32(1099776))
						panic("unreachable")
					}
					v1 = v33 << 2
				l52:
					{
						if v1 == 0 {
							goto l51
						}
						v9 = v18 + v1
						v1 = v1 + i32(-4)
						t75 := int32(load32(m.memory[uint32(v1+(v5+i32(8))):]))
						v13 = t75
						t76 := int32(load32(m.memory[uint32(v9):]))
						t77 := v13
						v9 = t76
						if t77 == v9 {
							goto l52
						}
					}
					v34 = i32(0)
					if uint32(v13) < uint32(v9) {
						goto l53
					}
				l51:
					v11 = i32(1)
					v14 = i32(0)
					{
						if v33 == i32(1) {
							goto l54
						}
						v35 = v33 & i32(1)
						v15 = v33 & i32(62)
						v14 = i32(0)
						v11 = i32(1)
						v1 = v5 + i32(8)
						v9 = v5 + i32(668)
					l55:
						{
							t78 := int32(load32(m.memory[uint32(v1):]))
							t79 := v1
							v12 = t78
							t80 := int32(load32(m.memory[uint32(v9):]))
							v13 = v12 + (t80 ^ i32(-1))
							v19 = v13 + v11&i32(1)
							store32(m.memory[uint32(t79):], uint32(v19))
							v11 = v1 + i32(4)
							t81 := int32(load32(m.memory[uint32(v11):]))
							t82 := v11
							v20 = t81
							t83 := int32(load32(m.memory[uint32(v9+i32(4)):]))
							v11 = v20 + (t83 ^ i32(-1))
							t84 := v11
							var p85 int32
							if uint32(v13) < uint32(v12) {
								p85 = 1
							}
							var p86 int32
							if uint32(v19) < uint32(v13) {
								p86 = 1
							}
							v13 = t84 + (p85 | p86)
							store32(m.memory[uint32(t82):], uint32(v13))
							var p87 int32
							if uint32(v11) < uint32(v20) {
								p87 = 1
							}
							var p88 int32
							if uint32(v13) < uint32(v11) {
								p88 = 1
							}
							v11 = p87 | p88
							v9 = v9 + i32(8)
							v1 = v1 + i32(8)
							t89 := v15
							v14 = v14 + i32(2)
							if t89 != v14 {
								goto l55
							}
						}
						if v35 == 0 {
							goto l56
						}
					l54:
						t90 := v5 + i32(8)
						v1 = v14 << 2
						v9 = t90 + v1
						t91 := int32(load32(m.memory[uint32(v9):]))
						t92 := v9
						v9 = t91
						t93 := int32(load32(m.memory[uint32(v27+v1):]))
						v1 = v9 + (t93 ^ i32(-1))
						v13 = v1 + v11
						store32(m.memory[uint32(t92):], uint32(v13))
						if uint32(v1) < uint32(v9) {
							goto l57
						}
						if uint32(v13) < uint32(v1) {
							goto l57
						}
						goto l58
					}
				l56:
					if v11 == 0 {
						goto l58
					}
				l57:
					store32(m.memory[int64(uint32(v5))+168:], uint32(v33))
					v34 = i32(8)
					v15 = v33
				l53:
					p94 := v15
					if uint32(v29) > uint32(v15) {
						p94 = v29
					}
					v33 = p94
					if uint32(v33) >= uint32(i32(41)) {
						m.fn124(i32(0), v33, i32(40), i32(1099776))
						panic("unreachable")
					}
					v1 = v33 << 2
				l61:
					{
						if v1 == 0 {
							goto l60
						}
						v9 = v16 + v1
						v1 = v1 + i32(-4)
						t95 := int32(load32(m.memory[uint32(v1+(v5+i32(8))):]))
						v13 = t95
						t96 := int32(load32(m.memory[uint32(v9):]))
						t97 := v13
						v9 = t96
						if t97 == v9 {
							goto l61
						}
					}
					if uint32(v13) >= uint32(v9) {
						goto l60
					}
					v33 = v15
					goto l62
				l60:
					if v33 == 0 {
						goto l63
					}
					v11 = i32(1)
					v14 = i32(0)
					{
						if v33 == i32(1) {
							goto l64
						}
						v35 = v33 & i32(1)
						v15 = v33 & i32(62)
						v14 = i32(0)
						v11 = i32(1)
						v1 = v5 + i32(8)
						v9 = v5 + i32(504)
					l65:
						{
							t98 := int32(load32(m.memory[uint32(v1):]))
							t99 := v1
							v12 = t98
							t100 := int32(load32(m.memory[uint32(v9):]))
							v13 = v12 + (t100 ^ i32(-1))
							v19 = v13 + v11&i32(1)
							store32(m.memory[uint32(t99):], uint32(v19))
							v11 = v1 + i32(4)
							t101 := int32(load32(m.memory[uint32(v11):]))
							t102 := v11
							v20 = t101
							t103 := int32(load32(m.memory[uint32(v9+i32(4)):]))
							v11 = v20 + (t103 ^ i32(-1))
							t104 := v11
							var p105 int32
							if uint32(v13) < uint32(v12) {
								p105 = 1
							}
							var p106 int32
							if uint32(v19) < uint32(v13) {
								p106 = 1
							}
							v13 = t104 + (p105 | p106)
							store32(m.memory[uint32(t102):], uint32(v13))
							var p107 int32
							if uint32(v11) < uint32(v20) {
								p107 = 1
							}
							var p108 int32
							if uint32(v13) < uint32(v11) {
								p108 = 1
							}
							v11 = p107 | p108
							v9 = v9 + i32(8)
							v1 = v1 + i32(8)
							t109 := v15
							v14 = v14 + i32(2)
							if t109 != v14 {
								goto l65
							}
						}
						if v35 == 0 {
							goto l66
						}
					l64:
						t110 := v5 + i32(8)
						v1 = v14 << 2
						v9 = t110 + v1
						t111 := int32(load32(m.memory[uint32(v9):]))
						t112 := v9
						v9 = t111
						t113 := int32(load32(m.memory[uint32(v26+v1):]))
						v1 = v9 + (t113 ^ i32(-1))
						v13 = v1 + v11
						store32(m.memory[uint32(t112):], uint32(v13))
						if uint32(v1) < uint32(v9) {
							goto l63
						}
						if uint32(v13) < uint32(v1) {
							goto l63
						}
						goto l67
					}
				l66:
					if v11 == 0 {
						goto l67
					}
				l63:
					store32(m.memory[int64(uint32(v5))+168:], uint32(v33))
					v34 = v34 | i32(4)
				l62:
					p114 := v33
					if uint32(v28) > uint32(v33) {
						p114 = v28
					}
					v35 = p114
					if uint32(v35) >= uint32(i32(41)) {
						m.fn124(i32(0), v35, i32(40), i32(1099776))
						panic("unreachable")
					}
					v1 = v35 << 2
				l70:
					{
						if v1 == 0 {
							goto l69
						}
						v9 = v4 + v1
						v1 = v1 + i32(-4)
						t115 := int32(load32(m.memory[uint32(v1+(v5+i32(8))):]))
						v13 = t115
						t116 := int32(load32(m.memory[uint32(v9):]))
						t117 := v13
						v9 = t116
						if t117 == v9 {
							goto l70
						}
					}
					if uint32(v13) >= uint32(v9) {
						goto l69
					}
					v35 = v33
					goto l71
				l69:
					if v35 == 0 {
						goto l72
					}
					v11 = i32(1)
					v14 = i32(0)
					{
						if v35 == i32(1) {
							goto l73
						}
						v33 = v35 & i32(1)
						v15 = v35 & i32(62)
						v14 = i32(0)
						v11 = i32(1)
						v1 = v5 + i32(8)
						v9 = v5 + i32(340)
					l74:
						{
							t118 := int32(load32(m.memory[uint32(v1):]))
							t119 := v1
							v12 = t118
							t120 := int32(load32(m.memory[uint32(v9):]))
							v13 = v12 + (t120 ^ i32(-1))
							v19 = v13 + v11&i32(1)
							store32(m.memory[uint32(t119):], uint32(v19))
							v11 = v1 + i32(4)
							t121 := int32(load32(m.memory[uint32(v11):]))
							t122 := v11
							v20 = t121
							t123 := int32(load32(m.memory[uint32(v9+i32(4)):]))
							v11 = v20 + (t123 ^ i32(-1))
							t124 := v11
							var p125 int32
							if uint32(v13) < uint32(v12) {
								p125 = 1
							}
							var p126 int32
							if uint32(v19) < uint32(v13) {
								p126 = 1
							}
							v13 = t124 + (p125 | p126)
							store32(m.memory[uint32(t122):], uint32(v13))
							var p127 int32
							if uint32(v11) < uint32(v20) {
								p127 = 1
							}
							var p128 int32
							if uint32(v13) < uint32(v11) {
								p128 = 1
							}
							v11 = p127 | p128
							v9 = v9 + i32(8)
							v1 = v1 + i32(8)
							t129 := v15
							v14 = v14 + i32(2)
							if t129 != v14 {
								goto l74
							}
						}
						if v33 == 0 {
							goto l75
						}
					l73:
						t130 := v5 + i32(8)
						v1 = v14 << 2
						v9 = t130 + v1
						t131 := int32(load32(m.memory[uint32(v9):]))
						t132 := v9
						v9 = t131
						t133 := int32(load32(m.memory[uint32(v25+v1):]))
						v1 = v9 + (t133 ^ i32(-1))
						v13 = v1 + v11
						store32(m.memory[uint32(t132):], uint32(v13))
						if uint32(v1) < uint32(v9) {
							goto l72
						}
						if uint32(v13) < uint32(v1) {
							goto l72
						}
						goto l76
					}
				l75:
					if v11 == 0 {
						goto l76
					}
				l72:
					store32(m.memory[int64(uint32(v5))+168:], uint32(v35))
					v34 = v34 + i32(2)
				l71:
					p134 := v35
					if uint32(v21) > uint32(v35) {
						p134 = v21
					}
					v15 = p134
					if uint32(v15) >= uint32(i32(41)) {
						m.fn124(i32(0), v15, i32(40), i32(1099776))
						panic("unreachable")
					}
					v1 = v15 << 2
				l79:
					{
						if v1 == 0 {
							goto l78
						}
						v9 = v17 + v1
						v1 = v1 + i32(-4)
						t135 := int32(load32(m.memory[uint32(v1+(v5+i32(8))):]))
						v13 = t135
						t136 := int32(load32(m.memory[uint32(v9):]))
						t137 := v13
						v9 = t136
						if t137 == v9 {
							goto l79
						}
					}
					if uint32(v13) >= uint32(v9) {
						goto l78
					}
					v15 = v35
					goto l80
				l78:
					if v15 == 0 {
						goto l81
					}
					v11 = i32(1)
					v14 = i32(0)
					{
						if v15 == i32(1) {
							goto l82
						}
						v35 = v15 & i32(1)
						v33 = v15 & i32(62)
						v14 = i32(0)
						v11 = i32(1)
						v1 = v5 + i32(8)
						v9 = v5 + i32(176)
					l83:
						{
							t138 := int32(load32(m.memory[uint32(v1):]))
							t139 := v1
							v12 = t138
							t140 := int32(load32(m.memory[uint32(v9):]))
							v13 = v12 + (t140 ^ i32(-1))
							v19 = v13 + v11&i32(1)
							store32(m.memory[uint32(t139):], uint32(v19))
							v11 = v1 + i32(4)
							t141 := int32(load32(m.memory[uint32(v11):]))
							t142 := v11
							v20 = t141
							t143 := int32(load32(m.memory[uint32(v9+i32(4)):]))
							v11 = v20 + (t143 ^ i32(-1))
							t144 := v11
							var p145 int32
							if uint32(v13) < uint32(v12) {
								p145 = 1
							}
							var p146 int32
							if uint32(v19) < uint32(v13) {
								p146 = 1
							}
							v13 = t144 + (p145 | p146)
							store32(m.memory[uint32(t142):], uint32(v13))
							var p147 int32
							if uint32(v11) < uint32(v20) {
								p147 = 1
							}
							var p148 int32
							if uint32(v13) < uint32(v11) {
								p148 = 1
							}
							v11 = p147 | p148
							v9 = v9 + i32(8)
							v1 = v1 + i32(8)
							t149 := v33
							v14 = v14 + i32(2)
							if t149 != v14 {
								goto l83
							}
						}
						if v35 == 0 {
							goto l84
						}
					l82:
						t150 := v5 + i32(8)
						v1 = v14 << 2
						v9 = t150 + v1
						t151 := int32(load32(m.memory[uint32(v9):]))
						t152 := v9
						v9 = t151
						t153 := int32(load32(m.memory[uint32(v5+i32(176)+v1):]))
						v1 = v9 + (t153 ^ i32(-1))
						v13 = v1 + v11
						store32(m.memory[uint32(t152):], uint32(v13))
						if uint32(v1) < uint32(v9) {
							goto l81
						}
						if uint32(v13) < uint32(v1) {
							goto l81
						}
						goto l85
					}
				l84:
					if v11 == 0 {
						goto l85
					}
				l81:
					store32(m.memory[int64(uint32(v5))+168:], uint32(v15))
					v34 = v34 + i32(1)
				l80:
					if v32 == v3 {
						m.fn36(v3, v3, i32(1121456))
						panic("unreachable")
					}
					m.memory[uint32(v2+v32)] = byte(v34 + i32(48))
					if v15 != 0 {
						goto l87
					}
					v15 = i32(0)
					goto l88
				l87:
					v14 = v15 << 2
					v1 = v14 + i32(-4)
					v9 = int32(uint32(v1)>>2) + i32(1)
					v11 = v9 & i32(3)
					if uint32(v1) >= uint32(i32(12)) {
						goto l89
					}
					v6 = i64(0)
					v1 = v5 + i32(8)
					goto l90
				l89:
					v9 = v9 & i32(0x7ffffffc)
					v6 = i64(0)
					v1 = v5 + i32(8)
				l91:
					{
						t154 := int64(load32(m.memory[uint32(v1):]))
						t155 := v1
						v6 = t154*i64(10) + v6
						store32(m.memory[uint32(t155):], uint32(v6))
						v13 = v1 + i32(4)
						t156 := int64(load32(m.memory[uint32(v13):]))
						t157 := v13
						v6 = t156*i64(10) + int64(uint64(v6)>>32)
						store32(m.memory[uint32(t157):], uint32(v6))
						v13 = v1 + i32(8)
						t158 := int64(load32(m.memory[uint32(v13):]))
						t159 := v13
						v6 = t158*i64(10) + int64(uint64(v6)>>32)
						store32(m.memory[uint32(t159):], uint32(v6))
						v13 = v1 + i32(12)
						t160 := int64(load32(m.memory[uint32(v13):]))
						t161 := v13
						v6 = t160*i64(10) + int64(uint64(v6)>>32)
						store32(m.memory[uint32(t161):], uint32(v6))
						v6 = int64(uint64(v6) >> 32)
						v1 = v1 + i32(16)
						v9 = v9 + i32(-4)
						if v9 != 0 {
							goto l91
						}
					}
					if v11 == 0 {
						goto l92
					}
				l90:
					v9 = v11 << 2
				l93:
					{
						t162 := int64(load32(m.memory[uint32(v1):]))
						t163 := v1
						v6 = t162*i64(10) + v6
						store32(m.memory[uint32(t163):], uint32(v6))
						v1 = v1 + i32(4)
						v6 = int64(uint64(v6) >> 32)
						v9 = v9 + i32(-4)
						if v9 != 0 {
							goto l93
						}
					}
				l92:
					if v6 == 0 {
						goto l88
					}
					if v15 == i32(40) {
						m.fn36(i32(40), i32(40), i32(1099776))
						panic("unreachable")
					}
					store32(m.memory[uint32(v5+i32(8)+v14):], uint32(int32(v6)))
					v15 = v15 + i32(1)
				l88:
					store32(m.memory[int64(uint32(v5))+168:], uint32(v15))
					if v31 != v24 {
						goto l95
					}
				}
				v12 = i32(0)
				goto l46
			}
		}
	l44:
		v24 = i32(0)
		goto l46
	}
l85:
	m.fn7(i32(1099747), i32(26), i32(1099776))
	panic("unreachable")
l76:
	m.fn7(i32(1099747), i32(26), i32(1099776))
	panic("unreachable")
l67:
	m.fn7(i32(1099747), i32(26), i32(1099776))
	panic("unreachable")
l58:
	m.fn7(i32(1099747), i32(26), i32(1099776))
	panic("unreachable")
l46:
	{
		if v21 == 0 {
			goto l99
		}
		v14 = v21 << 2
		v1 = v14 + i32(-4)
		v9 = int32(uint32(v1)>>2) + i32(1)
		v11 = v9 & i32(3)
		if uint32(v1) >= uint32(i32(12)) {
			goto l100
		}
		v6 = i64(0)
		v1 = v5 + i32(176)
		goto l101
	l100:
		v9 = v9 & i32(0x7ffffffc)
		v6 = i64(0)
		v1 = v5 + i32(176)
	l102:
		{
			t164 := int64(load32(m.memory[uint32(v1):]))
			t165 := v1
			v6 = t164*i64(5) + v6
			store32(m.memory[uint32(t165):], uint32(v6))
			v13 = v1 + i32(4)
			t166 := int64(load32(m.memory[uint32(v13):]))
			t167 := v13
			v6 = t166*i64(5) + int64(uint64(v6)>>32)
			store32(m.memory[uint32(t167):], uint32(v6))
			v13 = v1 + i32(8)
			t168 := int64(load32(m.memory[uint32(v13):]))
			t169 := v13
			v6 = t168*i64(5) + int64(uint64(v6)>>32)
			store32(m.memory[uint32(t169):], uint32(v6))
			v13 = v1 + i32(12)
			t170 := int64(load32(m.memory[uint32(v13):]))
			t171 := v13
			v6 = t170*i64(5) + int64(uint64(v6)>>32)
			store32(m.memory[uint32(t171):], uint32(v6))
			v6 = int64(uint64(v6) >> 32)
			v1 = v1 + i32(16)
			v9 = v9 + i32(-4)
			if v9 != 0 {
				goto l102
			}
		}
		if v11 == 0 {
			goto l103
		}
	l101:
		v9 = v11 << 2
	l104:
		{
			t172 := int64(load32(m.memory[uint32(v1):]))
			t173 := v1
			v6 = t172*i64(5) + v6
			store32(m.memory[uint32(t173):], uint32(v6))
			v1 = v1 + i32(4)
			v6 = int64(uint64(v6) >> 32)
			v9 = v9 + i32(-4)
			if v9 != 0 {
				goto l104
			}
		}
	l103:
		if !(v6 == 0) {
			goto l105
		}
		v22 = v21
		goto l99
	l105:
		if v21 == i32(40) {
			m.fn36(i32(40), i32(40), i32(1099776))
			panic("unreachable")
		}
		store32(m.memory[uint32(v5+i32(176)+v14):], uint32(int32(v6)))
		v22 = v21 + i32(1)
	l99:
		store32(m.memory[int64(uint32(v5))+336:], uint32(v22))
		p174 := v15
		if uint32(v22) > uint32(v15) {
			p174 = v22
		}
		v1 = p174
		if uint32(v1) >= uint32(i32(41)) {
			m.fn124(i32(0), v1, i32(40), i32(1099776))
			panic("unreachable")
		}
		v1 = v1 << 2
		v11 = v5 + i32(8) + i32(-4)
		v14 = v5 + i32(176) + i32(-4)
		{
			{
			l109:
				{
					if v1 == 0 {
						goto l108
					}
					v9 = v14 + v1
					v13 = v11 + v1
					v1 = v1 + i32(-4)
					t175 := int32(load32(m.memory[uint32(v13):]))
					v13 = t175
					t176 := int32(load32(m.memory[uint32(v9):]))
					t177 := v13
					v9 = t176
					if t177 == v9 {
						goto l109
					}
				}
				var p178 int32
				if uint32(v13) > uint32(v9) {
					p178 = 1
				}
				var p179 int32
				if uint32(v13) < uint32(v9) {
					p179 = 1
				}
				switch (p178 - p179) & i32(255) {
				case 0:
					goto l108
				case 1:
					goto l110
				default:
					goto l111
				}
			}
		l108:
			v1 = i32(0)
			if v12 != 0 {
				goto l112
			}
			v1 = v24 + i32(-1)
			if uint32(v1) >= uint32(v3) {
				m.fn36(v1, v3, i32(1121408))
				panic("unreachable")
			}
			t180 := int32(m.memory[uint32(v2+v1)])
			if t180&i32(1) == 0 {
				goto l111
			}
		}
	l110:
		if uint32(v24) > uint32(v3) {
			m.fn124(i32(0), v24, v3, i32(1121424))
			panic("unreachable")
		}
		v11 = v2 + v24
		v1 = v24
		{
		l116:
			{
				v9 = v1
				if v9 == 0 {
					v1 = i32(49)
					if v12 != 0 {
						goto l117
					}
					m.memory[uint32(v2)] = byte(i32(49))
					v1 = i32(48)
					v9 = v24 + i32(-1)
					if v9 == 0 {
						goto l117
					}
					memory_fill(m.memory, uint32(v2+i32(1)), i32(48), uint32(v9))
				l117:
					v10 = v10 + i32(1)
					if v23 != 0 {
						goto l111
					}
					if uint32(v24) >= uint32(v3) {
						goto l111
					}
					m.memory[uint32(v11)] = byte(v1)
					v24 = v24 + i32(1)
					goto l111
				}
				v1 = v9 + i32(-1)
				v13 = v1 + v2
				t181 := int32(m.memory[uint32(v13)])
				if t181 == i32(57) {
					goto l116
				}
			}
			t182 := int32(m.memory[uint32(v13)])
			m.memory[uint32(v13)] = byte(t182 + i32(1))
			v1 = v24 - v9
			if v1 == 0 {
				goto l111
			}
			memory_fill(m.memory, uint32(v2+v9), i32(48), uint32(v1))
			goto l111
		}
	}
l111:
	if uint32(v24) > uint32(v3) {
		m.fn124(i32(0), v24, v3, i32(1121440))
		panic("unreachable")
	}
	v1 = v24
l112:
	store16(m.memory[int64(uint32(v0))+8:], uint16(v10))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
	goto l98
l98:
	store32(m.memory[uint32(v0):], uint32(v2))
	m.g0 = v5 + i32(832)
}
func (m *Module) fn872(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7 int32
	var v8, v9 int64
	{
		if uint32(v1) < uint32(i32(8)) {
			t12 := int32(load32(m.memory[int64(uint32(v0))+160:]))
			v6 = t12
			if uint32(v6) > uint32(i32(40)) {
				m.fn124(i32(0), v6, i32(40), i32(1099776))
				panic("unreachable")
			}
			if v6 != 0 {
				t13 := int64(load32(m.memory[int64(uint32(v1<<2))+1121648:]))
				v8 = t13
				v7 = v6 << 2
				v5 = v7 + i32(-4)
				v1 = int32(uint32(v5)>>2) + i32(1)
				v3 = v1 & i32(3)
				v9 = i64(0)
				v2 = v0
				if uint32(v5) < uint32(i32(12)) {
					goto l12
				}
				v5 = v1 & i32(0x7ffffffc)
				v9 = i64(0)
				v2 = v0
			l13:
				{
					t14 := int64(load32(m.memory[uint32(v2):]))
					t15 := v2
					v9 = t14*v8 + v9
					store32(m.memory[uint32(t15):], uint32(v9))
					v1 = v2 + i32(4)
					t16 := int64(load32(m.memory[uint32(v1):]))
					t17 := v1
					v9 = t16*v8 + int64(uint64(v9)>>32)
					store32(m.memory[uint32(t17):], uint32(v9))
					v1 = v2 + i32(8)
					t18 := int64(load32(m.memory[uint32(v1):]))
					t19 := v1
					v9 = t18*v8 + int64(uint64(v9)>>32)
					store32(m.memory[uint32(t19):], uint32(v9))
					v1 = v2 + i32(12)
					t20 := int64(load32(m.memory[uint32(v1):]))
					t21 := v1
					v9 = t20*v8 + int64(uint64(v9)>>32)
					store32(m.memory[uint32(t21):], uint32(v9))
					v9 = int64(uint64(v9) >> 32)
					v2 = v2 + i32(16)
					v5 = v5 + i32(-4)
					if v5 != 0 {
						goto l13
					}
				}
				if v3 == 0 {
					goto l14
				}
			l12:
				v5 = v3 << 2
			l15:
				{
					t22 := int64(load32(m.memory[uint32(v2):]))
					t23 := v2
					v9 = t22*v8 + v9
					store32(m.memory[uint32(t23):], uint32(v9))
					v2 = v2 + i32(4)
					v9 = int64(uint64(v9) >> 32)
					v5 = v5 + i32(-4)
					if v5 != 0 {
						goto l15
					}
				}
			l14:
				if v9 == 0 {
					goto l16
				}
				if v6 == i32(40) {
					m.fn36(i32(40), i32(40), i32(1099776))
					panic("unreachable")
				}
				store32(m.memory[uint32(v0+v7):], uint32(int32(v9)))
				v6 = v6 + i32(1)
			l16:
				store32(m.memory[int64(uint32(v0))+160:], uint32(v6))
				return v0
			}
			store32(m.memory[int64(uint32(v0))+160:], uint32(i32(0)))
			return v0
		}
		v2 = v1 & i32(7)
		if v2 == 0 {
			goto l1
		}
		t0 := int32(load32(m.memory[int64(uint32(v0))+160:]))
		v3 = t0
		if uint32(v3) > uint32(i32(40)) {
			m.fn124(i32(0), v3, i32(40), i32(1099776))
			panic("unreachable")
		}
		if v3 != 0 {
			v4 = v3 << 2
			v5 = v4 + i32(-4)
			v6 = int32(uint32(v5)>>2) + i32(1)
			v7 = v6 & i32(3)
			t1 := int32(load32(m.memory[int64(uint32(v2<<2))+1121648:]))
			v8 = int64(uint32(i32_shr_u(t1, v2)))
			v9 = i64(0)
			v2 = v0
			if uint32(v5) < uint32(i32(12)) {
				goto l4
			}
			v5 = v6 & i32(0x7ffffffc)
			v9 = i64(0)
			v2 = v0
		l5:
			{
				t2 := int64(load32(m.memory[uint32(v2):]))
				t3 := v2
				v9 = t2*v8 + v9
				store32(m.memory[uint32(t3):], uint32(v9))
				v6 = v2 + i32(4)
				t4 := int64(load32(m.memory[uint32(v6):]))
				t5 := v6
				v9 = t4*v8 + int64(uint64(v9)>>32)
				store32(m.memory[uint32(t5):], uint32(v9))
				v6 = v2 + i32(8)
				t6 := int64(load32(m.memory[uint32(v6):]))
				t7 := v6
				v9 = t6*v8 + int64(uint64(v9)>>32)
				store32(m.memory[uint32(t7):], uint32(v9))
				v6 = v2 + i32(12)
				t8 := int64(load32(m.memory[uint32(v6):]))
				t9 := v6
				v9 = t8*v8 + int64(uint64(v9)>>32)
				store32(m.memory[uint32(t9):], uint32(v9))
				v9 = int64(uint64(v9) >> 32)
				v2 = v2 + i32(16)
				v5 = v5 + i32(-4)
				if v5 != 0 {
					goto l5
				}
			}
			if v7 == 0 {
				goto l6
			}
		l4:
			v5 = v7 << 2
		l7:
			{
				t10 := int64(load32(m.memory[uint32(v2):]))
				t11 := v2
				v9 = t10*v8 + v9
				store32(m.memory[uint32(t11):], uint32(v9))
				v2 = v2 + i32(4)
				v9 = int64(uint64(v9) >> 32)
				v5 = v5 + i32(-4)
				if v5 != 0 {
					goto l7
				}
			}
		l6:
			if v9 == 0 {
				goto l8
			}
			if v3 == i32(40) {
				m.fn36(i32(40), i32(40), i32(1099776))
				panic("unreachable")
			}
			store32(m.memory[uint32(v0+v4):], uint32(int32(v9)))
			v3 = v3 + i32(1)
		l8:
			store32(m.memory[int64(uint32(v0))+160:], uint32(v3))
			goto l1
		}
		store32(m.memory[int64(uint32(v0))+160:], uint32(i32(0)))
		goto l1
	}
l1:
	{
		if v1&i32(8) == 0 {
			goto l18
		}
		t24 := int32(load32(m.memory[int64(uint32(v0))+160:]))
		v3 = t24
		if uint32(v3) > uint32(i32(40)) {
			m.fn124(i32(0), v3, i32(40), i32(1099776))
			panic("unreachable")
		}
		if v3 != 0 {
			goto l20
		}
		v3 = i32(0)
		goto l21
	l20:
		v4 = v3 << 2
		v5 = v4 + i32(-4)
		v6 = int32(uint32(v5)>>2) + i32(1)
		v7 = v6 & i32(3)
		v8 = i64(0)
		v2 = v0
		if uint32(v5) < uint32(i32(12)) {
			goto l22
		}
		v5 = v6 & i32(0x7ffffffc)
		v8 = i64(0)
		v2 = v0
	l23:
		{
			t25 := int64(load32(m.memory[uint32(v2):]))
			t26 := v2
			v8 = t25*i64(390625) + v8
			store32(m.memory[uint32(t26):], uint32(v8))
			v6 = v2 + i32(4)
			t27 := int64(load32(m.memory[uint32(v6):]))
			t28 := v6
			v8 = t27*i64(390625) + int64(uint64(v8)>>32)
			store32(m.memory[uint32(t28):], uint32(v8))
			v6 = v2 + i32(8)
			t29 := int64(load32(m.memory[uint32(v6):]))
			t30 := v6
			v8 = t29*i64(390625) + int64(uint64(v8)>>32)
			store32(m.memory[uint32(t30):], uint32(v8))
			v6 = v2 + i32(12)
			t31 := int64(load32(m.memory[uint32(v6):]))
			t32 := v6
			v8 = t31*i64(390625) + int64(uint64(v8)>>32)
			store32(m.memory[uint32(t32):], uint32(v8))
			v8 = int64(uint64(v8) >> 32)
			v2 = v2 + i32(16)
			v5 = v5 + i32(-4)
			if v5 != 0 {
				goto l23
			}
		}
		if v7 == 0 {
			goto l24
		}
	l22:
		v5 = v7 << 2
	l25:
		{
			t33 := int64(load32(m.memory[uint32(v2):]))
			t34 := v2
			v8 = t33*i64(390625) + v8
			store32(m.memory[uint32(t34):], uint32(v8))
			v2 = v2 + i32(4)
			v8 = int64(uint64(v8) >> 32)
			v5 = v5 + i32(-4)
			if v5 != 0 {
				goto l25
			}
		}
	l24:
		if v8 == 0 {
			goto l21
		}
		if v3 == i32(40) {
			m.fn36(i32(40), i32(40), i32(1099776))
			panic("unreachable")
		}
		store32(m.memory[uint32(v0+v4):], uint32(int32(v8)))
		v3 = v3 + i32(1)
	l21:
		store32(m.memory[int64(uint32(v0))+160:], uint32(v3))
	}
l18:
	if v1&i32(16) == 0 {
		goto l27
	}
	_ = m.fn858(v0, i32(1121688), i32(2))
l27:
	if v1&i32(32) == 0 {
		goto l28
	}
	_ = m.fn858(v0, i32(1121696), i32(3))
l28:
	if v1&i32(64) == 0 {
		goto l29
	}
	_ = m.fn858(v0, i32(1121708), i32(5))
l29:
	if v1&i32(128) == 0 {
		goto l30
	}
	_ = m.fn858(v0, i32(1121728), i32(10))
l30:
	if v1&i32(256) == 0 {
		goto l31
	}
	_ = m.fn858(v0, i32(1121768), i32(19))
l31:
	_ = m.fn859(v0, v1)
	return v0
}
func (m *Module) fn873(v0, v1, v2 int32) {
	var v3 int32
	var v4, v5, v6, v7, v8, v9 int64
	var v10 int32
	var v11, v12, v13, v14 int64
	var v15 int32
	var v16, v17, v18, v19, v20 int64
	var v21, v22 int32
	var v23 int64
	var v24, v25, v26, v27 int32
	var v28, v29 int64
	t0 := m.g0
	v3 = t0 - i32(80)
	m.g0 = v3
	{
		t1 := int64(load64(m.memory[uint32(v1):]))
		v4 = t1
		if v4 == i64(0) {
			m.fn7(i32(0x111a88), i32(28), i32(1121012))
			panic("unreachable")
		}
		t2 := int64(load64(m.memory[int64(uint32(v1))+8:]))
		v5 = t2
		if v5 == i64(0) {
			m.fn7(i32(1121028), i32(29), i32(1121060))
			panic("unreachable")
		}
		t3 := int64(load64(m.memory[int64(uint32(v1))+16:]))
		v6 = t3
		if v6 == i64(0) {
			m.fn7(i32(1121076), i32(28), i32(1121104))
			panic("unreachable")
		}
		v7 = v6 + v4
		if uint64(v7) < uint64(v6) {
			m.fn7(i32(1121288), i32(54), i32(1121344))
			panic("unreachable")
		}
		if uint64(v4) < uint64(v5) {
			m.fn7(i32(1121216), i32(55), i32(1121272))
			panic("unreachable")
		}
		if uint64(v7) >= uint64(i64(0x2000000000000000)) {
			m.fn7(i32(1121120), i32(45), i32(1121168))
			panic("unreachable")
		}
		t4 := int32(load16(m.memory[int64(uint32(v1))+24:]))
		t5 := v3
		v1 = t4
		store16(m.memory[int64(uint32(t5))+64:], uint16(v1))
		t6 := v3
		v5 = v4 - v5
		store64(m.memory[int64(uint32(t6))+56:], uint64(v5))
		t7 := v3
		t8 := v5
		v6 = int64(bits.LeadingZeros64(uint64(v7)))
		v8 = i64_shl(t8, v6)
		v9 = i64_shr_u(v8, v6)
		store64(m.memory[int64(uint32(t7))+72:], uint64(v9))
		if v9 != v5 {
			m.fn854(v3+i32(72), v3+i32(56))
			panic("unreachable")
		}
		store16(m.memory[int64(uint32(v3))+64:], uint16(v1))
		store64(m.memory[int64(uint32(v3))+56:], uint64(v4))
		t9 := v3
		v9 = i64_shl(v4, v6)
		v5 = i64_shr_u(v9, v6)
		store64(m.memory[int64(uint32(t9))+72:], uint64(v5))
		if v5 != v4 {
			m.fn854(v3+i32(72), v3+i32(56))
			panic("unreachable")
		}
		v10 = v1 - int32(v6)
		v1 = (int32(int16(i32(-96)-v10))*i32(80) + i32(86960)) / i32(2126)
		if uint32(v1) > uint32(i32(80)) {
			m.fn36(v1, i32(81), i32(1120948))
			panic("unreachable")
		}
		t10 := v3 + i32(32)
		v1 = v1 << 4
		t11 := int64(load64(m.memory[int64(uint32(v1))+1119608:]))
		v4 = t11
		m.fn1911(t10, v4, i64(0), i64_shl(v7, v6), i64(0))
		m.fn1911(v3+i32(16), v4, i64(0), v8, i64(0))
		m.fn1911(v3, v4, i64(0), v9, i64(0))
		t12 := int32(load16(m.memory[int64(uint32(v1))+1119616:]))
		v10 = i32(0) - (v10 + t12)
		v4 = int64(uint32(v10))
		v8 = i64_shl(i64(1), v4)
		v11 = v8 + i64(-1)
		t13 := int64(load64(m.memory[int64(uint32(v3))+16:]))
		v12 = t13 >> 63
		t14 := int64(load64(m.memory[uint32(v3):]))
		v13 = int64(uint64(t14) >> 63)
		t15 := int64(load64(m.memory[int64(uint32(v3))+8:]))
		v14 = t15
		t16 := int32(load16(m.memory[int64(uint32(v1))+1119618:]))
		v1 = t16
		v15 = v10 & i32(63)
		t17 := int64(load64(m.memory[int64(uint32(v3))+24:]))
		v16 = t17
		{
			t18 := int64(load64(m.memory[int64(uint32(v3))+40:]))
			v17 = t18
			t19 := int64(load64(m.memory[int64(uint32(v3))+32:]))
			t20 := v17
			v18 = int64(uint64(t19) >> 63)
			v19 = t20 + v18
			v20 = v19 + i64(1)
			v21 = int32(i64_shr_u(v20, v4))
			if uint32(v21) < uint32(i32(10000)) {
				if uint32(v21) < uint32(i32(100)) {
					var p30 int32
					if uint32(v21) > uint32(i32(9)) {
						p30 = 1
					}
					v22 = p30
					p31 := i32(1)
					if v22 != 0 {
						p31 = i32(10)
					}
					v10 = p31
					goto l12
				}
				var p27 int32
				if uint32(v21) < uint32(i32(1000)) {
					p27 = 1
				}
				v10 = p27
				p28 := i32(3)
				if v10 != 0 {
					p28 = i32(2)
				}
				v22 = p28
				p29 := i32(1000)
				if v10 != 0 {
					p29 = i32(100)
				}
				v10 = p29
				goto l12
			}
			if uint32(v21) < uint32(i32(1000000)) {
				var p32 int32
				if uint32(v21) < uint32(i32(100000)) {
					p32 = 1
				}
				v10 = p32
				p33 := i32(5)
				if v10 != 0 {
					p33 = i32(4)
				}
				v22 = p33
				p34 := i32(100000)
				if v10 != 0 {
					p34 = i32(10000)
				}
				v10 = p34
				goto l12
			}
			{
				if uint32(v21) < uint32(i32(100000000)) {
					var p24 int32
					if uint32(v21) < uint32(i32(10000000)) {
						p24 = 1
					}
					v10 = p24
					p25 := i32(7)
					if v10 != 0 {
						p25 = i32(6)
					}
					v22 = p25
					p26 := i32(10000000)
					if v10 != 0 {
						p26 = i32(1000000)
					}
					v10 = p26
					goto l12
				}
				var p21 int32
				if uint32(v21) < uint32(i32(1000000000)) {
					p21 = 1
				}
				v10 = p21
				p22 := i32(9)
				if v10 != 0 {
					p22 = i32(8)
				}
				v22 = p22
				p23 := i32(1000000000)
				if v10 != 0 {
					p23 = i32(100000000)
				}
				v10 = p23
				goto l12
			}
		}
	}
l12:
	v4 = v20 & v11
	v23 = v13 + v14
	v6 = int64(uint32(v15))
	v24 = v22 - v1 + i32(1)
	v9 = v12 - v16 + v20 + i64(1)
	v5 = v9 & v11
	v1 = i32(0)
l20:
	{
		v25 = v2 + v1
		t35 := int32(uint32(v21) / uint32(v10))
		t36 := v25
		v15 = t35
		v26 = v15 + i32(48)
		m.memory[uint32(t36)] = byte(v26)
		v27 = v1 + i32(1)
		t37 := v9
		v21 = v21 - v15*v10
		v28 = i64_shl(int64(uint32(v21)), v6)
		v7 = v28 + v4
		if uint64(t37) > uint64(v7) {
			v8 = v9 - v7
			t46 := v8
			v6 = i64_shl(int64(uint32(v10)), v6)
			var p47 int32
			if uint64(t46) < uint64(v6) {
				p47 = 1
			}
			v1 = p47
			v5 = v20 - v23
			v29 = v5 + i64(1)
			t48 := v7
			v11 = v5 + i64(-1)
			if uint64(t48) >= uint64(v11) {
				goto l24
			}
			if uint64(v8) < uint64(v6) {
				goto l24
			}
			t49 := v19 + v12 - v16
			v4 = v4 + v6
			v20 = t49 - (v4 + v28) + i64(2)
			v12 = v19 - v23 - v7
			v8 = v4 + v13 + v14 - v18 - v17 + v28
			v4 = i64(0)
		l27:
			{
				v5 = v7 + v6
				if uint64(v5) < uint64(v11) {
					goto l25
				}
				if uint64(v12+v4) >= uint64(v8) {
					goto l25
				}
				v1 = i32(0)
				goto l24
			l25:
				t50 := v25
				v26 = v26 + i32(-1)
				m.memory[uint32(t50)] = byte(v26)
				v28 = v20 + v4
				var p51 int32
				if uint64(v28) < uint64(v6) {
					p51 = 1
				}
				v1 = p51
				if uint64(v5) >= uint64(v11) {
					goto l26
				}
				v8 = v8 + v6
				v4 = v4 - v6
				v7 = v5
				if uint64(v28) < uint64(v6) {
					goto l26
				}
				goto l27
			}
		}
		{
			if v22 != v1 {
				goto l15
			}
			v7 = i64(1)
		l17:
			{
				v9 = v7
				v10 = v1
				if v10 == i32(16) {
					m.fn36(i32(17), i32(17), i32(1121200))
					panic("unreachable")
				}
				t38 := v2 + v10 + i32(1)
				v4 = v4 * i64(10)
				v21 = int32(i64_shr_u(v4, v6)) + i32(48)
				m.memory[uint32(t38)] = byte(v21)
				v7 = v9 * i64(10)
				v1 = v10 + i32(1)
				v5 = v5 * i64(10)
				t39 := v5
				v4 = v4 & v11
				if uint64(t39) <= uint64(v4) {
					goto l17
				}
			}
			v28 = v5 - v4
			var p40 int32
			if uint64(v28) < uint64(v8) {
				p40 = 1
			}
			v15 = p40
			v6 = v7 * (v20 - v23)
			v13 = v6 + v7
			t41 := v4
			v11 = v6 - v7
			if uint64(t41) >= uint64(v11) {
				goto l18
			}
			if uint64(v28) >= uint64(v8) {
				v1 = v2 + v1
				v20 = v5 - v8
				v12 = v8 - v11
				v6 = i64(0) - v4
			l23:
				{
					v7 = v4 + v8
					if uint64(v7) < uint64(v11) {
						goto l21
					}
					if uint64(v11+v6) >= uint64(v12+v4) {
						goto l21
					}
					v15 = i32(0)
					goto l18
				l21:
					t44 := v1
					v21 = v21 + i32(-1)
					m.memory[uint32(t44)] = byte(v21)
					v28 = v20 + v6
					var p45 int32
					if uint64(v28) < uint64(v8) {
						p45 = 1
					}
					v15 = p45
					if uint64(v7) >= uint64(v11) {
						goto l22
					}
					v6 = v6 - v8
					v4 = v7
					if uint64(v28) < uint64(v8) {
						goto l22
					}
					goto l23
				}
			}
			goto l18
		}
	l15:
		;
		var p42 int32
		if uint32(v10) < uint32(i32(10)) {
			p42 = 1
		}
		v15 = p42
		t43 := int32(uint32(v10) / uint32(i32(10)))
		v10 = t43
		v1 = v27
		if v15 == 0 {
			goto l20
		}
	}
	m.fn687(i32(1121184))
	panic("unreachable")
l24:
	v5 = v7
l26:
	if uint64(v29) <= uint64(v5) {
		goto l28
	}
	if v1 != 0 {
		goto l28
	}
	v4 = v5 + v6
	if uint64(v4) < uint64(v29) {
		goto l29
	}
	if uint64(v29-v5) < uint64(v4-v29) {
		goto l28
	}
l29:
	store32(m.memory[uint32(v0):], uint32(i32(0)))
	goto l30
l28:
	if uint64(v5) < uint64(i64(2)) {
		goto l31
	}
	if uint64(v5) <= uint64(v9+i64(-4)) {
		store16(m.memory[int64(uint32(v0))+8:], uint16(v24))
		store32(m.memory[int64(uint32(v0))+4:], uint32(v27))
		goto l33
	}
l31:
	store32(m.memory[uint32(v0):], uint32(i32(0)))
	goto l30
l18:
	v7 = v4
l22:
	if uint64(v13) <= uint64(v7) {
		goto l34
	}
	if v15 != 0 {
		goto l34
	}
	v4 = v7 + v8
	if uint64(v4) < uint64(v13) {
		goto l35
	}
	if uint64(v13-v7) < uint64(v4-v13) {
		goto l34
	}
l35:
	store32(m.memory[uint32(v0):], uint32(i32(0)))
	goto l30
l34:
	if uint64(v9*i64(20)) > uint64(v7) {
		goto l36
	}
	if uint64(v7) <= uint64(v5+v9*i64(-40)) {
		goto l37
	}
l36:
	store32(m.memory[uint32(v0):], uint32(i32(0)))
	goto l30
l37:
	store16(m.memory[int64(uint32(v0))+8:], uint16(v24))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v10+i32(2)))
l33:
	store32(m.memory[uint32(v0):], uint32(v2))
l30:
	m.g0 = v3 + i32(80)
}
func (m *Module) fn874(v0, v1, v2 int32) {
	var v3 int32
	var v4, v5, v6, v7 int64
	var v8, v9, v10, v11, v12, v13, v14, v15, v16, v17, v18, v19, v20, v21, v22, v23, v24, v25, v26, v27, v28, v29, v30, v31, v32, v33, v34, v35, v36, v37 int32
	t0 := m.g0
	v3 = t0 - i32(1328)
	m.g0 = v3
	{
		t1 := int64(load64(m.memory[uint32(v1):]))
		v4 = t1
		if v4 == i64(0) {
			m.fn7(i32(0x111a88), i32(28), i32(1121520))
			panic("unreachable")
		}
		t2 := int64(load64(m.memory[int64(uint32(v1))+8:]))
		v5 = t2
		if v5 == i64(0) {
			m.fn7(i32(1121028), i32(29), i32(0x111d00))
			panic("unreachable")
		}
		t3 := int64(load64(m.memory[int64(uint32(v1))+16:]))
		v6 = t3
		if v6 == i64(0) {
			m.fn7(i32(1121076), i32(28), i32(1121552))
			panic("unreachable")
		}
		v7 = v6 + v4
		if uint64(v7) < uint64(v6) {
			m.fn7(i32(1121288), i32(54), i32(1121632))
			panic("unreachable")
		}
		if uint64(v4) < uint64(v5) {
			m.fn7(i32(1121216), i32(55), i32(1121616))
			panic("unreachable")
		}
		t4 := int32(int8(m.memory[int64(uint32(v1))+26]))
		v8 = t4
		t5 := int32(int16(load16(m.memory[int64(uint32(v1))+24:])))
		v1 = t5
		store64(m.memory[int64(uint32(v3))+8:], uint64(v4))
		t7 := v3
		p6 := i32(2)
		if uint64(v4) < uint64(i64(0x100000000)) {
			p6 = i32(1)
		}
		store32(m.memory[int64(uint32(t7))+168:], uint32(p6))
		memory_zero(m.memory, uint32(v3+i32(8)+i32(8)), uint32(i32(152)))
		store64(m.memory[int64(uint32(v3))+176:], uint64(v5))
		t9 := v3
		p8 := i32(2)
		if uint64(v5) < uint64(i64(0x100000000)) {
			p8 = i32(1)
		}
		store32(m.memory[int64(uint32(t9))+336:], uint32(p8))
		memory_zero(m.memory, uint32(v3+i32(176)+i32(8)), uint32(i32(152)))
		store64(m.memory[int64(uint32(v3))+344:], uint64(v6))
		t11 := v3
		p10 := i32(2)
		if uint64(v6) < uint64(i64(0x100000000)) {
			p10 = i32(1)
		}
		store32(m.memory[int64(uint32(t11))+504:], uint32(p10))
		memory_zero(m.memory, uint32(v3+i32(344)+i32(8)), uint32(i32(152)))
		memory_zero(m.memory, uint32(v3+i32(512)), uint32(i32(156)))
		store32(m.memory[int64(uint32(v3))+508:], uint32(i32(1)))
		store32(m.memory[int64(uint32(v3))+668:], uint32(i32(1)))
		v9 = int32(int64(uint64((int64(v1)-int64(bits.LeadingZeros64(uint64(v7+i64(-1)))))*i64(1292913986)+i64(82746495104)) >> 32))
		v10 = int32(int16(v9))
		if v1 < i32(0) {
			goto l5
		}
		_ = m.fn859(v3+i32(8), v1)
		_ = m.fn859(v3+i32(176), v1)
		_ = m.fn859(v3+i32(344), v1)
		goto l6
	l5:
		_ = m.fn859(v3+i32(508), int32(int16(i32(0)-v1)))
	l6:
		{
			if v10 > i32(-1) {
				goto l7
			}
			t16 := v3 + i32(8)
			v1 = (i32(0) - v10) & i32(0xffff)
			_ = m.fn872(t16, v1)
			_ = m.fn872(v3+i32(176), v1)
			_ = m.fn872(v3+i32(344), v1)
			goto l8
		}
	l7:
		_ = m.fn872(v3+i32(508), v9&i32(0x7fff))
	l8:
		memory_copy(m.memory, uint32(v3+i32(1164)), uint32(v3+i32(8)), uint32(i32(164)))
		t21 := int32(load32(m.memory[int64(uint32(v3))+504:]))
		v11 = t21
		t22 := int32(load32(m.memory[int64(uint32(v3))+1324:]))
		t23 := v11
		v1 = t22
		p24 := v1
		if uint32(v11) > uint32(v1) {
			p24 = t23
		}
		v12 = p24
		if uint32(v12) > uint32(i32(40)) {
			m.fn124(i32(0), v12, i32(40), i32(1099776))
			panic("unreachable")
		}
		if v12 != 0 {
			goto l10
		}
		v12 = i32(0)
		goto l11
	l10:
		v13 = i32(0)
		v14 = i32(0)
		{
			if v12 == i32(1) {
				goto l12
			}
			v15 = v12 & i32(1)
			v16 = v12 & i32(62)
			v13 = i32(0)
			v1 = v3 + i32(1164)
			v9 = v3 + i32(344)
			v14 = i32(0)
		l13:
			{
				t25 := int32(load32(m.memory[uint32(v9):]))
				t26 := v1
				v17 = t25
				t27 := int32(load32(m.memory[uint32(v1):]))
				v18 = v17 + t27
				v19 = v18 + v13&i32(1)
				store32(m.memory[uint32(t26):], uint32(v19))
				v13 = v1 + i32(4)
				t28 := int32(load32(m.memory[uint32(v9+i32(4)):]))
				t29 := v13
				v20 = t28
				t30 := int32(load32(m.memory[uint32(v13):]))
				v13 = v20 + t30
				t31 := v13
				var p32 int32
				if uint32(v18) < uint32(v17) {
					p32 = 1
				}
				var p33 int32
				if uint32(v19) < uint32(v18) {
					p33 = 1
				}
				v18 = t31 + (p32 | p33)
				store32(m.memory[uint32(t29):], uint32(v18))
				var p34 int32
				if uint32(v13) < uint32(v20) {
					p34 = 1
				}
				var p35 int32
				if uint32(v18) < uint32(v13) {
					p35 = 1
				}
				v13 = p34 | p35
				v9 = v9 + i32(8)
				v1 = v1 + i32(8)
				t36 := v16
				v14 = v14 + i32(2)
				if t36 != v14 {
					goto l13
				}
			}
			if v15 == 0 {
				goto l14
			}
		l12:
			t37 := v3 + i32(1164)
			v1 = v14 << 2
			v9 = t37 + v1
			t38 := int32(load32(m.memory[uint32(v3+i32(344)+v1):]))
			t39 := v9
			v18 = t38
			t40 := int32(load32(m.memory[uint32(v9):]))
			v1 = v18 + t40
			v9 = v1 + v13
			store32(m.memory[uint32(t39):], uint32(v9))
			if uint32(v1) < uint32(v18) {
				goto l15
			}
			if uint32(v9) < uint32(v1) {
				goto l15
			}
			goto l11
		}
	l14:
		if v13 == 0 {
			goto l11
		}
	l15:
		if v12 == i32(40) {
			m.fn36(i32(40), i32(40), i32(1099776))
			panic("unreachable")
		}
		store32(m.memory[uint32(v3+i32(1164)+v12<<2):], uint32(i32(1)))
		v12 = v12 + i32(1)
	l11:
		store32(m.memory[int64(uint32(v3))+1324:], uint32(v12))
		t41 := int32(load32(m.memory[int64(uint32(v3))+668:]))
		t42 := v12
		v21 = t41
		p43 := v21
		if uint32(v12) > uint32(v21) {
			p43 = t42
		}
		v1 = p43
		if uint32(v1) >= uint32(i32(41)) {
			m.fn124(i32(0), v1, i32(40), i32(1099776))
			panic("unreachable")
		}
		v1 = v1 << 2
		v18 = v3 + i32(1164) + i32(-4)
		{
		l20:
			{
				if v1 != 0 {
					goto l18
				}
				v1 = i32(0)
				goto l19
			l18:
				v9 = v18 + v1
				v1 = v1 + i32(-4)
				t44 := int32(load32(m.memory[uint32(v1+(v3+i32(508))):]))
				v13 = t44
				t45 := int32(load32(m.memory[uint32(v9):]))
				t46 := v13
				v9 = t45
				if t46 == v9 {
					goto l20
				}
			}
			var p47 int32
			if uint32(v13) > uint32(v9) {
				p47 = 1
			}
			var p48 int32
			if uint32(v13) < uint32(v9) {
				p48 = 1
			}
			v1 = p47 - p48
		}
	l19:
		{
			if v1 < v8 {
				goto l21
			}
			{
				t49 := int32(load32(m.memory[int64(uint32(v3))+168:]))
				v13 = t49
				if uint32(v13) > uint32(i32(40)) {
					m.fn124(i32(0), v13, i32(40), i32(1099776))
					panic("unreachable")
				}
				if v13 != 0 {
					v17 = v13 << 2
					v1 = v17 + i32(-4)
					v9 = int32(uint32(v1)>>2) + i32(1)
					v14 = v9 & i32(3)
					if uint32(v1) >= uint32(i32(12)) {
						goto l25
					}
					v4 = i64(0)
					v1 = v3 + i32(8)
					goto l26
				}
				v13 = i32(0)
				goto l24
			}
		l25:
			v9 = v9 & i32(0x7ffffffc)
			v4 = i64(0)
			v1 = v3 + i32(8)
		l27:
			{
				t50 := int64(load32(m.memory[uint32(v1):]))
				t51 := v1
				v4 = t50*i64(10) + v4
				store32(m.memory[uint32(t51):], uint32(v4))
				v18 = v1 + i32(4)
				t52 := int64(load32(m.memory[uint32(v18):]))
				t53 := v18
				v4 = t52*i64(10) + int64(uint64(v4)>>32)
				store32(m.memory[uint32(t53):], uint32(v4))
				v18 = v1 + i32(8)
				t54 := int64(load32(m.memory[uint32(v18):]))
				t55 := v18
				v4 = t54*i64(10) + int64(uint64(v4)>>32)
				store32(m.memory[uint32(t55):], uint32(v4))
				v18 = v1 + i32(12)
				t56 := int64(load32(m.memory[uint32(v18):]))
				t57 := v18
				v4 = t56*i64(10) + int64(uint64(v4)>>32)
				store32(m.memory[uint32(t57):], uint32(v4))
				v4 = int64(uint64(v4) >> 32)
				v1 = v1 + i32(16)
				v9 = v9 + i32(-4)
				if v9 != 0 {
					goto l27
				}
			}
			if v14 == 0 {
				goto l28
			}
		l26:
			v9 = v14 << 2
		l29:
			{
				t58 := int64(load32(m.memory[uint32(v1):]))
				t59 := v1
				v4 = t58*i64(10) + v4
				store32(m.memory[uint32(t59):], uint32(v4))
				v1 = v1 + i32(4)
				v4 = int64(uint64(v4) >> 32)
				v9 = v9 + i32(-4)
				if v9 != 0 {
					goto l29
				}
			}
		l28:
			if v4 == 0 {
				goto l24
			}
			if v13 == i32(40) {
				m.fn36(i32(40), i32(40), i32(1099776))
				panic("unreachable")
			}
			store32(m.memory[uint32(v3+i32(8)+v17):], uint32(int32(v4)))
			v13 = v13 + i32(1)
		l24:
			store32(m.memory[int64(uint32(v3))+168:], uint32(v13))
			{
				t60 := int32(load32(m.memory[int64(uint32(v3))+336:]))
				v14 = t60
				if uint32(v14) > uint32(i32(40)) {
					m.fn124(i32(0), v14, i32(40), i32(1099776))
					panic("unreachable")
				}
				v22 = i32(0)
				v1 = i32(0)
				if v14 == 0 {
					goto l32
				}
				v19 = v14 << 2
				v1 = v19 + i32(-4)
				v9 = int32(uint32(v1)>>2) + i32(1)
				v17 = v9 & i32(3)
				if uint32(v1) >= uint32(i32(12)) {
					goto l33
				}
				v4 = i64(0)
				v1 = v3 + i32(176)
				goto l34
			}
		l33:
			v9 = v9 & i32(0x7ffffffc)
			v4 = i64(0)
			v1 = v3 + i32(176)
		l35:
			{
				t61 := int64(load32(m.memory[uint32(v1):]))
				t62 := v1
				v4 = t61*i64(10) + v4
				store32(m.memory[uint32(t62):], uint32(v4))
				v18 = v1 + i32(4)
				t63 := int64(load32(m.memory[uint32(v18):]))
				t64 := v18
				v4 = t63*i64(10) + int64(uint64(v4)>>32)
				store32(m.memory[uint32(t64):], uint32(v4))
				v18 = v1 + i32(8)
				t65 := int64(load32(m.memory[uint32(v18):]))
				t66 := v18
				v4 = t65*i64(10) + int64(uint64(v4)>>32)
				store32(m.memory[uint32(t66):], uint32(v4))
				v18 = v1 + i32(12)
				t67 := int64(load32(m.memory[uint32(v18):]))
				t68 := v18
				v4 = t67*i64(10) + int64(uint64(v4)>>32)
				store32(m.memory[uint32(t68):], uint32(v4))
				v4 = int64(uint64(v4) >> 32)
				v1 = v1 + i32(16)
				v9 = v9 + i32(-4)
				if v9 != 0 {
					goto l35
				}
			}
			if v17 == 0 {
				goto l36
			}
		l34:
			v9 = v17 << 2
		l37:
			{
				t69 := int64(load32(m.memory[uint32(v1):]))
				t70 := v1
				v4 = t69*i64(10) + v4
				store32(m.memory[uint32(t70):], uint32(v4))
				v1 = v1 + i32(4)
				v4 = int64(uint64(v4) >> 32)
				v9 = v9 + i32(-4)
				if v9 != 0 {
					goto l37
				}
			}
		l36:
			if !(v4 == 0) {
				goto l38
			}
			v1 = v14
			goto l32
		l38:
			if v14 == i32(40) {
				m.fn36(i32(40), i32(40), i32(1099776))
				panic("unreachable")
			}
			store32(m.memory[uint32(v3+i32(176)+v19):], uint32(int32(v4)))
			v1 = v14 + i32(1)
		l32:
			store32(m.memory[int64(uint32(v3))+336:], uint32(v1))
			if v11 == 0 {
				goto l40
			}
			v17 = v11 << 2
			v1 = v17 + i32(-4)
			v9 = int32(uint32(v1)>>2) + i32(1)
			v14 = v9 & i32(3)
			if uint32(v1) >= uint32(i32(12)) {
				goto l41
			}
			v4 = i64(0)
			v1 = v3 + i32(344)
			goto l42
		l41:
			v9 = v9 & i32(0x7ffffffc)
			v4 = i64(0)
			v1 = v3 + i32(344)
		l43:
			{
				t71 := int64(load32(m.memory[uint32(v1):]))
				t72 := v1
				v4 = t71*i64(10) + v4
				store32(m.memory[uint32(t72):], uint32(v4))
				v18 = v1 + i32(4)
				t73 := int64(load32(m.memory[uint32(v18):]))
				t74 := v18
				v4 = t73*i64(10) + int64(uint64(v4)>>32)
				store32(m.memory[uint32(t74):], uint32(v4))
				v18 = v1 + i32(8)
				t75 := int64(load32(m.memory[uint32(v18):]))
				t76 := v18
				v4 = t75*i64(10) + int64(uint64(v4)>>32)
				store32(m.memory[uint32(t76):], uint32(v4))
				v18 = v1 + i32(12)
				t77 := int64(load32(m.memory[uint32(v18):]))
				t78 := v18
				v4 = t77*i64(10) + int64(uint64(v4)>>32)
				store32(m.memory[uint32(t78):], uint32(v4))
				v4 = int64(uint64(v4) >> 32)
				v1 = v1 + i32(16)
				v9 = v9 + i32(-4)
				if v9 != 0 {
					goto l43
				}
			}
			if v14 == 0 {
				goto l44
			}
		l42:
			v9 = v14 << 2
		l45:
			{
				t79 := int64(load32(m.memory[uint32(v1):]))
				t80 := v1
				v4 = t79*i64(10) + v4
				store32(m.memory[uint32(t80):], uint32(v4))
				v1 = v1 + i32(4)
				v4 = int64(uint64(v4) >> 32)
				v9 = v9 + i32(-4)
				if v9 != 0 {
					goto l45
				}
			}
		l44:
			{
				if !(v4 == 0) {
					goto l46
				}
				t81 := v3
				v22 = v11
				store32(m.memory[int64(uint32(t81))+504:], uint32(v22))
				goto l47
			}
		l46:
			if v11 == i32(40) {
				m.fn36(i32(40), i32(40), i32(1099776))
				panic("unreachable")
			}
			store32(m.memory[uint32(v3+i32(344)+v17):], uint32(int32(v4)))
			v22 = v11 + i32(1)
		l40:
			store32(m.memory[int64(uint32(v3))+504:], uint32(v22))
			goto l47
		l21:
			v10 = v10 + i32(1)
			t82 := int32(load32(m.memory[int64(uint32(v3))+168:]))
			v13 = t82
			v22 = v11
		}
	l47:
		memory_copy(m.memory, uint32(v3+i32(672)), uint32(v3+i32(508)), uint32(i32(164)))
		t83 := m.fn859(v3+i32(672), i32(1))
		v23 = t83
		memory_copy(m.memory, uint32(v3+i32(836)), uint32(v3+i32(508)), uint32(i32(164)))
		t84 := m.fn859(v3+i32(836), i32(2))
		v24 = t84
		memory_copy(m.memory, uint32(v3+i32(1000)), uint32(v3+i32(508)), uint32(i32(164)))
		{
			{
				t85 := m.fn859(v3+i32(1000), i32(3))
				v25 = t85
				t86 := int32(load32(m.memory[int64(uint32(v25))+160:]))
				v26 = t86
				p87 := v13
				if uint32(v26) > uint32(v13) {
					p87 = v26
				}
				v27 = p87
				if uint32(v27) > uint32(i32(40)) {
					goto l49
				}
				v15 = v3 + i32(508) + i32(-4)
				v28 = v3 + i32(1164) + i32(-4)
				v11 = v3 + i32(672) + i32(-4)
				v12 = v3 + i32(836) + i32(-4)
				v16 = v3 + i32(1000) + i32(-4)
				t88 := int32(load32(m.memory[int64(uint32(v23))+160:]))
				v29 = t88
				t89 := int32(load32(m.memory[int64(uint32(v24))+160:]))
				v30 = t89
				v31 = i32(0)
			l128:
				{
					v32 = v31
					v1 = v27 << 2
				l51:
					{
						if v1 == 0 {
							goto l50
						}
						v9 = v16 + v1
						v1 = v1 + i32(-4)
						t90 := int32(load32(m.memory[uint32(v1+(v3+i32(8))):]))
						v18 = t90
						t91 := int32(load32(m.memory[uint32(v9):]))
						t92 := v18
						v9 = t91
						if t92 == v9 {
							goto l51
						}
					}
					v33 = i32(0)
					if uint32(v18) < uint32(v9) {
						goto l52
					}
				l50:
					if v27 == 0 {
						goto l53
					}
					v13 = i32(1)
					v14 = i32(0)
					{
						if v27 == i32(1) {
							goto l54
						}
						v34 = v27 & i32(1)
						v35 = v27 & i32(62)
						v14 = i32(0)
						v13 = i32(1)
						v1 = v3 + i32(8)
						v9 = v3 + i32(1000)
					l55:
						{
							t93 := int32(load32(m.memory[uint32(v1):]))
							t94 := v1
							v17 = t93
							t95 := int32(load32(m.memory[uint32(v9):]))
							v18 = v17 + (t95 ^ i32(-1))
							v19 = v18 + v13&i32(1)
							store32(m.memory[uint32(t94):], uint32(v19))
							v13 = v1 + i32(4)
							t96 := int32(load32(m.memory[uint32(v13):]))
							t97 := v13
							v20 = t96
							t98 := int32(load32(m.memory[uint32(v9+i32(4)):]))
							v13 = v20 + (t98 ^ i32(-1))
							t99 := v13
							var p100 int32
							if uint32(v18) < uint32(v17) {
								p100 = 1
							}
							var p101 int32
							if uint32(v19) < uint32(v18) {
								p101 = 1
							}
							v18 = t99 + (p100 | p101)
							store32(m.memory[uint32(t97):], uint32(v18))
							var p102 int32
							if uint32(v13) < uint32(v20) {
								p102 = 1
							}
							var p103 int32
							if uint32(v18) < uint32(v13) {
								p103 = 1
							}
							v13 = p102 | p103
							v9 = v9 + i32(8)
							v1 = v1 + i32(8)
							t104 := v35
							v14 = v14 + i32(2)
							if t104 != v14 {
								goto l55
							}
						}
						if v34 == 0 {
							goto l56
						}
					l54:
						t105 := v3 + i32(8)
						v1 = v14 << 2
						v9 = t105 + v1
						t106 := int32(load32(m.memory[uint32(v9):]))
						t107 := v9
						v9 = t106
						t108 := int32(load32(m.memory[uint32(v25+v1):]))
						v1 = v9 + (t108 ^ i32(-1))
						v18 = v1 + v13
						store32(m.memory[uint32(t107):], uint32(v18))
						if uint32(v1) < uint32(v9) {
							goto l53
						}
						if uint32(v18) < uint32(v1) {
							goto l53
						}
						goto l57
					}
				l56:
					if v13 == 0 {
						goto l57
					}
				l53:
					store32(m.memory[int64(uint32(v3))+168:], uint32(v27))
					v33 = i32(8)
					v13 = v27
				l52:
					p109 := v13
					if uint32(v30) > uint32(v13) {
						p109 = v30
					}
					v27 = p109
					if uint32(v27) >= uint32(i32(41)) {
						m.fn124(i32(0), v27, i32(40), i32(1099776))
						panic("unreachable")
					}
					v1 = v27 << 2
				l60:
					{
						if v1 == 0 {
							goto l59
						}
						v9 = v12 + v1
						v1 = v1 + i32(-4)
						t110 := int32(load32(m.memory[uint32(v1+(v3+i32(8))):]))
						v18 = t110
						t111 := int32(load32(m.memory[uint32(v9):]))
						t112 := v18
						v9 = t111
						if t112 == v9 {
							goto l60
						}
					}
					if uint32(v18) >= uint32(v9) {
						goto l59
					}
					v27 = v13
					goto l61
				l59:
					if v27 == 0 {
						goto l62
					}
					v13 = i32(1)
					v14 = i32(0)
					{
						if v27 == i32(1) {
							goto l63
						}
						v34 = v27 & i32(1)
						v35 = v27 & i32(62)
						v14 = i32(0)
						v13 = i32(1)
						v1 = v3 + i32(8)
						v9 = v3 + i32(836)
					l64:
						{
							t113 := int32(load32(m.memory[uint32(v1):]))
							t114 := v1
							v17 = t113
							t115 := int32(load32(m.memory[uint32(v9):]))
							v18 = v17 + (t115 ^ i32(-1))
							v19 = v18 + v13&i32(1)
							store32(m.memory[uint32(t114):], uint32(v19))
							v13 = v1 + i32(4)
							t116 := int32(load32(m.memory[uint32(v13):]))
							t117 := v13
							v20 = t116
							t118 := int32(load32(m.memory[uint32(v9+i32(4)):]))
							v13 = v20 + (t118 ^ i32(-1))
							t119 := v13
							var p120 int32
							if uint32(v18) < uint32(v17) {
								p120 = 1
							}
							var p121 int32
							if uint32(v19) < uint32(v18) {
								p121 = 1
							}
							v18 = t119 + (p120 | p121)
							store32(m.memory[uint32(t117):], uint32(v18))
							var p122 int32
							if uint32(v13) < uint32(v20) {
								p122 = 1
							}
							var p123 int32
							if uint32(v18) < uint32(v13) {
								p123 = 1
							}
							v13 = p122 | p123
							v9 = v9 + i32(8)
							v1 = v1 + i32(8)
							t124 := v35
							v14 = v14 + i32(2)
							if t124 != v14 {
								goto l64
							}
						}
						if v34 == 0 {
							goto l65
						}
					l63:
						t125 := v3 + i32(8)
						v1 = v14 << 2
						v9 = t125 + v1
						t126 := int32(load32(m.memory[uint32(v9):]))
						t127 := v9
						v9 = t126
						t128 := int32(load32(m.memory[uint32(v24+v1):]))
						v1 = v9 + (t128 ^ i32(-1))
						v18 = v1 + v13
						store32(m.memory[uint32(t127):], uint32(v18))
						if uint32(v1) < uint32(v9) {
							goto l62
						}
						if uint32(v18) < uint32(v1) {
							goto l62
						}
						goto l66
					}
				l65:
					if v13 == 0 {
						goto l66
					}
				l62:
					store32(m.memory[int64(uint32(v3))+168:], uint32(v27))
					v33 = v33 | i32(4)
				l61:
					p129 := v27
					if uint32(v29) > uint32(v27) {
						p129 = v29
					}
					v34 = p129
					if uint32(v34) >= uint32(i32(41)) {
						m.fn124(i32(0), v34, i32(40), i32(1099776))
						panic("unreachable")
					}
					v1 = v34 << 2
				l69:
					{
						if v1 == 0 {
							goto l68
						}
						v9 = v11 + v1
						v1 = v1 + i32(-4)
						t130 := int32(load32(m.memory[uint32(v1+(v3+i32(8))):]))
						v18 = t130
						t131 := int32(load32(m.memory[uint32(v9):]))
						t132 := v18
						v9 = t131
						if t132 == v9 {
							goto l69
						}
					}
					if uint32(v18) >= uint32(v9) {
						goto l68
					}
					v34 = v27
					goto l70
				l68:
					if v34 == 0 {
						goto l71
					}
					v13 = i32(1)
					v14 = i32(0)
					{
						if v34 == i32(1) {
							goto l72
						}
						v27 = v34 & i32(1)
						v35 = v34 & i32(62)
						v14 = i32(0)
						v13 = i32(1)
						v1 = v3 + i32(8)
						v9 = v3 + i32(672)
					l73:
						{
							t133 := int32(load32(m.memory[uint32(v1):]))
							t134 := v1
							v17 = t133
							t135 := int32(load32(m.memory[uint32(v9):]))
							v18 = v17 + (t135 ^ i32(-1))
							v19 = v18 + v13&i32(1)
							store32(m.memory[uint32(t134):], uint32(v19))
							v13 = v1 + i32(4)
							t136 := int32(load32(m.memory[uint32(v13):]))
							t137 := v13
							v20 = t136
							t138 := int32(load32(m.memory[uint32(v9+i32(4)):]))
							v13 = v20 + (t138 ^ i32(-1))
							t139 := v13
							var p140 int32
							if uint32(v18) < uint32(v17) {
								p140 = 1
							}
							var p141 int32
							if uint32(v19) < uint32(v18) {
								p141 = 1
							}
							v18 = t139 + (p140 | p141)
							store32(m.memory[uint32(t137):], uint32(v18))
							var p142 int32
							if uint32(v13) < uint32(v20) {
								p142 = 1
							}
							var p143 int32
							if uint32(v18) < uint32(v13) {
								p143 = 1
							}
							v13 = p142 | p143
							v9 = v9 + i32(8)
							v1 = v1 + i32(8)
							t144 := v35
							v14 = v14 + i32(2)
							if t144 != v14 {
								goto l73
							}
						}
						if v27 == 0 {
							goto l74
						}
					l72:
						t145 := v3 + i32(8)
						v1 = v14 << 2
						v9 = t145 + v1
						t146 := int32(load32(m.memory[uint32(v9):]))
						t147 := v9
						v9 = t146
						t148 := int32(load32(m.memory[uint32(v23+v1):]))
						v1 = v9 + (t148 ^ i32(-1))
						v18 = v1 + v13
						store32(m.memory[uint32(t147):], uint32(v18))
						if uint32(v1) < uint32(v9) {
							goto l71
						}
						if uint32(v18) < uint32(v1) {
							goto l71
						}
						goto l75
					}
				l74:
					if v13 == 0 {
						goto l75
					}
				l71:
					store32(m.memory[int64(uint32(v3))+168:], uint32(v34))
					v33 = v33 + i32(2)
				l70:
					p149 := v34
					if uint32(v21) > uint32(v34) {
						p149 = v21
					}
					v27 = p149
					if uint32(v27) >= uint32(i32(41)) {
						m.fn124(i32(0), v27, i32(40), i32(1099776))
						panic("unreachable")
					}
					v1 = v27 << 2
				l78:
					{
						if v1 == 0 {
							goto l77
						}
						v1 = v1 + i32(-4)
						t150 := int32(load32(m.memory[uint32(v1+(v3+i32(8))):]))
						v9 = t150
						t151 := int32(load32(m.memory[uint32(v1+(v3+i32(508))):]))
						t152 := v9
						v18 = t151
						if t152 == v18 {
							goto l78
						}
					}
					if uint32(v9) >= uint32(v18) {
						goto l77
					}
					v27 = v34
					goto l79
				l77:
					if v27 == 0 {
						goto l80
					}
					v13 = i32(1)
					v14 = i32(0)
					{
						if v27 == i32(1) {
							goto l81
						}
						v34 = v27 & i32(1)
						v35 = v27 & i32(62)
						v14 = i32(0)
						v13 = i32(1)
						v1 = v3 + i32(8)
						v9 = v3 + i32(508)
					l82:
						{
							t153 := int32(load32(m.memory[uint32(v1):]))
							t154 := v1
							v17 = t153
							t155 := int32(load32(m.memory[uint32(v9):]))
							v18 = v17 + (t155 ^ i32(-1))
							v19 = v18 + v13&i32(1)
							store32(m.memory[uint32(t154):], uint32(v19))
							v13 = v1 + i32(4)
							t156 := int32(load32(m.memory[uint32(v13):]))
							t157 := v13
							v20 = t156
							t158 := int32(load32(m.memory[uint32(v9+i32(4)):]))
							v13 = v20 + (t158 ^ i32(-1))
							t159 := v13
							var p160 int32
							if uint32(v18) < uint32(v17) {
								p160 = 1
							}
							var p161 int32
							if uint32(v19) < uint32(v18) {
								p161 = 1
							}
							v18 = t159 + (p160 | p161)
							store32(m.memory[uint32(t157):], uint32(v18))
							var p162 int32
							if uint32(v13) < uint32(v20) {
								p162 = 1
							}
							var p163 int32
							if uint32(v18) < uint32(v13) {
								p163 = 1
							}
							v13 = p162 | p163
							v9 = v9 + i32(8)
							v1 = v1 + i32(8)
							t164 := v35
							v14 = v14 + i32(2)
							if t164 != v14 {
								goto l82
							}
						}
						if v34 == 0 {
							goto l83
						}
					l81:
						t165 := v3 + i32(8)
						v1 = v14 << 2
						v9 = t165 + v1
						t166 := int32(load32(m.memory[uint32(v9):]))
						t167 := v9
						v9 = t166
						t168 := int32(load32(m.memory[uint32(v3+i32(508)+v1):]))
						v1 = v9 + (t168 ^ i32(-1))
						v18 = v1 + v13
						store32(m.memory[uint32(t167):], uint32(v18))
						if uint32(v1) < uint32(v9) {
							goto l80
						}
						if uint32(v18) < uint32(v1) {
							goto l80
						}
						goto l84
					}
				l83:
					if v13 == 0 {
						goto l84
					}
				l80:
					store32(m.memory[int64(uint32(v3))+168:], uint32(v27))
					v33 = v33 + i32(1)
				l79:
					if v32 == i32(17) {
						m.fn36(i32(17), i32(17), i32(1121568))
						panic("unreachable")
					}
					m.memory[uint32(v2+v32)] = byte(v33 + i32(48))
					t169 := int32(load32(m.memory[int64(uint32(v3))+336:]))
					v33 = t169
					p170 := v27
					if uint32(v33) > uint32(v27) {
						p170 = v33
					}
					v1 = p170
					if uint32(v1) >= uint32(i32(41)) {
						m.fn124(i32(0), v1, i32(40), i32(1099776))
						panic("unreachable")
					}
					v31 = v32 + i32(1)
					v1 = v1 << 2
					{
					l89:
						{
							if v1 != 0 {
								goto l87
							}
							v36 = i32(0)
							goto l88
						l87:
							v1 = v1 + i32(-4)
							t171 := int32(load32(m.memory[uint32(v1+(v3+i32(8))):]))
							v9 = t171
							t172 := int32(load32(m.memory[uint32(v1+(v3+i32(176))):]))
							t173 := v9
							v18 = t172
							if t173 == v18 {
								goto l89
							}
						}
						var p174 int32
						if uint32(v9) > uint32(v18) {
							p174 = 1
						}
						var p175 int32
						if uint32(v9) < uint32(v18) {
							p175 = 1
						}
						v36 = p174 - p175
					}
				l88:
					memory_copy(m.memory, uint32(v3+i32(1164)), uint32(v3+i32(8)), uint32(i32(164)))
					t176 := int32(load32(m.memory[int64(uint32(v3))+1324:]))
					t177 := v22
					v1 = t176
					p178 := v1
					if uint32(v22) > uint32(v1) {
						p178 = t177
					}
					v34 = p178
					if uint32(v34) > uint32(i32(40)) {
						m.fn124(i32(0), v34, i32(40), i32(1099776))
						panic("unreachable")
					}
					if v34 != 0 {
						goto l91
					}
					v34 = i32(0)
					goto l92
				l91:
					v13 = i32(0)
					v14 = i32(0)
					{
						if v34 == i32(1) {
							goto l93
						}
						v37 = v34 & i32(1)
						v35 = v34 & i32(62)
						v13 = i32(0)
						v1 = v3 + i32(1164)
						v9 = v3 + i32(344)
						v14 = i32(0)
					l94:
						{
							t179 := int32(load32(m.memory[uint32(v9):]))
							t180 := v1
							v17 = t179
							t181 := int32(load32(m.memory[uint32(v1):]))
							v18 = v17 + t181
							v19 = v18 + v13&i32(1)
							store32(m.memory[uint32(t180):], uint32(v19))
							v13 = v1 + i32(4)
							t182 := int32(load32(m.memory[uint32(v9+i32(4)):]))
							t183 := v13
							v20 = t182
							t184 := int32(load32(m.memory[uint32(v13):]))
							v13 = v20 + t184
							t185 := v13
							var p186 int32
							if uint32(v18) < uint32(v17) {
								p186 = 1
							}
							var p187 int32
							if uint32(v19) < uint32(v18) {
								p187 = 1
							}
							v18 = t185 + (p186 | p187)
							store32(m.memory[uint32(t183):], uint32(v18))
							var p188 int32
							if uint32(v13) < uint32(v20) {
								p188 = 1
							}
							var p189 int32
							if uint32(v18) < uint32(v13) {
								p189 = 1
							}
							v13 = p188 | p189
							v9 = v9 + i32(8)
							v1 = v1 + i32(8)
							t190 := v35
							v14 = v14 + i32(2)
							if t190 != v14 {
								goto l94
							}
						}
						if v37 == 0 {
							goto l95
						}
					l93:
						t191 := v3 + i32(1164)
						v1 = v14 << 2
						v9 = t191 + v1
						t192 := int32(load32(m.memory[uint32(v3+i32(344)+v1):]))
						t193 := v9
						v18 = t192
						t194 := int32(load32(m.memory[uint32(v9):]))
						v1 = v18 + t194
						v9 = v1 + v13
						store32(m.memory[uint32(t193):], uint32(v9))
						if uint32(v1) < uint32(v18) {
							goto l96
						}
						if uint32(v9) < uint32(v1) {
							goto l96
						}
						goto l92
					}
				l95:
					if v13 == 0 {
						goto l92
					}
				l96:
					if v34 == i32(40) {
						m.fn36(i32(40), i32(40), i32(1099776))
						panic("unreachable")
					}
					store32(m.memory[uint32(v3+i32(1164)+v34<<2):], uint32(i32(1)))
					v34 = v34 + i32(1)
				l92:
					store32(m.memory[int64(uint32(v3))+1324:], uint32(v34))
					p195 := v21
					if uint32(v34) > uint32(v21) {
						p195 = v34
					}
					v1 = p195
					if uint32(v1) >= uint32(i32(41)) {
						m.fn124(i32(0), v1, i32(40), i32(1099776))
						panic("unreachable")
					}
					v1 = v1 << 2
					{
					l101:
						{
							if v1 != 0 {
								goto l99
							}
							v1 = i32(0)
							goto l100
						l99:
							v9 = v28 + v1
							v18 = v15 + v1
							v1 = v1 + i32(-4)
							t196 := int32(load32(m.memory[uint32(v18):]))
							v18 = t196
							t197 := int32(load32(m.memory[uint32(v9):]))
							t198 := v18
							v9 = t197
							if t198 == v9 {
								goto l101
							}
						}
						var p199 int32
						if uint32(v18) > uint32(v9) {
							p199 = 1
						}
						var p200 int32
						if uint32(v18) < uint32(v9) {
							p200 = 1
						}
						v1 = p199 - p200
					}
				l100:
					if v36 < v8 {
						goto l102
					}
					if v1 < v8 {
						goto l103
					}
					v14 = i32(0)
					v13 = i32(0)
					if v27 == 0 {
						goto l104
					}
					v17 = v27 << 2
					v1 = v17 + i32(-4)
					v9 = int32(uint32(v1)>>2) + i32(1)
					v13 = v9 & i32(3)
					if uint32(v1) >= uint32(i32(12)) {
						goto l105
					}
					v4 = i64(0)
					v1 = v3 + i32(8)
					goto l106
				l105:
					v9 = v9 & i32(0x7ffffffc)
					v4 = i64(0)
					v1 = v3 + i32(8)
				l107:
					{
						t201 := int64(load32(m.memory[uint32(v1):]))
						t202 := v1
						v4 = t201*i64(10) + v4
						store32(m.memory[uint32(t202):], uint32(v4))
						v18 = v1 + i32(4)
						t203 := int64(load32(m.memory[uint32(v18):]))
						t204 := v18
						v4 = t203*i64(10) + int64(uint64(v4)>>32)
						store32(m.memory[uint32(t204):], uint32(v4))
						v18 = v1 + i32(8)
						t205 := int64(load32(m.memory[uint32(v18):]))
						t206 := v18
						v4 = t205*i64(10) + int64(uint64(v4)>>32)
						store32(m.memory[uint32(t206):], uint32(v4))
						v18 = v1 + i32(12)
						t207 := int64(load32(m.memory[uint32(v18):]))
						t208 := v18
						v4 = t207*i64(10) + int64(uint64(v4)>>32)
						store32(m.memory[uint32(t208):], uint32(v4))
						v4 = int64(uint64(v4) >> 32)
						v1 = v1 + i32(16)
						v9 = v9 + i32(-4)
						if v9 != 0 {
							goto l107
						}
					}
					if v13 == 0 {
						goto l108
					}
				l106:
					v9 = v13 << 2
				l109:
					{
						t209 := int64(load32(m.memory[uint32(v1):]))
						t210 := v1
						v4 = t209*i64(10) + v4
						store32(m.memory[uint32(t210):], uint32(v4))
						v1 = v1 + i32(4)
						v4 = int64(uint64(v4) >> 32)
						v9 = v9 + i32(-4)
						if v9 != 0 {
							goto l109
						}
					}
				l108:
					if !(v4 == 0) {
						goto l110
					}
					v13 = v27
					goto l104
				l110:
					if v27 == i32(40) {
						m.fn36(i32(40), i32(40), i32(1099776))
						panic("unreachable")
					}
					store32(m.memory[uint32(v3+i32(8)+v17):], uint32(int32(v4)))
					v13 = v27 + i32(1)
				l104:
					store32(m.memory[int64(uint32(v3))+168:], uint32(v13))
					if v33 == 0 {
						goto l112
					}
					v17 = v33 << 2
					v1 = v17 + i32(-4)
					v9 = int32(uint32(v1)>>2) + i32(1)
					v14 = v9 & i32(3)
					if uint32(v1) >= uint32(i32(12)) {
						goto l113
					}
					v4 = i64(0)
					v1 = v3 + i32(176)
					goto l114
				l113:
					v9 = v9 & i32(0x7ffffffc)
					v4 = i64(0)
					v1 = v3 + i32(176)
				l115:
					{
						t211 := int64(load32(m.memory[uint32(v1):]))
						t212 := v1
						v4 = t211*i64(10) + v4
						store32(m.memory[uint32(t212):], uint32(v4))
						v18 = v1 + i32(4)
						t213 := int64(load32(m.memory[uint32(v18):]))
						t214 := v18
						v4 = t213*i64(10) + int64(uint64(v4)>>32)
						store32(m.memory[uint32(t214):], uint32(v4))
						v18 = v1 + i32(8)
						t215 := int64(load32(m.memory[uint32(v18):]))
						t216 := v18
						v4 = t215*i64(10) + int64(uint64(v4)>>32)
						store32(m.memory[uint32(t216):], uint32(v4))
						v18 = v1 + i32(12)
						t217 := int64(load32(m.memory[uint32(v18):]))
						t218 := v18
						v4 = t217*i64(10) + int64(uint64(v4)>>32)
						store32(m.memory[uint32(t218):], uint32(v4))
						v4 = int64(uint64(v4) >> 32)
						v1 = v1 + i32(16)
						v9 = v9 + i32(-4)
						if v9 != 0 {
							goto l115
						}
					}
					if v14 == 0 {
						goto l116
					}
				l114:
					v9 = v14 << 2
				l117:
					{
						t219 := int64(load32(m.memory[uint32(v1):]))
						t220 := v1
						v4 = t219*i64(10) + v4
						store32(m.memory[uint32(t220):], uint32(v4))
						v1 = v1 + i32(4)
						v4 = int64(uint64(v4) >> 32)
						v9 = v9 + i32(-4)
						if v9 != 0 {
							goto l117
						}
					}
				l116:
					if !(v4 == 0) {
						goto l118
					}
					v14 = v33
					goto l112
				l118:
					if v33 == i32(40) {
						m.fn36(i32(40), i32(40), i32(1099776))
						panic("unreachable")
					}
					store32(m.memory[uint32(v3+i32(176)+v17):], uint32(int32(v4)))
					v14 = v33 + i32(1)
				l112:
					store32(m.memory[int64(uint32(v3))+336:], uint32(v14))
					if v22 != 0 {
						goto l120
					}
					v22 = i32(0)
					goto l121
				l120:
					v17 = v22 << 2
					v1 = v17 + i32(-4)
					v9 = int32(uint32(v1)>>2) + i32(1)
					v14 = v9 & i32(3)
					if uint32(v1) >= uint32(i32(12)) {
						goto l122
					}
					v4 = i64(0)
					v1 = v3 + i32(344)
					goto l123
				l122:
					v9 = v9 & i32(0x7ffffffc)
					v4 = i64(0)
					v1 = v3 + i32(344)
				l124:
					{
						t221 := int64(load32(m.memory[uint32(v1):]))
						t222 := v1
						v4 = t221*i64(10) + v4
						store32(m.memory[uint32(t222):], uint32(v4))
						v18 = v1 + i32(4)
						t223 := int64(load32(m.memory[uint32(v18):]))
						t224 := v18
						v4 = t223*i64(10) + int64(uint64(v4)>>32)
						store32(m.memory[uint32(t224):], uint32(v4))
						v18 = v1 + i32(8)
						t225 := int64(load32(m.memory[uint32(v18):]))
						t226 := v18
						v4 = t225*i64(10) + int64(uint64(v4)>>32)
						store32(m.memory[uint32(t226):], uint32(v4))
						v18 = v1 + i32(12)
						t227 := int64(load32(m.memory[uint32(v18):]))
						t228 := v18
						v4 = t227*i64(10) + int64(uint64(v4)>>32)
						store32(m.memory[uint32(t228):], uint32(v4))
						v4 = int64(uint64(v4) >> 32)
						v1 = v1 + i32(16)
						v9 = v9 + i32(-4)
						if v9 != 0 {
							goto l124
						}
					}
					if v14 == 0 {
						goto l125
					}
				l123:
					v9 = v14 << 2
				l126:
					{
						t229 := int64(load32(m.memory[uint32(v1):]))
						t230 := v1
						v4 = t229*i64(10) + v4
						store32(m.memory[uint32(t230):], uint32(v4))
						v1 = v1 + i32(4)
						v4 = int64(uint64(v4) >> 32)
						v9 = v9 + i32(-4)
						if v9 != 0 {
							goto l126
						}
					}
				l125:
					if v4 == 0 {
						goto l121
					}
					if v22 == i32(40) {
						m.fn36(i32(40), i32(40), i32(1099776))
						panic("unreachable")
					}
					store32(m.memory[uint32(v3+i32(344)+v17):], uint32(int32(v4)))
					v22 = v22 + i32(1)
				l121:
					store32(m.memory[int64(uint32(v3))+504:], uint32(v22))
					p231 := v13
					if uint32(v26) > uint32(v13) {
						p231 = v26
					}
					v27 = p231
					if uint32(v27) < uint32(i32(41)) {
						goto l128
					}
				}
			}
		l49:
			m.fn124(i32(0), v27, i32(40), i32(1099776))
			panic("unreachable")
		l102:
			if v1 >= v8 {
				goto l129
			}
			_ = m.fn859(v3+i32(8), i32(1))
			t233 := int32(load32(m.memory[int64(uint32(v3))+168:]))
			t234 := v21
			v1 = t233
			p235 := v1
			if uint32(v21) > uint32(v1) {
				p235 = t234
			}
			v1 = p235
			if uint32(v1) >= uint32(i32(41)) {
				m.fn124(i32(0), v1, i32(40), i32(1099776))
				panic("unreachable")
			}
			v1 = v1 << 2
			v13 = v3 + i32(8) + i32(-4)
			v14 = v3 + i32(508) + i32(-4)
		l131:
			{
				if v1 == 0 {
					goto l103
				}
				v9 = v14 + v1
				v18 = v13 + v1
				v1 = v1 + i32(-4)
				t236 := int32(load32(m.memory[uint32(v18):]))
				v18 = t236
				t237 := int32(load32(m.memory[uint32(v9):]))
				t238 := v18
				v9 = t237
				if t238 == v9 {
					goto l131
				}
			}
			if uint32(v18) < uint32(v9) {
				goto l129
			}
		}
	l103:
		v13 = v2 + v31
		v1 = v31
		{
		l133:
			{
				v9 = v1
				if v9 == 0 {
					m.memory[uint32(v2)] = byte(i32(49))
					if v32 == 0 {
						goto l134
					}
					memory_fill(m.memory, uint32(v2+i32(1)), i32(48), uint32(v32))
				l134:
					if uint32(v32) > uint32(i32(15)) {
						m.fn36(v31, i32(17), i32(1121584))
						panic("unreachable")
					}
					m.memory[uint32(v13)] = byte(i32(48))
					v10 = v10 + i32(1)
					v31 = v32 + i32(2)
					goto l136
				}
				v1 = v9 + i32(-1)
				v18 = v1 + v2
				t239 := int32(m.memory[uint32(v18)])
				if t239 == i32(57) {
					goto l133
				}
			}
			t240 := int32(m.memory[uint32(v18)])
			m.memory[uint32(v18)] = byte(t240 + i32(1))
			v1 = v31 - v9
			if v1 == 0 {
				goto l129
			}
			memory_fill(m.memory, uint32(v2+v9), i32(48), uint32(v1))
			goto l129
		}
	}
l129:
	if uint32(v32) <= uint32(i32(16)) {
		goto l136
	}
	m.fn124(i32(0), v31, i32(17), i32(1121600))
	panic("unreachable")
l136:
	store16(m.memory[int64(uint32(v0))+8:], uint16(v10))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v31))
	store32(m.memory[uint32(v0):], uint32(v2))
	m.g0 = v3 + i32(1328)
	return
l84:
	m.fn7(i32(1099747), i32(26), i32(1099776))
	panic("unreachable")
l75:
	m.fn7(i32(1099747), i32(26), i32(1099776))
	panic("unreachable")
l66:
	m.fn7(i32(1099747), i32(26), i32(1099776))
	panic("unreachable")
l57:
	m.fn7(i32(1099747), i32(26), i32(1099776))
	panic("unreachable")
}
func (m *Module) fn875(v0, v1 int32) int32 {
	var v2, v3 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	t1 := int32(m.memory[uint32(v0)])
	v3 = t1
	v0 = i32(3)
l0:
	{
		t2 := int32(m.memory[uint32(v3&i32(15)+i32(1122456))])
		m.memory[uint32(v2+i32(14)+v0+i32(-2))] = byte(t2)
		v0 = v0 + i32(-1)
		v3 = int32(uint32(v3)>>4) & i32(15)
		if v3 != 0 {
			goto l0
		}
	}
	t3 := m.fn683(v1, i32(1), i32(1122454), i32(2), v2+i32(14)+v0+i32(-1), i32(3)-v0)
	v0 = t3
	m.g0 = v2 + i32(16)
	return v0
}
func (m *Module) fn876(v0, v1 int32) int32 {
	var v2, v3 int32
	var v4 int64
	var v5, v6, v7, v8, v9 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	{
		{
			t1 := int32(load16(m.memory[int64(uint32(v0))+12:]))
			v3 = t1
			if v3 != 0 {
				goto l0
			}
			t2 := int32(load32(m.memory[uint32(v0):]))
			t3 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t4 := m.fn877(t2, t3, v1)
			v1 = t4
			goto l1
		}
	l0:
		t5 := int64(load64(m.memory[int64(uint32(v1))+8:]))
		store64(m.memory[int64(uint32(v2))+8:], uint64(t5))
		t6 := int64(load64(m.memory[uint32(v1):]))
		store64(m.memory[uint32(v2):], uint64(t6))
		{
			{
				t7 := int64(load64(m.memory[int64(uint32(v0))+8:]))
				v4 = t7
				v5 = int32(v4)
				if v5&i32(0x1000000) != 0 {
					goto l2
				}
				t8 := int32(load32(m.memory[int64(uint32(v2))+4:]))
				v6 = t8
				goto l3
			}
		l2:
			t9 := int32(load32(m.memory[uint32(v0):]))
			t10 := int32(load32(m.memory[uint32(v2):]))
			t11 := int32(load32(m.memory[int64(uint32(v2))+4:]))
			v1 = t11
			t12 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t13 := int32(load32(m.memory[int64(uint32(t12))+12:]))
			t14 := m.t0[uint(t13)].(func(int32, int32, int32) int32)(t9, t10, v1)
			if t14 != 0 {
				goto l4
			}
			t15 := v0
			v5 = v5&i32(-0x60200000) | i32(0x20000030)
			store32(m.memory[int64(uint32(t15))+8:], uint32(v5))
			store64(m.memory[uint32(v2):], uint64(i64(1)))
			v6 = i32(0)
			v1 = v3 - v1&i32(0xffff)
			p16 := v1
			if uint32(v1) > uint32(v3) {
				p16 = i32(0)
			}
			v3 = p16
		}
	l3:
		{
			t17 := int32(load32(m.memory[int64(uint32(v2))+12:]))
			v7 = t17
			if v7 == 0 {
				goto l5
			}
			t18 := int32(load32(m.memory[int64(uint32(v2))+8:]))
			v1 = t18
		l11:
			{
				{
					t19 := int32(load16(m.memory[uint32(v1):]))
					switch t19 {
					default:
						t20 := int32(load32(m.memory[uint32(v1+i32(4)):]))
						v8 = t20
						goto l9
					case 1:
						t21 := int32(load16(m.memory[uint32(v1+i32(2)):]))
						v8 = t21
						if v8 != 0 {
							goto l10
						}
						v8 = i32(1)
						goto l9
					case 2:
						t22 := int32(load32(m.memory[uint32(v1+i32(8)):]))
						v8 = t22
						goto l9
					}
				}
			l10:
				v8 = int32(uint32((v8+i32(0x5fff6))&(v8+i32(524188))^(v8+i32(916504))&(v8+i32(514288)))>>17) + i32(1)
			l9:
				v8 = v6 + v8
				p23 := v8
				if uint32(v8) < uint32(v6) {
					p23 = i32(-1)
				}
				v6 = p23
				v1 = v1 + i32(12)
				v7 = v7 + i32(-1)
				if v7 != 0 {
					goto l11
				}
			}
		}
	l5:
		{
			if uint32(v6) < uint32(v3&i32(0xffff)) {
				v9 = v3 - v6
				v1 = i32(0)
				v3 = i32(0)
				switch int32(uint32(v5)>>29) & i32(3) {
				default:
					goto l13
				case 1, 3:
					v3 = v9
					goto l13
				case 2:
					v3 = int32(uint32(v9&i32(65534)) >> 1)
				}
			l13:
				v8 = v5 & i32(0x1fffff)
				t27 := int32(load32(m.memory[int64(uint32(v0))+4:]))
				v6 = t27
				t28 := int32(load32(m.memory[uint32(v0):]))
				v7 = t28
			l17:
				{
					if uint32(v1&i32(0xffff)) >= uint32(v3&i32(0xffff)) {
						t31 := m.fn877(v7, v6, v2)
						if t31 != 0 {
							goto l4
						}
						v5 = (v9 - v3) & i32(0xffff)
						v3 = i32(0)
					l19:
						{
							if uint32(v3&i32(0xffff)) < uint32(v5) {
								goto l18
							}
							v1 = i32(0)
							store64(m.memory[int64(uint32(v0))+8:], uint64(v4))
							goto l1
						l18:
							v1 = i32(1)
							v3 = v3 + i32(1)
							t32 := int32(load32(m.memory[int64(uint32(v6))+16:]))
							t33 := m.t0[uint(t32)].(func(int32, int32) int32)(v7, v8)
							if t33 == 0 {
								goto l19
							}
						}
						store64(m.memory[int64(uint32(v0))+8:], uint64(v4))
						goto l1
					}
					v1 = v1 + i32(1)
					t29 := int32(load32(m.memory[int64(uint32(v6))+16:]))
					t30 := m.t0[uint(t29)].(func(int32, int32) int32)(v7, v8)
					if t30 == 0 {
						goto l17
					}
					goto l4
				}
			}
			t24 := int32(load32(m.memory[uint32(v0):]))
			t25 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t26 := m.fn877(t24, t25, v2)
			v1 = t26
			store64(m.memory[int64(uint32(v0))+8:], uint64(v4))
			goto l1
		}
	l4:
		v1 = i32(1)
	}
l1:
	m.g0 = v2 + i32(16)
	return v1
}
func (m *Module) fn877(v0, v1, v2 int32) int32 {
	var v3, v4, v5, v6, v7, v8, v9 int32
	t0 := m.g0
	v3 = t0 - i32(16)
	m.g0 = v3
	{
		t1 := int32(load32(m.memory[int64(uint32(v2))+4:]))
		v4 = t1
		if v4 == 0 {
			goto l0
		}
		t2 := int32(load32(m.memory[uint32(v2):]))
		t3 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		t4 := m.t0[uint(t3)].(func(int32, int32, int32) int32)(v0, t2, v4)
		if t4 != 0 {
			goto l1
		}
	}
l0:
	{
		t5 := int32(load32(m.memory[int64(uint32(v2))+12:]))
		v5 = t5
		if v5 != 0 {
			t6 := int32(load32(m.memory[int64(uint32(v2))+8:]))
			v4 = t6
			v6 = v4 + v5*i32(12)
		l13:
			{
				{
					{
						{
							{
								t7 := int32(load16(m.memory[uint32(v4):]))
								switch t7 {
								default:
									t8 := int32(load32(m.memory[int64(uint32(v4))+4:]))
									v2 = t8
									if uint32(v2) < uint32(i32(65)) {
										goto l7
									}
									t9 := int32(load32(m.memory[uint32(v1+i32(12)):]))
									v5 = t9
								l8:
									{
										t10 := m.t0[uint(v5)].(func(int32, int32, int32) int32)(v0, i32(1099912), i32(64))
										if t10 != 0 {
											goto l1
										}
										v2 = v2 + i32(-64)
										if uint32(v2) > uint32(i32(64)) {
											goto l8
										}
										goto l9
									}
								case 1:
									t11 := int32(load16(m.memory[int64(uint32(v4))+2:]))
									v2 = t11
									m.memory[int64(uint32(v3))+12] = byte(i32(0))
									store32(m.memory[int64(uint32(v3))+8:], uint32(i32(0)))
									if v2 != 0 {
										goto l10
									}
									m.memory[int64(uint32(v3))+8] = byte(v2 | i32(48))
									v5 = i32(1)
									goto l11
								case 2:
									t12 := int32(load32(m.memory[int64(uint32(v4))+4:]))
									t13 := int32(load32(m.memory[int64(uint32(v4))+8:]))
									t14 := int32(load32(m.memory[uint32(v1+i32(12)):]))
									t15 := m.t0[uint(t14)].(func(int32, int32, int32) int32)(v0, t12, t13)
									if t15 == 0 {
										goto l12
									}
									goto l1
								}
							}
						l7:
							if v2 == 0 {
								goto l12
							}
							t16 := int32(load32(m.memory[uint32(v1+i32(12)):]))
							v5 = t16
						}
					l9:
						t17 := m.t0[uint(v5)].(func(int32, int32, int32) int32)(v0, i32(1099912), v2)
						if t17 != 0 {
							goto l1
						}
						goto l12
					}
				l10:
					t18 := v3 + i32(8)
					v7 = int32(uint32((v2+i32(0x5fff6))&(v2+i32(524188))^(v2+i32(916504))&(v2+i32(514288))) >> 17)
					v8 = t18 + v7
					t19 := int32(uint32(v2) / uint32(i32(10)))
					t20 := v8
					t21 := v2
					v9 = t19
					m.memory[uint32(t20)] = byte(t21 - v9*i32(10) | i32(48))
					v5 = v7 + i32(1)
					if v7 == 0 {
						goto l11
					}
					t22 := int32(uint32(v9) % uint32(i32(10)))
					m.memory[uint32(v8+i32(-1))] = byte(t22 | i32(48))
					if v5 == i32(2) {
						goto l11
					}
					t23 := int32(uint32(v2) / uint32(i32(100)))
					t24 := int32(uint32(t23) % uint32(i32(10)))
					m.memory[uint32(v8+i32(-2))] = byte(t24 | i32(48))
					if v5 == i32(3) {
						goto l11
					}
					t25 := int32(uint32(v2) / uint32(i32(1000)))
					t26 := int32(uint32(t25) % uint32(i32(10)))
					m.memory[uint32(v8+i32(-3))] = byte(t26 | i32(48))
					if v5 == i32(4) {
						goto l11
					}
					t27 := int32(uint32(v2) / uint32(i32(10000)))
					m.memory[uint32(v8+i32(-4))] = byte(t27 | i32(48))
				}
			l11:
				t28 := int32(load32(m.memory[uint32(v1+i32(12)):]))
				t29 := m.t0[uint(t28)].(func(int32, int32, int32) int32)(v0, v3+i32(8), v5)
				if t29 != 0 {
					goto l1
				}
			}
		l12:
			v4 = v4 + i32(12)
			if v4 != v6 {
				goto l13
			}
			v2 = i32(0)
			goto l3
		}
		v2 = i32(0)
		goto l3
	}
l1:
	v2 = i32(1)
l3:
	m.g0 = v3 + i32(16)
	return v2
}
func (m *Module) fn878(v0 int32, v1, v2 int64) {
	var v3, v4, v5 int32
	var v6, v7, v8, v9, v10 int64
	var v11 int32
	t0 := m.g0
	v3 = t0 - i32(32)
	m.g0 = v3
	store32(m.memory[int64(uint32(v0))+8:], uint32(i32(2047)))
	store64(m.memory[uint32(v0):], uint64(i64(0)))
	v4 = i32(0)
	{
		if v1 < i64(-342) {
			goto l0
		}
		if v2 == 0 {
			goto l0
		}
		if v1 > i64(308) {
			goto l1
		}
		t1 := v3 + i32(16)
		v5 = int32(v1)
		v4 = v5 << 4
		t2 := int64(load64(m.memory[uint32(v4+i32(1114568)):]))
		t3 := v2
		v6 = int64(bits.LeadingZeros64(uint64(v2)))
		v7 = i64_shl(t3, v6)
		m.fn1911(t1, t2, i64(0), v7, i64(0))
		t4 := int64(load64(m.memory[int64(uint32(v3))+16:]))
		v8 = t4
		{
			t5 := int64(load64(m.memory[int64(uint32(v3))+24:]))
			v2 = t5
			if v2&i64(511) != i64(511) {
				goto l2
			}
			t6 := int64(load64(m.memory[uint32(v4+i32(1109096)+i32(5480)):]))
			m.fn1911(v3, t6, i64(0), v7, i64(0))
			t7 := int64(load64(m.memory[int64(uint32(v3))+8:]))
			v7 = t7
			v8 = v7 + v8
			var p8 int32
			if uint64(v8) < uint64(v7) {
				p8 = 1
			}
			v2 = int64(uint32(p8)) + v2
		}
	l2:
		if uint64(v1+i64(27)) < uint64(i64(83)) {
			goto l3
		}
		v4 = i32(-1)
		if v8 == i64(-1) {
			goto l0
		}
	l3:
		t9 := v2
		v9 = int64(uint64(v2) >> 63)
		v10 = v9 + i64(9)
		v7 = i64_shr_u(t9, v10)
		{
			{
				v5 = v5*i32(217706)>>16 - int32(v6) + int32(v9) + i32(63)
				if v5 < i32(-1022) {
					goto l4
				}
				p10 := v7
				if i64_shl(v7, v10) == v2 {
					p10 = v7 & i64(0xfffffffffffffc)
				}
				p11 := v7
				if v7&i64(3) == i64(1) {
					p11 = p10
				}
				p12 := v7
				if uint64(v8) < uint64(i64(2)) {
					p12 = p11
				}
				p13 := v7
				if uint64(v1+i64(4)) < uint64(i64(28)) {
					p13 = p12
				}
				v1 = p13
				v1 = v1&i64(1) + v1
				var p14 int32
				if uint64(v1) > uint64(i64(0x3fffffffffffff)) {
					p14 = 1
				}
				v11 = p14
				p15 := i32(1023)
				if v11 != 0 {
					p15 = i32(1024)
				}
				v4 = p15 + v5
				if uint32(v4) > uint32(i32(2046)) {
					goto l1
				}
				p16 := int64(uint64(v1)>>1) & i64(0x7fefffffffffffff)
				if v11 != 0 {
					p16 = i64(0)
				}
				v1 = p16
				goto l5
			}
		l4:
			v4 = i32(0)
			if uint32(v5) < uint32(i32(-1085)) {
				goto l0
			}
			v1 = i64_shr_u(v7, int64(uint32(i32(-1022)-v5)))
			v1 = v1&i64(1) + v1
			var p17 int32
			if uint64(v1) > uint64(i64(0x1fffffffffffff)) {
				p17 = 1
			}
			v4 = p17
			v1 = int64(uint64(v1) >> 1)
		}
	l5:
		store64(m.memory[uint32(v0):], uint64(v1))
	}
l0:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v4))
l1:
	m.g0 = v3 + i32(32)
}
func (m *Module) fn879(v0, v1 int32) int32 {
	var v2, v3 int32
	var v4, v5 float64
	t0 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	v2 = t0
	v3 = v2 & i32(0x200000)
	t1 := math.Float64frombits(load64(m.memory[uint32(v0):]))
	v4 = t1
	{
		if v2&i32(0x10000000) != 0 {
			t12 := int32(load16(m.memory[int64(uint32(v1))+14:]))
			t13 := v1
			t14 := v4
			var p15 int32
			if v3 != i32(0) {
				p15 = 1
			}
			t16 := m.fn882(t13, t14, p15, t12)
			return t16
		}
		{
			v5 = math.Abs(v4)
			if v5 >= float64(1e+16) {
				goto l1
			}
			var p2 int32
			if v4 != float64(0) {
				p2 = 1
			}
			var p3 int32
			if v5 < float64(0.0001) {
				p3 = 1
			}
			if p2&p3 != 0 {
				goto l1
			}
			t4 := v1
			t5 := v4
			var p6 int32
			if v3 != i32(0) {
				p6 = 1
			}
			t7 := m.fn880(t4, t5, p6, i32(1))
			return t7
		}
	l1:
		t8 := v1
		t9 := v4
		var p10 int32
		if v3 != i32(0) {
			p10 = 1
		}
		t11 := m.fn881(t8, t9, p10)
		return t11
	}
}
func (m *Module) fn880(v0 int32, v1 float64, v2, v3 int32) int32 {
	var v4 int32
	var v5, v6 int64
	var v7 int32
	var v8, v9 int64
	var v10, v11, v12 int32
	var v13 int64
	t0 := m.g0
	v4 = t0 - i32(128)
	m.g0 = v4
	v5 = int64(math.Float64bits(v1))
	v6 = v5 & i64(0xfffffffffffff)
	t1 := v6 | i64(0x10000000000000)
	t2 := v5 << 1 & i64(0x1ffffffffffffe)
	v7 = int32(int64(uint64(v5)>>52)) & i32(2047)
	p3 := t2
	if v7 != 0 {
		p3 = t1
	}
	v8 = p3
	v9 = v8 & i64(1)
	v10 = i32(2)
	{
		var p4 int32
		if v6 == 0 {
			p4 = 1
		}
		v11 = p4
		t6 := v11
		p5 := i32(3)
		if v11 != 0 {
			p5 = i32(2)
		}
		v6 = v5 & i64(0x7ff0000000000000)
		p7 := i32(4)
		if v6 == 0 {
			p7 = p5
		}
		p8 := p7
		if v6 == i64(0x7ff0000000000000) {
			p8 = t6
		}
		switch p8 {
		default:
			goto l0
		case 1:
			v10 = i32(3)
			goto l0
		case 2:
			v10 = i32(4)
			goto l0
		case 3:
			v12 = v7 + i32(-1075)
			v10 = int32(v9) ^ i32(1)
			v13 = i64(1)
			goto l0
		case 4:
			t9 := v8 << 1
			var p10 int32
			if v8 == i64(0x10000000000000) {
				p10 = 1
			}
			v12 = p10
			p11 := t9
			if v12 != 0 {
				p11 = i64(0x40000000000000)
			}
			v8 = p11
			p12 := i64(1)
			if v12 != 0 {
				p12 = i64(2)
			}
			v13 = p12
			v10 = int32(v9) ^ i32(1)
			p13 := i32(-1076)
			if v12 != 0 {
				p13 = i32(-1077)
			}
			v12 = p13 + v7
		}
	}
l0:
	store16(m.memory[int64(uint32(v4))+120:], uint16(v12))
	store64(m.memory[int64(uint32(v4))+112:], uint64(v13))
	store64(m.memory[int64(uint32(v4))+104:], uint64(i64(1)))
	store64(m.memory[int64(uint32(v4))+96:], uint64(v8))
	m.memory[int64(uint32(v4))+122] = byte(v10)
	{
		{
			if uint32(v10&i32(255)) > uint32(i32(1)) {
				v11 = v10 + i32(-2)
				if v11&i32(255) == 0 {
					store32(m.memory[int64(uint32(v4))+40:], uint32(i32(3)))
					store32(m.memory[int64(uint32(v4))+36:], uint32(i32(1098786)))
					store16(m.memory[int64(uint32(v4))+32:], uint16(i32(2)))
					v12 = i32(1)
					v11 = v4 + i32(32)
					v2 = i32(0)
					v10 = i32(1)
					goto l11
				}
				v10 = i32(1)
				var p20 int32
				if v5 < i64(0) {
					p20 = 1
				}
				v12 = p20
				p21 := i32(1098785)
				if v12 != 0 {
					p21 = i32(1098784)
				}
				p22 := i32(1)
				if v12 != 0 {
					p22 = i32(1098784)
				}
				p23 := p22
				if v2 != 0 {
					p23 = p21
				}
				v12 = p23
				p24 := int32(int64(uint64(v5) >> 63))
				if v2 != 0 {
					p24 = i32(1)
				}
				v2 = p24
				if v11&i32(255) != i32(2) {
					store32(m.memory[int64(uint32(v4))+40:], uint32(i32(3)))
					store32(m.memory[int64(uint32(v4))+36:], uint32(i32(1098789)))
					store16(m.memory[int64(uint32(v4))+32:], uint16(i32(2)))
					v11 = v4 + i32(32)
					goto l11
				}
				store16(m.memory[int64(uint32(v4))+32:], uint16(i32(2)))
				if v3&i32(0xffff) != 0 {
					goto l10
				}
				v10 = i32(1)
				store32(m.memory[int64(uint32(v4))+40:], uint32(i32(1)))
				store32(m.memory[int64(uint32(v4))+36:], uint32(i32(1098792)))
				v11 = v4 + i32(32)
				goto l11
			}
			v10 = v3 & i32(0xffff)
			var p14 int32
			if v5 < i64(0) {
				p14 = 1
			}
			v12 = p14
			p15 := i32(1)
			if v12 != 0 {
				p15 = i32(1098784)
			}
			v11 = p15
			p16 := i32(1098785)
			if v12 != 0 {
				p16 = i32(1098784)
			}
			v12 = p16
			v3 = int32(int64(uint64(v5) >> 63))
			m.fn873(v4+i32(32), v4+i32(96), v4+i32(15))
			t17 := int32(load32(m.memory[int64(uint32(v4))+32:]))
			if t17 == 0 {
				goto l6
			}
			t18 := int32(load32(m.memory[int64(uint32(v4))+40:]))
			store32(m.memory[int64(uint32(v4))+88:], uint32(t18))
			t19 := int64(load64(m.memory[int64(uint32(v4))+32:]))
			store64(m.memory[int64(uint32(v4))+80:], uint64(t19))
			goto l7
		}
	l6:
		m.fn874(v4+i32(80), v4+i32(96), v4+i32(15))
	l7:
		p25 := v11
		if v2 != 0 {
			p25 = v12
		}
		v12 = p25
		p26 := v3
		if v2 != 0 {
			p26 = i32(1)
		}
		v2 = p26
		t27 := int32(load32(m.memory[int64(uint32(v4))+80:]))
		t28 := int32(load32(m.memory[int64(uint32(v4))+84:]))
		t29 := int32(load16(m.memory[int64(uint32(v4))+88:]))
		m.fn868(v4, t27, t28, t29, v10, v4+i32(32))
		t30 := int32(load32(m.memory[int64(uint32(v4))+4:]))
		v10 = t30
		t31 := int32(load32(m.memory[uint32(v4):]))
		v11 = t31
		goto l11
	}
l10:
	store32(m.memory[int64(uint32(v4))+48:], uint32(i32(1)))
	store16(m.memory[int64(uint32(v4))+44:], uint16(i32(0)))
	v10 = i32(2)
	store32(m.memory[int64(uint32(v4))+40:], uint32(i32(2)))
	store32(m.memory[int64(uint32(v4))+36:], uint32(i32(1098793)))
	v11 = v4 + i32(32)
l11:
	store32(m.memory[int64(uint32(v4))+92:], uint32(v10))
	store32(m.memory[int64(uint32(v4))+88:], uint32(v11))
	store32(m.memory[int64(uint32(v4))+84:], uint32(v2))
	store32(m.memory[int64(uint32(v4))+80:], uint32(v12))
	t32 := m.fn876(v0, v4+i32(80))
	v10 = t32
	m.g0 = v4 + i32(128)
	return v10
}
func (m *Module) fn881(v0 int32, v1 float64, v2 int32) int32 {
	var v3 int32
	var v4, v5 int64
	var v6 int32
	var v7, v8 int64
	var v9, v10, v11 int32
	var v12 int64
	var v13, v14, v15 int32
	t0 := m.g0
	v3 = t0 - i32(160)
	m.g0 = v3
	v4 = int64(math.Float64bits(v1))
	v5 = v4 & i64(0xfffffffffffff)
	t1 := v5 | i64(0x10000000000000)
	t2 := v4 << 1 & i64(0x1ffffffffffffe)
	v6 = int32(int64(uint64(v4)>>52)) & i32(2047)
	p3 := t2
	if v6 != 0 {
		p3 = t1
	}
	v7 = p3
	v8 = v7 & i64(1)
	v9 = i32(2)
	{
		var p4 int32
		if v5 == 0 {
			p4 = 1
		}
		v10 = p4
		t6 := v10
		p5 := i32(3)
		if v10 != 0 {
			p5 = i32(2)
		}
		v5 = v4 & i64(0x7ff0000000000000)
		p7 := i32(4)
		if v5 == 0 {
			p7 = p5
		}
		p8 := p7
		if v5 == i64(0x7ff0000000000000) {
			p8 = t6
		}
		switch p8 {
		default:
			goto l0
		case 1:
			v9 = i32(3)
			goto l0
		case 2:
			v9 = i32(4)
			goto l0
		case 3:
			v11 = v6 + i32(-1075)
			v9 = int32(v8) ^ i32(1)
			v12 = i64(1)
			goto l0
		case 4:
			t9 := v7 << 1
			var p10 int32
			if v7 == i64(0x10000000000000) {
				p10 = 1
			}
			v11 = p10
			p11 := t9
			if v11 != 0 {
				p11 = i64(0x40000000000000)
			}
			v7 = p11
			p12 := i64(1)
			if v11 != 0 {
				p12 = i64(2)
			}
			v12 = p12
			v9 = int32(v8) ^ i32(1)
			p13 := i32(-1076)
			if v11 != 0 {
				p13 = i32(-1077)
			}
			v11 = p13 + v6
		}
	}
l0:
	store16(m.memory[int64(uint32(v3))+136:], uint16(v11))
	store64(m.memory[int64(uint32(v3))+128:], uint64(v12))
	store64(m.memory[int64(uint32(v3))+120:], uint64(i64(1)))
	store64(m.memory[int64(uint32(v3))+112:], uint64(v7))
	m.memory[int64(uint32(v3))+138] = byte(v9)
	{
		{
			{
				if uint32(v9&i32(255)) > uint32(i32(1)) {
					v9 = v9 + i32(-2)
					if v9&i32(255) == 0 {
						store32(m.memory[int64(uint32(v3))+32:], uint32(i32(3)))
						store32(m.memory[int64(uint32(v3))+28:], uint32(i32(1098786)))
						store16(m.memory[int64(uint32(v3))+24:], uint16(i32(2)))
						v10 = i32(1)
						v2 = i32(0)
						v11 = i32(1)
						goto l10
					}
					v11 = i32(1)
					var p17 int32
					if v4 < i64(0) {
						p17 = 1
					}
					v10 = p17
					p18 := i32(1098785)
					if v10 != 0 {
						p18 = i32(1098784)
					}
					p19 := i32(1)
					if v10 != 0 {
						p19 = i32(1098784)
					}
					p20 := p19
					if v2 != 0 {
						p20 = p18
					}
					v10 = p20
					p21 := int32(int64(uint64(v4) >> 63))
					if v2 != 0 {
						p21 = i32(1)
					}
					v2 = p21
					store16(m.memory[int64(uint32(v3))+24:], uint16(i32(2)))
					if v9&i32(255) == i32(2) {
						store32(m.memory[int64(uint32(v3))+32:], uint32(i32(3)))
						store32(m.memory[int64(uint32(v3))+28:], uint32(i32(1098795)))
						goto l10
					}
					store32(m.memory[int64(uint32(v3))+32:], uint32(i32(3)))
					store32(m.memory[int64(uint32(v3))+28:], uint32(i32(1098789)))
					goto l10
				}
				m.fn873(v3+i32(96), v3+i32(112), v3+i32(7))
				t14 := int32(load32(m.memory[int64(uint32(v3))+96:]))
				if t14 == 0 {
					goto l6
				}
				t15 := int32(load32(m.memory[int64(uint32(v3))+104:]))
				store32(m.memory[int64(uint32(v3))+152:], uint32(t15))
				t16 := int64(load64(m.memory[int64(uint32(v3))+96:]))
				store64(m.memory[int64(uint32(v3))+144:], uint64(t16))
				goto l7
			}
		l6:
			m.fn874(v3+i32(144), v3+i32(112), v3+i32(7))
		l7:
			t22 := int32(load32(m.memory[int64(uint32(v3))+148:]))
			v9 = t22
			if v9 == 0 {
				m.fn7(i32(1102113), i32(33), i32(1102212))
				panic("unreachable")
			}
			t23 := int32(load32(m.memory[int64(uint32(v3))+144:]))
			v11 = t23
			t24 := int32(m.memory[uint32(v11)])
			if uint32(t24) <= uint32(i32(48)) {
				m.fn7(i32(1102164), i32(31), i32(1102228))
				panic("unreachable")
			}
			var p25 int32
			if v4 < i64(0) {
				p25 = 1
			}
			v10 = p25
			p26 := i32(1)
			if v10 != 0 {
				p26 = i32(1098784)
			}
			v13 = p26
			p27 := i32(1098785)
			if v10 != 0 {
				p27 = i32(1098784)
			}
			v10 = p27
			v14 = int32(int64(uint64(v4) >> 63))
			t28 := int32(int16(load16(m.memory[int64(uint32(v3))+152:])))
			v6 = t28
			store32(m.memory[int64(uint32(v3))+32:], uint32(i32(1)))
			store32(m.memory[int64(uint32(v3))+28:], uint32(v11))
			store16(m.memory[int64(uint32(v3))+24:], uint16(i32(2)))
			v15 = v9 + i32(-1)
			if v15 != 0 {
				v9 = v3 + i32(60)
				store32(m.memory[int64(uint32(v3))+56:], uint32(v15))
				store16(m.memory[int64(uint32(v3))+48:], uint16(i32(2)))
				store32(m.memory[int64(uint32(v3))+40:], uint32(i32(1100303)))
				store16(m.memory[int64(uint32(v3))+36:], uint16(i32(2)))
				store32(m.memory[int64(uint32(v3))+44:], uint32(i32(1)))
				store32(m.memory[int64(uint32(v3))+52:], uint32(v11+i32(1)))
				v11 = i32(5)
				goto l14
			}
			v9 = v3 + i32(36)
			v11 = i32(3)
			goto l14
		}
	l14:
		p29 := v13
		if v2 != 0 {
			p29 = v10
		}
		v10 = p29
		p30 := v14
		if v2 != 0 {
			p30 = i32(1)
		}
		v2 = p30
		store16(m.memory[int64(uint32(v9))+12:], uint16(i32(1)))
		store16(m.memory[uint32(v9):], uint16(i32(2)))
		t31 := v9
		var p32 int32
		if v6 < i32(1) {
			p32 = 1
		}
		v13 = p32
		p33 := i32(1)
		if v13 != 0 {
			p33 = i32(2)
		}
		store32(m.memory[int64(uint32(t31))+8:], uint32(p33))
		t35 := v9
		p34 := i32(1102244)
		if v13 != 0 {
			p34 = i32(1102245)
		}
		store32(m.memory[int64(uint32(t35))+4:], uint32(p34))
		t36 := v9
		v6 = v6 + i32(-1)
		t37 := v6
		v6 = v6 >> 31
		store16(m.memory[int64(uint32(t36))+14:], uint16(t37^v6-v6))
	}
l10:
	store32(m.memory[int64(uint32(v3))+108:], uint32(v11))
	store32(m.memory[int64(uint32(v3))+100:], uint32(v2))
	store32(m.memory[int64(uint32(v3))+96:], uint32(v10))
	store32(m.memory[int64(uint32(v3))+104:], uint32(v3+i32(24)))
	t38 := m.fn876(v0, v3+i32(96))
	v9 = t38
	m.g0 = v3 + i32(160)
	return v9
}
func (m *Module) fn882(v0 int32, v1 float64, v2, v3 int32) int32 {
	var v4 int32
	var v5, v6 int64
	var v7 int32
	var v8, v9 int64
	var v10, v11, v12 int32
	var v13 int64
	var v14, v15, v16 int32
	t0 := m.g0
	v4 = t0 - i32(1136)
	m.g0 = v4
	v5 = int64(math.Float64bits(v1))
	v6 = v5 & i64(0xfffffffffffff)
	t1 := v6 | i64(0x10000000000000)
	t2 := v5 << 1 & i64(0x1ffffffffffffe)
	v7 = int32(int64(uint64(v5)>>52)) & i32(2047)
	p3 := t2
	if v7 != 0 {
		p3 = t1
	}
	v8 = p3
	v9 = v8 & i64(1)
	v10 = i32(2)
	{
		var p4 int32
		if v6 == 0 {
			p4 = 1
		}
		v11 = p4
		t6 := v11
		p5 := i32(3)
		if v11 != 0 {
			p5 = i32(2)
		}
		v6 = v5 & i64(0x7ff0000000000000)
		p7 := i32(4)
		if v6 == 0 {
			p7 = p5
		}
		p8 := p7
		if v6 == i64(0x7ff0000000000000) {
			p8 = t6
		}
		switch p8 {
		default:
			goto l0
		case 1:
			v10 = i32(3)
			goto l0
		case 2:
			v10 = i32(4)
			goto l0
		case 3:
			v12 = v7 + i32(-1075)
			v10 = int32(v9) ^ i32(1)
			v13 = i64(1)
			goto l0
		case 4:
			t9 := v8 << 1
			var p10 int32
			if v8 == i64(0x10000000000000) {
				p10 = 1
			}
			v12 = p10
			p11 := t9
			if v12 != 0 {
				p11 = i64(0x40000000000000)
			}
			v8 = p11
			p12 := i64(1)
			if v12 != 0 {
				p12 = i64(2)
			}
			v13 = p12
			v10 = int32(v9) ^ i32(1)
			p13 := i32(-1076)
			if v12 != 0 {
				p13 = i32(-1077)
			}
			v12 = p13 + v7
		}
	}
l0:
	v11 = v3 & i32(0xffff)
	store16(m.memory[int64(uint32(v4))+1128:], uint16(v12))
	store64(m.memory[int64(uint32(v4))+1120:], uint64(v13))
	store64(m.memory[int64(uint32(v4))+1112:], uint64(i64(1)))
	store64(m.memory[int64(uint32(v4))+1104:], uint64(v8))
	m.memory[int64(uint32(v4))+1130] = byte(v10)
	{
		{
			if uint32(v10&i32(255)) > uint32(i32(1)) {
				v7 = v10 + i32(-2)
				if v7&i32(255) == 0 {
					store32(m.memory[int64(uint32(v4))+1048:], uint32(i32(3)))
					store32(m.memory[int64(uint32(v4))+1044:], uint32(i32(1098786)))
					store16(m.memory[int64(uint32(v4))+1040:], uint16(i32(2)))
					v12 = i32(1)
					v3 = v4 + i32(1040)
					v2 = i32(0)
					v10 = i32(1)
					goto l10
				}
				v10 = i32(1)
				var p15 int32
				if v5 < i64(0) {
					p15 = 1
				}
				v12 = p15
				p16 := i32(1098785)
				if v12 != 0 {
					p16 = i32(1098784)
				}
				p17 := i32(1)
				if v12 != 0 {
					p17 = i32(1098784)
				}
				p18 := p17
				if v2 != 0 {
					p18 = p16
				}
				v12 = p18
				p19 := int32(int64(uint64(v5) >> 63))
				if v2 != 0 {
					p19 = i32(1)
				}
				v2 = p19
				if v7&i32(255) != i32(2) {
					store32(m.memory[int64(uint32(v4))+1048:], uint32(i32(3)))
					store32(m.memory[int64(uint32(v4))+1044:], uint32(i32(1098789)))
					store16(m.memory[int64(uint32(v4))+1040:], uint16(i32(2)))
					v3 = v4 + i32(1040)
					goto l10
				}
				store16(m.memory[int64(uint32(v4))+1040:], uint16(i32(2)))
				if v3&i32(0xffff) != 0 {
					store32(m.memory[int64(uint32(v4))+1056:], uint32(v11))
					store16(m.memory[int64(uint32(v4))+1052:], uint16(i32(0)))
					v10 = i32(2)
					store32(m.memory[int64(uint32(v4))+1048:], uint32(i32(2)))
					store32(m.memory[int64(uint32(v4))+1044:], uint32(i32(1098793)))
					v3 = v4 + i32(1040)
					goto l10
				}
				v10 = i32(1)
				store32(m.memory[int64(uint32(v4))+1048:], uint32(i32(1)))
				store32(m.memory[int64(uint32(v4))+1044:], uint32(i32(1098792)))
				v3 = v4 + i32(1040)
				goto l10
			}
			v10 = int32(int16(v12))
			p14 := i32(5)
			if v10 < i32(0) {
				p14 = i32(-12)
			}
			v10 = p14 * v10
			if uint32(v10) < uint32(i32(16064)) {
				goto l6
			}
			m.fn7(i32(1098860), i32(37), i32(1098900))
			panic("unreachable")
		}
	l6:
		;
		var p20 int32
		if v5 < i64(0) {
			p20 = 1
		}
		v7 = p20
		p21 := i32(1)
		if v7 != 0 {
			p21 = i32(1098784)
		}
		v14 = p21
		p22 := i32(1098785)
		if v7 != 0 {
			p22 = i32(1098784)
		}
		v7 = p22
		v15 = int32(int64(uint64(v5) >> 63))
		t23 := v4 + i32(1040)
		t24 := v8
		t25 := v12
		t26 := v4 + i32(16)
		v16 = int32(uint32(v10)>>4) + i32(21)
		t28 := v16
		p27 := i32(-0x8000)
		if int32(int16(v3)) > i32(-1) {
			p27 = i32(0) - v3
		}
		v10 = p27
		m.fn869(t23, t24, t25, t26, t28, v10)
		v10 = int32(int16(v10))
		{
			t29 := int32(load32(m.memory[int64(uint32(v4))+1040:]))
			if t29 == 0 {
				goto l11
			}
			t30 := int32(load32(m.memory[int64(uint32(v4))+1048:]))
			store32(m.memory[int64(uint32(v4))+1096:], uint32(t30))
			t31 := int64(load64(m.memory[int64(uint32(v4))+1040:]))
			store64(m.memory[int64(uint32(v4))+1088:], uint64(t31))
			goto l12
		}
	l11:
		m.fn871(v4+i32(1088), v4+i32(1104), v4+i32(16), v16, v10)
	l12:
		p32 := v14
		if v2 != 0 {
			p32 = v7
		}
		v12 = p32
		p33 := v15
		if v2 != 0 {
			p33 = i32(1)
		}
		v2 = p33
		{
			t34 := int32(int16(load16(m.memory[int64(uint32(v4))+1096:])))
			v7 = t34
			if v7 <= v10 {
				goto l13
			}
			t35 := int32(load32(m.memory[int64(uint32(v4))+1088:]))
			t36 := int32(load32(m.memory[int64(uint32(v4))+0x444:]))
			m.fn868(v4+i32(8), t35, t36, v7, v11, v4+i32(1040))
			t37 := int32(load32(m.memory[int64(uint32(v4))+12:]))
			v10 = t37
			t38 := int32(load32(m.memory[int64(uint32(v4))+8:]))
			v3 = t38
			goto l10
		}
	l13:
		v10 = i32(2)
		store16(m.memory[int64(uint32(v4))+1040:], uint16(i32(2)))
		if v3&i32(0xffff) != 0 {
			goto l14
		}
		v10 = i32(1)
		store32(m.memory[int64(uint32(v4))+1048:], uint32(i32(1)))
		store32(m.memory[int64(uint32(v4))+1044:], uint32(i32(1098792)))
		v3 = v4 + i32(1040)
		goto l10
	l14:
		store32(m.memory[int64(uint32(v4))+1056:], uint32(v11))
		store16(m.memory[int64(uint32(v4))+1052:], uint16(i32(0)))
		store32(m.memory[int64(uint32(v4))+1048:], uint32(i32(2)))
		store32(m.memory[int64(uint32(v4))+1044:], uint32(i32(1098793)))
		v3 = v4 + i32(1040)
	}
l10:
	store32(m.memory[int64(uint32(v4))+1100:], uint32(v10))
	store32(m.memory[int64(uint32(v4))+1096:], uint32(v3))
	store32(m.memory[int64(uint32(v4))+0x444:], uint32(v2))
	store32(m.memory[int64(uint32(v4))+1088:], uint32(v12))
	t39 := m.fn876(v0, v4+i32(1088))
	v10 = t39
	m.g0 = v4 + i32(1136)
	return v10
}
func (m *Module) fn883(v0 int32, v1 float64, v2, v3 int32) int32 {
	var v4 int32
	var v5, v6 int64
	var v7 int32
	var v8, v9 int64
	var v10, v11, v12 int32
	var v13 int64
	var v14, v15, v16, v17, v18, v19 int32
	t0 := m.g0
	v4 = t0 - i32(1168)
	m.g0 = v4
	v5 = int64(math.Float64bits(v1))
	v6 = v5 & i64(0xfffffffffffff)
	t1 := v6 | i64(0x10000000000000)
	t2 := v5 << 1 & i64(0x1ffffffffffffe)
	v7 = int32(int64(uint64(v5)>>52)) & i32(2047)
	p3 := t2
	if v7 != 0 {
		p3 = t1
	}
	v8 = p3
	v9 = v8 & i64(1)
	v10 = i32(2)
	{
		var p4 int32
		if v6 == 0 {
			p4 = 1
		}
		v11 = p4
		t6 := v11
		p5 := i32(3)
		if v11 != 0 {
			p5 = i32(2)
		}
		v6 = v5 & i64(0x7ff0000000000000)
		p7 := i32(4)
		if v6 == 0 {
			p7 = p5
		}
		p8 := p7
		if v6 == i64(0x7ff0000000000000) {
			p8 = t6
		}
		switch p8 {
		default:
			goto l0
		case 1:
			v10 = i32(3)
			goto l0
		case 2:
			v10 = i32(4)
			goto l0
		case 3:
			v12 = v7 + i32(-1075)
			v10 = int32(v9) ^ i32(1)
			v13 = i64(1)
			goto l0
		case 4:
			t9 := v8 << 1
			var p10 int32
			if v8 == i64(0x10000000000000) {
				p10 = 1
			}
			v12 = p10
			p11 := t9
			if v12 != 0 {
				p11 = i64(0x40000000000000)
			}
			v8 = p11
			p12 := i64(1)
			if v12 != 0 {
				p12 = i64(2)
			}
			v13 = p12
			v10 = int32(v9) ^ i32(1)
			p13 := i32(-1076)
			if v12 != 0 {
				p13 = i32(-1077)
			}
			v12 = p13 + v7
		}
	}
l0:
	v11 = v3 & i32(0xffff)
	store16(m.memory[int64(uint32(v4))+1144:], uint16(v12))
	store64(m.memory[int64(uint32(v4))+1136:], uint64(v13))
	store64(m.memory[int64(uint32(v4))+1128:], uint64(i64(1)))
	store64(m.memory[int64(uint32(v4))+1120:], uint64(v8))
	m.memory[int64(uint32(v4))+1146] = byte(v10)
	{
		{
			if uint32(v10&i32(255)) > uint32(i32(1)) {
				v7 = v10 + i32(-2)
				if v7&i32(255) == 0 {
					store32(m.memory[int64(uint32(v4))+1040:], uint32(i32(3)))
					store32(m.memory[int64(uint32(v4))+1036:], uint32(i32(1098786)))
					store16(m.memory[int64(uint32(v4))+1032:], uint16(i32(2)))
					v12 = i32(1)
					v2 = i32(0)
					v10 = i32(1)
					goto l10
				}
				v10 = i32(1)
				var p17 int32
				if v5 < i64(0) {
					p17 = 1
				}
				v12 = p17
				p18 := i32(1098785)
				if v12 != 0 {
					p18 = i32(1098784)
				}
				p19 := i32(1)
				if v12 != 0 {
					p19 = i32(1098784)
				}
				p20 := p19
				if v2 != 0 {
					p20 = p18
				}
				v12 = p20
				p21 := int32(int64(uint64(v5) >> 63))
				if v2 != 0 {
					p21 = i32(1)
				}
				v2 = p21
				if v7&i32(255) != i32(2) {
					store32(m.memory[int64(uint32(v4))+1040:], uint32(i32(3)))
					store32(m.memory[int64(uint32(v4))+1036:], uint32(i32(1098789)))
					store16(m.memory[int64(uint32(v4))+1032:], uint16(i32(2)))
					goto l10
				}
				store16(m.memory[int64(uint32(v4))+1032:], uint16(i32(2)))
				if v3&i32(0xffff) != 0 {
					store32(m.memory[int64(uint32(v4))+1064:], uint32(i32(2)))
					store32(m.memory[int64(uint32(v4))+1060:], uint32(i32(1098798)))
					store16(m.memory[int64(uint32(v4))+1056:], uint16(i32(2)))
					store32(m.memory[int64(uint32(v4))+1048:], uint32(v11))
					store16(m.memory[int64(uint32(v4))+1044:], uint16(i32(0)))
					store32(m.memory[int64(uint32(v4))+1040:], uint32(i32(2)))
					store32(m.memory[int64(uint32(v4))+1036:], uint32(i32(1098793)))
					v10 = i32(3)
					goto l10
				}
				store32(m.memory[int64(uint32(v4))+1040:], uint32(i32(3)))
				store32(m.memory[int64(uint32(v4))+1036:], uint32(i32(1098795)))
				goto l10
			}
			t14 := v11 + i32(1)
			v10 = int32(int16(v12))
			p15 := i32(5)
			if v10 < i32(0) {
				p15 = i32(-12)
			}
			v10 = int32(uint32(p15*v10)>>4) + i32(21)
			p16 := v10
			if uint32(v10) > uint32(v11) {
				p16 = t14
			}
			v10 = p16
			if uint32(v10) < uint32(i32(1025)) {
				goto l6
			}
			m.fn7(i32(1098800), i32(41), i32(1098844))
			panic("unreachable")
		}
	l6:
		m.fn869(v4+i32(1104), v8, v12, v4+i32(8), v10, i32(0x8000))
		{
			t22 := int32(load32(m.memory[int64(uint32(v4))+1104:]))
			if t22 == 0 {
				goto l11
			}
			t23 := int32(load32(m.memory[int64(uint32(v4))+1112:]))
			store32(m.memory[int64(uint32(v4))+1160:], uint32(t23))
			t24 := int64(load64(m.memory[int64(uint32(v4))+1104:]))
			store64(m.memory[int64(uint32(v4))+1152:], uint64(t24))
			goto l12
		}
	l11:
		m.fn871(v4+i32(1152), v4+i32(1120), v4+i32(8), v10, i32(0x8000))
	l12:
		{
			t25 := int32(load32(m.memory[int64(uint32(v4))+1156:]))
			v14 = t25
			if v14 == 0 {
				m.fn7(i32(1102113), i32(33), i32(1102212))
				panic("unreachable")
			}
			t26 := int32(load32(m.memory[int64(uint32(v4))+1152:]))
			v15 = t26
			t27 := int32(m.memory[uint32(v15)])
			if uint32(t27) <= uint32(i32(48)) {
				m.fn7(i32(1102164), i32(31), i32(1102228))
				panic("unreachable")
			}
			var p28 int32
			if v5 < i64(0) {
				p28 = 1
			}
			v7 = p28
			v5 = int64(uint64(v5) >> 63)
			v12 = v4 + i32(1056)
			t29 := int32(int16(load16(m.memory[int64(uint32(v4))+1160:])))
			v16 = t29
			v10 = i32(1)
			store32(m.memory[int64(uint32(v4))+1040:], uint32(i32(1)))
			store32(m.memory[int64(uint32(v4))+1036:], uint32(v15))
			v17 = i32(2)
			store16(m.memory[int64(uint32(v4))+1032:], uint16(i32(2)))
			v14 = v14 + i32(-1)
			if v14 != 0 {
				v3 = v4 + i32(1064)
				v18 = v4 + i32(1060)
				store32(m.memory[int64(uint32(v4))+1048:], uint32(i32(1100303)))
				store16(m.memory[int64(uint32(v4))+1044:], uint16(i32(2)))
				store32(m.memory[int64(uint32(v4))+1052:], uint32(i32(1)))
				v19 = v15 + i32(1)
				v17 = i32(3)
				v15 = v12
				v10 = v14
				v12 = v4 + i32(1068)
				goto l17
			}
			if v3&i32(0xffff) == 0 {
				goto l16
			}
			v3 = v4 + i32(1052)
			v18 = v4 + i32(1048)
			v15 = v4 + i32(1044)
			v19 = i32(1100303)
			goto l17
		}
	l17:
		store32(m.memory[uint32(v18):], uint32(v19))
		store16(m.memory[uint32(v15):], uint16(i32(2)))
		store32(m.memory[uint32(v3):], uint32(v10))
		if uint32(v11) > uint32(v14) {
			goto l18
		}
		v10 = v17
		goto l16
	l18:
		store16(m.memory[uint32(v12):], uint16(i32(0)))
		store32(m.memory[int64(uint32(v12))+4:], uint32(v11-v14))
		v10 = v17 + i32(1)
	l16:
		p30 := i32(1)
		if v7 != 0 {
			p30 = i32(1098784)
		}
		v12 = p30
		p31 := i32(1098785)
		if v7 != 0 {
			p31 = i32(1098784)
		}
		v11 = p31
		v3 = int32(v5)
		if v16 < i32(1) {
			goto l19
		}
		v7 = v4 + i32(1032) + v10*i32(12)
		store32(m.memory[int64(uint32(v7))+8:], uint32(i32(1)))
		store32(m.memory[int64(uint32(v7))+4:], uint32(i32(1102244)))
		store16(m.memory[uint32(v7):], uint16(i32(2)))
		v7 = v16 + i32(-1)
		goto l20
	l19:
		v7 = v4 + i32(1032) + v10*i32(12)
		store32(m.memory[int64(uint32(v7))+8:], uint32(i32(2)))
		store32(m.memory[int64(uint32(v7))+4:], uint32(i32(1102245)))
		store16(m.memory[uint32(v7):], uint16(i32(2)))
		v7 = i32(1) - v16
	l20:
		p32 := v12
		if v2 != 0 {
			p32 = v11
		}
		v12 = p32
		p33 := v3
		if v2 != 0 {
			p33 = i32(1)
		}
		v2 = p33
		v11 = v4 + i32(1032) + v10*i32(12)
		store16(m.memory[int64(uint32(v11))+14:], uint16(v7))
		store16(m.memory[int64(uint32(v11))+12:], uint16(i32(1)))
		v10 = v10 + i32(2)
	}
l10:
	store32(m.memory[int64(uint32(v4))+1116:], uint32(v10))
	store32(m.memory[int64(uint32(v4))+1108:], uint32(v2))
	store32(m.memory[int64(uint32(v4))+1104:], uint32(v12))
	store32(m.memory[int64(uint32(v4))+1112:], uint32(v4+i32(1032)))
	t34 := m.fn876(v0, v4+i32(1104))
	v10 = t34
	m.g0 = v4 + i32(1168)
	return v10
}
func (m *Module) fn884(v0, v1 int32) int32 {
	var v2 int32
	t0 := m.g0
	v2 = t0 - i32(32)
	m.g0 = v2
	{
		{
			t1 := int32(m.memory[int64(uint32(v0))+4])
			if t1 != i32(1) {
				goto l0
			}
			t2 := int32(m.memory[int64(uint32(v0))+5])
			m.memory[int64(uint32(v2))+15] = byte(t2)
			store64(m.memory[int64(uint32(v2))+24:], uint64(int64(uint32(i32(3)))<<32|int64(uint32(v0))))
			store64(m.memory[int64(uint32(v2))+16:], uint64(int64(uint32(i32(18)))<<32|int64(uint32(v2+i32(15)))))
			t3 := int32(load32(m.memory[uint32(v1):]))
			t4 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t5 := m.fn49(t3, t4, i32(0x1002dd), v2+i32(16))
			v0 = t5
			goto l1
		}
	l0:
		store64(m.memory[int64(uint32(v2))+16:], uint64(int64(uint32(i32(3)))<<32|int64(uint32(v0))))
		t6 := int32(load32(m.memory[uint32(v1):]))
		t7 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t8 := m.fn49(t6, t7, i32(1049358), v2+i32(16))
		v0 = t8
	}
l1:
	m.g0 = v2 + i32(32)
	return v0
}
func (m *Module) fn885(v0, v1 int32) int32 {
	var v2, v3, v4 int32
	{
		if v0 <= i32(-1) {
			goto l0
		}
		{
			if v0 != 0 {
				goto l1
			}
			v2 = i32(1)
			goto l2
		l1:
			t0 := m.fn11(v0)
			v2 = t0
			if v2 == 0 {
				m.fn16(i32(1), v0)
				panic("unreachable")
			}
			t1 := int32(m.memory[uint32(v2+i32(-4))])
			if t1&i32(3) == 0 {
				goto l2
			}
			if v0 == 0 {
				goto l2
			}
			memory_zero(m.memory, uint32(v2), uint32(v0))
		}
	l2:
		if uint32(v1) > uint32(i32(0x3fffffff)) {
			goto l0
		}
		v3 = v1 << 2
		if uint32(v3) >= uint32(i32(0x7ffffffd)) {
			goto l0
		}
		t2 := m.fn11(v3)
		v4 = t2
		if v4 == 0 {
			m.fn16(i32(4), v3)
			panic("unreachable")
		}
		{
			t3 := int32(m.memory[uint32(v4+i32(-4))])
			if t3&i32(3) == 0 {
				goto l5
			}
			if v3 == 0 {
				goto l5
			}
			memory_zero(m.memory, uint32(v4), uint32(v3))
		}
	l5:
		t4 := m.fn11(i32(64))
		v3 = t4
		if v3 == 0 {
			m.fn27(i32(8), i32(64))
			panic("unreachable")
		}
		store32(m.memory[int64(uint32(v3))+56:], uint32(v0))
		store32(m.memory[int64(uint32(v3))+52:], uint32(v2))
		store32(m.memory[int64(uint32(v3))+48:], uint32(v0))
		store32(m.memory[int64(uint32(v3))+44:], uint32(i32(0)))
		store32(m.memory[int64(uint32(v3))+40:], uint32(v1))
		store32(m.memory[int64(uint32(v3))+36:], uint32(v4))
		store32(m.memory[int64(uint32(v3))+32:], uint32(v1))
		store64(m.memory[uint32(v3):], uint64(i64(0)))
		return v3
	}
l0:
	m.fn15()
	panic("unreachable")
}
func (m *Module) fn886(v0, v1, v2 int32) {
	var v3, v4, v5, v6, v7, v8, v9 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+44:]))
		v3 = t0
		t1 := int32(load32(m.memory[int64(uint32(v0))+40:]))
		t2 := v3
		v4 = t1
		if uint32(t2) > uint32(v4) {
			m.fn124(i32(0), v3, v4, i32(1139052))
			panic("unreachable")
		}
		{
			if v3 != 0 {
				goto l1
			}
			v5 = i32(0)
			goto l2
		l1:
			t3 := int32(load32(m.memory[int64(uint32(v0))+36:]))
			t4 := int32(load32(m.memory[uint32(t3+v3<<2+i32(-4)):]))
			v5 = t4
		}
	l2:
		{
			v6 = v5 + v2
			t5 := int32(load32(m.memory[int64(uint32(v0))+56:]))
			t6 := v6
			v3 = t5
			if uint32(t6) <= uint32(v3) {
				goto l3
			}
			v7 = v0 + i32(48)
		l6:
			{
				{
					v4 = v3 << 1
					p7 := i32(4)
					if uint32(v4) > uint32(i32(4)) {
						p7 = v4
					}
					v4 = p7 - v3
					t8 := int32(load32(m.memory[int64(uint32(v0))+48:]))
					if uint32(v4) <= uint32(t8-v3) {
						goto l4
					}
					m.fn887(v7, v3, v4, i32(1), i32(1))
					t9 := int32(load32(m.memory[int64(uint32(v0))+56:]))
					v3 = t9
				}
			l4:
				t10 := int32(load32(m.memory[int64(uint32(v0))+52:]))
				v8 = t10 + v3
				v9 = v4 + i32(-1)
				if v9 == 0 {
					goto l5
				}
				memory_zero(m.memory, uint32(v8), uint32(v9))
			l5:
				t11 := v0
				v3 = v3 + v4
				store32(m.memory[int64(uint32(t11))+56:], uint32(v3))
				m.memory[uint32(v8+v9)] = byte(i32(0))
				if uint32(v6) > uint32(v3) {
					goto l6
				}
			}
		}
	l3:
		if uint32(v6) < uint32(v5) {
			m.fn124(v5, v6, v3, i32(1139004))
			panic("unreachable")
		}
		{
			if v2 == 0 {
				goto l8
			}
			t12 := int32(load32(m.memory[int64(uint32(v0))+52:]))
			memory_copy(m.memory, uint32(t12+v5), uint32(v1), uint32(v2))
		}
	l8:
		{
			t13 := int32(load32(m.memory[int64(uint32(v0))+44:]))
			v4 = t13
			t14 := int32(load32(m.memory[int64(uint32(v0))+40:]))
			t15 := v4
			v3 = t14
			if uint32(t15) < uint32(v3) {
				goto l9
			}
			{
				v8 = v3 << 1
				p16 := i32(4)
				if uint32(v8) > uint32(i32(4)) {
					p16 = v8
				}
				v8 = p16 - v3
				t17 := int32(load32(m.memory[int64(uint32(v0))+32:]))
				if uint32(v8) <= uint32(t17-v3) {
					goto l10
				}
				m.fn887(v0+i32(32), v3, v8, i32(4), i32(4))
				t18 := int32(load32(m.memory[int64(uint32(v0))+44:]))
				v4 = t18
				t19 := int32(load32(m.memory[int64(uint32(v0))+40:]))
				v3 = t19
			}
		l10:
			t20 := int32(load32(m.memory[int64(uint32(v0))+36:]))
			v9 = t20
			v5 = v8<<2 + i32(-4)
			if v5 == 0 {
				goto l11
			}
			memory_zero(m.memory, uint32(v9+v3<<2), uint32(v5))
		l11:
			t21 := v0
			v3 = v3 + v8
			store32(m.memory[int64(uint32(t21))+40:], uint32(v3))
			store32(m.memory[uint32(v9+v3<<2+i32(-4)):], uint32(i32(0)))
		}
	l9:
		if uint32(v4) >= uint32(v3) {
			m.fn36(v4, v3, i32(1139036))
			panic("unreachable")
		}
		store32(m.memory[int64(uint32(v0))+44:], uint32(v4+i32(1)))
		t22 := int32(load32(m.memory[int64(uint32(v0))+36:]))
		store32(m.memory[uint32(t22+v4<<2):], uint32(v6))
		return
	}
}
func (m *Module) fn887(v0, v1, v2, v3, v4 int32) {
	var v5 int32
	t0 := m.g0
	v5 = t0 - i32(16)
	m.g0 = v5
	v1 = v2 + v1
	if uint32(v1) >= uint32(v2) {
		goto l0
	}
	m.fn16(i32(0), i32(0))
	panic("unreachable")
l0:
	t1 := int32(load32(m.memory[uint32(v0):]))
	t2 := v5 + i32(4)
	v2 = t1
	t3 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t4 := v2
	t5 := v1
	v2 = v2 << 1
	p6 := v2
	if uint32(v1) > uint32(v2) {
		p6 = t5
	}
	v2 = p6
	t8 := v2
	p7 := i32(4)
	if v4 == i32(1) {
		p7 = i32(8)
	}
	v1 = p7
	p9 := v1
	if uint32(v2) > uint32(v1) {
		p9 = t8
	}
	v2 = p9
	m.fn888(t2, t4, t3, v2, v3, v4)
	{
		t10 := int32(load32(m.memory[int64(uint32(v5))+4:]))
		if t10 != i32(1) {
			goto l1
		}
		t11 := int32(load32(m.memory[int64(uint32(v5))+8:]))
		t12 := int32(load32(m.memory[int64(uint32(v5))+12:]))
		m.fn16(t11, t12)
		panic("unreachable")
	}
l1:
	t13 := int32(load32(m.memory[int64(uint32(v5))+8:]))
	v4 = t13
	store32(m.memory[uint32(v0):], uint32(v2))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v4))
	m.g0 = v5 + i32(16)
}
func (m *Module) fn888(v0, v1, v2, v3, v4, v5 int32) {
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
			t0 := m.fn23(v2, v5*v1, v4, v3)
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
		t1 := m.fn11(v3)
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
func (m *Module) fn889(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7, v8 int32
	v2 = i32(0)
l19:
	v3 = i32(0)
	{
		{
			v4 = v0 + v2
			v5 = (i32(0) - v4) & i32(3)
			t0 := v5 | i32(8)
			v6 = v1 - v2
			if uint32(t0) > uint32(v6) {
				goto l0
			}
			{
				if v5 == 0 {
					goto l1
				}
				v3 = i32(0)
				t1 := int32(int8(m.memory[uint32(v4)]))
				v7 = t1
				if v7 < i32(0) {
					goto l2
				}
				if v5 == i32(1) {
					goto l1
				}
				{
					t2 := int32(int8(m.memory[int64(uint32(v4))+1]))
					v7 = t2
					if v7 >= i32(0) {
						if v5 == i32(2) {
							goto l1
						}
						t3 := int32(int8(m.memory[int64(uint32(v4))+2]))
						v7 = t3
						if v7 >= i32(0) {
							goto l1
						}
						v3 = i32(2)
						goto l2
					}
					v3 = i32(1)
					goto l2
				}
			}
		l1:
			v8 = v6 + i32(-8)
			v3 = v5
		l5:
			{
				v7 = v4 + v3
				t4 := int32(load32(m.memory[uint32(v7+i32(4)):]))
				v5 = t4 & i32(-2139062144)
				t5 := int32(load32(m.memory[uint32(v7):]))
				t6 := v5
				v7 = t5 & i32(-2139062144)
				if t6|v7 != 0 {
					goto l4
				}
				v3 = v3 + i32(8)
				if uint32(v3) <= uint32(v8) {
					goto l5
				}
			}
		}
	l0:
		if uint32(v3) < uint32(v6) {
		l7:
			{
				t7 := int32(int8(m.memory[uint32(v4+v3)]))
				v7 = t7
				if v7 < i32(0) {
					goto l2
				}
				v3 = v3 + i32(1)
				if uint32(v3) < uint32(v6) {
					goto l7
				}
			}
			return v1
		}
		return v1
	l4:
		if v7 == 0 {
			goto l8
		}
		v7 = int32(uint32(int32(bits.TrailingZeros32(uint32(v7)))) >> 3)
		goto l9
	l8:
		v7 = int32(uint32(int32(bits.TrailingZeros32(uint32(v5))))>>3) + i32(4)
	l9:
		t8 := v4
		v3 = v7 + v3
		t9 := int32(m.memory[uint32(t8+v3)])
		v7 = t9
	}
l2:
	v3 = v3 + v2
	if uint32(v3+i32(4)) > uint32(v1) {
		goto l10
	}
l16:
	{
		v4 = v3
		if uint32((v7+i32(62))&i32(255)) < uint32(i32(30)) {
			goto l11
		}
		v3 = v7 & i32(255)
		if uint32(v3) > uint32(i32(239)) {
			v7 = v4 + v0
			t15 := int32(m.memory[uint32(v7+i32(3))])
			t16 := int32(m.memory[int64(uint32(v3))+1270848])
			t17 := int32(m.memory[uint32(v7+i32(1))])
			t18 := int32(m.memory[int64(uint32(t17))+1270720])
			t19 := int32(m.memory[uint32(v7+i32(2))])
			if t15&i32(192)<<2|(t16&t18|int32(uint32(t19)>>6)) != i32(514) {
				goto l18
			}
			v3 = v4 + i32(4)
			if uint32(v4+i32(8)) > uint32(v1) {
				goto l10
			}
			t20 := int32(int8(m.memory[uint32(v0+v3)]))
			v7 = t20
			if v7 < i32(0) {
				goto l16
			}
			v2 = v4 + i32(5)
			goto l17
		}
		v3 = v4
	l15:
		{
			{
				t10 := int32(m.memory[int64(uint32(v7&i32(255)))+1270848])
				v7 = v0 + v3
				t11 := int32(m.memory[uint32(v7+i32(1))])
				t12 := int32(m.memory[int64(uint32(t11))+1270720])
				t13 := int32(m.memory[uint32(v7+i32(2))])
				if t10&t12|int32(uint32(t13)>>6) == i32(2) {
					goto l13
				}
				return v3
			}
		l13:
			if uint32(v3+i32(7)) > uint32(v1) {
				v3 = v3 + i32(3)
				goto l10
			}
			v3 = v3 + i32(3)
			t14 := int32(m.memory[uint32(v7+i32(3))])
			v4 = t14
			v7 = int32(int8(v4))
			if v4&i32(240) == i32(224) {
				goto l15
			}
		}
		if v7 < i32(0) {
			goto l16
		}
		v2 = v3 + i32(-3) + i32(4)
		goto l17
	l11:
		t21 := int32(int8(m.memory[uint32(v4+v0+i32(1))]))
		if t21 >= i32(-64) {
			goto l18
		}
		v3 = v4 + i32(2)
		if uint32(v4+i32(6)) > uint32(v1) {
			goto l10
		}
		t22 := int32(int8(m.memory[uint32(v0+v3)]))
		v7 = t22
		if v7 <= i32(-1) {
			goto l16
		}
	}
	v2 = v4 + i32(3)
l17:
	if uint32(v2) <= uint32(v1) {
		goto l19
	}
	m.fn124(v2, v1, v1, i32(1270688))
	panic("unreachable")
l10:
	if uint32(v3) < uint32(v1) {
		goto l27
	}
	return v3
l27:
	{
		{
			v6 = v0 + v3
			t23 := int32(int8(m.memory[uint32(v6)]))
			v7 = t23
			if v7 > i32(-1) {
				v3 = v3 + i32(1)
				goto l24
			}
			if uint32((v7+i32(62))&i32(255)) < uint32(i32(30)) {
				v7 = v3 + i32(2)
				if uint32(v7) <= uint32(v1) {
					goto l25
				}
				return v3
			}
			if uint32(v7) < uint32(i32(-16)) {
				v4 = v3 + i32(3)
				if uint32(v4) <= uint32(v1) {
					t24 := int32(m.memory[int64(uint32(v7&i32(255)))+1270848])
					t25 := int32(m.memory[int64(uint32(v6))+1])
					t26 := int32(m.memory[int64(uint32(t25))+1270720])
					t27 := int32(m.memory[int64(uint32(v6))+2])
					p28 := v3
					if t24&t26|int32(uint32(t27)>>6) == i32(2) {
						p28 = v4
					}
					v4 = p28
					goto l18
				}
				return v3
			}
			return v3
		}
	l25:
		v4 = v3
		v3 = v7
		t29 := int32(int8(m.memory[int64(uint32(v6))+1]))
		if t29 >= i32(-64) {
			goto l18
		}
	}
l24:
	v4 = v3
	if uint32(v3) < uint32(v1) {
		goto l27
	}
l18:
	return v4
}
func (m *Module) fn890(v0, v1, v2 int32) {
	var v3, v4, v5 int32
	var v6 int64
	var v7 int32
	t0 := m.g0
	v3 = t0 - i32(48)
	m.g0 = v3
	{
		t1 := int32(m.memory[int64(uint32(v1))+24])
		switch t1 {
		default:
			v4 = i32(0)
			if uint32(v2) <= uint32(i32(-4)) {
				v6 = int64(uint32(int32(uint32(v2+i32(1))>>1))) * i64(3)
				if int32(int64(uint64(v6)>>32)) == 0 {
					v5 = int32(v6)
					if v5 != i32(-1) {
						v5 = v5 + i32(1)
						t4 := v5
						v7 = v2 + i32(3)
						p5 := v7
						if uint32(v5) > uint32(v7) {
							p5 = t4
						}
						v5 = p5
						t6 := int32(load32(m.memory[int64(uint32(v1))+20:]))
						v7 = t6
						if v7 == i32(1143808) {
							goto l13
						}
						if v7 == i32(1139704) {
							goto l13
						}
						if v7 == i32(1143836) {
							goto l13
						}
						m.fn894(v3+i32(8), v1, v2)
						t7 := int32(load32(m.memory[int64(uint32(v3))+8:]))
						if t7 != i32(1) {
							goto l7
						}
						t8 := int32(load32(m.memory[int64(uint32(v3))+12:]))
						v1 = t8
						p9 := v5
						if uint32(v1) > uint32(v5) {
							p9 = v1
						}
						v5 = p9
						goto l13
					}
					goto l7
				}
				goto l7
			}
			goto l7
		case 1, 2, 3, 9:
			m.fn894(v3+i32(16), v1, v2)
			t2 := int32(load32(m.memory[int64(uint32(v3))+20:]))
			v5 = t2
			t3 := int32(load32(m.memory[int64(uint32(v3))+16:]))
			v4 = t3
			goto l7
		case 4, 5:
			v4 = i32(0)
			if uint32(v2) <= uint32(i32(-3)) {
				v5 = v2 + i32(5)
				t10 := v5
				v2 = v2 + i32(2)
				if uint32(t10) >= uint32(v2) {
					t11 := int32(load32(m.memory[int64(uint32(v1))+20:]))
					if t11 == i32(1139704) {
						goto l13
					}
					m.fn894(v3+i32(24), v1, v2)
					t12 := int32(load32(m.memory[int64(uint32(v3))+24:]))
					if t12 != i32(1) {
						goto l7
					}
					t13 := int32(load32(m.memory[int64(uint32(v3))+28:]))
					v1 = t13
					p14 := v5
					if uint32(v1) > uint32(v5) {
						p14 = v1
					}
					v5 = p14
					goto l13
				}
				goto l7
			}
			goto l7
		case 6, 7:
			v4 = i32(0)
			if uint32(v2) <= uint32(i32(-3)) {
				v5 = v2 + i32(3)
				t15 := v5
				v7 = v2 + i32(2)
				if uint32(t15) >= uint32(v7) {
					v6 = int64(uint32(int32(uint32(v5)>>1))) * i64(3)
					if int32(int64(uint64(v6)>>32)) == 0 {
						v5 = int32(v6) + i32(1)
						if v5 != 0 {
							t16 := int32(load32(m.memory[int64(uint32(v1))+20:]))
							v2 = t16
							if v2 == i32(1143836) {
								goto l13
							}
							if v2 == i32(1143808) {
								goto l13
							}
							m.fn894(v3+i32(32), v1, v7)
							t17 := int32(load32(m.memory[int64(uint32(v3))+32:]))
							if t17 != i32(1) {
								goto l7
							}
							t18 := int32(load32(m.memory[int64(uint32(v3))+36:]))
							v1 = t18
							p19 := v5
							if uint32(v1) > uint32(v5) {
								p19 = v1
							}
							v5 = p19
							goto l13
						}
						goto l7
					}
					goto l7
				}
				goto l7
			}
			goto l7
		case 8:
			if uint32(v2) <= uint32(i32(-3)) {
				m.fn894(v3+i32(40), v1, v2+i32(2))
				t20 := int32(load32(m.memory[int64(uint32(v3))+44:]))
				v5 = t20
				t21 := int32(load32(m.memory[int64(uint32(v3))+40:]))
				v4 = t21
				goto l7
			}
			v4 = i32(0)
			goto l7
		case 10:
			m.fn7(i32(1145880), i32(41), i32(1145972))
			panic("unreachable")
		}
	}
l13:
	v4 = i32(1)
l7:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
	store32(m.memory[uint32(v0):], uint32(v4))
	m.g0 = v3 + i32(48)
}
func (m *Module) fn891(v0, v1, v2 int32) {
	var v3, v4 int32
	var v5 int64
	var v6, v7 int32
	t0 := m.g0
	v3 = t0 - i32(48)
	m.g0 = v3
	{
		t1 := int32(m.memory[int64(uint32(v1))+24])
		switch t1 {
		default:
			v4 = i32(0)
			v5 = int64(uint32(v2)) * i64(3)
			if int32(int64(uint64(v5)>>32)) == 0 {
				v6 = int32(v5)
				if uint32(v6) <= uint32(i32(-4)) {
					v7 = int32(uint32(v2+i32(1))>>1)*i32(3) + i32(1)
					t4 := v7
					v6 = v6 + i32(3)
					p5 := v6
					if uint32(v7) > uint32(v6) {
						p5 = t4
					}
					v6 = p5
					t6 := int32(load32(m.memory[int64(uint32(v1))+20:]))
					v7 = t6
					if v7 == i32(1143808) {
						goto l12
					}
					if v7 == i32(1139704) {
						goto l12
					}
					if v7 == i32(1143836) {
						goto l12
					}
					m.fn895(v3+i32(8), v1, v2)
					t7 := int32(load32(m.memory[int64(uint32(v3))+8:]))
					if t7 != i32(1) {
						goto l7
					}
					t8 := int32(load32(m.memory[int64(uint32(v3))+12:]))
					v1 = t8
					p9 := v6
					if uint32(v1) > uint32(v6) {
						p9 = v1
					}
					v6 = p9
					goto l12
				}
				goto l7
			}
			goto l7
		case 1, 2, 3, 9:
			m.fn895(v3+i32(16), v1, v2)
			t2 := int32(load32(m.memory[int64(uint32(v3))+20:]))
			v6 = t2
			t3 := int32(load32(m.memory[int64(uint32(v3))+16:]))
			v4 = t3
			goto l7
		case 4, 5:
			v4 = i32(0)
			if uint32(v2) <= uint32(i32(-3)) {
				v7 = v2 + i32(2)
				v5 = int64(uint32(v7)) * i64(3)
				if int32(int64(uint64(v5)>>32)) == 0 {
					v2 = int32(v5)
					if uint32(v2) <= uint32(i32(-4)) {
						v6 = v2 + i32(3)
						t10 := int32(load32(m.memory[int64(uint32(v1))+20:]))
						if t10 == i32(1139704) {
							goto l12
						}
						m.fn895(v3+i32(24), v1, v7)
						t11 := int32(load32(m.memory[int64(uint32(v3))+24:]))
						if t11 != i32(1) {
							goto l7
						}
						t12 := int32(load32(m.memory[int64(uint32(v3))+28:]))
						v1 = t12
						p13 := v6
						if uint32(v1) > uint32(v6) {
							p13 = v1
						}
						v6 = p13
						goto l12
					}
					goto l7
				}
				goto l7
			}
			goto l7
		case 6, 7:
			v4 = i32(0)
			if uint32(v2) <= uint32(i32(-3)) {
				v6 = v2 + i32(3)
				t14 := v6
				v7 = v2 + i32(2)
				if uint32(t14) >= uint32(v7) {
					v5 = int64(uint32(int32(uint32(v6)>>1))) * i64(3)
					if int32(int64(uint64(v5)>>32)) == 0 {
						v6 = int32(v5) + i32(1)
						if v6 != 0 {
							t15 := int32(load32(m.memory[int64(uint32(v1))+20:]))
							v2 = t15
							if v2 == i32(1143836) {
								goto l12
							}
							if v2 == i32(1143808) {
								goto l12
							}
							m.fn895(v3+i32(32), v1, v7)
							t16 := int32(load32(m.memory[int64(uint32(v3))+32:]))
							if t16 != i32(1) {
								goto l7
							}
							t17 := int32(load32(m.memory[int64(uint32(v3))+36:]))
							v1 = t17
							p18 := v6
							if uint32(v1) > uint32(v6) {
								p18 = v1
							}
							v6 = p18
							goto l12
						}
						goto l7
					}
					goto l7
				}
				goto l7
			}
			goto l7
		case 8:
			if uint32(v2) <= uint32(i32(-3)) {
				m.fn895(v3+i32(40), v1, v2+i32(2))
				t19 := int32(load32(m.memory[int64(uint32(v3))+44:]))
				v6 = t19
				t20 := int32(load32(m.memory[int64(uint32(v3))+40:]))
				v4 = t20
				goto l7
			}
			v4 = i32(0)
			goto l7
		case 10:
			m.fn7(i32(1145880), i32(41), i32(1145924))
			panic("unreachable")
		}
	}
l12:
	v4 = i32(1)
l7:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v6))
	store32(m.memory[uint32(v0):], uint32(v4))
	m.g0 = v3 + i32(48)
}
func (m *Module) fn892(v0, v1, v2, v3, v4, v5 int32) {
	var v6, v7, v8 int32
	t0 := m.g0
	v6 = t0 - i32(16)
	m.g0 = v6
	t1 := int32(m.memory[int64(uint32(v1))+24])
	v7 = t1
	{
		{
			if v3 == 0 {
				v8 = i32(0)
				switch v7 {
				case 1:
					store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
					store32(m.memory[uint32(v0):], uint32(i32(0)))
					m.memory[int64(uint32(v0))+4] = byte(i32(0))
					goto l12
				case 2:
					store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
					store32(m.memory[uint32(v0):], uint32(i32(0)))
					m.memory[int64(uint32(v0))+4] = byte(i32(0))
					goto l12
				case 3:
					store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
					store32(m.memory[uint32(v0):], uint32(i32(0)))
					m.memory[int64(uint32(v0))+4] = byte(i32(0))
					goto l12
				case 4:
					goto l18
				case 5:
					goto l6
				case 6:
					goto l7
				case 7:
					goto l8
				case 9:
					goto l19
				case 10:
					goto l11
				default:
					store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
					store32(m.memory[uint32(v0):], uint32(i32(0)))
					m.memory[int64(uint32(v0))+4] = byte(i32(0))
					goto l12
				case 8:
					m.fn897(v0, v1, v2, i32(0), v4, v5, i32(0), i32(187))
					goto l12
				}
			}
			v8 = i32(0)
		l21:
			{
				{
					switch v7 & i32(255) {
					case 2:
						goto l3
					case 3:
						goto l4
					case 5:
						goto l6
					case 6:
						goto l7
					case 7:
						goto l8
					case 10:
						goto l11
					case 8:
						m.fn897(v0, v1, v2, v3, v4, v5, i32(0), i32(187))
						goto l12
					default:
						t2 := int32(m.memory[uint32(v2)])
						switch t2 + i32(-239) {
						case 0:
							goto l13
						case 15:
							goto l15
						case 16:
							goto l16
						default:
							goto l14
						}
					case 4:
						if uint32(v8) < uint32(v3) {
							t6 := int32(m.memory[uint32(v2+v8)])
							if t6 != i32(187) {
								m.fn897(v0, v1, v2, v3, v4, v5, v8, i32(239))
								goto l12
							}
							v7 = i32(5)
							goto l20
						}
						goto l18
					case 9:
						v8 = v3
						goto l19
					case 1:
						t3 := int32(m.memory[uint32(v2)])
						if t3 != i32(239) {
							goto l14
						}
					}
				l13:
					v7 = i32(4)
					goto l20
				l3:
					t4 := int32(m.memory[uint32(v2)])
					if t4 != i32(254) {
						goto l14
					}
				}
			l15:
				v7 = i32(6)
				goto l20
			l4:
				t5 := int32(m.memory[uint32(v2)])
				if t5 == i32(255) {
					goto l16
				}
			}
		l14:
			v7 = i32(9)
			m.memory[int64(uint32(v1))+24] = byte(i32(9))
			goto l21
		l16:
			v7 = i32(7)
		l20:
			m.memory[int64(uint32(v1))+24] = byte(v7)
			v8 = v8 + i32(1)
			goto l21
		l18:
			m.fn897(v0, v1, v2, v3, v4, v5, v8, i32(239))
			goto l12
		l19:
			m.fn898(v0, v1, v2, v8, v4, v5, i32(1))
			t7 := int32(m.memory[int64(uint32(v0))+4])
			if t7 != 0 {
				goto l12
			}
			m.memory[int64(uint32(v1))+24] = byte(i32(10))
			goto l12
		}
	l11:
		m.fn7(i32(1145880), i32(41), i32(1145940))
		panic("unreachable")
	l6:
		if uint32(v8) < uint32(v3) {
			goto l28
		}
		m.fn899(v0, v1, v2, v3, v4, v5, v8)
		goto l12
	l7:
		if uint32(v8) < uint32(v3) {
			{
				t15 := int32(m.memory[uint32(v2+v8)])
				if t15 != i32(255) {
					m.fn897(v0, v1, v2, v3, v4, v5, v8, i32(254))
					goto l12
				}
				m.memory[int64(uint32(v1))+24] = byte(i32(9))
				v7 = v8 + i32(1)
				t16 := int32(load32(m.memory[int64(uint32(v1))+20:]))
				if t16 == i32(1143808) {
					goto l35
				}
				store32(m.memory[int64(uint32(v1))+4:], uint32(i32(65536)))
				m.memory[int64(uint32(v1))+2] = byte(i32(0))
				m.memory[uint32(v1)] = byte(i32(10))
				store32(m.memory[int64(uint32(v1))+20:], uint32(i32(1143808)))
				goto l35
			}
		l35:
			m.fn898(v6+i32(4), v1, v2+v7, v3-v7, v4, v5, i32(1))
			v8 = v6 + i32(8)
			{
				t17 := int32(m.memory[int64(uint32(v6))+8])
				if t17 != 0 {
					goto l36
				}
				m.memory[int64(uint32(v1))+24] = byte(i32(10))
			}
		l36:
			t18 := int32(m.memory[int64(uint32(v8))+2])
			m.memory[int64(uint32(v0))+6] = byte(t18)
			t19 := int32(load16(m.memory[uint32(v8):]))
			store16(m.memory[int64(uint32(v0))+4:], uint16(t19))
			t20 := int32(load32(m.memory[int64(uint32(v6))+4:]))
			v1 = t20
			t21 := int32(load32(m.memory[int64(uint32(v6))+12:]))
			store32(m.memory[int64(uint32(v0))+8:], uint32(t21))
			store32(m.memory[uint32(v0):], uint32(v1+v7))
			goto l12
		}
		m.fn897(v0, v1, v2, v3, v4, v5, v8, i32(254))
		goto l12
	l8:
		if uint32(v8) < uint32(v3) {
			{
				t8 := int32(m.memory[uint32(v2+v8)])
				if t8 != i32(254) {
					m.fn897(v0, v1, v2, v3, v4, v5, v8, i32(255))
					goto l12
				}
				m.memory[int64(uint32(v1))+24] = byte(i32(9))
				v7 = v8 + i32(1)
				t9 := int32(load32(m.memory[int64(uint32(v1))+20:]))
				if t9 == i32(1143836) {
					goto l32
				}
				store32(m.memory[int64(uint32(v1))+4:], uint32(i32(0)))
				m.memory[int64(uint32(v1))+2] = byte(i32(0))
				m.memory[uint32(v1)] = byte(i32(10))
				store32(m.memory[int64(uint32(v1))+20:], uint32(i32(1143836)))
				goto l32
			}
		l32:
			m.fn898(v6+i32(4), v1, v2+v7, v3-v7, v4, v5, i32(1))
			v8 = v6 + i32(8)
			{
				t10 := int32(m.memory[int64(uint32(v6))+8])
				if t10 != 0 {
					goto l33
				}
				m.memory[int64(uint32(v1))+24] = byte(i32(10))
			}
		l33:
			t11 := int32(m.memory[int64(uint32(v8))+2])
			m.memory[int64(uint32(v0))+6] = byte(t11)
			t12 := int32(load16(m.memory[uint32(v8):]))
			store16(m.memory[int64(uint32(v0))+4:], uint16(t12))
			t13 := int32(load32(m.memory[int64(uint32(v6))+4:]))
			v1 = t13
			t14 := int32(load32(m.memory[int64(uint32(v6))+12:]))
			store32(m.memory[int64(uint32(v0))+8:], uint32(t14))
			store32(m.memory[uint32(v0):], uint32(v1+v7))
			goto l12
		}
		m.fn897(v0, v1, v2, v3, v4, v5, v8, i32(255))
		goto l12
	l28:
		{
			t22 := int32(m.memory[uint32(v2+v8)])
			if t22 != i32(191) {
				m.fn899(v0, v1, v2, v3, v4, v5, v8)
				goto l12
			}
			m.memory[int64(uint32(v1))+24] = byte(i32(9))
			v7 = v8 + i32(1)
			t23 := int32(load32(m.memory[int64(uint32(v1))+20:]))
			if t23 == i32(1139704) {
				goto l38
			}
			store16(m.memory[int64(uint32(v1))+16:], uint16(i32(49024)))
			store32(m.memory[int64(uint32(v1))+12:], uint32(i32(0)))
			store64(m.memory[int64(uint32(v1))+4:], uint64(i64(0)))
			m.memory[uint32(v1)] = byte(i32(1))
			store32(m.memory[int64(uint32(v1))+20:], uint32(i32(1139704)))
			goto l38
		}
	l38:
		m.fn898(v6+i32(4), v1, v2+v7, v3-v7, v4, v5, i32(1))
		v8 = v6 + i32(8)
		{
			t24 := int32(m.memory[int64(uint32(v6))+8])
			if t24 != 0 {
				goto l39
			}
			m.memory[int64(uint32(v1))+24] = byte(i32(10))
		}
	l39:
		t25 := int32(m.memory[int64(uint32(v8))+2])
		m.memory[int64(uint32(v0))+6] = byte(t25)
		t26 := int32(load16(m.memory[uint32(v8):]))
		store16(m.memory[int64(uint32(v0))+4:], uint16(t26))
		t27 := int32(load32(m.memory[int64(uint32(v6))+4:]))
		v1 = t27
		t28 := int32(load32(m.memory[int64(uint32(v6))+12:]))
		store32(m.memory[int64(uint32(v0))+8:], uint32(t28))
		store32(m.memory[uint32(v0):], uint32(v1+v7))
	}
l12:
	m.g0 = v6 + i32(16)
}
func (m *Module) fn893(v0, v1, v2 int32) {
	var v3 int32
	t0 := m.g0
	v3 = t0 - i32(16)
	m.g0 = v3
	v1 = v2 + v1
	if uint32(v1) >= uint32(v2) {
		goto l0
	}
	m.fn16(i32(0), i32(0))
	panic("unreachable")
l0:
	t1 := int32(load32(m.memory[uint32(v0):]))
	t2 := v3 + i32(4)
	v2 = t1
	t3 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t4 := v2
	t5 := v1
	v2 = v2 << 1
	p6 := v2
	if uint32(v1) > uint32(v2) {
		p6 = t5
	}
	v2 = p6
	p7 := i32(8)
	if uint32(v2) > uint32(i32(8)) {
		p7 = v2
	}
	v2 = p7
	m.fn896(t2, t4, t3, v2)
	{
		t8 := int32(load32(m.memory[int64(uint32(v3))+4:]))
		if t8 != i32(1) {
			goto l1
		}
		t9 := int32(load32(m.memory[int64(uint32(v3))+8:]))
		t10 := int32(load32(m.memory[int64(uint32(v3))+12:]))
		m.fn16(t9, t10)
		panic("unreachable")
	}
l1:
	t11 := int32(load32(m.memory[int64(uint32(v3))+8:]))
	v1 = t11
	store32(m.memory[uint32(v0):], uint32(v2))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
	m.g0 = v3 + i32(16)
}
func (m *Module) fn894(v0, v1, v2 int32) {
	var v3, v4 int32
	var v5 int64
	var v6 int32
	v3 = i32(1)
	v4 = i32(3)
	{
		t0 := int32(m.memory[uint32(v1)])
		switch t0 {
		case 8:
			goto l8
		default:
			v5 = int64(uint32(v2)) * i64(3)
			v4 = int32(v5)
			var p1 int32
			if int32(int64(uint64(v5)>>32)) == 0 {
				p1 = 1
			}
			v3 = p1
			goto l8
		case 2:
			v3 = i32(0)
			{
				t2 := int32(m.memory[int64(uint32(v1))+1])
				t3 := int32(m.memory[int64(uint32(v1))+9])
				t4 := int32(m.memory[int64(uint32(v1))+3])
				t5 := int32(m.memory[int64(uint32(v1))+5])
				t6 := int32(m.memory[int64(uint32(v1))+7])
				v1 = v2 + (t2+t3+t4+t5+t6)&i32(63)
				if uint32(v1) >= uint32(v2) {
					v5 = int64(uint32(v1)) * i64(3)
					if int32(int64(uint64(v5)>>32)) == 0 {
						v1 = int32(v5)
						v4 = v1 + i32(1)
						var p7 int32
						if v1 != i32(-1) {
							p7 = 1
						}
						v3 = p7
						goto l8
					}
					goto l8
				}
				goto l8
			}
		case 3:
			t8 := int32(m.memory[int64(uint32(v1))+1])
			v1 = v2 + t8&i32(255)
			var p9 int32
			if uint32(v1) >= uint32(v2) {
				p9 = 1
			}
			var p10 int32
			if uint32(v1) < uint32(i32(0x7fffffff)) {
				p10 = 1
			}
			v3 = p9 & p10
			v4 = v1<<1 + i32(2)
			goto l8
		case 4:
			v3 = i32(0)
			{
				t11 := int32(m.memory[int64(uint32(v1))+1])
				t12 := v2
				var p13 int32
				if t11 != i32(0) {
					p13 = 1
				}
				v1 = t12 + p13
				if uint32(v1) >= uint32(v2) {
					if v1 == i32(-1) {
						goto l8
					}
					v3 = int32(uint32(v1+i32(1)) >> 1)
					v1 = v3 + v1
					var p14 int32
					if uint32(v1) >= uint32(v3) {
						p14 = 1
					}
					var p15 int32
					if uint32(v1) < uint32(i32(-2)) {
						p15 = 1
					}
					v3 = p14 & p15
					v4 = v1 + i32(2)
					goto l8
				}
				goto l8
			}
		case 5:
			{
				{
					t16 := int32(m.memory[int64(uint32(v1))+5])
					if t16 != 0 {
						goto l14
					}
					t17 := int32(m.memory[int64(uint32(v1))+1])
					v6 = t17
					v4 = i32(0)
					goto l15
				}
			l14:
				t18 := int32(m.memory[int64(uint32(v1))+1])
				t19 := int32(m.memory[int64(uint32(v1))+2])
				v3 = t19
				v6 = t18 + v3
				v4 = (v3 ^ i32(-1)) & i32(1)
			}
		l15:
			v3 = i32(0)
			{
				t20 := int32(m.memory[int64(uint32(v1))+3])
				t21 := v4 + v2
				var p22 int32
				if uint32(t20) > uint32(i32(4)) {
					p22 = 1
				}
				v1 = t21 + p22
				if uint32(v1) >= uint32(v2) {
					v4 = v1 + v6&i32(255)
					if uint32(v4) >= uint32(v1) {
						v5 = int64(uint32(v4)) * i64(3)
						v4 = int32(v5)
						var p23 int32
						if int32(int64(uint64(v5)>>32)) == 0 {
							p23 = 1
						}
						v3 = p23
						goto l8
					}
					goto l8
				}
				goto l8
			}
		case 6:
			t24 := int32(m.memory[int64(uint32(v1))+1])
			v1 = v2 + t24
			if uint32(v1) >= uint32(v2) {
				v5 = int64(uint32(v1)) * i64(3)
				v4 = int32(v5)
				var p38 int32
				if int32(int64(uint64(v5)>>32)) == 0 {
					p38 = 1
				}
				v3 = p38
				goto l8
			}
			goto l19
		case 7:
			v3 = i32(0)
			{
				t25 := int32(m.memory[int64(uint32(v1))+1])
				v1 = v2 + t25
				if uint32(v1) >= uint32(v2) {
					if v1 == i32(-1) {
						goto l8
					}
					v3 = int32(uint32(v1+i32(1)) >> 1)
					v1 = v3 + v1
					var p26 int32
					if uint32(v1) >= uint32(v3) {
						p26 = 1
					}
					var p27 int32
					if uint32(v1) < uint32(i32(-2)) {
						p27 = 1
					}
					v3 = p26 & p27
					v4 = v1 + i32(2)
					goto l8
				}
				goto l8
			}
		case 9:
			v5 = int64(uint32(v2)) * i64(3)
			v4 = int32(v5)
			var p28 int32
			if int32(int64(uint64(v5)>>32)) == 0 {
				p28 = 1
			}
			v3 = p28
			goto l8
		case 10:
			v3 = i32(0)
			{
				t29 := int32(m.memory[int64(uint32(v1))+2])
				p30 := i32(1)
				if t29 != 0 {
					p30 = i32(2)
				}
				t31 := int32(load16(m.memory[int64(uint32(v1))+4:]))
				t33 := p30 + v2
				p32 := i32(0)
				if t31 != 0 {
					p32 = i32(2)
				}
				v1 = t33 + p32
				if uint32(v1) >= uint32(v2) {
					v5 = int64(uint32(int32(uint32(v1)>>1))) * i64(3)
					if int32(int64(uint64(v5)>>32)) == 0 {
						v1 = int32(v5)
						v4 = v1 + i32(1)
						var p34 int32
						if v1 != i32(-1) {
							p34 = 1
						}
						v3 = p34
						goto l8
					}
					goto l8
				}
				goto l8
			}
		case 1:
			t35 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			t36 := int32(load32(m.memory[int64(uint32(v1))+12:]))
			p37 := i32(3)
			if t36 != 0 {
				p37 = t35 + i32(4)
			}
			v1 = p37
			v4 = v1 + v2
			if uint32(v4) < uint32(v1) {
				goto l19
			}
			goto l8
		}
	}
l19:
	v3 = i32(0)
l8:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v4))
	store32(m.memory[uint32(v0):], uint32(v3))
}
func (m *Module) fn895(v0, v1, v2 int32) {
	var v3, v4 int32
	var v5 int64
	var v6 int32
	v3 = i32(1)
	v4 = i32(3)
	{
		t0 := int32(m.memory[uint32(v1)])
		switch t0 {
		case 8:
			goto l8
		default:
			v5 = int64(uint32(v2)) * i64(3)
			v4 = int32(v5)
			var p1 int32
			if int32(int64(uint64(v5)>>32)) == 0 {
				p1 = 1
			}
			v3 = p1
			goto l8
		case 1:
			v3 = i32(0)
			{
				t2 := int32(load32(m.memory[int64(uint32(v1))+8:]))
				t3 := int32(load32(m.memory[int64(uint32(v1))+12:]))
				p4 := i32(0)
				if t3 != 0 {
					p4 = t2 + i32(1)
				}
				v1 = p4
				v4 = v1 + v2
				if uint32(v4) >= uint32(v1) {
					v5 = int64(uint32(v4)) * i64(3)
					if int32(int64(uint64(v5)>>32)) == 0 {
						v1 = int32(v5)
						v4 = v1 + i32(3)
						var p5 int32
						if uint32(v1) < uint32(i32(-3)) {
							p5 = 1
						}
						v3 = p5
						goto l8
					}
					goto l8
				}
				goto l8
			}
		case 2:
			v3 = i32(0)
			{
				t6 := int32(m.memory[int64(uint32(v1))+1])
				t7 := int32(m.memory[int64(uint32(v1))+9])
				t8 := int32(m.memory[int64(uint32(v1))+3])
				t9 := int32(m.memory[int64(uint32(v1))+5])
				t10 := int32(m.memory[int64(uint32(v1))+7])
				v1 = v2 + (t6+t7+t8+t9+t10)&i32(63)
				if uint32(v1) >= uint32(v2) {
					v5 = int64(uint32(v1)) * i64(3)
					if int32(int64(uint64(v5)>>32)) == 0 {
						v1 = int32(v5)
						v4 = v1 + i32(1)
						var p11 int32
						if v1 != i32(-1) {
							p11 = 1
						}
						v3 = p11
						goto l8
					}
					goto l8
				}
				goto l8
			}
		case 3:
			v3 = i32(0)
			{
				t12 := int32(m.memory[int64(uint32(v1))+1])
				v1 = v2 + t12
				if uint32(v1) >= uint32(v2) {
					v5 = int64(uint32(v1)) * i64(3)
					if int32(int64(uint64(v5)>>32)) == 0 {
						v1 = int32(v5)
						v4 = v1 + i32(3)
						var p13 int32
						if uint32(v1) < uint32(i32(-3)) {
							p13 = 1
						}
						v3 = p13
						goto l8
					}
					goto l8
				}
				goto l8
			}
		case 4:
			v3 = i32(0)
			{
				t14 := int32(m.memory[int64(uint32(v1))+1])
				t15 := v2
				var p16 int32
				if t14 != i32(0) {
					p16 = 1
				}
				v1 = t15 + p16
				if uint32(v1) >= uint32(v2) {
					v5 = int64(uint32(v1)) * i64(3)
					v4 = int32(v5)
					var p17 int32
					if int32(int64(uint64(v5)>>32)) == 0 {
						p17 = 1
					}
					v3 = p17
					goto l8
				}
				goto l8
			}
		case 5:
			{
				{
					t18 := int32(m.memory[int64(uint32(v1))+5])
					if t18 != 0 {
						goto l18
					}
					t19 := int32(m.memory[int64(uint32(v1))+1])
					v6 = t19
					v4 = i32(0)
					goto l19
				}
			l18:
				t20 := int32(m.memory[int64(uint32(v1))+1])
				t21 := int32(m.memory[int64(uint32(v1))+2])
				v3 = t21
				v6 = t20 + v3
				v4 = (v3 ^ i32(-1)) & i32(1)
			}
		l19:
			v3 = i32(0)
			{
				t22 := int32(m.memory[int64(uint32(v1))+3])
				t23 := v4 + v2
				var p24 int32
				if uint32(t22) > uint32(i32(4)) {
					p24 = 1
				}
				v1 = t23 + p24
				if uint32(v1) >= uint32(v2) {
					v4 = v1 + v6&i32(255)
					if uint32(v4) >= uint32(v1) {
						v5 = int64(uint32(v4)) * i64(3)
						v4 = int32(v5)
						var p25 int32
						if int32(int64(uint64(v5)>>32)) == 0 {
							p25 = 1
						}
						v3 = p25
						goto l8
					}
					goto l8
				}
				goto l8
			}
		case 6:
			t26 := int32(m.memory[int64(uint32(v1))+1])
			v1 = v2 + t26
			if uint32(v1) >= uint32(v2) {
				v5 = int64(uint32(v1)) * i64(3)
				v4 = int32(v5)
				var p36 int32
				if int32(int64(uint64(v5)>>32)) == 0 {
					p36 = 1
				}
				v3 = p36
				goto l8
			}
			goto l23
		case 7:
			t27 := int32(m.memory[int64(uint32(v1))+1])
			v1 = v2 + t27
			if uint32(v1) < uint32(v2) {
				goto l23
			}
			v5 = int64(uint32(v1)) * i64(3)
			v4 = int32(v5)
			var p28 int32
			if int32(int64(uint64(v5)>>32)) == 0 {
				p28 = 1
			}
			v3 = p28
			goto l8
		case 9:
			v5 = int64(uint32(v2)) * i64(3)
			v4 = int32(v5)
			var p29 int32
			if int32(int64(uint64(v5)>>32)) == 0 {
				p29 = 1
			}
			v3 = p29
			goto l8
		case 10:
			v3 = i32(0)
			{
				t30 := int32(m.memory[int64(uint32(v1))+2])
				p31 := i32(1)
				if t30 != 0 {
					p31 = i32(2)
				}
				t32 := int32(load16(m.memory[int64(uint32(v1))+4:]))
				t34 := p31 + v2
				p33 := i32(0)
				if t32 != 0 {
					p33 = i32(2)
				}
				v1 = t34 + p33
				if uint32(v1) >= uint32(v2) {
					v5 = int64(uint32(int32(uint32(v1)>>1))) * i64(3)
					if int32(int64(uint64(v5)>>32)) == 0 {
						v1 = int32(v5)
						v4 = v1 + i32(1)
						var p35 int32
						if v1 != i32(-1) {
							p35 = 1
						}
						v3 = p35
						goto l8
					}
					goto l8
				}
				goto l8
			}
		}
	}
l23:
	v3 = i32(0)
l8:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v4))
	store32(m.memory[uint32(v0):], uint32(v3))
}
func (m *Module) fn896(v0, v1, v2, v3 int32) {
	var v4 int32
	v4 = i32(0)
	if v3 >= i32(0) {
		goto l0
	}
	v1 = i32(1)
	v2 = i32(4)
	goto l1
l0:
	{
		{
			if v1 == 0 {
				goto l2
			}
			t0 := m.fn23(v2, v1, i32(1), v3)
			v4 = t0
			goto l3
		}
	l2:
		t1 := m.fn11(v3)
		v4 = t1
	}
l3:
	if v4 != 0 {
		goto l4
	}
	v1 = i32(1)
	store32(m.memory[int64(uint32(v0))+4:], uint32(i32(1)))
	goto l5
l4:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v4))
	v1 = i32(0)
l5:
	v2 = i32(8)
	v4 = v3
l1:
	store32(m.memory[uint32(v0+v2):], uint32(v4))
	store32(m.memory[uint32(v0):], uint32(v1))
}
