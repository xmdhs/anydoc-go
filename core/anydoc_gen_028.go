package core

import (
	"math"
	"math/bits"
)

func (m *Module) fn1212(v0 int32) {
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
					m.fn766(v6 + (i32(0)-int32(uint32(int64(bits.TrailingZeros64(uint64(v5))))>>3))*i32(60) + i32(-48))
					v4 = v4 + i32(-1)
					v5 = (v5 + i64(-1)) & v5
					goto l4
				}
				v6 = v6 + i32(-480)
				t5 := int64(load64(m.memory[uint32(v0):]))
				v5 = (t5 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
				v0 = v0 + i32(8)
				goto l3
			}
		}
	l1:
		m.fn39(v1+i32(4), i32(60), i32(8), v2+i32(1))
		t6 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		t7 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t8 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		m.fn40(v3-t6, t7, t8)
	}
l0:
	m.g0 = v1 + i32(16)
}
func (m *Module) fn1213(v0, v1, v2, v3, v4 int32) {
	var v5, v6, v7, v8, v9, v10, v11, v12, v13, v14, v15, v16, v17, v18 int32
	var v19 int64
	t0 := m.g0
	v5 = t0 - i32(64)
	m.g0 = v5
	v6 = v4 + i32(8)
	v7 = i32(0)
l42:
	{
		t1 := v1
		v8 = v7 + i32(2)
		if uint32(t1) < uint32(v8) {
			goto l0
		}
		v9 = v0 + v7
		t2 := int32(load16(m.memory[uint32(v9):]))
		v10 = t2
		m.fn148(v5+i32(32), v8, v0, v1, i32(1072528))
		t3 := int32(load32(m.memory[int64(uint32(v5))+32:]))
		t4 := int32(load32(m.memory[int64(uint32(v5))+36:]))
		t5 := m.fn1430(v10, t3, t4)
		v7 = t5
		if uint32(v7) > uint32(v1-v8) {
			goto l0
		}
		v11 = v0 + v8
		{
			{
				v10 = v10 & i32(0xffff)
				switch v10 + i32(-26182) {
				case 0:
					if uint32(v7) < uint32(i32(4)) {
						goto l2
					}
					t61 := int32(load32(m.memory[uint32(v11):]))
					t62 := v3
					v10 = t61
					if uint32(t62) < uint32(v10) {
						goto l2
					}
					v11 = v3 - v10
					if uint32(v11) < uint32(i32(2)) {
						goto l2
					}
					t63 := v11 + i32(-2)
					v10 = v2 + v10
					t64 := int32(load16(m.memory[uint32(v10):]))
					v11 = t64
					if uint32(t63) < uint32(v11) {
						goto l2
					}
					m.fn1213(v10+i32(2), v11, i32(1), i32(0), v4)
					goto l2
				case 1, 2:
					goto l2
				case 4:
					if uint32(v7) < uint32(i32(4)) {
						goto l2
					}
					t69 := int32(load32(m.memory[uint32(v4):]))
					v9 = t69
					store32(m.memory[uint32(v4):], uint32(i32(1)))
					t70 := int32(load32(m.memory[int64(uint32(v4))+4:]))
					t71 := v4
					v14 = t70
					t72 := int32(load32(m.memory[uint32(v11):]))
					t73 := v14
					v10 = t72
					v11 = t73 + v10
					t74 := v11>>31 ^ i32(-0x80000000)
					t75 := v11
					var p76 int32
					if v10 < i32(0) {
						p76 = 1
					}
					var p77 int32
					if v11 < v14 {
						p77 = 1
					}
					p78 := t75
					if p76^p77 != 0 {
						p78 = t74
					}
					p79 := v10
					if v9 != 0 {
						p79 = p78
					}
					store32(m.memory[int64(uint32(t71))+4:], uint32(p79))
					goto l2
				default:
					switch v10 + i32(-9238) {
					default:
						switch v10 + i32(-9291) {
						case 0:
							{
								if v7 != 0 {
									goto l24
								}
								v10 = i32(0)
								goto l25
							l24:
								t24 := int32(m.memory[uint32(v11)])
								var p25 int32
								if t24 != i32(0) {
									p25 = 1
								}
								v10 = p25
							}
						l25:
							m.memory[int64(uint32(v4))+46] = byte(v10)
							goto l2
						case 1:
							{
								if v7 != 0 {
									goto l26
								}
								v10 = i32(0)
								goto l27
							l26:
								t26 := int32(m.memory[uint32(v11)])
								var p27 int32
								if t26 != i32(0) {
									p27 = 1
								}
								v10 = p27
							}
						l27:
							m.memory[int64(uint32(v4))+47] = byte(v10)
							goto l2
						default:
							if v10 == i32(9738) {
								{
									if v7 == 0 {
										goto l21
									}
									t19 := int32(m.memory[uint32(v11)])
									v10 = t19
								}
							l21:
								m.memory[int64(uint32(v4))+41] = byte(v10)
								t20 := v4
								var p21 int32
								if v7 != i32(0) {
									p21 = 1
								}
								m.memory[int64(uint32(t20))+40] = byte(p21)
								goto l2
							}
							if v10 == i32(9792) {
								if v7 == 0 {
									goto l2
								}
								t65 := int32(m.memory[uint32(v11)])
								t66 := v4
								v10 = t65
								m.memory[int64(uint32(t66))+43] = byte(v10 + i32(1))
								t67 := v4
								var p68 int32
								if uint32(v10&i32(255)) < uint32(i32(9)) {
									p68 = 1
								}
								m.memory[int64(uint32(t67))+42] = byte(p68)
								goto l2
							}
							if v10 == i32(13316) {
								if v7 == 0 {
									goto l39
								}
								t58 := int32(m.memory[uint32(v11)])
								v10 = t58
								t59 := int32(load32(m.memory[uint32(v6):]))
								if t59 == i32(-1) {
									goto l40
								}
								var p60 int32
								if v10 != i32(0) {
									p60 = 1
								}
								v10 = p60
								goto l41
							}
							if v10 == i32(17931) {
								v10 = i32(0)
								{
									if uint32(v7) < uint32(i32(2)) {
										goto l22
									}
									t22 := int32(load16(m.memory[uint32(v11):]))
									v9 = t22
									v10 = i32(1)
								}
							l22:
								store16(m.memory[int64(uint32(v4))+38:], uint16(v9))
								store16(m.memory[int64(uint32(v4))+36:], uint16(v10))
								goto l2
							}
							if v10 == i32(54792) {
								if uint32(v7) < uint32(i32(3)) {
									goto l2
								}
								t28 := int32(m.memory[int64(uint32(v11))+2])
								v12 = t28
								if uint32(v12) > uint32(i32(63)) {
									goto l2
								}
								v10 = i32(1)
								t29 := v5 + i32(24)
								v13 = v12 + i32(1)
								m.fn59(t29, v13, i32(2), i32(2))
								v14 = i32(0)
								store32(m.memory[int64(uint32(v5))+48:], uint32(i32(0)))
								t30 := int32(load32(m.memory[int64(uint32(v5))+28:]))
								t31 := v5
								v15 = t30
								store32(m.memory[int64(uint32(t31))+44:], uint32(v15))
								t32 := int32(load32(m.memory[int64(uint32(v5))+24:]))
								store32(m.memory[int64(uint32(v5))+40:], uint32(t32))
								m.memory[int64(uint32(v5))+60] = byte(i32(0))
								store32(m.memory[int64(uint32(v5))+56:], uint32(v12))
								store32(m.memory[int64(uint32(v5))+52:], uint32(i32(0)))
							l31:
								{
									m.fn254(v5+i32(16), v5+i32(52))
									t33 := int32(load32(m.memory[int64(uint32(v5))+16:]))
									if t33 != i32(1) {
										m.fn59(v5+i32(8), v12, i32(1), i32(4))
										v14 = i32(0)
										store32(m.memory[int64(uint32(v5))+60:], uint32(i32(0)))
										t39 := int32(load32(m.memory[int64(uint32(v5))+12:]))
										t40 := v5
										v17 = t39
										store32(m.memory[int64(uint32(t40))+56:], uint32(v17))
										t41 := int32(load32(m.memory[int64(uint32(v5))+8:]))
										t42 := v5
										v10 = t41
										store32(m.memory[int64(uint32(t42))+52:], uint32(v10))
										{
											if uint32(v10) >= uint32(v12) {
												goto l32
											}
											m.fn62(v5+i32(52), i32(0), v12, i32(1), i32(4))
											t43 := int32(load32(m.memory[int64(uint32(v5))+60:]))
											v14 = t43
											t44 := int32(load32(m.memory[int64(uint32(v5))+56:]))
											v17 = t44
										}
									l32:
										v11 = v17 + v14<<2
										p45 := i32(1)
										if uint32(v12) > uint32(i32(1)) {
											p45 = v12
										}
										v16 = p45
										v10 = v16 + i32(-1)
									l36:
										if v10 != 0 {
											store32(m.memory[uint32(v11):], uint32(i32(0)))
											v10 = v10 + i32(-1)
											v11 = v11 + i32(4)
											goto l36
										}
										v18 = v14 + v16
										if v12 != 0 {
											goto l34
										}
										v18 = v18 + i32(-1)
										goto l35
									l34:
										store32(m.memory[uint32(v11):], uint32(i32(0)))
									l35:
										v14 = v18 << 2
										v11 = v13<<1 + i32(3)
										v16 = v7 - v12<<1 + i32(-5)
										t46 := int32(load32(m.memory[int64(uint32(v5))+52:]))
										v13 = t46
										v10 = v17
									l38:
										{
											if v14 == 0 {
												goto l37
											}
											if uint32(v7) < uint32(v11) {
												goto l37
											}
											if uint32(v16) < uint32(i32(2)) {
												goto l37
											}
											t47 := int32(load16(m.memory[uint32(v9+v11+i32(2)):]))
											t48 := v10
											v15 = t47
											v12 = v15 & i32(3)
											var p49 int32
											if v12 == i32(1) {
												p49 = 1
											}
											m.memory[int64(uint32(t48))+1] = byte(p49)
											t50 := v10
											var p51 int32
											if uint32(v12) > uint32(i32(1)) {
												p51 = 1
											}
											m.memory[uint32(t50)] = byte(p51)
											t52 := v10
											v15 = int32(uint32(v15)>>5) & i32(3)
											var p53 int32
											if v15 == i32(1) {
												p53 = 1
											}
											m.memory[int64(uint32(t52))+3] = byte(p53)
											t54 := v10
											var p55 int32
											if v15 == i32(3) {
												p55 = 1
											}
											m.memory[int64(uint32(t54))+2] = byte(p55)
											v16 = v16 + i32(-20)
											v11 = v11 + i32(20)
											v14 = v14 + i32(-4)
											v10 = v10 + i32(4)
											goto l38
										}
									}
									t34 := int32(load32(m.memory[int64(uint32(v5))+20:]))
									t35 := v7
									v16 = t34<<1 + i32(3)
									if uint32(t35) < uint32(v16) {
										goto l29
									}
									if uint32(v7-v16) < uint32(i32(2)) {
										goto l29
									}
									t36 := int32(load16(m.memory[uint32(v11+v16):]))
									v16 = t36
									{
										t37 := int32(load32(m.memory[int64(uint32(v5))+40:]))
										if v10+i32(-1) != t37 {
											goto l30
										}
										m.fn1153(v5 + i32(40))
										t38 := int32(load32(m.memory[int64(uint32(v5))+44:]))
										v15 = t38
									}
								l30:
									store16(m.memory[uint32(v15+v14):], uint16(v16))
									v14 = v14 + i32(2)
									store32(m.memory[int64(uint32(v5))+48:], uint32(v10))
									v10 = v10 + i32(1)
									goto l31
								}
							l29:
								t56 := int32(load32(m.memory[int64(uint32(v5))+40:]))
								t57 := int32(load32(m.memory[int64(uint32(v5))+44:]))
								m.fn768(t56, t57)
								goto l2
							}
							if v10 != i32(54827) {
								goto l2
							}
							if uint32(v7) < uint32(i32(3)) {
								goto l2
							}
							t6 := int32(load32(m.memory[uint32(v6):]))
							if t6 == i32(-1) {
								goto l2
							}
							t7 := int32(load32(m.memory[int64(uint32(v4))+28:]))
							t8 := int32(m.memory[int64(uint32(v11))+1])
							v10 = t8
							if uint32(t7) <= uint32(v10) {
								goto l2
							}
							t9 := int32(load32(m.memory[int64(uint32(v4))+24:]))
							v10 = t9 + v10<<2
							t10 := int32(m.memory[int64(uint32(v11))+2])
							t11 := v10
							v11 = t10 & i32(255)
							var p12 int32
							if v11 == i32(1) {
								p12 = 1
							}
							m.memory[int64(uint32(t11))+3] = byte(p12)
							t13 := v10
							var p14 int32
							if v11 == i32(3) {
								p14 = 1
							}
							m.memory[int64(uint32(t13))+2] = byte(p14)
							goto l2
						}
					case 0:
						{
							if v7 != 0 {
								goto l17
							}
							v10 = i32(0)
							goto l18
						l17:
							t15 := int32(m.memory[uint32(v11)])
							var p16 int32
							if t15 != i32(0) {
								p16 = 1
							}
							v10 = p16
						}
					l18:
						m.memory[int64(uint32(v4))+44] = byte(v10)
						goto l2
					case 1:
						{
							if v7 != 0 {
								goto l19
							}
							v10 = i32(0)
							goto l20
						l19:
							t17 := int32(m.memory[uint32(v11)])
							var p18 int32
							if t17 != i32(0) {
								p18 = 1
							}
							v10 = p18
						}
					l20:
						m.memory[int64(uint32(v4))+45] = byte(v10)
						goto l2
					}
				case 3:
					v10 = i32(0)
					{
						if uint32(v7) < uint32(i32(4)) {
							goto l23
						}
						t23 := int32(load32(m.memory[uint32(v11):]))
						v9 = t23
						v10 = i32(1)
					}
				l23:
					store32(m.memory[int64(uint32(v4))+4:], uint32(v9))
					store32(m.memory[uint32(v4):], uint32(v10))
					goto l2
				}
			l37:
				t80 := int32(load32(m.memory[int64(uint32(v5))+40:]))
				v10 = t80
				if v10 == i32(-1) {
					goto l2
				}
				t81 := int64(load64(m.memory[int64(uint32(v5))+44:]))
				v19 = t81
				t82 := int32(m.memory[int64(uint32(v4))+32])
				v11 = t82
				t83 := int32(load32(m.memory[int64(uint32(v4))+8:]))
				v9 = t83
				m.fn766(v6)
				store32(m.memory[int64(uint32(v4))+28:], uint32(v18))
				store32(m.memory[int64(uint32(v4))+24:], uint32(v17))
				store32(m.memory[int64(uint32(v4))+20:], uint32(v13))
				store64(m.memory[int64(uint32(v4))+12:], uint64(v19))
				store32(m.memory[int64(uint32(v4))+8:], uint32(v10))
				t84 := v4
				t85 := v11
				var p86 int32
				if v9 != i32(-1) {
					p86 = 1
				}
				m.memory[int64(uint32(t84))+32] = byte(t85 & p86)
				goto l2
			}
		l39:
			v10 = i32(0)
			t87 := int32(load32(m.memory[uint32(v6):]))
			if t87 == i32(-1) {
				goto l2
			}
		}
	l41:
		m.memory[int64(uint32(v4))+32] = byte(v10)
		goto l2
	l40:
		if v10 == 0 {
			goto l2
		}
		m.fn766(v6)
		m.memory[int64(uint32(v4))+32] = byte(i32(1))
		store64(m.memory[int64(uint32(v4))+24:], uint64(i64(1)))
		store64(m.memory[int64(uint32(v4))+16:], uint64(i64(0)))
		store64(m.memory[int64(uint32(v4))+8:], uint64(i64(0x200000000)))
	l2:
		v7 = v7 + v8
		goto l42
	}
l0:
	m.g0 = v5 + i32(64)
}
func (m *Module) fn1214(v0, v1, v2 int32) {
	var v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14, v15, v16, v17, v18, v19, v20, v21, v22, v23, v24 int32
	var v25 int64
	t0 := m.g0
	v3 = t0 - i32(64)
	m.g0 = v3
	t1 := int32(load32(m.memory[uint32(v2):]))
	t2 := v2
	t3 := v1
	v4 = t1 & i32(1)
	p4 := t3
	if v4 != 0 {
		p4 = t2
	}
	t5 := int32(load32(m.memory[int64(uint32(p4))+4:]))
	v5 = t5
	t6 := int32(load16(m.memory[int64(uint32(v2))+36:]))
	t7 := v2
	t8 := v1
	v6 = t6 & i32(1)
	p9 := t8
	if v6 != 0 {
		p9 = t7
	}
	t10 := int32(load16(m.memory[int64(uint32(p9))+38:]))
	v7 = t10
	t11 := int32(m.memory[int64(uint32(v2))+40])
	t12 := v2
	t13 := v1
	v8 = t11
	p14 := t13
	if v8&i32(1) != 0 {
		p14 = t12
	}
	t15 := int32(m.memory[int64(uint32(p14))+41])
	v9 = t15
	t16 := int32(m.memory[int64(uint32(v2))+42])
	t17 := v1
	t18 := v2
	v10 = t16
	var p19 int32
	if v10 == i32(2) {
		p19 = 1
	}
	v11 = p19
	p20 := t18
	if v11 != 0 {
		p20 = t17
	}
	t21 := int32(m.memory[int64(uint32(p20))+43])
	v12 = t21
	t22 := int32(load32(m.memory[uint32(v1):]))
	v13 = t22
	t23 := int32(m.memory[int64(uint32(v1))+46])
	v14 = t23
	t24 := int32(m.memory[int64(uint32(v2))+46])
	v15 = t24
	t25 := int32(m.memory[int64(uint32(v1))+47])
	v16 = t25
	t26 := int32(m.memory[int64(uint32(v2))+47])
	v17 = t26
	t27 := int32(load16(m.memory[int64(uint32(v1))+36:]))
	v18 = t27
	t28 := int32(m.memory[int64(uint32(v1))+40])
	v19 = t28
	t29 := int32(m.memory[int64(uint32(v1))+42])
	v20 = t29
	t30 := int32(m.memory[int64(uint32(v1))+44])
	v21 = t30
	t31 := int32(m.memory[int64(uint32(v2))+44])
	v22 = t31
	t32 := int32(m.memory[int64(uint32(v1))+45])
	v23 = t32
	t33 := int32(m.memory[int64(uint32(v2))+45])
	v24 = t33
	t34 := int32(load32(m.memory[int64(uint32(v2))+32:]))
	store32(m.memory[int64(uint32(v3))+24:], uint32(t34))
	t35 := int64(load64(m.memory[int64(uint32(v2))+24:]))
	store64(m.memory[int64(uint32(v3))+16:], uint64(t35))
	t36 := int64(load64(m.memory[int64(uint32(v2))+16:]))
	store64(m.memory[int64(uint32(v3))+8:], uint64(t36))
	t37 := int64(load64(m.memory[int64(uint32(v2))+8:]))
	t38 := v3
	v25 = t37
	store64(m.memory[uint32(t38):], uint64(v25))
	t39 := int32(load32(m.memory[int64(uint32(v1))+32:]))
	store32(m.memory[int64(uint32(v3))+56:], uint32(t39))
	t40 := int64(load64(m.memory[int64(uint32(v1))+24:]))
	store64(m.memory[int64(uint32(v3))+48:], uint64(t40))
	t41 := int64(load64(m.memory[int64(uint32(v1))+16:]))
	store64(m.memory[int64(uint32(v3))+40:], uint64(t41))
	t42 := int64(load64(m.memory[int64(uint32(v1))+8:]))
	store64(m.memory[int64(uint32(v3))+32:], uint64(t42))
	t43 := v0
	t44 := v3 + i32(32)
	t45 := v3
	var p46 int32
	if int32(v25) == i32(-1) {
		p46 = 1
	}
	v1 = p46
	p47 := t45
	if v1 != 0 {
		p47 = t44
	}
	v2 = p47
	t48 := int64(load64(m.memory[uint32(v2):]))
	store64(m.memory[int64(uint32(t43))+8:], uint64(t48))
	t49 := int64(load64(m.memory[int64(uint32(v2))+8:]))
	store64(m.memory[int64(uint32(v0))+16:], uint64(t49))
	t50 := int64(load64(m.memory[int64(uint32(v2))+16:]))
	store64(m.memory[int64(uint32(v0))+24:], uint64(t50))
	t51 := int32(load32(m.memory[int64(uint32(v2))+24:]))
	store32(m.memory[int64(uint32(v0))+32:], uint32(t51))
	p52 := v3 + i32(32)
	if v1 != 0 {
		p52 = v3
	}
	m.fn766(p52)
	t54 := v0
	p53 := v24
	if v24 == i32(2) {
		p53 = v23
	}
	m.memory[int64(uint32(t54))+45] = byte(p53)
	t56 := v0
	p55 := v22
	if v22 == i32(2) {
		p55 = v21
	}
	m.memory[int64(uint32(t56))+44] = byte(p55)
	m.memory[int64(uint32(v0))+43] = byte(v12)
	t58 := v0
	p57 := v10
	if v11 != 0 {
		p57 = v20
	}
	m.memory[int64(uint32(t58))+42] = byte(p57)
	m.memory[int64(uint32(v0))+41] = byte(v9)
	m.memory[int64(uint32(v0))+40] = byte(v19 | v8)
	store16(m.memory[int64(uint32(v0))+38:], uint16(v7))
	t60 := v0
	p59 := v18
	if v6 != 0 {
		p59 = i32(1)
	}
	store16(m.memory[int64(uint32(t60))+36:], uint16(p59))
	t62 := v0
	p61 := v17
	if v17 == i32(2) {
		p61 = v16
	}
	m.memory[int64(uint32(t62))+47] = byte(p61)
	t64 := v0
	p63 := v15
	if v15 == i32(2) {
		p63 = v14
	}
	m.memory[int64(uint32(t64))+46] = byte(p63)
	store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
	t66 := v0
	p65 := v13
	if v4 != 0 {
		p65 = i32(1)
	}
	store32(m.memory[uint32(t66):], uint32(p65))
	m.g0 = v3 + i32(64)
}
func (m *Module) fn1215(v0, v1, v2, v3 int32) int32 {
	var v4, v5, v6, v7, v8, v9, v10 int32
	t0 := m.g0
	v4 = t0 - i32(16)
	m.g0 = v4
	m.memory[int64(uint32(v4))+15] = byte(v2)
	m.memory[int64(uint32(v4))+14] = byte(int32(uint32(v2) >> 8))
	m.memory[int64(uint32(v4))+13] = byte(int32(uint32(v2) >> 16))
	v5 = i32(0)
	var p1 int32
	if v3&i32(256) != i32(0) {
		p1 = 1
	}
	v6 = p1
	var p2 int32
	if v3&i32(65536) != i32(0) {
		p2 = 1
	}
	v7 = p2
	var _ int32
l8:
	{
		{
			t4 := v1
			v8 = v5 + i32(2)
			if uint32(t4) < uint32(v8) {
				goto l0
			}
			if uint32(v5) >= uint32(v1) {
				m.fn158(v5, v1, i32(1072496))
				panic("unreachable")
			}
			v9 = v5 + i32(1)
			if uint32(v9) >= uint32(v1) {
				m.fn158(v9, v1, i32(1072512))
				panic("unreachable")
			}
			t5 := int32(m.memory[uint32(v0+v5)])
			v5 = t5
			t6 := int32(m.memory[uint32(v0+v9)])
			v9 = t6
			m.fn148(v4, v8, v0, v1, i32(1072528))
			v9 = v5 | v9<<8
			t7 := int32(load32(m.memory[uint32(v4):]))
			t8 := int32(load32(m.memory[int64(uint32(v4))+4:]))
			t9 := m.fn1430(v9, t7, t8)
			v5 = t9
			if uint32(v5) > uint32(v1-v8) {
				goto l0
			}
			v10 = v0 + v8
			switch v9&i32(0xffff) + i32(-2101) {
			default:
				goto l6
			case 0:
				t10 := m.fn1431(v10, v5, v3&i32(1))
				v9 = t10 & i32(255)
				if v9 == i32(2) {
					goto l6
				}
				v10 = v4 + i32(15)
				goto l7
			case 1:
				t11 := m.fn1431(v10, v5, v6)
				v9 = t11 & i32(255)
				if v9 == i32(2) {
					goto l6
				}
				v10 = v4 + i32(14)
				goto l7
			case 2:
				t12 := m.fn1431(v10, v5, v7)
				v9 = t12 & i32(255)
				if v9 == i32(2) {
					goto l6
				}
				v10 = v4 + i32(13)
				goto l7
			}
		}
	l0:
		t13 := int32(m.memory[int64(uint32(v4))+13])
		v1 = t13
		t14 := int32(m.memory[int64(uint32(v4))+14])
		v5 = t14
		t15 := int32(m.memory[int64(uint32(v4))+15])
		v8 = t15
		m.g0 = v4 + i32(16)
		return v8 | (v2&i32(-0x1000000) | v1<<16 | v5<<8)
	}
l7:
	m.memory[uint32(v10)] = byte(v9)
l6:
	v5 = v5 + v8
	goto l8
}
func (m *Module) fn1216(v0 int32) {
	t0 := int32(load32(m.memory[uint32(v0):]))
	if t0 == i32(2) {
		return
	}
	m.fn766(v0 + i32(8))
}
func (m *Module) fn1217(v0, v1 int32) {
	m.fn76(v1, v0)
}
func (m *Module) fn1218(v0, v1, v2, v3 int32) {
	var v4, v5, v6, v7, v8 int32
	var v9 int64
	var v10, v11, v12, v13, v14, v15, v16, v17, v18 int32
	t0 := m.g0
	v4 = t0 - i32(96)
	m.g0 = v4
	{
		if uint32(v2) >= uint32(v3) {
			goto l0
		}
		store32(m.memory[uint32(v0):], uint32(i32(2)))
		goto l1
	l0:
		v5 = v2 - v3
		if uint32(v5) > uint32(i32(27)) {
			goto l2
		}
		store32(m.memory[uint32(v0):], uint32(i32(2)))
		goto l1
	l2:
		v6 = v1 + v3
		t1 := int32(m.memory[int64(uint32(v6))+5])
		v7 = t1
		t2 := int32(m.memory[int64(uint32(v6))+4])
		v8 = t2
		t3 := int64(load32(m.memory[uint32(v6):]))
		v9 = t3
		v2 = i32(0)
		m.memory[int64(uint32(v4))+80] = byte(i32(0))
		store32(m.memory[int64(uint32(v4))+76:], uint32(v6+i32(15)))
		store32(m.memory[int64(uint32(v4))+72:], uint32(v6+i32(6)))
		m.fn915(v4+i32(24), v4+i32(72))
		v10 = i32(1)
		v11 = i32(0)
		{
			t4 := int32(m.memory[int64(uint32(v4))+24])
			if t4 == 0 {
				goto l3
			}
			t5 := int32(m.memory[int64(uint32(v4))+25])
			v2 = t5
			m.fn59(v4+i32(16), i32(8), i32(1), i32(1))
			t6 := int32(load32(m.memory[int64(uint32(v4))+16:]))
			v1 = t6
			t7 := int32(load32(m.memory[int64(uint32(v4))+20:]))
			v10 = t7
			m.memory[uint32(v10)] = byte(v2)
			store32(m.memory[int64(uint32(v4))+92:], uint32(i32(1)))
			store32(m.memory[int64(uint32(v4))+88:], uint32(v10))
			store32(m.memory[int64(uint32(v4))+84:], uint32(v1))
			t8 := int32(load32(m.memory[int64(uint32(v4))+80:]))
			store32(m.memory[int64(uint32(v4))+40:], uint32(t8))
			t9 := int64(load64(m.memory[int64(uint32(v4))+72:]))
			store64(m.memory[int64(uint32(v4))+32:], uint64(t9))
			v2 = i32(1)
		l6:
			{
				m.fn915(v4+i32(8), v4+i32(32))
				t10 := int32(m.memory[int64(uint32(v4))+8])
				if t10 == 0 {
					goto l4
				}
				t11 := int32(m.memory[int64(uint32(v4))+9])
				v1 = t11
				{
					t12 := int32(load32(m.memory[int64(uint32(v4))+84:]))
					if v2 != t12 {
						goto l5
					}
					m.fn47(v4+i32(84), i32(1))
					t13 := int32(load32(m.memory[int64(uint32(v4))+88:]))
					v10 = t13
				}
			l5:
				m.memory[uint32(v10+v2)] = byte(v1)
				t14 := v4
				v2 = v2 + i32(1)
				store32(m.memory[int64(uint32(t14))+92:], uint32(v2))
				goto l6
			}
		l4:
			t15 := int32(load32(m.memory[int64(uint32(v4))+84:]))
			v11 = t15
		}
	l3:
		{
			{
				t16 := int32(m.memory[int64(uint32(v6))+25])
				t17 := v5
				v12 = t16
				t18 := int32(m.memory[int64(uint32(v6))+24])
				t19 := v12
				v13 = t18
				v1 = t19 + v13 + i32(28)
				if uint32(t17) < uint32(v1) {
					goto l7
				}
				v14 = i32(1)
				if uint32(v5-v1) > uint32(i32(1)) {
					goto l8
				}
			}
		l7:
			store32(m.memory[uint32(v0):], uint32(i32(2)))
			goto l9
		l8:
			v15 = v6 + v1
			switch v8 + i32(-1) {
			default:
				if v8 == i32(23) {
					goto l15
				}
				if v8 != i32(255) {
					goto l16
				}
				fallthrough
			case 2:
				v14 = v8
				goto l16
			case 0:
				v14 = i32(5)
				goto l16
			case 1:
				v14 = i32(4)
				goto l16
			case 3:
				v14 = i32(2)
				goto l16
			}
		l15:
			v14 = i32(0)
		l16:
			t20 := int32(m.memory[int64(uint32(v6))+26])
			v16 = t20
			v17 = v1 + i32(2)
			t21 := int32(load16(m.memory[uint32(v15):]))
			v15 = t21
			v8 = i32(0)
			store32(m.memory[int64(uint32(v4))+92:], uint32(i32(0)))
			store64(m.memory[int64(uint32(v4))+84:], uint64(i64(0x400000000)))
			if uint32((v14+i32(-1))&i32(255)) > uint32(i32(253)) {
				goto l17
			}
			v12 = v5 - v12 - v13 + i32(-30)
			v1 = v17
		l23:
			if v15 == v8 {
				goto l17
			}
			{
				if uint32(v5) < uint32(v1) {
					goto l18
				}
				if uint32(v12) < uint32(i32(2)) {
					goto l18
				}
				t22 := int32(load16(m.memory[uint32(v6+v1):]))
				v13 = t22
				if uint32(v13) >= uint32(i32(9)) {
					goto l19
				}
				t23 := m.fn1432(v10, v2, v8+i32(1))
				if t23 == 0 {
					goto l19
				}
				store32(m.memory[int64(uint32(v4))+32:], uint32(i32(-1)))
				m.memory[int64(uint32(v4))+36] = byte(v13)
				m.fn1321(v4+i32(84), v4+i32(32))
				goto l20
			}
		l18:
			store32(m.memory[uint32(v0):], uint32(i32(2)))
			m.fn763(v4 + i32(84))
			goto l9
		l19:
			if uint32(v13^i32(-1058816)) < uint32(i32(-1112064)) {
				goto l20
			}
			if uint32(v13) < uint32(i32(32)) {
				goto l20
			}
			if uint32((v13+i32(-160))&i32(0xffff)) >= uint32(i32(65503)) {
				goto l20
			}
			{
				t24 := int32(load32(m.memory[int64(uint32(v4))+92:]))
				v18 = t24
				if v18 == 0 {
					goto l21
				}
				t25 := int32(load32(m.memory[int64(uint32(v4))+88:]))
				v18 = t25 + v18*i32(12) + i32(-12)
				if v18 == 0 {
					goto l21
				}
				t26 := int32(load32(m.memory[uint32(v18):]))
				if t26 != i32(-1) {
					goto l22
				}
			}
		l21:
			m.fn1072(v4+i32(32), v13)
			m.fn1321(v4+i32(84), v4+i32(32))
			goto l20
		l22:
			m.fn74(v18, v13)
		l20:
			v8 = v8 + i32(1)
			v12 = v12 + i32(-2)
			v1 = v1 + i32(2)
			goto l23
		l17:
			t27 := int32(load32(m.memory[int64(uint32(v4))+92:]))
			store32(m.memory[int64(uint32(v4))+48:], uint32(t27))
			t28 := int64(load64(m.memory[int64(uint32(v4))+84:]))
			store64(m.memory[int64(uint32(v4))+40:], uint64(t28))
			m.memory[int64(uint32(v4))+64] = byte(v14)
			store64(m.memory[int64(uint32(v4))+56:], uint64(v9))
			store32(m.memory[int64(uint32(v4))+36:], uint32(v16))
			m.memory[int64(uint32(v4))+52] = byte(int32(uint32(v7)>>2) & i32(1))
			store32(m.memory[int64(uint32(v4))+32:], uint32(int32(uint32(v7)>>3)&i32(1)))
			v2 = v17 + v3 + v15<<1
			if uint32(v2) >= uint32(v3) {
				goto l24
			}
			store32(m.memory[uint32(v0):], uint32(i32(2)))
			m.fn763(v4 + i32(40))
			goto l9
		l24:
			memory_copy(m.memory, uint32(v0), uint32(v4+i32(32)), uint32(i32(40)))
			store32(m.memory[int64(uint32(v0))+40:], uint32(v2))
		}
	l9:
		m.fn16(v11, v10)
	}
l1:
	m.g0 = v4 + i32(96)
}
func (m *Module) fn1219(v0, v1 int32) {
	var v2, v3, v4 int32
	var v5 int64
	t0 := int32(load32(m.memory[uint32(v1):]))
	v2 = t0
	t1 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	v3 = t1
	t2 := int32(m.memory[int64(uint32(v1))+32])
	v4 = t2
	t3 := int64(load64(m.memory[int64(uint32(v1))+24:]))
	v5 = t3
	m.fn1138(v0+i32(8), v1+i32(8))
	store64(m.memory[int64(uint32(v0))+24:], uint64(v5))
	m.memory[int64(uint32(v0))+32] = byte(v4)
	store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
	store32(m.memory[uint32(v0):], uint32(v2))
}
func (m *Module) fn1220(v0, v1 int32) {
	var v2, v3, v4 int32
	t0 := m.g0
	v2 = t0 - i32(512)
	m.g0 = v2
	v3 = i32(0)
l1:
	if v3 == i32(360) {
		goto l0
	}
	v4 = v2 + i32(152) + v3
	store32(m.memory[uint32(v4):], uint32(i32(0)))
	m.memory[uint32(v4+i32(32))] = byte(i32(0))
	store64(m.memory[uint32(v4+i32(24)):], uint64(i64(1)))
	m.memory[uint32(v4+i32(20))] = byte(i32(0))
	store32(m.memory[uint32(v4+i32(16)):], uint32(i32(0)))
	store64(m.memory[uint32(v4+i32(8)):], uint64(i64(0x400000000)))
	v3 = v3 + i32(40)
	goto l1
l0:
	v4 = i32(0)
l3:
	if v4 == i32(144) {
		goto l2
	}
	store64(m.memory[uint32(v2+i32(8)+v4):], uint64(i64(0)))
	v4 = v4 + i32(16)
	goto l3
l2:
	store32(m.memory[int64(uint32(v0))+504:], uint32(v1^i32(-1)|i32(-65536)))
	memory_copy(m.memory, uint32(v0+i32(144)), uint32(v2+i32(152)), uint32(i32(360)))
	memory_copy(m.memory, uint32(v0), uint32(v2+i32(8)), uint32(i32(144)))
	m.g0 = v2 + i32(512)
}
func (m *Module) fn1221(v0 int32) {
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
					m.fn761(v6 + (i32(0)-int32(uint32(int64(bits.TrailingZeros64(uint64(v5))))>>3))*i32(368) + i32(-360))
					v4 = v4 + i32(-1)
					v5 = (v5 + i64(-1)) & v5
					goto l4
				}
				v6 = v6 + i32(-2944)
				t5 := int64(load64(m.memory[uint32(v0):]))
				v5 = (t5 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
				v0 = v0 + i32(8)
				goto l3
			}
		}
	l1:
		m.fn39(v1+i32(4), i32(368), i32(8), v2+i32(1))
		t6 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		t7 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t8 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		m.fn40(v3-t6, t7, t8)
	}
