package core

import (
	"math"
	"math/bits"
)

func (m *Module) fn1662(v0 int32, v1 int64, v2, v3, v4, v5 int32) {
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
				m.fn256(i32(1130120), i32(28), i32(1130148))
				panic("unreachable")
			}
			t1 := v2
			v7 = int64(bits.LeadingZeros64(uint64(v1)))
			v8 = t1 - int32(v7)
			v2 = (int32(int16(i32(-96)-v8))*i32(80) + i32(86960)) / i32(2126)
			if uint32(v2) > uint32(i32(80)) {
				m.fn158(v2, i32(81), i32(1130164))
				panic("unreachable")
			}
			t2 := v6
			v2 = v2 << 4
			t3 := int64(load64(m.memory[int64(uint32(v2))+1128824:]))
			m.fn1853(t2, t3, i64(0), i64_shl(v1, v7), i64(0))
			t4 := int64(load64(m.memory[uint32(v6):]))
			t5 := int64(load64(m.memory[int64(uint32(v6))+8:]))
			v1 = int64(uint64(t4)>>63) + t5
			t6 := int32(load16(m.memory[int64(uint32(v2))+1128832:]))
			t7 := v1
			v9 = i32(-64) - (v8 + t6)
			v7 = int64(uint32(v9))
			v10 = int32(i64_shr_u(t7, v7))
			t8 := int32(load16(m.memory[int64(uint32(v2))+1128834:]))
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
				t9 := int32(load32(m.memory[uint32(v4<<2+i32(1130860)):]))
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
			m.fn1663(v0, v3, v4, i32(0), v16, v5, t29, i64_shl(int64(uint32(v8)), v15), v11)
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
				m.fn158(v4, v4, i32(1130196))
				panic("unreachable")
			}
			v10 = v10 - v9*v8
			m.memory[uint32(v3+v2)] = byte(v9 + i32(48))
			if v13 == v2 {
				m.fn1663(v0, v3, v4, v18, v16, v5, i64_shl(int64(uint32(v10)), v15)+v7, i64_shl(int64(uint32(v8)), v15), v11)
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
						m.fn158(v2, v4, i32(1130212))
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
				m.fn1663(v0, v3, v4, v18, v16, v5, v7, v11, v1)
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
		m.fn494(i32(1130180))
		panic("unreachable")
	}
l3:
	store32(m.memory[uint32(v0):], uint32(i32(0)))
l14:
	m.g0 = v6 + i32(16)
}
func (m *Module) fn1663(v0, v1, v2, v3, v4, v5 int32, v6, v7, v8 int64) {
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
			m.fn151(i32(0), v3, v2, i32(1131092))
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
			m.fn151(i32(0), v3, v2, i32(1131076))
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
	m.fn151(i32(0), v3, v2, i32(1131060))
	panic("unreachable")
l11:
	store16(m.memory[int64(uint32(v0))+8:], uint16(v4))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
	store32(m.memory[uint32(v0):], uint32(v1))
	return
l4:
	store32(m.memory[uint32(v0):], uint32(i32(0)))
}
func (m *Module) fn1664(v0, v1, v2, v3, v4 int32) {
	var v5 int32
	var v6, v7, v8 int64
	var v9, v10, v11, v12, v13, v14, v15, v16, v17, v18, v19, v20, v21, v22, v23, v24, v25, v26, v27, v28, v29, v30, v31 int32
	t0 := m.g0
	v5 = t0 - i32(832)
	m.g0 = v5
	{
		t1 := int64(load64(m.memory[uint32(v1):]))
		v6 = t1
		if v6 == i64(0) {
			m.fn256(i32(1130120), i32(28), i32(1130576))
			panic("unreachable")
		}
		t2 := int64(load64(m.memory[int64(uint32(v1))+8:]))
		v7 = t2
		if v7 == i64(0) {
			m.fn256(i32(1130244), i32(29), i32(1130592))
			panic("unreachable")
		}
		t3 := int64(load64(m.memory[int64(uint32(v1))+16:]))
		v8 = t3
		if v8 == i64(0) {
			m.fn256(i32(1130292), i32(28), i32(1130608))
			panic("unreachable")
		}
		if uint64(v8) > uint64(v6^i64(-1)) {
			m.fn256(i32(1130504), i32(54), i32(1130720))
			panic("unreachable")
		}
		if uint64(v6) < uint64(v7) {
			m.fn256(i32(1130432), i32(55), i32(1130704))
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
		_ = m.fn1650(v5+i32(8), v1)
		goto l6
	l5:
		_ = m.fn1650(v5+i32(176), int32(int16(i32(0)-v1)))
	l6:
		if v10 > i32(-1) {
			goto l7
		}
		_ = m.fn1665(v5+i32(8), (i32(0)-v10)&i32(0xffff))
		goto l8
	l7:
		_ = m.fn1665(v5+i32(176), v9&i32(0x7fff))
	l8:
		memory_copy(m.memory, uint32(v5+i32(668)), uint32(v5+i32(176)), uint32(i32(164)))
		v11 = v3
		if uint32(v3) < uint32(i32(10)) {
			goto l9
		}
		v12 = v5 + i32(668) + i32(-4)
		v11 = v3
	l13:
		{
			t11 := int32(load32(m.memory[int64(uint32(v5))+828:]))
			v1 = t11
			if uint32(v1) > uint32(i32(40)) {
				m.fn151(i32(0), v1, i32(40), i32(1108992))
				panic("unreachable")
			}
			if v1 == 0 {
				goto l11
			}
			v1 = v1 << 2
			v6 = i64(0)
		l12:
			{
				v9 = v12 + v1
				t12 := int64(load32(m.memory[uint32(v9):]))
				t13 := v9
				v6 = v6<<32 | t12
				t14 := int64(uint64(v6) / uint64(i64(1000000000)))
				v7 = t14
				store32(m.memory[uint32(t13):], uint32(v7))
				v6 = v6 - v7*i64(1000000000)
				v1 = v1 + i32(-4)
				if v1 != 0 {
					goto l12
				}
			}
		l11:
			v11 = v11 + i32(-9)
			if uint32(v11) > uint32(i32(9)) {
				goto l13
			}
		}
	l9:
		t15 := int32(load32(m.memory[int64(uint32(v11<<2))+1130864:]))
		v9 = t15 << 1
		if v9 == 0 {
			m.fn256(i32(1108936), i32(27), i32(1108992))
			panic("unreachable")
		}
		{
			{
				t16 := int32(load32(m.memory[int64(uint32(v5))+828:]))
				v1 = t16
				if uint32(v1) > uint32(i32(40)) {
					m.fn151(i32(0), v1, i32(40), i32(1108992))
					panic("unreachable")
				}
				if v1 != 0 {
					goto l16
				}
				v1 = i32(0)
				goto l17
			}
		l16:
			v1 = v1 << 2
			v12 = v5 + i32(668) + i32(-4)
			v6 = int64(uint32(v9))
			v7 = i64(0)
		l18:
			{
				v9 = v12 + v1
				t17 := int64(load32(m.memory[uint32(v9):]))
				t18 := v9
				v7 = v7<<32 | t17
				t19 := int64(uint64(v7) / uint64(v6))
				v8 = t19
				store32(m.memory[uint32(t18):], uint32(v8))
				v7 = v7 - v8*v6
				v1 = v1 + i32(-4)
				if v1 != 0 {
					goto l18
				}
			}
			t20 := int32(load32(m.memory[int64(uint32(v5))+828:]))
			v1 = t20
		}
	l17:
		{
			t21 := int32(load32(m.memory[int64(uint32(v5))+168:]))
			v13 = t21
			p22 := v1
			if uint32(v13) > uint32(v1) {
				p22 = v13
			}
			v14 = p22
			if uint32(v14) > uint32(i32(40)) {
				m.fn151(i32(0), v14, i32(40), i32(1108992))
				panic("unreachable")
			}
			if v14 != 0 {
				goto l20
			}
			v14 = i32(0)
			goto l21
		}
	l20:
		v11 = i32(0)
		v9 = v5 + i32(8)
		v1 = v5 + i32(668)
		v15 = v14
	l22:
		{
			t23 := int32(load32(m.memory[uint32(v9):]))
			t24 := v1
			v16 = t23
			t25 := int32(load32(m.memory[uint32(v1):]))
			v12 = v16 + t25
			v11 = v12 + v11&i32(1)
			store32(m.memory[uint32(t24):], uint32(v11))
			var p26 int32
			if uint32(v12) < uint32(v16) {
				p26 = 1
			}
			var p27 int32
			if uint32(v11) < uint32(v12) {
				p27 = 1
			}
			v11 = p26 | p27
			v1 = v1 + i32(4)
			v9 = v9 + i32(4)
			v15 = v15 + i32(-1)
			if v15 != 0 {
				goto l22
			}
		}
		if v11 == 0 {
			goto l21
		}
		if v14 == i32(40) {
			m.fn158(i32(40), i32(40), i32(1108992))
			panic("unreachable")
		}
		store32(m.memory[uint32(v5+i32(668)+v14<<2):], uint32(i32(1)))
		v14 = v14 + i32(1)
	l21:
		store32(m.memory[int64(uint32(v5))+828:], uint32(v14))
		t28 := int32(load32(m.memory[int64(uint32(v5))+336:]))
		v17 = t28
		p29 := v14
		if uint32(v17) > uint32(v14) {
			p29 = v17
		}
		v1 = p29
		if uint32(v1) >= uint32(i32(41)) {
			m.fn151(i32(0), v1, i32(40), i32(1108992))
			panic("unreachable")
		}
		v1 = v1 << 2
		v9 = v5 + i32(668) + i32(-4)
	l26:
		{
			if v1 == 0 {
				goto l25
			}
			t30 := int32(load32(m.memory[uint32(v9+v1):]))
			v12 = t30
			t31 := v12
			v1 = v1 + i32(-4)
			t32 := int32(load32(m.memory[uint32(v1+(v5+i32(176))):]))
			v11 = t32
			if t31 == v11 {
				goto l26
			}
		}
		if uint32(v12) >= uint32(v11) {
			goto l25
		}
		if v13 != 0 {
			t33 := v5 + i32(8)
			v9 = v13 << 2
			v12 = t33 + v9
			v6 = i64(0)
			v1 = v5 + i32(8)
		l29:
			{
				t34 := int64(load32(m.memory[uint32(v1):]))
				t35 := v1
				v6 = t34*i64(10) + v6
				store32(m.memory[uint32(t35):], uint32(v6))
				v1 = v1 + i32(4)
				v6 = int64(uint64(v6) >> 32)
				v9 = v9 + i32(-4)
				if v9 != 0 {
					goto l29
				}
			}
			if v6 == 0 {
				goto l30
			}
			if v13 == i32(40) {
				m.fn158(i32(40), i32(40), i32(1108992))
				panic("unreachable")
			}
			store32(m.memory[uint32(v12):], uint32(int32(v6)))
			v13 = v13 + i32(1)
		l30:
			store32(m.memory[int64(uint32(v5))+168:], uint32(v13))
			goto l28
		}
		v13 = i32(0)
		store32(m.memory[int64(uint32(v5))+168:], uint32(i32(0)))
		goto l28
	l25:
		v10 = v10 + i32(1)
	l28:
		v18 = i32(0)
		v16 = i32(1)
		{
			v1 = int32(int16(v10))
			t36 := v1
			v9 = int32(int16(v4))
			var p37 int32
			if t36 < v9 {
				p37 = 1
			}
			v19 = p37
			if v19 != 0 {
				goto l32
			}
			p38 := v3
			if uint32(v1-v9) < uint32(v3) {
				p38 = int32(int16(v10 - v4))
			}
			v20 = p38
			if v20 != 0 {
				memory_copy(m.memory, uint32(v5+i32(340)), uint32(v5+i32(176)), uint32(i32(164)))
				t39 := m.fn1650(v5+i32(340), i32(1))
				v1 = t39
				memory_copy(m.memory, uint32(v5+i32(504)), uint32(v5+i32(176)), uint32(i32(164)))
				t40 := m.fn1650(v5+i32(504), i32(2))
				v9 = t40
				memory_copy(m.memory, uint32(v5+i32(668)), uint32(v5+i32(176)), uint32(i32(164)))
				v21 = v5 + i32(8) + i32(-4)
				v22 = v5 + i32(176) + i32(-4)
				v23 = v5 + i32(340) + i32(-4)
				v4 = v5 + i32(504) + i32(-4)
				v14 = v5 + i32(668) + i32(-4)
				t41 := m.fn1650(v5+i32(668), i32(3))
				v12 = t41
				t42 := int32(load32(m.memory[int64(uint32(v1))+160:]))
				v24 = t42
				t43 := int32(load32(m.memory[int64(uint32(v9))+160:]))
				v25 = t43
				t44 := int32(load32(m.memory[int64(uint32(v12))+160:]))
				v26 = t44
				v27 = i32(0)
			l70:
				{
					v28 = v27
					if uint32(v13) >= uint32(i32(41)) {
						m.fn151(i32(0), v13, i32(40), i32(1108992))
						panic("unreachable")
					}
					v27 = v28 + i32(1)
					v12 = v13 << 2
					v1 = i32(0)
				l37:
					{
						if v12 == v1 {
							if uint32(v20) > uint32(v3) {
								m.fn151(v28, v20, v3, i32(1130688))
								panic("unreachable")
							}
							if v20 == v28 {
								goto l72
							}
							v1 = v20 - v28
							if v1 == 0 {
								goto l72
							}
							memory_fill(m.memory, uint32(v2+v28), i32(48), uint32(v1))
						l72:
							store16(m.memory[int64(uint32(v0))+8:], uint16(v10))
							store32(m.memory[int64(uint32(v0))+4:], uint32(v20))
							goto l73
						}
						v9 = v5 + i32(8) + v1
						v1 = v1 + i32(4)
						t45 := int32(load32(m.memory[uint32(v9):]))
						if t45 == 0 {
							goto l37
						}
					}
					p46 := v13
					if uint32(v26) > uint32(v13) {
						p46 = v26
					}
					v29 = p46
					if uint32(v29) >= uint32(i32(41)) {
						m.fn151(i32(0), v29, i32(40), i32(1108992))
						panic("unreachable")
					}
					v1 = v29 << 2
				l40:
					{
						if v1 == 0 {
							goto l39
						}
						v9 = v14 + v1
						v1 = v1 + i32(-4)
						t47 := int32(load32(m.memory[uint32(v1+(v5+i32(8))):]))
						v12 = t47
						t48 := int32(load32(m.memory[uint32(v9):]))
						t49 := v12
						v9 = t48
						if t49 == v9 {
							goto l40
						}
					}
					v30 = i32(0)
					if uint32(v12) < uint32(v9) {
						goto l41
					}
				l39:
					v11 = i32(1)
					v9 = v5 + i32(668)
					v1 = v5 + i32(8)
					v15 = v29
				l42:
					{
						t50 := int32(load32(m.memory[uint32(v1):]))
						t51 := v1
						v16 = t50
						t52 := int32(load32(m.memory[uint32(v9):]))
						v12 = v16 + (t52 ^ i32(-1))
						v11 = v12 + v11&i32(1)
						store32(m.memory[uint32(t51):], uint32(v11))
						var p53 int32
						if uint32(v12) < uint32(v16) {
							p53 = 1
						}
						var p54 int32
						if uint32(v11) < uint32(v12) {
							p54 = 1
						}
						v11 = p53 | p54
						v1 = v1 + i32(4)
						v9 = v9 + i32(4)
						v15 = v15 + i32(-1)
						if v15 != 0 {
							goto l42
						}
					}
					if v11 == 0 {
						m.fn256(i32(1108963), i32(26), i32(1108992))
						panic("unreachable")
					}
					store32(m.memory[int64(uint32(v5))+168:], uint32(v29))
					v30 = i32(8)
					v13 = v29
				l41:
					p55 := v13
					if uint32(v25) > uint32(v13) {
						p55 = v25
					}
					v29 = p55
					if uint32(v29) >= uint32(i32(41)) {
						m.fn151(i32(0), v29, i32(40), i32(1108992))
						panic("unreachable")
					}
					v1 = v29 << 2
				l46:
					{
						if v1 == 0 {
							goto l45
						}
						v9 = v4 + v1
						v1 = v1 + i32(-4)
						t56 := int32(load32(m.memory[uint32(v1+(v5+i32(8))):]))
						v12 = t56
						t57 := int32(load32(m.memory[uint32(v9):]))
						t58 := v12
						v9 = t57
						if t58 == v9 {
							goto l46
						}
					}
					if uint32(v12) >= uint32(v9) {
						goto l45
					}
					v29 = v13
					goto l47
				l45:
					if v29 == 0 {
						goto l48
					}
					v11 = i32(1)
					v9 = v5 + i32(504)
					v1 = v5 + i32(8)
					v15 = v29
				l49:
					{
						t59 := int32(load32(m.memory[uint32(v1):]))
						t60 := v1
						v16 = t59
						t61 := int32(load32(m.memory[uint32(v9):]))
						v12 = v16 + (t61 ^ i32(-1))
						v11 = v12 + v11&i32(1)
						store32(m.memory[uint32(t60):], uint32(v11))
						var p62 int32
						if uint32(v12) < uint32(v16) {
							p62 = 1
						}
						var p63 int32
						if uint32(v11) < uint32(v12) {
							p63 = 1
						}
						v11 = p62 | p63
						v1 = v1 + i32(4)
						v9 = v9 + i32(4)
						v15 = v15 + i32(-1)
						if v15 != 0 {
							goto l49
						}
					}
					if v11 == 0 {
						m.fn256(i32(1108963), i32(26), i32(1108992))
						panic("unreachable")
					}
				l48:
					store32(m.memory[int64(uint32(v5))+168:], uint32(v29))
					v30 = v30 | i32(4)
				l47:
					p64 := v29
					if uint32(v24) > uint32(v29) {
						p64 = v24
					}
					v31 = p64
					if uint32(v31) >= uint32(i32(41)) {
						m.fn151(i32(0), v31, i32(40), i32(1108992))
						panic("unreachable")
					}
					v1 = v31 << 2
				l53:
					{
						if v1 == 0 {
							goto l52
						}
						v9 = v23 + v1
						v1 = v1 + i32(-4)
						t65 := int32(load32(m.memory[uint32(v1+(v5+i32(8))):]))
						v12 = t65
						t66 := int32(load32(m.memory[uint32(v9):]))
						t67 := v12
						v9 = t66
						if t67 == v9 {
							goto l53
						}
					}
					if uint32(v12) >= uint32(v9) {
						goto l52
					}
					v31 = v29
					goto l54
				l52:
					if v31 == 0 {
						goto l55
					}
					v11 = i32(1)
					v9 = v5 + i32(340)
					v1 = v5 + i32(8)
					v15 = v31
				l56:
					{
						t68 := int32(load32(m.memory[uint32(v1):]))
						t69 := v1
						v16 = t68
						t70 := int32(load32(m.memory[uint32(v9):]))
						v12 = v16 + (t70 ^ i32(-1))
						v11 = v12 + v11&i32(1)
						store32(m.memory[uint32(t69):], uint32(v11))
						var p71 int32
						if uint32(v12) < uint32(v16) {
							p71 = 1
						}
						var p72 int32
						if uint32(v11) < uint32(v12) {
							p72 = 1
						}
						v11 = p71 | p72
						v1 = v1 + i32(4)
						v9 = v9 + i32(4)
						v15 = v15 + i32(-1)
						if v15 != 0 {
							goto l56
						}
					}
					if v11 == 0 {
						m.fn256(i32(1108963), i32(26), i32(1108992))
						panic("unreachable")
					}
				l55:
					store32(m.memory[int64(uint32(v5))+168:], uint32(v31))
					v30 = v30 + i32(2)
				l54:
					p73 := v31
					if uint32(v17) > uint32(v31) {
						p73 = v17
					}
					v13 = p73
					if uint32(v13) >= uint32(i32(41)) {
						m.fn151(i32(0), v13, i32(40), i32(1108992))
						panic("unreachable")
					}
					v1 = v13 << 2
				l60:
					{
						if v1 == 0 {
							goto l59
						}
						v9 = v22 + v1
						v12 = v21 + v1
						v1 = v1 + i32(-4)
						t74 := int32(load32(m.memory[uint32(v12):]))
						v12 = t74
						t75 := int32(load32(m.memory[uint32(v9):]))
						t76 := v12
						v9 = t75
						if t76 == v9 {
							goto l60
						}
					}
					if uint32(v12) >= uint32(v9) {
						goto l59
					}
					v13 = v31
					goto l61
				l59:
					if v13 == 0 {
						goto l62
					}
					v11 = i32(1)
					v9 = v5 + i32(176)
					v1 = v5 + i32(8)
					v15 = v13
				l63:
					{
						t77 := int32(load32(m.memory[uint32(v1):]))
						t78 := v1
						v16 = t77
						t79 := int32(load32(m.memory[uint32(v9):]))
						v12 = v16 + (t79 ^ i32(-1))
						v11 = v12 + v11&i32(1)
						store32(m.memory[uint32(t78):], uint32(v11))
						var p80 int32
						if uint32(v12) < uint32(v16) {
							p80 = 1
						}
						var p81 int32
						if uint32(v11) < uint32(v12) {
							p81 = 1
						}
						v11 = p80 | p81
						v1 = v1 + i32(4)
						v9 = v9 + i32(4)
						v15 = v15 + i32(-1)
						if v15 != 0 {
							goto l63
						}
					}
					if v11 == 0 {
						m.fn256(i32(1108963), i32(26), i32(1108992))
						panic("unreachable")
					}
				l62:
					store32(m.memory[int64(uint32(v5))+168:], uint32(v13))
					v30 = v30 + i32(1)
				l61:
					if v28 == v3 {
						m.fn158(v3, v3, i32(1130672))
						panic("unreachable")
					}
					m.memory[uint32(v2+v28)] = byte(v30 + i32(48))
					{
						if v13 != 0 {
							goto l66
						}
						v13 = i32(0)
						goto l67
					l66:
						t82 := v5 + i32(8)
						v9 = v13 << 2
						v12 = t82 + v9
						v6 = i64(0)
						v1 = v5 + i32(8)
					l68:
						{
							t83 := int64(load32(m.memory[uint32(v1):]))
							t84 := v1
							v6 = t83*i64(10) + v6
							store32(m.memory[uint32(t84):], uint32(v6))
							v1 = v1 + i32(4)
							v6 = int64(uint64(v6) >> 32)
							v9 = v9 + i32(-4)
							if v9 != 0 {
								goto l68
							}
						}
						if v6 == 0 {
							goto l67
						}
						if v13 == i32(40) {
							m.fn158(i32(40), i32(40), i32(1108992))
							panic("unreachable")
						}
						store32(m.memory[uint32(v12):], uint32(int32(v6)))
						v13 = v13 + i32(1)
					}
				l67:
					store32(m.memory[int64(uint32(v5))+168:], uint32(v13))
					if v27 != v20 {
						goto l70
					}
				}
				v16 = i32(0)
				goto l34
			}
		}
	l32:
		v20 = i32(0)
		goto l34
	}
l34:
	{
		{
			if v17 == 0 {
				goto l74
			}
			t85 := v5 + i32(176)
			v9 = v17 << 2
			v12 = t85 + v9
			v6 = i64(0)
			v1 = v5 + i32(176)
		l75:
			{
				t86 := int64(load32(m.memory[uint32(v1):]))
				t87 := v1
				v6 = t86*i64(5) + v6
				store32(m.memory[uint32(t87):], uint32(v6))
				v1 = v1 + i32(4)
				v6 = int64(uint64(v6) >> 32)
				v9 = v9 + i32(-4)
				if v9 != 0 {
					goto l75
				}
			}
			if !(v6 == 0) {
				goto l76
			}
			v18 = v17
			goto l74
		l76:
			if v17 == i32(40) {
				m.fn158(i32(40), i32(40), i32(1108992))
				panic("unreachable")
			}
			store32(m.memory[uint32(v12):], uint32(int32(v6)))
			v18 = v17 + i32(1)
		}
	l74:
		store32(m.memory[int64(uint32(v5))+336:], uint32(v18))
		p88 := v13
		if uint32(v18) > uint32(v13) {
			p88 = v18
		}
		v1 = p88
		if uint32(v1) >= uint32(i32(41)) {
			m.fn151(i32(0), v1, i32(40), i32(1108992))
			panic("unreachable")
		}
		v1 = v1 << 2
		v11 = v5 + i32(8) + i32(-4)
		v15 = v5 + i32(176) + i32(-4)
		{
			{
			l80:
				{
					if v1 == 0 {
						goto l79
					}
					v9 = v15 + v1
					v12 = v11 + v1
					v1 = v1 + i32(-4)
					t89 := int32(load32(m.memory[uint32(v12):]))
					v12 = t89
					t90 := int32(load32(m.memory[uint32(v9):]))
					t91 := v12
					v9 = t90
					if t91 == v9 {
						goto l80
					}
				}
				var p92 int32
				if uint32(v12) > uint32(v9) {
					p92 = 1
				}
				var p93 int32
				if uint32(v12) < uint32(v9) {
					p93 = 1
				}
				switch (p92 - p93) & i32(255) {
				case 0:
					goto l79
				case 1:
					goto l81
				default:
					goto l82
				}
			}
		l79:
			v1 = i32(0)
			if v16 != 0 {
				goto l83
			}
			v1 = v20 + i32(-1)
			if uint32(v1) >= uint32(v3) {
				m.fn158(v1, v3, i32(1130624))
				panic("unreachable")
			}
			t94 := int32(m.memory[uint32(v2+v1)])
			if t94&i32(1) == 0 {
				goto l82
			}
		}
	l81:
		if uint32(v20) > uint32(v3) {
			m.fn151(i32(0), v20, v3, i32(1130640))
			panic("unreachable")
		}
		v11 = v2 + v20
		v1 = v20
		{
		l87:
			{
				v9 = v1
				if v9 == 0 {
					v1 = i32(49)
					if v16 != 0 {
						goto l88
					}
					m.memory[uint32(v2)] = byte(i32(49))
					v1 = i32(48)
					v9 = v20 + i32(-1)
					if v9 == 0 {
						goto l88
					}
					memory_fill(m.memory, uint32(v2+i32(1)), i32(48), uint32(v9))
				l88:
					v10 = v10 + i32(1)
					if v19 != 0 {
						goto l82
					}
					if uint32(v20) >= uint32(v3) {
						goto l82
					}
					m.memory[uint32(v11)] = byte(v1)
					v20 = v20 + i32(1)
					goto l82
				}
				v1 = v9 + i32(-1)
				v12 = v1 + v2
				t95 := int32(m.memory[uint32(v12)])
				if t95 == i32(57) {
					goto l87
				}
			}
			t96 := int32(m.memory[uint32(v12)])
			m.memory[uint32(v12)] = byte(t96 + i32(1))
			v1 = v20 - v9
			if v1 == 0 {
				goto l82
			}
			memory_fill(m.memory, uint32(v2+v9), i32(48), uint32(v1))
			goto l82
		}
	}
l82:
	if uint32(v20) > uint32(v3) {
		m.fn151(i32(0), v20, v3, i32(1130656))
		panic("unreachable")
	}
	v1 = v20
l83:
	store16(m.memory[int64(uint32(v0))+8:], uint16(v10))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
	goto l73
l73:
	store32(m.memory[uint32(v0):], uint32(v2))
	m.g0 = v5 + i32(832)
}
func (m *Module) fn1665(v0, v1 int32) int32 {
	var v2, v3, v4, v5 int32
	var v6, v7 int64
	if uint32(v1) < uint32(i32(8)) {
		t5 := int32(load32(m.memory[int64(uint32(v0))+160:]))
		v3 = t5
		if uint32(v3) > uint32(i32(40)) {
			m.fn151(i32(0), v3, i32(40), i32(1108992))
			panic("unreachable")
		}
		if v3 != 0 {
			t6 := int64(load32(m.memory[int64(uint32(v1<<2))+1130864:]))
			v6 = t6
			t7 := v0
			v4 = v3 << 2
			v1 = t7 + v4
			v7 = i64(0)
			v2 = v0
		l9:
			{
				t8 := int64(load32(m.memory[uint32(v2):]))
				t9 := v2
				v7 = t8*v6 + v7
				store32(m.memory[uint32(t9):], uint32(v7))
				v2 = v2 + i32(4)
				v7 = int64(uint64(v7) >> 32)
				v4 = v4 + i32(-4)
				if v4 != 0 {
					goto l9
				}
			}
			if v7 == 0 {
				goto l10
			}
			if v3 == i32(40) {
				m.fn158(i32(40), i32(40), i32(1108992))
				panic("unreachable")
			}
			store32(m.memory[uint32(v1):], uint32(int32(v7)))
			v3 = v3 + i32(1)
		l10:
			store32(m.memory[int64(uint32(v0))+160:], uint32(v3))
			return v0
		}
		store32(m.memory[int64(uint32(v0))+160:], uint32(i32(0)))
		return v0
	}
	v2 = v1 & i32(7)
	if v2 == 0 {
		goto l1
	}
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+160:]))
		v3 = t0
		if uint32(v3) > uint32(i32(40)) {
			m.fn151(i32(0), v3, i32(40), i32(1108992))
			panic("unreachable")
		}
		if v3 != 0 {
			t1 := v0
			v4 = v3 << 2
			v5 = t1 + v4
			t2 := int32(load32(m.memory[int64(uint32(v2<<2))+1130864:]))
			v6 = int64(uint32(i32_shr_u(t2, v2)))
			v7 = i64(0)
			v2 = v0
		l4:
			{
				t3 := int64(load32(m.memory[uint32(v2):]))
				t4 := v2
				v7 = t3*v6 + v7
				store32(m.memory[uint32(t4):], uint32(v7))
				v2 = v2 + i32(4)
				v7 = int64(uint64(v7) >> 32)
				v4 = v4 + i32(-4)
				if v4 != 0 {
					goto l4
				}
			}
			if v7 == 0 {
				goto l5
			}
			if v3 == i32(40) {
				m.fn158(i32(40), i32(40), i32(1108992))
				panic("unreachable")
			}
			store32(m.memory[uint32(v5):], uint32(int32(v7)))
			v3 = v3 + i32(1)
		l5:
			store32(m.memory[int64(uint32(v0))+160:], uint32(v3))
			goto l1
		}
		store32(m.memory[int64(uint32(v0))+160:], uint32(i32(0)))
		goto l1
	}
