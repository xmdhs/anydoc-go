package core

import (
	"math/bits"
)

func (m *Module) fn402(v0 int32) {
	var v1, v2, v3, v4 int32
	var v5 int64
	var v6, v7, v8, v9, v10 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+404:]))
		v1 = t0
		if v1 == 0 {
			goto l0
		}
		{
			t1 := int32(load32(m.memory[int64(uint32(v0))+412:]))
			v2 = t1
			if v2 == 0 {
				goto l1
			}
			t2 := int32(load32(m.memory[int64(uint32(v0))+400:]))
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
				v3 = v3 + i32(-224)
				t4 := int64(load64(m.memory[uint32(v6):]))
				v5 = t4 & i64(-0x7f7f7f7f7f7f7f80)
				if v5 == i64(-0x7f7f7f7f7f7f7f80) {
					goto l3
				}
			}
			v5 = v5 ^ i64(-0x7f7f7f7f7f7f7f80)
		l2:
			{
				v6 = v3 + (i32(0)-int32(uint32(int64(bits.TrailingZeros64(uint64(v5))))>>3))*i32(28)
				t5 := int32(load32(m.memory[uint32(v6+i32(-28)):]))
				v7 = t5
				if v7 == 0 {
					goto l4
				}
				t6 := int32(load32(m.memory[uint32(v6+i32(-24)):]))
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
				if uint32(t8) < uint32(p9+v7) {
					m.fn7(i32(1274404), i32(46), i32(1274452))
					panic("unreachable")
				}
				if v9 == 0 {
					goto l6
				}
				if uint32(v10) > uint32(v7+i32(39)) {
					m.fn7(i32(1274468), i32(46), i32(1274516))
					panic("unreachable")
				}
			l6:
				m.fn5(v8)
			}
		l4:
			{
				t10 := int32(load32(m.memory[uint32(v6+i32(-12)):]))
				v7 = t10
				if v7 < i32(1) {
					goto l8
				}
				t11 := int32(load32(m.memory[uint32(v6+i32(-8)):]))
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
				if uint32(t13) < uint32(p14+v7) {
					m.fn7(i32(1274404), i32(46), i32(1274452))
					panic("unreachable")
				}
				if v6 == 0 {
					goto l10
				}
				if uint32(v8) > uint32(v7+i32(39)) {
					m.fn7(i32(1274468), i32(46), i32(1274516))
					panic("unreachable")
				}
			l10:
				m.fn5(v9)
			}
		l8:
			v5 = (v5 + i64(-1)) & v5
			v2 = v2 + i32(-1)
			if v2 != 0 {
				goto l12
			}
		}
	l1:
		t15 := v1
		v4 = (v1*i32(28) + i32(35)) & i32(-8)
		v3 = t15 + v4 + i32(9)
		if v3 == 0 {
			goto l0
		}
		t16 := int32(load32(m.memory[int64(uint32(v0))+400:]))
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
			m.fn7(i32(1274404), i32(46), i32(1274452))
			panic("unreachable")
		}
		if v4 == 0 {
			goto l14
		}
		if uint32(v2) > uint32(v3+i32(39)) {
			m.fn7(i32(1274468), i32(46), i32(1274516))
			panic("unreachable")
		}
	l14:
		m.fn5(v6)
	}
l0:
	m.fn421(v0 + i32(504))
	{
		t20 := int32(load32(m.memory[int64(uint32(v0))+436:]))
		v1 = t20
		if v1 == 0 {
			goto l16
		}
		{
			t21 := int32(load32(m.memory[int64(uint32(v0))+444:]))
			v2 = t21
			if v2 == 0 {
				goto l17
			}
			t22 := int32(load32(m.memory[int64(uint32(v0))+432:]))
			v3 = t22
			v4 = v3 + i32(8)
			t23 := int64(load64(m.memory[uint32(v3):]))
			v5 = (t23 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
		l24:
			if v5 != i64(0) {
				goto l18
			}
		l19:
			{
				v6 = v4
				v4 = v6 + i32(8)
				v3 = v3 + i32(-3328)
				t24 := int64(load64(m.memory[uint32(v6):]))
				v5 = t24 & i64(-0x7f7f7f7f7f7f7f80)
				if v5 == i64(-0x7f7f7f7f7f7f7f80) {
					goto l19
				}
			}
			v5 = v5 ^ i64(-0x7f7f7f7f7f7f7f80)
		l18:
			{
				v6 = v3 + (i32(0)-int32(uint32(int64(bits.TrailingZeros64(uint64(v5))))>>3))*i32(416)
				t25 := int32(load32(m.memory[uint32(v6+i32(-416)):]))
				v7 = t25
				if v7 == 0 {
					goto l20
				}
				t26 := int32(load32(m.memory[uint32(v6+i32(-412)):]))
				v8 = t26
				t27 := int32(load32(m.memory[uint32(v8+i32(-4)):]))
				v9 = t27
				v10 = v9 & i32(-8)
				t28 := v10
				v9 = v9 & i32(3)
				p29 := i32(8)
				if v9 != 0 {
					p29 = i32(4)
				}
				if uint32(t28) < uint32(p29+v7) {
					m.fn7(i32(1274404), i32(46), i32(1274452))
					panic("unreachable")
				}
				if v9 == 0 {
					goto l22
				}
				if uint32(v10) > uint32(v7+i32(39)) {
					m.fn7(i32(1274468), i32(46), i32(1274516))
					panic("unreachable")
				}
			l22:
				m.fn5(v8)
			}
		l20:
			v5 = (v5 + i64(-1)) & v5
			m.fn420(v6 + i32(-400))
			v2 = v2 + i32(-1)
			if v2 != 0 {
				goto l24
			}
		}
	l17:
		v4 = v1 * i32(416)
		v3 = v4 + v1 + i32(425)
		if v3 == 0 {
			goto l16
		}
		t30 := int32(load32(m.memory[int64(uint32(v0))+432:]))
		v6 = t30 - v4
		t31 := int32(load32(m.memory[uint32(v6+i32(-420)):]))
		v4 = t31
		v2 = v4 & i32(-8)
		t32 := v2
		v4 = v4 & i32(3)
		p33 := i32(8)
		if v4 != 0 {
			p33 = i32(4)
		}
		if uint32(t32) < uint32(p33+v3) {
			m.fn7(i32(1274404), i32(46), i32(1274452))
			panic("unreachable")
		}
		if v4 == 0 {
			goto l26
		}
		if uint32(v2) > uint32(v3+i32(39)) {
			m.fn7(i32(1274468), i32(46), i32(1274516))
			panic("unreachable")
		}
	l26:
		m.fn5(v6 + i32(-416))
	}
l16:
	{
		t34 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		if t34 == i32(-1) {
			goto l28
		}
		m.fn420(v0)
	}
l28:
	m.fn421(v0 + i32(464))
}
func (m *Module) fn403(v0 int32) {
	m.fn413(v0 + i32(4))
	m.fn400(v0 + i32(16))
	m.fn401(v0 + i32(56))
}
func (m *Module) fn404(v0, v1, v2, v3, v4, v5, v6, v7, v8 int32) {
	var v9, v10, v11 int32
	var v12, v13 int64
	t0 := m.g0
	v9 = t0 - i32(80)
	m.g0 = v9
	v10 = v7
	v11 = v8
	{
		t1 := m.fn148(v1, v2, v5, v6)
		v6 = t1
		if v6 == 0 {
			goto l0
		}
		t2 := int32(load32(m.memory[int64(uint32(v6))+8:]))
		v11 = t2
		t3 := int32(load32(m.memory[int64(uint32(v6))+4:]))
		v10 = t3
	}
l0:
	m.fn149(v9+i32(12), v3, v4, v10, v11)
	v10 = v9 + i32(12) + i32(4)
	{
		t4 := int32(load32(m.memory[int64(uint32(v9))+12:]))
		if t4 != 0 {
			m.fn149(v9+i32(52), v3, v4, v7, v8)
			{
				t12 := int32(load32(m.memory[int64(uint32(v9))+52:]))
				if t12 != i32(1) {
					{
						t20 := int32(load32(m.memory[int64(uint32(v9))+68:]))
						v8 = t20
						if uint32(v8+i32(-1)) > uint32(i32(-3)) {
							goto l9
						}
						t21 := int32(load32(m.memory[int64(uint32(v9))+72:]))
						v11 = t21
						t22 := int32(load32(m.memory[uint32(v11+i32(-4)):]))
						v7 = t22
						v4 = v7 & i32(-8)
						t23 := v4
						v7 = v7 & i32(3)
						p24 := i32(8)
						if v7 != 0 {
							p24 = i32(4)
						}
						if uint32(t23) < uint32(p24+v8) {
							m.fn7(i32(1274404), i32(46), i32(1274452))
							panic("unreachable")
						}
						if v7 == 0 {
							goto l11
						}
						if uint32(v4) > uint32(v8+i32(39)) {
							m.fn7(i32(1274468), i32(46), i32(1274516))
							panic("unreachable")
						}
					l11:
						m.fn5(v11)
					}
				l9:
					t25 := v9
					v8 = v9 + i32(52) + i32(4)
					t26 := int32(load32(m.memory[int64(uint32(v8))+8:]))
					v7 = t26
					store32(m.memory[int64(uint32(t25))+48:], uint32(v7))
					t27 := int64(load64(m.memory[uint32(v8):]))
					t28 := v9
					v12 = t27
					store64(m.memory[int64(uint32(t28))+40:], uint64(v12))
					store32(m.memory[int64(uint32(v0))+8:], uint32(v7))
					store64(m.memory[uint32(v0):], uint64(v12))
					m.fn143(v10)
					goto l2
				}
				t13 := int64(load64(m.memory[int64(uint32(v9))+60:]))
				t14 := v9
				v12 = t13
				store64(m.memory[int64(uint32(t14))+40:], uint64(v12))
				t15 := int32(load32(m.memory[int64(uint32(v9))+68:]))
				t16 := v9
				v11 = t15
				store32(m.memory[int64(uint32(t16))+48:], uint32(v11))
				t17 := int32(load32(m.memory[int64(uint32(v9))+56:]))
				v4 = t17
				t18 := int64(load64(m.memory[int64(uint32(v9))+72:]))
				v13 = t18
				store32(m.memory[int64(uint32(v9))+64:], uint32(v11))
				store64(m.memory[int64(uint32(v9))+56:], uint64(v12))
				store64(m.memory[int64(uint32(v9))+68:], uint64(v13))
				store32(m.memory[int64(uint32(v9))+52:], uint32(v4))
				t19 := m.fn11(v8)
				v11 = t19
				if v11 == 0 {
					m.fn16(i32(1), v8)
					panic("unreachable")
				}
				if v8 == 0 {
					goto l8
				}
				memory_copy(m.memory, uint32(v11), uint32(v7), uint32(v8))
			l8:
				store32(m.memory[int64(uint32(v0))+8:], uint32(v8))
				store32(m.memory[int64(uint32(v0))+4:], uint32(v11))
				store32(m.memory[uint32(v0):], uint32(v8))
				m.fn143(v9 + i32(52))
				m.fn143(v10)
				goto l2
			}
		}
		t5 := int32(load32(m.memory[int64(uint32(v10))+8:]))
		store32(m.memory[int64(uint32(v0))+8:], uint32(t5))
		t6 := int64(load64(m.memory[uint32(v10):]))
		store64(m.memory[uint32(v0):], uint64(t6))
		t7 := int32(load32(m.memory[int64(uint32(v9))+28:]))
		v8 = t7
		if uint32(v8+i32(-1)) > uint32(i32(-3)) {
			goto l2
		}
		t8 := int32(load32(m.memory[int64(uint32(v9))+32:]))
		v7 = t8
		t9 := int32(load32(m.memory[uint32(v7+i32(-4)):]))
		v0 = t9
		v10 = v0 & i32(-8)
		t10 := v10
		v0 = v0 & i32(3)
		p11 := i32(8)
		if v0 != 0 {
			p11 = i32(4)
		}
		if uint32(t10) < uint32(p11+v8) {
			m.fn7(i32(1274404), i32(46), i32(1274452))
			panic("unreachable")
		}
		if v0 == 0 {
			goto l4
		}
		if uint32(v10) > uint32(v8+i32(39)) {
			m.fn7(i32(1274468), i32(46), i32(1274516))
			panic("unreachable")
		}
	l4:
		m.fn5(v7)
		goto l2
	}
l2:
	m.g0 = v9 + i32(80)
}
func (m *Module) fn405(v0, v1, v2, v3 int32) int32 {
	var v4, v5, v6 int32
	var v7, v8 int64
	var v9 int32
	t0 := m.g0
	v4 = t0 - i32(16)
	m.g0 = v4
	v5 = i32(2)
	{
		if v1 == 0 {
			goto l0
		}
		v1 = v1 * i32(44)
	l5:
		{
			t1 := int32(load32(m.memory[uint32(v0):]))
			if t1 == i32(-1) {
				goto l1
			}
			t2 := int32(load32(m.memory[uint32(v0+i32(8)):]))
			if t2 != v3 {
				goto l1
			}
			t3 := int32(load32(m.memory[uint32(v0+i32(4)):]))
			t4 := m.fn1909(t3, v2, v3)
			if t4 != 0 {
				goto l1
			}
			t5 := int32(load32(m.memory[uint32(v0+i32(36)):]))
			v6 = t5
			if v6 == 0 {
				goto l1
			}
			t6 := int32(load32(m.memory[uint32(v0+i32(40)):]))
			if t6 != i32(60) {
				goto l1
			}
			v7 = i64(0x687474703a2f2f73)
			{
				{
					t7 := int64(load64(m.memory[int64(uint32(v6))+8:]))
					v8 = t7
					v8 = v8<<56 | v8&i64(0xff00)<<40 | (v8&i64(0xff0000)<<24 | v8&i64(0xff000000)<<8) | (int64(uint64(v8)>>8)&i64(0xff000000) | int64(uint64(v8)>>24)&i64(0xff0000) | (int64(uint64(v8)>>40)&i64(0xff00) | int64(uint64(v8)>>56)))
					if v8 != i64(0x687474703a2f2f73) {
						goto l2
					}
					v7 = i64(7163086727793553007)
					t8 := int64(load64(m.memory[uint32(v6+i32(16)):]))
					v8 = t8
					v8 = v8<<56 | v8&i64(0xff00)<<40 | (v8&i64(0xff0000)<<24 | v8&i64(0xff000000)<<8) | (int64(uint64(v8)>>8)&i64(0xff000000) | int64(uint64(v8)>>24)&i64(0xff0000) | (int64(uint64(v8)>>40)&i64(0xff00) | int64(uint64(v8)>>56)))
					if v8 != i64(7163086727793553007) {
						goto l2
					}
					v7 = i64(8099000968406656623)
					t9 := int64(load64(m.memory[uint32(v6+i32(24)):]))
					v8 = t9
					v8 = v8<<56 | v8&i64(0xff00)<<40 | (v8&i64(0xff0000)<<24 | v8&i64(0xff000000)<<8) | (int64(uint64(v8)>>8)&i64(0xff000000) | int64(uint64(v8)>>24)&i64(0xff0000) | (int64(uint64(v8)>>40)&i64(0xff00) | int64(uint64(v8)>>56)))
					if v8 != i64(8099000968406656623) {
						goto l2
					}
					v7 = i64(8245353645561769842)
					t10 := int64(load64(m.memory[uint32(v6+i32(32)):]))
					v8 = t10
					v8 = v8<<56 | v8&i64(0xff00)<<40 | (v8&i64(0xff0000)<<24 | v8&i64(0xff000000)<<8) | (int64(uint64(v8)>>8)&i64(0xff000000) | int64(uint64(v8)>>24)&i64(0xff0000) | (int64(uint64(v8)>>40)&i64(0xff00) | int64(uint64(v8)>>56)))
					if v8 != i64(8245353645561769842) {
						goto l2
					}
					v7 = i64(0x672f776f72647072)
					t11 := int64(load64(m.memory[uint32(v6+i32(40)):]))
					v8 = t11
					v8 = v8<<56 | v8&i64(0xff00)<<40 | (v8&i64(0xff0000)<<24 | v8&i64(0xff000000)<<8) | (int64(uint64(v8)>>8)&i64(0xff000000) | int64(uint64(v8)>>24)&i64(0xff0000) | (int64(uint64(v8)>>40)&i64(0xff00) | int64(uint64(v8)>>56)))
					if v8 != i64(0x672f776f72647072) {
						goto l2
					}
					v7 = i64(0x6f63657373696e67)
					t12 := int64(load64(m.memory[uint32(v6+i32(48)):]))
					v8 = t12
					v8 = v8<<56 | v8&i64(0xff00)<<40 | (v8&i64(0xff0000)<<24 | v8&i64(0xff000000)<<8) | (int64(uint64(v8)>>8)&i64(0xff000000) | int64(uint64(v8)>>24)&i64(0xff0000) | (int64(uint64(v8)>>40)&i64(0xff00) | int64(uint64(v8)>>56)))
					if v8 != i64(0x6f63657373696e67) {
						goto l2
					}
					v7 = i64(7884728940222232111)
					t13 := int64(load64(m.memory[uint32(v6+i32(56)):]))
					v8 = t13
					v8 = v8<<56 | v8&i64(0xff00)<<40 | (v8&i64(0xff0000)<<24 | v8&i64(0xff000000)<<8) | (int64(uint64(v8)>>8)&i64(0xff000000) | int64(uint64(v8)>>24)&i64(0xff0000) | (int64(uint64(v8)>>40)&i64(0xff00) | int64(uint64(v8)>>56)))
					if v8 != i64(7884728940222232111) {
						goto l2
					}
					v9 = i32(0)
					t14 := int32(load32(m.memory[uint32(v6+i32(64)):]))
					v6 = t14
					v6 = i32_rotr(v6&i32(0xff00ff), i32(8)) | i32_rotr(v6, i32(24))&i32(0xff00ff)
					if v6 == i32(1835100526) {
						goto l3
					}
					v8 = int64(uint32(v6))
					v7 = i64(1835100526)
				}
			l2:
				p15 := i32(1)
				if uint64(v8) < uint64(v7) {
					p15 = i32(-1)
				}
				v9 = p15
			}
		l3:
			if v9 == 0 {
				goto l4
			}
		}
	l1:
		v0 = v0 + i32(44)
		v1 = v1 + i32(-44)
		if v1 != 0 {
			goto l5
		}
		goto l0
	l4:
		t16 := int32(load32(m.memory[uint32(v0+i32(16)):]))
		t17 := int32(load32(m.memory[uint32(v0+i32(20)):]))
		m.fn155(v4+i32(8), t16, t17, i32(1070016), i32(60), i32(1070079), i32(3))
		v5 = i32(1)
		t18 := int32(load32(m.memory[int64(uint32(v4))+8:]))
		v0 = t18
		if v0 == 0 {
			goto l0
		}
		{
			t19 := int32(load32(m.memory[int64(uint32(v4))+12:]))
			switch t19 + i32(-1) {
			default:
				goto l0
			case 0:
				t20 := int32(m.memory[uint32(v0)])
				var p21 int32
				if t20 != i32(48) {
					p21 = 1
				}
				v5 = p21
				goto l0
			case 4:
				t22 := int32(load32(m.memory[uint32(v0):]))
				t23 := int32(m.memory[uint32(v0+i32(4))])
				var p24 int32
				if t22^i32(1936482662)|(t23^i32(101)) != i32(0) {
					p24 = 1
				}
				v5 = p24
				goto l0
			case 2:
				t25 := int32(load16(m.memory[uint32(v0):]))
				t26 := int32(m.memory[uint32(v0+i32(2))])
				var p27 int32
				if (t25^i32(26223)|(t26^i32(102)))&i32(0xffff) != i32(0) {
					p27 = 1
				}
				v5 = p27
				goto l0
			case 3:
				t28 := int32(load32(m.memory[uint32(v0):]))
				var p29 int32
				if t28 != i32(1701736302) {
					p29 = 1
				}
				v5 = p29
			}
		}
	}
l0:
	m.g0 = v4 + i32(16)
	return v5
}
func (m *Module) fn406(v0, v1, v2 int32) {
	var v3, v4, v5, v6, v7, v8 int32
	var v9, v10 int64
	var v11, v12, v13 int32
	var v14 int64
	var v15, v16, v17, v18, v19, v20, v21, v22, v23, v24, v25 int32
	t0 := m.g0
	v3 = t0 - i32(64)
	m.g0 = v3
	v4 = v2 * i32(44)
	v5 = i32(1075328)
	{
		if v2 == 0 {
			goto l0
		}
		v6 = v4
		v7 = v1
	l5:
		{
			t1 := int32(load32(m.memory[uint32(v7):]))
			if t1 == i32(-1) {
				goto l1
			}
			t2 := int32(load32(m.memory[uint32(v7+i32(8)):]))
			if t2 != i32(6) {
				goto l1
			}
			t3 := int32(load32(m.memory[uint32(v7+i32(4)):]))
			v8 = t3
			t4 := int32(load32(m.memory[uint32(v8):]))
			t5 := int32(load16(m.memory[uint32(v8+i32(4)):]))
			if t4^i32(1181578606)|(t5^i32(29805)) != 0 {
				goto l1
			}
			t6 := int32(load32(m.memory[uint32(v7+i32(36)):]))
			v8 = t6
			if v8 == 0 {
				goto l1
			}
			t7 := int32(load32(m.memory[uint32(v7+i32(40)):]))
			if t7 != i32(60) {
				goto l1
			}
			v9 = i64(0x687474703a2f2f73)
			{
				{
					t8 := int64(load64(m.memory[int64(uint32(v8))+8:]))
					v10 = t8
					v10 = v10<<56 | v10&i64(0xff00)<<40 | (v10&i64(0xff0000)<<24 | v10&i64(0xff000000)<<8) | (int64(uint64(v10)>>8)&i64(0xff000000) | int64(uint64(v10)>>24)&i64(0xff0000) | (int64(uint64(v10)>>40)&i64(0xff00) | int64(uint64(v10)>>56)))
					if v10 != i64(0x687474703a2f2f73) {
						goto l2
					}
					v9 = i64(7163086727793553007)
					t9 := int64(load64(m.memory[uint32(v8+i32(16)):]))
					v10 = t9
					v10 = v10<<56 | v10&i64(0xff00)<<40 | (v10&i64(0xff0000)<<24 | v10&i64(0xff000000)<<8) | (int64(uint64(v10)>>8)&i64(0xff000000) | int64(uint64(v10)>>24)&i64(0xff0000) | (int64(uint64(v10)>>40)&i64(0xff00) | int64(uint64(v10)>>56)))
					if v10 != i64(7163086727793553007) {
						goto l2
					}
					v9 = i64(8099000968406656623)
					t10 := int64(load64(m.memory[uint32(v8+i32(24)):]))
					v10 = t10
					v10 = v10<<56 | v10&i64(0xff00)<<40 | (v10&i64(0xff0000)<<24 | v10&i64(0xff000000)<<8) | (int64(uint64(v10)>>8)&i64(0xff000000) | int64(uint64(v10)>>24)&i64(0xff0000) | (int64(uint64(v10)>>40)&i64(0xff00) | int64(uint64(v10)>>56)))
					if v10 != i64(8099000968406656623) {
						goto l2
					}
					v9 = i64(8245353645561769842)
					t11 := int64(load64(m.memory[uint32(v8+i32(32)):]))
					v10 = t11
					v10 = v10<<56 | v10&i64(0xff00)<<40 | (v10&i64(0xff0000)<<24 | v10&i64(0xff000000)<<8) | (int64(uint64(v10)>>8)&i64(0xff000000) | int64(uint64(v10)>>24)&i64(0xff0000) | (int64(uint64(v10)>>40)&i64(0xff00) | int64(uint64(v10)>>56)))
					if v10 != i64(8245353645561769842) {
						goto l2
					}
					v9 = i64(0x672f776f72647072)
					t12 := int64(load64(m.memory[uint32(v8+i32(40)):]))
					v10 = t12
					v10 = v10<<56 | v10&i64(0xff00)<<40 | (v10&i64(0xff0000)<<24 | v10&i64(0xff000000)<<8) | (int64(uint64(v10)>>8)&i64(0xff000000) | int64(uint64(v10)>>24)&i64(0xff0000) | (int64(uint64(v10)>>40)&i64(0xff00) | int64(uint64(v10)>>56)))
					if v10 != i64(0x672f776f72647072) {
						goto l2
					}
					v9 = i64(0x6f63657373696e67)
					t13 := int64(load64(m.memory[uint32(v8+i32(48)):]))
					v10 = t13
					v10 = v10<<56 | v10&i64(0xff00)<<40 | (v10&i64(0xff0000)<<24 | v10&i64(0xff000000)<<8) | (int64(uint64(v10)>>8)&i64(0xff000000) | int64(uint64(v10)>>24)&i64(0xff0000) | (int64(uint64(v10)>>40)&i64(0xff00) | int64(uint64(v10)>>56)))
					if v10 != i64(0x6f63657373696e67) {
						goto l2
					}
					v9 = i64(7884728940222232111)
					t14 := int64(load64(m.memory[uint32(v8+i32(56)):]))
					v10 = t14
					v10 = v10<<56 | v10&i64(0xff00)<<40 | (v10&i64(0xff0000)<<24 | v10&i64(0xff000000)<<8) | (int64(uint64(v10)>>8)&i64(0xff000000) | int64(uint64(v10)>>24)&i64(0xff0000) | (int64(uint64(v10)>>40)&i64(0xff00) | int64(uint64(v10)>>56)))
					if v10 != i64(7884728940222232111) {
						goto l2
					}
					v11 = i32(0)
					t15 := int32(load32(m.memory[uint32(v8+i32(64)):]))
					v8 = t15
					v8 = i32_rotr(v8&i32(0xff00ff), i32(8)) | i32_rotr(v8, i32(24))&i32(0xff00ff)
					if v8 == i32(1835100526) {
						goto l3
					}
					v10 = int64(uint32(v8))
					v9 = i64(1835100526)
				}
			l2:
				p16 := i32(1)
				if uint64(v10) < uint64(v9) {
					p16 = i32(-1)
				}
				v11 = p16
			}
		l3:
			if v11 == 0 {
				t17 := int32(load32(m.memory[uint32(v7+i32(16)):]))
				t18 := int32(load32(m.memory[uint32(v7+i32(20)):]))
				m.fn155(v3+i32(32), t17, t18, i32(1070016), i32(60), i32(1070079), i32(3))
				t19 := int32(load32(m.memory[int64(uint32(v3))+32:]))
				v7 = t19
				if v7 == 0 {
					goto l0
				}
				v11 = i32(1)
				v12 = i32(0)
				v5 = v7
				{
					t20 := int32(load32(m.memory[int64(uint32(v3))+36:]))
					switch t20 + i32(-4) {
					case 2:
						goto l0
					default:
						goto l7
					case 0:
						t21 := int32(load32(m.memory[uint32(v7):]))
						if t21 != i32(1701736302) {
							goto l10
						}
						v11 = i32(255)
						v12 = i32(1)
						goto l7
					case 7:
						v11 = i32(3)
						{
							t22 := int64(load64(m.memory[uint32(v7):]))
							t23 := t22 ^ i64(0x74654c7265776f6c)
							v6 = v7 + i32(3)
							t24 := int64(load64(m.memory[uint32(v6):]))
							if !(t23|(t24^i64(8243122736236098149)) == 0) {
								t25 := int64(load64(m.memory[uint32(v7):]))
								t26 := int64(load64(m.memory[uint32(v6):]))
								if t25^i64(0x74654c7265707075)|(t26^i64(8243122736236098149)) == 0 {
									goto l7
								}
								goto l10
							}
							v11 = i32(2)
							goto l7
						}
					case 6:
						t27 := int64(load64(m.memory[uint32(v7):]))
						t28 := t27 ^ i64(7885612123831103340)
						v6 = v7 + i32(8)
						t29 := int64(load16(m.memory[uint32(v6):]))
						if !(t28|(t29^i64(28257)) == 0) {
							t30 := int64(load64(m.memory[uint32(v7):]))
							t31 := int64(load16(m.memory[uint32(v6):]))
							p32 := i32(1)
							if t30^i64(7885612123830644853)|(t31^i64(28257)) == 0 {
								p32 = i32(5)
							}
							v11 = p32
							goto l7
						}
						v11 = i32(4)
						goto l7
					}
				}
			}
		}
	l1:
		v7 = v7 + i32(44)
		v6 = v6 + i32(-44)
		if v6 != 0 {
			goto l5
		}
		goto l0
	l0:
		t33 := int32(load32(m.memory[uint32(v5):]))
		t34 := int32(load16(m.memory[uint32(v5+i32(4)):]))
		if t33^i32(1819047266)|(t34^i32(29797)) != 0 {
			goto l10
		}
		v11 = i32(0)
		v12 = i32(1)
		goto l7
	}
l10:
	v11 = i32(1)
	v12 = i32(0)
l7:
	v13 = i32(0)
	if v2 != 0 {
		goto l13
	}
	v6 = i32(4)
	v14 = i64(1)
	v8 = i32(0)
	v7 = i32(0)
	goto l14
l13:
	v6 = v4
	v7 = v1
	{
	l19:
		{
			t35 := int32(load32(m.memory[uint32(v7):]))
			if t35 == i32(-1) {
				goto l15
			}
			t36 := int32(load32(m.memory[uint32(v7+i32(8)):]))
			if t36 != i32(5) {
				goto l15
			}
			t37 := int32(load32(m.memory[uint32(v7+i32(4)):]))
			v8 = t37
			t38 := int32(load32(m.memory[uint32(v8):]))
			t39 := int32(m.memory[uint32(v8+i32(4))])
			if t38^i32(1918989427)|(t39^i32(116)) != 0 {
				goto l15
			}
			t40 := int32(load32(m.memory[uint32(v7+i32(36)):]))
			v8 = t40
			if v8 == 0 {
				goto l15
			}
			t41 := int32(load32(m.memory[uint32(v7+i32(40)):]))
			if t41 != i32(60) {
				goto l15
			}
			v9 = i64(0x687474703a2f2f73)
			{
				{
					t42 := int64(load64(m.memory[int64(uint32(v8))+8:]))
					v10 = t42
					v10 = v10<<56 | v10&i64(0xff00)<<40 | (v10&i64(0xff0000)<<24 | v10&i64(0xff000000)<<8) | (int64(uint64(v10)>>8)&i64(0xff000000) | int64(uint64(v10)>>24)&i64(0xff0000) | (int64(uint64(v10)>>40)&i64(0xff00) | int64(uint64(v10)>>56)))
					if v10 != i64(0x687474703a2f2f73) {
						goto l16
					}
					v9 = i64(7163086727793553007)
					t43 := int64(load64(m.memory[uint32(v8+i32(16)):]))
					v10 = t43
					v10 = v10<<56 | v10&i64(0xff00)<<40 | (v10&i64(0xff0000)<<24 | v10&i64(0xff000000)<<8) | (int64(uint64(v10)>>8)&i64(0xff000000) | int64(uint64(v10)>>24)&i64(0xff0000) | (int64(uint64(v10)>>40)&i64(0xff00) | int64(uint64(v10)>>56)))
					if v10 != i64(7163086727793553007) {
						goto l16
					}
					v9 = i64(8099000968406656623)
					t44 := int64(load64(m.memory[uint32(v8+i32(24)):]))
					v10 = t44
					v10 = v10<<56 | v10&i64(0xff00)<<40 | (v10&i64(0xff0000)<<24 | v10&i64(0xff000000)<<8) | (int64(uint64(v10)>>8)&i64(0xff000000) | int64(uint64(v10)>>24)&i64(0xff0000) | (int64(uint64(v10)>>40)&i64(0xff00) | int64(uint64(v10)>>56)))
					if v10 != i64(8099000968406656623) {
						goto l16
					}
					v9 = i64(8245353645561769842)
					t45 := int64(load64(m.memory[uint32(v8+i32(32)):]))
					v10 = t45
					v10 = v10<<56 | v10&i64(0xff00)<<40 | (v10&i64(0xff0000)<<24 | v10&i64(0xff000000)<<8) | (int64(uint64(v10)>>8)&i64(0xff000000) | int64(uint64(v10)>>24)&i64(0xff0000) | (int64(uint64(v10)>>40)&i64(0xff00) | int64(uint64(v10)>>56)))
					if v10 != i64(8245353645561769842) {
						goto l16
					}
					v9 = i64(0x672f776f72647072)
					t46 := int64(load64(m.memory[uint32(v8+i32(40)):]))
					v10 = t46
					v10 = v10<<56 | v10&i64(0xff00)<<40 | (v10&i64(0xff0000)<<24 | v10&i64(0xff000000)<<8) | (int64(uint64(v10)>>8)&i64(0xff000000) | int64(uint64(v10)>>24)&i64(0xff0000) | (int64(uint64(v10)>>40)&i64(0xff00) | int64(uint64(v10)>>56)))
					if v10 != i64(0x672f776f72647072) {
						goto l16
					}
					v9 = i64(0x6f63657373696e67)
					t47 := int64(load64(m.memory[uint32(v8+i32(48)):]))
					v10 = t47
					v10 = v10<<56 | v10&i64(0xff00)<<40 | (v10&i64(0xff0000)<<24 | v10&i64(0xff000000)<<8) | (int64(uint64(v10)>>8)&i64(0xff000000) | int64(uint64(v10)>>24)&i64(0xff0000) | (int64(uint64(v10)>>40)&i64(0xff00) | int64(uint64(v10)>>56)))
					if v10 != i64(0x6f63657373696e67) {
						goto l16
					}
					v9 = i64(7884728940222232111)
					t48 := int64(load64(m.memory[uint32(v8+i32(56)):]))
					v10 = t48
					v10 = v10<<56 | v10&i64(0xff00)<<40 | (v10&i64(0xff0000)<<24 | v10&i64(0xff000000)<<8) | (int64(uint64(v10)>>8)&i64(0xff000000) | int64(uint64(v10)>>24)&i64(0xff0000) | (int64(uint64(v10)>>40)&i64(0xff00) | int64(uint64(v10)>>56)))
					if v10 != i64(7884728940222232111) {
						goto l16
					}
					v5 = i32(0)
					t49 := int32(load32(m.memory[uint32(v8+i32(64)):]))
					v8 = t49
					v8 = i32_rotr(v8&i32(0xff00ff), i32(8)) | i32_rotr(v8, i32(24))&i32(0xff00ff)
					if v8 == i32(1835100526) {
						goto l17
					}
					v10 = int64(uint32(v8))
					v9 = i64(1835100526)
				}
			l16:
				p50 := i32(1)
				if uint64(v10) < uint64(v9) {
					p50 = i32(-1)
				}
				v5 = p50
			}
		l17:
			if v5 == 0 {
				goto l18
			}
		}
	l15:
		v7 = v7 + i32(44)
		v6 = v6 + i32(-44)
		if v6 != 0 {
			goto l19
		}
		v14 = i64(1)
		goto l20
	l18:
		t51 := int32(load32(m.memory[uint32(v7+i32(16)):]))
		t52 := int32(load32(m.memory[uint32(v7+i32(20)):]))
		m.fn155(v3+i32(24), t51, t52, i32(1070016), i32(60), i32(1070079), i32(3))
		v14 = i64(1)
		t53 := int32(load32(m.memory[int64(uint32(v3))+24:]))
		v7 = t53
		if v7 == 0 {
			goto l20
		}
		t54 := int32(load32(m.memory[int64(uint32(v3))+28:]))
		m.fn409(v3+i32(40), v7, t54)
		t55 := int32(m.memory[int64(uint32(v3))+40])
		if t55 == i32(1) {
			goto l20
		}
		t56 := int64(load64(m.memory[int64(uint32(v3))+48:]))
		v10 = t56
		p57 := i64(0)
		if v10 > i64(0) {
			p57 = v10
		}
		v10 = p57
		p58 := i64(0x7fffffff)
		if v10 < i64(0x7fffffff) {
			p58 = v10
		}
		v14 = p58
	}
l20:
	v6 = v4
	v7 = v1
	{
	l25:
		{
			t59 := int32(load32(m.memory[uint32(v7):]))
			if t59 == i32(-1) {
				goto l21
			}
			t60 := int32(load32(m.memory[uint32(v7+i32(8)):]))
			if t60 != i32(10) {
				goto l21
			}
			t61 := int32(load32(m.memory[uint32(v7+i32(4)):]))
			v8 = t61
			t62 := int64(load64(m.memory[uint32(v8):]))
			t63 := int64(load16(m.memory[uint32(v8+i32(8)):]))
			if t62^i64(0x61747365526c766c)|(t63^i64(29810)) != i64(0) {
				goto l21
			}
			t64 := int32(load32(m.memory[uint32(v7+i32(36)):]))
			v8 = t64
			if v8 == 0 {
				goto l21
			}
			t65 := int32(load32(m.memory[uint32(v7+i32(40)):]))
			if t65 != i32(60) {
				goto l21
			}
			v9 = i64(0x687474703a2f2f73)
			{
				{
					t66 := int64(load64(m.memory[int64(uint32(v8))+8:]))
					v10 = t66
					v10 = v10<<56 | v10&i64(0xff00)<<40 | (v10&i64(0xff0000)<<24 | v10&i64(0xff000000)<<8) | (int64(uint64(v10)>>8)&i64(0xff000000) | int64(uint64(v10)>>24)&i64(0xff0000) | (int64(uint64(v10)>>40)&i64(0xff00) | int64(uint64(v10)>>56)))
					if v10 != i64(0x687474703a2f2f73) {
						goto l22
					}
					v9 = i64(7163086727793553007)
					t67 := int64(load64(m.memory[uint32(v8+i32(16)):]))
					v10 = t67
					v10 = v10<<56 | v10&i64(0xff00)<<40 | (v10&i64(0xff0000)<<24 | v10&i64(0xff000000)<<8) | (int64(uint64(v10)>>8)&i64(0xff000000) | int64(uint64(v10)>>24)&i64(0xff0000) | (int64(uint64(v10)>>40)&i64(0xff00) | int64(uint64(v10)>>56)))
					if v10 != i64(7163086727793553007) {
						goto l22
					}
					v9 = i64(8099000968406656623)
					t68 := int64(load64(m.memory[uint32(v8+i32(24)):]))
					v10 = t68
					v10 = v10<<56 | v10&i64(0xff00)<<40 | (v10&i64(0xff0000)<<24 | v10&i64(0xff000000)<<8) | (int64(uint64(v10)>>8)&i64(0xff000000) | int64(uint64(v10)>>24)&i64(0xff0000) | (int64(uint64(v10)>>40)&i64(0xff00) | int64(uint64(v10)>>56)))
					if v10 != i64(8099000968406656623) {
						goto l22
					}
					v9 = i64(8245353645561769842)
					t69 := int64(load64(m.memory[uint32(v8+i32(32)):]))
					v10 = t69
					v10 = v10<<56 | v10&i64(0xff00)<<40 | (v10&i64(0xff0000)<<24 | v10&i64(0xff000000)<<8) | (int64(uint64(v10)>>8)&i64(0xff000000) | int64(uint64(v10)>>24)&i64(0xff0000) | (int64(uint64(v10)>>40)&i64(0xff00) | int64(uint64(v10)>>56)))
					if v10 != i64(8245353645561769842) {
						goto l22
					}
					v9 = i64(0x672f776f72647072)
					t70 := int64(load64(m.memory[uint32(v8+i32(40)):]))
					v10 = t70
					v10 = v10<<56 | v10&i64(0xff00)<<40 | (v10&i64(0xff0000)<<24 | v10&i64(0xff000000)<<8) | (int64(uint64(v10)>>8)&i64(0xff000000) | int64(uint64(v10)>>24)&i64(0xff0000) | (int64(uint64(v10)>>40)&i64(0xff00) | int64(uint64(v10)>>56)))
					if v10 != i64(0x672f776f72647072) {
						goto l22
					}
					v9 = i64(0x6f63657373696e67)
					t71 := int64(load64(m.memory[uint32(v8+i32(48)):]))
					v10 = t71
					v10 = v10<<56 | v10&i64(0xff00)<<40 | (v10&i64(0xff0000)<<24 | v10&i64(0xff000000)<<8) | (int64(uint64(v10)>>8)&i64(0xff000000) | int64(uint64(v10)>>24)&i64(0xff0000) | (int64(uint64(v10)>>40)&i64(0xff00) | int64(uint64(v10)>>56)))
					if v10 != i64(0x6f63657373696e67) {
						goto l22
					}
					v9 = i64(7884728940222232111)
					t72 := int64(load64(m.memory[uint32(v8+i32(56)):]))
					v10 = t72
					v10 = v10<<56 | v10&i64(0xff00)<<40 | (v10&i64(0xff0000)<<24 | v10&i64(0xff000000)<<8) | (int64(uint64(v10)>>8)&i64(0xff000000) | int64(uint64(v10)>>24)&i64(0xff0000) | (int64(uint64(v10)>>40)&i64(0xff00) | int64(uint64(v10)>>56)))
					if v10 != i64(7884728940222232111) {
						goto l22
					}
					v5 = i32(0)
					t73 := int32(load32(m.memory[uint32(v8+i32(64)):]))
					v8 = t73
					v8 = i32_rotr(v8&i32(0xff00ff), i32(8)) | i32_rotr(v8, i32(24))&i32(0xff00ff)
					if v8 == i32(1835100526) {
						goto l23
					}
					v10 = int64(uint32(v8))
					v9 = i64(1835100526)
				}
			l22:
				p74 := i32(1)
				if uint64(v10) < uint64(v9) {
					p74 = i32(-1)
				}
				v5 = p74
			}
		l23:
			if v5 == 0 {
				goto l24
			}
		}
	l21:
		v7 = v7 + i32(44)
		v6 = v6 + i32(-44)
		if v6 != 0 {
			goto l25
		}
		v8 = i32(0)
		goto l26
	l24:
		t75 := int32(load32(m.memory[uint32(v7+i32(16)):]))
		t76 := int32(load32(m.memory[uint32(v7+i32(20)):]))
		m.fn155(v3+i32(16), t75, t76, i32(1070016), i32(60), i32(1070079), i32(3))
		v8 = i32(0)
		{
			t77 := int32(load32(m.memory[int64(uint32(v3))+16:]))
			v7 = t77
			if v7 != 0 {
				goto l27
			}
			goto l26
		}
	l27:
		t78 := int32(load32(m.memory[int64(uint32(v3))+20:]))
		v6 = t78
		v8 = v6
		switch v6 {
		case 0:
			goto l26
		case 1:
			v8 = i32(0)
			t79 := int32(m.memory[uint32(v7)])
			v15 = t79
			switch v15 + i32(-43) {
			case 0, 2:
				goto l26
			default:
				goto l30
			}
		default:
			t80 := int32(m.memory[uint32(v7)])
			v15 = t80
		}
	l30:
		t81 := v7
		var p82 int32
		if v15&i32(255) == i32(43) {
			p82 = 1
		}
		v8 = p82
		v15 = t81 + v8
		{
			v7 = v6 - v8
			if uint32(v7) < uint32(i32(9)) {
				goto l31
			}
			v5 = i32(0)
		l35:
			if v7 == 0 {
				goto l32
			}
			v8 = i32(0)
			v10 = int64(uint32(v5)) * i64(10)
			if int32(int64(uint64(v10)>>32)) == 0 {
				t83 := int32(m.memory[uint32(v15)])
				v6 = t83 + i32(-48)
				if uint32(v6) <= uint32(i32(9)) {
					v15 = v15 + i32(1)
					v7 = v7 + i32(-1)
					v5 = v6 + int32(v10)
					if uint32(v5) >= uint32(v6) {
						goto l35
					}
					goto l26
				}
				goto l26
			}
			goto l26
		l31:
			if v7 != 0 {
				goto l36
			}
			v5 = i32(0)
			goto l32
		l36:
			v8 = i32(0)
			{
				t84 := int32(m.memory[uint32(v15)])
				v5 = t84 + i32(-48)
				if uint32(v5) <= uint32(i32(9)) {
					goto l37
				}
				goto l26
			}
		l37:
			if v7 == i32(1) {
				goto l32
			}
			{
				t85 := int32(m.memory[int64(uint32(v15))+1])
				v6 = t85 + i32(-48)
				if uint32(v6) <= uint32(i32(9)) {
					goto l38
				}
				goto l26
			}
		l38:
			v5 = v6 + v5*i32(10)
			if v7 == i32(2) {
				goto l32
			}
			{
				t86 := int32(m.memory[int64(uint32(v15))+2])
				v6 = t86 + i32(-48)
				if uint32(v6) <= uint32(i32(9)) {
					goto l39
				}
				goto l26
			}
		l39:
			v5 = v6 + v5*i32(10)
			if v7 == i32(3) {
				goto l32
			}
			{
				t87 := int32(m.memory[int64(uint32(v15))+3])
				v6 = t87 + i32(-48)
				if uint32(v6) <= uint32(i32(9)) {
					goto l40
				}
				goto l26
			}
		l40:
			v5 = v6 + v5*i32(10)
			if v7 == i32(4) {
				goto l32
			}
			{
				t88 := int32(m.memory[int64(uint32(v15))+4])
				v6 = t88 + i32(-48)
				if uint32(v6) <= uint32(i32(9)) {
					goto l41
				}
				goto l26
			}
		l41:
			v5 = v6 + v5*i32(10)
			if v7 == i32(5) {
				goto l32
			}
			{
				t89 := int32(m.memory[int64(uint32(v15))+5])
				v6 = t89 + i32(-48)
				if uint32(v6) <= uint32(i32(9)) {
					goto l42
				}
				goto l26
			}
		l42:
			v5 = v6 + v5*i32(10)
			if v7 == i32(6) {
				goto l32
			}
			{
				t90 := int32(m.memory[int64(uint32(v15))+6])
				v6 = t90 + i32(-48)
				if uint32(v6) <= uint32(i32(9)) {
					goto l43
				}
				goto l26
			}
		l43:
			v5 = v6 + v5*i32(10)
			if v7 == i32(7) {
				goto l32
			}
			t91 := int32(m.memory[int64(uint32(v15))+7])
			v7 = t91 + i32(-48)
			if uint32(v7) > uint32(i32(9)) {
				goto l26
			}
			v5 = v7 + v5*i32(10)
		}
	l32:
		v8 = i32(1)
	}
l26:
	if v12 != 0 {
		goto l44
	}
	v7 = v1
l49:
	{
		t92 := int32(load32(m.memory[uint32(v7):]))
		if t92 == i32(-1) {
			goto l45
		}
		t93 := int32(load32(m.memory[uint32(v7+i32(8)):]))
		if t93 != i32(7) {
			goto l45
		}
		t94 := int32(load32(m.memory[uint32(v7+i32(4)):]))
		v6 = t94
		t95 := int32(load32(m.memory[uint32(v6):]))
		t96 := int32(load32(m.memory[uint32(v6+i32(3)):]))
		if t95^i32(0x546c766c)|(t96^i32(0x74786554)) != 0 {
			goto l45
		}
		t97 := int32(load32(m.memory[uint32(v7+i32(36)):]))
		v6 = t97
		if v6 == 0 {
			goto l45
		}
		t98 := int32(load32(m.memory[uint32(v7+i32(40)):]))
		if t98 != i32(60) {
			goto l45
		}
		v9 = i64(0x687474703a2f2f73)
		{
			{
				t99 := int64(load64(m.memory[int64(uint32(v6))+8:]))
				v10 = t99
				v10 = v10<<56 | v10&i64(0xff00)<<40 | (v10&i64(0xff0000)<<24 | v10&i64(0xff000000)<<8) | (int64(uint64(v10)>>8)&i64(0xff000000) | int64(uint64(v10)>>24)&i64(0xff0000) | (int64(uint64(v10)>>40)&i64(0xff00) | int64(uint64(v10)>>56)))
				if v10 != i64(0x687474703a2f2f73) {
					goto l46
				}
				v9 = i64(7163086727793553007)
				t100 := int64(load64(m.memory[uint32(v6+i32(16)):]))
				v10 = t100
				v10 = v10<<56 | v10&i64(0xff00)<<40 | (v10&i64(0xff0000)<<24 | v10&i64(0xff000000)<<8) | (int64(uint64(v10)>>8)&i64(0xff000000) | int64(uint64(v10)>>24)&i64(0xff0000) | (int64(uint64(v10)>>40)&i64(0xff00) | int64(uint64(v10)>>56)))
				if v10 != i64(7163086727793553007) {
					goto l46
				}
				v9 = i64(8099000968406656623)
				t101 := int64(load64(m.memory[uint32(v6+i32(24)):]))
				v10 = t101
				v10 = v10<<56 | v10&i64(0xff00)<<40 | (v10&i64(0xff0000)<<24 | v10&i64(0xff000000)<<8) | (int64(uint64(v10)>>8)&i64(0xff000000) | int64(uint64(v10)>>24)&i64(0xff0000) | (int64(uint64(v10)>>40)&i64(0xff00) | int64(uint64(v10)>>56)))
				if v10 != i64(8099000968406656623) {
					goto l46
				}
				v9 = i64(8245353645561769842)
				t102 := int64(load64(m.memory[uint32(v6+i32(32)):]))
				v10 = t102
				v10 = v10<<56 | v10&i64(0xff00)<<40 | (v10&i64(0xff0000)<<24 | v10&i64(0xff000000)<<8) | (int64(uint64(v10)>>8)&i64(0xff000000) | int64(uint64(v10)>>24)&i64(0xff0000) | (int64(uint64(v10)>>40)&i64(0xff00) | int64(uint64(v10)>>56)))
				if v10 != i64(8245353645561769842) {
					goto l46
				}
				v9 = i64(0x672f776f72647072)
				t103 := int64(load64(m.memory[uint32(v6+i32(40)):]))
				v10 = t103
				v10 = v10<<56 | v10&i64(0xff00)<<40 | (v10&i64(0xff0000)<<24 | v10&i64(0xff000000)<<8) | (int64(uint64(v10)>>8)&i64(0xff000000) | int64(uint64(v10)>>24)&i64(0xff0000) | (int64(uint64(v10)>>40)&i64(0xff00) | int64(uint64(v10)>>56)))
				if v10 != i64(0x672f776f72647072) {
					goto l46
				}
				v9 = i64(0x6f63657373696e67)
				t104 := int64(load64(m.memory[uint32(v6+i32(48)):]))
				v10 = t104
				v10 = v10<<56 | v10&i64(0xff00)<<40 | (v10&i64(0xff0000)<<24 | v10&i64(0xff000000)<<8) | (int64(uint64(v10)>>8)&i64(0xff000000) | int64(uint64(v10)>>24)&i64(0xff0000) | (int64(uint64(v10)>>40)&i64(0xff00) | int64(uint64(v10)>>56)))
				if v10 != i64(0x6f63657373696e67) {
					goto l46
				}
				v9 = i64(7884728940222232111)
				t105 := int64(load64(m.memory[uint32(v6+i32(56)):]))
				v10 = t105
				v10 = v10<<56 | v10&i64(0xff00)<<40 | (v10&i64(0xff0000)<<24 | v10&i64(0xff000000)<<8) | (int64(uint64(v10)>>8)&i64(0xff000000) | int64(uint64(v10)>>24)&i64(0xff0000) | (int64(uint64(v10)>>40)&i64(0xff00) | int64(uint64(v10)>>56)))
				if v10 != i64(7884728940222232111) {
					goto l46
				}
				v12 = i32(0)
				t106 := int32(load32(m.memory[uint32(v6+i32(64)):]))
				v6 = t106
				v6 = i32_rotr(v6&i32(0xff00ff), i32(8)) | i32_rotr(v6, i32(24))&i32(0xff00ff)
				if v6 == i32(1835100526) {
					goto l47
				}
				v10 = int64(uint32(v6))
				v9 = i64(1835100526)
			}
		l46:
			p107 := i32(1)
			if uint64(v10) < uint64(v9) {
				p107 = i32(-1)
			}
			v12 = p107
		}
	l47:
		if v12 == 0 {
			t108 := int32(load32(m.memory[uint32(v7+i32(16)):]))
			t109 := int32(load32(m.memory[uint32(v7+i32(20)):]))
			m.fn155(v3+i32(8), t108, t109, i32(1070016), i32(60), i32(1070079), i32(3))
			t110 := int32(load32(m.memory[int64(uint32(v3))+8:]))
			v4 = t110
			if v4 == 0 {
				goto l44
			}
			t111 := int32(load32(m.memory[int64(uint32(v3))+12:]))
			v7 = t111
			store32(m.memory[int64(uint32(v3))+48:], uint32(i32(0)))
			store64(m.memory[int64(uint32(v3))+40:], uint64(i64(0x400000000)))
			v16 = v4 + v7
			v6 = i32(-2)
			v7 = i32(0)
			v13 = i32(4)
		l64:
			{
				t112 := int32(load32(m.memory[int64(uint32(v3))+44:]))
				v17 = t112
				t113 := v17
				v18 = v7 * i32(12)
				v12 = t113 + v18
				v19 = v12 + i32(-8)
				v20 = v12 + i32(-4)
				v21 = v12 + i32(-12)
			l78:
				{
					switch v6 + i32(2) {
					case 1:
						goto l51
					default:
						goto l52
					case 0:
						if v4 == v16 {
							goto l51
						}
						{
							t114 := int32(int8(m.memory[uint32(v4)]))
							v6 = t114
							if v6 <= i32(-1) {
								goto l53
							}
							v4 = v4 + i32(1)
							v6 = v6 & i32(255)
							goto l52
						}
					l53:
						t115 := int32(m.memory[int64(uint32(v4))+1])
						v12 = t115 & i32(63)
						v15 = v6 & i32(31)
						if uint32(v6) > uint32(i32(-33)) {
							goto l54
						}
						v6 = v15<<6 | v12
						v4 = v4 + i32(2)
						goto l52
					l54:
						t116 := int32(m.memory[int64(uint32(v4))+2])
						v12 = v12<<6 | t116&i32(63)
						if uint32(v6) >= uint32(i32(-16)) {
							goto l55
						}
						v6 = v12 | v15<<12
						v4 = v4 + i32(3)
						goto l52
					l55:
						t117 := int32(m.memory[int64(uint32(v4))+3])
						v6 = v12<<6 | t117&i32(63) | v15<<18&i32(0x1c0000)
						v4 = v4 + i32(4)
					}
				l52:
					if v6 == i32(37) {
						if v4 != v16 {
							{
								{
									t118 := int32(int8(m.memory[uint32(v4)]))
									v12 = t118
									if v12 <= i32(-1) {
										goto l59
									}
									v4 = v4 + i32(1)
									v22 = v12 & i32(255)
									goto l60
								}
							l59:
								t119 := int32(m.memory[int64(uint32(v4))+1])
								v15 = t119 & i32(63)
								v22 = v12 & i32(31)
								if uint32(v12) > uint32(i32(-33)) {
									goto l61
								}
								v22 = v22<<6 | v15
								v4 = v4 + i32(2)
								goto l60
							l61:
								t120 := int32(m.memory[int64(uint32(v4))+2])
								v15 = v15<<6 | t120&i32(63)
								if uint32(v12) >= uint32(i32(-16)) {
									goto l62
								}
								v22 = v15 | v22<<12
								v4 = v4 + i32(3)
								goto l60
							l62:
								t121 := int32(m.memory[int64(uint32(v4))+3])
								v22 = v15<<6 | t121&i32(63) | v22<<18&i32(0x1c0000)
								v4 = v4 + i32(4)
							}
						l60:
							v12 = v22 + i32(-49)
							if uint32(v12) > uint32(i32(8)) {
								goto l57
							}
							{
								t122 := int32(load32(m.memory[int64(uint32(v3))+40:]))
								if v7 != t122 {
									goto l63
								}
								m.fn311(v3 + i32(40))
								t123 := int32(load32(m.memory[int64(uint32(v3))+44:]))
								v13 = t123
							}
						l63:
							v6 = v13 + v18
							m.memory[int64(uint32(v6))+4] = byte(v12)
							store32(m.memory[uint32(v6):], uint32(i32(-1)))
							t124 := v3
							v7 = v7 + i32(1)
							store32(m.memory[int64(uint32(t124))+48:], uint32(v7))
							v6 = i32(-2)
							goto l64
						}
						v22 = i32(-1)
						v4 = v16
						goto l57
					}
					v22 = i32(-2)
					goto l57
				l51:
					t125 := int32(load32(m.memory[int64(uint32(v3))+44:]))
					v6 = t125
					t126 := int32(load32(m.memory[int64(uint32(v3))+40:]))
					v13 = t126
					goto l14
				}
			l57:
				{
					{
						{
							if v7 == 0 {
								goto l65
							}
							t127 := int32(load32(m.memory[uint32(v21):]))
							v15 = t127
							if v15 == i32(-1) {
								goto l65
							}
							var p128 int32
							if uint32(v6) < uint32(i32(128)) {
								p128 = 1
							}
							v23 = p128
							if v23 == 0 {
								goto l66
							}
							v13 = i32(1)
							goto l67
						}
					l65:
						store32(m.memory[int64(uint32(v3))+60:], uint32(i32(0)))
						if uint32(v6) < uint32(i32(128)) {
							m.memory[int64(uint32(v3))+60] = byte(v6)
							v6 = i32(1)
							goto l70
						}
						v13 = v6&i32(63) | i32(-128)
						v12 = int32(uint32(v6) >> 6)
						if uint32(v6) >= uint32(i32(2048)) {
							v15 = int32(uint32(v6) >> 12)
							v12 = v12&i32(63) | i32(-128)
							if uint32(v6) > uint32(i32(0xffff)) {
								m.memory[int64(uint32(v3))+63] = byte(v13)
								m.memory[int64(uint32(v3))+62] = byte(v12)
								m.memory[int64(uint32(v3))+61] = byte(v15&i32(63) | i32(-128))
								m.memory[int64(uint32(v3))+60] = byte(int32(uint32(v6)>>18) | i32(-16))
								v6 = i32(4)
								goto l70
							}
							m.memory[int64(uint32(v3))+62] = byte(v13)
							m.memory[int64(uint32(v3))+61] = byte(v12)
							m.memory[int64(uint32(v3))+60] = byte(v15 | i32(224))
							v6 = i32(3)
							goto l70
						}
						m.memory[int64(uint32(v3))+61] = byte(v13)
						m.memory[int64(uint32(v3))+60] = byte(v12 | i32(192))
						v6 = i32(2)
						goto l70
					l66:
						if uint32(v6) >= uint32(i32(2048)) {
							goto l72
						}
						v13 = i32(2)
						goto l67
					l72:
						p129 := i32(4)
						if uint32(v6) < uint32(i32(65536)) {
							p129 = i32(3)
						}
						v13 = p129
					}
				l67:
					{
						t130 := int32(load32(m.memory[uint32(v20):]))
						t131 := v13
						t132 := v15
						v12 = t130
						if uint32(t131) <= uint32(t132-v12) {
							goto l73
						}
						m.fn197(v21, v12, v13, i32(1), i32(1))
					}
				l73:
					t133 := int32(load32(m.memory[uint32(v19):]))
					v15 = t133 + v12
					if v23 != 0 {
						goto l74
					}
					v23 = v6&i32(63) | i32(-128)
					v24 = int32(uint32(v6) >> 6)
					if uint32(v6) >= uint32(i32(2048)) {
						v25 = int32(uint32(v6) >> 12)
						v24 = v24&i32(63) | i32(-128)
						if uint32(v6) > uint32(i32(0xffff)) {
							m.memory[int64(uint32(v15))+3] = byte(v23)
							m.memory[int64(uint32(v15))+2] = byte(v24)
							m.memory[int64(uint32(v15))+1] = byte(v25&i32(63) | i32(-128))
							m.memory[uint32(v15)] = byte(int32(uint32(v6)>>18) | i32(-16))
							goto l76
						}
						m.memory[int64(uint32(v15))+2] = byte(v23)
						m.memory[int64(uint32(v15))+1] = byte(v24)
						m.memory[uint32(v15)] = byte(v25 | i32(224))
						goto l76
					}
					m.memory[int64(uint32(v15))+1] = byte(v23)
					m.memory[uint32(v15)] = byte(v24 | i32(192))
					goto l76
				}
			l74:
				m.memory[uint32(v15)] = byte(v6)
			l76:
				store32(m.memory[uint32(v20):], uint32(v13+v12))
				v13 = v17
				v6 = v22
				goto l78
			l70:
				{
					t134 := m.fn11(v6)
					v15 = t134
					if v15 != 0 {
						if v6 == 0 {
							goto l80
						}
						memory_copy(m.memory, uint32(v15), uint32(v3+i32(60)), uint32(v6))
					l80:
						{
							t135 := int32(load32(m.memory[int64(uint32(v3))+40:]))
							if v7 != t135 {
								goto l81
							}
							m.fn311(v3 + i32(40))
						}
					l81:
						t136 := int32(load32(m.memory[int64(uint32(v3))+44:]))
						v13 = t136
						v12 = v13 + v18
						store32(m.memory[int64(uint32(v12))+8:], uint32(v6))
						store32(m.memory[int64(uint32(v12))+4:], uint32(v15))
						store32(m.memory[uint32(v12):], uint32(v6))
						t137 := v3
						v7 = v7 + i32(1)
						store32(m.memory[int64(uint32(t137))+48:], uint32(v7))
						v6 = v22
						goto l64
					}
					m.fn16(i32(1), v6)
					panic("unreachable")
				}
			}
		}
	}
l45:
	v7 = v7 + i32(44)
	v4 = v4 + i32(-44)
	if v4 == 0 {
		goto l44
	}
	goto l49
l44:
	v6 = i32(4)
	v7 = i32(0)
l14:
	t138 := m.fn405(v1, v2, i32(1079676), i32(5))
	v4 = t138
	store64(m.memory[int64(uint32(v0))+24:], uint64(v14))
	m.memory[int64(uint32(v0))+32] = byte(v11)
	store32(m.memory[int64(uint32(v0))+16:], uint32(v7))
	store32(m.memory[int64(uint32(v0))+12:], uint32(v6))
	store32(m.memory[int64(uint32(v0))+8:], uint32(v13))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
	store32(m.memory[uint32(v0):], uint32(v8))
	m.memory[int64(uint32(v0))+20] = byte(v4 & i32(253))
	m.g0 = v3 + i32(64)
}
func (m *Module) fn407(v0 int32) {
	var v1, v2, v3, v4, v5, v6, v7 int32
	t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
	v1 = t0
	{
		{
			t1 := int32(load32(m.memory[int64(uint32(v0))+16:]))
			v2 = t1
			if v2 == 0 {
				goto l0
			}
			v3 = v1
		l5:
			{
				t2 := int32(load32(m.memory[uint32(v3):]))
				v4 = t2
				if v4 < i32(1) {
					goto l1
				}
				t3 := int32(load32(m.memory[uint32(v3+i32(4)):]))
				v5 = t3
				t4 := int32(load32(m.memory[uint32(v5+i32(-4)):]))
				v6 = t4
				v7 = v6 & i32(-8)
				t5 := v7
				v6 = v6 & i32(3)
				p6 := i32(8)
				if v6 != 0 {
					p6 = i32(4)
				}
				if uint32(t5) < uint32(p6+v4) {
					m.fn7(i32(1274404), i32(46), i32(1274452))
					panic("unreachable")
				}
				if v6 == 0 {
					goto l3
				}
				if uint32(v7) > uint32(v4+i32(39)) {
					m.fn7(i32(1274468), i32(46), i32(1274516))
					panic("unreachable")
				}
			l3:
				m.fn5(v5)
			}
		l1:
			v3 = v3 + i32(12)
			v2 = v2 + i32(-1)
			if v2 != 0 {
				goto l5
			}
		}
	l0:
		{
			t7 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			v3 = t7
			if v3 == 0 {
				goto l6
			}
			t8 := int32(load32(m.memory[uint32(v1+i32(-4)):]))
			v2 = t8
			v4 = v2 & i32(-8)
			t9 := v4
			v2 = v2 & i32(3)
			p10 := i32(8)
			if v2 != 0 {
				p10 = i32(4)
			}
			v3 = v3 * i32(12)
			if uint32(t9) < uint32(p10+v3) {
				m.fn7(i32(1274404), i32(46), i32(1274452))
				panic("unreachable")
			}
			if v2 == 0 {
				goto l8
			}
			if uint32(v4) > uint32(v3+i32(39)) {
				m.fn7(i32(1274468), i32(46), i32(1274516))
				panic("unreachable")
			}
		l8:
			m.fn5(v1)
		}
	l6:
		t11 := int32(load32(m.memory[int64(uint32(v0))+52:]))
		v1 = t11
		{
			t12 := int32(load32(m.memory[int64(uint32(v0))+56:]))
			v2 = t12
			if v2 == 0 {
				goto l10
			}
			v3 = v1
		l15:
			{
				t13 := int32(load32(m.memory[uint32(v3):]))
				v4 = t13
				if v4 < i32(1) {
					goto l11
				}
				t14 := int32(load32(m.memory[uint32(v3+i32(4)):]))
				v5 = t14
				t15 := int32(load32(m.memory[uint32(v5+i32(-4)):]))
				v6 = t15
				v7 = v6 & i32(-8)
				t16 := v7
				v6 = v6 & i32(3)
				p17 := i32(8)
				if v6 != 0 {
					p17 = i32(4)
				}
				if uint32(t16) < uint32(p17+v4) {
					m.fn7(i32(1274404), i32(46), i32(1274452))
					panic("unreachable")
				}
				if v6 == 0 {
					goto l13
				}
				if uint32(v7) > uint32(v4+i32(39)) {
					m.fn7(i32(1274468), i32(46), i32(1274516))
					panic("unreachable")
				}
			l13:
				m.fn5(v5)
			}
		l11:
			v3 = v3 + i32(12)
			v2 = v2 + i32(-1)
			if v2 != 0 {
				goto l15
			}
		}
	l10:
		{
			t18 := int32(load32(m.memory[int64(uint32(v0))+48:]))
			v3 = t18
			if v3 == 0 {
				goto l16
			}
			t19 := int32(load32(m.memory[uint32(v1+i32(-4)):]))
			v2 = t19
			v4 = v2 & i32(-8)
			t20 := v4
			v2 = v2 & i32(3)
			p21 := i32(8)
			if v2 != 0 {
				p21 = i32(4)
			}
			v3 = v3 * i32(12)
			if uint32(t20) < uint32(p21+v3) {
				m.fn7(i32(1274404), i32(46), i32(1274452))
				panic("unreachable")
			}
			if v2 == 0 {
				goto l18
			}
			if uint32(v4) > uint32(v3+i32(39)) {
				m.fn7(i32(1274468), i32(46), i32(1274516))
				panic("unreachable")
			}
		l18:
			m.fn5(v1)
		}
	l16:
		t22 := int32(load32(m.memory[int64(uint32(v0))+92:]))
		v1 = t22
		{
			t23 := int32(load32(m.memory[int64(uint32(v0))+96:]))
			v2 = t23
			if v2 == 0 {
				goto l20
			}
			v3 = v1
		l25:
			{
				t24 := int32(load32(m.memory[uint32(v3):]))
				v4 = t24
				if v4 < i32(1) {
					goto l21
				}
				t25 := int32(load32(m.memory[uint32(v3+i32(4)):]))
				v5 = t25
				t26 := int32(load32(m.memory[uint32(v5+i32(-4)):]))
				v6 = t26
				v7 = v6 & i32(-8)
				t27 := v7
				v6 = v6 & i32(3)
				p28 := i32(8)
				if v6 != 0 {
					p28 = i32(4)
				}
				if uint32(t27) < uint32(p28+v4) {
					m.fn7(i32(1274404), i32(46), i32(1274452))
					panic("unreachable")
				}
				if v6 == 0 {
					goto l23
				}
				if uint32(v7) > uint32(v4+i32(39)) {
					m.fn7(i32(1274468), i32(46), i32(1274516))
					panic("unreachable")
				}
			l23:
				m.fn5(v5)
			}
		l21:
			v3 = v3 + i32(12)
			v2 = v2 + i32(-1)
			if v2 != 0 {
				goto l25
			}
		}
	l20:
		{
			t29 := int32(load32(m.memory[int64(uint32(v0))+88:]))
			v3 = t29
			if v3 == 0 {
				goto l26
			}
			t30 := int32(load32(m.memory[uint32(v1+i32(-4)):]))
			v2 = t30
			v4 = v2 & i32(-8)
			t31 := v4
			v2 = v2 & i32(3)
			p32 := i32(8)
			if v2 != 0 {
				p32 = i32(4)
			}
			v3 = v3 * i32(12)
			if uint32(t31) < uint32(p32+v3) {
				m.fn7(i32(1274404), i32(46), i32(1274452))
				panic("unreachable")
			}
			if v2 == 0 {
				goto l28
			}
			if uint32(v4) > uint32(v3+i32(39)) {
				m.fn7(i32(1274468), i32(46), i32(1274516))
				panic("unreachable")
			}
		l28:
			m.fn5(v1)
		}
	l26:
		t33 := int32(load32(m.memory[int64(uint32(v0))+132:]))
		v1 = t33
		{
			t34 := int32(load32(m.memory[int64(uint32(v0))+136:]))
			v2 = t34
			if v2 == 0 {
				goto l30
			}
			v3 = v1
		l35:
			{
				t35 := int32(load32(m.memory[uint32(v3):]))
				v4 = t35
				if v4 < i32(1) {
					goto l31
				}
				t36 := int32(load32(m.memory[uint32(v3+i32(4)):]))
				v5 = t36
				t37 := int32(load32(m.memory[uint32(v5+i32(-4)):]))
				v6 = t37
				v7 = v6 & i32(-8)
				t38 := v7
				v6 = v6 & i32(3)
				p39 := i32(8)
				if v6 != 0 {
					p39 = i32(4)
				}
				if uint32(t38) < uint32(p39+v4) {
					m.fn7(i32(1274404), i32(46), i32(1274452))
					panic("unreachable")
				}
				if v6 == 0 {
					goto l33
				}
				if uint32(v7) > uint32(v4+i32(39)) {
					m.fn7(i32(1274468), i32(46), i32(1274516))
					panic("unreachable")
				}
			l33:
				m.fn5(v5)
			}
		l31:
			v3 = v3 + i32(12)
			v2 = v2 + i32(-1)
			if v2 != 0 {
				goto l35
			}
		}
	l30:
		{
			t40 := int32(load32(m.memory[int64(uint32(v0))+128:]))
			v3 = t40
			if v3 == 0 {
				goto l36
			}
			t41 := int32(load32(m.memory[uint32(v1+i32(-4)):]))
			v2 = t41
			v4 = v2 & i32(-8)
			t42 := v4
			v2 = v2 & i32(3)
			p43 := i32(8)
			if v2 != 0 {
				p43 = i32(4)
			}
			v3 = v3 * i32(12)
			if uint32(t42) < uint32(p43+v3) {
				m.fn7(i32(1274404), i32(46), i32(1274452))
				panic("unreachable")
			}
			if v2 == 0 {
				goto l38
			}
			if uint32(v4) > uint32(v3+i32(39)) {
				m.fn7(i32(1274468), i32(46), i32(1274516))
				panic("unreachable")
			}
		l38:
			m.fn5(v1)
		}
	l36:
		t44 := int32(load32(m.memory[int64(uint32(v0))+172:]))
		v1 = t44
		{
			t45 := int32(load32(m.memory[int64(uint32(v0))+176:]))
			v2 = t45
			if v2 == 0 {
				goto l40
			}
			v3 = v1
		l45:
			{
				t46 := int32(load32(m.memory[uint32(v3):]))
				v4 = t46
				if v4 < i32(1) {
					goto l41
				}
				t47 := int32(load32(m.memory[uint32(v3+i32(4)):]))
				v5 = t47
				t48 := int32(load32(m.memory[uint32(v5+i32(-4)):]))
				v6 = t48
				v7 = v6 & i32(-8)
				t49 := v7
				v6 = v6 & i32(3)
				p50 := i32(8)
				if v6 != 0 {
					p50 = i32(4)
				}
				if uint32(t49) < uint32(p50+v4) {
					m.fn7(i32(1274404), i32(46), i32(1274452))
					panic("unreachable")
				}
				if v6 == 0 {
					goto l43
				}
				if uint32(v7) > uint32(v4+i32(39)) {
					m.fn7(i32(1274468), i32(46), i32(1274516))
					panic("unreachable")
				}
			l43:
				m.fn5(v5)
			}
		l41:
			v3 = v3 + i32(12)
			v2 = v2 + i32(-1)
			if v2 != 0 {
				goto l45
			}
		}
	l40:
		{
			t51 := int32(load32(m.memory[int64(uint32(v0))+168:]))
			v3 = t51
			if v3 == 0 {
				goto l46
			}
			t52 := int32(load32(m.memory[uint32(v1+i32(-4)):]))
			v2 = t52
			v4 = v2 & i32(-8)
			t53 := v4
			v2 = v2 & i32(3)
			p54 := i32(8)
			if v2 != 0 {
				p54 = i32(4)
			}
			v3 = v3 * i32(12)
			if uint32(t53) < uint32(p54+v3) {
				m.fn7(i32(1274404), i32(46), i32(1274452))
				panic("unreachable")
			}
			if v2 == 0 {
				goto l48
			}
			if uint32(v4) > uint32(v3+i32(39)) {
				m.fn7(i32(1274468), i32(46), i32(1274516))
				panic("unreachable")
			}
		l48:
			m.fn5(v1)
		}
	l46:
		t55 := int32(load32(m.memory[int64(uint32(v0))+212:]))
		v1 = t55
		{
			t56 := int32(load32(m.memory[int64(uint32(v0))+216:]))
			v2 = t56
			if v2 == 0 {
				goto l50
			}
			v3 = v1
		l55:
			{
				t57 := int32(load32(m.memory[uint32(v3):]))
				v4 = t57
				if v4 < i32(1) {
					goto l51
				}
				t58 := int32(load32(m.memory[uint32(v3+i32(4)):]))
				v5 = t58
				t59 := int32(load32(m.memory[uint32(v5+i32(-4)):]))
				v6 = t59
				v7 = v6 & i32(-8)
				t60 := v7
				v6 = v6 & i32(3)
				p61 := i32(8)
				if v6 != 0 {
					p61 = i32(4)
				}
				if uint32(t60) < uint32(p61+v4) {
					m.fn7(i32(1274404), i32(46), i32(1274452))
					panic("unreachable")
				}
				if v6 == 0 {
					goto l53
				}
				if uint32(v7) > uint32(v4+i32(39)) {
					m.fn7(i32(1274468), i32(46), i32(1274516))
					panic("unreachable")
				}
			l53:
				m.fn5(v5)
			}
		l51:
			v3 = v3 + i32(12)
			v2 = v2 + i32(-1)
			if v2 != 0 {
				goto l55
			}
		}
	l50:
		{
			t62 := int32(load32(m.memory[int64(uint32(v0))+208:]))
			v3 = t62
			if v3 == 0 {
				goto l56
			}
			t63 := int32(load32(m.memory[uint32(v1+i32(-4)):]))
			v2 = t63
			v4 = v2 & i32(-8)
			t64 := v4
			v2 = v2 & i32(3)
			p65 := i32(8)
			if v2 != 0 {
				p65 = i32(4)
			}
			v3 = v3 * i32(12)
			if uint32(t64) < uint32(p65+v3) {
				m.fn7(i32(1274404), i32(46), i32(1274452))
				panic("unreachable")
			}
			if v2 == 0 {
				goto l58
			}
			if uint32(v4) > uint32(v3+i32(39)) {
				m.fn7(i32(1274468), i32(46), i32(1274516))
				panic("unreachable")
			}
		l58:
			m.fn5(v1)
		}
	l56:
		t66 := int32(load32(m.memory[int64(uint32(v0))+252:]))
		v1 = t66
		{
			t67 := int32(load32(m.memory[int64(uint32(v0))+256:]))
			v2 = t67
			if v2 == 0 {
				goto l60
			}
			v3 = v1
		l65:
			{
				t68 := int32(load32(m.memory[uint32(v3):]))
				v4 = t68
				if v4 < i32(1) {
					goto l61
				}
				t69 := int32(load32(m.memory[uint32(v3+i32(4)):]))
				v5 = t69
				t70 := int32(load32(m.memory[uint32(v5+i32(-4)):]))
				v6 = t70
				v7 = v6 & i32(-8)
				t71 := v7
				v6 = v6 & i32(3)
				p72 := i32(8)
				if v6 != 0 {
					p72 = i32(4)
				}
				if uint32(t71) < uint32(p72+v4) {
					m.fn7(i32(1274404), i32(46), i32(1274452))
					panic("unreachable")
				}
				if v6 == 0 {
					goto l63
				}
				if uint32(v7) > uint32(v4+i32(39)) {
					m.fn7(i32(1274468), i32(46), i32(1274516))
					panic("unreachable")
				}
			l63:
				m.fn5(v5)
			}
		l61:
			v3 = v3 + i32(12)
			v2 = v2 + i32(-1)
			if v2 != 0 {
				goto l65
			}
		}
	l60:
		{
			t73 := int32(load32(m.memory[int64(uint32(v0))+248:]))
			v3 = t73
			if v3 == 0 {
				goto l66
			}
			t74 := int32(load32(m.memory[uint32(v1+i32(-4)):]))
			v2 = t74
			v4 = v2 & i32(-8)
			t75 := v4
			v2 = v2 & i32(3)
			p76 := i32(8)
			if v2 != 0 {
				p76 = i32(4)
			}
			v3 = v3 * i32(12)
			if uint32(t75) < uint32(p76+v3) {
				m.fn7(i32(1274404), i32(46), i32(1274452))
				panic("unreachable")
			}
			if v2 == 0 {
				goto l68
			}
			if uint32(v4) > uint32(v3+i32(39)) {
				m.fn7(i32(1274468), i32(46), i32(1274516))
				panic("unreachable")
			}
		l68:
			m.fn5(v1)
		}
	l66:
		t77 := int32(load32(m.memory[int64(uint32(v0))+292:]))
		v1 = t77
		{
			t78 := int32(load32(m.memory[int64(uint32(v0))+296:]))
			v2 = t78
			if v2 == 0 {
				goto l70
			}
			v3 = v1
		l75:
			{
				t79 := int32(load32(m.memory[uint32(v3):]))
				v4 = t79
				if v4 < i32(1) {
					goto l71
				}
				t80 := int32(load32(m.memory[uint32(v3+i32(4)):]))
				v5 = t80
				t81 := int32(load32(m.memory[uint32(v5+i32(-4)):]))
				v6 = t81
				v7 = v6 & i32(-8)
				t82 := v7
				v6 = v6 & i32(3)
				p83 := i32(8)
				if v6 != 0 {
					p83 = i32(4)
				}
				if uint32(t82) < uint32(p83+v4) {
					m.fn7(i32(1274404), i32(46), i32(1274452))
					panic("unreachable")
				}
				if v6 == 0 {
					goto l73
				}
				if uint32(v7) > uint32(v4+i32(39)) {
					m.fn7(i32(1274468), i32(46), i32(1274516))
					panic("unreachable")
				}
			l73:
				m.fn5(v5)
			}
		l71:
			v3 = v3 + i32(12)
			v2 = v2 + i32(-1)
			if v2 != 0 {
				goto l75
			}
		}
	l70:
		{
			t84 := int32(load32(m.memory[int64(uint32(v0))+288:]))
			v3 = t84
			if v3 == 0 {
				goto l76
			}
			t85 := int32(load32(m.memory[uint32(v1+i32(-4)):]))
			v2 = t85
			v4 = v2 & i32(-8)
			t86 := v4
			v2 = v2 & i32(3)
			p87 := i32(8)
			if v2 != 0 {
				p87 = i32(4)
			}
			v3 = v3 * i32(12)
			if uint32(t86) < uint32(p87+v3) {
				m.fn7(i32(1274404), i32(46), i32(1274452))
				panic("unreachable")
			}
			if v2 == 0 {
				goto l78
			}
			if uint32(v4) > uint32(v3+i32(39)) {
				m.fn7(i32(1274468), i32(46), i32(1274516))
				panic("unreachable")
			}
		l78:
			m.fn5(v1)
		}
	l76:
		t88 := int32(load32(m.memory[int64(uint32(v0))+332:]))
		v1 = t88
		{
			t89 := int32(load32(m.memory[int64(uint32(v0))+336:]))
			v2 = t89
			if v2 == 0 {
				goto l80
			}
			v3 = v1
		l85:
			{
				t90 := int32(load32(m.memory[uint32(v3):]))
				v4 = t90
				if v4 < i32(1) {
					goto l81
				}
				t91 := int32(load32(m.memory[uint32(v3+i32(4)):]))
				v5 = t91
				t92 := int32(load32(m.memory[uint32(v5+i32(-4)):]))
				v6 = t92
				v7 = v6 & i32(-8)
				t93 := v7
				v6 = v6 & i32(3)
				p94 := i32(8)
				if v6 != 0 {
					p94 = i32(4)
				}
				if uint32(t93) < uint32(p94+v4) {
					m.fn7(i32(1274404), i32(46), i32(1274452))
					panic("unreachable")
				}
				if v6 == 0 {
					goto l83
				}
				if uint32(v7) > uint32(v4+i32(39)) {
					m.fn7(i32(1274468), i32(46), i32(1274516))
					panic("unreachable")
				}
			l83:
				m.fn5(v5)
			}
		l81:
			v3 = v3 + i32(12)
			v2 = v2 + i32(-1)
			if v2 != 0 {
				goto l85
			}
		}
	l80:
		{
			t95 := int32(load32(m.memory[int64(uint32(v0))+328:]))
			v3 = t95
			if v3 == 0 {
				return
			}
			t96 := int32(load32(m.memory[uint32(v1+i32(-4)):]))
			v2 = t96
			v4 = v2 & i32(-8)
			t97 := v4
			v2 = v2 & i32(3)
			p98 := i32(8)
			if v2 != 0 {
				p98 = i32(4)
			}
			v3 = v3 * i32(12)
			if uint32(t97) < uint32(p98+v3) {
				m.fn7(i32(1274404), i32(46), i32(1274452))
				panic("unreachable")
			}
			if v2 == 0 {
				goto l88
			}
			if uint32(v4) > uint32(v3+i32(39)) {
				m.fn7(i32(1274468), i32(46), i32(1274516))
				panic("unreachable")
			}
		l88:
			m.fn5(v1)
		}
		return
	}
}
func (m *Module) fn408(v0 int32) {
	var v1, v2, v3, v4 int32
	{
		t0 := int32(load32(m.memory[uint32(v0):]))
		v1 = t0
		if v1 == i32(-1) {
			goto l0
		}
		if v1 == 0 {
			goto l0
		}
		t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
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
			m.fn7(i32(1274404), i32(46), i32(1274452))
			panic("unreachable")
		}
		if v3 == 0 {
			goto l2
		}
		if uint32(v4) > uint32(v1+i32(39)) {
			m.fn7(i32(1274468), i32(46), i32(1274516))
			panic("unreachable")
		}
	l2:
		m.fn5(v2)
	}
l0:
	{
		t5 := int32(load32(m.memory[int64(uint32(v0))+12:]))
		v1 = t5
		if v1 == i32(-1) {
			goto l4
		}
		if v1 == 0 {
			goto l4
		}
		t6 := int32(load32(m.memory[int64(uint32(v0))+16:]))
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
		if uint32(t8) < uint32(p9+v1) {
			m.fn7(i32(1274404), i32(46), i32(1274452))
			panic("unreachable")
		}
		if v3 == 0 {
			goto l6
		}
		if uint32(v4) > uint32(v1+i32(39)) {
			m.fn7(i32(1274468), i32(46), i32(1274516))
			panic("unreachable")
		}
	l6:
		m.fn5(v2)
	}
l4:
	{
		t10 := int32(load32(m.memory[int64(uint32(v0))+24:]))
		v1 = t10
		if v1 == i32(-1) {
			goto l8
		}
		if v1 == 0 {
			goto l8
		}
		t11 := int32(load32(m.memory[int64(uint32(v0))+28:]))
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
		if uint32(t13) < uint32(p14+v1) {
			m.fn7(i32(1274404), i32(46), i32(1274452))
			panic("unreachable")
		}
		if v3 == 0 {
			goto l10
		}
		if uint32(v4) > uint32(v1+i32(39)) {
			m.fn7(i32(1274468), i32(46), i32(1274516))
			panic("unreachable")
		}
	l10:
		m.fn5(v2)
	}
l8:
	{
		t15 := int32(load32(m.memory[int64(uint32(v0))+36:]))
		v1 = t15
		if v1 == i32(-1) {
			goto l12
		}
		if v1 == 0 {
			goto l12
		}
		t16 := int32(load32(m.memory[int64(uint32(v0))+40:]))
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
		if uint32(t18) < uint32(p19+v1) {
			m.fn7(i32(1274404), i32(46), i32(1274452))
			panic("unreachable")
		}
		if v3 == 0 {
			goto l14
		}
		if uint32(v4) > uint32(v1+i32(39)) {
			m.fn7(i32(1274468), i32(46), i32(1274516))
			panic("unreachable")
		}
	l14:
		m.fn5(v2)
	}
l12:
	{
		t20 := int32(load32(m.memory[int64(uint32(v0))+48:]))
		v1 = t20
		if v1 == i32(-1) {
			goto l16
		}
		if v1 == 0 {
			goto l16
		}
		t21 := int32(load32(m.memory[int64(uint32(v0))+52:]))
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
		if uint32(t23) < uint32(p24+v1) {
			m.fn7(i32(1274404), i32(46), i32(1274452))
			panic("unreachable")
		}
		if v3 == 0 {
			goto l18
		}
		if uint32(v4) > uint32(v1+i32(39)) {
			m.fn7(i32(1274468), i32(46), i32(1274516))
			panic("unreachable")
		}
	l18:
		m.fn5(v2)
	}
l16:
	{
		t25 := int32(load32(m.memory[int64(uint32(v0))+60:]))
		v1 = t25
		if v1 == i32(-1) {
			goto l20
		}
		if v1 == 0 {
			goto l20
		}
		t26 := int32(load32(m.memory[int64(uint32(v0))+64:]))
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
		if uint32(t28) < uint32(p29+v1) {
			m.fn7(i32(1274404), i32(46), i32(1274452))
			panic("unreachable")
		}
		if v3 == 0 {
			goto l22
		}
		if uint32(v4) > uint32(v1+i32(39)) {
			m.fn7(i32(1274468), i32(46), i32(1274516))
			panic("unreachable")
		}
	l22:
		m.fn5(v2)
	}
l20:
	{
		t30 := int32(load32(m.memory[int64(uint32(v0))+72:]))
		v1 = t30
		if v1 == i32(-1) {
			goto l24
		}
		if v1 == 0 {
			goto l24
		}
		t31 := int32(load32(m.memory[int64(uint32(v0))+76:]))
		v2 = t31
		t32 := int32(load32(m.memory[uint32(v2+i32(-4)):]))
		v3 = t32
		v4 = v3 & i32(-8)
		t33 := v4
		v3 = v3 & i32(3)
		p34 := i32(8)
		if v3 != 0 {
			p34 = i32(4)
		}
		if uint32(t33) < uint32(p34+v1) {
			m.fn7(i32(1274404), i32(46), i32(1274452))
			panic("unreachable")
		}
		if v3 == 0 {
			goto l26
		}
		if uint32(v4) > uint32(v1+i32(39)) {
			m.fn7(i32(1274468), i32(46), i32(1274516))
			panic("unreachable")
		}
	l26:
		m.fn5(v2)
	}
l24:
	{
		t35 := int32(load32(m.memory[int64(uint32(v0))+84:]))
		v1 = t35
		if v1 == i32(-1) {
			goto l28
		}
		if v1 == 0 {
			goto l28
		}
		t36 := int32(load32(m.memory[int64(uint32(v0))+88:]))
		v2 = t36
		t37 := int32(load32(m.memory[uint32(v2+i32(-4)):]))
		v3 = t37
		v4 = v3 & i32(-8)
		t38 := v4
		v3 = v3 & i32(3)
		p39 := i32(8)
		if v3 != 0 {
			p39 = i32(4)
		}
		if uint32(t38) < uint32(p39+v1) {
			m.fn7(i32(1274404), i32(46), i32(1274452))
			panic("unreachable")
		}
		if v3 == 0 {
			goto l30
		}
		if uint32(v4) > uint32(v1+i32(39)) {
			m.fn7(i32(1274468), i32(46), i32(1274516))
			panic("unreachable")
		}
	l30:
		m.fn5(v2)
	}
l28:
	{
		t40 := int32(load32(m.memory[int64(uint32(v0))+96:]))
		v1 = t40
		if v1 == i32(-1) {
			return
		}
		if v1 == 0 {
			return
		}
		t41 := int32(load32(m.memory[int64(uint32(v0))+100:]))
		v3 = t41
		t42 := int32(load32(m.memory[uint32(v3+i32(-4)):]))
		v0 = t42
		v2 = v0 & i32(-8)
		t43 := v2
		v0 = v0 & i32(3)
		p44 := i32(8)
		if v0 != 0 {
			p44 = i32(4)
		}
		if uint32(t43) < uint32(p44+v1) {
			m.fn7(i32(1274404), i32(46), i32(1274452))
			panic("unreachable")
		}
		if v0 == 0 {
			goto l34
		}
		if uint32(v2) > uint32(v1+i32(39)) {
			m.fn7(i32(1274468), i32(46), i32(1274516))
			panic("unreachable")
		}
	l34:
		m.fn5(v3)
	}
}
func (m *Module) fn409(v0, v1, v2 int32) {
	var v3, v4 int32
	var v5, v6 int64
	t0 := m.g0
	v3 = t0 - i32(32)
	m.g0 = v3
	switch v2 {
	case 0:
		m.memory[int64(uint32(v0))+1] = byte(i32(0))
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
	switch v4&i32(255) + i32(-43) {
	case 2:
		v4 = v2 + i32(-1)
		v1 = v1 + i32(1)
		if uint32(v2) < uint32(i32(17)) {
			v5 = i64(0)
			if v4 != 0 {
			l15:
				{
					t8 := int32(m.memory[uint32(v1)])
					v2 = t8 + i32(-48)
					if uint32(v2) > uint32(i32(9)) {
						goto l4
					}
					v1 = v1 + i32(1)
					v5 = v5*i64(10) - int64(uint32(v2))
					v4 = v4 + i32(-1)
					if v4 == 0 {
						goto l13
					}
					goto l15
				}
			}
			goto l13
		}
		v5 = i64(0)
	l12:
		{
			m.fn1911(v3+i32(16), v5, v5>>63, i64(10), i64(0))
			t3 := int32(m.memory[uint32(v1)])
			v2 = t3
			t4 := int64(load64(m.memory[int64(uint32(v3))+24:]))
			t5 := int64(load64(m.memory[int64(uint32(v3))+16:]))
			v6 = t5
			if t4 != v6>>63 {
				if uint32((v2+i32(-48))&i32(255)) >= uint32(i32(10)) {
					goto l4
				}
				m.memory[int64(uint32(v0))+1] = byte(i32(3))
				goto l3
			}
			v2 = v2 + i32(-48)
			if uint32(v2) >= uint32(i32(10)) {
				goto l4
			}
			v5 = int64(uint32(v2))
			var p6 int32
			if v5 > i64(0) {
				p6 = 1
			}
			v5 = v6 - v5
			var p7 int32
			if v5 < v6 {
				p7 = 1
			}
			if p6^p7 != 0 {
				m.memory[int64(uint32(v0))+1] = byte(i32(3))
				goto l3
			}
			v1 = v1 + i32(1)
			v4 = v4 + i32(-1)
			if v4 != 0 {
				goto l12
			}
			goto l13
		}
	case 0:
		v2 = v2 + i32(-1)
		v1 = v1 + i32(1)
		fallthrough
	default:
		if uint32(v2) < uint32(i32(16)) {
			v5 = i64(0)
			if v2 == 0 {
				goto l13
			}
		l20:
			{
				t14 := int32(m.memory[uint32(v1)])
				v4 = t14 + i32(-48)
				if uint32(v4) > uint32(i32(9)) {
					goto l4
				}
				v1 = v1 + i32(1)
				v5 = v5*i64(10) + int64(uint32(v4))
				v2 = v2 + i32(-1)
				if v2 == 0 {
					goto l13
				}
				goto l20
			}
		}
		v5 = i64(0)
	l19:
		{
			m.fn1911(v3, v5, v5>>63, i64(10), i64(0))
			t9 := int32(m.memory[uint32(v1)])
			v4 = t9
			t10 := int64(load64(m.memory[int64(uint32(v3))+8:]))
			t11 := int64(load64(m.memory[uint32(v3):]))
			v6 = t11
			if t10 != v6>>63 {
				if uint32((v4+i32(-48))&i32(255)) >= uint32(i32(10)) {
					goto l4
				}
				m.memory[int64(uint32(v0))+1] = byte(i32(2))
				goto l3
			}
			v4 = v4 + i32(-48)
			if uint32(v4) >= uint32(i32(10)) {
				goto l4
			}
			v5 = int64(uint32(v4))
			var p12 int32
			if v5 < i64(0) {
				p12 = 1
			}
			v5 = v6 + v5
			var p13 int32
			if v5 < v6 {
				p13 = 1
			}
			if p12^p13 != 0 {
				m.memory[int64(uint32(v0))+1] = byte(i32(2))
				goto l3
			}
			v1 = v1 + i32(1)
			v2 = v2 + i32(-1)
			if v2 == 0 {
				goto l13
			}
			goto l19
		}
	}
l4:
	v1 = i32(1)
	m.memory[int64(uint32(v0))+1] = byte(i32(1))
	goto l21
l13:
	store64(m.memory[int64(uint32(v0))+8:], uint64(v5))
	v1 = i32(0)
	goto l21
l3:
	v1 = i32(1)
l21:
	m.memory[uint32(v0)] = byte(v1)
	m.g0 = v3 + i32(32)
}
func (m *Module) fn410(v0, v1, v2 int32) {
	var v3, v4 int32
	t0 := m.g0
	v3 = t0 - i32(64)
	m.g0 = v3
	store32(m.memory[int64(uint32(v3))+8:], uint32(i32(0)))
	store64(m.memory[uint32(v3):], uint64(i64(0x800000000)))
	store32(m.memory[int64(uint32(v3))+36:], uint32(i32(0)))
	store64(m.memory[int64(uint32(v3))+28:], uint64(i64(0x800000000)))
	store32(m.memory[int64(uint32(v3))+12:], uint32(i32(0)))
	m.fn746(v3+i32(40), v1, v2, v3, v3+i32(12))
	v2 = v3 + i32(28)
	{
		{
			t1 := int32(load32(m.memory[int64(uint32(v3))+40:]))
			if t1 == i32(-1) {
				goto l0
			}
			t2 := int64(load64(m.memory[int64(uint32(v3))+56:]))
			store64(m.memory[int64(uint32(v0))+16:], uint64(t2))
			t3 := int64(load64(m.memory[int64(uint32(v3))+48:]))
			store64(m.memory[int64(uint32(v0))+8:], uint64(t3))
			t4 := int64(load64(m.memory[int64(uint32(v3))+40:]))
			store64(m.memory[uint32(v0):], uint64(t4))
			m.fn438(v2)
			m.fn423(v3 + i32(12))
			t5 := int32(load32(m.memory[int64(uint32(v3))+4:]))
			v1 = t5
			{
				t6 := int32(load32(m.memory[int64(uint32(v3))+8:]))
				v2 = t6
				if v2 == 0 {
					goto l1
				}
				v0 = v1
			l2:
				m.fn330(v0)
				v0 = v0 + i32(32)
				v2 = v2 + i32(-1)
				if v2 != 0 {
					goto l2
				}
			}
		l1:
			t7 := int32(load32(m.memory[uint32(v3):]))
			v0 = t7
			if v0 == 0 {
				goto l3
			}
			t8 := int32(load32(m.memory[uint32(v1+i32(-4)):]))
			v2 = t8
			v4 = v2 & i32(-8)
			t9 := v4
			v2 = v2 & i32(3)
			p10 := i32(8)
			if v2 != 0 {
				p10 = i32(4)
			}
			v0 = v0 << 5
			if uint32(t9) < uint32(p10|v0) {
				m.fn7(i32(1274404), i32(46), i32(1274452))
				panic("unreachable")
			}
			if v2 == 0 {
				goto l5
			}
			if uint32(v4) > uint32(v0+i32(39)) {
				m.fn7(i32(1274468), i32(46), i32(1274516))
				panic("unreachable")
			}
		l5:
			m.fn5(v1)
			goto l3
		}
	l0:
		m.fn422(v3+i32(12), v3)
		m.fn436(v3, v2)
		store32(m.memory[uint32(v0):], uint32(i32(-1)))
		t11 := int32(load32(m.memory[int64(uint32(v3))+8:]))
		store32(m.memory[int64(uint32(v0))+12:], uint32(t11))
		t12 := int64(load64(m.memory[uint32(v3):]))
		store64(m.memory[int64(uint32(v0))+4:], uint64(t12))
		m.fn438(v2)
		m.fn423(v3 + i32(12))
	}
l3:
	m.g0 = v3 + i32(64)
}
func (m *Module) fn411(v0 int32) {
	var v1, v2, v3 int32
	m.fn153(v0)
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+52:]))
		v1 = t0
		if v1 == 0 {
			return
		}
		t1 := int32(load32(m.memory[int64(uint32(v0))+56:]))
		v2 = t1
		t2 := int32(load32(m.memory[uint32(v2+i32(-4)):]))
		v0 = t2
		v3 = v0 & i32(-8)
		t3 := v3
		v0 = v0 & i32(3)
		p4 := i32(8)
		if v0 != 0 {
			p4 = i32(4)
		}
		if uint32(t3) < uint32(p4+v1) {
			m.fn7(i32(1274404), i32(46), i32(1274452))
			panic("unreachable")
		}
		if v0 == 0 {
			goto l2
		}
		if uint32(v3) > uint32(v1+i32(39)) {
			m.fn7(i32(1274468), i32(46), i32(1274516))
			panic("unreachable")
		}
	l2:
		m.fn5(v2)
	}
}
func (m *Module) fn412(v0 int32) {
	var v1, v2, v3, v4, v5 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		v1 = t0
		t1 := int32(load32(m.memory[uint32(v0):]))
		t2 := v1
		v2 = t1
		if t2 == v2 {
			return
		}
		v1 = v1 - v2
		v0 = v2*i32(40) + v0 + i32(12)
	l5:
		{
			t3 := int32(load32(m.memory[uint32(v0+i32(-4)):]))
			v2 = t3
			if v2 == 0 {
				goto l1
			}
			t4 := int32(load32(m.memory[uint32(v0):]))
			v3 = t4
			t5 := int32(load32(m.memory[uint32(v3+i32(-4)):]))
			v4 = t5
			v5 = v4 & i32(-8)
			t6 := v5
			v4 = v4 & i32(3)
			p7 := i32(8)
			if v4 != 0 {
				p7 = i32(4)
			}
			if uint32(t6) < uint32(p7+v2) {
				m.fn7(i32(1274404), i32(46), i32(1274452))
				panic("unreachable")
			}
			if v4 == 0 {
				goto l3
			}
			if uint32(v5) > uint32(v2+i32(39)) {
				m.fn7(i32(1274468), i32(46), i32(1274516))
				panic("unreachable")
			}
		l3:
			m.fn5(v3)
		}
	l1:
		v0 = v0 + i32(40)
		v1 = v1 + i32(-1)
		if v1 != 0 {
			goto l5
		}
	}
}
func (m *Module) fn413(v0 int32) {
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
					m.fn7(i32(1274404), i32(46), i32(1274452))
					panic("unreachable")
				}
				if v7 == 0 {
					goto l3
				}
				if uint32(v8) > uint32(v5+i32(39)) {
					m.fn7(i32(1274468), i32(46), i32(1274516))
					panic("unreachable")
				}
			l3:
				m.fn5(v6)
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
				m.fn330(v5)
				v5 = v5 + i32(32)
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
				v5 = v5 << 5
				if uint32(t11) < uint32(p12|v5) {
					m.fn7(i32(1274404), i32(46), i32(1274452))
					panic("unreachable")
				}
				if v7 == 0 {
					goto l9
				}
				if uint32(v4) > uint32(v5+i32(39)) {
					m.fn7(i32(1274468), i32(46), i32(1274516))
					panic("unreachable")
				}
			l9:
				m.fn5(v6)
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
			m.fn7(i32(1274404), i32(46), i32(1274452))
			panic("unreachable")
		}
		if v7 == 0 {
			goto l14
		}
		if uint32(v4) > uint32(v5+i32(39)) {
			m.fn7(i32(1274468), i32(46), i32(1274516))
			panic("unreachable")
		}
	l14:
		m.fn5(v1)
	}
}
func (m *Module) fn414(v0, v1 int32) {
	var v2, v3 int32
	{
		if v1 == 0 {
			return
		}
		v2 = v1 * i32(104)
		v1 = v2 + v1 + i32(113)
		if v1 == 0 {
			return
		}
		v2 = v0 - v2
		t0 := int32(load32(m.memory[uint32(v2+i32(-108)):]))
		v0 = t0
		v3 = v0 & i32(-8)
		t1 := v3
		v0 = v0 & i32(3)
		p2 := i32(8)
		if v0 != 0 {
			p2 = i32(4)
		}
		if uint32(t1) < uint32(p2+v1) {
			m.fn7(i32(1274404), i32(46), i32(1274452))
			panic("unreachable")
		}
		if v0 == 0 {
			goto l2
		}
		if uint32(v3) > uint32(v1+i32(39)) {
			m.fn7(i32(1274468), i32(46), i32(1274516))
			panic("unreachable")
		}
	l2:
		m.fn5(v2 + i32(-104))
	}
}
func (m *Module) fn415(v0 int32) {
	var v1, v2, v3, v4 int32
	var v5 int64
	var v6 int32
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
		l4:
			if v5 != i64(0) {
				goto l2
			}
		l3:
			{
				v6 = v4
				v4 = v6 + i32(8)
				v3 = v3 + i32(-3840)
				t4 := int64(load64(m.memory[uint32(v6):]))
				v5 = t4 & i64(-0x7f7f7f7f7f7f7f80)
				if v5 == i64(-0x7f7f7f7f7f7f7f80) {
					goto l3
				}
			}
			v5 = v5 ^ i64(-0x7f7f7f7f7f7f7f80)
		l2:
			v6 = v3 + (i32(0)-int32(uint32(int64(bits.TrailingZeros64(uint64(v5))))>>3))*i32(480)
			m.fn407(v6 + i32(-472))
			m.fn408(v6 + i32(-112))
			v5 = (v5 + i64(-1)) & v5
			v2 = v2 + i32(-1)
			if v2 != 0 {
				goto l4
			}
		}
	l1:
		v4 = v1 * i32(480)
		v3 = v4 + v1 + i32(489)
		if v3 == 0 {
			return
		}
		t5 := int32(load32(m.memory[uint32(v0):]))
		v6 = t5 - v4
		t6 := int32(load32(m.memory[uint32(v6+i32(-484)):]))
		v4 = t6
		v2 = v4 & i32(-8)
		t7 := v2
		v4 = v4 & i32(3)
		p8 := i32(8)
		if v4 != 0 {
			p8 = i32(4)
		}
		if uint32(t7) < uint32(p8+v3) {
			m.fn7(i32(1274404), i32(46), i32(1274452))
			panic("unreachable")
		}
		if v4 == 0 {
			goto l6
		}
		if uint32(v2) > uint32(v3+i32(39)) {
			m.fn7(i32(1274468), i32(46), i32(1274516))
			panic("unreachable")
		}
	l6:
		m.fn5(v6 + i32(-480))
	}
}
func (m *Module) fn416(v0, v1 int32) {
	var v2, v3 int32
	{
		if v1 == 0 {
			return
		}
		t0 := v1
		v2 = (v1*i32(20) + i32(27)) & i32(-8)
		v1 = t0 + v2 + i32(9)
		if v1 == 0 {
			return
		}
		v2 = v0 - v2
		t1 := int32(load32(m.memory[uint32(v2+i32(-4)):]))
		v0 = t1
		v3 = v0 & i32(-8)
		t2 := v3
		v0 = v0 & i32(3)
		p3 := i32(8)
		if v0 != 0 {
			p3 = i32(4)
		}
		if uint32(t2) < uint32(p3+v1) {
			m.fn7(i32(1274404), i32(46), i32(1274452))
			panic("unreachable")
		}
		if v0 == 0 {
			goto l2
		}
		if uint32(v3) > uint32(v1+i32(39)) {
			m.fn7(i32(1274468), i32(46), i32(1274516))
			panic("unreachable")
		}
	l2:
		m.fn5(v2)
	}
}
func (m *Module) fn417(v0, v1 int32) int32 {
	var v2, v3, v4 int32
	var v5, v6 int64
	var v7, v8, v9, v10 int32
	t0 := m.g0
	v2 = t0 - i32(32)
	m.g0 = v2
	v3 = i32(33686018)
	{
		if v1 == 0 {
			goto l0
		}
		v1 = v1 * i32(44)
	l5:
		{
			t1 := int32(load32(m.memory[uint32(v0):]))
			if t1 == i32(-1) {
				goto l1
			}
			t2 := int32(load32(m.memory[uint32(v0+i32(8)):]))
			if t2 != i32(15) {
				goto l1
			}
			t3 := int32(load32(m.memory[uint32(v0+i32(4)):]))
			v4 = t3
			t4 := int64(load64(m.memory[uint32(v4):]))
			t5 := int64(load64(m.memory[uint32(v4+i32(7)):]))
			if t4^i64(8030604426084902260)|(t5^i64(8315168235865862255)) != i64(0) {
				goto l1
			}
			t6 := int32(load32(m.memory[uint32(v0+i32(36)):]))
			v4 = t6
			if v4 == 0 {
				goto l1
			}
			t7 := int32(load32(m.memory[uint32(v0+i32(40)):]))
			if t7 != i32(47) {
				goto l1
			}
			v5 = i64(8462947847038399337)
			{
				{
					t8 := int64(load64(m.memory[int64(uint32(v4))+8:]))
					v6 = t8
					v6 = v6<<56 | v6&i64(0xff00)<<40 | (v6&i64(0xff0000)<<24 | v6&i64(0xff000000)<<8) | (int64(uint64(v6)>>8)&i64(0xff000000) | int64(uint64(v6)>>24)&i64(0xff0000) | (int64(uint64(v6)>>40)&i64(0xff00) | int64(uint64(v6)>>56)))
					if v6 != i64(8462947847038399337) {
						goto l2
					}
					v5 = i64(0x733a6e616d65733a)
					t9 := int64(load64(m.memory[uint32(v4+i32(16)):]))
					v6 = t9
					v6 = v6<<56 | v6&i64(0xff00)<<40 | (v6&i64(0xff0000)<<24 | v6&i64(0xff000000)<<8) | (int64(uint64(v6)>>8)&i64(0xff000000) | int64(uint64(v6)>>24)&i64(0xff0000) | (int64(uint64(v6)>>40)&i64(0xff00) | int64(uint64(v6)>>56)))
					if v6 != i64(0x733a6e616d65733a) {
						goto l2
					}
					v5 = i64(8386611181395471972)
					t10 := int64(load64(m.memory[uint32(v4+i32(24)):]))
					v6 = t10
					v6 = v6<<56 | v6&i64(0xff00)<<40 | (v6&i64(0xff0000)<<24 | v6&i64(0xff000000)<<8) | (int64(uint64(v6)>>8)&i64(0xff000000) | int64(uint64(v6)>>24)&i64(0xff0000) | (int64(uint64(v6)>>40)&i64(0xff00) | int64(uint64(v6)>>56)))
					if v6 != i64(8386611181395471972) {
						goto l2
					}
					v5 = i64(8026388073617978426)
					t11 := int64(load64(m.memory[uint32(v4+i32(32)):]))
					v6 = t11
					v6 = v6<<56 | v6&i64(0xff00)<<40 | (v6&i64(0xff0000)<<24 | v6&i64(0xff000000)<<8) | (int64(uint64(v6)>>8)&i64(0xff000000) | int64(uint64(v6)>>24)&i64(0xff0000) | (int64(uint64(v6)>>40)&i64(0xff00) | int64(uint64(v6)>>56)))
					if v6 != i64(8026388073617978426) {
						goto l2
					}
					v5 = i64(8677711278648226676)
					t12 := int64(load64(m.memory[uint32(v4+i32(40)):]))
					v6 = t12
					v6 = v6<<56 | v6&i64(0xff00)<<40 | (v6&i64(0xff0000)<<24 | v6&i64(0xff000000)<<8) | (int64(uint64(v6)>>8)&i64(0xff000000) | int64(uint64(v6)>>24)&i64(0xff0000) | (int64(uint64(v6)>>40)&i64(0xff00) | int64(uint64(v6)>>56)))
					if v6 != i64(8677711278648226676) {
						goto l2
					}
					v5 = i64(0x74796c653a312e30)
					v7 = i32(0)
					t13 := int64(load64(m.memory[uint32(v4+i32(47)):]))
					v6 = t13
					v6 = v6<<56 | v6&i64(0xff00)<<40 | (v6&i64(0xff0000)<<24 | v6&i64(0xff000000)<<8) | (int64(uint64(v6)>>8)&i64(0xff000000) | int64(uint64(v6)>>24)&i64(0xff0000) | (int64(uint64(v6)>>40)&i64(0xff00) | int64(uint64(v6)>>56)))
					if v6 == i64(0x74796c653a312e30) {
						goto l3
					}
				}
			l2:
				p14 := i32(1)
				if uint64(v6) < uint64(v5) {
					p14 = i32(-1)
				}
				v7 = p14
			}
		l3:
			if v7 == 0 {
				goto l4
			}
		}
	l1:
		v0 = v0 + i32(44)
		v1 = v1 + i32(-44)
		if v1 != 0 {
			goto l5
		}
		goto l0
	l4:
		t15 := int32(load32(m.memory[uint32(v0+i32(16)):]))
		t16 := v2 + i32(24)
		v8 = t15
		t17 := int32(load32(m.memory[uint32(v0+i32(20)):]))
		t18 := v8
		v9 = t17
		m.fn155(t16, t18, v9, i32(1079489), i32(59), i32(1076356), i32(11))
		{
			{
				t19 := int32(load32(m.memory[int64(uint32(v2))+24:]))
				v1 = t19
				if v1 != 0 {
					goto l6
				}
				v1 = i32(2)
				goto l7
			}
		l6:
			v3 = i32(0)
			{
				{
					{
						t20 := int32(load32(m.memory[int64(uint32(v2))+28:]))
						v4 = t20
						switch v4 {
						case 0:
							goto l8
						case 4:
							goto l11
						case 1:
							t21 := int32(m.memory[uint32(v1)])
							v7 = t21
							switch v7 + i32(-43) {
							case 0, 2:
								goto l8
							default:
								goto l12
							}
						default:
							t22 := int32(m.memory[uint32(v1)])
							v7 = t22
						}
					}
				l12:
					t23 := v1
					var p24 int32
					if v7&i32(255) == i32(43) {
						p24 = 1
					}
					v0 = p24
					v7 = t23 + v0
					v1 = v4 - v0
					if uint32(v1) < uint32(i32(9)) {
						if v1 != 0 {
							v3 = i32(0)
							{
								t26 := int32(m.memory[uint32(v7)])
								v0 = t26 + i32(-48)
								if uint32(v0) <= uint32(i32(9)) {
									if v1 != i32(1) {
										t27 := int32(m.memory[int64(uint32(v7))+1])
										v4 = t27 + i32(-48)
										if uint32(v4) <= uint32(i32(9)) {
											v0 = v4 + v0*i32(10)
											if v1 != i32(2) {
												t28 := int32(m.memory[int64(uint32(v7))+2])
												v4 = t28 + i32(-48)
												if uint32(v4) <= uint32(i32(9)) {
													v0 = v4 + v0*i32(10)
													if v1 != i32(3) {
														t29 := int32(m.memory[int64(uint32(v7))+3])
														v4 = t29 + i32(-48)
														if uint32(v4) <= uint32(i32(9)) {
															v0 = v4 + v0*i32(10)
															if v1 != i32(4) {
																t30 := int32(m.memory[int64(uint32(v7))+4])
																v4 = t30 + i32(-48)
																if uint32(v4) <= uint32(i32(9)) {
																	v0 = v4 + v0*i32(10)
																	if v1 != i32(5) {
																		t31 := int32(m.memory[int64(uint32(v7))+5])
																		v4 = t31 + i32(-48)
																		if uint32(v4) <= uint32(i32(9)) {
																			v0 = v4 + v0*i32(10)
																			if v1 != i32(6) {
																				t32 := int32(m.memory[int64(uint32(v7))+6])
																				v4 = t32 + i32(-48)
																				if uint32(v4) <= uint32(i32(9)) {
																					v0 = v4 + v0*i32(10)
																					if v1 != i32(7) {
																						t33 := int32(m.memory[int64(uint32(v7))+7])
																						v1 = t33 + i32(-48)
																						if uint32(v1) > uint32(i32(9)) {
																							goto l8
																						}
																						v0 = v1 + v0*i32(10)
																						v3 = i32(1)
																						goto l8
																					}
																					v3 = i32(1)
																					goto l8
																				}
																				goto l8
																			}
																			v3 = i32(1)
																			goto l8
																		}
																		goto l8
																	}
																	v3 = i32(1)
																	goto l8
																}
																goto l8
															}
															v3 = i32(1)
															goto l8
														}
														goto l8
													}
													v3 = i32(1)
													goto l8
												}
												goto l8
											}
											v3 = i32(1)
											goto l8
										}
										goto l8
									}
									v3 = i32(1)
									goto l8
								}
								goto l8
							}
						}
						v3 = i32(1)
						v0 = i32(0)
						goto l8
					}
					v0 = i32(0)
				l17:
					if v1 != 0 {
						v3 = i32(0)
						v6 = int64(uint32(v0)) * i64(10)
						if int32(int64(uint64(v6)>>32)) == 0 {
							t25 := int32(m.memory[uint32(v7)])
							v4 = t25 + i32(-48)
							if uint32(v4) <= uint32(i32(9)) {
								v7 = v7 + i32(1)
								v1 = v1 + i32(-1)
								v0 = v4 + int32(v6)
								if uint32(v0) >= uint32(v4) {
									goto l17
								}
								goto l8
							}
							goto l8
						}
						goto l8
					}
					v3 = i32(1)
					goto l8
				}
			l11:
				{
					t34 := int32(load32(m.memory[uint32(v1):]))
					if t34 != i32(1684828002) {
						goto l33
					}
					v1 = i32(1)
					goto l7
				}
			l33:
				t35 := int32(m.memory[uint32(v1)])
				var p36 int32
				if t35 == i32(43) {
					p36 = 1
				}
				v0 = p36
				p37 := i32(4)
				if v0 != 0 {
					p37 = i32(3)
				}
				v4 = p37
				v1 = v1 + v0
				v0 = i32(0)
			l35:
				{
					t38 := int32(m.memory[uint32(v1)])
					v7 = t38 + i32(-48)
					if uint32(v7) <= uint32(i32(9)) {
						goto l34
					}
					v3 = i32(0)
					goto l8
				}
			l34:
				v3 = i32(1)
				v1 = v1 + i32(1)
				v0 = v7 + v0*i32(10)
				v4 = v4 + i32(-1)
				if v4 != 0 {
					goto l35
				}
			}
		l8:
			t39 := v3
			var p40 int32
			if uint32(v0) > uint32(i32(599)) {
				p40 = 1
			}
			v1 = t39 & p40
		}
	l7:
		m.fn155(v2+i32(16), v8, v9, i32(1079489), i32(59), i32(1076367), i32(10))
		{
			t41 := int32(load32(m.memory[int64(uint32(v2))+16:]))
			v4 = t41
			if v4 != 0 {
				goto l36
			}
			v0 = i32(512)
			goto l37
		}
	l36:
		v0 = i32(0)
		{
			{
				t42 := int32(load32(m.memory[int64(uint32(v2))+20:]))
				switch t42 + i32(-6) {
				default:
					goto l37
				case 0:
					v7 = i32(1769234796)
					t43 := int32(load32(m.memory[uint32(v4):]))
					v3 = t43
					v3 = i32_rotr(v3&i32(0xff00ff), i32(8)) | i32_rotr(v3, i32(24))&i32(0xff00ff)
					if v3 != i32(1769234796) {
						goto l40
					}
					v7 = i32(26979)
					t44 := int32(load16(m.memory[uint32(v4+i32(4)):]))
					v4 = t44
					v3 = (v4<<8 | int32(uint32(v4)>>8)) & i32(0xffff)
					if v3 != i32(26979) {
						goto l40
					}
					if i32(0) == 0 {
						goto l41
					}
					goto l37
				case 1:
					v7 = i32(1868721257)
					{
						{
							t45 := int32(load32(m.memory[uint32(v4):]))
							v3 = t45
							v3 = i32_rotr(v3&i32(0xff00ff), i32(8)) | i32_rotr(v3, i32(24))&i32(0xff00ff)
							if v3 != i32(1868721257) {
								goto l42
							}
							v7 = i32(1769043301)
							v10 = i32(0)
							t46 := int32(load32(m.memory[uint32(v4+i32(3)):]))
							v4 = t46
							v3 = i32_rotr(v4&i32(0xff00ff), i32(8)) | i32_rotr(v4, i32(24))&i32(0xff00ff)
							if v3 == i32(1769043301) {
								goto l43
							}
						}
					l42:
						p47 := i32(1)
						if uint32(v3) < uint32(v7) {
							p47 = i32(-1)
						}
						v10 = p47
					}
				l43:
					if v10 != 0 {
						goto l37
					}
					goto l41
				}
			}
		l40:
			p48 := i32(1)
			if uint32(v3) < uint32(v7) {
				p48 = i32(-1)
			}
			if p48 != 0 {
				goto l37
			}
		}
	l41:
		v0 = i32(256)
	l37:
		m.fn155(v2+i32(8), v8, v9, i32(1073889), i32(47), i32(1079548), i32(23))
		{
			{
				t49 := int32(load32(m.memory[int64(uint32(v2))+8:]))
				v4 = t49
				if v4 != 0 {
					goto l44
				}
				v4 = i32(33685504)
				goto l45
			}
		l44:
			{
				t50 := int32(load32(m.memory[int64(uint32(v2))+12:]))
				if t50 == i32(4) {
					goto l46
				}
				v4 = i32(33619968)
				goto l45
			}
		l46:
			t51 := int32(load32(m.memory[uint32(v4):]))
			p52 := i32(33619968)
			if t51 == i32(1701736302) {
				p52 = i32(0x2000000)
			}
			v4 = p52
		}
	l45:
		v3 = v0 | v1 | v4
	}
l0:
	m.g0 = v2 + i32(32)
	return v3
}
func (m *Module) fn418(v0, v1, v2 int32) {
	var v3, v4 int32
	var v5 int64
	var v6, v7 int32
	var v8, v9 int64
	var v10, v11, v12 int32
	var v13 int64
	var v14, v15 int32
	t0 := int64(load64(m.memory[int64(uint32(v0))+16:]))
	t1 := int64(load64(m.memory[int64(uint32(v0))+24:]))
	t2 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	v3 = t2
	t3 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	t4 := v3
	v4 = t3
	t5 := m.fn65(t0, t1, t4, v4)
	v5 = t5
	{
		t6 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		if t6 != 0 {
			goto l0
		}
		_ = m.fn71(v0, v0+i32(16))
	}
l0:
	t8 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	v6 = t8
	v7 = v6 & int32(v5)
	v8 = int64(uint64(v5) >> 25)
	v9 = v8 & i64(127) * i64(72340172838076673)
	t9 := int32(load32(m.memory[uint32(v0):]))
	v10 = t9
	v11 = i32(0)
	v12 = i32(0)
l14:
	{
		t10 := int64(load64(m.memory[uint32(v10+v7):]))
		v13 = t10
		v5 = v13 ^ v9
		v5 = (v5 ^ i64(-1)) & (v5 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
		if v5 == 0 {
			goto l1
		}
	l4:
		{
			t11 := v4
			v14 = v10 - (int32(uint32(int64(bits.TrailingZeros64(uint64(v5))))>>3)+v7)&v6<<4
			t12 := int32(load32(m.memory[uint32(v14+i32(-8)):]))
			if t11 != t12 {
				goto l2
			}
			t13 := int32(load32(m.memory[uint32(v14+i32(-12)):]))
			t14 := m.fn1909(v3, t13, v4)
			if t14 == 0 {
				store32(m.memory[uint32(v14+i32(-4)):], uint32(v2))
				{
					t24 := int32(load32(m.memory[uint32(v1):]))
					v0 = t24
					if v0 == 0 {
						return
					}
					t25 := int32(load32(m.memory[uint32(v3+i32(-4)):]))
					v1 = t25
					v10 = v1 & i32(-8)
					t26 := v10
					v1 = v1 & i32(3)
					p27 := i32(8)
					if v1 != 0 {
						p27 = i32(4)
					}
					if uint32(t26) < uint32(p27+v0) {
						m.fn7(i32(1274404), i32(46), i32(1274452))
						panic("unreachable")
					}
					if v1 == 0 {
						goto l12
					}
					if uint32(v10) > uint32(v0+i32(39)) {
						m.fn7(i32(1274468), i32(46), i32(1274516))
						panic("unreachable")
					}
				l12:
					m.fn5(v3)
				}
				return
			}
		}
	l2:
		v5 = (v5 + i64(-1)) & v5
		if !(v5 == 0) {
			goto l4
		}
	}
l1:
	v5 = v13 & i64(-0x7f7f7f7f7f7f7f80)
	if v11 == i32(1) {
		goto l5
	}
	if v5 == 0 {
		v11 = i32(0)
		goto l8
	}
	v15 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v5))))>>3) + v7) & v6
l5:
	if v5&(v13<<1) != i64(0) {
		{
			t15 := int32(int8(m.memory[uint32(v10+v15)]))
			v7 = t15
			if v7 < i32(0) {
				goto l9
			}
			t16 := int64(load64(m.memory[uint32(v10):]))
			t17 := v10
			v15 = int32(uint32(int64(bits.TrailingZeros64(uint64(t16&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
			t18 := int32(m.memory[uint32(t17+v15)])
			v7 = t18
		}
	l9:
		t19 := v10 + v15
		v3 = int32(v8) & i32(127)
		m.memory[uint32(t19)] = byte(v3)
		m.memory[uint32(v10+(v15+i32(-8))&v6+i32(8))] = byte(v3)
		t20 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		store32(m.memory[int64(uint32(v0))+8:], uint32(t20-v7&i32(1)))
		t21 := int32(load32(m.memory[int64(uint32(v0))+12:]))
		store32(m.memory[int64(uint32(v0))+12:], uint32(t21+i32(1)))
		v0 = v10 - v15<<4
		store32(m.memory[uint32(v0+i32(-4)):], uint32(v2))
		v0 = v0 + i32(-16)
		t22 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		store32(m.memory[int64(uint32(v0))+8:], uint32(t22))
		t23 := int64(load64(m.memory[uint32(v1):]))
		store64(m.memory[uint32(v0):], uint64(t23))
		return
	}
	v11 = i32(1)
	goto l8
l8:
	v12 = v12 + i32(8)
	v7 = (v12 + v7) & v6
	goto l14
}
func (m *Module) fn419(v0, v1, v2 int32) {
	var v3, v4, v5, v6, v7, v8, v9, v10, v11, v12 int32
	var v13, v14, v15, v16 int64
	t0 := m.g0
	v3 = t0 - i32(480)
	m.g0 = v3
	m.memory[int64(uint32(v3))+476] = byte(i32(0))
	store32(m.memory[int64(uint32(v3))+472:], uint32(i32(1)))
	store64(m.memory[int64(uint32(v3))+456:], uint64(i64(-0x100000000)))
	store64(m.memory[int64(uint32(v3))+448:], uint64(i64(0x100000000)))
	store64(m.memory[int64(uint32(v3))+440:], uint64(i64(1)))
	m.memory[int64(uint32(v3))+436] = byte(i32(0))
	store32(m.memory[int64(uint32(v3))+432:], uint32(i32(1)))
	store64(m.memory[int64(uint32(v3))+416:], uint64(i64(-0x100000000)))
	store64(m.memory[int64(uint32(v3))+408:], uint64(i64(0x100000000)))
	store64(m.memory[int64(uint32(v3))+400:], uint64(i64(1)))
	m.memory[int64(uint32(v3))+396] = byte(i32(0))
	store32(m.memory[int64(uint32(v3))+392:], uint32(i32(1)))
	store64(m.memory[int64(uint32(v3))+376:], uint64(i64(-0x100000000)))
	store64(m.memory[int64(uint32(v3))+368:], uint64(i64(0x100000000)))
	store64(m.memory[int64(uint32(v3))+360:], uint64(i64(1)))
	m.memory[int64(uint32(v3))+356] = byte(i32(0))
	store32(m.memory[int64(uint32(v3))+352:], uint32(i32(1)))
	store64(m.memory[int64(uint32(v3))+336:], uint64(i64(-0x100000000)))
	store64(m.memory[int64(uint32(v3))+328:], uint64(i64(0x100000000)))
	store64(m.memory[int64(uint32(v3))+320:], uint64(i64(1)))
	m.memory[int64(uint32(v3))+316] = byte(i32(0))
	store32(m.memory[int64(uint32(v3))+312:], uint32(i32(1)))
	store64(m.memory[int64(uint32(v3))+296:], uint64(i64(-0x100000000)))
	store64(m.memory[int64(uint32(v3))+288:], uint64(i64(0x100000000)))
	store64(m.memory[int64(uint32(v3))+280:], uint64(i64(1)))
	m.memory[int64(uint32(v3))+276] = byte(i32(0))
	store32(m.memory[int64(uint32(v3))+272:], uint32(i32(1)))
	store64(m.memory[int64(uint32(v3))+256:], uint64(i64(-0x100000000)))
	store64(m.memory[int64(uint32(v3))+248:], uint64(i64(0x100000000)))
	store64(m.memory[int64(uint32(v3))+240:], uint64(i64(1)))
	m.memory[int64(uint32(v3))+236] = byte(i32(0))
	store32(m.memory[int64(uint32(v3))+232:], uint32(i32(1)))
	store64(m.memory[int64(uint32(v3))+216:], uint64(i64(-0x100000000)))
	store64(m.memory[int64(uint32(v3))+208:], uint64(i64(0x100000000)))
	store64(m.memory[int64(uint32(v3))+200:], uint64(i64(1)))
	m.memory[int64(uint32(v3))+196] = byte(i32(0))
	store32(m.memory[int64(uint32(v3))+192:], uint32(i32(1)))
	store64(m.memory[int64(uint32(v3))+176:], uint64(i64(-0x100000000)))
	store64(m.memory[int64(uint32(v3))+168:], uint64(i64(0x100000000)))
	store64(m.memory[int64(uint32(v3))+160:], uint64(i64(1)))
	m.memory[int64(uint32(v3))+156] = byte(i32(0))
	store32(m.memory[int64(uint32(v3))+152:], uint32(i32(1)))
	store64(m.memory[int64(uint32(v3))+136:], uint64(i64(-0x100000000)))
	store64(m.memory[int64(uint32(v3))+128:], uint64(i64(0x100000000)))
	store64(m.memory[int64(uint32(v3))+120:], uint64(i64(1)))
	m.memory[int64(uint32(v3))+116] = byte(i32(0))
	store32(m.memory[int64(uint32(v3))+112:], uint32(i32(1)))
	store64(m.memory[int64(uint32(v3))+96:], uint64(i64(-0x100000000)))
	store64(m.memory[int64(uint32(v3))+88:], uint64(i64(0x100000000)))
	store64(m.memory[int64(uint32(v3))+80:], uint64(i64(1)))
	v4 = v1 + v2*i32(44)
l1:
	{
		{
			{
				v2 = v1
				if v2 == v4 {
					memory_copy(m.memory, uint32(v0), uint32(v3+i32(80)), uint32(i32(400)))
					m.g0 = v3 + i32(480)
					return
				}
				v1 = v2 + i32(44)
				t1 := int32(load32(m.memory[uint32(v2):]))
				if t1 == i32(-1) {
					goto l1
				}
				t2 := v3 + i32(72)
				v5 = v2 + i32(16)
				t3 := int32(load32(m.memory[uint32(v5):]))
				v6 = t3
				t4 := v6
				v7 = v2 + i32(20)
				t5 := int32(load32(m.memory[uint32(v7):]))
				v8 = t5
				m.fn155(t2, t4, v8, i32(1071795), i32(46), i32(1079440), i32(5))
				t6 := int32(load32(m.memory[int64(uint32(v3))+72:]))
				v9 = t6
				if v9 == 0 {
					goto l1
				}
				{
					t7 := int32(load32(m.memory[int64(uint32(v3))+76:]))
					v10 = t7
					switch v10 {
					case 0:
						goto l1
					case 1:
						t8 := int32(m.memory[uint32(v9)])
						v11 = t8
						switch v11 + i32(-43) {
						case 0, 2:
							goto l1
						default:
							goto l4
						}
					default:
						t9 := int32(m.memory[uint32(v9)])
						v11 = t9
					}
				}
			l4:
				t10 := v9
				var p11 int32
				if v11&i32(255) == i32(43) {
					p11 = 1
				}
				v12 = p11
				v11 = t10 + v12
				{
					v9 = v10 - v12
					if uint32(v9) < uint32(i32(9)) {
						goto l5
					}
					v10 = i32(0)
				l7:
					{
						if v9 == 0 {
							goto l6
						}
						v13 = int64(uint32(v10)) * i64(10)
						if int32(int64(uint64(v13)>>32)) != 0 {
							goto l1
						}
						t12 := int32(m.memory[uint32(v11)])
						v12 = t12 + i32(-48)
						if uint32(v12) > uint32(i32(9)) {
							goto l1
						}
						v11 = v11 + i32(1)
						v9 = v9 + i32(-1)
						v10 = v12 + int32(v13)
						if uint32(v10) < uint32(v12) {
							goto l1
						}
						goto l7
					}
				l5:
					if v9 == 0 {
						goto l1
					}
					t13 := int32(m.memory[uint32(v11)])
					v10 = t13 + i32(-48)
					if uint32(v10) > uint32(i32(9)) {
						goto l1
					}
					if v9 == i32(1) {
						goto l6
					}
					t14 := int32(m.memory[int64(uint32(v11))+1])
					v12 = t14 + i32(-48)
					if uint32(v12) > uint32(i32(9)) {
						goto l1
					}
					v10 = v12 + v10*i32(10)
					if v9 == i32(2) {
						goto l6
					}
					t15 := int32(m.memory[int64(uint32(v11))+2])
					v12 = t15 + i32(-48)
					if uint32(v12) > uint32(i32(9)) {
						goto l1
					}
					v10 = v12 + v10*i32(10)
					if v9 == i32(3) {
						goto l6
					}
					t16 := int32(m.memory[int64(uint32(v11))+3])
					v12 = t16 + i32(-48)
					if uint32(v12) > uint32(i32(9)) {
						goto l1
					}
					v10 = v12 + v10*i32(10)
					if v9 == i32(4) {
						goto l6
					}
					t17 := int32(m.memory[int64(uint32(v11))+4])
					v12 = t17 + i32(-48)
					if uint32(v12) > uint32(i32(9)) {
						goto l1
					}
					v10 = v12 + v10*i32(10)
					if v9 == i32(5) {
						goto l6
					}
					t18 := int32(m.memory[int64(uint32(v11))+5])
					v12 = t18 + i32(-48)
					if uint32(v12) > uint32(i32(9)) {
						goto l1
					}
					v10 = v12 + v10*i32(10)
					if v9 == i32(6) {
						goto l6
					}
					t19 := int32(m.memory[int64(uint32(v11))+6])
					v12 = t19 + i32(-48)
					if uint32(v12) > uint32(i32(9)) {
						goto l1
					}
					v10 = v12 + v10*i32(10)
					if v9 == i32(7) {
						goto l6
					}
					t20 := int32(m.memory[int64(uint32(v11))+7])
					v9 = t20 + i32(-48)
					if uint32(v9) > uint32(i32(9)) {
						goto l1
					}
					v10 = v9 + v10*i32(10)
				}
			l6:
				if uint32(v10+i32(-1)) > uint32(i32(9)) {
					goto l1
				}
				{
					{
						t21 := int32(load32(m.memory[int64(uint32(v2))+8:]))
						v9 = t21
						if v9 != i32(23) {
							goto l8
						}
						t22 := int32(load32(m.memory[int64(uint32(v2))+4:]))
						v9 = t22
						t23 := int64(load64(m.memory[uint32(v9):]))
						t24 := int64(load64(m.memory[uint32(v9+i32(8)):]))
						t25 := int64(load64(m.memory[uint32(v9+i32(15)):]))
						if t23^i64(8531344011606321516)|(t24^i64(7308349836370996325))|(t25^i64(0x7265626d756e2d65)) != i64(0) {
							goto l1
						}
						t26 := int32(load32(m.memory[int64(uint32(v2))+36:]))
						v9 = t26
						if v9 == 0 {
							goto l1
						}
						t27 := int32(load32(m.memory[int64(uint32(v2))+40:]))
						if t27 != i32(46) {
							goto l1
						}
						t28 := int64(load64(m.memory[int64(uint32(v9))+8:]))
						t29 := int64(load64(m.memory[uint32(v9+i32(16)):]))
						t30 := int64(load64(m.memory[uint32(v9+i32(24)):]))
						t31 := int64(load64(m.memory[uint32(v9+i32(32)):]))
						t32 := int64(load64(m.memory[uint32(v9+i32(40)):]))
						t33 := int64(load64(m.memory[uint32(v9+i32(46)):]))
						if !(t28^i64(7598524126653739637)|(t29^i64(4211821596982000243))|(t30^i64(7236833184807805812)|(t31^i64(4212112933405418351)))|(t32^i64(7310532362577407352)|(t33^i64(3471766489881142644))) == 0) {
							goto l1
						}
						goto l9
					}
				l8:
					if v9 != i32(19) {
						goto l1
					}
					t34 := int32(load32(m.memory[int64(uint32(v2))+4:]))
					v9 = t34
					t35 := int64(load64(m.memory[uint32(v9):]))
					t36 := int64(load64(m.memory[uint32(v9+i32(8)):]))
					t37 := int64(load64(m.memory[uint32(v9+i32(11)):]))
					if t35^i64(3271142103424726383)|(t36^i64(8391100474303341932))|(t37^i64(7308349836370996325)) != i64(0) {
						goto l1
					}
					t38 := int32(load32(m.memory[int64(uint32(v2))+36:]))
					v9 = t38
					if v9 == 0 {
						goto l1
					}
					t39 := int32(load32(m.memory[int64(uint32(v2))+40:]))
					if t39 != i32(46) {
						goto l1
					}
					t40 := int64(load64(m.memory[int64(uint32(v9))+8:]))
					t41 := int64(load64(m.memory[uint32(v9+i32(16)):]))
					t42 := int64(load64(m.memory[uint32(v9+i32(24)):]))
					t43 := int64(load64(m.memory[uint32(v9+i32(32)):]))
					t44 := int64(load64(m.memory[uint32(v9+i32(40)):]))
					t45 := int64(load64(m.memory[uint32(v9+i32(46)):]))
					if t40^i64(7598524126653739637)|(t41^i64(4211821596982000243))|(t42^i64(7236833184807805812)|(t43^i64(4212112933405418351)))|(t44^i64(7310532362577407352)|(t45^i64(3471766489881142644))) != i64(0) {
						goto l1
					}
				}
			l9:
				v10 = v3 + i32(80) + v10*i32(40)
				m.fn155(v3+i32(64), v6, v8, i32(1073889), i32(47), i32(1079445), i32(10))
				v2 = i32(1)
				{
					t46 := int32(load32(m.memory[int64(uint32(v3))+64:]))
					v9 = t46
					if v9 == 0 {
						goto l10
					}
					{
						t47 := int32(load32(m.memory[int64(uint32(v3))+68:]))
						switch t47 {
						default:
							goto l10
						case 0:
							v2 = i32(0)
							goto l10
						case 1:
							t48 := int32(m.memory[uint32(v9)])
							v9 = t48 + i32(-65)
							v9 = v9<<5 | int32(uint32(v9&i32(248))>>3)
							if uint32(v9&i32(255)) >= uint32(i32(6)) {
								goto l10
							}
							v2 = int32(i64_shr_u(i64(4406653289731), int64(uint32(v9<<3))&i64(248)))
						}
					}
				}
			l10:
				v8 = v10 + i32(-40)
				m.memory[uint32(v10+i32(-4))] = byte(v2)
				t49 := int32(load32(m.memory[uint32(v5):]))
				t50 := int32(load32(m.memory[uint32(v7):]))
				m.fn155(v3+i32(56), t49, t50, i32(1071795), i32(46), i32(0x107788), i32(11))
				v14 = i64(1)
				{
					t51 := int32(load32(m.memory[int64(uint32(v3))+56:]))
					v2 = t51
					if v2 == 0 {
						goto l13
					}
					{
						t52 := int32(load32(m.memory[int64(uint32(v3))+60:]))
						v9 = t52
						switch v9 {
						case 0:
							goto l13
						case 1:
							t53 := int32(m.memory[uint32(v2)])
							v6 = t53
							switch v6 + i32(-43) {
							case 0, 2:
								goto l13
							default:
								goto l16
							}
						default:
							t54 := int32(m.memory[uint32(v2)])
							v6 = t54
						}
					}
				l16:
					t55 := v2
					var p56 int32
					if v6&i32(255) == i32(43) {
						p56 = 1
					}
					v6 = p56
					v2 = t55 + v6
					v9 = v9 - v6
					if uint32(v9) < uint32(i32(17)) {
						goto l17
					}
					v13 = i64(0)
				l19:
					{
						if v9 == 0 {
							goto l18
						}
						m.fn1911(v3+i32(32), v13, i64(0), i64(10), i64(0))
						t57 := int64(load64(m.memory[int64(uint32(v3))+40:]))
						if t57 != i64(0) {
							goto l13
						}
						t58 := int32(m.memory[uint32(v2)])
						v6 = t58 + i32(-48)
						if uint32(v6) > uint32(i32(9)) {
							goto l13
						}
						v2 = v2 + i32(1)
						v9 = v9 + i32(-1)
						t59 := int64(load64(m.memory[int64(uint32(v3))+32:]))
						v15 = t59
						v13 = v15 + int64(uint32(v6))
						if uint64(v13) >= uint64(v15) {
							goto l19
						}
						goto l13
					}
				l17:
					v13 = i64(0)
					if v9 == 0 {
						goto l18
					}
				l20:
					{
						t60 := int32(m.memory[uint32(v2)])
						v6 = t60 + i32(-48)
						if uint32(v6) > uint32(i32(9)) {
							goto l13
						}
						v2 = v2 + i32(1)
						v13 = v13*i64(10) + int64(uint32(v6))
						v9 = v9 + i32(-1)
						if v9 != 0 {
							goto l20
						}
					}
				l18:
					p61 := i64(0xffffffff)
					if uint64(v13) < uint64(i64(0xffffffff)) {
						p61 = v13
					}
					v14 = p61
				}
			l13:
				store64(m.memory[uint32(v8):], uint64(v14))
				t62 := int32(load32(m.memory[uint32(v5):]))
				t63 := int32(load32(m.memory[uint32(v7):]))
				m.fn155(v3+i32(24), t62, t63, i32(1073889), i32(47), i32(1079455), i32(10))
				t64 := int32(load32(m.memory[int64(uint32(v3))+28:]))
				v9 = t64
				t65 := int32(load32(m.memory[int64(uint32(v3))+24:]))
				t66 := v9
				v6 = t65
				p67 := i32(0)
				if v6 != 0 {
					p67 = t66
				}
				v2 = p67
				if v2 <= i32(-1) {
					goto l21
				}
				if v2 != 0 {
					t68 := m.fn11(v2)
					v8 = t68
					if v8 == 0 {
						m.fn16(i32(1), v9)
						panic("unreachable")
					}
					if v2 == 0 {
						goto l23
					}
					t70 := v8
					p69 := i32(1)
					if v6 != 0 {
						p69 = v6
					}
					memory_copy(m.memory, uint32(t70), uint32(p69), uint32(v2))
					goto l23
				}
				v9 = i32(0)
				v8 = i32(1)
				goto l23
			}
		l23:
			{
				v2 = v10 + i32(-32)
				t71 := int32(load32(m.memory[uint32(v2):]))
				v6 = t71
				if v6 == 0 {
					goto l25
				}
				t72 := int32(load32(m.memory[uint32(v10+i32(-28)):]))
				m.fn21(t72, v6, i32(1))
			}
		l25:
			store32(m.memory[uint32(v2):], uint32(v9))
			store32(m.memory[uint32(v10+i32(-24)):], uint32(v9))
			store32(m.memory[uint32(v10+i32(-28)):], uint32(v8))
			t73 := int32(load32(m.memory[uint32(v5):]))
			t74 := int32(load32(m.memory[uint32(v7):]))
			m.fn155(v3+i32(16), t73, t74, i32(1073889), i32(47), i32(1079465), i32(10))
			{
				t75 := int32(load32(m.memory[int64(uint32(v3))+16:]))
				v9 = t75
				if v9 != 0 {
					t76 := int32(load32(m.memory[int64(uint32(v3))+20:]))
					v2 = t76
					if v2 <= i32(-1) {
						goto l21
					}
					if v2 != 0 {
						t77 := m.fn11(v2)
						v6 = t77
						if v6 == 0 {
							m.fn16(i32(1), v2)
							panic("unreachable")
						}
						if v2 == 0 {
							goto l30
						}
						memory_copy(m.memory, uint32(v6), uint32(v9), uint32(v2))
					l30:
						v16 = int64(uint32(v2))<<32 | int64(uint32(v6))
						goto l27
					}
					v16 = i64(1)
					v2 = i32(0)
					goto l27
				}
				v2 = i32(-1)
				goto l27
			}
		}
	l21:
		m.fn15()
		panic("unreachable")
	l27:
		{
			v6 = v10 + i32(-20)
			t78 := int32(load32(m.memory[uint32(v6):]))
			v9 = t78
			if v9 == i32(-1) {
				goto l31
			}
			if v9 == 0 {
				goto l31
			}
			t79 := int32(load32(m.memory[uint32(v10+i32(-16)):]))
			m.fn21(t79, v9, i32(1))
		}
	l31:
		store32(m.memory[uint32(v6):], uint32(v2))
		store64(m.memory[uint32(v10+i32(-16)):], uint64(v16))
		t80 := int32(load32(m.memory[uint32(v5):]))
		t81 := int32(load32(m.memory[uint32(v7):]))
		m.fn155(v3+i32(8), t80, t81, i32(1071795), i32(46), i32(1079475), i32(14))
		{
			{
				t82 := int32(load32(m.memory[int64(uint32(v3))+8:]))
				v9 = t82
				if v9 != 0 {
					goto l32
				}
				v2 = i32(1)
				goto l33
			}
		l32:
			v2 = i32(1)
			{
				t83 := int32(load32(m.memory[int64(uint32(v3))+12:]))
				v5 = t83
				switch v5 {
				case 0:
					goto l33
				case 1:
					v2 = i32(1)
					t84 := int32(m.memory[uint32(v9)])
					v7 = t84
					switch v7 + i32(-43) {
					case 0, 2:
						goto l33
					default:
						goto l36
					}
				default:
					t85 := int32(m.memory[uint32(v9)])
					v7 = t85
				}
			}
		l36:
			t86 := v9
			var p87 int32
			if v7&i32(255) == i32(43) {
				p87 = 1
			}
			v2 = p87
			v7 = t86 + v2
			v9 = v5 - v2
			if uint32(v9) < uint32(i32(9)) {
				goto l37
			}
			v6 = i32(0)
		l40:
			if v9 != 0 {
				v13 = int64(uint32(v6)) * i64(10)
				if int32(int64(uint64(v13)>>32)) == 0 {
					v2 = i32(1)
					t88 := int32(m.memory[uint32(v7)])
					v5 = t88 + i32(-48)
					if uint32(v5) > uint32(i32(9)) {
						goto l33
					}
					v7 = v7 + i32(1)
					v9 = v9 + i32(-1)
					v6 = v5 + int32(v13)
					if uint32(v6) >= uint32(v5) {
						goto l40
					}
					goto l33
				}
				v2 = i32(1)
				goto l33
			}
			v2 = v6
			goto l33
		l37:
			if v9 != 0 {
				goto l41
			}
			v2 = i32(0)
			goto l33
		l41:
			{
				t89 := int32(m.memory[uint32(v7)])
				v2 = t89 + i32(-48)
				if uint32(v2) <= uint32(i32(9)) {
					goto l42
				}
				v2 = i32(1)
				goto l33
			}
		l42:
			if v9 == i32(1) {
				goto l33
			}
			{
				t90 := int32(m.memory[int64(uint32(v7))+1])
				v5 = t90 + i32(-48)
				if uint32(v5) <= uint32(i32(9)) {
					goto l43
				}
				v2 = i32(1)
				goto l33
			}
		l43:
			v2 = v5 + v2*i32(10)
			if v9 == i32(2) {
				goto l33
			}
			{
				t91 := int32(m.memory[int64(uint32(v7))+2])
				v5 = t91 + i32(-48)
				if uint32(v5) <= uint32(i32(9)) {
					goto l44
				}
				v2 = i32(1)
				goto l33
			}
		l44:
			v2 = v5 + v2*i32(10)
			if v9 == i32(3) {
				goto l33
			}
			{
				t92 := int32(m.memory[int64(uint32(v7))+3])
				v5 = t92 + i32(-48)
				if uint32(v5) <= uint32(i32(9)) {
					goto l45
				}
				v2 = i32(1)
				goto l33
			}
		l45:
			v2 = v5 + v2*i32(10)
			if v9 == i32(4) {
				goto l33
			}
			{
				t93 := int32(m.memory[int64(uint32(v7))+4])
				v5 = t93 + i32(-48)
				if uint32(v5) <= uint32(i32(9)) {
					goto l46
				}
				v2 = i32(1)
				goto l33
			}
		l46:
			v2 = v5 + v2*i32(10)
			if v9 == i32(5) {
				goto l33
			}
			{
				t94 := int32(m.memory[int64(uint32(v7))+5])
				v5 = t94 + i32(-48)
				if uint32(v5) <= uint32(i32(9)) {
					goto l47
				}
				v2 = i32(1)
				goto l33
			}
		l47:
			v2 = v5 + v2*i32(10)
			if v9 == i32(6) {
				goto l33
			}
			{
				t95 := int32(m.memory[int64(uint32(v7))+6])
				v5 = t95 + i32(-48)
				if uint32(v5) <= uint32(i32(9)) {
					goto l48
				}
				v2 = i32(1)
				goto l33
			}
		l48:
			v5 = v5 + v2*i32(10)
			if v9 != i32(7) {
				goto l49
			}
			v2 = v5
			goto l33
		l49:
			v2 = i32(1)
			t96 := int32(m.memory[int64(uint32(v7))+7])
			v9 = t96 + i32(-48)
			if uint32(v9) > uint32(i32(9)) {
				goto l33
			}
			v2 = v9 + v5*i32(10)
		}
	l33:
		store32(m.memory[uint32(v10+i32(-8)):], uint32(v2))
		goto l1
	}
}
func (m *Module) fn420(v0 int32) {
	var v1, v2, v3, v4 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		v1 = t0
		if v1 == 0 {
			goto l0
		}
		t1 := int32(load32(m.memory[int64(uint32(v0))+12:]))
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
			m.fn7(i32(1274404), i32(46), i32(1274452))
			panic("unreachable")
		}
		if v3 == 0 {
			goto l2
		}
		if uint32(v4) > uint32(v1+i32(39)) {
			m.fn7(i32(1274468), i32(46), i32(1274516))
			panic("unreachable")
		}
	l2:
		m.fn5(v2)
	}
l0:
	{
		t5 := int32(load32(m.memory[int64(uint32(v0))+20:]))
		v1 = t5
		if v1 == i32(-1) {
			goto l4
		}
		if v1 == 0 {
			goto l4
		}
		t6 := int32(load32(m.memory[int64(uint32(v0))+24:]))
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
		if uint32(t8) < uint32(p9+v1) {
			m.fn7(i32(1274404), i32(46), i32(1274452))
			panic("unreachable")
		}
		if v3 == 0 {
			goto l6
		}
		if uint32(v4) > uint32(v1+i32(39)) {
			m.fn7(i32(1274468), i32(46), i32(1274516))
			panic("unreachable")
		}
	l6:
		m.fn5(v2)
	}
l4:
	{
		t10 := int32(load32(m.memory[int64(uint32(v0))+48:]))
		v1 = t10
		if v1 == 0 {
			goto l8
		}
		t11 := int32(load32(m.memory[int64(uint32(v0))+52:]))
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
		if uint32(t13) < uint32(p14+v1) {
			m.fn7(i32(1274404), i32(46), i32(1274452))
			panic("unreachable")
		}
		if v3 == 0 {
			goto l10
		}
		if uint32(v4) > uint32(v1+i32(39)) {
			m.fn7(i32(1274468), i32(46), i32(1274516))
			panic("unreachable")
		}
	l10:
		m.fn5(v2)
	}
l8:
	{
		t15 := int32(load32(m.memory[int64(uint32(v0))+60:]))
		v1 = t15
		if v1 == i32(-1) {
			goto l12
		}
		if v1 == 0 {
			goto l12
		}
		t16 := int32(load32(m.memory[int64(uint32(v0))+64:]))
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
		if uint32(t18) < uint32(p19+v1) {
			m.fn7(i32(1274404), i32(46), i32(1274452))
			panic("unreachable")
		}
		if v3 == 0 {
			goto l14
		}
		if uint32(v4) > uint32(v1+i32(39)) {
			m.fn7(i32(1274468), i32(46), i32(1274516))
			panic("unreachable")
		}
	l14:
		m.fn5(v2)
	}
l12:
	{
		t20 := int32(load32(m.memory[int64(uint32(v0))+88:]))
		v1 = t20
		if v1 == 0 {
			goto l16
		}
		t21 := int32(load32(m.memory[int64(uint32(v0))+92:]))
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
		if uint32(t23) < uint32(p24+v1) {
			m.fn7(i32(1274404), i32(46), i32(1274452))
			panic("unreachable")
		}
		if v3 == 0 {
			goto l18
		}
		if uint32(v4) > uint32(v1+i32(39)) {
			m.fn7(i32(1274468), i32(46), i32(1274516))
			panic("unreachable")
		}
	l18:
		m.fn5(v2)
	}
l16:
	{
		t25 := int32(load32(m.memory[int64(uint32(v0))+100:]))
		v1 = t25
		if v1 == i32(-1) {
			goto l20
		}
		if v1 == 0 {
			goto l20
		}
		t26 := int32(load32(m.memory[int64(uint32(v0))+104:]))
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
		if uint32(t28) < uint32(p29+v1) {
			m.fn7(i32(1274404), i32(46), i32(1274452))
			panic("unreachable")
		}
		if v3 == 0 {
			goto l22
		}
		if uint32(v4) > uint32(v1+i32(39)) {
			m.fn7(i32(1274468), i32(46), i32(1274516))
			panic("unreachable")
		}
	l22:
		m.fn5(v2)
	}
l20:
	{
		t30 := int32(load32(m.memory[int64(uint32(v0))+128:]))
		v1 = t30
		if v1 == 0 {
			goto l24
		}
		t31 := int32(load32(m.memory[int64(uint32(v0))+132:]))
		v2 = t31
		t32 := int32(load32(m.memory[uint32(v2+i32(-4)):]))
		v3 = t32
		v4 = v3 & i32(-8)
		t33 := v4
		v3 = v3 & i32(3)
		p34 := i32(8)
		if v3 != 0 {
			p34 = i32(4)
		}
		if uint32(t33) < uint32(p34+v1) {
			m.fn7(i32(1274404), i32(46), i32(1274452))
			panic("unreachable")
		}
		if v3 == 0 {
			goto l26
		}
		if uint32(v4) > uint32(v1+i32(39)) {
			m.fn7(i32(1274468), i32(46), i32(1274516))
			panic("unreachable")
		}
	l26:
		m.fn5(v2)
	}
l24:
	{
		t35 := int32(load32(m.memory[int64(uint32(v0))+140:]))
		v1 = t35
		if v1 == i32(-1) {
			goto l28
		}
		if v1 == 0 {
			goto l28
		}
		t36 := int32(load32(m.memory[int64(uint32(v0))+144:]))
		v2 = t36
		t37 := int32(load32(m.memory[uint32(v2+i32(-4)):]))
		v3 = t37
		v4 = v3 & i32(-8)
		t38 := v4
		v3 = v3 & i32(3)
		p39 := i32(8)
		if v3 != 0 {
			p39 = i32(4)
		}
		if uint32(t38) < uint32(p39+v1) {
			m.fn7(i32(1274404), i32(46), i32(1274452))
			panic("unreachable")
		}
		if v3 == 0 {
			goto l30
		}
		if uint32(v4) > uint32(v1+i32(39)) {
			m.fn7(i32(1274468), i32(46), i32(1274516))
			panic("unreachable")
		}
	l30:
		m.fn5(v2)
	}
l28:
	{
		t40 := int32(load32(m.memory[int64(uint32(v0))+168:]))
		v1 = t40
		if v1 == 0 {
			goto l32
		}
		t41 := int32(load32(m.memory[int64(uint32(v0))+172:]))
		v2 = t41
		t42 := int32(load32(m.memory[uint32(v2+i32(-4)):]))
		v3 = t42
		v4 = v3 & i32(-8)
		t43 := v4
		v3 = v3 & i32(3)
		p44 := i32(8)
		if v3 != 0 {
			p44 = i32(4)
		}
		if uint32(t43) < uint32(p44+v1) {
			m.fn7(i32(1274404), i32(46), i32(1274452))
			panic("unreachable")
		}
		if v3 == 0 {
			goto l34
		}
		if uint32(v4) > uint32(v1+i32(39)) {
			m.fn7(i32(1274468), i32(46), i32(1274516))
			panic("unreachable")
		}
	l34:
		m.fn5(v2)
	}
l32:
	{
		t45 := int32(load32(m.memory[int64(uint32(v0))+180:]))
		v1 = t45
		if v1 == i32(-1) {
			goto l36
		}
		if v1 == 0 {
			goto l36
		}
		t46 := int32(load32(m.memory[int64(uint32(v0))+184:]))
		v2 = t46
		t47 := int32(load32(m.memory[uint32(v2+i32(-4)):]))
		v3 = t47
		v4 = v3 & i32(-8)
		t48 := v4
		v3 = v3 & i32(3)
		p49 := i32(8)
		if v3 != 0 {
			p49 = i32(4)
		}
		if uint32(t48) < uint32(p49+v1) {
			m.fn7(i32(1274404), i32(46), i32(1274452))
			panic("unreachable")
		}
		if v3 == 0 {
			goto l38
		}
		if uint32(v4) > uint32(v1+i32(39)) {
			m.fn7(i32(1274468), i32(46), i32(1274516))
			panic("unreachable")
		}
	l38:
		m.fn5(v2)
	}
l36:
	{
		t50 := int32(load32(m.memory[int64(uint32(v0))+208:]))
		v1 = t50
		if v1 == 0 {
			goto l40
		}
		t51 := int32(load32(m.memory[int64(uint32(v0))+212:]))
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
		if uint32(t53) < uint32(p54+v1) {
			m.fn7(i32(1274404), i32(46), i32(1274452))
			panic("unreachable")
		}
		if v3 == 0 {
			goto l42
		}
		if uint32(v4) > uint32(v1+i32(39)) {
			m.fn7(i32(1274468), i32(46), i32(1274516))
			panic("unreachable")
		}
	l42:
		m.fn5(v2)
	}
l40:
	{
		t55 := int32(load32(m.memory[int64(uint32(v0))+220:]))
		v1 = t55
		if v1 == i32(-1) {
			goto l44
		}
		if v1 == 0 {
			goto l44
		}
		t56 := int32(load32(m.memory[int64(uint32(v0))+224:]))
		v2 = t56
		t57 := int32(load32(m.memory[uint32(v2+i32(-4)):]))
		v3 = t57
		v4 = v3 & i32(-8)
		t58 := v4
		v3 = v3 & i32(3)
		p59 := i32(8)
		if v3 != 0 {
			p59 = i32(4)
		}
		if uint32(t58) < uint32(p59+v1) {
			m.fn7(i32(1274404), i32(46), i32(1274452))
			panic("unreachable")
		}
		if v3 == 0 {
			goto l46
		}
		if uint32(v4) > uint32(v1+i32(39)) {
			m.fn7(i32(1274468), i32(46), i32(1274516))
			panic("unreachable")
		}
	l46:
		m.fn5(v2)
	}
l44:
	{
		t60 := int32(load32(m.memory[int64(uint32(v0))+248:]))
		v1 = t60
		if v1 == 0 {
			goto l48
		}
		t61 := int32(load32(m.memory[int64(uint32(v0))+252:]))
		v2 = t61
		t62 := int32(load32(m.memory[uint32(v2+i32(-4)):]))
		v3 = t62
		v4 = v3 & i32(-8)
		t63 := v4
		v3 = v3 & i32(3)
		p64 := i32(8)
		if v3 != 0 {
			p64 = i32(4)
		}
		if uint32(t63) < uint32(p64+v1) {
			m.fn7(i32(1274404), i32(46), i32(1274452))
			panic("unreachable")
		}
		if v3 == 0 {
			goto l50
		}
		if uint32(v4) > uint32(v1+i32(39)) {
			m.fn7(i32(1274468), i32(46), i32(1274516))
			panic("unreachable")
		}
	l50:
		m.fn5(v2)
	}
l48:
	{
		t65 := int32(load32(m.memory[int64(uint32(v0))+260:]))
		v1 = t65
		if v1 == i32(-1) {
			goto l52
		}
		if v1 == 0 {
			goto l52
		}
		t66 := int32(load32(m.memory[int64(uint32(v0))+264:]))
		v2 = t66
		t67 := int32(load32(m.memory[uint32(v2+i32(-4)):]))
		v3 = t67
		v4 = v3 & i32(-8)
		t68 := v4
		v3 = v3 & i32(3)
		p69 := i32(8)
		if v3 != 0 {
			p69 = i32(4)
		}
		if uint32(t68) < uint32(p69+v1) {
			m.fn7(i32(1274404), i32(46), i32(1274452))
			panic("unreachable")
		}
		if v3 == 0 {
			goto l54
		}
		if uint32(v4) > uint32(v1+i32(39)) {
			m.fn7(i32(1274468), i32(46), i32(1274516))
			panic("unreachable")
		}
	l54:
		m.fn5(v2)
	}
l52:
	{
		t70 := int32(load32(m.memory[int64(uint32(v0))+288:]))
		v1 = t70
		if v1 == 0 {
			goto l56
		}
		t71 := int32(load32(m.memory[int64(uint32(v0))+292:]))
		v2 = t71
		t72 := int32(load32(m.memory[uint32(v2+i32(-4)):]))
		v3 = t72
		v4 = v3 & i32(-8)
		t73 := v4
		v3 = v3 & i32(3)
		p74 := i32(8)
		if v3 != 0 {
			p74 = i32(4)
		}
		if uint32(t73) < uint32(p74+v1) {
			m.fn7(i32(1274404), i32(46), i32(1274452))
			panic("unreachable")
		}
		if v3 == 0 {
			goto l58
		}
		if uint32(v4) > uint32(v1+i32(39)) {
			m.fn7(i32(1274468), i32(46), i32(1274516))
			panic("unreachable")
		}
	l58:
		m.fn5(v2)
	}
l56:
	{
		t75 := int32(load32(m.memory[int64(uint32(v0))+300:]))
		v1 = t75
		if v1 == i32(-1) {
			goto l60
		}
		if v1 == 0 {
			goto l60
		}
		t76 := int32(load32(m.memory[int64(uint32(v0))+304:]))
		v2 = t76
		t77 := int32(load32(m.memory[uint32(v2+i32(-4)):]))
		v3 = t77
		v4 = v3 & i32(-8)
		t78 := v4
		v3 = v3 & i32(3)
		p79 := i32(8)
		if v3 != 0 {
			p79 = i32(4)
		}
		if uint32(t78) < uint32(p79+v1) {
			m.fn7(i32(1274404), i32(46), i32(1274452))
			panic("unreachable")
		}
		if v3 == 0 {
			goto l62
		}
		if uint32(v4) > uint32(v1+i32(39)) {
			m.fn7(i32(1274468), i32(46), i32(1274516))
			panic("unreachable")
		}
	l62:
		m.fn5(v2)
	}
l60:
	{
		t80 := int32(load32(m.memory[int64(uint32(v0))+328:]))
		v1 = t80
		if v1 == 0 {
			goto l64
		}
		t81 := int32(load32(m.memory[int64(uint32(v0))+332:]))
		v2 = t81
		t82 := int32(load32(m.memory[uint32(v2+i32(-4)):]))
		v3 = t82
		v4 = v3 & i32(-8)
		t83 := v4
		v3 = v3 & i32(3)
		p84 := i32(8)
		if v3 != 0 {
			p84 = i32(4)
		}
		if uint32(t83) < uint32(p84+v1) {
			m.fn7(i32(1274404), i32(46), i32(1274452))
			panic("unreachable")
		}
		if v3 == 0 {
			goto l66
		}
		if uint32(v4) > uint32(v1+i32(39)) {
			m.fn7(i32(1274468), i32(46), i32(1274516))
			panic("unreachable")
		}
	l66:
		m.fn5(v2)
	}
l64:
	{
		t85 := int32(load32(m.memory[int64(uint32(v0))+340:]))
		v1 = t85
		if v1 == i32(-1) {
			goto l68
		}
		if v1 == 0 {
			goto l68
		}
		t86 := int32(load32(m.memory[int64(uint32(v0))+344:]))
		v2 = t86
		t87 := int32(load32(m.memory[uint32(v2+i32(-4)):]))
		v3 = t87
		v4 = v3 & i32(-8)
		t88 := v4
		v3 = v3 & i32(3)
		p89 := i32(8)
		if v3 != 0 {
			p89 = i32(4)
		}
		if uint32(t88) < uint32(p89+v1) {
			m.fn7(i32(1274404), i32(46), i32(1274452))
			panic("unreachable")
		}
		if v3 == 0 {
			goto l70
		}
		if uint32(v4) > uint32(v1+i32(39)) {
			m.fn7(i32(1274468), i32(46), i32(1274516))
			panic("unreachable")
		}
	l70:
		m.fn5(v2)
	}
l68:
	{
		t90 := int32(load32(m.memory[int64(uint32(v0))+368:]))
		v1 = t90
		if v1 == 0 {
			goto l72
		}
		t91 := int32(load32(m.memory[int64(uint32(v0))+372:]))
		v2 = t91
		t92 := int32(load32(m.memory[uint32(v2+i32(-4)):]))
		v3 = t92
		v4 = v3 & i32(-8)
		t93 := v4
		v3 = v3 & i32(3)
		p94 := i32(8)
		if v3 != 0 {
			p94 = i32(4)
		}
		if uint32(t93) < uint32(p94+v1) {
			m.fn7(i32(1274404), i32(46), i32(1274452))
			panic("unreachable")
		}
		if v3 == 0 {
			goto l74
		}
		if uint32(v4) > uint32(v1+i32(39)) {
			m.fn7(i32(1274468), i32(46), i32(1274516))
			panic("unreachable")
		}
	l74:
		m.fn5(v2)
	}
l72:
	{
		t95 := int32(load32(m.memory[int64(uint32(v0))+380:]))
		v1 = t95
		if v1 == i32(-1) {
			return
		}
		if v1 == 0 {
			return
		}
		t96 := int32(load32(m.memory[int64(uint32(v0))+384:]))
		v3 = t96
		t97 := int32(load32(m.memory[uint32(v3+i32(-4)):]))
		v0 = t97
		v2 = v0 & i32(-8)
		t98 := v2
		v0 = v0 & i32(3)
		p99 := i32(8)
		if v0 != 0 {
			p99 = i32(4)
		}
		if uint32(t98) < uint32(p99+v1) {
			m.fn7(i32(1274404), i32(46), i32(1274452))
			panic("unreachable")
		}
		if v0 == 0 {
			goto l78
		}
		if uint32(v2) > uint32(v1+i32(39)) {
			m.fn7(i32(1274468), i32(46), i32(1274516))
			panic("unreachable")
		}
	l78:
		m.fn5(v3)
	}
}
func (m *Module) fn421(v0 int32) {
	var v1, v2, v3, v4 int32
	var v5 int64
	var v6, v7, v8, v9 int32
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
		l8:
			if v5 != i64(0) {
				goto l2
			}
		l3:
			{
				v6 = v4
				v4 = v6 + i32(8)
				v3 = v3 + i32(-128)
				t4 := int64(load64(m.memory[uint32(v6):]))
				v5 = t4 & i64(-0x7f7f7f7f7f7f7f80)
				if v5 == i64(-0x7f7f7f7f7f7f7f80) {
					goto l3
				}
			}
			v5 = v5 ^ i64(-0x7f7f7f7f7f7f7f80)
		l2:
			{
				v7 = v3 - int32(int64(bits.TrailingZeros64(uint64(v5))))<<1&i32(240)
				t5 := int32(load32(m.memory[uint32(v7+i32(-16)):]))
				v6 = t5
				if v6 == 0 {
					goto l4
				}
				t6 := int32(load32(m.memory[uint32(v7+i32(-12)):]))
				v8 = t6
				t7 := int32(load32(m.memory[uint32(v8+i32(-4)):]))
				v7 = t7
				v9 = v7 & i32(-8)
				t8 := v9
				v7 = v7 & i32(3)
				p9 := i32(8)
				if v7 != 0 {
					p9 = i32(4)
				}
				if uint32(t8) < uint32(p9+v6) {
					m.fn7(i32(1274404), i32(46), i32(1274452))
					panic("unreachable")
				}
				if v7 == 0 {
					goto l6
				}
				if uint32(v9) > uint32(v6+i32(39)) {
					m.fn7(i32(1274468), i32(46), i32(1274516))
					panic("unreachable")
				}
			l6:
				m.fn5(v8)
			}
		l4:
			v5 = (v5 + i64(-1)) & v5
			v2 = v2 + i32(-1)
			if v2 != 0 {
				goto l8
			}
		}
	l1:
		v4 = v1 << 4
		v3 = v4 + v1 + i32(25)
		if v3 == 0 {
			return
		}
		t10 := int32(load32(m.memory[uint32(v0):]))
		v6 = t10 - v4
		t11 := int32(load32(m.memory[uint32(v6+i32(-20)):]))
		v4 = t11
		v2 = v4 & i32(-8)
		t12 := v2
		v4 = v4 & i32(3)
		p13 := i32(8)
		if v4 != 0 {
			p13 = i32(4)
		}
		if uint32(t12) < uint32(p13+v3) {
			m.fn7(i32(1274404), i32(46), i32(1274452))
			panic("unreachable")
		}
		if v4 == 0 {
			goto l10
		}
		if uint32(v2) > uint32(v3+i32(39)) {
			m.fn7(i32(1274468), i32(46), i32(1274516))
			panic("unreachable")
		}
	l10:
		m.fn5(v6 + i32(-16))
	}
}
func (m *Module) fn422(v0, v1 int32) {
	var v2, v3, v4, v5, v6, v7, v8, v9, v10 int32
	t0 := m.g0
	v2 = t0 - i32(32)
	m.g0 = v2
	t1 := int32(load32(m.memory[uint32(v0):]))
	v3 = t1
	v4 = i32(0)
	store32(m.memory[uint32(v0):], uint32(i32(0)))
	t2 := int32(load32(m.memory[int64(uint32(v0))+12:]))
	v5 = t2
	t3 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	v6 = t3
	t4 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	v7 = t4
	{
		switch v3 {
		default:
			goto l0
		case 2:
			v8 = v5 * i32(12)
			if v5 == 0 {
				goto l4
			}
			v0 = v6 + i32(8)
			v9 = i32(0)
			v4 = i32(1)
			v3 = v8
		l5:
			{
				t5 := int32(load32(m.memory[uint32(v0+i32(-4)):]))
				t6 := int32(load32(m.memory[uint32(v0):]))
				m.fn144(v2+i32(8), t5, t6)
				t7 := int32(load32(m.memory[int64(uint32(v2))+12:]))
				if t7 != 0 {
					goto l4
				}
				v0 = v0 + i32(12)
				v9 = v9 + i32(1)
				v3 = v3 + i32(-12)
				if v3 != 0 {
					goto l5
				}
			}
			v4 = i32(0)
		l4:
			v0 = v6 + v8
			v3 = v5 * i32(-12)
			v8 = v5
		l7:
			{
				v10 = v8
				if v3 == 0 {
					goto l6
				}
				t8 := int32(load32(m.memory[uint32(v0+i32(-8)):]))
				t9 := int32(load32(m.memory[uint32(v0+i32(-4)):]))
				m.fn144(v2, t8, t9)
				v3 = v3 + i32(12)
				v8 = v10 + i32(-1)
				v0 = v0 + i32(-12)
				t10 := int32(load32(m.memory[int64(uint32(v2))+4:]))
				if t10 == 0 {
					goto l7
				}
			}
			if v4 == 0 {
				goto l8
			}
			if uint32(v10) < uint32(v9) {
				m.fn121(v9, v10, v5, i32(1072324))
				panic("unreachable")
			}
			m.fn203(v2+i32(20), v6+v9*i32(12), v10-v9, i32(1099582), i32(1))
			{
				t11 := int32(load32(m.memory[int64(uint32(v1))+8:]))
				v0 = t11
				t12 := int32(load32(m.memory[uint32(v1):]))
				if v0 != t12 {
					goto l10
				}
				m.fn310(v1)
			}
		l10:
			store32(m.memory[int64(uint32(v1))+8:], uint32(v0+i32(1)))
			t13 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			v0 = t13 + v0<<5
			store32(m.memory[uint32(v0):], uint32(i32(-0x7ffffffc)))
			t14 := int64(load64(m.memory[int64(uint32(v2))+20:]))
			store64(m.memory[int64(uint32(v0))+4:], uint64(t14))
			t15 := int32(load32(m.memory[int64(uint32(v2))+28:]))
			store32(m.memory[int64(uint32(v0))+12:], uint32(t15))
			store32(m.memory[int64(uint32(v0))+16:], uint32(i32(-1)))
			goto l8
		case 1:
			if v5 == 0 {
				if v7 == 0 {
					goto l0
				}
				t27 := int32(load32(m.memory[uint32(v6+i32(-4)):]))
				v0 = t27
				v5 = v0 & i32(-8)
				t28 := v5
				v0 = v0 & i32(3)
				p29 := i32(8)
				if v0 != 0 {
					p29 = i32(4)
				}
				v3 = v7 << 5
				if uint32(t28) < uint32(p29|v3) {
					m.fn7(i32(1274404), i32(46), i32(1274452))
					panic("unreachable")
				}
				if v0 == 0 {
					goto l20
				}
				if uint32(v5) <= uint32(v3+i32(39)) {
					goto l20
				}
				m.fn7(i32(1274468), i32(46), i32(1274516))
				panic("unreachable")
			}
			{
				t16 := int32(load32(m.memory[int64(uint32(v1))+8:]))
				v0 = t16
				t17 := int32(load32(m.memory[uint32(v1):]))
				if v0 != t17 {
					goto l12
				}
				m.fn310(v1)
			}
		l12:
			store32(m.memory[int64(uint32(v1))+8:], uint32(v0+i32(1)))
			t18 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			v0 = t18 + v0<<5
			store32(m.memory[int64(uint32(v0))+12:], uint32(v5))
			store32(m.memory[int64(uint32(v0))+8:], uint32(v6))
			store32(m.memory[int64(uint32(v0))+4:], uint32(v7))
			store32(m.memory[uint32(v0):], uint32(i32(-0x7ffffffd)))
			goto l0
		}
	l6:
		if v5 == 0 {
			goto l13
		}
	l8:
		v0 = v6
	l18:
		{
			t19 := int32(load32(m.memory[uint32(v0):]))
			v3 = t19
			if v3 == 0 {
				goto l14
			}
			t20 := int32(load32(m.memory[uint32(v0+i32(4)):]))
			v10 = t20
			t21 := int32(load32(m.memory[uint32(v10+i32(-4)):]))
			v8 = t21
			v9 = v8 & i32(-8)
			t22 := v9
			v8 = v8 & i32(3)
			p23 := i32(8)
			if v8 != 0 {
				p23 = i32(4)
			}
			if uint32(t22) < uint32(p23+v3) {
				m.fn7(i32(1274404), i32(46), i32(1274452))
				panic("unreachable")
			}
			if v8 == 0 {
				goto l16
			}
			if uint32(v9) > uint32(v3+i32(39)) {
				m.fn7(i32(1274468), i32(46), i32(1274516))
				panic("unreachable")
			}
		l16:
			m.fn5(v10)
		}
	l14:
		v0 = v0 + i32(12)
		v5 = v5 + i32(-1)
		if v5 != 0 {
			goto l18
		}
	l13:
		if v7 == 0 {
			goto l0
		}
		t24 := int32(load32(m.memory[uint32(v6+i32(-4)):]))
		v0 = t24
		v5 = v0 & i32(-8)
		t25 := v5
		v0 = v0 & i32(3)
		p26 := i32(8)
		if v0 != 0 {
			p26 = i32(4)
		}
		v3 = v7 * i32(12)
		if uint32(t25) < uint32(p26+v3) {
			m.fn7(i32(1274404), i32(46), i32(1274452))
			panic("unreachable")
		}
		if v0 == 0 {
			goto l20
		}
		if uint32(v5) > uint32(v3+i32(39)) {
			m.fn7(i32(1274468), i32(46), i32(1274516))
			panic("unreachable")
		}
		goto l20
	}
l20:
	m.fn5(v6)
l0:
	m.g0 = v2 + i32(32)
}
func (m *Module) fn423(v0 int32) {
	var v1, v2, v3, v4, v5, v6, v7 int32
	{
		t0 := int32(load32(m.memory[uint32(v0):]))
		switch t0 {
		case 0:
			return
		default:
			t1 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			v1 = t1
			{
				t2 := int32(load32(m.memory[int64(uint32(v0))+12:]))
				v2 = t2
				if v2 == 0 {
					goto l3
				}
				v3 = v1
			l8:
				{
					t3 := int32(load32(m.memory[uint32(v3):]))
					v4 = t3
					if v4 == 0 {
						goto l4
					}
					t4 := int32(load32(m.memory[uint32(v3+i32(4)):]))
					v5 = t4
					t5 := int32(load32(m.memory[uint32(v5+i32(-4)):]))
					v6 = t5
					v7 = v6 & i32(-8)
					t6 := v7
					v6 = v6 & i32(3)
					p7 := i32(8)
					if v6 != 0 {
						p7 = i32(4)
					}
					if uint32(t6) < uint32(p7+v4) {
						m.fn7(i32(1274404), i32(46), i32(1274452))
						panic("unreachable")
					}
					if v6 == 0 {
						goto l6
					}
					if uint32(v7) > uint32(v4+i32(39)) {
						m.fn7(i32(1274468), i32(46), i32(1274516))
						panic("unreachable")
					}
				l6:
					m.fn5(v5)
				}
			l4:
				v3 = v3 + i32(12)
				v2 = v2 + i32(-1)
				if v2 != 0 {
					goto l8
				}
			}
		l3:
			t8 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			v3 = t8
			if v3 == 0 {
				return
			}
			t9 := int32(load32(m.memory[uint32(v1+i32(-4)):]))
			v2 = t9
			v4 = v2 & i32(-8)
			t10 := v4
			v2 = v2 & i32(3)
			p11 := i32(8)
			if v2 != 0 {
				p11 = i32(4)
			}
			v3 = v3 * i32(12)
			if uint32(t10) < uint32(p11+v3) {
				m.fn7(i32(1274404), i32(46), i32(1274452))
				panic("unreachable")
			}
			if v2 == 0 {
				goto l10
			}
			if uint32(v4) > uint32(v3+i32(39)) {
				m.fn7(i32(1274468), i32(46), i32(1274516))
				panic("unreachable")
			}
			goto l10
		case 1:
			t12 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			v1 = t12
			{
				t13 := int32(load32(m.memory[int64(uint32(v0))+12:]))
				v2 = t13
				if v2 == 0 {
					goto l12
				}
				v3 = v1
			l13:
				m.fn330(v3)
				v3 = v3 + i32(32)
				v2 = v2 + i32(-1)
				if v2 != 0 {
					goto l13
				}
			}
		l12:
			t14 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			v3 = t14
			if v3 == 0 {
				return
			}
			t15 := int32(load32(m.memory[uint32(v1+i32(-4)):]))
			v2 = t15
			v4 = v2 & i32(-8)
			t16 := v4
			v2 = v2 & i32(3)
			p17 := i32(8)
			if v2 != 0 {
				p17 = i32(4)
			}
			v3 = v3 << 5
			if uint32(t16) < uint32(p17|v3) {
				m.fn7(i32(1274404), i32(46), i32(1274452))
				panic("unreachable")
			}
			if v2 == 0 {
				goto l10
			}
			if uint32(v4) <= uint32(v3+i32(39)) {
				goto l10
			}
			m.fn7(i32(1274468), i32(46), i32(1274516))
			panic("unreachable")
		}
	}
l10:
	m.fn5(v1)
}
func (m *Module) fn424(v0, v1, v2, v3, v4 int32) {
	var v5, v6, v7, v8, v9, v10, v11 int32
	t0 := m.g0
	v5 = t0 - i32(160)
	m.g0 = v5
	t1 := int32(load32(m.memory[int64(uint32(v1))+40:]))
	v6 = t1
	{
		{
			{
				{
					{
						{
							t2 := int32(load32(m.memory[int64(uint32(v1))+36:]))
							v7 = t2
							if v7 == 0 {
								goto l0
							}
							if v6 != i32(46) {
								goto l0
							}
							t3 := int64(load64(m.memory[int64(uint32(v7))+8:]))
							t4 := int64(load64(m.memory[uint32(v7+i32(16)):]))
							t5 := int64(load64(m.memory[uint32(v7+i32(24)):]))
							t6 := int64(load64(m.memory[uint32(v7+i32(32)):]))
							t7 := int64(load64(m.memory[uint32(v7+i32(40)):]))
							t8 := int64(load64(m.memory[uint32(v7+i32(46)):]))
							if t3^i64(7598524126653739637)|(t4^i64(4211821596982000243))|(t5^i64(7236833184807805812)|(t6^i64(4212112933405418351)))|(t7^i64(7310532362577407352)|(t8^i64(3471766489881142644))) == 0 {
								t24 := int32(load32(m.memory[int64(uint32(v1))+8:]))
								v7 = t24
								if v7 != i32(1) {
									goto l5
								}
								t25 := int32(load32(m.memory[int64(uint32(v1))+4:]))
								v7 = t25
								t26 := int32(m.memory[uint32(v7)])
								if t26 == i32(112) {
									goto l6
								}
								m.fn422(v4, v3)
								t27 := int32(m.memory[uint32(v7)])
								switch t27 + i32(-104) {
								case 0:
									t77 := int32(load32(m.memory[int64(uint32(v1))+16:]))
									t78 := v5 + i32(24)
									v8 = t77
									t79 := int32(load32(m.memory[int64(uint32(v1))+20:]))
									t80 := v8
									v9 = t79
									m.fn155(t78, t80, v9, i32(1071795), i32(46), i32(1079308), i32(13))
									v4 = i32(1)
									{
										t81 := int32(load32(m.memory[int64(uint32(v5))+24:]))
										v7 = t81
										if v7 == 0 {
											goto l27
										}
										t82 := int32(load32(m.memory[int64(uint32(v5))+28:]))
										m.fn733(v5+i32(16), v7, t82)
										t83 := int32(m.memory[int64(uint32(v5))+17])
										t84 := int32(m.memory[int64(uint32(v5))+16])
										p85 := i32(1)
										if t84 != 0 {
											p85 = t83
										}
										v4 = p85
									}
								l27:
									m.fn731(v5+i32(104), v1, v2)
									t86 := int64(load64(m.memory[int64(uint32(v5))+108:]))
									store64(m.memory[int64(uint32(v5))+80:], uint64(t86))
									t87 := int64(load64(m.memory[int64(uint32(v5))+116:]))
									store64(m.memory[int64(uint32(v5))+88:], uint64(t87))
									t88 := int64(load64(m.memory[int64(uint32(v5))+124:]))
									store64(m.memory[int64(uint32(v5))+96:], uint64(t88))
									{
										t89 := int32(load32(m.memory[int64(uint32(v5))+104:]))
										if t89 != i32(1) {
											t93 := int32(load32(m.memory[int64(uint32(v5))+88:]))
											t94 := v5
											v7 = t93
											store32(m.memory[int64(uint32(t94))+40:], uint32(v7))
											t95 := int64(load64(m.memory[int64(uint32(v5))+80:]))
											store64(m.memory[int64(uint32(v5))+32:], uint64(t95))
											t96 := int32(load32(m.memory[int64(uint32(v5))+100:]))
											store32(m.memory[int64(uint32(v5))+56:], uint32(t96))
											t97 := int64(load64(m.memory[int64(uint32(v5))+92:]))
											store64(m.memory[int64(uint32(v5))+48:], uint64(t97))
											v7 = v7 * i32(28)
											t98 := int32(load32(m.memory[int64(uint32(v5))+36:]))
											v6 = t98 + i32(-28)
											{
											l30:
												{
													v1 = v7
													if v1 == 0 {
														goto l29
													}
													v7 = v1 + i32(-28)
													v6 = v6 + i32(28)
													t99 := m.fn306(v6)
													if t99 != 0 {
														goto l30
													}
												}
												t100 := int64(load64(m.memory[int64(uint32(v5))+32:]))
												store64(m.memory[int64(uint32(v5))+64:], uint64(t100))
												t101 := int32(load32(m.memory[int64(uint32(v5))+40:]))
												t102 := v5
												v6 = t101
												store32(m.memory[int64(uint32(t102))+72:], uint32(v6))
												t103 := int32(load32(m.memory[int64(uint32(v5))+68:]))
												v10 = t103
												t104 := int32(load32(m.memory[uint32(v2+i32(200)):]))
												m.fn734(v5+i32(104), v8, v9, t104)
												t105 := int32(load32(m.memory[int64(uint32(v5))+108:]))
												v7 = t105
												{
													t106 := int32(load32(m.memory[int64(uint32(v5))+104:]))
													v11 = t106
													if v11 == i32(-1) {
														goto l31
													}
													t107 := int64(load64(m.memory[int64(uint32(v5))+120:]))
													store64(m.memory[int64(uint32(v0))+16:], uint64(t107))
													t108 := int64(load64(m.memory[int64(uint32(v5))+112:]))
													store64(m.memory[int64(uint32(v0))+8:], uint64(t108))
													store32(m.memory[int64(uint32(v0))+4:], uint32(v7))
													store32(m.memory[uint32(v0):], uint32(v11))
													m.fn718(v5 + i32(32))
													m.fn376(v5 + i32(48))
													goto l4
												}
											l31:
												t109 := fn735(v7)
												m.fn454(v10, v6, t109)
												store32(m.memory[int64(uint32(v5))+156:], uint32(i32(0)))
												store64(m.memory[int64(uint32(v5))+148:], uint64(i64(0x100000000)))
												m.fn455(v10, v6, v5+i32(148))
												m.fn736(v5+i32(80), v8, v9, v4, v2)
												{
													t110 := int32(load32(m.memory[int64(uint32(v5))+80:]))
													if t110 == i32(-1) {
														goto l32
													}
													t111 := int32(load32(m.memory[int64(uint32(v5))+88:]))
													store32(m.memory[int64(uint32(v5))+116:], uint32(t111))
													t112 := int64(load64(m.memory[int64(uint32(v5))+80:]))
													store64(m.memory[int64(uint32(v5))+108:], uint64(t112))
													store32(m.memory[int64(uint32(v5))+120:], uint32(i32(0)))
													store32(m.memory[int64(uint32(v5))+104:], uint32(i32(3)))
													m.fn706(v5+i32(64), v5+i32(104))
												}
											l32:
												t113 := int32(load32(m.memory[int64(uint32(v5))+72:]))
												store32(m.memory[int64(uint32(v5))+112:], uint32(t113))
												t114 := int64(load64(m.memory[int64(uint32(v5))+64:]))
												store64(m.memory[int64(uint32(v5))+104:], uint64(t114))
												t115 := int64(load64(m.memory[int64(uint32(v5))+148:]))
												store64(m.memory[int64(uint32(v5))+116:], uint64(t115))
												t116 := int32(load32(m.memory[int64(uint32(v5))+156:]))
												store32(m.memory[int64(uint32(v5))+124:], uint32(t116))
												{
													t117 := int32(load32(m.memory[int64(uint32(v3))+8:]))
													v7 = t117
													t118 := int32(load32(m.memory[uint32(v3):]))
													if v7 != t118 {
														goto l33
													}
													m.fn310(v3)
												}
											l33:
												store32(m.memory[int64(uint32(v3))+8:], uint32(v7+i32(1)))
												t119 := int32(load32(m.memory[int64(uint32(v3))+4:]))
												v7 = t119 + v7<<5
												t120 := int64(load64(m.memory[int64(uint32(v5))+104:]))
												store64(m.memory[uint32(v7):], uint64(t120))
												t121 := int64(load64(m.memory[int64(uint32(v5))+112:]))
												store64(m.memory[int64(uint32(v7))+8:], uint64(t121))
												t122 := int64(load64(m.memory[int64(uint32(v5))+120:]))
												store64(m.memory[int64(uint32(v7))+16:], uint64(t122))
												m.memory[int64(uint32(v7))+24] = byte(v4)
											}
										l29:
											t123 := int32(load32(m.memory[int64(uint32(v5))+56:]))
											v6 = t123
											t124 := int32(load32(m.memory[int64(uint32(v5))+52:]))
											v7 = t124
											t125 := int32(load32(m.memory[int64(uint32(v5))+48:]))
											store32(m.memory[int64(uint32(v5))+112:], uint32(t125))
											store32(m.memory[int64(uint32(v5))+104:], uint32(v7))
											store32(m.memory[int64(uint32(v5))+108:], uint32(v7))
											store32(m.memory[int64(uint32(v5))+116:], uint32(v7+v6<<5))
											m.fn450(v3, v5+i32(104))
											store32(m.memory[uint32(v0):], uint32(i32(-1)))
											if v1 != 0 {
												goto l4
											}
											m.fn718(v5 + i32(32))
											goto l4
										}
										t90 := int64(load64(m.memory[int64(uint32(v5))+96:]))
										store64(m.memory[int64(uint32(v0))+16:], uint64(t90))
										t91 := int64(load64(m.memory[int64(uint32(v5))+88:]))
										store64(m.memory[int64(uint32(v0))+8:], uint64(t91))
										t92 := int64(load64(m.memory[int64(uint32(v5))+80:]))
										store64(m.memory[uint32(v0):], uint64(t92))
										goto l4
									}
								case 8:
									goto l6
								default:
									goto l8
								}
							}
						}
					l0:
						t9 := int32(load32(m.memory[int64(uint32(v1))+8:]))
						if t9 != i32(5) {
							goto l2
						}
						t10 := int32(load32(m.memory[int64(uint32(v1))+4:]))
						v8 = t10
						t11 := int32(load32(m.memory[uint32(v8):]))
						t12 := int32(m.memory[uint32(v8+i32(4))])
						if t11^i32(1818386804)|(t12^i32(101)) != 0 {
							goto l2
						}
						if v7 == 0 {
							goto l2
						}
						if v6 != i32(47) {
							goto l2
						}
						t13 := int64(load64(m.memory[int64(uint32(v7))+8:]))
						t14 := int64(load64(m.memory[uint32(v7+i32(16)):]))
						t15 := int64(load64(m.memory[uint32(v7+i32(24)):]))
						t16 := int64(load64(m.memory[uint32(v7+i32(32)):]))
						t17 := int64(load64(m.memory[uint32(v7+i32(40)):]))
						t18 := int64(load64(m.memory[uint32(v7+i32(47)):]))
						if !(t13^i64(7598524126653739637)|(t14^i64(4211821596982000243))|(t15^i64(7236833184807805812)|(t16^i64(4212112933405418351)))|(t17^i64(7022301986425695608)|(t18^i64(3471766489628697185))) == 0) {
							goto l2
						}
						m.fn422(v4, v3)
						m.fn425(v5+i32(104), v1, v2)
						t19 := int32(load32(m.memory[int64(uint32(v5))+116:]))
						v7 = t19
						t20 := int32(load32(m.memory[int64(uint32(v5))+112:]))
						v6 = t20
						t21 := int32(load32(m.memory[int64(uint32(v5))+108:]))
						v1 = t21
						t22 := int32(load32(m.memory[int64(uint32(v5))+104:]))
						v4 = t22
						if v4 == i32(-1) {
							goto l3
						}
						t23 := int64(load64(m.memory[int64(uint32(v5))+120:]))
						store64(m.memory[int64(uint32(v0))+16:], uint64(t23))
						store32(m.memory[int64(uint32(v0))+12:], uint32(v7))
						store32(m.memory[int64(uint32(v0))+8:], uint32(v6))
						store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
						store32(m.memory[uint32(v0):], uint32(v4))
						goto l4
					}
				l3:
					{
						{
							t28 := int32(load32(m.memory[uint32(v3):]))
							t29 := int32(load32(m.memory[int64(uint32(v3))+8:]))
							t30 := v7
							v4 = t29
							if uint32(t30) <= uint32(t28-v4) {
								goto l9
							}
							m.fn197(v3, v4, v7, i32(8), i32(32))
							t31 := int32(load32(m.memory[int64(uint32(v3))+8:]))
							v4 = t31
							goto l10
						}
					l9:
						if v7 == 0 {
							goto l11
						}
					l10:
						v2 = v7 << 5
						if v2 == 0 {
							goto l11
						}
						t32 := int32(load32(m.memory[int64(uint32(v3))+4:]))
						memory_copy(m.memory, uint32(t32+v4<<5), uint32(v6), uint32(v2))
					}
				l11:
					store32(m.memory[int64(uint32(v3))+8:], uint32(v4+v7))
					if v1 == 0 {
						goto l2
					}
					m.fn21(v6, v1<<5, i32(8))
				l2:
					store32(m.memory[uint32(v0):], uint32(i32(-1)))
					goto l4
				l5:
					m.fn422(v4, v3)
					t33 := int32(load32(m.memory[int64(uint32(v1))+4:]))
					v6 = t33
					switch v7 + i32(-4) {
					default:
						goto l8
					case 0:
						t34 := int32(load32(m.memory[uint32(v6):]))
						if t34 != i32(1953720684) {
							goto l8
						}
						m.fn730(v5+i32(104), v1, v2, i32(0), i32(0), v7, i32(8), i32(0))
						t35 := int32(load32(m.memory[int64(uint32(v5))+116:]))
						v1 = t35
						t36 := int32(load32(m.memory[int64(uint32(v5))+112:]))
						v7 = t36
						t37 := int32(load32(m.memory[int64(uint32(v5))+108:]))
						v6 = t37
						t38 := int32(load32(m.memory[int64(uint32(v5))+104:]))
						v4 = t38
						if v4 == i32(-1) {
							store32(m.memory[int64(uint32(v5))+112:], uint32(v6))
							store32(m.memory[int64(uint32(v5))+104:], uint32(v7))
							store32(m.memory[int64(uint32(v5))+108:], uint32(v7))
							store32(m.memory[int64(uint32(v5))+116:], uint32(v7+v1<<5))
							m.fn450(v3, v5+i32(104))
							store32(m.memory[uint32(v0):], uint32(i32(-1)))
							goto l4
						}
						t39 := int64(load64(m.memory[int64(uint32(v5))+120:]))
						store64(m.memory[int64(uint32(v0))+16:], uint64(t39))
						store32(m.memory[int64(uint32(v0))+12:], uint32(v1))
						store32(m.memory[int64(uint32(v0))+8:], uint32(v7))
						store32(m.memory[int64(uint32(v0))+4:], uint32(v6))
						store32(m.memory[uint32(v0):], uint32(v4))
						goto l4
					case 3:
						t40 := int32(load32(m.memory[uint32(v6):]))
						t41 := int32(load32(m.memory[uint32(v6+i32(3)):]))
						if t40^i32(1952671091)|(t41^i32(1852795252)) != 0 {
							goto l8
						}
						goto l20
					case 6:
						t42 := int64(load64(m.memory[uint32(v6):]))
						t43 := int64(load16(m.memory[uint32(v6+i32(8)):]))
						if t42^i64(0x6f622d7865646e69)|(t43^i64(31076)) == 0 {
							goto l20
						}
						goto l8
					case 7:
						t44 := int64(load64(m.memory[uint32(v6):]))
						t45 := int64(load64(m.memory[uint32(v6+i32(3)):]))
						if t44^i64(7598748466401275497)|(t45^i64(7308344291584997477)) == 0 {
							goto l20
						}
						goto l8
					case 12:
						t46 := int64(load64(m.memory[uint32(v6):]))
						t47 := int64(load64(m.memory[uint32(v6+i32(8)):]))
						if t46^i64(0x666f2d656c626174)|(t47^i64(0x746e65746e6f632d)) != i64(0) {
							goto l8
						}
						goto l21
					case 8:
						t48 := int64(load64(m.memory[uint32(v6):]))
						t49 := int64(load32(m.memory[uint32(v6+i32(8)):]))
						if t48^i64(8243680141505620322)|(t49^i64(2036887649)) == 0 {
							goto l21
						}
						goto l8
					case 14:
						t50 := int64(load64(m.memory[uint32(v6):]))
						t51 := t50 ^ i64(8387218051550964833)
						v7 = v6 + i32(8)
						t52 := int64(load64(m.memory[uint32(v7):]))
						t53 := t51 | (t52 ^ i64(0x646e692d6c616369))
						v4 = v6 + i32(16)
						t54 := int64(load16(m.memory[uint32(v4):]))
						if t53|(t54^i64(30821)) == 0 {
							goto l21
						}
						t55 := int64(load64(m.memory[uint32(v6):]))
						t56 := int64(load64(m.memory[uint32(v7):]))
						t57 := int64(load16(m.memory[uint32(v4):]))
						if t55^i64(7021802808264125545)|(t56^i64(0x646e692d6e6f6974))|(t57^i64(30821)) == 0 {
							goto l21
						}
					}
				}
			l8:
				store32(m.memory[uint32(v0):], uint32(i32(-1)))
				goto l4
			l6:
				m.fn731(v5+i32(104), v1, v2)
				t58 := int64(load64(m.memory[int64(uint32(v5))+108:]))
				store64(m.memory[int64(uint32(v5))+80:], uint64(t58))
				t59 := int64(load64(m.memory[int64(uint32(v5))+116:]))
				store64(m.memory[int64(uint32(v5))+88:], uint64(t59))
				t60 := int64(load64(m.memory[int64(uint32(v5))+124:]))
				store64(m.memory[int64(uint32(v5))+96:], uint64(t60))
				t61 := int32(load32(m.memory[int64(uint32(v5))+104:]))
				if t61 != 0 {
					t129 := int64(load64(m.memory[int64(uint32(v5))+96:]))
					store64(m.memory[int64(uint32(v0))+16:], uint64(t129))
					t130 := int64(load64(m.memory[int64(uint32(v5))+88:]))
					store64(m.memory[int64(uint32(v0))+8:], uint64(t130))
					t131 := int64(load64(m.memory[int64(uint32(v5))+80:]))
					store64(m.memory[uint32(v0):], uint64(t131))
					goto l4
				}
				t62 := int32(load32(m.memory[int64(uint32(v5))+88:]))
				store32(m.memory[int64(uint32(v5))+144:], uint32(t62))
				t63 := int64(load64(m.memory[int64(uint32(v5))+80:]))
				store64(m.memory[int64(uint32(v5))+136:], uint64(t63))
				t64 := int32(load32(m.memory[int64(uint32(v5))+100:]))
				store32(m.memory[int64(uint32(v5))+72:], uint32(t64))
				t65 := int64(load64(m.memory[int64(uint32(v5))+92:]))
				store64(m.memory[int64(uint32(v5))+64:], uint64(t65))
				t66 := int32(load32(m.memory[int64(uint32(v1))+16:]))
				t67 := int32(load32(m.memory[int64(uint32(v1))+20:]))
				m.fn155(v5+i32(8), t66, t67, i32(1071795), i32(46), i32(1079068), i32(10))
				{
					t68 := int32(load32(m.memory[int64(uint32(v5))+8:]))
					v7 = t68
					if v7 == 0 {
						goto l23
					}
					t69 := int32(load32(m.memory[int64(uint32(v2))+200:]))
					t70 := int32(load32(m.memory[int64(uint32(v5))+12:]))
					t71 := m.fn732(t69, v7, t70)
					v7 = t71 & i32(255)
					if v7 != i32(2) {
						goto l24
					}
				}
			l23:
				m.fn422(v4, v3)
				{
					t72 := int32(load32(m.memory[int64(uint32(v3))+8:]))
					v7 = t72
					t73 := int32(load32(m.memory[uint32(v3):]))
					if v7 != t73 {
						goto l25
					}
					m.fn310(v3)
				}
			l25:
				store32(m.memory[int64(uint32(v3))+8:], uint32(v7+i32(1)))
				t74 := int32(load32(m.memory[int64(uint32(v3))+4:]))
				v7 = t74 + v7<<5
				store32(m.memory[uint32(v7):], uint32(i32(-0x80000000)))
				t75 := int64(load64(m.memory[int64(uint32(v5))+136:]))
				store64(m.memory[int64(uint32(v7))+4:], uint64(t75))
				t76 := int32(load32(m.memory[int64(uint32(v5))+144:]))
				store32(m.memory[int64(uint32(v7))+12:], uint32(t76))
				goto l26
			}
		l24:
			m.fn558(v4, v7&i32(1), v5+i32(136), v3)
		l26:
			{
				{
					t126 := int32(load32(m.memory[int64(uint32(v5))+72:]))
					v1 = t126
					if v1 != 0 {
						goto l34
					}
					m.fn376(v5 + i32(64))
					goto l35
				}
			l34:
				m.fn422(v4, v3)
				t127 := int32(load32(m.memory[int64(uint32(v5))+68:]))
				v7 = t127
				t128 := int32(load32(m.memory[int64(uint32(v5))+64:]))
				store32(m.memory[int64(uint32(v5))+112:], uint32(t128))
				store32(m.memory[int64(uint32(v5))+104:], uint32(v7))
				store32(m.memory[int64(uint32(v5))+116:], uint32(v7+v1<<5))
				store32(m.memory[int64(uint32(v5))+108:], uint32(v7))
				m.fn450(v3, v5+i32(104))
			}
		l35:
			store32(m.memory[uint32(v0):], uint32(i32(-1)))
			goto l4
		l21:
			m.fn346(v5+i32(104), v1, v2)
			t132 := int32(load32(m.memory[int64(uint32(v5))+116:]))
			v7 = t132
			t133 := int32(load32(m.memory[int64(uint32(v5))+112:]))
			v6 = t133
			t134 := int32(load32(m.memory[int64(uint32(v5))+108:]))
			v1 = t134
			{
				t135 := int32(load32(m.memory[int64(uint32(v5))+104:]))
				v4 = t135
				if v4 == i32(-1) {
					{
						{
							t137 := int32(load32(m.memory[uint32(v3):]))
							t138 := int32(load32(m.memory[int64(uint32(v3))+8:]))
							t139 := v7
							v4 = t138
							if uint32(t139) <= uint32(t137-v4) {
								goto l37
							}
							m.fn197(v3, v4, v7, i32(8), i32(32))
							t140 := int32(load32(m.memory[int64(uint32(v3))+8:]))
							v4 = t140
							goto l38
						}
					l37:
						if v7 == 0 {
							goto l39
						}
					l38:
						v2 = v7 << 5
						if v2 == 0 {
							goto l39
						}
						t141 := int32(load32(m.memory[int64(uint32(v3))+4:]))
						memory_copy(m.memory, uint32(t141+v4<<5), uint32(v6), uint32(v2))
					}
				l39:
					store32(m.memory[int64(uint32(v3))+8:], uint32(v4+v7))
					if v1 == 0 {
						goto l40
					}
					m.fn21(v6, v1<<5, i32(8))
				l40:
					store32(m.memory[uint32(v0):], uint32(i32(-1)))
					goto l4
				}
				t136 := int64(load64(m.memory[int64(uint32(v5))+120:]))
				store64(m.memory[int64(uint32(v0))+16:], uint64(t136))
				store32(m.memory[int64(uint32(v0))+12:], uint32(v7))
				store32(m.memory[int64(uint32(v0))+8:], uint32(v6))
				store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
				store32(m.memory[uint32(v0):], uint32(v4))
				goto l4
			}
		}
	l20:
		m.fn346(v5+i32(104), v1, v2)
		t142 := int32(load32(m.memory[int64(uint32(v5))+116:]))
		v1 = t142
		t143 := int32(load32(m.memory[int64(uint32(v5))+112:]))
		v7 = t143
		t144 := int32(load32(m.memory[int64(uint32(v5))+108:]))
		v6 = t144
		{
			t145 := int32(load32(m.memory[int64(uint32(v5))+104:]))
			v4 = t145
			if v4 == i32(-1) {
				goto l41
			}
			t146 := int64(load64(m.memory[int64(uint32(v5))+120:]))
			store64(m.memory[int64(uint32(v0))+16:], uint64(t146))
			store32(m.memory[int64(uint32(v0))+12:], uint32(v1))
			store32(m.memory[int64(uint32(v0))+8:], uint32(v7))
			store32(m.memory[int64(uint32(v0))+4:], uint32(v6))
			store32(m.memory[uint32(v0):], uint32(v4))
			goto l4
		}
	l41:
		store32(m.memory[int64(uint32(v5))+112:], uint32(v6))
		store32(m.memory[int64(uint32(v5))+104:], uint32(v7))
		store32(m.memory[int64(uint32(v5))+108:], uint32(v7))
		store32(m.memory[int64(uint32(v5))+116:], uint32(v7+v1<<5))
		m.fn450(v3, v5+i32(104))
		store32(m.memory[uint32(v0):], uint32(i32(-1)))
	}
l4:
	m.g0 = v5 + i32(160)
}
func (m *Module) fn425(v0, v1, v2 int32) {
	var v3 int32
	var v4, v5 int64
	var v6 int32
	t0 := m.g0
	v3 = t0 - i32(176)
	m.g0 = v3
	{
		{
			t1 := int32(m.memory[int64(uint32(i32(0)))+1294512])
			if t1 == 0 {
				goto l0
			}
			t2 := int64(load64(m.memory[int64(uint32(i32(0)))+1294504:]))
			v4 = t2
			t3 := int64(load64(m.memory[int64(uint32(i32(0)))+1294496:]))
			v5 = t3
			goto l1
		}
	l0:
		m.fn194(v3 + i32(8))
		m.memory[int64(uint32(i32(0)))+1294512] = byte(i32(1))
		t4 := int64(load64(m.memory[int64(uint32(v3))+16:]))
		v4 = t4
		store64(m.memory[int64(uint32(i32(0)))+1294504:], uint64(v4))
		t5 := int64(load64(m.memory[int64(uint32(v3))+8:]))
		v5 = t5
	}
l1:
	store64(m.memory[int64(uint32(v3))+48:], uint64(v5))
	store64(m.memory[int64(uint32(i32(0)))+1294496:], uint64(v5+i64(1)))
	store64(m.memory[int64(uint32(v3))+88:], uint64(i64(0)))
	store32(m.memory[int64(uint32(v3))+80:], uint32(i32(0)))
	store64(m.memory[int64(uint32(v3))+72:], uint64(i64(0x400000000)))
	store64(m.memory[int64(uint32(v3))+64:], uint64(i64(0)))
	store64(m.memory[int64(uint32(v3))+56:], uint64(v4))
	store64(m.memory[int64(uint32(v3))+8:], uint64(i64(0)))
	store64(m.memory[int64(uint32(v3))+16:], uint64(i64(0)))
	store64(m.memory[int64(uint32(v3))+24:], uint64(i64(0)))
	t6 := int64(load64(m.memory[int64(uint32(i32(0)))+1276264:]))
	store64(m.memory[int64(uint32(v3))+32:], uint64(t6))
	t7 := int64(load64(m.memory[int64(uint32(i32(0)))+1276272:]))
	store64(m.memory[int64(uint32(v3))+40:], uint64(t7))
	m.fn721(v3+i32(120), v1, v2, v3+i32(8), i32(1))
	{
		{
			t8 := int32(load32(m.memory[int64(uint32(v3))+120:]))
			if t8 == i32(-1) {
				goto l2
			}
			t9 := int64(load64(m.memory[int64(uint32(v3))+136:]))
			store64(m.memory[int64(uint32(v0))+16:], uint64(t9))
			t10 := int64(load64(m.memory[int64(uint32(v3))+128:]))
			store64(m.memory[int64(uint32(v0))+8:], uint64(t10))
			t11 := int64(load64(m.memory[int64(uint32(v3))+120:]))
			store64(m.memory[uint32(v0):], uint64(t11))
			m.fn357(v3 + i32(72))
			t12 := int32(load32(m.memory[int64(uint32(v3))+36:]))
			v0 = t12
			if v0 == 0 {
				goto l3
			}
			v2 = v0 << 4
			v0 = v2 + v0 + i32(25)
			if v0 == 0 {
				goto l3
			}
			t13 := int32(load32(m.memory[int64(uint32(v3))+32:]))
			v1 = t13 - v2
			t14 := int32(load32(m.memory[uint32(v1+i32(-20)):]))
			v2 = t14
			v6 = v2 & i32(-8)
			t15 := v6
			v2 = v2 & i32(3)
			p16 := i32(8)
			if v2 != 0 {
				p16 = i32(4)
			}
			if uint32(t15) < uint32(p16+v0) {
				m.fn7(i32(1274404), i32(46), i32(1274452))
				panic("unreachable")
			}
			if v2 == 0 {
				goto l5
			}
			if uint32(v6) > uint32(v0+i32(39)) {
				m.fn7(i32(1274468), i32(46), i32(1274516))
				panic("unreachable")
			}
		l5:
			m.fn5(v1 + i32(-16))
			goto l3
		}
	l2:
		t17 := v3
		v2 = v3 + i32(32)
		t18 := int64(load64(m.memory[int64(uint32(v2))+48:]))
		store64(m.memory[int64(uint32(t17))+168:], uint64(t18))
		t19 := int64(load64(m.memory[int64(uint32(v2))+40:]))
		store64(m.memory[int64(uint32(v3))+160:], uint64(t19))
		t20 := int64(load64(m.memory[int64(uint32(v2))+32:]))
		store64(m.memory[int64(uint32(v3))+152:], uint64(t20))
		t21 := int64(load64(m.memory[int64(uint32(v2))+24:]))
		store64(m.memory[int64(uint32(v3))+144:], uint64(t21))
		t22 := int64(load64(m.memory[int64(uint32(v2))+16:]))
		store64(m.memory[int64(uint32(v3))+136:], uint64(t22))
		t23 := int64(load64(m.memory[int64(uint32(v2))+8:]))
		store64(m.memory[int64(uint32(v3))+128:], uint64(t23))
		t24 := int64(load64(m.memory[uint32(v2):]))
		store64(m.memory[int64(uint32(v3))+120:], uint64(t24))
		m.fn331(v3+i32(100), v3+i32(120))
		{
			t25 := int32(load32(m.memory[int64(uint32(v3))+108:]))
			v2 = t25
			if v2 != 0 {
				goto l7
			}
			store64(m.memory[int64(uint32(v0))+8:], uint64(i64(8)))
			store64(m.memory[uint32(v0):], uint64(i64(0xffffffff)))
			m.fn357(v3 + i32(100))
			goto l3
		}
	l7:
		t26 := int32(load32(m.memory[int64(uint32(v3))+104:]))
		t27 := int32(load32(m.memory[int64(uint32(v3))+88:]))
		t28 := m.fn356(t26, v2, t27)
		store32(m.memory[int64(uint32(v3))+112:], uint32(t28))
		t29 := m.fn11(i32(32))
		v2 = t29
		if v2 == 0 {
			m.fn23(i32(8), i32(32))
			panic("unreachable")
		}
		t30 := int32(load32(m.memory[int64(uint32(v3))+116:]))
		store32(m.memory[int64(uint32(v2))+20:], uint32(t30))
		t31 := int64(load64(m.memory[int64(uint32(v3))+108:]))
		store64(m.memory[int64(uint32(v2))+12:], uint64(t31))
		t32 := int64(load64(m.memory[int64(uint32(v3))+100:]))
		store64(m.memory[int64(uint32(v2))+4:], uint64(t32))
		store32(m.memory[uint32(v2):], uint32(i32(-0x7ffffffe)))
		store32(m.memory[int64(uint32(v0))+12:], uint32(i32(1)))
		store32(m.memory[int64(uint32(v0))+8:], uint32(v2))
		store64(m.memory[uint32(v0):], uint64(i64(0x1ffffffff)))
	}
l3:
	m.g0 = v3 + i32(176)
}
func (m *Module) fn426(v0, v1, v2, v3, v4, v5 int32) {
	var v6, v7, v8, v9, v10, v11, v12, v13, v14, v15, v16, v17, v18, v19 int32
	var v20, v21 int64
	t0 := m.g0
	v6 = t0 - i32(96)
	m.g0 = v6
	t1 := int32(load32(m.memory[int64(uint32(v1))+28:]))
	v7 = t1
	t2 := int32(load32(m.memory[int64(uint32(v1))+32:]))
	v8 = v7 + t2*i32(44)
l1:
	{
		{
			{
				{
					v1 = v7
					if v1 == v8 {
						store32(m.memory[uint32(v0):], uint32(i32(-1)))
						goto l10
					}
					v7 = v1 + i32(44)
					t3 := int32(load32(m.memory[uint32(v1):]))
					if t3 == i32(-1) {
						goto l1
					}
					{
						t4 := int32(load32(m.memory[int64(uint32(v1))+8:]))
						v9 = t4
						if v9 != i32(5) {
							goto l2
						}
						t5 := int32(load32(m.memory[int64(uint32(v1))+4:]))
						v10 = t5
						t6 := int32(load32(m.memory[uint32(v10):]))
						t7 := int32(m.memory[uint32(v10+i32(4))])
						if t6^i32(1702129518)|(t7^i32(115)) != 0 {
							goto l2
						}
						t8 := int32(load32(m.memory[int64(uint32(v1))+36:]))
						v10 = t8
						if v10 == 0 {
							goto l2
						}
						t9 := int32(load32(m.memory[int64(uint32(v1))+40:]))
						if t9 != i32(54) {
							goto l2
						}
						t10 := int64(load64(m.memory[int64(uint32(v10))+8:]))
						t11 := int64(load64(m.memory[uint32(v10+i32(16)):]))
						t12 := int64(load64(m.memory[uint32(v10+i32(24)):]))
						t13 := int64(load64(m.memory[uint32(v10+i32(32)):]))
						t14 := int64(load64(m.memory[uint32(v10+i32(40)):]))
						t15 := int64(load64(m.memory[uint32(v10+i32(48)):]))
						t16 := int64(load64(m.memory[uint32(v10+i32(54)):]))
						if t10^i64(7598524126653739637)|(t11^i64(4211821596982000243))|(t12^i64(7236833184807805812)|(t13^i64(4212112933405418351)))|(t14^i64(8246155185163627896)|(t15^i64(7598805623977112421))|(t16^i64(3471766489779890548))) == 0 {
							{
								{
									t27 := int32(load32(m.memory[int64(uint32(v1))+32:]))
									v11 = t27
									if v11 != 0 {
										goto l11
									}
									v12 = i32(4)
									v9 = i32(0)
									goto l12
								}
							l11:
								t28 := int32(load32(m.memory[int64(uint32(v1))+28:]))
								v1 = t28
								v10 = v11 << 2
								t29 := m.fn11(v10)
								v12 = t29
								if v12 == 0 {
									m.fn16(i32(4), v10)
									panic("unreachable")
								}
								v10 = v11*i32(44) + i32(-44)
								t30 := int32(uint32(v10) / uint32(i32(44)))
								v13 = t30 + i32(1)
								v14 = v13 & i32(7)
								v9 = i32(0)
								if uint32(v10) < uint32(i32(308)) {
									goto l14
								}
								v9 = v13 & i32(0xffffff8)
								v15 = v13 << 2 & i32(0x3fffffe0)
								v13 = i32(0)
							l15:
								{
									v10 = v12 + v13
									store32(m.memory[uint32(v10):], uint32(v1))
									store32(m.memory[uint32(v10+i32(28)):], uint32(v1+i32(308)))
									store32(m.memory[uint32(v10+i32(24)):], uint32(v1+i32(264)))
									store32(m.memory[uint32(v10+i32(20)):], uint32(v1+i32(220)))
									store32(m.memory[uint32(v10+i32(16)):], uint32(v1+i32(176)))
									store32(m.memory[uint32(v10+i32(12)):], uint32(v1+i32(132)))
									store32(m.memory[uint32(v10+i32(8)):], uint32(v1+i32(88)))
									store32(m.memory[uint32(v10+i32(4)):], uint32(v1+i32(44)))
									v1 = v1 + i32(352)
									t31 := v15
									v13 = v13 + i32(32)
									if t31 != v13 {
										goto l15
									}
								}
								if v14 == 0 {
									goto l16
								}
							l14:
								v15 = v9 + v14
								v13 = v14 << 2
								v10 = v12 + v9<<2
							l17:
								store32(m.memory[uint32(v10):], uint32(v1))
								v10 = v10 + i32(4)
								v1 = v1 + i32(44)
								v13 = v13 + i32(-4)
								if v13 != 0 {
									goto l17
								}
								v9 = v15
							l16:
								v1 = int32(uint32(v9) >> 1)
								if v1 == 0 {
									goto l12
								}
								v16 = v12 + v9<<2
								v13 = i32(0)
								if v1 == i32(1) {
									goto l18
								}
								v17 = v1 & i32(1)
								v18 = v1 & i32(0xffffffe)
								v10 = v16 + i32(-4)
								v13 = i32(0)
								v1 = v12
							l19:
								{
									t32 := int32(load32(m.memory[uint32(v10):]))
									v15 = t32
									t33 := int32(load32(m.memory[uint32(v1):]))
									store32(m.memory[uint32(v10):], uint32(t33))
									store32(m.memory[uint32(v1):], uint32(v15))
									v15 = v16 + (v13^i32(0x3ffffffe))<<2
									t34 := int32(load32(m.memory[uint32(v15):]))
									v14 = t34
									t35 := v15
									v19 = v1 + i32(4)
									t36 := int32(load32(m.memory[uint32(v19):]))
									store32(m.memory[uint32(t35):], uint32(t36))
									store32(m.memory[uint32(v19):], uint32(v14))
									v10 = v10 + i32(-8)
									v1 = v1 + i32(8)
									t37 := v18
									v13 = v13 + i32(2)
									if t37 != v13 {
										goto l19
									}
								}
								if v17 == 0 {
									goto l12
								}
							l18:
								v1 = v12 + v13<<2
								t38 := int32(load32(m.memory[uint32(v1):]))
								v10 = t38
								t39 := v1
								v13 = v16 + (v13^i32(-1))<<2
								t40 := int32(load32(m.memory[uint32(v13):]))
								store32(m.memory[uint32(t39):], uint32(t40))
								store32(m.memory[uint32(v13):], uint32(v10))
							}
						l12:
							store32(m.memory[int64(uint32(v6))+40:], uint32(i32(5)))
							store32(m.memory[int64(uint32(v6))+36:], uint32(i32(1076783)))
							store32(m.memory[int64(uint32(v6))+32:], uint32(i32(49)))
							store32(m.memory[int64(uint32(v6))+28:], uint32(i32(1071318)))
							store32(m.memory[int64(uint32(v6))+20:], uint32(v12))
							store32(m.memory[int64(uint32(v6))+16:], uint32(v11))
							if v9 == 0 {
								goto l20
							}
						l45:
							{
								t41 := v6
								v15 = v9 + i32(-1)
								store32(m.memory[int64(uint32(t41))+24:], uint32(v15))
								{
									t42 := v12
									v14 = v15 << 2
									t43 := int32(load32(m.memory[uint32(t42+v14):]))
									v18 = t43
									t44 := int32(load32(m.memory[uint32(v18):]))
									if t44 == i32(-1) {
										goto l21
									}
									t45 := int32(load32(m.memory[int64(uint32(v18))+28:]))
									v19 = t45
									{
										{
											{
												t46 := int32(load32(m.memory[int64(uint32(v18))+32:]))
												v1 = t46
												t47 := int32(load32(m.memory[int64(uint32(v6))+16:]))
												if uint32(v1) <= uint32(t47-v15) {
													goto l22
												}
												m.fn197(v6+i32(16), v15, v1, i32(4), i32(4))
												t48 := int32(load32(m.memory[int64(uint32(v6))+20:]))
												v12 = t48
												t49 := int32(load32(m.memory[int64(uint32(v6))+24:]))
												v10 = t49
												goto l23
											}
										l22:
											v10 = v15
											v9 = v15
											if v1 == 0 {
												goto l24
											}
										l23:
											{
												{
													v11 = v1 * i32(44)
													v16 = v11 + i32(-44)
													t50 := int32(uint32(v16) / uint32(i32(44)))
													v1 = t50
													if v1&i32(7) != i32(7) {
														goto l25
													}
													v9 = v10
													v1 = v19
													goto l26
												}
											l25:
												t51 := v10
												v1 = (v1 + i32(1)) & i32(7)
												v9 = t51 + v1
												v13 = i32(0) - v1
												v10 = v12 + v10<<2
												v1 = v19
											l27:
												store32(m.memory[uint32(v10):], uint32(v1))
												v10 = v10 + i32(4)
												v1 = v1 + i32(44)
												v13 = v13 + i32(1)
												if v13 != 0 {
													goto l27
												}
											}
										l26:
											if uint32(v16) < uint32(i32(308)) {
												goto l28
											}
											v13 = v19 + v11
											v10 = v12 + v9<<2
										l29:
											store32(m.memory[uint32(v10):], uint32(v1))
											store32(m.memory[uint32(v10+i32(28)):], uint32(v1+i32(308)))
											store32(m.memory[uint32(v10+i32(24)):], uint32(v1+i32(264)))
											store32(m.memory[uint32(v10+i32(20)):], uint32(v1+i32(220)))
											store32(m.memory[uint32(v10+i32(16)):], uint32(v1+i32(176)))
											store32(m.memory[uint32(v10+i32(12)):], uint32(v1+i32(132)))
											store32(m.memory[uint32(v10+i32(8)):], uint32(v1+i32(88)))
											store32(m.memory[uint32(v10+i32(4)):], uint32(v1+i32(44)))
											v10 = v10 + i32(32)
											v9 = v9 + i32(8)
											v1 = v1 + i32(352)
											if v1 != v13 {
												goto l29
											}
										l28:
											store32(m.memory[int64(uint32(v6))+24:], uint32(v9))
											if uint32(v15) > uint32(v9) {
												m.fn121(v15, v9, v9, i32(1080576))
												panic("unreachable")
											}
										l24:
											{
												v1 = int32(uint32(v9-v15) >> 1)
												if v1 == 0 {
													goto l31
												}
												v11 = v12 + v14
												v19 = v12 + v9<<2
												v9 = i32(0)
												if v1 == i32(1) {
													goto l32
												}
												v17 = v1 & i32(1)
												v16 = v1 & i32(0x7ffffffe)
												v10 = v19 + i32(-4)
												v9 = i32(0)
												v1 = v11
											l33:
												{
													t52 := int32(load32(m.memory[uint32(v10):]))
													v13 = t52
													t53 := int32(load32(m.memory[uint32(v1):]))
													store32(m.memory[uint32(v10):], uint32(t53))
													store32(m.memory[uint32(v1):], uint32(v13))
													v13 = v19 + (v9^i32(0x3ffffffe))<<2
													t54 := int32(load32(m.memory[uint32(v13):]))
													v15 = t54
													t55 := v13
													v14 = v1 + i32(4)
													t56 := int32(load32(m.memory[uint32(v14):]))
													store32(m.memory[uint32(t55):], uint32(t56))
													store32(m.memory[uint32(v14):], uint32(v15))
													v10 = v10 + i32(-8)
													v1 = v1 + i32(8)
													t57 := v16
													v9 = v9 + i32(2)
													if t57 != v9 {
														goto l33
													}
												}
												if v17 == 0 {
													goto l31
												}
											l32:
												v1 = v11 + v9<<2
												t58 := int32(load32(m.memory[uint32(v1):]))
												v10 = t58
												t59 := v1
												v9 = v19 + (v9^i32(-1))<<2
												t60 := int32(load32(m.memory[uint32(v9):]))
												store32(m.memory[uint32(t59):], uint32(t60))
												store32(m.memory[uint32(v9):], uint32(v10))
											}
										l31:
											t61 := int32(load32(m.memory[uint32(v18):]))
											if t61 == i32(-1) {
												goto l21
											}
											t62 := int32(load32(m.memory[int64(uint32(v18))+8:]))
											if t62 != i32(5) {
												goto l21
											}
											t63 := int32(load32(m.memory[int64(uint32(v18))+4:]))
											v1 = t63
											t64 := int32(load32(m.memory[uint32(v1):]))
											t65 := int32(m.memory[uint32(v1+i32(4))])
											if t64^i32(1835102822)|(t65^i32(101)) != 0 {
												goto l21
											}
											t66 := int32(load32(m.memory[int64(uint32(v18))+36:]))
											v1 = t66
											if v1 == 0 {
												goto l21
											}
											t67 := int32(load32(m.memory[int64(uint32(v18))+40:]))
											if t67 != i32(49) {
												goto l21
											}
											v20 = i64(8462947847038399337)
											t68 := int64(load64(m.memory[int64(uint32(v1))+8:]))
											v21 = t68
											v21 = v21<<56 | v21&i64(0xff00)<<40 | (v21&i64(0xff0000)<<24 | v21&i64(0xff000000)<<8) | (int64(uint64(v21)>>8)&i64(0xff000000) | int64(uint64(v21)>>24)&i64(0xff0000) | (int64(uint64(v21)>>40)&i64(0xff00) | int64(uint64(v21)>>56)))
											if v21 != i64(8462947847038399337) {
												goto l34
											}
											v20 = i64(0x733a6e616d65733a)
											t69 := int64(load64(m.memory[uint32(v1+i32(16)):]))
											v21 = t69
											v21 = v21<<56 | v21&i64(0xff00)<<40 | (v21&i64(0xff0000)<<24 | v21&i64(0xff000000)<<8) | (int64(uint64(v21)>>8)&i64(0xff000000) | int64(uint64(v21)>>24)&i64(0xff0000) | (int64(uint64(v21)>>40)&i64(0xff00) | int64(uint64(v21)>>56)))
											if v21 != i64(0x733a6e616d65733a) {
												goto l34
											}
											v20 = i64(8386611181395471972)
											t70 := int64(load64(m.memory[uint32(v1+i32(24)):]))
											v21 = t70
											v21 = v21<<56 | v21&i64(0xff00)<<40 | (v21&i64(0xff0000)<<24 | v21&i64(0xff000000)<<8) | (int64(uint64(v21)>>8)&i64(0xff000000) | int64(uint64(v21)>>24)&i64(0xff0000) | (int64(uint64(v21)>>40)&i64(0xff00) | int64(uint64(v21)>>56)))
											if v21 != i64(8386611181395471972) {
												goto l34
											}
											v20 = i64(8026388073617978426)
											t71 := int64(load64(m.memory[uint32(v1+i32(32)):]))
											v21 = t71
											v21 = v21<<56 | v21&i64(0xff00)<<40 | (v21&i64(0xff0000)<<24 | v21&i64(0xff000000)<<8) | (int64(uint64(v21)>>8)&i64(0xff000000) | int64(uint64(v21)>>24)&i64(0xff0000) | (int64(uint64(v21)>>40)&i64(0xff00) | int64(uint64(v21)>>56)))
											if v21 != i64(8026388073617978426) {
												goto l34
											}
											v20 = i64(8677711278648222834)
											t72 := int64(load64(m.memory[uint32(v1+i32(40)):]))
											v21 = t72
											v21 = v21<<56 | v21&i64(0xff00)<<40 | (v21&i64(0xff0000)<<24 | v21&i64(0xff000000)<<8) | (int64(uint64(v21)>>8)&i64(0xff000000) | int64(uint64(v21)>>24)&i64(0xff0000) | (int64(uint64(v21)>>40)&i64(0xff00) | int64(uint64(v21)>>56)))
											if v21 != i64(8677711278648222834) {
												goto l34
											}
											v20 = i64(7023198066806763822)
											t73 := int64(load64(m.memory[uint32(v1+i32(48)):]))
											v21 = t73
											v21 = v21<<56 | v21&i64(0xff00)<<40 | (v21&i64(0xff0000)<<24 | v21&i64(0xff000000)<<8) | (int64(uint64(v21)>>8)&i64(0xff000000) | int64(uint64(v21)>>24)&i64(0xff0000) | (int64(uint64(v21)>>40)&i64(0xff00) | int64(uint64(v21)>>56)))
											if v21 != i64(7023198066806763822) {
												goto l34
											}
											t74 := int32(m.memory[uint32(v1+i32(56))])
											v1 = t74 + i32(-48)
											goto l35
										}
									l34:
										p75 := i32(1)
										if uint64(v21) < uint64(v20) {
											p75 = i32(-1)
										}
										v1 = p75
									}
								l35:
									if v1 != 0 {
										goto l21
									}
									t76 := int32(load32(m.memory[int64(uint32(v18))+32:]))
									v1 = t76
									if v1 == 0 {
										goto l21
									}
									v10 = v1 * i32(44)
									t77 := int32(load32(m.memory[int64(uint32(v18))+28:]))
									v1 = t77
								l40:
									{
										t78 := int32(load32(m.memory[uint32(v1):]))
										if t78 == i32(-1) {
											goto l36
										}
										t79 := int32(load32(m.memory[uint32(v1+i32(8)):]))
										if t79 != i32(8) {
											goto l36
										}
										t80 := int32(load32(m.memory[uint32(v1+i32(4)):]))
										t81 := int64(load64(m.memory[uint32(t80):]))
										if t81 != i64(8678262954333332852) {
											goto l36
										}
										t82 := int32(load32(m.memory[uint32(v1+i32(36)):]))
										v9 = t82
										if v9 == 0 {
											goto l36
										}
										t83 := int32(load32(m.memory[uint32(v1+i32(40)):]))
										if t83 != i32(49) {
											goto l36
										}
										v20 = i64(8462947847038399337)
										{
											{
												t84 := int64(load64(m.memory[int64(uint32(v9))+8:]))
												v21 = t84
												v21 = v21<<56 | v21&i64(0xff00)<<40 | (v21&i64(0xff0000)<<24 | v21&i64(0xff000000)<<8) | (int64(uint64(v21)>>8)&i64(0xff000000) | int64(uint64(v21)>>24)&i64(0xff0000) | (int64(uint64(v21)>>40)&i64(0xff00) | int64(uint64(v21)>>56)))
												if v21 != i64(8462947847038399337) {
													goto l37
												}
												v20 = i64(0x733a6e616d65733a)
												t85 := int64(load64(m.memory[uint32(v9+i32(16)):]))
												v21 = t85
												v21 = v21<<56 | v21&i64(0xff00)<<40 | (v21&i64(0xff0000)<<24 | v21&i64(0xff000000)<<8) | (int64(uint64(v21)>>8)&i64(0xff000000) | int64(uint64(v21)>>24)&i64(0xff0000) | (int64(uint64(v21)>>40)&i64(0xff00) | int64(uint64(v21)>>56)))
												if v21 != i64(0x733a6e616d65733a) {
													goto l37
												}
												v20 = i64(8386611181395471972)
												t86 := int64(load64(m.memory[uint32(v9+i32(24)):]))
												v21 = t86
												v21 = v21<<56 | v21&i64(0xff00)<<40 | (v21&i64(0xff0000)<<24 | v21&i64(0xff000000)<<8) | (int64(uint64(v21)>>8)&i64(0xff000000) | int64(uint64(v21)>>24)&i64(0xff0000) | (int64(uint64(v21)>>40)&i64(0xff00) | int64(uint64(v21)>>56)))
												if v21 != i64(8386611181395471972) {
													goto l37
												}
												v20 = i64(8026388073617978426)
												t87 := int64(load64(m.memory[uint32(v9+i32(32)):]))
												v21 = t87
												v21 = v21<<56 | v21&i64(0xff00)<<40 | (v21&i64(0xff0000)<<24 | v21&i64(0xff000000)<<8) | (int64(uint64(v21)>>8)&i64(0xff000000) | int64(uint64(v21)>>24)&i64(0xff0000) | (int64(uint64(v21)>>40)&i64(0xff00) | int64(uint64(v21)>>56)))
												if v21 != i64(8026388073617978426) {
													goto l37
												}
												v20 = i64(8677711278648222834)
												t88 := int64(load64(m.memory[uint32(v9+i32(40)):]))
												v21 = t88
												v21 = v21<<56 | v21&i64(0xff00)<<40 | (v21&i64(0xff0000)<<24 | v21&i64(0xff000000)<<8) | (int64(uint64(v21)>>8)&i64(0xff000000) | int64(uint64(v21)>>24)&i64(0xff0000) | (int64(uint64(v21)>>40)&i64(0xff00) | int64(uint64(v21)>>56)))
												if v21 != i64(8677711278648222834) {
													goto l37
												}
												v20 = i64(7023198066806763822)
												t89 := int64(load64(m.memory[uint32(v9+i32(48)):]))
												v21 = t89
												v21 = v21<<56 | v21&i64(0xff00)<<40 | (v21&i64(0xff0000)<<24 | v21&i64(0xff000000)<<8) | (int64(uint64(v21)>>8)&i64(0xff000000) | int64(uint64(v21)>>24)&i64(0xff0000) | (int64(uint64(v21)>>40)&i64(0xff00) | int64(uint64(v21)>>56)))
												if v21 != i64(7023198066806763822) {
													goto l37
												}
												t90 := int32(m.memory[uint32(v9+i32(56))])
												v9 = t90 + i32(-48)
												goto l38
											}
										l37:
											p91 := i32(1)
											if uint64(v21) < uint64(v20) {
												p91 = i32(-1)
											}
											v9 = p91
										}
									l38:
										if v9 == 0 {
											goto l39
										}
									}
								l36:
									v1 = v1 + i32(44)
									v10 = v10 + i32(-44)
									if v10 == 0 {
										goto l21
									}
									goto l40
								l39:
									m.fn346(v6+i32(48), v1, v2)
									t92 := int32(load32(m.memory[int64(uint32(v6))+60:]))
									v1 = t92
									t93 := int32(load32(m.memory[int64(uint32(v6))+56:]))
									v13 = t93
									t94 := int32(load32(m.memory[int64(uint32(v6))+52:]))
									v10 = t94
									{
										t95 := int32(load32(m.memory[int64(uint32(v6))+48:]))
										v9 = t95
										if v9 == i32(-1) {
											goto l41
										}
										t96 := int64(load64(m.memory[int64(uint32(v6))+64:]))
										store64(m.memory[int64(uint32(v0))+16:], uint64(t96))
										store32(m.memory[int64(uint32(v0))+12:], uint32(v1))
										store32(m.memory[int64(uint32(v0))+8:], uint32(v13))
										store32(m.memory[int64(uint32(v0))+4:], uint32(v10))
										store32(m.memory[uint32(v0):], uint32(v9))
										t97 := int32(load32(m.memory[int64(uint32(v6))+16:]))
										v1 = t97
										if v1 == 0 {
											goto l10
										}
										t98 := int32(load32(m.memory[int64(uint32(v6))+20:]))
										m.fn21(t98, v1<<2, i32(4))
										goto l10
									}
								l41:
									{
										{
											t99 := int32(load32(m.memory[uint32(v5):]))
											t100 := int32(load32(m.memory[int64(uint32(v5))+8:]))
											t101 := v1
											v9 = t100
											if uint32(t101) <= uint32(t99-v9) {
												goto l42
											}
											m.fn197(v5, v9, v1, i32(8), i32(32))
											t102 := int32(load32(m.memory[int64(uint32(v5))+8:]))
											v9 = t102
											goto l43
										}
									l42:
										if v1 == 0 {
											goto l44
										}
									l43:
										v15 = v1 << 5
										if v15 == 0 {
											goto l44
										}
										t103 := int32(load32(m.memory[int64(uint32(v5))+4:]))
										memory_copy(m.memory, uint32(t103+v9<<5), uint32(v13), uint32(v15))
									}
								l44:
									store32(m.memory[int64(uint32(v5))+8:], uint32(v9+v1))
									if v10 == 0 {
										goto l21
									}
									m.fn21(v13, v10<<5, i32(8))
								}
							l21:
								t104 := int32(load32(m.memory[int64(uint32(v6))+24:]))
								v9 = t104
								if v9 != 0 {
									goto l45
								}
							}
						l20:
							t105 := int32(load32(m.memory[int64(uint32(v6))+16:]))
							v1 = t105
							if v1 == 0 {
								goto l1
							}
							{
								t106 := int32(load32(m.memory[int64(uint32(v6))+20:]))
								v9 = t106
								t107 := int32(load32(m.memory[uint32(v9+i32(-4)):]))
								v10 = t107
								v13 = v10 & i32(-8)
								t108 := v13
								v10 = v10 & i32(3)
								p109 := i32(8)
								if v10 != 0 {
									p109 = i32(4)
								}
								v1 = v1 << 2
								if uint32(t108) < uint32(p109+v1) {
									m.fn7(i32(1274404), i32(46), i32(1274452))
									panic("unreachable")
								}
								if v10 == 0 {
									goto l47
								}
								if uint32(v13) > uint32(v1+i32(39)) {
									m.fn7(i32(1274468), i32(46), i32(1274516))
									panic("unreachable")
								}
							l47:
								m.fn5(v9)
								goto l1
							}
						}
					}
				l2:
					t17 := int32(load32(m.memory[int64(uint32(v1))+36:]))
					v10 = t17
					if v10 == 0 {
						goto l1
					}
					t18 := int32(load32(m.memory[int64(uint32(v1))+40:]))
					if t18 != i32(49) {
						goto l1
					}
					t19 := int64(load64(m.memory[int64(uint32(v10))+8:]))
					t20 := int64(load64(m.memory[uint32(v10+i32(16)):]))
					t21 := int64(load64(m.memory[uint32(v10+i32(24)):]))
					t22 := int64(load64(m.memory[uint32(v10+i32(32)):]))
					t23 := int64(load64(m.memory[uint32(v10+i32(40)):]))
					t24 := int64(load64(m.memory[uint32(v10+i32(48)):]))
					t25 := int64(m.memory[uint32(v10+i32(56))])
					if t19^i64(7598524126653739637)|(t20^i64(4211821596982000243))|(t21^i64(7236833184807805812)|(t22^i64(4212112933405418351)))|(t23^i64(8242777485443100024)|(t24^i64(3328505815511955297))|(t25^i64(48))) != i64(0) {
						goto l1
					}
					t26 := int32(load32(m.memory[int64(uint32(v1))+4:]))
					v10 = t26
					switch v9 + i32(-1) {
					case 0:
						t158 := int32(m.memory[uint32(v10)])
						if t158 != i32(103) {
							goto l1
						}
						m.fn426(v6+i32(16), v1, v2, v3, v4, v5)
						t159 := int32(load32(m.memory[int64(uint32(v6))+16:]))
						if t159 == i32(-1) {
							goto l1
						}
						t160 := int64(load64(m.memory[int64(uint32(v6))+32:]))
						store64(m.memory[int64(uint32(v0))+16:], uint64(t160))
						t161 := int64(load64(m.memory[int64(uint32(v6))+24:]))
						store64(m.memory[int64(uint32(v0))+8:], uint64(t161))
						t162 := int64(load64(m.memory[int64(uint32(v6))+16:]))
						store64(m.memory[uint32(v0):], uint64(t162))
						goto l10
					case 3:
						t112 := int32(load32(m.memory[uint32(v10):]))
						if t112 == i32(1952671090) {
							goto l49
						}
						t113 := int32(load32(m.memory[uint32(v10):]))
						if t113 == i32(1752457584) {
							goto l49
						}
						t114 := int32(load32(m.memory[uint32(v10):]))
						if t114 != i32(1701734764) {
							goto l1
						}
						goto l49
					case 4:
						goto l6
					case 6:
						t115 := int32(load32(m.memory[uint32(v10):]))
						t116 := t115 ^ i32(1768713317)
						v9 = v10 + i32(3)
						t117 := int32(load32(m.memory[uint32(v9):]))
						if t116|(t117^i32(1702064233)) == 0 {
							goto l49
						}
						t118 := int32(load32(m.memory[uint32(v10):]))
						t119 := int32(load32(m.memory[uint32(v9):]))
						if t118^i32(2037149552)|(t119^i32(0x6e6f6779)) == 0 {
							goto l49
						}
						t120 := int32(load32(m.memory[uint32(v10):]))
						t121 := int32(load32(m.memory[uint32(v9):]))
						if t120^i32(1953522019)|(t121^i32(1852795252)) != 0 {
							goto l1
						}
						goto l49
					case 8:
						t110 := int64(load64(m.memory[uint32(v10):]))
						t111 := int64(m.memory[uint32(v10+i32(8))])
						if !(t110^i64(8031153322804014947)|(t111^i64(114)) == 0) {
							goto l1
						}
						goto l49
					case 11:
						goto l9
					default:
						goto l1
					}
				}
			l9:
				t122 := int64(load64(m.memory[uint32(v10):]))
				t123 := int64(load32(m.memory[uint32(v10+i32(8)):]))
				if t122^i64(8299410013776213347)|(t123^i64(1701863784)) != i64(0) {
					goto l1
				}
			}
		l49:
			t124 := int32(load32(m.memory[uint32(v1+i32(32)):]))
			v9 = t124 * i32(44)
			t125 := int32(load32(m.memory[uint32(v1+i32(28)):]))
			v10 = t125
		l57:
			if v9 == 0 {
				goto l1
			}
			{
				t126 := int32(load32(m.memory[uint32(v10):]))
				if t126 == i32(-1) {
					goto l50
				}
				{
					{
						t127 := int32(load32(m.memory[uint32(v10+i32(8)):]))
						v13 = t127
						if v13 != i32(1) {
							goto l51
						}
						t128 := int32(load32(m.memory[uint32(v10+i32(4)):]))
						t129 := int32(m.memory[uint32(t128)])
						if t129 != i32(112) {
							goto l50
						}
						t130 := int32(load32(m.memory[uint32(v10+i32(36)):]))
						v13 = t130
						if v13 == 0 {
							goto l50
						}
						t131 := int32(load32(m.memory[uint32(v10+i32(40)):]))
						if t131 != i32(46) {
							goto l50
						}
						t132 := int64(load64(m.memory[int64(uint32(v13))+8:]))
						t133 := int64(load64(m.memory[uint32(v13+i32(16)):]))
						t134 := int64(load64(m.memory[uint32(v13+i32(24)):]))
						t135 := int64(load64(m.memory[uint32(v13+i32(32)):]))
						t136 := int64(load64(m.memory[uint32(v13+i32(40)):]))
						t137 := int64(load64(m.memory[uint32(v13+i32(46)):]))
						if !(t132^i64(7598524126653739637)|(t133^i64(4211821596982000243))|(t134^i64(7236833184807805812)|(t135^i64(4212112933405418351)))|(t136^i64(7310532362577407352)|(t137^i64(3471766489881142644))) == 0) {
							goto l50
						}
						goto l52
					}
				l51:
					if v13 != i32(4) {
						goto l50
					}
					t138 := int32(load32(m.memory[uint32(v10+i32(4)):]))
					t139 := int32(load32(m.memory[uint32(t138):]))
					if t139 != i32(1953720684) {
						goto l50
					}
					t140 := int32(load32(m.memory[uint32(v10+i32(36)):]))
					v13 = t140
					if v13 == 0 {
						goto l50
					}
					t141 := int32(load32(m.memory[uint32(v10+i32(40)):]))
					if t141 != i32(46) {
						goto l50
					}
					t142 := int64(load64(m.memory[int64(uint32(v13))+8:]))
					t143 := int64(load64(m.memory[uint32(v13+i32(16)):]))
					t144 := int64(load64(m.memory[uint32(v13+i32(24)):]))
					t145 := int64(load64(m.memory[uint32(v13+i32(32)):]))
					t146 := int64(load64(m.memory[uint32(v13+i32(40)):]))
					t147 := int64(load64(m.memory[uint32(v13+i32(46)):]))
					if t142^i64(7598524126653739637)|(t143^i64(4211821596982000243))|(t144^i64(7236833184807805812)|(t145^i64(4212112933405418351)))|(t146^i64(7310532362577407352)|(t147^i64(3471766489881142644))) != i64(0) {
						goto l50
					}
				}
			l52:
				m.fn346(v6+i32(16), v1, v2)
				t148 := int32(load32(m.memory[int64(uint32(v6))+28:]))
				v1 = t148
				t149 := int32(load32(m.memory[int64(uint32(v6))+24:]))
				v13 = t149
				t150 := int32(load32(m.memory[int64(uint32(v6))+20:]))
				v10 = t150
				{
					t151 := int32(load32(m.memory[int64(uint32(v6))+16:]))
					v9 = t151
					if v9 == i32(-1) {
						{
							{
								t153 := int32(load32(m.memory[uint32(v4):]))
								t154 := int32(load32(m.memory[int64(uint32(v4))+8:]))
								t155 := v1
								v9 = t154
								if uint32(t155) <= uint32(t153-v9) {
									goto l54
								}
								m.fn197(v4, v9, v1, i32(8), i32(32))
								t156 := int32(load32(m.memory[int64(uint32(v4))+8:]))
								v9 = t156
								goto l55
							}
						l54:
							if v1 == 0 {
								goto l56
							}
						l55:
							v15 = v1 << 5
							if v15 == 0 {
								goto l56
							}
							t157 := int32(load32(m.memory[int64(uint32(v4))+4:]))
							memory_copy(m.memory, uint32(t157+v9<<5), uint32(v13), uint32(v15))
						}
					l56:
						store32(m.memory[int64(uint32(v4))+8:], uint32(v9+v1))
						if v10 == 0 {
							goto l1
						}
						m.fn21(v13, v10<<5, i32(8))
						goto l1
					}
					t152 := int64(load64(m.memory[int64(uint32(v6))+32:]))
					store64(m.memory[int64(uint32(v0))+16:], uint64(t152))
					store32(m.memory[int64(uint32(v0))+12:], uint32(v1))
					store32(m.memory[int64(uint32(v0))+8:], uint32(v13))
					store32(m.memory[int64(uint32(v0))+4:], uint32(v10))
					store32(m.memory[uint32(v0):], uint32(v9))
					goto l10
				}
			}
		l50:
			v10 = v10 + i32(44)
			v9 = v9 + i32(-44)
			goto l57
		}
	l6:
		t163 := int32(load32(m.memory[uint32(v10):]))
		t164 := int32(m.memory[uint32(v10+i32(4))])
		if t163^i32(1835102822)|(t164^i32(101)) != 0 {
			goto l1
		}
		t165 := int32(load32(m.memory[uint32(v1+i32(16)):]))
		t166 := int32(load32(m.memory[uint32(v1+i32(20)):]))
		m.fn155(v6+i32(8), t165, t166, i32(1076729), i32(54), i32(1073833), i32(5))
		t167 := int32(load32(m.memory[int64(uint32(v6))+8:]))
		v10 = t167
		p168 := i32(1)
		if v10 != 0 {
			p168 = v10
		}
		v15 = p168
		{
			t169 := int32(load32(m.memory[int64(uint32(v6))+12:]))
			p170 := i32(0)
			if v10 != 0 {
				p170 = t169
			}
			v14 = p170
			switch v14 + i32(-6) {
			default:
				goto l59
			case 5:
				t171 := int64(load64(m.memory[uint32(v15):]))
				t172 := int64(load64(m.memory[uint32(v15+i32(3)):]))
				if t171^i64(7887331463663149424)|(t172^i64(0x7265626d756e2d65)) == 0 {
					goto l1
				}
				goto l59
			case 3:
				t173 := int64(load64(m.memory[uint32(v15):]))
				t174 := int64(m.memory[uint32(v15+i32(8))])
				if t173^i64(7883960361013240164)|(t174^i64(101)) == 0 {
					goto l1
				}
				goto l59
			case 0:
				t175 := int32(load32(m.memory[uint32(v15):]))
				t176 := t175 ^ i32(1953460070)
				v10 = v15 + i32(4)
				t177 := int32(load16(m.memory[uint32(v10):]))
				if t176|(t177^i32(29285)) == 0 {
					goto l1
				}
				t178 := int32(load32(m.memory[uint32(v15):]))
				t179 := int32(load16(m.memory[uint32(v10):]))
				if t178^i32(1684104552)|(t179^i32(29285)) == 0 {
					goto l1
				}
			}
		}
	l59:
		v16 = i32(0)
		store32(m.memory[int64(uint32(v6))+80:], uint32(i32(0)))
		store64(m.memory[int64(uint32(v6))+72:], uint64(i64(0x800000000)))
		t180 := int32(load32(m.memory[uint32(v1+i32(28)):]))
		v9 = t180
		t181 := int32(load32(m.memory[uint32(v1+i32(32)):]))
		v13 = v9 + t181*i32(44)
		v18 = i32(8)
	l63:
		{
			v10 = v9
			if v10 == v13 {
				goto l62
			}
			v9 = v10 + i32(44)
			t182 := int32(load32(m.memory[uint32(v10):]))
			if t182 == i32(-1) {
				goto l63
			}
			{
				t183 := int32(load32(m.memory[int64(uint32(v10))+8:]))
				v19 = t183
				if v19 != i32(8) {
					if v19 != i32(5) {
						goto l63
					}
					{
						{
							t200 := int32(load32(m.memory[int64(uint32(v10))+4:]))
							v19 = t200
							t201 := int32(load32(m.memory[uint32(v19):]))
							t202 := t201 ^ i32(1818386804)
							v12 = v19 + i32(4)
							t203 := int32(m.memory[uint32(v12)])
							if t202|(t203^i32(101)) != 0 {
								goto l67
							}
							t204 := int32(load32(m.memory[int64(uint32(v10))+36:]))
							v11 = t204
							if v11 == 0 {
								goto l67
							}
							t205 := int32(load32(m.memory[int64(uint32(v10))+40:]))
							if t205 != i32(47) {
								goto l67
							}
							t206 := int64(load64(m.memory[int64(uint32(v11))+8:]))
							t207 := int64(load64(m.memory[uint32(v11+i32(16)):]))
							t208 := int64(load64(m.memory[uint32(v11+i32(24)):]))
							t209 := int64(load64(m.memory[uint32(v11+i32(32)):]))
							t210 := int64(load64(m.memory[uint32(v11+i32(40)):]))
							t211 := int64(load64(m.memory[uint32(v11+i32(47)):]))
							if t206^i64(7598524126653739637)|(t207^i64(4211821596982000243))|(t208^i64(7236833184807805812)|(t209^i64(4212112933405418351)))|(t210^i64(7022301986425695608)|(t211^i64(3471766489628697185))) == 0 {
								m.fn425(v6+i32(16), v10, v2)
								t239 := int32(load32(m.memory[int64(uint32(v6))+28:]))
								v10 = t239
								t240 := int32(load32(m.memory[int64(uint32(v6))+24:]))
								v12 = t240
								t241 := int32(load32(m.memory[int64(uint32(v6))+20:]))
								v19 = t241
								{
									t242 := int32(load32(m.memory[int64(uint32(v6))+16:]))
									v11 = t242
									if v11 == i32(-1) {
										{
											t244 := int32(load32(m.memory[int64(uint32(v6))+72:]))
											if uint32(v10) <= uint32(t244-v16) {
												goto l75
											}
											m.fn197(v6+i32(72), v16, v10, i32(8), i32(32))
											t245 := int32(load32(m.memory[int64(uint32(v6))+76:]))
											v18 = t245
											t246 := int32(load32(m.memory[int64(uint32(v6))+80:]))
											v16 = t246
											goto l76
										}
									l75:
										if v10 == 0 {
											goto l77
										}
									l76:
										v11 = v10 << 5
										if v11 == 0 {
											goto l77
										}
										memory_copy(m.memory, uint32(v18+v16<<5), uint32(v12), uint32(v11))
									l77:
										t247 := v6
										v16 = v16 + v10
										store32(m.memory[int64(uint32(t247))+80:], uint32(v16))
										if v19 == 0 {
											goto l63
										}
										m.fn21(v12, v19<<5, i32(8))
										goto l63
									}
									t243 := int64(load64(m.memory[int64(uint32(v6))+32:]))
									store64(m.memory[int64(uint32(v0))+16:], uint64(t243))
									store32(m.memory[int64(uint32(v0))+12:], uint32(v10))
									store32(m.memory[int64(uint32(v0))+8:], uint32(v12))
									store32(m.memory[int64(uint32(v0))+4:], uint32(v19))
									store32(m.memory[uint32(v0):], uint32(v11))
									goto l66
								}
							}
						}
					l67:
						t212 := int32(load32(m.memory[uint32(v19):]))
						t213 := int32(m.memory[uint32(v12)])
						if t212^i32(1734438249)|(t213^i32(101)) != 0 {
							goto l63
						}
						t214 := int32(load32(m.memory[int64(uint32(v10))+36:]))
						v19 = t214
						if v19 == 0 {
							goto l63
						}
						t215 := int32(load32(m.memory[int64(uint32(v10))+40:]))
						if t215 != i32(49) {
							goto l63
						}
						t216 := int64(load64(m.memory[int64(uint32(v19))+8:]))
						t217 := int64(load64(m.memory[uint32(v19+i32(16)):]))
						t218 := int64(load64(m.memory[uint32(v19+i32(24)):]))
						t219 := int64(load64(m.memory[uint32(v19+i32(32)):]))
						t220 := int64(load64(m.memory[uint32(v19+i32(40)):]))
						t221 := int64(load64(m.memory[uint32(v19+i32(48)):]))
						t222 := int64(m.memory[uint32(v19+i32(56))])
						if !(t216^i64(7598524126653739637)|(t217^i64(4211821596982000243))|(t218^i64(7236833184807805812)|(t219^i64(4212112933405418351)))|(t220^i64(8242777485443100024)|(t221^i64(3328505815511955297))|(t222^i64(48))) == 0) {
							goto l63
						}
						store32(m.memory[int64(uint32(v6))+92:], uint32(i32(0)))
						store64(m.memory[int64(uint32(v6))+84:], uint64(i64(0x400000000)))
						store32(m.memory[int64(uint32(v6))+56:], uint32(i32(0)))
						store64(m.memory[int64(uint32(v6))+48:], uint64(i64(0x800000000)))
						m.fn717(v6+i32(16), v1, v2, v6+i32(84), v6+i32(48))
						{
							t223 := int32(load32(m.memory[int64(uint32(v6))+16:]))
							if t223 == i32(-1) {
								t227 := int32(load32(m.memory[int64(uint32(v6))+92:]))
								v1 = t227 * i32(28)
								t228 := int32(load32(m.memory[int64(uint32(v6))+88:]))
								v9 = t228 + i32(-28)
								{
								l71:
									{
										v10 = v1
										if v10 == 0 {
											goto l70
										}
										v1 = v10 + i32(-28)
										v9 = v9 + i32(28)
										t229 := m.fn306(v9)
										if t229 != 0 {
											goto l71
										}
									}
									t230 := int32(load32(m.memory[int64(uint32(v6))+92:]))
									store32(m.memory[int64(uint32(v6))+24:], uint32(t230))
									t231 := int64(load64(m.memory[int64(uint32(v6))+84:]))
									store64(m.memory[int64(uint32(v6))+16:], uint64(t231))
									{
										t232 := int32(load32(m.memory[int64(uint32(v6))+72:]))
										if v16 != t232 {
											goto l72
										}
										m.fn310(v6 + i32(72))
										t233 := int32(load32(m.memory[int64(uint32(v6))+76:]))
										v18 = t233
									}
								l72:
									v1 = v18 + v16<<5
									store32(m.memory[uint32(v1):], uint32(i32(-0x80000000)))
									t234 := int64(load64(m.memory[int64(uint32(v6))+16:]))
									store64(m.memory[int64(uint32(v1))+4:], uint64(t234))
									t235 := int32(load32(m.memory[int64(uint32(v6))+24:]))
									store32(m.memory[int64(uint32(v1))+12:], uint32(t235))
									store32(m.memory[int64(uint32(v6))+80:], uint32(v16+i32(1)))
								}
							l70:
								t236 := int32(load32(m.memory[int64(uint32(v6))+56:]))
								v9 = t236
								t237 := int32(load32(m.memory[int64(uint32(v6))+52:]))
								v1 = t237
								t238 := int32(load32(m.memory[int64(uint32(v6))+48:]))
								store32(m.memory[int64(uint32(v6))+24:], uint32(t238))
								store32(m.memory[int64(uint32(v6))+16:], uint32(v1))
								store32(m.memory[int64(uint32(v6))+20:], uint32(v1))
								store32(m.memory[int64(uint32(v6))+28:], uint32(v1+v9<<5))
								m.fn450(v6+i32(72), v6+i32(16))
								if v10 == 0 {
									goto l73
								}
								goto l62
							}
							t224 := int64(load64(m.memory[int64(uint32(v6))+32:]))
							store64(m.memory[int64(uint32(v0))+16:], uint64(t224))
							t225 := int64(load64(m.memory[int64(uint32(v6))+24:]))
							store64(m.memory[int64(uint32(v0))+8:], uint64(t225))
							t226 := int64(load64(m.memory[int64(uint32(v6))+16:]))
							store64(m.memory[uint32(v0):], uint64(t226))
							m.fn376(v6 + i32(48))
							m.fn718(v6 + i32(84))
							goto l66
						}
					}
				}
				t184 := int32(load32(m.memory[int64(uint32(v10))+4:]))
				t185 := int64(load64(m.memory[uint32(t184):]))
				if t185 != i64(8678262954333332852) {
					goto l63
				}
				t186 := int32(load32(m.memory[int64(uint32(v10))+36:]))
				v19 = t186
				if v19 == 0 {
					goto l63
				}
				t187 := int32(load32(m.memory[int64(uint32(v10))+40:]))
				if t187 != i32(49) {
					goto l63
				}
				t188 := int64(load64(m.memory[int64(uint32(v19))+8:]))
				t189 := int64(load64(m.memory[uint32(v19+i32(16)):]))
				t190 := int64(load64(m.memory[uint32(v19+i32(24)):]))
				t191 := int64(load64(m.memory[uint32(v19+i32(32)):]))
				t192 := int64(load64(m.memory[uint32(v19+i32(40)):]))
				t193 := int64(load64(m.memory[uint32(v19+i32(48)):]))
				t194 := int64(m.memory[uint32(v19+i32(56))])
				if !(t188^i64(7598524126653739637)|(t189^i64(4211821596982000243))|(t190^i64(7236833184807805812)|(t191^i64(4212112933405418351)))|(t192^i64(8242777485443100024)|(t193^i64(3328505815511955297))|(t194^i64(48))) == 0) {
					goto l63
				}
				m.fn346(v6+i32(16), v10, v2)
				t195 := int32(load32(m.memory[int64(uint32(v6))+28:]))
				v10 = t195
				t196 := int32(load32(m.memory[int64(uint32(v6))+24:]))
				v12 = t196
				t197 := int32(load32(m.memory[int64(uint32(v6))+20:]))
				v19 = t197
				t198 := int32(load32(m.memory[int64(uint32(v6))+16:]))
				v11 = t198
				if v11 == i32(-1) {
					{
						t248 := int32(load32(m.memory[int64(uint32(v6))+72:]))
						if uint32(v10) <= uint32(t248-v16) {
							goto l78
						}
						m.fn197(v6+i32(72), v16, v10, i32(8), i32(32))
						t249 := int32(load32(m.memory[int64(uint32(v6))+76:]))
						v18 = t249
						t250 := int32(load32(m.memory[int64(uint32(v6))+80:]))
						v16 = t250
						goto l79
					}
				l78:
					if v10 == 0 {
						goto l80
					}
				l79:
					v11 = v10 << 5
					if v11 == 0 {
						goto l80
					}
					memory_copy(m.memory, uint32(v18+v16<<5), uint32(v12), uint32(v11))
				l80:
					t251 := v6
					v16 = v16 + v10
					store32(m.memory[int64(uint32(t251))+80:], uint32(v16))
					if v19 == 0 {
						goto l63
					}
					m.fn21(v12, v19<<5, i32(8))
					goto l63
				}
				t199 := int64(load64(m.memory[int64(uint32(v6))+32:]))
				store64(m.memory[int64(uint32(v0))+16:], uint64(t199))
				store32(m.memory[int64(uint32(v0))+12:], uint32(v10))
				store32(m.memory[int64(uint32(v0))+8:], uint32(v12))
				store32(m.memory[int64(uint32(v0))+4:], uint32(v19))
				store32(m.memory[uint32(v0):], uint32(v11))
				goto l66
			}
		l73:
		}
		m.fn718(v6 + i32(84))
	l62:
		{
			{
				if v14 != i32(5) {
					goto l81
				}
				t252 := int32(load32(m.memory[uint32(v15):]))
				t253 := int32(m.memory[uint32(v15+i32(4))])
				if t252^i32(1819568500)|(t253^i32(101)) == 0 {
					m.fn720(v6+i32(72), v3)
					goto l1
				}
			}
		l81:
			t254 := int32(load32(m.memory[int64(uint32(v6))+76:]))
			t255 := int32(load32(m.memory[int64(uint32(v6))+80:]))
			m.fn719(v4, t254, t255)
			store32(m.memory[int64(uint32(v6))+80:], uint32(i32(0)))
			m.fn376(v6 + i32(72))
			goto l1
		}
	l66:
	}
	m.fn376(v6 + i32(72))
l10:
	m.g0 = v6 + i32(96)
}
func (m *Module) fn427(v0, v1, v2 int32) {
	var v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14, v15 int32
	t0 := m.g0
	v3 = t0 - i32(32)
	m.g0 = v3
	store32(m.memory[int64(uint32(v3))+16:], uint32(i32(0)))
	store64(m.memory[int64(uint32(v3))+8:], uint64(i64(0x400000000)))
	v4 = i32(4)
	t1 := int32(m.memory[uint32(v1)])
	var p2 int32
	if t1&i32(255) == i32(47) {
		p2 = 1
	}
	v5 = p2
	v6 = i32(1)
	v7 = i32(0)
	{
	l34:
		{
			{
				if v5 != 0 {
					v10 = v6
				l14:
					v11 = v1
					if v10&i32(255) == i32(1) {
						if v2 == 0 {
							m.fn121(i32(1), i32(0), i32(0), i32(1275324))
							panic("unreachable")
						}
						v1 = v11 + i32(1)
						v2 = v2 + i32(-1)
						v13 = i32(6)
						v6 = i32(2)
						goto l22
					}
					if v2 == 0 {
						goto l3
					}
					v12 = i32(0)
				l6:
					{
						v8 = i32(1)
						{
							t3 := int32(m.memory[uint32(v11+v12)])
							if t3 != i32(47) {
								goto l4
							}
							v10 = i32(1)
							goto l5
						}
					l4:
						t4 := v2
						v12 = v12 + i32(1)
						if t4 != v12 {
							goto l6
						}
					}
					v10 = i32(0)
					v12 = v2
				l5:
					v13 = i32(255)
					switch v12 {
					case 0:
						goto l7
					default:
						goto l10
					case 2:
						t5 := int32(m.memory[uint32(v11)])
						if t5 != i32(46) {
							goto l10
						}
						t6 := int32(m.memory[int64(uint32(v11))+1])
						if t6&i32(255) != i32(46) {
							goto l10
						}
						v13 = i32(8)
						goto l11
					case 1:
						t7 := int32(m.memory[uint32(v11)])
						if t7 == i32(46) {
							goto l7
						}
					}
				l10:
					v13 = i32(9)
				l11:
					v8 = i32(0)
				l7:
					{
						t8 := v2
						v14 = v10 + v12
						if uint32(t8) >= uint32(v14) {
							v1 = v11 + v14
							v2 = v2 - v14
							v10 = i32(2)
							v15 = v11
							if v8 != 0 {
								goto l14
							}
							v15 = v11
							goto l15
						}
						v8 = v2
						goto l13
					}
				}
				v8 = v2
				v9 = v6
			l31:
				v11 = v1
				if v9&i32(255) != i32(1) {
					if v8 == 0 {
						goto l3
					}
					v12 = i32(0)
				l25:
					{
						v10 = i32(1)
						{
							t12 := int32(m.memory[uint32(v11+v12)])
							if t12 != i32(47) {
								goto l23
							}
							v1 = i32(1)
							goto l24
						}
					l23:
						t13 := v8
						v12 = v12 + i32(1)
						if t13 != v12 {
							goto l25
						}
					}
					v1 = i32(0)
					v12 = v8
				l24:
					v13 = i32(255)
					switch v12 {
					case 0:
						goto l26
					default:
						goto l29
					case 1:
						t14 := int32(m.memory[uint32(v11)])
						if t14 == i32(46) {
							goto l26
						}
						goto l29
					case 2:
						t15 := int32(m.memory[uint32(v11)])
						if t15 != i32(46) {
							goto l29
						}
						t16 := int32(m.memory[int64(uint32(v11))+1])
						if t16&i32(255) != i32(46) {
							goto l29
						}
						v13 = i32(8)
						goto l30
					}
				l29:
					v13 = i32(9)
				l30:
					v10 = i32(0)
				l26:
					t17 := v8
					v14 = v1 + v12
					if uint32(t17) < uint32(v14) {
						goto l13
					}
					v1 = v11 + v14
					v9 = i32(2)
					v15 = v11
					v2 = v8 - v14
					v8 = v2
					if v10 != 0 {
						goto l31
					}
					v15 = v11
					goto l15
				}
				v6 = i32(2)
				v1 = v11
				switch v8 {
				case 0:
					goto l17
				default:
					v1 = v11
					t9 := int32(m.memory[uint32(v11)])
					if t9 != i32(46) {
						goto l17
					}
					v1 = v11
					t10 := int32(m.memory[int64(uint32(v11))+1])
					if t10 == i32(47) {
						goto l20
					}
					goto l17
				case 1:
					v8 = i32(1)
					v1 = v11
					t11 := int32(m.memory[uint32(v11)])
					if t11 != i32(46) {
						goto l17
					}
				}
			l20:
				if v2 == 0 {
					m.fn121(i32(1), i32(0), i32(0), i32(1275308))
					panic("unreachable")
				}
				v1 = v11 + i32(1)
				v2 = v2 + i32(-1)
				v13 = i32(7)
				v6 = i32(2)
				goto l22
			l17:
				v9 = i32(2)
				goto l31
			l15:
				if v13 != i32(255) {
					goto l22
				}
			l3:
				t18 := int32(load32(m.memory[int64(uint32(v3))+16:]))
				store32(m.memory[int64(uint32(v0))+8:], uint32(t18))
				t19 := int64(load64(m.memory[int64(uint32(v3))+8:]))
				store64(m.memory[uint32(v0):], uint64(t19))
				goto l32
			}
		l22:
			v11 = v13 + i32(-5)
			p20 := v11
			if uint32(v11) > uint32(v13) {
				p20 = i32(0)
			}
			v11 = p20
			if v11 == i32(2) {
				goto l34
			}
			switch v11 {
			case 2:
				panic("unreachable")
			case 1:
				v7 = i32(0)
				store32(m.memory[int64(uint32(v3))+16:], uint32(i32(0)))
				goto l34
			default:
				m.fn263(v0+i32(4), i32(20), i32(1093360), i32(35))
				goto l40
			case 3:
				if v7 == 0 {
					goto l41
				}
				t21 := v3
				v7 = v7 + i32(-1)
				store32(m.memory[int64(uint32(t21))+16:], uint32(v7))
				goto l34
			case 4:
				m.fn14(v3+i32(20), v15, v12)
				{
					t22 := int32(load32(m.memory[int64(uint32(v3))+20:]))
					if t22 != 0 {
						goto l42
					}
					t23 := int32(load32(m.memory[int64(uint32(v3))+28:]))
					v11 = t23
					t24 := int32(load32(m.memory[int64(uint32(v3))+24:]))
					v8 = t24
					{
						t25 := int32(load32(m.memory[int64(uint32(v3))+8:]))
						if v7 != t25 {
							goto l43
						}
						m.fn743(v3 + i32(8))
						t26 := int32(load32(m.memory[int64(uint32(v3))+12:]))
						v4 = t26
					}
				l43:
					v10 = v4 + v7<<3
					store32(m.memory[int64(uint32(v10))+4:], uint32(v11))
					store32(m.memory[uint32(v10):], uint32(v8))
					t27 := v3
					v7 = v7 + i32(1)
					store32(m.memory[int64(uint32(t27))+16:], uint32(v7))
					goto l34
				}
			l42:
			}
		}
		m.fn263(v0+i32(4), i32(20), i32(1093429), i32(14))
		goto l40
	l41:
		m.fn263(v0+i32(4), i32(20), i32(1093395), i32(34))
	l40:
		store32(m.memory[uint32(v0):], uint32(i32(-1)))
		t28 := int32(load32(m.memory[int64(uint32(v3))+8:]))
		v12 = t28
		if v12 == 0 {
			goto l32
		}
		{
			t29 := int32(load32(m.memory[int64(uint32(v3))+12:]))
			v8 = t29
			t30 := int32(load32(m.memory[uint32(v8+i32(-4)):]))
			v11 = t30
			v10 = v11 & i32(-8)
			t31 := v10
			v11 = v11 & i32(3)
			p32 := i32(8)
			if v11 != 0 {
				p32 = i32(4)
			}
			v12 = v12 << 3
			if uint32(t31) < uint32(p32+v12) {
				m.fn7(i32(1274404), i32(46), i32(1274452))
				panic("unreachable")
			}
			if v11 == 0 {
				goto l45
			}
			if uint32(v10) > uint32(v12+i32(39)) {
				m.fn7(i32(1274468), i32(46), i32(1274516))
				panic("unreachable")
			}
		l45:
			m.fn5(v8)
			goto l32
		}
	}
l32:
	m.g0 = v3 + i32(32)
	return
l13:
	m.fn121(v14, v8, v8, i32(1275340))
	panic("unreachable")
}
func (m *Module) fn428(v0, v1, v2, v3 int32) {
	var v4, v5, v6, v7, v8, v9, v10 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		v4 = t0
		if v4 <= i32(-1) {
			panic("unreachable")
		}
		{
			v5 = v4 + i32(1)
			if v5 < v4 {
				m.fn140(i32(1274532), i32(28), i32(1274560))
				panic("unreachable")
			}
			store32(m.memory[int64(uint32(v1))+8:], uint32(v5))
			{
				if v3 != 0 {
					goto l2
				}
				v6 = i32(0)
				v7 = i32(1)
				goto l3
			l2:
				v8 = v2 + v3<<3
				t1 := int32(load32(m.memory[uint32(v1+i32(100)):]))
				v5 = t1
				t2 := int32(load32(m.memory[uint32(v1+i32(96)):]))
				v9 = t2
				v6 = i32(0)
			l11:
				{
					if uint32(v6) >= uint32(v5) {
						m.fn33(v6, v5, i32(1069844))
						panic("unreachable")
					}
					v7 = i32(0)
					t3 := int32(load32(m.memory[int64(uint32(v9+v6*i32(80)))+48:]))
					v4 = t3
					if v4 == i32(-1) {
						goto l3
					}
					v10 = v2 + i32(8)
				l10:
					{
						if uint32(v4) >= uint32(v5) {
							m.fn33(v4, v5, i32(1069844))
							panic("unreachable")
						}
						{
							t4 := int32(load32(m.memory[uint32(v2):]))
							t5 := int32(load32(m.memory[int64(uint32(v2))+4:]))
							v3 = v9 + v4*i32(80)
							t6 := int32(load32(m.memory[int64(uint32(v3))+64:]))
							t7 := int32(load32(m.memory[int64(uint32(v3))+68:]))
							t8 := m.fn291(t4, t5, t6, t7)
							switch t8 & i32(255) {
							case 0:
								goto l6
							case 1:
								goto l7
							default:
								v4 = i32(40)
								goto l9
							}
						}
					l7:
						v4 = i32(44)
					l9:
						t9 := int32(load32(m.memory[uint32(v3+v4):]))
						v4 = t9
						if v4 != i32(-1) {
							goto l10
						}
						goto l3
					}
				}
			l6:
				v2 = v10
				v6 = v4
				if v10 != v8 {
					goto l11
				}
				v6 = v4
				v7 = i32(1)
			}
		l3:
			t10 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			t11 := v1
			v4 = t10
			store32(m.memory[int64(uint32(t11))+8:], uint32(v4+i32(-1)))
			if v4 <= i32(0) {
				m.fn28(i32(1275252), i32(77), i32(1275292))
				panic("unreachable")
			}
			store32(m.memory[int64(uint32(v0))+4:], uint32(v6))
			store32(m.memory[uint32(v0):], uint32(v7))
			return
		}
	}
}
func (m *Module) fn429(v0, v1, v2 int32) {
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
				t1 := int64(load64(m.memory[int64(uint32(v1))+64:]))
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
					m.fn432(t3, t4, t5, int32(p2))
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
						t9 := int64(load64(m.memory[int64(uint32(v1))+64:]))
						v4 = t9
						t10 := int32(load32(m.memory[int64(uint32(v3))+44:]))
						t11 := v4
						v5 = t10
						v7 = int64(uint32(v5))
						if uint64(t11) < uint64(v7) {
							m.fn28(i32(1080592), i32(69), i32(1080628))
							panic("unreachable")
						}
						store64(m.memory[int64(uint32(v1))+64:], uint64(v4-v7))
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
						m.fn121(i32(0), v5, i32(32), i32(1070132))
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
								m.fn7(i32(1274404), i32(46), i32(1274452))
								panic("unreachable")
							}
							if v10 == 0 {
								goto l14
							}
							if uint32(v11) > uint32(v9+i32(39)) {
								m.fn7(i32(1274468), i32(46), i32(1274516))
								panic("unreachable")
							}
						l14:
							m.fn5(v8)
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
							m.fn7(i32(1274404), i32(46), i32(1274452))
							panic("unreachable")
						}
						if v8 == 0 {
							goto l17
						}
						if uint32(v9) >= uint32(i32(52)) {
							m.fn7(i32(1274468), i32(46), i32(1274516))
							panic("unreachable")
						}
					l17:
						m.fn5(v5)
					}
				l9:
					t24 := int64(load64(m.memory[int64(uint32(v1))+64:]))
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
func (m *Module) fn430(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6 int32
	var v7 int64
	var v8, v9, v10, v11, v12, v13, v14, v15, v16 int32
	var v17 int64
	t0 := m.g0
	v2 = t0 - i32(64)
	m.g0 = v2
	t1 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	v3 = t1
	t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	v4 = t2
	v0 = i32(1)
	{
		t3 := int32(load32(m.memory[uint32(v1):]))
		v5 = t3
		t4 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t5 := v5
		v6 = t4
		t6 := int32(load32(m.memory[int64(uint32(v6))+16:]))
		v1 = t6
		t7 := m.t0[uint(v1)].(func(int32, int32) int32)(t5, i32(34))
		if t7 != 0 {
			goto l0
		}
		store32(m.memory[int64(uint32(v2))+12:], uint32(v3))
		store32(m.memory[int64(uint32(v2))+8:], uint32(v4))
		v7 = int64(uint32(i32(56)))<<32 | int64(uint32(v2+i32(32)))
	l27:
		{
			m.fn30(v2+i32(16), v2+i32(8))
			{
				{
					t8 := int32(load32(m.memory[int64(uint32(v2))+16:]))
					v8 = t8
					if v8 == 0 {
						t32 := m.t0[uint(v1)].(func(int32, int32) int32)(v5, i32(34))
						v0 = t32
						goto l0
					}
					t9 := int32(load32(m.memory[int64(uint32(v2))+28:]))
					v9 = t9
					t10 := int32(load32(m.memory[int64(uint32(v2))+24:]))
					v4 = t10
					{
						t11 := int32(load32(m.memory[int64(uint32(v2))+20:]))
						v10 = t11
						if v10 != 0 {
							v11 = v8 + v10
							v0 = i32(0)
							v12 = v8
							v13 = i32(0)
						l23:
							{
								{
									t12 := int32(int8(m.memory[uint32(v12)]))
									v3 = t12
									if v3 <= i32(-1) {
										goto l4
									}
									v14 = v12 + i32(1)
									v15 = v3 & i32(255)
									goto l5
								}
							l4:
								t13 := int32(m.memory[int64(uint32(v12))+1])
								v16 = t13 & i32(63)
								v14 = v3 & i32(31)
								if uint32(v3) > uint32(i32(-33)) {
									goto l6
								}
								v15 = v14<<6 | v16
								v14 = v12 + i32(2)
								goto l5
							l6:
								t14 := int32(m.memory[int64(uint32(v12))+2])
								v16 = v16<<6 | t14&i32(63)
								if uint32(v3) >= uint32(i32(-16)) {
									goto l7
								}
								v15 = v16 | v14<<12
								v14 = v12 + i32(3)
								goto l5
							l7:
								t15 := int32(m.memory[int64(uint32(v12))+3])
								v15 = v16<<6 | t15&i32(63) | v14<<18&i32(0x1c0000)
								v14 = v12 + i32(4)
							}
						l5:
							m.fn856(v2+i32(32), v15, i32(65537))
							{
								t16 := int32(m.memory[int64(uint32(v2))+45])
								t17 := int32(m.memory[int64(uint32(v2))+44])
								if (t16-t17)&i32(255) == i32(1) {
									goto l8
								}
								{
									if uint32(v13) < uint32(v0) {
										goto l9
									}
									{
										if v0 == 0 {
											goto l10
										}
										if uint32(v0) < uint32(v10) {
											goto l11
										}
										if v0 != v10 {
											goto l9
										}
										goto l10
									l11:
										t18 := int32(int8(m.memory[uint32(v8+v0)]))
										if t18 <= i32(-65) {
											goto l9
										}
									}
								l10:
									{
										if v13 == 0 {
											goto l12
										}
										if uint32(v13) < uint32(v10) {
											goto l13
										}
										if v13 == v10 {
											goto l12
										}
										goto l9
									l13:
										t19 := int32(int8(m.memory[uint32(v8+v13)]))
										if t19 <= i32(-65) {
											goto l9
										}
									}
								l12:
									t20 := int32(load32(m.memory[int64(uint32(v6))+12:]))
									t21 := m.t0[uint(t20)].(func(int32, int32, int32) int32)(v5, v8+v0, v13-v0)
									if t21 != 0 {
										goto l14
									}
									t22 := int64(load64(m.memory[int64(uint32(v2))+40:]))
									store64(m.memory[int64(uint32(v2))+56:], uint64(t22))
									t23 := int64(load64(m.memory[int64(uint32(v2))+32:]))
									t24 := v2
									v17 = t23
									store64(m.memory[int64(uint32(t24))+48:], uint64(v17))
									t25 := int32(m.memory[int64(uint32(v2))+60])
									v0 = t25
									t26 := int32(m.memory[int64(uint32(v2))+61])
									v3 = t26
									if uint32(v3) < uint32(i32(129)) {
										p28 := v3
										if uint32(v0) > uint32(v3) {
											p28 = v0
										}
										v16 = p28
									l18:
										{
											if v16 == v0 {
												goto l16
											}
											v3 = v2 + i32(48) + v0
											v0 = v0 + i32(1)
											t29 := int32(m.memory[uint32(v3)])
											t30 := m.t0[uint(v1)].(func(int32, int32) int32)(v5, t29)
											if t30 == 0 {
												goto l18
											}
											goto l14
										}
									}
									v16 = int32(v17)
								l17:
									{
										if uint32(v0&i32(255)) >= uint32(v3) {
											goto l16
										}
										v0 = v0 + i32(1)
										t27 := m.t0[uint(v1)].(func(int32, int32) int32)(v5, v16)
										if t27 == 0 {
											goto l17
										}
										goto l14
									}
								}
							l9:
								m.fn38(v8, v10, v0, v13, i32(1122556))
								panic("unreachable")
							l16:
								{
									if uint32(v15) >= uint32(i32(128)) {
										goto l19
									}
									v0 = i32(1)
									goto l20
								l19:
									if uint32(v15) >= uint32(i32(2048)) {
										goto l21
									}
									v0 = i32(2)
									goto l20
								l21:
									p31 := i32(4)
									if uint32(v15) < uint32(i32(65536)) {
										p31 = i32(3)
									}
									v0 = p31
								}
							l20:
								v0 = v0 + v13
							}
						l8:
							v13 = v13 - v12 + v14
							v12 = v14
							if v14 == v11 {
								goto l22
							}
							goto l23
						}
						v3 = i32(0)
						goto l3
					}
				}
			l22:
				if v0 != 0 {
					goto l24
				}
				v3 = i32(0)
				goto l3
			l24:
				if uint32(v0) < uint32(v10) {
					goto l25
				}
				v3 = v10
				if v0 == v10 {
					goto l3
				}
				goto l26
			l25:
				t33 := int32(int8(m.memory[uint32(v8+v0)]))
				if t33 < i32(-64) {
					goto l26
				}
				v3 = v0
			}
		l3:
			t34 := int32(load32(m.memory[int64(uint32(v6))+12:]))
			t35 := m.t0[uint(t34)].(func(int32, int32, int32) int32)(v5, v8+v3, v10-v3)
			if t35 != 0 {
				goto l14
			}
			if v9 == 0 {
				goto l27
			}
		l28:
			{
				t36 := int32(m.memory[uint32(v4)])
				m.memory[int64(uint32(v2))+32] = byte(t36)
				store64(m.memory[int64(uint32(v2))+48:], uint64(v7))
				t37 := m.fn46(v5, v6, i32(1122544), v2+i32(48))
				if t37 != 0 {
					goto l14
				}
				v4 = v4 + i32(1)
				v9 = v9 + i32(-1)
				if v9 == 0 {
					goto l27
				}
				goto l28
			}
		}
	l14:
		v0 = i32(1)
	}
l0:
	m.g0 = v2 + i32(64)
	return v0
l26:
	m.fn38(v8, v10, v0, v10, i32(1122528))
	panic("unreachable")
}
func (m *Module) fn431(v0 int32) {
	var v1, v2, v3, v4, v5, v6, v7 int32
	t0 := m.g0
	v1 = t0 - i32(16)
	m.g0 = v1
	t1 := int32(load32(m.memory[int64(uint32(v0))+40:]))
	v2 = t1
	store32(m.memory[int64(uint32(v0))+40:], uint32(i32(0)))
	{
		if v2 == 0 {
			goto l0
		}
		t2 := int32(load32(m.memory[int64(uint32(v0))+44:]))
		t3 := v1 + i32(8)
		t4 := v2
		t5 := v0
		v3 = t2
		t6 := int32(load32(m.memory[int64(uint32(v3))+12:]))
		m.t0[uint(t6)].(func(int32, int32, int32))(t3, t4, t5)
		{
			t7 := int32(m.memory[int64(uint32(v1))+8])
			if t7 == i32(255) {
				{
					t25 := int32(load32(m.memory[uint32(v3):]))
					v4 = t25
					if v4 == 0 {
						goto l15
					}
					m.t0[uint(v4)].(func(int32))(v2)
				}
			l15:
				t26 := int32(load32(m.memory[int64(uint32(v3))+4:]))
				v3 = t26
				if v3 == 0 {
					goto l0
				}
				t27 := int32(load32(m.memory[uint32(v2+i32(-4)):]))
				v4 = t27
				v5 = v4 & i32(-8)
				t28 := v5
				v4 = v4 & i32(3)
				p29 := i32(8)
				if v4 != 0 {
					p29 = i32(4)
				}
				if uint32(t28) < uint32(p29+v3) {
					m.fn7(i32(1274404), i32(46), i32(1274452))
					panic("unreachable")
				}
				if v4 == 0 {
					goto l14
				}
				if uint32(v5) <= uint32(v3+i32(39)) {
					goto l14
				}
				m.fn7(i32(1274468), i32(46), i32(1274516))
				panic("unreachable")
			}
			t8 := int32(load32(m.memory[int64(uint32(v1))+12:]))
			v4 = t8
			t9 := int32(m.memory[int64(uint32(v1))+8])
			v5 = t9
			{
				t10 := int32(load32(m.memory[uint32(v3):]))
				v6 = t10
				if v6 == 0 {
					goto l2
				}
				m.t0[uint(v6)].(func(int32))(v2)
			}
		l2:
			{
				t11 := int32(load32(m.memory[int64(uint32(v3))+4:]))
				v3 = t11
				if v3 == 0 {
					goto l3
				}
				t12 := int32(load32(m.memory[uint32(v2+i32(-4)):]))
				v6 = t12
				v7 = v6 & i32(-8)
				t13 := v7
				v6 = v6 & i32(3)
				p14 := i32(8)
				if v6 != 0 {
					p14 = i32(4)
				}
				if uint32(t13) < uint32(p14+v3) {
					m.fn7(i32(1274404), i32(46), i32(1274452))
					panic("unreachable")
				}
				if v6 == 0 {
					goto l5
				}
				if uint32(v7) > uint32(v3+i32(39)) {
					m.fn7(i32(1274468), i32(46), i32(1274516))
					panic("unreachable")
				}
			l5:
				m.fn5(v2)
			}
		l3:
			if v5 != i32(3) {
				goto l0
			}
			t15 := int32(load32(m.memory[uint32(v4):]))
			v2 = t15
			{
				t16 := int32(load32(m.memory[uint32(v4+i32(4)):]))
				v3 = t16
				t17 := int32(load32(m.memory[uint32(v3):]))
				v5 = t17
				if v5 == 0 {
					goto l7
				}
				m.t0[uint(v5)].(func(int32))(v2)
			}
		l7:
			{
				t18 := int32(load32(m.memory[int64(uint32(v3))+4:]))
				v3 = t18
				if v3 == 0 {
					goto l8
				}
				t19 := int32(load32(m.memory[uint32(v2+i32(-4)):]))
				v5 = t19
				v6 = v5 & i32(-8)
				t20 := v6
				v5 = v5 & i32(3)
				p21 := i32(8)
				if v5 != 0 {
					p21 = i32(4)
				}
				if uint32(t20) < uint32(p21+v3) {
					m.fn7(i32(1274404), i32(46), i32(1274452))
					panic("unreachable")
				}
				if v5 == 0 {
					goto l10
				}
				if uint32(v6) > uint32(v3+i32(39)) {
					m.fn7(i32(1274468), i32(46), i32(1274516))
					panic("unreachable")
				}
			l10:
				m.fn5(v2)
			}
		l8:
			t22 := int32(load32(m.memory[uint32(v4+i32(-4)):]))
			v2 = t22
			v3 = v2 & i32(-8)
			t23 := v3
			v2 = v2 & i32(3)
			p24 := i32(20)
			if v2 != 0 {
				p24 = i32(16)
			}
			if uint32(t23) < uint32(p24) {
				m.fn7(i32(1274404), i32(46), i32(1274452))
				panic("unreachable")
			}
			if v2 != 0 {
				goto l13
			}
			v2 = v4
			goto l14
		}
	l13:
		v2 = v4
		if uint32(v3) >= uint32(i32(52)) {
			m.fn7(i32(1274468), i32(46), i32(1274516))
			panic("unreachable")
		}
	l14:
		m.fn5(v2)
	}
l0:
	{
		t30 := int32(load32(m.memory[int64(uint32(v0))+48:]))
		v2 = t30
		if v2 == i32(-1) {
			goto l18
		}
		t31 := int32(load32(m.memory[int64(uint32(v2))+4:]))
		t32 := v2
		v3 = t31
		store32(m.memory[int64(uint32(t32))+4:], uint32(v3+i32(-1)))
		if v3 != i32(1) {
			goto l18
		}
		t33 := int32(load32(m.memory[uint32(v2+i32(-4)):]))
		v3 = t33
		t34 := v3 & i32(-8)
		v4 = v3 & i32(3)
		p35 := i32(144)
		if v4 != 0 {
			p35 = i32(140)
		}
		if uint32(t34) < uint32(p35) {
			m.fn7(i32(1274404), i32(46), i32(1274452))
			panic("unreachable")
		}
		if v4 == 0 {
			goto l20
		}
		if uint32(v3) >= uint32(i32(176)) {
			m.fn7(i32(1274468), i32(46), i32(1274516))
			panic("unreachable")
		}
	l20:
		m.fn5(v2)
	}
l18:
	{
		t36 := int32(load32(m.memory[uint32(v0):]))
		v2 = t36
		if v2 == 0 {
			goto l22
		}
		t37 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		v4 = t37
		t38 := int32(load32(m.memory[uint32(v4+i32(-4)):]))
		v3 = t38
		v5 = v3 & i32(-8)
		t39 := v5
		v3 = v3 & i32(3)
		p40 := i32(8)
		if v3 != 0 {
			p40 = i32(4)
		}
		if uint32(t39) < uint32(p40+v2) {
			m.fn7(i32(1274404), i32(46), i32(1274452))
			panic("unreachable")
		}
		if v3 == 0 {
			goto l24
		}
		if uint32(v5) > uint32(v2+i32(39)) {
			m.fn7(i32(1274468), i32(46), i32(1274516))
			panic("unreachable")
		}
	l24:
		m.fn5(v4)
	}
l22:
	{
		t41 := int32(load32(m.memory[int64(uint32(v0))+40:]))
		v2 = t41
		if v2 == 0 {
			goto l26
		}
		{
			t42 := int32(load32(m.memory[int64(uint32(v0))+44:]))
			v0 = t42
			t43 := int32(load32(m.memory[uint32(v0):]))
			v3 = t43
			if v3 == 0 {
				goto l27
			}
			m.t0[uint(v3)].(func(int32))(v2)
		}
	l27:
		t44 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		v0 = t44
		if v0 == 0 {
			goto l26
		}
		t45 := int32(load32(m.memory[uint32(v2+i32(-4)):]))
		v3 = t45
		v4 = v3 & i32(-8)
		t46 := v4
		v3 = v3 & i32(3)
		p47 := i32(8)
		if v3 != 0 {
			p47 = i32(4)
		}
		if uint32(t46) < uint32(p47+v0) {
			m.fn7(i32(1274404), i32(46), i32(1274452))
			panic("unreachable")
		}
		if v3 == 0 {
			goto l29
		}
		if uint32(v4) > uint32(v0+i32(39)) {
			m.fn7(i32(1274468), i32(46), i32(1274516))
			panic("unreachable")
		}
	l29:
		m.fn5(v2)
	}
l26:
	m.g0 = v1 + i32(16)
}
func (m *Module) fn432(v0, v1, v2, v3 int32) {
	var v4, v5, v6, v7 int32
	var v8 int64
	var v9, v10, v11 int32
	var v12 int64
	var v13, v14 int32
	var v15, v16 int64
	var v17, v18, v19, v20, v21, v22 int32
	var v23 int64
	var v24, v25, v26 int32
	var v27 int64
	t0 := m.g0
	v4 = t0 - i32(128)
	m.g0 = v4
	{
		{
			t1 := int32(load32(m.memory[int64(uint32(v1))+12:]))
			v5 = t1
			t2 := int32(load32(m.memory[int64(uint32(v1))+16:]))
			if uint32(v5) < uint32(t2) {
				goto l0
			}
			t3 := int64(load64(m.memory[int64(uint32(v1))+32:]))
			t4 := int64(load64(m.memory[int64(uint32(v1))+24:]))
			if uint64(t3+int64(uint32(v5))) >= uint64(t4) {
				goto l0
			}
			t5 := int32(load32(m.memory[int64(uint32(v1))+40:]))
			v6 = t5
			store32(m.memory[int64(uint32(v1))+40:], uint32(i32(0)))
			{
				{
					{
						{
							{
								if v6 == 0 {
									goto l1
								}
								t6 := int32(load32(m.memory[int64(uint32(v1))+44:]))
								t7 := v4 + i32(96)
								t8 := v6
								t9 := v1
								v7 = t6
								t10 := int32(load32(m.memory[int64(uint32(v7))+12:]))
								m.t0[uint(t10)].(func(int32, int32, int32))(t7, t8, t9)
								{
									t11 := int32(m.memory[int64(uint32(v4))+96])
									if t11 == i32(255) {
										goto l2
									}
									t12 := int64(m.memory[int64(uint32(v4))+96])
									v8 = t12
									t13 := int32(load32(m.memory[int64(uint32(v4))+100:]))
									v5 = t13
									t14 := int32(load32(m.memory[int64(uint32(v4))+96:]))
									v9 = t14
									{
										t15 := int32(load32(m.memory[uint32(v7):]))
										v10 = t15
										if v10 == 0 {
											goto l3
										}
										m.t0[uint(v10)].(func(int32))(v6)
									}
								l3:
									{
										t16 := int32(load32(m.memory[int64(uint32(v7))+4:]))
										v7 = t16
										if v7 == 0 {
											goto l4
										}
										t17 := int32(load32(m.memory[uint32(v6+i32(-4)):]))
										v10 = t17
										v11 = v10 & i32(-8)
										t18 := v11
										v10 = v10 & i32(3)
										p19 := i32(8)
										if v10 != 0 {
											p19 = i32(4)
										}
										if uint32(t18) < uint32(p19+v7) {
											m.fn7(i32(1274404), i32(46), i32(1274452))
											panic("unreachable")
										}
										if v10 == 0 {
											goto l6
										}
										if uint32(v11) > uint32(v7+i32(39)) {
											m.fn7(i32(1274468), i32(46), i32(1274516))
											panic("unreachable")
										}
									l6:
										m.fn5(v6)
									}
								l4:
									if v8 == i64(255) {
										goto l1
									}
									v8 = int64(uint32(int32(uint32(v9) >> 8)))
									goto l8
								}
							l2:
								{
									t20 := int32(load32(m.memory[uint32(v7):]))
									v5 = t20
									if v5 == 0 {
										goto l9
									}
									m.t0[uint(v5)].(func(int32))(v6)
								}
							l9:
								t21 := int32(load32(m.memory[int64(uint32(v7))+4:]))
								v5 = t21
								if v5 == 0 {
									goto l1
								}
								t22 := int32(load32(m.memory[uint32(v6+i32(-4)):]))
								v7 = t22
								v9 = v7 & i32(-8)
								t23 := v9
								v7 = v7 & i32(3)
								p24 := i32(8)
								if v7 != 0 {
									p24 = i32(4)
								}
								if uint32(t23) < uint32(p24+v5) {
									m.fn7(i32(1274404), i32(46), i32(1274452))
									panic("unreachable")
								}
								if v7 == 0 {
									goto l11
								}
								if uint32(v9) > uint32(v5+i32(39)) {
									m.fn7(i32(1274468), i32(46), i32(1274516))
									panic("unreachable")
								}
							l11:
								m.fn5(v6)
							}
						l1:
							t25 := int64(load64(m.memory[int64(uint32(v1))+32:]))
							t26 := int64(load32(m.memory[int64(uint32(v1))+12:]))
							t27 := v1
							v8 = t25 + t26
							store64(m.memory[int64(uint32(t27))+32:], uint64(v8))
							{
								{
									t28 := int32(load32(m.memory[int64(uint32(v1))+48:]))
									v7 = t28
									if v7 == i32(-1) {
										goto l13
									}
									t29 := int32(load32(m.memory[int64(uint32(v1))+52:]))
									v10 = t29
									t30 := int64(load64(m.memory[int64(uint32(v1))+24:]))
									v12 = t30
									t31 := int32(load32(m.memory[uint32(v7):]))
									v5 = t31
								l15:
									{
										if v5 == 0 {
											goto l13
										}
										if v5 <= i32(-1) {
											store64(m.memory[int64(uint32(v4))+96:], uint64(int64(uint32(i32(1)))<<32|int64(uint32(i32(1075772)))))
											m.fn28(i32(1052645), v4+i32(96), i32(1079688))
											panic("unreachable")
										}
										t32 := int32(load32(m.memory[uint32(v7):]))
										t33 := v7
										t34 := v5 + i32(1)
										v6 = t32
										t35 := v6
										var p36 int32
										if v6 == v5 {
											p36 = 1
										}
										v9 = p36
										p37 := t35
										if v9 != 0 {
											p37 = t34
										}
										store32(m.memory[uint32(t33):], uint32(p37))
										v5 = v6
										if v9 == 0 {
											goto l15
										}
									}
									store32(m.memory[int64(uint32(v1))+12:], uint32(i32(0)))
									v12 = v12 - v8
									t38 := int32(load32(m.memory[int64(uint32(v1))+8:]))
									t39 := v12
									v5 = t38
									if uint64(t39) > uint64(uint32(v5)) {
										goto l16
									}
									v13 = v5
									goto l17
								}
							l13:
								t40 := m.fn11(i32(24))
								v1 = t40
								if v1 == 0 {
									m.fn16(i32(1), i32(24))
									panic("unreachable")
								}
								t41 := int64(load64(m.memory[int64(uint32(i32(0)))+1070823:]))
								store64(m.memory[int64(uint32(v1))+16:], uint64(t41))
								t42 := int64(load64(m.memory[int64(uint32(i32(0)))+1070815:]))
								store64(m.memory[int64(uint32(v1))+8:], uint64(t42))
								t43 := int64(load64(m.memory[int64(uint32(i32(0)))+1070807:]))
								store64(m.memory[uint32(v1):], uint64(t43))
								t44 := m.fn11(i32(12))
								v6 = t44
								if v6 == 0 {
									m.fn23(i32(4), i32(12))
									panic("unreachable")
								}
								store32(m.memory[int64(uint32(v6))+8:], uint32(i32(24)))
								store32(m.memory[int64(uint32(v6))+4:], uint32(v1))
								store32(m.memory[uint32(v6):], uint32(i32(24)))
								t45 := m.fn11(i32(12))
								v5 = t45
								if v5 == 0 {
									m.fn23(i32(4), i32(12))
									panic("unreachable")
								}
								m.memory[int64(uint32(v5))+8] = byte(i32(40))
								store32(m.memory[int64(uint32(v5))+4:], uint32(i32(1070848)))
								store32(m.memory[uint32(v5):], uint32(v6))
								v9 = i32(3)
								v8 = i64(0)
								goto l8
							}
						l16:
							{
								t46 := int32(load32(m.memory[int64(uint32(v1))+20:]))
								v6 = t46
								t47 := v6
								t48 := v6
								v9 = int32(v12)
								p49 := v9
								if uint32(v6) < uint32(v9) {
									p49 = t48
								}
								p50 := p49
								if uint64(v12) > uint64(i64(0xffffffff)) {
									p50 = t47
								}
								v6 = p50
								p51 := i32(1024)
								if uint32(v6) > uint32(i32(1024)) {
									p51 = v6
								}
								v13 = p51
								if uint32(v13) <= uint32(v5) {
									goto l21
								}
								{
									v9 = v13 - v5
									t52 := int32(load32(m.memory[uint32(v1):]))
									if uint32(v9) <= uint32(t52-v5) {
										goto l22
									}
									m.fn715(v1, v5, v9, i32(1), i32(1))
									t53 := int32(load32(m.memory[int64(uint32(v1))+8:]))
									v5 = t53
								}
							l22:
								t54 := int32(load32(m.memory[int64(uint32(v1))+4:]))
								v11 = t54
								v6 = v11 + v5
								{
									if uint32(v9) < uint32(i32(2)) {
										goto l23
									}
									v9 = v9 + i32(-1)
									if v9 == 0 {
										goto l24
									}
									memory_zero(m.memory, uint32(v6), uint32(v9))
								l24:
									t55 := v11
									v5 = v5 + v9
									v6 = t55 + v5
								}
							l23:
								m.memory[uint32(v6)] = byte(i32(0))
								v13 = v5 + i32(1)
							}
						l21:
							store32(m.memory[int64(uint32(v1))+8:], uint32(v13))
						l17:
							t56 := int32(load32(m.memory[int64(uint32(v7))+8:]))
							if t56 != 0 {
								goto l25
							}
							t57 := int32(load32(m.memory[int64(uint32(v1))+4:]))
							v14 = t57
							store32(m.memory[int64(uint32(v7))+8:], uint32(i32(-1)))
							{
								t58 := int32(load32(m.memory[int64(uint32(v7))+100:]))
								t59 := v10
								v5 = t58
								if uint32(t59) >= uint32(v5) {
									m.fn33(v10, v5, i32(1069844))
									panic("unreachable")
								}
								t60 := int32(load32(m.memory[int64(uint32(v7))+96:]))
								v6 = t60 + v10*i32(80)
								t61 := int64(load64(m.memory[int64(uint32(v6))+32:]))
								v12 = t61
								v15 = v12 - v8
								t62 := v15
								v16 = int64(uint32(v13))
								p63 := v16
								if uint64(v15) < uint64(v16) {
									p63 = t62
								}
								v5 = int32(p63)
								p64 := i32(0)
								if uint64(v12) > uint64(v8) {
									p64 = v5
								}
								v17 = p64
								if v17 != 0 {
									v18 = v7 + i32(16)
									t65 := int32(load32(m.memory[int64(uint32(v6))+56:]))
									v11 = t65
									{
										{
											if uint64(v12) < uint64(i64(4096)) {
												v9 = i32(0)
												store32(m.memory[int64(uint32(v4))+56:], uint32(i32(0)))
												store64(m.memory[int64(uint32(v4))+48:], uint64(i64(0x400000000)))
												store32(m.memory[int64(uint32(v4))+60:], uint32(v11))
												v16 = int64(uint32(i32(3))) << 32
												v12 = v16 | int64(uint32(v4+i32(68)))
												v15 = v16 | int64(uint32(v4+i32(64)))
												t68 := int32(load32(m.memory[int64(uint32(v7))+116:]))
												v10 = t68
												t69 := int32(load32(m.memory[int64(uint32(v7))+112:]))
												v19 = t69
												v6 = i32(1)
												v20 = i32(4)
												v5 = v11
											l39:
												{
													{
														if v5 != i32(-2) {
															goto l32
														}
														t70 := int32(load32(m.memory[int64(uint32(v4))+56:]))
														store32(m.memory[int64(uint32(v4))+112:], uint32(t70))
														t71 := int64(load64(m.memory[int64(uint32(v4))+48:]))
														store64(m.memory[int64(uint32(v4))+104:], uint64(t71))
														store32(m.memory[int64(uint32(v4))+116:], uint32(v18))
														store64(m.memory[int64(uint32(v4))+96:], uint64(i64(0)))
														goto l33
													}
												l32:
													{
														t72 := int32(load32(m.memory[int64(uint32(v4))+48:]))
														if v6+i32(-1) != t72 {
															goto l34
														}
														m.fn174(v4 + i32(48))
														t73 := int32(load32(m.memory[int64(uint32(v4))+52:]))
														v20 = t73
													}
												l34:
													store32(m.memory[uint32(v20+v9):], uint32(v5))
													store32(m.memory[int64(uint32(v4))+56:], uint32(v6))
													t74 := int32(load32(m.memory[int64(uint32(v4))+60:]))
													t75 := v4
													v5 = t74
													store32(m.memory[int64(uint32(t75))+64:], uint32(v5))
													{
														{
															if uint32(v5) < uint32(v10) {
																goto l35
															}
															store32(m.memory[int64(uint32(v4))+68:], uint32(v10))
															store64(m.memory[int64(uint32(v4))+8:], uint64(v12))
															store64(m.memory[uint32(v4):], uint64(v15))
															m.fn17(v4+i32(80), i32(1064782), v4)
															m.fn163(v4+i32(72), i32(21), v4+i32(80))
															goto l36
														l35:
															t76 := int32(load32(m.memory[uint32(v19+v5<<2):]))
															t77 := v4
															v5 = t76
															store32(m.memory[int64(uint32(t77))+68:], uint32(v5))
															if v5 == i32(-2) {
																goto l37
															}
															if uint32(v5) < uint32(v10) {
																goto l37
															}
															store64(m.memory[uint32(v4):], uint64(v12))
															m.fn17(v4+i32(32), i32(1066159), v4)
															m.fn163(v4+i32(72), i32(21), v4+i32(32))
														}
													l36:
														t78 := int32(m.memory[int64(uint32(v4))+72])
														if t78 != i32(255) {
															t95 := int64(load64(m.memory[int64(uint32(v4))+72:]))
															store64(m.memory[int64(uint32(v4))+96:], uint64(t95))
															goto l40
														}
														t79 := int32(load32(m.memory[int64(uint32(v4))+76:]))
														v5 = t79
													}
												l37:
													v9 = v9 + i32(4)
													v6 = v6 + i32(1)
													store32(m.memory[int64(uint32(v4))+60:], uint32(v5))
													if v5 != v11 {
														goto l39
													}
												}
												store64(m.memory[int64(uint32(v4))+32:], uint64(v16|int64(uint32(v4+i32(60)))))
												m.fn17(v4, i32(1051009), v4+i32(32))
												m.fn163(v4+i32(96), i32(21), v4)
												goto l40
											}
											m.fn186(v4+i32(96), v18, v11, i32(0))
											t66 := int32(load32(m.memory[int64(uint32(v4))+104:]))
											v6 = t66
											if v6 != i32(-1) {
												goto l30
											}
											t67 := int64(load64(m.memory[int64(uint32(v4))+96:]))
											v12 = t67
											goto l31
										}
									l30:
										t80 := int32(load32(m.memory[int64(uint32(v4))+124:]))
										store32(m.memory[int64(uint32(v4))+28:], uint32(t80))
										t81 := int64(load64(m.memory[int64(uint32(v4))+116:]))
										t82 := v4
										v12 = t81
										store64(m.memory[int64(uint32(t82))+20:], uint64(v12))
										t83 := int64(load64(m.memory[int64(uint32(v4))+108:]))
										store64(m.memory[int64(uint32(v4))+12:], uint64(t83))
										store32(m.memory[int64(uint32(v4))+8:], uint32(v6))
										t84 := int64(load32(m.memory[int64(uint32(v4))+16:]))
										t85 := int32(m.memory[int64(uint32(int32(v12)))+20])
										t87 := v4
										p86 := i64(9)
										if t85 != 0 {
											p86 = i64(12)
										}
										v12 = i64_shl(t84, p86)
										store64(m.memory[int64(uint32(t87))+72:], uint64(v12))
										store64(m.memory[int64(uint32(v4))+48:], uint64(v8))
										{
											{
												if v8 < i64(0) {
													goto l41
												}
												if uint64(v8) <= uint64(v12) {
													goto l42
												}
											l41:
												store64(m.memory[int64(uint32(v4))+104:], uint64(int64(uint32(i32(10)))<<32|int64(uint32(v4+i32(72)))))
												store64(m.memory[int64(uint32(v4))+96:], uint64(int64(uint32(i32(57)))<<32|int64(uint32(v4+i32(48)))))
												m.fn17(v4+i32(80), i32(1064737), v4+i32(96))
												m.fn163(v4+i32(32)|i32(4), i32(20), v4+i32(80))
												t88 := int64(load64(m.memory[int64(uint32(v4))+36:]))
												v12 = t88
												goto l43
											}
										l42:
											store64(m.memory[uint32(v4):], uint64(v8))
											if uint32(v17) > uint32(v13) {
												m.fn121(i32(0), v17, v13, i32(1069860))
												panic("unreachable")
											}
											m.fn294(v4+i32(96), v4, v14, v17)
											t89 := int32(m.memory[int64(uint32(v4))+96])
											if t89 == i32(255) {
												{
													t93 := int32(load32(m.memory[int64(uint32(v4))+8:]))
													v6 = t93
													if v6 == 0 {
														goto l46
													}
													t94 := int32(load32(m.memory[int64(uint32(v4))+12:]))
													m.fn21(t94, v6<<2, i32(4))
												}
											l46:
												v9 = i32(255)
												v8 = i64(0)
												goto l28
											}
											t90 := int64(load64(m.memory[int64(uint32(v4))+96:]))
											v12 = t90
											t91 := int32(load32(m.memory[int64(uint32(v4))+8:]))
											v6 = t91
										}
									l43:
										if v6 == 0 {
											goto l31
										}
										t92 := int32(load32(m.memory[int64(uint32(v4))+12:]))
										m.fn21(t92, v6<<2, i32(4))
									}
								l31:
									v8 = int64(uint64(v12) >> 8)
									v5 = int32(int64(uint64(v12) >> 32))
									v9 = int32(v12)
									goto l28
								}
								v9 = i32(255)
								v8 = i64(0)
								v5 = i32(0)
								goto l28
							}
						}
					l40:
						store32(m.memory[int64(uint32(v4))+104:], uint32(i32(-1)))
						t96 := int32(load32(m.memory[int64(uint32(v4))+48:]))
						v5 = t96
						if v5 == 0 {
							goto l33
						}
						m.fn21(v20, v5<<2, i32(4))
					}
				l33:
					{
						t97 := int32(load32(m.memory[int64(uint32(v4))+104:]))
						v21 = t97
						if v21 != i32(-1) {
							goto l47
						}
						t98 := int64(load64(m.memory[int64(uint32(v4))+96:]))
						v12 = t98
						v8 = int64(uint64(v12) >> 8)
						v5 = int32(int64(uint64(v12) >> 32))
						v9 = int32(v12)
						goto l28
					}
				l47:
					t99 := int32(load32(m.memory[int64(uint32(v4))+116:]))
					v6 = t99
					t100 := int32(load32(m.memory[int64(uint32(v4))+108:]))
					v18 = t100
					t101 := int32(load32(m.memory[int64(uint32(v4))+112:]))
					t102 := v4
					v22 = t101
					v12 = int64(uint32(v22)) << 6
					store64(m.memory[int64(uint32(t102))+48:], uint64(v12))
					store64(m.memory[int64(uint32(v4))+80:], uint64(v8))
					{
						if v8 < i64(0) {
							goto l48
						}
						if uint64(v8) <= uint64(v12) {
							goto l49
						}
					l48:
						store64(m.memory[int64(uint32(v4))+104:], uint64(int64(uint32(i32(10)))<<32|int64(uint32(v4+i32(48)))))
						store64(m.memory[int64(uint32(v4))+96:], uint64(int64(uint32(i32(57)))<<32|int64(uint32(v4+i32(80)))))
						m.fn17(v4+i32(32), i32(1064737), v4+i32(96))
						m.fn163(v4|i32(4), i32(20), v4+i32(32))
						t103 := int64(load64(m.memory[int64(uint32(v4))+4:]))
						v12 = t103
						goto l50
					}
				l49:
					if uint32(v17) > uint32(v13) {
						m.fn121(i32(0), v17, v13, i32(1069876))
						panic("unreachable")
					}
					v23 = v16 | int64(uint32(v4+i32(80)))
					v24 = v4 + i32(32) + i32(4)
					v11 = v17
					{
					l83:
						if v12 == v8 {
							goto l52
						}
						{
							t104 := v22
							v5 = int32(int64(uint64(v8) >> 6))
							if uint32(t104) <= uint32(v5) {
								m.fn33(v5, v22, i32(1079776))
								panic("unreachable")
							}
							{
								t105 := int32(load32(m.memory[int64(uint32(v6))+84:]))
								if t105 == 0 {
									m.fn33(i32(0), i32(0), i32(1069844))
									panic("unreachable")
								}
								t106 := int32(load32(m.memory[uint32(v18+v5<<2):]))
								v20 = t106
								t107 := int32(load32(m.memory[int64(uint32(v6))+80:]))
								t108 := int32(load32(m.memory[int64(uint32(t107))+56:]))
								m.fn186(v4+i32(96), v6, t108, i32(1))
								{
									t109 := int32(load32(m.memory[int64(uint32(v4))+104:]))
									v5 = t109
									if v5 != i32(-1) {
										t112 := int32(load32(m.memory[int64(uint32(v4))+108:]))
										v13 = t112
										{
											{
												{
													t113 := int32(load32(m.memory[int64(uint32(v4))+116:]))
													t114 := v20
													v9 = t113
													t115 := int32(m.memory[int64(uint32(v9))+20])
													v25 = t115
													p116 := i32(3)
													if v25 != 0 {
														p116 = i32(6)
													}
													v26 = p116
													v19 = i32_shr_u(t114, v26)
													t117 := int32(load32(m.memory[int64(uint32(v4))+112:]))
													if uint32(v19) < uint32(t117) {
														goto l57
													}
													m.fn263(v4+i32(96), i32(21), i32(1075644), i32(17))
													t118 := int32(load32(m.memory[int64(uint32(v4))+96:]))
													v20 = t118
													t119 := int32(load32(m.memory[int64(uint32(v4))+100:]))
													v19 = t119
													v9 = i32(0)
													if v5 != 0 {
														goto l58
													}
													goto l59
												}
											l57:
												t120 := int32(load32(m.memory[uint32(v13+v19<<2):]))
												t121 := v4
												v19 = t120
												store32(m.memory[int64(uint32(t121))+80:], uint32(v19))
												{
													{
														t122 := int32(load32(m.memory[int64(uint32(v9))+16:]))
														if uint32(v19) < uint32(t122) {
															goto l60
														}
														store64(m.memory[int64(uint32(v4))+96:], uint64(v23))
														store64(m.memory[int64(uint32(v4))+104:], uint64(v16|int64(uint32(v9+i32(16)))))
														m.fn17(v4, i32(1049021), v4+i32(96))
														m.fn163(v24, i32(21), v4)
														t123 := int32(load32(m.memory[int64(uint32(v4))+36:]))
														v20 = t123
														t124 := int32(load32(m.memory[int64(uint32(v4))+40:]))
														v19 = t124
														v9 = i32(0)
														goto l61
													}
												l60:
													t126 := v9
													t127 := int64(uint32(v19 + i32(1)))
													p125 := i64(9)
													if v25 != 0 {
														p125 = i64(12)
													}
													t128 := i64_shl(t127, p125)
													t129 := v8 & i64(63)
													v20 = v20 & (i32_shl(i32(-1), v26) ^ i32(-1)) << 6
													v15 = t129 | int64(uint32(v20))
													store64(m.memory[int64(uint32(t126))+8:], uint64(t128+v15))
													v19 = int32(v15) - v20
													v20 = i32(64)
												}
											l61:
												if v5 == 0 {
													goto l62
												}
											}
										l58:
											t130 := int32(load32(m.memory[uint32(v13+i32(-4)):]))
											v25 = t130
											v26 = v25 & i32(-8)
											t131 := v26
											v25 = v25 & i32(3)
											p132 := i32(8)
											if v25 != 0 {
												p132 = i32(4)
											}
											v5 = v5 << 2
											if uint32(t131) < uint32(p132+v5) {
												m.fn7(i32(1274404), i32(46), i32(1274452))
												panic("unreachable")
											}
											if v25 == 0 {
												goto l64
											}
											if uint32(v26) > uint32(v5+i32(39)) {
												m.fn7(i32(1274468), i32(46), i32(1274516))
												panic("unreachable")
											}
										l64:
											m.fn5(v13)
										}
									l62:
										if v9 == 0 {
											goto l59
										}
										t133 := v11
										v15 = v12 - v8
										t134 := v15
										v27 = int64(uint32(v11))
										p135 := v27
										if uint64(v15) < uint64(v27) {
											p135 = t134
										}
										v13 = int32(p135)
										if uint32(t133) < uint32(v13) {
											m.fn121(i32(0), v13, v11, i32(1079792))
											panic("unreachable")
										}
										v5 = i32(0)
										{
											if v20 == v19 {
												goto l67
											}
											t136 := int32(load32(m.memory[uint32(v9):]))
											t137 := int64(load64(m.memory[int64(uint32(v9))+8:]))
											v15 = t137
											t138 := int32(load32(m.memory[int64(uint32(v9))+4:]))
											t139 := v15
											v5 = t138
											v27 = int64(uint32(v5))
											p140 := v27
											if uint64(v15) < uint64(v27) {
												p140 = t139
											}
											v26 = t136 + int32(p140)
											{
												t142 := v5
												p141 := i64(0xffffffff)
												if uint64(v15) < uint64(i64(0xffffffff)) {
													p141 = v15
												}
												v25 = t142 - int32(p141)
												p143 := v25
												if uint32(v25) > uint32(v5) {
													p143 = i32(0)
												}
												v5 = p143
												t144 := v5
												v20 = v20 - v19
												p145 := v13
												if uint32(v20) < uint32(v13) {
													p145 = v20
												}
												v20 = p145
												p146 := v20
												if uint32(v5) < uint32(v20) {
													p146 = t144
												}
												v5 = p146
												if v5 != i32(1) {
													goto l68
												}
												t147 := int32(m.memory[uint32(v26)])
												m.memory[uint32(v14)] = byte(t147)
												goto l69
											}
										l68:
											if v5 == 0 {
												goto l69
											}
											memory_copy(m.memory, uint32(v14), uint32(v26), uint32(v5))
										l69:
											store64(m.memory[int64(uint32(v9))+8:], uint64(v15+int64(uint32(v5))))
										}
									l67:
										v10 = v10 | i32(255)
										v8 = v8 + int64(uint32(v5))
										goto l56
									}
									t110 := int32(load32(m.memory[int64(uint32(v4))+100:]))
									v5 = t110
									t111 := int32(load32(m.memory[int64(uint32(v4))+96:]))
									v10 = t111
									goto l56
								}
							}
						}
					l59:
						v5 = v19
						v10 = v20
					l56:
						v9 = i32(0)
						{
							switch v10 & i32(255) {
							case 0:
								goto l70
							case 2:
								t148 := int32(m.memory[int64(uint32(v5))+8])
								if t148 == i32(35) {
									goto l75
								}
								v9 = i32(2)
								goto l70
							case 3:
								t149 := int32(m.memory[int64(uint32(v5))+8])
								if t149 == i32(35) {
									goto l76
								}
								v9 = i32(3)
								goto l70
							case 1:
								if v10&i32(0xff00) == i32(8960) {
									goto l75
								}
								v9 = i32(1)
								goto l70
							default:
								if v5 == 0 {
									goto l52
								}
								if uint32(v11) < uint32(v5) {
									m.fn121(v5, v11, v11, i32(1069360))
									panic("unreachable")
								}
								v14 = v14 + v5
								v11 = v11 - v5
								goto l75
							}
						l76:
							t150 := int32(load32(m.memory[uint32(v5):]))
							v20 = t150
							{
								t151 := int32(load32(m.memory[uint32(v5+i32(4)):]))
								v9 = t151
								t152 := int32(load32(m.memory[uint32(v9):]))
								v13 = t152
								if v13 == 0 {
									goto l78
								}
								m.t0[uint(v13)].(func(int32))(v20)
							}
						l78:
							{
								t153 := int32(load32(m.memory[int64(uint32(v9))+4:]))
								v13 = t153
								if v13 == 0 {
									goto l79
								}
								t154 := int32(load32(m.memory[int64(uint32(v9))+8:]))
								m.fn21(v20, v13, t154)
							}
						l79:
							t155 := int32(load32(m.memory[uint32(v5+i32(-4)):]))
							v9 = t155
							v20 = v9 & i32(-8)
							t156 := v20
							v9 = v9 & i32(3)
							p157 := i32(20)
							if v9 != 0 {
								p157 = i32(16)
							}
							if uint32(t156) < uint32(p157) {
								m.fn7(i32(1274404), i32(46), i32(1274452))
								panic("unreachable")
							}
							if v9 == 0 {
								goto l81
							}
							if uint32(v20) >= uint32(i32(52)) {
								m.fn7(i32(1274468), i32(46), i32(1274516))
								panic("unreachable")
							}
						l81:
							m.fn5(v5)
						}
					l75:
						if v11 != 0 {
							goto l83
						}
						goto l84
					l52:
						t158 := int64(load64(m.memory[int64(uint32(i32(0)))+1277248:]))
						v8 = t158
						v10 = int32(v8)
						if v10&i32(255) == i32(255) {
							goto l84
						}
						v5 = int32(int64(uint64(v8) >> 32))
						v9 = v10
					}
				l70:
					v12 = int64(uint32(v5))<<32 | int64(uint32(v10&i32(-256)|v9&i32(255)))
				l50:
					v8 = int64(uint64(v12) >> 8)
					v5 = int32(int64(uint64(v12) >> 32))
					v9 = int32(v12)
					if v21 == 0 {
						goto l28
					}
					m.fn21(v18, v21<<2, i32(4))
					goto l28
				l84:
					v9 = i32(255)
					v8 = i64(0)
					if v21 != 0 {
						goto l85
					}
					v5 = v17
					goto l28
				l85:
					m.fn21(v18, v21<<2, i32(4))
					v5 = v17
				}
			l28:
				t159 := int32(load32(m.memory[int64(uint32(v7))+8:]))
				v6 = t159
				store32(m.memory[int64(uint32(v7))+8:], uint32(i32(0)))
				store32(m.memory[int64(uint32(v4))+96:], uint32(v6))
				if v6 == i32(-1) {
					goto l86
				}
				m.fn716(v4 + i32(96))
			}
		l25:
			panic("unreachable")
		l86:
			if v9&i32(255) != i32(255) {
				t167 := int32(load32(m.memory[uint32(v7):]))
				t168 := v7
				v1 = t167
				store32(m.memory[uint32(t168):], uint32(v1+i32(-1)))
				if v1 != i32(1) {
					goto l8
				}
				m.fn161(v7)
				goto l8
			}
			store32(m.memory[int64(uint32(v1))+16:], uint32(v5))
			{
				t160 := int32(load32(m.memory[int64(uint32(v1))+12:]))
				if uint32(t160) <= uint32(v5) {
					goto l88
				}
				store32(m.memory[int64(uint32(v1))+12:], uint32(v5))
			}
		l88:
			t161 := int32(load32(m.memory[uint32(v7):]))
			t162 := v7
			v5 = t161
			store32(m.memory[uint32(t162):], uint32(v5+i32(-1)))
			if v5 != i32(1) {
				goto l0
			}
			m.fn161(v7)
		}
	l0:
		t163 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		v7 = t163
		{
			t164 := int32(load32(m.memory[int64(uint32(v1))+16:]))
			v6 = t164
			t165 := int32(load32(m.memory[int64(uint32(v1))+12:]))
			t166 := v6
			v5 = t165
			if uint32(t166) < uint32(v5) {
				goto l89
			}
			if uint32(v6) <= uint32(v7) {
				t169 := int32(load32(m.memory[int64(uint32(v1))+4:]))
				v7 = t169 + v5
				{
					v6 = v6 - v5
					p170 := v3
					if uint32(v6) < uint32(v3) {
						p170 = v6
					}
					v6 = p170
					if v6 != i32(1) {
						goto l91
					}
					t171 := int32(m.memory[uint32(v7)])
					m.memory[uint32(v2)] = byte(t171)
					goto l92
				}
			l91:
				if v6 == 0 {
					goto l92
				}
				memory_copy(m.memory, uint32(v2), uint32(v7), uint32(v6))
			l92:
				store32(m.memory[int64(uint32(v0))+4:], uint32(v6))
				m.memory[uint32(v0)] = byte(i32(255))
				store32(m.memory[int64(uint32(v1))+12:], uint32(v5+v6))
				goto l93
			}
		}
	l89:
		m.fn121(v5, v6, v7, i32(1093324))
		panic("unreachable")
	}
l8:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
	store32(m.memory[uint32(v0):], uint32(v8<<8|int64(uint32(v9))&i64(255)))
l93:
	m.g0 = v4 + i32(128)
}
func (m *Module) fn433(v0, v1, v2, v3 int32) {
	var v4, v5, v6, v7, v8 int32
	v4 = i32(255)
	v5 = i32(0)
	{
		if uint32(v2) < uint32(v3) {
			goto l0
		}
		if uint32(v2-v3) < uint32(i32(4)) {
			goto l0
		}
		v6 = v3 + i32(4)
		{
			{
				t0 := int32(load32(m.memory[uint32(v1+v3):]))
				v7 = t0
				if v7&i32(15) != 0 {
					goto l1
				}
				v8 = i32(2)
				goto l2
			}
		l1:
			if uint32(v2-v6) <= uint32(i32(1)) {
				goto l0
			}
			t1 := int32(m.memory[uint32(v1+v6)])
			p2 := i32(2)
			if v7&i32(1) != 0 {
				p2 = t1 & i32(1)
			}
			v8 = p2
			v6 = v3 + i32(6)
		}
	l2:
		v3 = int32(uint32(v7) >> 3)
		v3 = v3&i32(2) + int32(uint32(v7)>>6)&i32(2) + int32(uint32(v7)>>5)&i32(2) + v3&i32(4) + int32(uint32(v7)>>10)&i32(2) + int32(uint32(v7)>>11)&i32(2) + int32(uint32(v7)>>12)&i32(2) + int32(uint32(v7)>>13)&i32(2) + int32(uint32(v7)>>7)&i32(2) + int32(uint32(v7)>>9)&i32(2) + int32(uint32(v7)>>14)&i32(2) + v6
		{
			if v7&i32(0x100000) == 0 {
				goto l3
			}
			if uint32(v2) < uint32(v3) {
				goto l0
			}
			if uint32(v2-v3) <= uint32(i32(1)) {
				goto l0
			}
			t3 := int32(load16(m.memory[uint32(v1+v3):]))
			v3 = v3 + t3<<2 + i32(2)
		}
	l3:
		v3 = v3 + int32(uint32(v7)>>15)&i32(2)
		p4 := v3
		if v7&i32(0xe0000) != 0 {
			p4 = v3 + i32(2)
		}
		v7 = p4 + int32(uint32(v7)>>20)&i32(2)
		t5 := v7
		var p6 int32
		if uint32(v7) > uint32(v2) {
			p6 = 1
		}
		v7 = p6
		p7 := t5
		if v7 != 0 {
			p7 = i32(0)
		}
		v5 = p7
		p8 := v8
		if v7 != 0 {
			p8 = i32(-1)
		}
		v4 = p8
	}
l0:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
	m.memory[uint32(v0)] = byte(v4)
}
func (m *Module) fn434(v0, v1, v2, v3, v4, v5, v6, v7 int32) {
	var v8, v9, v10, v11, v12, v13 int32
	var v14 int64
	var v15 int32
	var v16, v17 int64
	var v18 int32
	t0 := m.g0
	v8 = t0 - i32(32)
	m.g0 = v8
	{
		t1 := int32(load32(m.memory[uint32(v2):]))
		if t1 != i32(1) {
			store64(m.memory[uint32(v0):], uint64(i64(0xffffffff)))
			goto l6
		}
		v9 = i32(0)
		t2 := int32(load32(m.memory[int64(uint32(v2))+8:]))
		v10 = t2
		var p3 int32
		if v10 != i32(0) {
			p3 = 1
		}
		v11 = p3
		p4 := v11
		if v7 != 0 {
			p4 = i32(0)
		}
		v12 = p4
		t5 := int32(load32(m.memory[int64(uint32(v3))+12:]))
		if t5 == 0 {
			goto l1
		}
		t6 := int64(load64(m.memory[int64(uint32(v3))+16:]))
		t7 := int64(load64(m.memory[int64(uint32(v3))+24:]))
		t8 := int32(load32(m.memory[int64(uint32(v2))+4:]))
		v13 = t8
		t9 := m.fn94(t6, t7, v13)
		v14 = t9
		t10 := int32(load32(m.memory[int64(uint32(v3))+4:]))
		v15 = t10
		v2 = v15 & int32(v14)
		v16 = int64(uint64(v14)>>25) & i64(127) * i64(72340172838076673)
		t11 := int32(load32(m.memory[uint32(v3):]))
		v3 = t11
	l5:
		{
			{
				t12 := int64(load64(m.memory[uint32(v3+v2):]))
				v17 = t12
				v14 = v17 ^ v16
				v14 = (v14 ^ i64(-1)) & (v14 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
				if v14 == 0 {
					goto l2
				}
			l4:
				{
					t13 := v13
					v18 = v3 - (int32(uint32(int64(bits.TrailingZeros64(uint64(v14))))>>3)+v2)&v15<<3
					t14 := int32(load32(m.memory[uint32(v18+i32(-8)):]))
					if t13 == t14 {
						t16 := int32(load32(m.memory[uint32(v18+i32(-4)):]))
						t17 := v5
						v2 = t16
						if uint32(t17) < uint32(v2) {
							goto l1
						}
						v5 = v5 - v2
						if uint32(v5) < uint32(i32(8)) {
							goto l1
						}
						v2 = v4 + v2
						t18 := int32(load32(m.memory[int64(uint32(v2))+4:]))
						v3 = t18
						if uint32(v3) > uint32(v5+i32(-8)) {
							goto l1
						}
						t19 := int32(load16(m.memory[int64(uint32(v2))+2:]))
						if t19 != v6&i32(0xffff) {
							goto l1
						}
						v6 = v2 + i32(8)
						{
							if v7 == 0 {
								goto l7
							}
							v11 = i32(0)
							v2 = i32(0)
						l11:
							{
								if uint32(v3) >= uint32(v2) {
									goto l8
								}
								goto l7
							l8:
								v5 = v3 - v2
								if uint32(v5) >= uint32(i32(8)) {
									goto l9
								}
								goto l7
							l9:
								{
									v7 = v6 + v2
									t20 := int32(load32(m.memory[int64(uint32(v7))+4:]))
									v4 = t20
									if uint32(v4) <= uint32(v5+i32(-8)) {
										goto l10
									}
									goto l7
								}
							l10:
								v2 = v2 + v4 + i32(8)
								t21 := int32(m.memory[int64(uint32(v7))+2])
								t22 := int32(m.memory[int64(uint32(v7))+3])
								if t21<<16|t22<<24 != i32(0x3f10000) {
									goto l11
								}
							}
							if uint32(v4) < uint32(i32(4)) {
								goto l7
							}
							t23 := int32(load32(m.memory[int64(uint32(v7))+8:]))
							v10 = t23
							var p24 int32
							if v10 != i32(0) {
								p24 = 1
							}
							v11 = p24
						}
					l7:
						m.fn396(v8+i32(8), v1, v6, v3)
						t25 := int32(load32(m.memory[int64(uint32(v8))+8:]))
						if t25 == i32(-1) {
							goto l12
						}
						t26 := int64(load64(m.memory[int64(uint32(v8))+24:]))
						store64(m.memory[int64(uint32(v0))+16:], uint64(t26))
						t27 := int64(load64(m.memory[int64(uint32(v8))+16:]))
						store64(m.memory[int64(uint32(v0))+8:], uint64(t27))
						t28 := int64(load64(m.memory[int64(uint32(v8))+8:]))
						store64(m.memory[uint32(v0):], uint64(t28))
						goto l6
					}
					v14 = (v14 + i64(-1)) & v14
					if !(v14 == 0) {
						goto l4
					}
				}
			}
		l2:
			if !(v17&(v17<<1)&i64(-0x7f7f7f7f7f7f7f80) == 0) {
				goto l1
			}
			t15 := v2
			v9 = v9 + i32(8)
			v2 = (t15 + v9) & v15
			goto l5
		}
	}
l1:
	v11 = v12
l12:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v10))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v11))
	store32(m.memory[uint32(v0):], uint32(i32(-1)))
l6:
	m.g0 = v8 + i32(32)
}
func (m *Module) fn435(v0 int32) {
	var v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11, v12 int32
	var v13 int64
	var v14, v15, v16 int32
	var v17 int64
	var v18, v19 int32
	var v20 int64
	var v21 int32
	var v22 int64
	var v23, v24, v25, v26, v27, v28, v29, v30, v31, v32, v33, v34, v35, v36 int32
	t0 := m.g0
	v1 = t0 - i32(80)
	m.g0 = v1
	t1 := int32(load32(m.memory[int64(uint32(v0))+16:]))
	v2 = t1
	store32(m.memory[int64(uint32(v0))+16:], uint32(i32(-1)))
	{
		if v2 == i32(-1) {
			goto l0
		}
		t2 := int32(load32(m.memory[int64(uint32(v0))+44:]))
		v3 = t2
		t3 := int32(load32(m.memory[int64(uint32(v0))+40:]))
		v4 = t3
		t4 := int32(load32(m.memory[int64(uint32(v0))+32:]))
		v5 = t4
		t5 := int32(load32(m.memory[int64(uint32(v0))+28:]))
		v6 = t5
		t6 := int32(load32(m.memory[int64(uint32(v0))+20:]))
		v7 = t6
		{
			t7 := int32(load32(m.memory[int64(uint32(v0))+24:]))
			v8 = t7
			if v8 != 0 {
				t17 := int32(m.memory[int64(uint32(v0))+52])
				v12 = t17
				t18 := int32(load32(m.memory[int64(uint32(v0))+48:]))
				v9 = t18
				t19 := int32(load32(m.memory[int64(uint32(v0))+36:]))
				v10 = t19
				t20 := int64(load64(m.memory[uint32(v0):]))
				t21 := v0
				v13 = t20 + i64(1)
				store64(m.memory[uint32(t21):], uint64(v13))
				if v6 != i32(-1) {
					goto l14
				}
				v5 = i32(4)
				v6 = i32(0)
				v10 = i32(0)
				v4 = i32(0)
				v3 = i32(4)
				v9 = i32(0)
			l14:
				v14 = i32(1)
				v15 = i32(0)
				{
					{
						t22 := int32(load32(m.memory[int64(uint32(v0))+56:]))
						v11 = t22
						t23 := int32(load32(m.memory[int64(uint32(v0))+104:]))
						if uint32(v11) < uint32(t23) {
							goto l15
						}
						v16 = i32(0)
						goto l16
					}
				l15:
					{
						t24 := int32(load32(m.memory[int64(uint32(v0))+100:]))
						v11 = t24 + v11*i32(40)
						t25 := int32(load32(m.memory[int64(uint32(v11))+20:]))
						if t25 != 0 {
							goto l17
						}
						v16 = i32(0)
						goto l16
					}
				l17:
					t26 := int64(load64(m.memory[int64(uint32(v11))+24:]))
					t27 := int64(load64(m.memory[int64(uint32(v11))+32:]))
					t28 := m.fn106(t26, t27, v12)
					v17 = t28
					t29 := int32(load32(m.memory[int64(uint32(v11))+12:]))
					v18 = t29
					v19 = v18 & int32(v17)
					v20 = int64(uint64(v17)>>25) & i64(127) * i64(72340172838076673)
					t30 := int32(load32(m.memory[int64(uint32(v11))+8:]))
					v11 = t30
					v21 = i32(0)
				l22:
					{
						t31 := int64(load64(m.memory[uint32(v11+v19):]))
						v22 = t31
						v17 = v22 ^ v20
						v17 = (v17 ^ i64(-1)) & (v17 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
						if v17 == 0 {
							goto l18
						}
					l20:
						{
							v23 = v11 - (int32(uint32(int64(bits.TrailingZeros64(uint64(v17))))>>3)+v19)&v18<<4
							t32 := int32(load16(m.memory[uint32(v23+i32(-16)):]))
							if t32 == v12 {
								goto l19
							}
							v17 = (v17 + i64(-1)) & v17
							if !(v17 == 0) {
								goto l20
							}
						}
					}
				l18:
					if v22&(v22<<1)&i64(-0x7f7f7f7f7f7f7f80) == 0 {
						t33 := v19
						v21 = v21 + i32(8)
						v19 = (t33 + v21) & v18
						goto l22
					}
					v16 = i32(0)
					goto l16
				l19:
					{
						t34 := int32(load32(m.memory[uint32(v23+i32(-4)):]))
						v16 = t34
						if v16 != 0 {
							goto l23
						}
						v16 = i32(0)
						goto l16
					}
				l23:
					t35 := int32(load32(m.memory[uint32(v23+i32(-8)):]))
					v19 = t35
					v11 = v16 * i32(3)
					t36 := m.fn11(v11)
					v14 = t36
					if v14 == 0 {
						m.fn16(i32(1), v11)
						panic("unreachable")
					}
					if v11 == 0 {
						goto l16
					}
					memory_copy(m.memory, uint32(v14), uint32(v19), uint32(v11))
				}
			l16:
				v19 = i32(-1)
				v23 = i32(-1)
				v24 = v3
				{
					if v9 == 0 {
						goto l25
					}
					v24 = v3 + i32(8)
					t37 := int32(load32(m.memory[uint32(v3):]))
					v23 = t37
					v15 = v3
				}
			l25:
				v9 = v9 << 3
				v18 = v10 << 3
				v11 = i32(0)
				v25 = v5
				{
					if v10 == 0 {
						goto l26
					}
					v25 = v5 + i32(8)
					t38 := int32(load32(m.memory[uint32(v5):]))
					v19 = t38
					v11 = v5
				}
			l26:
				v26 = v3 + v9
				v27 = v5 + v18
				store32(m.memory[int64(uint32(v1))+8:], uint32(i32(0)))
				store64(m.memory[uint32(v1):], uint64(i64(0x400000000)))
				store32(m.memory[int64(uint32(v1))+20:], uint32(i32(0)))
				store64(m.memory[int64(uint32(v1))+12:], uint64(i64(0x400000000)))
				store32(m.memory[int64(uint32(v1))+32:], uint32(i32(0)))
				store64(m.memory[int64(uint32(v1))+24:], uint64(i64(0x100000000)))
				v28 = v7 + v8
				v29 = i32(4)
				v30 = i32(0)
				v31 = i32(0)
				v8 = i32(0)
				v21 = i32(1)
				v32 = i32(0)
				v33 = i32(0)
				v34 = i32(0)
				v9 = v7
			l95:
				{
					{
						t39 := int32(int8(m.memory[uint32(v9)]))
						v10 = t39
						if v10 <= i32(-1) {
							goto l27
						}
						v9 = v9 + i32(1)
						v10 = v10 & i32(255)
						goto l28
					}
				l27:
					t40 := int32(m.memory[int64(uint32(v9))+1])
					v35 = t40 & i32(63)
					v18 = v10 & i32(31)
					if uint32(v10) > uint32(i32(-33)) {
						goto l29
					}
					v10 = v18<<6 | v35
					v9 = v9 + i32(2)
					goto l28
				l29:
					t41 := int32(m.memory[int64(uint32(v9))+2])
					v35 = v35<<6 | t41&i32(63)
					if uint32(v10) >= uint32(i32(-16)) {
						goto l30
					}
					v10 = v35 | v18<<12
					v9 = v9 + i32(3)
					goto l28
				l30:
					t42 := int32(m.memory[int64(uint32(v9))+3])
					v10 = v35<<6 | t42&i32(63) | v18<<18&i32(0x1c0000)
					v9 = v9 + i32(4)
				}
			l28:
				{
					if v11 != 0 {
						goto l31
					}
					v18 = i32(0)
					goto l32
				l31:
					t43 := int32(load16(m.memory[int64(uint32(v11))+4:]))
					v18 = t43
				}
			l32:
				{
					{
						if uint32(v18) < uint32(v16) {
							v35 = v14 + v18*i32(3)
							t44 := int32(m.memory[int64(uint32(v35))+2])
							v18 = t44
							t45 := int32(m.memory[int64(uint32(v35))+1])
							v35 = t45
							if v15 == 0 {
								goto l35
							}
							t46 := int32(m.memory[int64(uint32(v15))+4])
							t47 := v35
							v36 = t46
							p48 := v36
							if v36 == i32(2) {
								p48 = t47
							}
							v35 = p48
							goto l36
						}
						if v15 != 0 {
							goto l34
						}
						v35 = i32(0)
						v18 = i32(0)
						goto l35
					l34:
						t49 := int32(m.memory[int64(uint32(v15))+4])
						v35 = t49
						v18 = i32(2)
					}
				l36:
					t50 := int32(m.memory[int64(uint32(v15))+5])
					t51 := v18
					v36 = t50
					p52 := v36
					if v36 == i32(2) {
						p52 = t51
					}
					v18 = p52
				}
			l35:
				{
					{
						{
							switch v10 + i32(-11) {
							case 2:
								if v31 == 0 {
									goto l40
								}
								t53 := int32(load32(m.memory[int64(uint32(v1))+24:]))
								v10 = t53
								store32(m.memory[int64(uint32(v1))+24:], uint32(i32(0)))
								store64(m.memory[int64(uint32(v1))+28:], uint64(i64(1)))
								m.fn446(v1+i32(48), v21, v31)
								{
									if v10 == 0 {
										goto l41
									}
									t54 := int32(load32(m.memory[uint32(v21+i32(-4)):]))
									v18 = t54
									v35 = v18 & i32(-8)
									t55 := v35
									v18 = v18 & i32(3)
									p56 := i32(8)
									if v18 != 0 {
										p56 = i32(4)
									}
									if uint32(t55) < uint32(p56+v10) {
										m.fn7(i32(1274404), i32(46), i32(1274452))
										panic("unreachable")
									}
									if v18 == 0 {
										goto l43
									}
									if uint32(v35) > uint32(v10+i32(39)) {
										m.fn7(i32(1274468), i32(46), i32(1274516))
										panic("unreachable")
									}
								l43:
									m.fn5(v21)
								}
							l41:
								{
									{
										t57 := int32(load32(m.memory[int64(uint32(v1))+56:]))
										if t57 != 0 {
											goto l45
										}
										t58 := int32(load32(m.memory[int64(uint32(v1))+48:]))
										v10 = t58
										if v10 == 0 {
											goto l46
										}
										t59 := int32(load32(m.memory[int64(uint32(v1))+52:]))
										v18 = t59
										t60 := int32(load32(m.memory[uint32(v18+i32(-4)):]))
										v8 = t60
										v21 = v8 & i32(-8)
										t61 := v21
										v8 = v8 & i32(3)
										p62 := i32(8)
										if v8 != 0 {
											p62 = i32(4)
										}
										if uint32(t61) < uint32(p62+v10) {
											m.fn7(i32(1274404), i32(46), i32(1274452))
											panic("unreachable")
										}
										if v8 == 0 {
											goto l48
										}
										if uint32(v21) > uint32(v10+i32(39)) {
											m.fn7(i32(1274468), i32(46), i32(1274516))
											panic("unreachable")
										}
									l48:
										m.fn5(v18)
										goto l46
									}
								l45:
									t63 := int32(load32(m.memory[int64(uint32(v1))+56:]))
									store32(m.memory[int64(uint32(v1))+72:], uint32(t63))
									t64 := int64(load64(m.memory[int64(uint32(v1))+48:]))
									store64(m.memory[int64(uint32(v1))+64:], uint64(t64))
									v18 = v34&i32(255)<<8 | v33&i32(255)
									{
										t65 := int32(load32(m.memory[int64(uint32(v1))+12:]))
										if v8 != t65 {
											goto l50
										}
										m.fn315(v1 + i32(12))
										t66 := int32(load32(m.memory[int64(uint32(v1))+16:]))
										v29 = t66
									}
								l50:
									v10 = v29 + v8*i32(28)
									store32(m.memory[uint32(v10):], uint32(i32(3)))
									t67 := int64(load64(m.memory[int64(uint32(v1))+64:]))
									store64(m.memory[int64(uint32(v10))+4:], uint64(t67))
									t68 := int32(load32(m.memory[int64(uint32(v1))+72:]))
									store32(m.memory[int64(uint32(v10))+12:], uint32(t68))
									store32(m.memory[int64(uint32(v10))+16:], uint32(v18))
									store32(m.memory[int64(uint32(v1))+20:], uint32(v8+i32(1)))
								}
							l46:
								v21 = i32(1)
								v30 = i32(0)
								goto l40
							case 0:
								if v32 == 0 {
									goto l51
								}
								t69 := int32(load32(m.memory[int64(uint32(v1))+24:]))
								v10 = t69
								store32(m.memory[int64(uint32(v1))+24:], uint32(i32(0)))
								store64(m.memory[int64(uint32(v1))+28:], uint64(i64(1)))
								m.fn446(v1+i32(48), v21, v32)
								{
									if v10 == 0 {
										goto l52
									}
									t70 := int32(load32(m.memory[uint32(v21+i32(-4)):]))
									v18 = t70
									v35 = v18 & i32(-8)
									t71 := v35
									v18 = v18 & i32(3)
									p72 := i32(8)
									if v18 != 0 {
										p72 = i32(4)
									}
									if uint32(t71) < uint32(p72+v10) {
										m.fn7(i32(1274404), i32(46), i32(1274452))
										panic("unreachable")
									}
									if v18 == 0 {
										goto l54
									}
									if uint32(v35) > uint32(v10+i32(39)) {
										m.fn7(i32(1274468), i32(46), i32(1274516))
										panic("unreachable")
									}
								l54:
									m.fn5(v21)
								}
							l52:
								t73 := int32(load32(m.memory[int64(uint32(v1))+56:]))
								if t73 != 0 {
									t79 := int32(load32(m.memory[int64(uint32(v1))+56:]))
									store32(m.memory[int64(uint32(v1))+72:], uint32(t79))
									t80 := int64(load64(m.memory[int64(uint32(v1))+48:]))
									store64(m.memory[int64(uint32(v1))+64:], uint64(t80))
									v18 = v34&i32(255)<<8 | v33&i32(255)
									{
										t81 := int32(load32(m.memory[int64(uint32(v1))+12:]))
										if v8 != t81 {
											goto l64
										}
										m.fn315(v1 + i32(12))
										t82 := int32(load32(m.memory[int64(uint32(v1))+16:]))
										v29 = t82
									}
								l64:
									v10 = v29 + v8*i32(28)
									store32(m.memory[uint32(v10):], uint32(i32(3)))
									t83 := int64(load64(m.memory[int64(uint32(v1))+64:]))
									store64(m.memory[int64(uint32(v10))+4:], uint64(t83))
									t84 := int32(load32(m.memory[int64(uint32(v1))+72:]))
									store32(m.memory[int64(uint32(v10))+12:], uint32(t84))
									store32(m.memory[int64(uint32(v10))+16:], uint32(v18))
									t85 := v1
									v8 = v8 + i32(1)
									store32(m.memory[int64(uint32(t85))+20:], uint32(v8))
									goto l57
								}
								t74 := int32(load32(m.memory[int64(uint32(v1))+48:]))
								v10 = t74
								if v10 == 0 {
									goto l57
								}
								t75 := int32(load32(m.memory[int64(uint32(v1))+52:]))
								v21 = t75
								t76 := int32(load32(m.memory[uint32(v21+i32(-4)):]))
								v18 = t76
								v35 = v18 & i32(-8)
								t77 := v35
								v18 = v18 & i32(3)
								p78 := i32(8)
								if v18 != 0 {
									p78 = i32(4)
								}
								if uint32(t77) < uint32(p78+v10) {
									m.fn7(i32(1274404), i32(46), i32(1274452))
									panic("unreachable")
								}
								if v18 == 0 {
									goto l59
								}
								if uint32(v35) > uint32(v10+i32(39)) {
									m.fn7(i32(1274468), i32(46), i32(1274516))
									panic("unreachable")
								}
							l59:
								m.fn5(v21)
								goto l57
							default:
								if (v35^v33)&i32(1) != 0 {
									goto l61
								}
								if (v18^v34)&i32(1) == 0 {
									goto l62
								}
							l61:
								if v30 != 0 {
									t86 := int32(load32(m.memory[int64(uint32(v1))+24:]))
									v32 = t86
									store32(m.memory[int64(uint32(v1))+24:], uint32(i32(0)))
									store64(m.memory[int64(uint32(v1))+28:], uint64(i64(1)))
									m.fn446(v1+i32(48), v21, v30)
									{
										if v32 == 0 {
											goto l65
										}
										t87 := int32(load32(m.memory[uint32(v21+i32(-4)):]))
										v30 = t87
										v31 = v30 & i32(-8)
										t88 := v31
										v30 = v30 & i32(3)
										p89 := i32(8)
										if v30 != 0 {
											p89 = i32(4)
										}
										if uint32(t88) < uint32(p89+v32) {
											m.fn7(i32(1274404), i32(46), i32(1274452))
											panic("unreachable")
										}
										if v30 == 0 {
											goto l67
										}
										if uint32(v31) > uint32(v32+i32(39)) {
											m.fn7(i32(1274468), i32(46), i32(1274516))
											panic("unreachable")
										}
									l67:
										m.fn5(v21)
									}
								l65:
									{
										t90 := int32(load32(m.memory[int64(uint32(v1))+56:]))
										if t90 != 0 {
											t96 := int32(load32(m.memory[int64(uint32(v1))+56:]))
											store32(m.memory[int64(uint32(v1))+72:], uint32(t96))
											t97 := int64(load64(m.memory[int64(uint32(v1))+48:]))
											store64(m.memory[int64(uint32(v1))+64:], uint64(t97))
											{
												t98 := int32(load32(m.memory[int64(uint32(v1))+12:]))
												if v8 != t98 {
													goto l73
												}
												m.fn315(v1 + i32(12))
												t99 := int32(load32(m.memory[int64(uint32(v1))+16:]))
												v29 = t99
											}
										l73:
											v21 = v29 + v8*i32(28)
											store32(m.memory[uint32(v21):], uint32(i32(3)))
											t100 := int64(load64(m.memory[int64(uint32(v1))+64:]))
											store64(m.memory[int64(uint32(v21))+4:], uint64(t100))
											t101 := int32(load32(m.memory[int64(uint32(v1))+72:]))
											store32(m.memory[int64(uint32(v21))+12:], uint32(t101))
											store16(m.memory[int64(uint32(v21))+18:], uint16(i32(0)))
											m.memory[int64(uint32(v21))+17] = byte(v34)
											m.memory[int64(uint32(v21))+16] = byte(v33)
											t102 := v1
											v8 = v8 + i32(1)
											store32(m.memory[int64(uint32(t102))+20:], uint32(v8))
											t103 := int32(load32(m.memory[int64(uint32(v1))+32:]))
											v30 = t103
											goto l62
										}
										v30 = i32(0)
										t91 := int32(load32(m.memory[int64(uint32(v1))+48:]))
										v21 = t91
										if v21 == 0 {
											goto l62
										}
										t92 := int32(load32(m.memory[int64(uint32(v1))+52:]))
										v31 = t92
										t93 := int32(load32(m.memory[uint32(v31+i32(-4)):]))
										v32 = t93
										v36 = v32 & i32(-8)
										t94 := v36
										v32 = v32 & i32(3)
										p95 := i32(8)
										if v32 != 0 {
											p95 = i32(4)
										}
										if uint32(t94) < uint32(p95+v21) {
											m.fn7(i32(1274404), i32(46), i32(1274452))
											panic("unreachable")
										}
										if v32 == 0 {
											goto l71
										}
										if uint32(v36) > uint32(v21+i32(39)) {
											m.fn7(i32(1274468), i32(46), i32(1274516))
											panic("unreachable")
										}
									l71:
										m.fn5(v31)
										goto l62
									}
								}
								v30 = i32(0)
								goto l62
							}
						l62:
							{
								{
									var p104 int32
									if uint32(v10) < uint32(i32(128)) {
										p104 = 1
									}
									v36 = p104
									if v36 == 0 {
										goto l74
									}
									v32 = i32(1)
									goto l75
								}
							l74:
								if uint32(v10) >= uint32(i32(2048)) {
									goto l76
								}
								v32 = i32(2)
								goto l75
							l76:
								p105 := i32(4)
								if uint32(v10) < uint32(i32(65536)) {
									p105 = i32(3)
								}
								v32 = p105
							}
						l75:
							{
								t106 := int32(load32(m.memory[int64(uint32(v1))+24:]))
								if uint32(v32) <= uint32(t106-v30) {
									goto l77
								}
								m.fn197(v1+i32(24), v30, v32, i32(1), i32(1))
							}
						l77:
							t107 := int32(load32(m.memory[int64(uint32(v1))+28:]))
							v21 = t107
							v31 = v21 + v30
							if v36 != 0 {
								goto l78
							}
							v36 = v10&i32(63) | i32(-128)
							v33 = int32(uint32(v10) >> 6)
							if uint32(v10) >= uint32(i32(2048)) {
								v34 = int32(uint32(v10) >> 12)
								v33 = v33&i32(63) | i32(-128)
								if uint32(v10) > uint32(i32(0xffff)) {
									m.memory[int64(uint32(v31))+3] = byte(v36)
									m.memory[int64(uint32(v31))+2] = byte(v33)
									m.memory[int64(uint32(v31))+1] = byte(v34&i32(63) | i32(-128))
									m.memory[uint32(v31)] = byte(int32(uint32(v10)>>18) | i32(-16))
									goto l80
								}
								m.memory[int64(uint32(v31))+2] = byte(v36)
								m.memory[int64(uint32(v31))+1] = byte(v33)
								m.memory[uint32(v31)] = byte(v34 | i32(224))
								goto l80
							}
							m.memory[int64(uint32(v31))+1] = byte(v36)
							m.memory[uint32(v31)] = byte(v33 | i32(192))
							goto l80
						l78:
							m.memory[uint32(v31)] = byte(v10)
						l80:
							v34 = v18 & i32(1)
							v33 = v35 & i32(1)
							t108 := v1
							v30 = v32 + v30
							store32(m.memory[int64(uint32(t108))+32:], uint32(v30))
							p109 := i32(2)
							if uint32(v10) < uint32(i32(65536)) {
								p109 = i32(1)
							}
							v10 = p109
							v31 = v30
							v32 = v30
							goto l82
						}
					l57:
						v21 = i32(1)
						v30 = i32(0)
						v31 = i32(0)
					l51:
						{
							t110 := int32(load32(m.memory[int64(uint32(v1))+12:]))
							if v8 != t110 {
								goto l83
							}
							m.fn315(v1 + i32(12))
						}
					l83:
						t111 := int32(load32(m.memory[int64(uint32(v1))+16:]))
						v29 = t111
						store32(m.memory[uint32(v29+v8*i32(28)):], uint32(i32(8)))
						v10 = i32(1)
						t112 := v1
						v8 = v8 + i32(1)
						store32(m.memory[int64(uint32(t112))+20:], uint32(v8))
						goto l84
					}
				l40:
					{
						if v11 != 0 {
							goto l85
						}
						v18 = i32(2)
						v35 = i32(0)
						goto l86
					l85:
						t113 := int32(m.memory[int64(uint32(v11))+6])
						v18 = t113
						t114 := int32(load16(m.memory[int64(uint32(v11))+4:]))
						v35 = t114
					}
				l86:
					t115 := int64(load64(m.memory[int64(uint32(v1))+12:]))
					v17 = t115
					store64(m.memory[int64(uint32(v1))+12:], uint64(i64(0x400000000)))
					t116 := int32(load32(m.memory[int64(uint32(v1))+20:]))
					v10 = t116
					store32(m.memory[int64(uint32(v1))+20:], uint32(i32(0)))
					store32(m.memory[int64(uint32(v1))+72:], uint32(v10))
					store64(m.memory[int64(uint32(v1))+64:], uint64(v17))
					{
						t117 := int32(load32(m.memory[int64(uint32(v1))+8:]))
						v8 = t117
						t118 := int32(load32(m.memory[uint32(v1):]))
						if v8 != t118 {
							goto l87
						}
						m.fn312(v1)
					}
				l87:
					v29 = i32(4)
					t119 := int32(load32(m.memory[int64(uint32(v1))+4:]))
					v10 = t119 + v8<<4
					t120 := int64(load64(m.memory[int64(uint32(v1))+64:]))
					store64(m.memory[uint32(v10):], uint64(t120))
					t121 := int32(load32(m.memory[int64(uint32(v1))+72:]))
					store32(m.memory[int64(uint32(v10))+8:], uint32(t121))
					m.memory[int64(uint32(v10))+14] = byte(v18)
					store16(m.memory[int64(uint32(v10))+12:], uint16(v35))
					v10 = i32(1)
					store32(m.memory[int64(uint32(v1))+8:], uint32(v8+i32(1)))
					v31 = i32(0)
					v8 = i32(0)
				}
			l84:
				v32 = i32(0)
			l82:
				{
					{
						if uint32(v23) <= uint32(v10) {
							goto l88
						}
						v18 = v23 - v10
						p122 := v18
						if uint32(v18) > uint32(v23) {
							p122 = i32(0)
						}
						v23 = p122
						goto l89
					}
				l88:
					if v24 != v26 {
						goto l90
					}
					v23 = i32(-1)
					v15 = i32(0)
					goto l89
				l90:
					t123 := int32(load32(m.memory[uint32(v24):]))
					v23 = t123
					v15 = v24
					v24 = v24 + i32(8)
				}
			l89:
				{
					{
						if uint32(v19) <= uint32(v10) {
							goto l91
						}
						v10 = v19 - v10
						p124 := v10
						if uint32(v10) > uint32(v19) {
							p124 = i32(0)
						}
						v19 = p124
						goto l92
					}
				l91:
					if v25 != v27 {
						goto l93
					}
					v19 = i32(-1)
					v11 = i32(0)
					goto l92
				l93:
					t125 := int32(load32(m.memory[uint32(v25):]))
					v19 = t125
					v11 = v25
					v25 = v25 + i32(8)
				}
			l92:
				if v9 == v28 {
					{
						t126 := int32(load32(m.memory[int64(uint32(v1))+32:]))
						v9 = t126
						if v9 == 0 {
							goto l96
						}
						m.fn446(v1+i32(36), v21, v9)
						{
							t127 := int32(load32(m.memory[int64(uint32(v1))+44:]))
							if t127 != 0 {
								{
									t133 := int32(load32(m.memory[int64(uint32(v1))+12:]))
									if v8 != t133 {
										goto l101
									}
									m.fn315(v1 + i32(12))
								}
							l101:
								t134 := int32(load32(m.memory[int64(uint32(v1))+16:]))
								v9 = t134 + v8*i32(28)
								store32(m.memory[uint32(v9):], uint32(i32(3)))
								t135 := int64(load64(m.memory[int64(uint32(v1))+36:]))
								store64(m.memory[int64(uint32(v9))+4:], uint64(t135))
								t136 := int32(load32(m.memory[int64(uint32(v1))+44:]))
								store32(m.memory[int64(uint32(v9))+12:], uint32(t136))
								store16(m.memory[int64(uint32(v9))+18:], uint16(i32(0)))
								m.memory[int64(uint32(v9))+17] = byte(v34)
								m.memory[int64(uint32(v9))+16] = byte(v33)
								t137 := v1
								v8 = v8 + i32(1)
								store32(m.memory[int64(uint32(t137))+20:], uint32(v8))
								goto l96
							}
							t128 := int32(load32(m.memory[int64(uint32(v1))+36:]))
							v9 = t128
							if v9 == 0 {
								goto l96
							}
							t129 := int32(load32(m.memory[int64(uint32(v1))+40:]))
							v19 = t129
							t130 := int32(load32(m.memory[uint32(v19+i32(-4)):]))
							v10 = t130
							v23 = v10 & i32(-8)
							t131 := v23
							v10 = v10 & i32(3)
							p132 := i32(8)
							if v10 != 0 {
								p132 = i32(4)
							}
							if uint32(t131) < uint32(p132+v9) {
								m.fn7(i32(1274404), i32(46), i32(1274452))
								panic("unreachable")
							}
							if v10 == 0 {
								goto l99
							}
							if uint32(v23) > uint32(v9+i32(39)) {
								m.fn7(i32(1274468), i32(46), i32(1274516))
								panic("unreachable")
							}
						l99:
							m.fn5(v19)
							goto l96
						}
					}
				l96:
					{
						{
							if v8 != 0 {
								goto l102
							}
							t138 := int32(load32(m.memory[int64(uint32(v1))+8:]))
							v9 = t138
							t139 := int32(load32(m.memory[int64(uint32(v1))+4:]))
							v32 = t139
							goto l103
						}
					l102:
						{
							if v11 != 0 {
								goto l104
							}
							v19 = i32(2)
							v11 = i32(0)
							goto l105
						l104:
							t140 := int32(m.memory[int64(uint32(v11))+6])
							v19 = t140
							t141 := int32(load16(m.memory[int64(uint32(v11))+4:]))
							v11 = t141
						}
					l105:
						{
							t142 := int32(load32(m.memory[int64(uint32(v1))+8:]))
							v10 = t142
							t143 := int32(load32(m.memory[uint32(v1):]))
							if v10 != t143 {
								goto l106
							}
							m.fn312(v1)
						}
					l106:
						t144 := int32(load32(m.memory[int64(uint32(v1))+4:]))
						v32 = t144
						v9 = v32 + v10<<4
						t145 := int64(load64(m.memory[int64(uint32(v1))+12:]))
						store64(m.memory[uint32(v9):], uint64(t145))
						t146 := int32(load32(m.memory[int64(uint32(v1))+20:]))
						store32(m.memory[int64(uint32(v9))+8:], uint32(t146))
						m.memory[int64(uint32(v9))+14] = byte(v19)
						store16(m.memory[int64(uint32(v9))+12:], uint16(v11))
						t147 := v1
						v9 = v10 + i32(1)
						store32(m.memory[int64(uint32(t147))+8:], uint32(v9))
					}
				l103:
					v35 = v32 + v9<<4
					t148 := int32(load32(m.memory[uint32(v1):]))
					v31 = t148
					v15 = v32
					if v9 == 0 {
						goto l107
					}
					v25 = v0 + i32(84)
					v24 = v0 + i32(72)
					v15 = v32
				l129:
					{
						v9 = v15
						v15 = v9 + i32(16)
						t149 := int32(load32(m.memory[uint32(v9):]))
						v23 = t149
						if v23 == i32(-1) {
							goto l107
						}
						t150 := int32(m.memory[int64(uint32(v9))+14])
						v30 = t150
						t151 := int32(load16(m.memory[int64(uint32(v9))+12:]))
						v18 = t151
						t152 := int64(load64(m.memory[int64(uint32(v9))+4:]))
						v17 = t152
						v20 = int64(uint64(v17) >> 32)
						v10 = int32(v20)
						v19 = v10 * i32(28)
						v11 = int32(v17)
						v9 = i32(0)
						{
							{
							l109:
								{
									if v19 == v9 {
										m.fn436(v24, v25)
										if v20 == 0 {
											goto l112
										}
										v9 = v11
									l113:
										m.fn332(v9)
										v9 = v9 + i32(28)
										v10 = v10 + i32(-1)
										if v10 != 0 {
											goto l113
										}
									l112:
										if v23 == 0 {
											goto l114
										}
										t155 := int32(load32(m.memory[uint32(v11+i32(-4)):]))
										v9 = t155
										v10 = v9 & i32(-8)
										t156 := v10
										v9 = v9 & i32(3)
										p157 := i32(8)
										if v9 != 0 {
											p157 = i32(4)
										}
										v19 = v23 * i32(28)
										if uint32(t156) < uint32(p157+v19) {
											m.fn7(i32(1274404), i32(46), i32(1274452))
											panic("unreachable")
										}
										if v9 == 0 {
											goto l116
										}
										if uint32(v10) > uint32(v19+i32(39)) {
											m.fn7(i32(1274468), i32(46), i32(1274516))
											panic("unreachable")
										}
									l116:
										m.fn5(v11)
										goto l114
									}
									t153 := v11
									v9 = v9 + i32(28)
									t154 := m.fn306(t153 + v9 + i32(-28))
									if t154 != 0 {
										goto l109
									}
								}
								switch v12 {
								case 0, 6:
									m.fn436(v24, v25)
									v9 = i32(0)
									if uint32(v16) > uint32(v18) {
										v18 = v14 + v18*i32(3)
										t163 := int32(m.memory[int64(uint32(v18))+1])
										v19 = t163
										t164 := int32(m.memory[int64(uint32(v18))+2])
										v18 = t164
										if v18 != i32(2) {
											goto l126
										}
										goto l119
									}
									v19 = i32(0)
									goto l119
								default:
									{
										if uint32(v16) > uint32(v18) {
											goto l120
										}
										if v30&i32(255) != i32(2) {
											goto l121
										}
										goto l122
									l120:
										if v30&i32(255) != i32(2) {
											goto l121
										}
										t158 := int32(m.memory[uint32(v14+v18*i32(3))])
										v30 = t158
									}
								l121:
									switch v30 & i32(255) {
									case 0, 2:
										goto l122
									default:
										t159 := m.fn11(i32(32))
										v10 = t159
										if v10 == 0 {
											m.fn23(i32(8), i32(32))
											panic("unreachable")
										}
										store64(m.memory[int64(uint32(v10))+8:], uint64(v17))
										store32(m.memory[int64(uint32(v10))+4:], uint32(v23))
										store32(m.memory[uint32(v10):], uint32(i32(-0x80000000)))
										{
											t160 := int32(load32(m.memory[int64(uint32(v0))+92:]))
											v9 = t160
											t161 := int32(load32(m.memory[int64(uint32(v0))+84:]))
											if v9 != t161 {
												goto l125
											}
											m.fn319(v25)
										}
									l125:
										store32(m.memory[int64(uint32(v0))+92:], uint32(v9+i32(1)))
										t162 := int32(load32(m.memory[int64(uint32(v0))+88:]))
										v9 = t162 + v9*i32(56)
										store32(m.memory[int64(uint32(v9))+48:], uint32(i32(1)))
										store32(m.memory[int64(uint32(v9))+44:], uint32(v10))
										store32(m.memory[int64(uint32(v9))+40:], uint32(i32(1)))
										store32(m.memory[int64(uint32(v9))+28:], uint32(i32(-1)))
										store32(m.memory[int64(uint32(v9))+24:], uint32(v18))
										store64(m.memory[int64(uint32(v9))+16:], uint64(i64(0)))
										m.memory[int64(uint32(v9))+8] = byte(i32(0))
										store64(m.memory[uint32(v9):], uint64(v13))
										goto l114
									}
								}
							l126:
								v9 = v18 & i32(1) << 8
							l119:
								m.fn454(v11, v10, v9|v19&i32(1))
								store32(m.memory[int64(uint32(v1))+72:], uint32(i32(0)))
								store64(m.memory[int64(uint32(v1))+64:], uint64(i64(0x100000000)))
								m.fn455(v11, v10, v1+i32(64))
								t165 := int32(load32(m.memory[int64(uint32(v1))+72:]))
								store32(m.memory[int64(uint32(v1))+56:], uint32(t165))
								t166 := int64(load64(m.memory[int64(uint32(v1))+64:]))
								store64(m.memory[int64(uint32(v1))+48:], uint64(t166))
								{
									t167 := int32(load32(m.memory[int64(uint32(v0))+80:]))
									v9 = t167
									t168 := int32(load32(m.memory[int64(uint32(v0))+72:]))
									if v9 != t168 {
										goto l127
									}
									m.fn310(v24)
								}
							l127:
								store32(m.memory[int64(uint32(v0))+80:], uint32(v9+i32(1)))
								t169 := int32(load32(m.memory[int64(uint32(v0))+76:]))
								v9 = t169 + v9<<5
								store64(m.memory[int64(uint32(v9))+4:], uint64(v17))
								store32(m.memory[uint32(v9):], uint32(v23))
								t170 := int64(load64(m.memory[int64(uint32(v1))+48:]))
								store64(m.memory[int64(uint32(v9))+12:], uint64(t170))
								t171 := int32(load32(m.memory[int64(uint32(v1))+56:]))
								store32(m.memory[int64(uint32(v9))+20:], uint32(t171))
								m.memory[int64(uint32(v9))+24] = byte(i32(2))
								goto l114
							}
						l122:
							m.fn436(v24, v25)
							{
								t172 := int32(load32(m.memory[int64(uint32(v0))+80:]))
								v9 = t172
								t173 := int32(load32(m.memory[int64(uint32(v0))+72:]))
								if v9 != t173 {
									goto l128
								}
								m.fn310(v24)
							}
						l128:
							store32(m.memory[int64(uint32(v0))+80:], uint32(v9+i32(1)))
							t174 := int32(load32(m.memory[int64(uint32(v0))+76:]))
							v9 = t174 + v9<<5
							store64(m.memory[int64(uint32(v9))+8:], uint64(v17))
							store32(m.memory[int64(uint32(v9))+4:], uint32(v23))
							store32(m.memory[uint32(v9):], uint32(i32(-0x80000000)))
						}
					l114:
						if v15 != v35 {
							goto l129
						}
					}
				l107:
					{
						if v35 == v15 {
							goto l130
						}
						v18 = int32(uint32(v35-v15) >> 4)
						v11 = i32(0)
					l137:
						{
							v19 = v15 + v11<<4
							t175 := int32(load32(m.memory[int64(uint32(v19))+4:]))
							v23 = t175
							{
								t176 := int32(load32(m.memory[int64(uint32(v19))+8:]))
								v10 = t176
								if v10 == 0 {
									goto l131
								}
								v9 = v23
							l132:
								m.fn332(v9)
								v9 = v9 + i32(28)
								v10 = v10 + i32(-1)
								if v10 != 0 {
									goto l132
								}
							}
						l131:
							{
								t177 := int32(load32(m.memory[uint32(v19):]))
								v9 = t177
								if v9 == 0 {
									goto l133
								}
								t178 := int32(load32(m.memory[uint32(v23+i32(-4)):]))
								v10 = t178
								v19 = v10 & i32(-8)
								t179 := v19
								v10 = v10 & i32(3)
								p180 := i32(8)
								if v10 != 0 {
									p180 = i32(4)
								}
								v9 = v9 * i32(28)
								if uint32(t179) < uint32(p180+v9) {
									m.fn7(i32(1274404), i32(46), i32(1274452))
									panic("unreachable")
								}
								if v10 == 0 {
									goto l135
								}
								if uint32(v19) > uint32(v9+i32(39)) {
									m.fn7(i32(1274468), i32(46), i32(1274516))
									panic("unreachable")
								}
							l135:
								m.fn5(v23)
							}
						l133:
							v11 = v11 + i32(1)
							if v11 != v18 {
								goto l137
							}
						}
					l130:
						{
							if v31 == 0 {
								goto l138
							}
							t181 := int32(load32(m.memory[uint32(v32+i32(-4)):]))
							v9 = t181
							v10 = v9 & i32(-8)
							t182 := v10
							v9 = v9 & i32(3)
							p183 := i32(8)
							if v9 != 0 {
								p183 = i32(4)
							}
							v11 = v31 << 4
							if uint32(t182) < uint32(p183|v11) {
								m.fn7(i32(1274404), i32(46), i32(1274452))
								panic("unreachable")
							}
							if v9 == 0 {
								goto l140
							}
							if uint32(v10) > uint32(v11+i32(39)) {
								m.fn7(i32(1274468), i32(46), i32(1274516))
								panic("unreachable")
							}
						l140:
							m.fn5(v32)
						}
					l138:
						{
							t184 := int32(load32(m.memory[int64(uint32(v1))+24:]))
							v9 = t184
							if v9 == 0 {
								goto l142
							}
							t185 := int32(load32(m.memory[uint32(v21+i32(-4)):]))
							v10 = t185
							v11 = v10 & i32(-8)
							t186 := v11
							v10 = v10 & i32(3)
							p187 := i32(8)
							if v10 != 0 {
								p187 = i32(4)
							}
							if uint32(t186) < uint32(p187+v9) {
								m.fn7(i32(1274404), i32(46), i32(1274452))
								panic("unreachable")
							}
							if v10 == 0 {
								goto l144
							}
							if uint32(v11) > uint32(v9+i32(39)) {
								m.fn7(i32(1274468), i32(46), i32(1274516))
								panic("unreachable")
							}
						l144:
							m.fn5(v21)
						}
					l142:
						if v8 != 0 {
							goto l146
						}
						t188 := int32(load32(m.memory[int64(uint32(v1))+12:]))
						v9 = t188
						if v9 == 0 {
							goto l146
						}
						t189 := int32(load32(m.memory[int64(uint32(v1))+16:]))
						v11 = t189
						t190 := int32(load32(m.memory[uint32(v11+i32(-4)):]))
						v10 = t190
						v19 = v10 & i32(-8)
						t191 := v19
						v10 = v10 & i32(3)
						p192 := i32(8)
						if v10 != 0 {
							p192 = i32(4)
						}
						v9 = v9 * i32(28)
						if uint32(t191) < uint32(p192+v9) {
							m.fn7(i32(1274404), i32(46), i32(1274452))
							panic("unreachable")
						}
						if v10 == 0 {
							goto l148
						}
						if uint32(v19) > uint32(v9+i32(39)) {
							m.fn7(i32(1274468), i32(46), i32(1274516))
							panic("unreachable")
						}
					l148:
						m.fn5(v11)
						goto l146
					}
				l146:
					{
						{
							if v16 == 0 {
								goto l150
							}
							t193 := int32(load32(m.memory[uint32(v14+i32(-4)):]))
							v9 = t193
							v10 = v9 & i32(-8)
							t194 := v10
							v9 = v9 & i32(3)
							p195 := i32(8)
							if v9 != 0 {
								p195 = i32(4)
							}
							v11 = v16 * i32(3)
							if uint32(t194) < uint32(p195+v11) {
								m.fn7(i32(1274404), i32(46), i32(1274452))
								panic("unreachable")
							}
							if v9 == 0 {
								goto l152
							}
							if uint32(v10) > uint32(v11+i32(39)) {
								m.fn7(i32(1274468), i32(46), i32(1274516))
								panic("unreachable")
							}
						l152:
							m.fn5(v14)
						}
					l150:
						{
							if v6 == 0 {
								goto l154
							}
							t196 := int32(load32(m.memory[uint32(v5+i32(-4)):]))
							v9 = t196
							v10 = v9 & i32(-8)
							t197 := v10
							v9 = v9 & i32(3)
							p198 := i32(8)
							if v9 != 0 {
								p198 = i32(4)
							}
							v11 = v6 << 3
							if uint32(t197) < uint32(p198+v11) {
								m.fn7(i32(1274404), i32(46), i32(1274452))
								panic("unreachable")
							}
							if v9 == 0 {
								goto l156
							}
							if uint32(v10) > uint32(v11+i32(39)) {
								m.fn7(i32(1274468), i32(46), i32(1274516))
								panic("unreachable")
							}
						l156:
							m.fn5(v5)
						}
					l154:
						{
							if v4 == 0 {
								goto l158
							}
							t199 := int32(load32(m.memory[uint32(v3+i32(-4)):]))
							v9 = t199
							v10 = v9 & i32(-8)
							t200 := v10
							v9 = v9 & i32(3)
							p201 := i32(8)
							if v9 != 0 {
								p201 = i32(4)
							}
							v11 = v4 << 3
							if uint32(t200) < uint32(p201+v11) {
								m.fn7(i32(1274404), i32(46), i32(1274452))
								panic("unreachable")
							}
							if v9 == 0 {
								goto l160
							}
							if uint32(v10) > uint32(v11+i32(39)) {
								m.fn7(i32(1274468), i32(46), i32(1274516))
								panic("unreachable")
							}
						l160:
							m.fn5(v3)
						}
					l158:
						if v2 == 0 {
							goto l0
						}
						t202 := int32(load32(m.memory[uint32(v7+i32(-4)):]))
						v9 = t202
						v10 = v9 & i32(-8)
						t203 := v10
						v9 = v9 & i32(3)
						p204 := i32(8)
						if v9 != 0 {
							p204 = i32(4)
						}
						if uint32(t203) < uint32(p204+v2) {
							m.fn7(i32(1274404), i32(46), i32(1274452))
							panic("unreachable")
						}
						if v9 != 0 {
							v3 = v7
							if uint32(v10) <= uint32(v2+i32(39)) {
								goto l12
							}
							m.fn7(i32(1274468), i32(46), i32(1274516))
							panic("unreachable")
						}
						v3 = v7
						goto l12
					}
				}
				goto l95
			}
			{
				if v2 == 0 {
					goto l2
				}
				t8 := int32(load32(m.memory[uint32(v7+i32(-4)):]))
				v9 = t8
				v10 = v9 & i32(-8)
				t9 := v10
				v9 = v9 & i32(3)
				p10 := i32(8)
				if v9 != 0 {
					p10 = i32(4)
				}
				if uint32(t9) < uint32(p10+v2) {
					m.fn7(i32(1274404), i32(46), i32(1274452))
					panic("unreachable")
				}
				if v9 == 0 {
					goto l4
				}
				if uint32(v10) > uint32(v2+i32(39)) {
					m.fn7(i32(1274468), i32(46), i32(1274516))
					panic("unreachable")
				}
			l4:
				m.fn5(v7)
			}
		l2:
			switch v6 + i32(1) {
			case 0:
				goto l0
			default:
				t11 := int32(load32(m.memory[uint32(v5+i32(-4)):]))
				v9 = t11
				v10 = v9 & i32(-8)
				t12 := v10
				v9 = v9 & i32(3)
				p13 := i32(8)
				if v9 != 0 {
					p13 = i32(4)
				}
				v11 = v6 << 3
				if uint32(t12) < uint32(p13+v11) {
					m.fn7(i32(1274404), i32(46), i32(1274452))
					panic("unreachable")
				}
				if v9 == 0 {
					goto l9
				}
				if uint32(v10) > uint32(v11+i32(39)) {
					m.fn7(i32(1274468), i32(46), i32(1274516))
					panic("unreachable")
				}
			l9:
				m.fn5(v5)
				fallthrough
			case 1:
				if v4 == 0 {
					goto l0
				}
				t14 := int32(load32(m.memory[uint32(v3+i32(-4)):]))
				v9 = t14
				v10 = v9 & i32(-8)
				t15 := v10
				v9 = v9 & i32(3)
				p16 := i32(8)
				if v9 != 0 {
					p16 = i32(4)
				}
				v11 = v4 << 3
				if uint32(t15) < uint32(p16+v11) {
					m.fn7(i32(1274404), i32(46), i32(1274452))
					panic("unreachable")
				}
				if v9 == 0 {
					goto l12
				}
				if uint32(v10) > uint32(v11+i32(39)) {
					m.fn7(i32(1274468), i32(46), i32(1274516))
					panic("unreachable")
				}
				goto l12
			}
		}
	l12:
		m.fn5(v3)
	}
l0:
	m.g0 = v1 + i32(80)
}
func (m *Module) fn436(v0, v1 int32) {
	var v2, v3 int32
	var v4 int64
	var v5, v6, v7 int32
	t0 := m.g0
	v2 = t0 - i32(32)
	m.g0 = v2
	t1 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	v3 = t1
	store32(m.memory[int64(uint32(v1))+8:], uint32(i32(0)))
	t2 := int64(load64(m.memory[uint32(v1):]))
	v4 = t2
	store64(m.memory[uint32(v1):], uint64(i64(0x800000000)))
	store32(m.memory[int64(uint32(v2))+16:], uint32(v3))
	store64(m.memory[int64(uint32(v2))+8:], uint64(v4))
	{
		if v3 != 0 {
			goto l0
		}
		m.fn438(v2 + i32(8))
		goto l1
	l0:
		m.fn562(v2+i32(20), v2+i32(8))
		t3 := int32(load32(m.memory[int64(uint32(v2))+20:]))
		v3 = t3
		t4 := int32(load32(m.memory[int64(uint32(v2))+24:]))
		v5 = t4
		{
			{
				t5 := int32(load32(m.memory[int64(uint32(v2))+28:]))
				v1 = t5
				t6 := int32(load32(m.memory[uint32(v0):]))
				t7 := int32(load32(m.memory[int64(uint32(v0))+8:]))
				t8 := v1
				v6 = t7
				if uint32(t8) <= uint32(t6-v6) {
					goto l2
				}
				m.fn197(v0, v6, v1, i32(8), i32(32))
				t9 := int32(load32(m.memory[int64(uint32(v0))+8:]))
				v6 = t9
				goto l3
			}
		l2:
			if v1 == 0 {
				goto l4
			}
		l3:
			v7 = v1 << 5
			if v7 == 0 {
				goto l4
			}
			t10 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			memory_copy(m.memory, uint32(t10+v6<<5), uint32(v5), uint32(v7))
		}
	l4:
		store32(m.memory[int64(uint32(v0))+8:], uint32(v6+v1))
		if v3 == 0 {
			goto l1
		}
		t11 := int32(load32(m.memory[uint32(v5+i32(-4)):]))
		v1 = t11
		v0 = v1 & i32(-8)
		t12 := v0
		v1 = v1 & i32(3)
		p13 := i32(8)
		if v1 != 0 {
			p13 = i32(4)
		}
		v3 = v3 << 5
		if uint32(t12) < uint32(p13|v3) {
			m.fn7(i32(1274404), i32(46), i32(1274452))
			panic("unreachable")
		}
		if v1 == 0 {
			goto l6
		}
		if uint32(v0) > uint32(v3+i32(39)) {
			m.fn7(i32(1274468), i32(46), i32(1274516))
			panic("unreachable")
		}
	l6:
		m.fn5(v5)
	}
l1:
	m.g0 = v2 + i32(32)
}
func (m *Module) fn437(v0, v1, v2, v3 int32) {
	var v4, v5, v6, v7, v8, v9, v10, v11, v12, v13 int32
	t0 := m.g0
	v4 = t0 - i32(48)
	m.g0 = v4
	{
		switch v1&i32(0xffff) + i32(-3999) {
		case 2:
			t38 := int32(load32(m.memory[int64(uint32(v0))+16:]))
			if t38 == i32(-1) {
				goto l3
			}
			v5 = i32(0)
			v9 = i32(0)
			{
				t39 := int32(load32(m.memory[int64(uint32(v0))+24:]))
				v10 = t39
				if v10 == 0 {
					goto l42
				}
				t40 := int32(load32(m.memory[int64(uint32(v0))+20:]))
				v1 = t40
				v7 = v1 + v10
				v9 = i32(0)
			l47:
				{
					{
						{
							t41 := int32(int8(m.memory[uint32(v1)]))
							v10 = t41
							if v10 <= i32(-1) {
								goto l43
							}
							v1 = v1 + i32(1)
							v10 = v10 & i32(255)
							goto l44
						}
					l43:
						t42 := int32(m.memory[int64(uint32(v1))+1])
						v6 = t42 & i32(63)
						v11 = v10 & i32(31)
						if uint32(v10) > uint32(i32(-33)) {
							goto l45
						}
						v10 = v11<<6 | v6
						v1 = v1 + i32(2)
						goto l44
					l45:
						t43 := int32(m.memory[int64(uint32(v1))+2])
						v6 = v6<<6 | t43&i32(63)
						if uint32(v10) >= uint32(i32(-16)) {
							goto l46
						}
						v10 = v6 | v11<<12
						v1 = v1 + i32(3)
						goto l44
					l46:
						t44 := int32(m.memory[int64(uint32(v1))+3])
						v10 = v6<<6 | t44&i32(63) | v11<<18&i32(0x1c0000)
						v1 = v1 + i32(4)
					}
				l44:
					p45 := i32(2)
					if uint32(v10) < uint32(i32(65536)) {
						p45 = i32(1)
					}
					v9 = p45 + v9
					if v1 != v7 {
						goto l47
					}
				}
			}
		l42:
			store64(m.memory[int64(uint32(v4))+40:], uint64(i64(4)))
			store64(m.memory[int64(uint32(v4))+32:], uint64(i64(0)))
			store64(m.memory[int64(uint32(v4))+24:], uint64(i64(0x400000000)))
			v7 = i32(6)
			v1 = i32(1)
			v13 = i32(4)
			v11 = i32(0)
		l51:
			{
				if uint32(v3) < uint32(v5) {
					goto l48
				}
				if uint32(v3-v5) < uint32(i32(4)) {
					goto l48
				}
				t46 := v3
				v6 = v5 + i32(4)
				if uint32(t46-v6) < uint32(i32(2)) {
					goto l48
				}
				t47 := int32(load32(m.memory[uint32(v2+v5):]))
				v10 = t47
				t48 := int32(load16(m.memory[uint32(v2+v6):]))
				v8 = t48
				m.fn433(v4, v2, v3, v5+i32(6))
				t49 := int32(m.memory[uint32(v4)])
				v12 = t49
				if v12 == i32(255) {
					goto l49
				}
				t50 := int32(load32(m.memory[int64(uint32(v4))+4:]))
				v5 = t50
				{
					t51 := int32(load32(m.memory[int64(uint32(v4))+24:]))
					if v1+i32(-1) != t51 {
						goto l50
					}
					m.fn293(v4 + i32(24))
					t52 := int32(load32(m.memory[int64(uint32(v4))+28:]))
					v13 = t52
				}
			l50:
				v6 = v13 + v7
				m.memory[uint32(v6)] = byte(v12)
				store16(m.memory[uint32(v6+i32(-2)):], uint16(v8))
				store32(m.memory[uint32(v6+i32(-6)):], uint32(v10))
				store32(m.memory[int64(uint32(v4))+32:], uint32(v1))
				if v10 == 0 {
					goto l48
				}
				v7 = v7 + i32(8)
				v1 = v1 + i32(1)
				v11 = v10 + v11
				if uint32(v11) <= uint32(v9) {
					goto l51
				}
			}
		l48:
			v6 = i32(0)
			v12 = v4 + i32(36)
		l55:
			{
				if uint32(v3) < uint32(v5) {
					goto l49
				}
				if uint32(v3-v5) < uint32(i32(4)) {
					goto l49
				}
				t53 := v3
				v1 = v5 + i32(4)
				if uint32(t53-v1) < uint32(i32(4)) {
					goto l49
				}
				t54 := int32(load32(m.memory[uint32(v2+v5):]))
				v10 = t54
				v7 = v5 + i32(8)
				{
					{
						t55 := int32(load32(m.memory[uint32(v2+v1):]))
						v1 = t55
						if v1&i32(0xffff) != 0 {
							goto l52
						}
						v11 = i32(2)
						v8 = i32(2)
						goto l53
					}
				l52:
					if uint32(v3-v7) <= uint32(i32(1)) {
						goto l49
					}
					t56 := int32(m.memory[uint32(v2+v7)])
					v7 = t56
					p57 := i32(2)
					if v1&i32(1) != 0 {
						p57 = v7 & i32(1)
					}
					v11 = p57
					p58 := i32(2)
					if v1&i32(2) != 0 {
						p58 = int32(uint32(v7)>>1) & i32(1)
					}
					v8 = p58
					v7 = v5 + i32(10)
				}
			l53:
				t59 := int32(uint32(v1)>>20)&i32(2) + int32(uint32(v1)>>15)&i32(2) + int32(uint32(v1)>>21)&i32(2) + int32(uint32(v1)>>22)&i32(2)
				v5 = int32(uint32(v1) >> 16)
				v5 = t59 + v5&i32(2) + v5&i32(4) + int32(uint32(v1)>>18)&i32(2) + v7
				if uint32(v5) > uint32(v3) {
					goto l49
				}
				{
					t60 := int32(load32(m.memory[int64(uint32(v4))+44:]))
					v1 = t60
					t61 := int32(load32(m.memory[int64(uint32(v4))+36:]))
					if v1 != t61 {
						goto l54
					}
					m.fn293(v12)
				}
			l54:
				t62 := int32(load32(m.memory[int64(uint32(v4))+40:]))
				v7 = t62 + v1<<3
				m.memory[int64(uint32(v7))+5] = byte(v8)
				m.memory[int64(uint32(v7))+4] = byte(v11)
				store32(m.memory[uint32(v7):], uint32(v10))
				store32(m.memory[int64(uint32(v4))+44:], uint32(v1+i32(1)))
				if v10 == 0 {
					goto l49
				}
				v6 = v10 + v6
				if uint32(v6) <= uint32(v9) {
					goto l55
				}
			}
		l49:
			{
				{
					t63 := int32(load32(m.memory[int64(uint32(v0))+28:]))
					v1 = t63
					if v1 == i32(-1) {
						goto l56
					}
					{
						if v1 == 0 {
							goto l57
						}
						t64 := int32(load32(m.memory[int64(uint32(v0))+32:]))
						v3 = t64
						t65 := int32(load32(m.memory[uint32(v3+i32(-4)):]))
						v5 = t65
						v9 = v5 & i32(-8)
						t66 := v9
						v5 = v5 & i32(3)
						p67 := i32(8)
						if v5 != 0 {
							p67 = i32(4)
						}
						v1 = v1 << 3
						if uint32(t66) < uint32(p67+v1) {
							m.fn7(i32(1274404), i32(46), i32(1274452))
							panic("unreachable")
						}
						if v5 == 0 {
							goto l59
						}
						if uint32(v9) > uint32(v1+i32(39)) {
							m.fn7(i32(1274468), i32(46), i32(1274516))
							panic("unreachable")
						}
					l59:
						m.fn5(v3)
					}
				l57:
					t68 := int32(load32(m.memory[int64(uint32(v0))+40:]))
					v1 = t68
					if v1 == 0 {
						goto l56
					}
					t69 := int32(load32(m.memory[int64(uint32(v0))+44:]))
					v3 = t69
					t70 := int32(load32(m.memory[uint32(v3+i32(-4)):]))
					v5 = t70
					v9 = v5 & i32(-8)
					t71 := v9
					v5 = v5 & i32(3)
					p72 := i32(8)
					if v5 != 0 {
						p72 = i32(4)
					}
					v1 = v1 << 3
					if uint32(t71) < uint32(p72+v1) {
						m.fn7(i32(1274404), i32(46), i32(1274452))
						panic("unreachable")
					}
					if v5 == 0 {
						goto l62
					}
					if uint32(v9) > uint32(v1+i32(39)) {
						m.fn7(i32(1274468), i32(46), i32(1274516))
						panic("unreachable")
					}
				l62:
					m.fn5(v3)
				}
			l56:
				v1 = v0 + i32(28)
				t73 := int64(load64(m.memory[int64(uint32(v4))+40:]))
				store64(m.memory[int64(uint32(v1))+16:], uint64(t73))
				t74 := int64(load64(m.memory[int64(uint32(v4))+32:]))
				store64(m.memory[int64(uint32(v1))+8:], uint64(t74))
				t75 := int64(load64(m.memory[int64(uint32(v4))+24:]))
				store64(m.memory[uint32(v1):], uint64(t75))
				goto l3
			}
		case 9:
			goto l4
		default:
			goto l3
		case 0:
			m.fn435(v0)
			if v3 != 0 {
				goto l5
			}
			v5 = i32(1)
			goto l6
		case 1:
			v6 = i32(0)
			store32(m.memory[int64(uint32(v4))+20:], uint32(i32(0)))
			store64(m.memory[int64(uint32(v4))+12:], uint64(i64(0x100000000)))
			v1 = int32(uint32(v3)>>1)&i32(1) + int32(uint32(v3)>>2)
			if v1 == 0 {
				goto l7
			}
			m.fn197(v4+i32(12), i32(0), v1, i32(1), i32(1))
		l7:
			v7 = v3 & i32(0x7ffffffe)
		l29:
			{
				{
					if v6&i32(1) == 0 {
						goto l8
					}
					v3 = v8
					goto l9
				l8:
					if uint32(v7) < uint32(i32(2)) {
						t10 := int32(load32(m.memory[int64(uint32(v4))+20:]))
						v1 = t10
						t11 := int32(load32(m.memory[int64(uint32(v4))+16:]))
						v3 = t11
						t12 := int32(load32(m.memory[int64(uint32(v4))+12:]))
						v5 = t12
						{
							t13 := int32(load32(m.memory[int64(uint32(v0))+16:]))
							v9 = t13
							if v9 == i32(-1) {
								m.memory[int64(uint32(v0))+52] = byte(i32(1))
								store32(m.memory[int64(uint32(v0))+28:], uint32(i32(-1)))
								store32(m.memory[int64(uint32(v0))+24:], uint32(v1))
								store32(m.memory[int64(uint32(v0))+20:], uint32(v3))
								store32(m.memory[int64(uint32(v0))+16:], uint32(v5))
								goto l3
							}
							{
								{
									t14 := int32(load32(m.memory[int64(uint32(v0))+24:]))
									t15 := v1
									t16 := v9
									v10 = t14
									if uint32(t15) <= uint32(t16-v10) {
										goto l23
									}
									m.fn197(v0+i32(16), v10, v1, i32(1), i32(1))
									t17 := int32(load32(m.memory[int64(uint32(v0))+24:]))
									v10 = t17
									goto l24
								}
							l23:
								if v1 == 0 {
									goto l25
								}
							l24:
								if v1 == 0 {
									goto l25
								}
								t18 := int32(load32(m.memory[int64(uint32(v0))+20:]))
								memory_copy(m.memory, uint32(t18+v10), uint32(v3), uint32(v1))
							}
						l25:
							store32(m.memory[int64(uint32(v0))+24:], uint32(v10+v1))
							if v5 == 0 {
								goto l3
							}
							t19 := int32(load32(m.memory[uint32(v3+i32(-4)):]))
							v1 = t19
							v9 = v1 & i32(-8)
							t20 := v9
							v1 = v1 & i32(3)
							p21 := i32(8)
							if v1 != 0 {
								p21 = i32(4)
							}
							if uint32(t20) < uint32(p21+v5) {
								m.fn7(i32(1274404), i32(46), i32(1274452))
								panic("unreachable")
							}
							if v1 == 0 {
								goto l27
							}
							if uint32(v9) > uint32(v5+i32(39)) {
								m.fn7(i32(1274468), i32(46), i32(1274516))
								panic("unreachable")
							}
						l27:
							m.fn5(v3)
							goto l3
						}
					}
					v7 = v7 + i32(-2)
					t1 := int32(load16(m.memory[uint32(v2):]))
					v3 = t1
					v2 = v2 + i32(2)
				}
			l9:
				{
					{
						{
							if v3&i32(63488) != i32(55296) {
								goto l11
							}
							v1 = i32(65533)
							v9 = i32(0)
							{
								if uint32(v3&i32(0xffff)) > uint32(i32(56319)) {
									goto l12
								}
								if uint32(v7) < uint32(i32(2)) {
									goto l12
								}
								v7 = v7 + i32(-2)
								v5 = v2 + i32(2)
								t2 := int32(m.memory[int64(uint32(v2))+1])
								v10 = t2 << 8
								t3 := int32(m.memory[uint32(v2)])
								v6 = v10 | t3
								{
									if uint32((v10+i32(8192))&i32(0xff00)) >= uint32(i32(64512)) {
										goto l13
									}
									v2 = v5
									v8 = v6
									v9 = i32(1)
									t4 := int32(load32(m.memory[int64(uint32(v4))+20:]))
									v5 = t4
									goto l14
								}
							l13:
								v1 = v3&i32(1023)<<10 | v6&i32(1023) + i32(65536)
								v2 = v5
							}
						l12:
							t5 := int32(load32(m.memory[int64(uint32(v4))+20:]))
							v5 = t5
							goto l14
						}
					l11:
						t6 := int32(load32(m.memory[int64(uint32(v4))+20:]))
						v5 = t6
						v1 = v3 & i32(0xffff)
						if uint32(v1) >= uint32(i32(128)) {
							goto l15
						}
						v9 = i32(1)
						v6 = i32(0)
						v10 = i32(1)
						goto l16
					l15:
						v10 = i32(2)
						v9 = i32(0)
						v6 = i32(0)
						if uint32(v3&i32(0xffff)) < uint32(i32(2048)) {
							goto l16
						}
					}
				l14:
					v6 = v9
					p7 := i32(4)
					if uint32(v1) < uint32(i32(65536)) {
						p7 = i32(3)
					}
					v10 = p7
					v9 = i32(0)
				}
			l16:
				{
					t8 := int32(load32(m.memory[int64(uint32(v4))+12:]))
					if uint32(v10) <= uint32(t8-v5) {
						goto l17
					}
					m.fn197(v4+i32(12), v5, v10, i32(1), i32(1))
				}
			l17:
				t9 := int32(load32(m.memory[int64(uint32(v4))+16:]))
				v3 = t9 + v5
				if v9 != 0 {
					m.memory[uint32(v3)] = byte(v1)
					goto l20
				}
				v9 = v1&i32(63) | i32(-128)
				v11 = int32(uint32(v1) >> 6)
				if uint32(v1) >= uint32(i32(2048)) {
					v12 = int32(uint32(v1) >> 12)
					v11 = v11&i32(63) | i32(-128)
					if uint32(v1) > uint32(i32(0xffff)) {
						m.memory[int64(uint32(v3))+3] = byte(v9)
						m.memory[int64(uint32(v3))+2] = byte(v11)
						m.memory[int64(uint32(v3))+1] = byte(v12&i32(63) | i32(-128))
						m.memory[uint32(v3)] = byte(int32(uint32(v1)>>18) | i32(-16))
						goto l20
					}
					m.memory[int64(uint32(v3))+2] = byte(v9)
					m.memory[int64(uint32(v3))+1] = byte(v11)
					m.memory[uint32(v3)] = byte(v12 | i32(224))
					goto l20
				}
				m.memory[int64(uint32(v3))+1] = byte(v9)
				m.memory[uint32(v3)] = byte(v11 | i32(192))
				goto l20
			}
		l20:
			store32(m.memory[int64(uint32(v4))+20:], uint32(v10+v5))
			goto l29
		}
	l5:
		t22 := int32(m.memory[uint32(v2)])
		v5 = t22
	}
l6:
	{
		t23 := int32(load32(m.memory[int64(uint32(v0))+16:]))
		v1 = t23
		if v1 == i32(-1) {
			goto l30
		}
		{
			if v1 == 0 {
				goto l31
			}
			t24 := int32(load32(m.memory[int64(uint32(v0))+20:]))
			v9 = t24
			t25 := int32(load32(m.memory[uint32(v9+i32(-4)):]))
			v3 = t25
			v10 = v3 & i32(-8)
			t26 := v10
			v3 = v3 & i32(3)
			p27 := i32(8)
			if v3 != 0 {
				p27 = i32(4)
			}
			if uint32(t26) < uint32(p27+v1) {
				m.fn7(i32(1274404), i32(46), i32(1274452))
				panic("unreachable")
			}
			if v3 == 0 {
				goto l33
			}
			if uint32(v10) > uint32(v1+i32(39)) {
				m.fn7(i32(1274468), i32(46), i32(1274516))
				panic("unreachable")
			}
		l33:
			m.fn5(v9)
		}
	l31:
		t28 := int32(load32(m.memory[int64(uint32(v0))+28:]))
		v1 = t28
		if v1 == i32(-1) {
			goto l30
		}
		{
			if v1 == 0 {
				goto l35
			}
			t29 := int32(load32(m.memory[int64(uint32(v0))+32:]))
			v9 = t29
			t30 := int32(load32(m.memory[uint32(v9+i32(-4)):]))
			v3 = t30
			v10 = v3 & i32(-8)
			t31 := v10
			v3 = v3 & i32(3)
			p32 := i32(8)
			if v3 != 0 {
				p32 = i32(4)
			}
			v1 = v1 << 3
			if uint32(t31) < uint32(p32+v1) {
				m.fn7(i32(1274404), i32(46), i32(1274452))
				panic("unreachable")
			}
			if v3 == 0 {
				goto l37
			}
			if uint32(v10) > uint32(v1+i32(39)) {
				m.fn7(i32(1274468), i32(46), i32(1274516))
				panic("unreachable")
			}
		l37:
			m.fn5(v9)
		}
	l35:
		t33 := int32(load32(m.memory[int64(uint32(v0))+40:]))
		v1 = t33
		if v1 == 0 {
			goto l30
		}
		t34 := int32(load32(m.memory[int64(uint32(v0))+44:]))
		v9 = t34
		t35 := int32(load32(m.memory[uint32(v9+i32(-4)):]))
		v3 = t35
		v10 = v3 & i32(-8)
		t36 := v10
		v3 = v3 & i32(3)
		p37 := i32(8)
		if v3 != 0 {
			p37 = i32(4)
		}
		v1 = v1 << 3
		if uint32(t36) < uint32(p37+v1) {
			m.fn7(i32(1274404), i32(46), i32(1274452))
			panic("unreachable")
		}
		if v3 == 0 {
			goto l40
		}
		if uint32(v10) > uint32(v1+i32(39)) {
			m.fn7(i32(1274468), i32(46), i32(1274516))
			panic("unreachable")
		}
	l40:
		m.fn5(v9)
	}
l30:
	m.memory[int64(uint32(v0))+52] = byte(v5)
	store64(m.memory[int64(uint32(v0))+24:], uint64(i64(-0x100000000)))
	store64(m.memory[int64(uint32(v0))+16:], uint64(i64(0x100000000)))
	goto l3
l4:
	v5 = i32(0)
	store32(m.memory[int64(uint32(v4))+20:], uint32(i32(0)))
	store64(m.memory[int64(uint32(v4))+12:], uint64(i64(0x100000000)))
	v9 = i32(1)
	v1 = i32(0)
	{
		if v3 == 0 {
			goto l64
		}
		m.fn197(v4+i32(12), i32(0), v3, i32(1), i32(1))
		t76 := int32(load32(m.memory[int64(uint32(v4))+20:]))
		v1 = t76
	l67:
		{
			{
				t77 := int32(int8(m.memory[uint32(v2)]))
				v5 = t77
				var p78 int32
				if v5 > i32(-1) {
					p78 = 1
				}
				v6 = p78
				p79 := i32(2)
				if v6 != 0 {
					p79 = i32(1)
				}
				v7 = p79
				t80 := int32(load32(m.memory[int64(uint32(v4))+12:]))
				if uint32(v7) <= uint32(t80-v1) {
					goto l65
				}
				m.fn197(v4+i32(12), v1, v7, i32(1), i32(1))
			}
		l65:
			t81 := int32(load32(m.memory[int64(uint32(v4))+16:]))
			v9 = t81
			v10 = v9 + v1
			if v6 != 0 {
				goto l66
			}
			m.memory[int64(uint32(v10))+1] = byte(v5 & i32(191))
			v5 = int32(uint32(v5&i32(192))>>6) | i32(-64)
		l66:
			m.memory[uint32(v10)] = byte(v5)
			t82 := v4
			v1 = v7 + v1
			store32(m.memory[int64(uint32(t82))+20:], uint32(v1))
			v2 = v2 + i32(1)
			v3 = v3 + i32(-1)
			if v3 != 0 {
				goto l67
			}
		}
		t83 := int32(load32(m.memory[int64(uint32(v4))+12:]))
		v5 = t83
	}
l64:
	{
		t84 := int32(load32(m.memory[int64(uint32(v0))+16:]))
		v3 = t84
		if v3 == i32(-1) {
			goto l68
		}
		{
			{
				t85 := int32(load32(m.memory[int64(uint32(v0))+24:]))
				t86 := v1
				t87 := v3
				v10 = t85
				if uint32(t86) <= uint32(t87-v10) {
					goto l69
				}
				m.fn197(v0+i32(16), v10, v1, i32(1), i32(1))
				t88 := int32(load32(m.memory[int64(uint32(v0))+24:]))
				v10 = t88
				goto l70
			}
		l69:
			if v1 == 0 {
				goto l71
			}
		l70:
			if v1 == 0 {
				goto l71
			}
			t89 := int32(load32(m.memory[int64(uint32(v0))+20:]))
			memory_copy(m.memory, uint32(t89+v10), uint32(v9), uint32(v1))
		}
	l71:
		store32(m.memory[int64(uint32(v0))+24:], uint32(v10+v1))
		if v5 == 0 {
			goto l3
		}
		t90 := int32(load32(m.memory[uint32(v9+i32(-4)):]))
		v1 = t90
		v3 = v1 & i32(-8)
		t91 := v3
		v1 = v1 & i32(3)
		p92 := i32(8)
		if v1 != 0 {
			p92 = i32(4)
		}
		if uint32(t91) < uint32(p92+v5) {
			m.fn7(i32(1274404), i32(46), i32(1274452))
			panic("unreachable")
		}
		if v1 == 0 {
			goto l73
		}
		if uint32(v3) > uint32(v5+i32(39)) {
			m.fn7(i32(1274468), i32(46), i32(1274516))
			panic("unreachable")
		}
	l73:
		m.fn5(v9)
		goto l3
	}
l68:
	m.memory[int64(uint32(v0))+52] = byte(i32(1))
	store32(m.memory[int64(uint32(v0))+28:], uint32(i32(-1)))
	store32(m.memory[int64(uint32(v0))+24:], uint32(v1))
	store32(m.memory[int64(uint32(v0))+20:], uint32(v9))
	store32(m.memory[int64(uint32(v0))+16:], uint32(v5))
l3:
	m.g0 = v4 + i32(48)
}
func (m *Module) fn438(v0 int32) {
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
				v4 = v1 + v3*i32(56)
				t2 := int32(load32(m.memory[int64(uint32(v4))+28:]))
				v5 = t2
				if v5 == i32(-1) {
					goto l1
				}
				if v5 == 0 {
					goto l1
				}
				t3 := int32(load32(m.memory[int64(uint32(v4))+32:]))
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
					m.fn7(i32(1274404), i32(46), i32(1274452))
					panic("unreachable")
				}
				if v7 == 0 {
					goto l3
				}
				if uint32(v8) > uint32(v5+i32(39)) {
					m.fn7(i32(1274468), i32(46), i32(1274516))
					panic("unreachable")
				}
			l3:
				m.fn5(v6)
			}
		l1:
			t7 := int32(load32(m.memory[int64(uint32(v4))+44:]))
			v6 = t7
			{
				t8 := int32(load32(m.memory[int64(uint32(v4))+48:]))
				v7 = t8
				if v7 == 0 {
					goto l5
				}
				v5 = v6
			l6:
				m.fn330(v5)
				v5 = v5 + i32(32)
				v7 = v7 + i32(-1)
				if v7 != 0 {
					goto l6
				}
			}
		l5:
			{
				t9 := int32(load32(m.memory[int64(uint32(v4))+40:]))
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
				v5 = v5 << 5
				if uint32(t11) < uint32(p12|v5) {
					m.fn7(i32(1274404), i32(46), i32(1274452))
					panic("unreachable")
				}
				if v7 == 0 {
					goto l9
				}
				if uint32(v4) > uint32(v5+i32(39)) {
					m.fn7(i32(1274468), i32(46), i32(1274516))
					panic("unreachable")
				}
			l9:
				m.fn5(v6)
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
		v5 = v5 * i32(56)
		if uint32(t15) < uint32(p16+v5) {
			m.fn7(i32(1274404), i32(46), i32(1274452))
			panic("unreachable")
		}
		if v7 == 0 {
			goto l14
		}
		if uint32(v4) > uint32(v5+i32(39)) {
			m.fn7(i32(1274468), i32(46), i32(1274516))
			panic("unreachable")
		}
	l14:
		m.fn5(v1)
	}
}
func (m *Module) fn439(v0, v1, v2, v3, v4 int32) {
	var v5, v6 int32
	var v7 int64
	t0 := m.g0
	v5 = t0 - i32(160)
	m.g0 = v5
	v1 = int32(uint32(v1&i32(65520)) >> 4)
	switch v2&i32(0xffff) + i32(-61466) {
	default:
		store32(m.memory[uint32(v0):], uint32(i32(-2)))
		goto l3
	case 3, 4:
		v6 = i32(17)
		switch v1 + i32(-1761) {
		case 1:
			goto l5
		default:
			if v1 != i32(1131) {
				goto l5
			}
			fallthrough
		case 0, 2:
			v6 = i32(33)
		}
	l5:
		{
			if uint32(v4) < uint32(v6) {
				store32(m.memory[uint32(v0):], uint32(i32(-2)))
				goto l3
			}
			store32(m.memory[int64(uint32(v0))+24:], uint32(i32(3)))
			store32(m.memory[uint32(v0):], uint32(i32(-1)))
			store32(m.memory[int64(uint32(v0))+8:], uint32(v4-v6))
			store32(m.memory[int64(uint32(v0))+4:], uint32(v3+v6))
			t1 := v0
			var p2 int32
			if v2&i32(0xffff) == i32(61469) {
				p2 = 1
			}
			v1 = p2
			p3 := i32(1075388)
			if v1 != 0 {
				p3 = i32(1075376)
			}
			store32(m.memory[int64(uint32(t1))+20:], uint32(p3))
			t5 := v0
			p4 := i32(9)
			if v1 != 0 {
				p4 = i32(10)
			}
			store32(m.memory[int64(uint32(t5))+16:], uint32(p4))
			t7 := v0
			p6 := i32(1075379)
			if v1 != 0 {
				p6 = i32(1075366)
			}
			store32(m.memory[int64(uint32(t7))+12:], uint32(p6))
			goto l3
		}
	case 0, 1:
		{
			t9 := v4
			p8 := i32(16)
			if v1 == i32(535) {
				p8 = i32(32)
			}
			p10 := p8
			if v1 == i32(981) {
				p10 = i32(32)
			}
			v1 = p10
			if uint32(t9) < uint32(v1) {
				store32(m.memory[uint32(v0):], uint32(i32(-2)))
				goto l3
			}
			v4 = v4 - v1
			if uint32(v4) > uint32(i32(3)) {
				if uint32(v4) > uint32(i32(32)) {
					if v4 == i32(33) {
						goto l11
					}
					v1 = v3 + v1
					v6 = v1 + i32(34)
					v4 = v4 + i32(-34)
					v3 = i32(-1)
					{
						{
							t11 := int32(m.memory[int64(uint32(v1))+32])
							if t11 != 0 {
								goto l12
							}
							t12 := int32(load32(m.memory[uint32(v1):]))
							v1 = t12
							store32(m.memory[int64(uint32(v5))+12:], uint32(i32(0)))
							store64(m.memory[int64(uint32(v5))+4:], uint64(i64(0x100000000)))
							m.fn563(v5 + i32(148))
							m.fn564(v5+i32(16), v6, v4, v5+i32(148))
							t14 := v5
							p13 := i32(0x8000000)
							if uint32(v1) < uint32(i32(0x8000000)) {
								p13 = v1
							}
							v7 = int64(uint32(p13))
							store64(m.memory[int64(uint32(t14))+136:], uint64(v7))
							store64(m.memory[int64(uint32(v5))+128:], uint64(v7))
							m.fn565(v5+i32(148), v5+i32(16), v5+i32(4))
							t15 := int32(m.memory[int64(uint32(v5))+148])
							v1 = t15
							if v1 != i32(255) {
								t24 := int32(load32(m.memory[int64(uint32(v5))+152:]))
								m.fn567(v1, t24)
								store32(m.memory[uint32(v0):], uint32(i32(-2)))
								m.fn566(v5 + i32(16))
								t25 := int32(load32(m.memory[int64(uint32(v5))+4:]))
								v0 = t25
								if v0 == 0 {
									goto l3
								}
								t26 := int32(load32(m.memory[int64(uint32(v5))+8:]))
								m.fn21(t26, v0, i32(1))
								goto l3
							}
							t16 := int32(load32(m.memory[int64(uint32(v5))+4:]))
							v3 = t16
							t17 := int32(load32(m.memory[int64(uint32(v5))+8:]))
							v6 = t17
							t18 := int32(load32(m.memory[int64(uint32(v5))+12:]))
							v4 = t18
							m.fn566(v5 + i32(16))
						}
					l12:
						store32(m.memory[int64(uint32(v0))+24:], uint32(i32(3)))
						t19 := v0
						var p20 int32
						if v2&i32(0xffff) == i32(61466) {
							p20 = 1
						}
						v1 = p20
						p21 := i32(1075351)
						if v1 != 0 {
							p21 = i32(1075363)
						}
						store32(m.memory[int64(uint32(t19))+20:], uint32(p21))
						store32(m.memory[int64(uint32(v0))+16:], uint32(i32(9)))
						t23 := v0
						p22 := i32(1075342)
						if v1 != 0 {
							p22 = i32(1075354)
						}
						store32(m.memory[int64(uint32(t23))+12:], uint32(p22))
						store32(m.memory[int64(uint32(v0))+8:], uint32(v4))
						store32(m.memory[int64(uint32(v0))+4:], uint32(v6))
						store32(m.memory[uint32(v0):], uint32(v3))
						goto l3
					}
				}
				store32(m.memory[uint32(v0):], uint32(i32(-2)))
				goto l3
			}
			store32(m.memory[uint32(v0):], uint32(i32(-2)))
			goto l3
		}
	l11:
		store32(m.memory[uint32(v0):], uint32(i32(-2)))
	}
l3:
	m.g0 = v5 + i32(160)
}
func (m *Module) fn440(v0, v1, v2, v3, v4, v5 int32) {
	var v6, v7, v8 int32
	var v9 int64
	var v10, v11 int32
	var v12 int64
	var v13, v14 int32
	var v15 int64
	var v16 int32
	var v17 int64
	var v18, v19, v20 int32
	t0 := m.g0
	v6 = t0 - i32(32)
	m.g0 = v6
	{
		{
			{
				t1 := int32(load32(m.memory[int64(uint32(v1))+12:]))
				if t1 == 0 {
					goto l0
				}
				t2 := int64(load64(m.memory[int64(uint32(v1))+16:]))
				t3 := int64(load64(m.memory[int64(uint32(v1))+24:]))
				t4 := int32(load32(m.memory[int64(uint32(v3))+4:]))
				v7 = t4
				t5 := int32(load32(m.memory[int64(uint32(v3))+8:]))
				t6 := v7
				v8 = t5
				t7 := m.fn65(t2, t3, t6, v8)
				v9 = t7
				t8 := int32(load32(m.memory[int64(uint32(v1))+4:]))
				v10 = t8
				v11 = v10 & int32(v9)
				v12 = int64(uint64(v9)>>25) & i64(127) * i64(72340172838076673)
				t9 := int32(load32(m.memory[uint32(v1):]))
				v13 = t9
				v14 = i32(0)
			l5:
				{
					{
						t10 := int64(load64(m.memory[uint32(v13+v11):]))
						v15 = t10
						v9 = v15 ^ v12
						v9 = (v9 ^ i64(-1)) & (v9 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
						if v9 == 0 {
							goto l1
						}
					l4:
						{
							t11 := v8
							v16 = v13 - (int32(uint32(int64(bits.TrailingZeros64(uint64(v9))))>>3)+v11)&v10<<4
							t12 := int32(load32(m.memory[uint32(v16+i32(-8)):]))
							if t11 != t12 {
								goto l2
							}
							t13 := int32(load32(m.memory[uint32(v16+i32(-12)):]))
							t14 := m.fn1909(v7, t13, v8)
							if t14 == 0 {
								store32(m.memory[uint32(v0):], uint32(i32(-1)))
								t16 := int32(load32(m.memory[uint32(v16+i32(-4)):]))
								store32(m.memory[int64(uint32(v0))+4:], uint32(t16))
								goto l6
							}
						}
					l2:
						v9 = (v9 + i64(-1)) & v9
						if !(v9 == 0) {
							goto l4
						}
					}
				l1:
					if !(v15&(v15<<1)&i64(-0x7f7f7f7f7f7f7f80) == 0) {
						goto l0
					}
					t15 := v11
					v14 = v14 + i32(8)
					v11 = (t15 + v14) & v10
					goto l5
				}
			}
		l0:
			t17 := int32(load32(m.memory[int64(uint32(v1))+32:]))
			t18 := v1
			v11 = t17 + v5
			store32(m.memory[int64(uint32(t18))+32:], uint32(v11))
			{
				if uint32(v11) > uint32(i32(0x8000000)) {
					t56 := m.fn11(i32(45))
					v1 = t56
					if v1 == 0 {
						m.fn16(i32(1), i32(45))
						panic("unreachable")
					}
					store32(m.memory[int64(uint32(v0))+20:], uint32(i32(21)))
					store32(m.memory[int64(uint32(v0))+16:], uint32(i32(1072385)))
					store32(m.memory[int64(uint32(v0))+12:], uint32(i32(45)))
					store32(m.memory[int64(uint32(v0))+8:], uint32(v1))
					store64(m.memory[uint32(v0):], uint64(i64(0x2d80000003)))
					t57 := int64(load64(m.memory[int64(uint32(i32(0)))+1072377:]))
					store64(m.memory[int64(uint32(v1))+37:], uint64(t57))
					t58 := int64(load64(m.memory[int64(uint32(i32(0)))+1072372:]))
					store64(m.memory[int64(uint32(v1))+32:], uint64(t58))
					t59 := int64(load64(m.memory[int64(uint32(i32(0)))+1072364:]))
					store64(m.memory[int64(uint32(v1))+24:], uint64(t59))
					t60 := int64(load64(m.memory[int64(uint32(i32(0)))+1072356:]))
					store64(m.memory[int64(uint32(v1))+16:], uint64(t60))
					t61 := int64(load64(m.memory[int64(uint32(i32(0)))+1072348:]))
					store64(m.memory[int64(uint32(v1))+8:], uint64(t61))
					t62 := int64(load64(m.memory[int64(uint32(i32(0)))+1072340:]))
					store64(m.memory[uint32(v1):], uint64(t62))
					goto l6
				}
				t19 := int32(load32(m.memory[int64(uint32(v1))+44:]))
				v8 = t19
				{
					{
						{
							t20 := int32(load32(m.memory[int64(uint32(v3))+8:]))
							v11 = t20
							if v11 != 0 {
								goto l8
							}
							v13 = i32(1)
							goto l9
						}
					l8:
						t21 := int32(load32(m.memory[int64(uint32(v3))+4:]))
						v16 = t21
						t22 := m.fn11(v11)
						v13 = t22
						if v13 == 0 {
							m.fn16(i32(1), v11)
							panic("unreachable")
						}
						if v11 == 0 {
							goto l9
						}
						memory_copy(m.memory, uint32(v13), uint32(v16), uint32(v11))
					}
				l9:
					t23 := int64(load64(m.memory[int64(uint32(v1))+16:]))
					t24 := int64(load64(m.memory[int64(uint32(v1))+24:]))
					t25 := m.fn65(t23, t24, v13, v11)
					v9 = t25
					{
						t26 := int32(load32(m.memory[int64(uint32(v1))+8:]))
						if t26 != 0 {
							goto l11
						}
						_ = m.fn70(v1, v1+i32(16))
					}
				l11:
					t28 := int32(load32(m.memory[int64(uint32(v1))+4:]))
					v14 = t28
					v10 = v14 & int32(v9)
					v17 = int64(uint64(v9) >> 25)
					v12 = v17 & i64(127) * i64(72340172838076673)
					t29 := int32(load32(m.memory[uint32(v1):]))
					v16 = t29
					v18 = i32(0)
					v19 = i32(0)
				l30:
					{
						{
							{
								t30 := int64(load64(m.memory[uint32(v16+v10):]))
								v15 = t30
								v9 = v15 ^ v12
								v9 = (v9 ^ i64(-1)) & (v9 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
								if v9 == 0 {
									goto l12
								}
							l15:
								{
									t31 := v11
									v20 = v16 - (int32(uint32(int64(bits.TrailingZeros64(uint64(v9))))>>3)+v10)&v14<<4
									t32 := int32(load32(m.memory[uint32(v20+i32(-8)):]))
									if t31 != t32 {
										goto l13
									}
									t33 := int32(load32(m.memory[uint32(v20+i32(-12)):]))
									t34 := m.fn1909(v13, t33, v11)
									if t34 == 0 {
										goto l14
									}
								}
							l13:
								v9 = (v9 + i64(-1)) & v9
								if !(v9 == 0) {
									goto l15
								}
							}
						l12:
							v9 = v15 & i64(-0x7f7f7f7f7f7f7f80)
							if v18 == i32(1) {
								goto l16
							}
							if v9 == 0 {
								v18 = i32(0)
								goto l19
							}
							v7 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v9))))>>3) + v10) & v14
						l16:
							if v9&(v15<<1) != i64(0) {
								{
									t35 := int32(int8(m.memory[uint32(v16+v7)]))
									v10 = t35
									if v10 < i32(0) {
										goto l20
									}
									t36 := int64(load64(m.memory[uint32(v16):]))
									t37 := v16
									v7 = int32(uint32(int64(bits.TrailingZeros64(uint64(t36&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
									t38 := int32(m.memory[uint32(t37+v7)])
									v10 = t38
								}
							l20:
								t39 := v16 + v7
								v20 = int32(v17) & i32(127)
								m.memory[uint32(t39)] = byte(v20)
								m.memory[uint32(v16+(v7+i32(-8))&v14+i32(8))] = byte(v20)
								t40 := int32(load32(m.memory[int64(uint32(v1))+8:]))
								store32(m.memory[int64(uint32(v1))+8:], uint32(t40-v10&i32(1)))
								t41 := int32(load32(m.memory[int64(uint32(v1))+12:]))
								store32(m.memory[int64(uint32(v1))+12:], uint32(t41+i32(1)))
								v16 = v16 - v7<<4
								store32(m.memory[uint32(v16+i32(-16)):], uint32(v11))
								store32(m.memory[uint32(v16+i32(-12)):], uint32(v13))
								store32(m.memory[uint32(v16+i32(-8)):], uint32(v11))
								store32(m.memory[uint32(v16+i32(-4)):], uint32(v8))
								goto l21
							}
							v18 = i32(1)
							goto l19
						l14:
							store32(m.memory[uint32(v20+i32(-4)):], uint32(v8))
							if v11 == 0 {
								goto l21
							}
							t42 := int32(load32(m.memory[uint32(v13+i32(-4)):]))
							v16 = t42
							v10 = v16 & i32(-8)
							t43 := v10
							v16 = v16 & i32(3)
							p44 := i32(8)
							if v16 != 0 {
								p44 = i32(4)
							}
							if uint32(t43) < uint32(p44+v11) {
								m.fn7(i32(1274404), i32(46), i32(1274452))
								panic("unreachable")
							}
							if v16 == 0 {
								goto l23
							}
							if uint32(v10) > uint32(v11+i32(39)) {
								m.fn7(i32(1274468), i32(46), i32(1274516))
								panic("unreachable")
							}
						l23:
							m.fn5(v13)
						}
					l21:
						if v5 != 0 {
							t45 := m.fn11(v5)
							v11 = t45
							if v11 == 0 {
								m.fn16(i32(1), v5)
								panic("unreachable")
							}
							if v5 == 0 {
								goto l26
							}
							memory_copy(m.memory, uint32(v11), uint32(v4), uint32(v5))
							goto l26
						}
						v11 = i32(1)
						goto l26
					l26:
						t46 := int32(load32(m.memory[int64(uint32(v2))+8:]))
						store32(m.memory[int64(uint32(v6))+16:], uint32(t46))
						t47 := int64(load64(m.memory[uint32(v2):]))
						store64(m.memory[int64(uint32(v6))+8:], uint64(t47))
						t48 := int64(load64(m.memory[uint32(v3):]))
						store64(m.memory[int64(uint32(v6))+20:], uint64(t48))
						t49 := int32(load32(m.memory[int64(uint32(v3))+8:]))
						store32(m.memory[int64(uint32(v6))+28:], uint32(t49))
						{
							t50 := int32(load32(m.memory[int64(uint32(v1))+44:]))
							v3 = t50
							t51 := int32(load32(m.memory[int64(uint32(v1))+36:]))
							if v3 != t51 {
								goto l28
							}
							m.fn316(v1 + i32(36))
						}
					l28:
						store32(m.memory[int64(uint32(v1))+44:], uint32(v3+i32(1)))
						t52 := int32(load32(m.memory[int64(uint32(v1))+40:]))
						v1 = t52 + v3*i32(40)
						t53 := int64(load64(m.memory[int64(uint32(v6))+8:]))
						store64(m.memory[uint32(v1):], uint64(t53))
						t54 := int64(load64(m.memory[int64(uint32(v6))+16:]))
						store64(m.memory[int64(uint32(v1))+8:], uint64(t54))
						t55 := int64(load64(m.memory[int64(uint32(v6))+24:]))
						store64(m.memory[int64(uint32(v1))+16:], uint64(t55))
						store32(m.memory[int64(uint32(v1))+36:], uint32(v8))
						store32(m.memory[int64(uint32(v1))+32:], uint32(v5))
						store32(m.memory[int64(uint32(v1))+28:], uint32(v11))
						store32(m.memory[int64(uint32(v1))+24:], uint32(v5))
						store32(m.memory[uint32(v0):], uint32(i32(-1)))
						store32(m.memory[int64(uint32(v0))+4:], uint32(v8))
						goto l29
					}
				l19:
					v19 = v19 + i32(8)
					v10 = (v19 + v10) & v14
					goto l30
				}
			}
		}
	l6:
		{
			t63 := int32(load32(m.memory[uint32(v3):]))
			v1 = t63
			if v1 == 0 {
				goto l32
			}
			t64 := int32(load32(m.memory[int64(uint32(v3))+4:]))
			v0 = t64
			t65 := int32(load32(m.memory[uint32(v0+i32(-4)):]))
			v3 = t65
			v5 = v3 & i32(-8)
			t66 := v5
			v3 = v3 & i32(3)
			p67 := i32(8)
			if v3 != 0 {
				p67 = i32(4)
			}
			if uint32(t66) < uint32(p67+v1) {
				m.fn7(i32(1274404), i32(46), i32(1274452))
				panic("unreachable")
			}
			if v3 == 0 {
				goto l34
			}
			if uint32(v5) > uint32(v1+i32(39)) {
				m.fn7(i32(1274468), i32(46), i32(1274516))
				panic("unreachable")
			}
		l34:
			m.fn5(v0)
		}
	l32:
		t68 := int32(load32(m.memory[uint32(v2):]))
		v1 = t68
		if v1 == 0 {
			goto l29
		}
		t69 := int32(load32(m.memory[int64(uint32(v2))+4:]))
		v0 = t69
		t70 := int32(load32(m.memory[uint32(v0+i32(-4)):]))
		v3 = t70
		v5 = v3 & i32(-8)
		t71 := v5
		v3 = v3 & i32(3)
		p72 := i32(8)
		if v3 != 0 {
			p72 = i32(4)
		}
		if uint32(t71) < uint32(p72+v1) {
			m.fn7(i32(1274404), i32(46), i32(1274452))
			panic("unreachable")
		}
		if v3 == 0 {
			goto l37
		}
		if uint32(v5) > uint32(v1+i32(39)) {
			m.fn7(i32(1274468), i32(46), i32(1274516))
			panic("unreachable")
		}
	l37:
		m.fn5(v0)
	}
l29:
	m.g0 = v6 + i32(32)
}
func (m *Module) fn441(v0, v1, v2 int32) int32 {
	var v3, v4, v5 int32
	t0 := m.g0
	v3 = t0 - i32(32)
	m.g0 = v3
	m.fn427(v3+i32(20), v1, v2)
	{
		{
			t1 := int32(load32(m.memory[int64(uint32(v3))+20:]))
			v4 = t1
			if v4 == i32(-1) {
				goto l0
			}
			t2 := int32(load32(m.memory[int64(uint32(v3))+24:]))
			t3 := v3 + i32(8)
			t4 := v0
			v1 = t2
			t5 := int32(load32(m.memory[int64(uint32(v3))+28:]))
			m.fn428(t3, t4, v1, t5)
			t6 := int32(load32(m.memory[int64(uint32(v3))+8:]))
			var p7 int32
			if t6 == i32(1) {
				p7 = 1
			}
			v2 = p7
			if v4 == 0 {
				goto l1
			}
			t8 := int32(load32(m.memory[uint32(v1+i32(-4)):]))
			v0 = t8
			v5 = v0 & i32(-8)
			t9 := v5
			v0 = v0 & i32(3)
			p10 := i32(8)
			if v0 != 0 {
				p10 = i32(4)
			}
			v4 = v4 << 3
			if uint32(t9) < uint32(p10+v4) {
				m.fn7(i32(1274404), i32(46), i32(1274452))
				panic("unreachable")
			}
			if v0 == 0 {
				goto l3
			}
			if uint32(v5) <= uint32(v4+i32(39)) {
				goto l3
			}
			m.fn7(i32(1274468), i32(46), i32(1274516))
			panic("unreachable")
		}
	l0:
		v2 = i32(0)
		t11 := int32(m.memory[int64(uint32(v3))+24])
		if t11 != i32(3) {
			goto l1
		}
		t12 := int32(load32(m.memory[int64(uint32(v3))+28:]))
		v1 = t12
		t13 := int32(load32(m.memory[uint32(v1):]))
		v2 = t13
		{
			t14 := int32(load32(m.memory[uint32(v1+i32(4)):]))
			v4 = t14
			t15 := int32(load32(m.memory[uint32(v4):]))
			v0 = t15
			if v0 == 0 {
				goto l4
			}
			m.t0[uint(v0)].(func(int32))(v2)
		}
	l4:
		{
			t16 := int32(load32(m.memory[int64(uint32(v4))+4:]))
			v4 = t16
			if v4 == 0 {
				goto l5
			}
			t17 := int32(load32(m.memory[uint32(v2+i32(-4)):]))
			v0 = t17
			v5 = v0 & i32(-8)
			t18 := v5
			v0 = v0 & i32(3)
			p19 := i32(8)
			if v0 != 0 {
				p19 = i32(4)
			}
			if uint32(t18) < uint32(p19+v4) {
				m.fn7(i32(1274404), i32(46), i32(1274452))
				panic("unreachable")
			}
			if v0 == 0 {
				goto l7
			}
			if uint32(v5) > uint32(v4+i32(39)) {
				m.fn7(i32(1274468), i32(46), i32(1274516))
				panic("unreachable")
			}
		l7:
			m.fn5(v2)
		}
	l5:
		t20 := int32(load32(m.memory[uint32(v1+i32(-4)):]))
		v2 = t20
		v0 = v2 & i32(-8)
		t21 := v0
		v4 = v2 & i32(3)
		p22 := i32(20)
		if v4 != 0 {
			p22 = i32(16)
		}
		if uint32(t21) < uint32(p22) {
			m.fn7(i32(1274404), i32(46), i32(1274452))
			panic("unreachable")
		}
		v2 = i32(0)
		if v4 == 0 {
			goto l3
		}
		if uint32(v0) >= uint32(i32(52)) {
			m.fn7(i32(1274468), i32(46), i32(1274516))
			panic("unreachable")
		}
	}
l3:
	m.fn5(v1)
l1:
	m.g0 = v3 + i32(32)
	return v2
}
func (m *Module) fn442(v0, v1, v2 int32) {
	var v3, v4, v5, v6, v7 int32
	var v8, v9 int64
	var v10, v11, v12 int32
	var v13 int64
	t0 := m.g0
	v3 = t0 - i32(32)
	m.g0 = v3
	{
		if v2 != 0 {
			goto l0
		}
		v4 = i32(0)
		goto l1
	l0:
		v5 = v2 * i32(44)
		v6 = v5
		v2 = v1
	l6:
		{
			t1 := int32(load32(m.memory[uint32(v2):]))
			if t1 == i32(-1) {
				goto l2
			}
			t2 := int32(load32(m.memory[uint32(v2+i32(8)):]))
			if t2 != i32(6) {
				goto l2
			}
			t3 := int32(load32(m.memory[uint32(v2+i32(4)):]))
			v7 = t3
			t4 := int32(load32(m.memory[uint32(v7):]))
			t5 := int32(load16(m.memory[uint32(v7+i32(4)):]))
			if t4^i32(1867412834)|(t5^i32(25966)) != 0 {
				goto l2
			}
			t6 := int32(load32(m.memory[uint32(v2+i32(36)):]))
			v7 = t6
			if v7 == 0 {
				goto l2
			}
			t7 := int32(load32(m.memory[uint32(v2+i32(40)):]))
			if t7 != i32(53) {
				goto l2
			}
			v8 = i64(0x687474703a2f2f73)
			{
				{
					t8 := int64(load64(m.memory[int64(uint32(v7))+8:]))
					v9 = t8
					v9 = v9<<56 | v9&i64(0xff00)<<40 | (v9&i64(0xff0000)<<24 | v9&i64(0xff000000)<<8) | (int64(uint64(v9)>>8)&i64(0xff000000) | int64(uint64(v9)>>24)&i64(0xff0000) | (int64(uint64(v9)>>40)&i64(0xff00) | int64(uint64(v9)>>56)))
					if v9 != i64(0x687474703a2f2f73) {
						goto l3
					}
					v8 = i64(7163086727793553007)
					t9 := int64(load64(m.memory[uint32(v7+i32(16)):]))
					v9 = t9
					v9 = v9<<56 | v9&i64(0xff00)<<40 | (v9&i64(0xff0000)<<24 | v9&i64(0xff000000)<<8) | (int64(uint64(v9)>>8)&i64(0xff000000) | int64(uint64(v9)>>24)&i64(0xff0000) | (int64(uint64(v9)>>40)&i64(0xff00) | int64(uint64(v9)>>56)))
					if v9 != i64(7163086727793553007) {
						goto l3
					}
					v8 = i64(8099000968406656623)
					t10 := int64(load64(m.memory[uint32(v7+i32(24)):]))
					v9 = t10
					v9 = v9<<56 | v9&i64(0xff00)<<40 | (v9&i64(0xff0000)<<24 | v9&i64(0xff000000)<<8) | (int64(uint64(v9)>>8)&i64(0xff000000) | int64(uint64(v9)>>24)&i64(0xff0000) | (int64(uint64(v9)>>40)&i64(0xff00) | int64(uint64(v9)>>56)))
					if v9 != i64(8099000968406656623) {
						goto l3
					}
					v8 = i64(8245353645561769842)
					t11 := int64(load64(m.memory[uint32(v7+i32(32)):]))
					v9 = t11
					v9 = v9<<56 | v9&i64(0xff00)<<40 | (v9&i64(0xff0000)<<24 | v9&i64(0xff000000)<<8) | (int64(uint64(v9)>>8)&i64(0xff000000) | int64(uint64(v9)>>24)&i64(0xff0000) | (int64(uint64(v9)>>40)&i64(0xff00) | int64(uint64(v9)>>56)))
					if v9 != i64(8245353645561769842) {
						goto l3
					}
					v8 = i64(7435271952236243310)
					t12 := int64(load64(m.memory[uint32(v7+i32(40)):]))
					v9 = t12
					v9 = v9<<56 | v9&i64(0xff00)<<40 | (v9&i64(0xff0000)<<24 | v9&i64(0xff000000)<<8) | (int64(uint64(v9)>>8)&i64(0xff000000) | int64(uint64(v9)>>24)&i64(0xff0000) | (int64(uint64(v9)>>40)&i64(0xff00) | int64(uint64(v9)>>56)))
					if v9 != i64(7435271952236243310) {
						goto l3
					}
					v8 = i64(0x676d6c2f32303036)
					t13 := int64(load64(m.memory[uint32(v7+i32(48)):]))
					v9 = t13
					v9 = v9<<56 | v9&i64(0xff00)<<40 | (v9&i64(0xff0000)<<24 | v9&i64(0xff000000)<<8) | (int64(uint64(v9)>>8)&i64(0xff000000) | int64(uint64(v9)>>24)&i64(0xff0000) | (int64(uint64(v9)>>40)&i64(0xff00) | int64(uint64(v9)>>56)))
					if v9 != i64(0x676d6c2f32303036) {
						goto l3
					}
					v8 = i64(3472334890029115758)
					v4 = i32(0)
					t14 := int64(load64(m.memory[uint32(v7+i32(53)):]))
					v9 = t14
					v9 = v9<<56 | v9&i64(0xff00)<<40 | (v9&i64(0xff0000)<<24 | v9&i64(0xff000000)<<8) | (int64(uint64(v9)>>8)&i64(0xff000000) | int64(uint64(v9)>>24)&i64(0xff0000) | (int64(uint64(v9)>>40)&i64(0xff00) | int64(uint64(v9)>>56)))
					if v9 == i64(3472334890029115758) {
						goto l4
					}
				}
			l3:
				p15 := i32(1)
				if uint64(v9) < uint64(v8) {
					p15 = i32(-1)
				}
				v4 = p15
			}
		l4:
			if v4 != 0 {
				goto l2
			}
			v4 = i32(1)
			goto l32
		}
	l2:
		v2 = v2 + i32(44)
		v6 = v6 + i32(-44)
		if v6 != 0 {
			goto l6
		}
		v6 = v5
		v2 = v1
	l11:
		{
			t16 := int32(load32(m.memory[uint32(v2):]))
			if t16 == i32(-1) {
				goto l7
			}
			t17 := int32(load32(m.memory[uint32(v2+i32(8)):]))
			if t17 != i32(9) {
				goto l7
			}
			t18 := int32(load32(m.memory[uint32(v2+i32(4)):]))
			v7 = t18
			t19 := int64(load64(m.memory[uint32(v7):]))
			t20 := int64(m.memory[uint32(v7+i32(8))])
			if t19^i64(8452816096595113314)|(t20^i64(109)) != i64(0) {
				goto l7
			}
			t21 := int32(load32(m.memory[uint32(v2+i32(36)):]))
			v7 = t21
			if v7 == 0 {
				goto l7
			}
			t22 := int32(load32(m.memory[uint32(v2+i32(40)):]))
			if t22 != i32(53) {
				goto l7
			}
			v8 = i64(0x687474703a2f2f73)
			{
				{
					t23 := int64(load64(m.memory[int64(uint32(v7))+8:]))
					v9 = t23
					v9 = v9<<56 | v9&i64(0xff00)<<40 | (v9&i64(0xff0000)<<24 | v9&i64(0xff000000)<<8) | (int64(uint64(v9)>>8)&i64(0xff000000) | int64(uint64(v9)>>24)&i64(0xff0000) | (int64(uint64(v9)>>40)&i64(0xff00) | int64(uint64(v9)>>56)))
					if v9 != i64(0x687474703a2f2f73) {
						goto l8
					}
					v8 = i64(7163086727793553007)
					t24 := int64(load64(m.memory[uint32(v7+i32(16)):]))
					v9 = t24
					v9 = v9<<56 | v9&i64(0xff00)<<40 | (v9&i64(0xff0000)<<24 | v9&i64(0xff000000)<<8) | (int64(uint64(v9)>>8)&i64(0xff000000) | int64(uint64(v9)>>24)&i64(0xff0000) | (int64(uint64(v9)>>40)&i64(0xff00) | int64(uint64(v9)>>56)))
					if v9 != i64(7163086727793553007) {
						goto l8
					}
					v8 = i64(8099000968406656623)
					t25 := int64(load64(m.memory[uint32(v7+i32(24)):]))
					v9 = t25
					v9 = v9<<56 | v9&i64(0xff00)<<40 | (v9&i64(0xff0000)<<24 | v9&i64(0xff000000)<<8) | (int64(uint64(v9)>>8)&i64(0xff000000) | int64(uint64(v9)>>24)&i64(0xff0000) | (int64(uint64(v9)>>40)&i64(0xff00) | int64(uint64(v9)>>56)))
					if v9 != i64(8099000968406656623) {
						goto l8
					}
					v8 = i64(8245353645561769842)
					t26 := int64(load64(m.memory[uint32(v7+i32(32)):]))
					v9 = t26
					v9 = v9<<56 | v9&i64(0xff00)<<40 | (v9&i64(0xff0000)<<24 | v9&i64(0xff000000)<<8) | (int64(uint64(v9)>>8)&i64(0xff000000) | int64(uint64(v9)>>24)&i64(0xff0000) | (int64(uint64(v9)>>40)&i64(0xff00) | int64(uint64(v9)>>56)))
					if v9 != i64(8245353645561769842) {
						goto l8
					}
					v8 = i64(7435271952236243310)
					t27 := int64(load64(m.memory[uint32(v7+i32(40)):]))
					v9 = t27
					v9 = v9<<56 | v9&i64(0xff00)<<40 | (v9&i64(0xff0000)<<24 | v9&i64(0xff000000)<<8) | (int64(uint64(v9)>>8)&i64(0xff000000) | int64(uint64(v9)>>24)&i64(0xff0000) | (int64(uint64(v9)>>40)&i64(0xff00) | int64(uint64(v9)>>56)))
					if v9 != i64(7435271952236243310) {
						goto l8
					}
					v8 = i64(0x676d6c2f32303036)
					t28 := int64(load64(m.memory[uint32(v7+i32(48)):]))
					v9 = t28
					v9 = v9<<56 | v9&i64(0xff00)<<40 | (v9&i64(0xff0000)<<24 | v9&i64(0xff000000)<<8) | (int64(uint64(v9)>>8)&i64(0xff000000) | int64(uint64(v9)>>24)&i64(0xff0000) | (int64(uint64(v9)>>40)&i64(0xff00) | int64(uint64(v9)>>56)))
					if v9 != i64(0x676d6c2f32303036) {
						goto l8
					}
					v8 = i64(3472334890029115758)
					v4 = i32(0)
					t29 := int64(load64(m.memory[uint32(v7+i32(53)):]))
					v9 = t29
					v9 = v9<<56 | v9&i64(0xff00)<<40 | (v9&i64(0xff0000)<<24 | v9&i64(0xff000000)<<8) | (int64(uint64(v9)>>8)&i64(0xff000000) | int64(uint64(v9)>>24)&i64(0xff0000) | (int64(uint64(v9)>>40)&i64(0xff00) | int64(uint64(v9)>>56)))
					if v9 == i64(3472334890029115758) {
						goto l9
					}
				}
			l8:
				p30 := i32(1)
				if uint64(v9) < uint64(v8) {
					p30 = i32(-1)
				}
				v4 = p30
			}
		l9:
			if v4 == 0 {
				t46 := int32(load32(m.memory[uint32(v2+i32(16)):]))
				t47 := v3 + i32(8)
				v4 = t46
				t48 := int32(load32(m.memory[uint32(v2+i32(20)):]))
				t49 := v4
				v10 = t48
				m.fn155(t47, t49, v10, i32(1071585), i32(53), i32(1071578), i32(4))
				v7 = i32(1)
				t50 := int32(load32(m.memory[int64(uint32(v3))+8:]))
				v6 = t50
				p51 := i32(1)
				if v6 != 0 {
					p51 = v6
				}
				v2 = p51
				{
					{
						{
							{
								{
									{
										t52 := int32(load32(m.memory[int64(uint32(v3))+12:]))
										p53 := i32(0)
										if v6 != 0 {
											p53 = t52
										}
										v6 = p53
										if uint32(v6) < uint32(i32(7)) {
											if v6 == i32(6) {
												goto l22
											}
											v7 = i32(1)
											if uint32(v6) >= uint32(i32(5)) {
												goto l23
											}
											v11 = i32(0)
											goto l24
										}
										v7 = i32(3)
										{
											t54 := int32(load32(m.memory[uint32(v2):]))
											t55 := t54 ^ i32(1752198241)
											v11 = v2 + i32(3)
											t56 := int32(load32(m.memory[uint32(v11):]))
											if t55|(t56^i32(1665950056)) != 0 {
												t57 := int32(load32(m.memory[uint32(v2):]))
												t58 := int32(load32(m.memory[uint32(v11):]))
												if t57^i32(1752198241)|(t58^i32(1666539880)) == 0 {
													goto l18
												}
												{
													t59 := int32(load32(m.memory[uint32(v2):]))
													t60 := int32(load32(m.memory[uint32(v2+i32(3)):]))
													if t59^i32(1634561906)|(t60^i32(1665953377)) != 0 {
														v7 = i32(1919905121)
														t61 := int32(load32(m.memory[uint32(v2):]))
														v11 = t61
														v11 = i32_rotr(v11&i32(0xff00ff), i32(8)) | i32_rotr(v11, i32(24))&i32(0xff00ff)
														if v11 != i32(1919905121) {
															goto l20
														}
														v7 = i32(1634620771)
														v12 = i32(0)
														t62 := int32(load32(m.memory[uint32(v2+i32(3)):]))
														v11 = t62
														v11 = i32_rotr(v11&i32(0xff00ff), i32(8)) | i32_rotr(v11, i32(24))&i32(0xff00ff)
														if v11 != i32(1634620771) {
															goto l20
														}
														goto l21
													}
													v7 = i32(4)
													goto l18
												}
											}
											v7 = i32(2)
											goto l18
										}
									}
								l20:
									p63 := i32(1)
									if uint32(v7) < uint32(v11) {
										p63 = i32(-1)
									}
									v12 = p63
								}
							l21:
								p64 := i32(5)
								if v12 != 0 {
									p64 = i32(1)
								}
								v7 = p64
							}
						l18:
							if uint32(v6) < uint32(i32(9)) {
								goto l22
							}
							v11 = v2 + v6 + i32(-9)
							t65 := int64(load64(m.memory[uint32(v11):]))
							t66 := int64(m.memory[uint32(v11+i32(8))])
							if !(t65^i64(8389997672730354000)|(t66^i64(104)) == 0) {
								goto l22
							}
							v11 = i32(2)
							goto l24
						}
					l22:
						v11 = v2 + v6 + i32(-6)
						t67 := int32(load32(m.memory[uint32(v11):]))
						t68 := int32(load16(m.memory[uint32(v11+i32(4)):]))
						if t67^i32(1701994832)|(t68^i32(21102)) != 0 {
							goto l23
						}
						v11 = i32(1)
						goto l24
					}
				l23:
					{
						{
							v6 = v2 + v6 + i32(-5)
							t69 := int32(load32(m.memory[uint32(v6):]))
							v2 = t69
							v2 = i32_rotr(v2&i32(0xff00ff), i32(8)) | i32_rotr(v2, i32(24))&i32(0xff00ff)
							if v2 == i32(1349280105) {
								goto l25
							}
							p70 := i32(1)
							if uint32(v2) > uint32(i32(1349280105)) {
								p70 = i32(-1)
							}
							v2 = p70
							goto l26
						}
					l25:
						t71 := int32(m.memory[uint32(v6+i32(4))])
						v2 = i32(110) - t71
					}
				l26:
					p72 := i32(3)
					if v2 != 0 {
						p72 = i32(0)
					}
					v11 = p72
				}
			l24:
				m.fn155(v3, v4, v10, i32(1071585), i32(53), i32(1079681), i32(7))
				v4 = i32(3)
				v13 = i64(1)
				t73 := int32(load32(m.memory[uint32(v3):]))
				v2 = t73
				if v2 == 0 {
					goto l32
				}
				t74 := int32(load32(m.memory[int64(uint32(v3))+4:]))
				m.fn409(v3+i32(16), v2, t74)
				t75 := int32(m.memory[int64(uint32(v3))+16])
				if t75 == i32(1) {
					goto l32
				}
				t76 := int64(load64(m.memory[int64(uint32(v3))+24:]))
				v9 = t76
				p77 := i64(1)
				if v9 > i64(1) {
					p77 = v9
				}
				v9 = p77
				p78 := i64(0x7fff)
				if v9 < i64(0x7fff) {
					p78 = v9
				}
				v13 = p78
				goto l32
			}
		}
	l7:
		v2 = v2 + i32(44)
		v6 = v6 + i32(-44)
		if v6 != 0 {
			goto l11
		}
		v6 = v5
		v2 = v1
	l15:
		{
			t31 := int32(load32(m.memory[uint32(v2):]))
			if t31 == i32(-1) {
				goto l12
			}
			t32 := int32(load32(m.memory[uint32(v2+i32(8)):]))
			if t32 != i32(6) {
				goto l12
			}
			t33 := int32(load32(m.memory[uint32(v2+i32(4)):]))
			v7 = t33
			t34 := int32(load32(m.memory[uint32(v7):]))
			t35 := int32(load16(m.memory[uint32(v7+i32(4)):]))
			if t34^i32(1749251426)|(t35^i32(29281)) != 0 {
				goto l12
			}
			t36 := int32(load32(m.memory[uint32(v2+i32(36)):]))
			v7 = t36
			if v7 == 0 {
				goto l12
			}
			t37 := int32(load32(m.memory[uint32(v2+i32(40)):]))
			if t37 != i32(53) {
				goto l12
			}
			v8 = i64(0x687474703a2f2f73)
			{
				{
					t38 := int64(load64(m.memory[int64(uint32(v7))+8:]))
					v9 = t38
					v9 = v9<<56 | v9&i64(0xff00)<<40 | (v9&i64(0xff0000)<<24 | v9&i64(0xff000000)<<8) | (int64(uint64(v9)>>8)&i64(0xff000000) | int64(uint64(v9)>>24)&i64(0xff0000) | (int64(uint64(v9)>>40)&i64(0xff00) | int64(uint64(v9)>>56)))
					if v9 != i64(0x687474703a2f2f73) {
						goto l13
					}
					v8 = i64(7163086727793553007)
					t39 := int64(load64(m.memory[uint32(v7+i32(16)):]))
					v9 = t39
					v9 = v9<<56 | v9&i64(0xff00)<<40 | (v9&i64(0xff0000)<<24 | v9&i64(0xff000000)<<8) | (int64(uint64(v9)>>8)&i64(0xff000000) | int64(uint64(v9)>>24)&i64(0xff0000) | (int64(uint64(v9)>>40)&i64(0xff00) | int64(uint64(v9)>>56)))
					if v9 != i64(7163086727793553007) {
						goto l13
					}
					v8 = i64(8099000968406656623)
					t40 := int64(load64(m.memory[uint32(v7+i32(24)):]))
					v9 = t40
					v9 = v9<<56 | v9&i64(0xff00)<<40 | (v9&i64(0xff0000)<<24 | v9&i64(0xff000000)<<8) | (int64(uint64(v9)>>8)&i64(0xff000000) | int64(uint64(v9)>>24)&i64(0xff0000) | (int64(uint64(v9)>>40)&i64(0xff00) | int64(uint64(v9)>>56)))
					if v9 != i64(8099000968406656623) {
						goto l13
					}
					v8 = i64(8245353645561769842)
					t41 := int64(load64(m.memory[uint32(v7+i32(32)):]))
					v9 = t41
					v9 = v9<<56 | v9&i64(0xff00)<<40 | (v9&i64(0xff0000)<<24 | v9&i64(0xff000000)<<8) | (int64(uint64(v9)>>8)&i64(0xff000000) | int64(uint64(v9)>>24)&i64(0xff0000) | (int64(uint64(v9)>>40)&i64(0xff00) | int64(uint64(v9)>>56)))
					if v9 != i64(8245353645561769842) {
						goto l13
					}
					v8 = i64(7435271952236243310)
					t42 := int64(load64(m.memory[uint32(v7+i32(40)):]))
					v9 = t42
					v9 = v9<<56 | v9&i64(0xff00)<<40 | (v9&i64(0xff0000)<<24 | v9&i64(0xff000000)<<8) | (int64(uint64(v9)>>8)&i64(0xff000000) | int64(uint64(v9)>>24)&i64(0xff0000) | (int64(uint64(v9)>>40)&i64(0xff00) | int64(uint64(v9)>>56)))
					if v9 != i64(7435271952236243310) {
						goto l13
					}
					v8 = i64(0x676d6c2f32303036)
					t43 := int64(load64(m.memory[uint32(v7+i32(48)):]))
					v9 = t43
					v9 = v9<<56 | v9&i64(0xff00)<<40 | (v9&i64(0xff0000)<<24 | v9&i64(0xff000000)<<8) | (int64(uint64(v9)>>8)&i64(0xff000000) | int64(uint64(v9)>>24)&i64(0xff0000) | (int64(uint64(v9)>>40)&i64(0xff00) | int64(uint64(v9)>>56)))
					if v9 != i64(0x676d6c2f32303036) {
						goto l13
					}
					v8 = i64(3472334890029115758)
					v4 = i32(0)
					t44 := int64(load64(m.memory[uint32(v7+i32(53)):]))
					v9 = t44
					v9 = v9<<56 | v9&i64(0xff00)<<40 | (v9&i64(0xff0000)<<24 | v9&i64(0xff000000)<<8) | (int64(uint64(v9)>>8)&i64(0xff000000) | int64(uint64(v9)>>24)&i64(0xff0000) | (int64(uint64(v9)>>40)&i64(0xff00) | int64(uint64(v9)>>56)))
					if v9 == i64(3472334890029115758) {
						goto l14
					}
				}
			l13:
				p45 := i32(1)
				if uint64(v9) < uint64(v8) {
					p45 = i32(-1)
				}
				v4 = p45
			}
		l14:
			if v4 != 0 {
				goto l12
			}
			v4 = i32(2)
			goto l32
		}
	l12:
		v2 = v2 + i32(44)
		v6 = v6 + i32(-44)
		if v6 != 0 {
			goto l15
		}
		v4 = i32(0)
		goto l32
	l32:
		{
			t79 := int32(load32(m.memory[uint32(v1):]))
			if t79 == i32(-1) {
				goto l28
			}
			t80 := int32(load32(m.memory[uint32(v1+i32(8)):]))
			if t80 != i32(6) {
				goto l28
			}
			t81 := int32(load32(m.memory[uint32(v1+i32(4)):]))
			v2 = t81
			t82 := int32(load32(m.memory[uint32(v2):]))
			t83 := int32(load16(m.memory[uint32(v2+i32(4)):]))
			if t82^i32(0x52666564)|(t83^i32(29264)) != 0 {
				goto l28
			}
			t84 := int32(load32(m.memory[uint32(v1+i32(36)):]))
			v2 = t84
			if v2 == 0 {
				goto l28
			}
			t85 := int32(load32(m.memory[uint32(v1+i32(40)):]))
			if t85 != i32(53) {
				goto l28
			}
			v8 = i64(0x687474703a2f2f73)
			{
				{
					t86 := int64(load64(m.memory[int64(uint32(v2))+8:]))
					v9 = t86
					v9 = v9<<56 | v9&i64(0xff00)<<40 | (v9&i64(0xff0000)<<24 | v9&i64(0xff000000)<<8) | (int64(uint64(v9)>>8)&i64(0xff000000) | int64(uint64(v9)>>24)&i64(0xff0000) | (int64(uint64(v9)>>40)&i64(0xff00) | int64(uint64(v9)>>56)))
					if v9 != i64(0x687474703a2f2f73) {
						goto l29
					}
					v8 = i64(7163086727793553007)
					t87 := int64(load64(m.memory[uint32(v2+i32(16)):]))
					v9 = t87
					v9 = v9<<56 | v9&i64(0xff00)<<40 | (v9&i64(0xff0000)<<24 | v9&i64(0xff000000)<<8) | (int64(uint64(v9)>>8)&i64(0xff000000) | int64(uint64(v9)>>24)&i64(0xff0000) | (int64(uint64(v9)>>40)&i64(0xff00) | int64(uint64(v9)>>56)))
					if v9 != i64(7163086727793553007) {
						goto l29
					}
					v8 = i64(8099000968406656623)
					t88 := int64(load64(m.memory[uint32(v2+i32(24)):]))
					v9 = t88
					v9 = v9<<56 | v9&i64(0xff00)<<40 | (v9&i64(0xff0000)<<24 | v9&i64(0xff000000)<<8) | (int64(uint64(v9)>>8)&i64(0xff000000) | int64(uint64(v9)>>24)&i64(0xff0000) | (int64(uint64(v9)>>40)&i64(0xff00) | int64(uint64(v9)>>56)))
					if v9 != i64(8099000968406656623) {
						goto l29
					}
					v8 = i64(8245353645561769842)
					t89 := int64(load64(m.memory[uint32(v2+i32(32)):]))
					v9 = t89
					v9 = v9<<56 | v9&i64(0xff00)<<40 | (v9&i64(0xff0000)<<24 | v9&i64(0xff000000)<<8) | (int64(uint64(v9)>>8)&i64(0xff000000) | int64(uint64(v9)>>24)&i64(0xff0000) | (int64(uint64(v9)>>40)&i64(0xff00) | int64(uint64(v9)>>56)))
					if v9 != i64(8245353645561769842) {
						goto l29
					}
					v8 = i64(7435271952236243310)
					t90 := int64(load64(m.memory[uint32(v2+i32(40)):]))
					v9 = t90
					v9 = v9<<56 | v9&i64(0xff00)<<40 | (v9&i64(0xff0000)<<24 | v9&i64(0xff000000)<<8) | (int64(uint64(v9)>>8)&i64(0xff000000) | int64(uint64(v9)>>24)&i64(0xff0000) | (int64(uint64(v9)>>40)&i64(0xff00) | int64(uint64(v9)>>56)))
					if v9 != i64(7435271952236243310) {
						goto l29
					}
					v8 = i64(0x676d6c2f32303036)
					t91 := int64(load64(m.memory[uint32(v2+i32(48)):]))
					v9 = t91
					v9 = v9<<56 | v9&i64(0xff00)<<40 | (v9&i64(0xff0000)<<24 | v9&i64(0xff000000)<<8) | (int64(uint64(v9)>>8)&i64(0xff000000) | int64(uint64(v9)>>24)&i64(0xff0000) | (int64(uint64(v9)>>40)&i64(0xff00) | int64(uint64(v9)>>56)))
					if v9 != i64(0x676d6c2f32303036) {
						goto l29
					}
					v8 = i64(3472334890029115758)
					v6 = i32(0)
					t92 := int64(load64(m.memory[uint32(v2+i32(53)):]))
					v9 = t92
					v9 = v9<<56 | v9&i64(0xff00)<<40 | (v9&i64(0xff0000)<<24 | v9&i64(0xff000000)<<8) | (int64(uint64(v9)>>8)&i64(0xff000000) | int64(uint64(v9)>>24)&i64(0xff0000) | (int64(uint64(v9)>>40)&i64(0xff00) | int64(uint64(v9)>>56)))
					if v9 == i64(3472334890029115758) {
						goto l30
					}
				}
			l29:
				p93 := i32(1)
				if uint64(v9) < uint64(v8) {
					p93 = i32(-1)
				}
				v6 = p93
			}
		l30:
			if v6 == 0 {
				goto l31
			}
		}
	l28:
		v1 = v1 + i32(44)
		v5 = v5 + i32(-44)
		if v5 != 0 {
			goto l32
		}
	l1:
		v1 = i32(2)
		v2 = i32(2)
		v5 = i32(2)
		goto l33
	l31:
		t94 := int32(load32(m.memory[uint32(v1+i32(16)):]))
		t95 := int32(load32(m.memory[uint32(v1+i32(20)):]))
		t96 := m.fn710(t94, t95)
		v1 = t96
		v5 = int32(uint32(v1) >> 16)
		v2 = int32(uint32(v1) >> 8)
	}
l33:
	m.memory[int64(uint32(v0))+19] = byte(i32(2))
	m.memory[int64(uint32(v0))+18] = byte(v5)
	m.memory[int64(uint32(v0))+17] = byte(v2)
	m.memory[int64(uint32(v0))+16] = byte(v1)
	store64(m.memory[int64(uint32(v0))+8:], uint64(v13))
	m.memory[int64(uint32(v0))+2] = byte(v11)
	m.memory[int64(uint32(v0))+1] = byte(v7)
	m.memory[uint32(v0)] = byte(v4)
	m.g0 = v3 + i32(32)
}
func (m *Module) fn443(v0, v1 int32) int32 {
	var v2, v3 int32
	var v4 int64
	var v5, v6 int32
	var v7, v8 int64
	var v9, v10, v11 int32
	var v12 int64
	var v13, v14 int32
	t0 := int64(load64(m.memory[int64(uint32(v0))+16:]))
	t1 := int64(load64(m.memory[int64(uint32(v0))+24:]))
	t2 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	v2 = t2
	t3 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	t4 := v2
	v3 = t3
	t5 := m.fn65(t0, t1, t4, v3)
	v4 = t5
	{
		t6 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		if t6 != 0 {
			goto l0
		}
		_ = m.fn79(v0, v0+i32(16))
	}
l0:
	t8 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	v5 = t8
	v6 = v5 & int32(v4)
	v7 = int64(uint64(v4) >> 25)
	v8 = v7 & i64(127) * i64(72340172838076673)
	t9 := int32(load32(m.memory[uint32(v0):]))
	v9 = t9
	v10 = i32(0)
	v11 = i32(0)
	var _ int32
l14:
	{
		t11 := int64(load64(m.memory[uint32(v9+v6):]))
		v12 = t11
		v4 = v12 ^ v8
		v4 = (v4 ^ i64(-1)) & (v4 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
		if v4 == 0 {
			goto l1
		}
	l4:
		{
			t12 := v3
			v13 = v9 + (i32(0)-(int32(uint32(int64(bits.TrailingZeros64(uint64(v4))))>>3)+v6)&v5)*i32(12)
			t13 := int32(load32(m.memory[uint32(v13+i32(-4)):]))
			if t12 != t13 {
				goto l2
			}
			t14 := int32(load32(m.memory[uint32(v13+i32(-8)):]))
			t15 := m.fn1909(v2, t14, v3)
			if t15 == 0 {
				{
					t25 := int32(load32(m.memory[uint32(v1):]))
					v0 = t25
					if v0 == 0 {
						goto l10
					}
					t26 := int32(load32(m.memory[uint32(v2+i32(-4)):]))
					v1 = t26
					v9 = v1 & i32(-8)
					t27 := v9
					v1 = v1 & i32(3)
					p28 := i32(8)
					if v1 != 0 {
						p28 = i32(4)
					}
					if uint32(t27) < uint32(p28+v0) {
						m.fn7(i32(1274404), i32(46), i32(1274452))
						panic("unreachable")
					}
					if v1 == 0 {
						goto l12
					}
					if uint32(v9) > uint32(v0+i32(39)) {
						m.fn7(i32(1274468), i32(46), i32(1274516))
						panic("unreachable")
					}
				l12:
					m.fn5(v2)
				}
			l10:
				return i32(1)
			}
		}
	l2:
		v4 = (v4 + i64(-1)) & v4
		if !(v4 == 0) {
			goto l4
		}
	}
l1:
	v4 = v12 & i64(-0x7f7f7f7f7f7f7f80)
	if v10 == i32(1) {
		goto l5
	}
	if v4 == 0 {
		v10 = i32(0)
		goto l8
	}
	v14 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v4))))>>3) + v6) & v5
l5:
	if v4&(v12<<1) != i64(0) {
		{
			t16 := int32(int8(m.memory[uint32(v9+v14)]))
			v6 = t16
			if v6 < i32(0) {
				goto l9
			}
			t17 := int64(load64(m.memory[uint32(v9):]))
			t18 := v9
			v14 = int32(uint32(int64(bits.TrailingZeros64(uint64(t17&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
			t19 := int32(m.memory[uint32(t18+v14)])
			v6 = t19
		}
	l9:
		t20 := v9 + v14
		v2 = int32(v7) & i32(127)
		m.memory[uint32(t20)] = byte(v2)
		m.memory[uint32(v9+(v14+i32(-8))&v5+i32(8))] = byte(v2)
		t21 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		store32(m.memory[int64(uint32(v0))+8:], uint32(t21-v6&i32(1)))
		t22 := int32(load32(m.memory[int64(uint32(v0))+12:]))
		store32(m.memory[int64(uint32(v0))+12:], uint32(t22+i32(1)))
		v0 = v9 + (i32(0)-v14)*i32(12) + i32(-12)
		t23 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		store32(m.memory[int64(uint32(v0))+8:], uint32(t23))
		t24 := int64(load64(m.memory[uint32(v1):]))
		store64(m.memory[uint32(v0):], uint64(t24))
		return i32(0)
	}
	v10 = i32(1)
	goto l8
l8:
	v11 = v11 + i32(8)
	v6 = (v11 + v6) & v5
	goto l14
}
func (m *Module) fn444(v0 int32) int32 {
	var v1, v2, v3, v4, v5, v6 int32
	t0 := m.g0
	v1 = t0 - i32(16)
	m.g0 = v1
	{
		{
			t1 := int32(load32(m.memory[int64(uint32(v0))+48:]))
			v2 = t1
			if v2 == 0 {
				goto l0
			}
			v2 = v2 + i32(-1)
			t2 := int32(load32(m.memory[int64(uint32(v0))+44:]))
			v3 = t2
			goto l1
		}
	l0:
		{
			t3 := int32(load32(m.memory[int64(uint32(v0))+40:]))
			if t3 != 0 {
				goto l2
			}
			m.fn311(v0 + i32(40))
		}
	l2:
		store32(m.memory[int64(uint32(v0))+48:], uint32(i32(1)))
		v2 = i32(0)
		t4 := int32(load32(m.memory[int64(uint32(v0))+44:]))
		v3 = t4
		store32(m.memory[int64(uint32(v3))+8:], uint32(i32(0)))
		store64(m.memory[uint32(v3):], uint64(i64(0x400000000)))
	}
l1:
	t5 := int32(load32(m.memory[int64(uint32(v3+v2*i32(12)))+8:]))
	m.fn335(v1+i32(4), v0, v2, t5)
	{
		{
			t6 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			v3 = t6
			if v3 != i32(1) {
				goto l3
			}
			t7 := int32(load32(m.memory[int64(uint32(v0))+48:]))
			t8 := v2
			v4 = t7
			if uint32(t8) >= uint32(v4) {
				m.fn33(v2, v4, i32(1073788))
				panic("unreachable")
			}
			t9 := int32(load32(m.memory[int64(uint32(v1))+12:]))
			v5 = t9
			t10 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			v6 = t10
			{
				t11 := int32(load32(m.memory[int64(uint32(v0))+44:]))
				v0 = t11 + v2*i32(12)
				t12 := int32(load32(m.memory[int64(uint32(v0))+8:]))
				v2 = t12
				t13 := int32(load32(m.memory[uint32(v0):]))
				if v2 != t13 {
					goto l5
				}
				m.fn191(v0)
			}
		l5:
			t14 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			v4 = t14 + v2*i32(20)
			store32(m.memory[int64(uint32(v4))+8:], uint32(v5))
			store32(m.memory[int64(uint32(v4))+4:], uint32(v6))
			store32(m.memory[uint32(v4):], uint32(i32(-1)))
			store32(m.memory[int64(uint32(v0))+8:], uint32(v2+i32(1)))
			goto l6
		}
	l3:
		t15 := int32(load32(m.memory[int64(uint32(v0))+48:]))
		t16 := v2
		v4 = t15
		if uint32(t16) >= uint32(v4) {
			m.fn33(v2, v4, i32(1073772))
			panic("unreachable")
		}
		{
			t17 := int32(load32(m.memory[int64(uint32(v0))+44:]))
			v0 = t17 + v2*i32(12)
			t18 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			v2 = t18
			t19 := int32(load32(m.memory[uint32(v0):]))
			if v2 != t19 {
				goto l8
			}
			m.fn191(v0)
		}
	l8:
		t20 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		v4 = t20 + v2*i32(20)
		store32(m.memory[int64(uint32(v4))+16:], uint32(i32(0)))
		store64(m.memory[int64(uint32(v4))+8:], uint64(i64(0)))
		store64(m.memory[uint32(v4):], uint64(i64(0x800000000)))
		store32(m.memory[int64(uint32(v0))+8:], uint32(v2+i32(1)))
	}
l6:
	m.g0 = v1 + i32(16)
	return v3
}
func (m *Module) fn445(v0, v1, v2, v3 int32) int32 {
	var v4, v5, v6, v7, v8, v9, v10, v11 int32
	var v12, v13 int64
	t0 := m.g0
	v4 = t0 - i32(48)
	m.g0 = v4
	v5 = v3 << 3
	v6 = v0 + v1*i32(44)
	v3 = v0
l7:
	if v3 == v6 {
		goto l0
	}
	v7 = i32(0)
	{
		t1 := int32(load32(m.memory[uint32(v3):]))
		var p2 int32
		if t1 == i32(-1) {
			p2 = 1
		}
		v8 = p2
		if v8 != 0 {
			goto l1
		}
		t3 := int32(load32(m.memory[int64(uint32(v3))+8:]))
		if t3 != i32(6) {
			goto l1
		}
		t4 := int32(load32(m.memory[int64(uint32(v3))+4:]))
		v9 = t4
		t5 := int32(load32(m.memory[uint32(v9):]))
		t6 := int32(load16(m.memory[uint32(v9+i32(4)):]))
		if t5^i32(1768908867)|(t6^i32(25955)) != 0 {
			goto l1
		}
		t7 := int32(load32(m.memory[int64(uint32(v3))+36:]))
		v9 = t7
		if v9 == 0 {
			goto l1
		}
		t8 := int32(load32(m.memory[int64(uint32(v3))+40:]))
		if t8 != i32(59) {
			goto l1
		}
		t9 := int64(load64(m.memory[int64(uint32(v9))+8:]))
		t10 := int64(load64(m.memory[uint32(v9+i32(16)):]))
		t11 := int64(load64(m.memory[uint32(v9+i32(24)):]))
		t12 := int64(load64(m.memory[uint32(v9+i32(32)):]))
		t13 := int64(load64(m.memory[uint32(v9+i32(40)):]))
		t14 := int64(load64(m.memory[uint32(v9+i32(48)):]))
		t15 := int64(load64(m.memory[uint32(v9+i32(56)):]))
		t16 := int64(load64(m.memory[uint32(v9+i32(59)):]))
		if t9^i64(8299904566308402280)|(t10^i64(8011467649423075427))|(t11^i64(8027222603262223728)|(t12^i64(8245860516147326322)))|(t13^i64(0x70756b72616d2f67)|(t14^i64(7598805606781117229))|(t15^i64(3616242566693677410)|(t16^i64(3904673869033206889)))) != i64(0) {
			goto l1
		}
		t17 := int32(load32(m.memory[uint32(v3+i32(16)):]))
		t18 := int32(load32(m.memory[uint32(v3+i32(20)):]))
		m.fn155(v4+i32(8), t17, t18, i32(1072856), i32(59), i32(1071727), i32(8))
		{
			t19 := int32(load32(m.memory[int64(uint32(v4))+8:]))
			v9 = t19
			if v9 != 0 {
				goto l2
			}
			v7 = v3
			goto l3
		}
	l2:
		p20 := v3
		if v8 != 0 {
			p20 = i32(0)
		}
		v7 = p20
		t21 := int32(load32(m.memory[int64(uint32(v4))+12:]))
		v8 = t21
		store16(m.memory[int64(uint32(v4))+44:], uint16(i32(1)))
		store32(m.memory[int64(uint32(v4))+40:], uint32(i32(0)))
		store32(m.memory[int64(uint32(v4))+32:], uint32(v9))
		store32(m.memory[int64(uint32(v4))+24:], uint32(v9))
		store32(m.memory[int64(uint32(v4))+16:], uint32(i32(0)))
		store32(m.memory[int64(uint32(v4))+28:], uint32(v8))
		store32(m.memory[int64(uint32(v4))+20:], uint32(v8))
		store32(m.memory[int64(uint32(v4))+36:], uint32(v9+v8))
	l5:
		{
			m.fn247(v4, v4+i32(16))
			t22 := int32(load32(m.memory[uint32(v4):]))
			v10 = t22
			if v10 == 0 {
				goto l1
			}
			t23 := int32(load32(m.memory[int64(uint32(v4))+4:]))
			v11 = t23
			v9 = v5
			v8 = v2
		l6:
			{
				t24 := int32(load32(m.memory[uint32(v8+i32(4)):]))
				if t24 != v11 {
					goto l4
				}
				t25 := int32(load32(m.memory[uint32(v8):]))
				t26 := m.fn1909(t25, v10, v11)
				if t26 == 0 {
					goto l5
				}
			}
		l4:
			v8 = v8 + i32(8)
			v9 = v9 + i32(-8)
			if v9 != 0 {
				goto l6
			}
		}
		v7 = i32(0)
	}
l1:
	v3 = v3 + i32(44)
	if v7 == 0 {
		goto l7
	}
	goto l3
l0:
	v7 = i32(0)
	if v1 == 0 {
		goto l3
	}
	v3 = v1 * i32(44)
l11:
	{
		t27 := int32(load32(m.memory[uint32(v0):]))
		if t27 == i32(-1) {
			goto l8
		}
		t28 := int32(load32(m.memory[uint32(v0+i32(8)):]))
		if t28 != i32(8) {
			goto l8
		}
		t29 := int32(load32(m.memory[uint32(v0+i32(4)):]))
		t30 := int64(load64(m.memory[uint32(t29):]))
		if t30 != i64(7738135660106375494) {
			goto l8
		}
		t31 := int32(load32(m.memory[uint32(v0+i32(36)):]))
		v8 = t31
		if v8 == 0 {
			goto l8
		}
		t32 := int32(load32(m.memory[uint32(v0+i32(40)):]))
		if t32 != i32(59) {
			goto l8
		}
		v12 = i64(0x687474703a2f2f73)
		{
			{
				t33 := int64(load64(m.memory[int64(uint32(v8))+8:]))
				v13 = t33
				v13 = v13<<56 | v13&i64(0xff00)<<40 | (v13&i64(0xff0000)<<24 | v13&i64(0xff000000)<<8) | (int64(uint64(v13)>>8)&i64(0xff000000) | int64(uint64(v13)>>24)&i64(0xff0000) | (int64(uint64(v13)>>40)&i64(0xff00) | int64(uint64(v13)>>56)))
				if v13 != i64(0x687474703a2f2f73) {
					goto l9
				}
				v12 = i64(7163086727793553007)
				t34 := int64(load64(m.memory[uint32(v8+i32(16)):]))
				v13 = t34
				v13 = v13<<56 | v13&i64(0xff00)<<40 | (v13&i64(0xff0000)<<24 | v13&i64(0xff000000)<<8) | (int64(uint64(v13)>>8)&i64(0xff000000) | int64(uint64(v13)>>24)&i64(0xff0000) | (int64(uint64(v13)>>40)&i64(0xff00) | int64(uint64(v13)>>56)))
				if v13 != i64(7163086727793553007) {
					goto l9
				}
				v12 = i64(8099000968406656623)
				t35 := int64(load64(m.memory[uint32(v8+i32(24)):]))
				v13 = t35
				v13 = v13<<56 | v13&i64(0xff00)<<40 | (v13&i64(0xff0000)<<24 | v13&i64(0xff000000)<<8) | (int64(uint64(v13)>>8)&i64(0xff000000) | int64(uint64(v13)>>24)&i64(0xff0000) | (int64(uint64(v13)>>40)&i64(0xff00) | int64(uint64(v13)>>56)))
				if v13 != i64(8099000968406656623) {
					goto l9
				}
				v12 = i64(8245353645561769842)
				t36 := int64(load64(m.memory[uint32(v8+i32(32)):]))
				v13 = t36
				v13 = v13<<56 | v13&i64(0xff00)<<40 | (v13&i64(0xff0000)<<24 | v13&i64(0xff000000)<<8) | (int64(uint64(v13)>>8)&i64(0xff000000) | int64(uint64(v13)>>24)&i64(0xff0000) | (int64(uint64(v13)>>40)&i64(0xff00) | int64(uint64(v13)>>56)))
				if v13 != i64(8245353645561769842) {
					goto l9
				}
				v12 = i64(7435281775110878576)
				t37 := int64(load64(m.memory[uint32(v8+i32(40)):]))
				v13 = t37
				v13 = v13<<56 | v13&i64(0xff00)<<40 | (v13&i64(0xff0000)<<24 | v13&i64(0xff000000)<<8) | (int64(uint64(v13)>>8)&i64(0xff000000) | int64(uint64(v13)>>24)&i64(0xff0000) | (int64(uint64(v13)>>40)&i64(0xff00) | int64(uint64(v13)>>56)))
				if v13 != i64(7435281775110878576) {
					goto l9
				}
				v12 = i64(3270580270228665449)
				t38 := int64(load64(m.memory[uint32(v8+i32(48)):]))
				v13 = t38
				v13 = v13<<56 | v13&i64(0xff00)<<40 | (v13&i64(0xff0000)<<24 | v13&i64(0xff000000)<<8) | (int64(uint64(v13)>>8)&i64(0xff000000) | int64(uint64(v13)>>24)&i64(0xff0000) | (int64(uint64(v13)>>40)&i64(0xff00) | int64(uint64(v13)>>56)))
				if v13 != i64(3270580270228665449) {
					goto l9
				}
				v12 = i64(7091318288453021490)
				t39 := int64(load64(m.memory[uint32(v8+i32(56)):]))
				v13 = t39
				v13 = v13<<56 | v13&i64(0xff00)<<40 | (v13&i64(0xff0000)<<24 | v13&i64(0xff000000)<<8) | (int64(uint64(v13)>>8)&i64(0xff000000) | int64(uint64(v13)>>24)&i64(0xff0000) | (int64(uint64(v13)>>40)&i64(0xff00) | int64(uint64(v13)>>56)))
				if v13 != i64(7091318288453021490) {
					goto l9
				}
				v12 = i64(7598831714893312054)
				v9 = i32(0)
				t40 := int64(load64(m.memory[uint32(v8+i32(59)):]))
				v13 = t40
				v13 = v13<<56 | v13&i64(0xff00)<<40 | (v13&i64(0xff0000)<<24 | v13&i64(0xff000000)<<8) | (int64(uint64(v13)>>8)&i64(0xff000000) | int64(uint64(v13)>>24)&i64(0xff0000) | (int64(uint64(v13)>>40)&i64(0xff00) | int64(uint64(v13)>>56)))
				if v13 == i64(7598831714893312054) {
					goto l10
				}
			}
		l9:
			p41 := i32(1)
			if uint64(v13) < uint64(v12) {
				p41 = i32(-1)
			}
			v9 = p41
		}
	l10:
		if v9 != 0 {
			goto l8
		}
		v7 = v0
		goto l3
	}
l8:
	v0 = v0 + i32(44)
	v3 = v3 + i32(-44)
	if v3 != 0 {
		goto l11
	}
l3:
	m.g0 = v4 + i32(48)
	return v7
}
func (m *Module) fn446(v0, v1, v2 int32) {
	var v3, v4, v5, v6, v7, v8, v9, v10, v11 int32
	t0 := m.g0
	v3 = t0 - i32(16)
	m.g0 = v3
	if v2 <= i32(-1) {
		m.fn15()
		panic("unreachable")
	}
	v4 = i32(1)
	{
		if v2 == 0 {
			goto l1
		}
		t1 := m.fn11(v2)
		v4 = t1
		if v4 == 0 {
			m.fn16(i32(1), v2)
			panic("unreachable")
		}
	}
l1:
	v5 = i32(0)
	store32(m.memory[int64(uint32(v3))+12:], uint32(i32(0)))
	store32(m.memory[int64(uint32(v3))+8:], uint32(v4))
	store32(m.memory[int64(uint32(v3))+4:], uint32(v2))
	v6 = v1 + v2
	v2 = i32(-2)
l12:
	{
		{
			switch v2 + i32(2) {
			case 1:
				goto l4
			default:
				goto l5
			case 0:
				if v1 == v6 {
					goto l4
				}
				{
					t2 := int32(int8(m.memory[uint32(v1)]))
					v2 = t2
					if v2 <= i32(-1) {
						t3 := int32(m.memory[int64(uint32(v1))+1])
						v7 = t3 & i32(63)
						v8 = v2 & i32(31)
						if uint32(v2) > uint32(i32(-33)) {
							t4 := int32(m.memory[int64(uint32(v1))+2])
							v7 = v7<<6 | t4&i32(63)
							if uint32(v2) >= uint32(i32(-16)) {
								goto l8
							}
							v2 = v7 | v8<<12
							v1 = v1 + i32(3)
							goto l5
						}
						v2 = v8<<6 | v7
						v1 = v1 + i32(2)
						goto l5
					}
					v1 = v1 + i32(1)
					v2 = v2 & i32(255)
					goto l5
				}
			}
		l4:
			t5 := int32(load32(m.memory[int64(uint32(v3))+12:]))
			store32(m.memory[int64(uint32(v0))+8:], uint32(t5))
			t6 := int64(load64(m.memory[int64(uint32(v3))+4:]))
			store64(m.memory[uint32(v0):], uint64(t6))
			m.g0 = v3 + i32(16)
			return
		}
	l8:
		t7 := int32(m.memory[int64(uint32(v1))+3])
		v2 = v7<<6 | t7&i32(63) | v8<<18&i32(0x1c0000)
		v1 = v1 + i32(4)
	}
l5:
	v7 = v2
	v2 = i32(-2)
	{
		{
			{
				if v7 > i32(8202) {
					goto l9
				}
				switch v7 + i32(-160) {
				case 0:
					{
						t9 := int32(load32(m.memory[int64(uint32(v3))+4:]))
						if t9 != v5 {
							goto l19
						}
						m.fn197(v3+i32(4), v5, i32(1), i32(1), i32(1))
					}
				l19:
					t10 := int32(load32(m.memory[int64(uint32(v3))+8:]))
					v4 = t10
					m.memory[uint32(v4+v5)] = byte(i32(32))
					goto l20
				case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12:
					goto l11
				case 13:
					goto l12
				default:
					switch v7 + i32(-9) {
					case 0:
						{
							t11 := int32(load32(m.memory[int64(uint32(v3))+4:]))
							if t11 != v5 {
								goto l21
							}
							m.fn197(v3+i32(4), v5, i32(1), i32(1), i32(1))
						}
					l21:
						t12 := int32(load32(m.memory[int64(uint32(v3))+8:]))
						v4 = t12
						m.memory[uint32(v4+v5)] = byte(i32(9))
						goto l20
					case 1:
						goto l15
					case 4:
						{
							if v1 != v6 {
								goto l22
							}
							v2 = i32(-1)
							v1 = v6
							goto l23
						l22:
							{
								{
									t13 := int32(int8(m.memory[uint32(v1)]))
									v2 = t13
									if v2 <= i32(-1) {
										goto l24
									}
									v1 = v1 + i32(1)
									v2 = v2 & i32(255)
									goto l25
								}
							l24:
								t14 := int32(m.memory[int64(uint32(v1))+1])
								v7 = t14 & i32(63)
								v8 = v2 & i32(31)
								if uint32(v2) > uint32(i32(-33)) {
									goto l26
								}
								v2 = v8<<6 | v7
								v1 = v1 + i32(2)
								goto l25
							l26:
								t15 := int32(m.memory[int64(uint32(v1))+2])
								v7 = v7<<6 | t15&i32(63)
								if uint32(v2) >= uint32(i32(-16)) {
									goto l27
								}
								v2 = v7 | v8<<12
								v1 = v1 + i32(3)
								goto l25
							l27:
								t16 := int32(m.memory[int64(uint32(v1))+3])
								v2 = v7<<6 | t16&i32(63) | v8<<18&i32(0x1c0000)
								v1 = v1 + i32(4)
							}
						l25:
							p17 := v2
							if v2 == i32(10) {
								p17 = i32(-2)
							}
							v2 = p17
						}
					l23:
						{
							t18 := int32(load32(m.memory[int64(uint32(v3))+4:]))
							if t18 != v5 {
								goto l28
							}
							m.fn197(v3+i32(4), v5, i32(1), i32(1), i32(1))
						}
					l28:
						t19 := int32(load32(m.memory[int64(uint32(v3))+8:]))
						v4 = t19
						m.memory[uint32(v4+v5)] = byte(i32(32))
						t20 := v3
						v5 = v5 + i32(1)
						store32(m.memory[int64(uint32(t20))+12:], uint32(v5))
						goto l12
					default:
						goto l11
					}
				}
			l9:
				if v7 == i32(8203) {
					goto l12
				}
				if v7 == i32(65279) {
					goto l12
				}
			l11:
				if uint32(v7) < uint32(i32(32)) {
					goto l12
				}
				if uint32(v7+i32(-127)) < uint32(i32(33)) {
					goto l12
				}
				var p8 int32
				if uint32(v7) < uint32(i32(128)) {
					p8 = 1
				}
				v9 = p8
				if v9 == 0 {
					goto l17
				}
				v2 = i32(1)
				goto l18
			}
		l15:
			{
				t21 := int32(load32(m.memory[int64(uint32(v3))+4:]))
				if t21 != v5 {
					goto l29
				}
				m.fn197(v3+i32(4), v5, i32(1), i32(1), i32(1))
				t22 := int32(load32(m.memory[int64(uint32(v3))+8:]))
				v4 = t22
			}
		l29:
			m.memory[uint32(v4+v5)] = byte(i32(32))
		l20:
			v5 = v5 + i32(1)
			goto l30
		l17:
			if uint32(v7) >= uint32(i32(2048)) {
				goto l31
			}
			v2 = i32(2)
			goto l18
		l31:
			p23 := i32(4)
			if uint32(v7) < uint32(i32(65536)) {
				p23 = i32(3)
			}
			v2 = p23
		}
	l18:
		{
			t24 := int32(load32(m.memory[int64(uint32(v3))+4:]))
			if uint32(v2) <= uint32(t24-v5) {
				goto l32
			}
			m.fn197(v3+i32(4), v5, v2, i32(1), i32(1))
		}
	l32:
		t25 := int32(load32(m.memory[int64(uint32(v3))+8:]))
		v4 = t25
		v8 = v4 + v5
		if v9 != 0 {
			goto l33
		}
		v9 = v7&i32(63) | i32(-128)
		v10 = int32(uint32(v7) >> 6)
		if uint32(v7) >= uint32(i32(2048)) {
			v11 = int32(uint32(v7) >> 12)
			v10 = v10&i32(63) | i32(-128)
			if uint32(v7) > uint32(i32(0xffff)) {
				m.memory[int64(uint32(v8))+3] = byte(v9)
				m.memory[int64(uint32(v8))+2] = byte(v10)
				m.memory[int64(uint32(v8))+1] = byte(v11&i32(63) | i32(-128))
				m.memory[uint32(v8)] = byte(int32(uint32(v7)>>18) | i32(-16))
				v5 = v2 + v5
				goto l30
			}
			m.memory[int64(uint32(v8))+2] = byte(v9)
			m.memory[int64(uint32(v8))+1] = byte(v10)
			m.memory[uint32(v8)] = byte(v11 | i32(224))
			v5 = v2 + v5
			goto l30
		}
		m.memory[int64(uint32(v8))+1] = byte(v9)
		m.memory[uint32(v8)] = byte(v10 | i32(192))
		v5 = v2 + v5
		goto l30
	l33:
		m.memory[uint32(v8)] = byte(v7)
		v5 = v2 + v5
	}
l30:
	v2 = i32(-2)
	store32(m.memory[int64(uint32(v3))+12:], uint32(v5))
	goto l12
}