l0:
	m.g0 = v1 + i32(16)
}
func (m *Module) fn1222(v0, v1, v2, v3, v4, v5, v6 int32) {
	var v7, v8, v9 int32
	t0 := m.g0
	v7 = t0 - i32(32)
	m.g0 = v7
	v8 = i32(0)
	v9 = i32(0)
	{
		if uint32(v2) < uint32(v5) {
			goto l0
		}
		v9 = i32(0)
		if uint32(v2-v5) < uint32(i32(4)) {
			goto l0
		}
		t1 := int32(load32(m.memory[uint32(v1+v5):]))
		v9 = t1
	}
l0:
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
		v8 = t3
	}
l1:
	{
		{
			{
				v2 = v9 + v8
				p4 := v2
				if uint32(v2) < uint32(v9) {
					p4 = i32(-1)
				}
				v2 = p4
				if uint32(v2) > uint32(v4) {
					store64(m.memory[int64(uint32(v0))+8:], uint64(i64(0)))
					store64(m.memory[uint32(v0):], uint64(i64(0x400000000)))
					goto l6
				}
				if uint32(v8) < uint32(i32(8)) {
					store64(m.memory[int64(uint32(v0))+8:], uint64(i64(0)))
					store64(m.memory[uint32(v0):], uint64(i64(0x400000000)))
					goto l6
				}
				if v6 != 0 {
					goto l4
				}
				v4 = int32(uint32(v8)>>2) + i32(-1)
				goto l5
			}
		l4:
			t5 := int32(uint32(v8+i32(-4)) / uint32(v6|i32(4)))
			v4 = t5
		}
	l5:
		v5 = v2 - v9
		v8 = v3 + v9
		m.fn485(v7+i32(8), v4+i32(1))
		m.memory[int64(uint32(v7))+28] = byte(i32(0))
		store32(m.memory[int64(uint32(v7))+24:], uint32(v4))
		store32(m.memory[int64(uint32(v7))+20:], uint32(i32(0)))
	l9:
		{
			m.fn254(v7, v7+i32(20))
			t6 := int32(load32(m.memory[uint32(v7):]))
			if t6 != i32(1) {
				goto l7
			}
			v9 = i32(0)
			{
				t7 := int32(load32(m.memory[int64(uint32(v7))+4:]))
				t8 := v5
				v2 = t7 << 2
				if uint32(t8) < uint32(v2) {
					goto l8
				}
				if uint32(v5-v2) < uint32(i32(4)) {
					goto l8
				}
				t9 := int32(load32(m.memory[uint32(v8+v2):]))
				v9 = t9
			}
		l8:
			m.fn584(v7+i32(8), v9)
			goto l9
		}
	l7:
		t10 := int32(load32(m.memory[int64(uint32(v7))+16:]))
		store32(m.memory[int64(uint32(v0))+8:], uint32(t10))
		t11 := int64(load64(m.memory[int64(uint32(v7))+8:]))
		store64(m.memory[uint32(v0):], uint64(t11))
		store32(m.memory[int64(uint32(v0))+12:], uint32(v4))
	}
l6:
	m.g0 = v7 + i32(32)
}
func (m *Module) fn1223(v0, v1, v2 int32) int32 {
	var v3, v4, v5 int32
	if v1 != 0 {
		v3 = i32(0)
	l2:
		{
			if uint32(v1) < uint32(i32(2)) {
				t2 := int32(load32(m.memory[uint32(v0+v3<<2):]))
				t3 := v3
				var p4 int32
				if uint32(t2) < uint32(v2) {
					p4 = 1
				}
				return t3 + p4
			}
			v4 = int32(uint32(v1) >> 1)
			v5 = v4 + v3
			t0 := int32(load32(m.memory[uint32(v0+v5<<2):]))
			p1 := v3
			if uint32(t0) < uint32(v2) {
				p1 = v5
			}
			v3 = p1
			v1 = v1 - v4
			goto l2
		}
	}
	return i32(0)
}
func (m *Module) fn1224(v0, v1 int32) {
	var v2, v3, v4 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		v2 = t0
		if uint32(v2) < uint32(i32(2)) {
			goto l0
		}
		t1 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		v3 = t1
		if uint32(v2) < uint32(i32(21)) {
			v4 = v2 * i32(72)
			v2 = i32(72)
		l2:
			if v4 == v2 {
				goto l0
			}
			m.fn986(v3, v3+v2)
			v2 = v2 + i32(72)
			goto l2
		}
		m.fn977(v3, v2)
		goto l0
	}