l1:
	if v1&i32(8) == 0 {
		goto l12
	}
	{
		{
			t10 := int32(load32(m.memory[int64(uint32(v0))+160:]))
			v3 = t10
			if uint32(v3) > uint32(i32(40)) {
				m.fn151(i32(0), v3, i32(40), i32(1108992))
				panic("unreachable")
			}
			if v3 != 0 {
				goto l14
			}
			v3 = i32(0)
			goto l15
		}
	l14:
		t11 := v0
		v4 = v3 << 2
		v5 = t11 + v4
		v7 = i64(0)
		v2 = v0
	l16:
		{
			t12 := int64(load32(m.memory[uint32(v2):]))
			t13 := v2
			v7 = t12*i64(390625) + v7
			store32(m.memory[uint32(t13):], uint32(v7))
			v2 = v2 + i32(4)
			v7 = int64(uint64(v7) >> 32)
			v4 = v4 + i32(-4)
			if v4 != 0 {
				goto l16
			}
		}
		if v7 == 0 {
			goto l15
		}
		if v3 == i32(40) {
			m.fn158(i32(40), i32(40), i32(1108992))
			panic("unreachable")
		}
		store32(m.memory[uint32(v5):], uint32(int32(v7)))
		v3 = v3 + i32(1)
	}
l15:
	store32(m.memory[int64(uint32(v0))+160:], uint32(v3))
l12:
	if v1&i32(16) == 0 {
		goto l18
	}
	_ = m.fn1649(v0, i32(1130904), i32(2))
l18:
	if v1&i32(32) == 0 {
		goto l19
	}
	_ = m.fn1649(v0, i32(1130912), i32(3))
l19:
	if v1&i32(64) == 0 {
		goto l20
	}
	_ = m.fn1649(v0, i32(1130924), i32(5))
l20:
	if v1&i32(128) == 0 {
		goto l21
	}
	_ = m.fn1649(v0, i32(1130944), i32(10))
l21:
	if v1&i32(256) == 0 {
		goto l22
	}
	_ = m.fn1649(v0, i32(1130984), i32(19))
