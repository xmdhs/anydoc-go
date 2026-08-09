package core

import (
	"math/bits"
)

func (m *Module) fn1167(v0, v1, v2 int32) {
	var v3, v4, v5, v6 int32
	var v7, v8 int64
	var v9, v10, v11, v12, v13, v14, v15, v16 int32
	t0 := m.g0
	v3 = t0 - i32(48)
	m.g0 = v3
	t1 := int32(load32(m.memory[int64(uint32(v2))+16:]))
	t2 := v1
	v4 = t1
	p3 := i32(1)
	if uint32(v4) > uint32(i32(1)) {
		p3 = v4
	}
	v5 = p3
	t4 := int32(load32(m.memory[int64(uint32(v2))+12:]))
	t5 := int64(uint32(v5))
	v4 = t4
	p6 := i32(1)
	if uint32(v4) > uint32(i32(1)) {
		p6 = v4
	}
	v6 = p6
	t7 := int64(load64(m.memory[int64(uint32(v1))+32:]))
	t8 := t5 * int64(uint32(v6))
	v7 = t7
	v8 = t8 + v7 + i64(-1)
	p9 := v8
	if uint64(v8) < uint64(v7) {
		p9 = i64(-1)
	}
	v7 = p9
	store64(m.memory[int64(uint32(t2))+32:], uint64(v7))
	{
		if uint64(v7) > uint64(i64(4000000)) {
			m.fn51(v0+i32(4), i32(1076840), i32(47))
			store32(m.memory[int64(uint32(v0))+20:], uint32(i32(13)))
			store32(m.memory[int64(uint32(v0))+16:], uint32(i32(1075780)))
			store32(m.memory[uint32(v0):], uint32(i32(-0x7ffffffd)))
			m.fn969(v2)
			goto l3
		}
		t10 := m.fn1175(v1)
		v9 = t10
	l2:
		{
			t11 := int32(load32(m.memory[int64(uint32(v1))+44:]))
			t12 := int32(load32(m.memory[int64(uint32(v1))+48:]))
			t13 := m.fn857(t11, t12, v9, i32(1076776))
			v4 = t13
			store32(m.memory[int64(uint32(v3))+16:], uint32(v9))
			t14 := int32(load32(m.memory[int64(uint32(v4))+8:]))
			store32(m.memory[int64(uint32(v3))+20:], uint32(t14))
			m.fn653(v3+i32(4), v1, v3+i32(16))
			t15 := int32(load32(m.memory[int64(uint32(v3))+4:]))
			if t15 != i32(1) {
				t20 := int32(load32(m.memory[int64(uint32(v1))+44:]))
				v10 = t20
				t21 := int32(load32(m.memory[int64(uint32(v1))+48:]))
				t22 := v10
				v11 = t21
				t23 := m.fn857(t22, v11, v9, i32(1076808))
				t24 := int32(load32(m.memory[int64(uint32(t23))+8:]))
				v12 = t24
				v13 = i32(1)
			l6:
				if v6 != v13 {
					store32(m.memory[int64(uint32(v3))+24:], uint32(v9))
					store32(m.memory[int64(uint32(v3))+28:], uint32(v12+v13))
					t25 := m.fn650(v1, v3+i32(24))
					if t25 != 0 {
						goto l5
					}
					v13 = v13 + i32(1)
					goto l6
				}
				v13 = v6
				goto l5
			l5:
				v14 = i32(1)
			l10:
				if v14 != v5 {
					v15 = v14 + v9
					v16 = v14 + i32(1)
					v4 = v13
					v6 = v12
				l11:
					if v4 != 0 {
						store32(m.memory[int64(uint32(v3))+28:], uint32(v6))
						store32(m.memory[int64(uint32(v3))+24:], uint32(v15))
						t26 := m.fn650(v1, v3+i32(24))
						if t26 != 0 {
							goto l8
						}
						v4 = v4 + i32(-1)
						v6 = v6 + i32(1)
						goto l11
					}
					v14 = v16
					goto l10
				}
				v14 = v5
				goto l8
			l8:
				t27 := m.fn857(v10, v11, v9, i32(1076824))
				v4 = t27
				store32(m.memory[int64(uint32(v3))+40:], uint32(v14))
				store32(m.memory[int64(uint32(v3))+36:], uint32(v13))
				t28 := int32(load32(m.memory[int64(uint32(v2))+8:]))
				store32(m.memory[int64(uint32(v3))+32:], uint32(t28))
				t29 := int64(load64(m.memory[uint32(v2):]))
				store64(m.memory[int64(uint32(v3))+24:], uint64(t29))
				m.fn1172(v4, v3+i32(24))
				v15 = i32(0)
			l15:
				if v15 == v14 {
					store32(m.memory[uint32(v0):], uint32(i32(-1)))
					goto l3
				}
				v2 = v15 + v9
				v6 = i32(0)
			l14:
				v4 = v6
				if uint32(v4) >= uint32(v13) {
					v15 = v15 + i32(1)
					goto l15
				}
				v6 = i32(1)
				if v15|v4 == 0 {
					goto l14
				}
				m.fn1118(v3+i32(24), v1, v2, v4+v12, v9, v12)
				v6 = v4 + i32(1)
				goto l14
			}
			t16 := int64(load64(m.memory[int64(uint32(v3))+8:]))
			v7 = t16
			t17 := int32(load32(m.memory[int64(uint32(v1))+44:]))
			t18 := int32(load32(m.memory[int64(uint32(v1))+48:]))
			t19 := m.fn857(t17, t18, v9, i32(1076792))
			v4 = t19
			store32(m.memory[int64(uint32(v3))+24:], uint32(i32(-1)))
			store64(m.memory[int64(uint32(v3))+28:], uint64(v7))
			m.fn1172(v4, v3+i32(24))
			goto l2
		}
	}
l3:
	m.g0 = v3 + i32(48)
}
func (m *Module) fn1168(v0, v1 int32) {
	var v2, v3, v4 int32
	var v5 int64
	var v6, v7 int32
	var v8 int64
	var v9, v10, v11, v12, v13, v14, v15, v16 int32
	var v17, v18, v19 int64
	var v20 int32
	t0 := m.g0
	v2 = t0 - i32(112)
	m.g0 = v2
	m.fn22(v2+i32(32), i32(3))
	t1 := int64(load64(m.memory[int64(uint32(i32(0)))+1286264:]))
	store64(m.memory[int64(uint32(v2))+8:], uint64(t1))
	t2 := int64(load64(m.memory[int64(uint32(i32(0)))+1286256:]))
	store64(m.memory[uint32(v2):], uint64(t2))
	t3 := int64(load64(m.memory[int64(uint32(v2))+40:]))
	store64(m.memory[int64(uint32(v2))+24:], uint64(t3))
	t4 := int64(load64(m.memory[int64(uint32(v2))+32:]))
	store64(m.memory[int64(uint32(v2))+16:], uint64(t4))
	t5 := int32(load32(m.memory[uint32(v1):]))
	v3 = t5
	v4 = v3 + i32(8)
	t6 := int64(load64(m.memory[uint32(v3):]))
	v5 = (t6 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
	v6 = v2 + i32(16)
	t7 := int32(load32(m.memory[int64(uint32(v1))+12:]))
	v7 = t7
l12:
	if v7 == 0 {
		t9 := int32(load32(m.memory[uint32(v2):]))
		v3 = t9
		v9 = v3 + i32(8)
		t10 := int32(load32(m.memory[int64(uint32(v2))+4:]))
		t11 := v3
		v7 = t10
		v10 = t11 + v7 + i32(1)
		t12 := int64(load64(m.memory[uint32(v3):]))
		v8 = (t12 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
		t13 := int32(load32(m.memory[int64(uint32(v2))+12:]))
		v4 = t13
		v11 = i32(0)
		{
			if v7 == 0 {
				goto l3
			}
			m.fn39(v2+i32(92), i32(16), i32(8), v7+i32(1))
			t14 := int32(load32(m.memory[int64(uint32(v2))+96:]))
			store32(m.memory[int64(uint32(v2))+68:], uint32(t14))
			t15 := int32(load32(m.memory[int64(uint32(v2))+100:]))
			store32(m.memory[int64(uint32(v2))+72:], uint32(v3-t15))
			t16 := int32(load32(m.memory[int64(uint32(v2))+92:]))
			v11 = t16
		}
	l3:
		store32(m.memory[int64(uint32(v2))+56:], uint32(v4))
		store32(m.memory[int64(uint32(v2))+48:], uint32(v3))
		store32(m.memory[int64(uint32(v2))+44:], uint32(v10))
		store32(m.memory[int64(uint32(v2))+40:], uint32(v9))
		store64(m.memory[int64(uint32(v2))+32:], uint64(v8))
		store32(m.memory[int64(uint32(v2))+64:], uint32(v11))
		v12 = v1 + i32(40)
	l11:
		{
			if v4 == 0 {
				goto l4
			}
			t17 := m.fn1170(v2 + i32(32))
			v3 = t17
			t18 := int32(load32(m.memory[int64(uint32(v2))+56:]))
			t19 := v2
			v4 = t18 + i32(-1)
			store32(m.memory[int64(uint32(t19))+56:], uint32(v4))
			t20 := int32(load32(m.memory[uint32(v3+i32(-12)):]))
			v13 = t20
			if v13 == i32(-1) {
				goto l5
			}
			t21 := int32(load32(m.memory[uint32(v3+i32(-16)):]))
			v4 = t21
			t22 := int32(load32(m.memory[uint32(v3+i32(-8)):]))
			v14 = t22
			t23 := int32(load32(m.memory[uint32(v3+i32(-4)):]))
			t24 := v14
			v3 = t23
			m.fn1171(t24, v3)
			v10 = v14 + v3<<2
			v9 = v14
		l9:
			{
				if v9 == v10 {
					m.fn1173(v14, v13)
					t37 := int32(load32(m.memory[int64(uint32(v2))+56:]))
					v4 = t37
					goto l11
				}
				t25 := int32(load32(m.memory[uint32(v9):]))
				v11 = t25
			l10:
				{
					t26 := int32(load32(m.memory[int64(uint32(v1))+44:]))
					v3 = t26
					t27 := int32(load32(m.memory[int64(uint32(v1))+48:]))
					t28 := v3
					v7 = t27
					t29 := m.fn857(t28, v7, v4, i32(1076936))
					t30 := int32(load32(m.memory[int64(uint32(t29))+8:]))
					if uint32(t30) < uint32(v11) {
						t36 := m.fn857(v3, v7, v4, i32(1076984))
						v3 = t36
						store32(m.memory[int64(uint32(v2))+108:], uint32(i32(0)))
						store64(m.memory[int64(uint32(v2))+100:], uint64(i64(0)))
						store64(m.memory[int64(uint32(v2))+92:], uint64(i64(0x800000000)))
						m.fn1172(v3, v2+i32(92))
						goto l10
					}
					store32(m.memory[int64(uint32(v2))+88:], uint32(v11))
					store32(m.memory[int64(uint32(v2))+84:], uint32(v4))
					m.fn653(v2+i32(92), v1, v2+i32(84))
					t31 := int32(load32(m.memory[int64(uint32(v2))+92:]))
					if t31 == 0 {
						goto l8
					}
					t32 := int64(load64(m.memory[int64(uint32(v2))+96:]))
					v8 = t32
					t33 := int32(load32(m.memory[int64(uint32(v1))+44:]))
					t34 := int32(load32(m.memory[int64(uint32(v1))+48:]))
					t35 := m.fn857(t33, t34, v4, i32(1076968))
					v3 = t35
					store32(m.memory[int64(uint32(v2))+92:], uint32(i32(-1)))
					store64(m.memory[int64(uint32(v2))+96:], uint64(v8))
					m.fn1172(v3, v2+i32(92))
					v9 = v9 + i32(4)
					goto l9
				}
			}
		l8:
		}
		m.fn153(i32(1076952))
		panic("unreachable")
	}
	v8 = v5
l2:
	{
		if v8 != i64(0) {
			v7 = v7 + i32(-1)
			v5 = (v8 + i64(-1)) & v8
			v9 = v3 - int32(int64(bits.TrailingZeros64(uint64(v8))))<<1&i32(240)
			t38 := int32(load32(m.memory[uint32(v9+i32(-16)):]))
			v11 = t38
			t39 := int32(load32(m.memory[int64(uint32(v1))+48:]))
			if uint32(v11) >= uint32(t39) {
				goto l12
			}
			t40 := int32(load32(m.memory[uint32(v9+i32(-12)):]))
			v15 = t40
			t41 := int64(load64(m.memory[int64(uint32(v2))+16:]))
			t42 := int64(load64(m.memory[int64(uint32(v2))+24:]))
			t43 := m.fn66(t41, t42, v11)
			v8 = t43
			t44 := int32(load32(m.memory[int64(uint32(v2))+4:]))
			v13 = t44
			t45 := v13
			v16 = int32(v8)
			v10 = t45 & v16
			v17 = int64(uint64(v8)>>25) & i64(127) * i64(72340172838076673)
			v12 = i32(0)
			t46 := int32(load32(m.memory[uint32(v2):]))
			v9 = t46
		l18:
			{
				t47 := int64(load64(m.memory[uint32(v9+v10):]))
				v18 = t47
				v19 = v18 ^ v17
				v19 = (v19 ^ i64(-1)) & (v19 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
				{
				l15:
					{
						if v19 == 0 {
							goto l13
						}
						v14 = v9 - (int32(uint32(int64(bits.TrailingZeros64(uint64(v19))))>>3)+v10)&v13<<4
						t48 := int32(load32(m.memory[uint32(v14+i32(-16)):]))
						if t48 == v11 {
							goto l14
						}
						v19 = (v19 + i64(-1)) & v19
						goto l15
					}
				l13:
					if v18&(v18<<1)&i64(-0x7f7f7f7f7f7f7f80) == 0 {
						t59 := v10
						v12 = v12 + i32(8)
						v10 = (t59 + v12) & v13
						goto l18
					}
					{
						t49 := int32(load32(m.memory[int64(uint32(v2))+8:]))
						v12 = t49
						if v12 != 0 {
							goto l17
						}
						_ = m.fn704(v2, v6)
						t51 := int32(load32(m.memory[int64(uint32(v2))+8:]))
						v12 = t51
						t52 := int32(load32(m.memory[int64(uint32(v2))+4:]))
						v13 = t52
						t53 := int32(load32(m.memory[uint32(v2):]))
						v9 = t53
					}
				l17:
					t54 := m.fn26(v9, v13, v8)
					t55 := v9
					v10 = t54
					v14 = t55 + v10
					t56 := int32(m.memory[uint32(v14)])
					v20 = t56
					t57 := v14
					v16 = int32(uint32(v16) >> 25)
					m.memory[uint32(t57)] = byte(v16)
					m.memory[uint32(v9+v13&(v10+i32(-8))+i32(8))] = byte(v16)
					v14 = v9 - v10<<4
					store32(m.memory[uint32(v14+i32(-4)):], uint32(i32(0)))
					store64(m.memory[uint32(v14+i32(-12)):], uint64(i64(0x400000000)))
					store32(m.memory[uint32(v14+i32(-16)):], uint32(v11))
					t58 := int32(load32(m.memory[int64(uint32(v2))+12:]))
					store32(m.memory[int64(uint32(v2))+12:], uint32(t58+i32(1)))
					store32(m.memory[int64(uint32(v2))+8:], uint32(v12-v20&i32(1)))
				}
			l14:
				m.fn176(v14+i32(-12), v15)
				goto l12
			}
		}
		v3 = v3 + i32(-128)
		t8 := int64(load64(m.memory[uint32(v4):]))
		v8 = (t8 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
		v4 = v4 + i32(8)
		goto l2
	}
l5:
	if v4 == 0 {
		goto l4
	}
l19:
	{
		t60 := m.fn1170(v2 + i32(32))
		v4 = t60
		t61 := int32(load32(m.memory[int64(uint32(v2))+56:]))
		t62 := v2
		v3 = t61 + i32(-1)
		store32(m.memory[int64(uint32(t62))+56:], uint32(v3))
		t63 := int32(load32(m.memory[uint32(v4+i32(-12)):]))
		t64 := int32(load32(m.memory[uint32(v4+i32(-8)):]))
		m.fn188(t63, t64)
		if v3 != 0 {
			goto l19
		}
	}
l4:
	{
		t65 := int32(load32(m.memory[int64(uint32(v2))+64:]))
		v4 = t65
		if v4 == 0 {
			goto l20
		}
		t66 := int32(load32(m.memory[int64(uint32(v2))+72:]))
		t67 := int32(load32(m.memory[int64(uint32(v2))+68:]))
		m.fn40(t66, v4, t67)
	}
l20:
	t68 := int32(load32(m.memory[int64(uint32(v1))+48:]))
	v13 = t68
	t69 := int32(load32(m.memory[int64(uint32(v1))+44:]))
	v15 = t69
l31:
	{
		if v13 == 0 {
			goto l21
		}
		v4 = v15 + v13*i32(12)
		if v4+i32(-12) == 0 {
			goto l21
		}
		t70 := int32(load32(m.memory[uint32(v4+i32(-8)):]))
		v9 = t70
		t71 := int32(load32(m.memory[uint32(v4+i32(-4)):]))
		v10 = v9 + t71*i32(20)
	l23:
		{
			v4 = v9
			if v4 == v10 {
				t98 := v1
				v13 = v13 + i32(-1)
				store32(m.memory[int64(uint32(t98))+48:], uint32(v13))
				t99 := v2
				v4 = v15 + v13*i32(12)
				t100 := int64(load64(m.memory[uint32(v4):]))
				v8 = t100
				store64(m.memory[int64(uint32(t99))+32:], uint64(v8))
				t101 := int32(load32(m.memory[int64(uint32(v4))+8:]))
				store32(m.memory[int64(uint32(v2))+40:], uint32(t101))
				if int32(v8) == i32(-1) {
					goto l31
				}
				m.fn973(v2 + i32(32))
				goto l31
			}
			v9 = v4 + i32(20)
			t72 := int32(load32(m.memory[uint32(v4):]))
			if t72 == i32(-1) {
				goto l23
			}
			t73 := int32(load32(m.memory[uint32(v4+i32(8)):]))
			v3 = t73 << 5
			t74 := int32(load32(m.memory[uint32(v4+i32(4)):]))
			v4 = t74
		l24:
			{
				if v3 == 0 {
					goto l23
				}
				t75 := int32(load32(m.memory[uint32(v4):]))
				if t75 != i32(-0x80000000) {
					goto l21
				}
				v3 = v3 + i32(-32)
				t76 := int32(load32(m.memory[int64(uint32(v4))+12:]))
				v7 = t76
				t77 := int32(load32(m.memory[int64(uint32(v4))+8:]))
				v11 = t77
				v4 = v4 + i32(32)
				t78 := m.fn23(v11, v7)
				if t78 != 0 {
					goto l24
				}
			}
		}
	}
l21:
	v14 = i32(0)
l27:
	{
		if v13 == v14 {
			m.memory[int64(uint32(v0))+16] = byte(i32(0))
			store32(m.memory[int64(uint32(v0))+12:], uint32(i32(0)))
			t94 := int32(load32(m.memory[int64(uint32(v12))+8:]))
			store32(m.memory[int64(uint32(v0))+8:], uint32(t94))
			t95 := int64(load64(m.memory[uint32(v12):]))
			store64(m.memory[uint32(v0):], uint64(t95))
			t96 := int32(load32(m.memory[uint32(v1):]))
			t97 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			m.fn1174(t96, t97)
			m.g0 = v2 + i32(112)
			return
		}
		t79 := m.fn857(v15, v13, v14, i32(1076888))
		t80 := int32(load32(m.memory[int64(uint32(t79))+8:]))
		v4 = t80
		v9 = v13 - v14
		v16 = v14 + i32(1)
		v3 = i32(0)
		v7 = i32(12)
	l30:
		if v4 != 0 {
			t81 := m.fn857(v15, v13, v14, i32(1076904))
			t82 := v3
			v11 = t81
			t83 := int32(load32(m.memory[int64(uint32(v11))+8:]))
			v10 = t83
			if uint32(t82) >= uint32(v10) {
				m.fn158(v3, v10, i32(1076920))
				panic("unreachable")
			}
			{
				t84 := int32(load32(m.memory[int64(uint32(v11))+4:]))
				v11 = t84 + v7
				t85 := int32(load32(m.memory[uint32(v11+i32(-12)):]))
				if t85 == i32(-1) {
					goto l29
				}
				t86 := int32(load32(m.memory[uint32(v11):]))
				t87 := v11
				t88 := v4
				v10 = t86
				p89 := v10
				if uint32(v4) < uint32(v10) {
					p89 = t88
				}
				store32(m.memory[uint32(t87):], uint32(p89))
				v11 = v11 + i32(4)
				t90 := int32(load32(m.memory[uint32(v11):]))
				t91 := v11
				t92 := v9
				v11 = t90
				p93 := v11
				if uint32(v9) < uint32(v11) {
					p93 = t92
				}
				store32(m.memory[uint32(t91):], uint32(p93))
			}
		l29:
			v3 = v3 + i32(1)
			v4 = v4 + i32(-1)
			v7 = v7 + i32(20)
			goto l30
		}
		v14 = v16
		goto l27
	}
}
func (m *Module) fn1169(v0, v1 int32) {
	var v2 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		v2 = t0
		t1 := int32(load32(m.memory[uint32(v0):]))
		if v2 != t1 {
			goto l0
		}
		m.fn272(v0)
	}
l0:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v2+i32(1)))
	t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	v0 = t2 + v2*i32(12)
	t3 := int64(load64(m.memory[uint32(v1):]))
	store64(m.memory[uint32(v0):], uint64(t3))
	t4 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	store32(m.memory[int64(uint32(v0))+8:], uint32(t4))
}
func (m *Module) fn1170(v0 int32) int32 {
	var v1, v2 int32
	var v3 int64
	var v4 int32
	t0 := int32(load32(m.memory[int64(uint32(v0))+16:]))
	v1 = t0
	t1 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	v2 = t1
	t2 := int64(load64(m.memory[uint32(v0):]))
	v3 = t2
	var _ int32
l1:
	if v3 == 0 {
		t4 := v0
		v1 = v1 + i32(-128)
		store32(m.memory[int64(uint32(t4))+16:], uint32(v1))
		t5 := v0
		v4 = v2 + i32(8)
		store32(m.memory[int64(uint32(t5))+8:], uint32(v4))
		t6 := int64(load64(m.memory[uint32(v2):]))
		t7 := v0
		v3 = (t6 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
		store64(m.memory[uint32(t7):], uint64(v3))
		v2 = v4
		goto l1
	}
	store64(m.memory[uint32(v0):], uint64((v3+i64(-1))&v3))
	return v1 - int32(int64(bits.TrailingZeros64(uint64(v3))))<<1&i32(240)
}
func (m *Module) fn1171(v0, v1 int32) {
	if uint32(v1) < uint32(i32(2)) {
		return
	}
	if uint32(v1) < uint32(i32(21)) {
		goto l1
	}
	m.fn1178(v0, v1)
	return
l1:
	m.fn1179(v0, v1, i32(1))
}
func (m *Module) fn1172(v0, v1 int32) {
	var v2 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		v2 = t0
		t1 := int32(load32(m.memory[uint32(v0):]))
		if v2 != t1 {
			goto l0
		}
		m.fn418(v0)
	}
l0:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v2+i32(1)))
	t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	v0 = t2 + v2*i32(20)
	t3 := int64(load64(m.memory[uint32(v1):]))
	store64(m.memory[uint32(v0):], uint64(t3))
	t4 := int64(load64(m.memory[int64(uint32(v1))+8:]))
	store64(m.memory[int64(uint32(v0))+8:], uint64(t4))
	t5 := int32(load32(m.memory[int64(uint32(v1))+16:]))
	store32(m.memory[int64(uint32(v0))+16:], uint32(t5))
}
func (m *Module) fn1173(v0, v1 int32) {
	m.fn188(v1, v0)
}
func (m *Module) fn1174(v0, v1 int32) {
	var v2 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	{
		if v1 == 0 {
			goto l0
		}
		m.fn39(v2+i32(4), i32(16), i32(8), v1+i32(1))
		t1 := int32(load32(m.memory[int64(uint32(v2))+12:]))
		t2 := int32(load32(m.memory[int64(uint32(v2))+4:]))
		t3 := int32(load32(m.memory[int64(uint32(v2))+8:]))
		m.fn40(v0-t1, t2, t3)
	}
l0:
	m.g0 = v2 + i32(16)
}
func (m *Module) fn1175(v0 int32) int32 {
	var v1, v2 int32
	t0 := m.g0
	v1 = t0 - i32(16)
	m.g0 = v1
	{
		t1 := int32(load32(m.memory[int64(uint32(v0))+48:]))
		v2 = t1
		if v2 != 0 {
			goto l0
		}
		store32(m.memory[int64(uint32(v1))+12:], uint32(i32(0)))
		store64(m.memory[int64(uint32(v1))+4:], uint64(i64(0x400000000)))
		m.fn1169(v0+i32(40), v1+i32(4))
		t2 := int32(load32(m.memory[int64(uint32(v0))+48:]))
		v2 = t2
	}
l0:
	m.g0 = v1 + i32(16)
	return v2 + i32(-1)
}
func (m *Module) fn1176(v0, v1 int32) int32 {
	var v2, v3 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	{
		t1 := int32(load32(m.memory[uint32(v0):]))
		v3 = t1
		p2 := i32(1)
		if v3 < i32(0) {
			p2 = v3 ^ i32(-0x80000000)
		}
		switch p2 {
		default:
			store32(m.memory[int64(uint32(v2))+12:], uint32(v0+i32(4)))
			t3 := m.fn264(v1, i32(1285157), i32(11), v2+i32(12), i32(76))
			v0 = t3
			goto l6
		case 1:
			store32(m.memory[int64(uint32(v2))+12:], uint32(v0))
			t4 := m.fn459(v1, i32(1086568), i32(9), i32(1086577), i32(4), v0+i32(12), i32(162), i32(1086581), i32(6), v2+i32(12), i32(76))
			v0 = t4
			goto l6
		case 2:
			t5 := int32(load32(m.memory[uint32(v1):]))
			t6 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t7 := int32(load32(m.memory[int64(uint32(t6))+12:]))
			t8 := m.t0[uint(t7)].(func(int32, int32, int32) int32)(t5, i32(1086587), i32(9))
			v0 = t8
			goto l6
		case 3:
			store32(m.memory[int64(uint32(v2))+12:], uint32(v0+i32(4)))
			t9 := m.fn459(v1, i32(1086596), i32(13), i32(1086609), i32(5), v0+i32(16), i32(71), i32(1086581), i32(6), v2+i32(12), i32(76))
			v0 = t9
			goto l6
		case 4:
			store32(m.memory[int64(uint32(v2))+12:], uint32(v0+i32(4)))
			t10 := m.fn283(v1, i32(1086614), i32(11), i32(1086577), i32(4), v2+i32(12), i32(76))
			v0 = t10
			goto l6
		case 5:
			store32(m.memory[int64(uint32(v2))+12:], uint32(v0+i32(4)))
			t11 := m.fn264(v1, i32(1100477), i32(2), v2+i32(12), i32(68))
			v0 = t11
		}
	}
l6:
	m.g0 = v2 + i32(16)
	return v0
}
func (m *Module) fn1177(v0, v1 int32) int32 {
	var v2 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	{
		{
			t1 := int32(load32(m.memory[uint32(v0):]))
			if t1 == i32(-1) {
				goto l0
			}
			store32(m.memory[int64(uint32(v2))+12:], uint32(v0))
			t2 := m.fn264(v1, i32(1087236), i32(4), v2+i32(12), i32(76))
			v0 = t2
			goto l1
		}
	l0:
		t3 := int32(load32(m.memory[uint32(v1):]))
		t4 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t5 := int32(load32(m.memory[int64(uint32(t4))+12:]))
		t6 := m.t0[uint(t5)].(func(int32, int32, int32) int32)(t3, i32(1087232), i32(4))
		v0 = t6
	}
l1:
	m.g0 = v2 + i32(16)
	return v0
}
func (m *Module) fn1178(v0, v1 int32) {
	var v2, v3, v4, v5, v6, v7, v8 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	{
		t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		t2 := int32(load32(m.memory[uint32(v0):]))
		if uint32(t1) < uint32(t2) {
			v3 = v0 + i32(4)
			v4 = i32(2)
		l5:
			{
				if v1 == v4 {
					t7 := v2
					t8 := v0
					v6 = int32(uint32(v1) >> 1)
					m.fn1846(t7, t8, v6, v6, i32(1301108))
					t9 := int32(load32(m.memory[int64(uint32(v2))+4:]))
					v7 = t9
					t10 := int32(load32(m.memory[uint32(v2):]))
					v3 = t10
					t11 := v2
					t12 := v0 + v1<<2
					v4 = v6 << 2
					m.fn1846(t11, t12-v4, v6, v6, i32(1301124))
					t13 := int32(load32(m.memory[uint32(v2):]))
					v1 = v4 + t13 + i32(-4)
					v4 = i32(0)
					t14 := int32(load32(m.memory[int64(uint32(v2))+4:]))
					t15 := v6 + i32(-1)
					v8 = t14
					var p16 int32
					if uint32(t15) < uint32(v8) {
						p16 = 1
					}
					v0 = p16
				l8:
					v5 = v6 + v4
					if v5 == 0 {
						goto l1
					}
					if v7+v4 == 0 {
						m.fn158(v7, v7, i32(1301140))
						panic("unreachable")
					}
					{
						if v0 == 0 {
							m.fn158(v5+i32(-1), v8, i32(1301156))
							panic("unreachable")
						}
						t17 := int32(load32(m.memory[uint32(v3):]))
						v5 = t17
						t18 := int32(load32(m.memory[uint32(v1):]))
						store32(m.memory[uint32(v3):], uint32(t18))
						store32(m.memory[uint32(v1):], uint32(v5))
						v3 = v3 + i32(4)
						v1 = v1 + i32(-4)
						v4 = v4 + i32(-1)
						goto l8
					}
				}
				v5 = v3 + i32(4)
				t5 := int32(load32(m.memory[uint32(v5):]))
				t6 := int32(load32(m.memory[uint32(v3):]))
				if uint32(t5) >= uint32(t6) {
					goto l2
				}
				v4 = v4 + i32(1)
				v3 = v5
				goto l5
			}
		}
		v3 = v0 + i32(4)
		v4 = i32(2)
	l3:
		{
			if v1 == v4 {
				goto l1
			}
			v5 = v3 + i32(4)
			t3 := int32(load32(m.memory[uint32(v5):]))
			t4 := int32(load32(m.memory[uint32(v3):]))
			if uint32(t3) < uint32(t4) {
				goto l2
			}
			v4 = v4 + i32(1)
			v3 = v5
			goto l3
		}
	}
l2:
	m.fn1845(v0, v1, i32(0), int32(bits.LeadingZeros32(uint32(v1|i32(1))))<<1^i32(62))
	goto l1
l1:
	m.g0 = v2 + i32(16)
}
func (m *Module) fn1179(v0, v1, v2 int32) {
	var v3, v4, v5, v6, v7 int32
	{
		if uint32(v2+i32(-1)) >= uint32(v1) {
			panic("unreachable")
		}
		t0 := v0
		v3 = v2 << 2
		v4 = t0 + v3
		v5 = v0 + v1<<2
	l6:
		if v4 == v5 {
			return
		}
		{
			t1 := int32(load32(m.memory[uint32(v4):]))
			v6 = t1
			t2 := int32(load32(m.memory[uint32(v4+i32(-4)):]))
			t3 := v6
			v1 = t2
			if uint32(t3) >= uint32(v1) {
				goto l2
			}
			v2 = v3
		l5:
			{
				v7 = v0 + v2
				store32(m.memory[uint32(v7):], uint32(v1))
				if v2 != i32(4) {
					goto l3
				}
				v2 = v0
				goto l4
			l3:
				v2 = v2 + i32(-4)
				t4 := int32(load32(m.memory[uint32(v7+i32(-8)):]))
				t5 := v6
				v1 = t4
				if uint32(t5) < uint32(v1) {
					goto l5
				}
			}
			v2 = v0 + v2
		l4:
			store32(m.memory[uint32(v2):], uint32(v6))
		}
	l2:
		v3 = v3 + i32(4)
		v4 = v4 + i32(4)
		goto l6
	}
}
func (m *Module) fn1180(v0, v1, v2 int32) {
	var v3 int32
	t0 := m.g0
	v3 = t0 - i32(48)
	m.g0 = v3
	{
		t1 := m.fn159(v1, v2, i32(1070608), i32(8))
		if t1 != 0 {
			goto l0
		}
		store32(m.memory[uint32(v0):], uint32(i32(-1)))
		goto l1
	}
l0:
	store64(m.memory[int64(uint32(v3))+16:], uint64(i64(0)))
	store32(m.memory[int64(uint32(v3))+12:], uint32(v2))
	store32(m.memory[int64(uint32(v3))+8:], uint32(v1))
	m.fn578(v3+i32(28), v3+i32(8))
	{
		t2 := int32(load32(m.memory[int64(uint32(v3))+28:]))
		if t2 != 0 {
			m.fn1027(v3 + i32(28))
			goto l5
		}
		t3 := int32(load32(m.memory[int64(uint32(v3))+36:]))
		store32(m.memory[int64(uint32(v3))+44:], uint32(t3))
		t4 := int32(load32(m.memory[int64(uint32(v3))+32:]))
		t5 := v3
		v2 = t4
		store32(m.memory[int64(uint32(t5))+40:], uint32(v2))
		{
			t6 := m.fn636(v2, i32(1084557), i32(14))
			if t6 != 0 {
				goto l3
			}
			t7 := m.fn636(v2, i32(1072157), i32(16))
			if t7 == 0 {
				goto l4
			}
		}
	l3:
		store32(m.memory[uint32(v0):], uint32(i32(-0x7ffffffe)))
		m.fn956(v3 + i32(40))
		goto l1
	}
l4:
	m.fn956(v3 + i32(40))
l5:
	m.fn1200(v0, i32(1084571), i32(81))
l1:
	m.g0 = v3 + i32(48)
}
func (m *Module) fn1181(v0, v1, v2 int32) {
	{
		t0 := int32(load32(m.memory[uint32(v1):]))
		if t0 == i32(-1) {
			goto l0
		}
		t1 := int64(load64(m.memory[int64(uint32(v1))+16:]))
		store64(m.memory[int64(uint32(v0))+16:], uint64(t1))
		t2 := int64(load64(m.memory[int64(uint32(v1))+8:]))
		store64(m.memory[int64(uint32(v0))+8:], uint64(t2))
		t3 := int64(load64(m.memory[uint32(v1):]))
		store64(m.memory[uint32(v0):], uint64(t3))
		m.fn785(v2)
		return
	}
l0:
	t4 := int64(load64(m.memory[int64(uint32(v2))+16:]))
	store64(m.memory[int64(uint32(v0))+16:], uint64(t4))
	t5 := int64(load64(m.memory[int64(uint32(v2))+8:]))
	store64(m.memory[int64(uint32(v0))+8:], uint64(t5))
	t6 := int64(load64(m.memory[uint32(v2):]))
	store64(m.memory[uint32(v0):], uint64(t6))
}
func (m *Module) fn1182(v0, v1, v2 int32) {
	t0 := int32(load32(m.memory[uint32(v1):]))
	if t0 != 0 {
		m.fn1326(v2)
		panic("unreachable")
	}
	store32(m.memory[uint32(v1):], uint32(i32(-1)))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
	store32(m.memory[uint32(v0):], uint32(v1+i32(8)))
}
func (m *Module) fn1183(v0, v1, v2 int32) {
	var v3 int32
	t0 := m.g0
	v3 = t0 - i32(64)
	m.g0 = v3
	store32(m.memory[int64(uint32(v3))+12:], uint32(v2))
	store32(m.memory[int64(uint32(v3))+8:], uint32(v1))
	m.fn778(v3+i32(16), v1, v2, i32(47))
	{
		t1 := int32(load32(m.memory[int64(uint32(v3))+16:]))
		v2 = t1
		if v2 == 0 {
			goto l0
		}
		t2 := int32(load32(m.memory[int64(uint32(v3))+20:]))
		store32(m.memory[int64(uint32(v3))+36:], uint32(t2))
		store32(m.memory[int64(uint32(v3))+32:], uint32(v2))
		t3 := int64(load64(m.memory[int64(uint32(v3))+24:]))
		store64(m.memory[int64(uint32(v3))+40:], uint64(t3))
		store32(m.memory[int64(uint32(v3))+60:], uint32(i32(1)))
		store32(m.memory[int64(uint32(v3))+52:], uint32(i32(1)))
		store32(m.memory[int64(uint32(v3))+56:], uint32(v3+i32(40)))
		store32(m.memory[int64(uint32(v3))+48:], uint32(v3+i32(32)))
		m.fn73(v0, i32(1066707), v3+i32(48))
		goto l1
	}
l0:
	store32(m.memory[int64(uint32(v3))+52:], uint32(i32(1)))
	store32(m.memory[int64(uint32(v3))+48:], uint32(v3+i32(8)))
	m.fn73(v0, i32(1066724), v3+i32(48))
l1:
	m.g0 = v3 + i32(64)
}
func (m *Module) fn1184(v0, v1, v2, v3, v4, v5, v6, v7 int32) {
	var v8 int32
	var v9 int64
	t0 := m.g0
	v8 = t0 - i32(80)
	m.g0 = v8
	{
		{
			t1 := m.fn1039(v1, v4, v5)
			v5 = t1
			if v5 != 0 {
				goto l0
			}
			v5 = i32(0)
			goto l1
		}
	l0:
		t2 := int32(load32(m.memory[int64(uint32(v5))+8:]))
		v4 = t2
		t3 := int32(load32(m.memory[int64(uint32(v5))+4:]))
		v5 = t3
	}
l1:
	t5 := v8 + i32(12)
	t6 := v2
	t7 := v3
	p4 := v6
	if v5 != 0 {
		p4 = v5
	}
	p8 := v7
	if v5 != 0 {
		p8 = v4
	}
	m.fn774(t5, t6, t7, p4, p8)
	v5 = v8 + i32(12) + i32(4)
	{
		t9 := int32(load32(m.memory[int64(uint32(v8))+12:]))
		if t9 != 0 {
			goto l2
		}
		t10 := int64(load64(m.memory[uint32(v5):]))
		store64(m.memory[uint32(v0):], uint64(t10))
		t11 := int32(load32(m.memory[int64(uint32(v5))+8:]))
		store32(m.memory[int64(uint32(v0))+8:], uint32(t11))
		t12 := int32(load32(m.memory[int64(uint32(v8))+28:]))
		t13 := int32(load32(m.memory[int64(uint32(v8))+32:]))
		m.fn134(t12, t13)
		goto l3
	}
l2:
	m.fn774(v8+i32(52), v2, v3, v6, v7)
	{
		{
			t14 := int32(load32(m.memory[int64(uint32(v8))+52:]))
			if t14 != 0 {
				t20 := int64(load64(m.memory[int64(uint32(v8))+60:]))
				store64(m.memory[int64(uint32(v8))+40:], uint64(t20))
				t21 := int32(load32(m.memory[int64(uint32(v8))+68:]))
				store32(m.memory[int64(uint32(v8))+48:], uint32(t21))
				t22 := int32(load32(m.memory[int64(uint32(v8))+56:]))
				v3 = t22
				if v3 == i32(-1) {
					goto l5
				}
				t23 := int64(load64(m.memory[int64(uint32(v8))+72:]))
				v9 = t23
				store32(m.memory[int64(uint32(v8))+52:], uint32(v3))
				t24 := int64(load64(m.memory[int64(uint32(v8))+40:]))
				store64(m.memory[int64(uint32(v8))+56:], uint64(t24))
				t25 := int32(load32(m.memory[int64(uint32(v8))+48:]))
				store32(m.memory[int64(uint32(v8))+64:], uint32(t25))
				store64(m.memory[int64(uint32(v8))+68:], uint64(v9))
				m.fn51(v0, v6, v7)
				m.fn785(v8 + i32(52))
				goto l6
			}
			t15 := v8
			v7 = v8 + i32(52) + i32(4)
			t16 := int64(load64(m.memory[uint32(v7):]))
			store64(m.memory[int64(uint32(t15))+40:], uint64(t16))
			t17 := int32(load32(m.memory[int64(uint32(v7))+8:]))
			store32(m.memory[int64(uint32(v8))+48:], uint32(t17))
			t18 := int32(load32(m.memory[int64(uint32(v8))+68:]))
			t19 := int32(load32(m.memory[int64(uint32(v8))+72:]))
			m.fn134(t18, t19)
			goto l5
		}
	l5:
		t26 := int32(load32(m.memory[int64(uint32(v8))+48:]))
		store32(m.memory[int64(uint32(v0))+8:], uint32(t26))
		t27 := int64(load64(m.memory[int64(uint32(v8))+40:]))
		store64(m.memory[uint32(v0):], uint64(t27))
	}
l6:
	m.fn785(v5)
l3:
	m.g0 = v8 + i32(80)
}
func (m *Module) fn1185(v0 int32) {
	var v1 int32
	var v2, v3 int64
	t0 := m.g0
	v1 = t0 - i32(16)
	m.g0 = v1
	m.fn22(v1, i32(3))
	t1 := int64(load64(m.memory[uint32(v1):]))
	v2 = t1
	t2 := int64(load64(m.memory[int64(uint32(v1))+8:]))
	v3 = t2
	t3 := int64(load64(m.memory[int64(uint32(i32(0)))+1286264:]))
	store64(m.memory[int64(uint32(v0))+8:], uint64(t3))
	t4 := int64(load64(m.memory[int64(uint32(i32(0)))+1286256:]))
	store64(m.memory[uint32(v0):], uint64(t4))
	store64(m.memory[int64(uint32(v0))+24:], uint64(v3))
	store64(m.memory[int64(uint32(v0))+16:], uint64(v2))
	m.g0 = v1 + i32(16)
}
func (m *Module) fn1186(v0 int32) int32 {
	var v1, v2, v3, v4 int32
	v1 = v0 + i32(8)
	t0 := int32(load32(m.memory[uint32(v0):]))
	v2 = t0
	t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	v3 = t1
l1:
	{
		v4 = v2
		if v4 != v3 {
			goto l0
		}
		return i32(0)
	l0:
		t2 := v0
		v2 = v4 + i32(44)
		store32(m.memory[uint32(t2):], uint32(v2))
		t3 := int32(load32(m.memory[uint32(v4):]))
		if t3 == i32(-1) {
			goto l1
		}
		t4 := m.fn844(v1, v4)
		if t4 == 0 {
			goto l1
		}
	}
	return v4
}
func (m *Module) fn1187(v0, v1 int32) int32 {
	var v2, v3, v4 int32
	t0 := m.fn1318(v0, v1, i32(1073722), i32(6))
	v2 = t0
	t1 := m.fn1318(v0, v1, i32(1073728), i32(7))
	v3 = t1 & i32(255)
	t2 := m.fn1318(v0, v1, i32(1073720), i32(1))
	v4 = t2 & i32(255)
	t3 := m.fn1318(v0, v1, i32(1073721), i32(1))
	v0 = t3 & i32(255)
	v1 = i32(33619968)
	{
		switch v2 & i32(255) {
		case 0:
			goto l0
		default:
			goto l1
		case 2:
			if v3 != i32(2) {
				goto l0
			}
			v1 = i32(33685504)
			goto l1
		}
	l0:
		p4 := i32(0x2000000)
		if v3&i32(1) != 0 {
			p4 = i32(33619968)
		}
		v1 = p4
	}
l1:
	return v1 | v0<<8 | v4
}
func (m *Module) fn1188(v0 int32) int32 {
	t0 := fn1319(v0, i32(0))
	return t0
}
func (m *Module) fn1189(v0, v1 int32) {
	var v2 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	{
		if v1 == 0 {
			goto l0
		}
		m.fn39(v2+i32(4), i32(20), i32(8), v1+i32(1))
		t1 := int32(load32(m.memory[int64(uint32(v2))+12:]))
		t2 := int32(load32(m.memory[int64(uint32(v2))+4:]))
		t3 := int32(load32(m.memory[int64(uint32(v2))+8:]))
		m.fn40(v0-t1, t2, t3)
	}
l0:
	m.g0 = v2 + i32(16)
}
func (m *Module) fn1190(v0, v1, v2 int32) {
	var v3, v4 int32
	var v5, v6 int64
	t0 := m.g0
	v3 = t0 - i32(48)
	m.g0 = v3
	{
		switch v2 {
		case 0:
			m.memory[int64(uint32(v0))+1] = byte(i32(0))
			v2 = i32(1)
			goto l3
		case 1:
			t1 := int32(m.memory[uint32(v1)])
			v4 = t1
			switch v4 + i32(-43) {
			case 0, 2:
				goto l4
			default:
				goto l5
			}
		default:
			t2 := int32(m.memory[uint32(v1)])
			v4 = t2
		}
	l5:
		t3 := v1
		var p4 int32
		if v4&i32(255) == i32(43) {
			p4 = 1
		}
		v4 = p4
		v1 = t3 + v4
		v2 = v2 - v4
		if uint32(v2) < uint32(i32(17)) {
			v5 = i64(0)
		l10:
			{
				if v2 == 0 {
					goto l7
				}
				t10 := int32(m.memory[uint32(v1)])
				m.fn199(v3+i32(8), t10, i32(10))
				t11 := int32(load32(m.memory[int64(uint32(v3))+8:]))
				if t11 != i32(1) {
					goto l4
				}
				v1 = v1 + i32(1)
				v2 = v2 + i32(-1)
				t12 := int32(load32(m.memory[int64(uint32(v3))+12:]))
				v5 = v5*i64(10) + int64(uint32(t12))
				goto l10
			}
		}
		v5 = i64(0)
	l9:
		{
			if v2 == 0 {
				goto l7
			}
			m.fn1853(v3+i32(16), v5, i64(0), i64(10), i64(0))
			t5 := int32(m.memory[uint32(v1)])
			m.fn199(v3+i32(40), t5, i32(10))
			t6 := int32(load32(m.memory[int64(uint32(v3))+40:]))
			v4 = t6
			t7 := int64(load64(m.memory[int64(uint32(v3))+24:]))
			if t7 != i64(0) {
				v2 = i32(1)
				if v4&i32(1) == 0 {
					goto l4
				}
				m.memory[int64(uint32(v0))+1] = byte(i32(2))
				goto l3
			}
			if v4&i32(1) == 0 {
				goto l4
			}
			v1 = v1 + i32(1)
			v2 = v2 + i32(-1)
			t8 := int64(load64(m.memory[int64(uint32(v3))+16:]))
			v6 = t8
			t9 := int32(load32(m.memory[int64(uint32(v3))+44:]))
			v5 = v6 + int64(uint32(t9))
			if uint64(v5) >= uint64(v6) {
				goto l9
			}
		}
		m.memory[int64(uint32(v0))+1] = byte(i32(2))
		v2 = i32(1)
		goto l3
	}
l4:
	v2 = i32(1)
	m.memory[int64(uint32(v0))+1] = byte(i32(1))
	goto l3
l7:
	store64(m.memory[int64(uint32(v0))+8:], uint64(v5))
	v2 = i32(0)
l3:
	m.memory[uint32(v0)] = byte(v2)
	m.g0 = v3 + i32(48)
}
func (m *Module) fn1191(v0 int32) {
	var v1 int32
	var v2, v3 int64
	t0 := m.g0
	v1 = t0 - i32(16)
	m.g0 = v1
	m.fn34(v1)
	t1 := int64(load64(m.memory[uint32(v1):]))
	v2 = t1
	t2 := int64(load64(m.memory[int64(uint32(v1))+8:]))
	v3 = t2
	t3 := int64(load64(m.memory[int64(uint32(i32(0)))+1286264:]))
	store64(m.memory[int64(uint32(v0))+8:], uint64(t3))
	t4 := int64(load64(m.memory[int64(uint32(i32(0)))+1286256:]))
	store64(m.memory[uint32(v0):], uint64(t4))
	store64(m.memory[int64(uint32(v0))+24:], uint64(v3))
	store64(m.memory[int64(uint32(v0))+16:], uint64(v2))
	m.g0 = v1 + i32(16)
}
func (m *Module) fn1192(v0, v1, v2 int32) int32 {
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
			t6 := m.fn1320(t4, t5, v3, v1, v2)
			v0 = t6
			p7 := i32(0)
			if v0 != 0 {
				p7 = v0 + i32(-12)
			}
			return p7
		}
		return i32(0)
	}
}
func (m *Module) fn1193(v0 int32) {
	var v1, v2, v3, v4 int32
	var v5 int64
	var v6 int32
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
					m.fn771(v6 + (i32(0)-int32(uint32(int64(bits.TrailingZeros64(uint64(v5))))>>3))*i32(480) + i32(-472))
					v4 = v4 + i32(-1)
					v5 = (v5 + i64(-1)) & v5
					goto l4
				}
				v6 = v6 + i32(-3840)
				t5 := int64(load64(m.memory[uint32(v0):]))
				v5 = (t5 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
				v0 = v0 + i32(8)
				goto l3
			}
		}
	l1:
		m.fn39(v1+i32(4), i32(480), i32(8), v2+i32(1))
		t6 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		t7 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t8 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		m.fn40(v3-t6, t7, t8)
	}
