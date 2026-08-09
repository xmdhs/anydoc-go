package core

import (
	"math/bits"
)

func (m *Module) fn267(v0, v1 int32) int32 {
	var v2 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	t1 := int32(load16(m.memory[uint32(v0):]))
	store16(m.memory[int64(uint32(v2))+6:], uint16(t1))
	store64(m.memory[int64(uint32(v2))+8:], uint64(int64(uint32(i32(12)))<<32|int64(uint32(v2+i32(6)))))
	t2 := int32(load32(m.memory[uint32(v1):]))
	t3 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	t4 := m.fn46(t2, t3, i32(1277164), v2+i32(8))
	v1 = t4
	m.g0 = v2 + i32(16)
	return v1
}
func (m *Module) fn268(v0, v1, v2, v3 int32) {
	var v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14, v15, v16, v17, v18 int32
	var v19 int64
	var v20, v21, v22, v23 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v1))+40:]))
		t1 := v3
		v4 = t0
		if uint32(t1) < uint32(v4) {
			goto l0
		}
		v5 = i32(1)
		v6 = v3
		{
			t2 := int32(load32(m.memory[uint32(v1):]))
			v7 = t2
			p3 := v7 + i32(-2)
			if uint32(v7) < uint32(i32(2)) {
				p3 = i32(2)
			}
			switch p3 {
			default:
				goto l1
			case 1:
				if v3 == 0 {
					goto l0
				}
				v7 = v2 + v3
				t4 := int32(m.memory[int64(uint32(v1))+4])
				v6 = t4
				if uint32(v3) > uint32(i32(3)) {
					t7 := int32(load32(m.memory[uint32(v7+i32(-4)):]))
					v1 = v6 * i32(16843009)
					v4 = t7 ^ v1
					if (i32(16843008)-v4|v4)&i32(-2139062144) == i32(-2139062144) {
						v4 = v3 - v7&i32(3)
						if uint32(v3) < uint32(i32(9)) {
							v7 = v2 + v4
						l13:
							{
								if uint32(v7) <= uint32(v2) {
									goto l0
								}
								t12 := v6
								v7 = v7 + i32(-1)
								t13 := int32(m.memory[uint32(v7)])
								if t12 != t13 {
									goto l13
								}
								goto l6
							}
						}
					l11:
						{
							if v4 < i32(8) {
								goto l10
							}
							v7 = v2 + v4
							t10 := int32(load32(m.memory[uint32(v7+i32(-8)):]))
							v3 = t10 ^ v1
							if (i32(16843008)-v3|v3)&i32(-2139062144) != i32(-2139062144) {
								goto l10
							}
							v4 = v4 + i32(-8)
							t11 := int32(load32(m.memory[uint32(v7+i32(-4)):]))
							v3 = t11 ^ v1
							if (i32(16843008)-v3|v3)&i32(-2139062144) == i32(-2139062144) {
								goto l11
							}
							goto l51
						}
					}
				l8:
					{
						if uint32(v7) <= uint32(v2) {
							goto l0
						}
						t8 := v6
						v7 = v7 + i32(-1)
						t9 := int32(m.memory[uint32(v7)])
						if t8 != t9 {
							goto l8
						}
						goto l6
					}
				}
			l5:
				{
					if uint32(v7) <= uint32(v2) {
						goto l0
					}
					t5 := v6
					v7 = v7 + i32(-1)
					t6 := int32(m.memory[uint32(v7)])
					if t5 != t6 {
						goto l5
					}
					goto l6
				}
			case 2:
				t14 := int32(load32(m.memory[int64(uint32(v1))+36:]))
				v8 = t14
				{
					if uint32(v3) < uint32(i32(16)) {
						v3 = v2 + v3
						v7 = i32(0)
						if v4 == 0 {
							goto l17
						}
						v5 = v3 - v4
						v7 = i32(0)
						v6 = v3
					l18:
						{
							t16 := v7 << 1
							v6 = v6 + i32(-1)
							t17 := int32(m.memory[uint32(v6)])
							v7 = t16 + t17
							if uint32(v5) < uint32(v6) {
								goto l18
							}
						}
					l17:
						v5 = i32(0) - v4
						t18 := int32(load32(m.memory[int64(uint32(v1))+28:]))
						v10 = t18
						t19 := int32(load32(m.memory[int64(uint32(v1))+24:]))
						v1 = t19
					l21:
						v6 = v3 + v5
						{
							if v1 != v7 {
								goto l19
							}
							t20 := m.fn1813(v6, v8, v4)
							if t20 == 0 {
								goto l19
							}
							v6 = v6 - v2
							v5 = i32(1)
							goto l1
						}
					l19:
						if uint32(v6) > uint32(v2) {
							t21 := v7
							t22 := v10
							v3 = v3 + i32(-1)
							t23 := int32(m.memory[uint32(v3)])
							t24 := int32(m.memory[uint32(v3+v5)])
							v7 = (t21-t22*t23)<<1 + t24
							goto l21
						}
						v5 = i32(0)
						goto l1
					}
					v5 = i32(1)
					t15 := int32(load32(m.memory[int64(uint32(v1))+4:]))
					v9 = t15
					if v7&i32(1) == 0 {
						if v4 != 0 {
							v11 = v2 - v4
							t25 := int32(load32(m.memory[int64(uint32(v1))+16:]))
							v12 = t25
							v13 = i32(0) - v12
							v14 = v8 + v12
							v15 = v12 - v4
							v16 = v12 ^ i32(-1)
							p26 := v4
							if uint32(v12) > uint32(v4) {
								p26 = v12
							}
							v17 = p26
							v18 = v17 - v12
							t27 := int64(load64(m.memory[int64(uint32(v1))+8:]))
							v19 = t27
							t28 := int32(m.memory[uint32(v8)])
							v20 = t28 & i32(255)
							v7 = v3
							v1 = v4
						l36:
							v21 = v1
							{
								{
									v22 = v7
									v6 = v22 - v4
									if uint32(v6) >= uint32(v3) {
										m.fn33(v6, v3, i32(1276640))
										panic("unreachable")
									}
									v1 = v4
									v7 = v6
									t29 := int32(m.memory[uint32(v2+v6)])
									t30 := v19
									v23 = t29
									if int32(i64_shr_u(t30, int64(uint32(v23))))&i32(1) != 0 {
										goto l24
									}
									goto l25
								}
							l24:
								v10 = v11 + v22
								p31 := v12
								if uint32(v21) < uint32(v12) {
									p31 = v21
								}
								v7 = p31 + i32(1)
							l29:
								if v7 == i32(1) {
									v7 = i32(0)
									if v20 != v23 {
										goto l30
									}
									if uint32(v12) < uint32(v21) {
										v21 = v13 + v21
										v7 = v15 + v22
										v1 = v18
										v10 = v14
									l35:
										if v1 == 0 {
											m.fn33(v17, v4, i32(1276688))
											panic("unreachable")
										}
										if uint32(v7) >= uint32(v3) {
											m.fn33(v7, v3, i32(1276704))
											panic("unreachable")
										}
										{
											t34 := int32(m.memory[uint32(v10)])
											t35 := int32(m.memory[uint32(v2+v7)])
											if t34 != t35 {
												v7 = v22 - v9
												v1 = v9
												goto l25
											}
											v1 = v1 + i32(-1)
											v5 = i32(1)
											v10 = v10 + i32(1)
											v7 = v7 + i32(1)
											v21 = v21 + i32(-1)
											if v21 == 0 {
												goto l1
											}
											goto l35
										}
									}
									v5 = i32(1)
									goto l1
								}
								v5 = v7 + i32(-2)
								if uint32(v5) >= uint32(v4) {
									m.fn33(v5, v4, i32(1276656))
									panic("unreachable")
								}
								{
									v5 = v6 + v7 + i32(-2)
									if uint32(v5) >= uint32(v3) {
										goto l28
									}
									v5 = v10 + v7
									v1 = v8 + v7
									v7 = v7 + i32(-1)
									t32 := int32(m.memory[uint32(v1+i32(-2))])
									t33 := int32(m.memory[uint32(v5+i32(-2))])
									if t32 == t33 {
										goto l29
									}
									goto l30
								}
							l28:
								m.fn33(v5, v3, i32(1276672))
								panic("unreachable")
							l30:
								v7 = v22 + v16 + v7
								v1 = v4
							}
						l25:
							v5 = i32(0)
							if uint32(v7) >= uint32(v4) {
								goto l36
							}
							goto l1
						}
						v6 = v3
						goto l1
					}
					if v4 != 0 {
						v20 = v2 - v4
						t36 := int32(load32(m.memory[int64(uint32(v1))+16:]))
						v12 = t36
						v14 = v12 - v4
						v15 = v8 + v12
						v18 = i32(0) - v12
						v13 = v12 + i32(1)
						v23 = v12 + i32(-1)
						v16 = v12 ^ i32(-1)
						t37 := int64(load64(m.memory[int64(uint32(v1))+8:]))
						v19 = t37
						t38 := int32(m.memory[uint32(v8)])
						v11 = t38 & i32(255)
						v7 = v3
					l50:
						{
							v21 = v7
							v6 = v21 - v4
							if uint32(v6) >= uint32(v3) {
								m.fn33(v6, v3, i32(1276576))
								panic("unreachable")
							}
							v7 = v6
							t39 := int32(m.memory[uint32(v2+v6)])
							t40 := v19
							v22 = t39
							if int32(i64_shr_u(t40, int64(uint32(v22))))&i32(1) != 0 {
								goto l38
							}
							goto l39
						}
					l38:
						if uint32(v23) >= uint32(v4) {
							if v12 == 0 {
								goto l41
							}
							m.fn33(v23, v4, i32(1276592))
							panic("unreachable")
						}
						v10 = v20 + v21
						v7 = v13
					l43:
						{
							if v7 == i32(1) {
								goto l41
							}
							v5 = v6 + v7 + i32(-2)
							if uint32(v5) >= uint32(v3) {
								m.fn33(v5, v3, i32(1276608))
								panic("unreachable")
							}
							v5 = v10 + v7
							v1 = v8 + v7
							v7 = v7 + i32(-1)
							t41 := int32(m.memory[uint32(v1+i32(-2))])
							t42 := int32(m.memory[uint32(v5+i32(-2))])
							if t41 == t42 {
								goto l43
							}
							goto l44
						}
					l41:
						v7 = i32(0)
						if v11 != v22 {
							goto l44
						}
						v7 = v12
						if uint32(v12) >= uint32(v4) {
							goto l45
						}
						v7 = v14 + v21
						v1 = v15
						v10 = v18
					l48:
						if uint32(v7) >= uint32(v3) {
							m.fn33(v7, v3, i32(1276624))
							panic("unreachable")
						}
						{
							t43 := int32(m.memory[uint32(v1)])
							t44 := int32(m.memory[uint32(v2+v7)])
							if t43 == t44 {
								v5 = i32(1)
								v7 = v7 + i32(1)
								v1 = v1 + i32(1)
								t45 := v4
								v10 = v10 + i32(-1)
								if t45+v10 != 0 {
									goto l48
								}
								goto l1
							}
							v7 = i32(0) - v10
							goto l45
						}
					l45:
						if v7 != v4 {
							v7 = v21 - v9
							goto l39
						}
						v5 = i32(1)
						goto l1
					l44:
						v7 = v21 + v16 + v7
					l39:
						v5 = i32(0)
						if uint32(v7) >= uint32(v4) {
							goto l50
						}
						goto l1
					}
					v6 = v3
					goto l1
				}
			}
		}
	l10:
		v7 = v2 + v4
		goto l51
	l51:
		{
			if uint32(v7) <= uint32(v2) {
				goto l0
			}
			t46 := v6
			v7 = v7 + i32(-1)
			t47 := int32(m.memory[uint32(v7)])
			if t46 != t47 {
				goto l51
			}
		}
	l6:
		v6 = v7 - v2
		goto l1
	}
l0:
	v5 = i32(0)
