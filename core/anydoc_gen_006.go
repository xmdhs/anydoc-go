package core

import (
	"math/bits"
)

func (m *Module) fn222(v0, v1, v2, v3 int32) {
	var v4, v5, v6, v7, v8, v9, v10, v11, v12 int32
	t0 := m.g0
	v4 = t0 - i32(80)
	m.g0 = v4
	{
		{
			if uint32(v3) > uint32(i32(3)) {
				goto l0
			}
			v5 = i32(1)
			m.memory[int64(uint32(v0))+8] = byte(i32(1))
			store32(m.memory[int64(uint32(v0))+4:], uint32(i32(-0x7ffffff7)))
			t1 := int64(load64(m.memory[int64(uint32(v1))+16:]))
			store64(m.memory[int64(uint32(v1))+24:], uint64(t1-int64(uint32(v3))))
			goto l1
		}
	l0:
		v5 = v2 + i32(2)
		v6 = v3 + i32(-4)
		{
			if uint32(v3) < uint32(i32(7)) {
				if v6 != 0 {
					goto l4
				}
				v3 = i32(0)
				goto l5
			}
			t2 := int32(load16(m.memory[uint32(v5):]))
			t3 := int32(m.memory[uint32(v5+i32(2))])
			if (t2^i32(28024)|(t3^i32(108)))&i32(0xffff) == 0 {
				{
					if v6 == i32(3) {
						goto l6
					}
					t4 := int32(m.memory[int64(uint32(v2))+5])
					v3 = t4 + i32(-9)
					if uint32(v3) > uint32(i32(23)) {
						goto l4
					}
					if i32_shl(i32(1), v3)&i32(8388627) == 0 {
						goto l4
					}
				}
			l6:
				t5 := int32(load32(m.memory[int64(uint32(v1))+4:]))
				v7 = t5
				{
					t6 := int32(load32(m.memory[uint32(v1):]))
					switch t6 {
					case 1, 3:
						goto l8
					default:
						store32(m.memory[int64(uint32(v4))+52:], uint32(v6))
						store32(m.memory[int64(uint32(v4))+48:], uint32(v5))
						store32(m.memory[int64(uint32(v4))+56:], uint32(v7))
						store64(m.memory[int64(uint32(v4))+24:], uint64(i64(0)))
						store64(m.memory[int64(uint32(v4))+16:], uint64(i64(0x400000000)))
						store64(m.memory[int64(uint32(v4))+8:], uint64(i64(0x300000001)))
						store16(m.memory[int64(uint32(v4))+44:], uint16(i32(0)))
					l20:
						m.fn1823(v4+i32(60), v4+i32(8))
						{
							t7 := int32(load32(m.memory[int64(uint32(v4))+60:]))
							v3 = t7
							if v3 == i32(-3) {
								{
									t10 := int32(load32(m.memory[int64(uint32(v4))+16:]))
									v3 = t10
									if v3 == 0 {
										goto l12
									}
									t11 := int32(load32(m.memory[int64(uint32(v4))+20:]))
									v1 = t11
									t12 := int32(load32(m.memory[uint32(v1+i32(-4)):]))
									v2 = t12
									v8 = v2 & i32(-8)
									t13 := v8
									v2 = v2 & i32(3)
									p14 := i32(8)
									if v2 != 0 {
										p14 = i32(4)
									}
									v3 = v3 << 3
									if uint32(t13) < uint32(p14+v3) {
										m.fn7(i32(1274404), i32(46), i32(1274452))
										panic("unreachable")
									}
									if v2 == 0 {
										goto l14
									}
									if uint32(v8) > uint32(v3+i32(39)) {
										m.fn7(i32(1274468), i32(46), i32(1274516))
										panic("unreachable")
									}
								l14:
									m.fn5(v1)
								}
							l12:
								t15 := int32(load32(m.memory[int64(uint32(v4))+28:]))
								v2 = t15
								if v2 == 0 {
									goto l8
								}
								t16 := int32(load32(m.memory[int64(uint32(v4))+32:]))
								v3 = t16
								if v3 == 0 {
									goto l8
								}
								v1 = v3 << 3
								v3 = v1 + v3 + i32(17)
								if v3 == 0 {
									goto l8
								}
								v1 = v2 - v1
								t17 := int32(load32(m.memory[uint32(v1+i32(-12)):]))
								v2 = t17
								v8 = v2 & i32(-8)
								t18 := v8
								v2 = v2 & i32(3)
								p19 := i32(8)
								if v2 != 0 {
									p19 = i32(4)
								}
								if uint32(t18) < uint32(p19+v3) {
									m.fn7(i32(1274404), i32(46), i32(1274452))
									panic("unreachable")
								}
								if v2 == 0 {
									goto l17
								}
								if uint32(v8) > uint32(v3+i32(39)) {
									m.fn7(i32(1274468), i32(46), i32(1274516))
									panic("unreachable")
								}
							l17:
								m.fn5(v1 + i32(-8))
								goto l8
							}
							t8 := int32(load32(m.memory[int64(uint32(v4))+68:]))
							v8 = t8
							t9 := int32(load32(m.memory[int64(uint32(v4))+64:]))
							v2 = t9
							if v3 != i32(-2) {
								{
									t20 := int32(load32(m.memory[int64(uint32(v4))+76:]))
									if t20 != i32(8) {
										goto l19
									}
									t21 := int32(load32(m.memory[int64(uint32(v4))+72:]))
									t22 := int64(load64(m.memory[uint32(t21):]))
									if t22 == i64(7453010313431182949) {
										goto l11
									}
								}
							l19:
								if v3 < i32(1) {
									goto l20
								}
								t23 := int32(load32(m.memory[uint32(v2+i32(-4)):]))
								v8 = t23
								v9 = v8 & i32(-8)
								t24 := v9
								v8 = v8 & i32(3)
								p25 := i32(8)
								if v8 != 0 {
									p25 = i32(4)
								}
								if uint32(t24) < uint32(p25+v3) {
									m.fn7(i32(1274404), i32(46), i32(1274452))
									panic("unreachable")
								}
								if v8 == 0 {
									goto l22
								}
								if uint32(v9) > uint32(v3+i32(39)) {
									m.fn7(i32(1274468), i32(46), i32(1274516))
									panic("unreachable")
								}
							l22:
								m.fn5(v2)
								goto l20
							}
							v3 = i32(-3)
							goto l11
						}
					l11:
						{
							t26 := int32(load32(m.memory[int64(uint32(v4))+16:]))
							v9 = t26
							if v9 == 0 {
								goto l24
							}
							t27 := int32(load32(m.memory[int64(uint32(v4))+20:]))
							v10 = t27
							t28 := int32(load32(m.memory[uint32(v10+i32(-4)):]))
							v11 = t28
							v12 = v11 & i32(-8)
							t29 := v12
							v11 = v11 & i32(3)
							p30 := i32(8)
							if v11 != 0 {
								p30 = i32(4)
							}
							v9 = v9 << 3
							if uint32(t29) < uint32(p30+v9) {
								m.fn7(i32(1274404), i32(46), i32(1274452))
								panic("unreachable")
							}
							if v11 == 0 {
								goto l26
							}
							if uint32(v12) > uint32(v9+i32(39)) {
								m.fn7(i32(1274468), i32(46), i32(1274516))
								panic("unreachable")
							}
						l26:
							m.fn5(v10)
						}
					l24:
						{
							t31 := int32(load32(m.memory[int64(uint32(v4))+28:]))
							v11 = t31
							if v11 == 0 {
								goto l28
							}
							t32 := int32(load32(m.memory[int64(uint32(v4))+32:]))
							v9 = t32
							if v9 == 0 {
								goto l28
							}
							v10 = v9 << 3
							v9 = v10 + v9 + i32(17)
							if v9 == 0 {
								goto l28
							}
							v10 = v11 - v10
							t33 := int32(load32(m.memory[uint32(v10+i32(-12)):]))
							v11 = t33
							v12 = v11 & i32(-8)
							t34 := v12
							v11 = v11 & i32(3)
							p35 := i32(8)
							if v11 != 0 {
								p35 = i32(4)
							}
							if uint32(t34) < uint32(p35+v9) {
								m.fn7(i32(1274404), i32(46), i32(1274452))
								panic("unreachable")
							}
							if v11 == 0 {
								goto l30
							}
							if uint32(v12) > uint32(v9+i32(39)) {
								m.fn7(i32(1274468), i32(46), i32(1274516))
								panic("unreachable")
							}
						l30:
							m.fn5(v10 + i32(-8))
						}
					l28:
						if v3 < i32(-1) {
							goto l8
						}
						t36 := m.fn211(v2, v8)
						v8 = t36
						if uint32(v3+i32(-1)) > uint32(i32(-3)) {
							goto l32
						}
						m.fn21(v2, v3, i32(1))
					l32:
						if v8 == 0 {
							goto l8
						}
						store32(m.memory[int64(uint32(v1))+4:], uint32(v8))
						store32(m.memory[uint32(v1):], uint32(i32(3)))
					}
				}
			l8:
				store32(m.memory[int64(uint32(v0))+24:], uint32(i32(3)))
				store32(m.memory[int64(uint32(v0))+20:], uint32(v7))
				store32(m.memory[int64(uint32(v0))+16:], uint32(v6))
				store32(m.memory[int64(uint32(v0))+12:], uint32(v5))
				store64(m.memory[int64(uint32(v0))+4:], uint64(i64(-0xfffffffa)))
				v5 = i32(0)
				goto l1
			}
			goto l4
		}
	l4:
		v3 = i32(0)
	l34:
		{
			{
				t37 := int32(m.memory[uint32(v5+v3)])
				v2 = t37 + i32(-9)
				if uint32(v2) > uint32(i32(23)) {
					goto l33
				}
				if i32_shl(i32(1), v2)&i32(8388627) != 0 {
					goto l5
				}
			}
		l33:
			t38 := v6
			v3 = v3 + i32(1)
			if t38 != v3 {
				goto l34
			}
		}
		v3 = v6
	l5:
		store32(m.memory[int64(uint32(v0))+24:], uint32(v3))
		store32(m.memory[int64(uint32(v0))+16:], uint32(v6))
		store32(m.memory[int64(uint32(v0))+12:], uint32(v5))
		store64(m.memory[int64(uint32(v0))+4:], uint64(i64(-0xfffffff9)))
		t39 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		store32(m.memory[int64(uint32(v0))+20:], uint32(t39))
		v5 = i32(0)
	}
l1:
	store32(m.memory[uint32(v0):], uint32(v5))
	m.g0 = v4 + i32(80)
}
func (m *Module) fn223(v0, v1, v2, v3 int32) {
	var v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14 int32
	t0 := m.g0
	v4 = t0 - i32(16)
	m.g0 = v4
	{
		v5 = v3 + i32(-1)
		if uint32(v5) <= uint32(i32(1)) {
			m.fn121(i32(2), v5, v3, i32(1272572))
			panic("unreachable")
		}
		if v3 == 0 {
			m.fn121(i32(2), v5, i32(0), i32(1272572))
			panic("unreachable")
		}
		v6 = v3 + i32(-3)
		t1 := int32(m.memory[int64(uint32(v1))+13])
		if t1 != i32(1) {
			goto l2
		}
		if v6 == 0 {
			goto l2
		}
		v5 = v3 + i32(-2)
	l4:
		{
			t2 := int32(m.memory[uint32(v2+v5)])
			v7 = t2 + i32(-9)
			if uint32(v7) > uint32(i32(23)) {
				goto l3
			}
			if i32_shl(i32(1), v7)&i32(8388627) == 0 {
				goto l3
			}
			v5 = v5 + i32(-1)
			if v5 != i32(1) {
				goto l4
			}
			goto l2
		}
	l3:
		v5 = v5 + i32(-1)
		if uint32(v5) > uint32(v6) {
			m.fn121(i32(0), v5, v6, i32(1272540))
			panic("unreachable")
		}
		v6 = v5
		goto l2
	}
l2:
	v5 = v2 + i32(2)
	t3 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	v8 = t3
	{
		{
			{
				t4 := int32(load32(m.memory[int64(uint32(v1))+52:]))
				v2 = t4
				if v2 != 0 {
					v7 = i32(-1)
					t11 := v1
					v2 = v2 + i32(-1)
					store32(m.memory[int64(uint32(t11))+52:], uint32(v2))
					t12 := int32(load32(m.memory[int64(uint32(v1))+48:]))
					t13 := int32(load32(m.memory[uint32(t12+v2<<2):]))
					v2 = t13
					t14 := int32(load32(m.memory[int64(uint32(v1))+40:]))
					v14 = t14
					t15 := int32(m.memory[int64(uint32(v1))+11])
					if t15 == 0 {
						goto l15
					}
					if uint32(v14) < uint32(v2) {
						m.fn121(v2, v14, v14, i32(1272556))
						panic("unreachable")
					}
					t16 := int32(load32(m.memory[int64(uint32(v1))+36:]))
					v13 = t16 + v2
					{
						t17 := v6
						v9 = v14 - v2
						if t17 != v9 {
							goto l17
						}
						t18 := m.fn1909(v5, v13, v6)
						if t18 == 0 {
							goto l15
						}
					}
				l17:
					m.fn243(v4+i32(4), v8, v13, v9)
					t19 := int32(load32(m.memory[int64(uint32(v4))+4:]))
					v9 = t19
					if v9 == i32(-2) {
						goto l18
					}
					t20 := int32(load32(m.memory[int64(uint32(v4))+12:]))
					v13 = t20
					t21 := int32(load32(m.memory[int64(uint32(v4))+8:]))
					v7 = t21
					if v9 == i32(-1) {
						goto l19
					}
					v10 = v7
					goto l20
				}
				v9 = i32(16)
				v10 = i32(12)
				v11 = i32(8)
				v12 = i32(4)
				{
					t5 := int32(m.memory[int64(uint32(v1))+9])
					if t5 == 0 {
						t6 := int64(load64(m.memory[int64(uint32(v1))+16:]))
						store64(m.memory[int64(uint32(v1))+24:], uint64(t6-int64(uint32(v3))))
						m.fn243(v4+i32(4), v8, v5, v6)
						t7 := int32(load32(m.memory[int64(uint32(v4))+4:]))
						v7 = t7
						if v7 == i32(-2) {
							v5 = i32(1)
							v7 = i32(0)
							v13 = i32(-0x7ffffffc)
							v6 = i32(0)
							v3 = i32(1)
							goto l8
						}
						t8 := int32(load32(m.memory[int64(uint32(v4))+12:]))
						v6 = t8
						t9 := int32(load32(m.memory[int64(uint32(v4))+8:]))
						v2 = t9
						v13 = i32(-0x7ffffffc)
						if v7 == i32(-1) {
							if v6 <= i32(-1) {
								goto l11
							}
							v3 = i32(1)
							if v6 != 0 {
								t10 := m.fn11(v6)
								v5 = t10
								if v5 == 0 {
									m.fn16(i32(1), v6)
									panic("unreachable")
								}
								if v6 == 0 {
									goto l14
								}
								memory_copy(m.memory, uint32(v5), uint32(v2), uint32(v6))
							l14:
								v7 = v6
								goto l8
							}
							v7 = i32(0)
							v5 = i32(1)
							v6 = i32(0)
							goto l8
						}
						v3 = i32(1)
						v5 = v2
						goto l8
					}
					v3 = i32(0)
					v7 = i32(-1)
					v13 = i32(1)
					goto l8
				}
			}
		l19:
			if v13 <= i32(-1) {
				goto l11
			}
			if v13 != 0 {
				goto l21
			}
		l18:
			v10 = i32(1)
			v13 = i32(0)
			v9 = i32(0)
			goto l20
		l21:
			t22 := m.fn11(v13)
			v10 = t22
			if v10 == 0 {
				m.fn16(i32(1), v13)
				panic("unreachable")
			}
			if v13 == 0 {
				goto l23
			}
			memory_copy(m.memory, uint32(v10), uint32(v7), uint32(v13))
		l23:
			v9 = v13
		}
	l20:
		store32(m.memory[int64(uint32(v1))+40:], uint32(v2))
		t23 := int64(load64(m.memory[int64(uint32(v1))+16:]))
		store64(m.memory[int64(uint32(v1))+24:], uint64(t23-int64(uint32(v3))))
		m.fn243(v4+i32(4), v8, v5, v6)
		{
			{
				t24 := int32(load32(m.memory[int64(uint32(v4))+4:]))
				v7 = t24
				if v7 == i32(-2) {
					goto l24
				}
				t25 := int32(load32(m.memory[int64(uint32(v4))+12:]))
				v6 = t25
				t26 := int32(load32(m.memory[int64(uint32(v4))+8:]))
				v2 = t26
				if v7 == i32(-1) {
					goto l25
				}
				v5 = v2
				goto l26
			l25:
				if v6 <= i32(-1) {
					goto l11
				}
				if v6 != 0 {
					goto l27
				}
			}
		l24:
			v5 = i32(1)
			v6 = i32(0)
			v7 = i32(0)
			goto l26
		l27:
			t27 := m.fn11(v6)
			v5 = t27
			if v5 == 0 {
				m.fn16(i32(1), v6)
				panic("unreachable")
			}
			if v6 == 0 {
				goto l29
			}
			memory_copy(m.memory, uint32(v5), uint32(v2), uint32(v6))
		l29:
			v7 = v6
		}
	l26:
		store32(m.memory[int64(uint32(v0))+8:], uint32(v10))
		store32(m.memory[int64(uint32(v0))+4:], uint32(v9))
		v3 = i32(1)
		v9 = i32(24)
		v10 = i32(20)
		v11 = i32(16)
		v12 = i32(12)
		goto l8
	}
l11:
	m.fn15()
	panic("unreachable")
l15:
	v3 = i32(0)
	v9 = i32(16)
	v10 = i32(12)
	v11 = i32(8)
	v13 = i32(1)
	v12 = i32(4)
	if uint32(v14) < uint32(v2) {
		goto l8
	}
	store32(m.memory[int64(uint32(v1))+40:], uint32(v2))
l8:
	store32(m.memory[uint32(v0+v12):], uint32(v13))
	store32(m.memory[uint32(v0+v11):], uint32(v7))
	store32(m.memory[uint32(v0+v10):], uint32(v5))
	store32(m.memory[uint32(v0):], uint32(v3))
	store32(m.memory[uint32(v0+v9):], uint32(v6))
	m.g0 = v4 + i32(16)
}
func (m *Module) fn224(v0, v1, v2, v3, v4 int32) {
	var v5, v6, v7, v8, v9, v10 int32
	{
		{
			{
				t0 := v2 + i32(-9)
				v5 = v2 & i32(255)
				p1 := i32(2)
				if uint32(v5) > uint32(i32(8)) {
					p1 = t0
				}
				switch p1 & i32(255) {
				default:
					v2 = i32(5)
					if uint32(v4) < uint32(i32(9)) {
						goto l3
					}
					t2 := int64(load64(m.memory[uint32(v3):]))
					t3 := int64(m.memory[uint32(v3+i32(8))])
					if !(t2^i64(0x41544144435b213c)|(t3^i64(91)) == 0) {
						goto l3
					}
					if uint32(v4) >= uint32(i32(12)) {
						store64(m.memory[int64(uint32(v0))+4:], uint64(i64(-0xfffffffc)))
						t6 := int32(load32(m.memory[int64(uint32(v1))+4:]))
						store32(m.memory[int64(uint32(v0))+20:], uint32(t6))
						store32(m.memory[int64(uint32(v0))+16:], uint32(v4+i32(-12)))
						store32(m.memory[int64(uint32(v0))+12:], uint32(v3+i32(9)))
						store32(m.memory[uint32(v0):], uint32(i32(0)))
						return
					}
					m.fn121(i32(9), v4+i32(-3), v4, i32(1272588))
					panic("unreachable")
				case 1:
					v2 = i32(3)
					if uint32(v4) < uint32(i32(4)) {
						goto l3
					}
					t4 := int32(load32(m.memory[uint32(v3):]))
					if t4 != i32(757932348) {
						goto l3
					}
					t5 := int32(m.memory[int64(uint32(v1))+10])
					if t5 == 0 {
						goto l5
					}
					if uint32(v4) <= uint32(i32(6)) {
						m.fn121(i32(4), v4+i32(-3), v4, i32(1272636))
						panic("unreachable")
					}
					v6 = v4 + i32(-7)
					if v6 != 0 {
						goto l7
					}
					v4 = i32(0)
					goto l8
				case 2:
					v2 = i32(4)
					if v5 != i32(8) {
						goto l3
					}
					if uint32(v4) > uint32(i32(8)) {
						t18 := int32(m.memory[uint32(v3)])
						v5 = t18
						p19 := i32(0)
						if uint32((v5+i32(-65))&i32(255)) < uint32(i32(26)) {
							p19 = i32(32)
						}
						if (p19|v5)&i32(255) != i32(60) {
							goto l3
						}
						t20 := int32(m.memory[int64(uint32(v3))+1])
						v5 = t20
						p21 := i32(0)
						if uint32((v5+i32(-65))&i32(255)) < uint32(i32(26)) {
							p21 = i32(32)
						}
						if (p21|v5)&i32(255) != i32(33) {
							goto l3
						}
						t22 := int32(m.memory[int64(uint32(v3))+2])
						v5 = t22
						p23 := i32(0)
						if uint32((v5+i32(-65))&i32(255)) < uint32(i32(26)) {
							p23 = i32(32)
						}
						if (p23|v5)&i32(255) != i32(100) {
							goto l3
						}
						t24 := int32(m.memory[int64(uint32(v3))+3])
						v5 = t24
						p25 := i32(0)
						if uint32((v5+i32(-65))&i32(255)) < uint32(i32(26)) {
							p25 = i32(32)
						}
						if (p25|v5)&i32(255) != i32(111) {
							goto l3
						}
						t26 := int32(m.memory[int64(uint32(v3))+4])
						v5 = t26
						p27 := i32(0)
						if uint32((v5+i32(-65))&i32(255)) < uint32(i32(26)) {
							p27 = i32(32)
						}
						if (p27|v5)&i32(255) != i32(99) {
							goto l3
						}
						t28 := int32(m.memory[int64(uint32(v3))+5])
						v5 = t28
						p29 := i32(0)
						if uint32((v5+i32(-65))&i32(255)) < uint32(i32(26)) {
							p29 = i32(32)
						}
						if (p29|v5)&i32(255) != i32(116) {
							goto l3
						}
						t30 := int32(m.memory[int64(uint32(v3))+6])
						v5 = t30
						p31 := i32(0)
						if uint32((v5+i32(-65))&i32(255)) < uint32(i32(26)) {
							p31 = i32(32)
						}
						if (p31|v5)&i32(255) != i32(121) {
							goto l3
						}
						t32 := int32(m.memory[int64(uint32(v3))+7])
						v5 = t32
						p33 := i32(0)
						if uint32((v5+i32(-65))&i32(255)) < uint32(i32(26)) {
							p33 = i32(32)
						}
						if (p33|v5)&i32(255) != i32(112) {
							goto l3
						}
						t34 := int32(m.memory[int64(uint32(v3))+8])
						v5 = t34
						p35 := i32(0)
						if uint32((v5+i32(-65))&i32(255)) < uint32(i32(26)) {
							p35 = i32(32)
						}
						if (p35|v5)&i32(255) != i32(101) {
							goto l3
						}
						v8 = v4 + i32(-1)
						if v4 == i32(9) {
							m.fn121(i32(9), v8, i32(9), i32(1272668))
							panic("unreachable")
						}
						{
							if v4 == i32(10) {
								goto l26
							}
							v7 = v3 + i32(9)
							v6 = v4 + i32(-10)
							v2 = i32(0)
						l28:
							{
								t36 := int32(m.memory[uint32(v7+v2)])
								v5 = t36 + i32(-9)
								if uint32(v5) > uint32(i32(23)) {
									goto l27
								}
								if i32_shl(i32(1), v5)&i32(8388627) == 0 {
									goto l27
								}
								t37 := v6
								v2 = v2 + i32(1)
								if t37 != v2 {
									goto l28
								}
							}
						l26:
							store32(m.memory[int64(uint32(v0))+4:], uint32(i32(-0x7ffffffe)))
							t38 := int64(load64(m.memory[int64(uint32(v1))+16:]))
							store64(m.memory[int64(uint32(v1))+24:], uint64(t38+i64(-1)))
							goto l24
						}
					l27:
						{
							t39 := v8
							v2 = v2 + i32(9)
							if uint32(t39) < uint32(v2) {
								m.fn121(v2, v8, v4, i32(1272652))
								panic("unreachable")
							}
							store64(m.memory[int64(uint32(v0))+4:], uint64(i64(-0xfffffff8)))
							t40 := int32(load32(m.memory[int64(uint32(v1))+4:]))
							store32(m.memory[int64(uint32(v0))+20:], uint32(t40))
							store32(m.memory[int64(uint32(v0))+16:], uint32(v8-v2))
							store32(m.memory[int64(uint32(v0))+12:], uint32(v3+v2))
							store32(m.memory[uint32(v0):], uint32(i32(0)))
							return
						}
					}
					goto l3
				}
			}
		l7:
			v7 = v3 + i32(4)
			v5 = v7 + v6
			v8 = i32(0)
		l22:
			{
				v2 = v7
				if uint32(v6) <= uint32(i32(3)) {
				l18:
					{
						t13 := int32(m.memory[uint32(v2)])
						if t13 == i32(45) {
							goto l15
						}
						v2 = v2 + i32(1)
						if v2 == v5 {
							goto l5
						}
						goto l18
					}
				}
				v2 = v7
				{
					t7 := int32(load32(m.memory[uint32(v7):]))
					v9 = t7
					if (i32(16843008)-(v9^i32(757935405))|v9)&i32(-2139062144) != i32(-2139062144) {
					l17:
						{
							t12 := int32(m.memory[uint32(v2)])
							if t12 == i32(45) {
								goto l15
							}
							v2 = v2 + i32(1)
							if v2 == v5 {
								goto l5
							}
							goto l17
						}
					}
					t8 := v7
					v9 = i32(4) - v7&i32(3)
					v2 = t8 + v9
					if uint32(v6) < uint32(i32(9)) {
						if uint32(v9) >= uint32(v6) {
							goto l5
						}
					l16:
						{
							t11 := int32(m.memory[uint32(v2)])
							if t11 == i32(45) {
								goto l15
							}
							v2 = v2 + i32(1)
							if v2 == v5 {
								goto l5
							}
							goto l16
						}
					}
					if v9 > v6+i32(-8) {
						goto l13
					}
					v10 = v5 + i32(-8)
				l14:
					{
						t9 := int32(load32(m.memory[uint32(v2):]))
						v9 = t9
						if (i32(16843008)-(v9^i32(757935405))|v9)&i32(-2139062144) != i32(-2139062144) {
							goto l13
						}
						t10 := int32(load32(m.memory[uint32(v2+i32(4)):]))
						v9 = t10
						if (i32(16843008)-(v9^i32(757935405))|v9)&i32(-2139062144) != i32(-2139062144) {
							goto l13
						}
						v2 = v2 + i32(8)
						if uint32(v2) <= uint32(v10) {
							goto l14
						}
						goto l13
					}
				}
			l13:
				if uint32(v2) >= uint32(v5) {
					goto l5
				}
			l19:
				{
					t14 := int32(m.memory[uint32(v2)])
					if t14 == i32(45) {
						goto l15
					}
					v2 = v2 + i32(1)
					if v2 == v5 {
						goto l5
					}
					goto l19
				}
			l15:
				v9 = v2 - v7
				v2 = v9 + i32(1)
				v8 = v2 + v8
				v5 = v8 + i32(4)
				if uint32(v5) >= uint32(v4) {
					m.fn33(v5, v4, i32(1272604))
					panic("unreachable")
				}
				t15 := int32(m.memory[uint32(v3+v5)])
				if t15 == i32(45) {
					store32(m.memory[int64(uint32(v0))+4:], uint32(i32(-0x7ffffffa)))
					t17 := int64(load64(m.memory[int64(uint32(v1))+16:]))
					store64(m.memory[int64(uint32(v1))+24:], uint64(int64(uint32(v9))-int64(uint32(v4))+t17+i64(4)))
					goto l24
				}
				v5 = v7 + v6
				v7 = v7 + v2
				v6 = v6 - v2
				if v6 != 0 {
					goto l22
				}
			}
		l5:
			v2 = v4 + i32(-3)
			if uint32(v2) < uint32(i32(4)) {
				m.fn121(i32(4), v2, v4, i32(1272620))
				panic("unreachable")
			}
			v4 = v4 + i32(-7)
		l8:
			store32(m.memory[int64(uint32(v0))+16:], uint32(v4))
			store64(m.memory[int64(uint32(v0))+4:], uint64(i64(-0xfffffffb)))
			t16 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			store32(m.memory[int64(uint32(v0))+20:], uint32(t16))
			store32(m.memory[int64(uint32(v0))+12:], uint32(v3+i32(4)))
			store32(m.memory[uint32(v0):], uint32(i32(0)))
			return
		}
	l3:
		m.memory[int64(uint32(v0))+8] = byte(v2)
		store32(m.memory[int64(uint32(v0))+4:], uint32(i32(-0x7ffffff7)))
		t41 := int64(load64(m.memory[int64(uint32(v1))+16:]))
		store64(m.memory[int64(uint32(v1))+24:], uint64(t41-int64(uint32(v4))))
	}
l24:
	store32(m.memory[uint32(v0):], uint32(i32(1)))
}
func (m *Module) fn225(v0, v1, v2 int32) {
	var v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13 int32
	t0 := m.g0
	v3 = t0 - i32(80)
	m.g0 = v3
	t1 := int32(load16(m.memory[int64(uint32(v1))+28:]))
	store16(m.memory[int64(uint32(v1))+28:], uint16(t1+i32(1)))
	store64(m.memory[int64(uint32(v3))+24:], uint64(i64(0)))
	store64(m.memory[int64(uint32(v3))+16:], uint64(i64(0x400000000)))
	store32(m.memory[int64(uint32(v3))+8:], uint32(i32(1)))
	v4 = i32(0)
	store16(m.memory[int64(uint32(v3))+44:], uint16(i32(0)))
	t2 := int64(load64(m.memory[int64(uint32(v2))+4:]))
	store64(m.memory[int64(uint32(v3))+48:], uint64(t2))
	t3 := int32(load32(m.memory[int64(uint32(v2))+12:]))
	store32(m.memory[int64(uint32(v3))+56:], uint32(t3))
	t4 := int32(load32(m.memory[int64(uint32(v2))+16:]))
	store32(m.memory[int64(uint32(v3))+12:], uint32(t4))
	v5 = v1 + i32(12)
l60:
	m.fn1823(v3+i32(60), v3+i32(8))
	{
		{
			{
				{
					t5 := int32(load32(m.memory[int64(uint32(v3))+60:]))
					v2 = t5
					if v2 > i32(-2) {
						t16 := int32(load32(m.memory[int64(uint32(v3))+64:]))
						v6 = t16
						t17 := int32(load32(m.memory[int64(uint32(v3))+76:]))
						v7 = t17
						if uint32(v7) < uint32(i32(5)) {
							goto l10
						}
						t18 := int32(load32(m.memory[int64(uint32(v3))+72:]))
						v8 = t18
						t19 := int32(load32(m.memory[uint32(v8):]))
						t20 := int32(m.memory[uint32(v8+i32(4))])
						if t19^i32(1852599672)|(t20^i32(115)) != 0 {
							goto l10
						}
						t21 := int32(load32(m.memory[int64(uint32(v3))+68:]))
						v9 = t21
						{
							{
								if v7 == i32(5) {
									goto l11
								}
								t22 := int32(m.memory[int64(uint32(v8))+5])
								if t22 != i32(58) {
									goto l10
								}
								t23 := int32(load32(m.memory[int64(uint32(v1))+24:]))
								t24 := v4
								v10 = t23
								if uint32(t24) >= uint32(v10) {
									goto l12
								}
								v11 = v8 + i32(6)
								t25 := int32(load16(m.memory[int64(uint32(v1))+28:]))
								v12 = t25
								switch v7 + i32(-9) {
								default:
									goto l14
								case 0:
									t26 := int32(m.memory[uint32(v11)])
									if t26 != i32(120) {
										goto l14
									}
									t27 := int32(m.memory[int64(uint32(v8))+7])
									if t27 != i32(109) {
										goto l14
									}
									t28 := int32(m.memory[int64(uint32(v8))+8])
									if t28 != i32(108) {
										goto l14
									}
									{
										if v9 != i32(36) {
											v1 = i32(0)
											{
												if v9 < i32(0) {
													goto l19
												}
												if v9 != 0 {
													goto l20
												}
												v1 = i32(1)
												v9 = i32(0)
												v8 = i32(0)
												goto l21
											l20:
												v8 = v9
												t35 := m.fn11(v9)
												v1 = t35
												if v1 != 0 {
													goto l18
												}
												v1 = i32(1)
											}
										l19:
											m.fn16(v1, v9)
											panic("unreachable")
										}
										t29 := int64(load64(m.memory[uint32(v6):]))
										t30 := int64(load64(m.memory[uint32(v6+i32(8)):]))
										t31 := int64(load64(m.memory[uint32(v6+i32(16)):]))
										t32 := int64(load64(m.memory[uint32(v6+i32(24)):]))
										t33 := int64(load32(m.memory[uint32(v6+i32(32)):]))
										if t29^i64(8588134942460114024)|(t30^i64(0x726f2e33772e7777))|(t31^i64(4121127138782359399)|(t32^i64(8315172552237332537)))|(t33^i64(1701011824)) == 0 {
											goto l17
										}
										v8 = i32(36)
										t34 := m.fn11(i32(36))
										v1 = t34
										if v1 != 0 {
											goto l18
										}
										m.fn16(i32(1), i32(36))
										panic("unreachable")
									}
								l18:
									if v9 == 0 {
										goto l21
									}
									memory_copy(m.memory, uint32(v1), uint32(v6), uint32(v9))
								l21:
									v7 = i32(1)
									goto l22
								case 2:
									t36 := int32(m.memory[uint32(v11)])
									if t36 != i32(120) {
										goto l14
									}
									t37 := int32(m.memory[int64(uint32(v8))+7])
									if t37 != i32(109) {
										goto l14
									}
									t38 := int32(m.memory[int64(uint32(v8))+8])
									if t38 != i32(108) {
										goto l14
									}
									t39 := int32(m.memory[int64(uint32(v8))+9])
									if t39 != i32(110) {
										goto l14
									}
									t40 := int32(m.memory[int64(uint32(v8))+10])
									if t40 != i32(115) {
										goto l14
									}
									v1 = i32(0)
									{
										if v9 < i32(0) {
											goto l23
										}
										v7 = i32(2)
										if v9 == 0 {
											goto l24
										}
										t41 := m.fn11(v9)
										v1 = t41
										if v1 != 0 {
											if v9 == 0 {
												goto l26
											}
											memory_copy(m.memory, uint32(v1), uint32(v6), uint32(v9))
										l26:
											v8 = v9
											goto l22
										}
										v1 = i32(1)
									}
								l23:
									m.fn16(v1, v9)
									panic("unreachable")
								}
							}
						l11:
							t42 := int32(load32(m.memory[int64(uint32(v1))+24:]))
							t43 := v4
							v10 = t42
							if uint32(t43) < uint32(v10) {
								goto l27
							}
						}
					l12:
						store32(m.memory[uint32(v0):], uint32(i32(5)))
						v1 = i32(4)
						v9 = v10
						goto l28
					}
					{
						t6 := int32(load32(m.memory[int64(uint32(v3))+16:]))
						v2 = t6
						if v2 == 0 {
							goto l1
						}
						t7 := int32(load32(m.memory[int64(uint32(v3))+20:]))
						v1 = t7
						t8 := int32(load32(m.memory[uint32(v1+i32(-4)):]))
						v6 = t8
						v7 = v6 & i32(-8)
						t9 := v7
						v6 = v6 & i32(3)
						p10 := i32(8)
						if v6 != 0 {
							p10 = i32(4)
						}
						v2 = v2 << 3
						if uint32(t9) < uint32(p10+v2) {
							m.fn7(i32(1274404), i32(46), i32(1274452))
							panic("unreachable")
						}
						if v6 == 0 {
							goto l3
						}
						if uint32(v7) > uint32(v2+i32(39)) {
							m.fn7(i32(1274468), i32(46), i32(1274516))
							panic("unreachable")
						}
					l3:
						m.fn5(v1)
					}
				l1:
					{
						t11 := int32(load32(m.memory[int64(uint32(v3))+28:]))
						v6 = t11
						if v6 == 0 {
							goto l5
						}
						t12 := int32(load32(m.memory[int64(uint32(v3))+32:]))
						v2 = t12
						if v2 == 0 {
							goto l5
						}
						v1 = v2 << 3
						v2 = v1 + v2 + i32(17)
						if v2 == 0 {
							goto l5
						}
						v1 = v6 - v1
						t13 := int32(load32(m.memory[uint32(v1+i32(-12)):]))
						v6 = t13
						v7 = v6 & i32(-8)
						t14 := v7
						v6 = v6 & i32(3)
						p15 := i32(8)
						if v6 != 0 {
							p15 = i32(4)
						}
						if uint32(t14) < uint32(p15+v2) {
							m.fn7(i32(1274404), i32(46), i32(1274452))
							panic("unreachable")
						}
						if v6 == 0 {
							goto l7
						}
						if uint32(v7) > uint32(v2+i32(39)) {
							m.fn7(i32(1274468), i32(46), i32(1274516))
							panic("unreachable")
						}
					l7:
						m.fn5(v1 + i32(-8))
					}
				l5:
					store32(m.memory[uint32(v0):], uint32(i32(-1)))
					goto l9
				}
			l14:
				v8 = v7 + i32(-6)
				{
					switch v9 + i32(-29) {
					default:
						goto l30
					case 7:
						t44 := int64(load64(m.memory[uint32(v6):]))
						t45 := int64(load64(m.memory[uint32(v6+i32(8)):]))
						t46 := int64(load64(m.memory[uint32(v6+i32(16)):]))
						t47 := int64(load64(m.memory[uint32(v6+i32(24)):]))
						t48 := int64(load32(m.memory[uint32(v6+i32(32)):]))
						if t44^i64(8588134942460114024)|(t45^i64(0x726f2e33772e7777))|(t46^i64(4121127138782359399)|(t47^i64(8315172552237332537)))|(t48^i64(1701011824)) != i64(0) {
							goto l30
						}
						if v8 <= i32(-1) {
							goto l32
						}
						v7 = i32(3)
						if v8 == 0 {
							goto l24
						}
						t49 := m.fn11(v8)
						v1 = t49
						if v1 == 0 {
							m.fn16(i32(1), v8)
							panic("unreachable")
						}
						if v8 == 0 {
							goto l34
						}
						memory_copy(m.memory, uint32(v1), uint32(v11), uint32(v8))
					l34:
						v9 = v8
						goto l22
					case 0:
						t50 := int64(load64(m.memory[uint32(v6):]))
						t51 := int64(load64(m.memory[uint32(v6+i32(8)):]))
						t52 := int64(load64(m.memory[uint32(v6+i32(16)):]))
						t53 := int64(load64(m.memory[uint32(v6+i32(21)):]))
						if t50^i64(8588134942460114024)|(t51^i64(0x726f2e33772e7777))|(t52^i64(8660193591981911911)|(t53^i64(0x2f736e6c6d782f30))) == 0 {
							if v8 <= i32(-1) {
								goto l32
							}
							v7 = i32(4)
							if v8 == 0 {
								goto l24
							}
							t66 := m.fn11(v8)
							v1 = t66
							if v1 == 0 {
								m.fn16(i32(1), v8)
								panic("unreachable")
							}
							if v8 == 0 {
								goto l45
							}
							memory_copy(m.memory, uint32(v1), uint32(v11), uint32(v8))
						l45:
							v9 = v8
							goto l22
						}
					}
				l30:
					{
						{
							t54 := int32(load32(m.memory[uint32(v1):]))
							t55 := v8
							v13 = t54
							t56 := int32(load32(m.memory[int64(uint32(v1))+8:]))
							t57 := v13
							v10 = t56
							if uint32(t55) <= uint32(t57-v10) {
								goto l36
							}
							m.fn245(v1, v10, v8)
							t58 := int32(load32(m.memory[uint32(v1):]))
							v13 = t58
							t59 := int32(load32(m.memory[int64(uint32(v1))+8:]))
							v7 = t59
							goto l37
						}
					l36:
						v7 = v10
						if v8 == 0 {
							goto l38
						}
					l37:
						if v8 == 0 {
							goto l38
						}
						t60 := int32(load32(m.memory[int64(uint32(v1))+4:]))
						memory_copy(m.memory, uint32(t60+v7), uint32(v11), uint32(v8))
					}
				l38:
					t61 := v1
					v7 = v7 + v8
					store32(m.memory[int64(uint32(t61))+8:], uint32(v7))
					{
						{
							if uint32(v9) <= uint32(v13-v7) {
								goto l39
							}
							m.fn245(v1, v7, v9)
							t62 := int32(load32(m.memory[int64(uint32(v1))+8:]))
							v7 = t62
							goto l40
						}
					l39:
						if v9 == 0 {
							goto l41
						}
					l40:
						if v9 == 0 {
							goto l41
						}
						t63 := int32(load32(m.memory[int64(uint32(v1))+4:]))
						memory_copy(m.memory, uint32(t63+v7), uint32(v6), uint32(v9))
					}
				l41:
					store32(m.memory[int64(uint32(v1))+8:], uint32(v7+v9))
					t64 := int32(load32(m.memory[int64(uint32(v1))+20:]))
					v7 = t64
					t65 := int32(load32(m.memory[int64(uint32(v1))+12:]))
					if v7 == t65 {
						goto l42
					}
					goto l43
				}
			l32:
				m.fn15()
				panic("unreachable")
			l24:
				v1 = i32(1)
				v9 = i32(0)
				v8 = i32(0)
			l22:
				store32(m.memory[int64(uint32(v0))+8:], uint32(v1))
				store32(m.memory[int64(uint32(v0))+4:], uint32(v8))
				store32(m.memory[uint32(v0):], uint32(v7))
				v1 = i32(12)
			l28:
				store32(m.memory[uint32(v0+v1):], uint32(v9))
				{
					if v2 < i32(1) {
						goto l46
					}
					t67 := int32(load32(m.memory[uint32(v6+i32(-4)):]))
					v1 = t67
					v7 = v1 & i32(-8)
					t68 := v7
					v1 = v1 & i32(3)
					p69 := i32(8)
					if v1 != 0 {
						p69 = i32(4)
					}
					if uint32(t68) < uint32(p69+v2) {
						m.fn7(i32(1274404), i32(46), i32(1274452))
						panic("unreachable")
					}
					if v1 == 0 {
						goto l48
					}
					if uint32(v7) > uint32(v2+i32(39)) {
						m.fn7(i32(1274468), i32(46), i32(1274516))
						panic("unreachable")
					}
				l48:
					m.fn5(v6)
				}
			l46:
				{
					t70 := int32(load32(m.memory[int64(uint32(v3))+16:]))
					v2 = t70
					if v2 == 0 {
						goto l50
					}
					t71 := int32(load32(m.memory[int64(uint32(v3))+20:]))
					v1 = t71
					t72 := int32(load32(m.memory[uint32(v1+i32(-4)):]))
					v6 = t72
					v7 = v6 & i32(-8)
					t73 := v7
					v6 = v6 & i32(3)
					p74 := i32(8)
					if v6 != 0 {
						p74 = i32(4)
					}
					v2 = v2 << 3
					if uint32(t73) < uint32(p74+v2) {
						m.fn7(i32(1274404), i32(46), i32(1274452))
						panic("unreachable")
					}
					if v6 == 0 {
						goto l52
					}
					if uint32(v7) > uint32(v2+i32(39)) {
						m.fn7(i32(1274468), i32(46), i32(1274516))
						panic("unreachable")
					}
				l52:
					m.fn5(v1)
				}
			l50:
				t75 := int32(load32(m.memory[int64(uint32(v3))+28:]))
				v6 = t75
				if v6 == 0 {
					goto l9
				}
				t76 := int32(load32(m.memory[int64(uint32(v3))+32:]))
				v2 = t76
				if v2 == 0 {
					goto l9
				}
				v1 = v2 << 3
				v2 = v1 + v2 + i32(17)
				if v2 == 0 {
					goto l9
				}
				v1 = v6 - v1
				t77 := int32(load32(m.memory[uint32(v1+i32(-12)):]))
				v6 = t77
				v7 = v6 & i32(-8)
				t78 := v7
				v6 = v6 & i32(3)
				p79 := i32(8)
				if v6 != 0 {
					p79 = i32(4)
				}
				if uint32(t78) < uint32(p79+v2) {
					m.fn7(i32(1274404), i32(46), i32(1274452))
					panic("unreachable")
				}
				if v6 == 0 {
					goto l55
				}
				if uint32(v7) > uint32(v2+i32(39)) {
					m.fn7(i32(1274468), i32(46), i32(1274516))
					panic("unreachable")
				}
			l55:
				m.fn5(v1 + i32(-8))
			}
		l9:
			m.g0 = v3 + i32(80)
			return
		l27:
			t80 := int32(load16(m.memory[int64(uint32(v1))+28:]))
			v12 = t80
			{
				{
					t81 := int32(load32(m.memory[uint32(v1):]))
					t82 := int32(load32(m.memory[int64(uint32(v1))+8:]))
					t83 := v9
					v10 = t82
					if uint32(t83) <= uint32(t81-v10) {
						goto l57
					}
					m.fn245(v1, v10, v9)
					t84 := int32(load32(m.memory[int64(uint32(v1))+8:]))
					v7 = t84
					goto l58
				}
			l57:
				v7 = v10
				if v9 == 0 {
					goto l59
				}
			l58:
				if v9 == 0 {
					goto l59
				}
				t85 := int32(load32(m.memory[int64(uint32(v1))+4:]))
				memory_copy(m.memory, uint32(t85+v7), uint32(v6), uint32(v9))
			}
		l59:
			store32(m.memory[int64(uint32(v1))+8:], uint32(v7+v9))
			v8 = i32(0)
			t86 := int32(load32(m.memory[int64(uint32(v1))+20:]))
			v7 = t86
			t87 := int32(load32(m.memory[int64(uint32(v1))+12:]))
			if v7 != t87 {
				goto l43
			}
		}
	l42:
		m.fn1830(v5)
	l43:
		store32(m.memory[int64(uint32(v1))+20:], uint32(v7+i32(1)))
		t88 := int32(load32(m.memory[int64(uint32(v1))+16:]))
		v7 = t88 + v7<<4
		store16(m.memory[int64(uint32(v7))+12:], uint16(v12))
		store32(m.memory[int64(uint32(v7))+8:], uint32(v9))
		store32(m.memory[int64(uint32(v7))+4:], uint32(v8))
		store32(m.memory[uint32(v7):], uint32(v10))
	}
l17:
	v4 = v4 + i32(1)
l10:
	if v2 < i32(1) {
		goto l60
	}
	{
		t89 := int32(load32(m.memory[uint32(v6+i32(-4)):]))
		v7 = t89
		v8 = v7 & i32(-8)
		t90 := v8
		v7 = v7 & i32(3)
		p91 := i32(8)
		if v7 != 0 {
			p91 = i32(4)
		}
		if uint32(t90) < uint32(p91+v2) {
			m.fn7(i32(1274404), i32(46), i32(1274452))
			panic("unreachable")
		}
		if v7 == 0 {
			goto l62
		}
		if uint32(v8) > uint32(v2+i32(39)) {
			m.fn7(i32(1274468), i32(46), i32(1274516))
			panic("unreachable")
		}
	l62:
		m.fn5(v6)
		goto l60
	}
}
func (m *Module) fn226(v0, v1, v2, v3 int32) {
	var v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14, v15, v16, v17, v18, v19, v20, v21, v22, v23, v24, v25, v26, v27, v28, v29, v30, v31, v32, v33, v34 int32
	t0 := m.g0
	v4 = t0 - i32(192)
	m.g0 = v4
	{
		{
			t1 := int32(load32(m.memory[int64(uint32(v1))+16:]))
			v5 = t1
			t2 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			t3 := v5
			v6 = t2
			if uint32(t3) > uint32(v6) {
				m.fn121(i32(0), v5, v6, i32(1272488))
				panic("unreachable")
			}
			t4 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t5 := v4 + i32(40)
			v7 = t4
			m.fn238(t5, v7, v5)
			t6 := int32(load32(m.memory[int64(uint32(v4))+44:]))
			v8 = t6
			t7 := int32(load32(m.memory[int64(uint32(v4))+40:]))
			v9 = t7
			t8 := v4 + i32(132)
			v10 = v2 + i32(72)
			t9 := int32(load32(m.memory[int64(uint32(v4))+48:]))
			t10 := int32(load32(m.memory[int64(uint32(v4))+52:]))
			m.fn239(t8, v10, t9, t10, i32(1))
			t11 := int32(load32(m.memory[int64(uint32(v4))+136:]))
			v11 = t11
			t12 := int32(load32(m.memory[int64(uint32(v4))+132:]))
			v12 = t12
			if v12 == i32(-0x7fffffff) {
				goto l1
			}
			v13 = i32(0)
			goto l2
		}
	l1:
		t13 := int32(load32(m.memory[int64(uint32(v4))+140:]))
		m.fn240(v4+i32(16), v3, v11, t13)
		t14 := int32(load32(m.memory[int64(uint32(v4))+20:]))
		v14 = t14
		t15 := int32(load32(m.memory[int64(uint32(v4))+16:]))
		v13 = t15
	}
l2:
	m.fn29(v4+i32(40), v9, v8)
	t16 := int32(load32(m.memory[int64(uint32(v4))+48:]))
	v15 = t16
	t17 := int32(load32(m.memory[int64(uint32(v4))+44:]))
	v8 = t17
	{
		{
			{
				t18 := int32(load32(m.memory[int64(uint32(v4))+40:]))
				v16 = t18
				if v16 == i32(-1) {
					goto l3
				}
				v17 = v8
				goto l4
			}
		l3:
			if v15 <= i32(-1) {
				goto l5
			}
			if v15 != 0 {
				goto l6
			}
			v17 = i32(1)
			v15 = i32(0)
			v16 = i32(0)
			goto l4
		l6:
			t19 := m.fn11(v15)
			v17 = t19
			if v17 == 0 {
				m.fn16(i32(1), v15)
				panic("unreachable")
			}
			if v15 == 0 {
				goto l8
			}
			memory_copy(m.memory, uint32(v17), uint32(v8), uint32(v15))
		l8:
			v16 = v15
		}
	l4:
		store32(m.memory[int64(uint32(v4))+36:], uint32(i32(0)))
		store64(m.memory[int64(uint32(v4))+28:], uint64(i64(0x400000000)))
		t20 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		store32(m.memory[int64(uint32(v4))+88:], uint32(t20))
		store32(m.memory[int64(uint32(v4))+84:], uint32(v6))
		store32(m.memory[int64(uint32(v4))+80:], uint32(v7))
		store16(m.memory[int64(uint32(v4))+76:], uint16(i32(256)))
		store64(m.memory[int64(uint32(v4))+56:], uint64(i64(0)))
		store64(m.memory[int64(uint32(v4))+48:], uint64(i64(0x400000000)))
		store32(m.memory[int64(uint32(v4))+44:], uint32(v5))
		store32(m.memory[int64(uint32(v4))+40:], uint32(i32(1)))
		t21 := int32(load32(m.memory[int64(uint32(v2))+4:]))
		v18 = t21
		v19 = i32(0)
		v20 = i32(4)
		v21 = i32(0)
	l10:
		{
			t22 := int32(load32(m.memory[int64(uint32(v4))+80:]))
			t23 := int32(load32(m.memory[int64(uint32(v4))+84:]))
			m.fn241(v4+i32(96), v4+i32(40), t22, t23)
			{
				{
					{
						{
							{
								{
									t24 := int32(load32(m.memory[int64(uint32(v4))+96:]))
									v1 = t24
									switch v1 + i32(2) {
									case 1:
										goto l10
									default:
										t25 := int32(load32(m.memory[int64(uint32(v4))+104:]))
										v9 = t25
										t26 := int32(load32(m.memory[int64(uint32(v4))+100:]))
										t27 := v9
										v6 = t26
										t28 := int32(load32(m.memory[int64(uint32(v4))+84:]))
										var p29 int32
										if uint32(t27) < uint32(v6) {
											p29 = 1
										}
										t30 := v9
										v22 = t28
										var p31 int32
										if uint32(t30) > uint32(v22) {
											p31 = 1
										}
										v23 = p29 | p31
										t32 := int32(load32(m.memory[int64(uint32(v4))+80:]))
										v7 = t32
										t33 := int32(load32(m.memory[int64(uint32(v4))+112:]))
										v5 = t33
										t34 := int32(load32(m.memory[int64(uint32(v4))+108:]))
										v8 = t34
										switch v1 {
										default:
											if v23 != 0 {
												m.fn121(v6, v9, v22, i32(1272332))
												panic("unreachable")
											}
											if uint32(v5) < uint32(v8) {
												goto l17
											}
											if uint32(v5) <= uint32(v22) {
												goto l18
											}
										l17:
											m.fn121(v8, v5, v22, i32(1272332))
											panic("unreachable")
										case 1:
											if v23 != 0 {
												m.fn121(v6, v9, v22, i32(1272332))
												panic("unreachable")
											}
											if uint32(v5) < uint32(v8) {
												goto l20
											}
											if uint32(v5) <= uint32(v22) {
												goto l18
											}
										l20:
											m.fn121(v8, v5, v22, i32(1272332))
											panic("unreachable")
										case 2:
											if v23 != 0 {
												m.fn121(v6, v9, v22, i32(1272332))
												panic("unreachable")
											}
											if uint32(v5) < uint32(v8) {
												goto l22
											}
											if uint32(v5) <= uint32(v22) {
												goto l18
											}
										l22:
											m.fn121(v8, v5, v22, i32(1272332))
											panic("unreachable")
										case 3:
											if v23 != 0 {
												m.fn121(v6, v9, v22, i32(1272332))
												panic("unreachable")
											}
											v8 = i32(0)
											v23 = i32(1)
											goto l24
										}
									case 0:
										{
											t35 := int32(load32(m.memory[int64(uint32(v4))+48:]))
											v1 = t35
											if v1 == 0 {
												goto l25
											}
											t36 := int32(load32(m.memory[int64(uint32(v4))+52:]))
											v8 = t36
											t37 := int32(load32(m.memory[uint32(v8+i32(-4)):]))
											v5 = t37
											v9 = v5 & i32(-8)
											t38 := v9
											v5 = v5 & i32(3)
											p39 := i32(8)
											if v5 != 0 {
												p39 = i32(4)
											}
											v1 = v1 << 3
											if uint32(t38) < uint32(p39+v1) {
												m.fn7(i32(1274404), i32(46), i32(1274452))
												panic("unreachable")
											}
											if v5 == 0 {
												goto l27
											}
											if uint32(v9) > uint32(v1+i32(39)) {
												m.fn7(i32(1274468), i32(46), i32(1274516))
												panic("unreachable")
											}
										l27:
											m.fn5(v8)
										}
									l25:
										{
											t40 := int32(load32(m.memory[int64(uint32(v4))+60:]))
											v5 = t40
											if v5 == 0 {
												goto l29
											}
											t41 := int32(load32(m.memory[int64(uint32(v4))+64:]))
											v1 = t41
											if v1 == 0 {
												goto l29
											}
											v8 = v1 << 3
											v1 = v8 + v1 + i32(17)
											if v1 == 0 {
												goto l29
											}
											v8 = v5 - v8
											t42 := int32(load32(m.memory[uint32(v8+i32(-12)):]))
											v5 = t42
											v9 = v5 & i32(-8)
											t43 := v9
											v5 = v5 & i32(3)
											p44 := i32(8)
											if v5 != 0 {
												p44 = i32(4)
											}
											if uint32(t43) < uint32(p44+v1) {
												m.fn7(i32(1274404), i32(46), i32(1274452))
												panic("unreachable")
											}
											if v5 == 0 {
												goto l31
											}
											if uint32(v9) > uint32(v1+i32(39)) {
												m.fn7(i32(1274468), i32(46), i32(1274516))
												panic("unreachable")
											}
										l31:
											m.fn5(v8 + i32(-8))
										}
									l29:
										if v15 != i32(6) {
											goto l33
										}
										t45 := int32(load32(m.memory[uint32(v17):]))
										t46 := int32(load16(m.memory[uint32(v17+i32(4)):]))
										if t45^i32(1768908867)|(t46^i32(25955)) != 0 {
											goto l33
										}
										if v13 == 0 {
											goto l33
										}
										if v14 != i32(59) {
											goto l33
										}
										t47 := int64(load64(m.memory[int64(uint32(v13))+8:]))
										t48 := int64(load64(m.memory[uint32(v13+i32(16)):]))
										t49 := int64(load64(m.memory[uint32(v13+i32(24)):]))
										t50 := int64(load64(m.memory[uint32(v13+i32(32)):]))
										t51 := int64(load64(m.memory[uint32(v13+i32(40)):]))
										t52 := int64(load64(m.memory[uint32(v13+i32(48)):]))
										t53 := int64(load64(m.memory[uint32(v13+i32(56)):]))
										t54 := int64(load64(m.memory[uint32(v13+i32(59)):]))
										if t47^i64(8299904566308402280)|(t48^i64(8011467649423075427))|(t49^i64(8027222603262223728)|(t50^i64(8245860516147326322)))|(t51^i64(0x70756b72616d2f67)|(t52^i64(7598805606781117229))|(t53^i64(3616242566693677410)|(t54^i64(3904673869033206889)))) != i64(0) {
											goto l33
										}
										if v21 == 0 {
											goto l33
										}
									l36:
										{
											t55 := int32(load32(m.memory[uint32(v20+i32(8)):]))
											if t55 != i32(8) {
												goto l34
											}
											t56 := int32(load32(m.memory[uint32(v20+i32(4)):]))
											t57 := int64(load64(m.memory[uint32(t56):]))
											if t57 == i64(0x7365726975716552) {
												v22 = i32(0)
												store32(m.memory[int64(uint32(v4))+100:], uint32(i32(0)))
												store32(m.memory[int64(uint32(v4))+96:], uint32(v2))
												t58 := int32(load32(m.memory[int64(uint32(v20))+20:]))
												v1 = t58
												t59 := int32(load32(m.memory[int64(uint32(v20))+16:]))
												v5 = t59
												store16(m.memory[int64(uint32(v4))+128:], uint16(i32(1)))
												store32(m.memory[int64(uint32(v4))+124:], uint32(i32(0)))
												store32(m.memory[int64(uint32(v4))+116:], uint32(v5))
												store32(m.memory[int64(uint32(v4))+112:], uint32(v1))
												store32(m.memory[int64(uint32(v4))+108:], uint32(v5))
												store32(m.memory[int64(uint32(v4))+104:], uint32(v1))
												store32(m.memory[int64(uint32(v4))+120:], uint32(v5+v1))
												m.fn242(v4+i32(168), v4+i32(96))
												{
													{
														t60 := int32(load32(m.memory[int64(uint32(v4))+168:]))
														if t60 != i32(-1) {
															goto l37
														}
														v23 = i32(4)
														v1 = i32(0)
														goto l38
													}
												l37:
													t61 := m.fn11(i32(48))
													v9 = t61
													if v9 == 0 {
														m.fn16(i32(4), i32(48))
														panic("unreachable")
													}
													t62 := int32(load32(m.memory[int64(uint32(v4))+176:]))
													store32(m.memory[int64(uint32(v9))+8:], uint32(t62))
													t63 := int64(load64(m.memory[int64(uint32(v4))+168:]))
													store64(m.memory[uint32(v9):], uint64(t63))
													store32(m.memory[int64(uint32(v4))+164:], uint32(i32(1)))
													store32(m.memory[int64(uint32(v4))+160:], uint32(v9))
													store32(m.memory[int64(uint32(v4))+156:], uint32(i32(4)))
													t64 := int32(load32(m.memory[int64(uint32(v4))+128:]))
													store32(m.memory[int64(uint32(v4))+72:], uint32(t64))
													t65 := int64(load64(m.memory[int64(uint32(v4))+120:]))
													store64(m.memory[int64(uint32(v4))+64:], uint64(t65))
													t66 := int64(load64(m.memory[int64(uint32(v4))+112:]))
													store64(m.memory[int64(uint32(v4))+56:], uint64(t66))
													t67 := int64(load64(m.memory[int64(uint32(v4))+104:]))
													store64(m.memory[int64(uint32(v4))+48:], uint64(t67))
													t68 := int64(load64(m.memory[int64(uint32(v4))+96:]))
													store64(m.memory[int64(uint32(v4))+40:], uint64(t68))
													v5 = i32(12)
													v1 = i32(1)
												l42:
													{
														m.fn242(v4+i32(180), v4+i32(40))
														t69 := int32(load32(m.memory[int64(uint32(v4))+180:]))
														if t69 == i32(-1) {
															goto l40
														}
														{
															t70 := int32(load32(m.memory[int64(uint32(v4))+156:]))
															if v1 != t70 {
																goto l41
															}
															m.fn197(v4+i32(156), v1, i32(1), i32(4), i32(12))
															t71 := int32(load32(m.memory[int64(uint32(v4))+160:]))
															v9 = t71
														}
													l41:
														v8 = v9 + v5
														t72 := int32(load32(m.memory[int64(uint32(v4))+188:]))
														store32(m.memory[int64(uint32(v8))+8:], uint32(t72))
														t73 := int64(load64(m.memory[int64(uint32(v4))+180:]))
														store64(m.memory[uint32(v8):], uint64(t73))
														t74 := v4
														v1 = v1 + i32(1)
														store32(m.memory[int64(uint32(t74))+164:], uint32(v1))
														v5 = v5 + i32(12)
														goto l42
													}
												l40:
													t75 := int32(load32(m.memory[int64(uint32(v4))+160:]))
													v23 = t75
													t76 := int32(load32(m.memory[int64(uint32(v4))+156:]))
													v22 = t76
												}
											l38:
												m.fn203(v4+i32(40), v23, v1, i32(1089929), i32(1))
												{
													v5 = v20 + i32(12)
													t77 := int32(load32(m.memory[uint32(v5):]))
													v8 = t77
													if v8 == 0 {
														goto l43
													}
													t78 := int32(load32(m.memory[int64(uint32(v20))+16:]))
													m.fn21(t78, v8, i32(1))
												}
											l43:
												t79 := int32(load32(m.memory[int64(uint32(v4))+48:]))
												store32(m.memory[int64(uint32(v5))+8:], uint32(t79))
												t80 := int64(load64(m.memory[int64(uint32(v4))+40:]))
												store64(m.memory[uint32(v5):], uint64(t80))
												if v1 == 0 {
													goto l44
												}
												v5 = v23
											l49:
												{
													t81 := int32(load32(m.memory[uint32(v5):]))
													v8 = t81
													if v8 == 0 {
														goto l45
													}
													t82 := int32(load32(m.memory[uint32(v5+i32(4)):]))
													v6 = t82
													t83 := int32(load32(m.memory[uint32(v6+i32(-4)):]))
													v9 = t83
													v7 = v9 & i32(-8)
													t84 := v7
													v9 = v9 & i32(3)
													p85 := i32(8)
													if v9 != 0 {
														p85 = i32(4)
													}
													if uint32(t84) < uint32(p85+v8) {
														m.fn7(i32(1274404), i32(46), i32(1274452))
														panic("unreachable")
													}
													if v9 == 0 {
														goto l47
													}
													if uint32(v7) > uint32(v8+i32(39)) {
														m.fn7(i32(1274468), i32(46), i32(1274516))
														panic("unreachable")
													}
												l47:
													m.fn5(v6)
												}
											l45:
												v5 = v5 + i32(12)
												v1 = v1 + i32(-1)
												if v1 != 0 {
													goto l49
												}
											l44:
												if v22 == 0 {
													goto l33
												}
												m.fn21(v23, v22*i32(12), i32(4))
												goto l33
											}
										}
									l34:
										v20 = v20 + i32(32)
										v19 = v19 + i32(-32)
										if v19 == 0 {
											goto l33
										}
										goto l36
									}
								}
							l18:
								v23 = v7 + v8
								v8 = v5 - v8
							l24:
								v1 = v7 + v6
								{
									v5 = v9 - v6
									if uint32(v5) < uint32(i32(5)) {
										goto l50
									}
									t86 := int32(load32(m.memory[uint32(v1):]))
									t87 := int32(m.memory[uint32(v1+i32(4))])
									if t86^i32(1852599672)|(t87^i32(115)) != 0 {
										goto l50
									}
									if v5 == i32(5) {
										goto l10
									}
									t88 := int32(m.memory[int64(uint32(v1))+5])
									if t88 == i32(58) {
										goto l10
									}
								}
							l50:
								m.fn238(v4+i32(96), v1, v5)
								t89 := int32(load32(m.memory[int64(uint32(v4))+100:]))
								v24 = t89
								t90 := int32(load32(m.memory[int64(uint32(v4))+96:]))
								v25 = t90
								v26 = i32(0)
								t91 := int32(load32(m.memory[int64(uint32(v4))+104:]))
								t92 := int32(load32(m.memory[int64(uint32(v4))+108:]))
								m.fn239(v4+i32(144), v10, t91, t92, i32(0))
								t93 := int32(load32(m.memory[int64(uint32(v4))+148:]))
								v27 = t93
								{
									t94 := int32(load32(m.memory[int64(uint32(v4))+144:]))
									v28 = t94
									if v28 != i32(-0x7fffffff) {
										goto l51
									}
									t95 := int32(load32(m.memory[int64(uint32(v4))+152:]))
									m.fn240(v4+i32(8), v3, v27, t95)
									t96 := int32(load32(m.memory[int64(uint32(v4))+12:]))
									v29 = t96
									t97 := int32(load32(m.memory[int64(uint32(v4))+8:]))
									v26 = t97
								}
							l51:
								m.fn243(v4+i32(96), v18, v23, v8)
								{
									t98 := int32(load32(m.memory[int64(uint32(v4))+96:]))
									v30 = t98
									if v30 != i32(-2) {
										t100 := int32(load32(m.memory[int64(uint32(v4))+100:]))
										t101 := v4
										v9 = t100
										t102 := int32(load32(m.memory[int64(uint32(v4))+104:]))
										t103 := v9
										v1 = t102
										store32(m.memory[int64(uint32(t101))+160:], uint32(t103+v1))
										if v1 != 0 {
											v5 = i32(0)
										l58:
											{
												{
													v7 = v9 + v5
													t104 := int32(m.memory[uint32(v7)])
													v6 = t104 + i32(-9)
													if uint32(v6) > uint32(i32(29)) {
														goto l56
													}
													if i32_shl(i32(1), v6)&i32(0x20000013) != 0 {
														store32(m.memory[int64(uint32(v4))+156:], uint32(v7+i32(1)))
														if v1 <= i32(-1) {
															goto l5
														}
														{
															t106 := m.fn11(v1)
															v6 = t106
															if v6 == 0 {
																m.fn16(i32(1), v1)
																panic("unreachable")
															}
															store32(m.memory[int64(uint32(v4))+188:], uint32(i32(0)))
															store32(m.memory[int64(uint32(v4))+184:], uint32(v6))
															store32(m.memory[int64(uint32(v4))+180:], uint32(v1))
															m.fn244(v4+i32(96), v4+i32(180), v4+i32(156), v9, v1, i32(0), v5)
															t107 := int32(load32(m.memory[int64(uint32(v4))+100:]))
															v22 = t107
															t108 := int32(load32(m.memory[int64(uint32(v4))+96:]))
															v5 = t108
															if v5 != i32(-1) {
																goto l60
															}
															t109 := int32(load32(m.memory[int64(uint32(v4))+156:]))
															v6 = t109
															t110 := int32(load32(m.memory[int64(uint32(v4))+160:]))
															t111 := v6
															v31 = t110
															if t111 == v31 {
																goto l61
															}
														l65:
															v5 = i32(0)
														l64:
															{
																{
																	v32 = v6 + v5
																	t112 := int32(m.memory[uint32(v32)])
																	v7 = t112 + i32(-9)
																	if uint32(v7) > uint32(i32(29)) {
																		goto l62
																	}
																	if i32_shl(i32(1), v7)&i32(0x20000013) != 0 {
																		store32(m.memory[int64(uint32(v4))+156:], uint32(v32+i32(1)))
																		m.fn244(v4+i32(96), v4+i32(180), v4+i32(156), v9, v1, v22, v5+v22)
																		t114 := int32(load32(m.memory[int64(uint32(v4))+100:]))
																		v22 = t114
																		t115 := int32(load32(m.memory[int64(uint32(v4))+96:]))
																		v5 = t115
																		if v5 != i32(-1) {
																			goto l60
																		}
																		t116 := int32(load32(m.memory[int64(uint32(v4))+156:]))
																		v6 = t116
																		t117 := int32(load32(m.memory[int64(uint32(v4))+160:]))
																		t118 := v6
																		v31 = t117
																		if t118 != v31 {
																			goto l65
																		}
																		goto l61
																	}
																}
															l62:
																t113 := v6
																v5 = v5 + i32(1)
																if t113+v5 != v31 {
																	goto l64
																}
																goto l61
															}
														}
													}
												}
											l56:
												t105 := v1
												v5 = v5 + i32(1)
												if t105 == v5 {
													goto l55
												}
												goto l58
											}
										}
										v1 = i32(0)
										goto l55
									}
									v6 = i32(-0x7ffffff4)
									v22 = i32(2)
									t99 := int32(load32(m.memory[int64(uint32(v4))+168:]))
									v1 = t99
									v5 = v18
									goto l53
								}
							}
						l60:
							t119 := int32(load32(m.memory[int64(uint32(v4))+112:]))
							v6 = t119
							t120 := int32(load32(m.memory[int64(uint32(v4))+108:]))
							v7 = t120
							t121 := int32(load32(m.memory[int64(uint32(v4))+104:]))
							v1 = t121
							{
								t122 := int32(load32(m.memory[int64(uint32(v4))+180:]))
								v32 = t122
								if v32 == 0 {
									goto l66
								}
								t123 := int32(load32(m.memory[int64(uint32(v4))+184:]))
								v33 = t123
								t124 := int32(load32(m.memory[uint32(v33+i32(-4)):]))
								v31 = t124
								v34 = v31 & i32(-8)
								t125 := v34
								v31 = v31 & i32(3)
								p126 := i32(8)
								if v31 != 0 {
									p126 = i32(4)
								}
								if uint32(t125) < uint32(p126+v32) {
									m.fn7(i32(1274404), i32(46), i32(1274452))
									panic("unreachable")
								}
								if v31 == 0 {
									goto l68
								}
								if uint32(v34) > uint32(v32+i32(39)) {
									m.fn7(i32(1274468), i32(46), i32(1274516))
									panic("unreachable")
								}
							l68:
								m.fn5(v33)
							}
						l66:
							store32(m.memory[int64(uint32(v4))+176:], uint32(v6))
							store64(m.memory[int64(uint32(v4))+168:], uint64(int64(uint32(v7))<<32|int64(uint32(v1))))
							v6 = i32(-0x7ffffff3)
							if v30 < i32(1) {
								goto l53
							}
							t127 := int32(load32(m.memory[uint32(v9+i32(-4)):]))
							v7 = t127
							v32 = v7 & i32(-8)
							t128 := v32
							v7 = v7 & i32(3)
							p129 := i32(8)
							if v7 != 0 {
								p129 = i32(4)
							}
							if uint32(t128) < uint32(p129+v30) {
								m.fn7(i32(1274404), i32(46), i32(1274452))
								panic("unreachable")
							}
							if v7 == 0 {
								goto l71
							}
							if uint32(v32) <= uint32(v30+i32(39)) {
								goto l71
							}
							m.fn7(i32(1274468), i32(46), i32(1274516))
							panic("unreachable")
						l71:
							m.fn5(v9)
						}
					l53:
						t130 := int64(load64(m.memory[int64(uint32(v4))+172:]))
						store64(m.memory[int64(uint32(v4))+112:], uint64(t130))
						store32(m.memory[int64(uint32(v4))+108:], uint32(v1))
						store32(m.memory[int64(uint32(v4))+104:], uint32(v22))
						store32(m.memory[int64(uint32(v4))+100:], uint32(v5))
						store32(m.memory[int64(uint32(v4))+96:], uint32(v6))
						m.fn29(v4+i32(180), v23, v8)
						t131 := int32(load32(m.memory[int64(uint32(v4))+188:]))
						v1 = t131
						t132 := int32(load32(m.memory[int64(uint32(v4))+184:]))
						v5 = t132
						{
							{
								t133 := int32(load32(m.memory[int64(uint32(v4))+180:]))
								v6 = t133
								if v6 == i32(-1) {
									goto l72
								}
								v7 = v5
								goto l73
							}
						l72:
							if v1 <= i32(-1) {
								goto l5
							}
							if v1 != 0 {
								goto l74
							}
							v7 = i32(1)
							v6 = i32(0)
							v1 = i32(0)
							goto l73
						l74:
							t134 := m.fn11(v1)
							v7 = t134
							if v7 == 0 {
								m.fn16(i32(1), v1)
								panic("unreachable")
							}
							if v1 == 0 {
								goto l76
							}
							memory_copy(m.memory, uint32(v7), uint32(v5), uint32(v1))
						l76:
							v6 = v1
						}
					l73:
						m.fn232(v4 + i32(96))
						goto l77
					}
				l61:
					{
						if v22 == 0 {
							goto l78
						}
						if uint32(v1) > uint32(v22) {
							goto l79
						}
						if v1 != v22 {
							goto l80
						}
						goto l78
					l79:
						t135 := int32(int8(m.memory[uint32(v9+v22)]))
						if t135 <= i32(-65) {
							goto l80
						}
					}
				l78:
					{
						{
							v5 = v1 - v22
							t136 := int32(load32(m.memory[int64(uint32(v4))+180:]))
							t137 := int32(load32(m.memory[int64(uint32(v4))+188:]))
							t138 := v5
							v8 = t137
							if uint32(t138) <= uint32(t136-v8) {
								goto l81
							}
							m.fn245(v4+i32(180), v8, v5)
							t139 := int32(load32(m.memory[int64(uint32(v4))+188:]))
							v8 = t139
							goto l82
						}
					l81:
						if v1 == v22 {
							goto l83
						}
					l82:
						if v5 == 0 {
							goto l83
						}
						t140 := int32(load32(m.memory[int64(uint32(v4))+184:]))
						memory_copy(m.memory, uint32(t140+v8), uint32(v9+v22), uint32(v5))
					}
				l83:
					store32(m.memory[int64(uint32(v4))+188:], uint32(v8+v5))
				l80:
					t141 := int32(load32(m.memory[int64(uint32(v4))+180:]))
					v6 = t141
					if v6 == i32(-1) {
						goto l55
					}
					t142 := int32(load32(m.memory[int64(uint32(v4))+188:]))
					v1 = t142
					t143 := int32(load32(m.memory[int64(uint32(v4))+184:]))
					v7 = t143
					if v30 < i32(1) {
						goto l77
					}
					{
						t144 := int32(load32(m.memory[uint32(v9+i32(-4)):]))
						v5 = t144
						v8 = v5 & i32(-8)
						t145 := v8
						v5 = v5 & i32(3)
						p146 := i32(8)
						if v5 != 0 {
							p146 = i32(4)
						}
						if uint32(t145) < uint32(p146+v30) {
							m.fn7(i32(1274404), i32(46), i32(1274452))
							panic("unreachable")
						}
						if v5 == 0 {
							goto l85
						}
						if uint32(v8) > uint32(v30+i32(39)) {
							m.fn7(i32(1274468), i32(46), i32(1274516))
							panic("unreachable")
						}
					l85:
						m.fn5(v9)
						goto l77
					}
				}
			l55:
				if v30 == i32(-1) {
					goto l87
				}
				v6 = v30
				v7 = v9
				goto l77
			l87:
				if v1 <= i32(-1) {
					goto l5
				}
				if v1 != 0 {
					goto l88
				}
				v7 = i32(1)
				v6 = i32(0)
				goto l77
			l88:
				t147 := m.fn11(v1)
				v7 = t147
				if v7 == 0 {
					m.fn16(i32(1), v1)
					panic("unreachable")
				}
				if v1 == 0 {
					goto l90
				}
				memory_copy(m.memory, uint32(v7), uint32(v9), uint32(v1))
			l90:
				v6 = v1
			}
		l77:
			m.fn29(v4+i32(96), v25, v24)
			t148 := int32(load32(m.memory[int64(uint32(v4))+104:]))
			v8 = t148
			t149 := int32(load32(m.memory[int64(uint32(v4))+100:]))
			v5 = t149
			{
				{
					t150 := int32(load32(m.memory[int64(uint32(v4))+96:]))
					v9 = t150
					if v9 == i32(-1) {
						goto l91
					}
					v23 = v5
					goto l92
				}
			l91:
				if v8 <= i32(-1) {
					goto l5
				}
				if v8 != 0 {
					goto l93
				}
				v23 = i32(1)
				v9 = i32(0)
				v8 = i32(0)
				goto l92
			l93:
				t151 := m.fn11(v8)
				v23 = t151
				if v23 == 0 {
					m.fn16(i32(1), v8)
					panic("unreachable")
				}
				if v8 == 0 {
					goto l95
				}
				memory_copy(m.memory, uint32(v23), uint32(v5), uint32(v8))
			l95:
				v9 = v8
			}
		l92:
			{
				t152 := int32(load32(m.memory[int64(uint32(v4))+28:]))
				if v21 != t152 {
					goto l96
				}
				m.fn246(v4 + i32(28))
				t153 := int32(load32(m.memory[int64(uint32(v4))+32:]))
				v20 = t153
			}
		l96:
			v5 = v20 + v21<<5
			store32(m.memory[int64(uint32(v5))+28:], uint32(v29))
			store32(m.memory[int64(uint32(v5))+24:], uint32(v26))
			store32(m.memory[int64(uint32(v5))+20:], uint32(v1))
			store32(m.memory[int64(uint32(v5))+16:], uint32(v7))
			store32(m.memory[int64(uint32(v5))+12:], uint32(v6))
			store32(m.memory[int64(uint32(v5))+8:], uint32(v8))
			store32(m.memory[int64(uint32(v5))+4:], uint32(v23))
			store32(m.memory[uint32(v5):], uint32(v9))
			t154 := v4
			v21 = v21 + i32(1)
			store32(m.memory[int64(uint32(t154))+36:], uint32(v21))
			{
				if v28 < i32(-0x7ffffffe) {
					goto l97
				}
				if v28 == 0 {
					goto l97
				}
				t155 := int32(load32(m.memory[uint32(v27+i32(-4)):]))
				v1 = t155
				v5 = v1 & i32(-8)
				t156 := v5
				v1 = v1 & i32(3)
				p157 := i32(8)
				if v1 != 0 {
					p157 = i32(4)
				}
				if uint32(t156) < uint32(p157+v28) {
					m.fn7(i32(1274404), i32(46), i32(1274452))
					panic("unreachable")
				}
				if v1 == 0 {
					goto l99
				}
				if uint32(v5) > uint32(v28+i32(39)) {
					m.fn7(i32(1274468), i32(46), i32(1274516))
					panic("unreachable")
				}
			l99:
				m.fn5(v27)
			}
		l97:
			v19 = v19 + i32(32)
			goto l10
		}
	}
l5:
	m.fn15()
	panic("unreachable")
l33:
	t158 := int32(load32(m.memory[int64(uint32(v4))+36:]))
	store32(m.memory[int64(uint32(v0))+20:], uint32(t158))
	t159 := int64(load64(m.memory[int64(uint32(v4))+28:]))
	store64(m.memory[int64(uint32(v0))+12:], uint64(t159))
	store32(m.memory[int64(uint32(v0))+40:], uint32(v14))
	store32(m.memory[int64(uint32(v0))+36:], uint32(v13))
	store32(m.memory[int64(uint32(v0))+32:], uint32(i32(0)))
	store64(m.memory[int64(uint32(v0))+24:], uint64(i64(0x400000000)))
	store32(m.memory[int64(uint32(v0))+8:], uint32(v15))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v17))
	store32(m.memory[uint32(v0):], uint32(v16))
	{
		if v12 < i32(-0x7ffffffe) {
			goto l101
		}
		if v12 == 0 {
			goto l101
		}
		t160 := int32(load32(m.memory[uint32(v11+i32(-4)):]))
		v1 = t160
		v5 = v1 & i32(-8)
		t161 := v5
		v1 = v1 & i32(3)
		p162 := i32(8)
		if v1 != 0 {
			p162 = i32(4)
		}
		if uint32(t161) < uint32(p162+v12) {
			m.fn7(i32(1274404), i32(46), i32(1274452))
			panic("unreachable")
		}
		if v1 == 0 {
			goto l103
		}
		if uint32(v5) > uint32(v12+i32(39)) {
			m.fn7(i32(1274468), i32(46), i32(1274516))
			panic("unreachable")
		}
	l103:
		m.fn5(v11)
	}
l101:
	m.g0 = v4 + i32(192)
}
func (m *Module) fn227(v0 int32) {
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
	m.fn208(t2, t4, t3, v2, i32(4), i32(44))
	{
		t6 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		if t6 != i32(1) {
			goto l0
		}
		t7 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		t8 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		m.fn16(t7, t8)
		panic("unreachable")
	}
l0:
	t9 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	v3 = t9
	store32(m.memory[uint32(v0):], uint32(v2))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
	m.g0 = v1 + i32(16)
}
func (m *Module) fn228(v0, v1, v2, v3 int32) {
	var v4, v5, v6 int32
	{
		{
			p0 := v2 + i32(24)
			if v1 != 0 {
				p0 = v0 + v1*i32(44) + i32(-20)
			}
			v1 = p0
			t1 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			v2 = t1
			if v2 == 0 {
				goto l0
			}
			t2 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			v0 = t2 + v2*i32(44)
			t3 := int32(load32(m.memory[uint32(v0+i32(-44)):]))
			if t3 == i32(-1) {
				t8 := int32(load32(m.memory[int64(uint32(v3))+4:]))
				v2 = t8
				{
					{
						t9 := int32(load32(m.memory[int64(uint32(v3))+8:]))
						v1 = t9
						t10 := v1
						v4 = v0 + i32(-40)
						t11 := int32(load32(m.memory[uint32(v4):]))
						v5 = v0 + i32(-32)
						t12 := int32(load32(m.memory[uint32(v5):]))
						v6 = t12
						if uint32(t10) <= uint32(t11-v6) {
							goto l3
						}
						m.fn197(v4, v6, v1, i32(1), i32(1))
						t13 := int32(load32(m.memory[uint32(v5):]))
						v6 = t13
						goto l4
					}
				l3:
					if v1 == 0 {
						goto l5
					}
				l4:
					if v1 == 0 {
						goto l5
					}
					t14 := int32(load32(m.memory[uint32(v0+i32(-36)):]))
					memory_copy(m.memory, uint32(t14+v6), uint32(v2), uint32(v1))
				}
			l5:
				store32(m.memory[uint32(v5):], uint32(v6+v1))
				{
					t15 := int32(load32(m.memory[uint32(v3):]))
					v1 = t15
					if v1 == 0 {
						return
					}
					t16 := int32(load32(m.memory[uint32(v2+i32(-4)):]))
					v3 = t16
					v0 = v3 & i32(-8)
					t17 := v0
					v3 = v3 & i32(3)
					p18 := i32(8)
					if v3 != 0 {
						p18 = i32(4)
					}
					if uint32(t17) < uint32(p18+v1) {
						m.fn7(i32(1274404), i32(46), i32(1274452))
						panic("unreachable")
					}
					if v3 == 0 {
						goto l8
					}
					if uint32(v0) > uint32(v1+i32(39)) {
						m.fn7(i32(1274468), i32(46), i32(1274516))
						panic("unreachable")
					}
				l8:
					m.fn5(v2)
				}
				return
			}
		}
	l0:
		{
			t4 := int32(load32(m.memory[uint32(v1):]))
			if v2 != t4 {
				goto l2
			}
			m.fn227(v1)
		}
	l2:
		store32(m.memory[int64(uint32(v1))+8:], uint32(v2+i32(1)))
		t5 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		v1 = t5 + v2*i32(44)
		store32(m.memory[uint32(v1):], uint32(i32(-1)))
		t6 := int64(load64(m.memory[uint32(v3):]))
		store64(m.memory[int64(uint32(v1))+4:], uint64(t6))
		t7 := int32(load32(m.memory[int64(uint32(v3))+8:]))
		store32(m.memory[int64(uint32(v1))+12:], uint32(t7))
		return
	}
}
func (m *Module) fn229(v0, v1, v2 int32) {
	var v3, v4, v5 int32
	var v6 int64
	t0 := m.g0
	v3 = t0 - i32(16)
	m.g0 = v3
	t1 := int32(load32(m.memory[int64(uint32(v2))+8:]))
	v4 = t1
	t2 := int32(load32(m.memory[int64(uint32(v2))+4:]))
	v5 = t2
	{
		t3 := int32(load32(m.memory[uint32(v2):]))
		if t3 == i32(-1) {
			m.fn243(v3+i32(4), v1, v5, v4)
			{
				t5 := int32(load32(m.memory[int64(uint32(v3))+4:]))
				if t5 == i32(-2) {
					m.memory[int64(uint32(v0))+8] = byte(i32(2))
					store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
					store32(m.memory[uint32(v0):], uint32(i32(-2)))
					goto l2
				}
				t6 := int32(load32(m.memory[int64(uint32(v3))+12:]))
				store32(m.memory[int64(uint32(v0))+8:], uint32(t6))
				t7 := int64(load64(m.memory[int64(uint32(v3))+4:]))
				store64(m.memory[uint32(v0):], uint64(t7))
				goto l2
			}
		}
		m.fn243(v3+i32(4), v1, v5, v4)
		t4 := int32(load32(m.memory[int64(uint32(v3))+4:]))
		v2 = t4
		if v2 != i32(-2) {
			t8 := int32(load32(m.memory[int64(uint32(v3))+8:]))
			v5 = t8
			t9 := int64(load32(m.memory[int64(uint32(v3))+12:]))
			v6 = t9
			v1 = int32(v6)
			{
				if v2 == i32(-1) {
					goto l4
				}
				v4 = v5
				goto l5
			l4:
				if v1 <= i32(-1) {
					m.fn15()
					panic("unreachable")
				}
				if !(v6 == 0) {
					goto l7
				}
				v2 = i32(0)
				v4 = i32(1)
				v1 = i32(0)
				goto l5
			l7:
				t10 := m.fn11(v1)
				v4 = t10
				if v4 == 0 {
					m.fn16(i32(1), v1)
					panic("unreachable")
				}
				if v1 == 0 {
					goto l9
				}
				memory_copy(m.memory, uint32(v4), uint32(v5), uint32(v1))
			l9:
				v2 = v1
			}
		l5:
			store32(m.memory[int64(uint32(v0))+8:], uint32(v1))
			store32(m.memory[int64(uint32(v0))+4:], uint32(v4))
			store32(m.memory[uint32(v0):], uint32(v2))
			goto l2
		}
		m.memory[int64(uint32(v0))+8] = byte(i32(2))
		store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
		store32(m.memory[uint32(v0):], uint32(i32(-2)))
		goto l2
	}
l2:
	m.g0 = v3 + i32(16)
}
func (m *Module) fn230(v0 int32) {
	var v1, v2, v3, v4 int32
	var v5 int64
	var v6, v7, v8, v9, v10 int32
	var v11 int64
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
		l9:
			{
				if v5 != i64(0) {
					goto l2
				}
			l3:
				{
					v6 = v4
					v4 = v6 + i32(8)
					v3 = v3 + i32(-160)
					t4 := int64(load64(m.memory[uint32(v6):]))
					v5 = t4 & i64(-0x7f7f7f7f7f7f7f80)
					if v5 == i64(-0x7f7f7f7f7f7f7f80) {
						goto l3
					}
				}
				v5 = v5 ^ i64(-0x7f7f7f7f7f7f7f80)
			l2:
				{
					v6 = v3 + (i32(0)-int32(uint32(int64(bits.TrailingZeros64(uint64(v5))))>>3))*i32(20)
					t5 := int32(load32(m.memory[uint32(v6+i32(-20)):]))
					v7 = t5
					if v7 == 0 {
						goto l4
					}
					t6 := int32(load32(m.memory[uint32(v6+i32(-16)):]))
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
				v11 = v5 + i64(-1)
				v9 = v6 + i32(-8)
				t10 := int32(load32(m.memory[uint32(v9):]))
				v7 = t10
				t11 := int32(load32(m.memory[uint32(v7):]))
				t12 := v7
				v7 = t11 + i32(-1)
				store32(m.memory[uint32(t12):], uint32(v7))
				{
					if v7 != 0 {
						goto l8
					}
					t13 := int32(load32(m.memory[uint32(v9):]))
					t14 := int32(load32(m.memory[uint32(v6+i32(-4)):]))
					m.fn146(t13, t14)
				}
			l8:
				v5 = v11 & v5
				v2 = v2 + i32(-1)
				if v2 != 0 {
					goto l9
				}
			}
		}
	l1:
		t15 := v1
		v4 = (v1*i32(20) + i32(27)) & i32(-8)
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
			m.fn7(i32(1274404), i32(46), i32(1274452))
			panic("unreachable")
		}
		if v4 == 0 {
			goto l11
		}
		if uint32(v2) > uint32(v3+i32(39)) {
			m.fn7(i32(1274468), i32(46), i32(1274516))
			panic("unreachable")
		}
	l11:
		m.fn5(v6)
	}
}
func (m *Module) fn231(v0, v1 int32) int32 {
	var v2, v3, v4, v5 int32
	t0 := m.g0
	v2 = t0 - i32(32)
	m.g0 = v2
	{
		t1 := int32(load32(m.memory[uint32(v0):]))
		v3 = t1
		p2 := i32(2)
		if uint32(v3) > uint32(i32(-0x7ffffff9)) {
			p2 = v3 + i32(0x7ffffff8)
		}
		switch p2 {
		default:
			store32(m.memory[int64(uint32(v2))+4:], uint32(v0+i32(4)))
			store64(m.memory[int64(uint32(v2))+16:], uint64(int64(uint32(i32(26)))<<32|int64(uint32(v2+i32(4)))))
			t3 := int32(load32(m.memory[uint32(v1):]))
			t4 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t5 := m.fn46(t3, t4, i32(0x100c77), v2+i32(16))
			v1 = t5
			goto l7
		case 1:
			store32(m.memory[int64(uint32(v2))+4:], uint32(v0+i32(4)))
			store64(m.memory[int64(uint32(v2))+16:], uint64(int64(uint32(i32(27)))<<32|int64(uint32(v2+i32(4)))))
			t6 := int32(load32(m.memory[uint32(v1):]))
			t7 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t8 := m.fn46(t6, t7, i32(1051487), v2+i32(16))
			v1 = t8
			goto l7
		case 2:
			store32(m.memory[int64(uint32(v2))+4:], uint32(v0))
			store64(m.memory[int64(uint32(v2))+16:], uint64(int64(uint32(i32(28)))<<32|int64(uint32(v2+i32(4)))))
			t9 := int32(load32(m.memory[uint32(v1):]))
			t10 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t11 := m.fn46(t9, t10, i32(1051290), v2+i32(16))
			v1 = t11
			goto l7
		case 3:
			store32(m.memory[int64(uint32(v2))+4:], uint32(v0+i32(4)))
			store64(m.memory[int64(uint32(v2))+16:], uint64(int64(uint32(i32(29)))<<32|int64(uint32(v2+i32(4)))))
			t12 := int32(load32(m.memory[uint32(v1):]))
			t13 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t14 := m.fn46(t12, t13, i32(1052122), v2+i32(16))
			v1 = t14
			goto l7
		case 4:
			v3 = v0 + i32(4)
			{
				t15 := int32(m.memory[int64(uint32(v0))+8])
				if t15 != i32(2) {
					store32(m.memory[int64(uint32(v2))+4:], uint32(v3))
					store64(m.memory[int64(uint32(v2))+16:], uint64(int64(uint32(i32(30)))<<32|int64(uint32(v2+i32(4)))))
					t21 := int32(load32(m.memory[uint32(v1):]))
					t22 := int32(load32(m.memory[int64(uint32(v1))+4:]))
					t23 := m.fn46(t21, t22, i32(1052413), v2+i32(16))
					v1 = t23
					goto l7
				}
				t16 := int32(load32(m.memory[uint32(v3):]))
				t17 := int64(load64(m.memory[int64(uint32(t16))+12:]))
				store64(m.memory[int64(uint32(v2))+4:], uint64(t17))
				store64(m.memory[int64(uint32(v2))+16:], uint64(int64(uint32(i32(1)))<<32|int64(uint32(v2+i32(4)))))
				t18 := int32(load32(m.memory[uint32(v1):]))
				t19 := int32(load32(m.memory[int64(uint32(v1))+4:]))
				t20 := m.fn46(t18, t19, i32(1050647), v2+i32(16))
				v1 = t20
				goto l7
			}
		case 5:
			t24 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			v3 = t24
			t25 := int32(load32(m.memory[uint32(v1):]))
			v1 = t25
			{
				t26 := int32(load32(m.memory[int64(uint32(v0))+4:]))
				v4 = t26
				switch v4 >> 31 & (v4 + i32(-0x7fffffff)) {
				default:
					store32(m.memory[int64(uint32(v2))+12:], uint32(v0+i32(16)))
					store32(m.memory[int64(uint32(v2))+4:], uint32(v0+i32(4)))
					store64(m.memory[int64(uint32(v2))+24:], uint64(int64(uint32(i32(31)))<<32|int64(uint32(v2+i32(4)))))
					store64(m.memory[int64(uint32(v2))+16:], uint64(int64(uint32(i32(32)))<<32|int64(uint32(v2+i32(12)))))
					t27 := m.fn46(v1, v3, i32(1066410), v2+i32(16))
					v1 = t27
					goto l7
				case 1:
					store32(m.memory[int64(uint32(v2))+4:], uint32(v0+i32(8)))
					store64(m.memory[int64(uint32(v2))+16:], uint64(int64(uint32(i32(32)))<<32|int64(uint32(v2+i32(4)))))
					t28 := m.fn46(v1, v3, i32(1067624), v2+i32(16))
					v1 = t28
					goto l7
				case 2:
					store32(m.memory[int64(uint32(v2))+4:], uint32(v0+i32(8)))
					store64(m.memory[int64(uint32(v2))+16:], uint64(int64(uint32(i32(33)))<<32|int64(uint32(v2+i32(4)))))
					t29 := m.fn46(v1, v3, i32(1052289), v2+i32(16))
					v1 = t29
					goto l7
				case 3:
					t30 := int32(load32(m.memory[int64(uint32(v3))+12:]))
					t31 := m.t0[uint(t30)].(func(int32, int32, int32) int32)(v1, i32(1274107), i32(46))
					v1 = t31
					goto l7
				}
			}
		case 6:
			t32 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			v4 = t32
			t33 := int32(load32(m.memory[uint32(v1):]))
			v3 = t33
			{
				t34 := int32(load32(m.memory[int64(uint32(v0))+4:]))
				switch t34 {
				default:
					v1 = i32(1)
					t35 := int32(load32(m.memory[int64(uint32(v4))+12:]))
					t36 := v3
					v5 = t35
					t37 := m.t0[uint(v5)].(func(int32, int32, int32) int32)(t36, i32(1273008), i32(26))
					if t37 != 0 {
						goto l7
					}
					t38 := int32(load32(m.memory[int64(uint32(v0))+12:]))
					t39 := int32(load32(m.memory[int64(uint32(v0))+16:]))
					t40 := m.fn674(v3, v4, t38, t39)
					if t40 != 0 {
						goto l7
					}
					t41 := m.t0[uint(v5)].(func(int32, int32, int32) int32)(v3, i32(1272907), i32(1))
					v1 = t41
					goto l7
				case 1:
					v1 = i32(1)
					t42 := int32(load32(m.memory[int64(uint32(v4))+12:]))
					t43 := v3
					v5 = t42
					t44 := m.t0[uint(v5)].(func(int32, int32, int32) int32)(t43, i32(1273034), i32(47))
					if t44 != 0 {
						goto l7
					}
					t45 := int32(load32(m.memory[int64(uint32(v0))+12:]))
					t46 := int32(load32(m.memory[int64(uint32(v0))+16:]))
					t47 := m.fn674(v3, v4, t45, t46)
					if t47 != 0 {
						goto l7
					}
					t48 := m.t0[uint(v5)].(func(int32, int32, int32) int32)(v3, i32(1272907), i32(1))
					v1 = t48
					goto l7
				case 2:
					v1 = i32(1)
					t49 := int32(load32(m.memory[int64(uint32(v4))+12:]))
					t50 := v3
					v5 = t49
					t51 := m.t0[uint(v5)].(func(int32, int32, int32) int32)(t50, i32(1273081), i32(49))
					if t51 != 0 {
						goto l7
					}
					t52 := int32(load32(m.memory[int64(uint32(v0))+12:]))
					t53 := int32(load32(m.memory[int64(uint32(v0))+16:]))
					t54 := m.fn674(v3, v4, t52, t53)
					if t54 != 0 {
						goto l7
					}
					t55 := m.t0[uint(v5)].(func(int32, int32, int32) int32)(v3, i32(1272907), i32(1))
					v1 = t55
					goto l7
				case 3:
					v1 = i32(1)
					t56 := int32(load32(m.memory[int64(uint32(v4))+12:]))
					t57 := v3
					v5 = t56
					t58 := m.t0[uint(v5)].(func(int32, int32, int32) int32)(t57, i32(1273130), i32(22))
					if t58 != 0 {
						goto l7
					}
					t59 := int32(load32(m.memory[int64(uint32(v0))+12:]))
					t60 := int32(load32(m.memory[int64(uint32(v0))+16:]))
					t61 := m.fn674(v3, v4, t59, t60)
					if t61 != 0 {
						goto l7
					}
					t62 := m.t0[uint(v5)].(func(int32, int32, int32) int32)(v3, i32(1273152), i32(59))
					v1 = t62
					goto l7
				case 4:
					v1 = i32(1)
					t63 := int32(load32(m.memory[int64(uint32(v4))+12:]))
					t64 := v3
					v5 = t63
					t65 := m.t0[uint(v5)].(func(int32, int32, int32) int32)(t64, i32(1273130), i32(22))
					if t65 != 0 {
						goto l7
					}
					t66 := int32(load32(m.memory[int64(uint32(v0))+12:]))
					t67 := int32(load32(m.memory[int64(uint32(v0))+16:]))
					t68 := m.fn674(v3, v4, t66, t67)
					if t68 != 0 {
						goto l7
					}
					t69 := m.t0[uint(v5)].(func(int32, int32, int32) int32)(v3, i32(1273211), i32(52))
					v1 = t69
					goto l7
				case 5:
					store32(m.memory[int64(uint32(v2))+4:], uint32(v0+i32(8)))
					store64(m.memory[int64(uint32(v2))+16:], uint64(int64(uint32(i32(34)))<<32|int64(uint32(v2+i32(4)))))
					t70 := m.fn46(v3, v4, i32(1052762), v2+i32(16))
					v1 = t70
				}
			}
		}
	}
l7:
	m.g0 = v2 + i32(32)
	return v1
}
func (m *Module) fn232(v0 int32) {
	var v1, v2, v3 int32
	{
		t0 := int32(load32(m.memory[uint32(v0):]))
		v1 = t0
		p1 := i32(2)
		if uint32(v1) > uint32(i32(-0x7ffffff9)) {
			p1 = v1 + i32(0x7ffffff8)
		}
		switch p1 {
		case 1, 3, 4:
			return
		default:
			t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			switch t2 {
			default:
				return
			case 0, 1, 2, 3, 4:
				t3 := int32(load32(m.memory[int64(uint32(v0))+8:]))
				v1 = t3
				if v1 == 0 {
					return
				}
				t4 := int32(load32(m.memory[int64(uint32(v0))+12:]))
				v2 = t4
				t5 := int32(load32(m.memory[uint32(v2+i32(-4)):]))
				v0 = t5
				v3 = v0 & i32(-8)
				t6 := v3
				v0 = v0 & i32(3)
				p7 := i32(8)
				if v0 != 0 {
					p7 = i32(4)
				}
				if uint32(t6) < uint32(p7+v1) {
					m.fn7(i32(1274404), i32(46), i32(1274452))
					panic("unreachable")
				}
				if v0 == 0 {
					goto l7
				}
				if uint32(v3) > uint32(v1+i32(39)) {
					m.fn7(i32(1274468), i32(46), i32(1274516))
					panic("unreachable")
				}
			l7:
				m.fn5(v2)
				return
			}
		case 0:
			t8 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			v1 = t8
			t9 := int32(load32(m.memory[uint32(v1):]))
			t10 := v1
			v1 = t9
			store32(m.memory[uint32(t10):], uint32(v1+i32(-1)))
			if v1 != i32(1) {
				return
			}
			t11 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			m.fn237(t11)
			return
		case 2:
			p12 := i32(5)
			if v1 < i32(0) {
				p12 = v1 ^ i32(-0x80000000)
			}
			switch p12 {
			default:
				return
			case 0:
				t13 := int32(load32(m.memory[int64(uint32(v0))+4:]))
				v1 = t13
				if v1 <= i32(0) {
					return
				}
				v2 = i32(8)
				goto l13
			case 3:
				t14 := int32(load32(m.memory[int64(uint32(v0))+4:]))
				v1 = t14
				if v1 == 0 {
					return
				}
				v2 = i32(8)
				goto l13
			case 4:
				t15 := int32(load32(m.memory[int64(uint32(v0))+4:]))
				v1 = t15
				if v1 == 0 {
					return
				}
				v2 = i32(8)
				goto l13
			case 5:
				{
					if v1 == 0 {
						goto l14
					}
					t16 := int32(load32(m.memory[int64(uint32(v0))+4:]))
					m.fn21(t16, v1, i32(1))
				}
			l14:
				t17 := int32(load32(m.memory[int64(uint32(v0))+12:]))
				v1 = t17
				if v1 == 0 {
					return
				}
				v2 = i32(16)
				goto l13
			}
		case 5:
			t18 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			v1 = t18
			if v1 < i32(1) {
				return
			}
			t19 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			v2 = t19
			t20 := int32(load32(m.memory[uint32(v2+i32(-4)):]))
			v0 = t20
			v3 = v0 & i32(-8)
			t21 := v3
			v0 = v0 & i32(3)
			p22 := i32(8)
			if v0 != 0 {
				p22 = i32(4)
			}
			if uint32(t21) < uint32(p22+v1) {
				m.fn7(i32(1274404), i32(46), i32(1274452))
				panic("unreachable")
			}
			if v0 == 0 {
				goto l16
			}
			if uint32(v3) > uint32(v1+i32(39)) {
				m.fn7(i32(1274468), i32(46), i32(1274516))
				panic("unreachable")
			}
		l16:
			m.fn5(v2)
			return
		}
	}
l13:
	{
		t23 := int32(load32(m.memory[uint32(v0+v2):]))
		v2 = t23
		t24 := int32(load32(m.memory[uint32(v2+i32(-4)):]))
		v0 = t24
		v3 = v0 & i32(-8)
		t25 := v3
		v0 = v0 & i32(3)
		p26 := i32(8)
		if v0 != 0 {
			p26 = i32(4)
		}
		if uint32(t25) < uint32(p26+v1) {
			m.fn7(i32(1274404), i32(46), i32(1274452))
			panic("unreachable")
		}
		if v0 == 0 {
			goto l19
		}
		if uint32(v3) > uint32(v1+i32(39)) {
			m.fn7(i32(1274468), i32(46), i32(1274516))
			panic("unreachable")
		}
	l19:
		m.fn5(v2)
		return
	}
}
func (m *Module) fn233(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := v1
	v0 = t0
	t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t3 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	t4 := m.fn9(t1, t2, t3)
	return t4
}
func (m *Module) fn234(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := v1
	v0 = t0
	t2 := int32(load32(m.memory[uint32(v0):]))
	t3 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t4 := m.fn9(t1, t2, t3)
	return t4
}
func (m *Module) fn235(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := m.fn255(t0, v1)
	return t1
}
func (m *Module) fn236(v0, v1, v2, v3 int32) {
	var v4, v5, v6, v7, v8 int32
	{
		v4 = (v2 + i32(3)) & i32(-4)
		if v4 != v2 {
			goto l0
		}
		v5 = v3 + i32(-8)
		v4 = i32(0)
		goto l1
	l0:
		t0 := v3
		v4 = v4 - v2
		p1 := v4
		if uint32(v3) < uint32(v4) {
			p1 = t0
		}
		v4 = p1
		if v3 == 0 {
			goto l2
		}
		v6 = i32(0)
		v7 = v1 & i32(255)
		v8 = i32(1)
	l4:
		{
			t2 := int32(m.memory[uint32(v2+v6)])
			if t2 == v7 {
				goto l3
			}
			t3 := v4
			v6 = v6 + i32(1)
			if t3 != v6 {
				goto l4
			}
		}
	l2:
		t4 := v4
		v5 = v3 + i32(-8)
		if uint32(t4) > uint32(v5) {
			goto l5
		}
	}
l1:
	v6 = v1 & i32(255) * i32(16843009)
l6:
	{
		v7 = v2 + v4
		t5 := int32(load32(m.memory[uint32(v7):]))
		v8 = t5 ^ v6
		t6 := int32(load32(m.memory[uint32(v7+i32(4)):]))
		t7 := i32(16843008) - v8 | v8
		v7 = t6 ^ v6
		if t7&(i32(16843008)-v7|v7)&i32(-2139062144) != i32(-2139062144) {
			goto l5
		}
		v4 = v4 + i32(8)
		if uint32(v4) <= uint32(v5) {
			goto l6
		}
	}
l5:
	if v3 == v4 {
		goto l7
	}
	v5 = v3 - v4
	v7 = v2 + v4
	v6 = i32(0)
	v8 = v1 & i32(255)
l9:
	{
		t8 := int32(m.memory[uint32(v7+v6)])
		if t8 == v8 {
			v6 = v6 + v4
			v8 = i32(1)
			goto l3
		}
		t9 := v5
		v6 = v6 + i32(1)
		if t9 == v6 {
			goto l7
		}
		goto l9
	}
l7:
	v8 = i32(0)
l3:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v6))
	store32(m.memory[uint32(v0):], uint32(v8))
}
func (m *Module) fn237(v0 int32) {
	var v1, v2, v3, v4, v5 int32
	{
		t0 := int32(m.memory[int64(uint32(v0))+8])
		if t0 != i32(3) {
			goto l0
		}
		t1 := int32(load32(m.memory[uint32(v0+i32(12)):]))
		v1 = t1
		t2 := int32(load32(m.memory[uint32(v1):]))
		v2 = t2
		{
			t3 := int32(load32(m.memory[uint32(v1+i32(4)):]))
			v3 = t3
			t4 := int32(load32(m.memory[uint32(v3):]))
			v4 = t4
			if v4 == 0 {
				goto l1
			}
			m.t0[uint(v4)].(func(int32))(v2)
		}
	l1:
		{
			t5 := int32(load32(m.memory[int64(uint32(v3))+4:]))
			v3 = t5
			if v3 == 0 {
				goto l2
			}
			t6 := int32(load32(m.memory[uint32(v2+i32(-4)):]))
			v4 = t6
			v5 = v4 & i32(-8)
			t7 := v5
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
				goto l4
			}
			if uint32(v5) > uint32(v3+i32(39)) {
				m.fn7(i32(1274468), i32(46), i32(1274516))
				panic("unreachable")
			}
		l4:
			m.fn5(v2)
		}
	l2:
		t9 := int32(load32(m.memory[uint32(v1+i32(-4)):]))
		v2 = t9
		v3 = v2 & i32(-8)
		t10 := v3
		v2 = v2 & i32(3)
		p11 := i32(20)
		if v2 != 0 {
			p11 = i32(16)
		}
		if uint32(t10) < uint32(p11) {
			m.fn7(i32(1274404), i32(46), i32(1274452))
			panic("unreachable")
		}
		if v2 == 0 {
			goto l7
		}
		if uint32(v3) >= uint32(i32(52)) {
			m.fn7(i32(1274468), i32(46), i32(1274516))
			panic("unreachable")
		}
	l7:
		m.fn5(v1)
	}
l0:
	{
		if v0 == i32(-1) {
			return
		}
		t12 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		t13 := v0
		v1 = t12
		store32(m.memory[int64(uint32(t13))+4:], uint32(v1+i32(-1)))
		if v1 != i32(1) {
			return
		}
		t14 := int32(load32(m.memory[uint32(v0+i32(-4)):]))
		v1 = t14
		t15 := v1 & i32(-8)
		v2 = v1 & i32(3)
		p16 := i32(24)
		if v2 != 0 {
			p16 = i32(20)
		}
		if uint32(t15) < uint32(p16) {
			m.fn7(i32(1274404), i32(46), i32(1274452))
			panic("unreachable")
		}
		if v2 == 0 {
			goto l11
		}
		if uint32(v1) >= uint32(i32(56)) {
			m.fn7(i32(1274468), i32(46), i32(1274516))
			panic("unreachable")
		}
	l11:
		m.fn5(v0)
	}
}
func (m *Module) fn238(v0, v1, v2 int32) {
	var v3, v4, v5, v6 int32
	{
		if v2 == 0 {
			goto l0
		}
		if uint32(v2) < uint32(i32(4)) {
			v3 = v1
			t5 := int32(m.memory[uint32(v1)])
			if t5 == i32(58) {
				goto l3
			}
			if v2 == i32(1) {
				goto l0
			}
			{
				t6 := int32(m.memory[int64(uint32(v1))+1])
				if t6 != i32(58) {
					if v2 == i32(2) {
						goto l0
					}
					t7 := int32(m.memory[int64(uint32(v1))+2])
					if t7 != i32(58) {
						goto l0
					}
					v3 = v1 + i32(2)
					goto l3
				}
				v3 = v1 + i32(1)
				goto l3
			}
		}
		{
			t0 := int32(load32(m.memory[uint32(v1):]))
			v3 = t0
			if (i32(16843008)-(v3^i32(976894522))|v3)&i32(-2139062144) == i32(-2139062144) {
				t2 := v1
				v5 = v1 & i32(3)
				v4 = i32(4) - v5
				v3 = t2 + v4
				if uint32(v2) < uint32(i32(9)) {
					if uint32(v4) >= uint32(v2) {
						goto l0
					}
					v4 = v2 + v5 + i32(-4)
				l9:
					{
						t8 := int32(m.memory[uint32(v3)])
						if t8 == i32(58) {
							goto l3
						}
						v3 = v3 + i32(1)
						v4 = v4 + i32(-1)
						if v4 == 0 {
							goto l0
						}
						goto l9
					}
				}
				v5 = v1 + v2
				if uint32(v4) > uint32(v2+i32(-8)) {
					goto l6
				}
				v6 = v5 + i32(-8)
			l7:
				{
					t3 := int32(load32(m.memory[uint32(v3):]))
					v4 = t3
					if (i32(16843008)-(v4^i32(976894522))|v4)&i32(-2139062144) != i32(-2139062144) {
						goto l6
					}
					t4 := int32(load32(m.memory[uint32(v3+i32(4)):]))
					v4 = t4
					if (i32(16843008)-(v4^i32(976894522))|v4)&i32(-2139062144) != i32(-2139062144) {
						goto l6
					}
					v3 = v3 + i32(8)
					if uint32(v3) <= uint32(v6) {
						goto l7
					}
					goto l6
				}
			}
			v4 = v2
			v3 = v1
		l4:
			{
				t1 := int32(m.memory[uint32(v3)])
				if t1 == i32(58) {
					goto l3
				}
				v3 = v3 + i32(1)
				v4 = v4 + i32(-1)
				if v4 != 0 {
					goto l4
				}
				goto l0
			}
		}
	l6:
		if uint32(v3) >= uint32(v5) {
			goto l0
		}
	l10:
		{
			t9 := int32(m.memory[uint32(v3)])
			if t9 == i32(58) {
				goto l3
			}
			v3 = v3 + i32(1)
			if v3 == v5 {
				goto l0
			}
			goto l10
		}
	l3:
		store32(m.memory[int64(uint32(v0))+8:], uint32(v1))
		t10 := v0
		v4 = v3 - v1
		store32(m.memory[int64(uint32(t10))+12:], uint32(v4))
		store32(m.memory[uint32(v0):], uint32(v3+i32(1)))
		store32(m.memory[int64(uint32(v0))+4:], uint32(v4^i32(-1)+v2))
		return
	}
l0:
	store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v2))
	store32(m.memory[uint32(v0):], uint32(v1))
}
func (m *Module) fn239(v0, v1, v2, v3, v4 int32) {
	var v5, v6, v7, v8, v9 int32
	t0 := int32(load32(m.memory[int64(uint32(v1))+20:]))
	v5 = t0
	v6 = v5 << 4
	t1 := int32(load32(m.memory[int64(uint32(v1))+16:]))
	v7 = t1
	{
		{
			if v2 == 0 {
				if v4 == 0 {
					store32(m.memory[uint32(v0):], uint32(i32(-0x80000000)))
					return
				}
			l9:
				{
					if v6 == 0 {
						store32(m.memory[uint32(v0):], uint32(i32(-0x80000000)))
						return
					}
					v5 = v7 + v6
					v4 = v6 + i32(-16)
					v6 = v4
					t8 := int32(load32(m.memory[uint32(v5+i32(-12)):]))
					if t8 != 0 {
						goto l9
					}
				}
				v5 = v7 + v4
				t9 := int32(load32(m.memory[uint32(v5+i32(8)):]))
				v6 = t9
				if v6 != 0 {
					t10 := int32(load32(m.memory[int64(uint32(v1))+8:]))
					v4 = t10
					t11 := int32(load32(m.memory[uint32(v5):]))
					t12 := v4
					v5 = t11
					if uint32(t12) < uint32(v5) {
						m.fn28(i32(1272348), i32(19), i32(1272780))
						panic("unreachable")
					}
					if uint32(v6) > uint32(v4-v5) {
						m.fn28(i32(1272348), i32(19), i32(1272796))
						panic("unreachable")
					}
					store32(m.memory[int64(uint32(v0))+8:], uint32(v6))
					store32(m.memory[uint32(v0):], uint32(i32(-0x7fffffff)))
					t13 := int32(load32(m.memory[int64(uint32(v1))+4:]))
					store32(m.memory[int64(uint32(v0))+4:], uint32(t13+v5))
					return
				}
				store32(m.memory[uint32(v0):], uint32(i32(-0x80000000)))
				return
			}
			if v5 == 0 {
				goto l1
			}
			t2 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			v8 = t2
			t3 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			v9 = t3
		l6:
			{
				v1 = v7 + v6
				t4 := int32(load32(m.memory[uint32(v1+i32(-12)):]))
				v5 = t4
				if v5 == 0 {
					goto l2
				}
				t5 := int32(load32(m.memory[uint32(v1+i32(-16)):]))
				t6 := v9
				v4 = t5
				if uint32(t6) < uint32(v4) {
					m.fn28(i32(1272348), i32(19), i32(1272300))
					panic("unreachable")
				}
				if uint32(v5) > uint32(v9-v4) {
					m.fn28(i32(1272348), i32(19), i32(1272316))
					panic("unreachable")
				}
				if v5 != v3 {
					goto l2
				}
				t7 := m.fn1909(v8+v4, v2, v3)
				if t7 == 0 {
					goto l5
				}
			}
		l2:
			v6 = v6 + i32(-16)
			if v6 != 0 {
				goto l6
			}
			goto l1
		}
	l5:
		t14 := int32(load32(m.memory[uint32(v1+i32(-8)):]))
		v6 = t14
		if v6 != 0 {
			t16 := v9
			v5 = v3 + v4
			if uint32(t16) < uint32(v5) {
				m.fn28(i32(1272348), i32(19), i32(1272780))
				panic("unreachable")
			}
			if uint32(v6) > uint32(v9-v5) {
				m.fn28(i32(1272348), i32(19), i32(1272796))
				panic("unreachable")
			}
			store32(m.memory[int64(uint32(v0))+8:], uint32(v6))
			store32(m.memory[uint32(v0):], uint32(i32(-0x7fffffff)))
			store32(m.memory[int64(uint32(v0))+4:], uint32(v8+v5))
			return
		}
	}
l1:
	if v3 <= i32(-1) {
		m.fn15()
		panic("unreachable")
	}
	{
		if v3 != 0 {
			goto l15
		}
		v6 = i32(1)
		goto l16
	l15:
		t15 := m.fn11(v3)
		v6 = t15
		if v6 == 0 {
			m.fn16(i32(1), v3)
			panic("unreachable")
		}
		if v3 == 0 {
			goto l16
		}
		memory_copy(m.memory, uint32(v6), uint32(v2), uint32(v3))
	}
l16:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v3))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v6))
	store32(m.memory[uint32(v0):], uint32(v3))
}
func (m *Module) fn240(v0, v1, v2, v3 int32) {
	var v4 int32
	var v5 int64
	var v6, v7 int32
	var v8 int64
	var v9, v10 int32
	var v11 int64
	var v12, v13, v14, v15 int32
	var v16 int64
	var v17, v18 int32
	t0 := m.g0
	v4 = t0 - i32(16)
	m.g0 = v4
	{
		{
			{
				{
					t1 := int32(load32(m.memory[int64(uint32(v1))+12:]))
					if t1 == 0 {
						goto l0
					}
					t2 := int64(load64(m.memory[int64(uint32(v1))+16:]))
					t3 := int64(load64(m.memory[int64(uint32(v1))+24:]))
					t4 := m.fn207(t2, t3, v2, v3)
					v5 = t4
					t5 := int32(load32(m.memory[int64(uint32(v1))+4:]))
					v6 = t5
					v7 = v6 & int32(v5)
					v8 = int64(uint64(v5)>>25) & i64(127) * i64(72340172838076673)
					t6 := int32(load32(m.memory[uint32(v1):]))
					v9 = t6
					v10 = i32(0)
				l5:
					{
						{
							t7 := int64(load64(m.memory[uint32(v9+v7):]))
							v11 = t7
							v5 = v11 ^ v8
							v5 = (v5 ^ i64(-1)) & (v5 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
							if v5 == 0 {
								goto l1
							}
						l4:
							{
								t8 := v3
								v12 = v9 + (i32(0)-(int32(uint32(int64(bits.TrailingZeros64(uint64(v5))))>>3)+v7)&v6)*i32(20)
								t9 := int32(load32(m.memory[uint32(v12+i32(-12)):]))
								if t8 != t9 {
									goto l2
								}
								t10 := int32(load32(m.memory[uint32(v12+i32(-16)):]))
								t11 := m.fn1909(v2, t10, v3)
								if t11 == 0 {
									t16 := int32(load32(m.memory[uint32(v12+i32(-4)):]))
									v7 = t16
									t17 := int32(load32(m.memory[uint32(v12+i32(-8)):]))
									v9 = t17
									t18 := int32(load32(m.memory[uint32(v9):]))
									t19 := v9
									v1 = t18 + i32(1)
									store32(m.memory[uint32(t19):], uint32(v1))
									if v1 == 0 {
										goto l8
									}
									goto l9
								}
							}
						l2:
							v5 = (v5 + i64(-1)) & v5
							if !(v5 == 0) {
								goto l4
							}
						}
					l1:
						if !(v11&(v11<<1)&i64(-0x7f7f7f7f7f7f7f80) == 0) {
							goto l0
						}
						t12 := v7
						v10 = v10 + i32(8)
						v7 = (t12 + v10) & v6
						goto l5
					}
				}
			l0:
				m.fn29(v4, v2, v3)
				t13 := int32(load32(m.memory[int64(uint32(v4))+8:]))
				v7 = t13
				t14 := int32(load32(m.memory[int64(uint32(v4))+4:]))
				v12 = t14
				t15 := int32(load32(m.memory[uint32(v4):]))
				v6 = t15
				if v6 == i32(-1) {
					goto l6
				}
				v9 = v12
				goto l7
			}
		l6:
			if v7 <= i32(-1) {
				m.fn15()
				panic("unreachable")
			}
			if v7 != 0 {
				goto l11
			}
			v9 = i32(1)
			v7 = i32(0)
			v6 = i32(0)
			goto l7
		l11:
			t20 := m.fn11(v7)
			v9 = t20
			if v9 == 0 {
				m.fn16(i32(1), v7)
				panic("unreachable")
			}
			if v7 == 0 {
				goto l13
			}
			memory_copy(m.memory, uint32(v9), uint32(v12), uint32(v7))
		l13:
			v6 = v7
		}
	l7:
		m.fn198(v4, v9, v7)
		{
			{
				t21 := int32(load32(m.memory[uint32(v4):]))
				v12 = t21
				if v12 != i32(-1) {
					goto l14
				}
				v10 = v9
				v12 = v6
				goto l15
			}
		l14:
			t22 := int32(load32(m.memory[int64(uint32(v4))+8:]))
			v7 = t22
			t23 := int32(load32(m.memory[int64(uint32(v4))+4:]))
			v10 = t23
			if v6 == 0 {
				goto l15
			}
			t24 := int32(load32(m.memory[uint32(v9+i32(-4)):]))
			v13 = t24
			v14 = v13 & i32(-8)
			t25 := v14
			v13 = v13 & i32(3)
			p26 := i32(8)
			if v13 != 0 {
				p26 = i32(4)
			}
			if uint32(t25) < uint32(p26+v6) {
				m.fn7(i32(1274404), i32(46), i32(1274452))
				panic("unreachable")
			}
			if v13 == 0 {
				goto l17
			}
			if uint32(v14) > uint32(v6+i32(39)) {
				m.fn7(i32(1274468), i32(46), i32(1274516))
				panic("unreachable")
			}
		l17:
			m.fn5(v9)
		}
	l15:
		if uint32(v7) >= uint32(i32(0x7ffffff5)) {
			m.fn42(i32(1284936), i32(43), v4+i32(15), i32(1068072), i32(1068088))
			panic("unreachable")
		}
		v6 = (v7 + i32(11)) & i32(0x7ffffffc)
		t27 := m.fn11(v6)
		v9 = t27
		if v9 == 0 {
			m.fn23(i32(4), v6)
			panic("unreachable")
		}
		store64(m.memory[uint32(v9):], uint64(i64(0x100000001)))
		if v7 == 0 {
			goto l21
		}
		memory_copy(m.memory, uint32(v9+i32(8)), uint32(v10), uint32(v7))
	l21:
		{
			if v12 == 0 {
				goto l22
			}
			t28 := int32(load32(m.memory[uint32(v10+i32(-4)):]))
			v6 = t28
			v13 = v6 & i32(-8)
			t29 := v13
			v6 = v6 & i32(3)
			p30 := i32(8)
			if v6 != 0 {
				p30 = i32(4)
			}
			if uint32(t29) < uint32(p30+v12) {
				m.fn7(i32(1274404), i32(46), i32(1274452))
				panic("unreachable")
			}
			if v6 == 0 {
				goto l24
			}
			if uint32(v13) > uint32(v12+i32(39)) {
				goto l25
			}
		l24:
			m.fn5(v10)
		}
	l22:
		{
			{
				if v3 != 0 {
					goto l26
				}
				v2 = i32(0)
				v12 = i32(1)
				goto l27
			l26:
				t31 := m.fn11(v3)
				v12 = t31
				if v12 == 0 {
					m.fn16(i32(1), v3)
					panic("unreachable")
				}
				if v3 == 0 {
					goto l29
				}
				memory_copy(m.memory, uint32(v12), uint32(v2), uint32(v3))
			l29:
				v2 = v3
			}
		l27:
			t32 := int32(load32(m.memory[uint32(v9):]))
			t33 := v9
			v6 = t32 + i32(1)
			store32(m.memory[uint32(t33):], uint32(v6))
			if v6 == 0 {
				goto l8
			}
			t34 := int64(load64(m.memory[int64(uint32(v1))+16:]))
			t35 := int64(load64(m.memory[int64(uint32(v1))+24:]))
			t36 := m.fn61(t34, t35, v12, v2)
			v5 = t36
			{
				t37 := int32(load32(m.memory[int64(uint32(v1))+8:]))
				if t37 != 0 {
					goto l30
				}
				_ = m.fn60(v1, v1+i32(16))
			}
		l30:
			t39 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			v15 = t39
			v10 = v15 & int32(v5)
			v16 = int64(uint64(v5) >> 25)
			v8 = v16 & i64(127) * i64(72340172838076673)
			t40 := int32(load32(m.memory[uint32(v1):]))
			v6 = t40
			v17 = i32(0)
			v18 = i32(0)
		l44:
			{
				t41 := int64(load64(m.memory[uint32(v6+v10):]))
				v11 = t41
				v5 = v11 ^ v8
				v5 = (v5 ^ i64(-1)) & (v5 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
				if v5 == 0 {
					goto l31
				}
			l34:
				{
					t42 := v2
					v13 = v6 + (i32(0)-(int32(uint32(int64(bits.TrailingZeros64(uint64(v5))))>>3)+v10)&v15)*i32(20)
					t43 := int32(load32(m.memory[uint32(v13+i32(-12)):]))
					if t42 != t43 {
						goto l32
					}
					t44 := int32(load32(m.memory[uint32(v13+i32(-16)):]))
					t45 := m.fn1909(v12, t44, v2)
					if t45 == 0 {
						v1 = v13 + i32(-4)
						t53 := int32(load32(m.memory[uint32(v1):]))
						v6 = t53
						store32(m.memory[uint32(v1):], uint32(v7))
						v2 = v13 + i32(-8)
						t54 := int32(load32(m.memory[uint32(v2):]))
						v1 = t54
						store32(m.memory[uint32(v2):], uint32(v9))
						{
							if v3 == 0 {
								goto l40
							}
							t55 := int32(load32(m.memory[uint32(v12+i32(-4)):]))
							v2 = t55
							v10 = v2 & i32(-8)
							t56 := v10
							v2 = v2 & i32(3)
							p57 := i32(8)
							if v2 != 0 {
								p57 = i32(4)
							}
							if uint32(t56) < uint32(p57+v3) {
								m.fn7(i32(1274404), i32(46), i32(1274452))
								panic("unreachable")
							}
							if v2 == 0 {
								goto l42
							}
							if uint32(v10) > uint32(v3+i32(39)) {
								m.fn7(i32(1274468), i32(46), i32(1274516))
								panic("unreachable")
							}
						l42:
							m.fn5(v12)
						}
					l40:
						t58 := int32(load32(m.memory[uint32(v1):]))
						t59 := v1
						v3 = t58 + i32(-1)
						store32(m.memory[uint32(t59):], uint32(v3))
						if v3 != 0 {
							goto l9
						}
						m.fn146(v1, v6)
						goto l9
					}
				}
			l32:
				v5 = (v5 + i64(-1)) & v5
				if !(v5 == 0) {
					goto l34
				}
			}
		l31:
			v5 = v11 & i64(-0x7f7f7f7f7f7f7f80)
			if v17 == i32(1) {
				goto l35
			}
			if v5 == 0 {
				v17 = i32(0)
				goto l38
			}
			v14 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v5))))>>3) + v10) & v15
		l35:
			if v5&(v11<<1) != i64(0) {
				{
					t46 := int32(int8(m.memory[uint32(v6+v14)]))
					v3 = t46
					if v3 < i32(0) {
						goto l39
					}
					t47 := int64(load64(m.memory[uint32(v6):]))
					t48 := v6
					v14 = int32(uint32(int64(bits.TrailingZeros64(uint64(t47&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
					t49 := int32(m.memory[uint32(t48+v14)])
					v3 = t49
				}
			l39:
				t50 := v6 + v14
				v10 = int32(v16) & i32(127)
				m.memory[uint32(t50)] = byte(v10)
				m.memory[uint32(v6+(v14+i32(-8))&v15+i32(8))] = byte(v10)
				t51 := int32(load32(m.memory[int64(uint32(v1))+8:]))
				store32(m.memory[int64(uint32(v1))+8:], uint32(t51-v3&i32(1)))
				t52 := int32(load32(m.memory[int64(uint32(v1))+12:]))
				store32(m.memory[int64(uint32(v1))+12:], uint32(t52+i32(1)))
				v1 = v6 + (i32(0)-v14)*i32(20)
				store32(m.memory[uint32(v1+i32(-20)):], uint32(v2))
				store32(m.memory[uint32(v1+i32(-16)):], uint32(v12))
				store32(m.memory[uint32(v1+i32(-12)):], uint32(v2))
				store32(m.memory[uint32(v1+i32(-8)):], uint32(v9))
				store32(m.memory[uint32(v1+i32(-4)):], uint32(v7))
				goto l9
			}
			v17 = i32(1)
			goto l38
		l38:
			v18 = v18 + i32(8)
			v10 = (v18 + v10) & v15
			goto l44
		}
	l25:
		m.fn7(i32(1274468), i32(46), i32(1274516))
	}
l8:
	panic("unreachable")
l9:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v7))
	store32(m.memory[uint32(v0):], uint32(v9))
	m.g0 = v4 + i32(16)
}
func (m *Module) fn241(v0, v1, v2, v3 int32) {
	var v4, v5, v6, v7, v8, v9, v10, v11, v12, v13 int32
	t0 := m.g0
	v4 = t0 - i32(16)
	m.g0 = v4
	t1 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	v5 = t1
	{
		t2 := int32(load32(m.memory[uint32(v1):]))
		switch t2 {
		case 1:
			goto l1
		default:
			goto l0
		case 2:
			if uint32(v5) > uint32(v3) {
				m.fn121(v5, v3, v3, i32(1272708))
				panic("unreachable")
			}
		l5:
			{
				if v3 == v5 {
					goto l0
				}
				v6 = v2 + v5
				v5 = v5 + i32(1)
				t3 := int32(m.memory[uint32(v6)])
				v6 = t3 + i32(-9)
				if uint32(v6) > uint32(i32(23)) {
					goto l5
				}
				if i32_shl(i32(1), v6)&i32(8388627) == 0 {
					goto l5
				}
			}
			v5 = v5 + i32(-1)
			goto l1
		case 3:
			if uint32(v5) > uint32(v3) {
				m.fn121(v5, v3, v3, i32(1272724))
				panic("unreachable")
			}
			if v5 == v3 {
				goto l0
			}
			v7 = i32(0) - v3
		l12:
			{
				{
					v8 = v2 + v5
					t4 := int32(m.memory[uint32(v8)])
					v9 = t4
					v6 = v9 + i32(-9)
					if uint32(v6) > uint32(i32(30)) {
						goto l7
					}
					if i32_shl(i32(1), v6)&i32(8388627) != 0 {
						goto l8
					}
					if i32_shl(i32(1), v6)&i32(0x42000000) == 0 {
						goto l7
					}
					v6 = v3 + i32(-1)
					v7 = i32(0)
				l9:
					{
						if v5 == v6 {
							goto l0
						}
						v10 = v8 + v7
						v6 = v6 + i32(-1)
						v7 = v7 + i32(1)
						t5 := int32(m.memory[uint32(v10+i32(1))])
						if t5 != v9 {
							goto l9
						}
					}
					v5 = v5 + v7
					goto l1
				}
			l7:
				if uint32(v5) > uint32(v3) {
					m.fn121(v5, v3, v3, i32(1272708))
					panic("unreachable")
				}
				v7 = v5 - v3
				v6 = i32(0)
			l11:
				{
					if v7+v6 == 0 {
						goto l0
					}
					v9 = v8 + v6
					v6 = v6 + i32(1)
					t6 := int32(m.memory[uint32(v9)])
					v9 = t6 + i32(-9)
					if uint32(v9) > uint32(i32(23)) {
						goto l11
					}
					if i32_shl(i32(1), v9)&i32(8388627) == 0 {
						goto l11
					}
				}
				v5 = v5 + v6 + i32(-1)
				goto l1
			l8:
				t7 := v7
				v5 = v5 + i32(1)
				if t7+v5 != 0 {
					goto l12
				}
			}
		}
	}
l0:
	store32(m.memory[uint32(v0):], uint32(i32(-2)))
	goto l13
l1:
	if uint32(v5) > uint32(v3) {
		m.fn121(v5, v3, v3, i32(1272756))
		panic("unreachable")
	}
	if v5 == v3 {
		goto l15
	}
	v11 = v2 + v3
l47:
	{
		{
			t8 := int32(m.memory[uint32(v2+v5)])
			v6 = t8 + i32(-9)
			if uint32(v6) > uint32(i32(23)) {
				goto l16
			}
			if i32_shl(i32(1), v6)&i32(8388627) != 0 {
				goto l17
			}
		}
	l16:
		v12 = v3 + i32(-1)
		v7 = v2 + i32(1)
		v8 = v5
	l35:
		{
			if v12 == v8 {
				v6 = i32(0)
				store32(m.memory[uint32(v1):], uint32(i32(0)))
				v8 = i32(-1)
				{
					t10 := int32(m.memory[int64(uint32(v1))+36])
					if t10 == i32(1) {
						m.fn1825(v4+i32(4), v1, v2, v3, v5, v3)
						{
							t11 := int32(m.memory[int64(uint32(v4))+4])
							v6 = t11
							if v6 == i32(255) {
								t16 := int32(load32(m.memory[int64(uint32(v4))+8:]))
								v6 = t16
								v5 = int32(uint32(v6) >> 8)
								v8 = i32(3)
								t17 := int32(load32(m.memory[int64(uint32(v4))+12:]))
								v3 = t17
								goto l22
							}
							t12 := int32(load16(m.memory[int64(uint32(v4))+5:]))
							t13 := int32(m.memory[int64(uint32(v4))+7])
							v5 = t12 | t13<<16
							t14 := int32(load32(m.memory[int64(uint32(v4))+12:]))
							v2 = t14
							t15 := int32(load32(m.memory[int64(uint32(v4))+8:]))
							v3 = t15
							goto l22
						}
					}
					goto l22
				}
			}
			v10 = v8 + i32(1)
			t9 := int32(m.memory[uint32(v7+v8)])
			v6 = t9
			v9 = v6 + i32(-9)
			if uint32(v9) <= uint32(i32(23)) {
				if i32_shl(i32(1), v9)&i32(8388627) == 0 {
					goto l20
				}
				v6 = v7 + v10
				if v6 == v11 {
					goto l24
				}
				v12 = v8 + i32(2)
				v8 = v10 + i32(2)
				v9 = v3
			l26:
				{
					t18 := int32(m.memory[uint32(v6)])
					v13 = t18
					v7 = v13 + i32(-9)
					if uint32(v7) > uint32(i32(23)) {
						goto l25
					}
					if i32_shl(i32(1), v7)&i32(8388627) == 0 {
						goto l25
					}
					v6 = v6 + i32(1)
					v8 = v8 + i32(1)
					t19 := v12
					v9 = v9 + i32(-1)
					if t19 != v9 {
						goto l26
					}
				}
			l24:
				v6 = i32(0)
				store32(m.memory[uint32(v1):], uint32(i32(0)))
				v8 = i32(-1)
				{
					{
						t20 := int32(m.memory[int64(uint32(v1))+36])
						if t20 == i32(1) {
							goto l27
						}
						goto l28
					}
				l27:
					m.fn1825(v4+i32(4), v1, v2, v3, v5, v10)
					{
						t21 := int32(m.memory[int64(uint32(v4))+4])
						v6 = t21
						if v6 == i32(255) {
							goto l29
						}
						t22 := int32(load16(m.memory[int64(uint32(v4))+5:]))
						t23 := int32(m.memory[int64(uint32(v4))+7])
						v5 = t22 | t23<<16
						t24 := int32(load32(m.memory[int64(uint32(v4))+12:]))
						v2 = t24
						t25 := int32(load32(m.memory[int64(uint32(v4))+8:]))
						v3 = t25
						goto l28
					}
				l29:
					t26 := int32(load32(m.memory[int64(uint32(v4))+8:]))
					v6 = t26
					v5 = int32(uint32(v6) >> 8)
					v8 = i32(3)
					t27 := int32(load32(m.memory[int64(uint32(v4))+12:]))
					v3 = t27
				}
			l28:
				store16(m.memory[int64(uint32(v0))+5:], uint16(v5))
				store32(m.memory[int64(uint32(v0))+12:], uint32(v2))
				store32(m.memory[int64(uint32(v0))+8:], uint32(v3))
				m.memory[int64(uint32(v0))+4] = byte(v6)
				store32(m.memory[uint32(v0):], uint32(v8))
				m.memory[uint32(v0+i32(7))] = byte(int32(uint32(v5) >> 16))
				goto l13
			l25:
				if v13 != i32(61) {
					t28 := v1
					v6 = v8 + i32(-1)
					store32(m.memory[int64(uint32(t28))+4:], uint32(v6))
					store32(m.memory[uint32(v1):], uint32(i32(1)))
					v8 = i32(-1)
					{
						{
							t29 := int32(m.memory[int64(uint32(v1))+36])
							if t29 != 0 {
								goto l32
							}
							v2 = i32(0)
							goto l33
						}
					l32:
						m.fn1825(v4+i32(4), v1, v2, v3, v5, v10)
						{
							t30 := int32(m.memory[int64(uint32(v4))+4])
							v2 = t30
							if v2 == i32(255) {
								goto l34
							}
							t31 := int32(load16(m.memory[int64(uint32(v4))+5:]))
							t32 := int32(m.memory[int64(uint32(v4))+7])
							v5 = t31 | t32<<16
							t33 := int32(load32(m.memory[int64(uint32(v4))+12:]))
							v3 = t33
							t34 := int32(load32(m.memory[int64(uint32(v4))+8:]))
							v6 = t34
							goto l33
						}
					l34:
						t35 := int32(load32(m.memory[int64(uint32(v4))+8:]))
						v2 = t35
						v5 = int32(uint32(v2) >> 8)
						v8 = i32(3)
						t36 := int32(load32(m.memory[int64(uint32(v4))+12:]))
						v6 = t36
					}
				l33:
					store16(m.memory[int64(uint32(v0))+5:], uint16(v5))
					store32(m.memory[int64(uint32(v0))+12:], uint32(v3))
					store32(m.memory[int64(uint32(v0))+8:], uint32(v6))
					m.memory[int64(uint32(v0))+4] = byte(v2)
					store32(m.memory[uint32(v0):], uint32(v8))
					m.memory[uint32(v0+i32(7))] = byte(int32(uint32(v5) >> 16))
					goto l13
				}
				v6 = v6 + i32(1)
				v9 = v8 + i32(-1)
				goto l31
			}
			goto l20
		}
	l22:
		store16(m.memory[int64(uint32(v0))+5:], uint16(v5))
		store32(m.memory[int64(uint32(v0))+12:], uint32(v2))
		store32(m.memory[int64(uint32(v0))+8:], uint32(v3))
		m.memory[int64(uint32(v0))+4] = byte(v6)
		store32(m.memory[uint32(v0):], uint32(v8))
		m.memory[uint32(v0+i32(7))] = byte(int32(uint32(v5) >> 16))
		goto l13
	l20:
		v8 = v10
		if v6 != i32(61) {
			goto l35
		}
		v8 = v10 + i32(1)
		v6 = v8 + v2
		v9 = v10
	l31:
		m.fn1825(v4+i32(4), v1, v2, v3, v5, v10)
		{
			t37 := int32(m.memory[int64(uint32(v4))+4])
			if t37 == i32(255) {
				{
					if v6 == v11 {
						goto l37
					}
					t40 := int32(load32(m.memory[int64(uint32(v4))+12:]))
					v12 = t40
					t41 := int32(load32(m.memory[int64(uint32(v4))+8:]))
					v13 = t41
					v5 = i32(0)
				l43:
					{
						v10 = v8 + v5
						v9 = v6 + v5
						t42 := int32(m.memory[uint32(v9)])
						v7 = t42
						v2 = v7 + i32(-9)
						if uint32(v2) > uint32(i32(30)) {
							goto l38
						}
						{
							if i32_shl(i32(1), v2)&i32(8388627) != 0 {
								goto l39
							}
							if i32_shl(i32(1), v2)&i32(0x42000000) == 0 {
								goto l38
							}
							v2 = v10 + i32(1)
							v5 = i32(0)
						l41:
							{
								v6 = v9 + v5 + i32(1)
								if v6 == v11 {
									store32(m.memory[int64(uint32(v0))+8:], uint32(v3))
									m.memory[int64(uint32(v0))+5] = byte(v7)
									m.memory[int64(uint32(v0))+4] = byte(i32(3))
									store32(m.memory[uint32(v0):], uint32(i32(-1)))
									store32(m.memory[uint32(v1):], uint32(i32(0)))
									goto l13
								}
								v5 = v5 + i32(1)
								t43 := int32(m.memory[uint32(v6)])
								if t43 != v7 {
									goto l41
								}
							}
							t44 := v1
							v5 = v10 + v5
							store32(m.memory[int64(uint32(t44))+4:], uint32(v5+i32(1)))
							store32(m.memory[uint32(v1):], uint32(i32(1)))
							if v7 != i32(34) {
								store32(m.memory[int64(uint32(v0))+16:], uint32(v5))
								store32(m.memory[int64(uint32(v0))+12:], uint32(v2))
								store32(m.memory[int64(uint32(v0))+8:], uint32(v12))
								store32(m.memory[int64(uint32(v0))+4:], uint32(v13))
								store32(m.memory[uint32(v0):], uint32(i32(1)))
								goto l13
							}
							store32(m.memory[int64(uint32(v0))+16:], uint32(v5))
							store32(m.memory[int64(uint32(v0))+12:], uint32(v2))
							store32(m.memory[int64(uint32(v0))+8:], uint32(v12))
							store32(m.memory[int64(uint32(v0))+4:], uint32(v13))
							store32(m.memory[uint32(v0):], uint32(i32(0)))
							goto l13
						}
					l39:
						t45 := v6
						v5 = v5 + i32(1)
						if t45+v5 != v11 {
							goto l43
						}
					}
				}
			l37:
				store32(m.memory[int64(uint32(v0))+8:], uint32(v3))
				m.memory[int64(uint32(v0))+4] = byte(i32(1))
				store32(m.memory[uint32(v0):], uint32(i32(-1)))
				store32(m.memory[uint32(v1):], uint32(i32(0)))
				goto l13
			l38:
				t46 := int32(m.memory[int64(uint32(v1))+36])
				if t46 != i32(1) {
					store32(m.memory[int64(uint32(v1))+4:], uint32(v10))
					store32(m.memory[uint32(v1):], uint32(i32(2)))
					store32(m.memory[int64(uint32(v0))+8:], uint32(v10))
					m.memory[int64(uint32(v0))+4] = byte(i32(2))
					store32(m.memory[uint32(v0):], uint32(i32(-1)))
					goto l13
				}
				v5 = i32(0)
			l46:
				{
					v6 = v9 + v5 + i32(1)
					if v6 == v11 {
						goto l45
					}
					v5 = v5 + i32(1)
					t47 := int32(m.memory[uint32(v6)])
					v6 = t47 + i32(-9)
					if uint32(v6) > uint32(i32(23)) {
						goto l46
					}
					if i32_shl(i32(1), v6)&i32(8388627) == 0 {
						goto l46
					}
				}
				v3 = v10 + v5
			l45:
				store32(m.memory[int64(uint32(v1))+4:], uint32(v3))
				store32(m.memory[uint32(v1):], uint32(i32(1)))
				store32(m.memory[int64(uint32(v0))+16:], uint32(v3))
				store32(m.memory[int64(uint32(v0))+12:], uint32(v10))
				store32(m.memory[int64(uint32(v0))+8:], uint32(v12))
				store32(m.memory[int64(uint32(v0))+4:], uint32(v13))
				store32(m.memory[uint32(v0):], uint32(i32(2)))
				goto l13
			}
			t38 := int32(load32(m.memory[int64(uint32(v4))+12:]))
			store32(m.memory[int64(uint32(v0))+12:], uint32(t38))
			t39 := int64(load64(m.memory[int64(uint32(v4))+4:]))
			store64(m.memory[int64(uint32(v0))+4:], uint64(t39))
			store32(m.memory[int64(uint32(v1))+4:], uint32(v9))
			store32(m.memory[uint32(v1):], uint32(i32(3)))
			store32(m.memory[uint32(v0):], uint32(i32(-1)))
			goto l13
		}
	l17:
		t48 := v3
		v5 = v5 + i32(1)
		if t48 != v5 {
			goto l47
		}
	}
l15:
	store32(m.memory[uint32(v0):], uint32(i32(-2)))
	store32(m.memory[uint32(v1):], uint32(i32(0)))
l13:
	m.g0 = v4 + i32(16)
}
func (m *Module) fn242(v0, v1 int32) {
	var v2, v3, v4, v5, v6, v7, v8, v9, v10 int32
	t0 := m.g0
	v2 = t0 - i32(48)
	m.g0 = v2
	m.fn247(v2+i32(8), v1+i32(4))
	{
		t1 := int32(load32(m.memory[int64(uint32(v2))+8:]))
		v3 = t1
		if v3 == 0 {
			store32(m.memory[uint32(v0):], uint32(i32(-1)))
			goto l16
		}
		t2 := int32(load32(m.memory[uint32(v1):]))
		v1 = t2
		t3 := int32(load32(m.memory[int64(uint32(v2))+12:]))
		store32(m.memory[int64(uint32(v2))+20:], uint32(t3))
		store32(m.memory[int64(uint32(v2))+16:], uint32(v3))
		store64(m.memory[int64(uint32(v2))+24:], uint64(int64(uint32(i32(1)))<<32|int64(uint32(v2+i32(16)))))
		m.fn17(v2+i32(32), i32(1052665), v2+i32(24))
		t4 := int32(load32(m.memory[int64(uint32(v2))+32:]))
		v3 = t4
		t5 := int32(load32(m.memory[int64(uint32(v2))+36:]))
		t6 := v2 + i32(32)
		v4 = t5
		t7 := int32(load32(m.memory[int64(uint32(v2))+40:]))
		m.fn238(t6, v4, t7)
		t8 := int32(load32(m.memory[int64(uint32(v2))+40:]))
		t9 := int32(load32(m.memory[int64(uint32(v2))+44:]))
		m.fn239(v2+i32(32), v1+i32(72), t8, t9, i32(1))
		t10 := int32(load32(m.memory[int64(uint32(v2))+36:]))
		v5 = t10
		{
			{
				t11 := int32(load32(m.memory[int64(uint32(v2))+32:]))
				v6 = t11
				if v6 != i32(-0x7fffffff) {
					t16 := int32(load32(m.memory[int64(uint32(v2))+20:]))
					v1 = t16
					if v1 <= i32(-1) {
						goto l4
					}
					if v1 != 0 {
						t17 := int32(load32(m.memory[int64(uint32(v2))+16:]))
						v7 = t17
						t18 := m.fn11(v1)
						v8 = t18
						if v8 != 0 {
							if v1 == 0 {
								goto l6
							}
							memory_copy(m.memory, uint32(v8), uint32(v7), uint32(v1))
							goto l6
						}
						m.fn16(i32(1), v1)
						panic("unreachable")
					}
					v8 = i32(1)
					goto l6
				}
				t12 := int32(load32(m.memory[int64(uint32(v2))+40:]))
				m.fn29(v2+i32(32), v5, t12)
				t13 := int32(load32(m.memory[int64(uint32(v2))+40:]))
				v1 = t13
				t14 := int32(load32(m.memory[int64(uint32(v2))+36:]))
				v5 = t14
				t15 := int32(load32(m.memory[int64(uint32(v2))+32:]))
				v7 = t15
				if v7 == i32(-1) {
					goto l2
				}
				v6 = v5
				goto l3
			}
		l2:
			if v1 <= i32(-1) {
				goto l4
			}
			if v1 != 0 {
				goto l8
			}
			v6 = i32(1)
			v1 = i32(0)
			v7 = i32(0)
			goto l3
		l8:
			t19 := m.fn11(v1)
			v6 = t19
			if v6 == 0 {
				m.fn16(i32(1), v1)
				panic("unreachable")
			}
			if v1 == 0 {
				goto l10
			}
			memory_copy(m.memory, uint32(v6), uint32(v5), uint32(v1))
		l10:
			v7 = v1
		}
	l3:
		m.fn198(v2+i32(32), v6, v1)
		{
			t20 := int32(load32(m.memory[int64(uint32(v2))+32:]))
			v5 = t20
			if v5 != i32(-1) {
				t21 := int32(load32(m.memory[int64(uint32(v2))+40:]))
				v1 = t21
				t22 := int32(load32(m.memory[int64(uint32(v2))+36:]))
				v8 = t22
				if v7 == 0 {
					goto l12
				}
				t23 := int32(load32(m.memory[uint32(v6+i32(-4)):]))
				v9 = t23
				v10 = v9 & i32(-8)
				t24 := v10
				v9 = v9 & i32(3)
				p25 := i32(8)
				if v9 != 0 {
					p25 = i32(4)
				}
				if uint32(t24) < uint32(p25+v7) {
					m.fn7(i32(1274404), i32(46), i32(1274452))
					panic("unreachable")
				}
				if v9 == 0 {
					goto l14
				}
				if uint32(v10) > uint32(v7+i32(39)) {
					m.fn7(i32(1274468), i32(46), i32(1274516))
					panic("unreachable")
				}
			l14:
				m.fn5(v6)
				goto l12
			}
			v8 = v6
			v5 = v7
			goto l12
		}
	}
l4:
	m.fn15()
	panic("unreachable")
l6:
	{
		if v6|i32(-0x80000000) == i32(-0x80000000) {
			goto l17
		}
		t26 := int32(load32(m.memory[uint32(v5+i32(-4)):]))
		v7 = t26
		v9 = v7 & i32(-8)
		t27 := v9
		v7 = v7 & i32(3)
		p28 := i32(8)
		if v7 != 0 {
			p28 = i32(4)
		}
		if uint32(t27) < uint32(p28+v6) {
			m.fn7(i32(1274404), i32(46), i32(1274452))
			panic("unreachable")
		}
		if v7 == 0 {
			goto l19
		}
		if uint32(v9) > uint32(v6+i32(39)) {
			m.fn7(i32(1274468), i32(46), i32(1274516))
			panic("unreachable")
		}
	l19:
		m.fn5(v5)
	}
l17:
	v5 = v1
l12:
	{
		if v3 == 0 {
			goto l21
		}
		t29 := int32(load32(m.memory[uint32(v4+i32(-4)):]))
		v6 = t29
		v7 = v6 & i32(-8)
		t30 := v7
		v6 = v6 & i32(3)
		p31 := i32(8)
		if v6 != 0 {
			p31 = i32(4)
		}
		if uint32(t30) < uint32(p31+v3) {
			m.fn7(i32(1274404), i32(46), i32(1274452))
			panic("unreachable")
		}
		if v6 == 0 {
			goto l23
		}
		if uint32(v7) > uint32(v3+i32(39)) {
			m.fn7(i32(1274468), i32(46), i32(1274516))
			panic("unreachable")
		}
	l23:
		m.fn5(v4)
	}
l21:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v1))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v8))
	store32(m.memory[uint32(v0):], uint32(v5))
l16:
	m.g0 = v2 + i32(48)
}
func (m *Module) fn243(v0, v1, v2, v3 int32) {
	var v4, v5, v6, v7 int32
	var v8, v9, v10 int64
	var v11 int32
	t0 := m.g0
	v4 = t0 - i32(80)
	m.g0 = v4
	{
		{
			if v1 == i32(1140336) {
				t19 := m.fn889(v2, v3)
				if t19 == v3 {
					store32(m.memory[int64(uint32(v0))+8:], uint32(v3))
					store32(m.memory[int64(uint32(v0))+4:], uint32(v2))
					store32(m.memory[uint32(v0):], uint32(i32(-1)))
					goto l29
				}
				store32(m.memory[uint32(v0):], uint32(i32(-2)))
				goto l29
			}
			{
				if v1 == i32(1144468) {
					goto l1
				}
				if v1 == i32(1144956) {
					goto l1
				}
				if v1 != i32(1144440) {
					if v1 == i32(1144776) {
						v6 = i32(0)
					l22:
						{
							t7 := v3
							v5 = v6
							if t7 == v5 {
								goto l21
							}
							t8 := int32(int8(m.memory[uint32(v2+v5)]))
							v7 = t8
							if v7 < i32(0) {
								goto l20
							}
							v6 = v5 + i32(1)
							v7 = v7 & i32(255)
							if uint32(v7) > uint32(i32(27)) {
								goto l22
							}
							if i32_shl(i32(1), v7)&i32(0x800c000) == 0 {
								goto l22
							}
							goto l20
						}
					}
					v5 = i32(0)
					v6 = (i32(0) - v2) & i32(3)
					if uint32(v6|i32(8)) > uint32(v3) {
						goto l18
					}
					if v6 == 0 {
						goto l19
					}
					v5 = i32(0)
					t4 := int32(int8(m.memory[uint32(v2)]))
					if t4 < i32(0) {
						goto l20
					}
					v5 = i32(1)
					if v6 == i32(1) {
						goto l19
					}
					t5 := int32(int8(m.memory[int64(uint32(v2))+1]))
					if t5 < i32(0) {
						goto l20
					}
					v5 = i32(2)
					if v6 == i32(2) {
						goto l19
					}
					t6 := int32(int8(m.memory[int64(uint32(v2))+2]))
					if t6 >= i32(0) {
						goto l19
					}
					goto l20
				}
			l1:
				v5 = i32(0)
				v6 = i32(9)
				{
					t1 := int32(m.memory[uint32(v1)])
					switch t1 {
					case 12:
						goto l14
					default:
						t2 := int32(load32(m.memory[int64(uint32(v1))+8:]))
						v5 = t2
						v6 = i32(0)
						goto l14
					case 1:
						v6 = i32(1)
						goto l15
					case 2, 3:
						v6 = i32(2)
						goto l15
					case 4:
						v6 = i32(3)
						goto l15
					case 5:
						v6 = i32(4)
						goto l15
					case 6:
						v6 = i32(5)
						goto l15
					case 7:
						v6 = i32(6)
						goto l15
					case 8:
						v6 = i32(7)
						goto l15
					case 9:
						v6 = i32(8)
						goto l15
					case 10:
						v6 = i32(10)
						v5 = i32(65536)
						goto l14
					case 11:
						v6 = i32(10)
					}
				}
			l15:
				v5 = i32(0)
			l14:
				m.memory[int64(uint32(v4))+40] = byte(i32(9))
				store16(m.memory[int64(uint32(v4))+32:], uint16(i32(49024)))
				store64(m.memory[int64(uint32(v4))+24:], uint64(i64(0)))
				store32(m.memory[int64(uint32(v4))+20:], uint32(v5))
				m.memory[int64(uint32(v4))+19] = byte(i32(0))
				store16(m.memory[int64(uint32(v4))+17:], uint16(i32(0)))
				m.memory[int64(uint32(v4))+16] = byte(v6)
				store32(m.memory[int64(uint32(v4))+36:], uint32(v1))
				m.fn890(v4+i32(8), v4+i32(16), v3)
				t3 := int32(load32(m.memory[int64(uint32(v4))+8:]))
				if t3&i32(1) != 0 {
					t9 := int32(load32(m.memory[int64(uint32(v4))+12:]))
					v1 = t9
					if v1 <= i32(-1) {
						goto l23
					}
					{
						if v1 != 0 {
							goto l24
						}
						v6 = i32(1)
						goto l25
					l24:
						t10 := m.fn11(v1)
						v6 = t10
						if v6 == 0 {
							m.fn16(i32(1), v1)
							panic("unreachable")
						}
					}
				l25:
					t11 := int32(load32(m.memory[int64(uint32(v4))+40:]))
					t12 := v4
					v5 = t11
					store32(m.memory[int64(uint32(t12))+72:], uint32(v5))
					t13 := int64(load64(m.memory[int64(uint32(v4))+32:]))
					t14 := v4
					v8 = t13
					store64(m.memory[int64(uint32(t14))+64:], uint64(v8))
					t15 := int64(load64(m.memory[int64(uint32(v4))+24:]))
					t16 := v4
					v9 = t15
					store64(m.memory[int64(uint32(t16))+56:], uint64(v9))
					t17 := int64(load64(m.memory[int64(uint32(v4))+16:]))
					t18 := v4
					v10 = t17
					store64(m.memory[int64(uint32(t18))+48:], uint64(v10))
					store32(m.memory[int64(uint32(v4))+40:], uint32(v5))
					store64(m.memory[int64(uint32(v4))+32:], uint64(v8))
					store64(m.memory[int64(uint32(v4))+24:], uint64(v9))
					store64(m.memory[int64(uint32(v4))+16:], uint64(v10))
					v5 = i32(0)
					goto l27
				}
				m.fn219(i32(1146304))
				panic("unreachable")
			}
		l19:
			v7 = v3 + i32(-8)
			v5 = v6
		l31:
			{
				v6 = v2 + v5
				t20 := int32(load32(m.memory[uint32(v6+i32(4)):]))
				v11 = t20 & i32(-2139062144)
				t21 := int32(load32(m.memory[uint32(v6):]))
				t22 := v11
				v6 = t21 & i32(-2139062144)
				if t22|v6 != 0 {
					goto l30
				}
				v5 = v5 + i32(8)
				if uint32(v5) <= uint32(v7) {
					goto l31
				}
			}
		l18:
			if uint32(v5) >= uint32(v3) {
				goto l21
			}
		l32:
			{
				t23 := int32(int8(m.memory[uint32(v2+v5)]))
				if t23 < i32(0) {
					goto l20
				}
				t24 := v3
				v5 = v5 + i32(1)
				if t24 != v5 {
					goto l32
				}
				goto l21
			}
		l30:
			if v6 == 0 {
				goto l33
			}
			v5 = int32(uint32(int32(bits.TrailingZeros32(uint32(v6))))>>3) + v5
			goto l20
		l33:
			v5 = int32(uint32(int32(bits.TrailingZeros32(uint32(v11))))>>3) + i32(4) + v5
		l20:
			if v3 == v5 {
				goto l21
			}
			v6 = i32(0)
			v7 = i32(9)
			{
				t25 := int32(m.memory[uint32(v1)])
				switch t25 {
				case 12:
					goto l45
				default:
					t26 := int32(load32(m.memory[int64(uint32(v1))+8:]))
					v6 = t26
					v7 = i32(0)
					goto l45
				case 1:
					v7 = i32(1)
					goto l46
				case 2, 3:
					v7 = i32(2)
					goto l46
				case 4:
					v7 = i32(3)
					goto l46
				case 5:
					v7 = i32(4)
					goto l46
				case 6:
					v7 = i32(5)
					goto l46
				case 7:
					v7 = i32(6)
					goto l46
				case 8:
					v7 = i32(7)
					goto l46
				case 9:
					v7 = i32(8)
					goto l46
				case 10:
					v7 = i32(10)
					v6 = i32(65536)
					goto l45
				case 11:
					v7 = i32(10)
				}
			}
		l46:
			v6 = i32(0)
		l45:
			m.memory[int64(uint32(v4))+40] = byte(i32(9))
			store16(m.memory[int64(uint32(v4))+32:], uint16(i32(49024)))
			store64(m.memory[int64(uint32(v4))+24:], uint64(i64(0)))
			store32(m.memory[int64(uint32(v4))+20:], uint32(v6))
			m.memory[int64(uint32(v4))+19] = byte(i32(0))
			store16(m.memory[int64(uint32(v4))+17:], uint16(i32(0)))
			m.memory[int64(uint32(v4))+16] = byte(v7)
			store32(m.memory[int64(uint32(v4))+36:], uint32(v1))
			t27 := v4
			t28 := v4 + i32(16)
			v7 = v3 - v5
			m.fn890(t27, t28, v7)
			t29 := int32(load32(m.memory[uint32(v4):]))
			if t29&i32(1) == 0 {
				goto l47
			}
			t30 := int32(load32(m.memory[int64(uint32(v4))+4:]))
			v6 = t30
			v1 = v6 + v5
			if uint32(v1) < uint32(v6) {
				goto l47
			}
			if v1 <= i32(-1) {
				goto l23
			}
			{
				if v1 != 0 {
					goto l48
				}
				v6 = i32(1)
				goto l49
			l48:
				t31 := m.fn11(v1)
				v6 = t31
				if v6 == 0 {
					m.fn16(i32(1), v1)
					panic("unreachable")
				}
			}
		l49:
			if v5 == 0 {
				goto l51
			}
			memory_copy(m.memory, uint32(v6), uint32(v2), uint32(v5))
		l51:
			if uint32(v3) < uint32(v5) {
				m.fn121(v5, v3, v3, i32(1146272))
				panic("unreachable")
			}
			t32 := int32(load32(m.memory[int64(uint32(v4))+40:]))
			t33 := v4
			v3 = t32
			store32(m.memory[int64(uint32(t33))+72:], uint32(v3))
			t34 := int64(load64(m.memory[int64(uint32(v4))+32:]))
			t35 := v4
			v8 = t34
			store64(m.memory[int64(uint32(t35))+64:], uint64(v8))
			t36 := int64(load64(m.memory[int64(uint32(v4))+24:]))
			t37 := v4
			v9 = t36
			store64(m.memory[int64(uint32(t37))+56:], uint64(v9))
			t38 := int64(load64(m.memory[int64(uint32(v4))+16:]))
			t39 := v4
			v10 = t38
			store64(m.memory[int64(uint32(t39))+48:], uint64(v10))
			store32(m.memory[int64(uint32(v4))+40:], uint32(v3))
			store64(m.memory[int64(uint32(v4))+32:], uint64(v8))
			store64(m.memory[int64(uint32(v4))+24:], uint64(v9))
			store64(m.memory[int64(uint32(v4))+16:], uint64(v10))
			if uint32(v5) > uint32(v1) {
				m.fn121(v5, v1, v1, i32(1146588))
				panic("unreachable")
			}
			v2 = v2 + v5
			v3 = v7
		}
	l27:
		m.fn892(v4+i32(48), v4+i32(16), v2, v3, v6+v5, v1-v5)
		t40 := int32(m.memory[int64(uint32(v4))+52])
		switch t40 {
		case 1:
			m.fn7(i32(1274576), i32(40), i32(1146320))
			panic("unreachable")
		case 2:
			store32(m.memory[uint32(v0):], uint32(i32(-2)))
			if v1 == 0 {
				goto l29
			}
			t42 := int32(load32(m.memory[uint32(v6+i32(-4)):]))
			v5 = t42
			v2 = v5 & i32(-8)
			t43 := v2
			v5 = v5 & i32(3)
			p44 := i32(8)
			if v5 != 0 {
				p44 = i32(4)
			}
			if uint32(t43) < uint32(p44+v1) {
				m.fn7(i32(1274404), i32(46), i32(1274452))
				panic("unreachable")
			}
			if v5 == 0 {
				goto l58
			}
			if uint32(v2) > uint32(v1+i32(39)) {
				m.fn7(i32(1274468), i32(46), i32(1274516))
				panic("unreachable")
			}
		l58:
			m.fn5(v6)
			goto l29
		default:
			t41 := int32(load32(m.memory[int64(uint32(v4))+56:]))
			v2 = t41
			store32(m.memory[int64(uint32(v0))+4:], uint32(v6))
			store32(m.memory[uint32(v0):], uint32(v1))
			store32(m.memory[int64(uint32(v0))+8:], uint32(v2+v5))
			goto l29
		}
	}
l47:
	m.fn219(i32(1146288))
	panic("unreachable")
l23:
	m.fn15()
	panic("unreachable")
l21:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v3))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v2))
	store32(m.memory[uint32(v0):], uint32(i32(-1)))
l29:
	m.g0 = v4 + i32(80)
}
func (m *Module) fn244(v0, v1, v2, v3, v4, v5, v6 int32) {
	var v7, v8, v9, v10, v11, v12 int32
	t0 := m.g0
	v7 = t0 - i32(16)
	m.g0 = v7
	{
		if uint32(v6) < uint32(v5) {
			goto l0
		}
		{
			if v5 == 0 {
				goto l1
			}
			if uint32(v5) < uint32(v4) {
				goto l2
			}
			if v5 != v4 {
				goto l0
			}
			goto l1
		l2:
			t1 := int32(int8(m.memory[uint32(v3+v5)]))
			if t1 <= i32(-65) {
				goto l0
			}
		}
	l1:
		{
			if v6 == 0 {
				goto l3
			}
			if uint32(v6) < uint32(v4) {
				goto l4
			}
			if v6 == v4 {
				goto l3
			}
			goto l0
		l4:
			t2 := int32(int8(m.memory[uint32(v3+v6)]))
			if t2 <= i32(-65) {
				goto l0
			}
		}
	l3:
		{
			{
				v8 = v6 - v5
				t3 := int32(load32(m.memory[uint32(v1):]))
				t4 := int32(load32(m.memory[int64(uint32(v1))+8:]))
				t5 := v8
				v9 = t4
				if uint32(t5) <= uint32(t3-v9) {
					goto l5
				}
				m.fn245(v1, v9, v8)
				t6 := int32(load32(m.memory[int64(uint32(v1))+8:]))
				v9 = t6
				goto l6
			}
		l5:
			if v6 == v5 {
				goto l7
			}
		l6:
			if v8 == 0 {
				goto l7
			}
			t7 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			memory_copy(m.memory, uint32(t7+v9), uint32(v3+v5), uint32(v8))
		}
	l7:
		t8 := v1
		v5 = v9 + v8
		store32(m.memory[int64(uint32(t8))+8:], uint32(v5))
		if uint32(v6) < uint32(v4) {
			goto l8
		}
		m.fn33(v6, v4, i32(1272188))
		panic("unreachable")
	}
l0:
	m.fn38(v3, v4, v5, v6, i32(1272172))
	panic("unreachable")
l8:
	{
		t9 := int32(m.memory[uint32(v3+v6)])
		v8 = t9
		if v8 == i32(9) {
			{
				t27 := int32(load32(m.memory[uint32(v1):]))
				if t27 != v5 {
					goto l22
				}
				m.fn245(v1, v5, i32(1))
			}
		l22:
			store32(m.memory[uint32(v0):], uint32(i32(-1)))
			store32(m.memory[int64(uint32(v1))+8:], uint32(v5+i32(1)))
			store32(m.memory[int64(uint32(v0))+4:], uint32(v6+i32(1)))
			t28 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			m.memory[uint32(t28+v5)] = byte(i32(32))
			goto l18
		}
		if v8 == i32(38) {
			{
				t22 := int32(load32(m.memory[int64(uint32(v2))+4:]))
				v9 = t22
				t23 := int32(load32(m.memory[uint32(v2):]))
				t24 := v9
				v8 = t23
				if t24 == v8 {
					goto l19
				}
				v10 = v6 + i32(1)
				v12 = v8 + (v9 - v8)
				v5 = i32(0)
			l21:
				{
					v11 = v8 + v5
					t25 := int32(m.memory[uint32(v11)])
					if t25 == i32(59) {
						goto l20
					}
					t26 := v8
					v5 = v5 + i32(1)
					if t26+v5 != v9 {
						goto l21
					}
				}
				store32(m.memory[uint32(v2):], uint32(v12))
			}
		l19:
			store32(m.memory[int64(uint32(v0))+8:], uint32(v4))
			store32(m.memory[int64(uint32(v0))+4:], uint32(v6))
			store32(m.memory[uint32(v0):], uint32(i32(-0x80000000)))
			goto l18
		}
		{
			t10 := m.fn1818(v1, v3, v4, v6, i32(32))
			v10 = t10
			v4 = v10 + (v6 ^ i32(-1))
			if v4 == 0 {
				goto l11
			}
			v1 = v4 & i32(3)
			t11 := int32(load32(m.memory[int64(uint32(v2))+4:]))
			v8 = t11
			t12 := int32(load32(m.memory[uint32(v2):]))
			v5 = t12
			if uint32(v10-v6+i32(-2)) < uint32(i32(3)) {
				goto l17
			}
			v4 = v4 & i32(-4)
		l15:
			{
				t13 := v5
				var p14 int32
				if v5 != v8 {
					p14 = 1
				}
				v9 = p14
				v5 = t13 + v9
				t15 := v5
				var p16 int32
				if v5 != v8 {
					p16 = 1
				}
				v11 = p16
				v5 = t15 + v11
				t17 := v5
				var p18 int32
				if v5 != v8 {
					p18 = 1
				}
				v3 = p18
				v6 = t17 + v3
				t19 := v6
				var p20 int32
				if v6 != v8 {
					p20 = 1
				}
				v5 = t19 + p20
				if v9 != 0 {
					goto l13
				}
				if v11 != 0 {
					goto l13
				}
				if v3 != 0 {
					goto l13
				}
				if v6 == v8 {
					goto l14
				}
			l13:
				store32(m.memory[uint32(v2):], uint32(v5))
			l14:
				v4 = v4 + i32(-4)
				if v4 != 0 {
					goto l15
				}
			}
			if v1 == 0 {
				goto l11
			}
		l17:
			{
				if v5 == v8 {
					goto l16
				}
				t21 := v2
				v5 = v5 + i32(1)
				store32(m.memory[uint32(t21):], uint32(v5))
			}
		l16:
			v1 = v1 + i32(-1)
			if v1 != 0 {
				goto l17
			}
		}
	l11:
		store32(m.memory[uint32(v0):], uint32(i32(-1)))
		store32(m.memory[int64(uint32(v0))+4:], uint32(v10))
		goto l18
	}
l20:
	store32(m.memory[uint32(v2):], uint32(v11+i32(1)))
	{
		t29 := v6
		v8 = v5 + v10
		if uint32(t29) >= uint32(v8) {
			goto l23
		}
		{
			if uint32(v10) >= uint32(v4) {
				goto l24
			}
			t30 := int32(int8(m.memory[uint32(v3+v10)]))
			if t30 < i32(-64) {
				goto l23
			}
		}
	l24:
		if uint32(v8) < uint32(v4) {
			goto l25
		}
		if v8 != v4 {
			goto l23
		}
		goto l26
	l25:
		t31 := int32(int8(m.memory[uint32(v3+v8)]))
		if t31 > i32(-65) {
			goto l26
		}
	}
l23:
	m.fn38(v3, v4, v10, v8, i32(1272204))
	panic("unreachable")
l26:
	if v5 != 0 {
		goto l27
	}
	v1 = i32(1)
	goto l28
l27:
	{
		{
			{
				{
					v6 = v3 + v10
					t32 := int32(m.memory[uint32(v6)])
					v2 = t32
					if v2 == i32(35) {
						m.fn625(v7+i32(8), v6+i32(1), v5+i32(-1))
						{
							t40 := int32(m.memory[int64(uint32(v7))+8])
							if t40 == i32(255) {
								t42 := int32(load32(m.memory[int64(uint32(v7))+12:]))
								v5 = t42
								store32(m.memory[int64(uint32(v7))+8:], uint32(i32(0)))
								m.fn481(v7, v5, v7+i32(8))
								t43 := int32(load32(m.memory[uint32(v7):]))
								t44 := int32(load32(m.memory[int64(uint32(v7))+4:]))
								m.fn1820(v1, t43, t44)
								goto l43
							}
							t41 := int64(load64(m.memory[int64(uint32(v7))+8:]))
							store64(m.memory[int64(uint32(v0))+4:], uint64(t41))
							store32(m.memory[uint32(v0):], uint32(i32(-0x7fffffff)))
							goto l18
						}
					}
					switch v5 + i32(-2) {
					case 0:
						v4 = i32(1272904)
						v9 = i32(116)
						v11 = i32(1)
						switch v2 + i32(-103) {
						case 5:
							goto l36
						default:
							goto l35
						case 0:
							v4 = i32(1272905)
							goto l36
						}
					case 2:
						switch v2 + i32(-97) {
						case 0:
							t39 := int32(m.memory[int64(uint32(v6))+1])
							if t39 == i32(112) {
								goto l42
							}
							goto l35
						default:
							goto l35
						case 16:
							t33 := int32(m.memory[int64(uint32(v6))+1])
							if t33 != i32(117) {
								goto l35
							}
							t34 := int32(m.memory[int64(uint32(v6))+2])
							if t34 != i32(111) {
								goto l35
							}
							v4 = i32(1272893)
							v9 = i32(116)
							goto l39
						}
					case 1:
						t35 := int32(load16(m.memory[uint32(v6):]))
						t36 := int32(m.memory[uint32(v6+i32(2))])
						if (t35^i32(28001)|(t36^i32(112)))&i32(0xffff) == 0 {
							m.fn1819(v1)
							goto l43
						}
						if v2 != i32(97) {
							goto l35
						}
						t37 := int32(m.memory[int64(uint32(v6))+1])
						if t37 != i32(109) {
							goto l35
						}
						v4 = i32(1272906)
						t38 := int32(m.memory[int64(uint32(v6))+2])
						if t38 != i32(112) {
							goto l35
						}
						goto l41
					default:
						if v5 > i32(-1) {
							goto l35
						}
						m.fn15()
						panic("unreachable")
					}
				}
			l42:
				t45 := int32(m.memory[int64(uint32(v6))+2])
				if t45 != i32(111) {
					goto l35
				}
				v4 = i32(1272907)
				v9 = i32(115)
			}
		l39:
			v11 = i32(3)
		l36:
			t46 := int32(m.memory[uint32(v6+v11)])
			if t46 != v9 {
				goto l35
			}
		}
	l41:
		{
			t47 := int32(m.memory[uint32(v4)])
			v5 = t47 + i32(-9)
			if uint32(v5) > uint32(i32(29)) {
				goto l45
			}
			if i32_shl(i32(1), v5)&i32(0x20000013) != 0 {
				store32(m.memory[uint32(v0):], uint32(i32(-0x7ffffffe)))
				goto l18
			}
		}
	l45:
		m.fn1820(v1, v4, i32(1))
		goto l43
	l35:
		t48 := m.fn11(v5)
		v1 = t48
		if v1 != 0 {
			goto l47
		}
		m.fn16(i32(1), v5)
		panic("unreachable")
	}
l47:
	if v5 == 0 {
		goto l28
	}
	memory_copy(m.memory, uint32(v1), uint32(v6), uint32(v5))
l28:
	store32(m.memory[int64(uint32(v0))+16:], uint32(v8))
	store32(m.memory[int64(uint32(v0))+12:], uint32(v10))
	store32(m.memory[int64(uint32(v0))+8:], uint32(v5))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
	store32(m.memory[uint32(v0):], uint32(v5))
	goto l18
l43:
	store32(m.memory[uint32(v0):], uint32(i32(-1)))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v8+i32(1)))
l18:
	m.g0 = v7 + i32(16)
}
func (m *Module) fn245(v0, v1, v2 int32) {
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
	m.fn1817(t2, t4, t3, v2, i32(1), i32(1))
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
func (m *Module) fn246(v0 int32) {
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
	m.fn208(t2, t4, t3, v2, i32(4), i32(32))
	{
		t6 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		if t6 != i32(1) {
			goto l0
		}
		t7 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		t8 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		m.fn16(t7, t8)
		panic("unreachable")
	}
l0:
	t9 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	v3 = t9
	store32(m.memory[uint32(v0):], uint32(v2))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
	m.g0 = v1 + i32(16)
}
func (m *Module) fn247(v0, v1 int32) {
	var v2, v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13 int32
	t0 := int32(m.memory[int64(uint32(v1))+29])
	v2 = t0
	t1 := int32(load32(m.memory[int64(uint32(v1))+24:]))
	v3 = t1
	t2 := int32(load32(m.memory[int64(uint32(v1))+16:]))
	v4 = t2
	t3 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	v5 = t3
	t4 := int32(m.memory[int64(uint32(v1))+28])
	v6 = t4
	t5 := int32(load32(m.memory[int64(uint32(v1))+20:]))
	v7 = t5
	t6 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	v8 = t6
	t7 := int32(load32(m.memory[uint32(v1):]))
	v9 = t7
l17:
	v10 = v9
	v11 = i32(0)
	if v2&i32(1) == 0 {
		goto l0
	}
	goto l1
l0:
	if v4 == v7 {
		goto l2
	}
l15:
	v12 = v3
	{
		{
			v3 = v4
			t8 := int32(int8(m.memory[uint32(v3)]))
			v2 = t8
			if v2 <= i32(-1) {
				goto l3
			}
			v4 = v3 + i32(1)
			v2 = v2 & i32(255)
			goto l4
		}
	l3:
		t9 := int32(m.memory[int64(uint32(v3))+1])
		v4 = t9 & i32(63)
		v9 = v2 & i32(31)
		if uint32(v2) > uint32(i32(-33)) {
			goto l5
		}
		v2 = v9<<6 | v4
		v4 = v3 + i32(2)
		goto l4
	l5:
		t10 := int32(m.memory[int64(uint32(v3))+2])
		v4 = v4<<6 | t10&i32(63)
		if uint32(v2) >= uint32(i32(-16)) {
			goto l6
		}
		v2 = v4 | v9<<12
		v4 = v3 + i32(3)
		goto l4
	l6:
		t11 := int32(m.memory[int64(uint32(v3))+3])
		v2 = v4<<6 | t11&i32(63) | v9<<18&i32(0x1c0000)
		v4 = v3 + i32(4)
	}
l4:
	v3 = v4 - v3 + v12
	v9 = v2 + i32(-9)
	if uint32(v9) > uint32(i32(23)) {
		goto l7
	}
	if i32_shl(i32(1), v9)&i32(8388639) != 0 {
		goto l8
	}
l7:
	if uint32(v2) < uint32(i32(133)) {
		goto l9
	}
	v9 = int32(uint32(v2) >> 8)
	switch v9 + i32(-22) {
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25:
		goto l9
	case 26:
		if v2 != i32(12288) {
			goto l9
		}
		goto l8
	case 10:
		t12 := int32(m.memory[int64(uint32(v2&i32(255)))+1139700])
		if t12&i32(2) == 0 {
			goto l9
		}
		goto l8
	default:
		if v9 != 0 {
			goto l9
		}
		t13 := int32(m.memory[int64(uint32(v2&i32(255)))+1139700])
		if t13&i32(1) != 0 {
			goto l8
		}
		goto l9
	case 0:
		if v2 != i32(5760) {
			goto l9
		}
	}
l8:
	store32(m.memory[int64(uint32(v1))+24:], uint32(v3))
	store32(m.memory[int64(uint32(v1))+16:], uint32(v4))
	store32(m.memory[uint32(v1):], uint32(v3))
	v2 = i32(0)
	v9 = v3
	goto l14
l9:
	if v4 != v7 {
		goto l15
	}
	store32(m.memory[int64(uint32(v1))+24:], uint32(v3))
	store32(m.memory[int64(uint32(v1))+16:], uint32(v4))
l2:
	v2 = i32(1)
	m.memory[int64(uint32(v1))+29] = byte(i32(1))
	if v6&i32(1) == 0 {
		goto l16
	}
	v9 = v10
	v12 = v5
	goto l14
l16:
	v9 = v10
	v12 = v5
	if v5 == v10 {
		goto l1
	}
l14:
	v13 = v12 - v10
	if v13 == 0 {
		goto l17
	}
	v11 = v8 + v10
l1:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v13))
	store32(m.memory[uint32(v0):], uint32(v11))
}
func (m *Module) fn248(v0 int32) {
	var v1, v2, v3, v4 int32
	{
		{
			t0 := int32(load32(m.memory[int64(uint32(v0))+44:]))
			v1 = t0
			if v1 == 0 {
				goto l0
			}
			t1 := int32(load32(m.memory[int64(uint32(v0))+40:]))
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
			t5 := int32(load32(m.memory[int64(uint32(v0))+52:]))
			v1 = t5
			if v1 == 0 {
				goto l4
			}
			t6 := int32(load32(m.memory[int64(uint32(v0))+48:]))
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
			t10 := int32(load32(m.memory[int64(uint32(v0))+80:]))
			v1 = t10
			if v1 == 0 {
				goto l8
			}
			t11 := int32(load32(m.memory[uint32(v1):]))
			t12 := v1
			v3 = t11
			store32(m.memory[uint32(t12):], uint32(v3+i32(-1)))
			if v3 != i32(1) {
				goto l8
			}
			t13 := int32(load32(m.memory[int64(uint32(v0))+80:]))
			t14 := int32(load32(m.memory[int64(uint32(v0))+84:]))
			m.fn249(t13, t14)
		}
	l8:
		{
			t15 := int32(load32(m.memory[int64(uint32(v0))+88:]))
			v1 = t15
			if v1 == 0 {
				goto l9
			}
			t16 := int32(load32(m.memory[uint32(v1):]))
			t17 := v1
			v3 = t16
			store32(m.memory[uint32(t17):], uint32(v3+i32(-1)))
			if v3 != i32(1) {
				goto l9
			}
			t18 := int32(load32(m.memory[int64(uint32(v0))+88:]))
			t19 := int32(load32(m.memory[int64(uint32(v0))+92:]))
			m.fn249(t18, t19)
		}
	l9:
		{
			t20 := int32(load32(m.memory[int64(uint32(v0))+60:]))
			v1 = t20
			if v1 == 0 {
				goto l10
			}
			t21 := int32(load32(m.memory[int64(uint32(v0))+56:]))
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
				goto l12
			}
			if uint32(v4) > uint32(v1+i32(39)) {
				m.fn7(i32(1274468), i32(46), i32(1274516))
				panic("unreachable")
			}
		l12:
			m.fn5(v2)
		}
	l10:
		t25 := int32(m.memory[int64(uint32(v0))+120])
		if t25 == i32(2) {
			m.fn28(i32(1276344), i32(121), i32(1276404))
			panic("unreachable")
		}
		{
			t26 := int32(load32(m.memory[int64(uint32(v0))+136:]))
			v1 = t26
			if v1 == 0 {
				return
			}
			t27 := int32(load32(m.memory[int64(uint32(v0))+140:]))
			v3 = t27
			t28 := int32(load32(m.memory[uint32(v3+i32(-4)):]))
			v0 = t28
			v2 = v0 & i32(-8)
			t29 := v2
			v0 = v0 & i32(3)
			p30 := i32(8)
			if v0 != 0 {
				p30 = i32(4)
			}
			v1 = v1 << 5
			if uint32(t29) < uint32(p30|v1) {
				m.fn7(i32(1274404), i32(46), i32(1274452))
				panic("unreachable")
			}
			if v0 == 0 {
				goto l17
			}
			if uint32(v2) > uint32(v1+i32(39)) {
				m.fn7(i32(1274468), i32(46), i32(1274516))
				panic("unreachable")
			}
		l17:
			m.fn5(v3)
		}
		return
	}
}
func (m *Module) fn249(v0, v1 int32) {
	var v2, v3 int32
	{
		if v0 == i32(-1) {
			return
		}
		t0 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		t1 := v0
		v2 = t0
		store32(m.memory[int64(uint32(t1))+4:], uint32(v2+i32(-1)))
		if v2 != i32(1) {
			return
		}
		v1 = (v1 + i32(11)) & i32(-4)
		if v1 == 0 {
			return
		}
		t2 := int32(load32(m.memory[uint32(v0+i32(-4)):]))
		v2 = t2
		v3 = v2 & i32(-8)
		t3 := v3
		v2 = v2 & i32(3)
		p4 := i32(8)
		if v2 != 0 {
			p4 = i32(4)
		}
		if uint32(t3) < uint32(p4+v1) {
			m.fn7(i32(1274404), i32(46), i32(1274452))
			panic("unreachable")
		}
		if v2 == 0 {
			goto l2
		}
		if uint32(v3) > uint32(v1+i32(39)) {
			m.fn7(i32(1274468), i32(46), i32(1274516))
			panic("unreachable")
		}
	l2:
		m.fn5(v0)
	}
}
func (m *Module) fn250(v0, v1 int64, v2, v3 int32) int64 {
	var v4 int32
	var v5, v6, v7, v8 int64
	t0 := m.g0
	v4 = t0 - i32(80)
	m.g0 = v4
	store64(m.memory[int64(uint32(v4))+56:], uint64(i64(0)))
	store64(m.memory[int64(uint32(v4))+64:], uint64(i64(0)))
	store64(m.memory[int64(uint32(v4))+48:], uint64(v1))
	store64(m.memory[int64(uint32(v4))+32:], uint64(v1^i64(8387220255154660723)))
	store64(m.memory[int64(uint32(v4))+24:], uint64(v1^i64(7237128888997146477)))
	store64(m.memory[int64(uint32(v4))+40:], uint64(v0))
	store64(m.memory[int64(uint32(v4))+16:], uint64(v0^i64(0x6c7967656e657261)))
	store64(m.memory[int64(uint32(v4))+8:], uint64(v0^i64(8317987319222330741)))
	m.fn59(v4+i32(8), v2, v3)
	m.memory[int64(uint32(v4))+79] = byte(i32(255))
	m.fn59(v4+i32(8), v4+i32(79), i32(1))
	t1 := int64(load64(m.memory[int64(uint32(v4))+8:]))
	v0 = t1
	t2 := int64(load64(m.memory[int64(uint32(v4))+24:]))
	v1 = t2
	t3 := int64(load32(m.memory[int64(uint32(v4))+64:]))
	v5 = t3
	t4 := int64(load64(m.memory[int64(uint32(v4))+56:]))
	v6 = t4
	t5 := int64(load64(m.memory[int64(uint32(v4))+32:]))
	v7 = t5
	t6 := int64(load64(m.memory[int64(uint32(v4))+16:]))
	v8 = t6
	m.g0 = v4 + i32(80)
	t7 := v7
	v5 = v6 | v5<<56
	v6 = t7 ^ v5
	t8 := i64_rotl(v6, i64(16))
	v6 = v6 + v8
	v7 = t8 ^ v6
	t9 := i64_rotl(v7, i64(21))
	t10 := v7
	v0 = v1 + v0
	v7 = t10 + i64_rotl(v0, i64(32))
	v8 = t9 ^ v7
	t11 := i64_rotl(v8, i64(16))
	t12 := v8
	t13 := v6
	v1 = i64_rotl(v1, i64(13)) ^ v0
	v0 = t13 + v1
	v6 = t12 + (i64_rotl(v0, i64(32)) ^ i64(255))
	v8 = t11 ^ v6
	t14 := i64_rotl(v8, i64(21))
	t15 := v8
	t16 := v7 ^ v5
	v1 = v0 ^ i64_rotl(v1, i64(17))
	v0 = t16 + v1
	v5 = t15 + i64_rotl(v0, i64(32))
	v7 = t14 ^ v5
	t17 := i64_rotl(v7, i64(16))
	t18 := v7
	v1 = v0 ^ i64_rotl(v1, i64(13))
	v0 = v1 + v6
	v6 = t18 + i64_rotl(v0, i64(32))
	v7 = t17 ^ v6
	t19 := i64_rotl(v7, i64(21))
	t20 := v7
	v1 = v0 ^ i64_rotl(v1, i64(17))
	v0 = v1 + v5
	v5 = t20 + i64_rotl(v0, i64(32))
	v7 = t19 ^ v5
	t21 := i64_rotl(v7, i64(16))
	t22 := v7
	v1 = i64_rotl(v1, i64(13)) ^ v0
	v0 = v1 + v6
	v6 = t22 + i64_rotl(v0, i64(32))
	t23 := i64_rotl(t21^v6, i64(21))
	v1 = i64_rotl(v1, i64(17)) ^ v0
	v1 = i64_rotl(v1, i64(13)) ^ (v1 + v5)
	t24 := t23 ^ i64_rotl(v1, i64(17))
	v1 = v1 + v6
	return t24 ^ i64_rotl(v1, i64(32)) ^ v1
}
func (m *Module) fn251(v0, v1, v2, v3 int32) {
	var v4, v5 int32
	var v6, v7, v8 int64
	var v9, v10, v11, v12, v13 int32
	t0 := m.g0
	v4 = t0 - i32(256)
	m.g0 = v4
	t1 := int32(load32(m.memory[int64(uint32(v1))+16:]))
	t2 := v4 + i32(8)
	v5 = t1
	m.fn154(t2, v5+i32(24), v2, v3)
	{
		t3 := int32(load32(m.memory[int64(uint32(v4))+8:]))
		if t3 != i32(1) {
			store64(m.memory[uint32(v0):], uint64(i64(-1)))
			store32(m.memory[int64(uint32(v0))+8:], uint32(i32(-0x7ffffffd)))
			goto l2
		}
		{
			t4 := int32(load32(m.memory[int64(uint32(v4))+12:]))
			v3 = t4
			t5 := int32(load32(m.memory[int64(uint32(v5))+48:]))
			if uint32(v3) < uint32(t5) {
				t6 := int32(load32(m.memory[int64(uint32(v5))+44:]))
				v3 = t6 + v3*i32(192)
				t7 := int32(m.memory[int64(uint32(v3))+168])
				if t7 != 0 {
					store32(m.memory[int64(uint32(v0))+16:], uint32(i32(33)))
					store32(m.memory[int64(uint32(v0))+12:], uint32(i32(1073610)))
					store32(m.memory[int64(uint32(v0))+8:], uint32(i32(-0x7ffffffe)))
					store64(m.memory[uint32(v0):], uint64(i64(-1)))
					goto l2
				}
				{
					{
						t8 := int32(m.memory[int64(uint32(v3))+120])
						if t8 != i32(3) {
							goto l4
						}
						t9 := int64(load64(m.memory[int64(uint32(v3))+112:]))
						v6 = t9
						goto l5
					}
				l4:
					t10 := int32(load32(m.memory[int64(uint32(v1))+4:]))
					v2 = t10
					v7 = int64(uint32(v2))
					{
						{
							{
								{
									t11 := int64(load64(m.memory[int64(uint32(v3))+96:]))
									t12 := v2
									v6 = t11
									p13 := i64(0xffffffff)
									if uint64(v6) < uint64(i64(0xffffffff)) {
										p13 = v6
									}
									v5 = t12 - int32(p13)
									p14 := v5
									if uint32(v5) > uint32(v2) {
										p14 = i32(0)
									}
									if uint32(p14) > uint32(i32(29)) {
										goto l6
									}
									t15 := int64(load64(m.memory[int64(uint32(i32(0)))+1277248:]))
									v8 = t15
									v3 = int32(v8)
									if v8&i64(255) != i64(255) {
										goto l7
									}
									store64(m.memory[int64(uint32(v1))+8:], uint64(v6+i64(30)))
									if v3&i32(255) != i32(255) {
										goto l8
									}
									goto l9
								l7:
									store64(m.memory[int64(uint32(v1))+8:], uint64(v7))
									if v3&i32(255) == i32(255) {
										goto l9
									}
								l8:
									v1 = int32(int64(uint64(v8) >> 32))
									v6 = int64(uint64(v8) >> 8)
									v2 = int32(v6)
									switch v3 & i32(255) {
									case 2, 3:
										t16 := int32(m.memory[int64(uint32(v1))+8])
										v2 = t16
										fallthrough
									case 1:
										if v2&i32(255) == i32(37) {
											store32(m.memory[int64(uint32(v4))+20:], uint32(i32(29)))
											store32(m.memory[int64(uint32(v4))+16:], uint32(i32(1070356)))
											store64(m.memory[int64(uint32(v4))+24:], uint64(int64(uint32(i32(1)))<<32|int64(uint32(v4+i32(16)))))
											m.fn17(v4+i32(128)+i32(4), i32(1050747), v4+i32(24))
											if v3&i32(255) != i32(3) {
												goto l14
											}
											t17 := int32(load32(m.memory[uint32(v1):]))
											v2 = t17
											{
												t18 := int32(load32(m.memory[uint32(v1+i32(4)):]))
												v3 = t18
												t19 := int32(load32(m.memory[uint32(v3):]))
												v5 = t19
												if v5 == 0 {
													goto l15
												}
												m.t0[uint(v5)].(func(int32))(v2)
											}
										l15:
											{
												t20 := int32(load32(m.memory[int64(uint32(v3))+4:]))
												v5 = t20
												if v5 == 0 {
													goto l16
												}
												t21 := int32(load32(m.memory[int64(uint32(v3))+8:]))
												m.fn21(v2, v5, t21)
											}
										l16:
											m.fn21(v1, i32(12), i32(4))
											goto l14
										}
										fallthrough
									default:
										store32(m.memory[int64(uint32(v4))+140:], uint32(v1))
										store32(m.memory[int64(uint32(v4))+132:], uint32(i32(-0x80000000)))
										store32(m.memory[int64(uint32(v4))+136:], uint32(int32(v6)<<8&i32(0xff00)|v3&i32(-65281)))
										goto l14
									}
								}
							l6:
								t22 := v1
								v8 = v6 + i64(30)
								store64(m.memory[int64(uint32(t22))+8:], uint64(v8))
								t23 := int32(load32(m.memory[uint32(v1):]))
								p24 := v7
								if uint64(v6) < uint64(v7) {
									p24 = v6
								}
								v2 = t23 + int32(p24)
								t25 := int32(load32(m.memory[uint32(v2):]))
								if t25 == i32(67324752) {
									goto l17
								}
							}
						l9:
							t26 := int32(load32(m.memory[int64(uint32(i32(0)))+0x105554:]))
							store32(m.memory[int64(uint32(v4))+140:], uint32(t26))
							t27 := int64(load64(m.memory[int64(uint32(i32(0)))+1070412:]))
							store64(m.memory[int64(uint32(v4))+132:], uint64(t27))
						}
					l14:
						t28 := int32(load32(m.memory[int64(uint32(v4))+133:]))
						store32(m.memory[int64(uint32(v4))+89:], uint32(t28))
						t29 := int32(m.memory[int64(uint32(v4))+132])
						m.memory[int64(uint32(v4))+88] = byte(t29)
						t30 := int32(m.memory[int64(uint32(v4))+143])
						m.memory[int64(uint32(v4))+99] = byte(t30)
						t31 := int32(load32(m.memory[int64(uint32(v4))+139:]))
						store32(m.memory[int64(uint32(v4))+95:], uint32(t31))
						t32 := int32(load16(m.memory[int64(uint32(v4))+137:]))
						store16(m.memory[int64(uint32(v4))+93:], uint16(t32))
						t33 := int32(load32(m.memory[int64(uint32(v4))+88:]))
						v1 = t33
						t34 := int64(load64(m.memory[int64(uint32(v4))+92:]))
						store64(m.memory[int64(uint32(v0))+12:], uint64(t34))
						store32(m.memory[int64(uint32(v0))+8:], uint32(v1))
						store64(m.memory[uint32(v0):], uint64(i64(-1)))
						goto l2
					}
				l17:
					t35 := int32(load16(m.memory[int64(uint32(v2))+28:]))
					store16(m.memory[int64(uint32(v4))+153:], uint16(t35))
					t36 := int64(load64(m.memory[int64(uint32(v2))+20:]))
					store64(m.memory[int64(uint32(v4))+145:], uint64(t36))
					t37 := int64(load16(m.memory[int64(uint32(v4))+153:]))
					v6 = t37
					t38 := int64(load16(m.memory[int64(uint32(v4))+151:]))
					v7 = t38
					store64(m.memory[int64(uint32(v4))+128:], uint64(i64(1)))
					t39 := v4
					v6 = v6 + (v8 + v7)
					store64(m.memory[int64(uint32(t39))+136:], uint64(v6))
					m.fn256(v3+i32(112), v4+i32(128))
				}
			l5:
				store64(m.memory[int64(uint32(v1))+8:], uint64(v6))
				t40 := int64(load64(m.memory[int64(uint32(v3))+64:]))
				t41 := v4
				v6 = t40
				store64(m.memory[int64(uint32(t41))+40:], uint64(v6))
				store64(m.memory[int64(uint32(v4))+32:], uint64(v6))
				{
					{
						t42 := int32(load16(m.memory[int64(uint32(v3))+148:]))
						v2 = t42
						if v2 != i32(2) {
							goto l18
						}
						t43 := int32(load16(m.memory[int64(uint32(v3))+150:]))
						v1 = t43
						v3 = i32(-0x7ffffffb)
						goto l19
					}
				l18:
					t44 := int32(load16(m.memory[int64(uint32(v3))+32:]))
					if t44 == 0 {
						t45 := int64(load64(m.memory[int64(uint32(v4))+36:]))
						v7 = t45
						v5 = int32(int64(uint64(v6) >> 32))
						v9 = int32(v6)
						t46 := int32(load32(m.memory[int64(uint32(v3))+152:]))
						v10 = t46
						{
							{
								if v2 == 0 {
									t48 := m.fn11(i32(72))
									v2 = t48
									if v2 == 0 {
										m.fn23(i32(8), i32(72))
										panic("unreachable")
									}
									m.memory[int64(uint32(v2))+68] = byte(i32(1))
									store32(m.memory[int64(uint32(v2))+64:], uint32(v10))
									store32(m.memory[int64(uint32(v2))+56:], uint32(i32(0)))
									store64(m.memory[int64(uint32(v2))+48:], uint64(i64(0)))
									store32(m.memory[int64(uint32(v2))+24:], uint32(v1))
									store32(m.memory[int64(uint32(v2))+20:], uint32(v5))
									store64(m.memory[int64(uint32(v2))+12:], uint64(v7))
									store32(m.memory[int64(uint32(v2))+8:], uint32(v9))
									store64(m.memory[uint32(v2):], uint64(i64(0)))
									v1 = i32(2)
									goto l24
								}
								t47 := m.fn11(i32(8192))
								v11 = t47
								if v11 != 0 {
									goto l22
								}
								m.fn16(i32(1), i32(8192))
								panic("unreachable")
							}
						l22:
							store32(m.memory[int64(uint32(v4))+144:], uint32(v1))
							store32(m.memory[int64(uint32(v4))+140:], uint32(v5))
							store64(m.memory[int64(uint32(v4))+132:], uint64(v7))
							store32(m.memory[int64(uint32(v4))+128:], uint32(v9))
							m.fn257(v4 + i32(168))
							t49 := int32(load32(m.memory[int64(uint32(v4))+145:]))
							store32(m.memory[int64(uint32(v4))+24:], uint32(t49))
							t50 := int32(load32(m.memory[int64(uint32(v4))+148:]))
							store32(m.memory[int64(uint32(v4))+27:], uint32(t50))
							t51 := int64(load64(m.memory[int64(uint32(v4))+160:]))
							store64(m.memory[int64(uint32(v4))+88:], uint64(t51))
							t52 := int64(load64(m.memory[int64(uint32(v4))+168:]))
							store64(m.memory[int64(uint32(v4))+96:], uint64(t52))
							t53 := int64(load64(m.memory[int64(uint32(v4))+176:]))
							store64(m.memory[int64(uint32(v4))+104:], uint64(t53))
							t54 := int64(load64(m.memory[int64(uint32(v4))+184:]))
							store64(m.memory[int64(uint32(v4))+112:], uint64(t54))
							t55 := int64(load64(m.memory[int64(uint32(v4))+192:]))
							store64(m.memory[int64(uint32(v4))+120:], uint64(t55))
							t56 := int32(m.memory[int64(uint32(v4))+144])
							v1 = t56
							t57 := int32(load32(m.memory[int64(uint32(v4))+140:]))
							v5 = t57
							t58 := int32(load32(m.memory[int64(uint32(v4))+136:]))
							v9 = t58
							t59 := int32(load32(m.memory[int64(uint32(v4))+132:]))
							v12 = t59
							t60 := int32(load32(m.memory[int64(uint32(v4))+128:]))
							v13 = t60
							t61 := int64(load64(m.memory[int64(uint32(v4))+152:]))
							v6 = t61
							t62 := int64(load64(m.memory[int64(uint32(v4))+248:]))
							store64(m.memory[int64(uint32(v4))+80:], uint64(t62))
							t63 := int64(load64(m.memory[int64(uint32(v4))+240:]))
							store64(m.memory[int64(uint32(v4))+72:], uint64(t63))
							t64 := int64(load64(m.memory[int64(uint32(v4))+232:]))
							store64(m.memory[int64(uint32(v4))+64:], uint64(t64))
							t65 := int64(load64(m.memory[int64(uint32(v4))+224:]))
							store64(m.memory[int64(uint32(v4))+56:], uint64(t65))
							t66 := int64(load64(m.memory[int64(uint32(v4))+216:]))
							store64(m.memory[int64(uint32(v4))+48:], uint64(t66))
							t67 := int64(load64(m.memory[int64(uint32(v4))+208:]))
							store64(m.memory[int64(uint32(v4))+40:], uint64(t67))
							t68 := int64(load64(m.memory[int64(uint32(v4))+200:]))
							store64(m.memory[int64(uint32(v4))+32:], uint64(t68))
							t69 := m.fn11(i32(184))
							v2 = t69
							if v2 == 0 {
								m.fn23(i32(8), i32(184))
								panic("unreachable")
							}
							store64(m.memory[int64(uint32(v2))+6:], uint64(i64(0)))
							store16(m.memory[int64(uint32(v2))+4:], uint16(i32(8192)))
							store32(m.memory[uint32(v2):], uint32(v11))
							store32(m.memory[int64(uint32(v2))+13:], uint32(i32(0)))
							m.memory[int64(uint32(v2))+48] = byte(v1)
							store32(m.memory[int64(uint32(v2))+44:], uint32(v5))
							store32(m.memory[int64(uint32(v2))+40:], uint32(v9))
							store32(m.memory[int64(uint32(v2))+36:], uint32(v12))
							store32(m.memory[int64(uint32(v2))+32:], uint32(v13))
							store64(m.memory[int64(uint32(v2))+24:], uint64(i64(0)))
							t70 := int32(load32(m.memory[int64(uint32(v4))+24:]))
							store32(m.memory[int64(uint32(v2))+49:], uint32(t70))
							t71 := int32(load32(m.memory[int64(uint32(v4))+27:]))
							store32(m.memory[int64(uint32(v2))+52:], uint32(t71))
							store64(m.memory[int64(uint32(v2))+56:], uint64(v6))
							t72 := int64(load64(m.memory[int64(uint32(v4))+88:]))
							store64(m.memory[int64(uint32(v2))+64:], uint64(t72))
							t73 := int64(load64(m.memory[int64(uint32(v4))+96:]))
							store64(m.memory[int64(uint32(v2))+72:], uint64(t73))
							t74 := int64(load64(m.memory[int64(uint32(v4))+104:]))
							store64(m.memory[int64(uint32(v2))+80:], uint64(t74))
							t75 := int64(load64(m.memory[int64(uint32(v4))+112:]))
							store64(m.memory[int64(uint32(v2))+88:], uint64(t75))
							t76 := int64(load64(m.memory[int64(uint32(v4))+120:]))
							store64(m.memory[int64(uint32(v2))+96:], uint64(t76))
							store64(m.memory[int64(uint32(v2))+160:], uint64(i64(0)))
							store32(m.memory[int64(uint32(v2))+168:], uint32(i32(0)))
							store32(m.memory[int64(uint32(v2))+176:], uint32(v10))
							m.memory[int64(uint32(v2))+180] = byte(i32(1))
							t77 := int64(load64(m.memory[int64(uint32(v4))+80:]))
							store64(m.memory[int64(uint32(v2))+152:], uint64(t77))
							t78 := int64(load64(m.memory[int64(uint32(v4))+72:]))
							store64(m.memory[int64(uint32(v2))+144:], uint64(t78))
							t79 := int64(load64(m.memory[int64(uint32(v4))+64:]))
							store64(m.memory[int64(uint32(v2))+136:], uint64(t79))
							t80 := int64(load64(m.memory[int64(uint32(v4))+56:]))
							store64(m.memory[int64(uint32(v2))+128:], uint64(t80))
							t81 := int64(load64(m.memory[int64(uint32(v4))+48:]))
							store64(m.memory[int64(uint32(v2))+120:], uint64(t81))
							t82 := int64(load64(m.memory[int64(uint32(v4))+40:]))
							store64(m.memory[int64(uint32(v2))+112:], uint64(t82))
							t83 := int64(load64(m.memory[int64(uint32(v4))+32:]))
							store64(m.memory[int64(uint32(v2))+104:], uint64(t83))
							v1 = i32(3)
						}
					l24:
						store32(m.memory[int64(uint32(v0))+188:], uint32(i32(0)))
						store16(m.memory[int64(uint32(v0))+186:], uint16(i32(0)))
						store32(m.memory[int64(uint32(v0))+180:], uint32(v2))
						store32(m.memory[int64(uint32(v0))+176:], uint32(v1))
						store32(m.memory[int64(uint32(v0))+8:], uint32(v3))
						store64(m.memory[uint32(v0):], uint64(i64(2)))
						goto l2
					}
					v3 = i32(-0x7ffffffc)
					v1 = i32(0)
				}
			l19:
				store32(m.memory[int64(uint32(v0))+12:], uint32(v1))
				store32(m.memory[int64(uint32(v0))+8:], uint32(v3))
				store64(m.memory[uint32(v0):], uint64(i64(-1)))
				goto l2
			}
			store64(m.memory[uint32(v0):], uint64(i64(-1)))
			store32(m.memory[int64(uint32(v0))+8:], uint32(i32(-0x7ffffffd)))
			goto l2
		}
	}
l2:
	m.g0 = v4 + i32(256)
}
func (m *Module) fn252(v0, v1, v2 int32) {
	var v3 int32
	var v4 int64
	var v5, v6, v7 int32
	var v8 int64
	var v9, v10, v11, v12 int32
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
				t1 := int64(load64(m.memory[int64(uint32(v1))+8:]))
				v4 = t1
				if v4 == 0 {
					goto l0
				}
				t2 := int32(load32(m.memory[int64(uint32(v1))+16:]))
				v5 = t2
			l19:
				{
					t4 := v3 + i32(40)
					t5 := v5
					t6 := v3 + i32(8)
					p3 := i64(32)
					if uint64(v4) < uint64(i64(32)) {
						p3 = v4
					}
					m.fn253(t4, t5, t6, int32(p3))
					{
						{
							t7 := int32(m.memory[int64(uint32(v3))+40])
							if t7 == i32(255) {
								goto l1
							}
							t8 := int32(load32(m.memory[int64(uint32(v3))+44:]))
							v6 = t8
							t9 := int32(load32(m.memory[int64(uint32(v3))+40:]))
							v7 = t9
							goto l2
						}
					l1:
						t10 := int32(load32(m.memory[int64(uint32(v3))+44:]))
						t11 := v4
						v6 = t10
						v8 = int64(uint32(v6))
						if uint64(t11) < uint64(v8) {
							m.fn28(i32(1080592), i32(69), i32(1080628))
							panic("unreachable")
						}
						t12 := v1
						v4 = v4 - v8
						store64(m.memory[int64(uint32(t12))+8:], uint64(v4))
						v7 = v7 | i32(255)
					}
				l2:
					switch v7 & i32(255) {
					case 0:
						goto l4
					case 1:
						if v7&i32(0xff00) != i32(8960) {
							goto l4
						}
						goto l9
					default:
						if uint32(v6) < uint32(i32(33)) {
							goto l10
						}
						m.fn121(i32(0), v6, i32(32), i32(1070132))
						panic("unreachable")
					case 2:
						t13 := int32(m.memory[int64(uint32(v6))+8])
						if t13 == i32(35) {
							goto l9
						}
						goto l4
					case 3:
						t14 := int32(m.memory[int64(uint32(v6))+8])
						if t14 != i32(35) {
							goto l4
						}
						t15 := int32(load32(m.memory[uint32(v6):]))
						v9 = t15
						{
							t16 := int32(load32(m.memory[uint32(v6+i32(4)):]))
							v10 = t16
							t17 := int32(load32(m.memory[uint32(v10):]))
							v11 = t17
							if v11 == 0 {
								goto l11
							}
							m.t0[uint(v11)].(func(int32))(v9)
						}
					l11:
						{
							t18 := int32(load32(m.memory[int64(uint32(v10))+4:]))
							v10 = t18
							if v10 == 0 {
								goto l12
							}
							t19 := int32(load32(m.memory[uint32(v9+i32(-4)):]))
							v11 = t19
							v12 = v11 & i32(-8)
							t20 := v12
							v11 = v11 & i32(3)
							p21 := i32(8)
							if v11 != 0 {
								p21 = i32(4)
							}
							if uint32(t20) < uint32(p21+v10) {
								m.fn7(i32(1274404), i32(46), i32(1274452))
								panic("unreachable")
							}
							if v11 == 0 {
								goto l14
							}
							if uint32(v12) > uint32(v10+i32(39)) {
								m.fn7(i32(1274468), i32(46), i32(1274516))
								panic("unreachable")
							}
						l14:
							m.fn5(v9)
						}
					l12:
						t22 := int32(load32(m.memory[uint32(v6+i32(-4)):]))
						v9 = t22
						v10 = v9 & i32(-8)
						t23 := v10
						v9 = v9 & i32(3)
						p24 := i32(20)
						if v9 != 0 {
							p24 = i32(16)
						}
						if uint32(t23) < uint32(p24) {
							m.fn7(i32(1274404), i32(46), i32(1274452))
							panic("unreachable")
						}
						if v9 == 0 {
							goto l17
						}
						if uint32(v10) >= uint32(i32(52)) {
							m.fn7(i32(1274468), i32(46), i32(1274516))
							panic("unreachable")
						}
					l17:
						m.fn5(v6)
					}
				l9:
					if !(v4 == 0) {
						goto l19
					}
				}
			}
		l0:
			v5 = v2 + i32(8)
			t25 := int32(load32(m.memory[int64(uint32(v2))+8:]))
			v7 = t25
			goto l20
		}
	l4:
		store32(m.memory[int64(uint32(v0))+4:], uint32(v6))
		store32(m.memory[uint32(v0):], uint32(v7))
		goto l21
	l10:
		v5 = v2 + i32(8)
		{
			t26 := int32(load32(m.memory[uint32(v2):]))
			t27 := int32(load32(m.memory[int64(uint32(v2))+8:]))
			t28 := v6
			v7 = t27
			if uint32(t28) <= uint32(t26-v7) {
				goto l22
			}
			m.fn197(v2, v7, v6, i32(1), i32(1))
			t29 := int32(load32(m.memory[int64(uint32(v2))+8:]))
			v7 = t29
			goto l23
		}
	l22:
		if v6 != 0 {
			goto l23
		}
	l20:
		v6 = i32(0)
		goto l24
	l23:
		if v6 == 0 {
			goto l24
		}
		t30 := int32(load32(m.memory[int64(uint32(v2))+4:]))
		memory_copy(m.memory, uint32(t30+v7), uint32(v3+i32(8)), uint32(v6))
	}
l24:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v6))
	m.memory[uint32(v0)] = byte(i32(255))
	store32(m.memory[uint32(v5):], uint32(v7+v6))
l21:
	m.g0 = v3 + i32(48)
}
func (m *Module) fn253(v0, v1, v2, v3 int32) {
	var v4 int64
	var v5 int32
	var v6 int64
	var v7 int32
	var v8 int64
	var v9, v10 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v1))+176:]))
		switch t0 {
		case 1:
			{
				{
					t1 := int64(load64(m.memory[int64(uint32(v1))+192:]))
					v4 = t1
					if !(v4 == 0) {
						goto l4
					}
					v3 = i32(0)
					goto l5
				}
			l4:
				t2 := int32(load32(m.memory[int64(uint32(v1))+200:]))
				v5 = t2
				t3 := int32(load32(m.memory[uint32(v5):]))
				t4 := int64(load64(m.memory[int64(uint32(v5))+8:]))
				v6 = t4
				t5 := int32(load32(m.memory[int64(uint32(v5))+4:]))
				t6 := v6
				v7 = t5
				v8 = int64(uint32(v7))
				p7 := v8
				if uint64(v6) < uint64(v8) {
					p7 = t6
				}
				v9 = t3 + int32(p7)
				{
					t9 := v7
					p8 := i64(0xffffffff)
					if uint64(v6) < uint64(i64(0xffffffff)) {
						p8 = v6
					}
					v10 = t9 - int32(p8)
					p10 := v10
					if uint32(v10) > uint32(v7) {
						p10 = i32(0)
					}
					v7 = p10
					t11 := v7
					t12 := v4
					v8 = int64(uint32(v3))
					p13 := v8
					if uint64(v4) < uint64(v8) {
						p13 = t12
					}
					v3 = int32(p13)
					p14 := v3
					if uint32(v7) < uint32(v3) {
						p14 = t11
					}
					v3 = p14
					if v3 != i32(1) {
						goto l6
					}
					t15 := int32(m.memory[uint32(v9)])
					m.memory[uint32(v2)] = byte(t15)
					goto l7
				}
			l6:
				if v3 == 0 {
					goto l7
				}
				memory_copy(m.memory, uint32(v2), uint32(v9), uint32(v3))
			l7:
				t16 := v1
				t17 := v4
				v8 = int64(uint32(v3))
				store64(m.memory[int64(uint32(t16))+192:], uint64(t17-v8))
				store64(m.memory[int64(uint32(v5))+8:], uint64(v6+v8))
			}
		l5:
			m.memory[uint32(v0)] = byte(i32(255))
			store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
			return
		case 2:
			t18 := int32(load32(m.memory[int64(uint32(v1))+180:]))
			m.fn258(v0, t18, v2, v3)
			return
		case 3:
			t19 := int32(load32(m.memory[int64(uint32(v1))+180:]))
			m.fn259(v0, t19, v2, v3)
			return
		default:
			t20 := m.fn11(i32(37))
			v1 = t20
			if v1 == 0 {
				m.fn16(i32(1), i32(37))
				panic("unreachable")
			}
			t21 := int64(load64(m.memory[int64(uint32(i32(0)))+1075585:]))
			store64(m.memory[int64(uint32(v1))+29:], uint64(t21))
			t22 := int64(load64(m.memory[int64(uint32(i32(0)))+1075580:]))
			store64(m.memory[int64(uint32(v1))+24:], uint64(t22))
			t23 := int64(load64(m.memory[int64(uint32(i32(0)))+1075572:]))
			store64(m.memory[int64(uint32(v1))+16:], uint64(t23))
			t24 := int64(load64(m.memory[int64(uint32(i32(0)))+1075564:]))
			store64(m.memory[int64(uint32(v1))+8:], uint64(t24))
			t25 := int64(load64(m.memory[int64(uint32(i32(0)))+1075556:]))
			store64(m.memory[uint32(v1):], uint64(t25))
			{
				t26 := m.fn11(i32(12))
				v3 = t26
				if v3 == 0 {
					m.fn23(i32(4), i32(12))
					panic("unreachable")
				}
				store32(m.memory[int64(uint32(v3))+8:], uint32(i32(37)))
				store32(m.memory[int64(uint32(v3))+4:], uint32(v1))
				store32(m.memory[uint32(v3):], uint32(i32(37)))
				t27 := m.fn11(i32(12))
				v1 = t27
				if v1 == 0 {
					m.fn23(i32(4), i32(12))
					panic("unreachable")
				}
				m.memory[int64(uint32(v1))+8] = byte(i32(40))
				store32(m.memory[int64(uint32(v1))+4:], uint32(i32(1070848)))
				store32(m.memory[uint32(v1):], uint32(v3))
				store64(m.memory[uint32(v0):], uint64(int64(uint32(v1))<<32|i64(3)))
				return
			}
		}
	}
}
func (m *Module) fn254(v0 int32) {
	var v1, v2, v3 int32
	var v4 int64
	var v5, v6, v7, v8, v9, v10 int32
	var v11, v12 int64
	t0 := m.g0
	v1 = t0 - i32(160)
	m.g0 = v1
	{
		{
			{
				t1 := int64(load64(m.memory[uint32(v0):]))
				if t1 == i64(2) {
					goto l0
				}
				t2 := int32(load32(m.memory[int64(uint32(v0))+176:]))
				v2 = t2
				store32(m.memory[int64(uint32(v0))+176:], uint32(i32(0)))
				t3 := int32(load32(m.memory[int64(uint32(v0))+180:]))
				v3 = t3
				{
					switch v2 {
					case 1:
						t4 := int32(load32(m.memory[int64(uint32(v0))+200:]))
						v2 = t4
						t5 := int64(load64(m.memory[int64(uint32(v0))+192:]))
						v4 = t5
						t6 := int32(load32(m.memory[int64(uint32(v0))+184:]))
						v5 = t6
						t7 := int32(load32(m.memory[int64(uint32(v0))+188:]))
						v6 = t7
						goto l5
					case 2:
						t8 := int32(load32(m.memory[uint32(v3+i32(-4)):]))
						v7 = t8
						t9 := v7 & i32(-8)
						v8 = v7 & i32(3)
						p10 := i32(80)
						if v8 != 0 {
							p10 = i32(76)
						}
						if uint32(t9) < uint32(p10) {
							m.fn7(i32(1274404), i32(46), i32(1274452))
							panic("unreachable")
						}
						t11 := int32(load32(m.memory[int64(uint32(v3))+24:]))
						v2 = t11
						t12 := int64(load64(m.memory[int64(uint32(v3))+16:]))
						v4 = t12
						t13 := int32(load32(m.memory[int64(uint32(v3))+12:]))
						v6 = t13
						t14 := int32(load32(m.memory[int64(uint32(v3))+8:]))
						v5 = t14
						if v8 == 0 {
							goto l7
						}
						if uint32(v7) < uint32(i32(112)) {
							goto l7
						}
						m.fn7(i32(1274468), i32(46), i32(1274516))
						panic("unreachable")
					default:
						t15 := m.fn11(i32(37))
						v3 = t15
						if v3 == 0 {
							m.fn16(i32(1), i32(37))
							panic("unreachable")
						}
						t16 := int64(load64(m.memory[int64(uint32(i32(0)))+1075585:]))
						store64(m.memory[int64(uint32(v3))+29:], uint64(t16))
						t17 := int64(load64(m.memory[int64(uint32(i32(0)))+1075580:]))
						store64(m.memory[int64(uint32(v3))+24:], uint64(t17))
						t18 := int64(load64(m.memory[int64(uint32(i32(0)))+1075572:]))
						store64(m.memory[int64(uint32(v3))+16:], uint64(t18))
						t19 := int64(load64(m.memory[int64(uint32(i32(0)))+1075564:]))
						store64(m.memory[int64(uint32(v3))+8:], uint64(t19))
						t20 := int64(load64(m.memory[int64(uint32(i32(0)))+1075556:]))
						store64(m.memory[uint32(v3):], uint64(t20))
						t21 := m.fn11(i32(12))
						v2 = t21
						if v2 == 0 {
							m.fn23(i32(4), i32(12))
							panic("unreachable")
						}
						store32(m.memory[int64(uint32(v2))+8:], uint32(i32(37)))
						store32(m.memory[int64(uint32(v2))+4:], uint32(v3))
						store32(m.memory[uint32(v2):], uint32(i32(37)))
						t22 := m.fn11(i32(12))
						v6 = t22
						if v6 == 0 {
							m.fn23(i32(4), i32(12))
							panic("unreachable")
						}
						m.memory[int64(uint32(v6))+8] = byte(i32(40))
						store32(m.memory[int64(uint32(v6))+4:], uint32(i32(1070848)))
						store32(m.memory[uint32(v6):], uint32(v2))
						goto l11
					case 3:
						t23 := int32(load32(m.memory[int64(uint32(v3))+36:]))
						v8 = t23
						t24 := int32(load32(m.memory[int64(uint32(v3))+32:]))
						v9 = t24
						{
							{
								t25 := int64(load64(m.memory[int64(uint32(v3))+24:]))
								v4 = t25
								if v4 == i64(2) {
									goto l12
								}
								t26 := int64(load64(m.memory[int64(uint32(v3))+16:]))
								store64(m.memory[int64(uint32(v1))+16:], uint64(t26))
								t27 := int64(load64(m.memory[int64(uint32(v3))+8:]))
								store64(m.memory[int64(uint32(v1))+8:], uint64(t27))
								t28 := int64(load64(m.memory[uint32(v3):]))
								store64(m.memory[uint32(v1):], uint64(t28))
								t29 := int64(load64(m.memory[int64(uint32(v3))+40:]))
								store64(m.memory[int64(uint32(v1))+40:], uint64(t29))
								t30 := int64(load64(m.memory[int64(uint32(v3))+48:]))
								store64(m.memory[int64(uint32(v1))+48:], uint64(t30))
								t31 := int64(load64(m.memory[int64(uint32(v3))+56:]))
								store64(m.memory[int64(uint32(v1))+56:], uint64(t31))
								t32 := int64(load64(m.memory[int64(uint32(v3))+64:]))
								store64(m.memory[int64(uint32(v1))+64:], uint64(t32))
								t33 := int64(load64(m.memory[int64(uint32(v3))+72:]))
								store64(m.memory[int64(uint32(v1))+72:], uint64(t33))
								t34 := int64(load64(m.memory[int64(uint32(v3))+80:]))
								store64(m.memory[int64(uint32(v1))+80:], uint64(t34))
								t35 := int64(load64(m.memory[int64(uint32(v3))+96:]))
								store64(m.memory[int64(uint32(v1))+96:], uint64(t35))
								t36 := int64(load64(m.memory[int64(uint32(v3))+88:]))
								store64(m.memory[int64(uint32(v1))+88:], uint64(t36))
								t37 := int64(load64(m.memory[int64(uint32(v3))+152:]))
								store64(m.memory[int64(uint32(v1))+152:], uint64(t37))
								t38 := int64(load64(m.memory[int64(uint32(v3))+144:]))
								store64(m.memory[int64(uint32(v1))+144:], uint64(t38))
								t39 := int64(load64(m.memory[int64(uint32(v3))+136:]))
								store64(m.memory[int64(uint32(v1))+136:], uint64(t39))
								t40 := int64(load64(m.memory[int64(uint32(v3))+128:]))
								store64(m.memory[int64(uint32(v1))+128:], uint64(t40))
								t41 := int64(load64(m.memory[int64(uint32(v3))+120:]))
								store64(m.memory[int64(uint32(v1))+120:], uint64(t41))
								t42 := int64(load64(m.memory[int64(uint32(v3))+112:]))
								store64(m.memory[int64(uint32(v1))+112:], uint64(t42))
								t43 := int64(load64(m.memory[int64(uint32(v3))+104:]))
								store64(m.memory[int64(uint32(v1))+104:], uint64(t43))
								store32(m.memory[int64(uint32(v1))+36:], uint32(v8))
								store32(m.memory[int64(uint32(v1))+32:], uint32(v9))
								store64(m.memory[int64(uint32(v1))+24:], uint64(v4))
								t44 := int32(load32(m.memory[int64(uint32(v1))+48:]))
								v2 = t44
								t45 := int64(load64(m.memory[int64(uint32(v1))+40:]))
								v4 = t45
								t46 := int32(load32(m.memory[int64(uint32(v1))+4:]))
								v7 = t46
								t47 := int32(load32(m.memory[uint32(v1):]))
								v10 = t47
								m.fn260(v1 + i32(88))
								v6 = v8
								v5 = v9
								goto l13
							}
						l12:
							t48 := int32(load32(m.memory[int64(uint32(v3))+80:]))
							v2 = t48
							t49 := int64(load64(m.memory[int64(uint32(v3))+72:]))
							v4 = t49
							t50 := int32(load32(m.memory[int64(uint32(v3))+68:]))
							v6 = t50
							t51 := int32(load32(m.memory[int64(uint32(v3))+64:]))
							v5 = t51
							v10 = v9
							v7 = v8
						}
					l13:
						{
							if v7 == 0 {
								goto l14
							}
							t52 := int32(load32(m.memory[uint32(v10+i32(-4)):]))
							v8 = t52
							v9 = v8 & i32(-8)
							t53 := v9
							v8 = v8 & i32(3)
							p54 := i32(8)
							if v8 != 0 {
								p54 = i32(4)
							}
							if uint32(t53) < uint32(p54+v7) {
								m.fn7(i32(1274404), i32(46), i32(1274452))
								panic("unreachable")
							}
							if v8 == 0 {
								goto l16
							}
							if uint32(v9) > uint32(v7+i32(39)) {
								m.fn7(i32(1274468), i32(46), i32(1274516))
								panic("unreachable")
							}
						l16:
							m.fn5(v10)
						}
					l14:
						t55 := int32(load32(m.memory[uint32(v3+i32(-4)):]))
						v7 = t55
						t56 := v7 & i32(-8)
						v8 = v7 & i32(3)
						p57 := i32(192)
						if v8 != 0 {
							p57 = i32(188)
						}
						if uint32(t56) < uint32(p57) {
							m.fn7(i32(1274404), i32(46), i32(1274452))
							panic("unreachable")
						}
						if v8 == 0 {
							goto l7
						}
						if uint32(v7) >= uint32(i32(224)) {
							m.fn7(i32(1274468), i32(46), i32(1274516))
							panic("unreachable")
						}
					}
				l7:
					m.fn5(v3)
				l5:
					if v2 == 0 {
						goto l20
					}
				l22:
					{
						if v4 == 0 {
							goto l21
						}
						t58 := int64(load64(m.memory[int64(uint32(v2))+8:]))
						t59 := v2
						v11 = t58
						t60 := int32(load32(m.memory[int64(uint32(v2))+4:]))
						t61 := v11
						v3 = t60
						t63 := v3
						p62 := i64(0xffffffff)
						if uint64(v11) < uint64(i64(0xffffffff)) {
							p62 = v11
						}
						v6 = t63 - int32(p62)
						p64 := v6
						if uint32(v6) > uint32(v3) {
							p64 = i32(0)
						}
						v6 = p64
						t66 := v6
						p65 := i64(8192)
						if uint64(v4) < uint64(i64(8192)) {
							p65 = v4
						}
						v5 = int32(p65)
						p67 := v5
						if uint32(v6) < uint32(v5) {
							p67 = t66
						}
						v12 = int64(uint32(p67))
						store64(m.memory[int64(uint32(t59))+8:], uint64(t61+v12))
						v4 = v4 - v12
						t68 := v3
						t69 := v11
						v12 = int64(uint32(v3))
						p70 := v12
						if uint64(v11) < uint64(v12) {
							p70 = t69
						}
						if t68 != int32(p70) {
							goto l22
						}
						goto l21
					}
				l20:
					if v5&i32(255) != i32(3) {
						goto l21
					}
				l11:
					t71 := int32(load32(m.memory[uint32(v6):]))
					v3 = t71
					{
						t72 := int32(load32(m.memory[uint32(v6+i32(4)):]))
						v2 = t72
						t73 := int32(load32(m.memory[uint32(v2):]))
						v5 = t73
						if v5 == 0 {
							goto l23
						}
						m.t0[uint(v5)].(func(int32))(v3)
					}
				l23:
					{
						t74 := int32(load32(m.memory[int64(uint32(v2))+4:]))
						v2 = t74
						if v2 == 0 {
							goto l24
						}
						t75 := int32(load32(m.memory[uint32(v3+i32(-4)):]))
						v5 = t75
						v7 = v5 & i32(-8)
						t76 := v7
						v5 = v5 & i32(3)
						p77 := i32(8)
						if v5 != 0 {
							p77 = i32(4)
						}
						if uint32(t76) < uint32(p77+v2) {
							m.fn7(i32(1274404), i32(46), i32(1274452))
							panic("unreachable")
						}
						if v5 == 0 {
							goto l26
						}
						if uint32(v7) > uint32(v2+i32(39)) {
							m.fn7(i32(1274468), i32(46), i32(1274516))
							panic("unreachable")
						}
					l26:
						m.fn5(v3)
					}
				l24:
					t78 := int32(load32(m.memory[uint32(v6+i32(-4)):]))
					v3 = t78
					v2 = v3 & i32(-8)
					t79 := v2
					v3 = v3 & i32(3)
					p80 := i32(20)
					if v3 != 0 {
						p80 = i32(16)
					}
					if uint32(t79) < uint32(p80) {
						m.fn7(i32(1274404), i32(46), i32(1274452))
						panic("unreachable")
					}
					if v3 == 0 {
						goto l29
					}
					if uint32(v2) >= uint32(i32(52)) {
						m.fn7(i32(1274468), i32(46), i32(1274516))
						panic("unreachable")
					}
				l29:
					m.fn5(v6)
				}
			l21:
				m.fn248(v0)
			}
		l0:
			t81 := int32(load32(m.memory[int64(uint32(v0))+180:]))
			v3 = t81
			{
				t82 := int32(load32(m.memory[int64(uint32(v0))+176:]))
				switch t82 {
				case 2:
					goto l32
				default:
					goto l31
				case 3:
					t83 := int64(load64(m.memory[int64(uint32(v3))+24:]))
					if t83 != i64(2) {
						{
							t89 := int32(load32(m.memory[int64(uint32(v3))+4:]))
							v2 = t89
							if v2 == 0 {
								goto l39
							}
							t90 := int32(load32(m.memory[uint32(v3):]))
							v5 = t90
							t91 := int32(load32(m.memory[uint32(v5+i32(-4)):]))
							v6 = t91
							v0 = v6 & i32(-8)
							t92 := v0
							v6 = v6 & i32(3)
							p93 := i32(8)
							if v6 != 0 {
								p93 = i32(4)
							}
							if uint32(t92) < uint32(p93+v2) {
								m.fn7(i32(1274404), i32(46), i32(1274452))
								panic("unreachable")
							}
							if v6 == 0 {
								goto l41
							}
							if uint32(v0) > uint32(v2+i32(39)) {
								m.fn7(i32(1274468), i32(46), i32(1274516))
								panic("unreachable")
							}
						l41:
							m.fn5(v5)
						}
					l39:
						m.fn260(v3 + i32(88))
						v2 = i32(184)
						goto l35
					}
					v2 = i32(184)
					t84 := int32(load32(m.memory[int64(uint32(v3))+36:]))
					v6 = t84
					if v6 == 0 {
						goto l35
					}
					t85 := int32(load32(m.memory[int64(uint32(v3))+32:]))
					v0 = t85
					t86 := int32(load32(m.memory[uint32(v0+i32(-4)):]))
					v5 = t86
					v7 = v5 & i32(-8)
					t87 := v7
					v5 = v5 & i32(3)
					p88 := i32(8)
					if v5 != 0 {
						p88 = i32(4)
					}
					if uint32(t87) < uint32(p88+v6) {
						m.fn7(i32(1274404), i32(46), i32(1274452))
						panic("unreachable")
					}
					if v5 == 0 {
						goto l37
					}
					if uint32(v7) > uint32(v6+i32(39)) {
						m.fn7(i32(1274468), i32(46), i32(1274516))
						panic("unreachable")
					}
				l37:
					m.fn5(v0)
					goto l35
				}
			}
		}
	l32:
		v2 = i32(72)
	l35:
		t94 := int32(load32(m.memory[uint32(v3+i32(-4)):]))
		v6 = t94
		v5 = v6 & i32(-8)
		t95 := v5
		v6 = v6 & i32(3)
		p96 := i32(8)
		if v6 != 0 {
			p96 = i32(4)
		}
		if uint32(t95) < uint32(p96+v2) {
			m.fn7(i32(1274404), i32(46), i32(1274452))
			panic("unreachable")
		}
		if v6 == 0 {
			goto l44
		}
		if uint32(v5) > uint32(v2+i32(39)) {
			m.fn7(i32(1274468), i32(46), i32(1274516))
			panic("unreachable")
		}
	l44:
		m.fn5(v3)
	}
l31:
	m.g0 = v1 + i32(160)
}
func (m *Module) fn255(v0, v1 int32) int32 {
	var v2, v3, v4, v5 int32
	t0 := m.g0
	v2 = t0 - i32(32)
	m.g0 = v2
	{
		t1 := int32(m.memory[uint32(v0)])
		switch t1 {
		default:
			t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			store32(m.memory[uint32(v2):], uint32(t2))
			t3 := m.fn11(i32(20))
			v0 = t3
			if v0 == 0 {
				m.fn16(i32(1), i32(20))
				panic("unreachable")
			}
			t4 := int32(load32(m.memory[int64(uint32(i32(0)))+1274632:]))
			store32(m.memory[int64(uint32(v0))+16:], uint32(t4))
			t5 := int64(load64(m.memory[int64(uint32(i32(0)))+1274624:]))
			store64(m.memory[int64(uint32(v0))+8:], uint64(t5))
			t6 := int64(load64(m.memory[int64(uint32(i32(0)))+1274616:]))
			store64(m.memory[uint32(v0):], uint64(t6))
			store32(m.memory[int64(uint32(v2))+12:], uint32(i32(20)))
			store32(m.memory[int64(uint32(v2))+8:], uint32(v0))
			store32(m.memory[int64(uint32(v2))+4:], uint32(i32(20)))
			store64(m.memory[int64(uint32(v2))+24:], uint64(int64(uint32(i32(35)))<<32|int64(uint32(v2))))
			store64(m.memory[int64(uint32(v2))+16:], uint64(int64(uint32(i32(36)))<<32|int64(uint32(v2+i32(4)))))
			t7 := int32(load32(m.memory[uint32(v1):]))
			t8 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t9 := m.fn46(t7, t8, i32(1067253), v2+i32(16))
			v0 = t9
			t10 := int32(load32(m.memory[int64(uint32(v2))+4:]))
			v1 = t10
			if v1 == 0 {
				goto l5
			}
			t11 := int32(load32(m.memory[int64(uint32(v2))+8:]))
			v3 = t11
			t12 := int32(load32(m.memory[uint32(v3+i32(-4)):]))
			v4 = t12
			v5 = v4 & i32(-8)
			t13 := v5
			v4 = v4 & i32(3)
			p14 := i32(8)
			if v4 != 0 {
				p14 = i32(4)
			}
			if uint32(t13) < uint32(p14+v1) {
				m.fn7(i32(1274404), i32(46), i32(1274452))
				panic("unreachable")
			}
			if v4 == 0 {
				goto l7
			}
			if uint32(v5) > uint32(v1+i32(39)) {
				m.fn7(i32(1274468), i32(46), i32(1274516))
				panic("unreachable")
			}
		l7:
			m.fn5(v3)
			goto l5
		case 1:
			t15 := int32(load32(m.memory[uint32(v1):]))
			t16 := int32(m.memory[int64(uint32(v0))+1])
			v0 = t16 << 2
			t17 := int32(load32(m.memory[int64(uint32(v0))+1291560:]))
			t18 := int32(load32(m.memory[int64(uint32(v0))+1291392:]))
			t19 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t20 := int32(load32(m.memory[int64(uint32(t19))+12:]))
			t21 := m.t0[uint(t20)].(func(int32, int32, int32) int32)(t15, t17, t18)
			v0 = t21
			goto l5
		case 2:
			t22 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t23 := v1
			v0 = t22
			t24 := int32(load32(m.memory[uint32(v0):]))
			t25 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t26 := m.fn9(t23, t24, t25)
			v0 = t26
			goto l5
		case 3:
			t27 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			v0 = t27
			t28 := int32(load32(m.memory[uint32(v0):]))
			t29 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t30 := int32(load32(m.memory[int64(uint32(t29))+16:]))
			t31 := m.t0[uint(t30)].(func(int32, int32) int32)(t28, v1)
			v0 = t31
		}
	}
l5:
	m.g0 = v2 + i32(32)
	return v0
}
func (m *Module) fn256(v0, v1 int32) {
	var v2 int32
	var v3 int64
	{
		t0 := int32(m.memory[int64(uint32(v0))+8])
		switch t0 + i32(-2) {
		case 0:
			m.fn28(i32(1092188), i32(113), i32(1092172))
			panic("unreachable")
		default:
			m.memory[int64(uint32(v0))+8] = byte(i32(2))
			t1 := int32(load32(m.memory[uint32(v1):]))
			v2 = t1
			store64(m.memory[uint32(v1):], uint64(i64(0)))
			if v2 == 0 {
				m.fn219(i32(1070892))
				panic("unreachable")
			}
			t2 := int64(load64(m.memory[int64(uint32(v1))+8:]))
			v3 = t2
			m.memory[int64(uint32(v0))+8] = byte(i32(3))
			store64(m.memory[uint32(v0):], uint64(v3))
			fallthrough
		case 1:
			return
		}
	}
}
func (m *Module) fn257(v0 int32) {
	var v1, v2, v3 int32
	t0 := m.g0
	v1 = t0 - i32(5344)
	m.g0 = v1
	memory_zero(m.memory, uint32(v1+i32(16)), uint32(i32(5328)))
	{
		t1 := m.fn832(i32(64), i32(47360))
		v2 = t1
		if v2 == 0 {
			store32(m.memory[int64(uint32(v1))+8:], uint32(i32(-4)))
			m.fn904(v1 + i32(8))
			panic("unreachable")
		}
		v3 = (v2 + i32(63)) & i32(-64)
		store64(m.memory[int64(uint32(v3))+48:], uint64(i64(0)))
		store32(m.memory[int64(uint32(v3))+12:], uint32(i32(32832)))
		m.memory[int64(uint32(v3))+4] = byte(i32(0))
		store32(m.memory[uint32(v3):], uint32(i32(1024)))
		store64(m.memory[int64(uint32(v3))+16:], uint64(i64(0)))
		store64(m.memory[int64(uint32(v3))+24:], uint64(i64(0)))
		store64(m.memory[int64(uint32(v3))+32:], uint64(i64(0)))
		store32(m.memory[int64(uint32(v3))+40:], uint32(i32(0)))
		store64(m.memory[int64(uint32(v3))+76:], uint64(i64(0)))
		store32(m.memory[int64(uint32(v3))+72:], uint32(i32(1)))
		m.memory[int64(uint32(v3))+64] = byte(i32(0))
		store64(m.memory[int64(uint32(v3))+56:], uint64(i64(0x100000001)))
		store64(m.memory[int64(uint32(v3))+84:], uint64(i64(0)))
		store64(m.memory[int64(uint32(v3))+92:], uint64(i64(0)))
		store64(m.memory[int64(uint32(v3))+100:], uint64(i64(0)))
		m.memory[int64(uint32(v3))+160] = byte(i32(0))
		store32(m.memory[int64(uint32(v3))+156:], uint32(i32(0)))
		store32(m.memory[int64(uint32(v3))+108:], uint32(i32(32)))
		store32(m.memory[int64(uint32(v3))+8:], uint32(v3+i32(14464)))
		store64(m.memory[int64(uint32(v3))+128:], uint64(i64(0)))
		store64(m.memory[int64(uint32(v3))+120:], uint64(i64(0)))
		store64(m.memory[int64(uint32(v3))+112:], uint64(i64(0)))
		store64(m.memory[int64(uint32(v3))+140:], uint64(i64(0)))
		store64(m.memory[int64(uint32(v3))+145:], uint64(i64(0)))
		memory_copy(m.memory, uint32(v3+i32(161)), uint32(v1+i32(13)), uint32(i32(5331)))
		memory_zero(m.memory, uint32(v3+i32(5492)), uint32(i32(8920)))
		store32(m.memory[int64(uint32(v3))+14408:], uint32(i32(47360)))
		store32(m.memory[int64(uint32(v3))+14404:], uint32(v2))
		store32(m.memory[int64(uint32(v3))+84:], uint32(i32(0)))
		store32(m.memory[int64(uint32(v3))+20:], uint32(i32(0)))
		store64(m.memory[int64(uint32(v3))+120:], uint64(i64(0x1ffffffff)))
		store32(m.memory[uint32(v3):], uint32(i32(984064)))
		m.memory[int64(uint32(v3))+160] = byte(i32(0))
		store32(m.memory[int64(uint32(v3))+156:], uint32(i32(0)))
		store64(m.memory[int64(uint32(v3))+144:], uint64(i64(0x8000)))
		m.memory[int64(uint32(v3))+64] = byte(i32(0))
		store32(m.memory[int64(uint32(v3))+56:], uint32(i32(1)))
		store64(m.memory[int64(uint32(v3))+48:], uint64(i64(0)))
		store32(m.memory[int64(uint32(v3))+100:], uint32(i32(-1)))
		store32(m.memory[int64(uint32(v0))+84:], uint32(i32(0)))
		store64(m.memory[int64(uint32(v0))+76:], uint64(i64(0)))
		store32(m.memory[int64(uint32(v0))+56:], uint32(i32(0)))
		store64(m.memory[int64(uint32(v0))+48:], uint64(i64(0)))
		store64(m.memory[int64(uint32(v0))+40:], uint64(i64(0)))
		store64(m.memory[int64(uint32(v0))+32:], uint64(i64(0)))
		store64(m.memory[int64(uint32(v0))+24:], uint64(i64(0)))
		store64(m.memory[int64(uint32(v0))+16:], uint64(i64(0)))
		store64(m.memory[int64(uint32(v0))+8:], uint64(i64(0)))
		store64(m.memory[uint32(v0):], uint64(i64(0)))
		store32(m.memory[int64(uint32(v0))+72:], uint32(i32(0)))
		store32(m.memory[int64(uint32(v0))+68:], uint32(i32(37)))
		store32(m.memory[int64(uint32(v0))+64:], uint32(i32(38)))
		store32(m.memory[int64(uint32(v0))+60:], uint32(v3))
		m.g0 = v1 + i32(5344)
		return
	}
}
func (m *Module) fn258(v0, v1, v2, v3 int32) {
	var v4, v5 int32
	t0 := m.g0
	v4 = t0 - i32(16)
	m.g0 = v4
	m.fn262(v4+i32(8), v1, v2, v3)
	{
		{
			t1 := int32(m.memory[int64(uint32(v4))+8])
			if t1 == i32(255) {
				goto l0
			}
			t2 := int64(load64(m.memory[int64(uint32(v4))+8:]))
			store64(m.memory[uint32(v0):], uint64(t2))
			goto l1
		}
	l0:
		t3 := int32(load32(m.memory[int64(uint32(v4))+12:]))
		v5 = t3
		{
			t4 := int32(m.memory[int64(uint32(v1))+68])
			if t4 == 0 {
				goto l2
			}
			{
				if v3 == 0 {
					goto l3
				}
				if v5 != 0 {
					goto l3
				}
				t5 := int32(load32(m.memory[int64(uint32(v1))+64:]))
				t6 := int32(load32(m.memory[int64(uint32(v1))+56:]))
				if t5 == t6 {
					goto l3
				}
				m.fn263(v0, i32(21), i32(1080480), i32(16))
				goto l1
			}
		l3:
			if uint32(v5) > uint32(v3) {
				m.fn121(i32(0), v5, v3, i32(1080496))
				panic("unreachable")
			}
			m.fn264(v1+i32(48), v2, v5)
			goto l2
		}
	l2:
		m.memory[uint32(v0)] = byte(i32(255))
		store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
	}
l1:
	m.g0 = v4 + i32(16)
}
func (m *Module) fn259(v0, v1, v2, v3 int32) {
	var v4, v5, v6, v7, v8, v9, v10, v11 int32
	var v12, v13 int64
	t0 := m.g0
	v4 = t0 - i32(32)
	m.g0 = v4
	{
		{
			{
				{
					t1 := int64(load64(m.memory[int64(uint32(v1))+24:]))
					if t1 == i64(2) {
						goto l0
					}
					v5 = v1 + i32(24)
					v6 = v1 + i32(72)
					t2 := int32(load32(m.memory[int64(uint32(v1))+12:]))
					v7 = t2
					t3 := int32(load32(m.memory[int64(uint32(v1))+8:]))
					v8 = t3
				l10:
					{
						t4 := int32(load32(m.memory[uint32(v1):]))
						v9 = t4
						{
							if uint32(v8) < uint32(v7) {
								goto l1
							}
							t5 := int32(load32(m.memory[int64(uint32(v1))+4:]))
							v8 = t5
							{
								t6 := int32(m.memory[int64(uint32(v1))+16])
								if t6 != 0 {
									goto l2
								}
								if v8 == 0 {
									goto l2
								}
								memory_zero(m.memory, uint32(v9), uint32(v8))
							}
						l2:
							m.fn262(v4+i32(16), v5, v9, v8)
							{
								t7 := int32(m.memory[int64(uint32(v4))+16])
								if t7 != i32(255) {
									t9 := int32(load32(m.memory[int64(uint32(v4))+20:]))
									v10 = t9
									t10 := int32(load32(m.memory[int64(uint32(v4))+16:]))
									v11 = t10
									t11 := int64(m.memory[int64(uint32(v4))+16])
									v12 = t11
									m.memory[int64(uint32(v1))+16] = byte(i32(1))
									store64(m.memory[int64(uint32(v1))+8:], uint64(i64(0)))
									v7 = i32(0)
									v8 = i32(0)
									if v12 == i64(255) {
										goto l1
									}
									store32(m.memory[int64(uint32(v4))+12:], uint32(v10))
									store32(m.memory[int64(uint32(v4))+8:], uint32(v11))
									goto l5
								}
								{
									t8 := int32(load32(m.memory[int64(uint32(v4))+20:]))
									v7 = t8
									if uint32(v7) > uint32(v8) {
										m.fn7(i32(1069306), i32(36), i32(1069344))
										panic("unreachable")
									}
									m.memory[int64(uint32(v1))+16] = byte(i32(1))
									store32(m.memory[int64(uint32(v1))+12:], uint32(v7))
									v8 = i32(0)
									store32(m.memory[int64(uint32(v1))+8:], uint32(i32(0)))
									goto l1
								}
							}
						}
					l1:
						t12 := int64(load64(m.memory[int64(uint32(v1))+80:]))
						v13 = t12
						t13 := int64(load64(m.memory[int64(uint32(v1))+72:]))
						v12 = t13
						t14 := v4 + i32(16)
						t15 := v6
						t16 := v9 + v8
						t17 := v7 - v8
						t18 := v2
						t19 := v3
						var p20 int32
						if v7 == v8 {
							p20 = 1
						}
						v11 = p20
						p21 := i32(0)
						if v11 != 0 {
							p21 = i32(4)
						}
						m.fn265(t14, t15, t16, t17, t18, t19, p21)
						t22 := int32(m.memory[int64(uint32(v4))+20])
						v10 = t22
						t23 := int32(load32(m.memory[int64(uint32(v4))+16:]))
						v9 = t23
						t24 := int32(load32(m.memory[int64(uint32(v1))+12:]))
						t25 := v1
						v7 = t24
						t26 := int32(load32(m.memory[int64(uint32(v1))+8:]))
						t27 := int64(load64(m.memory[int64(uint32(v1))+72:]))
						t28 := v7
						v8 = t26 + int32(t27-v12)
						p29 := v8
						if uint32(v7) < uint32(v8) {
							p29 = t28
						}
						v8 = p29
						store32(m.memory[int64(uint32(t25))+8:], uint32(v8))
						if v9 == i32(2) {
							t30 := int64(load64(m.memory[int64(uint32(v1))+80:]))
							v9 = int32(t30 - v13)
							switch v10 {
							case 2:
								goto l9
							default:
								if v11 != 0 {
									goto l9
								}
								if v3 == 0 {
									goto l9
								}
								if v9 == 0 {
									goto l10
								}
								goto l9
							case 1:
								if v11 != 0 {
									goto l9
								}
								if v3 == 0 {
									goto l9
								}
								if v9 == 0 {
									goto l10
								}
								goto l9
							}
						}
						m.fn263(v4+i32(8), i32(20), i32(1069820), i32(22))
						goto l5
					}
				}
			l0:
				t31 := int32(load32(m.memory[int64(uint32(v1))+36:]))
				v9 = t31
				{
					{
						{
							{
								t32 := int32(load32(m.memory[int64(uint32(v1))+40:]))
								v8 = t32
								t33 := int32(load32(m.memory[int64(uint32(v1))+44:]))
								t34 := v8
								v7 = t33
								if t34 != v7 {
									goto l11
								}
								if uint32(v3) >= uint32(v9) {
									store64(m.memory[int64(uint32(v1))+40:], uint64(i64(0)))
									m.fn262(v4+i32(8), v1+i32(56), v2, v3)
									goto l5
								}
							}
						l11:
							t35 := int32(load32(m.memory[int64(uint32(v1))+32:]))
							v10 = t35
							if uint32(v8) < uint32(v7) {
								goto l13
							}
							v7 = v1 + i32(56)
							{
								t36 := int32(m.memory[int64(uint32(v1))+48])
								if t36 != 0 {
									goto l14
								}
								if v9 == 0 {
									goto l14
								}
								memory_zero(m.memory, uint32(v10), uint32(v9))
							}
						l14:
							m.fn262(v4+i32(16), v7, v10, v9)
							t37 := int32(m.memory[int64(uint32(v4))+16])
							if t37 != i32(255) {
								goto l15
							}
							{
								t38 := int32(load32(m.memory[int64(uint32(v4))+20:]))
								v7 = t38
								if uint32(v7) > uint32(v9) {
									m.fn7(i32(1069306), i32(36), i32(1069344))
									panic("unreachable")
								}
								m.memory[int64(uint32(v1))+48] = byte(i32(1))
								store32(m.memory[int64(uint32(v1))+44:], uint32(v7))
								v8 = i32(0)
								goto l13
							}
						}
					l15:
						t39 := int32(load32(m.memory[int64(uint32(v4))+20:]))
						v9 = t39
						t40 := int32(load32(m.memory[int64(uint32(v4))+16:]))
						v11 = t40
						t41 := int64(m.memory[int64(uint32(v4))+16])
						v12 = t41
						m.memory[int64(uint32(v1))+48] = byte(i32(1))
						store64(m.memory[int64(uint32(v1))+40:], uint64(i64(0)))
						v7 = i32(0)
						v8 = i32(0)
						if v12 != i64(255) {
							goto l17
						}
					}
				l13:
					v10 = v10 + v8
					{
						v9 = v7 - v8
						p42 := v3
						if uint32(v9) < uint32(v3) {
							p42 = v9
						}
						v9 = p42
						if v9 != i32(1) {
							goto l18
						}
						t43 := int32(m.memory[uint32(v10)])
						m.memory[uint32(v2)] = byte(t43)
						goto l19
					}
				l18:
					if v9 == 0 {
						goto l19
					}
					memory_copy(m.memory, uint32(v2), uint32(v10), uint32(v9))
				l19:
					t44 := v1
					t45 := v7
					v8 = v9 + v8
					p46 := v8
					if uint32(v7) < uint32(v8) {
						p46 = t45
					}
					store32(m.memory[int64(uint32(t44))+40:], uint32(p46))
					goto l9
				}
			l17:
				store32(m.memory[int64(uint32(v4))+12:], uint32(v9))
				store32(m.memory[int64(uint32(v4))+8:], uint32(v11))
			}
		l5:
			t47 := int32(m.memory[int64(uint32(v4))+8])
			if t47 != i32(255) {
				t52 := int64(load64(m.memory[int64(uint32(v4))+8:]))
				store64(m.memory[uint32(v0):], uint64(t52))
				goto l23
			}
			t48 := int32(load32(m.memory[int64(uint32(v4))+12:]))
			v9 = t48
		}
	l9:
		t49 := int32(m.memory[int64(uint32(v1))+180])
		if t49 == 0 {
			goto l21
		}
		if v3 == 0 {
			goto l22
		}
		if v9 != 0 {
			goto l22
		}
		t50 := int32(load32(m.memory[int64(uint32(v1))+176:]))
		t51 := int32(load32(m.memory[int64(uint32(v1))+168:]))
		if t50 == t51 {
			goto l22
		}
		m.fn263(v0, i32(21), i32(1080480), i32(16))
		goto l23
	}
l22:
	if uint32(v9) > uint32(v3) {
		m.fn121(i32(0), v9, v3, i32(1080496))
		panic("unreachable")
	}
	m.fn264(v1+i32(160), v2, v9)
	goto l21
l21:
	m.memory[uint32(v0)] = byte(i32(255))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v9))
l23:
	m.g0 = v4 + i32(32)
}
func (m *Module) fn260(v0 int32) {
	var v1, v2, v3, v4 int32
	t0 := m.g0
	v1 = t0 - i32(32)
	m.g0 = v1
	t1 := int32(load32(m.memory[int64(uint32(v0))+44:]))
	v2 = t1
	store64(m.memory[int64(uint32(v2))+16:], uint64(i64(0)))
	store64(m.memory[int64(uint32(v2))+8:], uint64(i64(1)))
	store32(m.memory[int64(uint32(v0))+44:], uint32(i32(0)))
	t2 := int32(load32(m.memory[int64(uint32(v0))+56:]))
	v3 = t2
	t3 := int32(load32(m.memory[int64(uint32(v0))+52:]))
	v4 = t3
	t4 := int32(load32(m.memory[int64(uint32(v2))+14404:]))
	v0 = t4
	t5 := int32(load32(m.memory[int64(uint32(v2))+14408:]))
	t6 := v1
	v2 = t5
	store32(m.memory[int64(uint32(t6))+12:], uint32(v2))
	store32(m.memory[int64(uint32(v1))+8:], uint32(v0))
	{
		if v0 == 0 {
			goto l0
		}
		{
			if v4 == i32(37) {
				goto l1
			}
			t7 := int32(load32(m.memory[uint32(v0+i32(-4)):]))
			m.t0[uint(v4)].(func(int32, int32))(v3, t7)
			goto l0
		}
	l1:
		if v2 == 0 {
			store64(m.memory[int64(uint32(v1))+16:], uint64(int64(uint32(i32(39)))<<32|int64(uint32(v1+i32(8)))))
			m.fn633(i32(1), v1+i32(12), i32(1277668), i32(1050185), v1+i32(16), i32(1277672))
			panic("unreachable")
		}
		if uint32(v2) >= uint32(i32(0x7fffffc1)) {
			m.fn42(i32(1284936), i32(43), v1+i32(31), i32(1285048), i32(1285064))
			panic("unreachable")
		}
		t8 := int32(load32(m.memory[uint32(v0+i32(-4)):]))
		v4 = t8
		v3 = v4 & i32(-8)
		t9 := v3
		v4 = v4 & i32(3)
		p10 := i32(8)
		if v4 != 0 {
			p10 = i32(4)
		}
		if uint32(t9) < uint32(p10+v2) {
			m.fn7(i32(1274404), i32(46), i32(1274452))
			panic("unreachable")
		}
		if v4 == 0 {
			goto l5
		}
		if uint32(v3) > uint32(v2+i32(39)) {
			m.fn7(i32(1274468), i32(46), i32(1274516))
			panic("unreachable")
		}
	l5:
		m.fn5(v0)
	}
l0:
	m.g0 = v1 + i32(32)
}
func (m *Module) fn261(v0 int32) {
	var v1, v2, v3 int32
	{
		t0 := int32(load32(m.memory[uint32(v0):]))
		v1 = t0
		if v1 == 0 {
			return
		}
		t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
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
func (m *Module) fn262(v0, v1, v2, v3 int32) {
	var v4 int64
	var v5, v6 int32
	var v7 int64
	var v8 int32
	var v9 int64
	var v10, v11, v12 int32
	{
		{
			t0 := int64(load64(m.memory[uint32(v1):]))
			if t0 != i64(1) {
				goto l0
			}
			{
				t1 := int64(load64(m.memory[int64(uint32(v1))+16:]))
				v4 = t1
				if !(v4 == 0) {
					t2 := int32(load32(m.memory[int64(uint32(v1))+24:]))
					v6 = t2
					t3 := int32(load32(m.memory[uint32(v6):]))
					t4 := int64(load64(m.memory[int64(uint32(v6))+8:]))
					v7 = t4
					t5 := int32(load32(m.memory[int64(uint32(v6))+4:]))
					t6 := v7
					v8 = t5
					v9 = int64(uint32(v8))
					p7 := v9
					if uint64(v7) < uint64(v9) {
						p7 = t6
					}
					v10 = t3 + int32(p7)
					{
						t9 := v8
						p8 := i64(0xffffffff)
						if uint64(v7) < uint64(i64(0xffffffff)) {
							p8 = v7
						}
						v11 = t9 - int32(p8)
						p10 := v11
						if uint32(v11) > uint32(v8) {
							p10 = i32(0)
						}
						v8 = p10
						t11 := v8
						t12 := v4
						v9 = int64(uint32(v3))
						p13 := v9
						if uint64(v4) < uint64(v9) {
							p13 = t12
						}
						v11 = int32(p13)
						p14 := v11
						if uint32(v8) < uint32(v11) {
							p14 = t11
						}
						v5 = p14
						if v5 != i32(1) {
							goto l3
						}
						t15 := int32(m.memory[uint32(v10)])
						m.memory[uint32(v2)] = byte(t15)
						goto l4
					}
				l3:
					if v5 == 0 {
						goto l4
					}
					memory_copy(m.memory, uint32(v2), uint32(v10), uint32(v5))
				l4:
					t16 := v1
					t17 := v4
					v9 = int64(uint32(v5))
					store64(m.memory[int64(uint32(t16))+16:], uint64(t17-v9))
					store64(m.memory[int64(uint32(v6))+8:], uint64(v7+v9))
					if v5 != 0 {
						v11 = v3 + i32(-1)
						v8 = i32(1) - v5
						t18 := int32(load32(m.memory[int64(uint32(v1))+36:]))
						v10 = t18
						t19 := int32(load32(m.memory[int64(uint32(v1))+32:]))
						v6 = t19
						t20 := int32(load32(m.memory[int64(uint32(v1))+40:]))
						v3 = t20
					l7:
						{
							t21 := int32(m.memory[uint32(v2)])
							t22 := v2
							v12 = v3 | i32(2)
							v12 = t21 ^ int32(uint32((v12^i32(1))*v12&i32(0xff00))>>8)
							m.memory[uint32(t22)] = byte(v12)
							t23 := int32(load32(m.memory[int64(uint32((v12^v6)&i32(255)<<2))+1285096:]))
							t24 := v10
							v6 = int32(uint32(v6)>>8) ^ t23
							v10 = (t24+v6&i32(255))*i32(134775813) + i32(1)
							t25 := int32(load32(m.memory[int64(uint32((int32(uint32(v10)>>24)^v3&i32(255))<<2))+1285096:]))
							v3 = t25 ^ int32(uint32(v3)>>8)
							if v8 == 0 {
								goto l6
							}
							v8 = v8 + i32(1)
							v2 = v2 + i32(1)
							v12 = v11
							v11 = v11 + i32(-1)
							if v12 != 0 {
								goto l7
							}
						}
					l6:
						store32(m.memory[int64(uint32(v1))+36:], uint32(v10))
						store32(m.memory[int64(uint32(v1))+32:], uint32(v6))
						store32(m.memory[int64(uint32(v1))+40:], uint32(v3))
						goto l2
					}
					v5 = i32(0)
					goto l2
				}
				v5 = i32(0)
				goto l2
			}
		}
	l0:
		{
			t26 := int64(load64(m.memory[int64(uint32(v1))+16:]))
			v4 = t26
			if !(v4 == 0) {
				goto l8
			}
			v5 = i32(0)
			goto l2
		}
	l8:
		t27 := int32(load32(m.memory[int64(uint32(v1))+24:]))
		v6 = t27
		t28 := int32(load32(m.memory[uint32(v6):]))
		t29 := int64(load64(m.memory[int64(uint32(v6))+8:]))
		v7 = t29
		t30 := int32(load32(m.memory[int64(uint32(v6))+4:]))
		t31 := v7
		v8 = t30
		v9 = int64(uint32(v8))
		p32 := v9
		if uint64(v7) < uint64(v9) {
			p32 = t31
		}
		v10 = t28 + int32(p32)
		{
			t34 := v8
			p33 := i64(0xffffffff)
			if uint64(v7) < uint64(i64(0xffffffff)) {
				p33 = v7
			}
			v11 = t34 - int32(p33)
			p35 := v11
			if uint32(v11) > uint32(v8) {
				p35 = i32(0)
			}
			v8 = p35
			t36 := v8
			t37 := v4
			v9 = int64(uint32(v3))
			p38 := v9
			if uint64(v4) < uint64(v9) {
				p38 = t37
			}
			v3 = int32(p38)
			p39 := v3
			if uint32(v8) < uint32(v3) {
				p39 = t36
			}
			v5 = p39
			if v5 != i32(1) {
				goto l9
			}
			t40 := int32(m.memory[uint32(v10)])
			m.memory[uint32(v2)] = byte(t40)
			goto l10
		}
	l9:
		if v5 == 0 {
			goto l10
		}
		memory_copy(m.memory, uint32(v2), uint32(v10), uint32(v5))
	l10:
		t41 := v1
		t42 := v4
		v9 = int64(uint32(v5))
		store64(m.memory[int64(uint32(t41))+16:], uint64(t42-v9))
		store64(m.memory[int64(uint32(v6))+8:], uint64(v7+v9))
	}
l2:
	m.memory[uint32(v0)] = byte(i32(255))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
}
func (m *Module) fn263(v0, v1, v2, v3 int32) {
	var v4 int32
	{
		t0 := m.fn11(v3)
		v4 = t0
		if v4 == 0 {
			m.fn16(i32(1), v3)
			panic("unreachable")
		}
		if v3 == 0 {
			goto l1
		}
		memory_copy(m.memory, uint32(v4), uint32(v2), uint32(v3))
	l1:
		t1 := m.fn11(i32(12))
		v2 = t1
		if v2 == 0 {
			m.fn23(i32(4), i32(12))
			panic("unreachable")
		}
		store32(m.memory[int64(uint32(v2))+8:], uint32(v3))
		store32(m.memory[int64(uint32(v2))+4:], uint32(v4))
		store32(m.memory[uint32(v2):], uint32(v3))
		t2 := m.fn11(i32(12))
		v3 = t2
		if v3 == 0 {
			m.fn23(i32(4), i32(12))
			panic("unreachable")
		}
		m.memory[int64(uint32(v3))+8] = byte(v1)
		store32(m.memory[int64(uint32(v3))+4:], uint32(i32(1093460)))
		store32(m.memory[uint32(v3):], uint32(v2))
		store64(m.memory[uint32(v0):], uint64(int64(uint32(v3))<<32|i64(3)))
		return
	}
}
func (m *Module) fn264(v0, v1, v2 int32) {
	var v3, v4, v5 int32
	t0 := int64(load64(m.memory[uint32(v0):]))
	store64(m.memory[uint32(v0):], uint64(t0+int64(uint32(v2))))
	t1 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	v3 = t1 ^ i32(-1)
	if uint32(v2) < uint32(i32(64)) {
		goto l0
	}
l1:
	{
		t2 := int32(m.memory[uint32(v1+i32(62))])
		t3 := int32(load32(m.memory[int64(uint32(t2<<2))+1124276:]))
		t4 := int32(m.memory[uint32(v1+i32(63))])
		t5 := int32(load32(m.memory[int64(uint32(t4<<2))+1123252:]))
		t6 := int32(m.memory[uint32(v1+i32(61))])
		t7 := int32(load32(m.memory[int64(uint32(t6<<2))+1125300:]))
		t8 := int32(m.memory[uint32(v1+i32(60))])
		t9 := int32(load32(m.memory[int64(uint32(t8<<2))+1126324:]))
		t10 := int32(m.memory[uint32(v1+i32(59))])
		t11 := int32(load32(m.memory[int64(uint32(t10<<2))+1127348:]))
		t12 := int32(m.memory[uint32(v1+i32(58))])
		t13 := int32(load32(m.memory[int64(uint32(t12<<2))+1128372:]))
		t14 := int32(m.memory[uint32(v1+i32(57))])
		t15 := int32(load32(m.memory[int64(uint32(t14<<2))+1129396:]))
		t16 := int32(m.memory[uint32(v1+i32(56))])
		t17 := int32(load32(m.memory[int64(uint32(t16<<2))+1130420:]))
		t18 := int32(m.memory[uint32(v1+i32(55))])
		t19 := int32(load32(m.memory[int64(uint32(t18<<2))+1131444:]))
		t20 := int32(m.memory[uint32(v1+i32(54))])
		t21 := int32(load32(m.memory[int64(uint32(t20<<2))+1132468:]))
		t22 := int32(m.memory[uint32(v1+i32(53))])
		t23 := int32(load32(m.memory[int64(uint32(t22<<2))+1133492:]))
		t24 := int32(m.memory[uint32(v1+i32(52))])
		t25 := int32(load32(m.memory[int64(uint32(t24<<2))+1134516:]))
		t26 := int32(m.memory[uint32(v1+i32(47))])
		t27 := int32(load32(m.memory[int64(uint32(t26<<2))+1123252:]))
		t28 := int32(m.memory[uint32(v1+i32(46))])
		t29 := int32(load32(m.memory[int64(uint32(t28<<2))+1124276:]))
		t30 := int32(m.memory[uint32(v1+i32(45))])
		t31 := int32(load32(m.memory[int64(uint32(t30<<2))+1125300:]))
		t32 := int32(m.memory[uint32(v1+i32(44))])
		t33 := int32(load32(m.memory[int64(uint32(t32<<2))+1126324:]))
		t34 := int32(m.memory[uint32(v1+i32(43))])
		t35 := int32(load32(m.memory[int64(uint32(t34<<2))+1127348:]))
		t36 := int32(m.memory[uint32(v1+i32(42))])
		t37 := int32(load32(m.memory[int64(uint32(t36<<2))+1128372:]))
		t38 := int32(m.memory[uint32(v1+i32(41))])
		t39 := int32(load32(m.memory[int64(uint32(t38<<2))+1129396:]))
		t40 := int32(m.memory[uint32(v1+i32(40))])
		t41 := int32(load32(m.memory[int64(uint32(t40<<2))+1130420:]))
		t42 := int32(m.memory[uint32(v1+i32(39))])
		t43 := int32(load32(m.memory[int64(uint32(t42<<2))+1131444:]))
		t44 := int32(m.memory[uint32(v1+i32(38))])
		t45 := int32(load32(m.memory[int64(uint32(t44<<2))+1132468:]))
		t46 := int32(m.memory[uint32(v1+i32(37))])
		t47 := int32(load32(m.memory[int64(uint32(t46<<2))+1133492:]))
		t48 := int32(m.memory[uint32(v1+i32(36))])
		t49 := int32(load32(m.memory[int64(uint32(t48<<2))+1134516:]))
		t50 := int32(m.memory[uint32(v1+i32(31))])
		t51 := int32(load32(m.memory[int64(uint32(t50<<2))+1123252:]))
		t52 := int32(m.memory[uint32(v1+i32(30))])
		t53 := int32(load32(m.memory[int64(uint32(t52<<2))+1124276:]))
		t54 := int32(m.memory[uint32(v1+i32(29))])
		t55 := int32(load32(m.memory[int64(uint32(t54<<2))+1125300:]))
		t56 := int32(m.memory[uint32(v1+i32(28))])
		t57 := int32(load32(m.memory[int64(uint32(t56<<2))+1126324:]))
		t58 := int32(m.memory[uint32(v1+i32(27))])
		t59 := int32(load32(m.memory[int64(uint32(t58<<2))+1127348:]))
		t60 := int32(m.memory[uint32(v1+i32(26))])
		t61 := int32(load32(m.memory[int64(uint32(t60<<2))+1128372:]))
		t62 := int32(m.memory[uint32(v1+i32(25))])
		t63 := int32(load32(m.memory[int64(uint32(t62<<2))+1129396:]))
		t64 := int32(m.memory[uint32(v1+i32(24))])
		t65 := int32(load32(m.memory[int64(uint32(t64<<2))+1130420:]))
		t66 := int32(m.memory[uint32(v1+i32(23))])
		t67 := int32(load32(m.memory[int64(uint32(t66<<2))+1131444:]))
		t68 := int32(m.memory[uint32(v1+i32(22))])
		t69 := int32(load32(m.memory[int64(uint32(t68<<2))+1132468:]))
		t70 := int32(m.memory[uint32(v1+i32(21))])
		t71 := int32(load32(m.memory[int64(uint32(t70<<2))+1133492:]))
		t72 := int32(m.memory[uint32(v1+i32(20))])
		t73 := int32(load32(m.memory[int64(uint32(t72<<2))+1134516:]))
		t74 := int32(m.memory[uint32(v1+i32(15))])
		t75 := int32(load32(m.memory[int64(uint32(t74<<2))+1123252:]))
		t76 := int32(m.memory[uint32(v1+i32(14))])
		t77 := int32(load32(m.memory[int64(uint32(t76<<2))+1124276:]))
		t78 := int32(m.memory[uint32(v1+i32(13))])
		t79 := int32(load32(m.memory[int64(uint32(t78<<2))+1125300:]))
		t80 := int32(m.memory[uint32(v1+i32(12))])
		t81 := int32(load32(m.memory[int64(uint32(t80<<2))+1126324:]))
		t82 := int32(m.memory[uint32(v1+i32(11))])
		t83 := int32(load32(m.memory[int64(uint32(t82<<2))+1127348:]))
		t84 := int32(m.memory[uint32(v1+i32(10))])
		t85 := int32(load32(m.memory[int64(uint32(t84<<2))+1128372:]))
		t86 := int32(m.memory[uint32(v1+i32(9))])
		t87 := int32(load32(m.memory[int64(uint32(t86<<2))+1129396:]))
		t88 := int32(m.memory[uint32(v1+i32(8))])
		t89 := int32(load32(m.memory[int64(uint32(t88<<2))+1130420:]))
		t90 := int32(m.memory[uint32(v1+i32(7))])
		t91 := int32(load32(m.memory[int64(uint32(t90<<2))+1131444:]))
		t92 := int32(m.memory[uint32(v1+i32(6))])
		t93 := int32(load32(m.memory[int64(uint32(t92<<2))+1132468:]))
		t94 := int32(m.memory[uint32(v1+i32(5))])
		t95 := int32(load32(m.memory[int64(uint32(t94<<2))+1133492:]))
		t96 := int32(m.memory[uint32(v1+i32(4))])
		t97 := int32(load32(m.memory[int64(uint32(t96<<2))+1134516:]))
		t98 := int32(m.memory[uint32(v1+i32(3))])
		t99 := int32(load32(m.memory[int64(uint32((int32(uint32(v3)>>24)^t98)<<2))+1135540:]))
		t100 := int32(m.memory[uint32(v1+i32(2))])
		t101 := int32(load32(m.memory[int64(uint32((int32(uint32(v3)>>16)&i32(255)^t100)<<2))+1136564:]))
		t102 := int32(m.memory[uint32(v1+i32(1))])
		t103 := int32(load32(m.memory[int64(uint32((int32(uint32(v3)>>8)&i32(255)^t102)<<2))+1137588:]))
		t104 := int32(m.memory[uint32(v1)])
		t105 := int32(load32(m.memory[int64(uint32((v3&i32(255)^t104)<<2))+1138612:]))
		t106 := t3 ^ t5 ^ t7 ^ t9 ^ t11 ^ t13 ^ t15 ^ t17 ^ t19 ^ t21 ^ t23 ^ t25
		t107 := t27 ^ t29 ^ t31 ^ t33 ^ t35 ^ t37 ^ t39 ^ t41 ^ t43 ^ t45 ^ t47 ^ t49
		t108 := t51 ^ t53 ^ t55 ^ t57 ^ t59 ^ t61 ^ t63 ^ t65 ^ t67 ^ t69 ^ t71 ^ t73
		v3 = t75 ^ t77 ^ t79 ^ t81 ^ t83 ^ t85 ^ t87 ^ t89 ^ t91 ^ t93 ^ t95 ^ t97 ^ t99 ^ t101 ^ t103 ^ t105
		t109 := int32(m.memory[uint32(v1+i32(19))])
		t110 := int32(load32(m.memory[int64(uint32((int32(uint32(v3)>>24)^t109)<<2))+1135540:]))
		t111 := int32(m.memory[uint32(v1+i32(18))])
		t112 := int32(load32(m.memory[int64(uint32((int32(uint32(v3)>>16)&i32(255)^t111)<<2))+1136564:]))
		t113 := int32(m.memory[uint32(v1+i32(17))])
		t114 := int32(load32(m.memory[int64(uint32((int32(uint32(v3)>>8)&i32(255)^t113)<<2))+1137588:]))
		t115 := int32(m.memory[uint32(v1+i32(16))])
		t116 := int32(load32(m.memory[int64(uint32((v3&i32(255)^t115)<<2))+1138612:]))
		v3 = t108 ^ t110 ^ t112 ^ t114 ^ t116
		t117 := int32(m.memory[uint32(v1+i32(35))])
		t118 := int32(load32(m.memory[int64(uint32((int32(uint32(v3)>>24)^t117)<<2))+1135540:]))
		t119 := int32(m.memory[uint32(v1+i32(34))])
		t120 := int32(load32(m.memory[int64(uint32((int32(uint32(v3)>>16)&i32(255)^t119)<<2))+1136564:]))
		t121 := int32(m.memory[uint32(v1+i32(33))])
		t122 := int32(load32(m.memory[int64(uint32((int32(uint32(v3)>>8)&i32(255)^t121)<<2))+1137588:]))
		t123 := int32(m.memory[uint32(v1+i32(32))])
		t124 := int32(load32(m.memory[int64(uint32((v3&i32(255)^t123)<<2))+1138612:]))
		v3 = t107 ^ t118 ^ t120 ^ t122 ^ t124
		t125 := int32(m.memory[uint32(v1+i32(51))])
		t126 := int32(load32(m.memory[int64(uint32((int32(uint32(v3)>>24)^t125)<<2))+1135540:]))
		t127 := int32(m.memory[uint32(v1+i32(50))])
		t128 := int32(load32(m.memory[int64(uint32((int32(uint32(v3)>>16)&i32(255)^t127)<<2))+1136564:]))
		t129 := int32(m.memory[uint32(v1+i32(49))])
		t130 := int32(load32(m.memory[int64(uint32((int32(uint32(v3)>>8)&i32(255)^t129)<<2))+1137588:]))
		t131 := int32(m.memory[uint32(v1+i32(48))])
		t132 := int32(load32(m.memory[int64(uint32((v3&i32(255)^t131)<<2))+1138612:]))
		v3 = t106 ^ t126 ^ t128 ^ t130 ^ t132
		v1 = v1 + i32(64)
		v2 = v2 + i32(-64)
		if uint32(v2) > uint32(i32(63)) {
			goto l1
		}
	}
l0:
	if v2 == 0 {
		goto l2
	}
	v4 = v2 & i32(3)
	if v4 != 0 {
		goto l3
	}
	v5 = v1
	goto l4
l3:
	v5 = v1
l5:
	{
		t133 := int32(m.memory[uint32(v5)])
		t134 := int32(load32(m.memory[int64(uint32((t133^v3)&i32(255)<<2))+1123252:]))
		v3 = t134 ^ int32(uint32(v3)>>8)
		v5 = v5 + i32(1)
		v4 = v4 + i32(-1)
		if v4 != 0 {
			goto l5
		}
	}
l4:
	if uint32(v2) < uint32(i32(4)) {
		goto l2
	}
	v2 = v1 + v2
l6:
	{
		t135 := int32(m.memory[uint32(v5+i32(3))])
		t136 := int32(m.memory[uint32(v5+i32(2))])
		t137 := int32(m.memory[uint32(v5+i32(1))])
		t138 := int32(m.memory[uint32(v5)])
		t139 := int32(load32(m.memory[int64(uint32((t138^v3)&i32(255)<<2))+1123252:]))
		v1 = t139 ^ int32(uint32(v3)>>8)
		t140 := int32(load32(m.memory[int64(uint32((t137^v1)&i32(255)<<2))+1123252:]))
		v1 = t140 ^ int32(uint32(v1)>>8)
		t141 := int32(load32(m.memory[int64(uint32((t136^v1)&i32(255)<<2))+1123252:]))
		v1 = t141 ^ int32(uint32(v1)>>8)
		t142 := int32(load32(m.memory[int64(uint32((t135^v1)&i32(255)<<2))+1123252:]))
		v3 = t142 ^ int32(uint32(v1)>>8)
		v5 = v5 + i32(4)
		if v5 != v2 {
			goto l6
		}
	}
l2:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v3^i32(-1)))
}
func (m *Module) fn265(v0, v1, v2, v3, v4, v5, v6 int32) {
	var v7, v8, v9 int32
	var v10, v11 int64
	var v12, v13, v14, v15, v16, v17, v18, v19, v20, v21, v22, v23, v24 int32
	var v25, v26 int64
	var v27, v28, v29 int32
	var v30, v31 int64
	var v32, v33, v34 int32
	t0 := m.g0
	v7 = t0 - i32(32)
	m.g0 = v7
	store32(m.memory[int64(uint32(v1))+48:], uint32(v5))
	store32(m.memory[int64(uint32(v1))+36:], uint32(v3))
	store32(m.memory[int64(uint32(v1))+44:], uint32(v4))
	store32(m.memory[int64(uint32(v1))+32:], uint32(v2))
	{
		t1 := int32(load32(m.memory[int64(uint32(v1))+60:]))
		v8 = t1
		t2 := int32(m.memory[uint32(v8)])
		v9 = t2
		if v9 != i32(12) {
			goto l0
		}
		v9 = i32(13)
		m.memory[uint32(v8)] = byte(i32(13))
	}
l0:
	t3 := int64(load64(m.memory[int64(uint32(v1))+24:]))
	v10 = t3
	t4 := int64(load64(m.memory[int64(uint32(v1))+16:]))
	v11 = t4
	store32(m.memory[int64(uint32(v8))+116:], uint32(v5))
	store32(m.memory[int64(uint32(v8))+112:], uint32(v3))
	store32(m.memory[int64(uint32(v8))+80:], uint32(i32(0)))
	store32(m.memory[int64(uint32(v8))+76:], uint32(v5))
	store32(m.memory[int64(uint32(v8))+72:], uint32(v4))
	store32(m.memory[int64(uint32(v8))+56:], uint32(v2))
	m.memory[int64(uint32(v8))+4] = byte(v6)
	store32(m.memory[int64(uint32(v8))+60:], uint32(v2+v3))
	v12 = v8 + i32(65)
	v13 = v8 + i32(8)
	v14 = v8 + i32(13828)
	v15 = v8 + i32(13188)
	v16 = v8 + i32(164)
	v17 = v8 + i32(5492)
	v18 = v8 + i32(10820)
	v19 = v8 + i32(72)
l18:
	v5 = v9 & i32(255)
	v9 = i32(18)
	v20 = i32(29)
	v3 = i32(1)
	switch v5 {
	case 1:
		t629 := int64(load64(m.memory[int64(uint32(v8))+48:]))
		v25 = t629
		{
			t630 := int32(m.memory[int64(uint32(v8))+64])
			v9 = t630
			if uint32(v9) > uint32(i32(15)) {
				goto l264
			}
			v3 = i32(0)
			{
				t631 := int32(load32(m.memory[int64(uint32(v8))+56:]))
				v5 = t631
				t632 := int32(load32(m.memory[int64(uint32(v8))+60:]))
				t633 := v5
				v29 = t632
				if t633 != v29 {
					goto l265
				}
				v20 = i32(1)
				goto l30
			}
		l265:
			v20 = i32(1)
			t634 := v8
			v27 = v5 + i32(1)
			store32(m.memory[int64(uint32(t634))+56:], uint32(v27))
			t635 := v8
			v30 = int64(uint32(v9))
			v26 = v30 + i64(8)
			m.memory[int64(uint32(t635))+64] = byte(v26)
			t636 := int64(m.memory[uint32(v5)])
			t637 := v8
			v25 = i64_shl(t636, v30) | v25
			store64(m.memory[int64(uint32(t637))+48:], uint64(v25))
			if uint32(v9) > uint32(i32(7)) {
				goto l264
			}
			if v27 == v29 {
				goto l30
			}
			store32(m.memory[int64(uint32(v8))+56:], uint32(v5+i32(2)))
			m.memory[int64(uint32(v8))+64] = byte(v9 | i32(16))
			t638 := int64(m.memory[int64(uint32(v5))+1])
			t639 := v8
			v25 = i64_shl(t638, v26) | v25
			store64(m.memory[int64(uint32(t639))+48:], uint64(v25))
		}
	l264:
		t640 := v8
		v9 = int32(v25)
		store32(m.memory[int64(uint32(t640))+120:], uint32(v9))
		{
			if v9&i32(255) != i32(8) {
				store32(m.memory[int64(uint32(v8))+136:], uint32(i32(27)))
				store32(m.memory[int64(uint32(v8))+132:], uint32(i32(1066020)))
				goto l57
			}
			if v9&i32(57344) != 0 {
				store32(m.memory[int64(uint32(v8))+136:], uint32(i32(25)))
				store32(m.memory[int64(uint32(v8))+132:], uint32(i32(1052941)))
				goto l57
			}
			{
				t641 := int32(load32(m.memory[int64(uint32(v8))+140:]))
				v5 = t641
				if v5 == 0 {
					goto l268
				}
				store32(m.memory[uint32(v5):], uint32(int32(uint32(v9)>>8)&i32(1)))
				t642 := int32(load32(m.memory[int64(uint32(v8))+120:]))
				v9 = t642
			}
		l268:
			if v9&i32(512) == 0 {
				goto l269
			}
			t643 := int32(m.memory[int64(uint32(v8))+3])
			if t643&i32(4) == 0 {
				goto l269
			}
			t644 := int32(load32(m.memory[int64(uint32(v8))+124:]))
			v9 = t644
			t645 := int64(load64(m.memory[int64(uint32(v8))+48:]))
			store16(m.memory[int64(uint32(v7))+16:], uint16(t645))
			t646 := m.fn907(v9, v7+i32(16), i32(2))
			store32(m.memory[int64(uint32(v8))+124:], uint32(t646))
			goto l269
		}
	l269:
		m.memory[int64(uint32(v8))+64] = byte(i32(0))
		store64(m.memory[int64(uint32(v8))+48:], uint64(i64(0)))
		v9 = i32(2)
		goto l18
	case 2:
		{
			t604 := int32(m.memory[int64(uint32(v8))+64])
			v9 = t604
			if uint32(v9) > uint32(i32(31)) {
				goto l261
			}
			v3 = i32(0)
			v20 = i32(2)
			t605 := int32(load32(m.memory[int64(uint32(v8))+56:]))
			v5 = t605
			t606 := int32(load32(m.memory[int64(uint32(v8))+60:]))
			t607 := v5
			v29 = t606
			if t607 == v29 {
				goto l30
			}
			t608 := v8
			v27 = v5 + i32(1)
			store32(m.memory[int64(uint32(t608))+56:], uint32(v27))
			t609 := v8
			v25 = int64(uint32(v9))
			v30 = v25 + i64(8)
			m.memory[int64(uint32(t609))+64] = byte(v30)
			t610 := int64(m.memory[uint32(v5)])
			t611 := int64(load64(m.memory[int64(uint32(v8))+48:]))
			t612 := v8
			v26 = i64_shl(t610, v25) | t611
			store64(m.memory[int64(uint32(t612))+48:], uint64(v26))
			if uint32(v9) > uint32(i32(23)) {
				goto l261
			}
			if v27 == v29 {
				goto l30
			}
			v20 = i32(2)
			t613 := v8
			v27 = v5 + i32(2)
			store32(m.memory[int64(uint32(t613))+56:], uint32(v27))
			t614 := v8
			v31 = v25 + i64(16)
			m.memory[int64(uint32(t614))+64] = byte(v31)
			t615 := int64(m.memory[int64(uint32(v5))+1])
			t616 := v8
			v30 = i64_shl(t615, v30) | v26
			store64(m.memory[int64(uint32(t616))+48:], uint64(v30))
			if uint32(v9) > uint32(i32(15)) {
				goto l261
			}
			if v27 == v29 {
				goto l30
			}
			t617 := v8
			v27 = v5 + i32(3)
			store32(m.memory[int64(uint32(t617))+56:], uint32(v27))
			t618 := v8
			v25 = v25 + i64(24)
			m.memory[int64(uint32(t618))+64] = byte(v25)
			t619 := int64(m.memory[int64(uint32(v5))+2])
			t620 := v8
			v30 = i64_shl(t619, v31) | v30
			store64(m.memory[int64(uint32(t620))+48:], uint64(v30))
			if uint32(v9) > uint32(i32(7)) {
				goto l261
			}
			if v27 == v29 {
				goto l30
			}
			store32(m.memory[int64(uint32(v8))+56:], uint32(v5+i32(4)))
			m.memory[int64(uint32(v8))+64] = byte(v9 | i32(32))
			t621 := int64(m.memory[int64(uint32(v5))+3])
			store64(m.memory[int64(uint32(v8))+48:], uint64(i64_shl(t621, v25)|v30))
		}
	l261:
		{
			t622 := int32(load32(m.memory[int64(uint32(v8))+140:]))
			v9 = t622
			if v9 == 0 {
				goto l262
			}
			t623 := int64(load64(m.memory[int64(uint32(v8))+48:]))
			store32(m.memory[int64(uint32(v9))+4:], uint32(t623))
		}
	l262:
		{
			t624 := int32(m.memory[int64(uint32(v8))+121])
			if t624&i32(2) == 0 {
				goto l263
			}
			t625 := int32(m.memory[int64(uint32(v8))+3])
			if t625&i32(4) == 0 {
				goto l263
			}
			t626 := int64(load64(m.memory[int64(uint32(v8))+48:]))
			store32(m.memory[int64(uint32(v7))+16:], uint32(t626))
			t627 := int32(load32(m.memory[int64(uint32(v8))+124:]))
			t628 := m.fn907(t627, v7+i32(16), i32(4))
			store32(m.memory[int64(uint32(v8))+124:], uint32(t628))
		}
	l263:
		m.memory[int64(uint32(v8))+64] = byte(i32(0))
		store64(m.memory[int64(uint32(v8))+48:], uint64(i64(0)))
		v9 = i32(3)
		goto l18
	case 3:
		{
			t585 := int32(m.memory[int64(uint32(v8))+64])
			v9 = t585
			if uint32(v9) > uint32(i32(15)) {
				goto l258
			}
			v3 = i32(0)
			v20 = i32(3)
			t586 := int32(load32(m.memory[int64(uint32(v8))+56:]))
			v5 = t586
			t587 := int32(load32(m.memory[int64(uint32(v8))+60:]))
			t588 := v5
			v29 = t587
			if t588 == v29 {
				goto l30
			}
			t589 := v8
			v27 = v5 + i32(1)
			store32(m.memory[int64(uint32(t589))+56:], uint32(v27))
			t590 := v8
			v25 = int64(uint32(v9))
			v30 = v25 + i64(8)
			m.memory[int64(uint32(t590))+64] = byte(v30)
			t591 := int64(m.memory[uint32(v5)])
			t592 := int64(load64(m.memory[int64(uint32(v8))+48:]))
			t593 := v8
			v25 = i64_shl(t591, v25) | t592
			store64(m.memory[int64(uint32(t593))+48:], uint64(v25))
			if uint32(v9) > uint32(i32(7)) {
				goto l258
			}
			if v27 == v29 {
				goto l30
			}
			store32(m.memory[int64(uint32(v8))+56:], uint32(v5+i32(2)))
			m.memory[int64(uint32(v8))+64] = byte(v9 | i32(16))
			t594 := int64(m.memory[int64(uint32(v5))+1])
			store64(m.memory[int64(uint32(v8))+48:], uint64(i64_shl(t594, v30)|v25))
		}
	l258:
		{
			t595 := int32(load32(m.memory[int64(uint32(v8))+140:]))
			v9 = t595
			if v9 == 0 {
				goto l259
			}
			t596 := int32(m.memory[int64(uint32(v8))+48])
			store32(m.memory[int64(uint32(v9))+8:], uint32(t596))
			t597 := int32(load32(m.memory[int64(uint32(v8))+140:]))
			t598 := int64(load64(m.memory[int64(uint32(v8))+48:]))
			store32(m.memory[int64(uint32(t597))+12:], uint32(int64(uint64(t598)>>8)))
		}
	l259:
		{
			t599 := int32(m.memory[int64(uint32(v8))+121])
			if t599&i32(2) == 0 {
				goto l260
			}
			t600 := int32(m.memory[int64(uint32(v8))+3])
			if t600&i32(4) == 0 {
				goto l260
			}
			t601 := int64(load64(m.memory[int64(uint32(v8))+48:]))
			store16(m.memory[int64(uint32(v7))+16:], uint16(t601))
			t602 := int32(load32(m.memory[int64(uint32(v8))+124:]))
			t603 := m.fn907(t602, v7+i32(16), i32(2))
			store32(m.memory[int64(uint32(v8))+124:], uint32(t603))
		}
	l260:
		m.memory[int64(uint32(v8))+64] = byte(i32(0))
		store64(m.memory[int64(uint32(v8))+48:], uint64(i64(0)))
		v9 = i32(4)
		goto l18
	case 4:
		{
			{
				t257 := int32(load32(m.memory[int64(uint32(v8))+120:]))
				v9 = t257
				if v9&i32(1024) == 0 {
					v9 = i32(5)
					t273 := int32(load32(m.memory[int64(uint32(v8))+140:]))
					v5 = t273
					if v5 == 0 {
						goto l18
					}
					store32(m.memory[int64(uint32(v5))+16:], uint32(i32(0)))
					goto l18
				}
				t258 := int64(load64(m.memory[int64(uint32(v8))+48:]))
				v25 = t258
				{
					t259 := int32(m.memory[int64(uint32(v8))+64])
					v5 = t259
					if uint32(v5) > uint32(i32(15)) {
						goto l154
					}
					v3 = i32(0)
					v20 = i32(4)
					t260 := int32(load32(m.memory[int64(uint32(v8))+56:]))
					v29 = t260
					t261 := int32(load32(m.memory[int64(uint32(v8))+60:]))
					t262 := v29
					v27 = t261
					if t262 == v27 {
						goto l30
					}
					t263 := v8
					v28 = v29 + i32(1)
					store32(m.memory[int64(uint32(t263))+56:], uint32(v28))
					t264 := v8
					v30 = int64(uint32(v5))
					v26 = v30 + i64(8)
					m.memory[int64(uint32(t264))+64] = byte(v26)
					t265 := int64(m.memory[uint32(v29)])
					t266 := v8
					v25 = i64_shl(t265, v30) | v25
					store64(m.memory[int64(uint32(t266))+48:], uint64(v25))
					if uint32(v5) > uint32(i32(7)) {
						goto l154
					}
					if v28 == v27 {
						goto l30
					}
					store32(m.memory[int64(uint32(v8))+56:], uint32(v29+i32(2)))
					m.memory[int64(uint32(v8))+64] = byte(v5 | i32(16))
					t267 := int64(m.memory[int64(uint32(v29))+1])
					t268 := v8
					v25 = i64_shl(t267, v26) | v25
					store64(m.memory[int64(uint32(t268))+48:], uint64(v25))
				}
			l154:
				t269 := v8
				v5 = int32(v25)
				store32(m.memory[int64(uint32(t269))+88:], uint32(v5))
				{
					t270 := int32(load32(m.memory[int64(uint32(v8))+140:]))
					v3 = t270
					if v3 == 0 {
						goto l155
					}
					store32(m.memory[int64(uint32(v3))+20:], uint32(v5))
					t271 := int32(load32(m.memory[int64(uint32(v8))+120:]))
					v9 = t271
				}
			l155:
				if v9&i32(512) == 0 {
					goto l156
				}
				t272 := int32(m.memory[int64(uint32(v8))+3])
				if t272&i32(4) != 0 {
					goto l157
				}
				goto l156
			}
		l157:
			t274 := int64(load64(m.memory[int64(uint32(v8))+48:]))
			store16(m.memory[int64(uint32(v7))+16:], uint16(t274))
			t275 := int32(load32(m.memory[int64(uint32(v8))+124:]))
			t276 := m.fn907(t275, v7+i32(16), i32(2))
			store32(m.memory[int64(uint32(v8))+124:], uint32(t276))
		}
	l156:
		m.memory[int64(uint32(v8))+64] = byte(i32(0))
		store64(m.memory[int64(uint32(v8))+48:], uint64(i64(0)))
		v9 = i32(5)
		goto l18
	case 5:
		{
			t226 := int32(load32(m.memory[int64(uint32(v8))+120:]))
			v3 = t226
			if v3&i32(1024) == 0 {
				goto l147
			}
			{
				t227 := int32(load32(m.memory[int64(uint32(v8))+60:]))
				t228 := int32(load32(m.memory[int64(uint32(v8))+56:]))
				v29 = t228
				v5 = t227 - v29
				t229 := int32(load32(m.memory[int64(uint32(v8))+88:]))
				t230 := v5
				v9 = t229
				p231 := v9
				if uint32(v5) < uint32(v9) {
					p231 = t230
				}
				v5 = p231
				if v5 == 0 {
					goto l148
				}
				{
					t232 := int32(load32(m.memory[int64(uint32(v8))+140:]))
					v20 = t232
					if v20 == 0 {
						goto l149
					}
					t233 := int32(load32(m.memory[int64(uint32(v20))+16:]))
					v27 = t233
					if v27 == 0 {
						goto l149
					}
					{
						t234 := int32(load32(m.memory[int64(uint32(v20))+24:]))
						t235 := v5
						v3 = t234
						t236 := int32(load32(m.memory[int64(uint32(v20))+20:]))
						t237 := v3
						v9 = t236 - v9
						v20 = t237 - v9
						p238 := v20
						if uint32(v20) > uint32(v3) {
							p238 = i32(0)
						}
						v20 = p238
						p239 := v20
						if uint32(v5) < uint32(v20) {
							p239 = t235
						}
						v20 = p239
						if v20 == 0 {
							goto l150
						}
						t241 := v27
						p240 := v9
						if uint32(v3) < uint32(v9) {
							p240 = v3
						}
						memory_copy(m.memory, uint32(t241+p240), uint32(v29), uint32(v20))
					}
				l150:
					t242 := int32(load32(m.memory[int64(uint32(v8))+120:]))
					v3 = t242
				}
			l149:
				{
					if v3&i32(512) == 0 {
						goto l151
					}
					t243 := int32(m.memory[int64(uint32(v8))+3])
					if t243&i32(4) == 0 {
						goto l151
					}
					{
						t244 := int32(load32(m.memory[int64(uint32(v8))+60:]))
						t245 := int32(load32(m.memory[int64(uint32(v8))+56:]))
						t246 := v5
						v9 = t245
						v3 = t244 - v9
						if uint32(t246) > uint32(v3) {
							m.fn121(i32(0), v5, v3, i32(1284604))
							panic("unreachable")
						}
						t247 := int32(load32(m.memory[int64(uint32(v8))+124:]))
						t248 := m.fn907(t247, v9, v5)
						store32(m.memory[int64(uint32(v8))+124:], uint32(t248))
						goto l151
					}
				}
			l151:
				t249 := int32(load32(m.memory[int64(uint32(v8))+112:]))
				store32(m.memory[int64(uint32(v8))+112:], uint32(t249-v5))
				t250 := int32(load32(m.memory[int64(uint32(v8))+88:]))
				t251 := v8
				v9 = t250 - v5
				store32(m.memory[int64(uint32(t251))+88:], uint32(v9))
				t252 := int32(load32(m.memory[int64(uint32(v8))+60:]))
				t253 := v8
				v3 = t252
				t254 := int32(load32(m.memory[int64(uint32(v8))+56:]))
				t255 := v3
				v5 = t254 + v5
				p256 := v5
				if uint32(v3) < uint32(v5) {
					p256 = t255
				}
				store32(m.memory[int64(uint32(t253))+56:], uint32(p256))
			}
		l148:
			if v9 == 0 {
				goto l147
			}
			v3 = i32(0)
			v20 = i32(5)
			goto l30
		}
	l147:
		store32(m.memory[int64(uint32(v8))+88:], uint32(i32(0)))
		v9 = i32(6)
		goto l18
	case 6:
		{
			t196 := int32(load32(m.memory[int64(uint32(v8))+120:]))
			v27 = t196
			if v27&i32(2048) == 0 {
				t198 := int32(load32(m.memory[int64(uint32(v8))+140:]))
				v9 = t198
				if v9 == 0 {
					goto l136
				}
				store32(m.memory[int64(uint32(v9))+28:], uint32(i32(0)))
				goto l136
			}
			v20 = i32(6)
			t197 := int32(load32(m.memory[int64(uint32(v8))+112:]))
			if t197 != 0 {
				{
					t199 := int32(load32(m.memory[int64(uint32(v8))+60:]))
					v3 = t199
					t200 := int32(load32(m.memory[int64(uint32(v8))+56:]))
					t201 := v3
					v5 = t200
					if t201 != v5 {
						goto l137
					}
					v29 = i32(0)
					goto l138
				}
			l137:
				v29 = v3 - v5
				v9 = i32(0)
			l140:
				{
					t202 := int32(m.memory[uint32(v5+v9)])
					if t202 == 0 {
						goto l139
					}
					t203 := v5
					v9 = v9 + i32(1)
					if t203+v9 != v3 {
						goto l140
					}
					goto l138
				}
			l139:
				v29 = v9 + i32(1)
			l138:
				{
					t204 := int32(load32(m.memory[int64(uint32(v8))+140:]))
					v9 = t204
					if v9 == 0 {
						goto l141
					}
					t205 := int32(load32(m.memory[int64(uint32(v9))+28:]))
					v28 = t205
					if v28 == 0 {
						goto l141
					}
					{
						t206 := int32(load32(m.memory[int64(uint32(v9))+32:]))
						v3 = t206
						t207 := int32(load32(m.memory[int64(uint32(v8))+88:]))
						t208 := v3
						v9 = t207
						if uint32(t208) < uint32(v9) {
							m.fn140(i32(1284620), i32(18), i32(1284640))
							panic("unreachable")
						}
						{
							v3 = v3 - v9
							p209 := v29
							if uint32(v3) < uint32(v29) {
								p209 = v3
							}
							v3 = p209
							if v3 == 0 {
								goto l143
							}
							memory_copy(m.memory, uint32(v28+v9), uint32(v5), uint32(v3))
						}
					l143:
						t210 := int32(load32(m.memory[int64(uint32(v8))+88:]))
						store32(m.memory[int64(uint32(v8))+88:], uint32(t210+v3))
						t211 := int32(load32(m.memory[int64(uint32(v8))+120:]))
						v27 = t211
						goto l141
					}
				}
			l141:
				{
					if v27&i32(512) == 0 {
						goto l144
					}
					t212 := int32(m.memory[int64(uint32(v8))+3])
					if t212&i32(4) == 0 {
						goto l144
					}
					t213 := int32(load32(m.memory[int64(uint32(v8))+124:]))
					t214 := m.fn907(t213, v5, v29)
					store32(m.memory[int64(uint32(v8))+124:], uint32(t214))
				}
			l144:
				{
					{
						if v29 != 0 {
							goto l145
						}
						t215 := int32(load32(m.memory[int64(uint32(v8))+60:]))
						t216 := v8
						v9 = t215
						t217 := int32(load32(m.memory[int64(uint32(v8))+56:]))
						t218 := v9
						v5 = t217
						p219 := v5
						if uint32(v9) < uint32(v5) {
							p219 = t218
						}
						v5 = p219
						store32(m.memory[int64(uint32(t216))+56:], uint32(v5))
						goto l146
					}
				l145:
					t220 := int32(m.memory[uint32(v5+v29+i32(-1))])
					v3 = t220
					t221 := int32(load32(m.memory[int64(uint32(v8))+60:]))
					t222 := v8
					v9 = t221
					t223 := int32(load32(m.memory[int64(uint32(v8))+56:]))
					t224 := v9
					v5 = t223 + v29
					p225 := v5
					if uint32(v9) < uint32(v5) {
						p225 = t224
					}
					v5 = p225
					store32(m.memory[int64(uint32(t222))+56:], uint32(v5))
					if v3 == 0 {
						goto l136
					}
				}
			l146:
				if v9 != v5 {
					goto l136
				}
				v3 = i32(0)
				goto l30
			}
			v3 = i32(0)
			goto l30
		}
	l136:
		store32(m.memory[int64(uint32(v8))+88:], uint32(i32(0)))
		v9 = i32(7)
		goto l18
	case 7:
		t166 := int32(load32(m.memory[int64(uint32(v8))+120:]))
		v27 = t166
		if v27&i32(4096) == 0 {
			v9 = i32(8)
			t168 := int32(load32(m.memory[int64(uint32(v8))+140:]))
			v5 = t168
			if v5 == 0 {
				goto l18
			}
			store32(m.memory[int64(uint32(v5))+36:], uint32(i32(0)))
			goto l18
		}
		v20 = i32(7)
		t167 := int32(load32(m.memory[int64(uint32(v8))+112:]))
		if t167 != 0 {
			{
				t169 := int32(load32(m.memory[int64(uint32(v8))+60:]))
				v3 = t169
				t170 := int32(load32(m.memory[int64(uint32(v8))+56:]))
				t171 := v3
				v5 = t170
				if t171 != v5 {
					goto l124
				}
				v29 = i32(0)
				goto l125
			}
		l124:
			v29 = v3 - v5
			v9 = i32(0)
		l127:
			{
				t172 := int32(m.memory[uint32(v5+v9)])
				if t172 == 0 {
					goto l126
				}
				t173 := v5
				v9 = v9 + i32(1)
				if t173+v9 != v3 {
					goto l127
				}
				goto l125
			}
		l126:
			v29 = v9 + i32(1)
		l125:
			{
				t174 := int32(load32(m.memory[int64(uint32(v8))+140:]))
				v9 = t174
				if v9 == 0 {
					goto l128
				}
				t175 := int32(load32(m.memory[int64(uint32(v9))+36:]))
				v28 = t175
				if v28 == 0 {
					goto l128
				}
				{
					t176 := int32(load32(m.memory[int64(uint32(v9))+40:]))
					v3 = t176
					t177 := int32(load32(m.memory[int64(uint32(v8))+88:]))
					t178 := v3
					v9 = t177
					if uint32(t178) < uint32(v9) {
						m.fn140(i32(1284656), i32(18), i32(1284676))
						panic("unreachable")
					}
					{
						v3 = v3 - v9
						p179 := v29
						if uint32(v3) < uint32(v29) {
							p179 = v3
						}
						v3 = p179
						if v3 == 0 {
							goto l130
						}
						memory_copy(m.memory, uint32(v28+v9), uint32(v5), uint32(v3))
					}
				l130:
					t180 := int32(load32(m.memory[int64(uint32(v8))+88:]))
					store32(m.memory[int64(uint32(v8))+88:], uint32(t180+v3))
					t181 := int32(load32(m.memory[int64(uint32(v8))+120:]))
					v27 = t181
					goto l128
				}
			}
		l128:
			{
				if v27&i32(512) == 0 {
					goto l131
				}
				t182 := int32(m.memory[int64(uint32(v8))+3])
				if t182&i32(4) == 0 {
					goto l131
				}
				t183 := int32(load32(m.memory[int64(uint32(v8))+124:]))
				t184 := m.fn907(t183, v5, v29)
				store32(m.memory[int64(uint32(v8))+124:], uint32(t184))
			}
		l131:
			{
				{
					if v29 != 0 {
						goto l132
					}
					t185 := int32(load32(m.memory[int64(uint32(v8))+60:]))
					t186 := v8
					v5 = t185
					t187 := int32(load32(m.memory[int64(uint32(v8))+56:]))
					t188 := v5
					v9 = t187
					p189 := v9
					if uint32(v5) < uint32(v9) {
						p189 = t188
					}
					v3 = p189
					store32(m.memory[int64(uint32(t186))+56:], uint32(v3))
					goto l133
				}
			l132:
				t190 := int32(m.memory[uint32(v5+v29+i32(-1))])
				v27 = t190
				t191 := int32(load32(m.memory[int64(uint32(v8))+60:]))
				t192 := v8
				v5 = t191
				t193 := int32(load32(m.memory[int64(uint32(v8))+56:]))
				t194 := v5
				v9 = t193 + v29
				p195 := v9
				if uint32(v5) < uint32(v9) {
					p195 = t194
				}
				v3 = p195
				store32(m.memory[int64(uint32(t192))+56:], uint32(v3))
				v9 = i32(8)
				if v27 == 0 {
					goto l18
				}
			}
		l133:
			v9 = i32(8)
			if v5 != v3 {
				goto l18
			}
			v3 = i32(0)
			goto l30
		}
		v3 = i32(0)
		goto l30
	case 8:
		{
			t148 := int32(load32(m.memory[int64(uint32(v8))+120:]))
			v5 = t148
			if v5&i32(512) == 0 {
				goto l117
			}
			{
				t149 := int32(m.memory[int64(uint32(v8))+64])
				v9 = t149
				if uint32(v9) > uint32(i32(15)) {
					goto l118
				}
				v3 = i32(0)
				v20 = i32(8)
				t150 := int32(load32(m.memory[int64(uint32(v8))+56:]))
				v29 = t150
				t151 := int32(load32(m.memory[int64(uint32(v8))+60:]))
				t152 := v29
				v27 = t151
				if t152 == v27 {
					goto l30
				}
				t153 := v8
				v28 = v29 + i32(1)
				store32(m.memory[int64(uint32(t153))+56:], uint32(v28))
				t154 := v8
				v25 = int64(uint32(v9))
				v30 = v25 + i64(8)
				m.memory[int64(uint32(t154))+64] = byte(v30)
				t155 := int64(m.memory[uint32(v29)])
				t156 := int64(load64(m.memory[int64(uint32(v8))+48:]))
				t157 := v8
				v25 = i64_shl(t155, v25) | t156
				store64(m.memory[int64(uint32(t157))+48:], uint64(v25))
				if uint32(v9) > uint32(i32(7)) {
					goto l118
				}
				if v28 == v27 {
					goto l30
				}
				store32(m.memory[int64(uint32(v8))+56:], uint32(v29+i32(2)))
				m.memory[int64(uint32(v8))+64] = byte(v9 | i32(16))
				t158 := int64(m.memory[int64(uint32(v29))+1])
				store64(m.memory[int64(uint32(v8))+48:], uint64(i64_shl(t158, v30)|v25))
			}
		l118:
			{
				t159 := int32(m.memory[int64(uint32(v8))+3])
				if t159&i32(4) == 0 {
					goto l119
				}
				t160 := int32(load16(m.memory[int64(uint32(v8))+124:]))
				t161 := int32(load32(m.memory[int64(uint32(v8))+48:]))
				if t160 != t161 {
					store32(m.memory[int64(uint32(v8))+136:], uint32(i32(20)))
					store32(m.memory[int64(uint32(v8))+132:], uint32(i32(1065455)))
					goto l57
				}
			}
		l119:
			m.memory[int64(uint32(v8))+64] = byte(i32(0))
			store64(m.memory[int64(uint32(v8))+48:], uint64(i64(0)))
		}
	l117:
		{
			t162 := int32(load32(m.memory[int64(uint32(v8))+140:]))
			v9 = t162
			if v9 == 0 {
				goto l121
			}
			store32(m.memory[int64(uint32(v9))+44:], uint32(int32(uint32(v5)>>9)&i32(1)))
			t163 := int32(load32(m.memory[int64(uint32(v8))+140:]))
			store32(m.memory[int64(uint32(t163))+48:], uint32(i32(1)))
			t164 := int32(load32(m.memory[int64(uint32(v8))+120:]))
			v5 = t164
		}
	l121:
		v9 = i32(12)
		t165 := int32(m.memory[int64(uint32(v8))+3])
		if t165&i32(4) == 0 {
			goto l18
		}
		if v5 == 0 {
			goto l18
		}
		store64(m.memory[int64(uint32(v8))+124:], uint64(i64(0)))
		goto l18
	case 11:
		{
			t563 := int32(m.memory[int64(uint32(v8))+3])
			v9 = t563
			if v9 == 0 {
				goto l254
			}
			t564 := int32(load32(m.memory[int64(uint32(v8))+120:]))
			if t564 == 0 {
				goto l254
			}
			{
				t565 := int32(m.memory[int64(uint32(v8))+64])
				v5 = t565
				if uint32(v5) > uint32(i32(31)) {
					goto l255
				}
				v3 = i32(0)
				v20 = i32(11)
				t566 := int32(load32(m.memory[int64(uint32(v8))+56:]))
				v15 = t566
				t567 := int32(load32(m.memory[int64(uint32(v8))+60:]))
				t568 := v15
				v29 = t567
				if t568 == v29 {
					goto l30
				}
				t569 := v8
				v27 = v15 + i32(1)
				store32(m.memory[int64(uint32(t569))+56:], uint32(v27))
				t570 := v8
				v25 = int64(uint32(v5))
				v30 = v25 + i64(8)
				m.memory[int64(uint32(t570))+64] = byte(v30)
				t571 := int64(m.memory[uint32(v15)])
				t572 := int64(load64(m.memory[int64(uint32(v8))+48:]))
				t573 := v8
				v26 = i64_shl(t571, v25) | t572
				store64(m.memory[int64(uint32(t573))+48:], uint64(v26))
				if uint32(v5) > uint32(i32(23)) {
					goto l255
				}
				if v27 == v29 {
					goto l30
				}
				t574 := v8
				v27 = v15 + i32(2)
				store32(m.memory[int64(uint32(t574))+56:], uint32(v27))
				t575 := v8
				v31 = v25 + i64(16)
				m.memory[int64(uint32(t575))+64] = byte(v31)
				t576 := int64(m.memory[int64(uint32(v15))+1])
				t577 := v8
				v30 = i64_shl(t576, v30) | v26
				store64(m.memory[int64(uint32(t577))+48:], uint64(v30))
				if uint32(v5) > uint32(i32(15)) {
					goto l255
				}
				if v27 == v29 {
					goto l30
				}
				t578 := v8
				v27 = v15 + i32(3)
				store32(m.memory[int64(uint32(t578))+56:], uint32(v27))
				t579 := v8
				v25 = v25 + i64(24)
				m.memory[int64(uint32(t579))+64] = byte(v25)
				t580 := int64(m.memory[int64(uint32(v15))+2])
				t581 := v8
				v30 = i64_shl(t580, v31) | v30
				store64(m.memory[int64(uint32(t581))+48:], uint64(v30))
				if uint32(v5) > uint32(i32(7)) {
					goto l255
				}
				if v27 == v29 {
					goto l30
				}
				store32(m.memory[int64(uint32(v8))+56:], uint32(v15+i32(4)))
				m.memory[int64(uint32(v8))+64] = byte(v5 | i32(32))
				t582 := int64(m.memory[int64(uint32(v15))+3])
				store64(m.memory[int64(uint32(v8))+48:], uint64(i64_shl(t582, v25)|v30))
			}
		l255:
			{
				if v9&i32(4) == 0 {
					goto l256
				}
				t583 := int32(load32(m.memory[int64(uint32(v8))+84:]))
				t584 := int32(load32(m.memory[int64(uint32(v8))+48:]))
				if t583 != t584 {
					store32(m.memory[int64(uint32(v8))+136:], uint32(i32(23)))
					store32(m.memory[int64(uint32(v8))+132:], uint32(i32(1065381)))
					goto l57
				}
			}
		l256:
			m.memory[int64(uint32(v8))+64] = byte(i32(0))
			store64(m.memory[int64(uint32(v8))+48:], uint64(i64(0)))
			goto l254
		}
	l254:
		v20 = i32(29)
		v3 = i32(1)
		goto l30
	case 12:
		v9 = i32(13)
		t562 := int32(m.memory[int64(uint32(v8))+4])
		if uint32(t562) <= uint32(i32(4)) {
			goto l18
		}
		v3 = i32(0)
		v20 = i32(12)
		goto l30
	case 13:
		t137 := int32(m.memory[int64(uint32(v8))+1])
		v5 = t137
		if v5&i32(1) == 0 {
			t141 := int64(load64(m.memory[int64(uint32(v8))+48:]))
			v25 = t141
			{
				t142 := int32(m.memory[int64(uint32(v8))+64])
				v9 = t142
				if uint32(v9) > uint32(i32(2)) {
					goto l111
				}
				{
					t143 := int32(load32(m.memory[int64(uint32(v8))+56:]))
					v3 = t143
					t144 := int32(load32(m.memory[int64(uint32(v8))+60:]))
					if v3 != t144 {
						goto l112
					}
					v3 = i32(0)
					v20 = i32(13)
					goto l30
				}
			l112:
				v30 = int64(uint32(v9))
				store32(m.memory[int64(uint32(v8))+56:], uint32(v3+i32(1)))
				v9 = v9 | i32(8)
				t145 := int64(m.memory[uint32(v3)])
				v25 = i64_shl(t145, v30) | v25
			}
		l111:
			m.memory[int64(uint32(v8))+64] = byte(v9 + i32(-1))
			t146 := v8
			v30 = int64(uint64(v25) >> 1)
			store64(m.memory[int64(uint32(t146))+48:], uint64(v30))
			m.memory[int64(uint32(v8))+1] = byte(int32(v25)&i32(1) | v5)
			switch int32(v30) & i32(3) {
			default:
				m.memory[int64(uint32(v8))+64] = byte(v9 + i32(-3))
				store64(m.memory[int64(uint32(v8))+48:], uint64(int64(uint64(v25)>>3)))
				v9 = i32(14)
				goto l18
			case 1:
				v3 = i32(0)
				m.memory[int64(uint32(v8))+160] = byte(i32(0))
				store32(m.memory[int64(uint32(v8))+156:], uint32(i32(5)))
				m.memory[int64(uint32(v8))+152] = byte(i32(0))
				store32(m.memory[int64(uint32(v8))+148:], uint32(i32(9)))
				m.memory[int64(uint32(v8))+64] = byte(v9 + i32(-3))
				store64(m.memory[int64(uint32(v8))+48:], uint64(int64(uint64(v25)>>3)))
				v9 = i32(17)
				v20 = i32(17)
				t147 := int32(m.memory[int64(uint32(v8))+4])
				if t147 != i32(6) {
					goto l18
				}
				goto l30
			case 2:
				m.memory[int64(uint32(v8))+64] = byte(v9 + i32(-3))
				store64(m.memory[int64(uint32(v8))+48:], uint64(int64(uint64(v25)>>3)))
				v9 = i32(24)
				goto l18
			case 3:
				store32(m.memory[int64(uint32(v8))+136:], uint32(i32(19)))
				store32(m.memory[int64(uint32(v8))+132:], uint32(i32(1065720)))
				v3 = i32(-3)
				m.memory[int64(uint32(v8))+64] = byte(v9 + i32(-3))
				store64(m.memory[int64(uint32(v8))+48:], uint64(int64(uint64(v25)>>3)))
				v20 = i32(30)
				goto l30
			}
		}
		t138 := int32(m.memory[int64(uint32(v8))+64])
		t139 := v8
		v9 = t138
		m.memory[int64(uint32(t139))+64] = byte(v9 & i32(248))
		t140 := int64(load64(m.memory[int64(uint32(v8))+48:]))
		store64(m.memory[int64(uint32(v8))+48:], uint64(i64_shr_u(t140, int64(uint32(v9&i32(7))))))
		v9 = i32(16)
		goto l18
	case 14:
		t540 := int32(m.memory[int64(uint32(v8))+64])
		t541 := v8
		v9 = t540
		m.memory[int64(uint32(t541))+64] = byte(v9 & i32(248))
		t542 := int64(load64(m.memory[int64(uint32(v8))+48:]))
		t543 := v8
		v25 = i64_shr_u(t542, int64(uint32(v9&i32(7))))
		store64(m.memory[int64(uint32(t543))+48:], uint64(v25))
		{
			if uint32(v9) > uint32(i32(31)) {
				goto l252
			}
			v3 = i32(0)
			v20 = i32(14)
			t544 := int32(load32(m.memory[int64(uint32(v8))+56:]))
			v5 = t544
			t545 := int32(load32(m.memory[int64(uint32(v8))+60:]))
			t546 := v5
			v29 = t545
			if t546 == v29 {
				goto l30
			}
			t547 := v8
			v27 = v5 + i32(1)
			store32(m.memory[int64(uint32(t547))+56:], uint32(v27))
			t548 := v8
			v30 = int64(uint32(v9)) & i64(24)
			v26 = v30 + i64(8)
			m.memory[int64(uint32(t548))+64] = byte(v26)
			t549 := int64(m.memory[uint32(v5)])
			t550 := v8
			v25 = i64_shl(t549, v30) | v25
			store64(m.memory[int64(uint32(t550))+48:], uint64(v25))
			v9 = v9 & i32(24)
			if v9 == i32(24) {
				goto l252
			}
			if v27 == v29 {
				goto l30
			}
			t551 := v8
			v27 = v5 + i32(2)
			store32(m.memory[int64(uint32(t551))+56:], uint32(v27))
			t552 := v8
			v31 = v30 + i64(16)
			m.memory[int64(uint32(t552))+64] = byte(v31)
			t553 := int64(m.memory[int64(uint32(v5))+1])
			t554 := v8
			v25 = i64_shl(t553, v26) | v25
			store64(m.memory[int64(uint32(t554))+48:], uint64(v25))
			if uint32(v9) > uint32(i32(15)) {
				goto l252
			}
			if v27 == v29 {
				goto l30
			}
			t555 := v8
			v27 = v5 + i32(3)
			store32(m.memory[int64(uint32(t555))+56:], uint32(v27))
			t556 := v8
			v30 = v30 + i64(24)
			m.memory[int64(uint32(t556))+64] = byte(v30)
			t557 := int64(m.memory[int64(uint32(v5))+2])
			t558 := v8
			v25 = i64_shl(t557, v31) | v25
			store64(m.memory[int64(uint32(t558))+48:], uint64(v25))
			if v9 != 0 {
				goto l252
			}
			if v27 == v29 {
				goto l30
			}
			m.memory[int64(uint32(v8))+64] = byte(i32(32))
			store32(m.memory[int64(uint32(v8))+56:], uint32(v5+i32(4)))
			t559 := int64(m.memory[int64(uint32(v5))+3])
			t560 := v8
			v25 = i64_shl(t559, v30) | v25
			store64(m.memory[int64(uint32(t560))+48:], uint64(v25))
		}
	l252:
		{
			if (int64(uint64(v25)>>16)^v25)&i64(0xffff) != i64(0xffff) {
				store32(m.memory[int64(uint32(v8))+136:], uint32(i32(29)))
				store32(m.memory[int64(uint32(v8))+132:], uint32(i32(1064674)))
				goto l57
			}
			v3 = i32(0)
			m.memory[int64(uint32(v8))+64] = byte(i32(0))
			store64(m.memory[int64(uint32(v8))+48:], uint64(i64(0)))
			store32(m.memory[int64(uint32(v8))+88:], uint32(int32(v25)&i32(0xffff)))
			v9 = i32(15)
			t561 := int32(m.memory[int64(uint32(v8))+4])
			if t561 != i32(6) {
				goto l18
			}
			v20 = i32(14)
			goto l30
		}
	case 15:
		v9 = i32(12)
		t647 := int32(load32(m.memory[int64(uint32(v8))+88:]))
		v20 = t647
		if v20 == 0 {
			goto l18
		}
		t648 := int32(load32(m.memory[int64(uint32(v8))+56:]))
		v27 = t648
		t649 := int32(load32(m.memory[int64(uint32(v8))+60:]))
		v29 = t649
		t650 := int32(load32(m.memory[int64(uint32(v8))+80:]))
		v3 = t650
		t651 := int32(load32(m.memory[int64(uint32(v8))+72:]))
		v21 = t651
		t652 := int32(load32(m.memory[int64(uint32(v8))+76:]))
		v28 = t652
	l274:
		{
			v5 = v29 - v27
			t653 := int32(load32(m.memory[int64(uint32(v8))+76:]))
			t654 := v5
			v29 = t653 - v3
			p655 := v20
			if uint32(v29) < uint32(v20) {
				p655 = v29
			}
			v20 = p655
			p656 := v20
			if uint32(v5) < uint32(v20) {
				p656 = t654
			}
			v5 = p656
			if v5 != 0 {
				if uint32(v28) < uint32(v3) {
					m.fn121(v3, v28, v28, i32(1284400))
					panic("unreachable")
				}
				t657 := v5
				v20 = v28 - v3
				if uint32(t657) > uint32(v20) {
					m.fn121(i32(0), v5, v20, i32(1284384))
					panic("unreachable")
				}
				if v5 == 0 {
					goto l273
				}
				memory_copy(m.memory, uint32(v21+v3), uint32(v27), uint32(v5))
			l273:
				t658 := v8
				v3 = v5 + v3
				store32(m.memory[int64(uint32(t658))+80:], uint32(v3))
				t659 := int32(load32(m.memory[int64(uint32(v8))+88:]))
				t660 := v8
				v20 = t659 - v5
				store32(m.memory[int64(uint32(t660))+88:], uint32(v20))
				t661 := int32(load32(m.memory[int64(uint32(v8))+60:]))
				t662 := v8
				v29 = t661
				t663 := int32(load32(m.memory[int64(uint32(v8))+56:]))
				t664 := v29
				v5 = t663 + v5
				p665 := v5
				if uint32(v29) < uint32(v5) {
					p665 = t664
				}
				v27 = p665
				store32(m.memory[int64(uint32(t662))+56:], uint32(v27))
				if v20 == 0 {
					goto l18
				}
				goto l274
			}
			v3 = i32(0)
			v20 = i32(15)
			goto l30
		}
	case 16:
		v9 = i32(11)
		t100 := int32(m.memory[int64(uint32(v8))+3])
		v5 = t100
		if v5 == 0 {
			goto l18
		}
		{
			t101 := int32(m.memory[int64(uint32(v8))+64])
			v29 = t101
			if uint32(v29) > uint32(i32(31)) {
				goto l102
			}
			v3 = i32(0)
			v20 = i32(16)
			t102 := int32(load32(m.memory[int64(uint32(v8))+56:]))
			v27 = t102
			t103 := int32(load32(m.memory[int64(uint32(v8))+60:]))
			t104 := v27
			v28 = t103
			if t104 == v28 {
				goto l30
			}
			t105 := v8
			v21 = v27 + i32(1)
			store32(m.memory[int64(uint32(t105))+56:], uint32(v21))
			t106 := v8
			v25 = int64(uint32(v29))
			v30 = v25 + i64(8)
			m.memory[int64(uint32(t106))+64] = byte(v30)
			t107 := int64(m.memory[uint32(v27)])
			t108 := int64(load64(m.memory[int64(uint32(v8))+48:]))
			t109 := v8
			v26 = i64_shl(t107, v25) | t108
			store64(m.memory[int64(uint32(t109))+48:], uint64(v26))
			if uint32(v29) > uint32(i32(23)) {
				goto l102
			}
			if v21 == v28 {
				goto l30
			}
			t110 := v8
			v21 = v27 + i32(2)
			store32(m.memory[int64(uint32(t110))+56:], uint32(v21))
			t111 := v8
			v31 = v25 + i64(16)
			m.memory[int64(uint32(t111))+64] = byte(v31)
			t112 := int64(m.memory[int64(uint32(v27))+1])
			t113 := v8
			v30 = i64_shl(t112, v30) | v26
			store64(m.memory[int64(uint32(t113))+48:], uint64(v30))
			if uint32(v29) > uint32(i32(15)) {
				goto l102
			}
			if v21 == v28 {
				goto l30
			}
			t114 := v8
			v21 = v27 + i32(3)
			store32(m.memory[int64(uint32(t114))+56:], uint32(v21))
			t115 := v8
			v25 = v25 + i64(24)
			m.memory[int64(uint32(t115))+64] = byte(v25)
			t116 := int64(m.memory[int64(uint32(v27))+2])
			t117 := v8
			v30 = i64_shl(t116, v31) | v30
			store64(m.memory[int64(uint32(t117))+48:], uint64(v30))
			if uint32(v29) > uint32(i32(7)) {
				goto l102
			}
			if v21 == v28 {
				goto l30
			}
			store32(m.memory[int64(uint32(v8))+56:], uint32(v27+i32(4)))
			m.memory[int64(uint32(v8))+64] = byte(v29 | i32(32))
			t118 := int64(m.memory[int64(uint32(v27))+3])
			store64(m.memory[int64(uint32(v8))+48:], uint64(i64_shl(t118, v25)|v30))
		}
	l102:
		t119 := int32(load32(m.memory[int64(uint32(v8))+84:]))
		t120 := int32(load32(m.memory[int64(uint32(v8))+80:]))
		t121 := v8
		v20 = t120
		store32(m.memory[int64(uint32(t121))+84:], uint32(t119+v20))
		t122 := int32(load32(m.memory[int64(uint32(v8))+120:]))
		v3 = t122
		{
			{
				if v5&i32(4) == 0 {
					goto l103
				}
				{
					if v3 != 0 {
						goto l104
					}
					t123 := int32(load32(m.memory[int64(uint32(v8))+124:]))
					t124 := int32(load32(m.memory[int64(uint32(v8))+72:]))
					t125 := m.fn906(t123, t124, v20)
					store32(m.memory[int64(uint32(v8))+124:], uint32(t125))
					t126 := int64(load64(m.memory[int64(uint32(v8))+48:]))
					v25 = t126
					goto l105
				}
			l104:
				t127 := int32(load32(m.memory[int64(uint32(v8))+128:]))
				t128 := int32(load32(m.memory[int64(uint32(v8))+72:]))
				t129 := m.fn907(t127, t128, v20)
				t130 := v8
				v5 = t129
				store32(m.memory[int64(uint32(t130))+124:], uint32(v5))
				store32(m.memory[int64(uint32(v8))+128:], uint32(v5))
				t131 := int32(load32(m.memory[int64(uint32(v8))+120:]))
				v3 = t131
			}
		l103:
			t132 := int64(load64(m.memory[int64(uint32(v8))+48:]))
			v25 = t132
			if v3 != 0 {
				goto l106
			}
		}
	l105:
		v5 = int32(v25)
		v5 = i32_rotr(v5&i32(0xff00ff), i32(8)) | i32_rotr(v5, i32(24))&i32(0xff00ff)
		goto l107
	l106:
		v5 = int32(v25)
	l107:
		t133 := int32(load32(m.memory[int64(uint32(v8))+76:]))
		t134 := int32(load32(m.memory[int64(uint32(v8))+80:]))
		store32(m.memory[int64(uint32(v8))+116:], uint32(t133-t134))
		{
			t135 := int32(m.memory[int64(uint32(v8))+3])
			if t135&i32(4) == 0 {
				goto l108
			}
			t136 := int32(load32(m.memory[int64(uint32(v8))+124:]))
			if v5 != t136 {
				store32(m.memory[int64(uint32(v8))+136:], uint32(i32(21)))
				store32(m.memory[int64(uint32(v8))+132:], uint32(i32(1065404)))
				goto l57
			}
		}
	l108:
		m.memory[int64(uint32(v8))+64] = byte(i32(0))
		store64(m.memory[int64(uint32(v8))+48:], uint64(i64(0)))
		goto l18
	case 17:
		goto l18
	case 18:
		m.memory[uint32(v8)] = byte(i32(18))
		{
			t387 := int32(load32(m.memory[int64(uint32(v8))+60:]))
			v3 = t387
			t388 := int32(load32(m.memory[int64(uint32(v8))+56:]))
			t389 := v3
			v24 = t388
			if uint32(t389-v24) < uint32(i32(15)) {
				goto l189
			}
			t390 := int32(load32(m.memory[int64(uint32(v8))+76:]))
			t391 := int32(load32(m.memory[int64(uint32(v8))+80:]))
			if uint32(t390-t391) <= uint32(i32(259)) {
				goto l189
			}
			m.fn910(v8)
			t392 := int32(m.memory[uint32(v8)])
			v9 = t392
			if v9 != i32(18) {
				goto l18
			}
			t393 := int32(load32(m.memory[int64(uint32(v8))+60:]))
			v3 = t393
			t394 := int32(load32(m.memory[int64(uint32(v8))+56:]))
			v24 = t394
		}
	l189:
		t395 := int32(load32(m.memory[int64(uint32(v19))+8:]))
		v9 = t395
		store32(m.memory[int64(uint32(v8))+80:], uint32(i32(0)))
		t396 := int64(load64(m.memory[uint32(v19):]))
		v25 = t396
		store64(m.memory[int64(uint32(v8))+72:], uint64(i64(1)))
		store32(m.memory[int64(uint32(v7))+24:], uint32(v9))
		store64(m.memory[int64(uint32(v7))+16:], uint64(v25))
		t397 := int32(load32(m.memory[uint32(v12):]))
		store32(m.memory[int64(uint32(v7))+8:], uint32(t397))
		t398 := int32(load32(m.memory[int64(uint32(v12))+3:]))
		store32(m.memory[int64(uint32(v7))+11:], uint32(t398))
		v28 = i32(512)
		v21 = i32(1277688)
		{
			t399 := int32(m.memory[int64(uint32(v8))+152])
			switch t399 {
			default:
				goto l190
			case 1:
				v28 = i32(1332)
				v21 = v16
				goto l190
			case 2:
				v28 = i32(1332)
				v21 = v17
				goto l190
			case 3:
				v28 = i32(592)
				v21 = v18
			}
		}
	l190:
		v29 = i32(32)
		v27 = i32(1279736)
		{
			t400 := int32(m.memory[int64(uint32(v8))+160])
			switch t400 {
			default:
				goto l194
			case 1:
				v29 = i32(1332)
				v27 = v16
				goto l194
			case 2:
				v29 = i32(1332)
				v27 = v17
				goto l194
			case 3:
				v29 = i32(592)
				v27 = v18
			}
		}
	l194:
		t401 := int32(m.memory[int64(uint32(v8))+64])
		v22 = t401
		t402 := int64(load64(m.memory[int64(uint32(v8))+48:]))
		v25 = t402
		{
		l220:
			v20 = i32(18)
		l212:
			{
				{
					{
						{
							{
								switch v20&i32(255) + i32(-18) {
								case 1:
									t418 := int32(load32(m.memory[int64(uint32(v7))+24:]))
									v9 = t418
									t419 := int32(load32(m.memory[int64(uint32(v7))+20:]))
									t420 := v9
									v5 = t419
									if t420 == v5 {
										v20 = i32(19)
										m.memory[uint32(v8)] = byte(i32(19))
										t483 := int64(load64(m.memory[int64(uint32(v7))+16:]))
										store64(m.memory[uint32(v19):], uint64(t483))
										t484 := int32(load32(m.memory[int64(uint32(v7))+24:]))
										store32(m.memory[int64(uint32(v19))+8:], uint32(t484))
										m.memory[int64(uint32(v8))+64] = byte(v22)
										store32(m.memory[int64(uint32(v8))+60:], uint32(v3))
										store32(m.memory[int64(uint32(v8))+56:], uint32(v24))
										store64(m.memory[int64(uint32(v8))+48:], uint64(v25))
										t485 := int32(load32(m.memory[int64(uint32(v7))+8:]))
										store32(m.memory[uint32(v12):], uint32(t485))
										t486 := int32(load32(m.memory[int64(uint32(v7))+11:]))
										store32(m.memory[int64(uint32(v12))+3:], uint32(t486))
										v3 = i32(0)
										goto l30
									}
									{
										if uint32(v9) >= uint32(v5) {
											m.fn33(v9, v5, i32(1279896))
											panic("unreachable")
										}
										t421 := int32(load32(m.memory[int64(uint32(v7))+16:]))
										t422 := int32(load32(m.memory[int64(uint32(v8))+88:]))
										m.memory[uint32(t421+v9)] = byte(t422)
										store32(m.memory[int64(uint32(v7))+24:], uint32(v9+i32(1)))
										goto l220
									}
								case 3:
									v5 = v24
									v9 = v22
									t429 := int64(load32(m.memory[int64(uint32(v8))+156:]))
									t430 := v29
									t431 := v25
									v31 = i64_shl(i64(-1), t429) ^ i64(-1)
									v20 = int32(t431 & v31)
									if uint32(t430) <= uint32(v20) {
										goto l223
									}
								l226:
									{
										v23 = v27 + v20<<2
										t432 := int32(m.memory[int64(uint32(v23))+3])
										v20 = t432
										if uint32(v20) <= uint32(v9&i32(255)) {
											goto l224
										}
										{
											if v5 == v3 {
												goto l225
											}
											v30 = int64(uint32(v9))
											t433 := int64(m.memory[uint32(v5)])
											v26 = t433
											v5 = v5 + i32(1)
											v9 = v9 + i32(8)
											t434 := v29
											v25 = i64_shl(v26, v30) | v25
											v20 = int32(v25 & v31)
											if uint32(t434) <= uint32(v20) {
												goto l223
											}
											goto l226
										}
									l225:
									}
									v20 = i32(21)
									m.memory[uint32(v8)] = byte(i32(21))
									t435 := int64(load64(m.memory[int64(uint32(v7))+16:]))
									store64(m.memory[uint32(v19):], uint64(t435))
									t436 := int32(load32(m.memory[int64(uint32(v7))+24:]))
									store32(m.memory[int64(uint32(v19))+8:], uint32(t436))
									store32(m.memory[int64(uint32(v8))+60:], uint32(v3))
									store32(m.memory[int64(uint32(v8))+56:], uint32(v3))
									store64(m.memory[int64(uint32(v8))+48:], uint64(v25))
									t437 := int32(load32(m.memory[int64(uint32(v7))+8:]))
									store32(m.memory[uint32(v12):], uint32(t437))
									t438 := int32(load32(m.memory[int64(uint32(v7))+11:]))
									store32(m.memory[int64(uint32(v12))+3:], uint32(t438))
									m.memory[int64(uint32(v8))+64] = byte(v22 + v3<<3 - v24<<3)
									v3 = i32(0)
									goto l30
								case 5:
									t445 := int32(load32(m.memory[int64(uint32(v7))+20:]))
									v5 = t445
									t446 := int32(load32(m.memory[int64(uint32(v7))+24:]))
									t447 := v5
									v9 = t446
									if t447 == v9 {
										v20 = i32(23)
										m.memory[uint32(v8)] = byte(i32(23))
										t495 := int64(load64(m.memory[int64(uint32(v7))+16:]))
										store64(m.memory[uint32(v19):], uint64(t495))
										t496 := int32(load32(m.memory[int64(uint32(v7))+24:]))
										store32(m.memory[int64(uint32(v19))+8:], uint32(t496))
										m.memory[int64(uint32(v8))+64] = byte(v22)
										store32(m.memory[int64(uint32(v8))+60:], uint32(v3))
										store32(m.memory[int64(uint32(v8))+56:], uint32(v24))
										store64(m.memory[int64(uint32(v8))+48:], uint64(v25))
										t497 := int32(load32(m.memory[int64(uint32(v7))+8:]))
										store32(m.memory[uint32(v12):], uint32(t497))
										t498 := int32(load32(m.memory[int64(uint32(v7))+11:]))
										store32(m.memory[int64(uint32(v12))+3:], uint32(t498))
										v3 = i32(0)
										goto l30
									}
									v5 = v5 - v9
									{
										t448 := int32(load32(m.memory[int64(uint32(v8))+92:]))
										v20 = t448
										if uint32(v20) > uint32(v9) {
											v9 = v20 - v9
											t454 := int32(load32(m.memory[int64(uint32(v8))+16:]))
											if uint32(v9) > uint32(t454) {
												t473 := int32(m.memory[int64(uint32(v8))+1])
												if t473&i32(4) == 0 {
													m.fn28(i32(1279992), i32(85), i32(1284588))
													panic("unreachable")
												}
												t474 := int32(load32(m.memory[int64(uint32(v7))+24:]))
												store32(m.memory[int64(uint32(v19))+8:], uint32(t474))
												t475 := int64(load64(m.memory[int64(uint32(v7))+16:]))
												store64(m.memory[uint32(v19):], uint64(t475))
												m.memory[int64(uint32(v8))+64] = byte(v22)
												store32(m.memory[int64(uint32(v8))+60:], uint32(v3))
												store32(m.memory[int64(uint32(v8))+56:], uint32(v24))
												store64(m.memory[int64(uint32(v8))+48:], uint64(v25))
												t476 := int32(load32(m.memory[int64(uint32(v7))+8:]))
												store32(m.memory[uint32(v12):], uint32(t476))
												t477 := int32(load32(m.memory[int64(uint32(v7))+11:]))
												store32(m.memory[int64(uint32(v12))+3:], uint32(t477))
												v20 = i32(30)
												store32(m.memory[int64(uint32(v8))+136:], uint32(i32(30)))
												store32(m.memory[int64(uint32(v8))+132:], uint32(i32(1065425)))
												goto l235
											}
											t455 := int32(load32(m.memory[int64(uint32(v8))+12:]))
											v23 = t455
											v20 = v23 + i32(-64)
											if uint32(v20) >= uint32(i32(-63)) {
												m.fn7(i32(1284212), i32(74), i32(1284288))
												panic("unreachable")
											}
											t457 := v7 + i32(16)
											t458 := v13
											p456 := v20
											if uint32(v20) > uint32(v23) {
												p456 = i32(0)
											}
											t459 := int32(load32(m.memory[int64(uint32(v8))+20:]))
											t460 := v9
											v20 = t459
											v23 = t460 - v20
											t461 := p456 - v23
											t462 := v20 - v9
											var p463 int32
											if uint32(v9) > uint32(v20) {
												p463 = 1
											}
											v20 = p463
											p464 := t462
											if v20 != 0 {
												p464 = t461
											}
											v32 = p464
											t465 := int32(load32(m.memory[int64(uint32(v8))+88:]))
											t466 := v32
											t467 := v32
											t468 := v5
											v33 = t465
											t470 := v33
											p469 := v9
											if v20 != 0 {
												p469 = v23
											}
											v9 = p469
											p471 := v9
											if uint32(v33) < uint32(v9) {
												p471 = t470
											}
											v9 = p471
											p472 := v9
											if uint32(v5) < uint32(v9) {
												p472 = t468
											}
											v9 = p472
											m.fn909(t457, t458, t466, t467+v9)
											goto l231
										}
										t449 := int32(load32(m.memory[int64(uint32(v8))+88:]))
										t450 := v7 + i32(16)
										t451 := v20
										t452 := v5
										v9 = t449
										p453 := v9
										if uint32(v5) < uint32(v9) {
											p453 = t452
										}
										v9 = p453
										m.fn908(t450, t451, v9)
										goto l231
									}
								default:
									{
										if uint32(v3-v24) < uint32(i32(15)) {
											goto l204
										}
										t403 := int32(load32(m.memory[int64(uint32(v7))+20:]))
										t404 := int32(load32(m.memory[int64(uint32(v7))+24:]))
										if uint32(t403-t404) > uint32(i32(259)) {
											m.memory[uint32(v8)] = byte(i32(18))
											t478 := int64(load64(m.memory[int64(uint32(v7))+16:]))
											store64(m.memory[uint32(v19):], uint64(t478))
											t479 := int32(load32(m.memory[int64(uint32(v7))+24:]))
											store32(m.memory[int64(uint32(v19))+8:], uint32(t479))
											m.memory[int64(uint32(v8))+64] = byte(v22)
											store32(m.memory[int64(uint32(v8))+60:], uint32(v3))
											store32(m.memory[int64(uint32(v8))+56:], uint32(v24))
											store64(m.memory[int64(uint32(v8))+48:], uint64(v25))
											t480 := int32(load32(m.memory[int64(uint32(v7))+8:]))
											store32(m.memory[uint32(v12):], uint32(t480))
											t481 := int32(load32(m.memory[int64(uint32(v7))+11:]))
											store32(m.memory[int64(uint32(v12))+3:], uint32(t481))
											m.fn910(v8)
											t482 := int32(m.memory[uint32(v8)])
											v9 = t482
											goto l18
										}
									}
								l204:
									store32(m.memory[int64(uint32(v8))+100:], uint32(i32(0)))
									v5 = v24
									v9 = v22
									t405 := int64(load32(m.memory[int64(uint32(v8))+148:]))
									t406 := v28
									t407 := v25
									v31 = i64_shl(i64(-1), t405) ^ i64(-1)
									v20 = int32(t407 & v31)
									if uint32(t406) > uint32(v20) {
									l217:
										{
											v23 = v21 + v20<<2
											t411 := int32(m.memory[int64(uint32(v23))+3])
											v20 = t411
											if uint32(v20) <= uint32(v9&i32(255)) {
												goto l215
											}
											{
												if v5 == v3 {
													goto l216
												}
												v30 = int64(uint32(v9))
												t412 := int64(m.memory[uint32(v5)])
												v26 = t412
												v5 = v5 + i32(1)
												v9 = v9 + i32(8)
												t413 := v28
												v25 = i64_shl(v26, v30) | v25
												v20 = int32(v25 & v31)
												if uint32(t413) <= uint32(v20) {
													goto l207
												}
												goto l217
											}
										l216:
										}
										v20 = i32(18)
										m.memory[uint32(v8)] = byte(i32(18))
										t414 := int64(load64(m.memory[int64(uint32(v7))+16:]))
										store64(m.memory[uint32(v19):], uint64(t414))
										t415 := int32(load32(m.memory[int64(uint32(v7))+24:]))
										store32(m.memory[int64(uint32(v19))+8:], uint32(t415))
										store32(m.memory[int64(uint32(v8))+60:], uint32(v3))
										store32(m.memory[int64(uint32(v8))+56:], uint32(v3))
										store64(m.memory[int64(uint32(v8))+48:], uint64(v25))
										t416 := int32(load32(m.memory[int64(uint32(v7))+8:]))
										store32(m.memory[uint32(v12):], uint32(t416))
										t417 := int32(load32(m.memory[int64(uint32(v7))+11:]))
										store32(m.memory[int64(uint32(v12))+3:], uint32(t417))
										m.memory[int64(uint32(v8))+64] = byte(v22 + v3<<3 - v24<<3)
										v3 = i32(0)
										goto l30
									}
									goto l207
								case 2:
									t408 := int32(load32(m.memory[int64(uint32(v8))+96:]))
									v20 = t408
									if v20 != 0 {
										v9 = v24
										v5 = v22
										if uint32(v20) > uint32(v22&i32(255)) {
											goto l222
										}
										v5 = v22
										v23 = v24
										goto l211
									}
									t409 := int32(load32(m.memory[int64(uint32(v8))+88:]))
									v9 = t409
									goto l209
								case 4:
									v20 = i32(23)
									t410 := int32(load32(m.memory[int64(uint32(v8))+96:]))
									v23 = t410
									if v23 == 0 {
										goto l212
									}
									v9 = v24
									v5 = v22
									if uint32(v23) > uint32(v22&i32(255)) {
										goto l228
									}
									v5 = v22
									v32 = v24
									goto l214
								}
							l222:
								{
									if v9 == v3 {
										v20 = i32(20)
										m.memory[uint32(v8)] = byte(i32(20))
										t487 := int64(load64(m.memory[int64(uint32(v7))+16:]))
										store64(m.memory[uint32(v19):], uint64(t487))
										t488 := int32(load32(m.memory[int64(uint32(v7))+24:]))
										store32(m.memory[int64(uint32(v19))+8:], uint32(t488))
										store32(m.memory[int64(uint32(v8))+60:], uint32(v3))
										store32(m.memory[int64(uint32(v8))+56:], uint32(v3))
										store64(m.memory[int64(uint32(v8))+48:], uint64(v25))
										t489 := int32(load32(m.memory[int64(uint32(v7))+8:]))
										store32(m.memory[uint32(v12):], uint32(t489))
										t490 := int32(load32(m.memory[int64(uint32(v7))+11:]))
										store32(m.memory[int64(uint32(v12))+3:], uint32(t490))
										m.memory[int64(uint32(v8))+64] = byte(v22 + v3<<3 - v24<<3)
										v3 = i32(0)
										goto l30
									}
									t423 := int64(m.memory[uint32(v9)])
									v25 = i64_shl(t423, int64(uint32(v5))) | v25
									v23 = v9 + i32(1)
									v9 = v23
									t424 := v20
									v5 = v5 + i32(8)
									if uint32(t424) > uint32(v5&i32(255)) {
										goto l222
									}
								}
							l211:
								t425 := int32(load32(m.memory[int64(uint32(v8))+100:]))
								store32(m.memory[int64(uint32(v8))+100:], uint32(t425+v20))
								t426 := int32(load32(m.memory[int64(uint32(v8))+88:]))
								t427 := v8
								t428 := v25
								v30 = int64(uint32(v20))
								v9 = t426 + int32(t428&(i64_shl(i64(-1), v30)^i64(-1)))
								store32(m.memory[int64(uint32(t427))+88:], uint32(v9))
								v22 = v5 - v20
								v25 = i64_shr_u(v25, v30)
								v24 = v23
							}
						l209:
							store32(m.memory[int64(uint32(v8))+104:], uint32(v9))
							v20 = i32(21)
							goto l212
						l228:
							{
								if v9 == v3 {
									v20 = i32(22)
									m.memory[uint32(v8)] = byte(i32(22))
									t491 := int64(load64(m.memory[int64(uint32(v7))+16:]))
									store64(m.memory[uint32(v19):], uint64(t491))
									t492 := int32(load32(m.memory[int64(uint32(v7))+24:]))
									store32(m.memory[int64(uint32(v19))+8:], uint32(t492))
									store32(m.memory[int64(uint32(v8))+60:], uint32(v3))
									store32(m.memory[int64(uint32(v8))+56:], uint32(v3))
									store64(m.memory[int64(uint32(v8))+48:], uint64(v25))
									t493 := int32(load32(m.memory[int64(uint32(v7))+8:]))
									store32(m.memory[uint32(v12):], uint32(t493))
									t494 := int32(load32(m.memory[int64(uint32(v7))+11:]))
									store32(m.memory[int64(uint32(v12))+3:], uint32(t494))
									m.memory[int64(uint32(v8))+64] = byte(v22 + v3<<3 - v24<<3)
									v3 = i32(0)
									goto l30
								}
								t439 := int64(m.memory[uint32(v9)])
								v25 = i64_shl(t439, int64(uint32(v5))) | v25
								v32 = v9 + i32(1)
								v9 = v32
								t440 := v23
								v5 = v5 + i32(8)
								if uint32(t440) > uint32(v5&i32(255)) {
									goto l228
								}
							}
						l214:
							t441 := int32(load32(m.memory[int64(uint32(v8))+100:]))
							store32(m.memory[int64(uint32(v8))+100:], uint32(t441+v23))
							t442 := int32(load32(m.memory[int64(uint32(v8))+92:]))
							t443 := v8
							t444 := v25
							v30 = int64(uint32(v23))
							store32(m.memory[int64(uint32(t443))+92:], uint32(t442+int32(t444&(i64_shl(i64(-1), v30)^i64(-1)))))
							v22 = v5 - v23
							v25 = i64_shr_u(v25, v30)
							v24 = v32
							goto l212
						}
					l231:
						t499 := int32(load32(m.memory[int64(uint32(v8))+88:]))
						t500 := v8
						v5 = t499
						store32(m.memory[int64(uint32(t500))+88:], uint32(v5-v9))
						p501 := i32(23)
						if v5 == v9 {
							p501 = i32(18)
						}
						v20 = p501
						goto l212
					}
				l224:
					t502 := int32(load16(m.memory[uint32(v23):]))
					v32 = t502
					{
						{
							t503 := int32(m.memory[int64(uint32(v23))+2])
							v34 = t503
							if uint32(v34) <= uint32(i32(15)) {
								goto l236
							}
							v23 = v20
							goto l237
						}
					l236:
						t504 := v25
						v31 = i64_shl(i64(-1), int64(uint32(v34+v20))) ^ i64(-1)
						v23 = i32_shr_u(int32(t504&v31), v20) + v32
						if uint32(v23) >= uint32(v29) {
							goto l238
						}
						v34 = v20 & i32(31)
						{
						l241:
							{
								v33 = v27 + v23<<2
								t505 := int32(m.memory[int64(uint32(v33))+3])
								v23 = t505
								if uint32((v23+v20)&i32(255)) <= uint32(v9&i32(255)) {
									goto l239
								}
								{
									if v5 == v3 {
										goto l240
									}
									v30 = int64(uint32(v9))
									t506 := int64(m.memory[uint32(v5)])
									v26 = t506
									v5 = v5 + i32(1)
									v9 = v9 + i32(8)
									v25 = i64_shl(v26, v30) | v25
									v23 = i32_shr_u(int32(v25&v31), v34) + v32
									if uint32(v23) >= uint32(v29) {
										goto l238
									}
									goto l241
								}
							l240:
							}
							v20 = i32(21)
							m.memory[uint32(v8)] = byte(i32(21))
							t507 := int64(load64(m.memory[int64(uint32(v7))+16:]))
							store64(m.memory[uint32(v19):], uint64(t507))
							t508 := int32(load32(m.memory[int64(uint32(v7))+24:]))
							store32(m.memory[int64(uint32(v19))+8:], uint32(t508))
							store32(m.memory[int64(uint32(v8))+60:], uint32(v3))
							store32(m.memory[int64(uint32(v8))+56:], uint32(v3))
							store64(m.memory[int64(uint32(v8))+48:], uint64(v25))
							t509 := int32(load32(m.memory[int64(uint32(v7))+8:]))
							store32(m.memory[uint32(v12):], uint32(t509))
							t510 := int32(load32(m.memory[int64(uint32(v7))+11:]))
							store32(m.memory[int64(uint32(v12))+3:], uint32(t510))
							m.memory[int64(uint32(v8))+64] = byte(v22 + v3<<3 - v24<<3)
							v3 = i32(0)
							goto l30
						}
					l239:
						t511 := int32(m.memory[int64(uint32(v33))+2])
						v34 = t511
						t512 := int32(load16(m.memory[uint32(v33):]))
						v32 = t512
						t513 := int32(load32(m.memory[int64(uint32(v8))+100:]))
						store32(m.memory[int64(uint32(v8))+100:], uint32(t513+v20))
						v9 = v9 - v20
						v25 = i64_shr_u(v25, int64(uint32(v20)))
					}
				l237:
					v22 = v9 - v23
					v25 = i64_shr_u(v25, int64(uint32(v23)))
					if v34&i32(64) != 0 {
						goto l242
					}
					store32(m.memory[int64(uint32(v8))+96:], uint32(v34&i32(15)))
					store32(m.memory[int64(uint32(v8))+92:], uint32(v32&i32(0xffff)))
					v20 = i32(22)
					v24 = v5
					goto l212
				l242:
					t514 := int32(load32(m.memory[int64(uint32(v7))+24:]))
					store32(m.memory[int64(uint32(v19))+8:], uint32(t514))
					t515 := int64(load64(m.memory[int64(uint32(v7))+16:]))
					store64(m.memory[uint32(v19):], uint64(t515))
					m.memory[int64(uint32(v8))+64] = byte(v22)
					store32(m.memory[int64(uint32(v8))+60:], uint32(v3))
					store32(m.memory[int64(uint32(v8))+56:], uint32(v5))
					store64(m.memory[int64(uint32(v8))+48:], uint64(v25))
					t516 := int32(load32(m.memory[int64(uint32(v7))+8:]))
					store32(m.memory[uint32(v12):], uint32(t516))
					t517 := int32(load32(m.memory[int64(uint32(v7))+11:]))
					store32(m.memory[int64(uint32(v12))+3:], uint32(t517))
					store32(m.memory[int64(uint32(v8))+136:], uint32(i32(22)))
					store32(m.memory[int64(uint32(v8))+132:], uint32(i32(1065867)))
					v20 = i32(30)
				}
			l235:
				m.memory[uint32(v8)] = byte(v20)
				v3 = i32(-3)
				goto l30
			l238:
				m.fn33(v23, v29, i32(1284572))
				panic("unreachable")
			l223:
				m.fn33(v20, v29, i32(1284556))
				panic("unreachable")
			l215:
				t518 := int32(load16(m.memory[uint32(v23):]))
				v32 = t518
				{
					t519 := int32(m.memory[int64(uint32(v23))+2])
					v34 = t519
					if v34 != 0 {
						v23 = i32(0)
						if uint32(v34) <= uint32(i32(15)) {
							{
								t520 := v28
								t521 := v25
								v31 = i64_shl(i64(-1), int64(uint32(v34+v20))) ^ i64(-1)
								t522 := int32(t521&v31) & i32(0xffff)
								v34 = v20 & i32(15)
								v23 = (i32_shr_u(t522, v34) + v32) & i32(0xffff)
								if uint32(t520) <= uint32(v23) {
									goto l246
								}
								{
								l249:
									{
										v23 = v21 + v23<<2
										t523 := int32(m.memory[int64(uint32(v23))+3])
										v33 = t523
										if uint32((v33+v20)&i32(255)) <= uint32(v9&i32(255)) {
											v9 = v9 - v20
											v25 = i64_shr_u(v25, int64(uint32(v20)))
											t530 := int32(m.memory[int64(uint32(v23))+2])
											v34 = t530
											t531 := int32(load16(m.memory[uint32(v23):]))
											v32 = t531
											v23 = v20
											v24 = v5
											v20 = v33
											goto l244
										}
										{
											if v5 == v3 {
												goto l248
											}
											v30 = int64(uint32(v9))
											t524 := int64(m.memory[uint32(v5)])
											v26 = t524
											v5 = v5 + i32(1)
											v9 = v9 + i32(8)
											t525 := v28
											v25 = i64_shl(v26, v30) | v25
											v23 = (i32_shr_u(int32(v25&v31)&i32(0xffff), v34) + v32) & i32(0xffff)
											if uint32(t525) <= uint32(v23) {
												goto l246
											}
											goto l249
										}
									l248:
									}
									v20 = i32(18)
									m.memory[uint32(v8)] = byte(i32(18))
									t526 := int64(load64(m.memory[int64(uint32(v7))+16:]))
									store64(m.memory[uint32(v19):], uint64(t526))
									t527 := int32(load32(m.memory[int64(uint32(v7))+24:]))
									store32(m.memory[int64(uint32(v19))+8:], uint32(t527))
									store32(m.memory[int64(uint32(v8))+60:], uint32(v3))
									store32(m.memory[int64(uint32(v8))+56:], uint32(v3))
									store64(m.memory[int64(uint32(v8))+48:], uint64(v25))
									t528 := int32(load32(m.memory[int64(uint32(v7))+8:]))
									store32(m.memory[uint32(v12):], uint32(t528))
									t529 := int32(load32(m.memory[int64(uint32(v7))+11:]))
									store32(m.memory[int64(uint32(v12))+3:], uint32(t529))
									m.memory[int64(uint32(v8))+64] = byte(v22 + v3<<3 - v24<<3)
									v3 = i32(0)
									goto l30
								}
							}
						l246:
							m.fn33(v23, v28, i32(1284540))
							panic("unreachable")
						}
						v24 = v5
						goto l244
					}
					v23 = i32(0)
					v24 = v5
					v34 = i32(0)
					goto l244
				}
			l244:
				store32(m.memory[int64(uint32(v8))+88:], uint32(v32&i32(0xffff)))
				store32(m.memory[int64(uint32(v8))+100:], uint32(v23+v20&i32(255)))
				v22 = v9 - v20
				v25 = i64_shr_u(v25, int64(uint32(v20)))
				v20 = i32(19)
				if v34&i32(255) == 0 {
					goto l212
				}
				if v34&i32(32) != 0 {
					v9 = i32(12)
					m.memory[uint32(v8)] = byte(i32(12))
					store32(m.memory[int64(uint32(v8))+100:], uint32(i32(-1)))
					t536 := int64(load64(m.memory[int64(uint32(v7))+16:]))
					store64(m.memory[uint32(v19):], uint64(t536))
					t537 := int32(load32(m.memory[int64(uint32(v7))+24:]))
					store32(m.memory[int64(uint32(v19))+8:], uint32(t537))
					m.memory[int64(uint32(v8))+64] = byte(v22)
					store32(m.memory[int64(uint32(v8))+60:], uint32(v3))
					store32(m.memory[int64(uint32(v8))+56:], uint32(v24))
					store64(m.memory[int64(uint32(v8))+48:], uint64(v25))
					t538 := int32(load32(m.memory[int64(uint32(v7))+8:]))
					store32(m.memory[uint32(v12):], uint32(t538))
					t539 := int32(load32(m.memory[int64(uint32(v7))+11:]))
					store32(m.memory[int64(uint32(v12))+3:], uint32(t539))
					goto l18
				}
				if v34&i32(64) != 0 {
					goto l251
				}
				store32(m.memory[int64(uint32(v8))+96:], uint32(v34&i32(15)))
				v20 = i32(20)
				goto l212
			l251:
			}
			v20 = i32(30)
			m.memory[uint32(v8)] = byte(i32(30))
			t532 := int64(load64(m.memory[int64(uint32(v7))+16:]))
			store64(m.memory[uint32(v19):], uint64(t532))
			t533 := int32(load32(m.memory[int64(uint32(v7))+24:]))
			store32(m.memory[int64(uint32(v19))+8:], uint32(t533))
			m.memory[int64(uint32(v8))+64] = byte(v22)
			store32(m.memory[int64(uint32(v8))+60:], uint32(v3))
			store32(m.memory[int64(uint32(v8))+56:], uint32(v24))
			store64(m.memory[int64(uint32(v8))+48:], uint64(v25))
			t534 := int32(load32(m.memory[int64(uint32(v7))+8:]))
			store32(m.memory[uint32(v12):], uint32(t534))
			t535 := int32(load32(m.memory[int64(uint32(v7))+11:]))
			store32(m.memory[int64(uint32(v12))+3:], uint32(t535))
			store32(m.memory[int64(uint32(v8))+136:], uint32(i32(28)))
			store32(m.memory[int64(uint32(v8))+132:], uint32(i32(1065839)))
			v3 = i32(-3)
			goto l30
		}
	l207:
		m.fn33(v20, v28, i32(1284524))
		panic("unreachable")
	case 19:
		t381 := int32(load32(m.memory[int64(uint32(v8))+80:]))
		v5 = t381
		t382 := int32(load32(m.memory[int64(uint32(v8))+76:]))
		t383 := v5
		v3 = t382
		if t383 != v3 {
			if uint32(v5) >= uint32(v3) {
				m.fn33(v5, v3, i32(1279896))
				panic("unreachable")
			}
			t384 := int32(load32(m.memory[int64(uint32(v8))+72:]))
			t385 := int32(load32(m.memory[int64(uint32(v8))+88:]))
			m.memory[uint32(t384+v5)] = byte(t385)
			t386 := int32(load32(m.memory[int64(uint32(v8))+80:]))
			store32(m.memory[int64(uint32(v8))+80:], uint32(t386+i32(1)))
			goto l18
		}
		v3 = i32(0)
		v20 = i32(19)
		goto l30
	case 20:
		{
			{
				t365 := int32(load32(m.memory[int64(uint32(v8))+96:]))
				v29 = t365
				if v29 != 0 {
					goto l181
				}
				t366 := int32(load32(m.memory[int64(uint32(v8))+88:]))
				v9 = t366
				goto l182
			}
		l181:
			t367 := int64(load64(m.memory[int64(uint32(v8))+48:]))
			v25 = t367
			{
				{
					t368 := int32(m.memory[int64(uint32(v8))+64])
					t369 := v29
					v5 = t368
					if uint32(t369) > uint32(v5) {
						goto l183
					}
					v3 = v5
					goto l184
				}
			l183:
				t370 := int32(load32(m.memory[int64(uint32(v8))+56:]))
				v9 = t370
				t371 := int32(load32(m.memory[int64(uint32(v8))+60:]))
				v27 = t371
			l186:
				{
					if v9 != v27 {
						goto l185
					}
					v3 = i32(0)
					v20 = i32(20)
					goto l30
				l185:
					t372 := v8
					v20 = v9 + i32(1)
					store32(m.memory[int64(uint32(t372))+56:], uint32(v20))
					t373 := v8
					v3 = v5 + i32(8)
					m.memory[int64(uint32(t373))+64] = byte(v3)
					t374 := int64(m.memory[uint32(v9)])
					t375 := v8
					v25 = i64_shl(t374, int64(uint32(v5))) | v25
					store64(m.memory[int64(uint32(t375))+48:], uint64(v25))
					v9 = v20
					v5 = v3
					if uint32(v29) > uint32(v3&i32(255)) {
						goto l186
					}
				}
			}
		l184:
			m.memory[int64(uint32(v8))+64] = byte(v3 - v29)
			t376 := v8
			t377 := v25
			v30 = int64(uint32(v29))
			store64(m.memory[int64(uint32(t376))+48:], uint64(i64_shr_u(t377, v30)))
			t378 := int32(load32(m.memory[int64(uint32(v8))+100:]))
			store32(m.memory[int64(uint32(v8))+100:], uint32(t378+v29))
			t379 := int32(load32(m.memory[int64(uint32(v8))+88:]))
			t380 := v8
			v9 = t379 + int32(v25&(i64_shl(i64(-1), v30)^i64(-1)))
			store32(m.memory[int64(uint32(t380))+88:], uint32(v9))
		}
	l182:
		store32(m.memory[int64(uint32(v8))+104:], uint32(v9))
		v9 = i32(21)
		goto l18
	case 21:
		t79 := int64(load32(m.memory[int64(uint32(v8))+156:]))
		v26 = i64_shl(i64(-1), t79) ^ i64(-1)
		t80 := int32(load32(m.memory[int64(uint32(v8))+56:]))
		v5 = t80
		t81 := int32(m.memory[int64(uint32(v8))+64])
		v9 = t81
		t82 := int32(load32(m.memory[int64(uint32(v8))+60:]))
		v28 = t82
		t83 := int64(load64(m.memory[int64(uint32(v8))+48:]))
		v25 = t83
		t84 := int32(m.memory[int64(uint32(v8))+160])
		v27 = t84
		{
		l93:
			v3 = int32(v25 & v26)
			v20 = i32(32)
			v29 = i32(1279736)
			switch v27 {
			default:
				goto l83
			case 1:
				v20 = i32(1332)
				v29 = v16
				goto l83
			case 2:
				v20 = i32(1332)
				v29 = v17
				goto l83
			case 3:
				v20 = i32(592)
				v29 = v18
			}
		l83:
			if uint32(v20) <= uint32(v3) {
				m.fn33(v3, v20, i32(1284692))
				panic("unreachable")
			}
			{
				v20 = v29 + v3<<2
				t85 := int32(m.memory[int64(uint32(v20))+3])
				v3 = t85
				if uint32(v3) <= uint32(v9&i32(255)) {
					t86 := int32(load16(m.memory[uint32(v20):]))
					v23 = t86
					t87 := int32(m.memory[int64(uint32(v20))+2])
					v21 = t87
					if uint32(v21&i32(255)) <= uint32(i32(15)) {
						goto l91
					}
					v20 = v3
					goto l92
				}
				if v5 != v28 {
					t88 := int64(m.memory[uint32(v5)])
					v30 = t88
					t89 := v8
					v3 = v9 + i32(8)
					m.memory[int64(uint32(t89))+64] = byte(v3)
					t90 := v8
					v5 = v5 + i32(1)
					store32(m.memory[int64(uint32(t90))+56:], uint32(v5))
					t91 := v8
					v25 = i64_shl(v30, int64(uint32(v9))) | v25
					store64(m.memory[int64(uint32(t91))+48:], uint64(v25))
					v9 = v3
					goto l93
				}
				goto l90
			}
		l91:
			v24 = v3 & i32(31)
			v26 = i64_shl(i64(-1), int64(uint32(v3+v21))) ^ i64(-1)
		l100:
			v20 = i32_shr_u(int32(v25&v26), v24) + v23
			v29 = i32(32)
			v21 = i32(1279736)
			switch v27 {
			default:
				goto l94
			case 1:
				v29 = i32(1332)
				v21 = v16
				goto l94
			case 2:
				v29 = i32(1332)
				v21 = v17
				goto l94
			case 3:
				v29 = i32(592)
				v21 = v18
			}
		l94:
			if uint32(v20) >= uint32(v29) {
				m.fn33(v20, v29, i32(1284692))
				panic("unreachable")
			}
			{
				v29 = v21 + v20<<2
				t92 := int32(m.memory[int64(uint32(v29))+3])
				v20 = t92
				if uint32((v20+v3)&i32(255)) <= uint32(v9&i32(255)) {
					goto l99
				}
				if v5 == v28 {
					goto l90
				}
				t93 := int64(m.memory[uint32(v5)])
				v30 = t93
				t94 := v8
				v20 = v9 + i32(8)
				m.memory[int64(uint32(t94))+64] = byte(v20)
				t95 := v8
				v5 = v5 + i32(1)
				store32(m.memory[int64(uint32(t95))+56:], uint32(v5))
				t96 := v8
				v25 = i64_shl(v30, int64(uint32(v9))) | v25
				store64(m.memory[int64(uint32(t96))+48:], uint64(v25))
				v9 = v20
				goto l100
			}
		l99:
			t97 := int32(m.memory[int64(uint32(v29))+2])
			v21 = t97
			t98 := int32(load16(m.memory[uint32(v29):]))
			v23 = t98
			t99 := int32(load32(m.memory[int64(uint32(v8))+100:]))
			store32(m.memory[int64(uint32(v8))+100:], uint32(t99+v3))
			v9 = v9 - v3
			v25 = i64_shr_u(v25, int64(uint32(v3)))
		}
	l92:
		m.memory[int64(uint32(v8))+64] = byte(v9 - v20)
		store64(m.memory[int64(uint32(v8))+48:], uint64(i64_shr_u(v25, int64(uint32(v20)))))
		if v21&i32(64) != 0 {
			store32(m.memory[int64(uint32(v8))+136:], uint32(i32(22)))
			store32(m.memory[int64(uint32(v8))+132:], uint32(i32(1065867)))
			goto l57
		}
		store32(m.memory[int64(uint32(v8))+96:], uint32(v21&i32(15)))
		store32(m.memory[int64(uint32(v8))+92:], uint32(v23&i32(0xffff)))
		v9 = i32(22)
		goto l18
	l90:
		v3 = i32(0)
		v20 = i32(21)
		goto l30
	case 22:
		v9 = i32(23)
		t65 := int32(load32(m.memory[int64(uint32(v8))+96:]))
		v27 = t65
		if v27 == 0 {
			goto l18
		}
		t66 := int64(load64(m.memory[int64(uint32(v8))+48:]))
		v25 = t66
		{
			{
				t67 := int32(m.memory[int64(uint32(v8))+64])
				t68 := v27
				v3 = t67
				if uint32(t68) > uint32(v3) {
					goto l79
				}
				v20 = v3
				goto l80
			}
		l79:
			t69 := int32(load32(m.memory[int64(uint32(v8))+56:]))
			v5 = t69
			t70 := int32(load32(m.memory[int64(uint32(v8))+60:]))
			v28 = t70
		l82:
			{
				if v5 != v28 {
					goto l81
				}
				v3 = i32(0)
				v20 = i32(22)
				goto l30
			l81:
				t71 := v8
				v29 = v5 + i32(1)
				store32(m.memory[int64(uint32(t71))+56:], uint32(v29))
				t72 := v8
				v20 = v3 + i32(8)
				m.memory[int64(uint32(t72))+64] = byte(v20)
				t73 := int64(m.memory[uint32(v5)])
				t74 := v8
				v25 = i64_shl(t73, int64(uint32(v3))) | v25
				store64(m.memory[int64(uint32(t74))+48:], uint64(v25))
				v5 = v29
				v3 = v20
				if uint32(v27) > uint32(v20&i32(255)) {
					goto l82
				}
			}
		}
	l80:
		m.memory[int64(uint32(v8))+64] = byte(v20 - v27)
		t75 := v8
		t76 := v25
		v30 = int64(uint32(v27))
		store64(m.memory[int64(uint32(t75))+48:], uint64(i64_shr_u(t76, v30)))
		t77 := int32(load32(m.memory[int64(uint32(v8))+100:]))
		store32(m.memory[int64(uint32(v8))+100:], uint32(t77+v27))
		t78 := int32(load32(m.memory[int64(uint32(v8))+92:]))
		store32(m.memory[int64(uint32(v8))+92:], uint32(t78+int32(v25&(i64_shl(i64(-1), v30)^i64(-1)))))
		goto l18
	case 23:
	l180:
		{
			t334 := int32(load32(m.memory[int64(uint32(v8))+76:]))
			v3 = t334
			t335 := int32(load32(m.memory[int64(uint32(v8))+80:]))
			t336 := v3
			v5 = t335
			if t336 != v5 {
				v3 = v3 - v5
				{
					t337 := int32(load32(m.memory[int64(uint32(v8))+92:]))
					v20 = t337
					if uint32(v20) > uint32(v5) {
						v5 = v20 - v5
						t343 := int32(load32(m.memory[int64(uint32(v8))+16:]))
						if uint32(v5) > uint32(t343) {
							t362 := int32(m.memory[int64(uint32(v8))+1])
							if t362&i32(4) == 0 {
								m.fn28(i32(1279992), i32(85), i32(1284708))
								panic("unreachable")
							}
							v20 = i32(30)
							store32(m.memory[int64(uint32(v8))+136:], uint32(i32(30)))
							store32(m.memory[int64(uint32(v8))+132:], uint32(i32(1065425)))
							v3 = i32(-3)
							goto l30
						}
						t344 := int32(load32(m.memory[int64(uint32(v8))+12:]))
						v29 = t344
						v20 = v29 + i32(-64)
						if uint32(v20) >= uint32(i32(-63)) {
							m.fn7(i32(1284212), i32(74), i32(1284288))
							panic("unreachable")
						}
						t346 := v19
						t347 := v13
						p345 := v20
						if uint32(v20) > uint32(v29) {
							p345 = i32(0)
						}
						t348 := int32(load32(m.memory[int64(uint32(v8))+20:]))
						t349 := v5
						v20 = t348
						v29 = t349 - v20
						t350 := p345 - v29
						t351 := v20 - v5
						var p352 int32
						if uint32(v5) > uint32(v20) {
							p352 = 1
						}
						v20 = p352
						p353 := t351
						if v20 != 0 {
							p353 = t350
						}
						v27 = p353
						t354 := int32(load32(m.memory[int64(uint32(v8))+88:]))
						t355 := v27
						t356 := v27
						t357 := v3
						v28 = t354
						t359 := v28
						p358 := v5
						if v20 != 0 {
							p358 = v29
						}
						v5 = p358
						p360 := v5
						if uint32(v28) < uint32(v5) {
							p360 = t359
						}
						v5 = p360
						p361 := v5
						if uint32(v3) < uint32(v5) {
							p361 = t357
						}
						v5 = p361
						m.fn909(t346, t347, t355, t356+v5)
						goto l176
					}
					t338 := int32(load32(m.memory[int64(uint32(v8))+88:]))
					t339 := v19
					t340 := v20
					t341 := v3
					v5 = t338
					p342 := v5
					if uint32(v3) < uint32(v5) {
						p342 = t341
					}
					v5 = p342
					m.fn908(t339, t340, v5)
					goto l176
				}
			l176:
				t363 := int32(load32(m.memory[int64(uint32(v8))+88:]))
				t364 := v8
				v3 = t363
				store32(m.memory[int64(uint32(t364))+88:], uint32(v3-v5))
				if v3 == v5 {
					goto l18
				}
				goto l180
			}
			v3 = i32(0)
			v20 = i32(23)
			goto l30
		}
	case 24:
		t320 := int64(load64(m.memory[int64(uint32(v8))+48:]))
		v25 = t320
		{
			t321 := int32(m.memory[int64(uint32(v8))+64])
			v9 = t321
			if uint32(v9) > uint32(i32(13)) {
				goto l171
			}
			v3 = i32(0)
			v20 = i32(24)
			t322 := int32(load32(m.memory[int64(uint32(v8))+56:]))
			v5 = t322
			t323 := int32(load32(m.memory[int64(uint32(v8))+60:]))
			t324 := v5
			v29 = t323
			if t324 == v29 {
				goto l30
			}
			t325 := v8
			v27 = v5 + i32(1)
			store32(m.memory[int64(uint32(t325))+56:], uint32(v27))
			t326 := v8
			v30 = int64(uint32(v9))
			v26 = v30 + i64(8)
			m.memory[int64(uint32(t326))+64] = byte(v26)
			t327 := int64(m.memory[uint32(v5)])
			t328 := v8
			v25 = i64_shl(t327, v30) | v25
			store64(m.memory[int64(uint32(t328))+48:], uint64(v25))
			if uint32(v9) <= uint32(i32(5)) {
				goto l172
			}
			v9 = int32(v26)
			goto l171
		l172:
			if v27 == v29 {
				goto l30
			}
			store32(m.memory[int64(uint32(v8))+56:], uint32(v5+i32(2)))
			t329 := v8
			v9 = v9 | i32(16)
			m.memory[int64(uint32(t329))+64] = byte(v9)
			t330 := int64(m.memory[int64(uint32(v5))+1])
			t331 := v8
			v25 = i64_shl(t330, v26) | v25
			store64(m.memory[int64(uint32(t331))+48:], uint64(v25))
		}
	l171:
		m.memory[int64(uint32(v8))+64] = byte(v9 + i32(-14))
		store64(m.memory[int64(uint32(v8))+48:], uint64(int64(uint64(v25)>>14)))
		t332 := v8
		v9 = int32(v25)
		v5 = v9 & i32(31)
		store32(m.memory[int64(uint32(t332))+28:], uint32(v5+i32(257)))
		t333 := v8
		v3 = int32(uint32(v9)>>5) & i32(31)
		store32(m.memory[int64(uint32(t333))+32:], uint32(v3+i32(1)))
		store32(m.memory[int64(uint32(v8))+24:], uint32(int32(uint32(v9)>>10)&i32(15)+i32(4)))
		if uint32(v5) > uint32(i32(29)) {
			goto l173
		}
		if uint32(v3) > uint32(i32(29)) {
			goto l173
		}
		store32(m.memory[int64(uint32(v8))+36:], uint32(i32(0)))
		v9 = i32(25)
		goto l18
	l173:
		store32(m.memory[int64(uint32(v8))+136:], uint32(i32(36)))
		store32(m.memory[int64(uint32(v8))+132:], uint32(i32(1064580)))
		goto l57
	case 25:
		goto l26
	case 27:
		t300 := int64(load64(m.memory[int64(uint32(v8))+48:]))
		v25 = t300
		{
			t301 := int32(m.memory[int64(uint32(v8))+64])
			v9 = t301
			if uint32(v9) > uint32(i32(31)) {
				goto l170
			}
			v3 = i32(0)
			v20 = i32(27)
			t302 := int32(load32(m.memory[int64(uint32(v8))+56:]))
			v5 = t302
			t303 := int32(load32(m.memory[int64(uint32(v8))+60:]))
			t304 := v5
			v29 = t303
			if t304 == v29 {
				goto l30
			}
			t305 := v8
			v27 = v5 + i32(1)
			store32(m.memory[int64(uint32(t305))+56:], uint32(v27))
			t306 := v8
			v30 = int64(uint32(v9))
			v26 = v30 + i64(8)
			m.memory[int64(uint32(t306))+64] = byte(v26)
			t307 := int64(m.memory[uint32(v5)])
			t308 := v8
			v25 = i64_shl(t307, v30) | v25
			store64(m.memory[int64(uint32(t308))+48:], uint64(v25))
			if uint32(v9) > uint32(i32(23)) {
				goto l170
			}
			if v27 == v29 {
				goto l30
			}
			t309 := v8
			v27 = v5 + i32(2)
			store32(m.memory[int64(uint32(t309))+56:], uint32(v27))
			t310 := v8
			v31 = v30 + i64(16)
			m.memory[int64(uint32(t310))+64] = byte(v31)
			t311 := int64(m.memory[int64(uint32(v5))+1])
			t312 := v8
			v25 = i64_shl(t311, v26) | v25
			store64(m.memory[int64(uint32(t312))+48:], uint64(v25))
			if uint32(v9) > uint32(i32(15)) {
				goto l170
			}
			if v27 == v29 {
				goto l30
			}
			t313 := v8
			v27 = v5 + i32(3)
			store32(m.memory[int64(uint32(t313))+56:], uint32(v27))
			t314 := v8
			v30 = v30 + i64(24)
			m.memory[int64(uint32(t314))+64] = byte(v30)
			t315 := int64(m.memory[int64(uint32(v5))+2])
			t316 := v8
			v25 = i64_shl(t315, v31) | v25
			store64(m.memory[int64(uint32(t316))+48:], uint64(v25))
			if uint32(v9) > uint32(i32(7)) {
				goto l170
			}
			if v27 == v29 {
				goto l30
			}
			store32(m.memory[int64(uint32(v8))+56:], uint32(v5+i32(4)))
			m.memory[int64(uint32(v8))+64] = byte(v9 | i32(32))
			t317 := int64(m.memory[int64(uint32(v5))+3])
			t318 := v8
			v25 = i64_shl(t317, v30) | v25
			store64(m.memory[int64(uint32(t318))+48:], uint64(v25))
		}
	l170:
		m.memory[int64(uint32(v8))+64] = byte(i32(0))
		store64(m.memory[int64(uint32(v8))+48:], uint64(i64(0)))
		t319 := v8
		v9 = int32(v25)
		store32(m.memory[int64(uint32(t319))+124:], uint32(i32_rotr(v9&i32(0xff00ff), i32(8))|i32_rotr(v9, i32(24))&i32(0xff00ff)))
		v9 = i32(28)
		goto l18
	case 28:
		v3 = i32(2)
		{
			t299 := int32(m.memory[int64(uint32(v8))+1])
			if t299&i32(2) != 0 {
				store32(m.memory[int64(uint32(v8))+124:], uint32(i32(1)))
				v9 = i32(12)
				goto l18
			}
			v20 = i32(28)
			goto l30
		}
	case 29:
		goto l30
	case 30:
		store32(m.memory[int64(uint32(v8))+136:], uint32(i32(29)))
		store32(m.memory[int64(uint32(v8))+132:], uint32(i32(1065691)))
		goto l57
	default:
		v9 = i32(13)
		t277 := int32(m.memory[int64(uint32(v8))+3])
		v5 = t277
		if v5 == 0 {
			goto l18
		}
		t278 := int64(load64(m.memory[int64(uint32(v8))+48:]))
		v25 = t278
		{
			t279 := int32(m.memory[int64(uint32(v8))+64])
			v9 = t279
			if uint32(v9) > uint32(i32(15)) {
				goto l158
			}
			t280 := int32(load32(m.memory[int64(uint32(v8))+56:]))
			v3 = t280
			t281 := int32(load32(m.memory[int64(uint32(v8))+60:]))
			t282 := v3
			v20 = t281
			if t282 == v20 {
				goto l159
			}
			t283 := v8
			v29 = v3 + i32(1)
			store32(m.memory[int64(uint32(t283))+56:], uint32(v29))
			t284 := v8
			v30 = int64(uint32(v9))
			v26 = v30 + i64(8)
			m.memory[int64(uint32(t284))+64] = byte(v26)
			t285 := int64(m.memory[uint32(v3)])
			t286 := v8
			v25 = i64_shl(t285, v30) | v25
			store64(m.memory[int64(uint32(t286))+48:], uint64(v25))
			if uint32(v9) > uint32(i32(7)) {
				goto l158
			}
			if v29 == v20 {
				goto l159
			}
			store32(m.memory[int64(uint32(v8))+56:], uint32(v3+i32(2)))
			m.memory[int64(uint32(v8))+64] = byte(v9 | i32(16))
			t287 := int64(m.memory[int64(uint32(v3))+1])
			t288 := v8
			v25 = i64_shl(t287, v26) | v25
			store64(m.memory[int64(uint32(t288))+48:], uint64(v25))
		}
	l158:
		if v5&i32(2) == 0 {
			goto l160
		}
		if v25 == i64(35615) {
			{
				t293 := int32(m.memory[int64(uint32(v8))+2])
				if t293 != 0 {
					goto l165
				}
				m.memory[int64(uint32(v8))+2] = byte(i32(15))
			}
		l165:
			store16(m.memory[int64(uint32(v7))+16:], uint16(i32(35615)))
			t294 := m.fn907(i32(0), v7+i32(16), i32(2))
			store32(m.memory[int64(uint32(v8))+124:], uint32(t294))
			m.memory[int64(uint32(v8))+64] = byte(i32(0))
			store64(m.memory[int64(uint32(v8))+48:], uint64(i64(0)))
			v9 = i32(1)
			goto l18
		}
	l160:
		{
			t289 := int32(load32(m.memory[int64(uint32(v8))+140:]))
			v9 = t289
			if v9 == 0 {
				goto l162
			}
			store32(m.memory[int64(uint32(v9))+48:], uint32(i32(-1)))
			t290 := int32(m.memory[int64(uint32(v8))+3])
			v5 = t290
		}
	l162:
		{
			if v5&i32(1) == 0 {
				goto l163
			}
			t291 := int64(load64(m.memory[int64(uint32(v8))+48:]))
			v25 = t291
			t292 := int64(uint64(v25<<8&i64(0xff00)+int64(uint64(v25)>>8)) % uint64(i64(31)))
			if t292 == 0 {
				if v25&i64(15) != i64(8) {
					store32(m.memory[int64(uint32(v8))+136:], uint32(i32(27)))
					store32(m.memory[int64(uint32(v8))+132:], uint32(i32(1066020)))
					goto l57
				}
				t295 := v8
				v30 = int64(uint64(v25) >> 4)
				store64(m.memory[int64(uint32(t295))+48:], uint64(v30))
				t296 := int32(m.memory[int64(uint32(v8))+64])
				m.memory[int64(uint32(v8))+64] = byte(t296 + i32(-4))
				v3 = int32(v30) & i32(15)
				v9 = v3 + i32(8)
				{
					t297 := int32(m.memory[int64(uint32(v8))+2])
					v5 = t297
					if v5 != 0 {
						goto l167
					}
					m.memory[int64(uint32(v8))+2] = byte(v9)
					v5 = v9
				}
			l167:
				{
					if uint32(v3) > uint32(i32(7)) {
						goto l168
					}
					if uint32(v9) > uint32(v5) {
						goto l168
					}
					store64(m.memory[int64(uint32(v8))+120:], uint64(i64(0x100000000)))
					m.memory[int64(uint32(v8))+64] = byte(i32(0))
					store64(m.memory[int64(uint32(v8))+48:], uint64(i64(0)))
					store32(m.memory[int64(uint32(v8))+144:], uint32(i32_shl(i32(1), v9)))
					p298 := i32(27)
					if v25&i64(8192) == 0 {
						p298 = i32(12)
					}
					v9 = p298
					goto l18
				}
			l168:
				store32(m.memory[int64(uint32(v8))+136:], uint32(i32(20)))
				store32(m.memory[int64(uint32(v8))+132:], uint32(i32(1065522)))
				goto l57
			}
		}
	l163:
		store32(m.memory[int64(uint32(v8))+136:], uint32(i32(23)))
		store32(m.memory[int64(uint32(v8))+132:], uint32(i32(1065358)))
		goto l57
	l159:
		v20 = i32(0)
		v3 = i32(0)
		goto l30
	case 26:
		t5 := int32(load32(m.memory[int64(uint32(v8))+36:]))
		v21 = t5
		t6 := int32(load32(m.memory[int64(uint32(v8))+32:]))
		t7 := v21
		v22 = t6
		t8 := int32(load32(m.memory[int64(uint32(v8))+28:]))
		t9 := v22
		v23 = t8
		v24 = t9 + v23
		if uint32(t7) >= uint32(v24) {
			goto l32
		}
		t10 := int32(load32(m.memory[int64(uint32(v8))+56:]))
		v5 = t10
		t11 := int32(m.memory[int64(uint32(v8))+64])
		v9 = t11
		t12 := int64(load64(m.memory[int64(uint32(v8))+48:]))
		v25 = t12
	l72:
		{
			t13 := int64(load32(m.memory[int64(uint32(v8))+148:]))
			v26 = i64_shl(i64(-1), t13) ^ i64(-1)
			t14 := int32(load32(m.memory[int64(uint32(v8))+60:]))
			v27 = t14
			t15 := int32(m.memory[int64(uint32(v8))+152])
			v28 = t15
		l40:
			v3 = int32(v25 & v26)
			v20 = i32(512)
			v29 = i32(1277688)
			switch v28 {
			default:
				goto l33
			case 1:
				v20 = i32(1332)
				v29 = v16
				goto l33
			case 2:
				v20 = i32(1332)
				v29 = v17
				goto l33
			case 3:
				v20 = i32(592)
				v29 = v18
			}
		l33:
			if uint32(v20) <= uint32(v3) {
				m.fn33(v3, v20, i32(1284808))
				panic("unreachable")
			}
			{
				t16 := v9 & i32(255)
				v3 = v29 + v3<<2
				t17 := int32(m.memory[int64(uint32(v3))+3])
				v20 = t17
				if uint32(t16) >= uint32(v20) {
					{
						t22 := int32(load16(m.memory[uint32(v3):]))
						v3 = t22
						if uint32(v3) < uint32(i32(16)) {
							t50 := v8
							v9 = v9 - v20
							m.memory[int64(uint32(t50))+64] = byte(v9)
							t51 := v8
							v25 = i64_shr_u(v25, int64(uint32(v20)))
							store64(m.memory[int64(uint32(t51))+48:], uint64(v25))
							if uint32(v21) < uint32(i32(320)) {
								store16(m.memory[uint32(v15+v21<<1):], uint16(v3))
								v21 = v21 + i32(1)
								goto l70
							}
							m.fn33(v21, i32(320), i32(1284920))
							panic("unreachable")
						}
						v29 = v9 & i32(255)
						switch v3 + i32(-16) {
						default:
							v28 = v20 + i32(7)
							if uint32(v28) > uint32(v29) {
								goto l47
							}
							v29 = v5
							v3 = v9
							goto l46
						l47:
							{
								if v5 == v27 {
									goto l39
								}
								t23 := v8
								v29 = v5 + i32(1)
								store32(m.memory[int64(uint32(t23))+56:], uint32(v29))
								t24 := v8
								v3 = v9 + i32(8)
								m.memory[int64(uint32(t24))+64] = byte(v3)
								t25 := int64(m.memory[uint32(v5)])
								t26 := v8
								v25 = i64_shl(t25, int64(uint32(v9))) | v25
								store64(m.memory[int64(uint32(t26))+48:], uint64(v25))
								v5 = v29
								v9 = v3
								if uint32(v28) > uint32(v3&i32(255)) {
									goto l47
								}
							}
						l46:
							t27 := v8
							v9 = v3 - v20 + i32(-7)
							m.memory[int64(uint32(t27))+64] = byte(v9)
							t28 := v8
							v30 = i64_shr_u(v25, int64(uint32(v20)))
							v25 = int64(uint64(v30) >> 7)
							store64(m.memory[int64(uint32(t28))+48:], uint64(v25))
							v5 = int32(v30)&i32(127) + i32(11)
							v3 = v5 + v21
							if uint32(v3) > uint32(v24) {
								store32(m.memory[int64(uint32(v8))+136:], uint32(i32(26)))
								store32(m.memory[int64(uint32(v8))+132:], uint32(i32(1053089)))
								goto l57
							}
							if uint32(v21) > uint32(i32(320)) {
								m.fn121(v21, i32(320), i32(320), i32(1284904))
								panic("unreachable")
							}
							t29 := v5
							v20 = i32(320) - v21
							if uint32(t29) > uint32(v20) {
								m.fn121(i32(0), v5, v20, i32(1284888))
								panic("unreachable")
							}
							v5 = v5 << 1
							if v5 != 0 {
								memory_zero(m.memory, uint32(v15+v21<<1), uint32(v5))
								goto l52
							}
							goto l52
						case 0:
							v28 = v20 + i32(2)
							if uint32(v28) > uint32(v29) {
								goto l53
							}
							v29 = v9
							goto l54
						l53:
							v3 = v5
						l55:
							{
								if v3 == v27 {
									goto l39
								}
								t30 := v8
								v5 = v3 + i32(1)
								store32(m.memory[int64(uint32(t30))+56:], uint32(v5))
								t31 := v8
								v29 = v9 + i32(8)
								m.memory[int64(uint32(t31))+64] = byte(v29)
								t32 := int64(m.memory[uint32(v3)])
								t33 := v8
								v25 = i64_shl(t32, int64(uint32(v9))) | v25
								store64(m.memory[int64(uint32(t33))+48:], uint64(v25))
								v3 = v5
								v9 = v29
								if uint32(v28) > uint32(v29&i32(255)) {
									goto l55
								}
							}
						l54:
							t34 := v8
							v9 = v29 - v20
							m.memory[int64(uint32(t34))+64] = byte(v9)
							t35 := v8
							v30 = i64_shr_u(v25, int64(uint32(v20)))
							store64(m.memory[int64(uint32(t35))+48:], uint64(v30))
							if v21 != 0 {
								v3 = v21 + i32(-1)
								if uint32(v21) > uint32(i32(320)) {
									m.fn33(v3, i32(320), i32(1284824))
									panic("unreachable")
								}
								t43 := int32(load16(m.memory[uint32(v15+v3<<1):]))
								v3 = t43
								t44 := v8
								v9 = v9 + i32(-2)
								m.memory[int64(uint32(t44))+64] = byte(v9)
								t45 := v8
								v25 = int64(uint64(v30) >> 2)
								store64(m.memory[int64(uint32(t45))+48:], uint64(v25))
								v27 = int32(v30) & i32(3)
								v29 = v27 + i32(3)
								if uint32(v29+v21) <= uint32(v24) {
									t46 := v29
									v20 = i32(320) - v21
									if uint32(t46) > uint32(v20) {
										m.fn121(i32(0), v29, v20, i32(1284840))
										panic("unreachable")
									}
									v20 = v15 + v21<<1
									store16(m.memory[int64(uint32(v20))+4:], uint16(v3))
									store16(m.memory[int64(uint32(v20))+2:], uint16(v3))
									store16(m.memory[uint32(v20):], uint16(v3))
									if v27 == 0 {
										goto l69
									}
									store16(m.memory[int64(uint32(v20))+6:], uint16(v3))
									if v29 == i32(4) {
										goto l69
									}
									store16(m.memory[int64(uint32(v20))+8:], uint16(v3))
									if v29 == i32(5) {
										goto l69
									}
									store16(m.memory[int64(uint32(v20))+10:], uint16(v3))
								l69:
									t47 := int32(load32(m.memory[int64(uint32(v8))+36:]))
									v21 = t47 + v29
									t48 := int32(load32(m.memory[int64(uint32(v8))+32:]))
									v22 = t48
									t49 := int32(load32(m.memory[int64(uint32(v8))+28:]))
									v23 = t49
									goto l70
								}
								store32(m.memory[int64(uint32(v8))+136:], uint32(i32(26)))
								store32(m.memory[int64(uint32(v8))+132:], uint32(i32(1053089)))
								goto l57
							}
							store32(m.memory[int64(uint32(v8))+136:], uint32(i32(26)))
							store32(m.memory[int64(uint32(v8))+132:], uint32(i32(1053089)))
							goto l57
						case 1:
							v3 = v20 + i32(3)
							if uint32(v3) > uint32(v29) {
								goto l60
							}
							v28 = v5
							v29 = v9
							goto l59
						l60:
							{
								if v5 == v27 {
									goto l39
								}
								t36 := v8
								v28 = v5 + i32(1)
								store32(m.memory[int64(uint32(t36))+56:], uint32(v28))
								t37 := v8
								v29 = v9 + i32(8)
								m.memory[int64(uint32(t37))+64] = byte(v29)
								t38 := int64(m.memory[uint32(v5)])
								t39 := v8
								v25 = i64_shl(t38, int64(uint32(v9))) | v25
								store64(m.memory[int64(uint32(t39))+48:], uint64(v25))
								v5 = v28
								v9 = v29
								if uint32(v3) > uint32(v29&i32(255)) {
									goto l60
								}
							}
						l59:
							v3 = i32(-3)
							t40 := v8
							v9 = v29 - v20 + i32(-3)
							m.memory[int64(uint32(t40))+64] = byte(v9)
							t41 := v8
							v30 = i64_shr_u(v25, int64(uint32(v20)))
							v25 = int64(uint64(v30) >> 3)
							store64(m.memory[int64(uint32(t41))+48:], uint64(v25))
							v5 = int32(v30)&i32(7) + i32(3)
							v20 = v5 + v21
							if uint32(v20) > uint32(v24) {
								store32(m.memory[int64(uint32(v8))+136:], uint32(i32(26)))
								store32(m.memory[int64(uint32(v8))+132:], uint32(i32(1053089)))
								v20 = i32(30)
								goto l30
							}
							if uint32(v21) > uint32(i32(320)) {
								m.fn121(v21, i32(320), i32(320), i32(1284872))
								panic("unreachable")
							}
							t42 := v5
							v3 = i32(320) - v21
							if uint32(t42) > uint32(v3) {
								m.fn121(i32(0), v5, v3, i32(1284856))
								panic("unreachable")
							}
							v5 = v5 << 1
							if v5 != 0 {
								memory_zero(m.memory, uint32(v15+v21<<1), uint32(v5))
								goto l65
							}
							goto l65
						}
					}
				l65:
					v5 = v28
					v21 = v20
					goto l70
				l52:
					v5 = v29
					v21 = v3
				l70:
					store32(m.memory[int64(uint32(v8))+36:], uint32(v21))
					t52 := v21
					v24 = v23 + v22
					if uint32(t52) < uint32(v24) {
						goto l72
					}
					goto l32
				}
				if v5 == v27 {
					goto l39
				}
				t18 := int64(m.memory[uint32(v5)])
				v30 = t18
				t19 := v8
				v3 = v9 + i32(8)
				m.memory[int64(uint32(t19))+64] = byte(v3)
				t20 := v8
				v5 = v5 + i32(1)
				store32(m.memory[int64(uint32(t20))+56:], uint32(v5))
				t21 := v8
				v25 = i64_shl(v30, int64(uint32(v9))) | v25
				store64(m.memory[int64(uint32(t21))+48:], uint64(v25))
				v9 = v3
				goto l40
			}
		}
	case 10:
		v3 = i32(-4)
		v20 = i32(10)
		goto l30
	case 9:
		v3 = i32(-2)
		v20 = i32(9)
		goto l30
	}
l39:
	v3 = i32(0)
	v20 = i32(26)
	goto l30
l32:
	{
		t53 := int32(load16(m.memory[int64(uint32(v8))+13700:]))
		if t53 != 0 {
			if uint32(v23) >= uint32(i32(321)) {
				m.fn121(i32(0), v23, i32(320), i32(1284792))
				panic("unreachable")
			}
			m.fn905(v7+i32(16), i32(1), v15, v23, v17, i32(1332), i32(10), v14)
			t54 := int32(load32(m.memory[int64(uint32(v7))+16:]))
			if t54 != 0 {
				store32(m.memory[int64(uint32(v8))+136:], uint32(i32(28)))
				store32(m.memory[int64(uint32(v8))+132:], uint32(i32(1052888)))
				goto l57
			}
			t55 := int32(load32(m.memory[int64(uint32(v7))+24:]))
			v9 = t55
			t56 := int32(load32(m.memory[int64(uint32(v7))+20:]))
			store32(m.memory[int64(uint32(v8))+148:], uint32(t56))
			m.memory[int64(uint32(v8))+152] = byte(i32(2))
			store32(m.memory[int64(uint32(v8))+40:], uint32(v9))
			t57 := int32(load32(m.memory[int64(uint32(v8))+28:]))
			v9 = t57
			if uint32(v9) > uint32(i32(320)) {
				m.fn121(v9, i32(320), i32(320), i32(1284776))
				panic("unreachable")
			}
			t58 := int32(load32(m.memory[int64(uint32(v8))+32:]))
			v5 = t58
			t59 := v5
			v3 = i32(320) - v9
			if uint32(t59) > uint32(v3) {
				m.fn121(i32(0), v5, v3, i32(1284760))
				panic("unreachable")
			}
			m.fn905(v7+i32(16), i32(2), v15+v9<<1, v5, v18, i32(592), i32(9), v14)
			{
				t60 := int32(load32(m.memory[int64(uint32(v7))+16:]))
				if t60 != 0 {
					store32(m.memory[int64(uint32(v8))+136:], uint32(i32(22)))
					store32(m.memory[int64(uint32(v8))+132:], uint32(i32(1052966)))
					goto l57
				}
				t61 := int32(load32(m.memory[int64(uint32(v7))+24:]))
				v9 = t61
				t62 := int32(load32(m.memory[int64(uint32(v7))+20:]))
				v5 = t62
				m.memory[int64(uint32(v8))+160] = byte(i32(3))
				store32(m.memory[int64(uint32(v8))+156:], uint32(v5))
				t63 := int32(load32(m.memory[int64(uint32(v8))+40:]))
				store32(m.memory[int64(uint32(v8))+40:], uint32(v9+t63))
				v9 = i32(17)
				t64 := int32(m.memory[int64(uint32(v8))+4])
				if t64 != i32(6) {
					goto l18
				}
				v3 = i32(0)
				v20 = i32(17)
				goto l30
			}
		}
		store32(m.memory[int64(uint32(v8))+136:], uint32(i32(37)))
		store32(m.memory[int64(uint32(v8))+132:], uint32(i32(1065321)))
		goto l57
	}
l26:
	{
		t666 := int32(load32(m.memory[int64(uint32(v8))+36:]))
		v9 = t666
		t667 := int32(load32(m.memory[int64(uint32(v8))+24:]))
		t668 := v9
		v29 = t667
		if uint32(t668) >= uint32(v29) {
			goto l275
		}
		p669 := i32(19)
		if uint32(v9) > uint32(i32(19)) {
			p669 = v9
		}
		v27 = p669
		t670 := int32(load32(m.memory[int64(uint32(v8))+56:]))
		v20 = t670
		t671 := int64(load64(m.memory[int64(uint32(v8))+48:]))
		v30 = t671
		t672 := int32(m.memory[int64(uint32(v8))+64])
		v5 = t672
		t673 := int32(load32(m.memory[int64(uint32(v8))+60:]))
		v21 = t673
	l280:
		{
			{
				if uint32(v5&i32(255)) <= uint32(i32(2)) {
					goto l276
				}
				v25 = v30
				v3 = v5
				goto l277
			l276:
				if v20 != v21 {
					goto l278
				}
				v3 = i32(0)
				v20 = i32(25)
				goto l30
			l278:
				t674 := v8
				v28 = v20 + i32(1)
				store32(m.memory[int64(uint32(t674))+56:], uint32(v28))
				t675 := v8
				v3 = v5 | i32(8)
				m.memory[int64(uint32(t675))+64] = byte(v3)
				t676 := int64(m.memory[uint32(v20)])
				t677 := v8
				v25 = i64_shl(t676, int64(uint32(v5))&i64(255)) | v30
				store64(m.memory[int64(uint32(t677))+48:], uint64(v25))
				v20 = v28
			}
		l277:
			if v27 == v9 {
				m.fn33(v27, i32(19), i32(1284744))
				panic("unreachable")
			}
			t678 := v8
			v5 = v3 + i32(-3)
			m.memory[int64(uint32(t678))+64] = byte(v5)
			t679 := v8
			v30 = int64(uint64(v25) >> 3)
			store64(m.memory[int64(uint32(t679))+48:], uint64(v30))
			t680 := v8
			v3 = v9 + i32(1)
			store32(m.memory[int64(uint32(t680))+36:], uint32(v3))
			t681 := int32(m.memory[uint32(v9+i32(1284724))])
			store16(m.memory[uint32(v15+t681<<1):], uint16(int32(v25)&i32(7)))
			v9 = v3
			if v29 != v3 {
				goto l280
			}
		}
		v9 = v29
		goto l275
	}
l275:
	if uint32(v9) > uint32(i32(18)) {
		goto l281
	}
	v5 = v9
	v3 = v9 & i32(3)
	if v3 == i32(3) {
		goto l282
	}
	v5 = i32(0)
l283:
	{
		t682 := int32(m.memory[uint32(v9+v5+i32(1284724))])
		store16(m.memory[uint32(v15+t682<<1):], uint16(i32(0)))
		t683 := v3
		v5 = v5 + i32(1)
		if t683^v5 != i32(3) {
			goto l283
		}
	}
	v5 = v9 + v5
l282:
	if uint32(v9) > uint32(i32(15)) {
		goto l284
	}
l285:
	{
		t684 := int32(m.memory[uint32(v5+i32(1284724))])
		store16(m.memory[uint32(v15+t684<<1):], uint16(i32(0)))
		t685 := int32(m.memory[uint32(v5+i32(1284725))])
		store16(m.memory[uint32(v15+t685<<1):], uint16(i32(0)))
		t686 := int32(m.memory[uint32(v5+i32(1284726))])
		store16(m.memory[uint32(v15+t686<<1):], uint16(i32(0)))
		t687 := int32(m.memory[uint32(v5+i32(1284727))])
		store16(m.memory[uint32(v15+t687<<1):], uint16(i32(0)))
		v5 = v5 + i32(4)
		if v5 != i32(19) {
			goto l285
		}
	}
l284:
	store32(m.memory[int64(uint32(v8))+36:], uint32(i32(19)))
l281:
	m.fn905(v7+i32(16), i32(0), v15, i32(19), v16, i32(1332), i32(7), v14)
	{
		t688 := int32(load32(m.memory[int64(uint32(v7))+16:]))
		if t688 != 0 {
			goto l286
		}
		t689 := int32(load32(m.memory[int64(uint32(v7))+20:]))
		v9 = t689
		t690 := int32(load32(m.memory[int64(uint32(v7))+24:]))
		v5 = t690
		m.memory[int64(uint32(v8))+152] = byte(i32(1))
		store32(m.memory[int64(uint32(v8))+40:], uint32(v5))
		store32(m.memory[int64(uint32(v8))+148:], uint32(v9))
		store32(m.memory[int64(uint32(v8))+36:], uint32(i32(0)))
		v9 = i32(26)
		goto l18
	}
l286:
	store32(m.memory[int64(uint32(v8))+136:], uint32(i32(25)))
	store32(m.memory[int64(uint32(v8))+132:], uint32(i32(1052916)))
l57:
	v3 = i32(-3)
	v20 = i32(30)
l30:
	m.memory[uint32(v8)] = byte(v20)
	t691 := int32(load32(m.memory[int64(uint32(v1))+32:]))
	v29 = t691
	t692 := int32(load32(m.memory[int64(uint32(v1))+60:]))
	t693 := v1
	v8 = t692
	t694 := int32(load32(m.memory[int64(uint32(v8))+56:]))
	v5 = t694
	store32(m.memory[int64(uint32(t693))+32:], uint32(v5))
	t695 := int32(load32(m.memory[int64(uint32(v8))+60:]))
	store32(m.memory[int64(uint32(v1))+36:], uint32(t695-v5))
	t696 := int32(load32(m.memory[int64(uint32(v8))+76:]))
	t697 := v1
	v9 = t696
	t698 := int32(load32(m.memory[int64(uint32(v8))+80:]))
	t699 := v9
	v20 = t698
	store32(m.memory[int64(uint32(t697))+48:], uint32(t699-v20))
	t700 := int32(load32(m.memory[int64(uint32(v8))+72:]))
	t701 := v1
	t702 := v20
	v27 = t700
	v28 = t702 + v27
	store32(m.memory[int64(uint32(t701))+44:], uint32(v28))
	t703 := int32(load32(m.memory[int64(uint32(v1))+40:]))
	store32(m.memory[int64(uint32(v1))+40:], uint32(v5-v29+t703))
	t704 := int32(load32(m.memory[int64(uint32(v8))+84:]))
	t705 := int32(load32(m.memory[int64(uint32(v8))+116:]))
	t706 := v8
	v9 = v20 - v9 + t705
	v15 = t704 + v9
	store32(m.memory[int64(uint32(t706))+84:], uint32(v15))
	store32(m.memory[int64(uint32(v1))+52:], uint32(v15))
	t707 := int32(load32(m.memory[int64(uint32(v8))+124:]))
	t708 := v1
	v16 = t707
	store32(m.memory[int64(uint32(t708))+80:], uint32(v16))
	{
		{
			t709 := int32(load32(m.memory[uint32(v8+i32(12)):]))
			v15 = t709
			if uint32(v15+i32(-64)) >= uint32(i32(-63)) {
				m.fn7(i32(1284212), i32(74), i32(1284288))
				panic("unreachable")
			}
			{
				if uint32(v15) > uint32(i32(64)) {
					goto l288
				}
				if v9 == 0 {
					goto l289
				}
				t710 := int32(m.memory[uint32(v8)])
				v19 = t710
				if uint32(v19) > uint32(i32(30)) {
					goto l288
				}
				v19 = i32_shl(i32(1), v19)
				if v19&i32(0x40000600) != 0 {
					goto l289
				}
				if v19&i32(67584) == 0 {
					goto l288
				}
				t711 := int32(m.memory[int64(uint32(v8))+4])
				if t711 == i32(4) {
					goto l289
				}
			}
		l288:
			if uint32(v9) > uint32(v20) {
				m.fn121(i32(0), v9, v20, i32(1285032))
				panic("unreachable")
			}
			t712 := int32(m.memory[int64(uint32(v8))+3])
			v18 = t712 & i32(4)
			t713 := int32(load32(m.memory[int64(uint32(v8))+120:]))
			v21 = t713
			{
				t714 := v9
				v20 = v15 + i32(-64)
				p715 := v20
				if uint32(v20) > uint32(v15) {
					p715 = i32(0)
				}
				v20 = p715
				if uint32(t714) >= uint32(v20) {
					t732 := v27
					v19 = v9 - v20
					v15 = t732 + v19
					{
						if v18 == 0 {
							if v20 == 0 {
								goto l307
							}
							t736 := int32(load32(m.memory[int64(uint32(v8))+8:]))
							memory_copy(m.memory, uint32(t736), uint32(v15), uint32(v20))
							goto l307
						}
						if v21 != 0 {
							t737 := int32(load32(m.memory[int64(uint32(v8))+128:]))
							t738 := m.fn907(t737, v27, v19)
							v27 = t738
							t739 := int32(load32(m.memory[int64(uint32(v8))+8:]))
							v19 = t739
							t740 := m.fn907(v27, v15, v20)
							store32(m.memory[int64(uint32(v8))+128:], uint32(t740))
							if v20 == 0 {
								goto l307
							}
							memory_copy(m.memory, uint32(v19), uint32(v15), uint32(v20))
							goto l307
						}
						t733 := m.fn906(v16, v27, v19)
						v27 = t733
						{
							if v20 == 0 {
								goto l306
							}
							t734 := int32(load32(m.memory[int64(uint32(v8))+8:]))
							memory_copy(m.memory, uint32(t734), uint32(v15), uint32(v20))
						}
					l306:
						t735 := m.fn906(v27, v15, v20)
						store32(m.memory[int64(uint32(v8))+124:], uint32(t735))
						goto l307
					}
				}
				t716 := int32(load32(m.memory[int64(uint32(v8))+20:]))
				t717 := v9
				t718 := v9
				t719 := v20
				v17 = t716
				v13 = t719 - v17
				p720 := v13
				if uint32(v9) < uint32(v13) {
					p720 = t718
				}
				v19 = p720
				v23 = t717 - v19
				v12 = v27 + v19
				{
					if v18 != 0 {
						if uint32(v15) < uint32(v17) {
							m.fn121(v17, v15, v15, i32(0x139900))
							panic("unreachable")
						}
						t723 := v19
						v15 = v15 - v17
						if uint32(t723) > uint32(v15) {
							m.fn121(i32(0), v19, v15, i32(1284336))
							panic("unreachable")
						}
						t724 := int32(load32(m.memory[int64(uint32(v8))+8:]))
						v18 = t724
						v15 = v18 + v17
						{
							if v21 != 0 {
								t728 := int32(load32(m.memory[int64(uint32(v8))+128:]))
								t729 := m.fn907(t728, v27, v19)
								store32(m.memory[int64(uint32(v8))+128:], uint32(t729))
								if v19 == 0 {
									goto l303
								}
								memory_copy(m.memory, uint32(v15), uint32(v27), uint32(v19))
							l303:
								if uint32(v9) <= uint32(v13) {
									goto l296
								}
								t730 := int32(load32(m.memory[int64(uint32(v8))+128:]))
								t731 := m.fn907(t730, v12, v23)
								store32(m.memory[int64(uint32(v8))+128:], uint32(t731))
								if v23 == 0 {
									goto l297
								}
								memory_copy(m.memory, uint32(v18), uint32(v12), uint32(v23))
								goto l297
							}
							if v19 == 0 {
								goto l301
							}
							memory_copy(m.memory, uint32(v15), uint32(v27), uint32(v19))
						l301:
							t725 := m.fn906(v16, v27, v19)
							t726 := v8
							v15 = t725
							store32(m.memory[int64(uint32(t726))+124:], uint32(v15))
							if uint32(v9) <= uint32(v13) {
								goto l296
							}
							if v23 == 0 {
								goto l302
							}
							memory_copy(m.memory, uint32(v18), uint32(v12), uint32(v23))
						l302:
							t727 := m.fn906(v15, v12, v23)
							store32(m.memory[int64(uint32(v8))+124:], uint32(t727))
							goto l297
						}
					}
					if uint32(v15) < uint32(v17) {
						m.fn121(v17, v15, v15, i32(1284320))
						panic("unreachable")
					}
					t721 := v19
					v15 = v15 - v17
					if uint32(t721) > uint32(v15) {
						m.fn121(i32(0), v19, v15, i32(1284304))
						panic("unreachable")
					}
					t722 := int32(load32(m.memory[int64(uint32(v8))+8:]))
					v15 = t722
					if v19 == 0 {
						goto l295
					}
					memory_copy(m.memory, uint32(v15+v17), uint32(v27), uint32(v19))
				l295:
					if uint32(v9) <= uint32(v13) {
						goto l296
					}
					if v23 == 0 {
						goto l297
					}
					memory_copy(m.memory, uint32(v15), uint32(v12), uint32(v23))
					goto l297
				}
			}
		}
	l307:
		store32(m.memory[int64(uint32(v8))+16:], uint32(v20))
		store32(m.memory[int64(uint32(v8))+20:], uint32(i32(0)))
		goto l289
	l297:
		store32(m.memory[int64(uint32(v8))+16:], uint32(v20))
		store32(m.memory[int64(uint32(v8))+20:], uint32(v23))
		goto l289
	l296:
		t741 := v8
		v15 = v19 + v17
		p742 := v15
		if v15 == v20 {
			p742 = i32(0)
		}
		store32(m.memory[int64(uint32(t741))+20:], uint32(p742))
		t743 := int32(load32(m.memory[int64(uint32(v8))+16:]))
		v15 = t743
		if uint32(v15) >= uint32(v20) {
			goto l289
		}
		store32(m.memory[int64(uint32(v8))+16:], uint32(v15+v19))
	}
l289:
	{
		t744 := int32(load32(m.memory[int64(uint32(v8))+132:]))
		v15 = t744
		if v15 == 0 {
			goto l308
		}
		{
			t745 := int32(load32(m.memory[int64(uint32(v8))+136:]))
			v20 = t745
			if v20 == 0 {
				goto l309
			}
			t746 := int32(m.memory[uint32(v15+v20+i32(-1))])
			if t746 == 0 {
				goto l310
			}
		}
	l309:
		m.fn7(i32(1284979), i32(37), i32(1285016))
		panic("unreachable")
	l310:
		store32(m.memory[int64(uint32(v1))+56:], uint32(v15))
	}
l308:
	t747 := int32(m.memory[int64(uint32(v8))+1])
	t748 := int32(m.memory[int64(uint32(v8))+64])
	v15 = t747<<6&i32(64) | t748
	{
		t749 := int32(m.memory[uint32(v8)])
		v8 = (t749 + i32(-12)) & i32(255)
		if uint32(v8) >= uint32(i32(6)) {
			goto l311
		}
		t750 := int32(load32(m.memory[int64(uint32(v8<<2))+1291320:]))
		v15 = t750 | v15
	}
l311:
	store32(m.memory[int64(uint32(v1))+76:], uint32(v15))
	t751 := int64(load64(m.memory[int64(uint32(v1))+16:]))
	t752 := v1
	v25 = t751 + int64(uint32(v5-v2))
	store64(m.memory[int64(uint32(t752))+16:], uint64(v25))
	t753 := int64(load64(m.memory[int64(uint32(v1))+24:]))
	t754 := v1
	v30 = t753 + int64(uint32(v28-v4))
	store64(m.memory[int64(uint32(t754))+24:], uint64(v30))
	{
		{
			if v3 != 0 {
				goto l312
			}
			var p755 int32
			if v6&i32(255) == i32(4) {
				p755 = 1
			}
			var p756 int32
			if v5 == v29 {
				p756 = 1
			}
			var p757 int32
			if v9 == 0 {
				p757 = 1
			}
			if p755|p756&p757 != 0 {
				goto l313
			}
		}
	l312:
		switch v3 + i32(5) {
		case 1, 2, 3:
			t762 := int64(load64(m.memory[uint32(v1):]))
			store64(m.memory[uint32(v1):], uint64(v25-v11+t762))
			t763 := int64(load64(m.memory[int64(uint32(v1))+8:]))
			store64(m.memory[int64(uint32(v1))+8:], uint64(v30-v10+t763))
			v8 = v3 + i32(-2)
			switch v8 {
			case 0:
				goto l320
			case 1:
				goto l319
			default:
				{
					{
						t764 := int32(load32(m.memory[int64(uint32(v1))+56:]))
						v8 = t764
						if v8 != 0 {
							goto l322
						}
						v1 = i32(0)
						goto l323
					}
				l322:
					t765 := m.fn1910(v8)
					m.fn14(v7+i32(16), v8, t765)
					t766 := int32(load32(m.memory[int64(uint32(v7))+20:]))
					t767 := int32(load32(m.memory[int64(uint32(v7))+16:]))
					p768 := t766
					if t767 != 0 {
						p768 = i32(0)
					}
					v1 = p768
					t769 := int32(load32(m.memory[int64(uint32(v7))+24:]))
					v8 = t769
				}
			l323:
				store32(m.memory[int64(uint32(v0))+8:], uint32(v8))
				store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
				store32(m.memory[uint32(v0):], uint32(i32(0)))
				goto l324
			}
		case 5:
			goto l316
		case 7:
			t760 := int64(load64(m.memory[uint32(v1):]))
			store64(m.memory[uint32(v1):], uint64(v25-v11+t760))
			t761 := int64(load64(m.memory[int64(uint32(v1))+8:]))
			store64(m.memory[int64(uint32(v1))+8:], uint64(v30-v10+t761))
			goto l320
		default:
			goto l313
		case 6:
			v3 = i32(2)
			goto l316
		case 4:
			m.fn28(i32(1284416), i32(147), i32(1284492))
			panic("unreachable")
		}
	l313:
		v3 = i32(1)
	l316:
		t758 := int64(load64(m.memory[uint32(v1):]))
		store64(m.memory[uint32(v1):], uint64(v25-v11+t758))
		t759 := int64(load64(m.memory[int64(uint32(v1))+8:]))
		store64(m.memory[int64(uint32(v1))+8:], uint64(v30-v10+t759))
		goto l319
	}
l320:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v16))
	store32(m.memory[uint32(v0):], uint32(i32(1)))
	goto l324
l319:
	store32(m.memory[uint32(v0):], uint32(i32(2)))
	m.memory[int64(uint32(v0))+4] = byte(v3)
l324:
	m.g0 = v7 + i32(32)
}
func (m *Module) fn266(v0, v1 int32) int32 {
	var v2, v3, v4 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	t1 := int32(m.memory[uint32(v0)])
	t2 := v2
	v3 = t1
	t3 := int32(m.memory[int64(uint32(v3&i32(1)))+1122873])
	m.memory[int64(uint32(t2))+15] = byte(t3)
	v4 = int32(uint32(v3) >> 1)
	v0 = int32(bits.LeadingZeros32(uint32(v4))) + i32(-24)
	{
		if v4 != 0 {
			goto l0
		}
		v4 = v2 + i32(8) + i32(7)
		goto l1
	l0:
		t4 := int32(m.memory[uint32(v4&i32(1)+i32(1122873))])
		m.memory[int64(uint32(v2))+14] = byte(t4)
		if v0 != i32(7) {
			goto l2
		}
		v4 = v2 + i32(8) + i32(6)
		goto l1
	l2:
		t5 := int32(m.memory[uint32(int32(uint32(v3)>>2)&i32(1)+i32(1122873))])
		m.memory[int64(uint32(v2))+13] = byte(t5)
		if v0 != i32(6) {
			goto l3
		}
		v4 = v2 + i32(8) + i32(5)
		goto l1
	l3:
		t6 := int32(m.memory[uint32(int32(uint32(v3)>>3)&i32(1)+i32(1122873))])
		m.memory[int64(uint32(v2))+12] = byte(t6)
		if v0 != i32(5) {
			goto l4
		}
		v4 = v2 + i32(8) + i32(4)
		goto l1
	l4:
		t7 := int32(m.memory[uint32(int32(uint32(v3)>>4)&i32(1)+i32(1122873))])
		m.memory[int64(uint32(v2))+11] = byte(t7)
		if v0 != i32(4) {
			goto l5
		}
		v4 = v2 + i32(8) + i32(3)
		goto l1
	l5:
		t8 := int32(m.memory[uint32(int32(uint32(v3)>>5)&i32(1)+i32(1122873))])
		m.memory[int64(uint32(v2))+10] = byte(t8)
		if v0 != i32(3) {
			goto l6
		}
		v4 = v2 + i32(8) + i32(2)
		goto l1
	l6:
		t9 := int32(m.memory[uint32(int32(uint32(v3)>>6)&i32(1)+i32(1122873))])
		m.memory[int64(uint32(v2))+9] = byte(t9)
		if v0 != i32(2) {
			goto l7
		}
		v4 = v2 + i32(8) + i32(1)
		goto l1
	l7:
		t10 := int32(m.memory[int64(uint32(int32(uint32(v3)>>7)))+1122873])
		m.memory[int64(uint32(v2))+8] = byte(t10)
		v4 = v2 + i32(8)
	}
l1:
	t11 := m.fn681(v1, i32(1), i32(1122875), i32(2), v4, (i32(9)-v0)&i32(255))
	v0 = t11
	m.g0 = v2 + i32(16)
	return v0
}