l0:
	m.g0 = v1 + i32(16)
}
func (m *Module) fn1194(v0, v1, v2 int32) {
	var v3, v4, v5, v6, v7, v8 int32
	var v9, v10 int64
	var v11, v12 int32
	t0 := m.g0
	v3 = t0 - i32(80)
	m.g0 = v3
	{
		{
			t1 := m.fn886(v1, v2, i32(1072544), i32(60), i32(1077730), i32(6))
			v4 = t1
			if v4 != 0 {
				goto l0
			}
			v4 = i32(0)
			goto l1
		}
	l0:
		t2 := int32(load32(m.memory[uint32(v4+i32(16)):]))
		t3 := int32(load32(m.memory[uint32(v4+i32(20)):]))
		m.fn1046(v3+i32(32), t2, t3, i32(1072544), i32(60), i32(1073156), i32(3))
		t4 := int32(load32(m.memory[int64(uint32(v3))+36:]))
		v5 = t4
		t5 := int32(load32(m.memory[int64(uint32(v3))+32:]))
		v4 = t5
	}
l1:
	{
		{
			p6 := i32(1079292)
			if v4 != 0 {
				p6 = v4
			}
			v6 = p6
			t8 := v6
			p7 := i32(6)
			if v4 != 0 {
				p7 = v5
			}
			v4 = p7
			t9 := m.fn15(t8, v4, i32(1074851), i32(4))
			v5 = t9
			if v5 == 0 {
				goto l2
			}
			v6 = i32(255)
			goto l3
		}
	l2:
		{
			t10 := m.fn15(v6, v4, i32(1079292), i32(6))
			if t10 == 0 {
				goto l4
			}
			v6 = i32(0)
			goto l3
		}
	l4:
		{
			t11 := m.fn15(v6, v4, i32(1086243), i32(11))
			if t11 == 0 {
				goto l5
			}
			v6 = i32(2)
			goto l3
		}
	l5:
		{
			t12 := m.fn15(v6, v4, i32(1086254), i32(11))
			if t12 == 0 {
				goto l6
			}
			v6 = i32(3)
			goto l3
		}
	l6:
		{
			t13 := m.fn15(v6, v4, i32(1086265), i32(10))
			if t13 == 0 {
				goto l7
			}
			v6 = i32(4)
			goto l3
		}
	l7:
		t14 := m.fn15(v6, v4, i32(1086275), i32(10))
		p15 := i32(1)
		if t14 != 0 {
			p15 = i32(5)
		}
		v6 = p15
	}
l3:
	v7 = i32(0)
	{
		{
			t16 := m.fn886(v1, v2, i32(1072544), i32(60), i32(1077061), i32(5))
			v4 = t16
			if v4 != 0 {
				goto l8
			}
			v4 = i32(0)
			goto l9
		}
	l8:
		t17 := int32(load32(m.memory[uint32(v4+i32(16)):]))
		t18 := int32(load32(m.memory[uint32(v4+i32(20)):]))
		m.fn1046(v3+i32(24), t17, t18, i32(1072544), i32(60), i32(1073156), i32(3))
		t19 := int32(load32(m.memory[int64(uint32(v3))+28:]))
		v8 = t19
		t20 := int32(load32(m.memory[int64(uint32(v3))+24:]))
		v4 = t20
	}
l9:
	m.fn1196(v3+i32(40), v4, v8)
	t21 := int64(load64(m.memory[int64(uint32(v3))+48:]))
	v9 = t21
	t22 := int64(load64(m.memory[int64(uint32(v3))+40:]))
	v10 = t22
	{
		t23 := m.fn886(v1, v2, i32(1072544), i32(60), i32(1086285), i32(10))
		v4 = t23
		if v4 == 0 {
			goto l11
		}
		t24 := int32(load32(m.memory[uint32(v4+i32(16)):]))
		t25 := int32(load32(m.memory[uint32(v4+i32(20)):]))
		m.fn1046(v3+i32(16), t24, t25, i32(1072544), i32(60), i32(1073156), i32(3))
		t26 := int32(load32(m.memory[int64(uint32(v3))+16:]))
		v4 = t26
		if v4 == 0 {
			goto l11
		}
		t27 := int32(load32(m.memory[int64(uint32(v3))+20:]))
		m.fn1071(v3+i32(40), v4, t27)
		t28 := int32(m.memory[int64(uint32(v3))+40])
		v7 = (t28 ^ i32(-1)) & i32(1)
		t29 := int32(load32(m.memory[int64(uint32(v3))+44:]))
		v11 = t29
		goto l11
	}
l11:
	v8 = i32(4)
	v12 = i32(0)
	{
		t30 := v5
		var p31 int32
		if v6 == 0 {
			p31 = 1
		}
		if t30|p31 != 0 {
			goto l12
		}
		t32 := m.fn886(v1, v2, i32(1072544), i32(60), i32(1086295), i32(7))
		v4 = t32
		if v4 == 0 {
			goto l12
		}
		t33 := int32(load32(m.memory[uint32(v4+i32(16)):]))
		t34 := int32(load32(m.memory[uint32(v4+i32(20)):]))
		m.fn1046(v3+i32(8), t33, t34, i32(1072544), i32(60), i32(1073156), i32(3))
		t35 := int32(load32(m.memory[int64(uint32(v3))+8:]))
		v4 = t35
		if v4 == 0 {
			goto l12
		}
		t36 := int32(load32(m.memory[int64(uint32(v3))+12:]))
		v5 = t36
		store32(m.memory[int64(uint32(v3))+64:], uint32(i32(0)))
		store64(m.memory[int64(uint32(v3))+56:], uint64(i64(0x400000000)))
		store32(m.memory[int64(uint32(v3))+68:], uint32(i32(-2)))
		store32(m.memory[int64(uint32(v3))+72:], uint32(v4))
		store32(m.memory[int64(uint32(v3))+76:], uint32(v4+v5))
	l19:
		{
			{
				t37 := m.fn869(v3 + i32(68))
				v4 = t37
				if v4 == i32(37) {
					goto l13
				}
				if v4 != i32(-1) {
					goto l14
				}
				t38 := int32(load32(m.memory[int64(uint32(v3))+56:]))
				v4 = t38
				if v4 == i32(-1) {
					goto l12
				}
				t39 := int32(load32(m.memory[int64(uint32(v3))+64:]))
				v12 = t39
				t40 := int32(load32(m.memory[int64(uint32(v3))+60:]))
				v8 = t40
				goto l15
			}
		l13:
			t41 := m.fn870(v3 + i32(68))
			v5 = t41
			if v5 == 0 {
				goto l14
			}
			t42 := int32(load32(m.memory[uint32(v5):]))
			m.fn199(v3, t42, i32(10))
			t43 := int32(load32(m.memory[uint32(v3):]))
			if t43 != i32(1) {
				goto l14
			}
			t44 := int32(load32(m.memory[int64(uint32(v3))+4:]))
			v5 = t44 + i32(-1)
			if uint32(v5) < uint32(i32(9)) {
				_ = m.fn869(v3 + i32(68))
				store32(m.memory[int64(uint32(v3))+40:], uint32(i32(-1)))
				m.memory[int64(uint32(v3))+44] = byte(v5)
				m.fn1321(v3+i32(56), v3+i32(40))
				goto l19
			}
		}
	l14:
		{
			t45 := int32(load32(m.memory[int64(uint32(v3))+64:]))
			v5 = t45
			if v5 == 0 {
				goto l17
			}
			t46 := int32(load32(m.memory[int64(uint32(v3))+60:]))
			v5 = t46 + v5*i32(12) + i32(-12)
			if v5 == 0 {
				goto l17
			}
			t47 := int32(load32(m.memory[uint32(v5):]))
			if t47 != i32(-1) {
				m.fn74(v5, v4)
				goto l19
			}
		}
	l17:
		m.fn1072(v3+i32(40), v4)
		m.fn1321(v3+i32(56), v3+i32(40))
		goto l19
	}
l12:
	v4 = i32(0)
l15:
	t49 := m.fn1318(v1, v2, i32(1086302), i32(5))
	v5 = t49
	m.memory[int64(uint32(v0))+32] = byte(v6)
	store32(m.memory[int64(uint32(v0))+16:], uint32(v12))
	store32(m.memory[int64(uint32(v0))+12:], uint32(v8))
	store32(m.memory[int64(uint32(v0))+8:], uint32(v4))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v11))
	store32(m.memory[uint32(v0):], uint32(v7))
	t51 := v0
	p50 := i64(1)
	if int32(v10) != 0 {
		p50 = v9
	}
	store64(m.memory[int64(uint32(t51))+24:], uint64(p50))
	t52 := v0
	var p53 int32
	if v5&i32(255) == i32(1) {
		p53 = 1
	}
	m.memory[int64(uint32(t52))+20] = byte(p53)
	m.g0 = v3 + i32(80)
}
func (m *Module) fn1195(v0, v1, v2 int32) {
	var v3 int32
	t0 := m.g0
	v3 = t0 - i32(16)
	m.g0 = v3
	{
		{
			t1 := m.fn886(v1, v2, i32(1072544), i32(60), i32(1074872), i32(6))
			v2 = t1
			if v2 != 0 {
				goto l0
			}
			v2 = i32(0)
			goto l1
		}
	l0:
		t2 := int32(load32(m.memory[uint32(v2+i32(16)):]))
		t3 := int32(load32(m.memory[uint32(v2+i32(20)):]))
		m.fn1046(v3+i32(8), t2, t3, i32(1072544), i32(60), i32(1073156), i32(3))
		t4 := int32(load32(m.memory[int64(uint32(v3))+12:]))
		v1 = t4
		t5 := int32(load32(m.memory[int64(uint32(v3))+8:]))
		v2 = t5
	}
l1:
	m.fn1041(v0, v2, v1)
	m.g0 = v3 + i32(16)
}
func (m *Module) fn1196(v0, v1, v2 int32) {
	var v3 int32
	var v4 int64
	t0 := m.g0
	v3 = t0 - i32(16)
	m.g0 = v3
	v4 = i64(0)
	{
		if v1 == 0 {
			goto l0
		}
		m.fn1322(v3, v1, v2)
		t1 := int32(m.memory[uint32(v3)])
		if t1 == i32(1) {
			goto l0
		}
		t2 := int64(load64(m.memory[int64(uint32(v3))+8:]))
		t3 := v0
		v4 = t2
		p4 := i64(0)
		if v4 > i64(0) {
			p4 = v4
		}
		v4 = p4
		p5 := i64(0x7fffffff)
		if v4 < i64(0x7fffffff) {
			p5 = v4
		}
		store64(m.memory[int64(uint32(t3))+8:], uint64(p5))
		v4 = i64(1)
	}
l0:
	store64(m.memory[uint32(v0):], uint64(v4))
	m.g0 = v3 + i32(16)
}
func (m *Module) fn1197(v0, v1, v2 int32) {
	var v3, v4, v5, v6 int32
	var v7, v8, v9, v10 int64
	var v11, v12, v13, v14, v15, v16, v17, v18, v19, v20, v21, v22, v23, v24, v25, v26, v27, v28, v29, v30, v31, v32 int32
	t0 := m.g0
	v3 = t0 - i32(1664)
	m.g0 = v3
	{
		{
			t1 := m.fn159(v1, v2, i32(1080188), i32(5))
			if t1 != 0 {
				goto l0
			}
			m.fn1200(v0+i32(4), i32(1082060), i32(15))
			store32(m.memory[uint32(v0):], uint32(i32(-1)))
			goto l1
		}
	l0:
		store32(m.memory[int64(uint32(v3))+184:], uint32(i32(0)))
		store32(m.memory[int64(uint32(v3))+180:], uint32(v2))
		store32(m.memory[int64(uint32(v3))+176:], uint32(v1))
		{
		l4:
			{
				m.fn1371(v3+i32(640), v3+i32(176))
				{
					t2 := int32(load32(m.memory[int64(uint32(v3))+640:]))
					v4 = t2
					if v4 != i32(-1) {
						goto l2
					}
					v5 = i32(1153692)
					goto l3
				}
			l2:
				if uint32(v4) > uint32(i32(1)) {
					goto l4
				}
				t3 := int32(load32(m.memory[int64(uint32(v3))+644:]))
				v6 = t3
				t4 := int32(load32(m.memory[int64(uint32(v3))+648:]))
				t5 := int32(load32(m.memory[int64(uint32(v3))+652:]))
				t6 := m.fn15(t4, t5, i32(1082053), i32(7))
				if t6 == 0 {
					goto l4
				}
				if v4&i32(1) == 0 {
					goto l4
				}
			}
			v4 = i32(1148716)
			{
				p7 := i32(0)
				if v6 > i32(0) {
					p7 = v6
				}
				v6 = p7
				switch v6 + i32(-1250) {
				case 2:
					goto l7
				default:
					switch v6 + i32(-932) {
					case 1, 2, 3:
						goto l7
					default:
						switch v6 + i32(-949) {
						case 0:
							v4 = i32(1148752)
							goto l7
						case 1:
							v4 = i32(1148748)
							goto l7
						default:
							goto l20
						}
					case 0:
						v4 = i32(1148756)
						goto l7
					case 4:
						v4 = i32(1148744)
						goto l7
					}
				case 0:
					v4 = i32(1148708)
					goto l7
				case 1:
					v4 = i32(1148712)
					goto l7
				case 3:
					v4 = i32(1148720)
					goto l7
				case 4:
					v4 = i32(1148724)
					goto l7
				case 5:
					v4 = i32(1148728)
					goto l7
				case 6:
					v4 = i32(1148732)
					goto l7
				case 7:
					v4 = i32(1148736)
					goto l7
				case 8:
					v4 = i32(1148740)
					goto l7
				}
			}
		l20:
			if v6 != i32(874) {
				goto l7
			}
			v4 = i32(1148704)
		l7:
			t8 := int32(load32(m.memory[uint32(v4):]))
			v5 = t8
		}
	l3:
		m.fn34(v3 + i32(640))
		t9 := int64(load64(m.memory[int64(uint32(v3))+640:]))
		v7 = t9
		t10 := int64(load64(m.memory[int64(uint32(v3))+648:]))
		v8 = t10
		m.fn34(v3 + i32(640))
		t11 := int64(load64(m.memory[int64(uint32(v3))+640:]))
		v9 = t11
		t12 := int64(load64(m.memory[int64(uint32(v3))+648:]))
		v10 = t12
		m.fn34(v3 + i32(640))
		store64(m.memory[int64(uint32(v3))+1128:], uint64(v8))
		store64(m.memory[int64(uint32(v3))+1120:], uint64(v7))
		store64(m.memory[int64(uint32(v3))+1160:], uint64(v10))
		store64(m.memory[int64(uint32(v3))+1152:], uint64(v9))
		t13 := int64(load64(m.memory[int64(uint32(i32(0)))+1286256:]))
		t14 := v3
		v7 = t13
		store64(m.memory[int64(uint32(t14))+1104:], uint64(v7))
		t15 := int64(load64(m.memory[int64(uint32(i32(0)))+1286264:]))
		t16 := v3
		v8 = t15
		store64(m.memory[int64(uint32(t16))+1112:], uint64(v8))
		store64(m.memory[int64(uint32(v3))+1136:], uint64(v7))
		store64(m.memory[int64(uint32(v3))+1144:], uint64(v8))
		store64(m.memory[int64(uint32(v3))+1168:], uint64(v7))
		store64(m.memory[int64(uint32(v3))+1176:], uint64(v8))
		t17 := int64(load64(m.memory[int64(uint32(v3))+648:]))
		store64(m.memory[int64(uint32(v3))+1192:], uint64(t17))
		t18 := int64(load64(m.memory[int64(uint32(v3))+640:]))
		store64(m.memory[int64(uint32(v3))+1184:], uint64(t18))
		m.fn1372(v3+i32(640), v1, v2, i32(1078256), i32(7))
		t19 := int32(load32(m.memory[int64(uint32(v3))+644:]))
		v11 = t19
		t20 := int32(load32(m.memory[int64(uint32(v3))+648:]))
		v12 = v11 + t20<<3
		t21 := int32(load32(m.memory[int64(uint32(v3))+640:]))
		v13 = t21
		v14 = v3 + i32(1168)
		v15 = v3 + i32(1152)
		v16 = v3 + i32(1136)
		v17 = v3 + i32(1120)
		v18 = v11
	l43:
		{
			if v18 == v12 {
				goto l21
			}
			t22 := int64(load64(m.memory[uint32(v18):]))
			v7 = t22
			v19 = i32(0)
			store32(m.memory[int64(uint32(v3))+184:], uint32(i32(0)))
			store64(m.memory[int64(uint32(v3))+176:], uint64(v7))
		l23:
			{
				m.fn1371(v3+i32(640), v3+i32(176))
				t23 := int32(load32(m.memory[int64(uint32(v3))+640:]))
				v4 = t23
				if v4 == i32(-1) {
					v18 = v18 + i32(8)
					goto l43
				}
				if uint32(v4) > uint32(i32(1)) {
					goto l23
				}
				t24 := int32(load32(m.memory[int64(uint32(v3))+644:]))
				v6 = t24
				{
					t25 := int32(load32(m.memory[int64(uint32(v3))+648:]))
					v20 = t25
					t26 := int32(load32(m.memory[int64(uint32(v3))+652:]))
					t27 := v20
					v21 = t26
					t28 := m.fn15(t27, v21, i32(1079211), i32(1))
					if t28 == 0 {
						t29 := m.fn15(v20, v21, i32(1085932), i32(8))
						if t29 == 0 {
							goto l23
						}
						var p30 int32
						if v19 != i32(1) {
							p30 = 1
						}
						v20 = p30
						v19 = i32(0)
						if v20 != 0 {
							goto l23
						}
						v19 = i32(1)
						if v4 != i32(1) {
							goto l23
						}
						switch v6 + i32(-128) {
						default:
							switch v6 + i32(-161) {
							case 0:
								v20 = i32(1153712)
								goto l35
							case 1:
								v20 = i32(1153732)
								goto l35
							case 2:
								v20 = i32(1153812)
								goto l35
							default:
								goto l34
							}
						l34:
							v20 = v5
							if uint32(v6) < uint32(i32(2)) {
								goto l35
							}
							if v6 == i32(177) {
								v20 = i32(1153752)
								goto l35
							}
							if v6 == i32(186) {
								v20 = i32(1153792)
								goto l35
							}
							if v6 == i32(204) {
								v20 = i32(1153672)
								goto l35
							}
							if v6 == i32(222) {
								v20 = i32(1153600)
								goto l35
							}
							if v6 == i32(238) {
								goto l40
							}
							fallthrough
						case 2, 3, 4, 5, 7:
							p31 := v5
							if uint32(v6+i32(-178)) < uint32(i32(3)) {
								p31 = i32(1153772)
							}
							v20 = p31
							goto l35
						case 0:
							v20 = i32(1153144)
							goto l35
						case 1:
							v20 = i32(1149016)
							goto l35
						case 6:
							v20 = i32(1154788)
							goto l35
						case 8:
							v20 = i32(1154812)
							goto l35
						}
					l40:
						v20 = i32(1153652)
					l35:
						store32(m.memory[int64(uint32(v3))+1520:], uint32(v22))
						t32 := int64(load64(m.memory[int64(uint32(v3))+1120:]))
						t33 := int64(load64(m.memory[int64(uint32(v3))+1128:]))
						t34 := m.fn66(t32, t33, v22)
						v7 = t34
						store32(m.memory[int64(uint32(v3))+1200:], uint32(v3+i32(1520)))
						{
							t35 := int32(load32(m.memory[int64(uint32(v3))+1112:]))
							if t35 != 0 {
								goto l41
							}
							_ = m.fn719(v3+i32(1104), v17)
						}
					l41:
						store32(m.memory[int64(uint32(v3))+1236:], uint32(v3+i32(1104)))
						store32(m.memory[int64(uint32(v3))+1232:], uint32(v3+i32(1200)))
						t37 := int32(load32(m.memory[int64(uint32(v3))+1104:]))
						t38 := int32(load32(m.memory[int64(uint32(v3))+1108:]))
						m.fn69(v3+i32(72), t37, t38, v7, v3+i32(1232), i32(163))
						t39 := int32(load32(m.memory[int64(uint32(v3))+76:]))
						v4 = t39
						t40 := int32(load32(m.memory[int64(uint32(v3))+1104:]))
						v6 = t40
						{
							t41 := int32(load32(m.memory[int64(uint32(v3))+72:]))
							if t41 != i32(1) {
								goto l42
							}
							v21 = v6 + v4
							t42 := int32(m.memory[uint32(v21)])
							v23 = t42
							t43 := v21
							v24 = int32(uint32(int32(v7)) >> 25)
							m.memory[uint32(t43)] = byte(v24)
							t44 := int32(load32(m.memory[int64(uint32(v3))+1108:]))
							m.memory[uint32(v6+t44&(v4+i32(-8))+i32(8))] = byte(v24)
							store32(m.memory[uint32(v6-v4<<3+i32(-8)):], uint32(v22))
							t45 := int32(load32(m.memory[int64(uint32(v3))+1116:]))
							store32(m.memory[int64(uint32(v3))+1116:], uint32(t45+i32(1)))
							t46 := int32(load32(m.memory[int64(uint32(v3))+1112:]))
							store32(m.memory[int64(uint32(v3))+1112:], uint32(t46-v23&i32(1)))
						}
					l42:
						store32(m.memory[uint32(v6+(i32(0)-v4)<<3+i32(-4)):], uint32(v20))
						goto l23
					}
					v22 = v6
					v19 = v4
					goto l23
				}
			}
		}
	l21:
		m.fn1217(v11, v13)
		m.fn1372(v3+i32(640), v1, v2, i32(1074279), i32(10))
		t47 := int32(load32(m.memory[int64(uint32(v3))+644:]))
		v25 = t47
		t48 := int32(load32(m.memory[int64(uint32(v3))+648:]))
		v26 = v25 + t48<<3
		t49 := int32(load32(m.memory[int64(uint32(v3))+640:]))
		v27 = t49
		v28 = v3 + i32(640) + i32(16)
		v29 = v3 + i32(176) + i32(16)
		v30 = v25
	l77:
		{
			if v30 == v26 {
				goto l44
			}
			t50 := int64(load64(m.memory[uint32(v30):]))
			v7 = t50
			v6 = i32(0)
			store32(m.memory[int64(uint32(v3))+1624:], uint32(i32(0)))
			store64(m.memory[int64(uint32(v3))+1616:], uint64(v7))
			m.fn22(v3+i32(640), i32(3))
			t51 := int64(load64(m.memory[int64(uint32(i32(0)))+1286256:]))
			store64(m.memory[int64(uint32(v3))+176:], uint64(t51))
			t52 := int64(load64(m.memory[int64(uint32(i32(0)))+1286264:]))
			store64(m.memory[int64(uint32(v3))+184:], uint64(t52))
			t53 := int64(load64(m.memory[int64(uint32(v3))+648:]))
			store64(m.memory[int64(uint32(v3))+200:], uint64(t53))
			t54 := int64(load64(m.memory[int64(uint32(v3))+640:]))
			store64(m.memory[int64(uint32(v3))+192:], uint64(t54))
			store32(m.memory[int64(uint32(v3))+1528:], uint32(i32(0)))
			store64(m.memory[int64(uint32(v3))+1520:], uint64(i64(0x100000000)))
			v21 = i32(2)
		l49:
			m.fn1371(v3+i32(1232), v3+i32(1616))
			{
				t55 := int32(load32(m.memory[int64(uint32(v3))+1232:]))
				v4 = t55
				if v4 == i32(-1) {
					t105 := int32(load32(m.memory[int64(uint32(v3))+176:]))
					v22 = t105
					v20 = v22 + i32(8)
					t106 := int64(load64(m.memory[uint32(v22):]))
					v7 = (t106 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
					t107 := int32(load32(m.memory[int64(uint32(v3))+188:]))
					v31 = t107
				l84:
					{
						if v31 == 0 {
							t148 := int32(load32(m.memory[int64(uint32(v3))+1520:]))
							t149 := int32(load32(m.memory[int64(uint32(v3))+1524:]))
							m.fn16(t148, t149)
							{
								t150 := int32(load32(m.memory[int64(uint32(v3))+180:]))
								v4 = t150
								if v4 == 0 {
									goto l76
								}
								t151 := int32(load32(m.memory[int64(uint32(v3))+176:]))
								v6 = t151
								m.fn39(v3+i32(640), i32(20), i32(8), v4+i32(1))
								t152 := int32(load32(m.memory[int64(uint32(v3))+648:]))
								t153 := int32(load32(m.memory[int64(uint32(v3))+640:]))
								t154 := int32(load32(m.memory[int64(uint32(v3))+644:]))
								m.fn40(v6-t152, t153, t154)
							}
						l76:
							v30 = v30 + i32(8)
							goto l77
						}
					l66:
						{
							if v7 != i64(0) {
								v6 = i32(0)
								t109 := int32(load32(m.memory[uint32(v22+(i32(0)-int32(uint32(int64(bits.TrailingZeros64(uint64(v7))))>>3))*i32(20)+i32(-20)):]))
								v17 = t109
								store32(m.memory[int64(uint32(v3))+1208:], uint32(i32(0)))
								store64(m.memory[int64(uint32(v3))+1200:], uint64(i64(0x400000000)))
								m.fn22(v3+i32(1232), i32(3))
								t110 := int64(load64(m.memory[int64(uint32(i32(0)))+1286256:]))
								store64(m.memory[int64(uint32(v3))+640:], uint64(t110))
								t111 := int64(load64(m.memory[int64(uint32(i32(0)))+1286264:]))
								store64(m.memory[int64(uint32(v3))+648:], uint64(t111))
								t112 := int64(load64(m.memory[int64(uint32(v3))+1240:]))
								store64(m.memory[int64(uint32(v3))+664:], uint64(t112))
								t113 := int64(load64(m.memory[int64(uint32(v3))+1232:]))
								store64(m.memory[int64(uint32(v3))+656:], uint64(t113))
								v31 = v31 + i32(-1)
								v7 = (v7 + i64(-1)) & v7
								v21 = i32(1)
								v11 = i32(4)
								v4 = v17
							l74:
								{
									if v21 != i32(1) {
										goto l67
									}
									store32(m.memory[int64(uint32(v3))+1652:], uint32(v4))
									t114 := int64(load64(m.memory[int64(uint32(v3))+656:]))
									t115 := int64(load64(m.memory[int64(uint32(v3))+664:]))
									t116 := m.fn66(t114, t115, v4)
									v8 = t116
									store32(m.memory[int64(uint32(v3))+1600:], uint32(v3+i32(1652)))
									{
										t117 := int32(load32(m.memory[int64(uint32(v3))+648:]))
										if t117 != 0 {
											goto l68
										}
										_ = m.fn725(v3+i32(640), v28)
									}
								l68:
									store32(m.memory[int64(uint32(v3))+1236:], uint32(v3+i32(640)))
									store32(m.memory[int64(uint32(v3))+1232:], uint32(v3+i32(1600)))
									t119 := int32(load32(m.memory[int64(uint32(v3))+640:]))
									t120 := int32(load32(m.memory[int64(uint32(v3))+644:]))
									m.fn69(v3+i32(48), t119, t120, v8, v3+i32(1232), i32(4))
									t121 := int32(load32(m.memory[int64(uint32(v3))+48:]))
									if t121&i32(1) == 0 {
										goto l67
									}
									t122 := int32(load32(m.memory[int64(uint32(v3))+640:]))
									v21 = t122
									t123 := int32(load32(m.memory[int64(uint32(v3))+52:]))
									t124 := v21
									v19 = t123
									v12 = t124 + v19
									t125 := int32(m.memory[uint32(v12)])
									v23 = t125
									t126 := v12
									v24 = int32(uint32(int32(v8)) >> 25)
									m.memory[uint32(t126)] = byte(v24)
									t127 := int32(load32(m.memory[int64(uint32(v3))+644:]))
									m.memory[uint32(v21+t127&(v19+i32(-8))+i32(8))] = byte(v24)
									store32(m.memory[uint32(v21-v19<<2+i32(-4)):], uint32(v4))
									t128 := int32(load32(m.memory[int64(uint32(v3))+652:]))
									store32(m.memory[int64(uint32(v3))+652:], uint32(t128+i32(1)))
									t129 := int32(load32(m.memory[int64(uint32(v3))+648:]))
									store32(m.memory[int64(uint32(v3))+648:], uint32(t129-v23&i32(1)))
									t130 := int32(load32(m.memory[int64(uint32(v3))+188:]))
									if t130 == 0 {
										goto l67
									}
									t131 := int64(load64(m.memory[int64(uint32(v3))+192:]))
									t132 := int64(load64(m.memory[int64(uint32(v3))+200:]))
									t133 := m.fn66(t131, t132, v4)
									v8 = t133
									t134 := int32(load32(m.memory[int64(uint32(v3))+180:]))
									v23 = t134
									v19 = v23 & int32(v8)
									v9 = int64(uint64(v8)>>25) & i64(127) * i64(72340172838076673)
									v13 = i32(0)
									t135 := int32(load32(m.memory[int64(uint32(v3))+176:]))
									v21 = t135
								l75:
									{
										t136 := int64(load64(m.memory[uint32(v21+v19):]))
										v10 = t136
										v8 = v10 ^ v9
										v8 = (v8 ^ i64(-1)) & (v8 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
										{
										l71:
											{
												var p137 int32
												if v8 == 0 {
													p137 = 1
												}
												v12 = p137
												if v12 != 0 {
													goto l69
												}
												t138 := v4
												t139 := v21
												v24 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v8))))>>3) + v19) & v23
												t140 := int32(load32(m.memory[uint32(t139+(i32(0)-v24)*i32(20)+i32(-20)):]))
												if t138 == t140 {
													goto l70
												}
												v8 = (v8 + i64(-1)) & v8
												goto l71
											}
										l69:
											if v10&(v10<<1)&i64(-0x7f7f7f7f7f7f7f80) == 0 {
												t147 := v19
												v13 = v13 + i32(8)
												v19 = (t147 + v13) & v23
												goto l75
											}
										l70:
											if v12 != 0 {
												goto l67
											}
											v21 = v21 + (i32(0)-v24)*i32(20)
											p141 := v21
											if v12 != 0 {
												p141 = i32(0)
											}
											v4 = p141 + i32(-16)
											{
												t142 := int32(load32(m.memory[int64(uint32(v3))+1200:]))
												if v6 != t142 {
													goto l73
												}
												m.fn618(v3 + i32(1200))
												t143 := int32(load32(m.memory[int64(uint32(v3))+1204:]))
												v11 = t143
											}
										l73:
											store32(m.memory[uint32(v11+v6<<2):], uint32(v4))
											t144 := v3
											v6 = v6 + i32(1)
											store32(m.memory[int64(uint32(t144))+1208:], uint32(v6))
											t145 := int32(load32(m.memory[uint32(v21+i32(-4)):]))
											v4 = t145
											t146 := int32(load32(m.memory[uint32(v21+i32(-8)):]))
											v21 = t146
											goto l74
										}
									}
								}
							}
							v22 = v22 + i32(-160)
							t108 := int64(load64(m.memory[uint32(v20):]))
							v7 = (t108 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
							v20 = v20 + i32(8)
							goto l66
						}
					l67:
						v21 = i32(2)
						v6 = v6 << 2
						v12 = i32(0)
						v19 = i32(33686018)
						t155 := int32(load32(m.memory[int64(uint32(v3))+1204:]))
						v24 = t155
						v4 = i32(0)
					l79:
						{
							if v6 == 0 {
								store32(m.memory[int64(uint32(v3))+1652:], uint32(v17))
								t167 := int64(load64(m.memory[int64(uint32(v3))+1152:]))
								t168 := int64(load64(m.memory[int64(uint32(v3))+1160:]))
								t169 := m.fn66(t167, t168, v17)
								v8 = t169
								store32(m.memory[int64(uint32(v3))+1600:], uint32(v3+i32(1652)))
								{
									t170 := int32(load32(m.memory[int64(uint32(v3))+1144:]))
									if t170 != 0 {
										goto l80
									}
									_ = m.fn717(v16, v15)
								}
							l80:
								store32(m.memory[int64(uint32(v3))+1236:], uint32(v16))
								store32(m.memory[int64(uint32(v3))+1232:], uint32(v3+i32(1600)))
								t172 := int32(load32(m.memory[int64(uint32(v3))+1136:]))
								t173 := int32(load32(m.memory[int64(uint32(v3))+1140:]))
								m.fn69(v3+i32(40), t172, t173, v8, v3+i32(1232), i32(165))
								t174 := int32(load32(m.memory[int64(uint32(v3))+44:]))
								v6 = t174
								t175 := int32(load32(m.memory[int64(uint32(v3))+1136:]))
								v12 = t175
								{
									t176 := int32(load32(m.memory[int64(uint32(v3))+40:]))
									if t176 != i32(1) {
										goto l81
									}
									v23 = v12 + v6
									t177 := int32(m.memory[uint32(v23)])
									v11 = t177
									t178 := v23
									v13 = int32(uint32(int32(v8)) >> 25)
									m.memory[uint32(t178)] = byte(v13)
									t179 := int32(load32(m.memory[int64(uint32(v3))+1140:]))
									m.memory[uint32(v12+t179&(v6+i32(-8))+i32(8))] = byte(v13)
									v6 = v12 + (i32(0)-v6)*i32(12)
									store32(m.memory[uint32(v6+i32(-12)):], uint32(v17))
									t180 := int32(load32(m.memory[int64(uint32(v3))+1148:]))
									store32(m.memory[int64(uint32(v3))+1148:], uint32(t180+i32(1)))
									t181 := int32(load32(m.memory[int64(uint32(v3))+1144:]))
									store32(m.memory[int64(uint32(v3))+1144:], uint32(t181-v11&i32(1)))
									goto l82
								}
							l81:
								v6 = v12 + (i32(0)-v6)*i32(12)
							l82:
								m.memory[uint32(v6+i32(-2))] = byte(v21)
								store32(m.memory[uint32(v6+i32(-6)):], uint32(v19))
								m.memory[uint32(v6+i32(-7))] = byte(v18)
								m.memory[uint32(v6+i32(-8))] = byte(v4)
								{
									t182 := int32(load32(m.memory[int64(uint32(v3))+644:]))
									v4 = t182
									if v4 == 0 {
										goto l83
									}
									t183 := int32(load32(m.memory[int64(uint32(v3))+640:]))
									v6 = t183
									m.fn39(v3+i32(1232), i32(4), i32(8), v4+i32(1))
									t184 := int32(load32(m.memory[int64(uint32(v3))+1240:]))
									t185 := int32(load32(m.memory[int64(uint32(v3))+1232:]))
									t186 := int32(load32(m.memory[int64(uint32(v3))+1236:]))
									m.fn40(v6-t184, t185, t186)
								}
							l83:
								t187 := int32(load32(m.memory[int64(uint32(v3))+1200:]))
								m.fn136(t187, v24, i32(4), i32(4))
								goto l84
							}
							t156 := v19
							v6 = v6 + i32(-4)
							t157 := int32(load32(m.memory[uint32(v6+v24):]))
							v4 = t157
							t158 := int32(load32(m.memory[int64(uint32(v4))+2:]))
							t159 := fn1373(t156, t158)
							v19 = t159
							t160 := int32(m.memory[int64(uint32(v4))+1])
							t161 := int32(m.memory[uint32(v4)])
							t162 := v18
							v23 = t161
							p163 := t162
							if v23 != 0 {
								p163 = t160
							}
							v18 = p163
							t164 := int32(m.memory[int64(uint32(v4))+6])
							t165 := v21
							v4 = t164
							p166 := v4
							if v4 == i32(2) {
								p166 = t165
							}
							v21 = p166
							v12 = v12 | v23
							v4 = v12 & i32(1)
							goto l79
						}
					}
				}
				t56 := int32(load32(m.memory[int64(uint32(v3))+1236:]))
				v20 = t56
				{
					{
						p57 := v4 + i32(-2)
						if uint32(v4) < uint32(i32(2)) {
							p57 = i32(2)
						}
						switch p57 {
						case 2:
							t84 := int32(load32(m.memory[int64(uint32(v3))+1240:]))
							v19 = t84
							t85 := int32(load32(m.memory[int64(uint32(v3))+1244:]))
							t86 := v19
							v22 = t85
							t87 := m.fn15(t86, v22, i32(1079224), i32(1))
							if t87 != 0 {
								v12 = i32(0)
								store32(m.memory[int64(uint32(v3))+1528:], uint32(i32(0)))
								p188 := i32(0)
								if v4&i32(1) != 0 {
									p188 = v20
								}
								v18 = p188
								v23 = i32(33686018)
								v21 = i32(0)
								goto l49
							}
							{
								t88 := m.fn15(v19, v22, i32(1086035), i32(8))
								if t88 != 0 {
									if v21 == i32(2) {
										goto l63
									}
									t94 := v4
									var p95 int32
									if v20 != i32(222) {
										p95 = 1
									}
									v21 = t94 & p95
									v11 = v20
									goto l49
								}
								t89 := m.fn15(v19, v22, i32(1079144), i32(12))
								if t89 != 0 {
									t96 := v12
									var p97 int32
									if v21 != i32(2) {
										p97 = 1
									}
									t98 := p97 & v4
									var p99 int32
									if uint32(v20) < uint32(i32(9)) {
										p99 = 1
									}
									v4 = t98 & p99
									p100 := t96
									if v4 != 0 {
										p100 = i32(1)
									}
									v12 = p100
									p101 := v24
									if v4 != 0 {
										p101 = v20 + i32(1)
									}
									v24 = p101
									goto l49
								}
								t90 := m.fn15(v19, v22, i32(1073720), i32(1))
								if t90 != 0 {
									if v21 == i32(2) {
										goto l63
									}
									t102 := v23 & i32(-256)
									t103 := v4 ^ i32(-1)
									var p104 int32
									if v20 != i32(0) {
										p104 = 1
									}
									v23 = t102 | (t103|p104)&i32(1)
									goto l49
								}
								t91 := m.fn15(v19, v22, i32(1073721), i32(1))
								if t91 == 0 {
									goto l49
								}
								if v21 == i32(2) {
									goto l49
								}
								p92 := i32(0)
								if v20 != 0 {
									p92 = i32(256)
								}
								p93 := i32(256)
								if v4&i32(1) != 0 {
									p93 = p92
								}
								v23 = p93 | v23&i32(-65281)
								goto l49
							}
						default:
							goto l49
						case 0:
							v6 = v6 + i32(1)
							goto l49
						case 1:
							if v6 != i32(1) {
								goto l52
							}
							if v21 == i32(2) {
								goto l53
							}
							t58 := int32(load32(m.memory[int64(uint32(v3))+1524:]))
							t59 := int32(load32(m.memory[int64(uint32(v3))+1528:]))
							m.fn510(v3+i32(640), v5, t58, t59)
							t60 := int32(load32(m.memory[int64(uint32(v3))+640:]))
							v19 = t60
							t61 := int32(load32(m.memory[int64(uint32(v3))+644:]))
							t62 := v3 + i32(64)
							v22 = t61
							t63 := int32(load32(m.memory[int64(uint32(v3))+648:]))
							m.fn856(t62, v22, t63, i32(59))
							t64 := int32(load32(m.memory[int64(uint32(v3))+64:]))
							t65 := int32(load32(m.memory[int64(uint32(v3))+68:]))
							t66 := m.fn1208(t64, t65)
							v4 = t66
							store32(m.memory[int64(uint32(v3))+1600:], uint32(v18))
							t67 := int64(load64(m.memory[int64(uint32(v3))+192:]))
							t68 := int64(load64(m.memory[int64(uint32(v3))+200:]))
							t69 := m.fn66(t67, t68, v18)
							v7 = t69
							store32(m.memory[int64(uint32(v3))+1200:], uint32(v3+i32(1600)))
							{
								t70 := int32(load32(m.memory[int64(uint32(v3))+184:]))
								if t70 != 0 {
									goto l54
								}
								_ = m.fn723(v3+i32(176), v29)
							}
						l54:
							v13 = v4 & i32(255)
							store32(m.memory[int64(uint32(v3))+644:], uint32(v3+i32(176)))
							store32(m.memory[int64(uint32(v3))+640:], uint32(v3+i32(1200)))
							t72 := int32(load32(m.memory[int64(uint32(v3))+176:]))
							t73 := int32(load32(m.memory[int64(uint32(v3))+180:]))
							m.fn69(v3+i32(56), t72, t73, v7, v3+i32(640), i32(164))
							t74 := int32(load32(m.memory[int64(uint32(v3))+60:]))
							v4 = t74
							t75 := int32(load32(m.memory[int64(uint32(v3))+176:]))
							v20 = t75
							t76 := int32(load32(m.memory[int64(uint32(v3))+56:]))
							if t76 != i32(1) {
								goto l55
							}
							v17 = v20 + v4
							t77 := int32(m.memory[uint32(v17)])
							v31 = t77
							t78 := v17
							v32 = int32(uint32(int32(v7)) >> 25)
							m.memory[uint32(t78)] = byte(v32)
							t79 := int32(load32(m.memory[int64(uint32(v3))+180:]))
							m.memory[uint32(v20+t79&(v4+i32(-8))+i32(8))] = byte(v32)
							v4 = v20 + (i32(0)-v4)*i32(20)
							store32(m.memory[uint32(v4+i32(-20)):], uint32(v18))
							t80 := int32(load32(m.memory[int64(uint32(v3))+188:]))
							store32(m.memory[int64(uint32(v3))+188:], uint32(t80+i32(1)))
							t81 := int32(load32(m.memory[int64(uint32(v3))+184:]))
							store32(m.memory[int64(uint32(v3))+184:], uint32(t81-v31&i32(1)))
							goto l56
						case 4:
							if v6 != i32(1) {
								goto l49
							}
							if v21 != i32(2) {
								goto l57
							}
							goto l58
						case 5:
							if v6 != i32(1) {
								goto l49
							}
							if v21 == i32(2) {
								goto l58
							}
						}
					}
				l57:
					m.fn145(v3+i32(1520), v20)
					v6 = i32(1)
					goto l49
				l55:
					v4 = v20 + (i32(0)-v4)*i32(20)
				l56:
					store32(m.memory[uint32(v4+i32(-4)):], uint32(v11))
					store32(m.memory[uint32(v4+i32(-8)):], uint32(v21))
					m.memory[uint32(v4+i32(-10))] = byte(v13)
					store32(m.memory[uint32(v4+i32(-14)):], uint32(v23))
					m.memory[uint32(v4+i32(-15))] = byte(v24)
					m.memory[uint32(v4+i32(-16))] = byte(v12)
					m.fn134(v19, v22)
				l53:
					v21 = i32(2)
				l52:
					store32(m.memory[int64(uint32(v3))+1528:], uint32(i32(0)))
					t82 := v6
					var p83 int32
					if v6 != i32(0) {
						p83 = 1
					}
					v6 = t82 - p83
					goto l49
				}
			}
		l58:
			v6 = i32(1)
		l63:
			v21 = i32(2)
			goto l49
		}
	l44:
		m.fn1217(v25, v27)
		m.fn22(v3+i32(640), i32(3))
		t189 := int64(load64(m.memory[int64(uint32(i32(0)))+1286256:]))
		store64(m.memory[int64(uint32(v3))+1200:], uint64(t189))
		t190 := int64(load64(m.memory[int64(uint32(i32(0)))+1286264:]))
		store64(m.memory[int64(uint32(v3))+1208:], uint64(t190))
		t191 := int64(load64(m.memory[int64(uint32(v3))+648:]))
		store64(m.memory[int64(uint32(v3))+1224:], uint64(t191))
		t192 := int64(load64(m.memory[int64(uint32(v3))+640:]))
		store64(m.memory[int64(uint32(v3))+1216:], uint64(t192))
		m.fn1372(v3+i32(640), v1, v2, i32(1078413), i32(9))
		t193 := int32(load32(m.memory[int64(uint32(v3))+644:]))
		v30 = t193
		t194 := int32(load32(m.memory[int64(uint32(v3))+648:]))
		v16 = v30 + t194<<3
		t195 := int32(load32(m.memory[int64(uint32(v3))+640:]))
		v32 = t195
		v17 = v30
	l111:
		{
			if v17 == v16 {
				goto l85
			}
			t196 := int64(load64(m.memory[uint32(v17):]))
			v7 = t196
			v6 = i32(0)
			store32(m.memory[int64(uint32(v3))+1608:], uint32(i32(0)))
			store64(m.memory[int64(uint32(v3))+1600:], uint64(v7))
		l87:
			if v6 == i32(288) {
				memory_copy(m.memory, uint32(v3+i32(1232)), uint32(v3+i32(640)), uint32(i32(288)))
				store64(m.memory[int64(uint32(v3))+1536:], uint64(i64(1)))
				store64(m.memory[int64(uint32(v3))+1528:], uint64(i64(0)))
				store64(m.memory[int64(uint32(v3))+1520:], uint64(i64(0x100000000)))
				store16(m.memory[int64(uint32(v3))+1544:], uint16(i32(512)))
				v11 = i32(0)
				v21 = i32(0)
				v19 = i32(0)
				v12 = i32(0)
				v20 = i32(0)
				v23 = i32(0)
			l92:
				{
					m.fn1371(v3+i32(1616), v3+i32(1600))
					{
						t197 := int32(load32(m.memory[int64(uint32(v3))+1616:]))
						v4 = t197
						if v4 == i32(-1) {
							v17 = v17 + i32(8)
							m.fn759(v3 + i32(1520))
							m.fn764(v3 + i32(1232))
							goto l111
						}
						t198 := int32(load32(m.memory[int64(uint32(v3))+1620:]))
						v6 = t198
						{
							p199 := v4 + i32(-2)
							if uint32(v4) < uint32(i32(2)) {
								p199 = i32(2)
							}
							switch p199 {
							default:
								goto l92
							case 0:
								v20 = v20 + i32(1)
								goto l92
							case 1:
								m.memory[int64(uint32(v3))+1545] = byte(i32(2))
								if v19&v23&i32(1) == 0 {
									goto l94
								}
								v19 = i32(1)
								if v20 != v24+i32(1) {
									goto l94
								}
								{
									if uint32(v12) > uint32(i32(8)) {
										goto l95
									}
									{
										{
											v4 = v3 + i32(1232) + v12<<5
											t200 := int32(m.memory[int64(uint32(v4))+24])
											if uint32((t200+i32(-1))&i32(255)) < uint32(i32(254)) {
												goto l96
											}
											v6 = i32(0)
											store32(m.memory[int64(uint32(v3))+648:], uint32(i32(0)))
											store64(m.memory[int64(uint32(v3))+640:], uint64(i64(0x400000000)))
											goto l97
										}
									l96:
										t201 := int32(load32(m.memory[int64(uint32(v3))+1524:]))
										t202 := int32(load32(m.memory[int64(uint32(v3))+1528:]))
										t203 := int32(load32(m.memory[int64(uint32(v3))+1536:]))
										t204 := int32(load32(m.memory[int64(uint32(v3))+1540:]))
										m.fn1374(v3+i32(640), t201, t202, t203, t204, v5)
										t205 := int32(m.memory[int64(uint32(v3))+1544])
										v6 = t205
									}
								l97:
									m.memory[int64(uint32(v3))+652] = byte(v6)
									m.fn759(v3 + i32(1520))
									m.memory[int64(uint32(v3))+1544] = byte(i32(0))
									store64(m.memory[int64(uint32(v3))+1536:], uint64(i64(1)))
									store64(m.memory[int64(uint32(v3))+1528:], uint64(i64(0)))
									store64(m.memory[int64(uint32(v3))+1520:], uint64(i64(0x100000000)))
									m.fn763(v4)
									t206 := int64(load64(m.memory[int64(uint32(v3))+648:]))
									store64(m.memory[int64(uint32(v4))+8:], uint64(t206))
									t207 := int64(load64(m.memory[int64(uint32(v3))+640:]))
									store64(m.memory[uint32(v4):], uint64(t207))
								}
							l95:
								v12 = v12 + i32(1)
								v19 = i32(0)
								goto l94
							case 2:
								t208 := int32(load32(m.memory[int64(uint32(v3))+1624:]))
								v22 = t208
								t209 := int32(load32(m.memory[int64(uint32(v3))+1628:]))
								t210 := v22
								v18 = t209
								t211 := m.fn15(t210, v18, i32(1081789), i32(4))
								if t211 != 0 {
									v6 = i32(0)
								l104:
									if v6 == i32(288) {
										memory_copy(m.memory, uint32(v3+i32(176)), uint32(v3+i32(640)), uint32(i32(288)))
										m.fn764(v3 + i32(1232))
										memory_copy(m.memory, uint32(v3+i32(1232)), uint32(v3+i32(176)), uint32(i32(288)))
										v21 = i32(1)
										v11 = i32(0)
										v13 = v20
										v12 = i32(0)
										v24 = v20
										v23 = i32(1)
										goto l92
									}
									v4 = v3 + i32(640) + v6
									store64(m.memory[uint32(v4):], uint64(i64(0x400000000)))
									m.memory[uint32(v4+i32(24))] = byte(i32(0))
									store64(m.memory[uint32(v4+i32(16)):], uint64(i64(1)))
									m.memory[uint32(v4+i32(12))] = byte(i32(0))
									store32(m.memory[uint32(v4+i32(8)):], uint32(i32(0)))
									v6 = v6 + i32(32)
									goto l104
								}
								t212 := m.fn15(v22, v18, i32(1085940), i32(6))
								if t212 != 0 {
									if v21 != 0 {
										v21 = i32(1)
										v31 = v6
										v11 = v4
										goto l92
									}
									v21 = i32(0)
									goto l92
								}
								t213 := m.fn15(v22, v18, i32(1085946), i32(9))
								if t213 != 0 {
									m.fn759(v3 + i32(1520))
									store16(m.memory[int64(uint32(v3))+1544:], uint16(i32(512)))
									store64(m.memory[int64(uint32(v3))+1536:], uint64(i64(1)))
									store64(m.memory[int64(uint32(v3))+1528:], uint64(i64(0)))
									store64(m.memory[int64(uint32(v3))+1520:], uint64(i64(0x100000000)))
									v19 = i32(1)
									goto l92
								}
								{
									t214 := m.fn15(v22, v18, i32(1085955), i32(8))
									if t214 != 0 {
										goto l101
									}
									t215 := m.fn15(v22, v18, i32(1085963), i32(9))
									if t215 == 0 {
										t221 := m.fn15(v22, v18, i32(1085972), i32(12))
										if t221 != 0 {
											t228 := v19
											var p229 int32
											if uint32(v12) < uint32(i32(9)) {
												p229 = 1
											}
											if t228&p229 == 0 {
												goto l92
											}
											v19 = i32(1)
											if v4 != i32(1) {
												goto l92
											}
											t231 := v3 + i32(1232) + v12<<5
											p230 := i32(0)
											if v6 > i32(0) {
												p230 = v6
											}
											store64(m.memory[int64(uint32(t231))+16:], uint64(uint32(p230)))
											goto l92
										}
										t222 := m.fn15(v22, v18, i32(1085984), i32(9))
										if t222 != 0 {
											if v19&i32(1) != 0 {
												v19 = i32(1)
												m.memory[int64(uint32(v3))+1545] = byte(i32(1))
												goto l92
											}
											goto l110
										}
										t223 := m.fn15(v22, v18, i32(1085993), i32(12))
										if t223 != 0 {
											if v19&i32(1) == 0 {
												goto l110
											}
											m.memory[int64(uint32(v3))+1545] = byte(i32(0))
											v19 = i32(1)
											goto l92
										}
										t224 := m.fn15(v22, v18, i32(1086005), i32(10))
										if t224&v19 == 0 {
											goto l92
										}
										v19 = i32(1)
										t225 := v3
										t226 := v4 ^ i32(-1)
										var p227 int32
										if v6 != i32(0) {
											p227 = 1
										}
										m.memory[int64(uint32(t225))+1544] = byte((t226 | p227) & i32(1))
										goto l92
									}
								}
							l101:
								t216 := v19
								var p217 int32
								if uint32(v12) < uint32(i32(9)) {
									p217 = 1
								}
								if t216&p217 == 0 {
									goto l92
								}
								v19 = i32(1)
								t219 := v3 + i32(1232) + v12<<5
								p218 := i32(0)
								if v4&i32(1) != 0 {
									p218 = v6
								}
								t220 := fn1375(p218)
								m.memory[int64(uint32(t219))+24] = byte(t220)
								goto l92
							case 4, 5:
								m.fn1376(v3+i32(1520), v6)
								goto l92
							}
						}
					}
				l110:
					v19 = i32(0)
					goto l92
				l94:
					if v21 == i32(1) {
						goto l112
					}
					v21 = i32(0)
					goto l113
				l112:
					if v13 == v20 {
						goto l114
					}
					v21 = i32(1)
					goto l113
				l114:
					if v11&i32(1) == 0 {
						goto l115
					}
					m.fn1137(v3+i32(640), v3+i32(1232))
					m.fn1124(v3+i32(176), v3+i32(1200), v31, v3+i32(640))
					m.fn1377(v3 + i32(176))
				l115:
					v6 = i32(0)
				l117:
					if v6 == i32(288) {
						goto l116
					}
					v4 = v3 + i32(640) + v6
					store64(m.memory[uint32(v4):], uint64(i64(0x400000000)))
					m.memory[uint32(v4+i32(24))] = byte(i32(0))
					store64(m.memory[uint32(v4+i32(16)):], uint64(i64(1)))
					m.memory[uint32(v4+i32(12))] = byte(i32(0))
					store32(m.memory[uint32(v4+i32(8)):], uint32(i32(0)))
					v6 = v6 + i32(32)
					goto l117
				l116:
					memory_copy(m.memory, uint32(v3+i32(176)), uint32(v3+i32(640)), uint32(i32(288)))
					m.fn764(v3 + i32(1232))
					memory_copy(m.memory, uint32(v3+i32(1232)), uint32(v3+i32(176)), uint32(i32(288)))
					v11 = i32(0)
					v21 = i32(0)
					v12 = i32(0)
					v23 = i32(0)
				l113:
					t232 := v20
					var p233 int32
					if v20 != i32(0) {
						p233 = 1
					}
					v20 = t232 - p233
					goto l92
				}
			}
			v4 = v3 + i32(640) + v6
			store64(m.memory[uint32(v4):], uint64(i64(0x400000000)))
			m.memory[uint32(v4+i32(24))] = byte(i32(0))
			store64(m.memory[uint32(v4+i32(16)):], uint64(i64(1)))
			m.memory[uint32(v4+i32(12))] = byte(i32(0))
			store32(m.memory[uint32(v4+i32(8)):], uint32(i32(0)))
			v6 = v6 + i32(32)
			goto l87
		}
	l85:
		m.fn1217(v30, v32)
		m.fn1372(v3+i32(640), v1, v2, i32(1078422), i32(17))
		t234 := int32(load32(m.memory[int64(uint32(v3))+644:]))
		v16 = t234
		t235 := int32(load32(m.memory[int64(uint32(v3))+648:]))
		v13 = v16 + t235<<3
		t236 := int32(load32(m.memory[int64(uint32(v3))+640:]))
		v30 = t236
		v17 = v3 + i32(1520) + i32(12)
		v31 = v3 + i32(640) + i32(12)
		v11 = v16
	l146:
		{
			if v11 == v13 {
				goto l118
			}
			t237 := int64(load64(m.memory[uint32(v11):]))
			v7 = t237
			store32(m.memory[int64(uint32(v3))+1552:], uint32(v5))
			v4 = i32(0)
			store32(m.memory[int64(uint32(v3))+1564:], uint32(i32(0)))
			store64(m.memory[int64(uint32(v3))+1556:], uint64(v7))
			store32(m.memory[int64(uint32(v3))+1568:], uint32(i32(0)))
			store32(m.memory[int64(uint32(v3))+1576:], uint32(i32(0)))
		l120:
			if v4 == i32(144) {
				m.memory[int64(uint32(v3))+1592] = byte(i32(-2))
				store64(m.memory[int64(uint32(v3))+1584:], uint64(i64(-72340172838076674)))
				v4 = i32(0)
			l122:
				if v4 == i32(252) {
					memory_copy(m.memory, uint32(v3+i32(176)), uint32(v3+i32(640)), uint32(i32(252)))
					store64(m.memory[int64(uint32(v3))+656:], uint64(i64(1)))
					store64(m.memory[int64(uint32(v3))+648:], uint64(i64(0)))
					store64(m.memory[int64(uint32(v3))+640:], uint64(i64(0x100000000)))
					store16(m.memory[int64(uint32(v3))+664:], uint16(i32(512)))
					store32(m.memory[int64(uint32(v3))+1660:], uint32(v14))
					store32(m.memory[int64(uint32(v3))+1656:], uint32(v3+i32(1552)))
					store32(m.memory[int64(uint32(v3))+1652:], uint32(v3+i32(1200)))
					v20 = i32(0)
					v19 = i32(0)
					v21 = i32(0)
					v12 = i32(0)
				l127:
					{
						m.fn1371(v3+i32(1600), v3+i32(1556))
						{
							t238 := int32(load32(m.memory[int64(uint32(v3))+1600:]))
							v4 = t238
							if v4 == i32(-1) {
								v11 = v11 + i32(8)
								m.fn1379(v3+i32(1652), v3+i32(1576), v3+i32(1568), v3+i32(1232), v3+i32(1584), v3+i32(176))
								m.fn759(v3 + i32(640))
								m.fn1380(v3 + i32(176))
								goto l146
							}
							t239 := int32(load32(m.memory[int64(uint32(v3))+1604:]))
							v6 = t239
							{
								p240 := v4 + i32(-2)
								if uint32(v4) < uint32(i32(2)) {
									p240 = i32(2)
								}
								switch p240 {
								default:
									goto l127
								case 0:
									v20 = v20 + i32(1)
									goto l127
								case 1:
									m.memory[int64(uint32(v3))+665] = byte(i32(2))
									if v19 == i32(1) {
										v19 = i32(1)
										if v24 != v20 {
											goto l130
										}
										if uint32(v12) >= uint32(i32(9)) {
											goto l131
										}
										t241 := int32(load32(m.memory[int64(uint32(v3))+648:]))
										if t241 == 0 {
											goto l131
										}
										t242 := int64(load64(m.memory[uint32(v31):]))
										store64(m.memory[uint32(v17):], uint64(t242))
										t243 := int32(load32(m.memory[int64(uint32(v31))+8:]))
										store32(m.memory[int64(uint32(v17))+8:], uint32(t243))
										t244 := int32(load32(m.memory[int64(uint32(v3))+648:]))
										store32(m.memory[int64(uint32(v3))+1528:], uint32(t244))
										t245 := int64(load64(m.memory[int64(uint32(v3))+640:]))
										v7 = t245
										store64(m.memory[int64(uint32(v3))+640:], uint64(i64(0x100000000)))
										store32(m.memory[int64(uint32(v3))+648:], uint32(i32(0)))
										store64(m.memory[int64(uint32(v3))+1616:], uint64(v7))
										store64(m.memory[int64(uint32(v3))+652:], uint64(i64(0x100000000)))
										store32(m.memory[int64(uint32(v3))+660:], uint32(i32(0)))
										t246 := int64(load64(m.memory[int64(uint32(v3))+1528:]))
										store64(m.memory[int64(uint32(v3))+1624:], uint64(t246))
										t247 := int64(load64(m.memory[int64(uint32(v3))+1536:]))
										store64(m.memory[int64(uint32(v3))+1632:], uint64(t247))
										t248 := int32(m.memory[int64(uint32(v3))+664])
										v6 = t248
										v4 = v3 + i32(176) + v12*i32(28)
										m.fn1378(v4)
										m.memory[int64(uint32(v4))+24] = byte(v6)
										t249 := int64(load64(m.memory[int64(uint32(v3))+1632:]))
										store64(m.memory[int64(uint32(v4))+16:], uint64(t249))
										t250 := int64(load64(m.memory[int64(uint32(v3))+1624:]))
										store64(m.memory[int64(uint32(v4))+8:], uint64(t250))
										t251 := int64(load64(m.memory[int64(uint32(v3))+1616:]))
										store64(m.memory[uint32(v4):], uint64(t251))
										goto l131
									}
									v19 = i32(0)
									goto l130
								case 2:
									t252 := int32(load32(m.memory[int64(uint32(v3))+1608:]))
									v22 = t252
									t253 := int32(load32(m.memory[int64(uint32(v3))+1612:]))
									t254 := v22
									v18 = t253
									t255 := m.fn15(t254, v18, i32(1086015), i32(12))
									if t255 != 0 {
										m.fn1379(v3+i32(1652), v3+i32(1576), v3+i32(1568), v3+i32(1232), v3+i32(1584), v3+i32(176))
										v12 = i32(0)
										v21 = i32(1)
										v23 = v20
										goto l127
									}
									{
										{
											t256 := m.fn15(v22, v18, i32(1085940), i32(6))
											if t256 != 0 {
												if v21 != 0 {
													store32(m.memory[int64(uint32(v3))+1572:], uint32(v6))
													store32(m.memory[int64(uint32(v3))+1568:], uint32(v4))
													v21 = i32(1)
													goto l127
												}
												goto l143
											}
											t257 := m.fn15(v22, v18, i32(0x107774), i32(2))
											if t257 != 0 {
												if v21 == 0 {
													goto l143
												}
												store32(m.memory[int64(uint32(v3))+1580:], uint32(v6))
												store32(m.memory[int64(uint32(v3))+1576:], uint32(v4))
												v21 = i32(1)
												goto l127
											}
											t258 := m.fn15(v22, v18, i32(1086027), i32(8))
											if t258 != 0 {
												m.fn759(v3 + i32(640))
												store16(m.memory[int64(uint32(v3))+664:], uint16(i32(512)))
												store64(m.memory[int64(uint32(v3))+656:], uint64(i64(1)))
												store64(m.memory[int64(uint32(v3))+648:], uint64(i64(0)))
												store64(m.memory[int64(uint32(v3))+640:], uint64(i64(0x100000000)))
												v19 = i32(1)
												v24 = v20
												goto l127
											}
											t259 := m.fn15(v22, v18, i32(1085972), i32(12))
											if t259 != 0 {
												if v19 == 0 {
													goto l127
												}
												if uint32(v12) > uint32(i32(8)) {
													goto l127
												}
												if v4&i32(1) == 0 {
													goto l127
												}
												v4 = v3 + i32(1232) + v12<<4
												store64(m.memory[uint32(v4):], uint64(i64(1)))
												t269 := v4
												p268 := i32(0)
												if v6 > i32(0) {
													p268 = v6
												}
												store64(m.memory[int64(uint32(t269))+8:], uint64(uint32(p268)))
												v19 = i32(1)
												goto l127
											}
											t260 := m.fn15(v22, v18, i32(1085955), i32(8))
											if t260 != 0 {
												if v19 == 0 {
													goto l127
												}
												if uint32(v12) >= uint32(i32(9)) {
													goto l127
												}
												goto l144
											}
											t261 := m.fn15(v22, v18, i32(1085963), i32(9))
											if t261 != 0 {
												goto l138
											}
											t262 := m.fn15(v22, v18, i32(1085984), i32(9))
											if t262 != 0 {
												if v19 != 0 {
													v19 = i32(1)
													m.memory[int64(uint32(v3))+665] = byte(i32(1))
													goto l127
												}
												goto l141
											}
											t263 := m.fn15(v22, v18, i32(1085993), i32(12))
											if t263 != 0 {
												if v19 == 0 {
													goto l141
												}
												m.memory[int64(uint32(v3))+665] = byte(i32(0))
												v19 = i32(1)
												goto l127
											}
											t264 := m.fn15(v22, v18, i32(1086005), i32(10))
											if t264 == 0 {
												goto l127
											}
											if v19 == 0 {
												goto l141
											}
											v19 = i32(1)
											t265 := v3
											t266 := v4 ^ i32(-1)
											var p267 int32
											if v6 != i32(0) {
												p267 = 1
											}
											m.memory[int64(uint32(t265))+664] = byte((t266 | p267) & i32(1))
											goto l127
										}
									l138:
										if v19 == 0 {
											goto l127
										}
										if uint32(v12) >= uint32(i32(9)) {
											goto l127
										}
									l144:
										t271 := v3 + i32(1584) + v12
										p270 := i32(0)
										if v4&i32(1) != 0 {
											p270 = v6
										}
										t272 := fn1375(p270)
										m.memory[uint32(t271)] = byte(t272)
										goto l127
									}
								case 4, 5:
									m.fn1376(v3+i32(640), v6)
									goto l127
								}
							}
						}
					l141:
						v19 = i32(0)
						goto l127
					l143:
						v21 = i32(0)
						goto l127
					l131:
						m.fn759(v3 + i32(640))
						v19 = i32(0)
						m.memory[int64(uint32(v3))+664] = byte(i32(0))
						store64(m.memory[int64(uint32(v3))+656:], uint64(i64(1)))
						store64(m.memory[int64(uint32(v3))+648:], uint64(i64(0)))
						store64(m.memory[int64(uint32(v3))+640:], uint64(i64(0x100000000)))
						v12 = v12 + i32(1)
					l130:
						if v21 == i32(1) {
							goto l147
						}
						v21 = i32(0)
						goto l148
					l147:
						v21 = i32(1)
						if v23 != v20 {
							goto l148
						}
						m.fn1379(v3+i32(1652), v3+i32(1576), v3+i32(1568), v3+i32(1232), v3+i32(1584), v3+i32(176))
						v21 = i32(0)
						v12 = i32(0)
					l148:
						t273 := v20
						var p274 int32
						if v20 != i32(0) {
							p274 = 1
						}
						v20 = t273 - p274
						goto l127
					}
				}
				store32(m.memory[uint32(v3+i32(640)+v4):], uint32(i32(-1)))
				v4 = v4 + i32(28)
				goto l122
			}
			store64(m.memory[uint32(v3+i32(1232)+v4):], uint64(i64(0)))
			v4 = v4 + i32(16)
			goto l120
		}
	l118:
		m.fn1217(v16, v30)
		memory_copy(m.memory, uint32(v3+i32(80)), uint32(v3+i32(1104)), uint32(i32(96)))
		m.fn1381(v3 + i32(1200))
		m.fn34(v3 + i32(640))
		t275 := int64(load64(m.memory[int64(uint32(v3))+640:]))
		v7 = t275
		t276 := int64(load64(m.memory[int64(uint32(v3))+648:]))
		v8 = t276
		t277 := m.fn113(i32(4), i32(12))
		v4 = t277
		store32(m.memory[int64(uint32(v4))+8:], uint32(i32(0)))
		store64(m.memory[uint32(v4):], uint64(i64(0x800000000)))
		t278 := m.fn113(i32(4), i32(16))
		v6 = t278
		store32(m.memory[uint32(v6):], uint32(i32(0)))
		m.fn1225(v3 + i32(640))
		store16(m.memory[int64(uint32(v3))+634:], uint16(i32(0xff00)))
		store32(m.memory[int64(uint32(v3))+630:], uint32(i32(33685504)))
		m.memory[int64(uint32(v3))+628] = byte(i32(0))
		store32(m.memory[int64(uint32(v3))+624:], uint32(i32(0)))
		store64(m.memory[int64(uint32(v3))+616:], uint64(i64(0x100000001)))
		store64(m.memory[int64(uint32(v3))+608:], uint64(i64(0)))
		store32(m.memory[int64(uint32(v3))+600:], uint32(i32(0)))
		store32(m.memory[int64(uint32(v3))+592:], uint32(i32(0)))
		store32(m.memory[int64(uint32(v3))+552:], uint32(i32(0)))
		store64(m.memory[int64(uint32(v3))+544:], uint64(i64(0x400000000)))
		store32(m.memory[int64(uint32(v3))+504:], uint32(i32(0)))
		store32(m.memory[int64(uint32(v3))+500:], uint32(v2))
		store32(m.memory[int64(uint32(v3))+496:], uint32(v1))
		memory_copy(m.memory, uint32(v3+i32(176)), uint32(v3+i32(80)), uint32(i32(96)))
		m.memory[int64(uint32(v3))+636] = byte(i32(0))
		store32(m.memory[int64(uint32(v3))+492:], uint32(i32(0)))
		store32(m.memory[int64(uint32(v3))+488:], uint32(v5))
		store16(m.memory[int64(uint32(v3))+484:], uint16(i32(0)))
		store32(m.memory[int64(uint32(v3))+480:], uint32(i32(0)))
		store64(m.memory[int64(uint32(v3))+472:], uint64(i64(0x100000000)))
		store32(m.memory[int64(uint32(v3))+588:], uint32(i32(0)))
		store64(m.memory[int64(uint32(v3))+580:], uint64(i64(0x800000000)))
		store64(m.memory[int64(uint32(v3))+572:], uint64(i64(8)))
		store64(m.memory[int64(uint32(v3))+564:], uint64(i64(0)))
		store64(m.memory[int64(uint32(v3))+556:], uint64(i64(0x400000000)))
		store32(m.memory[int64(uint32(v3))+352:], uint32(i32(0)))
		store64(m.memory[int64(uint32(v3))+528:], uint64(i64(0x100000001)))
		store64(m.memory[int64(uint32(v3))+516:], uint64(i64(0x100000000)))
		store64(m.memory[int64(uint32(v3))+508:], uint64(i64(0x800000000)))
		store32(m.memory[int64(uint32(v3))+540:], uint32(i32(1)))
		store32(m.memory[int64(uint32(v3))+536:], uint32(v6))
		store32(m.memory[int64(uint32(v3))+524:], uint32(v4))
		store64(m.memory[int64(uint32(v3))+296:], uint64(v8))
		store64(m.memory[int64(uint32(v3))+288:], uint64(v7))
		t279 := int64(load64(m.memory[int64(uint32(i32(0)))+1286256:]))
		store64(m.memory[int64(uint32(v3))+272:], uint64(t279))
		t280 := int64(load64(m.memory[int64(uint32(i32(0)))+1286264:]))
		store64(m.memory[int64(uint32(v3))+280:], uint64(t280))
		store32(m.memory[int64(uint32(v3))+428:], uint32(i32(-1)))
		store32(m.memory[int64(uint32(v3))+416:], uint32(i32(-1)))
		store64(m.memory[int64(uint32(v3))+408:], uint64(i64(1)))
		store64(m.memory[int64(uint32(v3))+400:], uint64(i64(0)))
		store64(m.memory[int64(uint32(v3))+392:], uint64(i64(0x400000000)))
		store64(m.memory[int64(uint32(v3))+384:], uint64(i64(4)))
		store64(m.memory[int64(uint32(v3))+376:], uint64(i64(0)))
		store64(m.memory[int64(uint32(v3))+368:], uint64(i64(0x400000000)))
		v2 = v3 + i32(304)
		memory_copy(m.memory, uint32(v2), uint32(v3+i32(640)), uint32(i32(48)))
		v14 = v3 + i32(1232) + i32(8)
		v17 = v3 + i32(640) + i32(4)
		v30 = v3 + i32(596)
		v5 = v3 + i32(432)
		v31 = v3 + i32(440)
		v32 = v3 + i32(428)
		v24 = v3 + i32(392)
		v28 = v3 + i32(380)
		v25 = v3 + i32(368)
		v27 = v3 + i32(508)
		v6 = v3 + i32(556)
		v26 = v3 + i32(472)
		v13 = v3 + i32(592)
		v16 = v3 + i32(544)
		v11 = v3 + i32(496)
	l157:
		{
			m.fn1371(v3+i32(1616), v11)
			{
				{
					{
						t281 := int32(load32(m.memory[int64(uint32(v3))+1616:]))
						v4 = t281
						if v4 == i32(-1) {
							{
								t526 := int32(load32(m.memory[int64(uint32(v3))+552:]))
								if t526 == 0 {
									goto l245
								}
								m.memory[int64(uint32(v3))+636] = byte(i32(1))
							}
						l245:
							m.fn1382(v3 + i32(176))
							m.fn1383(v3+i32(1104), v3+i32(176))
							t527 := int32(load32(m.memory[int64(uint32(v3))+1104:]))
							v4 = t527
							if v4 != i32(-1) {
								goto l212
							}
							memory_copy(m.memory, uint32(v3+i32(640)), uint32(v3+i32(176)), uint32(i32(464)))
							t528 := int32(load32(m.memory[int64(uint32(v3))+980:]))
							v4 = t528
							m.memory[int64(uint32(v3))+1112] = byte(i32(0))
							store32(m.memory[int64(uint32(v3))+1108:], uint32(v4))
							store32(m.memory[int64(uint32(v3))+1104:], uint32(i32(1)))
							v6 = v3 + i32(972)
							{
								{
								l248:
									{
										m.fn1391(v3+i32(16), v3+i32(1104))
										t529 := int32(load32(m.memory[int64(uint32(v3))+16:]))
										if t529 != i32(1) {
											goto l246
										}
										{
											t530 := int32(load32(m.memory[int64(uint32(v3))+20:]))
											t531 := v6
											v4 = t530
											t532 := m.fn1392(t531, v4)
											if t532 != 0 {
												goto l247
											}
											v20 = v4 + i32(-1)
											t533 := int32(load32(m.memory[int64(uint32(v3))+980:]))
											if uint32(v20) >= uint32(t533) {
												goto l248
											}
											t534 := int32(load32(m.memory[int64(uint32(v3))+976:]))
											t535 := int32(load32(m.memory[int64(uint32(t534+v20<<6))+8:]))
											if t535 == 0 {
												goto l248
											}
										}
									l247:
										m.fn1387(v3+i32(1232), v3+i32(640), v4)
										t536 := int32(load32(m.memory[int64(uint32(v3))+1232:]))
										v4 = t536
										if v4 == i32(-1) {
											goto l248
										}
									}
									t537 := int32(load32(m.memory[int64(uint32(v3))+1252:]))
									store32(m.memory[int64(uint32(v0))+24:], uint32(t537))
									t538 := int64(load64(m.memory[int64(uint32(v3))+1244:]))
									store64(m.memory[int64(uint32(v0))+16:], uint64(t538))
									t539 := int64(load64(m.memory[int64(uint32(v3))+1236:]))
									store64(m.memory[int64(uint32(v0))+8:], uint64(t539))
									store32(m.memory[uint32(v0):], uint32(i32(-1)))
									store32(m.memory[int64(uint32(v0))+4:], uint32(v4))
									goto l249
								}
							l246:
								t540 := int32(load32(m.memory[int64(uint32(v3))+980:]))
								v4 = t540
								m.memory[int64(uint32(v3))+1112] = byte(i32(0))
								store32(m.memory[int64(uint32(v3))+1108:], uint32(v4))
								store32(m.memory[int64(uint32(v3))+1104:], uint32(i32(2)))
								{
								l251:
									{
										m.fn1391(v3+i32(8), v3+i32(1104))
										t541 := int32(load32(m.memory[int64(uint32(v3))+8:]))
										if t541 != i32(1) {
											goto l250
										}
										t542 := int32(load32(m.memory[int64(uint32(v3))+12:]))
										t543 := v3 + i32(1232)
										t544 := v6
										v4 = t542
										m.fn1393(t543, t544, v4, v4+i32(-1))
										t545 := int32(load32(m.memory[int64(uint32(v3))+1232:]))
										v4 = t545
										if v4 == i32(-1) {
											goto l251
										}
									}
									t546 := int32(load32(m.memory[int64(uint32(v3))+1252:]))
									store32(m.memory[int64(uint32(v0))+24:], uint32(t546))
									t547 := int64(load64(m.memory[int64(uint32(v3))+1244:]))
									store64(m.memory[int64(uint32(v0))+16:], uint64(t547))
									t548 := int64(load64(m.memory[int64(uint32(v3))+1236:]))
									store64(m.memory[int64(uint32(v0))+8:], uint64(t548))
									store32(m.memory[uint32(v0):], uint32(i32(-1)))
									store32(m.memory[int64(uint32(v0))+4:], uint32(v4))
									goto l249
								}
							l250:
								m.fn1394(v3+i32(1232), v3+i32(640))
								t549 := int32(load32(m.memory[int64(uint32(v3))+1232:]))
								v4 = t549
								if v4 == i32(-1) {
									m.fn1396(v3 + i32(640))
									t553 := int32(load32(m.memory[int64(uint32(v3))+1040:]))
									store32(m.memory[int64(uint32(v3))+1240:], uint32(t553))
									t554 := int64(load64(m.memory[int64(uint32(v3))+1032:]))
									store64(m.memory[int64(uint32(v3))+1232:], uint64(t554))
									t555 := int64(load64(m.memory[int64(uint32(v3))+856:]))
									store64(m.memory[int64(uint32(v3))+1244:], uint64(t555))
									t556 := int32(load32(m.memory[int64(uint32(v3))+864:]))
									store32(m.memory[int64(uint32(v3))+1252:], uint32(t556))
									t557 := int64(load64(m.memory[int64(uint32(v3))+804:]))
									store64(m.memory[int64(uint32(v3))+1256:], uint64(t557))
									t558 := int32(load32(m.memory[int64(uint32(v3))+812:]))
									store32(m.memory[int64(uint32(v3))+1264:], uint32(t558))
									memory_copy(m.memory, uint32(v0), uint32(v3+i32(1232)), uint32(i32(36)))
									t559 := int32(load32(m.memory[int64(uint32(v3))+1008:]))
									t560 := int32(load32(m.memory[int64(uint32(v3))+1012:]))
									m.fn1397(t559, t560)
									m.fn1398(v3 + i32(640))
									t561 := int32(load32(m.memory[int64(uint32(v3))+936:]))
									t562 := int32(load32(m.memory[int64(uint32(v3))+940:]))
									m.fn16(t561, t562)
									m.fn894(v3 + i32(1020))
									m.fn1302(v3 + i32(1044))
									m.fn1332(v3 + i32(816))
									t563 := int32(load32(m.memory[int64(uint32(v3))+736:]))
									t564 := int32(load32(m.memory[int64(uint32(v3))+740:]))
									m.fn1399(t563, t564)
									m.fn1400(v6)
									m.fn448(v3 + i32(832))
									t565 := int32(load32(m.memory[int64(uint32(v3))+844:]))
									t566 := int32(load32(m.memory[int64(uint32(v3))+848:]))
									m.fn911(t565, t566)
									t567 := int32(load32(m.memory[int64(uint32(v3))+868:]))
									t568 := int32(load32(m.memory[int64(uint32(v3))+872:]))
									m.fn16(t567, t568)
									t569 := int32(load32(m.memory[int64(uint32(v3))+880:]))
									t570 := int32(load32(m.memory[int64(uint32(v3))+884:]))
									m.fn134(t569, t570)
									m.fn1388(v3 + i32(892))
									m.fn57(v3 + i32(768))
									goto l1
								}
								t550 := int32(load32(m.memory[int64(uint32(v3))+1252:]))
								store32(m.memory[int64(uint32(v0))+24:], uint32(t550))
								t551 := int64(load64(m.memory[int64(uint32(v3))+1244:]))
								store64(m.memory[int64(uint32(v0))+16:], uint64(t551))
								t552 := int64(load64(m.memory[int64(uint32(v3))+1236:]))
								store64(m.memory[int64(uint32(v0))+8:], uint64(t552))
								store32(m.memory[uint32(v0):], uint32(i32(-1)))
								store32(m.memory[int64(uint32(v0))+4:], uint32(v4))
							}
						l249:
							m.fn1395(v3 + i32(640))
							goto l1
						}
						t282 := int32(load32(m.memory[int64(uint32(v3))+1624:]))
						v21 = t282
						t283 := int32(load32(m.memory[int64(uint32(v3))+1620:]))
						v20 = t283
						{
							p284 := v4 + i32(-2)
							if uint32(v4) < uint32(i32(2)) {
								p284 = i32(2)
							}
							switch p284 {
							default:
								m.fn1382(v3 + i32(176))
								memory_copy(m.memory, uint32(v3+i32(640)), uint32(v13), uint32(i32(44)))
								{
									t285 := int32(load32(m.memory[int64(uint32(v3))+552:]))
									v4 = t285
									t286 := int32(load32(m.memory[int64(uint32(v3))+544:]))
									if v4 != t286 {
										goto l156
									}
									m.fn1075(v16)
								}
							l156:
								t287 := int32(load32(m.memory[int64(uint32(v3))+548:]))
								memory_copy(m.memory, uint32(t287+v4*i32(44)), uint32(v3+i32(640)), uint32(i32(44)))
								store32(m.memory[int64(uint32(v3))+552:], uint32(v4+i32(1)))
								goto l157
							case 1:
								m.fn1382(v3 + i32(176))
								{
									t288 := int32(load32(m.memory[int64(uint32(v3))+552:]))
									v4 = t288
									if v4 != 0 {
										t289 := v3
										v19 = v4 + i32(-1)
										store32(m.memory[int64(uint32(t289))+552:], uint32(v19))
										t290 := int32(load32(m.memory[int64(uint32(v3))+548:]))
										v4 = t290 + v19*i32(44)
										t291 := int32(load32(m.memory[uint32(v4):]))
										v20 = t291
										if v20 == i32(2) {
											goto l159
										}
										memory_copy(m.memory, uint32(v3+i32(640)), uint32(v4+i32(4)), uint32(i32(40)))
										store32(m.memory[int64(uint32(v3))+592:], uint32(v20))
										memory_copy(m.memory, uint32(v30), uint32(v3+i32(640)), uint32(i32(40)))
										goto l160
									}
									v19 = i32(0)
									goto l159
								}
							case 2:
								{
									{
										t292 := int32(load32(m.memory[int64(uint32(v3))+1628:]))
										t293 := v21
										v19 = t292
										t294 := m.fn15(t293, v19, i32(1079208), i32(1))
										if t294 != 0 {
											{
												t303 := int32(m.memory[int64(uint32(v3))+634])
												if t303 != 0 {
													goto l168
												}
												t304 := int32(m.memory[int64(uint32(v3))+631])
												if t304&i32(1) != 0 {
													goto l157
												}
											}
										l168:
											m.fn1382(v3 + i32(176))
											store32(m.memory[int64(uint32(v3))+492:], uint32(i32(0)))
											if v4 != i32(1) {
												goto l157
											}
											t305 := int32(load32(m.memory[int64(uint32(v3))+616:]))
											v4 = t305
											t306 := int32(load16(m.memory[int64(uint32(v3))+484:]))
											v21 = t306
											store16(m.memory[int64(uint32(v3))+484:], uint16(i32(0)))
											v19 = v20 & i32(64512)
											if v21 != i32(1) {
												goto l169
											}
											if v19 != i32(56320) {
												goto l170
											}
											t307 := int32(load16(m.memory[int64(uint32(v3))+486:]))
											v20 = v20&i32(0xffff) + t307<<10 + i32(-56613888)
											if uint32(v20^i32(55296)+i32(-1114112)) > uint32(i32(-1112065)) {
												goto l171
											}
											goto l172
										}
										t295 := m.fn15(v21, v19, i32(1079209), i32(2))
										if t295 != 0 {
											t309 := v3
											p308 := i32(0)
											if v20 > i32(0) {
												p308 = v20
											}
											p310 := i32(1)
											if v4&i32(1) != 0 {
												p310 = p308
											}
											store32(m.memory[int64(uint32(t309))+616:], uint32(p310))
											goto l157
										}
										t296 := m.fn15(v21, v19, i32(1079211), i32(1))
										if t296 != 0 {
											m.fn1382(v3 + i32(176))
											store32(m.memory[int64(uint32(v3))+596:], uint32(v20))
											store32(m.memory[int64(uint32(v3))+592:], uint32(v4))
											goto l157
										}
										t297 := v4 ^ i32(-1)
										var p298 int32
										if v20 != i32(0) {
											p298 = 1
										}
										v22 = (t297 | p298) & i32(1)
										t299 := m.fn15(v21, v19, i32(1073720), i32(1))
										if t299 != 0 {
											m.fn1382(v3 + i32(176))
											m.memory[int64(uint32(v3))+608] = byte(v22)
											goto l157
										}
										t300 := m.fn15(v21, v19, i32(1073721), i32(1))
										if t300 != 0 {
											m.fn1382(v3 + i32(176))
											m.memory[int64(uint32(v3))+609] = byte(v22)
											goto l157
										}
										{
											t301 := m.fn15(v21, v19, i32(1073722), i32(6))
											if t301 != 0 {
												goto l166
											}
											t302 := m.fn15(v21, v19, i32(1079212), i32(7))
											if t302 == 0 {
												t311 := m.fn15(v21, v19, i32(1079219), i32(5))
												if t311 != 0 {
													m.fn1382(v3 + i32(176))
													store32(m.memory[int64(uint32(v3))+608:], uint32(i32(0)))
													goto l157
												}
												{
													t312 := m.fn15(v21, v19, i32(1079224), i32(1))
													if t312 != 0 {
														if v4 != i32(1) {
															goto l157
														}
														t321 := int32(load32(m.memory[int64(uint32(v3))+220:]))
														if t321 == 0 {
															goto l157
														}
														t322 := int64(load64(m.memory[int64(uint32(v3))+224:]))
														t323 := int64(load64(m.memory[int64(uint32(v3))+232:]))
														t324 := m.fn66(t322, t323, v20)
														v7 = t324
														t325 := int32(load32(m.memory[int64(uint32(v3))+212:]))
														v22 = t325
														v21 = v22 & int32(v7)
														v8 = int64(uint64(v7)>>25) & i64(127) * i64(72340172838076673)
														v18 = i32(0)
														t326 := int32(load32(m.memory[int64(uint32(v3))+208:]))
														v19 = t326
													l182:
														{
															t327 := int64(load64(m.memory[uint32(v19+v21):]))
															v9 = t327
															v7 = v9 ^ v8
															v7 = (v7 ^ i64(-1)) & (v7 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
														l181:
															{
																if v7 == 0 {
																	if !(v9&(v9<<1)&i64(-0x7f7f7f7f7f7f7f80) == 0) {
																		goto l157
																	}
																	t330 := v21
																	v18 = v18 + i32(8)
																	v21 = (t330 + v18) & v22
																	goto l182
																}
																t328 := v20
																v4 = v19 + (i32(0)-(int32(uint32(int64(bits.TrailingZeros64(uint64(v7))))>>3)+v21)&v22)*i32(12)
																t329 := int32(load32(m.memory[uint32(v4+i32(-12)):]))
																if t328 == t329 {
																	goto l180
																}
																v7 = (v7 + i64(-1)) & v7
																goto l181
															}
														l180:
														}
														t331 := int32(m.memory[uint32(v4+i32(-8))])
														v20 = t331
														if v20 == i32(2) {
															goto l157
														}
														t332 := int32(load32(m.memory[uint32(v4+i32(-6)):]))
														v21 = t332
														t333 := int32(m.memory[uint32(v4+i32(-7))])
														v19 = t333
														t334 := int32(m.memory[uint32(v4+i32(-2))])
														v4 = t334
														m.fn1382(v3 + i32(176))
														m.memory[int64(uint32(v3))+632] = byte(v4)
														m.memory[int64(uint32(v3))+629] = byte(v19)
														m.memory[int64(uint32(v3))+628] = byte(v20 & i32(1))
														t335 := int32(load32(m.memory[int64(uint32(v3))+608:]))
														t336 := fn1319(v21, t335)
														t337 := v3
														v4 = t336
														store32(m.memory[int64(uint32(t337))+612:], uint32(v4))
														store32(m.memory[int64(uint32(v3))+608:], uint32(v4))
														goto l157
													}
													{
														t313 := m.fn15(v21, v19, i32(1075793), i32(3))
														if t313 != 0 {
															goto l175
														}
														t314 := m.fn15(v21, v19, i32(1079225), i32(4))
														if t314 == 0 {
															t338 := m.fn15(v21, v19, i32(1079229), i32(4))
															if t338 != 0 {
																m.fn1382(v3 + i32(176))
																store64(m.memory[int64(uint32(v3))+620:], uint64(i64(1)))
																m.memory[int64(uint32(v3))+630] = byte(i32(0))
																m.memory[int64(uint32(v3))+635] = byte(i32(255))
																store32(m.memory[int64(uint32(v3))+600:], uint32(i32(0)))
																m.memory[int64(uint32(v3))+632] = byte(i32(2))
																m.memory[int64(uint32(v3))+628] = byte(i32(0))
																store32(m.memory[int64(uint32(v3))+612:], uint32(i32(0)))
																goto l157
															}
															t339 := m.fn15(v21, v19, i32(1079233), i32(4))
															if t339 != 0 {
																goto l184
															}
															t340 := m.fn15(v21, v19, i32(1079237), i32(3))
															if t340 != 0 {
																goto l184
															}
															t341 := m.fn15(v21, v19, i32(1079240), i32(4))
															if t341 != 0 {
																goto l184
															}
															t342 := m.fn15(v21, v19, i32(1079244), i32(6))
															if t342 != 0 {
																goto l184
															}
															t343 := m.fn15(v21, v19, i32(1077507), i32(3))
															if t343 != 0 {
																m.fn1389(v3+i32(176), i32(32))
																goto l157
															}
															t344 := m.fn15(v21, v19, i32(1079250), i32(6))
															if t344 != 0 {
																m.fn1389(v3+i32(176), i32(8212))
																goto l157
															}
															t345 := m.fn15(v21, v19, i32(1079256), i32(6))
															if t345 != 0 {
																m.fn1389(v3+i32(176), i32(8211))
																goto l157
															}
															t346 := m.fn15(v21, v19, i32(1079262), i32(6))
															if t346 != 0 {
																m.fn1389(v3+i32(176), i32(8216))
																goto l157
															}
															t347 := m.fn15(v21, v19, i32(1079268), i32(6))
															if t347 != 0 {
																m.fn1389(v3+i32(176), i32(8217))
																goto l157
															}
															t348 := m.fn15(v21, v19, i32(1079274), i32(9))
															if t348 != 0 {
																m.fn1389(v3+i32(176), i32(8220))
																goto l157
															}
															t349 := m.fn15(v21, v19, i32(1079283), i32(9))
															if t349 != 0 {
																m.fn1389(v3+i32(176), i32(8221))
																goto l157
															}
															t350 := m.fn15(v21, v19, i32(1079292), i32(6))
															if t350 != 0 {
																m.fn1389(v3+i32(176), i32(8226))
																goto l157
															}
															t351 := m.fn15(v21, v19, i32(1079298), i32(7))
															if t351 != 0 {
																goto l193
															}
															t352 := m.fn15(v21, v19, i32(1079305), i32(7))
															if t352 == 0 {
																goto l194
															}
															goto l193
														}
													}
												l175:
													m.fn1382(v3 + i32(176))
													t315 := int32(m.memory[int64(uint32(v3))+633])
													if t315 != i32(2) {
														store32(m.memory[int64(uint32(v3))+640:], uint32(i32(8)))
														m.fn1340(v6, v3+i32(640))
														goto l157
													}
													t316 := int32(m.memory[int64(uint32(v3))+631])
													if t316 != 0 {
														goto l157
													}
													m.fn1383(v3+i32(640), v3+i32(176))
													t317 := int32(load32(m.memory[int64(uint32(v3))+640:]))
													v4 = t317
													if v4 == i32(-1) {
														goto l157
													}
													t318 := int64(load64(m.memory[int64(uint32(v3))+645:]))
													store64(m.memory[int64(uint32(v3))+1200:], uint64(t318))
													t319 := int64(load64(m.memory[int64(uint32(v3))+653:]))
													store64(m.memory[int64(uint32(v3))+1208:], uint64(t319))
													t320 := int32(load32(m.memory[int64(uint32(v3))+660:]))
													store32(m.memory[int64(uint32(v3))+1215:], uint32(t320))
													goto l178
												}
											}
										}
									l166:
										m.fn1382(v3 + i32(176))
										m.memory[int64(uint32(v3))+610] = byte(v22)
										goto l157
									}
								l169:
									if v19 == i32(55296) {
										goto l195
									}
								l170:
									v20 = int32(uint32(v20)>>15)&i32(65536) + v20
									if uint32(v20^i32(55296)+i32(-1114112)) < uint32(i32(-1112064)) {
										goto l172
									}
								l171:
									store32(m.memory[int64(uint32(v3))+492:], uint32(v4))
									if v20 == i32(-1) {
										goto l157
									}
									m.fn1072(v3+i32(640), v20)
									m.fn1384(v3+i32(176), v3+i32(640))
									goto l157
								l195:
									store16(m.memory[int64(uint32(v3))+486:], uint16(v20))
									store16(m.memory[int64(uint32(v3))+484:], uint16(i32(1)))
								l172:
									store32(m.memory[int64(uint32(v3))+492:], uint32(v4))
									goto l157
								l194:
									t353 := m.fn15(v21, v19, i32(1079312), i32(7))
									if t353 != 0 {
										goto l193
									}
									{
										t354 := m.fn15(v21, v19, i32(1079319), i32(5))
										if t354 != 0 {
											goto l196
										}
										{
											t355 := m.fn15(v21, v19, i32(1079324), i32(4))
											if t355 != 0 {
												goto l197
											}
											t356 := m.fn15(v21, v19, i32(1079328), i32(5))
											if t356 != 0 {
												t372 := int32(m.memory[int64(uint32(v3))+631])
												if t372 != 0 {
													goto l157
												}
												t373 := int32(m.memory[int64(uint32(v3))+633])
												if t373&i32(255) != i32(2) {
													goto l157
												}
												t374 := int32(load32(m.memory[int64(uint32(v3))+620:]))
												t375 := v27
												v4 = t374
												p376 := i32(1)
												if uint32(v4) > uint32(i32(1)) {
													p376 = v4
												}
												t377 := m.fn1385(t375, p376)
												v4 = t377
												store32(m.memory[int64(uint32(v4))+24:], uint32(i32(0)))
												store64(m.memory[int64(uint32(v4))+16:], uint64(i64(0)))
												store64(m.memory[int64(uint32(v4))+52:], uint64(i64(0)))
												m.memory[int64(uint32(v4))+60] = byte(i32(0))
												goto l157
											}
											t357 := m.fn15(v21, v19, i32(1079333), i32(5))
											if t357 != 0 {
												t378 := int32(m.memory[int64(uint32(v3))+631])
												if t378 != 0 {
													goto l157
												}
												t379 := int32(m.memory[int64(uint32(v3))+633])
												if t379&i32(255) != i32(2) {
													goto l157
												}
												t380 := int32(load32(m.memory[int64(uint32(v3))+620:]))
												t381 := v27
												v4 = t380
												p382 := i32(1)
												if uint32(v4) > uint32(i32(1)) {
													p382 = v4
												}
												t383 := m.fn1385(t381, p382)
												m.memory[int64(uint32(t383))+60] = byte(i32(1))
												goto l157
											}
											t358 := m.fn15(v21, v19, i32(1079338), i32(5))
											if t358 != 0 {
												t384 := int32(m.memory[int64(uint32(v3))+631])
												if t384 != 0 {
													goto l157
												}
												t385 := int32(m.memory[int64(uint32(v3))+633])
												if t385&i32(255) != i32(2) {
													goto l157
												}
												t386 := int32(load32(m.memory[int64(uint32(v3))+620:]))
												t387 := v27
												v4 = t386
												p388 := i32(1)
												if uint32(v4) > uint32(i32(1)) {
													p388 = v4
												}
												t389 := m.fn1385(t387, p388)
												m.memory[int64(uint32(t389))+24] = byte(i32(1))
												goto l157
											}
											t359 := m.fn15(v21, v19, i32(1079343), i32(5))
											if t359 != 0 {
												t390 := int32(m.memory[int64(uint32(v3))+631])
												if t390 != 0 {
													goto l157
												}
												t391 := int32(m.memory[int64(uint32(v3))+633])
												if t391&i32(255) != i32(2) {
													goto l157
												}
												t392 := int32(load32(m.memory[int64(uint32(v3))+620:]))
												t393 := v27
												v4 = t392
												p394 := i32(1)
												if uint32(v4) > uint32(i32(1)) {
													p394 = v4
												}
												t395 := m.fn1385(t393, p394)
												m.memory[int64(uint32(t395))+25] = byte(i32(1))
												goto l157
											}
											t360 := m.fn15(v21, v19, i32(1079348), i32(6))
											if t360 != 0 {
												t396 := int32(m.memory[int64(uint32(v3))+631])
												if t396 != 0 {
													goto l157
												}
												t397 := int32(m.memory[int64(uint32(v3))+633])
												if t397&i32(255) != i32(2) {
													goto l157
												}
												t398 := int32(load32(m.memory[int64(uint32(v3))+620:]))
												t399 := v27
												v4 = t398
												p400 := i32(1)
												if uint32(v4) > uint32(i32(1)) {
													p400 = v4
												}
												t401 := m.fn1385(t399, p400)
												m.memory[int64(uint32(t401))+26] = byte(i32(1))
												goto l157
											}
											t361 := m.fn15(v21, v19, i32(1079354), i32(6))
											if t361 != 0 {
												t402 := int32(m.memory[int64(uint32(v3))+631])
												if t402 != 0 {
													goto l157
												}
												t403 := int32(m.memory[int64(uint32(v3))+633])
												if t403&i32(255) != i32(2) {
													goto l157
												}
												t404 := int32(load32(m.memory[int64(uint32(v3))+620:]))
												t405 := v27
												v4 = t404
												p406 := i32(1)
												if uint32(v4) > uint32(i32(1)) {
													p406 = v4
												}
												t407 := m.fn1385(t405, p406)
												m.memory[int64(uint32(t407))+27] = byte(i32(1))
												goto l157
											}
											t362 := m.fn15(v21, v19, i32(1079360), i32(5))
											if t362 != 0 {
												t408 := int32(m.memory[int64(uint32(v3))+631])
												if t408 != 0 {
													goto l157
												}
												t409 := int32(m.memory[int64(uint32(v3))+633])
												if t409&i32(255) != i32(2) {
													goto l157
												}
												t410 := int32(load32(m.memory[int64(uint32(v3))+620:]))
												t411 := v27
												v21 = t410
												p412 := i32(1)
												if uint32(v21) > uint32(i32(1)) {
													p412 = v21
												}
												t413 := m.fn1385(t411, p412)
												v21 = t413
												store64(m.memory[int64(uint32(v21))+16:], uint64(i64(0)))
												t414 := int64(load64(m.memory[int64(uint32(v21))+24:]))
												v7 = t414
												store32(m.memory[int64(uint32(v21))+24:], uint32(i32(0)))
												p415 := i64(0)
												if v4&i32(1) != 0 {
													p415 = int64(v20)
												}
												v8 = p415
												{
													t416 := int32(load32(m.memory[int64(uint32(v21))+52:]))
													v4 = t416
													t417 := int32(load32(m.memory[int64(uint32(v21))+44:]))
													if v4 != t417 {
														goto l211
													}
													m.fn1145(v21 + i32(44))
												}
											l211:
												t418 := int32(load32(m.memory[int64(uint32(v21))+48:]))
												v20 = t418 + v4<<4
												store64(m.memory[int64(uint32(v20))+8:], uint64(v7))
												store64(m.memory[uint32(v20):], uint64(v8))
												store32(m.memory[int64(uint32(v21))+52:], uint32(v4+i32(1)))
												goto l157
											}
											t363 := m.fn15(v21, v19, i32(1079365), i32(4))
											if t363 != 0 {
												m.fn1382(v3 + i32(176))
												t419 := int32(m.memory[int64(uint32(v3))+631])
												if t419 != 0 {
													goto l157
												}
												t420 := int32(m.memory[int64(uint32(v3))+633])
												if t420&i32(255) != i32(2) {
													goto l157
												}
												m.fn1386(v3+i32(640), v3+i32(176), i32(1))
												t421 := int32(load32(m.memory[int64(uint32(v3))+640:]))
												v4 = t421
												if v4 == i32(-1) {
													goto l157
												}
												t422 := int64(load64(m.memory[int64(uint32(v3))+645:]))
												store64(m.memory[int64(uint32(v3))+1200:], uint64(t422))
												t423 := int64(load64(m.memory[int64(uint32(v3))+653:]))
												store64(m.memory[int64(uint32(v3))+1208:], uint64(t423))
												t424 := int32(load32(m.memory[int64(uint32(v3))+660:]))
												store32(m.memory[int64(uint32(v3))+1215:], uint32(t424))
												goto l178
											}
											t364 := m.fn15(v21, v19, i32(1079369), i32(8))
											if t364 != 0 {
												m.fn1382(v3 + i32(176))
												t425 := int32(m.memory[int64(uint32(v3))+631])
												if t425 != 0 {
													goto l157
												}
												t426 := int32(m.memory[int64(uint32(v3))+633])
												if t426&i32(255) != i32(2) {
													goto l157
												}
												t427 := int32(load32(m.memory[int64(uint32(v3))+620:]))
												t428 := v3 + i32(640)
												t429 := v3 + i32(176)
												v4 = t427
												p430 := i32(2)
												if uint32(v4) > uint32(i32(2)) {
													p430 = v4
												}
												m.fn1386(t428, t429, p430)
												t431 := int32(load32(m.memory[int64(uint32(v3))+640:]))
												v4 = t431
												if v4 == i32(-1) {
													goto l157
												}
												t432 := int64(load64(m.memory[int64(uint32(v3))+645:]))
												store64(m.memory[int64(uint32(v3))+1200:], uint64(t432))
												t433 := int64(load64(m.memory[int64(uint32(v3))+653:]))
												store64(m.memory[int64(uint32(v3))+1208:], uint64(t433))
												t434 := int32(load32(m.memory[int64(uint32(v3))+660:]))
												store32(m.memory[int64(uint32(v3))+1215:], uint32(t434))
												goto l178
											}
											t365 := m.fn15(v21, v19, i32(1077929), i32(3))
											if t365 != 0 {
												m.fn1382(v3 + i32(176))
												t435 := int32(m.memory[int64(uint32(v3))+631])
												if t435 != 0 {
													goto l157
												}
												t436 := int32(m.memory[int64(uint32(v3))+633])
												if t436&i32(255) != i32(2) {
													goto l157
												}
												m.fn1387(v3+i32(640), v3+i32(176), i32(1))
												t437 := int32(load32(m.memory[int64(uint32(v3))+640:]))
												v4 = t437
												if v4 == i32(-1) {
													goto l157
												}
												t438 := int64(load64(m.memory[int64(uint32(v3))+645:]))
												store64(m.memory[int64(uint32(v3))+1200:], uint64(t438))
												t439 := int64(load64(m.memory[int64(uint32(v3))+653:]))
												store64(m.memory[int64(uint32(v3))+1208:], uint64(t439))
												t440 := int32(load32(m.memory[int64(uint32(v3))+660:]))
												store32(m.memory[int64(uint32(v3))+1215:], uint32(t440))
												goto l178
											}
											t366 := m.fn15(v21, v19, i32(1079377), i32(7))
											if t366 != 0 {
												goto l208
											}
											t367 := m.fn15(v21, v19, i32(1079384), i32(14))
											if t367 == 0 {
												t455 := m.fn15(v21, v19, i32(1079144), i32(12))
												if t455 != 0 {
													if v4 != i32(1) {
														goto l157
													}
													if uint32(v20) >= uint32(i32(9)) {
														goto l157
													}
													m.memory[int64(uint32(v3))+628] = byte(i32(1))
													m.memory[int64(uint32(v3))+629] = byte(v20 + i32(1))
													goto l157
												}
												t456 := m.fn15(v21, v19, i32(1072629), i32(4))
												if t456 != 0 {
													t463 := v3
													p462 := i32(0)
													if v20 > i32(0) {
														p462 = v20
													}
													v20 = p462
													p464 := i32(8)
													if v20 < i32(8) {
														p464 = v20
													}
													p465 := i32(0)
													if v4&i32(1) != 0 {
														p465 = p464
													}
													store32(m.memory[int64(uint32(t463))+624:], uint32(p465))
													goto l157
												}
												t457 := m.fn15(v21, v19, i32(0x107774), i32(2))
												if t457 != 0 {
													store32(m.memory[int64(uint32(v3))+604:], uint32(v20))
													store32(m.memory[int64(uint32(v3))+600:], uint32(v4))
													goto l157
												}
												{
													t458 := m.fn15(v21, v19, i32(0x107776), i32(8))
													if t458 != 0 {
														goto l216
													}
													t459 := m.fn15(v21, v19, i32(1079166), i32(6))
													if t459 == 0 {
														t466 := m.fn15(v21, v19, i32(1079172), i32(8))
														if t466 != 0 {
															m.memory[int64(uint32(v3))+635] = byte(i32(0))
															goto l157
														}
														{
															t467 := m.fn15(v21, v19, i32(1079180), i32(9))
															if t467 != 0 {
																goto l219
															}
															t468 := m.fn15(v21, v19, i32(1079189), i32(5))
															if t468 == 0 {
																t469 := m.fn15(v21, v19, i32(1079398), i32(5))
																if t469 != 0 {
																	m.fn1382(v3 + i32(176))
																	t486 := int32(m.memory[int64(uint32(v3))+631])
																	if t486 != 0 {
																		goto l157
																	}
																	t487 := int32(load32(m.memory[int64(uint32(v3))+564:]))
																	v21 = t487
																	t488 := int32(load32(m.memory[int64(uint32(v3))+552:]))
																	v19 = t488
																	{
																		t489 := int32(load32(m.memory[int64(uint32(v3))+376:]))
																		v20 = t489
																		t490 := int32(load32(m.memory[int64(uint32(v3))+368:]))
																		if v20 != t490 {
																			goto l234
																		}
																		m.fn418(v25)
																	}
																l234:
																	t491 := int32(load32(m.memory[int64(uint32(v3))+372:]))
																	v4 = t491 + v20*i32(20)
																	store32(m.memory[int64(uint32(v4))+16:], uint32(v21))
																	store32(m.memory[int64(uint32(v4))+12:], uint32(v19))
																	store32(m.memory[int64(uint32(v4))+8:], uint32(i32(0)))
																	store64(m.memory[uint32(v4):], uint64(i64(0x100000000)))
																	store32(m.memory[int64(uint32(v3))+376:], uint32(v20+i32(1)))
																	goto l157
																}
																t470 := m.fn15(v21, v19, i32(1079403), i32(7))
																if t470 != 0 {
																	m.fn1382(v3 + i32(176))
																	t509 := int32(load32(m.memory[int64(uint32(v3))+376:]))
																	if t509 == 0 {
																		goto l236
																	}
																	m.memory[int64(uint32(v3))+634] = byte(i32(2))
																	goto l236
																}
																t471 := m.fn15(v21, v19, i32(1079410), i32(7))
																if t471 != 0 {
																	m.fn1382(v3 + i32(176))
																	m.memory[int64(uint32(v3))+634] = byte(i32(0))
																	goto l157
																}
																t472 := m.fn15(v21, v19, i32(1079417), i32(8))
																if t472 != 0 {
																	m.fn1382(v3 + i32(176))
																	m.memory[int64(uint32(v3))+631] = byte(i32(0))
																	store16(m.memory[int64(uint32(v3))+633:], uint16(i32(0)))
																	t492 := int32(load32(m.memory[int64(uint32(v3))+564:]))
																	v21 = t492
																	t493 := int32(load32(m.memory[int64(uint32(v3))+552:]))
																	v19 = t493
																	{
																		t494 := int32(load32(m.memory[int64(uint32(v3))+388:]))
																		v4 = t494
																		t495 := int32(load32(m.memory[int64(uint32(v3))+380:]))
																		if v4 != t495 {
																			goto l235
																		}
																		m.fn272(v28)
																	}
																l235:
																	t496 := int32(load32(m.memory[int64(uint32(v3))+384:]))
																	v20 = t496 + v4*i32(12)
																	m.memory[int64(uint32(v20))+8] = byte(i32(0))
																	store32(m.memory[int64(uint32(v20))+4:], uint32(v21))
																	store32(m.memory[uint32(v20):], uint32(v19))
																	store32(m.memory[int64(uint32(v3))+388:], uint32(v4+i32(1)))
																	goto l157
																}
																t473 := m.fn15(v21, v19, i32(1079425), i32(6))
																if t473 != 0 {
																	t497 := int32(load32(m.memory[int64(uint32(v3))+388:]))
																	v4 = t497
																	if v4 == 0 {
																		goto l157
																	}
																	t498 := int32(load32(m.memory[int64(uint32(v3))+384:]))
																	v4 = t498 + v4*i32(12)
																	if v4+i32(-12) == 0 {
																		goto l157
																	}
																	m.memory[uint32(v4+i32(-4))] = byte(i32(1))
																	goto l157
																}
																t474 := m.fn15(v21, v19, i32(1079431), i32(5))
																if t474 != 0 {
																	goto l157
																}
																t475 := m.fn15(v21, v19, i32(1079436), i32(9))
																if t475 != 0 {
																	m.fn1382(v3 + i32(176))
																	m.memory[int64(uint32(v3))+631] = byte(i32(0))
																	m.memory[int64(uint32(v3))+634] = byte(i32(3))
																	goto l157
																}
																t476 := m.fn15(v21, v19, i32(1079445), i32(7))
																if t476 != 0 {
																	goto l210
																}
																t477 := m.fn15(v21, v19, i32(1079452), i32(6))
																if t477 != 0 {
																	goto l210
																}
																t478 := m.fn15(v21, v19, i32(1079458), i32(6))
																if t478 != 0 {
																	goto l210
																}
																t479 := m.fn15(v21, v19, i32(1077556), i32(4))
																if t479 != 0 {
																	t499 := int32(m.memory[int64(uint32(v3))+631])
																	if t499 != 0 {
																		goto l157
																	}
																	m.fn1382(v3 + i32(176))
																	m.memory[int64(uint32(v3))+634] = byte(i32(4))
																	t500 := int32(load32(m.memory[int64(uint32(v3))+552:]))
																	v4 = t500
																	m.fn1388(v32)
																	store32(m.memory[int64(uint32(v3))+456:], uint32(i32(0)))
																	store32(m.memory[int64(uint32(v3))+452:], uint32(v4))
																	store64(m.memory[int64(uint32(v3))+436:], uint64(i64(-0x100000000)))
																	store64(m.memory[int64(uint32(v3))+428:], uint64(i64(0x100000000)))
																	goto l157
																}
																t480 := m.fn15(v21, v19, i32(1079464), i32(7))
																if t480 != 0 {
																	t501 := int32(m.memory[int64(uint32(v3))+634])
																	if t501 != i32(4) {
																		goto l157
																	}
																	t502 := int32(load32(m.memory[int64(uint32(v3))+428:]))
																	if t502 == i32(-1) {
																		goto l157
																	}
																	store32(m.memory[int64(uint32(v3))+468:], uint32(i32(3)))
																	store32(m.memory[int64(uint32(v3))+464:], uint32(i32(1079563)))
																	store32(m.memory[int64(uint32(v3))+460:], uint32(i32(9)))
																	store32(m.memory[int64(uint32(v3))+456:], uint32(i32(1079554)))
																	goto l157
																}
																t481 := m.fn15(v21, v19, i32(1079471), i32(8))
																if t481 != 0 {
																	t503 := int32(m.memory[int64(uint32(v3))+634])
																	if t503 != i32(4) {
																		goto l157
																	}
																	t504 := int32(load32(m.memory[int64(uint32(v3))+428:]))
																	if t504 == i32(-1) {
																		goto l157
																	}
																	store32(m.memory[int64(uint32(v3))+468:], uint32(i32(3)))
																	store32(m.memory[int64(uint32(v3))+464:], uint32(i32(1079551)))
																	store32(m.memory[int64(uint32(v3))+460:], uint32(i32(10)))
																	store32(m.memory[int64(uint32(v3))+456:], uint32(i32(1079541)))
																	goto l157
																}
																t482 := m.fn15(v21, v19, i32(1079479), i32(7))
																if t482 != 0 {
																	t505 := int32(m.memory[int64(uint32(v3))+634])
																	if t505 != i32(4) {
																		goto l157
																	}
																	t506 := int32(load32(m.memory[int64(uint32(v3))+428:]))
																	if t506 == i32(-1) {
																		goto l157
																	}
																	store32(m.memory[int64(uint32(v3))+468:], uint32(i32(3)))
																	store32(m.memory[int64(uint32(v3))+464:], uint32(i32(1079538)))
																	store32(m.memory[int64(uint32(v3))+460:], uint32(i32(9)))
																	store32(m.memory[int64(uint32(v3))+456:], uint32(i32(1079529)))
																	goto l157
																}
																t483 := m.fn15(v21, v19, i32(1079486), i32(9))
																if t483 != 0 {
																	t507 := int32(m.memory[int64(uint32(v3))+634])
																	if t507 != i32(4) {
																		goto l157
																	}
																	t508 := int32(load32(m.memory[int64(uint32(v3))+428:]))
																	if t508 == i32(-1) {
																		goto l157
																	}
																	store32(m.memory[int64(uint32(v3))+468:], uint32(i32(3)))
																	store32(m.memory[int64(uint32(v3))+464:], uint32(i32(1079526)))
																	store32(m.memory[int64(uint32(v3))+460:], uint32(i32(9)))
																	store32(m.memory[int64(uint32(v3))+456:], uint32(i32(1079517)))
																	goto l157
																}
																t484 := m.fn15(v21, v19, i32(1079495), i32(7))
																if t484 != 0 {
																	goto l232
																}
																t485 := m.fn15(v21, v19, i32(1079502), i32(8))
																if t485 == 0 {
																	t510 := m.fn15(v21, v19, i32(1079510), i32(7))
																	if t510 != 0 {
																		goto l232
																	}
																	t511 := m.fn914(v21, v19, i32(1078672), i32(59))
																	if t511 == 0 {
																		goto l157
																	}
																	m.fn1382(v3 + i32(176))
																	m.memory[int64(uint32(v3))+634] = byte(i32(0))
																	goto l236
																}
																goto l232
															}
														}
													l219:
														m.memory[int64(uint32(v3))+635] = byte(i32(1))
														goto l157
													}
												}
											l216:
												m.fn1382(v3 + i32(176))
												m.memory[int64(uint32(v3))+634] = byte(i32(1))
												t460 := int32(load32(m.memory[int64(uint32(v3))+416:]))
												if t460 != i32(-1) {
													goto l157
												}
												t461 := int32(load32(m.memory[int64(uint32(v3))+420:]))
												m.fn134(i32(-1), t461)
												store32(m.memory[int64(uint32(v3))+424:], uint32(i32(0)))
												store64(m.memory[int64(uint32(v3))+416:], uint64(i64(0x100000000)))
												goto l157
											}
											goto l210
										}
									l197:
										t369 := v3
										p368 := i32(0)
										if v20 > i32(0) {
											p368 = v20
										}
										v20 = p368
										p370 := i32(8)
										if v20 < i32(8) {
											p370 = v20
										}
										p371 := i32(1)
										if v4&i32(1) != 0 {
											p371 = p370
										}
										v4 = p371
										store32(m.memory[int64(uint32(t369))+620:], uint32(v4))
										if uint32(v4) <= uint32(i32(1)) {
											goto l157
										}
									}
								l196:
									m.memory[int64(uint32(v3))+630] = byte(i32(1))
									goto l157
								l208:
									m.fn1382(v3 + i32(176))
									t441 := int32(m.memory[int64(uint32(v3))+631])
									if t441 != 0 {
										goto l157
									}
									t442 := int32(m.memory[int64(uint32(v3))+633])
									if t442&i32(255) != i32(2) {
										goto l157
									}
									t443 := int32(load32(m.memory[int64(uint32(v3))+620:]))
									t444 := v3 + i32(640)
									t445 := v3 + i32(176)
									v4 = t443
									p446 := i32(2)
									if uint32(v4) > uint32(i32(2)) {
										p446 = v4
									}
									m.fn1387(t444, t445, p446)
									t447 := int32(load32(m.memory[int64(uint32(v3))+640:]))
									v4 = t447
									if v4 == i32(-1) {
										goto l157
									}
									t448 := int64(load64(m.memory[int64(uint32(v3))+645:]))
									store64(m.memory[int64(uint32(v3))+1200:], uint64(t448))
									t449 := int64(load64(m.memory[int64(uint32(v3))+653:]))
									store64(m.memory[int64(uint32(v3))+1208:], uint64(t449))
									t450 := int32(load32(m.memory[int64(uint32(v3))+660:]))
									store32(m.memory[int64(uint32(v3))+1215:], uint32(t450))
								}
							l178:
								t451 := int32(m.memory[int64(uint32(v3))+644])
								m.memory[int64(uint32(v3))+1108] = byte(t451)
								t452 := int64(load64(m.memory[int64(uint32(v3))+1200:]))
								store64(m.memory[int64(uint32(v3))+1109:], uint64(t452))
								t453 := int64(load64(m.memory[int64(uint32(v3))+1208:]))
								store64(m.memory[int64(uint32(v3))+1117:], uint64(t453))
								t454 := int32(load32(m.memory[int64(uint32(v3))+1215:]))
								store32(m.memory[int64(uint32(v3))+1124:], uint32(t454))
								goto l212
							case 3:
								v4 = v20 & i32(255)
								switch v4 + i32(-123) {
								case 0, 2:
									goto l237
								case 1:
									goto l157
								default:
									switch v4 + i32(-92) {
									case 0:
										goto l237
									case 1, 2:
										goto l157
									case 3:
										m.fn1389(v3+i32(176), i32(45))
										goto l157
									default:
										if v4 != i32(42) {
											goto l157
										}
										goto l236
									}
								case 3:
									m.fn1389(v3+i32(176), i32(160))
									goto l157
								}
							l237:
								m.fn1389(v3+i32(176), v4)
								goto l157
							case 6:
								t512 := int32(m.memory[int64(uint32(v3))+634])
								if t512 != i32(4) {
									goto l157
								}
								t513 := int32(load32(m.memory[int64(uint32(v3))+428:]))
								if t513 == i32(-1) {
									goto l157
								}
								t514 := int32(load32(m.memory[int64(uint32(v3))+552:]))
								t515 := int32(load32(m.memory[int64(uint32(v3))+452:]))
								if t514 != t515 {
									goto l157
								}
								m.fn51(v3+i32(640), v20, v21)
								t516 := int32(load32(m.memory[int64(uint32(v3))+440:]))
								t517 := int32(load32(m.memory[int64(uint32(v3))+444:]))
								m.fn1390(t516, t517)
								t518 := int32(load32(m.memory[int64(uint32(v3))+648:]))
								store32(m.memory[int64(uint32(v31))+8:], uint32(t518))
								t519 := int64(load64(m.memory[int64(uint32(v3))+640:]))
								store64(m.memory[uint32(v31):], uint64(t519))
								goto l157
							case 4, 5:
								t520 := int32(m.memory[int64(uint32(v3))+634])
								v4 = t520
								if v4 == i32(4) {
									t523 := int32(load32(m.memory[int64(uint32(v3))+428:]))
									if t523 == i32(-1) {
										goto l157
									}
									t524 := int32(load32(m.memory[int64(uint32(v3))+552:]))
									t525 := int32(load32(m.memory[int64(uint32(v3))+452:]))
									if t524 != t525 {
										goto l157
									}
									m.fn145(v32, v20)
									goto l157
								}
								{
									if v4 != 0 {
										goto l243
									}
									t521 := int32(m.memory[int64(uint32(v3))+631])
									if t521&i32(1) != 0 {
										goto l157
									}
								}
							l243:
								t522 := int32(load32(m.memory[int64(uint32(v3))+492:]))
								v4 = t522
								if v4 != 0 {
									store32(m.memory[int64(uint32(v3))+492:], uint32(v4+i32(-1)))
									goto l157
								}
								m.fn145(v26, v20)
								goto l157
							}
						}
					}
				l236:
					m.memory[int64(uint32(v3))+631] = byte(i32(1))
					goto l157
				l232:
					t571 := int32(m.memory[int64(uint32(v3))+634])
					if t571 != i32(4) {
						goto l157
					}
					t572 := int32(load32(m.memory[int64(uint32(v3))+428:]))
					if t572 == i32(-1) {
						goto l157
					}
					store32(m.memory[int64(uint32(v3))+456:], uint32(i32(0)))
					goto l157
				}
			l210:
				m.memory[int64(uint32(v3))+631] = byte(i32(0))
				goto l157
			l193:
				m.fn1389(v3+i32(176), i32(32))
				goto l157
			l184:
				m.fn1382(v3 + i32(176))
				t573 := int32(m.memory[int64(uint32(v3))+631])
				if t573 != 0 {
					goto l157
				}
				store32(m.memory[int64(uint32(v3))+640:], uint32(i32(8)))
				m.fn1340(v6, v3+i32(640))
				goto l157
			}
		l159:
			m.memory[int64(uint32(v3))+636] = byte(i32(1))
		l160:
			t574 := int32(load32(m.memory[int64(uint32(v3))+372:]))
			t575 := int32(load32(m.memory[int64(uint32(v3))+376:]))
			v20 = t575
			v4 = t574 + v20*i32(20) + i32(-20)
		l274:
			{
				if v20 == 0 {
					goto l253
				}
				if v4 == 0 {
					goto l253
				}
				t576 := int32(load32(m.memory[uint32(v4+i32(12)):]))
				if uint32(t576) > uint32(v19) {
					goto l254
				}
			}
		l253:
			store32(m.memory[int64(uint32(v3))+376:], uint32(v20))
			{
				{
				l262:
					{
						{
							t577 := int32(load32(m.memory[int64(uint32(v3))+388:]))
							v4 = t577
							if v4 == 0 {
								goto l255
							}
							t578 := int32(load32(m.memory[int64(uint32(v3))+384:]))
							v21 = t578
							v20 = v21 + v4*i32(12) + i32(-12)
							if v20 == 0 {
								goto l255
							}
							t579 := int32(load32(m.memory[uint32(v20):]))
							if uint32(t579) > uint32(v19) {
								t589 := v3
								v4 = v4 + i32(-1)
								store32(m.memory[int64(uint32(t589))+388:], uint32(v4))
								v4 = v21 + v4*i32(12)
								t590 := int32(m.memory[int64(uint32(v4))+8])
								v21 = t590
								if v21 == i32(2) {
									m.fn153(i32(1078192))
									panic("unreachable")
								}
								t591 := int32(load32(m.memory[int64(uint32(v3))+564:]))
								t592 := v3 + i32(640)
								t593 := v6
								v20 = t591
								t594 := int32(load32(m.memory[int64(uint32(v4))+4:]))
								t595 := v20
								v4 = t594
								p596 := v4
								if uint32(v20) < uint32(v4) {
									p596 = t595
								}
								m.fn749(t592, t593, p596)
								m.fn1401(v3+i32(1640), v3+i32(640))
								{
									t597 := int32(load32(m.memory[int64(uint32(v3))+1644:]))
									t598 := int32(load32(m.memory[int64(uint32(v3))+1648:]))
									t599 := m.fn23(t597, t598)
									if t599 != 0 {
										m.fn894(v3 + i32(1640))
										goto l262
									}
									t600 := int32(load32(m.memory[int64(uint32(v3))+400:]))
									store32(m.memory[int64(uint32(v3))+1520:], uint32(t600))
									store32(m.memory[int64(uint32(v3))+1236:], uint32(i32(5)))
									store32(m.memory[int64(uint32(v3))+1232:], uint32(v3+i32(1520)))
									m.fn73(v3+i32(640), i32(0x100061), v3+i32(1232))
									t601 := int32(load32(m.memory[int64(uint32(v3))+640:]))
									v20 = t601
									t602 := int32(load32(m.memory[int64(uint32(v3))+644:]))
									t603 := v3 + i32(640)
									v22 = t602
									t604 := int32(load32(m.memory[int64(uint32(v3))+648:]))
									t605 := v22
									v18 = t604
									m.fn31(t603, t605, v18)
									t606 := m.fn113(i32(8), i32(32))
									v4 = t606
									store32(m.memory[uint32(v4):], uint32(i32(-0x80000000)))
									t607 := int64(load64(m.memory[int64(uint32(v3))+1640:]))
									store64(m.memory[int64(uint32(v4))+4:], uint64(t607))
									t608 := int32(load32(m.memory[int64(uint32(v3))+1648:]))
									store32(m.memory[int64(uint32(v4))+12:], uint32(t608))
									m.memory[int64(uint32(v3))+664] = byte(v21)
									store32(m.memory[int64(uint32(v3))+660:], uint32(i32(1)))
									store32(m.memory[int64(uint32(v3))+656:], uint32(v4))
									store32(m.memory[int64(uint32(v3))+652:], uint32(i32(1)))
									m.fn1230(v24, v3+i32(640))
									store32(m.memory[int64(uint32(v3))+652:], uint32(v18))
									store32(m.memory[int64(uint32(v3))+648:], uint32(v22))
									store32(m.memory[int64(uint32(v3))+644:], uint32(v20))
									store32(m.memory[int64(uint32(v3))+640:], uint32(i32(7)))
									m.fn1340(v6, v3+i32(640))
									goto l262
								}
							}
						}
					l255:
						t580 := int32(m.memory[int64(uint32(v3))+634])
						if t580 == i32(3) {
							goto l257
						}
						t581 := int32(load32(m.memory[int64(uint32(v3))+412:]))
						v20 = t581
						if v20 == 0 {
							goto l257
						}
						store32(m.memory[int64(uint32(v3))+412:], uint32(i32(0)))
						t582 := int32(load32(m.memory[int64(uint32(v3))+404:]))
						v21 = t582
						t583 := int32(load32(m.memory[int64(uint32(v3))+408:]))
						v4 = t583
						store64(m.memory[int64(uint32(v3))+404:], uint64(i64(0x100000000)))
						m.fn46(v3+i32(32), v4, v20)
						t584 := int32(load32(m.memory[int64(uint32(v3))+32:]))
						t585 := int32(load32(m.memory[int64(uint32(v3))+36:]))
						m.fn51(v3+i32(1232), t584, t585)
						t586 := int32(load32(m.memory[int64(uint32(v3))+1240:]))
						if t586 == 0 {
							goto l258
						}
						t587 := int32(load32(m.memory[int64(uint32(v3))+1240:]))
						store32(m.memory[int64(uint32(v17))+8:], uint32(t587))
						t588 := int64(load64(m.memory[int64(uint32(v3))+1232:]))
						store64(m.memory[uint32(v17):], uint64(t588))
						store32(m.memory[int64(uint32(v3))+640:], uint32(i32(6)))
						m.fn1340(v6, v3+i32(640))
						goto l259
					}
				l258:
					t609 := int32(load32(m.memory[int64(uint32(v3))+1232:]))
					t610 := int32(load32(m.memory[int64(uint32(v3))+1236:]))
					m.fn16(t609, t610)
				}
			l259:
				m.fn16(v21, v4)
			l257:
				t611 := int32(load32(m.memory[int64(uint32(v3))+428:]))
				v21 = t611
				if v21 == i32(-1) {
					goto l157
				}
				t612 := int32(load32(m.memory[int64(uint32(v3))+552:]))
				t613 := int32(load32(m.memory[int64(uint32(v3))+452:]))
				if uint32(t612) >= uint32(t613) {
					goto l157
				}
				store32(m.memory[int64(uint32(v3))+428:], uint32(i32(-1)))
				store32(m.memory[int64(uint32(v3))+640:], uint32(v21))
				memory_copy(m.memory, uint32(v17), uint32(v5), uint32(i32(40)))
				t614 := int32(load32(m.memory[int64(uint32(v3))+668:]))
				v22 = t614
				if v22 == 0 {
					m.fn1402(v3 + i32(640))
					goto l157
				}
				t615 := int32(load32(m.memory[int64(uint32(v3))+672:]))
				v12 = t615
				t616 := int64(load64(m.memory[int64(uint32(v3))+676:]))
				store64(m.memory[int64(uint32(v3))+1556:], uint64(t616))
				t617 := int32(load32(m.memory[int64(uint32(v3))+644:]))
				v19 = t617
				{
					{
						t618 := int32(load32(m.memory[int64(uint32(v3))+652:]))
						if t618 == i32(-1) {
							goto l264
						}
						t619 := int32(load32(m.memory[int64(uint32(v3))+448:]))
						v4 = t619
						t620 := int32(load32(m.memory[int64(uint32(v3))+444:]))
						v20 = t620
						t621 := int32(load32(m.memory[int64(uint32(v3))+440:]))
						v18 = t621
						goto l265
					}
				l264:
					t622 := int32(load32(m.memory[int64(uint32(v3))+648:]))
					t623 := v3 + i32(1232)
					v4 = t622
					m.fn140(t623, int32(uint32(v4)>>1))
					v20 = v19 + v4
					v18 = i32(0)
					v4 = v19
				l267:
					{
						if v4 == v20 {
							goto l266
						}
						t624 := int32(m.memory[uint32(v4)])
						m.fn199(v3+i32(24), t624, i32(16))
						v4 = v4 + i32(1)
						t625 := int32(load32(m.memory[int64(uint32(v3))+24:]))
						if t625&i32(1) == 0 {
							goto l267
						}
						t626 := int32(load32(m.memory[int64(uint32(v3))+28:]))
						v15 = t626
						if v18&i32(1) == 0 {
							goto l268
						}
						m.fn145(v3+i32(1232), v23<<4|v15)
						goto l269
					l268:
						v23 = v15
					l269:
						v18 = v18 ^ i32(1)
						goto l267
					}
				l266:
					t627 := int32(load32(m.memory[int64(uint32(v3))+1240:]))
					v4 = t627
					t628 := int32(load32(m.memory[int64(uint32(v3))+1236:]))
					v20 = t628
					t629 := int32(load32(m.memory[int64(uint32(v3))+1232:]))
					v18 = t629
				}
			l265:
				m.fn16(v21, v19)
				{
					{
						if v4 != 0 {
							goto l270
						}
						v4 = i32(-1)
						goto l271
					l270:
						t630 := int32(load32(m.memory[int64(uint32(v3))+348:]))
						store32(m.memory[int64(uint32(v3))+1600:], uint32(t630))
						store32(m.memory[int64(uint32(v3))+1244:], uint32(i32(1)))
						store32(m.memory[int64(uint32(v3))+1236:], uint32(i32(5)))
						store32(m.memory[int64(uint32(v3))+1240:], uint32(v3+i32(1556)))
						store32(m.memory[int64(uint32(v3))+1232:], uint32(v3+i32(1600)))
						m.fn73(v3+i32(1652), i32(0x1000ac), v3+i32(1232))
						m.fn51(v3+i32(1600), v22, v12)
						m.fn1296(v3+i32(1232), v2, v3+i32(1600), v3+i32(1652), v20, v4)
						t631 := int32(load32(m.memory[int64(uint32(v3))+1236:]))
						v21 = t631
						t632 := int32(load32(m.memory[int64(uint32(v3))+1232:]))
						v4 = t632
						if v4 == i32(-1) {
							store32(m.memory[int64(uint32(v3))+1252:], uint32(v21))
							store32(m.memory[int64(uint32(v3))+1248:], uint32(i32(-0x80000000)))
							store64(m.memory[int64(uint32(v3))+1240:], uint64(i64(1)))
							store64(m.memory[int64(uint32(v3))+1232:], uint64(i64(5)))
							m.fn1340(v6, v3+i32(1232))
							m.fn16(v18, v20)
							goto l157
						}
						t633 := int64(load64(m.memory[int64(uint32(v14))+8:]))
						store64(m.memory[int64(uint32(v3))+1528:], uint64(t633))
						t634 := int64(load64(m.memory[uint32(v14):]))
						store64(m.memory[int64(uint32(v3))+1520:], uint64(t634))
						v1 = v21
					}
				l271:
					m.fn16(v18, v20)
					if v4 == i32(-1) {
						goto l157
					}
					t635 := int64(load64(m.memory[int64(uint32(v3))+1528:]))
					store64(m.memory[int64(uint32(v3))+1120:], uint64(t635))
					t636 := int64(load64(m.memory[int64(uint32(v3))+1520:]))
					store64(m.memory[int64(uint32(v3))+1112:], uint64(t636))
					store32(m.memory[int64(uint32(v3))+1108:], uint32(v1))
					goto l212
				}
			}
		l254:
			v20 = v20 + i32(-1)
			{
				t637 := int32(load32(m.memory[uint32(v4):]))
				v22 = t637
				if v22 == i32(-1) {
					goto l273
				}
				t638 := int32(load32(m.memory[uint32(v4+i32(8)):]))
				v18 = t638
				t639 := int32(load32(m.memory[uint32(v4+i32(4)):]))
				v21 = t639
				t640 := int32(load32(m.memory[int64(uint32(v3))+564:]))
				t641 := v3 + i32(640)
				t642 := v6
				v12 = t640
				t643 := int32(load32(m.memory[uint32(v4+i32(16)):]))
				t644 := v12
				v23 = t643
				p645 := v23
				if uint32(v12) < uint32(v23) {
					p645 = t644
				}
				m.fn749(t641, t642, p645)
				m.fn1401(v3+i32(1232), v3+i32(640))
				m.fn1403(v3+i32(640), v21, v18, v3+i32(1232))
				m.fn1341(v6, v3+i32(640))
				m.fn16(v22, v21)
				v4 = v4 + i32(-20)
				goto l274
			}
		l273:
		}
		store32(m.memory[int64(uint32(v3))+376:], uint32(v20))
		m.fn153(i32(1078208))
		panic("unreachable")
	l212:
		t646 := int32(load32(m.memory[int64(uint32(v3))+1124:]))
		store32(m.memory[int64(uint32(v0))+24:], uint32(t646))
		t647 := int64(load64(m.memory[int64(uint32(v3))+1116:]))
		store64(m.memory[int64(uint32(v0))+16:], uint64(t647))
		t648 := int64(load64(m.memory[int64(uint32(v3))+1108:]))
		store64(m.memory[int64(uint32(v0))+8:], uint64(t648))
		store32(m.memory[uint32(v0):], uint32(i32(-1)))
		store32(m.memory[int64(uint32(v0))+4:], uint32(v4))
		m.fn1395(v3 + i32(176))
	}