l1:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v6))
	store32(m.memory[uint32(v0):], uint32(v5))
}
func (m *Module) fn269(v0, v1 int32) int32 {
	var v2, v3 int32
	var v4 int64
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	{
		{
			t1 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			v3 = t1
			if v3&i32(0x2000000) != 0 {
				t3 := int64(load64(m.memory[uint32(v0):]))
				v4 = t3
				v0 = i32(17)
			l3:
				{
					t4 := int32(m.memory[int64(uint32(int32(v4)&i32(15)))+1099352])
					m.memory[uint32(v2+v0+i32(-2))] = byte(t4)
					v0 = v0 + i32(-1)
					v4 = int64(uint64(v4) >> 4)
					if v4 != i64(0) {
						goto l3
					}
				}
				t5 := m.fn681(v1, i32(1), i32(1123086), i32(2), v2+v0+i32(-1), i32(17)-v0)
				v0 = t5
				goto l2
			}
			if v3&i32(0x4000000) != 0 {
				goto l1
			}
			t2 := m.fn162(v0, v1)
			v0 = t2
			goto l2
		}
	l1:
		t6 := int64(load64(m.memory[uint32(v0):]))
		v4 = t6
		v0 = i32(17)
	l4:
		{
			t7 := int32(m.memory[int64(uint32(int32(v4)&i32(15)))+1123088])
			m.memory[uint32(v2+v0+i32(-2))] = byte(t7)
			v0 = v0 + i32(-1)
			v4 = int64(uint64(v4) >> 4)
			if v4 != i64(0) {
				goto l4
			}
		}
		t8 := m.fn681(v1, i32(1), i32(1123086), i32(2), v2+v0+i32(-1), i32(17)-v0)
		v0 = t8
	}
l2:
	m.g0 = v2 + i32(16)
	return v0
}
func (m *Module) fn270(v0, v1 int32) {
	var v2, v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14, v15, v16, v17, v18 int32
	t0 := m.g0
	v2 = t0 - i32(48)
	m.g0 = v2
	t1 := int32(m.memory[int64(uint32(v1))+3])
	v3 = t1
	t2 := int32(m.memory[int64(uint32(v1))+2])
	v4 = t2
	t3 := int32(m.memory[int64(uint32(v1))+1])
	v5 = t3
	t4 := int32(m.memory[uint32(v1)])
	v6 = t4
	m.memory[int64(uint32(v2))+22] = byte(i32(0))
	m.memory[int64(uint32(v2))+23] = byte(i32(1))
	v7 = i32(1)
	v8 = i32(0)
	v9 = v5
	v10 = v6
	{
		t5 := int32(m.memory[int64(uint32(v5))+1276908])
		t6 := int32(m.memory[int64(uint32(v6))+1276908])
		if uint32(t5) >= uint32(t6) {
			goto l0
		}
		v7 = i32(0)
		m.memory[int64(uint32(v2))+23] = byte(i32(0))
		v8 = i32(1)
		m.memory[int64(uint32(v2))+22] = byte(i32(1))
		v9 = v6
		v10 = v5
	}
l0:
	v11 = i32(2)
	store32(m.memory[int64(uint32(v2))+40:], uint32(i32(2)))
	store64(m.memory[int64(uint32(v2))+32:], uint64(i64(0xff00000000)))
	store32(m.memory[int64(uint32(v2))+24:], uint32(v1))
	store32(m.memory[int64(uint32(v2))+28:], uint32(v1+i32(4)))
	{
	l7:
		{
			{
				{
					if v11 != 0 {
						goto l1
					}
					t7 := int32(load32(m.memory[int64(uint32(v2))+36:]))
					v11 = t7
					if v11 == 0 {
						goto l2
					}
					store32(m.memory[int64(uint32(v2))+36:], uint32(v11+i32(-1)))
					t8 := int32(load32(m.memory[int64(uint32(v2))+24:]))
					v11 = t8
					t9 := int32(load32(m.memory[int64(uint32(v2))+28:]))
					if v11 == t9 {
						goto l2
					}
					store32(m.memory[int64(uint32(v2))+24:], uint32(v11+i32(1)))
					t10 := int32(load32(m.memory[int64(uint32(v2))+32:]))
					t11 := v2
					v12 = t10
					store32(m.memory[int64(uint32(t11))+32:], uint32(v12+i32(1)))
					goto l3
				}
			l1:
				store32(m.memory[int64(uint32(v2))+40:], uint32(i32(0)))
				m.fn1900(v2+i32(8), v2+i32(24), v11)
				t12 := int32(load32(m.memory[int64(uint32(v2))+12:]))
				v11 = t12
				if v11 == 0 {
					goto l2
				}
				t13 := int32(load32(m.memory[int64(uint32(v2))+8:]))
				v12 = t13
			}
		l3:
			t14 := int32(m.memory[uint32(v11)])
			v11 = t14
			t15 := int32(m.memory[int64(uint32(v11))+1276908])
			v13 = t15
			t16 := v13
			v14 = v10 & i32(255)
			t17 := int32(m.memory[int64(uint32(v14))+1276908])
			if uint32(t16) < uint32(t17) {
				goto l4
			}
			{
				if v11 == v14 {
					goto l5
				}
				t18 := int32(m.memory[int64(uint32(v9&i32(255)))+1276908])
				if uint32(v13) >= uint32(t18) {
					goto l5
				}
				if uint32(v12) > uint32(i32(255)) {
					m.memory[int64(uint32(v2))+47] = byte(i32(2))
					m.fn42(i32(1284936), i32(43), v2+i32(47), i32(1276204), i32(1276220))
					panic("unreachable")
				}
				m.memory[int64(uint32(v2))+23] = byte(v12)
				v7 = v12
				v9 = v11
			}
		l5:
			t19 := int32(load32(m.memory[int64(uint32(v2))+40:]))
			v11 = t19
			goto l7
		}
	l2:
		v11 = v8 & i32(255)
		if v11 != v7&i32(255) {
			{
				{
					if uint32(v11) > uint32(i32(3)) {
						m.fn33(v11, i32(4), i32(1275976))
						panic("unreachable")
					}
					t21 := int32(m.memory[uint32(v1+v11)])
					v15 = t21
					t22 := int32(m.memory[int64(uint32(v15))+1276908])
					if uint32(t22) <= uint32(i32(250)) {
						goto l11
					}
					v16 = i32(40)
					goto l12
				}
			l11:
				v12 = v7 & i32(255)
				if uint32(v12) >= uint32(i32(4)) {
					m.fn33(v12, i32(4), i32(1275992))
					panic("unreachable")
				}
				t23 := int32(m.memory[uint32(v1+v12)])
				v17 = v15<<16 | t23<<24 | v12<<8 | v11
				v16 = i32(41)
			}
		l12:
			v14 = i32(1)
			v13 = i32(0)
			v7 = i32(1)
			v12 = i32(1)
			v11 = i32(0)
		l18:
			v10 = v12
			v12 = v11 + v13
			if uint32(v12) > uint32(i32(3)) {
				m.fn33(v12, i32(4), i32(1276512))
				panic("unreachable")
			}
			{
				t24 := int32(m.memory[uint32(v1+v14)])
				v14 = t24
				t25 := int32(m.memory[uint32(v1+v12)])
				t26 := v14
				v12 = t25
				if uint32(t26) < uint32(v12) {
					goto l15
				}
				v11 = v11 + i32(1)
				{
					if uint32(v14) > uint32(v12) {
						v12 = v11 + v10
						v7 = v12 - v13
						v11 = i32(0)
						goto l17
					}
					t27 := v11
					var p28 int32
					if v11 == v7 {
						p28 = 1
					}
					v12 = p28
					p29 := t27
					if v12 != 0 {
						p29 = i32(0)
					}
					v11 = p29
					p30 := i32(0)
					if v12 != 0 {
						p30 = v7
					}
					v12 = p30 + v10
					goto l17
				}
			}
		l15:
			v7 = i32(1)
			v12 = v10 + i32(1)
			v11 = i32(0)
			v13 = v10
		l17:
			v14 = v12 + v11
			if uint32(v14) < uint32(i32(4)) {
				goto l18
			}
			v14 = i32(1)
			v9 = i32(0)
			v18 = i32(1)
			v12 = i32(1)
			v11 = i32(0)
		l23:
			v10 = v12
			v12 = v11 + v9
			if uint32(v12) > uint32(i32(3)) {
				m.fn33(v12, i32(4), i32(1276512))
				panic("unreachable")
			}
			{
				t31 := int32(m.memory[uint32(v1+v14)])
				v14 = t31
				t32 := int32(m.memory[uint32(v1+v12)])
				t33 := v14
				v12 = t32
				if uint32(t33) > uint32(v12) {
					goto l20
				}
				v11 = v11 + i32(1)
				{
					if uint32(v14) < uint32(v12) {
						v12 = v11 + v10
						v18 = v12 - v9
						v11 = i32(0)
						goto l22
					}
					t34 := v11
					var p35 int32
					if v11 == v18 {
						p35 = 1
					}
					v12 = p35
					p36 := t34
					if v12 != 0 {
						p36 = i32(0)
					}
					v11 = p36
					p37 := i32(0)
					if v12 != 0 {
						p37 = v18
					}
					v12 = p37 + v10
					goto l22
				}
			}
		l20:
			v18 = i32(1)
			v12 = v10 + i32(1)
			v11 = i32(0)
			v9 = v10
		l22:
			v14 = v12 + v11
			if uint32(v14) < uint32(i32(4)) {
				goto l23
			}
			t38 := v13
			t39 := v9
			var p40 int32
			if uint32(v13) > uint32(v9) {
				p40 = 1
			}
			v12 = p40
			p41 := t39
			if v12 != 0 {
				p41 = t38
			}
			v11 = p41
			v10 = i32(4) - v11
			p42 := v11
			if uint32(v10) > uint32(v11) {
				p42 = v10
			}
			v14 = p42
			if v11&i32(0x7ffffffe) != 0 {
				goto l24
			}
			if uint32(v11) >= uint32(i32(5)) {
				m.fn28(i32(1272348), i32(19), i32(1272048))
				panic("unreachable")
			}
			p43 := v18
			if v12 != 0 {
				p43 = v7
			}
			v12 = p43
			if uint32(v12) > uint32(v10) {
				m.fn121(i32(0), v12, v10, i32(1272064))
				panic("unreachable")
			}
			if uint32(v11) > uint32(v12) {
				goto l24
			}
			v10 = v1 + v12
			{
				if uint32(v11) < uint32(i32(2)) {
					v13 = v6
					if v11 != 0 {
						goto l28
					}
					v10 = i32(0)
					goto l29
				}
				t44 := int32(load16(m.memory[uint32(v10):]))
				t45 := int32(load16(m.memory[uint32(v1):]))
				if t44 != t45 {
					goto l24
				}
				v10 = v10 + i32(2)
				v13 = v4
				goto l28
			}
		}
		m.fn1901(v2+i32(22), v2+i32(23))
		panic("unreachable")
	l4:
		m.memory[int64(uint32(v2))+23] = byte(v8)
		{
			if uint32(v12) > uint32(i32(255)) {
				goto l9
			}
			m.memory[int64(uint32(v2))+22] = byte(v12)
			v7 = v8
			v8 = v12
			v9 = v10
			v10 = v11
			t20 := int32(load32(m.memory[int64(uint32(v2))+40:]))
			v11 = t20
			goto l7
		}
	l9:
		m.memory[int64(uint32(v2))+47] = byte(i32(2))
		m.fn42(i32(1284936), i32(43), v2+i32(47), i32(1276204), i32(1276236))
		panic("unreachable")
	l28:
		t46 := int32(m.memory[uint32(v10)])
		t47 := v14
		t48 := v12
		var p49 int32
		if t46 != v13 {
			p49 = 1
		}
		v10 = p49
		p50 := t48
		if v10 != 0 {
			p50 = t47
		}
		v12 = p50
		goto l29
	}
l24:
	v12 = v14
	v10 = i32(1)
l29:
	store32(m.memory[int64(uint32(v0))+64:], uint32(i32(4)))
	store32(m.memory[int64(uint32(v0))+60:], uint32(v1))
	store32(m.memory[int64(uint32(v0))+56:], uint32(i32(0)))
	store32(m.memory[int64(uint32(v0))+48:], uint32(v16))
	store32(m.memory[int64(uint32(v0))+44:], uint32(i32(8)))
	m.memory[int64(uint32(v0))+33] = byte(v8)
	m.memory[int64(uint32(v0))+32] = byte(v15)
	store32(m.memory[int64(uint32(v0))+28:], uint32(v17))
	store32(m.memory[int64(uint32(v0))+24:], uint32(i32(42)))
	store32(m.memory[int64(uint32(v0))+16:], uint32(v11))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v12))
	store32(m.memory[uint32(v0):], uint32(v10))
	store32(m.memory[int64(uint32(v0))+40:], uint32((v5<<1+v6<<2+v4)<<1+v3))
	store64(m.memory[int64(uint32(v0))+8:], uint64(i64_shl(i64(1), int64(uint32(v3)))|(i64_shl(i64(1), int64(uint32(v4)))|(i64_shl(i64(1), int64(uint32(v5)))|i64_shl(i64(1), int64(uint32(v6)))))))
	m.g0 = v2 + i32(48)
}
func (m *Module) fn271(v0, v1, v2 int32) {
	var v3, v4 int32
	var v5, v6 int64
	var v7, v8, v9 int32
	var v10, v11 int64
	var v12, v13 int32
	var v14, v15 int64
	var v16, v17, v18, v19 int32
	t0 := m.g0
	v3 = t0 - i32(32)
	m.g0 = v3
	{
		{
			t1 := int32(m.memory[int64(uint32(v1))+1136])
			v4 = t1
			if v4 == i32(2) {
				goto l0
			}
			t2 := int64(load64(m.memory[int64(uint32(v1))+1128:]))
			t3 := v2
			v5 = t2
			store64(m.memory[int64(uint32(t3))+8:], uint64(v5))
			{
				v6 = v5 + i64(4)
				p4 := v6
				if uint64(v6) < uint64(v5) {
					p4 = i64(-1)
				}
				t5 := int64(load64(m.memory[int64(uint32(v1))+1120:]))
				if uint64(p4) > uint64(t5) {
					goto l1
				}
				v7 = i32(0)
				t6 := int32(load32(m.memory[int64(uint32(v2))+4:]))
				v8 = t6
				v6 = int64(uint32(v8))
				{
					{
						t8 := v8
						p7 := i64(0xffffffff)
						if uint64(v5) < uint64(i64(0xffffffff)) {
							p7 = v5
						}
						v9 = t8 - int32(p7)
						p9 := v9
						if uint32(v9) > uint32(v8) {
							p9 = i32(0)
						}
						if uint32(p9) < uint32(i32(4)) {
							goto l2
						}
						t10 := int32(load32(m.memory[uint32(v2):]))
						p11 := v6
						if uint64(v5) < uint64(v6) {
							p11 = v5
						}
						t12 := int32(load32(m.memory[uint32(t10+int32(p11)):]))
						v7 = t12
						v10 = i64(0)
						v8 = i32(255)
						goto l3
					}
				l2:
					t13 := int64(load64(m.memory[int64(uint32(i32(0)))+1277248:]))
					v11 = t13
					v10 = int64(uint64(v11) >> 8)
					v8 = int32(v11)
					if v11&i64(255) != i64(255) {
						goto l4
					}
				}
			l3:
				v6 = v5 + i64(4)
			l4:
				store64(m.memory[int64(uint32(v2))+8:], uint64(v6))
				v9 = v8 & i32(255)
				if v9 == i32(255) {
					t15 := int32(load32(m.memory[int64(uint32(v1))+72:]))
					if t15 == i32(4) {
						t16 := int32(load32(m.memory[int64(uint32(v1))+68:]))
						t17 := int32(load32(m.memory[uint32(t16):]))
						if t17 != v7 {
							goto l1
						}
						store64(m.memory[int64(uint32(v2))+8:], uint64(v5))
						m.memory[int64(uint32(v1))+1136] = byte(i32(2))
						store64(m.memory[int64(uint32(v0))+16:], uint64(v5))
						store64(m.memory[int64(uint32(v0))+8:], uint64(i64(1)))
						store32(m.memory[uint32(v0):], uint32(i32(0)))
						goto l12
					}
					goto l1
				}
				v7 = int32(int64(uint64(v10) >> 24))
				v12 = int32(v10)
				v13 = v12
				switch v9 {
				default:
					goto l6
				case 2, 3:
					t14 := int32(m.memory[int64(uint32(v7))+8])
					v13 = t14
					fallthrough
				case 1:
					if v13&i32(255) != i32(37) {
						goto l6
					}
					if v4&i32(1) != 0 {
						store32(m.memory[uint32(v0):], uint32(i32(0)))
						store64(m.memory[int64(uint32(v0))+8:], uint64(i64(0)))
						m.fn272(v8, v7)
						goto l12
					}
					m.fn272(v8, v7)
					goto l10
				}
			l6:
				store32(m.memory[int64(uint32(v0))+12:], uint32(v7))
				store64(m.memory[uint32(v0):], uint64(i64(-0x7fffffffffffffff)))
				store32(m.memory[int64(uint32(v0))+8:], uint32(v12<<8|v8&i32(255)))
				goto l12
			}
		l1:
			if v4&i32(1) == 0 {
				goto l10
			}
			store32(m.memory[uint32(v0):], uint32(i32(0)))
			store64(m.memory[int64(uint32(v0))+8:], uint64(i64(0)))
			goto l12
		l10:
			m.memory[int64(uint32(v1))+1136] = byte(i32(2))
		}
	l0:
		{
			t18 := int64(load64(m.memory[int64(uint32(v1))+1104:]))
			v5 = t18
			t19 := int64(load64(m.memory[int64(uint32(v1))+1112:]))
			if uint64(v5) < uint64(t19) {
				goto l13
			}
			t20 := int64(load64(m.memory[int64(uint32(v1))+1120:]))
			t21 := v5
			v6 = t20
			if uint64(t21) >= uint64(v6) {
				goto l13
			}
			v10 = v5 + i64(1024)
			p22 := v10
			if uint64(v10) < uint64(v5) {
				p22 = i64(-1)
			}
			v10 = p22
			if uint64(v10) <= uint64(v5) {
				goto l13
			}
			{
				p23 := v10
				if uint64(v6) < uint64(v10) {
					p23 = v6
				}
				v11 = p23 - v5
				v4 = int32(v11)
				if uint32(v4) > uint32(i32(1024)) {
					goto l14
				}
				v9 = v1 + i32(80)
				t24 := int64(load64(m.memory[int64(uint32(i32(0)))+1277248:]))
				v6 = t24
				v14 = v6 & i64(255)
				v15 = int64(uint64(v6) >> 8)
				t25 := int32(load32(m.memory[uint32(v2):]))
				v12 = t25
				v16 = int32(v6)
				t26 := int32(load32(m.memory[int64(uint32(v2))+4:]))
				v13 = t26
				v10 = int64(uint32(v13))
				{
					t27 := int32(load32(m.memory[uint32(v1):]))
					if t27 == 0 {
						goto l15
					}
					t28 := int32(load32(m.memory[int64(uint32(v1))+4:]))
					t29 := v4
					v8 = t28
					if uint32(t29) < uint32(v8) {
						m.fn121(v8, v4, v4, i32(1276892))
						panic("unreachable")
					}
					v17 = v9 + v8
					v4 = v4 - v8
					goto l17
				}
			l15:
				{
					t31 := v13
					p30 := v10
					if uint64(v5) < uint64(v10) {
						p30 = v5
					}
					v8 = int32(p30)
					if uint32(t31-v8) < uint32(v4) {
						goto l18
					}
					v8 = v12 + v8
					if v4 == i32(1) {
						t32 := int32(m.memory[uint32(v8)])
						m.memory[uint32(v9)] = byte(t32)
						v8 = i32(255)
						goto l21
					}
					if v4 == 0 {
						goto l20
					}
					memory_copy(m.memory, uint32(v9), uint32(v8), uint32(v4))
				l20:
					v8 = i32(255)
					goto l21
				}
			l18:
				v8 = v16
				v6 = v10
				if v14 != i64(255) {
					goto l22
				}
			l21:
				v6 = v11&i64(2047) + v5
			l22:
				store64(m.memory[int64(uint32(v2))+8:], uint64(v6))
				if v8&i32(255) != i32(255) {
					goto l23
				}
				v8 = i32(0)
				v17 = v9
			l17:
				v18 = v1 + i32(8)
				store64(m.memory[int64(uint32(v3))+24:], uint64(i64(1)))
				{
					t33 := int32(load32(m.memory[int64(uint32(v1))+72:]))
					t34 := v4
					v7 = t33
					if uint32(t34) < uint32(v7) {
						goto l24
					}
					t35 := int32(load32(m.memory[int64(uint32(v1))+68:]))
					t36 := int32(load32(m.memory[int64(uint32(v1))+56:]))
					m.t0[uint(t36)].(func(int32, int32, int32, int32, int32, int32, int32))(v3+i32(16), v18, v3+i32(24), v17, v4, t35, v7)
					t37 := int32(load32(m.memory[int64(uint32(v3))+16:]))
					if t37&i32(1) == 0 {
						goto l24
					}
					t38 := int32(load32(m.memory[int64(uint32(v3))+20:]))
					v4 = t38
					goto l25
				}
			l24:
				store32(m.memory[uint32(v1):], uint32(i32(0)))
				{
					t39 := int64(load64(m.memory[int64(uint32(v1))+1104:]))
					v5 = t39
					v6 = v5 + i64(1021)
					p40 := v6
					if uint64(v6) < uint64(v5) {
						p40 = i64(-1)
					}
					v5 = p40
					t41 := int64(load64(m.memory[int64(uint32(v1))+1120:]))
					t42 := v5
					v6 = t41
					if uint64(t42) >= uint64(v6) {
						goto l26
					}
					store64(m.memory[int64(uint32(v1))+1104:], uint64(v5))
					t43 := int64(load64(m.memory[int64(uint32(v1))+1112:]))
					if uint64(v5) < uint64(t43) {
						goto l13
					}
					t44 := int32(load32(m.memory[int64(uint32(v1))+56:]))
					v17 = t44
					t45 := int32(load32(m.memory[int64(uint32(v1))+68:]))
					v19 = t45
				l33:
					{
						if uint64(v5) >= uint64(v6) {
							goto l13
						}
						v11 = v5 + i64(1024)
						p46 := v11
						if uint64(v11) < uint64(v5) {
							p46 = i64(-1)
						}
						v11 = p46
						if uint64(v11) <= uint64(v5) {
							goto l13
						}
						p47 := v11
						if uint64(v6) < uint64(v11) {
							p47 = v6
						}
						v11 = p47 - v5
						v4 = int32(v11)
						if uint32(v4) >= uint32(i32(1025)) {
							goto l14
						}
						{
							t49 := v13
							p48 := v10
							if uint64(v5) < uint64(v10) {
								p48 = v5
							}
							v8 = int32(p48)
							if uint32(t49-v8) < uint32(v4) {
								goto l27
							}
							v8 = v12 + v8
							if v4 == i32(1) {
								t50 := int32(m.memory[uint32(v8)])
								m.memory[uint32(v9)] = byte(t50)
								v8 = i32(255)
								goto l30
							}
							if v4 == 0 {
								goto l29
							}
							memory_copy(m.memory, uint32(v9), uint32(v8), uint32(v4))
						l29:
							v8 = i32(255)
							goto l30
						}
					l27:
						v8 = v16
						v6 = v10
						if v14 != i64(255) {
							goto l31
						}
					l30:
						v6 = v11&i64(2047) + v5
					l31:
						store64(m.memory[int64(uint32(v2))+8:], uint64(v6))
						if v8&i32(255) != i32(255) {
							goto l23
						}
						store64(m.memory[int64(uint32(v3))+24:], uint64(i64(1)))
						{
							if uint32(v7) > uint32(v4) {
								goto l32
							}
							m.t0[uint(v17)].(func(int32, int32, int32, int32, int32, int32, int32))(v3+i32(8), v18, v3+i32(24), v9, v4, v19, v7)
							t51 := int32(load32(m.memory[int64(uint32(v3))+8:]))
							if t51&i32(1) == 0 {
								goto l32
							}
							t52 := int32(load32(m.memory[int64(uint32(v3))+12:]))
							v4 = t52
							v8 = i32(0)
							goto l25
						}
					l32:
						store32(m.memory[uint32(v1):], uint32(i32(0)))
						t53 := int64(load64(m.memory[int64(uint32(v1))+1104:]))
						v5 = t53
						v6 = v5 + i64(1021)
						p54 := v6
						if uint64(v6) < uint64(v5) {
							p54 = i64(-1)
						}
						v5 = p54
						t55 := int64(load64(m.memory[int64(uint32(v1))+1120:]))
						t56 := v5
						v6 = t55
						if uint64(t56) >= uint64(v6) {
							goto l26
						}
						store64(m.memory[int64(uint32(v1))+1104:], uint64(v5))
						t57 := int64(load64(m.memory[int64(uint32(v1))+1112:]))
						if uint64(v5) >= uint64(t57) {
							goto l33
						}
						goto l13
					}
				}
			l26:
				store64(m.memory[int64(uint32(v1))+1112:], uint64(v6))
				goto l13
			}
		l14:
			m.fn121(i32(0), v4, i32(1024), i32(1068408))
			panic("unreachable")
		l23:
			v1 = int32(int64(uint64(v15) >> 24))
			v4 = int32(v15)
			v2 = v4
			v8 = v8 & i32(255)
			switch v8 {
			default:
				goto l34
			case 2:
				t58 := int32(m.memory[int64(uint32(v1))+8])
				v2 = t58
				fallthrough
			case 1:
				if v2&i32(255) == i32(37) {
					goto l13
				}
				goto l34
			case 3:
				t59 := int32(m.memory[int64(uint32(v1))+8])
				if t59 != i32(37) {
					goto l34
				}
				t60 := int32(load32(m.memory[uint32(v1):]))
				v4 = t60
				{
					t61 := int32(load32(m.memory[uint32(v1+i32(4)):]))
					v2 = t61
					t62 := int32(load32(m.memory[uint32(v2):]))
					v8 = t62
					if v8 == 0 {
						goto l38
					}
					m.t0[uint(v8)].(func(int32))(v4)
				}
			l38:
				{
					t63 := int32(load32(m.memory[int64(uint32(v2))+4:]))
					v8 = t63
					if v8 == 0 {
						goto l39
					}
					t64 := int32(load32(m.memory[int64(uint32(v2))+8:]))
					m.fn21(v4, v8, t64)
				}
			l39:
				m.fn21(v1, i32(12), i32(4))
			}
		}
	l13:
		store64(m.memory[int64(uint32(v0))+8:], uint64(i64(0)))
		store32(m.memory[uint32(v0):], uint32(i32(0)))
		goto l12
	l34:
		store32(m.memory[int64(uint32(v0))+12:], uint32(v1))
		store32(m.memory[int64(uint32(v0))+4:], uint32(i32(-0x80000000)))
		store32(m.memory[int64(uint32(v0))+8:], uint32(v4<<8|v8))
		store32(m.memory[uint32(v0):], uint32(i32(1)))
		goto l12
	l25:
		store32(m.memory[uint32(v1):], uint32(i32(1)))
		store64(m.memory[int64(uint32(v0))+8:], uint64(i64(1)))
		store32(m.memory[int64(uint32(v1))+4:], uint32(v8+v4+i32(4)))
		t65 := v2
		v5 = v5 + int64(uint32(v8)) + int64(uint32(v4))
		store64(m.memory[int64(uint32(t65))+8:], uint64(v5))
		store64(m.memory[int64(uint32(v0))+16:], uint64(v5))
		store32(m.memory[uint32(v0):], uint32(i32(0)))
	}
l12:
	m.g0 = v3 + i32(32)
}
func (m *Module) fn272(v0, v1 int32) {
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
				m.fn7(i32(1274404), i32(46), i32(1274452))
				panic("unreachable")
			}
			if v3 == 0 {
				goto l4
			}
			if uint32(v4) > uint32(v2+i32(39)) {
				m.fn7(i32(1274468), i32(46), i32(1274516))
				panic("unreachable")
			}
		l4:
			m.fn5(v0)
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
			m.fn7(i32(1274404), i32(46), i32(1274452))
			panic("unreachable")
		}
		if v0 == 0 {
			goto l7
		}
		if uint32(v2) >= uint32(i32(52)) {
			m.fn7(i32(1274468), i32(46), i32(1274516))
			panic("unreachable")
		}
	l7:
		m.fn5(v1)
	}
}
func (m *Module) fn273(v0, v1, v2 int32) {
	var v3, v4 int32
	var v5, v6 int64
	var v7 int32
	var v8, v9 int64
	var v10, v11 int32
	{
		{
			if v2 != 0 {
				goto l0
			}
			v3 = i32(1)
			goto l1
		l0:
			t0 := m.fn11(v2)
			v3 = t0
			if v3 == 0 {
				m.fn16(i32(1), v2)
				panic("unreachable")
			}
			t1 := int32(m.memory[uint32(v3+i32(-4))])
			if t1&i32(3) == 0 {
				goto l1
			}
			if v2 == 0 {
				goto l1
			}
			memory_zero(m.memory, uint32(v3), uint32(v2))
		}
	l1:
		t2 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		v4 = t2
		v5 = int64(uint32(v4))
		{
			t3 := int64(load64(m.memory[int64(uint32(v1))+8:]))
			t4 := v2
			t5 := v4
			v6 = t3
			p6 := i64(0xffffffff)
			if uint64(v6) < uint64(i64(0xffffffff)) {
				p6 = v6
			}
			v7 = t5 - int32(p6)
			p7 := v7
			if uint32(v7) > uint32(v4) {
				p7 = i32(0)
			}
			if uint32(t4) > uint32(p7) {
				t11 := int64(load64(m.memory[int64(uint32(i32(0)))+1277248:]))
				v8 = t11
				v9 = int64(uint64(v8) >> 8)
				v4 = int32(v8)
				if v8&i64(255) != i64(255) {
					goto l6
				}
				goto l7
			}
			t8 := int32(load32(m.memory[uint32(v1):]))
			p9 := v5
			if uint64(v6) < uint64(v5) {
				p9 = v6
			}
			v4 = t8 + int32(p9)
			if v2 == i32(1) {
				t10 := int32(m.memory[uint32(v4)])
				m.memory[uint32(v3)] = byte(t10)
				goto l5
			}
			if v2 == 0 {
				goto l5
			}
			memory_copy(m.memory, uint32(v3), uint32(v4), uint32(v2))
			goto l5
		}
	}
l5:
	v9 = i64(0)
	v4 = i32(255)
l7:
	v5 = v6 + int64(uint32(v2))
l6:
	store64(m.memory[int64(uint32(v1))+8:], uint64(v5))
	{
		v7 = v4 & i32(255)
		if v7 == i32(255) {
			store32(m.memory[int64(uint32(v0))+8:], uint32(v2))
			store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
			store32(m.memory[uint32(v0):], uint32(i32(-2)))
			return
		}
		v1 = int32(int64(uint64(v9) >> 24))
		v10 = int32(v9)
		v11 = v10
		switch v7 {
		default:
			goto l9
		case 2:
			t12 := int32(m.memory[int64(uint32(v1))+8])
			v11 = t12
			fallthrough
		case 1:
			if v11&i32(255) != i32(37) {
				goto l9
			}
			store32(m.memory[int64(uint32(v0))+8:], uint32(i32(50)))
			store32(m.memory[int64(uint32(v0))+4:], uint32(i32(1069376)))
			store32(m.memory[uint32(v0):], uint32(i32(-1)))
			goto l13
		case 3:
			t13 := int32(m.memory[int64(uint32(v1))+8])
			if t13 == i32(37) {
				goto l14
			}
		}
	l9:
		store32(m.memory[int64(uint32(v0))+8:], uint32(v1))
		store32(m.memory[uint32(v0):], uint32(i32(-0x80000000)))
		store32(m.memory[int64(uint32(v0))+4:], uint32(v10<<8|v4&i32(255)))
		goto l13
	l14:
		store32(m.memory[int64(uint32(v0))+8:], uint32(i32(50)))
		store32(m.memory[int64(uint32(v0))+4:], uint32(i32(1069376)))
		store32(m.memory[uint32(v0):], uint32(i32(-1)))
		t14 := int32(load32(m.memory[uint32(v1):]))
		v0 = t14
		{
			t15 := int32(load32(m.memory[uint32(v1+i32(4)):]))
			v4 = t15
			t16 := int32(load32(m.memory[uint32(v4):]))
			v7 = t16
			if v7 == 0 {
				goto l15
			}
			m.t0[uint(v7)].(func(int32))(v0)
		}
	l15:
		{
			t17 := int32(load32(m.memory[int64(uint32(v4))+4:]))
			v4 = t17
			if v4 == 0 {
				goto l16
			}
			t18 := int32(load32(m.memory[uint32(v0+i32(-4)):]))
			v7 = t18
			v11 = v7 & i32(-8)
			t19 := v11
			v7 = v7 & i32(3)
			p20 := i32(8)
			if v7 != 0 {
				p20 = i32(4)
			}
			if uint32(t19) < uint32(p20+v4) {
				m.fn7(i32(1274404), i32(46), i32(1274452))
				panic("unreachable")
			}
			if v7 == 0 {
				goto l18
			}
			if uint32(v11) > uint32(v4+i32(39)) {
				m.fn7(i32(1274468), i32(46), i32(1274516))
				panic("unreachable")
			}
		l18:
			m.fn5(v0)
		}
	l16:
		t21 := int32(load32(m.memory[uint32(v1+i32(-4)):]))
		v0 = t21
		v4 = v0 & i32(-8)
		t22 := v4
		v0 = v0 & i32(3)
		p23 := i32(20)
		if v0 != 0 {
			p23 = i32(16)
		}
		if uint32(t22) < uint32(p23) {
			m.fn7(i32(1274404), i32(46), i32(1274452))
			panic("unreachable")
		}
		if v0 == 0 {
			goto l21
		}
		if uint32(v4) >= uint32(i32(52)) {
			m.fn7(i32(1274468), i32(46), i32(1274516))
			panic("unreachable")
		}
	l21:
		m.fn5(v1)
	}
l13:
	{
		if v2 == 0 {
			return
		}
		t24 := int32(load32(m.memory[uint32(v3+i32(-4)):]))
		v1 = t24
		v0 = v1 & i32(-8)
		t25 := v0
		v1 = v1 & i32(3)
		p26 := i32(8)
		if v1 != 0 {
			p26 = i32(4)
		}
		if uint32(t25) < uint32(p26+v2) {
			m.fn7(i32(1274404), i32(46), i32(1274452))
			panic("unreachable")
		}
		if v1 == 0 {
			goto l25
		}
		if uint32(v0) > uint32(v2+i32(39)) {
			m.fn7(i32(1274468), i32(46), i32(1274516))
			panic("unreachable")
		}
	l25:
		m.fn5(v3)
	}
}
func (m *Module) fn274(v0, v1, v2 int32) {
	var v3, v4, v5, v6, v7, v8, v9, v10 int32
	var v11 int64
	t0 := m.g0
	v3 = t0 - i32(48)
	m.g0 = v3
	v4 = i32(0)
	{
		{
			{
			l1:
				{
					if v2 == v4 {
						goto l0
					}
					v5 = v1 + v4
					v4 = v4 + i32(1)
					t1 := int32(int8(m.memory[uint32(v5)]))
					if t1 >= i32(0) {
						goto l1
					}
				}
				store32(m.memory[int64(uint32(v3))+36:], uint32(i32(0)))
				store64(m.memory[int64(uint32(v3))+28:], uint64(i64(0x100000000)))
				m.fn284(v3+i32(28), i32(0), v2)
			l8:
				{
					{
						t2 := int32(int8(m.memory[uint32(v1)]))
						v4 = t2
						var p3 int32
						if v4 > i32(-1) {
							p3 = 1
						}
						v6 = p3
						if v6 != 0 {
							goto l2
						}
						v4 = v4 & i32(127) << 2
						t4 := int32(load32(m.memory[int64(uint32(v4))+1293344:]))
						v5 = t4
						t5 := int32(load32(m.memory[int64(uint32(v4))+1292832:]))
						v4 = t5
						goto l3
					}
				l2:
					v5 = i32(1)
				l3:
					{
						t6 := int32(load32(m.memory[int64(uint32(v3))+28:]))
						t7 := int32(load32(m.memory[int64(uint32(v3))+36:]))
						t8 := v5
						v7 = t7
						if uint32(t8) <= uint32(t6-v7) {
							goto l4
						}
						m.fn284(v3+i32(28), v7, v5)
					}
				l4:
					t9 := int32(load32(m.memory[int64(uint32(v3))+32:]))
					v8 = t9
					v9 = v8 + v7
					if v6 != 0 {
						goto l5
					}
					v6 = int32(uint32(v4) >> 6)
					v10 = v4&i32(63) | i32(-128)
					if uint32(v4) >= uint32(i32(2048)) {
						m.memory[int64(uint32(v9))+2] = byte(v10)
						m.memory[int64(uint32(v9))+1] = byte(v6 | i32(128))
						m.memory[uint32(v9)] = byte(int32(uint32(v4)>>12) | i32(224))
						goto l7
					}
					m.memory[int64(uint32(v9))+1] = byte(v10)
					m.memory[uint32(v9)] = byte(v6 | i32(192))
					goto l7
				l5:
					m.memory[uint32(v9)] = byte(v4)
				l7:
					t10 := v3
					v4 = v7 + v5
					store32(m.memory[int64(uint32(t10))+36:], uint32(v4))
					v1 = v1 + i32(1)
					v2 = v2 + i32(-1)
					if v2 != 0 {
						goto l8
					}
				}
				t11 := int32(load32(m.memory[int64(uint32(v3))+28:]))
				v5 = t11
				goto l9
			}
		l0:
			m.fn14(v3+i32(4), v1, v2)
			t12 := int32(load32(m.memory[int64(uint32(v3))+4:]))
			if t12 != 0 {
				goto l10
			}
			t13 := int32(load32(m.memory[int64(uint32(v3))+12:]))
			v4 = t13
			t14 := int32(load32(m.memory[int64(uint32(v3))+8:]))
			v8 = t14
			v5 = i32(-1)
		}
	l9:
		store32(m.memory[int64(uint32(v0))+8:], uint32(v4))
		store32(m.memory[int64(uint32(v0))+4:], uint32(v8))
		store32(m.memory[uint32(v0):], uint32(v5))
		goto l11
	l10:
		t15 := int64(load64(m.memory[int64(uint32(v3))+8:]))
		store64(m.memory[int64(uint32(v3))+16:], uint64(t15))
		store64(m.memory[int64(uint32(v3))+40:], uint64(int64(uint32(i32(43)))<<32|int64(uint32(v3+i32(16)))))
		m.fn17(v3+i32(28), i32(1052449), v3+i32(40))
		m.fn1898(v3+i32(40), v3+i32(28))
		t16 := int64(load64(m.memory[int64(uint32(v3))+40:]))
		v11 = t16
		store32(m.memory[uint32(v0):], uint32(i32(-2)))
		store64(m.memory[int64(uint32(v0))+4:], uint64(v11))
	}
l11:
	m.g0 = v3 + i32(48)
}
func (m *Module) fn275(v0, v1 int32) {
	var v2 int32
	var v3, v4 int64
	var v5 int32
	var v6, v7, v8 int64
	t0 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	v2 = t0
	v3 = int64(uint32(v2))
	{
		{
			t1 := int64(load64(m.memory[int64(uint32(v1))+8:]))
			t2 := v2
			v4 = t1
			p3 := i64(0xffffffff)
			if uint64(v4) < uint64(i64(0xffffffff)) {
				p3 = v4
			}
			v5 = t2 - int32(p3)
			p4 := v5
			if uint32(v5) > uint32(v2) {
				p4 = i32(0)
			}
			if uint32(p4) < uint32(i32(8)) {
				goto l0
			}
			t5 := int32(load32(m.memory[uint32(v1):]))
			p6 := v3
			if uint64(v4) < uint64(v3) {
				p6 = v4
			}
			t7 := int64(load64(m.memory[uint32(t5+int32(p6)):]))
			v6 = t7
			v2 = i32(255)
			v7 = i64(0)
			goto l1
		}
	l0:
		t8 := int64(load64(m.memory[int64(uint32(i32(0)))+1277248:]))
		v8 = t8
		v7 = int64(uint64(v8) >> 8)
		v2 = int32(v8)
		v6 = i64(0)
		if v8&i64(255) != i64(255) {
			goto l2
		}
	}
l1:
	v3 = v4 + i64(8)
l2:
	store64(m.memory[int64(uint32(v1))+8:], uint64(v3))
	if v2&i32(255) == i32(255) {
		goto l3
	}
	store64(m.memory[int64(uint32(v0))+4:], uint64(v7<<8|int64(uint32(v2))&i64(255)))
	store32(m.memory[uint32(v0):], uint32(i32(1)))
	return
l3:
	store64(m.memory[int64(uint32(v0))+8:], uint64(v6))
	store32(m.memory[uint32(v0):], uint32(i32(0)))
}
func (m *Module) fn276(v0 int32) {
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
	m.fn287(t2, t4, t3, v2, i32(8), i32(32))
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
func (m *Module) fn277(v0, v1, v2 int32) {
	var v3 int32
	var v4, v5 int64
	var v6 int32
	var v7, v8, v9 int64
	var v10, v11, v12, v13 int32
	{
		{
			t0 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			v3 = t0
			t1 := int64(load64(m.memory[int64(uint32(v1))+8:]))
			t2 := v3
			v4 = t1
			t3 := v4
			v5 = int64(uint32(v3))
			p4 := v5
			if uint64(v4) < uint64(v5) {
				p4 = t3
			}
			if t2 == int32(p4) {
				goto l0
			}
			v6 = i32(255)
			v7 = i64(0)
			goto l1
		}
	l0:
		t5 := int64(load64(m.memory[int64(uint32(i32(0)))+1277248:]))
		v8 = t5
		v7 = int64(uint64(v8) >> 8)
		v6 = int32(v8)
		v9 = v5
		if v8&i64(255) != i64(255) {
			goto l2
		}
	}
l1:
	v9 = v4 + i64(1)
l2:
	t6 := int32(load32(m.memory[uint32(v1):]))
	v10 = t6
	store64(m.memory[int64(uint32(v1))+8:], uint64(v9))
	if v6&i32(255) == i32(255) {
		{
			{
				t8 := v3
				p7 := v5
				if uint64(v9) < uint64(v5) {
					p7 = v9
				}
				v6 = int32(p7)
				if uint32(t8-v6) < uint32(i32(4)) {
					goto l4
				}
				t9 := int32(load32(m.memory[uint32(v10+v6):]))
				v11 = t9
				v6 = i32(255)
				v7 = i64(0)
				goto l5
			}
		l4:
			v11 = i32(0)
			t10 := int64(load64(m.memory[int64(uint32(i32(0)))+1277248:]))
			v8 = t10
			v7 = int64(uint64(v8) >> 8)
			v6 = int32(v8)
			v4 = v5
			if v8&i64(255) != i64(255) {
				goto l6
			}
		}
	l5:
		v4 = v9 + i64(4)
	l6:
		store64(m.memory[int64(uint32(v1))+8:], uint64(v4))
		{
			if v6&i32(255) == i32(255) {
				v12 = v2 & i32(0xffff)
				if uint32(v12) < uint32(i32(5)) {
					store32(m.memory[int64(uint32(v0))+12:], uint32(i32(32)))
					store32(m.memory[int64(uint32(v0))+8:], uint32(i32(1276008)))
					store64(m.memory[uint32(v0):], uint64(i64(-0xffffffff)))
					return
				}
				v6 = v12 + i32(-5)
				if v6 == 0 {
					t20 := v10
					p19 := v5
					if uint64(v4) < uint64(v5) {
						p19 = v4
					}
					v3 = t20 + int32(p19)
					v2 = i32(1)
					goto l13
				}
				{
					t12 := m.fn11(v6)
					v2 = t12
					if v2 == 0 {
						m.fn16(i32(1), v6)
						panic("unreachable")
					}
					{
						t13 := int32(m.memory[uint32(v2+i32(-4))])
						if t13&i32(3) == 0 {
							goto l11
						}
						if v6 == 0 {
							goto l11
						}
						memory_zero(m.memory, uint32(v2), uint32(v6))
					}
				l11:
					{
						t15 := v6
						t16 := v3
						p14 := v5
						if uint64(v4) < uint64(v5) {
							p14 = v4
						}
						v13 = int32(p14)
						if uint32(t15) > uint32(t16-v13) {
							t18 := int64(load64(m.memory[int64(uint32(i32(0)))+1277248:]))
							v7 = t18
							v9 = int64(uint64(v7) >> 8)
							v3 = int32(v7)
							if v7&i64(255) != i64(255) {
								goto l15
							}
							goto l16
						}
						v3 = v10 + v13
						if v6 != i32(1) {
							goto l13
						}
						t17 := int32(m.memory[uint32(v3)])
						m.memory[uint32(v2)] = byte(t17)
						goto l14
					}
				}
			l13:
				if v6 == 0 {
					goto l14
				}
				memory_copy(m.memory, uint32(v2), uint32(v3), uint32(v6))
			l14:
				v9 = i64(0)
				v3 = i32(255)
			l16:
				v5 = v4 + int64(uint32(v6))
			l15:
				store64(m.memory[int64(uint32(v1))+8:], uint64(v5))
				{
					if v3&i32(255) == i32(255) {
						goto l17
					}
					store64(m.memory[uint32(v0):], uint64(i64(-0x7fffffffffffffff)))
					store64(m.memory[int64(uint32(v0))+8:], uint64(v9<<8|int64(uint32(v3))&i64(255)))
					if v6 == 0 {
						return
					}
					t21 := int32(load32(m.memory[uint32(v2+i32(-4)):]))
					v1 = t21
					v0 = v1 & i32(-8)
					t22 := v0
					v1 = v1 & i32(3)
					p23 := i32(8)
					if v1 != 0 {
						p23 = i32(4)
					}
					if uint32(t22) < uint32(p23+v6) {
						m.fn7(i32(1274404), i32(46), i32(1274452))
						panic("unreachable")
					}
					if v1 == 0 {
						goto l20
					}
					if uint32(v0) > uint32(v12+i32(34)) {
						m.fn7(i32(1274468), i32(46), i32(1274516))
						panic("unreachable")
					}
				l20:
					m.fn5(v2)
					return
				}
			l17:
				store32(m.memory[int64(uint32(v0))+12:], uint32(v11))
				store32(m.memory[int64(uint32(v0))+8:], uint32(v6))
				store32(m.memory[int64(uint32(v0))+4:], uint32(v2))
				store32(m.memory[uint32(v0):], uint32(i32(0)))
				return
			}
			m.memory[int64(uint32(v0))+8] = byte(v6)
			store64(m.memory[uint32(v0):], uint64(i64(-0x7fffffffffffffff)))
			t11 := v0
			v1 = int32(v7)
			store16(m.memory[int64(uint32(t11))+9:], uint16(v1))
			store32(m.memory[int64(uint32(v0))+12:], uint32(int64(uint64(v7)>>24)))
			m.memory[uint32(v0+i32(11))] = byte(int32(uint32(v1) >> 16))
			return
		}
	}
	store64(m.memory[uint32(v0):], uint64(i64(-0x7fffffffffffffff)))
	store64(m.memory[int64(uint32(v0))+8:], uint64(v7<<8|int64(uint32(v6))&i64(255)))
}
func (m *Module) fn278(v0, v1 int32) {
	var v2, v3 int32
	var v4, v5 int64
	var v6 int32
	var v7, v8 int64
	v2 = i32(0)
	t0 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	v3 = t0
	v4 = int64(uint32(v3))
	{
		{
			t1 := int64(load64(m.memory[int64(uint32(v1))+8:]))
			t2 := v3
			v5 = t1
			p3 := i64(0xffffffff)
			if uint64(v5) < uint64(i64(0xffffffff)) {
				p3 = v5
			}
			v6 = t2 - int32(p3)
			p4 := v6
			if uint32(v6) > uint32(v3) {
				p4 = i32(0)
			}
			if uint32(p4) < uint32(i32(2)) {
				goto l0
			}
			t5 := int32(load32(m.memory[uint32(v1):]))
			p6 := v4
			if uint64(v5) < uint64(v4) {
				p6 = v5
			}
			t7 := int32(load16(m.memory[uint32(t5+int32(p6)):]))
			v2 = t7
			v3 = i32(255)
			v7 = i64(0)
			goto l1
		}
	l0:
		t8 := int64(load64(m.memory[int64(uint32(i32(0)))+1277248:]))
		v8 = t8
		v7 = int64(uint64(v8) >> 8)
		v3 = int32(v8)
		if v8&i64(255) != i64(255) {
			goto l2
		}
	}
l1:
	v4 = v5 + i64(2)
l2:
	store64(m.memory[int64(uint32(v1))+8:], uint64(v4))
	if v3&i32(255) == i32(255) {
		goto l3
	}
	store64(m.memory[uint32(v0):], uint64(v7<<8|int64(uint32(v3))&i64(255)))
	return
l3:
	m.memory[uint32(v0)] = byte(i32(255))
	store16(m.memory[int64(uint32(v0))+2:], uint16(v2))
}
func (m *Module) fn279(v0, v1, v2, v3 int32) {
	var v4 int32
	var v5 int64
	t0 := m.g0
	v4 = t0 - i32(32)
	m.g0 = v4
	store32(m.memory[int64(uint32(v4))+24:], uint32(i32(0)))
	store64(m.memory[int64(uint32(v4))+16:], uint64(i64(0)))
	m.fn264(v4+i32(16), v2, v3)
	t1 := int32(load32(m.memory[int64(uint32(v4))+24:]))
	t2 := v4
	v3 = t1
	store32(m.memory[int64(uint32(t2))+12:], uint32(v3))
	{
		{
			t3 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			if v3 != t3 {
				goto l0
			}
			store32(m.memory[uint32(v0):], uint32(i32(-2)))
			t4 := int64(load64(m.memory[uint32(v1):]))
			store64(m.memory[int64(uint32(v0))+4:], uint64(t4))
			goto l1
		}
	l0:
		t5 := v4
		v5 = int64(uint32(i32(14))) << 32
		store64(m.memory[int64(uint32(t5))+24:], uint64(v5|int64(uint32(v4+i32(12)))))
		store64(m.memory[int64(uint32(v4))+16:], uint64(v5|int64(uint32(v1+i32(8)))))
		m.fn17(v0, i32(1276420), v4+i32(16))
		t6 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		v0 = t6
		if v0 == 0 {
			goto l1
		}
		t7 := int32(load32(m.memory[uint32(v1):]))
		v3 = t7
		t8 := int32(load32(m.memory[uint32(v3+i32(-4)):]))
		v1 = t8
		v2 = v1 & i32(-8)
		t9 := v2
		v1 = v1 & i32(3)
		p10 := i32(8)
		if v1 != 0 {
			p10 = i32(4)
		}
		if uint32(t9) < uint32(p10+v0) {
			m.fn7(i32(1274404), i32(46), i32(1274452))
			panic("unreachable")
		}
		if v1 == 0 {
			goto l3
		}
		if uint32(v2) > uint32(v0+i32(39)) {
			m.fn7(i32(1274468), i32(46), i32(1274516))
			panic("unreachable")
		}
	l3:
		m.fn5(v3)
	}
l1:
	m.g0 = v4 + i32(32)
}
func (m *Module) fn280(v0, v1, v2 int32) {
	var v3 int32
	{
		if v2 != 0 {
			goto l0
		}
		v3 = i32(1)
		goto l1
	l0:
		t0 := m.fn11(v2)
		v3 = t0
		if v3 == 0 {
			m.fn16(i32(1), v2)
			panic("unreachable")
		}
		if v2 == 0 {
			goto l1
		}
		memory_copy(m.memory, uint32(v3), uint32(v1), uint32(v2))
	}
l1:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v2))
	store32(m.memory[uint32(v0):], uint32(v3))
}
func (m *Module) fn281(v0, v1 int32) {
	var v2, v3, v4, v5 int32
	{
		t0 := int32(load32(m.memory[uint32(v1):]))
		v2 = t0
		t1 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		t2 := v2
		v3 = t1
		if uint32(t2) > uint32(v3) {
			t4 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			v1 = t4
			{
				if v3 != 0 {
					t8 := m.fn26(v1, v2, i32(1), v3)
					v1 = t8
					if v1 != 0 {
						goto l1
					}
					m.fn16(i32(1), v3)
					panic("unreachable")
				}
				t5 := int32(load32(m.memory[uint32(v1+i32(-4)):]))
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
					goto l4
				}
				if uint32(v5) > uint32(v2+i32(39)) {
					m.fn7(i32(1274468), i32(46), i32(1274516))
					panic("unreachable")
				}
			l4:
				m.fn5(v1)
				v1 = i32(1)
				goto l1
			}
		}
		t3 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		v1 = t3
		goto l1
	}
