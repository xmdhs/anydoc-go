package core

import (
	"math/bits"
)

func (m *Module) fn897(v0, v1, v2, v3, v4, v5, v6, v7 int32) {
	var v8, v9, v10 int32
	t0 := m.g0
	v8 = t0 - i32(16)
	m.g0 = v8
	m.memory[int64(uint32(v1))+24] = byte(i32(9))
	{
		{
			if v6 != 0 {
				m.fn898(v0, v1, v2, v3, v4, v5, i32(1))
				t4 := int32(m.memory[int64(uint32(v0))+4])
				if t4 != 0 {
					goto l4
				}
				m.memory[int64(uint32(v1))+24] = byte(i32(10))
				goto l4
			}
			m.memory[int64(uint32(v8))+3] = byte(v7)
			v7 = i32(0)
			m.fn898(v8+i32(4), v1, v8+i32(3), i32(1), v4, v5, i32(0))
			t1 := int32(load32(m.memory[int64(uint32(v8))+12:]))
			v6 = t1
			t2 := int32(load16(m.memory[int64(uint32(v8))+9:]))
			v9 = t2
			v10 = i32(2)
			t3 := int32(m.memory[int64(uint32(v8))+8])
			switch t3 {
			case 1:
				m.fn7(i32(1146620), i32(39), i32(1146572))
				panic("unreachable")
			case 2:
				goto l3
			default:
				goto l1
			}
		}
	l1:
		if uint32(v5) < uint32(v6) {
			m.fn121(v6, v5, v5, i32(1146572))
			panic("unreachable")
		}
		m.fn898(v8+i32(4), v1, v2, v3, v4+v6, v5-v6, i32(1))
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
func (m *Module) fn898(v0, v1, v2, v3, v4, v5, v6 int32) {
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
			m.fn121(v12, v5, v5, i32(1146724))
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
						m.fn121(v11, v3, v3, i32(1140168))
						panic("unreachable")
					}
					{
						if uint32(v5) < uint32(v12) {
							m.fn121(v12, v5, v5, i32(1140152))
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
															m.fn33(v26, v9, i32(1140312))
															panic("unreachable")
														}
														v27 = i32(2)
														v16 = v15
														goto l80
													}
													m.fn33(v13, v9, i32(1140296))
													panic("unreachable")
												}
												if uint32(v13) >= uint32(v9) {
													m.fn33(v13, v9, i32(1140248))
													panic("unreachable")
												}
												m.memory[uint32(v14+v13)] = byte(int32(uint32(v15&i32(61440))>>12) | i32(224))
												v16 = v13 + i32(1)
												if uint32(v16) >= uint32(v9) {
													m.fn33(v16, v9, i32(1140264))
													panic("unreachable")
												}
												m.memory[uint32(v14+v16)] = byte(int32(uint32(v15)>>6)&i32(63) | i32(128))
												v26 = v13 + i32(2)
												if uint32(v26) >= uint32(v9) {
													m.fn33(v26, v9, i32(1140280))
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
												m.fn33(v13, v9, i32(1140184))
												panic("unreachable")
											}
											t49 := v14 + v13
											v15 = v15&i32(0xffff)<<10 + v16&i32(0xffff) + i32(-56613888)
											m.memory[uint32(t49)] = byte(int32(uint32(v15)>>18) | i32(240))
											v26 = v13 + i32(1)
											if uint32(v26) >= uint32(v9) {
												m.fn33(v26, v9, i32(1140200))
												panic("unreachable")
											}
											m.memory[uint32(v14+v26)] = byte(int32(uint32(v15)>>12)&i32(63) | i32(128))
											v26 = v13 + i32(2)
											if uint32(v26) >= uint32(v9) {
												m.fn33(v26, v9, i32(1140216))
												panic("unreachable")
											}
											m.memory[uint32(v14+v26)] = byte(int32(uint32(v15)>>6)&i32(63) | i32(128))
											v26 = v13 + i32(3)
											if uint32(v26) >= uint32(v9) {
												m.fn33(v26, v9, i32(1140232))
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
											m.fn7(i32(1140104), i32(30), i32(1140136))
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
					m.fn121(v11, v3, v3, i32(1140168))
					panic("unreachable")
				}
				{
					if uint32(v5) < uint32(v12) {
						m.fn121(v12, v5, v5, i32(1140152))
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
														m.fn33(v27, v9, i32(1140312))
														panic("unreachable")
													}
													v25 = i32(2)
													goto l63
												}
												m.fn33(v13, v9, i32(1140296))
												panic("unreachable")
											}
											if uint32(v13) >= uint32(v9) {
												m.fn33(v13, v9, i32(1140248))
												panic("unreachable")
											}
											m.memory[uint32(v14+v13)] = byte(int32(uint32(v15&i32(61440))>>12) | i32(224))
											v16 = v13 + i32(1)
											if uint32(v16) >= uint32(v9) {
												m.fn33(v16, v9, i32(1140264))
												panic("unreachable")
											}
											m.memory[uint32(v14+v16)] = byte(int32(uint32(v15)>>6)&i32(63) | i32(128))
											v27 = v13 + i32(2)
											if uint32(v27) >= uint32(v9) {
												m.fn33(v27, v9, i32(1140280))
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
											m.fn33(v13, v9, i32(1140184))
											panic("unreachable")
										}
										t38 := v14 + v13
										v15 = v15&i32(0xffff)<<10 + v26 + i32(-56613888)
										m.memory[uint32(t38)] = byte(int32(uint32(v15)>>18) | i32(240))
										v16 = v13 + i32(1)
										if uint32(v16) >= uint32(v9) {
											m.fn33(v16, v9, i32(1140200))
											panic("unreachable")
										}
										m.memory[uint32(v14+v16)] = byte(int32(uint32(v15)>>12)&i32(63) | i32(128))
										v16 = v13 + i32(2)
										if uint32(v16) >= uint32(v9) {
											m.fn33(v16, v9, i32(1140216))
											panic("unreachable")
										}
										m.memory[uint32(v14+v16)] = byte(int32(uint32(v15)>>6)&i32(63) | i32(128))
										v27 = v13 + i32(3)
										if uint32(v27) >= uint32(v9) {
											m.fn33(v27, v9, i32(1140232))
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
										m.fn7(i32(1140104), i32(30), i32(1140136))
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
								t76 := int32(load16(m.memory[int64(uint32(v12<<1))+1219988:]))
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
									t62 := int32(load16(m.memory[int64(uint32(v15<<1))+1227524:]))
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
								t63 := int32(load16(m.memory[int64(uint32(v12<<1))+1210024:]))
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
									t65 := int32(load16(m.memory[int64(uint32(v11<<1))+1235360:]))
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
								t67 := int32(load16(m.memory[int64(uint32(v11<<1))+1235328:]))
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
								t69 := int32(load16(m.memory[int64(uint32(v11<<1))+1147006:]))
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
								m.fn900(v7+i32(32), i32(1242550), i32(77), v13)
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
											m.fn33(v10, i32(77), i32(1227872))
											panic("unreachable")
										}
										v10 = v10 << 1
										t73 := int32(load16(m.memory[int64(uint32(v10))+1263672:]))
										t74 := int32(load16(m.memory[int64(uint32(v10))+1242550:]))
										v13 = t73 + v13 - t74
										goto l147
									}
								l145:
									if uint32(v10) > uint32(i32(76)) {
										m.fn33(v10, i32(77), i32(1227856))
										panic("unreachable")
									}
									t75 := int32(load16(m.memory[int64(uint32(v10<<1))+1263672:]))
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
					m.fn900(v7+i32(40), i32(1252008), i32(535), v10)
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
								m.fn33(i32(-1), i32(535), i32(1227872))
								panic("unreachable")
							}
							t80 := v10
							v13 = v13 << 1
							t81 := int32(load16(m.memory[int64(uint32(v13))+1252008:]))
							t82 := int32(load16(m.memory[int64(uint32(v13))+1244948:]))
							v13 = t80 - t81 + t82
							v15 = v13 & i32(0xffff)
							v10 = int32(uint32(v15) >> 12)
							v15 = int32(uint32(v15) >> 6)
							goto l157
						}
					l155:
						if uint32(v13) > uint32(i32(534)) {
							m.fn33(i32(535), i32(535), i32(1227856))
							panic("unreachable")
						}
						t83 := int32(load16(m.memory[int64(uint32(v13<<1))+1244948:]))
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
						t86 := int32(load16(m.memory[int64(uint32(v10<<1))+1246018:]))
						p87 := v10
						if uint32(t86) > uint32(v13) {
							p87 = t85
						}
						v10 = p87
						t88 := v10
						v10 = v10 + i32(135)
						t89 := int32(load16(m.memory[int64(uint32(v10<<1))+1246018:]))
						p90 := v10
						if uint32(t89) > uint32(v13) {
							p90 = t88
						}
						v10 = p90
						t91 := v10
						v10 = v10 + i32(67)
						t92 := int32(load16(m.memory[int64(uint32(v10<<1))+1246018:]))
						p93 := v10
						if uint32(t92) > uint32(v13) {
							p93 = t91
						}
						v10 = p93
						t94 := v10
						v10 = v10 + i32(34)
						t95 := int32(load16(m.memory[int64(uint32(v10<<1))+1246018:]))
						p96 := v10
						if uint32(t95) > uint32(v13) {
							p96 = t94
						}
						v10 = p96
						t97 := v10
						v10 = v10 + i32(17)
						t98 := int32(load16(m.memory[int64(uint32(v10<<1))+1246018:]))
						p99 := v10
						if uint32(t98) > uint32(v13) {
							p99 = t97
						}
						v10 = p99
						t100 := v10
						v10 = v10 + i32(8)
						t101 := int32(load16(m.memory[int64(uint32(v10<<1))+1246018:]))
						p102 := v10
						if uint32(t101) > uint32(v13) {
							p102 = t100
						}
						v10 = p102
						t103 := v10
						v10 = v10 + i32(4)
						t104 := int32(load16(m.memory[int64(uint32(v10<<1))+1246018:]))
						p105 := v10
						if uint32(t104) > uint32(v13) {
							p105 = t103
						}
						v10 = p105
						t106 := v10
						v10 = v10 + i32(2)
						t107 := int32(load16(m.memory[int64(uint32(v10<<1))+1246018:]))
						p108 := v10
						if uint32(t107) > uint32(v13) {
							p108 = t106
						}
						v10 = p108
						t109 := v10
						v10 = v10 + i32(1)
						t110 := int32(load16(m.memory[int64(uint32(v10<<1))+1246018:]))
						p111 := v10
						if uint32(t110) > uint32(v13) {
							p111 = t109
						}
						v10 = p111
						t112 := v10
						v10 = v10 + i32(1)
						t113 := int32(load16(m.memory[int64(uint32(v10<<1))+1246018:]))
						p114 := v10
						if uint32(t113) > uint32(v13) {
							p114 = t112
						}
						v10 = p114
						v15 = v10 << 1
						t115 := int32(load16(m.memory[int64(uint32(v15))+1246018:]))
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
								m.fn33(i32(-1), i32(1079), i32(1227872))
								panic("unreachable")
							}
							t118 := v12
							v13 = v13 << 1
							t119 := int32(load16(m.memory[int64(uint32(v13))+1246018:]))
							t120 := int32(load16(m.memory[int64(uint32(v13))+1242724:]))
							v13 = t118 - t119 + t120
							v15 = v13 & i32(0xffff)
							v10 = int32(uint32(v15) >> 12)
							v15 = int32(uint32(v15) >> 6)
							goto l161
						}
					}
				l159:
					t121 := int32(load16(m.memory[int64(uint32(v15))+1242724:]))
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
									t167 := int32(load16(m.memory[int64(uint32(v17<<1))+1219988:]))
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
										t140 := int32(load16(m.memory[int64(uint32(v10<<1))+1227524:]))
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
									t141 := int32(load16(m.memory[int64(uint32(v17<<1))+1210024:]))
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
										t146 := int32(load16(m.memory[int64(uint32(v11<<1))+1235360:]))
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
									t151 := int32(load16(m.memory[int64(uint32(v11<<1))+1235328:]))
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
									t156 := int32(load16(m.memory[int64(uint32(v11<<1))+1147006:]))
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
									m.fn900(v7+i32(24), i32(1242550), i32(77), v13)
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
												m.fn33(v10, i32(77), i32(1227872))
												panic("unreachable")
											}
											v10 = v10 << 1
											t164 := int32(load16(m.memory[int64(uint32(v10))+1263672:]))
											t165 := int32(load16(m.memory[int64(uint32(v10))+1242550:]))
											v13 = t164 + v13 - t165
											goto l206
										}
									l204:
										if uint32(v10) > uint32(i32(76)) {
											m.fn33(v10, i32(77), i32(1227856))
											panic("unreachable")
										}
										t166 := int32(load16(m.memory[int64(uint32(v10<<1))+1263672:]))
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
									t174 := int32(load16(m.memory[int64(uint32(v13<<1))+1252008:]))
									p175 := v13
									if uint32(t174) > uint32(v10) {
										p175 = t173
									}
									v13 = p175
									t176 := v13
									v13 = v13 + i32(67)
									t177 := int32(load16(m.memory[int64(uint32(v13<<1))+1252008:]))
									p178 := v13
									if uint32(t177) > uint32(v10) {
										p178 = t176
									}
									v13 = p178
									t179 := v13
									v13 = v13 + i32(33)
									t180 := int32(load16(m.memory[int64(uint32(v13<<1))+1252008:]))
									p181 := v13
									if uint32(t180) > uint32(v10) {
										p181 = t179
									}
									v13 = p181
									t182 := v13
									v13 = v13 + i32(17)
									t183 := int32(load16(m.memory[int64(uint32(v13<<1))+1252008:]))
									p184 := v13
									if uint32(t183) > uint32(v10) {
										p184 = t182
									}
									v13 = p184
									t185 := v13
									v13 = v13 + i32(8)
									t186 := int32(load16(m.memory[int64(uint32(v13<<1))+1252008:]))
									p187 := v13
									if uint32(t186) > uint32(v10) {
										p187 = t185
									}
									v13 = p187
									t188 := v13
									v13 = v13 + i32(4)
									t189 := int32(load16(m.memory[int64(uint32(v13<<1))+1252008:]))
									p190 := v13
									if uint32(t189) > uint32(v10) {
										p190 = t188
									}
									v13 = p190
									t191 := v13
									v13 = v13 + i32(2)
									t192 := int32(load16(m.memory[int64(uint32(v13<<1))+1252008:]))
									p193 := v13
									if uint32(t192) > uint32(v10) {
										p193 = t191
									}
									v13 = p193
									t194 := v13
									v13 = v13 + i32(1)
									t195 := int32(load16(m.memory[int64(uint32(v13<<1))+1252008:]))
									p196 := v13
									if uint32(t195) > uint32(v10) {
										p196 = t194
									}
									v13 = p196
									t197 := v13
									v13 = v13 + i32(1)
									t198 := int32(load16(m.memory[int64(uint32(v13<<1))+1252008:]))
									p199 := v13
									if uint32(t198) > uint32(v10) {
										p199 = t197
									}
									v13 = p199
									v12 = v13 << 1
									t200 := int32(load16(m.memory[int64(uint32(v12))+1252008:]))
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
											m.fn33(i32(-1), i32(535), i32(1227872))
											panic("unreachable")
										}
										t203 := v10
										v13 = v13 << 1
										t204 := int32(load16(m.memory[int64(uint32(v13))+1252008:]))
										t205 := int32(load16(m.memory[int64(uint32(v13))+1244948:]))
										v10 = t203 - t204 + t205
										v12 = v10 & i32(0xffff)
										v13 = int32(uint32(v12) >> 12)
										v12 = int32(uint32(v12) >> 6)
										goto l214
									}
								}
							l212:
								t206 := int32(load16(m.memory[int64(uint32(v12))+1244948:]))
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
								t213 := int32(load16(m.memory[int64(uint32(v10<<1))+1246018:]))
								p214 := v10
								if uint32(t213) > uint32(v13) {
									p214 = t212
								}
								v10 = p214
								t215 := v10
								v10 = v10 + i32(135)
								t216 := int32(load16(m.memory[int64(uint32(v10<<1))+1246018:]))
								p217 := v10
								if uint32(t216) > uint32(v13) {
									p217 = t215
								}
								v10 = p217
								t218 := v10
								v10 = v10 + i32(67)
								t219 := int32(load16(m.memory[int64(uint32(v10<<1))+1246018:]))
								p220 := v10
								if uint32(t219) > uint32(v13) {
									p220 = t218
								}
								v10 = p220
								t221 := v10
								v10 = v10 + i32(34)
								t222 := int32(load16(m.memory[int64(uint32(v10<<1))+1246018:]))
								p223 := v10
								if uint32(t222) > uint32(v13) {
									p223 = t221
								}
								v10 = p223
								t224 := v10
								v10 = v10 + i32(17)
								t225 := int32(load16(m.memory[int64(uint32(v10<<1))+1246018:]))
								p226 := v10
								if uint32(t225) > uint32(v13) {
									p226 = t224
								}
								v10 = p226
								t227 := v10
								v10 = v10 + i32(8)
								t228 := int32(load16(m.memory[int64(uint32(v10<<1))+1246018:]))
								p229 := v10
								if uint32(t228) > uint32(v13) {
									p229 = t227
								}
								v10 = p229
								t230 := v10
								v10 = v10 + i32(4)
								t231 := int32(load16(m.memory[int64(uint32(v10<<1))+1246018:]))
								p232 := v10
								if uint32(t231) > uint32(v13) {
									p232 = t230
								}
								v10 = p232
								t233 := v10
								v10 = v10 + i32(2)
								t234 := int32(load16(m.memory[int64(uint32(v10<<1))+1246018:]))
								p235 := v10
								if uint32(t234) > uint32(v13) {
									p235 = t233
								}
								v10 = p235
								t236 := v10
								v10 = v10 + i32(1)
								t237 := int32(load16(m.memory[int64(uint32(v10<<1))+1246018:]))
								p238 := v10
								if uint32(t237) > uint32(v13) {
									p238 = t236
								}
								v10 = p238
								t239 := v10
								v10 = v10 + i32(1)
								t240 := int32(load16(m.memory[int64(uint32(v10<<1))+1246018:]))
								p241 := v10
								if uint32(t240) > uint32(v13) {
									p241 = t239
								}
								v10 = p241
								v12 = v10 << 1
								t242 := int32(load16(m.memory[int64(uint32(v12))+1246018:]))
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
										m.fn33(i32(-1), i32(1079), i32(1227872))
										panic("unreachable")
									}
									t245 := v17
									v13 = v13 << 1
									t246 := int32(load16(m.memory[int64(uint32(v13))+1246018:]))
									t247 := int32(load16(m.memory[int64(uint32(v13))+1242724:]))
									v10 = t245 - t246 + t247
									v12 = v10 & i32(0xffff)
									v13 = int32(uint32(v12) >> 12)
									v12 = int32(uint32(v12) >> 6)
									goto l217
								}
							}
						l215:
							t248 := int32(load16(m.memory[int64(uint32(v12))+1242724:]))
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
			m.fn121(v15, v5, v5, i32(1146724))
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
								t273 := int32(load16(m.memory[int64(uint32(v16<<1))+1235510:]))
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
								t275 := int32(load16(m.memory[int64(uint32(v16<<1))+1263826:]))
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
								t277 := int32(load16(m.memory[int64(uint32(v16<<1))+1270606:]))
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
							m.fn901(v7+i32(16), v10)
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
						t271 := int32(load16(m.memory[int64(uint32(v12<<1))+1241946:]))
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
					t279 := int32(load16(m.memory[int64(uint32(v16<<1))+1270606:]))
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
							m.fn121(v17, v5, v5, i32(1146724))
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
							t321 := int32(load16(m.memory[int64(uint32(v11<<1))+1235510:]))
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
							t326 := int32(load16(m.memory[int64(uint32(v11<<1))+1263826:]))
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
							t331 := int32(load16(m.memory[int64(uint32(v11<<1))+1270606:]))
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
							t336 := int32(load16(m.memory[int64(uint32(v11<<1))+1270606:]))
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
							t308 := int32(load16(m.memory[int64(uint32(v17<<1))+1241946:]))
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
						t342 := int32(load16(m.memory[int64(uint32(v17<<1))+1242484:]))
						v5 = v11 + t342
						if uint32(v5) < uint32(i32(240)) {
							goto l338
						}
						m.fn33(v5, i32(240), i32(1242056))
						panic("unreachable")
					}
				l338:
					v11 = v16 + i32(1)
					v13 = v4 + v16
					{
						t343 := int32(load16(m.memory[int64(uint32(v5<<1))+1227044:]))
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
			m.fn121(v15, v3, v3, i32(1146740))
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
						m.fn7(i32(1274576), i32(40), i32(1146400))
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
									t366 := int32(load16(m.memory[int64(uint32(v15<<1))+1235510:]))
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
									t368 := int32(load16(m.memory[int64(uint32(v15<<1))+1263826:]))
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
									t359 := int32(load16(m.memory[int64(uint32(v15<<1))+1270606:]))
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
									t361 := int32(load16(m.memory[int64(uint32(v12<<1))+1241946:]))
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
							t370 := int32(load16(m.memory[int64(uint32(v12<<1))+1242484:]))
							v15 = v15 + t370
							if uint32(v15) < uint32(i32(240)) {
								goto l420
							}
							m.fn33(v15, i32(240), i32(1242056))
							panic("unreachable")
						}
					l420:
						v10 = v4 + v16
						t371 := int32(load16(m.memory[int64(uint32(v15<<1))+1227044:]))
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
								t389 := int32(load16(m.memory[int64(uint32(v10<<1))+1235510:]))
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
								t391 := int32(load16(m.memory[int64(uint32(v10<<1))+1263826:]))
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
								t381 := int32(load16(m.memory[int64(uint32(v10<<1))+1270606:]))
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
								t383 := int32(load16(m.memory[int64(uint32(v16<<1))+1241946:]))
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
								t393 := int32(load16(m.memory[int64(uint32(v10<<1))+1198422:]))
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
					t384 := int32(load16(m.memory[int64(uint32(v16<<1))+1244882:]))
					v10 = v10 + t384
					if uint32(v10) >= uint32(i32(255)) {
						m.fn33(v10, i32(255), i32(1242708))
						panic("unreachable")
					}
					t385 := int32(load16(m.memory[int64(uint32(v10<<1))+1227888:]))
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
					t395 := int32(load16(m.memory[int64(uint32(v15<<1))+1242484:]))
					v13 = v10 + t395
					if uint32(v13) < uint32(i32(240)) {
						goto l498
					}
					m.fn33(v13, i32(240), i32(1242056))
					panic("unreachable")
				}
			l498:
				{
					t396 := int32(load16(m.memory[int64(uint32(v13<<1))+1227044:]))
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
								t454 := int32(load16(m.memory[int64(uint32(v10<<1))+1235510:]))
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
								t460 := int32(load16(m.memory[int64(uint32(v10<<1))+1263826:]))
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
								t434 := int32(load16(m.memory[int64(uint32(v10<<1))+1270606:]))
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
								t440 := int32(load16(m.memory[int64(uint32(v12<<1))+1241946:]))
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
											t427 := int32(load16(m.memory[int64(uint32(v10<<1))+1198422:]))
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
											t416 := int32(load16(m.memory[int64(uint32(v15<<1))+1244882:]))
											v10 = v10 + t416
											if uint32(v10) >= uint32(i32(255)) {
												m.fn33(v10, i32(255), i32(1242708))
												panic("unreachable")
											}
											t417 := int32(load16(m.memory[int64(uint32(v10<<1))+1227888:]))
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
						t466 := int32(load16(m.memory[int64(uint32(v15<<1))+1242484:]))
						v13 = v10 + t466
						if uint32(v13) < uint32(i32(240)) {
							goto l591
						}
						m.fn33(v13, i32(240), i32(1242056))
						panic("unreachable")
					}
				l591:
					v15 = v16 + i32(1)
					v10 = v4 + v16
					{
						t467 := int32(load16(m.memory[int64(uint32(v13<<1))+1227044:]))
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
			m.fn121(v17, v5, v5, i32(1146724))
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
					t477 := int32(load16(m.memory[int64(uint32(v10<<1))+1160678:]))
					v16 = t477
					if v16 != 0 {
						t478 := int32(load32(m.memory[int64(uint32(int32(uint32(v10)>>3)&i32(0x1ffffffc)))+1224688:]))
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
					m.fn121(v17, v5, v5, i32(1146756))
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
						t497 := int32(load16(m.memory[int64(uint32(v10<<1))+1160678:]))
						v18 = t497
						if v18 != 0 {
							v16 = v13 + i32(1)
							{
								t498 := int32(load32(m.memory[int64(uint32(int32(uint32(v10)>>3)&i32(0x1ffffffc)))+1224688:]))
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
										t620 := int32(load16(m.memory[int64(uint32(v16<<1))+1242072:]))
										t621 := v16
										v16 = v10 & i32(0xffff)
										p622 := t621
										if uint32(t620) > uint32(v16) {
											p622 = t619
										}
										v12 = p622
										t623 := v12
										v12 = v12 + i32(26)
										t624 := int32(load16(m.memory[int64(uint32(v12<<1))+1242072:]))
										p625 := v12
										if uint32(t624) > uint32(v16) {
											p625 = t623
										}
										v12 = p625
										t626 := v12
										v12 = v12 + i32(13)
										t627 := int32(load16(m.memory[int64(uint32(v12<<1))+1242072:]))
										p628 := v12
										if uint32(t627) > uint32(v16) {
											p628 = t626
										}
										v12 = p628
										t629 := v12
										v12 = v12 + i32(6)
										t630 := int32(load16(m.memory[int64(uint32(v12<<1))+1242072:]))
										p631 := v12
										if uint32(t630) > uint32(v16) {
											p631 = t629
										}
										v12 = p631
										t632 := v12
										v12 = v12 + i32(3)
										t633 := int32(load16(m.memory[int64(uint32(v12<<1))+1242072:]))
										p634 := v12
										if uint32(t633) > uint32(v16) {
											p634 = t632
										}
										v12 = p634
										t635 := v12
										v12 = v12 + i32(2)
										t636 := int32(load16(m.memory[int64(uint32(v12<<1))+1242072:]))
										p637 := v12
										if uint32(t636) > uint32(v16) {
											p637 = t635
										}
										v12 = p637
										t638 := v12
										v12 = v12 + i32(1)
										t639 := int32(load16(m.memory[int64(uint32(v12<<1))+1242072:]))
										p640 := v12
										if uint32(t639) > uint32(v16) {
											p640 = t638
										}
										v12 = p640
										v11 = v12 << 1
										t641 := int32(load16(m.memory[int64(uint32(v11))+1242072:]))
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
												m.fn33(i32(-1), i32(206), i32(1227872))
												panic("unreachable")
											}
											v16 = v16 << 1
											t644 := int32(load16(m.memory[int64(uint32(v16))+1241440:]))
											t645 := int32(load16(m.memory[int64(uint32(v16))+1242072:]))
											v10 = t644 + v10 - t645
											goto l703
										}
									}
								l701:
									t646 := int32(load16(m.memory[int64(uint32(v11))+1241440:]))
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
											t614 := int32(load16(m.memory[int64(uint32(v16&i32(255)*i32(188)+v10&i32(255)<<1))+1147142:]))
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
												t609 := int32(load16(m.memory[int64(uint32(v10&i32(255)<<1))+1219800:]))
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
											t611 := int32(load16(m.memory[int64(uint32(v16<<1))+1253078:]))
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
													t549 := int32(load16(m.memory[int64(uint32(v10<<1))+1260326:]))
													t550 := v10
													v10 = v16 & i32(0xffff)
													p551 := t550
													if uint32(t549) > uint32(v10) {
														p551 = t548
													}
													v12 = p551
													t552 := v12
													v12 = v12 + i32(203)
													t553 := int32(load16(m.memory[int64(uint32(v12<<1))+1260326:]))
													p554 := v12
													if uint32(t553) > uint32(v10) {
														p554 = t552
													}
													v12 = p554
													t555 := v12
													v12 = v12 + i32(102)
													t556 := int32(load16(m.memory[int64(uint32(v12<<1))+1260326:]))
													p557 := v12
													if uint32(t556) > uint32(v10) {
														p557 = t555
													}
													v12 = p557
													t558 := v12
													v12 = v12 + i32(51)
													t559 := int32(load16(m.memory[int64(uint32(v12<<1))+1260326:]))
													p560 := v12
													if uint32(t559) > uint32(v10) {
														p560 = t558
													}
													v12 = p560
													t561 := v12
													v12 = v12 + i32(25)
													t562 := int32(load16(m.memory[int64(uint32(v12<<1))+1260326:]))
													p563 := v12
													if uint32(t562) > uint32(v10) {
														p563 = t561
													}
													v12 = p563
													t564 := v12
													v12 = v12 + i32(13)
													t565 := int32(load16(m.memory[int64(uint32(v12<<1))+1260326:]))
													p566 := v12
													if uint32(t565) > uint32(v10) {
														p566 = t564
													}
													v12 = p566
													t567 := v12
													v12 = v12 + i32(6)
													t568 := int32(load16(m.memory[int64(uint32(v12<<1))+1260326:]))
													p569 := v12
													if uint32(t568) > uint32(v10) {
														p569 = t567
													}
													v12 = p569
													t570 := v12
													v12 = v12 + i32(3)
													t571 := int32(load16(m.memory[int64(uint32(v12<<1))+1260326:]))
													p572 := v12
													if uint32(t571) > uint32(v10) {
														p572 = t570
													}
													v12 = p572
													t573 := v12
													v12 = v12 + i32(2)
													t574 := int32(load16(m.memory[int64(uint32(v12<<1))+1260326:]))
													p575 := v12
													if uint32(t574) > uint32(v10) {
														p575 = t573
													}
													v12 = p575
													t576 := v12
													v12 = v12 + i32(1)
													t577 := int32(load16(m.memory[int64(uint32(v12<<1))+1260326:]))
													p578 := v12
													if uint32(t577) > uint32(v10) {
														p578 = t576
													}
													v12 = p578
													v11 = v12 << 1
													t579 := int32(load16(m.memory[int64(uint32(v11))+1260326:]))
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
															m.fn33(i32(-1), i32(1627), i32(1227872))
															panic("unreachable")
														}
														t582 := v16
														v10 = v10 << 1
														t583 := int32(load16(m.memory[int64(uint32(v10))+1260326:]))
														t584 := int32(load16(m.memory[int64(uint32(v10))+1253122:]))
														v10 = t582 - t583 + t584
														v12 = v10 & i32(0xffff)
														v16 = int32(uint32(v12) >> 12)
														v12 = int32(uint32(v12) >> 6)
														goto l683
													}
												}
											l681:
												t585 := int32(load16(m.memory[int64(uint32(v11))+1253122:]))
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
													t588 := int32(load16(m.memory[int64(uint32(v16<<1))+1235390:]))
													p589 := v16
													if uint32(t588) > uint32(v10) {
														p589 = t587
													}
													v16 = p589
													t590 := v16
													v16 = v16 + i32(7)
													t591 := int32(load16(m.memory[int64(uint32(v16<<1))+1235390:]))
													p592 := v16
													if uint32(t591) > uint32(v10) {
														p592 = t590
													}
													v16 = p592
													t593 := v16
													v16 = v16 + i32(4)
													t594 := int32(load16(m.memory[int64(uint32(v16<<1))+1235390:]))
													p595 := v16
													if uint32(t594) > uint32(v10) {
														p595 = t593
													}
													v16 = p595
													t596 := v16
													v16 = v16 + i32(2)
													t597 := int32(load16(m.memory[int64(uint32(v16<<1))+1235390:]))
													p598 := v16
													if uint32(t597) > uint32(v10) {
														p598 = t596
													}
													v16 = p598
													t599 := v16
													v16 = v16 + i32(1)
													t600 := int32(load16(m.memory[int64(uint32(v16<<1))+1235390:]))
													p601 := v16
													if uint32(t600) > uint32(v10) {
														p601 = t599
													}
													v16 = p601
													v12 = v16 << 1
													t602 := int32(load16(m.memory[int64(uint32(v12))+1235390:]))
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
															m.fn33(i32(-1), i32(59), i32(1227872))
															panic("unreachable")
														}
														v16 = v16 << 1
														t605 := int32(load16(m.memory[int64(uint32(v16))+1256376:]))
														t606 := int32(load16(m.memory[int64(uint32(v16))+1235390:]))
														v10 = t605 + v10 - t606
														goto l686
													}
												}
											l684:
												t607 := int32(load16(m.memory[int64(uint32(v12))+1256376:]))
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
											m.fn33(v10, i32(101), i32(1146788))
											panic("unreachable")
										}
										v13 = v4 + v13
										t545 := int32(load16(m.memory[int64(uint32(v10<<1))+1146804:]))
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
										t509 := int32(load16(m.memory[int64(uint32(v16<<1))+1256494:]))
										p510 := v16
										if uint32(t509) > uint32(v10) {
											p510 = t508
										}
										v16 = p510
										t511 := v16
										v16 = v16 + i32(239)
										t512 := int32(load16(m.memory[int64(uint32(v16<<1))+1256494:]))
										p513 := v16
										if uint32(t512) > uint32(v10) {
											p513 = t511
										}
										v16 = p513
										t514 := v16
										v16 = v16 + i32(120)
										t515 := int32(load16(m.memory[int64(uint32(v16<<1))+1256494:]))
										p516 := v16
										if uint32(t515) > uint32(v10) {
											p516 = t514
										}
										v16 = p516
										t517 := v16
										v16 = v16 + i32(60)
										t518 := int32(load16(m.memory[int64(uint32(v16<<1))+1256494:]))
										p519 := v16
										if uint32(t518) > uint32(v10) {
											p519 = t517
										}
										v16 = p519
										t520 := v16
										v16 = v16 + i32(30)
										t521 := int32(load16(m.memory[int64(uint32(v16<<1))+1256494:]))
										p522 := v16
										if uint32(t521) > uint32(v10) {
											p522 = t520
										}
										v16 = p522
										t523 := v16
										v16 = v16 + i32(15)
										t524 := int32(load16(m.memory[int64(uint32(v16<<1))+1256494:]))
										p525 := v16
										if uint32(t524) > uint32(v10) {
											p525 = t523
										}
										v16 = p525
										t526 := v16
										v16 = v16 + i32(7)
										t527 := int32(load16(m.memory[int64(uint32(v16<<1))+1256494:]))
										p528 := v16
										if uint32(t527) > uint32(v10) {
											p528 = t526
										}
										v16 = p528
										t529 := v16
										v16 = v16 + i32(4)
										t530 := int32(load16(m.memory[int64(uint32(v16<<1))+1256494:]))
										p531 := v16
										if uint32(t530) > uint32(v10) {
											p531 = t529
										}
										v16 = p531
										t532 := v16
										v16 = v16 + i32(2)
										t533 := int32(load16(m.memory[int64(uint32(v16<<1))+1256494:]))
										p534 := v16
										if uint32(t533) > uint32(v10) {
											p534 = t532
										}
										v16 = p534
										t535 := v16
										v16 = v16 + i32(1)
										t536 := int32(load16(m.memory[int64(uint32(v16<<1))+1256494:]))
										p537 := v16
										if uint32(t536) > uint32(v10) {
											p537 = t535
										}
										v16 = p537
										v12 = v16 << 1
										t538 := int32(load16(m.memory[int64(uint32(v12))+1256494:]))
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
												m.fn33(i32(-1), i32(1916), i32(1227872))
												panic("unreachable")
											}
											t541 := v18
											v10 = v10 << 1
											t542 := int32(load16(m.memory[int64(uint32(v10))+1256494:]))
											t543 := int32(load16(m.memory[int64(uint32(v10))+1248176:]))
											v10 = t541 - t542 + t543
											v12 = v10 & i32(0xffff)
											v16 = int32(uint32(v12) >> 12)
											v12 = int32(uint32(v12) >> 6)
											goto l673
										}
									}
								l671:
									t544 := int32(load16(m.memory[int64(uint32(v12))+1248176:]))
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
						t649 := int32(load16(m.memory[int64(uint32(v16<<1))+1241852:]))
						p650 := v16
						if uint32(t649) > uint32(v10) {
							p650 = t648
						}
						v16 = p650
						t651 := v16
						v16 = v16 + i32(6)
						t652 := int32(load16(m.memory[int64(uint32(v16<<1))+1241852:]))
						p653 := v16
						if uint32(t652) > uint32(v10) {
							p653 = t651
						}
						v16 = p653
						t654 := v16
						v16 = v16 + i32(3)
						t655 := int32(load16(m.memory[int64(uint32(v16<<1))+1241852:]))
						p656 := v16
						if uint32(t655) > uint32(v10) {
							p656 = t654
						}
						v16 = p656
						t657 := v16
						v16 = v16 + i32(1)
						t658 := int32(load16(m.memory[int64(uint32(v16<<1))+1241852:]))
						p659 := v16
						if uint32(t658) > uint32(v10) {
							p659 = t657
						}
						v16 = p659
						t660 := v16
						v16 = v16 + i32(1)
						t661 := int32(load16(m.memory[int64(uint32(v16<<1))+1241852:]))
						p662 := v16
						if uint32(t661) > uint32(v10) {
							p662 = t660
						}
						v16 = p662
						v12 = v16 << 1
						t663 := int32(load16(m.memory[int64(uint32(v12))+1241852:]))
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
								m.fn33(i32(-1), i32(46), i32(1227872))
								panic("unreachable")
							}
							v16 = v16 << 1
							t666 := int32(load16(m.memory[int64(uint32(v16))+1263580:]))
							t667 := int32(load16(m.memory[int64(uint32(v16))+1241852:]))
							v10 = t666 + v10 - t667
							goto l707
						}
					}
				l705:
					t668 := int32(load16(m.memory[int64(uint32(v12))+1263580:]))
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
				t670 := int32(load16(m.memory[int64(uint32(v16<<1))+1198358:]))
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
							m.fn121(v17, v5, v5, i32(1146756))
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
														t843 := int32(load16(m.memory[int64(uint32(v15*i32(188)+v11<<1))+1147142:]))
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
															t835 := int32(load16(m.memory[int64(uint32(v11<<1))+1219800:]))
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
														t836 := int32(load16(m.memory[int64(uint32(v13<<1))+1253078:]))
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
																t772 := int32(load16(m.memory[int64(uint32(v13<<1))+1260326:]))
																t773 := v13
																v13 = v10 & i32(0xffff)
																p774 := t773
																if uint32(t772) > uint32(v13) {
																	p774 = t771
																}
																v15 = p774
																t775 := v15
																v15 = v15 + i32(203)
																t776 := int32(load16(m.memory[int64(uint32(v15<<1))+1260326:]))
																p777 := v15
																if uint32(t776) > uint32(v13) {
																	p777 = t775
																}
																v15 = p777
																t778 := v15
																v15 = v15 + i32(102)
																t779 := int32(load16(m.memory[int64(uint32(v15<<1))+1260326:]))
																p780 := v15
																if uint32(t779) > uint32(v13) {
																	p780 = t778
																}
																v15 = p780
																t781 := v15
																v15 = v15 + i32(51)
																t782 := int32(load16(m.memory[int64(uint32(v15<<1))+1260326:]))
																p783 := v15
																if uint32(t782) > uint32(v13) {
																	p783 = t781
																}
																v15 = p783
																t784 := v15
																v15 = v15 + i32(25)
																t785 := int32(load16(m.memory[int64(uint32(v15<<1))+1260326:]))
																p786 := v15
																if uint32(t785) > uint32(v13) {
																	p786 = t784
																}
																v15 = p786
																t787 := v15
																v15 = v15 + i32(13)
																t788 := int32(load16(m.memory[int64(uint32(v15<<1))+1260326:]))
																p789 := v15
																if uint32(t788) > uint32(v13) {
																	p789 = t787
																}
																v15 = p789
																t790 := v15
																v15 = v15 + i32(6)
																t791 := int32(load16(m.memory[int64(uint32(v15<<1))+1260326:]))
																p792 := v15
																if uint32(t791) > uint32(v13) {
																	p792 = t790
																}
																v15 = p792
																t793 := v15
																v15 = v15 + i32(3)
																t794 := int32(load16(m.memory[int64(uint32(v15<<1))+1260326:]))
																p795 := v15
																if uint32(t794) > uint32(v13) {
																	p795 = t793
																}
																v15 = p795
																t796 := v15
																v15 = v15 + i32(2)
																t797 := int32(load16(m.memory[int64(uint32(v15<<1))+1260326:]))
																p798 := v15
																if uint32(t797) > uint32(v13) {
																	p798 = t796
																}
																v15 = p798
																t799 := v15
																v15 = v15 + i32(1)
																t800 := int32(load16(m.memory[int64(uint32(v15<<1))+1260326:]))
																p801 := v15
																if uint32(t800) > uint32(v13) {
																	p801 = t799
																}
																v15 = p801
																v12 = v15 << 1
																t802 := int32(load16(m.memory[int64(uint32(v12))+1260326:]))
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
																		m.fn33(i32(-1), i32(1627), i32(1227872))
																		panic("unreachable")
																	}
																	t805 := v10
																	v13 = v13 << 1
																	t806 := int32(load16(m.memory[int64(uint32(v13))+1260326:]))
																	t807 := int32(load16(m.memory[int64(uint32(v13))+1253122:]))
																	v10 = t805 - t806 + t807
																	v15 = v10 & i32(0xffff)
																	v13 = int32(uint32(v15) >> 12)
																	v15 = int32(uint32(v15) >> 6)
																	goto l763
																}
															}
														l761:
															t808 := int32(load16(m.memory[int64(uint32(v12))+1253122:]))
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
																t815 := int32(load16(m.memory[int64(uint32(v10<<1))+1235390:]))
																p816 := v10
																if uint32(t815) > uint32(v13) {
																	p816 = t814
																}
																v10 = p816
																t817 := v10
																v10 = v10 + i32(7)
																t818 := int32(load16(m.memory[int64(uint32(v10<<1))+1235390:]))
																p819 := v10
																if uint32(t818) > uint32(v13) {
																	p819 = t817
																}
																v10 = p819
																t820 := v10
																v10 = v10 + i32(4)
																t821 := int32(load16(m.memory[int64(uint32(v10<<1))+1235390:]))
																p822 := v10
																if uint32(t821) > uint32(v13) {
																	p822 = t820
																}
																v10 = p822
																t823 := v10
																v10 = v10 + i32(2)
																t824 := int32(load16(m.memory[int64(uint32(v10<<1))+1235390:]))
																p825 := v10
																if uint32(t824) > uint32(v13) {
																	p825 = t823
																}
																v10 = p825
																t826 := v10
																v10 = v10 + i32(1)
																t827 := int32(load16(m.memory[int64(uint32(v10<<1))+1235390:]))
																p828 := v10
																if uint32(t827) > uint32(v13) {
																	p828 = t826
																}
																v10 = p828
																v15 = v10 << 1
																t829 := int32(load16(m.memory[int64(uint32(v15))+1235390:]))
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
																		m.fn33(i32(-1), i32(59), i32(1227872))
																		panic("unreachable")
																	}
																	v10 = v10 << 1
																	t832 := int32(load16(m.memory[int64(uint32(v10))+1256376:]))
																	t833 := int32(load16(m.memory[int64(uint32(v10))+1235390:]))
																	v13 = t832 + v13 - t833
																	goto l766
																}
															}
														l764:
															t834 := int32(load16(m.memory[int64(uint32(v15))+1256376:]))
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
														m.fn33(v13, i32(101), i32(1146788))
														panic("unreachable")
													}
													t764 := int32(load16(m.memory[int64(uint32(v13<<1))+1146804:]))
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
														t724 := int32(load16(m.memory[int64(uint32(v10<<1))+1256494:]))
														p725 := v10
														if uint32(t724) > uint32(v13) {
															p725 = t723
														}
														v10 = p725
														t726 := v10
														v10 = v10 + i32(239)
														t727 := int32(load16(m.memory[int64(uint32(v10<<1))+1256494:]))
														p728 := v10
														if uint32(t727) > uint32(v13) {
															p728 = t726
														}
														v10 = p728
														t729 := v10
														v10 = v10 + i32(120)
														t730 := int32(load16(m.memory[int64(uint32(v10<<1))+1256494:]))
														p731 := v10
														if uint32(t730) > uint32(v13) {
															p731 = t729
														}
														v10 = p731
														t732 := v10
														v10 = v10 + i32(60)
														t733 := int32(load16(m.memory[int64(uint32(v10<<1))+1256494:]))
														p734 := v10
														if uint32(t733) > uint32(v13) {
															p734 = t732
														}
														v10 = p734
														t735 := v10
														v10 = v10 + i32(30)
														t736 := int32(load16(m.memory[int64(uint32(v10<<1))+1256494:]))
														p737 := v10
														if uint32(t736) > uint32(v13) {
															p737 = t735
														}
														v10 = p737
														t738 := v10
														v10 = v10 + i32(15)
														t739 := int32(load16(m.memory[int64(uint32(v10<<1))+1256494:]))
														p740 := v10
														if uint32(t739) > uint32(v13) {
															p740 = t738
														}
														v10 = p740
														t741 := v10
														v10 = v10 + i32(7)
														t742 := int32(load16(m.memory[int64(uint32(v10<<1))+1256494:]))
														p743 := v10
														if uint32(t742) > uint32(v13) {
															p743 = t741
														}
														v10 = p743
														t744 := v10
														v10 = v10 + i32(4)
														t745 := int32(load16(m.memory[int64(uint32(v10<<1))+1256494:]))
														p746 := v10
														if uint32(t745) > uint32(v13) {
															p746 = t744
														}
														v10 = p746
														t747 := v10
														v10 = v10 + i32(2)
														t748 := int32(load16(m.memory[int64(uint32(v10<<1))+1256494:]))
														p749 := v10
														if uint32(t748) > uint32(v13) {
															p749 = t747
														}
														v10 = p749
														t750 := v10
														v10 = v10 + i32(1)
														t751 := int32(load16(m.memory[int64(uint32(v10<<1))+1256494:]))
														p752 := v10
														if uint32(t751) > uint32(v13) {
															p752 = t750
														}
														v10 = p752
														v15 = v10 << 1
														t753 := int32(load16(m.memory[int64(uint32(v15))+1256494:]))
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
																m.fn33(i32(-1), i32(1916), i32(1227872))
																panic("unreachable")
															}
															t756 := v11
															v13 = v13 << 1
															t757 := int32(load16(m.memory[int64(uint32(v13))+1256494:]))
															t758 := int32(load16(m.memory[int64(uint32(v13))+1248176:]))
															v10 = t756 - t757 + t758
															v15 = v10 & i32(0xffff)
															v13 = int32(uint32(v15) >> 12)
															v15 = int32(uint32(v15) >> 6)
															goto l753
														}
													}
												l751:
													t759 := int32(load16(m.memory[int64(uint32(v15))+1248176:]))
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
															t695 := int32(load16(m.memory[int64(uint32(v10<<1))+1242072:]))
															t696 := v10
															v10 = v13 & i32(0xffff)
															p697 := t696
															if uint32(t695) > uint32(v10) {
																p697 = t694
															}
															v15 = p697
															t698 := v15
															v15 = v15 + i32(26)
															t699 := int32(load16(m.memory[int64(uint32(v15<<1))+1242072:]))
															p700 := v15
															if uint32(t699) > uint32(v10) {
																p700 = t698
															}
															v15 = p700
															t701 := v15
															v15 = v15 + i32(13)
															t702 := int32(load16(m.memory[int64(uint32(v15<<1))+1242072:]))
															p703 := v15
															if uint32(t702) > uint32(v10) {
																p703 = t701
															}
															v15 = p703
															t704 := v15
															v15 = v15 + i32(6)
															t705 := int32(load16(m.memory[int64(uint32(v15<<1))+1242072:]))
															p706 := v15
															if uint32(t705) > uint32(v10) {
																p706 = t704
															}
															v15 = p706
															t707 := v15
															v15 = v15 + i32(3)
															t708 := int32(load16(m.memory[int64(uint32(v15<<1))+1242072:]))
															p709 := v15
															if uint32(t708) > uint32(v10) {
																p709 = t707
															}
															v15 = p709
															t710 := v15
															v15 = v15 + i32(2)
															t711 := int32(load16(m.memory[int64(uint32(v15<<1))+1242072:]))
															p712 := v15
															if uint32(t711) > uint32(v10) {
																p712 = t710
															}
															v15 = p712
															t713 := v15
															v15 = v15 + i32(1)
															t714 := int32(load16(m.memory[int64(uint32(v15<<1))+1242072:]))
															p715 := v15
															if uint32(t714) > uint32(v10) {
																p715 = t713
															}
															v15 = p715
															v12 = v15 << 1
															t716 := int32(load16(m.memory[int64(uint32(v12))+1242072:]))
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
																	m.fn33(i32(-1), i32(206), i32(1227872))
																	panic("unreachable")
																}
																v10 = v10 << 1
																t719 := int32(load16(m.memory[int64(uint32(v10))+1241440:]))
																t720 := int32(load16(m.memory[int64(uint32(v10))+1242072:]))
																v13 = t719 + v13 - t720
																goto l745
															}
														}
													l743:
														t721 := int32(load16(m.memory[int64(uint32(v12))+1241440:]))
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
								m.fn900(t853, i32(1241852), i32(46), v10)
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
											m.fn33(v13, i32(46), i32(1227872))
											panic("unreachable")
										}
										v13 = v13 << 1
										t856 := int32(load16(m.memory[int64(uint32(v13))+1263580:]))
										t857 := int32(load16(m.memory[int64(uint32(v13))+1241852:]))
										v13 = t856 + v10 - t857
										goto l781
									}
								l779:
									if uint32(v13) > uint32(i32(45)) {
										m.fn33(v13, i32(46), i32(1227856))
										panic("unreachable")
									}
									t858 := int32(load16(m.memory[int64(uint32(v13<<1))+1263580:]))
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
							t859 := int32(load16(m.memory[int64(uint32(v11<<1))+1198358:]))
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
			m.fn121(v15, v3, v3, i32(1146772))
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
							m.fn121(v15, v3, v3, i32(1146708))
							panic("unreachable")
						}
						if uint32(v5) < uint32(v13) {
							m.fn121(v13, v5, v5, i32(1146692))
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
						t877 := m.fn889(t874, p876)
						v10 = t877
						if uint32(v10) > uint32(v12) {
							m.fn121(i32(0), v10, v12, i32(1146676))
							panic("unreachable")
						}
						if uint32(v10) > uint32(v11) {
							m.fn121(i32(0), v10, v11, i32(1146660))
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
func (m *Module) fn899(v0, v1, v2, v3, v4, v5, v6 int32) {
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
		m.fn898(v7+i32(4), v1, v7+i32(2), i32(2), v4, v5, i32(0))
		t1 := int32(load32(m.memory[int64(uint32(v7))+12:]))
		v6 = t1
		{
			t2 := int32(m.memory[int64(uint32(v7))+8])
			switch t2 {
			case 1:
				m.fn7(i32(1146620), i32(39), i32(1146572))
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
					m.fn121(v6, v5, v5, i32(1146572))
					panic("unreachable")
				}
				m.fn898(v7+i32(4), v1, v2, v3, v4+v6, v5-v6, i32(1))
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
		m.fn897(v0, v1, v2, v3, v4, v5, i32(0), i32(239))
		goto l9
	default:
		m.fn898(v0, v1, v2, v3, v4, v5, i32(1))
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
func (m *Module) fn900(v0, v1, v2, v3 int32) {
	var v4, v5, v6, v7 int32
	v4 = i32(0)
	v5 = v3 & i32(0xffff)
l0:
	{
		t0 := v4
		v6 = int32(uint32(v2) >> 1)
		v7 = v6 + v4
		t1 := int32(load16(m.memory[uint32(v1+v7<<1):]))
		p2 := v7
		if uint32(t1) > uint32(v5) {
			p2 = t0
		}
		v4 = p2
		v2 = v2 - v6
		if uint32(v2) > uint32(i32(1)) {
			goto l0
		}
	}
	v2 = i32(1)
	{
		{
			t3 := int32(load16(m.memory[uint32(v1+v4<<1):]))
			v6 = t3
			t4 := v6
			v7 = v3 & i32(0xffff)
			if t4 != v7 {
				goto l1
			}
			v2 = i32(0)
			goto l2
		}
	l1:
		t5 := v4
		var p6 int32
		if uint32(v6) < uint32(v7) {
			p6 = 1
		}
		v4 = t5 + p6
	}
l2:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v4))
	store32(m.memory[uint32(v0):], uint32(v2))
}
func (m *Module) fn901(v0, v1 int32) {
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
		t0 := int32(load16(m.memory[int64(uint32(v2<<1))+1242484:]))
		v3 = v3 + t0
		if uint32(v3) > uint32(i32(239)) {
			m.fn33(v3, i32(240), i32(1242056))
			panic("unreachable")
		}
		t1 := int32(load16(m.memory[int64(uint32(v3<<1))+1227044:]))
		v3 = t1
	}
l10:
	store16(m.memory[int64(uint32(v0))+2:], uint16(v3))
	store16(m.memory[uint32(v0):], uint16(v1))
}
func (m *Module) fn902(v0, v1 int32) {
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
			m.fn42(i32(1284936), i32(43), v2+i32(15), i32(1285048), i32(1285064))
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
			m.fn7(i32(1274404), i32(46), i32(1274452))
			panic("unreachable")
		}
		if v3 == 0 {
			goto l3
		}
		if uint32(v4) > uint32(v0+i32(39)) {
			m.fn7(i32(1274468), i32(46), i32(1274516))
			panic("unreachable")
		}
	l3:
		m.fn5(v1)
	}
