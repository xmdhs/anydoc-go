package core

import (
	"math/bits"
)

func (m *Module) fn897(v0, v1 int32) {
	var v2, v3 int32
	{
		if uint32(v1) >= uint32(i32(108)) {
			goto l0
		}
		v2 = i32(2)
		v3 = v1
		goto l1
	l0:
		v2 = i32(8)
		v3 = v1 + i32(-119)
		if uint32(v3) >= uint32(i32(8)) {
			goto l2
		}
		v2 = i32(5)
		goto l1
	l2:
		v3 = v1 + i32(-135)
		if uint32(v3) < uint32(i32(7)) {
			goto l1
		}
		v3 = v1 + i32(-153)
		if uint32(v3) >= uint32(i32(15)) {
			goto l3
		}
		v2 = i32(11)
		goto l1
	l3:
		v3 = v1 + i32(-175)
		if uint32(v3) >= uint32(i32(8)) {
			goto l4
		}
		v2 = i32(14)
		goto l1
	l4:
		if v1 != i32(187) {
			goto l5
		}
		v2 = i32(17)
		v3 = i32(0)
		goto l1
	l5:
		v3 = v1 + i32(-658)
		if uint32(v3) >= uint32(i32(32)) {
			goto l6
		}
		v2 = i32(20)
		goto l1
	l6:
		v2 = i32(23)
		v3 = v1 + i32(-1159)
		if uint32(v3) < uint32(i32(23)) {
			goto l1
		}
		v3 = v1 + i32(-1190)
		if uint32(v3) >= uint32(i32(30)) {
			goto l7
		}
		v2 = i32(26)
		goto l1
	l7:
		v3 = v1 + i32(-10736)
		if uint32(v3) >= uint32(i32(8)) {
			goto l8
		}
		v2 = i32(29)
		goto l1
	l8:
		v3 = v1 + i32(-8644)
		if uint32(v3) <= uint32(i32(3)) {
			goto l9
		}
		v1 = i32(0)
		goto l10
	l9:
		v2 = i32(32)
	l1:
		v1 = i32(1)
		t0 := int32(load16(m.memory[int64(uint32(v2<<1))+1242356:]))
		v3 = v3 + t0
		if uint32(v3) > uint32(i32(239)) {
			m.fn32(v3, i32(240), i32(1241928))
			panic("unreachable")
		}
		t1 := int32(load16(m.memory[int64(uint32(v3<<1))+1226932:]))
		v3 = t1
	}
l10:
	store16(m.memory[int64(uint32(v0))+2:], uint16(v3))
	store16(m.memory[uint32(v0):], uint16(v1))
}
func (m *Module) fn898(v0, v1 int32) {
	var v2, v3, v4 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	{
		if v0 == 0 {
			goto l0
		}
		if v1 == 0 {
			goto l0
		}
		t1 := int32(load32(m.memory[uint32(v0):]))
		v0 = t1
		if v0 == 0 {
			goto l0
		}
		if uint32(v0) >= uint32(i32(0x7fffffc1)) {
			m.fn41(i32(1284720), i32(43), v2+i32(15), i32(1284832), i32(1284848))
			panic("unreachable")
		}
		t2 := int32(load32(m.memory[uint32(v1+i32(-4)):]))
		v3 = t2
		v4 = v3 & i32(-8)
		t3 := v4
		v3 = v3 & i32(3)
		p4 := i32(8)
		if v3 != 0 {
			p4 = i32(4)
		}
		if uint32(t3) < uint32(p4+v0) {
			m.fn3(i32(1274224), i32(46), i32(1274272))
			panic("unreachable")
		}
		if v3 == 0 {
			goto l3
		}
		if uint32(v4) > uint32(v0+i32(39)) {
			m.fn3(i32(1274288), i32(46), i32(1274336))
			panic("unreachable")
		}
	l3:
		m.fn1(v1)
	}
l0:
	m.g0 = v2 + i32(16)
}
func (m *Module) fn899(v0, v1, v2 int32) int32 {
	var v3 int32
	t0 := m.g0
	v3 = t0 - i32(16)
	m.g0 = v3
	{
		v1 = v2 * v1
		if v1 != 0 {
			goto l0
		}
		v1 = i32(0)
		goto l1
	l0:
		if uint32(v1) < uint32(i32(0x7fffffc1)) {
			goto l2
		}
		m.fn41(i32(1284720), i32(43), v3+i32(15), i32(1284832), i32(1284864))
		panic("unreachable")
	l2:
		t1 := m.fn19(v1, i32(64))
		v1 = t1
	}
l1:
	m.g0 = v3 + i32(16)
	return v1
}
func (m *Module) fn900(v0 int32) {
	var v1 int32
	t0 := m.g0
	v1 = t0 - i32(16)
	m.g0 = v1
	store32(m.memory[int64(uint32(v1))+12:], uint32(i32(1277452)))
	store32(m.memory[int64(uint32(v1))+8:], uint32(v0))
	m.fn841(i32(0), v1+i32(8), i32(1279836), v1+i32(12), i32(1279836), i32(0), v1, i32(1284292))
	panic("unreachable")
}
func (m *Module) fn901(v0, v1, v2, v3, v4, v5, v6, v7 int32) {
	var v8, v9, v10, v11, v12, v13, v14, v15, v16, v17, v18, v19, v20, v21, v22, v23, v24, v25, v26, v27, v28, v29, v30, v31, v32 int32
	var v33 int64
	t0 := m.g0
	v8 = t0 - i32(80)
	m.g0 = v8
	store64(m.memory[int64(uint32(v8))+32:], uint64(i64(0)))
	store64(m.memory[int64(uint32(v8))+24:], uint64(i64(0)))
	store64(m.memory[int64(uint32(v8))+16:], uint64(i64(0)))
	store64(m.memory[int64(uint32(v8))+8:], uint64(i64(0)))
	if v3 == 0 {
		goto l0
	}
	v9 = i32(0)
	v10 = v3 << 1
	v11 = v10
	v12 = v2
	v13 = i32(15)
l3:
	{
		t1 := int32(load16(m.memory[uint32(v12):]))
		v14 = t1
		if v14 == 0 {
			goto l1
		}
		{
			if uint32(v14) > uint32(i32(15)) {
				m.fn32(v14, i32(16), i32(1290000))
				panic("unreachable")
			}
			v15 = v8 + i32(8) + v14<<1
			t2 := int32(load16(m.memory[uint32(v15):]))
			store16(m.memory[uint32(v15):], uint16(t2+i32(1)))
			p3 := v14
			if uint32(v13) < uint32(v14) {
				p3 = v13
			}
			v13 = p3
			p4 := v14
			if uint32(v9) > uint32(v14) {
				p4 = v9
			}
			v9 = p4
			goto l1
		}
	}
l1:
	v12 = v12 + i32(2)
	v11 = v11 + i32(-2)
	if v11 != 0 {
		goto l3
	}
	if v9 != 0 {
		goto l4
	}
l0:
	store32(m.memory[int64(uint32(v0))+8:], uint32(i32(2)))
	store64(m.memory[uint32(v4):], uint64(i64(0x140000001400000)))
	store64(m.memory[uint32(v0):], uint64(i64(0x100000000)))
	goto l5
l4:
	{
		if uint32(v13) > uint32(v9) {
			store32(m.memory[int64(uint32(v8))+72:], uint32(v13))
			store32(m.memory[int64(uint32(v8))+76:], uint32(v9))
			t73 := v8
			v33 = int64(uint32(i32(81))) << 32
			store64(m.memory[int64(uint32(t73))+48:], uint64(v33|int64(uint32(v8+i32(76)))))
			store64(m.memory[int64(uint32(v8))+40:], uint64(v33|int64(uint32(v8+i32(72)))))
			m.fn27(i32(1051134), v8+i32(40), i32(1290016))
			panic("unreachable")
		}
		t5 := int32(load16(m.memory[int64(uint32(v8))+10:]))
		v14 = t5
		if uint32(v14) > uint32(i32(2)) {
			goto l7
		}
		v11 = i32(4) - v14<<1
		t6 := int32(load16(m.memory[int64(uint32(v8))+12:]))
		t7 := v11 & i32(65534)
		v12 = t6
		if uint32(t7) < uint32(v12) {
			goto l7
		}
		v15 = (v11 - v12) << 1
		t8 := int32(load16(m.memory[int64(uint32(v8))+14:]))
		t9 := v15 & i32(65534)
		v11 = t8
		if uint32(t9) < uint32(v11) {
			goto l7
		}
		v16 = (v15 - v11) << 1
		t10 := int32(load16(m.memory[int64(uint32(v8))+16:]))
		t11 := v16 & i32(65534)
		v15 = t10
		if uint32(t11) < uint32(v15) {
			goto l7
		}
		v17 = (v16 - v15) << 1
		t12 := int32(load16(m.memory[int64(uint32(v8))+18:]))
		t13 := v17 & i32(65534)
		v16 = t12
		if uint32(t13) < uint32(v16) {
			goto l7
		}
		v18 = (v17 - v16) << 1
		t14 := int32(load16(m.memory[int64(uint32(v8))+20:]))
		t15 := v18 & i32(65534)
		v17 = t14
		if uint32(t15) < uint32(v17) {
			goto l7
		}
		v19 = (v18 - v17) << 1
		t16 := int32(load16(m.memory[int64(uint32(v8))+22:]))
		t17 := v19 & i32(65534)
		v18 = t16
		if uint32(t17) < uint32(v18) {
			goto l7
		}
		v20 = (v19 - v18) << 1
		t18 := int32(load16(m.memory[int64(uint32(v8))+24:]))
		t19 := v20 & i32(65534)
		v19 = t18
		if uint32(t19) < uint32(v19) {
			goto l7
		}
		v21 = (v20 - v19) & i32(0xffff) << 1
		t20 := int32(load16(m.memory[int64(uint32(v8))+26:]))
		t21 := v21
		v20 = t20
		if uint32(t21) < uint32(v20) {
			goto l7
		}
		v22 = (v21 - v20) << 1
		t22 := int32(load16(m.memory[int64(uint32(v8))+28:]))
		t23 := v22
		v21 = t22
		if uint32(t23) < uint32(v21) {
			goto l7
		}
		v23 = (v22 - v21) << 1
		t24 := int32(load16(m.memory[int64(uint32(v8))+30:]))
		t25 := v23
		v22 = t24
		if uint32(t25) < uint32(v22) {
			goto l7
		}
		v23 = (v23 - v22) << 1
		t26 := int32(load16(m.memory[int64(uint32(v8))+32:]))
		t27 := v23
		v24 = t26
		if uint32(t27) < uint32(v24) {
			goto l7
		}
		v23 = (v23 - v24) << 1
		t28 := int32(load16(m.memory[int64(uint32(v8))+34:]))
		t29 := v23
		v25 = t28
		if uint32(t29) < uint32(v25) {
			goto l7
		}
		v23 = (v23 - v25) << 1
		t30 := int32(load16(m.memory[int64(uint32(v8))+36:]))
		t31 := v23
		v26 = t30
		if uint32(t31) < uint32(v26) {
			goto l7
		}
		v23 = (v23 - v26) << 1
		t32 := int32(load16(m.memory[int64(uint32(v8))+38:]))
		t33 := v23
		v27 = t32
		if uint32(t33) < uint32(v27) {
			goto l7
		}
		{
			if v23 == v27 {
				goto l8
			}
			if v1&i32(255) == 0 {
				goto l9
			}
			if v9 != i32(1) {
				goto l9
			}
		l8:
			t35 := v13
			p34 := v9
			if uint32(v6) < uint32(v9) {
				p34 = v6
			}
			p36 := p34
			if uint32(v6) < uint32(v13) {
				p36 = t35
			}
			v23 = p36
			v6 = i32(0)
			store32(m.memory[int64(uint32(v8))+40:], uint32(i32(0)))
			store16(m.memory[int64(uint32(v8))+44:], uint16(v14))
			t37 := v8
			v14 = v14 + v12
			store16(m.memory[int64(uint32(t37))+46:], uint16(v14))
			t38 := v8
			v14 = v14 + v11
			store16(m.memory[int64(uint32(t38))+48:], uint16(v14))
			t39 := v8
			v14 = v14 + v15
			store16(m.memory[int64(uint32(t39))+50:], uint16(v14))
			t40 := v8
			v14 = v14 + v16
			store16(m.memory[int64(uint32(t40))+52:], uint16(v14))
			t41 := v8
			v14 = v14 + v17
			store16(m.memory[int64(uint32(t41))+54:], uint16(v14))
			t42 := v8
			v14 = v14 + v18
			store16(m.memory[int64(uint32(t42))+56:], uint16(v14))
			t43 := v8
			v14 = v14 + v19
			store16(m.memory[int64(uint32(t43))+58:], uint16(v14))
			t44 := v8
			v14 = v14 + v20
			store16(m.memory[int64(uint32(t44))+60:], uint16(v14))
			t45 := v8
			v14 = v14 + v21
			store16(m.memory[int64(uint32(t45))+62:], uint16(v14))
			t46 := v8
			v14 = v14 + v22
			store16(m.memory[int64(uint32(t46))+64:], uint16(v14))
			t47 := v8
			v14 = v14 + v24
			store16(m.memory[int64(uint32(t47))+66:], uint16(v14))
			t48 := v8
			v14 = v14 + v25
			store16(m.memory[int64(uint32(t48))+68:], uint16(v14))
			store16(m.memory[int64(uint32(v8))+70:], uint16(v14+v26))
			v14 = v2
		l13:
			{
				t49 := int32(load16(m.memory[uint32(v14):]))
				v12 = t49
				if v12 == 0 {
					goto l10
				}
				{
					if uint32(v12) > uint32(i32(15)) {
						m.fn32(v12, i32(16), i32(1290032))
						panic("unreachable")
					}
					v12 = v8 + i32(40) + v12<<1
					t50 := int32(load16(m.memory[uint32(v12):]))
					t51 := v12
					v12 = t50
					store16(m.memory[uint32(t51):], uint16(v12+i32(1)))
					if uint32(v12) >= uint32(i32(288)) {
						m.fn32(v12, i32(288), i32(1290048))
						panic("unreachable")
					}
					store16(m.memory[uint32(v7+v12<<1):], uint16(v6))
					goto l10
				}
			}
		l10:
			v14 = v14 + i32(2)
			v6 = v6 + i32(1)
			v10 = v10 + i32(-2)
			if v10 != 0 {
				goto l13
			}
			v22 = i32(20)
			v28 = i32(1)
			v29 = i32(2)
			v14 = v1 & i32(255)
			v27 = v14
			switch v14 {
			default:
				goto l14
			case 1:
				if uint32(v23) > uint32(i32(10)) {
					goto l17
				}
				v22 = i32(257)
				v27 = i32(31)
				v29 = i32(1290064)
				v28 = i32(1290126)
				goto l14
			case 2:
				if uint32(v23) > uint32(i32(9)) {
					goto l17
				}
				v22 = i32(0)
				v27 = i32(32)
				v29 = i32(1290158)
				v28 = i32(1290222)
			}
		l14:
			v26 = i32_shl(i32(1), v23)
			v30 = v26 + i32(-1)
			v31 = (v22 + i32(-1)) & i32(0xffff)
			v32 = v14 + i32(-1)
			v25 = i32(-1)
			v18 = i32(0)
			v21 = v23
			v24 = i32(0)
			v15 = i32(0)
			v17 = i32(0)
			v16 = i32(0)
		l41:
			{
				{
					t52 := v7
					v19 = v15
					t53 := int32(load16(m.memory[uint32(t52+v19<<1):]))
					v14 = t53
					if uint32(v14) >= uint32(v22) {
						t57 := v27
						v14 = (v14 - v22) & i32(0xffff)
						if uint32(t57) <= uint32(v14) {
							m.fn32(v14, v27, i32(1290272))
							panic("unreachable")
						}
						t58 := int32(m.memory[uint32(v28+v14)])
						v1 = t58
						t59 := int32(load16(m.memory[uint32(v29+v14<<1):]))
						v10 = t59
						goto l19
					}
					var p54 int32
					if uint32(v14) < uint32(v31) {
						p54 = 1
					}
					v12 = p54
					p55 := i32(96)
					if v12 != 0 {
						p55 = i32(0)
					}
					v1 = p55
					p56 := i32(0)
					if v12 != 0 {
						p56 = v14
					}
					v10 = p56
					goto l19
				}
			l19:
				t60 := v4
				t61 := v24 + i32_shr_u(v16, v18)
				v20 = i32_shl(i32(1), v21)
				t62 := t61 + v20
				v6 = v13 - v18
				v11 = i32_shl(i32(-1), v6)
				v12 = t62 + v11
				v14 = t60 + v12<<2
				v16 = v11 << 2
				v15 = v20
				{
				l22:
					if uint32(v12) >= uint32(v5) {
						m.fn32(v12, v5, i32(1290288))
						panic("unreachable")
					}
					store16(m.memory[uint32(v14):], uint16(v10))
					m.memory[uint32(v14+i32(3))] = byte(v6)
					m.memory[uint32(v14+i32(2))] = byte(v1)
					v14 = v14 + v16
					v12 = v12 + v11
					v15 = v15 + v11
					if v15 != 0 {
						goto l22
					}
					if uint32(v13) > uint32(i32(15)) {
						m.fn32(v13, i32(16), i32(1290304))
						panic("unreachable")
					}
					v17 = i32_shr_u(i32(-0x80000000), v13+i32(-1)) + v17
					v14 = i32_rotr(v17&i32(0xff00ff), i32(8)) | i32_rotr(v17, i32(24))&i32(0xff00ff)
					v14 = int32(uint32(v14)>>4)&i32(252645135) | v14&i32(252645135)<<4
					v14 = int32(uint32(v14)>>2)&i32(0x33333333) | v14&i32(0x33333333)<<2
					v16 = int32(uint32(v14)>>1)&i32(0x55555555) | v14&i32(0x55555555)<<1
					v15 = v19 + i32(1)
					v14 = v8 + i32(8) + v13<<1
					t63 := int32(load16(m.memory[uint32(v14):]))
					t64 := v14
					v14 = t63 + i32(-1)
					store16(m.memory[uint32(t64):], uint16(v14))
					{
						if v14&i32(0xffff) != 0 {
							goto l24
						}
						if v13 == v9 {
							if v17 == 0 {
								goto l30
							}
							if uint32(v5) < uint32(v24) {
								m.fn120(v24, v5, v5, i32(1290368))
								panic("unreachable")
							}
							{
								t68 := v16
								v14 = v5 - v24
								if uint32(t68) >= uint32(v14) {
									m.fn32(v16, v14, i32(1290352))
									panic("unreachable")
								}
								v14 = v4 + v24<<2 + v16<<2
								m.memory[int64(uint32(v14))+3] = byte(v6)
								m.memory[int64(uint32(v14))+2] = byte(i32(64))
								store16(m.memory[uint32(v14):], uint16(i32(0)))
								goto l30
							}
						}
						if v19 == i32(287) {
							m.fn32(i32(288), i32(288), i32(1290320))
							panic("unreachable")
						}
						t65 := int32(load16(m.memory[uint32(v7+v15<<1):]))
						t66 := v3
						v14 = t65
						if uint32(t66) <= uint32(v14) {
							m.fn32(v14, v3, i32(1290336))
							panic("unreachable")
						}
						t67 := int32(load16(m.memory[uint32(v2+v14<<1):]))
						v13 = t67
					}
				l24:
					if uint32(v13) <= uint32(v23) {
						goto l28
					}
					v6 = v16 & v30
					if v6 != v25 {
						t70 := v13
						p69 := v23
						if v18 != 0 {
							p69 = v18
						}
						v18 = p69
						v21 = t70 - v18
						v12 = i32_shl(i32(1), v21)
						if uint32(v13) >= uint32(v9) {
							goto l33
						}
						v21 = v9 - v18
						v14 = v8 + i32(8) + v13<<1
						v11 = v13
					l36:
						{
							t71 := int32(load16(m.memory[uint32(v14):]))
							v12 = v12 - t71
							if v12 >= i32(1) {
								v14 = v14 + i32(2)
								v12 = v12 << 1
								v11 = v11 + i32(1)
								if uint32(v11) < uint32(v9) {
									goto l36
								}
								goto l35
							}
							v21 = v11 - v18
							goto l35
						}
					}
					goto l28
				}
			l35:
				v12 = i32_shl(i32(1), v21)
			l33:
				v26 = v12 + v26
				switch v32 {
				default:
					goto l39
				case 0:
					if uint32(v26) <= uint32(i32(1332)) {
						goto l39
					}
					goto l17
				case 1:
					if uint32(v26) > uint32(i32(592)) {
						goto l17
					}
				}
			l39:
				{
					if uint32(v6) >= uint32(v5) {
						m.fn32(v6, v5, i32(1290384))
						panic("unreachable")
					}
					v14 = v4 + v6<<2
					m.memory[int64(uint32(v14))+3] = byte(v23)
					m.memory[int64(uint32(v14))+2] = byte(v21)
					t72 := v14
					v24 = v20 + v24
					store16(m.memory[uint32(t72):], uint16(v24))
					v25 = v6
					goto l28
				}
			l30:
				store32(m.memory[int64(uint32(v0))+8:], uint32(v26))
				store32(m.memory[int64(uint32(v0))+4:], uint32(v23))
				store32(m.memory[uint32(v0):], uint32(i32(0)))
				goto l5
			l28:
				if v15 != i32(288) {
					goto l41
				}
			}
			m.fn32(i32(288), i32(288), i32(1290256))
			panic("unreachable")
		}
	l9:
		store32(m.memory[uint32(v0):], uint32(i32(-1)))
		goto l5
	}
l17:
	store32(m.memory[uint32(v0):], uint32(i32(1)))
	goto l5
l7:
	store32(m.memory[uint32(v0):], uint32(i32(-1)))
l5:
	m.g0 = v8 + i32(80)
}
func (m *Module) fn902(v0, v1, v2 int32) int32 {
	var v3, v4, v5, v6, v7, v8, v9, v10 int32
	v3 = v0 & i32(0xffff)
	v4 = int32(uint32(v0) >> 16)
	switch v2 {
	default:
		if uint32(v2) < uint32(i32(16)) {
			v5 = v2 + i32(-1)
			v8 = v2 & i32(7)
			if v8 != 0 {
				goto l6
			}
			v0 = v1
			goto l7
		l6:
			v0 = v1
		l8:
			{
				v7 = v0
				v0 = v7 + i32(1)
				t2 := int32(m.memory[uint32(v7)])
				v3 = v3 + t2
				v4 = v3 + v4
				v8 = v8 + i32(-1)
				if v8 != 0 {
					goto l8
				}
			}
		l7:
			if uint32(v5) < uint32(i32(7)) {
				goto l9
			}
			v9 = v1 + v2
		l10:
			{
				t3 := int32(m.memory[uint32(v0)])
				v8 = v3 + t3
				t4 := int32(m.memory[uint32(v0+i32(1))])
				v7 = v8 + t4
				t5 := int32(m.memory[uint32(v0+i32(2))])
				v2 = v7 + t5
				t6 := int32(m.memory[uint32(v0+i32(3))])
				v1 = v2 + t6
				t7 := int32(m.memory[uint32(v0+i32(4))])
				v5 = v1 + t7
				t8 := int32(m.memory[uint32(v0+i32(5))])
				v6 = v5 + t8
				t9 := int32(m.memory[uint32(v0+i32(6))])
				v10 = v6 + t9
				t10 := int32(m.memory[uint32(v0+i32(7))])
				v3 = v10 + t10
				v4 = v3 + (v10 + (v6 + (v5 + (v1 + (v2 + (v7 + (v8 + v4)))))))
				v0 = v0 + i32(8)
				if v0 != v9 {
					goto l10
				}
			}
		l9:
			t11 := int32(uint32(v4) % uint32(i32(65521)))
			t12 := int32(uint32(v3) % uint32(i32(65521)))
			return t11<<16 | t12
		}
		t0 := int32(uint32(v2) % uint32(i32(5552)))
		t1 := v2
		v5 = t0
		v6 = t1 - v5
		if uint32(v6) < uint32(i32(5552)) {
			goto l4
		}
		v7 = v1
		v2 = v6
	l12:
		{
			v8 = i32(0)
		l11:
			{
				t13 := v3
				v0 = v7 + v8
				t14 := int32(m.memory[uint32(v0)])
				v3 = t13 + t14
				t15 := int32(m.memory[uint32(v0+i32(1))])
				t16 := v3 + v4
				v3 = v3 + t15
				t17 := int32(m.memory[uint32(v0+i32(2))])
				t18 := t16 + v3
				v3 = v3 + t17
				t19 := int32(m.memory[uint32(v0+i32(3))])
				t20 := t18 + v3
				v3 = v3 + t19
				t21 := int32(m.memory[uint32(v0+i32(4))])
				t22 := t20 + v3
				v3 = v3 + t21
				t23 := int32(m.memory[uint32(v0+i32(5))])
				t24 := t22 + v3
				v3 = v3 + t23
				t25 := int32(m.memory[uint32(v0+i32(6))])
				t26 := t24 + v3
				v3 = v3 + t25
				t27 := int32(m.memory[uint32(v0+i32(7))])
				t28 := t26 + v3
				v3 = v3 + t27
				t29 := int32(m.memory[uint32(v0+i32(8))])
				t30 := t28 + v3
				v3 = v3 + t29
				t31 := int32(m.memory[uint32(v0+i32(9))])
				t32 := t30 + v3
				v3 = v3 + t31
				t33 := int32(m.memory[uint32(v0+i32(10))])
				t34 := t32 + v3
				v3 = v3 + t33
				t35 := int32(m.memory[uint32(v0+i32(11))])
				t36 := t34 + v3
				v3 = v3 + t35
				t37 := int32(m.memory[uint32(v0+i32(12))])
				t38 := t36 + v3
				v3 = v3 + t37
				t39 := int32(m.memory[uint32(v0+i32(13))])
				t40 := t38 + v3
				v3 = v3 + t39
				t41 := int32(m.memory[uint32(v0+i32(14))])
				t42 := t40 + v3
				v3 = v3 + t41
				t43 := int32(m.memory[uint32(v0+i32(15))])
				t44 := t42 + v3
				v3 = v3 + t43
				v4 = t44 + v3
				v8 = v8 + i32(16)
				if v8 != i32(5552) {
					goto l11
				}
			}
			t45 := int32(uint32(v4) % uint32(i32(65521)))
			v4 = t45
			t46 := int32(uint32(v3) % uint32(i32(65521)))
			v3 = t46
			v7 = v7 + i32(5552)
			v2 = v2 + i32(-5552)
			if uint32(v2) < uint32(i32(5552)) {
				goto l4
			}
			goto l12
		}
	case 1:
		t47 := int32(m.memory[uint32(v1)])
		v0 = v3 + t47
		p48 := v0 + i32(-65521)
		if uint32(v0) < uint32(i32(65521)) {
			p48 = v0
		}
		v0 = p48
		t49 := int32(uint32(v0+v4) % uint32(i32(65521)))
		v0 = t49<<16 + v0
		fallthrough
	case 0:
		return v0
	}
l4:
	v7 = v1 + v6
	v2 = v5 & i32(15)
	v1 = v5 & i32(8176)
	if v1 == 0 {
		goto l13
	}
	v8 = i32(0) - v1
	v0 = v7
l14:
	{
		t50 := int32(m.memory[uint32(v0)])
		v3 = v3 + t50
		t51 := int32(m.memory[uint32(v0+i32(1))])
		t52 := v3 + v4
		v3 = v3 + t51
		t53 := int32(m.memory[uint32(v0+i32(2))])
		t54 := t52 + v3
		v3 = v3 + t53
		t55 := int32(m.memory[uint32(v0+i32(3))])
		t56 := t54 + v3
		v3 = v3 + t55
		t57 := int32(m.memory[uint32(v0+i32(4))])
		t58 := t56 + v3
		v3 = v3 + t57
		t59 := int32(m.memory[uint32(v0+i32(5))])
		t60 := t58 + v3
		v3 = v3 + t59
		t61 := int32(m.memory[uint32(v0+i32(6))])
		t62 := t60 + v3
		v3 = v3 + t61
		t63 := int32(m.memory[uint32(v0+i32(7))])
		t64 := t62 + v3
		v3 = v3 + t63
		t65 := int32(m.memory[uint32(v0+i32(8))])
		t66 := t64 + v3
		v3 = v3 + t65
		t67 := int32(m.memory[uint32(v0+i32(9))])
		t68 := t66 + v3
		v3 = v3 + t67
		t69 := int32(m.memory[uint32(v0+i32(10))])
		t70 := t68 + v3
		v3 = v3 + t69
		t71 := int32(m.memory[uint32(v0+i32(11))])
		t72 := t70 + v3
		v3 = v3 + t71
		t73 := int32(m.memory[uint32(v0+i32(12))])
		t74 := t72 + v3
		v3 = v3 + t73
		t75 := int32(m.memory[uint32(v0+i32(13))])
		t76 := t74 + v3
		v3 = v3 + t75
		t77 := int32(m.memory[uint32(v0+i32(14))])
		t78 := t76 + v3
		v3 = v3 + t77
		t79 := int32(m.memory[uint32(v0+i32(15))])
		t80 := t78 + v3
		v3 = v3 + t79
		v4 = t80 + v3
		v0 = v0 + i32(16)
		v8 = v8 + i32(16)
		if v8 != 0 {
			goto l14
		}
	}
l13:
	if v2 == 0 {
		goto l15
	}
	v1 = v7 + v1
	v8 = v5 & i32(7)
	if v8 != 0 {
		goto l16
	}
	v0 = v1
	goto l17
l16:
	v0 = v1
l18:
	{
		v7 = v0
		v0 = v7 + i32(1)
		t81 := int32(m.memory[uint32(v7)])
		v3 = v3 + t81
		v4 = v3 + v4
		v8 = v8 + i32(-1)
		if v8 != 0 {
			goto l18
		}
	}
l17:
	if uint32(v2) < uint32(i32(8)) {
		goto l15
	}
	v9 = v1 + v2
l19:
	{
		t82 := int32(m.memory[uint32(v0)])
		v8 = v3 + t82
		t83 := int32(m.memory[uint32(v0+i32(1))])
		v7 = v8 + t83
		t84 := int32(m.memory[uint32(v0+i32(2))])
		v2 = v7 + t84
		t85 := int32(m.memory[uint32(v0+i32(3))])
		v1 = v2 + t85
		t86 := int32(m.memory[uint32(v0+i32(4))])
		v5 = v1 + t86
		t87 := int32(m.memory[uint32(v0+i32(5))])
		v6 = v5 + t87
		t88 := int32(m.memory[uint32(v0+i32(6))])
		v10 = v6 + t88
		t89 := int32(m.memory[uint32(v0+i32(7))])
		v3 = v10 + t89
		v4 = v3 + (v10 + (v6 + (v5 + (v1 + (v2 + (v7 + (v8 + v4)))))))
		v0 = v0 + i32(8)
		if v0 != v9 {
			goto l19
		}
	}
l15:
	t90 := int32(uint32(v3) % uint32(i32(65521)))
	t91 := int32(uint32(v4) % uint32(i32(65521)))
	return t90 | t91<<16
}
func (m *Module) fn903(v0, v1, v2 int32) int32 {
	var v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14, v15 int32
	t0 := m.g0
	v3 = t0 - i32(1024)
	m.g0 = v3
	{
		{
			t1 := v2
			v4 = (v1+i32(3))&i32(-4) - v1
			if uint32(t1) >= uint32(v4) {
				goto l0
			}
			v5 = i32(1)
			v6 = i32(0)
			v7 = i32(4)
			v8 = i32(0)
			goto l1
		}
	l0:
		v7 = v1 + v4
		t2 := v7
		v2 = v2 - v4
		v5 = t2 + v2&i32(0x7ffffffc)
		v8 = v2 & i32(3)
		v6 = int32(uint32(v2) >> 2)
		v2 = v4
	}
l1:
	v0 = v0 ^ i32(-1)
	{
		if v2 == 0 {
			goto l2
		}
		v4 = i32(0)
		{
			if v2 == i32(1) {
				goto l3
			}
			v4 = i32(2)
			t3 := int32(m.memory[int64(uint32(v1))+1])
			t4 := int32(m.memory[uint32(v1)])
			t5 := int32(load32(m.memory[int64(uint32((t4^v0)&i32(255)<<2))+1284880:]))
			v0 = t5 ^ int32(uint32(v0)>>8)
			t6 := int32(load32(m.memory[int64(uint32((t3^v0)&i32(255)<<2))+1284880:]))
			v0 = t6 ^ int32(uint32(v0)>>8)
			if v2&i32(1) == 0 {
				goto l2
			}
		}
	l3:
		t7 := int32(m.memory[uint32(v1+v4)])
		t8 := int32(load32(m.memory[int64(uint32((t7^v0)&i32(255)<<2))+1284880:]))
		v0 = t8 ^ int32(uint32(v0)>>8)
	}
l2:
	v9 = i32(0)
	t9 := int32(uint32(v6) / uint32(i32(5)))
	v10 = t9
	t10 := v10
	var p11 int32
	if v10 != i32(0) {
		p11 = 1
	}
	v11 = t10 - p11
	if uint32(v6) >= uint32(i32(10)) {
		goto l4
	}
	v12 = i32(0)
	v13 = i32(0)
	v14 = i32(0)
	goto l5
l4:
	v1 = i32(0)
	v2 = v7
	v9 = i32(0)
	v12 = i32(0)
	v13 = i32(0)
	v14 = i32(0)
	v15 = i32(0)
l11:
	{
		if uint32(v6) <= uint32(v1) {
			goto l6
		}
		{
			v4 = v6 - v1
			p12 := v4
			if uint32(v4) > uint32(v6) {
				p12 = i32(0)
			}
			v4 = p12
			if v4 == i32(1) {
				goto l7
			}
			if v4 == i32(2) {
				v1 = v1 + i32(2)
				goto l6
			}
			if v4 == i32(3) {
				v1 = v1 + i32(3)
				goto l6
			}
			if v4 != i32(4) {
				goto l10
			}
			v1 = v1 + i32(4)
			goto l6
		}
	l7:
		v1 = v1 + i32(1)
	l6:
		m.fn32(v1, v6, i32(1283980))
		panic("unreachable")
	l10:
		t13 := int32(load32(m.memory[uint32(v2+i32(8)):]))
		v4 = v12 ^ t13
		t14 := int32(load32(m.memory[int64(uint32(v4&i32(255)<<2))+1279868:]))
		t15 := int32(load32(m.memory[int64(uint32(int32(uint32(v4)>>6)&i32(1020)))+1280892:]))
		t16 := int32(load32(m.memory[int64(uint32(int32(uint32(v4)>>14)&i32(1020)))+1281916:]))
		t17 := int32(load32(m.memory[int64(uint32(int32(uint32(v4)>>22)&i32(1020)))+1282940:]))
		v12 = t14 ^ t15 ^ t16 ^ t17
		t18 := int32(load32(m.memory[uint32(v2+i32(4)):]))
		v4 = v9 ^ t18
		t19 := int32(load32(m.memory[int64(uint32(v4&i32(255)<<2))+1279868:]))
		t20 := int32(load32(m.memory[int64(uint32(int32(uint32(v4)>>6)&i32(1020)))+1280892:]))
		t21 := int32(load32(m.memory[int64(uint32(int32(uint32(v4)>>14)&i32(1020)))+1281916:]))
		t22 := int32(load32(m.memory[int64(uint32(int32(uint32(v4)>>22)&i32(1020)))+1282940:]))
		v9 = t19 ^ t20 ^ t21 ^ t22
		t23 := int32(load32(m.memory[uint32(v2):]))
		v4 = v0 ^ t23
		t24 := int32(load32(m.memory[int64(uint32(v4&i32(255)<<2))+1279868:]))
		t25 := int32(load32(m.memory[int64(uint32(int32(uint32(v4)>>6)&i32(1020)))+1280892:]))
		t26 := int32(load32(m.memory[int64(uint32(int32(uint32(v4)>>14)&i32(1020)))+1281916:]))
		t27 := int32(load32(m.memory[int64(uint32(int32(uint32(v4)>>22)&i32(1020)))+1282940:]))
		v0 = t24 ^ t25 ^ t26 ^ t27
		t28 := int32(load32(m.memory[uint32(v2+i32(16)):]))
		v4 = v14 ^ t28
		t29 := int32(load32(m.memory[int64(uint32(v4&i32(255)<<2))+1279868:]))
		t30 := int32(load32(m.memory[int64(uint32(int32(uint32(v4)>>6)&i32(1020)))+1280892:]))
		t31 := int32(load32(m.memory[int64(uint32(int32(uint32(v4)>>14)&i32(1020)))+1281916:]))
		t32 := int32(load32(m.memory[int64(uint32(int32(uint32(v4)>>22)&i32(1020)))+1282940:]))
		v14 = t29 ^ t30 ^ t31 ^ t32
		t33 := int32(load32(m.memory[uint32(v2+i32(12)):]))
		v4 = t33 ^ v13
		t34 := int32(load32(m.memory[int64(uint32(v4&i32(255)<<2))+1279868:]))
		t35 := int32(load32(m.memory[int64(uint32(int32(uint32(v4)>>6)&i32(1020)))+1280892:]))
		t36 := int32(load32(m.memory[int64(uint32(int32(uint32(v4)>>14)&i32(1020)))+1281916:]))
		t37 := int32(load32(m.memory[int64(uint32(int32(uint32(v4)>>22)&i32(1020)))+1282940:]))
		v13 = t34 ^ t35 ^ t36 ^ t37
		v1 = v1 + i32(5)
		v2 = v2 + i32(20)
		v15 = v15 + i32(1)
		if uint32(v15) < uint32(v11) {
			goto l11
		}
	}
l5:
	{
		t38 := v6
		v2 = v11 * i32(5)
		if uint32(t38) < uint32(v2) {
			m.fn120(v2, v6, v6, i32(1283964))
			panic("unreachable")
		}
		{
			if v6 == v2 {
				goto l13
			}
			v11 = v7 + v2<<2
			t39 := int32(load32(m.memory[uint32(v11):]))
			v1 = t39
			memory_copy(m.memory, uint32(v3), uint32(i32(1285904)), uint32(i32(1024)))
			t40 := v3
			v1 = v1 ^ v0
			t41 := int32(load32(m.memory[uint32(t40+v1&i32(255)<<2):]))
			v4 = t41
			memory_copy(m.memory, uint32(v3), uint32(i32(1286928)), uint32(i32(1024)))
			t42 := int32(load32(m.memory[uint32(v3+int32(uint32(v1)>>6)&i32(1020)):]))
			v0 = t42
			memory_copy(m.memory, uint32(v3), uint32(i32(1287952)), uint32(i32(1024)))
			t43 := int32(load32(m.memory[uint32(v3+int32(uint32(v1)>>14)&i32(1020)):]))
			v15 = t43
			memory_copy(m.memory, uint32(v3), uint32(i32(1288976)), uint32(i32(1024)))
			t44 := int32(load32(m.memory[uint32(v3+int32(uint32(v1)>>22)&i32(1020)):]))
			v0 = v15 ^ (v0 ^ v4) ^ t44
			v1 = v6 - v2
			if v1 == i32(1) {
				goto l13
			}
			t45 := int32(load32(m.memory[int64(uint32(v11))+4:]))
			v2 = t45
			memory_copy(m.memory, uint32(v3), uint32(i32(1285904)), uint32(i32(1024)))
			t46 := v3
			v2 = v2 ^ v9 ^ v0
			t47 := int32(load32(m.memory[uint32(t46+v2&i32(255)<<2):]))
			v4 = t47
			memory_copy(m.memory, uint32(v3), uint32(i32(1286928)), uint32(i32(1024)))
			t48 := int32(load32(m.memory[uint32(v3+int32(uint32(v2)>>6)&i32(1020)):]))
			v0 = t48
			memory_copy(m.memory, uint32(v3), uint32(i32(1287952)), uint32(i32(1024)))
			t49 := int32(load32(m.memory[uint32(v3+int32(uint32(v2)>>14)&i32(1020)):]))
			v9 = t49
			memory_copy(m.memory, uint32(v3), uint32(i32(1288976)), uint32(i32(1024)))
			t50 := int32(load32(m.memory[uint32(v3+int32(uint32(v2)>>22)&i32(1020)):]))
			v0 = v9 ^ (v0 ^ v4) ^ t50
			if v1 == i32(2) {
				goto l13
			}
			t51 := int32(load32(m.memory[int64(uint32(v11))+8:]))
			v2 = t51
			memory_copy(m.memory, uint32(v3), uint32(i32(1285904)), uint32(i32(1024)))
			t52 := v3
			v2 = v2 ^ v12 ^ v0
			t53 := int32(load32(m.memory[uint32(t52+v2&i32(255)<<2):]))
			v4 = t53
			memory_copy(m.memory, uint32(v3), uint32(i32(1286928)), uint32(i32(1024)))
			t54 := int32(load32(m.memory[uint32(v3+int32(uint32(v2)>>6)&i32(1020)):]))
			v0 = t54
			memory_copy(m.memory, uint32(v3), uint32(i32(1287952)), uint32(i32(1024)))
			t55 := int32(load32(m.memory[uint32(v3+int32(uint32(v2)>>14)&i32(1020)):]))
			v9 = t55
			memory_copy(m.memory, uint32(v3), uint32(i32(1288976)), uint32(i32(1024)))
			t56 := int32(load32(m.memory[uint32(v3+int32(uint32(v2)>>22)&i32(1020)):]))
			v0 = v9 ^ (v0 ^ v4) ^ t56
			if v1 == i32(3) {
				goto l13
			}
			t57 := int32(load32(m.memory[int64(uint32(v11))+12:]))
			v2 = t57
			memory_copy(m.memory, uint32(v3), uint32(i32(1285904)), uint32(i32(1024)))
			t58 := v3
			v2 = v2 ^ v13 ^ v0
			t59 := int32(load32(m.memory[uint32(t58+v2&i32(255)<<2):]))
			v4 = t59
			memory_copy(m.memory, uint32(v3), uint32(i32(1286928)), uint32(i32(1024)))
			t60 := int32(load32(m.memory[uint32(v3+int32(uint32(v2)>>6)&i32(1020)):]))
			v0 = t60
			memory_copy(m.memory, uint32(v3), uint32(i32(1287952)), uint32(i32(1024)))
			t61 := int32(load32(m.memory[uint32(v3+int32(uint32(v2)>>14)&i32(1020)):]))
			v9 = t61
			memory_copy(m.memory, uint32(v3), uint32(i32(1288976)), uint32(i32(1024)))
			t62 := int32(load32(m.memory[uint32(v3+int32(uint32(v2)>>22)&i32(1020)):]))
			v0 = v9 ^ (v0 ^ v4) ^ t62
			if v1 == i32(4) {
				goto l13
			}
			t63 := int32(load32(m.memory[int64(uint32(v11))+16:]))
			v2 = t63
			memory_copy(m.memory, uint32(v3), uint32(i32(1285904)), uint32(i32(1024)))
			t64 := v3
			v2 = v2 ^ v14 ^ v0
			t65 := int32(load32(m.memory[uint32(t64+v2&i32(255)<<2):]))
			v4 = t65
			memory_copy(m.memory, uint32(v3), uint32(i32(1286928)), uint32(i32(1024)))
			t66 := int32(load32(m.memory[uint32(v3+int32(uint32(v2)>>6)&i32(1020)):]))
			v0 = t66
			memory_copy(m.memory, uint32(v3), uint32(i32(1287952)), uint32(i32(1024)))
			t67 := int32(load32(m.memory[uint32(v3+int32(uint32(v2)>>14)&i32(1020)):]))
			v9 = t67
			memory_copy(m.memory, uint32(v3), uint32(i32(1288976)), uint32(i32(1024)))
			t68 := int32(load32(m.memory[uint32(v3+int32(uint32(v2)>>22)&i32(1020)):]))
			v0 = v9 ^ (v0 ^ v4) ^ t68
			if v1 == i32(5) {
				goto l13
			}
			t70 := v10 * i32(20)
			p69 := v10
			if v10 != 0 {
				p69 = i32(1)
			}
			v2 = p69
			v1 = t70 - v2*i32(20) + v7 + i32(20)
			v4 = v6 + v2*i32(5) - v10*i32(5) + i32(-5)
		l14:
			{
				t71 := int32(load32(m.memory[uint32(v1):]))
				v2 = t71
				memory_copy(m.memory, uint32(v3), uint32(i32(1285904)), uint32(i32(1024)))
				t72 := v3
				v2 = v2 ^ v0
				t73 := int32(load32(m.memory[uint32(t72+v2&i32(255)<<2):]))
				v6 = t73
				memory_copy(m.memory, uint32(v3), uint32(i32(1286928)), uint32(i32(1024)))
				t74 := int32(load32(m.memory[uint32(v3+int32(uint32(v2)>>6)&i32(1020)):]))
				v0 = t74
				memory_copy(m.memory, uint32(v3), uint32(i32(1287952)), uint32(i32(1024)))
				t75 := int32(load32(m.memory[uint32(v3+int32(uint32(v2)>>14)&i32(1020)):]))
				v9 = t75
				memory_copy(m.memory, uint32(v3), uint32(i32(1288976)), uint32(i32(1024)))
				t76 := int32(load32(m.memory[uint32(v3+int32(uint32(v2)>>22)&i32(1020)):]))
				v0 = v9 ^ (v0 ^ v6) ^ t76
				v1 = v1 + i32(4)
				v4 = v4 + i32(-1)
				if v4 != 0 {
					goto l14
				}
			}
		}
	l13:
		{
			if v8 == 0 {
				goto l15
			}
			t77 := int32(m.memory[uint32(v5)])
			t78 := int32(load32(m.memory[int64(uint32((t77^v0)&i32(255)<<2))+1284880:]))
			v0 = t78 ^ int32(uint32(v0)>>8)
			if v8 == i32(1) {
				goto l15
			}
			t79 := int32(m.memory[int64(uint32(v5))+1])
			t80 := int32(load32(m.memory[int64(uint32((t79^v0)&i32(255)<<2))+1284880:]))
			v0 = t80 ^ int32(uint32(v0)>>8)
			if v8 == i32(2) {
				goto l15
			}
			t81 := int32(m.memory[int64(uint32(v5))+2])
			t82 := int32(load32(m.memory[int64(uint32((t81^v0)&i32(255)<<2))+1284880:]))
			v0 = t82 ^ int32(uint32(v0)>>8)
		}
	l15:
		m.g0 = v3 + i32(1024)
		return v0 ^ i32(-1)
	}
}
func (m *Module) fn904(v0, v1, v2 int32) {
	var v3, v4, v5, v6, v7 int32
	t0 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	t1 := v0
	v3 = t0
	v4 = v3 + v2
	store32(m.memory[int64(uint32(t1))+8:], uint32(v4))
	t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	v5 = t2
	t3 := v5
	v6 = v4 + i32(8)
	p4 := v6
	if uint32(v5) < uint32(v6) {
		p4 = t3
	}
	v4 = p4
	t5 := int32(load32(m.memory[uint32(v0):]))
	v7 = t5
	{
		if uint32(v2) > uint32(v1) {
			goto l0
		}
		if uint32(v3) < uint32(v1) {
			m.fn139(i32(1277376), i32(9), i32(1277388))
			panic("unreachable")
		}
		v0 = v3 - v1
		if uint32(v6) < uint32(v5) {
			if v2 == 0 {
				return
			}
			v3 = v7 + v3
			t6 := v3
			v0 = v7 + v0
			t7 := int64(load64(m.memory[uint32(v0):]))
			store64(m.memory[uint32(t6):], uint64(t7))
			if uint32(v2) < uint32(i32(9)) {
				return
			}
			v2 = v0 + v2
			v0 = i32(0) - v1
			v3 = v3 + i32(8)
		l7:
			{
				t8 := int64(load64(m.memory[uint32(v3+v0):]))
				store64(m.memory[uint32(v3):], uint64(t8))
				v3 = v3 + i32(8)
				if uint32(v3+v0) < uint32(v2) {
					goto l7
				}
				return
			}
		}
		v1 = v0 + v2
		if uint32(v1) > uint32(v4) {
			m.fn120(i32(0), v1, v4, i32(1290400))
			panic("unreachable")
		}
		if uint32(v0) > uint32(v1) {
			m.fn120(v0, v1, v4, i32(1290416))
			panic("unreachable")
		}
		if uint32(v3) > uint32(v4-v2) {
			m.fn27(i32(1277256), i32(43), i32(1277360))
			panic("unreachable")
		}
		if v2 == 0 {
			return
		}
		memory_copy(m.memory, uint32(v7+v3), uint32(v7+v0), uint32(v2))
		return
	l0:
		{
			if v1 == i32(1) {
				goto l8
			}
			v0 = v4 - v3
			p9 := v0
			if uint32(v0) > uint32(v4) {
				p9 = i32(0)
			}
			v0 = p9
			v5 = v7 - v1
			v1 = i32(0) - v1
		l11:
			{
				v6 = v1 + v3
				if uint32(v6) >= uint32(v4) {
					m.fn32(v6, v4, i32(1277328))
					panic("unreachable")
				}
				if v0 == 0 {
					m.fn32(v3, v4, i32(1277344))
					panic("unreachable")
				}
				t10 := int32(m.memory[uint32(v5+v3)])
				m.memory[uint32(v7+v3)] = byte(t10)
				v3 = v3 + i32(1)
				v0 = v0 + i32(-1)
				v2 = v2 + i32(-1)
				if v2 != 0 {
					goto l11
				}
				return
			}
		}
	l8:
		v0 = v3 + i32(-1)
		if uint32(v0) >= uint32(v4) {
			m.fn32(v0, v4, i32(1277280))
			panic("unreachable")
		}
		if uint32(v4) < uint32(v3) {
			m.fn120(v3, v4, v4, i32(1277312))
			panic("unreachable")
		}
		t11 := v2
		v4 = v4 - v3
		if uint32(t11) > uint32(v4) {
			m.fn120(i32(0), v2, v4, i32(1277296))
			panic("unreachable")
		}
		if v2 == 0 {
			return
		}
		t12 := int32(m.memory[uint32(v7+v0)])
		memory_fill(m.memory, uint32(v7+v3), t12, uint32(v2))
	}
}
func (m *Module) fn905(v0, v1, v2, v3 int32) {
	var v4, v5, v6, v7, v8, v9 int32
	{
		{
			t0 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			v4 = t0
			t1 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			t2 := v4
			v5 = t1
			v6 = t2 - v5
			t3 := v6
			v7 = v3 - v2
			if uint32(t3) >= uint32(v7+i32(8)) {
				if v3 == v2 {
					goto l5
				}
				t7 := int32(load32(m.memory[uint32(v0):]))
				v4 = t7 + v5
				t8 := int32(load32(m.memory[uint32(v1):]))
				t9 := v4
				v1 = t8
				v2 = v1 + v2
				t10 := int64(load64(m.memory[uint32(v2):]))
				store64(m.memory[uint32(t9):], uint64(t10))
				v2 = v2 + i32(8)
				t11 := v2
				v1 = v1 + v3
				if uint32(t11) >= uint32(v1) {
					goto l5
				}
				v3 = v4 + i32(8)
			l6:
				{
					t12 := int64(load64(m.memory[uint32(v2):]))
					store64(m.memory[uint32(v3):], uint64(t12))
					v3 = v3 + i32(8)
					v2 = v2 + i32(8)
					if uint32(v2) < uint32(v1) {
						goto l6
					}
					goto l5
				}
			}
			t4 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			v8 = t4
			t5 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t6 := v8
			v9 = t5
			if uint32(t6) > uint32(v9) {
				m.fn120(i32(0), v8, v9, i32(1284152))
				panic("unreachable")
			}
			if uint32(v3) < uint32(v2) {
				goto l2
			}
			if uint32(v3) > uint32(v8) {
				goto l2
			}
			if uint32(v4) < uint32(v5) {
				m.fn120(v5, v4, v4, i32(1277420))
				panic("unreachable")
			}
			if uint32(v7) <= uint32(v6) {
				goto l4
			}
			m.fn120(i32(0), v7, v6, i32(1277404))
			panic("unreachable")
		}
	l2:
		m.fn120(v2, v3, v8, i32(1277436))
		panic("unreachable")
	l4:
		if v7 == 0 {
			goto l5
		}
		t13 := int32(load32(m.memory[uint32(v0):]))
		t14 := int32(load32(m.memory[uint32(v1):]))
		memory_copy(m.memory, uint32(t13+v5), uint32(t14+v2), uint32(v7))
	}
l5:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v5+v7))
}
func (m *Module) fn906(v0 int32) {
	var v1 int32
	var v2 int64
	var v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13 int32
	var v14, v15 int64
	var v16, v17, v18, v19, v20, v21 int32
	var v22 int64
	var v23, v24, v25, v26, v27 int32
	var v28, v29 int64
	t0 := m.g0
	v1 = t0 - i32(16)
	m.g0 = v1
	t1 := int64(load64(m.memory[int64(uint32(v0))+48:]))
	v2 = t1
	store64(m.memory[int64(uint32(v0))+48:], uint64(i64(0)))
	t2 := int32(load32(m.memory[int64(uint32(v0))+56:]))
	v3 = t2
	store32(m.memory[int64(uint32(v0))+56:], uint32(i32(1)))
	t3 := int32(load32(m.memory[int64(uint32(v0))+60:]))
	v4 = t3
	t4 := int32(load32(m.memory[int64(uint32(v0))+64:]))
	v5 = t4
	store64(m.memory[int64(uint32(v0))+60:], uint64(i64(1)))
	t5 := int32(load32(m.memory[int64(uint32(v0))+72:]))
	v6 = t5
	store32(m.memory[int64(uint32(v0))+72:], uint32(i32(1)))
	t6 := int32(load32(m.memory[int64(uint32(v0))+80:]))
	v7 = t6
	t7 := int32(load32(m.memory[int64(uint32(v0))+76:]))
	v8 = t7
	store64(m.memory[int64(uint32(v0))+76:], uint64(i64(0)))
	store32(m.memory[int64(uint32(v1))+4:], uint32(v6))
	store32(m.memory[int64(uint32(v1))+8:], uint32(v8))
	store32(m.memory[int64(uint32(v1))+12:], uint32(v7))
	v9 = i32(1277472)
	v10 = i32(512)
	{
		t8 := int32(m.memory[int64(uint32(v0))+152])
		switch t8 {
		default:
			goto l0
		case 1:
			v9 = v0 + i32(164)
			v10 = i32(1332)
			goto l0
		case 2:
			v9 = v0 + i32(5492)
			v10 = i32(1332)
			goto l0
		case 3:
			v9 = v0 + i32(10820)
			v10 = i32(592)
		}
	}
l0:
	v11 = i32(1279520)
	v12 = i32(32)
	{
		t9 := int32(m.memory[int64(uint32(v0))+160])
		switch t9 {
		default:
			goto l4
		case 1:
			v11 = v0 + i32(164)
			v12 = i32(1332)
			goto l4
		case 2:
			v11 = v0 + i32(5492)
			v12 = i32(1332)
			goto l4
		case 3:
			v11 = v0 + i32(10820)
			v12 = i32(592)
		}
	}
l4:
	{
		t10 := int32(load32(m.memory[int64(uint32(v0))+12:]))
		v6 = t10
		v13 = v6 + i32(-64)
		if uint32(v13) >= uint32(i32(-63)) {
			m.fn3(i32(1283996), i32(74), i32(1284072))
			panic("unreachable")
		}
		t11 := int64(load32(m.memory[int64(uint32(v0))+156:]))
		v14 = i64_shl(i64(-1), t11)
		t12 := int64(load32(m.memory[int64(uint32(v0))+148:]))
		v15 = i64_shl(i64(-1), t12) ^ i64(-1)
		var p13 int32
		if uint32(v13) > uint32(v6) {
			p13 = 1
		}
		v16 = p13
		{
			if uint32(v5&i32(255)) <= uint32(i32(9)) {
				goto l9
			}
			v6 = v5
			v17 = v3
			goto l10
		l9:
			v6 = v5 | i32(56)
			v17 = v3 + (int32(uint32(v5)>>3)&i32(1) ^ i32(7))
			t14 := int64(load64(m.memory[uint32(v3):]))
			v2 = i64_shl(t14, int64(uint32(v5&i32(15)))) | v2
		}
	l10:
		v18 = v0 + i32(72)
		v19 = v5 & i32(-256)
		t15 := int32(load32(m.memory[int64(uint32(v0))+68:]))
		v20 = t15
		v21 = v0 + i32(8)
		v22 = v14 ^ i64(-1)
		p16 := v13
		if v16 != 0 {
			p16 = i32(0)
		}
		v23 = p16
		v14 = int64(uint64(v2) >> 32)
		v24 = int32(v15)
		v13 = int32(v2)
		{
		l46:
			{
				t17 := int64(load64(m.memory[uint32(v17):]))
				t18 := v13
				v2 = i64_shl(t17, int64(uint32(v6)))
				v25 = t18 | int32(v2)
				{
					{
						{
							{
								{
									t19 := int32(load32(m.memory[int64(uint32(v0))+148:]))
									v26 = v6 & i32(255)
									if uint32(t19) <= uint32(v26) {
										goto l11
									}
									t20 := v10
									v5 = v25 & v24
									if uint32(t20) > uint32(v5) {
										goto l12
									}
									m.fn32(v5, v10, i32(1279648))
									panic("unreachable")
								}
							l11:
								t21 := v10
								v5 = v13 & v24
								if uint32(t21) <= uint32(v5) {
									m.fn32(v5, v10, i32(1279664))
									panic("unreachable")
								}
							}
						l12:
							v6 = v6 | i32(56)
							v2 = int64(uint64(v2)>>32) | v14
							v5 = v9 + v5<<2
							t22 := int32(m.memory[int64(uint32(v5))+3])
							v13 = t22
							t23 := int32(load16(m.memory[uint32(v5):]))
							v3 = t23
							t24 := int32(m.memory[int64(uint32(v5))+2])
							v5 = t24
							if v5 == 0 {
								goto l14
							}
							v16 = v7
							goto l15
						}
					l14:
						if uint32(v7) >= uint32(v8) {
							m.fn32(v7, v8, i32(1279680))
							panic("unreachable")
						}
						t25 := int32(load32(m.memory[int64(uint32(v1))+4:]))
						v27 = t25
						m.memory[uint32(v27+v7)] = byte(v3)
						t26 := v1
						v16 = v7 + i32(1)
						store32(m.memory[int64(uint32(t26))+12:], uint32(v16))
						{
							t27 := v10
							v14 = i64_shr_u(v2<<32|int64(uint32(v25)), int64(uint32(v13)))
							v5 = int32(v14 & v15)
							if uint32(t27) <= uint32(v5) {
								m.fn32(v5, v10, i32(1279696))
								panic("unreachable")
							}
							v6 = v6 - v13
							v5 = v9 + v5<<2
							t28 := int32(load16(m.memory[uint32(v5):]))
							v3 = t28
							t29 := int32(m.memory[int64(uint32(v5))+3])
							v13 = t29
							t30 := int32(m.memory[int64(uint32(v5))+2])
							v5 = t30
							if v5 == 0 {
								goto l18
							}
							v2 = int64(uint64(v14) >> 32)
							v25 = int32(v14)
							goto l15
						}
					l18:
						if uint32(v16) >= uint32(v8) {
							m.fn32(v16, v8, i32(1279680))
							panic("unreachable")
						}
						m.memory[uint32(v27+v16)] = byte(v3)
						t31 := v1
						v16 = v7 + i32(2)
						store32(m.memory[int64(uint32(t31))+12:], uint32(v16))
						t32 := v10
						v14 = i64_shr_u(v14, int64(uint32(v13)))
						v5 = int32(v14 & v15)
						if uint32(t32) <= uint32(v5) {
							m.fn32(v5, v10, i32(1279712))
							panic("unreachable")
						}
						v2 = int64(uint64(v14) >> 32)
						v6 = v6 - v13
						v5 = v9 + v5<<2
						t33 := int32(load16(m.memory[uint32(v5):]))
						v3 = t33
						t34 := int32(m.memory[int64(uint32(v5))+3])
						v13 = t34
						t35 := int32(m.memory[int64(uint32(v5))+2])
						v5 = t35
						v25 = int32(v14)
					}
				l15:
					v17 = v17 + (int32(uint32(v26)>>3) ^ i32(7))
					v6 = v6 - v13
					v2 = i64_shr_u(v2<<32|int64(uint32(v25)), int64(uint32(v13)))
					v14 = int64(uint64(v2) >> 32)
					v13 = int32(v2)
					if v5&i32(255) == 0 {
						goto l21
					}
				l25:
					if v5&i32(16) != 0 {
						t41 := v12
						t42 := v2
						v28 = int64(uint32(v5)) & i64(15)
						v14 = i64_shr_u(t42, v28)
						v13 = int32(v14 & v22)
						if uint32(t41) > uint32(v13) {
							v25 = v11 + v13<<2
							t43 := int32(m.memory[int64(uint32(v25))+3])
							v7 = t43
							t44 := int32(m.memory[int64(uint32(v25))+2])
							v13 = t44
							{
								v5 = v6 - v5&i32(15)
								if uint32(v5&i32(255)) < uint32(i32(28)) {
									goto l29
								}
								v6 = v5
								v8 = v17
								goto l30
							l29:
								v6 = v5 | i32(56)
								v8 = v17 + (int32(uint32(v5&i32(248))>>3) ^ i32(7))
								t45 := int64(load64(m.memory[uint32(v17):]))
								v14 = i64_shl(t45, int64(uint32(v5))&i64(255)) | v14
							}
						l30:
							t46 := int32(load16(m.memory[uint32(v25):]))
							v17 = t46
							v6 = v6 - v7
							v29 = i64_shr_u(v14, int64(uint32(v7)))
							{
								{
									if v13&i32(16) != 0 {
										goto l31
									}
								l34:
									{
										if v13&i32(64) != 0 {
											m.memory[uint32(v0)] = byte(i32(30))
											v13 = int32(v29)
											v9 = i32(1065754)
											v3 = i32(1)
											v10 = i32(22)
											goto l42
										}
										t47 := v12
										v5 = (v17 + int32(v29&(i64_shl(i64(-1), int64(uint32(v13))&i64(47))^i64(-1)))) & i32(0xffff)
										if uint32(t47) <= uint32(v5) {
											m.fn32(v5, v12, i32(1279760))
											panic("unreachable")
										}
										t48 := v6
										v5 = v11 + v5<<2
										t49 := int32(m.memory[int64(uint32(v5))+3])
										v13 = t49
										v6 = t48 - v13
										v29 = i64_shr_u(v29, int64(uint32(v13)))
										t50 := int32(load16(m.memory[uint32(v5):]))
										v17 = t50
										t51 := int32(m.memory[int64(uint32(v5))+2])
										v13 = t51
										if v13&i32(16) == 0 {
											goto l34
										}
									}
								l31:
									v3 = v3 + int32(v2&(i64_shl(i64(-1), v28)^i64(-1)))
									v6 = v6 - v13&i32(15)
									t52 := v29
									v2 = int64(uint32(v13)) & i64(15)
									v28 = i64_shr_u(t52, v2)
									v14 = int64(uint64(v28) >> 32)
									v13 = int32(v28)
									v17 = (v17 + int32(v29&(i64_shl(i64(-1), v2)^i64(-1)))) & i32(0xffff)
									if uint32(v17) > uint32(v16) {
										v5 = v17 - v16
										t53 := int32(load32(m.memory[int64(uint32(v0))+16:]))
										if uint32(v5) > uint32(t53) {
											t55 := int32(m.memory[int64(uint32(v0))+1])
											if t55&i32(4) == 0 {
												m.fn27(i32(1279776), i32(85), i32(1279820))
												panic("unreachable")
											}
											v10 = i32(30)
											m.memory[uint32(v0)] = byte(i32(30))
											v9 = i32(1065312)
											v3 = i32(1)
											goto l42
										}
										{
											t54 := int32(load32(m.memory[int64(uint32(v0))+20:]))
											v25 = t54
											if v25 != 0 {
												if uint32(v25) < uint32(v5) {
													goto l40
												}
												v16 = v25 - v5
												goto l39
											}
											v16 = v23 - v5
											goto l39
										}
									}
									m.fn904(v1+i32(4), v17, v3&i32(0xffff))
									v17 = v8
									goto l36
								}
							l42:
								v17 = v8
								goto l43
							l40:
								t56 := v23
								v5 = v5 - v25
								v16 = t56 - v5
								if uint32(v5) >= uint32(v3&i32(0xffff)) {
									goto l39
								}
								m.fn905(v1+i32(4), v21, v16, v23)
								v3 = v3 - v5
								v16 = i32(0)
								v5 = v25
							}
						l39:
							t57 := v1 + i32(4)
							t58 := v21
							t59 := v16
							t60 := v5
							v3 = v3 & i32(0xffff)
							p61 := v3
							if uint32(v5) < uint32(v3) {
								p61 = t60
							}
							m.fn905(t57, t58, t59, p61+v16)
							if uint32(v3) > uint32(v5) {
								m.fn904(v1+i32(4), v17, v3-v5)
								v17 = v8
								goto l36
							}
							v17 = v8
							goto l36
						}
						m.fn32(v13, v12, i32(1279744))
						panic("unreachable")
					}
					if v5&i32(64) != 0 {
						if v5&i32(32) == 0 {
							m.memory[uint32(v0)] = byte(i32(30))
							v9 = i32(1065726)
							v3 = i32(1)
							v10 = i32(28)
							goto l43
						}
						m.memory[uint32(v0)] = byte(i32(12))
						goto l27
					}
					{
						t36 := v10
						v5 = (v3 + int32(v2&(i64_shl(i64(-1), int64(uint32(v5))&i64(47))^i64(-1)))) & i32(0xffff)
						if uint32(t36) <= uint32(v5) {
							goto l24
						}
						t37 := v14<<32 | int64(uint32(v13))
						v5 = v9 + v5<<2
						t38 := int32(m.memory[int64(uint32(v5))+3])
						v3 = t38
						v2 = i64_shr_u(t37, int64(uint32(v3)))
						v14 = int64(uint64(v2) >> 32)
						v13 = int32(v2)
						v6 = v6 - v3
						t39 := int32(load16(m.memory[uint32(v5):]))
						v3 = t39
						t40 := int32(m.memory[int64(uint32(v5))+2])
						v5 = t40
						if v5 == 0 {
							goto l21
						}
						goto l25
					}
				l24:
					m.fn32(v5, v10, i32(1279728))
					panic("unreachable")
				l21:
					if uint32(v16) >= uint32(v8) {
						m.fn32(v16, v8, i32(1279680))
						panic("unreachable")
					}
					t62 := int32(load32(m.memory[int64(uint32(v1))+4:]))
					m.memory[uint32(t62+v16)] = byte(v3)
					store32(m.memory[int64(uint32(v1))+12:], uint32(v16+i32(1)))
				}
			l36:
				if uint32(v4-v17+int32(uint32(v6&i32(248))>>3)) <= uint32(i32(14)) {
					goto l27
				}
				t63 := int32(load32(m.memory[int64(uint32(v1))+8:]))
				v8 = t63
				t64 := int32(load32(m.memory[int64(uint32(v1))+12:]))
				t65 := v8
				v7 = t64
				if uint32(t65-v7) > uint32(i32(259)) {
					goto l46
				}
			}
		l27:
			v3 = i32(0)
			v9 = i32(0)
		l43:
			store32(m.memory[int64(uint32(v0))+68:], uint32(v20))
			store32(m.memory[int64(uint32(v0))+60:], uint32(v4))
			t66 := int64(load64(m.memory[int64(uint32(v1))+4:]))
			store64(m.memory[uint32(v18):], uint64(t66))
			t67 := int32(load32(m.memory[int64(uint32(v1))+12:]))
			store32(m.memory[int64(uint32(v18))+8:], uint32(t67))
			t68 := v0
			t69 := v19
			v5 = v6 & i32(7)
			store32(m.memory[int64(uint32(t68))+64:], uint32(t69|v5))
			store32(m.memory[int64(uint32(v0))+56:], uint32(v17-int32(uint32(v6&i32(248))>>3)))
			store64(m.memory[int64(uint32(v0))+48:], uint64(uint32(v13&(i32_shl(i32(-1), v5)^i32(-1)))))
			if v3 == 0 {
				goto l47
			}
			store32(m.memory[int64(uint32(v0))+136:], uint32(v10))
			store32(m.memory[int64(uint32(v0))+132:], uint32(v9))
		l47:
			m.g0 = v1 + i32(16)
			return
		}
	}
}
func (m *Module) fn907(v0, v1, v2, v3, v4, v5, v6 int32) {
	var v7, v8, v9, v10, v11, v12, v13, v14 int32
	var v15 int64
	var v16, v17 int32
	{
		if uint32(v4) < uint32(i32(16)) {
			if uint32(v6) > uint32(v4) {
				goto l4
			}
			v8 = v3 + v4
			if v6 != 0 {
				goto l5
			}
			v4 = i32(0)
			goto l6
		l5:
			v9 = v3 + v6
			v4 = i32(0)
			v10 = v3
		l7:
			{
				t2 := int32(m.memory[uint32(v10)])
				v4 = v4<<1 + t2
				v10 = v10 + i32(1)
				if uint32(v10) < uint32(v9) {
					goto l7
				}
			}
		l6:
			t3 := int32(load32(m.memory[int64(uint32(v1))+44:]))
			v11 = t3
			t4 := int32(load32(m.memory[int64(uint32(v1))+40:]))
			v9 = t4
			v1 = v8 - v6
			v10 = v3
		l10:
			{
				{
					if v9 != v4 {
						goto l8
					}
					t5 := m.fn908(v10, v5, v6)
					if t5 == 0 {
						goto l8
					}
					v9 = v10 - v3
					v12 = i32(1)
					goto l9
				}
			l8:
				if uint32(v10) >= uint32(v1) {
					goto l4
				}
				t6 := int32(m.memory[uint32(v10)])
				t7 := int32(m.memory[uint32(v10+v6)])
				v4 = (v4-v11*t6)<<1 + t7
				v10 = v10 + i32(1)
				goto l10
			}
		}
		t0 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		v7 = t0
		t1 := int32(load32(m.memory[uint32(v1):]))
		if t1 == 0 {
			if v6 == 0 {
				goto l3
			}
			if uint32(v6) > uint32(v4) {
				goto l4
			}
			v13 = v6 + i32(-1)
			v14 = v6 - v7
			t8 := int64(load64(m.memory[int64(uint32(v1))+8:]))
			v15 = t8
			t9 := int32(load32(m.memory[int64(uint32(v1))+16:]))
			v16 = t9
			v10 = i32(0)
			v8 = v6
			v1 = i32(0)
			{
			l25:
				{
					t10 := v13
					v9 = v10
					v17 = t10 + v9
					if uint32(v17) >= uint32(v4) {
						m.fn32(v17, v4, i32(1271756))
						panic("unreachable")
					}
					v12 = i32(0)
					v11 = i32(0)
					v10 = v8
					{
						t11 := int64(m.memory[uint32(v3+v17)])
						if i64_shr_u(v15, t11)&i64(1) == 0 {
							goto l12
						}
						{
							p12 := v16
							if uint32(v1) > uint32(v16) {
								p12 = v1
							}
							v8 = p12
							if uint32(v8) >= uint32(v6) {
								goto l13
							}
							v11 = v3 + v9
							v10 = v8
						l16:
							{
								if uint32(v9+v10) >= uint32(v4) {
									goto l14
								}
								t13 := int32(m.memory[uint32(v5+v10)])
								t14 := int32(m.memory[uint32(v11+v10)])
								if t13 != t14 {
									goto l15
								}
								t15 := v6
								v10 = v10 + i32(1)
								if t15 != v10 {
									goto l16
								}
							}
						}
					l13:
						if uint32(v16) <= uint32(v1) {
							goto l17
						}
						v11 = v3 + v9
						v10 = v16
					l22:
						if uint32(v10) >= uint32(v6) {
							m.fn32(v10, v6, i32(1271788))
							panic("unreachable")
						}
						{
							v8 = v9 + v10
							if uint32(v8) >= uint32(v4) {
								m.fn32(v8, v4, i32(1271804))
								panic("unreachable")
							}
							t16 := int32(m.memory[uint32(v5+v10)])
							t17 := int32(m.memory[uint32(v11+v10)])
							if t16 == t17 {
								goto l20
							}
							v11 = v14
							v10 = v7
							goto l21
						}
					l20:
						v10 = v10 + i32(-1)
						if uint32(v10) > uint32(v1) {
							goto l22
						}
					l17:
						if uint32(v1) >= uint32(v6) {
							m.fn32(v1, v6, i32(1271820))
							panic("unreachable")
						}
						{
							v8 = v9 + v1
							if uint32(v8) >= uint32(v4) {
								m.fn32(v8, v4, i32(1271836))
								panic("unreachable")
							}
							v11 = v14
							v10 = v7
							t18 := int32(m.memory[uint32(v5+v1)])
							t19 := int32(m.memory[uint32(v3+v8)])
							if t18 != t19 {
								goto l21
							}
							v12 = i32(1)
							goto l9
						}
					l15:
						v10 = v10 - v16 + i32(1)
						v11 = i32(0)
					l21:
						v10 = v10 + v9
					}
				l12:
					v1 = v11
					v8 = v10 + v6
					if uint32(v8) <= uint32(v4) {
						goto l25
					}
					goto l9
				l14:
				}
				t20 := v4
				v10 = v9 + v8
				p21 := v10
				if uint32(v4) > uint32(v10) {
					p21 = t20
				}
				m.fn32(p21, v4, i32(1271772))
				panic("unreachable")
			}
		}
		if v6 != 0 {
			if uint32(v6) > uint32(v4) {
				goto l4
			}
			v16 = v6 + i32(-1)
			t22 := int64(load64(m.memory[int64(uint32(v1))+8:]))
			v15 = t22
			v9 = i32(0)
			t23 := int32(load32(m.memory[int64(uint32(v1))+16:]))
			v8 = t23
			var p24 int32
			if uint32(v8+i32(-1)) >= uint32(v6) {
				p24 = 1
			}
			v11 = p24
		l37:
			{
				v10 = v16 + v9
				if uint32(v10) >= uint32(v4) {
					m.fn32(v10, v4, i32(1271676))
					panic("unreachable")
				}
				v1 = v6
				t25 := int64(m.memory[uint32(v3+v10)])
				if i64_shr_u(v15, t25)&i64(1) == 0 {
					goto l27
				}
				v10 = v8
				if uint32(v8) >= uint32(v6) {
					goto l35
				}
				v1 = v3 + v9
				v10 = v8
			l31:
				{
					if uint32(v9+v10) >= uint32(v4) {
						t31 := v4
						v10 = v8 + v9
						p32 := v10
						if uint32(v4) > uint32(v10) {
							p32 = t31
						}
						m.fn32(p32, v4, i32(1271692))
						panic("unreachable")
					}
					t26 := int32(m.memory[uint32(v5+v10)])
					t27 := int32(m.memory[uint32(v1+v10)])
					if t26 != t27 {
						v9 = v9 - v8 + v10 + i32(1)
						goto l36
					}
					t28 := v6
					v10 = v10 + i32(1)
					if t28 != v10 {
						goto l31
					}
				}
				v10 = v8
			l35:
				{
					if v10 != 0 {
						goto l32
					}
					v12 = i32(1)
					goto l9
				l32:
					v10 = v10 + i32(-1)
					if v11 != 0 {
						m.fn32(v10, v6, i32(1271708))
						panic("unreachable")
					}
					v1 = v10 + v9
					if uint32(v1) >= uint32(v4) {
						m.fn32(v1, v4, i32(1271724))
						panic("unreachable")
					}
					t29 := int32(m.memory[uint32(v5+v10)])
					t30 := int32(m.memory[uint32(v3+v1)])
					if t29 == t30 {
						goto l35
					}
				}
				v1 = v7
				goto l27
			}
		l27:
			v9 = v1 + v9
		l36:
			v12 = i32(0)
			if uint32(v9+v6) <= uint32(v4) {
				goto l37
			}
			goto l9
		}
		goto l3
	}
l4:
	v12 = i32(0)
	goto l9
l3:
	v9 = i32(0)
	v12 = i32(1)
l9:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v9))
	store32(m.memory[uint32(v0):], uint32(v12))
}
func (m *Module) fn908(v0, v1, v2 int32) int32 {
	var v3 int32
	if uint32(v2) < uint32(i32(4)) {
		goto l0
	}
l2:
	{
		t0 := int32(load32(m.memory[uint32(v0):]))
		t1 := int32(load32(m.memory[uint32(v1):]))
		if t0 == t1 {
			goto l1
		}
		return i32(0)
	}
l1:
	v1 = v1 + i32(4)
	v0 = v0 + i32(4)
	v2 = v2 + i32(-4)
	if uint32(v2) > uint32(i32(3)) {
		goto l2
	}
l0:
	v3 = i32(1)
	if uint32(v2) <= uint32(i32(1)) {
		goto l3
	}
	{
		t2 := int32(load16(m.memory[uint32(v0):]))
		t3 := int32(load16(m.memory[uint32(v1):]))
		if t2 == t3 {
			goto l4
		}
		return i32(0)
	}
l4:
	v2 = v2 + i32(-2)
	v1 = v1 + i32(2)
	v0 = v0 + i32(2)
l3:
	{
		if v2 == 0 {
			goto l5
		}
		t4 := int32(m.memory[uint32(v0)])
		t5 := int32(m.memory[uint32(v1)])
		var p6 int32
		if t4 == t5 {
			p6 = 1
		}
		v3 = p6
	}
l5:
	return v3
}
func (m *Module) fn909(v0, v1, v2, v3 int32) {
	var v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14, v15, v16, v17, v18 int32
	v4 = v3 + i32(-4)
	v5 = v3 + i32(-8)
	v6 = v2 + v3
	v7 = v6 + i32(-8)
	t0 := int32(m.memory[int64(uint32(v1))+6])
	v8 = t0
	v9 = v8 * i32(16843009)
	t1 := int32(m.memory[int64(uint32(v1))+5])
	v10 = t1
	t2 := int32(m.memory[int64(uint32(v1))+4])
	v11 = t2
	v12 = i32(0)
	t3 := int32(m.memory[int64(uint32(v1))+7])
	v13 = t3 & i32(255)
	v14 = i32(0)
l14:
	{
		if uint32(v3) > uint32(v14) {
			goto l0
		}
		goto l1
	l0:
		v15 = v2 + v14
		v16 = v3 - v14
		if uint32(v16) > uint32(i32(3)) {
			t5 := int32(load32(m.memory[uint32(v15):]))
			v1 = t5 ^ v9
			if (i32(16843008)-v1|v1)&i32(-2139062144) == i32(-2139062144) {
				v17 = v15 & i32(3)
				v1 = i32(4) - v17
				v18 = v1 + v14
				v1 = v15 + v1
				if uint32(v16) < uint32(i32(9)) {
					if uint32(v18) < uint32(v3) {
						v16 = v4 + v17 - v14
					l11:
						{
							t9 := int32(m.memory[uint32(v1)])
							if v8 == t9 {
								goto l3
							}
							v1 = v1 + i32(1)
							v16 = v16 + i32(-1)
							if v16 != 0 {
								goto l11
							}
						}
						goto l1
					}
					goto l1
				}
				if v18 > v5 {
					goto l8
				}
			l9:
				{
					t7 := int32(load32(m.memory[uint32(v1):]))
					v16 = t7 ^ v9
					if (i32(16843008)-v16|v16)&i32(-2139062144) != i32(-2139062144) {
						goto l8
					}
					t8 := int32(load32(m.memory[uint32(v1+i32(4)):]))
					v16 = t8 ^ v9
					if (i32(16843008)-v16|v16)&i32(-2139062144) != i32(-2139062144) {
						goto l8
					}
					v1 = v1 + i32(8)
					if uint32(v1) <= uint32(v7) {
						goto l9
					}
					goto l8
				}
			l8:
				if uint32(v1) < uint32(v6) {
				l13:
					{
						t10 := int32(m.memory[uint32(v1)])
						if v8 == t10 {
							goto l3
						}
						v1 = v1 + i32(1)
						if v1 != v6 {
							goto l13
						}
					}
					goto l1
				}
				goto l1
			}
			v1 = v15
		l6:
			{
				t6 := int32(m.memory[uint32(v1)])
				if v8 == t6 {
					goto l3
				}
				v1 = v1 + i32(1)
				v16 = v16 + i32(-1)
				if v16 != 0 {
					goto l6
				}
			}
			goto l1
		}
		v1 = v15
	l4:
		{
			t4 := int32(m.memory[uint32(v1)])
			if v8 == t4 {
				goto l3
			}
			v1 = v1 + i32(1)
			v16 = v16 + i32(-1)
			if v16 != 0 {
				goto l4
			}
		}
		goto l1
	l3:
		v1 = v1 - v15 + v14
		v14 = v1 + i32(1)
		if uint32(v1) < uint32(v11) {
			goto l14
		}
		v1 = v1 - v11
		v16 = v1 + v10
		if uint32(v16) < uint32(v1) {
			goto l14
		}
		if uint32(v16) >= uint32(v3) {
			goto l14
		}
		t11 := int32(m.memory[uint32(v2+v16)])
		if t11 != v13 {
			goto l14
		}
	}
	v12 = i32(1)
l1:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
	store32(m.memory[uint32(v0):], uint32(v12))
}
func (m *Module) fn910(v0, v1, v2, v3, v4, v5, v6 int32) {
	var v7, v8, v9, v10, v11, v12, v13, v14, v15, v16 int32
	var v17 int64
	var v18, v19, v20 int32
	t0 := m.g0
	v7 = t0 - i32(16)
	m.g0 = v7
	{
		if uint32(v4) < uint32(i32(16)) {
			if uint32(v6) > uint32(v4) {
				goto l4
			}
			v10 = v3 + v4
			if v6 != 0 {
				goto l5
			}
			v11 = i32(0)
			goto l6
		l5:
			v4 = v3 + v6
			v11 = i32(0)
			v12 = v3
		l7:
			{
				t3 := int32(m.memory[uint32(v12)])
				v11 = v11<<1 + t3
				v12 = v12 + i32(1)
				if uint32(v12) < uint32(v4) {
					goto l7
				}
			}
		l6:
			t4 := int32(load32(m.memory[int64(uint32(v1))+44:]))
			v13 = t4
			t5 := int32(load32(m.memory[int64(uint32(v1))+40:]))
			v4 = t5
			v1 = v10 - v6
			v12 = v3
		l10:
			{
				{
					if v4 != v11 {
						goto l8
					}
					t6 := m.fn908(v12, v5, v6)
					if t6 == 0 {
						goto l8
					}
					v11 = v12 - v3
					v1 = i32(1)
					goto l9
				}
			l8:
				if uint32(v12) >= uint32(v1) {
					goto l4
				}
				t7 := int32(m.memory[uint32(v12)])
				t8 := int32(m.memory[uint32(v12+v6)])
				v11 = (v11-v13*t7)<<1 + t8
				v12 = v12 + i32(1)
				goto l10
			}
		}
		v8 = v1 + i32(24)
		t1 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		v9 = t1
		t2 := int32(load32(m.memory[uint32(v1):]))
		if t2 == 0 {
			if v6 == 0 {
				goto l3
			}
			if uint32(v6) > uint32(v4) {
				goto l4
			}
			v14 = v6 + i32(-1)
			v15 = v6 - v9
			t9 := int32(load32(m.memory[int64(uint32(v2))+4:]))
			v16 = t9
			t10 := int64(load64(m.memory[int64(uint32(v1))+8:]))
			v17 = t10
			t11 := int32(load32(m.memory[int64(uint32(v1))+24:]))
			v18 = t11
			t12 := int32(load32(m.memory[int64(uint32(v1))+16:]))
			v19 = t12
			t13 := int32(load32(m.memory[uint32(v2):]))
			v20 = t13
			v11 = i32(0)
			v1 = i32(0)
		l31:
			{
				p14 := v19
				if uint32(v1) > uint32(v19) {
					p14 = v1
				}
				v10 = p14
				if v20 != 0 {
					if uint32(v20) < uint32(i32(51)) {
						goto l13
					}
					if uint32(v16) >= uint32(v20<<3+i32(-8)) {
						goto l13
					}
					v20 = i32(0)
					store32(m.memory[uint32(v2):], uint32(i32(0)))
					goto l12
				l13:
					{
						if uint32(v4) < uint32(v11) {
							m.fn120(v11, v4, v4, i32(1271852))
							panic("unreachable")
						}
						t15 := v7 + i32(8)
						t16 := v8
						t17 := v3 + v11
						v12 = v4 - v11
						m.t0[uint(v18)].(func(int32, int32, int32, int32))(t15, t16, t17, v12)
						{
							t18 := int32(load32(m.memory[int64(uint32(v7))+8:]))
							if t18&i32(1) == 0 {
								t24 := v2
								v11 = v20 + i32(1)
								p25 := i32(-1)
								if v11 != 0 {
									p25 = v11
								}
								store32(m.memory[uint32(t24):], uint32(p25))
								t26 := v2
								v12 = v16 + v12
								p27 := v12
								if uint32(v12) < uint32(v16) {
									p27 = i32(-1)
								}
								store32(m.memory[int64(uint32(t26))+4:], uint32(p27))
								v1 = i32(0)
								goto l9
							}
							t19 := int32(load32(m.memory[int64(uint32(v7))+12:]))
							v12 = t19
							t20 := v2
							v1 = v20 + i32(1)
							p21 := i32(-1)
							if v1 != 0 {
								p21 = v1
							}
							v20 = p21
							store32(m.memory[uint32(t20):], uint32(v20))
							t22 := v2
							v1 = v16 + v12
							p23 := v1
							if uint32(v1) < uint32(v16) {
								p23 = i32(-1)
							}
							v16 = p23
							store32(m.memory[int64(uint32(t22))+4:], uint32(v16))
							v1 = i32(0)
							v10 = v19
							v11 = v12 + v11
							if uint32(v11+v6) <= uint32(v4) {
								goto l12
							}
							goto l9
						}
					}
				}
				v20 = i32(0)
				goto l12
			l12:
				v12 = v14 + v11
				if uint32(v12) < uint32(v4) {
					{
						t28 := int64(m.memory[uint32(v3+v12)])
						if i64_shr_u(v17, t28)&i64(1) == 0 {
							goto l17
						}
						if uint32(v10) >= uint32(v6) {
							goto l18
						}
						v13 = v3 + v11
						v12 = v10
					l21:
						{
							if uint32(v11+v12) >= uint32(v4) {
								t36 := v4
								v12 = v10 + v11
								p37 := v12
								if uint32(v4) > uint32(v12) {
									p37 = t36
								}
								m.fn32(p37, v4, i32(1271772))
								panic("unreachable")
							}
							t29 := int32(m.memory[uint32(v5+v12)])
							t30 := int32(m.memory[uint32(v13+v12)])
							if t29 != t30 {
								goto l20
							}
							t31 := v6
							v12 = v12 + i32(1)
							if t31 != v12 {
								goto l21
							}
						}
					l18:
						if uint32(v19) <= uint32(v1) {
							goto l22
						}
						v13 = v3 + v11
						v12 = v19
					l27:
						if uint32(v12) >= uint32(v6) {
							m.fn32(v12, v6, i32(1271788))
							panic("unreachable")
						}
						{
							v10 = v11 + v12
							if uint32(v10) >= uint32(v4) {
								m.fn32(v10, v4, i32(1271804))
								panic("unreachable")
							}
							t32 := int32(m.memory[uint32(v5+v12)])
							t33 := int32(m.memory[uint32(v13+v12)])
							if t32 == t33 {
								goto l25
							}
							v1 = v15
							v11 = v9 + v11
							goto l26
						}
					l25:
						v12 = v12 + i32(-1)
						if uint32(v12) > uint32(v1) {
							goto l27
						}
					l22:
						if uint32(v1) >= uint32(v6) {
							m.fn32(v1, v6, i32(1271820))
							panic("unreachable")
						}
						{
							v13 = v11 + v1
							if uint32(v13) >= uint32(v4) {
								m.fn32(v13, v4, i32(1271836))
								panic("unreachable")
							}
							v10 = v5 + v1
							v1 = v15
							v12 = v9
							t34 := int32(m.memory[uint32(v10)])
							t35 := int32(m.memory[uint32(v3+v13)])
							if t34 != t35 {
								goto l30
							}
							v1 = i32(1)
							goto l9
						}
					l20:
						v12 = v12 - v19 + i32(1)
						v1 = i32(0)
					l30:
						v11 = v12 + v11
						goto l26
					}
				l17:
					v11 = v11 + v6
					v1 = i32(0)
				l26:
					if uint32(v11+v6) > uint32(v4) {
						goto l4
					}
					goto l31
				}
				m.fn32(v12, v4, i32(1271756))
				panic("unreachable")
			}
		}
		if v6 != 0 {
			if uint32(v6) > uint32(v4) {
				goto l4
			}
			v16 = v6 + i32(-1)
			t38 := int32(load32(m.memory[int64(uint32(v2))+4:]))
			v19 = t38
			t39 := int64(load64(m.memory[int64(uint32(v1))+8:]))
			v17 = t39
			t40 := int32(load32(m.memory[int64(uint32(v1))+24:]))
			v14 = t40
			t41 := int32(load32(m.memory[uint32(v2):]))
			v10 = t41
			v11 = i32(0)
			t42 := int32(load32(m.memory[int64(uint32(v1))+16:]))
			v20 = t42
			var p43 int32
			if uint32(v20+i32(-1)) >= uint32(v6) {
				p43 = 1
			}
			v13 = p43
		l48:
			if v10 != 0 {
				if uint32(v10) < uint32(i32(51)) {
					goto l34
				}
				if uint32(v19) >= uint32(v10<<3+i32(-8)) {
					goto l34
				}
				v10 = i32(0)
				store32(m.memory[uint32(v2):], uint32(i32(0)))
				goto l33
			l34:
				{
					if uint32(v4) < uint32(v11) {
						m.fn120(v11, v4, v4, i32(1271740))
						panic("unreachable")
					}
					t44 := v7
					t45 := v8
					t46 := v3 + v11
					v12 = v4 - v11
					m.t0[uint(v14)].(func(int32, int32, int32, int32))(t44, t45, t46, v12)
					{
						t47 := int32(load32(m.memory[uint32(v7):]))
						if t47&i32(1) == 0 {
							t53 := v2
							v11 = v10 + i32(1)
							p54 := i32(-1)
							if v11 != 0 {
								p54 = v11
							}
							store32(m.memory[uint32(t53):], uint32(p54))
							t55 := v2
							v12 = v19 + v12
							p56 := v12
							if uint32(v12) < uint32(v19) {
								p56 = i32(-1)
							}
							store32(m.memory[int64(uint32(t55))+4:], uint32(p56))
							v1 = i32(0)
							goto l9
						}
						t48 := int32(load32(m.memory[int64(uint32(v7))+4:]))
						v12 = t48
						t49 := v2
						v1 = v10 + i32(1)
						p50 := i32(-1)
						if v1 != 0 {
							p50 = v1
						}
						v10 = p50
						store32(m.memory[uint32(t49):], uint32(v10))
						t51 := v2
						v1 = v19 + v12
						p52 := v1
						if uint32(v1) < uint32(v19) {
							p52 = i32(-1)
						}
						v19 = p52
						store32(m.memory[int64(uint32(t51))+4:], uint32(v19))
						v11 = v12 + v11
						if uint32(v11+v6) > uint32(v4) {
							goto l4
						}
						goto l33
					}
				}
			}
			v10 = i32(0)
			goto l33
		l33:
			v12 = v16 + v11
			if uint32(v12) >= uint32(v4) {
				m.fn32(v12, v4, i32(1271676))
				panic("unreachable")
			}
			v1 = v6
			{
				t57 := int64(m.memory[uint32(v3+v12)])
				if i64_shr_u(v17, t57)&i64(1) == 0 {
					goto l38
				}
				v12 = v20
				if uint32(v20) >= uint32(v6) {
					goto l46
				}
				v1 = v3 + v11
				v12 = v20
			l42:
				{
					if uint32(v11+v12) >= uint32(v4) {
						t63 := v4
						v12 = v20 + v11
						p64 := v12
						if uint32(v4) > uint32(v12) {
							p64 = t63
						}
						m.fn32(p64, v4, i32(1271692))
						panic("unreachable")
					}
					t58 := int32(m.memory[uint32(v5+v12)])
					t59 := int32(m.memory[uint32(v1+v12)])
					if t58 != t59 {
						v11 = v11 - v20 + v12 + i32(1)
						goto l47
					}
					t60 := v6
					v12 = v12 + i32(1)
					if t60 != v12 {
						goto l42
					}
				}
				v12 = v20
			l46:
				{
					if v12 != 0 {
						goto l43
					}
					v1 = i32(1)
					goto l9
				l43:
					v12 = v12 + i32(-1)
					if v13 != 0 {
						m.fn32(v12, v6, i32(1271708))
						panic("unreachable")
					}
					v1 = v12 + v11
					if uint32(v1) >= uint32(v4) {
						m.fn32(v1, v4, i32(1271724))
						panic("unreachable")
					}
					t61 := int32(m.memory[uint32(v5+v12)])
					t62 := int32(m.memory[uint32(v3+v1)])
					if t61 == t62 {
						goto l46
					}
				}
				v1 = v9
				goto l38
			}
		l38:
			v11 = v1 + v11
		l47:
			v1 = i32(0)
			if uint32(v11+v6) <= uint32(v4) {
				goto l48
			}
			goto l9
		}
		goto l3
	}
l4:
	v1 = i32(0)
	goto l9
l3:
	v11 = i32(0)
	v1 = i32(1)
l9:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v11))
	store32(m.memory[uint32(v0):], uint32(v1))
	m.g0 = v7 + i32(16)
}
func (m *Module) fn911(v0 int32) int32 {
	var v1, v2, v3, v4, v5, v6, v7, v8, v9 int32
	var v10 int64
	var v11, v12, v13, v14 int32
	var v15 int64
	var v16 int32
	var v17 int64
	{
		{
			{
				t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
				v1 = t0
				v2 = v1 + i32(1)
				if v2 == 0 {
					m.fn27(i32(1271632), i32(57), i32(1271660))
					panic("unreachable")
				}
				{
					t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
					t2 := v2
					v3 = t1
					t3 := v3
					v4 = v3 + i32(1)
					v5 = int32(uint32(v4) >> 3)
					p4 := v5 * i32(7)
					if uint32(v3) < uint32(i32(8)) {
						p4 = t3
					}
					v6 = p4
					if uint32(t2) <= uint32(int32(uint32(v6)>>1)) {
						{
							if v4 == 0 {
								goto l6
							}
							t7 := int32(load32(m.memory[uint32(v0):]))
							v7 = t7
							v2 = i32(0)
							{
								{
									t8 := v5
									var p9 int32
									if v4&i32(7) != i32(0) {
										p9 = 1
									}
									v5 = t8 + p9
									if v5 == i32(1) {
										goto l7
									}
									v8 = v5 & i32(1)
									v9 = v5 & i32(0x3ffffffe)
									v2 = i32(0)
								l8:
									{
										v5 = v7 + v2
										t10 := int64(load64(m.memory[uint32(v5):]))
										t11 := v5
										v10 = t10
										store64(m.memory[uint32(t11):], uint64(int64(uint64(v10^i64(-1))>>7)&i64(72340172838076673)+(v10|i64(0x7f7f7f7f7f7f7f7f))))
										v5 = v5 + i32(8)
										t12 := int64(load64(m.memory[uint32(v5):]))
										t13 := v5
										v10 = t12
										store64(m.memory[uint32(t13):], uint64(int64(uint64(v10^i64(-1))>>7)&i64(72340172838076673)+(v10|i64(0x7f7f7f7f7f7f7f7f))))
										v2 = v2 + i32(16)
										v9 = v9 + i32(-2)
										if v9 != 0 {
											goto l8
										}
									}
									if v8 == 0 {
										goto l9
									}
								}
							l7:
								v2 = v7 + v2
								t14 := int64(load64(m.memory[uint32(v2):]))
								t15 := v2
								v10 = t14
								store64(m.memory[uint32(t15):], uint64(int64(uint64(v10^i64(-1))>>7)&i64(72340172838076673)+(v10|i64(0x7f7f7f7f7f7f7f7f))))
							}
						l9:
							{
								if uint32(v4) < uint32(i32(8)) {
									goto l10
								}
								t16 := int64(load64(m.memory[uint32(v7):]))
								store64(m.memory[uint32(v7+v4):], uint64(t16))
								goto l11
							}
						l10:
							if v4 == 0 {
								goto l11
							}
							memory_copy(m.memory, uint32(v7+i32(8)), uint32(v7), uint32(v4))
						l11:
							v5 = i32(0)
						l19:
							{
								t17 := v7
								v2 = v5
								v9 = t17 + v2
								t18 := int32(m.memory[uint32(v9)])
								if t18 != i32(128) {
									goto l12
								}
								v11 = v7 - v2<<3 + i32(-8)
								v12 = v7 + (v2^i32(-1))<<3
								{
								l18:
									{
										t19 := int64(load64(m.memory[uint32(v11):]))
										t20 := v3
										v13 = int32(t19)
										v5 = t20 & v13
										v8 = v5
										{
											t21 := int64(load64(m.memory[uint32(v7+v5):]))
											v10 = t21 & i64(-0x7f7f7f7f7f7f7f80)
											if v10 != i64(0) {
												goto l13
											}
											v4 = i32(8)
											v8 = v5
										l14:
											{
												v8 = v8 + v4
												v4 = v4 + i32(8)
												t22 := v7
												v8 = v8 & v3
												t23 := int64(load64(m.memory[uint32(t22+v8):]))
												v10 = t23 & i64(-0x7f7f7f7f7f7f7f80)
												if v10 == 0 {
													goto l14
												}
											}
										}
									l13:
										{
											t24 := v7
											v8 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v10))))>>3) + v8) & v3
											t25 := int32(int8(m.memory[uint32(t24+v8)]))
											if t25 < i32(0) {
												goto l15
											}
											t26 := int64(load64(m.memory[uint32(v7):]))
											v8 = int32(uint32(int64(bits.TrailingZeros64(uint64(t26&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
										}
									l15:
										{
											if uint32((v8-v5^(v2-v5))&v3) < uint32(i32(8)) {
												goto l16
											}
											v5 = v7 + v8
											t27 := int32(m.memory[uint32(v5)])
											v4 = t27
											t28 := v5
											v13 = int32(uint32(v13) >> 25)
											m.memory[uint32(t28)] = byte(v13)
											m.memory[uint32(v7+(v8+i32(-8))&v3+i32(8))] = byte(v13)
											v5 = v7 - v8<<3 + i32(-8)
											if v4 == i32(255) {
												goto l17
											}
											t29 := int32(load32(m.memory[uint32(v5):]))
											v8 = t29
											t30 := int32(load32(m.memory[uint32(v12):]))
											store32(m.memory[uint32(v5):], uint32(t30))
											store32(m.memory[uint32(v12):], uint32(v8))
											t31 := int32(load32(m.memory[int64(uint32(v12))+4:]))
											v8 = t31
											t32 := int32(load32(m.memory[int64(uint32(v5))+4:]))
											store32(m.memory[int64(uint32(v12))+4:], uint32(t32))
											store32(m.memory[int64(uint32(v5))+4:], uint32(v8))
											goto l18
										}
									l16:
									}
									t33 := v9
									v5 = int32(uint32(v13) >> 25)
									m.memory[uint32(t33)] = byte(v5)
									m.memory[uint32(v7+(v2+i32(-8))&v3+i32(8))] = byte(v5)
									goto l12
								}
							l17:
								m.memory[uint32(v9)] = byte(i32(255))
								m.memory[uint32(v7+(v2+i32(-8))&v3+i32(8))] = byte(i32(255))
								t34 := int64(load64(m.memory[uint32(v12):]))
								store64(m.memory[uint32(v5):], uint64(t34))
							}
						l12:
							v5 = v2 + i32(1)
							if v2 != v3 {
								goto l19
							}
						}
					l6:
						store32(m.memory[int64(uint32(v0))+8:], uint32(v6-v1))
						goto l20
					}
					v7 = v6 + i32(1)
					p5 := v2
					if uint32(v7) > uint32(v2) {
						p5 = v7
					}
					v2 = p5
					if uint32(v2) < uint32(i32(15)) {
						goto l2
					}
					{
						if uint32(v2) > uint32(i32(0x1fffffff)) {
							m.fn27(i32(1271632), i32(57), i32(1271660))
							panic("unreachable")
						}
						t6 := int32(uint32(v2<<3) / uint32(i32(7)))
						v2 = i32_shr_u(i32(-1), int32(bits.LeadingZeros32(uint32(t6+i32(-1)))))
						if uint32(v2) > uint32(i32(0x1ffffffe)) {
							goto l4
						}
						v2 = v2 + i32(1)
						goto l5
					}
				}
			}
		l2:
			p35 := v2&i32(8) + i32(8)
			if uint32(v2) < uint32(i32(4)) {
				p35 = i32(4)
			}
			v2 = p35
		}
	l5:
		v7 = v2 + i32(8)
		t36 := v7
		v9 = v2 << 3
		v5 = t36 + v9
		if uint32(v5) < uint32(v7) {
			goto l4
		}
		if uint32(v5) > uint32(i32(0x7ffffff8)) {
			goto l4
		}
		{
			t37 := m.fn7(v5)
			v8 = t37
			if v8 != 0 {
				v5 = v8 + v9
				if v7 == 0 {
					goto l22
				}
				memory_fill(m.memory, uint32(v5), i32(255), uint32(v7))
			l22:
				v8 = v2 + i32(-1)
				p38 := int32(uint32(v2)>>3) * i32(7)
				if uint32(v2) < uint32(i32(9)) {
					p38 = v8
				}
				v14 = p38
				t39 := int32(load32(m.memory[uint32(v0):]))
				v6 = t39
				{
					if v1 == 0 {
						goto l23
					}
					t40 := int64(load64(m.memory[uint32(v6):]))
					v10 = (t40 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
					v7 = v6
					v2 = i32(0)
					v13 = v1
				l29:
					{
						if v10 != i64(0) {
							goto l24
						}
					l25:
						{
							v2 = v2 + i32(8)
							v7 = v7 + i32(8)
							t41 := int64(load64(m.memory[uint32(v7):]))
							v10 = t41 & i64(-0x7f7f7f7f7f7f7f80)
							if v10 == i64(-0x7f7f7f7f7f7f7f80) {
								goto l25
							}
						}
						v10 = v10 ^ i64(-0x7f7f7f7f7f7f7f80)
					l24:
						{
							t42 := v5
							t43 := v8
							v11 = v6 - (int32(uint32(int64(bits.TrailingZeros64(uint64(v10))))>>3)+v2)<<3 + i32(-8)
							t44 := int64(load64(m.memory[uint32(v11):]))
							v12 = int32(t44)
							v9 = t43 & v12
							t45 := int64(load64(m.memory[uint32(t42+v9):]))
							v15 = t45 & i64(-0x7f7f7f7f7f7f7f80)
							if v15 != i64(0) {
								goto l26
							}
							v16 = i32(8)
						l27:
							{
								v9 = v9 + v16
								v16 = v16 + i32(8)
								t46 := v5
								v9 = v9 & v8
								t47 := int64(load64(m.memory[uint32(t46+v9):]))
								v15 = t47 & i64(-0x7f7f7f7f7f7f7f80)
								if v15 == 0 {
									goto l27
								}
							}
						}
					l26:
						v17 = v10 + i64(-1)
						{
							t48 := v5
							v9 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v15))))>>3) + v9) & v8
							t49 := int32(int8(m.memory[uint32(t48+v9)]))
							if t49 < i32(0) {
								goto l28
							}
							t50 := int64(load64(m.memory[uint32(v5):]))
							v9 = int32(uint32(int64(bits.TrailingZeros64(uint64(t50&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
						}
					l28:
						v10 = v17 & v10
						t51 := v5 + v9
						v12 = int32(uint32(v12) >> 25)
						m.memory[uint32(t51)] = byte(v12)
						m.memory[uint32(v5+(v9+i32(-8))&v8+i32(8))] = byte(v12)
						t52 := int64(load64(m.memory[uint32(v11):]))
						store64(m.memory[uint32(v5-v9<<3+i32(-8)):], uint64(t52))
						v13 = v13 + i32(-1)
						if v13 != 0 {
							goto l29
						}
					}
				}
			l23:
				store32(m.memory[int64(uint32(v0))+4:], uint32(v8))
				store32(m.memory[uint32(v0):], uint32(v5))
				store32(m.memory[int64(uint32(v0))+8:], uint32(v14-v1))
				if v3 == 0 {
					goto l20
				}
				t53 := v3
				v7 = v4 << 3
				v2 = t53 + v7 + i32(9)
				if v2 == 0 {
					goto l20
				}
				v3 = v6 - v7
				t54 := int32(load32(m.memory[uint32(v3+i32(-4)):]))
				v7 = t54
				v5 = v7 & i32(-8)
				t55 := v5
				v7 = v7 & i32(3)
				p56 := i32(8)
				if v7 != 0 {
					p56 = i32(4)
				}
				if uint32(t55) < uint32(p56+v2) {
					m.fn3(i32(1274224), i32(46), i32(1274272))
					panic("unreachable")
				}
				if v7 == 0 {
					goto l31
				}
				if uint32(v5) > uint32(v2+i32(39)) {
					m.fn3(i32(1274288), i32(46), i32(1274336))
					panic("unreachable")
				}
			l31:
				m.fn1(v3)
				return i32(-1)
			}
			m.fn23(i32(8), v5)
			panic("unreachable")
		}
	}