l22:
	_ = m.fn1650(v0, v1)
	return v0
}
func (m *Module) fn1666(v0, v1, v2 int32) {
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
			m.fn256(i32(1130120), i32(28), i32(1130228))
			panic("unreachable")
		}
		t2 := int64(load64(m.memory[int64(uint32(v1))+8:]))
		v5 = t2
		if v5 == i64(0) {
			m.fn256(i32(1130244), i32(29), i32(1130276))
			panic("unreachable")
		}
		t3 := int64(load64(m.memory[int64(uint32(v1))+16:]))
		v6 = t3
		if v6 == i64(0) {
			m.fn256(i32(1130292), i32(28), i32(1130320))
			panic("unreachable")
		}
		v7 = v6 + v4
		if uint64(v7) < uint64(v6) {
			m.fn256(i32(1130504), i32(54), i32(1130560))
			panic("unreachable")
		}
		if uint64(v4) < uint64(v5) {
			m.fn256(i32(1130432), i32(55), i32(1130488))
			panic("unreachable")
		}
		if uint64(v7) >= uint64(i64(0x2000000000000000)) {
			m.fn256(i32(1130336), i32(45), i32(1130384))
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
			m.fn1643(v3+i32(72), v3+i32(56))
			panic("unreachable")
		}
		store16(m.memory[int64(uint32(v3))+64:], uint16(v1))
		store64(m.memory[int64(uint32(v3))+56:], uint64(v4))
		t9 := v3
		v9 = i64_shl(v4, v6)
		v5 = i64_shr_u(v9, v6)
		store64(m.memory[int64(uint32(t9))+72:], uint64(v5))
		if v5 != v4 {
			m.fn1643(v3+i32(72), v3+i32(56))
			panic("unreachable")
		}
		v10 = v1 - int32(v6)
		v1 = (int32(int16(i32(-96)-v10))*i32(80) + i32(86960)) / i32(2126)
		if uint32(v1) > uint32(i32(80)) {
			m.fn158(v1, i32(81), i32(1130164))
			panic("unreachable")
		}
		t10 := v3 + i32(32)
		v1 = v1 << 4
		t11 := int64(load64(m.memory[int64(uint32(v1))+1128824:]))
		v4 = t11
		m.fn1853(t10, v4, i64(0), i64_shl(v7, v6), i64(0))
		m.fn1853(v3+i32(16), v4, i64(0), v8, i64(0))
		m.fn1853(v3, v4, i64(0), v9, i64(0))
		t12 := int32(load16(m.memory[int64(uint32(v1))+1128832:]))
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
		t16 := int32(load16(m.memory[int64(uint32(v1))+1128834:]))
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
					m.fn158(i32(17), i32(17), i32(1130416))
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
	m.fn494(i32(1130400))
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
func (m *Module) fn1667(v0, v1, v2 int32) {
	var v3 int32
	var v4, v5, v6, v7 int64
	var v8, v9, v10, v11, v12, v13, v14, v15, v16, v17, v18, v19, v20, v21, v22, v23, v24, v25, v26, v27, v28, v29, v30, v31 int32
	t0 := m.g0
	v3 = t0 - i32(1328)
	m.g0 = v3
	{
		t1 := int64(load64(m.memory[uint32(v1):]))
		v4 = t1
		if v4 == i64(0) {
			m.fn256(i32(1130120), i32(28), i32(1130736))
			panic("unreachable")
		}
		t2 := int64(load64(m.memory[int64(uint32(v1))+8:]))
		v5 = t2
		if v5 == i64(0) {
			m.fn256(i32(1130244), i32(29), i32(1130752))
			panic("unreachable")
		}
		t3 := int64(load64(m.memory[int64(uint32(v1))+16:]))
		v6 = t3
		if v6 == i64(0) {
			m.fn256(i32(1130292), i32(28), i32(1130768))
			panic("unreachable")
		}
		v7 = v6 + v4
		if uint64(v7) < uint64(v6) {
			m.fn256(i32(1130504), i32(54), i32(1130848))
			panic("unreachable")
		}
		if uint64(v4) < uint64(v5) {
			m.fn256(i32(1130432), i32(55), i32(1130832))
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
		_ = m.fn1650(v3+i32(8), v1)
		_ = m.fn1650(v3+i32(176), v1)
		_ = m.fn1650(v3+i32(344), v1)
		goto l6
	l5:
		_ = m.fn1650(v3+i32(508), int32(int16(i32(0)-v1)))
	l6:
		{
			if v10 > i32(-1) {
				goto l7
			}
			t16 := v3 + i32(8)
			v1 = (i32(0) - v10) & i32(0xffff)
			_ = m.fn1665(t16, v1)
			_ = m.fn1665(v3+i32(176), v1)
			_ = m.fn1665(v3+i32(344), v1)
			goto l8
		}
	l7:
		_ = m.fn1665(v3+i32(508), v9&i32(0x7fff))
	l8:
		memory_copy(m.memory, uint32(v3+i32(1164)), uint32(v3+i32(8)), uint32(i32(164)))
		{
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
				m.fn151(i32(0), v12, i32(40), i32(1108992))
				panic("unreachable")
			}
			if v12 != 0 {
				goto l10
			}
			v12 = i32(0)
			goto l11
		}
	l10:
		v13 = i32(0)
		v9 = v3 + i32(344)
		v1 = v3 + i32(1164)
		v14 = v12
	l12:
		{
			t25 := int32(load32(m.memory[uint32(v9):]))
			t26 := v1
			v15 = t25
			t27 := int32(load32(m.memory[uint32(v1):]))
			v16 = v15 + t27
			v13 = v16 + v13&i32(1)
			store32(m.memory[uint32(t26):], uint32(v13))
			var p28 int32
			if uint32(v16) < uint32(v15) {
				p28 = 1
			}
			var p29 int32
			if uint32(v13) < uint32(v16) {
				p29 = 1
			}
			v13 = p28 | p29
			v1 = v1 + i32(4)
			v9 = v9 + i32(4)
			v14 = v14 + i32(-1)
			if v14 != 0 {
				goto l12
			}
		}
		if v13 == 0 {
			goto l11
		}
		if v12 == i32(40) {
			m.fn158(i32(40), i32(40), i32(1108992))
			panic("unreachable")
		}
		store32(m.memory[uint32(v3+i32(1164)+v12<<2):], uint32(i32(1)))
		v12 = v12 + i32(1)
	l11:
		store32(m.memory[int64(uint32(v3))+1324:], uint32(v12))
		t30 := int32(load32(m.memory[int64(uint32(v3))+668:]))
		t31 := v12
		v17 = t30
		p32 := v17
		if uint32(v12) > uint32(v17) {
			p32 = t31
		}
		v1 = p32
		if uint32(v1) >= uint32(i32(41)) {
			m.fn151(i32(0), v1, i32(40), i32(1108992))
			panic("unreachable")
		}
		v1 = v1 << 2
		v16 = v3 + i32(1164) + i32(-4)
		{
		l17:
			{
				if v1 != 0 {
					goto l15
				}
				v1 = i32(0)
				goto l16
			l15:
				v9 = v16 + v1
				v1 = v1 + i32(-4)
				t33 := int32(load32(m.memory[uint32(v1+(v3+i32(508))):]))
				v13 = t33
				t34 := int32(load32(m.memory[uint32(v9):]))
				t35 := v13
				v9 = t34
				if t35 == v9 {
					goto l17
				}
			}
			var p36 int32
			if uint32(v13) > uint32(v9) {
				p36 = 1
			}
			var p37 int32
			if uint32(v13) < uint32(v9) {
				p37 = 1
			}
			v1 = p36 - p37
		}
	l16:
		{
			{
				if v1 < v8 {
					goto l18
				}
				{
					{
						t38 := int32(load32(m.memory[int64(uint32(v3))+168:]))
						v13 = t38
						if uint32(v13) > uint32(i32(40)) {
							m.fn151(i32(0), v13, i32(40), i32(1108992))
							panic("unreachable")
						}
						if v13 != 0 {
							goto l20
						}
						v13 = i32(0)
						goto l21
					}
				l20:
					t39 := v3 + i32(8)
					v9 = v13 << 2
					v16 = t39 + v9
					v4 = i64(0)
					v1 = v3 + i32(8)
				l22:
					{
						t40 := int64(load32(m.memory[uint32(v1):]))
						t41 := v1
						v4 = t40*i64(10) + v4
						store32(m.memory[uint32(t41):], uint32(v4))
						v1 = v1 + i32(4)
						v4 = int64(uint64(v4) >> 32)
						v9 = v9 + i32(-4)
						if v9 != 0 {
							goto l22
						}
					}
					if v4 == 0 {
						goto l21
					}
					if v13 == i32(40) {
						m.fn158(i32(40), i32(40), i32(1108992))
						panic("unreachable")
					}
					store32(m.memory[uint32(v16):], uint32(int32(v4)))
					v13 = v13 + i32(1)
				}
			l21:
				store32(m.memory[int64(uint32(v3))+168:], uint32(v13))
				t42 := int32(load32(m.memory[int64(uint32(v3))+336:]))
				v16 = t42
				if uint32(v16) > uint32(i32(40)) {
					m.fn151(i32(0), v16, i32(40), i32(1108992))
					panic("unreachable")
				}
				v18 = i32(0)
				v1 = i32(0)
				{
					if v16 == 0 {
						goto l25
					}
					t43 := v3 + i32(176)
					v9 = v16 << 2
					v14 = t43 + v9
					v4 = i64(0)
					v1 = v3 + i32(176)
				l26:
					{
						t44 := int64(load32(m.memory[uint32(v1):]))
						t45 := v1
						v4 = t44*i64(10) + v4
						store32(m.memory[uint32(t45):], uint32(v4))
						v1 = v1 + i32(4)
						v4 = int64(uint64(v4) >> 32)
						v9 = v9 + i32(-4)
						if v9 != 0 {
							goto l26
						}
					}
					if !(v4 == 0) {
						goto l27
					}
					v1 = v16
					goto l25
				l27:
					if v16 == i32(40) {
						m.fn158(i32(40), i32(40), i32(1108992))
						panic("unreachable")
					}
					store32(m.memory[uint32(v14):], uint32(int32(v4)))
					v1 = v16 + i32(1)
				}
			l25:
				store32(m.memory[int64(uint32(v3))+336:], uint32(v1))
				{
					if v11 == 0 {
						goto l29
					}
					t46 := v3 + i32(344)
					v9 = v11 << 2
					v16 = t46 + v9
					v4 = i64(0)
					v1 = v3 + i32(344)
				l30:
					{
						t47 := int64(load32(m.memory[uint32(v1):]))
						t48 := v1
						v4 = t47*i64(10) + v4
						store32(m.memory[uint32(t48):], uint32(v4))
						v1 = v1 + i32(4)
						v4 = int64(uint64(v4) >> 32)
						v9 = v9 + i32(-4)
						if v9 != 0 {
							goto l30
						}
					}
					{
						if !(v4 == 0) {
							goto l31
						}
						t49 := v3
						v18 = v11
						store32(m.memory[int64(uint32(t49))+504:], uint32(v18))
						goto l32
					}
				l31:
					if v11 == i32(40) {
						m.fn158(i32(40), i32(40), i32(1108992))
						panic("unreachable")
					}
					store32(m.memory[uint32(v16):], uint32(int32(v4)))
					v18 = v11 + i32(1)
				}
			l29:
				store32(m.memory[int64(uint32(v3))+504:], uint32(v18))
				goto l32
			}
		l18:
			v10 = v10 + i32(1)
			t50 := int32(load32(m.memory[int64(uint32(v3))+168:]))
			v13 = t50
			v18 = v11
		}
	l32:
		memory_copy(m.memory, uint32(v3+i32(672)), uint32(v3+i32(508)), uint32(i32(164)))
		t51 := m.fn1650(v3+i32(672), i32(1))
		v1 = t51
		memory_copy(m.memory, uint32(v3+i32(836)), uint32(v3+i32(508)), uint32(i32(164)))
		t52 := m.fn1650(v3+i32(836), i32(2))
		v9 = t52
		memory_copy(m.memory, uint32(v3+i32(1000)), uint32(v3+i32(508)), uint32(i32(164)))
		{
			{
				t53 := m.fn1650(v3+i32(1000), i32(3))
				t54 := int32(load32(m.memory[int64(uint32(t53))+160:]))
				v19 = t54
				p55 := v13
				if uint32(v19) > uint32(v13) {
					p55 = v19
				}
				v20 = p55
				if uint32(v20) > uint32(i32(40)) {
					goto l34
				}
				v21 = v3 + i32(508) + i32(-4)
				v22 = v3 + i32(1164) + i32(-4)
				v23 = v3 + i32(176) + i32(-4)
				v24 = v3 + i32(672) + i32(-4)
				v11 = v3 + i32(836) + i32(-4)
				v12 = v3 + i32(1000) + i32(-4)
				t56 := int32(load32(m.memory[int64(uint32(v1))+160:]))
				v25 = t56
				t57 := int32(load32(m.memory[int64(uint32(v9))+160:]))
				v26 = t57
				v27 = i32(0)
			l90:
				{
					v28 = v27
					v1 = v20 << 2
				l36:
					{
						if v1 == 0 {
							goto l35
						}
						v9 = v12 + v1
						v1 = v1 + i32(-4)
						t58 := int32(load32(m.memory[uint32(v1+(v3+i32(8))):]))
						v16 = t58
						t59 := int32(load32(m.memory[uint32(v9):]))
						t60 := v16
						v9 = t59
						if t60 == v9 {
							goto l36
						}
					}
					v29 = i32(0)
					if uint32(v16) < uint32(v9) {
						goto l37
					}
				l35:
					if v20 == 0 {
						goto l38
					}
					v13 = i32(1)
					v9 = v3 + i32(1000)
					v1 = v3 + i32(8)
					v14 = v20
				l39:
					{
						t61 := int32(load32(m.memory[uint32(v1):]))
						t62 := v1
						v15 = t61
						t63 := int32(load32(m.memory[uint32(v9):]))
						v16 = v15 + (t63 ^ i32(-1))
						v13 = v16 + v13&i32(1)
						store32(m.memory[uint32(t62):], uint32(v13))
						var p64 int32
						if uint32(v16) < uint32(v15) {
							p64 = 1
						}
						var p65 int32
						if uint32(v13) < uint32(v16) {
							p65 = 1
						}
						v13 = p64 | p65
						v1 = v1 + i32(4)
						v9 = v9 + i32(4)
						v14 = v14 + i32(-1)
						if v14 != 0 {
							goto l39
						}
					}
					if v13 == 0 {
						m.fn256(i32(1108963), i32(26), i32(1108992))
						panic("unreachable")
					}
				l38:
					store32(m.memory[int64(uint32(v3))+168:], uint32(v20))
					v29 = i32(8)
					v13 = v20
				l37:
					p66 := v13
					if uint32(v26) > uint32(v13) {
						p66 = v26
					}
					v20 = p66
					if uint32(v20) >= uint32(i32(41)) {
						m.fn151(i32(0), v20, i32(40), i32(1108992))
						panic("unreachable")
					}
					v1 = v20 << 2
				l43:
					{
						if v1 == 0 {
							goto l42
						}
						v9 = v11 + v1
						v1 = v1 + i32(-4)
						t67 := int32(load32(m.memory[uint32(v1+(v3+i32(8))):]))
						v16 = t67
						t68 := int32(load32(m.memory[uint32(v9):]))
						t69 := v16
						v9 = t68
						if t69 == v9 {
							goto l43
						}
					}
					if uint32(v16) >= uint32(v9) {
						goto l42
					}
					v20 = v13
					goto l44
				l42:
					if v20 == 0 {
						goto l45
					}
					v13 = i32(1)
					v9 = v3 + i32(836)
					v1 = v3 + i32(8)
					v14 = v20
				l46:
					{
						t70 := int32(load32(m.memory[uint32(v1):]))
						t71 := v1
						v15 = t70
						t72 := int32(load32(m.memory[uint32(v9):]))
						v16 = v15 + (t72 ^ i32(-1))
						v13 = v16 + v13&i32(1)
						store32(m.memory[uint32(t71):], uint32(v13))
						var p73 int32
						if uint32(v16) < uint32(v15) {
							p73 = 1
						}
						var p74 int32
						if uint32(v13) < uint32(v16) {
							p74 = 1
						}
						v13 = p73 | p74
						v1 = v1 + i32(4)
						v9 = v9 + i32(4)
						v14 = v14 + i32(-1)
						if v14 != 0 {
							goto l46
						}
					}
					if v13 == 0 {
						m.fn256(i32(1108963), i32(26), i32(1108992))
						panic("unreachable")
					}
				l45:
					store32(m.memory[int64(uint32(v3))+168:], uint32(v20))
					v29 = v29 | i32(4)
				l44:
					p75 := v20
					if uint32(v25) > uint32(v20) {
						p75 = v25
					}
					v30 = p75
					if uint32(v30) >= uint32(i32(41)) {
						m.fn151(i32(0), v30, i32(40), i32(1108992))
						panic("unreachable")
					}
					v1 = v30 << 2
				l50:
					{
						if v1 == 0 {
							goto l49
						}
						v9 = v24 + v1
						v1 = v1 + i32(-4)
						t76 := int32(load32(m.memory[uint32(v1+(v3+i32(8))):]))
						v16 = t76
						t77 := int32(load32(m.memory[uint32(v9):]))
						t78 := v16
						v9 = t77
						if t78 == v9 {
							goto l50
						}
					}
					if uint32(v16) >= uint32(v9) {
						goto l49
					}
					v30 = v20
					goto l51
				l49:
					if v30 == 0 {
						goto l52
					}
					v13 = i32(1)
					v9 = v3 + i32(672)
					v1 = v3 + i32(8)
					v14 = v30
				l53:
					{
						t79 := int32(load32(m.memory[uint32(v1):]))
						t80 := v1
						v15 = t79
						t81 := int32(load32(m.memory[uint32(v9):]))
						v16 = v15 + (t81 ^ i32(-1))
						v13 = v16 + v13&i32(1)
						store32(m.memory[uint32(t80):], uint32(v13))
						var p82 int32
						if uint32(v16) < uint32(v15) {
							p82 = 1
						}
						var p83 int32
						if uint32(v13) < uint32(v16) {
							p83 = 1
						}
						v13 = p82 | p83
						v1 = v1 + i32(4)
						v9 = v9 + i32(4)
						v14 = v14 + i32(-1)
						if v14 != 0 {
							goto l53
						}
					}
					if v13 == 0 {
						m.fn256(i32(1108963), i32(26), i32(1108992))
						panic("unreachable")
					}
				l52:
					store32(m.memory[int64(uint32(v3))+168:], uint32(v30))
					v29 = v29 + i32(2)
				l51:
					p84 := v30
					if uint32(v17) > uint32(v30) {
						p84 = v17
					}
					v20 = p84
					if uint32(v20) >= uint32(i32(41)) {
						m.fn151(i32(0), v20, i32(40), i32(1108992))
						panic("unreachable")
					}
					v1 = v20 << 2
				l57:
					{
						if v1 == 0 {
							goto l56
						}
						v1 = v1 + i32(-4)
						t85 := int32(load32(m.memory[uint32(v1+(v3+i32(8))):]))
						v9 = t85
						t86 := int32(load32(m.memory[uint32(v1+(v3+i32(508))):]))
						t87 := v9
						v16 = t86
						if t87 == v16 {
							goto l57
						}
					}
					if uint32(v9) >= uint32(v16) {
						goto l56
					}
					v20 = v30
					goto l58
				l56:
					if v20 == 0 {
						goto l59
					}
					v13 = i32(1)
					v9 = v3 + i32(508)
					v1 = v3 + i32(8)
					v14 = v20
				l60:
					{
						t88 := int32(load32(m.memory[uint32(v1):]))
						t89 := v1
						v15 = t88
						t90 := int32(load32(m.memory[uint32(v9):]))
						v16 = v15 + (t90 ^ i32(-1))
						v13 = v16 + v13&i32(1)
						store32(m.memory[uint32(t89):], uint32(v13))
						var p91 int32
						if uint32(v16) < uint32(v15) {
							p91 = 1
						}
						var p92 int32
						if uint32(v13) < uint32(v16) {
							p92 = 1
						}
						v13 = p91 | p92
						v1 = v1 + i32(4)
						v9 = v9 + i32(4)
						v14 = v14 + i32(-1)
						if v14 != 0 {
							goto l60
						}
					}
					if v13 == 0 {
						m.fn256(i32(1108963), i32(26), i32(1108992))
						panic("unreachable")
					}
				l59:
					store32(m.memory[int64(uint32(v3))+168:], uint32(v20))
					v29 = v29 + i32(1)
				l58:
					if v28 == i32(17) {
						m.fn158(i32(17), i32(17), i32(1130784))
						panic("unreachable")
					}
					m.memory[uint32(v2+v28)] = byte(v29 + i32(48))
					t93 := int32(load32(m.memory[int64(uint32(v3))+336:]))
					v30 = t93
					p94 := v20
					if uint32(v30) > uint32(v20) {
						p94 = v30
					}
					v1 = p94
					if uint32(v1) >= uint32(i32(41)) {
						m.fn151(i32(0), v1, i32(40), i32(1108992))
						panic("unreachable")
					}
					v27 = v28 + i32(1)
					v1 = v1 << 2
					{
					l66:
						{
							if v1 != 0 {
								goto l64
							}
							v31 = i32(0)
							goto l65
						l64:
							v9 = v23 + v1
							v1 = v1 + i32(-4)
							t95 := int32(load32(m.memory[uint32(v1+(v3+i32(8))):]))
							v16 = t95
							t96 := int32(load32(m.memory[uint32(v9):]))
							t97 := v16
							v9 = t96
							if t97 == v9 {
								goto l66
							}
						}
						var p98 int32
						if uint32(v16) > uint32(v9) {
							p98 = 1
						}
						var p99 int32
						if uint32(v16) < uint32(v9) {
							p99 = 1
						}
						v31 = p98 - p99
					}
				l65:
					memory_copy(m.memory, uint32(v3+i32(1164)), uint32(v3+i32(8)), uint32(i32(164)))
					{
						t100 := int32(load32(m.memory[int64(uint32(v3))+1324:]))
						t101 := v18
						v1 = t100
						p102 := v1
						if uint32(v18) > uint32(v1) {
							p102 = t101
						}
						v29 = p102
						if uint32(v29) > uint32(i32(40)) {
							m.fn151(i32(0), v29, i32(40), i32(1108992))
							panic("unreachable")
						}
						if v29 != 0 {
							goto l68
						}
						v29 = i32(0)
						goto l69
					}
				l68:
					v13 = i32(0)
					v9 = v3 + i32(344)
					v1 = v3 + i32(1164)
					v14 = v29
				l70:
					{
						t103 := int32(load32(m.memory[uint32(v9):]))
						t104 := v1
						v15 = t103
						t105 := int32(load32(m.memory[uint32(v1):]))
						v16 = v15 + t105
						v13 = v16 + v13&i32(1)
						store32(m.memory[uint32(t104):], uint32(v13))
						var p106 int32
						if uint32(v16) < uint32(v15) {
							p106 = 1
						}
						var p107 int32
						if uint32(v13) < uint32(v16) {
							p107 = 1
						}
						v13 = p106 | p107
						v1 = v1 + i32(4)
						v9 = v9 + i32(4)
						v14 = v14 + i32(-1)
						if v14 != 0 {
							goto l70
						}
					}
					if v13 == 0 {
						goto l69
					}
					if v29 == i32(40) {
						m.fn158(i32(40), i32(40), i32(1108992))
						panic("unreachable")
					}
					store32(m.memory[uint32(v3+i32(1164)+v29<<2):], uint32(i32(1)))
					v29 = v29 + i32(1)
				l69:
					store32(m.memory[int64(uint32(v3))+1324:], uint32(v29))
					p108 := v17
					if uint32(v29) > uint32(v17) {
						p108 = v29
					}
					v1 = p108
					if uint32(v1) >= uint32(i32(41)) {
						m.fn151(i32(0), v1, i32(40), i32(1108992))
						panic("unreachable")
					}
					v1 = v1 << 2
					{
					l75:
						{
							if v1 != 0 {
								goto l73
							}
							v1 = i32(0)
							goto l74
						l73:
							v9 = v22 + v1
							v16 = v21 + v1
							v1 = v1 + i32(-4)
							t109 := int32(load32(m.memory[uint32(v16):]))
							v16 = t109
							t110 := int32(load32(m.memory[uint32(v9):]))
							t111 := v16
							v9 = t110
							if t111 == v9 {
								goto l75
							}
						}
						var p112 int32
						if uint32(v16) > uint32(v9) {
							p112 = 1
						}
						var p113 int32
						if uint32(v16) < uint32(v9) {
							p113 = 1
						}
						v1 = p112 - p113
					}
				l74:
					if v31 < v8 {
						goto l76
					}
					if v1 < v8 {
						goto l77
					}
					v16 = i32(0)
					v13 = i32(0)
					{
						if v20 == 0 {
							goto l78
						}
						t114 := v3 + i32(8)
						v9 = v20 << 2
						v13 = t114 + v9
						v4 = i64(0)
						v1 = v3 + i32(8)
					l79:
						{
							t115 := int64(load32(m.memory[uint32(v1):]))
							t116 := v1
							v4 = t115*i64(10) + v4
							store32(m.memory[uint32(t116):], uint32(v4))
							v1 = v1 + i32(4)
							v4 = int64(uint64(v4) >> 32)
							v9 = v9 + i32(-4)
							if v9 != 0 {
								goto l79
							}
						}
						if !(v4 == 0) {
							goto l80
						}
						v13 = v20
						goto l78
					l80:
						if v20 == i32(40) {
							m.fn158(i32(40), i32(40), i32(1108992))
							panic("unreachable")
						}
						store32(m.memory[uint32(v13):], uint32(int32(v4)))
						v13 = v20 + i32(1)
					}
				l78:
					store32(m.memory[int64(uint32(v3))+168:], uint32(v13))
					{
						if v30 == 0 {
							goto l82
						}
						t117 := v3 + i32(176)
						v9 = v30 << 2
						v16 = t117 + v9
						v4 = i64(0)
						v1 = v3 + i32(176)
					l83:
						{
							t118 := int64(load32(m.memory[uint32(v1):]))
							t119 := v1
							v4 = t118*i64(10) + v4
							store32(m.memory[uint32(t119):], uint32(v4))
							v1 = v1 + i32(4)
							v4 = int64(uint64(v4) >> 32)
							v9 = v9 + i32(-4)
							if v9 != 0 {
								goto l83
							}
						}
						if !(v4 == 0) {
							goto l84
						}
						v16 = v30
						goto l82
					l84:
						if v30 == i32(40) {
							m.fn158(i32(40), i32(40), i32(1108992))
							panic("unreachable")
						}
						store32(m.memory[uint32(v16):], uint32(int32(v4)))
						v16 = v30 + i32(1)
					}
				l82:
					store32(m.memory[int64(uint32(v3))+336:], uint32(v16))
					{
						if v18 != 0 {
							goto l86
						}
						v18 = i32(0)
						goto l87
					l86:
						t120 := v3 + i32(344)
						v9 = v18 << 2
						v16 = t120 + v9
						v4 = i64(0)
						v1 = v3 + i32(344)
					l88:
						{
							t121 := int64(load32(m.memory[uint32(v1):]))
							t122 := v1
							v4 = t121*i64(10) + v4
							store32(m.memory[uint32(t122):], uint32(v4))
							v1 = v1 + i32(4)
							v4 = int64(uint64(v4) >> 32)
							v9 = v9 + i32(-4)
							if v9 != 0 {
								goto l88
							}
						}
						if v4 == 0 {
							goto l87
						}
						if v18 == i32(40) {
							m.fn158(i32(40), i32(40), i32(1108992))
							panic("unreachable")
						}
						store32(m.memory[uint32(v16):], uint32(int32(v4)))
						v18 = v18 + i32(1)
					}
				l87:
					store32(m.memory[int64(uint32(v3))+504:], uint32(v18))
					p123 := v13
					if uint32(v19) > uint32(v13) {
						p123 = v19
					}
					v20 = p123
					if uint32(v20) < uint32(i32(41)) {
						goto l90
					}
				}
			}
		l34:
			m.fn151(i32(0), v20, i32(40), i32(1108992))
			panic("unreachable")
		l76:
			if v1 >= v8 {
				goto l91
			}
			_ = m.fn1650(v3+i32(8), i32(1))
			t125 := int32(load32(m.memory[int64(uint32(v3))+168:]))
			t126 := v17
			v1 = t125
			p127 := v1
			if uint32(v17) > uint32(v1) {
				p127 = t126
			}
			v1 = p127
			if uint32(v1) >= uint32(i32(41)) {
				m.fn151(i32(0), v1, i32(40), i32(1108992))
				panic("unreachable")
			}
			v1 = v1 << 2
			v13 = v3 + i32(8) + i32(-4)
			v14 = v3 + i32(508) + i32(-4)
		l93:
			{
				if v1 == 0 {
					goto l77
				}
				v9 = v14 + v1
				v16 = v13 + v1
				v1 = v1 + i32(-4)
				t128 := int32(load32(m.memory[uint32(v16):]))
				v16 = t128
				t129 := int32(load32(m.memory[uint32(v9):]))
				t130 := v16
				v9 = t129
				if t130 == v9 {
					goto l93
				}
			}
			if uint32(v16) < uint32(v9) {
				goto l91
			}
		}
	l77:
		v13 = v2 + v27
		v1 = v27
		{
		l95:
			{
				v9 = v1
				if v9 == 0 {
					m.memory[uint32(v2)] = byte(i32(49))
					if v28 == 0 {
						goto l96
					}
					memory_fill(m.memory, uint32(v2+i32(1)), i32(48), uint32(v28))
				l96:
					if uint32(v28) > uint32(i32(15)) {
						m.fn158(v27, i32(17), i32(1130800))
						panic("unreachable")
					}
					m.memory[uint32(v13)] = byte(i32(48))
					v10 = v10 + i32(1)
					v27 = v28 + i32(2)
					goto l98
				}
				v1 = v9 + i32(-1)
				v16 = v1 + v2
				t131 := int32(m.memory[uint32(v16)])
				if t131 == i32(57) {
					goto l95
				}
			}
			t132 := int32(m.memory[uint32(v16)])
			m.memory[uint32(v16)] = byte(t132 + i32(1))
			v1 = v27 - v9
			if v1 == 0 {
				goto l91
			}
			memory_fill(m.memory, uint32(v2+v9), i32(48), uint32(v1))
			goto l91
		}
	}
l91:
	if uint32(v28) <= uint32(i32(16)) {
		goto l98
	}
	m.fn151(i32(0), v27, i32(17), i32(1130816))
	panic("unreachable")
l98:
	store16(m.memory[int64(uint32(v0))+8:], uint16(v10))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v27))
	store32(m.memory[uint32(v0):], uint32(v2))
	m.g0 = v3 + i32(1328)
}
func (m *Module) fn1668(v0, v1, v2 int32) {
	var v3 int32
	var v4 int64
	t0 := m.g0
	v3 = t0 - i32(32)
	m.g0 = v3
	store32(m.memory[int64(uint32(v3))+8:], uint32(v1))
	store32(m.memory[int64(uint32(v3))+12:], uint32(v0))
	t1 := v3
	v4 = int64(uint32(i32(5))) << 32
	store64(m.memory[int64(uint32(t1))+24:], uint64(v4|int64(uint32(v3+i32(12)))))
	store64(m.memory[int64(uint32(v3))+16:], uint64(v4|int64(uint32(v3+i32(8)))))
	m.fn91(i32(1068969), v3+i32(16), v2)
	panic("unreachable")
}
func (m *Module) fn1669(v0, v1 int32) int32 {
	var v2, v3 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	t1 := int32(m.memory[uint32(v0)])
	v3 = t1
	v0 = i32(3)
l0:
	{
		t2 := int32(m.memory[uint32(v3&i32(15)+i32(1131672))])
		m.memory[uint32(v2+i32(14)+v0+i32(-2))] = byte(t2)
		v0 = v0 + i32(-1)
		v3 = int32(uint32(v3)>>4) & i32(15)
		if v3 != 0 {
			goto l0
		}
	}
	t3 := m.fn1638(v1, i32(1), i32(1131670), i32(2), v2+i32(14)+v0+i32(-1), i32(3)-v0)
	v0 = t3
	m.g0 = v2 + i32(16)
	return v0
}
func (m *Module) fn1670(v0, v1 int32) int32 {
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
			t4 := m.fn1671(t2, t3, v1)
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
						t31 := m.fn1671(v7, v6, v2)
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
			t26 := m.fn1671(t24, t25, v2)
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
func (m *Module) fn1671(v0, v1, v2 int32) int32 {
	var v3, v4, v5, v6, v7, v8 int32
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
		v4 = t5
		if v4 != 0 {
			t6 := int32(load32(m.memory[int64(uint32(v2))+8:]))
			v5 = t6
			v6 = v5 + v4*i32(12)
		l14:
			{
				{
					{
						{
							t7 := int32(load16(m.memory[uint32(v5):]))
							switch t7 {
							default:
								t8 := int32(load32(m.memory[int64(uint32(v5))+4:]))
								v2 = t8
								if uint32(v2) < uint32(i32(65)) {
									goto l7
								}
								t9 := int32(load32(m.memory[uint32(v1+i32(12)):]))
								v4 = t9
							l8:
								{
									t10 := m.t0[uint(v4)].(func(int32, int32, int32) int32)(v0, i32(1109128), i32(64))
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
								t11 := int32(load16(m.memory[int64(uint32(v5))+2:]))
								v2 = t11
								m.memory[int64(uint32(v3))+12] = byte(i32(0))
								store32(m.memory[int64(uint32(v3))+8:], uint32(i32(0)))
								if v2 != 0 {
									goto l10
								}
								v7 = i32(1)
								goto l11
							case 2:
								t12 := int32(load32(m.memory[int64(uint32(v5))+4:]))
								t13 := int32(load32(m.memory[int64(uint32(v5))+8:]))
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
						v4 = t16
					}
				l9:
					t17 := m.t0[uint(v4)].(func(int32, int32, int32) int32)(v0, i32(1109128), v2)
					if t17 != 0 {
						goto l1
					}
					goto l12
				}
			l10:
				v7 = int32(uint32((v2+i32(0x5fff6))&(v2+i32(524188))^(v2+i32(916504))&(v2+i32(514288)))>>17) + i32(1)
			l11:
				v4 = v7
			l13:
				{
					v4 = v4 + i32(-1)
					t18 := int32(uint32(v2&i32(0xffff)) / uint32(i32(10)))
					t19 := v4 + (v3 + i32(8))
					t20 := v2
					v8 = t18
					m.memory[uint32(t19)] = byte(t20 - v8*i32(10) | i32(48))
					v2 = v8
					if v4 != 0 {
						goto l13
					}
				}
				t21 := int32(load32(m.memory[uint32(v1+i32(12)):]))
				t22 := m.t0[uint(t21)].(func(int32, int32, int32) int32)(v0, v3+i32(8), v7)
				if t22 != 0 {
					goto l1
				}
			}
		l12:
			v5 = v5 + i32(12)
			if v5 != v6 {
				goto l14
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
func (m *Module) fn1672(v0 int32, v1, v2 int64) {
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
		t2 := int64(load64(m.memory[uint32(v4+i32(1123784)):]))
		t3 := v2
		v6 = int64(bits.LeadingZeros64(uint64(v2)))
		v7 = i64_shl(t3, v6)
		m.fn1853(t1, t2, i64(0), v7, i64(0))
		t4 := int64(load64(m.memory[int64(uint32(v3))+16:]))
		v8 = t4
		{
			t5 := int64(load64(m.memory[int64(uint32(v3))+24:]))
			v2 = t5
			if v2&i64(511) != i64(511) {
				goto l2
			}
			t6 := int64(load64(m.memory[uint32(v4+i32(1118312)+i32(5480)):]))
			m.fn1853(v3, t6, i64(0), v7, i64(0))
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
func (m *Module) fn1673(v0, v1 int32) int32 {
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
			t16 := m.fn1676(t13, t14, p15, t12)
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
			t7 := m.fn1674(t4, t5, p6, i32(1))
			return t7
		}
	l1:
		t8 := v1
		t9 := v4
		var p10 int32
		if v3 != i32(0) {
			p10 = 1
		}
		t11 := m.fn1675(t8, t9, p10)
		return t11
	}
}
func (m *Module) fn1674(v0 int32, v1 float64, v2, v3 int32) int32 {
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
					store32(m.memory[int64(uint32(v4))+36:], uint32(i32(1108002)))
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
				p21 := i32(1108001)
				if v12 != 0 {
					p21 = i32(1108000)
				}
				p22 := i32(1)
				if v12 != 0 {
					p22 = i32(1108000)
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
					store32(m.memory[int64(uint32(v4))+36:], uint32(i32(1108005)))
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
				store32(m.memory[int64(uint32(v4))+36:], uint32(i32(1108008)))
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
				p15 = i32(1108000)
			}
			v11 = p15
			p16 := i32(1108001)
			if v12 != 0 {
				p16 = i32(1108000)
			}
			v12 = p16
			v3 = int32(int64(uint64(v5) >> 63))
			m.fn1666(v4+i32(32), v4+i32(96), v4+i32(15))
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
		m.fn1667(v4+i32(80), v4+i32(96), v4+i32(15))
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
		m.fn1661(v4, t27, t28, t29, v10, v4+i32(32))
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
	store32(m.memory[int64(uint32(v4))+36:], uint32(i32(1108009)))
	v11 = v4 + i32(32)
l11:
	store32(m.memory[int64(uint32(v4))+92:], uint32(v10))
	store32(m.memory[int64(uint32(v4))+88:], uint32(v11))
	store32(m.memory[int64(uint32(v4))+84:], uint32(v2))
	store32(m.memory[int64(uint32(v4))+80:], uint32(v12))
	t32 := m.fn1670(v0, v4+i32(80))
	v10 = t32
	m.g0 = v4 + i32(128)
	return v10
}
func (m *Module) fn1675(v0 int32, v1 float64, v2 int32) int32 {
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
						store32(m.memory[int64(uint32(v3))+28:], uint32(i32(1108002)))
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
					p18 := i32(1108001)
					if v10 != 0 {
						p18 = i32(1108000)
					}
					p19 := i32(1)
					if v10 != 0 {
						p19 = i32(1108000)
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
						store32(m.memory[int64(uint32(v3))+28:], uint32(i32(1108011)))
						goto l10
					}
					store32(m.memory[int64(uint32(v3))+32:], uint32(i32(3)))
					store32(m.memory[int64(uint32(v3))+28:], uint32(i32(1108005)))
					goto l10
				}
				m.fn1666(v3+i32(96), v3+i32(112), v3+i32(7))
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
			m.fn1667(v3+i32(144), v3+i32(112), v3+i32(7))
		l7:
			t22 := int32(load32(m.memory[int64(uint32(v3))+148:]))
			v9 = t22
			if v9 == 0 {
				m.fn256(i32(1111329), i32(33), i32(1111428))
				panic("unreachable")
			}
			t23 := int32(load32(m.memory[int64(uint32(v3))+144:]))
			v11 = t23
			t24 := int32(m.memory[uint32(v11)])
			if uint32(t24) <= uint32(i32(48)) {
				m.fn256(i32(1111380), i32(31), i32(1111444))
				panic("unreachable")
			}
			var p25 int32
			if v4 < i64(0) {
				p25 = 1
			}
			v10 = p25
			p26 := i32(1)
			if v10 != 0 {
				p26 = i32(1108000)
			}
			v13 = p26
			p27 := i32(1108001)
			if v10 != 0 {
				p27 = i32(1108000)
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
				store32(m.memory[int64(uint32(v3))+40:], uint32(i32(1109519)))
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
		p34 := i32(1111460)
		if v13 != 0 {
			p34 = i32(1111461)
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
	t38 := m.fn1670(v0, v3+i32(96))
	v9 = t38
	m.g0 = v3 + i32(160)
	return v9
}
func (m *Module) fn1676(v0 int32, v1 float64, v2, v3 int32) int32 {
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
					store32(m.memory[int64(uint32(v4))+1044:], uint32(i32(1108002)))
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
				p16 := i32(1108001)
				if v12 != 0 {
					p16 = i32(1108000)
				}
				p17 := i32(1)
				if v12 != 0 {
					p17 = i32(1108000)
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
					store32(m.memory[int64(uint32(v4))+1044:], uint32(i32(1108005)))
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
					store32(m.memory[int64(uint32(v4))+1044:], uint32(i32(1108009)))
					v3 = v4 + i32(1040)
					goto l10
				}
				v10 = i32(1)
				store32(m.memory[int64(uint32(v4))+1048:], uint32(i32(1)))
				store32(m.memory[int64(uint32(v4))+1044:], uint32(i32(1108008)))
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
			m.fn256(i32(1108076), i32(37), i32(1108116))
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
			p21 = i32(1108000)
		}
		v14 = p21
		p22 := i32(1108001)
		if v7 != 0 {
			p22 = i32(1108000)
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
		m.fn1662(t23, t24, t25, t26, t28, v10)
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
		m.fn1664(v4+i32(1088), v4+i32(1104), v4+i32(16), v16, v10)
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
			m.fn1661(v4+i32(8), t35, t36, v7, v11, v4+i32(1040))
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
		store32(m.memory[int64(uint32(v4))+1044:], uint32(i32(1108008)))
		v3 = v4 + i32(1040)
		goto l10
	l14:
		store32(m.memory[int64(uint32(v4))+1056:], uint32(v11))
		store16(m.memory[int64(uint32(v4))+1052:], uint16(i32(0)))
		store32(m.memory[int64(uint32(v4))+1048:], uint32(i32(2)))
		store32(m.memory[int64(uint32(v4))+1044:], uint32(i32(1108009)))
		v3 = v4 + i32(1040)
	}
l10:
	store32(m.memory[int64(uint32(v4))+1100:], uint32(v10))
	store32(m.memory[int64(uint32(v4))+1096:], uint32(v3))
	store32(m.memory[int64(uint32(v4))+0x444:], uint32(v2))
	store32(m.memory[int64(uint32(v4))+1088:], uint32(v12))
	t39 := m.fn1670(v0, v4+i32(1088))
	v10 = t39
	m.g0 = v4 + i32(1136)
	return v10
}
func (m *Module) fn1677(v0 int32, v1 float64, v2, v3 int32) int32 {
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
					store32(m.memory[int64(uint32(v4))+1036:], uint32(i32(1108002)))
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
				p18 := i32(1108001)
				if v12 != 0 {
					p18 = i32(1108000)
				}
				p19 := i32(1)
				if v12 != 0 {
					p19 = i32(1108000)
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
					store32(m.memory[int64(uint32(v4))+1036:], uint32(i32(1108005)))
					store16(m.memory[int64(uint32(v4))+1032:], uint16(i32(2)))
					goto l10
				}
				store16(m.memory[int64(uint32(v4))+1032:], uint16(i32(2)))
				if v3&i32(0xffff) != 0 {
					store32(m.memory[int64(uint32(v4))+1064:], uint32(i32(2)))
					store32(m.memory[int64(uint32(v4))+1060:], uint32(i32(1108014)))
					store16(m.memory[int64(uint32(v4))+1056:], uint16(i32(2)))
					store32(m.memory[int64(uint32(v4))+1048:], uint32(v11))
					store16(m.memory[int64(uint32(v4))+1044:], uint16(i32(0)))
					store32(m.memory[int64(uint32(v4))+1040:], uint32(i32(2)))
					store32(m.memory[int64(uint32(v4))+1036:], uint32(i32(1108009)))
					v10 = i32(3)
					goto l10
				}
				store32(m.memory[int64(uint32(v4))+1040:], uint32(i32(3)))
				store32(m.memory[int64(uint32(v4))+1036:], uint32(i32(1108011)))
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
			m.fn256(i32(1108016), i32(41), i32(1108060))
			panic("unreachable")
		}
	l6:
		m.fn1662(v4+i32(1104), v8, v12, v4+i32(8), v10, i32(0x8000))
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
		m.fn1664(v4+i32(1152), v4+i32(1120), v4+i32(8), v10, i32(0x8000))
	l12:
		{
			t25 := int32(load32(m.memory[int64(uint32(v4))+1156:]))
			v14 = t25
			if v14 == 0 {
				m.fn256(i32(1111329), i32(33), i32(1111428))
				panic("unreachable")
			}
			t26 := int32(load32(m.memory[int64(uint32(v4))+1152:]))
			v15 = t26
			t27 := int32(m.memory[uint32(v15)])
			if uint32(t27) <= uint32(i32(48)) {
				m.fn256(i32(1111380), i32(31), i32(1111444))
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
				store32(m.memory[int64(uint32(v4))+1048:], uint32(i32(1109519)))
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
			v19 = i32(1109519)
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
			p30 = i32(1108000)
		}
		v12 = p30
		p31 := i32(1108001)
		if v7 != 0 {
			p31 = i32(1108000)
		}
		v11 = p31
		v3 = int32(v5)
		if v16 < i32(1) {
			goto l19
		}
		v7 = v4 + i32(1032) + v10*i32(12)
		store32(m.memory[int64(uint32(v7))+8:], uint32(i32(1)))
		store32(m.memory[int64(uint32(v7))+4:], uint32(i32(1111460)))
		store16(m.memory[uint32(v7):], uint16(i32(2)))
		v7 = v16 + i32(-1)
		goto l20
	l19:
		v7 = v4 + i32(1032) + v10*i32(12)
		store32(m.memory[int64(uint32(v7))+8:], uint32(i32(2)))
		store32(m.memory[int64(uint32(v7))+4:], uint32(i32(1111461)))
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
	t34 := m.fn1670(v0, v4+i32(1104))
	v10 = t34
	m.g0 = v4 + i32(1168)
	return v10
}
func (m *Module) fn1678(v0, v1 int32) int32 {
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
			store64(m.memory[int64(uint32(v2))+24:], uint64(int64(uint32(i32(5)))<<32|int64(uint32(v0))))
			store64(m.memory[int64(uint32(v2))+16:], uint64(int64(uint32(i32(96)))<<32|int64(uint32(v2+i32(15)))))
			t3 := int32(load32(m.memory[uint32(v1):]))
			t4 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t5 := m.fn100(t3, t4, i32(1049281), v2+i32(16))
			v0 = t5
			goto l1
		}
	l0:
		store64(m.memory[int64(uint32(v2))+16:], uint64(int64(uint32(i32(5)))<<32|int64(uint32(v0))))
		t6 := int32(load32(m.memory[uint32(v1):]))
		t7 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t8 := m.fn100(t6, t7, i32(1049330), v2+i32(16))
		v0 = t8
	}
l1:
	m.g0 = v2 + i32(32)
	return v0
}
func (m *Module) fn1679(v0 int32) int32 {
	var v1, v2 int32
	t0 := m.g0
	v1 = t0 - i32(16)
	m.g0 = v1
	m.fn1680(v1+i32(8), i32(8), v0, i32(0))
	{
		t1 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		v2 = t1
		if v2 != 0 {
			m.g0 = v1 + i32(16)
			return v2
		}
		m.fn85(i32(8), v0)
		panic("unreachable")
	}
}
func (m *Module) fn1680(v0, v1, v2, v3 int32) {
	{
		if v2 == 0 {
			goto l0
		}
		{
			if v3 != 0 {
				goto l1
			}
			t0 := m.fn248(v2, v1)
			v1 = t0
			goto l0
		}
	l1:
		t1 := m.fn1561(v2, v1)
		v1 = t1
	}
l0:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v2))
	store32(m.memory[uint32(v0):], uint32(v1))
}
func (m *Module) fn1681(v0, v1 int32) {
	var v2, v3 int32
	t0 := int32(load32(m.memory[uint32(v1):]))
	v2 = t0
	t1 := m.fn1682(v2)
	v3 = t1
	store64(m.memory[int64(uint32(v0))+8:], uint64(i64(0)))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
	store32(m.memory[uint32(v0):], uint32(v1))
	t2 := int32(load32(m.memory[int64(uint32(v2))+44:]))
	store32(m.memory[int64(uint32(v0))+16:], uint32(t2))
}
func (m *Module) fn1682(v0 int32) int32 {
	var v1 int32
	{
		t0 := m.fn936(v0 + i32(32))
		v1 = t0
		t1 := int32(load32(m.memory[int64(uint32(v0))+56:]))
		t2 := v1
		v0 = t1
		if uint32(t2) <= uint32(v0) {
			return v1
		}
		m.fn151(i32(0), v1, v0, i32(1148268))
		panic("unreachable")
	}
}
func (m *Module) fn1683(v0, v1 int32) {
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
		v5 = t12
		t13 := int32(load32(m.memory[int64(uint32(v5))+56:]))
		v1 = t13
		if uint32(v3) < uint32(v4) {
			goto l3
		}
		if uint32(v3) > uint32(v1) {
			goto l3
		}
		v3 = v3 - v4
		t14 := int32(load32(m.memory[int64(uint32(v5))+52:]))
		v1 = t14 + v4
	}
l1:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
	store32(m.memory[uint32(v0):], uint32(v1))
	m.g0 = v2 + i32(16)
	return
l3:
	m.fn151(v4, v3, v1, i32(1148588))
	panic("unreachable")
}
func (m *Module) fn1684(v0, v1, v2, v3, v4 int32) {
	var v5, v6, v7, v8 int32
	var v9 int64
	t0 := m.g0
	v5 = t0 - i32(32)
	m.g0 = v5
	{
		v1 = v2 + v1
		if uint32(v1) >= uint32(v2) {
			goto l0
		}
		v5 = i32(0)
		goto l1
	l0:
		v6 = i32(0)
		v7 = v5 + i32(20)
		{
			t1 := int32(load32(m.memory[uint32(v0):]))
			t2 := int64(uint32(v4))
			t3 := v1
			v8 = t1
			v2 = v8 << 1
			p4 := v2
			if uint32(v1) > uint32(v2) {
				p4 = t3
			}
			v2 = p4
			t6 := v2
			p5 := i32(4)
			if v4 == i32(1) {
				p5 = i32(8)
			}
			v1 = p5
			p7 := v1
			if uint32(v2) > uint32(v1) {
				p7 = t6
			}
			v2 = p7
			v9 = t2 * int64(uint32(v2))
			if int32(int64(uint64(v9)>>32)) != 0 {
				goto l2
			}
			v1 = int32(v9)
			if uint32(v1) > uint32(i32(-0x80000000)-v3) {
				goto l2
			}
			{
				if v8 != 0 {
					goto l3
				}
				v4 = i32(0)
				v8 = v5 + i32(28)
				goto l4
			l3:
				t8 := int32(load32(m.memory[int64(uint32(v0))+4:]))
				v6 = t8
				store32(m.memory[int64(uint32(v5))+28:], uint32(v3))
				v4 = v8 * v4
				v8 = v5 + i32(24)
			}
		l4:
			store32(m.memory[uint32(v8):], uint32(v4))
			{
				{
					t9 := int32(load32(m.memory[int64(uint32(v5))+28:]))
					if t9 == 0 {
						goto l5
					}
					{
						t10 := int32(load32(m.memory[int64(uint32(v5))+24:]))
						v4 = t10
						if v4 != 0 {
							t12 := m.fn89(v6, v4, v3, v1)
							v4 = t12
							goto l7
						}
						m.fn1680(v5+i32(8), v3, v1, i32(0))
						t11 := int32(load32(m.memory[int64(uint32(v5))+8:]))
						v4 = t11
						goto l7
					}
				}
			l5:
				m.fn1680(v5, v3, v1, i32(0))
				t13 := int32(load32(m.memory[uint32(v5):]))
				v4 = t13
			}
		l7:
			if v4 != 0 {
				goto l8
			}
			store32(m.memory[int64(uint32(v5))+20:], uint32(v3))
			v7 = v5 + i32(16)
			v6 = v1
		}
	l2:
		store32(m.memory[uint32(v7):], uint32(v6))
		t14 := int32(load32(m.memory[int64(uint32(v5))+16:]))
		v4 = t14
		t15 := int32(load32(m.memory[int64(uint32(v5))+20:]))
		v5 = t15
	}
l1:
	m.fn2(v5, v4)
	panic("unreachable")
l8:
	store32(m.memory[uint32(v0):], uint32(v2))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v4))
	m.g0 = v5 + i32(32)
}
func (m *Module) fn1685(v0, v1, v2, v3 int32) {
	var v4 int32
	var v5 int64
	var v6, v7 int32
	t0 := m.g0
	v4 = t0 - i32(16)
	m.g0 = v4
	v5 = int64(uint32(v3)) * int64(uint32(v1))
	if int32(int64(uint64(v5)>>32)) != 0 {
		goto l0
	}
	v6 = int32(v5)
	if uint32(v6) > uint32(i32(-0x80000000)-v2) {
		goto l0
	}
	if v6 != 0 {
		goto l1
	}
	store32(m.memory[int64(uint32(v0))+8:], uint32(v2))
	v3 = i32(0)
	store32(m.memory[int64(uint32(v0))+4:], uint32(i32(0)))
	goto l2
l0:
	store32(m.memory[int64(uint32(v0))+4:], uint32(i32(0)))
	v3 = i32(1)
	goto l2
l1:
	v3 = i32(1)
	m.fn1680(v4+i32(8), v2, v6, i32(1))
	{
		t1 := int32(load32(m.memory[int64(uint32(v4))+8:]))
		v7 = t1
		if v7 == 0 {
			goto l3
		}
		store32(m.memory[int64(uint32(v0))+8:], uint32(v7))
		store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
		v3 = i32(0)
		goto l2
	}
l3:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v6))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v2))
l2:
	store32(m.memory[uint32(v0):], uint32(v3))
	m.g0 = v4 + i32(16)
}
func (m *Module) fn1686(v0, v1 int32) int32 {
	var v2, v3, v4 int32
	t0 := m.g0
	v2 = t0 - i32(64)
	m.g0 = v2
	m.fn1685(v2+i32(8), v0, i32(1), i32(1))
	t1 := int32(load32(m.memory[int64(uint32(v2))+12:]))
	v3 = t1
	{
		t2 := int32(load32(m.memory[int64(uint32(v2))+8:]))
		if t2 != i32(1) {
			t4 := int32(load32(m.memory[int64(uint32(v2))+16:]))
			v4 = t4
			m.fn937(v2+i32(48), v1)
			store32(m.memory[int64(uint32(v2))+60:], uint32(i32(0)))
			t5 := int64(load64(m.memory[int64(uint32(v2))+48:]))
			store64(m.memory[int64(uint32(v2))+32:], uint64(t5))
			t6 := int64(load64(m.memory[int64(uint32(v2))+56:]))
			store64(m.memory[int64(uint32(v2))+40:], uint64(t6))
			t7 := m.fn1679(i32(64))
			v1 = t7
			store64(m.memory[uint32(v1):], uint64(i64(0)))
			memory_copy(m.memory, uint32(v1+i32(8)), uint32(v2+i32(8)), uint32(i32(40)))
			store32(m.memory[int64(uint32(v1))+56:], uint32(v0))
			store32(m.memory[int64(uint32(v1))+52:], uint32(v4))
			store32(m.memory[int64(uint32(v1))+48:], uint32(v3))
			m.g0 = v2 + i32(64)
			return v1
		}
		t3 := int32(load32(m.memory[int64(uint32(v2))+16:]))
		m.fn2(v3, t3)
		panic("unreachable")
	}
}
func (m *Module) fn1687(v0 int32) {
	t0 := int32(load32(m.memory[int64(uint32(v0))+48:]))
	t1 := int32(load32(m.memory[uint32(v0+i32(52)):]))
	m.fn608(t0, t1, i32(1), i32(1))
	t2 := int32(load32(m.memory[int64(uint32(v0))+32:]))
	t3 := int32(load32(m.memory[uint32(v0+i32(36)):]))
	m.fn188(t2, t3)
	m.fn10(v0, i32(64), i32(8))
}
func (m *Module) fn1688(v0, v1, v2 int32) {
	var v3, v4, v5, v6, v7 int32
	v3 = v0 + i32(48)
	v4 = v0 + i32(32)
	t0 := m.fn936(v4)
	v5 = t0
	t1 := m.fn936(v4)
	v6 = t1 + v2
l4:
	{
		t2 := int32(load32(m.memory[int64(uint32(v0))+56:]))
		t3 := v6
		v7 = t2
		if uint32(t3) > uint32(v7) {
			t13 := v3
			v7 = v7 << 1
			p14 := i32(4)
			if uint32(v7) > uint32(i32(4)) {
				p14 = v7
			}
			m.fn482(t13, p14)
			goto l4
		}
		if uint32(v6) < uint32(v5) {
			m.fn151(v5, v6, v7, i32(1148236))
			panic("unreachable")
		}
		t4 := int32(load32(m.memory[int64(uint32(v0))+52:]))
		m.fn1689(t4+v5, v6-v5, v1, v2, i32(1148252))
		{
			t5 := int32(load32(m.memory[int64(uint32(v0))+44:]))
			v7 = t5
			t6 := int32(load32(m.memory[int64(uint32(v0))+40:]))
			t7 := v7
			v3 = t6
			if uint32(t7) < uint32(v3) {
				goto l2
			}
			t8 := v4
			v7 = v3 << 1
			p9 := i32(4)
			if uint32(v7) > uint32(i32(4)) {
				p9 = v7
			}
			m.fn939(t8, p9)
			t10 := int32(load32(m.memory[int64(uint32(v0))+40:]))
			v3 = t10
			t11 := int32(load32(m.memory[int64(uint32(v0))+44:]))
			v7 = t11
		}
	l2:
		if uint32(v7) >= uint32(v3) {
			m.fn158(v7, v3, i32(1148284))
			panic("unreachable")
		}
		store32(m.memory[int64(uint32(v0))+44:], uint32(v7+i32(1)))
		t12 := int32(load32(m.memory[int64(uint32(v0))+36:]))
		store32(m.memory[uint32(t12+v7<<2):], uint32(v6))
		return
	}
}
func (m *Module) fn1689(v0, v1, v2, v3, v4 int32) {
	if v1 != v3 {
		m.fn1668(v1, v3, v4)
		panic("unreachable")
	}
	if v1 == 0 {
		return
	}
	memory_copy(m.memory, uint32(v0), uint32(v2), uint32(v1))
}
func (m *Module) fn1690(v0, v1, v2, v3, v4, v5 int32) {
	var v6, v7, v8, v9 int32
	var v10 int64
	t0 := m.g0
	v6 = t0 - i32(32)
	m.g0 = v6
	v2 = v3 + v2
	if uint32(v2) >= uint32(v3) {
		goto l0
	}
	v3 = i32(0)
	goto l9
l0:
	v7 = i32(0)
	v8 = v6 + i32(20)
	{
		{
			t1 := int32(load32(m.memory[uint32(v1):]))
			t2 := int64(uint32(v5))
			t3 := v2
			v9 = t1
			v3 = v9 << 1
			p4 := v3
			if uint32(v2) > uint32(v3) {
				p4 = t3
			}
			v3 = p4
			t6 := v3
			p5 := i32(4)
			if v5 == i32(1) {
				p5 = i32(8)
			}
			v2 = p5
			p7 := v2
			if uint32(v3) > uint32(v2) {
				p7 = t6
			}
			v2 = p7
			v10 = t2 * int64(uint32(v2))
			if int32(int64(uint64(v10)>>32)) != 0 {
				goto l2
			}
			v3 = int32(v10)
			if uint32(v3) > uint32(i32(-0x80000000)-v4) {
				goto l2
			}
			{
				if v9 != 0 {
					goto l3
				}
				v5 = i32(0)
				v7 = v6 + i32(28)
				goto l4
			l3:
				t8 := int32(load32(m.memory[int64(uint32(v1))+4:]))
				v8 = t8
				store32(m.memory[int64(uint32(v6))+28:], uint32(v4))
				v5 = v9 * v5
				v7 = v6 + i32(24)
			}
		l4:
			store32(m.memory[uint32(v7):], uint32(v5))
			{
				{
					t9 := int32(load32(m.memory[int64(uint32(v6))+28:]))
					if t9 == 0 {
						goto l5
					}
					{
						t10 := int32(load32(m.memory[int64(uint32(v6))+24:]))
						v5 = t10
						if v5 != 0 {
							t12 := m.fn89(v8, v5, v4, v3)
							v5 = t12
							goto l7
						}
						m.fn1691(v6+i32(8), v4, v3)
						t11 := int32(load32(m.memory[int64(uint32(v6))+8:]))
						v5 = t11
						goto l7
					}
				}
			l5:
				m.fn1691(v6, v4, v3)
				t13 := int32(load32(m.memory[uint32(v6):]))
				v5 = t13
			}
		l7:
			if v5 != 0 {
				goto l8
			}
			store32(m.memory[int64(uint32(v6))+20:], uint32(v4))
			v8 = v6 + i32(16)
			v7 = v3
		}
	l2:
		store32(m.memory[uint32(v8):], uint32(v7))
		t14 := int32(load32(m.memory[int64(uint32(v6))+16:]))
		v5 = t14
		t15 := int32(load32(m.memory[int64(uint32(v6))+20:]))
		v3 = t15
		goto l9
	}
l8:
	store32(m.memory[uint32(v1):], uint32(v2))
	store32(m.memory[int64(uint32(v1))+4:], uint32(v5))
	v3 = i32(-1)
l9:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
	store32(m.memory[uint32(v0):], uint32(v3))
	m.g0 = v6 + i32(32)
}
func (m *Module) fn1691(v0, v1, v2 int32) {
	{
		if v2 == 0 {
			goto l0
		}
		t0 := m.fn248(v2, v1)
		v1 = t0
	}
l0:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v2))
	store32(m.memory[uint32(v0):], uint32(v1))
}
func (m *Module) fn1692(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7, v8, v9 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	v3 = i32(0)
l18:
	m.fn148(v2+i32(8), v3, v0, v1, i32(1280160))
	v4 = i32(0)
	{
		{
			t1 := int32(load32(m.memory[int64(uint32(v2))+8:]))
			v5 = t1
			v6 = (i32(0) - v5) & i32(3)
			t2 := int32(load32(m.memory[int64(uint32(v2))+12:]))
			t3 := v6 | i32(8)
			v7 = t2
			if uint32(t3) > uint32(v7) {
				goto l0
			}
			v4 = i32(0)
		l7:
			if v6 != v4 {
				t9 := int32(int8(m.memory[uint32(v5+v4)]))
				v9 = t9
				if v9 < i32(0) {
					goto l5
				}
				v4 = v4 + i32(1)
				goto l7
			}
			v8 = v7 + i32(-8)
			v4 = v6
		l6:
			{
				v9 = v5 + v4
				t4 := int32(load32(m.memory[uint32(v9+i32(4)):]))
				v6 = t4 & i32(-2139062144)
				t5 := int32(load32(m.memory[uint32(v9):]))
				t6 := v6
				v9 = t5 & i32(-2139062144)
				if t6|v9 == 0 {
					v4 = v4 + i32(8)
					if uint32(v4) <= uint32(v8) {
						goto l6
					}
					goto l0
				}
				if v9 != 0 {
					goto l3
				}
				v9 = int32(uint32(int32(bits.TrailingZeros32(uint32(v6))))>>3) + i32(4)
				goto l4
			l3:
				v9 = int32(uint32(int32(bits.TrailingZeros32(uint32(v9)))) >> 3)
			l4:
				t7 := v5
				v4 = v9 + v4
				t8 := int32(m.memory[uint32(t7+v4)])
				v9 = t8
				goto l5
			}
		}
	l0:
		p10 := v4
		if uint32(v7) > uint32(v4) {
			p10 = v7
		}
		v6 = p10
	l10:
		if v6 != v4 {
			t11 := int32(int8(m.memory[uint32(v5+v4)]))
			v9 = t11
			if v9 < i32(0) {
				goto l5
			}
			v4 = v4 + i32(1)
			goto l10
		}
		v5 = v1
		goto l9
	}
l5:
	v4 = v4 + v3
	if uint32(v4+i32(4)) > uint32(v1) {
		goto l21
	}
l17:
	{
		v5 = v4
		if uint32((v9+i32(62))&i32(255)) < uint32(i32(30)) {
			goto l12
		}
		v4 = v9 & i32(255)
		if uint32(v4) > uint32(i32(239)) {
			v9 = v5 + v0
			t17 := int32(m.memory[uint32(v9+i32(3))])
			t18 := int32(m.memory[int64(uint32(v4))+1280320])
			t19 := int32(m.memory[uint32(v9+i32(1))])
			t20 := int32(m.memory[int64(uint32(t19))+1280192])
			t21 := int32(m.memory[uint32(v9+i32(2))])
			if t17&i32(192)<<2|(t18&t20|int32(uint32(t21)>>6)) != i32(514) {
				goto l9
			}
			v4 = v5 + i32(4)
			if uint32(v5+i32(8)) > uint32(v1) {
				goto l21
			}
			t22 := int32(int8(m.memory[uint32(v0+v4)]))
			v9 = t22
			if v9 < i32(0) {
				goto l17
			}
			v3 = v5 + i32(5)
			goto l18
		}
		v4 = v5
	l16:
		{
			t12 := int32(m.memory[int64(uint32(v9&i32(255)))+1280320])
			v9 = v0 + v4
			t13 := int32(m.memory[uint32(v9+i32(1))])
			t14 := int32(m.memory[int64(uint32(t13))+1280192])
			t15 := int32(m.memory[uint32(v9+i32(2))])
			if t12&t14|int32(uint32(t15)>>6) != i32(2) {
				goto l14
			}
			if uint32(v4+i32(7)) <= uint32(v1) {
				goto l15
			}
			v4 = v4 + i32(3)
			goto l21
		l15:
			v4 = v4 + i32(3)
			t16 := int32(m.memory[uint32(v9+i32(3))])
			v5 = t16
			v9 = int32(int8(v5))
			if v5&i32(240) == i32(224) {
				goto l16
			}
		}
		if v9 < i32(0) {
			goto l17
		}
		v3 = v4 + i32(-3) + i32(4)
		goto l18
	l12:
		t23 := int32(int8(m.memory[uint32(v5+v0+i32(1))]))
		if t23 >= i32(-64) {
			goto l9
		}
		v4 = v5 + i32(2)
		if uint32(v5+i32(6)) > uint32(v1) {
			goto l21
		}
		t24 := int32(int8(m.memory[uint32(v0+v4)]))
		v9 = t24
		if v9 <= i32(-1) {
			goto l17
		}
	}
	v3 = v5 + i32(3)
	goto l18
l21:
	if uint32(v4) >= uint32(v1) {
		goto l14
	}
	{
		v6 = v0 + v4
		t25 := int32(int8(m.memory[uint32(v6)]))
		v9 = t25
		if v9 > i32(-1) {
			v4 = v4 + i32(1)
			goto l21
		}
		{
			if uint32((v9+i32(62))&i32(255)) < uint32(i32(30)) {
				v9 = v4 + i32(2)
				if uint32(v9) > uint32(v1) {
					goto l14
				}
				v5 = v4
				v4 = v9
				t31 := int32(int8(m.memory[int64(uint32(v6))+1]))
				if t31 < i32(-64) {
					goto l21
				}
				goto l9
			}
			if uint32(v9) >= uint32(i32(-16)) {
				goto l14
			}
			v5 = v4 + i32(3)
			if uint32(v5) > uint32(v1) {
				goto l14
			}
			t26 := int32(m.memory[int64(uint32(v9&i32(255)))+1280320])
			t27 := int32(m.memory[int64(uint32(v6))+1])
			t28 := int32(m.memory[int64(uint32(t27))+1280192])
			t29 := int32(m.memory[int64(uint32(v6))+2])
			p30 := v4
			if t26&t28|int32(uint32(t29)>>6) == i32(2) {
				p30 = v5
			}
			v5 = p30
			goto l9
		}
	}
l14:
	v5 = v4
l9:
	m.g0 = v2 + i32(16)
	return v5
}
func (m *Module) fn1693(v0, v1 int32) int32 {
	var v2, v3, v4, v5 int32
	v2 = i32(0)
	{
		v3 = (i32(0) - v0) & i32(3)
		if uint32(v3|i32(8)) > uint32(v1) {
			goto l0
		}
		v2 = i32(0)
	l8:
		if v3 != v2 {
			goto l1
		}
		v4 = v1 + i32(-8)
		v2 = v3
	l6:
		{
			v3 = v0 + v2
			t0 := int32(load32(m.memory[uint32(v3+i32(4)):]))
			v5 = t0 & i32(-2139062144)
			t1 := int32(load32(m.memory[uint32(v3):]))
			t2 := v5
			v3 = t1 & i32(-2139062144)
			if t2|v3 == 0 {
				v2 = v2 + i32(8)
				if uint32(v2) <= uint32(v4) {
					goto l6
				}
				goto l0
			}
			if v3 != 0 {
				goto l3
			}
			v0 = int32(uint32(int32(bits.TrailingZeros32(uint32(v5))))>>3) + i32(4)
			goto l4
		l3:
			v0 = int32(uint32(int32(bits.TrailingZeros32(uint32(v3)))) >> 3)
		l4:
			v2 = v0 + v2
			goto l5
		}
	l1:
		{
			if v1 == v2 {
				goto l7
			}
			t3 := int32(int8(m.memory[uint32(v0+v2)]))
			if t3 < i32(0) {
				goto l5
			}
			v2 = v2 + i32(1)
			goto l8
		}
	l7:
		m.fn158(v1, v1, i32(1280144))
		panic("unreachable")
	l0:
		p4 := v1
		if uint32(v2) > uint32(v1) {
			p4 = v2
		}
		v3 = p4
	l10:
		if v3 != v2 {
			t5 := int32(int8(m.memory[uint32(v0+v2)]))
			if t5 < i32(0) {
				goto l5
			}
			v2 = v2 + i32(1)
			goto l10
		}
		return v1
	}
l5:
	return v2
}
func (m *Module) fn1694(v0, v1, v2 int32) {
	var v3 int32
	v3 = i32(9)
	switch v1 & i32(255) {
	case 12:
		goto l11
	default:
		store32(m.memory[int64(uint32(v0))+4:], uint32(v2))
		v3 = i32(0)
		goto l11
	case 1:
		store16(m.memory[int64(uint32(v0))+16:], uint16(i32(49024)))
		store32(m.memory[int64(uint32(v0))+12:], uint32(i32(0)))
		store64(m.memory[int64(uint32(v0))+4:], uint64(i64(0)))
		v3 = i32(1)
		goto l11
	case 2, 3:
		store32(m.memory[int64(uint32(v0))+9:], uint32(i32(0)))
		m.memory[int64(uint32(v0))+7] = byte(i32(0))
		m.memory[int64(uint32(v0))+5] = byte(i32(0))
		m.memory[int64(uint32(v0))+3] = byte(i32(0))
		m.memory[int64(uint32(v0))+1] = byte(i32(0))
		v3 = i32(2)
		goto l11
	case 4:
		m.memory[int64(uint32(v0))+1] = byte(i32(0))
		v3 = i32(3)
		goto l11
	case 5:
		m.memory[int64(uint32(v0))+1] = byte(i32(0))
		v3 = i32(4)
		goto l11
	case 6:
		m.memory[int64(uint32(v0))+5] = byte(i32(0))
		store32(m.memory[int64(uint32(v0))+1:], uint32(i32(0)))
		v3 = i32(5)
		goto l11
	case 7:
		m.memory[int64(uint32(v0))+1] = byte(i32(0))
		v3 = i32(6)
		goto l11
	case 8:
		m.memory[int64(uint32(v0))+1] = byte(i32(0))
		v3 = i32(7)
		goto l11
	case 9:
		m.memory[int64(uint32(v0))+1] = byte(i32(0))
		v3 = i32(8)
		goto l11
	case 10:
		store32(m.memory[int64(uint32(v0))+4:], uint32(i32(65536)))
		m.memory[int64(uint32(v0))+2] = byte(i32(0))
		goto l12
	case 11:
		store32(m.memory[int64(uint32(v0))+4:], uint32(i32(0)))
		m.memory[int64(uint32(v0))+2] = byte(i32(0))
	}
l12:
	v3 = i32(10)
l11:
	m.memory[uint32(v0)] = byte(v3)
}
func (m *Module) fn1695(v0, v1 int32) {
	t0 := int32(m.memory[uint32(v1)])
	t1 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	m.fn1694(v0, t0, t1)
	m.memory[int64(uint32(v0))+24] = byte(i32(9))
	store32(m.memory[int64(uint32(v0))+20:], uint32(v1))
}
func (m *Module) fn1696(v0, v1, v2 int32) {
	var v3, v4 int32
	var v5 int64
	var v6, v7 int32
	t0 := m.g0
	v3 = t0 - i32(48)
	m.g0 = v3
	{
		{
			t1 := int32(m.memory[int64(uint32(v1))+24])
			switch t1 {
			default:
				v4 = i32(0)
				if uint32(v2) > uint32(i32(-4)) {
					goto l8
				}
				v5 = int64(uint32(int32(uint32(v2+i32(1))>>1))) * i64(3)
				if int32(int64(uint64(v5)>>32)) != 0 {
					goto l8
				}
				v6 = int32(v5)
				if v6 != i32(-1) {
					goto l7
				}
				goto l8
			case 1, 2, 3, 9:
				m.fn1703(v3+i32(16), v1, v2)
				t2 := int32(load32(m.memory[int64(uint32(v3))+20:]))
				v6 = t2
				t3 := int32(load32(m.memory[int64(uint32(v3))+16:]))
				v4 = t3
				goto l8
			case 4, 5:
				v4 = i32(0)
				if uint32(v2) > uint32(i32(-3)) {
					goto l8
				}
				v6 = v2 + i32(5)
				t4 := v6
				v2 = v2 + i32(2)
				if uint32(t4) < uint32(v2) {
					goto l8
				}
				t5 := int32(load32(m.memory[int64(uint32(v1))+20:]))
				if t5 == i32(1148960) {
					goto l9
				}
				m.fn1703(v3+i32(24), v1, v2)
				t6 := int32(load32(m.memory[int64(uint32(v3))+24:]))
				if t6 != i32(1) {
					goto l8
				}
				t7 := int32(load32(m.memory[int64(uint32(v3))+28:]))
				v1 = t7
				p8 := v6
				if uint32(v1) > uint32(v6) {
					p8 = v1
				}
				v6 = p8
				goto l9
			case 6, 7:
				v4 = i32(0)
				if uint32(v2) > uint32(i32(-3)) {
					goto l8
				}
				v6 = v2 + i32(3)
				t9 := v6
				v2 = v2 + i32(2)
				if uint32(t9) < uint32(v2) {
					goto l8
				}
				v5 = int64(uint32(int32(uint32(v6)>>1))) * i64(3)
				if int32(int64(uint64(v5)>>32)) != 0 {
					goto l8
				}
				v6 = int32(v5) + i32(1)
				if v6 == 0 {
					goto l8
				}
				t10 := int32(load32(m.memory[int64(uint32(v1))+20:]))
				v7 = t10
				if v7 == i32(1153092) {
					goto l9
				}
				if v7 == i32(1153064) {
					goto l9
				}
				m.fn1703(v3+i32(32), v1, v2)
				t11 := int32(load32(m.memory[int64(uint32(v3))+32:]))
				if t11 != i32(1) {
					goto l8
				}
				t12 := int32(load32(m.memory[int64(uint32(v3))+36:]))
				v1 = t12
				p13 := v6
				if uint32(v1) > uint32(v6) {
					p13 = v1
				}
				v6 = p13
				goto l9
			case 8:
				if uint32(v2) <= uint32(i32(-3)) {
					m.fn1703(v3+i32(40), v1, v2+i32(2))
					t20 := int32(load32(m.memory[int64(uint32(v3))+44:]))
					v6 = t20
					t21 := int32(load32(m.memory[int64(uint32(v3))+40:]))
					v4 = t21
					goto l8
				}
				v4 = i32(0)
				goto l8
			case 10:
				m.fn256(i32(1155236), i32(41), i32(1155328))
				panic("unreachable")
			}
		}
	l7:
		v6 = v6 + i32(1)
		t14 := v6
		v7 = v2 + i32(3)
		p15 := v7
		if uint32(v6) > uint32(v7) {
			p15 = t14
		}
		v6 = p15
		t16 := int32(load32(m.memory[int64(uint32(v1))+20:]))
		v7 = t16
		if v7 == i32(1153064) {
			goto l9
		}
		if v7 == i32(1148960) {
			goto l9
		}
		if v7 == i32(1153092) {
			goto l9
		}
		m.fn1703(v3+i32(8), v1, v2)
		t17 := int32(load32(m.memory[int64(uint32(v3))+8:]))
		if t17 != i32(1) {
			goto l8
		}
		t18 := int32(load32(m.memory[int64(uint32(v3))+12:]))
		v1 = t18
		p19 := v6
		if uint32(v1) > uint32(v6) {
			p19 = v1
		}
		v6 = p19
	}
l9:
	v4 = i32(1)
	goto l8
l8:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v6))
	store32(m.memory[uint32(v0):], uint32(v4))
	m.g0 = v3 + i32(48)
}
func (m *Module) fn1697(v0, v1, v2 int32) {
	var v3 int32
	v3 = i32(1)
	{
		if v1 == i32(1) {
			goto l0
		}
		v3 = i32(0)
		goto l1
	l0:
		p0 := i32(1)
		if uint32(v2) > uint32(i32(1)) {
			p0 = i32_shr_u(i32(-1), int32(bits.LeadingZeros32(uint32(v2+i32(-1))))) + i32(1)
		}
		v1 = p0
	}
l1:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
	store32(m.memory[uint32(v0):], uint32(v3))
}
func (m *Module) fn1698(v0, v1, v2 int32) {
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
			if int32(int64(uint64(v5)>>32)) != 0 {
				goto l6
			}
			v6 = int32(v5)
			if uint32(v6) >= uint32(i32(-3)) {
				goto l6
			}
			v7 = int32(uint32(v2+i32(1))>>1)*i32(3) + i32(1)
			t2 := v7
			v6 = v6 + i32(3)
			p3 := v6
			if uint32(v7) > uint32(v6) {
				p3 = t2
			}
			v6 = p3
			t4 := int32(load32(m.memory[int64(uint32(v1))+20:]))
			v7 = t4
			if v7 == i32(1153064) {
				goto l7
			}
			if v7 == i32(1148960) {
				goto l7
			}
			if v7 == i32(1153092) {
				goto l7
			}
			m.fn1704(v3+i32(8), v1, v2)
			t5 := int32(load32(m.memory[int64(uint32(v3))+8:]))
			if t5 != i32(1) {
				goto l8
			}
			t6 := int32(load32(m.memory[int64(uint32(v3))+12:]))
			v1 = t6
			p7 := v6
			if uint32(v1) > uint32(v6) {
				p7 = v1
			}
			v6 = p7
			goto l7
		case 1, 2, 3, 9:
			m.fn1704(v3+i32(16), v1, v2)
			t8 := int32(load32(m.memory[int64(uint32(v3))+20:]))
			v6 = t8
			t9 := int32(load32(m.memory[int64(uint32(v3))+16:]))
			v4 = t9
			goto l8
		case 4, 5:
			v4 = i32(0)
			if uint32(v2) > uint32(i32(-3)) {
				goto l6
			}
			v2 = v2 + i32(2)
			v5 = int64(uint32(v2)) * i64(3)
			if int32(int64(uint64(v5)>>32)) != 0 {
				goto l6
			}
			v6 = int32(v5)
			if uint32(v6) > uint32(i32(-4)) {
				goto l6
			}
			v6 = v6 + i32(3)
			t10 := int32(load32(m.memory[int64(uint32(v1))+20:]))
			if t10 == i32(1148960) {
				goto l7
			}
			m.fn1704(v3+i32(24), v1, v2)
			t11 := int32(load32(m.memory[int64(uint32(v3))+24:]))
			if t11 != i32(1) {
				goto l8
			}
			t12 := int32(load32(m.memory[int64(uint32(v3))+28:]))
			v1 = t12
			p13 := v6
			if uint32(v1) > uint32(v6) {
				p13 = v1
			}
			v6 = p13
			goto l7
		case 6, 7:
			v4 = i32(0)
			if uint32(v2) <= uint32(i32(-3)) {
				v6 = v2 + i32(3)
				t14 := v6
				v2 = v2 + i32(2)
				if uint32(t14) < uint32(v2) {
					goto l6
				}
				v5 = int64(uint32(int32(uint32(v6)>>1))) * i64(3)
				if int32(int64(uint64(v5)>>32)) != 0 {
					goto l6
				}
				v6 = int32(v5) + i32(1)
				if v6 == 0 {
					goto l6
				}
				t15 := int32(load32(m.memory[int64(uint32(v1))+20:]))
				v7 = t15
				if v7 == i32(1153092) {
					goto l7
				}
				if v7 == i32(1153064) {
					goto l7
				}
				m.fn1704(v3+i32(32), v1, v2)
				t16 := int32(load32(m.memory[int64(uint32(v3))+32:]))
				if t16 != i32(1) {
					goto l8
				}
				t17 := int32(load32(m.memory[int64(uint32(v3))+36:]))
				v1 = t17
				p18 := v6
				if uint32(v1) > uint32(v6) {
					p18 = v1
				}
				v6 = p18
				goto l7
			}
			goto l6
		case 8:
			if uint32(v2) <= uint32(i32(-3)) {
				m.fn1704(v3+i32(40), v1, v2+i32(2))
				t19 := int32(load32(m.memory[int64(uint32(v3))+44:]))
				v6 = t19
				t20 := int32(load32(m.memory[int64(uint32(v3))+40:]))
				v4 = t20
				goto l8
			}
			v4 = i32(0)
			goto l8
		case 10:
			m.fn256(i32(1155236), i32(41), i32(1155280))
			panic("unreachable")
		}
	}