l1:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
	store32(m.memory[uint32(v0):], uint32(v1))
}
func (m *Module) fn282(v0 int32) {
	var v1, v2, v3, v4 int32
	{
		t0 := int32(load32(m.memory[uint32(v0):]))
		v1 = t0
		v2 = v1 ^ i32(-0x80000000)
		p1 := i32(1)
		if uint32(v2) < uint32(i32(6)) {
			p1 = v2
		}
		switch p1 {
		default:
			return
		case 0:
			t2 := int32(m.memory[int64(uint32(v0))+4])
			if t2 != i32(3) {
				return
			}
			t3 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			v0 = t3
			t4 := int32(load32(m.memory[uint32(v0):]))
			v2 = t4
			{
				t5 := int32(load32(m.memory[uint32(v0+i32(4)):]))
				v1 = t5
				t6 := int32(load32(m.memory[uint32(v1):]))
				v3 = t6
				if v3 == 0 {
					goto l3
				}
				m.t0[uint(v3)].(func(int32))(v2)
			}
		l3:
			{
				t7 := int32(load32(m.memory[int64(uint32(v1))+4:]))
				v1 = t7
				if v1 == 0 {
					goto l4
				}
				t8 := int32(load32(m.memory[uint32(v2+i32(-4)):]))
				v3 = t8
				v4 = v3 & i32(-8)
				t9 := v4
				v3 = v3 & i32(3)
				p10 := i32(8)
				if v3 != 0 {
					p10 = i32(4)
				}
				if uint32(t9) < uint32(p10+v1) {
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
			t11 := int32(load32(m.memory[uint32(v0+i32(-4)):]))
			v2 = t11
			v1 = v2 & i32(-8)
			t12 := v1
			v2 = v2 & i32(3)
			p13 := i32(20)
			if v2 != 0 {
				p13 = i32(16)
			}
			if uint32(t12) < uint32(p13) {
				m.fn7(i32(1274404), i32(46), i32(1274452))
				panic("unreachable")
			}
			if v2 == 0 {
				goto l9
			}
			if uint32(v1) < uint32(i32(52)) {
				goto l9
			}
			m.fn7(i32(1274468), i32(46), i32(1274516))
			panic("unreachable")
		case 1:
			if uint32(v1+i32(-1)) > uint32(i32(-3)) {
				return
			}
			t14 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			v0 = t14
			t15 := int32(load32(m.memory[uint32(v0+i32(-4)):]))
			v2 = t15
			v3 = v2 & i32(-8)
			t16 := v3
			v2 = v2 & i32(3)
			p17 := i32(8)
			if v2 != 0 {
				p17 = i32(4)
			}
			if uint32(t16) < uint32(p17+v1) {
				m.fn7(i32(1274404), i32(46), i32(1274452))
				panic("unreachable")
			}
			if v2 == 0 {
				goto l9
			}
			if uint32(v3) <= uint32(v1+i32(39)) {
				goto l9
			}
			m.fn7(i32(1274468), i32(46), i32(1274516))
			panic("unreachable")
		}
	}
l9:
	m.fn5(v0)
}
func (m *Module) fn283(v0, v1 int32) {
	var v2, v3 int32
	var v4, v5 int64
	var v6 int32
	var v7, v8 int64
	v2 = i32(0)
	t0 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	v3 = t0
	v4 = int64(uint32(v3))
	{
		{
			t1 := int64(load64(m.memory[int64(uint32(v1))+8:]))
			t2 := v3
			v5 = t1
			p3 := i64(0xffffffff)
			if uint64(v5) < uint64(i64(0xffffffff)) {
				p3 = v5
			}
			v6 = t2 - int32(p3)
			p4 := v6
			if uint32(v6) > uint32(v3) {
				p4 = i32(0)
			}
			if uint32(p4) < uint32(i32(4)) {
				goto l0
			}
			t5 := int32(load32(m.memory[uint32(v1):]))
			p6 := v4
			if uint64(v5) < uint64(v4) {
				p6 = v5
			}
			t7 := int32(load32(m.memory[uint32(t5+int32(p6)):]))
			v2 = t7
			v3 = i32(255)
			v7 = i64(0)
			goto l1
		}
	l0:
		t8 := int64(load64(m.memory[int64(uint32(i32(0)))+1277248:]))
		v8 = t8
		v7 = int64(uint64(v8) >> 8)
		v3 = int32(v8)
		if v8&i64(255) != i64(255) {
			goto l2
		}
	}
l1:
	v4 = v5 + i64(4)
l2:
	store64(m.memory[int64(uint32(v1))+8:], uint64(v4))
	if v3&i32(255) == i32(255) {
		goto l3
	}
	store64(m.memory[uint32(v0):], uint64(v7<<8|int64(uint32(v3))&i64(255)))
	return
l3:
	m.memory[uint32(v0)] = byte(i32(255))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v2))
}
func (m *Module) fn284(v0, v1, v2 int32) {
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
	m.fn287(t2, t4, t3, v2, i32(1), i32(1))
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
func (m *Module) fn285(v0 int32) {
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
	m.fn208(t2, t4, t3, v2, i32(8), i32(176))
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
func (m *Module) fn286(v0, v1, v2 int32) int32 {
	var v3, v4, v5, v6, v7, v8, v9, v10, v11, v12 int32
	var v13 int64
	var v14, v15, v16 int32
	var v17 int64
	var v18 int32
	var v19 int64
	{
		{
			{
				t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
				v3 = t0
				v4 = v3 + i32(1)
				if v4 == 0 {
					m.fn28(i32(1271760), i32(57), i32(1271744))
					panic("unreachable")
				}
				{
					t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
					t2 := v4
					v5 = t1
					t3 := v5
					v6 = v5 + i32(1)
					v7 = int32(uint32(v6) >> 3)
					p4 := v7 * i32(7)
					if uint32(v5) < uint32(i32(8)) {
						p4 = t3
					}
					v8 = p4
					if uint32(t2) <= uint32(int32(uint32(v8)>>1)) {
						{
							if v6 == 0 {
								goto l6
							}
							t7 := int32(load32(m.memory[uint32(v0):]))
							v9 = t7
							v4 = i32(0)
							{
								{
									t8 := v7
									var p9 int32
									if v6&i32(7) != i32(0) {
										p9 = 1
									}
									v7 = t8 + p9
									if v7 == i32(1) {
										goto l7
									}
									v11 = v7 & i32(1)
									v12 = v7 & i32(0x3ffffffe)
									v4 = i32(0)
								l8:
									{
										v7 = v9 + v4
										t10 := int64(load64(m.memory[uint32(v7):]))
										t11 := v7
										v13 = t10
										store64(m.memory[uint32(t11):], uint64(int64(uint64(v13^i64(-1))>>7)&i64(72340172838076673)+(v13|i64(0x7f7f7f7f7f7f7f7f))))
										v7 = v7 + i32(8)
										t12 := int64(load64(m.memory[uint32(v7):]))
										t13 := v7
										v13 = t12
										store64(m.memory[uint32(t13):], uint64(int64(uint64(v13^i64(-1))>>7)&i64(72340172838076673)+(v13|i64(0x7f7f7f7f7f7f7f7f))))
										v4 = v4 + i32(16)
										v12 = v12 + i32(-2)
										if v12 != 0 {
											goto l8
										}
									}
									if v11 == 0 {
										goto l9
									}
								}
							l7:
								v4 = v9 + v4
								t14 := int64(load64(m.memory[uint32(v4):]))
								t15 := v4
								v13 = t14
								store64(m.memory[uint32(t15):], uint64(int64(uint64(v13^i64(-1))>>7)&i64(72340172838076673)+(v13|i64(0x7f7f7f7f7f7f7f7f))))
							}
						l9:
							{
								if uint32(v6) < uint32(i32(8)) {
									goto l10
								}
								t16 := int64(load64(m.memory[uint32(v9):]))
								store64(m.memory[uint32(v9+v6):], uint64(t16))
								goto l11
							}
						l10:
							if v6 == 0 {
								goto l11
							}
							memory_copy(m.memory, uint32(v9+i32(8)), uint32(v9), uint32(v6))
						l11:
							v7 = i32(0)
						l20:
							{
								t17 := v9
								v4 = v7
								v12 = t17 + v4
								t18 := int32(m.memory[uint32(v12)])
								if t18 != i32(128) {
									goto l12
								}
								{
									v14 = v9 - v4<<2 + i32(-4)
									t19 := int32(load32(m.memory[uint32(v14):]))
									v7 = t19
									if uint32(v7) >= uint32(v2) {
										goto l13
									}
									v15 = v9 + (v4^i32(-1))<<2
								l19:
									{
										t20 := int32(load32(m.memory[int64(uint32(v1+v7*i32(192)))+184:]))
										v16 = t20
										v7 = v16 & v5
										v11 = v7
										{
											t21 := int64(load64(m.memory[uint32(v9+v7):]))
											v13 = t21 & i64(-0x7f7f7f7f7f7f7f80)
											if v13 != i64(0) {
												goto l14
											}
											v6 = i32(8)
											v11 = v7
										l15:
											{
												v11 = v11 + v6
												v6 = v6 + i32(8)
												t22 := v9
												v11 = v11 & v5
												t23 := int64(load64(m.memory[uint32(t22+v11):]))
												v13 = t23 & i64(-0x7f7f7f7f7f7f7f80)
												if v13 == 0 {
													goto l15
												}
											}
										}
									l14:
										{
											t24 := v9
											v11 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v13))))>>3) + v11) & v5
											t25 := int32(int8(m.memory[uint32(t24+v11)]))
											if t25 < i32(0) {
												goto l16
											}
											t26 := int64(load64(m.memory[uint32(v9):]))
											v11 = int32(uint32(int64(bits.TrailingZeros64(uint64(t26&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
										}
									l16:
										if uint32((v11-v7^(v4-v7))&v5) < uint32(i32(8)) {
											t32 := v12
											v7 = int32(uint32(v16) >> 25)
											m.memory[uint32(t32)] = byte(v7)
											m.memory[uint32(v9+(v4+i32(-8))&v5+i32(8))] = byte(v7)
											goto l12
										}
										v7 = v9 + v11
										t27 := int32(m.memory[uint32(v7)])
										v6 = t27
										t28 := v7
										v16 = int32(uint32(v16) >> 25)
										m.memory[uint32(t28)] = byte(v16)
										m.memory[uint32(v9+(v11+i32(-8))&v5+i32(8))] = byte(v16)
										v7 = v9 - v11<<2 + i32(-4)
										if v6 == i32(255) {
											goto l18
										}
										t29 := int32(load32(m.memory[uint32(v15):]))
										v11 = t29
										t30 := int32(load32(m.memory[uint32(v7):]))
										store32(m.memory[uint32(v15):], uint32(t30))
										store32(m.memory[uint32(v7):], uint32(v11))
										t31 := int32(load32(m.memory[uint32(v14):]))
										v7 = t31
										if uint32(v7) < uint32(v2) {
											goto l19
										}
									}
								}
							l13:
								m.fn33(v7, v2, i32(1276328))
								panic("unreachable")
							l18:
								m.memory[uint32(v12)] = byte(i32(255))
								m.memory[uint32(v9+(v4+i32(-8))&v5+i32(8))] = byte(i32(255))
								t33 := int32(load32(m.memory[uint32(v15):]))
								store32(m.memory[uint32(v7):], uint32(t33))
							}
						l12:
							v7 = v4 + i32(1)
							if v4 != v5 {
								goto l20
							}
						}
					l6:
						store32(m.memory[int64(uint32(v0))+8:], uint32(v8-v3))
						goto l21
					}
					v9 = v8 + i32(1)
					p5 := v4
					if uint32(v9) > uint32(v4) {
						p5 = v9
					}
					v4 = p5
					if uint32(v4) < uint32(i32(15)) {
						goto l2
					}
					{
						if uint32(v4) > uint32(i32(0x1fffffff)) {
							m.fn28(i32(1271760), i32(57), i32(1271744))
							panic("unreachable")
						}
						t6 := int32(uint32(v4<<3) / uint32(i32(7)))
						v4 = i32_shr_u(i32(-1), int32(bits.LeadingZeros32(uint32(t6+i32(-1)))))
						if uint32(v4) > uint32(i32(0x3ffffffd)) {
							goto l4
						}
						v10 = v4 + i32(1)
						goto l5
					}
				}
			}
		l2:
			p34 := v4&i32(8) + i32(8)
			if uint32(v4) < uint32(i32(4)) {
				p34 = i32(4)
			}
			v10 = p34
		}
	l5:
		v4 = (v10<<2 + i32(7)) & i32(-8)
		t35 := v4
		v12 = v10 + i32(8)
		v9 = t35 + v12
		if uint32(v9) < uint32(v4) {
			goto l4
		}
		if uint32(v9) > uint32(i32(0x7ffffff8)) {
			goto l4
		}
		{
			t36 := m.fn11(v9)
			v7 = t36
			if v7 != 0 {
				v7 = v7 + v4
				if v12 == 0 {
					goto l23
				}
				memory_fill(m.memory, uint32(v7), i32(255), uint32(v12))
			l23:
				v11 = v10 + i32(-1)
				t37 := int32(load32(m.memory[uint32(v0):]))
				v8 = t37
				{
					if v3 == 0 {
						goto l24
					}
					t38 := int64(load64(m.memory[uint32(v8):]))
					v13 = (t38 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
					v9 = v8
					v4 = i32(0)
					v16 = v3
				l31:
					{
						if v13 != i64(0) {
							goto l25
						}
					l26:
						{
							v4 = v4 + i32(8)
							v9 = v9 + i32(8)
							t39 := int64(load64(m.memory[uint32(v9):]))
							v13 = t39 & i64(-0x7f7f7f7f7f7f7f80)
							if v13 == i64(-0x7f7f7f7f7f7f7f80) {
								goto l26
							}
						}
						v13 = v13 ^ i64(-0x7f7f7f7f7f7f7f80)
					l25:
						v15 = v8 - (int32(uint32(int64(bits.TrailingZeros64(uint64(v13))))>>3)+v4)<<2 + i32(-4)
						t40 := int32(load32(m.memory[uint32(v15):]))
						v12 = t40
						if uint32(v12) >= uint32(v2) {
							m.fn33(v12, v2, i32(1276328))
							panic("unreachable")
						}
						{
							t41 := int32(load32(m.memory[int64(uint32(v1+v12*i32(192)))+184:]))
							t42 := v7
							v14 = t41
							v12 = v14 & v11
							t43 := int64(load64(m.memory[uint32(t42+v12):]))
							v17 = t43 & i64(-0x7f7f7f7f7f7f7f80)
							if v17 != i64(0) {
								goto l28
							}
							v18 = i32(8)
						l29:
							{
								v12 = v12 + v18
								v18 = v18 + i32(8)
								t44 := v7
								v12 = v12 & v11
								t45 := int64(load64(m.memory[uint32(t44+v12):]))
								v17 = t45 & i64(-0x7f7f7f7f7f7f7f80)
								if v17 == 0 {
									goto l29
								}
							}
						}
					l28:
						v19 = v13 + i64(-1)
						{
							t46 := v7
							v12 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v17))))>>3) + v12) & v11
							t47 := int32(int8(m.memory[uint32(t46+v12)]))
							if t47 < i32(0) {
								goto l30
							}
							t48 := int64(load64(m.memory[uint32(v7):]))
							v12 = int32(uint32(int64(bits.TrailingZeros64(uint64(t48&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
						}
					l30:
						v13 = v19 & v13
						t49 := v7 + v12
						v14 = int32(uint32(v14) >> 25)
						m.memory[uint32(t49)] = byte(v14)
						m.memory[uint32(v7+(v12+i32(-8))&v11+i32(8))] = byte(v14)
						t50 := int32(load32(m.memory[uint32(v15):]))
						store32(m.memory[uint32(v7-v12<<2+i32(-4)):], uint32(t50))
						v16 = v16 + i32(-1)
						if v16 != 0 {
							goto l31
						}
					}
				}
			l24:
				store32(m.memory[int64(uint32(v0))+4:], uint32(v11))
				store32(m.memory[uint32(v0):], uint32(v7))
				t52 := v0
				p51 := int32(uint32(v10)>>3) * i32(7)
				if uint32(v10) < uint32(i32(9)) {
					p51 = v11
				}
				store32(m.memory[int64(uint32(t52))+8:], uint32(p51-v3))
				if v5 == 0 {
					goto l21
				}
				t53 := v8
				v4 = (v6<<2 + i32(7)) & i32(-8)
				v7 = t53 - v4
				t54 := int32(load32(m.memory[uint32(v7+i32(-4)):]))
				v9 = t54
				v12 = v9 & i32(-8)
				t55 := v12
				v9 = v9 & i32(3)
				p56 := i32(8)
				if v9 != 0 {
					p56 = i32(4)
				}
				v4 = v5 + v4 + i32(9)
				if uint32(t55) < uint32(p56+v4) {
					m.fn7(i32(1274404), i32(46), i32(1274452))
					panic("unreachable")
				}
				if v9 == 0 {
					goto l33
				}
				if uint32(v12) > uint32(v4+i32(39)) {
					m.fn7(i32(1274468), i32(46), i32(1274516))
					panic("unreachable")
				}
			l33:
				m.fn5(v7)
				return i32(-1)
			}
			m.fn23(i32(8), v9)
			panic("unreachable")
		}
	}
l21:
	return i32(-1)
l4:
	m.fn28(i32(1271760), i32(57), i32(1271744))
	panic("unreachable")
}
func (m *Module) fn287(v0, v1, v2, v3, v4, v5 int32) {
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
			t0 := m.fn26(v2, v5*v1, v4, v3)
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
func (m *Module) fn288(v0 int32) {
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
	m.fn287(t2, t4, t3, v2, i32(8), i32(192))
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
func (m *Module) fn289(v0, v1 int32) int32 {
	var v2, v3, v4 int32
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
				t11 := int32(m.memory[uint32(v0)])
				v3 = t11
				v0 = i32(3)
			l6:
				{
					t12 := int32(m.memory[uint32(v3&i32(15)+i32(1099352))])
					m.memory[uint32(v2+i32(9)+v0+i32(-2))] = byte(t12)
					v0 = v0 + i32(-1)
					v3 = int32(uint32(v3)>>4) & i32(15)
					if v3 != 0 {
						goto l6
					}
				}
				t13 := m.fn681(v1, i32(1), i32(1123086), i32(2), v2+i32(9)+v0+i32(-1), i32(3)-v0)
				v0 = t13
				goto l5
			}
			if v3&i32(0x4000000) != 0 {
				goto l1
			}
			v3 = i32(3)
			t3 := int32(m.memory[uint32(v0)])
			v0 = t3
			v4 = v0
			{
				if uint32(v0) < uint32(i32(10)) {
					goto l2
				}
				v3 = i32(1)
				t4 := int32(uint32(v0) / uint32(i32(100)))
				t5 := v2
				t6 := v0
				v4 = t4
				t7 := int32(load16(m.memory[int64(uint32((t6-v4*i32(100))&i32(255)<<1))+1100735:]))
				store16(m.memory[int64(uint32(t5))+12:], uint16(t7))
			}
		l2:
			{
				if v0 == 0 {
					goto l3
				}
				if v4 == 0 {
					goto l4
				}
			l3:
				t8 := v2 + i32(11)
				v3 = v3 + i32(-1)
				t9 := int32(m.memory[int64(uint32(v4<<1))+1100736])
				m.memory[uint32(t8+v3)] = byte(t9)
			}
		l4:
			t10 := m.fn681(v1, i32(1), i32(1), i32(0), v2+i32(11)+v3, i32(3)-v3)
			v0 = t10
			goto l5
		}
	l1:
		t14 := int32(m.memory[uint32(v0)])
		v3 = t14
		v0 = i32(3)
	l7:
		{
			t15 := int32(m.memory[uint32(v3&i32(15)+i32(1123088))])
			m.memory[uint32(v2+i32(14)+v0+i32(-2))] = byte(t15)
			v0 = v0 + i32(-1)
			v3 = int32(uint32(v3)>>4) & i32(15)
			if v3 != 0 {
				goto l7
			}
		}
		t16 := m.fn681(v1, i32(1), i32(1123086), i32(2), v2+i32(14)+v0+i32(-1), i32(3)-v0)
		v0 = t16
	}
l5:
	m.g0 = v2 + i32(16)
	return v0
}
func (m *Module) fn290(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v1):]))
	t1 := int32(m.memory[uint32(v0)])
	v0 = t1 << 2
	t2 := int32(load32(m.memory[int64(uint32(v0))+1290772:]))
	t3 := int32(load32(m.memory[int64(uint32(v0))+1290756:]))
	t4 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	t5 := int32(load32(m.memory[int64(uint32(t4))+12:]))
	t6 := m.t0[uint(t5)].(func(int32, int32, int32) int32)(t0, t2, t3)
	return t6
}
func (m *Module) fn291(v0, v1, v2, v3 int32) int32 {
	var v4, v5, v6, v7, v8, v9 int32
	{
		{
			if uint32(v1) < uint32(i32(4)) {
				goto l0
			}
			t0 := int32(load32(m.memory[uint32(v0):]))
			if t0&i32(-2139062144) != 0 {
				goto l1
			}
			{
				v4 = (v0 + i32(3)) & i32(-4)
				p1 := v4 - v0
				if v4 == v0 {
					p1 = i32(4)
				}
				v4 = p1
				t2 := v4
				v5 = v1 + i32(-4)
				if uint32(t2) >= uint32(v5) {
					goto l2
				}
			l3:
				{
					t3 := int32(load32(m.memory[uint32(v0+v4):]))
					if t3&i32(-2139062144) != 0 {
						goto l1
					}
					v4 = v4 + i32(4)
					if uint32(v4) < uint32(v5) {
						goto l3
					}
				}
			}
		l2:
			t4 := int32(load32(m.memory[uint32(v0+v5):]))
			if t4&i32(-2139062144) != 0 {
				goto l1
			}
			goto l4
		}
	l0:
		if v1 == 0 {
			goto l4
		}
		t5 := v0
		v4 = v1 + i32(-1)
		t6 := int32(int8(m.memory[uint32(t5+v4)]))
		if t6 < i32(0) {
			goto l1
		}
		if v4 == 0 {
			goto l4
		}
		t7 := v0
		v4 = v1 + i32(-2)
		t8 := int32(int8(m.memory[uint32(t7+v4)]))
		if t8 < i32(0) {
			goto l1
		}
		if v4 == 0 {
			goto l4
		}
		t9 := v0
		v4 = v1 + i32(-3)
		t10 := int32(int8(m.memory[uint32(t9+v4)]))
		if t10 < i32(0) {
			goto l1
		}
		if v4 != 0 {
			goto l1
		}
	}
l4:
	{
		{
			if uint32(v3) < uint32(i32(4)) {
				goto l5
			}
			t11 := int32(load32(m.memory[uint32(v2):]))
			if t11&i32(-2139062144) != 0 {
				goto l1
			}
			{
				v4 = (v2 + i32(3)) & i32(-4)
				p12 := v4 - v2
				if v4 == v2 {
					p12 = i32(4)
				}
				v4 = p12
				t13 := v4
				v5 = v3 + i32(-4)
				if uint32(t13) >= uint32(v5) {
					goto l6
				}
			l7:
				{
					t14 := int32(load32(m.memory[uint32(v2+v4):]))
					if t14&i32(-2139062144) != 0 {
						goto l1
					}
					v4 = v4 + i32(4)
					if uint32(v4) < uint32(v5) {
						goto l7
					}
				}
			}
		l6:
			t15 := int32(load32(m.memory[uint32(v2+v5):]))
			if t15&i32(-2139062144) != 0 {
				goto l1
			}
			goto l8
		}
	l5:
		if v3 == 0 {
			goto l8
		}
		t16 := v2
		v4 = v3 + i32(-1)
		t17 := int32(int8(m.memory[uint32(t16+v4)]))
		if t17 < i32(0) {
			goto l1
		}
		if v4 == 0 {
			goto l8
		}
		t18 := v2
		v4 = v3 + i32(-2)
		t19 := int32(int8(m.memory[uint32(t18+v4)]))
		if t19 < i32(0) {
			goto l1
		}
		if v4 == 0 {
			goto l8
		}
		t20 := v2
		v4 = v3 + i32(-3)
		t21 := int32(int8(m.memory[uint32(t20+v4)]))
		if t21 < i32(0) {
			goto l1
		}
		if v4 != 0 {
			goto l1
		}
	}
l8:
	{
		if v1 == v3 {
			v6 = i32(0)
			{
			l11:
				{
					if v1 == 0 {
						goto l10
					}
					t24 := int32(m.memory[uint32(v2)])
					v4 = t24
					t25 := int32(m.memory[uint32(v0)])
					v5 = t25
					v1 = v1 + i32(-1)
					v2 = v2 + i32(1)
					v0 = v0 + i32(1)
					t27 := v5
					p26 := i32(0)
					if uint32((v5+i32(-97))&i32(255)) < uint32(i32(26)) {
						p26 = i32(32)
					}
					v5 = t27 ^ p26
					t29 := v5 & i32(255)
					t30 := v4
					p28 := i32(0)
					if uint32((v4+i32(-97))&i32(255)) < uint32(i32(26)) {
						p28 = i32(32)
					}
					v4 = t30 ^ p28
					if t29 == v4&i32(255) {
						goto l11
					}
				}
				v1 = v5 & i32(255)
				t31 := v1
				v4 = v4 & i32(255)
				var p32 int32
				if uint32(t31) > uint32(v4) {
					p32 = 1
				}
				var p33 int32
				if uint32(v1) < uint32(v4) {
					p33 = 1
				}
				v6 = p32 - p33
			}
		l10:
			return v6
		}
		var p22 int32
		if uint32(v1) > uint32(v3) {
			p22 = 1
		}
		var p23 int32
		if uint32(v1) < uint32(v3) {
			p23 = 1
		}
		return p22 - p23
	}
l1:
	v7 = v0 + v1
	v1 = i32(0)
	v5 = v0
	v4 = i32(0)
l13:
	if v4&i32(0xffff) == 0 {
		if v5 == v7 {
			v8 = v2 + v3
			v4 = i32(0)
			v6 = v2
			v5 = i32(0)
		l21:
			if v5&i32(0xffff) == 0 {
				if v6 == v8 {
					if v1 == v4 {
					l39:
						{
							if v0 == v7 {
								if v2 != v8 {
									t57 := int32(int8(m.memory[uint32(v2)]))
									v1 = t57
									if v1 <= i32(-1) {
										t59 := int32(m.memory[int64(uint32(v2))+1])
										v4 = t59 & i32(63)
										v5 = v1 & i32(31)
										if uint32(v1) > uint32(i32(-33)) {
											t61 := int32(m.memory[int64(uint32(v2))+2])
											v4 = v4<<6 | t61&i32(63)
											if uint32(v1) >= uint32(i32(-16)) {
												t63 := int32(m.memory[int64(uint32(v2))+3])
												_ = m.fn842(v4<<6 | t63&i32(63) | v5<<18&i32(0x1c0000))
												return i32(255)
											}
											_ = m.fn842(v4 | v5<<12)
											return i32(255)
										}
										_ = m.fn842(v5<<6 | v4)
										return i32(255)
									}
									_ = m.fn842(v1 & i32(255))
									return i32(255)
								}
								return i32(0)
							}
							{
								{
									t44 := int32(int8(m.memory[uint32(v0)]))
									v1 = t44
									if v1 <= i32(-1) {
										goto l30
									}
									v0 = v0 + i32(1)
									v1 = v1 & i32(255)
									goto l31
								}
							l30:
								t45 := int32(m.memory[int64(uint32(v0))+1])
								v4 = t45 & i32(63)
								v5 = v1 & i32(31)
								if uint32(v1) > uint32(i32(-33)) {
									goto l32
								}
								v1 = v5<<6 | v4
								v0 = v0 + i32(2)
								goto l31
							l32:
								t46 := int32(m.memory[int64(uint32(v0))+2])
								v4 = v4<<6 | t46&i32(63)
								if uint32(v1) >= uint32(i32(-16)) {
									goto l33
								}
								v1 = v4 | v5<<12
								v0 = v0 + i32(3)
								goto l31
							l33:
								t47 := int32(m.memory[int64(uint32(v0))+3])
								v1 = v4<<6 | t47&i32(63) | v5<<18&i32(0x1c0000)
								v0 = v0 + i32(4)
							}
						l31:
							t48 := m.fn842(v1)
							v4 = t48
							if v2 != v8 {
								goto l34
							}
							return i32(1)
						l34:
							{
								{
									t49 := int32(int8(m.memory[uint32(v2)]))
									v1 = t49
									if v1 <= i32(-1) {
										goto l35
									}
									v2 = v2 + i32(1)
									v1 = v1 & i32(255)
									goto l36
								}
							l35:
								t50 := int32(m.memory[int64(uint32(v2))+1])
								v5 = t50 & i32(63)
								v6 = v1 & i32(31)
								if uint32(v1) > uint32(i32(-33)) {
									goto l37
								}
								v1 = v6<<6 | v5
								v2 = v2 + i32(2)
								goto l36
							l37:
								t51 := int32(m.memory[int64(uint32(v2))+2])
								v5 = v5<<6 | t51&i32(63)
								if uint32(v1) >= uint32(i32(-16)) {
									goto l38
								}
								v1 = v5 | v6<<12
								v2 = v2 + i32(3)
								goto l36
							l38:
								t52 := int32(m.memory[int64(uint32(v2))+3])
								v1 = v5<<6 | t52&i32(63) | v6<<18&i32(0x1c0000)
								v2 = v2 + i32(4)
							}
						l36:
							t53 := m.fn842(v1)
							t54 := v4
							v1 = t53
							if t54 == v1 {
								goto l39
							}
						}
						var p55 int32
						if uint32(v4) > uint32(v1) {
							p55 = 1
						}
						var p56 int32
						if uint32(v4) < uint32(v1) {
							p56 = 1
						}
						return p55 - p56
					}
					var p42 int32
					if uint32(v1) > uint32(v4) {
						p42 = 1
					}
					var p43 int32
					if uint32(v1) < uint32(v4) {
						p43 = 1
					}
					return p42 - p43
				}
				{
					t38 := int32(int8(m.memory[uint32(v6)]))
					v5 = t38
					if v5 <= i32(-1) {
						if uint32(v5) >= uint32(i32(-32)) {
							v3 = v5 & i32(31)
							t39 := int32(m.memory[int64(uint32(v6))+1])
							t40 := int32(m.memory[int64(uint32(v6))+2])
							v9 = t39&i32(63)<<6 | t40&i32(63)
							{
								if uint32(v5) >= uint32(i32(-16)) {
									goto l25
								}
								v5 = v9 | v3<<12
								v6 = v6 + i32(3)
								goto l26
							l25:
								t41 := int32(m.memory[int64(uint32(v6))+3])
								v5 = v9<<6 | t41&i32(63) | v3<<18&i32(0x1c0000)
								v6 = v6 + i32(4)
							}
						l26:
							if uint32(v5) >= uint32(i32(65536)) {
								v5 = v5&i32(1023) | i32(-9216)
								v4 = v4 + i32(1)
								goto l21
							}
							v5 = i32(0)
							v4 = v4 + i32(1)
							goto l21
						}
						v6 = v6 + i32(2)
						v5 = i32(0)
						v4 = v4 + i32(1)
						goto l21
					}
					v6 = v6 + i32(1)
					v5 = i32(0)
					v4 = v4 + i32(1)
					goto l21
				}
			}
			v5 = i32(0)
			v4 = v4 + i32(1)
			goto l21
		}
		{
			t34 := int32(int8(m.memory[uint32(v5)]))
			v4 = t34
			if v4 <= i32(-1) {
				if uint32(v4) >= uint32(i32(-32)) {
					v6 = v4 & i32(31)
					t35 := int32(m.memory[int64(uint32(v5))+1])
					t36 := int32(m.memory[int64(uint32(v5))+2])
					v8 = t35&i32(63)<<6 | t36&i32(63)
					{
						if uint32(v4) >= uint32(i32(-16)) {
							goto l17
						}
						v4 = v8 | v6<<12
						v5 = v5 + i32(3)
						goto l18
					l17:
						t37 := int32(m.memory[int64(uint32(v5))+3])
						v4 = v8<<6 | t37&i32(63) | v6<<18&i32(0x1c0000)
						v5 = v5 + i32(4)
					}
				l18:
					if uint32(v4) >= uint32(i32(65536)) {
						v4 = v4&i32(1023) | i32(-9216)
						v1 = v1 + i32(1)
						goto l13
					}
					v4 = i32(0)
					v1 = v1 + i32(1)
					goto l13
				}
				v5 = v5 + i32(2)
				v4 = i32(0)
				v1 = v1 + i32(1)
				goto l13
			}
			v5 = v5 + i32(1)
			v4 = i32(0)
			v1 = v1 + i32(1)
			goto l13
		}
	}
	v4 = i32(0)
	v1 = v1 + i32(1)
	goto l13
}
func (m *Module) fn292(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t1 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	t2 := int32(load32(m.memory[uint32(v1):]))
	t3 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	t4 := m.fn53(t0, t1, t2, t3)
	return t4
}
func (m *Module) fn293(v0 int32) {
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
	m.fn208(t2, t4, t3, v2, i32(4), i32(8))
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
func (m *Module) fn294(v0, v1, v2, v3 int32) {
	var v4 int32
	var v5 int64
	var v6, v7 int32
	var v8, v9 int64
	var v10, v11, v12 int32
	var v13 int64
	var v14 int32
	var v15, v16 int64
	var v17 int32
	var v18 int64
	var v19, v20, v21, v22, v23 int32
	t0 := m.g0
	v4 = t0 - i32(48)
	m.g0 = v4
	v5 = int64(uint32(i32(3))) << 32
	t1 := int32(load32(m.memory[int64(uint32(v1))+20:]))
	t2 := v5
	v6 = t1
	v7 = v6 + i32(16)
	v8 = t2 | int64(uint32(v7))
	v9 = v5 | int64(uint32(v4+i32(16)))
	v10 = v4 + i32(4) + i32(4)
	t3 := int32(load32(m.memory[int64(uint32(v1))+12:]))
	v11 = t3
	t4 := int64(load64(m.memory[uint32(v1):]))
	v5 = t4
	t5 := int32(load32(m.memory[int64(uint32(v1))+16:]))
	v12 = t5
	v13 = int64(uint32(v12))
l24:
	{
		{
			{
				t6 := int32(m.memory[int64(uint32(v6))+20])
				t7 := v13
				v14 = t6
				p8 := i64(9)
				if v14 != 0 {
					p8 = i64(12)
				}
				v15 = p8
				v16 = i64_shl(t7, v15)
				if v16 == v5 {
					goto l0
				}
				t9 := v12
				v17 = int32(i64_shr_u(v5, v15))
				if uint32(t9) <= uint32(v17) {
					m.fn33(v17, v12, i32(1079960))
					panic("unreachable")
				}
				t10 := int32(load32(m.memory[uint32(v11+v17<<2):]))
				t11 := v4
				v17 = t10
				store32(m.memory[int64(uint32(t11))+16:], uint32(v17))
				{
					{
						t12 := int32(load32(m.memory[uint32(v7):]))
						if uint32(v17) >= uint32(t12) {
							goto l2
						}
						t14 := v4
						p13 := i32(512)
						if v14 != 0 {
							p13 = i32(4096)
						}
						v14 = p13
						store32(m.memory[int64(uint32(t14))+8:], uint32(v14))
						t15 := v4
						v18 = (int64(uint32(v14)) + i64(-1)) & v5
						store32(m.memory[int64(uint32(t15))+12:], uint32(v18))
						t16 := v6
						v15 = i64_shl(int64(uint32(v17+i32(1))), v15) + v18
						store64(m.memory[int64(uint32(t16))+8:], uint64(v15))
						t17 := v3
						v16 = v16 - v5
						t18 := v16
						v18 = int64(uint32(v3))
						p19 := v18
						if uint64(v16) < uint64(v18) {
							p19 = t18
						}
						v14 = int32(p19)
						if uint32(t17) < uint32(v14) {
							m.fn121(i32(0), v14, v3, i32(1079976))
							panic("unreachable")
						}
						v17 = i32(0)
						{
							t20 := int64(load64(m.memory[int64(uint32(v4))+8:]))
							v16 = t20
							v19 = int32(v16)
							t21 := v19
							v20 = int32(int64(uint64(v16) >> 32))
							if t21 == v20 {
								goto l4
							}
							t22 := int32(load32(m.memory[uint32(v6):]))
							t23 := int32(load32(m.memory[int64(uint32(v6))+4:]))
							t24 := v15
							v17 = t23
							v16 = int64(uint32(v17))
							p25 := v16
							if uint64(v15) < uint64(v16) {
								p25 = t24
							}
							v21 = t22 + int32(p25)
							{
								t27 := v17
								p26 := i64(0xffffffff)
								if uint64(v15) < uint64(i64(0xffffffff)) {
									p26 = v15
								}
								v22 = t27 - int32(p26)
								p28 := v22
								if uint32(v22) > uint32(v17) {
									p28 = i32(0)
								}
								v17 = p28
								t29 := v17
								v19 = v19 - v20
								p30 := v14
								if uint32(v19) < uint32(v14) {
									p30 = v19
								}
								v14 = p30
								p31 := v14
								if uint32(v17) < uint32(v14) {
									p31 = t29
								}
								v17 = p31
								if v17 != i32(1) {
									goto l5
								}
								t32 := int32(m.memory[uint32(v21)])
								m.memory[uint32(v2)] = byte(t32)
								goto l6
							}
						l5:
							if v17 == 0 {
								goto l6
							}
							memory_copy(m.memory, uint32(v2), uint32(v21), uint32(v17))
						l6:
							store64(m.memory[int64(uint32(v6))+8:], uint64(v15+int64(uint32(v17))))
						}
					l4:
						t33 := v1
						v5 = v5 + int64(uint32(v17))
						store64(m.memory[uint32(t33):], uint64(v5))
						v23 = v23 | i32(255)
						goto l7
					}
				l2:
					store64(m.memory[int64(uint32(v4))+40:], uint64(v8))
					store64(m.memory[int64(uint32(v4))+32:], uint64(v9))
					m.fn17(v4+i32(20), i32(1049021), v4+i32(32))
					m.fn163(v10, i32(21), v4+i32(20))
					t34 := int32(load32(m.memory[int64(uint32(v4))+8:]))
					v23 = t34
					t35 := int32(load32(m.memory[int64(uint32(v4))+12:]))
					v17 = t35
				}
			l7:
				switch v23 & i32(255) {
				case 0:
					goto l8
				case 2:
					t37 := int32(m.memory[int64(uint32(v17))+8])
					if t37 == i32(35) {
						goto l13
					}
					goto l8
				case 3:
					goto l11
				case 1:
					if v23&i32(0xff00) != i32(8960) {
						goto l8
					}
					goto l13
				default:
					if v17 == 0 {
						goto l0
					}
					if uint32(v3) >= uint32(v17) {
						v2 = v2 + v17
						v3 = v3 - v17
						goto l13
					}
					m.fn121(v17, v3, v3, i32(1069360))
					panic("unreachable")
				}
			}
		l0:
			t36 := int64(load64(m.memory[int64(uint32(i32(0)))+1277248:]))
			store64(m.memory[uint32(v0):], uint64(t36))
			goto l15
		}
	l11:
		t38 := int32(m.memory[int64(uint32(v17))+8])
		if t38 != i32(35) {
			goto l8
		}
		t39 := int32(load32(m.memory[uint32(v17):]))
		v14 = t39
		{
			t40 := int32(load32(m.memory[uint32(v17+i32(4)):]))
			v19 = t40
			t41 := int32(load32(m.memory[uint32(v19):]))
			v20 = t41
			if v20 == 0 {
				goto l16
			}
			m.t0[uint(v20)].(func(int32))(v14)
		}
	l16:
		{
			t42 := int32(load32(m.memory[int64(uint32(v19))+4:]))
			v19 = t42
			if v19 == 0 {
				goto l17
			}
			t43 := int32(load32(m.memory[uint32(v14+i32(-4)):]))
			v20 = t43
			v22 = v20 & i32(-8)
			t44 := v22
			v20 = v20 & i32(3)
			p45 := i32(8)
			if v20 != 0 {
				p45 = i32(4)
			}
			if uint32(t44) < uint32(p45+v19) {
				m.fn7(i32(1274404), i32(46), i32(1274452))
				panic("unreachable")
			}
			if v20 == 0 {
				goto l19
			}
			if uint32(v22) > uint32(v19+i32(39)) {
				m.fn7(i32(1274468), i32(46), i32(1274516))
				panic("unreachable")
			}
		l19:
			m.fn5(v14)
		}
	l17:
		t46 := int32(load32(m.memory[uint32(v17+i32(-4)):]))
		v14 = t46
		v19 = v14 & i32(-8)
		t47 := v19
		v14 = v14 & i32(3)
		p48 := i32(20)
		if v14 != 0 {
			p48 = i32(16)
		}
		if uint32(t47) < uint32(p48) {
			m.fn7(i32(1274404), i32(46), i32(1274452))
			panic("unreachable")
		}
		if v14 == 0 {
			goto l22
		}
		if uint32(v19) >= uint32(i32(52)) {
			m.fn7(i32(1274468), i32(46), i32(1274516))
			panic("unreachable")
		}
	l22:
		m.fn5(v17)
	}
l13:
	if v3 != 0 {
		goto l24
	}
	m.memory[uint32(v0)] = byte(i32(255))
	goto l15
l8:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v17))
	store32(m.memory[uint32(v0):], uint32(v23))