l20:
	return i32(-1)
l4:
	m.fn27(i32(1271632), i32(57), i32(1271660))
	panic("unreachable")
}
func (m *Module) fn912(v0, v1, v2, v3, v4, v5 int32) {
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
func (m *Module) fn913(v0, v1, v2, v3, v4 int32) int32 {
	var v5, v6 int32
	var v7 int64
	t0 := m.g0
	v5 = t0 - i32(48)
	m.g0 = v5
	store32(m.memory[int64(uint32(v5))+4:], uint32(v3))
	{
		{
			if uint32(v3) >= uint32(v2) {
				m.fn32(v3, v2, i32(1272776))
				panic("unreachable")
			}
			t1 := int32(m.memory[uint32(v1+v3)])
			v6 = t1
			switch v6 + i32(-10) {
			case 0:
				goto l1
			case 3:
				{
					t4 := int32(load32(m.memory[uint32(v0):]))
					t5 := int32(load32(m.memory[int64(uint32(v0))+8:]))
					v6 = t5
					if t4 != v6 {
						goto l4
					}
					m.fn244(v0, v6, i32(1))
				}
			l4:
				store32(m.memory[int64(uint32(v0))+8:], uint32(v6+i32(1)))
				t6 := int32(load32(m.memory[int64(uint32(v0))+4:]))
				m.memory[uint32(t6+v6)] = byte(v4)
				v0 = v3 + i32(1)
				if uint32(v0) >= uint32(v2) {
					goto l5
				}
				t7 := int32(m.memory[uint32(v1+v0)])
				p8 := v0
				if t7 == i32(10) {
					p8 = v3 + i32(2)
				}
				v0 = p8
				goto l5
			default:
				m.memory[int64(uint32(v5))+11] = byte(v6)
				store32(m.memory[int64(uint32(v5))+12:], uint32(v6))
				t2 := v5
				t3 := int64(uint32(i32(56))) << 32
				v7 = int64(uint32(v5 + i32(11)))
				store64(m.memory[int64(uint32(t2))+40:], uint64(t3|v7))
				store64(m.memory[int64(uint32(v5))+32:], uint64(int64(uint32(i32(18)))<<32|v7))
				store64(m.memory[int64(uint32(v5))+24:], uint64(int64(uint32(i32(44)))<<32|int64(uint32(v5+i32(12)))))
				store64(m.memory[int64(uint32(v5))+16:], uint64(int64(uint32(i32(3)))<<32|int64(uint32(v5+i32(4)))))
				m.fn27(i32(1066198), v5+i32(16), i32(1272792))
				panic("unreachable")
			}
		}
	l1:
		{
			t9 := int32(load32(m.memory[uint32(v0):]))
			t10 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			v2 = t10
			if t9 != v2 {
				goto l6
			}
			m.fn244(v0, v2, i32(1))
		}
	l6:
		store32(m.memory[int64(uint32(v0))+8:], uint32(v2+i32(1)))
		t11 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		m.memory[uint32(t11+v2)] = byte(v4)
		v0 = v3 + i32(1)
	}
l5:
	m.g0 = v5 + i32(48)
	return v0
}
func (m *Module) fn914(v0 int32) {
	var v1 int32
	{
		t0 := int32(load32(m.memory[uint32(v0):]))
		t1 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		v1 = t1
		if t0 != v1 {
			goto l0
		}
		m.fn244(v0, v1, i32(1))
	}
l0:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v1+i32(1)))
	t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	m.memory[uint32(t2+v1)] = byte(i32(38))
}
func (m *Module) fn915(v0, v1, v2 int32) {
	var v3 int32
	{
		t0 := int32(load32(m.memory[uint32(v0):]))
		t1 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		t2 := v2
		v3 = t1
		if uint32(t2) <= uint32(t0-v3) {
			goto l0
		}
		m.fn244(v0, v3, v2)
		t3 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		v3 = t3
	}
l0:
	{
		if v2 == 0 {
			goto l1
		}
		t4 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		memory_copy(m.memory, uint32(t4+v3), uint32(v1), uint32(v2))
	}
l1:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v3+v2))
}
func (m *Module) fn916(v0, v1, v2, v3 int32) {
	var v4, v5, v6, v7 int32
	v4 = i32(0)
	if v3 != 0 {
		goto l0
	}
	goto l1
l0:
	{
		t0 := int32(m.memory[uint32(v2)])
		v5 = t0
		if v5 != i32(63) {
			goto l2
		}
		store16(m.memory[uint32(v1):], uint16(i32(4)))
		v5 = i32(1)
		v4 = i32(1)
		goto l1
	}
l2:
	{
		if uint32(v3) < uint32(i32(3)) {
			goto l3
		}
		if v5 != i32(33) {
			goto l3
		}
		t1 := int32(m.memory[int64(uint32(v2))+2])
		v5 = t1
		{
			t2 := int32(m.memory[int64(uint32(v2))+1])
			v6 = t2
			if v6 != i32(45) {
				goto l4
			}
			if v5 == i32(45) {
				v5 = i32(3)
				store16(m.memory[uint32(v1):], uint16(i32(3)))
				v4 = i32(1)
				goto l1
			}
		}
	l4:
		{
			if uint32(v3) > uint32(i32(7)) {
				switch v6 + i32(-65) {
				case 4:
					switch v5 + i32(-76) {
					default:
						goto l7
					case 0:
						t7 := int32(m.memory[int64(uint32(v2))+3])
						if t7 != i32(69) {
							goto l7
						}
						t8 := int32(m.memory[int64(uint32(v2))+4])
						if t8 != i32(77) {
							goto l7
						}
						t9 := int32(m.memory[int64(uint32(v2))+5])
						if t9 != i32(69) {
							goto l7
						}
						t10 := int32(m.memory[int64(uint32(v2))+6])
						if t10 != i32(78) {
							goto l7
						}
						t11 := int32(m.memory[int64(uint32(v2))+7])
						if t11 != i32(84) {
							goto l7
						}
						m.memory[uint32(v1)] = byte(i32(5))
						goto l14
					case 2:
						t12 := int32(m.memory[int64(uint32(v2))+3])
						if t12 != i32(84) {
							goto l7
						}
						t13 := int32(m.memory[int64(uint32(v2))+4])
						if t13 != i32(73) {
							goto l7
						}
						t14 := int32(m.memory[int64(uint32(v2))+5])
						if t14 != i32(84) {
							goto l7
						}
						t15 := int32(m.memory[int64(uint32(v2))+6])
						if t15 != i32(89) {
							goto l7
						}
						goto l8
					}
				default:
					if v3 == i32(8) {
						goto l7
					}
					if v6 != i32(78) {
						goto l7
					}
					if v5 != i32(79) {
						goto l7
					}
					t16 := int32(m.memory[int64(uint32(v2))+3])
					if t16 != i32(84) {
						goto l7
					}
					t17 := int32(m.memory[int64(uint32(v2))+4])
					if t17 != i32(65) {
						goto l7
					}
					t18 := int32(m.memory[int64(uint32(v2))+5])
					if t18 != i32(84) {
						goto l7
					}
					t19 := int32(m.memory[int64(uint32(v2))+6])
					if t19 != i32(73) {
						goto l7
					}
					t20 := int32(m.memory[int64(uint32(v2))+7])
					if t20 != i32(79) {
						goto l7
					}
					t21 := int32(m.memory[int64(uint32(v2))+8])
					if t21 != i32(78) {
						goto l7
					}
					store16(m.memory[uint32(v1):], uint16(i32(6)))
					v4 = i32(1)
					v5 = i32(9)
					goto l1
				case 0:
					if v5 != i32(84) {
						goto l7
					}
					t22 := int32(m.memory[int64(uint32(v2))+3])
					if t22&i32(255) != i32(84) {
						goto l7
					}
					t23 := int32(m.memory[int64(uint32(v2))+4])
					if t23 != i32(76) {
						goto l7
					}
					t24 := int32(m.memory[int64(uint32(v2))+5])
					if t24 != i32(73) {
						goto l7
					}
					t25 := int32(m.memory[int64(uint32(v2))+6])
					if t25 != i32(83) {
						goto l7
					}
					t26 := int32(m.memory[int64(uint32(v2))+7])
					if t26 != i32(84) {
						goto l7
					}
					store16(m.memory[uint32(v1):], uint16(i32(6)))
					goto l14
				}
			}
			if v3 != i32(7) {
				goto l3
			}
			if v6 != i32(69) {
				goto l7
			}
			if v5 != i32(78) {
				goto l7
			}
			t3 := int32(m.memory[int64(uint32(v2))+3])
			if t3 != i32(84) {
				goto l7
			}
			t4 := int32(m.memory[int64(uint32(v2))+4])
			if t4 != i32(73) {
				goto l7
			}
			t5 := int32(m.memory[int64(uint32(v2))+5])
			if t5 != i32(84) {
				goto l7
			}
			t6 := int32(m.memory[int64(uint32(v2))+6])
			if t6 != i32(89) {
				goto l7
			}
			goto l8
		}
	l14:
		v4 = i32(1)
		v5 = i32(8)
		goto l1
	l8:
		store16(m.memory[uint32(v1):], uint16(i32(6)))
		v4 = i32(1)
		v5 = i32(7)
		goto l1
	}
l3:
	if uint32(v3) <= uint32(i32(3)) {
		v5 = v2
	l22:
		{
			t33 := int32(m.memory[uint32(v5)])
			if t33 == i32(62) {
				goto l17
			}
			v5 = v5 + i32(1)
			v3 = v3 + i32(-1)
			if v3 != 0 {
				goto l22
			}
		}
		goto l1
	}
l7:
	v6 = v2 + v3
	{
		t27 := int32(load32(m.memory[uint32(v2):]))
		v5 = t27
		if (i32(16843008)-(v5^i32(1044266558))|v5)&i32(-2139062144) == i32(-2139062144) {
			t29 := v2
			v7 = i32(4) - v2&i32(3)
			v5 = t29 + v7
			if uint32(v3) < uint32(i32(9)) {
				if uint32(v7) < uint32(v3) {
				l24:
					{
						t34 := int32(m.memory[uint32(v5)])
						if t34 == i32(62) {
							goto l17
						}
						v5 = v5 + i32(1)
						if v5 != v6 {
							goto l24
						}
					}
					goto l1
				}
				goto l1
			}
			t30 := v7
			v3 = v3 + i32(-8)
			if uint32(t30) > uint32(v3) {
				goto l20
			}
			v7 = v3 + v2
		l21:
			{
				t31 := int32(load32(m.memory[uint32(v5):]))
				v3 = t31
				if (i32(16843008)-(v3^i32(1044266558))|v3)&i32(-2139062144) != i32(-2139062144) {
					goto l20
				}
				t32 := int32(load32(m.memory[uint32(v5+i32(4)):]))
				v3 = t32
				if (i32(16843008)-(v3^i32(1044266558))|v3)&i32(-2139062144) != i32(-2139062144) {
					goto l20
				}
				v5 = v5 + i32(8)
				if uint32(v5) <= uint32(v7) {
					goto l21
				}
				goto l20
			}
		}
		v5 = v2
	l18:
		{
			t28 := int32(m.memory[uint32(v5)])
			if t28 == i32(62) {
				goto l17
			}
			v5 = v5 + i32(1)
			if v5 != v6 {
				goto l18
			}
		}
		goto l1
	}
l20:
	if uint32(v5) < uint32(v6) {
	l26:
		{
			t35 := int32(m.memory[uint32(v5)])
			if t35 == i32(62) {
				goto l17
			}
			v5 = v5 + i32(1)
			if v5 != v6 {
				goto l26
			}
		}
		goto l1
	}
	goto l1
l17:
	v4 = i32(1)
	m.memory[uint32(v1)] = byte(i32(1))
	v5 = v5 - v2 + i32(1)
l1:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
	store32(m.memory[uint32(v0):], uint32(v4))
}
func (m *Module) fn917(v0, v1, v2 int32) {
	var v3, v4, v5, v6, v7, v8 int32
	v3 = i32(0)
	{
		t0 := v1
		v4 = v1 + v2
		if uint32(t0) >= uint32(v4) {
			goto l15
		}
		v5 = v4 + i32(-8)
		v6 = v1
	l14:
		v7 = v4 - v6
		if uint32(v7) <= uint32(i32(3)) {
		l8:
			{
				t5 := int32(m.memory[uint32(v6)])
				if t5 == i32(62) {
					goto l6
				}
				v6 = v6 + i32(1)
				if v6 != v4 {
					goto l8
				}
				goto l15
			}
		}
		{
			t1 := int32(load32(m.memory[uint32(v6):]))
			v8 = t1
			if (i32(16843008)-(v8^i32(1044266558))|v8)&i32(-2139062144) != i32(-2139062144) {
			l7:
				{
					t4 := int32(m.memory[uint32(v6)])
					if t4 == i32(62) {
						goto l6
					}
					v6 = v6 + i32(1)
					if v6 != v4 {
						goto l7
					}
					goto l15
				}
			}
			v6 = v6&i32(-4) + i32(4)
			if uint32(v7) < uint32(i32(9)) {
				if uint32(v6) >= uint32(v4) {
					goto l15
				}
			l9:
				{
					t6 := int32(m.memory[uint32(v6)])
					if t6 == i32(62) {
						goto l6
					}
					v6 = v6 + i32(1)
					if v6 != v4 {
						goto l9
					}
					goto l15
				}
			}
			if uint32(v6) > uint32(v5) {
				goto l4
			}
		l5:
			{
				t2 := int32(load32(m.memory[uint32(v6):]))
				v7 = t2
				if (i32(16843008)-(v7^i32(1044266558))|v7)&i32(-2139062144) != i32(-2139062144) {
					goto l4
				}
				t3 := int32(load32(m.memory[uint32(v6+i32(4)):]))
				v7 = t3
				if (i32(16843008)-(v7^i32(1044266558))|v7)&i32(-2139062144) != i32(-2139062144) {
					goto l4
				}
				v6 = v6 + i32(8)
				if uint32(v6) <= uint32(v5) {
					goto l5
				}
				goto l4
			}
		}
	l4:
		if uint32(v6) >= uint32(v4) {
			goto l15
		}
	l10:
		{
			t7 := int32(m.memory[uint32(v6)])
			if t7 == i32(62) {
				goto l6
			}
			v6 = v6 + i32(1)
			if v6 != v4 {
				goto l10
			}
			goto l15
		}
	l6:
		v7 = v6 - v1
		if uint32(v7) > uint32(v2) {
			m.fn120(i32(0), v7, v2, i32(1272812))
			panic("unreachable")
		}
		{
			if uint32(v7) < uint32(i32(2)) {
				goto l12
			}
			t8 := int32(load16(m.memory[uint32(v1+v7+i32(-2)):]))
			if t8 == i32(11565) {
				v3 = i32(1)
				v6 = v7 + i32(1)
				goto l15
			}
		}
	l12:
		v6 = v6 + i32(1)
		if uint32(v6) < uint32(v4) {
			goto l14
		}
		goto l15
	}
l15:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v6))
	store32(m.memory[uint32(v0):], uint32(v3))
}
func (m *Module) fn918(v0, v1 int32) {
	var v2, v3, v4, v5, v6, v7, v8, v9 int32
	t0 := m.g0
	v2 = t0 - i32(32)
	m.g0 = v2
	t1 := int32(load32(m.memory[int64(uint32(v1))+40:]))
	t2 := int32(load32(m.memory[int64(uint32(v1))+44:]))
	m.fn240(v2+i32(12), v1, t1, t2)
	{
		t3 := int32(load32(m.memory[int64(uint32(v2))+12:]))
		v3 = t3
		switch v3 + i32(2) {
		case 0:
			store32(m.memory[uint32(v0):], uint32(i32(-3)))
			goto l3
		case 1:
			t4 := int32(load32(m.memory[int64(uint32(v2))+24:]))
			store32(m.memory[int64(uint32(v0))+12:], uint32(t4))
			t5 := int64(load64(m.memory[int64(uint32(v2))+16:]))
			store64(m.memory[int64(uint32(v0))+4:], uint64(t5))
			store32(m.memory[uint32(v0):], uint32(i32(-2)))
			goto l3
		default:
			t6 := int32(load32(m.memory[int64(uint32(v2))+20:]))
			v4 = t6
			t7 := int32(load32(m.memory[int64(uint32(v2))+16:]))
			t8 := v4
			v5 = t7
			t9 := int32(load32(m.memory[int64(uint32(v1))+44:]))
			var p10 int32
			if uint32(t8) < uint32(v5) {
				p10 = 1
			}
			t11 := v4
			v6 = t9
			var p12 int32
			if uint32(t11) > uint32(v6) {
				p12 = 1
			}
			v7 = p10 | p12
			t13 := int32(load32(m.memory[int64(uint32(v1))+40:]))
			v8 = t13
			t14 := int32(load32(m.memory[int64(uint32(v2))+28:]))
			v1 = t14
			t15 := int32(load32(m.memory[int64(uint32(v2))+24:]))
			v9 = t15
			switch v3 {
			default:
				if v7 != 0 {
					m.fn120(v5, v4, v6, i32(1272152))
					panic("unreachable")
				}
				if uint32(v1) < uint32(v9) {
					goto l9
				}
				if uint32(v1) <= uint32(v6) {
					goto l10
				}
			l9:
				m.fn120(v9, v1, v6, i32(1272152))
				panic("unreachable")
			case 1:
				if v7 != 0 {
					m.fn120(v5, v4, v6, i32(1272152))
					panic("unreachable")
				}
				if uint32(v1) < uint32(v9) {
					goto l12
				}
				if uint32(v1) <= uint32(v6) {
					goto l10
				}
			l12:
				m.fn120(v9, v1, v6, i32(1272152))
				panic("unreachable")
			case 2:
				if v7 != 0 {
					m.fn120(v5, v4, v6, i32(1272152))
					panic("unreachable")
				}
				if uint32(v1) < uint32(v9) {
					goto l14
				}
				if uint32(v1) <= uint32(v6) {
					goto l10
				}
			l14:
				m.fn120(v9, v1, v6, i32(1272152))
				panic("unreachable")
			case 3:
				if v7 != 0 {
					m.fn120(v5, v4, v6, i32(1272152))
					panic("unreachable")
				}
				v1 = i32(0)
				v3 = i32(1)
				goto l16
			}
		l10:
			v3 = v8 + v9
			v1 = v1 - v9
		l16:
			store32(m.memory[int64(uint32(v0))+8:], uint32(v1))
			store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
			store32(m.memory[uint32(v0):], uint32(i32(-1)))
			store32(m.memory[int64(uint32(v0))+16:], uint32(v4-v5))
			store32(m.memory[int64(uint32(v0))+12:], uint32(v8+v5))
			goto l3
		}
	}
l3:
	m.g0 = v2 + i32(32)
}
func (m *Module) fn919(v0 int32) {
	var v1, v2, v3 int32
	t0 := m.g0
	v1 = t0 - i32(16)
	m.g0 = v1
	t1 := int32(load32(m.memory[uint32(v0):]))
	t2 := v1 + i32(4)
	v2 = t1
	t3 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t4 := v2
	v2 = v2 << 1
	p5 := i32(4)
	if uint32(v2) > uint32(i32(4)) {
		p5 = v2
	}
	v2 = p5
	m.fn912(t2, t4, t3, v2, i32(4), i32(4))
	{
		t6 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		if t6 != i32(1) {
			goto l0
		}
		t7 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		t8 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		m.fn12(t7, t8)
		panic("unreachable")
	}
l0:
	t9 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	v3 = t9
	store32(m.memory[uint32(v0):], uint32(v2))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
	m.g0 = v1 + i32(16)
}
func (m *Module) fn920(v0, v1, v2, v3, v4, v5 int32) {
	var v6, v7, v8, v9, v10, v11, v12, v13 int32
	{
		t0 := int32(m.memory[int64(uint32(v1))+37])
		if t0 == 0 {
			goto l0
		}
		{
			t1 := int32(load32(m.memory[int64(uint32(v1))+16:]))
			v6 = t1
			if uint32(v6) > uint32(i32(31)) {
				m.fn922(v0, v1, v2, v3, v4, v5)
				return
			}
			t2 := int32(load32(m.memory[int64(uint32(v1))+12:]))
			v7 = t2
			{
				if v6 == 0 {
					goto l2
				}
				if uint32(v5) < uint32(v4) {
					goto l3
				}
				if uint32(v5) > uint32(v3) {
					goto l3
				}
				v8 = v7 + v6<<3
				v9 = v2 + v4
				v10 = v5 - v4
				v11 = v7
			l7:
				{
					t3 := int32(load32(m.memory[uint32(v11+i32(4)):]))
					v12 = t3
					t4 := int32(load32(m.memory[uint32(v11):]))
					t5 := v12
					v13 = t4
					if uint32(t5) < uint32(v13) {
						goto l4
					}
					if uint32(v12) > uint32(v3) {
						goto l4
					}
					{
						if v12-v13 != v10 {
							goto l5
						}
						t6 := m.fn973(v2+v13, v9, v10)
						if t6 == 0 {
							store32(m.memory[int64(uint32(v0))+8:], uint32(v13))
							store32(m.memory[int64(uint32(v0))+4:], uint32(v4))
							m.memory[uint32(v0)] = byte(i32(4))
							return
						}
					}
				l5:
					v11 = v11 + i32(8)
					if v11 != v8 {
						goto l7
					}
				}
			l2:
				{
					t7 := int32(load32(m.memory[int64(uint32(v1))+8:]))
					if v6 != t7 {
						goto l8
					}
					m.fn921(v1 + i32(8))
					t8 := int32(load32(m.memory[int64(uint32(v1))+12:]))
					v7 = t8
				}
			l8:
				store32(m.memory[int64(uint32(v1))+16:], uint32(v6+i32(1)))
				v11 = v7 + v6<<3
				store32(m.memory[int64(uint32(v11))+4:], uint32(v5))
				store32(m.memory[uint32(v11):], uint32(v4))
				goto l0
			l3:
				t9 := int32(load32(m.memory[int64(uint32(v7))+4:]))
				v12 = t9
				t10 := int32(load32(m.memory[uint32(v7):]))
				t11 := v12
				v13 = t10
				if uint32(t11) < uint32(v13) {
					goto l4
				}
				if uint32(v12) <= uint32(v3) {
					m.fn120(v4, v5, v3, i32(1272056))
					panic("unreachable")
				}
			}
		l4:
			m.fn120(v13, v12, v3, i32(1272072))
			panic("unreachable")
		}
	}
l0:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v5))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v4))
	m.memory[uint32(v0)] = byte(i32(255))
}
func (m *Module) fn921(v0 int32) {
	var v1, v2, v3 int32
	t0 := m.g0
	v1 = t0 - i32(16)
	m.g0 = v1
	t1 := int32(load32(m.memory[uint32(v0):]))
	t2 := v1 + i32(4)
	v2 = t1
	t3 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t4 := v2
	v2 = v2 << 1
	p5 := i32(4)
	if uint32(v2) > uint32(i32(4)) {
		p5 = v2
	}
	v2 = p5
	m.fn912(t2, t4, t3, v2, i32(4), i32(8))
	{
		t6 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		if t6 != i32(1) {
			goto l0
		}
		t7 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		t8 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		m.fn12(t7, t8)
		panic("unreachable")
	}
l0:
	t9 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	v3 = t9
	store32(m.memory[uint32(v0):], uint32(v2))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
	m.g0 = v1 + i32(16)
}
func (m *Module) fn922(v0, v1, v2, v3, v4, v5 int32) {
	var v6, v7, v8, v9, v10, v11, v12, v13 int32
	var v14, v15 int64
	t0 := m.g0
	v6 = t0 - i32(32)
	m.g0 = v6
	v7 = v1 + i32(20)
	{
		{
			t1 := int32(load32(m.memory[int64(uint32(v1))+20:]))
			if t1 != 0 {
				goto l0
			}
			{
				{
					t2 := int32(load32(m.memory[int64(uint32(v1))+16:]))
					v8 = t2
					if v8 != 0 {
						goto l1
					}
					v9 = i32(0x137888)
					v10 = i32(0)
					v11 = i32(0)
					goto l2
				}
			l1:
				{
					{
						if uint32(v8) > uint32(i32(7)) {
							goto l3
						}
						p3 := v8<<1&i32(8) + i32(8)
						if v8 == i32(1) {
							p3 = i32(4)
						}
						v10 = p3
						goto l4
					}
				l3:
					t4 := int32(uint32(v8<<4) / uint32(i32(7)))
					v10 = i32_shr_u(i32(-1), int32(bits.LeadingZeros32(uint32(t4+i32(-1)))))
					if uint32(v10) > uint32(i32(0x1ffffffe)) {
						goto l5
					}
					v10 = v10 + i32(1)
				}
			l4:
				v11 = v10 + i32(8)
				t5 := v11
				v12 = v10 << 3
				v9 = t5 + v12
				if uint32(v9) < uint32(v11) {
					goto l5
				}
				if uint32(v9) > uint32(i32(0x7ffffff8)) {
					goto l5
				}
				t6 := m.fn7(v9)
				v13 = t6
				if v13 == 0 {
					m.fn23(i32(8), v9)
					panic("unreachable")
				}
				v9 = v13 + v12
				if v11 == 0 {
					goto l7
				}
				memory_fill(m.memory, uint32(v9), i32(255), uint32(v11))
			l7:
				v11 = v10 + i32(-1)
				p7 := int32(uint32(v10)>>3) * i32(7)
				if uint32(v10) < uint32(i32(9)) {
					p7 = v11
				}
				v10 = p7
			}
		l2:
			store32(m.memory[int64(uint32(v6))+28:], uint32(i32(0)))
			store32(m.memory[int64(uint32(v6))+24:], uint32(v10))
			store32(m.memory[int64(uint32(v6))+20:], uint32(v11))
			store32(m.memory[int64(uint32(v6))+16:], uint32(v9))
			{
				if v8 == 0 {
					goto l8
				}
				t8 := int32(load32(m.memory[int64(uint32(v1))+12:]))
				v10 = t8
				v9 = v10 + v8<<3
			l10:
				{
					t9 := int32(load32(m.memory[uint32(v10+i32(4)):]))
					v8 = t9
					t10 := int32(load32(m.memory[uint32(v10):]))
					t11 := v8
					v11 = t10
					if uint32(t11) < uint32(v11) {
						goto l9
					}
					if uint32(v8) > uint32(v3) {
						goto l9
					}
					t12 := m.fn923(v2+v11, v8-v11)
					_ = m.fn924(v6+i32(16), t12)
					v10 = v10 + i32(8)
					if v10 != v9 {
						goto l10
					}
				}
			}
		l8:
			t14 := int64(load64(m.memory[int64(uint32(v6))+24:]))
			t15 := v6
			v14 = t14
			store64(m.memory[int64(uint32(t15))+8:], uint64(v14))
			t16 := int64(load64(m.memory[int64(uint32(v6))+16:]))
			t17 := v6
			v15 = t16
			store64(m.memory[uint32(t17):], uint64(v15))
			store64(m.memory[int64(uint32(v7))+8:], uint64(v14))
			store64(m.memory[uint32(v7):], uint64(v15))
		}
	l0:
		{
			if uint32(v5) < uint32(v4) {
				goto l11
			}
			if uint32(v5) > uint32(v3) {
				goto l11
			}
			t18 := v7
			v12 = v2 + v4
			t19 := v12
			v9 = v5 - v4
			t20 := m.fn923(t19, v9)
			t21 := m.fn924(t18, t20)
			if t21 != 0 {
				t23 := int32(load32(m.memory[int64(uint32(v1))+16:]))
				v13 = t23
				if v13 != 0 {
					t24 := int32(load32(m.memory[int64(uint32(v1))+12:]))
					v10 = t24
					v7 = v10 + v13<<3
				l18:
					{
						t25 := int32(load32(m.memory[uint32(v10+i32(4)):]))
						v8 = t25
						t26 := int32(load32(m.memory[uint32(v10):]))
						t27 := v8
						v11 = t26
						if uint32(t27) < uint32(v11) {
							goto l15
						}
						if uint32(v8) > uint32(v3) {
							goto l15
						}
						{
							if v8-v11 != v9 {
								goto l16
							}
							t28 := m.fn973(v2+v11, v12, v9)
							if t28 == 0 {
								goto l17
							}
						}
					l16:
						v10 = v10 + i32(8)
						if v10 == v7 {
							goto l13
						}
						goto l18
					l17:
					}
					store32(m.memory[int64(uint32(v0))+8:], uint32(v11))
					store32(m.memory[int64(uint32(v0))+4:], uint32(v4))
					m.memory[uint32(v0)] = byte(i32(4))
					goto l19
				}
				v13 = i32(0)
				goto l13
			}
			t22 := int32(load32(m.memory[int64(uint32(v1))+16:]))
			v13 = t22
			goto l13
		}
	l11:
		m.fn120(v4, v5, v3, i32(1272560))
		panic("unreachable")
	l5:
		m.fn27(i32(1271632), i32(57), i32(1271660))
		panic("unreachable")
	l9:
		m.fn120(v11, v8, v3, i32(1272088))
		panic("unreachable")
	l15:
		m.fn120(v11, v8, v3, i32(1272104))
		panic("unreachable")
	l13:
		{
			t29 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			if v13 != t29 {
				goto l20
			}
			m.fn921(v1 + i32(8))
		}
	l20:
		store32(m.memory[int64(uint32(v0))+8:], uint32(v5))
		store32(m.memory[int64(uint32(v0))+4:], uint32(v4))
		m.memory[uint32(v0)] = byte(i32(255))
		store32(m.memory[int64(uint32(v1))+16:], uint32(v13+i32(1)))
		t30 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		v10 = t30 + v13<<3
		store32(m.memory[int64(uint32(v10))+4:], uint32(v5))
		store32(m.memory[uint32(v10):], uint32(v4))
	}
l19:
	m.g0 = v6 + i32(32)
}
func (m *Module) fn923(v0, v1 int32) int64 {
	var v2, v3 int32
	var v4 int64
	var v5 int32
	var v6, v7, v8 int64
	var v9 int32
	var v10, v11 int64
	v2 = v1 & i32(7)
	v3 = i32(0)
	v4 = i64(8317987319222330741)
	v5 = v1 & i32(0x7ffffff8)
	if v5 != 0 {
		goto l0
	}
	v6 = i64(0x6c7967656e657261)
	v7 = i64(7237128888997146477)
	v8 = i64(8387220255154660723)
	v9 = i32(0)
	goto l1
l0:
	v9 = i32(0)
	v8 = i64(8387220255154660723)
	v7 = i64(7237128888997146477)
	v6 = i64(0x6c7967656e657261)
l2:
	{
		t0 := int64(load64(m.memory[uint32(v0+v9):]))
		v10 = t0
		v8 = v10 ^ v8
		v6 = v8 + v6
		t1 := v6
		v4 = v4 + v7
		v7 = v4 ^ i64_rotl(v7, i64(13))
		v11 = t1 + v7
		v7 = v11 ^ i64_rotl(v7, i64(17))
		v6 = v6 ^ i64_rotl(v8, i64(16))
		t2 := i64_rotl(v6, i64(21))
		v4 = v6 + i64_rotl(v4, i64(32))
		v8 = t2 ^ v4
		v6 = i64_rotl(v11, i64(32))
		v4 = v4 ^ v10
		v9 = v9 + i32(8)
		if uint32(v9) < uint32(v5) {
			goto l2
		}
	}
	v9 = (v5+i32(-1))&i32(-8) + i32(8)
l1:
	v10 = i64(0)
	{
		if uint32(v2) < uint32(i32(4)) {
			goto l3
		}
		t3 := int64(load32(m.memory[uint32(v0+v9):]))
		v10 = t3
		v3 = i32(4)
	}
l3:
	{
		if uint32(v3|i32(1)) >= uint32(v2) {
			goto l4
		}
		t4 := int64(load16(m.memory[uint32(v0+v9+v3):]))
		v10 = i64_shl(t4, int64(uint32(v3<<3))) | v10
		v3 = v3 | i32(2)
	}
l4:
	{
		if uint32(v3) >= uint32(v2) {
			goto l5
		}
		t5 := int64(m.memory[uint32(v0+(v3+v9))])
		v10 = i64_shl(t5, int64(uint32(v3<<3))) | v10
	}
l5:
	v10 = v10 | int64(uint32(v1))<<56
	v8 = v10 ^ v8
	t6 := i64_rotl(v8, i64(16))
	v6 = v8 + v6
	v8 = t6 ^ v6
	t7 := i64_rotl(v8, i64(21))
	t8 := v8
	v4 = v4 + v7
	v8 = t8 + i64_rotl(v4, i64(32))
	v11 = t7 ^ v8
	t9 := i64_rotl(v11, i64(16))
	t10 := v11
	t11 := v6
	v7 = v4 ^ i64_rotl(v7, i64(13))
	v4 = t11 + v7
	v6 = t10 + (i64_rotl(v4, i64(32)) ^ i64(255))
	v11 = t9 ^ v6
	t12 := i64_rotl(v11, i64(21))
	t13 := v11
	t14 := v8 ^ v10
	v7 = v4 ^ i64_rotl(v7, i64(17))
	v4 = t14 + v7
	v8 = t13 + i64_rotl(v4, i64(32))
	v10 = t12 ^ v8
	t15 := i64_rotl(v10, i64(16))
	t16 := v10
	v7 = v4 ^ i64_rotl(v7, i64(13))
	v4 = v7 + v6
	v6 = t16 + i64_rotl(v4, i64(32))
	v10 = t15 ^ v6
	t17 := i64_rotl(v10, i64(21))
	t18 := v10
	v7 = v4 ^ i64_rotl(v7, i64(17))
	v4 = v7 + v8
	v8 = t18 + i64_rotl(v4, i64(32))
	v10 = t17 ^ v8
	t19 := i64_rotl(v10, i64(16))
	t20 := v10
	v7 = i64_rotl(v7, i64(13)) ^ v4
	v4 = v7 + v6
	v6 = t20 + i64_rotl(v4, i64(32))
	t21 := i64_rotl(t19^v6, i64(21))
	v7 = i64_rotl(v7, i64(17)) ^ v4
	v7 = i64_rotl(v7, i64(13)) ^ (v7 + v8)
	t22 := t21 ^ i64_rotl(v7, i64(17))
	v7 = v7 + v6
	return t22 ^ i64_rotl(v7, i64(32)) ^ v7
}
func (m *Module) fn924(v0 int32, v1 int64) int32 {
	var v2, v3 int32
	var v4, v5 int64
	var v6, v7, v8 int32
	var v9, v10 int64
	var v11, v12 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		if t0 != 0 {
			goto l0
		}
		_ = m.fn911(v0)
	}