l0:
	m.g0 = v2 + i32(16)
}
func (m *Module) fn903(v0, v1, v2 int32) int32 {
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
		m.fn42(i32(1284936), i32(43), v3+i32(15), i32(1285048), i32(1285080))
		panic("unreachable")
	l2:
		t1 := m.fn832(i32(64), v1)
		v1 = t1
	}
l1:
	m.g0 = v3 + i32(16)
	return v1
}
func (m *Module) fn904(v0 int32) {
	var v1 int32
	t0 := m.g0
	v1 = t0 - i32(16)
	m.g0 = v1
	store32(m.memory[int64(uint32(v1))+12:], uint32(i32(1277668)))
	store32(m.memory[int64(uint32(v1))+8:], uint32(v0))
	m.fn845(i32(0), v1+i32(8), i32(1280052), v1+i32(12), i32(1280052), i32(0), v1, i32(0x13999c))
	panic("unreachable")
}
func (m *Module) fn905(v0, v1, v2, v3, v4, v5, v6, v7 int32) {
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
				m.fn33(v14, i32(16), i32(1290216))
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
			m.fn28(i32(1051134), v8+i32(40), i32(1290232))
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
						m.fn33(v12, i32(16), i32(1290248))
						panic("unreachable")
					}
					v12 = v8 + i32(40) + v12<<1
					t50 := int32(load16(m.memory[uint32(v12):]))
					t51 := v12
					v12 = t50
					store16(m.memory[uint32(t51):], uint16(v12+i32(1)))
					if uint32(v12) >= uint32(i32(288)) {
						m.fn33(v12, i32(288), i32(1290264))
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
				v29 = i32(1290280)
				v28 = i32(1290342)
				goto l14
			case 2:
				if uint32(v23) > uint32(i32(9)) {
					goto l17
				}
				v22 = i32(0)
				v27 = i32(32)
				v29 = i32(1290374)
				v28 = i32(1290438)
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
							m.fn33(v14, v27, i32(1290488))
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
						m.fn33(v12, v5, i32(1290504))
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
						m.fn33(v13, i32(16), i32(1290520))
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
								m.fn121(v24, v5, v5, i32(1290584))
								panic("unreachable")
							}
							{
								t68 := v16
								v14 = v5 - v24
								if uint32(t68) >= uint32(v14) {
									m.fn33(v16, v14, i32(1290568))
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
							m.fn33(i32(288), i32(288), i32(1290536))
							panic("unreachable")
						}
						t65 := int32(load16(m.memory[uint32(v7+v15<<1):]))
						t66 := v3
						v14 = t65
						if uint32(t66) <= uint32(v14) {
							m.fn33(v14, v3, i32(1290552))
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
						m.fn33(v6, v5, i32(1290600))
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
			m.fn33(i32(288), i32(288), i32(1290472))
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
func (m *Module) fn906(v0, v1, v2 int32) int32 {
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
func (m *Module) fn907(v0, v1, v2 int32) int32 {
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
			t5 := int32(load32(m.memory[int64(uint32((t4^v0)&i32(255)<<2))+1285096:]))
			v0 = t5 ^ int32(uint32(v0)>>8)
			t6 := int32(load32(m.memory[int64(uint32((t3^v0)&i32(255)<<2))+1285096:]))
			v0 = t6 ^ int32(uint32(v0)>>8)
			if v2&i32(1) == 0 {
				goto l2
			}
		}
	l3:
		t7 := int32(m.memory[uint32(v1+v4)])
		t8 := int32(load32(m.memory[int64(uint32((t7^v0)&i32(255)<<2))+1285096:]))
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
		m.fn33(v1, v6, i32(1284196))
		panic("unreachable")
	l10:
		t13 := int32(load32(m.memory[uint32(v2+i32(8)):]))
		v4 = v12 ^ t13
		t14 := int32(load32(m.memory[int64(uint32(v4&i32(255)<<2))+1280084:]))
		t15 := int32(load32(m.memory[int64(uint32(int32(uint32(v4)>>6)&i32(1020)))+1281108:]))
		t16 := int32(load32(m.memory[int64(uint32(int32(uint32(v4)>>14)&i32(1020)))+1282132:]))
		t17 := int32(load32(m.memory[int64(uint32(int32(uint32(v4)>>22)&i32(1020)))+1283156:]))
		v12 = t14 ^ t15 ^ t16 ^ t17
		t18 := int32(load32(m.memory[uint32(v2+i32(4)):]))
		v4 = v9 ^ t18
		t19 := int32(load32(m.memory[int64(uint32(v4&i32(255)<<2))+1280084:]))
		t20 := int32(load32(m.memory[int64(uint32(int32(uint32(v4)>>6)&i32(1020)))+1281108:]))
		t21 := int32(load32(m.memory[int64(uint32(int32(uint32(v4)>>14)&i32(1020)))+1282132:]))
		t22 := int32(load32(m.memory[int64(uint32(int32(uint32(v4)>>22)&i32(1020)))+1283156:]))
		v9 = t19 ^ t20 ^ t21 ^ t22
		t23 := int32(load32(m.memory[uint32(v2):]))
		v4 = v0 ^ t23
		t24 := int32(load32(m.memory[int64(uint32(v4&i32(255)<<2))+1280084:]))
		t25 := int32(load32(m.memory[int64(uint32(int32(uint32(v4)>>6)&i32(1020)))+1281108:]))
		t26 := int32(load32(m.memory[int64(uint32(int32(uint32(v4)>>14)&i32(1020)))+1282132:]))
		t27 := int32(load32(m.memory[int64(uint32(int32(uint32(v4)>>22)&i32(1020)))+1283156:]))
		v0 = t24 ^ t25 ^ t26 ^ t27
		t28 := int32(load32(m.memory[uint32(v2+i32(16)):]))
		v4 = v14 ^ t28
		t29 := int32(load32(m.memory[int64(uint32(v4&i32(255)<<2))+1280084:]))
		t30 := int32(load32(m.memory[int64(uint32(int32(uint32(v4)>>6)&i32(1020)))+1281108:]))
		t31 := int32(load32(m.memory[int64(uint32(int32(uint32(v4)>>14)&i32(1020)))+1282132:]))
		t32 := int32(load32(m.memory[int64(uint32(int32(uint32(v4)>>22)&i32(1020)))+1283156:]))
		v14 = t29 ^ t30 ^ t31 ^ t32
		t33 := int32(load32(m.memory[uint32(v2+i32(12)):]))
		v4 = t33 ^ v13
		t34 := int32(load32(m.memory[int64(uint32(v4&i32(255)<<2))+1280084:]))
		t35 := int32(load32(m.memory[int64(uint32(int32(uint32(v4)>>6)&i32(1020)))+1281108:]))
		t36 := int32(load32(m.memory[int64(uint32(int32(uint32(v4)>>14)&i32(1020)))+1282132:]))
		t37 := int32(load32(m.memory[int64(uint32(int32(uint32(v4)>>22)&i32(1020)))+1283156:]))
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
			m.fn121(v2, v6, v6, i32(1284180))
			panic("unreachable")
		}
		{
			if v6 == v2 {
				goto l13
			}
			v11 = v7 + v2<<2
			t39 := int32(load32(m.memory[uint32(v11):]))
			v1 = t39
			memory_copy(m.memory, uint32(v3), uint32(i32(1286120)), uint32(i32(1024)))
			t40 := v3
			v1 = v1 ^ v0
			t41 := int32(load32(m.memory[uint32(t40+v1&i32(255)<<2):]))
			v4 = t41
			memory_copy(m.memory, uint32(v3), uint32(i32(1287144)), uint32(i32(1024)))
			t42 := int32(load32(m.memory[uint32(v3+int32(uint32(v1)>>6)&i32(1020)):]))
			v0 = t42
			memory_copy(m.memory, uint32(v3), uint32(i32(1288168)), uint32(i32(1024)))
			t43 := int32(load32(m.memory[uint32(v3+int32(uint32(v1)>>14)&i32(1020)):]))
			v15 = t43
			memory_copy(m.memory, uint32(v3), uint32(i32(1289192)), uint32(i32(1024)))
			t44 := int32(load32(m.memory[uint32(v3+int32(uint32(v1)>>22)&i32(1020)):]))
			v0 = v15 ^ (v0 ^ v4) ^ t44
			v1 = v6 - v2
			if v1 == i32(1) {
				goto l13
			}
			t45 := int32(load32(m.memory[int64(uint32(v11))+4:]))
			v2 = t45
			memory_copy(m.memory, uint32(v3), uint32(i32(1286120)), uint32(i32(1024)))
			t46 := v3
			v2 = v2 ^ v9 ^ v0
			t47 := int32(load32(m.memory[uint32(t46+v2&i32(255)<<2):]))
			v4 = t47
			memory_copy(m.memory, uint32(v3), uint32(i32(1287144)), uint32(i32(1024)))
			t48 := int32(load32(m.memory[uint32(v3+int32(uint32(v2)>>6)&i32(1020)):]))
			v0 = t48
			memory_copy(m.memory, uint32(v3), uint32(i32(1288168)), uint32(i32(1024)))
			t49 := int32(load32(m.memory[uint32(v3+int32(uint32(v2)>>14)&i32(1020)):]))
			v9 = t49
			memory_copy(m.memory, uint32(v3), uint32(i32(1289192)), uint32(i32(1024)))
			t50 := int32(load32(m.memory[uint32(v3+int32(uint32(v2)>>22)&i32(1020)):]))
			v0 = v9 ^ (v0 ^ v4) ^ t50
			if v1 == i32(2) {
				goto l13
			}
			t51 := int32(load32(m.memory[int64(uint32(v11))+8:]))
			v2 = t51
			memory_copy(m.memory, uint32(v3), uint32(i32(1286120)), uint32(i32(1024)))
			t52 := v3
			v2 = v2 ^ v12 ^ v0
			t53 := int32(load32(m.memory[uint32(t52+v2&i32(255)<<2):]))
			v4 = t53
			memory_copy(m.memory, uint32(v3), uint32(i32(1287144)), uint32(i32(1024)))
			t54 := int32(load32(m.memory[uint32(v3+int32(uint32(v2)>>6)&i32(1020)):]))
			v0 = t54
			memory_copy(m.memory, uint32(v3), uint32(i32(1288168)), uint32(i32(1024)))
			t55 := int32(load32(m.memory[uint32(v3+int32(uint32(v2)>>14)&i32(1020)):]))
			v9 = t55
			memory_copy(m.memory, uint32(v3), uint32(i32(1289192)), uint32(i32(1024)))
			t56 := int32(load32(m.memory[uint32(v3+int32(uint32(v2)>>22)&i32(1020)):]))
			v0 = v9 ^ (v0 ^ v4) ^ t56
			if v1 == i32(3) {
				goto l13
			}
			t57 := int32(load32(m.memory[int64(uint32(v11))+12:]))
			v2 = t57
			memory_copy(m.memory, uint32(v3), uint32(i32(1286120)), uint32(i32(1024)))
			t58 := v3
			v2 = v2 ^ v13 ^ v0
			t59 := int32(load32(m.memory[uint32(t58+v2&i32(255)<<2):]))
			v4 = t59
			memory_copy(m.memory, uint32(v3), uint32(i32(1287144)), uint32(i32(1024)))
			t60 := int32(load32(m.memory[uint32(v3+int32(uint32(v2)>>6)&i32(1020)):]))
			v0 = t60
			memory_copy(m.memory, uint32(v3), uint32(i32(1288168)), uint32(i32(1024)))
			t61 := int32(load32(m.memory[uint32(v3+int32(uint32(v2)>>14)&i32(1020)):]))
			v9 = t61
			memory_copy(m.memory, uint32(v3), uint32(i32(1289192)), uint32(i32(1024)))
			t62 := int32(load32(m.memory[uint32(v3+int32(uint32(v2)>>22)&i32(1020)):]))
			v0 = v9 ^ (v0 ^ v4) ^ t62
			if v1 == i32(4) {
				goto l13
			}
			t63 := int32(load32(m.memory[int64(uint32(v11))+16:]))
			v2 = t63
			memory_copy(m.memory, uint32(v3), uint32(i32(1286120)), uint32(i32(1024)))
			t64 := v3
			v2 = v2 ^ v14 ^ v0
			t65 := int32(load32(m.memory[uint32(t64+v2&i32(255)<<2):]))
			v4 = t65
			memory_copy(m.memory, uint32(v3), uint32(i32(1287144)), uint32(i32(1024)))
			t66 := int32(load32(m.memory[uint32(v3+int32(uint32(v2)>>6)&i32(1020)):]))
			v0 = t66
			memory_copy(m.memory, uint32(v3), uint32(i32(1288168)), uint32(i32(1024)))
			t67 := int32(load32(m.memory[uint32(v3+int32(uint32(v2)>>14)&i32(1020)):]))
			v9 = t67
			memory_copy(m.memory, uint32(v3), uint32(i32(1289192)), uint32(i32(1024)))
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
				memory_copy(m.memory, uint32(v3), uint32(i32(1286120)), uint32(i32(1024)))
				t72 := v3
				v2 = v2 ^ v0
				t73 := int32(load32(m.memory[uint32(t72+v2&i32(255)<<2):]))
				v6 = t73
				memory_copy(m.memory, uint32(v3), uint32(i32(1287144)), uint32(i32(1024)))
				t74 := int32(load32(m.memory[uint32(v3+int32(uint32(v2)>>6)&i32(1020)):]))
				v0 = t74
				memory_copy(m.memory, uint32(v3), uint32(i32(1288168)), uint32(i32(1024)))
				t75 := int32(load32(m.memory[uint32(v3+int32(uint32(v2)>>14)&i32(1020)):]))
				v9 = t75
				memory_copy(m.memory, uint32(v3), uint32(i32(1289192)), uint32(i32(1024)))
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
			t78 := int32(load32(m.memory[int64(uint32((t77^v0)&i32(255)<<2))+1285096:]))
			v0 = t78 ^ int32(uint32(v0)>>8)
			if v8 == i32(1) {
				goto l15
			}
			t79 := int32(m.memory[int64(uint32(v5))+1])
			t80 := int32(load32(m.memory[int64(uint32((t79^v0)&i32(255)<<2))+1285096:]))
			v0 = t80 ^ int32(uint32(v0)>>8)
			if v8 == i32(2) {
				goto l15
			}
			t81 := int32(m.memory[int64(uint32(v5))+2])
			t82 := int32(load32(m.memory[int64(uint32((t81^v0)&i32(255)<<2))+1285096:]))
			v0 = t82 ^ int32(uint32(v0)>>8)
		}
	l15:
		m.g0 = v3 + i32(1024)
		return v0 ^ i32(-1)
	}
}
func (m *Module) fn908(v0, v1, v2 int32) {
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
			m.fn140(i32(1277592), i32(9), i32(1277604))
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
			m.fn121(i32(0), v1, v4, i32(1290616))
			panic("unreachable")
		}
		if uint32(v0) > uint32(v1) {
			m.fn121(v0, v1, v4, i32(1290632))
			panic("unreachable")
		}
		if uint32(v3) > uint32(v4-v2) {
			m.fn28(i32(1277472), i32(43), i32(1277576))
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
					m.fn33(v6, v4, i32(1277544))
					panic("unreachable")
				}
				if v0 == 0 {
					m.fn33(v3, v4, i32(1277560))
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
			m.fn33(v0, v4, i32(1277496))
			panic("unreachable")
		}
		if uint32(v4) < uint32(v3) {
			m.fn121(v3, v4, v4, i32(1277528))
			panic("unreachable")
		}
		t11 := v2
		v4 = v4 - v3
		if uint32(t11) > uint32(v4) {
			m.fn121(i32(0), v2, v4, i32(1277512))
			panic("unreachable")
		}
		if v2 == 0 {
			return
		}
		t12 := int32(m.memory[uint32(v7+v0)])
		memory_fill(m.memory, uint32(v7+v3), t12, uint32(v2))
	}
}
func (m *Module) fn909(v0, v1, v2, v3 int32) {
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
				m.fn121(i32(0), v8, v9, i32(1284368))
				panic("unreachable")
			}
			if uint32(v3) < uint32(v2) {
				goto l2
			}
			if uint32(v3) > uint32(v8) {
				goto l2
			}
			if uint32(v4) < uint32(v5) {
				m.fn121(v5, v4, v4, i32(1277636))
				panic("unreachable")
			}
			if uint32(v7) <= uint32(v6) {
				goto l4
			}
			m.fn121(i32(0), v7, v6, i32(1277620))
			panic("unreachable")
		}
	l2:
		m.fn121(v2, v3, v8, i32(1277652))
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
func (m *Module) fn910(v0 int32) {
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
	v9 = i32(1277688)
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
	v11 = i32(1279736)
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
			m.fn7(i32(1284212), i32(74), i32(1284288))
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
									m.fn33(v5, v10, i32(1279864))
									panic("unreachable")
								}
							l11:
								t21 := v10
								v5 = v13 & v24
								if uint32(t21) <= uint32(v5) {
									m.fn33(v5, v10, i32(1279880))
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
							m.fn33(v7, v8, i32(1279896))
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
								m.fn33(v5, v10, i32(1279912))
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
							m.fn33(v16, v8, i32(1279896))
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
							m.fn33(v5, v10, i32(1279928))
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
											v9 = i32(1065867)
											v3 = i32(1)
											v10 = i32(22)
											goto l42
										}
										t47 := v12
										v5 = (v17 + int32(v29&(i64_shl(i64(-1), int64(uint32(v13))&i64(47))^i64(-1)))) & i32(0xffff)
										if uint32(t47) <= uint32(v5) {
											m.fn33(v5, v12, i32(1279976))
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
												m.fn28(i32(1279992), i32(85), i32(1280036))
												panic("unreachable")
											}
											v10 = i32(30)
											m.memory[uint32(v0)] = byte(i32(30))
											v9 = i32(1065425)
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
									m.fn908(v1+i32(4), v17, v3&i32(0xffff))
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
								m.fn909(v1+i32(4), v21, v16, v23)
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
							m.fn909(t57, t58, t59, p61+v16)
							if uint32(v3) > uint32(v5) {
								m.fn908(v1+i32(4), v17, v3-v5)
								v17 = v8
								goto l36
							}
							v17 = v8
							goto l36
						}
						m.fn33(v13, v12, i32(1279960))
						panic("unreachable")
					}
					if v5&i32(64) != 0 {
						if v5&i32(32) == 0 {
							m.memory[uint32(v0)] = byte(i32(30))
							v9 = i32(1065839)
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
					m.fn33(v5, v10, i32(1279944))
					panic("unreachable")
				l21:
					if uint32(v16) >= uint32(v8) {
						m.fn33(v16, v8, i32(1279896))
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
func (m *Module) fn911(v0, v1, v2, v3 int32) {
	var v4 int32
	t0 := m.g0
	v4 = t0 - i32(16)
	m.g0 = v4
	if v1 != 0 {
		goto l0
	}
	m.fn912()
	panic("unreachable")
l0:
	t1 := int32(load32(m.memory[int64(uint32(v2))+16:]))
	m.t0[uint(t1)].(func(int32, int32, int32))(v4+i32(8), v1, v3)
	t2 := int32(load32(m.memory[int64(uint32(v4))+12:]))
	v1 = t2
	t3 := int32(load32(m.memory[int64(uint32(v4))+8:]))
	t4 := v0
	v3 = t3
	store32(m.memory[int64(uint32(t4))+4:], uint32(v3))
	t6 := v0
	p5 := i32(0)
	if v3&i32(1) != 0 {
		p5 = v1
	}
	store32(m.memory[uint32(t6):], uint32(p5))
	m.g0 = v4 + i32(16)
}
func (m *Module) fn912() {
	m.fn1867()
	panic("unreachable")
}
func (m *Module) fn913() {
	m.fn0(i32(25))
	m.fn0(i32(12))
	m.fn0(i32(70))
	m.fn0(i32(108))
	m.fn0(i32(111))
	m.fn0(i32(97))
	m.fn0(i32(116))
	m.fn0(i32(54))
	m.fn0(i32(52))
	m.fn0(i32(65))
	m.fn0(i32(114))
	m.fn0(i32(114))
	m.fn0(i32(97))
	m.fn0(i32(121))
}
func (m *Module) fn914() {
	m.fn0(i32(25))
	m.fn0(i32(12))
	m.fn0(i32(70))
	m.fn0(i32(108))
	m.fn0(i32(111))
	m.fn0(i32(97))
	m.fn0(i32(116))
	m.fn0(i32(51))
	m.fn0(i32(50))
	m.fn0(i32(65))
	m.fn0(i32(114))
	m.fn0(i32(114))
	m.fn0(i32(97))
	m.fn0(i32(121))
}
func (m *Module) fn915() {
	m.fn0(i32(25))
	m.fn0(i32(10))
	m.fn0(i32(85))
	m.fn0(i32(105))
	m.fn0(i32(110))
	m.fn0(i32(116))
	m.fn0(i32(56))
	m.fn0(i32(65))
	m.fn0(i32(114))
	m.fn0(i32(114))
	m.fn0(i32(97))
	m.fn0(i32(121))
}
func (m *Module) fn916() {
	m.fn0(i32(25))
	m.fn0(i32(17))
	m.fn0(i32(85))
	m.fn0(i32(105))
	m.fn0(i32(110))
	m.fn0(i32(116))
	m.fn0(i32(56))
	m.fn0(i32(67))
	m.fn0(i32(108))
	m.fn0(i32(97))
	m.fn0(i32(109))
	m.fn0(i32(112))
	m.fn0(i32(101))
	m.fn0(i32(100))
	m.fn0(i32(65))
	m.fn0(i32(114))
	m.fn0(i32(114))
	m.fn0(i32(97))
	m.fn0(i32(121))
}
func (m *Module) fn917() {
	m.fn0(i32(25))
	m.fn0(i32(10))
	m.fn0(i32(73))
	m.fn0(i32(110))
	m.fn0(i32(116))
	m.fn0(i32(51))
	m.fn0(i32(50))
	m.fn0(i32(65))
	m.fn0(i32(114))
	m.fn0(i32(114))
	m.fn0(i32(97))
	m.fn0(i32(121))
}
func (m *Module) fn918() {
	m.fn0(i32(25))
	m.fn0(i32(11))
	m.fn0(i32(85))
	m.fn0(i32(105))
	m.fn0(i32(110))
	m.fn0(i32(116))
	m.fn0(i32(51))
	m.fn0(i32(50))
	m.fn0(i32(65))
	m.fn0(i32(114))
	m.fn0(i32(114))
	m.fn0(i32(97))
	m.fn0(i32(121))
}
func (m *Module) fn919() {
	m.fn0(i32(25))
	m.fn0(i32(10))
	m.fn0(i32(73))
	m.fn0(i32(110))
	m.fn0(i32(116))
	m.fn0(i32(49))
	m.fn0(i32(54))
	m.fn0(i32(65))
	m.fn0(i32(114))
	m.fn0(i32(114))
	m.fn0(i32(97))
	m.fn0(i32(121))
}
func (m *Module) fn920() {
	m.fn0(i32(25))
	m.fn0(i32(11))
	m.fn0(i32(85))
	m.fn0(i32(105))
	m.fn0(i32(110))
	m.fn0(i32(116))
	m.fn0(i32(49))
	m.fn0(i32(54))
	m.fn0(i32(65))
	m.fn0(i32(114))
	m.fn0(i32(114))
	m.fn0(i32(97))
	m.fn0(i32(121))
}
func (m *Module) fn921() {
	m.fn0(i32(25))
	m.fn0(i32(13))
	m.fn0(i32(66))
	m.fn0(i32(105))
	m.fn0(i32(103))
	m.fn0(i32(73))
	m.fn0(i32(110))
	m.fn0(i32(116))
	m.fn0(i32(54))
	m.fn0(i32(52))
	m.fn0(i32(65))
	m.fn0(i32(114))
	m.fn0(i32(114))
	m.fn0(i32(97))
	m.fn0(i32(121))
}
func (m *Module) fn922() {
	m.fn0(i32(25))
	m.fn0(i32(14))
	m.fn0(i32(66))
	m.fn0(i32(105))
	m.fn0(i32(103))
	m.fn0(i32(85))
	m.fn0(i32(105))
	m.fn0(i32(110))
	m.fn0(i32(116))
	m.fn0(i32(54))
	m.fn0(i32(52))
	m.fn0(i32(65))
	m.fn0(i32(114))
	m.fn0(i32(114))
	m.fn0(i32(97))
	m.fn0(i32(121))
}
func (m *Module) fn923() {
	m.fn0(i32(25))
	m.fn0(i32(17))
	m.fn0(i32(83))
	m.fn0(i32(104))
	m.fn0(i32(97))
	m.fn0(i32(114))
	m.fn0(i32(101))
	m.fn0(i32(100))
	m.fn0(i32(65))
	m.fn0(i32(114))
	m.fn0(i32(114))
	m.fn0(i32(97))
	m.fn0(i32(121))
	m.fn0(i32(66))
	m.fn0(i32(117))
	m.fn0(i32(102))
	m.fn0(i32(102))
	m.fn0(i32(101))
	m.fn0(i32(114))
}
func (m *Module) fn924() {
	m.fn0(i32(25))
	m.fn0(i32(18))
	m.fn0(i32(65))
	m.fn0(i32(114))
	m.fn0(i32(114))
	m.fn0(i32(97))
	m.fn0(i32(121))
	m.fn0(i32(66))
	m.fn0(i32(117))
	m.fn0(i32(102))
	m.fn0(i32(102))
	m.fn0(i32(101))
	m.fn0(i32(114))
	m.fn0(i32(79))
	m.fn0(i32(112))
	m.fn0(i32(116))
	m.fn0(i32(105))
	m.fn0(i32(111))
	m.fn0(i32(110))
	m.fn0(i32(115))
}
func (m *Module) fn925() {
	m.fn0(i32(25))
	m.fn0(i32(11))
	m.fn0(i32(65))
	m.fn0(i32(114))
	m.fn0(i32(114))
	m.fn0(i32(97))
	m.fn0(i32(121))
	m.fn0(i32(66))
	m.fn0(i32(117))
	m.fn0(i32(102))
	m.fn0(i32(102))
	m.fn0(i32(101))
	m.fn0(i32(114))
}
func (m *Module) fn926() {
	m.fn0(i32(25))
	m.fn0(i32(25))
	m.fn0(i32(70))
	m.fn0(i32(105))
	m.fn0(i32(110))
	m.fn0(i32(97))
	m.fn0(i32(108))
	m.fn0(i32(105))
	m.fn0(i32(122))
	m.fn0(i32(97))
	m.fn0(i32(116))
	m.fn0(i32(105))
	m.fn0(i32(111))
	m.fn0(i32(110))
	m.fn0(i32(82))
	m.fn0(i32(101))
	m.fn0(i32(103))
	m.fn0(i32(105))
	m.fn0(i32(115))
	m.fn0(i32(116))
	m.fn0(i32(114))
	m.fn0(i32(121))
	m.fn0(i32(60))
	m.fn0(i32(97))
	m.fn0(i32(110))
	m.fn0(i32(121))
	m.fn0(i32(62))
}
func (m *Module) fn927() {
	m.fn0(i32(25))
	m.fn0(i32(16))
	m.fn0(i32(82))
	m.fn0(i32(101))
	m.fn0(i32(103))
	m.fn0(i32(69))
	m.fn0(i32(120))
	m.fn0(i32(112))
	m.fn0(i32(77))
	m.fn0(i32(97))
	m.fn0(i32(116))
	m.fn0(i32(99))
	m.fn0(i32(104))
	m.fn0(i32(65))
	m.fn0(i32(114))
	m.fn0(i32(114))
	m.fn0(i32(97))
	m.fn0(i32(121))
}
func (m *Module) fn928() {
	m.fn0(i32(25))
	m.fn0(i32(16))
	m.fn0(i32(80))
	m.fn0(i32(114))
	m.fn0(i32(111))
	m.fn0(i32(120))
	m.fn0(i32(121))
	m.fn0(i32(67))
	m.fn0(i32(111))
	m.fn0(i32(110))
	m.fn0(i32(115))
	m.fn0(i32(116))
	m.fn0(i32(114))
	m.fn0(i32(117))
	m.fn0(i32(99))
	m.fn0(i32(116))
	m.fn0(i32(111))
	m.fn0(i32(114))
}
func (m *Module) fn929() {
	m.fn0(i32(25))
	m.fn0(i32(10))
	m.fn0(i32(82))
	m.fn0(i32(97))
	m.fn0(i32(110))
	m.fn0(i32(103))
	m.fn0(i32(101))
	m.fn0(i32(69))
	m.fn0(i32(114))
	m.fn0(i32(114))
	m.fn0(i32(111))
	m.fn0(i32(114))
}
func (m *Module) fn930() {
	m.fn0(i32(25))
	m.fn0(i32(14))
	m.fn0(i32(82))
	m.fn0(i32(101))
	m.fn0(i32(102))
	m.fn0(i32(101))
	m.fn0(i32(114))
	m.fn0(i32(101))
	m.fn0(i32(110))
	m.fn0(i32(99))
	m.fn0(i32(101))
	m.fn0(i32(69))
	m.fn0(i32(114))
	m.fn0(i32(114))
	m.fn0(i32(111))
	m.fn0(i32(114))
}
func (m *Module) fn931() {
	m.fn0(i32(25))
	m.fn0(i32(14))
	m.fn0(i32(65))
	m.fn0(i32(103))
	m.fn0(i32(103))
	m.fn0(i32(114))
	m.fn0(i32(101))
	m.fn0(i32(103))
	m.fn0(i32(97))
	m.fn0(i32(116))
	m.fn0(i32(101))
	m.fn0(i32(69))
	m.fn0(i32(114))
	m.fn0(i32(114))
	m.fn0(i32(111))
	m.fn0(i32(114))
}
func (m *Module) fn932() {
	m.fn0(i32(25))
	m.fn0(i32(11))
	m.fn0(i32(83))
	m.fn0(i32(121))
	m.fn0(i32(110))
	m.fn0(i32(116))
	m.fn0(i32(97))
	m.fn0(i32(120))
	m.fn0(i32(69))
	m.fn0(i32(114))
	m.fn0(i32(114))
	m.fn0(i32(111))
	m.fn0(i32(114))
}
func (m *Module) fn933() {
	m.fn0(i32(25))
	m.fn0(i32(12))
	m.fn0(i32(70))
	m.fn0(i32(108))
	m.fn0(i32(111))
	m.fn0(i32(97))
	m.fn0(i32(116))
	m.fn0(i32(49))
	m.fn0(i32(54))
	m.fn0(i32(65))
	m.fn0(i32(114))
	m.fn0(i32(114))
	m.fn0(i32(97))
	m.fn0(i32(121))
}
func (m *Module) fn934() {
	m.fn0(i32(25))
	m.fn0(i32(12))
	m.fn0(i32(69))
	m.fn0(i32(114))
	m.fn0(i32(114))
	m.fn0(i32(111))
	m.fn0(i32(114))
	m.fn0(i32(79))
	m.fn0(i32(112))
	m.fn0(i32(116))
	m.fn0(i32(105))
	m.fn0(i32(111))
	m.fn0(i32(110))
	m.fn0(i32(115))
}
func (m *Module) fn935() {
	m.fn0(i32(25))
	m.fn0(i32(15))
	m.fn0(i32(87))
	m.fn0(i32(101))
	m.fn0(i32(98))
	m.fn0(i32(65))
	m.fn0(i32(115))
	m.fn0(i32(115))
	m.fn0(i32(101))
	m.fn0(i32(109))
	m.fn0(i32(98))
	m.fn0(i32(108))
	m.fn0(i32(121))
	m.fn0(i32(46))
	m.fn0(i32(84))
	m.fn0(i32(97))
	m.fn0(i32(103))
}
func (m *Module) fn936() {
	m.fn0(i32(25))
	m.fn0(i32(21))
	m.fn0(i32(87))
	m.fn0(i32(101))
	m.fn0(i32(98))
	m.fn0(i32(65))
	m.fn0(i32(115))
	m.fn0(i32(115))
	m.fn0(i32(101))
	m.fn0(i32(109))
	m.fn0(i32(98))
	m.fn0(i32(108))
	m.fn0(i32(121))
	m.fn0(i32(46))
	m.fn0(i32(69))
	m.fn0(i32(120))
	m.fn0(i32(99))
	m.fn0(i32(101))
	m.fn0(i32(112))
	m.fn0(i32(116))
	m.fn0(i32(105))
	m.fn0(i32(111))
	m.fn0(i32(110))
}
func (m *Module) fn937() {
	m.fn0(i32(25))
	m.fn0(i32(18))
	m.fn0(i32(87))
	m.fn0(i32(101))
	m.fn0(i32(98))
	m.fn0(i32(65))
	m.fn0(i32(115))
	m.fn0(i32(115))
	m.fn0(i32(101))
	m.fn0(i32(109))
	m.fn0(i32(98))
	m.fn0(i32(108))
	m.fn0(i32(121))
	m.fn0(i32(46))
	m.fn0(i32(71))
	m.fn0(i32(108))
	m.fn0(i32(111))
	m.fn0(i32(98))
	m.fn0(i32(97))
	m.fn0(i32(108))
}
func (m *Module) fn938() {
	m.fn0(i32(25))
	m.fn0(i32(18))
	m.fn0(i32(87))
	m.fn0(i32(101))
	m.fn0(i32(98))
	m.fn0(i32(65))
	m.fn0(i32(115))
	m.fn0(i32(115))
	m.fn0(i32(101))
	m.fn0(i32(109))
	m.fn0(i32(98))
	m.fn0(i32(108))
	m.fn0(i32(121))
	m.fn0(i32(46))
	m.fn0(i32(77))
	m.fn0(i32(101))
	m.fn0(i32(109))
	m.fn0(i32(111))
	m.fn0(i32(114))
	m.fn0(i32(121))
}
func (m *Module) fn939() {
	m.fn0(i32(25))
	m.fn0(i32(24))
	m.fn0(i32(87))
	m.fn0(i32(101))
	m.fn0(i32(98))
	m.fn0(i32(65))
	m.fn0(i32(115))
	m.fn0(i32(115))
	m.fn0(i32(101))
	m.fn0(i32(109))
	m.fn0(i32(98))
	m.fn0(i32(108))
	m.fn0(i32(121))
	m.fn0(i32(46))
	m.fn0(i32(67))
	m.fn0(i32(111))
	m.fn0(i32(109))
	m.fn0(i32(112))
	m.fn0(i32(105))
	m.fn0(i32(108))
	m.fn0(i32(101))
	m.fn0(i32(69))
	m.fn0(i32(114))
	m.fn0(i32(114))
	m.fn0(i32(111))
	m.fn0(i32(114))
}
func (m *Module) fn940() {
	m.fn0(i32(25))
	m.fn0(i32(20))
	m.fn0(i32(87))
	m.fn0(i32(101))
	m.fn0(i32(98))
	m.fn0(i32(65))
	m.fn0(i32(115))
	m.fn0(i32(115))
	m.fn0(i32(101))
	m.fn0(i32(109))
	m.fn0(i32(98))
	m.fn0(i32(108))
	m.fn0(i32(121))
	m.fn0(i32(46))
	m.fn0(i32(73))
	m.fn0(i32(110))
	m.fn0(i32(115))
	m.fn0(i32(116))
	m.fn0(i32(97))
	m.fn0(i32(110))
	m.fn0(i32(99))
	m.fn0(i32(101))
}
func (m *Module) fn941() {
	m.fn0(i32(25))
	m.fn0(i32(21))
	m.fn0(i32(87))
	m.fn0(i32(101))
	m.fn0(i32(98))
	m.fn0(i32(65))
	m.fn0(i32(115))
	m.fn0(i32(115))
	m.fn0(i32(101))
	m.fn0(i32(109))
	m.fn0(i32(98))
	m.fn0(i32(108))
	m.fn0(i32(121))
	m.fn0(i32(46))
	m.fn0(i32(76))
	m.fn0(i32(105))
	m.fn0(i32(110))
	m.fn0(i32(107))
	m.fn0(i32(69))
	m.fn0(i32(114))
	m.fn0(i32(114))
	m.fn0(i32(111))
	m.fn0(i32(114))
}