l15:
	m.g0 = v4 + i32(48)
}
func (m *Module) fn295(v0 int32) {
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
	m.fn1908(t2, t4, t3, v2, i32(2), i32(2))
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
func (m *Module) fn296(v0, v1 int32) int32 {
	var v2, v3, v4 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	v3 = i32(3)
	t1 := int32(m.memory[uint32(v0)])
	v0 = t1
	v4 = v0
	{
		if uint32(v0) < uint32(i32(10)) {
			goto l0
		}
		v3 = i32(1)
		t2 := int32(uint32(v0) / uint32(i32(100)))
		t3 := v2
		t4 := v0
		v4 = t2
		t5 := int32(load16(m.memory[int64(uint32((t4-v4*i32(100))&i32(255)<<1))+1100735:]))
		store16(m.memory[int64(uint32(t3))+14:], uint16(t5))
	}
l0:
	{
		if v0 == 0 {
			goto l1
		}
		if v4 == 0 {
			goto l2
		}
	l1:
		t6 := v2 + i32(13)
		v3 = v3 + i32(-1)
		t7 := int32(m.memory[int64(uint32(v4<<1))+1100736])
		m.memory[uint32(t6+v3)] = byte(t7)
	}
l2:
	t8 := m.fn681(v1, i32(1), i32(1), i32(0), v2+i32(13)+v3, i32(3)-v3)
	v3 = t8
	m.g0 = v2 + i32(16)
	return v3
}
func (m *Module) fn297(v0, v1, v2 int32) {
	var v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14, v15, v16, v17, v18, v19, v20, v21, v22, v23, v24, v25, v26, v27 int32
	var v28 int64
	t0 := m.g0
	v3 = t0 - i32(112)
	m.g0 = v3
	{
		{
			{
				{
					if v2 != 0 {
						goto l0
					}
					v4 = i32(2)
					v5 = i32(0)
					v6 = i32(0)
					goto l1
				l0:
					v7 = v1 + v2
					{
						{
							t1 := int32(int8(m.memory[uint32(v1)]))
							v8 = t1
							if v8 <= i32(-1) {
								goto l2
							}
							v9 = v1 + i32(1)
							v8 = v8 & i32(255)
							v10 = i32(0)
							goto l3
						}
					l2:
						t2 := int32(m.memory[int64(uint32(v1))+1])
						v10 = t2 & i32(63)
						v11 = v8 & i32(31)
						if uint32(v8) > uint32(i32(-33)) {
							goto l4
						}
						v8 = v11<<6 | v10
						v9 = v1 + i32(2)
						v10 = i32(0)
						goto l3
					l4:
						t3 := int32(m.memory[int64(uint32(v1))+2])
						v10 = v10<<6 | t3&i32(63)
						{
							if uint32(v8) >= uint32(i32(-16)) {
								goto l5
							}
							v8 = v10 | v11<<12
							v9 = v1 + i32(3)
							goto l6
						l5:
							t4 := int32(m.memory[int64(uint32(v1))+3])
							v8 = v10<<6 | t4&i32(63) | v11<<18&i32(0x1c0000)
							v9 = v1 + i32(4)
						}
					l6:
						if uint32(v8) > uint32(i32(0xffff)) {
							goto l7
						}
						v10 = i32(0)
						goto l3
					l7:
						v10 = v8&i32(1023) | i32(-9216)
						v8 = int32(uint32(v8+i32(0xff0000))>>10) | i32(-10240)
					}
				l3:
					v11 = v7 - v9
					t5 := int32(uint32(v11) / uint32(i32(3)))
					v12 = t5
					t6 := v12
					var p7 int32
					if v10 != i32(0) {
						p7 = 1
					}
					t8 := t6 + p7
					var p9 int32
					if v11-v12*i32(3) != i32(0) {
						p9 = 1
					}
					v11 = t8 + p9
					p10 := i32(3)
					if uint32(v11) > uint32(i32(3)) {
						p10 = v11
					}
					v11 = p10
					p11 := i32(31)
					if uint32(v11) < uint32(i32(31)) {
						p11 = v11
					}
					v11 = p11 + i32(1)
					v12 = v11 << 1
					t12 := m.fn11(v12)
					v5 = t12
					if v5 == 0 {
						m.fn16(i32(2), v12)
						panic("unreachable")
					}
					store16(m.memory[uint32(v5):], uint16(v8))
					store32(m.memory[int64(uint32(v3))+56:], uint32(i32(1)))
					store32(m.memory[int64(uint32(v3))+52:], uint32(v5))
					store32(m.memory[int64(uint32(v3))+48:], uint32(v11))
					v12 = i32(30)
					v11 = i32(2)
					v8 = i32(2)
					{
					l17:
						v6 = v8 + i32(-1)
						v13 = i32(0)
						{
							if v10&i32(0xffff) != 0 {
								goto l9
							}
							if v9 == v7 {
								goto l10
							}
							{
								t13 := int32(int8(m.memory[uint32(v9)]))
								v10 = t13
								if v10 <= i32(-1) {
									goto l11
								}
								v9 = v9 + i32(1)
								v10 = v10 & i32(255)
								goto l9
							}
						l11:
							t14 := int32(m.memory[int64(uint32(v9))+1])
							v4 = t14 & i32(63)
							v14 = v10 & i32(31)
							if uint32(v10) > uint32(i32(-33)) {
								goto l12
							}
							v10 = v14<<6 | v4
							v9 = v9 + i32(2)
							goto l9
						l12:
							t15 := int32(m.memory[int64(uint32(v9))+2])
							v4 = v4<<6 | t15&i32(63)
							{
								if uint32(v10) >= uint32(i32(-16)) {
									goto l13
								}
								v10 = v4 | v14<<12
								v9 = v9 + i32(3)
								goto l14
							l13:
								t16 := int32(m.memory[int64(uint32(v9))+3])
								v10 = v4<<6 | t16&i32(63) | v14<<18&i32(0x1c0000)
								v9 = v9 + i32(4)
							}
						l14:
							if uint32(v10) <= uint32(i32(0xffff)) {
								goto l9
							}
							v13 = v10&i32(1023) | i32(-9216)
							v10 = int32(uint32(v10+i32(0xff0000))>>10) | i32(-10240)
						}
					l9:
						{
							t17 := int32(load32(m.memory[int64(uint32(v3))+48:]))
							if v6 != t17 {
								goto l15
							}
							v5 = i32(1)
							{
								if v12 == 0 {
									goto l16
								}
								t18 := v12
								v5 = v7 - v9
								t19 := int32(uint32(v5) / uint32(i32(3)))
								v4 = t19
								t20 := v4
								var p21 int32
								if v13 != i32(0) {
									p21 = 1
								}
								t22 := t20 + p21
								var p23 int32
								if v5-v4*i32(3) != i32(0) {
									p23 = 1
								}
								v5 = t22 + p23
								p24 := v5
								if uint32(v12) < uint32(v5) {
									p24 = t18
								}
								v5 = p24 + i32(1)
							}
						l16:
							m.fn715(v3+i32(48), v6, v5, i32(2), i32(2))
							t25 := int32(load32(m.memory[int64(uint32(v3))+52:]))
							v5 = t25
						}
					l15:
						store16(m.memory[uint32(v5+v11):], uint16(v10))
						v11 = v11 + i32(2)
						store32(m.memory[int64(uint32(v3))+56:], uint32(v8))
						v8 = v8 + i32(1)
						v10 = v13
						v12 = v12 + i32(-1)
						if v12 != i32(-1) {
							goto l17
						}
						t26 := int32(load32(m.memory[int64(uint32(v3))+52:]))
						v4 = t26
						t27 := int32(load32(m.memory[int64(uint32(v3))+48:]))
						v5 = t27
						goto l18
					}
				l10:
					t28 := int32(load32(m.memory[int64(uint32(v3))+52:]))
					v4 = t28
					t29 := int32(load32(m.memory[int64(uint32(v3))+48:]))
					v5 = t29
					if uint32(v6) > uint32(i32(31)) {
						goto l18
					}
				}
			l1:
				v7 = v2 + i32(-8)
				t30 := v1
				v15 = (v1 + i32(3)) & i32(-4)
				v16 = t30 - v15
				v17 = v15 - v1
				v18 = v3 + i32(48) + i32(8)
				v19 = v3 + i32(32) | i32(3)
				v20 = v3 + i32(32) | i32(2)
				v21 = v3 + i32(32) | i32(1)
				var p31 int32
				if uint32(v2) > uint32(i32(7)) {
					p31 = 1
				}
				v14 = p31
				var p32 int32
				if v2 == i32(4) {
					p32 = 1
				}
				v22 = p32
				var p33 int32
				if v2 == i32(5) {
					p33 = 1
				}
				v23 = p33
				var p34 int32
				if v2 == i32(6) {
					p34 = 1
				}
				v24 = p34
				v9 = i32(0)
			l51:
				{
					t35 := int32(load32(m.memory[int64(uint32(v9))+1093340:]))
					t36 := v3
					v10 = t35
					store32(m.memory[int64(uint32(t36))+16:], uint32(v10))
					{
						if uint32(v10) < uint32(i32(128)) {
							if v14 != 0 {
								v8 = i32(0)
								v11 = v1
								v12 = v16
								if v15 == v1 {
									goto l45
								}
							l46:
								{
									t65 := int32(m.memory[uint32(v11)])
									if t65 == v10&i32(255) {
										goto l25
									}
									v11 = v11 + i32(1)
									v12 = v12 + i32(1)
									if v12 != 0 {
										goto l46
									}
								}
								v8 = v17
								if uint32(v17) > uint32(v7) {
									goto l47
								}
							l45:
								v11 = v10 * i32(16843009)
							l48:
								{
									v12 = v1 + v8
									t66 := int32(load32(m.memory[uint32(v12):]))
									v13 = t66 ^ v11
									t67 := int32(load32(m.memory[uint32(v12+i32(4)):]))
									t68 := i32(16843008) - v13 | v13
									v12 = t67 ^ v11
									if t68&(i32(16843008)-v12|v12)&i32(-2139062144) != i32(-2139062144) {
										goto l47
									}
									v8 = v8 + i32(8)
									if uint32(v8) <= uint32(v7) {
										goto l48
									}
								}
							l47:
								if v2 == v8 {
									goto l24
								}
								v11 = v2 - v8
								v8 = v1 + v8
							l49:
								{
									t69 := int32(m.memory[uint32(v8)])
									if t69 == v10&i32(255) {
										goto l25
									}
									v8 = v8 + i32(1)
									v11 = v11 + i32(-1)
									if v11 == 0 {
										goto l24
									}
									goto l49
								}
							}
							if v2 == 0 {
								goto l24
							}
							t58 := int32(m.memory[uint32(v1)])
							v8 = v10 & i32(255)
							if t58 == v8 {
								goto l25
							}
							if v2 == i32(1) {
								goto l24
							}
							t59 := int32(m.memory[int64(uint32(v1))+1])
							if t59 == v8 {
								goto l25
							}
							if v2 == i32(2) {
								goto l24
							}
							t60 := int32(m.memory[int64(uint32(v1))+2])
							if t60 == v8 {
								goto l25
							}
							if v2 == i32(3) {
								goto l24
							}
							t61 := int32(m.memory[int64(uint32(v1))+3])
							if t61 == v8 {
								goto l25
							}
							if v22 != 0 {
								goto l24
							}
							t62 := int32(m.memory[int64(uint32(v1))+4])
							if t62 == v8 {
								goto l25
							}
							if v23 != 0 {
								goto l24
							}
							t63 := int32(m.memory[int64(uint32(v1))+5])
							if t63 == v8 {
								goto l25
							}
							if v24 != 0 {
								goto l24
							}
							t64 := int32(m.memory[int64(uint32(v1))+6])
							if t64 != v8 {
								goto l24
							}
							goto l25
						}
						store32(m.memory[int64(uint32(v3))+32:], uint32(i32(0)))
						v8 = int32(uint32(v10) >> 6)
						if uint32(v10) > uint32(i32(2047)) {
							goto l20
						}
						m.memory[int64(uint32(v3))+32] = byte(v8 | i32(192))
						v8 = i32(2)
						v11 = v21
						goto l21
					l20:
						v11 = int32(uint32(v10) >> 12)
						v8 = v8&i32(63) | i32(-128)
						if uint32(v10) > uint32(i32(0xffff)) {
							goto l22
						}
						m.memory[int64(uint32(v3))+33] = byte(v8)
						m.memory[int64(uint32(v3))+32] = byte(v11 | i32(224))
						v8 = i32(3)
						v11 = v20
						goto l21
					l22:
						m.memory[int64(uint32(v3))+34] = byte(v8)
						m.memory[int64(uint32(v3))+33] = byte(v11&i32(63) | i32(-128))
						m.memory[int64(uint32(v3))+32] = byte(int32(uint32(v10)>>18) | i32(-16))
						v8 = i32(4)
						v11 = v19
					l21:
						m.memory[uint32(v11)] = byte(v10&i32(63) | i32(128))
						{
							if uint32(v8) < uint32(v2) {
								m.fn158(v3+i32(48), v1, v2, v3+i32(32), v8)
								{
									t38 := int32(load32(m.memory[int64(uint32(v3))+48:]))
									if t38 != 0 {
										t53 := int32(load32(m.memory[int64(uint32(v3))+108:]))
										v12 = t53
										t54 := int32(load32(m.memory[int64(uint32(v3))+104:]))
										v11 = t54
										t55 := int32(load32(m.memory[int64(uint32(v3))+100:]))
										v10 = t55
										t56 := int32(load32(m.memory[int64(uint32(v3))+96:]))
										v8 = t56
										t57 := int32(load32(m.memory[int64(uint32(v3))+84:]))
										if t57 == i32(-1) {
											goto l43
										}
										m.fn206(v3+i32(36), v18, v8, v10, v11, v12, i32(0))
										goto l42
									}
									v10 = i32(0)
									{
										t39 := int32(m.memory[int64(uint32(v3))+62])
										if t39 != 0 {
											goto l27
										}
										t40 := int32(m.memory[int64(uint32(v3))+60])
										v13 = t40
										t41 := int32(load32(m.memory[int64(uint32(v3))+100:]))
										v11 = t41
										t42 := int32(load32(m.memory[int64(uint32(v3))+96:]))
										v12 = t42
										{
											t43 := int32(load32(m.memory[int64(uint32(v3))+52:]))
											v8 = t43
											if v8 == 0 {
												goto l28
											}
											if uint32(v8) < uint32(v11) {
												goto l29
											}
											if v8 == v11 {
												goto l28
											}
											goto l30
										l29:
											t44 := int32(int8(m.memory[uint32(v12+v8)]))
											if t44 < i32(-64) {
												goto l30
											}
										}
									l28:
										{
											if v8 == v11 {
												goto l31
											}
											{
												v25 = v12 + v8
												t45 := int32(int8(m.memory[uint32(v25)]))
												v10 = t45
												if v10 > i32(-1) {
													goto l32
												}
												t46 := int32(m.memory[int64(uint32(v25))+1])
												v26 = t46 & i32(63)
												v27 = v10 & i32(31)
												if uint32(v10) >= uint32(i32(-32)) {
													t47 := int32(m.memory[int64(uint32(v25))+2])
													v26 = v26<<6 | t47&i32(63)
													if uint32(v10) >= uint32(i32(-16)) {
														t48 := int32(m.memory[int64(uint32(v25))+3])
														v10 = v26<<6 | t48&i32(63) | v27<<18&i32(0x1c0000)
														goto l34
													}
													v10 = v26 | v27<<12
													goto l34
												}
												v10 = v27<<6 | v26
												goto l34
											}
										l32:
											v10 = v10 & i32(255)
										l34:
											if v13&i32(1) != 0 {
												goto l36
											}
											{
												if uint32(v10) >= uint32(i32(128)) {
													goto l37
												}
												v10 = i32(1)
												goto l38
											l37:
												if uint32(v10) >= uint32(i32(2048)) {
													goto l39
												}
												v10 = i32(2)
												goto l38
											l39:
												p49 := i32(4)
												if uint32(v10) < uint32(i32(65536)) {
													p49 = i32(3)
												}
												v10 = p49
											}
										l38:
											{
												v8 = v10 + v8
												if v8 == 0 {
													goto l40
												}
												if uint32(v8) < uint32(v11) {
													goto l41
												}
												if v8 != v11 {
													goto l30
												}
												goto l40
											l41:
												t50 := int32(int8(m.memory[uint32(v12+v8)]))
												if t50 < i32(-64) {
													goto l30
												}
											}
										l40:
											if v8 == v11 {
												goto l36
											}
											t51 := int32(int8(m.memory[uint32(v12+v8)]))
											var p52 int32
											if t51 > i32(-1) {
												p52 = 1
											}
											_ = p52
											goto l36
										}
									l31:
										if v13&i32(1) == 0 {
											goto l27
										}
									l36:
										v10 = i32(1)
									}
								l27:
									store32(m.memory[int64(uint32(v3))+36:], uint32(v10))
									goto l42
								}
							}
							if v8 != v2 {
								goto l24
							}
							t37 := m.fn1909(v3+i32(32), v1, v2)
							if t37 == 0 {
								goto l25
							}
							goto l24
						}
					l43:
						m.fn206(v3+i32(36), v18, v8, v10, v11, v12, i32(1))
					l42:
						t70 := int32(load32(m.memory[int64(uint32(v3))+36:]))
						if t70 == 0 {
							goto l24
						}
					}
				l25:
					store64(m.memory[int64(uint32(v3))+48:], uint64(int64(uint32(i32(44)))<<32|int64(uint32(v3+i32(16)))))
					m.fn17(v3+i32(20), i32(1065113), v3+i32(48))
					m.fn163(v0+i32(4), i32(20), v3+i32(20))
					store32(m.memory[uint32(v0):], uint32(i32(-1)))
					goto l50
				l30:
					m.fn38(v12, v11, v8, v11, i32(1093504))
					panic("unreachable")
				l24:
					v9 = v9 + i32(4)
					if v9 != i32(16) {
						goto l51
					}
				}
				store32(m.memory[int64(uint32(v0))+8:], uint32(v6))
				store32(m.memory[int64(uint32(v0))+4:], uint32(v4))
				store32(m.memory[uint32(v0):], uint32(v5))
				goto l52
			}
		l18:
			v8 = i32(0)
			v10 = i32(0)
		l54:
			if v10&i32(0xffff) == 0 {
				goto l53
			}
			v10 = i32(0)
			v8 = v8 + i32(1)
			goto l54
		l53:
			if v1 == v7 {
				goto l55
			}
			{
				t71 := int32(int8(m.memory[uint32(v1)]))
				v10 = t71
				if v10 <= i32(-1) {
					if uint32(v10) >= uint32(i32(-32)) {
						v11 = v10 & i32(31)
						t72 := int32(m.memory[int64(uint32(v1))+1])
						t73 := int32(m.memory[int64(uint32(v1))+2])
						v12 = t72&i32(63)<<6 | t73&i32(63)
						{
							if uint32(v10) >= uint32(i32(-16)) {
								goto l58
							}
							v10 = v12 | v11<<12
							v1 = v1 + i32(3)
							goto l59
						l58:
							t74 := int32(m.memory[int64(uint32(v1))+3])
							v10 = v12<<6 | t74&i32(63) | v11<<18&i32(0x1c0000)
							v1 = v1 + i32(4)
						}
					l59:
						if uint32(v10) >= uint32(i32(65536)) {
							v10 = v10&i32(1023) | i32(-9216)
							v8 = v8 + i32(1)
							goto l54
						}
						v10 = i32(0)
						v8 = v8 + i32(1)
						goto l54
					}
					v1 = v1 + i32(2)
					v10 = i32(0)
					v8 = v8 + i32(1)
					goto l54
				}
				v1 = v1 + i32(1)
				v10 = i32(0)
				v8 = v8 + i32(1)
				goto l54
			}
		l55:
			store32(m.memory[int64(uint32(v3))+36:], uint32(v8))
			t75 := v3
			v28 = int64(uint32(i32(3))) << 32
			store64(m.memory[int64(uint32(t75))+56:], uint64(v28|int64(uint32(v3+i32(36)))))
			store64(m.memory[int64(uint32(v3))+48:], uint64(v28|int64(uint32(i32(1093356)))))
			m.fn17(v3+i32(4), i32(1067190), v3+i32(48))
			m.fn163(v0+i32(4), i32(20), v3+i32(4))
			store32(m.memory[uint32(v0):], uint32(i32(-1)))
		}
	l50:
		if v5 == 0 {
			goto l52
		}
		t76 := int32(load32(m.memory[uint32(v4+i32(-4)):]))
		v8 = t76
		v10 = v8 & i32(-8)
		t77 := v10
		v8 = v8 & i32(3)
		p78 := i32(8)
		if v8 != 0 {
			p78 = i32(4)
		}
		v1 = v5 << 1
		if uint32(t77) < uint32(p78+v1) {
			m.fn7(i32(1274404), i32(46), i32(1274452))
			panic("unreachable")
		}
		if v8 == 0 {
			goto l62
		}
		if uint32(v10) > uint32(v1+i32(39)) {
			m.fn7(i32(1274468), i32(46), i32(1274516))
			panic("unreachable")
		}
	l62:
		m.fn5(v4)
	}
l52:
	m.g0 = v3 + i32(112)
}
func (m *Module) fn298(v0, v1, v2 int32) {
	var v3, v4, v5 int32
	var v6, v7 int64
	var v8, v9 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t1 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		v3 = t1
		if t0 == v3 {
			goto l0
		}
		t2 := int32(load32(m.memory[uint32(v1):]))
		v4 = t2
		t3 := int32(load32(m.memory[int64(uint32(v4))+4:]))
		v5 = t3
		t4 := int64(load64(m.memory[int64(uint32(v4))+8:]))
		t5 := v5
		v6 = t4
		t6 := v6
		v7 = int64(uint32(v5))
		p7 := v7
		if uint64(v6) < uint64(v7) {
			p7 = t6
		}
		v8 = int32(p7)
		var p8 int32
		if t5 != v8 {
			p8 = 1
		}
		v9 = p8
		{
			var p9 int32
			if v5 == v8 {
				p9 = 1
			}
			v5 = p9
			if v5 != 0 {
				goto l1
			}
			t10 := int32(load32(m.memory[uint32(v4):]))
			t11 := int32(m.memory[uint32(t10+v8)])
			m.memory[uint32(v2)] = byte(t11)
		}
	l1:
		store32(m.memory[int64(uint32(v1))+8:], uint32(v3+v9))
		store64(m.memory[int64(uint32(v4))+8:], uint64(v6+int64(uint32(v9))))
		if v5 != 0 {
			goto l0
		}
		m.memory[uint32(v0)] = byte(i32(255))
		return
	}
l0:
	t12 := int64(load64(m.memory[int64(uint32(i32(0)))+1277248:]))
	store64(m.memory[uint32(v0):], uint64(t12))
}
func (m *Module) fn299(v0, v1 int32) {
	var v2, v3, v4, v5, v6, v7 int32
	var v8 int64
	var v9 int32
	var v10 int64
	var v11, v12, v13, v14, v15, v16, v17, v18, v19 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	store32(m.memory[int64(uint32(v2))+8:], uint32(i32(0)))
	t1 := int32(load32(m.memory[uint32(v1):]))
	v3 = t1
	t2 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	v4 = t2
	t3 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	v5 = t3
	v6 = i32(4)
	v7 = v2 + i32(8)
	{
	l4:
		if v5 != v4 {
			t4 := int32(load32(m.memory[uint32(v3):]))
			t5 := int64(load64(m.memory[int64(uint32(v3))+8:]))
			v8 = t5
			t6 := int32(load32(m.memory[int64(uint32(v3))+4:]))
			t7 := v8
			v9 = t6
			v10 = int64(uint32(v9))
			p8 := v10
			if uint64(v8) < uint64(v10) {
				p8 = t7
			}
			v11 = int32(p8)
			v12 = t4 + v11
			{
				t10 := v9
				p9 := i64(0xffffffff)
				if uint64(v8) < uint64(i64(0xffffffff)) {
					p9 = v8
				}
				v13 = t10 - int32(p9)
				p11 := v13
				if uint32(v13) > uint32(v9) {
					p11 = i32(0)
				}
				v13 = p11
				t12 := v13
				v14 = v5 - v4
				p13 := v6
				if uint32(v14) < uint32(v6) {
					p13 = v14
				}
				v14 = p13
				p14 := v14
				if uint32(v13) < uint32(v14) {
					p14 = t12
				}
				v13 = p14
				if v13 != i32(1) {
					goto l2
				}
				t15 := int32(m.memory[uint32(v12)])
				m.memory[uint32(v7)] = byte(t15)
				goto l3
			}
		l2:
			if v13 == 0 {
				goto l3
			}
			memory_copy(m.memory, uint32(v7), uint32(v12), uint32(v13))
		l3:
			t16 := v1
			v4 = v13 + v4
			store32(m.memory[int64(uint32(t16))+8:], uint32(v4))
			store64(m.memory[int64(uint32(v3))+8:], uint64(v8+int64(uint32(v13))))
			if v9 == v11 {
				goto l1
			}
			v7 = v7 + v13
			v6 = v6 - v13
			if v6 != 0 {
				goto l4
			}
			goto l5
		}
		v4 = v5
		goto l1
	l1:
		t17 := int64(load64(m.memory[int64(uint32(i32(0)))+1277248:]))
		v8 = t17
		if v8&i64(255) == i64(255) {
			goto l5
		}
		m.memory[uint32(v0)] = byte(i32(1))
		store64(m.memory[int64(uint32(v0))+4:], uint64(v8))
		return
	}
l5:
	t18 := int32(load32(m.memory[int64(uint32(v2))+8:]))
	v15 = t18
	store16(m.memory[int64(uint32(v2))+8:], uint16(i32(0)))
	v6 = i32(2)
	v7 = v2 + i32(8)
	{
	l10:
		if v5 != v4 {
			t19 := int32(load32(m.memory[uint32(v3):]))
			t20 := int64(load64(m.memory[int64(uint32(v3))+8:]))
			v8 = t20
			t21 := int32(load32(m.memory[int64(uint32(v3))+4:]))
			t22 := v8
			v9 = t21
			v10 = int64(uint32(v9))
			p23 := v10
			if uint64(v8) < uint64(v10) {
				p23 = t22
			}
			v11 = int32(p23)
			v12 = t19 + v11
			{
				t25 := v9
				p24 := i64(0xffffffff)
				if uint64(v8) < uint64(i64(0xffffffff)) {
					p24 = v8
				}
				v13 = t25 - int32(p24)
				p26 := v13
				if uint32(v13) > uint32(v9) {
					p26 = i32(0)
				}
				v13 = p26
				t27 := v13
				v14 = v5 - v4
				p28 := v6
				if uint32(v14) < uint32(v6) {
					p28 = v14
				}
				v14 = p28
				p29 := v14
				if uint32(v13) < uint32(v14) {
					p29 = t27
				}
				v13 = p29
				if v13 != i32(1) {
					goto l8
				}
				t30 := int32(m.memory[uint32(v12)])
				m.memory[uint32(v7)] = byte(t30)
				goto l9
			}
		l8:
			if v13 == 0 {
				goto l9
			}
			memory_copy(m.memory, uint32(v7), uint32(v12), uint32(v13))
		l9:
			t31 := v1
			v4 = v13 + v4
			store32(m.memory[int64(uint32(t31))+8:], uint32(v4))
			store64(m.memory[int64(uint32(v3))+8:], uint64(v8+int64(uint32(v13))))
			if v9 == v11 {
				goto l7
			}
			v7 = v7 + v13
			v6 = v6 - v13
			if v6 != 0 {
				goto l10
			}
			goto l11
		}
		v4 = v5
		goto l7
	l7:
		t32 := int64(load64(m.memory[int64(uint32(i32(0)))+1277248:]))
		v8 = t32
		if v8&i64(255) == i64(255) {
			goto l11
		}
		m.memory[uint32(v0)] = byte(i32(1))
		store64(m.memory[int64(uint32(v0))+4:], uint64(v8))
		return
	}
l11:
	t33 := int32(load16(m.memory[int64(uint32(v2))+8:]))
	v16 = t33
	store16(m.memory[int64(uint32(v2))+8:], uint16(i32(0)))
	v17 = int32(uint32(v16) >> 8)
	v6 = i32(2)
	v7 = v2 + i32(8)
	{
	l16:
		if v5 != v4 {
			t34 := int32(load32(m.memory[uint32(v3):]))
			t35 := int64(load64(m.memory[int64(uint32(v3))+8:]))
			v8 = t35
			t36 := int32(load32(m.memory[int64(uint32(v3))+4:]))
			t37 := v8
			v9 = t36
			v10 = int64(uint32(v9))
			p38 := v10
			if uint64(v8) < uint64(v10) {
				p38 = t37
			}
			v11 = int32(p38)
			v12 = t34 + v11
			{
				t40 := v9
				p39 := i64(0xffffffff)
				if uint64(v8) < uint64(i64(0xffffffff)) {
					p39 = v8
				}
				v13 = t40 - int32(p39)
				p41 := v13
				if uint32(v13) > uint32(v9) {
					p41 = i32(0)
				}
				v13 = p41
				t42 := v13
				v14 = v5 - v4
				p43 := v6
				if uint32(v14) < uint32(v6) {
					p43 = v14
				}
				v14 = p43
				p44 := v14
				if uint32(v13) < uint32(v14) {
					p44 = t42
				}
				v13 = p44
				if v13 != i32(1) {
					goto l14
				}
				t45 := int32(m.memory[uint32(v12)])
				m.memory[uint32(v7)] = byte(t45)
				goto l15
			}
		l14:
			if v13 == 0 {
				goto l15
			}
			memory_copy(m.memory, uint32(v7), uint32(v12), uint32(v13))
		l15:
			t46 := v1
			v4 = v13 + v4
			store32(m.memory[int64(uint32(t46))+8:], uint32(v4))
			store64(m.memory[int64(uint32(v3))+8:], uint64(v8+int64(uint32(v13))))
			if v9 == v11 {
				goto l13
			}
			v7 = v7 + v13
			v6 = v6 - v13
			if v6 != 0 {
				goto l16
			}
			goto l17
		}
		v4 = v5
		goto l13
	l13:
		t47 := int64(load64(m.memory[int64(uint32(i32(0)))+1277248:]))
		v8 = t47
		if v8&i64(255) == i64(255) {
			goto l17
		}
		m.memory[uint32(v0)] = byte(i32(1))
		store64(m.memory[int64(uint32(v0))+4:], uint64(v8))
		return
	}
l17:
	t48 := int32(load16(m.memory[int64(uint32(v2))+8:]))
	v18 = t48
	store64(m.memory[int64(uint32(v2))+8:], uint64(i64(0)))
	v6 = i32(8)
	v19 = int32(uint32(v18) >> 8)
	v7 = v2 + i32(8)
	{
	l21:
		{
			if v5 == v4 {
				goto l18
			}
			t49 := int32(load32(m.memory[uint32(v3):]))
			t50 := int64(load64(m.memory[int64(uint32(v3))+8:]))
			v8 = t50
			t51 := int32(load32(m.memory[int64(uint32(v3))+4:]))
			t52 := v8
			v9 = t51
			v10 = int64(uint32(v9))
			p53 := v10
			if uint64(v8) < uint64(v10) {
				p53 = t52
			}
			v11 = int32(p53)
			v12 = t49 + v11
			{
				t55 := v9
				p54 := i64(0xffffffff)
				if uint64(v8) < uint64(i64(0xffffffff)) {
					p54 = v8
				}
				v13 = t55 - int32(p54)
				p56 := v13
				if uint32(v13) > uint32(v9) {
					p56 = i32(0)
				}
				v13 = p56
				t57 := v13
				v14 = v5 - v4
				p58 := v6
				if uint32(v14) < uint32(v6) {
					p58 = v14
				}
				v14 = p58
				p59 := v14
				if uint32(v13) < uint32(v14) {
					p59 = t57
				}
				v13 = p59
				if v13 != i32(1) {
					goto l19
				}
				t60 := int32(m.memory[uint32(v12)])
				m.memory[uint32(v7)] = byte(t60)
				goto l20
			}
		l19:
			if v13 == 0 {
				goto l20
			}
			memory_copy(m.memory, uint32(v7), uint32(v12), uint32(v13))
		l20:
			t61 := v1
			v4 = v13 + v4
			store32(m.memory[int64(uint32(t61))+8:], uint32(v4))
			store64(m.memory[int64(uint32(v3))+8:], uint64(v8+int64(uint32(v13))))
			if v9 == v11 {
				goto l18
			}
			v7 = v7 + v13
			v6 = v6 - v13
			if v6 != 0 {
				goto l21
			}
			goto l22
		}
	l18:
		t62 := int64(load64(m.memory[int64(uint32(i32(0)))+1277248:]))
		v8 = t62
		if v8&i64(255) == i64(255) {
			goto l22
		}
		m.memory[uint32(v0)] = byte(i32(1))
		store64(m.memory[int64(uint32(v0))+4:], uint64(v8))
		return
	}
l22:
	t63 := int64(load64(m.memory[int64(uint32(v2))+8:]))
	store64(m.memory[int64(uint32(v0))+9:], uint64(t63))
	m.memory[int64(uint32(v0))+8] = byte(v18)
	m.memory[int64(uint32(v0))+7] = byte(v19)
	m.memory[int64(uint32(v0))+6] = byte(v16)
	m.memory[int64(uint32(v0))+5] = byte(v17)
	m.memory[int64(uint32(v0))+4] = byte(v15)
	m.memory[uint32(v0)] = byte(i32(0))
	m.memory[int64(uint32(v0))+3] = byte(int32(uint32(v15) >> 8))
	m.memory[int64(uint32(v0))+2] = byte(int32(uint32(v15) >> 16))
	m.memory[int64(uint32(v0))+1] = byte(int32(uint32(v15) >> 24))
}
func (m *Module) fn300(v0, v1 int32) {
	var v2, v3, v4, v5, v6, v7 int32
	var v8 int64
	var v9 int32
	var v10 int64
	var v11, v12, v13, v14 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	store64(m.memory[int64(uint32(v2))+8:], uint64(i64(0)))
	t1 := int32(load32(m.memory[uint32(v1):]))
	v3 = t1
	t2 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	v4 = t2
	t3 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	v5 = t3
	v6 = i32(8)
	v7 = v2 + i32(8)
	{
	l3:
		{
			if v5 == v4 {
				goto l0
			}
			t4 := int32(load32(m.memory[uint32(v3):]))
			t5 := int64(load64(m.memory[int64(uint32(v3))+8:]))
			v8 = t5
			t6 := int32(load32(m.memory[int64(uint32(v3))+4:]))
			t7 := v8
			v9 = t6
			v10 = int64(uint32(v9))
			p8 := v10
			if uint64(v8) < uint64(v10) {
				p8 = t7
			}
			v11 = int32(p8)
			v12 = t4 + v11
			{
				t10 := v9
				p9 := i64(0xffffffff)
				if uint64(v8) < uint64(i64(0xffffffff)) {
					p9 = v8
				}
				v13 = t10 - int32(p9)
				p11 := v13
				if uint32(v13) > uint32(v9) {
					p11 = i32(0)
				}
				v13 = p11
				t12 := v13
				v14 = v5 - v4
				p13 := v6
				if uint32(v14) < uint32(v6) {
					p13 = v14
				}
				v14 = p13
				p14 := v14
				if uint32(v13) < uint32(v14) {
					p14 = t12
				}
				v13 = p14
				if v13 != i32(1) {
					goto l1
				}
				t15 := int32(m.memory[uint32(v12)])
				m.memory[uint32(v7)] = byte(t15)
				goto l2
			}
		l1:
			if v13 == 0 {
				goto l2
			}
			memory_copy(m.memory, uint32(v7), uint32(v12), uint32(v13))
		l2:
			t16 := v1
			v4 = v13 + v4
			store32(m.memory[int64(uint32(t16))+8:], uint32(v4))
			store64(m.memory[int64(uint32(v3))+8:], uint64(v8+int64(uint32(v13))))
			if v9 == v11 {
				goto l0
			}
			v7 = v7 + v13
			v6 = v6 - v13
			if v6 != 0 {
				goto l3
			}
			goto l4
		}
	l0:
		t17 := int64(load64(m.memory[int64(uint32(i32(0)))+1277248:]))
		v8 = t17
		if v8&i64(255) == i64(255) {
			goto l4
		}
		store64(m.memory[int64(uint32(v0))+4:], uint64(v8))
		store32(m.memory[uint32(v0):], uint32(i32(1)))
		return
	}
l4:
	t18 := int64(load64(m.memory[int64(uint32(v2))+8:]))
	store64(m.memory[int64(uint32(v0))+8:], uint64(t18))
	store32(m.memory[uint32(v0):], uint32(i32(0)))
}
func (m *Module) fn301(v0, v1 int32) {
	var v2, v3, v4, v5, v6, v7 int32
	var v8 int64
	var v9 int32
	var v10 int64
	var v11, v12, v13, v14 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	store64(m.memory[int64(uint32(v2))+8:], uint64(i64(0)))
	t1 := int32(load32(m.memory[uint32(v1):]))
	v3 = t1
	t2 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	v4 = t2
	t3 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	v5 = t3
	v6 = i32(8)
	v7 = v2 + i32(8)
	{
	l3:
		{
			if v5 == v4 {
				goto l0
			}
			t4 := int32(load32(m.memory[uint32(v3):]))
			t5 := int64(load64(m.memory[int64(uint32(v3))+8:]))
			v8 = t5
			t6 := int32(load32(m.memory[int64(uint32(v3))+4:]))
			t7 := v8
			v9 = t6
			v10 = int64(uint32(v9))
			p8 := v10
			if uint64(v8) < uint64(v10) {
				p8 = t7
			}
			v11 = int32(p8)
			v12 = t4 + v11
			{
				t10 := v9
				p9 := i64(0xffffffff)
				if uint64(v8) < uint64(i64(0xffffffff)) {
					p9 = v8
				}
				v13 = t10 - int32(p9)
				p11 := v13
				if uint32(v13) > uint32(v9) {
					p11 = i32(0)
				}
				v13 = p11
				t12 := v13
				v14 = v5 - v4
				p13 := v6
				if uint32(v14) < uint32(v6) {
					p13 = v14
				}
				v14 = p13
				p14 := v14
				if uint32(v13) < uint32(v14) {
					p14 = t12
				}
				v13 = p14
				if v13 != i32(1) {
					goto l1
				}
				t15 := int32(m.memory[uint32(v12)])
				m.memory[uint32(v7)] = byte(t15)
				goto l2
			}
		l1:
			if v13 == 0 {
				goto l2
			}
			memory_copy(m.memory, uint32(v7), uint32(v12), uint32(v13))
		l2:
			t16 := v1
			v4 = v13 + v4
			store32(m.memory[int64(uint32(t16))+8:], uint32(v4))
			store64(m.memory[int64(uint32(v3))+8:], uint64(v8+int64(uint32(v13))))
			if v9 == v11 {
				goto l0
			}
			v7 = v7 + v13
			v6 = v6 - v13
			if v6 != 0 {
				goto l3
			}
			goto l4
		}
	l0:
		t17 := int64(load64(m.memory[int64(uint32(i32(0)))+1277248:]))
		v8 = t17
		if v8&i64(255) == i64(255) {
			goto l4
		}
		store64(m.memory[int64(uint32(v0))+4:], uint64(v8))
		store32(m.memory[uint32(v0):], uint32(i32(1)))
		return
	}
l4:
	t18 := int64(load64(m.memory[int64(uint32(v2))+8:]))
	store64(m.memory[int64(uint32(v0))+8:], uint64(t18))
	store32(m.memory[uint32(v0):], uint32(i32(0)))
}
func (m *Module) fn302(v0, v1 int32, v2 int64) {
	var v3, v4, v5 int32
	t0 := m.g0
	v3 = t0 - i32(32)
	m.g0 = v3
	{
		if v1&i32(255) != 0 {
			m.fn303(v3+i32(12), v1, v2)
			store64(m.memory[int64(uint32(v3))+24:], uint64(int64(uint32(i32(17)))<<32|int64(uint32(v3+i32(12)))))
			m.fn17(v0, i32(1066600), v3+i32(24))
			t2 := int32(load32(m.memory[int64(uint32(v3))+12:]))
			v0 = t2
			if v0 == 0 {
				goto l2
			}
			t3 := int32(load32(m.memory[int64(uint32(v3))+16:]))
			v4 = t3
			t4 := int32(load32(m.memory[uint32(v4+i32(-4)):]))
			v1 = t4
			v5 = v1 & i32(-8)
			t5 := v5
			v1 = v1 & i32(3)
			p6 := i32(8)
			if v1 != 0 {
				p6 = i32(4)
			}
			if uint32(t5) < uint32(p6+v0) {
				m.fn7(i32(1274404), i32(46), i32(1274452))
				panic("unreachable")
			}
			if v1 == 0 {
				goto l4
			}
			if uint32(v5) > uint32(v0+i32(39)) {
				m.fn7(i32(1274468), i32(46), i32(1274516))
				panic("unreachable")
			}
		l4:
			m.fn5(v4)
			goto l2
		}
		t1 := m.fn11(i32(1))
		v1 = t1
		if v1 != 0 {
			goto l1
		}
		m.fn16(i32(1), i32(1))
		panic("unreachable")
	}
l1:
	store32(m.memory[int64(uint32(v0))+8:], uint32(i32(1)))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
	store32(m.memory[uint32(v0):], uint32(i32(1)))
	m.memory[uint32(v1)] = byte(i32(45))
l2:
	m.g0 = v3 + i32(32)
}
func (m *Module) fn303(v0, v1 int32, v2 int64) {
	var v3 int32
	var v4 int64
	var v5 int32
	var v6 int64
	var v7, v8, v9, v10 int32
	t0 := m.g0
	v3 = t0 - i32(32)
	m.g0 = v3
	{
		{
			{
				{
					switch v1 & i32(255) {
					case 1:
						v1 = i32(20)
						v4 = v2
						if uint64(v2) < uint64(i64(1000)) {
							goto l6
						}
						v1 = i32(20)
						v4 = v2
					l7:
						{
							v5 = v3 + i32(12) + v1
							t1 := v5 + i32(-4)
							v6 = v4
							t2 := int64(uint64(v6) / uint64(i64(10000)))
							t3 := v6
							v4 = t2
							v7 = int32(t3 - v4*i64(10000))
							t4 := int32(uint32(v7&i32(0xffff)) / uint32(i32(100)))
							v8 = t4
							t5 := int32(load16(m.memory[int64(uint32(v8<<1))+1100735:]))
							store16(m.memory[uint32(t1):], uint16(t5))
							t6 := int32(load16(m.memory[int64(uint32((v7-v8*i32(100))&i32(0xffff)<<1))+1100735:]))
							store16(m.memory[uint32(v5+i32(-2)):], uint16(t6))
							v1 = v1 + i32(-4)
							if uint64(v6) > uint64(i64(9999999)) {
								goto l7
							}
						}
					l6:
						{
							if uint64(v4) <= uint64(i64(9)) {
								goto l8
							}
							t7 := v3 + i32(12)
							v1 = v1 + i32(-2)
							t8 := t7 + v1
							v5 = int32(v4)
							t9 := int32(uint32(v5&i32(0xffff)) / uint32(i32(100)))
							t10 := v5
							v5 = t9
							t11 := int32(load16(m.memory[int64(uint32((t10-v5*i32(100))&i32(0xffff)<<1))+1100735:]))
							store16(m.memory[uint32(t8):], uint16(t11))
							v4 = int64(uint32(v5))
						}
					l8:
						{
							if v2 == 0 {
								goto l9
							}
							if v4 == 0 {
								goto l10
							}
						l9:
							v8 = v3 + i32(12) + v1 + i32(-1)
							t12 := int32(m.memory[int64(uint32(int32(v4)<<1))+1100736])
							m.memory[uint32(v8)] = byte(t12)
							v5 = i32(21) - v1
							goto l11
						}
					l10:
						v5 = i32(20) - v1
						v7 = i32(1)
						if v1 == i32(20) {
							goto l12
						}
						v8 = v3 + i32(12) + v1
					l11:
						t13 := m.fn11(v5)
						v7 = t13
						if v7 != 0 {
							goto l13
						}
						m.fn16(i32(1), v5)
						panic("unreachable")
					case 2:
						m.fn304(v0, v2)
						goto l14
					case 3:
						m.fn304(v3+i32(12), v2)
						t14 := int32(load32(m.memory[int64(uint32(v3))+20:]))
						v5 = t14
						if v5 <= i32(-1) {
							goto l15
						}
						t15 := int32(load32(m.memory[int64(uint32(v3))+16:]))
						v9 = t15
						if v5 != 0 {
							t16 := m.fn11(v5)
							v8 = t16
							if v8 != 0 {
								goto l18
							}
							m.fn16(i32(1), v5)
							panic("unreachable")
						}
						store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
						store64(m.memory[uint32(v0):], uint64(i64(0x100000000)))
						goto l17
					case 4:
						m.fn305(v0, v2)
						goto l14
					case 5:
						m.fn305(v3+i32(12), v2)
						t17 := int32(load32(m.memory[int64(uint32(v3))+20:]))
						v5 = t17
						if v5 <= i32(-1) {
							goto l15
						}
						t18 := int32(load32(m.memory[int64(uint32(v3))+16:]))
						v9 = t18
						if v5 != 0 {
							t19 := m.fn11(v5)
							v8 = t19
							if v8 != 0 {
								goto l21
							}
							m.fn16(i32(1), v5)
							panic("unreachable")
						}
						store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
						store64(m.memory[uint32(v0):], uint64(i64(0x100000000)))
						goto l20
					default:
						t20 := m.fn11(i32(1))
						v1 = t20
						if v1 == 0 {
							m.fn16(i32(1), i32(1))
							panic("unreachable")
						}
						store32(m.memory[int64(uint32(v0))+8:], uint32(i32(1)))
						store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
						store32(m.memory[uint32(v0):], uint32(i32(1)))
						m.memory[uint32(v1)] = byte(i32(45))
						goto l14
					}
				l13:
					if v5 == 0 {
						goto l12
					}
					memory_copy(m.memory, uint32(v7), uint32(v8), uint32(v5))
				l12:
					store32(m.memory[int64(uint32(v0))+8:], uint32(v5))
					store32(m.memory[int64(uint32(v0))+4:], uint32(v7))
					store32(m.memory[uint32(v0):], uint32(v5))
					goto l14
				l18:
					if v5 == 0 {
						goto l23
					}
					memory_copy(m.memory, uint32(v8), uint32(v9), uint32(v5))
				l23:
					store32(m.memory[int64(uint32(v0))+8:], uint32(v5))
					store32(m.memory[int64(uint32(v0))+4:], uint32(v8))
					store32(m.memory[uint32(v0):], uint32(v5))
					v1 = i32(0)
					if v5 == i32(1) {
						goto l24
					}
					v10 = v5 & i32(1)
					v0 = v5 & i32(0x7ffffffe)
					v1 = i32(0)
				l25:
					{
						v5 = v8 + v1
						t21 := int32(m.memory[uint32(v5)])
						t22 := v5
						v7 = t21
						p23 := i32(0)
						if uint32((v7+i32(-97))&i32(255)) < uint32(i32(26)) {
							p23 = i32(32)
						}
						m.memory[uint32(t22)] = byte(p23 ^ v7)
						v5 = v5 + i32(1)
						t24 := int32(m.memory[uint32(v5)])
						t25 := v5
						v5 = t24
						p26 := i32(0)
						if uint32((v5+i32(-97))&i32(255)) < uint32(i32(26)) {
							p26 = i32(32)
						}
						m.memory[uint32(t25)] = byte(p26 ^ v5)
						t27 := v0
						v1 = v1 + i32(2)
						if t27 != v1 {
							goto l25
						}
					}
					if v10 == 0 {
						goto l17
					}
				l24:
					v1 = v8 + v1
					t28 := int32(m.memory[uint32(v1)])
					t29 := v1
					v1 = t28
					p30 := i32(0)
					if uint32((v1+i32(-97))&i32(255)) < uint32(i32(26)) {
						p30 = i32(32)
					}
					m.memory[uint32(t29)] = byte(p30 ^ v1)
				}
			l17:
				t31 := int32(load32(m.memory[int64(uint32(v3))+12:]))
				v1 = t31
				if v1 == 0 {
					goto l14
				}
				t32 := int32(load32(m.memory[uint32(v9+i32(-4)):]))
				v5 = t32
				v7 = v5 & i32(-8)
				t33 := v7
				v5 = v5 & i32(3)
				p34 := i32(8)
				if v5 != 0 {
					p34 = i32(4)
				}
				if uint32(t33) < uint32(p34+v1) {
					m.fn7(i32(1274404), i32(46), i32(1274452))
					panic("unreachable")
				}
				if v5 == 0 {
					goto l27
				}
				if uint32(v7) > uint32(v1+i32(39)) {
					m.fn7(i32(1274468), i32(46), i32(1274516))
					panic("unreachable")
				}
			l27:
				m.fn5(v9)
				goto l14
			}
		l21:
			if v5 == 0 {
				goto l29
			}
			memory_copy(m.memory, uint32(v8), uint32(v9), uint32(v5))
		l29:
			store32(m.memory[int64(uint32(v0))+8:], uint32(v5))
			store32(m.memory[int64(uint32(v0))+4:], uint32(v8))
			store32(m.memory[uint32(v0):], uint32(v5))
			v1 = i32(0)
			if v5 == i32(1) {
				goto l30
			}
			v10 = v5 & i32(1)
			v0 = v5 & i32(0x7ffffffe)
			v1 = i32(0)
		l31:
			{
				v5 = v8 + v1
				t35 := int32(m.memory[uint32(v5)])
				t36 := v5
				v7 = t35
				p37 := i32(0)
				if uint32((v7+i32(-97))&i32(255)) < uint32(i32(26)) {
					p37 = i32(32)
				}
				m.memory[uint32(t36)] = byte(p37 ^ v7)
				v5 = v5 + i32(1)
				t38 := int32(m.memory[uint32(v5)])
				t39 := v5
				v5 = t38
				p40 := i32(0)
				if uint32((v5+i32(-97))&i32(255)) < uint32(i32(26)) {
					p40 = i32(32)
				}
				m.memory[uint32(t39)] = byte(p40 ^ v5)
				t41 := v0
				v1 = v1 + i32(2)
				if t41 != v1 {
					goto l31
				}
			}
			if v10 == 0 {
				goto l20
			}
		l30:
			v1 = v8 + v1
			t42 := int32(m.memory[uint32(v1)])
			t43 := v1
			v1 = t42
			p44 := i32(0)
			if uint32((v1+i32(-97))&i32(255)) < uint32(i32(26)) {
				p44 = i32(32)
			}
			m.memory[uint32(t43)] = byte(p44 ^ v1)
		}
	l20:
		t45 := int32(load32(m.memory[int64(uint32(v3))+12:]))
		v1 = t45
		if v1 == 0 {
			goto l14
		}
		t46 := int32(load32(m.memory[uint32(v9+i32(-4)):]))
		v5 = t46
		v7 = v5 & i32(-8)
		t47 := v7
		v5 = v5 & i32(3)
		p48 := i32(8)
		if v5 != 0 {
			p48 = i32(4)
		}
		if uint32(t47) < uint32(p48+v1) {
			m.fn7(i32(1274404), i32(46), i32(1274452))
			panic("unreachable")
		}
		if v5 == 0 {
			goto l33
		}
		if uint32(v7) > uint32(v1+i32(39)) {
			m.fn7(i32(1274468), i32(46), i32(1274516))
			panic("unreachable")
		}
	l33:
		m.fn5(v9)
		goto l14
	}
l15:
	m.fn15()
	panic("unreachable")
l14:
	m.g0 = v3 + i32(32)
}
func (m *Module) fn304(v0 int32, v1 int64) {
	var v2, v3, v4 int32
	var v5 int64
	var v6, v7, v8, v9, v10, v11, v12, v13 int32
	t0 := m.g0
	v2 = t0 - i32(32)
	m.g0 = v2
	{
		if v1 != i64(0) {
			v3 = i32(0)
			store32(m.memory[int64(uint32(v2))+16:], uint32(i32(0)))
			store64(m.memory[int64(uint32(v2))+8:], uint64(i64(0x100000000)))
			v4 = i32(1)
		l5:
			{
				v5 = v1 + i64(-1)
				t2 := int64(uint64(v5) / uint64(i64(26)))
				t3 := v5
				v1 = t2
				v6 = int32(t3-v1*i64(26)) + i32(97)
				{
					t4 := int32(load32(m.memory[int64(uint32(v2))+8:]))
					if v3 != t4 {
						goto l3
					}
					m.fn39(v2 + i32(8))
					t5 := int32(load32(m.memory[int64(uint32(v2))+12:]))
					v4 = t5
				}
			l3:
				m.memory[uint32(v4+v3)] = byte(v6)
				t6 := v2
				v3 = v3 + i32(1)
				store32(m.memory[int64(uint32(t6))+16:], uint32(v3))
				if uint64(v5) < uint64(i64(26)) {
					t7 := int32(load32(m.memory[int64(uint32(v2))+12:]))
					v7 = t7
					{
						v8 = int32(uint32(v3) >> 1)
						if v8 == 0 {
							goto l6
						}
						v9 = v7 + v3
						v6 = i32(0)
						if v8 == i32(1) {
							goto l7
						}
						v4 = v3 + i32(-1)
						v10 = v8 & i32(1)
						v11 = int32(uint32(v3)>>1) & i32(0x7ffffffe)
						v6 = i32(0)
					l8:
						{
							v12 = v7 + v4
							t8 := int32(m.memory[uint32(v12)])
							v13 = t8
							t9 := v12
							v8 = v7 + v6
							t10 := int32(m.memory[uint32(v8)])
							m.memory[uint32(t9)] = byte(t10)
							m.memory[uint32(v8)] = byte(v13)
							v12 = v9 + (v6 ^ i32(-2))
							t11 := int32(m.memory[uint32(v12)])
							v13 = t11
							t12 := v12
							v8 = v8 + i32(1)
							t13 := int32(m.memory[uint32(v8)])
							m.memory[uint32(t12)] = byte(t13)
							m.memory[uint32(v8)] = byte(v13)
							v4 = v4 + i32(-2)
							t14 := v11
							v6 = v6 + i32(2)
							if t14 != v6 {
								goto l8
							}
						}
						if v10 == 0 {
							goto l6
						}
					l7:
						v4 = v7 + v6
						t15 := int32(m.memory[uint32(v4)])
						v8 = t15
						t16 := v4
						v6 = v9 + (v6 ^ i32(-1))
						t17 := int32(m.memory[uint32(v6)])
						m.memory[uint32(t16)] = byte(t17)
						m.memory[uint32(v6)] = byte(v8)
					}
				l6:
					t18 := int32(load32(m.memory[int64(uint32(v2))+8:]))
					v6 = t18
					m.fn14(v2+i32(20), v7, v3)
					{
						{
							t19 := int32(load32(m.memory[int64(uint32(v2))+20:]))
							if t19 != 0 {
								goto l9
							}
							v5 = int64(uint32(v3))
							goto l10
						}
					l9:
						if v6 != i32(-1) {
							store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
							store64(m.memory[uint32(v0):], uint64(i64(0x100000000)))
							if v6 == 0 {
								goto l2
							}
							{
								t21 := int32(load32(m.memory[uint32(v7+i32(-4)):]))
								v3 = t21
								v4 = v3 & i32(-8)
								t22 := v4
								v3 = v3 & i32(3)
								p23 := i32(8)
								if v3 != 0 {
									p23 = i32(4)
								}
								if uint32(t22) < uint32(p23+v6) {
									m.fn7(i32(1274404), i32(46), i32(1274452))
									panic("unreachable")
								}
								if v3 == 0 {
									goto l13
								}
								if uint32(v4) > uint32(v6+i32(39)) {
									m.fn7(i32(1274468), i32(46), i32(1274516))
									panic("unreachable")
								}
							l13:
								m.fn5(v7)
								goto l2
							}
						}
						t20 := int64(load64(m.memory[int64(uint32(v2))+24:]))
						v5 = t20
						v6 = v7
						v7 = v3
					}
				l10:
					store32(m.memory[int64(uint32(v0))+8:], uint32(v5))
					store32(m.memory[int64(uint32(v0))+4:], uint32(v7))
					store32(m.memory[uint32(v0):], uint32(v6))
					goto l2
				}
				goto l5
			}
		}
		t1 := m.fn11(i32(1))
		v3 = t1
		if v3 == 0 {
			m.fn16(i32(1), i32(1))
			panic("unreachable")
		}
		store32(m.memory[int64(uint32(v0))+8:], uint32(i32(1)))
		store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
		store32(m.memory[uint32(v0):], uint32(i32(1)))
		m.memory[uint32(v3)] = byte(i32(48))
		goto l2
	}
l2:
	m.g0 = v2 + i32(32)
}
func (m *Module) fn305(v0 int32, v1 int64) {
	var v2, v3, v4 int32
	var v5 int64
	var v6, v7, v8, v9 int32
	t0 := m.g0
	v2 = t0 - i32(240)
	m.g0 = v2
	{
		if uint64(v1+i64(-4000)) > uint64(i64(-4000)) {
			v3 = i32(0)
			store32(m.memory[int64(uint32(v2))+20:], uint32(i32(0)))
			store64(m.memory[int64(uint32(v2))+12:], uint64(i64(0x100000000)))
			v8 = v2 + i32(32)
			memory_copy(m.memory, uint32(v8), uint32(i32(1075952)), uint32(i32(208)))
			v7 = i32(1)
			v9 = i32(0)
		l13:
			{
				v4 = v8 + v9<<4
				t17 := int32(load32(m.memory[int64(uint32(v4))+8:]))
				v6 = t17
				if v6 == 0 {
					goto l6
				}
				{
					t18 := int64(load64(m.memory[uint32(v4):]))
					t19 := v1
					v5 = t18
					if uint64(t19) < uint64(v5) {
						goto l7
					}
					{
						t20 := int32(load32(m.memory[int64(uint32(v4))+12:]))
						v4 = t20
						if v4 != 0 {
							goto l12
						}
						store32(m.memory[int64(uint32(v2))+20:], uint32(v3))
					l9:
						v1 = v1 - v5
						if uint64(v1) >= uint64(v5) {
							goto l9
						}
						goto l7
					}
				l12:
					{
						{
							t21 := int32(load32(m.memory[int64(uint32(v2))+12:]))
							if uint32(v4) <= uint32(t21-v3) {
								goto l10
							}
							m.fn197(v2+i32(12), v3, v4, i32(1), i32(1))
							t22 := int32(load32(m.memory[int64(uint32(v2))+16:]))
							v7 = t22
							t23 := int32(load32(m.memory[int64(uint32(v2))+20:]))
							v3 = t23
						}
					l10:
						if v4 == 0 {
							goto l11
						}
						memory_copy(m.memory, uint32(v7+v3), uint32(v6), uint32(v4))
					l11:
						t24 := v2
						v3 = v3 + v4
						store32(m.memory[int64(uint32(t24))+20:], uint32(v3))
						v1 = v1 - v5
						if uint64(v1) >= uint64(v5) {
							goto l12
						}
					}
				}
			l7:
				v9 = v9 + i32(1)
				if v9 != i32(13) {
					goto l13
				}
			}
		l6:
			t25 := int32(load32(m.memory[int64(uint32(v2))+20:]))
			store32(m.memory[int64(uint32(v0))+8:], uint32(t25))
			t26 := int64(load64(m.memory[int64(uint32(v2))+12:]))
			store64(m.memory[uint32(v0):], uint64(t26))
			goto l14
		}
		v3 = i32(20)
		{
			if uint64(v1) < uint64(i64(1000)) {
				goto l1
			}
			v4 = i32(24)
		l2:
			{
				t1 := v2 + i32(24)
				v3 = v4
				v4 = t1 + v3
				t2 := v4 + i32(-8)
				v5 = v1
				t3 := int64(uint64(v5) / uint64(i64(10000)))
				t4 := v5
				v1 = t3
				v6 = int32(t4 - v1*i64(10000))
				t5 := int32(uint32(v6&i32(0xffff)) / uint32(i32(100)))
				v7 = t5
				t6 := int32(load16(m.memory[int64(uint32(v7<<1))+1100735:]))
				store16(m.memory[uint32(t2):], uint16(t6))
				t7 := int32(load16(m.memory[int64(uint32((v6-v7*i32(100))&i32(0xffff)<<1))+1100735:]))
				store16(m.memory[uint32(v4+i32(-6)):], uint16(t7))
				v4 = v3 + i32(-4)
				if uint64(v5) > uint64(i64(9999999)) {
					goto l2
				}
			}
			v3 = v3 + i32(-8)
			{
				if uint64(v5) <= uint64(i64(99999)) {
					goto l3
				}
				t8 := v3 + (v2 + i32(24)) + i32(-1)
				v3 = int32(v1)
				t9 := int32(uint32(v3&i32(0xffff)) / uint32(i32(100)))
				t10 := v3
				v6 = t9
				v7 = (t10 - v6*i32(100)) & i32(0xffff) << 1
				t11 := int32(m.memory[int64(uint32(v7))+1100736])
				m.memory[uint32(t8)] = byte(t11)
				t12 := v2 + i32(24)
				v3 = v4 + i32(-6)
				t13 := int32(m.memory[int64(uint32(v7))+1100735])
				m.memory[uint32(t12+v3)] = byte(t13)
				v1 = int64(uint32(v6))
			}
		l3:
			if v1 == 0 {
				goto l4
			}
		l1:
			t14 := v2 + i32(24)
			v3 = v3 + i32(-1)
			t15 := int32(m.memory[int64(uint32(int32(v1)<<1))+1100736])
			m.memory[uint32(t14+v3)] = byte(t15)
		}
	l4:
		v4 = i32(20) - v3
		t16 := m.fn11(v4)
		v6 = t16
		if v6 != 0 {
			goto l5
		}
		m.fn16(i32(1), v4)
		panic("unreachable")
	}
l5:
	if v4 == 0 {
		goto l15
	}
	memory_copy(m.memory, uint32(v6), uint32(v2+i32(24)+v3), uint32(v4))
l15:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v4))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v6))
	store32(m.memory[uint32(v0):], uint32(v4))