l1:
	m.g0 = v3 + i32(1664)
}
func (m *Module) fn1198(v0, v1, v2, v3, v4 int32) {
	var v5 int32
	var v6 int64
	var v7, v8, v9, v10, v11, v12 int32
	var v13, v14 int64
	var v15, v16 int32
	t0 := m.g0
	v5 = t0 - i32(240)
	m.g0 = v5
	store32(m.memory[int64(uint32(v5))+28:], uint32(v4))
	store32(m.memory[int64(uint32(v5))+24:], uint32(v3))
	m.fn637(v5+i32(120), v3, v4)
	t1 := int64(load64(m.memory[int64(uint32(v5))+124:]))
	v6 = t1
	{
		{
			{
				t2 := int32(load32(m.memory[int64(uint32(v5))+120:]))
				v7 = t2
				if v7 != i32(-1) {
					goto l0
				}
				v8 = int32(int64(uint64(v6) >> 32))
				v9 = int32(v6)
				goto l1
			}
		l0:
			t3 := int32(load32(m.memory[int64(uint32(v5))+124:]))
			v10 = t3
			t4 := v5 + i32(216)
			v8 = int32(v6)
			t5 := v8
			v9 = int32(int64(uint64(v6) >> 32))
			m.fn1028(t4, t5, v9)
			m.fn638(v5+i32(16), v1, v8, v9)
			{
				{
					{
						t6 := int32(load32(m.memory[int64(uint32(v5))+16:]))
						if t6 != i32(1) {
							store32(m.memory[int64(uint32(v5))+108:], uint32(i32(166)))
							store32(m.memory[int64(uint32(v5))+104:], uint32(v5+i32(216)))
							m.fn73(v5+i32(120), i32(1051893), v5+i32(104))
							m.fn580(v5+i32(32)|i32(4), i32(0), v5+i32(120))
							goto l8
						}
						t7 := int32(load32(m.memory[int64(uint32(v5))+20:]))
						v11 = t7
						m.fn640(v5+i32(8), v1)
						t8 := int32(load32(m.memory[int64(uint32(v5))+12:]))
						v8 = t8
						t9 := int32(load32(m.memory[int64(uint32(v5))+8:]))
						v9 = t9
						t10 := int32(load32(m.memory[uint32(v9+i32(80)):]))
						t11 := int32(load32(m.memory[uint32(v9+i32(84)):]))
						t12 := m.fn590(t10, t11, v11)
						t13 := int32(m.memory[int64(uint32(t12))+72])
						v9 = t13
						m.fn641(v8)
						{
							if v9 != i32(2) {
								store32(m.memory[int64(uint32(v5))+108:], uint32(i32(166)))
								store32(m.memory[int64(uint32(v5))+104:], uint32(v5+i32(216)))
								m.fn73(v5+i32(120), i32(1051934), v5+i32(104))
								m.fn580(v5+i32(32)|i32(4), i32(20), v5+i32(120))
								goto l8
							}
							m.fn642(v5+i32(120), v1+i32(8))
							t14 := int32(load32(m.memory[int64(uint32(v5))+128:]))
							v8 = t14
							t15 := int32(load32(m.memory[int64(uint32(v5))+124:]))
							v9 = t15
							t16 := int32(load32(m.memory[uint32(v9+i32(80)):]))
							t17 := int32(load32(m.memory[uint32(v9+i32(84)):]))
							t18 := m.fn590(t16, t17, v11)
							t19 := int64(load64(m.memory[int64(uint32(t18))+32:]))
							v6 = t19
							m.fn641(v8)
						l4:
							{
								t20 := int32(load32(m.memory[int64(uint32(v1))+4:]))
								v8 = t20
							l6:
								{
									if v8 == i32(-1) {
										goto l4
									}
									if v8 <= i32(-1) {
										store64(m.memory[int64(uint32(v5))+120:], uint64(int64(uint32(i32(1)))<<32|int64(uint32(i32(1080020)))))
										m.fn91(i32(1052692), v5+i32(120), i32(1080028))
										panic("unreachable")
									}
									t21 := int32(load32(m.memory[int64(uint32(v1))+4:]))
									t22 := v1
									t23 := v8 + i32(1)
									v9 = t21
									p24 := v9
									if v9 == v8 {
										p24 = t23
									}
									store32(m.memory[int64(uint32(t22))+4:], uint32(p24))
									var p25 int32
									if v9 != v8 {
										p25 = 1
									}
									v12 = p25
									v8 = v9
									if v12 != 0 {
										goto l6
									}
								}
							}
							m.fn1342(v5+i32(120), i32(1024), i32(1), i32(1), i32(1))
							t26 := int32(load32(m.memory[int64(uint32(v5))+124:]))
							v8 = t26
							t27 := int32(load32(m.memory[int64(uint32(v5))+120:]))
							if t27 != i32(1) {
								store32(m.memory[int64(uint32(v5))+84:], uint32(v11))
								store32(m.memory[int64(uint32(v5))+80:], uint32(v1))
								v12 = i32(0)
								store32(m.memory[int64(uint32(v5))+48:], uint32(i32(0)))
								store64(m.memory[int64(uint32(v5))+40:], uint64(i64(1024)))
								t29 := int32(load32(m.memory[int64(uint32(v5))+128:]))
								store32(m.memory[int64(uint32(v5))+36:], uint32(t29))
								store32(m.memory[int64(uint32(v5))+32:], uint32(v8))
								store32(m.memory[int64(uint32(v5))+72:], uint32(i32(0)))
								store64(m.memory[int64(uint32(v5))+64:], uint64(i64(0)))
								store64(m.memory[int64(uint32(v5))+56:], uint64(v6))
								t31 := v5
								p30 := i32(1024)
								if uint32(v2) > uint32(i32(1024)) {
									p30 = v2
								}
								store32(m.memory[int64(uint32(t31))+52:], uint32(p30))
								t32 := int32(load32(m.memory[int64(uint32(v5))+216:]))
								t33 := int32(load32(m.memory[int64(uint32(v5))+220:]))
								m.fn16(t32, t33)
								m.fn639(v7, v10)
								t34 := int64(load64(m.memory[int64(uint32(v5))+52:]))
								t35 := v5
								v6 = t34
								store64(m.memory[int64(uint32(t35))+208:], uint64(v6))
								t36 := int64(load64(m.memory[int64(uint32(v5))+44:]))
								t37 := v5
								v13 = t36
								store64(m.memory[int64(uint32(t37))+200:], uint64(v13))
								t38 := int64(load64(m.memory[int64(uint32(v5))+36:]))
								t39 := v5
								v14 = t38
								store64(m.memory[int64(uint32(t39))+192:], uint64(v14))
								t40 := int32(load32(m.memory[int64(uint32(v5))+84:]))
								store32(m.memory[int64(uint32(v5))+172:], uint32(t40))
								t41 := int64(load64(m.memory[int64(uint32(v5))+76:]))
								store64(m.memory[int64(uint32(v5))+164:], uint64(t41))
								t42 := int64(load64(m.memory[int64(uint32(v5))+68:]))
								store64(m.memory[int64(uint32(v5))+156:], uint64(t42))
								t43 := int64(load64(m.memory[int64(uint32(v5))+60:]))
								store64(m.memory[int64(uint32(v5))+148:], uint64(t43))
								store64(m.memory[int64(uint32(v5))+124:], uint64(v14))
								store64(m.memory[int64(uint32(v5))+132:], uint64(v13))
								store64(m.memory[int64(uint32(v5))+140:], uint64(v6))
								store32(m.memory[int64(uint32(v5))+100:], uint32(i32(0)))
								store64(m.memory[int64(uint32(v5))+92:], uint64(i64(0x100000000)))
								store32(m.memory[int64(uint32(v5))+120:], uint32(v8))
								store64(m.memory[int64(uint32(v5))+176:], uint64(i64(0x8000001)))
								store64(m.memory[int64(uint32(v5))+184:], uint64(i64(0x8000001)))
								m.fn962(v5+i32(32), v5+i32(120), v5+i32(92))
								{
									t44 := int32(m.memory[int64(uint32(v5))+32])
									if t44 == i32(255) {
										t46 := int32(load32(m.memory[int64(uint32(v5))+36:]))
										if t46 == 0 {
											goto l11
										}
										v11 = i32(8192)
										t47 := int32(load32(m.memory[int64(uint32(v5))+92:]))
										v10 = t47
										t48 := int32(load32(m.memory[int64(uint32(v5))+100:]))
										v12 = t48
									l27:
										{
											if v10|v12 != 0 {
												goto l12
											}
											m.fn962(v5+i32(32), v5+i32(120), v5+i32(92))
											{
												t49 := int32(m.memory[int64(uint32(v5))+32])
												if t49 == i32(255) {
													goto l13
												}
												t50 := int64(load64(m.memory[int64(uint32(v5))+32:]))
												v6 = t50
												v12 = int32(int64(uint64(v6) >> 32))
												goto l10
											}
										l13:
											t51 := int32(load32(m.memory[int64(uint32(v5))+100:]))
											v12 = t51
											t52 := int32(load32(m.memory[int64(uint32(v5))+36:]))
											if t52 == 0 {
												goto l11
											}
											t53 := int32(load32(m.memory[int64(uint32(v5))+92:]))
											v10 = t53
										}
									l12:
										{
											{
												if v12 != v10 {
													goto l14
												}
												t54 := m.fn351(v5+i32(92), v10, i32(32))
												if t54 != i32(-1) {
													v8 = i32(1)
													v12 = i32(0)
													v6 = i64(9728)
													goto l29
												}
												t55 := int32(load32(m.memory[int64(uint32(v5))+92:]))
												v10 = t55
												t56 := int32(load32(m.memory[int64(uint32(v5))+100:]))
												v12 = t56
											}
										l14:
											t57 := int32(load32(m.memory[int64(uint32(v5))+96:]))
											v9 = t57
											v8 = i32(0)
											m.memory[int64(uint32(v5))+228] = byte(i32(0))
											store32(m.memory[int64(uint32(v5))+224:], uint32(i32(0)))
											store32(m.memory[int64(uint32(v5))+216:], uint32(v9+v12))
											t58 := v5
											t59 := v11
											v15 = v10 - v12
											p60 := v15
											if uint32(v11) < uint32(v15) {
												p60 = t59
											}
											v16 = p60
											store32(m.memory[int64(uint32(t58))+220:], uint32(v16))
										l28:
											{
												{
													{
														t61 := int64(load64(m.memory[int64(uint32(v5))+184:]))
														v6 = t61
														if v6 != i64(0) {
															goto l16
														}
														m.memory[int64(uint32(v5))+232] = byte(i32(255))
														goto l17
													}
												l16:
													{
														{
															t62 := int32(load32(m.memory[int64(uint32(v5))+220:]))
															t63 := v6
															v9 = t62 - v8
															if uint64(t63) < uint64(uint32(v9)) {
																goto l18
															}
															m.fn952(v5+i32(232), v5+i32(120), v5+i32(216))
															t64 := int64(load64(m.memory[int64(uint32(v5))+184:]))
															t65 := int32(load32(m.memory[int64(uint32(v5))+224:]))
															t66 := v5
															v9 = t65
															store64(m.memory[int64(uint32(t66))+184:], uint64(t64-int64(uint32(v9-v8))))
															goto l19
														}
													l18:
														t67 := int32(m.memory[int64(uint32(v5))+228])
														v1 = t67
														m.memory[int64(uint32(v5))+44] = byte(i32(0))
														store32(m.memory[int64(uint32(v5))+40:], uint32(i32(0)))
														t68 := v5
														v7 = int32(v6)
														store32(m.memory[int64(uint32(t68))+36:], uint32(v7))
														t69 := int32(load32(m.memory[int64(uint32(v5))+216:]))
														t70 := v5
														v2 = t69 + v8
														store32(m.memory[int64(uint32(t70))+32:], uint32(v2))
														{
															{
																if v1 != 0 {
																	goto l20
																}
																m.fn952(v5+i32(232), v5+i32(120), v5+i32(32))
																t71 := int32(load32(m.memory[int64(uint32(v5))+40:]))
																v1 = t71
																t72 := int32(m.memory[int64(uint32(v5))+44])
																if t72 == 0 {
																	goto l21
																}
																m.fn1094(v2+v7, v9-v7)
																m.memory[int64(uint32(v5))+228] = byte(i32(1))
																goto l21
															}
														l20:
															m.memory[int64(uint32(v5))+44] = byte(i32(1))
															m.fn952(v5+i32(232), v5+i32(120), v5+i32(32))
															t73 := int32(load32(m.memory[int64(uint32(v5))+40:]))
															v1 = t73
														}
													l21:
														t74 := v5
														v9 = v8 + v1
														store32(m.memory[int64(uint32(t74))+224:], uint32(v9))
														t75 := int64(load64(m.memory[int64(uint32(v5))+184:]))
														store64(m.memory[int64(uint32(v5))+184:], uint64(t75-int64(uint32(v1))))
													}
												l19:
													t76 := int32(m.memory[int64(uint32(v5))+232])
													if t76 != i32(255) {
														t78 := m.fn313(v5 + i32(232))
														if t78 != 0 {
															t81 := int32(load32(m.memory[int64(uint32(v5))+232:]))
															t82 := int32(load32(m.memory[int64(uint32(v5))+236:]))
															m.fn119(t81, t82)
															v8 = v9
															goto l28
														}
														store32(m.memory[int64(uint32(v5))+100:], uint32(v12+v9))
														t79 := int64(load64(m.memory[int64(uint32(v5))+232:]))
														v6 = t79
														v12 = int32(int64(uint64(v6) >> 32))
														goto l10
													}
													v8 = v9
												}
											l17:
												t77 := v5
												v12 = v12 + v8
												store32(m.memory[int64(uint32(t77))+100:], uint32(v12))
												if v8 != 0 {
													{
														t80 := int32(m.memory[int64(uint32(v5))+228])
														if t80&i32(1) == 0 {
															goto l26
														}
														if uint32(v15) < uint32(v11) {
															goto l27
														}
														if v8 != v16 {
															goto l27
														}
														if v11 <= i32(-1) {
															goto l26
														}
														v11 = v11 << 1
														goto l27
													}
												l26:
													v11 = i32(-1)
													goto l27
												}
												v8 = i32(255)
												v6 = i64(0)
												goto l24
											}
										}
									}
									t45 := int64(load64(m.memory[int64(uint32(v5))+32:]))
									v6 = t45
									v12 = int32(int64(uint64(v6) >> 32))
									goto l10
								}
							}
							t28 := int32(load32(m.memory[int64(uint32(v5))+128:]))
							m.fn2(v8, t28)
							panic("unreachable")
						}
					}
				l10:
					v8 = int32(v6)
				l24:
					if v8&i32(255) == i32(255) {
						goto l11
					}
				l29:
					store64(m.memory[int64(uint32(v5))+232:], uint64(int64(uint32(v12))<<32|v6&i64(0xffffff00)|int64(uint32(v8))&i64(255)))
					m.fn51(v5+i32(44), v3, v4)
					store32(m.memory[int64(uint32(v5))+220:], uint32(i32(11)))
					store32(m.memory[int64(uint32(v5))+216:], uint32(v5+i32(232)))
					m.fn73(v5+i32(32), i32(1051912), v5+i32(216))
					t83 := int32(m.memory[int64(uint32(v5))+232])
					t84 := int32(load32(m.memory[int64(uint32(v5))+236:]))
					m.fn119(t83, t84)
					t85 := int64(load64(m.memory[int64(uint32(v5))+40:]))
					store64(m.memory[int64(uint32(v5))+104:], uint64(t85))
					t86 := int64(load64(m.memory[int64(uint32(v5))+48:]))
					store64(m.memory[int64(uint32(v5))+112:], uint64(t86))
					t87 := int32(load32(m.memory[int64(uint32(v5))+36:]))
					v12 = t87
					t88 := int32(load32(m.memory[int64(uint32(v5))+32:]))
					v8 = t88
					if v8 == i32(-1) {
						goto l11
					}
					t89 := int64(load64(m.memory[int64(uint32(v5))+112:]))
					store64(m.memory[int64(uint32(v0))+16:], uint64(t89))
					t90 := int64(load64(m.memory[int64(uint32(v5))+104:]))
					store64(m.memory[int64(uint32(v0))+8:], uint64(t90))
					store32(m.memory[int64(uint32(v0))+4:], uint32(v12))
					store32(m.memory[uint32(v0):], uint32(v8))
					m.fn1344(v5 + i32(120))
					goto l30
				}
			l11:
				m.fn1344(v5 + i32(120))
				{
					if uint32(v12) > uint32(i32(0x8000000)) {
						goto l31
					}
					t91 := int32(load32(m.memory[int64(uint32(v5))+100:]))
					store32(m.memory[int64(uint32(v0))+12:], uint32(t91))
					t92 := int64(load64(m.memory[int64(uint32(v5))+92:]))
					store64(m.memory[int64(uint32(v0))+4:], uint64(t92))
					store32(m.memory[uint32(v0):], uint32(i32(-1)))
					goto l32
				}
			l31:
				store32(m.memory[int64(uint32(v5))+124:], uint32(i32(1)))
				store32(m.memory[int64(uint32(v5))+120:], uint32(v5+i32(24)))
				m.fn73(v0+i32(4), i32(1067262), v5+i32(120))
				store32(m.memory[int64(uint32(v0))+20:], uint32(i32(15)))
				store32(m.memory[int64(uint32(v0))+16:], uint32(i32(1072424)))
				store32(m.memory[uint32(v0):], uint32(i32(-0x7ffffffd)))
			l30:
				t93 := int32(load32(m.memory[int64(uint32(v5))+92:]))
				t94 := int32(load32(m.memory[int64(uint32(v5))+96:]))
				m.fn16(t93, t94)
				goto l32
			}
		l8:
			t95 := int32(load32(m.memory[int64(uint32(v5))+216:]))
			t96 := int32(load32(m.memory[int64(uint32(v5))+220:]))
			m.fn16(t95, t96)
			m.fn639(v7, v10)
			t97 := int32(load32(m.memory[int64(uint32(v5))+40:]))
			v8 = t97
			t98 := int32(load32(m.memory[int64(uint32(v5))+36:]))
			v9 = t98
		}
	l1:
		m.fn51(v5+i32(192)|i32(4), v3, v4)
		store32(m.memory[int64(uint32(v5))+192:], uint32(i32(-0x7ffffffc)))
		m.fn119(v9, v8)
		t99 := int64(load64(m.memory[int64(uint32(v5))+208:]))
		store64(m.memory[int64(uint32(v0))+16:], uint64(t99))
		t100 := int64(load64(m.memory[int64(uint32(v5))+200:]))
		store64(m.memory[int64(uint32(v0))+8:], uint64(t100))
		t101 := int64(load64(m.memory[int64(uint32(v5))+192:]))
		store64(m.memory[uint32(v0):], uint64(t101))
	}
l32:
	m.g0 = v5 + i32(240)
}
func (m *Module) fn1199(v0, v1 int32) {
	{
		t0 := int32(load32(m.memory[uint32(v1):]))
		if t0 != i32(-1) {
			goto l0
		}
		t1 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		store32(m.memory[int64(uint32(v0))+8:], uint32(t1))
		t2 := int64(load64(m.memory[int64(uint32(v1))+4:]))
		store64(m.memory[uint32(v0):], uint64(t2))
		return
	}
l0:
	store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
	store64(m.memory[uint32(v0):], uint64(i64(0x100000000)))
	m.fn1297(v1)
}
func (m *Module) fn1200(v0, v1, v2 int32) {
	m.fn51(v0, v1, v2)
	store32(m.memory[int64(uint32(v0))+12:], uint32(i32(-1)))
}
func (m *Module) fn1201(v0, v1 int32) {
	m.fn136(v0, v1, i32(4), i32(24))
}
func (m *Module) fn1202(v0 int32) {
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
		v3 = v3 + i32(12)
		goto l1
	}