l6:
	goto l8
l7:
	v4 = i32(1)
l8:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v6))
	store32(m.memory[uint32(v0):], uint32(v4))
	m.g0 = v3 + i32(48)
}
func (m *Module) fn1699(v0, v1 int32) {
	var v2, v3 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	m.fn1705(v2+i32(4), v1)
	t1 := int32(load32(m.memory[int64(uint32(v2))+8:]))
	v1 = t1
	{
		t2 := int32(load32(m.memory[int64(uint32(v2))+4:]))
		if t2 != i32(1) {
			goto l0
		}
		t3 := int32(load32(m.memory[int64(uint32(v2))+12:]))
		m.fn2(v1, t3)
		panic("unreachable")
	}
l0:
	t4 := int32(load32(m.memory[int64(uint32(v2))+12:]))
	v3 = t4
	store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
	store32(m.memory[uint32(v0):], uint32(v1))
	m.g0 = v2 + i32(16)
}
func (m *Module) fn1700(v0, v1 int32) int32 {
	var v2, v3, v4 int32
	v2 = i32(0)
l2:
	{
		t0 := v1
		v3 = v2
		if t0 != v3 {
			{
				t1 := int32(int8(m.memory[uint32(v0+v3)]))
				v4 = t1
				if v4 < i32(0) {
					goto l1
				}
				v2 = v3 + i32(1)
				v4 = v4 & i32(255)
				if uint32(v4) > uint32(i32(27)) {
					goto l2
				}
				if i32_shl(i32(1), v4)&i32(0x800c000) == 0 {
					goto l2
				}
			}
		l1:
			return v3
		}
		return v1
	}
}
func (m *Module) fn1701(v0, v1, v2, v3, v4, v5 int32) {
	var v6, v7 int32
	t0 := int32(m.memory[int64(uint32(v1))+24])
	v6 = t0
	v7 = i32(0)
l32:
	{
		switch v6 & i32(255) {
		case 3:
			goto l3
		case 4:
			if uint32(v7) < uint32(v3) {
				t11 := int32(m.memory[uint32(v2+v7)])
				if t11 != i32(187) {
					m.fn1707(v0, v1, v2, v3, v4, v5, v7, i32(239))
					return
				}
				v6 = i32(5)
				goto l19
			}
			m.fn1707(v0, v1, v2, v3, v4, v5, v7, i32(239))
			return
		case 5:
			if uint32(v7) < uint32(v3) {
				t9 := int32(m.memory[uint32(v2+v7)])
				if t9 != i32(191) {
					goto l23
				}
				m.memory[int64(uint32(v1))+24] = byte(i32(9))
				v6 = v7 + i32(1)
				{
					t10 := int32(load32(m.memory[int64(uint32(v1))+20:]))
					if t10 == i32(1148960) {
						goto l30
					}
					store16(m.memory[int64(uint32(v1))+16:], uint16(i32(49024)))
					store32(m.memory[int64(uint32(v1))+12:], uint32(i32(0)))
					store64(m.memory[int64(uint32(v1))+4:], uint64(i64(0)))
					m.memory[uint32(v1)] = byte(i32(1))
					store32(m.memory[int64(uint32(v1))+20:], uint32(i32(1148960)))
				}
			l30:
				m.fn1708(v0, v1, v2, v3, v4, v5, v6)
				return
			}
			goto l23
		case 6:
			if uint32(v7) < uint32(v3) {
				t7 := int32(m.memory[uint32(v2+v7)])
				if t7 != i32(255) {
					m.fn1707(v0, v1, v2, v3, v4, v5, v7, i32(254))
					return
				}
				m.memory[int64(uint32(v1))+24] = byte(i32(9))
				v6 = v7 + i32(1)
				{
					t8 := int32(load32(m.memory[int64(uint32(v1))+20:]))
					if t8 == i32(1153064) {
						goto l29
					}
					store32(m.memory[int64(uint32(v1))+4:], uint32(i32(65536)))
					m.memory[int64(uint32(v1))+2] = byte(i32(0))
					m.memory[uint32(v1)] = byte(i32(10))
					store32(m.memory[int64(uint32(v1))+20:], uint32(i32(1153064)))
				}
			l29:
				m.fn1708(v0, v1, v2, v3, v4, v5, v6)
				return
			}
			m.fn1707(v0, v1, v2, v3, v4, v5, v7, i32(254))
			return
		case 7:
			if uint32(v7) < uint32(v3) {
				t5 := int32(m.memory[uint32(v2+v7)])
				if t5 != i32(254) {
					m.fn1707(v0, v1, v2, v3, v4, v5, v7, i32(255))
					return
				}
				m.memory[int64(uint32(v1))+24] = byte(i32(9))
				v6 = v7 + i32(1)
				{
					t6 := int32(load32(m.memory[int64(uint32(v1))+20:]))
					if t6 == i32(1153092) {
						goto l27
					}
					store32(m.memory[int64(uint32(v1))+1:], uint32(i32(0)))
					m.memory[uint32(v1)] = byte(i32(10))
					store32(m.memory[int64(uint32(v1))+20:], uint32(i32(1153092)))
					store32(m.memory[int64(uint32(v1))+4:], uint32(i32(0)))
				}
			l27:
				m.fn1708(v0, v1, v2, v3, v4, v5, v6)
				return
			}
			m.fn1707(v0, v1, v2, v3, v4, v5, v7, i32(255))
			return
		case 9:
			m.fn1706(v0, v1, v2, v3, v4, v5)
			return
		case 10:
			m.fn256(i32(1155236), i32(41), i32(1155296))
			panic("unreachable")
		default:
			if v3 == 0 {
				goto l11
			}
			{
				t1 := int32(m.memory[uint32(v2)])
				v6 = t1
				switch v6 + i32(-254) {
				case 0:
					goto l12
				case 1:
					goto l13
				default:
					if v6 == i32(239) {
						goto l15
					}
					goto l16
				}
			}
		case 8:
			m.fn1707(v0, v1, v2, v3, v4, v5, i32(0), i32(187))
			return
		case 1:
			if v3 == 0 {
				store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
				store32(m.memory[uint32(v0):], uint32(i32(0)))
				m.memory[int64(uint32(v0))+4] = byte(i32(0))
				return
			}
			t2 := int32(m.memory[uint32(v2)])
			if t2 != i32(239) {
				goto l16
			}
			goto l15
		case 2:
			if v3 == 0 {
				store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
				store32(m.memory[uint32(v0):], uint32(i32(0)))
				m.memory[int64(uint32(v0))+4] = byte(i32(0))
				return
			}
			t3 := int32(m.memory[uint32(v2)])
			if t3 != i32(254) {
				goto l16
			}
		}
	l12:
		v6 = i32(6)
		goto l19
	l3:
		if v3 == 0 {
			store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
			store32(m.memory[uint32(v0):], uint32(i32(0)))
			m.memory[int64(uint32(v0))+4] = byte(i32(0))
			return
		}
		t4 := int32(m.memory[uint32(v2)])
		if t4 != i32(255) {
			goto l16
		}
	}
l13:
	v6 = i32(7)
	goto l19
l23:
	m.fn1709(v0, v1, v2, v3, v4, v5, v7)
	return
l15:
	v6 = i32(4)
l19:
	m.memory[int64(uint32(v1))+24] = byte(v6)
	v7 = v7 + i32(1)
	goto l32
l16:
	v6 = i32(9)
	m.memory[int64(uint32(v1))+24] = byte(i32(9))
	goto l32
l11:
	store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
	store32(m.memory[uint32(v0):], uint32(i32(0)))
	m.memory[int64(uint32(v0))+4] = byte(i32(0))
}
func (m *Module) fn1702(v0, v1 int32) {
	var v2, v3 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	{
		t1 := int32(load32(m.memory[uint32(v0):]))
		t2 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		t3 := v1
		v3 = t2
		if uint32(t3) <= uint32(t1-v3) {
			goto l0
		}
		m.fn1690(v2+i32(8), v0, v3, v1, i32(1), i32(1))
		t4 := int32(load32(m.memory[int64(uint32(v2))+8:]))
		v0 = t4
		if v0 == i32(-1) {
			goto l0
		}
		t5 := int32(load32(m.memory[int64(uint32(v2))+12:]))
		m.fn2(v0, t5)
		panic("unreachable")
	}
l0:
	m.g0 = v2 + i32(16)
}
func (m *Module) fn1703(v0, v1, v2 int32) {
	var v3, v4, v5 int32
	var v6 int64
	t0 := m.g0
	v3 = t0 - i32(32)
	m.g0 = v3
	v4 = i32(1)
	v5 = i32(3)
	{
		{
			t1 := int32(m.memory[uint32(v1)])
			switch t1 {
			case 8:
				goto l8
			default:
				v6 = int64(uint32(v2)) * i64(3)
				v5 = int32(v6)
				var p2 int32
				if int32(int64(uint64(v6)>>32)) == 0 {
					p2 = 1
				}
				v4 = p2
				goto l8
			case 1:
				t3 := int32(load32(m.memory[int64(uint32(v1))+8:]))
				t4 := int32(load32(m.memory[int64(uint32(v1))+12:]))
				p5 := i32(3)
				if t4 != 0 {
					p5 = t3 + i32(4)
				}
				v5 = p5 + v2
				var p6 int32
				if uint32(v5) >= uint32(v2) {
					p6 = 1
				}
				v4 = p6
				goto l8
			case 2:
				m.fn1714(v3+i32(8), v1+i32(1), v2)
				t7 := int32(load32(m.memory[int64(uint32(v3))+12:]))
				v5 = t7
				t8 := int32(load32(m.memory[int64(uint32(v3))+8:]))
				v4 = t8
				goto l8
			case 3:
				t9 := int32(m.memory[int64(uint32(v1))+1])
				v5 = v2 + t9&i32(255)
				var p10 int32
				if v5 != i32(0x7fffffff) {
					p10 = 1
				}
				var p11 int32
				if uint32(v5) >= uint32(v2) {
					p11 = 1
				}
				var p12 int32
				if v5 > i32(-1) {
					p12 = 1
				}
				v4 = p10 & (p11 & p12)
				v5 = v5<<1 + i32(2)
				goto l8
			case 4:
				v4 = i32(0)
				t13 := int32(m.memory[int64(uint32(v1))+1])
				t14 := v2
				var p15 int32
				if t13 != i32(0) {
					p15 = 1
				}
				v1 = t14 + p15
				if uint32(v1) < uint32(v2) {
					goto l11
				}
				if v1 == i32(-1) {
					goto l11
				}
				v2 = int32(uint32(v1+i32(1)) >> 1)
				v1 = v2 + v1
				if uint32(v1) < uint32(v2) {
					goto l8
				}
				var p16 int32
				if uint32(v1) < uint32(i32(-2)) {
					p16 = 1
				}
				v4 = p16
				v5 = v1 + i32(2)
				goto l8
			case 5:
				m.fn1715(v3+i32(16), v1+i32(1), v2)
				t17 := int32(load32(m.memory[int64(uint32(v3))+20:]))
				v5 = t17
				t18 := int32(load32(m.memory[int64(uint32(v3))+16:]))
				v4 = t18
				goto l8
			case 6:
				t19 := int32(m.memory[int64(uint32(v1))+1])
				m.fn1716(v3+i32(24), t19, v2)
				t20 := int32(load32(m.memory[int64(uint32(v3))+28:]))
				v5 = t20
				t21 := int32(load32(m.memory[int64(uint32(v3))+24:]))
				v4 = t21
				goto l8
			case 7:
				v4 = i32(0)
				t22 := int32(m.memory[int64(uint32(v1))+1])
				v1 = v2 + t22
				if uint32(v1) < uint32(v2) {
					goto l11
				}
				if v1 == i32(-1) {
					goto l11
				}
				v2 = int32(uint32(v1+i32(1)) >> 1)
				v1 = v2 + v1
				if uint32(v1) < uint32(v2) {
					goto l8
				}
				var p23 int32
				if uint32(v1) < uint32(i32(-2)) {
					p23 = 1
				}
				v4 = p23
				v5 = v1 + i32(2)
				goto l8
			case 9:
				v6 = int64(uint32(v2)) * i64(3)
				v5 = int32(v6)
				var p24 int32
				if int32(int64(uint64(v6)>>32)) == 0 {
					p24 = 1
				}
				v4 = p24
				goto l8
			case 10:
				v4 = i32(0)
				t25 := int32(m.memory[int64(uint32(v1))+2])
				p26 := i32(1)
				if t25 != 0 {
					p26 = i32(2)
				}
				t27 := int32(load16(m.memory[int64(uint32(v1))+4:]))
				t29 := p26 + v2
				p28 := i32(0)
				if t27 != 0 {
					p28 = i32(2)
				}
				v1 = t29 + p28
				if uint32(v1) >= uint32(v2) {
					goto l12
				}
			}
		}
	l11:
		goto l8
	l12:
		v6 = int64(uint32(int32(uint32(v1)>>1))) * i64(3)
		if int32(int64(uint64(v6)>>32)) != 0 {
			goto l8
		}
		v5 = int32(v6)
		var p30 int32
		if v5 != i32(-1) {
			p30 = 1
		}
		v4 = p30
		v5 = v5 + i32(1)
	}
l8:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
	store32(m.memory[uint32(v0):], uint32(v4))
	m.g0 = v3 + i32(32)
}
func (m *Module) fn1704(v0, v1, v2 int32) {
	var v3, v4, v5 int32
	var v6 int64
	t0 := m.g0
	v3 = t0 - i32(32)
	m.g0 = v3
	v4 = i32(1)
	v5 = i32(3)
	{
		{
			t1 := int32(m.memory[uint32(v1)])
			switch t1 {
			case 8:
				goto l8
			default:
				v6 = int64(uint32(v2)) * i64(3)
				v5 = int32(v6)
				var p2 int32
				if int32(int64(uint64(v6)>>32)) == 0 {
					p2 = 1
				}
				v4 = p2
				goto l8
			case 1:
				v4 = i32(0)
				t3 := int32(load32(m.memory[int64(uint32(v1))+8:]))
				t4 := int32(load32(m.memory[int64(uint32(v1))+12:]))
				p5 := i32(0)
				if t4 != 0 {
					p5 = t3 + i32(1)
				}
				v5 = p5
				v1 = v5 + v2
				if uint32(v1) < uint32(v5) {
					goto l11
				}
				v6 = int64(uint32(v1)) * i64(3)
				if int32(int64(uint64(v6)>>32)) != 0 {
					goto l8
				}
				v5 = int32(v6)
				var p6 int32
				if uint32(v5) < uint32(i32(-3)) {
					p6 = 1
				}
				v4 = p6
				v5 = v5 + i32(3)
				goto l8
			case 2:
				m.fn1714(v3+i32(8), v1+i32(1), v2)
				t7 := int32(load32(m.memory[int64(uint32(v3))+12:]))
				v5 = t7
				t8 := int32(load32(m.memory[int64(uint32(v3))+8:]))
				v4 = t8
				goto l8
			case 3:
				v4 = i32(0)
				t9 := int32(m.memory[int64(uint32(v1))+1])
				v1 = v2 + t9
				if uint32(v1) < uint32(v2) {
					goto l11
				}
				v6 = int64(uint32(v1)) * i64(3)
				if int32(int64(uint64(v6)>>32)) != 0 {
					goto l8
				}
				v5 = int32(v6)
				var p10 int32
				if uint32(v5) < uint32(i32(-3)) {
					p10 = 1
				}
				v4 = p10
				v5 = v5 + i32(3)
				goto l8
			case 4:
				v4 = i32(0)
				t11 := int32(m.memory[int64(uint32(v1))+1])
				t12 := v2
				var p13 int32
				if t11 != i32(0) {
					p13 = 1
				}
				v1 = t12 + p13
				if uint32(v1) < uint32(v2) {
					goto l8
				}
				v6 = int64(uint32(v1)) * i64(3)
				v5 = int32(v6)
				var p14 int32
				if int32(int64(uint64(v6)>>32)) == 0 {
					p14 = 1
				}
				v4 = p14
				goto l8
			case 5:
				m.fn1715(v3+i32(16), v1+i32(1), v2)
				t15 := int32(load32(m.memory[int64(uint32(v3))+20:]))
				v5 = t15
				t16 := int32(load32(m.memory[int64(uint32(v3))+16:]))
				v4 = t16
				goto l8
			case 6:
				t17 := int32(m.memory[int64(uint32(v1))+1])
				m.fn1716(v3+i32(24), t17, v2)
				t18 := int32(load32(m.memory[int64(uint32(v3))+28:]))
				v5 = t18
				t19 := int32(load32(m.memory[int64(uint32(v3))+24:]))
				v4 = t19
				goto l8
			case 7:
				t20 := int32(m.memory[int64(uint32(v1))+1])
				v4 = v2 + t20
				if uint32(v4) >= uint32(v2) {
					v6 = int64(uint32(v4)) * i64(3)
					v5 = int32(v6)
					var p21 int32
					if int32(int64(uint64(v6)>>32)) == 0 {
						p21 = 1
					}
					v4 = p21
					goto l8
				}
				v4 = i32(0)
				goto l8
			case 9:
				v6 = int64(uint32(v2)) * i64(3)
				v5 = int32(v6)
				var p22 int32
				if int32(int64(uint64(v6)>>32)) == 0 {
					p22 = 1
				}
				v4 = p22
				goto l8
			case 10:
				v4 = i32(0)
				t23 := int32(m.memory[int64(uint32(v1))+2])
				p24 := i32(1)
				if t23 != 0 {
					p24 = i32(2)
				}
				t25 := int32(load16(m.memory[int64(uint32(v1))+4:]))
				t27 := p24 + v2
				p26 := i32(0)
				if t25 != 0 {
					p26 = i32(2)
				}
				v1 = t27 + p26
				if uint32(v1) >= uint32(v2) {
					goto l13
				}
			}
		}
	l11:
		goto l8
	l13:
		v6 = int64(uint32(int32(uint32(v1)>>1))) * i64(3)
		if int32(int64(uint64(v6)>>32)) != 0 {
			goto l8
		}
		v5 = int32(v6)
		var p28 int32
		if v5 != i32(-1) {
			p28 = 1
		}
		v4 = p28
		v5 = v5 + i32(1)
	}
l8:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
	store32(m.memory[uint32(v0):], uint32(v4))
	m.g0 = v3 + i32(32)
}
func (m *Module) fn1705(v0, v1 int32) {
	var v2, v3, v4 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	if v1 < i32(0) {
		store32(m.memory[int64(uint32(v0))+4:], uint32(i32(0)))
		v3 = i32(1)
		goto l2
	}
	if v1 != 0 {
		goto l1
	}
	store64(m.memory[int64(uint32(v0))+4:], uint64(i64(0x100000000)))
	v3 = i32(0)
	goto l2
l1:
	v3 = i32(1)
	m.fn1691(v2+i32(8), i32(1), v1)
	{
		t1 := int32(load32(m.memory[int64(uint32(v2))+8:]))
		v4 = t1
		if v4 == 0 {
			goto l3
		}
		store32(m.memory[int64(uint32(v0))+8:], uint32(v4))
		store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
		v3 = i32(0)
		goto l2
	}
l3:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v1))
	store32(m.memory[int64(uint32(v0))+4:], uint32(i32(1)))
l2:
	store32(m.memory[uint32(v0):], uint32(v3))
	m.g0 = v2 + i32(16)
}
func (m *Module) fn1706(v0, v1, v2, v3, v4, v5 int32) {
	m.fn1710(v0, v1, v2, v3, v4, v5, i32(1))
	{
		t0 := int32(m.memory[int64(uint32(v0))+4])
		if t0 != 0 {
			return
		}
		m.memory[int64(uint32(v1))+24] = byte(i32(10))
	}
}