l0:
	t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	v2 = t2
	v3 = v2 & int32(v1)
	v4 = int64(uint64(v1) >> 25)
	v5 = v4 & i64(127) * i64(72340172838076673)
	t3 := int32(load32(m.memory[uint32(v0):]))
	v6 = t3
	v7 = i32(0)
	v8 = i32(0)
	var _ int32
l9:
	{
		{
			t5 := int64(load64(m.memory[uint32(v6+v3):]))
			v9 = t5
			v10 = v9 ^ v5
			v10 = (v10 ^ i64(-1)) & (v10 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
			if v10 == 0 {
				goto l1
			}
		l3:
			{
				v11 = i32(1)
				t6 := int64(load64(m.memory[uint32(v6-(int32(uint32(int64(bits.TrailingZeros64(uint64(v10))))>>3)+v3)&v2<<3+i32(-8)):]))
				if v1 == t6 {
					goto l2
				}
				v10 = (v10 + i64(-1)) & v10
				if !(v10 == 0) {
					goto l3
				}
			}
		}
	l1:
		v10 = v9 & i64(-0x7f7f7f7f7f7f7f80)
		if v7 == i32(1) {
			goto l4
		}
		if v10 == 0 {
			goto l5
		}
		v12 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v10))))>>3) + v3) & v2
	l4:
		if v10&(v9<<1) != i64(0) {
			goto l6
		}
		v7 = i32(1)
		goto l7
	l6:
		v11 = i32(0)
		{
			t7 := int32(int8(m.memory[uint32(v6+v12)]))
			v3 = t7
			if v3 < i32(0) {
				goto l8
			}
			t8 := int64(load64(m.memory[uint32(v6):]))
			t9 := v6
			v12 = int32(uint32(int64(bits.TrailingZeros64(uint64(t8&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
			t10 := int32(m.memory[uint32(t9+v12)])
			v3 = t10
		}
	l8:
		t11 := v6 + v12
		v7 = int32(v4) & i32(127)
		m.memory[uint32(t11)] = byte(v7)
		m.memory[uint32(v6+(v12+i32(-8))&v2+i32(8))] = byte(v7)
		t12 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		store32(m.memory[int64(uint32(v0))+8:], uint32(t12-v3&i32(1)))
		t13 := int32(load32(m.memory[int64(uint32(v0))+12:]))
		store32(m.memory[int64(uint32(v0))+12:], uint32(t13+i32(1)))
		store64(m.memory[uint32(v6-v12<<3+i32(-8)):], uint64(v1))
	}
l2:
	return v11
l5:
	v7 = i32(0)
l7:
	v8 = v8 + i32(8)
	v3 = (v8 + v3) & v2
	goto l9
}
func (m *Module) fn925(v0 int32) {
	var v1, v2, v3 int32
	t0 := m.g0
	v1 = t0 - i32(16)
	m.g0 = v1
	t1 := int32(load32(m.memory[uint32(v0):]))
	t2 := v1 + i32(4)
	v2 = t1
	t3 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t4 := v2
	v2 = v2 << 1
	p5 := i32(4)
	if uint32(v2) > uint32(i32(4)) {
		p5 = v2
	}
	v2 = p5
	m.fn912(t2, t4, t3, v2, i32(4), i32(16))
	{
		t6 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		if t6 != i32(1) {
			goto l0
		}
		t7 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		t8 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		m.fn12(t7, t8)
		panic("unreachable")
	}
l0:
	t9 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	v3 = t9
	store32(m.memory[uint32(v0):], uint32(v2))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
	m.g0 = v1 + i32(16)
}
func (m *Module) fn926(v0, v1, v2 int32) {
	var v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14 int32
	t0 := m.g0
	v3 = t0 - i32(16)
	m.g0 = v3
	if v2 == 0 {
		goto l0
	}
	v4 = v1 + v2
	if uint32(v2) < uint32(i32(4)) {
		v5 = v1
		t6 := int32(m.memory[uint32(v1)])
		if t6 == i32(13) {
			goto l3
		}
		if v2 == i32(1) {
			goto l0
		}
		{
			t7 := int32(m.memory[int64(uint32(v1))+1])
			if t7 != i32(13) {
				if v2 == i32(2) {
					goto l0
				}
				t8 := int32(m.memory[int64(uint32(v1))+2])
				if t8 != i32(13) {
					goto l0
				}
				v5 = v1 + i32(2)
				goto l3
			}
			v5 = v1 + i32(1)
			goto l3
		}
	}
	{
		t1 := int32(load32(m.memory[uint32(v1):]))
		v5 = t1
		if (i32(16843008)-(v5^i32(218959117))|v5)&i32(-2139062144) == i32(-2139062144) {
			t3 := v1
			v7 = v1 & i32(3)
			v6 = i32(4) - v7
			v5 = t3 + v6
			if uint32(v2) < uint32(i32(9)) {
				if uint32(v6) >= uint32(v2) {
					goto l0
				}
				v6 = v2 + v7 + i32(-4)
			l9:
				{
					t9 := int32(m.memory[uint32(v5)])
					if t9 == i32(13) {
						goto l3
					}
					v5 = v5 + i32(1)
					v6 = v6 + i32(-1)
					if v6 == 0 {
						goto l0
					}
					goto l9
				}
			}
			if uint32(v6) > uint32(v2+i32(-8)) {
				goto l6
			}
			v7 = v4 + i32(-8)
		l7:
			{
				t4 := int32(load32(m.memory[uint32(v5):]))
				v6 = t4
				if (i32(16843008)-(v6^i32(218959117))|v6)&i32(-2139062144) != i32(-2139062144) {
					goto l6
				}
				t5 := int32(load32(m.memory[uint32(v5+i32(4)):]))
				v6 = t5
				if (i32(16843008)-(v6^i32(218959117))|v6)&i32(-2139062144) != i32(-2139062144) {
					goto l6
				}
				v5 = v5 + i32(8)
				if uint32(v5) <= uint32(v7) {
					goto l7
				}
				goto l6
			}
		}
		v6 = v2
		v5 = v1
	l4:
		{
			t2 := int32(m.memory[uint32(v5)])
			if t2 == i32(13) {
				goto l3
			}
			v5 = v5 + i32(1)
			v6 = v6 + i32(-1)
			if v6 != 0 {
				goto l4
			}
			goto l0
		}
	}
l6:
	if uint32(v5) >= uint32(v4) {
		goto l0
	}
l10:
	{
		t10 := int32(m.memory[uint32(v5)])
		if t10 == i32(13) {
			goto l3
		}
		v5 = v5 + i32(1)
		if v5 == v4 {
			goto l0
		}
		goto l10
	}
l3:
	{
		t11 := m.fn7(v2)
		v7 = t11
		if v7 == 0 {
			m.fn12(i32(1), v2)
			panic("unreachable")
		}
		v6 = v5 - v1
		store32(m.memory[int64(uint32(v3))+8:], uint32(v7))
		store32(m.memory[int64(uint32(v3))+4:], uint32(v2))
		{
			{
				if v5 == v1 {
					goto l12
				}
				t12 := int32(int8(m.memory[uint32(v1+v6)]))
				if t12 < i32(-64) {
					m.fn37(v1, v2, i32(0), v6, i32(1272728))
					panic("unreachable")
				}
				if v6 == 0 {
					goto l12
				}
				memory_copy(m.memory, uint32(v7), uint32(v1), uint32(v6))
			}
		l12:
			store32(m.memory[int64(uint32(v3))+12:], uint32(v6))
			{
				{
					t13 := m.fn913(v3+i32(4), v1, v2, v6, i32(10))
					v7 = t13
					if uint32(v7) > uint32(v2) {
						goto l14
					}
					v8 = v2 + i32(-4)
					v9 = v2 + i32(-8)
					v10 = v4 + i32(-8)
				l35:
					{
						v11 = v1 + v7
						v12 = v2 - v7
						{
							if uint32(v2) <= uint32(v7) {
								goto l15
							}
							v6 = v12
							v5 = v11
							if uint32(v12) <= uint32(i32(3)) {
							l25:
								{
									t19 := int32(m.memory[uint32(v5)])
									if t19 == i32(13) {
										t21 := int32(m.memory[uint32(v11)])
										v13 = t21
										goto l21
									}
									v5 = v5 + i32(1)
									v6 = v6 + i32(-1)
									if v6 != 0 {
										goto l25
									}
								}
								t20 := int32(m.memory[uint32(v11)])
								v13 = t20
								goto l23
							}
							v6 = v2
							v5 = v11
							{
								t14 := int32(load32(m.memory[uint32(v11):]))
								v13 = t14
								if (i32(16843008)-(v13^i32(218959117))|v13)&i32(-2139062144) != i32(-2139062144) {
								l22:
									{
										t17 := int32(m.memory[uint32(v5)])
										if t17 == i32(13) {
											goto l21
										}
										v5 = v5 + i32(1)
										t18 := v7
										v6 = v6 + i32(-1)
										if t18 != v6 {
											goto l22
										}
										goto l23
									}
								}
								v14 = v11 & i32(3)
								v5 = i32(4) - v14
								v6 = v5 + v7
								v5 = v11 + v5
								if uint32(v12) < uint32(i32(9)) {
									if uint32(v6) >= uint32(v2) {
										goto l23
									}
									v6 = v8 + v14
								l26:
									{
										t22 := int32(m.memory[uint32(v5)])
										if t22 == i32(13) {
											goto l21
										}
										v5 = v5 + i32(1)
										t23 := v7
										v6 = v6 + i32(-1)
										if t23 != v6 {
											goto l26
										}
										goto l23
									}
								}
								if v6 > v9 {
									goto l19
								}
							l20:
								{
									t15 := int32(load32(m.memory[uint32(v5):]))
									v6 = t15
									if (i32(16843008)-(v6^i32(218959117))|v6)&i32(-2139062144) != i32(-2139062144) {
										goto l19
									}
									t16 := int32(load32(m.memory[uint32(v5+i32(4)):]))
									v6 = t16
									if (i32(16843008)-(v6^i32(218959117))|v6)&i32(-2139062144) != i32(-2139062144) {
										goto l19
									}
									v5 = v5 + i32(8)
									if uint32(v5) <= uint32(v10) {
										goto l20
									}
									goto l19
								}
							}
						l19:
							if uint32(v5) >= uint32(v4) {
								goto l23
							}
						l27:
							{
								t24 := int32(m.memory[uint32(v5)])
								if t24 == i32(13) {
									goto l21
								}
								v5 = v5 + i32(1)
								if v5 != v4 {
									goto l27
								}
							}
						l23:
							if int32(int8(v13)) <= i32(-65) {
								goto l28
							}
						l15:
							t25 := int32(load32(m.memory[int64(uint32(v3))+4:]))
							t26 := int32(load32(m.memory[int64(uint32(v3))+12:]))
							t27 := v12
							v5 = t26
							if uint32(t27) <= uint32(t25-v5) {
								goto l29
							}
							m.fn244(v3+i32(4), v5, v12)
							t28 := int32(load32(m.memory[int64(uint32(v3))+12:]))
							v5 = t28
							goto l30
						}
					l21:
						v6 = v5 - v11
						v12 = v6 + v7
						if int32(int8(v13)) < i32(-64) {
							goto l31
						}
						t29 := int32(int8(m.memory[uint32(v1+v12)]))
						if t29 <= i32(-65) {
							goto l31
						}
						{
							{
								t30 := int32(load32(m.memory[int64(uint32(v3))+4:]))
								t31 := int32(load32(m.memory[int64(uint32(v3))+12:]))
								t32 := v6
								v7 = t31
								if uint32(t32) <= uint32(t30-v7) {
									goto l32
								}
								m.fn244(v3+i32(4), v7, v6)
								t33 := int32(load32(m.memory[int64(uint32(v3))+12:]))
								v7 = t33
								goto l33
							}
						l32:
							if v5 == v11 {
								goto l34
							}
						l33:
							if v6 == 0 {
								goto l34
							}
							t34 := int32(load32(m.memory[int64(uint32(v3))+8:]))
							memory_copy(m.memory, uint32(t34+v7), uint32(v11), uint32(v6))
						}
					l34:
						store32(m.memory[int64(uint32(v3))+12:], uint32(v7+v6))
						t35 := m.fn913(v3+i32(4), v1, v2, v12, i32(10))
						v7 = t35
						if uint32(v7) <= uint32(v2) {
							goto l35
						}
					}
				}
			l14:
				m.fn120(v7, v2, v2, i32(1272760))
				panic("unreachable")
			l31:
				m.fn37(v1, v2, v7, v12, i32(1272744))
				panic("unreachable")
			l29:
				if v2 == v7 {
					goto l36
				}
			l30:
				if v12 == 0 {
					goto l36
				}
				t36 := int32(load32(m.memory[int64(uint32(v3))+8:]))
				memory_copy(m.memory, uint32(t36+v5), uint32(v11), uint32(v12))
			}
		l36:
			store32(m.memory[int64(uint32(v3))+12:], uint32(v5+v12))
		l28:
			t37 := int32(load32(m.memory[int64(uint32(v3))+12:]))
			store32(m.memory[int64(uint32(v0))+8:], uint32(t37))
			t38 := int64(load64(m.memory[int64(uint32(v3))+4:]))
			store64(m.memory[uint32(v0):], uint64(t38))
			goto l37
		}
	}
l0:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v2))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
	store32(m.memory[uint32(v0):], uint32(i32(-1)))
l37:
	m.g0 = v3 + i32(16)
}
func (m *Module) fn927(v0, v1 int32) int32 {
	var v2, v3 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	t1 := int32(load32(m.memory[uint32(v0):]))
	t2 := int32(load32(m.memory[uint32(t1):]))
	v3 = t2
	v0 = i32(9)
l0:
	{
		t3 := int32(m.memory[int64(uint32(v3&i32(15)))+1099240])
		m.memory[uint32(v2+i32(8)+v0+i32(-2))] = byte(t3)
		v0 = v0 + i32(-1)
		v3 = int32(uint32(v3) >> 4)
		if v3 != 0 {
			goto l0
		}
	}
	t4 := m.fn679(v1, i32(1), i32(1122974), i32(2), v2+i32(8)+v0+i32(-1), i32(9)-v0)
	v0 = t4
	m.g0 = v2 + i32(16)
	return v0
}
func (m *Module) fn928(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := v1
	v0 = t0
	t2 := int32(load32(m.memory[uint32(v0+i32(4)):]))
	t3 := int32(load32(m.memory[uint32(v0+i32(8)):]))
	t4 := m.fn5(t1, t2, t3)
	return t4
}
func (m *Module) fn929(v0, v1 int32) int32 {
	var v2, v3, v4 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	t1 := int32(load32(m.memory[uint32(v0):]))
	v3 = t1
	{
		{
			{
				t2 := int32(load32(m.memory[int64(uint32(v1))+8:]))
				v0 = t2
				if v0&i32(0x2000000) != 0 {
					t4 := int32(load32(m.memory[uint32(v3):]))
					v4 = t4
					v0 = i32(9)
				l4:
					{
						t5 := int32(m.memory[int64(uint32(v4&i32(15)))+1099240])
						m.memory[uint32(v2+i32(8)+v0+i32(-2))] = byte(t5)
						v0 = v0 + i32(-1)
						v4 = int32(uint32(v4) >> 4)
						if v4 != 0 {
							goto l4
						}
					}
					v4 = i32(1)
					t6 := m.fn679(v1, i32(1), i32(1122974), i32(2), v2+i32(8)+v0+i32(-1), i32(9)-v0)
					if t6 == 0 {
						goto l2
					}
					goto l3
				}
				if v0&i32(0x4000000) != 0 {
					goto l1
				}
				t3 := m.fn43(v3, v1)
				if t3 == 0 {
					goto l2
				}
				v4 = i32(1)
				goto l3
			}
		l1:
			t7 := int32(load32(m.memory[uint32(v3):]))
			v4 = t7
			v0 = i32(9)
		l5:
			{
				t8 := int32(m.memory[int64(uint32(v4&i32(15)))+1122976])
				m.memory[uint32(v2+i32(8)+v0+i32(-2))] = byte(t8)
				v0 = v0 + i32(-1)
				v4 = int32(uint32(v4) >> 4)
				if v4 != 0 {
					goto l5
				}
			}
			v4 = i32(1)
			t9 := m.fn679(v1, i32(1), i32(1122974), i32(2), v2+i32(8)+v0+i32(-1), i32(9)-v0)
			if t9 != 0 {
				goto l3
			}
		}
	l2:
		{
			t10 := int32(load32(m.memory[uint32(v1):]))
			t11 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t12 := int32(load32(m.memory[int64(uint32(t11))+12:]))
			t13 := m.t0[uint(t12)].(func(int32, int32, int32) int32)(t10, i32(1274031), i32(2))
			if t13 == 0 {
				goto l6
			}
			v4 = i32(1)
			goto l3
		}
	l6:
		v0 = v3 + i32(4)
		{
			t14 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			v4 = t14
			if v4&i32(0x2000000) != 0 {
				t16 := int32(load32(m.memory[uint32(v0):]))
				v4 = t16
				v0 = i32(9)
			l9:
				{
					t17 := int32(m.memory[int64(uint32(v4&i32(15)))+1099240])
					m.memory[uint32(v2+i32(8)+v0+i32(-2))] = byte(t17)
					v0 = v0 + i32(-1)
					v4 = int32(uint32(v4) >> 4)
					if v4 != 0 {
						goto l9
					}
				}
				t18 := m.fn679(v1, i32(1), i32(1122974), i32(2), v2+i32(8)+v0+i32(-1), i32(9)-v0)
				v4 = t18
				goto l3
			}
			if v4&i32(0x4000000) != 0 {
				goto l8
			}
			t15 := m.fn43(v0, v1)
			v4 = t15
			goto l3
		}
	l8:
		t19 := int32(load32(m.memory[uint32(v0):]))
		v4 = t19
		v0 = i32(9)
	l10:
		{
			t20 := int32(m.memory[int64(uint32(v4&i32(15)))+1122976])
			m.memory[uint32(v2+i32(8)+v0+i32(-2))] = byte(t20)
			v0 = v0 + i32(-1)
			v4 = int32(uint32(v4) >> 4)
			if v4 != 0 {
				goto l10
			}
		}
		t21 := m.fn679(v1, i32(1), i32(1122974), i32(2), v2+i32(8)+v0+i32(-1), i32(9)-v0)
		v4 = t21
	}
l3:
	m.g0 = v2 + i32(16)
	return v4
}
func (m *Module) fn930(v0, v1 int32) int32 {
	var v2 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	{
		t1 := int32(load32(m.memory[uint32(v0):]))
		v0 = t1
		t2 := int32(m.memory[uint32(v0)])
		switch t2 {
		default:
			t3 := int32(load32(m.memory[uint32(v1):]))
			t4 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t5 := int32(load32(m.memory[int64(uint32(t4))+12:]))
			t6 := m.t0[uint(t5)].(func(int32, int32, int32) int32)(t3, i32(1273580), i32(22))
			v1 = t6
			goto l4
		case 1:
			t7 := int32(m.memory[int64(uint32(v0))+1])
			t8 := v1
			v0 = t7 << 2
			t9 := int32(load32(m.memory[int64(uint32(v0))+1291152:]))
			t10 := int32(load32(m.memory[int64(uint32(v0))+1291128:]))
			t11 := m.fn5(t8, t9, t10)
			v1 = t11
			goto l4
		case 2:
			store32(m.memory[int64(uint32(v2))+4:], uint32(v0+i32(4)))
			store64(m.memory[int64(uint32(v2))+8:], uint64(int64(uint32(i32(34)))<<32|int64(uint32(v2+i32(4)))))
			t12 := int32(load32(m.memory[uint32(v1):]))
			t13 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t14 := m.fn45(t12, t13, i32(1052731), v2+i32(8))
			v1 = t14
			goto l4
		case 3:
			store32(m.memory[int64(uint32(v2))+4:], uint32(v0+i32(4)))
			store64(m.memory[int64(uint32(v2))+8:], uint64(int64(uint32(i32(91)))<<32|int64(uint32(v2+i32(4)))))
			t15 := int32(load32(m.memory[uint32(v1):]))
			t16 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t17 := m.fn45(t15, t16, i32(1066418), v2+i32(8))
			v1 = t17
		}
	}
l4:
	m.g0 = v2 + i32(16)
	return v1
}
func (m *Module) fn931(v0, v1, v2 int32) int32 {
	var v3, v4 int32
	var v5 int64
	t0 := m.g0
	v3 = t0 - i32(32)
	m.g0 = v3
	v4 = v0 + i32(4)
	{
		t1 := int32(m.memory[uint32(v0)])
		switch t1 {
		default:
			store32(m.memory[int64(uint32(v3))+12:], uint32(v4))
			store64(m.memory[int64(uint32(v3))+16:], uint64(int64(uint32(i32(34)))<<32|int64(uint32(v3+i32(12)))))
			t2 := m.fn45(v1, v2, i32(1065837), v3+i32(16))
			v0 = t2
			goto l5
		case 1:
			store32(m.memory[int64(uint32(v3))+12:], uint32(v4))
			store64(m.memory[int64(uint32(v3))+16:], uint64(int64(uint32(i32(34)))<<32|int64(uint32(v3+i32(12)))))
			t3 := m.fn45(v1, v2, i32(1065521), v3+i32(16))
			v0 = t3
			goto l5
		case 2:
			store32(m.memory[int64(uint32(v3))+12:], uint32(v4))
			store64(m.memory[int64(uint32(v3))+16:], uint64(int64(uint32(i32(34)))<<32|int64(uint32(v3+i32(12)))))
			t4 := m.fn45(v1, v2, i32(1066137), v3+i32(16))
			v0 = t4
			goto l5
		case 3:
			store32(m.memory[int64(uint32(v3))+8:], uint32(v4))
			t5 := int32(m.memory[int64(uint32(v0))+1])
			store32(m.memory[int64(uint32(v3))+12:], uint32(t5))
			store64(m.memory[int64(uint32(v3))+24:], uint64(int64(uint32(i32(44)))<<32|int64(uint32(v3+i32(12)))))
			store64(m.memory[int64(uint32(v3))+16:], uint64(int64(uint32(i32(34)))<<32|int64(uint32(v3+i32(8)))))
			t6 := m.fn45(v1, v2, i32(1065461), v3+i32(16))
			v0 = t6
			goto l5
		case 4:
			store32(m.memory[int64(uint32(v3))+8:], uint32(v4))
			store32(m.memory[int64(uint32(v3))+12:], uint32(v0+i32(8)))
			t7 := v3
			v5 = int64(uint32(i32(34))) << 32
			store64(m.memory[int64(uint32(t7))+24:], uint64(v5|int64(uint32(v3+i32(12)))))
			store64(m.memory[int64(uint32(v3))+16:], uint64(v5|int64(uint32(v3+i32(8)))))
			t8 := m.fn45(v1, v2, i32(1050308), v3+i32(16))
			v0 = t8
		}
	}
l5:
	m.g0 = v3 + i32(32)
	return v0
}
func (m *Module) fn932(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(t0):]))
	t2 := m.fn254(t1+i32(8), v1)
	return t2
}
func (m *Module) fn933(v0, v1 int32) int32 {
	var v2 int32
	t0 := int32(load32(m.memory[uint32(v1):]))
	v2 = t0
	t1 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	t2 := int32(load32(m.memory[int64(uint32(t1))+12:]))
	v1 = t2
	{
		t3 := int32(load32(m.memory[uint32(v0):]))
		t4 := int32(m.memory[uint32(t3)])
		switch t4 {
		default:
			t5 := m.t0[uint(v1)].(func(int32, int32, int32) int32)(v2, i32(1273083), i32(34))
			return t5
		case 1:
			t6 := m.t0[uint(v1)].(func(int32, int32, int32) int32)(v2, i32(1273117), i32(69))
			return t6
		case 2:
			t7 := m.t0[uint(v1)].(func(int32, int32, int32) int32)(v2, i32(1273186), i32(62))
			return t7
		case 3:
			t8 := m.t0[uint(v1)].(func(int32, int32, int32) int32)(v2, i32(1273248), i32(55))
			return t8
		case 4:
			t9 := m.t0[uint(v1)].(func(int32, int32, int32) int32)(v2, i32(1273303), i32(53))
			return t9
		case 5:
			t10 := m.t0[uint(v1)].(func(int32, int32, int32) int32)(v2, i32(1273356), i32(53))
			return t10
		case 6:
			t11 := m.t0[uint(v1)].(func(int32, int32, int32) int32)(v2, i32(1273409), i32(49))
			return t11
		case 7:
			t12 := m.t0[uint(v1)].(func(int32, int32, int32) int32)(v2, i32(1273458), i32(61))
			return t12
		case 8:
			t13 := m.t0[uint(v1)].(func(int32, int32, int32) int32)(v2, i32(1273519), i32(61))
			return t13
		}
	}
}
func (m *Module) fn934(v0, v1 int32) int32 {
	var v2, v3, v4 int32
	var v5 int64
	t0 := m.g0
	v2 = t0 - i32(32)
	m.g0 = v2
	t1 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	v3 = t1
	t2 := int32(load32(m.memory[uint32(v1):]))
	v1 = t2
	{
		t3 := int32(load32(m.memory[uint32(v0):]))
		v4 = t3
		t4 := int32(load32(m.memory[uint32(v4):]))
		v0 = t4
		p5 := i32(5)
		if v0 < i32(0) {
			p5 = v0 ^ i32(-0x80000000)
		}
		switch p5 {
		case 1:
			t6 := int32(load32(m.memory[int64(uint32(v3))+12:]))
			t7 := m.t0[uint(t6)].(func(int32, int32, int32) int32)(v1, i32(1273691), i32(50))
			v1 = t7
			goto l8
		case 2:
			t8 := int32(load32(m.memory[int64(uint32(v3))+12:]))
			t9 := m.t0[uint(t8)].(func(int32, int32, int32) int32)(v1, i32(1273741), i32(67))
			v1 = t9
			goto l8
		case 3:
			store32(m.memory[int64(uint32(v2))+12:], uint32(v4+i32(4)))
			store64(m.memory[int64(uint32(v2))+16:], uint64(int64(uint32(i32(31)))<<32|int64(uint32(v2+i32(12)))))
			t10 := m.fn45(v1, v3, i32(0x100ffe), v2+i32(16))
			v1 = t10
			goto l8
		case 4:
			store32(m.memory[int64(uint32(v2))+12:], uint32(v4+i32(4)))
			store64(m.memory[int64(uint32(v2))+16:], uint64(int64(uint32(i32(31)))<<32|int64(uint32(v2+i32(12)))))
			t11 := m.fn45(v1, v3, i32(1065362), v2+i32(16))
			v1 = t11
			goto l8
		case 5:
			store32(m.memory[int64(uint32(v2))+8:], uint32(v4))
			store32(m.memory[int64(uint32(v2))+12:], uint32(v4+i32(12)))
			t12 := v2
			v5 = int64(uint32(i32(31))) << 32
			store64(m.memory[int64(uint32(t12))+24:], uint64(v5|int64(uint32(v2+i32(12)))))
			store64(m.memory[int64(uint32(v2))+16:], uint64(v5|int64(uint32(v2+i32(8)))))
			t13 := m.fn45(v1, v3, i32(1066005), v2+i32(16))
			v1 = t13
			goto l8
		case 6:
			t14 := int32(load32(m.memory[int64(uint32(v3))+12:]))
			t15 := m.t0[uint(t14)].(func(int32, int32, int32) int32)(v1, i32(1273808), i32(44))
			v1 = t15
			goto l8
		case 7:
			t16 := int32(load32(m.memory[int64(uint32(v3))+12:]))
			t17 := m.t0[uint(t16)].(func(int32, int32, int32) int32)(v1, i32(1273852), i32(75))
			v1 = t17
			goto l8
		default:
			{
				t18 := int32(load32(m.memory[int64(uint32(v4))+4:]))
				if t18 == i32(-1) {
					goto l9
				}
				store32(m.memory[int64(uint32(v2))+12:], uint32(v4+i32(4)))
				store64(m.memory[int64(uint32(v2))+16:], uint64(int64(uint32(i32(31)))<<32|int64(uint32(v2+i32(12)))))
				t19 := m.fn45(v1, v3, i32(1066330), v2+i32(16))
				v1 = t19
				goto l8
			}
		l9:
			t20 := int32(load32(m.memory[int64(uint32(v3))+12:]))
			t21 := m.t0[uint(t20)].(func(int32, int32, int32) int32)(v1, i32(1273636), i32(55))
			v1 = t21
		}
	}
l8:
	m.g0 = v2 + i32(32)
	return v1
}
func (m *Module) fn935(v0, v1, v2, v3 int32) {
	var v4, v5 int32
	t0 := m.g0
	v4 = t0 - i32(16)
	m.g0 = v4
	t1 := int32(load32(m.memory[int64(uint32(i32(0)))+1293784:]))
	v5 = t1
	store32(m.memory[int64(uint32(i32(0)))+1293784:], uint32(v5+i32(1)))
	if v5 < i32(0) {
		goto l0
	}
	{
		t2 := int32(m.memory[int64(uint32(i32(0)))+1293776])
		if t2 != 0 {
			t5 := int32(load32(m.memory[int64(uint32(v1))+24:]))
			m.t0[uint(t5)].(func(int32, int32))(v4+i32(8), v0)
			panic("unreachable")
		}
		m.memory[int64(uint32(i32(0)))+1293776] = byte(i32(1))
		t3 := int32(load32(m.memory[int64(uint32(i32(0)))+1293772:]))
		store32(m.memory[int64(uint32(i32(0)))+1293772:], uint32(t3+i32(1)))
		t4 := int32(load32(m.memory[int64(uint32(i32(0)))+1293780:]))
		v5 = t4
		if v5 <= i32(-1) {
			goto l0
		}
		v1 = v5 + i32(1)
		if v1 >= v5 {
			goto l2
		}
		m.fn139(i32(1274352), i32(28), i32(1274380))
		panic("unreachable")
	}
l2:
	store32(m.memory[int64(uint32(i32(0)))+1293780:], uint32(v1+i32(-1)))
	if v1 <= i32(0) {
		m.fn27(i32(1275072), i32(77), i32(1275112))
		panic("unreachable")
	}
	m.memory[int64(uint32(i32(0)))+1293776] = byte(i32(0))
	if v2 != 0 {
		fn936()
		panic("unreachable")
	}
l0:
	panic("unreachable")
}
func fn936() {
	panic("unreachable")
}
func (m *Module) fn937(v0, v1 int32) {
	store32(m.memory[uint32(v0):], uint32(i32(0)))
}
func (m *Module) fn938(v0, v1 int32) {
	t0 := int64(load64(m.memory[int64(uint32(i32(0)))+1274200:]))
	store64(m.memory[int64(uint32(v0))+8:], uint64(t0))
	t1 := int64(load64(m.memory[int64(uint32(i32(0)))+1274192:]))
	store64(m.memory[uint32(v0):], uint64(t1))
}
func (m *Module) fn939(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7, v8 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	t1 := int32(load32(m.memory[uint32(v0):]))
	v0 = t1
	{
		{
			{
				{
					t2 := int32(load32(m.memory[int64(uint32(v1))+8:]))
					v3 = t2
					if v3&i32(0x2000000) != 0 {
						t11 := int32(load32(m.memory[uint32(v0):]))
						v3 = t11
						v0 = i32(9)
					l6:
						{
							t12 := int32(m.memory[int64(uint32(v3&i32(15)))+1099240])
							m.memory[uint32(v2+i32(6)+v0+i32(-2))] = byte(t12)
							v0 = v0 + i32(-1)
							v3 = int32(uint32(v3) >> 4)
							if v3 != 0 {
								goto l6
							}
						}
						t13 := m.fn679(v1, i32(1), i32(1122974), i32(2), v2+i32(6)+v0+i32(-1), i32(9)-v0)
						v0 = t13
						goto l7
					}
					t3 := int32(load32(m.memory[uint32(v0):]))
					v0 = t3
					if v3&i32(0x4000000) != 0 {
						goto l1
					}
					v3 = i32(10)
					{
						t4 := v0
						v4 = v0 >> 31
						v5 = t4 ^ v4 - v4
						if uint32(v5) < uint32(i32(1000)) {
							goto l2
						}
						v3 = i32(10)
					l3:
						{
							v6 = v2 + i32(6) + v3
							t5 := v6 + i32(-4)
							v4 = v5
							t6 := int32(uint32(v4) / uint32(i32(10000)))
							t7 := v4
							v5 = t6
							v7 = t7 - v5*i32(10000)
							t8 := int32(uint32(v7&i32(0xffff)) / uint32(i32(100)))
							v8 = t8
							t9 := int32(load16(m.memory[int64(uint32(v8<<1))+1100623:]))
							store16(m.memory[uint32(t5):], uint16(t9))
							t10 := int32(load16(m.memory[int64(uint32((v7-v8*i32(100))&i32(0xffff)<<1))+1100623:]))
							store16(m.memory[uint32(v6+i32(-2)):], uint16(t10))
							v3 = v3 + i32(-4)
							if uint32(v4) > uint32(i32(9999999)) {
								goto l3
							}
						}
					}
				l2:
					if uint32(v5) > uint32(i32(9)) {
						goto l4
					}
					v4 = v5
					goto l5
				}
			l4:
				t14 := v2 + i32(6)
				v3 = v3 + i32(-2)
				t15 := int32(uint32(v5&i32(0xffff)) / uint32(i32(100)))
				t16 := t14 + v3
				t17 := v5
				v4 = t15
				t18 := int32(load16(m.memory[int64(uint32((t17-v4*i32(100))&i32(0xffff)<<1))+1100623:]))
				store16(m.memory[uint32(t16):], uint16(t18))
			}
		l5:
			{
				if v0 == 0 {
					goto l8
				}
				if v4 == 0 {
					goto l9
				}
			l8:
				t19 := v2 + i32(6)
				v3 = v3 + i32(-1)
				t20 := int32(m.memory[int64(uint32(v4<<1))+1100624])
				m.memory[uint32(t19+v3)] = byte(t20)
			}
		l9:
			t21 := m.fn679(v1, int32(uint32(v0^i32(-1))>>31), i32(1), i32(0), v2+i32(6)+v3, i32(10)-v3)
			v0 = t21
			goto l7
		}
	l1:
		v3 = i32(9)
	l10:
		{
			t22 := int32(m.memory[int64(uint32(v0&i32(15)))+1122976])
			m.memory[uint32(v2+i32(6)+v3+i32(-2))] = byte(t22)
			v3 = v3 + i32(-1)
			v0 = int32(uint32(v0) >> 4)
			if v0 != 0 {
				goto l10
			}
		}
		t23 := m.fn679(v1, i32(1), i32(1122974), i32(2), v2+i32(6)+v3+i32(-1), i32(9)-v3)
		v0 = t23
	}
l7:
	m.g0 = v2 + i32(16)
	return v0
}
func (m *Module) fn940(v0, v1 int32) {
	m.memory[int64(uint32(i32(0)))+1294240] = byte(i32(1))
	panic("unreachable")
}
func (m *Module) fn941(v0 int32) {
	m.fn942(v0)
	panic("unreachable")
}