l0:
	t4 := int32(load32(m.memory[uint32(v0):]))
	m.fn136(t4, v2, i32(4), i32(12))
}
func (m *Module) fn1203(v0, v1 int32) {
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
func (m *Module) fn1204(v0, v1, v2, v3, v4, v5, v6, v7, v8 int32) {
	var v9, v10, v11, v12, v13, v14, v15, v16, v17, v18, v19, v20, v21, v22, v23, v24, v25 int32
	t0 := m.g0
	v9 = t0 - i32(96)
	m.g0 = v9
	v10 = i32(0)
	store32(m.memory[int64(uint32(v9))+28:], uint32(i32(0)))
	store64(m.memory[int64(uint32(v9))+20:], uint64(i64(0x400000000)))
	{
		if uint32(v2) < uint32(v5) {
			goto l0
		}
		if uint32(v2-v5) < uint32(i32(4)) {
			goto l0
		}
		t1 := int32(load32(m.memory[uint32(v1+v5):]))
		v10 = t1
	}
l0:
	v11 = v10
	{
		t2 := v2
		v5 = v5 + i32(4)
		if uint32(t2) < uint32(v5) {
			goto l1
		}
		v11 = v10
		if uint32(v2-v5) < uint32(i32(4)) {
			goto l1
		}
		t3 := int32(load32(m.memory[uint32(v1+v5):]))
		v5 = v10 + t3
		p4 := v5
		if uint32(v5) < uint32(v10) {
			p4 = i32(-1)
		}
		v11 = p4
	}
l1:
	{
		if uint32(v11) > uint32(v4) {
			goto l2
		}
		v12 = v11 - v10
		if uint32(v12) < uint32(i32(8)) {
			goto l2
		}
		v13 = v3 + v10
		p5 := i32(1)
		if v6 != 0 {
			p5 = i32(13)
		}
		v14 = p5
		v15 = v12 + i32(-4)
		v3 = int32(uint32(v15) >> 3)
		v16 = v9 + i32(80)
		v17 = i32(4)
		v18 = i32(0)
		v5 = i32(0)
	l6:
		{
			t6 := v15
			v10 = (v3 + v5) << 2
			v11 = t6 - v10
			v10 = v10 + i32(4)
		l5:
			if v3 == v5 {
				goto l2
			}
			{
				if uint32(v12) < uint32(v10) {
					goto l3
				}
				if uint32(v11) < uint32(i32(4)) {
					goto l3
				}
				t7 := int32(load32(m.memory[uint32(v13+v10):]))
				t8 := v2
				v4 = t7 << 9 & i32(0x7ffffe00)
				var p9 int32
				if uint32(t8) < uint32(v4) {
					p9 = 1
				}
				var p10 int32
				if uint32(v2-v4) < uint32(i32(512)) {
					p10 = 1
				}
				v19 = p9 | p10
				if v19 != 0 {
					goto l3
				}
				v20 = v1 + v4
				t11 := int32(m.memory[int64(uint32(v20))+511])
				v21 = t11
				if v21 != 0 {
					v5 = v5 + i32(1)
					v11 = i32(0)
					p12 := v20
					if v19 != 0 {
						p12 = i32(0)
					}
					v22 = p12
					v19 = v21<<2 + i32(4)
				l15:
					{
						v10 = v11 << 2
						v4 = v19 + v14*v11
					l9:
						if v21 == v11 {
							goto l6
						}
						if uint32(v10) > uint32(i32(511)) {
							goto l7
						}
						if v10 == i32(508) {
							goto l7
						}
						if uint32(v4) < uint32(i32(512)) {
							store16(m.memory[int64(uint32(v9))+92:], uint16(i32(0)))
							store64(m.memory[int64(uint32(v9))+84:], uint64(i64(1)))
							store64(m.memory[int64(uint32(v9))+76:], uint64(i64(33686018)))
							m.memory[int64(uint32(v9))+74] = byte(i32(2))
							m.memory[int64(uint32(v9))+72] = byte(i32(0))
							store16(m.memory[int64(uint32(v9))+68:], uint16(i32(0)))
							store32(m.memory[int64(uint32(v9))+40:], uint32(i32(-1)))
							store32(m.memory[int64(uint32(v9))+32:], uint32(i32(0)))
							v10 = v20 + v10
							t13 := int32(load32(m.memory[int64(uint32(v10))+4:]))
							v23 = t13
							t14 := int32(load32(m.memory[uint32(v10):]))
							v24 = t14
							{
								t15 := int32(m.memory[uint32(v20+v4)])
								v10 = t15
								if v10 == 0 {
									goto l10
								}
								t16 := v22
								v4 = v10 << 1
								v25 = t16 + v4
								t17 := int32(m.memory[uint32(v25)])
								v10 = t17
								if v6 != 0 {
									goto l11
								}
								if uint32(v4^i32(511)) < uint32(v10) {
									goto l10
								}
								m.fn51(v16, v25+i32(1), v10)
								m.fn16(i32(0), i32(1))
								goto l10
							l11:
								{
									if v10 != 0 {
										goto l12
									}
									v4 = v4 + i32(2)
									t18 := int32(m.memory[int64(uint32(v25))+1])
									v10 = t18 << 1
									goto l13
								}
							l12:
								v4 = v4 | i32(1)
								v10 = v10<<1 + i32(-1)
							l13:
								if uint32(i32(512)-v4) < uint32(v10) {
									goto l10
								}
								if uint32(v10) < uint32(i32(2)) {
									goto l10
								}
								t19 := v9
								v4 = v22 + v4
								t20 := int32(load16(m.memory[uint32(v4):]))
								store16(m.memory[int64(uint32(t19))+92:], uint16(t20))
								m.fn148(v9+i32(8), i32(2), v4, v10, i32(1081552))
								t21 := int32(load32(m.memory[int64(uint32(v9))+8:]))
								t22 := int32(load32(m.memory[int64(uint32(v9))+12:]))
								m.fn1213(t21, t22, v7, v8, v9+i32(32))
							}
						l10:
							{
								t23 := int32(load32(m.memory[int64(uint32(v9))+20:]))
								if v18 != t23 {
									goto l14
								}
								m.fn1148(v9 + i32(20))
								t24 := int32(load32(m.memory[int64(uint32(v9))+24:]))
								v17 = t24
							}
						l14:
							v11 = v11 + i32(1)
							v10 = v17 + v18*i32(72)
							memory_copy(m.memory, uint32(v10), uint32(v9+i32(32)), uint32(i32(64)))
							store32(m.memory[int64(uint32(v10))+68:], uint32(v23))
							store32(m.memory[int64(uint32(v10))+64:], uint32(v24))
							t25 := v9
							v18 = v18 + i32(1)
							store32(m.memory[int64(uint32(t25))+28:], uint32(v18))
							goto l15
						}
					l7:
						v11 = v11 + i32(1)
						v4 = v4 + v14
						v10 = v10 + i32(4)
						goto l9
					}
				}
			}
		l3:
			v5 = v5 + i32(1)
			v11 = v11 + i32(-4)
			v10 = v10 + i32(4)
			goto l5
		}
	}
l2:
	t26 := int32(load32(m.memory[int64(uint32(v9))+28:]))
	store32(m.memory[int64(uint32(v0))+8:], uint32(t26))
	t27 := int64(load64(m.memory[int64(uint32(v9))+20:]))
	store64(m.memory[uint32(v0):], uint64(t27))
	m.g0 = v9 + i32(96)
}
func (m *Module) fn1205(v0 int32) {
	var v1 int32
	var v2, v3 int64
	t0 := m.g0
	v1 = t0 - i32(16)
	m.g0 = v1
	m.fn34(v1)
	t1 := int64(load64(m.memory[uint32(v1):]))
	v2 = t1
	t2 := int64(load64(m.memory[int64(uint32(v1))+8:]))
	v3 = t2
	t3 := int64(load64(m.memory[int64(uint32(i32(0)))+1286264:]))
	store64(m.memory[int64(uint32(v0))+64:], uint64(t3))
	t4 := int64(load64(m.memory[int64(uint32(i32(0)))+1286256:]))
	store64(m.memory[int64(uint32(v0))+56:], uint64(t4))
	store64(m.memory[int64(uint32(v0))+80:], uint64(v3))
	store64(m.memory[int64(uint32(v0))+72:], uint64(v2))
	m.memory[int64(uint32(v0))+54] = byte(i32(2))
	m.memory[int64(uint32(v0))+52] = byte(i32(0))
	store64(m.memory[int64(uint32(v0))+44:], uint64(i64(33686018)))
	m.memory[int64(uint32(v0))+42] = byte(i32(2))
	m.memory[int64(uint32(v0))+40] = byte(i32(0))
	store16(m.memory[int64(uint32(v0))+36:], uint16(i32(0)))
	store32(m.memory[int64(uint32(v0))+8:], uint32(i32(-1)))
	store32(m.memory[uint32(v0):], uint32(i32(0)))
	m.g0 = v1 + i32(16)
}
func (m *Module) fn1206(v0, v1 int32) {
	var v2 int32
	{
		t0 := int32(load32(m.memory[uint32(v0):]))
		t1 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		t2 := v1
		v2 = t1
		if uint32(t2) <= uint32(t0-v2) {
			return
		}
		m.fn1842(v0, v2, v1, i32(2), i32(2))
	}
}
func (m *Module) fn1207(v0, v1, v2 int32) {
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
	store32(m.memory[int64(uint32(v0))+4:], uint32(v2))
	store32(m.memory[uint32(v0):], uint32(v1))
}
func (m *Module) fn1208(v0, v1 int32) int32 {
	var v2, v3, v4, v5 int32
	t0 := m.g0
	v2 = t0 - i32(112)
	m.g0 = v2
	v3 = i32(0)
	m.fn59(v2+i32(8), i32(0), i32(1), i32(1))
	store32(m.memory[int64(uint32(v2))+28:], uint32(i32(0)))
	t1 := int64(load64(m.memory[int64(uint32(v2))+8:]))
	store64(m.memory[int64(uint32(v2))+20:], uint64(t1))
	m.fn601(v2+i32(32), v0, v1, i32(1080596), i32(4))
l1:
	{
		m.fn790(v2+i32(100), v2+i32(32))
		t2 := int32(load32(m.memory[int64(uint32(v2))+100:]))
		if t2 != i32(1) {
			m.fn75(v2+i32(20), v0+v3, v1-v3)
			t5 := int32(load32(m.memory[int64(uint32(v2))+20:]))
			v5 = t5
			t6 := int32(load32(m.memory[int64(uint32(v2))+24:]))
			t7 := v2
			v0 = t6
			t8 := int32(load32(m.memory[int64(uint32(v2))+28:]))
			m.fn46(t7, v0, t8)
			t9 := int32(load32(m.memory[uint32(v2):]))
			t10 := int32(load32(m.memory[int64(uint32(v2))+4:]))
			m.fn14(v2+i32(32), t9, t10)
			v4 = i32(0)
			{
				t11 := int32(load32(m.memory[int64(uint32(v2))+36:]))
				v3 = t11
				t12 := int32(load32(m.memory[int64(uint32(v2))+40:]))
				t13 := v3
				v1 = t12
				t14 := m.fn15(t13, v1, i32(1080600), i32(5))
				if t14 != 0 {
					goto l2
				}
				t15 := m.fn15(v3, v1, i32(1080605), i32(13))
				if t15 != 0 {
					goto l2
				}
				t16 := m.fn15(v3, v1, i32(1080618), i32(10))
				if t16 != 0 {
					goto l2
				}
				t17 := m.fn15(v3, v1, i32(1080628), i32(10))
				if t17 != 0 {
					goto l2
				}
				v4 = i32(1)
				t18 := m.fn15(v3, v1, i32(1080638), i32(17))
				if t18 != 0 {
					goto l2
				}
				t19 := m.fn15(v3, v1, i32(1080655), i32(11))
				if t19 != 0 {
					goto l2
				}
				t20 := m.fn15(v3, v1, i32(1080666), i32(17))
				p21 := i32(2)
				if t20 != 0 {
					p21 = i32(1)
				}
				v4 = p21
			}
		l2:
			t22 := int32(load32(m.memory[int64(uint32(v2))+32:]))
			m.fn16(t22, v3)
			m.fn16(v5, v0)
			m.g0 = v2 + i32(112)
			return v4
		}
		t3 := int32(load32(m.memory[int64(uint32(v2))+104:]))
		v4 = t3 - v3
		v5 = v0 + v3
		t4 := int32(load32(m.memory[int64(uint32(v2))+108:]))
		v3 = t4
		m.fn75(v2+i32(20), v5, v4)
		m.fn75(v2+i32(20), i32(1097368), i32(1))
		goto l1
	}
}
func (m *Module) fn1209(v0, v1 int32) int32 {
	var v2 int64
	var v3, v4 int32
	var v5 int64
	var v6, v7 int32
	var v8 int64
	var v9 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
		if t0 != 0 {
			t1 := int64(load64(m.memory[int64(uint32(v0))+16:]))
			t2 := int64(load64(m.memory[int64(uint32(v0))+24:]))
			t3 := m.fn529(t1, t2, v1)
			v2 = t3
			t4 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			v3 = t4
			v4 = v3 & int32(v2)
			v5 = int64(uint64(v2)>>25) & i64(127) * i64(72340172838076673)
			t5 := int32(load32(m.memory[uint32(v0):]))
			v0 = t5
			v6 = v1 & i32(0xffff)
			v7 = i32(0)
			var _ int32
		l5:
			{
				t7 := int64(load64(m.memory[uint32(v0+v4):]))
				v8 = t7
				v2 = v8 ^ v5
				v2 = (v2 ^ i64(-1)) & (v2 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
				{
				l3:
					{
						var p8 int32
						if v2 == 0 {
							p8 = 1
						}
						v1 = p8
						if v1 != 0 {
							goto l1
						}
						t9 := v6
						t10 := v0
						v9 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v2))))>>3) + v4) & v3
						t11 := int32(load16(m.memory[uint32(t10+(i32(0)-v9)*i32(60)+i32(-60)):]))
						if t9 == t11 {
							goto l2
						}
						v2 = (v2 + i64(-1)) & v2
						goto l3
					}
				l1:
					if v8&(v8<<1)&i64(-0x7f7f7f7f7f7f7f80) == 0 {
						t14 := v4
						v7 = v7 + i32(8)
						v4 = (t14 + v7) & v3
						goto l5
					}
				l2:
					p12 := v0 + (i32(0)-v9)*i32(60)
					if v1 != 0 {
						p12 = i32(0)
					}
					p13 := p12 + i32(-56)
					if v1 != 0 {
						p13 = i32(0)
					}
					return p13
				}
			}
		}
		return i32(0)
	}
}
func (m *Module) fn1210(v0, v1 int32) {
	var v2 int32
	t0 := int32(load32(m.memory[int64(uint32(v1))+48:]))
	v2 = t0
	m.fn1429(v0, v1)
	store32(m.memory[int64(uint32(v0))+48:], uint32(v2))
	t1 := int32(m.memory[int64(uint32(v1))+54])
	m.memory[int64(uint32(v0))+54] = byte(t1)
	t2 := int32(m.memory[int64(uint32(v1))+52])
	m.memory[int64(uint32(v0))+52] = byte(t2)
	t3 := int32(m.memory[int64(uint32(v1))+53])
	m.memory[int64(uint32(v0))+53] = byte(t3)
}
func (m *Module) fn1211(v0, v1 int32) int32 {
	var v2 int64
	var v3, v4 int32
	var v5 int64
	var v6, v7 int32
	var v8 int64
	var v9 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
		if t0 != 0 {
			t1 := int64(load64(m.memory[int64(uint32(v0))+16:]))
			t2 := int64(load64(m.memory[int64(uint32(v0))+24:]))
			t3 := m.fn529(t1, t2, v1)
			v2 = t3
			t4 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			v3 = t4
			v4 = v3 & int32(v2)
			v5 = int64(uint64(v2)>>25) & i64(127) * i64(72340172838076673)
			t5 := int32(load32(m.memory[uint32(v0):]))
			v0 = t5
			v6 = v1 & i32(0xffff)
			v7 = i32(0)
			var _ int32
		l5:
			{
				t7 := int64(load64(m.memory[uint32(v0+v4):]))
				v8 = t7
				v2 = v8 ^ v5
				v2 = (v2 ^ i64(-1)) & (v2 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
				{
				l3:
					{
						var p8 int32
						if v2 == 0 {
							p8 = 1
						}
						v1 = p8
						if v1 != 0 {
							goto l1
						}
						t9 := v6
						t10 := v0
						v9 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v2))))>>3) + v4) & v3
						t11 := int32(load16(m.memory[uint32(t10+(i32(0)-v9)*i32(36)+i32(-36)):]))
						if t9 == t11 {
							goto l2
						}
						v2 = (v2 + i64(-1)) & v2
						goto l3
					}
				l1:
					if v8&(v8<<1)&i64(-0x7f7f7f7f7f7f7f80) == 0 {
						t14 := v4
						v7 = v7 + i32(8)
						v4 = (t14 + v7) & v3
						goto l5
					}
				l2:
					p12 := v0 + (i32(0)-v9)*i32(36)
					if v1 != 0 {
						p12 = i32(0)
					}
					p13 := p12 + i32(-32)
					if v1 != 0 {
						p13 = i32(0)
					}
					return p13
				}
			}
		}
		return i32(0)
	}
}
