package core

import (
	"math"
	"math/bits"
)

func (m *Module) fn852(v0, v1 int32) int32 {
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
					t5 := int32(m.memory[int64(uint32(int32(v4)&i32(15)))+1098816])
					m.memory[uint32(v2+v0+i32(-2))] = byte(t5)
					v0 = v0 + i32(-1)
					v4 = int64(uint64(v4) >> 4)
					if v4 != i64(0) {
						goto l3
					}
				}
				t6 := m.fn306(v1, i32(1), i32(1122550), i32(2), v2+v0+i32(-1), i32(17)-v0)
				v0 = t6
				goto l2
			}
			if v3&i32(0x4000000) != 0 {
				goto l1
			}
			t3 := m.fn162(v0, v1)
			v0 = t3
			goto l2
		}
	l1:
		t7 := int64(load64(m.memory[uint32(v0):]))
		v4 = t7
		v0 = i32(17)
	l4:
		{
			t8 := int32(m.memory[int64(uint32(int32(v4)&i32(15)))+1122552])
			m.memory[uint32(v2+v0+i32(-2))] = byte(t8)
			v0 = v0 + i32(-1)
			v4 = int64(uint64(v4) >> 4)
			if v4 != i64(0) {
				goto l4
			}
		}
		t9 := m.fn306(v1, i32(1), i32(1122550), i32(2), v2+v0+i32(-1), i32(17)-v0)
		v0 = t9
	}
l2:
	m.g0 = v2 + i32(16)
	return v0
}
func (m *Module) fn853(v0, v1, v2 int32) {
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
				t1 := m.fn854(v1)
				if t1 == 0 {
					goto l13
				}
				m.memory[int64(uint32(v3))+14] = byte(i32(0))
				store16(m.memory[int64(uint32(v3))+12:], uint16(i32(0)))
				t2 := int32(m.memory[int64(uint32(int32(uint32(v1)>>20)))+1098816])
				m.memory[int64(uint32(v3))+15] = byte(t2)
				t3 := int32(m.memory[int64(uint32(int32(uint32(v1)>>4)&i32(15)))+1098816])
				m.memory[int64(uint32(v3))+19] = byte(t3)
				t4 := int32(m.memory[int64(uint32(int32(uint32(v1)>>8)&i32(15)))+1098816])
				m.memory[int64(uint32(v3))+18] = byte(t4)
				t5 := int32(m.memory[int64(uint32(int32(uint32(v1)>>12)&i32(15)))+1098816])
				m.memory[int64(uint32(v3))+17] = byte(t5)
				t6 := int32(m.memory[int64(uint32(int32(uint32(v1)>>16)&i32(15)))+1098816])
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
				t10 := int32(m.memory[int64(uint32(v1&i32(15)))+1098816])
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
				t12 := int32(m.memory[int64(uint32(v2))+1101501])
				t13 := v5
				v4 = t12
				v8 = t13 + v4
				{
					t14 := int32(m.memory[int64(uint32(v2))+1101500])
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
						v2 = v5 + i32(1101576)
						goto l25
					l22:
						m.fn121(v5, v8, i32(284), i32(1102152))
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
				t16 := int32(m.memory[int64(uint32(v2))+1100693])
				t17 := v5
				v4 = t16
				v8 = t17 + v4
				{
					t18 := int32(m.memory[int64(uint32(v2))+1100692])
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
				v2 = v5 + i32(1100784)
				goto l32
			l29:
				m.fn121(v5, v8, i32(212), i32(1102152))
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
				t20 := int32(int8(m.memory[int64(uint32(v2))+1100996]))
				v8 = t20
				if v8 < i32(0) {
					if v7 == i32(504) {
						m.fn219(i32(1102168))
						panic("unreachable")
					}
					t21 := int32(m.memory[uint32(v2+i32(1100997))])
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
			t22 := int32(int8(m.memory[int64(uint32(v2))+1101860]))
			v8 = t22
			if v8 < i32(0) {
				if v7 == i32(292) {
					m.fn219(i32(1102168))
					panic("unreachable")
				}
				t23 := int32(m.memory[uint32(v2+i32(1101861))])
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
		t24 := int32(m.memory[int64(uint32(int32(uint32(v1)>>20)))+1098816])
		m.memory[int64(uint32(v3))+25] = byte(t24)
		t25 := int32(m.memory[int64(uint32(int32(uint32(v1)>>4)&i32(15)))+1098816])
		m.memory[int64(uint32(v3))+29] = byte(t25)
		t26 := int32(m.memory[int64(uint32(int32(uint32(v1)>>8)&i32(15)))+1098816])
		m.memory[int64(uint32(v3))+28] = byte(t26)
		t27 := int32(m.memory[int64(uint32(int32(uint32(v1)>>12)&i32(15)))+1098816])
		m.memory[int64(uint32(v3))+27] = byte(t27)
		t28 := int32(m.memory[int64(uint32(int32(uint32(v1)>>16)&i32(15)))+1098816])
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
		t32 := int32(m.memory[int64(uint32(v1&i32(15)))+1098816])
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
func (m *Module) fn854(v0 int32) int32 {
	var v1, v2, v3, v4, v5 int32
	v1 = i32(0)
	p0 := i32(16)
	if uint32(v0) < uint32(i32(69291)) {
		p0 = i32(0)
	}
	v2 = p0
	t1 := v2
	v2 = v2 | i32(8)
	t2 := int32(load32(m.memory[int64(uint32(v2<<2))+1105868:]))
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
	t7 := int32(load32(m.memory[int64(uint32(v3<<2))+1105868:]))
	p8 := v3
	if uint32(t7<<11) > uint32(v2) {
		p8 = t6
	}
	v3 = p8
	t9 := v3
	v3 = v3 | i32(2)
	t10 := int32(load32(m.memory[int64(uint32(v3<<2))+1105868:]))
	p11 := v3
	if uint32(t10<<11) > uint32(v2) {
		p11 = t9
	}
	v3 = p11
	t12 := v3
	v3 = v3 + i32(1)
	t13 := int32(load32(m.memory[int64(uint32(v3<<2))+1105868:]))
	p14 := v3
	if uint32(t13<<11) > uint32(v2) {
		p14 = t12
	}
	v3 = p14
	t15 := v3
	v3 = v3 + i32(1)
	t16 := int32(load32(m.memory[int64(uint32(v3<<2))+1105868:]))
	p17 := v3
	if uint32(t16<<11) > uint32(v2) {
		p17 = t15
	}
	v3 = p17
	t18 := int32(load32(m.memory[int64(uint32(v3<<2))+1105868:]))
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
	v5 = v2 + i32(1105868)
	t21 := int32(load32(m.memory[int64(uint32(v2))+1105868:]))
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
		t24 := int32(m.memory[uint32(v2+i32(1097358))])
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
func (m *Module) fn855(v0, v1, v2 int32) int32 {
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
					m.fn33(v11, i32(40), i32(1099872))
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
					m.fn33(v12, i32(40), i32(1099872))
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
			m.fn121(i32(0), v5, i32(40), i32(1099872))
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
				m.fn33(v1, i32(40), i32(1099872))
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
				m.fn33(v12, i32(40), i32(1099872))
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
func (m *Module) fn856(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7, v8, v9, v10 int32
	if uint32(v1) >= uint32(i32(1280)) {
		m.fn3(i32(1099888), i32(29), i32(1099872))
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
					m.fn33(v4, i32(40), i32(1099872))
					panic("unreachable")
				}
				v7 = v2 + v4
				if uint32(v7) >= uint32(i32(40)) {
					m.fn33(v7, i32(40), i32(1099872))
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
				m.fn33(v6, i32(40), i32(1099872))
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
				m.fn33(v4, i32(40), i32(1099872))
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
func (m *Module) fn857(v0, v1 int32) {
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
		t1 := int32(load16(m.memory[int64(uint32(v1))+1107720:]))
		v4 = t1
		v5 = v4 & i32(2047)
		t2 := int32(load16(m.memory[int64(uint32(v1))+1107722:]))
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
					m.fn33(i32(768), i32(768), i32(1109160))
					panic("unreachable")
				}
				v10 = v8 + v1
				v1 = v1 + i32(1)
				t3 := int32(m.memory[uint32(v10+i32(1308))])
				v10 = t3
				t4 := int32(m.memory[uint32(v9+i32(1109158))])
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
				m.fn33(v2, i32(768), i32(1099960))
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
				m.fn33(v1, i32(768), i32(1099944))
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
func (m *Module) fn858(v0, v1 int32) {
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
				m.fn33(i32(768), i32(768), i32(1099976))
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
					m.fn33(v8+v7, i32(768), i32(1099992))
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
					m.fn33(v7, i32(768), i32(1099944))
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
func (m *Module) fn859(v0 int32) int64 {
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
func (m *Module) fn860(v0, v1, v2, v3, v4 int32) {
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
				v7 = int64(uint32(i32(2))) << 32
				store64(m.memory[int64(uint32(t5))+32:], uint64(v7|int64(uint32(v5+i32(8)))))
				store64(m.memory[int64(uint32(v5))+24:], uint64(v7|int64(uint32(v5))))
				m.fn28(i32(1050346), v5+i32(24), v4)
				panic("unreachable")
			}
			if uint32(v3) > uint32(v1) {
				t6 := v5
				v7 = int64(uint32(i32(2))) << 32
				store64(m.memory[int64(uint32(t6))+32:], uint64(v7|int64(uint32(v5+i32(8)))))
				store64(m.memory[int64(uint32(v5))+24:], uint64(v7|int64(uint32(v5+i32(4)))))
				m.fn28(i32(1050407), v5+i32(24), v4)
				panic("unreachable")
			}
			if uint32(v2) > uint32(v3) {
				t7 := v5
				v7 = int64(uint32(i32(2))) << 32
				store64(m.memory[int64(uint32(t7))+32:], uint64(v7|int64(uint32(v5+i32(4)))))
				store64(m.memory[int64(uint32(v5))+24:], uint64(v7|int64(uint32(v5))))
				m.fn28(i32(1049763), v5+i32(24), v4)
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
		store64(m.memory[int64(uint32(v5))+40:], uint64(int64(uint32(i32(90)))<<32|int64(uint32(v5+i32(12)))))
		store64(m.memory[int64(uint32(v5))+32:], uint64(int64(uint32(i32(91)))<<32|int64(uint32(v5+i32(20)))))
		store64(m.memory[int64(uint32(v5))+24:], uint64(int64(uint32(i32(2)))<<32|int64(uint32(v5))))
		m.fn28(i32(1066107), v5+i32(24), v4)
		panic("unreachable")
	l8:
		m.fn38(v0, v1, v6, v2, v4)
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
		store64(m.memory[int64(uint32(v5))+40:], uint64(int64(uint32(i32(90)))<<32|int64(uint32(v5+i32(12)))))
		store64(m.memory[int64(uint32(v5))+32:], uint64(int64(uint32(i32(91)))<<32|int64(uint32(v5+i32(20)))))
		store64(m.memory[int64(uint32(v5))+24:], uint64(int64(uint32(i32(2)))<<32|int64(uint32(v5+i32(4)))))
		m.fn28(i32(1066189), v5+i32(24), v4)
		panic("unreachable")
	}
l13:
	m.fn219(v4)
	panic("unreachable")
l23:
	m.fn38(v0, v1, v6, v3, v4)
	panic("unreachable")
l18:
	t24 := v5
	v7 = int64(uint32(i32(2))) << 32
	store64(m.memory[int64(uint32(t24))+32:], uint64(v7|int64(uint32(v5+i32(8)))))
	store64(m.memory[int64(uint32(v5))+24:], uint64(v7|int64(uint32(v5+i32(4)))))
	m.fn28(i32(1050407), v5+i32(24), v4)
	panic("unreachable")
}
func (m *Module) fn861(v0, v1 int32) int32 {
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
						t4 := int32(m.memory[int64(uint32(v4&i32(15)))+1098816])
						m.memory[uint32(v2+i32(8)+v3+i32(-2))] = byte(t4)
						v3 = v3 + i32(-1)
						v4 = int32(uint32(v4) >> 4)
						if v4 != 0 {
							goto l4
						}
					}
					v4 = i32(1)
					t5 := m.fn306(v1, i32(1), i32(1122550), i32(2), v2+i32(8)+v3+i32(-1), i32(9)-v3)
					if t5 == 0 {
						goto l2
					}
					goto l3
				}
				if v3&i32(0x4000000) != 0 {
					goto l1
				}
				t2 := m.fn14(v0, v1)
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
				t7 := int32(m.memory[int64(uint32(v4&i32(15)))+1122552])
				m.memory[uint32(v2+i32(8)+v3+i32(-2))] = byte(t7)
				v3 = v3 + i32(-1)
				v4 = int32(uint32(v4) >> 4)
				if v4 != 0 {
					goto l5
				}
			}
			v4 = i32(1)
			t8 := m.fn306(v1, i32(1), i32(1122550), i32(2), v2+i32(8)+v3+i32(-1), i32(9)-v3)
			if t8 != 0 {
				goto l3
			}
		}
	l2:
		{
			t9 := int32(load32(m.memory[uint32(v1):]))
			t10 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t11 := int32(load32(m.memory[int64(uint32(t10))+12:]))
			t12 := m.t0[uint(t11)].(func(int32, int32, int32) int32)(t9, i32(1273647), i32(2))
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
					t16 := int32(m.memory[int64(uint32(v4&i32(15)))+1098816])
					m.memory[uint32(v2+i32(8)+v3+i32(-2))] = byte(t16)
					v3 = v3 + i32(-1)
					v4 = int32(uint32(v4) >> 4)
					if v4 != 0 {
						goto l9
					}
				}
				t17 := m.fn306(v1, i32(1), i32(1122550), i32(2), v2+i32(8)+v3+i32(-1), i32(9)-v3)
				v4 = t17
				goto l3
			}
			if v4&i32(0x4000000) != 0 {
				goto l8
			}
			t14 := m.fn14(v3, v1)
			v4 = t14
			goto l3
		}
	l8:
		t18 := int32(load32(m.memory[uint32(v3):]))
		v4 = t18
		v3 = i32(9)
	l10:
		{
			t19 := int32(m.memory[int64(uint32(v4&i32(15)))+1122552])
			m.memory[uint32(v2+i32(8)+v3+i32(-2))] = byte(t19)
			v3 = v3 + i32(-1)
			v4 = int32(uint32(v4) >> 4)
			if v4 != 0 {
				goto l10
			}
		}
		t20 := m.fn306(v1, i32(1), i32(1122550), i32(2), v2+i32(8)+v3+i32(-1), i32(9)-v3)
		v4 = t20
	}
l3:
	m.g0 = v2 + i32(16)
	return v4
}
func (m *Module) fn862(v0, v1 int32) int32 {
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
		m.fn853(v2, t6, i32(257))
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
func (m *Module) fn863(v0, v1 int32) int32 {
	t0 := m.fn56(v1, i32(1122692), i32(24))
	return t0
}
func (m *Module) fn864(v0, v1 int32) int32 {
	t0 := m.fn56(v1, i32(1122660), i32(32))
	return t0
}
func (m *Module) fn865(v0, v1, v2, v3, v4, v5 int32) {
	var v6 int32
	{
		if v2 == 0 {
			m.fn3(i32(1102209), i32(33), i32(1102244))
			panic("unreachable")
		}
		t0 := int32(m.memory[uint32(v1)])
		if uint32(t0) <= uint32(i32(48)) {
			m.fn3(i32(1102260), i32(31), i32(1102292))
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
					store32(m.memory[int64(uint32(v5))+4:], uint32(i32(1098889)))
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
					store32(m.memory[int64(uint32(v5))+28:], uint32(i32(1100399)))
					store16(m.memory[int64(uint32(v5))+24:], uint16(i32(2)))
					goto l6
				}
				v1 = i32(2)
				goto l5
			}
		l3:
			store16(m.memory[int64(uint32(v5))+24:], uint16(i32(2)))
			store32(m.memory[int64(uint32(v5))+20:], uint32(i32(1)))
			store32(m.memory[int64(uint32(v5))+16:], uint32(i32(1100399)))
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
func (m *Module) fn866(v0 int32, v1 int64, v2, v3, v4, v5 int32) {
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
				m.fn3(i32(1121000), i32(28), i32(1121028))
				panic("unreachable")
			}
			t1 := v2
			v7 = int64(bits.LeadingZeros64(uint64(v1)))
			v8 = t1 - int32(v7)
			v2 = (int32(int16(i32(-96)-v8))*i32(80) + i32(86960)) / i32(2126)
			if uint32(v2) > uint32(i32(80)) {
				m.fn33(v2, i32(81), i32(1121044))
				panic("unreachable")
			}
			t2 := v6
			v2 = v2 << 4
			t3 := int64(load64(m.memory[int64(uint32(v2))+1119704:]))
			m.fn976(t2, t3, i64(0), i64_shl(v1, v7), i64(0))
			t4 := int64(load64(m.memory[uint32(v6):]))
			t5 := int64(load64(m.memory[int64(uint32(v6))+8:]))
			v1 = int64(uint64(t4)>>63) + t5
			t6 := int32(load16(m.memory[int64(uint32(v2))+1119712:]))
			t7 := v1
			v9 = i32(-64) - (v8 + t6)
			v7 = int64(uint32(v9))
			v10 = int32(i64_shr_u(t7, v7))
			t8 := int32(load16(m.memory[int64(uint32(v2))+1119714:]))
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
				t9 := int32(load32(m.memory[uint32(v4<<2+i32(0x111dcc)):]))
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
			m.fn867(v0, v3, v4, i32(0), v16, v5, t29, i64_shl(int64(uint32(v8)), v15), v11)
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
				m.fn33(v4, v4, i32(1121076))
				panic("unreachable")
			}
			v10 = v10 - v9*v8
			m.memory[uint32(v3+v2)] = byte(v9 + i32(48))
			if v13 == v2 {
				m.fn867(v0, v3, v4, v18, v16, v5, i64_shl(int64(uint32(v10)), v15)+v7, i64_shl(int64(uint32(v8)), v15), v11)
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
						m.fn33(v2, v4, i32(0x111b44))
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
				m.fn867(v0, v3, v4, v18, v16, v5, v7, v11, v1)
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
		m.fn689(i32(1121060))
		panic("unreachable")
	}
l3:
	store32(m.memory[uint32(v0):], uint32(i32(0)))
l14:
	m.g0 = v6 + i32(16)
}
func (m *Module) fn867(v0, v1, v2, v3, v4, v5 int32, v6, v7, v8 int64) {
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
			m.fn121(i32(0), v3, v2, i32(1121972))
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
			m.fn121(i32(0), v3, v2, i32(1121956))
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
	m.fn121(i32(0), v3, v2, i32(1121940))
	panic("unreachable")
l11:
	store16(m.memory[int64(uint32(v0))+8:], uint16(v4))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
	store32(m.memory[uint32(v0):], uint32(v1))
	return
l4:
	store32(m.memory[uint32(v0):], uint32(i32(0)))
}
func (m *Module) fn868(v0, v1, v2, v3, v4 int32) {
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
			m.fn3(i32(1121000), i32(28), i32(1121456))
			panic("unreachable")
		}
		t2 := int64(load64(m.memory[int64(uint32(v1))+8:]))
		v7 = t2
		if v7 == i64(0) {
			m.fn3(i32(1121124), i32(29), i32(0x111cc0))
			panic("unreachable")
		}
		t3 := int64(load64(m.memory[int64(uint32(v1))+16:]))
		v8 = t3
		if v8 == i64(0) {
			m.fn3(i32(1121172), i32(28), i32(1121488))
			panic("unreachable")
		}
		if uint64(v8) > uint64(v6^i64(-1)) {
			m.fn3(i32(1121384), i32(54), i32(1121600))
			panic("unreachable")
		}
		if uint64(v6) < uint64(v7) {
			m.fn3(i32(1121312), i32(55), i32(1121584))
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
		_ = m.fn856(v5+i32(8), v1)
		goto l6
	l5:
		_ = m.fn856(v5+i32(176), int32(int16(i32(0)-v1)))
	l6:
		if v10 > i32(-1) {
			goto l7
		}
		_ = m.fn869(v5+i32(8), (i32(0)-v10)&i32(0xffff))
		goto l8
	l7:
		_ = m.fn869(v5+i32(176), v9&i32(0x7fff))
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
					m.fn121(i32(0), v1, i32(40), i32(1099872))
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
		t20 := int32(load32(m.memory[int64(uint32(v11<<2))+1121744:]))
		v9 = t20 << 1
		if v9 == 0 {
			m.fn3(i32(1099816), i32(27), i32(1099872))
			panic("unreachable")
		}
		{
			{
				{
					t21 := int32(load32(m.memory[int64(uint32(v5))+828:]))
					v1 = t21
					if uint32(v1) > uint32(i32(40)) {
						m.fn121(i32(0), v1, i32(40), i32(1099872))
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
			m.fn121(i32(0), v16, i32(40), i32(1099872))
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
			m.fn33(i32(40), i32(40), i32(1099872))
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
			m.fn121(i32(0), v1, i32(40), i32(1099872))
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
				m.fn33(i32(40), i32(40), i32(1099872))
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
				t67 := m.fn856(v5+i32(340), i32(1))
				v25 = t67
				memory_copy(m.memory, uint32(v5+i32(504)), uint32(v5+i32(176)), uint32(i32(164)))
				t68 := m.fn856(v5+i32(504), i32(2))
				v26 = t68
				memory_copy(m.memory, uint32(v5+i32(668)), uint32(v5+i32(176)), uint32(i32(164)))
				v17 = v5 + i32(176) + i32(-4)
				v4 = v5 + i32(340) + i32(-4)
				v16 = v5 + i32(504) + i32(-4)
				v18 = v5 + i32(668) + i32(-4)
				t69 := m.fn856(v5+i32(668), i32(3))
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
						m.fn121(i32(0), v15, i32(40), i32(1099872))
						panic("unreachable")
					}
					v31 = v32 + i32(1)
					v13 = v15 << 2
					v1 = i32(0)
				l49:
					{
						if v13 == v1 {
							if uint32(v24) > uint32(v3) {
								m.fn121(v32, v24, v3, i32(1121568))
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
						m.fn121(i32(0), v33, i32(40), i32(1099872))
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
						m.fn121(i32(0), v33, i32(40), i32(1099872))
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
						m.fn121(i32(0), v35, i32(40), i32(1099872))
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
						m.fn121(i32(0), v15, i32(40), i32(1099872))
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
						m.fn33(v3, v3, i32(1121552))
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
						m.fn33(i32(40), i32(40), i32(1099872))
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
	m.fn3(i32(1099843), i32(26), i32(1099872))
	panic("unreachable")
l76:
	m.fn3(i32(1099843), i32(26), i32(1099872))
	panic("unreachable")
l67:
	m.fn3(i32(1099843), i32(26), i32(1099872))
	panic("unreachable")
l58:
	m.fn3(i32(1099843), i32(26), i32(1099872))
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
			m.fn33(i32(40), i32(40), i32(1099872))
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
			m.fn121(i32(0), v1, i32(40), i32(1099872))
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
				m.fn33(v1, v3, i32(1121504))
				panic("unreachable")
			}
			t180 := int32(m.memory[uint32(v2+v1)])
			if t180&i32(1) == 0 {
				goto l111
			}
		}
	l110:
		if uint32(v24) > uint32(v3) {
			m.fn121(i32(0), v24, v3, i32(1121520))
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
		m.fn121(i32(0), v24, v3, i32(0x111d00))
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
func (m *Module) fn869(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7 int32
	var v8, v9 int64
	{
		if uint32(v1) < uint32(i32(8)) {
			t12 := int32(load32(m.memory[int64(uint32(v0))+160:]))
			v6 = t12
			if uint32(v6) > uint32(i32(40)) {
				m.fn121(i32(0), v6, i32(40), i32(1099872))
				panic("unreachable")
			}
			if v6 != 0 {
				t13 := int64(load32(m.memory[int64(uint32(v1<<2))+1121744:]))
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
					m.fn33(i32(40), i32(40), i32(1099872))
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
			m.fn121(i32(0), v3, i32(40), i32(1099872))
			panic("unreachable")
		}
		if v3 != 0 {
			v4 = v3 << 2
			v5 = v4 + i32(-4)
			v6 = int32(uint32(v5)>>2) + i32(1)
			v7 = v6 & i32(3)
			t1 := int32(load32(m.memory[int64(uint32(v2<<2))+1121744:]))
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
				m.fn33(i32(40), i32(40), i32(1099872))
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
			m.fn121(i32(0), v3, i32(40), i32(1099872))
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
			m.fn33(i32(40), i32(40), i32(1099872))
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
	_ = m.fn855(v0, i32(1121784), i32(2))
l27:
	if v1&i32(32) == 0 {
		goto l28
	}
	_ = m.fn855(v0, i32(0x111e00), i32(3))
l28:
	if v1&i32(64) == 0 {
		goto l29
	}
	_ = m.fn855(v0, i32(1121804), i32(5))
l29:
	if v1&i32(128) == 0 {
		goto l30
	}
	_ = m.fn855(v0, i32(1121824), i32(10))
l30:
	if v1&i32(256) == 0 {
		goto l31
	}
	_ = m.fn855(v0, i32(1121864), i32(19))
l31:
	_ = m.fn856(v0, v1)
	return v0
}
func (m *Module) fn870(v0, v1, v2 int32) {
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
			m.fn3(i32(1121000), i32(28), i32(1121108))
			panic("unreachable")
		}
		t2 := int64(load64(m.memory[int64(uint32(v1))+8:]))
		v5 = t2
		if v5 == i64(0) {
			m.fn3(i32(1121124), i32(29), i32(1121156))
			panic("unreachable")
		}
		t3 := int64(load64(m.memory[int64(uint32(v1))+16:]))
		v6 = t3
		if v6 == i64(0) {
			m.fn3(i32(1121172), i32(28), i32(1121200))
			panic("unreachable")
		}
		v7 = v6 + v4
		if uint64(v7) < uint64(v6) {
			m.fn3(i32(1121384), i32(54), i32(1121440))
			panic("unreachable")
		}
		if uint64(v4) < uint64(v5) {
			m.fn3(i32(1121312), i32(55), i32(1121368))
			panic("unreachable")
		}
		if uint64(v7) >= uint64(i64(0x2000000000000000)) {
			m.fn3(i32(1121216), i32(45), i32(1121264))
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
			m.fn851(v3+i32(72), v3+i32(56))
			panic("unreachable")
		}
		store16(m.memory[int64(uint32(v3))+64:], uint16(v1))
		store64(m.memory[int64(uint32(v3))+56:], uint64(v4))
		t9 := v3
		v9 = i64_shl(v4, v6)
		v5 = i64_shr_u(v9, v6)
		store64(m.memory[int64(uint32(t9))+72:], uint64(v5))
		if v5 != v4 {
			m.fn851(v3+i32(72), v3+i32(56))
			panic("unreachable")
		}
		v10 = v1 - int32(v6)
		v1 = (int32(int16(i32(-96)-v10))*i32(80) + i32(86960)) / i32(2126)
		if uint32(v1) > uint32(i32(80)) {
			m.fn33(v1, i32(81), i32(1121044))
			panic("unreachable")
		}
		t10 := v3 + i32(32)
		v1 = v1 << 4
		t11 := int64(load64(m.memory[int64(uint32(v1))+1119704:]))
		v4 = t11
		m.fn976(t10, v4, i64(0), i64_shl(v7, v6), i64(0))
		m.fn976(v3+i32(16), v4, i64(0), v8, i64(0))
		m.fn976(v3, v4, i64(0), v9, i64(0))
		t12 := int32(load16(m.memory[int64(uint32(v1))+1119712:]))
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
		t16 := int32(load16(m.memory[int64(uint32(v1))+1119714:]))
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
					m.fn33(i32(17), i32(17), i32(1121296))
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
	m.fn689(i32(0x111c00))
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
func (m *Module) fn871(v0, v1, v2 int32) {
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
			m.fn3(i32(1121000), i32(28), i32(1121616))
			panic("unreachable")
		}
		t2 := int64(load64(m.memory[int64(uint32(v1))+8:]))
		v5 = t2
		if v5 == i64(0) {
			m.fn3(i32(1121124), i32(29), i32(1121632))
			panic("unreachable")
		}
		t3 := int64(load64(m.memory[int64(uint32(v1))+16:]))
		v6 = t3
		if v6 == i64(0) {
			m.fn3(i32(1121172), i32(28), i32(1121648))
			panic("unreachable")
		}
		v7 = v6 + v4
		if uint64(v7) < uint64(v6) {
			m.fn3(i32(1121384), i32(54), i32(1121728))
			panic("unreachable")
		}
		if uint64(v4) < uint64(v5) {
			m.fn3(i32(1121312), i32(55), i32(1121712))
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
		_ = m.fn856(v3+i32(8), v1)
		_ = m.fn856(v3+i32(176), v1)
		_ = m.fn856(v3+i32(344), v1)
		goto l6
	l5:
		_ = m.fn856(v3+i32(508), int32(int16(i32(0)-v1)))
	l6:
		{
			if v10 > i32(-1) {
				goto l7
			}
			t16 := v3 + i32(8)
			v1 = (i32(0) - v10) & i32(0xffff)
			_ = m.fn869(t16, v1)
			_ = m.fn869(v3+i32(176), v1)
			_ = m.fn869(v3+i32(344), v1)
			goto l8
		}
	l7:
		_ = m.fn869(v3+i32(508), v9&i32(0x7fff))
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
			m.fn121(i32(0), v12, i32(40), i32(1099872))
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
			m.fn33(i32(40), i32(40), i32(1099872))
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
			m.fn121(i32(0), v1, i32(40), i32(1099872))
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
					m.fn121(i32(0), v13, i32(40), i32(1099872))
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
				m.fn33(i32(40), i32(40), i32(1099872))
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
					m.fn121(i32(0), v14, i32(40), i32(1099872))
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
				m.fn33(i32(40), i32(40), i32(1099872))
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
				m.fn33(i32(40), i32(40), i32(1099872))
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
		t83 := m.fn856(v3+i32(672), i32(1))
		v23 = t83
		memory_copy(m.memory, uint32(v3+i32(836)), uint32(v3+i32(508)), uint32(i32(164)))
		t84 := m.fn856(v3+i32(836), i32(2))
		v24 = t84
		memory_copy(m.memory, uint32(v3+i32(1000)), uint32(v3+i32(508)), uint32(i32(164)))
		{
			{
				t85 := m.fn856(v3+i32(1000), i32(3))
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
						m.fn121(i32(0), v27, i32(40), i32(1099872))
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
						m.fn121(i32(0), v34, i32(40), i32(1099872))
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
						m.fn121(i32(0), v27, i32(40), i32(1099872))
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
						m.fn33(i32(17), i32(17), i32(1121664))
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
						m.fn121(i32(0), v1, i32(40), i32(1099872))
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
						m.fn121(i32(0), v34, i32(40), i32(1099872))
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
						m.fn33(i32(40), i32(40), i32(1099872))
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
						m.fn121(i32(0), v1, i32(40), i32(1099872))
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
						m.fn33(i32(40), i32(40), i32(1099872))
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
						m.fn33(i32(40), i32(40), i32(1099872))
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
						m.fn33(i32(40), i32(40), i32(1099872))
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
			m.fn121(i32(0), v27, i32(40), i32(1099872))
			panic("unreachable")
		l102:
			if v1 >= v8 {
				goto l129
			}
			_ = m.fn856(v3+i32(8), i32(1))
			t233 := int32(load32(m.memory[int64(uint32(v3))+168:]))
			t234 := v21
			v1 = t233
			p235 := v1
			if uint32(v21) > uint32(v1) {
				p235 = t234
			}
			v1 = p235
			if uint32(v1) >= uint32(i32(41)) {
				m.fn121(i32(0), v1, i32(40), i32(1099872))
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
						m.fn33(v31, i32(17), i32(1121680))
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
	m.fn121(i32(0), v31, i32(17), i32(1121696))
	panic("unreachable")
l136:
	store16(m.memory[int64(uint32(v0))+8:], uint16(v10))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v31))
	store32(m.memory[uint32(v0):], uint32(v2))
	m.g0 = v3 + i32(1328)
	return
l84:
	m.fn3(i32(1099843), i32(26), i32(1099872))
	panic("unreachable")
l75:
	m.fn3(i32(1099843), i32(26), i32(1099872))
	panic("unreachable")
l66:
	m.fn3(i32(1099843), i32(26), i32(1099872))
	panic("unreachable")
l57:
	m.fn3(i32(1099843), i32(26), i32(1099872))
	panic("unreachable")
}
func (m *Module) fn872(v0, v1 int32) int32 {
	var v2, v3 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	t1 := int32(m.memory[uint32(v0)])
	v3 = t1
	v0 = i32(3)
l0:
	{
		t2 := int32(m.memory[uint32(v3&i32(15)+i32(1122552))])
		m.memory[uint32(v2+i32(14)+v0+i32(-2))] = byte(t2)
		v0 = v0 + i32(-1)
		v3 = int32(uint32(v3)>>4) & i32(15)
		if v3 != 0 {
			goto l0
		}
	}
	t3 := m.fn306(v1, i32(1), i32(1122550), i32(2), v2+i32(14)+v0+i32(-1), i32(3)-v0)
	v0 = t3
	m.g0 = v2 + i32(16)
	return v0
}
func (m *Module) fn873(v0, v1 int32) int32 {
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
			t4 := m.fn874(t2, t3, v1)
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
						t31 := m.fn874(v7, v6, v2)
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
			t26 := m.fn874(t24, t25, v2)
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
func (m *Module) fn874(v0, v1, v2 int32) int32 {
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
										t10 := m.t0[uint(v5)].(func(int32, int32, int32) int32)(v0, i32(1100008), i32(64))
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
						t17 := m.t0[uint(v5)].(func(int32, int32, int32) int32)(v0, i32(1100008), v2)
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
func (m *Module) fn875(v0 int32, v1, v2 int64) {
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
		t2 := int64(load64(m.memory[uint32(v4+i32(1114664)):]))
		t3 := v2
		v6 = int64(bits.LeadingZeros64(uint64(v2)))
		v7 = i64_shl(t3, v6)
		m.fn976(t1, t2, i64(0), v7, i64(0))
		t4 := int64(load64(m.memory[int64(uint32(v3))+16:]))
		v8 = t4
		{
			t5 := int64(load64(m.memory[int64(uint32(v3))+24:]))
			v2 = t5
			if v2&i64(511) != i64(511) {
				goto l2
			}
			t6 := int64(load64(m.memory[uint32(v4+i32(1109192)+i32(5480)):]))
			m.fn976(v3, t6, i64(0), v7, i64(0))
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
func (m *Module) fn876(v0, v1 int32) int32 {
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
			t16 := m.fn879(t13, t14, p15, t12)
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
			t7 := m.fn877(t4, t5, p6, i32(1))
			return t7
		}
	l1:
		t8 := v1
		t9 := v4
		var p10 int32
		if v3 != i32(0) {
			p10 = 1
		}
		t11 := m.fn878(t8, t9, p10)
		return t11
	}
}
func (m *Module) fn877(v0 int32, v1 float64, v2, v3 int32) int32 {
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
					store32(m.memory[int64(uint32(v4))+36:], uint32(i32(1098882)))
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
				p21 := i32(1098881)
				if v12 != 0 {
					p21 = i32(1098880)
				}
				p22 := i32(1)
				if v12 != 0 {
					p22 = i32(1098880)
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
					store32(m.memory[int64(uint32(v4))+36:], uint32(i32(1098885)))
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
				store32(m.memory[int64(uint32(v4))+36:], uint32(i32(1098888)))
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
				p15 = i32(1098880)
			}
			v11 = p15
			p16 := i32(1098881)
			if v12 != 0 {
				p16 = i32(1098880)
			}
			v12 = p16
			v3 = int32(int64(uint64(v5) >> 63))
			m.fn870(v4+i32(32), v4+i32(96), v4+i32(15))
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
		m.fn871(v4+i32(80), v4+i32(96), v4+i32(15))
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
		m.fn865(v4, t27, t28, t29, v10, v4+i32(32))
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
	store32(m.memory[int64(uint32(v4))+36:], uint32(i32(1098889)))
	v11 = v4 + i32(32)
l11:
	store32(m.memory[int64(uint32(v4))+92:], uint32(v10))
	store32(m.memory[int64(uint32(v4))+88:], uint32(v11))
	store32(m.memory[int64(uint32(v4))+84:], uint32(v2))
	store32(m.memory[int64(uint32(v4))+80:], uint32(v12))
	t32 := m.fn873(v0, v4+i32(80))
	v10 = t32
	m.g0 = v4 + i32(128)
	return v10
}
func (m *Module) fn878(v0 int32, v1 float64, v2 int32) int32 {
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
						store32(m.memory[int64(uint32(v3))+28:], uint32(i32(1098882)))
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
					p18 := i32(1098881)
					if v10 != 0 {
						p18 = i32(1098880)
					}
					p19 := i32(1)
					if v10 != 0 {
						p19 = i32(1098880)
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
						store32(m.memory[int64(uint32(v3))+28:], uint32(i32(1098891)))
						goto l10
					}
					store32(m.memory[int64(uint32(v3))+32:], uint32(i32(3)))
					store32(m.memory[int64(uint32(v3))+28:], uint32(i32(1098885)))
					goto l10
				}
				m.fn870(v3+i32(96), v3+i32(112), v3+i32(7))
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
			m.fn871(v3+i32(144), v3+i32(112), v3+i32(7))
		l7:
			t22 := int32(load32(m.memory[int64(uint32(v3))+148:]))
			v9 = t22
			if v9 == 0 {
				m.fn3(i32(1102209), i32(33), i32(1102308))
				panic("unreachable")
			}
			t23 := int32(load32(m.memory[int64(uint32(v3))+144:]))
			v11 = t23
			t24 := int32(m.memory[uint32(v11)])
			if uint32(t24) <= uint32(i32(48)) {
				m.fn3(i32(1102260), i32(31), i32(1102324))
				panic("unreachable")
			}
			var p25 int32
			if v4 < i64(0) {
				p25 = 1
			}
			v10 = p25
			p26 := i32(1)
			if v10 != 0 {
				p26 = i32(1098880)
			}
			v13 = p26
			p27 := i32(1098881)
			if v10 != 0 {
				p27 = i32(1098880)
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
				store32(m.memory[int64(uint32(v3))+40:], uint32(i32(1100399)))
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
		p34 := i32(1102340)
		if v13 != 0 {
			p34 = i32(1102341)
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
	t38 := m.fn873(v0, v3+i32(96))
	v9 = t38
	m.g0 = v3 + i32(160)
	return v9
}
func (m *Module) fn879(v0 int32, v1 float64, v2, v3 int32) int32 {
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
					store32(m.memory[int64(uint32(v4))+1044:], uint32(i32(1098882)))
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
				p16 := i32(1098881)
				if v12 != 0 {
					p16 = i32(1098880)
				}
				p17 := i32(1)
				if v12 != 0 {
					p17 = i32(1098880)
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
					store32(m.memory[int64(uint32(v4))+1044:], uint32(i32(1098885)))
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
					store32(m.memory[int64(uint32(v4))+1044:], uint32(i32(1098889)))
					v3 = v4 + i32(1040)
					goto l10
				}
				v10 = i32(1)
				store32(m.memory[int64(uint32(v4))+1048:], uint32(i32(1)))
				store32(m.memory[int64(uint32(v4))+1044:], uint32(i32(1098888)))
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
			m.fn3(i32(1098956), i32(37), i32(1098996))
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
			p21 = i32(1098880)
		}
		v14 = p21
		p22 := i32(1098881)
		if v7 != 0 {
			p22 = i32(1098880)
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
		m.fn866(t23, t24, t25, t26, t28, v10)
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
		m.fn868(v4+i32(1088), v4+i32(1104), v4+i32(16), v16, v10)
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
			m.fn865(v4+i32(8), t35, t36, v7, v11, v4+i32(1040))
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
		store32(m.memory[int64(uint32(v4))+1044:], uint32(i32(1098888)))
		v3 = v4 + i32(1040)
		goto l10
	l14:
		store32(m.memory[int64(uint32(v4))+1056:], uint32(v11))
		store16(m.memory[int64(uint32(v4))+1052:], uint16(i32(0)))
		store32(m.memory[int64(uint32(v4))+1048:], uint32(i32(2)))
		store32(m.memory[int64(uint32(v4))+1044:], uint32(i32(1098889)))
		v3 = v4 + i32(1040)
	}
l10:
	store32(m.memory[int64(uint32(v4))+1100:], uint32(v10))
	store32(m.memory[int64(uint32(v4))+1096:], uint32(v3))
	store32(m.memory[int64(uint32(v4))+0x444:], uint32(v2))
	store32(m.memory[int64(uint32(v4))+1088:], uint32(v12))
	t39 := m.fn873(v0, v4+i32(1088))
	v10 = t39
	m.g0 = v4 + i32(1136)
	return v10
}
func (m *Module) fn880(v0 int32, v1 float64, v2, v3 int32) int32 {
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
					store32(m.memory[int64(uint32(v4))+1036:], uint32(i32(1098882)))
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
				p18 := i32(1098881)
				if v12 != 0 {
					p18 = i32(1098880)
				}
				p19 := i32(1)
				if v12 != 0 {
					p19 = i32(1098880)
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
					store32(m.memory[int64(uint32(v4))+1036:], uint32(i32(1098885)))
					store16(m.memory[int64(uint32(v4))+1032:], uint16(i32(2)))
					goto l10
				}
				store16(m.memory[int64(uint32(v4))+1032:], uint16(i32(2)))
				if v3&i32(0xffff) != 0 {
					store32(m.memory[int64(uint32(v4))+1064:], uint32(i32(2)))
					store32(m.memory[int64(uint32(v4))+1060:], uint32(i32(1098894)))
					store16(m.memory[int64(uint32(v4))+1056:], uint16(i32(2)))
					store32(m.memory[int64(uint32(v4))+1048:], uint32(v11))
					store16(m.memory[int64(uint32(v4))+1044:], uint16(i32(0)))
					store32(m.memory[int64(uint32(v4))+1040:], uint32(i32(2)))
					store32(m.memory[int64(uint32(v4))+1036:], uint32(i32(1098889)))
					v10 = i32(3)
					goto l10
				}
				store32(m.memory[int64(uint32(v4))+1040:], uint32(i32(3)))
				store32(m.memory[int64(uint32(v4))+1036:], uint32(i32(1098891)))
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
			m.fn3(i32(1098896), i32(41), i32(1098940))
			panic("unreachable")
		}
	l6:
		m.fn866(v4+i32(1104), v8, v12, v4+i32(8), v10, i32(0x8000))
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
		m.fn868(v4+i32(1152), v4+i32(1120), v4+i32(8), v10, i32(0x8000))
	l12:
		{
			t25 := int32(load32(m.memory[int64(uint32(v4))+1156:]))
			v14 = t25
			if v14 == 0 {
				m.fn3(i32(1102209), i32(33), i32(1102308))
				panic("unreachable")
			}
			t26 := int32(load32(m.memory[int64(uint32(v4))+1152:]))
			v15 = t26
			t27 := int32(m.memory[uint32(v15)])
			if uint32(t27) <= uint32(i32(48)) {
				m.fn3(i32(1102260), i32(31), i32(1102324))
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
				store32(m.memory[int64(uint32(v4))+1048:], uint32(i32(1100399)))
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
			v19 = i32(1100399)
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
			p30 = i32(1098880)
		}
		v12 = p30
		p31 := i32(1098881)
		if v7 != 0 {
			p31 = i32(1098880)
		}
		v11 = p31
		v3 = int32(v5)
		if v16 < i32(1) {
			goto l19
		}
		v7 = v4 + i32(1032) + v10*i32(12)
		store32(m.memory[int64(uint32(v7))+8:], uint32(i32(1)))
		store32(m.memory[int64(uint32(v7))+4:], uint32(i32(1102340)))
		store16(m.memory[uint32(v7):], uint16(i32(2)))
		v7 = v16 + i32(-1)
		goto l20
	l19:
		v7 = v4 + i32(1032) + v10*i32(12)
		store32(m.memory[int64(uint32(v7))+8:], uint32(i32(2)))
		store32(m.memory[int64(uint32(v7))+4:], uint32(i32(1102341)))
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
	t34 := m.fn873(v0, v4+i32(1104))
	v10 = t34
	m.g0 = v4 + i32(1168)
	return v10
}
func (m *Module) fn881(v0, v1 int32) int32 {
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
			store64(m.memory[int64(uint32(v2))+24:], uint64(int64(uint32(i32(2)))<<32|int64(uint32(v0))))
			store64(m.memory[int64(uint32(v2))+16:], uint64(int64(uint32(i32(19)))<<32|int64(uint32(v2+i32(15)))))
			t3 := int32(load32(m.memory[uint32(v1):]))
			t4 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t5 := m.fn45(t3, t4, i32(1049276), v2+i32(16))
			v0 = t5
			goto l1
		}
	l0:
		store64(m.memory[int64(uint32(v2))+16:], uint64(int64(uint32(i32(2)))<<32|int64(uint32(v0))))
		t6 := int32(load32(m.memory[uint32(v1):]))
		t7 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t8 := m.fn45(t6, t7, i32(1049325), v2+i32(16))
		v0 = t8
	}
l1:
	m.g0 = v2 + i32(32)
	return v0
}
func (m *Module) fn882(v0, v1 int32) int32 {
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
			t0 := m.fn5(v0)
			v2 = t0
			if v2 == 0 {
				m.fn10(i32(1), v0)
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
		t2 := m.fn5(v3)
		v4 = t2
		if v4 == 0 {
			m.fn10(i32(4), v3)
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
		t4 := m.fn5(i32(64))
		v3 = t4
		if v3 == 0 {
			m.fn24(i32(8), i32(64))
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
	m.fn9()
	panic("unreachable")
}
func (m *Module) fn883(v0, v1, v2 int32) {
	var v3, v4, v5, v6, v7, v8, v9 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+44:]))
		v3 = t0
		t1 := int32(load32(m.memory[int64(uint32(v0))+40:]))
		t2 := v3
		v4 = t1
		if uint32(t2) > uint32(v4) {
			m.fn121(i32(0), v3, v4, i32(1139148))
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
					m.fn884(v7, v3, v4, i32(1), i32(1))
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
			m.fn121(v5, v6, v3, i32(1139100))
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
				m.fn884(v0+i32(32), v3, v8, i32(4), i32(4))
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
			m.fn33(v4, v3, i32(1139132))
			panic("unreachable")
		}
		store32(m.memory[int64(uint32(v0))+44:], uint32(v4+i32(1)))
		t22 := int32(load32(m.memory[int64(uint32(v0))+36:]))
		store32(m.memory[uint32(t22+v4<<2):], uint32(v6))
		return
	}
}
func (m *Module) fn884(v0, v1, v2, v3, v4 int32) {
	var v5 int32
	t0 := m.g0
	v5 = t0 - i32(16)
	m.g0 = v5
	v1 = v2 + v1
	if uint32(v1) >= uint32(v2) {
		goto l0
	}
	m.fn10(i32(0), i32(0))
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
	m.fn885(t2, t4, t3, v2, v3, v4)
	{
		t10 := int32(load32(m.memory[int64(uint32(v5))+4:]))
		if t10 != i32(1) {
			goto l1
		}
		t11 := int32(load32(m.memory[int64(uint32(v5))+8:]))
		t12 := int32(load32(m.memory[int64(uint32(v5))+12:]))
		m.fn10(t11, t12)
		panic("unreachable")
	}
l1:
	t13 := int32(load32(m.memory[int64(uint32(v5))+8:]))
	v4 = t13
	store32(m.memory[uint32(v0):], uint32(v2))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v4))
	m.g0 = v5 + i32(16)
}
func (m *Module) fn885(v0, v1, v2, v3, v4, v5 int32) {
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
			t0 := m.fn22(v2, v5*v1, v4, v3)
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
		t1 := m.fn20(v3, v4)
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
func (m *Module) fn886(v0, v1 int32) int32 {
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
			t16 := int32(m.memory[int64(uint32(v3))+1270976])
			t17 := int32(m.memory[uint32(v7+i32(1))])
			t18 := int32(m.memory[int64(uint32(t17))+1270848])
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
				t10 := int32(m.memory[int64(uint32(v7&i32(255)))+1270976])
				v7 = v0 + v3
				t11 := int32(m.memory[uint32(v7+i32(1))])
				t12 := int32(m.memory[int64(uint32(t11))+1270848])
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
	m.fn121(v2, v1, v1, i32(1270816))
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
					t24 := int32(m.memory[int64(uint32(v7&i32(255)))+1270976])
					t25 := int32(m.memory[int64(uint32(v6))+1])
					t26 := int32(m.memory[int64(uint32(t25))+1270848])
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
func (m *Module) fn887(v0, v1, v2 int32) {
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
						if v7 == i32(1143904) {
							goto l13
						}
						if v7 == i32(1139800) {
							goto l13
						}
						if v7 == i32(1143932) {
							goto l13
						}
						m.fn891(v3+i32(8), v1, v2)
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
			m.fn891(v3+i32(16), v1, v2)
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
					if t11 == i32(1139800) {
						goto l13
					}
					m.fn891(v3+i32(24), v1, v2)
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
							if v2 == i32(1143932) {
								goto l13
							}
							if v2 == i32(1143904) {
								goto l13
							}
							m.fn891(v3+i32(32), v1, v7)
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
				m.fn891(v3+i32(40), v1, v2+i32(2))
				t20 := int32(load32(m.memory[int64(uint32(v3))+44:]))
				v5 = t20
				t21 := int32(load32(m.memory[int64(uint32(v3))+40:]))
				v4 = t21
				goto l7
			}
			v4 = i32(0)
			goto l7
		case 10:
			m.fn3(i32(1145976), i32(41), i32(1146068))
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
func (m *Module) fn888(v0, v1, v2 int32) {
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
					if v7 == i32(1143904) {
						goto l12
					}
					if v7 == i32(1139800) {
						goto l12
					}
					if v7 == i32(1143932) {
						goto l12
					}
					m.fn892(v3+i32(8), v1, v2)
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
			m.fn892(v3+i32(16), v1, v2)
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
						if t10 == i32(1139800) {
							goto l12
						}
						m.fn892(v3+i32(24), v1, v7)
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
							if v2 == i32(1143932) {
								goto l12
							}
							if v2 == i32(1143904) {
								goto l12
							}
							m.fn892(v3+i32(32), v1, v7)
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
				m.fn892(v3+i32(40), v1, v2+i32(2))
				t19 := int32(load32(m.memory[int64(uint32(v3))+44:]))
				v6 = t19
				t20 := int32(load32(m.memory[int64(uint32(v3))+40:]))
				v4 = t20
				goto l7
			}
			v4 = i32(0)
			goto l7
		case 10:
			m.fn3(i32(1145976), i32(41), i32(1146020))
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
func (m *Module) fn889(v0, v1, v2, v3, v4, v5 int32) {
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
					m.fn894(v0, v1, v2, i32(0), v4, v5, i32(0), i32(187))
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
						m.fn894(v0, v1, v2, v3, v4, v5, i32(0), i32(187))
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
								m.fn894(v0, v1, v2, v3, v4, v5, v8, i32(239))
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
			m.fn894(v0, v1, v2, v3, v4, v5, v8, i32(239))
			goto l12
		l19:
			m.fn895(v0, v1, v2, v8, v4, v5, i32(1))
			t7 := int32(m.memory[int64(uint32(v0))+4])
			if t7 != 0 {
				goto l12
			}
			m.memory[int64(uint32(v1))+24] = byte(i32(10))
			goto l12
		}
	l11:
		m.fn3(i32(1145976), i32(41), i32(1146036))
		panic("unreachable")
	l6:
		if uint32(v8) < uint32(v3) {
			goto l28
		}
		m.fn896(v0, v1, v2, v3, v4, v5, v8)
		goto l12
	l7:
		if uint32(v8) < uint32(v3) {
			{
				t15 := int32(m.memory[uint32(v2+v8)])
				if t15 != i32(255) {
					m.fn894(v0, v1, v2, v3, v4, v5, v8, i32(254))
					goto l12
				}
				m.memory[int64(uint32(v1))+24] = byte(i32(9))
				v7 = v8 + i32(1)
				t16 := int32(load32(m.memory[int64(uint32(v1))+20:]))
				if t16 == i32(1143904) {
					goto l35
				}
				store32(m.memory[int64(uint32(v1))+4:], uint32(i32(65536)))
				m.memory[int64(uint32(v1))+2] = byte(i32(0))
				m.memory[uint32(v1)] = byte(i32(10))
				store32(m.memory[int64(uint32(v1))+20:], uint32(i32(1143904)))
				goto l35
			}
		l35:
			m.fn895(v6+i32(4), v1, v2+v7, v3-v7, v4, v5, i32(1))
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
		m.fn894(v0, v1, v2, v3, v4, v5, v8, i32(254))
		goto l12
	l8:
		if uint32(v8) < uint32(v3) {
			{
				t8 := int32(m.memory[uint32(v2+v8)])
				if t8 != i32(254) {
					m.fn894(v0, v1, v2, v3, v4, v5, v8, i32(255))
					goto l12
				}
				m.memory[int64(uint32(v1))+24] = byte(i32(9))
				v7 = v8 + i32(1)
				t9 := int32(load32(m.memory[int64(uint32(v1))+20:]))
				if t9 == i32(1143932) {
					goto l32
				}
				store32(m.memory[int64(uint32(v1))+4:], uint32(i32(0)))
				m.memory[int64(uint32(v1))+2] = byte(i32(0))
				m.memory[uint32(v1)] = byte(i32(10))
				store32(m.memory[int64(uint32(v1))+20:], uint32(i32(1143932)))
				goto l32
			}
		l32:
			m.fn895(v6+i32(4), v1, v2+v7, v3-v7, v4, v5, i32(1))
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
		m.fn894(v0, v1, v2, v3, v4, v5, v8, i32(255))
		goto l12
	l28:
		{
			t22 := int32(m.memory[uint32(v2+v8)])
			if t22 != i32(191) {
				m.fn896(v0, v1, v2, v3, v4, v5, v8)
				goto l12
			}
			m.memory[int64(uint32(v1))+24] = byte(i32(9))
			v7 = v8 + i32(1)
			t23 := int32(load32(m.memory[int64(uint32(v1))+20:]))
			if t23 == i32(1139800) {
				goto l38
			}
			store16(m.memory[int64(uint32(v1))+16:], uint16(i32(49024)))
			store32(m.memory[int64(uint32(v1))+12:], uint32(i32(0)))
			store64(m.memory[int64(uint32(v1))+4:], uint64(i64(0)))
			m.memory[uint32(v1)] = byte(i32(1))
			store32(m.memory[int64(uint32(v1))+20:], uint32(i32(1139800)))
			goto l38
		}
	l38:
		m.fn895(v6+i32(4), v1, v2+v7, v3-v7, v4, v5, i32(1))
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
func (m *Module) fn890(v0, v1, v2 int32) {
	var v3 int32
	t0 := m.g0
	v3 = t0 - i32(16)
	m.g0 = v3
	v1 = v2 + v1
	if uint32(v1) >= uint32(v2) {
		goto l0
	}
	m.fn10(i32(0), i32(0))
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
	m.fn893(t2, t4, t3, v2)
	{
		t8 := int32(load32(m.memory[int64(uint32(v3))+4:]))
		if t8 != i32(1) {
			goto l1
		}
		t9 := int32(load32(m.memory[int64(uint32(v3))+8:]))
		t10 := int32(load32(m.memory[int64(uint32(v3))+12:]))
		m.fn10(t9, t10)
		panic("unreachable")
	}
l1:
	t11 := int32(load32(m.memory[int64(uint32(v3))+8:]))
	v1 = t11
	store32(m.memory[uint32(v0):], uint32(v2))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
	m.g0 = v3 + i32(16)
}
func (m *Module) fn891(v0, v1, v2 int32) {
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
func (m *Module) fn892(v0, v1, v2 int32) {
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
func (m *Module) fn893(v0, v1, v2, v3 int32) {
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
			t0 := m.fn22(v2, v1, i32(1), v3)
			v4 = t0
			goto l3
		}
	l2:
		t1 := m.fn5(v3)
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
func (m *Module) fn894(v0, v1, v2, v3, v4, v5, v6, v7 int32) {
	var v8, v9, v10 int32
	t0 := m.g0
	v8 = t0 - i32(16)
	m.g0 = v8
	m.memory[int64(uint32(v1))+24] = byte(i32(9))
	{
		{
			if v6 != 0 {
				m.fn895(v0, v1, v2, v3, v4, v5, i32(1))
				t4 := int32(m.memory[int64(uint32(v0))+4])
				if t4 != 0 {
					goto l4
				}
				m.memory[int64(uint32(v1))+24] = byte(i32(10))
				goto l4
			}
			m.memory[int64(uint32(v8))+3] = byte(v7)
			v7 = i32(0)
			m.fn895(v8+i32(4), v1, v8+i32(3), i32(1), v4, v5, i32(0))
			t1 := int32(load32(m.memory[int64(uint32(v8))+12:]))
			v6 = t1
			t2 := int32(load16(m.memory[int64(uint32(v8))+9:]))
			v9 = t2
			v10 = i32(2)
			t3 := int32(m.memory[int64(uint32(v8))+8])
			switch t3 {
			case 1:
				m.fn3(i32(1146084), i32(39), i32(1146036))
				panic("unreachable")
			case 2:
				goto l3
			default:
				goto l1
			}
		}
	l1:
		if uint32(v5) < uint32(v6) {
			m.fn121(v6, v5, v5, i32(1146036))
			panic("unreachable")
		}
		m.fn895(v8+i32(4), v1, v2, v3, v4+v6, v5-v6, i32(1))
		{
			t5 := int32(m.memory[int64(uint32(v8))+8])
			v10 = t5
			if v10 != 0 {
				goto l6
			}
			m.memory[int64(uint32(v1))+24] = byte(i32(10))
		}
	l6:
		t6 := int32(load32(m.memory[int64(uint32(v8))+12:]))
		v6 = t6 + v6
		t7 := int32(load32(m.memory[int64(uint32(v8))+4:]))
		v7 = t7
		t8 := int32(load16(m.memory[int64(uint32(v8))+9:]))
		v9 = t8
	}
l3:
	store16(m.memory[int64(uint32(v0))+5:], uint16(v9))
	m.memory[int64(uint32(v0))+4] = byte(v10)
	store32(m.memory[int64(uint32(v0))+8:], uint32(v6))
	store32(m.memory[uint32(v0):], uint32(v7))
l4:
	m.g0 = v8 + i32(16)
}
func (m *Module) fn895(v0, v1, v2, v3, v4, v5, v6 int32) {
	var v7, v8, v9, v10, v11, v12, v13, v14, v15, v16, v17, v18, v19, v20, v21, v22, v23, v24, v25, v26, v27 int32
	t0 := m.g0
	v7 = t0 - i32(64)
	m.g0 = v7
	{
		t1 := int32(m.memory[uint32(v1)])
		switch t1 {
		default:
			v8 = v3 + i32(-1)
			t2 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			v9 = t2
			v10 = v2
			v11 = i32(0)
			v12 = i32(0)
		l36:
			{
				v13 = v5 - v12
				t3 := v13
				v1 = v3 - v11
				t4 := v1
				var p5 int32
				if uint32(v13) < uint32(v1) {
					p5 = 1
				}
				v14 = p5
				p6 := t4
				if v14 != 0 {
					p6 = t3
				}
				v15 = p6
				v13 = i32(0)
				v16 = v4 + v12
				if (v16^v10)&i32(3) != 0 {
					goto l11
				}
				v13 = i32(0)
				v17 = (i32(0) - v10) & i32(3)
				if uint32(v17|i32(8)) > uint32(v15) {
					goto l11
				}
				{
					if v17 != 0 {
						goto l12
					}
					v13 = i32(0)
					goto l13
				l12:
					v13 = i32(0)
					t7 := int32(int8(m.memory[uint32(v10)]))
					v1 = t7
					if v1 < i32(0) {
						goto l14
					}
					m.memory[uint32(v16)] = byte(v1)
					v13 = i32(1)
					if v17 == i32(1) {
						goto l13
					}
					{
						t8 := int32(int8(m.memory[int64(uint32(v10))+1]))
						v1 = t8
						if v1 >= i32(0) {
							goto l15
						}
						v13 = i32(1)
						goto l14
					}
				l15:
					m.memory[int64(uint32(v16))+1] = byte(v1)
					v13 = i32(2)
					if v17 == i32(2) {
						goto l13
					}
					{
						t9 := int32(int8(m.memory[int64(uint32(v10))+2]))
						v1 = t9
						if v1 >= i32(0) {
							goto l16
						}
						v13 = i32(2)
						goto l14
					}
				l16:
					m.memory[int64(uint32(v16))+2] = byte(v1)
					v13 = i32(3)
				}
			l13:
				v18 = v15 + i32(-8)
			l20:
				{
					v1 = v16 + v13
					t10 := v1
					v17 = v10 + v13
					t11 := int32(load32(m.memory[uint32(v17):]))
					v6 = t11
					store32(m.memory[uint32(t10):], uint32(v6))
					t12 := int32(load32(m.memory[uint32(v17+i32(4)):]))
					t13 := v1 + i32(4)
					v1 = t12
					store32(m.memory[uint32(t13):], uint32(v1))
					{
						v17 = v1 & i32(-2139062144)
						t14 := v17
						v1 = v6 & i32(-2139062144)
						if t14|v1 == 0 {
							goto l17
						}
						if v1 != 0 {
							goto l18
						}
						v1 = int32(uint32(int32(bits.TrailingZeros32(uint32(v17))))>>3) + i32(4)
						goto l19
					l18:
						v1 = int32(uint32(int32(bits.TrailingZeros32(uint32(v1)))) >> 3)
					l19:
						t15 := v10
						v13 = v1 + v13
						t16 := int32(m.memory[uint32(t15+v13)])
						v1 = t16
						goto l14
					}
				l17:
					v13 = v13 + i32(8)
					if uint32(v13) <= uint32(v18) {
						goto l20
					}
				}
			l11:
				if uint32(v13) >= uint32(v15) {
					goto l21
				}
			l22:
				{
					t17 := int32(int8(m.memory[uint32(v10+v13)]))
					v1 = t17
					if v1 < i32(0) {
						goto l14
					}
					m.memory[uint32(v16+v13)] = byte(v1)
					t18 := v15
					v13 = v13 + i32(1)
					if t18 != v13 {
						goto l22
					}
				}
			l21:
				v13 = v15 + v12
				v10 = v15 + v11
				goto l23
			l14:
				v10 = v13 + v11
				v13 = v13 + v12
				if uint32(v13+i32(2)) < uint32(v5) {
					goto l24
				}
				v14 = i32(1)
			l23:
				store32(m.memory[uint32(v0):], uint32(v10))
				store32(m.memory[int64(uint32(v0))+4:], uint32(v14|(v7+i32(52))&i32(-256)))
				store32(m.memory[int64(uint32(v0))+8:], uint32(v13))
				goto l25
			l24:
				v15 = v10 + i32(1)
				{
					t19 := int32(load16(m.memory[uint32(v9+v1&i32(255)<<1+i32(-256)):]))
					v1 = t19
					if v1 == 0 {
						goto l26
					}
				l35:
					{
						v10 = v13 + i32(1)
						if uint32(v1&i32(0xffff)) < uint32(i32(2048)) {
							goto l27
						}
						m.memory[uint32(v4+v10)] = byte(int32(uint32(v1)>>6)&i32(63) | i32(128))
						v10 = v13 + i32(2)
						v16 = int32(uint32(v1&i32(61440))>>12) | i32(-32)
						v12 = i32(3)
						goto l28
					l27:
						v16 = int32(uint32(v1)>>6) | i32(-64)
						v12 = i32(2)
					l28:
						m.memory[uint32(v4+v13)] = byte(v16)
						m.memory[uint32(v4+v10)] = byte(v1&i32(63) | i32(128))
						v10 = v12 + v13
						if uint32(v15) < uint32(v3) {
							goto l29
						}
						store32(m.memory[uint32(v0):], uint32(v15))
						m.memory[int64(uint32(v0))+4] = byte(i32(0))
						store32(m.memory[int64(uint32(v0))+8:], uint32(v10))
						goto l25
					l29:
						if uint32(v10+i32(2)) >= uint32(v5) {
							store32(m.memory[uint32(v0):], uint32(v15))
							m.memory[int64(uint32(v0))+4] = byte(i32(1))
							store32(m.memory[int64(uint32(v0))+8:], uint32(v10))
							goto l25
						}
						v16 = v2 + v15
						v6 = v8 - v15
						v17 = v4 + v10
						v13 = i32(0)
					l34:
						{
							v11 = v15 + v13 + i32(1)
							t20 := int32(int8(m.memory[uint32(v16+v13)]))
							v1 = t20
							if v1 < i32(0) {
								goto l31
							}
							m.memory[uint32(v17+v13)] = byte(v1)
							v12 = v10 + v13 + i32(1)
							if uint32(v1) > uint32(i32(59)) {
								goto l32
							}
							if v6 != v13 {
								goto l33
							}
							store32(m.memory[uint32(v0):], uint32(v3))
							m.memory[int64(uint32(v0))+4] = byte(i32(0))
							store32(m.memory[int64(uint32(v0))+8:], uint32(v12))
							goto l25
						l33:
							t21 := v10
							v13 = v13 + i32(1)
							v1 = t21 + v13
							if uint32(v1+i32(2)) < uint32(v5) {
								goto l34
							}
						}
						store32(m.memory[uint32(v0):], uint32(v15+v13))
						m.memory[int64(uint32(v0))+4] = byte(i32(1))
						store32(m.memory[int64(uint32(v0))+8:], uint32(v1))
						goto l25
					l31:
						v13 = v10 + v13
						v15 = v11
						t22 := int32(load16(m.memory[uint32(v9+v1&i32(255)<<1+i32(-256)):]))
						v1 = t22
						if v1 != 0 {
							goto l35
						}
					}
				}
			l26:
				m.memory[int64(uint32(v0))+6] = byte(i32(0))
				store16(m.memory[int64(uint32(v0))+4:], uint16(i32(258)))
				store32(m.memory[uint32(v0):], uint32(v15))
				store32(m.memory[int64(uint32(v0))+8:], uint32(v13))
				goto l25
			l32:
				v10 = v2 + v11
				if uint32(v12) <= uint32(v5) {
					goto l36
				}
			}
			m.fn121(v12, v5, v5, i32(1146188))
			panic("unreachable")
		case 10:
			{
				{
					t23 := int32(m.memory[int64(uint32(v1))+7])
					if t23 != 0 {
						goto l37
					}
					t24 := int32(load16(m.memory[int64(uint32(v1))+4:]))
					v8 = t24
					v12 = i32(0)
					goto l38
				}
			l37:
				if uint32(v5) > uint32(i32(2)) {
					goto l39
				}
				m.memory[int64(uint32(v0))+4] = byte(i32(1))
				goto l40
			l39:
				{
					t25 := int32(load16(m.memory[int64(uint32(v1))+4:]))
					v13 = t25
					if uint32(v13) < uint32(i32(128)) {
						goto l41
					}
					if uint32(v13) < uint32(i32(2048)) {
						m.memory[int64(uint32(v4))+1] = byte(v13&i32(63) | i32(128))
						m.memory[uint32(v4)] = byte(int32(uint32(v13)>>6) | i32(192))
						v12 = i32(2)
						goto l43
					}
					m.memory[int64(uint32(v4))+2] = byte(v13&i32(63) | i32(128))
					m.memory[uint32(v4)] = byte(int32(uint32(v13)>>12) | i32(224))
					m.memory[int64(uint32(v4))+1] = byte(int32(uint32(v13)>>6)&i32(63) | i32(128))
					v12 = i32(3)
					goto l43
				}
			l41:
				m.memory[uint32(v4)] = byte(v13)
				v12 = i32(1)
			l43:
				v8 = i32(0)
				store16(m.memory[int64(uint32(v1))+4:], uint16(i32(0)))
				m.memory[int64(uint32(v1))+7] = byte(i32(0))
			l38:
				t26 := int32(m.memory[int64(uint32(v1))+2])
				v17 = t26
				t27 := int32(m.memory[int64(uint32(v1))+6])
				v19 = t27
				t28 := int32(m.memory[int64(uint32(v1))+3])
				v20 = t28
				v11 = i32(0)
			l104:
				v18 = v17 & i32(1)
				if v18 != 0 {
					goto l44
				}
				if v8&i32(0xffff) != 0 {
					goto l44
				}
				if v19&i32(1) != 0 {
					if uint32(v3) < uint32(v11) {
						m.fn121(v11, v3, v3, i32(1139632))
						panic("unreachable")
					}
					{
						if uint32(v5) < uint32(v12) {
							m.fn121(v12, v5, v5, i32(1139616))
							panic("unreachable")
						}
						v15 = int32(uint32(v3-v11) >> 1)
						if v15 == 0 {
							goto l44
						}
						v13 = i32(0)
						v10 = i32(0)
						v9 = v5 - v12
						if uint32(v9) < uint32(i32(4)) {
							goto l71
						}
						v14 = v4 + v12
						v13 = v15 + i32(-1)
						t40 := v13
						t41 := v15
						v25 = v2 + v11
						t42 := int32(load16(m.memory[uint32(v25+v13<<1):]))
						p43 := t41
						if t42&i32(252) == i32(216) {
							p43 = t40
						}
						v22 = p43
						v21 = v9 + i32(-3)
						v24 = i32(0)
						v23 = i32(0)
						v10 = i32(0)
					l89:
						{
							v13 = v9 - v23
							t44 := v13
							v15 = v22 - v10
							p45 := v15
							if uint32(v13) < uint32(v15) {
								p45 = t44
							}
							v26 = p45
							if v26 == 0 {
								goto l72
							}
							v27 = v14 + v23
							v16 = v25 + v10<<1
							v13 = i32(0)
						l74:
							{
								t46 := int32(load16(m.memory[uint32(v16):]))
								v15 = t46
								v15 = v15<<8 | int32(uint32(v15)>>8)
								if uint32(v15&i32(0xffff)) > uint32(i32(127)) {
									v10 = v13 + v10
									v13 = v13 + v23
									if uint32(v13) >= uint32(v21) {
										goto l71
									}
									v10 = v10 + i32(1)
								l88:
									{
										{
											v16 = (v15 + i32(10240)) & i32(0xffff)
											if uint32(v16) > uint32(i32(2047)) {
												if uint32(v15&i32(0xffff)) < uint32(i32(2048)) {
													if uint32(v13) < uint32(v9) {
														m.memory[uint32(v14+v13)] = byte(int32(uint32(v15)>>6) | i32(192))
														v26 = v13 + i32(1)
														if uint32(v26) >= uint32(v9) {
															m.fn33(v26, v9, i32(1139776))
															panic("unreachable")
														}
														v27 = i32(2)
														v16 = v15
														goto l80
													}
													m.fn33(v13, v9, i32(1139760))
													panic("unreachable")
												}
												if uint32(v13) >= uint32(v9) {
													m.fn33(v13, v9, i32(1139712))
													panic("unreachable")
												}
												m.memory[uint32(v14+v13)] = byte(int32(uint32(v15&i32(61440))>>12) | i32(224))
												v16 = v13 + i32(1)
												if uint32(v16) >= uint32(v9) {
													m.fn33(v16, v9, i32(1139728))
													panic("unreachable")
												}
												m.memory[uint32(v14+v16)] = byte(int32(uint32(v15)>>6)&i32(63) | i32(128))
												v26 = v13 + i32(2)
												if uint32(v26) >= uint32(v9) {
													m.fn33(v26, v9, i32(1139744))
													panic("unreachable")
												}
												v27 = i32(3)
												v16 = v15
												goto l80
											}
											if uint32(v16) > uint32(i32(1023)) {
												goto l53
											}
											if uint32(v10) >= uint32(v22) {
												goto l53
											}
											t48 := int32(load16(m.memory[uint32(v25+v10<<1):]))
											v16 = t48
											v16 = v16<<8 | int32(uint32(v16)>>8)
											if v16&i32(64512) != i32(56320) {
												goto l53
											}
											if uint32(v13) >= uint32(v9) {
												m.fn33(v13, v9, i32(1139648))
												panic("unreachable")
											}
											t49 := v14 + v13
											v15 = v15&i32(0xffff)<<10 + v16&i32(0xffff) + i32(-56613888)
											m.memory[uint32(t49)] = byte(int32(uint32(v15)>>18) | i32(240))
											v26 = v13 + i32(1)
											if uint32(v26) >= uint32(v9) {
												m.fn33(v26, v9, i32(1139664))
												panic("unreachable")
											}
											m.memory[uint32(v14+v26)] = byte(int32(uint32(v15)>>12)&i32(63) | i32(128))
											v26 = v13 + i32(2)
											if uint32(v26) >= uint32(v9) {
												m.fn33(v26, v9, i32(1139680))
												panic("unreachable")
											}
											m.memory[uint32(v14+v26)] = byte(int32(uint32(v15)>>6)&i32(63) | i32(128))
											v26 = v13 + i32(3)
											if uint32(v26) >= uint32(v9) {
												m.fn33(v26, v9, i32(1139696))
												panic("unreachable")
											}
											v10 = v10 + i32(1)
											v27 = i32(4)
											goto l80
										}
									l80:
										m.memory[uint32(v14+v26)] = byte(v16&i32(63) | i32(128))
										v13 = v27 + v13
										if uint32(v13) >= uint32(v21) {
											goto l71
										}
										if v10 == v22 {
											goto l71
										}
										if uint32(v10) >= uint32(v22) {
											m.fn3(i32(1139568), i32(30), i32(1139600))
											panic("unreachable")
										}
										v15 = v10 << 1
										v10 = v10 + i32(1)
										t50 := int32(load16(m.memory[uint32(v25+v15):]))
										v15 = t50
										v15 = v15<<8 | int32(uint32(v15)>>8)
										if uint32(v15&i32(0xffff)) > uint32(i32(127)) {
											goto l88
										}
									}
									m.memory[uint32(v14+v13)] = byte(v15)
									v23 = v13 + i32(1)
									goto l89
								}
								m.memory[uint32(v27+v13)] = byte(v15)
								v16 = v16 + i32(2)
								t47 := v26
								v13 = v13 + i32(1)
								if t47 != v13 {
									goto l74
								}
							}
							v24 = v26
						}
					l72:
						v13 = v24 + v23
						v10 = v24 + v10
						goto l71
					}
				}
				if uint32(v3) < uint32(v11) {
					m.fn121(v11, v3, v3, i32(1139632))
					panic("unreachable")
				}
				{
					if uint32(v5) < uint32(v12) {
						m.fn121(v12, v5, v5, i32(1139616))
						panic("unreachable")
					}
					v10 = int32(uint32(v3-v11) >> 1)
					if v10 == 0 {
						goto l44
					}
					v13 = i32(0)
					v16 = i32(0)
					v9 = v5 - v12
					if uint32(v9) < uint32(i32(4)) {
						goto l48
					}
					v14 = v4 + v12
					v13 = v10 + i32(-1)
					t29 := v13
					t30 := v10
					v21 = v2 + v11
					t31 := int32(load16(m.memory[uint32(v21+v13<<1):]))
					p32 := t30
					if t31&i32(64512) == i32(55296) {
						p32 = t29
					}
					v22 = p32
					v23 = v9 + i32(-3)
					v24 = i32(0)
					v25 = i32(0)
					v10 = i32(0)
				l68:
					{
						v13 = v9 - v25
						t33 := v13
						v15 = v22 - v10
						p34 := v15
						if uint32(v13) < uint32(v15) {
							p34 = t33
						}
						v26 = p34
						if v26 == 0 {
							goto l49
						}
						v27 = v14 + v25
						v16 = v21 + v10<<1
						v13 = i32(0)
					l51:
						{
							t35 := int32(load16(m.memory[uint32(v16):]))
							v15 = t35
							if uint32(v15) > uint32(i32(127)) {
								v16 = v13 + v10
								v13 = v13 + v25
								if uint32(v13) >= uint32(v23) {
									goto l48
								}
								v10 = v16 + i32(1)
							l67:
								{
									{
										v16 = (v15 + i32(10240)) & i32(0xffff)
										if uint32(v16) > uint32(i32(2047)) {
											if uint32(v15&i32(0xffff)) < uint32(i32(2048)) {
												if uint32(v13) < uint32(v9) {
													m.memory[uint32(v14+v13)] = byte(int32(uint32(v15)>>6) | i32(192))
													v27 = v13 + i32(1)
													if uint32(v27) >= uint32(v9) {
														m.fn33(v27, v9, i32(1139776))
														panic("unreachable")
													}
													v25 = i32(2)
													goto l63
												}
												m.fn33(v13, v9, i32(1139760))
												panic("unreachable")
											}
											if uint32(v13) >= uint32(v9) {
												m.fn33(v13, v9, i32(1139712))
												panic("unreachable")
											}
											m.memory[uint32(v14+v13)] = byte(int32(uint32(v15&i32(61440))>>12) | i32(224))
											v16 = v13 + i32(1)
											if uint32(v16) >= uint32(v9) {
												m.fn33(v16, v9, i32(1139728))
												panic("unreachable")
											}
											m.memory[uint32(v14+v16)] = byte(int32(uint32(v15)>>6)&i32(63) | i32(128))
											v27 = v13 + i32(2)
											if uint32(v27) >= uint32(v9) {
												m.fn33(v27, v9, i32(1139744))
												panic("unreachable")
											}
											v25 = i32(3)
											goto l63
										}
										if uint32(v16) > uint32(i32(1023)) {
											goto l53
										}
										if uint32(v10) >= uint32(v22) {
											goto l53
										}
										t37 := int32(load16(m.memory[uint32(v21+v10<<1):]))
										v26 = t37
										if v26&i32(64512) != i32(56320) {
											goto l53
										}
										if uint32(v13) >= uint32(v9) {
											m.fn33(v13, v9, i32(1139648))
											panic("unreachable")
										}
										t38 := v14 + v13
										v15 = v15&i32(0xffff)<<10 + v26 + i32(-56613888)
										m.memory[uint32(t38)] = byte(int32(uint32(v15)>>18) | i32(240))
										v16 = v13 + i32(1)
										if uint32(v16) >= uint32(v9) {
											m.fn33(v16, v9, i32(1139664))
											panic("unreachable")
										}
										m.memory[uint32(v14+v16)] = byte(int32(uint32(v15)>>12)&i32(63) | i32(128))
										v16 = v13 + i32(2)
										if uint32(v16) >= uint32(v9) {
											m.fn33(v16, v9, i32(1139680))
											panic("unreachable")
										}
										m.memory[uint32(v14+v16)] = byte(int32(uint32(v15)>>6)&i32(63) | i32(128))
										v27 = v13 + i32(3)
										if uint32(v27) >= uint32(v9) {
											m.fn33(v27, v9, i32(1139696))
											panic("unreachable")
										}
										v16 = v10 + i32(1)
										v25 = i32(4)
										goto l58
									}
								l63:
									v26 = v15
									v16 = v10
								l58:
									m.memory[uint32(v14+v27)] = byte(v26&i32(63) | i32(128))
									v13 = v25 + v13
									if uint32(v13) >= uint32(v23) {
										goto l48
									}
									if v16 == v22 {
										goto l48
									}
									if uint32(v16) >= uint32(v22) {
										m.fn3(i32(1139568), i32(30), i32(1139600))
										panic("unreachable")
									}
									v10 = v16 + i32(1)
									t39 := int32(load16(m.memory[uint32(v21+v16<<1):]))
									v15 = t39
									if uint32(v15) > uint32(i32(127)) {
										goto l67
									}
								}
								m.memory[uint32(v14+v13)] = byte(v15)
								v25 = v13 + i32(1)
								goto l68
							}
							m.memory[uint32(v27+v13)] = byte(v15)
							v16 = v16 + i32(2)
							t36 := v26
							v13 = v13 + i32(1)
							if t36 != v13 {
								goto l51
							}
						}
						v24 = v26
					}
				l49:
					v13 = v24 + v25
					v16 = v24 + v10
					goto l48
				}
			l53:
				m.memory[int64(uint32(v0))+6] = byte(i32(0))
				store16(m.memory[int64(uint32(v0))+4:], uint16(i32(514)))
				v12 = v13 + v12
				v13 = v10<<1 + v11
				goto l90
			l71:
				v12 = v13 + v12
				v13 = v10<<1 + v11
				goto l91
			l48:
				v12 = v13 + v12
				v13 = v16<<1 + v11
				goto l91
			l44:
				v13 = v11
			l91:
				if uint32(v13) < uint32(v3) {
					v10 = v12 + i32(3)
					if uint32(v10) < uint32(v5) {
						v11 = v13 + i32(1)
						t53 := int32(m.memory[uint32(v2+v13)])
						v13 = t53
						{
							if v18 == 0 {
								m.memory[int64(uint32(v1))+3] = byte(v13)
								m.memory[int64(uint32(v1))+2] = byte(i32(1))
								v20 = v13
								v17 = v17 ^ i32(1)
								goto l104
							}
							m.memory[int64(uint32(v1))+2] = byte(i32(0))
							v13 = v20<<8 | v13&i32(255)
							p54 := v13<<8 | v20&i32(255)
							if v19&i32(1) != 0 {
								p54 = v13
							}
							v13 = p54
							v15 = v13 & i32(64512)
							if v15 == i32(55296) {
								store16(m.memory[int64(uint32(v1))+4:], uint16(v13))
								if v8&i32(0xffff) != 0 {
									m.memory[int64(uint32(v0))+6] = byte(i32(2))
									store16(m.memory[int64(uint32(v0))+4:], uint16(i32(514)))
									goto l103
								}
								v8 = v13
								v17 = v17 ^ i32(1)
								goto l104
							}
							if v15 != i32(56320) {
								goto l101
							}
							v15 = v8 & i32(0xffff)
							if v15 != 0 {
								v8 = i32(0)
								store16(m.memory[int64(uint32(v1))+4:], uint16(i32(0)))
								v10 = v4 + v12
								m.memory[uint32(v10+i32(3))] = byte(v13&i32(63) | i32(128))
								t55 := v10
								v13 = v15<<10 + v13&i32(0xffff) + i32(-56613888)
								m.memory[uint32(t55)] = byte(int32(uint32(v13)>>18) | i32(240))
								m.memory[uint32(v10+i32(2))] = byte(int32(uint32(v13)>>6)&i32(63) | i32(128))
								m.memory[uint32(v10+i32(1))] = byte(int32(uint32(v13)>>12)&i32(63) | i32(128))
								v12 = v12 + i32(4)
								v17 = v17 ^ i32(1)
								goto l104
							}
							m.memory[int64(uint32(v0))+6] = byte(i32(0))
							store16(m.memory[int64(uint32(v0))+4:], uint16(i32(514)))
							goto l103
						}
					l101:
						if v8&i32(0xffff) == 0 {
							v15 = v4 + v12
							v16 = v13 & i32(0xffff)
							if uint32(v16) < uint32(i32(128)) {
								m.memory[uint32(v15)] = byte(v13)
								v12 = v12 + i32(1)
								v8 = i32(0)
								v17 = v17 ^ i32(1)
								goto l104
							}
							if uint32(v16) < uint32(i32(2048)) {
								m.memory[uint32(v15+i32(1))] = byte(v13&i32(63) | i32(128))
								m.memory[uint32(v15)] = byte(int32(uint32(v13)>>6) | i32(192))
								v12 = v12 + i32(2)
								v8 = i32(0)
								v17 = v17 ^ i32(1)
								goto l104
							}
							m.memory[uint32(v15+i32(2))] = byte(v13&i32(63) | i32(128))
							m.memory[uint32(v15+i32(1))] = byte(int32(uint32(v13)>>6)&i32(63) | i32(128))
							m.memory[uint32(v15)] = byte(int32(uint32(v13&i32(61440))>>12) | i32(224))
							v8 = i32(0)
							v12 = v10
							v17 = v17 ^ i32(1)
							goto l104
						}
						m.memory[int64(uint32(v1))+7] = byte(i32(1))
						store16(m.memory[int64(uint32(v1))+4:], uint16(v13))
						m.memory[int64(uint32(v0))+6] = byte(i32(2))
						store16(m.memory[int64(uint32(v0))+4:], uint16(i32(514)))
					l103:
						v13 = v11
						goto l90
					}
					m.memory[int64(uint32(v0))+4] = byte(i32(1))
					goto l90
				}
				{
					if v6 == 0 {
						goto l93
					}
					t51 := v17
					var p52 int32
					if v8&i32(0xffff) != i32(0) {
						p52 = 1
					}
					if (t51|p52)&i32(1) != 0 {
						if uint32(v12+i32(2)) < uint32(v5) {
							if v8&i32(0xffff) != 0 {
								store16(m.memory[int64(uint32(v1))+4:], uint16(i32(0)))
								if v17&i32(1) == 0 {
									m.memory[int64(uint32(v0))+6] = byte(i32(0))
									store16(m.memory[int64(uint32(v0))+4:], uint16(i32(514)))
									goto l90
								}
								m.memory[int64(uint32(v0))+6] = byte(i32(0))
								store16(m.memory[int64(uint32(v0))+4:], uint16(i32(770)))
								m.memory[int64(uint32(v1))+2] = byte(i32(0))
								goto l90
							}
							m.memory[int64(uint32(v0))+6] = byte(i32(0))
							store16(m.memory[int64(uint32(v0))+4:], uint16(i32(258)))
							m.memory[int64(uint32(v1))+2] = byte(i32(0))
							goto l90
						}
						m.memory[int64(uint32(v0))+4] = byte(i32(1))
						goto l40
					}
				}
			l93:
				m.memory[int64(uint32(v0))+4] = byte(i32(0))
				goto l90
			}
		l40:
			v13 = i32(0)
			v12 = i32(0)
		l90:
			store32(m.memory[int64(uint32(v0))+8:], uint32(v12))
			store32(m.memory[uint32(v0):], uint32(v13))
			goto l25
		case 9:
			if v3 != 0 {
				goto l109
			}
			v16 = i32(0)
			v3 = i32(0)
			v13 = i32(0)
			goto l110
		l109:
			v13 = i32(0)
			v16 = i32(1)
			v1 = i32(0)
		l113:
			if uint32(v13+i32(2)) < uint32(v5) {
				goto l111
			}
			v3 = v1
			goto l110
		l111:
			v15 = i32(1)
			v10 = v1 + i32(1)
			{
				t56 := int32(int8(m.memory[uint32(v2+v1)]))
				v1 = t56
				if v1 > i32(-1) {
					goto l112
				}
				v15 = v4 + v13
				m.memory[uint32(v15+i32(2))] = byte(v1 & i32(191))
				m.memory[uint32(v15+i32(1))] = byte(int32(uint32(v1&i32(192))>>6) | i32(156))
				v1 = i32(239)
				v15 = i32(3)
			}
		l112:
			m.memory[uint32(v4+v13)] = byte(v1)
			v13 = v15 + v13
			v1 = v10
			if v3 != v10 {
				goto l113
			}
			v16 = i32(0)
		l110:
			store32(m.memory[int64(uint32(v0))+8:], uint32(v13))
			store32(m.memory[uint32(v0):], uint32(v3))
			m.memory[int64(uint32(v0))+4] = byte(v16)
			goto l25
		case 8:
			{
				if v3 == 0 {
					goto l114
				}
				t57 := int32(m.memory[int64(uint32(v1))+1])
				if t57&i32(1) == 0 {
					if uint32(v5) < uint32(i32(3)) {
						store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
						store32(m.memory[uint32(v0):], uint32(i32(0)))
						m.memory[int64(uint32(v0))+4] = byte(i32(1))
						goto l25
					}
					m.memory[int64(uint32(v0))+6] = byte(i32(0))
					store16(m.memory[int64(uint32(v0))+4:], uint16(i32(258)))
					m.memory[int64(uint32(v1))+1] = byte(i32(1))
					store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
					store32(m.memory[uint32(v0):], uint32(i32(1)))
					goto l25
				}
			}
		l114:
			store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
			store32(m.memory[uint32(v0):], uint32(v3))
			m.memory[int64(uint32(v0))+4] = byte(i32(0))
			goto l25
		case 7:
			store32(m.memory[int64(uint32(v7))+56:], uint32(v5))
			store32(m.memory[int64(uint32(v7))+52:], uint32(v4))
			v15 = i32(0)
			v17 = i32(0)
			{
				t58 := int32(m.memory[int64(uint32(v1))+1])
				if t58 == 0 {
					goto l117
				}
				m.memory[int64(uint32(v1))+1] = byte(i32(0))
				if v3 == 0 {
					if v6 != 0 {
						m.memory[int64(uint32(v0))+6] = byte(i32(0))
						store16(m.memory[int64(uint32(v0))+4:], uint16(i32(258)))
						store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
						store32(m.memory[uint32(v0):], uint32(i32(0)))
						goto l25
					}
					store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
					store32(m.memory[uint32(v0):], uint32(i32(0)))
					m.memory[int64(uint32(v0))+4] = byte(i32(0))
					goto l25
				}
				if uint32(v5) > uint32(i32(2)) {
					goto l119
				}
				store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
				store32(m.memory[uint32(v0):], uint32(i32(0)))
				m.memory[int64(uint32(v0))+4] = byte(i32(1))
				goto l25
			l119:
				t59 := int32(int8(m.memory[uint32(v2)]))
				v13 = t59
				{
					t60 := int32(m.memory[int64(uint32(v1))+2])
					v10 = t60
					if uint32(v10&i32(255)) > uint32(i32(31)) {
						v15 = v10 + i32(-32)
						v16 = v13 + i32(95)
						if uint32(v16&i32(255)) < uint32(i32(94)) {
							t61 := v15 & i32(255) * i32(94)
							v11 = v16 & i32(255)
							v15 = t61 + v11
							v12 = v15 + i32(-1410)
							if uint32(v12) < uint32(i32(2350)) {
								v17 = i32(1)
								t76 := int32(load16(m.memory[int64(uint32(v12<<1))+1219452:]))
								t77 := v4
								v13 = t76
								m.memory[int64(uint32(t77))+2] = byte(v13&i32(63) | i32(128))
								m.memory[uint32(v4)] = byte(int32(uint32(v13)>>12) | i32(224))
								m.memory[int64(uint32(v4))+1] = byte(int32(uint32(v13)>>6)&i32(63) | i32(128))
								v15 = i32(3)
								goto l117
							}
							if uint32(v15) < uint32(i32(165)) {
								v17 = i32(1)
								{
									t62 := int32(load16(m.memory[int64(uint32(v15<<1))+1226988:]))
									v13 = t62
									if uint32(v13) < uint32(i32(2048)) {
										m.memory[uint32(v4)] = byte(int32(uint32(v13)>>6) | i32(192))
										m.memory[int64(uint32(v4))+1] = byte(v13&i32(63) | i32(128))
										v17 = i32(1)
										store32(m.memory[int64(uint32(v7))+60:], uint32(i32(1)))
										v15 = i32(2)
										goto l117
									}
									m.memory[int64(uint32(v4))+2] = byte(v13&i32(63) | i32(128))
									m.memory[uint32(v4)] = byte(int32(uint32(v13)>>12) | i32(224))
									m.memory[int64(uint32(v4))+1] = byte(int32(uint32(v13)>>6)&i32(63) | i32(128))
									goto l142
								}
							}
							v12 = v15 + i32(-3854)
							if uint32(v12) < uint32(i32(4888)) {
								v17 = i32(1)
								t63 := int32(load16(m.memory[int64(uint32(v12<<1))+1209488:]))
								t64 := v4
								v13 = t63
								m.memory[int64(uint32(t64))+2] = byte(v13&i32(63) | i32(128))
								m.memory[uint32(v4)] = byte(int32(uint32(v13)>>12) | i32(224))
								m.memory[int64(uint32(v4))+1] = byte(int32(uint32(v13)>>6)&i32(63) | i32(128))
								goto l142
							}
							v12 = v10 & i32(255)
							if v12 != i32(39) {
								goto l134
							}
							if uint32(v16&i32(255)) < uint32(i32(15)) {
								if v13&i32(-83) != i32(-91) {
									v17 = i32(1)
									t65 := int32(load16(m.memory[int64(uint32(v11<<1))+1234848:]))
									t66 := v4
									v13 = t65
									m.memory[int64(uint32(t66))+1] = byte(v13&i32(63) | i32(128))
									m.memory[uint32(v4)] = byte(int32(uint32(v13)>>6) | i32(192))
									goto l144
								}
								m.memory[int64(uint32(v0))+6] = byte(i32(0))
								store16(m.memory[int64(uint32(v0))+4:], uint16(i32(514)))
								store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
								store32(m.memory[uint32(v0):], uint32(i32(1)))
								goto l25
							}
						l134:
							if v12 != i32(40) {
								goto l136
							}
							if uint32(v16&i32(255)) < uint32(i32(16)) {
								v17 = i32(1)
								t67 := int32(load16(m.memory[int64(uint32(v11<<1))+1234816:]))
								t68 := v4
								v13 = t67
								m.memory[int64(uint32(t68))+1] = byte(v13&i32(63) | i32(128))
								m.memory[uint32(v4)] = byte(int32(uint32(v13)>>6) | i32(192))
								goto l144
							}
						l136:
							if v10&i32(255) != i32(37) {
								goto l138
							}
							if uint32(v16&i32(255)) < uint32(i32(68)) {
								v17 = i32(1)
								t69 := int32(load16(m.memory[int64(uint32(v11<<1))+1146470:]))
								t70 := v4
								v13 = t69
								m.memory[int64(uint32(t70))+2] = byte(v13&i32(63) | i32(128))
								m.memory[uint32(v4)] = byte(int32(uint32(v13)>>12) | i32(224))
								m.memory[int64(uint32(v4))+1] = byte(int32(uint32(v13)>>6)&i32(63) | i32(128))
								goto l142
							}
						l138:
							v13 = v15 + i32(-188)
							if uint32(v13) < uint32(i32(927)) {
								m.fn897(v7+i32(32), i32(1242038), i32(77), v13)
								t71 := int32(load32(m.memory[int64(uint32(v7))+36:]))
								v10 = t71
								{
									{
										t72 := int32(load32(m.memory[int64(uint32(v7))+32:]))
										if t72 != i32(1) {
											goto l145
										}
										v10 = v10 + i32(-1)
										if uint32(v10) >= uint32(i32(77)) {
											m.fn33(v10, i32(77), i32(1227336))
											panic("unreachable")
										}
										v10 = v10 << 1
										t73 := int32(load16(m.memory[int64(uint32(v10))+1263160:]))
										t74 := int32(load16(m.memory[int64(uint32(v10))+1242038:]))
										v13 = t73 + v13 - t74
										goto l147
									}
								l145:
									if uint32(v10) > uint32(i32(76)) {
										m.fn33(v10, i32(77), i32(1227320))
										panic("unreachable")
									}
									t75 := int32(load16(m.memory[int64(uint32(v10<<1))+1263160:]))
									v13 = t75
								}
							l147:
								v10 = v13 & i32(0xffff)
								if uint32(v10) < uint32(i32(128)) {
									m.memory[int64(uint32(v0))+6] = byte(i32(0))
									store16(m.memory[int64(uint32(v0))+4:], uint16(i32(514)))
									store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
									store32(m.memory[uint32(v0):], uint32(i32(1)))
									goto l25
								}
								if uint32(v10) < uint32(i32(2048)) {
									m.memory[int64(uint32(v4))+1] = byte(v13&i32(63) | i32(128))
									m.memory[uint32(v4)] = byte(int32(uint32(v13)>>6) | i32(192))
									v15 = i32(2)
									goto l152
								}
								m.memory[int64(uint32(v4))+2] = byte(v13&i32(63) | i32(128))
								m.memory[int64(uint32(v4))+1] = byte(int32(uint32(v13)>>6)&i32(63) | i32(128))
								m.memory[uint32(v4)] = byte(int32(uint32(v13&i32(61440))>>12) | i32(224))
								goto l151
							}
							m.memory[int64(uint32(v0))+6] = byte(i32(0))
							store16(m.memory[int64(uint32(v0))+4:], uint16(i32(514)))
							store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
							store32(m.memory[uint32(v0):], uint32(i32(1)))
							goto l25
						}
						if uint32((v13+i32(127))&i32(255)) < uint32(i32(32)) {
							v10 = v13 + i32(-77)
							goto l129
						}
						if uint32((v13+i32(-97))&i32(255)) < uint32(i32(26)) {
							v10 = v13 + i32(-71)
							goto l129
						}
						v10 = v13 + i32(-65)
						if uint32(v10&i32(255)) < uint32(i32(26)) {
							goto l129
						}
						m.memory[int64(uint32(v0))+4] = byte(i32(2))
						if v13 > i32(-1) {
							store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
							store32(m.memory[uint32(v0):], uint32(i32(0)))
							store16(m.memory[int64(uint32(v0))+5:], uint16(i32(1)))
							goto l25
						}
						store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
						store32(m.memory[uint32(v0):], uint32(i32(1)))
						store16(m.memory[int64(uint32(v0))+5:], uint16(i32(2)))
						goto l25
					}
					if uint32((v13+i32(127))&i32(255)) < uint32(i32(126)) {
						v15 = v13 + i32(-77)
						goto l124
					}
					if uint32((v13+i32(-97))&i32(255)) < uint32(i32(26)) {
						v15 = v13 + i32(-71)
						goto l124
					}
					v15 = v13 + i32(-65)
					if uint32(v15&i32(255)) < uint32(i32(26)) {
						goto l124
					}
					m.memory[int64(uint32(v0))+4] = byte(i32(2))
					if v13 > i32(-1) {
						store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
						store32(m.memory[uint32(v0):], uint32(i32(0)))
						store16(m.memory[int64(uint32(v0))+5:], uint16(i32(1)))
						goto l25
					}
					store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
					store32(m.memory[uint32(v0):], uint32(i32(1)))
					store16(m.memory[int64(uint32(v0))+5:], uint16(i32(2)))
					goto l25
				}
			l144:
				v15 = i32(2)
				store32(m.memory[int64(uint32(v7))+60:], uint32(i32(2)))
				goto l117
			l142:
				v15 = i32(3)
				store32(m.memory[int64(uint32(v7))+60:], uint32(i32(3)))
				goto l117
			l129:
				v10 = v15&i32(255)*i32(84) + v10&i32(255)
				if uint32(v10) < uint32(i32(3126)) {
					m.fn897(v7+i32(40), i32(1251496), i32(535), v10)
					t78 := int32(load32(m.memory[int64(uint32(v7))+44:]))
					v13 = t78
					{
						{
							t79 := int32(load32(m.memory[int64(uint32(v7))+40:]))
							if t79 != i32(1) {
								goto l155
							}
							v13 = v13 + i32(-1)
							if uint32(v13) >= uint32(i32(535)) {
								m.fn33(i32(-1), i32(535), i32(1227336))
								panic("unreachable")
							}
							t80 := v10
							v13 = v13 << 1
							t81 := int32(load16(m.memory[int64(uint32(v13))+1251496:]))
							t82 := int32(load16(m.memory[int64(uint32(v13))+1244436:]))
							v13 = t80 - t81 + t82
							v15 = v13 & i32(0xffff)
							v10 = int32(uint32(v15) >> 12)
							v15 = int32(uint32(v15) >> 6)
							goto l157
						}
					l155:
						if uint32(v13) > uint32(i32(534)) {
							m.fn33(i32(535), i32(535), i32(1227320))
							panic("unreachable")
						}
						t83 := int32(load16(m.memory[int64(uint32(v13<<1))+1244436:]))
						v13 = t83
						v10 = int32(uint32(v13) >> 12)
						v15 = int32(uint32(v13) >> 6)
					}
				l157:
					m.memory[uint32(v4)] = byte(v10 | i32(224))
					m.memory[int64(uint32(v4))+2] = byte(v13&i32(63) | i32(128))
					m.memory[int64(uint32(v4))+1] = byte(v15&i32(63) | i32(128))
					goto l151
				}
				m.memory[int64(uint32(v0))+4] = byte(i32(2))
				if v13 > i32(-1) {
					store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
					store32(m.memory[uint32(v0):], uint32(i32(0)))
					store16(m.memory[int64(uint32(v0))+5:], uint16(i32(1)))
					goto l25
				}
				store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
				store32(m.memory[uint32(v0):], uint32(i32(1)))
				store16(m.memory[int64(uint32(v0))+5:], uint16(i32(2)))
				goto l25
			l124:
				{
					{
						v12 = v10&i32(255)*i32(178) + v15&i32(255)
						v13 = v12 & i32(0xffff)
						p84 := i32(539)
						if uint32(v13) < uint32(i32(2868)) {
							p84 = i32(0)
						}
						v10 = p84
						t85 := v10
						v10 = v10 + i32(270)
						t86 := int32(load16(m.memory[int64(uint32(v10<<1))+1245506:]))
						p87 := v10
						if uint32(t86) > uint32(v13) {
							p87 = t85
						}
						v10 = p87
						t88 := v10
						v10 = v10 + i32(135)
						t89 := int32(load16(m.memory[int64(uint32(v10<<1))+1245506:]))
						p90 := v10
						if uint32(t89) > uint32(v13) {
							p90 = t88
						}
						v10 = p90
						t91 := v10
						v10 = v10 + i32(67)
						t92 := int32(load16(m.memory[int64(uint32(v10<<1))+1245506:]))
						p93 := v10
						if uint32(t92) > uint32(v13) {
							p93 = t91
						}
						v10 = p93
						t94 := v10
						v10 = v10 + i32(34)
						t95 := int32(load16(m.memory[int64(uint32(v10<<1))+1245506:]))
						p96 := v10
						if uint32(t95) > uint32(v13) {
							p96 = t94
						}
						v10 = p96
						t97 := v10
						v10 = v10 + i32(17)
						t98 := int32(load16(m.memory[int64(uint32(v10<<1))+1245506:]))
						p99 := v10
						if uint32(t98) > uint32(v13) {
							p99 = t97
						}
						v10 = p99
						t100 := v10
						v10 = v10 + i32(8)
						t101 := int32(load16(m.memory[int64(uint32(v10<<1))+1245506:]))
						p102 := v10
						if uint32(t101) > uint32(v13) {
							p102 = t100
						}
						v10 = p102
						t103 := v10
						v10 = v10 + i32(4)
						t104 := int32(load16(m.memory[int64(uint32(v10<<1))+1245506:]))
						p105 := v10
						if uint32(t104) > uint32(v13) {
							p105 = t103
						}
						v10 = p105
						t106 := v10
						v10 = v10 + i32(2)
						t107 := int32(load16(m.memory[int64(uint32(v10<<1))+1245506:]))
						p108 := v10
						if uint32(t107) > uint32(v13) {
							p108 = t106
						}
						v10 = p108
						t109 := v10
						v10 = v10 + i32(1)
						t110 := int32(load16(m.memory[int64(uint32(v10<<1))+1245506:]))
						p111 := v10
						if uint32(t110) > uint32(v13) {
							p111 = t109
						}
						v10 = p111
						t112 := v10
						v10 = v10 + i32(1)
						t113 := int32(load16(m.memory[int64(uint32(v10<<1))+1245506:]))
						p114 := v10
						if uint32(t113) > uint32(v13) {
							p114 = t112
						}
						v10 = p114
						v15 = v10 << 1
						t115 := int32(load16(m.memory[int64(uint32(v15))+1245506:]))
						v16 = t115
						if v16 == v13 {
							goto l159
						}
						{
							t116 := v10
							var p117 int32
							if uint32(v16) >= uint32(v13) {
								p117 = 1
							}
							v13 = t116 - p117
							if uint32(v13) >= uint32(i32(1079)) {
								m.fn33(i32(-1), i32(1079), i32(1227336))
								panic("unreachable")
							}
							t118 := v12
							v13 = v13 << 1
							t119 := int32(load16(m.memory[int64(uint32(v13))+1245506:]))
							t120 := int32(load16(m.memory[int64(uint32(v13))+1242212:]))
							v13 = t118 - t119 + t120
							v15 = v13 & i32(0xffff)
							v10 = int32(uint32(v15) >> 12)
							v15 = int32(uint32(v15) >> 6)
							goto l161
						}
					}
				l159:
					t121 := int32(load16(m.memory[int64(uint32(v15))+1242212:]))
					v13 = t121
					v10 = int32(uint32(v13) >> 12)
					v15 = int32(uint32(v13) >> 6)
				}
			l161:
				m.memory[uint32(v4)] = byte(v10 | i32(224))
				m.memory[int64(uint32(v4))+2] = byte(v13&i32(63) | i32(128))
				m.memory[int64(uint32(v4))+1] = byte(v15&i32(63) | i32(128))
			l151:
				v15 = i32(3)
			l152:
				store32(m.memory[int64(uint32(v7))+60:], uint32(v15))
				v17 = i32(1)
			}
		l117:
			v26 = v3 + i32(-1)
		l228:
			{
				v13 = v5 - v15
				t122 := v13
				v10 = v3 - v17
				t123 := v10
				var p124 int32
				if uint32(v13) < uint32(v10) {
					p124 = 1
				}
				v14 = p124
				p125 := t123
				if v14 != 0 {
					p125 = t122
				}
				v12 = p125
				v13 = i32(0)
				{
					{
						v11 = v4 + v15
						t126 := v11
						v16 = v2 + v17
						if (t126^v16)&i32(3) != 0 {
							goto l162
						}
						v13 = i32(0)
						v18 = (i32(0) - v16) & i32(3)
						if uint32(v18|i32(8)) > uint32(v12) {
							goto l162
						}
						{
							if v18 != 0 {
								goto l163
							}
							v13 = i32(0)
							goto l164
						l163:
							v13 = i32(0)
							t127 := int32(int8(m.memory[uint32(v16)]))
							v10 = t127
							if v10 < i32(0) {
								goto l165
							}
							m.memory[uint32(v11)] = byte(v10)
							v13 = i32(1)
							if v18 == i32(1) {
								goto l164
							}
							{
								t128 := int32(int8(m.memory[int64(uint32(v16))+1]))
								v10 = t128
								if v10 >= i32(0) {
									goto l166
								}
								v13 = i32(1)
								goto l165
							}
						l166:
							m.memory[int64(uint32(v11))+1] = byte(v10)
							v13 = i32(2)
							if v18 == i32(2) {
								goto l164
							}
							{
								t129 := int32(int8(m.memory[int64(uint32(v16))+2]))
								v10 = t129
								if v10 >= i32(0) {
									goto l167
								}
								v13 = i32(2)
								goto l165
							}
						l167:
							m.memory[int64(uint32(v11))+2] = byte(v10)
							v13 = i32(3)
						}
					l164:
						v8 = v12 + i32(-8)
					l171:
						{
							v18 = v16 + v13
							t130 := int32(load32(m.memory[uint32(v18):]))
							v10 = t130
							v9 = v11 + v13
							t131 := int32(load32(m.memory[uint32(v18+i32(4)):]))
							t132 := v9 + i32(4)
							v18 = t131
							store32(m.memory[uint32(t132):], uint32(v18))
							store32(m.memory[uint32(v9):], uint32(v10))
							{
								v18 = v18 & i32(-2139062144)
								t133 := v18
								v10 = v10 & i32(-2139062144)
								if t133|v10 == 0 {
									goto l168
								}
								if v10 != 0 {
									goto l169
								}
								v10 = int32(uint32(int32(bits.TrailingZeros32(uint32(v18))))>>3) + i32(4)
								goto l170
							l169:
								v10 = int32(uint32(int32(bits.TrailingZeros32(uint32(v10)))) >> 3)
							l170:
								t134 := v16
								v13 = v10 + v13
								t135 := int32(m.memory[uint32(t134+v13)])
								v10 = t135
								goto l165
							}
						l168:
							v13 = v13 + i32(8)
							if uint32(v13) <= uint32(v8) {
								goto l171
							}
						}
					}
				l162:
					if uint32(v13) >= uint32(v12) {
						goto l172
					}
				l173:
					{
						t136 := int32(int8(m.memory[uint32(v16+v13)]))
						v10 = t136
						if v10 < i32(0) {
							goto l165
						}
						m.memory[uint32(v11+v13)] = byte(v10)
						t137 := v12
						v13 = v13 + i32(1)
						if t137 != v13 {
							goto l173
						}
					}
				l172:
					v15 = v12 + v15
					v13 = v12 + v17
					goto l174
				l165:
					t138 := v7
					v15 = v13 + v15
					store32(m.memory[int64(uint32(t138))+60:], uint32(v15))
					v13 = v13 + v17
					if uint32(v15+i32(2)) < uint32(v5) {
						goto l175
					}
					v14 = i32(1)
				}
			l174:
				store32(m.memory[int64(uint32(v0))+8:], uint32(v15))
				store32(m.memory[uint32(v0):], uint32(v13))
				store32(m.memory[int64(uint32(v0))+4:], uint32(v14|(v7+i32(52))&i32(-256)))
				goto l25
			l175:
				v5 = v13 + i32(1)
				v12 = v10 + i32(127)
				if uint32(v12&i32(255)) > uint32(i32(125)) {
					goto l176
				}
			l227:
				{
					if uint32(v5) < uint32(v3) {
						goto l177
					}
					if v6 != 0 {
						m.memory[int64(uint32(v0))+6] = byte(i32(0))
						store16(m.memory[int64(uint32(v0))+4:], uint16(i32(258)))
						goto l179
					}
					m.memory[int64(uint32(v1))+2] = byte(v12)
					m.memory[int64(uint32(v1))+1] = byte(i32(1))
					m.memory[int64(uint32(v0))+4] = byte(i32(0))
					goto l179
				l177:
					v16 = v5 + i32(1)
					t139 := int32(int8(m.memory[uint32(v2+v5)]))
					v13 = t139
					{
						v12 = v12 & i32(255)
						if uint32(v12) > uint32(i32(31)) {
							v10 = v10 + i32(95)
							v11 = (v13 + i32(95)) & i32(255)
							if uint32(v11) < uint32(i32(94)) {
								v10 = v10&i32(255)*i32(94) + v11
								v17 = v10 + i32(-1410)
								if uint32(v17) < uint32(i32(2350)) {
									t167 := int32(load16(m.memory[int64(uint32(v17<<1))+1219452:]))
									t168 := v4 + v15
									v13 = t167
									m.memory[uint32(t168)] = byte(int32(uint32(v13)>>12) | i32(224))
									t169 := int32(load32(m.memory[int64(uint32(v7))+52:]))
									v4 = t169
									t170 := int32(load32(m.memory[int64(uint32(v7))+60:]))
									t171 := v4
									v10 = t170
									v15 = t171 + v10
									m.memory[uint32(v15+i32(2))] = byte(v13&i32(63) | i32(128))
									m.memory[uint32(v15+i32(1))] = byte(int32(uint32(v13)>>6)&i32(63) | i32(128))
									v15 = v10 + i32(3)
									goto l202
								}
								if uint32(v10) < uint32(i32(165)) {
									v11 = v15 + i32(1)
									v12 = v4 + v15
									{
										t140 := int32(load16(m.memory[int64(uint32(v10<<1))+1226988:]))
										v13 = t140
										if uint32(v13) < uint32(i32(2048)) {
											m.memory[uint32(v4+v11)] = byte(v13&i32(63) | i32(128))
											m.memory[uint32(v12)] = byte(int32(uint32(v13)>>6) | i32(192))
											v15 = v15 + i32(2)
											goto l202
										}
										m.memory[uint32(v12)] = byte(int32(uint32(v13)>>12) | i32(224))
										m.memory[uint32(v12+i32(2))] = byte(v13&i32(63) | i32(128))
										m.memory[uint32(v4+v11)] = byte(int32(uint32(v13)>>6)&i32(63) | i32(128))
										v15 = v15 + i32(3)
										goto l202
									}
								}
								v17 = v10 + i32(-3854)
								if uint32(v17) < uint32(i32(4888)) {
									t141 := int32(load16(m.memory[int64(uint32(v17<<1))+1209488:]))
									t142 := v4 + v15
									v13 = t141
									m.memory[uint32(t142)] = byte(int32(uint32(v13)>>12) | i32(224))
									t143 := int32(load32(m.memory[int64(uint32(v7))+52:]))
									v4 = t143
									t144 := int32(load32(m.memory[int64(uint32(v7))+60:]))
									t145 := v4
									v10 = t144
									v15 = t145 + v10
									m.memory[uint32(v15+i32(2))] = byte(v13&i32(63) | i32(128))
									m.memory[uint32(v15+i32(1))] = byte(int32(uint32(v13)>>6)&i32(63) | i32(128))
									v15 = v10 + i32(3)
									goto l202
								}
								if v12 != i32(39) {
									goto l193
								}
								if uint32(v11) < uint32(i32(15)) {
									if v13&i32(-83) != i32(-91) {
										t146 := int32(load16(m.memory[int64(uint32(v11<<1))+1234848:]))
										t147 := v4 + v15
										v13 = t146
										m.memory[uint32(t147)] = byte(int32(uint32(v13)>>6) | i32(192))
										t148 := int32(load32(m.memory[int64(uint32(v7))+52:]))
										v4 = t148
										t149 := int32(load32(m.memory[int64(uint32(v7))+60:]))
										t150 := v4
										v10 = t149
										m.memory[uint32(t150+v10+i32(1))] = byte(v13&i32(63) | i32(128))
										v15 = v10 + i32(2)
										goto l202
									}
									m.memory[int64(uint32(v0))+6] = byte(i32(0))
									store16(m.memory[int64(uint32(v0))+4:], uint16(i32(514)))
									goto l200
								}
							l193:
								if v12 != i32(40) {
									goto l195
								}
								if uint32(v11) < uint32(i32(16)) {
									t151 := int32(load16(m.memory[int64(uint32(v11<<1))+1234816:]))
									t152 := v4 + v15
									v13 = t151
									m.memory[uint32(t152)] = byte(int32(uint32(v13)>>6) | i32(192))
									t153 := int32(load32(m.memory[int64(uint32(v7))+52:]))
									v4 = t153
									t154 := int32(load32(m.memory[int64(uint32(v7))+60:]))
									t155 := v4
									v10 = t154
									m.memory[uint32(t155+v10+i32(1))] = byte(v13&i32(63) | i32(128))
									v15 = v10 + i32(2)
									goto l202
								}
							l195:
								if v12 != i32(37) {
									goto l197
								}
								if uint32(v11) < uint32(i32(68)) {
									t156 := int32(load16(m.memory[int64(uint32(v11<<1))+1146470:]))
									t157 := v4 + v15
									v13 = t156
									m.memory[uint32(t157)] = byte(int32(uint32(v13)>>12) | i32(224))
									t158 := int32(load32(m.memory[int64(uint32(v7))+60:]))
									t159 := v7
									v10 = t158
									v15 = v10 + i32(1)
									store32(m.memory[int64(uint32(t159))+60:], uint32(v15))
									t160 := int32(load32(m.memory[int64(uint32(v7))+52:]))
									t161 := v10
									v4 = t160
									m.memory[uint32(t161+v4+i32(2))] = byte(v13&i32(63) | i32(128))
									m.memory[uint32(v4+v15)] = byte(int32(uint32(v13)>>6)&i32(63) | i32(128))
									v15 = v10 + i32(3)
									goto l202
								}
							l197:
								v13 = v10 + i32(-188)
								if uint32(v13) < uint32(i32(927)) {
									m.fn897(v7+i32(24), i32(1242038), i32(77), v13)
									t162 := int32(load32(m.memory[int64(uint32(v7))+28:]))
									v10 = t162
									{
										{
											t163 := int32(load32(m.memory[int64(uint32(v7))+24:]))
											if t163 != i32(1) {
												goto l204
											}
											v10 = v10 + i32(-1)
											if uint32(v10) >= uint32(i32(77)) {
												m.fn33(v10, i32(77), i32(1227336))
												panic("unreachable")
											}
											v10 = v10 << 1
											t164 := int32(load16(m.memory[int64(uint32(v10))+1263160:]))
											t165 := int32(load16(m.memory[int64(uint32(v10))+1242038:]))
											v13 = t164 + v13 - t165
											goto l206
										}
									l204:
										if uint32(v10) > uint32(i32(76)) {
											m.fn33(v10, i32(77), i32(1227320))
											panic("unreachable")
										}
										t166 := int32(load16(m.memory[int64(uint32(v10<<1))+1263160:]))
										v13 = t166
									}
								l206:
									v12 = v13 & i32(0xffff)
									if uint32(v12) < uint32(i32(128)) {
										m.memory[int64(uint32(v0))+6] = byte(i32(0))
										store16(m.memory[int64(uint32(v0))+4:], uint16(i32(514)))
										goto l200
									}
									v11 = v15 + i32(1)
									v10 = v4 + v15
									if uint32(v12) < uint32(i32(2048)) {
										m.memory[uint32(v4+v11)] = byte(v13&i32(63) | i32(128))
										m.memory[uint32(v10)] = byte(int32(uint32(v13)>>6) | i32(192))
										v15 = v15 + i32(2)
										goto l202
									}
									m.memory[uint32(v10+i32(2))] = byte(v13&i32(63) | i32(128))
									m.memory[uint32(v4+v11)] = byte(int32(uint32(v13)>>6)&i32(63) | i32(128))
									m.memory[uint32(v10)] = byte(int32(uint32(v13&i32(61440))>>12) | i32(224))
									v15 = v15 + i32(3)
									goto l202
								}
								m.memory[int64(uint32(v0))+6] = byte(i32(0))
								store16(m.memory[int64(uint32(v0))+4:], uint16(i32(514)))
								goto l200
							}
							if uint32((v13+i32(127))&i32(255)) < uint32(i32(32)) {
								v12 = v13 + i32(-77)
								goto l188
							}
							if uint32((v13+i32(-97))&i32(255)) < uint32(i32(26)) {
								v12 = v13 + i32(-71)
								goto l188
							}
							v12 = v13 + i32(-65)
							if uint32(v12&i32(255)) < uint32(i32(26)) {
								goto l188
							}
							m.memory[int64(uint32(v0))+4] = byte(i32(2))
							if v13 > i32(-1) {
								store32(m.memory[int64(uint32(v0))+8:], uint32(v15))
								store32(m.memory[uint32(v0):], uint32(v5))
								store16(m.memory[int64(uint32(v0))+5:], uint16(i32(1)))
								goto l25
							}
							store32(m.memory[int64(uint32(v0))+8:], uint32(v15))
							store32(m.memory[uint32(v0):], uint32(v16))
							store16(m.memory[int64(uint32(v0))+5:], uint16(i32(2)))
							goto l25
						}
						if uint32((v13+i32(127))&i32(255)) < uint32(i32(126)) {
							v10 = v13 + i32(-77)
							goto l183
						}
						if uint32((v13+i32(-97))&i32(255)) < uint32(i32(26)) {
							v10 = v13 + i32(-71)
							goto l183
						}
						v10 = v13 + i32(-65)
						if uint32(v10&i32(255)) < uint32(i32(26)) {
							goto l183
						}
						m.memory[int64(uint32(v0))+4] = byte(i32(2))
						if v13 > i32(-1) {
							store32(m.memory[int64(uint32(v0))+8:], uint32(v15))
							store32(m.memory[uint32(v0):], uint32(v5))
							store16(m.memory[int64(uint32(v0))+5:], uint16(i32(1)))
							goto l25
						}
						store32(m.memory[int64(uint32(v0))+8:], uint32(v15))
						store32(m.memory[uint32(v0):], uint32(v16))
						store16(m.memory[int64(uint32(v0))+5:], uint16(i32(2)))
						goto l25
					l200:
						store32(m.memory[int64(uint32(v0))+8:], uint32(v15))
						store32(m.memory[uint32(v0):], uint32(v16))
						goto l25
					l188:
						v10 = v10&i32(255)*i32(84) + v12&i32(255)
						if uint32(v10) < uint32(i32(3126)) {
							{
								{
									p172 := i32(267)
									if uint32(v10) < uint32(i32(1715)) {
										p172 = i32(0)
									}
									v13 = p172
									t173 := v13
									v13 = v13 + i32(134)
									t174 := int32(load16(m.memory[int64(uint32(v13<<1))+1251496:]))
									p175 := v13
									if uint32(t174) > uint32(v10) {
										p175 = t173
									}
									v13 = p175
									t176 := v13
									v13 = v13 + i32(67)
									t177 := int32(load16(m.memory[int64(uint32(v13<<1))+1251496:]))
									p178 := v13
									if uint32(t177) > uint32(v10) {
										p178 = t176
									}
									v13 = p178
									t179 := v13
									v13 = v13 + i32(33)
									t180 := int32(load16(m.memory[int64(uint32(v13<<1))+1251496:]))
									p181 := v13
									if uint32(t180) > uint32(v10) {
										p181 = t179
									}
									v13 = p181
									t182 := v13
									v13 = v13 + i32(17)
									t183 := int32(load16(m.memory[int64(uint32(v13<<1))+1251496:]))
									p184 := v13
									if uint32(t183) > uint32(v10) {
										p184 = t182
									}
									v13 = p184
									t185 := v13
									v13 = v13 + i32(8)
									t186 := int32(load16(m.memory[int64(uint32(v13<<1))+1251496:]))
									p187 := v13
									if uint32(t186) > uint32(v10) {
										p187 = t185
									}
									v13 = p187
									t188 := v13
									v13 = v13 + i32(4)
									t189 := int32(load16(m.memory[int64(uint32(v13<<1))+1251496:]))
									p190 := v13
									if uint32(t189) > uint32(v10) {
										p190 = t188
									}
									v13 = p190
									t191 := v13
									v13 = v13 + i32(2)
									t192 := int32(load16(m.memory[int64(uint32(v13<<1))+1251496:]))
									p193 := v13
									if uint32(t192) > uint32(v10) {
										p193 = t191
									}
									v13 = p193
									t194 := v13
									v13 = v13 + i32(1)
									t195 := int32(load16(m.memory[int64(uint32(v13<<1))+1251496:]))
									p196 := v13
									if uint32(t195) > uint32(v10) {
										p196 = t194
									}
									v13 = p196
									t197 := v13
									v13 = v13 + i32(1)
									t198 := int32(load16(m.memory[int64(uint32(v13<<1))+1251496:]))
									p199 := v13
									if uint32(t198) > uint32(v10) {
										p199 = t197
									}
									v13 = p199
									v12 = v13 << 1
									t200 := int32(load16(m.memory[int64(uint32(v12))+1251496:]))
									v11 = t200
									if v11 == v10 {
										goto l212
									}
									{
										t201 := v13
										var p202 int32
										if uint32(v11) >= uint32(v10) {
											p202 = 1
										}
										v13 = t201 - p202
										if uint32(v13) >= uint32(i32(535)) {
											m.fn33(i32(-1), i32(535), i32(1227336))
											panic("unreachable")
										}
										t203 := v10
										v13 = v13 << 1
										t204 := int32(load16(m.memory[int64(uint32(v13))+1251496:]))
										t205 := int32(load16(m.memory[int64(uint32(v13))+1244436:]))
										v10 = t203 - t204 + t205
										v12 = v10 & i32(0xffff)
										v13 = int32(uint32(v12) >> 12)
										v12 = int32(uint32(v12) >> 6)
										goto l214
									}
								}
							l212:
								t206 := int32(load16(m.memory[int64(uint32(v12))+1244436:]))
								v10 = t206
								v13 = int32(uint32(v10) >> 12)
								v12 = int32(uint32(v10) >> 6)
							}
						l214:
							m.memory[uint32(v4+v15)] = byte(v13 | i32(224))
							t207 := int32(load32(m.memory[int64(uint32(v7))+52:]))
							v4 = t207
							t208 := int32(load32(m.memory[int64(uint32(v7))+60:]))
							t209 := v4
							v13 = t208
							m.memory[uint32(t209+v13+i32(1))] = byte(v12&i32(63) | i32(128))
							t210 := v7
							v15 = v13 + i32(2)
							store32(m.memory[int64(uint32(t210))+60:], uint32(v15))
							m.memory[uint32(v4+v15)] = byte(v10&i32(63) | i32(128))
							v15 = v13 + i32(3)
							goto l202
						}
						m.memory[int64(uint32(v0))+4] = byte(i32(2))
						if v13 > i32(-1) {
							store32(m.memory[int64(uint32(v0))+8:], uint32(v15))
							store32(m.memory[uint32(v0):], uint32(v5))
							store16(m.memory[int64(uint32(v0))+5:], uint16(i32(1)))
							goto l25
						}
						store32(m.memory[int64(uint32(v0))+8:], uint32(v15))
						store32(m.memory[uint32(v0):], uint32(v16))
						store16(m.memory[int64(uint32(v0))+5:], uint16(i32(2)))
						goto l25
					l183:
						{
							{
								v17 = v12*i32(178) + v10&i32(255)
								v13 = v17 & i32(0xffff)
								p211 := i32(539)
								if uint32(v13) < uint32(i32(2868)) {
									p211 = i32(0)
								}
								v10 = p211
								t212 := v10
								v10 = v10 + i32(270)
								t213 := int32(load16(m.memory[int64(uint32(v10<<1))+1245506:]))
								p214 := v10
								if uint32(t213) > uint32(v13) {
									p214 = t212
								}
								v10 = p214
								t215 := v10
								v10 = v10 + i32(135)
								t216 := int32(load16(m.memory[int64(uint32(v10<<1))+1245506:]))
								p217 := v10
								if uint32(t216) > uint32(v13) {
									p217 = t215
								}
								v10 = p217
								t218 := v10
								v10 = v10 + i32(67)
								t219 := int32(load16(m.memory[int64(uint32(v10<<1))+1245506:]))
								p220 := v10
								if uint32(t219) > uint32(v13) {
									p220 = t218
								}
								v10 = p220
								t221 := v10
								v10 = v10 + i32(34)
								t222 := int32(load16(m.memory[int64(uint32(v10<<1))+1245506:]))
								p223 := v10
								if uint32(t222) > uint32(v13) {
									p223 = t221
								}
								v10 = p223
								t224 := v10
								v10 = v10 + i32(17)
								t225 := int32(load16(m.memory[int64(uint32(v10<<1))+1245506:]))
								p226 := v10
								if uint32(t225) > uint32(v13) {
									p226 = t224
								}
								v10 = p226
								t227 := v10
								v10 = v10 + i32(8)
								t228 := int32(load16(m.memory[int64(uint32(v10<<1))+1245506:]))
								p229 := v10
								if uint32(t228) > uint32(v13) {
									p229 = t227
								}
								v10 = p229
								t230 := v10
								v10 = v10 + i32(4)
								t231 := int32(load16(m.memory[int64(uint32(v10<<1))+1245506:]))
								p232 := v10
								if uint32(t231) > uint32(v13) {
									p232 = t230
								}
								v10 = p232
								t233 := v10
								v10 = v10 + i32(2)
								t234 := int32(load16(m.memory[int64(uint32(v10<<1))+1245506:]))
								p235 := v10
								if uint32(t234) > uint32(v13) {
									p235 = t233
								}
								v10 = p235
								t236 := v10
								v10 = v10 + i32(1)
								t237 := int32(load16(m.memory[int64(uint32(v10<<1))+1245506:]))
								p238 := v10
								if uint32(t237) > uint32(v13) {
									p238 = t236
								}
								v10 = p238
								t239 := v10
								v10 = v10 + i32(1)
								t240 := int32(load16(m.memory[int64(uint32(v10<<1))+1245506:]))
								p241 := v10
								if uint32(t240) > uint32(v13) {
									p241 = t239
								}
								v10 = p241
								v12 = v10 << 1
								t242 := int32(load16(m.memory[int64(uint32(v12))+1245506:]))
								v11 = t242
								if v11 == v13 {
									goto l215
								}
								{
									t243 := v10
									var p244 int32
									if uint32(v11) >= uint32(v13) {
										p244 = 1
									}
									v13 = t243 - p244
									if uint32(v13) >= uint32(i32(1079)) {
										m.fn33(i32(-1), i32(1079), i32(1227336))
										panic("unreachable")
									}
									t245 := v17
									v13 = v13 << 1
									t246 := int32(load16(m.memory[int64(uint32(v13))+1245506:]))
									t247 := int32(load16(m.memory[int64(uint32(v13))+1242212:]))
									v10 = t245 - t246 + t247
									v12 = v10 & i32(0xffff)
									v13 = int32(uint32(v12) >> 12)
									v12 = int32(uint32(v12) >> 6)
									goto l217
								}
							}
						l215:
							t248 := int32(load16(m.memory[int64(uint32(v12))+1242212:]))
							v10 = t248
							v13 = int32(uint32(v10) >> 12)
							v12 = int32(uint32(v10) >> 6)
						}
					l217:
						m.memory[uint32(v4+v15)] = byte(v13 | i32(224))
						t249 := int32(load32(m.memory[int64(uint32(v7))+52:]))
						v4 = t249
						t250 := int32(load32(m.memory[int64(uint32(v7))+60:]))
						t251 := v4
						v13 = t250
						m.memory[uint32(t251+v13+i32(1))] = byte(v12&i32(63) | i32(128))
						t252 := v7
						v15 = v13 + i32(2)
						store32(m.memory[int64(uint32(t252))+60:], uint32(v15))
						m.memory[uint32(v4+v15)] = byte(v10&i32(63) | i32(128))
						v15 = v13 + i32(3)
					}
				l202:
					store32(m.memory[int64(uint32(v7))+60:], uint32(v15))
					if uint32(v16) < uint32(v3) {
						goto l218
					}
					store32(m.memory[int64(uint32(v0))+8:], uint32(v15))
					store32(m.memory[uint32(v0):], uint32(v16))
					m.memory[int64(uint32(v0))+4] = byte(i32(0))
					goto l25
				l218:
					{
						t253 := int32(load32(m.memory[int64(uint32(v7))+56:]))
						if uint32(v15+i32(2)) >= uint32(t253) {
							store32(m.memory[int64(uint32(v0))+8:], uint32(v15))
							store32(m.memory[uint32(v0):], uint32(v16))
							m.memory[int64(uint32(v0))+4] = byte(i32(1))
							goto l25
						}
						v17 = v5 + i32(2)
						{
							t254 := int32(int8(m.memory[uint32(v2+v16)]))
							v10 = t254
							if v10 >= i32(0) {
								m.memory[uint32(v4+v15)] = byte(v10)
								t255 := int32(load32(m.memory[int64(uint32(v7))+60:]))
								t256 := v7
								v13 = t255
								v15 = v13 + i32(1)
								store32(m.memory[int64(uint32(t256))+60:], uint32(v15))
								if uint32(v10) > uint32(i32(59)) {
									goto l222
								}
								if v16 == v26 {
									goto l223
								}
								{
									t257 := int32(load32(m.memory[int64(uint32(v7))+56:]))
									if uint32(v13+i32(3)) >= uint32(t257) {
										goto l224
									}
									v5 = v5 + i32(3)
								l226:
									{
										t258 := int32(int8(m.memory[uint32(v2+v5+i32(-1))]))
										v10 = t258
										if v10 < i32(0) {
											goto l221
										}
										m.memory[uint32(v4+v15)] = byte(v10)
										t259 := int32(load32(m.memory[int64(uint32(v7))+60:]))
										t260 := v7
										v13 = t259
										v15 = v13 + i32(1)
										store32(m.memory[int64(uint32(t260))+60:], uint32(v15))
										if uint32(v10) <= uint32(i32(59)) {
											goto l225
										}
										v17 = v5
										goto l222
									l225:
										if v3 == v5 {
											goto l223
										}
										v5 = v5 + i32(1)
										t261 := int32(load32(m.memory[int64(uint32(v7))+56:]))
										if uint32(v13+i32(3)) < uint32(t261) {
											goto l226
										}
									}
									v17 = v5 + i32(-1)
								}
							l224:
								store32(m.memory[int64(uint32(v0))+8:], uint32(v15))
								store32(m.memory[uint32(v0):], uint32(v17))
								m.memory[int64(uint32(v0))+4] = byte(i32(1))
								goto l25
							}
							v5 = v17
							goto l221
						}
					}
				l223:
					store32(m.memory[int64(uint32(v0))+8:], uint32(v15))
					store32(m.memory[uint32(v0):], uint32(v3))
					m.memory[int64(uint32(v0))+4] = byte(i32(0))
					goto l25
				l221:
					v12 = v10 + i32(127)
					if uint32(v12&i32(255)) <= uint32(i32(125)) {
						goto l227
					}
				}
			l176:
				m.memory[int64(uint32(v0))+6] = byte(i32(0))
				store16(m.memory[int64(uint32(v0))+4:], uint16(i32(258)))
			l179:
				store32(m.memory[int64(uint32(v0))+8:], uint32(v15))
				store32(m.memory[uint32(v0):], uint32(v5))
				goto l25
			l222:
				t262 := int32(load32(m.memory[int64(uint32(v7))+56:]))
				t263 := v15
				v5 = t262
				if uint32(t263) <= uint32(v5) {
					goto l228
				}
			}
			m.fn121(v15, v5, v5, i32(1146188))
			panic("unreachable")
		case 6:
			store32(m.memory[int64(uint32(v7))+56:], uint32(v5))
			store32(m.memory[int64(uint32(v7))+52:], uint32(v4))
			v17 = i32(0)
			v15 = i32(0)
			{
				t264 := int32(m.memory[int64(uint32(v1))+1])
				if t264 == 0 {
					goto l343
				}
				m.memory[int64(uint32(v1))+1] = byte(i32(0))
				if v3 == 0 {
					if v6 != 0 {
						m.memory[int64(uint32(v0))+6] = byte(i32(0))
						store16(m.memory[int64(uint32(v0))+4:], uint16(i32(258)))
						store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
						store32(m.memory[uint32(v0):], uint32(i32(0)))
						goto l25
					}
					store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
					store32(m.memory[uint32(v0):], uint32(i32(0)))
					m.memory[int64(uint32(v0))+4] = byte(i32(0))
					goto l25
				}
				if uint32(v5) > uint32(i32(2)) {
					goto l231
				}
				store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
				store32(m.memory[uint32(v0):], uint32(i32(0)))
				m.memory[int64(uint32(v0))+4] = byte(i32(1))
				goto l25
			l231:
				t265 := int32(int8(m.memory[uint32(v2)]))
				v13 = t265
				{
					{
						{
							{
								t266 := int32(m.memory[int64(uint32(v1))+2])
								v10 = t266
								if v10 != i32(1) {
									goto l233
								}
								v15 = (v13 + i32(97)) & i32(255)
								if uint32(v15) < uint32(i32(83)) {
									m.memory[uint32(v4)] = byte(i32(227))
									t268 := v4
									v13 = v15 + i32(12353)
									m.memory[int64(uint32(t268))+2] = byte(v13&i32(63) | i32(128))
									m.memory[int64(uint32(v4))+1] = byte(int32(uint32(v13)>>6) & i32(131))
									goto l247
								}
							}
						l233:
							v15 = v13 + i32(-64)
							if uint32(v15&i32(255)) <= uint32(i32(62)) {
								goto l235
							}
							if v13 > i32(-4) {
								m.memory[int64(uint32(v0))+4] = byte(i32(2))
								if v13 > i32(-1) {
									store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
									store32(m.memory[uint32(v0):], uint32(i32(0)))
									store16(m.memory[int64(uint32(v0))+5:], uint16(i32(1)))
									goto l25
								}
								store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
								store32(m.memory[uint32(v0):], uint32(i32(1)))
								store16(m.memory[int64(uint32(v0))+5:], uint16(i32(2)))
								goto l25
							}
							v15 = v13 + i32(-65)
						l235:
							if v10 != i32(2) {
								goto l237
							}
							v16 = v15 & i32(255)
							if uint32(v16) < uint32(i32(86)) {
								m.memory[uint32(v4)] = byte(i32(227))
								t272 := v4
								v13 = v16 + i32(12449)
								m.memory[int64(uint32(t272))+1] = byte(int32(uint32(v13)>>6) & i32(135))
								m.memory[int64(uint32(v4))+2] = byte(v13&i32(63) | i32(128))
								store32(m.memory[int64(uint32(v7))+60:], uint32(i32(2)))
								goto l247
							}
						l237:
							v10 = v10*i32(188) + v15&i32(255)
							v16 = v10 + i32(-1410)
							if uint32(v16) < uint32(i32(2965)) {
								v15 = i32(1)
								t273 := int32(load16(m.memory[int64(uint32(v16<<1))+1234998:]))
								t274 := v4
								v13 = t273
								m.memory[int64(uint32(t274))+2] = byte(v13&i32(63) | i32(128))
								m.memory[uint32(v4)] = byte(int32(uint32(v13)>>12) | i32(224))
								m.memory[int64(uint32(v4))+1] = byte(int32(uint32(v13)>>6)&i32(63) | i32(128))
								goto l268
							}
							v16 = v10 + i32(-4418)
							if uint32(v16) < uint32(i32(3390)) {
								v15 = i32(1)
								t275 := int32(load16(m.memory[int64(uint32(v16<<1))+1263314:]))
								t276 := v4
								v13 = t275
								m.memory[int64(uint32(t276))+2] = byte(v13&i32(63) | i32(128))
								m.memory[uint32(v4)] = byte(int32(uint32(v13)>>12) | i32(224))
								m.memory[int64(uint32(v4))+1] = byte(int32(uint32(v13)>>6)&i32(63) | i32(128))
								goto l268
							}
							v16 = v10 + i32(-10744)
							if uint32(v16) < uint32(i32(360)) {
								v15 = i32(1)
								t277 := int32(load16(m.memory[int64(uint32(v16<<1))+1270094:]))
								t278 := v4
								v13 = t277
								m.memory[int64(uint32(t278))+2] = byte(v13&i32(63) | i32(128))
								m.memory[uint32(v4)] = byte(int32(uint32(v13)>>12) | i32(224))
								m.memory[int64(uint32(v4))+1] = byte(int32(uint32(v13)>>6)&i32(63) | i32(128))
								goto l268
							}
							v16 = v10 + i32(-8272)
							if uint32(v16) < uint32(i32(360)) {
								goto l242
							}
							if uint32(v10+i32(-8836)) < uint32(i32(1880)) {
								t269 := v4
								v13 = v10 + i32(-17028)
								m.memory[int64(uint32(t269))+2] = byte(v13&i32(63) | i32(128))
								m.memory[int64(uint32(v4))+1] = byte(int32(uint32(v13)>>6)&i32(63) | i32(128))
								m.memory[uint32(v4)] = byte(int32(uint32(v13&i32(61440))>>12) | i32(224))
								goto l248
							}
							m.fn898(v7+i32(16), v10)
							t267 := int32(load16(m.memory[int64(uint32(v7))+16:]))
							if t267&i32(1) != 0 {
								goto l244
							}
							v16 = v10 + i32(-203)
							if uint32(v16) >= uint32(i32(10)) {
								goto l245
							}
							v12 = i32(2)
							goto l246
						}
					l244:
						t270 := int32(load16(m.memory[int64(uint32(v7))+18:]))
						v13 = t270
						if uint32(v13&i32(0xffff)) < uint32(i32(2048)) {
							m.memory[int64(uint32(v4))+1] = byte(v13&i32(63) | i32(128))
							m.memory[uint32(v4)] = byte(int32(uint32(v13)>>6) | i32(192))
							v17 = i32(2)
							goto l250
						}
						m.memory[int64(uint32(v4))+2] = byte(v13&i32(63) | i32(128))
						m.memory[int64(uint32(v4))+1] = byte(int32(uint32(v13)>>6)&i32(63) | i32(128))
						m.memory[uint32(v4)] = byte(int32(uint32(v13&i32(61440))>>12) | i32(224))
					}
				l248:
					v17 = i32(3)
					goto l250
				l245:
					v16 = v10 + i32(-220)
					if uint32(v16) >= uint32(i32(26)) {
						goto l251
					}
					v12 = i32(5)
					goto l246
				l251:
					v16 = v10 + i32(-252)
					if uint32(v16) >= uint32(i32(26)) {
						goto l252
					}
					v12 = i32(8)
					goto l246
				l252:
					v16 = v10 + i32(-470)
					if uint32(v16) >= uint32(i32(17)) {
						goto l253
					}
					v12 = i32(11)
					goto l246
				l253:
					v16 = v10 + i32(-487)
					if uint32(v16) >= uint32(i32(7)) {
						goto l254
					}
					v12 = i32(14)
					goto l246
				l254:
					v12 = i32(17)
					v16 = v10 + i32(-502)
					if uint32(v16) < uint32(i32(17)) {
						goto l246
					}
					v16 = v10 + i32(-519)
					if uint32(v16) >= uint32(i32(7)) {
						goto l255
					}
					v12 = i32(20)
					goto l246
				l255:
					v16 = v10 + i32(-564)
					if uint32(v16) >= uint32(i32(6)) {
						goto l256
					}
					v12 = i32(23)
					goto l246
				l256:
					v16 = i32(0)
					if v10 != i32(570) {
						goto l257
					}
					v12 = i32(26)
					goto l246
				l257:
					v15 = v10 + i32(-571)
					if uint32(v15) >= uint32(i32(26)) {
						goto l258
					}
					v12 = i32(29)
					v16 = v15
					goto l246
				l258:
					v15 = v10 + i32(-612)
					if uint32(v15) >= uint32(i32(6)) {
						goto l259
					}
					v12 = i32(32)
					v16 = v15
					goto l246
				l259:
					if v10 != i32(618) {
						goto l260
					}
					v12 = i32(35)
					goto l246
				l260:
					v16 = v10 + i32(-619)
					if uint32(v16) >= uint32(i32(26)) {
						goto l261
					}
					v12 = i32(38)
					goto l246
				l261:
					v16 = v10 + i32(-1128)
					if uint32(v16) >= uint32(i32(20)) {
						goto l262
					}
					v12 = i32(41)
					goto l246
				l262:
					v16 = v10 + i32(-1148)
					if uint32(v16) >= uint32(i32(10)) {
						goto l263
					}
					v12 = i32(44)
					goto l246
				l263:
					v16 = v10 + i32(-8634)
					if uint32(v16) >= uint32(i32(10)) {
						goto l264
					}
					v12 = i32(47)
					goto l246
				l264:
					v16 = v10 + i32(-10716)
					if uint32(v16) >= uint32(i32(10)) {
						goto l265
					}
					v12 = i32(50)
					goto l246
				l265:
					v16 = v10 + i32(-10726)
					if uint32(v16) >= uint32(i32(10)) {
						m.memory[int64(uint32(v0))+4] = byte(i32(2))
						if v13 > i32(-1) {
							store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
							store32(m.memory[uint32(v0):], uint32(i32(0)))
							store16(m.memory[int64(uint32(v0))+5:], uint16(i32(1)))
							goto l25
						}
						store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
						store32(m.memory[uint32(v0):], uint32(i32(1)))
						store16(m.memory[int64(uint32(v0))+5:], uint16(i32(2)))
						goto l25
					}
					v12 = i32(53)
				l246:
					v15 = i32(1)
					{
						t271 := int32(load16(m.memory[int64(uint32(v12<<1))+1241434:]))
						v13 = t271 + v16
						if uint32(v13&i32(0xffff)) < uint32(i32(2048)) {
							m.memory[int64(uint32(v4))+1] = byte(v13&i32(63) | i32(128))
							m.memory[uint32(v4)] = byte(int32(uint32(v13)>>6) | i32(192))
							v17 = i32(2)
							store32(m.memory[int64(uint32(v7))+60:], uint32(i32(2)))
							goto l343
						}
						m.memory[int64(uint32(v4))+2] = byte(v13&i32(63) | i32(128))
						m.memory[int64(uint32(v4))+1] = byte(int32(uint32(v13)>>6)&i32(63) | i32(128))
						m.memory[uint32(v4)] = byte(int32(uint32(v13&i32(61440))>>12) | i32(224))
						goto l268
					}
				l242:
					v15 = i32(1)
					t279 := int32(load16(m.memory[int64(uint32(v16<<1))+1270094:]))
					t280 := v4
					v13 = t279
					m.memory[int64(uint32(t280))+2] = byte(v13&i32(63) | i32(128))
					m.memory[uint32(v4)] = byte(int32(uint32(v13)>>12) | i32(224))
					m.memory[int64(uint32(v4))+1] = byte(int32(uint32(v13)>>6)&i32(63) | i32(128))
				}
			l268:
				v17 = i32(3)
				store32(m.memory[int64(uint32(v7))+60:], uint32(i32(3)))
				goto l343
			l250:
				store32(m.memory[int64(uint32(v7))+60:], uint32(v17))
				v15 = i32(1)
				goto l343
			l247:
				v15 = i32(1)
				v17 = i32(3)
			}
		l343:
			{
				{
					{
						if uint32(v5) < uint32(v17) {
							m.fn121(v17, v5, v5, i32(1146188))
							panic("unreachable")
						}
						v13 = v5 - v17
						t281 := v13
						v10 = v3 - v15
						t282 := v10
						var p283 int32
						if uint32(v13) < uint32(v10) {
							p283 = 1
						}
						v14 = p283
						p284 := t282
						if v14 != 0 {
							p284 = t281
						}
						v12 = p284
						v13 = i32(0)
						v11 = v4 + v17
						t285 := v11
						v16 = v2 + v15
						if (t285^v16)&i32(3) != 0 {
							goto l272
						}
						v13 = i32(0)
						v18 = (i32(0) - v16) & i32(3)
						if uint32(v18|i32(8)) > uint32(v12) {
							goto l272
						}
						if v18 != 0 {
							v13 = i32(0)
							t286 := int32(int8(m.memory[uint32(v16)]))
							v10 = t286
							if v10 < i32(0) {
								goto l275
							}
							m.memory[uint32(v11)] = byte(v10)
							v13 = i32(1)
							if v18 == i32(1) {
								goto l274
							}
							{
								t287 := int32(int8(m.memory[int64(uint32(v16))+1]))
								v10 = t287
								if v10 >= i32(0) {
									m.memory[int64(uint32(v11))+1] = byte(v10)
									v13 = i32(2)
									if v18 == i32(2) {
										goto l274
									}
									{
										t288 := int32(int8(m.memory[int64(uint32(v16))+2]))
										v10 = t288
										if v10 >= i32(0) {
											m.memory[int64(uint32(v11))+2] = byte(v10)
											v13 = i32(3)
											goto l274
										}
										v13 = i32(2)
										goto l275
									}
								}
								v13 = i32(1)
								goto l275
							}
						}
						v13 = i32(0)
						goto l274
					}
				l274:
					v8 = v12 + i32(-8)
				l281:
					{
						v18 = v16 + v13
						t289 := int32(load32(m.memory[uint32(v18):]))
						v10 = t289
						v9 = v11 + v13
						t290 := int32(load32(m.memory[uint32(v18+i32(4)):]))
						t291 := v9 + i32(4)
						v18 = t290
						store32(m.memory[uint32(t291):], uint32(v18))
						store32(m.memory[uint32(v9):], uint32(v10))
						{
							v18 = v18 & i32(-2139062144)
							t292 := v18
							v10 = v10 & i32(-2139062144)
							if t292|v10 == 0 {
								goto l278
							}
							if v10 != 0 {
								goto l279
							}
							v10 = int32(uint32(int32(bits.TrailingZeros32(uint32(v18))))>>3) + i32(4)
							goto l280
						l279:
							v10 = int32(uint32(int32(bits.TrailingZeros32(uint32(v10)))) >> 3)
						l280:
							t293 := v16
							v13 = v10 + v13
							t294 := int32(m.memory[uint32(t293+v13)])
							v10 = t294
							goto l275
						}
					l278:
						v13 = v13 + i32(8)
						if uint32(v13) <= uint32(v8) {
							goto l281
						}
					}
				l272:
					if uint32(v13) >= uint32(v12) {
						goto l282
					}
				l283:
					{
						t295 := int32(int8(m.memory[uint32(v16+v13)]))
						v10 = t295
						if v10 < i32(0) {
							goto l275
						}
						m.memory[uint32(v11+v13)] = byte(v10)
						t296 := v12
						v13 = v13 + i32(1)
						if t296 != v13 {
							goto l283
						}
					}
				l282:
					v16 = v12 + v17
					v13 = v12 + v15
					goto l284
				l275:
					t297 := v7
					v16 = v13 + v17
					store32(m.memory[int64(uint32(t297))+60:], uint32(v16))
					v13 = v13 + v15
					if uint32(v16+i32(2)) < uint32(v5) {
						goto l285
					}
					v14 = i32(1)
				}
			l284:
				store32(m.memory[int64(uint32(v0))+8:], uint32(v16))
				store32(m.memory[uint32(v0):], uint32(v13))
				store32(m.memory[int64(uint32(v0))+4:], uint32(v14|(v7+i32(52))&i32(-256)))
				goto l25
			l285:
				v15 = v13 + i32(1)
			l342:
				{
					v13 = v10 + i32(127)
					if uint32(v13&i32(255)) < uint32(i32(31)) {
						goto l286
					}
					if uint32((v10+i32(3))&i32(255)) < uint32(i32(227)) {
						v5 = (v10 + i32(95)) & i32(255)
						if uint32(v5) > uint32(i32(62)) {
							if v10&i32(255) != i32(128) {
								m.memory[int64(uint32(v0))+6] = byte(i32(0))
								store16(m.memory[int64(uint32(v0))+4:], uint16(i32(258)))
								goto l290
							}
							m.memory[uint32(v4+v16)] = byte(i32(194))
							t314 := int32(load32(m.memory[int64(uint32(v7))+60:]))
							t315 := v7
							v16 = t314 + i32(1)
							store32(m.memory[int64(uint32(t315))+60:], uint32(v16))
							v10 = i32(128)
							t316 := int32(load32(m.memory[int64(uint32(v7))+52:]))
							v4 = t316
							goto l336
						}
						m.memory[uint32(v4+v16)] = byte(i32(239))
						t309 := int32(load32(m.memory[int64(uint32(v7))+52:]))
						v4 = t309
						t310 := int32(load32(m.memory[int64(uint32(v7))+60:]))
						t311 := v4
						v13 = t310
						t312 := t311 + v13 + i32(1)
						v5 = v5 + i32(-159)
						m.memory[uint32(t312)] = byte(int32(uint32(v5)>>6) & i32(191))
						t313 := v7
						v16 = v13 + i32(2)
						store32(m.memory[int64(uint32(t313))+60:], uint32(v16))
						v10 = v5&i32(63) | i32(-128)
						goto l336
					}
					v13 = v10 + i32(63)
				l286:
					if uint32(v15) < uint32(v3) {
						v10 = v15 + i32(1)
						v12 = v2 + v15
						t298 := int32(int8(m.memory[uint32(v12)]))
						v5 = t298
						v13 = v13 & i32(255)
						if v13 != i32(1) {
							goto l291
						}
						v11 = (v5 + i32(97)) & i32(255)
						if uint32(v11) < uint32(i32(83)) {
							m.memory[uint32(v4+v16)] = byte(i32(227))
							t299 := int32(load32(m.memory[int64(uint32(v7))+52:]))
							v4 = t299
							t300 := int32(load32(m.memory[int64(uint32(v7))+60:]))
							t301 := v4
							v5 = t300
							v13 = t301 + v5
							t302 := v13 + i32(2)
							v16 = v11 + i32(12353)
							m.memory[uint32(t302)] = byte(v16&i32(63) | i32(128))
							m.memory[uint32(v13+i32(1))] = byte(int32(uint32(v16)>>6) & i32(131))
							v16 = v5 + i32(3)
							goto l304
						}
					l291:
						v11 = v5 + i32(-64)
						if uint32(v11&i32(255)) <= uint32(i32(62)) {
							goto l293
						}
						if v5 > i32(-4) {
							m.memory[int64(uint32(v0))+4] = byte(i32(2))
							if v5 > i32(-1) {
								store32(m.memory[int64(uint32(v0))+8:], uint32(v16))
								store32(m.memory[uint32(v0):], uint32(v15))
								store16(m.memory[int64(uint32(v0))+5:], uint16(i32(1)))
								goto l25
							}
							store32(m.memory[int64(uint32(v0))+8:], uint32(v16))
							store32(m.memory[uint32(v0):], uint32(v10))
							store16(m.memory[int64(uint32(v0))+5:], uint16(i32(2)))
							goto l25
						}
						v11 = v5 + i32(-65)
					l293:
						if v13 != i32(2) {
							goto l295
						}
						v17 = v11 & i32(255)
						if uint32(v17) < uint32(i32(86)) {
							m.memory[uint32(v4+v16)] = byte(i32(227))
							t317 := int32(load32(m.memory[int64(uint32(v7))+52:]))
							v4 = t317
							t318 := int32(load32(m.memory[int64(uint32(v7))+60:]))
							t319 := v4
							v5 = t318
							v13 = t319 + v5
							t320 := v13 + i32(2)
							v16 = v17 + i32(12449)
							m.memory[uint32(t320)] = byte(v16&i32(63) | i32(128))
							m.memory[uint32(v13+i32(1))] = byte(int32(uint32(v16)>>6) & i32(135))
							v16 = v5 + i32(3)
							goto l304
						}
					l295:
						v13 = v13*i32(188) + v11&i32(255)
						v11 = v13 + i32(-1410)
						if uint32(v11) < uint32(i32(2965)) {
							t321 := int32(load16(m.memory[int64(uint32(v11<<1))+1234998:]))
							t322 := v4 + v16
							v5 = t321
							m.memory[uint32(t322)] = byte(int32(uint32(v5)>>12) | i32(224))
							t323 := int32(load32(m.memory[int64(uint32(v7))+52:]))
							v4 = t323
							t324 := int32(load32(m.memory[int64(uint32(v7))+60:]))
							t325 := v4
							v13 = t324
							v16 = t325 + v13
							m.memory[uint32(v16+i32(2))] = byte(v5&i32(63) | i32(128))
							m.memory[uint32(v16+i32(1))] = byte(int32(uint32(v5)>>6)&i32(63) | i32(128))
							v16 = v13 + i32(3)
							goto l304
						}
						v11 = v13 + i32(-4418)
						if uint32(v11) < uint32(i32(3390)) {
							t326 := int32(load16(m.memory[int64(uint32(v11<<1))+1263314:]))
							t327 := v4 + v16
							v5 = t326
							m.memory[uint32(t327)] = byte(int32(uint32(v5)>>12) | i32(224))
							t328 := int32(load32(m.memory[int64(uint32(v7))+52:]))
							v4 = t328
							t329 := int32(load32(m.memory[int64(uint32(v7))+60:]))
							t330 := v4
							v13 = t329
							v16 = t330 + v13
							m.memory[uint32(v16+i32(2))] = byte(v5&i32(63) | i32(128))
							m.memory[uint32(v16+i32(1))] = byte(int32(uint32(v5)>>6)&i32(63) | i32(128))
							v16 = v13 + i32(3)
							goto l304
						}
						v11 = v13 + i32(-10744)
						if uint32(v11) < uint32(i32(360)) {
							t331 := int32(load16(m.memory[int64(uint32(v11<<1))+1270094:]))
							t332 := v4 + v16
							v5 = t331
							m.memory[uint32(t332)] = byte(int32(uint32(v5)>>12) | i32(224))
							t333 := int32(load32(m.memory[int64(uint32(v7))+52:]))
							v4 = t333
							t334 := int32(load32(m.memory[int64(uint32(v7))+60:]))
							t335 := v4
							v13 = t334
							v16 = t335 + v13
							m.memory[uint32(v16+i32(2))] = byte(v5&i32(63) | i32(128))
							m.memory[uint32(v16+i32(1))] = byte(int32(uint32(v5)>>6)&i32(63) | i32(128))
							v16 = v13 + i32(3)
							goto l304
						}
						v11 = v13 + i32(-8272)
						if uint32(v11) < uint32(i32(360)) {
							t336 := int32(load16(m.memory[int64(uint32(v11<<1))+1270094:]))
							t337 := v4 + v16
							v5 = t336
							m.memory[uint32(t337)] = byte(int32(uint32(v5)>>12) | i32(224))
							t338 := int32(load32(m.memory[int64(uint32(v7))+60:]))
							t339 := v7
							v13 = t338
							v16 = v13 + i32(1)
							store32(m.memory[int64(uint32(t339))+60:], uint32(v16))
							t340 := int32(load32(m.memory[int64(uint32(v7))+52:]))
							t341 := v13
							v4 = t340
							m.memory[uint32(t341+v4+i32(2))] = byte(v5&i32(63) | i32(128))
							m.memory[uint32(v4+v16)] = byte(int32(uint32(v5)>>6)&i32(63) | i32(128))
							v16 = v13 + i32(3)
							goto l304
						}
						if uint32(v13+i32(-8836)) < uint32(i32(1880)) {
							t303 := v4 + v16
							v5 = v13 + i32(-17028)
							m.memory[uint32(t303)] = byte(int32(uint32(v5&i32(61440))>>12) | i32(224))
							t304 := int32(load32(m.memory[int64(uint32(v7))+60:]))
							t305 := v7
							v13 = t304
							v16 = v13 + i32(1)
							store32(m.memory[int64(uint32(t305))+60:], uint32(v16))
							t306 := int32(load32(m.memory[int64(uint32(v7))+52:]))
							t307 := v13
							v4 = t306
							m.memory[uint32(t307+v4+i32(2))] = byte(v5&i32(63) | i32(128))
							m.memory[uint32(v4+v16)] = byte(int32(uint32(v5)>>6)&i32(63) | i32(128))
							v16 = v13 + i32(3)
							goto l304
						}
						if uint32(v13) >= uint32(i32(108)) {
							v11 = v13 + i32(-119)
							if uint32(v11) >= uint32(i32(8)) {
								v11 = v13 + i32(-135)
								if uint32(v11) >= uint32(i32(7)) {
									v11 = v13 + i32(-153)
									if uint32(v11) >= uint32(i32(15)) {
										v11 = v13 + i32(-175)
										if uint32(v11) >= uint32(i32(8)) {
											if v13 != i32(187) {
												v11 = v13 + i32(-658)
												if uint32(v11) >= uint32(i32(32)) {
													v17 = i32(23)
													v11 = v13 + i32(-1159)
													if uint32(v11) < uint32(i32(23)) {
														goto l303
													}
													v11 = v13 + i32(-1190)
													if uint32(v11) >= uint32(i32(30)) {
														v11 = v13 + i32(-10736)
														if uint32(v11) >= uint32(i32(8)) {
															v11 = v13 + i32(-8644)
															if uint32(v11) >= uint32(i32(4)) {
																v11 = v13 + i32(-203)
																if uint32(v11) >= uint32(i32(10)) {
																	goto l314
																}
																v17 = i32(2)
																goto l315
															}
															v17 = i32(32)
															goto l303
														}
														v17 = i32(29)
														goto l303
													}
													v17 = i32(26)
													goto l303
												}
												v17 = i32(20)
												goto l303
											}
											v17 = i32(17)
											v11 = i32(0)
											goto l303
										}
										v17 = i32(14)
										goto l303
									}
									v17 = i32(11)
									goto l303
								}
								v17 = i32(8)
								goto l303
							}
							v17 = i32(5)
							goto l303
						}
						v17 = i32(2)
						v11 = v13
						goto l303
					l314:
						v11 = v13 + i32(-220)
						if uint32(v11) >= uint32(i32(26)) {
							goto l316
						}
						v17 = i32(5)
						goto l315
					l316:
						v11 = v13 + i32(-252)
						if uint32(v11) >= uint32(i32(26)) {
							goto l317
						}
						v17 = i32(8)
						goto l315
					l317:
						v11 = v13 + i32(-470)
						if uint32(v11) >= uint32(i32(17)) {
							goto l318
						}
						v17 = i32(11)
						goto l315
					l318:
						v11 = v13 + i32(-487)
						if uint32(v11) >= uint32(i32(7)) {
							goto l319
						}
						v17 = i32(14)
						goto l315
					l319:
						v17 = i32(17)
						v11 = v13 + i32(-502)
						if uint32(v11) < uint32(i32(17)) {
							goto l315
						}
						v11 = v13 + i32(-519)
						if uint32(v11) >= uint32(i32(7)) {
							goto l320
						}
						v17 = i32(20)
						goto l315
					l320:
						v11 = v13 + i32(-564)
						if uint32(v11) >= uint32(i32(6)) {
							goto l321
						}
						v17 = i32(23)
						goto l315
					l321:
						v11 = i32(0)
						if v13 != i32(570) {
							goto l322
						}
						v17 = i32(26)
						goto l315
					l322:
						v18 = v13 + i32(-571)
						if uint32(v18) >= uint32(i32(26)) {
							goto l323
						}
						v17 = i32(29)
						v11 = v18
						goto l315
					l323:
						v18 = v13 + i32(-612)
						if uint32(v18) >= uint32(i32(6)) {
							goto l324
						}
						v17 = i32(32)
						v11 = v18
						goto l315
					l324:
						if v13 != i32(618) {
							goto l325
						}
						v17 = i32(35)
						goto l315
					l325:
						v11 = v13 + i32(-619)
						if uint32(v11) >= uint32(i32(26)) {
							goto l326
						}
						v17 = i32(38)
						goto l315
					l326:
						v11 = v13 + i32(-1128)
						if uint32(v11) >= uint32(i32(20)) {
							goto l327
						}
						v17 = i32(41)
						goto l315
					l327:
						v11 = v13 + i32(-1148)
						if uint32(v11) >= uint32(i32(10)) {
							goto l328
						}
						v17 = i32(44)
						goto l315
					l328:
						v11 = v13 + i32(-8634)
						if uint32(v11) >= uint32(i32(10)) {
							goto l329
						}
						v17 = i32(47)
						goto l315
					l329:
						v11 = v13 + i32(-10716)
						if uint32(v11) >= uint32(i32(10)) {
							goto l330
						}
						v17 = i32(50)
						goto l315
					l330:
						v11 = v13 + i32(-10726)
						if uint32(v11) >= uint32(i32(10)) {
							m.memory[int64(uint32(v0))+4] = byte(i32(2))
							if v5 > i32(-1) {
								store32(m.memory[int64(uint32(v0))+8:], uint32(v16))
								store32(m.memory[uint32(v0):], uint32(v15))
								store16(m.memory[int64(uint32(v0))+5:], uint16(i32(1)))
								goto l25
							}
							store32(m.memory[int64(uint32(v0))+8:], uint32(v16))
							store32(m.memory[uint32(v0):], uint32(v10))
							store16(m.memory[int64(uint32(v0))+5:], uint16(i32(2)))
							goto l25
						}
						v17 = i32(53)
					l315:
						v18 = v16 + i32(1)
						v13 = v4 + v16
						{
							t308 := int32(load16(m.memory[int64(uint32(v17<<1))+1241434:]))
							v5 = t308 + v11
							if uint32(v5&i32(0xffff)) < uint32(i32(2048)) {
								m.memory[uint32(v4+v18)] = byte(v5&i32(63) | i32(128))
								m.memory[uint32(v13)] = byte(int32(uint32(v5)>>6) | i32(192))
								v16 = v16 + i32(2)
								goto l304
							}
							m.memory[uint32(v13+i32(2))] = byte(v5&i32(63) | i32(128))
							m.memory[uint32(v4+v18)] = byte(int32(uint32(v5)>>6)&i32(63) | i32(128))
							m.memory[uint32(v13)] = byte(int32(uint32(v5&i32(61440))>>12) | i32(224))
							v16 = v16 + i32(3)
							goto l304
						}
					}
					if v6 != 0 {
						m.memory[int64(uint32(v0))+6] = byte(i32(0))
						store16(m.memory[int64(uint32(v0))+4:], uint16(i32(258)))
						goto l290
					}
					m.memory[int64(uint32(v1))+2] = byte(v13)
					m.memory[int64(uint32(v1))+1] = byte(i32(1))
					m.memory[int64(uint32(v0))+4] = byte(i32(0))
					goto l290
				l303:
					{
						t342 := int32(load16(m.memory[int64(uint32(v17<<1))+1241972:]))
						v5 = v11 + t342
						if uint32(v5) < uint32(i32(240)) {
							goto l338
						}
						m.fn33(v5, i32(240), i32(1241544))
						panic("unreachable")
					}
				l338:
					v11 = v16 + i32(1)
					v13 = v4 + v16
					{
						t343 := int32(load16(m.memory[int64(uint32(v5<<1))+1226508:]))
						v5 = t343
						if uint32(v5) < uint32(i32(2048)) {
							goto l339
						}
						m.memory[uint32(v13)] = byte(int32(uint32(v5)>>12) | i32(224))
						m.memory[uint32(v13+i32(2))] = byte(v5&i32(63) | i32(128))
						m.memory[uint32(v4+v11)] = byte(int32(uint32(v5)>>6)&i32(63) | i32(128))
						v16 = v16 + i32(3)
						goto l304
					}
				l339:
					m.memory[uint32(v4+v11)] = byte(v5&i32(63) | i32(128))
					m.memory[uint32(v13)] = byte(int32(uint32(v5)>>6) | i32(192))
					v16 = v16 + i32(2)
				l304:
					store32(m.memory[int64(uint32(v7))+60:], uint32(v16))
					if uint32(v10) < uint32(v3) {
						goto l340
					}
					store32(m.memory[int64(uint32(v0))+8:], uint32(v16))
					store32(m.memory[uint32(v0):], uint32(v10))
					m.memory[int64(uint32(v0))+4] = byte(i32(0))
					goto l25
				l340:
					{
						t344 := int32(load32(m.memory[int64(uint32(v7))+56:]))
						if uint32(v16+i32(2)) < uint32(t344) {
							goto l341
						}
						store32(m.memory[int64(uint32(v0))+8:], uint32(v16))
						store32(m.memory[uint32(v0):], uint32(v10))
						m.memory[int64(uint32(v0))+4] = byte(i32(1))
						goto l25
					}
				l341:
					v15 = v15 + i32(2)
					t345 := int32(int8(m.memory[uint32(v12+i32(1))]))
					v10 = t345
					if v10 < i32(0) {
						goto l342
					}
				}
			l336:
				m.memory[uint32(v4+v16)] = byte(v10)
				t346 := int32(load32(m.memory[int64(uint32(v7))+60:]))
				t347 := v7
				v17 = t346 + i32(1)
				store32(m.memory[int64(uint32(t347))+60:], uint32(v17))
				t348 := int32(load32(m.memory[int64(uint32(v7))+56:]))
				v5 = t348
				if uint32(v15) <= uint32(v3) {
					goto l343
				}
			}
			m.fn121(v15, v3, v3, i32(1146204))
			panic("unreachable")
		l290:
			store32(m.memory[int64(uint32(v0))+8:], uint32(v16))
			store32(m.memory[uint32(v0):], uint32(v15))
			goto l25
		case 5:
			v16 = i32(0)
			{
				t349 := int32(m.memory[int64(uint32(v1))+2])
				if t349 == 0 {
					goto l344
				}
				if uint32(v5) > uint32(i32(2)) {
					goto l345
				}
				m.memory[int64(uint32(v0))+4] = byte(i32(1))
				v3 = i32(0)
				v16 = i32(0)
				goto l346
			l345:
				v16 = i32(0)
				store16(m.memory[int64(uint32(v1))+1:], uint16(i32(0)))
				{
					t350 := int32(m.memory[int64(uint32(v1))+3])
					switch t350 {
					default:
						m.fn3(i32(1274012), i32(40), i32(1145864))
						panic("unreachable")
					case 0, 1:
						t351 := int32(m.memory[int64(uint32(v1))+5])
						m.memory[uint32(v4)] = byte(t351)
						m.memory[int64(uint32(v1))+5] = byte(i32(0))
						v16 = i32(1)
						goto l344
					case 2:
						t352 := int32(m.memory[int64(uint32(v1))+5])
						v13 = t352
						m.memory[int64(uint32(v1))+5] = byte(i32(0))
						m.memory[int64(uint32(v4))+2] = byte(v13&i32(63) | i32(128))
						t353 := v4
						v13 = v13 + i32(-192)
						m.memory[int64(uint32(t353))+1] = byte(int32(uint32(v13)>>6)&i32(63) | i32(128))
						m.memory[uint32(v4)] = byte(int32(uint32(v13&i32(61440))>>12) | i32(224))
						v16 = i32(3)
						goto l344
					case 3:
						m.memory[int64(uint32(v1))+3] = byte(i32(4))
					}
				}
			}
		l344:
			{
				if v3 == 0 {
					goto l351
				}
				t354 := int32(m.memory[int64(uint32(v1))+1])
				v9 = t354
				t355 := int32(m.memory[int64(uint32(v1))+5])
				v18 = t355
				t356 := int32(m.memory[int64(uint32(v1))+3])
				v12 = t356
				v13 = i32(0)
			l422:
				{
					v10 = v13
					v17 = v16 + i32(2)
					if uint32(v17) < uint32(v5) {
						goto l352
					}
					m.memory[int64(uint32(v0))+4] = byte(i32(1))
					v3 = v10
					goto l346
				l352:
					v13 = v10 + i32(1)
					t357 := int32(m.memory[uint32(v2+v10)])
					v11 = t357
					v15 = int32(int8(v11))
					{
						{
							switch v12 & i32(255) {
							case 5:
								switch v11 + i32(-36) {
								case 0, 4:
									v12 = i32(6)
									m.memory[int64(uint32(v1))+3] = byte(i32(6))
									m.memory[int64(uint32(v1))+5] = byte(v15)
									v18 = v15
									goto l368
								default:
									m.memory[int64(uint32(v1))+1] = byte(i32(0))
									m.memory[int64(uint32(v0))+6] = byte(i32(0))
									store16(m.memory[int64(uint32(v0))+4:], uint16(i32(258)))
									t362 := int32(m.memory[int64(uint32(v1))+4])
									m.memory[int64(uint32(v1))+3] = byte(t362)
									v3 = v10
									goto l346
								}
							case 6:
								v17 = v18 & i32(255)
								var p358 int32
								if v17 != i32(40) {
									p358 = 1
								}
								v12 = p358
								if v12 != 0 {
									goto l362
								}
								if v15 != i32(66) {
									goto l362
								}
								v12 = i32(0)
								goto l363
							default:
								if v15 == i32(27) {
									goto l364
								}
								m.memory[int64(uint32(v1))+1] = byte(i32(0))
								if v15 < i32(0) {
									goto l365
								}
								if v15&i32(254) != i32(14) {
									m.memory[uint32(v4+v16)] = byte(v15)
									v16 = v16 + i32(1)
									v12 = i32(0)
									v9 = i32(0)
									goto l368
								}
							l365:
								m.memory[int64(uint32(v0))+6] = byte(i32(0))
								store16(m.memory[int64(uint32(v0))+4:], uint16(i32(258)))
								goto l367
							case 1:
								if v15 == i32(27) {
									goto l364
								}
								m.memory[int64(uint32(v1))+1] = byte(i32(0))
								if v11 == i32(126) {
									v10 = v4 + v16
									store16(m.memory[uint32(v10):], uint16(i32(32994)))
									m.memory[uint32(v10+i32(2))] = byte(i32(190))
									v16 = v16 + i32(3)
									v9 = i32(0)
									v12 = i32(1)
									goto l368
								}
								if v11 != i32(92) {
									if v15 < i32(0) {
										goto l371
									}
									if v15&i32(254) != i32(14) {
										m.memory[uint32(v4+v16)] = byte(v15)
										v12 = i32(1)
										v16 = v16 + i32(1)
										v9 = i32(0)
										goto l368
									}
								l371:
									m.memory[int64(uint32(v0))+6] = byte(i32(0))
									store16(m.memory[int64(uint32(v0))+4:], uint16(i32(258)))
									goto l367
								}
								store16(m.memory[uint32(v4+v16):], uint16(i32(42434)))
								v9 = i32(0)
								v12 = i32(1)
								v16 = v17
								goto l368
							case 2:
								if v15 == i32(27) {
									goto l364
								}
								m.memory[int64(uint32(v1))+1] = byte(i32(0))
								if uint32((v15+i32(-33))&i32(255)) < uint32(i32(63)) {
									v10 = v4 + v16
									m.memory[uint32(v10)] = byte(i32(239))
									v12 = i32(2)
									m.memory[uint32(v10+i32(2))] = byte(v15&i32(63) | i32(128))
									m.memory[uint32(v10+i32(1))] = byte(int32(uint32(v15+i32(16192))>>6) & i32(191))
									v16 = v16 + i32(3)
									v9 = i32(0)
									goto l368
								}
								m.memory[int64(uint32(v0))+6] = byte(i32(0))
								store16(m.memory[int64(uint32(v0))+4:], uint16(i32(258)))
								goto l367
							case 3:
								if v15 == i32(27) {
									goto l364
								}
								m.memory[int64(uint32(v1))+1] = byte(i32(0))
								if uint32((v15+i32(-33))&i32(255)) < uint32(i32(94)) {
									v12 = i32(4)
									m.memory[int64(uint32(v1))+3] = byte(i32(4))
									m.memory[int64(uint32(v1))+5] = byte(v15)
									v9 = i32(0)
									v18 = v15
									goto l368
								}
								m.memory[int64(uint32(v0))+6] = byte(i32(0))
								store16(m.memory[int64(uint32(v0))+4:], uint16(i32(258)))
								goto l367
							case 4:
								if v15 == i32(27) {
									goto l375
								}
								m.memory[int64(uint32(v1))+3] = byte(i32(3))
								v10 = v15 + i32(-33)
								v15 = (v18 + i32(-33)) & i32(255)
								if v15 != i32(3) {
									goto l376
								}
								v12 = v10 & i32(255)
								if uint32(v12) < uint32(i32(83)) {
									v10 = v4 + v16
									m.memory[uint32(v10)] = byte(i32(227))
									t364 := v10 + i32(2)
									v15 = v12 + i32(12353)
									m.memory[uint32(t364)] = byte(v15&i32(63) | i32(128))
									m.memory[uint32(v10+i32(1))] = byte(int32(uint32(v15)>>6) & i32(131))
									v12 = i32(3)
									v16 = v16 + i32(3)
									v18 = i32(36)
									goto l368
								}
							l376:
								if v15 != i32(4) {
									goto l378
								}
								v12 = v10 & i32(255)
								if uint32(v12) < uint32(i32(86)) {
									v10 = v4 + v16
									m.memory[uint32(v10)] = byte(i32(227))
									t365 := v10 + i32(2)
									v15 = v12 + i32(12449)
									m.memory[uint32(t365)] = byte(v15&i32(63) | i32(128))
									m.memory[uint32(v10+i32(1))] = byte(int32(uint32(v15)>>6) & i32(135))
									v12 = i32(3)
									v16 = v16 + i32(3)
									v18 = i32(37)
									goto l368
								}
							l378:
								v10 = v10 & i32(255)
								if uint32(v10) > uint32(i32(93)) {
									m.memory[int64(uint32(v0))+6] = byte(i32(0))
									store16(m.memory[int64(uint32(v0))+4:], uint16(i32(514)))
									goto l367
								}
								v10 = v15*i32(94) + v10
								v15 = v10 + i32(-1410)
								if uint32(v15) < uint32(i32(2965)) {
									v10 = v4 + v16
									t366 := int32(load16(m.memory[int64(uint32(v15<<1))+1234998:]))
									t367 := v10 + i32(2)
									v15 = t366
									m.memory[uint32(t367)] = byte(v15&i32(63) | i32(128))
									m.memory[uint32(v10)] = byte(int32(uint32(v15)>>12) | i32(224))
									m.memory[uint32(v10+i32(1))] = byte(int32(uint32(v15)>>6)&i32(63) | i32(128))
									goto l397
								}
								v15 = v10 + i32(-4418)
								if uint32(v15) < uint32(i32(3390)) {
									v10 = v4 + v16
									t368 := int32(load16(m.memory[int64(uint32(v15<<1))+1263314:]))
									t369 := v10 + i32(2)
									v15 = t368
									m.memory[uint32(t369)] = byte(v15&i32(63) | i32(128))
									m.memory[uint32(v10)] = byte(int32(uint32(v15)>>12) | i32(224))
									m.memory[uint32(v10+i32(1))] = byte(int32(uint32(v15)>>6)&i32(63) | i32(128))
									goto l397
								}
								v15 = v10 + i32(-8272)
								if uint32(v15) < uint32(i32(360)) {
									v10 = v4 + v16
									t359 := int32(load16(m.memory[int64(uint32(v15<<1))+1270094:]))
									t360 := v10 + i32(2)
									v15 = t359
									m.memory[uint32(t360)] = byte(v15&i32(63) | i32(128))
									m.memory[uint32(v10)] = byte(int32(uint32(v15)>>12) | i32(224))
									m.memory[uint32(v10+i32(1))] = byte(int32(uint32(v15)>>6)&i32(63) | i32(128))
									goto l397
								}
								if uint32(v10) >= uint32(i32(108)) {
									v15 = v10 + i32(-119)
									if uint32(v15) >= uint32(i32(8)) {
										v15 = v10 + i32(-135)
										if uint32(v15) >= uint32(i32(7)) {
											v15 = v10 + i32(-153)
											if uint32(v15) >= uint32(i32(15)) {
												v15 = v10 + i32(-175)
												if uint32(v15) >= uint32(i32(8)) {
													if v10 != i32(187) {
														v15 = v10 + i32(-658)
														if uint32(v15) >= uint32(i32(32)) {
															v12 = i32(23)
															v15 = v10 + i32(-1159)
															if uint32(v15) < uint32(i32(23)) {
																goto l385
															}
															v15 = v10 + i32(-1190)
															if uint32(v15) >= uint32(i32(30)) {
																v15 = v10 + i32(-10736)
																if uint32(v15) >= uint32(i32(8)) {
																	v15 = v10 + i32(-8644)
																	if uint32(v15) >= uint32(i32(4)) {
																		v15 = v10 + i32(-203)
																		if uint32(v15) >= uint32(i32(10)) {
																			goto l395
																		}
																		v12 = i32(2)
																		goto l396
																	}
																	v12 = i32(32)
																	goto l385
																}
																v12 = i32(29)
																goto l385
															}
															v12 = i32(26)
															goto l385
														}
														v12 = i32(20)
														goto l385
													}
													v12 = i32(17)
													v15 = i32(0)
													goto l385
												}
												v12 = i32(14)
												goto l385
											}
											v12 = i32(11)
											goto l385
										}
										v12 = i32(8)
										goto l385
									}
									v12 = i32(5)
									goto l385
								}
								v12 = i32(2)
								v15 = v10
								goto l385
							l395:
								v15 = v10 + i32(-220)
								if uint32(v15) >= uint32(i32(26)) {
									goto l398
								}
								v12 = i32(5)
								goto l396
							l398:
								v15 = v10 + i32(-252)
								if uint32(v15) >= uint32(i32(26)) {
									goto l399
								}
								v12 = i32(8)
								goto l396
							l399:
								v15 = v10 + i32(-470)
								if uint32(v15) >= uint32(i32(17)) {
									goto l400
								}
								v12 = i32(11)
								goto l396
							l400:
								v15 = v10 + i32(-487)
								if uint32(v15) >= uint32(i32(7)) {
									goto l401
								}
								v12 = i32(14)
								goto l396
							l401:
								v12 = i32(17)
								v15 = v10 + i32(-502)
								if uint32(v15) < uint32(i32(17)) {
									goto l396
								}
								v15 = v10 + i32(-519)
								if uint32(v15) >= uint32(i32(7)) {
									goto l402
								}
								v12 = i32(20)
								goto l396
							l402:
								v15 = v10 + i32(-564)
								if uint32(v15) >= uint32(i32(6)) {
									goto l403
								}
								v12 = i32(23)
								goto l396
							l403:
								v15 = i32(0)
								if v10 != i32(570) {
									goto l404
								}
								v12 = i32(26)
								goto l396
							l404:
								v11 = v10 + i32(-571)
								if uint32(v11) >= uint32(i32(26)) {
									goto l405
								}
								v12 = i32(29)
								v15 = v11
								goto l396
							l405:
								v11 = v10 + i32(-612)
								if uint32(v11) >= uint32(i32(6)) {
									goto l406
								}
								v12 = i32(32)
								v15 = v11
								goto l396
							l406:
								if v10 != i32(618) {
									goto l407
								}
								v12 = i32(35)
								goto l396
							l407:
								v15 = v10 + i32(-619)
								if uint32(v15) >= uint32(i32(26)) {
									goto l408
								}
								v12 = i32(38)
								goto l396
							l408:
								v15 = v10 + i32(-1128)
								if uint32(v15) >= uint32(i32(20)) {
									goto l409
								}
								v12 = i32(41)
								goto l396
							l409:
								v15 = v10 + i32(-1148)
								if uint32(v15) >= uint32(i32(10)) {
									goto l410
								}
								v12 = i32(44)
								goto l396
							l410:
								v15 = v10 + i32(-8634)
								if uint32(v15) >= uint32(i32(10)) {
									goto l411
								}
								v12 = i32(47)
								goto l396
							l411:
								v15 = v10 + i32(-10716)
								if uint32(v15) >= uint32(i32(10)) {
									goto l412
								}
								v12 = i32(50)
								goto l396
							l412:
								v15 = v10 + i32(-10726)
								if uint32(v15) >= uint32(i32(10)) {
									m.memory[int64(uint32(v0))+6] = byte(i32(0))
									store16(m.memory[int64(uint32(v0))+4:], uint16(i32(514)))
									goto l367
								}
								v12 = i32(53)
							l396:
								v10 = v4 + v16
								{
									t361 := int32(load16(m.memory[int64(uint32(v12<<1))+1241434:]))
									v15 = t361 + v15
									if uint32(v15&i32(0xffff)) < uint32(i32(2048)) {
										m.memory[uint32(v10+i32(1))] = byte(v15&i32(63) | i32(128))
										m.memory[uint32(v10)] = byte(int32(uint32(v15)>>6) | i32(192))
										goto l415
									}
									m.memory[uint32(v10+i32(2))] = byte(v15&i32(63) | i32(128))
									m.memory[uint32(v10+i32(1))] = byte(int32(uint32(v15)>>6)&i32(63) | i32(128))
									m.memory[uint32(v10)] = byte(int32(uint32(v15&i32(61440))>>12) | i32(224))
									goto l397
								}
							}
						l362:
							if v12 != 0 {
								goto l416
							}
							if v15 != i32(74) {
								goto l416
							}
							v12 = i32(1)
							goto l363
						l416:
							if v12 != 0 {
								goto l417
							}
							if v15 != i32(73) {
								goto l417
							}
							v12 = i32(2)
							goto l363
						l417:
							if v17 != i32(36) {
								goto l418
							}
							v12 = i32(3)
							switch v11 + i32(-64) {
							case 0, 2:
								goto l363
							default:
								goto l418
							}
						l418:
							store16(m.memory[int64(uint32(v1))+1:], uint16(i32(256)))
							m.memory[int64(uint32(v0))+6] = byte(i32(1))
							store16(m.memory[int64(uint32(v0))+4:], uint16(i32(258)))
							t363 := int32(m.memory[int64(uint32(v1))+4])
							m.memory[int64(uint32(v1))+3] = byte(t363)
							v3 = v10
							goto l346
						}
					l363:
						m.memory[int64(uint32(v1))+4] = byte(v12)
						m.memory[int64(uint32(v1))+3] = byte(v12)
						v18 = i32(0)
						m.memory[int64(uint32(v1))+5] = byte(i32(0))
						m.memory[int64(uint32(v1))+1] = byte(i32(1))
						if v9&i32(1) != 0 {
							m.memory[int64(uint32(v0))+6] = byte(i32(3))
							store16(m.memory[int64(uint32(v0))+4:], uint16(i32(770)))
							goto l367
						}
						v9 = i32(1)
						goto l368
					l385:
						{
							t370 := int32(load16(m.memory[int64(uint32(v12<<1))+1241972:]))
							v15 = v15 + t370
							if uint32(v15) < uint32(i32(240)) {
								goto l420
							}
							m.fn33(v15, i32(240), i32(1241544))
							panic("unreachable")
						}
					l420:
						v10 = v4 + v16
						t371 := int32(load16(m.memory[int64(uint32(v15<<1))+1226508:]))
						v15 = t371
						if uint32(v15) < uint32(i32(2048)) {
							goto l421
						}
						m.memory[uint32(v10+i32(2))] = byte(v15&i32(63) | i32(128))
						m.memory[uint32(v10)] = byte(int32(uint32(v15)>>12) | i32(224))
						m.memory[uint32(v10+i32(1))] = byte(int32(uint32(v15)>>6)&i32(63) | i32(128))
					}
				l397:
					v12 = i32(3)
					v16 = v16 + i32(3)
					goto l368
				l421:
					m.memory[uint32(v10+i32(1))] = byte(v15&i32(63) | i32(128))
					m.memory[uint32(v10)] = byte(int32(uint32(v15)>>6) | i32(192))
				l415:
					v12 = i32(3)
					v16 = v17
					goto l368
				l375:
					m.memory[int64(uint32(v0))+6] = byte(i32(1))
					store16(m.memory[int64(uint32(v0))+4:], uint16(i32(258)))
					m.memory[int64(uint32(v1))+3] = byte(i32(5))
				l367:
					v3 = v13
					goto l346
				l364:
					v12 = i32(5)
					m.memory[int64(uint32(v1))+3] = byte(i32(5))
				l368:
					if v3 != v13 {
						goto l422
					}
				}
			}
		l351:
			if v6 == 0 {
				goto l423
			}
			{
				t372 := int32(m.memory[int64(uint32(v1))+3])
				switch t372 + i32(-4) {
				default:
					goto l423
				case 0, 1:
					m.memory[int64(uint32(v0))+6] = byte(i32(0))
					store16(m.memory[int64(uint32(v0))+4:], uint16(i32(258)))
					t373 := int32(m.memory[int64(uint32(v1))+4])
					m.memory[int64(uint32(v1))+3] = byte(t373)
					goto l346
				case 2:
					m.memory[int64(uint32(v1))+2] = byte(i32(1))
					m.memory[int64(uint32(v0))+6] = byte(i32(1))
					store16(m.memory[int64(uint32(v0))+4:], uint16(i32(258)))
					t374 := int32(m.memory[int64(uint32(v1))+4])
					m.memory[int64(uint32(v1))+3] = byte(t374)
					goto l346
				}
			}
		l423:
			m.memory[int64(uint32(v0))+4] = byte(i32(0))
		l346:
			store32(m.memory[int64(uint32(v0))+8:], uint32(v16))
			store32(m.memory[uint32(v0):], uint32(v3))
			goto l25
		case 4:
			store32(m.memory[int64(uint32(v7))+56:], uint32(v5))
			store32(m.memory[int64(uint32(v7))+52:], uint32(v4))
			v15 = i32(0)
			v17 = i32(0)
			{
				t375 := int32(m.memory[int64(uint32(v1))+1])
				v13 = t375
				if v13 == 0 {
					goto l596
				}
				{
					{
						if v3 == 0 {
							goto l427
						}
						if uint32(v5) < uint32(i32(3)) {
							store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
							store32(m.memory[uint32(v0):], uint32(i32(0)))
							m.memory[int64(uint32(v0))+4] = byte(i32(1))
							goto l25
						}
						t376 := int32(m.memory[int64(uint32(v1))+2])
						v16 = t376
						t377 := int32(int8(m.memory[uint32(v2)]))
						v10 = t377
						v15 = i32(1)
						v12 = i32(0)
						switch v13 + i32(-1) {
						case 3:
							m.memory[int64(uint32(v1))+1] = byte(i32(0))
							{
								v13 = (v10 + i32(95)) & i32(255)
								if uint32(v13) > uint32(i32(62)) {
									m.memory[int64(uint32(v0))+4] = byte(i32(2))
									if v10 > i32(-1) {
										store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
										store32(m.memory[uint32(v0):], uint32(i32(0)))
										store16(m.memory[int64(uint32(v0))+5:], uint16(i32(1)))
										goto l25
									}
									store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
									store32(m.memory[uint32(v0):], uint32(i32(1)))
									store16(m.memory[int64(uint32(v0))+5:], uint16(i32(2)))
									goto l25
								}
								m.memory[uint32(v4)] = byte(i32(239))
								t386 := v4
								v13 = v13 + i32(-159)
								m.memory[int64(uint32(t386))+2] = byte(v13&i32(63) | i32(128))
								m.memory[int64(uint32(v4))+1] = byte(int32(uint32(v13)>>6) & i32(191))
								goto l496
							}
						default:
							m.memory[int64(uint32(v1))+1] = byte(i32(0))
							v13 = v10 + i32(95)
							if v16 != i32(3) {
								goto l439
							}
							v15 = v13 & i32(255)
							if uint32(v15) < uint32(i32(83)) {
								m.memory[uint32(v4)] = byte(i32(227))
								t387 := v4
								v13 = v15 + i32(12353)
								m.memory[int64(uint32(t387))+1] = byte(int32(uint32(v13)>>6) & i32(131))
								m.memory[int64(uint32(v4))+2] = byte(v13&i32(63) | i32(128))
								store32(m.memory[int64(uint32(v7))+60:], uint32(i32(2)))
								goto l496
							}
						l439:
							if v16 != i32(4) {
								goto l441
							}
							v15 = v13 & i32(255)
							if uint32(v15) < uint32(i32(86)) {
								m.memory[uint32(v4)] = byte(i32(227))
								t388 := v4
								v13 = v15 + i32(12449)
								m.memory[int64(uint32(t388))+1] = byte(int32(uint32(v13)>>6) & i32(135))
								m.memory[int64(uint32(v4))+2] = byte(v13&i32(63) | i32(128))
								store32(m.memory[int64(uint32(v7))+60:], uint32(i32(2)))
								goto l496
							}
						l441:
							v13 = v13 & i32(255)
							if uint32(v13) > uint32(i32(93)) {
								m.memory[int64(uint32(v0))+4] = byte(i32(2))
								if v10 > i32(-1) {
									store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
									store32(m.memory[uint32(v0):], uint32(i32(0)))
									store16(m.memory[int64(uint32(v0))+5:], uint16(i32(1)))
									goto l25
								}
								store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
								store32(m.memory[uint32(v0):], uint32(i32(1)))
								store16(m.memory[int64(uint32(v0))+5:], uint16(i32(2)))
								goto l25
							}
							v13 = v16*i32(94) + v13
							v10 = v13 + i32(-1410)
							if uint32(v10) < uint32(i32(2965)) {
								v15 = i32(1)
								t389 := int32(load16(m.memory[int64(uint32(v10<<1))+1234998:]))
								t390 := v4
								v13 = t389
								m.memory[uint32(t390)] = byte(int32(uint32(v13)>>12) | i32(224))
								m.memory[int64(uint32(v4))+2] = byte(v13&i32(63) | i32(128))
								m.memory[int64(uint32(v4))+1] = byte(int32(uint32(v13)>>6)&i32(63) | i32(128))
								store32(m.memory[int64(uint32(v7))+60:], uint32(i32(2)))
								v17 = i32(3)
								goto l596
							}
							v10 = v13 + i32(-4418)
							if uint32(v10) < uint32(i32(3390)) {
								v15 = i32(1)
								t391 := int32(load16(m.memory[int64(uint32(v10<<1))+1263314:]))
								t392 := v4
								v13 = t391
								m.memory[int64(uint32(t392))+2] = byte(v13&i32(63) | i32(128))
								m.memory[uint32(v4)] = byte(int32(uint32(v13)>>12) | i32(224))
								m.memory[int64(uint32(v4))+1] = byte(int32(uint32(v13)>>6)&i32(63) | i32(128))
								goto l459
							}
							v10 = v13 + i32(-8272)
							if uint32(v10) < uint32(i32(360)) {
								v15 = i32(1)
								t381 := int32(load16(m.memory[int64(uint32(v10<<1))+1270094:]))
								t382 := v4
								v13 = t381
								m.memory[int64(uint32(t382))+2] = byte(v13&i32(63) | i32(128))
								m.memory[uint32(v4)] = byte(int32(uint32(v13)>>12) | i32(224))
								m.memory[int64(uint32(v4))+1] = byte(int32(uint32(v13)>>6)&i32(63) | i32(128))
								goto l459
							}
							if uint32(v13) >= uint32(i32(108)) {
								v15 = i32(8)
								v10 = v13 + i32(-119)
								if uint32(v10) >= uint32(i32(8)) {
									v10 = v13 + i32(-135)
									if uint32(v10) < uint32(i32(7)) {
										goto l448
									}
									v10 = v13 + i32(-153)
									if uint32(v10) >= uint32(i32(15)) {
										v10 = v13 + i32(-175)
										if uint32(v10) >= uint32(i32(8)) {
											if v13 != i32(187) {
												v10 = v13 + i32(-658)
												if uint32(v10) >= uint32(i32(32)) {
													v15 = i32(23)
													v10 = v13 + i32(-1159)
													if uint32(v10) < uint32(i32(23)) {
														goto l448
													}
													v10 = v13 + i32(-1190)
													if uint32(v10) >= uint32(i32(30)) {
														v10 = v13 + i32(-10736)
														if uint32(v10) >= uint32(i32(8)) {
															v10 = v13 + i32(-8644)
															if uint32(v10) >= uint32(i32(4)) {
																v10 = v13 + i32(-203)
																if uint32(v10) >= uint32(i32(10)) {
																	goto l457
																}
																v16 = i32(2)
																goto l458
															}
															v15 = i32(32)
															goto l448
														}
														v15 = i32(29)
														goto l448
													}
													v15 = i32(26)
													goto l448
												}
												v15 = i32(20)
												goto l448
											}
											v15 = i32(17)
											v10 = i32(0)
											goto l448
										}
										v15 = i32(14)
										goto l448
									}
									v15 = i32(11)
									goto l448
								}
								v15 = i32(5)
								goto l448
							}
							v15 = i32(2)
							v10 = v13
							goto l448
						l457:
							v10 = v13 + i32(-220)
							if uint32(v10) >= uint32(i32(26)) {
								goto l460
							}
							v16 = i32(5)
							goto l458
						l460:
							v10 = v13 + i32(-252)
							if uint32(v10) >= uint32(i32(26)) {
								goto l461
							}
							v16 = i32(8)
							goto l458
						l461:
							v10 = v13 + i32(-470)
							if uint32(v10) >= uint32(i32(17)) {
								goto l462
							}
							v16 = i32(11)
							goto l458
						l462:
							v10 = v13 + i32(-487)
							if uint32(v10) >= uint32(i32(7)) {
								goto l463
							}
							v16 = i32(14)
							goto l458
						l463:
							v16 = i32(17)
							v10 = v13 + i32(-502)
							if uint32(v10) < uint32(i32(17)) {
								goto l458
							}
							v10 = v13 + i32(-519)
							if uint32(v10) >= uint32(i32(7)) {
								goto l464
							}
							v16 = i32(20)
							goto l458
						l464:
							v10 = v13 + i32(-564)
							if uint32(v10) >= uint32(i32(6)) {
								goto l465
							}
							v16 = i32(23)
							goto l458
						l465:
							v10 = i32(0)
							if v13 != i32(570) {
								goto l466
							}
							v16 = i32(26)
							goto l458
						l466:
							v15 = v13 + i32(-571)
							if uint32(v15) >= uint32(i32(26)) {
								goto l467
							}
							v16 = i32(29)
							v10 = v15
							goto l458
						l467:
							v15 = v13 + i32(-612)
							if uint32(v15) >= uint32(i32(6)) {
								goto l468
							}
							v16 = i32(32)
							v10 = v15
							goto l458
						l468:
							if v13 != i32(618) {
								goto l469
							}
							v16 = i32(35)
							goto l458
						l469:
							v10 = v13 + i32(-619)
							if uint32(v10) >= uint32(i32(26)) {
								goto l470
							}
							v16 = i32(38)
							goto l458
						l470:
							v10 = v13 + i32(-1128)
							if uint32(v10) >= uint32(i32(20)) {
								goto l471
							}
							v16 = i32(41)
							goto l458
						l471:
							v10 = v13 + i32(-1148)
							if uint32(v10) >= uint32(i32(10)) {
								goto l472
							}
							v16 = i32(44)
							goto l458
						l472:
							v10 = v13 + i32(-8634)
							if uint32(v10) >= uint32(i32(10)) {
								goto l473
							}
							v16 = i32(47)
							goto l458
						l473:
							v10 = v13 + i32(-10716)
							if uint32(v10) >= uint32(i32(10)) {
								goto l474
							}
							v16 = i32(50)
							goto l458
						l474:
							v10 = v13 + i32(-10726)
							if uint32(v10) >= uint32(i32(10)) {
								m.memory[int64(uint32(v0))+6] = byte(i32(0))
								store16(m.memory[int64(uint32(v0))+4:], uint16(i32(514)))
								store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
								store32(m.memory[uint32(v0):], uint32(i32(1)))
								goto l25
							}
							v16 = i32(53)
						l458:
							v15 = i32(1)
							{
								t383 := int32(load16(m.memory[int64(uint32(v16<<1))+1241434:]))
								v13 = t383 + v10
								if uint32(v13&i32(0xffff)) < uint32(i32(2048)) {
									m.memory[int64(uint32(v4))+1] = byte(v13&i32(63) | i32(128))
									m.memory[uint32(v4)] = byte(int32(uint32(v13)>>6) | i32(192))
									goto l477
								}
								m.memory[int64(uint32(v4))+2] = byte(v13&i32(63) | i32(128))
								m.memory[int64(uint32(v4))+1] = byte(int32(uint32(v13)>>6)&i32(63) | i32(128))
								m.memory[uint32(v4)] = byte(int32(uint32(v13&i32(61440))>>12) | i32(224))
								goto l459
							}
						case 1:
							m.memory[int64(uint32(v1))+1] = byte(i32(0))
							v16 = v10 + i32(95)
							if uint32(v16&i32(255)) > uint32(i32(93)) {
								m.memory[int64(uint32(v0))+4] = byte(i32(2))
								if v10 > i32(-1) {
									store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
									store32(m.memory[uint32(v0):], uint32(i32(0)))
									store16(m.memory[int64(uint32(v0))+5:], uint16(i32(1)))
									goto l25
								}
								store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
								store32(m.memory[uint32(v0):], uint32(i32(1)))
								store16(m.memory[int64(uint32(v0))+5:], uint16(i32(2)))
								goto l25
							}
							m.memory[int64(uint32(v1))+2] = byte(v16)
							v13 = i32(3)
							m.memory[int64(uint32(v1))+1] = byte(i32(3))
							v12 = i32(1)
							if v3 == i32(1) {
								goto l427
							}
							t378 := int32(m.memory[int64(uint32(v2))+1])
							v10 = t378
							v15 = i32(2)
							fallthrough
						case 2:
							m.memory[int64(uint32(v1))+1] = byte(i32(0))
							v13 = (v10 + i32(95)) & i32(255)
							if uint32(v13) > uint32(i32(93)) {
								m.memory[int64(uint32(v0))+4] = byte(i32(2))
								if int32(int8(v10)) > i32(-1) {
									store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
									store32(m.memory[uint32(v0):], uint32(v12))
									store16(m.memory[int64(uint32(v0))+5:], uint16(i32(2)))
									goto l25
								}
								store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
								store32(m.memory[uint32(v0):], uint32(v15))
								store16(m.memory[int64(uint32(v0))+5:], uint16(i32(3)))
								goto l25
							}
							v13 = v16&i32(255)*i32(94) + v13
							v10 = v13 + i32(-1410)
							if uint32(v10) < uint32(i32(5801)) {
								t393 := int32(load16(m.memory[int64(uint32(v10<<1))+1197886:]))
								t394 := v4
								v13 = t393
								m.memory[int64(uint32(t394))+2] = byte(v13&i32(63) | i32(128))
								m.memory[uint32(v4)] = byte(int32(uint32(v13)>>12) | i32(224))
								m.memory[int64(uint32(v4))+1] = byte(int32(uint32(v13)>>6)&i32(63) | i32(128))
								goto l459
							}
							v10 = v13 + i32(-108)
							if uint32(v10) >= uint32(i32(11)) {
								goto l436
							}
							v16 = i32(2)
							goto l437
						}
					}
				l427:
					if v6 != 0 {
						m.memory[int64(uint32(v0))+6] = byte(i32(0))
						m.memory[int64(uint32(v0))+4] = byte(i32(2))
						m.memory[int64(uint32(v1))+1] = byte(i32(0))
						store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
						store32(m.memory[uint32(v0):], uint32(v3))
						t380 := v0
						p379 := i32(1)
						if v13 == i32(3) {
							p379 = i32(2)
						}
						m.memory[int64(uint32(t380))+5] = byte(p379)
						goto l25
					}
					store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
					store32(m.memory[uint32(v0):], uint32(v3))
					m.memory[int64(uint32(v0))+4] = byte(i32(0))
					goto l25
				l436:
					v10 = v13 + i32(-127)
					if uint32(v10) >= uint32(i32(3)) {
						goto l480
					}
					v16 = i32(5)
					goto l437
				l480:
					v10 = v13 + i32(-168)
					if uint32(v10) >= uint32(i32(7)) {
						goto l481
					}
					v16 = i32(8)
					goto l437
				l481:
					v10 = v13 + i32(-534)
					if uint32(v10) >= uint32(i32(12)) {
						goto l482
					}
					v16 = i32(11)
					goto l437
				l482:
					v10 = v13 + i32(-550)
					if uint32(v10) >= uint32(i32(12)) {
						goto l483
					}
					v16 = i32(14)
					goto l437
				l483:
					v10 = v13 + i32(-608)
					if uint32(v10) >= uint32(i32(2)) {
						goto l484
					}
					v16 = i32(17)
					goto l437
				l484:
					v10 = v13 + i32(-656)
					if uint32(v10) >= uint32(i32(2)) {
						goto l485
					}
					v16 = i32(20)
					goto l437
				l485:
					v10 = v13 + i32(-752)
					if uint32(v10) >= uint32(i32(16)) {
						goto l486
					}
					v16 = i32(23)
					goto l437
				l486:
					v10 = v13 + i32(-784)
					if uint32(v10) >= uint32(i32(16)) {
						goto l487
					}
					v16 = i32(26)
					goto l437
				l487:
					v10 = v13 + i32(-846)
					if uint32(v10) >= uint32(i32(87)) {
						goto l488
					}
					v16 = i32(29)
					goto l437
				l488:
					v10 = v13 + i32(-940)
					if uint32(v10) > uint32(i32(86)) {
						goto l489
					}
					v16 = i32(32)
				l437:
					t384 := int32(load16(m.memory[int64(uint32(v16<<1))+1244370:]))
					v10 = v10 + t384
					if uint32(v10) >= uint32(i32(255)) {
						m.fn33(v10, i32(255), i32(1242196))
						panic("unreachable")
					}
					t385 := int32(load16(m.memory[int64(uint32(v10<<1))+1227352:]))
					v10 = t385
					if v10 == 0 {
						goto l489
					}
					if uint32(v10) < uint32(i32(2048)) {
						m.memory[int64(uint32(v4))+1] = byte(v10&i32(63) | i32(128))
						m.memory[uint32(v4)] = byte(int32(uint32(v10)>>6) | i32(192))
						goto l477
					}
					m.memory[int64(uint32(v4))+2] = byte(v10&i32(63) | i32(128))
					m.memory[uint32(v4)] = byte(int32(uint32(v10)>>12) | i32(224))
					m.memory[int64(uint32(v4))+1] = byte(int32(uint32(v10)>>6)&i32(63) | i32(128))
					goto l459
				}
			l489:
				v10 = v13 + i32(-597)
				if uint32(v10) < uint32(i32(11)) {
					m.memory[uint32(v4)] = byte(i32(208))
					m.memory[int64(uint32(v4))+1] = byte(v10 + i32(-126))
					goto l477
				}
				v13 = v13 + i32(-645)
				if uint32(v13) < uint32(i32(11)) {
					m.memory[uint32(v4)] = byte(i32(209))
					m.memory[int64(uint32(v4))+1] = byte(v13 + i32(-110))
					goto l477
				}
				m.memory[int64(uint32(v0))+6] = byte(i32(0))
				store16(m.memory[int64(uint32(v0))+4:], uint16(i32(770)))
				store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
				store32(m.memory[uint32(v0):], uint32(v15))
				goto l25
			l496:
				v17 = i32(3)
				v15 = i32(1)
				goto l596
			l477:
				v17 = i32(2)
				store32(m.memory[int64(uint32(v7))+60:], uint32(i32(2)))
				goto l596
			l459:
				v17 = i32(3)
				store32(m.memory[int64(uint32(v7))+60:], uint32(i32(3)))
				goto l596
			l448:
				{
					t395 := int32(load16(m.memory[int64(uint32(v15<<1))+1241972:]))
					v13 = v10 + t395
					if uint32(v13) < uint32(i32(240)) {
						goto l498
					}
					m.fn33(v13, i32(240), i32(1241544))
					panic("unreachable")
				}
			l498:
				{
					t396 := int32(load16(m.memory[int64(uint32(v13<<1))+1226508:]))
					v13 = t396
					if uint32(v13) < uint32(i32(2048)) {
						goto l499
					}
					m.memory[int64(uint32(v4))+2] = byte(v13&i32(63) | i32(128))
					m.memory[uint32(v4)] = byte(int32(uint32(v13)>>12) | i32(224))
					m.memory[int64(uint32(v4))+1] = byte(int32(uint32(v13)>>6)&i32(63) | i32(128))
					v17 = i32(3)
					goto l500
				}
			l499:
				m.memory[int64(uint32(v4))+1] = byte(v13&i32(63) | i32(128))
				m.memory[uint32(v4)] = byte(int32(uint32(v13)>>6) | i32(192))
				v17 = i32(2)
			l500:
				store32(m.memory[int64(uint32(v7))+60:], uint32(v17))
				v15 = i32(1)
			}
		l596:
			{
				v13 = v5 - v17
				t397 := v13
				v10 = v3 - v15
				t398 := v10
				var p399 int32
				if uint32(v13) < uint32(v10) {
					p399 = 1
				}
				v14 = p399
				p400 := t398
				if v14 != 0 {
					p400 = t397
				}
				v12 = p400
				v13 = i32(0)
				{
					{
						v11 = v4 + v17
						t401 := v11
						v16 = v2 + v15
						if (t401^v16)&i32(3) != 0 {
							goto l501
						}
						v13 = i32(0)
						v18 = (i32(0) - v16) & i32(3)
						if uint32(v18|i32(8)) > uint32(v12) {
							goto l501
						}
						{
							if v18 != 0 {
								goto l502
							}
							v13 = i32(0)
							v8 = v12 + i32(-8)
							goto l511
						l502:
							v13 = i32(0)
							t402 := int32(int8(m.memory[uint32(v16)]))
							v10 = t402
							if v10 < i32(0) {
								goto l504
							}
							m.memory[uint32(v11)] = byte(v10)
							v13 = i32(1)
							if v18 == i32(1) {
								goto l505
							}
							{
								t403 := int32(int8(m.memory[int64(uint32(v16))+1]))
								v10 = t403
								if v10 >= i32(0) {
									goto l506
								}
								v13 = i32(1)
								goto l504
							}
						l506:
							m.memory[int64(uint32(v11))+1] = byte(v10)
							v13 = i32(2)
							if v18 == i32(2) {
								goto l505
							}
							{
								t404 := int32(int8(m.memory[int64(uint32(v16))+2]))
								v10 = t404
								if v10 >= i32(0) {
									goto l507
								}
								v13 = i32(2)
								goto l504
							}
						l507:
							m.memory[int64(uint32(v11))+2] = byte(v10)
							v13 = i32(3)
						l505:
							v8 = v12 + i32(-8)
						}
					l511:
						{
							v18 = v16 + v13
							t405 := int32(load32(m.memory[uint32(v18):]))
							v10 = t405
							v9 = v11 + v13
							t406 := int32(load32(m.memory[uint32(v18+i32(4)):]))
							t407 := v9 + i32(4)
							v18 = t406
							store32(m.memory[uint32(t407):], uint32(v18))
							store32(m.memory[uint32(v9):], uint32(v10))
							{
								v18 = v18 & i32(-2139062144)
								t408 := v18
								v10 = v10 & i32(-2139062144)
								if t408|v10 == 0 {
									goto l508
								}
								if v10 != 0 {
									goto l509
								}
								v10 = int32(uint32(int32(bits.TrailingZeros32(uint32(v18))))>>3) + i32(4)
								goto l510
							l509:
								v10 = int32(uint32(int32(bits.TrailingZeros32(uint32(v10)))) >> 3)
							l510:
								t409 := v16
								v13 = v10 + v13
								t410 := int32(m.memory[uint32(t409+v13)])
								v10 = t410
								goto l504
							}
						l508:
							v13 = v13 + i32(8)
							if uint32(v13) <= uint32(v8) {
								goto l511
							}
						}
					}
				l501:
					if uint32(v13) >= uint32(v12) {
						goto l512
					}
				l513:
					{
						t411 := int32(int8(m.memory[uint32(v16+v13)]))
						v10 = t411
						if v10 < i32(0) {
							goto l504
						}
						m.memory[uint32(v11+v13)] = byte(v10)
						t412 := v12
						v13 = v13 + i32(1)
						if t412 != v13 {
							goto l513
						}
					}
				l512:
					v16 = v12 + v17
					v13 = v12 + v15
					goto l514
				l504:
					t413 := v7
					v16 = v13 + v17
					store32(m.memory[int64(uint32(t413))+60:], uint32(v16))
					v13 = v13 + v15
					if uint32(v16+i32(2)) < uint32(v5) {
						goto l515
					}
					v14 = i32(1)
				}
			l514:
				store32(m.memory[int64(uint32(v0))+8:], uint32(v16))
				store32(m.memory[uint32(v0):], uint32(v13))
				store32(m.memory[int64(uint32(v0))+4:], uint32(v14|(v7+i32(52))&i32(-256)))
				goto l25
			l515:
				v15 = v13 + i32(1)
			l595:
				{
					v5 = v10 + i32(95)
					v13 = v5 & i32(255)
					if uint32(v13) < uint32(i32(94)) {
						if uint32(v15) < uint32(v3) {
							v5 = v15 + i32(1)
							t433 := int32(int8(m.memory[uint32(v2+v15)]))
							v11 = t433
							v10 = v11 + i32(95)
							if v13 != i32(3) {
								goto l555
							}
							v12 = v10 & i32(255)
							if uint32(v12) < uint32(i32(83)) {
								m.memory[uint32(v4+v16)] = byte(i32(227))
								t444 := int32(load32(m.memory[int64(uint32(v7))+60:]))
								t445 := v7
								v13 = t444
								v10 = v13 + i32(1)
								store32(m.memory[int64(uint32(t445))+60:], uint32(v10))
								t446 := int32(load32(m.memory[int64(uint32(v7))+52:]))
								t447 := v13
								v4 = t446
								t448 := t447 + v4 + i32(2)
								v15 = v12 + i32(12353)
								m.memory[uint32(t448)] = byte(v15&i32(63) | i32(128))
								m.memory[uint32(v4+v10)] = byte(int32(uint32(v15)>>6) & i32(131))
								v16 = v13 + i32(3)
								goto l543
							}
						l555:
							if v13 != i32(4) {
								goto l557
							}
							v12 = v10 & i32(255)
							if uint32(v12) < uint32(i32(86)) {
								m.memory[uint32(v4+v16)] = byte(i32(227))
								t449 := int32(load32(m.memory[int64(uint32(v7))+60:]))
								t450 := v7
								v13 = t449
								v10 = v13 + i32(1)
								store32(m.memory[int64(uint32(t450))+60:], uint32(v10))
								t451 := int32(load32(m.memory[int64(uint32(v7))+52:]))
								t452 := v13
								v4 = t451
								t453 := t452 + v4 + i32(2)
								v15 = v12 + i32(12449)
								m.memory[uint32(t453)] = byte(v15&i32(63) | i32(128))
								m.memory[uint32(v4+v10)] = byte(int32(uint32(v15)>>6) & i32(135))
								v16 = v13 + i32(3)
								goto l543
							}
						l557:
							v10 = v10 & i32(255)
							if uint32(v10) > uint32(i32(93)) {
								m.memory[int64(uint32(v0))+4] = byte(i32(2))
								if v11 > i32(-1) {
									store32(m.memory[int64(uint32(v0))+8:], uint32(v16))
									store32(m.memory[uint32(v0):], uint32(v15))
									store16(m.memory[int64(uint32(v0))+5:], uint16(i32(1)))
									goto l25
								}
								store32(m.memory[int64(uint32(v0))+8:], uint32(v16))
								store32(m.memory[uint32(v0):], uint32(v5))
								store16(m.memory[int64(uint32(v0))+5:], uint16(i32(2)))
								goto l25
							}
							v13 = v13*i32(94) + v10
							v10 = v13 + i32(-1410)
							if uint32(v10) < uint32(i32(2965)) {
								t454 := int32(load16(m.memory[int64(uint32(v10<<1))+1234998:]))
								t455 := v4 + v16
								v13 = t454
								m.memory[uint32(t455)] = byte(int32(uint32(v13)>>12) | i32(224))
								t456 := int32(load32(m.memory[int64(uint32(v7))+60:]))
								t457 := v7
								v10 = t456
								v15 = v10 + i32(1)
								store32(m.memory[int64(uint32(t457))+60:], uint32(v15))
								t458 := int32(load32(m.memory[int64(uint32(v7))+52:]))
								t459 := v10
								v4 = t458
								m.memory[uint32(t459+v4+i32(2))] = byte(v13&i32(63) | i32(128))
								m.memory[uint32(v4+v15)] = byte(int32(uint32(v13)>>6)&i32(63) | i32(128))
								v16 = v10 + i32(3)
								goto l543
							}
							v10 = v13 + i32(-4418)
							if uint32(v10) < uint32(i32(3390)) {
								t460 := int32(load16(m.memory[int64(uint32(v10<<1))+1263314:]))
								t461 := v4 + v16
								v13 = t460
								m.memory[uint32(t461)] = byte(int32(uint32(v13)>>12) | i32(224))
								t462 := int32(load32(m.memory[int64(uint32(v7))+60:]))
								t463 := v7
								v10 = t462
								v15 = v10 + i32(1)
								store32(m.memory[int64(uint32(t463))+60:], uint32(v15))
								t464 := int32(load32(m.memory[int64(uint32(v7))+52:]))
								t465 := v10
								v4 = t464
								m.memory[uint32(t465+v4+i32(2))] = byte(v13&i32(63) | i32(128))
								m.memory[uint32(v4+v15)] = byte(int32(uint32(v13)>>6)&i32(63) | i32(128))
								v16 = v10 + i32(3)
								goto l543
							}
							v10 = v13 + i32(-8272)
							if uint32(v10) < uint32(i32(360)) {
								t434 := int32(load16(m.memory[int64(uint32(v10<<1))+1270094:]))
								t435 := v4 + v16
								v13 = t434
								m.memory[uint32(t435)] = byte(int32(uint32(v13)>>12) | i32(224))
								t436 := int32(load32(m.memory[int64(uint32(v7))+60:]))
								t437 := v7
								v10 = t436
								v15 = v10 + i32(1)
								store32(m.memory[int64(uint32(t437))+60:], uint32(v15))
								t438 := int32(load32(m.memory[int64(uint32(v7))+52:]))
								t439 := v10
								v4 = t438
								m.memory[uint32(t439+v4+i32(2))] = byte(v13&i32(63) | i32(128))
								m.memory[uint32(v4+v15)] = byte(int32(uint32(v13)>>6)&i32(63) | i32(128))
								v16 = v10 + i32(3)
								goto l543
							}
							if uint32(v13) >= uint32(i32(108)) {
								v10 = v13 + i32(-119)
								if uint32(v10) >= uint32(i32(8)) {
									v10 = v13 + i32(-135)
									if uint32(v10) >= uint32(i32(7)) {
										v10 = v13 + i32(-153)
										if uint32(v10) >= uint32(i32(15)) {
											v10 = v13 + i32(-175)
											if uint32(v10) >= uint32(i32(8)) {
												if v13 != i32(187) {
													v10 = v13 + i32(-658)
													if uint32(v10) >= uint32(i32(32)) {
														v15 = i32(23)
														v10 = v13 + i32(-1159)
														if uint32(v10) < uint32(i32(23)) {
															goto l564
														}
														v10 = v13 + i32(-1190)
														if uint32(v10) >= uint32(i32(30)) {
															v10 = v13 + i32(-8644)
															if uint32(v10) >= uint32(i32(4)) {
																v10 = v13 + i32(-203)
																if uint32(v10) >= uint32(i32(10)) {
																	goto l573
																}
																v12 = i32(2)
																goto l574
															}
															v15 = i32(32)
															goto l564
														}
														v15 = i32(26)
														goto l564
													}
													v15 = i32(20)
													goto l564
												}
												v15 = i32(17)
												v10 = i32(0)
												goto l564
											}
											v15 = i32(14)
											goto l564
										}
										v15 = i32(11)
										goto l564
									}
									v15 = i32(8)
									goto l564
								}
								v15 = i32(5)
								goto l564
							}
							v15 = i32(2)
							v10 = v13
							goto l564
						l573:
							v10 = v13 + i32(-220)
							if uint32(v10) >= uint32(i32(26)) {
								goto l575
							}
							v12 = i32(5)
							goto l574
						l575:
							v10 = v13 + i32(-252)
							if uint32(v10) >= uint32(i32(26)) {
								goto l576
							}
							v12 = i32(8)
							goto l574
						l576:
							v10 = v13 + i32(-470)
							if uint32(v10) >= uint32(i32(17)) {
								goto l577
							}
							v12 = i32(11)
							goto l574
						l577:
							v10 = v13 + i32(-487)
							if uint32(v10) >= uint32(i32(7)) {
								goto l578
							}
							v12 = i32(14)
							goto l574
						l578:
							v12 = i32(17)
							v10 = v13 + i32(-502)
							if uint32(v10) < uint32(i32(17)) {
								goto l574
							}
							v10 = v13 + i32(-519)
							if uint32(v10) >= uint32(i32(7)) {
								goto l579
							}
							v12 = i32(20)
							goto l574
						l579:
							v10 = v13 + i32(-564)
							if uint32(v10) >= uint32(i32(6)) {
								goto l580
							}
							v12 = i32(23)
							goto l574
						l580:
							v10 = i32(0)
							if v13 != i32(570) {
								goto l581
							}
							v12 = i32(26)
							goto l574
						l581:
							v15 = v13 + i32(-571)
							if uint32(v15) >= uint32(i32(26)) {
								goto l582
							}
							v12 = i32(29)
							v10 = v15
							goto l574
						l582:
							v15 = v13 + i32(-612)
							if uint32(v15) >= uint32(i32(6)) {
								goto l583
							}
							v12 = i32(32)
							v10 = v15
							goto l574
						l583:
							if v13 != i32(618) {
								goto l584
							}
							v12 = i32(35)
							goto l574
						l584:
							v10 = v13 + i32(-619)
							if uint32(v10) >= uint32(i32(26)) {
								goto l585
							}
							v12 = i32(38)
							goto l574
						l585:
							v10 = v13 + i32(-1128)
							if uint32(v10) >= uint32(i32(20)) {
								goto l586
							}
							v12 = i32(41)
							goto l574
						l586:
							v10 = v13 + i32(-1148)
							if uint32(v10) >= uint32(i32(10)) {
								goto l587
							}
							v12 = i32(44)
							goto l574
						l587:
							v10 = v13 + i32(-8634)
							if uint32(v10) >= uint32(i32(10)) {
								m.memory[int64(uint32(v0))+6] = byte(i32(0))
								store16(m.memory[int64(uint32(v0))+4:], uint16(i32(514)))
								goto l546
							}
							v12 = i32(47)
						l574:
							v11 = v16 + i32(1)
							v15 = v4 + v16
							{
								t440 := int32(load16(m.memory[int64(uint32(v12<<1))+1241434:]))
								v13 = t440 + v10
								if uint32(v13&i32(0xffff)) < uint32(i32(2048)) {
									m.memory[uint32(v4+v11)] = byte(v13&i32(63) | i32(128))
									m.memory[uint32(v15)] = byte(int32(uint32(v13)>>6) | i32(192))
									v16 = v16 + i32(2)
									goto l543
								}
								m.memory[uint32(v15+i32(2))] = byte(v13&i32(63) | i32(128))
								m.memory[uint32(v4+v11)] = byte(int32(uint32(v13)>>6)&i32(63) | i32(128))
								m.memory[uint32(v15)] = byte(int32(uint32(v13&i32(61440))>>12) | i32(224))
								v16 = v16 + i32(3)
								goto l543
							}
						}
						if v6 != 0 {
							m.memory[int64(uint32(v0))+6] = byte(i32(0))
							store16(m.memory[int64(uint32(v0))+4:], uint16(i32(258)))
							goto l522
						}
						m.memory[int64(uint32(v1))+2] = byte(v5)
						m.memory[int64(uint32(v1))+1] = byte(i32(1))
						m.memory[int64(uint32(v0))+4] = byte(i32(0))
						goto l522
					}
					switch v10&i32(255) + i32(-142) {
					case 1:
						if uint32(v15) < uint32(v3) {
							v13 = v15 + i32(1)
							{
								t414 := int32(int8(m.memory[uint32(v2+v15)]))
								v5 = t414
								v12 = v5 + i32(95)
								v10 = v12 & i32(255)
								if uint32(v10) > uint32(i32(93)) {
									m.memory[int64(uint32(v0))+4] = byte(i32(2))
									if v5 > i32(-1) {
										store32(m.memory[int64(uint32(v0))+8:], uint32(v16))
										store32(m.memory[uint32(v0):], uint32(v15))
										store16(m.memory[int64(uint32(v0))+5:], uint16(i32(1)))
										goto l25
									}
									store32(m.memory[int64(uint32(v0))+8:], uint32(v16))
									store32(m.memory[uint32(v0):], uint32(v13))
									store16(m.memory[int64(uint32(v0))+5:], uint16(i32(2)))
									goto l25
								}
								if uint32(v13) < uint32(v3) {
									v5 = v15 + i32(2)
									{
										t415 := int32(int8(m.memory[uint32(v2+v13)]))
										v12 = t415
										v15 = (v12 + i32(95)) & i32(255)
										if uint32(v15) > uint32(i32(93)) {
											m.memory[int64(uint32(v0))+4] = byte(i32(2))
											if v12 > i32(-1) {
												store32(m.memory[int64(uint32(v0))+8:], uint32(v16))
												store32(m.memory[uint32(v0):], uint32(v13))
												store16(m.memory[int64(uint32(v0))+5:], uint16(i32(2)))
												goto l25
											}
											store32(m.memory[int64(uint32(v0))+8:], uint32(v16))
											store32(m.memory[uint32(v0):], uint32(v5))
											store16(m.memory[int64(uint32(v0))+5:], uint16(i32(3)))
											goto l25
										}
										v13 = v10*i32(94) + v15
										v10 = v13 + i32(-1410)
										if uint32(v10) < uint32(i32(5801)) {
											t427 := int32(load16(m.memory[int64(uint32(v10<<1))+1197886:]))
											t428 := v4 + v16
											v13 = t427
											m.memory[uint32(t428)] = byte(int32(uint32(v13)>>12) | i32(224))
											t429 := int32(load32(m.memory[int64(uint32(v7))+52:]))
											v4 = t429
											t430 := int32(load32(m.memory[int64(uint32(v7))+60:]))
											t431 := v4
											v10 = t430
											m.memory[uint32(t431+v10+i32(1))] = byte(int32(uint32(v13)>>6)&i32(63) | i32(128))
											t432 := v7
											v15 = v10 + i32(2)
											store32(m.memory[int64(uint32(t432))+60:], uint32(v15))
											m.memory[uint32(v4+v15)] = byte(v13&i32(63) | i32(128))
											v16 = v10 + i32(3)
											goto l543
										}
										{
											v10 = v13 + i32(-108)
											if uint32(v10) >= uint32(i32(11)) {
												goto l529
											}
											v15 = i32(2)
											goto l530
										l529:
											v10 = v13 + i32(-127)
											if uint32(v10) >= uint32(i32(3)) {
												goto l531
											}
											v15 = i32(5)
											goto l530
										l531:
											v10 = v13 + i32(-168)
											if uint32(v10) >= uint32(i32(7)) {
												goto l532
											}
											v15 = i32(8)
											goto l530
										l532:
											v10 = v13 + i32(-534)
											if uint32(v10) >= uint32(i32(12)) {
												goto l533
											}
											v15 = i32(11)
											goto l530
										l533:
											v10 = v13 + i32(-550)
											if uint32(v10) >= uint32(i32(12)) {
												goto l534
											}
											v15 = i32(14)
											goto l530
										l534:
											v10 = v13 + i32(-608)
											if uint32(v10) >= uint32(i32(2)) {
												goto l535
											}
											v15 = i32(17)
											goto l530
										l535:
											v10 = v13 + i32(-656)
											if uint32(v10) >= uint32(i32(2)) {
												goto l536
											}
											v15 = i32(20)
											goto l530
										l536:
											v10 = v13 + i32(-752)
											if uint32(v10) >= uint32(i32(16)) {
												goto l537
											}
											v15 = i32(23)
											goto l530
										l537:
											v10 = v13 + i32(-784)
											if uint32(v10) >= uint32(i32(16)) {
												goto l538
											}
											v15 = i32(26)
											goto l530
										l538:
											v10 = v13 + i32(-846)
											if uint32(v10) >= uint32(i32(87)) {
												goto l539
											}
											v15 = i32(29)
											goto l530
										l539:
											v10 = v13 + i32(-940)
											if uint32(v10) > uint32(i32(86)) {
												goto l540
											}
											v15 = i32(32)
										l530:
											t416 := int32(load16(m.memory[int64(uint32(v15<<1))+1244370:]))
											v10 = v10 + t416
											if uint32(v10) >= uint32(i32(255)) {
												m.fn33(v10, i32(255), i32(1242196))
												panic("unreachable")
											}
											t417 := int32(load16(m.memory[int64(uint32(v10<<1))+1227352:]))
											v10 = t417
											if v10 == 0 {
												goto l540
											}
											v15 = v16 + i32(1)
											v13 = v4 + v16
											if uint32(v10) < uint32(i32(2048)) {
												m.memory[uint32(v4+v15)] = byte(v10&i32(63) | i32(128))
												m.memory[uint32(v13)] = byte(int32(uint32(v10)>>6) | i32(192))
												v16 = v16 + i32(2)
												goto l543
											}
											m.memory[uint32(v13)] = byte(int32(uint32(v10)>>12) | i32(224))
											m.memory[uint32(v13+i32(2))] = byte(v10&i32(63) | i32(128))
											m.memory[uint32(v4+v15)] = byte(int32(uint32(v10)>>6)&i32(63) | i32(128))
											v16 = v16 + i32(3)
											goto l543
										}
									l540:
										v10 = v13 + i32(-597)
										if uint32(v10) < uint32(i32(11)) {
											m.memory[uint32(v4+v16)] = byte(i32(208))
											t441 := int32(load32(m.memory[int64(uint32(v7))+60:]))
											t442 := v7
											v13 = t441
											v15 = v13 + i32(1)
											store32(m.memory[int64(uint32(t442))+60:], uint32(v15))
											t443 := int32(load32(m.memory[int64(uint32(v7))+52:]))
											v4 = t443
											m.memory[uint32(v4+v15)] = byte(v10 + i32(-126))
											v16 = v13 + i32(2)
											goto l543
										}
										v13 = v13 + i32(-645)
										if uint32(v13) < uint32(i32(11)) {
											m.memory[uint32(v4+v16)] = byte(i32(209))
											t418 := int32(load32(m.memory[int64(uint32(v7))+60:]))
											t419 := v7
											v10 = t418
											v15 = v10 + i32(1)
											store32(m.memory[int64(uint32(t419))+60:], uint32(v15))
											t420 := int32(load32(m.memory[int64(uint32(v7))+52:]))
											v4 = t420
											m.memory[uint32(v4+v15)] = byte(v13 + i32(-110))
											v16 = v10 + i32(2)
											goto l543
										}
										m.memory[int64(uint32(v0))+6] = byte(i32(0))
										store16(m.memory[int64(uint32(v0))+4:], uint16(i32(770)))
										goto l546
									}
								}
								if v6 != 0 {
									goto l525
								}
								m.memory[int64(uint32(v1))+2] = byte(v12)
								m.memory[int64(uint32(v1))+1] = byte(i32(3))
								m.memory[int64(uint32(v0))+4] = byte(i32(0))
								goto l526
							l525:
								m.memory[int64(uint32(v0))+6] = byte(i32(0))
								store16(m.memory[int64(uint32(v0))+4:], uint16(i32(514)))
							l526:
								store32(m.memory[int64(uint32(v0))+8:], uint32(v16))
								store32(m.memory[uint32(v0):], uint32(v13))
								goto l25
							}
						}
						if v6 != 0 {
							m.memory[int64(uint32(v0))+6] = byte(i32(0))
							store16(m.memory[int64(uint32(v0))+4:], uint16(i32(258)))
							goto l522
						}
						m.memory[int64(uint32(v0))+4] = byte(i32(0))
						m.memory[int64(uint32(v1))+1] = byte(i32(2))
						goto l522
					default:
						m.memory[int64(uint32(v0))+6] = byte(i32(0))
						store16(m.memory[int64(uint32(v0))+4:], uint16(i32(258)))
						goto l522
					case 0:
						if uint32(v15) < uint32(v3) {
							v5 = v15 + i32(1)
							t421 := int32(int8(m.memory[uint32(v2+v15)]))
							v13 = t421
							v10 = (v13 + i32(95)) & i32(255)
							if uint32(v10) > uint32(i32(62)) {
								m.memory[int64(uint32(v0))+4] = byte(i32(2))
								if v13 > i32(-1) {
									store32(m.memory[int64(uint32(v0))+8:], uint32(v16))
									store32(m.memory[uint32(v0):], uint32(v15))
									store16(m.memory[int64(uint32(v0))+5:], uint16(i32(1)))
									goto l25
								}
								store32(m.memory[int64(uint32(v0))+8:], uint32(v16))
								store32(m.memory[uint32(v0):], uint32(v5))
								store16(m.memory[int64(uint32(v0))+5:], uint16(i32(2)))
								goto l25
							}
							m.memory[uint32(v4+v16)] = byte(i32(239))
							t422 := int32(load32(m.memory[int64(uint32(v7))+52:]))
							v4 = t422
							t423 := int32(load32(m.memory[int64(uint32(v7))+60:]))
							t424 := v4
							v13 = t423
							t425 := t424 + v13 + i32(1)
							v10 = v10 + i32(-159)
							m.memory[uint32(t425)] = byte(int32(uint32(v10)>>6) & i32(191))
							t426 := v7
							v15 = v13 + i32(2)
							store32(m.memory[int64(uint32(t426))+60:], uint32(v15))
							m.memory[uint32(v4+v15)] = byte(v10&i32(63) | i32(128))
							v16 = v13 + i32(3)
							goto l543
						}
						if v6 != 0 {
							m.memory[int64(uint32(v0))+6] = byte(i32(0))
							store16(m.memory[int64(uint32(v0))+4:], uint16(i32(258)))
							goto l522
						}
						m.memory[int64(uint32(v0))+4] = byte(i32(0))
						m.memory[int64(uint32(v1))+1] = byte(i32(4))
						goto l522
					}
				l564:
					{
						t466 := int32(load16(m.memory[int64(uint32(v15<<1))+1241972:]))
						v13 = v10 + t466
						if uint32(v13) < uint32(i32(240)) {
							goto l591
						}
						m.fn33(v13, i32(240), i32(1241544))
						panic("unreachable")
					}
				l591:
					v15 = v16 + i32(1)
					v10 = v4 + v16
					{
						t467 := int32(load16(m.memory[int64(uint32(v13<<1))+1226508:]))
						v13 = t467
						if uint32(v13) < uint32(i32(2048)) {
							goto l592
						}
						m.memory[uint32(v10)] = byte(int32(uint32(v13)>>12) | i32(224))
						m.memory[uint32(v10+i32(2))] = byte(v13&i32(63) | i32(128))
						m.memory[uint32(v4+v15)] = byte(int32(uint32(v13)>>6)&i32(63) | i32(128))
						v16 = v16 + i32(3)
						goto l543
					}
				l592:
					m.memory[uint32(v4+v15)] = byte(v13&i32(63) | i32(128))
					m.memory[uint32(v10)] = byte(int32(uint32(v13)>>6) | i32(192))
					v16 = v16 + i32(2)
				l543:
					store32(m.memory[int64(uint32(v7))+60:], uint32(v16))
					if uint32(v5) < uint32(v3) {
						goto l593
					}
					store32(m.memory[int64(uint32(v0))+8:], uint32(v16))
					store32(m.memory[uint32(v0):], uint32(v5))
					m.memory[int64(uint32(v0))+4] = byte(i32(0))
					goto l25
				l593:
					{
						t468 := int32(load32(m.memory[int64(uint32(v7))+56:]))
						if uint32(v16+i32(2)) < uint32(t468) {
							goto l594
						}
						store32(m.memory[int64(uint32(v0))+8:], uint32(v16))
						store32(m.memory[uint32(v0):], uint32(v5))
						m.memory[int64(uint32(v0))+4] = byte(i32(1))
						goto l25
					}
				l594:
					v15 = v5 + i32(1)
					t469 := int32(int8(m.memory[uint32(v2+v5)]))
					v10 = t469
					if v10 < i32(0) {
						goto l595
					}
				}
				m.memory[uint32(v4+v16)] = byte(v10)
				t470 := int32(load32(m.memory[int64(uint32(v7))+60:]))
				t471 := v7
				v17 = t470 + i32(1)
				store32(m.memory[int64(uint32(t471))+60:], uint32(v17))
				t472 := int32(load32(m.memory[int64(uint32(v7))+56:]))
				t473 := v17
				v5 = t472
				if uint32(t473) <= uint32(v5) {
					goto l596
				}
			}
			m.fn121(v17, v5, v5, i32(1146188))
			panic("unreachable")
		l546:
			store32(m.memory[int64(uint32(v0))+8:], uint32(v16))
			store32(m.memory[uint32(v0):], uint32(v5))
			goto l25
		l522:
			store32(m.memory[int64(uint32(v0))+8:], uint32(v16))
			store32(m.memory[uint32(v0):], uint32(v15))
			goto l25
		case 3:
			v17 = i32(0)
			v15 = i32(0)
			{
				t474 := int32(m.memory[int64(uint32(v1))+1])
				if t474 != i32(1) {
					goto l653
				}
				m.memory[int64(uint32(v1))+1] = byte(i32(0))
				if v3 != 0 {
					goto l598
				}
				if v6 != 0 {
					m.memory[int64(uint32(v0))+6] = byte(i32(0))
					store16(m.memory[int64(uint32(v0))+4:], uint16(i32(258)))
					store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
					store32(m.memory[uint32(v0):], uint32(i32(0)))
					goto l25
				}
				store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
				store32(m.memory[uint32(v0):], uint32(i32(0)))
				m.memory[int64(uint32(v0))+4] = byte(i32(0))
				goto l25
			l598:
				if uint32(v5) > uint32(i32(3)) {
					goto l600
				}
				store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
				store32(m.memory[uint32(v0):], uint32(i32(0)))
				m.memory[int64(uint32(v0))+4] = byte(i32(1))
				goto l25
			l600:
				t475 := int32(m.memory[int64(uint32(v1))+2])
				v15 = t475
				{
					t476 := int32(int8(m.memory[uint32(v2)]))
					v13 = t476
					v10 = v13 + i32(-64)
					if uint32(v10&i32(255)) < uint32(i32(63)) {
						goto l601
					}
					if uint32((v13+i32(1))&i32(255)) < uint32(i32(162)) {
						m.memory[int64(uint32(v0))+4] = byte(i32(2))
						if v13 > i32(-1) {
							store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
							store32(m.memory[uint32(v0):], uint32(i32(0)))
							store16(m.memory[int64(uint32(v0))+5:], uint16(i32(1)))
							goto l25
						}
						store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
						store32(m.memory[uint32(v0):], uint32(i32(1)))
						store16(m.memory[int64(uint32(v0))+5:], uint16(i32(2)))
						goto l25
					}
					v10 = v13 + i32(-98)
				}
			l601:
				{
					v15 = v15*i32(157) + v10&i32(255)
					v10 = v15 + i32(-942)
					if uint32(v10) >= uint32(i32(18840)) {
						goto l603
					}
					t477 := int32(load16(m.memory[int64(uint32(v10<<1))+1160142:]))
					v16 = t477
					if v16 != 0 {
						t478 := int32(load32(m.memory[int64(uint32(int32(uint32(v10)>>3)&i32(0x1ffffffc)))+1224152:]))
						if i32_shr_u(t478, v10)&i32(1) != 0 {
							m.memory[uint32(v4)] = byte(i32(240))
							m.memory[int64(uint32(v4))+3] = byte(v16&i32(63) | i32(128))
							m.memory[int64(uint32(v4))+1] = byte(int32(uint32(v16)>>12) | i32(160))
							m.memory[int64(uint32(v4))+2] = byte(int32(uint32(v16)>>6)&i32(63) | i32(128))
							goto l612
						}
						if uint32(v16) < uint32(i32(2048)) {
							m.memory[int64(uint32(v4))+1] = byte(v16&i32(63) | i32(128))
							m.memory[uint32(v4)] = byte(int32(uint32(v16)>>6) | i32(192))
							v15 = i32(1)
							v17 = i32(2)
							goto l653
						}
						m.memory[int64(uint32(v4))+2] = byte(v16&i32(63) | i32(128))
						m.memory[uint32(v4)] = byte(int32(uint32(v16)>>12) | i32(224))
						m.memory[int64(uint32(v4))+1] = byte(int32(uint32(v16)>>6)&i32(63) | i32(128))
						v15 = i32(1)
						v17 = i32(3)
						goto l653
					}
				}
			l603:
				switch v15 + i32(-1133) {
				case 0:
					store32(m.memory[uint32(v4):], uint32(i32(-2066969917)))
					goto l612
				case 2:
					store32(m.memory[uint32(v4):], uint32(i32(-0x7333753d)))
					goto l612
				case 31:
					store32(m.memory[uint32(v4):], uint32(i32(-0x7b33553d)))
					goto l612
				case 33:
					goto l609
				default:
					m.memory[int64(uint32(v0))+4] = byte(i32(2))
					if v13 > i32(-1) {
						store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
						store32(m.memory[uint32(v0):], uint32(i32(0)))
						store16(m.memory[int64(uint32(v0))+5:], uint16(i32(1)))
						goto l25
					}
					store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
					store32(m.memory[uint32(v0):], uint32(i32(1)))
					store16(m.memory[int64(uint32(v0))+5:], uint16(i32(2)))
					goto l25
				}
			l609:
				store32(m.memory[uint32(v4):], uint32(i32(-0x7333553d)))
			l612:
				v15 = i32(1)
				v17 = i32(4)
			}
		l653:
			{
				if uint32(v5) < uint32(v17) {
					m.fn121(v17, v5, v5, i32(1146220))
					panic("unreachable")
				}
				v13 = v5 - v17
				t479 := v13
				v10 = v3 - v15
				t480 := v10
				var p481 int32
				if uint32(v13) < uint32(v10) {
					p481 = 1
				}
				v14 = p481
				p482 := t480
				if v14 != 0 {
					p482 = t479
				}
				v12 = p482
				v13 = i32(0)
				v11 = v4 + v17
				t483 := v11
				v16 = v2 + v15
				if (t483^v16)&i32(3) != 0 {
					goto l617
				}
				v13 = i32(0)
				v18 = (i32(0) - v16) & i32(3)
				if uint32(v18|i32(8)) > uint32(v12) {
					goto l617
				}
				if v18 != 0 {
					v13 = i32(0)
					t484 := int32(int8(m.memory[uint32(v16)]))
					v10 = t484
					if v10 < i32(0) {
						goto l620
					}
					m.memory[uint32(v11)] = byte(v10)
					v13 = i32(1)
					if v18 == i32(1) {
						goto l619
					}
					{
						t485 := int32(int8(m.memory[int64(uint32(v16))+1]))
						v10 = t485
						if v10 >= i32(0) {
							m.memory[int64(uint32(v11))+1] = byte(v10)
							v13 = i32(2)
							if v18 == i32(2) {
								goto l619
							}
							{
								t486 := int32(int8(m.memory[int64(uint32(v16))+2]))
								v10 = t486
								if v10 >= i32(0) {
									m.memory[int64(uint32(v11))+2] = byte(v10)
									v13 = i32(3)
									goto l619
								}
								v13 = i32(2)
								goto l620
							}
						}
						v13 = i32(1)
						goto l620
					}
				}
				v13 = i32(0)
				goto l619
			}
		l619:
			v8 = v12 + i32(-8)
		l626:
			{
				v10 = v11 + v13
				t487 := v10
				v18 = v16 + v13
				t488 := int32(load32(m.memory[uint32(v18):]))
				v9 = t488
				store32(m.memory[uint32(t487):], uint32(v9))
				t489 := int32(load32(m.memory[uint32(v18+i32(4)):]))
				t490 := v10 + i32(4)
				v10 = t489
				store32(m.memory[uint32(t490):], uint32(v10))
				{
					v18 = v10 & i32(-2139062144)
					t491 := v18
					v10 = v9 & i32(-2139062144)
					if t491|v10 == 0 {
						goto l623
					}
					if v10 != 0 {
						goto l624
					}
					v10 = int32(uint32(int32(bits.TrailingZeros32(uint32(v18))))>>3) + i32(4)
					goto l625
				l624:
					v10 = int32(uint32(int32(bits.TrailingZeros32(uint32(v10)))) >> 3)
				l625:
					t492 := v16
					v13 = v10 + v13
					t493 := int32(m.memory[uint32(t492+v13)])
					v10 = t493
					goto l620
				}
			l623:
				v13 = v13 + i32(8)
				if uint32(v13) <= uint32(v8) {
					goto l626
				}
			}
		l617:
			if uint32(v13) >= uint32(v12) {
				goto l627
			}
		l628:
			{
				t494 := int32(int8(m.memory[uint32(v16+v13)]))
				v10 = t494
				if v10 < i32(0) {
					goto l620
				}
				m.memory[uint32(v11+v13)] = byte(v10)
				t495 := v12
				v13 = v13 + i32(1)
				if t495 != v13 {
					goto l628
				}
			}
		l627:
			v13 = v12 + v17
			v15 = v12 + v15
			goto l629
		l620:
			v15 = v13 + v15
			v13 = v13 + v17
			if uint32(v13+i32(3)) < uint32(v5) {
				v15 = v15 + i32(1)
			l652:
				{
					v10 = v10 + i32(127)
					v16 = v10 & i32(255)
					if uint32(v16) > uint32(i32(125)) {
						m.memory[int64(uint32(v0))+6] = byte(i32(0))
						store16(m.memory[int64(uint32(v0))+4:], uint16(i32(258)))
						goto l642
					}
					if uint32(v15) >= uint32(v3) {
						if v6 != 0 {
							m.memory[int64(uint32(v0))+6] = byte(i32(0))
							store16(m.memory[int64(uint32(v0))+4:], uint16(i32(258)))
							goto l642
						}
						m.memory[int64(uint32(v1))+2] = byte(v10)
						m.memory[int64(uint32(v1))+1] = byte(i32(1))
						m.memory[int64(uint32(v0))+4] = byte(i32(0))
						goto l642
					}
					v12 = v15 + i32(1)
					{
						v17 = v2 + v15
						t496 := int32(int8(m.memory[uint32(v17)]))
						v11 = t496
						v10 = v11 + i32(-64)
						if uint32(v10&i32(255)) < uint32(i32(63)) {
							goto l633
						}
						if uint32((v11+i32(1))&i32(255)) < uint32(i32(162)) {
							m.memory[int64(uint32(v0))+4] = byte(i32(2))
							if v11 > i32(-1) {
								store32(m.memory[int64(uint32(v0))+8:], uint32(v13))
								store32(m.memory[uint32(v0):], uint32(v15))
								store16(m.memory[int64(uint32(v0))+5:], uint16(i32(1)))
								goto l25
							}
							store32(m.memory[int64(uint32(v0))+8:], uint32(v13))
							store32(m.memory[uint32(v0):], uint32(v12))
							store16(m.memory[int64(uint32(v0))+5:], uint16(i32(2)))
							goto l25
						}
						v10 = v11 + i32(-98)
					}
				l633:
					{
						v16 = v16*i32(157) + v10&i32(255)
						v10 = v16 + i32(-942)
						if uint32(v10) >= uint32(i32(18840)) {
							goto l635
						}
						t497 := int32(load16(m.memory[int64(uint32(v10<<1))+1160142:]))
						v18 = t497
						if v18 != 0 {
							v16 = v13 + i32(1)
							{
								t498 := int32(load32(m.memory[int64(uint32(int32(uint32(v10)>>3)&i32(0x1ffffffc)))+1224152:]))
								if i32_shr_u(t498, v10)&i32(1) != 0 {
									v10 = v4 + v13
									m.memory[uint32(v10)] = byte(i32(240))
									m.memory[uint32(v4+v16)] = byte(int32(uint32(v18)>>12) | i32(160))
									m.memory[uint32(v10+i32(3))] = byte(v18&i32(63) | i32(128))
									m.memory[uint32(v10+i32(2))] = byte(int32(uint32(v18)>>6)&i32(63) | i32(128))
									goto l647
								}
								v10 = v4 + v13
								if uint32(v18) < uint32(i32(2048)) {
									m.memory[uint32(v4+v16)] = byte(v18&i32(63) | i32(128))
									m.memory[uint32(v10)] = byte(int32(uint32(v18)>>6) | i32(192))
									v10 = i32(2)
									goto l646
								}
								m.memory[uint32(v10)] = byte(int32(uint32(v18)>>12) | i32(224))
								m.memory[uint32(v10+i32(2))] = byte(v18&i32(63) | i32(128))
								m.memory[uint32(v4+v16)] = byte(int32(uint32(v18)>>6)&i32(63) | i32(128))
								v10 = i32(3)
								goto l646
							}
						}
					}
				l635:
					switch v16 + i32(-1133) {
					case 0:
						store32(m.memory[uint32(v4+v13):], uint32(i32(-2066969917)))
						goto l647
					case 2:
						store32(m.memory[uint32(v4+v13):], uint32(i32(-0x7333753d)))
						goto l647
					case 31:
						store32(m.memory[uint32(v4+v13):], uint32(i32(-0x7b33553d)))
						goto l647
					case 33:
						store32(m.memory[uint32(v4+v13):], uint32(i32(-0x7333553d)))
						goto l647
					default:
						m.memory[int64(uint32(v0))+4] = byte(i32(2))
						if v11 > i32(-1) {
							store32(m.memory[int64(uint32(v0))+8:], uint32(v13))
							store32(m.memory[uint32(v0):], uint32(v15))
							store16(m.memory[int64(uint32(v0))+5:], uint16(i32(1)))
							goto l25
						}
						store32(m.memory[int64(uint32(v0))+8:], uint32(v13))
						store32(m.memory[uint32(v0):], uint32(v12))
						store16(m.memory[int64(uint32(v0))+5:], uint16(i32(2)))
						goto l25
					}
				l647:
					v10 = i32(4)
				l646:
					v13 = v10 + v13
					if uint32(v12) < uint32(v3) {
						goto l650
					}
					store32(m.memory[int64(uint32(v0))+8:], uint32(v13))
					store32(m.memory[uint32(v0):], uint32(v12))
					m.memory[int64(uint32(v0))+4] = byte(i32(0))
					goto l25
				l650:
					if uint32(v13+i32(3)) < uint32(v5) {
						goto l651
					}
					store32(m.memory[int64(uint32(v0))+8:], uint32(v13))
					store32(m.memory[uint32(v0):], uint32(v12))
					m.memory[int64(uint32(v0))+4] = byte(i32(1))
					goto l25
				l651:
					v15 = v15 + i32(2)
					t499 := int32(int8(m.memory[uint32(v17+i32(1))]))
					v10 = t499
					if v10 < i32(0) {
						goto l652
					}
				}
				m.memory[uint32(v4+v13)] = byte(v10)
				v17 = v13 + i32(1)
				goto l653
			l642:
				store32(m.memory[int64(uint32(v0))+8:], uint32(v13))
				store32(m.memory[uint32(v0):], uint32(v15))
				goto l25
			}
			v14 = i32(1)
		l629:
			store32(m.memory[int64(uint32(v0))+8:], uint32(v13))
			store32(m.memory[uint32(v0):], uint32(v15))
			store32(m.memory[int64(uint32(v0))+4:], uint32(v14|(v7+i32(52))&i32(-256)))
			goto l25
		case 2:
			store32(m.memory[int64(uint32(v7))+56:], uint32(v5))
			store32(m.memory[int64(uint32(v7))+52:], uint32(v4))
			v15 = i32(0)
			v13 = i32(0)
			{
				t500 := int32(m.memory[int64(uint32(v1))+7])
				if t500 != i32(1) {
					goto l654
				}
				if uint32(v5) > uint32(i32(2)) {
					goto l655
				}
				store32(m.memory[uint32(v0):], uint32(i32(0)))
				m.memory[int64(uint32(v0))+4] = byte(i32(1))
				store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
				goto l25
			l655:
				m.memory[int64(uint32(v1))+7] = byte(i32(0))
				t501 := int32(m.memory[int64(uint32(v1))+8])
				m.memory[uint32(v4)] = byte(t501)
				v13 = i32(1)
			}
		l654:
			{
				{
					t502 := int32(m.memory[int64(uint32(v1))+9])
					v10 = t502
					if v10 != 0 {
						goto l656
					}
					v17 = v13
					goto l788
				}
			l656:
				{
					if v3 == 0 {
						goto l658
					}
					v17 = v13 + i32(3)
					if uint32(v17) >= uint32(v5) {
						store32(m.memory[uint32(v0):], uint32(i32(0)))
						m.memory[int64(uint32(v0))+4] = byte(i32(1))
						store32(m.memory[int64(uint32(v0))+8:], uint32(v13))
						goto l25
					}
					t503 := int32(m.memory[int64(uint32(v1))+10])
					v9 = t503
					t504 := int32(m.memory[int64(uint32(v1))+12])
					v18 = t504
					t505 := int32(m.memory[int64(uint32(v1))+11])
					v11 = t505
					v16 = i32(0)
				l666:
					{
						v15 = v16 + i32(1)
						t506 := int32(int8(m.memory[uint32(v2+v16)]))
						v12 = t506
						switch v10&i32(255) + i32(-1) {
						case 2:
							m.memory[int64(uint32(v1))+9] = byte(i32(0))
							v10 = (v12 + i32(-48)) & i32(255)
							if uint32(v10) > uint32(i32(9)) {
								m.memory[int64(uint32(v1))+10] = byte(v18)
								m.memory[int64(uint32(v1))+9] = byte(i32(1))
								m.memory[int64(uint32(v1))+7] = byte(i32(1))
								m.memory[int64(uint32(v0))+6] = byte(i32(2))
								store16(m.memory[int64(uint32(v0))+4:], uint16(i32(258)))
								store32(m.memory[uint32(v0):], uint32(v16))
								m.memory[int64(uint32(v1))+8] = byte(v11 + i32(48))
								store32(m.memory[int64(uint32(v0))+8:], uint32(v13))
								goto l25
							}
							v10 = v11&i32(255)*i32(1260) + v18&i32(255)*i32(10) + v10 + v9*i32(12600)
							if uint32(v10) < uint32(i32(39420)) {
								if v10 == i32(7457) {
									v13 = v4 + v13
									m.memory[int64(uint32(v13))+2] = byte(i32(135))
									store16(m.memory[uint32(v13):], uint16(i32(40942)))
									goto l788
								}
								{
									{
										p618 := i32(103)
										if uint32(v10) < uint32(i32(11334)) {
											p618 = i32(0)
										}
										v16 = p618
										t619 := v16
										v16 = v16 + i32(51)
										t620 := int32(load16(m.memory[int64(uint32(v16<<1))+1241560:]))
										t621 := v16
										v16 = v10 & i32(0xffff)
										p622 := t621
										if uint32(t620) > uint32(v16) {
											p622 = t619
										}
										v12 = p622
										t623 := v12
										v12 = v12 + i32(26)
										t624 := int32(load16(m.memory[int64(uint32(v12<<1))+1241560:]))
										p625 := v12
										if uint32(t624) > uint32(v16) {
											p625 = t623
										}
										v12 = p625
										t626 := v12
										v12 = v12 + i32(13)
										t627 := int32(load16(m.memory[int64(uint32(v12<<1))+1241560:]))
										p628 := v12
										if uint32(t627) > uint32(v16) {
											p628 = t626
										}
										v12 = p628
										t629 := v12
										v12 = v12 + i32(6)
										t630 := int32(load16(m.memory[int64(uint32(v12<<1))+1241560:]))
										p631 := v12
										if uint32(t630) > uint32(v16) {
											p631 = t629
										}
										v12 = p631
										t632 := v12
										v12 = v12 + i32(3)
										t633 := int32(load16(m.memory[int64(uint32(v12<<1))+1241560:]))
										p634 := v12
										if uint32(t633) > uint32(v16) {
											p634 = t632
										}
										v12 = p634
										t635 := v12
										v12 = v12 + i32(2)
										t636 := int32(load16(m.memory[int64(uint32(v12<<1))+1241560:]))
										p637 := v12
										if uint32(t636) > uint32(v16) {
											p637 = t635
										}
										v12 = p637
										t638 := v12
										v12 = v12 + i32(1)
										t639 := int32(load16(m.memory[int64(uint32(v12<<1))+1241560:]))
										p640 := v12
										if uint32(t639) > uint32(v16) {
											p640 = t638
										}
										v12 = p640
										v11 = v12 << 1
										t641 := int32(load16(m.memory[int64(uint32(v11))+1241560:]))
										v18 = t641
										if v18 == v16 {
											goto l701
										}
										{
											t642 := v12
											var p643 int32
											if uint32(v18) >= uint32(v16) {
												p643 = 1
											}
											v16 = t642 - p643
											if uint32(v16) >= uint32(i32(206)) {
												m.fn33(i32(-1), i32(206), i32(1227336))
												panic("unreachable")
											}
											v16 = v16 << 1
											t644 := int32(load16(m.memory[int64(uint32(v16))+1240928:]))
											t645 := int32(load16(m.memory[int64(uint32(v16))+1241560:]))
											v10 = t644 + v10 - t645
											goto l703
										}
									}
								l701:
									t646 := int32(load16(m.memory[int64(uint32(v11))+1240928:]))
									v10 = t646
								}
							l703:
								v16 = v13 | i32(2)
								v13 = v4 + v13
								if uint32(v10&i32(0xffff)) < uint32(i32(2048)) {
									m.memory[int64(uint32(v13))+1] = byte(v10&i32(63) | i32(128))
									m.memory[uint32(v13)] = byte(int32(uint32(v10)>>6) | i32(192))
									v17 = v16
									goto l788
								}
								m.memory[uint32(v4+v16)] = byte(v10&i32(63) | i32(128))
								m.memory[int64(uint32(v13))+1] = byte(int32(uint32(v10)>>6)&i32(63) | i32(128))
								m.memory[uint32(v13)] = byte(int32(uint32(v10&i32(61440))>>12) | i32(224))
								goto l788
							}
							if uint32(v10+i32(-189000)) < uint32(i32(0x100000)) {
								v16 = v4 + v13
								t616 := v16
								v10 = v10 + i32(-123464)
								m.memory[uint32(t616)] = byte(int32(uint32(v10)>>18) | i32(240))
								m.memory[uint32(v4+v17)] = byte(v10&i32(63) | i32(128))
								m.memory[int64(uint32(v16))+2] = byte(int32(uint32(v10)>>6)&i32(63) | i32(128))
								m.memory[int64(uint32(v16))+1] = byte(int32(uint32(v10)>>12)&i32(63) | i32(128))
								t617 := v7
								v17 = v13 | i32(4)
								store32(m.memory[int64(uint32(t617))+60:], uint32(v17))
								goto l788
							}
							m.memory[int64(uint32(v0))+6] = byte(i32(0))
							store16(m.memory[int64(uint32(v0))+4:], uint16(i32(1026)))
							store32(m.memory[uint32(v0):], uint32(v15))
							store32(m.memory[int64(uint32(v0))+8:], uint32(v13))
							goto l25
						case 1:
							v18 = v12 + i32(127)
							if uint32(v18&i32(255)) > uint32(i32(125)) {
								m.memory[int64(uint32(v1))+7] = byte(i32(1))
								m.memory[int64(uint32(v1))+9] = byte(i32(0))
								m.memory[int64(uint32(v0))+6] = byte(i32(1))
								store16(m.memory[int64(uint32(v0))+4:], uint16(i32(258)))
								store32(m.memory[uint32(v0):], uint32(v16))
								m.memory[int64(uint32(v1))+8] = byte(v11 + i32(48))
								store32(m.memory[int64(uint32(v0))+8:], uint32(v13))
								goto l25
							}
							m.memory[int64(uint32(v1))+12] = byte(v18)
							v10 = i32(3)
							goto l664
						default:
							v11 = v12 + i32(-48)
							if uint32(v11&i32(255)) > uint32(i32(9)) {
								m.memory[int64(uint32(v1))+9] = byte(i32(0))
								if uint32(v9) > uint32(i32(31)) {
									v10 = v12 + i32(95)
									if uint32(v10&i32(255)) < uint32(i32(94)) {
										v16 = v9 + i32(-47)
										if uint32(v16&i32(255)) < uint32(i32(72)) {
											v13 = v4 + v13
											t614 := int32(load16(m.memory[int64(uint32(v16&i32(255)*i32(188)+v10&i32(255)<<1))+1146606:]))
											t615 := v13
											v10 = t614
											m.memory[int64(uint32(t615))+2] = byte(v10&i32(63) | i32(128))
											m.memory[uint32(v13)] = byte(int32(uint32(v10)>>12) | i32(224))
											m.memory[int64(uint32(v13))+1] = byte(int32(uint32(v10)>>6)&i32(63) | i32(128))
											store32(m.memory[int64(uint32(v7))+60:], uint32(v17))
											goto l788
										}
										switch v9 + i32(-32) {
										case 0:
											v12 = v13 + i32(1)
											v16 = v4 + v13
											{
												t609 := int32(load16(m.memory[int64(uint32(v10&i32(255)<<1))+1219264:]))
												v10 = t609
												if uint32(v10) < uint32(i32(2048)) {
													m.memory[uint32(v4+v12)] = byte(v10&i32(63) | i32(128))
													m.memory[uint32(v16)] = byte(int32(uint32(v10)>>6) | i32(192))
													t610 := v7
													v17 = v13 | i32(2)
													store32(m.memory[int64(uint32(t610))+60:], uint32(v17))
													goto l788
												}
												m.memory[uint32(v16)] = byte(int32(uint32(v10)>>12) | i32(224))
												m.memory[int64(uint32(v16))+2] = byte(v10&i32(63) | i32(128))
												m.memory[uint32(v4+v12)] = byte(int32(uint32(v10)>>6)&i32(63) | i32(128))
												store32(m.memory[int64(uint32(v7))+60:], uint32(v17))
												goto l788
											}
										case 5:
											v16 = (v12 + i32(32)) & i32(255)
											if uint32(v16) > uint32(i32(21)) {
												goto l695
											}
											v13 = v4 + v13
											t611 := int32(load16(m.memory[int64(uint32(v16<<1))+1252566:]))
											t612 := v13
											v10 = t611
											m.memory[int64(uint32(t612))+2] = byte(v10&i32(63) | i32(128))
											m.memory[uint32(v13)] = byte(int32(uint32(v10)>>12) | i32(224))
											m.memory[int64(uint32(v13))+1] = byte(int32(uint32(v10)>>6)&i32(63) | i32(128))
											store32(m.memory[int64(uint32(v7))+60:], uint32(v17))
											goto l788
										case 7:
											v16 = v10 & i32(255)
											if uint32(v16) < uint32(i32(32)) {
												goto l696
											}
											goto l695
										default:
											if uint32(v9) <= uint32(i32(118)) {
												goto l695
											}
											v13 = v4 + v13
											t613 := v13
											v10 = (v9+i32(-119))&i32(255)*i32(94) + v10&i32(255) + i32(-7628)
											m.memory[int64(uint32(t613))+2] = byte(v10&i32(63) | i32(128))
											m.memory[int64(uint32(v13))+1] = byte(int32(uint32(v10)>>6)&i32(63) | i32(128))
											m.memory[uint32(v13)] = byte(int32(uint32(v10&i32(61440))>>12) | i32(224))
											store32(m.memory[int64(uint32(v7))+60:], uint32(v17))
											goto l788
										}
									}
									v10 = v12 + i32(-64)
									if uint32(v10&i32(255)) <= uint32(i32(62)) {
										goto l676
									}
									if v12 > i32(-96) {
										m.memory[int64(uint32(v0))+4] = byte(i32(2))
										if v12 > i32(-1) {
											store32(m.memory[uint32(v0):], uint32(v16))
											store16(m.memory[int64(uint32(v0))+5:], uint16(i32(1)))
											store32(m.memory[int64(uint32(v0))+8:], uint32(v13))
											goto l25
										}
										store32(m.memory[uint32(v0):], uint32(v15))
										store16(m.memory[int64(uint32(v0))+5:], uint16(i32(2)))
										store32(m.memory[int64(uint32(v0))+8:], uint32(v13))
										goto l25
									}
									v10 = v12 + i32(-65)
								l676:
									{
										v10 = (v9+i32(-32))&i32(255)*i32(96) + v10&i32(255)
										v16 = v10 + i32(-864)
										if uint32(v16) < uint32(i32(8059)) {
											{
												{
													p547 := i32(813)
													if uint32(v10) < uint32(i32(3734)) {
														p547 = i32(0)
													}
													v10 = p547
													t548 := v10
													v10 = v10 + i32(407)
													t549 := int32(load16(m.memory[int64(uint32(v10<<1))+1259814:]))
													t550 := v10
													v10 = v16 & i32(0xffff)
													p551 := t550
													if uint32(t549) > uint32(v10) {
														p551 = t548
													}
													v12 = p551
													t552 := v12
													v12 = v12 + i32(203)
													t553 := int32(load16(m.memory[int64(uint32(v12<<1))+1259814:]))
													p554 := v12
													if uint32(t553) > uint32(v10) {
														p554 = t552
													}
													v12 = p554
													t555 := v12
													v12 = v12 + i32(102)
													t556 := int32(load16(m.memory[int64(uint32(v12<<1))+1259814:]))
													p557 := v12
													if uint32(t556) > uint32(v10) {
														p557 = t555
													}
													v12 = p557
													t558 := v12
													v12 = v12 + i32(51)
													t559 := int32(load16(m.memory[int64(uint32(v12<<1))+1259814:]))
													p560 := v12
													if uint32(t559) > uint32(v10) {
														p560 = t558
													}
													v12 = p560
													t561 := v12
													v12 = v12 + i32(25)
													t562 := int32(load16(m.memory[int64(uint32(v12<<1))+1259814:]))
													p563 := v12
													if uint32(t562) > uint32(v10) {
														p563 = t561
													}
													v12 = p563
													t564 := v12
													v12 = v12 + i32(13)
													t565 := int32(load16(m.memory[int64(uint32(v12<<1))+1259814:]))
													p566 := v12
													if uint32(t565) > uint32(v10) {
														p566 = t564
													}
													v12 = p566
													t567 := v12
													v12 = v12 + i32(6)
													t568 := int32(load16(m.memory[int64(uint32(v12<<1))+1259814:]))
													p569 := v12
													if uint32(t568) > uint32(v10) {
														p569 = t567
													}
													v12 = p569
													t570 := v12
													v12 = v12 + i32(3)
													t571 := int32(load16(m.memory[int64(uint32(v12<<1))+1259814:]))
													p572 := v12
													if uint32(t571) > uint32(v10) {
														p572 = t570
													}
													v12 = p572
													t573 := v12
													v12 = v12 + i32(2)
													t574 := int32(load16(m.memory[int64(uint32(v12<<1))+1259814:]))
													p575 := v12
													if uint32(t574) > uint32(v10) {
														p575 = t573
													}
													v12 = p575
													t576 := v12
													v12 = v12 + i32(1)
													t577 := int32(load16(m.memory[int64(uint32(v12<<1))+1259814:]))
													p578 := v12
													if uint32(t577) > uint32(v10) {
														p578 = t576
													}
													v12 = p578
													v11 = v12 << 1
													t579 := int32(load16(m.memory[int64(uint32(v11))+1259814:]))
													v18 = t579
													if v18 == v10 {
														goto l681
													}
													{
														t580 := v12
														var p581 int32
														if uint32(v18) >= uint32(v10) {
															p581 = 1
														}
														v10 = t580 - p581
														if uint32(v10) >= uint32(i32(1627)) {
															m.fn33(i32(-1), i32(1627), i32(1227336))
															panic("unreachable")
														}
														t582 := v16
														v10 = v10 << 1
														t583 := int32(load16(m.memory[int64(uint32(v10))+1259814:]))
														t584 := int32(load16(m.memory[int64(uint32(v10))+1252610:]))
														v10 = t582 - t583 + t584
														v12 = v10 & i32(0xffff)
														v16 = int32(uint32(v12) >> 12)
														v12 = int32(uint32(v12) >> 6)
														goto l683
													}
												}
											l681:
												t585 := int32(load16(m.memory[int64(uint32(v11))+1252610:]))
												v10 = t585
												v16 = int32(uint32(v10) >> 12)
												v12 = int32(uint32(v10) >> 6)
											}
										l683:
											v13 = v4 + v13
											m.memory[uint32(v13)] = byte(v16 | i32(224))
											m.memory[int64(uint32(v13))+2] = byte(v10&i32(63) | i32(128))
											m.memory[int64(uint32(v13))+1] = byte(v12&i32(63) | i32(128))
											store32(m.memory[int64(uint32(v7))+60:], uint32(v17))
											goto l788
										}
										if uint32(v10) < uint32(i32(864)) {
											{
												{
													p586 := i32(29)
													if uint32(v10) < uint32(i32(777)) {
														p586 = i32(0)
													}
													v16 = p586
													t587 := v16
													v16 = v16 + i32(15)
													t588 := int32(load16(m.memory[int64(uint32(v16<<1))+1234878:]))
													p589 := v16
													if uint32(t588) > uint32(v10) {
														p589 = t587
													}
													v16 = p589
													t590 := v16
													v16 = v16 + i32(7)
													t591 := int32(load16(m.memory[int64(uint32(v16<<1))+1234878:]))
													p592 := v16
													if uint32(t591) > uint32(v10) {
														p592 = t590
													}
													v16 = p592
													t593 := v16
													v16 = v16 + i32(4)
													t594 := int32(load16(m.memory[int64(uint32(v16<<1))+1234878:]))
													p595 := v16
													if uint32(t594) > uint32(v10) {
														p595 = t593
													}
													v16 = p595
													t596 := v16
													v16 = v16 + i32(2)
													t597 := int32(load16(m.memory[int64(uint32(v16<<1))+1234878:]))
													p598 := v16
													if uint32(t597) > uint32(v10) {
														p598 = t596
													}
													v16 = p598
													t599 := v16
													v16 = v16 + i32(1)
													t600 := int32(load16(m.memory[int64(uint32(v16<<1))+1234878:]))
													p601 := v16
													if uint32(t600) > uint32(v10) {
														p601 = t599
													}
													v16 = p601
													v12 = v16 << 1
													t602 := int32(load16(m.memory[int64(uint32(v12))+1234878:]))
													v11 = t602
													if v11 == v10 {
														goto l684
													}
													{
														t603 := v16
														var p604 int32
														if uint32(v11) >= uint32(v10) {
															p604 = 1
														}
														v16 = t603 - p604
														if uint32(v16) >= uint32(i32(59)) {
															m.fn33(i32(-1), i32(59), i32(1227336))
															panic("unreachable")
														}
														v16 = v16 << 1
														t605 := int32(load16(m.memory[int64(uint32(v16))+1255864:]))
														t606 := int32(load16(m.memory[int64(uint32(v16))+1234878:]))
														v10 = t605 + v10 - t606
														goto l686
													}
												}
											l684:
												t607 := int32(load16(m.memory[int64(uint32(v12))+1255864:]))
												v10 = t607
											}
										l686:
											v12 = v13 + i32(1)
											v16 = v4 + v13
											if uint32(v10&i32(0xffff)) < uint32(i32(2048)) {
												m.memory[uint32(v4+v12)] = byte(v10&i32(63) | i32(128))
												m.memory[uint32(v16)] = byte(int32(uint32(v10)>>6) | i32(192))
												t608 := v7
												v17 = v13 | i32(2)
												store32(m.memory[int64(uint32(t608))+60:], uint32(v17))
												goto l788
											}
											m.memory[int64(uint32(v16))+2] = byte(v10&i32(63) | i32(128))
											m.memory[uint32(v4+v12)] = byte(int32(uint32(v10)>>6)&i32(63) | i32(128))
											m.memory[uint32(v16)] = byte(int32(uint32(v10&i32(61440))>>12) | i32(224))
											store32(m.memory[int64(uint32(v7))+60:], uint32(v17))
											goto l788
										}
										v10 = v10 + i32(-8923)
										if uint32(v10) >= uint32(i32(101)) {
											m.fn33(v10, i32(101), i32(1146252))
											panic("unreachable")
										}
										v13 = v4 + v13
										t545 := int32(load16(m.memory[int64(uint32(v10<<1))+1146268:]))
										t546 := v13
										v10 = t545
										m.memory[int64(uint32(t546))+2] = byte(v10&i32(63) | i32(128))
										m.memory[uint32(v13)] = byte(int32(uint32(v10)>>12) | i32(224))
										m.memory[int64(uint32(v13))+1] = byte(int32(uint32(v10)>>6)&i32(63) | i32(128))
										store32(m.memory[int64(uint32(v7))+60:], uint32(v17))
										goto l788
									}
								}
								v10 = v12 + i32(-64)
								if uint32(v10&i32(255)) <= uint32(i32(62)) {
									goto l669
								}
								if v12 > i32(-2) {
									m.memory[int64(uint32(v0))+4] = byte(i32(2))
									if v12 > i32(-1) {
										store32(m.memory[uint32(v0):], uint32(v16))
										store16(m.memory[int64(uint32(v0))+5:], uint16(i32(1)))
										store32(m.memory[int64(uint32(v0))+8:], uint32(v13))
										goto l25
									}
									store32(m.memory[uint32(v0):], uint32(v15))
									store16(m.memory[int64(uint32(v0))+5:], uint16(i32(2)))
									store32(m.memory[int64(uint32(v0))+8:], uint32(v13))
									goto l25
								}
								v10 = v12 + i32(-65)
							l669:
								{
									{
										v18 = v9*i32(190) + v10&i32(255)
										v10 = v18 & i32(0xffff)
										p507 := i32(958)
										if uint32(v10) < uint32(i32(2880)) {
											p507 = i32(0)
										}
										v16 = p507
										t508 := v16
										v16 = v16 + i32(479)
										t509 := int32(load16(m.memory[int64(uint32(v16<<1))+1255982:]))
										p510 := v16
										if uint32(t509) > uint32(v10) {
											p510 = t508
										}
										v16 = p510
										t511 := v16
										v16 = v16 + i32(239)
										t512 := int32(load16(m.memory[int64(uint32(v16<<1))+1255982:]))
										p513 := v16
										if uint32(t512) > uint32(v10) {
											p513 = t511
										}
										v16 = p513
										t514 := v16
										v16 = v16 + i32(120)
										t515 := int32(load16(m.memory[int64(uint32(v16<<1))+1255982:]))
										p516 := v16
										if uint32(t515) > uint32(v10) {
											p516 = t514
										}
										v16 = p516
										t517 := v16
										v16 = v16 + i32(60)
										t518 := int32(load16(m.memory[int64(uint32(v16<<1))+1255982:]))
										p519 := v16
										if uint32(t518) > uint32(v10) {
											p519 = t517
										}
										v16 = p519
										t520 := v16
										v16 = v16 + i32(30)
										t521 := int32(load16(m.memory[int64(uint32(v16<<1))+1255982:]))
										p522 := v16
										if uint32(t521) > uint32(v10) {
											p522 = t520
										}
										v16 = p522
										t523 := v16
										v16 = v16 + i32(15)
										t524 := int32(load16(m.memory[int64(uint32(v16<<1))+1255982:]))
										p525 := v16
										if uint32(t524) > uint32(v10) {
											p525 = t523
										}
										v16 = p525
										t526 := v16
										v16 = v16 + i32(7)
										t527 := int32(load16(m.memory[int64(uint32(v16<<1))+1255982:]))
										p528 := v16
										if uint32(t527) > uint32(v10) {
											p528 = t526
										}
										v16 = p528
										t529 := v16
										v16 = v16 + i32(4)
										t530 := int32(load16(m.memory[int64(uint32(v16<<1))+1255982:]))
										p531 := v16
										if uint32(t530) > uint32(v10) {
											p531 = t529
										}
										v16 = p531
										t532 := v16
										v16 = v16 + i32(2)
										t533 := int32(load16(m.memory[int64(uint32(v16<<1))+1255982:]))
										p534 := v16
										if uint32(t533) > uint32(v10) {
											p534 = t532
										}
										v16 = p534
										t535 := v16
										v16 = v16 + i32(1)
										t536 := int32(load16(m.memory[int64(uint32(v16<<1))+1255982:]))
										p537 := v16
										if uint32(t536) > uint32(v10) {
											p537 = t535
										}
										v16 = p537
										v12 = v16 << 1
										t538 := int32(load16(m.memory[int64(uint32(v12))+1255982:]))
										v11 = t538
										if v11 == v10 {
											goto l671
										}
										{
											t539 := v16
											var p540 int32
											if uint32(v11) >= uint32(v10) {
												p540 = 1
											}
											v10 = t539 - p540
											if uint32(v10) >= uint32(i32(1916)) {
												m.fn33(i32(-1), i32(1916), i32(1227336))
												panic("unreachable")
											}
											t541 := v18
											v10 = v10 << 1
											t542 := int32(load16(m.memory[int64(uint32(v10))+1255982:]))
											t543 := int32(load16(m.memory[int64(uint32(v10))+1247664:]))
											v10 = t541 - t542 + t543
											v12 = v10 & i32(0xffff)
											v16 = int32(uint32(v12) >> 12)
											v12 = int32(uint32(v12) >> 6)
											goto l673
										}
									}
								l671:
									t544 := int32(load16(m.memory[int64(uint32(v12))+1247664:]))
									v10 = t544
									v16 = int32(uint32(v10) >> 12)
									v12 = int32(uint32(v10) >> 6)
								}
							l673:
								v13 = v4 + v13
								m.memory[uint32(v13)] = byte(v16 | i32(224))
								m.memory[int64(uint32(v13))+2] = byte(v10&i32(63) | i32(128))
								m.memory[int64(uint32(v13))+1] = byte(v12&i32(63) | i32(128))
								store32(m.memory[int64(uint32(v7))+60:], uint32(v17))
								goto l788
							}
							m.memory[int64(uint32(v1))+11] = byte(v11)
							v10 = i32(2)
						}
					l664:
						v16 = v15
						if v3 != v15 {
							goto l666
						}
					}
					m.memory[int64(uint32(v1))+9] = byte(v10)
				}
			l658:
				if v6 != 0 {
					m.memory[int64(uint32(v0))+6] = byte(i32(0))
					m.memory[int64(uint32(v0))+5] = byte(v10)
					m.memory[int64(uint32(v0))+4] = byte(i32(2))
					m.memory[int64(uint32(v1))+9] = byte(i32(0))
					store32(m.memory[uint32(v0):], uint32(v3))
					store32(m.memory[int64(uint32(v0))+8:], uint32(v13))
					goto l25
				}
				store32(m.memory[uint32(v0):], uint32(v3))
				m.memory[int64(uint32(v0))+4] = byte(i32(0))
				store32(m.memory[int64(uint32(v0))+8:], uint32(v13))
				goto l25
			l695:
				{
					{
						v10 = (v9+i32(-33))&i32(255)*i32(94) + v10&i32(255)
						p647 := i32(23)
						if uint32(v10) < uint32(i32(425)) {
							p647 = i32(0)
						}
						v16 = p647
						t648 := v16
						v16 = v16 + i32(11)
						t649 := int32(load16(m.memory[int64(uint32(v16<<1))+1241340:]))
						p650 := v16
						if uint32(t649) > uint32(v10) {
							p650 = t648
						}
						v16 = p650
						t651 := v16
						v16 = v16 + i32(6)
						t652 := int32(load16(m.memory[int64(uint32(v16<<1))+1241340:]))
						p653 := v16
						if uint32(t652) > uint32(v10) {
							p653 = t651
						}
						v16 = p653
						t654 := v16
						v16 = v16 + i32(3)
						t655 := int32(load16(m.memory[int64(uint32(v16<<1))+1241340:]))
						p656 := v16
						if uint32(t655) > uint32(v10) {
							p656 = t654
						}
						v16 = p656
						t657 := v16
						v16 = v16 + i32(1)
						t658 := int32(load16(m.memory[int64(uint32(v16<<1))+1241340:]))
						p659 := v16
						if uint32(t658) > uint32(v10) {
							p659 = t657
						}
						v16 = p659
						t660 := v16
						v16 = v16 + i32(1)
						t661 := int32(load16(m.memory[int64(uint32(v16<<1))+1241340:]))
						p662 := v16
						if uint32(t661) > uint32(v10) {
							p662 = t660
						}
						v16 = p662
						v12 = v16 << 1
						t663 := int32(load16(m.memory[int64(uint32(v12))+1241340:]))
						v11 = t663
						if v11 == v10 {
							goto l705
						}
						{
							t664 := v16
							var p665 int32
							if uint32(v11) >= uint32(v10) {
								p665 = 1
							}
							v16 = t664 - p665
							if uint32(v16) >= uint32(i32(46)) {
								m.fn33(i32(-1), i32(46), i32(1227336))
								panic("unreachable")
							}
							v16 = v16 << 1
							t666 := int32(load16(m.memory[int64(uint32(v16))+1263068:]))
							t667 := int32(load16(m.memory[int64(uint32(v16))+1241340:]))
							v10 = t666 + v10 - t667
							goto l707
						}
					}
				l705:
					t668 := int32(load16(m.memory[int64(uint32(v12))+1263068:]))
					v10 = t668
				}
			l707:
				v12 = v13 + i32(1)
				v16 = v4 + v13
				if uint32(v10&i32(0xffff)) < uint32(i32(2048)) {
					m.memory[uint32(v4+v12)] = byte(v10&i32(63) | i32(128))
					m.memory[uint32(v16)] = byte(int32(uint32(v10)>>6) | i32(192))
					t669 := v7
					v17 = v13 | i32(2)
					store32(m.memory[int64(uint32(t669))+60:], uint32(v17))
					goto l788
				}
				m.memory[int64(uint32(v16))+2] = byte(v10&i32(63) | i32(128))
				m.memory[uint32(v4+v12)] = byte(int32(uint32(v10)>>6)&i32(63) | i32(128))
				m.memory[uint32(v16)] = byte(int32(uint32(v10&i32(61440))>>12) | i32(224))
				store32(m.memory[int64(uint32(v7))+60:], uint32(v17))
				goto l788
			l696:
				v11 = v13 + i32(1)
				v12 = v4 + v13
				t670 := int32(load16(m.memory[int64(uint32(v16<<1))+1197822:]))
				v10 = t670
				if v16 != i32(27) {
					goto l709
				}
				m.memory[int64(uint32(v12))+2] = byte(v10&i32(63) | i32(128))
				m.memory[uint32(v4+v11)] = byte(int32(uint32(v10)>>6)&i32(63) | i32(128))
				m.memory[uint32(v12)] = byte(int32(uint32(v10&i32(61440))>>12) | i32(224))
				store32(m.memory[int64(uint32(v7))+60:], uint32(v17))
				goto l788
			l709:
				m.memory[uint32(v4+v11)] = byte(v10&i32(63) | i32(128))
				m.memory[uint32(v12)] = byte(int32(uint32(v10)>>6) | i32(192))
				t671 := v7
				v17 = v13 | i32(2)
				store32(m.memory[int64(uint32(t671))+60:], uint32(v17))
			}
		l788:
			{
				{
					{
						if uint32(v5) < uint32(v17) {
							m.fn121(v17, v5, v5, i32(1146220))
							panic("unreachable")
						}
						v13 = v5 - v17
						t672 := v13
						v10 = v3 - v15
						t673 := v10
						var p674 int32
						if uint32(v13) < uint32(v10) {
							p674 = 1
						}
						v14 = p674
						p675 := t673
						if v14 != 0 {
							p675 = t672
						}
						v12 = p675
						v13 = i32(0)
						v11 = v4 + v17
						t676 := v11
						v16 = v2 + v15
						if (t676^v16)&i32(3) != 0 {
							goto l711
						}
						v13 = i32(0)
						v18 = (i32(0) - v16) & i32(3)
						if uint32(v18|i32(8)) > uint32(v12) {
							goto l711
						}
						if v18 != 0 {
							v13 = i32(0)
							t677 := int32(int8(m.memory[uint32(v16)]))
							v10 = t677
							if v10 < i32(0) {
								goto l714
							}
							m.memory[uint32(v11)] = byte(v10)
							v13 = i32(1)
							if v18 == i32(1) {
								goto l715
							}
							{
								t678 := int32(int8(m.memory[int64(uint32(v16))+1]))
								v10 = t678
								if v10 >= i32(0) {
									goto l716
								}
								v13 = i32(1)
								goto l714
							}
						l716:
							m.memory[int64(uint32(v11))+1] = byte(v10)
							v13 = i32(2)
							if v18 == i32(2) {
								goto l715
							}
							{
								t679 := int32(int8(m.memory[int64(uint32(v16))+2]))
								v10 = t679
								if v10 >= i32(0) {
									goto l717
								}
								v13 = i32(2)
								goto l714
							}
						l717:
							m.memory[int64(uint32(v11))+2] = byte(v10)
							v13 = i32(3)
						l715:
							v8 = v12 + i32(-8)
							goto l721
						}
						v13 = i32(0)
						v8 = v12 + i32(-8)
						goto l721
					}
				l721:
					{
						v18 = v16 + v13
						t680 := int32(load32(m.memory[uint32(v18):]))
						v10 = t680
						v9 = v11 + v13
						t681 := int32(load32(m.memory[uint32(v18+i32(4)):]))
						t682 := v9 + i32(4)
						v18 = t681
						store32(m.memory[uint32(t682):], uint32(v18))
						store32(m.memory[uint32(v9):], uint32(v10))
						{
							v18 = v18 & i32(-2139062144)
							t683 := v18
							v10 = v10 & i32(-2139062144)
							if t683|v10 == 0 {
								goto l718
							}
							if v10 != 0 {
								goto l719
							}
							v10 = int32(uint32(int32(bits.TrailingZeros32(uint32(v18))))>>3) + i32(4)
							goto l720
						l719:
							v10 = int32(uint32(int32(bits.TrailingZeros32(uint32(v10)))) >> 3)
						l720:
							t684 := v16
							v13 = v10 + v13
							t685 := int32(m.memory[uint32(t684+v13)])
							v10 = t685
							goto l714
						}
					l718:
						v13 = v13 + i32(8)
						if uint32(v13) <= uint32(v8) {
							goto l721
						}
					}
				l711:
					if uint32(v13) >= uint32(v12) {
						goto l722
					}
				l723:
					{
						t686 := int32(int8(m.memory[uint32(v16+v13)]))
						v10 = t686
						if v10 < i32(0) {
							goto l714
						}
						m.memory[uint32(v11+v13)] = byte(v10)
						t687 := v12
						v13 = v13 + i32(1)
						if t687 != v13 {
							goto l723
						}
					}
				l722:
					v16 = v12 + v17
					v13 = v12 + v15
					goto l724
				l714:
					t688 := v7
					v16 = v13 + v17
					store32(m.memory[int64(uint32(t688))+60:], uint32(v16))
					v13 = v13 + v15
					if uint32(v16+i32(3)) < uint32(v5) {
						goto l725
					}
					v14 = i32(1)
				}
			l724:
				store32(m.memory[uint32(v0):], uint32(v13))
				store32(m.memory[int64(uint32(v0))+4:], uint32(v14|(v7+i32(52))&i32(-256)))
				store32(m.memory[int64(uint32(v0))+8:], uint32(v16))
				goto l25
			l725:
				v15 = v13 + i32(1)
			l787:
				{
					{
						{
							{
								v17 = v10 + i32(127)
								v13 = v17 & i32(255)
								if uint32(v13) > uint32(i32(125)) {
									if v10&i32(255) == i32(128) {
										m.memory[uint32(v4+v16)] = byte(i32(226))
										t849 := int32(load32(m.memory[int64(uint32(v7))+52:]))
										v4 = t849
										t850 := int32(load32(m.memory[int64(uint32(v7))+60:]))
										t851 := v4
										v5 = t850
										m.memory[uint32(t851+v5+i32(1))] = byte(i32(130))
										t852 := v7
										v16 = v5 + i32(2)
										store32(m.memory[int64(uint32(t852))+60:], uint32(v16))
										v10 = i32(172)
										goto l778
									}
									m.memory[int64(uint32(v0))+6] = byte(i32(0))
									store16(m.memory[int64(uint32(v0))+4:], uint16(i32(258)))
									goto l729
								}
								if uint32(v15) < uint32(v3) {
									v5 = v15 + i32(1)
									{
										t689 := int32(int8(m.memory[uint32(v2+v15)]))
										v12 = t689
										v18 = v12 + i32(-48)
										v11 = v18 & i32(255)
										if uint32(v11) > uint32(i32(9)) {
											if uint32(v13) > uint32(i32(31)) {
												v11 = (v12 + i32(95)) & i32(255)
												if uint32(v11) < uint32(i32(94)) {
													v15 = (v10 + i32(80)) & i32(255)
													if uint32(v15) < uint32(i32(72)) {
														t843 := int32(load16(m.memory[int64(uint32(v15*i32(188)+v11<<1))+1146606:]))
														t844 := v4 + v16
														v13 = t843
														m.memory[uint32(t844)] = byte(int32(uint32(v13)>>12) | i32(224))
														t845 := int32(load32(m.memory[int64(uint32(v7))+60:]))
														t846 := v7
														v10 = t845
														v15 = v10 + i32(1)
														store32(m.memory[int64(uint32(t846))+60:], uint32(v15))
														t847 := int32(load32(m.memory[int64(uint32(v7))+52:]))
														t848 := v10
														v4 = t847
														m.memory[uint32(t848+v4+i32(2))] = byte(v13&i32(63) | i32(128))
														m.memory[uint32(v4+v15)] = byte(int32(uint32(v13)>>6)&i32(63) | i32(128))
														v16 = v10 + i32(3)
														goto l741
													}
													switch v10&i32(255) + i32(-161) {
													case 0:
														v15 = v16 + i32(1)
														v10 = v4 + v16
														{
															t835 := int32(load16(m.memory[int64(uint32(v11<<1))+1219264:]))
															v13 = t835
															if uint32(v13) < uint32(i32(2048)) {
																m.memory[uint32(v4+v15)] = byte(v13&i32(63) | i32(128))
																m.memory[uint32(v10)] = byte(int32(uint32(v13)>>6) | i32(192))
																goto l747
															}
															m.memory[uint32(v10)] = byte(int32(uint32(v13)>>12) | i32(224))
															m.memory[uint32(v10+i32(2))] = byte(v13&i32(63) | i32(128))
															m.memory[uint32(v4+v15)] = byte(int32(uint32(v13)>>6)&i32(63) | i32(128))
															v16 = v16 + i32(3)
															goto l741
														}
													case 5:
														v13 = (v12 + i32(32)) & i32(255)
														if uint32(v13) > uint32(i32(21)) {
															goto l775
														}
														v10 = v4 + v16
														t836 := int32(load16(m.memory[int64(uint32(v13<<1))+1252566:]))
														t837 := v10 + i32(2)
														v13 = t836
														m.memory[uint32(t837)] = byte(v13&i32(63) | i32(128))
														m.memory[uint32(v10)] = byte(int32(uint32(v13)>>12) | i32(224))
														m.memory[uint32(v10+i32(1))] = byte(int32(uint32(v13)>>6)&i32(63) | i32(128))
														v16 = v16 + i32(3)
														goto l741
													case 7:
														if uint32(v11) < uint32(i32(32)) {
															goto l776
														}
														goto l775
													default:
														if uint32(v13) <= uint32(i32(118)) {
															goto l775
														}
														m.memory[uint32(v4+v16)] = byte(i32(238))
														t838 := int32(load32(m.memory[int64(uint32(v7))+60:]))
														t839 := v7
														v13 = t838
														v15 = v13 + i32(1)
														store32(m.memory[int64(uint32(t839))+60:], uint32(v15))
														t840 := int32(load32(m.memory[int64(uint32(v7))+52:]))
														t841 := v13
														v4 = t840
														t842 := t841 + v4 + i32(2)
														v10 = (v10+i32(8))&i32(255)*i32(94) + v11 + i32(-7628)
														m.memory[uint32(t842)] = byte(v10&i32(63) | i32(128))
														m.memory[uint32(v4+v15)] = byte(int32(uint32(v10)>>6) & i32(191))
														v16 = v13 + i32(3)
														goto l741
													}
												}
												v13 = v12 + i32(-64)
												if uint32(v13&i32(255)) <= uint32(i32(62)) {
													goto l756
												}
												if v12 > i32(-96) {
													m.memory[int64(uint32(v0))+4] = byte(i32(2))
													if v12 > i32(-1) {
														store32(m.memory[uint32(v0):], uint32(v15))
														store16(m.memory[int64(uint32(v0))+5:], uint16(i32(1)))
														store32(m.memory[int64(uint32(v0))+8:], uint32(v16))
														goto l25
													}
													store32(m.memory[uint32(v0):], uint32(v5))
													store16(m.memory[int64(uint32(v0))+5:], uint16(i32(2)))
													store32(m.memory[int64(uint32(v0))+8:], uint32(v16))
													goto l25
												}
												v13 = v12 + i32(-65)
											l756:
												{
													v13 = (v10+i32(95))&i32(255)*i32(96) + v13&i32(255)
													v10 = v13 + i32(-864)
													if uint32(v10) < uint32(i32(8059)) {
														{
															{
																p770 := i32(813)
																if uint32(v13) < uint32(i32(3734)) {
																	p770 = i32(0)
																}
																v13 = p770
																t771 := v13
																v13 = v13 + i32(407)
																t772 := int32(load16(m.memory[int64(uint32(v13<<1))+1259814:]))
																t773 := v13
																v13 = v10 & i32(0xffff)
																p774 := t773
																if uint32(t772) > uint32(v13) {
																	p774 = t771
																}
																v15 = p774
																t775 := v15
																v15 = v15 + i32(203)
																t776 := int32(load16(m.memory[int64(uint32(v15<<1))+1259814:]))
																p777 := v15
																if uint32(t776) > uint32(v13) {
																	p777 = t775
																}
																v15 = p777
																t778 := v15
																v15 = v15 + i32(102)
																t779 := int32(load16(m.memory[int64(uint32(v15<<1))+1259814:]))
																p780 := v15
																if uint32(t779) > uint32(v13) {
																	p780 = t778
																}
																v15 = p780
																t781 := v15
																v15 = v15 + i32(51)
																t782 := int32(load16(m.memory[int64(uint32(v15<<1))+1259814:]))
																p783 := v15
																if uint32(t782) > uint32(v13) {
																	p783 = t781
																}
																v15 = p783
																t784 := v15
																v15 = v15 + i32(25)
																t785 := int32(load16(m.memory[int64(uint32(v15<<1))+1259814:]))
																p786 := v15
																if uint32(t785) > uint32(v13) {
																	p786 = t784
																}
																v15 = p786
																t787 := v15
																v15 = v15 + i32(13)
																t788 := int32(load16(m.memory[int64(uint32(v15<<1))+1259814:]))
																p789 := v15
																if uint32(t788) > uint32(v13) {
																	p789 = t787
																}
																v15 = p789
																t790 := v15
																v15 = v15 + i32(6)
																t791 := int32(load16(m.memory[int64(uint32(v15<<1))+1259814:]))
																p792 := v15
																if uint32(t791) > uint32(v13) {
																	p792 = t790
																}
																v15 = p792
																t793 := v15
																v15 = v15 + i32(3)
																t794 := int32(load16(m.memory[int64(uint32(v15<<1))+1259814:]))
																p795 := v15
																if uint32(t794) > uint32(v13) {
																	p795 = t793
																}
																v15 = p795
																t796 := v15
																v15 = v15 + i32(2)
																t797 := int32(load16(m.memory[int64(uint32(v15<<1))+1259814:]))
																p798 := v15
																if uint32(t797) > uint32(v13) {
																	p798 = t796
																}
																v15 = p798
																t799 := v15
																v15 = v15 + i32(1)
																t800 := int32(load16(m.memory[int64(uint32(v15<<1))+1259814:]))
																p801 := v15
																if uint32(t800) > uint32(v13) {
																	p801 = t799
																}
																v15 = p801
																v12 = v15 << 1
																t802 := int32(load16(m.memory[int64(uint32(v12))+1259814:]))
																v11 = t802
																if v11 == v13 {
																	goto l761
																}
																{
																	t803 := v15
																	var p804 int32
																	if uint32(v11) >= uint32(v13) {
																		p804 = 1
																	}
																	v13 = t803 - p804
																	if uint32(v13) >= uint32(i32(1627)) {
																		m.fn33(i32(-1), i32(1627), i32(1227336))
																		panic("unreachable")
																	}
																	t805 := v10
																	v13 = v13 << 1
																	t806 := int32(load16(m.memory[int64(uint32(v13))+1259814:]))
																	t807 := int32(load16(m.memory[int64(uint32(v13))+1252610:]))
																	v10 = t805 - t806 + t807
																	v15 = v10 & i32(0xffff)
																	v13 = int32(uint32(v15) >> 12)
																	v15 = int32(uint32(v15) >> 6)
																	goto l763
																}
															}
														l761:
															t808 := int32(load16(m.memory[int64(uint32(v12))+1252610:]))
															v10 = t808
															v13 = int32(uint32(v10) >> 12)
															v15 = int32(uint32(v10) >> 6)
														}
													l763:
														m.memory[uint32(v4+v16)] = byte(v13 | i32(224))
														t809 := int32(load32(m.memory[int64(uint32(v7))+52:]))
														v4 = t809
														t810 := int32(load32(m.memory[int64(uint32(v7))+60:]))
														t811 := v4
														v13 = t810
														m.memory[uint32(t811+v13+i32(1))] = byte(v15&i32(63) | i32(128))
														t812 := v7
														v15 = v13 + i32(2)
														store32(m.memory[int64(uint32(t812))+60:], uint32(v15))
														m.memory[uint32(v4+v15)] = byte(v10&i32(63) | i32(128))
														v16 = v13 + i32(3)
														goto l741
													}
													if uint32(v13) < uint32(i32(864)) {
														{
															{
																p813 := i32(29)
																if uint32(v13) < uint32(i32(777)) {
																	p813 = i32(0)
																}
																v10 = p813
																t814 := v10
																v10 = v10 + i32(15)
																t815 := int32(load16(m.memory[int64(uint32(v10<<1))+1234878:]))
																p816 := v10
																if uint32(t815) > uint32(v13) {
																	p816 = t814
																}
																v10 = p816
																t817 := v10
																v10 = v10 + i32(7)
																t818 := int32(load16(m.memory[int64(uint32(v10<<1))+1234878:]))
																p819 := v10
																if uint32(t818) > uint32(v13) {
																	p819 = t817
																}
																v10 = p819
																t820 := v10
																v10 = v10 + i32(4)
																t821 := int32(load16(m.memory[int64(uint32(v10<<1))+1234878:]))
																p822 := v10
																if uint32(t821) > uint32(v13) {
																	p822 = t820
																}
																v10 = p822
																t823 := v10
																v10 = v10 + i32(2)
																t824 := int32(load16(m.memory[int64(uint32(v10<<1))+1234878:]))
																p825 := v10
																if uint32(t824) > uint32(v13) {
																	p825 = t823
																}
																v10 = p825
																t826 := v10
																v10 = v10 + i32(1)
																t827 := int32(load16(m.memory[int64(uint32(v10<<1))+1234878:]))
																p828 := v10
																if uint32(t827) > uint32(v13) {
																	p828 = t826
																}
																v10 = p828
																v15 = v10 << 1
																t829 := int32(load16(m.memory[int64(uint32(v15))+1234878:]))
																v12 = t829
																if v12 == v13 {
																	goto l764
																}
																{
																	t830 := v10
																	var p831 int32
																	if uint32(v12) >= uint32(v13) {
																		p831 = 1
																	}
																	v10 = t830 - p831
																	if uint32(v10) >= uint32(i32(59)) {
																		m.fn33(i32(-1), i32(59), i32(1227336))
																		panic("unreachable")
																	}
																	v10 = v10 << 1
																	t832 := int32(load16(m.memory[int64(uint32(v10))+1255864:]))
																	t833 := int32(load16(m.memory[int64(uint32(v10))+1234878:]))
																	v13 = t832 + v13 - t833
																	goto l766
																}
															}
														l764:
															t834 := int32(load16(m.memory[int64(uint32(v15))+1255864:]))
															v13 = t834
														}
													l766:
														v15 = v16 + i32(1)
														v10 = v4 + v16
														if uint32(v13&i32(0xffff)) < uint32(i32(2048)) {
															m.memory[uint32(v4+v15)] = byte(v13&i32(63) | i32(128))
															m.memory[uint32(v10)] = byte(int32(uint32(v13)>>6) | i32(192))
															goto l747
														}
														m.memory[uint32(v10+i32(2))] = byte(v13&i32(63) | i32(128))
														m.memory[uint32(v4+v15)] = byte(int32(uint32(v13)>>6)&i32(63) | i32(128))
														m.memory[uint32(v10)] = byte(int32(uint32(v13&i32(61440))>>12) | i32(224))
														v16 = v16 + i32(3)
														goto l741
													}
													v13 = v13 + i32(-8923)
													if uint32(v13) >= uint32(i32(101)) {
														m.fn33(v13, i32(101), i32(1146252))
														panic("unreachable")
													}
													t764 := int32(load16(m.memory[int64(uint32(v13<<1))+1146268:]))
													t765 := v4 + v16
													v13 = t764
													m.memory[uint32(t765)] = byte(int32(uint32(v13)>>12) | i32(224))
													t766 := int32(load32(m.memory[int64(uint32(v7))+52:]))
													v4 = t766
													t767 := int32(load32(m.memory[int64(uint32(v7))+60:]))
													t768 := v4
													v10 = t767
													m.memory[uint32(t768+v10+i32(1))] = byte(int32(uint32(v13)>>6)&i32(63) | i32(128))
													t769 := v7
													v15 = v10 + i32(2)
													store32(m.memory[int64(uint32(t769))+60:], uint32(v15))
													m.memory[uint32(v4+v15)] = byte(v13&i32(63) | i32(128))
													v16 = v10 + i32(3)
													goto l741
												}
											}
											{
												v10 = v12 + i32(-64)
												if uint32(v10&i32(255)) <= uint32(i32(62)) {
													goto l749
												}
												if v12 > i32(-2) {
													m.memory[int64(uint32(v0))+4] = byte(i32(2))
													if v12 > i32(-1) {
														store32(m.memory[uint32(v0):], uint32(v15))
														store16(m.memory[int64(uint32(v0))+5:], uint16(i32(1)))
														store32(m.memory[int64(uint32(v0))+8:], uint32(v16))
														goto l25
													}
													store32(m.memory[uint32(v0):], uint32(v5))
													store16(m.memory[int64(uint32(v0))+5:], uint16(i32(2)))
													store32(m.memory[int64(uint32(v0))+8:], uint32(v16))
													goto l25
												}
												v10 = v12 + i32(-65)
											l749:
												{
													{
														v11 = v13*i32(190) + v10&i32(255)
														v13 = v11 & i32(0xffff)
														p722 := i32(958)
														if uint32(v13) < uint32(i32(2880)) {
															p722 = i32(0)
														}
														v10 = p722
														t723 := v10
														v10 = v10 + i32(479)
														t724 := int32(load16(m.memory[int64(uint32(v10<<1))+1255982:]))
														p725 := v10
														if uint32(t724) > uint32(v13) {
															p725 = t723
														}
														v10 = p725
														t726 := v10
														v10 = v10 + i32(239)
														t727 := int32(load16(m.memory[int64(uint32(v10<<1))+1255982:]))
														p728 := v10
														if uint32(t727) > uint32(v13) {
															p728 = t726
														}
														v10 = p728
														t729 := v10
														v10 = v10 + i32(120)
														t730 := int32(load16(m.memory[int64(uint32(v10<<1))+1255982:]))
														p731 := v10
														if uint32(t730) > uint32(v13) {
															p731 = t729
														}
														v10 = p731
														t732 := v10
														v10 = v10 + i32(60)
														t733 := int32(load16(m.memory[int64(uint32(v10<<1))+1255982:]))
														p734 := v10
														if uint32(t733) > uint32(v13) {
															p734 = t732
														}
														v10 = p734
														t735 := v10
														v10 = v10 + i32(30)
														t736 := int32(load16(m.memory[int64(uint32(v10<<1))+1255982:]))
														p737 := v10
														if uint32(t736) > uint32(v13) {
															p737 = t735
														}
														v10 = p737
														t738 := v10
														v10 = v10 + i32(15)
														t739 := int32(load16(m.memory[int64(uint32(v10<<1))+1255982:]))
														p740 := v10
														if uint32(t739) > uint32(v13) {
															p740 = t738
														}
														v10 = p740
														t741 := v10
														v10 = v10 + i32(7)
														t742 := int32(load16(m.memory[int64(uint32(v10<<1))+1255982:]))
														p743 := v10
														if uint32(t742) > uint32(v13) {
															p743 = t741
														}
														v10 = p743
														t744 := v10
														v10 = v10 + i32(4)
														t745 := int32(load16(m.memory[int64(uint32(v10<<1))+1255982:]))
														p746 := v10
														if uint32(t745) > uint32(v13) {
															p746 = t744
														}
														v10 = p746
														t747 := v10
														v10 = v10 + i32(2)
														t748 := int32(load16(m.memory[int64(uint32(v10<<1))+1255982:]))
														p749 := v10
														if uint32(t748) > uint32(v13) {
															p749 = t747
														}
														v10 = p749
														t750 := v10
														v10 = v10 + i32(1)
														t751 := int32(load16(m.memory[int64(uint32(v10<<1))+1255982:]))
														p752 := v10
														if uint32(t751) > uint32(v13) {
															p752 = t750
														}
														v10 = p752
														v15 = v10 << 1
														t753 := int32(load16(m.memory[int64(uint32(v15))+1255982:]))
														v12 = t753
														if v12 == v13 {
															goto l751
														}
														{
															t754 := v10
															var p755 int32
															if uint32(v12) >= uint32(v13) {
																p755 = 1
															}
															v13 = t754 - p755
															if uint32(v13) >= uint32(i32(1916)) {
																m.fn33(i32(-1), i32(1916), i32(1227336))
																panic("unreachable")
															}
															t756 := v11
															v13 = v13 << 1
															t757 := int32(load16(m.memory[int64(uint32(v13))+1255982:]))
															t758 := int32(load16(m.memory[int64(uint32(v13))+1247664:]))
															v10 = t756 - t757 + t758
															v15 = v10 & i32(0xffff)
															v13 = int32(uint32(v15) >> 12)
															v15 = int32(uint32(v15) >> 6)
															goto l753
														}
													}
												l751:
													t759 := int32(load16(m.memory[int64(uint32(v15))+1247664:]))
													v10 = t759
													v13 = int32(uint32(v10) >> 12)
													v15 = int32(uint32(v10) >> 6)
												}
											l753:
												m.memory[uint32(v4+v16)] = byte(v13 | i32(224))
												t760 := int32(load32(m.memory[int64(uint32(v7))+52:]))
												v4 = t760
												t761 := int32(load32(m.memory[int64(uint32(v7))+60:]))
												t762 := v4
												v13 = t761
												m.memory[uint32(t762+v13+i32(1))] = byte(v15&i32(63) | i32(128))
												t763 := v7
												v15 = v13 + i32(2)
												store32(m.memory[int64(uint32(t763))+60:], uint32(v15))
												m.memory[uint32(v4+v15)] = byte(v10&i32(63) | i32(128))
												v16 = v13 + i32(3)
												goto l741
											}
										}
										if uint32(v5) < uint32(v3) {
											t690 := int32(m.memory[uint32(v2+v5)])
											v9 = t690 + i32(127)
											v10 = v9 & i32(255)
											if uint32(v10) > uint32(i32(125)) {
												m.memory[int64(uint32(v1))+8] = byte(v12)
												m.memory[int64(uint32(v1))+7] = byte(i32(1))
												m.memory[int64(uint32(v0))+6] = byte(i32(1))
												store16(m.memory[int64(uint32(v0))+4:], uint16(i32(258)))
												goto l733
											}
											v5 = v15 + i32(2)
											if uint32(v5) < uint32(v3) {
												t691 := int32(m.memory[uint32(v2+v5)])
												v17 = (t691 + i32(-48)) & i32(255)
												if uint32(v17) > uint32(i32(9)) {
													m.memory[int64(uint32(v1))+10] = byte(v9)
													m.memory[int64(uint32(v1))+9] = byte(i32(1))
													m.memory[int64(uint32(v1))+8] = byte(v12)
													m.memory[int64(uint32(v1))+7] = byte(i32(1))
													m.memory[int64(uint32(v0))+6] = byte(i32(2))
													store16(m.memory[int64(uint32(v0))+4:], uint16(i32(258)))
													goto l737
												}
												v5 = v15 + i32(3)
												v13 = v11*i32(1260) + v13*i32(12600) + v10*i32(10) + v17
												if uint32(v13) < uint32(i32(39420)) {
													if v13 == i32(7457) {
														goto l742
													}
													{
														{
															p693 := i32(103)
															if uint32(v13) < uint32(i32(11334)) {
																p693 = i32(0)
															}
															v10 = p693
															t694 := v10
															v10 = v10 + i32(51)
															t695 := int32(load16(m.memory[int64(uint32(v10<<1))+1241560:]))
															t696 := v10
															v10 = v13 & i32(0xffff)
															p697 := t696
															if uint32(t695) > uint32(v10) {
																p697 = t694
															}
															v15 = p697
															t698 := v15
															v15 = v15 + i32(26)
															t699 := int32(load16(m.memory[int64(uint32(v15<<1))+1241560:]))
															p700 := v15
															if uint32(t699) > uint32(v10) {
																p700 = t698
															}
															v15 = p700
															t701 := v15
															v15 = v15 + i32(13)
															t702 := int32(load16(m.memory[int64(uint32(v15<<1))+1241560:]))
															p703 := v15
															if uint32(t702) > uint32(v10) {
																p703 = t701
															}
															v15 = p703
															t704 := v15
															v15 = v15 + i32(6)
															t705 := int32(load16(m.memory[int64(uint32(v15<<1))+1241560:]))
															p706 := v15
															if uint32(t705) > uint32(v10) {
																p706 = t704
															}
															v15 = p706
															t707 := v15
															v15 = v15 + i32(3)
															t708 := int32(load16(m.memory[int64(uint32(v15<<1))+1241560:]))
															p709 := v15
															if uint32(t708) > uint32(v10) {
																p709 = t707
															}
															v15 = p709
															t710 := v15
															v15 = v15 + i32(2)
															t711 := int32(load16(m.memory[int64(uint32(v15<<1))+1241560:]))
															p712 := v15
															if uint32(t711) > uint32(v10) {
																p712 = t710
															}
															v15 = p712
															t713 := v15
															v15 = v15 + i32(1)
															t714 := int32(load16(m.memory[int64(uint32(v15<<1))+1241560:]))
															p715 := v15
															if uint32(t714) > uint32(v10) {
																p715 = t713
															}
															v15 = p715
															v12 = v15 << 1
															t716 := int32(load16(m.memory[int64(uint32(v12))+1241560:]))
															v11 = t716
															if v11 == v10 {
																goto l743
															}
															{
																t717 := v15
																var p718 int32
																if uint32(v11) >= uint32(v10) {
																	p718 = 1
																}
																v10 = t717 - p718
																if uint32(v10) >= uint32(i32(206)) {
																	m.fn33(i32(-1), i32(206), i32(1227336))
																	panic("unreachable")
																}
																v10 = v10 << 1
																t719 := int32(load16(m.memory[int64(uint32(v10))+1240928:]))
																t720 := int32(load16(m.memory[int64(uint32(v10))+1241560:]))
																v13 = t719 + v13 - t720
																goto l745
															}
														}
													l743:
														t721 := int32(load16(m.memory[int64(uint32(v12))+1240928:]))
														v13 = t721
													}
												l745:
													v15 = v16 + i32(1)
													v10 = v4 + v16
													if uint32(v13&i32(0xffff)) < uint32(i32(2048)) {
														m.memory[uint32(v4+v15)] = byte(v13&i32(63) | i32(128))
														m.memory[uint32(v10)] = byte(int32(uint32(v13)>>6) | i32(192))
														goto l747
													}
													m.memory[uint32(v10+i32(2))] = byte(v13&i32(63) | i32(128))
													m.memory[uint32(v4+v15)] = byte(int32(uint32(v13)>>6)&i32(63) | i32(128))
													m.memory[uint32(v10)] = byte(int32(uint32(v13&i32(61440))>>12) | i32(224))
													v16 = v16 + i32(3)
													goto l741
												}
												if uint32(v13+i32(-189000)) < uint32(i32(0x100000)) {
													v10 = v4 + v16
													t692 := v10 + i32(3)
													v13 = v13 + i32(-123464)
													m.memory[uint32(t692)] = byte(v13&i32(63) | i32(128))
													m.memory[uint32(v10)] = byte(int32(uint32(v13)>>18) | i32(240))
													m.memory[uint32(v10+i32(2))] = byte(int32(uint32(v13)>>6)&i32(63) | i32(128))
													m.memory[uint32(v10+i32(1))] = byte(int32(uint32(v13)>>12)&i32(63) | i32(128))
													v16 = v16 + i32(4)
													goto l741
												}
												m.memory[int64(uint32(v0))+6] = byte(i32(0))
												store16(m.memory[int64(uint32(v0))+4:], uint16(i32(1026)))
												goto l733
											}
											if v6 != 0 {
												m.memory[int64(uint32(v0))+6] = byte(i32(0))
												store16(m.memory[int64(uint32(v0))+4:], uint16(i32(770)))
												goto l737
											}
											m.memory[int64(uint32(v1))+12] = byte(v9)
											m.memory[int64(uint32(v1))+11] = byte(v18)
											m.memory[int64(uint32(v1))+10] = byte(v17)
											m.memory[int64(uint32(v1))+9] = byte(i32(3))
											m.memory[int64(uint32(v0))+4] = byte(i32(0))
											goto l737
										}
										if v6 != 0 {
											m.memory[int64(uint32(v0))+6] = byte(i32(0))
											store16(m.memory[int64(uint32(v0))+4:], uint16(i32(514)))
											goto l733
										}
										m.memory[int64(uint32(v1))+11] = byte(v18)
										m.memory[int64(uint32(v1))+10] = byte(v17)
										m.memory[int64(uint32(v1))+9] = byte(i32(2))
										m.memory[int64(uint32(v0))+4] = byte(i32(0))
										goto l733
									}
								}
								if v6 != 0 {
									m.memory[int64(uint32(v0))+6] = byte(i32(0))
									store16(m.memory[int64(uint32(v0))+4:], uint16(i32(258)))
									goto l729
								}
								m.memory[int64(uint32(v1))+10] = byte(v17)
								m.memory[int64(uint32(v1))+9] = byte(i32(1))
								m.memory[int64(uint32(v0))+4] = byte(i32(0))
								goto l729
							l775:
								t853 := v7 + i32(8)
								v10 = (v10+i32(94))&i32(255)*i32(94) + v11
								m.fn897(t853, i32(1241340), i32(46), v10)
								t854 := int32(load32(m.memory[int64(uint32(v7))+12:]))
								v13 = t854
								{
									{
										t855 := int32(load32(m.memory[int64(uint32(v7))+8:]))
										if t855 != i32(1) {
											goto l779
										}
										v13 = v13 + i32(-1)
										if uint32(v13) >= uint32(i32(46)) {
											m.fn33(v13, i32(46), i32(1227336))
											panic("unreachable")
										}
										v13 = v13 << 1
										t856 := int32(load16(m.memory[int64(uint32(v13))+1263068:]))
										t857 := int32(load16(m.memory[int64(uint32(v13))+1241340:]))
										v13 = t856 + v10 - t857
										goto l781
									}
								l779:
									if uint32(v13) > uint32(i32(45)) {
										m.fn33(v13, i32(46), i32(1227320))
										panic("unreachable")
									}
									t858 := int32(load16(m.memory[int64(uint32(v13<<1))+1263068:]))
									v13 = t858
								}
							l781:
								v15 = v16 + i32(1)
								v10 = v4 + v16
								if uint32(v13&i32(0xffff)) < uint32(i32(2048)) {
									m.memory[uint32(v4+v15)] = byte(v13&i32(63) | i32(128))
									m.memory[uint32(v10)] = byte(int32(uint32(v13)>>6) | i32(192))
									goto l747
								}
								m.memory[uint32(v10+i32(2))] = byte(v13&i32(63) | i32(128))
								m.memory[uint32(v4+v15)] = byte(int32(uint32(v13)>>6)&i32(63) | i32(128))
								m.memory[uint32(v10)] = byte(int32(uint32(v13&i32(61440))>>12) | i32(224))
								v16 = v16 + i32(3)
								goto l741
							}
						l776:
							v15 = v16 + i32(1)
							v10 = v4 + v16
							t859 := int32(load16(m.memory[int64(uint32(v11<<1))+1197822:]))
							v13 = t859
							if v11 != i32(27) {
								goto l784
							}
							m.memory[uint32(v10+i32(2))] = byte(v13&i32(63) | i32(128))
							m.memory[uint32(v4+v15)] = byte(int32(uint32(v13)>>6)&i32(63) | i32(128))
							m.memory[uint32(v10)] = byte(int32(uint32(v13&i32(61440))>>12) | i32(224))
							v16 = v16 + i32(3)
							goto l741
						l784:
							m.memory[uint32(v4+v15)] = byte(v13&i32(63) | i32(128))
							m.memory[uint32(v10)] = byte(int32(uint32(v13)>>6) | i32(192))
						}
					l747:
						v16 = v16 + i32(2)
						goto l741
					l742:
						m.memory[uint32(v4+v16)] = byte(i32(238))
						t860 := int32(load32(m.memory[int64(uint32(v7))+52:]))
						v4 = t860
						t861 := int32(load32(m.memory[int64(uint32(v7))+60:]))
						t862 := v4
						v13 = t861
						m.memory[uint32(t862+v13+i32(1))] = byte(i32(159))
						t863 := v7
						v10 = v13 + i32(2)
						store32(m.memory[int64(uint32(t863))+60:], uint32(v10))
						m.memory[uint32(v4+v10)] = byte(i32(135))
						v16 = v13 + i32(3)
					}
				l741:
					store32(m.memory[int64(uint32(v7))+60:], uint32(v16))
					if uint32(v5) < uint32(v3) {
						goto l785
					}
					store32(m.memory[uint32(v0):], uint32(v5))
					m.memory[int64(uint32(v0))+4] = byte(i32(0))
					store32(m.memory[int64(uint32(v0))+8:], uint32(v16))
					goto l25
				l785:
					{
						t864 := int32(load32(m.memory[int64(uint32(v7))+56:]))
						if uint32(v16+i32(3)) < uint32(t864) {
							goto l786
						}
						store32(m.memory[uint32(v0):], uint32(v5))
						m.memory[int64(uint32(v0))+4] = byte(i32(1))
						store32(m.memory[int64(uint32(v0))+8:], uint32(v16))
						goto l25
					}
				l786:
					v15 = v5 + i32(1)
					t865 := int32(int8(m.memory[uint32(v2+v5)]))
					v10 = t865
					if v10 < i32(0) {
						goto l787
					}
				}
			l778:
				m.memory[uint32(v4+v16)] = byte(v10)
				t866 := int32(load32(m.memory[int64(uint32(v7))+60:]))
				t867 := v7
				v17 = t866 + i32(1)
				store32(m.memory[int64(uint32(t867))+60:], uint32(v17))
				t868 := int32(load32(m.memory[int64(uint32(v7))+56:]))
				v5 = t868
				if uint32(v15) <= uint32(v3) {
					goto l788
				}
			}
			m.fn121(v15, v3, v3, i32(1146236))
			panic("unreachable")
		l737:
			store32(m.memory[uint32(v0):], uint32(v5))
			store32(m.memory[int64(uint32(v0))+8:], uint32(v16))
			goto l25
		l733:
			store32(m.memory[uint32(v0):], uint32(v5))
			store32(m.memory[int64(uint32(v0))+8:], uint32(v16))
			goto l25
		l729:
			store32(m.memory[uint32(v0):], uint32(v15))
			store32(m.memory[int64(uint32(v0))+8:], uint32(v16))
			goto l25
		case 1:
			t869 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			v18 = t869
			t870 := int32(m.memory[int64(uint32(v1))+17])
			v14 = t870
			t871 := int32(m.memory[int64(uint32(v1))+16])
			v8 = t871
			t872 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			v10 = t872
			t873 := int32(load32(m.memory[int64(uint32(v1))+12:]))
			v16 = t873
			v15 = i32(0)
			v13 = i32(0)
		l816:
			v9 = v10
		l809:
			{
				{
					{
						if v16 == 0 {
							goto l789
						}
						v10 = v15
						goto l790
					l789:
						if uint32(v3) < uint32(v15) {
							m.fn121(v15, v3, v3, i32(1146172))
							panic("unreachable")
						}
						if uint32(v5) < uint32(v13) {
							m.fn121(v13, v5, v5, i32(1146156))
							panic("unreachable")
						}
						v17 = v2 + v15
						t874 := v17
						v12 = v5 - v13
						t875 := v12
						v11 = v3 - v15
						p876 := v11
						if uint32(v12) < uint32(v11) {
							p876 = t875
						}
						t877 := m.fn886(t874, p876)
						v10 = t877
						if uint32(v10) > uint32(v12) {
							m.fn121(i32(0), v10, v12, i32(1146140))
							panic("unreachable")
						}
						if uint32(v10) > uint32(v11) {
							m.fn121(i32(0), v10, v11, i32(1146124))
							panic("unreachable")
						}
						if v10 == 0 {
							goto l795
						}
						memory_copy(m.memory, uint32(v4+v13), uint32(v17), uint32(v10))
					l795:
						v13 = v10 + v13
						v10 = v10 + v15
					}
				l790:
					if uint32(v10) < uint32(v3) {
						v11 = v13 + i32(3)
						if uint32(v11) < uint32(v5) {
							v15 = v10 + i32(1)
							t878 := int32(m.memory[uint32(v2+v10)])
							v12 = t878
							if v16 != 0 {
								if uint32(v12) < uint32(v8&i32(255)) {
									goto l807
								}
								if uint32(v12) <= uint32(v14&i32(255)) {
									store16(m.memory[int64(uint32(v1))+16:], uint16(i32(49024)))
									t880 := v1
									v18 = v18 + i32(1)
									store32(m.memory[int64(uint32(t880))+8:], uint32(v18))
									t881 := v1
									v17 = v9 << 6
									t882 := v17
									v12 = v12 & i32(63)
									v10 = t882 | v12
									store32(m.memory[int64(uint32(t881))+4:], uint32(v10))
									v8 = i32(128)
									v14 = i32(191)
									if v18 != v16 {
										goto l816
									}
									if v16 != i32(3) {
										goto l817
									}
									v10 = v4 + v13
									m.memory[uint32(v10+i32(3))] = byte(v12 | i32(128))
									m.memory[uint32(v10+i32(2))] = byte(v9&i32(63) | i32(128))
									m.memory[uint32(v10)] = byte(int32(uint32(v17)>>18) | i32(240))
									m.memory[uint32(v10+i32(1))] = byte(int32(uint32(v17)>>12)&i32(63) | i32(128))
									v13 = v13 + i32(4)
									goto l818
								l817:
									if uint32(v10&i32(0xffff)) < uint32(i32(2048)) {
										goto l819
									}
									v13 = v4 + v13
									m.memory[uint32(v13+i32(2))] = byte(v12 | i32(128))
									m.memory[uint32(v13+i32(1))] = byte(v9&i32(63) | i32(128))
									m.memory[uint32(v13)] = byte(int32(uint32(v10&i32(61440))>>12) | i32(224))
									v13 = v11
									goto l818
								l819:
									v10 = v4 + v13
									m.memory[uint32(v10)] = byte(v9 | i32(192))
									m.memory[uint32(v10+i32(1))] = byte(v12 | i32(128))
									v13 = v13 + i32(2)
								l818:
									v18 = i32(0)
									store32(m.memory[int64(uint32(v1))+12:], uint32(i32(0)))
									store64(m.memory[int64(uint32(v1))+4:], uint64(i64(0)))
									v10 = i32(0)
									v16 = i32(0)
									goto l816
								}
							l807:
								store32(m.memory[int64(uint32(v1))+12:], uint32(i32(0)))
								store64(m.memory[int64(uint32(v1))+4:], uint64(i64(0)))
								store16(m.memory[int64(uint32(v1))+16:], uint16(i32(49024)))
								m.memory[int64(uint32(v0))+6] = byte(i32(0))
								m.memory[int64(uint32(v0))+4] = byte(i32(2))
								m.memory[int64(uint32(v0))+5] = byte(v18 + i32(1))
								goto l799
							}
							if int32(int8(v12)) > i32(-1) {
								m.memory[uint32(v4+v13)] = byte(v12)
								v13 = v13 + i32(1)
								v16 = i32(0)
								goto l809
							}
							if uint32(v12) < uint32(i32(194)) {
								m.memory[int64(uint32(v0))+6] = byte(i32(0))
								store16(m.memory[int64(uint32(v0))+4:], uint16(i32(258)))
								v10 = v15
								goto l799
							}
							if uint32(v12) < uint32(i32(224)) {
								v16 = i32(1)
								store32(m.memory[int64(uint32(v1))+12:], uint32(i32(1)))
								t879 := v1
								v9 = v12 & i32(31)
								store32(m.memory[int64(uint32(t879))+4:], uint32(v9))
								goto l809
							}
							if uint32(v12) < uint32(i32(240)) {
								switch v12 + i32(-224) {
								case 0:
									v8 = i32(160)
									m.memory[int64(uint32(v1))+16] = byte(i32(160))
									goto l811
								case 13:
									v14 = i32(159)
									m.memory[int64(uint32(v1))+17] = byte(i32(159))
									goto l811
								default:
									goto l811
								}
							}
							if uint32(v12) < uint32(i32(245)) {
								switch v12 + i32(-240) {
								default:
									goto l814
								case 0:
									v8 = i32(144)
									m.memory[int64(uint32(v1))+16] = byte(i32(144))
									goto l814
								case 4:
									v14 = i32(143)
									m.memory[int64(uint32(v1))+17] = byte(i32(143))
									goto l814
								}
							}
							m.memory[int64(uint32(v0))+6] = byte(i32(0))
							store16(m.memory[int64(uint32(v0))+4:], uint16(i32(258)))
							v10 = v15
							goto l799
						}
						m.memory[int64(uint32(v0))+4] = byte(i32(1))
						goto l799
					}
					if v6 == 0 {
						goto l797
					}
					if v16 != 0 {
						goto l798
					}
				l797:
					m.memory[int64(uint32(v0))+4] = byte(i32(0))
					goto l799
				l798:
					store32(m.memory[int64(uint32(v1))+12:], uint32(i32(0)))
					store64(m.memory[int64(uint32(v1))+4:], uint64(i64(0)))
					m.memory[int64(uint32(v0))+6] = byte(i32(0))
					m.memory[int64(uint32(v0))+4] = byte(i32(2))
					m.memory[int64(uint32(v0))+5] = byte(v18 + i32(1))
				l799:
					store32(m.memory[int64(uint32(v0))+8:], uint32(v13))
					store32(m.memory[uint32(v0):], uint32(v10))
					goto l25
				l811:
					v16 = i32(2)
					store32(m.memory[int64(uint32(v1))+12:], uint32(i32(2)))
					t883 := v1
					v9 = v12 & i32(15)
					store32(m.memory[int64(uint32(t883))+4:], uint32(v9))
					goto l809
				}
			l814:
				v16 = i32(3)
				store32(m.memory[int64(uint32(v1))+12:], uint32(i32(3)))
				t884 := v1
				v9 = v12 & i32(7)
				store32(m.memory[int64(uint32(t884))+4:], uint32(v9))
				goto l809
			}
		}
	}
l25:
	m.g0 = v7 + i32(64)
}
func (m *Module) fn896(v0, v1, v2, v3, v4, v5, v6 int32) {
	var v7, v8, v9 int32
	t0 := m.g0
	v7 = t0 - i32(16)
	m.g0 = v7
	m.memory[int64(uint32(v1))+24] = byte(i32(9))
	switch v6 {
	case 0:
		store16(m.memory[int64(uint32(v7))+2:], uint16(i32(48111)))
		v8 = i32(0)
		v9 = i32(2)
		m.fn895(v7+i32(4), v1, v7+i32(2), i32(2), v4, v5, i32(0))
		t1 := int32(load32(m.memory[int64(uint32(v7))+12:]))
		v6 = t1
		{
			t2 := int32(m.memory[int64(uint32(v7))+8])
			switch t2 {
			case 1:
				m.fn3(i32(1146084), i32(39), i32(1146036))
				panic("unreachable")
			case 2:
				t3 := int32(load16(m.memory[int64(uint32(v7))+9:]))
				v5 = t3
				t4 := int32(load32(m.memory[int64(uint32(v7))+4:]))
				if t4 != i32(1) {
					goto l6
				}
				m.memory[int64(uint32(v1))+24] = byte(i32(8))
				goto l6
			default:
				if uint32(v5) < uint32(v6) {
					m.fn121(v6, v5, v5, i32(1146036))
					panic("unreachable")
				}
				m.fn895(v7+i32(4), v1, v2, v3, v4+v6, v5-v6, i32(1))
				{
					t5 := int32(m.memory[int64(uint32(v7))+8])
					v9 = t5
					if v9 != 0 {
						goto l8
					}
					m.memory[int64(uint32(v1))+24] = byte(i32(10))
				}
			l8:
				t6 := int32(load32(m.memory[int64(uint32(v7))+12:]))
				v6 = t6 + v6
				t7 := int32(load32(m.memory[int64(uint32(v7))+4:]))
				v8 = t7
				t8 := int32(load16(m.memory[int64(uint32(v7))+9:]))
				v5 = t8
			}
		}
	l6:
		store16(m.memory[int64(uint32(v0))+5:], uint16(v5))
		m.memory[int64(uint32(v0))+4] = byte(v9)
		store32(m.memory[int64(uint32(v0))+8:], uint32(v6))
		store32(m.memory[uint32(v0):], uint32(v8))
		goto l9
	case 1:
		m.fn894(v0, v1, v2, v3, v4, v5, i32(0), i32(239))
		goto l9
	default:
		m.fn895(v0, v1, v2, v3, v4, v5, i32(1))
		t9 := int32(m.memory[int64(uint32(v0))+4])
		if t9 != 0 {
			goto l9
		}
		m.memory[int64(uint32(v1))+24] = byte(i32(10))
		goto l9
	}
l9:
	m.g0 = v7 + i32(16)
}
