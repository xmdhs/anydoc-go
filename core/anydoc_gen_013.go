package core

import (
	"math"
	"math/bits"
)

func (m *Module) fn537(v0, v1, v2 int32) {
	var v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13 int32
	t0 := m.g0
	v3 = t0 - i32(32)
	m.g0 = v3
	v4 = v2 - v1
	t1 := int32(uint32(v4) / uint32(i32(20)))
	v5 = t1
	if uint32(v4) >= uint32(i32(-0x2aaaaab7)) {
		m.fn9()
		panic("unreachable")
	}
	{
		if v2 != v1 {
			goto l1
		}
		v5 = i32(0)
		v6 = i32(4)
		goto l2
	l1:
		v2 = v5 * i32(12)
		t2 := m.fn5(v2)
		v6 = t2
		if v6 == 0 {
			m.fn10(i32(4), v2)
			panic("unreachable")
		}
		v7 = i32(0)
	l22:
		v4 = i32(0)
		v8 = i32(1)
		v9 = i32(0)
		{
			v2 = v1 + v7*i32(20)
			t3 := int32(load32(m.memory[uint32(v2):]))
			if t3 == i32(-1) {
				goto l4
			}
			v4 = i32(0)
			store32(m.memory[int64(uint32(v3))+16:], uint32(i32(0)))
			store64(m.memory[int64(uint32(v3))+8:], uint64(i64(0x100000000)))
			{
				t4 := int32(load32(m.memory[int64(uint32(v2))+8:]))
				v8 = t4
				if v8 != 0 {
					goto l5
				}
				v8 = i32(1)
				v9 = i32(0)
				goto l4
			}
		l5:
			t5 := int32(load32(m.memory[int64(uint32(v2))+4:]))
			v2 = t5
			v10 = v2 + v8<<5
			v4 = i32(0)
			v11 = i32(1)
		l21:
			{
				{
					{
						t6 := int32(load32(m.memory[uint32(v2):]))
						if t6 != i32(-0x80000000) {
							v8 = i32(1)
							v4 = i32(0)
							{
								t8 := int32(load32(m.memory[int64(uint32(v3))+8:]))
								v2 = t8
								if v2 != 0 {
									t9 := int32(load32(m.memory[uint32(v11+i32(-4)):]))
									v9 = t9
									v12 = v9 & i32(-8)
									t10 := v12
									v9 = v9 & i32(3)
									p11 := i32(8)
									if v9 != 0 {
										p11 = i32(4)
									}
									if uint32(t10) < uint32(p11+v2) {
										m.fn3(i32(1273840), i32(46), i32(1273888))
										panic("unreachable")
									}
									if v9 == 0 {
										goto l11
									}
									if uint32(v12) > uint32(v2+i32(39)) {
										m.fn3(i32(1273904), i32(46), i32(1273952))
										panic("unreachable")
									}
								l11:
									m.fn1(v11)
									v9 = i32(0)
									goto l4
								}
								v9 = i32(0)
								goto l4
							}
						}
						t7 := int32(load32(m.memory[int64(uint32(v3))+8:]))
						v9 = t7
						if v4 != 0 {
							goto l7
						}
						v8 = i32(0)
						goto l8
					}
				l7:
					{
						if v9 != v4 {
							goto l13
						}
						m.fn197(v3+i32(8), v4, i32(1), i32(1), i32(1))
						t12 := int32(load32(m.memory[int64(uint32(v3))+12:]))
						v11 = t12
					}
				l13:
					m.memory[uint32(v11+v4)] = byte(i32(10))
					t13 := v3
					v8 = v4 + i32(1)
					store32(m.memory[int64(uint32(t13))+16:], uint32(v8))
					t14 := int32(load32(m.memory[int64(uint32(v3))+8:]))
					v9 = t14
				}
			l8:
				t15 := int32(load32(m.memory[uint32(v2+i32(12)):]))
				v4 = t15
				t16 := int32(load32(m.memory[uint32(v2+i32(8)):]))
				v12 = t16
				store32(m.memory[int64(uint32(v3))+28:], uint32(i32(0)))
				store64(m.memory[int64(uint32(v3))+20:], uint64(i64(0x100000000)))
				m.fn460(v12, v4, v3+i32(20))
				t17 := int32(load32(m.memory[int64(uint32(v3))+24:]))
				v13 = t17
				t18 := int32(load32(m.memory[int64(uint32(v3))+20:]))
				v12 = t18
				{
					{
						t19 := int32(load32(m.memory[int64(uint32(v3))+28:]))
						v4 = t19
						if uint32(v4) <= uint32(v9-v8) {
							goto l14
						}
						m.fn197(v3+i32(8), v8, v4, i32(1), i32(1))
						t20 := int32(load32(m.memory[int64(uint32(v3))+16:]))
						v8 = t20
						goto l15
					}
				l14:
					if v4 == 0 {
						goto l16
					}
				l15:
					t21 := int32(load32(m.memory[int64(uint32(v3))+12:]))
					v11 = t21
					if v4 == 0 {
						goto l16
					}
					memory_copy(m.memory, uint32(v11+v8), uint32(v13), uint32(v4))
				}
			l16:
				t22 := v3
				v4 = v8 + v4
				store32(m.memory[int64(uint32(t22))+16:], uint32(v4))
				{
					if v12 == 0 {
						goto l17
					}
					t23 := int32(load32(m.memory[uint32(v13+i32(-4)):]))
					v8 = t23
					v9 = v8 & i32(-8)
					t24 := v9
					v8 = v8 & i32(3)
					p25 := i32(8)
					if v8 != 0 {
						p25 = i32(4)
					}
					if uint32(t24) < uint32(p25+v12) {
						m.fn3(i32(1273840), i32(46), i32(1273888))
						panic("unreachable")
					}
					if v8 == 0 {
						goto l19
					}
					if uint32(v9) > uint32(v12+i32(39)) {
						m.fn3(i32(1273904), i32(46), i32(1273952))
						panic("unreachable")
					}
				l19:
					m.fn1(v13)
				}
			l17:
				v2 = v2 + i32(32)
				if v2 != v10 {
					goto l21
				}
			}
			t26 := int32(load32(m.memory[int64(uint32(v3))+12:]))
			v8 = t26
			t27 := int32(load32(m.memory[int64(uint32(v3))+8:]))
			v9 = t27
		}
	l4:
		v2 = v6 + v7*i32(12)
		store32(m.memory[int64(uint32(v2))+8:], uint32(v4))
		store32(m.memory[int64(uint32(v2))+4:], uint32(v8))
		store32(m.memory[uint32(v2):], uint32(v9))
		v7 = v7 + i32(1)
		if v7 != v5 {
			goto l22
		}
	}
l2:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v5))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v6))
	store32(m.memory[uint32(v0):], uint32(v5))
	m.g0 = v3 + i32(32)
}
func (m *Module) fn538(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13 int32
	t0 := m.g0
	v2 = t0 - i32(80)
	m.g0 = v2
	m.fn144(v2+i32(16), v0, v1)
	{
		{
			t1 := int32(load32(m.memory[int64(uint32(v2))+20:]))
			v3 = t1
			if v3 != 0 {
				goto l0
			}
			v1 = i32(255)
			goto l1
		}
	l0:
		t2 := int32(load32(m.memory[int64(uint32(v2))+16:]))
		v4 = t2
		v5 = v4 + v3
		t3 := int32(m.memory[uint32(v5+i32(-1))])
		v1 = t3
		v6 = i32(0)
		store32(m.memory[int64(uint32(v2))+32:], uint32(i32(0)))
		store64(m.memory[int64(uint32(v2))+24:], uint64(i64(0x100000000)))
		v7 = i32(1)
		v8 = i32(0)
		{
			t4 := v3
			var p5 int32
			if v1 == i32(37) {
				p5 = 1
			}
			v1 = t4 - p5
			if v1 == 0 {
				goto l2
			}
			v9 = v4 + v1
			v6 = i32(0)
			v8 = i32(1)
			v1 = v4
		l18:
			{
				{
					t6 := int32(int8(m.memory[uint32(v1)]))
					v0 = t6
					if v0 <= i32(-1) {
						goto l3
					}
					v1 = v1 + i32(1)
					v0 = v0 & i32(255)
					goto l4
				}
			l3:
				t7 := int32(m.memory[int64(uint32(v1))+1])
				v10 = t7 & i32(63)
				v11 = v0 & i32(31)
				if uint32(v0) > uint32(i32(-33)) {
					goto l5
				}
				v0 = v11<<6 | v10
				v1 = v1 + i32(2)
				goto l4
			l5:
				t8 := int32(m.memory[int64(uint32(v1))+2])
				v10 = v10<<6 | t8&i32(63)
				if uint32(v0) >= uint32(i32(-16)) {
					goto l6
				}
				v0 = v10 | v11<<12
				v1 = v1 + i32(3)
				goto l4
			l6:
				t9 := int32(m.memory[int64(uint32(v1))+3])
				v0 = v10<<6 | t9&i32(63) | v11<<18&i32(0x1c0000)
				v1 = v1 + i32(4)
			}
		l4:
			switch v0 + i32(-32) {
			case 0, 12:
				goto l7
			default:
				if v0 == i32(95) {
					goto l7
				}
				if v0 == i32(160) {
					goto l7
				}
				fallthrough
			case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11:
				{
					{
						var p10 int32
						if uint32(v0) < uint32(i32(128)) {
							p10 = 1
						}
						v7 = p10
						if v7 == 0 {
							goto l10
						}
						v10 = i32(1)
						goto l11
					}
				l10:
					if uint32(v0) >= uint32(i32(2048)) {
						goto l12
					}
					v10 = i32(2)
					goto l11
				l12:
					p11 := i32(4)
					if uint32(v0) < uint32(i32(65536)) {
						p11 = i32(3)
					}
					v10 = p11
				}
			l11:
				{
					t12 := int32(load32(m.memory[int64(uint32(v2))+24:]))
					if uint32(v10) <= uint32(t12-v6) {
						goto l13
					}
					m.fn197(v2+i32(24), v6, v10, i32(1), i32(1))
					t13 := int32(load32(m.memory[int64(uint32(v2))+28:]))
					v8 = t13
				}
			l13:
				v11 = v8 + v6
				if v7 != 0 {
					goto l14
				}
				v7 = v0&i32(63) | i32(-128)
				v12 = int32(uint32(v0) >> 6)
				if uint32(v0) >= uint32(i32(2048)) {
					v13 = int32(uint32(v0) >> 12)
					v12 = v12&i32(63) | i32(-128)
					if uint32(v0) > uint32(i32(0xffff)) {
						m.memory[int64(uint32(v11))+3] = byte(v7)
						m.memory[int64(uint32(v11))+2] = byte(v12)
						m.memory[int64(uint32(v11))+1] = byte(v13&i32(63) | i32(-128))
						m.memory[uint32(v11)] = byte(int32(uint32(v0)>>18) | i32(-16))
						goto l16
					}
					m.memory[int64(uint32(v11))+2] = byte(v7)
					m.memory[int64(uint32(v11))+1] = byte(v12)
					m.memory[uint32(v11)] = byte(v13 | i32(224))
					goto l16
				}
				m.memory[int64(uint32(v11))+1] = byte(v7)
				m.memory[uint32(v11)] = byte(v12 | i32(192))
				goto l16
			l14:
				m.memory[uint32(v11)] = byte(v0)
			l16:
				t14 := v2
				v6 = v10 + v6
				store32(m.memory[int64(uint32(t14))+32:], uint32(v6))
			}
		l7:
			if v1 != v9 {
				goto l18
			}
			t15 := int32(load32(m.memory[int64(uint32(v2))+28:]))
			v7 = t15
			t16 := int32(load32(m.memory[int64(uint32(v2))+24:]))
			v8 = t16
		}
	l2:
		v9 = v7 + v6
		v1 = v7
		{
		l25:
			if v1 != v9 {
				goto l19
			}
			v1 = i32(0)
			goto l20
		l19:
			{
				{
					t17 := int32(int8(m.memory[uint32(v1)]))
					v0 = t17
					if v0 <= i32(-1) {
						goto l21
					}
					v1 = v1 + i32(1)
					v0 = v0 & i32(255)
					goto l22
				}
			l21:
				t18 := int32(m.memory[int64(uint32(v1))+1])
				v10 = t18 & i32(63)
				v11 = v0 & i32(31)
				if uint32(v0) > uint32(i32(-33)) {
					goto l23
				}
				v0 = v11<<6 | v10
				v1 = v1 + i32(2)
				goto l22
			l23:
				t19 := int32(m.memory[int64(uint32(v1))+2])
				v10 = v10<<6 | t19&i32(63)
				if uint32(v0) >= uint32(i32(-16)) {
					goto l24
				}
				v0 = v10 | v11<<12
				v1 = v1 + i32(3)
				goto l22
			l24:
				t20 := int32(m.memory[int64(uint32(v1))+3])
				v0 = v10<<6 | t20&i32(63) | v11<<18&i32(0x1c0000)
				v1 = v1 + i32(4)
			}
		l22:
			if uint32(v0+i32(-48)) > uint32(i32(9)) {
				goto l25
			}
			m.fn578(v2+i32(24), v7, v6)
			t21 := int32(m.memory[int64(uint32(v2))+24])
			v1 = t21 ^ i32(1)
		}
	l20:
		{
			if v8 == 0 {
				goto l26
			}
			t22 := int32(load32(m.memory[uint32(v7+i32(-4)):]))
			v0 = t22
			v9 = v0 & i32(-8)
			t23 := v9
			v0 = v0 & i32(3)
			p24 := i32(8)
			if v0 != 0 {
				p24 = i32(4)
			}
			if uint32(t23) < uint32(p24+v8) {
				m.fn3(i32(1273840), i32(46), i32(1273888))
				panic("unreachable")
			}
			if v0 == 0 {
				goto l28
			}
			if uint32(v9) > uint32(v8+i32(39)) {
				m.fn3(i32(1273904), i32(46), i32(1273952))
				panic("unreachable")
			}
		l28:
			m.fn1(v7)
		}
	l26:
		if v1&i32(1) == 0 {
			m.fn144(v2+i32(8), v4, v3)
			t25 := int32(load32(m.memory[int64(uint32(v2))+8:]))
			t26 := int32(load32(m.memory[int64(uint32(v2))+12:]))
			m.fn31(v2+i32(24), t25, t26)
			t27 := int32(load32(m.memory[int64(uint32(v2))+28:]))
			v1 = t27
			{
				t28 := int32(load32(m.memory[int64(uint32(v2))+32:]))
				switch t28 + i32(-2) {
				default:
					goto l35
				case 3:
					t29 := int32(load32(m.memory[uint32(v1):]))
					t30 := int32(m.memory[uint32(v1+i32(4))])
					if t29^i32(1936482662)|(t30^i32(101)) != 0 {
						goto l35
					}
					goto l36
				case 1:
					t31 := int32(load16(m.memory[uint32(v1):]))
					t32 := int32(m.memory[uint32(v1+i32(2))])
					if (t31^i32(25977)|(t32^i32(115)))&i32(0xffff) != 0 {
						goto l35
					}
					goto l36
				case 0:
					t33 := int32(load16(m.memory[uint32(v1):]))
					if t33 == i32(28526) {
						goto l36
					}
					goto l35
				case 2:
					t34 := int32(load32(m.memory[uint32(v1):]))
					if t34 != i32(1702195828) {
						goto l35
					}
				}
			}
		l36:
			{
				t35 := int32(load32(m.memory[int64(uint32(v2))+24:]))
				v0 = t35
				if v0 == 0 {
					goto l37
				}
				t36 := int32(load32(m.memory[uint32(v1+i32(-4)):]))
				v9 = t36
				v6 = v9 & i32(-8)
				t37 := v6
				v9 = v9 & i32(3)
				p38 := i32(8)
				if v9 != 0 {
					p38 = i32(4)
				}
				if uint32(t37) < uint32(p38+v0) {
					m.fn3(i32(1273840), i32(46), i32(1273888))
					panic("unreachable")
				}
				if v9 == 0 {
					goto l39
				}
				if uint32(v6) > uint32(v0+i32(39)) {
					m.fn3(i32(1273904), i32(46), i32(1273952))
					panic("unreachable")
				}
			l39:
				m.fn1(v1)
			}
		l37:
			v1 = i32(1)
			goto l1
		}
		v1 = i32(0)
		goto l1
	l35:
		{
			t39 := int32(load32(m.memory[int64(uint32(v2))+24:]))
			v0 = t39
			if v0 == 0 {
				goto l41
			}
			t40 := int32(load32(m.memory[uint32(v1+i32(-4)):]))
			v9 = t40
			v6 = v9 & i32(-8)
			t41 := v6
			v9 = v9 & i32(3)
			p42 := i32(8)
			if v9 != 0 {
				p42 = i32(4)
			}
			if uint32(t41) < uint32(p42+v0) {
				m.fn3(i32(1273840), i32(46), i32(1273888))
				panic("unreachable")
			}
			if v9 == 0 {
				goto l43
			}
			if uint32(v6) > uint32(v0+i32(39)) {
				m.fn3(i32(1273904), i32(46), i32(1273952))
				panic("unreachable")
			}
		l43:
			m.fn1(v1)
		}
	l41:
		v0 = i32(0)
		v6 = i32(1)
		v11 = i32(0)
		v1 = v4
	l52:
		v7 = v11
		if v1 != v5 {
			goto l45
		}
		v7 = v3
		goto l60
	l45:
		{
			{
				t43 := int32(int8(m.memory[uint32(v1)]))
				v10 = t43
				if v10 <= i32(-1) {
					goto l47
				}
				v9 = v1 + i32(1)
				v10 = v10 & i32(255)
				goto l48
			}
		l47:
			t44 := int32(m.memory[int64(uint32(v1))+1])
			v9 = t44 & i32(63)
			v11 = v10 & i32(31)
			if uint32(v10) > uint32(i32(-33)) {
				goto l49
			}
			v10 = v11<<6 | v9
			v9 = v1 + i32(2)
			goto l48
		l49:
			t45 := int32(m.memory[int64(uint32(v1))+2])
			v9 = v9<<6 | t45&i32(63)
			if uint32(v10) >= uint32(i32(-16)) {
				goto l50
			}
			v10 = v9 | v11<<12
			v9 = v1 + i32(3)
			goto l48
		l50:
			t46 := int32(m.memory[int64(uint32(v1))+3])
			v10 = v9<<6 | t46&i32(63) | v11<<18&i32(0x1c0000)
			v9 = v1 + i32(4)
		}
	l48:
		v11 = v9 - v1 + v7
		if v10 == i32(84) {
			goto l51
		}
		v1 = v9
		if v10 != i32(32) {
			goto l52
		}
	l51:
		v6 = v4 + v11
		v0 = v3 - v11
	l60:
		v9 = v0
		if v9 != 0 {
			goto l53
		}
		v9 = i32(0)
		goto l54
	l53:
		{
			v10 = v6 + v9
			v0 = v10 + i32(-1)
			t47 := int32(int8(m.memory[uint32(v0)]))
			v1 = t47
			if v1 > i32(-1) {
				goto l55
			}
			{
				v0 = v10 + i32(-2)
				t48 := int32(m.memory[uint32(v0)])
				v11 = t48
				v5 = int32(int8(v11))
				if v5 < i32(-64) {
					goto l56
				}
				v10 = v11 & i32(31)
				goto l57
			}
		l56:
			{
				{
					v0 = v10 + i32(-3)
					t49 := int32(m.memory[uint32(v0)])
					v11 = t49
					v8 = int32(int8(v11))
					if v8 < i32(-64) {
						goto l58
					}
					v10 = v11 & i32(15)
					goto l59
				}
			l58:
				v0 = v10 + i32(-4)
				t50 := int32(m.memory[uint32(v0)])
				v10 = t50&i32(7)<<6 | v8&i32(63)
			}
		l59:
			v10 = v10<<6 | v5&i32(63)
		l57:
			v1 = v10<<6 | v1&i32(63)
		}
	l55:
		v0 = v0 - v6
		if v1 == i32(90) {
			goto l60
		}
	l54:
		store16(m.memory[int64(uint32(v2))+60:], uint16(i32(1)))
		store32(m.memory[int64(uint32(v2))+56:], uint32(v9))
		store32(m.memory[int64(uint32(v2))+52:], uint32(i32(0)))
		m.memory[int64(uint32(v2))+48] = byte(i32(1))
		store32(m.memory[int64(uint32(v2))+44:], uint32(i32(46)))
		store32(m.memory[int64(uint32(v2))+40:], uint32(v9))
		store32(m.memory[int64(uint32(v2))+36:], uint32(i32(0)))
		store32(m.memory[int64(uint32(v2))+32:], uint32(v9))
		store32(m.memory[int64(uint32(v2))+28:], uint32(v6))
		store32(m.memory[int64(uint32(v2))+24:], uint32(i32(46)))
		m.fn199(v2+i32(68), v2+i32(24))
		{
			{
				t51 := int32(load32(m.memory[int64(uint32(v2))+68:]))
				if t51 != i32(1) {
					goto l61
				}
				t52 := int32(load32(m.memory[int64(uint32(v2))+52:]))
				t53 := v6
				v1 = t52
				v0 = t53 + v1
				t54 := int32(load32(m.memory[int64(uint32(v2))+72:]))
				v9 = t54 - v1
				goto l62
			}
		l61:
			v0 = i32(0)
			{
				t55 := int32(m.memory[int64(uint32(v2))+61])
				if t55 == 0 {
					goto l63
				}
				goto l62
			}
		l63:
			{
				{
					t56 := int32(m.memory[int64(uint32(v2))+60])
					if t56 != i32(1) {
						goto l64
					}
					t57 := int32(load32(m.memory[int64(uint32(v2))+56:]))
					v6 = t57
					t58 := int32(load32(m.memory[int64(uint32(v2))+52:]))
					v1 = t58
					goto l65
				}
			l64:
				t59 := int32(load32(m.memory[int64(uint32(v2))+56:]))
				v6 = t59
				t60 := int32(load32(m.memory[int64(uint32(v2))+52:]))
				t61 := v6
				v1 = t60
				if t61 == v1 {
					goto l62
				}
			}
		l65:
			t62 := int32(load32(m.memory[int64(uint32(v2))+28:]))
			v0 = t62 + v1
			v9 = v6 - v1
		}
	l62:
		v1 = i32(3)
		p63 := i32(0)
		if v0 != 0 {
			p63 = v9
		}
		v9 = p63
		{
			{
				t64 := m.fn579(v4, v7, i32(1075992), i32(3))
				if t64 != i32(3) {
					goto l66
				}
				if v9 == 0 {
					goto l67
				}
				p65 := i32(1)
				if v0 != 0 {
					p65 = v0
				}
				t66 := m.fn579(p65, v9, i32(1076004), i32(1))
				if t66&i32(0xffffffe) == i32(2) {
					goto l67
				}
				goto l1
			}
		l66:
			if v9 != 0 {
				goto l1
			}
			t67 := m.fn579(v4, v7, i32(1076004), i32(1))
			if t67&i32(0xffffffe) != i32(2) {
				goto l1
			}
		}
	l67:
		v1 = i32(2)
	}
l1:
	m.g0 = v2 + i32(80)
	return v1
}
func (m *Module) fn539(v0 int32) int32 {
	var v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11 int32
	{
		{
			t0 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			v1 = t0
			if v1 != 0 {
				goto l0
			}
			v2 = i32(0)
			goto l1
		}
	l0:
		t1 := int32(load32(m.memory[uint32(v0):]))
		v3 = t1
		t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		v4 = t2
	l8:
		{
			v2 = v3
			if v2 != v4 {
				goto l2
			}
			v2 = i32(0)
			goto l1
		l2:
			t3 := v0
			v3 = v2 + i32(12)
			store32(m.memory[uint32(t3):], uint32(v3))
			t4 := v0
			v1 = v1 + i32(-1)
			store32(m.memory[int64(uint32(t4))+8:], uint32(v1))
			t5 := int32(load32(m.memory[uint32(v2+i32(4)):]))
			v5 = t5
			{
				{
					t6 := int32(load32(m.memory[uint32(v2+i32(8)):]))
					v2 = t6
					if uint32(v2) > uint32(i32(15)) {
						goto l3
					}
					v6 = i32(2)
					v7 = i32(0)
					if v2 == 0 {
						goto l4
					}
					v8 = v2 & i32(3)
					if uint32(v2) < uint32(i32(4)) {
						goto l7
					}
					v9 = v2 & i32(12)
					v10 = i32(0)
					v11 = i32(0)
				l6:
					{
						t7 := v11
						v2 = v5 + v10
						t8 := int32(int8(m.memory[uint32(v2)]))
						var p9 int32
						if t8 > i32(-65) {
							p9 = 1
						}
						t10 := int32(int8(m.memory[uint32(v2+i32(1))]))
						t11 := t7 + p9
						var p12 int32
						if t10 > i32(-65) {
							p12 = 1
						}
						t13 := int32(int8(m.memory[uint32(v2+i32(2))]))
						t14 := t11 + p12
						var p15 int32
						if t13 > i32(-65) {
							p15 = 1
						}
						t16 := int32(int8(m.memory[uint32(v2+i32(3))]))
						t17 := t14 + p15
						var p18 int32
						if t16 > i32(-65) {
							p18 = 1
						}
						v11 = t17 + p18
						t19 := v9
						v10 = v10 + i32(4)
						if t19 != v10 {
							goto l6
						}
					}
					if v8 != 0 {
						goto l7
					}
					v7 = i32(0)
					goto l4
				l7:
					v8 = v8 + i32(-1)
					if v8 != 0 {
						goto l7
					}
					goto l4
				}
			l3:
				t20 := m.fn580(v5, v2)
				var p21 int32
				if uint32(t20) > uint32(i32(64)) {
					p21 = 1
				}
				v7 = p21
				p22 := i32(2)
				if v7 != 0 {
					p22 = i32(1)
				}
				v6 = p22
			}
		l4:
			p23 := v7
			if v1 != 0 {
				p23 = v6
			}
			v2 = p23
			if v2 == i32(2) {
				goto l8
			}
		}
	}
l1:
	return v2 & i32(1)
}
func (m *Module) fn540(v0 int32) {
	var v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11 int32
	t0 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	v1 = t0
	{
		t1 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		v2 = t1
		if v2 == 0 {
			goto l0
		}
		v3 = i32(0)
	l11:
		{
			v4 = v1 + v3*i32(12)
			t2 := int32(load32(m.memory[int64(uint32(v4))+4:]))
			v5 = t2
			{
				t3 := int32(load32(m.memory[int64(uint32(v4))+8:]))
				v6 = t3
				if v6 == 0 {
					goto l1
				}
				v7 = v5
			l6:
				{
					t4 := int32(load32(m.memory[uint32(v7):]))
					v8 = t4
					if v8 == 0 {
						goto l2
					}
					t5 := int32(load32(m.memory[uint32(v7+i32(4)):]))
					v9 = t5
					t6 := int32(load32(m.memory[uint32(v9+i32(-4)):]))
					v10 = t6
					v11 = v10 & i32(-8)
					t7 := v11
					v10 = v10 & i32(3)
					p8 := i32(8)
					if v10 != 0 {
						p8 = i32(4)
					}
					if uint32(t7) < uint32(p8+v8) {
						m.fn3(i32(1273840), i32(46), i32(1273888))
						panic("unreachable")
					}
					if v10 == 0 {
						goto l4
					}
					if uint32(v11) > uint32(v8+i32(39)) {
						m.fn3(i32(1273904), i32(46), i32(1273952))
						panic("unreachable")
					}
				l4:
					m.fn1(v9)
				}
			l2:
				v7 = v7 + i32(12)
				v6 = v6 + i32(-1)
				if v6 != 0 {
					goto l6
				}
			}
		l1:
			{
				t9 := int32(load32(m.memory[uint32(v4):]))
				v7 = t9
				if v7 == 0 {
					goto l7
				}
				t10 := int32(load32(m.memory[uint32(v5+i32(-4)):]))
				v6 = t10
				v8 = v6 & i32(-8)
				t11 := v8
				v6 = v6 & i32(3)
				p12 := i32(8)
				if v6 != 0 {
					p12 = i32(4)
				}
				v7 = v7 * i32(12)
				if uint32(t11) < uint32(p12+v7) {
					m.fn3(i32(1273840), i32(46), i32(1273888))
					panic("unreachable")
				}
				if v6 == 0 {
					goto l9
				}
				if uint32(v8) > uint32(v7+i32(39)) {
					m.fn3(i32(1273904), i32(46), i32(1273952))
					panic("unreachable")
				}
			l9:
				m.fn1(v5)
			}
		l7:
			v3 = v3 + i32(1)
			if v3 != v2 {
				goto l11
			}
		}
	}
l0:
	{
		t13 := int32(load32(m.memory[uint32(v0):]))
		v7 = t13
		if v7 == 0 {
			return
		}
		t14 := int32(load32(m.memory[uint32(v1+i32(-4)):]))
		v6 = t14
		v8 = v6 & i32(-8)
		t15 := v8
		v6 = v6 & i32(3)
		p16 := i32(8)
		if v6 != 0 {
			p16 = i32(4)
		}
		v7 = v7 * i32(12)
		if uint32(t15) < uint32(p16+v7) {
			m.fn3(i32(1273840), i32(46), i32(1273888))
			panic("unreachable")
		}
		if v6 == 0 {
			goto l14
		}
		if uint32(v8) > uint32(v7+i32(39)) {
			m.fn3(i32(1273904), i32(46), i32(1273952))
			panic("unreachable")
		}
	l14:
		m.fn1(v1)
	}
}
func (m *Module) fn541(v0, v1, v2, v3 int32) {
	var v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14, v15, v16, v17, v18, v19, v20, v21, v22, v23, v24 int32
	t0 := m.g0
	v4 = t0 - i32(16)
	m.g0 = v4
	{
		{
			if uint32(v2) < uint32(i32(16)) {
				t30 := m.fn5(i32(17))
				v5 = t30
				if v5 == 0 {
					m.fn10(i32(1), i32(17))
					panic("unreachable")
				}
				store64(m.memory[int64(uint32(v0))+8:], uint64(i64(-0xffffffef)))
				store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
				store32(m.memory[uint32(v0):], uint32(i32(17)))
				t31 := int32(m.memory[int64(uint32(i32(0)))+1076084])
				m.memory[int64(uint32(v5))+16] = byte(t31)
				t32 := int64(load64(m.memory[int64(uint32(i32(0)))+1076076:]))
				store64(m.memory[int64(uint32(v5))+8:], uint64(t32))
				t33 := int64(load64(m.memory[int64(uint32(i32(0)))+1076068:]))
				store64(m.memory[uint32(v5):], uint64(t33))
				goto l26
			}
			v5 = v2 + i32(-4)
			t1 := int32(uint32(v5) / uint32(i32(12)))
			v6 = t1
			if uint32(v2) >= uint32(i32(0x4000000c)) {
				m.fn9()
				panic("unreachable")
			}
			{
				v7 = v6 * i32(24)
				t2 := m.fn5(v7)
				v8 = t2
				if v8 == 0 {
					m.fn10(i32(4), v7)
					panic("unreachable")
				}
				v9 = i32(0)
				store32(m.memory[int64(uint32(v4))+12:], uint32(i32(0)))
				store32(m.memory[int64(uint32(v4))+8:], uint32(v8))
				v10 = int32(uint32(v5) >> 2)
				v11 = v2 + i32(-1)
				v12 = int32(uint32(v2+i32(4)) >> 2)
				store32(m.memory[int64(uint32(v4))+4:], uint32(v6))
				t3 := v2
				v13 = v6 << 2
				v14 = t3 - v13 + i32(-6)
				v15 = i32(0)
				v5 = i32(0)
				{
				l25:
					{
						{
							{
								{
									if v12 == v5 {
										v7 = i32(6)
										t8 := m.fn5(i32(6))
										v5 = t8
										if v5 == 0 {
											m.fn10(i32(1), i32(6))
											panic("unreachable")
										}
										t9 := int32(load16(m.memory[int64(uint32(i32(0)))+1070708:]))
										store16(m.memory[int64(uint32(v5))+4:], uint16(t9))
										t10 := int32(load32(m.memory[int64(uint32(i32(0)))+1070704:]))
										store32(m.memory[uint32(v5):], uint32(t10))
										goto l9
									}
									if v10 == v5 {
										v7 = i32(6)
										t11 := m.fn5(i32(6))
										v5 = t11
										if v5 == 0 {
											m.fn10(i32(1), i32(6))
											panic("unreachable")
										}
										t12 := int32(load16(m.memory[int64(uint32(i32(0)))+1070708:]))
										store16(m.memory[int64(uint32(v5))+4:], uint16(t12))
										t13 := int32(load32(m.memory[int64(uint32(i32(0)))+1070704:]))
										store32(m.memory[uint32(v5):], uint32(t13))
										goto l9
									}
									if uint32(v2) < uint32(v13+i32(6)) {
										goto l5
									}
									if uint32(v14) <= uint32(i32(3)) {
										goto l5
									}
									v7 = v1 + v9
									t4 := int32(load32(m.memory[uint32(v7):]))
									v16 = t4
									t5 := int32(load32(m.memory[uint32(v7+i32(4)):]))
									v17 = t5
									v18 = v1 + v13
									t6 := int32(load32(m.memory[uint32(v18+i32(6)):]))
									v7 = t6
									if uint32(v13+i32(10)) < uint32(v11) {
										t17 := v7 & i32(0x3fffffff)
										v19 = int32(uint32(v7)>>30) & i32(1)
										v20 = i32_shr_u(t17, v19)
										t18 := int32(load16(m.memory[uint32(v18+i32(10)):]))
										v7 = t18
										if v7&i32(1) == 0 {
											goto l12
										}
										t19 := int32(load32(m.memory[int64(uint32(v3))+8:]))
										v18 = int32(uint32(v7) >> 1)
										var p20 int32
										if uint32(t19) > uint32(v18) {
											p20 = 1
										}
										v21 = p20
										goto l7
									}
									t7 := v7 & i32(0x3fffffff)
									v19 = int32(uint32(v7)>>30) & i32(1)
									v20 = i32_shr_u(t7, v19)
									v21 = i32(0)
									goto l7
								}
							l5:
								v7 = i32(7)
								t14 := m.fn5(i32(7))
								v5 = t14
								if v5 == 0 {
									m.fn10(i32(1), i32(7))
									panic("unreachable")
								}
								t15 := int32(load32(m.memory[int64(uint32(i32(0)))+1070713:]))
								store32(m.memory[int64(uint32(v5))+3:], uint32(t15))
								t16 := int32(load32(m.memory[int64(uint32(i32(0)))+1070710:]))
								store32(m.memory[uint32(v5):], uint32(t16))
								goto l9
							}
						l12:
							v21 = i32(0)
							if v7 != 0 {
								goto l13
							}
							goto l7
						l13:
							v22 = i32(38)
							v23 = int32(uint32(v7)>>1) & i32(127)
							switch v23 + i32(-85) {
							case 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34:
								goto l7
							default:
								v24 = i32(10)
								switch v23 + i32(-12) {
								case 0:
									goto l19
								default:
									goto l7
								case 12:
									v22 = i32(36)
									v24 = i32(22)
									goto l19
								case 13:
									v22 = i32(36)
									v24 = i32(23)
									goto l19
								}
							case 0:
								v22 = i32(8)
								v24 = i32(53)
								goto l19
							case 1:
								v22 = i32(8)
								v24 = i32(54)
								goto l19
							case 2:
								v22 = i32(8)
								v24 = i32(55)
								goto l19
							case 35:
								v24 = i32(64)
							}
						l19:
							t21 := m.fn5(i32(3))
							v23 = t21
							if v23 == 0 {
								m.fn24(i32(1), i32(3))
								panic("unreachable")
							}
							m.memory[int64(uint32(v23))+2] = byte(int32(uint32(v7) >> 8))
							m.memory[int64(uint32(v23))+1] = byte(v22)
							m.memory[uint32(v23)] = byte(v24)
							{
								t22 := int32(load32(m.memory[int64(uint32(v3))+8:]))
								v18 = t22
								t23 := int32(load32(m.memory[uint32(v3):]))
								if v18 != t23 {
									goto l23
								}
								m.fn316(v3)
							}
						l23:
							v21 = i32(1)
							store32(m.memory[int64(uint32(v3))+8:], uint32(v18+i32(1)))
							t24 := int32(load32(m.memory[int64(uint32(v3))+4:]))
							v7 = t24 + v18*i32(12)
							store32(m.memory[int64(uint32(v7))+8:], uint32(i32(3)))
							store32(m.memory[int64(uint32(v7))+4:], uint32(v23))
							store32(m.memory[uint32(v7):], uint32(i32(3)))
						}
					l7:
						{
							t25 := int32(load32(m.memory[int64(uint32(v4))+4:]))
							if v5 != t25 {
								goto l24
							}
							m.fn326(v4 + i32(4))
							t26 := int32(load32(m.memory[int64(uint32(v4))+8:]))
							v8 = t26
						}
					l24:
						v7 = v8 + v15
						store32(m.memory[uint32(v7):], uint32(v21))
						m.memory[uint32(v7+i32(20))] = byte(v19)
						store32(m.memory[uint32(v7+i32(16)):], uint32(v20))
						store32(m.memory[uint32(v7+i32(12)):], uint32(v17))
						store32(m.memory[uint32(v7+i32(8)):], uint32(v16))
						store32(m.memory[uint32(v7+i32(4)):], uint32(v18))
						t27 := v4
						v5 = v5 + i32(1)
						store32(m.memory[int64(uint32(t27))+12:], uint32(v5))
						v9 = v9 + i32(4)
						v15 = v15 + i32(24)
						v14 = v14 + i32(-8)
						v13 = v13 + i32(8)
						if v6 != v5 {
							goto l25
						}
					}
					t28 := int32(load32(m.memory[int64(uint32(v4))+12:]))
					store32(m.memory[int64(uint32(v0))+12:], uint32(t28))
					t29 := int64(load64(m.memory[int64(uint32(v4))+4:]))
					store64(m.memory[int64(uint32(v0))+4:], uint64(t29))
					store32(m.memory[uint32(v0):], uint32(i32(-1)))
					goto l26
				}
			}
		}
	l9:
		store32(m.memory[int64(uint32(v0))+12:], uint32(i32(-1)))
		store32(m.memory[int64(uint32(v0))+8:], uint32(v7))
		store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
		store32(m.memory[uint32(v0):], uint32(v7))
		t34 := int32(load32(m.memory[int64(uint32(v4))+4:]))
		v5 = t34
		if v5 == 0 {
			goto l26
		}
		t35 := int32(load32(m.memory[uint32(v8+i32(-4)):]))
		v7 = t35
		v13 = v7 & i32(-8)
		t36 := v13
		v7 = v7 & i32(3)
		p37 := i32(8)
		if v7 != 0 {
			p37 = i32(4)
		}
		v5 = v5 * i32(24)
		if uint32(t36) < uint32(p37+v5) {
			m.fn3(i32(1273840), i32(46), i32(1273888))
			panic("unreachable")
		}
		if v7 == 0 {
			goto l29
		}
		if uint32(v13) > uint32(v5+i32(39)) {
			m.fn3(i32(1273904), i32(46), i32(1273952))
			panic("unreachable")
		}
	l29:
		m.fn1(v8)
	}
l26:
	m.g0 = v4 + i32(16)
}
func (m *Module) fn542(v0 int32) {
	var v1 int32
	{
		t0 := m.fn5(i32(20))
		v1 = t0
		if v1 != 0 {
			goto l0
		}
		m.fn10(i32(1), i32(20))
		panic("unreachable")
	}
l0:
	store64(m.memory[int64(uint32(v0))+8:], uint64(i64(-0xffffffec)))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
	store32(m.memory[uint32(v0):], uint32(i32(20)))
	t1 := int32(load32(m.memory[int64(uint32(i32(0)))+1070753:]))
	store32(m.memory[int64(uint32(v1))+16:], uint32(t1))
	t2 := int64(load64(m.memory[int64(uint32(i32(0)))+1070745:]))
	store64(m.memory[int64(uint32(v1))+8:], uint64(t2))
	t3 := int64(load64(m.memory[int64(uint32(i32(0)))+1070737:]))
	store64(m.memory[uint32(v1):], uint64(t3))
}
func (m *Module) fn543(v0, v1, v2, v3, v4, v5, v6, v7, v8 int32) {
	var v9, v10, v11, v12, v13, v14, v15, v16, v17, v18, v19, v20, v21, v22, v23 int32
	t0 := m.g0
	v9 = t0 - i32(80)
	m.g0 = v9
	v10 = i32(0)
	store32(m.memory[int64(uint32(v9))+12:], uint32(i32(0)))
	store64(m.memory[int64(uint32(v9))+4:], uint64(i64(0x400000000)))
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
		v13 = int32(uint32(v12+i32(-4)) >> 3)
		if v13 == 0 {
			goto l2
		}
		v14 = v3 + v10
		p5 := i32(1)
		if v6 != 0 {
			p5 = i32(13)
		}
		v15 = p5
		v16 = i32(4)
		v17 = i32(0)
		v18 = i32(0)
	l14:
		{
			t6 := v12
			v10 = (v18+v13)<<2 + i32(4)
			if uint32(t6) < uint32(v10) {
				goto l3
			}
			if uint32(v12-v10) <= uint32(i32(3)) {
				goto l3
			}
			t7 := int32(load32(m.memory[uint32(v14+v10):]))
			t8 := v2
			v10 = t7 << 9 & i32(0x7ffffe00)
			if uint32(t8) < uint32(v10) {
				goto l3
			}
			if uint32(v2-v10) < uint32(i32(512)) {
				goto l3
			}
			v19 = v1 + v10
			t9 := int32(m.memory[int64(uint32(v19))+511])
			v10 = t9
			if v10 == 0 {
				goto l3
			}
			v3 = v10 << 2
			v11 = v3 + i32(4)
			v5 = i32(0)
			v10 = i32(0)
		l13:
			{
				{
					if uint32(v10) > uint32(i32(127)) {
						goto l4
					}
					if v5 == i32(508) {
						goto l4
					}
					if uint32(v11) > uint32(i32(511)) {
						goto l4
					}
					v4 = v19 + v5
					t10 := int32(load32(m.memory[uint32(v4):]))
					v20 = t10
					t11 := int32(load32(m.memory[uint32(v4+i32(4)):]))
					v21 = t11
					store16(m.memory[int64(uint32(v9))+76:], uint16(i32(0)))
					store64(m.memory[int64(uint32(v9))+68:], uint64(i64(1)))
					store64(m.memory[int64(uint32(v9))+60:], uint64(i64(33686018)))
					m.memory[int64(uint32(v9))+58] = byte(i32(2))
					m.memory[int64(uint32(v9))+56] = byte(i32(0))
					store16(m.memory[int64(uint32(v9))+52:], uint16(i32(0)))
					store32(m.memory[int64(uint32(v9))+24:], uint32(i32(-1)))
					store32(m.memory[int64(uint32(v9))+16:], uint32(i32(0)))
					{
						t12 := int32(m.memory[uint32(v19+v11)])
						v4 = t12
						if v4 == 0 {
							goto l5
						}
						t13 := v19
						v22 = v4 << 1
						v23 = t13 + v22
						t14 := int32(m.memory[uint32(v23)])
						v4 = t14
						if v6 != 0 {
							{
								if v4 == 0 {
									goto l10
								}
								v22 = v22 | i32(1)
								v4 = v4<<1 + i32(-1)
								goto l11
							l10:
								v22 = v22 + i32(2)
								t16 := int32(m.memory[int64(uint32(v23))+1])
								v4 = t16 << 1
							}
						l11:
							if uint32(v4) > uint32(i32(512)-v22) {
								goto l5
							}
							if uint32(v4) < uint32(i32(2)) {
								goto l5
							}
							t17 := v9
							v22 = v19 + v22
							t18 := int32(load16(m.memory[uint32(v22):]))
							store16(m.memory[int64(uint32(t17))+76:], uint16(t18))
							m.fn546(v22+i32(2), v4+i32(-2), v7, v8, v9+i32(16))
							goto l5
						}
						if uint32(v22^i32(511)) < uint32(v4) {
							goto l5
						}
						{
							if v4 != 0 {
								goto l7
							}
							v22 = i32(1)
							goto l8
						l7:
							t15 := m.fn5(v4)
							v22 = t15
							if v22 == 0 {
								m.fn10(i32(1), v4)
								panic("unreachable")
							}
							if v4 == 0 {
								goto l8
							}
							memory_copy(m.memory, uint32(v22), uint32(v23+i32(1)), uint32(v4))
						}
					l8:
						store32(m.memory[int64(uint32(v9))+72:], uint32(v4))
						store32(m.memory[int64(uint32(v9))+68:], uint32(v22))
						store32(m.memory[int64(uint32(v9))+64:], uint32(v4))
						goto l5
					}
				l5:
					{
						t19 := int32(load32(m.memory[int64(uint32(v9))+4:]))
						if v17 != t19 {
							goto l12
						}
						m.fn325(v9 + i32(4))
						t20 := int32(load32(m.memory[int64(uint32(v9))+8:]))
						v16 = t20
					}
				l12:
					v4 = v16 + v17*i32(72)
					t21 := int64(load64(m.memory[int64(uint32(v9))+72:]))
					store64(m.memory[int64(uint32(v4))+56:], uint64(t21))
					t22 := int64(load64(m.memory[int64(uint32(v9))+64:]))
					store64(m.memory[int64(uint32(v4))+48:], uint64(t22))
					t23 := int64(load64(m.memory[int64(uint32(v9))+56:]))
					store64(m.memory[int64(uint32(v4))+40:], uint64(t23))
					t24 := int64(load64(m.memory[int64(uint32(v9))+48:]))
					store64(m.memory[int64(uint32(v4))+32:], uint64(t24))
					t25 := int64(load64(m.memory[int64(uint32(v9))+40:]))
					store64(m.memory[int64(uint32(v4))+24:], uint64(t25))
					t26 := int64(load64(m.memory[int64(uint32(v9))+32:]))
					store64(m.memory[int64(uint32(v4))+16:], uint64(t26))
					t27 := int64(load64(m.memory[int64(uint32(v9))+24:]))
					store64(m.memory[int64(uint32(v4))+8:], uint64(t27))
					t28 := int64(load64(m.memory[int64(uint32(v9))+16:]))
					store64(m.memory[uint32(v4):], uint64(t28))
					store32(m.memory[int64(uint32(v4))+68:], uint32(v21))
					store32(m.memory[int64(uint32(v4))+64:], uint32(v20))
					t29 := v9
					v17 = v17 + i32(1)
					store32(m.memory[int64(uint32(t29))+12:], uint32(v17))
				}
			l4:
				v10 = v10 + i32(1)
				v11 = v11 + v15
				t30 := v3
				v5 = v5 + i32(4)
				if t30 != v5 {
					goto l13
				}
			}
		}
	l3:
		v18 = v18 + i32(1)
		if v18 != v13 {
			goto l14
		}
	}
l2:
	t31 := int32(load32(m.memory[int64(uint32(v9))+12:]))
	store32(m.memory[int64(uint32(v0))+8:], uint32(t31))
	t32 := int64(load64(m.memory[int64(uint32(v9))+4:]))
	store64(m.memory[uint32(v0):], uint64(t32))
	m.g0 = v9 + i32(80)
}
func (m *Module) fn544(v0 int32) {
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
	m.fn805(t2, t4, t3, v2, i32(4), i32(8))
	{
		t6 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		if t6 != i32(1) {
			goto l0
		}
		t7 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		t8 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		m.fn10(t7, t8)
		panic("unreachable")
	}
l0:
	t9 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	v3 = t9
	store32(m.memory[uint32(v0):], uint32(v2))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
	m.g0 = v1 + i32(16)
}
func (m *Module) fn545(v0, v1 int32) {
	var v2, v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14, v15 int32
	var v16 int64
	var v17, v18, v19 int32
	v2 = i32(-1)
	t0 := int32(m.memory[int64(uint32(v1))+47])
	v3 = t0
	t1 := int32(m.memory[int64(uint32(v1))+46])
	v4 = t1
	t2 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	v5 = t2
	t3 := int32(load32(m.memory[uint32(v1):]))
	v6 = t3
	t4 := int32(m.memory[int64(uint32(v1))+41])
	v7 = t4
	t5 := int32(m.memory[int64(uint32(v1))+40])
	v8 = t5
	t6 := int32(load16(m.memory[int64(uint32(v1))+38:]))
	v9 = t6
	t7 := int32(load16(m.memory[int64(uint32(v1))+36:]))
	v10 = t7
	t8 := int32(m.memory[int64(uint32(v1))+43])
	v11 = t8
	t9 := int32(m.memory[int64(uint32(v1))+42])
	v12 = t9
	t10 := int32(m.memory[int64(uint32(v1))+45])
	v13 = t10
	t11 := int32(m.memory[int64(uint32(v1))+44])
	v14 = t11
	{
		{
			t12 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			if t12 != i32(-1) {
				goto l0
			}
			goto l1
		}
	l0:
		v15 = i32(1)
		{
			{
				t13 := int32(load32(m.memory[int64(uint32(v1))+16:]))
				v2 = t13
				if v2 != 0 {
					goto l2
				}
				v16 = i64(2)
				goto l3
			}
		l2:
			t14 := int32(load32(m.memory[int64(uint32(v1))+12:]))
			v17 = t14
			v18 = v2 << 1
			t15 := m.fn5(v18)
			v19 = t15
			if v19 == 0 {
				m.fn10(i32(2), v18)
				panic("unreachable")
			}
			if v18 == 0 {
				goto l5
			}
			memory_copy(m.memory, uint32(v19), uint32(v17), uint32(v18))
		l5:
			v16 = int64(uint32(v19))
		}
	l3:
		{
			t16 := int32(load32(m.memory[int64(uint32(v1))+28:]))
			v18 = t16
			if v18 == 0 {
				goto l6
			}
			t17 := int32(load32(m.memory[int64(uint32(v1))+24:]))
			v17 = t17
			v19 = v18 << 2
			t18 := m.fn5(v19)
			v15 = t18
			if v15 == 0 {
				m.fn10(i32(1), v19)
				panic("unreachable")
			}
			if v19 == 0 {
				goto l6
			}
			memory_copy(m.memory, uint32(v15), uint32(v17), uint32(v19))
		}
	l6:
		v16 = v16 | int64(uint32(v2))<<32
		t19 := int32(m.memory[int64(uint32(v1))+32])
		v1 = t19
	}
l1:
	m.memory[int64(uint32(v0))+45] = byte(v13)
	m.memory[int64(uint32(v0))+44] = byte(v14)
	m.memory[int64(uint32(v0))+43] = byte(v11)
	m.memory[int64(uint32(v0))+42] = byte(v12)
	m.memory[int64(uint32(v0))+41] = byte(v7)
	m.memory[int64(uint32(v0))+40] = byte(v8)
	store16(m.memory[int64(uint32(v0))+38:], uint16(v9))
	store16(m.memory[int64(uint32(v0))+36:], uint16(v10))
	m.memory[int64(uint32(v0))+47] = byte(v3)
	m.memory[int64(uint32(v0))+46] = byte(v4)
	store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
	store32(m.memory[uint32(v0):], uint32(v6))
	m.memory[int64(uint32(v0))+32] = byte(v1)
	store32(m.memory[int64(uint32(v0))+28:], uint32(v18))
	store32(m.memory[int64(uint32(v0))+24:], uint32(v15))
	store32(m.memory[int64(uint32(v0))+20:], uint32(v18))
	store64(m.memory[int64(uint32(v0))+12:], uint64(v16))
	store32(m.memory[int64(uint32(v0))+8:], uint32(v2))
}
func (m *Module) fn546(v0, v1, v2, v3, v4 int32) {
	var v5, v6, v7, v8, v9, v10, v11, v12, v13, v14, v15, v16, v17 int32
	var v18 int64
	t0 := m.g0
	v5 = t0 - i32(16)
	m.g0 = v5
	if uint32(v1) < uint32(i32(2)) {
		goto l0
	}
	v6 = i32(2)
	v7 = i32(0)
l57:
	{
		{
			{
				if uint32(v7) >= uint32(v1) {
					m.fn33(v7, v1, i32(1069384))
					panic("unreachable")
				}
				v8 = v7 + i32(1)
				if uint32(v8) >= uint32(v1) {
					m.fn33(v8, v1, i32(1069400))
					panic("unreachable")
				}
				t1 := int32(m.memory[uint32(v0+v8)])
				v9 = t1
				t2 := int32(m.memory[uint32(v0+v7)])
				v8 = v9<<8 | t2
				v10 = v0 + v6
				v11 = v1 - v6
				v7 = i32(1)
				switch int32(uint32(v9) >> 5) {
				default:
					goto l3
				case 2, 4, 5:
					v7 = i32(2)
					goto l3
				case 3:
					v7 = i32(4)
					goto l3
				case 7:
					v7 = i32(3)
					goto l3
				case 6:
					if v8 == i32(54792) {
						goto l8
					}
					v7 = i32(0)
					if v1 == v6 {
						goto l9
					}
					t3 := int32(m.memory[uint32(v10)])
					v7 = t3 + i32(1)
					goto l3
				}
			}
		l8:
			if uint32(v11) > uint32(i32(1)) {
				goto l10
			}
			v7 = i32(0)
			goto l11
		l10:
			t4 := int32(load16(m.memory[uint32(v10):]))
			v7 = t4 + i32(1)
		}
	l3:
		if uint32(v7) > uint32(v11) {
			goto l0
		}
	l9:
		if v8 > i32(13315) {
			if v8 > i32(26181) {
				switch v8 + i32(-26182) {
				case 0:
					if uint32(v7) <= uint32(i32(3)) {
						goto l11
					}
					t15 := int32(load32(m.memory[uint32(v10):]))
					t16 := v3
					v8 = t15
					if uint32(t16) < uint32(v8) {
						goto l11
					}
					v9 = v3 - v8
					if uint32(v9) < uint32(i32(2)) {
						goto l11
					}
					t17 := v9 + i32(-2)
					v8 = v2 + v8
					t18 := int32(load16(m.memory[uint32(v8):]))
					v9 = t18
					if uint32(t17) < uint32(v9) {
						goto l11
					}
					m.fn546(v8+i32(2), v9, i32(1), i32(0), v4)
					goto l11
				case 1, 2:
					goto l11
				case 3:
					{
						if uint32(v7) < uint32(i32(4)) {
							goto l31
						}
						t19 := int32(load32(m.memory[uint32(v10):]))
						v8 = t19
					}
				l31:
					store32(m.memory[int64(uint32(v4))+4:], uint32(v8))
					t20 := v4
					var p21 int32
					if uint32(v7) > uint32(i32(3)) {
						p21 = 1
					}
					store32(m.memory[uint32(t20):], uint32(p21))
					goto l11
				case 4:
					if uint32(v7) <= uint32(i32(3)) {
						goto l11
					}
					t22 := int32(load32(m.memory[uint32(v10):]))
					v8 = t22
					{
						t23 := int32(load32(m.memory[uint32(v4):]))
						if t23 != i32(1) {
							goto l32
						}
						t24 := int32(load32(m.memory[int64(uint32(v4))+4:]))
						v11 = t24
						v9 = v11 + v8
						t25 := v9>>31 ^ i32(-0x80000000)
						t26 := v9
						var p27 int32
						if v8 < i32(0) {
							p27 = 1
						}
						var p28 int32
						if v9 < v11 {
							p28 = 1
						}
						p29 := t26
						if p27^p28 != 0 {
							p29 = t25
						}
						v8 = p29
					}
				l32:
					store32(m.memory[int64(uint32(v4))+4:], uint32(v8))
					store32(m.memory[uint32(v4):], uint32(i32(1)))
					goto l11
				default:
					if v8 == i32(54792) {
						if uint32(v7) < uint32(i32(3)) {
							goto l11
						}
						t30 := int32(m.memory[int64(uint32(v10))+2])
						v11 = t30
						if uint32(v11) > uint32(i32(63)) {
							goto l11
						}
						v8 = v11 + i32(1)
						v12 = v8 << 1
						t31 := m.fn5(v12)
						v13 = t31
						if v13 == 0 {
							m.fn10(i32(2), v12)
							panic("unreachable")
						}
						v14 = i32(0)
						store32(m.memory[int64(uint32(v5))+12:], uint32(i32(0)))
						store32(m.memory[int64(uint32(v5))+8:], uint32(v13))
						store32(m.memory[int64(uint32(v5))+4:], uint32(v8))
						v9 = i32(1)
						v8 = i32(0)
					l43:
						{
							{
								t32 := v7
								v15 = v8<<1 + i32(3)
								if uint32(t32) < uint32(v15) {
									goto l36
								}
								if uint32(v7-v15) <= uint32(i32(1)) {
									goto l36
								}
								t33 := int32(load16(m.memory[uint32(v10+v15):]))
								v15 = t33
								{
									t34 := int32(load32(m.memory[int64(uint32(v5))+4:]))
									if v9+i32(-1) != t34 {
										goto l37
									}
									m.fn331(v5 + i32(4))
									t35 := int32(load32(m.memory[int64(uint32(v5))+8:]))
									v13 = t35
								}
							l37:
								store16(m.memory[uint32(v13+v14):], uint16(v15))
								store32(m.memory[int64(uint32(v5))+12:], uint32(v9))
								if uint32(v8) < uint32(v11) {
									v14 = v14 + i32(2)
									v9 = v9 + i32(1)
									t40 := v8
									var p41 int32
									if uint32(v8) < uint32(v11) {
										p41 = 1
									}
									v8 = t40 + p41
									if uint32(v8) <= uint32(v11) {
										goto l43
									}
									goto l39
								}
								goto l39
							}
						l36:
							t36 := int32(load32(m.memory[int64(uint32(v5))+4:]))
							v8 = t36
							if v8 == 0 {
								goto l11
							}
							t37 := int32(load32(m.memory[uint32(v13+i32(-4)):]))
							v9 = t37
							v11 = v9 & i32(-8)
							t38 := v11
							v9 = v9 & i32(3)
							p39 := i32(8)
							if v9 != 0 {
								p39 = i32(4)
							}
							v8 = v8 << 1
							if uint32(t38) < uint32(p39+v8) {
								m.fn3(i32(1273840), i32(46), i32(1273888))
								panic("unreachable")
							}
							if v9 == 0 {
								goto l41
							}
							if uint32(v11) > uint32(v8+i32(39)) {
								m.fn3(i32(1273904), i32(46), i32(1273952))
								panic("unreachable")
							}
						l41:
							m.fn1(v13)
							goto l11
						}
					}
					if v8 != i32(54827) {
						goto l11
					}
					if uint32(v7) < uint32(i32(3)) {
						goto l11
					}
					t6 := int32(load32(m.memory[int64(uint32(v4))+8:]))
					if t6 == i32(-1) {
						goto l11
					}
					t7 := int32(load32(m.memory[int64(uint32(v4))+28:]))
					t8 := int32(m.memory[int64(uint32(v10))+1])
					v8 = t8
					if uint32(t7) <= uint32(v8) {
						goto l11
					}
					t9 := int32(load32(m.memory[int64(uint32(v4))+24:]))
					v8 = t9 + v8<<2
					t10 := int32(m.memory[int64(uint32(v10))+2])
					t11 := v8
					v9 = t10 & i32(255)
					var p12 int32
					if v9 == i32(1) {
						p12 = 1
					}
					m.memory[int64(uint32(t11))+3] = byte(p12)
					t13 := v8
					var p14 int32
					if v9 == i32(3) {
						p14 = 1
					}
					m.memory[int64(uint32(t13))+2] = byte(p14)
					goto l11
				}
			}
			if v8 == i32(13316) {
				if v7 == 0 {
					v8 = i32(0)
					t64 := int32(load32(m.memory[int64(uint32(v4))+8:]))
					if t64 != i32(-1) {
						goto l47
					}
					v7 = i32(0)
					goto l11
				}
				t61 := int32(m.memory[uint32(v10)])
				v8 = t61
				{
					t62 := int32(load32(m.memory[int64(uint32(v4))+8:]))
					if t62 == i32(-1) {
						if v8 == 0 {
							goto l11
						}
						m.memory[int64(uint32(v4))+32] = byte(i32(1))
						store64(m.memory[int64(uint32(v4))+24:], uint64(i64(1)))
						store64(m.memory[int64(uint32(v4))+16:], uint64(i64(0)))
						store64(m.memory[int64(uint32(v4))+8:], uint64(i64(0x200000000)))
						goto l11
					}
					var p63 int32
					if v8 != i32(0) {
						p63 = 1
					}
					v8 = p63
					goto l47
				}
			}
			if v8 != i32(17931) {
				goto l11
			}
			v8 = i32(0)
			{
				if uint32(v7) < uint32(i32(2)) {
					goto l23
				}
				t5 := int32(load16(m.memory[uint32(v10):]))
				v9 = t5
				v8 = i32(1)
			}
		l23:
			store16(m.memory[int64(uint32(v4))+38:], uint16(v9))
			store16(m.memory[int64(uint32(v4))+36:], uint16(v8))
			goto l11
		}
		if v8 > i32(9737) {
			if v8 == i32(9738) {
				{
					if v7 == 0 {
						goto l44
					}
					t52 := int32(m.memory[uint32(v10)])
					v8 = t52
				}
			l44:
				m.memory[int64(uint32(v4))+41] = byte(v8)
				t53 := v4
				var p54 int32
				if v7 != i32(0) {
					p54 = 1
				}
				m.memory[int64(uint32(t53))+40] = byte(p54)
				goto l11
			}
			if v8 != i32(9792) {
				goto l11
			}
			if v7 != 0 {
				t48 := int32(m.memory[uint32(v10)])
				t49 := v4
				v8 = t48
				m.memory[int64(uint32(t49))+43] = byte(v8 + i32(1))
				t50 := v4
				var p51 int32
				if uint32(v8&i32(255)) < uint32(i32(9)) {
					p51 = 1
				}
				m.memory[int64(uint32(t50))+42] = byte(p51)
				goto l11
			}
			v7 = i32(0)
			goto l11
		}
		switch v8 + i32(-9238) {
		case 0:
			if v7 != 0 {
				t42 := int32(m.memory[uint32(v10)])
				t43 := v4
				var p44 int32
				if t42 != i32(0) {
					p44 = 1
				}
				m.memory[int64(uint32(t43))+44] = byte(p44)
				goto l11
			}
			m.memory[int64(uint32(v4))+44] = byte(i32(0))
			goto l11
		case 1:
			if v7 != 0 {
				t45 := int32(m.memory[uint32(v10)])
				t46 := v4
				var p47 int32
				if t45 != i32(0) {
					p47 = 1
				}
				m.memory[int64(uint32(t46))+45] = byte(p47)
				goto l11
			}
			m.memory[int64(uint32(v4))+45] = byte(i32(0))
			goto l11
		default:
			switch v8 + i32(-9291) {
			case 0:
				if v7 != 0 {
					t55 := int32(m.memory[uint32(v10)])
					t56 := v4
					var p57 int32
					if t55 != i32(0) {
						p57 = 1
					}
					m.memory[int64(uint32(t56))+46] = byte(p57)
					goto l11
				}
				m.memory[int64(uint32(v4))+46] = byte(i32(0))
				goto l11
			case 1:
				if v7 != 0 {
					t58 := int32(m.memory[uint32(v10)])
					t59 := v4
					var p60 int32
					if t58 != i32(0) {
						p60 = 1
					}
					m.memory[int64(uint32(t59))+47] = byte(p60)
					goto l11
				}
				m.memory[int64(uint32(v4))+47] = byte(i32(0))
				goto l11
			default:
				goto l11
			}
		}
	l47:
		m.memory[int64(uint32(v4))+32] = byte(v8)
		goto l11
	l39:
		if v11 != 0 {
			v17 = v11 << 2
			t65 := m.fn5(v17)
			v16 = t65
			if v16 == 0 {
				m.fn10(i32(1), v17)
				panic("unreachable")
			}
			v8 = v16
			if v11 == i32(1) {
				goto l51
			}
			v8 = v17 + i32(-4)
			if v8 == 0 {
				goto l52
			}
			memory_zero(m.memory, uint32(v16), uint32(v8))
		l52:
			v8 = v8 + v16
		l51:
			v14 = i32(0)
			store32(m.memory[uint32(v8):], uint32(i32(0)))
			v8 = v12 + i32(3)
			v15 = v7 - v11<<1 + i32(-5)
		l53:
			{
				if uint32(v7) < uint32(v8) {
					goto l49
				}
				if uint32(v15) < uint32(i32(2)) {
					goto l49
				}
				v9 = v16 + v14
				t66 := int32(load16(m.memory[uint32(v10+v8):]))
				t67 := v9 + i32(1)
				v13 = t66
				v12 = v13 & i32(3)
				var p68 int32
				if v12 == i32(1) {
					p68 = 1
				}
				m.memory[uint32(t67)] = byte(p68)
				t69 := v9
				var p70 int32
				if uint32(v12) > uint32(i32(1)) {
					p70 = 1
				}
				m.memory[uint32(t69)] = byte(p70)
				t71 := v9 + i32(3)
				v13 = int32(uint32(v13)>>5) & i32(3)
				var p72 int32
				if v13 == i32(1) {
					p72 = 1
				}
				m.memory[uint32(t71)] = byte(p72)
				t73 := v9 + i32(2)
				var p74 int32
				if v13 == i32(3) {
					p74 = 1
				}
				m.memory[uint32(t73)] = byte(p74)
				v15 = v15 + i32(-20)
				v8 = v8 + i32(20)
				t75 := v17
				v14 = v14 + i32(4)
				if t75 == v14 {
					goto l49
				}
				goto l53
			}
		}
		v16 = i32(1)
		goto l49
	l49:
		t76 := int32(load32(m.memory[int64(uint32(v5))+4:]))
		v8 = t76
		if v8 == i32(-1) {
			goto l11
		}
		t77 := int64(load64(m.memory[int64(uint32(v5))+8:]))
		v18 = t77
		{
			{
				t78 := int32(load32(m.memory[int64(uint32(v4))+8:]))
				v9 = t78
				if v9 != i32(-1) {
					goto l54
				}
				v10 = i32(0)
				goto l55
			}
		l54:
			t79 := int32(m.memory[int64(uint32(v4))+32])
			v10 = t79
			{
				if v9 == 0 {
					goto l56
				}
				t80 := int32(load32(m.memory[int64(uint32(v4))+12:]))
				m.fn18(t80, v9<<1, i32(2))
			}
		l56:
			t81 := int32(load32(m.memory[int64(uint32(v4))+20:]))
			v9 = t81
			if v9 == 0 {
				goto l55
			}
			t82 := int32(load32(m.memory[int64(uint32(v4))+24:]))
			m.fn18(t82, v9<<2, i32(1))
		}
	l55:
		m.memory[int64(uint32(v4))+32] = byte(v10)
		store32(m.memory[int64(uint32(v4))+28:], uint32(v11))
		store32(m.memory[int64(uint32(v4))+24:], uint32(v16))
		store32(m.memory[int64(uint32(v4))+20:], uint32(v11))
		store64(m.memory[int64(uint32(v4))+12:], uint64(v18))
		store32(m.memory[int64(uint32(v4))+8:], uint32(v8))
	}
l11:
	v7 = v7 + v6
	v6 = v7 + i32(2)
	if uint32(v6) <= uint32(v1) {
		goto l57
	}
l0:
	m.g0 = v5 + i32(16)
}
func (m *Module) fn547(v0, v1, v2 int32) {
	var v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14, v15, v16, v17, v18, v19, v20, v21, v22, v23, v24, v25, v26, v27 int32
	var v28 int64
	var v29, v30, v31, v32 int32
	t0 := int32(load32(m.memory[uint32(v2):]))
	t1 := v2
	t2 := v1
	v3 = t0
	p3 := t2
	if v3 != 0 {
		p3 = t1
	}
	t4 := int32(load32(m.memory[int64(uint32(p3))+4:]))
	v4 = t4
	t5 := int32(m.memory[int64(uint32(v2))+40])
	t6 := v2
	t7 := v1
	v5 = t5
	p8 := t7
	if v5 != 0 {
		p8 = t6
	}
	t9 := int32(m.memory[int64(uint32(p8))+41])
	v6 = t9
	t10 := int32(load16(m.memory[int64(uint32(v2))+36:]))
	t11 := v2
	t12 := v1
	v7 = t10
	p13 := t12
	if v7 != 0 {
		p13 = t11
	}
	t14 := int32(load16(m.memory[int64(uint32(p13))+38:]))
	v8 = t14
	t15 := int32(m.memory[int64(uint32(v2))+42])
	t16 := v1
	t17 := v2
	v9 = t15
	p18 := t17
	if v9 == i32(2) {
		p18 = t16
	}
	t19 := int32(m.memory[int64(uint32(p18))+43])
	v10 = t19
	t20 := int32(load32(m.memory[int64(uint32(v1))+24:]))
	v11 = t20
	t21 := int32(load32(m.memory[int64(uint32(v1))+20:]))
	v12 = t21
	t22 := int32(load32(m.memory[int64(uint32(v1))+12:]))
	v13 = t22
	t23 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	v14 = t23
	t24 := int32(m.memory[int64(uint32(v1))+47])
	v15 = t24
	t25 := int32(m.memory[int64(uint32(v2))+47])
	v16 = t25
	t26 := int32(m.memory[int64(uint32(v1))+46])
	v17 = t26
	t27 := int32(m.memory[int64(uint32(v2))+46])
	v18 = t27
	t28 := int32(load32(m.memory[uint32(v1):]))
	v19 = t28
	t29 := int32(m.memory[int64(uint32(v1))+40])
	v20 = t29
	t30 := int32(load16(m.memory[int64(uint32(v1))+36:]))
	v21 = t30
	t31 := int32(m.memory[int64(uint32(v1))+42])
	v22 = t31
	t32 := int32(m.memory[int64(uint32(v1))+45])
	v23 = t32
	t33 := int32(m.memory[int64(uint32(v2))+45])
	v24 = t33
	t34 := int32(m.memory[int64(uint32(v1))+44])
	v25 = t34
	t35 := int32(m.memory[int64(uint32(v2))+44])
	v26 = t35
	{
		{
			{
				t36 := int32(load32(m.memory[int64(uint32(v2))+8:]))
				v27 = t36
				if v27 == i32(-1) {
					goto l0
				}
				t37 := int64(load64(m.memory[int64(uint32(v2))+28:]))
				v28 = t37
				t38 := int32(load32(m.memory[int64(uint32(v2))+24:]))
				v1 = t38
				t39 := int32(load32(m.memory[int64(uint32(v2))+20:]))
				v29 = t39
				t40 := int32(load32(m.memory[int64(uint32(v2))+16:]))
				v30 = t40
				t41 := int32(load32(m.memory[int64(uint32(v2))+12:]))
				v2 = t41
				switch v14 + i32(1) {
				case 0:
					goto l1
				default:
					t42 := int32(load32(m.memory[uint32(v13+i32(-4)):]))
					v31 = t42
					v32 = v31 & i32(-8)
					t43 := v32
					v31 = v31 & i32(3)
					p44 := i32(8)
					if v31 != 0 {
						p44 = i32(4)
					}
					v14 = v14 << 1
					if uint32(t43) < uint32(p44+v14) {
						m.fn3(i32(1273840), i32(46), i32(1273888))
						panic("unreachable")
					}
					if v31 == 0 {
						goto l5
					}
					if uint32(v32) > uint32(v14+i32(39)) {
						m.fn3(i32(1273904), i32(46), i32(1273952))
						panic("unreachable")
					}
				l5:
					m.fn1(v13)
					fallthrough
				case 1:
					if v12 == 0 {
						goto l1
					}
					t45 := int32(load32(m.memory[uint32(v11+i32(-4)):]))
					v14 = t45
					v13 = v14 & i32(-8)
					t46 := v13
					v14 = v14 & i32(3)
					p47 := i32(8)
					if v14 != 0 {
						p47 = i32(4)
					}
					v12 = v12 << 2
					if uint32(t46) < uint32(p47+v12) {
						m.fn3(i32(1273840), i32(46), i32(1273888))
						panic("unreachable")
					}
					if v14 == 0 {
						goto l8
					}
					if uint32(v13) > uint32(v12+i32(39)) {
						m.fn3(i32(1273904), i32(46), i32(1273952))
						panic("unreachable")
					}
				l8:
					m.fn1(v11)
					goto l1
				}
			}
		l0:
			t48 := int32(load32(m.memory[int64(uint32(v1))+16:]))
			v30 = t48
			t49 := int64(load64(m.memory[int64(uint32(v1))+28:]))
			v28 = t49
			v27 = v14
			v2 = v13
			v29 = v12
			v1 = v11
		}
	l1:
		m.memory[int64(uint32(v0))+43] = byte(v10)
		m.memory[int64(uint32(v0))+41] = byte(v6)
		store16(m.memory[int64(uint32(v0))+38:], uint16(v8))
		store32(m.memory[int64(uint32(v0))+4:], uint32(v4))
		store64(m.memory[int64(uint32(v0))+28:], uint64(v28))
		store32(m.memory[int64(uint32(v0))+24:], uint32(v1))
		store32(m.memory[int64(uint32(v0))+20:], uint32(v29))
		store32(m.memory[int64(uint32(v0))+16:], uint32(v30))
		store32(m.memory[int64(uint32(v0))+12:], uint32(v2))
		store32(m.memory[int64(uint32(v0))+8:], uint32(v27))
		m.memory[int64(uint32(v0))+40] = byte(v20 | v5)
		t51 := v0
		p50 := v21
		if v7 != 0 {
			p50 = i32(1)
		}
		store16(m.memory[int64(uint32(t51))+36:], uint16(p50))
		t53 := v0
		p52 := v19
		if v3 != 0 {
			p52 = i32(1)
		}
		store32(m.memory[uint32(t53):], uint32(p52))
		t55 := v0
		p54 := v9
		if v9 == i32(2) {
			p54 = v22
		}
		m.memory[int64(uint32(t55))+42] = byte(p54)
		t57 := v0
		p56 := v24
		if v24&i32(255) == i32(2) {
			p56 = v23
		}
		m.memory[int64(uint32(t57))+45] = byte(p56)
		t59 := v0
		p58 := v26
		if v26&i32(255) == i32(2) {
			p58 = v25
		}
		m.memory[int64(uint32(t59))+44] = byte(p58)
		t61 := v0
		p60 := v16
		if v16&i32(255) == i32(2) {
			p60 = v15
		}
		m.memory[int64(uint32(t61))+47] = byte(p60)
		t63 := v0
		p62 := v18
		if v18&i32(255) == i32(2) {
			p62 = v17
		}
		m.memory[int64(uint32(t63))+46] = byte(p62)
		return
	}
}
func (m *Module) fn548(v0, v1, v2, v3 int32) int32 {
	var v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14, v15 int32
	t0 := m.g0
	v4 = t0 - i32(16)
	m.g0 = v4
	m.memory[int64(uint32(v4))+15] = byte(v2)
	t1 := v4
	v5 = int32(uint32(v2) >> 8)
	m.memory[int64(uint32(t1))+14] = byte(v5)
	t2 := v4
	v6 = int32(uint32(v2) >> 16)
	m.memory[int64(uint32(t2))+13] = byte(v6)
	v7 = v2
	{
		if uint32(v1) < uint32(i32(2)) {
			goto l0
		}
		v8 = v3 ^ i32(1)
		v5 = v3 & i32(256)
		v9 = int32(uint32(v5) >> 8)
		v6 = v3 & i32(65536)
		v10 = int32(uint32(v6) >> 16)
		var p3 int32
		if v5 == 0 {
			p3 = 1
		}
		v11 = p3
		var p4 int32
		if v6 == 0 {
			p4 = 1
		}
		v12 = p4
		v5 = i32(0)
		v6 = i32(2)
	l32:
		{
			{
				if uint32(v5) >= uint32(v1) {
					m.fn33(v5, v1, i32(1069384))
					panic("unreachable")
				}
				v7 = v5 + i32(1)
				if uint32(v7) >= uint32(v1) {
					m.fn33(v7, v1, i32(1069400))
					panic("unreachable")
				}
				t5 := int32(m.memory[uint32(v0+v7)])
				v13 = t5
				t6 := int32(m.memory[uint32(v0+v5)])
				v7 = v13<<8 | t6
				v14 = v0 + v6
				v15 = v1 - v6
				v5 = i32(1)
				switch int32(uint32(v13) >> 5) {
				default:
					goto l3
				case 2, 4, 5:
					v5 = i32(2)
					goto l3
				case 3:
					v5 = i32(4)
					goto l3
				case 7:
					v5 = i32(3)
					goto l3
				case 6:
					if v7 == i32(54792) {
						goto l8
					}
					v5 = i32(0)
					if v1 == v6 {
						goto l9
					}
					t7 := int32(m.memory[uint32(v14)])
					v5 = t7 + i32(1)
					goto l3
				}
			}
		l8:
			if uint32(v15) > uint32(i32(1)) {
				goto l10
			}
			v5 = i32(0)
			goto l11
		l10:
			t8 := int32(load16(m.memory[uint32(v14):]))
			v5 = t8 + i32(1)
		}
	l3:
		if uint32(v5) > uint32(v15) {
			goto l12
		}
	l9:
		switch v7 + i32(-2101) {
		default:
			goto l11
		case 0:
			if v5 != 0 {
				v7 = i32(0)
				v13 = v4 + i32(15)
				{
					t9 := int32(m.memory[uint32(v14)])
					v15 = t9
					switch v15 {
					case 0:
						goto l19
					default:
						switch v15 + i32(-128) {
						case 0:
							v7 = v3
							goto l19
						case 1:
							v7 = v8
							goto l19
						default:
							goto l11
						}
					case 1:
						v7 = i32(1)
						goto l19
					}
				}
			}
			v5 = i32(0)
			goto l11
		case 1:
			if v5 != 0 {
				v7 = i32(0)
				v13 = v4 + i32(14)
				{
					t10 := int32(m.memory[uint32(v14)])
					v15 = t10
					switch v15 {
					case 0:
						goto l19
					default:
						switch v15 + i32(-128) {
						case 0:
							v7 = v9
							goto l19
						case 1:
							v7 = v11
							goto l19
						default:
							goto l11
						}
					case 1:
						v7 = i32(1)
						goto l19
					}
				}
			}
			v5 = i32(0)
			goto l11
		case 2:
			if v5 != 0 {
				goto l18
			}
			v5 = i32(0)
			goto l11
		}
	l18:
		v7 = i32(0)
		v13 = v4 + i32(13)
		{
			t11 := int32(m.memory[uint32(v14)])
			v15 = t11
			switch v15 {
			case 0:
				goto l19
			default:
				switch v15 + i32(-128) {
				case 0:
					v7 = v10
					goto l19
				case 1:
					goto l31
				default:
					goto l11
				}
			case 1:
				v7 = i32(1)
				goto l19
			}
		}
	l31:
		v7 = v12
	l19:
		m.memory[uint32(v13)] = byte(v7 & i32(1))
	l11:
		v5 = v5 + v6
		v6 = v5 + i32(2)
		if uint32(v6) <= uint32(v1) {
			goto l32
		}
	l12:
		t12 := int32(m.memory[int64(uint32(v4))+13])
		v6 = t12
		t13 := int32(m.memory[int64(uint32(v4))+14])
		v5 = t13
		t14 := int32(m.memory[int64(uint32(v4))+15])
		v7 = t14
	}
l0:
	m.g0 = v4 + i32(16)
	return v2&i32(-0x1000000) | v6&i32(255)<<16 | v5&i32(255)<<8 | v7&i32(255)
}
func (m *Module) fn549(v0, v1, v2, v3 int32) {
	var v4 int32
	var v5 int64
	var v6, v7 int32
	var v8, v9 int64
	var v10, v11, v12, v13 int32
	var v14 int64
	var v15, v16 int32
	t0 := m.g0
	v4 = t0 - i32(64)
	m.g0 = v4
	t1 := int64(load64(m.memory[int64(uint32(v1))+16:]))
	t2 := int64(load64(m.memory[int64(uint32(v1))+24:]))
	t3 := m.fn106(t1, t2, v2)
	v5 = t3
	{
		t4 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		if t4 != 0 {
			goto l0
		}
		_ = m.fn108(v1, v1+i32(16))
	}
l0:
	t6 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	v6 = t6
	v7 = v6 & int32(v5)
	v8 = int64(uint64(v5) >> 25)
	v9 = v8 & i64(127) * i64(72340172838076673)
	t7 := int32(load32(m.memory[uint32(v1):]))
	v10 = t7
	v11 = v2 & i32(0xffff)
	v12 = i32(0)
	v13 = i32(0)
l10:
	{
		{
			t8 := int64(load64(m.memory[uint32(v10+v7):]))
			v14 = t8
			v5 = v14 ^ v9
			v5 = (v5 ^ i64(-1)) & (v5 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
			if v5 == 0 {
				goto l1
			}
		l3:
			{
				t9 := v11
				v15 = v10 + (i32(0)-(int32(uint32(int64(bits.TrailingZeros64(uint64(v5))))>>3)+v7)&v6)*i32(60)
				t10 := int32(load16(m.memory[uint32(v15+i32(-60)):]))
				if t9 == t10 {
					goto l2
				}
				v5 = (v5 + i64(-1)) & v5
				if !(v5 == 0) {
					goto l3
				}
			}
		}
	l1:
		v5 = v14 & i64(-0x7f7f7f7f7f7f7f80)
		if v12 == i32(1) {
			goto l4
		}
		if v5 == 0 {
			goto l5
		}
		v16 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v5))))>>3) + v7) & v6
	l4:
		if v5&(v14<<1) != i64(0) {
			{
				t11 := int32(int8(m.memory[uint32(v10+v16)]))
				v7 = t11
				if v7 < i32(0) {
					goto l8
				}
				t12 := int64(load64(m.memory[uint32(v10):]))
				t13 := v10
				v16 = int32(uint32(int64(bits.TrailingZeros64(uint64(t12&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
				t14 := int32(m.memory[uint32(t13+v16)])
				v7 = t14
			}
		l8:
			t15 := v10 + v16
			v11 = int32(v8) & i32(127)
			m.memory[uint32(t15)] = byte(v11)
			m.memory[uint32(v10+(v16+i32(-8))&v6+i32(8))] = byte(v11)
			t16 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			store32(m.memory[int64(uint32(v1))+8:], uint32(t16-v7&i32(1)))
			v10 = v10 + (i32(0)-v16)*i32(60)
			store16(m.memory[uint32(v10+i32(-60)):], uint16(v2))
			t17 := int64(load64(m.memory[int64(uint32(v3))+48:]))
			store64(m.memory[int64(uint32(v4))+56:], uint64(t17))
			t18 := int64(load64(m.memory[int64(uint32(v3))+40:]))
			store64(m.memory[int64(uint32(v4))+48:], uint64(t18))
			t19 := int64(load64(m.memory[int64(uint32(v3))+32:]))
			store64(m.memory[int64(uint32(v4))+40:], uint64(t19))
			t20 := int64(load64(m.memory[int64(uint32(v3))+24:]))
			store64(m.memory[int64(uint32(v4))+32:], uint64(t20))
			t21 := int64(load64(m.memory[int64(uint32(v3))+16:]))
			store64(m.memory[int64(uint32(v4))+24:], uint64(t21))
			t22 := int64(load64(m.memory[int64(uint32(v3))+8:]))
			store64(m.memory[int64(uint32(v4))+16:], uint64(t22))
			t23 := int64(load64(m.memory[uint32(v3):]))
			store64(m.memory[int64(uint32(v4))+8:], uint64(t23))
			t24 := int32(load32(m.memory[int64(uint32(v1))+12:]))
			store32(m.memory[int64(uint32(v1))+12:], uint32(t24+i32(1)))
			v1 = v10 + i32(-58)
			t25 := int64(load64(m.memory[int64(uint32(v4))+6:]))
			store64(m.memory[uint32(v1):], uint64(t25))
			t26 := int64(load64(m.memory[int64(uint32(v4))+14:]))
			store64(m.memory[int64(uint32(v1))+8:], uint64(t26))
			t27 := int64(load64(m.memory[int64(uint32(v4))+22:]))
			store64(m.memory[int64(uint32(v1))+16:], uint64(t27))
			t28 := int64(load64(m.memory[int64(uint32(v4))+30:]))
			store64(m.memory[int64(uint32(v1))+24:], uint64(t28))
			t29 := int64(load64(m.memory[int64(uint32(v4))+38:]))
			store64(m.memory[int64(uint32(v1))+32:], uint64(t29))
			t30 := int64(load64(m.memory[int64(uint32(v4))+46:]))
			store64(m.memory[int64(uint32(v1))+40:], uint64(t30))
			t31 := int64(load64(m.memory[int64(uint32(v4))+54:]))
			store64(m.memory[int64(uint32(v1))+48:], uint64(t31))
			t32 := int32(load16(m.memory[int64(uint32(v4))+62:]))
			store16(m.memory[int64(uint32(v1))+56:], uint16(t32))
			store32(m.memory[uint32(v0):], uint32(i32(2)))
			goto l9
		}
		v12 = i32(1)
		goto l7
	l2:
		t33 := v0
		v1 = v15 + i32(-56)
		t34 := int64(load64(m.memory[int64(uint32(v1))+48:]))
		store64(m.memory[int64(uint32(t33))+48:], uint64(t34))
		t35 := int64(load64(m.memory[int64(uint32(v1))+40:]))
		store64(m.memory[int64(uint32(v0))+40:], uint64(t35))
		t36 := int64(load64(m.memory[int64(uint32(v1))+32:]))
		store64(m.memory[int64(uint32(v0))+32:], uint64(t36))
		t37 := int64(load64(m.memory[int64(uint32(v1))+24:]))
		store64(m.memory[int64(uint32(v0))+24:], uint64(t37))
		t38 := int64(load64(m.memory[int64(uint32(v1))+16:]))
		store64(m.memory[int64(uint32(v0))+16:], uint64(t38))
		t39 := int64(load64(m.memory[int64(uint32(v1))+8:]))
		store64(m.memory[int64(uint32(v0))+8:], uint64(t39))
		t40 := int64(load64(m.memory[uint32(v1):]))
		store64(m.memory[uint32(v0):], uint64(t40))
		t41 := int64(load64(m.memory[uint32(v3):]))
		store64(m.memory[uint32(v1):], uint64(t41))
		t42 := int64(load64(m.memory[int64(uint32(v3))+8:]))
		store64(m.memory[int64(uint32(v1))+8:], uint64(t42))
		t43 := int64(load64(m.memory[int64(uint32(v3))+16:]))
		store64(m.memory[int64(uint32(v1))+16:], uint64(t43))
		t44 := int64(load64(m.memory[int64(uint32(v3))+24:]))
		store64(m.memory[int64(uint32(v1))+24:], uint64(t44))
		t45 := int64(load64(m.memory[int64(uint32(v3))+32:]))
		store64(m.memory[int64(uint32(v1))+32:], uint64(t45))
		t46 := int64(load64(m.memory[int64(uint32(v3))+40:]))
		store64(m.memory[int64(uint32(v1))+40:], uint64(t46))
		t47 := int64(load64(m.memory[int64(uint32(v3))+48:]))
		store64(m.memory[int64(uint32(v1))+48:], uint64(t47))
	}
l9:
	m.g0 = v4 + i32(64)
	return
l5:
	v12 = i32(0)
l7:
	v13 = v13 + i32(8)
	v7 = (v13 + v7) & v6
	goto l10
}
func (m *Module) fn550(v0 int32) {
	var v1, v2, v3, v4 int32
	var v5 int64
	var v6, v7, v8, v9, v10 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		v1 = t0
		if v1 == 0 {
			return
		}
		{
			t1 := int32(load32(m.memory[int64(uint32(v0))+12:]))
			v2 = t1
			if v2 == 0 {
				goto l1
			}
			t2 := int32(load32(m.memory[uint32(v0):]))
			v3 = t2
			v4 = v3 + i32(8)
			t3 := int64(load64(m.memory[uint32(v3):]))
			v5 = (t3 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
		l12:
			if v5 != i64(0) {
				goto l2
			}
		l3:
			{
				v6 = v4
				v4 = v6 + i32(8)
				v3 = v3 + i32(-480)
				t4 := int64(load64(m.memory[uint32(v6):]))
				v5 = t4 & i64(-0x7f7f7f7f7f7f7f80)
				if v5 == i64(-0x7f7f7f7f7f7f7f80) {
					goto l3
				}
			}
			v5 = v5 ^ i64(-0x7f7f7f7f7f7f7f80)
		l2:
			{
				v6 = v3 + (i32(0)-int32(uint32(int64(bits.TrailingZeros64(uint64(v5))))>>3))*i32(60)
				t5 := int32(load32(m.memory[uint32(v6+i32(-48)):]))
				v7 = t5
				if v7 == i32(-1) {
					goto l4
				}
				{
					if v7 == 0 {
						goto l5
					}
					t6 := int32(load32(m.memory[uint32(v6+i32(-44)):]))
					v8 = t6
					t7 := int32(load32(m.memory[uint32(v8+i32(-4)):]))
					v9 = t7
					v10 = v9 & i32(-8)
					t8 := v10
					v9 = v9 & i32(3)
					p9 := i32(8)
					if v9 != 0 {
						p9 = i32(4)
					}
					v7 = v7 << 1
					if uint32(t8) < uint32(p9+v7) {
						m.fn3(i32(1273840), i32(46), i32(1273888))
						panic("unreachable")
					}
					if v9 == 0 {
						goto l7
					}
					if uint32(v10) > uint32(v7+i32(39)) {
						m.fn3(i32(1273904), i32(46), i32(1273952))
						panic("unreachable")
					}
				l7:
					m.fn1(v8)
				}
			l5:
				t10 := int32(load32(m.memory[uint32(v6+i32(-36)):]))
				v7 = t10
				if v7 == 0 {
					goto l4
				}
				t11 := int32(load32(m.memory[uint32(v6+i32(-32)):]))
				v9 = t11
				t12 := int32(load32(m.memory[uint32(v9+i32(-4)):]))
				v6 = t12
				v8 = v6 & i32(-8)
				t13 := v8
				v6 = v6 & i32(3)
				p14 := i32(8)
				if v6 != 0 {
					p14 = i32(4)
				}
				v7 = v7 << 2
				if uint32(t13) < uint32(p14+v7) {
					m.fn3(i32(1273840), i32(46), i32(1273888))
					panic("unreachable")
				}
				if v6 == 0 {
					goto l10
				}
				if uint32(v8) > uint32(v7+i32(39)) {
					m.fn3(i32(1273904), i32(46), i32(1273952))
					panic("unreachable")
				}
			l10:
				m.fn1(v9)
			}
		l4:
			v5 = (v5 + i64(-1)) & v5
			v2 = v2 + i32(-1)
			if v2 != 0 {
				goto l12
			}
		}
	l1:
		t15 := v1
		v4 = (v1*i32(60) + i32(67)) & i32(-8)
		v3 = t15 + v4 + i32(9)
		if v3 == 0 {
			return
		}
		t16 := int32(load32(m.memory[uint32(v0):]))
		v6 = t16 - v4
		t17 := int32(load32(m.memory[uint32(v6+i32(-4)):]))
		v4 = t17
		v2 = v4 & i32(-8)
		t18 := v2
		v4 = v4 & i32(3)
		p19 := i32(8)
		if v4 != 0 {
			p19 = i32(4)
		}
		if uint32(t18) < uint32(p19+v3) {
			m.fn3(i32(1273840), i32(46), i32(1273888))
			panic("unreachable")
		}
		if v4 == 0 {
			goto l14
		}
		if uint32(v2) > uint32(v3+i32(39)) {
			m.fn3(i32(1273904), i32(46), i32(1273952))
			panic("unreachable")
		}
	l14:
		m.fn1(v6)
	}
}
func (m *Module) fn551(v0, v1, v2, v3 int32) {
	var v4, v5, v6, v7 int32
	var v8 int64
	var v9, v10, v11, v12, v13, v14, v15, v16, v17, v18, v19, v20, v21, v22, v23, v24, v25, v26, v27, v28, v29, v30 int32
	t0 := m.g0
	v4 = t0 - i32(16)
	m.g0 = v4
	{
		if uint32(v2) < uint32(v3) {
			store32(m.memory[uint32(v0):], uint32(i32(2)))
			goto l2
		}
		v5 = v2 - v3
		if uint32(v5) > uint32(i32(27)) {
			v6 = v1 + v3
			t1 := int32(m.memory[int64(uint32(v6))+5])
			v7 = t1
			t2 := int32(m.memory[int64(uint32(v6))+4])
			v1 = t2
			t3 := int64(load32(m.memory[uint32(v6):]))
			v8 = t3
			{
				{
					{
						t4 := int32(m.memory[int64(uint32(v6))+6])
						v9 = t4
						if v9 != 0 {
							goto l3
						}
						v10 = i32(1)
						v11 = i32(0)
						v12 = i32(1)
						v13 = i32(0)
						goto l4
					}
				l3:
					v13 = i32(8)
					t5 := m.fn5(i32(8))
					v12 = t5
					if v12 == 0 {
						m.fn10(i32(1), i32(8))
						panic("unreachable")
					}
					m.memory[uint32(v12)] = byte(v9)
					store32(m.memory[int64(uint32(v4))+8:], uint32(v12))
					store32(m.memory[int64(uint32(v4))+4:], uint32(i32(8)))
					{
						t6 := int32(m.memory[int64(uint32(v6))+7])
						v2 = t6
						if v2 != 0 {
							goto l6
						}
						v10 = i32(1)
						v11 = i32(1)
						goto l4
					}
				l6:
					m.memory[int64(uint32(v12))+1] = byte(v2)
					v10 = i32(1)
					{
						t7 := int32(m.memory[int64(uint32(v6))+8])
						v2 = t7
						if v2 != 0 {
							goto l7
						}
						v11 = i32(2)
						goto l4
					}
				l7:
					m.memory[int64(uint32(v12))+2] = byte(v2)
					{
						t8 := int32(m.memory[int64(uint32(v6))+9])
						v2 = t8
						if v2 != 0 {
							goto l8
						}
						v11 = i32(3)
						goto l4
					}
				l8:
					m.memory[int64(uint32(v12))+3] = byte(v2)
					{
						t9 := int32(m.memory[int64(uint32(v6))+10])
						v2 = t9
						if v2 != 0 {
							goto l9
						}
						v11 = i32(4)
						goto l4
					}
				l9:
					m.memory[int64(uint32(v12))+4] = byte(v2)
					{
						t10 := int32(m.memory[int64(uint32(v6))+11])
						v2 = t10
						if v2 != 0 {
							goto l10
						}
						v11 = i32(5)
						goto l4
					}
				l10:
					m.memory[int64(uint32(v12))+5] = byte(v2)
					{
						t11 := int32(m.memory[int64(uint32(v6))+12])
						v2 = t11
						if v2 != 0 {
							goto l11
						}
						v11 = i32(6)
						goto l4
					}
				l11:
					m.memory[int64(uint32(v12))+6] = byte(v2)
					{
						t12 := int32(m.memory[int64(uint32(v6))+13])
						v2 = t12
						if v2 != 0 {
							goto l12
						}
						v11 = i32(7)
						goto l4
					}
				l12:
					m.memory[int64(uint32(v12))+7] = byte(v2)
					v11 = i32(8)
					store32(m.memory[int64(uint32(v4))+12:], uint32(i32(8)))
					v10 = i32(0)
					{
						t13 := int32(m.memory[int64(uint32(v6))+14])
						v2 = t13
						if v2 != 0 {
							goto l13
						}
						v13 = i32(8)
						goto l4
					}
				l13:
					m.fn197(v4+i32(4), i32(8), i32(1), i32(1), i32(1))
					t14 := int32(load32(m.memory[int64(uint32(v4))+8:]))
					v12 = t14
					m.memory[int64(uint32(v12))+8] = byte(v2)
					v11 = i32(9)
					t15 := int32(load32(m.memory[int64(uint32(v4))+4:]))
					v13 = t15
				}
			l4:
				t16 := int32(m.memory[int64(uint32(v6))+25])
				t17 := int32(m.memory[int64(uint32(v6))+24])
				t18 := v5
				v2 = t16 + t17 + i32(28)
				if uint32(t18) < uint32(v2) {
					goto l14
				}
				if uint32(v5-v2) < uint32(i32(2)) {
					goto l14
				}
				t19 := int32(m.memory[int64(uint32(v6))+26])
				v14 = t19
				v15 = v2 + i32(2)
				t20 := int32(load16(m.memory[uint32(v6+v2):]))
				v16 = t20
				v17 = i32(1)
				v2 = i32(0)
				switch v1 + i32(-1) {
				case 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21:
					goto l19
				case 22:
					goto l20
				case 0:
					v17 = i32(5)
					goto l19
				case 1:
					v17 = i32(4)
					goto l19
				case 2:
					v17 = i32(3)
					goto l19
				case 3:
					v17 = i32(2)
					goto l19
				default:
					v2 = i32(255)
					if v1 == i32(255) {
						goto l20
					}
				}
			l19:
				store32(m.memory[int64(uint32(v4))+12:], uint32(i32(0)))
				store64(m.memory[int64(uint32(v4))+4:], uint64(i64(0x400000000)))
				if v16 != 0 {
					t21 := v11
					v18 = (v12 + i32(3)) & i32(-4)
					v2 = v18 - v12
					p22 := v2
					if uint32(v11) < uint32(v2) {
						p22 = t21
					}
					v19 = p22
					v20 = v11 + i32(-8)
					v21 = i32(4)
					v22 = i32(0)
					v1 = i32(0)
				l57:
					{
						{
							{
								{
									t23 := v5
									v2 = v1<<1 + v15
									if uint32(t23) < uint32(v2) {
										goto l24
									}
									if uint32(v5-v2) <= uint32(i32(1)) {
										goto l24
									}
									v1 = v1 + i32(1)
									t24 := int32(load16(m.memory[uint32(v6+v2):]))
									v23 = t24
									if uint32(v23) < uint32(i32(9)) {
										if v10 == 0 {
											v2 = i32(0)
											if v18 == v12 {
												goto l37
											}
											v2 = v12
											v25 = v19
											if v9 == 0 {
												goto l38
											}
										l39:
											{
												t32 := int32(m.memory[uint32(v2)])
												if t32 == v1&i32(255) {
													goto l35
												}
												v2 = v2 + i32(1)
												v25 = v25 + i32(-1)
												if v25 == 0 {
													goto l38
												}
												goto l39
											}
										}
										if v9 == 0 {
											goto l26
										}
										v24 = v1 & i32(255)
										v2 = v12
										v25 = v11
									l36:
										{
											t31 := int32(m.memory[uint32(v2)])
											if t31 == v24 {
												goto l35
											}
											v2 = v2 + i32(1)
											v25 = v25 + i32(-1)
											if v25 == 0 {
												goto l26
											}
											goto l36
										}
									}
									goto l26
								}
							l24:
								store32(m.memory[uint32(v0):], uint32(i32(2)))
								if v22 == 0 {
									goto l27
								}
								v2 = v21
							l32:
								{
									t25 := int32(load32(m.memory[uint32(v2):]))
									v1 = t25
									if v1 < i32(1) {
										goto l28
									}
									t26 := int32(load32(m.memory[uint32(v2+i32(4)):]))
									v24 = t26
									t27 := int32(load32(m.memory[uint32(v24+i32(-4)):]))
									v25 = t27
									v26 = v25 & i32(-8)
									t28 := v26
									v25 = v25 & i32(3)
									p29 := i32(8)
									if v25 != 0 {
										p29 = i32(4)
									}
									if uint32(t28) < uint32(p29+v1) {
										m.fn3(i32(1273840), i32(46), i32(1273888))
										panic("unreachable")
									}
									if v25 == 0 {
										goto l30
									}
									if uint32(v26) > uint32(v1+i32(39)) {
										m.fn3(i32(1273904), i32(46), i32(1273952))
										panic("unreachable")
									}
								l30:
									m.fn1(v24)
								}
							l28:
								v2 = v2 + i32(12)
								v22 = v22 + i32(-1)
								if v22 != 0 {
									goto l32
								}
							l27:
								t30 := int32(load32(m.memory[int64(uint32(v4))+4:]))
								v2 = t30
								if v2 == 0 {
									goto l33
								}
								m.fn18(v21, v2*i32(12), i32(4))
								goto l33
							}
						l38:
							v2 = v19
							if uint32(v19) > uint32(v20) {
								goto l40
							}
						l37:
							v25 = v1 & i32(255) * i32(16843009)
						l41:
							{
								v24 = v12 + v2
								t33 := int32(load32(m.memory[uint32(v24):]))
								v26 = t33 ^ v25
								t34 := int32(load32(m.memory[uint32(v24+i32(4)):]))
								t35 := i32(16843008) - v26 | v26
								v24 = t34 ^ v25
								if t35&(i32(16843008)-v24|v24)&i32(-2139062144) != i32(-2139062144) {
									goto l40
								}
								v2 = v2 + i32(8)
								if uint32(v2) <= uint32(v20) {
									goto l41
								}
							}
						l40:
							if v11 == v2 {
								goto l26
							}
							v25 = v11 - v2
							v2 = v12 + v2
						l42:
							{
								t36 := int32(m.memory[uint32(v2)])
								if t36 == v1&i32(255) {
									goto l35
								}
								v2 = v2 + i32(1)
								v25 = v25 + i32(-1)
								if v25 == 0 {
									goto l26
								}
								goto l42
							}
						l35:
							{
								t37 := int32(load32(m.memory[int64(uint32(v4))+4:]))
								if v22 != t37 {
									goto l43
								}
								m.fn316(v4 + i32(4))
							}
						l43:
							t38 := int32(load32(m.memory[int64(uint32(v4))+8:]))
							v21 = t38
							v2 = v21 + v22*i32(12)
							m.memory[int64(uint32(v2))+4] = byte(v23)
							store32(m.memory[uint32(v2):], uint32(i32(-1)))
							goto l44
						}
					l26:
						if uint32((v23+i32(-160))&i32(0xffff)) > uint32(i32(65502)) {
							goto l45
						}
						if uint32(v23) < uint32(i32(32)) {
							goto l45
						}
						if uint32(v23^i32(-1058816)) <= uint32(i32(-1112065)) {
							goto l45
						}
						{
							if v22 == 0 {
								goto l46
							}
							v2 = v21 + v22*i32(12)
							v27 = v2 + i32(-12)
							t39 := int32(load32(m.memory[uint32(v27):]))
							v24 = t39
							if v24 == i32(-1) {
								goto l46
							}
							{
								var p40 int32
								if uint32(v23) < uint32(i32(2048)) {
									p40 = 1
								}
								v28 = p40
								p41 := i32(3)
								if v28 != 0 {
									p41 = i32(2)
								}
								var p42 int32
								if uint32(v23) < uint32(i32(128)) {
									p42 = 1
								}
								v29 = p42
								p43 := p41
								if v29 != 0 {
									p43 = i32(1)
								}
								v26 = p43
								t44 := v26
								t45 := v24
								v30 = v2 + i32(-4)
								t46 := int32(load32(m.memory[uint32(v30):]))
								v25 = t46
								if uint32(t44) <= uint32(t45-v25) {
									goto l47
								}
								m.fn197(v27, v25, v26, i32(1), i32(1))
							}
						l47:
							t47 := int32(load32(m.memory[uint32(v2+i32(-8)):]))
							v2 = t47 + v25
							if v29 != 0 {
								goto l48
							}
							v24 = int32(uint32(v23) >> 6)
							v29 = v23&i32(63) | i32(-128)
							if v28 == 0 {
								m.memory[int64(uint32(v2))+2] = byte(v29)
								m.memory[int64(uint32(v2))+1] = byte(v24&i32(63) | i32(128))
								m.memory[uint32(v2)] = byte(int32(uint32(v23)>>12) | i32(224))
								goto l50
							}
							m.memory[int64(uint32(v2))+1] = byte(v29)
							m.memory[uint32(v2)] = byte(v24 | i32(192))
							goto l50
						}
					l46:
						store32(m.memory[uint32(v4):], uint32(i32(0)))
						if uint32(v23) < uint32(i32(128)) {
							m.memory[uint32(v4)] = byte(v23)
							v2 = i32(1)
							goto l53
						}
						v2 = int32(uint32(v23) >> 6)
						v25 = v23&i32(63) | i32(-128)
						if uint32(v23) >= uint32(i32(2048)) {
							m.memory[int64(uint32(v4))+2] = byte(v25)
							m.memory[int64(uint32(v4))+1] = byte(v2&i32(63) | i32(128))
							m.memory[uint32(v4)] = byte(int32(uint32(v23)>>12) | i32(224))
							v2 = i32(3)
							goto l53
						}
						m.memory[int64(uint32(v4))+1] = byte(v25)
						m.memory[uint32(v4)] = byte(v2 | i32(192))
						v2 = i32(2)
						goto l53
					l48:
						m.memory[uint32(v2)] = byte(v23)
					l50:
						store32(m.memory[uint32(v30):], uint32(v25+v26))
						goto l45
					l53:
						{
							t48 := m.fn5(v2)
							v24 = t48
							if v24 != 0 {
								goto l54
							}
							m.fn10(i32(1), v2)
							panic("unreachable")
						}
					l54:
						if v2 == 0 {
							goto l55
						}
						memory_copy(m.memory, uint32(v24), uint32(v4), uint32(v2))
					l55:
						{
							t49 := int32(load32(m.memory[int64(uint32(v4))+4:]))
							if v22 != t49 {
								goto l56
							}
							m.fn316(v4 + i32(4))
							t50 := int32(load32(m.memory[int64(uint32(v4))+8:]))
							v21 = t50
						}
					l56:
						v25 = v21 + v22*i32(12)
						store32(m.memory[int64(uint32(v25))+8:], uint32(v2))
						store32(m.memory[int64(uint32(v25))+4:], uint32(v24))
						store32(m.memory[uint32(v25):], uint32(v2))
					l44:
						t51 := v4
						v22 = v22 + i32(1)
						store32(m.memory[int64(uint32(t51))+12:], uint32(v22))
					}
				l45:
					if v1 != v16 {
						goto l57
					}
					v2 = v17
					goto l23
				}
				v2 = v17
				goto l23
			}
		}
		store32(m.memory[uint32(v0):], uint32(i32(2)))
		goto l2
	l14:
		store32(m.memory[uint32(v0):], uint32(i32(2)))
	l33:
		if v13 == 0 {
			goto l2
		}
		{
			t52 := int32(load32(m.memory[uint32(v12+i32(-4)):]))
			v2 = t52
			v1 = v2 & i32(-8)
			t53 := v1
			v2 = v2 & i32(3)
			p54 := i32(8)
			if v2 != 0 {
				p54 = i32(4)
			}
			if uint32(t53) < uint32(p54+v13) {
				m.fn3(i32(1273840), i32(46), i32(1273888))
				panic("unreachable")
			}
			if v2 == 0 {
				goto l59
			}
			if uint32(v1) > uint32(v13+i32(39)) {
				m.fn3(i32(1273904), i32(46), i32(1273952))
				panic("unreachable")
			}
		l59:
			m.fn1(v12)
			goto l2
		}
	l20:
		store32(m.memory[int64(uint32(v4))+12:], uint32(i32(0)))
		store64(m.memory[int64(uint32(v4))+4:], uint64(i64(0x400000000)))
	l23:
		t55 := int32(load32(m.memory[int64(uint32(v4))+12:]))
		store32(m.memory[int64(uint32(v0))+16:], uint32(t55))
		t56 := int64(load64(m.memory[int64(uint32(v4))+4:]))
		store64(m.memory[int64(uint32(v0))+8:], uint64(t56))
		m.memory[int64(uint32(v0))+32] = byte(v2)
		store64(m.memory[int64(uint32(v0))+24:], uint64(v8))
		store32(m.memory[int64(uint32(v0))+4:], uint32(v14))
		store32(m.memory[int64(uint32(v0))+40:], uint32(v15+v3+v16<<1))
		m.memory[int64(uint32(v0))+20] = byte(int32(uint32(v7)>>2) & i32(1))
		store32(m.memory[uint32(v0):], uint32(int32(uint32(v7)>>3)&i32(1)))
		if v13 == 0 {
			goto l2
		}
		{
			t57 := int32(load32(m.memory[uint32(v12+i32(-4)):]))
			v2 = t57
			v1 = v2 & i32(-8)
			t58 := v1
			v2 = v2 & i32(3)
			p59 := i32(8)
			if v2 != 0 {
				p59 = i32(4)
			}
			if uint32(t58) < uint32(p59+v13) {
				m.fn3(i32(1273840), i32(46), i32(1273888))
				panic("unreachable")
			}
			if v2 == 0 {
				goto l62
			}
			if uint32(v1) > uint32(v13+i32(39)) {
				m.fn3(i32(1273904), i32(46), i32(1273952))
				panic("unreachable")
			}
		l62:
			m.fn1(v12)
			goto l2
		}
	}
l2:
	m.g0 = v4 + i32(16)
}
func (m *Module) fn552(v0, v1, v2, v3, v4, v5, v6 int32) {
	var v7, v8, v9, v10 int32
	t0 := m.g0
	v7 = t0 - i32(16)
	m.g0 = v7
	v8 = i32(0)
	{
		if uint32(v2) < uint32(v5) {
			goto l0
		}
		if uint32(v2-v5) < uint32(i32(4)) {
			goto l0
		}
		t1 := int32(load32(m.memory[uint32(v1+v5):]))
		v8 = t1
	}
l0:
	{
		{
			{
				{
					t2 := v2
					v5 = v5 + i32(4)
					if uint32(t2) < uint32(v5) {
						goto l1
					}
					if uint32(v2-v5) < uint32(i32(4)) {
						goto l1
					}
					t3 := int32(load32(m.memory[uint32(v1+v5):]))
					t4 := v8
					v5 = t3
					v2 = t4 + v5
					p5 := v2
					if uint32(v2) < uint32(v8) {
						p5 = i32(-1)
					}
					v2 = p5
					if uint32(v2) > uint32(v4) {
						goto l2
					}
					if uint32(v5) < uint32(i32(8)) {
						goto l3
					}
					if v6 != 0 {
						goto l4
					}
					v5 = int32(uint32(v5)>>2) + i32(-1)
					goto l5
				}
			l1:
				if uint32(v8) > uint32(v4) {
					goto l2
				}
			l3:
				v5 = i32(0)
				store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
				store64(m.memory[uint32(v0):], uint64(i64(0x400000000)))
				goto l6
			l4:
				t6 := int32(uint32(v5+i32(-4)) / uint32(v6|i32(4)))
				v5 = t6
			}
		l5:
			if uint32(v5) >= uint32(i32(0x1fffffff)) {
				m.fn9()
				panic("unreachable")
			}
			v4 = v5 + i32(1)
			v1 = v4 << 2
			t7 := m.fn5(v1)
			v9 = t7
			if v9 == 0 {
				m.fn10(i32(4), v1)
				panic("unreachable")
			}
			v10 = v2 - v8
			v3 = v3 + v8
			v1 = i32(0)
			store32(m.memory[int64(uint32(v7))+12:], uint32(i32(0)))
			store32(m.memory[int64(uint32(v7))+8:], uint32(v9))
			store32(m.memory[int64(uint32(v7))+4:], uint32(v4))
			v8 = i32(1)
			v2 = i32(0)
		l12:
			{
				v4 = i32(0)
				{
					t8 := v10
					v6 = v2 << 2
					if uint32(t8) < uint32(v6) {
						goto l9
					}
					v4 = i32(0)
					if uint32(v10-v6) < uint32(i32(4)) {
						goto l9
					}
					t9 := int32(load32(m.memory[uint32(v3+v6):]))
					v4 = t9
				}
			l9:
				{
					t10 := int32(load32(m.memory[int64(uint32(v7))+4:]))
					if v8+i32(-1) != t10 {
						goto l10
					}
					m.fn174(v7 + i32(4))
					t11 := int32(load32(m.memory[int64(uint32(v7))+8:]))
					v9 = t11
				}
			l10:
				store32(m.memory[uint32(v9+v1):], uint32(v4))
				store32(m.memory[int64(uint32(v7))+12:], uint32(v8))
				if uint32(v2) >= uint32(v5) {
					goto l11
				}
				v1 = v1 + i32(4)
				v8 = v8 + i32(1)
				t12 := v2
				var p13 int32
				if uint32(v2) < uint32(v5) {
					p13 = 1
				}
				v2 = t12 + p13
				if uint32(v2) <= uint32(v5) {
					goto l12
				}
				goto l11
			}
		}
	l2:
		v5 = i32(0)
		store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
		store64(m.memory[uint32(v0):], uint64(i64(0x400000000)))
		goto l6
	l11:
		t14 := int32(load32(m.memory[int64(uint32(v7))+12:]))
		store32(m.memory[int64(uint32(v0))+8:], uint32(t14))
		t15 := int64(load64(m.memory[int64(uint32(v7))+4:]))
		store64(m.memory[uint32(v0):], uint64(t15))
	}
l6:
	store32(m.memory[int64(uint32(v0))+12:], uint32(v5))
	m.g0 = v7 + i32(16)
}
func (m *Module) fn553(v0, v1 int32) {
	var v2, v3, v4, v5, v6, v7, v8, v9 int32
	t0 := m.g0
	v2 = t0 - i32(64)
	v3 = v0 + i32(72)
	v4 = v0 + v1*i32(72)
	v5 = i32(0)
	v1 = v0
l4:
	v6 = v3
	{
		t1 := int32(load32(m.memory[uint32(v1+i32(136)):]))
		v7 = t1
		t2 := int32(load32(m.memory[uint32(v1+i32(64)):]))
		if uint32(v7) >= uint32(t2) {
			goto l0
		}
		t3 := int64(load64(m.memory[int64(uint32(v6))+56:]))
		store64(m.memory[int64(uint32(v2))+56:], uint64(t3))
		t4 := int64(load64(m.memory[int64(uint32(v6))+48:]))
		store64(m.memory[int64(uint32(v2))+48:], uint64(t4))
		t5 := int64(load64(m.memory[int64(uint32(v6))+40:]))
		store64(m.memory[int64(uint32(v2))+40:], uint64(t5))
		t6 := int64(load64(m.memory[int64(uint32(v6))+32:]))
		store64(m.memory[int64(uint32(v2))+32:], uint64(t6))
		t7 := int64(load64(m.memory[int64(uint32(v6))+24:]))
		store64(m.memory[int64(uint32(v2))+24:], uint64(t7))
		t8 := int64(load64(m.memory[int64(uint32(v6))+16:]))
		store64(m.memory[int64(uint32(v2))+16:], uint64(t8))
		t9 := int64(load64(m.memory[int64(uint32(v6))+8:]))
		store64(m.memory[int64(uint32(v2))+8:], uint64(t9))
		t10 := int64(load64(m.memory[uint32(v6):]))
		store64(m.memory[uint32(v2):], uint64(t10))
		t11 := int32(load32(m.memory[int64(uint32(v1))+140:]))
		v8 = t11
		v1 = v5
	l2:
		{
			v3 = v0 + v1
			v9 = v3 + i32(72)
			memory_copy(m.memory, uint32(v9), uint32(v3), uint32(i32(72)))
			if v1 == 0 {
				goto l1
			}
			v1 = v1 + i32(-72)
			t12 := int32(load32(m.memory[uint32(v3+i32(-8)):]))
			if uint32(v7) < uint32(t12) {
				goto l2
			}
		}
		v1 = v0 + v1 + i32(72)
		goto l3
	l1:
		v1 = v0
	l3:
		t13 := int64(load64(m.memory[int64(uint32(v2))+56:]))
		store64(m.memory[int64(uint32(v1))+56:], uint64(t13))
		t14 := int64(load64(m.memory[int64(uint32(v2))+48:]))
		store64(m.memory[int64(uint32(v1))+48:], uint64(t14))
		t15 := int64(load64(m.memory[int64(uint32(v2))+40:]))
		store64(m.memory[int64(uint32(v1))+40:], uint64(t15))
		t16 := int64(load64(m.memory[int64(uint32(v2))+32:]))
		store64(m.memory[int64(uint32(v1))+32:], uint64(t16))
		t17 := int64(load64(m.memory[int64(uint32(v2))+24:]))
		store64(m.memory[int64(uint32(v1))+24:], uint64(t17))
		t18 := int64(load64(m.memory[int64(uint32(v2))+16:]))
		store64(m.memory[int64(uint32(v1))+16:], uint64(t18))
		t19 := int64(load64(m.memory[int64(uint32(v2))+8:]))
		store64(m.memory[int64(uint32(v1))+8:], uint64(t19))
		t20 := int64(load64(m.memory[uint32(v2):]))
		store64(m.memory[uint32(v1):], uint64(t20))
		store32(m.memory[uint32(v9+i32(-4)):], uint32(v8))
		store32(m.memory[uint32(v9+i32(-8)):], uint32(v7))
	}
l0:
	v5 = v5 + i32(72)
	v1 = v6
	v3 = v6 + i32(72)
	if v3 != v4 {
		goto l4
	}
}
func (m *Module) fn554(v0, v1, v2, v3 int32) {
	var v4, v5, v6, v7, v8 int32
	var v9, v10, v11 int64
	var v12, v13, v14, v15, v16, v17, v18, v19, v20, v21 int32
	var v22 int64
	var v23 int32
	var v24 int64
	var v25, v26 int32
	var v27 int64
	var v28, v29, v30, v31, v32, v33, v34, v35, v36, v37, v38, v39, v40, v41, v42, v43, v44, v45, v46, v47 int32
	var v48, v49 int64
	var v50, v51, v52, v53, v54, v55, v56 int32
	t0 := m.g0
	v4 = t0 - i32(800)
	m.g0 = v4
	store32(m.memory[int64(uint32(v4))+12:], uint32(i32(0)))
	store64(m.memory[int64(uint32(v4))+4:], uint64(i64(0x800000000)))
	store32(m.memory[int64(uint32(v4))+24:], uint32(i32(0)))
	store64(m.memory[int64(uint32(v4))+16:], uint64(i64(0x800000000)))
	store32(m.memory[int64(uint32(v4))+28:], uint32(i32(0)))
	store32(m.memory[int64(uint32(v4))+52:], uint32(i32(0)))
	store64(m.memory[int64(uint32(v4))+44:], uint64(i64(0x800000000)))
	store32(m.memory[int64(uint32(v4))+56:], uint32(i32(0)))
	store32(m.memory[int64(uint32(v4))+80:], uint32(i32(0)))
	store64(m.memory[int64(uint32(v4))+72:], uint64(i64(0x400000000)))
	store32(m.memory[int64(uint32(v4))+92:], uint32(i32(0)))
	store64(m.memory[int64(uint32(v4))+84:], uint64(i64(0x400000000)))
	store64(m.memory[int64(uint32(v4))+128:], uint64(i64(0)))
	store64(m.memory[int64(uint32(v4))+120:], uint64(i64(0x100000000)))
	store64(m.memory[int64(uint32(v4))+112:], uint64(i64(4)))
	store64(m.memory[int64(uint32(v4))+104:], uint64(i64(0)))
	store64(m.memory[int64(uint32(v4))+96:], uint64(i64(0x400000000)))
	{
		{
			{
				{
					t1 := int32(load32(m.memory[int64(uint32(v1))+96:]))
					t2 := v2
					v5 = t1
					p3 := v3
					if uint32(v5) < uint32(v3) {
						p3 = v5
					}
					if uint32(t2) >= uint32(p3) {
						goto l0
					}
					v6 = v4 + i32(96) + i32(12)
					v7 = v1 + i32(208)
					v8 = v1 + i32(280)
					v9 = int64(uint32(i32(18)))<<32 | int64(uint32(v4+i32(280)))
					v10 = int64(uint32(i32(1)))<<32 | int64(uint32(v4+i32(252)))
					v11 = int64(uint32(i32(2)))<<32 | int64(uint32(v4+i32(792)))
					v12 = v4 + i32(232) + i32(8)
					v13 = v4 + i32(424)
					v14 = v4 + i32(280) + i32(12)
					v15 = v1 + i32(336)
					v16 = v1 + i32(332)
					v17 = v1 + i32(348)
					v18 = v1 + i32(344)
				l15:
					{
						{
							{
								{
									{
										{
											{
												{
													{
														{
															t4 := int32(load32(m.memory[int64(uint32(v1))+108:]))
															t5 := v2
															v5 = t4
															if uint32(t5) >= uint32(v5) {
																m.fn33(v2, v5, i32(1073816))
																panic("unreachable")
															}
															t6 := int32(load32(m.memory[int64(uint32(v1))+92:]))
															v19 = v2 << 2
															t7 := int32(load32(m.memory[uint32(t6+v19):]))
															v20 = t7
															t8 := int32(load32(m.memory[int64(uint32(v1))+104:]))
															t9 := int32(load32(m.memory[uint32(t8+v19):]))
															v21 = t9
															t10 := int32(load32(m.memory[int64(uint32(v1))+180:]))
															if t10 == 0 {
																goto l2
															}
															t11 := int64(load64(m.memory[int64(uint32(v1))+184:]))
															t12 := int64(load64(m.memory[int64(uint32(v1))+192:]))
															t13 := m.fn94(t11, t12, v2)
															v22 = t13
															t14 := int32(load32(m.memory[int64(uint32(v1))+172:]))
															v23 = t14
															v5 = v23 & int32(v22)
															v24 = int64(uint64(v22)>>25) & i64(127) * i64(72340172838076673)
															t15 := int32(load32(m.memory[int64(uint32(v1))+168:]))
															v25 = t15
															v26 = i32(0)
														l6:
															{
																{
																	t16 := int64(load64(m.memory[uint32(v25+v5):]))
																	v27 = t16
																	v22 = v27 ^ v24
																	v22 = (v22 ^ i64(-1)) & (v22 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
																	if v22 == 0 {
																		goto l3
																	}
																l5:
																	{
																		t17 := v2
																		v28 = v25 - (int32(uint32(int64(bits.TrailingZeros64(uint64(v22))))>>3)+v5)&v23<<4
																		t18 := int32(load32(m.memory[uint32(v28+i32(-16)):]))
																		if t17 == t18 {
																			{
																				{
																					t20 := int32(load32(m.memory[uint32(v28+i32(-4)):]))
																					v5 = t20
																					if v5 != 0 {
																						goto l7
																					}
																					v25 = i32(1)
																					goto l8
																				}
																			l7:
																				t21 := int32(load32(m.memory[uint32(v28+i32(-8)):]))
																				v21 = t21
																				t22 := m.fn5(v5)
																				v25 = t22
																				if v25 == 0 {
																					m.fn10(i32(1), v5)
																					panic("unreachable")
																				}
																				if v5 == 0 {
																					goto l8
																				}
																				memory_copy(m.memory, uint32(v25), uint32(v21), uint32(v5))
																			}
																		l8:
																			store32(m.memory[int64(uint32(v4))+292:], uint32(v5))
																			store32(m.memory[int64(uint32(v4))+288:], uint32(v25))
																			store32(m.memory[int64(uint32(v4))+284:], uint32(v5))
																			store32(m.memory[int64(uint32(v4))+280:], uint32(i32(7)))
																			m.fn557(v4 + i32(96))
																			{
																				{
																					t23 := int32(load32(m.memory[int64(uint32(v4))+116:]))
																					v5 = t23
																					if v5 != 0 {
																						goto l10
																					}
																					{
																						t24 := int32(load32(m.memory[int64(uint32(v4))+104:]))
																						v25 = t24
																						t25 := int32(load32(m.memory[int64(uint32(v4))+96:]))
																						if v25 != t25 {
																							goto l11
																						}
																						m.fn318(v4 + i32(96))
																					}
																				l11:
																					t26 := int32(load32(m.memory[int64(uint32(v4))+100:]))
																					v5 = t26 + v25*i32(28)
																					t27 := int64(load64(m.memory[int64(uint32(v4))+280:]))
																					store64(m.memory[uint32(v5):], uint64(t27))
																					t28 := int64(load64(m.memory[int64(uint32(v4))+288:]))
																					store64(m.memory[int64(uint32(v5))+8:], uint64(t28))
																					t29 := int64(load64(m.memory[int64(uint32(v4))+296:]))
																					store64(m.memory[int64(uint32(v5))+16:], uint64(t29))
																					t30 := int32(load32(m.memory[int64(uint32(v4))+304:]))
																					store32(m.memory[int64(uint32(v5))+24:], uint32(t30))
																					store32(m.memory[int64(uint32(v4))+104:], uint32(v25+i32(1)))
																					goto l12
																				}
																			l10:
																				{
																					t31 := int32(load32(m.memory[int64(uint32(v4))+112:]))
																					v5 = t31 + v5*i32(28)
																					t32 := int32(m.memory[uint32(v5+i32(-4))])
																					if t32 != 0 {
																						goto l13
																					}
																					m.fn337(v4 + i32(280))
																					goto l12
																				}
																			l13:
																				{
																					v21 = v5 + i32(-8)
																					t33 := int32(load32(m.memory[uint32(v21):]))
																					v25 = t33
																					t34 := v25
																					v19 = v5 + i32(-16)
																					t35 := int32(load32(m.memory[uint32(v19):]))
																					if t34 != t35 {
																						goto l14
																					}
																					m.fn318(v19)
																				}
																			l14:
																				t36 := int32(load32(m.memory[uint32(v5+i32(-12)):]))
																				v5 = t36 + v25*i32(28)
																				t37 := int32(load32(m.memory[int64(uint32(v4))+304:]))
																				store32(m.memory[int64(uint32(v5))+24:], uint32(t37))
																				t38 := int64(load64(m.memory[int64(uint32(v4))+296:]))
																				store64(m.memory[int64(uint32(v5))+16:], uint64(t38))
																				t39 := int64(load64(m.memory[int64(uint32(v4))+288:]))
																				store64(m.memory[int64(uint32(v5))+8:], uint64(t39))
																				t40 := int64(load64(m.memory[int64(uint32(v4))+280:]))
																				store64(m.memory[uint32(v5):], uint64(t40))
																				store32(m.memory[uint32(v21):], uint32(v25+i32(1)))
																			}
																		l12:
																			v2 = v2 + i32(1)
																			t41 := int32(load32(m.memory[int64(uint32(v1))+96:]))
																			t42 := v2
																			v5 = t41
																			p43 := v3
																			if uint32(v5) < uint32(v3) {
																				p43 = v5
																			}
																			if uint32(t42) < uint32(p43) {
																				goto l15
																			}
																			goto l0
																		}
																		v22 = (v22 + i64(-1)) & v22
																		if !(v22 == 0) {
																			goto l5
																		}
																	}
																}
															l3:
																if !(v27&(v27<<1)&i64(-0x7f7f7f7f7f7f7f80) == 0) {
																	goto l2
																}
																t19 := v5
																v26 = v26 + i32(8)
																v5 = (t19 + v26) & v23
																goto l6
															}
														}
													l2:
														switch v20 + i32(-1) {
														case 1, 4, 7, 30:
															goto l17
														default:
															if uint32(v20) < uint32(i32(32)) {
																goto l17
															}
															if uint32(v20+i32(-127)) < uint32(i32(33)) {
																goto l17
															}
															{
																t44 := m.fn558(v1, v21, v2)
																v5 = t44
																t45 := int32(m.memory[int64(uint32(v4))+132])
																if v5&i32(1) != t45 {
																	goto l26
																}
																t46 := int32(m.memory[int64(uint32(v4))+133])
																if (int32(uint32(v5&i32(256))>>8)^t46)&i32(1) != 0 {
																	goto l26
																}
																t47 := int32(m.memory[int64(uint32(v4))+134])
																if (int32(uint32(v5&i32(65536))>>16)^t47)&i32(1) != 0 {
																	goto l26
																}
																t48 := int32(m.memory[int64(uint32(v4))+135])
																if (int32(uint32(v5&i32(0x1000000))>>24)^t48)&i32(1) == 0 {
																	goto l27
																}
															}
														l26:
															m.fn557(v4 + i32(96))
															store32(m.memory[int64(uint32(v4))+132:], uint32(v5))
														l27:
															t49 := int32(load32(m.memory[int64(uint32(v4))+128:]))
															v5 = t49
															var p50 int32
															if uint32(v20) < uint32(i32(128)) {
																p50 = 1
															}
															v19 = p50
															if v19 == 0 {
																goto l28
															}
															v25 = i32(1)
															goto l29
														case 6, 11, 12, 13:
															t51 := int32(load32(m.memory[uint32(v18):]))
															v26 = t51
															v5 = i32(0)
															{
																t52 := int32(load32(m.memory[uint32(v17):]))
																v25 = t52
																switch v25 {
																case 0:
																	goto l30
																default:
																	v5 = i32(0)
																l33:
																	{
																		t53 := v5
																		v28 = int32(uint32(v25) >> 1)
																		v23 = v28 + v5
																		t54 := int32(load32(m.memory[uint32(v26+v23*i32(72)+i32(64)):]))
																		p55 := v23
																		if uint32(t54) > uint32(v21) {
																			p55 = t53
																		}
																		v5 = p55
																		v25 = v25 - v28
																		if uint32(v25) > uint32(i32(1)) {
																			goto l33
																		}
																	}
																	fallthrough
																case 1:
																	t56 := int32(load32(m.memory[uint32(v26+v5*i32(72)+i32(64)):]))
																	t57 := v5
																	var p58 int32
																	if uint32(t56) <= uint32(v21) {
																		p58 = 1
																	}
																	v5 = t57 + p58
																	if v5 == 0 {
																		goto l30
																	}
																	v5 = v26 + v5*i32(72)
																	v25 = v5 + i32(-72)
																	if v25 == 0 {
																		goto l30
																	}
																	t59 := int32(load32(m.memory[uint32(v5+i32(-4)):]))
																	if uint32(v21) >= uint32(t59) {
																		goto l30
																	}
																	v21 = i32(-1)
																	t60 := int32(m.memory[uint32(v5+i32(-25))])
																	v29 = t60
																	t61 := int32(m.memory[uint32(v5+i32(-26))])
																	v30 = t61
																	t62 := int32(load32(m.memory[uint32(v5+i32(-68)):]))
																	v31 = t62
																	t63 := int32(m.memory[uint32(v5+i32(-31))])
																	v32 = t63
																	t64 := int32(m.memory[uint32(v5+i32(-32))])
																	v33 = t64
																	t65 := int32(load16(m.memory[uint32(v5+i32(-34)):]))
																	v34 = t65
																	t66 := int32(load16(m.memory[uint32(v5+i32(-36)):]))
																	v35 = t66
																	t67 := int32(m.memory[uint32(v5+i32(-29))])
																	v36 = t67
																	t68 := int32(m.memory[uint32(v5+i32(-30))])
																	v37 = t68
																	t69 := int32(m.memory[uint32(v5+i32(-27))])
																	v38 = t69
																	t70 := int32(m.memory[uint32(v5+i32(-28))])
																	v39 = t70
																	t71 := int32(load16(m.memory[uint32(v5+i32(-12)):]))
																	v28 = t71
																	t72 := int32(load32(m.memory[uint32(v25):]))
																	v40 = t72
																	{
																		t73 := int32(load32(m.memory[uint32(v5+i32(-64)):]))
																		if t73 == i32(-1) {
																			goto l34
																		}
																		{
																			{
																				t74 := int32(load32(m.memory[uint32(v5+i32(-56)):]))
																				v21 = t74
																				if v21 != 0 {
																					goto l35
																				}
																				v22 = i64(2)
																				goto l36
																			}
																		l35:
																			t75 := int32(load32(m.memory[uint32(v5+i32(-60)):]))
																			v26 = t75
																			v25 = v21 << 1
																			t76 := m.fn5(v25)
																			v23 = t76
																			if v23 == 0 {
																				m.fn10(i32(2), v25)
																				panic("unreachable")
																			}
																			if v25 == 0 {
																				goto l38
																			}
																			memory_copy(m.memory, uint32(v23), uint32(v26), uint32(v25))
																		l38:
																			v22 = int64(uint32(v23))
																		}
																	l36:
																		{
																			{
																				t77 := int32(load32(m.memory[uint32(v5+i32(-44)):]))
																				v25 = t77
																				if v25 != 0 {
																					goto l39
																				}
																				v23 = i32(1)
																				goto l40
																			}
																		l39:
																			t78 := int32(load32(m.memory[uint32(v5+i32(-48)):]))
																			v41 = t78
																			v26 = v25 << 2
																			t79 := m.fn5(v26)
																			v23 = t79
																			if v23 == 0 {
																				m.fn10(i32(1), v26)
																				panic("unreachable")
																			}
																			if v26 == 0 {
																				goto l40
																			}
																			memory_copy(m.memory, uint32(v23), uint32(v41), uint32(v26))
																		}
																	l40:
																		v22 = v22 | int64(uint32(v21))<<32
																		t80 := int32(m.memory[uint32(v5+i32(-40))])
																		v26 = t80
																	}
																l34:
																	m.memory[int64(uint32(v4))+181] = byte(v38)
																	m.memory[int64(uint32(v4))+180] = byte(v39)
																	m.memory[int64(uint32(v4))+179] = byte(v36)
																	m.memory[int64(uint32(v4))+178] = byte(v37)
																	m.memory[int64(uint32(v4))+177] = byte(v32)
																	m.memory[int64(uint32(v4))+176] = byte(v33)
																	store16(m.memory[int64(uint32(v4))+174:], uint16(v34))
																	store16(m.memory[int64(uint32(v4))+172:], uint16(v35))
																	m.memory[int64(uint32(v4))+183] = byte(v29)
																	m.memory[int64(uint32(v4))+182] = byte(v30)
																	store32(m.memory[int64(uint32(v4))+140:], uint32(v31))
																	store32(m.memory[int64(uint32(v4))+136:], uint32(v40))
																	m.memory[int64(uint32(v4))+168] = byte(v26)
																	store32(m.memory[int64(uint32(v4))+164:], uint32(v25))
																	store32(m.memory[int64(uint32(v4))+160:], uint32(v23))
																	store32(m.memory[int64(uint32(v4))+156:], uint32(v25))
																	store64(m.memory[int64(uint32(v4))+148:], uint64(v22))
																	store32(m.memory[int64(uint32(v4))+144:], uint32(v21))
																	goto l42
																}
															}
														l30:
															m.memory[int64(uint32(v4))+178] = byte(i32(2))
															v28 = i32(0)
															m.memory[int64(uint32(v4))+176] = byte(i32(0))
															store16(m.memory[int64(uint32(v4))+172:], uint16(i32(0)))
															store32(m.memory[int64(uint32(v4))+180:], uint32(i32(33686018)))
															store32(m.memory[int64(uint32(v4))+136:], uint32(i32(0)))
															store32(m.memory[int64(uint32(v4))+144:], uint32(i32(-1)))
														l42:
															{
																{
																	t81 := int32(load32(m.memory[int64(uint32(v1))+68:]))
																	if t81 != 0 {
																		goto l43
																	}
																	v5 = i32(0)
																	goto l44
																}
															l43:
																t82 := int64(load64(m.memory[int64(uint32(v1))+72:]))
																t83 := int64(load64(m.memory[int64(uint32(v1))+80:]))
																t84 := m.fn106(t82, t83, v28)
																v22 = t84
																t85 := int32(load32(m.memory[int64(uint32(v1))+60:]))
																v23 = t85
																v25 = v23 & int32(v22)
																v24 = int64(uint64(v22)>>25) & i64(127) * i64(72340172838076673)
																t86 := int32(load32(m.memory[int64(uint32(v1))+56:]))
																v21 = t86
																v26 = i32(0)
															l48:
																{
																	{
																		t87 := int64(load64(m.memory[uint32(v21+v25):]))
																		v27 = t87
																		v22 = v27 ^ v24
																		v22 = (v22 ^ i64(-1)) & (v22 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
																		if v22 == 0 {
																			goto l45
																		}
																	l47:
																		{
																			t88 := v28 & i32(0xffff)
																			v5 = v21 + (i32(0)-(int32(uint32(int64(bits.TrailingZeros64(uint64(v22))))>>3)+v25)&v23)*i32(60)
																			t89 := int32(load16(m.memory[uint32(v5+i32(-60)):]))
																			if t88 == t89 {
																				goto l46
																			}
																			v22 = (v22 + i64(-1)) & v22
																			if !(v22 == 0) {
																				goto l47
																			}
																		}
																	}
																l45:
																	v5 = i32(0)
																	if !(v27&(v27<<1)&i64(-0x7f7f7f7f7f7f7f80) == 0) {
																		goto l46
																	}
																	t90 := v25
																	v26 = v26 + i32(8)
																	v25 = (t90 + v26) & v23
																	goto l48
																}
															l46:
																p91 := i32(0)
																if v5 != 0 {
																	p91 = v5 + i32(-56)
																}
																v5 = p91
															}
														l44:
															v21 = i32(-1)
															p92 := v1
															if v5 != 0 {
																p92 = v5
															}
															v5 = p92
															t93 := int32(m.memory[int64(uint32(v5))+47])
															v29 = t93
															t94 := int32(m.memory[int64(uint32(v5))+46])
															v30 = t94
															t95 := int32(load32(m.memory[int64(uint32(v5))+4:]))
															v31 = t95
															t96 := int32(load32(m.memory[uint32(v5):]))
															v32 = t96
															t97 := int32(m.memory[int64(uint32(v5))+41])
															v33 = t97
															t98 := int32(m.memory[int64(uint32(v5))+40])
															v34 = t98
															t99 := int32(load16(m.memory[int64(uint32(v5))+38:]))
															v35 = t99
															t100 := int32(load16(m.memory[int64(uint32(v5))+36:]))
															v36 = t100
															t101 := int32(m.memory[int64(uint32(v5))+43])
															v37 = t101
															t102 := int32(m.memory[int64(uint32(v5))+42])
															v38 = t102
															t103 := int32(m.memory[int64(uint32(v5))+45])
															v39 = t103
															t104 := int32(m.memory[int64(uint32(v5))+44])
															v40 = t104
															{
																t105 := int32(load32(m.memory[int64(uint32(v5))+8:]))
																if t105 == i32(-1) {
																	goto l49
																}
																{
																	{
																		t106 := int32(load32(m.memory[int64(uint32(v5))+16:]))
																		v21 = t106
																		if v21 != 0 {
																			goto l50
																		}
																		v22 = i64(2)
																		goto l51
																	}
																l50:
																	t107 := int32(load32(m.memory[int64(uint32(v5))+12:]))
																	v26 = t107
																	v25 = v21 << 1
																	t108 := m.fn5(v25)
																	v23 = t108
																	if v23 == 0 {
																		m.fn10(i32(2), v25)
																		panic("unreachable")
																	}
																	if v25 == 0 {
																		goto l53
																	}
																	memory_copy(m.memory, uint32(v23), uint32(v26), uint32(v25))
																l53:
																	v22 = int64(uint32(v23))
																}
															l51:
																{
																	{
																		t109 := int32(load32(m.memory[int64(uint32(v5))+28:]))
																		v25 = t109
																		if v25 != 0 {
																			goto l54
																		}
																		v23 = i32(1)
																		goto l55
																	}
																l54:
																	t110 := int32(load32(m.memory[int64(uint32(v5))+24:]))
																	v41 = t110
																	v26 = v25 << 2
																	t111 := m.fn5(v26)
																	v23 = t111
																	if v23 == 0 {
																		m.fn10(i32(1), v26)
																		panic("unreachable")
																	}
																	if v26 == 0 {
																		goto l55
																	}
																	memory_copy(m.memory, uint32(v23), uint32(v41), uint32(v26))
																}
															l55:
																v22 = v22 | int64(uint32(v21))<<32
																t112 := int32(m.memory[int64(uint32(v5))+32])
																v26 = t112
															}
														l49:
															m.memory[int64(uint32(v4))+325] = byte(v39)
															m.memory[int64(uint32(v4))+324] = byte(v40)
															m.memory[int64(uint32(v4))+323] = byte(v37)
															m.memory[int64(uint32(v4))+322] = byte(v38)
															m.memory[int64(uint32(v4))+321] = byte(v33)
															m.memory[int64(uint32(v4))+320] = byte(v34)
															store16(m.memory[int64(uint32(v4))+318:], uint16(v35))
															store16(m.memory[int64(uint32(v4))+316:], uint16(v36))
															m.memory[int64(uint32(v4))+327] = byte(v29)
															m.memory[int64(uint32(v4))+326] = byte(v30)
															store32(m.memory[int64(uint32(v4))+284:], uint32(v31))
															store32(m.memory[int64(uint32(v4))+280:], uint32(v32))
															m.memory[int64(uint32(v4))+312] = byte(v26)
															store32(m.memory[int64(uint32(v4))+308:], uint32(v25))
															store32(m.memory[int64(uint32(v4))+304:], uint32(v23))
															store32(m.memory[int64(uint32(v4))+300:], uint32(v25))
															store64(m.memory[int64(uint32(v4))+292:], uint64(v22))
															store32(m.memory[int64(uint32(v4))+288:], uint32(v21))
															m.fn547(v4+i32(184), v4+i32(280), v4+i32(136))
															{
																t113 := int32(load32(m.memory[int64(uint32(v1))+132:]))
																if uint32(v2) >= uint32(t113) {
																	goto l57
																}
																t114 := int32(load32(m.memory[int64(uint32(v1))+128:]))
																t115 := int32(load32(m.memory[uint32(t114+v19):]))
																v5 = t115
																t116 := int32(load32(m.memory[int64(uint32(v1))+372:]))
																if uint32(v5) >= uint32(t116) {
																	goto l57
																}
																t117 := int32(load32(m.memory[int64(uint32(v1))+368:]))
																v5 = t117 + v5<<3
																t118 := int32(load32(m.memory[uint32(v5):]))
																if t118 != i32(1) {
																	goto l57
																}
																t119 := int32(load32(m.memory[int64(uint32(v5))+4:]))
																v5 = t119
																t120 := int32(load32(m.memory[int64(uint32(v1))+360:]))
																if uint32(v5) >= uint32(t120) {
																	goto l57
																}
																m.memory[int64(uint32(v4))+274] = byte(i32(2))
																m.memory[int64(uint32(v4))+272] = byte(i32(0))
																store16(m.memory[int64(uint32(v4))+268:], uint16(i32(0)))
																store32(m.memory[int64(uint32(v4))+276:], uint32(i32(33686018)))
																store32(m.memory[int64(uint32(v4))+232:], uint32(i32(0)))
																store32(m.memory[int64(uint32(v4))+240:], uint32(i32(-1)))
																t121 := int32(load32(m.memory[int64(uint32(v1))+356:]))
																v5 = t121 + v5*i32(12)
																t122 := int32(load32(m.memory[int64(uint32(v5))+4:]))
																t123 := int32(load32(m.memory[int64(uint32(v5))+8:]))
																m.fn546(t122, t123, i32(1), i32(0), v4+i32(232))
																m.fn547(v4+i32(280), v4+i32(184), v4+i32(232))
																t124 := int64(load64(m.memory[int64(uint32(v4))+320:]))
																store64(m.memory[int64(uint32(v4))+224:], uint64(t124))
																t125 := int64(load64(m.memory[int64(uint32(v4))+312:]))
																store64(m.memory[int64(uint32(v4))+216:], uint64(t125))
																t126 := int64(load64(m.memory[int64(uint32(v4))+304:]))
																store64(m.memory[int64(uint32(v4))+208:], uint64(t126))
																t127 := int64(load64(m.memory[int64(uint32(v4))+296:]))
																store64(m.memory[int64(uint32(v4))+200:], uint64(t127))
																t128 := int64(load64(m.memory[int64(uint32(v4))+288:]))
																store64(m.memory[int64(uint32(v4))+192:], uint64(t128))
																t129 := int64(load64(m.memory[int64(uint32(v4))+280:]))
																store64(m.memory[int64(uint32(v4))+184:], uint64(t129))
															}
														l57:
															t130 := int32(m.memory[int64(uint32(v4))+231])
															v37 = t130
															t131 := int32(m.memory[int64(uint32(v4))+230])
															v35 = t131
															t132 := int32(m.memory[int64(uint32(v4))+229])
															v41 = t132
															t133 := int32(m.memory[int64(uint32(v4))+228])
															v19 = t133
															t134 := int32(m.memory[int64(uint32(v4))+227])
															v42 = t134
															t135 := int32(m.memory[int64(uint32(v4))+226])
															v38 = t135
															t136 := int32(m.memory[int64(uint32(v4))+225])
															v43 = t136
															t137 := int32(m.memory[int64(uint32(v4))+224])
															v44 = t137
															t138 := int32(load16(m.memory[int64(uint32(v4))+222:]))
															v39 = t138
															t139 := int32(load16(m.memory[int64(uint32(v4))+220:]))
															v40 = t139
															t140 := int32(m.memory[int64(uint32(v4))+216])
															v45 = t140
															t141 := int32(load32(m.memory[int64(uint32(v4))+212:]))
															v46 = t141
															t142 := int32(load32(m.memory[int64(uint32(v4))+208:]))
															v30 = t142
															t143 := int32(load32(m.memory[int64(uint32(v4))+204:]))
															v29 = t143
															t144 := int32(load32(m.memory[int64(uint32(v4))+200:]))
															v47 = t144
															t145 := int32(load32(m.memory[int64(uint32(v4))+196:]))
															v32 = t145
															t146 := int32(load32(m.memory[int64(uint32(v4))+192:]))
															v23 = t146
															t147 := int32(load32(m.memory[int64(uint32(v4))+188:]))
															v33 = t147
															t148 := int32(load32(m.memory[int64(uint32(v4))+184:]))
															v31 = t148
															t149 := int64(load64(m.memory[int64(uint32(v4))+96:]))
															v22 = t149
															store64(m.memory[int64(uint32(v4))+96:], uint64(i64(0x400000000)))
															t150 := int64(load64(m.memory[int64(uint32(v4))+104:]))
															v24 = t150
															store64(m.memory[int64(uint32(v4))+104:], uint64(i64(0)))
															t151 := int64(load64(m.memory[int64(uint32(v4))+112:]))
															v27 = t151
															store64(m.memory[int64(uint32(v4))+112:], uint64(i64(4)))
															t152 := int64(load64(m.memory[int64(uint32(v4))+120:]))
															v48 = t152
															store64(m.memory[int64(uint32(v4))+120:], uint64(i64(0x100000000)))
															t153 := int64(load64(m.memory[int64(uint32(v4))+128:]))
															v49 = t153
															store64(m.memory[int64(uint32(v4))+128:], uint64(i64(0)))
															store64(m.memory[int64(uint32(v4))+312:], uint64(v49))
															store64(m.memory[int64(uint32(v4))+304:], uint64(v48))
															store64(m.memory[int64(uint32(v4))+296:], uint64(v27))
															store64(m.memory[int64(uint32(v4))+288:], uint64(v24))
															store64(m.memory[int64(uint32(v4))+280:], uint64(v22))
															m.fn557(v4 + i32(280))
															t154 := int32(load32(m.memory[int64(uint32(v4))+300:]))
															if t154 != 0 {
															l119:
																{
																	m.fn559(v4 + i32(280))
																	t282 := int32(load32(m.memory[int64(uint32(v4))+300:]))
																	if t282 != 0 {
																		goto l119
																	}
																	goto l59
																}
															}
															goto l59
														case 10:
															store32(m.memory[int64(uint32(v4))+280:], uint32(i32(8)))
															m.fn557(v4 + i32(96))
															{
																t155 := int32(load32(m.memory[int64(uint32(v4))+116:]))
																v5 = t155
																if v5 != 0 {
																	t163 := int32(load32(m.memory[int64(uint32(v4))+112:]))
																	v5 = t163 + v5*i32(28)
																	t164 := int32(m.memory[uint32(v5+i32(-4))])
																	if t164 != 0 {
																		{
																			v21 = v5 + i32(-8)
																			t165 := int32(load32(m.memory[uint32(v21):]))
																			v25 = t165
																			t166 := v25
																			v19 = v5 + i32(-16)
																			t167 := int32(load32(m.memory[uint32(v19):]))
																			if t166 != t167 {
																				goto l63
																			}
																			m.fn318(v19)
																		}
																	l63:
																		t168 := int32(load32(m.memory[uint32(v5+i32(-12)):]))
																		v5 = t168 + v25*i32(28)
																		t169 := int32(load32(m.memory[int64(uint32(v4))+304:]))
																		store32(m.memory[int64(uint32(v5))+24:], uint32(t169))
																		t170 := int64(load64(m.memory[int64(uint32(v4))+296:]))
																		store64(m.memory[int64(uint32(v5))+16:], uint64(t170))
																		t171 := int64(load64(m.memory[int64(uint32(v4))+288:]))
																		store64(m.memory[int64(uint32(v5))+8:], uint64(t171))
																		t172 := int64(load64(m.memory[int64(uint32(v4))+280:]))
																		store64(m.memory[uint32(v5):], uint64(t172))
																		store32(m.memory[uint32(v21):], uint32(v25+i32(1)))
																		goto l17
																	}
																	m.fn337(v4 + i32(280))
																	goto l17
																}
																{
																	t156 := int32(load32(m.memory[int64(uint32(v4))+104:]))
																	v25 = t156
																	t157 := int32(load32(m.memory[int64(uint32(v4))+96:]))
																	if v25 != t157 {
																		goto l61
																	}
																	m.fn318(v4 + i32(96))
																}
															l61:
																t158 := int32(load32(m.memory[int64(uint32(v4))+100:]))
																v5 = t158 + v25*i32(28)
																t159 := int64(load64(m.memory[int64(uint32(v4))+280:]))
																store64(m.memory[uint32(v5):], uint64(t159))
																t160 := int64(load64(m.memory[int64(uint32(v4))+288:]))
																store64(m.memory[int64(uint32(v5))+8:], uint64(t160))
																t161 := int64(load64(m.memory[int64(uint32(v4))+296:]))
																store64(m.memory[int64(uint32(v5))+16:], uint64(t161))
																t162 := int32(load32(m.memory[int64(uint32(v4))+304:]))
																store32(m.memory[int64(uint32(v5))+24:], uint32(t162))
																store32(m.memory[int64(uint32(v4))+104:], uint32(v25+i32(1)))
																goto l17
															}
														case 18:
															m.fn557(v4 + i32(96))
															{
																t173 := int32(load32(m.memory[int64(uint32(v4))+116:]))
																v25 = t173
																t174 := int32(load32(m.memory[int64(uint32(v4))+108:]))
																if v25 != t174 {
																	goto l64
																}
																m.fn318(v6)
															}
														l64:
															t175 := int32(load32(m.memory[int64(uint32(v4))+112:]))
															v5 = t175 + v25*i32(28)
															m.memory[int64(uint32(v5))+24] = byte(i32(0))
															store64(m.memory[int64(uint32(v5))+16:], uint64(i64(4)))
															store64(m.memory[int64(uint32(v5))+8:], uint64(i64(0)))
															store64(m.memory[uint32(v5):], uint64(i64(0x100000000)))
															store32(m.memory[int64(uint32(v4))+116:], uint32(v25+i32(1)))
															goto l17
														case 19:
															m.fn557(v4 + i32(96))
															t176 := int32(load32(m.memory[int64(uint32(v4))+116:]))
															v5 = t176
															if v5 == 0 {
																goto l17
															}
															t177 := int32(load32(m.memory[int64(uint32(v4))+112:]))
															m.memory[uint32(t177+v5*i32(28)+i32(-4))] = byte(i32(1))
															goto l17
														case 20:
															m.fn559(v4 + i32(96))
															goto l17
														case 8:
															{
																t178 := m.fn558(v1, v21, v2)
																v5 = t178
																t179 := int32(m.memory[int64(uint32(v4))+132])
																if v5&i32(1) != t179 {
																	goto l65
																}
																t180 := int32(m.memory[int64(uint32(v4))+133])
																if (int32(uint32(v5&i32(256))>>8)^t180)&i32(1) != 0 {
																	goto l65
																}
																t181 := int32(m.memory[int64(uint32(v4))+134])
																if (int32(uint32(v5&i32(65536))>>16)^t181)&i32(1) != 0 {
																	goto l65
																}
																t182 := int32(m.memory[int64(uint32(v4))+135])
																if (int32(uint32(v5&i32(0x1000000))>>24)^t182)&i32(1) == 0 {
																	goto l66
																}
															}
														l65:
															m.fn557(v4 + i32(96))
															store32(m.memory[int64(uint32(v4))+132:], uint32(v5))
														l66:
															{
																t183 := int32(load32(m.memory[int64(uint32(v4))+120:]))
																t184 := int32(load32(m.memory[int64(uint32(v4))+128:]))
																v5 = t184
																if t183 != v5 {
																	goto l67
																}
																m.fn197(v4+i32(96)+i32(24), v5, i32(1), i32(1), i32(1))
															}
														l67:
															t185 := int32(load32(m.memory[int64(uint32(v4))+124:]))
															m.memory[uint32(t185+v5)] = byte(i32(32))
															store32(m.memory[int64(uint32(v4))+128:], uint32(v5+i32(1)))
															goto l17
														case 29:
															{
																t186 := m.fn558(v1, v21, v2)
																v5 = t186
																t187 := int32(m.memory[int64(uint32(v4))+132])
																if v5&i32(1) != t187 {
																	goto l68
																}
																t188 := int32(m.memory[int64(uint32(v4))+133])
																if (int32(uint32(v5&i32(256))>>8)^t188)&i32(1) != 0 {
																	goto l68
																}
																t189 := int32(m.memory[int64(uint32(v4))+134])
																if (int32(uint32(v5&i32(65536))>>16)^t189)&i32(1) != 0 {
																	goto l68
																}
																t190 := int32(m.memory[int64(uint32(v4))+135])
																if (int32(uint32(v5&i32(0x1000000))>>24)^t190)&i32(1) == 0 {
																	goto l69
																}
															}
														l68:
															m.fn557(v4 + i32(96))
															store32(m.memory[int64(uint32(v4))+132:], uint32(v5))
														l69:
															{
																t191 := int32(load32(m.memory[int64(uint32(v4))+120:]))
																t192 := int32(load32(m.memory[int64(uint32(v4))+128:]))
																v5 = t192
																if t191 != v5 {
																	goto l70
																}
																m.fn197(v4+i32(96)+i32(24), v5, i32(1), i32(1), i32(1))
															}
														l70:
															t193 := int32(load32(m.memory[int64(uint32(v4))+124:]))
															m.memory[uint32(t193+v5)] = byte(i32(45))
															store32(m.memory[int64(uint32(v4))+128:], uint32(v5+i32(1)))
															goto l17
														case 0:
															t194 := int32(load32(m.memory[uint32(v16):]))
															v23 = t194
															v5 = i32(0)
															{
																t195 := int32(load32(m.memory[uint32(v15):]))
																v25 = t195
																switch v25 {
																case 0:
																	goto l17
																default:
																	v5 = i32(0)
																l73:
																	{
																		t196 := v5
																		v19 = int32(uint32(v25) >> 1)
																		v28 = v19 + v5
																		t197 := int32(load32(m.memory[uint32(v23+v28*i32(72)+i32(64)):]))
																		p198 := v28
																		if uint32(t197) > uint32(v21) {
																			p198 = t196
																		}
																		v5 = p198
																		v25 = v25 - v19
																		if uint32(v25) > uint32(i32(1)) {
																			goto l73
																		}
																	}
																	fallthrough
																case 1:
																	t199 := int32(load32(m.memory[uint32(v23+v5*i32(72)+i32(64)):]))
																	t200 := v5
																	var p201 int32
																	if uint32(t199) <= uint32(v21) {
																		p201 = 1
																	}
																	v5 = t200 + p201
																	if v5 == 0 {
																		goto l17
																	}
																	v25 = v23 + v5*i32(72)
																	if v25 == i32(72) {
																		goto l17
																	}
																	t202 := int32(load32(m.memory[uint32(v25+i32(-4)):]))
																	if uint32(v21) >= uint32(t202) {
																		goto l17
																	}
																	t203 := int32(load32(m.memory[uint32(v25+i32(-16)):]))
																	v5 = t203
																	if uint32(v5) < uint32(i32(2)) {
																		goto l17
																	}
																	t204 := int32(load32(m.memory[uint32(v25+i32(-20)):]))
																	v19 = t204
																	v25 = i32(0)
																	v21 = i32(2)
																	v30 = i32(0)
																l87:
																	{
																		{
																			if uint32(v25) >= uint32(v5) {
																				m.fn33(v25, v5, i32(1069384))
																				panic("unreachable")
																			}
																			v28 = v25 + i32(1)
																			if uint32(v28) >= uint32(v5) {
																				m.fn33(v28, v5, i32(1069400))
																				panic("unreachable")
																			}
																			t205 := int32(m.memory[uint32(v19+v28)])
																			v23 = t205
																			t206 := int32(m.memory[uint32(v19+v25)])
																			v28 = v23<<8 | t206
																			v26 = v19 + v21
																			v20 = v5 - v21
																			v25 = i32(1)
																			switch int32(uint32(v23) >> 5) {
																			default:
																				goto l76
																			case 2, 4, 5:
																				v25 = i32(2)
																				goto l76
																			case 3:
																				v25 = i32(4)
																				goto l76
																			case 7:
																				v25 = i32(3)
																				goto l76
																			case 6:
																				if v28 == i32(54792) {
																					goto l81
																				}
																				if v5 == v21 {
																					var p208 int32
																					if v28 != i32(27139) {
																						p208 = 1
																					}
																					v30 = p208 & v30
																					v25 = i32(0)
																					goto l84
																				}
																				t207 := int32(m.memory[uint32(v26)])
																				v25 = t207 + i32(1)
																			}
																		l76:
																			if uint32(v25) > uint32(v20) {
																				goto l83
																			}
																			if v28 != i32(27139) {
																				goto l84
																			}
																			if uint32(v25) >= uint32(i32(4)) {
																				t209 := int32(load32(m.memory[uint32(v26):]))
																				v29 = t209
																				v30 = i32(1)
																				goto l84
																			}
																			v30 = i32(0)
																			goto l84
																		}
																	l81:
																		if uint32(v20) > uint32(i32(1)) {
																			goto l86
																		}
																		v25 = i32(0)
																		goto l84
																	l86:
																		t210 := int32(load16(m.memory[uint32(v26):]))
																		v25 = t210
																		if uint32(v25) >= uint32(v20) {
																			goto l83
																		}
																		v25 = v25 + i32(1)
																	}
																l84:
																	v25 = v25 + v21
																	v21 = v25 + i32(2)
																	if uint32(v21) <= uint32(v5) {
																		goto l87
																	}
																l83:
																	if v30&i32(1) == 0 {
																		goto l17
																	}
																	store32(m.memory[int64(uint32(v4))+792:], uint32(v29))
																	t211 := int32(load32(m.memory[int64(uint32(v1))+384:]))
																	v5 = t211
																	if uint32(v5) < uint32(v29) {
																		goto l17
																	}
																	if uint32(v5-v29) < uint32(i32(4)) {
																		goto l17
																	}
																	t212 := v5
																	v25 = v29 + i32(4)
																	if uint32(t212) < uint32(v25) {
																		goto l17
																	}
																	if uint32(v5-v25) < uint32(i32(2)) {
																		goto l17
																	}
																	t213 := int32(load32(m.memory[int64(uint32(v1))+380:]))
																	t214 := v29
																	v19 = t213
																	v23 = v19 + v29
																	t215 := int32(load32(m.memory[uint32(v23):]))
																	v21 = t214 + t215
																	p216 := v21
																	if uint32(v21) < uint32(v29) {
																		p216 = i32(-1)
																	}
																	v21 = p216
																	if uint32(v21) > uint32(v5) {
																		goto l17
																	}
																	t217 := int32(load16(m.memory[uint32(v19+v25):]))
																	v5 = t217
																	{
																		t218 := m.fn5(i32(8))
																		v28 = t218
																		if v28 == 0 {
																			m.fn24(i32(4), i32(8))
																			panic("unreachable")
																		}
																		t219 := v23
																		v25 = v21 - v29
																		p220 := v5
																		if uint32(v25) < uint32(v5) {
																			p220 = v25
																		}
																		v5 = p220
																		v30 = t219 + v5
																		t221 := v28
																		v26 = v25 - v5
																		store32(m.memory[int64(uint32(t221))+4:], uint32(v26))
																		v31 = i32(0)
																		store32(m.memory[uint32(v28):], uint32(i32(0)))
																		v5 = i32(1)
																		store32(m.memory[int64(uint32(v4))+240:], uint32(i32(1)))
																		store32(m.memory[int64(uint32(v4))+236:], uint32(v28))
																		store32(m.memory[int64(uint32(v4))+232:], uint32(i32(1)))
																	l103:
																		{
																			{
																				{
																					{
																						t222 := v28
																						v19 = v5 + i32(-1)
																						v25 = t222 + v19<<3
																						t223 := int32(load32(m.memory[int64(uint32(v25))+4:]))
																						v21 = t223
																						t224 := int32(load32(m.memory[uint32(v25):]))
																						t225 := v21
																						v25 = t224
																						if uint32(t225) <= uint32(v25) {
																							goto l89
																						}
																						if uint32(v21) > uint32(v26) {
																							m.fn121(i32(0), v21, v26, i32(1076020))
																							panic("unreachable")
																						}
																						v23 = v21 - v25
																						if uint32(v23) < uint32(i32(8)) {
																							goto l89
																						}
																						v21 = v30 + v25
																						t226 := int32(load32(m.memory[int64(uint32(v21))+4:]))
																						v20 = t226
																						if uint32(v20) <= uint32(v23+i32(-8)) {
																							goto l91
																						}
																					}
																				l89:
																					store32(m.memory[int64(uint32(v4))+240:], uint32(v19))
																					goto l92
																				l91:
																					if uint32(v25) > uint32(i32(-9)) {
																						goto l93
																					}
																					t227 := v20
																					v33 = v25 + i32(8)
																					v19 = t227 + v33
																					if uint32(v19) < uint32(v20) {
																						goto l93
																					}
																					t228 := int32(m.memory[int64(uint32(v21))+1])
																					v32 = t228
																					t229 := int32(m.memory[uint32(v21)])
																					v29 = t229
																					t230 := int32(load16(m.memory[int64(uint32(v21))+2:]))
																					v23 = t230
																					t231 := v28
																					v34 = v5 << 3
																					store32(m.memory[uint32(t231+v34+i32(-8)):], uint32(v19))
																					v31 = v31 + i32(1)
																					if uint32(v31) > uint32(i32(10000)) {
																						goto l93
																					}
																					if uint32(v5) > uint32(i32(16)) {
																						goto l93
																					}
																					m.fn444(v4+i32(280), v32<<8|v29, v23, v21+i32(8), v20)
																					{
																						t232 := int32(load32(m.memory[int64(uint32(v4))+280:]))
																						if t232 == i32(-2) {
																							goto l94
																						}
																						v5 = i32(0)
																						goto l95
																					}
																				l94:
																					{
																						if v23&i32(0xffff) != i32(61447) {
																							goto l96
																						}
																						v23 = v19
																						{
																							if uint32(v20) <= uint32(i32(33)) {
																								goto l97
																							}
																							t233 := int32(m.memory[int64(uint32(v21))+41])
																							v23 = v25 + t233 + i32(44)
																						}
																					l97:
																						if uint32(v23) < uint32(v33) {
																							goto l93
																						}
																						if uint32(v23) >= uint32(v19) {
																							goto l92
																						}
																						{
																							t234 := int32(load32(m.memory[int64(uint32(v4))+232:]))
																							if v5 != t234 {
																								goto l98
																							}
																							m.fn297(v4 + i32(232))
																						}
																					l98:
																						t235 := int32(load32(m.memory[int64(uint32(v4))+236:]))
																						v28 = t235
																						v25 = v28 + v34
																						store32(m.memory[int64(uint32(v25))+4:], uint32(v19))
																						store32(m.memory[uint32(v25):], uint32(v23))
																						goto l99
																					}
																				l96:
																					if v29&i32(15) == i32(15) {
																						goto l100
																					}
																				}
																			l92:
																				t236 := int32(load32(m.memory[int64(uint32(v4))+240:]))
																				v5 = t236
																				goto l101
																			}
																		l100:
																			{
																				t237 := int32(load32(m.memory[int64(uint32(v4))+232:]))
																				if v5 != t237 {
																					goto l102
																				}
																				m.fn297(v4 + i32(232))
																				t238 := int32(load32(m.memory[int64(uint32(v4))+236:]))
																				v28 = t238
																			}
																		l102:
																			v25 = v28 + v34
																			store32(m.memory[int64(uint32(v25))+4:], uint32(v19))
																			store32(m.memory[uint32(v25):], uint32(v33))
																		l99:
																			t239 := v4
																			v5 = v5 + i32(1)
																			store32(m.memory[int64(uint32(t239))+240:], uint32(v5))
																		}
																	l101:
																		if v5 != 0 {
																			goto l103
																		}
																	l93:
																		store32(m.memory[int64(uint32(v4))+280:], uint32(i32(-2)))
																		v5 = i32(1)
																	l95:
																		{
																			t240 := int32(load32(m.memory[int64(uint32(v4))+232:]))
																			v25 = t240
																			if v25 == 0 {
																				goto l104
																			}
																			t241 := int32(load32(m.memory[int64(uint32(v4))+236:]))
																			m.fn18(t241, v25<<3, i32(4))
																		}
																	l104:
																		if v5 != 0 {
																			goto l17
																		}
																		t242 := int32(load32(m.memory[int64(uint32(v4))+304:]))
																		store32(m.memory[int64(uint32(v4))+256:], uint32(t242))
																		t243 := int64(load64(m.memory[int64(uint32(v4))+296:]))
																		store64(m.memory[int64(uint32(v4))+248:], uint64(t243))
																		t244 := int64(load64(m.memory[int64(uint32(v4))+288:]))
																		store64(m.memory[int64(uint32(v4))+240:], uint64(t244))
																		t245 := int64(load64(m.memory[int64(uint32(v4))+280:]))
																		store64(m.memory[int64(uint32(v4))+232:], uint64(t245))
																		store64(m.memory[int64(uint32(v4))+288:], uint64(v10))
																		store64(m.memory[int64(uint32(v4))+280:], uint64(v11))
																		m.fn167(v4+i32(136), i32(0x100099), v4+i32(280))
																		t246 := int32(load32(m.memory[int64(uint32(v1))+272:]))
																		if t246 != 0 {
																			m.fn355(i32(1073800))
																			panic("unreachable")
																		}
																		store32(m.memory[int64(uint32(v1))+272:], uint32(i32(-1)))
																		v25 = i32(0)
																		{
																			t247 := int32(load32(m.memory[int64(uint32(v4))+248:]))
																			v5 = t247
																			if v5 < i32(0) {
																				goto l106
																			}
																			if v5 != 0 {
																				goto l107
																			}
																			v25 = i32(0)
																			v50 = i32(1)
																			goto l108
																		l107:
																			t248 := int32(load32(m.memory[int64(uint32(v4))+244:]))
																			v25 = t248
																			t249 := m.fn5(v5)
																			v50 = t249
																			if v50 != 0 {
																				goto l109
																			}
																			v25 = i32(1)
																			v50 = v5
																		}
																	l106:
																		m.fn10(v25, v50)
																		panic("unreachable")
																	l109:
																		if v5 == 0 {
																			goto l110
																		}
																		memory_copy(m.memory, uint32(v50), uint32(v25), uint32(v5))
																	l110:
																		v25 = v5
																	l108:
																		store32(m.memory[int64(uint32(v4))+192:], uint32(v5))
																		store32(m.memory[int64(uint32(v4))+188:], uint32(v50))
																		store32(m.memory[int64(uint32(v4))+184:], uint32(v25))
																		t250 := int32(load32(m.memory[int64(uint32(v4))+236:]))
																		t251 := int32(load32(m.memory[int64(uint32(v4))+240:]))
																		m.fn445(v4+i32(280), v8, v4+i32(184), v4+i32(136), t250, t251)
																		t252 := int32(load32(m.memory[int64(uint32(v4))+284:]))
																		v5 = t252
																		{
																			t253 := int32(load32(m.memory[int64(uint32(v4))+280:]))
																			v25 = t253
																			if v25 == i32(-1) {
																				t261 := int32(load32(m.memory[int64(uint32(v1))+272:]))
																				store32(m.memory[int64(uint32(v1))+272:], uint32(t261+i32(1)))
																				{
																					t262 := int32(load32(m.memory[int64(uint32(v4))+232:]))
																					v25 = t262
																					if v25 < i32(1) {
																						goto l114
																					}
																					t263 := int32(load32(m.memory[int64(uint32(v4))+236:]))
																					m.fn18(t263, v25, i32(1))
																				}
																			l114:
																				store32(m.memory[int64(uint32(v4))+300:], uint32(v5))
																				store32(m.memory[int64(uint32(v4))+296:], uint32(i32(-0x80000000)))
																				store64(m.memory[int64(uint32(v4))+288:], uint64(i64(1)))
																				store64(m.memory[int64(uint32(v4))+280:], uint64(i64(5)))
																				m.fn557(v4 + i32(96))
																				{
																					t264 := int32(load32(m.memory[int64(uint32(v4))+116:]))
																					v5 = t264
																					if v5 != 0 {
																						t272 := int32(load32(m.memory[int64(uint32(v4))+112:]))
																						v5 = t272 + v5*i32(28)
																						t273 := int32(m.memory[uint32(v5+i32(-4))])
																						if t273 != 0 {
																							{
																								v21 = v5 + i32(-8)
																								t274 := int32(load32(m.memory[uint32(v21):]))
																								v25 = t274
																								t275 := v25
																								v19 = v5 + i32(-16)
																								t276 := int32(load32(m.memory[uint32(v19):]))
																								if t275 != t276 {
																									goto l118
																								}
																								m.fn318(v19)
																							}
																						l118:
																							t277 := int32(load32(m.memory[uint32(v5+i32(-12)):]))
																							v5 = t277 + v25*i32(28)
																							t278 := int32(load32(m.memory[int64(uint32(v4))+304:]))
																							store32(m.memory[int64(uint32(v5))+24:], uint32(t278))
																							t279 := int64(load64(m.memory[int64(uint32(v4))+296:]))
																							store64(m.memory[int64(uint32(v5))+16:], uint64(t279))
																							t280 := int64(load64(m.memory[int64(uint32(v4))+288:]))
																							store64(m.memory[int64(uint32(v5))+8:], uint64(t280))
																							t281 := int64(load64(m.memory[int64(uint32(v4))+280:]))
																							store64(m.memory[uint32(v5):], uint64(t281))
																							store32(m.memory[uint32(v21):], uint32(v25+i32(1)))
																							goto l17
																						}
																						m.fn337(v4 + i32(280))
																						goto l17
																					}
																					{
																						t265 := int32(load32(m.memory[int64(uint32(v4))+104:]))
																						v25 = t265
																						t266 := int32(load32(m.memory[int64(uint32(v4))+96:]))
																						if v25 != t266 {
																							goto l116
																						}
																						m.fn318(v4 + i32(96))
																					}
																				l116:
																					t267 := int32(load32(m.memory[int64(uint32(v4))+100:]))
																					v5 = t267 + v25*i32(28)
																					t268 := int64(load64(m.memory[int64(uint32(v4))+280:]))
																					store64(m.memory[uint32(v5):], uint64(t268))
																					t269 := int64(load64(m.memory[int64(uint32(v4))+288:]))
																					store64(m.memory[int64(uint32(v5))+8:], uint64(t269))
																					t270 := int64(load64(m.memory[int64(uint32(v4))+296:]))
																					store64(m.memory[int64(uint32(v5))+16:], uint64(t270))
																					t271 := int32(load32(m.memory[int64(uint32(v4))+304:]))
																					store32(m.memory[int64(uint32(v5))+24:], uint32(t271))
																					store32(m.memory[int64(uint32(v4))+104:], uint32(v25+i32(1)))
																					goto l17
																				}
																			}
																			t254 := int32(load32(m.memory[int64(uint32(v4))+300:]))
																			v2 = t254
																			t255 := int32(load32(m.memory[int64(uint32(v4))+296:]))
																			v3 = t255
																			t256 := int32(load32(m.memory[int64(uint32(v4))+292:]))
																			v21 = t256
																			t257 := int32(load32(m.memory[int64(uint32(v4))+288:]))
																			v19 = t257
																			t258 := int32(load32(m.memory[int64(uint32(v1))+272:]))
																			store32(m.memory[int64(uint32(v1))+272:], uint32(t258+i32(1)))
																			{
																				t259 := int32(load32(m.memory[int64(uint32(v4))+232:]))
																				v1 = t259
																				if v1 < i32(1) {
																					goto l112
																				}
																				t260 := int32(load32(m.memory[int64(uint32(v4))+236:]))
																				m.fn18(t260, v1, i32(1))
																			}
																		l112:
																			store32(m.memory[int64(uint32(v0))+20:], uint32(v2))
																			store32(m.memory[int64(uint32(v0))+16:], uint32(v3))
																			store32(m.memory[int64(uint32(v0))+12:], uint32(v21))
																			store32(m.memory[int64(uint32(v0))+8:], uint32(v19))
																			store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
																			store32(m.memory[uint32(v0):], uint32(v25))
																			goto l113
																		}
																	}
																}
															}
														}
													l28:
														v25 = i32(2)
														if uint32(v20) < uint32(i32(2048)) {
															goto l29
														}
														p283 := i32(4)
														if uint32(v20) < uint32(i32(65536)) {
															p283 = i32(3)
														}
														v25 = p283
													}
												l29:
													{
														t284 := int32(load32(m.memory[int64(uint32(v4))+120:]))
														if uint32(v25) <= uint32(t284-v5) {
															goto l120
														}
														m.fn197(v4+i32(96)+i32(24), v5, v25, i32(1), i32(1))
													}
												l120:
													t285 := int32(load32(m.memory[int64(uint32(v4))+124:]))
													v21 = t285 + v5
													if v19 != 0 {
														goto l121
													}
													v19 = v20&i32(63) | i32(-128)
													v28 = int32(uint32(v20) >> 6)
													if uint32(v20) >= uint32(i32(2048)) {
														v23 = int32(uint32(v20) >> 12)
														v28 = v28&i32(63) | i32(-128)
														if uint32(v20) > uint32(i32(0xffff)) {
															m.memory[int64(uint32(v21))+3] = byte(v19)
															m.memory[int64(uint32(v21))+2] = byte(v28)
															m.memory[int64(uint32(v21))+1] = byte(v23&i32(63) | i32(-128))
															m.memory[uint32(v21)] = byte(int32(uint32(v20)>>18) | i32(-16))
															goto l123
														}
														m.memory[int64(uint32(v21))+2] = byte(v19)
														m.memory[int64(uint32(v21))+1] = byte(v28)
														m.memory[uint32(v21)] = byte(v23 | i32(224))
														goto l123
													}
													m.memory[int64(uint32(v21))+1] = byte(v19)
													m.memory[uint32(v21)] = byte(v28 | i32(192))
													goto l123
												l121:
													m.memory[uint32(v21)] = byte(v20)
												l123:
													store32(m.memory[int64(uint32(v4))+128:], uint32(v25+v5))
													goto l17
												}
											l59:
												t286 := int32(load32(m.memory[int64(uint32(v4))+288:]))
												v25 = t286
												t287 := int32(load32(m.memory[int64(uint32(v4))+284:]))
												v21 = t287
												t288 := int32(load32(m.memory[int64(uint32(v4))+280:]))
												v26 = t288
												m.fn560(v14)
												{
													t289 := int32(load32(m.memory[int64(uint32(v4))+304:]))
													v5 = t289
													if v5 == 0 {
														goto l125
													}
													t290 := int32(load32(m.memory[int64(uint32(v4))+308:]))
													v36 = t290
													t291 := int32(load32(m.memory[uint32(v36+i32(-4)):]))
													v34 = t291
													v51 = v34 & i32(-8)
													t292 := v51
													v34 = v34 & i32(3)
													p293 := i32(8)
													if v34 != 0 {
														p293 = i32(4)
													}
													if uint32(t292) < uint32(p293+v5) {
														m.fn3(i32(1273840), i32(46), i32(1273888))
														panic("unreachable")
													}
													if v34 == 0 {
														goto l127
													}
													if uint32(v51) > uint32(v5+i32(39)) {
														m.fn3(i32(1273904), i32(46), i32(1273952))
														panic("unreachable")
													}
												l127:
													m.fn1(v36)
												}
											l125:
												{
													var p294 int32
													if v20 == i32(7) {
														p294 = 1
													}
													v5 = p294
													if v5 != 0 {
														goto l129
													}
													if v19&i32(1) != 0 {
														goto l129
													}
													m.fn561(v4+i32(280), v4+i32(4), v4+i32(84), v4+i32(72), v4+i32(44))
													t295 := int32(load32(m.memory[int64(uint32(v4))+280:]))
													if t295 == i32(-1) {
														store32(m.memory[int64(uint32(v4))+144:], uint32(v25))
														store32(m.memory[int64(uint32(v4))+140:], uint32(v21))
														store32(m.memory[int64(uint32(v4))+136:], uint32(v26))
														{
															{
																t304 := int32(load32(m.memory[int64(uint32(v1))+68:]))
																if t304 != 0 {
																	goto l144
																}
																v5 = i32(0)
																goto l145
															}
														l144:
															t305 := int64(load64(m.memory[int64(uint32(v1))+72:]))
															t306 := int64(load64(m.memory[int64(uint32(v1))+80:]))
															t307 := m.fn106(t305, t306, v28)
															v22 = t307
															t308 := int32(load32(m.memory[int64(uint32(v1))+60:]))
															v31 = t308
															v19 = v31 & int32(v22)
															v24 = int64(uint64(v22)>>25) & i64(127) * i64(72340172838076673)
															t309 := int32(load32(m.memory[int64(uint32(v1))+56:]))
															v20 = t309
															v33 = i32(0)
														l149:
															{
																{
																	t310 := int64(load64(m.memory[uint32(v20+v19):]))
																	v27 = t310
																	v22 = v27 ^ v24
																	v22 = (v22 ^ i64(-1)) & (v22 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
																	if v22 == 0 {
																		goto l146
																	}
																l148:
																	{
																		t311 := v28 & i32(0xffff)
																		v5 = v20 + (i32(0)-(int32(uint32(int64(bits.TrailingZeros64(uint64(v22))))>>3)+v19)&v31)*i32(60)
																		t312 := int32(load16(m.memory[uint32(v5+i32(-60)):]))
																		if t311 == t312 {
																			goto l147
																		}
																		v22 = (v22 + i64(-1)) & v22
																		if !(v22 == 0) {
																			goto l148
																		}
																	}
																}
															l146:
																v5 = i32(0)
																if !(v27&(v27<<1)&i64(-0x7f7f7f7f7f7f7f80) == 0) {
																	goto l147
																}
																t313 := v19
																v33 = v33 + i32(8)
																v19 = (t313 + v33) & v31
																goto l149
															}
														l147:
															p314 := i32(0)
															if v5 != 0 {
																p314 = v5 + i32(-56)
															}
															v5 = p314
														}
													l145:
														{
															p315 := v1
															if v5 != 0 {
																p315 = v5
															}
															v28 = p315
															t316 := int32(m.memory[int64(uint32(v28))+54])
															v5 = t316
															if v5 == i32(2) {
																v19 = v25 * i32(28)
																v5 = i32(0)
																{
																l152:
																	{
																		if v19 == v5 {
																			m.fn427(v4+i32(28), v4+i32(4))
																			m.fn441(v4+i32(4), v4+i32(16))
																			if v25 == 0 {
																				goto l156
																			}
																			v5 = v21
																		l157:
																			m.fn337(v5)
																			v5 = v5 + i32(28)
																			v25 = v25 + i32(-1)
																			if v25 != 0 {
																				goto l157
																			}
																		l156:
																			if v26 == 0 {
																				goto l139
																			}
																			t322 := int32(load32(m.memory[uint32(v21+i32(-4)):]))
																			v5 = t322
																			v25 = v5 & i32(-8)
																			t323 := v25
																			v5 = v5 & i32(3)
																			p324 := i32(8)
																			if v5 != 0 {
																				p324 = i32(4)
																			}
																			v19 = v26 * i32(28)
																			if uint32(t323) < uint32(p324+v19) {
																				m.fn3(i32(1273840), i32(46), i32(1273888))
																				panic("unreachable")
																			}
																			if v5 == 0 {
																				goto l159
																			}
																			if uint32(v25) > uint32(v19+i32(39)) {
																				m.fn3(i32(1273904), i32(46), i32(1273952))
																				panic("unreachable")
																			}
																		l159:
																			m.fn1(v21)
																			goto l139
																		}
																		t317 := v21
																		v5 = v5 + i32(28)
																		t318 := m.fn311(t317 + v5 + i32(-28))
																		if t318 != 0 {
																			goto l152
																		}
																	}
																	t319 := int32(m.memory[int64(uint32(v28))+53])
																	v25 = t319
																	t320 := int32(m.memory[int64(uint32(v28))+52])
																	v5 = t320
																	if v38&i32(255) == i32(2) {
																		goto l153
																	}
																	if (v5|v38)&i32(255) == 0 {
																		goto l154
																	}
																	p321 := v42
																	if v5&i32(1) != 0 {
																		p321 = v25
																	}
																	v25 = p321
																	goto l155
																}
															l153:
																if v5&i32(1) != 0 {
																	goto l155
																}
															l154:
																if v40&i32(1) == 0 {
																	goto l161
																}
																if v39 == i32(63489) {
																	goto l161
																}
																if v39 == 0 {
																	goto l161
																}
																p325 := i32(0)
																if v44&i32(1) != 0 {
																	p325 = v43
																}
																v25 = p325
																{
																	{
																		{
																			t326 := int32(load32(m.memory[int64(uint32(v1))+148:]))
																			if t326 == 0 {
																				goto l162
																			}
																			t327 := int64(load64(m.memory[int64(uint32(v1))+152:]))
																			t328 := int64(load64(m.memory[int64(uint32(v1))+160:]))
																			t329 := m.fn106(t327, t328, v39)
																			v22 = t329
																			t330 := int32(load32(m.memory[int64(uint32(v1))+140:]))
																			v19 = t330
																			v5 = v19 & int32(v22)
																			v24 = int64(uint64(v22)>>25) & i64(127) * i64(72340172838076673)
																			t331 := int32(load32(m.memory[int64(uint32(v1))+136:]))
																			v21 = t331
																			v20 = i32(0)
																		l166:
																			{
																				{
																					t332 := int64(load64(m.memory[uint32(v21+v5):]))
																					v27 = t332
																					v22 = v27 ^ v24
																					v22 = (v22 ^ i64(-1)) & (v22 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
																					if v22 == 0 {
																						goto l163
																					}
																				l165:
																					{
																						t333 := v39
																						v28 = v21 + (i32(0)-(int32(uint32(int64(bits.TrailingZeros64(uint64(v22))))>>3)+v5)&v19)*i32(520)
																						t334 := int32(load16(m.memory[uint32(v28+i32(-520)):]))
																						if t333 == t334 {
																							goto l164
																						}
																						v22 = (v22 + i64(-1)) & v22
																						if !(v22 == 0) {
																							goto l165
																						}
																					}
																				}
																			l163:
																				if !(v27&(v27<<1)&i64(-0x7f7f7f7f7f7f7f80) == 0) {
																					goto l162
																				}
																				t335 := v5
																				v20 = v20 + i32(8)
																				v5 = (t335 + v20) & v19
																				goto l166
																			}
																		}
																	l162:
																		m.memory[int64(uint32(v4))+776] = byte(i32(0))
																		store64(m.memory[int64(uint32(v4))+768:], uint64(i64(1)))
																		m.memory[int64(uint32(v4))+764] = byte(i32(0))
																		store32(m.memory[int64(uint32(v4))+760:], uint32(i32(0)))
																		store64(m.memory[int64(uint32(v4))+752:], uint64(i64(0x400000000)))
																		store32(m.memory[int64(uint32(v4))+744:], uint32(i32(0)))
																		m.memory[int64(uint32(v4))+736] = byte(i32(0))
																		store64(m.memory[int64(uint32(v4))+728:], uint64(i64(1)))
																		m.memory[int64(uint32(v4))+724] = byte(i32(0))
																		store32(m.memory[int64(uint32(v4))+720:], uint32(i32(0)))
																		store64(m.memory[int64(uint32(v4))+712:], uint64(i64(0x400000000)))
																		store32(m.memory[int64(uint32(v4))+704:], uint32(i32(0)))
																		m.memory[int64(uint32(v4))+696] = byte(i32(0))
																		store64(m.memory[int64(uint32(v4))+688:], uint64(i64(1)))
																		m.memory[int64(uint32(v4))+684] = byte(i32(0))
																		store32(m.memory[int64(uint32(v4))+680:], uint32(i32(0)))
																		store64(m.memory[int64(uint32(v4))+672:], uint64(i64(0x400000000)))
																		store32(m.memory[int64(uint32(v4))+664:], uint32(i32(0)))
																		m.memory[int64(uint32(v4))+656] = byte(i32(0))
																		store64(m.memory[int64(uint32(v4))+648:], uint64(i64(1)))
																		m.memory[int64(uint32(v4))+644] = byte(i32(0))
																		store32(m.memory[int64(uint32(v4))+640:], uint32(i32(0)))
																		store64(m.memory[int64(uint32(v4))+632:], uint64(i64(0x400000000)))
																		store32(m.memory[int64(uint32(v4))+624:], uint32(i32(0)))
																		m.memory[int64(uint32(v4))+616] = byte(i32(0))
																		store64(m.memory[int64(uint32(v4))+608:], uint64(i64(1)))
																		m.memory[int64(uint32(v4))+604] = byte(i32(0))
																		store32(m.memory[int64(uint32(v4))+600:], uint32(i32(0)))
																		store64(m.memory[int64(uint32(v4))+592:], uint64(i64(0x400000000)))
																		store32(m.memory[int64(uint32(v4))+584:], uint32(i32(0)))
																		m.memory[int64(uint32(v4))+576] = byte(i32(0))
																		store64(m.memory[int64(uint32(v4))+568:], uint64(i64(1)))
																		m.memory[int64(uint32(v4))+564] = byte(i32(0))
																		store32(m.memory[int64(uint32(v4))+560:], uint32(i32(0)))
																		store64(m.memory[int64(uint32(v4))+552:], uint64(i64(0x400000000)))
																		store32(m.memory[int64(uint32(v4))+544:], uint32(i32(0)))
																		m.memory[int64(uint32(v4))+536] = byte(i32(0))
																		store64(m.memory[int64(uint32(v4))+528:], uint64(i64(1)))
																		m.memory[int64(uint32(v4))+524] = byte(i32(0))
																		store32(m.memory[int64(uint32(v4))+520:], uint32(i32(0)))
																		store64(m.memory[int64(uint32(v4))+512:], uint64(i64(0x400000000)))
																		store32(m.memory[int64(uint32(v4))+504:], uint32(i32(0)))
																		m.memory[int64(uint32(v4))+496] = byte(i32(0))
																		store64(m.memory[int64(uint32(v4))+488:], uint64(i64(1)))
																		m.memory[int64(uint32(v4))+484] = byte(i32(0))
																		store32(m.memory[int64(uint32(v4))+480:], uint32(i32(0)))
																		store64(m.memory[int64(uint32(v4))+472:], uint64(i64(0x400000000)))
																		store32(m.memory[int64(uint32(v4))+464:], uint32(i32(0)))
																		m.memory[int64(uint32(v4))+456] = byte(i32(0))
																		store64(m.memory[int64(uint32(v4))+448:], uint64(i64(1)))
																		m.memory[int64(uint32(v4))+444] = byte(i32(0))
																		store32(m.memory[int64(uint32(v4))+440:], uint32(i32(0)))
																		store64(m.memory[int64(uint32(v4))+432:], uint64(i64(0x400000000)))
																		store32(m.memory[int64(uint32(v4))+424:], uint32(i32(0)))
																		store64(m.memory[int64(uint32(v4))+408:], uint64(i64(0)))
																		store64(m.memory[int64(uint32(v4))+392:], uint64(i64(0)))
																		store64(m.memory[int64(uint32(v4))+376:], uint64(i64(0)))
																		store64(m.memory[int64(uint32(v4))+360:], uint64(i64(0)))
																		store64(m.memory[int64(uint32(v4))+344:], uint64(i64(0)))
																		store64(m.memory[int64(uint32(v4))+328:], uint64(i64(0)))
																		store64(m.memory[int64(uint32(v4))+312:], uint64(i64(0)))
																		store64(m.memory[int64(uint32(v4))+296:], uint64(i64(0)))
																		store64(m.memory[int64(uint32(v4))+280:], uint64(i64(0)))
																		store32(m.memory[int64(uint32(v4))+784:], uint32(v39^i32(-1)))
																		t337 := v4 + i32(280)
																		p336 := i32(8)
																		if uint32(v25) < uint32(i32(8)) {
																			p336 = v25
																		}
																		t338 := int32(m.memory[int64(uint32(t337+p336*i32(40)))+176])
																		v19 = t338
																		if v19 == i32(255) {
																			m.fn412(v13)
																			goto l161
																		}
																		v5 = v4 + i32(280)
																		v20 = i32(1)
																		goto l168
																	}
																l164:
																	v5 = v28 + i32(-512)
																	t340 := v5
																	p339 := i32(8)
																	if uint32(v25) < uint32(i32(8)) {
																		p339 = v25
																	}
																	t341 := int32(m.memory[int64(uint32(t340+p339*i32(40)))+176])
																	v19 = t341
																	if v19 == i32(255) {
																		goto l161
																	}
																	v20 = i32(0)
																}
															l168:
																if v19 != 0 {
																	t342 := int32(load32(m.memory[int64(uint32(v1))+200:]))
																	if t342 == 0 {
																		goto l171
																	}
																	m.fn355(i32(1073848))
																	panic("unreachable")
																}
																store32(m.memory[int64(uint32(v4))+240:], uint32(i32(-1)))
																v22 = i64(0)
																goto l170
															}
															m.fn441(v4+i32(4), v4+i32(16))
															m.fn563(v4+i32(28), v5&i32(1), v4+i32(136), v4+i32(4))
															goto l139
														}
													}
													t296 := int64(load64(m.memory[int64(uint32(v4))+296:]))
													store64(m.memory[int64(uint32(v0))+16:], uint64(t296))
													t297 := int64(load64(m.memory[int64(uint32(v4))+288:]))
													store64(m.memory[int64(uint32(v0))+8:], uint64(t297))
													t298 := int64(load64(m.memory[int64(uint32(v4))+280:]))
													store64(m.memory[uint32(v0):], uint64(t298))
													if v25 == 0 {
														goto l131
													}
													v1 = v21
												l132:
													m.fn337(v1)
													v1 = v1 + i32(28)
													v25 = v25 + i32(-1)
													if v25 != 0 {
														goto l132
													}
												l131:
													if v26 == 0 {
														goto l133
													}
													m.fn18(v21, v26*i32(28), i32(4))
												l133:
													switch v23 + i32(1) {
													case 0:
														goto l113
													case 1:
														goto l134
													default:
														goto l135
													}
												}
											l129:
												m.fn427(v4+i32(28), v4+i32(4))
												m.fn441(v4+i32(4), v4+i32(16))
												if v31 != i32(1) {
													goto l136
												}
												if v33 > i32(1) {
													goto l137
												}
											l136:
												if v35&i32(1) != 0 {
													goto l137
												}
												if v37&i32(255) == i32(1) {
													goto l137
												}
												if v5 != 0 {
													if v41&i32(1) != 0 {
														m.fn427(v4+i32(56), v4+i32(44))
														t350 := int32(load32(m.memory[int64(uint32(v4))+80:]))
														if t350 == 0 {
															goto l174
														}
														t351 := int64(load64(m.memory[int64(uint32(v4))+72:]))
														v22 = t351
														store64(m.memory[int64(uint32(v4))+72:], uint64(i64(0x400000000)))
														t352 := int32(load32(m.memory[int64(uint32(v4))+80:]))
														v5 = t352
														store32(m.memory[int64(uint32(v4))+80:], uint32(i32(0)))
														store32(m.memory[int64(uint32(v4))+288:], uint32(v5))
														store64(m.memory[int64(uint32(v4))+280:], uint64(v22))
														v28 = i32(-1)
														if v23 == i32(-1) {
															goto l175
														}
														{
															if v47 != 0 {
																goto l176
															}
															v52 = i32(2)
															goto l177
														l176:
															v5 = v47 << 1
															t353 := m.fn5(v5)
															v52 = t353
															if v52 == 0 {
																m.fn10(i32(2), v5)
																panic("unreachable")
															}
															if v5 == 0 {
																goto l177
															}
															memory_copy(m.memory, uint32(v52), uint32(v32), uint32(v5))
														}
													l177:
														{
															if v46 != 0 {
																goto l179
															}
															v53 = i32(1)
															v54 = i32(0)
															v55 = v47
															goto l180
														l179:
															v5 = v46 << 2
															t354 := m.fn5(v5)
															v53 = t354
															if v53 == 0 {
																m.fn10(i32(1), v5)
																panic("unreachable")
															}
															if v5 == 0 {
																goto l182
															}
															memory_copy(m.memory, uint32(v53), uint32(v30), uint32(v5))
														l182:
															v55 = v47
															v54 = v46
														}
													l180:
														v56 = v45
														v28 = v47
													l175:
														t355 := int32(load32(m.memory[int64(uint32(v4))+288:]))
														store32(m.memory[int64(uint32(v4))+240:], uint32(t355))
														t356 := int64(load64(m.memory[int64(uint32(v4))+280:]))
														store64(m.memory[int64(uint32(v4))+232:], uint64(t356))
														{
															t357 := int32(load32(m.memory[int64(uint32(v4))+92:]))
															v19 = t357
															t358 := int32(load32(m.memory[int64(uint32(v4))+84:]))
															if v19 != t358 {
																goto l183
															}
															m.fn321(v4 + i32(84))
														}
													l183:
														t359 := int32(load32(m.memory[int64(uint32(v4))+88:]))
														v5 = t359 + v19*i32(40)
														t360 := int64(load64(m.memory[int64(uint32(v4))+232:]))
														store64(m.memory[uint32(v5):], uint64(t360))
														t361 := int32(load32(m.memory[int64(uint32(v4))+240:]))
														store32(m.memory[int64(uint32(v5))+8:], uint32(t361))
														m.memory[int64(uint32(v5))+36] = byte(v56)
														store32(m.memory[int64(uint32(v5))+32:], uint32(v54))
														store32(m.memory[int64(uint32(v5))+28:], uint32(v53))
														store32(m.memory[int64(uint32(v5))+24:], uint32(v54))
														store32(m.memory[int64(uint32(v5))+20:], uint32(v55))
														store32(m.memory[int64(uint32(v5))+16:], uint32(v52))
														store32(m.memory[int64(uint32(v5))+12:], uint32(v28))
														store32(m.memory[int64(uint32(v4))+92:], uint32(v19+i32(1)))
														goto l174
													}
													store32(m.memory[int64(uint32(v4))+288:], uint32(v25))
													store32(m.memory[int64(uint32(v4))+284:], uint32(v21))
													store32(m.memory[int64(uint32(v4))+280:], uint32(v26))
													m.fn562(v1, v28, v4+i32(280), v4+i32(44), v4+i32(56))
													m.fn427(v4+i32(56), v4+i32(44))
													t343 := int64(load64(m.memory[int64(uint32(v4))+44:]))
													v22 = t343
													store64(m.memory[int64(uint32(v4))+44:], uint64(i64(0x800000000)))
													t344 := int32(load32(m.memory[int64(uint32(v4))+52:]))
													v5 = t344
													store32(m.memory[int64(uint32(v4))+52:], uint32(i32(0)))
													store32(m.memory[int64(uint32(v4))+288:], uint32(v5))
													store64(m.memory[int64(uint32(v4))+280:], uint64(v22))
													{
														t345 := int32(load32(m.memory[int64(uint32(v4))+80:]))
														v5 = t345
														t346 := int32(load32(m.memory[int64(uint32(v4))+72:]))
														if v5 != t346 {
															goto l173
														}
														m.fn316(v4 + i32(72))
													}
												l173:
													t347 := int32(load32(m.memory[int64(uint32(v4))+76:]))
													v25 = t347 + v5*i32(12)
													t348 := int64(load64(m.memory[int64(uint32(v4))+280:]))
													store64(m.memory[uint32(v25):], uint64(t348))
													t349 := int32(load32(m.memory[int64(uint32(v4))+288:]))
													store32(m.memory[int64(uint32(v25))+8:], uint32(t349))
													store32(m.memory[int64(uint32(v4))+80:], uint32(v5+i32(1)))
													goto l139
												}
												store32(m.memory[int64(uint32(v4))+288:], uint32(v25))
												store32(m.memory[int64(uint32(v4))+284:], uint32(v21))
												store32(m.memory[int64(uint32(v4))+280:], uint32(v26))
												m.fn562(v1, v28, v4+i32(280), v4+i32(44), v4+i32(56))
												goto l139
											l135:
												m.fn18(v32, v23<<1, i32(2))
											l134:
												if v29 == 0 {
													goto l113
												}
												m.fn18(v30, v29<<2, i32(1))
											}
										l113:
											t299 := int32(load32(m.memory[int64(uint32(v4))+100:]))
											v5 = t299
											{
												t300 := int32(load32(m.memory[int64(uint32(v4))+104:]))
												v2 = t300
												if v2 == 0 {
													goto l140
												}
												v1 = v5
											l141:
												m.fn337(v1)
												v1 = v1 + i32(28)
												v2 = v2 + i32(-1)
												if v2 != 0 {
													goto l141
												}
											}
										l140:
											{
												t301 := int32(load32(m.memory[int64(uint32(v4))+96:]))
												v1 = t301
												if v1 == 0 {
													goto l142
												}
												m.fn18(v5, v1*i32(28), i32(4))
											}
										l142:
											m.fn560(v6)
											t302 := int32(load32(m.memory[int64(uint32(v4))+120:]))
											v1 = t302
											if v1 == 0 {
												goto l143
											}
											t303 := int32(load32(m.memory[int64(uint32(v4))+124:]))
											m.fn18(t303, v1, i32(1))
											goto l143
										}
									l137:
										store32(m.memory[int64(uint32(v4))+288:], uint32(v25))
										store32(m.memory[int64(uint32(v4))+284:], uint32(v21))
										store32(m.memory[int64(uint32(v4))+280:], uint32(v26))
										m.fn562(v1, v28, v4+i32(280), v4+i32(44), v4+i32(56))
										goto l139
									l174:
										if v25 == 0 {
											goto l184
										}
										v5 = v21
									l185:
										m.fn337(v5)
										v5 = v5 + i32(28)
										v25 = v25 + i32(-1)
										if v25 != 0 {
											goto l185
										}
									l184:
										if v26 == 0 {
											goto l139
										}
										{
											t362 := int32(load32(m.memory[uint32(v21+i32(-4)):]))
											v5 = t362
											v25 = v5 & i32(-8)
											t363 := v25
											v5 = v5 & i32(3)
											p364 := i32(8)
											if v5 != 0 {
												p364 = i32(4)
											}
											v19 = v26 * i32(28)
											if uint32(t363) < uint32(p364+v19) {
												m.fn3(i32(1273840), i32(46), i32(1273888))
												panic("unreachable")
											}
											if v5 == 0 {
												goto l187
											}
											if uint32(v25) > uint32(v19+i32(39)) {
												m.fn3(i32(1273904), i32(46), i32(1273952))
												panic("unreachable")
											}
										l187:
											m.fn1(v21)
											goto l139
										}
									l171:
										store32(m.memory[int64(uint32(v1))+200:], uint32(i32(-1)))
										m.fn564(v4+i32(232), v7, v39, v5, v25)
										t365 := int32(load32(m.memory[int64(uint32(v1))+200:]))
										store32(m.memory[int64(uint32(v1))+200:], uint32(t365+i32(1)))
										t366 := int64(load64(m.memory[int64(uint32(v4))+232:]))
										v22 = t366
									}
								l170:
									t367 := int32(load32(m.memory[int64(uint32(v12))+8:]))
									store32(m.memory[int64(uint32(v4))+192:], uint32(t367))
									t368 := int64(load64(m.memory[uint32(v12):]))
									store64(m.memory[int64(uint32(v4))+184:], uint64(t368))
									m.fn427(v4+i32(28), v4+i32(4))
									t369 := int64(load32(m.memory[int64(uint32(v5))+504:]))
									v24 = t369
									t370 := m.fn190(i32(8), i32(32))
									v21 = t370
									store32(m.memory[uint32(v21):], uint32(i32(-0x80000000)))
									t371 := int64(load64(m.memory[int64(uint32(v4))+136:]))
									store64(m.memory[int64(uint32(v21))+4:], uint64(t371))
									t372 := int32(load32(m.memory[int64(uint32(v4))+144:]))
									store32(m.memory[int64(uint32(v21))+12:], uint32(t372))
									{
										t373 := int32(load32(m.memory[int64(uint32(v4))+24:]))
										v28 = t373
										t374 := int32(load32(m.memory[int64(uint32(v4))+16:]))
										if v28 != t374 {
											goto l189
										}
										m.fn324(v4 + i32(16))
									}
								l189:
									t375 := int32(load32(m.memory[int64(uint32(v4))+20:]))
									v5 = t375 + v28*i32(56)
									store32(m.memory[int64(uint32(v5))+24:], uint32(v25))
									store64(m.memory[int64(uint32(v5))+16:], uint64(v22))
									m.memory[int64(uint32(v5))+8] = byte(v19)
									store64(m.memory[uint32(v5):], uint64(v24))
									t376 := int64(load64(m.memory[int64(uint32(v4))+184:]))
									store64(m.memory[int64(uint32(v5))+28:], uint64(t376))
									t377 := int32(load32(m.memory[int64(uint32(v4))+192:]))
									store32(m.memory[int64(uint32(v5))+36:], uint32(t377))
									store32(m.memory[int64(uint32(v5))+48:], uint32(i32(1)))
									store32(m.memory[int64(uint32(v5))+44:], uint32(v21))
									store32(m.memory[int64(uint32(v5))+40:], uint32(i32(1)))
									store32(m.memory[int64(uint32(v4))+24:], uint32(v28+i32(1)))
									if v20 == 0 {
										goto l139
									}
									m.fn412(v13)
									goto l139
								}
							l161:
								m.fn427(v4+i32(28), v4+i32(4))
								m.fn441(v4+i32(4), v4+i32(16))
								{
									t378 := int32(load32(m.memory[int64(uint32(v4))+12:]))
									v5 = t378
									t379 := int32(load32(m.memory[int64(uint32(v4))+4:]))
									if v5 != t379 {
										goto l190
									}
									m.fn315(v4 + i32(4))
								}
							l190:
								t380 := int32(load32(m.memory[int64(uint32(v4))+8:]))
								v25 = t380 + v5<<5
								store32(m.memory[uint32(v25):], uint32(i32(-0x80000000)))
								t381 := int64(load64(m.memory[int64(uint32(v4))+136:]))
								store64(m.memory[int64(uint32(v25))+4:], uint64(t381))
								t382 := int32(load32(m.memory[int64(uint32(v4))+144:]))
								store32(m.memory[int64(uint32(v25))+12:], uint32(t382))
								store32(m.memory[int64(uint32(v4))+12:], uint32(v5+i32(1)))
								goto l139
							}
						l155:
							m.fn427(v4+i32(28), v4+i32(4))
							m.fn441(v4+i32(4), v4+i32(16))
							t383 := int64(load64(m.memory[int64(uint32(v4))+136:]))
							store64(m.memory[int64(uint32(v4))+184:], uint64(t383))
							t384 := int32(load32(m.memory[int64(uint32(v4))+144:]))
							t385 := v4
							v5 = t384
							store32(m.memory[int64(uint32(t385))+192:], uint32(v5))
							t386 := int32(load32(m.memory[int64(uint32(v4))+188:]))
							v21 = t386
							t387 := int32(load32(m.memory[int64(uint32(v28))+48:]))
							m.fn459(v21, v5, t387)
							{
								if v40&i32(1) == 0 {
									goto l191
								}
								if v39 == 0 {
									goto l191
								}
								if v39 == i32(63489) {
									goto l191
								}
								t388 := int32(load32(m.memory[int64(uint32(v1))+148:]))
								if t388 == 0 {
									goto l191
								}
								p389 := i32(0)
								if v44&i32(1) != 0 {
									p389 = v43
								}
								v20 = p389
								t390 := int64(load64(m.memory[int64(uint32(v1))+152:]))
								t391 := int64(load64(m.memory[int64(uint32(v1))+160:]))
								t392 := m.fn106(t390, t391, v39)
								v22 = t392
								t393 := int32(load32(m.memory[int64(uint32(v1))+140:]))
								v26 = t393
								v19 = v26 & int32(v22)
								v24 = int64(uint64(v22)>>25) & i64(127) * i64(72340172838076673)
								t394 := int32(load32(m.memory[int64(uint32(v1))+136:]))
								v28 = t394
								v33 = i32(0)
							l201:
								{
									{
										t395 := int64(load64(m.memory[uint32(v28+v19):]))
										v27 = t395
										v22 = v27 ^ v24
										v22 = (v22 ^ i64(-1)) & (v22 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
										if v22 == 0 {
											goto l192
										}
									l194:
										{
											t396 := v39
											v31 = v28 + (i32(0)-(int32(uint32(int64(bits.TrailingZeros64(uint64(v22))))>>3)+v19)&v26)*i32(520)
											t397 := int32(load16(m.memory[uint32(v31+i32(-520)):]))
											if t396 == t397 {
												v19 = v31 + i32(-512)
												t399 := v19
												p398 := i32(8)
												if uint32(v20) < uint32(i32(8)) {
													p398 = v20
												}
												t400 := int32(m.memory[int64(uint32(t399+p398*i32(40)))+176])
												v28 = t400
												if uint32((v28+i32(-1))&i32(255)) > uint32(i32(253)) {
													goto l191
												}
												{
													t401 := int32(load32(m.memory[int64(uint32(v1))+200:]))
													if t401 != 0 {
														m.fn355(i32(1073832))
														panic("unreachable")
													}
													store32(m.memory[int64(uint32(v1))+200:], uint32(i32(-1)))
													m.fn564(v4+i32(280), v7, v39, v19, v20)
													t402 := int32(load32(m.memory[int64(uint32(v1))+200:]))
													store32(m.memory[int64(uint32(v1))+200:], uint32(t402+i32(1)))
													{
														{
															t403 := int32(load32(m.memory[int64(uint32(v4))+288:]))
															v19 = t403
															if v19 == i32(-1) {
																goto l196
															}
															t404 := int64(load64(m.memory[int64(uint32(v4))+292:]))
															store64(m.memory[int64(uint32(v4))+284:], uint64(t404))
															store32(m.memory[int64(uint32(v4))+280:], uint32(v19))
															goto l197
														}
													l196:
														t405 := int64(load64(m.memory[int64(uint32(v4))+280:]))
														m.fn307(v4+i32(280), v28, t405)
													}
												l197:
													store64(m.memory[int64(uint32(v4))+792:], uint64(v9))
													m.fn12(v4+i32(232), i32(1067493), v4+i32(792))
													{
														t406 := int32(load32(m.memory[int64(uint32(v4))+280:]))
														v19 = t406
														if v19 == 0 {
															goto l198
														}
														t407 := int32(load32(m.memory[int64(uint32(v4))+284:]))
														m.fn18(t407, v19, i32(1))
													}
												l198:
													t408 := int32(load32(m.memory[int64(uint32(v4))+232:]))
													v19 = t408
													if v19 == i32(-1) {
														goto l191
													}
													t409 := int64(load64(m.memory[int64(uint32(v4))+236:]))
													v22 = t409
													{
														t410 := int32(load32(m.memory[int64(uint32(v4))+184:]))
														if v5 != t410 {
															goto l199
														}
														m.fn318(v4 + i32(184))
														t411 := int32(load32(m.memory[int64(uint32(v4))+188:]))
														v21 = t411
													}
												l199:
													if v5 == 0 {
														goto l200
													}
													v28 = v5 * i32(28)
													if v28 == 0 {
														goto l200
													}
													memory_copy(m.memory, uint32(v21+i32(28)), uint32(v21), uint32(v28))
												l200:
													store32(m.memory[int64(uint32(v21))+16:], uint32(i32(0)))
													store64(m.memory[int64(uint32(v21))+8:], uint64(v22))
													store32(m.memory[int64(uint32(v21))+4:], uint32(v19))
													store32(m.memory[uint32(v21):], uint32(i32(3)))
													store32(m.memory[int64(uint32(v4))+192:], uint32(v5+i32(1)))
													goto l191
												}
											}
											v22 = (v22 + i64(-1)) & v22
											if v22 == 0 {
												goto l192
											}
											goto l194
										}
									}
								l192:
									if !(v27&(v27<<1)&i64(-0x7f7f7f7f7f7f7f80) == 0) {
										goto l191
									}
									t412 := v19
									v33 = v33 + i32(8)
									v19 = (t412 + v33) & v26
									goto l201
								}
							}
						l191:
							{
								t413 := int32(load32(m.memory[int64(uint32(v4))+12:]))
								v21 = t413
								t414 := int32(load32(m.memory[int64(uint32(v4))+4:]))
								if v21 != t414 {
									goto l202
								}
								m.fn315(v4 + i32(4))
							}
						l202:
							t415 := int32(load32(m.memory[int64(uint32(v4))+8:]))
							v5 = t415 + v21<<5
							t416 := int64(load64(m.memory[int64(uint32(v4))+184:]))
							store64(m.memory[uint32(v5):], uint64(t416))
							t417 := int32(load32(m.memory[int64(uint32(v4))+192:]))
							store32(m.memory[int64(uint32(v5))+8:], uint32(t417))
							m.memory[int64(uint32(v5))+24] = byte(v25)
							store32(m.memory[int64(uint32(v5))+12:], uint32(i32(-1)))
							store32(m.memory[int64(uint32(v4))+12:], uint32(v21+i32(1)))
						}
					l139:
						switch v23 + i32(1) {
						case 0:
							goto l17
						default:
							t418 := int32(load32(m.memory[uint32(v32+i32(-4)):]))
							v5 = t418
							v25 = v5 & i32(-8)
							t419 := v25
							v5 = v5 & i32(3)
							p420 := i32(8)
							if v5 != 0 {
								p420 = i32(4)
							}
							v21 = v23 << 1
							if uint32(t419) < uint32(p420+v21) {
								m.fn3(i32(1273840), i32(46), i32(1273888))
								panic("unreachable")
							}
							if v5 == 0 {
								goto l206
							}
							if uint32(v25) > uint32(v21+i32(39)) {
								m.fn3(i32(1273904), i32(46), i32(1273952))
								panic("unreachable")
							}
						l206:
							m.fn1(v32)
							fallthrough
						case 1:
							if v29 == 0 {
								goto l17
							}
							t421 := int32(load32(m.memory[uint32(v30+i32(-4)):]))
							v5 = t421
							v25 = v5 & i32(-8)
							t422 := v25
							v5 = v5 & i32(3)
							p423 := i32(8)
							if v5 != 0 {
								p423 = i32(4)
							}
							v21 = v29 << 2
							if uint32(t422) < uint32(p423+v21) {
								m.fn3(i32(1273840), i32(46), i32(1273888))
								panic("unreachable")
							}
							if v5 == 0 {
								goto l209
							}
							if uint32(v25) > uint32(v21+i32(39)) {
								m.fn3(i32(1273904), i32(46), i32(1273952))
								panic("unreachable")
							}
						l209:
							m.fn1(v30)
						}
					l17:
						v2 = v2 + i32(1)
						t424 := int32(load32(m.memory[int64(uint32(v1))+96:]))
						t425 := v2
						v5 = t424
						p426 := v3
						if uint32(v5) < uint32(v3) {
							p426 = v5
						}
						if uint32(t425) < uint32(p426) {
							goto l15
						}
						goto l0
					}
				}
			l0:
				t427 := int64(load64(m.memory[int64(uint32(v4))+128:]))
				store64(m.memory[int64(uint32(v4))+312:], uint64(t427))
				t428 := int64(load64(m.memory[int64(uint32(v4))+120:]))
				store64(m.memory[int64(uint32(v4))+304:], uint64(t428))
				t429 := int64(load64(m.memory[int64(uint32(v4))+112:]))
				store64(m.memory[int64(uint32(v4))+296:], uint64(t429))
				t430 := int64(load64(m.memory[int64(uint32(v4))+104:]))
				store64(m.memory[int64(uint32(v4))+288:], uint64(t430))
				t431 := int64(load64(m.memory[int64(uint32(v4))+96:]))
				store64(m.memory[int64(uint32(v4))+280:], uint64(t431))
				m.fn557(v4 + i32(280))
				{
					t432 := int32(load32(m.memory[int64(uint32(v4))+300:]))
					if t432 == 0 {
						goto l211
					}
				l212:
					{
						m.fn559(v4 + i32(280))
						t433 := int32(load32(m.memory[int64(uint32(v4))+300:]))
						if t433 != 0 {
							goto l212
						}
					}
				}
			l211:
				t434 := int32(load32(m.memory[int64(uint32(v4))+288:]))
				v2 = t434
				t435 := int32(load32(m.memory[int64(uint32(v4))+284:]))
				v5 = t435
				t436 := int32(load32(m.memory[int64(uint32(v4))+280:]))
				v3 = t436
				m.fn560(v4 + i32(292))
				{
					t437 := int32(load32(m.memory[int64(uint32(v4))+304:]))
					v1 = t437
					if v1 == 0 {
						goto l213
					}
					t438 := int32(load32(m.memory[int64(uint32(v4))+308:]))
					v21 = t438
					t439 := int32(load32(m.memory[uint32(v21+i32(-4)):]))
					v25 = t439
					v19 = v25 & i32(-8)
					t440 := v19
					v25 = v25 & i32(3)
					p441 := i32(8)
					if v25 != 0 {
						p441 = i32(4)
					}
					if uint32(t440) < uint32(p441+v1) {
						m.fn3(i32(1273840), i32(46), i32(1273888))
						panic("unreachable")
					}
					if v25 == 0 {
						goto l215
					}
					if uint32(v19) > uint32(v1+i32(39)) {
						m.fn3(i32(1273904), i32(46), i32(1273952))
						panic("unreachable")
					}
				l215:
					m.fn1(v21)
				}
			l213:
				m.fn427(v4+i32(56), v4+i32(44))
				m.fn561(v4+i32(280), v4+i32(4), v4+i32(84), v4+i32(72), v4+i32(44))
				t442 := int32(load32(m.memory[int64(uint32(v4))+280:]))
				if t442 == i32(-1) {
					goto l217
				}
				t443 := int64(load64(m.memory[int64(uint32(v4))+296:]))
				store64(m.memory[int64(uint32(v0))+16:], uint64(t443))
				t444 := int64(load64(m.memory[int64(uint32(v4))+288:]))
				store64(m.memory[int64(uint32(v0))+8:], uint64(t444))
				t445 := int64(load64(m.memory[int64(uint32(v4))+280:]))
				store64(m.memory[uint32(v0):], uint64(t445))
				if v2 == 0 {
					goto l218
				}
				v1 = v5
			l219:
				m.fn337(v1)
				v1 = v1 + i32(28)
				v2 = v2 + i32(-1)
				if v2 != 0 {
					goto l219
				}
			l218:
				if v3 == 0 {
					goto l143
				}
				t446 := int32(load32(m.memory[uint32(v5+i32(-4)):]))
				v1 = t446
				v2 = v1 & i32(-8)
				t447 := v2
				v1 = v1 & i32(3)
				p448 := i32(8)
				if v1 != 0 {
					p448 = i32(4)
				}
				v25 = v3 * i32(28)
				if uint32(t447) < uint32(p448+v25) {
					m.fn3(i32(1273840), i32(46), i32(1273888))
					panic("unreachable")
				}
				if v1 == 0 {
					goto l221
				}
				if uint32(v2) > uint32(v25+i32(39)) {
					m.fn3(i32(1273904), i32(46), i32(1273952))
					panic("unreachable")
				}
			l221:
				m.fn1(v5)
			}
		l143:
			m.fn565(v4 + i32(84))
			t449 := int32(load32(m.memory[int64(uint32(v4))+76:]))
			v21 = t449
			{
				t450 := int32(load32(m.memory[int64(uint32(v4))+80:]))
				v19 = t450
				if v19 == 0 {
					goto l223
				}
				v5 = i32(0)
			l230:
				{
					v25 = v21 + v5*i32(12)
					t451 := int32(load32(m.memory[int64(uint32(v25))+4:]))
					v3 = t451
					{
						t452 := int32(load32(m.memory[int64(uint32(v25))+8:]))
						v2 = t452
						if v2 == 0 {
							goto l224
						}
						v1 = v3
					l225:
						m.fn335(v1)
						v1 = v1 + i32(32)
						v2 = v2 + i32(-1)
						if v2 != 0 {
							goto l225
						}
					}
				l224:
					{
						t453 := int32(load32(m.memory[uint32(v25):]))
						v1 = t453
						if v1 == 0 {
							goto l226
						}
						t454 := int32(load32(m.memory[uint32(v3+i32(-4)):]))
						v2 = t454
						v25 = v2 & i32(-8)
						t455 := v25
						v2 = v2 & i32(3)
						p456 := i32(8)
						if v2 != 0 {
							p456 = i32(4)
						}
						v1 = v1 << 5
						if uint32(t455) < uint32(p456|v1) {
							m.fn3(i32(1273840), i32(46), i32(1273888))
							panic("unreachable")
						}
						if v2 == 0 {
							goto l228
						}
						if uint32(v25) > uint32(v1+i32(39)) {
							m.fn3(i32(1273904), i32(46), i32(1273952))
							panic("unreachable")
						}
					l228:
						m.fn1(v3)
					}
				l226:
					v5 = v5 + i32(1)
					if v5 != v19 {
						goto l230
					}
				}
			}
		l223:
			{
				t457 := int32(load32(m.memory[int64(uint32(v4))+72:]))
				v1 = t457
				if v1 == 0 {
					goto l231
				}
				t458 := int32(load32(m.memory[uint32(v21+i32(-4)):]))
				v2 = t458
				v5 = v2 & i32(-8)
				t459 := v5
				v2 = v2 & i32(3)
				p460 := i32(8)
				if v2 != 0 {
					p460 = i32(4)
				}
				v1 = v1 * i32(12)
				if uint32(t459) < uint32(p460+v1) {
					m.fn3(i32(1273840), i32(46), i32(1273888))
					panic("unreachable")
				}
				if v2 == 0 {
					goto l233
				}
				if uint32(v5) > uint32(v1+i32(39)) {
					m.fn3(i32(1273904), i32(46), i32(1273952))
					panic("unreachable")
				}
			l233:
				m.fn1(v21)
			}
		l231:
			m.fn428(v4 + i32(56))
			t461 := int32(load32(m.memory[int64(uint32(v4))+48:]))
			v5 = t461
			{
				t462 := int32(load32(m.memory[int64(uint32(v4))+52:]))
				v2 = t462
				if v2 == 0 {
					goto l235
				}
				v1 = v5
			l236:
				m.fn335(v1)
				v1 = v1 + i32(32)
				v2 = v2 + i32(-1)
				if v2 != 0 {
					goto l236
				}
			}
		l235:
			{
				t463 := int32(load32(m.memory[int64(uint32(v4))+44:]))
				v1 = t463
				if v1 == 0 {
					goto l237
				}
				t464 := int32(load32(m.memory[uint32(v5+i32(-4)):]))
				v2 = t464
				v25 = v2 & i32(-8)
				t465 := v25
				v2 = v2 & i32(3)
				p466 := i32(8)
				if v2 != 0 {
					p466 = i32(4)
				}
				v1 = v1 << 5
				if uint32(t465) < uint32(p466|v1) {
					m.fn3(i32(1273840), i32(46), i32(1273888))
					panic("unreachable")
				}
				if v2 == 0 {
					goto l239
				}
				if uint32(v25) > uint32(v1+i32(39)) {
					m.fn3(i32(1273904), i32(46), i32(1273952))
					panic("unreachable")
				}
			l239:
				m.fn1(v5)
			}
		l237:
			m.fn428(v4 + i32(28))
			m.fn443(v4 + i32(16))
			t467 := int32(load32(m.memory[int64(uint32(v4))+8:]))
			v5 = t467
			{
				t468 := int32(load32(m.memory[int64(uint32(v4))+12:]))
				v2 = t468
				if v2 == 0 {
					goto l241
				}
				v1 = v5
			l242:
				m.fn335(v1)
				v1 = v1 + i32(32)
				v2 = v2 + i32(-1)
				if v2 != 0 {
					goto l242
				}
			}
		l241:
			t469 := int32(load32(m.memory[int64(uint32(v4))+4:]))
			v1 = t469
			if v1 == 0 {
				goto l243
			}
			t470 := int32(load32(m.memory[uint32(v5+i32(-4)):]))
			v2 = t470
			v25 = v2 & i32(-8)
			t471 := v25
			v2 = v2 & i32(3)
			p472 := i32(8)
			if v2 != 0 {
				p472 = i32(4)
			}
			v1 = v1 << 5
			if uint32(t471) < uint32(p472|v1) {
				m.fn3(i32(1273840), i32(46), i32(1273888))
				panic("unreachable")
			}
			if v2 == 0 {
				goto l245
			}
			if uint32(v25) > uint32(v1+i32(39)) {
				m.fn3(i32(1273904), i32(46), i32(1273952))
				panic("unreachable")
			}
		l245:
			m.fn1(v5)
			goto l243
		}
	l217:
		v25 = v2 * i32(28)
		v1 = i32(0)
		{
			{
			l248:
				{
					if v25 == v1 {
						goto l247
					}
					t473 := v5
					v1 = v1 + i32(28)
					t474 := m.fn311(t473 + v1 + i32(-28))
					if t474 != 0 {
						goto l248
					}
				}
				m.fn427(v4+i32(28), v4+i32(4))
				m.fn441(v4+i32(4), v4+i32(16))
				{
					t475 := int32(load32(m.memory[int64(uint32(v4))+12:]))
					v25 = t475
					t476 := int32(load32(m.memory[int64(uint32(v4))+4:]))
					if v25 != t476 {
						goto l249
					}
					m.fn315(v4 + i32(4))
				}
			l249:
				t477 := int32(load32(m.memory[int64(uint32(v4))+8:]))
				v1 = t477 + v25<<5
				store32(m.memory[int64(uint32(v1))+12:], uint32(v2))
				store32(m.memory[int64(uint32(v1))+8:], uint32(v5))
				store32(m.memory[int64(uint32(v1))+4:], uint32(v3))
				store32(m.memory[uint32(v1):], uint32(i32(-0x80000000)))
				store32(m.memory[int64(uint32(v4))+12:], uint32(v25+i32(1)))
				m.fn427(v4+i32(28), v4+i32(4))
				m.fn441(v4+i32(4), v4+i32(16))
				store32(m.memory[uint32(v0):], uint32(i32(-1)))
				t478 := int32(load32(m.memory[int64(uint32(v4))+12:]))
				store32(m.memory[int64(uint32(v0))+12:], uint32(t478))
				t479 := int64(load64(m.memory[int64(uint32(v4))+4:]))
				store64(m.memory[int64(uint32(v0))+4:], uint64(t479))
				goto l250
			}
		l247:
			m.fn427(v4+i32(28), v4+i32(4))
			m.fn441(v4+i32(4), v4+i32(16))
			store32(m.memory[uint32(v0):], uint32(i32(-1)))
			t480 := int32(load32(m.memory[int64(uint32(v4))+12:]))
			store32(m.memory[int64(uint32(v0))+12:], uint32(t480))
			t481 := int64(load64(m.memory[int64(uint32(v4))+4:]))
			store64(m.memory[int64(uint32(v0))+4:], uint64(t481))
			if v2 == 0 {
				goto l251
			}
			v1 = v5
		l252:
			m.fn337(v1)
			v1 = v1 + i32(28)
			v2 = v2 + i32(-1)
			if v2 != 0 {
				goto l252
			}
		l251:
			if v3 == 0 {
				goto l250
			}
			t482 := int32(load32(m.memory[uint32(v5+i32(-4)):]))
			v1 = t482
			v2 = v1 & i32(-8)
			t483 := v2
			v1 = v1 & i32(3)
			p484 := i32(8)
			if v1 != 0 {
				p484 = i32(4)
			}
			v25 = v3 * i32(28)
			if uint32(t483) < uint32(p484+v25) {
				m.fn3(i32(1273840), i32(46), i32(1273888))
				panic("unreachable")
			}
			if v1 == 0 {
				goto l254
			}
			if uint32(v2) > uint32(v25+i32(39)) {
				m.fn3(i32(1273904), i32(46), i32(1273952))
				panic("unreachable")
			}
		l254:
			m.fn1(v5)
		}
	l250:
		m.fn565(v4 + i32(84))
		t485 := int32(load32(m.memory[int64(uint32(v4))+76:]))
		v21 = t485
		{
			t486 := int32(load32(m.memory[int64(uint32(v4))+80:]))
			v19 = t486
			if v19 == 0 {
				goto l256
			}
			v5 = i32(0)
		l263:
			{
				v25 = v21 + v5*i32(12)
				t487 := int32(load32(m.memory[int64(uint32(v25))+4:]))
				v3 = t487
				{
					t488 := int32(load32(m.memory[int64(uint32(v25))+8:]))
					v2 = t488
					if v2 == 0 {
						goto l257
					}
					v1 = v3
				l258:
					m.fn335(v1)
					v1 = v1 + i32(32)
					v2 = v2 + i32(-1)
					if v2 != 0 {
						goto l258
					}
				}
			l257:
				{
					t489 := int32(load32(m.memory[uint32(v25):]))
					v1 = t489
					if v1 == 0 {
						goto l259
					}
					t490 := int32(load32(m.memory[uint32(v3+i32(-4)):]))
					v2 = t490
					v25 = v2 & i32(-8)
					t491 := v25
					v2 = v2 & i32(3)
					p492 := i32(8)
					if v2 != 0 {
						p492 = i32(4)
					}
					v1 = v1 << 5
					if uint32(t491) < uint32(p492|v1) {
						m.fn3(i32(1273840), i32(46), i32(1273888))
						panic("unreachable")
					}
					if v2 == 0 {
						goto l261
					}
					if uint32(v25) > uint32(v1+i32(39)) {
						m.fn3(i32(1273904), i32(46), i32(1273952))
						panic("unreachable")
					}
				l261:
					m.fn1(v3)
				}
			l259:
				v5 = v5 + i32(1)
				if v5 != v19 {
					goto l263
				}
			}
		}
	l256:
		{
			t493 := int32(load32(m.memory[int64(uint32(v4))+72:]))
			v1 = t493
			if v1 == 0 {
				goto l264
			}
			t494 := int32(load32(m.memory[uint32(v21+i32(-4)):]))
			v2 = t494
			v5 = v2 & i32(-8)
			t495 := v5
			v2 = v2 & i32(3)
			p496 := i32(8)
			if v2 != 0 {
				p496 = i32(4)
			}
			v1 = v1 * i32(12)
			if uint32(t495) < uint32(p496+v1) {
				m.fn3(i32(1273840), i32(46), i32(1273888))
				panic("unreachable")
			}
			if v2 == 0 {
				goto l266
			}
			if uint32(v5) > uint32(v1+i32(39)) {
				m.fn3(i32(1273904), i32(46), i32(1273952))
				panic("unreachable")
			}
		l266:
			m.fn1(v21)
		}
	l264:
		m.fn428(v4 + i32(56))
		t497 := int32(load32(m.memory[int64(uint32(v4))+48:]))
		v5 = t497
		{
			t498 := int32(load32(m.memory[int64(uint32(v4))+52:]))
			v2 = t498
			if v2 == 0 {
				goto l268
			}
			v1 = v5
		l269:
			m.fn335(v1)
			v1 = v1 + i32(32)
			v2 = v2 + i32(-1)
			if v2 != 0 {
				goto l269
			}
		}
	l268:
		{
			t499 := int32(load32(m.memory[int64(uint32(v4))+44:]))
			v1 = t499
			if v1 == 0 {
				goto l270
			}
			t500 := int32(load32(m.memory[uint32(v5+i32(-4)):]))
			v2 = t500
			v25 = v2 & i32(-8)
			t501 := v25
			v2 = v2 & i32(3)
			p502 := i32(8)
			if v2 != 0 {
				p502 = i32(4)
			}
			v1 = v1 << 5
			if uint32(t501) < uint32(p502|v1) {
				m.fn3(i32(1273840), i32(46), i32(1273888))
				panic("unreachable")
			}
			if v2 == 0 {
				goto l272
			}
			if uint32(v25) > uint32(v1+i32(39)) {
				m.fn3(i32(1273904), i32(46), i32(1273952))
				panic("unreachable")
			}
		l272:
			m.fn1(v5)
		}
	l270:
		m.fn428(v4 + i32(28))
		m.fn443(v4 + i32(16))
	}
l243:
	m.g0 = v4 + i32(800)
}
func (m *Module) fn555(v0 int32) {
	var v1, v2, v3, v4, v5 int32
	var v6 int64
	var v7, v8, v9 int32
	{
		{
			t0 := int32(load32(m.memory[int64(uint32(v0))+88:]))
			v1 = t0
			if v1 == 0 {
				goto l0
			}
			t1 := int32(load32(m.memory[int64(uint32(v0))+92:]))
			v2 = t1
			t2 := int32(load32(m.memory[uint32(v2+i32(-4)):]))
			v3 = t2
			v4 = v3 & i32(-8)
			t3 := v4
			v3 = v3 & i32(3)
			p4 := i32(8)
			if v3 != 0 {
				p4 = i32(4)
			}
			v1 = v1 << 2
			if uint32(t3) < uint32(p4+v1) {
				m.fn3(i32(1273840), i32(46), i32(1273888))
				panic("unreachable")
			}
			if v3 == 0 {
				goto l2
			}
			if uint32(v4) > uint32(v1+i32(39)) {
				m.fn3(i32(1273904), i32(46), i32(1273952))
				panic("unreachable")
			}
		l2:
			m.fn1(v2)
		}
	l0:
		{
			t5 := int32(load32(m.memory[int64(uint32(v0))+100:]))
			v1 = t5
			if v1 == 0 {
				goto l4
			}
			t6 := int32(load32(m.memory[int64(uint32(v0))+104:]))
			v2 = t6
			t7 := int32(load32(m.memory[uint32(v2+i32(-4)):]))
			v3 = t7
			v4 = v3 & i32(-8)
			t8 := v4
			v3 = v3 & i32(3)
			p9 := i32(8)
			if v3 != 0 {
				p9 = i32(4)
			}
			v1 = v1 << 2
			if uint32(t8) < uint32(p9+v1) {
				m.fn3(i32(1273840), i32(46), i32(1273888))
				panic("unreachable")
			}
			if v3 == 0 {
				goto l6
			}
			if uint32(v4) > uint32(v1+i32(39)) {
				m.fn3(i32(1273904), i32(46), i32(1273952))
				panic("unreachable")
			}
		l6:
			m.fn1(v2)
		}
	l4:
		{
			t10 := int32(load32(m.memory[int64(uint32(v0))+112:]))
			v1 = t10
			if v1 == 0 {
				goto l8
			}
			t11 := int32(load32(m.memory[int64(uint32(v0))+116:]))
			v2 = t11
			t12 := int32(load32(m.memory[uint32(v2+i32(-4)):]))
			v3 = t12
			v4 = v3 & i32(-8)
			t13 := v4
			v3 = v3 & i32(3)
			p14 := i32(8)
			if v3 != 0 {
				p14 = i32(4)
			}
			v1 = v1 << 2
			if uint32(t13) < uint32(p14+v1) {
				m.fn3(i32(1273840), i32(46), i32(1273888))
				panic("unreachable")
			}
			if v3 == 0 {
				goto l10
			}
			if uint32(v4) > uint32(v1+i32(39)) {
				m.fn3(i32(1273904), i32(46), i32(1273952))
				panic("unreachable")
			}
		l10:
			m.fn1(v2)
		}
	l8:
		{
			t15 := int32(load32(m.memory[int64(uint32(v0))+124:]))
			v1 = t15
			if v1 == 0 {
				goto l12
			}
			t16 := int32(load32(m.memory[int64(uint32(v0))+128:]))
			v2 = t16
			t17 := int32(load32(m.memory[uint32(v2+i32(-4)):]))
			v3 = t17
			v4 = v3 & i32(-8)
			t18 := v4
			v3 = v3 & i32(3)
			p19 := i32(8)
			if v3 != 0 {
				p19 = i32(4)
			}
			v1 = v1 << 2
			if uint32(t18) < uint32(p19+v1) {
				m.fn3(i32(1273840), i32(46), i32(1273888))
				panic("unreachable")
			}
			if v3 == 0 {
				goto l14
			}
			if uint32(v4) > uint32(v1+i32(39)) {
				m.fn3(i32(1273904), i32(46), i32(1273952))
				panic("unreachable")
			}
		l14:
			m.fn1(v2)
		}
	l12:
		m.fn118(v0 + i32(328))
		m.fn118(v0 + i32(340))
		m.fn550(v0 + i32(56))
		{
			t20 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			v1 = t20
			if v1 == i32(-1) {
				goto l16
			}
			{
				if v1 == 0 {
					goto l17
				}
				t21 := int32(load32(m.memory[int64(uint32(v0))+12:]))
				v2 = t21
				t22 := int32(load32(m.memory[uint32(v2+i32(-4)):]))
				v3 = t22
				v4 = v3 & i32(-8)
				t23 := v4
				v3 = v3 & i32(3)
				p24 := i32(8)
				if v3 != 0 {
					p24 = i32(4)
				}
				v1 = v1 << 1
				if uint32(t23) < uint32(p24+v1) {
					m.fn3(i32(1273840), i32(46), i32(1273888))
					panic("unreachable")
				}
				if v3 == 0 {
					goto l19
				}
				if uint32(v4) > uint32(v1+i32(39)) {
					m.fn3(i32(1273904), i32(46), i32(1273952))
					panic("unreachable")
				}
			l19:
				m.fn1(v2)
			}
		l17:
			t25 := int32(load32(m.memory[int64(uint32(v0))+20:]))
			v1 = t25
			if v1 == 0 {
				goto l16
			}
			t26 := int32(load32(m.memory[int64(uint32(v0))+24:]))
			v2 = t26
			t27 := int32(load32(m.memory[uint32(v2+i32(-4)):]))
			v3 = t27
			v4 = v3 & i32(-8)
			t28 := v4
			v3 = v3 & i32(3)
			p29 := i32(8)
			if v3 != 0 {
				p29 = i32(4)
			}
			v1 = v1 << 2
			if uint32(t28) < uint32(p29+v1) {
				m.fn3(i32(1273840), i32(46), i32(1273888))
				panic("unreachable")
			}
			if v3 == 0 {
				goto l22
			}
			if uint32(v4) > uint32(v1+i32(39)) {
				m.fn3(i32(1273904), i32(46), i32(1273952))
				panic("unreachable")
			}
		l22:
			m.fn1(v2)
		}
	l16:
		{
			t30 := int32(load32(m.memory[int64(uint32(v0))+140:]))
			v5 = t30
			if v5 == 0 {
				goto l24
			}
			{
				t31 := int32(load32(m.memory[int64(uint32(v0))+148:]))
				v4 = t31
				if v4 == 0 {
					goto l25
				}
				t32 := int32(load32(m.memory[int64(uint32(v0))+136:]))
				v1 = t32
				v3 = v1 + i32(8)
				t33 := int64(load64(m.memory[uint32(v1):]))
				v6 = (t33 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
			l28:
				if v6 != i64(0) {
					goto l26
				}
			l27:
				{
					v2 = v3
					v3 = v2 + i32(8)
					v1 = v1 + i32(-4160)
					t34 := int64(load64(m.memory[uint32(v2):]))
					v6 = t34 & i64(-0x7f7f7f7f7f7f7f80)
					if v6 == i64(-0x7f7f7f7f7f7f7f80) {
						goto l27
					}
				}
				v6 = v6 ^ i64(-0x7f7f7f7f7f7f7f80)
			l26:
				m.fn412(v1 + (i32(0)-int32(uint32(int64(bits.TrailingZeros64(uint64(v6))))>>3))*i32(520) + i32(-368))
				v6 = (v6 + i64(-1)) & v6
				v4 = v4 + i32(-1)
				if v4 != 0 {
					goto l28
				}
			}
		l25:
			v3 = v5 * i32(520)
			v1 = v3 + v5 + i32(529)
			if v1 == 0 {
				goto l24
			}
			t35 := int32(load32(m.memory[int64(uint32(v0))+136:]))
			v2 = t35 - v3
			t36 := int32(load32(m.memory[uint32(v2+i32(-524)):]))
			v3 = t36
			v4 = v3 & i32(-8)
			t37 := v4
			v3 = v3 & i32(3)
			p38 := i32(8)
			if v3 != 0 {
				p38 = i32(4)
			}
			if uint32(t37) < uint32(p38+v1) {
				m.fn3(i32(1273840), i32(46), i32(1273888))
				panic("unreachable")
			}
			if v3 == 0 {
				goto l30
			}
			if uint32(v4) > uint32(v1+i32(39)) {
				m.fn3(i32(1273904), i32(46), i32(1273952))
				panic("unreachable")
			}
		l30:
			m.fn1(v2 + i32(-520))
		}
	l24:
		t39 := int32(load32(m.memory[int64(uint32(v0))+356:]))
		v7 = t39
		{
			t40 := int32(load32(m.memory[int64(uint32(v0))+360:]))
			v3 = t40
			if v3 == 0 {
				goto l32
			}
			v1 = v7
		l37:
			{
				t41 := int32(load32(m.memory[uint32(v1):]))
				v2 = t41
				if v2 == 0 {
					goto l33
				}
				t42 := int32(load32(m.memory[uint32(v1+i32(4)):]))
				v5 = t42
				t43 := int32(load32(m.memory[uint32(v5+i32(-4)):]))
				v4 = t43
				v8 = v4 & i32(-8)
				t44 := v8
				v4 = v4 & i32(3)
				p45 := i32(8)
				if v4 != 0 {
					p45 = i32(4)
				}
				if uint32(t44) < uint32(p45+v2) {
					m.fn3(i32(1273840), i32(46), i32(1273888))
					panic("unreachable")
				}
				if v4 == 0 {
					goto l35
				}
				if uint32(v8) > uint32(v2+i32(39)) {
					m.fn3(i32(1273904), i32(46), i32(1273952))
					panic("unreachable")
				}
			l35:
				m.fn1(v5)
			}
		l33:
			v1 = v1 + i32(12)
			v3 = v3 + i32(-1)
			if v3 != 0 {
				goto l37
			}
		}
	l32:
		{
			t46 := int32(load32(m.memory[int64(uint32(v0))+352:]))
			v1 = t46
			if v1 == 0 {
				goto l38
			}
			t47 := int32(load32(m.memory[uint32(v7+i32(-4)):]))
			v3 = t47
			v2 = v3 & i32(-8)
			t48 := v2
			v3 = v3 & i32(3)
			p49 := i32(8)
			if v3 != 0 {
				p49 = i32(4)
			}
			v1 = v1 * i32(12)
			if uint32(t48) < uint32(p49+v1) {
				m.fn3(i32(1273840), i32(46), i32(1273888))
				panic("unreachable")
			}
			if v3 == 0 {
				goto l40
			}
			if uint32(v2) > uint32(v1+i32(39)) {
				m.fn3(i32(1273904), i32(46), i32(1273952))
				panic("unreachable")
			}
		l40:
			m.fn1(v7)
		}
	l38:
		{
			t50 := int32(load32(m.memory[int64(uint32(v0))+364:]))
			v1 = t50
			if v1 == 0 {
				goto l42
			}
			t51 := int32(load32(m.memory[int64(uint32(v0))+368:]))
			v2 = t51
			t52 := int32(load32(m.memory[uint32(v2+i32(-4)):]))
			v3 = t52
			v4 = v3 & i32(-8)
			t53 := v4
			v3 = v3 & i32(3)
			p54 := i32(8)
			if v3 != 0 {
				p54 = i32(4)
			}
			v1 = v1 << 3
			if uint32(t53) < uint32(p54+v1) {
				m.fn3(i32(1273840), i32(46), i32(1273888))
				panic("unreachable")
			}
			if v3 == 0 {
				goto l44
			}
			if uint32(v4) > uint32(v1+i32(39)) {
				m.fn3(i32(1273904), i32(46), i32(1273952))
				panic("unreachable")
			}
		l44:
			m.fn1(v2)
		}
	l42:
		{
			t55 := int32(load32(m.memory[int64(uint32(v0))+172:]))
			v9 = t55
			if v9 == 0 {
				goto l46
			}
			{
				t56 := int32(load32(m.memory[int64(uint32(v0))+180:]))
				v4 = t56
				if v4 == 0 {
					goto l47
				}
				t57 := int32(load32(m.memory[int64(uint32(v0))+168:]))
				v1 = t57
				v3 = v1 + i32(8)
				t58 := int64(load64(m.memory[uint32(v1):]))
				v6 = (t58 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
			l54:
				if v6 != i64(0) {
					goto l48
				}
			l49:
				{
					v2 = v3
					v3 = v2 + i32(8)
					v1 = v1 + i32(-128)
					t59 := int64(load64(m.memory[uint32(v2):]))
					v6 = t59 & i64(-0x7f7f7f7f7f7f7f80)
					if v6 == i64(-0x7f7f7f7f7f7f7f80) {
						goto l49
					}
				}
				v6 = v6 ^ i64(-0x7f7f7f7f7f7f7f80)
			l48:
				{
					v5 = v1 - int32(int64(bits.TrailingZeros64(uint64(v6))))<<1&i32(240)
					t60 := int32(load32(m.memory[uint32(v5+i32(-12)):]))
					v2 = t60
					if v2 == 0 {
						goto l50
					}
					t61 := int32(load32(m.memory[uint32(v5+i32(-8)):]))
					v8 = t61
					t62 := int32(load32(m.memory[uint32(v8+i32(-4)):]))
					v5 = t62
					v7 = v5 & i32(-8)
					t63 := v7
					v5 = v5 & i32(3)
					p64 := i32(8)
					if v5 != 0 {
						p64 = i32(4)
					}
					if uint32(t63) < uint32(p64+v2) {
						m.fn3(i32(1273840), i32(46), i32(1273888))
						panic("unreachable")
					}
					if v5 == 0 {
						goto l52
					}
					if uint32(v7) > uint32(v2+i32(39)) {
						m.fn3(i32(1273904), i32(46), i32(1273952))
						panic("unreachable")
					}
				l52:
					m.fn1(v8)
				}
			l50:
				v6 = (v6 + i64(-1)) & v6
				v4 = v4 + i32(-1)
				if v4 != 0 {
					goto l54
				}
			}
		l47:
			v3 = v9 << 4
			v1 = v3 + v9 + i32(25)
			if v1 == 0 {
				goto l46
			}
			t65 := int32(load32(m.memory[int64(uint32(v0))+168:]))
			v2 = t65 - v3
			t66 := int32(load32(m.memory[uint32(v2+i32(-20)):]))
			v3 = t66
			v4 = v3 & i32(-8)
			t67 := v4
			v3 = v3 & i32(3)
			p68 := i32(8)
			if v3 != 0 {
				p68 = i32(4)
			}
			if uint32(t67) < uint32(p68+v1) {
				m.fn3(i32(1273840), i32(46), i32(1273888))
				panic("unreachable")
			}
			if v3 == 0 {
				goto l56
			}
			if uint32(v4) > uint32(v1+i32(39)) {
				m.fn3(i32(1273904), i32(46), i32(1273952))
				panic("unreachable")
			}
		l56:
			m.fn1(v2 + i32(-16))
		}
	l46:
		{
			t69 := int32(load32(m.memory[int64(uint32(v0))+212:]))
			v1 = t69
			if v1 == 0 {
				goto l58
			}
			v3 = v1 * i32(96)
			v1 = v3 + v1 + i32(105)
			if v1 == 0 {
				goto l58
			}
			t70 := int32(load32(m.memory[int64(uint32(v0))+208:]))
			v2 = t70 - v3
			t71 := int32(load32(m.memory[uint32(v2+i32(-100)):]))
			v3 = t71
			v4 = v3 & i32(-8)
			t72 := v4
			v3 = v3 & i32(3)
			p73 := i32(8)
			if v3 != 0 {
				p73 = i32(4)
			}
			if uint32(t72) < uint32(p73+v1) {
				m.fn3(i32(1273840), i32(46), i32(1273888))
				panic("unreachable")
			}
			if v3 == 0 {
				goto l60
			}
			if uint32(v4) > uint32(v1+i32(39)) {
				m.fn3(i32(1273904), i32(46), i32(1273952))
				panic("unreachable")
			}
		l60:
			m.fn1(v2 + i32(-96))
		}
	l58:
		{
			t74 := int32(load32(m.memory[int64(uint32(v0))+244:]))
			v1 = t74
			if v1 == 0 {
				goto l62
			}
			v3 = v1 << 3
			v1 = v3 + v1 + i32(17)
			if v1 == 0 {
				goto l62
			}
			t75 := int32(load32(m.memory[int64(uint32(v0))+240:]))
			v2 = t75 - v3
			t76 := int32(load32(m.memory[uint32(v2+i32(-12)):]))
			v3 = t76
			v4 = v3 & i32(-8)
			t77 := v4
			v3 = v3 & i32(3)
			p78 := i32(8)
			if v3 != 0 {
				p78 = i32(4)
			}
			if uint32(t77) < uint32(p78+v1) {
				m.fn3(i32(1273840), i32(46), i32(1273888))
				panic("unreachable")
			}
			if v3 == 0 {
				goto l64
			}
			if uint32(v4) > uint32(v1+i32(39)) {
				m.fn3(i32(1273904), i32(46), i32(1273952))
				panic("unreachable")
			}
		l64:
			m.fn1(v2 + i32(-8))
		}
	l62:
		{
			t79 := int32(load32(m.memory[int64(uint32(v0))+376:]))
			v1 = t79
			if v1 == 0 {
				goto l66
			}
			t80 := int32(load32(m.memory[int64(uint32(v0))+380:]))
			v2 = t80
			t81 := int32(load32(m.memory[uint32(v2+i32(-4)):]))
			v3 = t81
			v4 = v3 & i32(-8)
			t82 := v4
			v3 = v3 & i32(3)
			p83 := i32(8)
			if v3 != 0 {
				p83 = i32(4)
			}
			if uint32(t82) < uint32(p83+v1) {
				m.fn3(i32(1273840), i32(46), i32(1273888))
				panic("unreachable")
			}
			if v3 == 0 {
				goto l68
			}
			if uint32(v4) > uint32(v1+i32(39)) {
				m.fn3(i32(1273904), i32(46), i32(1273952))
				panic("unreachable")
			}
		l68:
			m.fn1(v2)
		}
	l66:
		m.fn388(v0 + i32(316))
		m.fn389(v0 + i32(280))
		return
	}
}
func (m *Module) fn556(v0 int32) {
	var v1, v2, v3, v4, v5, v6 int32
	t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
	v1 = t0
	t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t2 := v1
	v2 = t1
	t3 := int32(uint32(t2-v2) / uint32(i32(24)))
	v3 = t3
	if v1 == v2 {
		goto l0
	}
l5:
	{
		t4 := int32(load32(m.memory[uint32(v2):]))
		v1 = t4
		if v1 == 0 {
			goto l1
		}
		t5 := int32(load32(m.memory[uint32(v2+i32(4)):]))
		v4 = t5
		t6 := int32(load32(m.memory[uint32(v4+i32(-4)):]))
		v5 = t6
		v6 = v5 & i32(-8)
		t7 := v6
		v5 = v5 & i32(3)
		p8 := i32(8)
		if v5 != 0 {
			p8 = i32(4)
		}
		if uint32(t7) < uint32(p8+v1) {
			m.fn3(i32(1273840), i32(46), i32(1273888))
			panic("unreachable")
		}
		if v5 == 0 {
			goto l3
		}
		if uint32(v6) > uint32(v1+i32(39)) {
			m.fn3(i32(1273904), i32(46), i32(1273952))
			panic("unreachable")
		}
	l3:
		m.fn1(v4)
	}
l1:
	v2 = v2 + i32(24)
	v3 = v3 + i32(-1)
	if v3 != 0 {
		goto l5
	}
l0:
	{
		t9 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		v2 = t9
		if v2 == 0 {
			return
		}
		t10 := int32(load32(m.memory[uint32(v0):]))
		v1 = t10
		t11 := int32(load32(m.memory[uint32(v1+i32(-4)):]))
		v3 = t11
		v5 = v3 & i32(-8)
		t12 := v5
		v3 = v3 & i32(3)
		p13 := i32(8)
		if v3 != 0 {
			p13 = i32(4)
		}
		v2 = v2 * i32(24)
		if uint32(t12) < uint32(p13+v2) {
			m.fn3(i32(1273840), i32(46), i32(1273888))
			panic("unreachable")
		}
		if v3 == 0 {
			goto l8
		}
		if uint32(v5) > uint32(v2+i32(39)) {
			m.fn3(i32(1273904), i32(46), i32(1273952))
			panic("unreachable")
		}
	l8:
		m.fn1(v1)
	}
}
func (m *Module) fn557(v0 int32) {
	var v1, v2 int32
	var v3 int64
	var v4, v5, v6, v7, v8 int32
	t0 := m.g0
	v1 = t0 - i32(64)
	m.g0 = v1
	{
		t1 := int32(load32(m.memory[int64(uint32(v0))+32:]))
		if t1 == 0 {
			goto l0
		}
		t2 := int32(load32(m.memory[int64(uint32(v0))+32:]))
		v2 = t2
		store32(m.memory[int64(uint32(v0))+32:], uint32(i32(0)))
		t3 := int64(load64(m.memory[int64(uint32(v0))+24:]))
		v3 = t3
		store64(m.memory[int64(uint32(v0))+24:], uint64(i64(0x100000000)))
		store32(m.memory[int64(uint32(v1))+16:], uint32(v2))
		store64(m.memory[int64(uint32(v1))+8:], uint64(v3))
		t4 := int32(load32(m.memory[int64(uint32(v0))+36:]))
		v2 = t4
		{
			t5 := int32(load32(m.memory[int64(uint32(v0))+20:]))
			v4 = t5
			if v4 != 0 {
				t11 := int32(load32(m.memory[int64(uint32(v0))+16:]))
				v0 = t11 + v4*i32(28)
				t12 := int32(m.memory[uint32(v0+i32(-4))])
				if t12 != 0 {
					{
						v5 = v0 + i32(-8)
						t26 := int32(load32(m.memory[uint32(v5):]))
						v4 = t26
						t27 := v4
						v8 = v0 + i32(-16)
						t28 := int32(load32(m.memory[uint32(v8):]))
						if t27 != t28 {
							goto l11
						}
						m.fn318(v8)
					}
				l11:
					t29 := int32(load32(m.memory[uint32(v0+i32(-12)):]))
					v0 = t29 + v4*i32(28)
					t30 := int64(load64(m.memory[int64(uint32(v1))+8:]))
					store64(m.memory[int64(uint32(v0))+4:], uint64(t30))
					store32(m.memory[uint32(v0):], uint32(i32(3)))
					t31 := int32(load32(m.memory[int64(uint32(v1))+16:]))
					store32(m.memory[int64(uint32(v0))+12:], uint32(t31))
					store32(m.memory[int64(uint32(v0))+16:], uint32(v2))
					store32(m.memory[uint32(v5):], uint32(v4+i32(1)))
					goto l0
				}
				store32(m.memory[int64(uint32(v1))+24:], uint32(i32(3)))
				t13 := int64(load64(m.memory[int64(uint32(v1))+8:]))
				store64(m.memory[int64(uint32(v1))+28:], uint64(t13))
				t14 := int32(load32(m.memory[int64(uint32(v1))+16:]))
				store32(m.memory[int64(uint32(v1))+36:], uint32(t14))
				store32(m.memory[int64(uint32(v1))+40:], uint32(v2))
				store32(m.memory[int64(uint32(v1))+60:], uint32(i32(0)))
				store64(m.memory[int64(uint32(v1))+52:], uint64(i64(0x100000000)))
				m.fn460(v1+i32(24), i32(1), v1+i32(52))
				t15 := int32(load32(m.memory[int64(uint32(v1))+56:]))
				v5 = t15
				t16 := int32(load32(m.memory[int64(uint32(v1))+52:]))
				v4 = t16
				{
					{
						t17 := int32(load32(m.memory[int64(uint32(v1))+60:]))
						v2 = t17
						t18 := v2
						v6 = v0 + i32(-28)
						t19 := int32(load32(m.memory[uint32(v6):]))
						v7 = v0 + i32(-20)
						t20 := int32(load32(m.memory[uint32(v7):]))
						v8 = t20
						if uint32(t18) <= uint32(t19-v8) {
							goto l4
						}
						m.fn197(v6, v8, v2, i32(1), i32(1))
						t21 := int32(load32(m.memory[uint32(v7):]))
						v8 = t21
						goto l5
					}
				l4:
					if v2 == 0 {
						goto l6
					}
				l5:
					if v2 == 0 {
						goto l6
					}
					t22 := int32(load32(m.memory[uint32(v0+i32(-24)):]))
					memory_copy(m.memory, uint32(t22+v8), uint32(v5), uint32(v2))
				}
			l6:
				store32(m.memory[uint32(v7):], uint32(v8+v2))
				{
					if v4 == 0 {
						goto l7
					}
					t23 := int32(load32(m.memory[uint32(v5+i32(-4)):]))
					v0 = t23
					v2 = v0 & i32(-8)
					t24 := v2
					v0 = v0 & i32(3)
					p25 := i32(8)
					if v0 != 0 {
						p25 = i32(4)
					}
					if uint32(t24) < uint32(p25+v4) {
						m.fn3(i32(1273840), i32(46), i32(1273888))
						panic("unreachable")
					}
					if v0 == 0 {
						goto l9
					}
					if uint32(v2) > uint32(v4+i32(39)) {
						m.fn3(i32(1273904), i32(46), i32(1273952))
						panic("unreachable")
					}
				l9:
					m.fn1(v5)
				}
			l7:
				m.fn337(v1 + i32(24))
				goto l0
			}
			{
				t6 := int32(load32(m.memory[int64(uint32(v0))+8:]))
				v4 = t6
				t7 := int32(load32(m.memory[uint32(v0):]))
				if v4 != t7 {
					goto l2
				}
				m.fn318(v0)
			}
		l2:
			store32(m.memory[int64(uint32(v0))+8:], uint32(v4+i32(1)))
			t8 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			v0 = t8 + v4*i32(28)
			store32(m.memory[uint32(v0):], uint32(i32(3)))
			t9 := int64(load64(m.memory[int64(uint32(v1))+8:]))
			store64(m.memory[int64(uint32(v0))+4:], uint64(t9))
			t10 := int32(load32(m.memory[int64(uint32(v1))+16:]))
			store32(m.memory[int64(uint32(v0))+12:], uint32(t10))
			store32(m.memory[int64(uint32(v0))+16:], uint32(v2))
			goto l0
		}
	}
l0:
	m.g0 = v1 + i32(64)
}
func (m *Module) fn558(v0, v1, v2 int32) int32 {
	var v3, v4, v5, v6, v7, v8, v9, v10, v11, v12 int32
	var v13, v14, v15 int64
	t0 := int32(load32(m.memory[uint32(v0+i32(344)):]))
	v3 = t0
	v4 = i32(0)
	v5 = i32(0)
	{
		t1 := int32(load32(m.memory[uint32(v0+i32(348)):]))
		v6 = t1
		switch v6 {
		case 0:
			goto l0
		default:
			v5 = i32(0)
		l3:
			{
				t2 := v5
				v7 = int32(uint32(v6) >> 1)
				v8 = v7 + v5
				t3 := int32(load32(m.memory[uint32(v3+v8*i32(72)+i32(64)):]))
				p4 := v8
				if uint32(t3) > uint32(v1) {
					p4 = t2
				}
				v5 = p4
				v6 = v6 - v7
				if uint32(v6) > uint32(i32(1)) {
					goto l3
				}
			}
			fallthrough
		case 1:
			t5 := int32(load32(m.memory[uint32(v3+v5*i32(72)+i32(64)):]))
			t6 := v5
			var p7 int32
			if uint32(t5) <= uint32(v1) {
				p7 = 1
			}
			v5 = t6 + p7
			if v5 == 0 {
				goto l0
			}
			v5 = v3 + v5*i32(72)
			if v5 == i32(72) {
				goto l0
			}
			t8 := int32(load32(m.memory[uint32(v5+i32(-4)):]))
			if uint32(v1) >= uint32(t8) {
				goto l0
			}
			t9 := int32(load16(m.memory[uint32(v5+i32(-12)):]))
			v4 = t9
		}
	}
l0:
	v9 = i32(1)
	t10 := int32(load32(m.memory[uint32(v0+i32(332)):]))
	v3 = t10
	v6 = i32(0)
	{
		t11 := int32(load32(m.memory[uint32(v0+i32(336)):]))
		v5 = t11
		switch v5 {
		case 0:
			goto l4
		default:
			v6 = i32(0)
		l7:
			{
				t12 := v6
				v7 = int32(uint32(v5) >> 1)
				v8 = v7 + v6
				t13 := int32(load32(m.memory[uint32(v3+v8*i32(72)+i32(64)):]))
				p14 := v8
				if uint32(t13) > uint32(v1) {
					p14 = t12
				}
				v6 = p14
				v5 = v5 - v7
				if uint32(v5) > uint32(i32(1)) {
					goto l7
				}
			}
			fallthrough
		case 1:
			{
				t15 := int32(load32(m.memory[uint32(v3+v6*i32(72)+i32(64)):]))
				t16 := v6
				var p17 int32
				if uint32(t15) <= uint32(v1) {
					p17 = 1
				}
				v5 = t16 + p17
				if v5 != 0 {
					goto l8
				}
				v5 = i32(0)
				goto l4
			}
		l8:
			v5 = v3 + v5*i32(72)
			if v5 != i32(72) {
				goto l9
			}
			v5 = i32(0)
			goto l4
		l9:
			{
				t18 := int32(load32(m.memory[uint32(v5+i32(-4)):]))
				if uint32(v1) < uint32(t18) {
					goto l10
				}
				v5 = i32(0)
				goto l4
			}
		l10:
			t19 := int32(load32(m.memory[uint32(v5+i32(-20)):]))
			v9 = t19
			t20 := int32(load32(m.memory[uint32(v5+i32(-16)):]))
			v5 = t20
			if uint32(v5) < uint32(i32(2)) {
				goto l4
			}
			v6 = i32(0)
			v10 = i32(0)
			v1 = i32(2)
		l26:
			{
				{
					if uint32(v6) >= uint32(v5) {
						m.fn33(v6, v5, i32(1069384))
						panic("unreachable")
					}
					v7 = v6 + i32(1)
					if uint32(v7) >= uint32(v5) {
						m.fn33(v7, v5, i32(1069400))
						panic("unreachable")
					}
					t21 := int32(m.memory[uint32(v9+v7)])
					v8 = t21
					t22 := int32(m.memory[uint32(v9+v6)])
					v7 = v8<<8 | t22
					v11 = v9 + v1
					v3 = v5 - v1
					v6 = i32(1)
					switch int32(uint32(v8) >> 5) {
					default:
						goto l13
					case 2, 4, 5:
						v6 = i32(2)
						goto l13
					case 3:
						v6 = i32(4)
						goto l13
					case 7:
						v6 = i32(3)
						goto l13
					case 6:
						if v7 == i32(54792) {
							goto l18
						}
						if v5 == v1 {
							v6 = i32(0)
							p24 := v10
							if v7 == i32(18992) {
								p24 = i32(0)
							}
							v10 = p24
							goto l21
						}
						t23 := int32(m.memory[uint32(v11)])
						v6 = t23 + i32(1)
					}
				l13:
					if uint32(v6) > uint32(v3) {
						if v10&i32(1) == 0 {
							goto l4
						}
						goto l23
					}
					if v7 != i32(18992) {
						goto l21
					}
					if uint32(v6) >= uint32(i32(2)) {
						t25 := int32(load16(m.memory[uint32(v11):]))
						v12 = t25
						v10 = i32(1)
						goto l21
					}
					v6 = i32(1)
					v10 = i32(0)
					goto l21
				}
			l18:
				if uint32(v3) > uint32(i32(1)) {
					goto l24
				}
				v6 = i32(0)
				goto l21
			l24:
				t26 := int32(load16(m.memory[uint32(v11):]))
				v6 = t26
				if uint32(v6) >= uint32(v3) {
					goto l25
				}
				v6 = v6 + i32(1)
			}
		l21:
			v6 = v6 + v1
			v1 = v6 + i32(2)
			if uint32(v1) <= uint32(v5) {
				goto l26
			}
			if v10&i32(1) != 0 {
				goto l23
			}
			goto l4
		l25:
			if v10&i32(1) == 0 {
				goto l4
			}
		l23:
			v4 = v12
		}
	}
l4:
	{
		{
			t27 := int32(load32(m.memory[int64(uint32(v0))+68:]))
			if t27 != 0 {
				goto l27
			}
			v6 = i32(0)
			goto l28
		}
	l27:
		t28 := int64(load64(m.memory[int64(uint32(v0))+72:]))
		t29 := int64(load64(m.memory[int64(uint32(v0))+80:]))
		t30 := m.fn106(t28, t29, v4)
		v13 = t30
		t31 := int32(load32(m.memory[int64(uint32(v0))+60:]))
		v8 = t31
		v1 = v8 & int32(v13)
		v14 = int64(uint64(v13)>>25) & i64(127) * i64(72340172838076673)
		t32 := int32(load32(m.memory[int64(uint32(v0))+56:]))
		v7 = t32
		v3 = v4 & i32(0xffff)
		v4 = i32(0)
	l32:
		{
			{
				t33 := int64(load64(m.memory[uint32(v7+v1):]))
				v15 = t33
				v13 = v15 ^ v14
				v13 = (v13 ^ i64(-1)) & (v13 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
				if v13 == 0 {
					goto l29
				}
			l31:
				{
					t34 := v3
					v6 = v7 + (i32(0)-(int32(uint32(int64(bits.TrailingZeros64(uint64(v13))))>>3)+v1)&v8)*i32(60)
					t35 := int32(load16(m.memory[uint32(v6+i32(-60)):]))
					if t34 == t35 {
						goto l30
					}
					v13 = (v13 + i64(-1)) & v13
					if !(v13 == 0) {
						goto l31
					}
				}
			}
		l29:
			v6 = i32(0)
			if !(v15&(v15<<1)&i64(-0x7f7f7f7f7f7f7f80) == 0) {
				goto l30
			}
			t36 := v1
			v4 = v4 + i32(8)
			v1 = (t36 + v4) & v8
			goto l32
		}
	l30:
		p37 := i32(0)
		if v6 != 0 {
			p37 = v6 + i32(-56)
		}
		v6 = p37
	}
l28:
	t39 := v9
	t40 := v5
	p38 := v0
	if v6 != 0 {
		p38 = v6
	}
	t41 := int32(load32(m.memory[int64(uint32(p38))+48:]))
	v6 = t41
	t42 := m.fn548(t39, t40, v6, v6)
	v5 = t42
	{
		t43 := int32(load32(m.memory[int64(uint32(v0))+132:]))
		if uint32(v2) >= uint32(t43) {
			goto l33
		}
		t44 := int32(load32(m.memory[int64(uint32(v0))+128:]))
		t45 := int32(load32(m.memory[uint32(t44+v2<<2):]))
		v1 = t45
		t46 := int32(load32(m.memory[int64(uint32(v0))+372:]))
		if uint32(v1) >= uint32(t46) {
			goto l33
		}
		t47 := int32(load32(m.memory[int64(uint32(v0))+368:]))
		v1 = t47 + v1<<3
		t48 := int32(load32(m.memory[uint32(v1):]))
		if t48 != i32(1) {
			goto l33
		}
		t49 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		v1 = t49
		t50 := int32(load32(m.memory[int64(uint32(v0))+360:]))
		if uint32(v1) >= uint32(t50) {
			goto l33
		}
		t51 := int32(load32(m.memory[int64(uint32(v0))+356:]))
		v1 = t51 + v1*i32(12)
		t52 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t53 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		t54 := m.fn548(t52, t53, v5, v6)
		v5 = t54
	}
l33:
	return v5
}
func (m *Module) fn559(v0 int32) {
	var v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11, v12 int32
	t0 := m.g0
	v1 = t0 - i32(64)
	m.g0 = v1
	m.fn557(v0)
	{
		t1 := int32(load32(m.memory[int64(uint32(v0))+20:]))
		v2 = t1
		if v2 == 0 {
			goto l0
		}
		t2 := v0
		v2 = v2 + i32(-1)
		store32(m.memory[int64(uint32(t2))+20:], uint32(v2))
		t3 := int32(load32(m.memory[int64(uint32(v0))+16:]))
		v2 = t3 + v2*i32(28)
		t4 := int32(load32(m.memory[uint32(v2):]))
		v3 = t4
		t5 := int32(load32(m.memory[int64(uint32(v2))+4:]))
		v4 = t5
		t6 := int32(load32(m.memory[int64(uint32(v2))+8:]))
		v5 = t6
		t7 := int32(load32(m.memory[int64(uint32(v2))+20:]))
		store32(m.memory[int64(uint32(v1))+32:], uint32(t7))
		t8 := int64(load64(m.memory[int64(uint32(v2))+12:]))
		store64(m.memory[int64(uint32(v1))+24:], uint64(t8))
		m.fn479(v1+i32(12), v4, v5, v1+i32(24))
		t9 := int32(load32(m.memory[int64(uint32(v1))+16:]))
		v6 = t9
		t10 := int32(load32(m.memory[int64(uint32(v1))+20:]))
		t11 := v6
		v5 = t10
		v7 = t11 + v5*i32(28)
		t12 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		v8 = t12
		v2 = v6
		{
			if v5 == 0 {
				goto l1
			}
			v5 = v1 + i32(36) + i32(4)
			v2 = v6
		l8:
			{
				{
					t13 := int32(load32(m.memory[uint32(v2):]))
					v9 = t13
					if v9 == i32(-1) {
						v2 = v2 + i32(28)
						goto l1
					}
					t14 := int64(load64(m.memory[uint32(v2+i32(4)):]))
					store64(m.memory[uint32(v5):], uint64(t14))
					t15 := int64(load64(m.memory[uint32(v2+i32(12)):]))
					store64(m.memory[int64(uint32(v5))+8:], uint64(t15))
					t16 := int64(load64(m.memory[uint32(v2+i32(20)):]))
					store64(m.memory[int64(uint32(v5))+16:], uint64(t16))
					store32(m.memory[int64(uint32(v1))+36:], uint32(v9))
					t17 := int32(load32(m.memory[int64(uint32(v0))+20:]))
					v9 = t17
					if v9 != 0 {
						goto l3
					}
					{
						t18 := int32(load32(m.memory[int64(uint32(v0))+8:]))
						v9 = t18
						t19 := int32(load32(m.memory[uint32(v0):]))
						if v9 != t19 {
							goto l4
						}
						m.fn318(v0)
					}
				l4:
					store32(m.memory[int64(uint32(v0))+8:], uint32(v9+i32(1)))
					t20 := int32(load32(m.memory[int64(uint32(v0))+4:]))
					v9 = t20 + v9*i32(28)
					t21 := int64(load64(m.memory[int64(uint32(v1))+36:]))
					store64(m.memory[uint32(v9):], uint64(t21))
					t22 := int64(load64(m.memory[int64(uint32(v1))+44:]))
					store64(m.memory[int64(uint32(v9))+8:], uint64(t22))
					t23 := int64(load64(m.memory[int64(uint32(v1))+52:]))
					store64(m.memory[int64(uint32(v9))+16:], uint64(t23))
					t24 := int32(load32(m.memory[int64(uint32(v1))+60:]))
					store32(m.memory[int64(uint32(v9))+24:], uint32(t24))
					goto l5
				}
			l3:
				{
					t25 := int32(load32(m.memory[int64(uint32(v0))+16:]))
					v9 = t25 + v9*i32(28)
					t26 := int32(m.memory[uint32(v9+i32(-4))])
					if t26 != 0 {
						goto l6
					}
					m.fn337(v1 + i32(36))
					goto l5
				}
			l6:
				{
					v10 = v9 + i32(-8)
					t27 := int32(load32(m.memory[uint32(v10):]))
					v11 = t27
					t28 := v11
					v12 = v9 + i32(-16)
					t29 := int32(load32(m.memory[uint32(v12):]))
					if t28 != t29 {
						goto l7
					}
					m.fn318(v12)
				}
			l7:
				t30 := int32(load32(m.memory[uint32(v9+i32(-12)):]))
				v9 = t30 + v11*i32(28)
				t31 := int32(load32(m.memory[int64(uint32(v1))+60:]))
				store32(m.memory[int64(uint32(v9))+24:], uint32(t31))
				t32 := int64(load64(m.memory[int64(uint32(v1))+52:]))
				store64(m.memory[int64(uint32(v9))+16:], uint64(t32))
				t33 := int64(load64(m.memory[int64(uint32(v1))+44:]))
				store64(m.memory[int64(uint32(v9))+8:], uint64(t33))
				t34 := int64(load64(m.memory[int64(uint32(v1))+36:]))
				store64(m.memory[uint32(v9):], uint64(t34))
				store32(m.memory[uint32(v10):], uint32(v11+i32(1)))
			}
		l5:
			v2 = v2 + i32(28)
			if v2 != v7 {
				goto l8
			}
			goto l9
		l1:
			t35 := int32(uint32(v7-v2) / uint32(i32(28)))
			v0 = t35
			if v7 == v2 {
				goto l9
			}
		l10:
			m.fn337(v2)
			v2 = v2 + i32(28)
			v0 = v0 + i32(-1)
			if v0 != 0 {
				goto l10
			}
		}
	l9:
		{
			if v8 == 0 {
				goto l11
			}
			t36 := int32(load32(m.memory[uint32(v6+i32(-4)):]))
			v2 = t36
			v0 = v2 & i32(-8)
			t37 := v0
			v2 = v2 & i32(3)
			p38 := i32(8)
			if v2 != 0 {
				p38 = i32(4)
			}
			v5 = v8 * i32(28)
			if uint32(t37) < uint32(p38+v5) {
				m.fn3(i32(1273840), i32(46), i32(1273888))
				panic("unreachable")
			}
			if v2 == 0 {
				goto l13
			}
			if uint32(v0) > uint32(v5+i32(39)) {
				m.fn3(i32(1273904), i32(46), i32(1273952))
				panic("unreachable")
			}
		l13:
			m.fn1(v6)
		}
	l11:
		if v3 == 0 {
			goto l0
		}
		t39 := int32(load32(m.memory[uint32(v4+i32(-4)):]))
		v2 = t39
		v0 = v2 & i32(-8)
		t40 := v0
		v2 = v2 & i32(3)
		p41 := i32(8)
		if v2 != 0 {
			p41 = i32(4)
		}
		if uint32(t40) < uint32(p41+v3) {
			m.fn3(i32(1273840), i32(46), i32(1273888))
			panic("unreachable")
		}
		if v2 == 0 {
			goto l16
		}
		if uint32(v0) > uint32(v3+i32(39)) {
			m.fn3(i32(1273904), i32(46), i32(1273952))
			panic("unreachable")
		}
	l16:
		m.fn1(v4)
	}
l0:
	m.g0 = v1 + i32(64)
}
func (m *Module) fn560(v0 int32) {
	var v1, v2, v3, v4, v5, v6, v7, v8 int32
	t0 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	v1 = t0
	{
		t1 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		v2 = t1
		if v2 == 0 {
			goto l0
		}
		v3 = i32(0)
	l11:
		{
			{
				v4 = v1 + v3*i32(28)
				t2 := int32(load32(m.memory[uint32(v4):]))
				v5 = t2
				if v5 == 0 {
					goto l1
				}
				t3 := int32(load32(m.memory[int64(uint32(v4))+4:]))
				v6 = t3
				t4 := int32(load32(m.memory[uint32(v6+i32(-4)):]))
				v7 = t4
				v8 = v7 & i32(-8)
				t5 := v8
				v7 = v7 & i32(3)
				p6 := i32(8)
				if v7 != 0 {
					p6 = i32(4)
				}
				if uint32(t5) < uint32(p6+v5) {
					m.fn3(i32(1273840), i32(46), i32(1273888))
					panic("unreachable")
				}
				if v7 == 0 {
					goto l3
				}
				if uint32(v8) > uint32(v5+i32(39)) {
					m.fn3(i32(1273904), i32(46), i32(1273952))
					panic("unreachable")
				}
			l3:
				m.fn1(v6)
			}
		l1:
			t7 := int32(load32(m.memory[int64(uint32(v4))+16:]))
			v6 = t7
			{
				t8 := int32(load32(m.memory[int64(uint32(v4))+20:]))
				v7 = t8
				if v7 == 0 {
					goto l5
				}
				v5 = v6
			l6:
				m.fn337(v5)
				v5 = v5 + i32(28)
				v7 = v7 + i32(-1)
				if v7 != 0 {
					goto l6
				}
			}
		l5:
			{
				t9 := int32(load32(m.memory[int64(uint32(v4))+12:]))
				v5 = t9
				if v5 == 0 {
					goto l7
				}
				t10 := int32(load32(m.memory[uint32(v6+i32(-4)):]))
				v7 = t10
				v4 = v7 & i32(-8)
				t11 := v4
				v7 = v7 & i32(3)
				p12 := i32(8)
				if v7 != 0 {
					p12 = i32(4)
				}
				v5 = v5 * i32(28)
				if uint32(t11) < uint32(p12+v5) {
					m.fn3(i32(1273840), i32(46), i32(1273888))
					panic("unreachable")
				}
				if v7 == 0 {
					goto l9
				}
				if uint32(v4) > uint32(v5+i32(39)) {
					m.fn3(i32(1273904), i32(46), i32(1273952))
					panic("unreachable")
				}
			l9:
				m.fn1(v6)
			}
		l7:
			v3 = v3 + i32(1)
			if v3 != v2 {
				goto l11
			}
		}
	}
l0:
	{
		t13 := int32(load32(m.memory[uint32(v0):]))
		v5 = t13
		if v5 == 0 {
			return
		}
		t14 := int32(load32(m.memory[uint32(v1+i32(-4)):]))
		v7 = t14
		v4 = v7 & i32(-8)
		t15 := v4
		v7 = v7 & i32(3)
		p16 := i32(8)
		if v7 != 0 {
			p16 = i32(4)
		}
		v5 = v5 * i32(28)
		if uint32(t15) < uint32(p16+v5) {
			m.fn3(i32(1273840), i32(46), i32(1273888))
			panic("unreachable")
		}
		if v7 == 0 {
			goto l14
		}
		if uint32(v4) > uint32(v5+i32(39)) {
			m.fn3(i32(1273904), i32(46), i32(1273952))
			panic("unreachable")
		}
	l14:
		m.fn1(v1)
	}
}
func (m *Module) fn561(v0, v1, v2, v3, v4 int32) {
	var v5, v6 int32
	var v7 int64
	var v8, v9, v10, v11, v12, v13 int32
	var v14 int64
	var v15, v16 int32
	var v17 int64
	var v18, v19, v20, v21, v22, v23, v24, v25 int32
	var v26 int64
	var v27, v28, v29, v30 int32
	var v31, v32 int64
	t0 := m.g0
	v5 = t0 - i32(80)
	m.g0 = v5
	{
		{
			t1 := int32(load32(m.memory[int64(uint32(v4))+8:]))
			if t1 != 0 {
				goto l0
			}
			t2 := int32(load32(m.memory[int64(uint32(v3))+8:]))
			v4 = t2
			goto l1
		}
	l0:
		t3 := int32(load32(m.memory[int64(uint32(v4))+8:]))
		v6 = t3
		store32(m.memory[int64(uint32(v4))+8:], uint32(i32(0)))
		t4 := int64(load64(m.memory[uint32(v4):]))
		v7 = t4
		store64(m.memory[uint32(v4):], uint64(i64(0x800000000)))
		store32(m.memory[int64(uint32(v5))+56:], uint32(v6))
		store64(m.memory[int64(uint32(v5))+48:], uint64(v7))
		{
			t5 := int32(load32(m.memory[int64(uint32(v3))+8:]))
			v6 = t5
			t6 := int32(load32(m.memory[uint32(v3):]))
			if v6 != t6 {
				goto l2
			}
			m.fn316(v3)
		}
	l2:
		t7 := v3
		v4 = v6 + i32(1)
		store32(m.memory[int64(uint32(t7))+8:], uint32(v4))
		t8 := int32(load32(m.memory[int64(uint32(v3))+4:]))
		v6 = t8 + v6*i32(12)
		t9 := int64(load64(m.memory[int64(uint32(v5))+48:]))
		store64(m.memory[uint32(v6):], uint64(t9))
		t10 := int32(load32(m.memory[int64(uint32(v5))+56:]))
		store32(m.memory[int64(uint32(v6))+8:], uint32(t10))
	}
l1:
	{
		{
			{
				if v4 != 0 {
					goto l3
				}
				t11 := int32(load32(m.memory[int64(uint32(v2))+8:]))
				v3 = t11
				if v3 == 0 {
					goto l4
				}
				goto l5
			}
		l3:
			t12 := int32(load32(m.memory[int64(uint32(v3))+8:]))
			v4 = t12
			store32(m.memory[int64(uint32(v3))+8:], uint32(i32(0)))
			t13 := int64(load64(m.memory[uint32(v3):]))
			v7 = t13
			store64(m.memory[uint32(v3):], uint64(i64(0x400000000)))
			store32(m.memory[int64(uint32(v5))+56:], uint32(v4))
			store64(m.memory[int64(uint32(v5))+48:], uint64(v7))
			{
				t14 := int32(load32(m.memory[int64(uint32(v2))+8:]))
				v4 = t14
				t15 := int32(load32(m.memory[uint32(v2):]))
				if v4 != t15 {
					goto l6
				}
				m.fn321(v2)
			}
		l6:
			t16 := v2
			v3 = v4 + i32(1)
			store32(m.memory[int64(uint32(t16))+8:], uint32(v3))
			t17 := int32(load32(m.memory[int64(uint32(v2))+4:]))
			v4 = t17 + v4*i32(40)
			t18 := int64(load64(m.memory[int64(uint32(v5))+48:]))
			store64(m.memory[uint32(v4):], uint64(t18))
			t19 := int32(load32(m.memory[int64(uint32(v5))+56:]))
			store32(m.memory[int64(uint32(v4))+8:], uint32(t19))
			store32(m.memory[int64(uint32(v4))+12:], uint32(i32(-1)))
			if v3 != 0 {
				goto l5
			}
		}
	l4:
		store32(m.memory[uint32(v0):], uint32(i32(-1)))
		goto l7
	l5:
		t20 := int32(load32(m.memory[int64(uint32(v2))+4:]))
		v8 = t20
		store64(m.memory[int64(uint32(v2))+4:], uint64(i64(4)))
		t21 := int32(load32(m.memory[uint32(v2):]))
		v4 = t21
		store32(m.memory[uint32(v2):], uint32(i32(0)))
		v9 = v4 * i32(40)
		v10 = v8 + v3*i32(40)
		v11 = v8
		v12 = v8
	l30:
		{
			{
				{
					t22 := int32(load32(m.memory[int64(uint32(v12))+12:]))
					v13 = t22
					if v13 != i32(-1) {
						goto l8
					}
					v14 = i64(0)
					v15 = i32(1)
					v16 = i32(0)
					v17 = i64(0)
					v18 = i32(2)
					v13 = i32(0)
					v19 = i32(0)
					goto l9
				}
			l8:
				t23 := int64(load32(m.memory[int64(uint32(v12))+20:]))
				v17 = t23
				t24 := int64(load32(m.memory[int64(uint32(v12))+32:]))
				v14 = t24
				t25 := int32(m.memory[int64(uint32(v12))+36])
				v19 = t25
				t26 := int32(load32(m.memory[int64(uint32(v12))+16:]))
				v18 = t26
				t27 := int32(load32(m.memory[int64(uint32(v12))+24:]))
				v16 = t27
				t28 := int32(load32(m.memory[int64(uint32(v12))+28:]))
				v15 = t28
			}
		l9:
			t29 := int32(load32(m.memory[int64(uint32(v12))+8:]))
			v3 = t29
			if uint32(v3) > uint32(i32(0x7ffffff)) {
				goto l10
			}
			v2 = v3 << 5
			if uint32(v2) >= uint32(i32(0x7ffffff9)) {
				goto l10
			}
			t30 := int32(load32(m.memory[int64(uint32(v12))+4:]))
			v20 = t30
			t31 := int32(load32(m.memory[uint32(v12):]))
			v21 = t31
			v22 = i32(0)
			{
				if v2 != 0 {
					goto l11
				}
				v23 = i32(8)
				v24 = i32(0)
				goto l12
			l11:
				v24 = v3
				t32 := m.fn5(v2)
				v23 = t32
				if v23 == 0 {
					m.fn10(i32(8), v2)
					panic("unreachable")
				}
			}
		l12:
			{
				if v3 == 0 {
					goto l14
				}
				v6 = v18 + i32(2)
				v25 = v3 * i32(12)
				t33 := int32(uint32(v25+i32(-12)) / uint32(i32(12)))
				v22 = t33 + i32(1)
				v26 = i64(1000)
				v7 = i64(1)
				v3 = v23
				v2 = v15
				v4 = v20
			l17:
				{
					v27 = i32(0)
					v28 = i32(0)
					v29 = i32(0)
					v30 = i32(0)
					{
						if uint64(v7+i64(-1)) >= uint64(v14) {
							goto l15
						}
						t34 := int32(m.memory[uint32(v2+i32(3))])
						v30 = t34
						t35 := int32(m.memory[uint32(v2+i32(2))])
						v29 = t35
						t36 := int32(m.memory[uint32(v2+i32(1))])
						v28 = t36
						t37 := int32(m.memory[uint32(v2)])
						v27 = t37
					}
				l15:
					v31 = v26
					{
						if uint64(v7) >= uint64(v17) {
							goto l16
						}
						t38 := int64(int16(load16(m.memory[uint32(v6):])))
						v31 = t38
					}
				l16:
					t39 := int64(load64(m.memory[uint32(v4):]))
					v32 = t39
					t40 := int32(load32(m.memory[int64(uint32(v4))+8:]))
					store32(m.memory[int64(uint32(v3))+8:], uint32(t40))
					store64(m.memory[uint32(v3):], uint64(v32))
					m.memory[uint32(v3+i32(27))] = byte(v30 & i32(1))
					m.memory[uint32(v3+i32(26))] = byte(v29 & i32(1))
					m.memory[uint32(v3+i32(25))] = byte(v28 & i32(1))
					m.memory[uint32(v3+i32(24))] = byte(v27 & i32(1))
					store64(m.memory[uint32(v3+i32(16)):], uint64(v31))
					v3 = v3 + i32(32)
					v7 = v7 + i64(1)
					v6 = v6 + i32(2)
					v26 = v26 + i64(1000)
					v2 = v2 + i32(4)
					v4 = v4 + i32(12)
					v25 = v25 + i32(-12)
					if v25 != 0 {
						goto l17
					}
				}
			}
		l14:
			{
				if v21 == 0 {
					goto l18
				}
				t41 := int32(load32(m.memory[uint32(v20+i32(-4)):]))
				v3 = t41
				v2 = v3 & i32(-8)
				t42 := v2
				v3 = v3 & i32(3)
				p43 := i32(8)
				if v3 != 0 {
					p43 = i32(4)
				}
				v4 = v21 * i32(12)
				if uint32(t42) < uint32(p43+v4) {
					m.fn3(i32(1273840), i32(46), i32(1273888))
					panic("unreachable")
				}
				if v3 == 0 {
					goto l20
				}
				if uint32(v2) > uint32(v4+i32(39)) {
					m.fn3(i32(1273904), i32(46), i32(1273952))
					panic("unreachable")
				}
			l20:
				m.fn1(v20)
			}
		l18:
			{
				if v13 == 0 {
					goto l22
				}
				t44 := int32(load32(m.memory[uint32(v18+i32(-4)):]))
				v3 = t44
				v2 = v3 & i32(-8)
				t45 := v2
				v3 = v3 & i32(3)
				p46 := i32(8)
				if v3 != 0 {
					p46 = i32(4)
				}
				v4 = v13 << 1
				if uint32(t45) < uint32(p46+v4) {
					m.fn3(i32(1273840), i32(46), i32(1273888))
					panic("unreachable")
				}
				if v3 == 0 {
					goto l24
				}
				if uint32(v2) > uint32(v4+i32(39)) {
					m.fn3(i32(1273904), i32(46), i32(1273952))
					panic("unreachable")
				}
			l24:
				m.fn1(v18)
			}
		l22:
			{
				if v16 == 0 {
					goto l26
				}
				t47 := int32(load32(m.memory[uint32(v15+i32(-4)):]))
				v3 = t47
				v2 = v3 & i32(-8)
				t48 := v2
				v3 = v3 & i32(3)
				p49 := i32(8)
				if v3 != 0 {
					p49 = i32(4)
				}
				v4 = v16 << 2
				if uint32(t48) < uint32(p49+v4) {
					m.fn3(i32(1273840), i32(46), i32(1273888))
					panic("unreachable")
				}
				if v3 == 0 {
					goto l28
				}
				if uint32(v2) > uint32(v4+i32(39)) {
					m.fn3(i32(1273904), i32(46), i32(1273952))
					panic("unreachable")
				}
			l28:
				m.fn1(v15)
			}
		l26:
			store32(m.memory[int64(uint32(v11))+8:], uint32(v22))
			store32(m.memory[int64(uint32(v11))+4:], uint32(v23))
			store32(m.memory[uint32(v11):], uint32(v24))
			m.memory[int64(uint32(v11))+12] = byte(v19 & i32(1))
			v11 = v11 + i32(16)
			v12 = v12 + i32(40)
			if v12 != v10 {
				goto l30
			}
		}
		v3 = v8
		{
			if v9&i32(8) == 0 {
				goto l31
			}
			v2 = v9 & i32(-16)
			if v2 != 0 {
				goto l32
			}
			v3 = i32(4)
			m.fn18(v8, v9, i32(4))
			goto l31
		l32:
			t50 := m.fn22(v8, v9, i32(4), v2)
			v3 = t50
			if v3 == 0 {
				m.fn24(i32(4), v2)
				panic("unreachable")
			}
		}
	l31:
		store32(m.memory[int64(uint32(v5))+16:], uint32(v3))
		store32(m.memory[int64(uint32(v5))+12:], uint32(int32(uint32(v9)>>4)))
		store32(m.memory[int64(uint32(v5))+20:], uint32(int32(uint32(v11-v8)>>4)))
		m.fn566(v5+i32(48), v5+i32(12))
		t51 := int64(load64(m.memory[int64(uint32(v5))+52:]))
		store64(m.memory[int64(uint32(v5))+24:], uint64(t51))
		t52 := int64(load64(m.memory[int64(uint32(v5))+60:]))
		store64(m.memory[int64(uint32(v5))+32:], uint64(t52))
		t53 := int64(load64(m.memory[int64(uint32(v5))+68:]))
		store64(m.memory[int64(uint32(v5))+40:], uint64(t53))
		{
			t54 := int32(load32(m.memory[int64(uint32(v5))+48:]))
			v3 = t54
			if v3 != i32(-2) {
				goto l34
			}
			t55 := int64(load64(m.memory[int64(uint32(v5))+40:]))
			store64(m.memory[int64(uint32(v0))+16:], uint64(t55))
			t56 := int64(load64(m.memory[int64(uint32(v5))+32:]))
			store64(m.memory[int64(uint32(v0))+8:], uint64(t56))
			t57 := int64(load64(m.memory[int64(uint32(v5))+24:]))
			store64(m.memory[uint32(v0):], uint64(t57))
			goto l7
		}
	l34:
		{
			if v3 == i32(-1) {
				goto l35
			}
			t58 := int32(load32(m.memory[int64(uint32(v5))+76:]))
			v4 = t58
			{
				t59 := int32(load32(m.memory[int64(uint32(v1))+8:]))
				v2 = t59
				t60 := int32(load32(m.memory[uint32(v1):]))
				if v2 != t60 {
					goto l36
				}
				m.fn315(v1)
			}
		l36:
			store32(m.memory[int64(uint32(v1))+8:], uint32(v2+i32(1)))
			t61 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			v2 = t61 + v2<<5
			store32(m.memory[uint32(v2):], uint32(v3))
			t62 := int64(load64(m.memory[int64(uint32(v5))+24:]))
			store64(m.memory[int64(uint32(v2))+4:], uint64(t62))
			t63 := int64(load64(m.memory[int64(uint32(v5))+32:]))
			store64(m.memory[int64(uint32(v2))+12:], uint64(t63))
			t64 := int64(load64(m.memory[int64(uint32(v5))+40:]))
			store64(m.memory[int64(uint32(v2))+20:], uint64(t64))
			store32(m.memory[int64(uint32(v2))+28:], uint32(v4))
		}
	l35:
		store32(m.memory[uint32(v0):], uint32(i32(-1)))
	}
l7:
	m.g0 = v5 + i32(80)
	return
l10:
	m.fn9()
	panic("unreachable")
}
func (m *Module) fn562(v0, v1, v2, v3, v4 int32) {
	var v5 int64
	var v6, v7 int32
	var v8 int64
	var v9, v10, v11 int32
	var v12 int64
	{
		{
			t0 := int32(load32(m.memory[int64(uint32(v0))+68:]))
			if t0 != 0 {
				goto l0
			}
			v1 = i32(0)
			goto l1
		}
	l0:
		t1 := int64(load64(m.memory[int64(uint32(v0))+72:]))
		t2 := int64(load64(m.memory[int64(uint32(v0))+80:]))
		t3 := m.fn106(t1, t2, v1)
		v5 = t3
		t4 := int32(load32(m.memory[int64(uint32(v0))+60:]))
		v6 = t4
		v7 = v6 & int32(v5)
		v8 = int64(uint64(v5)>>25) & i64(127) * i64(72340172838076673)
		t5 := int32(load32(m.memory[int64(uint32(v0))+56:]))
		v9 = t5
		v10 = v1 & i32(0xffff)
		v11 = i32(0)
	l5:
		{
			{
				t6 := int64(load64(m.memory[uint32(v9+v7):]))
				v12 = t6
				v5 = v12 ^ v8
				v5 = (v5 ^ i64(-1)) & (v5 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
				if v5 == 0 {
					goto l2
				}
			l4:
				{
					t7 := v10
					v1 = v9 + (i32(0)-(int32(uint32(int64(bits.TrailingZeros64(uint64(v5))))>>3)+v7)&v6)*i32(60)
					t8 := int32(load16(m.memory[uint32(v1+i32(-60)):]))
					if t7 == t8 {
						goto l3
					}
					v5 = (v5 + i64(-1)) & v5
					if !(v5 == 0) {
						goto l4
					}
				}
			}
		l2:
			v1 = i32(0)
			if !(v12&(v12<<1)&i64(-0x7f7f7f7f7f7f7f80) == 0) {
				goto l3
			}
			t9 := v7
			v11 = v11 + i32(8)
			v7 = (t9 + v11) & v6
			goto l5
		}
	l3:
		p10 := i32(0)
		if v1 != 0 {
			p10 = v1 + i32(-56)
		}
		v1 = p10
	}
l1:
	{
		p11 := v0
		if v1 != 0 {
			p11 = v1
		}
		t12 := int32(m.memory[int64(uint32(p11))+54])
		v0 = t12
		if v0 == i32(2) {
			m.fn427(v4, v3)
			t13 := int32(load32(m.memory[int64(uint32(v2))+8:]))
			v1 = t13
			v7 = v1 * i32(28)
			t14 := int32(load32(m.memory[int64(uint32(v2))+4:]))
			v4 = t14
			v0 = i32(0)
			{
			l9:
				{
					if v7 == v0 {
						if v1 == 0 {
							goto l11
						}
						v0 = v4
					l12:
						m.fn337(v0)
						v0 = v0 + i32(28)
						v1 = v1 + i32(-1)
						if v1 != 0 {
							goto l12
						}
					l11:
						t22 := int32(load32(m.memory[uint32(v2):]))
						v0 = t22
						if v0 == 0 {
							return
						}
						{
							t23 := int32(load32(m.memory[uint32(v4+i32(-4)):]))
							v1 = t23
							v7 = v1 & i32(-8)
							t24 := v7
							v1 = v1 & i32(3)
							p25 := i32(8)
							if v1 != 0 {
								p25 = i32(4)
							}
							v0 = v0 * i32(28)
							if uint32(t24) < uint32(p25+v0) {
								m.fn3(i32(1273840), i32(46), i32(1273888))
								panic("unreachable")
							}
							if v1 == 0 {
								goto l14
							}
							if uint32(v7) > uint32(v0+i32(39)) {
								m.fn3(i32(1273904), i32(46), i32(1273952))
								panic("unreachable")
							}
						l14:
							m.fn1(v4)
							return
						}
					}
					t15 := v4
					v0 = v0 + i32(28)
					t16 := m.fn311(t15 + v0 + i32(-28))
					if t16 != 0 {
						goto l9
					}
				}
				{
					t17 := int32(load32(m.memory[int64(uint32(v3))+8:]))
					v0 = t17
					t18 := int32(load32(m.memory[uint32(v3):]))
					if v0 != t18 {
						goto l10
					}
					m.fn315(v3)
				}
			l10:
				store32(m.memory[int64(uint32(v3))+8:], uint32(v0+i32(1)))
				t19 := int32(load32(m.memory[int64(uint32(v3))+4:]))
				v0 = t19 + v0<<5
				store32(m.memory[uint32(v0):], uint32(i32(-0x80000000)))
				t20 := int64(load64(m.memory[uint32(v2):]))
				store64(m.memory[int64(uint32(v0))+4:], uint64(t20))
				t21 := int32(load32(m.memory[int64(uint32(v2))+8:]))
				store32(m.memory[int64(uint32(v0))+12:], uint32(t21))
				return
			}
		}
		m.fn563(v4, v0&i32(1), v2, v3)
		return
	}
}
func (m *Module) fn563(v0, v1, v2, v3 int32) {
	var v4, v5, v6 int32
	t0 := m.g0
	v4 = t0 - i32(16)
	m.g0 = v4
	v5 = i32(0)
	{
		t1 := int32(load32(m.memory[uint32(v0):]))
		v6 = t1
		switch v6 {
		case 2:
			v5 = i32(1)
			fallthrough
		case 1:
			if v5 == v1 {
				goto l3
			}
			fallthrough
		default:
			m.fn427(v0, v3)
			m.fn428(v0)
			store32(m.memory[int64(uint32(v0))+12:], uint32(i32(0)))
			t3 := v0
			p2 := i32(8)
			if v1 != 0 {
				p2 = i32(4)
			}
			store32(m.memory[int64(uint32(t3))+8:], uint32(p2))
			store32(m.memory[int64(uint32(v0))+4:], uint32(i32(0)))
			t5 := v0
			p4 := i32(1)
			if v1 != 0 {
				p4 = i32(2)
			}
			v6 = p4
			store32(m.memory[uint32(t5):], uint32(v6))
		}
	}
l3:
	t6 := int32(load32(m.memory[int64(uint32(v2))+8:]))
	v5 = t6
	t7 := int32(load32(m.memory[int64(uint32(v2))+4:]))
	v3 = t7
	{
		{
			{
				if v6 == i32(2) {
					goto l4
				}
				v1 = v5 * i32(28)
				v6 = v3 + i32(-28)
			l6:
				{
					if v1 == 0 {
						goto l5
					}
					v1 = v1 + i32(-28)
					v6 = v6 + i32(28)
					t8 := m.fn311(v6)
					if t8 != 0 {
						goto l6
					}
				}
				{
					t9 := int32(load32(m.memory[int64(uint32(v0))+12:]))
					v1 = t9
					t10 := int32(load32(m.memory[int64(uint32(v0))+4:]))
					if v1 != t10 {
						goto l7
					}
					m.fn315(v0 + i32(4))
				}
			l7:
				store32(m.memory[int64(uint32(v0))+12:], uint32(v1+i32(1)))
				t11 := int32(load32(m.memory[int64(uint32(v0))+8:]))
				v1 = t11 + v1<<5
				store32(m.memory[uint32(v1):], uint32(i32(-0x80000000)))
				t12 := int64(load64(m.memory[uint32(v2):]))
				store64(m.memory[int64(uint32(v1))+4:], uint64(t12))
				t13 := int32(load32(m.memory[int64(uint32(v2))+8:]))
				store32(m.memory[int64(uint32(v1))+12:], uint32(t13))
				goto l8
			}
		l4:
			store32(m.memory[int64(uint32(v4))+12:], uint32(i32(0)))
			store64(m.memory[int64(uint32(v4))+4:], uint64(i64(0x100000000)))
			m.fn460(v3, v5, v4+i32(4))
			{
				t14 := int32(load32(m.memory[int64(uint32(v0))+12:]))
				v1 = t14
				t15 := int32(load32(m.memory[int64(uint32(v0))+4:]))
				if v1 != t15 {
					goto l9
				}
				m.fn202(v0 + i32(4))
			}
		l9:
			store32(m.memory[int64(uint32(v0))+12:], uint32(v1+i32(1)))
			t16 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			v1 = t16 + v1*i32(12)
			t17 := int64(load64(m.memory[int64(uint32(v4))+4:]))
			store64(m.memory[uint32(v1):], uint64(t17))
			t18 := int32(load32(m.memory[int64(uint32(v4))+12:]))
			store32(m.memory[int64(uint32(v1))+8:], uint32(t18))
		}
	l5:
		if v5 == 0 {
			goto l10
		}
		v1 = v3
	l11:
		m.fn337(v1)
		v1 = v1 + i32(28)
		v5 = v5 + i32(-1)
		if v5 != 0 {
			goto l11
		}
	l10:
		t19 := int32(load32(m.memory[uint32(v2):]))
		v1 = t19
		if v1 == 0 {
			goto l8
		}
		t20 := int32(load32(m.memory[uint32(v3+i32(-4)):]))
		v5 = t20
		v6 = v5 & i32(-8)
		t21 := v6
		v5 = v5 & i32(3)
		p22 := i32(8)
		if v5 != 0 {
			p22 = i32(4)
		}
		v1 = v1 * i32(28)
		if uint32(t21) < uint32(p22+v1) {
			m.fn3(i32(1273840), i32(46), i32(1273888))
			panic("unreachable")
		}
		if v5 == 0 {
			goto l13
		}
		if uint32(v6) > uint32(v1+i32(39)) {
			m.fn3(i32(1273904), i32(46), i32(1273952))
			panic("unreachable")
		}
	l13:
		m.fn1(v3)
	}
l8:
	m.g0 = v4 + i32(16)
}
func (m *Module) fn564(v0, v1, v2, v3, v4 int32) {
	var v5, v6 int32
	var v7 int64
	var v8, v9 int32
	var v10, v11 int64
	var v12, v13, v14, v15 int32
	var v16 int64
	var v17, v18, v19 int32
	t0 := m.g0
	v5 = t0 - i32(112)
	m.g0 = v5
	t1 := int64(load64(m.memory[int64(uint32(v1))+48:]))
	t2 := int64(load64(m.memory[int64(uint32(v1))+56:]))
	t4 := v2
	p3 := i32(8)
	if uint32(v4) < uint32(i32(8)) {
		p3 = v4
	}
	v6 = p3
	t5 := m.fn92(t1, t2, t4, v6)
	v7 = t5
	{
		t6 := int32(load32(m.memory[int64(uint32(v1))+40:]))
		if t6 != 0 {
			goto l0
		}
		_ = m.fn91(v1+i32(32), v1+i32(48))
	}
l0:
	t8 := int32(load32(m.memory[int64(uint32(v1))+36:]))
	v8 = t8
	v9 = v8 & int32(v7)
	v10 = int64(uint64(v7) >> 25)
	v11 = v10 & i64(127) * i64(72340172838076673)
	t9 := int32(load32(m.memory[int64(uint32(v1))+32:]))
	v12 = t9
	v13 = i32(0)
	v14 = v2 & i32(0xffff)
	v15 = i32(0)
l61:
	{
		{
			{
				t10 := int64(load64(m.memory[uint32(v12+v9):]))
				v16 = t10
				v7 = v16 ^ v11
				v7 = (v7 ^ i64(-1)) & (v7 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
				if v7 == 0 {
					goto l1
				}
			l4:
				{
					t11 := v14
					v17 = v12 - (int32(uint32(int64(bits.TrailingZeros64(uint64(v7))))>>3)+v9)&v8<<3
					t12 := int32(load16(m.memory[uint32(v17+i32(-8)):]))
					if t11 != t12 {
						goto l2
					}
					v18 = i32(0)
					t13 := int32(load32(m.memory[uint32(v17+i32(-4)):]))
					if v6 == t13 {
						goto l3
					}
				}
			l2:
				v7 = (v7 + i64(-1)) & v7
				if !(v7 == 0) {
					goto l4
				}
			}
		l1:
			v7 = v16 & i64(-0x7f7f7f7f7f7f7f80)
			if v13 == i32(1) {
				goto l5
			}
			if v7 == 0 {
				v13 = i32(0)
				goto l8
			}
			v19 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v7))))>>3) + v9) & v8
		l5:
			if v7&(v16<<1) != i64(0) {
				goto l7
			}
			v13 = i32(1)
			goto l8
		l7:
			{
				t14 := int32(int8(m.memory[uint32(v12+v19)]))
				v9 = t14
				if v9 < i32(0) {
					goto l9
				}
				t15 := int64(load64(m.memory[uint32(v12):]))
				t16 := v12
				v19 = int32(uint32(int64(bits.TrailingZeros64(uint64(t15&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
				t17 := int32(m.memory[uint32(t16+v19)])
				v9 = t17
			}
		l9:
			t18 := v12 + v19
			v18 = int32(v10) & i32(127)
			m.memory[uint32(t18)] = byte(v18)
			m.memory[uint32(v12+(v19+i32(-8))&v8+i32(8))] = byte(v18)
			v18 = i32(1)
			t19 := int32(load32(m.memory[int64(uint32(v1))+40:]))
			store32(m.memory[int64(uint32(v1))+40:], uint32(t19-v9&i32(1)))
			t20 := int32(load32(m.memory[int64(uint32(v1))+44:]))
			store32(m.memory[int64(uint32(v1))+44:], uint32(t20+i32(1)))
			v12 = v12 - v19<<3
			store32(m.memory[uint32(v12+i32(-4)):], uint32(v6))
			store16(m.memory[uint32(v12+i32(-8)):], uint16(v2))
		}
	l3:
		t21 := int64(load64(m.memory[int64(uint32(v1))+16:]))
		t22 := int64(load64(m.memory[int64(uint32(v1))+24:]))
		t23 := int32(load32(m.memory[int64(uint32(v3))+504:]))
		v8 = t23
		t24 := m.fn94(t21, t22, v8)
		v7 = t24
		v10 = int64(uint64(v7) >> 25)
		v11 = v10 & i64(127) * i64(72340172838076673)
		t25 := int32(load32(m.memory[uint32(v1):]))
		v12 = t25
		v13 = i32(0)
		t26 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		v19 = t26
		t27 := v19
		v17 = int32(v7)
		v14 = t27 & v17
		v2 = v14
		{
		l14:
			{
				t28 := int64(load64(m.memory[uint32(v12+v2):]))
				v16 = t28
				v7 = v16 ^ v11
				v7 = (v7 ^ i64(-1)) & (v7 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
				if v7 == 0 {
					goto l10
				}
			l12:
				{
					v9 = v12 + (i32(0)-(int32(uint32(int64(bits.TrailingZeros64(uint64(v7))))>>3)+v2)&v19)*i32(96)
					t29 := int32(load32(m.memory[uint32(v9+i32(-96)):]))
					if t29 == v8 {
						goto l11
					}
					v7 = (v7 + i64(-1)) & v7
					if !(v7 == 0) {
						goto l12
					}
				}
			}
		l10:
			{
				if !(v16&(v16<<1)&i64(-0x7f7f7f7f7f7f7f80) == 0) {
					goto l13
				}
				t30 := v2
				v13 = v13 + i32(8)
				v2 = (t30 + v13) & v19
				goto l14
			}
		l13:
			{
				t31 := int32(load32(m.memory[int64(uint32(v1))+8:]))
				if t31 != 0 {
					goto l15
				}
				_ = m.fn101(v1, v1+i32(16))
				t33 := int32(load32(m.memory[int64(uint32(v1))+4:]))
				v19 = t33
				v14 = v19 & v17
				t34 := int32(load32(m.memory[uint32(v1):]))
				v12 = t34
			}
		l15:
			memory_zero(m.memory, uint32(v5+i32(12)), uint32(i32(81)))
			{
				t35 := int64(load64(m.memory[uint32(v12+v14):]))
				v7 = t35 & i64(-0x7f7f7f7f7f7f7f80)
				if v7 != i64(0) {
					goto l16
				}
				v9 = i32(8)
			l17:
				{
					v2 = v14 + v9
					v9 = v9 + i32(8)
					t36 := v12
					v14 = v2 & v19
					t37 := int64(load64(m.memory[uint32(t36+v14):]))
					v7 = t37 & i64(-0x7f7f7f7f7f7f7f80)
					if v7 == 0 {
						goto l17
					}
				}
			}
		l16:
			{
				t38 := v12
				v9 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v7))))>>3) + v14) & v19
				t39 := int32(int8(m.memory[uint32(t38+v9)]))
				v2 = t39
				if v2 < i32(0) {
					goto l18
				}
				t40 := int64(load64(m.memory[uint32(v12):]))
				t41 := v12
				v9 = int32(uint32(int64(bits.TrailingZeros64(uint64(t40&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
				t42 := int32(m.memory[uint32(t41+v9)])
				v2 = t42
			}
		l18:
			t43 := v12 + v9
			v14 = int32(v10) & i32(127)
			m.memory[uint32(t43)] = byte(v14)
			m.memory[uint32(v12+(v9+i32(-8))&v19+i32(8))] = byte(v14)
			t44 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			store32(m.memory[int64(uint32(v1))+8:], uint32(t44-v2&i32(1)))
			v9 = v12 + (i32(0)-v9)*i32(96)
			store32(m.memory[uint32(v9+i32(-96)):], uint32(v8))
			memory_copy(m.memory, uint32(v9+i32(-92)), uint32(v5+i32(8)), uint32(i32(92)))
			t45 := int32(load32(m.memory[int64(uint32(v1))+12:]))
			store32(m.memory[int64(uint32(v1))+12:], uint32(t45+i32(1)))
		}
	l11:
		v8 = v9 + i32(-88)
		{
			{
				t46 := v18
				v1 = v3 + v6<<4
				t47 := int32(load32(m.memory[uint32(v1):]))
				if t46&t47 != i32(1) {
					goto l19
				}
				t48 := int64(load64(m.memory[int64(uint32(v1))+8:]))
				v7 = t48
				goto l20
			}
		l19:
			{
				t49 := int32(m.memory[uint32(v9+v6+i32(-16))])
				if t49 != 0 {
					goto l21
				}
				t50 := int64(load64(m.memory[int64(uint32(v3+v6*i32(40)))+168:]))
				v7 = t50
				goto l20
			}
		l21:
			t51 := int64(load64(m.memory[uint32(v8+v6<<3):]))
			v7 = t51 + i64(1)
			p52 := v7
			if v7 == 0 {
				p52 = i64(-1)
			}
			v7 = p52
		}
	l20:
		v19 = v9 + i32(-16)
		m.memory[uint32(v19+v6)] = byte(i32(1))
		store64(m.memory[uint32(v8+v6<<3):], uint64(v7))
		if uint32(v4) > uint32(i32(7)) {
			goto l22
		}
		v4 = v6 + i32(2)
		v12 = v6 + i32(1)
		v1 = v12 * i32(40)
		v9 = v1 + i32(184)
		{
			v1 = v3 + v1
			t53 := int32(load32(m.memory[int64(uint32(v1))+144:]))
			if t53 != i32(1) {
				goto l23
			}
			t54 := int32(load32(m.memory[int64(uint32(v1))+148:]))
			if uint32(v6) >= uint32(t54) {
				goto l24
			}
		}
	l23:
		v1 = i32(0)
		goto l28
	}
l24:
	v1 = i32(1)
l28:
	switch v1 {
	case 0:
		m.memory[uint32(v19+v12)] = byte(i32(0))
		v1 = i32(1)
		goto l28
	default:
	l30:
		{
			v12 = v4
			v1 = v9
			if v1 == i32(504) {
				goto l22
			}
			v2 = v3 + v1
			t55 := int32(load32(m.memory[uint32(v2):]))
			if t55 == 0 {
				goto l29
			}
			v9 = v1 + i32(40)
			v4 = v12 + i32(1)
			t56 := int32(load32(m.memory[uint32(v2+i32(4)):]))
			if uint32(v6) >= uint32(t56) {
				goto l30
			}
		}
		v9 = v1 + i32(40)
		v4 = v12 + i32(1)
		goto l31
	l29:
		v9 = v1 + i32(40)
		v4 = v12 + i32(1)
	l31:
		if uint32(v12) < uint32(i32(9)) {
			v1 = i32(0)
			goto l28
		}
		m.fn33(v12, i32(9), i32(1073212))
		panic("unreachable")
	}
l22:
	v12 = i32(-1)
	{
		v1 = v3 + i32(144) + v6*i32(40)
		t57 := int32(m.memory[int64(uint32(v1))+32])
		v17 = t57
		if v17 == i32(255) {
			goto l53
		}
		t58 := int32(load32(m.memory[int64(uint32(v1))+16:]))
		v4 = t58
		if v4 == 0 {
			goto l53
		}
		v12 = i32(0)
		store32(m.memory[int64(uint32(v5))+108:], uint32(i32(0)))
		store64(m.memory[int64(uint32(v5))+100:], uint64(i64(0x100000000)))
		v9 = v4 * i32(12)
		t59 := int32(m.memory[int64(uint32(v1))+20])
		v18 = t59
		t60 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		v1 = t60
		v14 = i32(1)
	l46:
		{
			{
				t61 := int32(load32(m.memory[uint32(v1):]))
				if t61 != i32(-1) {
					t76 := int32(load32(m.memory[uint32(v1+i32(4)):]))
					v2 = t76
					{
						t77 := int32(load32(m.memory[uint32(v1+i32(8)):]))
						v4 = t77
						t78 := int32(load32(m.memory[int64(uint32(v5))+100:]))
						if uint32(v4) <= uint32(t78-v12) {
							goto l38
						}
						m.fn197(v5+i32(100), v12, v4, i32(1), i32(1))
						t79 := int32(load32(m.memory[int64(uint32(v5))+104:]))
						v14 = t79
						t80 := int32(load32(m.memory[int64(uint32(v5))+108:]))
						v12 = t80
						goto l39
					}
				l38:
					if v4 == 0 {
						goto l40
					}
				l39:
					if v4 == 0 {
						goto l40
					}
					memory_copy(m.memory, uint32(v14+v12), uint32(v2), uint32(v4))
				l40:
					t81 := v5
					v12 = v12 + v4
					store32(m.memory[int64(uint32(t81))+108:], uint32(v12))
					goto l41
				}
				t62 := int32(m.memory[uint32(v1+i32(4))])
				v4 = t62
				p63 := i32(8)
				if uint32(v4) < uint32(i32(8)) {
					p63 = v4
				}
				v4 = p63
				v2 = i32(1)
				{
					if v18&i32(1) != 0 {
						goto l35
					}
					t64 := int32(m.memory[int64(uint32(v3+v4*i32(40)))+176])
					v2 = t64
					p65 := v2
					if v2 == i32(255) {
						p65 = i32(1)
					}
					v2 = p65
				}
			l35:
				t66 := int32(m.memory[uint32(v19+v4)])
				t68 := v5 + i32(8)
				t69 := v2
				p67 := v3 + v4*i32(40) + i32(168)
				if t66 != 0 {
					p67 = v8 + v4<<3
				}
				t70 := int64(load64(m.memory[uint32(p67):]))
				m.fn308(t68, t69, t70)
				t71 := int32(load32(m.memory[int64(uint32(v5))+12:]))
				v2 = t71
				t72 := int32(load32(m.memory[int64(uint32(v5))+16:]))
				v4 = t72
				t73 := int32(load32(m.memory[int64(uint32(v5))+100:]))
				if uint32(v4) <= uint32(t73-v12) {
					goto l36
				}
				m.fn197(v5+i32(100), v12, v4, i32(1), i32(1))
				t74 := int32(load32(m.memory[int64(uint32(v5))+104:]))
				v14 = t74
				t75 := int32(load32(m.memory[int64(uint32(v5))+108:]))
				v12 = t75
				goto l37
			}
		l36:
			if v4 == 0 {
				goto l42
			}
		l37:
			if v4 == 0 {
				goto l42
			}
			memory_copy(m.memory, uint32(v14+v12), uint32(v2), uint32(v4))
		l42:
			t82 := v5
			v12 = v12 + v4
			store32(m.memory[int64(uint32(t82))+108:], uint32(v12))
			t83 := int32(load32(m.memory[int64(uint32(v5))+8:]))
			v4 = t83
			if v4 == 0 {
				goto l41
			}
			{
				t84 := int32(load32(m.memory[uint32(v2+i32(-4)):]))
				v6 = t84
				v13 = v6 & i32(-8)
				t85 := v13
				v6 = v6 & i32(3)
				p86 := i32(8)
				if v6 != 0 {
					p86 = i32(4)
				}
				if uint32(t85) < uint32(p86+v4) {
					m.fn3(i32(1273840), i32(46), i32(1273888))
					panic("unreachable")
				}
				if v6 == 0 {
					goto l44
				}
				if uint32(v13) > uint32(v4+i32(39)) {
					m.fn3(i32(1273904), i32(46), i32(1273952))
					panic("unreachable")
				}
			l44:
				m.fn1(v2)
				goto l41
			}
		}
	l41:
		v1 = v1 + i32(12)
		v9 = v9 + i32(-12)
		if v9 != 0 {
			goto l46
		}
		m.fn307(v5+i32(8), v17, v7)
		{
			{
				t87 := int32(load32(m.memory[int64(uint32(v5))+16:]))
				if v12 != t87 {
					goto l47
				}
				t88 := int32(load32(m.memory[int64(uint32(v5))+12:]))
				t89 := v14
				v1 = t88
				t90 := m.fn974(t89, v1, v12)
				if t90 == 0 {
					goto l48
				}
			}
		l47:
			{
				t91 := int32(load32(m.memory[int64(uint32(v5))+8:]))
				v1 = t91
				if v1 == 0 {
					goto l49
				}
				t92 := int32(load32(m.memory[int64(uint32(v5))+12:]))
				v4 = t92
				t93 := int32(load32(m.memory[uint32(v4+i32(-4)):]))
				v12 = t93
				v9 = v12 & i32(-8)
				t94 := v9
				v12 = v12 & i32(3)
				p95 := i32(8)
				if v12 != 0 {
					p95 = i32(4)
				}
				if uint32(t94) < uint32(p95+v1) {
					m.fn3(i32(1273840), i32(46), i32(1273888))
					panic("unreachable")
				}
				if v12 == 0 {
					goto l51
				}
				if uint32(v9) > uint32(v1+i32(39)) {
					m.fn3(i32(1273904), i32(46), i32(1273952))
					panic("unreachable")
				}
			l51:
				m.fn1(v4)
			}
		l49:
			t96 := int64(load64(m.memory[int64(uint32(v5))+104:]))
			v11 = t96
			t97 := int32(load32(m.memory[int64(uint32(v5))+100:]))
			v12 = t97
			goto l53
		}
	l48:
		{
			t98 := int32(load32(m.memory[int64(uint32(v5))+8:]))
			v12 = t98
			if v12 == 0 {
				goto l54
			}
			t99 := int32(load32(m.memory[uint32(v1+i32(-4)):]))
			v4 = t99
			v9 = v4 & i32(-8)
			t100 := v9
			v4 = v4 & i32(3)
			p101 := i32(8)
			if v4 != 0 {
				p101 = i32(4)
			}
			if uint32(t100) < uint32(p101+v12) {
				m.fn3(i32(1273840), i32(46), i32(1273888))
				panic("unreachable")
			}
			if v4 == 0 {
				goto l56
			}
			if uint32(v9) > uint32(v12+i32(39)) {
				m.fn3(i32(1273904), i32(46), i32(1273952))
				panic("unreachable")
			}
		l56:
			m.fn1(v1)
		}
	l54:
		v12 = i32(-1)
		t102 := int32(load32(m.memory[int64(uint32(v5))+100:]))
		v1 = t102
		if v1 == 0 {
			goto l53
		}
		t103 := int32(load32(m.memory[uint32(v14+i32(-4)):]))
		v4 = t103
		v9 = v4 & i32(-8)
		t104 := v9
		v4 = v4 & i32(3)
		p105 := i32(8)
		if v4 != 0 {
			p105 = i32(4)
		}
		if uint32(t104) < uint32(p105+v1) {
			m.fn3(i32(1273840), i32(46), i32(1273888))
			panic("unreachable")
		}
		if v4 == 0 {
			goto l59
		}
		if uint32(v9) > uint32(v1+i32(39)) {
			m.fn3(i32(1273904), i32(46), i32(1273952))
			panic("unreachable")
		}
	l59:
		m.fn1(v14)
	}
l53:
	store64(m.memory[int64(uint32(v0))+12:], uint64(v11))
	store32(m.memory[int64(uint32(v0))+8:], uint32(v12))
	store64(m.memory[uint32(v0):], uint64(v7))
	m.g0 = v5 + i32(112)
	return
l8:
	v15 = v15 + i32(8)
	v9 = (v15 + v9) & v8
	goto l61
}
func (m *Module) fn565(v0 int32) {
	var v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11 int32
	t0 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	v1 = t0
	{
		t1 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		v2 = t1
		if v2 == 0 {
			goto l0
		}
		v3 = i32(0)
	l21:
		{
			v4 = v1 + v3*i32(40)
			t2 := int32(load32(m.memory[int64(uint32(v4))+4:]))
			v5 = t2
			{
				t3 := int32(load32(m.memory[int64(uint32(v4))+8:]))
				v6 = t3
				if v6 == 0 {
					goto l1
				}
				v7 = i32(0)
			l8:
				{
					v8 = v5 + v7*i32(12)
					t4 := int32(load32(m.memory[int64(uint32(v8))+4:]))
					v9 = t4
					{
						t5 := int32(load32(m.memory[int64(uint32(v8))+8:]))
						v10 = t5
						if v10 == 0 {
							goto l2
						}
						v11 = v9
					l3:
						m.fn335(v11)
						v11 = v11 + i32(32)
						v10 = v10 + i32(-1)
						if v10 != 0 {
							goto l3
						}
					}
				l2:
					{
						t6 := int32(load32(m.memory[uint32(v8):]))
						v11 = t6
						if v11 == 0 {
							goto l4
						}
						t7 := int32(load32(m.memory[uint32(v9+i32(-4)):]))
						v10 = t7
						v8 = v10 & i32(-8)
						t8 := v8
						v10 = v10 & i32(3)
						p9 := i32(8)
						if v10 != 0 {
							p9 = i32(4)
						}
						v11 = v11 << 5
						if uint32(t8) < uint32(p9|v11) {
							m.fn3(i32(1273840), i32(46), i32(1273888))
							panic("unreachable")
						}
						if v10 == 0 {
							goto l6
						}
						if uint32(v8) > uint32(v11+i32(39)) {
							m.fn3(i32(1273904), i32(46), i32(1273952))
							panic("unreachable")
						}
					l6:
						m.fn1(v9)
					}
				l4:
					v7 = v7 + i32(1)
					if v7 != v6 {
						goto l8
					}
				}
			}
		l1:
			{
				t10 := int32(load32(m.memory[uint32(v4):]))
				v11 = t10
				if v11 == 0 {
					goto l9
				}
				t11 := int32(load32(m.memory[uint32(v5+i32(-4)):]))
				v10 = t11
				v7 = v10 & i32(-8)
				t12 := v7
				v10 = v10 & i32(3)
				p13 := i32(8)
				if v10 != 0 {
					p13 = i32(4)
				}
				v11 = v11 * i32(12)
				if uint32(t12) < uint32(p13+v11) {
					m.fn3(i32(1273840), i32(46), i32(1273888))
					panic("unreachable")
				}
				if v10 == 0 {
					goto l11
				}
				if uint32(v7) > uint32(v11+i32(39)) {
					m.fn3(i32(1273904), i32(46), i32(1273952))
					panic("unreachable")
				}
			l11:
				m.fn1(v5)
			}
		l9:
			{
				t14 := int32(load32(m.memory[int64(uint32(v4))+12:]))
				v11 = t14
				if v11 == i32(-1) {
					goto l13
				}
				{
					if v11 == 0 {
						goto l14
					}
					t15 := int32(load32(m.memory[int64(uint32(v4))+16:]))
					v7 = t15
					t16 := int32(load32(m.memory[uint32(v7+i32(-4)):]))
					v10 = t16
					v8 = v10 & i32(-8)
					t17 := v8
					v10 = v10 & i32(3)
					p18 := i32(8)
					if v10 != 0 {
						p18 = i32(4)
					}
					v11 = v11 << 1
					if uint32(t17) < uint32(p18+v11) {
						m.fn3(i32(1273840), i32(46), i32(1273888))
						panic("unreachable")
					}
					if v10 == 0 {
						goto l16
					}
					if uint32(v8) > uint32(v11+i32(39)) {
						m.fn3(i32(1273904), i32(46), i32(1273952))
						panic("unreachable")
					}
				l16:
					m.fn1(v7)
				}
			l14:
				t19 := int32(load32(m.memory[int64(uint32(v4))+24:]))
				v11 = t19
				if v11 == 0 {
					goto l13
				}
				t20 := int32(load32(m.memory[int64(uint32(v4))+28:]))
				v7 = t20
				t21 := int32(load32(m.memory[uint32(v7+i32(-4)):]))
				v10 = t21
				v8 = v10 & i32(-8)
				t22 := v8
				v10 = v10 & i32(3)
				p23 := i32(8)
				if v10 != 0 {
					p23 = i32(4)
				}
				v11 = v11 << 2
				if uint32(t22) < uint32(p23+v11) {
					m.fn3(i32(1273840), i32(46), i32(1273888))
					panic("unreachable")
				}
				if v10 == 0 {
					goto l19
				}
				if uint32(v8) > uint32(v11+i32(39)) {
					m.fn3(i32(1273904), i32(46), i32(1273952))
					panic("unreachable")
				}
			l19:
				m.fn1(v7)
			}
		l13:
			v3 = v3 + i32(1)
			if v3 != v2 {
				goto l21
			}
		}
	}
l0:
	{
		t24 := int32(load32(m.memory[uint32(v0):]))
		v11 = t24
		if v11 == 0 {
			return
		}
		t25 := int32(load32(m.memory[uint32(v1+i32(-4)):]))
		v10 = t25
		v7 = v10 & i32(-8)
		t26 := v7
		v10 = v10 & i32(3)
		p27 := i32(8)
		if v10 != 0 {
			p27 = i32(4)
		}
		v11 = v11 * i32(40)
		if uint32(t26) < uint32(p27+v11) {
			m.fn3(i32(1273840), i32(46), i32(1273888))
			panic("unreachable")
		}
		if v10 == 0 {
			goto l24
		}
		if uint32(v7) > uint32(v11+i32(39)) {
			m.fn3(i32(1273904), i32(46), i32(1273952))
			panic("unreachable")
		}
	l24:
		m.fn1(v1)
	}
}
func (m *Module) fn566(v0, v1 int32) {
	var v2, v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14 int32
	var v15, v16 int64
	var v17, v18, v19, v20, v21, v22, v23, v24, v25, v26, v27, v28, v29, v30, v31 int32
	var v32, v33, v34, v35, v36, v37, v38, v39, v40 int64
	t0 := m.g0
	v2 = t0 - i32(208)
	m.g0 = v2
	t1 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	v3 = t1
	{
		{
			t2 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			v4 = t2
			if v4 == 0 {
				goto l0
			}
			t3 := v3
			v5 = v4 << 4
			v6 = t3 + v5
			v4 = v3 + i32(12)
			v7 = int32(uint32(v5+i32(-16))>>4) + i32(1)
			v8 = i32(0)
		l2:
			{
				t4 := int32(m.memory[uint32(v4)])
				if t4 != i32(1) {
					goto l1
				}
				v4 = v4 + i32(16)
				v8 = v8 + i32(1)
				v5 = v5 + i32(-16)
				if v5 != 0 {
					goto l2
				}
			}
			v8 = v7
		l1:
			t5 := int32(load32(m.memory[uint32(v1):]))
			v9 = t5
			v7 = v3
			v10 = v3
		l7:
			{
				t6 := int32(load32(m.memory[uint32(v10):]))
				v11 = t6
				t7 := int32(load32(m.memory[int64(uint32(v10))+4:]))
				v12 = t7
				v13 = v12
				{
					t8 := int32(load32(m.memory[int64(uint32(v10))+8:]))
					v4 = t8
					if v4 == 0 {
						goto l3
					}
					t9 := v12
					v4 = v4 << 5
					v13 = t9 + v4
					v5 = v4 + i32(-32)
					v1 = int32(uint32(v5)>>5) + i32(1)
					v14 = v1 & i32(3)
					v15 = i64(-0x8000000000000000)
					v4 = v12
					if uint32(v5) < uint32(i32(96)) {
						goto l4
					}
					v5 = v1 & i32(0xffffffc)
					v15 = i64(-0x8000000000000000)
					v4 = v12
				l5:
					{
						v1 = v4 + i32(16)
						t10 := int64(load64(m.memory[uint32(v1):]))
						t11 := v1
						v16 = t10
						p12 := v15 + i64(1)
						if v16 > v15 {
							p12 = v16
						}
						v15 = p12
						store64(m.memory[uint32(t11):], uint64(v15))
						v1 = v4 + i32(48)
						t13 := int64(load64(m.memory[uint32(v1):]))
						t14 := v1
						v16 = t13
						p15 := v15 + i64(1)
						if v16 > v15 {
							p15 = v16
						}
						v15 = p15
						store64(m.memory[uint32(t14):], uint64(v15))
						v1 = v4 + i32(80)
						t16 := int64(load64(m.memory[uint32(v1):]))
						t17 := v1
						v16 = t16
						p18 := v15 + i64(1)
						if v16 > v15 {
							p18 = v16
						}
						v15 = p18
						store64(m.memory[uint32(t17):], uint64(v15))
						v1 = v4 + i32(112)
						t19 := int64(load64(m.memory[uint32(v1):]))
						t20 := v1
						v16 = t19
						p21 := v15 + i64(1)
						if v16 > v15 {
							p21 = v16
						}
						v15 = p21
						store64(m.memory[uint32(t20):], uint64(v15))
						v4 = v4 + i32(128)
						v5 = v5 + i32(-4)
						if v5 != 0 {
							goto l5
						}
					}
					if v14 == 0 {
						goto l3
					}
				l4:
					v5 = v14 << 5
					v4 = v4 + i32(16)
				l6:
					{
						t22 := int64(load64(m.memory[uint32(v4):]))
						t23 := v4
						v16 = t22
						p24 := v15 + i64(1)
						if v16 > v15 {
							p24 = v16
						}
						v15 = p24
						store64(m.memory[uint32(t23):], uint64(v15))
						v4 = v4 + i32(32)
						v5 = v5 + i32(-32)
						if v5 != 0 {
							goto l6
						}
					}
				}
			l3:
				store32(m.memory[int64(uint32(v7))+4:], uint32(v12))
				store32(m.memory[uint32(v7):], uint32(v11))
				store32(m.memory[int64(uint32(v7))+8:], uint32(int32(uint32(v13-v12)>>5)))
				v7 = v7 + i32(12)
				v10 = v10 + i32(16)
				if v10 != v6 {
					goto l7
				}
				goto l8
			}
		}
	l0:
		t25 := int32(load32(m.memory[uint32(v1):]))
		v9 = t25
		v8 = i32(0)
		v7 = v3
	}
l8:
	v17 = v9 << 4
	t26 := int32(uint32(v17) % uint32(i32(12)))
	t27 := v17
	v4 = t26
	v18 = t27 - v4
	v19 = v3
	{
		{
			if v4 == 0 {
				goto l9
			}
			t28 := m.fn22(v3, v17, i32(4), v18)
			v19 = t28
			if v19 == 0 {
				m.fn24(i32(4), v18)
				panic("unreachable")
			}
		}
	l9:
		t29 := v19
		v10 = v7 - v3
		v20 = t29 + v10
		v5 = i32(0)
		v4 = i32(0)
		{
		l15:
			if v5 == 0 {
				goto l11
			}
			if v5 != v1 {
				goto l12
			}
		l11:
			if v10 != v4 {
				v1 = v19 + v4
				t30 := int32(load32(m.memory[uint32(v1+i32(4)):]))
				v5 = t30
				t31 := int32(load32(m.memory[uint32(v1+i32(8)):]))
				v1 = v5 + t31<<5
				v4 = v4 + i32(12)
				goto l15
			}
			v12 = i32(8)
			v13 = i32(0)
			v10 = i32(0)
			goto l14
		l12:
			t32 := int64(load64(m.memory[uint32(v5+i32(16)):]))
			v15 = t32
			t33 := v1
			v5 = v5 + i32(32)
			v10 = int32(uint32(t33-v5) >> 5)
			p34 := i32(3)
			if uint32(v10) > uint32(i32(3)) {
				p34 = v10
			}
			v10 = p34 + i32(1)
			v12 = v10 << 3
			t35 := m.fn5(v12)
			v13 = t35
			if v13 == 0 {
				m.fn10(i32(8), v12)
				panic("unreachable")
			}
			v4 = v19 + v4
			store64(m.memory[uint32(v13):], uint64(v15))
			store32(m.memory[int64(uint32(v2))+160:], uint32(i32(1)))
			store32(m.memory[int64(uint32(v2))+156:], uint32(v13))
			store32(m.memory[int64(uint32(v2))+152:], uint32(v10))
			v10 = i32(1)
		l23:
			{
				{
					if v5 == v1 {
						goto l17
					}
					v12 = v5
					goto l18
				l17:
					if v4 == v20 {
						goto l19
					}
				l21:
					{
						v5 = v4 + i32(12)
						t36 := int32(load32(m.memory[uint32(v4+i32(8)):]))
						v1 = t36
						if v1 != 0 {
							goto l20
						}
						v4 = v5
						if v5 == v20 {
							goto l19
						}
						goto l21
					}
				l20:
					t37 := int32(load32(m.memory[uint32(v4+i32(4)):]))
					v12 = t37
					v1 = v12 + v1<<5
					v4 = v5
				}
			l18:
				v5 = v12 + i32(32)
				t38 := int64(load64(m.memory[uint32(v12+i32(16)):]))
				v15 = t38
				{
					t39 := int32(load32(m.memory[int64(uint32(v2))+152:]))
					if v10 != t39 {
						goto l22
					}
					m.fn197(v2+i32(152), v10, int32(uint32(v1-v5)>>5)+i32(1), i32(8), i32(8))
					t40 := int32(load32(m.memory[int64(uint32(v2))+156:]))
					v13 = t40
				}
			l22:
				store64(m.memory[uint32(v13+v10<<3):], uint64(v15))
				t41 := v2
				v10 = v10 + i32(1)
				store32(m.memory[int64(uint32(t41))+160:], uint32(v10))
				goto l23
			}
		l19:
			t42 := int32(load32(m.memory[int64(uint32(v2))+156:]))
			v12 = t42
			t43 := int32(load32(m.memory[int64(uint32(v2))+152:]))
			v13 = t43
			if uint32(v10) < uint32(i32(2)) {
				goto l14
			}
			if uint32(v10) < uint32(i32(21)) {
				goto l24
			}
			m.fn133(v12, v10)
			goto l14
		l24:
			v1 = v12 + i32(8)
			v14 = v12 + v10<<3
			v11 = i32(0)
		l29:
			{
				t44 := int64(load64(m.memory[uint32(v1):]))
				v16 = t44
				t45 := int64(load64(m.memory[uint32(v1+i32(-8)):]))
				t46 := v16
				v15 = t45
				if t46 >= v15 {
					goto l25
				}
				v4 = v11
			l28:
				{
					store64(m.memory[uint32(v12+v4+i32(8)):], uint64(v15))
					if v4 != 0 {
						goto l26
					}
					v4 = v12
					goto l27
				l26:
					t47 := v16
					v4 = v4 + i32(-8)
					v5 = v4 + v12
					t48 := int64(load64(m.memory[uint32(v5):]))
					v15 = t48
					if t47 < v15 {
						goto l28
					}
				}
				v4 = v5 + i32(8)
			l27:
				store64(m.memory[uint32(v4):], uint64(v16))
			}
		l25:
			v11 = v11 + i32(8)
			v1 = v1 + i32(8)
			if v1 != v14 {
				goto l29
			}
		}
	l14:
		v21 = i32(0)
		store32(m.memory[int64(uint32(v2))+16:], uint32(i32(0)))
		store64(m.memory[int64(uint32(v2))+8:], uint64(i64(0x800000000)))
		if v10 == 0 {
			goto l30
		}
		v5 = v10 << 3
		v21 = i32(0)
		v1 = i32(8)
		v4 = v12
	l34:
		{
			t49 := int64(load64(m.memory[uint32(v4):]))
			v15 = t49
			{
				{
					if v21 == 0 {
						goto l31
					}
					t50 := int64(load64(m.memory[uint32(v1+v21<<3+i32(-8)):]))
					if v15-t50 < i64(11) {
						goto l32
					}
				}
			l31:
				{
					t51 := int32(load32(m.memory[int64(uint32(v2))+8:]))
					if v21 != t51 {
						goto l33
					}
					m.fn332(v2 + i32(8))
					t52 := int32(load32(m.memory[int64(uint32(v2))+12:]))
					v1 = t52
				}
			l33:
				store64(m.memory[uint32(v1+v21<<3):], uint64(v15))
				t53 := v2
				v21 = v21 + i32(1)
				store32(m.memory[int64(uint32(t53))+16:], uint32(v21))
			}
		l32:
			v4 = v4 + i32(8)
			v5 = v5 + i32(-8)
			if v5 != 0 {
				goto l34
			}
		}
	l30:
		{
			if v13 == 0 {
				goto l35
			}
			t54 := int32(load32(m.memory[uint32(v12+i32(-4)):]))
			v4 = t54
			v5 = v4 & i32(-8)
			t55 := v5
			v4 = v4 & i32(3)
			p56 := i32(8)
			if v4 != 0 {
				p56 = i32(4)
			}
			v1 = v13 << 3
			if uint32(t55) < uint32(p56+v1) {
				m.fn3(i32(1273840), i32(46), i32(1273888))
				panic("unreachable")
			}
			if v4 == 0 {
				goto l37
			}
			if uint32(v5) > uint32(v1+i32(39)) {
				m.fn3(i32(1273904), i32(46), i32(1273952))
				panic("unreachable")
			}
		l37:
			m.fn1(v12)
		}
	l35:
		v22 = i32(0)
		store32(m.memory[int64(uint32(v2))+28:], uint32(i32(0)))
		store64(m.memory[int64(uint32(v2))+20:], uint64(i64(0x400000000)))
		if v7 != v3 {
			v22 = i32(0)
			t57 := int32(load32(m.memory[int64(uint32(v2))+12:]))
			v12 = t57
			v23 = i32(4)
			v4 = v19
		l77:
			{
				v24 = v4 + i32(12)
				t58 := int32(load32(m.memory[uint32(v4):]))
				v25 = t58
				if v25 == i32(-1) {
					goto l40
				}
				t59 := int32(load32(m.memory[int64(uint32(v4))+4:]))
				v26 = t59
				v27 = i32(0)
				store32(m.memory[int64(uint32(v2))+40:], uint32(i32(0)))
				store64(m.memory[int64(uint32(v2))+32:], uint64(i64(0x400000000)))
				t60 := int32(load32(m.memory[int64(uint32(v4))+8:]))
				v6 = v26 + t60<<5
				v14 = i32(-2)
				v28 = i32(4)
				v3 = i32(0)
				v11 = v26
			l63:
				{
					{
						if v14 == i32(-2) {
							goto l41
						}
						v29 = v13
						v4 = v10
						v22 = v9
						v23 = v30
						v1 = v31
						v15 = v16
						v5 = v11
						goto l42
					l41:
						if v11 != v6 {
							goto l43
						}
						v5 = v6
						goto l44
					l43:
						v5 = v11 + i32(32)
						t61 := int64(load64(m.memory[int64(uint32(v11))+4:]))
						v32 = t61
						v4 = int32(int64(uint64(v32) >> 32))
						t62 := int32(m.memory[int64(uint32(v11))+24])
						v1 = t62
						t63 := int32(m.memory[int64(uint32(v11))+27])
						v22 = t63
						t64 := int32(m.memory[int64(uint32(v11))+26])
						v23 = t64
						t65 := int64(load64(m.memory[int64(uint32(v11))+16:]))
						v15 = t65
						t66 := int32(load32(m.memory[uint32(v11):]))
						v14 = t66
						v29 = int32(v32)
					}
				l42:
					if v14 == i32(-1) {
						goto l44
					}
					store32(m.memory[int64(uint32(v2))+44:], uint32(v14))
					store64(m.memory[int64(uint32(v2))+48:], uint64(int64(uint32(v4))<<32|int64(uint32(v29))))
					v14 = i32(-2)
					if v1&i32(1) != 0 {
						goto l45
					}
					v11 = v5
					goto l46
				l45:
					v14 = i32(-1)
					if v5 != v6 {
						goto l57
					}
					v11 = v6
					goto l46
				l57:
					{
						v11 = v5 + i32(32)
						t67 := int32(m.memory[uint32(v5+i32(27))])
						v9 = t67
						t68 := int32(m.memory[uint32(v5+i32(26))])
						v30 = t68
						t69 := int32(m.memory[uint32(v5+i32(24))])
						v31 = t69
						t70 := int64(load64(m.memory[uint32(v5+i32(16)):]))
						v16 = t70
						t71 := int32(load32(m.memory[uint32(v5+i32(8)):]))
						v10 = t71
						t72 := int32(load32(m.memory[uint32(v5+i32(4)):]))
						v13 = t72
						t73 := int32(load32(m.memory[uint32(v5):]))
						v1 = t73
						if v1 == i32(-1) {
							goto l46
						}
						{
							t74 := int32(m.memory[uint32(v5+i32(25))])
							if t74&i32(1) != 0 {
								goto l48
							}
							v14 = v1
							goto l46
						}
					l48:
						{
							t75 := int32(load32(m.memory[int64(uint32(v2))+44:]))
							if uint32(v10) <= uint32(t75-v4) {
								goto l49
							}
							m.fn197(v2+i32(44), v4, v10, i32(8), i32(32))
							t76 := int32(load32(m.memory[int64(uint32(v2))+48:]))
							v29 = t76
							t77 := int32(load32(m.memory[int64(uint32(v2))+52:]))
							v4 = t77
							goto l50
						}
					l49:
						if v10 == 0 {
							goto l51
						}
					l50:
						v7 = v10 << 5
						if v7 == 0 {
							goto l51
						}
						memory_copy(m.memory, uint32(v29+v4<<5), uint32(v13), uint32(v7))
					l51:
						t78 := v2
						v4 = v4 + v10
						store32(m.memory[int64(uint32(t78))+52:], uint32(v4))
						{
							if v1 == 0 {
								goto l52
							}
							t79 := int32(load32(m.memory[uint32(v13+i32(-4)):]))
							v7 = t79
							v11 = v7 & i32(-8)
							t80 := v11
							v7 = v7 & i32(3)
							p81 := i32(8)
							if v7 != 0 {
								p81 = i32(4)
							}
							v1 = v1 << 5
							if uint32(t80) < uint32(p81|v1) {
								goto l53
							}
							if v7 == 0 {
								goto l54
							}
							if uint32(v11) > uint32(v1+i32(39)) {
								m.fn3(i32(1273904), i32(46), i32(1273952))
								panic("unreachable")
							}
						l54:
							m.fn1(v13)
						}
					l52:
						v15 = v16
						v5 = v5 + i32(32)
						if v5 == v6 {
							goto l56
						}
						goto l57
					l53:
					}
					m.fn3(i32(1273840), i32(46), i32(1273888))
					panic("unreachable")
				l56:
					v11 = v6
					v15 = v16
				l46:
					{
						if v21 != 0 {
							goto l58
						}
						v4 = i32(1)
						goto l59
					l58:
						v15 = v15 + i64(-10)
						v4 = i32(0)
						if v21 == i32(1) {
							goto l60
						}
						v4 = i32(0)
						v5 = v21
					l61:
						{
							v1 = int32(uint32(v5) >> 1)
							v7 = v1 + v4
							t82 := int64(load64(m.memory[uint32(v12+v7<<3):]))
							p83 := v4
							if t82 < v15 {
								p83 = v7
							}
							v4 = p83
							v5 = v5 - v1
							if uint32(v5) > uint32(i32(1)) {
								goto l61
							}
						}
					l60:
						t84 := int64(load64(m.memory[uint32(v12+v4<<3):]))
						t85 := v4
						var p86 int32
						if t84 < v15 {
							p86 = 1
						}
						v4 = t85 + p86 + i32(1)
					}
				l59:
					v5 = v27 + i32(1)
					p87 := v4
					if uint32(v5) > uint32(v4) {
						p87 = v5
					}
					v5 = p87
					{
						t88 := int32(load32(m.memory[int64(uint32(v2))+32:]))
						if v3 != t88 {
							goto l62
						}
						m.fn318(v2 + i32(32))
						t89 := int32(load32(m.memory[int64(uint32(v2))+36:]))
						v28 = t89
					}
				l62:
					t90 := int64(load64(m.memory[int64(uint32(v2))+44:]))
					v15 = t90
					v4 = v28 + v3*i32(28)
					t91 := int32(load32(m.memory[int64(uint32(v2))+52:]))
					store32(m.memory[int64(uint32(v4))+8:], uint32(t91))
					store64(m.memory[uint32(v4):], uint64(v15))
					m.memory[int64(uint32(v4))+25] = byte(v22)
					m.memory[int64(uint32(v4))+24] = byte(v23)
					store32(m.memory[int64(uint32(v4))+20:], uint32(i32(1)))
					store32(m.memory[int64(uint32(v4))+16:], uint32(v5))
					store32(m.memory[int64(uint32(v4))+12:], uint32(v27))
					t92 := v2
					v3 = v3 + i32(1)
					store32(m.memory[int64(uint32(t92))+40:], uint32(v3))
					v27 = v5
					goto l63
				}
			l44:
				{
					t93 := int32(load32(m.memory[int64(uint32(v2))+28:]))
					v4 = t93
					t94 := int32(load32(m.memory[int64(uint32(v2))+20:]))
					if v4 != t94 {
						goto l64
					}
					m.fn316(v2 + i32(20))
				}
			l64:
				t95 := int32(load32(m.memory[int64(uint32(v2))+24:]))
				v23 = t95
				v1 = v23 + v4*i32(12)
				t96 := int64(load64(m.memory[int64(uint32(v2))+32:]))
				store64(m.memory[uint32(v1):], uint64(t96))
				t97 := int32(load32(m.memory[int64(uint32(v2))+40:]))
				store32(m.memory[int64(uint32(v1))+8:], uint32(t97))
				t98 := v2
				v22 = v4 + i32(1)
				store32(m.memory[int64(uint32(t98))+28:], uint32(v22))
				if v6 == v5 {
					goto l65
				}
				v6 = int32(uint32(v6-v5) >> 5)
				v7 = i32(0)
			l72:
				{
					v11 = v5 + v7<<5
					t99 := int32(load32(m.memory[int64(uint32(v11))+4:]))
					v14 = t99
					{
						t100 := int32(load32(m.memory[int64(uint32(v11))+8:]))
						v1 = t100
						if v1 == 0 {
							goto l66
						}
						v4 = v14
					l67:
						m.fn335(v4)
						v4 = v4 + i32(32)
						v1 = v1 + i32(-1)
						if v1 != 0 {
							goto l67
						}
					}
				l66:
					{
						t101 := int32(load32(m.memory[uint32(v11):]))
						v4 = t101
						if v4 == 0 {
							goto l68
						}
						t102 := int32(load32(m.memory[uint32(v14+i32(-4)):]))
						v1 = t102
						v11 = v1 & i32(-8)
						t103 := v11
						v1 = v1 & i32(3)
						p104 := i32(8)
						if v1 != 0 {
							p104 = i32(4)
						}
						v4 = v4 << 5
						if uint32(t103) < uint32(p104|v4) {
							goto l69
						}
						if v1 == 0 {
							goto l70
						}
						if uint32(v11) > uint32(v4+i32(39)) {
							m.fn3(i32(1273904), i32(46), i32(1273952))
							panic("unreachable")
						}
					l70:
						m.fn1(v14)
					}
				l68:
					v7 = v7 + i32(1)
					if v7 != v6 {
						goto l72
					}
				}
			l65:
				{
					if v25 == 0 {
						goto l73
					}
					t105 := int32(load32(m.memory[uint32(v26+i32(-4)):]))
					v4 = t105
					v5 = v4 & i32(-8)
					t106 := v5
					v4 = v4 & i32(3)
					p107 := i32(8)
					if v4 != 0 {
						p107 = i32(4)
					}
					v1 = v25 << 5
					if uint32(t106) < uint32(p107|v1) {
						m.fn3(i32(1273840), i32(46), i32(1273888))
						panic("unreachable")
					}
					if v4 == 0 {
						goto l75
					}
					if uint32(v5) > uint32(v1+i32(39)) {
						m.fn3(i32(1273904), i32(46), i32(1273952))
						panic("unreachable")
					}
				l75:
					m.fn1(v26)
				}
			l73:
				v4 = v24
				if v24 == v20 {
					goto l40
				}
				goto l77
			l69:
			}
			m.fn3(i32(1273840), i32(46), i32(1273888))
			panic("unreachable")
		}
		v23 = i32(4)
		v24 = v19
		goto l40
	}
l40:
	t108 := int32(uint32(v20-v24) / uint32(i32(12)))
	v6 = t108
	{
		{
			{
				if v20 == v24 {
					goto l78
				}
				v11 = i32(0)
			l91:
				{
					v14 = v24 + v11*i32(12)
					t109 := int32(load32(m.memory[int64(uint32(v14))+4:]))
					v12 = t109
					{
						t110 := int32(load32(m.memory[int64(uint32(v14))+8:]))
						v13 = t110
						if v13 == 0 {
							goto l79
						}
						v1 = i32(0)
					l86:
						{
							v7 = v12 + v1<<5
							t111 := int32(load32(m.memory[int64(uint32(v7))+4:]))
							v10 = t111
							{
								t112 := int32(load32(m.memory[int64(uint32(v7))+8:]))
								v5 = t112
								if v5 == 0 {
									goto l80
								}
								v4 = v10
							l81:
								m.fn335(v4)
								v4 = v4 + i32(32)
								v5 = v5 + i32(-1)
								if v5 != 0 {
									goto l81
								}
							}
						l80:
							{
								t113 := int32(load32(m.memory[uint32(v7):]))
								v4 = t113
								if v4 == 0 {
									goto l82
								}
								t114 := int32(load32(m.memory[uint32(v10+i32(-4)):]))
								v5 = t114
								v7 = v5 & i32(-8)
								t115 := v7
								v5 = v5 & i32(3)
								p116 := i32(8)
								if v5 != 0 {
									p116 = i32(4)
								}
								v4 = v4 << 5
								if uint32(t115) < uint32(p116|v4) {
									m.fn3(i32(1273840), i32(46), i32(1273888))
									panic("unreachable")
								}
								if v5 == 0 {
									goto l84
								}
								if uint32(v7) > uint32(v4+i32(39)) {
									m.fn3(i32(1273904), i32(46), i32(1273952))
									panic("unreachable")
								}
							l84:
								m.fn1(v10)
							}
						l82:
							v1 = v1 + i32(1)
							if v1 != v13 {
								goto l86
							}
						}
					}
				l79:
					{
						t117 := int32(load32(m.memory[uint32(v14):]))
						v4 = t117
						if v4 == 0 {
							goto l87
						}
						t118 := int32(load32(m.memory[uint32(v12+i32(-4)):]))
						v5 = t118
						v1 = v5 & i32(-8)
						t119 := v1
						v5 = v5 & i32(3)
						p120 := i32(8)
						if v5 != 0 {
							p120 = i32(4)
						}
						v4 = v4 << 5
						if uint32(t119) < uint32(p120|v4) {
							m.fn3(i32(1273840), i32(46), i32(1273888))
							panic("unreachable")
						}
						if v5 == 0 {
							goto l89
						}
						if uint32(v1) > uint32(v4+i32(39)) {
							m.fn3(i32(1273904), i32(46), i32(1273952))
							panic("unreachable")
						}
					l89:
						m.fn1(v12)
					}
				l87:
					v11 = v11 + i32(1)
					if v11 != v6 {
						goto l91
					}
				}
			l78:
				{
					if v17 == 0 {
						goto l92
					}
					t121 := int32(load32(m.memory[uint32(v19+i32(-4)):]))
					v4 = t121
					v5 = v4 & i32(-8)
					t122 := v5
					v4 = v4 & i32(3)
					p123 := i32(8)
					if v4 != 0 {
						p123 = i32(4)
					}
					if uint32(t122) < uint32(p123+v18) {
						m.fn3(i32(1273840), i32(46), i32(1273888))
						panic("unreachable")
					}
					if v4 == 0 {
						goto l94
					}
					if uint32(v5) > uint32(v18+i32(39)) {
						m.fn3(i32(1273904), i32(46), i32(1273952))
						panic("unreachable")
					}
				l94:
					m.fn1(v19)
				}
			l92:
				{
					{
						t124 := int32(m.memory[int64(uint32(i32(0)))+1293880])
						if t124 == 0 {
							goto l96
						}
						t125 := int64(load64(m.memory[int64(uint32(i32(0)))+1293872:]))
						v33 = t125
						t126 := int64(load64(m.memory[int64(uint32(i32(0)))+1293864:]))
						v32 = t126
						goto l97
					}
				l96:
					m.fn194(v2 + i32(152))
					m.memory[int64(uint32(i32(0)))+1293880] = byte(i32(1))
					t127 := int64(load64(m.memory[int64(uint32(v2))+160:]))
					v33 = t127
					store64(m.memory[int64(uint32(i32(0)))+1293872:], uint64(v33))
					t128 := int64(load64(m.memory[int64(uint32(v2))+152:]))
					v32 = t128
				}
			l97:
				store64(m.memory[int64(uint32(i32(0)))+1293864:], uint64(v32+i64(1)))
				if v22 != 0 {
					v21 = i32(1275656)
					v24 = v2 + i32(168)
					v31 = i32(0)
					v29 = i32(0)
					v27 = i32(0)
				l144:
					{
						v3 = v29
						v30 = v21
						{
							{
								t129 := int32(m.memory[int64(uint32(i32(0)))+1293880])
								if t129 == 0 {
									goto l100
								}
								t130 := int64(load64(m.memory[int64(uint32(i32(0)))+1293872:]))
								v16 = t130
								t131 := int64(load64(m.memory[int64(uint32(i32(0)))+1293864:]))
								v15 = t131
								goto l101
							}
						l100:
							m.fn194(v2 + i32(152))
							m.memory[int64(uint32(i32(0)))+1293880] = byte(i32(1))
							t132 := int64(load64(m.memory[int64(uint32(v2))+160:]))
							v16 = t132
							store64(m.memory[int64(uint32(i32(0)))+1293872:], uint64(v16))
							t133 := int64(load64(m.memory[int64(uint32(v2))+152:]))
							v15 = t133
						}
					l101:
						store64(m.memory[int64(uint32(v2))+168:], uint64(v15))
						v29 = i32(0)
						store64(m.memory[int64(uint32(i32(0)))+1293864:], uint64(v15+i64(1)))
						store64(m.memory[int64(uint32(v2))+176:], uint64(v16))
						t134 := int64(load64(m.memory[int64(uint32(i32(0)))+1275672:]))
						store64(m.memory[int64(uint32(v2))+160:], uint64(t134))
						t135 := int64(load64(m.memory[int64(uint32(i32(0)))+1275664:]))
						store64(m.memory[int64(uint32(v2))+152:], uint64(t135))
						{
							{
								v1 = v23 + v31*i32(12)
								t136 := int32(load32(m.memory[int64(uint32(v1))+8:]))
								v14 = t136
								if v14 != 0 {
									goto l102
								}
								v21 = i32(1275656)
								v33 = v16
								v32 = v15
								v27 = i32(0)
								goto l103
							}
						l102:
							t137 := v32 ^ i64(8317987319222330741)
							v15 = v33 ^ i64(7237128888997146477)
							v16 = t137 + v15
							v34 = i64_rotl(v16, i64(32))
							v35 = v16 ^ i64_rotl(v15, i64(13))
							v36 = i64_rotl(v35, i64(17))
							v37 = v33 ^ i64(8387220255154660723)
							v38 = v32 ^ i64(0x6c7967656e657261)
							v4 = i32(0)
							v12 = i32(1275656)
						l138:
							{
								{
									t138 := int32(load32(m.memory[int64(uint32(v1))+8:]))
									t139 := v4
									v13 = t138
									if uint32(t139) >= uint32(v13) {
										m.fn33(v4, v13, i32(1075692))
										panic("unreachable")
									}
									t140 := int32(load32(m.memory[int64(uint32(v1))+4:]))
									v11 = v4 * i32(28)
									v5 = t140 + v11
									t141 := int32(load32(m.memory[int64(uint32(v5))+16:]))
									v10 = t141
									t142 := int32(load32(m.memory[int64(uint32(v5))+12:]))
									v7 = t142
									t143 := int32(m.memory[int64(uint32(v5))+25])
									if t143 != i32(1) {
										goto l105
									}
									if v27 == 0 {
										goto l106
									}
									t144 := v3
									v15 = int64(uint32(v10))<<32 | int64(uint32(v7))
									v16 = v15 ^ v37
									v32 = v16 + v38
									v16 = v32 ^ i64_rotl(v16, i64(16))
									v33 = v16 + v34
									v16 = v33 ^ i64_rotl(v16, i64(21)) ^ i64(0x800000000000000)
									t145 := i64_rotl(v16, i64(16))
									t146 := v16
									v32 = v32 + v35
									v16 = t146 + i64_rotl(v32, i64(32))
									v39 = t145 ^ v16
									t147 := i64_rotl(v39, i64(21))
									t148 := v39
									t149 := v33 ^ v15
									v32 = v32 ^ v36
									v33 = t149 + v32
									v39 = t148 + i64_rotl(v33, i64(32))
									v40 = t147 ^ v39
									t150 := i64_rotl(v40, i64(16))
									t151 := v40
									t152 := v16
									v32 = v33 ^ i64_rotl(v32, i64(13))
									v16 = t152 + v32
									v33 = t151 + (i64_rotl(v16, i64(32)) ^ i64(255))
									v40 = t150 ^ v33
									t153 := i64_rotl(v40, i64(21))
									t154 := v40
									t155 := v39 ^ i64(0x800000000000000)
									v16 = v16 ^ i64_rotl(v32, i64(17))
									v32 = t155 + v16
									v39 = t154 + i64_rotl(v32, i64(32))
									v40 = t153 ^ v39
									t156 := i64_rotl(v40, i64(16))
									t157 := v40
									v16 = v32 ^ i64_rotl(v16, i64(13))
									v32 = v16 + v33
									v33 = t157 + i64_rotl(v32, i64(32))
									v40 = t156 ^ v33
									t158 := i64_rotl(v40, i64(21))
									t159 := v40
									v16 = v32 ^ i64_rotl(v16, i64(17))
									v32 = v16 + v39
									v39 = t159 + i64_rotl(v32, i64(32))
									v40 = t158 ^ v39
									t160 := i64_rotl(v40, i64(16))
									t161 := v40
									v16 = i64_rotl(v16, i64(13)) ^ v32
									v32 = v16 + v33
									v33 = t161 + i64_rotl(v32, i64(32))
									t162 := i64_rotl(t160^v33, i64(21))
									v16 = i64_rotl(v16, i64(17)) ^ v32
									v16 = i64_rotl(v16, i64(13)) ^ (v16 + v39)
									t163 := t162 ^ i64_rotl(v16, i64(17))
									v16 = v16 + v33
									v16 = t163 ^ int64(uint64(v16)>>32) ^ v16
									v13 = t144 & int32(v16)
									v32 = int64(uint64(v16)>>25) & i64(127) * i64(72340172838076673)
									v9 = i32(0)
								l111:
									{
										{
											t164 := int64(load64(m.memory[uint32(v30+v13):]))
											v33 = t164
											v16 = v33 ^ v32
											v16 = (v16 ^ i64(-1)) & (v16 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
											if v16 == 0 {
												goto l107
											}
										l110:
											{
												t165 := v7
												v6 = v30 - (int32(uint32(int64(bits.TrailingZeros64(uint64(v16))))>>3)+v13)&v3<<4
												t166 := int32(load32(m.memory[uint32(v6+i32(-16)):]))
												if t165 != t166 {
													goto l108
												}
												t167 := int32(load32(m.memory[uint32(v6+i32(-12)):]))
												if v10 == t167 {
													t169 := int32(load32(m.memory[uint32(v6+i32(-8)):]))
													v13 = t169
													if uint32(v13) >= uint32(v22) {
														m.fn33(v13, v22, i32(1075708))
														panic("unreachable")
													}
													t170 := int32(load32(m.memory[uint32(v6+i32(-4)):]))
													v11 = t170
													t171 := v11
													v5 = v23 + v13*i32(12)
													t172 := int32(load32(m.memory[int64(uint32(v5))+8:]))
													v6 = t172
													if uint32(t171) >= uint32(v6) {
														m.fn33(v11, v6, i32(1075724))
														panic("unreachable")
													}
													t173 := int32(load32(m.memory[int64(uint32(v5))+4:]))
													v5 = t173 + v11*i32(28)
													t174 := int32(load32(m.memory[int64(uint32(v5))+20:]))
													store32(m.memory[int64(uint32(v5))+20:], uint32(t174+i32(1)))
													t175 := int64(load64(m.memory[int64(uint32(v2))+176:]))
													t176 := v15
													v16 = t175
													v32 = t176 ^ v16 ^ i64(8387220255154660723)
													t177 := int64(load64(m.memory[int64(uint32(v2))+168:]))
													t178 := v32
													v33 = t177
													v39 = t178 + (v33 ^ i64(0x6c7967656e657261))
													v32 = v39 ^ i64_rotl(v32, i64(16))
													t179 := v32
													v16 = v16 ^ i64(7237128888997146477)
													v33 = v16 + (v33 ^ i64(8317987319222330741))
													v40 = t179 + i64_rotl(v33, i64(32))
													v32 = v40 ^ i64_rotl(v32, i64(21)) ^ i64(0x800000000000000)
													t180 := i64_rotl(v32, i64(16))
													t181 := v32
													v16 = i64_rotl(v16, i64(13)) ^ v33
													v33 = v16 + v39
													v32 = t181 + i64_rotl(v33, i64(32))
													v39 = t180 ^ v32
													t182 := i64_rotl(v39, i64(21))
													t183 := v39
													v16 = v33 ^ i64_rotl(v16, i64(17))
													v15 = v16 + (v40 ^ v15)
													v33 = t183 + i64_rotl(v15, i64(32))
													v39 = t182 ^ v33
													t184 := i64_rotl(v39, i64(16))
													t185 := v39
													t186 := v32
													v15 = i64_rotl(v16, i64(13)) ^ v15
													v16 = t186 + v15
													v32 = t185 + (i64_rotl(v16, i64(32)) ^ i64(255))
													v39 = t184 ^ v32
													t187 := i64_rotl(v39, i64(21))
													t188 := v39
													t189 := v33 ^ i64(0x800000000000000)
													v15 = v16 ^ i64_rotl(v15, i64(17))
													v16 = t189 + v15
													v33 = t188 + i64_rotl(v16, i64(32))
													v39 = t187 ^ v33
													t190 := i64_rotl(v39, i64(16))
													t191 := v39
													v15 = v16 ^ i64_rotl(v15, i64(13))
													v16 = v15 + v32
													v32 = t191 + i64_rotl(v16, i64(32))
													v39 = t190 ^ v32
													t192 := i64_rotl(v39, i64(21))
													t193 := v39
													v15 = v16 ^ i64_rotl(v15, i64(17))
													v16 = v15 + v33
													v33 = t193 + i64_rotl(v16, i64(32))
													v39 = t192 ^ v33
													t194 := i64_rotl(v39, i64(16))
													t195 := v39
													v15 = i64_rotl(v15, i64(13)) ^ v16
													v16 = v15 + v32
													v32 = t195 + i64_rotl(v16, i64(32))
													t196 := i64_rotl(t194^v32, i64(21))
													v15 = i64_rotl(v15, i64(17)) ^ v16
													v15 = i64_rotl(v15, i64(13)) ^ (v15 + v33)
													t197 := t196 ^ i64_rotl(v15, i64(17))
													v15 = v15 + v32
													v15 = t197 ^ int64(uint64(v15)>>32) ^ v15
													{
														t198 := int32(load32(m.memory[int64(uint32(v2))+160:]))
														if t198 != 0 {
															goto l114
														}
														_ = m.fn88(v2+i32(152), v24)
														t200 := int32(load32(m.memory[int64(uint32(v2))+152:]))
														v12 = t200
													}
												l114:
													t201 := int32(load32(m.memory[int64(uint32(v2))+156:]))
													v29 = t201
													v5 = v29 & int32(v15)
													v33 = int64(uint64(v15) >> 25)
													v16 = v33 & i64(127) * i64(72340172838076673)
													v21 = i32(0)
													v28 = i32(0)
												l125:
													{
														t202 := int64(load64(m.memory[uint32(v12+v5):]))
														v32 = t202
														v15 = v32 ^ v16
														v15 = (v15 ^ i64(-1)) & (v15 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
														if v15 == 0 {
															goto l115
														}
													l118:
														{
															t203 := v7
															v6 = v12 - (int32(uint32(int64(bits.TrailingZeros64(uint64(v15))))>>3)+v5)&v29<<4
															t204 := int32(load32(m.memory[uint32(v6+i32(-16)):]))
															if t203 != t204 {
																goto l116
															}
															t205 := int32(load32(m.memory[uint32(v6+i32(-12)):]))
															if v10 == t205 {
																store32(m.memory[uint32(v6+i32(-4)):], uint32(v11))
																store32(m.memory[uint32(v6+i32(-8)):], uint32(v13))
																goto l124
															}
														}
													l116:
														v15 = (v15 + i64(-1)) & v15
														if !(v15 == 0) {
															goto l118
														}
													}
												l115:
													v15 = v32 & i64(-0x7f7f7f7f7f7f7f80)
													if v21 == i32(1) {
														goto l119
													}
													if v15 == 0 {
														goto l120
													}
													v9 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v15))))>>3) + v5) & v29
												l119:
													if v15&(v32<<1) != i64(0) {
														{
															t206 := int32(int8(m.memory[uint32(v12+v9)]))
															v6 = t206
															if v6 < i32(0) {
																goto l123
															}
															t207 := int64(load64(m.memory[uint32(v12):]))
															t208 := v12
															v9 = int32(uint32(int64(bits.TrailingZeros64(uint64(t207&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
															t209 := int32(m.memory[uint32(t208+v9)])
															v6 = t209
														}
													l123:
														t210 := v12 + v9
														v5 = int32(v33) & i32(127)
														m.memory[uint32(t210)] = byte(v5)
														m.memory[uint32(v12+(v9+i32(-8))&v29+i32(8))] = byte(v5)
														v5 = v12 - v9<<4
														store32(m.memory[uint32(v5+i32(-16)):], uint32(v7))
														store32(m.memory[uint32(v5+i32(-12)):], uint32(v10))
														store32(m.memory[uint32(v5+i32(-8)):], uint32(v13))
														store32(m.memory[uint32(v5+i32(-4)):], uint32(v11))
														t211 := int32(load32(m.memory[int64(uint32(v2))+164:]))
														store32(m.memory[int64(uint32(v2))+164:], uint32(t211+i32(1)))
														t212 := int32(load32(m.memory[int64(uint32(v2))+160:]))
														store32(m.memory[int64(uint32(v2))+160:], uint32(t212-v6&i32(1)))
														goto l124
													}
													v21 = i32(1)
													goto l122
												l120:
													v21 = i32(0)
												l122:
													v28 = v28 + i32(8)
													v5 = (v28 + v5) & v29
													goto l125
												}
											}
										l108:
											v16 = (v16 + i64(-1)) & v16
											if !(v16 == 0) {
												goto l110
											}
										}
									l107:
										if !(v33&(v33<<1)&i64(-0x7f7f7f7f7f7f7f80) == 0) {
											goto l106
										}
										t168 := v13
										v9 = v9 + i32(8)
										v13 = (t168 + v9) & v3
										goto l111
									}
								}
							l106:
								m.memory[int64(uint32(v5))+25] = byte(i32(0))
								t213 := int32(load32(m.memory[int64(uint32(v1))+8:]))
								v13 = t213
							}
						l105:
							{
								if uint32(v4) >= uint32(v13) {
									m.fn33(v4, v13, i32(1075740))
									panic("unreachable")
								}
								t214 := int32(load32(m.memory[int64(uint32(v1))+4:]))
								t215 := int32(m.memory[int64(uint32(t214+v11))+24])
								if t215 == 0 {
									goto l124
								}
								v15 = int64(uint32(v10))<<32 | int64(uint32(v7))
								t216 := int64(load64(m.memory[int64(uint32(v2))+176:]))
								t217 := v15
								v16 = t216
								v32 = t217 ^ v16 ^ i64(8387220255154660723)
								t218 := int64(load64(m.memory[int64(uint32(v2))+168:]))
								t219 := v32
								v33 = t218
								v39 = t219 + (v33 ^ i64(0x6c7967656e657261))
								v32 = v39 ^ i64_rotl(v32, i64(16))
								t220 := v32
								v16 = v16 ^ i64(7237128888997146477)
								v33 = v16 + (v33 ^ i64(8317987319222330741))
								v40 = t220 + i64_rotl(v33, i64(32))
								v32 = v40 ^ i64_rotl(v32, i64(21)) ^ i64(0x800000000000000)
								t221 := i64_rotl(v32, i64(16))
								t222 := v32
								v16 = i64_rotl(v16, i64(13)) ^ v33
								v33 = v16 + v39
								v32 = t222 + i64_rotl(v33, i64(32))
								v39 = t221 ^ v32
								t223 := i64_rotl(v39, i64(21))
								t224 := v39
								v16 = v33 ^ i64_rotl(v16, i64(17))
								v15 = v16 + (v40 ^ v15)
								v33 = t224 + i64_rotl(v15, i64(32))
								v39 = t223 ^ v33
								t225 := i64_rotl(v39, i64(16))
								t226 := v39
								t227 := v32
								v15 = i64_rotl(v16, i64(13)) ^ v15
								v16 = t227 + v15
								v32 = t226 + (i64_rotl(v16, i64(32)) ^ i64(255))
								v39 = t225 ^ v32
								t228 := i64_rotl(v39, i64(21))
								t229 := v39
								t230 := v33 ^ i64(0x800000000000000)
								v15 = v16 ^ i64_rotl(v15, i64(17))
								v16 = t230 + v15
								v33 = t229 + i64_rotl(v16, i64(32))
								v39 = t228 ^ v33
								t231 := i64_rotl(v39, i64(16))
								t232 := v39
								v15 = v16 ^ i64_rotl(v15, i64(13))
								v16 = v15 + v32
								v32 = t232 + i64_rotl(v16, i64(32))
								v39 = t231 ^ v32
								t233 := i64_rotl(v39, i64(21))
								t234 := v39
								v15 = v16 ^ i64_rotl(v15, i64(17))
								v16 = v15 + v33
								v33 = t234 + i64_rotl(v16, i64(32))
								v39 = t233 ^ v33
								t235 := i64_rotl(v39, i64(16))
								t236 := v39
								v15 = i64_rotl(v15, i64(13)) ^ v16
								v16 = v15 + v32
								v32 = t236 + i64_rotl(v16, i64(32))
								t237 := i64_rotl(t235^v32, i64(21))
								v15 = i64_rotl(v15, i64(17)) ^ v16
								v15 = i64_rotl(v15, i64(13)) ^ (v15 + v33)
								t238 := t237 ^ i64_rotl(v15, i64(17))
								v15 = v15 + v32
								v15 = t238 ^ int64(uint64(v15)>>32) ^ v15
								{
									t239 := int32(load32(m.memory[int64(uint32(v2))+160:]))
									if t239 != 0 {
										goto l127
									}
									_ = m.fn88(v2+i32(152), v24)
								}
							l127:
								t241 := int32(load32(m.memory[int64(uint32(v2))+156:]))
								v6 = t241
								v5 = v6 & int32(v15)
								v33 = int64(uint64(v15) >> 25)
								v16 = v33 & i64(127) * i64(72340172838076673)
								v9 = i32(0)
								t242 := int32(load32(m.memory[int64(uint32(v2))+152:]))
								v12 = t242
								v29 = i32(0)
							l137:
								{
									t243 := int64(load64(m.memory[uint32(v12+v5):]))
									v32 = t243
									v15 = v32 ^ v16
									v15 = (v15 ^ i64(-1)) & (v15 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
									if v15 == 0 {
										goto l128
									}
								l131:
									{
										t244 := v7
										v13 = v12 - (int32(uint32(int64(bits.TrailingZeros64(uint64(v15))))>>3)+v5)&v6<<4
										t245 := int32(load32(m.memory[uint32(v13+i32(-16)):]))
										if t244 != t245 {
											goto l129
										}
										t246 := int32(load32(m.memory[uint32(v13+i32(-12)):]))
										if v10 == t246 {
											store32(m.memory[uint32(v13+i32(-4)):], uint32(v4))
											store32(m.memory[uint32(v13+i32(-8)):], uint32(v31))
											goto l124
										}
									}
								l129:
									v15 = (v15 + i64(-1)) & v15
									if !(v15 == 0) {
										goto l131
									}
								}
							l128:
								v15 = v32 & i64(-0x7f7f7f7f7f7f7f80)
								if v9 == i32(1) {
									goto l132
								}
								if v15 == 0 {
									goto l133
								}
								v11 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v15))))>>3) + v5) & v6
							l132:
								if v15&(v32<<1) != i64(0) {
									{
										t247 := int32(int8(m.memory[uint32(v12+v11)]))
										v13 = t247
										if v13 < i32(0) {
											goto l136
										}
										t248 := int64(load64(m.memory[uint32(v12):]))
										t249 := v12
										v11 = int32(uint32(int64(bits.TrailingZeros64(uint64(t248&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
										t250 := int32(m.memory[uint32(t249+v11)])
										v13 = t250
									}
								l136:
									t251 := v12 + v11
									v5 = int32(v33) & i32(127)
									m.memory[uint32(t251)] = byte(v5)
									m.memory[uint32(v12+(v11+i32(-8))&v6+i32(8))] = byte(v5)
									v5 = v12 - v11<<4
									store32(m.memory[uint32(v5+i32(-16)):], uint32(v7))
									store32(m.memory[uint32(v5+i32(-12)):], uint32(v10))
									store32(m.memory[uint32(v5+i32(-8)):], uint32(v31))
									store32(m.memory[uint32(v5+i32(-4)):], uint32(v4))
									t252 := int32(load32(m.memory[int64(uint32(v2))+164:]))
									store32(m.memory[int64(uint32(v2))+164:], uint32(t252+i32(1)))
									t253 := int32(load32(m.memory[int64(uint32(v2))+160:]))
									store32(m.memory[int64(uint32(v2))+160:], uint32(t253-v13&i32(1)))
									goto l124
								}
								v9 = i32(1)
								goto l135
							l133:
								v9 = i32(0)
							l135:
								v29 = v29 + i32(8)
								v5 = (v29 + v5) & v6
								goto l137
							}
						l124:
							v4 = v4 + i32(1)
							if v4 != v14 {
								goto l138
							}
							t254 := int64(load64(m.memory[int64(uint32(v2))+176:]))
							v33 = t254
							t255 := int64(load64(m.memory[int64(uint32(v2))+168:]))
							v32 = t255
							t256 := int32(load32(m.memory[int64(uint32(v2))+164:]))
							v27 = t256
							t257 := int32(load32(m.memory[int64(uint32(v2))+156:]))
							v29 = t257
							t258 := int32(load32(m.memory[int64(uint32(v2))+152:]))
							v21 = t258
						}
					l103:
						{
							if v3 == 0 {
								goto l139
							}
							v5 = v3 << 4
							v4 = v5 + v3 + i32(25)
							if v4 == 0 {
								goto l139
							}
							v1 = v30 - v5
							t259 := int32(load32(m.memory[uint32(v1+i32(-20)):]))
							v5 = t259
							v7 = v5 & i32(-8)
							t260 := v7
							v5 = v5 & i32(3)
							p261 := i32(8)
							if v5 != 0 {
								p261 = i32(4)
							}
							if uint32(t260) < uint32(p261+v4) {
								goto l140
							}
							if v5 == 0 {
								goto l141
							}
							if uint32(v7) > uint32(v4+i32(39)) {
								m.fn3(i32(1273904), i32(46), i32(1273952))
								panic("unreachable")
							}
						l141:
							m.fn1(v1 + i32(-16))
						}
					l139:
						v31 = v31 + i32(1)
						if v31 == v22 {
							goto l143
						}
						goto l144
					l140:
					}
					m.fn3(i32(1273840), i32(46), i32(1273888))
					panic("unreachable")
				}
				v29 = i32(0)
				v21 = i32(1275656)
				goto l99
			l143:
				t262 := int32(m.memory[int64(uint32(i32(0)))+1293880])
				if t262 == 0 {
					goto l145
				}
			}
		l99:
			t263 := int64(load64(m.memory[int64(uint32(i32(0)))+1293872:]))
			v16 = t263
			t264 := int64(load64(m.memory[int64(uint32(i32(0)))+1293864:]))
			v15 = t264
			goto l146
		}
	l145:
		m.fn194(v2 + i32(152))
		m.memory[int64(uint32(i32(0)))+1293880] = byte(i32(1))
		t265 := int64(load64(m.memory[int64(uint32(v2))+160:]))
		v16 = t265
		store64(m.memory[int64(uint32(i32(0)))+1293872:], uint64(v16))
		t266 := int64(load64(m.memory[int64(uint32(v2))+152:]))
		v15 = t266
	}
l146:
	store64(m.memory[int64(uint32(v2))+72:], uint64(v15))
	store64(m.memory[int64(uint32(i32(0)))+1293864:], uint64(v15+i64(1)))
	store32(m.memory[int64(uint32(v2))+104:], uint32(i32(0)))
	store64(m.memory[int64(uint32(v2))+96:], uint64(i64(0x400000000)))
	store64(m.memory[int64(uint32(v2))+88:], uint64(i64(0)))
	store64(m.memory[int64(uint32(v2))+80:], uint64(v16))
	t267 := int64(load64(m.memory[int64(uint32(i32(0)))+1275664:]))
	store64(m.memory[int64(uint32(v2))+56:], uint64(t267))
	t268 := int64(load64(m.memory[int64(uint32(i32(0)))+1275672:]))
	store64(m.memory[int64(uint32(v2))+64:], uint64(t268))
	t269 := int32(load32(m.memory[int64(uint32(v2))+20:]))
	v5 = t269
	t270 := int32(load32(m.memory[int64(uint32(v2))+24:]))
	t271 := v2
	v4 = t270
	v20 = v4 + v22*i32(12)
	store32(m.memory[int64(uint32(t271))+128:], uint32(v20))
	store32(m.memory[int64(uint32(v2))+124:], uint32(v5))
	store32(m.memory[int64(uint32(v2))+120:], uint32(v4))
	store32(m.memory[int64(uint32(v2))+116:], uint32(v4))
	{
		{
			{
				if v22 == 0 {
					goto l147
				}
				v26 = v2 + i32(96)
			l203:
				{
					t272 := v2
					v28 = v4 + i32(12)
					store32(m.memory[int64(uint32(t272))+120:], uint32(v28))
					t273 := int32(load32(m.memory[uint32(v4):]))
					v24 = t273
					if v24 == i32(-1) {
						goto l147
					}
					t274 := int32(load32(m.memory[int64(uint32(v4))+8:]))
					v5 = t274
					t275 := int32(load32(m.memory[int64(uint32(v4))+4:]))
					v23 = t275
					{
						t276 := int32(load32(m.memory[int64(uint32(v2))+104:]))
						v4 = t276
						t277 := int32(load32(m.memory[int64(uint32(v2))+96:]))
						if v4 != t277 {
							goto l148
						}
						m.fn316(v26)
					}
				l148:
					t278 := int32(load32(m.memory[int64(uint32(v2))+100:]))
					v1 = t278 + v4*i32(12)
					store32(m.memory[int64(uint32(v1))+8:], uint32(i32(0)))
					store64(m.memory[uint32(v1):], uint64(i64(0x400000000)))
					store32(m.memory[int64(uint32(v2))+104:], uint32(v4+i32(1)))
					v22 = v23 + v5*i32(28)
					v30 = v23
					{
						if v5 == 0 {
							goto l149
						}
					l190:
						{
							v4 = v30
							v30 = v4 + i32(28)
							t279 := int32(load32(m.memory[uint32(v4):]))
							v27 = t279
							if v27 == i32(-1) {
								goto l149
							}
							t280 := int32(load32(m.memory[int64(uint32(v4))+8:]))
							v7 = t280
							t281 := int32(load32(m.memory[int64(uint32(v4))+4:]))
							v31 = t281
							t282 := int32(load32(m.memory[int64(uint32(v4))+16:]))
							v5 = t282
							t283 := int32(load32(m.memory[int64(uint32(v4))+12:]))
							t284 := v5
							v1 = t283
							v3 = t284 - v1
							{
								t285 := int32(m.memory[int64(uint32(v4))+25])
								if t285&i32(1) == 0 {
									t346 := int32(load32(m.memory[int64(uint32(v4))+20:]))
									v4 = t346
									store32(m.memory[int64(uint32(v2))+140:], uint32(v7))
									store32(m.memory[int64(uint32(v2))+136:], uint32(v31))
									store32(m.memory[int64(uint32(v2))+132:], uint32(v27))
									t348 := v2
									p347 := i32(1)
									if uint32(v4) > uint32(i32(1)) {
										p347 = v4
									}
									store32(m.memory[int64(uint32(t348))+148:], uint32(p347))
									t350 := v2
									p349 := i32(1)
									if uint32(v3) > uint32(i32(1)) {
										p349 = v3
									}
									store32(m.memory[int64(uint32(t350))+144:], uint32(p349))
									m.fn334(v2+i32(152), v2+i32(56), v2+i32(132))
									t351 := int32(load32(m.memory[int64(uint32(v2))+152:]))
									if t351 == i32(-1) {
										goto l170
									}
									t352 := int64(load64(m.memory[int64(uint32(v2))+168:]))
									store64(m.memory[int64(uint32(v0))+20:], uint64(t352))
									t353 := int64(load64(m.memory[int64(uint32(v2))+160:]))
									store64(m.memory[int64(uint32(v0))+12:], uint64(t353))
									t354 := int64(load64(m.memory[int64(uint32(v2))+152:]))
									store64(m.memory[int64(uint32(v0))+4:], uint64(t354))
									store32(m.memory[uint32(v0):], uint32(i32(-2)))
									t355 := int32(uint32(v22-v30) / uint32(i32(28)))
									v12 = t355
									if v22 == v30 {
										goto l174
									}
									v1 = i32(0)
								l181:
									{
										v7 = v30 + v1*i32(28)
										t356 := int32(load32(m.memory[int64(uint32(v7))+4:]))
										v10 = t356
										{
											t357 := int32(load32(m.memory[int64(uint32(v7))+8:]))
											v5 = t357
											if v5 == 0 {
												goto l175
											}
											v4 = v10
										l176:
											m.fn335(v4)
											v4 = v4 + i32(32)
											v5 = v5 + i32(-1)
											if v5 != 0 {
												goto l176
											}
										}
									l175:
										{
											t358 := int32(load32(m.memory[uint32(v7):]))
											v4 = t358
											if v4 == 0 {
												goto l177
											}
											t359 := int32(load32(m.memory[uint32(v10+i32(-4)):]))
											v5 = t359
											v7 = v5 & i32(-8)
											t360 := v7
											v5 = v5 & i32(3)
											p361 := i32(8)
											if v5 != 0 {
												p361 = i32(4)
											}
											v4 = v4 << 5
											if uint32(t360) < uint32(p361|v4) {
												goto l178
											}
											if v5 == 0 {
												goto l179
											}
											if uint32(v7) > uint32(v4+i32(39)) {
												m.fn3(i32(1273904), i32(46), i32(1273952))
												panic("unreachable")
											}
										l179:
											m.fn1(v10)
										}
									l177:
										v1 = v1 + i32(1)
										if v1 != v12 {
											goto l181
										}
									}
								l174:
									{
										if v24 == 0 {
											goto l182
										}
										t362 := int32(load32(m.memory[uint32(v23+i32(-4)):]))
										v4 = t362
										v5 = v4 & i32(-8)
										t363 := v5
										v4 = v4 & i32(3)
										p364 := i32(8)
										if v4 != 0 {
											p364 = i32(4)
										}
										v1 = v24 * i32(28)
										if uint32(t363) < uint32(p364+v1) {
											m.fn3(i32(1273840), i32(46), i32(1273888))
											panic("unreachable")
										}
										if v4 == 0 {
											goto l184
										}
										if uint32(v5) > uint32(v1+i32(39)) {
											m.fn3(i32(1273904), i32(46), i32(1273952))
											panic("unreachable")
										}
									l184:
										m.fn1(v23)
									}
								l182:
									m.fn576(v2 + i32(116))
									m.fn362(v26)
									t365 := int32(load32(m.memory[int64(uint32(v2))+60:]))
									v4 = t365
									if v4 == 0 {
										goto l186
									}
									v5 = v4 << 4
									v4 = v5 + v4 + i32(25)
									if v4 == 0 {
										goto l186
									}
									t366 := int32(load32(m.memory[int64(uint32(v2))+56:]))
									v1 = t366 - v5
									t367 := int32(load32(m.memory[uint32(v1+i32(-20)):]))
									v5 = t367
									v7 = v5 & i32(-8)
									t368 := v7
									v5 = v5 & i32(3)
									p369 := i32(8)
									if v5 != 0 {
										p369 = i32(4)
									}
									if uint32(t368) < uint32(p369+v4) {
										m.fn3(i32(1273840), i32(46), i32(1273888))
										panic("unreachable")
									}
									if v5 == 0 {
										goto l188
									}
									if uint32(v7) > uint32(v4+i32(39)) {
										m.fn3(i32(1273904), i32(46), i32(1273952))
										panic("unreachable")
									}
								l188:
									m.fn1(v1 + i32(-16))
									goto l186
								}
								{
									if v5 == v1 {
										goto l151
									}
									v1 = i32(0)
									t286 := int32(load32(m.memory[int64(uint32(v2))+104:]))
									v5 = t286
								l167:
									{
										{
											{
												if v5 == 0 {
													goto l152
												}
												v4 = v5 + i32(-1)
												t287 := int32(load32(m.memory[int64(uint32(v2))+100:]))
												v5 = t287
												goto l153
											}
										l152:
											{
												t288 := int32(load32(m.memory[int64(uint32(v2))+96:]))
												if t288 != 0 {
													goto l154
												}
												m.fn316(v26)
											}
										l154:
											v4 = i32(0)
											t289 := int32(load32(m.memory[int64(uint32(v2))+100:]))
											v5 = t289
											store32(m.memory[int64(uint32(v5))+8:], uint32(i32(0)))
											store64(m.memory[uint32(v5):], uint64(i64(0x400000000)))
											store32(m.memory[int64(uint32(v2))+104:], uint32(i32(1)))
										}
									l153:
										v1 = v1 + i32(1)
										t290 := int32(load32(m.memory[int64(uint32(v2))+60:]))
										v13 = t290
										t291 := v13
										t292 := v5
										v11 = v4 * i32(12)
										t293 := int32(load32(m.memory[int64(uint32(t292+v11))+8:]))
										v6 = t293
										v15 = int64(uint32(v6))<<32 | int64(uint32(v4))
										t294 := int64(load64(m.memory[int64(uint32(v2))+80:]))
										t295 := v15
										v16 = t294
										v32 = t295 ^ v16 ^ i64(8387220255154660723)
										t296 := int64(load64(m.memory[int64(uint32(v2))+72:]))
										t297 := v32
										v33 = t296
										v39 = t297 + (v33 ^ i64(0x6c7967656e657261))
										v32 = v39 ^ i64_rotl(v32, i64(16))
										t298 := v32
										v16 = v16 ^ i64(7237128888997146477)
										v33 = v16 + (v33 ^ i64(8317987319222330741))
										v40 = t298 + i64_rotl(v33, i64(32))
										v32 = v40 ^ i64_rotl(v32, i64(21)) ^ i64(0x800000000000000)
										t299 := i64_rotl(v32, i64(16))
										t300 := v32
										v16 = i64_rotl(v16, i64(13)) ^ v33
										v33 = v16 + v39
										v32 = t300 + i64_rotl(v33, i64(32))
										v39 = t299 ^ v32
										t301 := i64_rotl(v39, i64(21))
										t302 := v39
										v16 = v33 ^ i64_rotl(v16, i64(17))
										v15 = v16 + (v40 ^ v15)
										v33 = t302 + i64_rotl(v15, i64(32))
										v39 = t301 ^ v33
										t303 := i64_rotl(v39, i64(16))
										t304 := v39
										t305 := v32
										v15 = i64_rotl(v16, i64(13)) ^ v15
										v16 = t305 + v15
										v32 = t304 + (i64_rotl(v16, i64(32)) ^ i64(255))
										v39 = t303 ^ v32
										t306 := i64_rotl(v39, i64(21))
										t307 := v39
										t308 := v33 ^ i64(0x800000000000000)
										v15 = v16 ^ i64_rotl(v15, i64(17))
										v16 = t308 + v15
										v33 = t307 + i64_rotl(v16, i64(32))
										v39 = t306 ^ v33
										t309 := i64_rotl(v39, i64(16))
										t310 := v39
										v15 = v16 ^ i64_rotl(v15, i64(13))
										v16 = v15 + v32
										v32 = t310 + i64_rotl(v16, i64(32))
										v39 = t309 ^ v32
										t311 := i64_rotl(v39, i64(21))
										t312 := v39
										v15 = v16 ^ i64_rotl(v15, i64(17))
										v16 = v15 + v33
										v33 = t312 + i64_rotl(v16, i64(32))
										v39 = t311 ^ v33
										t313 := i64_rotl(v39, i64(16))
										t314 := v39
										v15 = i64_rotl(v15, i64(13)) ^ v16
										v16 = v15 + v32
										v32 = t314 + i64_rotl(v16, i64(32))
										t315 := i64_rotl(t313^v32, i64(21))
										v15 = i64_rotl(v15, i64(17)) ^ v16
										v15 = i64_rotl(v15, i64(13)) ^ (v15 + v33)
										t316 := t315 ^ i64_rotl(v15, i64(17))
										v15 = v15 + v32
										v15 = t316 ^ int64(uint64(v15)>>32) ^ v15
										v10 = t291 & int32(v15)
										v16 = int64(uint64(v15)>>25) & i64(127) * i64(72340172838076673)
										v9 = i32(0)
										t317 := int32(load32(m.memory[int64(uint32(v2))+56:]))
										v5 = t317
									l160:
										{
											{
												t318 := int64(load64(m.memory[uint32(v5+v10):]))
												v32 = t318
												v15 = v32 ^ v16
												v15 = (v15 ^ i64(-1)) & (v15 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
												if v15 == 0 {
													goto l155
												}
											l158:
												{
													t319 := v4
													t320 := v5
													v14 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v15))))>>3) + v10) & v13
													v12 = t320 - v14<<4
													t321 := int32(load32(m.memory[uint32(v12+i32(-16)):]))
													if t319 != t321 {
														goto l156
													}
													t322 := int32(load32(m.memory[uint32(v12+i32(-12)):]))
													if v6 == t322 {
														v10 = i32(128)
														{
															v6 = v5 + v14
															t324 := int64(load64(m.memory[uint32(v6):]))
															v15 = t324
															t325 := int32(uint32(int64(bits.TrailingZeros64(uint64(v15&(v15<<1)&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
															v5 = v5 + (v14+i32(-8))&v13
															t326 := int64(load64(m.memory[uint32(v5):]))
															v15 = t326
															if uint32(t325+int32(uint32(int64(bits.LeadingZeros64(uint64(v15&(v15<<1)&i64(-0x7f7f7f7f7f7f7f80)))))>>3)) > uint32(i32(7)) {
																goto l161
															}
															t327 := int32(load32(m.memory[int64(uint32(v2))+64:]))
															store32(m.memory[int64(uint32(v2))+64:], uint32(t327+i32(1)))
															v10 = i32(255)
														}
													l161:
														m.memory[uint32(v6)] = byte(v10)
														m.memory[uint32(v5+i32(8))] = byte(v10)
														t328 := int32(load32(m.memory[int64(uint32(v2))+68:]))
														store32(m.memory[int64(uint32(v2))+68:], uint32(t328+i32(-1)))
														t329 := int32(load32(m.memory[int64(uint32(v2))+104:]))
														t330 := v4
														v5 = t329
														if uint32(t330) >= uint32(v5) {
															m.fn33(v4, v5, i32(1073136))
															panic("unreachable")
														}
														t331 := int32(load32(m.memory[uint32(v12+i32(-4)):]))
														v13 = t331
														t332 := int32(load32(m.memory[uint32(v12+i32(-8)):]))
														v14 = t332
														{
															t333 := int32(load32(m.memory[int64(uint32(v2))+100:]))
															v4 = t333 + v11
															t334 := int32(load32(m.memory[int64(uint32(v4))+8:]))
															v10 = t334
															t335 := int32(load32(m.memory[uint32(v4):]))
															if v10 != t335 {
																goto l163
															}
															m.fn191(v4)
														}
													l163:
														t336 := int32(load32(m.memory[int64(uint32(v4))+4:]))
														v12 = t336 + v10*i32(20)
														store32(m.memory[int64(uint32(v12))+8:], uint32(v13))
														store32(m.memory[int64(uint32(v12))+4:], uint32(v14))
														store32(m.memory[uint32(v12):], uint32(i32(-1)))
														store32(m.memory[int64(uint32(v4))+8:], uint32(v10+i32(1)))
														goto l164
													}
												}
											l156:
												v15 = (v15 + i64(-1)) & v15
												if !(v15 == 0) {
													goto l158
												}
											}
										l155:
											if !(v32&(v32<<1)&i64(-0x7f7f7f7f7f7f7f80) == 0) {
												t337 := int32(load32(m.memory[int64(uint32(v2))+104:]))
												t338 := v4
												v5 = t337
												if uint32(t338) >= uint32(v5) {
													m.fn33(v4, v5, i32(1073120))
													panic("unreachable")
												}
												{
													t339 := int32(load32(m.memory[int64(uint32(v2))+100:]))
													v4 = t339 + v11
													t340 := int32(load32(m.memory[int64(uint32(v4))+8:]))
													v10 = t340
													t341 := int32(load32(m.memory[uint32(v4):]))
													if v10 != t341 {
														goto l166
													}
													m.fn191(v4)
												}
											l166:
												t342 := int32(load32(m.memory[int64(uint32(v4))+4:]))
												v12 = t342 + v10*i32(20)
												store32(m.memory[int64(uint32(v12))+16:], uint32(i32(0)))
												store64(m.memory[int64(uint32(v12))+8:], uint64(i64(0)))
												store64(m.memory[uint32(v12):], uint64(i64(0x800000000)))
												store32(m.memory[int64(uint32(v4))+8:], uint32(v10+i32(1)))
												goto l164
											}
											t323 := v10
											v9 = v9 + i32(8)
											v10 = (t323 + v9) & v13
											goto l160
										}
									l164:
										if v1 != v3 {
											goto l167
										}
									}
								}
							l151:
								if v7 == 0 {
									goto l168
								}
								v4 = v31
							l169:
								m.fn335(v4)
								v4 = v4 + i32(32)
								v7 = v7 + i32(-1)
								if v7 != 0 {
									goto l169
								}
							l168:
								if v27 == 0 {
									goto l170
								}
								{
									t343 := int32(load32(m.memory[uint32(v31+i32(-4)):]))
									v4 = t343
									v5 = v4 & i32(-8)
									t344 := v5
									v4 = v4 & i32(3)
									p345 := i32(8)
									if v4 != 0 {
										p345 = i32(4)
									}
									v1 = v27 << 5
									if uint32(t344) < uint32(p345|v1) {
										m.fn3(i32(1273840), i32(46), i32(1273888))
										panic("unreachable")
									}
									if v4 == 0 {
										goto l172
									}
									if uint32(v5) > uint32(v1+i32(39)) {
										m.fn3(i32(1273904), i32(46), i32(1273952))
										panic("unreachable")
									}
								l172:
									m.fn1(v31)
									goto l170
								}
							}
						l170:
							if v30 != v22 {
								goto l190
							}
							goto l191
						l178:
						}
						m.fn3(i32(1273840), i32(46), i32(1273888))
						panic("unreachable")
					l149:
						t370 := int32(uint32(v22-v30) / uint32(i32(28)))
						v12 = t370
						if v22 == v30 {
							goto l191
						}
						v1 = i32(0)
					l198:
						{
							v7 = v30 + v1*i32(28)
							t371 := int32(load32(m.memory[int64(uint32(v7))+4:]))
							v10 = t371
							{
								t372 := int32(load32(m.memory[int64(uint32(v7))+8:]))
								v5 = t372
								if v5 == 0 {
									goto l192
								}
								v4 = v10
							l193:
								m.fn335(v4)
								v4 = v4 + i32(32)
								v5 = v5 + i32(-1)
								if v5 != 0 {
									goto l193
								}
							}
						l192:
							{
								t373 := int32(load32(m.memory[uint32(v7):]))
								v4 = t373
								if v4 == 0 {
									goto l194
								}
								t374 := int32(load32(m.memory[uint32(v10+i32(-4)):]))
								v5 = t374
								v7 = v5 & i32(-8)
								t375 := v7
								v5 = v5 & i32(3)
								p376 := i32(8)
								if v5 != 0 {
									p376 = i32(4)
								}
								v4 = v4 << 5
								if uint32(t375) < uint32(p376|v4) {
									m.fn3(i32(1273840), i32(46), i32(1273888))
									panic("unreachable")
								}
								if v5 == 0 {
									goto l196
								}
								if uint32(v7) > uint32(v4+i32(39)) {
									m.fn3(i32(1273904), i32(46), i32(1273952))
									panic("unreachable")
								}
							l196:
								m.fn1(v10)
							}
						l194:
							v1 = v1 + i32(1)
							if v1 != v12 {
								goto l198
							}
						}
					}
				l191:
					{
						if v24 == 0 {
							goto l199
						}
						t377 := int32(load32(m.memory[uint32(v23+i32(-4)):]))
						v4 = t377
						v5 = v4 & i32(-8)
						t378 := v5
						v4 = v4 & i32(3)
						p379 := i32(8)
						if v4 != 0 {
							p379 = i32(4)
						}
						v1 = v24 * i32(28)
						if uint32(t378) < uint32(p379+v1) {
							m.fn3(i32(1273840), i32(46), i32(1273888))
							panic("unreachable")
						}
						if v4 == 0 {
							goto l201
						}
						if uint32(v5) > uint32(v1+i32(39)) {
							m.fn3(i32(1273904), i32(46), i32(1273952))
							panic("unreachable")
						}
					l201:
						m.fn1(v23)
					}
				l199:
					v4 = v28
					if v28 != v20 {
						goto l203
					}
				}
			l147:
				m.fn576(v2 + i32(116))
				t380 := int64(load64(m.memory[int64(uint32(v2))+104:]))
				store64(m.memory[int64(uint32(v2))+200:], uint64(t380))
				t381 := int64(load64(m.memory[int64(uint32(v2))+96:]))
				store64(m.memory[int64(uint32(v2))+192:], uint64(t381))
				t382 := int64(load64(m.memory[int64(uint32(v2))+88:]))
				store64(m.memory[int64(uint32(v2))+184:], uint64(t382))
				t383 := int64(load64(m.memory[int64(uint32(v2))+80:]))
				store64(m.memory[int64(uint32(v2))+176:], uint64(t383))
				t384 := int64(load64(m.memory[int64(uint32(v2))+72:]))
				store64(m.memory[int64(uint32(v2))+168:], uint64(t384))
				t385 := int64(load64(m.memory[int64(uint32(v2))+64:]))
				store64(m.memory[int64(uint32(v2))+160:], uint64(t385))
				t386 := int64(load64(m.memory[int64(uint32(v2))+56:]))
				store64(m.memory[int64(uint32(v2))+152:], uint64(t386))
				m.fn336(v2+i32(132), v2+i32(152))
				t387 := int32(load32(m.memory[int64(uint32(v2))+140:]))
				v4 = t387
				if v4 != 0 {
					goto l204
				}
				store32(m.memory[uint32(v0):], uint32(i32(-1)))
				m.fn362(v2 + i32(132))
			}
		l186:
			if v29 == 0 {
				goto l205
			}
			v5 = v29 << 4
			v4 = v5 + v29 + i32(25)
			if v4 == 0 {
				goto l205
			}
			v1 = v21 - v5
			t388 := int32(load32(m.memory[uint32(v1+i32(-20)):]))
			v5 = t388
			v7 = v5 & i32(-8)
			t389 := v7
			v5 = v5 & i32(3)
			p390 := i32(8)
			if v5 != 0 {
				p390 = i32(4)
			}
			if uint32(t389) < uint32(p390+v4) {
				m.fn3(i32(1273840), i32(46), i32(1273888))
				panic("unreachable")
			}
			if v5 == 0 {
				goto l207
			}
			if uint32(v7) > uint32(v4+i32(39)) {
				m.fn3(i32(1273904), i32(46), i32(1273952))
				panic("unreachable")
			}
		l207:
			m.fn1(v1 + i32(-16))
			goto l205
		}
	l204:
		t391 := int32(load32(m.memory[int64(uint32(v2))+136:]))
		t392 := m.fn361(t391, v4, v8)
		v4 = t392
		store32(m.memory[uint32(v0):], uint32(i32(-0x7ffffffe)))
		t393 := int64(load64(m.memory[int64(uint32(v2))+132:]))
		store64(m.memory[int64(uint32(v0))+4:], uint64(t393))
		t394 := int32(load32(m.memory[int64(uint32(v2))+148:]))
		store32(m.memory[int64(uint32(v0))+20:], uint32(t394))
		store32(m.memory[int64(uint32(v2))+144:], uint32(v4))
		t395 := int64(load64(m.memory[int64(uint32(v2))+140:]))
		store64(m.memory[int64(uint32(v0))+12:], uint64(t395))
		if v29 == 0 {
			goto l205
		}
		v5 = v29 << 4
		v4 = v5 + v29 + i32(25)
		if v4 == 0 {
			goto l205
		}
		v1 = v21 - v5
		t396 := int32(load32(m.memory[uint32(v1+i32(-20)):]))
		v5 = t396
		v7 = v5 & i32(-8)
		t397 := v7
		v5 = v5 & i32(3)
		p398 := i32(8)
		if v5 != 0 {
			p398 = i32(4)
		}
		if uint32(t397) < uint32(p398+v4) {
			m.fn3(i32(1273840), i32(46), i32(1273888))
			panic("unreachable")
		}
		if v5 == 0 {
			goto l210
		}
		if uint32(v7) > uint32(v4+i32(39)) {
			m.fn3(i32(1273904), i32(46), i32(1273952))
			panic("unreachable")
		}
	l210:
		m.fn1(v1 + i32(-16))
	}
l205:
	{
		t399 := int32(load32(m.memory[int64(uint32(v2))+8:]))
		v4 = t399
		if v4 == 0 {
			goto l212
		}
		t400 := int32(load32(m.memory[int64(uint32(v2))+12:]))
		v1 = t400
		t401 := int32(load32(m.memory[uint32(v1+i32(-4)):]))
		v5 = t401
		v7 = v5 & i32(-8)
		t402 := v7
		v5 = v5 & i32(3)
		p403 := i32(8)
		if v5 != 0 {
			p403 = i32(4)
		}
		v4 = v4 << 3
		if uint32(t402) < uint32(p403+v4) {
			m.fn3(i32(1273840), i32(46), i32(1273888))
			panic("unreachable")
		}
		if v5 == 0 {
			goto l214
		}
		if uint32(v7) > uint32(v4+i32(39)) {
			m.fn3(i32(1273904), i32(46), i32(1273952))
			panic("unreachable")
		}
	l214:
		m.fn1(v1)
	}
l212:
	m.g0 = v2 + i32(208)
}
func (m *Module) fn567(v0, v1 int32) {
	var v2, v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14, v15, v16, v17, v18, v19, v20, v21, v22, v23, v24, v25, v26, v27, v28 int32
	var v29 int64
	var v30 int32
	var v31 int64
	var v32, v33, v34, v35 int32
	var v36 int64
	var v37 int32
	var v38 int64
	var v39, v40 int32
	var v41 int64
	var v42, v43, v44 int32
	var v45 int64
	var v46, v47 int32
	var v48 int64
	t0 := m.g0
	v2 = t0 - i32(144)
	m.g0 = v2
	{
		{
			t1 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			v3 = t1
			if v3 == 0 {
				store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
				store64(m.memory[uint32(v0):], uint64(i64(0x800000000)))
				m.fn443(v1)
				goto l40
			}
			v4 = v3 * i32(56)
			t2 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			v5 = t2
			t3 := int32(load32(m.memory[uint32(v5+i32(24)):]))
			v6 = t3
			{
				if v3 == i32(1) {
					goto l1
				}
				t4 := int32(uint32(v4+i32(-56)) / uint32(i32(56)))
				v7 = t4
				v8 = v7 & i32(3)
				v9 = i32(0)
				if uint32(v7+i32(-1)) < uint32(i32(3)) {
					goto l2
				}
				v3 = v5 + i32(248)
				v10 = v7 & i32(0x7fffffc)
				v9 = i32(0)
			l3:
				{
					t5 := int32(load32(m.memory[uint32(v3+i32(-168)):]))
					t6 := v6
					v7 = t5
					p7 := v7
					if uint32(v6) < uint32(v7) {
						p7 = t6
					}
					v7 = p7
					t8 := int32(load32(m.memory[uint32(v3+i32(-112)):]))
					t9 := v7
					v11 = t8
					p10 := v11
					if uint32(v7) < uint32(v11) {
						p10 = t9
					}
					v7 = p10
					t11 := int32(load32(m.memory[uint32(v3+i32(-56)):]))
					t12 := v7
					v11 = t11
					p13 := v11
					if uint32(v7) < uint32(v11) {
						p13 = t12
					}
					v7 = p13
					t14 := int32(load32(m.memory[uint32(v3):]))
					t15 := v7
					v11 = t14
					p16 := v11
					if uint32(v7) < uint32(v11) {
						p16 = t15
					}
					v6 = p16
					v3 = v3 + i32(224)
					t17 := v10
					v9 = v9 + i32(4)
					if t17 != v9 {
						goto l3
					}
				}
				if v8 == 0 {
					goto l1
				}
			l2:
				v3 = v9*i32(56) + v5 + i32(80)
			l4:
				{
					t18 := int32(load32(m.memory[uint32(v3):]))
					t19 := v6
					v9 = t18
					p20 := v9
					if uint32(v6) < uint32(v9) {
						p20 = t19
					}
					v6 = p20
					v3 = v3 + i32(56)
					v8 = v8 + i32(-1)
					if v8 != 0 {
						goto l4
					}
				}
			}
		l1:
			v12 = v5 + v4
			store32(m.memory[int64(uint32(v2))+12:], uint32(i32(0)))
			store64(m.memory[int64(uint32(v2))+4:], uint64(i64(0x800000000)))
			store32(m.memory[int64(uint32(v2))+24:], uint32(i32(-1)))
			v13 = v2 + i32(116) + i32(4)
			v14 = v2 + i32(88) + i32(12)
			v15 = v2 + i32(88) + i32(8)
			v16 = v2 + i32(16) + i32(12)
			t21 := int32(load32(m.memory[uint32(v1):]))
			v17 = t21
			v18 = i32(0)
			v19 = i32(8)
			v20 = i32(-2)
			v21 = v5
		l39:
			{
				if v20 != i32(-2) {
					goto l5
				}
				{
					if v21 != v12 {
						goto l6
					}
					v20 = i32(-1)
					v21 = v12
					goto l7
				l6:
					t22 := int32(load32(m.memory[int64(uint32(v21))+9:]))
					store32(m.memory[int64(uint32(v2))+116:], uint32(t22))
					t23 := int32(load32(m.memory[int64(uint32(v21))+12:]))
					store32(m.memory[int64(uint32(v2))+119:], uint32(t23))
					t24 := int32(load32(m.memory[int64(uint32(v21))+52:]))
					v22 = t24
					t25 := int32(load32(m.memory[int64(uint32(v21))+48:]))
					v23 = t25
					t26 := int32(load32(m.memory[int64(uint32(v21))+44:]))
					v24 = t26
					t27 := int32(load32(m.memory[int64(uint32(v21))+40:]))
					v20 = t27
					t28 := int32(load32(m.memory[int64(uint32(v21))+36:]))
					v25 = t28
					t29 := int32(load32(m.memory[int64(uint32(v21))+32:]))
					v26 = t29
					t30 := int32(load32(m.memory[int64(uint32(v21))+28:]))
					v27 = t30
					t31 := int32(load32(m.memory[int64(uint32(v21))+24:]))
					v28 = t31
					t32 := int64(load64(m.memory[int64(uint32(v21))+16:]))
					v29 = t32
					t33 := int32(m.memory[int64(uint32(v21))+8])
					v30 = t33
					t34 := int64(load64(m.memory[uint32(v21):]))
					v31 = t34
					v21 = v21 + i32(56)
				}
			l7:
				t35 := int32(load32(m.memory[int64(uint32(v2))+119:]))
				store32(m.memory[int64(uint32(v2))+71:], uint32(t35))
				t36 := int32(load32(m.memory[int64(uint32(v2))+116:]))
				store32(m.memory[int64(uint32(v2))+68:], uint32(t36))
				v32 = v22
				v33 = v23
				v34 = v24
				v3 = v25
				v8 = v26
				v7 = v27
				v35 = v28
				v36 = v29
				v37 = v30
				v38 = v31
			}
		l5:
			if v20 == i32(-1) {
				goto l8
			}
			{
				if uint32(v35) <= uint32(v6) {
					{
						{
							{
								t80 := int32(load32(m.memory[int64(uint32(v2))+24:]))
								v9 = t80
								if v9 == i32(-1) {
									goto l32
								}
								{
									t81 := int64(load64(m.memory[int64(uint32(v2))+40:]))
									if t81 != v38 {
										goto l33
									}
									t82 := int32(m.memory[int64(uint32(v2))+48])
									t83 := t82 & i32(255)
									v10 = v37 & i32(255)
									if t83 != v10 {
										goto l33
									}
									if v10 == 0 {
										goto l34
									}
									t84 := int64(load64(m.memory[int64(uint32(v2))+56:]))
									v48 = t84
									if v48 == i64(-1) {
										goto l33
									}
									if v48+i64(1) == v36 {
										goto l34
									}
								}
							l33:
								t85 := int32(load32(m.memory[int64(uint32(v16))+8:]))
								store32(m.memory[int64(uint32(v14))+8:], uint32(t85))
								t86 := int64(load64(m.memory[uint32(v16):]))
								store64(m.memory[uint32(v14):], uint64(t86))
								store32(m.memory[int64(uint32(v2))+96:], uint32(v9))
								t87 := int64(load64(m.memory[int64(uint32(v2))+16:]))
								store64(m.memory[int64(uint32(v2))+88:], uint64(t87))
								{
									t88 := int32(load32(m.memory[int64(uint32(v2))+104:]))
									if t88 != 0 {
										goto l35
									}
									m.fn575(v15)
									goto l32
								}
							l35:
								t89 := int64(load64(m.memory[int64(uint32(v2))+104:]))
								store64(m.memory[int64(uint32(v13))+16:], uint64(t89))
								t90 := int64(load64(m.memory[int64(uint32(v2))+96:]))
								store64(m.memory[int64(uint32(v13))+8:], uint64(t90))
								t91 := int64(load64(m.memory[int64(uint32(v2))+88:]))
								store64(m.memory[uint32(v13):], uint64(t91))
								{
									t92 := int32(load32(m.memory[int64(uint32(v2))+4:]))
									if v18 != t92 {
										goto l36
									}
									m.fn315(v2 + i32(4))
									t93 := int32(load32(m.memory[int64(uint32(v2))+8:]))
									v19 = t93
								}
							l36:
								v9 = v19 + v18<<5
								store32(m.memory[uint32(v9):], uint32(i32(-0x7fffffff)))
								t94 := int64(load64(m.memory[int64(uint32(v2))+116:]))
								store64(m.memory[int64(uint32(v9))+4:], uint64(t94))
								t95 := int64(load64(m.memory[int64(uint32(v2))+124:]))
								store64(m.memory[int64(uint32(v9))+12:], uint64(t95))
								t96 := int64(load64(m.memory[int64(uint32(v2))+132:]))
								store64(m.memory[int64(uint32(v9))+20:], uint64(t96))
								t97 := int32(load32(m.memory[int64(uint32(v2))+140:]))
								store32(m.memory[int64(uint32(v9))+28:], uint32(t97))
								t98 := v2
								v18 = v18 + i32(1)
								store32(m.memory[int64(uint32(t98))+12:], uint32(v18))
							}
						l32:
							m.memory[int64(uint32(v2))+48] = byte(v37)
							store64(m.memory[int64(uint32(v2))+40:], uint64(v38))
							m.memory[int64(uint32(v2))+36] = byte(v37)
							v46 = i32(0)
							store32(m.memory[int64(uint32(v2))+32:], uint32(i32(0)))
							store64(m.memory[int64(uint32(v2))+24:], uint64(i64(0x400000000)))
							store64(m.memory[int64(uint32(v2))+56:], uint64(v36))
							t100 := v2
							p99 := i64(1)
							if v37&i32(255) != 0 {
								p99 = v36
							}
							store64(m.memory[int64(uint32(t100))+16:], uint64(p99))
							goto l37
						}
					l34:
						if v46 != v9 {
							goto l38
						}
					l37:
						m.fn318(v2 + i32(16) + i32(8))
						t101 := int32(load32(m.memory[int64(uint32(v2))+28:]))
						v47 = t101
					}
				l38:
					v9 = v47 + v46*i32(28)
					m.memory[int64(uint32(v9))+24] = byte(i32(2))
					store32(m.memory[int64(uint32(v9))+20:], uint32(v3))
					store32(m.memory[int64(uint32(v9))+16:], uint32(v8))
					store32(m.memory[int64(uint32(v9))+12:], uint32(v7))
					store32(m.memory[int64(uint32(v9))+8:], uint32(v33))
					store32(m.memory[int64(uint32(v9))+4:], uint32(v34))
					store32(m.memory[uint32(v9):], uint32(v20))
					store64(m.memory[int64(uint32(v2))+56:], uint64(v36))
					t102 := v2
					v46 = v46 + i32(1)
					store32(m.memory[int64(uint32(t102))+32:], uint32(v46))
					v20 = i32(-2)
					goto l39
				}
				v9 = i32(0)
				store32(m.memory[int64(uint32(v2))+84:], uint32(i32(0)))
				store64(m.memory[int64(uint32(v2))+76:], uint64(i64(0x800000000)))
				m.fn324(v2 + i32(76))
				t37 := int32(load32(m.memory[int64(uint32(v2))+80:]))
				v10 = t37
				m.memory[int64(uint32(v10))+8] = byte(v37)
				store64(m.memory[uint32(v10):], uint64(v38))
				t38 := int32(load32(m.memory[int64(uint32(v2))+68:]))
				store32(m.memory[int64(uint32(v10))+9:], uint32(t38))
				t39 := int32(load32(m.memory[int64(uint32(v2))+71:]))
				store32(m.memory[int64(uint32(v10))+12:], uint32(t39))
				store32(m.memory[int64(uint32(v10))+52:], uint32(v32))
				store32(m.memory[int64(uint32(v10))+48:], uint32(v33))
				store32(m.memory[int64(uint32(v10))+44:], uint32(v34))
				store32(m.memory[int64(uint32(v10))+40:], uint32(v20))
				store32(m.memory[int64(uint32(v10))+36:], uint32(v3))
				store32(m.memory[int64(uint32(v10))+32:], uint32(v8))
				store32(m.memory[int64(uint32(v10))+28:], uint32(v7))
				store32(m.memory[int64(uint32(v10))+24:], uint32(v35))
				store64(m.memory[int64(uint32(v10))+16:], uint64(v36))
				v7 = i32(1)
				v33 = v21
			l31:
				store32(m.memory[int64(uint32(v2))+84:], uint32(v7))
				{
					v3 = v21 + v9
					if v3 != v12 {
						goto l10
					}
					v20 = i32(-1)
					v3 = v39
					v21 = v12
					goto l11
				l10:
					t40 := int32(m.memory[uint32(v3+i32(8))])
					v40 = t40
					t41 := v2
					v8 = v3 + i32(12)
					t42 := int32(load32(m.memory[uint32(v8):]))
					store32(m.memory[int64(uint32(t41))+71:], uint32(t42))
					t43 := v2
					v37 = v3 + i32(9)
					t44 := int32(load32(m.memory[uint32(v37):]))
					store32(m.memory[int64(uint32(t43))+68:], uint32(t44))
					v20 = i32(-1)
					t45 := int64(load64(m.memory[uint32(v3+i32(16)):]))
					v41 = t45
					t46 := int32(load32(m.memory[uint32(v3+i32(52)):]))
					v42 = t46
					t47 := int32(load32(m.memory[uint32(v3+i32(48)):]))
					v43 = t47
					t48 := int32(load32(m.memory[uint32(v3+i32(44)):]))
					v44 = t48
					t49 := int32(load32(m.memory[uint32(v3+i32(36)):]))
					v39 = t49
					t50 := int32(load32(m.memory[uint32(v3+i32(32)):]))
					v4 = t50
					t51 := int32(load32(m.memory[uint32(v3+i32(28)):]))
					v1 = t51
					t52 := int32(load32(m.memory[uint32(v3+i32(24)):]))
					v11 = t52
					t53 := int64(load64(m.memory[uint32(v3):]))
					v45 = t53
					t54 := int32(load32(m.memory[uint32(v3+i32(40)):]))
					v35 = t54
					if v35 != i32(-1) {
						goto l12
					}
					v21 = v33 + i32(56)
					v3 = v39
				}
			l11:
				v8 = v4
				v7 = v1
				goto l13
			l12:
				if uint32(v11) > uint32(v6) {
					t74 := int32(load32(m.memory[uint32(v8):]))
					store32(m.memory[int64(uint32(v2))+119:], uint32(t74))
					t75 := int32(load32(m.memory[uint32(v37):]))
					store32(m.memory[int64(uint32(v2))+116:], uint32(t75))
					{
						t76 := int32(load32(m.memory[int64(uint32(v2))+76:]))
						if v7 != t76 {
							goto l30
						}
						m.fn324(v2 + i32(76))
						t77 := int32(load32(m.memory[int64(uint32(v2))+80:]))
						v10 = t77
					}
				l30:
					v8 = v10 + v9
					m.memory[uint32(v8+i32(64))] = byte(v40)
					store64(m.memory[uint32(v8+i32(56)):], uint64(v45))
					store64(m.memory[uint32(v8+i32(72)):], uint64(v41))
					t78 := int32(load32(m.memory[int64(uint32(v2))+119:]))
					v37 = t78
					t79 := int32(load32(m.memory[int64(uint32(v2))+116:]))
					v20 = t79
					store32(m.memory[uint32(v8+i32(108)):], uint32(v42))
					store32(m.memory[uint32(v8+i32(104)):], uint32(v43))
					store32(m.memory[uint32(v8+i32(100)):], uint32(v44))
					store32(m.memory[uint32(v8+i32(96)):], uint32(v35))
					store32(m.memory[uint32(v8+i32(92)):], uint32(v39))
					store32(m.memory[uint32(v8+i32(88)):], uint32(v4))
					store32(m.memory[uint32(v8+i32(84)):], uint32(v1))
					store32(m.memory[uint32(v8+i32(80)):], uint32(v11))
					store32(m.memory[uint32(v8+i32(65)):], uint32(v20))
					store32(m.memory[uint32(v8+i32(68)):], uint32(v37))
					v33 = v3 + i32(56)
					v9 = v9 + i32(56)
					v7 = v7 + i32(1)
					goto l31
				}
				v21 = v3 + i32(56)
				v3 = v39
				v8 = v4
				v7 = v1
				v20 = v35
			l13:
				m.fn567(v2+i32(116), v2+i32(76))
				{
					t55 := int32(load32(m.memory[int64(uint32(v2))+124:]))
					v9 = t55
					if v9 != 0 {
						{
							{
								t61 := int32(load32(m.memory[int64(uint32(v2))+24:]))
								v10 = t61
								if v10 != i32(-1) {
									goto l20
								}
								store64(m.memory[int64(uint32(v2))+56:], uint64(i64(0)))
								m.memory[int64(uint32(v2))+48] = byte(i32(0))
								store64(m.memory[int64(uint32(v2))+40:], uint64(i64(-1)))
								m.memory[int64(uint32(v2))+36] = byte(i32(0))
								store32(m.memory[int64(uint32(v2))+32:], uint32(i32(0)))
								store64(m.memory[int64(uint32(v2))+24:], uint64(i64(0x400000000)))
								store64(m.memory[int64(uint32(v2))+16:], uint64(i64(1)))
								goto l21
							}
						l20:
							{
								if v46 == 0 {
									goto l22
								}
								t62 := int32(load32(m.memory[int64(uint32(v2))+28:]))
								v47 = t62
								goto l23
							}
						l22:
							if v10 != 0 {
								goto l24
							}
						l21:
							m.fn318(v2 + i32(16) + i32(8))
						l24:
							t63 := int32(load32(m.memory[int64(uint32(v2))+28:]))
							v47 = t63
							m.memory[int64(uint32(v47))+24] = byte(i32(2))
							store64(m.memory[int64(uint32(v47))+8:], uint64(i64(-0x100000000)))
							store64(m.memory[uint32(v47):], uint64(i64(0x800000000)))
							v46 = i32(1)
							store32(m.memory[int64(uint32(v2))+32:], uint32(i32(1)))
						}
					l23:
						v39 = v9 << 5
						t64 := int32(load32(m.memory[int64(uint32(v2))+116:]))
						v1 = t64
						t65 := int32(load32(m.memory[int64(uint32(v2))+120:]))
						v10 = t65
						{
							t66 := v9
							v35 = v47 + v46*i32(28)
							v33 = v35 + i32(-28)
							t67 := int32(load32(m.memory[uint32(v33):]))
							v37 = v35 + i32(-20)
							t68 := int32(load32(m.memory[uint32(v37):]))
							v4 = t68
							if uint32(t66) <= uint32(t67-v4) {
								goto l25
							}
							m.fn197(v33, v4, v9, i32(8), i32(32))
							t69 := int32(load32(m.memory[uint32(v37):]))
							v4 = t69
						}
					l25:
						{
							if v39 == 0 {
								goto l26
							}
							t70 := int32(load32(m.memory[uint32(v35+i32(-24)):]))
							memory_copy(m.memory, uint32(t70+v4<<5), uint32(v10), uint32(v39))
						}
					l26:
						store32(m.memory[uint32(v37):], uint32(v4+v9))
						if v1 == 0 {
							goto l16
						}
						t71 := int32(load32(m.memory[uint32(v10+i32(-4)):]))
						v9 = t71
						v4 = v9 & i32(-8)
						t72 := v4
						v9 = v9 & i32(3)
						p73 := i32(8)
						if v9 != 0 {
							p73 = i32(4)
						}
						v1 = v1 << 5
						if uint32(t72) < uint32(p73|v1) {
							m.fn3(i32(1273840), i32(46), i32(1273888))
							panic("unreachable")
						}
						if v9 == 0 {
							goto l28
						}
						if uint32(v4) > uint32(v1+i32(39)) {
							m.fn3(i32(1273904), i32(46), i32(1273952))
							panic("unreachable")
						}
					l28:
						m.fn1(v10)
						goto l16
					}
					t56 := int32(load32(m.memory[int64(uint32(v2))+116:]))
					v9 = t56
					if v9 == 0 {
						goto l16
					}
					t57 := int32(load32(m.memory[int64(uint32(v2))+120:]))
					v1 = t57
					t58 := int32(load32(m.memory[uint32(v1+i32(-4)):]))
					v10 = t58
					v4 = v10 & i32(-8)
					t59 := v4
					v10 = v10 & i32(3)
					p60 := i32(8)
					if v10 != 0 {
						p60 = i32(4)
					}
					v9 = v9 << 5
					if uint32(t59) < uint32(p60|v9) {
						m.fn3(i32(1273840), i32(46), i32(1273888))
						panic("unreachable")
					}
					if v10 == 0 {
						goto l18
					}
					if uint32(v4) > uint32(v9+i32(39)) {
						m.fn3(i32(1273904), i32(46), i32(1273952))
						panic("unreachable")
					}
				l18:
					m.fn1(v1)
					goto l16
				}
			}
		l16:
			v32 = v42
			v33 = v43
			v34 = v44
			v35 = v11
			v36 = v41
			v37 = v40
			v38 = v45
			v1 = v7
			v4 = v8
			v39 = v3
			goto l39
		}
	l8:
		t103 := int32(load32(m.memory[int64(uint32(v2))+24:]))
		v3 = t103
		store32(m.memory[int64(uint32(v2))+24:], uint32(i32(-1)))
		{
			if v3 == i32(-1) {
				goto l41
			}
			t104 := int64(load64(m.memory[int64(uint32(v2))+16:]))
			v45 = t104
			t105 := int64(load64(m.memory[uint32(v16):]))
			store64(m.memory[int64(uint32(v2))+100:], uint64(t105))
			store32(m.memory[int64(uint32(v2))+96:], uint32(v3))
			store64(m.memory[int64(uint32(v2))+88:], uint64(v45))
			t106 := int32(load32(m.memory[int64(uint32(v16))+8:]))
			store32(m.memory[int64(uint32(v2))+108:], uint32(t106))
			{
				t107 := int32(load32(m.memory[int64(uint32(v2))+104:]))
				if t107 != 0 {
					goto l42
				}
				m.fn575(v2 + i32(96))
				goto l41
			}
		l42:
			t108 := int64(load64(m.memory[int64(uint32(v2))+104:]))
			store64(m.memory[int64(uint32(v2))+136:], uint64(t108))
			t109 := int64(load64(m.memory[int64(uint32(v2))+96:]))
			store64(m.memory[int64(uint32(v2))+128:], uint64(t109))
			t110 := int64(load64(m.memory[int64(uint32(v2))+88:]))
			store64(m.memory[int64(uint32(v2))+120:], uint64(t110))
			{
				t111 := int32(load32(m.memory[int64(uint32(v2))+12:]))
				v8 = t111
				t112 := int32(load32(m.memory[int64(uint32(v2))+4:]))
				if v8 != t112 {
					goto l43
				}
				m.fn315(v2 + i32(4))
			}
		l43:
			t113 := int32(load32(m.memory[int64(uint32(v2))+8:]))
			v3 = t113 + v8<<5
			store32(m.memory[uint32(v3):], uint32(i32(-0x7fffffff)))
			t114 := int64(load64(m.memory[int64(uint32(v2))+116:]))
			store64(m.memory[int64(uint32(v3))+4:], uint64(t114))
			t115 := int64(load64(m.memory[int64(uint32(v2))+124:]))
			store64(m.memory[int64(uint32(v3))+12:], uint64(t115))
			t116 := int64(load64(m.memory[int64(uint32(v2))+132:]))
			store64(m.memory[int64(uint32(v3))+20:], uint64(t116))
			t117 := int32(load32(m.memory[int64(uint32(v2))+140:]))
			store32(m.memory[int64(uint32(v3))+28:], uint32(t117))
			store32(m.memory[int64(uint32(v2))+12:], uint32(v8+i32(1)))
		}
	l41:
		t118 := int32(load32(m.memory[int64(uint32(v2))+12:]))
		store32(m.memory[int64(uint32(v0))+8:], uint32(t118))
		t119 := int64(load64(m.memory[int64(uint32(v2))+4:]))
		store64(m.memory[uint32(v0):], uint64(t119))
		t120 := int32(uint32(v12-v21) / uint32(i32(56)))
		v10 = t120
		if v12 == v21 {
			goto l44
		}
		v7 = i32(0)
	l55:
		{
			{
				v9 = v21 + v7*i32(56)
				t121 := int32(load32(m.memory[int64(uint32(v9))+28:]))
				v3 = t121
				if v3 == i32(-1) {
					goto l45
				}
				if v3 == 0 {
					goto l45
				}
				t122 := int32(load32(m.memory[int64(uint32(v9))+32:]))
				v11 = t122
				t123 := int32(load32(m.memory[uint32(v11+i32(-4)):]))
				v8 = t123
				v1 = v8 & i32(-8)
				t124 := v1
				v8 = v8 & i32(3)
				p125 := i32(8)
				if v8 != 0 {
					p125 = i32(4)
				}
				if uint32(t124) < uint32(p125+v3) {
					m.fn3(i32(1273840), i32(46), i32(1273888))
					panic("unreachable")
				}
				if v8 == 0 {
					goto l47
				}
				if uint32(v1) > uint32(v3+i32(39)) {
					m.fn3(i32(1273904), i32(46), i32(1273952))
					panic("unreachable")
				}
			l47:
				m.fn1(v11)
			}
		l45:
			t126 := int32(load32(m.memory[int64(uint32(v9))+44:]))
			v11 = t126
			{
				t127 := int32(load32(m.memory[int64(uint32(v9))+48:]))
				v8 = t127
				if v8 == 0 {
					goto l49
				}
				v3 = v11
			l50:
				m.fn335(v3)
				v3 = v3 + i32(32)
				v8 = v8 + i32(-1)
				if v8 != 0 {
					goto l50
				}
			}
		l49:
			{
				t128 := int32(load32(m.memory[int64(uint32(v9))+40:]))
				v3 = t128
				if v3 == 0 {
					goto l51
				}
				t129 := int32(load32(m.memory[uint32(v11+i32(-4)):]))
				v8 = t129
				v9 = v8 & i32(-8)
				t130 := v9
				v8 = v8 & i32(3)
				p131 := i32(8)
				if v8 != 0 {
					p131 = i32(4)
				}
				v3 = v3 << 5
				if uint32(t130) < uint32(p131|v3) {
					m.fn3(i32(1273840), i32(46), i32(1273888))
					panic("unreachable")
				}
				if v8 == 0 {
					goto l53
				}
				if uint32(v9) > uint32(v3+i32(39)) {
					m.fn3(i32(1273904), i32(46), i32(1273952))
					panic("unreachable")
				}
			l53:
				m.fn1(v11)
			}
		l51:
			v7 = v7 + i32(1)
			if v7 != v10 {
				goto l55
			}
		}
	l44:
		if v17 == 0 {
			goto l40
		}
		t132 := int32(load32(m.memory[uint32(v5+i32(-4)):]))
		v3 = t132
		v8 = v3 & i32(-8)
		t133 := v8
		v3 = v3 & i32(3)
		p134 := i32(8)
		if v3 != 0 {
			p134 = i32(4)
		}
		v9 = v17 * i32(56)
		if uint32(t133) < uint32(p134+v9) {
			m.fn3(i32(1273840), i32(46), i32(1273888))
			panic("unreachable")
		}
		if v3 == 0 {
			goto l57
		}
		if uint32(v8) > uint32(v9+i32(39)) {
			m.fn3(i32(1273904), i32(46), i32(1273952))
			panic("unreachable")
		}
	l57:
		m.fn1(v5)
	}
l40:
	m.g0 = v2 + i32(144)
}
func (m *Module) fn568(v0 int32) {
	var v1 int32
	{
		t0 := m.fn5(i32(0x8000))
		v1 = t0
		if v1 == 0 {
			m.fn10(i32(1), i32(0x8000))
			panic("unreachable")
		}
		{
			t1 := int32(m.memory[uint32(v1+i32(-4))])
			if t1&i32(3) == 0 {
				goto l1
			}
			memory_zero(m.memory, uint32(v1), uint32(i32(0x8000)))
		}
	l1:
		store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0x8000)))
		store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
		store32(m.memory[uint32(v0):], uint32(i32(0x8000)))
		return
	}
}
func (m *Module) fn569(v0, v1, v2, v3 int32) {
	var v4, v5, v6, v7 int32
	{
		{
			t0 := int32(load32(m.memory[uint32(v3):]))
			v4 = t0
			t1 := int32(load32(m.memory[int64(uint32(v3))+8:]))
			t2 := v4
			v5 = t1
			if uint32(t2) > uint32(v5) {
				goto l0
			}
			t3 := int32(load32(m.memory[int64(uint32(v3))+4:]))
			v3 = t3
			goto l1
		}
	l0:
		t4 := int32(load32(m.memory[int64(uint32(v3))+4:]))
		v3 = t4
		{
			if v5 != 0 {
				goto l2
			}
			t5 := int32(load32(m.memory[uint32(v3+i32(-4)):]))
			v6 = t5
			v7 = v6 & i32(-8)
			t6 := v7
			v6 = v6 & i32(3)
			p7 := i32(8)
			if v6 != 0 {
				p7 = i32(4)
			}
			if uint32(t6) < uint32(p7+v4) {
				m.fn3(i32(1273840), i32(46), i32(1273888))
				panic("unreachable")
			}
			if v6 == 0 {
				goto l4
			}
			if uint32(v7) > uint32(v4+i32(39)) {
				m.fn3(i32(1273904), i32(46), i32(1273952))
				panic("unreachable")
			}
		l4:
			m.fn1(v3)
			v3 = i32(1)
			goto l1
		}
	l2:
		t8 := m.fn22(v3, v4, i32(1), v5)
		v3 = t8
		if v3 == 0 {
			m.fn10(i32(1), v5)
			panic("unreachable")
		}
	}
l1:
	m.fn257(v0 + i32(24))
	store64(m.memory[int64(uint32(v0))+16:], uint64(i64(0)))
	store32(m.memory[int64(uint32(v0))+12:], uint32(v5))
	store32(m.memory[int64(uint32(v0))+8:], uint32(v3))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v2))
	store32(m.memory[uint32(v0):], uint32(v1))
}
func (m *Module) fn570(v0, v1, v2 int32) {
	var v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14 int32
	var v15 int64
	var v16, v17, v18, v19, v20 int32
	t0 := m.g0
	v3 = t0 - i32(16)
	m.g0 = v3
	t1 := int32(load32(m.memory[uint32(v2):]))
	v4 = t1
	v5 = v4
	t2 := int32(load32(m.memory[int64(uint32(v2))+8:]))
	v6 = t2
	v7 = v6
	if uint32(v4-v6) >= uint32(i32(32)) {
		goto l0
	}
	m.fn573(v3, v1, v2)
	{
		t3 := int32(m.memory[uint32(v3)])
		if t3 == i32(255) {
			t5 := int32(load32(m.memory[int64(uint32(v3))+4:]))
			if t5 == 0 {
				m.memory[uint32(v0)] = byte(i32(255))
				store32(m.memory[int64(uint32(v0))+4:], uint32(i32(0)))
				goto l2
			}
			t6 := int32(load32(m.memory[uint32(v2):]))
			v5 = t6
			t7 := int32(load32(m.memory[int64(uint32(v2))+8:]))
			v7 = t7
			goto l0
		}
		t4 := int64(load64(m.memory[uint32(v3):]))
		store64(m.memory[uint32(v0):], uint64(t4))
		goto l2
	}
l0:
	v8 = i32(8192)
l42:
	{
		if v7 != v5 {
			goto l4
		}
		if v5 != v4 {
			goto l4
		}
		m.fn573(v3, v1, v2)
		{
			t8 := int32(m.memory[uint32(v3)])
			if t8 == i32(255) {
				t10 := int32(load32(m.memory[int64(uint32(v2))+8:]))
				v7 = t10
				{
					t11 := int32(load32(m.memory[int64(uint32(v3))+4:]))
					if t11 == 0 {
						m.memory[uint32(v0)] = byte(i32(255))
						store32(m.memory[int64(uint32(v0))+4:], uint32(v7-v6))
						goto l2
					}
					t12 := int32(load32(m.memory[uint32(v2):]))
					v5 = t12
					goto l4
				}
			}
			t9 := int64(load64(m.memory[uint32(v3):]))
			store64(m.memory[uint32(v0):], uint64(t9))
			goto l2
		}
	l4:
		t13 := int32(load32(m.memory[int64(uint32(v2))+4:]))
		v9 = t13
		{
			{
				if v7 == v5 {
					goto l7
				}
				t14 := int32(load32(m.memory[uint32(v2):]))
				v5 = t14
				goto l8
			}
		l7:
			t15 := v3
			t16 := v5
			t17 := v9
			v10 = v5 + i32(32)
			t18 := v10
			v11 = v5 << 1
			p19 := v11
			if uint32(v10) > uint32(v11) {
				p19 = t18
			}
			v10 = p19
			m.fn208(t15, t16, t17, v10, i32(1), i32(1))
			t20 := int32(load32(m.memory[uint32(v3):]))
			if t20 != 0 {
				goto l9
			}
			t21 := int32(load32(m.memory[int64(uint32(v3))+4:]))
			v9 = t21
			store32(m.memory[uint32(v2):], uint32(v10))
			store32(m.memory[int64(uint32(v2))+4:], uint32(v9))
			v5 = v10
		}
	l8:
		t22 := v8
		v12 = v5 - v7
		p23 := v12
		if uint32(v8) < uint32(v12) {
			p23 = t22
		}
		v13 = p23
		v14 = v9 + v7
		t24 := int64(load64(m.memory[int64(uint32(v1))+120:]))
		v15 = t24
		v9 = i32(0)
		v16 = i32(0)
	l31:
		if v15 != i64(0) {
			goto l10
		}
		v17 = v17 | i32(255)
		v15 = i64(0)
		v10 = v9
		goto l11
	l10:
		v10 = v14 + v9
		{
			{
				t25 := v15
				v11 = v13 - v9
				if uint64(t25) < uint64(uint32(v11)) {
					v19 = int32(v15)
					if v16&i32(1) != 0 {
						m.fn574(v3, v1, v10, v19)
						{
							t36 := int32(m.memory[uint32(v3)])
							if t36 == i32(255) {
								t39 := int32(load32(m.memory[int64(uint32(v3))+4:]))
								v16 = t39
								if uint32(v16) > uint32(v19) {
									m.fn3(i32(1068762), i32(36), i32(1068800))
									panic("unreachable")
								}
								v17 = v17 | i32(255)
								goto l24
							}
							t37 := int32(load32(m.memory[int64(uint32(v3))+4:]))
							v18 = t37
							t38 := int32(load32(m.memory[uint32(v3):]))
							v17 = t38
							v16 = i32(0)
							goto l24
						}
					}
					if v19 == 0 {
						goto l19
					}
					memory_zero(m.memory, uint32(v10), uint32(v19))
				l19:
					m.fn574(v3, v1, v10, v19)
					{
						t32 := int32(m.memory[uint32(v3)])
						if t32 == i32(255) {
							t35 := int32(load32(m.memory[int64(uint32(v3))+4:]))
							v16 = t35
							if uint32(v16) > uint32(v19) {
								m.fn3(i32(1068762), i32(36), i32(1068800))
								panic("unreachable")
							}
							v17 = v17 | i32(255)
							goto l21
						}
						t33 := int32(load32(m.memory[int64(uint32(v3))+4:]))
						v18 = t33
						t34 := int32(load32(m.memory[uint32(v3):]))
						v17 = t34
						v16 = i32(0)
						goto l21
					}
				}
				if v16&i32(1) != 0 {
					goto l13
				}
				if v11 == 0 {
					goto l13
				}
				memory_zero(m.memory, uint32(v10), uint32(v11))
			l13:
				m.fn574(v3, v1, v10, v11)
				{
					{
						t26 := int32(m.memory[uint32(v3)])
						if t26 == i32(255) {
							goto l14
						}
						t27 := int32(load32(m.memory[int64(uint32(v3))+4:]))
						v18 = t27
						t28 := int32(load32(m.memory[uint32(v3):]))
						v17 = t28
						v10 = v9
						goto l15
					}
				l14:
					t29 := int32(load32(m.memory[int64(uint32(v3))+4:]))
					v10 = t29
					if uint32(v10) > uint32(v11) {
						m.fn3(i32(1068762), i32(36), i32(1068800))
						panic("unreachable")
					}
					v17 = v17 | i32(255)
					v10 = v10 + v9
				}
			l15:
				t30 := int64(load64(m.memory[int64(uint32(v1))+120:]))
				t31 := v1
				v15 = t30 - int64(uint32(v10-v9))
				store64(m.memory[int64(uint32(t31))+120:], uint64(v15))
				goto l17
			}
		l21:
			v11 = v11 - v19
			if v11 == 0 {
				goto l24
			}
			memory_zero(m.memory, uint32(v10+v19), uint32(v11))
		l24:
			t40 := int64(load64(m.memory[int64(uint32(v1))+120:]))
			t41 := v1
			v15 = t40 - int64(uint32(v16))
			store64(m.memory[int64(uint32(t41))+120:], uint64(v15))
			v10 = v16 + v9
		}
	l17:
		v16 = i32(1)
	l11:
		switch v17 & i32(255) {
		case 0:
			goto l26
		default:
			t54 := v2
			v7 = v10 + v7
			store32(m.memory[int64(uint32(t54))+8:], uint32(v7))
			if v10 != 0 {
				if v16&i32(1) != 0 {
					if uint32(v12) < uint32(v8) {
						goto l42
					}
					if v10 != v13 {
						goto l42
					}
					var p55 int32
					if v8 < i32(0) {
						p55 = 1
					}
					v9 = p55
					v8 = v8 << 1
					if v9 == 0 {
						goto l42
					}
					v8 = i32(-1)
					goto l42
				}
				v8 = i32(-1)
				goto l42
			}
			m.memory[uint32(v0)] = byte(i32(255))
			store32(m.memory[int64(uint32(v0))+4:], uint32(v7-v6))
			goto l2
		case 1:
			v9 = v10
			if v17&i32(0xff00) == i32(8960) {
				goto l31
			}
			goto l26
		case 3:
			t42 := int32(m.memory[int64(uint32(v18))+8])
			if t42 != i32(35) {
				goto l26
			}
			t43 := int32(load32(m.memory[uint32(v18):]))
			v9 = t43
			{
				t44 := int32(load32(m.memory[uint32(v18+i32(4)):]))
				v11 = t44
				t45 := int32(load32(m.memory[uint32(v11):]))
				v19 = t45
				if v19 == 0 {
					goto l32
				}
				m.t0[uint(v19)].(func(int32))(v9)
			}
		l32:
			{
				t46 := int32(load32(m.memory[int64(uint32(v11))+4:]))
				v11 = t46
				if v11 == 0 {
					goto l33
				}
				t47 := int32(load32(m.memory[uint32(v9+i32(-4)):]))
				v19 = t47
				v20 = v19 & i32(-8)
				t48 := v20
				v19 = v19 & i32(3)
				p49 := i32(8)
				if v19 != 0 {
					p49 = i32(4)
				}
				if uint32(t48) < uint32(p49+v11) {
					goto l34
				}
				if v19 == 0 {
					goto l35
				}
				if uint32(v20) > uint32(v11+i32(39)) {
					m.fn3(i32(1273904), i32(46), i32(1273952))
					panic("unreachable")
				}
			l35:
				m.fn1(v9)
			}
		l33:
			t50 := int32(load32(m.memory[uint32(v18+i32(-4)):]))
			v9 = t50
			v11 = v9 & i32(-8)
			t51 := v11
			v9 = v9 & i32(3)
			p52 := i32(20)
			if v9 != 0 {
				p52 = i32(16)
			}
			if uint32(t51) < uint32(p52) {
				m.fn3(i32(1273840), i32(46), i32(1273888))
				panic("unreachable")
			}
			if v9 == 0 {
				goto l38
			}
			if uint32(v11) >= uint32(i32(52)) {
				m.fn3(i32(1273904), i32(46), i32(1273952))
				panic("unreachable")
			}
		l38:
			m.fn1(v18)
			v9 = v10
			goto l31
		case 2:
			v9 = v10
			t53 := int32(m.memory[int64(uint32(v18))+8])
			if t53 == i32(35) {
				goto l31
			}
		}
	l26:
		store32(m.memory[int64(uint32(v0))+4:], uint32(v18))
		store32(m.memory[uint32(v0):], uint32(v17))
		store32(m.memory[int64(uint32(v2))+8:], uint32(v10+v7))
		goto l2
	l34:
	}
	m.fn3(i32(1273840), i32(46), i32(1273888))
	panic("unreachable")
l9:
	store64(m.memory[uint32(v0):], uint64(i64(9729)))
l2:
	m.g0 = v3 + i32(16)
}
func (m *Module) fn571(v0 int32) {
	var v1, v2, v3, v4 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
		v1 = t0
		if v1 == 0 {
			goto l0
		}
		t1 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		v2 = t1
		t2 := int32(load32(m.memory[uint32(v2+i32(-4)):]))
		v3 = t2
		v4 = v3 & i32(-8)
		t3 := v4
		v3 = v3 & i32(3)
		p4 := i32(8)
		if v3 != 0 {
			p4 = i32(4)
		}
		if uint32(t3) < uint32(p4+v1) {
			m.fn3(i32(1273840), i32(46), i32(1273888))
			panic("unreachable")
		}
		if v3 == 0 {
			goto l2
		}
		if uint32(v4) > uint32(v1+i32(39)) {
			m.fn3(i32(1273904), i32(46), i32(1273952))
			panic("unreachable")
		}
	l2:
		m.fn1(v2)
	}
l0:
	m.fn260(v0 + i32(40))
}
func (m *Module) fn572(v0, v1 int32) {
	var v2, v3, v4 int32
	{
		if v0&i32(255) != i32(3) {
			return
		}
		t0 := int32(load32(m.memory[uint32(v1):]))
		v0 = t0
		{
			t1 := int32(load32(m.memory[uint32(v1+i32(4)):]))
			v2 = t1
			t2 := int32(load32(m.memory[uint32(v2):]))
			v3 = t2
			if v3 == 0 {
				goto l1
			}
			m.t0[uint(v3)].(func(int32))(v0)
		}
	l1:
		{
			t3 := int32(load32(m.memory[int64(uint32(v2))+4:]))
			v2 = t3
			if v2 == 0 {
				goto l2
			}
			t4 := int32(load32(m.memory[uint32(v0+i32(-4)):]))
			v3 = t4
			v4 = v3 & i32(-8)
			t5 := v4
			v3 = v3 & i32(3)
			p6 := i32(8)
			if v3 != 0 {
				p6 = i32(4)
			}
			if uint32(t5) < uint32(p6+v2) {
				m.fn3(i32(1273840), i32(46), i32(1273888))
				panic("unreachable")
			}
			if v3 == 0 {
				goto l4
			}
			if uint32(v4) > uint32(v2+i32(39)) {
				m.fn3(i32(1273904), i32(46), i32(1273952))
				panic("unreachable")
			}
		l4:
			m.fn1(v0)
		}
	l2:
		t7 := int32(load32(m.memory[uint32(v1+i32(-4)):]))
		v0 = t7
		v2 = v0 & i32(-8)
		t8 := v2
		v0 = v0 & i32(3)
		p9 := i32(20)
		if v0 != 0 {
			p9 = i32(16)
		}
		if uint32(t8) < uint32(p9) {
			m.fn3(i32(1273840), i32(46), i32(1273888))
			panic("unreachable")
		}
		if v0 == 0 {
			goto l7
		}
		if uint32(v2) >= uint32(i32(52)) {
			m.fn3(i32(1273904), i32(46), i32(1273952))
			panic("unreachable")
		}
	l7:
		m.fn1(v1)
	}
}
func (m *Module) fn573(v0, v1, v2 int32) {
	var v3 int32
	var v4 int64
	var v5, v6 int32
	var v7 int64
	var v8, v9, v10, v11 int32
	t0 := m.g0
	v3 = t0 - i32(48)
	m.g0 = v3
	store64(m.memory[int64(uint32(v3))+32:], uint64(i64(0)))
	store64(m.memory[int64(uint32(v3))+24:], uint64(i64(0)))
	store64(m.memory[int64(uint32(v3))+16:], uint64(i64(0)))
	store64(m.memory[int64(uint32(v3))+8:], uint64(i64(0)))
	{
		{
			{
				t1 := int64(load64(m.memory[int64(uint32(v1))+120:]))
				v4 = t1
				if v4 == 0 {
					goto l0
				}
			l19:
				{
					t3 := v3 + i32(40)
					t4 := v1
					t5 := v3 + i32(8)
					p2 := i64(32)
					if uint64(v4) < uint64(i64(32)) {
						p2 = v4
					}
					m.fn574(t3, t4, t5, int32(p2))
					{
						{
							t6 := int32(m.memory[int64(uint32(v3))+40])
							if t6 == i32(255) {
								goto l1
							}
							t7 := int32(load32(m.memory[int64(uint32(v3))+44:]))
							v5 = t7
							t8 := int32(load32(m.memory[int64(uint32(v3))+40:]))
							v6 = t8
							goto l2
						}
					l1:
						t9 := int64(load64(m.memory[int64(uint32(v1))+120:]))
						v4 = t9
						t10 := int32(load32(m.memory[int64(uint32(v3))+44:]))
						t11 := v4
						v5 = t10
						v7 = int64(uint32(v5))
						if uint64(t11) < uint64(v7) {
							m.fn28(i32(1080060), i32(69), i32(1080096))
							panic("unreachable")
						}
						store64(m.memory[int64(uint32(v1))+120:], uint64(v4-v7))
						v6 = v6 | i32(255)
					}
				l2:
					switch v6 & i32(255) {
					case 0:
						goto l4
					case 1:
						if v6&i32(0xff00) != i32(8960) {
							goto l4
						}
						goto l9
					default:
						if uint32(v5) < uint32(i32(33)) {
							goto l10
						}
						m.fn121(i32(0), v5, i32(32), i32(1069588))
						panic("unreachable")
					case 2:
						t12 := int32(m.memory[int64(uint32(v5))+8])
						if t12 == i32(35) {
							goto l9
						}
						goto l4
					case 3:
						t13 := int32(m.memory[int64(uint32(v5))+8])
						if t13 != i32(35) {
							goto l4
						}
						t14 := int32(load32(m.memory[uint32(v5):]))
						v8 = t14
						{
							t15 := int32(load32(m.memory[uint32(v5+i32(4)):]))
							v9 = t15
							t16 := int32(load32(m.memory[uint32(v9):]))
							v10 = t16
							if v10 == 0 {
								goto l11
							}
							m.t0[uint(v10)].(func(int32))(v8)
						}
					l11:
						{
							t17 := int32(load32(m.memory[int64(uint32(v9))+4:]))
							v9 = t17
							if v9 == 0 {
								goto l12
							}
							t18 := int32(load32(m.memory[uint32(v8+i32(-4)):]))
							v10 = t18
							v11 = v10 & i32(-8)
							t19 := v11
							v10 = v10 & i32(3)
							p20 := i32(8)
							if v10 != 0 {
								p20 = i32(4)
							}
							if uint32(t19) < uint32(p20+v9) {
								m.fn3(i32(1273840), i32(46), i32(1273888))
								panic("unreachable")
							}
							if v10 == 0 {
								goto l14
							}
							if uint32(v11) > uint32(v9+i32(39)) {
								m.fn3(i32(1273904), i32(46), i32(1273952))
								panic("unreachable")
							}
						l14:
							m.fn1(v8)
						}
					l12:
						t21 := int32(load32(m.memory[uint32(v5+i32(-4)):]))
						v8 = t21
						v9 = v8 & i32(-8)
						t22 := v9
						v8 = v8 & i32(3)
						p23 := i32(20)
						if v8 != 0 {
							p23 = i32(16)
						}
						if uint32(t22) < uint32(p23) {
							m.fn3(i32(1273840), i32(46), i32(1273888))
							panic("unreachable")
						}
						if v8 == 0 {
							goto l17
						}
						if uint32(v9) >= uint32(i32(52)) {
							m.fn3(i32(1273904), i32(46), i32(1273952))
							panic("unreachable")
						}
					l17:
						m.fn1(v5)
					}
				l9:
					t24 := int64(load64(m.memory[int64(uint32(v1))+120:]))
					v4 = t24
					if !(v4 == 0) {
						goto l19
					}
				}
			}
		l0:
			v6 = v2 + i32(8)
			t25 := int32(load32(m.memory[int64(uint32(v2))+8:]))
			v1 = t25
			goto l20
		}
	l4:
		store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
		store32(m.memory[uint32(v0):], uint32(v6))
		goto l21
	l10:
		v6 = v2 + i32(8)
		{
			t26 := int32(load32(m.memory[uint32(v2):]))
			t27 := int32(load32(m.memory[int64(uint32(v2))+8:]))
			t28 := v5
			v1 = t27
			if uint32(t28) <= uint32(t26-v1) {
				goto l22
			}
			m.fn197(v2, v1, v5, i32(1), i32(1))
			t29 := int32(load32(m.memory[int64(uint32(v2))+8:]))
			v1 = t29
			goto l23
		}
	l22:
		if v5 != 0 {
			goto l23
		}
	l20:
		v5 = i32(0)
		goto l24
	l23:
		if v5 == 0 {
			goto l24
		}
		t30 := int32(load32(m.memory[int64(uint32(v2))+4:]))
		memory_copy(m.memory, uint32(t30+v1), uint32(v3+i32(8)), uint32(v5))
	}
l24:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
	m.memory[uint32(v0)] = byte(i32(255))
	store32(m.memory[uint32(v6):], uint32(v1+v5))
l21:
	m.g0 = v3 + i32(48)
}
func (m *Module) fn574(v0, v1, v2, v3 int32) {
	var v4, v5, v6, v7, v8, v9, v10, v11 int32
	var v12, v13 int64
	var v14, v15, v16 int32
	t0 := m.g0
	v4 = t0 - i32(16)
	m.g0 = v4
	v5 = v1 + i32(24)
	t1 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	v6 = t1
	t2 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	v7 = t2
	t3 := int32(load32(m.memory[int64(uint32(v1))+12:]))
	v8 = t3
	t4 := int32(load32(m.memory[uint32(v1):]))
	v9 = t4
	t5 := int32(load32(m.memory[int64(uint32(v1))+20:]))
	v10 = t5
	t6 := int32(load32(m.memory[int64(uint32(v1))+16:]))
	v11 = t6
l9:
	{
		{
			if v11 != v10 {
				goto l0
			}
			{
				p7 := v8
				if uint32(v6) < uint32(v8) {
					p7 = v6
				}
				v10 = p7
				if v10 != i32(1) {
					goto l1
				}
				t8 := int32(m.memory[uint32(v9)])
				m.memory[uint32(v7)] = byte(t8)
				goto l2
			}
		l1:
			if v10 == 0 {
				goto l2
			}
			memory_copy(m.memory, uint32(v7), uint32(v9), uint32(v10))
		l2:
			store32(m.memory[int64(uint32(v1))+20:], uint32(v10))
			t9 := v1
			v6 = v6 - v10
			store32(m.memory[int64(uint32(t9))+4:], uint32(v6))
			t10 := v1
			v9 = v9 + v10
			store32(m.memory[uint32(t10):], uint32(v9))
			v11 = i32(0)
		}
	l0:
		if uint32(v10) > uint32(v8) {
			goto l3
		}
		if uint32(v10) < uint32(v11) {
			goto l3
		}
		t11 := int64(load64(m.memory[int64(uint32(v1))+32:]))
		v12 = t11
		t12 := int64(load64(m.memory[int64(uint32(v1))+24:]))
		v13 = t12
		t13 := v4 + i32(4)
		t14 := v5
		t15 := v7 + v11
		t16 := v10 - v11
		t17 := v2
		t18 := v3
		var p19 int32
		if v10 == v11 {
			p19 = 1
		}
		v14 = p19
		p20 := i32(0)
		if v14 != 0 {
			p20 = i32(4)
		}
		m.fn269(t13, t14, t15, t16, t17, t18, p20)
		t21 := int32(m.memory[int64(uint32(v4))+8])
		v15 = t21
		t22 := int32(load32(m.memory[int64(uint32(v4))+4:]))
		v16 = t22
		t23 := int64(load64(m.memory[int64(uint32(v1))+24:]))
		t24 := v1
		t25 := v10
		v11 = v11 + int32(t23-v13)
		p26 := v11
		if uint32(v10) < uint32(v11) {
			p26 = t25
		}
		v11 = p26
		store32(m.memory[int64(uint32(t24))+16:], uint32(v11))
		if v16 == i32(2) {
			goto l4
		}
		m.fn267(v0, i32(20), i32(1069276), i32(22))
		goto l5
	l4:
		t27 := int64(load64(m.memory[int64(uint32(v1))+32:]))
		v16 = int32(t27 - v12)
		switch v15 {
		case 2:
			goto l8
		default:
			if v14 != 0 {
				goto l8
			}
			if v3 == 0 {
				goto l8
			}
			if v16 == 0 {
				goto l9
			}
			goto l8
		case 1:
			if v14 != 0 {
				goto l8
			}
			if v3 == 0 {
				goto l8
			}
			if v16 == 0 {
				goto l9
			}
		}
	l8:
	}
	m.memory[uint32(v0)] = byte(i32(255))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v16))
l5:
	m.g0 = v4 + i32(16)
	return
l3:
	m.fn121(v11, v10, v8, i32(1079412))
	panic("unreachable")
}
func (m *Module) fn575(v0 int32) {
	var v1, v2, v3, v4, v5, v6, v7, v8 int32
	t0 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	v1 = t0
	{
		t1 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		v2 = t1
		if v2 == 0 {
			goto l0
		}
		v3 = i32(0)
	l11:
		{
			v4 = v1 + v3*i32(28)
			t2 := int32(load32(m.memory[int64(uint32(v4))+4:]))
			v5 = t2
			{
				t3 := int32(load32(m.memory[int64(uint32(v4))+8:]))
				v6 = t3
				if v6 == 0 {
					goto l1
				}
				v7 = v5
			l2:
				m.fn335(v7)
				v7 = v7 + i32(32)
				v6 = v6 + i32(-1)
				if v6 != 0 {
					goto l2
				}
			}
		l1:
			{
				t4 := int32(load32(m.memory[uint32(v4):]))
				v7 = t4
				if v7 == 0 {
					goto l3
				}
				t5 := int32(load32(m.memory[uint32(v5+i32(-4)):]))
				v6 = t5
				v8 = v6 & i32(-8)
				t6 := v8
				v6 = v6 & i32(3)
				p7 := i32(8)
				if v6 != 0 {
					p7 = i32(4)
				}
				v7 = v7 << 5
				if uint32(t6) < uint32(p7|v7) {
					m.fn3(i32(1273840), i32(46), i32(1273888))
					panic("unreachable")
				}
				if v6 == 0 {
					goto l5
				}
				if uint32(v8) > uint32(v7+i32(39)) {
					m.fn3(i32(1273904), i32(46), i32(1273952))
					panic("unreachable")
				}
			l5:
				m.fn1(v5)
			}
		l3:
			{
				t8 := int32(load32(m.memory[int64(uint32(v4))+12:]))
				v7 = t8
				if v7 == i32(-1) {
					goto l7
				}
				if v7 == 0 {
					goto l7
				}
				t9 := int32(load32(m.memory[int64(uint32(v4))+16:]))
				v4 = t9
				t10 := int32(load32(m.memory[uint32(v4+i32(-4)):]))
				v6 = t10
				v5 = v6 & i32(-8)
				t11 := v5
				v6 = v6 & i32(3)
				p12 := i32(8)
				if v6 != 0 {
					p12 = i32(4)
				}
				if uint32(t11) < uint32(p12+v7) {
					m.fn3(i32(1273840), i32(46), i32(1273888))
					panic("unreachable")
				}
				if v6 == 0 {
					goto l9
				}
				if uint32(v5) > uint32(v7+i32(39)) {
					m.fn3(i32(1273904), i32(46), i32(1273952))
					panic("unreachable")
				}
			l9:
				m.fn1(v4)
			}
		l7:
			v3 = v3 + i32(1)
			if v3 != v2 {
				goto l11
			}
		}
	}
l0:
	{
		t13 := int32(load32(m.memory[uint32(v0):]))
		v7 = t13
		if v7 == 0 {
			return
		}
		t14 := int32(load32(m.memory[uint32(v1+i32(-4)):]))
		v6 = t14
		v4 = v6 & i32(-8)
		t15 := v4
		v6 = v6 & i32(3)
		p16 := i32(8)
		if v6 != 0 {
			p16 = i32(4)
		}
		v7 = v7 * i32(28)
		if uint32(t15) < uint32(p16+v7) {
			m.fn3(i32(1273840), i32(46), i32(1273888))
			panic("unreachable")
		}
		if v6 == 0 {
			goto l14
		}
		if uint32(v4) > uint32(v7+i32(39)) {
			m.fn3(i32(1273904), i32(46), i32(1273952))
			panic("unreachable")
		}
	l14:
		m.fn1(v1)
	}
}
func (m *Module) fn576(v0 int32) {
	var v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11 int32
	t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
	v1 = t0
	t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t2 := v1
	v2 = t1
	t3 := int32(uint32(t2-v2) / uint32(i32(12)))
	v3 = t3
	if v1 == v2 {
		goto l0
	}
	v4 = i32(0)
l13:
	{
		v5 = v2 + v4*i32(12)
		t4 := int32(load32(m.memory[int64(uint32(v5))+4:]))
		v6 = t4
		{
			t5 := int32(load32(m.memory[int64(uint32(v5))+8:]))
			v7 = t5
			if v7 == 0 {
				goto l1
			}
			v8 = i32(0)
		l8:
			{
				v9 = v6 + v8*i32(28)
				t6 := int32(load32(m.memory[int64(uint32(v9))+4:]))
				v10 = t6
				{
					t7 := int32(load32(m.memory[int64(uint32(v9))+8:]))
					v11 = t7
					if v11 == 0 {
						goto l2
					}
					v1 = v10
				l3:
					m.fn335(v1)
					v1 = v1 + i32(32)
					v11 = v11 + i32(-1)
					if v11 != 0 {
						goto l3
					}
				}
			l2:
				{
					t8 := int32(load32(m.memory[uint32(v9):]))
					v1 = t8
					if v1 == 0 {
						goto l4
					}
					t9 := int32(load32(m.memory[uint32(v10+i32(-4)):]))
					v11 = t9
					v9 = v11 & i32(-8)
					t10 := v9
					v11 = v11 & i32(3)
					p11 := i32(8)
					if v11 != 0 {
						p11 = i32(4)
					}
					v1 = v1 << 5
					if uint32(t10) < uint32(p11|v1) {
						m.fn3(i32(1273840), i32(46), i32(1273888))
						panic("unreachable")
					}
					if v11 == 0 {
						goto l6
					}
					if uint32(v9) > uint32(v1+i32(39)) {
						m.fn3(i32(1273904), i32(46), i32(1273952))
						panic("unreachable")
					}
				l6:
					m.fn1(v10)
				}
			l4:
				v8 = v8 + i32(1)
				if v8 != v7 {
					goto l8
				}
			}
		}
	l1:
		{
			t12 := int32(load32(m.memory[uint32(v5):]))
			v1 = t12
			if v1 == 0 {
				goto l9
			}
			t13 := int32(load32(m.memory[uint32(v6+i32(-4)):]))
			v11 = t13
			v8 = v11 & i32(-8)
			t14 := v8
			v11 = v11 & i32(3)
			p15 := i32(8)
			if v11 != 0 {
				p15 = i32(4)
			}
			v1 = v1 * i32(28)
			if uint32(t14) < uint32(p15+v1) {
				m.fn3(i32(1273840), i32(46), i32(1273888))
				panic("unreachable")
			}
			if v11 == 0 {
				goto l11
			}
			if uint32(v8) > uint32(v1+i32(39)) {
				m.fn3(i32(1273904), i32(46), i32(1273952))
				panic("unreachable")
			}
		l11:
			m.fn1(v6)
		}
	l9:
		v4 = v4 + i32(1)
		if v4 != v3 {
			goto l13
		}
	}
l0:
	{
		t16 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		v1 = t16
		if v1 == 0 {
			return
		}
		t17 := int32(load32(m.memory[uint32(v0):]))
		v8 = t17
		t18 := int32(load32(m.memory[uint32(v8+i32(-4)):]))
		v11 = t18
		v9 = v11 & i32(-8)
		t19 := v9
		v11 = v11 & i32(3)
		p20 := i32(8)
		if v11 != 0 {
			p20 = i32(4)
		}
		v1 = v1 * i32(12)
		if uint32(t19) < uint32(p20+v1) {
			m.fn3(i32(1273840), i32(46), i32(1273888))
			panic("unreachable")
		}
		if v11 == 0 {
			goto l16
		}
		if uint32(v9) > uint32(v1+i32(39)) {
			m.fn3(i32(1273904), i32(46), i32(1273952))
			panic("unreachable")
		}
	l16:
		m.fn1(v8)
	}
}
func (m *Module) fn577(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7, v8, v9, v10 int32
	t0 := m.g0
	v2 = t0 - i32(48)
	m.g0 = v2
	store32(m.memory[int64(uint32(v2))+36:], uint32(v1))
	v3 = i32(0)
	store32(m.memory[int64(uint32(v2))+32:], uint32(i32(0)))
	store32(m.memory[int64(uint32(v2))+28:], uint32(v1))
	store32(m.memory[int64(uint32(v2))+24:], uint32(v0))
	store32(m.memory[int64(uint32(v2))+20:], uint32(i32(58)))
	store32(m.memory[int64(uint32(v2))+40:], uint32(i32(58)))
	m.memory[int64(uint32(v2))+44] = byte(i32(1))
	m.fn199(v2+i32(8), v2+i32(20))
	{
		t1 := int32(load32(m.memory[int64(uint32(v2))+8:]))
		if t1 != i32(1) {
			goto l0
		}
		t2 := int32(load32(m.memory[int64(uint32(v2))+12:]))
		v4 = t2
		if v4 == 0 {
			goto l0
		}
		{
			if uint32(v4) < uint32(v1) {
				goto l1
			}
			if v4 != v1 {
				goto l2
			}
			goto l3
		l1:
			t3 := int32(int8(m.memory[uint32(v0+v4)]))
			if t3 > i32(-65) {
				goto l3
			}
		}
	l2:
		m.fn38(v0, v1, i32(0), v4, i32(1075676))
		panic("unreachable")
	l3:
		v5 = v0 + i32(1)
		{
			{
				t4 := int32(int8(m.memory[uint32(v0)]))
				v6 = t4
				if v6 <= i32(-1) {
					goto l4
				}
				v7 = v6 & i32(255)
				v8 = v5
				goto l5
			}
		l4:
			t5 := int32(m.memory[int64(uint32(v0))+1])
			v8 = t5 & i32(63)
			v7 = v6 & i32(31)
			if uint32(v6) > uint32(i32(-33)) {
				goto l6
			}
			v7 = v7<<6 | v8
			v8 = v0 + i32(2)
			goto l5
		l6:
			t6 := int32(m.memory[int64(uint32(v0))+2])
			v8 = v8<<6 | t6&i32(63)
			if uint32(v6) >= uint32(i32(-16)) {
				goto l7
			}
			v7 = v8 | v7<<12
			v8 = v0 + i32(3)
			goto l5
		l7:
			t7 := int32(m.memory[int64(uint32(v0))+3])
			v7 = v8<<6 | t7&i32(63) | v7<<18&i32(0x1c0000)
			v8 = v0 + i32(4)
		}
	l5:
		if uint32(v7&i32(2097119)+i32(-65)) > uint32(i32(25)) {
			goto l0
		}
		{
			t8 := v8
			v7 = v0 + v4
			if t8 == v7 {
				goto l8
			}
		l14:
			{
				{
					t9 := int32(int8(m.memory[uint32(v8)]))
					v4 = t9
					if v4 <= i32(-1) {
						goto l9
					}
					v8 = v8 + i32(1)
					v4 = v4 & i32(255)
					goto l10
				}
			l9:
				t10 := int32(m.memory[int64(uint32(v8))+1])
				v9 = t10 & i32(63)
				v10 = v4 & i32(31)
				if uint32(v4) > uint32(i32(-33)) {
					goto l11
				}
				v4 = v10<<6 | v9
				v8 = v8 + i32(2)
				goto l10
			l11:
				t11 := int32(m.memory[int64(uint32(v8))+2])
				v9 = v9<<6 | t11&i32(63)
				if uint32(v4) >= uint32(i32(-16)) {
					goto l12
				}
				v4 = v9 | v10<<12
				v8 = v8 + i32(3)
				goto l10
			l12:
				t12 := int32(m.memory[int64(uint32(v8))+3])
				v4 = v9<<6 | t12&i32(63) | v10<<18&i32(0x1c0000)
				v8 = v8 + i32(4)
			}
		l10:
			if uint32(v4+i32(-48)) < uint32(i32(10)) {
				goto l13
			}
			if uint32(v4&i32(2097119)+i32(-65)) < uint32(i32(26)) {
				goto l13
			}
			v4 = v4 + i32(-43)
			if uint32(v4) > uint32(i32(3)) {
				goto l0
			}
			if v4 == i32(1) {
				goto l0
			}
		l13:
			if v8 != v7 {
				goto l14
			}
		}
	l8:
		v3 = i32(1)
		if uint32(v1) < uint32(i32(3)) {
			goto l0
		}
		if uint32((v6&i32(223)+i32(-65))&i32(255)) > uint32(i32(25)) {
			goto l0
		}
		t13 := int32(m.memory[uint32(v5)])
		if t13 != i32(58) {
			goto l0
		}
		t14 := int32(m.memory[int64(uint32(v0))+2])
		v8 = t14
		var p15 int32
		if v8 != i32(92) {
			p15 = 1
		}
		var p16 int32
		if v8 != i32(47) {
			p16 = 1
		}
		v3 = p15 & p16
	}
l0:
	m.g0 = v2 + i32(48)
	return v3
}
func (m *Module) fn578(v0, v1, v2 int32) {
	var v3, v4, v5 int32
	var v6 int64
	var v7 int32
	var v8, v9 int64
	var v10, v11, v12, v13, v14, v15 int32
	var v16, v17 int64
	var v18 float64
	var v19, v20, v21 int32
	t0 := m.g0
	v3 = t0 - i32(1600)
	m.g0 = v3
	{
		if v2 != 0 {
			goto l0
		}
		m.memory[int64(uint32(v0))+1] = byte(i32(0))
		v4 = i32(1)
		goto l1
	l0:
		{
			{
				{
					t1 := int32(m.memory[uint32(v1)])
					v5 = t1
					switch v5 + i32(-43) {
					case 0, 2:
						v4 = i32(1)
						v2 = v2 + i32(-1)
						if v2 == 0 {
							goto l4
						}
						v1 = v1 + i32(1)
						fallthrough
					default:
						v6 = i64(0)
						v4 = v1
						v7 = v2
						{
							{
								{
									{
										{
											if uint32(v2) < uint32(i32(8)) {
												goto l9
											}
											v6 = i64(0)
											v4 = v1
											v7 = v2
										l6:
											{
												t2 := int64(load64(m.memory[uint32(v4):]))
												v8 = t2
												t3 := v8 + i64(5063812098665367110)
												v8 = v8 + i64(-3472328296227680304)
												if !((t3|v8)&i64(-0x7f7f7f7f7f7f7f80) == 0) {
													goto l9
												}
												v8 = v8*i64(10) + int64(uint64(v8)>>8)
												v6 = int64(uint64(int64(uint64(v8)>>16)&i64(0xff000000ff)*i64(0x271000000001)+v8&i64(0xff000000ff)*i64(0xf424000000064))>>32) + v6*i64(100000000)
												v4 = v4 + i32(8)
												v7 = v7 + i32(-8)
												if uint32(v7) > uint32(i32(7)) {
													goto l6
												}
											}
											if v7 != 0 {
												goto l9
											}
											v9 = i64(0)
											v10 = i32(1)
											goto l7
										l9:
											{
												t4 := int32(m.memory[uint32(v4)])
												v11 = t4
												v12 = v11 + i32(-48)
												if uint32(v12&i32(255)) > uint32(i32(9)) {
													goto l8
												}
												v6 = v6*i64(10) + int64(uint32(v12))&i64(255)
												v10 = i32(1)
												v4 = v4 + i32(1)
												v7 = v7 + i32(-1)
												if v7 != 0 {
													goto l9
												}
											}
											v9 = i64(0)
										l7:
											v7 = i32(0)
											v11 = v2
											v8 = i64(0)
											goto l10
										l8:
											v13 = v2 - v7
											if v11&i32(255) == i32(46) {
												goto l11
											}
											v8 = i64(0)
											v11 = i32(0)
											v12 = v7
											goto l12
										l11:
											v4 = v4 + i32(1)
											v10 = v7 + i32(-1)
											if v7 >= i32(9) {
												goto l13
											}
											v12 = v10
											goto l14
										l13:
											v12 = v10
										l16:
											{
												t5 := int64(load64(m.memory[uint32(v4):]))
												v8 = t5
												t6 := v8 + i64(5063812098665367110)
												v8 = v8 + i64(-3472328296227680304)
												if !((t6|v8)&i64(-0x7f7f7f7f7f7f7f80) == 0) {
													goto l15
												}
												v8 = v8*i64(10) + int64(uint64(v8)>>8)
												v6 = int64(uint64(int64(uint64(v8)>>16)&i64(0xff000000ff)*i64(0x271000000001)+v8&i64(0xff000000ff)*i64(0xf424000000064))>>32) + v6*i64(100000000)
												v4 = v4 + i32(8)
												v12 = v12 + i32(-8)
												if uint32(v12) > uint32(i32(7)) {
													goto l16
												}
											}
										l14:
											if v12 == 0 {
												goto l17
											}
										l15:
											v11 = v4
											v4 = v11 + v12
										l20:
											{
												t7 := int32(m.memory[uint32(v11)])
												v14 = t7 + i32(-48)
												if uint32(v14&i32(255)) <= uint32(i32(9)) {
													goto l18
												}
												v4 = v11
												goto l19
											}
										l18:
											v6 = v6*i64(10) + int64(uint32(v14))&i64(255)
											v11 = v11 + i32(1)
											v12 = v12 + i32(-1)
											if v12 != 0 {
												goto l20
											}
										l17:
											v12 = i32(0)
										l19:
											v11 = v10 - v12
											v8 = int64(i32(0) - v11)
										l12:
											v11 = v11 + v13
											if v11 == 0 {
												goto l21
											}
											v9 = i64(0)
											if v12 != 0 {
												goto l22
											}
											v10 = i32(1)
											goto l10
										l22:
											{
												t8 := int32(m.memory[uint32(v4)])
												if t8|i32(32) == i32(101) {
													goto l23
												}
												v10 = i32(0)
												goto l10
											}
										l23:
											v13 = v12 + i32(-1)
											if v13 == 0 {
												goto l21
											}
											v14 = v4 + i32(1)
											t9 := int32(m.memory[int64(uint32(v4))+1])
											v10 = t9
											v15 = v10
											switch v10 + i32(-43) {
											case 0, 2:
												v13 = v12 + i32(-2)
												if v13 == 0 {
													goto l21
												}
												v14 = v4 + i32(2)
												t10 := int32(m.memory[int64(uint32(v4))+2])
												v15 = t10
												fallthrough
											default:
												if uint32((v15+i32(-48))&i32(255)) > uint32(i32(9)) {
													goto l21
												}
												v16 = i64(0)
												v9 = i64(0)
											l27:
												{
													t11 := int32(m.memory[uint32(v14)])
													v4 = t11 + i32(-48)
													if uint32(v4&i32(255)) > uint32(i32(9)) {
														goto l26
													}
													v17 = v9*i64(10) + int64(uint32(v4))&i64(255)
													t12 := v17
													t13 := v9
													var p14 int32
													if v9 < i64(65536) {
														p14 = 1
													}
													v4 = p14
													p15 := t13
													if v4 != 0 {
														p15 = t12
													}
													v9 = p15
													p16 := v16
													if v4 != 0 {
														p16 = v17
													}
													v16 = p16
													v14 = v14 + i32(1)
													v13 = v13 + i32(-1)
													if v13 != 0 {
														goto l27
													}
												}
												v13 = i32(0)
											l26:
												p17 := v16
												if v10 == i32(45) {
													p17 = i64(0) - v16
												}
												v9 = p17
												v8 = v9 + v8
												var p18 int32
												if v13 == 0 {
													p18 = 1
												}
												v10 = p18
											}
										}
									l10:
										if v11 >= i32(20) {
											goto l28
										}
										v4 = i32(0)
										goto l29
									l28:
										v11 = v11 + i32(-19)
										v14 = v2
										v4 = v1
									l32:
										{
											t19 := int32(m.memory[uint32(v4)])
											v12 = t19
											switch v12 + i32(-46) {
											default:
												goto l31
											case 0, 2:
												t20 := v11
												v13 = v12 + i32(-47)
												p21 := v13
												if uint32(v13) > uint32(v12) {
													p21 = i32(0)
												}
												v11 = t20 - p21
												v4 = v4 + i32(1)
												v14 = v14 + i32(-1)
												if v14 != 0 {
													goto l32
												}
											}
										}
									l31:
										if v11 >= i32(1) {
											goto l33
										}
										v4 = i32(0)
										goto l29
									l33:
										v12 = i32(0) - v2
										v6 = i64(0)
										v4 = v1
									l36:
										{
											v11 = v12
											t22 := int32(m.memory[uint32(v4)])
											v14 = t22 + i32(-48)
											if uint32(v14&i32(255)) > uint32(i32(9)) {
												goto l34
											}
											v4 = v4 + i32(1)
											v12 = v11 + i32(1)
											v6 = v6*i64(10) + int64(uint32(v14))&i64(255)
											if uint64(v6) > uint64(i64(999999999999999999)) {
												goto l35
											}
											if v12 != 0 {
												goto l36
											}
										l35:
										}
										if uint64(v6) > uint64(i64(999999999999999999)) {
											goto l37
										}
										if v11 == i32(-1) {
											m.fn121(i32(1), i32(0), i32(0), i32(1109176))
											panic("unreachable")
										}
										v7 = i32(0) - v12
										goto l39
									l34:
										v7 = i32(0) - v11
									l39:
										v14 = v7 + i32(-1)
										if v14 != 0 {
											v4 = v4 + i32(1)
											v7 = v14
										l44:
											{
												t23 := int32(m.memory[uint32(v4)])
												v12 = t23 + i32(-48)
												if uint32(v12&i32(255)) <= uint32(i32(9)) {
													v11 = v7 + i32(-1)
													{
														v6 = v6*i64(10) + int64(uint32(v12))&i64(255)
														if uint64(v6) > uint64(i64(999999999999999999)) {
															goto l43
														}
														v4 = v4 + i32(1)
														var p24 int32
														if v7 != i32(1) {
															p24 = 1
														}
														v12 = p24
														v7 = v11
														if v12 != 0 {
															goto l44
														}
													}
												l43:
													v4 = v11 - v14
													goto l41
												}
												v4 = v7 - v14
												goto l41
											}
										}
										v4 = i32(0) - v14
										goto l41
									l37:
										v4 = i32(0) - (v7 + v12)
									l41:
										v8 = v9 + int64(v4)
										v4 = i32(1)
									l29:
										if v10 == 0 {
											goto l21
										}
										var p25 int32
										if uint64(v8+i64(-38)) < uint64(i64(-60)) {
											p25 = 1
										}
										var p26 int32
										if uint64(v6) > uint64(i64(0x20000000000000)) {
											p26 = 1
										}
										if p25|p26|v4 != 0 {
											goto l45
										}
										{
											if v8 > i64(22) {
												t28 := int64(load64(m.memory[uint32(int32(v8)<<3+i32(1098512)):]))
												m.fn976(v3, v6, i64(0), t28, i64(0))
												t29 := int64(load64(m.memory[int64(uint32(v3))+8:]))
												if t29 != i64(0) {
													goto l45
												}
												t30 := int64(load64(m.memory[uint32(v3):]))
												v9 = t30
												if uint64(v9) > uint64(i64(0x20000000000000)) {
													goto l45
												}
												v18 = float64(float64(uint64(v9)) * float64(1e+22))
												goto l48
											}
											v4 = int32(v8)
											v18 = float64(uint64(v6))
											if v8 < i64(0) {
												goto l47
											}
											t27 := math.Float64frombits(load64(m.memory[int64(uint32(v4<<3))+1122040:]))
											v18 = float64(t27 * v18)
											goto l48
										}
									}
								l21:
									switch v2 + i32(-3) {
									default:
										goto l50
									case 5:
										t31 := int64(load64(m.memory[uint32(v1):]))
										if t31&i64(-2314885530818453537) != i64(0x5954494e49464e49) {
											goto l50
										}
										v18 = math.Float64frombits(0x7ff0000000000000)
										goto l52
									case 0:
										{
											t32 := int64(load16(m.memory[uint32(v1):]))
											t33 := int64(m.memory[int64(uint32(v1))+2])
											v6 = (t32 | t33<<16) & i64(14671839)
											if v6 != i64(4607561) {
												goto l53
											}
											v18 = math.Float64frombits(0x7ff0000000000000)
											goto l52
										}
									l53:
										if v6 != i64(5128526) {
											goto l50
										}
										v18 = math.Float64frombits(0x7ff8000000000000)
									}
								l52:
									t35 := v0
									p34 := v18
									if v5 == i32(45) {
										p34 = -v18
									}
									store64(m.memory[int64(uint32(t35))+8:], math.Float64bits(p34))
									v4 = i32(0)
									goto l1
								}
							l47:
								t36 := math.Float64frombits(load64(m.memory[uint32(i32(1122040)-v4<<3):]))
								v18 = float64(v18 / t36)
							}
						l48:
							t38 := v0
							p37 := v18
							if v5 == i32(45) {
								p37 = -v18
							}
							store64(m.memory[int64(uint32(t38))+8:], math.Float64bits(p37))
							v4 = i32(0)
							goto l1
						}
					l45:
						m.fn875(v3+i32(16), v8, v6)
						{
							{
								t39 := int32(load32(m.memory[int64(uint32(v3))+24:]))
								t40 := v4
								v11 = t39
								var p41 int32
								if v11 > i32(-1) {
									p41 = 1
								}
								if t40&p41 != 0 {
									goto l54
								}
								if v11 < i32(0) {
									goto l55
								}
								t42 := int64(load64(m.memory[int64(uint32(v3))+16:]))
								v6 = t42
								goto l56
							}
						l54:
							m.fn875(v3+i32(816), v8, v6+i64(1))
							t43 := int64(load64(m.memory[int64(uint32(v3))+16:]))
							t44 := int64(load64(m.memory[int64(uint32(v3))+816:]))
							v6 = t44
							if t43 != v6 {
								goto l55
							}
							t45 := int32(load32(m.memory[int64(uint32(v3))+824:]))
							if v11 == t45 {
								goto l56
							}
						}
					l55:
						v19 = v3 + i32(816)
						v7 = i32(0)
						memory_zero(m.memory, uint32(v3+i32(816)), uint32(i32(777)))
						v15 = v3 + i32(824)
						v4 = i32(0)
					l64:
						{
							{
								v12 = v1 + v4
								t46 := int32(m.memory[uint32(v12)])
								v11 = t46
								if v11 == i32(48) {
									goto l57
								}
								v14 = v2 + v7
								v13 = v11 + i32(-48)
								if uint32(v13&i32(255)) > uint32(i32(9)) {
									if v11 == i32(46) {
										v11 = v12 + i32(1)
										v10 = v14 + i32(-1)
										goto l69
									}
									v10 = i32(0)
									v15 = i32(0)
									goto l68
								}
								v10 = v1 + v4
								v12 = v4 ^ i32(-1) + v2
								v4 = i32(0)
							l62:
								if uint32(v4) > uint32(i32(767)) {
									goto l59
								}
								m.memory[uint32(v15+v4)] = byte(v13)
							l59:
								v11 = v10 + v4 + i32(1)
								v7 = v4 + i32(1)
								{
									if v12 == v4 {
										store32(m.memory[uint32(v19):], uint32(v7))
										v15 = i32(0)
										v13 = i32(0)
										goto l63
									}
									v14 = v14 + i32(-1)
									v4 = v7
									t47 := int32(m.memory[uint32(v11)])
									v11 = t47
									v13 = v11 + i32(-48)
									if uint32(v13&i32(255)) > uint32(i32(9)) {
										v12 = v10 + v7
										store32(m.memory[int64(uint32(v3))+816:], uint32(v7))
										v15 = i32(0)
										if v11&i32(255) == i32(46) {
											goto l66
										}
										v13 = v14
										v11 = v12
										goto l63
									}
									goto l62
								}
							}
						l57:
							v7 = v7 + i32(-1)
							t48 := v2
							v4 = v4 + i32(1)
							if t48 != v4 {
								goto l64
							}
						}
						v10 = i32(0)
						goto l65
					l66:
						v11 = v10 + v7 + i32(-1) + i32(2)
						v10 = v14 + i32(1) + i32(-2)
						v13 = v10
						if v7 != 0 {
							goto l70
						}
					l69:
						if v10 != 0 {
							goto l71
						}
						v10 = i32(0)
						v7 = i32(0)
						v13 = i32(0)
						goto l72
					l71:
						v14 = v12 + v14
						v4 = i32(0)
					l74:
						{
							v12 = v11 + v4
							t49 := int32(m.memory[uint32(v12)])
							if t49 != i32(48) {
								goto l73
							}
							t50 := v10
							v4 = v4 + i32(1)
							if t50 != v4 {
								goto l74
							}
						}
						v7 = i32(0)
						v13 = i32(0)
						v11 = v14
						goto l72
					l73:
						v13 = v10 - v4
						v7 = i32(0)
						v11 = v12
					l70:
						if uint32(v13) < uint32(i32(8)) {
							goto l75
						}
						v4 = v7 + i32(8)
					l81:
						{
							v7 = v4
							if uint32(v7) < uint32(i32(768)) {
								goto l76
							}
							v7 = v7 + i32(-8)
							goto l77
						l76:
							t51 := int64(load64(m.memory[uint32(v11):]))
							v6 = t51
							t52 := v6 + i64(5063812098665367110)
							v6 = v6 + i64(-3472328296227680304)
							if (t52|v6)&i64(-0x7f7f7f7f7f7f7f80) != i64(0) {
								goto l78
							}
							v4 = v7 + i32(-8)
							if uint32(v4) > uint32(i32(768)) {
								goto l79
							}
							store64(m.memory[uint32(v3+i32(816)+v7):], uint64(v6))
							v4 = v7 + i32(8)
							v11 = v11 + i32(8)
							v13 = v13 + i32(-8)
							if uint32(v13) <= uint32(i32(7)) {
								goto l80
							}
							goto l81
						l79:
						}
						m.fn121(v4, i32(768), i32(768), i32(1107688))
						panic("unreachable")
					l50:
						v4 = i32(1)
					}
				}
			l4:
				m.memory[int64(uint32(v0))+1] = byte(v4)
				goto l1
			l78:
				v7 = v7 + i32(-8)
			l77:
				store32(m.memory[int64(uint32(v3))+816:], uint32(v7))
				goto l82
			l80:
				store32(m.memory[int64(uint32(v3))+816:], uint32(v7))
			l75:
				if v13 != 0 {
					goto l82
				}
				v13 = i32(0)
				goto l72
			l82:
				{
					t53 := int32(m.memory[uint32(v11)])
					v14 = t53 + i32(-48)
					if uint32(v14&i32(255)) > uint32(i32(9)) {
						goto l83
					}
					v20 = v11 + i32(1)
					v15 = v13 + i32(-1)
					v21 = v7 + (v3 + i32(816)) + i32(8)
					v12 = i32(0)
				l87:
					{
						t54 := v7
						v4 = v12
						v19 = t54 + v4
						if uint32(v19) > uint32(i32(767)) {
							goto l84
						}
						m.memory[uint32(v21+v4)] = byte(v14)
					}
				l84:
					{
						if v15 == v4 {
							goto l85
						}
						v13 = v13 + i32(-1)
						v12 = v4 + i32(1)
						t55 := int32(m.memory[uint32(v20+v4)])
						v14 = t55 + i32(-48)
						if uint32(v14&i32(255)) > uint32(i32(9)) {
							goto l86
						}
						goto l87
					}
				l85:
					v13 = i32(0)
				l86:
					v11 = v11 + v4 + i32(1)
					v7 = v19 + i32(1)
				}
			l83:
				store32(m.memory[int64(uint32(v3))+816:], uint32(v7))
			l72:
				t56 := v3
				v15 = v13 - v10
				store32(m.memory[int64(uint32(t56))+820:], uint32(v15))
			}
		l63:
			if v7 != 0 {
				v4 = v2 - v13
				{
					if uint32(v2) < uint32(v13) {
						m.fn121(i32(0), v4, v2, i32(1107704))
						panic("unreachable")
					}
					v12 = i32(0)
					if v2 == v13 {
						goto l91
					}
					v14 = v1 + i32(-1)
					v12 = i32(0)
				l94:
					{
						t57 := int32(m.memory[uint32(v14+v4)])
						switch t57 + i32(-46) {
						default:
							goto l91
						case 2:
							v12 = v12 + i32(1)
							fallthrough
						case 0:
							v4 = v4 + i32(-1)
							if v4 != 0 {
								goto l94
							}
						}
					}
				l91:
					t58 := v3
					v15 = v15 + v7
					store32(m.memory[int64(uint32(t58))+820:], uint32(v15))
					t59 := v3
					v10 = v7 - v12
					store32(m.memory[int64(uint32(t59))+816:], uint32(v10))
					if uint32(v10) < uint32(i32(769)) {
						goto l89
					}
					v10 = i32(768)
					store32(m.memory[int64(uint32(v3))+816:], uint32(i32(768)))
					m.memory[int64(uint32(v3))+1592] = byte(i32(1))
					goto l89
				}
			}
			v10 = i32(0)
			goto l89
		l89:
			v12 = v11
			v14 = v13
		l68:
			{
				if v14 == 0 {
					goto l95
				}
				t60 := int32(m.memory[uint32(v12)])
				if t60|i32(32) != i32(101) {
					goto l95
				}
				{
					v11 = v14 + i32(-1)
					if v11 != 0 {
						goto l96
					}
					v4 = i32(0)
					goto l97
				l96:
					{
						v7 = v12 + i32(1)
						t61 := int32(m.memory[uint32(v7)])
						v2 = t61
						switch v2 + i32(-43) {
						case 0, 2:
							v11 = v14 + i32(-2)
							if v11 == 0 {
								goto l100
							}
							v7 = v12 + i32(2)
							fallthrough
						default:
							v12 = i32(0)
							v4 = i32(0)
						l102:
							{
								t62 := int32(m.memory[uint32(v7)])
								v14 = (t62 + i32(-48)) & i32(255)
								if uint32(v14) > uint32(i32(9)) {
									goto l101
								}
								v14 = v4*i32(10) + v14
								t63 := v14
								t64 := v4
								var p65 int32
								if v4 < i32(65536) {
									p65 = 1
								}
								v13 = p65
								p66 := t64
								if v13 != 0 {
									p66 = t63
								}
								v4 = p66
								p67 := v12
								if v13 != 0 {
									p67 = v14
								}
								v12 = p67
								v7 = v7 + i32(1)
								v11 = v11 + i32(-1)
								if v11 != 0 {
									goto l102
								}
								goto l101
							}
						}
					}
				l100:
					v12 = i32(0)
				l101:
					p68 := v12
					if v2 == i32(45) {
						p68 = i32(0) - v12
					}
					v4 = p68
				}
			l97:
				store32(m.memory[int64(uint32(v3))+820:], uint32(v15+v4))
			}
		l95:
			if uint32(v10) > uint32(i32(18)) {
				goto l103
			}
		l65:
			v4 = i32(19) - v10
			if v4 == 0 {
				goto l103
			}
			memory_zero(m.memory, uint32(v3+i32(816)+v10+i32(8)), uint32(v4))
		l103:
			memory_copy(m.memory, uint32(v3+i32(36)), uint32(v3+i32(816)), uint32(i32(780)))
			v6 = i64(0)
			v11 = i32(0)
			t69 := int32(load32(m.memory[int64(uint32(v3))+36:]))
			if t69 == 0 {
				goto l56
			}
			t70 := int32(load32(m.memory[int64(uint32(v3))+40:]))
			v4 = t70
			if v4 < i32(-324) {
				goto l56
			}
			v11 = i32(2047)
			if v4 > i32(309) {
				goto l56
			}
			if v4 >= i32(1) {
				v7 = i32(0)
			l108:
				v12 = i32(60)
				{
					if uint32(v4) >= uint32(i32(19)) {
						goto l106
					}
					t71 := int32(m.memory[int64(uint32(v4))+1099012])
					v12 = t71
				}
			l106:
				m.fn858(v3+i32(36), v12)
				{
					t72 := int32(load32(m.memory[int64(uint32(v3))+40:]))
					v4 = t72
					if v4 <= i32(-2048) {
						v11 = i32(0)
						goto l56
					}
					v7 = v12 + v7
					if v4 < i32(1) {
						goto l105
					}
					goto l108
				}
			}
			v7 = i32(0)
			goto l105
		l105:
			v13 = v3 + i32(44)
		l113:
			{
				{
					if v4 != 0 {
						goto l109
					}
					t73 := int32(m.memory[int64(uint32(v3))+44])
					v4 = t73
					if uint32(v4) > uint32(i32(4)) {
						goto l110
					}
					p74 := i32(1)
					if uint32(v4) < uint32(i32(2)) {
						p74 = i32(2)
					}
					v12 = p74
					goto l111
				}
			l109:
				v12 = i32(60)
				v4 = i32(0) - v4
				if uint32(v4) >= uint32(i32(19)) {
					goto l111
				}
				t75 := int32(m.memory[int64(uint32(v4))+1099012])
				v12 = t75
			}
		l111:
			m.fn857(v3+i32(36), v12)
			{
				t76 := int32(load32(m.memory[int64(uint32(v3))+40:]))
				v4 = t76
				if v4 <= i32(2047) {
					goto l112
				}
				v11 = i32(2047)
				goto l56
			}
		l112:
			v7 = v7 - v12
			if v4 < i32(1) {
				goto l113
			}
		l110:
			v4 = v7 + i32(-1)
			if v4 > i32(-1023) {
				goto l114
			}
		l115:
			{
				t77 := v3 + i32(36)
				v7 = i32(-1022) - v4
				p78 := i32(60)
				if uint32(v7) < uint32(i32(60)) {
					p78 = v7
				}
				v7 = p78
				m.fn858(t77, v7)
				v4 = v7 + v4
				if uint32(v4) < uint32(i32(-1022)) {
					goto l115
				}
			}
		l114:
			if v4+i32(1023) > i32(2046) {
				goto l56
			}
			m.fn857(v3+i32(36), i32(53))
			{
				{
					{
						t79 := int32(load32(m.memory[int64(uint32(v3))+36:]))
						v12 = t79
						if v12 == 0 {
							goto l116
						}
						t80 := int32(load32(m.memory[int64(uint32(v3))+40:]))
						v2 = t80
						if v2 < i32(0) {
							goto l116
						}
						if uint32(v2) > uint32(i32(18)) {
							goto l117
						}
						if v2 != 0 {
							if v2 != i32(1) {
								v1 = v2 & i32(1)
								v10 = v2 & i32(30)
								v14 = i32(0)
								v8 = i64(0)
							l125:
								v8 = v8 * i64(10)
								{
									v7 = v14
									if uint32(v7) >= uint32(v12) {
										goto l122
									}
									t81 := int64(m.memory[uint32(v3+i32(36)+v7+i32(8))])
									v8 = v8 + t81
								}
							l122:
								v8 = v8 * i64(10)
								{
									v14 = v7 + i32(1)
									if uint32(v14) >= uint32(v12) {
										goto l123
									}
									t82 := int64(m.memory[uint32(v3+i32(36)+v7+i32(9))])
									v8 = v8 + t82
								}
							l123:
								v14 = v14 + i32(1)
								if v14 == v10 {
									goto l124
								}
								goto l125
							}
							v7 = i32(0)
							v8 = i64(0)
							goto l121
						}
						v8 = i64(0)
						goto l119
					}
				l116:
					v11 = v4 + i32(1022)
					goto l56
				l124:
					if v1 == 0 {
						goto l119
					}
					v7 = v7 + i32(2)
				l121:
					v8 = v8 * i64(10)
					if uint32(v7) >= uint32(v12) {
						goto l119
					}
					t83 := int64(m.memory[uint32(v13+v7)])
					v8 = v8 + t83
				}
			l119:
				{
					if uint32(v2) >= uint32(v12) {
						goto l126
					}
					v14 = v13 + v2
					t84 := int32(m.memory[uint32(v14)])
					v7 = t84
					{
						if v2+i32(1) != v12 {
							goto l127
						}
						if v7&i32(255) == i32(5) {
							goto l128
						}
					l127:
						if uint32(v7&i32(255)) > uint32(i32(4)) {
							goto l129
						}
						goto l126
					l128:
						t85 := int32(m.memory[int64(uint32(v3))+812])
						if t85 != 0 {
							goto l129
						}
						if v2 == 0 {
							goto l126
						}
						t86 := int32(m.memory[uint32(v14+i32(-1))])
						if t86&i32(1) == 0 {
							goto l126
						}
					}
				l129:
					v8 = v8 + i64(1)
				}
			l126:
				if uint64(v8) < uint64(i64(0x20000000000000)) {
					goto l130
				}
			l117:
				m.fn858(v3+i32(36), i32(1))
				t87 := m.fn859(v3 + i32(36))
				v8 = t87
				if v4+i32(1024) > i32(2046) {
					goto l56
				}
				v4 = v4 + i32(1)
			}
		l130:
			v6 = v8 & i64(0xfffffffffffff)
			p88 := i32(1023)
			if uint64(v8) < uint64(i64(0x10000000000000)) {
				p88 = i32(1022)
			}
			v11 = p88 + v4
		}
	l56:
		t89 := v0
		v18 = math.Float64frombits(uint64(int64(uint32(v11))<<52 | v6))
		p90 := v18
		if v5 == i32(45) {
			p90 = -v18
		}
		store64(m.memory[int64(uint32(t89))+8:], math.Float64bits(p90))
		v4 = i32(0)
	}
l1:
	m.memory[uint32(v0)] = byte(v4)
	m.g0 = v3 + i32(1600)
}
func (m *Module) fn579(v0, v1, v2, v3 int32) int32 {
	var v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14, v15, v16, v17, v18, v19, v20, v21, v22, v23, v24 int32
	t0 := m.g0
	v4 = t0 - i32(128)
	m.g0 = v4
	v5 = v1 + i32(-8)
	v6 = v2 + v3<<2
	t1 := v0
	v7 = (v0 + i32(3)) & i32(-4)
	v8 = t1 - v7
	v9 = v7 - v0
	v10 = v4 + i32(52) | i32(3)
	v11 = v4 + i32(52) | i32(2)
	v12 = v4 + i32(52) | i32(1)
	var p2 int32
	if uint32(v1) > uint32(i32(7)) {
		p2 = 1
	}
	v13 = p2
	var p3 int32
	if v1 == i32(5) {
		p3 = 1
	}
	v14 = p3
	var p4 int32
	if v1 == i32(6) {
		p4 = 1
	}
	v15 = p4
	v16 = i32(0)
	{
	l7:
		if v2 == v6 {
			goto l0
		}
		v17 = v2 + i32(4)
		{
			t5 := int32(load32(m.memory[uint32(v2):]))
			v18 = t5
			if uint32(v18) < uint32(i32(128)) {
				if v13 != 0 {
					v3 = i32(0)
					v20 = v0
					v2 = v8
					if v7 == v0 {
						goto l13
					}
				l15:
					{
						t15 := int32(m.memory[uint32(v20)])
						if t15 == v18&i32(255) {
							goto l12
						}
						v20 = v20 + i32(1)
						v2 = v2 + i32(1)
						if v2 == 0 {
							goto l14
						}
						goto l15
					}
				}
				v2 = v17
				if v1 == 0 {
					goto l7
				}
				t8 := int32(m.memory[uint32(v0)])
				v3 = v18 & i32(255)
				if t8 == v3 {
					goto l12
				}
				v2 = v17
				if v1 == i32(1) {
					goto l7
				}
				t9 := int32(m.memory[int64(uint32(v0))+1])
				if t9 == v3 {
					goto l12
				}
				v2 = v17
				if v1 == i32(2) {
					goto l7
				}
				t10 := int32(m.memory[int64(uint32(v0))+2])
				if t10 == v3 {
					goto l12
				}
				v2 = v17
				if v1 == i32(3) {
					goto l7
				}
				t11 := int32(m.memory[int64(uint32(v0))+3])
				if t11 == v3 {
					goto l12
				}
				v2 = v17
				if v1 == i32(4) {
					goto l7
				}
				t12 := int32(m.memory[int64(uint32(v0))+4])
				if t12 == v3 {
					goto l12
				}
				v2 = v17
				if v14 != 0 {
					goto l7
				}
				t13 := int32(m.memory[int64(uint32(v0))+5])
				if t13 == v3 {
					goto l12
				}
				v2 = v17
				if v15 != 0 {
					goto l7
				}
				v2 = v17
				t14 := int32(m.memory[int64(uint32(v0))+6])
				if t14 != v3 {
					goto l7
				}
				goto l12
			}
			store32(m.memory[int64(uint32(v4))+52:], uint32(i32(0)))
			v19 = int32(uint32(v18) >> 18)
			v20 = v18&i32(63) | i32(-128)
			v21 = int32(uint32(v18) >> 12)
			v22 = v21&i32(63) | i32(-128)
			v23 = int32(uint32(v18) >> 6)
			v24 = v23&i32(63) | i32(-128)
			if uint32(v18) > uint32(i32(2047)) {
				goto l2
			}
			m.memory[int64(uint32(v4))+52] = byte(v23 | i32(192))
			v3 = i32(2)
			v2 = v12
			goto l3
		l2:
			if uint32(v18) > uint32(i32(0xffff)) {
				goto l4
			}
			m.memory[int64(uint32(v4))+53] = byte(v24)
			m.memory[int64(uint32(v4))+52] = byte(v21 | i32(224))
			v3 = i32(3)
			v2 = v11
			goto l3
		l4:
			m.memory[int64(uint32(v4))+54] = byte(v24)
			m.memory[int64(uint32(v4))+53] = byte(v22)
			m.memory[int64(uint32(v4))+52] = byte(v19 | i32(-16))
			v3 = i32(4)
			v2 = v10
		l3:
			m.memory[uint32(v2)] = byte(v20)
			{
				if uint32(v3) < uint32(v1) {
					goto l5
				}
				{
					if v3 != v1 {
						v2 = v17
						goto l7
					}
					v2 = v17
					t6 := m.fn974(v4+i32(52), v0, v1)
					if t6 != 0 {
						goto l7
					}
					goto l8
				}
			l5:
				m.fn158(v4+i32(64), v0, v1, v4+i32(52), v3)
				m.fn159(v4, v4+i32(64))
				v2 = v17
				t7 := int32(load32(m.memory[uint32(v4):]))
				if t7 == 0 {
					goto l7
				}
			}
		l8:
			if uint32(v18) >= uint32(i32(2048)) {
				if uint32(v18) >= uint32(i32(65536)) {
					v3 = v19 | i32(240)
					v23 = v20 << 24
					v2 = i32(4)
					v5 = v24
					v20 = v22
					goto l10
				}
				v3 = v21 | i32(224)
				v2 = i32(3)
				v23 = i32(0)
				v5 = v20
				v20 = v24
				goto l10
			}
			v3 = v23 | i32(192)
			v2 = i32(2)
			v23 = i32(0)
			v5 = i32(0)
			goto l10
		}
	l14:
		v3 = v9
		if uint32(v9) > uint32(v5) {
			goto l17
		}
	l13:
		v20 = v18 * i32(16843009)
	l18:
		{
			v2 = v0 + v3
			t16 := int32(load32(m.memory[uint32(v2):]))
			v23 = t16 ^ v20
			t17 := int32(load32(m.memory[uint32(v2+i32(4)):]))
			t18 := i32(16843008) - v23 | v23
			v2 = t17 ^ v20
			if t18&(i32(16843008)-v2|v2)&i32(-2139062144) != i32(-2139062144) {
				goto l17
			}
			v3 = v3 + i32(8)
			if uint32(v3) <= uint32(v5) {
				goto l18
			}
		}
	l17:
		v23 = v1 - v3
		v20 = v0 + v3
		v2 = v17
		if v1 == v3 {
			goto l7
		}
	l19:
		{
			t19 := int32(m.memory[uint32(v20)])
			if t19 == v18&i32(255) {
				goto l12
			}
			v20 = v20 + i32(1)
			v23 = v23 + i32(-1)
			if v23 != 0 {
				goto l19
			}
		}
		v2 = v17
		goto l7
	l12:
		v2 = i32(1)
		v23 = i32(0)
		v5 = i32(0)
		v20 = i32(0)
		v3 = v18
	l10:
		store32(m.memory[int64(uint32(v4))+32:], uint32(v1))
		store32(m.memory[int64(uint32(v4))+28:], uint32(i32(0)))
		m.memory[int64(uint32(v4))+24] = byte(v2)
		store32(m.memory[int64(uint32(v4))+16:], uint32(v1))
		store32(m.memory[int64(uint32(v4))+12:], uint32(i32(0)))
		store32(m.memory[int64(uint32(v4))+8:], uint32(v1))
		store32(m.memory[int64(uint32(v4))+4:], uint32(v0))
		store32(m.memory[uint32(v4):], uint32(v18))
		store32(m.memory[int64(uint32(v4))+20:], uint32(v5&i32(255)<<16|v23|v20&i32(255)<<8|v3&i32(255)))
		store16(m.memory[int64(uint32(v4))+36:], uint16(i32(1)))
		m.fn199(v4+i32(64), v4)
		{
			{
				{
					t20 := int32(load32(m.memory[int64(uint32(v4))+64:]))
					if t20 != i32(1) {
						goto l20
					}
					t21 := int32(load32(m.memory[int64(uint32(v4))+28:]))
					v3 = t21
					t22 := int32(load32(m.memory[int64(uint32(v4))+72:]))
					store32(m.memory[int64(uint32(v4))+28:], uint32(t22))
					v18 = v0 + v3
					t23 := int32(load32(m.memory[int64(uint32(v4))+68:]))
					v3 = t23 - v3
					goto l21
				}
			l20:
				t24 := int32(m.memory[int64(uint32(v4))+37])
				if t24 != 0 {
					goto l22
				}
				m.memory[int64(uint32(v4))+37] = byte(i32(1))
				{
					{
						t25 := int32(m.memory[int64(uint32(v4))+36])
						if t25 != i32(1) {
							goto l23
						}
						t26 := int32(load32(m.memory[int64(uint32(v4))+32:]))
						v20 = t26
						t27 := int32(load32(m.memory[int64(uint32(v4))+28:]))
						v3 = t27
						goto l24
					}
				l23:
					t28 := int32(load32(m.memory[int64(uint32(v4))+32:]))
					v20 = t28
					t29 := int32(load32(m.memory[int64(uint32(v4))+28:]))
					t30 := v20
					v3 = t29
					if t30 == v3 {
						goto l22
					}
				}
			l24:
				t31 := int32(load32(m.memory[int64(uint32(v4))+4:]))
				v18 = t31 + v3
				v3 = v20 - v3
			}
		l21:
			t32 := m.fn5(i32(32))
			v0 = t32
			if v0 == 0 {
				m.fn10(i32(4), i32(32))
				panic("unreachable")
			}
			store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
			store32(m.memory[uint32(v0):], uint32(v18))
			store32(m.memory[int64(uint32(v4))+48:], uint32(i32(1)))
			store32(m.memory[int64(uint32(v4))+44:], uint32(v0))
			store32(m.memory[int64(uint32(v4))+40:], uint32(i32(4)))
			t33 := int64(load64(m.memory[int64(uint32(v4))+32:]))
			store64(m.memory[int64(uint32(v4))+96:], uint64(t33))
			t34 := int64(load64(m.memory[int64(uint32(v4))+24:]))
			store64(m.memory[int64(uint32(v4))+88:], uint64(t34))
			t35 := int64(load64(m.memory[int64(uint32(v4))+16:]))
			store64(m.memory[int64(uint32(v4))+80:], uint64(t35))
			t36 := int64(load64(m.memory[int64(uint32(v4))+8:]))
			store64(m.memory[int64(uint32(v4))+72:], uint64(t36))
			t37 := int64(load64(m.memory[uint32(v4):]))
			store64(m.memory[int64(uint32(v4))+64:], uint64(t37))
			v2 = i32(1)
			{
				t38 := int32(m.memory[int64(uint32(v4))+101])
				if t38 != 0 {
					goto l26
				}
				v3 = i32(12)
				v2 = i32(1)
			l32:
				{
					t39 := int32(load32(m.memory[int64(uint32(v4))+68:]))
					v18 = t39
					m.fn199(v4+i32(52), v4+i32(64))
					{
						{
							t40 := int32(load32(m.memory[int64(uint32(v4))+52:]))
							if t40 != i32(1) {
								goto l27
							}
							t41 := int32(load32(m.memory[int64(uint32(v4))+92:]))
							v20 = t41
							t42 := int32(load32(m.memory[int64(uint32(v4))+60:]))
							store32(m.memory[int64(uint32(v4))+92:], uint32(t42))
							v18 = v18 + v20
							t43 := int32(load32(m.memory[int64(uint32(v4))+56:]))
							v20 = t43 - v20
							goto l28
						}
					l27:
						t44 := int32(m.memory[int64(uint32(v4))+101])
						if t44 != 0 {
							goto l26
						}
						m.memory[int64(uint32(v4))+101] = byte(i32(1))
						{
							{
								t45 := int32(m.memory[int64(uint32(v4))+100])
								if t45 != i32(1) {
									goto l29
								}
								t46 := int32(load32(m.memory[int64(uint32(v4))+96:]))
								v23 = t46
								t47 := int32(load32(m.memory[int64(uint32(v4))+92:]))
								v20 = t47
								goto l30
							}
						l29:
							t48 := int32(load32(m.memory[int64(uint32(v4))+96:]))
							v23 = t48
							t49 := int32(load32(m.memory[int64(uint32(v4))+92:]))
							t50 := v23
							v20 = t49
							if t50 == v20 {
								goto l26
							}
						}
					l30:
						t51 := int32(load32(m.memory[int64(uint32(v4))+68:]))
						v18 = t51 + v20
						v20 = v23 - v20
					}
				l28:
					{
						t52 := int32(load32(m.memory[int64(uint32(v4))+40:]))
						if v2 != t52 {
							goto l31
						}
						m.fn197(v4+i32(40), v2, i32(1), i32(4), i32(8))
						t53 := int32(load32(m.memory[int64(uint32(v4))+44:]))
						v0 = t53
					}
				l31:
					v23 = v0 + v3
					store32(m.memory[uint32(v23):], uint32(v20))
					store32(m.memory[uint32(v23+i32(-4)):], uint32(v18))
					t54 := v4
					v2 = v2 + i32(1)
					store32(m.memory[int64(uint32(t54))+48:], uint32(v2))
					v3 = v3 + i32(8)
					t55 := int32(m.memory[int64(uint32(v4))+101])
					if t55 == 0 {
						goto l32
					}
				}
			}
		l26:
			t56 := int32(load32(m.memory[int64(uint32(v4))+44:]))
			v1 = t56
			v6 = v1 + v2<<3
			t57 := int32(load32(m.memory[int64(uint32(v4))+40:]))
			v5 = t57
			v3 = v1
		l35:
			{
				v16 = i32(0)
				t58 := int32(load32(m.memory[uint32(v3+i32(4)):]))
				v18 = t58
				if uint32(v18+i32(-1)) > uint32(i32(3)) {
					goto l33
				}
				v17 = v3 + i32(8)
				t59 := int32(load32(m.memory[uint32(v3):]))
				v3 = t59
				v20 = v3 + v18
			l40:
				if v3 != v20 {
					{
						{
							t60 := int32(int8(m.memory[uint32(v3)]))
							v18 = t60
							if v18 <= i32(-1) {
								goto l36
							}
							v3 = v3 + i32(1)
							v18 = v18 & i32(255)
							goto l37
						}
					l36:
						t61 := int32(m.memory[int64(uint32(v3))+1])
						v23 = t61 & i32(63)
						v0 = v18 & i32(31)
						if uint32(v18) > uint32(i32(-33)) {
							goto l38
						}
						v18 = v0<<6 | v23
						v3 = v3 + i32(2)
						goto l37
					l38:
						t62 := int32(m.memory[int64(uint32(v3))+2])
						v23 = v23<<6 | t62&i32(63)
						if uint32(v18) >= uint32(i32(-16)) {
							goto l39
						}
						v18 = v23 | v0<<12
						v3 = v3 + i32(3)
						goto l37
					l39:
						t63 := int32(m.memory[int64(uint32(v3))+3])
						v18 = v23<<6 | t63&i32(63) | v0<<18&i32(0x1c0000)
						v3 = v3 + i32(4)
					}
				l37:
					if uint32(v18+i32(-58)) >= uint32(i32(-10)) {
						goto l40
					}
					goto l33
				}
				v3 = v17
				if v17 != v6 {
					goto l35
				}
				v16 = v2
				goto l33
			}
		}
	l22:
		v1 = i32(4)
		v5 = i32(0)
		v16 = i32(0)
	l33:
		if v5 == 0 {
			goto l0
		}
		t64 := int32(load32(m.memory[uint32(v1+i32(-4)):]))
		v3 = t64
		v18 = v3 & i32(-8)
		t65 := v18
		v3 = v3 & i32(3)
		p66 := i32(8)
		if v3 != 0 {
			p66 = i32(4)
		}
		v20 = v5 << 3
		if uint32(t65) < uint32(p66+v20) {
			m.fn3(i32(1273840), i32(46), i32(1273888))
			panic("unreachable")
		}
		if v3 == 0 {
			goto l42
		}
		if uint32(v18) > uint32(v20+i32(39)) {
			m.fn3(i32(1273904), i32(46), i32(1273952))
			panic("unreachable")
		}
	l42:
		m.fn1(v1)
	}
l0:
	m.g0 = v4 + i32(128)
	return v16
}
func (m *Module) fn580(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7, v8, v9 int32
	t0 := v1
	t1 := v0
	v2 = (v0 + i32(3)) & i32(-4)
	v3 = t1 - v2
	v4 = t0 + v3
	v5 = v4 & i32(3)
	v6 = i32(0)
	v1 = i32(0)
	if v0 == v2 {
		goto l0
	}
	v7 = i32(0)
	v1 = i32(0)
	if uint32(v3) > uint32(i32(-4)) {
		goto l1
	}
	v7 = i32(0)
	v1 = i32(0)
l2:
	{
		t2 := v1
		v8 = v0 + v7
		t3 := int32(int8(m.memory[uint32(v8)]))
		var p4 int32
		if t3 > i32(-65) {
			p4 = 1
		}
		t5 := int32(int8(m.memory[uint32(v8+i32(1))]))
		t6 := t2 + p4
		var p7 int32
		if t5 > i32(-65) {
			p7 = 1
		}
		t8 := int32(int8(m.memory[uint32(v8+i32(2))]))
		t9 := t6 + p7
		var p10 int32
		if t8 > i32(-65) {
			p10 = 1
		}
		t11 := int32(int8(m.memory[uint32(v8+i32(3))]))
		t12 := t9 + p10
		var p13 int32
		if t11 > i32(-65) {
			p13 = 1
		}
		v1 = t12 + p13
		v7 = v7 + i32(4)
		if v7 != 0 {
			goto l2
		}
	}
l1:
	v8 = v0 + v7
l3:
	{
		t14 := int32(int8(m.memory[uint32(v8)]))
		t15 := v1
		var p16 int32
		if t14 > i32(-65) {
			p16 = 1
		}
		v1 = t15 + p16
		v8 = v8 + i32(1)
		v3 = v3 + i32(1)
		if v3 != 0 {
			goto l3
		}
	}
l0:
	{
		if v5 == 0 {
			goto l4
		}
		v8 = v2 + v4&i32(0x7ffffffc)
		t17 := int32(int8(m.memory[uint32(v8)]))
		var p18 int32
		if t17 > i32(-65) {
			p18 = 1
		}
		v6 = p18
		if v5 == i32(1) {
			goto l4
		}
		t19 := int32(int8(m.memory[int64(uint32(v8))+1]))
		t20 := v6
		var p21 int32
		if t19 > i32(-65) {
			p21 = 1
		}
		v6 = t20 + p21
		if v5 == i32(2) {
			goto l4
		}
		t22 := int32(int8(m.memory[int64(uint32(v8))+2]))
		t23 := v6
		var p24 int32
		if t22 > i32(-65) {
			p24 = 1
		}
		v6 = t23 + p24
	}
l4:
	v7 = int32(uint32(v4) >> 2)
	v5 = v6 + v1
	{
	l9:
		{
			v0 = v2
			if v7 == 0 {
				goto l5
			}
			p25 := i32(192)
			if uint32(v7) < uint32(i32(192)) {
				p25 = v7
			}
			v6 = p25
			v4 = v6 & i32(3)
			v9 = v6 << 2
			v1 = v9 & i32(1008)
			if v1 != 0 {
				goto l6
			}
			v8 = i32(0)
			goto l7
		l6:
			v3 = v0 + v1
			v8 = i32(0)
			v1 = v0
		l8:
			{
				t26 := int32(load32(m.memory[uint32(v1+i32(12)):]))
				v2 = t26
				t27 := int32(load32(m.memory[uint32(v1+i32(8)):]))
				t28 := (int32(uint32(v2^i32(-1))>>7) | int32(uint32(v2)>>6)) & i32(16843009)
				v2 = t27
				t29 := int32(load32(m.memory[uint32(v1+i32(4)):]))
				t30 := (int32(uint32(v2^i32(-1))>>7) | int32(uint32(v2)>>6)) & i32(16843009)
				v2 = t29
				t31 := int32(load32(m.memory[uint32(v1):]))
				t32 := (int32(uint32(v2^i32(-1))>>7) | int32(uint32(v2)>>6)) & i32(16843009)
				v2 = t31
				v8 = t28 + (t30 + (t32 + ((int32(uint32(v2^i32(-1))>>7)|int32(uint32(v2)>>6))&i32(16843009) + v8)))
				v1 = v1 + i32(16)
				if v1 != v3 {
					goto l8
				}
			}
		l7:
			v7 = v7 - v6
			v2 = v0 + v9
			v5 = int32(uint32((int32(uint32(v8)>>8)&i32(0xff00ff)+v8&i32(0xff00ff))*i32(65537))>>16) + v5
			if v4 == 0 {
				goto l9
			}
		}
		v8 = v0 + v6&i32(252)<<2
		t33 := int32(load32(m.memory[uint32(v8):]))
		v1 = t33
		v1 = (int32(uint32(v1^i32(-1))>>7) | int32(uint32(v1)>>6)) & i32(16843009)
		{
			if v4 == i32(1) {
				goto l10
			}
			t34 := int32(load32(m.memory[int64(uint32(v8))+4:]))
			v2 = t34
			v1 = (int32(uint32(v2^i32(-1))>>7)|int32(uint32(v2)>>6))&i32(16843009) + v1
			if v4 == i32(2) {
				goto l10
			}
			t35 := int32(load32(m.memory[int64(uint32(v8))+8:]))
			v8 = t35
			v1 = (int32(uint32(v8^i32(-1))>>7)|int32(uint32(v8)>>6))&i32(16843009) + v1
		}
	l10:
		v5 = int32(uint32((int32(uint32(v1)>>8)&i32(459007)+v1&i32(0xff00ff))*i32(65537))>>16) + v5
	}
l5:
	return v5
}
func (m *Module) fn581(v0, v1, v2 int32) {
	var v3, v4 int32
	var v5 int64
	var v6, v7 int32
	var v8 int64
	var v9, v10, v11, v12 int32
	var v13 int64
	var v14, v15, v16 int32
	var v17 int64
	var v18, v19, v20, v21, v22, v23, v24, v25, v26, v27, v28 int32
	t0 := m.g0
	v3 = t0 - i32(592)
	m.g0 = v3
	v4 = i32(0)
	memory_zero(m.memory, uint32(v3+i32(56)), uint32(i32(512)))
	t1 := int64(load64(m.memory[int64(uint32(v1))+8:]))
	v5 = t1
	t2 := int32(load32(m.memory[uint32(v1):]))
	v6 = t2
	t3 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	v7 = t3
	v8 = int64(uint32(v7))
l76:
	{
		t5 := v6
		p4 := v8
		if uint64(v5) < uint64(v8) {
			p4 = v5
		}
		v9 = int32(p4)
		v10 = t5 + v9
		v11 = v3 + i32(56) + v4
		{
			v9 = v7 - v9
			t6 := v9
			v12 = i32(512) - v4
			p7 := v12
			if uint32(v9) < uint32(v12) {
				p7 = t6
			}
			v9 = p7
			if v9 != i32(1) {
				if v9 == 0 {
					goto l2
				}
				memory_copy(m.memory, uint32(v11), uint32(v10), uint32(v9))
			l2:
				v5 = v5 + int64(uint32(v9))
				if v9 != 0 {
					goto l1
				}
				store64(m.memory[int64(uint32(v1))+8:], uint64(v5))
				t9 := int64(load64(m.memory[int64(uint32(v3))+56:]))
				v13 = t9
				{
					{
						{
							{
								{
									if v4 != i32(512) {
										goto l3
									}
									if v13 != i64(-0x1ee54e5e1fee3030) {
										goto l3
									}
									v14 = i32(512)
									v15 = i32(1)
									t10 := int32(load16(m.memory[int64(uint32(v3))+82:]))
									v16 = t10
									t11 := int32(load16(m.memory[int64(uint32(v3))+86:]))
									v4 = t11
									switch v4 + i32(-9) {
									case 0:
										goto l4
									case 3:
										goto l6
									default:
										v12 = i32(1067916)
										v14 = int32(uint32(i32(1067916)) >> 24)
										v6 = int32(uint32(i32(1067916)) >> 16)
										v1 = int32(uint32(i32(1067916)) >> 8)
										v11 = v4<<16 | i32(4)
										v4 = i32(1067928)
										v10 = i32(12)
										v2 = i32(0)
										v7 = i32(0)
										v9 = i32(12)
										goto l7
									}
								}
							l3:
								t13 := int32(v13) << 16
								p12 := i32(257)
								if v13 == i64(-0x1ee54e5e1fee3030) {
									p12 = i32(1)
								}
								v11 = t13 | p12
								v7 = int32(int64(uint64(v13) >> 56))
								v10 = int32(int64(uint64(v13) >> 48))
								v14 = int32(int64(uint64(v13) >> 40))
								v6 = int32(int64(uint64(v13) >> 32))
								v1 = int32(int64(uint64(v13) >> 24))
								v12 = int32(int64(uint64(v13) >> 16))
								v2 = i32(0)
								goto l7
							}
						l6:
							{
								{
									t15 := v7
									p14 := v8
									if uint64(v5) < uint64(v8) {
										p14 = v5
									}
									if uint32(t15-int32(p14)) <= uint32(i32(3583)) {
										goto l8
									}
									v13 = i64(0)
									v12 = i32(255)
									goto l9
								}
							l8:
								t16 := int64(load64(m.memory[int64(uint32(i32(0)))+1276648:]))
								v17 = t16
								v13 = int64(uint64(v17) >> 8)
								v12 = int32(v17)
								if v17&i64(255) != i64(255) {
									goto l10
								}
							}
						l9:
							v8 = v5 + i64(3584)
						l10:
							store64(m.memory[int64(uint32(v1))+8:], uint64(v8))
							if v12&i32(255) != i32(255) {
								goto l11
							}
							v14 = i32(4096)
							v15 = i32(0)
						l4:
							{
								t17 := int32(load16(m.memory[int64(uint32(v3))+88:]))
								v4 = t17
								if v4 != i32(6) {
									v12 = i32(1067896)
									v14 = int32(uint32(i32(1067896)) >> 24)
									v10 = i32(16)
									v6 = int32(uint32(i32(1067896)) >> 16)
									v1 = int32(uint32(i32(1067896)) >> 8)
									v9 = i32(4)
									v11 = v4<<16 | i32(4)
									v4 = i32(1067912)
									v2 = i32(0)
									v7 = i32(0)
									goto l7
								}
								t18 := int32(load32(m.memory[int64(uint32(v3))+118:]))
								v7 = t18
								if uint32(v7) > uint32(i32(0x3fffffff)) {
									goto l13
								}
								v6 = v7 << 2
								if uint32(v6) >= uint32(i32(0x7ffffffd)) {
									goto l13
								}
								t19 := int32(load32(m.memory[int64(uint32(v3))+124:]))
								v9 = t19
								t20 := int32(load32(m.memory[int64(uint32(v3))+120:]))
								v10 = t20
								t21 := int32(load32(m.memory[int64(uint32(v3))+116:]))
								v4 = t21
								t22 := int32(load32(m.memory[int64(uint32(v3))+104:]))
								v11 = t22
								t23 := int32(load32(m.memory[int64(uint32(v3))+100:]))
								v12 = t23
								if v6 != 0 {
									t24 := m.fn5(v6)
									v18 = t24
									if v18 != 0 {
										goto l15
									}
									m.fn10(i32(4), v6)
									panic("unreachable")
								}
								v18 = i32(4)
								v7 = i32(0)
								goto l15
							}
						l11:
							v2 = int32(int64(uint64(v13) >> 40))
							v7 = int32(int64(uint64(v13) >> 32))
							v10 = int32(int64(uint64(v13) >> 24))
							v14 = int32(int64(uint64(v13) >> 16))
							v6 = int32(int64(uint64(v13) >> 8))
							v1 = int32(v13)
							v11 = i32(0)
						l7:
							v14 = v14 & i32(255)
							goto l16
						l15:
							store32(m.memory[int64(uint32(v3))+20:], uint32(i32(0)))
							store32(m.memory[int64(uint32(v3))+16:], uint32(v18))
							store32(m.memory[int64(uint32(v3))+12:], uint32(v7))
							store64(m.memory[int64(uint32(v3))+584:], uint64(i64(0x400000000)))
							store32(m.memory[int64(uint32(v3))+580:], uint32(v3+i32(568)))
							store32(m.memory[int64(uint32(v3))+576:], uint32(i32(436)))
							store32(m.memory[int64(uint32(v3))+572:], uint32(v3+i32(132)))
							m.fn636(v3+i32(12), v3+i32(572))
							t25 := int32(load32(m.memory[int64(uint32(v3))+12:]))
							v18 = t25
							if v18 != i32(-1) {
								t26 := int64(load64(m.memory[int64(uint32(v3))+16:]))
								v5 = t26
								store32(m.memory[uint32(v3):], uint32(v18))
								store64(m.memory[int64(uint32(v3))+4:], uint64(v5))
								{
									t27 := m.fn5(i32(1024))
									v6 = t27
									if v6 == 0 {
										m.fn10(i32(1), i32(1024))
										panic("unreachable")
									}
									v7 = int32(v5)
									t28 := v3
									v19 = i32_shr_u(v2, int32(bits.TrailingZeros32(uint32(v14))))
									store32(m.memory[int64(uint32(t28))+28:], uint32(v19))
									store32(m.memory[int64(uint32(v3))+24:], uint32(v14))
									store32(m.memory[int64(uint32(v3))+20:], uint32(i32(0)))
									store32(m.memory[int64(uint32(v3))+16:], uint32(v6))
									store32(m.memory[int64(uint32(v3))+12:], uint32(i32(1024)))
									if uint32(v9) <= uint32(i32(-7)) {
									l26:
										m.fn637(v3+i32(56), v3+i32(12), v9, v1)
										{
											t29 := int32(m.memory[int64(uint32(v3))+56])
											v9 = t29
											if v9 == i32(255) {
												t36 := int32(load32(m.memory[int64(uint32(v3))+60:]))
												v9 = t36
												t37 := int32(load32(m.memory[int64(uint32(v3))+64:]))
												t38 := v3
												v7 = t37
												v6 = v7 & i32(3)
												store32(m.memory[int64(uint32(t38))+56:], uint32(v6))
												if v6 != 0 {
													m.fn638(i32(0), v3+i32(56), i32(1277068), i32(0), v4, i32(1090920))
													panic("unreachable")
												}
												store64(m.memory[int64(uint32(v3))+584:], uint64(i64(0x400000000)))
												store32(m.memory[int64(uint32(v3))+572:], uint32(v9))
												store32(m.memory[int64(uint32(v3))+576:], uint32(v7))
												store32(m.memory[int64(uint32(v3))+580:], uint32(v9+v7))
												m.fn636(v3, v3+i32(572))
												t39 := int32(load32(m.memory[int64(uint32(v3))+8:]))
												v9 = t39
												if v9 == 0 {
													m.fn219(i32(1067880))
													panic("unreachable")
												}
												t40 := v3
												v6 = v9 + i32(-1)
												store32(m.memory[int64(uint32(t40))+8:], uint32(v6))
												t41 := int32(load32(m.memory[uint32(v3):]))
												v18 = t41
												t42 := int32(load32(m.memory[int64(uint32(v3))+4:]))
												v7 = t42
												t43 := int32(load32(m.memory[uint32(v7+v6<<2):]))
												v9 = t43
												if uint32(v9) >= uint32(i32(-6)) {
													goto l21
												}
												goto l26
											}
											t30 := int32(m.memory[int64(uint32(v3))+59])
											m.memory[int64(uint32(v0))+7] = byte(t30)
											t31 := int32(load16(m.memory[int64(uint32(v3))+57:]))
											store16(m.memory[int64(uint32(v0))+5:], uint16(t31))
											t32 := int64(load64(m.memory[int64(uint32(v3))+60:]))
											v5 = t32
											t33 := int64(load64(m.memory[int64(uint32(v3))+68:]))
											store64(m.memory[int64(uint32(v0))+16:], uint64(t33))
											store64(m.memory[int64(uint32(v0))+8:], uint64(v5))
											m.memory[int64(uint32(v0))+4] = byte(v9)
											store32(m.memory[uint32(v0):], uint32(i32(-1)))
											{
												t34 := int32(load32(m.memory[int64(uint32(v3))+12:]))
												v4 = t34
												if v4 == 0 {
													goto l23
												}
												t35 := int32(load32(m.memory[int64(uint32(v3))+16:]))
												m.fn18(t35, v4, i32(1))
											}
										l23:
											if v18 == 0 {
												goto l18
											}
											m.fn18(v7, v18<<2, i32(4))
											goto l18
										}
									}
									v6 = int32(int64(uint64(v5) >> 32))
									goto l21
								}
							l21:
								if uint32(v12) > uint32(i32(0x3fffffff)) {
									goto l13
								}
								v2 = v12 << 2
								if uint32(v2) >= uint32(i32(0x7ffffffd)) {
									goto l13
								}
								v9 = i32(0)
								v20 = i32(0)
								v21 = i32(4)
								{
									if v2 == 0 {
										goto l27
									}
									t44 := m.fn5(v2)
									v21 = t44
									if v21 == 0 {
										m.fn10(i32(4), v2)
										panic("unreachable")
									}
									v20 = v12
								}
							l27:
								store32(m.memory[int64(uint32(v3))+40:], uint32(i32(0)))
								store32(m.memory[int64(uint32(v3))+36:], uint32(v21))
								store32(m.memory[int64(uint32(v3))+32:], uint32(v20))
								v6 = v6 << 2
							l31:
								{
									if v6 == v9 {
										if v18 == 0 {
											goto l32
										}
										m.fn18(v7, v18<<2, i32(4))
									l32:
										t46 := int32(load32(m.memory[int64(uint32(v3))+36:]))
										t47 := v3 + i32(56)
										t48 := v3 + i32(12)
										t49 := v11
										v22 = t46
										t50 := int32(load32(m.memory[int64(uint32(v3))+40:]))
										t51 := v22
										v23 = t50
										m.fn639(t47, t48, t49, t51, v23, v1, i32(-1))
										{
											t52 := int32(m.memory[int64(uint32(v3))+56])
											v9 = t52
											if v9 == i32(255) {
												t57 := int32(load32(m.memory[int64(uint32(v3))+64:]))
												v24 = t57
												t58 := int32(load32(m.memory[int64(uint32(v3))+60:]))
												v25 = t58
												{
													t59 := int32(load32(m.memory[int64(uint32(v3))+68:]))
													v2 = t59
													if v2 != 0 {
														v18 = i32(0)
														t60 := int32(uint32(v2) >> 7)
														var p61 int32
														if v2&i32(127) != i32(0) {
															p61 = 1
														}
														v26 = t60 + p61
														v9 = v26 * i32(20)
														t62 := m.fn5(v9)
														v27 = t62
														if v27 == 0 {
															m.fn10(i32(4), v9)
															panic("unreachable")
														}
														v6 = v24
													l60:
														{
															p63 := i32(128)
															if uint32(v2) < uint32(i32(128)) {
																p63 = v2
															}
															v21 = p63
															if uint32(v2) <= uint32(i32(63)) {
																m.fn121(i32(0), i32(64), v21, i32(1080552))
																panic("unreachable")
															}
															{
																{
																	t64 := int32(load16(m.memory[uint32(v6):]))
																	t65 := int32(m.memory[uint32(v6+i32(2))])
																	if (t64^i32(48111)|(t65^i32(191)))&i32(0xffff) != 0 {
																		goto l39
																	}
																	v7 = i32(1271548)
																	v9 = i32(3)
																	goto l40
																}
															l39:
																v9 = i32(2)
																{
																	t66 := int32(load16(m.memory[uint32(v6):]))
																	if t66 != i32(65279) {
																		goto l41
																	}
																	v7 = i32(1271552)
																	goto l40
																}
															l41:
																{
																	t67 := int32(load16(m.memory[uint32(v6):]))
																	v12 = t67
																	if (v12<<8|int32(uint32(v12)>>8))&i32(0xffff) == i32(65279) {
																		goto l42
																	}
																	v9 = i32(1143932)
																	v12 = i32(64)
																	v11 = v6
																	goto l43
																}
															l42:
																v7 = i32(1271556)
															l40:
																v11 = v6 + v9
																v12 = i32(64) - v9
																t68 := int32(load32(m.memory[uint32(v7):]))
																v9 = t68
															}
														l43:
															m.fn209(v3+i32(56), v9, v11, v12)
															t69 := int32(load32(m.memory[int64(uint32(v3))+64:]))
															v12 = t69
															t70 := int32(load32(m.memory[int64(uint32(v3))+60:]))
															v7 = t70
															{
																t71 := int32(load32(m.memory[int64(uint32(v3))+56:]))
																v20 = t71
																if v20 != i32(-1) {
																	goto l44
																}
																if v12 <= i32(-1) {
																	goto l13
																}
																if v12 != 0 {
																	t72 := m.fn5(v12)
																	v11 = t72
																	if v11 == 0 {
																		m.fn10(i32(1), v12)
																		panic("unreachable")
																	}
																	if v12 == 0 {
																		goto l48
																	}
																	memory_copy(m.memory, uint32(v11), uint32(v7), uint32(v12))
																l48:
																	v20 = v12
																	goto l49
																}
																v11 = i32(1)
																v20 = i32(0)
																v9 = i32(0)
																goto l46
															}
														l44:
															if v12 == 0 {
																goto l50
															}
															v11 = v7
														l49:
															v9 = i32(0)
														l52:
															{
																v7 = v11 + v9
																t73 := int32(m.memory[uint32(v7)])
																if t73 == 0 {
																	if v9 != 0 {
																		t75 := int32(int8(m.memory[uint32(v7)]))
																		if t75 > i32(-65) {
																			goto l46
																		}
																		m.fn3(i32(1080397), i32(48), i32(1080568))
																		panic("unreachable")
																	}
																	v9 = i32(0)
																	goto l46
																}
																t74 := v12
																v9 = v9 + i32(1)
																if t74 != v9 {
																	goto l52
																}
															}
															v9 = v12
															goto l46
														l50:
															v9 = i32(0)
															v11 = v7
														l46:
															if uint32(v2) <= uint32(i32(119)) {
																m.fn121(i32(116), i32(120), v21, i32(1080584))
																panic("unreachable")
															}
															t76 := int32(load32(m.memory[int64(uint32(v6))+116:]))
															v7 = t76
															{
																{
																	if v15 == 0 {
																		goto l55
																	}
																	if uint32(v2) <= uint32(i32(123)) {
																		m.fn121(i32(120), i32(124), v21, i32(1080600))
																		panic("unreachable")
																	}
																	t77 := int32(load32(m.memory[int64(uint32(v6))+120:]))
																	v28 = t77
																	goto l57
																}
															l55:
																if uint32(v2) <= uint32(i32(127)) {
																	m.fn121(i32(120), i32(128), v21, i32(1080616))
																	panic("unreachable")
																}
																t78 := int64(load64(m.memory[int64(uint32(v6))+120:]))
																v5 = t78
																if uint64(v5) > uint64(i64(0xffffffff)) {
																	m.memory[int64(uint32(v3))+56] = byte(i32(2))
																	m.fn42(i32(1284336), i32(43), v3+i32(56), i32(1080632), i32(1080648))
																	panic("unreachable")
																}
																v28 = int32(v5)
															}
														l57:
															v6 = v6 + v21
															v12 = v27 + v18*i32(20)
															store32(m.memory[int64(uint32(v12))+16:], uint32(v28))
															store32(m.memory[int64(uint32(v12))+12:], uint32(v7))
															store32(m.memory[int64(uint32(v12))+8:], uint32(v9))
															store32(m.memory[int64(uint32(v12))+4:], uint32(v11))
															store32(m.memory[uint32(v12):], uint32(v20))
															v18 = v18 + i32(1)
															v2 = v2 - v21
															if v2 != 0 {
																goto l60
															}
														}
														if v18 == 0 {
															goto l36
														}
														{
															if v16&i32(0xffff) == i32(3) {
																goto l61
															}
															t79 := int32(load32(m.memory[int64(uint32(v27))+12:]))
															if t79 != i32(-2) {
																goto l61
															}
															store32(m.memory[uint32(v0):], uint32(i32(-1)))
															m.memory[int64(uint32(v0))+4] = byte(i32(2))
															goto l62
														}
													l61:
														if v10 != 0 {
															t80 := int32(load32(m.memory[int64(uint32(v27))+12:]))
															t81 := int32(load32(m.memory[int64(uint32(v27))+16:]))
															m.fn639(v3+i32(56), v3+i32(12), t80, v22, v23, v1, t81)
															{
																t82 := int32(m.memory[int64(uint32(v3))+56])
																v9 = t82
																if v9 == i32(255) {
																	t87 := int32(load32(m.memory[int64(uint32(v3))+68:]))
																	v11 = t87
																	t88 := int32(load32(m.memory[int64(uint32(v3))+64:]))
																	v12 = t88
																	t89 := int32(load32(m.memory[int64(uint32(v3))+60:]))
																	v9 = t89
																	m.fn639(v3+i32(56), v3+i32(12), v4, v22, v23, v1, v10*v14)
																	{
																		t90 := int32(m.memory[int64(uint32(v3))+56])
																		v4 = t90
																		if v4 == i32(255) {
																			t95 := int32(load32(m.memory[int64(uint32(v3))+60:]))
																			v4 = t95
																			t96 := int32(load32(m.memory[int64(uint32(v3))+64:]))
																			t97 := v3 + i32(56)
																			v10 = t96
																			t98 := int32(load32(m.memory[int64(uint32(v3))+68:]))
																			m.fn640(t97, v10, t98)
																			m.fn641(v3+i32(44), v3+i32(56))
																			if v4 == 0 {
																				goto l64
																			}
																			m.fn18(v10, v4, i32(1))
																			goto l64
																		}
																		t91 := int32(m.memory[int64(uint32(v3))+59])
																		m.memory[int64(uint32(v0))+7] = byte(t91)
																		t92 := int32(load16(m.memory[int64(uint32(v3))+57:]))
																		store16(m.memory[int64(uint32(v0))+5:], uint16(t92))
																		t93 := int64(load64(m.memory[int64(uint32(v3))+60:]))
																		v5 = t93
																		t94 := int64(load64(m.memory[int64(uint32(v3))+68:]))
																		store64(m.memory[int64(uint32(v0))+16:], uint64(t94))
																		store64(m.memory[int64(uint32(v0))+8:], uint64(v5))
																		m.memory[int64(uint32(v0))+4] = byte(v4)
																		store32(m.memory[uint32(v0):], uint32(i32(-1)))
																		if v9 == 0 {
																			goto l62
																		}
																		m.fn18(v12, v9, i32(1))
																		goto l62
																	}
																}
																t83 := int32(m.memory[int64(uint32(v3))+59])
																m.memory[int64(uint32(v0))+7] = byte(t83)
																t84 := int32(load16(m.memory[int64(uint32(v3))+57:]))
																store16(m.memory[int64(uint32(v0))+5:], uint16(t84))
																t85 := int64(load64(m.memory[int64(uint32(v3))+60:]))
																v5 = t85
																t86 := int64(load64(m.memory[int64(uint32(v3))+68:]))
																store64(m.memory[int64(uint32(v0))+16:], uint64(t86))
																store64(m.memory[int64(uint32(v0))+8:], uint64(v5))
																m.memory[int64(uint32(v0))+4] = byte(v9)
																store32(m.memory[uint32(v0):], uint32(i32(-1)))
																goto l62
															}
														}
														v9 = i32(0)
														store32(m.memory[int64(uint32(v3))+52:], uint32(i32(0)))
														store64(m.memory[int64(uint32(v3))+44:], uint64(i64(0x400000000)))
														v12 = i32(1)
														v11 = i32(0)
														goto l64
													}
													v26 = i32(0)
													v27 = i32(4)
													goto l36
												}
											}
											t53 := int32(m.memory[int64(uint32(v3))+59])
											m.memory[int64(uint32(v0))+7] = byte(t53)
											t54 := int32(load16(m.memory[int64(uint32(v3))+57:]))
											store16(m.memory[int64(uint32(v0))+5:], uint16(t54))
											t55 := int64(load64(m.memory[int64(uint32(v3))+60:]))
											v5 = t55
											t56 := int64(load64(m.memory[int64(uint32(v3))+68:]))
											store64(m.memory[int64(uint32(v0))+16:], uint64(t56))
											store64(m.memory[int64(uint32(v0))+8:], uint64(v5))
											m.memory[int64(uint32(v0))+4] = byte(v9)
											store32(m.memory[uint32(v0):], uint32(i32(-1)))
											goto l34
										}
									}
									t45 := int32(load32(m.memory[uint32(v7+v9):]))
									v12 = t45
									if uint32(v12) < uint32(i32(-4)) {
										m.fn637(v3+i32(56), v3+i32(12), v12, v1)
										{
											t99 := int32(m.memory[int64(uint32(v3))+56])
											v12 = t99
											if v12 == i32(255) {
												t104 := int32(load32(m.memory[int64(uint32(v3))+60:]))
												t105 := int32(load32(m.memory[int64(uint32(v3))+64:]))
												m.fn640(v3+i32(572), t104, t105)
												m.fn636(v3+i32(32), v3+i32(572))
												v9 = v9 + i32(4)
												goto l31
											}
											t100 := int32(m.memory[int64(uint32(v3))+59])
											m.memory[int64(uint32(v0))+7] = byte(t100)
											t101 := int32(load16(m.memory[int64(uint32(v3))+57:]))
											store16(m.memory[int64(uint32(v0))+5:], uint16(t101))
											t102 := int64(load64(m.memory[int64(uint32(v3))+60:]))
											v5 = t102
											t103 := int64(load64(m.memory[int64(uint32(v3))+68:]))
											store64(m.memory[int64(uint32(v0))+16:], uint64(t103))
											store64(m.memory[int64(uint32(v0))+8:], uint64(v5))
											m.memory[int64(uint32(v0))+4] = byte(v12)
											store32(m.memory[uint32(v0):], uint32(i32(-1)))
											if v18 == 0 {
												goto l34
											}
											m.fn18(v7, v18<<2, i32(4))
											goto l34
										}
									}
									v9 = v9 + i32(4)
									goto l31
								}
							}
							v2 = int32(uint32(v10) >> 16)
							v7 = int32(uint32(v10) >> 8)
							v14 = int32(uint32(v12) >> 24)
							v6 = int32(uint32(v12) >> 16)
							v1 = int32(uint32(v12) >> 8)
						}
					l16:
						store32(m.memory[int64(uint32(v0))+20:], uint32(v9))
						store32(m.memory[int64(uint32(v0))+16:], uint32(v4))
						store32(m.memory[int64(uint32(v0))+4:], uint32(v11))
						store32(m.memory[uint32(v0):], uint32(i32(-1)))
						store32(m.memory[int64(uint32(v0))+12:], uint32(v2<<16|v7<<8&i32(0xff00)|v10&i32(255)))
						store32(m.memory[int64(uint32(v0))+8:], uint32(v14<<24|v6<<16&i32(0xff0000)|v1<<8&i32(0xff00)|v12&i32(255)))
						goto l18
					l13:
						m.fn9()
						panic("unreachable")
					l64:
						t106 := int32(load32(m.memory[int64(uint32(v3))+28:]))
						store32(m.memory[int64(uint32(v0))+28:], uint32(t106))
						t107 := int64(load64(m.memory[int64(uint32(v3))+20:]))
						store64(m.memory[int64(uint32(v0))+20:], uint64(t107))
						t108 := int64(load64(m.memory[int64(uint32(v3))+12:]))
						store64(m.memory[int64(uint32(v0))+12:], uint64(t108))
						t109 := int64(load64(m.memory[int64(uint32(v3))+32:]))
						store64(m.memory[int64(uint32(v0))+32:], uint64(t109))
						t110 := int32(load32(m.memory[int64(uint32(v3))+40:]))
						store32(m.memory[int64(uint32(v0))+40:], uint32(t110))
						store32(m.memory[int64(uint32(v0))+56:], uint32(i32(64)))
						store32(m.memory[int64(uint32(v0))+52:], uint32(v11))
						store32(m.memory[int64(uint32(v0))+48:], uint32(v12))
						store32(m.memory[int64(uint32(v0))+44:], uint32(v9))
						store32(m.memory[int64(uint32(v0))+8:], uint32(v18))
						store32(m.memory[int64(uint32(v0))+4:], uint32(v27))
						store32(m.memory[uint32(v0):], uint32(v26))
						store32(m.memory[int64(uint32(v0))+60:], uint32(v19<<3))
						t111 := int32(load32(m.memory[int64(uint32(v3))+52:]))
						store32(m.memory[int64(uint32(v0))+72:], uint32(t111))
						t112 := int64(load64(m.memory[int64(uint32(v3))+44:]))
						store64(m.memory[int64(uint32(v0))+64:], uint64(t112))
						if v25 == 0 {
							goto l18
						}
						m.fn18(v24, v25, i32(1))
						goto l18
					}
				l62:
					v4 = v27
				l72:
					{
						t113 := int32(load32(m.memory[uint32(v4):]))
						v9 = t113
						if v9 == 0 {
							goto l68
						}
						t114 := int32(load32(m.memory[uint32(v4+i32(4)):]))
						v10 = t114
						t115 := int32(load32(m.memory[uint32(v10+i32(-4)):]))
						v12 = t115
						v11 = v12 & i32(-8)
						t116 := v11
						v12 = v12 & i32(3)
						p117 := i32(8)
						if v12 != 0 {
							p117 = i32(4)
						}
						if uint32(t116) < uint32(p117+v9) {
							m.fn3(i32(1273840), i32(46), i32(1273888))
							panic("unreachable")
						}
						if v12 == 0 {
							goto l70
						}
						if uint32(v11) > uint32(v9+i32(39)) {
							m.fn3(i32(1273904), i32(46), i32(1273952))
							panic("unreachable")
						}
					l70:
						m.fn1(v10)
					}
				l68:
					v4 = v4 + i32(20)
					v18 = v18 + i32(-1)
					if v18 != 0 {
						goto l72
					}
					goto l73
				l36:
					store32(m.memory[uint32(v0):], uint32(i32(-1)))
					m.memory[int64(uint32(v0))+4] = byte(i32(2))
				l73:
					if v26 == 0 {
						goto l74
					}
					m.fn18(v27, v26*i32(20), i32(4))
				l74:
					if v25 == 0 {
						goto l34
					}
					m.fn18(v24, v25, i32(1))
				l34:
					{
						t118 := int32(load32(m.memory[int64(uint32(v3))+32:]))
						v4 = t118
						if v4 == 0 {
							goto l75
						}
						t119 := int32(load32(m.memory[int64(uint32(v3))+36:]))
						m.fn18(t119, v4<<2, i32(4))
					}
				l75:
					t120 := int32(load32(m.memory[int64(uint32(v3))+12:]))
					v4 = t120
					if v4 == 0 {
						goto l18
					}
					t121 := int32(load32(m.memory[int64(uint32(v3))+16:]))
					m.fn18(t121, v4, i32(1))
				}
			l18:
				m.g0 = v3 + i32(592)
				return
			}
			t8 := int32(m.memory[uint32(v10)])
			m.memory[uint32(v11)] = byte(t8)
			v5 = v5 + i64(1)
			goto l1
		}
	l1:
		v4 = v9 + v4
		goto l76
	}
}