l0:
	t2 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	store32(m.memory[int64(uint32(v0))+8:], uint32(t2))
	t3 := int64(load64(m.memory[uint32(v1):]))
	store64(m.memory[uint32(v0):], uint64(t3))
}
func (m *Module) fn1225(v0 int32) {
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
	store64(m.memory[int64(uint32(v0))+40:], uint64(i64(4)))
	store64(m.memory[int64(uint32(v0))+32:], uint64(i64(0)))
	t3 := int64(load64(m.memory[int64(uint32(i32(0)))+1286256:]))
	store64(m.memory[uint32(v0):], uint64(t3))
	t4 := int64(load64(m.memory[int64(uint32(i32(0)))+1286264:]))
	store64(m.memory[int64(uint32(v0))+8:], uint64(t4))
	store64(m.memory[int64(uint32(v0))+24:], uint64(v3))
	store64(m.memory[int64(uint32(v0))+16:], uint64(v2))
	m.g0 = v1 + i32(16)
}
func (m *Module) fn1226(v0, v1, v2, v3 int32) {
	var v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14, v15, v16, v17, v18, v19, v20, v21, v22, v23, v24, v25, v26, v27, v28 int32
	var v29 int64
	var v30 int32
	var v31 int64
	var v32, v33 int32
	var v34 int64
	var v35, v36, v37, v38, v39, v40, v41 int32
	t0 := m.g0
	v4 = t0 - i32(928)
	m.g0 = v4
	store32(m.memory[int64(uint32(v4))+68:], uint32(i32(0)))
	store64(m.memory[int64(uint32(v4))+60:], uint64(i64(0x800000000)))
	store32(m.memory[int64(uint32(v4))+80:], uint32(i32(0)))
	store64(m.memory[int64(uint32(v4))+72:], uint64(i64(0x800000000)))
	store32(m.memory[int64(uint32(v4))+84:], uint32(i32(0)))
	store32(m.memory[int64(uint32(v4))+108:], uint32(i32(0)))
	store64(m.memory[int64(uint32(v4))+100:], uint64(i64(0x800000000)))
	store32(m.memory[int64(uint32(v4))+112:], uint32(i32(0)))
	store32(m.memory[int64(uint32(v4))+136:], uint32(i32(0)))
	store64(m.memory[int64(uint32(v4))+128:], uint64(i64(0x400000000)))
	store32(m.memory[int64(uint32(v4))+148:], uint32(i32(0)))
	store64(m.memory[int64(uint32(v4))+140:], uint64(i64(0x400000000)))
	store64(m.memory[int64(uint32(v4))+184:], uint64(i64(0)))
	store64(m.memory[int64(uint32(v4))+176:], uint64(i64(0x100000000)))
	store64(m.memory[int64(uint32(v4))+168:], uint64(i64(4)))
	store64(m.memory[int64(uint32(v4))+160:], uint64(i64(0)))
	store64(m.memory[int64(uint32(v4))+152:], uint64(i64(0x400000000)))
	v5 = v1 + i32(200)
	v6 = v1 + i32(136)
	v7 = v1 + i32(272)
	v8 = v4 + i32(352) + i32(12)
	v9 = v4 + i32(192) + i32(8)
	v10 = v4 + i32(352) | i32(4)
	v11 = v4 + i32(496)
	v12 = v4 + i32(864) + i32(28)
	v13 = v4 + i32(256) + i32(8)
	v14 = v4 + i32(352) + i32(4)
	v15 = v4 + i32(884)
	v16 = v4 + i32(152) + i32(12)
	v17 = v1 + i32(96)
	v18 = v1 + i32(92)
	v19 = v1 + i32(108)
	v20 = v1 + i32(104)
	v21 = v1 + i32(336)
	v22 = v1 + i32(332)
	v23 = v1 + i32(348)
	v24 = v1 + i32(344)
	v25 = v1 + i32(192)
l56:
	{
		t1 := int32(load32(m.memory[uint32(v17):]))
		t2 := v2
		v26 = t1
		p3 := v3
		if uint32(v26) < uint32(v3) {
			p3 = v26
		}
		if uint32(t2) < uint32(p3) {
			t8 := int32(load32(m.memory[uint32(v18):]))
			t9 := m.fn622(t8, v26, v2, i32(1078128))
			t10 := int32(load32(m.memory[uint32(t9):]))
			v27 = t10
			t11 := int32(load32(m.memory[uint32(v20):]))
			t12 := int32(load32(m.memory[uint32(v19):]))
			t13 := m.fn622(t11, t12, v2, i32(1078144))
			t14 := int32(load32(m.memory[uint32(t13):]))
			v28 = t14
			{
				t15 := int32(load32(m.memory[int64(uint32(v1))+180:]))
				if t15 == 0 {
					goto l3
				}
				t16 := int64(load64(m.memory[int64(uint32(v1))+184:]))
				t17 := int64(load64(m.memory[uint32(v25):]))
				t18 := m.fn66(t16, t17, v2)
				v29 = t18
				t19 := int32(load32(m.memory[int64(uint32(v1))+172:]))
				v30 = t19
				v26 = v30 & int32(v29)
				v31 = int64(uint64(v29)>>25) & i64(127) * i64(72340172838076673)
				t20 := int32(load32(m.memory[int64(uint32(v1))+168:]))
				v32 = t20
				v33 = i32(0)
			l8:
				{
					t21 := int64(load64(m.memory[uint32(v32+v26):]))
					v34 = t21
					v29 = v34 ^ v31
					v29 = (v29 ^ i64(-1)) & (v29 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
					{
					l6:
						if v29 == 0 {
							if !(v34&(v34<<1)&i64(-0x7f7f7f7f7f7f7f80) == 0) {
								goto l3
							}
							t26 := v26
							v33 = v33 + i32(8)
							v26 = (t26 + v33) & v30
							goto l8
						}
						{
							t22 := v2
							v35 = v32 - (int32(uint32(int64(bits.TrailingZeros64(uint64(v29))))>>3)+v26)&v30<<4
							t23 := int32(load32(m.memory[uint32(v35+i32(-16)):]))
							if t22 == t23 {
								t24 := int32(load32(m.memory[uint32(v35+i32(-8)):]))
								t25 := int32(load32(m.memory[uint32(v35+i32(-4)):]))
								m.fn31(v14, t24, t25)
								store32(m.memory[int64(uint32(v4))+352:], uint32(i32(7)))
								m.fn1435(v4+i32(152), v4+i32(352))
								goto l7
							}
							v29 = (v29 + i64(-1)) & v29
							goto l6
						}
					}
				}
			}
		l3:
			switch v27 + i32(-1) {
			case 1, 4, 7, 30:
				goto l7
			default:
				if uint32(v27) < uint32(i32(32)) {
					goto l7
				}
				if uint32(v27+i32(-127)) < uint32(i32(33)) {
					goto l7
				}
				t27 := m.fn1436(v1, v28, v2)
				m.fn1437(v4+i32(152), v27, t27)
				goto l7
			case 6, 11, 12, 13:
				{
					t28 := int32(load32(m.memory[uint32(v24):]))
					t29 := int32(load32(m.memory[uint32(v23):]))
					t30 := m.fn1438(t28, t29, v28)
					v32 = t30
					if v32 == 0 {
						goto l18
					}
					t31 := int32(load16(m.memory[int64(uint32(v32))+60:]))
					v26 = t31
					m.fn1429(v4+i32(256), v32)
					goto l19
				}
			l18:
				m.memory[int64(uint32(v4))+298] = byte(i32(2))
				v26 = i32(0)
				m.memory[int64(uint32(v4))+296] = byte(i32(0))
				store16(m.memory[int64(uint32(v4))+292:], uint16(i32(0)))
				store32(m.memory[int64(uint32(v4))+300:], uint32(i32(33686018)))
				store32(m.memory[int64(uint32(v4))+256:], uint32(i32(0)))
				store32(m.memory[int64(uint32(v4))+264:], uint32(i32(-1)))
			l19:
				t32 := m.fn1439(v1, v26)
				m.fn1429(v4+i32(352), t32)
				m.fn1214(v4+i32(864), v4+i32(352), v4+i32(256))
				{
					t33 := int32(load32(m.memory[int64(uint32(v1))+132:]))
					if uint32(v2) >= uint32(t33) {
						goto l20
					}
					t34 := int32(load32(m.memory[int64(uint32(v1))+128:]))
					t35 := int32(load32(m.memory[uint32(t34+v2<<2):]))
					m.fn1440(v4+i32(24), v1, t35)
					t36 := int32(load32(m.memory[int64(uint32(v4))+24:]))
					v32 = t36
					if v32 == 0 {
						goto l20
					}
					t37 := int32(load32(m.memory[int64(uint32(v4))+28:]))
					v35 = t37
					m.memory[int64(uint32(v4))+346] = byte(i32(2))
					m.memory[int64(uint32(v4))+344] = byte(i32(0))
					store16(m.memory[int64(uint32(v4))+340:], uint16(i32(0)))
					store32(m.memory[int64(uint32(v4))+348:], uint32(i32(33686018)))
					store32(m.memory[int64(uint32(v4))+304:], uint32(i32(0)))
					store32(m.memory[int64(uint32(v4))+312:], uint32(i32(-1)))
					m.fn1213(v32, v35, i32(1), i32(0), v4+i32(304))
					m.fn1214(v4+i32(352), v4+i32(864), v4+i32(304))
					memory_copy(m.memory, uint32(v4+i32(864)), uint32(v4+i32(352)), uint32(i32(48)))
				}
			l20:
				store16(m.memory[int64(uint32(v4))+240:], uint16(v26))
				memory_copy(m.memory, uint32(v4+i32(192)), uint32(v4+i32(864)), uint32(i32(48)))
				memory_copy(m.memory, uint32(v4+i32(352)), uint32(v4+i32(152)), uint32(i32(40)))
				store64(m.memory[int64(uint32(v4))+184:], uint64(i64(0)))
				store64(m.memory[int64(uint32(v4))+176:], uint64(i64(0x100000000)))
				store64(m.memory[int64(uint32(v4))+168:], uint64(i64(4)))
				store64(m.memory[int64(uint32(v4))+160:], uint64(i64(0)))
				store64(m.memory[int64(uint32(v4))+152:], uint64(i64(0x400000000)))
				m.fn1433(v4+i32(244), v4+i32(352))
				var p38 int32
				if v27 == i32(7) {
					p38 = 1
				}
				v32 = p38
				if v32 != 0 {
					goto l21
				}
				t39 := int32(m.memory[int64(uint32(v4))+236])
				if t39&i32(1) != 0 {
					goto l21
				}
				m.fn1434(v4+i32(352), v4+i32(60), v4+i32(140), v4+i32(128), v4+i32(100))
				t40 := int32(load32(m.memory[int64(uint32(v4))+352:]))
				v32 = t40
				if v32 == i32(-1) {
					t109 := int32(load32(m.memory[int64(uint32(v4))+248:]))
					v32 = t109
					t110 := int32(load32(m.memory[int64(uint32(v4))+252:]))
					v35 = t110
					{
						t111 := m.fn1439(v1, v26)
						v26 = t111
						t112 := int32(m.memory[int64(uint32(v26))+54])
						v27 = t112
						if v27 == i32(2) {
							{
								t113 := m.fn23(v32, v35)
								if t113 != 0 {
									m.fn1333(v4+i32(84), v4+i32(60))
									m.fn1351(v4+i32(60), v4+i32(72))
									m.fn894(v4 + i32(244))
									goto l40
								}
								{
									t114 := int32(m.memory[int64(uint32(v4))+234])
									t115 := int32(m.memory[int64(uint32(v26))+52])
									t116 := t114 & i32(1)
									v32 = t115
									if t116|v32 == 0 {
										t148 := int32(load16(m.memory[int64(uint32(v4))+230:]))
										t149 := int32(load16(m.memory[int64(uint32(v4))+228:]))
										p150 := i32(0)
										if t149 != 0 {
											p150 = t148
										}
										v32 = p150
										v26 = v32 & i32(0xffff)
										if v26 == 0 {
											goto l46
										}
										if v26 == i32(63489) {
											goto l46
										}
										t151 := int32(m.memory[int64(uint32(v4))+233])
										t152 := int32(m.memory[int64(uint32(v4))+232])
										p153 := i32(0)
										if t152 != 0 {
											p153 = t151
										}
										v26 = p153
										t154 := m.fn1446(v6, v32)
										v27 = t154
										v35 = v27
										if v27 != 0 {
											goto l47
										}
										m.fn1220(v4+i32(352), v32)
										v35 = v4 + i32(352)
									l47:
										{
											t156 := v35
											p155 := i32(8)
											if uint32(v26) < uint32(i32(8)) {
												p155 = v26
											}
											t157 := int32(m.memory[int64(uint32(t156+p155*i32(40)))+176])
											v30 = t157
											if v30 == i32(255) {
												if v27 != 0 {
													goto l46
												}
												m.fn761(v11)
												goto l46
											}
											{
												if v30 != 0 {
													goto l49
												}
												store32(m.memory[int64(uint32(v4))+264:], uint32(i32(-1)))
												v29 = i64(0)
												goto l50
											l49:
												m.fn1182(v4+i32(8), v5, i32(1078176))
												t158 := int32(load32(m.memory[int64(uint32(v4))+12:]))
												v28 = t158
												t159 := int32(load32(m.memory[int64(uint32(v4))+8:]))
												m.fn1447(v4+i32(256), t159, v32, v35, v26)
												t160 := int32(load32(m.memory[uint32(v28):]))
												store32(m.memory[uint32(v28):], uint32(t160+i32(1)))
												t161 := int64(load64(m.memory[int64(uint32(v4))+256:]))
												v29 = t161
											}
										l50:
											t162 := int32(load32(m.memory[int64(uint32(v13))+8:]))
											store32(m.memory[int64(uint32(v12))+8:], uint32(t162))
											t163 := int64(load64(m.memory[uint32(v13):]))
											store64(m.memory[uint32(v12):], uint64(t163))
											m.fn1333(v4+i32(84), v4+i32(60))
											t164 := int64(load32(m.memory[int64(uint32(v35))+504:]))
											v31 = t164
											t165 := m.fn113(i32(8), i32(32))
											v32 = t165
											store32(m.memory[uint32(v32):], uint32(i32(-0x80000000)))
											t166 := int64(load64(m.memory[int64(uint32(v4))+244:]))
											store64(m.memory[int64(uint32(v32))+4:], uint64(t166))
											t167 := int32(load32(m.memory[int64(uint32(v4))+252:]))
											store32(m.memory[int64(uint32(v32))+12:], uint32(t167))
											store32(m.memory[int64(uint32(v4))+888:], uint32(v26))
											store32(m.memory[int64(uint32(v4))+912:], uint32(i32(1)))
											store32(m.memory[int64(uint32(v4))+908:], uint32(v32))
											store32(m.memory[int64(uint32(v4))+904:], uint32(i32(1)))
											store64(m.memory[int64(uint32(v4))+880:], uint64(v29))
											m.memory[int64(uint32(v4))+872] = byte(v30)
											store64(m.memory[int64(uint32(v4))+864:], uint64(v31))
											m.fn1369(v4+i32(72), v4+i32(864))
											if v27 != 0 {
												goto l40
											}
											m.fn761(v11)
											goto l40
										}
									}
									t117 := int32(m.memory[int64(uint32(v26))+53])
									t118 := int32(m.memory[int64(uint32(v4))+235])
									p119 := t118
									if v32 != 0 {
										p119 = t117
									}
									v32 = p119
									m.fn1333(v4+i32(84), v4+i32(60))
									m.fn1351(v4+i32(60), v4+i32(72))
									t120 := int64(load64(m.memory[int64(uint32(v4))+244:]))
									store64(m.memory[int64(uint32(v4))+256:], uint64(t120))
									t121 := int32(load32(m.memory[int64(uint32(v4))+252:]))
									t122 := v4
									v35 = t121
									store32(m.memory[int64(uint32(t122))+264:], uint32(v35))
									t123 := int32(load32(m.memory[int64(uint32(v4))+260:]))
									t124 := int32(load32(m.memory[int64(uint32(v26))+48:]))
									m.fn1368(t123, v35, t124)
									{
										t125 := int32(load16(m.memory[int64(uint32(v4))+230:]))
										t126 := int32(load16(m.memory[int64(uint32(v4))+228:]))
										p127 := i32(0)
										if t126 != 0 {
											p127 = t125
										}
										v35 = p127
										v26 = v35 & i32(0xffff)
										if v26 == 0 {
											goto l43
										}
										if v26 == i32(63489) {
											goto l43
										}
										t128 := int32(m.memory[int64(uint32(v4))+232])
										v27 = t128
										t129 := int32(m.memory[int64(uint32(v4))+233])
										v30 = t129
										t130 := m.fn1446(v6, v35)
										v26 = t130
										if v26 == 0 {
											goto l43
										}
										t132 := v26
										p131 := i32(0)
										if v27 != 0 {
											p131 = v30
										}
										v27 = p131
										p133 := i32(8)
										if uint32(v27) < uint32(i32(8)) {
											p133 = v27
										}
										t134 := int32(m.memory[int64(uint32(t132+p133*i32(40)))+176])
										v28 = t134
										if uint32((v28+i32(-1))&i32(255)) > uint32(i32(253)) {
											goto l43
										}
										m.fn1182(v4+i32(16), v5, i32(1078160))
										t135 := int32(load32(m.memory[int64(uint32(v4))+20:]))
										v30 = t135
										t136 := int32(load32(m.memory[int64(uint32(v4))+16:]))
										m.fn1447(v4+i32(352), t136, v35, v26, v27)
										t137 := int32(load32(m.memory[uint32(v30):]))
										store32(m.memory[uint32(v30):], uint32(t137+i32(1)))
										{
											{
												t138 := int32(load32(m.memory[int64(uint32(v4))+360:]))
												v26 = t138
												if v26 == i32(-1) {
													goto l44
												}
												t139 := int64(load64(m.memory[int64(uint32(v4))+364:]))
												store64(m.memory[int64(uint32(v4))+356:], uint64(t139))
												store32(m.memory[int64(uint32(v4))+352:], uint32(v26))
												goto l45
											}
										l44:
											t140 := int64(load64(m.memory[int64(uint32(v4))+352:]))
											m.fn800(v4+i32(352), v28, t140)
										}
									l45:
										store32(m.memory[int64(uint32(v4))+924:], uint32(i32(25)))
										store32(m.memory[int64(uint32(v4))+920:], uint32(v4+i32(352)))
										m.fn73(v4+i32(864), i32(1070105), v4+i32(920))
										t141 := int32(load32(m.memory[int64(uint32(v4))+352:]))
										t142 := int32(load32(m.memory[int64(uint32(v4))+356:]))
										m.fn16(t141, t142)
										t143 := int32(load32(m.memory[int64(uint32(v4))+864:]))
										if t143 == i32(-1) {
											goto l43
										}
										t144 := int32(load32(m.memory[int64(uint32(v4))+872:]))
										store32(m.memory[int64(uint32(v14))+8:], uint32(t144))
										t145 := int64(load64(m.memory[int64(uint32(v4))+864:]))
										store64(m.memory[uint32(v14):], uint64(t145))
										store32(m.memory[int64(uint32(v4))+368:], uint32(i32(0)))
										store32(m.memory[int64(uint32(v4))+352:], uint32(i32(3)))
										m.fn1163(v4+i32(256), v4+i32(352))
									}
								l43:
									t146 := int32(load32(m.memory[int64(uint32(v4))+264:]))
									store32(m.memory[int64(uint32(v4))+360:], uint32(t146))
									t147 := int64(load64(m.memory[int64(uint32(v4))+256:]))
									store64(m.memory[int64(uint32(v4))+352:], uint64(t147))
									m.memory[int64(uint32(v4))+376] = byte(v32)
									store32(m.memory[int64(uint32(v4))+364:], uint32(i32(-1)))
									m.fn338(v4+i32(60), v4+i32(352))
									goto l40
								}
							}
						l46:
							m.fn1333(v4+i32(84), v4+i32(60))
							m.fn1351(v4+i32(60), v4+i32(72))
							t168 := int32(load32(m.memory[int64(uint32(v4))+252:]))
							store32(m.memory[int64(uint32(v10))+8:], uint32(t168))
							t169 := int64(load64(m.memory[int64(uint32(v4))+244:]))
							store64(m.memory[uint32(v10):], uint64(t169))
							store32(m.memory[int64(uint32(v4))+352:], uint32(i32(-0x80000000)))
							m.fn338(v4+i32(60), v4+i32(352))
							goto l40
						}
						m.fn1351(v4+i32(60), v4+i32(72))
						m.fn1445(v4+i32(84), v27&i32(1), v4+i32(244), v4+i32(60))
						goto l40
					}
				}
				t41 := int32(load32(m.memory[int64(uint32(v4))+372:]))
				store32(m.memory[int64(uint32(v0))+20:], uint32(t41))
				t42 := int64(load64(m.memory[int64(uint32(v4))+364:]))
				store64(m.memory[int64(uint32(v0))+12:], uint64(t42))
				t43 := int64(load64(m.memory[int64(uint32(v4))+356:]))
				store64(m.memory[int64(uint32(v0))+4:], uint64(t43))
				store32(m.memory[uint32(v0):], uint32(v32))
				m.fn894(v4 + i32(244))
				m.fn766(v9)
				v2 = i32(0)
				v1 = i32(1)
				goto l23
			case 10:
				store32(m.memory[int64(uint32(v4))+352:], uint32(i32(8)))
				m.fn1435(v4+i32(152), v4+i32(352))
				goto l7
			case 18:
				m.fn1441(v4 + i32(152))
				m.memory[int64(uint32(v4))+376] = byte(i32(0))
				store32(m.memory[int64(uint32(v4))+360:], uint32(i32(0)))
				store64(m.memory[int64(uint32(v4))+352:], uint64(i64(0x100000000)))
				store32(m.memory[int64(uint32(v4))+372:], uint32(i32(0)))
				store64(m.memory[int64(uint32(v4))+364:], uint64(i64(0x400000000)))
				m.fn1230(v16, v4+i32(352))
				goto l7
			case 19:
				m.fn1441(v4 + i32(152))
				t44 := int32(load32(m.memory[int64(uint32(v4))+172:]))
				v26 = t44
				if v26 == 0 {
					goto l7
				}
				t45 := int32(load32(m.memory[int64(uint32(v4))+168:]))
				v26 = t45 + v26*i32(28)
				if v26+i32(-28) == 0 {
					goto l7
				}
				m.memory[uint32(v26+i32(-4))] = byte(i32(1))
				goto l7
			case 20:
				m.fn1442(v4 + i32(152))
				goto l7
			case 8:
				t46 := m.fn1436(v1, v28, v2)
				m.fn1437(v4+i32(152), i32(32), t46)
				goto l7
			case 29:
				t47 := m.fn1436(v1, v28, v2)
				m.fn1437(v4+i32(152), i32(45), t47)
				goto l7
			case 0:
				t48 := int32(load32(m.memory[uint32(v22):]))
				t49 := int32(load32(m.memory[uint32(v21):]))
				t50 := m.fn1438(t48, t49, v28)
				v32 = t50
				if v32 == 0 {
					goto l7
				}
				t51 := int32(load32(m.memory[int64(uint32(v32))+56:]))
				v26 = t51
				t52 := int32(load32(m.memory[int64(uint32(v32))+52:]))
				v27 = t52
				v32 = i32(0)
				v33 = i32(0)
				{
				l28:
					{
						t53 := v26
						v35 = v32 + i32(2)
						if uint32(t53) < uint32(v35) {
							goto l24
						}
						if uint32(v32) >= uint32(v26) {
							m.fn158(v32, v26, i32(1072496))
							panic("unreachable")
						}
						v30 = v32 + i32(1)
						if uint32(v30) >= uint32(v26) {
							m.fn158(v30, v26, i32(1072512))
							panic("unreachable")
						}
						t54 := int32(m.memory[uint32(v27+v32)])
						v32 = t54
						t55 := int32(m.memory[uint32(v27+v30)])
						v30 = t55
						m.fn148(v4+i32(48), v35, v27, v26, i32(1072528))
						v30 = v32 | v30<<8
						t56 := int32(load32(m.memory[int64(uint32(v4))+48:]))
						t57 := int32(load32(m.memory[int64(uint32(v4))+52:]))
						t58 := m.fn1430(v30, t56, t57)
						v32 = t58
						if uint32(v32) > uint32(v26-v35) {
							goto l24
						}
						{
							if v30&i32(0xffff) != i32(27139) {
								goto l27
							}
							v33 = i32(0)
							if uint32(v32) < uint32(i32(4)) {
								goto l27
							}
							t59 := int32(load32(m.memory[uint32(v27+v35):]))
							v28 = t59
							v33 = i32(1)
						}
					l27:
						v32 = v32 + v35
						goto l28
					}
				l24:
					if v33&i32(1) == 0 {
						goto l7
					}
					store32(m.memory[int64(uint32(v4))+920:], uint32(v28))
					t60 := int32(load32(m.memory[int64(uint32(v1))+384:]))
					v26 = t60
					if uint32(v26) < uint32(v28) {
						goto l7
					}
					if uint32(v26-v28) < uint32(i32(4)) {
						goto l7
					}
					t61 := v26
					v32 = v28 + i32(4)
					if uint32(t61) < uint32(v32) {
						goto l7
					}
					if uint32(v26-v32) < uint32(i32(2)) {
						goto l7
					}
					t62 := int32(load32(m.memory[int64(uint32(v1))+380:]))
					t63 := v28
					v27 = t62
					v30 = v27 + v28
					t64 := int32(load32(m.memory[uint32(v30):]))
					v35 = t63 + t64
					p65 := v35
					if uint32(v35) < uint32(v28) {
						p65 = i32(-1)
					}
					v35 = p65
					if uint32(v35) > uint32(v26) {
						goto l7
					}
					t66 := v4 + i32(40)
					v26 = v35 - v28
					t67 := int32(load16(m.memory[uint32(v27+v32):]))
					t68 := v26
					v32 = t67
					p69 := v32
					if uint32(v26) < uint32(v32) {
						p69 = t68
					}
					m.fn148(t66, p69, v30, v26, i32(1078096))
					t70 := int32(load32(m.memory[int64(uint32(v4))+40:]))
					v36 = t70
					t71 := int32(load32(m.memory[int64(uint32(v4))+44:]))
					v37 = t71
					t72 := m.fn113(i32(4), i32(8))
					v35 = t72
					store32(m.memory[int64(uint32(v35))+4:], uint32(v37))
					v38 = i32(0)
					store32(m.memory[uint32(v35):], uint32(i32(0)))
					v26 = i32(1)
					store32(m.memory[int64(uint32(v4))+200:], uint32(i32(1)))
					store32(m.memory[int64(uint32(v4))+196:], uint32(v35))
					store32(m.memory[int64(uint32(v4))+192:], uint32(i32(1)))
				l38:
					{
						{
							if v26 == 0 {
								goto l29
							}
							v30 = v35 + v26<<3
							v32 = v30 + i32(-8)
							if v32 == 0 {
								goto l29
							}
							{
								t73 := int32(load32(m.memory[uint32(v32):]))
								v27 = t73
								t74 := int32(load32(m.memory[uint32(v30+i32(-4)):]))
								t75 := v27
								v30 = t74
								if uint32(t75) < uint32(v30) {
									if uint32(v30) > uint32(v37) {
										m.fn151(i32(0), v30, v37, i32(1081484))
										panic("unreachable")
									}
									m.fn1285(v4+i32(864), v36, v30, v27)
									t76 := int32(load32(m.memory[int64(uint32(v4))+868:]))
									v33 = t76
									if v33 == 0 {
										goto l33
									}
									if uint32(v27) > uint32(i32(-9)) {
										goto l29
									}
									t77 := int32(load32(m.memory[int64(uint32(v4))+872:]))
									v30 = t77
									t78 := v30
									v39 = v27 + i32(8)
									v28 = t78 + v39
									if uint32(v28) < uint32(v30) {
										goto l29
									}
									t79 := int32(load16(m.memory[int64(uint32(v4))+866:]))
									v40 = t79
									t80 := int32(load16(m.memory[int64(uint32(v4))+864:]))
									v41 = t80
									store32(m.memory[uint32(v32):], uint32(v28))
									v38 = v38 + i32(1)
									if uint32(v38) > uint32(i32(10000)) {
										goto l29
									}
									if uint32(v26) > uint32(i32(16)) {
										goto l29
									}
									m.fn1295(v4+i32(352), v41, v40, v33, v30)
									t81 := int32(load32(m.memory[int64(uint32(v4))+352:]))
									if t81 != i32(-2) {
										t84 := int32(load32(m.memory[int64(uint32(v4))+192:]))
										m.fn76(t84, v35)
										t85 := int32(load32(m.memory[int64(uint32(v4))+376:]))
										store32(m.memory[int64(uint32(v4))+888:], uint32(t85))
										t86 := int64(load64(m.memory[int64(uint32(v4))+368:]))
										store64(m.memory[int64(uint32(v4))+880:], uint64(t86))
										t87 := int64(load64(m.memory[int64(uint32(v4))+360:]))
										store64(m.memory[int64(uint32(v4))+872:], uint64(t87))
										t88 := int64(load64(m.memory[int64(uint32(v4))+352:]))
										store64(m.memory[int64(uint32(v4))+864:], uint64(t88))
										store32(m.memory[int64(uint32(v4))+364:], uint32(i32(1)))
										store32(m.memory[int64(uint32(v4))+360:], uint32(v15))
										store32(m.memory[int64(uint32(v4))+356:], uint32(i32(5)))
										store32(m.memory[int64(uint32(v4))+352:], uint32(v4+i32(920)))
										m.fn73(v4+i32(256), i32(0x10009e), v4+i32(352))
										m.fn1182(v4+i32(32), v7, i32(1078112))
										t89 := int32(load32(m.memory[int64(uint32(v4))+36:]))
										v26 = t89
										t90 := int32(load32(m.memory[int64(uint32(v4))+32:]))
										v32 = t90
										t91 := int32(load32(m.memory[int64(uint32(v4))+876:]))
										t92 := int32(load32(m.memory[int64(uint32(v4))+880:]))
										m.fn51(v4+i32(192), t91, t92)
										t93 := int32(load32(m.memory[int64(uint32(v4))+868:]))
										t94 := int32(load32(m.memory[int64(uint32(v4))+872:]))
										m.fn1296(v4+i32(352), v32, v4+i32(192), v4+i32(256), t93, t94)
										t95 := int32(load32(m.memory[int64(uint32(v4))+356:]))
										v32 = t95
										{
											t96 := int32(load32(m.memory[int64(uint32(v4))+352:]))
											v35 = t96
											if v35 == i32(-1) {
												t104 := int32(load32(m.memory[uint32(v26):]))
												store32(m.memory[uint32(v26):], uint32(t104+i32(1)))
												t105 := int32(load32(m.memory[int64(uint32(v4))+864:]))
												t106 := int32(load32(m.memory[int64(uint32(v4))+868:]))
												m.fn134(t105, t106)
												store32(m.memory[int64(uint32(v4))+372:], uint32(v32))
												store32(m.memory[int64(uint32(v4))+368:], uint32(i32(-0x80000000)))
												store64(m.memory[int64(uint32(v4))+360:], uint64(i64(1)))
												store64(m.memory[int64(uint32(v4))+352:], uint64(i64(5)))
												m.fn1435(v4+i32(152), v4+i32(352))
												goto l7
											}
											t97 := int64(load64(m.memory[int64(uint32(v4))+360:]))
											v29 = t97
											t98 := int64(load64(m.memory[int64(uint32(v4))+368:]))
											v31 = t98
											t99 := int32(load32(m.memory[uint32(v26):]))
											store32(m.memory[uint32(v26):], uint32(t99+i32(1)))
											t100 := int32(load32(m.memory[int64(uint32(v4))+864:]))
											t101 := int32(load32(m.memory[int64(uint32(v4))+868:]))
											m.fn134(t100, t101)
											store64(m.memory[int64(uint32(v0))+16:], uint64(v31))
											store64(m.memory[int64(uint32(v0))+8:], uint64(v29))
											store32(m.memory[int64(uint32(v0))+4:], uint32(v32))
											store32(m.memory[uint32(v0):], uint32(v35))
											t102 := int32(load32(m.memory[int64(uint32(v4))+180:]))
											v1 = t102
											t103 := int32(load32(m.memory[int64(uint32(v4))+176:]))
											v2 = t103
											goto l23
										}
									}
									if v40&i32(0xffff) != i32(61447) {
										if v41&i32(15) != i32(15) {
											goto l31
										}
										m.fn1443(v4+i32(192), v39, v28)
										goto l31
									}
									v26 = v28
									{
										if uint32(v30) < uint32(i32(34)) {
											goto l36
										}
										t82 := int32(m.memory[int64(uint32(v33))+33])
										v26 = v27 + t82 + i32(44)
									}
								l36:
									if uint32(v26) < uint32(v39) {
										goto l29
									}
									if uint32(v26) >= uint32(v28) {
										goto l31
									}
									m.fn1443(v4+i32(192), v26, v28)
									goto l31
								}
								store32(m.memory[int64(uint32(v4))+200:], uint32(v26+i32(-1)))
								goto l31
							}
						l29:
							t83 := int32(load32(m.memory[int64(uint32(v4))+192:]))
							m.fn76(t83, v35)
							goto l7
						}
					l33:
						store32(m.memory[int64(uint32(v4))+200:], uint32(v26+i32(-1)))
					l31:
						t107 := int32(load32(m.memory[int64(uint32(v4))+196:]))
						v35 = t107
						t108 := int32(load32(m.memory[int64(uint32(v4))+200:]))
						v26 = t108
						goto l38
					}
				}
			}
		l23:
			m.fn894(v4 + i32(152))
			m.fn1444(v16)
			m.fn16(v2, v1)
			goto l2
		l21:
			m.fn1333(v4+i32(84), v4+i32(60))
			m.fn1351(v4+i32(60), v4+i32(72))
			{
				{
					t170 := int32(load32(m.memory[int64(uint32(v4))+192:]))
					if t170 != i32(1) {
						goto l51
					}
					t171 := int32(load32(m.memory[int64(uint32(v4))+196:]))
					if t171 > i32(1) {
						goto l52
					}
				}
			l51:
				t172 := int32(m.memory[int64(uint32(v4))+238])
				if t172&i32(1) != 0 {
					goto l52
				}
				t173 := int32(m.memory[int64(uint32(v4))+239])
				if t173&i32(1) != 0 {
					goto l52
				}
				if v32 != 0 {
					goto l53
				}
				m.fn1448(v1, v26, v4+i32(244), v4+i32(100), v4+i32(112))
				goto l40
			}
		l52:
			m.fn1448(v1, v26, v4+i32(244), v4+i32(100), v4+i32(112))
			goto l40
		l53:
			{
				t174 := int32(m.memory[int64(uint32(v4))+237])
				if t174&i32(1) != 0 {
					goto l54
				}
				m.fn1448(v1, v26, v4+i32(244), v4+i32(100), v4+i32(112))
				m.fn1333(v4+i32(112), v4+i32(100))
				t175 := int64(load64(m.memory[int64(uint32(v4))+100:]))
				v29 = t175
				store64(m.memory[int64(uint32(v4))+100:], uint64(i64(0x800000000)))
				t176 := int32(load32(m.memory[int64(uint32(v4))+108:]))
				v26 = t176
				store32(m.memory[int64(uint32(v4))+108:], uint32(i32(0)))
				store32(m.memory[int64(uint32(v4))+360:], uint32(v26))
				store64(m.memory[int64(uint32(v4))+352:], uint64(v29))
				m.fn1169(v4+i32(128), v4+i32(352))
				goto l40
			}
		l54:
			m.fn1333(v4+i32(112), v4+i32(100))
			{
				t177 := int32(load32(m.memory[int64(uint32(v4))+136:]))
				if t177 == 0 {
					goto l55
				}
				t178 := int64(load64(m.memory[int64(uint32(v4))+128:]))
				v29 = t178
				store64(m.memory[int64(uint32(v4))+128:], uint64(i64(0x400000000)))
				t179 := int32(load32(m.memory[int64(uint32(v4))+136:]))
				v26 = t179
				store32(m.memory[int64(uint32(v4))+136:], uint32(i32(0)))
				store32(m.memory[int64(uint32(v4))+360:], uint32(v26))
				store64(m.memory[int64(uint32(v4))+352:], uint64(v29))
				m.fn1449(v8, v9)
				m.fn1450(v4+i32(140), v4+i32(352))
			}
		l55:
			m.fn894(v4 + i32(244))
		l40:
			m.fn766(v9)
		l7:
			v2 = v2 + i32(1)
			goto l56
		}
		memory_copy(m.memory, uint32(v4+i32(352)), uint32(v4+i32(152)), uint32(i32(40)))
		m.fn1433(v4+i32(864), v4+i32(352))
		m.fn1333(v4+i32(112), v4+i32(100))
		m.fn1434(v4+i32(352), v4+i32(60), v4+i32(140), v4+i32(128), v4+i32(100))
		t4 := int32(load32(m.memory[int64(uint32(v4))+352:]))
		v2 = t4
		if v2 == i32(-1) {
			t180 := int32(load32(m.memory[int64(uint32(v4))+868:]))
			t181 := int32(load32(m.memory[int64(uint32(v4))+872:]))
			t182 := m.fn23(t180, t181)
			v2 = t182
			m.fn1333(v4+i32(84), v4+i32(60))
			m.fn1351(v4+i32(60), v4+i32(72))
			{
				{
					if v2 != 0 {
						goto l57
					}
					t183 := int32(load32(m.memory[int64(uint32(v4))+872:]))
					store32(m.memory[int64(uint32(v4))+364:], uint32(t183))
					t184 := int64(load64(m.memory[int64(uint32(v4))+864:]))
					store64(m.memory[int64(uint32(v4))+356:], uint64(t184))
					store32(m.memory[int64(uint32(v4))+352:], uint32(i32(-0x80000000)))
					m.fn338(v4+i32(60), v4+i32(352))
					m.fn1333(v4+i32(84), v4+i32(60))
					m.fn1351(v4+i32(60), v4+i32(72))
					t185 := int32(load32(m.memory[int64(uint32(v4))+68:]))
					store32(m.memory[int64(uint32(v0))+12:], uint32(t185))
					t186 := int64(load64(m.memory[int64(uint32(v4))+60:]))
					store64(m.memory[int64(uint32(v0))+4:], uint64(t186))
					goto l58
				}
			l57:
				t187 := int32(load32(m.memory[int64(uint32(v4))+68:]))
				store32(m.memory[int64(uint32(v0))+12:], uint32(t187))
				t188 := int64(load64(m.memory[int64(uint32(v4))+60:]))
				store64(m.memory[int64(uint32(v0))+4:], uint64(t188))
				m.fn894(v4 + i32(864))
			}
		l58:
			store32(m.memory[uint32(v0):], uint32(i32(-1)))
			m.fn1451(v4 + i32(140))
			m.fn976(v4 + i32(128))
			m.fn1332(v4 + i32(112))
			m.fn969(v4 + i32(100))
			m.fn1332(v4 + i32(84))
			m.fn1302(v4 + i32(72))
			goto l59
		}
		t5 := int32(load32(m.memory[int64(uint32(v4))+372:]))
		store32(m.memory[int64(uint32(v0))+20:], uint32(t5))
		t6 := int64(load64(m.memory[int64(uint32(v4))+364:]))
		store64(m.memory[int64(uint32(v0))+12:], uint64(t6))
		t7 := int64(load64(m.memory[int64(uint32(v4))+356:]))
		store64(m.memory[int64(uint32(v0))+4:], uint64(t7))
		store32(m.memory[uint32(v0):], uint32(v2))
		m.fn894(v4 + i32(864))
		goto l2
	}
l2:
	m.fn1451(v4 + i32(140))
	m.fn976(v4 + i32(128))
	m.fn1332(v4 + i32(112))
	m.fn969(v4 + i32(100))
	m.fn1332(v4 + i32(84))
	m.fn1302(v4 + i32(72))
	m.fn969(v4 + i32(60))
l59:
	m.g0 = v4 + i32(928)
}
func (m *Module) fn1227(v0 int32) {
	var v1, v2, v3, v4, v5 int32
	var v6 int64
	var v7, v8 int32
	t0 := m.g0
	v1 = t0 - i32(16)
	m.g0 = v1
	t1 := int32(load32(m.memory[int64(uint32(v0))+88:]))
	t2 := int32(load32(m.memory[int64(uint32(v0))+92:]))
	m.fn44(t1, t2)
	t3 := int32(load32(m.memory[int64(uint32(v0))+100:]))
	t4 := int32(load32(m.memory[int64(uint32(v0))+104:]))
	m.fn449(t3, t4)
	t5 := int32(load32(m.memory[int64(uint32(v0))+112:]))
	t6 := int32(load32(m.memory[int64(uint32(v0))+116:]))
	m.fn449(t5, t6)
	t7 := int32(load32(m.memory[int64(uint32(v0))+124:]))
	t8 := int32(load32(m.memory[int64(uint32(v0))+128:]))
	m.fn449(t7, t8)
	m.fn979(v0 + i32(328))
	m.fn979(v0 + i32(340))
	m.fn1212(v0 + i32(56))
	m.fn766(v0 + i32(8))
	{
		t9 := int32(load32(m.memory[int64(uint32(v0))+140:]))
		v2 = t9
		if v2 == 0 {
			goto l0
		}
		t10 := int32(load32(m.memory[int64(uint32(v0))+136:]))
		v3 = t10
		{
			t11 := int32(load32(m.memory[int64(uint32(v0))+148:]))
			v4 = t11
			if v4 == 0 {
				goto l1
			}
			v5 = v3 + i32(8)
			t12 := int64(load64(m.memory[uint32(v3):]))
			v6 = (t12 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
			v7 = v3
		l4:
			if v4 == 0 {
				goto l1
			}
		l3:
			{
				if v6 != i64(0) {
					m.fn761(v7 + (i32(0)-int32(uint32(int64(bits.TrailingZeros64(uint64(v6))))>>3))*i32(520) + i32(-368))
					v4 = v4 + i32(-1)
					v6 = (v6 + i64(-1)) & v6
					goto l4
				}
				v7 = v7 + i32(-4160)
				t13 := int64(load64(m.memory[uint32(v5):]))
				v6 = (t13 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
				v5 = v5 + i32(8)
				goto l3
			}
		}
	l1:
		m.fn39(v1+i32(4), i32(520), i32(8), v2+i32(1))
		t14 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		t15 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t16 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		m.fn40(v3-t14, t15, t16)
	}
l0:
	m.fn1202(v0 + i32(352))
	t17 := int32(load32(m.memory[int64(uint32(v0))+364:]))
	t18 := int32(load32(m.memory[int64(uint32(v0))+368:]))
	m.fn136(t17, t18, i32(4), i32(8))
	{
		t19 := int32(load32(m.memory[int64(uint32(v0))+172:]))
		v8 = t19
		if v8 == 0 {
			goto l5
		}
		t20 := int32(load32(m.memory[int64(uint32(v0))+168:]))
		v2 = t20
		{
			t21 := int32(load32(m.memory[int64(uint32(v0))+180:]))
			v4 = t21
			if v4 == 0 {
				goto l6
			}
			v5 = v2 + i32(8)
			t22 := int64(load64(m.memory[uint32(v2):]))
			v6 = (t22 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
			v7 = v2
		l9:
			if v4 == 0 {
				goto l6
			}
		l8:
			{
				if v6 != i64(0) {
					v3 = v7 - int32(int64(bits.TrailingZeros64(uint64(v6))))<<1&i32(240)
					t24 := int32(load32(m.memory[uint32(v3+i32(-12)):]))
					t25 := int32(load32(m.memory[uint32(v3+i32(-8)):]))
					m.fn16(t24, t25)
					v4 = v4 + i32(-1)
					v6 = (v6 + i64(-1)) & v6
					goto l9
				}
				v7 = v7 + i32(-128)
				t23 := int64(load64(m.memory[uint32(v5):]))
				v6 = (t23 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
				v5 = v5 + i32(8)
				goto l8
			}
		}
	l6:
		m.fn39(v1+i32(4), i32(16), i32(8), v8+i32(1))
		t26 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		t27 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t28 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		m.fn40(v2-t26, t27, t28)
	}
l5:
	{
		t29 := int32(load32(m.memory[int64(uint32(v0))+212:]))
		v5 = t29
		if v5 == 0 {
			goto l10
		}
		t30 := int32(load32(m.memory[int64(uint32(v0))+208:]))
		v7 = t30
		m.fn39(v1+i32(4), i32(96), i32(8), v5+i32(1))
		t31 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		t32 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t33 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		m.fn40(v7-t31, t32, t33)
	}
l10:
	{
		t34 := int32(load32(m.memory[int64(uint32(v0))+244:]))
		v5 = t34
		if v5 == 0 {
			goto l11
		}
		t35 := int32(load32(m.memory[int64(uint32(v0))+240:]))
		v7 = t35
		m.fn39(v1+i32(4), i32(8), i32(8), v5+i32(1))
		t36 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		t37 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t38 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		m.fn40(v7-t36, t37, t38)
	}
l11:
	t39 := int32(load32(m.memory[int64(uint32(v0))+376:]))
	t40 := int32(load32(m.memory[int64(uint32(v0))+380:]))
	m.fn16(t39, t40)
	m.fn1274(v0 + i32(280))
	m.g0 = v1 + i32(16)
}
func (m *Module) fn1228(v0 int32) {
	var v1, v2 int32
	t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
	t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	v1 = t1
	t2 := int32(uint32(t0-v1) / uint32(i32(24)))
	v2 = t2
l1:
	{
		if v2 == 0 {
			goto l0
		}
		t3 := int32(load32(m.memory[uint32(v1):]))
		t4 := int32(load32(m.memory[uint32(v1+i32(4)):]))
		m.fn16(t3, t4)
		v2 = v2 + i32(-1)
		v1 = v1 + i32(24)
		goto l1
	}
l0:
	t5 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	t6 := int32(load32(m.memory[uint32(v0):]))
	m.fn1201(t5, t6)
}
func (m *Module) fn1229(v0 int32) {
	var v1 int32
	t0 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	v1 = t0
	t1 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	m.fn81(v1, t1)
	t2 := int32(load32(m.memory[uint32(v0):]))
	m.fn82(t2, v1)
}
func (m *Module) fn1230(v0, v1 int32) {
	var v2 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		v2 = t0
		t1 := int32(load32(m.memory[uint32(v0):]))
		if v2 != t1 {
			goto l0
		}
		m.fn1143(v0)
	}
l0:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v2+i32(1)))
	t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	v0 = t2 + v2*i32(28)
	t3 := int64(load64(m.memory[uint32(v1):]))
	store64(m.memory[uint32(v0):], uint64(t3))
	t4 := int64(load64(m.memory[int64(uint32(v1))+8:]))
	store64(m.memory[int64(uint32(v0))+8:], uint64(t4))
	t5 := int64(load64(m.memory[int64(uint32(v1))+16:]))
	store64(m.memory[int64(uint32(v0))+16:], uint64(t5))
	t6 := int32(load32(m.memory[int64(uint32(v1))+24:]))
	store32(m.memory[int64(uint32(v0))+24:], uint32(t6))
}
func (m *Module) fn1231(v0 int32) {
	var v1 int32
	t0 := m.fn1679(i32(432))
	v1 = t0
	store64(m.memory[uint32(v1):], uint64(i64(1)))
	memory_zero(m.memory, uint32(v1+i32(8)), uint32(i32(260)))
	store32(m.memory[int64(uint32(v1))+268:], uint32(i32(1)))
	memory_zero(m.memory, uint32(v1+i32(272)), uint32(i32(145)))
	store32(m.memory[int64(uint32(v1))+426:], uint32(i32(257)))
	m.memory[int64(uint32(v1))+424] = byte(i32(0))
	m.memory[int64(uint32(v1))+422] = byte(i32(0))
	store32(m.memory[int64(uint32(v1))+417:], uint32(i32(8748)))
	m.memory[int64(uint32(v0))+10] = byte(i32(1))
	store32(m.memory[uint32(v0):], uint32(i32(8192)))
	store16(m.memory[int64(uint32(v0))+8:], uint16(i32(0)))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
}
func (m *Module) fn1232(v0 int32) {
	m.fn10(v0, i32(432), i32(8))
}
func (m *Module) fn1233(v0, v1 int32) {
	var v2 int32
	{
		t0 := int32(load32(m.memory[uint32(v0):]))
		t1 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		t2 := v1
		v2 = t1
		if uint32(t2) <= uint32(t0-v2) {
			return
		}
		m.fn1684(v0, v2, v1, i32(4), i32(4))
	}
}
func (m *Module) fn1234(v0, v1, v2 int32) int32 {
	var v3, v4, v5, v6, v7, v8, v9, v10, v11 int32
	var v12 int64
	var v13, v14 int32
	var v15 int64
	var v16 int32
	var v17, v18 int64
	var v19, v20, v21, v22, v23 int32
	t0 := m.g0
	v3 = t0 - i32(272)
	m.g0 = v3
	if v2 != 0 {
		p1 := v2
		if uint32(v1) < uint32(v2) {
			p1 = v1
		}
		v4 = p1
		goto l3
	}
	if uint32(v1) >= uint32(i32(2)) {
		t2 := m.fn857(v0, v1, i32(0), i32(1081392))
		v2 = t2
		t3 := int32(load32(m.memory[int64(uint32(v2))+8:]))
		v5 = t3 * i32(20)
		t4 := int32(load32(m.memory[int64(uint32(v2))+4:]))
		v2 = t4
	l5:
		{
			if v5 == 0 {
				p8 := i32(51)
				if uint32(v1) < uint32(i32(51)) {
					p8 = v1
				}
				v5 = p8
				if uint32(v5+i32(-1)) >= uint32(v1) {
					m.fn151(i32(1), v5, v1, i32(1081408))
					panic("unreachable")
				}
				t9 := m.fn857(v0, v1, i32(0), i32(1081424))
				t10 := int32(load32(m.memory[int64(uint32(t9))+8:]))
				v7 = t10
				v2 = v0 + i32(12)
				t11 := v0
				v8 = v5 * i32(12)
				v9 = t11 + v8
				m.fn27(v3 + i32(96))
				v10 = v3 + i32(112)
			l12:
				{
					if v2 == v9 {
						t32 := int32(load32(m.memory[int64(uint32(v3))+96:]))
						v2 = t32
						v6 = v2 + i32(8)
						t33 := int32(load32(m.memory[int64(uint32(v3))+100:]))
						t34 := v2
						v5 = t33
						v4 = t34 + v5 + i32(1)
						t35 := int64(load64(m.memory[uint32(v2):]))
						v12 = (t35 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
						t36 := int32(load32(m.memory[int64(uint32(v3))+108:]))
						v11 = t36
						{
							if v5 != 0 {
								goto l14
							}
							v5 = i32(0)
							goto l15
						l14:
							m.fn39(v3+i32(176), i32(8), i32(8), v5+i32(1))
							t37 := int32(load32(m.memory[int64(uint32(v3))+180:]))
							store32(m.memory[int64(uint32(v3))+164:], uint32(t37))
							t38 := int32(load32(m.memory[int64(uint32(v3))+184:]))
							store32(m.memory[int64(uint32(v3))+168:], uint32(v2-t38))
							t39 := int32(load32(m.memory[int64(uint32(v3))+176:]))
							v5 = t39
						}
					l15:
						store32(m.memory[int64(uint32(v3))+152:], uint32(v11))
						store32(m.memory[int64(uint32(v3))+144:], uint32(v2))
						store32(m.memory[int64(uint32(v3))+140:], uint32(v4))
						store32(m.memory[int64(uint32(v3))+136:], uint32(v6))
						store64(m.memory[int64(uint32(v3))+128:], uint64(v12))
						store32(m.memory[int64(uint32(v3))+160:], uint32(v5))
						m.fn1023(v3+i32(176), v3+i32(128))
						{
							t40 := int32(load32(m.memory[int64(uint32(v3))+176:]))
							if t40 != i32(1) {
								goto l16
							}
							t41 := int32(load32(m.memory[int64(uint32(v3))+184:]))
							v5 = t41
							t42 := int32(load32(m.memory[int64(uint32(v3))+180:]))
							v11 = t42
							memory_copy(m.memory, uint32(v3+i32(176)), uint32(v3+i32(128)), uint32(i32(48)))
							v9 = v3 + i32(256)
							v4 = v11
							v13 = v5
						l21:
							{
								m.fn1023(v3+i32(228), v3+i32(176))
								t43 := int32(load32(m.memory[int64(uint32(v3))+228:]))
								if t43 != i32(1) {
									m.fn1024(v3 + i32(176))
									goto l22
								}
								t44 := int32(load32(m.memory[int64(uint32(v3))+232:]))
								v6 = t44
								t45 := int32(load32(m.memory[int64(uint32(v3))+236:]))
								t46 := v3
								v2 = t45
								store32(m.memory[int64(uint32(t46))+268:], uint32(v2))
								store32(m.memory[int64(uint32(v3))+264:], uint32(v6))
								store32(m.memory[int64(uint32(v3))+260:], uint32(v6))
								store32(m.memory[int64(uint32(v3))+256:], uint32(v2))
								store32(m.memory[int64(uint32(v3))+252:], uint32(v13))
								store32(m.memory[int64(uint32(v3))+248:], uint32(v11))
								store32(m.memory[int64(uint32(v3))+244:], uint32(v4))
								store32(m.memory[int64(uint32(v3))+240:], uint32(v5))
								if v5 != v2 {
									goto l18
								}
								if uint32(v4) > uint32(v6) {
									goto l19
								}
								v2 = v5
								v6 = v9
								goto l20
							l18:
								v6 = v9
								if uint32(v5) <= uint32(v2) {
									goto l20
								}
							l19:
								v6 = v3 + i32(240)
								v2 = v5
							l20:
								t47 := int32(load32(m.memory[int64(uint32(v6))+12:]))
								v13 = t47
								t48 := int32(load32(m.memory[int64(uint32(v6))+8:]))
								v11 = t48
								t49 := int32(load32(m.memory[int64(uint32(v6))+4:]))
								v4 = t49
								v5 = v2
								goto l21
							}
						}
					l16:
						m.fn1024(v3 + i32(128))
						v11 = i32(0)
					l22:
						if v7 != v11 {
							goto l2
						}
						v5 = i32(0)
						t50 := m.fn857(v0, v1, i32(0), i32(1081440))
						t51 := v3 + i32(68)
						v2 = t50
						t52 := int32(load32(m.memory[int64(uint32(v2))+4:]))
						v6 = t52
						t53 := int32(load32(m.memory[int64(uint32(v2))+8:]))
						m.fn792(t51, v6, v6+t53*i32(20))
						t54 := int32(uint32(v8+i32(-12)) / uint32(i32(12)))
						t55 := v3 + i32(56)
						v6 = t54
						m.fn59(t55, v6, i32(4), i32(12))
						store32(m.memory[int64(uint32(v3))+136:], uint32(i32(0)))
						t56 := int32(load32(m.memory[int64(uint32(v3))+60:]))
						t57 := v3
						v4 = t56
						store32(m.memory[int64(uint32(t57))+132:], uint32(v4))
						t58 := int32(load32(m.memory[int64(uint32(v3))+56:]))
						t59 := v3
						v2 = t58
						store32(m.memory[int64(uint32(t59))+128:], uint32(v2))
						{
							if uint32(v6) <= uint32(v2) {
								goto l23
							}
							m.fn62(v3+i32(128), i32(0), v6, i32(4), i32(12))
							t60 := int32(load32(m.memory[int64(uint32(v3))+136:]))
							v5 = t60
							t61 := int32(load32(m.memory[int64(uint32(v3))+132:]))
							v4 = t61
						}
					l23:
						v11 = v5 + v6
						v2 = v0 + i32(20)
						v5 = v4 + v5*i32(12)
					l24:
						{
							t62 := int32(load32(m.memory[uint32(v2+i32(-4)):]))
							t63 := v3 + i32(176)
							v4 = t62
							t64 := int32(load32(m.memory[uint32(v2):]))
							m.fn792(t63, v4, v4+t64*i32(20))
							t65 := int32(load32(m.memory[int64(uint32(v3))+184:]))
							store32(m.memory[int64(uint32(v5))+8:], uint32(t65))
							t66 := int64(load64(m.memory[int64(uint32(v3))+176:]))
							store64(m.memory[uint32(v5):], uint64(t66))
							v2 = v2 + i32(12)
							v5 = v5 + i32(12)
							v6 = v6 + i32(-1)
							if v6 != 0 {
								goto l24
							}
						}
						t67 := int64(load64(m.memory[int64(uint32(v3))+128:]))
						store64(m.memory[int64(uint32(v3))+80:], uint64(t67))
						store32(m.memory[int64(uint32(v3))+88:], uint32(v11))
						t68 := int32(load32(m.memory[int64(uint32(v3))+84:]))
						v14 = t68
						t69 := int32(load32(m.memory[int64(uint32(v3))+72:]))
						v1 = t69
						t70 := int32(load32(m.memory[int64(uint32(v3))+76:]))
						t71 := v1
						v6 = t70
						t72 := m.fn1017(t71, v6)
						v13 = t72
						if v11 == 0 {
							goto l25
						}
						v2 = v14 + i32(8)
						v5 = v11
					l26:
						{
							t73 := int32(load32(m.memory[uint32(v2+i32(-4)):]))
							t74 := int32(load32(m.memory[uint32(v2):]))
							t75 := m.fn1016(v13, t73, t74)
							v13 = t75
							v2 = v2 + i32(12)
							v5 = v5 + i32(-1)
							if v5 != 0 {
								goto l26
							}
						}
					l25:
						v4 = i32(0)
						{
							if v13 == 0 {
								goto l27
							}
							if uint32(v13) > uint32(v6) {
								goto l27
							}
							v0 = v14 + v11*i32(12)
							m.fn27(v3 + i32(176))
							store32(m.memory[int64(uint32(v3))+140:], uint32(v13))
							store32(m.memory[int64(uint32(v3))+136:], uint32(i32(0)))
							store32(m.memory[int64(uint32(v3))+128:], uint32(v1))
							t76 := v3
							v11 = v1 + v6*i32(12)
							store32(m.memory[int64(uint32(t76))+132:], uint32(v11))
						l31:
							m.fn1424(v3+i32(48), v3+i32(128))
							{
								t77 := int32(load32(m.memory[int64(uint32(v3))+52:]))
								v2 = t77
								if v2 == 0 {
									store32(m.memory[int64(uint32(v3))+140:], uint32(v13))
									store32(m.memory[int64(uint32(v3))+136:], uint32(i32(0)))
									store32(m.memory[int64(uint32(v3))+132:], uint32(v11))
									store32(m.memory[int64(uint32(v3))+128:], uint32(v1))
									v10 = v3 + i32(240) + i32(8)
									v19 = i32(0)
									v20 = i32(0)
								l37:
									m.fn1424(v3+i32(32), v3+i32(128))
									{
										t88 := int32(load32(m.memory[int64(uint32(v3))+36:]))
										v16 = t88
										if v16 == 0 {
											if v19|v20 != 0 {
												var p109 int32
												if uint32(v19) > uint32(v20) {
													p109 = 1
												}
												v4 = p109
												m.fn38(v3 + i32(176))
												m.fn1426(v3 + i32(80))
												m.fn78(v3 + i32(68))
												goto l3
											}
											t101 := int32(load32(m.memory[int64(uint32(v3))+76:]))
											v5 = t101 * i32(12)
											v6 = v13 + i32(-1)
											t102 := int32(load32(m.memory[int64(uint32(v3))+72:]))
											v2 = t102
											v4 = i32(1)
										l40:
											{
												if v5 == 0 {
													goto l39
												}
												t103 := int32(load32(m.memory[int64(uint32(v2))+4:]))
												v11 = t103
												t104 := int32(load32(m.memory[int64(uint32(v2))+8:]))
												t105 := m.fn853(v11, v11+t104)
												var p106 int32
												if uint32(t105) > uint32(i32(64)) {
													p106 = 1
												}
												v11 = p106
												p107 := i32(2)
												if v11 != 0 {
													p107 = i32(1)
												}
												p108 := v11
												if v6 != 0 {
													p108 = p107
												}
												v11 = p108
												v5 = v5 + i32(-12)
												v6 = v6 + i32(-1)
												v2 = v2 + i32(12)
												if v11 == i32(2) {
													goto l40
												}
											}
											v4 = v11 ^ i32(1)
											goto l39
										}
										t89 := int32(load32(m.memory[int64(uint32(v3))+32:]))
										store32(m.memory[int64(uint32(v3))+92:], uint32(t89))
										store32(m.memory[int64(uint32(v3))+232:], uint32(v0))
										store32(m.memory[int64(uint32(v3))+228:], uint32(v14))
										store32(m.memory[int64(uint32(v3))+236:], uint32(v3+i32(92)))
										m.fn910(v3+i32(24), v3+i32(228))
										{
											t90 := int32(load32(m.memory[int64(uint32(v3))+24:]))
											v2 = t90
											if v2 == 0 {
												m.fn639(i32(0), i32(4))
												goto l37
											}
											t91 := int32(load32(m.memory[int64(uint32(v3))+28:]))
											v5 = t91
											v9 = i32(8)
											m.fn59(v3+i32(16), i32(4), i32(4), i32(8))
											t92 := int32(load32(m.memory[int64(uint32(v3))+16:]))
											v6 = t92
											t93 := int32(load32(m.memory[int64(uint32(v3))+20:]))
											v11 = t93
											store32(m.memory[int64(uint32(v11))+4:], uint32(v5))
											store32(m.memory[uint32(v11):], uint32(v2))
											v1 = i32(1)
											store32(m.memory[int64(uint32(v3))+104:], uint32(i32(1)))
											store32(m.memory[int64(uint32(v3))+100:], uint32(v11))
											store32(m.memory[int64(uint32(v3))+96:], uint32(v6))
											t94 := int32(load32(m.memory[int64(uint32(v3))+236:]))
											store32(m.memory[int64(uint32(v3))+248:], uint32(t94))
											t95 := int64(load64(m.memory[int64(uint32(v3))+228:]))
											store64(m.memory[int64(uint32(v3))+240:], uint64(t95))
											v2 = i32(12)
										l36:
											{
												m.fn910(v3+i32(8), v3+i32(240))
												t96 := int32(load32(m.memory[int64(uint32(v3))+8:]))
												v5 = t96
												if v5 == 0 {
													t110 := int32(load32(m.memory[int64(uint32(v3))+96:]))
													v21 = t110
													t111 := int32(load32(m.memory[int64(uint32(v3))+100:]))
													v7 = t111
													t112 := int32(load32(m.memory[int64(uint32(v16))+4:]))
													t113 := int32(load32(m.memory[int64(uint32(v16))+8:]))
													m.fn46(v3, t112, t113)
													t114 := int32(load32(m.memory[int64(uint32(v3))+4:]))
													v22 = t114
													t115 := int32(load32(m.memory[uint32(v3):]))
													v23 = t115
													store32(m.memory[int64(uint32(v3))+248:], uint32(i32(50462976)))
													v8 = v1 * i32(9)
													v2 = i32(0)
													{
													l43:
														{
															if v2 == i32(4) {
																goto l41
															}
															v16 = v2 + i32(1)
															t116 := int32(m.memory[uint32(v10+v2)])
															v4 = t116
															v5 = i32(0)
															v2 = v7
															v6 = v1
														l42:
															{
																t117 := v5
																v11 = v4 & i32(255)
																t118 := int32(load32(m.memory[uint32(v2):]))
																t119 := int32(load32(m.memory[uint32(v2+i32(4)):]))
																t120 := m.fn1427(t118, t119)
																var p121 int32
																if v11 == t120&i32(255) {
																	p121 = 1
																}
																v5 = t117 + p121
																v2 = v2 + i32(8)
																v6 = v6 + i32(-1)
																if v6 != 0 {
																	goto l42
																}
															}
															v2 = v16
															if uint32(v5*i32(10)) < uint32(v8) {
																goto l43
															}
														}
														if v11 != i32(3) {
															goto l44
														}
													l41:
														v2 = v7
													l46:
														{
															if v9 == 0 {
																goto l45
															}
															t122 := int32(load32(m.memory[uint32(v2):]))
															t123 := int32(load32(m.memory[int64(uint32(v2))+4:]))
															m.fn1425(v3+i32(96), t122, t123)
															m.fn1425(v3+i32(240), v23, v22)
															t124 := int32(load32(m.memory[int64(uint32(v3))+100:]))
															v5 = t124
															t125 := int32(load32(m.memory[int64(uint32(v3))+104:]))
															t126 := int32(load32(m.memory[int64(uint32(v3))+244:]))
															t127 := v5
															v6 = t126
															t128 := int32(load32(m.memory[int64(uint32(v3))+248:]))
															t129 := m.fn191(t127, t125, v6, t128)
															v4 = t129
															t130 := int32(load32(m.memory[int64(uint32(v3))+240:]))
															m.fn16(t130, v6)
															t131 := int32(load32(m.memory[int64(uint32(v3))+96:]))
															m.fn16(t131, v5)
															v9 = v9 + i32(-8)
															v2 = v2 + i32(8)
															if v4 == 0 {
																goto l46
															}
															goto l47
														}
													l44:
														t132 := m.fn1427(v23, v22)
														if t132&i32(255) == i32(3) {
															goto l48
														}
													}
												l47:
													v20 = v20 + i32(1)
													goto l45
												l48:
													v19 = v19 + i32(1)
												l45:
													m.fn639(v21, v7)
													goto l37
												}
												t97 := int32(load32(m.memory[int64(uint32(v3))+12:]))
												v6 = t97
												{
													t98 := int32(load32(m.memory[int64(uint32(v3))+96:]))
													if v1 != t98 {
														goto l35
													}
													m.fn797(v3 + i32(96))
													t99 := int32(load32(m.memory[int64(uint32(v3))+100:]))
													v11 = t99
												}
											l35:
												v4 = v11 + v2
												store32(m.memory[uint32(v4):], uint32(v6))
												store32(m.memory[uint32(v4+i32(-4)):], uint32(v5))
												t100 := v3
												v1 = v1 + i32(1)
												store32(m.memory[int64(uint32(t100))+104:], uint32(v1))
												v9 = v9 + i32(8)
												v2 = v2 + i32(8)
												goto l36
											}
										}
									}
								}
								t78 := int32(load32(m.memory[int64(uint32(v3))+48:]))
								v4 = t78
								t79 := int32(load32(m.memory[int64(uint32(v2))+4:]))
								t80 := v3 + i32(40)
								v5 = t79
								t81 := int32(load32(m.memory[int64(uint32(v2))+8:]))
								t82 := v5
								v6 = t81
								m.fn46(t80, t82, v6)
								t83 := int32(load32(m.memory[int64(uint32(v3))+44:]))
								if t83 == 0 {
									goto l29
								}
								t84 := m.fn779(v5, v6, i32(10))
								if t84 != 0 {
									goto l30
								}
								t85 := int32(load32(m.memory[int64(uint32(v2))+4:]))
								t86 := int32(load32(m.memory[int64(uint32(v2))+8:]))
								m.fn1425(v3+i32(240), t85, t86)
								t87 := m.fn32(v3+i32(176), v3+i32(240))
								if t87 != 0 {
									goto l31
								}
								goto l30
							}
						l29:
							if v4 == 0 {
								goto l31
							}
						l30:
							v4 = i32(0)
						l39:
							m.fn38(v3 + i32(176))
						}
					l27:
						m.fn1426(v3 + i32(80))
						m.fn78(v3 + i32(68))
						goto l3
					}
					t12 := int64(load64(m.memory[int64(uint32(v3))+112:]))
					t13 := int64(load64(m.memory[int64(uint32(v3))+120:]))
					t14 := int32(load32(m.memory[int64(uint32(v2))+8:]))
					v11 = t14
					t15 := m.fn66(t12, t13, v11)
					v12 = t15
					t16 := int32(load32(m.memory[int64(uint32(v3))+100:]))
					v13 = t16
					t17 := v13
					v14 = int32(v12)
					v5 = t17 & v14
					v15 = int64(uint64(v12)>>25) & i64(127) * i64(72340172838076673)
					v2 = v2 + i32(12)
					v16 = i32(0)
					t18 := int32(load32(m.memory[int64(uint32(v3))+96:]))
					v6 = t18
				l13:
					{
						t19 := int64(load64(m.memory[uint32(v6+v5):]))
						v17 = t19
						v18 = v17 ^ v15
						v18 = (v18 ^ i64(-1)) & (v18 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
						{
							{
							l10:
								{
									if v18 == 0 {
										goto l8
									}
									v4 = v6 - (int32(uint32(int64(bits.TrailingZeros64(uint64(v18))))>>3)+v5)&v13<<3
									t20 := int32(load32(m.memory[uint32(v4+i32(-8)):]))
									if t20 == v11 {
										goto l9
									}
									v18 = (v18 + i64(-1)) & v18
									goto l10
								}
							l8:
								if v17&(v17<<1)&i64(-0x7f7f7f7f7f7f7f80) == 0 {
									t31 := v5
									v16 = v16 + i32(8)
									v5 = (t31 + v16) & v13
									goto l13
								}
								m.fn713(v3+i32(96), v10)
								t21 := int32(load32(m.memory[int64(uint32(v3))+96:]))
								v5 = t21
								t22 := int32(load32(m.memory[int64(uint32(v3))+100:]))
								t23 := v5
								t24 := v5
								v4 = t22
								t25 := m.fn26(t24, v4, v12)
								v6 = t25
								v13 = t23 + v6
								t26 := int32(m.memory[uint32(v13)])
								v16 = t26
								t27 := v13
								v14 = int32(uint32(v14) >> 25)
								m.memory[uint32(t27)] = byte(v14)
								m.memory[uint32(v5+v4&(v6+i32(-8))+i32(8))] = byte(v14)
								v4 = v5 - v6<<3
								store32(m.memory[uint32(v4+i32(-4)):], uint32(i32(0)))
								store32(m.memory[uint32(v4+i32(-8)):], uint32(v11))
								t28 := int32(load32(m.memory[int64(uint32(v3))+108:]))
								store32(m.memory[int64(uint32(v3))+108:], uint32(t28+i32(1)))
								t29 := int32(load32(m.memory[int64(uint32(v3))+104:]))
								store32(m.memory[int64(uint32(v3))+104:], uint32(t29-v16&i32(1)))
							}
						l9:
							v5 = v4 + i32(-4)
							t30 := int32(load32(m.memory[uint32(v5):]))
							store32(m.memory[uint32(v5):], uint32(t30+i32(1)))
							goto l12
						}
					}
				}
			}
			v4 = i32(0)
			t5 := int32(load32(m.memory[uint32(v2):]))
			if t5 == i32(-1) {
				goto l3
			}
			t6 := int32(load32(m.memory[int64(uint32(v2))+12:]))
			if t6 != i32(1) {
				goto l3
			}
			v5 = v5 + i32(-20)
			t7 := int32(load32(m.memory[int64(uint32(v2))+16:]))
			v6 = t7
			v2 = v2 + i32(20)
			if v6 == i32(1) {
				goto l5
			}
			goto l3
		}
	}
	goto l2
l2:
	v4 = i32(0)
l3:
	m.g0 = v3 + i32(272)
	return v4
}
func (m *Module) fn1235(v0 int32) {
	t0 := int32(load32(m.memory[int64(uint32(v0))+92:]))
	m.fn10(t0, i32(432), i32(8))
	t1 := int32(load32(m.memory[int64(uint32(v0))+64:]))
	t2 := int32(load32(m.memory[int64(uint32(v0))+68:]))
	m.fn128(t1, t2)
	m.fn933(v0 + i32(16))
}
func (m *Module) fn1236(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(v1):]))
	t2 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	t3 := m.fn1243(t0, t1, t2)
	return t3
}
func (m *Module) fn1237(v0, v1 int32) int32 {
	var v2, v3, v4 int32
	t0 := m.g0
	v2 = t0 - i32(32)
	m.g0 = v2
	t1 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	v3 = t1
	t2 := int32(load32(m.memory[uint32(v1):]))
	v4 = t2
	{
		t3 := int32(load32(m.memory[uint32(v0):]))
		v1 = t3
		t4 := int32(load32(m.memory[uint32(v1):]))
		v0 = t4
		p5 := i32(2)
		if uint32(v0) > uint32(i32(-0x7ffffff2)) {
			p5 = v0 + i32(0x7ffffff1)
		}
		switch p5 {
		default:
			store32(m.memory[int64(uint32(v2))+28:], uint32(v1+i32(4)))
			store32(m.memory[int64(uint32(v2))+16:], uint32(i32(24)))
			store32(m.memory[int64(uint32(v2))+12:], uint32(v2+i32(28)))
			t6 := m.fn284(v4, v3, i32(1051739), v2+i32(12))
			v1 = t6
			goto l19
		case 1:
			store32(m.memory[int64(uint32(v2))+28:], uint32(v1+i32(4)))
			store32(m.memory[int64(uint32(v2))+16:], uint32(i32(167)))
			store32(m.memory[int64(uint32(v2))+12:], uint32(v2+i32(28)))
			t7 := m.fn284(v4, v3, i32(0x100bb5), v2+i32(12))
			v1 = t7
			goto l19
		case 2:
			store32(m.memory[int64(uint32(v2))+28:], uint32(v1))
			store32(m.memory[int64(uint32(v2))+16:], uint32(i32(168)))
			store32(m.memory[int64(uint32(v2))+12:], uint32(v2+i32(28)))
			t8 := m.fn284(v4, v3, i32(1051635), v2+i32(12))
			v1 = t8
			goto l19
		case 3:
			store32(m.memory[int64(uint32(v2))+28:], uint32(v1+i32(4)))
			store32(m.memory[int64(uint32(v2))+16:], uint32(i32(154)))
			store32(m.memory[int64(uint32(v2))+12:], uint32(v2+i32(28)))
			t9 := m.fn284(v4, v3, i32(1051672), v2+i32(12))
			v1 = t9
			goto l19
		case 4:
			store32(m.memory[int64(uint32(v2))+28:], uint32(v1+i32(4)))
			store32(m.memory[int64(uint32(v2))+16:], uint32(i32(20)))
			store32(m.memory[int64(uint32(v2))+12:], uint32(v2+i32(28)))
			t10 := m.fn284(v4, v3, i32(1051725), v2+i32(12))
			v1 = t10
			goto l19
		case 5:
			store32(m.memory[int64(uint32(v2))+8:], uint32(v1+i32(4)))
			store32(m.memory[int64(uint32(v2))+28:], uint32(v1+i32(12)))
			store32(m.memory[int64(uint32(v2))+24:], uint32(i32(169)))
			store32(m.memory[int64(uint32(v2))+16:], uint32(i32(22)))
			store32(m.memory[int64(uint32(v2))+20:], uint32(v2+i32(28)))
			store32(m.memory[int64(uint32(v2))+12:], uint32(v2+i32(8)))
			t11 := m.fn284(v4, v3, i32(1049503), v2+i32(12))
			v1 = t11
			goto l19
		case 6:
			store32(m.memory[int64(uint32(v2))+28:], uint32(v1+i32(4)))
			store32(m.memory[int64(uint32(v2))+16:], uint32(i32(36)))
			store32(m.memory[int64(uint32(v2))+12:], uint32(v2+i32(28)))
			t12 := m.fn284(v4, v3, i32(1070047), v2+i32(12))
			v1 = t12
			goto l19
		case 7:
			t13 := int32(load32(m.memory[int64(uint32(v3))+12:]))
			t14 := m.t0[uint(t13)].(func(int32, int32, int32) int32)(v4, i32(1100038), i32(20))
			v1 = t14
			goto l19
		case 8:
			store32(m.memory[int64(uint32(v2))+28:], uint32(v1+i32(4)))
			store32(m.memory[int64(uint32(v2))+16:], uint32(i32(169)))
			store32(m.memory[int64(uint32(v2))+12:], uint32(v2+i32(28)))
			t15 := m.fn284(v4, v3, i32(1050805), v2+i32(12))
			v1 = t15
			goto l19
		case 9:
			store32(m.memory[int64(uint32(v2))+28:], uint32(v1+i32(4)))
			store32(m.memory[int64(uint32(v2))+16:], uint32(i32(170)))
			store32(m.memory[int64(uint32(v2))+12:], uint32(v2+i32(28)))
			t16 := m.fn284(v4, v3, i32(1050599), v2+i32(12))
			v1 = t16
			goto l19
		case 10:
			store32(m.memory[int64(uint32(v2))+28:], uint32(v1+i32(4)))
			store32(m.memory[int64(uint32(v2))+16:], uint32(i32(171)))
			store32(m.memory[int64(uint32(v2))+12:], uint32(v2+i32(28)))
			t17 := m.fn284(v4, v3, i32(1051080), v2+i32(12))
			v1 = t17
			goto l19
		case 11:
			store32(m.memory[int64(uint32(v2))+28:], uint32(v1+i32(4)))
			store32(m.memory[int64(uint32(v2))+16:], uint32(i32(170)))
			store32(m.memory[int64(uint32(v2))+12:], uint32(v2+i32(28)))
			t18 := m.fn284(v4, v3, i32(1050137), v2+i32(12))
			v1 = t18
			goto l19
		case 12:
			store32(m.memory[int64(uint32(v2))+28:], uint32(v1+i32(4)))
			store32(m.memory[int64(uint32(v2))+16:], uint32(i32(170)))
			store32(m.memory[int64(uint32(v2))+12:], uint32(v2+i32(28)))
			t19 := m.fn284(v4, v3, i32(1050648), v2+i32(12))
			v1 = t19
			goto l19
		case 13:
			store32(m.memory[int64(uint32(v2))+28:], uint32(v1+i32(4)))
			store32(m.memory[int64(uint32(v2))+16:], uint32(i32(170)))
			store32(m.memory[int64(uint32(v2))+12:], uint32(v2+i32(28)))
			t20 := m.fn284(v4, v3, i32(1050825), v2+i32(12))
			v1 = t20
			goto l19
		case 14:
			store32(m.memory[int64(uint32(v2))+8:], uint32(v1+i32(4)))
			store32(m.memory[int64(uint32(v2))+28:], uint32(v1+i32(8)))
			store32(m.memory[int64(uint32(v2))+24:], uint32(i32(141)))
			store32(m.memory[int64(uint32(v2))+16:], uint32(i32(141)))
			store32(m.memory[int64(uint32(v2))+20:], uint32(v2+i32(28)))
			store32(m.memory[int64(uint32(v2))+12:], uint32(v2+i32(8)))
			t21 := m.fn284(v4, v3, i32(1069564), v2+i32(12))
			v1 = t21
			goto l19
		case 15:
			store32(m.memory[int64(uint32(v2))+8:], uint32(v1+i32(16)))
			store32(m.memory[int64(uint32(v2))+28:], uint32(v1+i32(4)))
			store32(m.memory[int64(uint32(v2))+24:], uint32(i32(36)))
			store32(m.memory[int64(uint32(v2))+16:], uint32(i32(22)))
			store32(m.memory[int64(uint32(v2))+20:], uint32(v2+i32(28)))
			store32(m.memory[int64(uint32(v2))+12:], uint32(v2+i32(8)))
			t22 := m.fn284(v4, v3, i32(1052669), v2+i32(12))
			v1 = t22
			goto l19
		case 16:
			t23 := int32(load32(m.memory[int64(uint32(v3))+12:]))
			t24 := m.t0[uint(t23)].(func(int32, int32, int32) int32)(v4, i32(1099985), i32(30))
			v1 = t24
			goto l19
		case 17:
			store32(m.memory[int64(uint32(v2))+28:], uint32(v1+i32(4)))
			store32(m.memory[int64(uint32(v2))+16:], uint32(i32(36)))
			store32(m.memory[int64(uint32(v2))+12:], uint32(v2+i32(28)))
			t25 := m.fn284(v4, v3, i32(1068112), v2+i32(12))
			v1 = t25
			goto l19
		case 18:
			store32(m.memory[int64(uint32(v2))+28:], uint32(v1+i32(4)))
			store32(m.memory[int64(uint32(v2))+16:], uint32(i32(172)))
			store32(m.memory[int64(uint32(v2))+12:], uint32(v2+i32(28)))
			t26 := m.fn284(v4, v3, i32(1051649), v2+i32(12))
			v1 = t26
		}
	}
l19:
	m.g0 = v2 + i32(32)
	return v1
}
func (m *Module) fn1238(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(v1):]))
	t2 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	t3 := m.fn1262(t0, t1, t2)
	return t3
}
func (m *Module) fn1239(v0, v1 int32) int32 {
	var v2, v3, v4 int32
	t0 := m.g0
	v2 = t0 - i32(32)
	m.g0 = v2
	t1 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	v3 = t1
	t2 := int32(load32(m.memory[uint32(v1):]))
	v4 = t2
	{
		t3 := int32(load32(m.memory[uint32(v0):]))
		v1 = t3
		t4 := int32(m.memory[uint32(v1)])
		v0 = t4
		p5 := i32(0)
		if uint32(v0) > uint32(i32(6)) {
			p5 = v0 + i32(-6)
		}
		switch p5 {
		default:
			store32(m.memory[int64(uint32(v2))+12:], uint32(v1))
			store32(m.memory[int64(uint32(v2))+20:], uint32(i32(173)))
			store32(m.memory[int64(uint32(v2))+16:], uint32(v2+i32(12)))
			t6 := m.fn284(v4, v3, i32(1051711), v2+i32(16))
			v1 = t6
			goto l6
		case 1:
			store32(m.memory[int64(uint32(v2))+12:], uint32(v1+i32(4)))
			store32(m.memory[int64(uint32(v2))+20:], uint32(i32(24)))
			store32(m.memory[int64(uint32(v2))+16:], uint32(v2+i32(12)))
			t7 := m.fn284(v4, v3, i32(1051739), v2+i32(16))
			v1 = t7
			goto l6
		case 2:
			store32(m.memory[int64(uint32(v2))+12:], uint32(v1+i32(4)))
			store32(m.memory[int64(uint32(v2))+20:], uint32(i32(36)))
			store32(m.memory[int64(uint32(v2))+16:], uint32(v2+i32(12)))
			t8 := m.fn284(v4, v3, i32(1069916), v2+i32(16))
			v1 = t8
			goto l6
		case 3:
			store32(m.memory[int64(uint32(v2))+8:], uint32(v1+i32(4)))
			store32(m.memory[int64(uint32(v2))+12:], uint32(v1+i32(12)))
			store32(m.memory[int64(uint32(v2))+28:], uint32(i32(169)))
			store32(m.memory[int64(uint32(v2))+20:], uint32(i32(22)))
			store32(m.memory[int64(uint32(v2))+24:], uint32(v2+i32(12)))
			store32(m.memory[int64(uint32(v2))+16:], uint32(v2+i32(8)))
			t9 := m.fn284(v4, v3, i32(1070069), v2+i32(16))
			v1 = t9
			goto l6
		case 4:
			t10 := int32(load32(m.memory[int64(uint32(v3))+12:]))
			t11 := m.t0[uint(t10)].(func(int32, int32, int32) int32)(v4, i32(1100015), i32(23))
			v1 = t11
			goto l6
		case 5:
			store32(m.memory[int64(uint32(v2))+8:], uint32(v1+i32(2)))
			store32(m.memory[int64(uint32(v2))+12:], uint32(v1+i32(4)))
			store32(m.memory[int64(uint32(v2))+28:], uint32(i32(169)))
			store32(m.memory[int64(uint32(v2))+20:], uint32(i32(169)))
			store32(m.memory[int64(uint32(v2))+24:], uint32(v2+i32(12)))
			store32(m.memory[int64(uint32(v2))+16:], uint32(v2+i32(8)))
			t12 := m.fn284(v4, v3, i32(1050906), v2+i32(16))
			v1 = t12
		}
	}
l6:
	m.g0 = v2 + i32(32)
	return v1
}
func (m *Module) fn1240(v0, v1 int32) int32 {
	var v2, v3, v4 int32
	t0 := m.g0
	v2 = t0 - i32(32)
	m.g0 = v2
	t1 := int32(load32(m.memory[uint32(v0):]))
	v0 = t1
	v3 = v0 + i32(4)
	t2 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	v4 = t2
	t3 := int32(load32(m.memory[uint32(v1):]))
	v1 = t3
	{
		t4 := int32(m.memory[uint32(v0)])
		switch t4 {
		default:
			store32(m.memory[int64(uint32(v2))+8:], uint32(v3))
			store32(m.memory[int64(uint32(v2))+28:], uint32(v0+i32(12)))
			store32(m.memory[int64(uint32(v2))+24:], uint32(i32(174)))
			store32(m.memory[int64(uint32(v2))+16:], uint32(i32(174)))
			store32(m.memory[int64(uint32(v2))+20:], uint32(v2+i32(28)))
			store32(m.memory[int64(uint32(v2))+12:], uint32(v2+i32(8)))
			t5 := m.fn284(v1, v4, i32(1069760), v2+i32(12))
			v1 = t5
			goto l5
		case 1:
			store32(m.memory[int64(uint32(v2))+8:], uint32(v3))
			store32(m.memory[int64(uint32(v2))+28:], uint32(v0+i32(1)))
			store32(m.memory[int64(uint32(v2))+24:], uint32(i32(175)))
			store32(m.memory[int64(uint32(v2))+16:], uint32(i32(174)))
			store32(m.memory[int64(uint32(v2))+20:], uint32(v2+i32(28)))
			store32(m.memory[int64(uint32(v2))+12:], uint32(v2+i32(8)))
			t6 := m.fn284(v1, v4, i32(0x100fdd), v2+i32(12))
			v1 = t6
			goto l5
		case 2:
			store32(m.memory[int64(uint32(v2))+28:], uint32(v3))
			store32(m.memory[int64(uint32(v2))+16:], uint32(i32(174)))
			store32(m.memory[int64(uint32(v2))+12:], uint32(v2+i32(28)))
			t7 := m.fn284(v1, v4, i32(1069821), v2+i32(12))
			v1 = t7
			goto l5
		case 3:
			store32(m.memory[int64(uint32(v2))+28:], uint32(v3))
			store32(m.memory[int64(uint32(v2))+16:], uint32(i32(36)))
			store32(m.memory[int64(uint32(v2))+12:], uint32(v2+i32(28)))
			t8 := m.fn284(v1, v4, i32(1069990), v2+i32(12))
			v1 = t8
			goto l5
		case 4:
			store32(m.memory[int64(uint32(v2))+28:], uint32(v3))
			store32(m.memory[int64(uint32(v2))+16:], uint32(i32(36)))
			store32(m.memory[int64(uint32(v2))+12:], uint32(v2+i32(28)))
			t9 := m.fn284(v1, v4, i32(1052692), v2+i32(12))
			v1 = t9
		}
	}
l5:
	m.g0 = v2 + i32(32)
	return v1
}
func (m *Module) fn1241(v0, v1 int32) int32 {
	var v2, v3, v4 int32
	t0 := m.g0
	v2 = t0 - i32(32)
	m.g0 = v2
	t1 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	v3 = t1
	t2 := int32(load32(m.memory[uint32(v1):]))
	v4 = t2
	{
		t3 := int32(load32(m.memory[uint32(v0):]))
		v1 = t3
		t4 := int32(load32(m.memory[uint32(v1):]))
		v0 = t4
		p5 := i32(2)
		if uint32(v0) > uint32(i32(-0x7ffffff2)) {
			p5 = v0 + i32(0x7ffffff1)
		}
		switch p5 {
		case 4:
			panic("unreachable")
		default:
			store32(m.memory[int64(uint32(v2))+12:], uint32(v1+i32(4)))
			store32(m.memory[int64(uint32(v2))+20:], uint32(i32(24)))
			store32(m.memory[int64(uint32(v2))+16:], uint32(v2+i32(12)))
			t6 := m.fn284(v4, v3, i32(1051739), v2+i32(16))
			v1 = t6
			goto l17
		case 1:
			store32(m.memory[int64(uint32(v2))+12:], uint32(v1+i32(4)))
			store32(m.memory[int64(uint32(v2))+20:], uint32(i32(176)))
			store32(m.memory[int64(uint32(v2))+16:], uint32(v2+i32(12)))
			t7 := m.fn284(v4, v3, i32(0x100bb5), v2+i32(16))
			v1 = t7
			goto l17
		case 2:
			store32(m.memory[int64(uint32(v2))+12:], uint32(v1))
			store32(m.memory[int64(uint32(v2))+20:], uint32(i32(168)))
			store32(m.memory[int64(uint32(v2))+16:], uint32(v2+i32(12)))
			t8 := m.fn284(v4, v3, i32(1051635), v2+i32(16))
			v1 = t8
			goto l17
		case 3:
			store32(m.memory[int64(uint32(v2))+12:], uint32(v1+i32(4)))
			store32(m.memory[int64(uint32(v2))+20:], uint32(i32(154)))
			store32(m.memory[int64(uint32(v2))+16:], uint32(v2+i32(12)))
			t9 := m.fn284(v4, v3, i32(1051672), v2+i32(16))
			v1 = t9
			goto l17
		case 5:
			store32(m.memory[int64(uint32(v2))+12:], uint32(v1+i32(4)))
			store32(m.memory[int64(uint32(v2))+20:], uint32(i32(177)))
			store32(m.memory[int64(uint32(v2))+16:], uint32(v2+i32(12)))
			t10 := m.fn284(v4, v3, i32(1051549), v2+i32(16))
			v1 = t10
			goto l17
		case 6:
			store32(m.memory[int64(uint32(v2))+12:], uint32(v1+i32(4)))
			store32(m.memory[int64(uint32(v2))+20:], uint32(i32(178)))
			store32(m.memory[int64(uint32(v2))+16:], uint32(v2+i32(12)))
			t11 := m.fn284(v4, v3, i32(1051476), v2+i32(16))
			v1 = t11
			goto l17
		case 7:
			store32(m.memory[int64(uint32(v2))+12:], uint32(v1))
			store32(m.memory[int64(uint32(v2))+20:], uint32(i32(179)))
			store32(m.memory[int64(uint32(v2))+16:], uint32(v2+i32(12)))
			t12 := m.fn284(v4, v3, i32(1051614), v2+i32(16))
			v1 = t12
			goto l17
		case 8:
			store32(m.memory[int64(uint32(v2))+12:], uint32(v1+i32(4)))
			store32(m.memory[int64(uint32(v2))+20:], uint32(i32(180)))
			store32(m.memory[int64(uint32(v2))+16:], uint32(v2+i32(12)))
			t13 := m.fn284(v4, v3, i32(1052258), v2+i32(16))
			v1 = t13
			goto l17
		case 9:
			store32(m.memory[int64(uint32(v2))+12:], uint32(v1+i32(4)))
			store32(m.memory[int64(uint32(v2))+20:], uint32(i32(22)))
			store32(m.memory[int64(uint32(v2))+16:], uint32(v2+i32(12)))
			t14 := m.fn284(v4, v3, i32(1067607), v2+i32(16))
			v1 = t14
			goto l17
		case 10:
			store32(m.memory[int64(uint32(v2))+12:], uint32(v1+i32(4)))
			store32(m.memory[int64(uint32(v2))+20:], uint32(i32(22)))
			store32(m.memory[int64(uint32(v2))+16:], uint32(v2+i32(12)))
			t15 := m.fn284(v4, v3, i32(1067804), v2+i32(16))
			v1 = t15
			goto l17
		case 11:
			store32(m.memory[int64(uint32(v2))+8:], uint32(v1+i32(16)))
			store32(m.memory[int64(uint32(v2))+12:], uint32(v1+i32(4)))
			store32(m.memory[int64(uint32(v2))+28:], uint32(i32(36)))
			store32(m.memory[int64(uint32(v2))+20:], uint32(i32(22)))
			store32(m.memory[int64(uint32(v2))+24:], uint32(v2+i32(12)))
			store32(m.memory[int64(uint32(v2))+16:], uint32(v2+i32(8)))
			t16 := m.fn284(v4, v3, i32(1069962), v2+i32(16))
			v1 = t16
			goto l17
		case 12:
			t17 := int32(load32(m.memory[int64(uint32(v3))+12:]))
			t18 := m.t0[uint(t17)].(func(int32, int32, int32) int32)(v4, i32(1099985), i32(30))
			v1 = t18
			goto l17
		case 13:
			store32(m.memory[int64(uint32(v2))+12:], uint32(v1+i32(4)))
			store32(m.memory[int64(uint32(v2))+20:], uint32(i32(36)))
			store32(m.memory[int64(uint32(v2))+16:], uint32(v2+i32(12)))
			t19 := m.fn284(v4, v3, i32(1068112), v2+i32(16))
			v1 = t19
			goto l17
		case 14:
			store32(m.memory[int64(uint32(v2))+12:], uint32(v1+i32(4)))
			store32(m.memory[int64(uint32(v2))+20:], uint32(i32(154)))
			store32(m.memory[int64(uint32(v2))+16:], uint32(v2+i32(12)))
			t20 := m.fn284(v4, v3, i32(1051776), v2+i32(16))
			v1 = t20
			goto l17
		case 15:
			store32(m.memory[int64(uint32(v2))+12:], uint32(v1+i32(4)))
			store32(m.memory[int64(uint32(v2))+20:], uint32(i32(172)))
			store32(m.memory[int64(uint32(v2))+16:], uint32(v2+i32(12)))
			t21 := m.fn284(v4, v3, i32(1051753), v2+i32(16))
			v1 = t21
			goto l17
		case 16:
			store32(m.memory[int64(uint32(v2))+8:], uint32(v1+i32(4)))
			store32(m.memory[int64(uint32(v2))+12:], uint32(v1+i32(8)))
			store32(m.memory[int64(uint32(v2))+28:], uint32(i32(141)))
			store32(m.memory[int64(uint32(v2))+20:], uint32(i32(141)))
			store32(m.memory[int64(uint32(v2))+24:], uint32(v2+i32(12)))
			store32(m.memory[int64(uint32(v2))+16:], uint32(v2+i32(8)))
			t22 := m.fn284(v4, v3, i32(1069107), v2+i32(16))
			v1 = t22
		}
	}
l17:
	m.g0 = v2 + i32(32)
	return v1
}
func (m *Module) fn1242(v0 int32) {
	t0 := int32(load32(m.memory[uint32(v0):]))
	switch t0 {
	case 7:
		return
	default:
		t1 := int32(m.memory[int64(uint32(v0))+4])
		t2 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		m.fn119(t1, t2)
		return
	case 1:
		m.fn948(v0 + i32(4))
		return
	case 2:
		m.fn417(v0 + i32(4))
		return
	case 3:
		m.fn534(v0 + i32(4))
		return
	case 4:
		m.fn564(v0 + i32(4))
		return
	case 5:
		m.fn451(v0 + i32(4))
		return
	case 6:
		t3 := int32(m.memory[int64(uint32(v0))+4])
		if uint32(t3) < uint32(i32(3)) {
			return
		}
		t4 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		t5 := int32(load32(m.memory[int64(uint32(v0))+12:]))
		m.fn16(t4, t5)
	}
}
func (m *Module) fn1243(v0, v1, v2 int32) int32 {
	var v3 int32
	t0 := m.g0
	v3 = t0 - i32(48)
	m.g0 = v3
	{
		t1 := int32(m.memory[uint32(v0)])
		switch t1 {
		default:
			store32(m.memory[int64(uint32(v3))+44:], uint32(v0+i32(4)))
			store32(m.memory[int64(uint32(v3))+24:], uint32(i32(24)))
			store32(m.memory[int64(uint32(v3))+20:], uint32(v3+i32(44)))
			t2 := m.fn284(v1, v2, i32(1051739), v3+i32(20))
			v0 = t2
			goto l15
		case 1:
			store32(m.memory[int64(uint32(v3))+44:], uint32(v0+i32(4)))
			store32(m.memory[int64(uint32(v3))+24:], uint32(i32(173)))
			store32(m.memory[int64(uint32(v3))+20:], uint32(v3+i32(44)))
			t3 := m.fn284(v1, v2, i32(1051711), v3+i32(20))
			v0 = t3
			goto l15
		case 2:
			store32(m.memory[int64(uint32(v3))+44:], uint32(v0+i32(4)))
			store32(m.memory[int64(uint32(v3))+24:], uint32(i32(20)))
			store32(m.memory[int64(uint32(v3))+20:], uint32(v3+i32(44)))
			t4 := m.fn284(v1, v2, i32(1051725), v3+i32(20))
			v0 = t4
			goto l15
		case 3:
			t5 := int32(load32(m.memory[int64(uint32(v2))+12:]))
			t6 := m.t0[uint(t5)].(func(int32, int32, int32) int32)(v1, i32(1100038), i32(20))
			v0 = t6
			goto l15
		case 4:
			store32(m.memory[int64(uint32(v3))+16:], uint32(v0+i32(4)))
			store32(m.memory[int64(uint32(v3))+44:], uint32(v0+i32(1)))
			store32(m.memory[int64(uint32(v3))+32:], uint32(i32(170)))
			store32(m.memory[int64(uint32(v3))+24:], uint32(i32(22)))
			store32(m.memory[int64(uint32(v3))+28:], uint32(v3+i32(44)))
			store32(m.memory[int64(uint32(v3))+20:], uint32(v3+i32(16)))
			t7 := m.fn284(v1, v2, i32(1100058), v3+i32(20))
			v0 = t7
			goto l15
		case 5:
			t8 := int32(load32(m.memory[int64(uint32(v2))+12:]))
			t9 := m.t0[uint(t8)].(func(int32, int32, int32) int32)(v1, i32(1099985), i32(30))
			v0 = t9
			goto l15
		case 6:
			store32(m.memory[int64(uint32(v3))+12:], uint32(v0+i32(4)))
			store32(m.memory[int64(uint32(v3))+16:], uint32(v0+i32(8)))
			store32(m.memory[int64(uint32(v3))+44:], uint32(v0+i32(12)))
			store32(m.memory[int64(uint32(v3))+40:], uint32(i32(141)))
			store32(m.memory[int64(uint32(v3))+32:], uint32(i32(141)))
			store32(m.memory[int64(uint32(v3))+24:], uint32(i32(22)))
			store32(m.memory[int64(uint32(v3))+36:], uint32(v3+i32(16)))
			store32(m.memory[int64(uint32(v3))+28:], uint32(v3+i32(12)))
			store32(m.memory[int64(uint32(v3))+20:], uint32(v3+i32(44)))
			t10 := m.fn284(v1, v2, i32(1050856), v3+i32(20))
			v0 = t10
			goto l15
		case 7:
			t11 := int32(load32(m.memory[int64(uint32(v2))+12:]))
			t12 := m.t0[uint(t11)].(func(int32, int32, int32) int32)(v1, i32(1100084), i32(56))
			v0 = t12
			goto l15
		case 8:
			store32(m.memory[int64(uint32(v3))+44:], uint32(v0+i32(4)))
			store32(m.memory[int64(uint32(v3))+24:], uint32(i32(22)))
			store32(m.memory[int64(uint32(v3))+20:], uint32(v3+i32(44)))
			t13 := m.fn284(v1, v2, i32(1069861), v3+i32(20))
			v0 = t13
			goto l15
		case 9:
			store32(m.memory[int64(uint32(v3))+44:], uint32(v0+i32(4)))
			store32(m.memory[int64(uint32(v3))+24:], uint32(i32(141)))
			store32(m.memory[int64(uint32(v3))+20:], uint32(v3+i32(44)))
			t14 := m.fn284(v1, v2, i32(1069613), v3+i32(20))
			v0 = t14
			goto l15
		case 10:
			store32(m.memory[int64(uint32(v3))+44:], uint32(v0+i32(4)))
			store32(m.memory[int64(uint32(v3))+24:], uint32(i32(171)))
			store32(m.memory[int64(uint32(v3))+20:], uint32(v3+i32(44)))
			t15 := m.fn284(v1, v2, i32(1051063), v3+i32(20))
			v0 = t15
			goto l15
		case 11:
			store32(m.memory[int64(uint32(v3))+44:], uint32(v0+i32(1)))
			store32(m.memory[int64(uint32(v3))+24:], uint32(i32(170)))
			store32(m.memory[int64(uint32(v3))+20:], uint32(v3+i32(44)))
			t16 := m.fn284(v1, v2, i32(1050583), v3+i32(20))
			v0 = t16
			goto l15
		case 12:
			t17 := int32(load32(m.memory[int64(uint32(v2))+12:]))
			t18 := m.t0[uint(t17)].(func(int32, int32, int32) int32)(v1, i32(1100140), i32(14))
			v0 = t18
			goto l15
		case 13:
			store32(m.memory[int64(uint32(v3))+44:], uint32(v0+i32(4)))
			store32(m.memory[int64(uint32(v3))+24:], uint32(i32(36)))
			store32(m.memory[int64(uint32(v3))+20:], uint32(v3+i32(44)))
			t19 := m.fn284(v1, v2, i32(1068112), v3+i32(20))
			v0 = t19
			goto l15
		case 14:
			store32(m.memory[int64(uint32(v3))+44:], uint32(v0+i32(2)))
			store32(m.memory[int64(uint32(v3))+24:], uint32(i32(148)))
			store32(m.memory[int64(uint32(v3))+20:], uint32(v3+i32(44)))
			t20 := m.fn284(v1, v2, i32(1070021), v3+i32(20))
			v0 = t20
		}
	}
l15:
	m.g0 = v3 + i32(48)
	return v0
}
func (m *Module) fn1244(v0, v1 int32) {
	var v2, v3, v4, v5, v6, v7, v8, v9, v10, v11, v12 int32
	t0 := m.g0
	v2 = t0 - i32(32)
	m.g0 = v2
	t1 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	v3 = t1
	t2 := int32(load32(m.memory[int64(uint32(v1))+24:]))
	v4 = t2
	t3 := int32(load32(m.memory[int64(uint32(v1))+20:]))
	v5 = t3
	t4 := int32(load32(m.memory[int64(uint32(v1))+16:]))
	v6 = t4
	t5 := int32(load32(m.memory[int64(uint32(v1))+12:]))
	v7 = t5
	t6 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	t7 := v2
	v8 = t6
	m.fn59(t7, v8, i32(8), i32(24))
	t8 := int32(load32(m.memory[uint32(v2):]))
	v9 = t8
	t9 := v9
	v1 = v8 & i32(0x1fffffff)
	p10 := v1
	if uint32(v9) < uint32(v1) {
		p10 = t9
	}
	v10 = p10
	v1 = i32(0)
	t11 := int32(load32(m.memory[int64(uint32(v2))+4:]))
	v11 = t11
l1:
	{
		if v10 == 0 {
			goto l0
		}
		m.fn219(v2+i32(8), v3+v1)
		v12 = v11 + v1
		t12 := int64(load64(m.memory[int64(uint32(v2))+24:]))
		store64(m.memory[int64(uint32(v12))+16:], uint64(t12))
		t13 := int64(load64(m.memory[int64(uint32(v2))+16:]))
		store64(m.memory[int64(uint32(v12))+8:], uint64(t13))
		t14 := int64(load64(m.memory[int64(uint32(v2))+8:]))
		store64(m.memory[uint32(v12):], uint64(t14))
		v10 = v10 + i32(-1)
		v1 = v1 + i32(24)
		goto l1
	}
l0:
	store32(m.memory[int64(uint32(v0))+24:], uint32(v4))
	store32(m.memory[int64(uint32(v0))+20:], uint32(v5))
	store32(m.memory[int64(uint32(v0))+16:], uint32(v6))
	store32(m.memory[int64(uint32(v0))+12:], uint32(v7))
	store32(m.memory[int64(uint32(v0))+8:], uint32(v8))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v11))
	store32(m.memory[uint32(v0):], uint32(v9))
	m.g0 = v2 + i32(32)
}
func (m *Module) fn1245(v0, v1, v2 int32) {
	var v3, v4 int32
	var v5 int64
	var v6, v7, v8 int32
	var v9 int64
	t0 := m.g0
	v3 = t0 - i32(96)
	m.g0 = v3
	store32(m.memory[int64(uint32(v3))+12:], uint32(i32(-1)))
	v4 = i32(0)
	m.memory[int64(uint32(v3))+48] = byte(i32(0))
	store32(m.memory[int64(uint32(v3))+44:], uint32(v2))
	store32(m.memory[int64(uint32(v3))+40:], uint32(v1))
	store32(m.memory[int64(uint32(v3))+36:], uint32(v3+i32(12)))
	m.fn1563(v3+i32(64), v3+i32(36))
	{
		t1 := int32(load32(m.memory[int64(uint32(v3))+64:]))
		if t1 != i32(1) {
			goto l0
		}
		t2 := int64(load64(m.memory[int64(uint32(v3))+68:]))
		v5 = t2
		m.fn382(v3, i32(4), i32(4), i32(8))
		t3 := int32(load32(m.memory[uint32(v3):]))
		v2 = t3
		t4 := int32(load32(m.memory[int64(uint32(v3))+4:]))
		v6 = t4
		store64(m.memory[uint32(v6):], uint64(v5))
		store32(m.memory[int64(uint32(v3))+60:], uint32(i32(1)))
		store32(m.memory[int64(uint32(v3))+56:], uint32(v6))
		store32(m.memory[int64(uint32(v3))+52:], uint32(v2))
		t5 := int64(load64(m.memory[int64(uint32(v3))+44:]))
		store64(m.memory[int64(uint32(v3))+72:], uint64(t5))
		t6 := int64(load64(m.memory[int64(uint32(v3))+36:]))
		store64(m.memory[int64(uint32(v3))+64:], uint64(t6))
		v4 = i32(12)
		v2 = i32(1)
	l3:
		{
			m.fn1563(v3+i32(84), v3+i32(64))
			t7 := int32(load32(m.memory[int64(uint32(v3))+84:]))
			if t7 != i32(1) {
				t13 := int32(load32(m.memory[int64(uint32(v3))+52:]))
				v4 = t13
				goto l4
			}
			t8 := int32(load32(m.memory[int64(uint32(v3))+92:]))
			v1 = t8
			t9 := int32(load32(m.memory[int64(uint32(v3))+88:]))
			v7 = t9
			{
				t10 := int32(load32(m.memory[int64(uint32(v3))+52:]))
				if v2 != t10 {
					goto l2
				}
				m.fn62(v3+i32(52), v2, i32(1), i32(4), i32(8))
				t11 := int32(load32(m.memory[int64(uint32(v3))+56:]))
				v6 = t11
			}
		l2:
			v8 = v6 + v4
			store32(m.memory[uint32(v8):], uint32(v1))
			store32(m.memory[uint32(v8+i32(-4)):], uint32(v7))
			t12 := v3
			v2 = v2 + i32(1)
			store32(m.memory[int64(uint32(t12))+60:], uint32(v2))
			v4 = v4 + i32(8)
			goto l3
		}
	}
l0:
	v6 = i32(4)
	v2 = i32(0)
l4:
	{
		t14 := int32(load32(m.memory[int64(uint32(v3))+12:]))
		v1 = t14
		if v1 == i32(-1) {
			goto l5
		}
		t15 := int64(load64(m.memory[int64(uint32(v3))+28:]))
		v5 = t15
		t16 := int64(load64(m.memory[int64(uint32(v3))+16:]))
		v9 = t16
		t17 := int32(load32(m.memory[int64(uint32(v3))+24:]))
		v2 = t17
		m.fn76(v4, v6)
		store32(m.memory[int64(uint32(v0))+12:], uint32(v2))
		store64(m.memory[int64(uint32(v0))+4:], uint64(v9))
		store64(m.memory[int64(uint32(v0))+16:], uint64(v5))
		store32(m.memory[uint32(v0):], uint32(v1))
		goto l6
	}
l5:
	switch v2 {
	default:
		store32(m.memory[int64(uint32(v0))+4:], uint32(v2))
		store32(m.memory[uint32(v0):], uint32(i32(-0x7fffffdf)))
		goto l11
	case 0:
		store64(m.memory[uint32(v0):], uint64(i64(0x80000021)))
		goto l11
	case 1:
		t18 := m.fn1580(v6, i32(1), i32(0), i32(1099504))
		t19 := int64(load64(m.memory[uint32(t18):]))
		v5 = t19
		t20 := m.fn1580(v6, i32(1), i32(0), i32(1099520))
		t21 := int64(load64(m.memory[uint32(t20):]))
		v9 = t21
		store32(m.memory[uint32(v0):], uint32(i32(-1)))
		store64(m.memory[int64(uint32(v0))+12:], uint64(v9))
		store64(m.memory[int64(uint32(v0))+4:], uint64(v5))
		goto l11
	case 2:
		_ = m.fn1580(v6, i32(2), i32(1), i32(1099536))
		_ = m.fn1580(v6, i32(2), i32(0), i32(1099552))
		_ = m.fn1580(v6, i32(2), i32(1), i32(1099568))
		_ = m.fn1580(v6, i32(2), i32(0), i32(1099584))
		t26 := m.fn1580(v6, i32(2), i32(0), i32(1099600))
		t27 := int64(load64(m.memory[uint32(t26):]))
		v5 = t27
		t28 := m.fn1580(v6, i32(2), i32(1), i32(1099616))
		t29 := int64(load64(m.memory[uint32(t28):]))
		v9 = t29
		store32(m.memory[uint32(v0):], uint32(i32(-1)))
		store64(m.memory[int64(uint32(v0))+12:], uint64(v9))
		store64(m.memory[int64(uint32(v0))+4:], uint64(v5))
	}
l11:
	m.fn76(v4, v6)
l6:
	m.g0 = v3 + i32(96)
}
func (m *Module) fn1246(v0 int32) {
	var v1 int32
	var v2 int64
	t0 := m.g0
	v1 = t0 - i32(16)
	m.g0 = v1
	m.fn382(v1+i32(8), i32(1024), i32(1), i32(1))
	t1 := int64(load64(m.memory[int64(uint32(v1))+8:]))
	v2 = t1
	store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
	store64(m.memory[uint32(v0):], uint64(v2))
	m.g0 = v1 + i32(16)
}
func (m *Module) fn1247(v0 int32) {
	var v1, v2, v3, v4 int32
	m.fn227(v0)
	t0 := int32(load32(m.memory[int64(uint32(v0))+328:]))
	t1 := int32(load32(m.memory[int64(uint32(v0))+332:]))
	m.fn16(t0, t1)
	t2 := int32(load32(m.memory[int64(uint32(v0))+340:]))
	t3 := int32(load32(m.memory[int64(uint32(v0))+344:]))
	m.fn16(t2, t3)
	t4 := int32(load32(m.memory[int64(uint32(v0))+352:]))
	t5 := int32(load32(m.memory[int64(uint32(v0))+356:]))
	m.fn16(t4, t5)
	t6 := int32(load32(m.memory[int64(uint32(v0))+364:]))
	t7 := int32(load32(m.memory[int64(uint32(v0))+368:]))
	m.fn16(t6, t7)
	t8 := int32(load32(m.memory[int64(uint32(v0))+376:]))
	t9 := int32(load32(m.memory[int64(uint32(v0))+380:]))
	m.fn16(t8, t9)
	t10 := int32(load32(m.memory[int64(uint32(v0))+396:]))
	v1 = t10
	t11 := int32(load32(m.memory[int64(uint32(v0))+392:]))
	v2 = t11
	v3 = v2
l2:
	if v1 == 0 {
		goto l0
	}
	{
		t12 := int32(load32(m.memory[uint32(v3):]))
		v4 = t12
		if v4 == i32(-1) {
			goto l1
		}
		t13 := int32(load32(m.memory[uint32(v3+i32(4)):]))
		m.fn16(v4, t13)
	}
l1:
	v1 = v1 + i32(-1)
	v3 = v3 + i32(28)
	goto l2
l0:
	t14 := int32(load32(m.memory[int64(uint32(v0))+388:]))
	m.fn136(t14, v2, i32(4), i32(28))
}
func (m *Module) fn1248(v0 int32) {
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
					v7 = v6 + (i32(0)-int32(uint32(int64(bits.TrailingZeros64(uint64(v5))))>>3))*i32(24)
					t6 := int32(load32(m.memory[uint32(v7+i32(-24)):]))
					t7 := int32(load32(m.memory[uint32(v7+i32(-20)):]))
					m.fn16(t6, t7)
					t8 := int32(load32(m.memory[uint32(v7+i32(-12)):]))
					t9 := int32(load32(m.memory[uint32(v7+i32(-8)):]))
					m.fn419(t8, t9)
					v4 = v4 + i32(-1)
					v5 = (v5 + i64(-1)) & v5
					goto l4
				}
				v6 = v6 + i32(-192)
				t5 := int64(load64(m.memory[uint32(v0):]))
				v5 = (t5 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
				v0 = v0 + i32(8)
				goto l3
			}
		}
	l1:
		m.fn39(v1+i32(4), i32(24), i32(8), v2+i32(1))
		t10 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		t11 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t12 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		m.fn40(v3-t10, t11, t12)
	}