l14:
	m.g0 = v2 + i32(240)
}
func (m *Module) fn306(v0 int32) int32 {
	var v1, v2, v3 int32
	t0 := m.g0
	v1 = t0 - i32(16)
	m.g0 = v1
	v2 = i32(0)
	{
		t1 := int32(load32(m.memory[uint32(v0):]))
		v3 = t1
		p2 := i32(1)
		if uint32(v3) > uint32(i32(2)) {
			p2 = v3 + i32(-3)
		}
		switch p2 {
		case 2, 4:
			goto l2
		case 3, 5:
			goto l3
		default:
			t3 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			t4 := int32(load32(m.memory[int64(uint32(v0))+12:]))
			m.fn144(v1+i32(8), t3, t4)
			t5 := int32(load32(m.memory[int64(uint32(v1))+12:]))
			var p6 int32
			if t5 == 0 {
				p6 = 1
			}
			v2 = p6
			goto l2
		case 1:
			t7 := int32(load32(m.memory[int64(uint32(v0))+12:]))
			if t7 != 0 {
				goto l2
			}
			t8 := int32(load32(m.memory[int64(uint32(v0))+24:]))
			v3 = t8 * i32(28)
			t9 := int32(load32(m.memory[int64(uint32(v0))+20:]))
			v0 = t9 + i32(-28)
		l4:
			{
				if v3 == 0 {
					goto l3
				}
				v3 = v3 + i32(-28)
				v0 = v0 + i32(28)
				t10 := m.fn306(v0)
				if t10 == 0 {
					goto l2
				}
				goto l4
			}
		}
	}
l3:
	v2 = i32(1)
l2:
	m.g0 = v1 + i32(16)
	return v2
}
func (m *Module) fn307(v0, v1, v2, v3, v4, v5 int32) int32 {
	var v6, v7, v8, v9, v10, v11, v12, v13, v14, v15, v16 int32
	t0 := m.g0
	v6 = t0 - i32(16)
	m.g0 = v6
	{
		if v1 != 0 {
			goto l0
		}
		v7 = i32(4)
		v8 = i32(0)
		goto l1
	l0:
		v9 = v1 << 2
		t1 := m.fn11(v9)
		v7 = t1
		if v7 == 0 {
			m.fn16(i32(4), v9)
			panic("unreachable")
		}
		v9 = v1*i32(44) + i32(-44)
		t2 := int32(uint32(v9) / uint32(i32(44)))
		v10 = t2 + i32(1)
		v11 = v10 & i32(7)
		v8 = i32(0)
		if uint32(v9) < uint32(i32(308)) {
			goto l3
		}
		v8 = v10 & i32(0xffffff8)
		v12 = v10 << 2 & i32(0x3fffffe0)
		v13 = i32(0)
	l4:
		{
			v9 = v7 + v13
			store32(m.memory[uint32(v9):], uint32(v0))
			store32(m.memory[uint32(v9+i32(28)):], uint32(v0+i32(308)))
			store32(m.memory[uint32(v9+i32(24)):], uint32(v0+i32(264)))
			store32(m.memory[uint32(v9+i32(20)):], uint32(v0+i32(220)))
			store32(m.memory[uint32(v9+i32(16)):], uint32(v0+i32(176)))
			store32(m.memory[uint32(v9+i32(12)):], uint32(v0+i32(132)))
			store32(m.memory[uint32(v9+i32(8)):], uint32(v0+i32(88)))
			store32(m.memory[uint32(v9+i32(4)):], uint32(v0+i32(44)))
			v0 = v0 + i32(352)
			t3 := v12
			v13 = v13 + i32(32)
			if t3 != v13 {
				goto l4
			}
		}
		if v11 == 0 {
			goto l5
		}
	l3:
		v12 = v8 + v11
		v13 = v11 << 2
		v9 = v7 + v8<<2
	l6:
		store32(m.memory[uint32(v9):], uint32(v0))
		v9 = v9 + i32(4)
		v0 = v0 + i32(44)
		v13 = v13 + i32(-4)
		if v13 != 0 {
			goto l6
		}
		v8 = v12
		if uint32(v12) >= uint32(i32(2)) {
			goto l5
		}
		v8 = i32(1)
		goto l1
	l5:
		v14 = v7 + v8<<2
		v13 = i32(0)
		v0 = int32(uint32(v10) >> 1)
		if v0 == i32(1) {
			goto l7
		}
		v15 = v0 & i32(1)
		v16 = v0 & i32(0x7fffffe)
		v9 = v14 + i32(-4)
		v13 = i32(0)
		v0 = v7
	l8:
		{
			t4 := int32(load32(m.memory[uint32(v9):]))
			v12 = t4
			t5 := int32(load32(m.memory[uint32(v0):]))
			store32(m.memory[uint32(v9):], uint32(t5))
			store32(m.memory[uint32(v0):], uint32(v12))
			v12 = v14 + (v13^i32(0x3ffffffe))<<2
			t6 := int32(load32(m.memory[uint32(v12):]))
			v11 = t6
			t7 := v12
			v10 = v0 + i32(4)
			t8 := int32(load32(m.memory[uint32(v10):]))
			store32(m.memory[uint32(t7):], uint32(t8))
			store32(m.memory[uint32(v10):], uint32(v11))
			v9 = v9 + i32(-8)
			v0 = v0 + i32(8)
			t9 := v16
			v13 = v13 + i32(2)
			if t9 != v13 {
				goto l8
			}
		}
		if v15 == 0 {
			goto l1
		}
	l7:
		v0 = v7 + v13<<2
		t10 := int32(load32(m.memory[uint32(v0):]))
		v9 = t10
		t11 := v0
		v13 = v14 + (v13^i32(-1))<<2
		t12 := int32(load32(m.memory[uint32(v13):]))
		store32(m.memory[uint32(t11):], uint32(t12))
		store32(m.memory[uint32(v13):], uint32(v9))
	}
l1:
	store32(m.memory[int64(uint32(v6))+12:], uint32(v8))
	store32(m.memory[int64(uint32(v6))+8:], uint32(v7))
	store32(m.memory[int64(uint32(v6))+4:], uint32(v1))
l11:
	{
		{
			t13 := m.fn151(v6 + i32(4))
			v0 = t13
			if v0 != 0 {
				goto l9
			}
			v0 = i32(0)
			goto l10
		}
	l9:
		t14 := int32(load32(m.memory[uint32(v0):]))
		if t14 == i32(-1) {
			goto l11
		}
		t15 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		if t15 != v5 {
			goto l11
		}
		t16 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		t17 := m.fn1909(t16, v4, v5)
		if t17 != 0 {
			goto l11
		}
		t18 := int32(load32(m.memory[int64(uint32(v0))+36:]))
		v9 = t18
		if v9 == 0 {
			goto l11
		}
		t19 := int32(load32(m.memory[int64(uint32(v0))+40:]))
		if t19 != v3 {
			goto l11
		}
		t20 := m.fn1909(v9+i32(8), v2, v3)
		if t20 != 0 {
			goto l11
		}
	}
l10:
	{
		t21 := int32(load32(m.memory[int64(uint32(v6))+4:]))
		v9 = t21
		if v9 == 0 {
			goto l12
		}
		t22 := int32(load32(m.memory[int64(uint32(v6))+8:]))
		v5 = t22
		t23 := int32(load32(m.memory[uint32(v5+i32(-4)):]))
		v13 = t23
		v7 = v13 & i32(-8)
		t24 := v7
		v13 = v13 & i32(3)
		p25 := i32(8)
		if v13 != 0 {
			p25 = i32(4)
		}
		v9 = v9 << 2
		if uint32(t24) < uint32(p25+v9) {
			m.fn7(i32(1274404), i32(46), i32(1274452))
			panic("unreachable")
		}
		if v13 == 0 {
			goto l14
		}
		if uint32(v7) > uint32(v9+i32(39)) {
			m.fn7(i32(1274468), i32(46), i32(1274516))
			panic("unreachable")
		}
	l14:
		m.fn5(v5)
	}
l12:
	m.g0 = v6 + i32(16)
	return v0
}
func (m *Module) fn308(v0, v1, v2, v3, v4, v5 int32) int32 {
	var v6 int32
	if v1 == 0 {
		goto l0
	}
	v1 = v1 * i32(44)
l2:
	{
		t0 := int32(load32(m.memory[uint32(v0):]))
		if t0 == i32(-1) {
			goto l1
		}
		t1 := int32(load32(m.memory[uint32(v0+i32(8)):]))
		if t1 != v5 {
			goto l1
		}
		t2 := int32(load32(m.memory[uint32(v0+i32(4)):]))
		t3 := m.fn1909(t2, v4, v5)
		if t3 != 0 {
			goto l1
		}
		t4 := int32(load32(m.memory[uint32(v0+i32(36)):]))
		v6 = t4
		if v6 == 0 {
			goto l1
		}
		t5 := int32(load32(m.memory[uint32(v0+i32(40)):]))
		if t5 != v3 {
			goto l1
		}
		t6 := m.fn1909(v6+i32(8), v2, v3)
		if t6 != 0 {
			goto l1
		}
		return v0
	}
l1:
	v0 = v0 + i32(44)
	v1 = v1 + i32(-44)
	if v1 != 0 {
		goto l2
	}
l0:
	return i32(0)
}
func (m *Module) fn309(v0, v1, v2 int32) {
	var v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14 int32
	t0 := m.g0
	v3 = t0 - i32(32)
	m.g0 = v3
	v4 = i32(0)
	store32(m.memory[int64(uint32(v3))+16:], uint32(i32(0)))
	store64(m.memory[int64(uint32(v3))+8:], uint64(i64(0x100000000)))
	{
		if v2 != 0 {
			goto l0
		}
		v5 = i32(4)
		v6 = i32(0)
		goto l1
	l0:
		v7 = v2 << 2
		t1 := m.fn11(v7)
		v5 = t1
		if v5 == 0 {
			m.fn16(i32(4), v7)
			panic("unreachable")
		}
		v7 = v2*i32(44) + i32(-44)
		t2 := int32(uint32(v7) / uint32(i32(44)))
		v8 = t2 + i32(1)
		v9 = v8 & i32(7)
		v6 = i32(0)
		if uint32(v7) < uint32(i32(308)) {
			goto l3
		}
		v6 = v8 & i32(0xffffff8)
		v10 = v8 << 2 & i32(0x3fffffe0)
		v11 = i32(0)
	l4:
		{
			v7 = v5 + v11
			store32(m.memory[uint32(v7):], uint32(v1))
			store32(m.memory[uint32(v7+i32(28)):], uint32(v1+i32(308)))
			store32(m.memory[uint32(v7+i32(24)):], uint32(v1+i32(264)))
			store32(m.memory[uint32(v7+i32(20)):], uint32(v1+i32(220)))
			store32(m.memory[uint32(v7+i32(16)):], uint32(v1+i32(176)))
			store32(m.memory[uint32(v7+i32(12)):], uint32(v1+i32(132)))
			store32(m.memory[uint32(v7+i32(8)):], uint32(v1+i32(88)))
			store32(m.memory[uint32(v7+i32(4)):], uint32(v1+i32(44)))
			v1 = v1 + i32(352)
			t3 := v10
			v11 = v11 + i32(32)
			if t3 != v11 {
				goto l4
			}
		}
		if v9 == 0 {
			goto l5
		}
	l3:
		v10 = v6 + v9
		v11 = v9 << 2
		v7 = v5 + v6<<2
	l6:
		store32(m.memory[uint32(v7):], uint32(v1))
		v7 = v7 + i32(4)
		v1 = v1 + i32(44)
		v11 = v11 + i32(-4)
		if v11 != 0 {
			goto l6
		}
		v6 = v10
		if uint32(v10) >= uint32(i32(2)) {
			goto l5
		}
		v6 = i32(1)
		goto l1
	l5:
		v12 = v5 + v6<<2
		v11 = i32(0)
		v1 = int32(uint32(v8) >> 1)
		if v1 == i32(1) {
			goto l7
		}
		v13 = v1 & i32(1)
		v14 = v1 & i32(0x7fffffe)
		v7 = v12 + i32(-4)
		v11 = i32(0)
		v1 = v5
	l8:
		{
			t4 := int32(load32(m.memory[uint32(v7):]))
			v10 = t4
			t5 := int32(load32(m.memory[uint32(v1):]))
			store32(m.memory[uint32(v7):], uint32(t5))
			store32(m.memory[uint32(v1):], uint32(v10))
			v10 = v12 + (v11^i32(0x3ffffffe))<<2
			t6 := int32(load32(m.memory[uint32(v10):]))
			v9 = t6
			t7 := v10
			v8 = v1 + i32(4)
			t8 := int32(load32(m.memory[uint32(v8):]))
			store32(m.memory[uint32(t7):], uint32(t8))
			store32(m.memory[uint32(v8):], uint32(v9))
			v7 = v7 + i32(-8)
			v1 = v1 + i32(8)
			t9 := v14
			v11 = v11 + i32(2)
			if t9 != v11 {
				goto l8
			}
		}
		if v13 == 0 {
			goto l1
		}
	l7:
		v1 = v5 + v11<<2
		t10 := int32(load32(m.memory[uint32(v1):]))
		v7 = t10
		t11 := v1
		v11 = v12 + (v11^i32(-1))<<2
		t12 := int32(load32(m.memory[uint32(v11):]))
		store32(m.memory[uint32(t11):], uint32(t12))
		store32(m.memory[uint32(v11):], uint32(v7))
	}
l1:
	store32(m.memory[int64(uint32(v3))+28:], uint32(v6))
	store32(m.memory[int64(uint32(v3))+24:], uint32(v5))
	store32(m.memory[int64(uint32(v3))+20:], uint32(v2))
	v7 = i32(1)
l10:
	{
		{
			t13 := m.fn151(v3 + i32(20))
			v1 = t13
			if v1 == 0 {
				{
					t20 := int32(load32(m.memory[int64(uint32(v3))+20:]))
					v1 = t20
					if v1 == 0 {
						goto l13
					}
					t21 := int32(load32(m.memory[int64(uint32(v3))+24:]))
					v11 = t21
					t22 := int32(load32(m.memory[uint32(v11+i32(-4)):]))
					v7 = t22
					v5 = v7 & i32(-8)
					t23 := v5
					v7 = v7 & i32(3)
					p24 := i32(8)
					if v7 != 0 {
						p24 = i32(4)
					}
					v1 = v1 << 2
					if uint32(t23) < uint32(p24+v1) {
						m.fn7(i32(1274404), i32(46), i32(1274452))
						panic("unreachable")
					}
					if v7 == 0 {
						goto l15
					}
					if uint32(v5) > uint32(v1+i32(39)) {
						m.fn7(i32(1274468), i32(46), i32(1274516))
						panic("unreachable")
					}
				l15:
					m.fn5(v11)
				}
			l13:
				t25 := int32(load32(m.memory[int64(uint32(v3))+16:]))
				store32(m.memory[int64(uint32(v0))+8:], uint32(t25))
				t26 := int64(load64(m.memory[int64(uint32(v3))+8:]))
				store64(m.memory[uint32(v0):], uint64(t26))
				m.g0 = v3 + i32(32)
				return
			}
			t14 := int32(load32(m.memory[uint32(v1):]))
			if t14 != i32(-1) {
				goto l10
			}
			t15 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			v11 = t15
			t16 := int32(load32(m.memory[int64(uint32(v1))+12:]))
			v1 = t16
			t17 := int32(load32(m.memory[int64(uint32(v3))+8:]))
			if uint32(v1) <= uint32(t17-v4) {
				goto l11
			}
			m.fn197(v3+i32(8), v4, v1, i32(1), i32(1))
			t18 := int32(load32(m.memory[int64(uint32(v3))+12:]))
			v7 = t18
			t19 := int32(load32(m.memory[int64(uint32(v3))+16:]))
			v4 = t19
			goto l12
		}
	l11:
		if v1 == 0 {
			goto l17
		}
	l12:
		if v1 == 0 {
			goto l17
		}
		memory_copy(m.memory, uint32(v7+v4), uint32(v11), uint32(v1))
	l17:
		t27 := v3
		v4 = v4 + v1
		store32(m.memory[int64(uint32(t27))+16:], uint32(v4))
		goto l10
	}
}
func (m *Module) fn310(v0 int32) {
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
	m.fn208(t2, t4, t3, v2, i32(8), i32(32))
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
func (m *Module) fn311(v0 int32) {
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
	m.fn208(t2, t4, t3, v2, i32(4), i32(12))
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