l0:
	m.g0 = v1 + i32(16)
}
func (m *Module) fn1249(v0 int32) {
	var v1 int32
	{
		t0 := int32(load32(m.memory[uint32(v0):]))
		v1 = t0
		p1 := i32(1)
		if uint32(v1) > uint32(i32(1)) {
			p1 = v1 + i32(-2)
		}
		switch p1 {
		default:
			m.fn231(v0 + i32(36))
			m.fn445(v0 + i32(12))
			return
		case 0:
			m.fn444(v0 + i32(8))
			return
		case 1:
			m.fn568(v0)
			return
		case 2:
			m.fn501(v0 + i32(8))
		}
	}
}
func (m *Module) fn1250(v0 int32) {
	m.fn228(v0)
	t0 := int32(load32(m.memory[int64(uint32(v0))+288:]))
	t1 := int32(load32(m.memory[int64(uint32(v0))+292:]))
	m.fn16(t0, t1)
}
func (m *Module) fn1251(v0, v1 int32) {
	var v2 int32
	v2 = i32(8)
	{
		t0 := int32(m.memory[uint32(v1)])
		switch t0 {
		case 9:
			goto l9
		default:
			t1 := int64(load64(m.memory[int64(uint32(v1))+8:]))
			store64(m.memory[int64(uint32(v0))+8:], uint64(t1))
			v2 = i32(0)
			goto l9
		case 1:
			t2 := math.Float64frombits(load64(m.memory[int64(uint32(v1))+8:]))
			store64(m.memory[int64(uint32(v0))+8:], math.Float64bits(t2))
			v2 = i32(1)
			goto l9
		case 2:
			t3 := int32(load32(m.memory[int64(uint32(v1))+12:]))
			store32(m.memory[int64(uint32(v0))+12:], uint32(t3))
			t4 := int64(load64(m.memory[int64(uint32(v1))+4:]))
			store64(m.memory[int64(uint32(v0))+4:], uint64(t4))
			v2 = i32(2)
			goto l9
		case 3:
			t5 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t6 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			m.fn377(v0+i32(4), t5, t6)
			v2 = i32(2)
			goto l9
		case 4:
			t7 := int32(m.memory[int64(uint32(v1))+1])
			m.memory[int64(uint32(v0))+1] = byte(t7)
			v2 = i32(3)
			goto l9
		case 5:
			t8 := int64(load64(m.memory[int64(uint32(v1))+16:]))
			store64(m.memory[int64(uint32(v0))+16:], uint64(t8))
			t9 := int64(load64(m.memory[int64(uint32(v1))+8:]))
			store64(m.memory[int64(uint32(v0))+8:], uint64(t9))
			v2 = i32(4)
			goto l9
		case 6:
			t10 := int32(load32(m.memory[int64(uint32(v1))+12:]))
			store32(m.memory[int64(uint32(v0))+12:], uint32(t10))
			t11 := int64(load64(m.memory[int64(uint32(v1))+4:]))
			store64(m.memory[int64(uint32(v0))+4:], uint64(t11))
			v2 = i32(5)
			goto l9
		case 7:
			t12 := int32(load32(m.memory[int64(uint32(v1))+12:]))
			store32(m.memory[int64(uint32(v0))+12:], uint32(t12))
			t13 := int64(load64(m.memory[int64(uint32(v1))+4:]))
			store64(m.memory[int64(uint32(v0))+4:], uint64(t13))
			v2 = i32(6)
			goto l9
		case 8:
			t14 := int32(m.memory[int64(uint32(v1))+1])
			m.memory[int64(uint32(v0))+1] = byte(t14)
			v2 = i32(7)
		}
	}
l9:
	m.memory[uint32(v0)] = byte(v2)
}
func (m *Module) fn1252(v0 int32, v1 float64) {
	var v2, v3 int32
	t0 := m.g0
	v2 = t0 - i32(64)
	m.g0 = v2
	store64(m.memory[int64(uint32(v2))+8:], math.Float64bits(v1))
	store32(m.memory[int64(uint32(v2))+20:], uint32(i32(181)))
	store32(m.memory[int64(uint32(v2))+16:], uint32(v2+i32(8)))
	m.fn73(v2+i32(36), i32(1083631), v2+i32(16))
	t1 := int32(load32(m.memory[int64(uint32(v2))+40:]))
	t2 := v2 + i32(16)
	v3 = t1
	t3 := int32(load32(m.memory[int64(uint32(v2))+44:]))
	m.fn217(t2, v3, t3)
	{
		{
			t4 := int32(m.memory[int64(uint32(v2))+16])
			if t4 != i32(1) {
				goto l0
			}
			store32(m.memory[int64(uint32(v2))+60:], uint32(i32(66)))
			store32(m.memory[int64(uint32(v2))+56:], uint32(v2+i32(8)))
			m.fn73(v0, i32(1052692), v2+i32(56))
			goto l1
		}
	l0:
		t5 := math.Float64frombits(load64(m.memory[int64(uint32(v2))+24:]))
		store64(m.memory[int64(uint32(v2))+48:], math.Float64bits(t5))
		store32(m.memory[int64(uint32(v2))+60:], uint32(i32(66)))
		store32(m.memory[int64(uint32(v2))+56:], uint32(v2+i32(48)))
		m.fn73(v0, i32(1052692), v2+i32(56))
	}
l1:
	t6 := int32(load32(m.memory[int64(uint32(v2))+36:]))
	m.fn16(t6, v3)
	m.g0 = v2 + i32(64)
}
func (m *Module) fn1253(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v1):]))
	t1 := int32(load32(m.memory[uint32(v0):]))
	t2 := int32(m.memory[uint32(t1)])
	v0 = t2 << 2
	t3 := int32(load32(m.memory[int64(uint32(v0))+1301548:]))
	t4 := int32(load32(m.memory[int64(uint32(v0))+1301516:]))
	t5 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	t6 := int32(load32(m.memory[int64(uint32(t5))+12:]))
	t7 := m.t0[uint(t6)].(func(int32, int32, int32) int32)(t0, t3, t4)
	return t7
}
func (m *Module) fn1254(v0 int32) {
	var v1 int32
	t0 := m.g0
	v1 = t0 - i32(16)
	m.g0 = v1
	store64(m.memory[int64(uint32(v1))+8:], uint64(int64(uint32(i32(182)))<<32|int64(uint32(v0))))
	m.fn91(i32(1052692), v1+i32(8), i32(1102060))
	panic("unreachable")
}
func (m *Module) fn1255(v0, v1, v2 int32) {
	var v3 int32
	t0 := i32_div_s(v1, v2)
	t1 := v0
	v3 = t0
	t2 := v3
	v1 = v1 - v3*v2
	v3 = v1 >> 31
	store32(m.memory[uint32(t1):], uint32(t2+v3))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v3&v2+v1))
}
func (m *Module) fn1256(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7, v8 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	v3 = i32(10)
	{
		t1 := int32(load32(m.memory[uint32(v0):]))
		v4 = t1
		t2 := v4
		v0 = v4 >> 31
		v5 = t2 ^ v0 - v0
		if uint32(v5) < uint32(i32(1000)) {
			goto l0
		}
		v3 = i32(10)
	l1:
		{
			v6 = v2 + i32(6) + v3
			t3 := v6 + i32(-4)
			v0 = v5
			t4 := int32(uint32(v0) / uint32(i32(10000)))
			t5 := v0
			v5 = t4
			v7 = t5 - v5*i32(10000)
			t6 := int32(uint32(v7&i32(0xffff)) / uint32(i32(100)))
			v8 = t6
			t7 := int32(load16(m.memory[int64(uint32(v8<<1))+1109319:]))
			store16(m.memory[uint32(t3):], uint16(t7))
			t8 := int32(load16(m.memory[int64(uint32((v7-v8*i32(100))&i32(0xffff)<<1))+1109319:]))
			store16(m.memory[uint32(v6+i32(-2)):], uint16(t8))
			v3 = v3 + i32(-4)
			if uint32(v0) > uint32(i32(9999999)) {
				goto l1
			}
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
		t9 := v2 + i32(6)
		v3 = v3 + i32(-2)
		t10 := int32(uint32(v5&i32(0xffff)) / uint32(i32(100)))
		t11 := t9 + v3
		t12 := v5
		v0 = t10
		t13 := int32(load16(m.memory[int64(uint32((t12-v0*i32(100))&i32(0xffff)<<1))+1109319:]))
		store16(m.memory[uint32(t11):], uint16(t13))
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
		t14 := v2 + i32(6)
		v3 = v3 + i32(-1)
		t15 := int32(m.memory[int64(uint32(v0<<1))+1109320])
		m.memory[uint32(t14+v3)] = byte(t15)
	}
l5:
	t16 := m.fn1638(v1, int32(uint32(v4^i32(-1))>>31), i32(1), i32(0), v2+i32(6)+v3, i32(10)-v3)
	v3 = t16
	m.g0 = v2 + i32(16)
	return v3
}
