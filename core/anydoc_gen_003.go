package core

import (
	"math/bits"
)

func (m *Module) fn87(v0, v1 int64, v2 int32) int64 {
	var v3 int32
	var v4, v5, v6, v7 int64
	t0 := m.g0
	v3 = t0 - i32(80)
	m.g0 = v3
	store64(m.memory[int64(uint32(v3))+56:], uint64(i64(0)))
	store64(m.memory[int64(uint32(v3))+64:], uint64(i64(0)))
	store64(m.memory[int64(uint32(v3))+48:], uint64(v1))
	store64(m.memory[int64(uint32(v3))+32:], uint64(v1^i64(8387220255154660723)))
	store64(m.memory[int64(uint32(v3))+24:], uint64(v1^i64(7237128888997146477)))
	store64(m.memory[int64(uint32(v3))+40:], uint64(v0))
	store64(m.memory[int64(uint32(v3))+16:], uint64(v0^i64(0x6c7967656e657261)))
	store64(m.memory[int64(uint32(v3))+8:], uint64(v0^i64(8317987319222330741)))
	t1 := int32(load32(m.memory[int64(uint32(v2))+4:]))
	t2 := int32(load32(m.memory[int64(uint32(v2))+8:]))
	m.fn59(v3+i32(8), t1, t2)
	m.memory[int64(uint32(v3))+76] = byte(i32(255))
	m.fn59(v3+i32(8), v3+i32(76), i32(1))
	t3 := int32(load32(m.memory[int64(uint32(v2))+12:]))
	store32(m.memory[int64(uint32(v3))+76:], uint32(t3))
	m.fn59(v3+i32(8), v3+i32(76), i32(4))
	t4 := int64(load64(m.memory[int64(uint32(v3))+8:]))
	v0 = t4
	t5 := int64(load64(m.memory[int64(uint32(v3))+24:]))
	v1 = t5
	t6 := int64(load32(m.memory[int64(uint32(v3))+64:]))
	v4 = t6
	t7 := int64(load64(m.memory[int64(uint32(v3))+56:]))
	v5 = t7
	t8 := int64(load64(m.memory[int64(uint32(v3))+32:]))
	v6 = t8
	t9 := int64(load64(m.memory[int64(uint32(v3))+16:]))
	v7 = t9
	m.g0 = v3 + i32(80)
	t10 := v6
	v4 = v5 | v4<<56
	v5 = t10 ^ v4
	t11 := i64_rotl(v5, i64(16))
	v5 = v5 + v7
	v6 = t11 ^ v5
	t12 := i64_rotl(v6, i64(21))
	t13 := v6
	v0 = v1 + v0
	v6 = t13 + i64_rotl(v0, i64(32))
	v7 = t12 ^ v6
	t14 := i64_rotl(v7, i64(16))
	t15 := v7
	t16 := v5
	v1 = i64_rotl(v1, i64(13)) ^ v0
	v0 = t16 + v1
	v5 = t15 + (i64_rotl(v0, i64(32)) ^ i64(255))
	v7 = t14 ^ v5
	t17 := i64_rotl(v7, i64(21))
	t18 := v7
	t19 := v6 ^ v4
	v1 = v0 ^ i64_rotl(v1, i64(17))
	v0 = t19 + v1
	v4 = t18 + i64_rotl(v0, i64(32))
	v6 = t17 ^ v4
	t20 := i64_rotl(v6, i64(16))
	t21 := v6
	v1 = v0 ^ i64_rotl(v1, i64(13))
	v0 = v1 + v5
	v5 = t21 + i64_rotl(v0, i64(32))
	v6 = t20 ^ v5
	t22 := i64_rotl(v6, i64(21))
	t23 := v6
	v1 = v0 ^ i64_rotl(v1, i64(17))
	v0 = v1 + v4
	v4 = t23 + i64_rotl(v0, i64(32))
	v6 = t22 ^ v4
	t24 := i64_rotl(v6, i64(16))
	t25 := v6
	v1 = i64_rotl(v1, i64(13)) ^ v0
	v0 = v1 + v5
	v5 = t25 + i64_rotl(v0, i64(32))
	t26 := i64_rotl(t24^v5, i64(21))
	v1 = i64_rotl(v1, i64(17)) ^ v0
	v1 = i64_rotl(v1, i64(13)) ^ (v1 + v4)
	t27 := t26 ^ i64_rotl(v1, i64(17))
	v1 = v1 + v5
	return t27 ^ i64_rotl(v1, i64(32)) ^ v1
}
func (m *Module) fn88(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7, v8, v9, v10 int32
	var v11, v12, v13 int64
	var v14, v15, v16 int32
	var v17, v18 int64
	{
		{
			{
				t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
				v2 = t0
				v3 = v2 + i32(1)
				if v3 == 0 {
					m.fn28(i32(1271248), i32(57), i32(1271276))
					panic("unreachable")
				}
				{
					t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
					t2 := v3
					v4 = t1
					t3 := v4
					v5 = v4 + i32(1)
					v6 = int32(uint32(v5) >> 3)
					v7 = v6 * i32(7)
					p4 := v7
					if uint32(v4) < uint32(i32(8)) {
						p4 = t3
					}
					v8 = p4
					if uint32(t2) <= uint32(int32(uint32(v8)>>1)) {
						{
							if v5 != 0 {
								goto l6
							}
							v3 = i32(0)
							goto l7
						l6:
							t7 := int32(load32(m.memory[uint32(v0):]))
							v8 = t7
							v3 = i32(0)
							{
								{
									t8 := v6
									var p9 int32
									if v5&i32(7) != i32(0) {
										p9 = 1
									}
									v6 = t8 + p9
									if v6 == i32(1) {
										goto l8
									}
									v9 = v6 & i32(1)
									v10 = v6 & i32(0x3ffffffe)
									v3 = i32(0)
								l9:
									{
										v6 = v8 + v3
										t10 := int64(load64(m.memory[uint32(v6):]))
										t11 := v6
										v11 = t10
										store64(m.memory[uint32(t11):], uint64(int64(uint64(v11^i64(-1))>>7)&i64(72340172838076673)+(v11|i64(0x7f7f7f7f7f7f7f7f))))
										v6 = v6 + i32(8)
										t12 := int64(load64(m.memory[uint32(v6):]))
										t13 := v6
										v11 = t12
										store64(m.memory[uint32(t13):], uint64(int64(uint64(v11^i64(-1))>>7)&i64(72340172838076673)+(v11|i64(0x7f7f7f7f7f7f7f7f))))
										v3 = v3 + i32(16)
										v10 = v10 + i32(-2)
										if v10 != 0 {
											goto l9
										}
									}
									if v9 == 0 {
										goto l10
									}
								}
							l8:
								v3 = v8 + v3
								t14 := int64(load64(m.memory[uint32(v3):]))
								t15 := v3
								v11 = t14
								store64(m.memory[uint32(t15):], uint64(int64(uint64(v11^i64(-1))>>7)&i64(72340172838076673)+(v11|i64(0x7f7f7f7f7f7f7f7f))))
							}
						l10:
							{
								if uint32(v5) < uint32(i32(8)) {
									goto l11
								}
								t16 := int64(load64(m.memory[uint32(v8):]))
								store64(m.memory[uint32(v8+v5):], uint64(t16))
								goto l12
							}
						l11:
							if v5 == 0 {
								goto l12
							}
							memory_copy(m.memory, uint32(v8+i32(8)), uint32(v8), uint32(v5))
						l12:
							t17 := int64(load64(m.memory[int64(uint32(v1))+8:]))
							v12 = t17
							t18 := int64(load64(m.memory[uint32(v1):]))
							v13 = t18
							v6 = i32(0)
						l20:
							{
								t19 := v8
								v3 = v6
								v10 = t19 + v3
								t20 := int32(m.memory[uint32(v10)])
								if t20 != i32(128) {
									goto l13
								}
								v6 = v8 - v3<<4
								v1 = v6 + i32(-12)
								v14 = v6 + i32(-16)
								v15 = v8 + (v3^i32(-1))<<4
							l19:
								{
									t21 := int32(load32(m.memory[uint32(v14):]))
									t22 := int32(load32(m.memory[uint32(v1):]))
									t23 := m.fn89(v13, v12, t21, t22)
									t24 := v4
									v9 = int32(t23)
									v6 = t24 & v9
									v5 = v6
									{
										t25 := int64(load64(m.memory[uint32(v8+v6):]))
										v11 = t25 & i64(-0x7f7f7f7f7f7f7f80)
										if v11 != i64(0) {
											goto l14
										}
										v16 = i32(8)
										v5 = v6
									l15:
										{
											v5 = v5 + v16
											v16 = v16 + i32(8)
											t26 := v8
											v5 = v5 & v4
											t27 := int64(load64(m.memory[uint32(t26+v5):]))
											v11 = t27 & i64(-0x7f7f7f7f7f7f7f80)
											if v11 == 0 {
												goto l15
											}
										}
									}
								l14:
									{
										t28 := v8
										v5 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v11))))>>3) + v5) & v4
										t29 := int32(int8(m.memory[uint32(t28+v5)]))
										if t29 < i32(0) {
											goto l16
										}
										t30 := int64(load64(m.memory[uint32(v8):]))
										v5 = int32(uint32(int64(bits.TrailingZeros64(uint64(t30&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
									}
								l16:
									{
										if uint32((v5-v6^(v3-v6))&v4) < uint32(i32(8)) {
											goto l17
										}
										v6 = v8 + v5
										t31 := int32(m.memory[uint32(v6)])
										v16 = t31
										t32 := v6
										v9 = int32(uint32(v9) >> 25)
										m.memory[uint32(t32)] = byte(v9)
										m.memory[uint32(v8+(v5+i32(-8))&v4+i32(8))] = byte(v9)
										v6 = v8 + (v5^i32(-1))<<4
										{
											if v16 != i32(255) {
												t35 := int64(load64(m.memory[uint32(v15):]))
												v11 = t35
												t36 := int64(load64(m.memory[uint32(v6):]))
												store64(m.memory[uint32(v15):], uint64(t36))
												store64(m.memory[uint32(v6):], uint64(v11))
												t37 := int64(load64(m.memory[int64(uint32(v15))+8:]))
												v11 = t37
												t38 := int64(load64(m.memory[int64(uint32(v6))+8:]))
												store64(m.memory[int64(uint32(v15))+8:], uint64(t38))
												store64(m.memory[int64(uint32(v6))+8:], uint64(v11))
												goto l19
											}
											m.memory[uint32(v10)] = byte(i32(255))
											m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(i32(255))
											t33 := int64(load64(m.memory[int64(uint32(v15))+8:]))
											store64(m.memory[int64(uint32(v6))+8:], uint64(t33))
											t34 := int64(load64(m.memory[uint32(v15):]))
											store64(m.memory[uint32(v6):], uint64(t34))
											goto l13
										}
									}
								l17:
								}
								t39 := v10
								v6 = int32(uint32(v9) >> 25)
								m.memory[uint32(t39)] = byte(v6)
								m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(v6)
							}
						l13:
							v6 = v3 + i32(1)
							if v3 != v4 {
								goto l20
							}
							p40 := v7
							if uint32(v4) < uint32(i32(8)) {
								p40 = v4
							}
							v3 = p40
						}
					l7:
						store32(m.memory[int64(uint32(v0))+8:], uint32(v3-v2))
						goto l21
					}
					v8 = v8 + i32(1)
					p5 := v3
					if uint32(v8) > uint32(v3) {
						p5 = v8
					}
					v3 = p5
					if uint32(v3) < uint32(i32(15)) {
						goto l2
					}
					{
						if uint32(v3) > uint32(i32(0x1fffffff)) {
							m.fn28(i32(1271248), i32(57), i32(1271276))
							panic("unreachable")
						}
						t6 := int32(uint32(v3<<3) / uint32(i32(7)))
						v3 = i32_shr_u(i32(-1), int32(bits.LeadingZeros32(uint32(t6+i32(-1)))))
						if uint32(v3) > uint32(i32(0xffffffe)) {
							goto l4
						}
						v3 = v3 + i32(1)
						goto l5
					}
				}
			}
		l2:
			p41 := v3&i32(8) + i32(8)
			if uint32(v3) < uint32(i32(4)) {
				p41 = i32(4)
			}
			v3 = p41
		}
	l5:
		v8 = v3 + i32(8)
		t42 := v8
		v10 = v3 << 4
		v6 = t42 + v10
		if uint32(v6) < uint32(v8) {
			goto l4
		}
		if uint32(v6) > uint32(i32(0x7ffffff8)) {
			goto l4
		}
		{
			t43 := m.fn5(v6)
			v5 = t43
			if v5 != 0 {
				v6 = v5 + v10
				if v8 == 0 {
					goto l23
				}
				memory_fill(m.memory, uint32(v6), i32(255), uint32(v8))
			l23:
				v5 = v3 + i32(-1)
				p44 := int32(uint32(v3)>>3) * i32(7)
				if uint32(v3) < uint32(i32(9)) {
					p44 = v5
				}
				v7 = p44
				t45 := int32(load32(m.memory[uint32(v0):]))
				v9 = t45
				{
					if v2 == 0 {
						goto l24
					}
					t46 := int64(load64(m.memory[uint32(v9):]))
					v11 = (t46 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
					t47 := int64(load64(m.memory[int64(uint32(v1))+8:]))
					v17 = t47
					t48 := int64(load64(m.memory[uint32(v1):]))
					v18 = t48
					v8 = v9
					v1 = v2
					v3 = i32(0)
				l30:
					{
						if v11 != i64(0) {
							goto l25
						}
					l26:
						{
							v3 = v3 + i32(8)
							v8 = v8 + i32(8)
							t49 := int64(load64(m.memory[uint32(v8):]))
							v11 = t49 & i64(-0x7f7f7f7f7f7f7f80)
							if v11 == i64(-0x7f7f7f7f7f7f7f80) {
								goto l26
							}
						}
						v11 = v11 ^ i64(-0x7f7f7f7f7f7f7f80)
					l25:
						{
							t50 := v6
							t51 := v5
							t52 := v18
							t53 := v17
							t54 := v9
							v14 = int32(uint32(int64(bits.TrailingZeros64(uint64(v11))))>>3) + v3
							v10 = t54 - v14<<4
							t55 := int32(load32(m.memory[uint32(v10+i32(-16)):]))
							t56 := int32(load32(m.memory[uint32(v10+i32(-12)):]))
							t57 := m.fn89(t52, t53, t55, t56)
							v15 = int32(t57)
							v10 = t51 & v15
							t58 := int64(load64(m.memory[uint32(t50+v10):]))
							v12 = t58 & i64(-0x7f7f7f7f7f7f7f80)
							if v12 != i64(0) {
								goto l27
							}
							v16 = i32(8)
						l28:
							{
								v10 = v10 + v16
								v16 = v16 + i32(8)
								t59 := v6
								v10 = v10 & v5
								t60 := int64(load64(m.memory[uint32(t59+v10):]))
								v12 = t60 & i64(-0x7f7f7f7f7f7f7f80)
								if v12 == 0 {
									goto l28
								}
							}
						}
					l27:
						v13 = v11 + i64(-1)
						{
							t61 := v6
							v10 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v12))))>>3) + v10) & v5
							t62 := int32(int8(m.memory[uint32(t61+v10)]))
							if t62 < i32(0) {
								goto l29
							}
							t63 := int64(load64(m.memory[uint32(v6):]))
							v10 = int32(uint32(int64(bits.TrailingZeros64(uint64(t63&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
						}
					l29:
						v11 = v13 & v11
						t64 := v6 + v10
						v15 = int32(uint32(v15) >> 25)
						m.memory[uint32(t64)] = byte(v15)
						m.memory[uint32(v6+(v10+i32(-8))&v5+i32(8))] = byte(v15)
						v10 = v6 + (v10^i32(-1))<<4
						t65 := v10
						v14 = v9 + (v14^i32(-1))<<4
						t66 := int64(load64(m.memory[int64(uint32(v14))+8:]))
						store64(m.memory[int64(uint32(t65))+8:], uint64(t66))
						t67 := int64(load64(m.memory[uint32(v14):]))
						store64(m.memory[uint32(v10):], uint64(t67))
						v1 = v1 + i32(-1)
						if v1 != 0 {
							goto l30
						}
					}
				}
			l24:
				store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
				store32(m.memory[uint32(v0):], uint32(v6))
				store32(m.memory[int64(uint32(v0))+8:], uint32(v7-v2))
				if v4 == 0 {
					goto l21
				}
				t68 := v4
				v8 = (v4<<4 + i32(23)) & i32(-16)
				v3 = t68 + v8 + i32(9)
				if v3 == 0 {
					goto l21
				}
				v4 = v9 - v8
				t69 := int32(load32(m.memory[uint32(v4+i32(-4)):]))
				v8 = t69
				v6 = v8 & i32(-8)
				t70 := v6
				v8 = v8 & i32(3)
				p71 := i32(8)
				if v8 != 0 {
					p71 = i32(4)
				}
				if uint32(t70) < uint32(p71+v3) {
					m.fn3(i32(1273840), i32(46), i32(1273888))
					panic("unreachable")
				}
				if v8 == 0 {
					goto l32
				}
				if uint32(v6) > uint32(v3+i32(39)) {
					m.fn3(i32(1273904), i32(46), i32(1273952))
					panic("unreachable")
				}
			l32:
				m.fn1(v4)
				return i32(-1)
			}
			m.fn24(i32(8), v6)
			panic("unreachable")
		}
	}
l21:
	return i32(-1)
l4:
	m.fn28(i32(1271248), i32(57), i32(1271276))
	panic("unreachable")
}
func (m *Module) fn89(v0, v1 int64, v2, v3 int32) int64 {
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
	store32(m.memory[int64(uint32(v4))+76:], uint32(v2))
	m.fn59(v4+i32(8), v4+i32(76), i32(4))
	store32(m.memory[int64(uint32(v4))+76:], uint32(v3))
	m.fn59(v4+i32(8), v4+i32(76), i32(4))
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
func (m *Module) fn90(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7, v8, v9, v10 int32
	var v11, v12, v13 int64
	var v14, v15, v16 int32
	var v17, v18 int64
	{
		{
			{
				t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
				v2 = t0
				v3 = v2 + i32(1)
				if v3 == 0 {
					m.fn28(i32(1271248), i32(57), i32(1271276))
					panic("unreachable")
				}
				{
					t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
					t2 := v3
					v4 = t1
					t3 := v4
					v5 = v4 + i32(1)
					v6 = int32(uint32(v5) >> 3)
					v7 = v6 * i32(7)
					p4 := v7
					if uint32(v4) < uint32(i32(8)) {
						p4 = t3
					}
					v8 = p4
					if uint32(t2) <= uint32(int32(uint32(v8)>>1)) {
						{
							if v5 != 0 {
								goto l6
							}
							v3 = i32(0)
							goto l7
						l6:
							t7 := int32(load32(m.memory[uint32(v0):]))
							v8 = t7
							v3 = i32(0)
							{
								{
									t8 := v6
									var p9 int32
									if v5&i32(7) != i32(0) {
										p9 = 1
									}
									v6 = t8 + p9
									if v6 == i32(1) {
										goto l8
									}
									v9 = v6 & i32(1)
									v10 = v6 & i32(0x3ffffffe)
									v3 = i32(0)
								l9:
									{
										v6 = v8 + v3
										t10 := int64(load64(m.memory[uint32(v6):]))
										t11 := v6
										v11 = t10
										store64(m.memory[uint32(t11):], uint64(int64(uint64(v11^i64(-1))>>7)&i64(72340172838076673)+(v11|i64(0x7f7f7f7f7f7f7f7f))))
										v6 = v6 + i32(8)
										t12 := int64(load64(m.memory[uint32(v6):]))
										t13 := v6
										v11 = t12
										store64(m.memory[uint32(t13):], uint64(int64(uint64(v11^i64(-1))>>7)&i64(72340172838076673)+(v11|i64(0x7f7f7f7f7f7f7f7f))))
										v3 = v3 + i32(16)
										v10 = v10 + i32(-2)
										if v10 != 0 {
											goto l9
										}
									}
									if v9 == 0 {
										goto l10
									}
								}
							l8:
								v3 = v8 + v3
								t14 := int64(load64(m.memory[uint32(v3):]))
								t15 := v3
								v11 = t14
								store64(m.memory[uint32(t15):], uint64(int64(uint64(v11^i64(-1))>>7)&i64(72340172838076673)+(v11|i64(0x7f7f7f7f7f7f7f7f))))
							}
						l10:
							{
								if uint32(v5) < uint32(i32(8)) {
									goto l11
								}
								t16 := int64(load64(m.memory[uint32(v8):]))
								store64(m.memory[uint32(v8+v5):], uint64(t16))
								goto l12
							}
						l11:
							if v5 == 0 {
								goto l12
							}
							memory_copy(m.memory, uint32(v8+i32(8)), uint32(v8), uint32(v5))
						l12:
							t17 := int64(load64(m.memory[int64(uint32(v1))+8:]))
							v12 = t17
							t18 := int64(load64(m.memory[uint32(v1):]))
							v13 = t18
							v6 = i32(0)
						l20:
							{
								t19 := v8
								v3 = v6
								v10 = t19 + v3
								t20 := int32(m.memory[uint32(v10)])
								if t20 != i32(128) {
									goto l13
								}
								v6 = v8 - v3<<3
								v1 = v6 + i32(-4)
								v14 = v6 + i32(-8)
								v15 = v8 + (v3^i32(-1))<<3
							l19:
								{
									t21 := int32(load32(m.memory[uint32(v14):]))
									t22 := int32(load32(m.memory[uint32(v1):]))
									t23 := m.fn89(v13, v12, t21, t22)
									t24 := v4
									v9 = int32(t23)
									v6 = t24 & v9
									v5 = v6
									{
										t25 := int64(load64(m.memory[uint32(v8+v6):]))
										v11 = t25 & i64(-0x7f7f7f7f7f7f7f80)
										if v11 != i64(0) {
											goto l14
										}
										v16 = i32(8)
										v5 = v6
									l15:
										{
											v5 = v5 + v16
											v16 = v16 + i32(8)
											t26 := v8
											v5 = v5 & v4
											t27 := int64(load64(m.memory[uint32(t26+v5):]))
											v11 = t27 & i64(-0x7f7f7f7f7f7f7f80)
											if v11 == 0 {
												goto l15
											}
										}
									}
								l14:
									{
										t28 := v8
										v5 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v11))))>>3) + v5) & v4
										t29 := int32(int8(m.memory[uint32(t28+v5)]))
										if t29 < i32(0) {
											goto l16
										}
										t30 := int64(load64(m.memory[uint32(v8):]))
										v5 = int32(uint32(int64(bits.TrailingZeros64(uint64(t30&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
									}
								l16:
									{
										if uint32((v5-v6^(v3-v6))&v4) < uint32(i32(8)) {
											goto l17
										}
										v6 = v8 + v5
										t31 := int32(m.memory[uint32(v6)])
										v16 = t31
										t32 := v6
										v9 = int32(uint32(v9) >> 25)
										m.memory[uint32(t32)] = byte(v9)
										m.memory[uint32(v8+(v5+i32(-8))&v4+i32(8))] = byte(v9)
										v6 = v8 + (v5^i32(-1))<<3
										{
											if v16 != i32(255) {
												t34 := int64(load64(m.memory[uint32(v15):]))
												v11 = t34
												t35 := int64(load64(m.memory[uint32(v6):]))
												store64(m.memory[uint32(v15):], uint64(t35))
												store64(m.memory[uint32(v6):], uint64(v11))
												goto l19
											}
											m.memory[uint32(v10)] = byte(i32(255))
											m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(i32(255))
											t33 := int64(load64(m.memory[uint32(v15):]))
											store64(m.memory[uint32(v6):], uint64(t33))
											goto l13
										}
									}
								l17:
								}
								t36 := v10
								v6 = int32(uint32(v9) >> 25)
								m.memory[uint32(t36)] = byte(v6)
								m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(v6)
							}
						l13:
							v6 = v3 + i32(1)
							if v3 != v4 {
								goto l20
							}
							p37 := v7
							if uint32(v4) < uint32(i32(8)) {
								p37 = v4
							}
							v3 = p37
						}
					l7:
						store32(m.memory[int64(uint32(v0))+8:], uint32(v3-v2))
						goto l21
					}
					v8 = v8 + i32(1)
					p5 := v3
					if uint32(v8) > uint32(v3) {
						p5 = v8
					}
					v3 = p5
					if uint32(v3) < uint32(i32(15)) {
						goto l2
					}
					{
						if uint32(v3) > uint32(i32(0x1fffffff)) {
							m.fn28(i32(1271248), i32(57), i32(1271276))
							panic("unreachable")
						}
						t6 := int32(uint32(v3<<3) / uint32(i32(7)))
						v3 = i32_shr_u(i32(-1), int32(bits.LeadingZeros32(uint32(t6+i32(-1)))))
						if uint32(v3) > uint32(i32(0x1ffffffe)) {
							goto l4
						}
						v3 = v3 + i32(1)
						goto l5
					}
				}
			}
		l2:
			p38 := v3&i32(8) + i32(8)
			if uint32(v3) < uint32(i32(4)) {
				p38 = i32(4)
			}
			v3 = p38
		}
	l5:
		v8 = v3 + i32(8)
		t39 := v8
		v10 = v3 << 3
		v6 = t39 + v10
		if uint32(v6) < uint32(v8) {
			goto l4
		}
		if uint32(v6) > uint32(i32(0x7ffffff8)) {
			goto l4
		}
		{
			t40 := m.fn5(v6)
			v5 = t40
			if v5 != 0 {
				v6 = v5 + v10
				if v8 == 0 {
					goto l23
				}
				memory_fill(m.memory, uint32(v6), i32(255), uint32(v8))
			l23:
				v5 = v3 + i32(-1)
				p41 := int32(uint32(v3)>>3) * i32(7)
				if uint32(v3) < uint32(i32(9)) {
					p41 = v5
				}
				v7 = p41
				t42 := int32(load32(m.memory[uint32(v0):]))
				v9 = t42
				{
					if v2 == 0 {
						goto l24
					}
					t43 := int64(load64(m.memory[uint32(v9):]))
					v11 = (t43 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
					t44 := int64(load64(m.memory[int64(uint32(v1))+8:]))
					v17 = t44
					t45 := int64(load64(m.memory[uint32(v1):]))
					v18 = t45
					v8 = v9
					v1 = v2
					v3 = i32(0)
				l30:
					{
						if v11 != i64(0) {
							goto l25
						}
					l26:
						{
							v3 = v3 + i32(8)
							v8 = v8 + i32(8)
							t46 := int64(load64(m.memory[uint32(v8):]))
							v11 = t46 & i64(-0x7f7f7f7f7f7f7f80)
							if v11 == i64(-0x7f7f7f7f7f7f7f80) {
								goto l26
							}
						}
						v11 = v11 ^ i64(-0x7f7f7f7f7f7f7f80)
					l25:
						{
							t47 := v6
							t48 := v5
							t49 := v18
							t50 := v17
							t51 := v9
							v14 = int32(uint32(int64(bits.TrailingZeros64(uint64(v11))))>>3) + v3
							v10 = t51 - v14<<3
							t52 := int32(load32(m.memory[uint32(v10+i32(-8)):]))
							t53 := int32(load32(m.memory[uint32(v10+i32(-4)):]))
							t54 := m.fn89(t49, t50, t52, t53)
							v15 = int32(t54)
							v10 = t48 & v15
							t55 := int64(load64(m.memory[uint32(t47+v10):]))
							v12 = t55 & i64(-0x7f7f7f7f7f7f7f80)
							if v12 != i64(0) {
								goto l27
							}
							v16 = i32(8)
						l28:
							{
								v10 = v10 + v16
								v16 = v16 + i32(8)
								t56 := v6
								v10 = v10 & v5
								t57 := int64(load64(m.memory[uint32(t56+v10):]))
								v12 = t57 & i64(-0x7f7f7f7f7f7f7f80)
								if v12 == 0 {
									goto l28
								}
							}
						}
					l27:
						v13 = v11 + i64(-1)
						{
							t58 := v6
							v10 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v12))))>>3) + v10) & v5
							t59 := int32(int8(m.memory[uint32(t58+v10)]))
							if t59 < i32(0) {
								goto l29
							}
							t60 := int64(load64(m.memory[uint32(v6):]))
							v10 = int32(uint32(int64(bits.TrailingZeros64(uint64(t60&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
						}
					l29:
						v11 = v13 & v11
						t61 := v6 + v10
						v15 = int32(uint32(v15) >> 25)
						m.memory[uint32(t61)] = byte(v15)
						m.memory[uint32(v6+(v10+i32(-8))&v5+i32(8))] = byte(v15)
						t62 := int64(load64(m.memory[uint32(v9+(v14^i32(-1))<<3):]))
						store64(m.memory[uint32(v6+(v10^i32(-1))<<3):], uint64(t62))
						v1 = v1 + i32(-1)
						if v1 != 0 {
							goto l30
						}
					}
				}
			l24:
				store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
				store32(m.memory[uint32(v0):], uint32(v6))
				store32(m.memory[int64(uint32(v0))+8:], uint32(v7-v2))
				if v4 == 0 {
					goto l21
				}
				t63 := v4
				v8 = (v4<<3 + i32(15)) & i32(-8)
				v3 = t63 + v8 + i32(9)
				if v3 == 0 {
					goto l21
				}
				v4 = v9 - v8
				t64 := int32(load32(m.memory[uint32(v4+i32(-4)):]))
				v8 = t64
				v6 = v8 & i32(-8)
				t65 := v6
				v8 = v8 & i32(3)
				p66 := i32(8)
				if v8 != 0 {
					p66 = i32(4)
				}
				if uint32(t65) < uint32(p66+v3) {
					m.fn3(i32(1273840), i32(46), i32(1273888))
					panic("unreachable")
				}
				if v8 == 0 {
					goto l32
				}
				if uint32(v6) > uint32(v3+i32(39)) {
					m.fn3(i32(1273904), i32(46), i32(1273952))
					panic("unreachable")
				}
			l32:
				m.fn1(v4)
				return i32(-1)
			}
			m.fn24(i32(8), v6)
			panic("unreachable")
		}
	}
l21:
	return i32(-1)
l4:
	m.fn28(i32(1271248), i32(57), i32(1271276))
	panic("unreachable")
}
func (m *Module) fn91(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7, v8, v9, v10 int32
	var v11, v12, v13 int64
	var v14, v15, v16 int32
	var v17, v18 int64
	{
		{
			{
				t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
				v2 = t0
				v3 = v2 + i32(1)
				if v3 == 0 {
					m.fn28(i32(1271248), i32(57), i32(1271276))
					panic("unreachable")
				}
				{
					t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
					t2 := v3
					v4 = t1
					t3 := v4
					v5 = v4 + i32(1)
					v6 = int32(uint32(v5) >> 3)
					v7 = v6 * i32(7)
					p4 := v7
					if uint32(v4) < uint32(i32(8)) {
						p4 = t3
					}
					v8 = p4
					if uint32(t2) <= uint32(int32(uint32(v8)>>1)) {
						{
							if v5 != 0 {
								goto l6
							}
							v3 = i32(0)
							goto l7
						l6:
							t7 := int32(load32(m.memory[uint32(v0):]))
							v8 = t7
							v3 = i32(0)
							{
								{
									t8 := v6
									var p9 int32
									if v5&i32(7) != i32(0) {
										p9 = 1
									}
									v6 = t8 + p9
									if v6 == i32(1) {
										goto l8
									}
									v9 = v6 & i32(1)
									v10 = v6 & i32(0x3ffffffe)
									v3 = i32(0)
								l9:
									{
										v6 = v8 + v3
										t10 := int64(load64(m.memory[uint32(v6):]))
										t11 := v6
										v11 = t10
										store64(m.memory[uint32(t11):], uint64(int64(uint64(v11^i64(-1))>>7)&i64(72340172838076673)+(v11|i64(0x7f7f7f7f7f7f7f7f))))
										v6 = v6 + i32(8)
										t12 := int64(load64(m.memory[uint32(v6):]))
										t13 := v6
										v11 = t12
										store64(m.memory[uint32(t13):], uint64(int64(uint64(v11^i64(-1))>>7)&i64(72340172838076673)+(v11|i64(0x7f7f7f7f7f7f7f7f))))
										v3 = v3 + i32(16)
										v10 = v10 + i32(-2)
										if v10 != 0 {
											goto l9
										}
									}
									if v9 == 0 {
										goto l10
									}
								}
							l8:
								v3 = v8 + v3
								t14 := int64(load64(m.memory[uint32(v3):]))
								t15 := v3
								v11 = t14
								store64(m.memory[uint32(t15):], uint64(int64(uint64(v11^i64(-1))>>7)&i64(72340172838076673)+(v11|i64(0x7f7f7f7f7f7f7f7f))))
							}
						l10:
							{
								if uint32(v5) < uint32(i32(8)) {
									goto l11
								}
								t16 := int64(load64(m.memory[uint32(v8):]))
								store64(m.memory[uint32(v8+v5):], uint64(t16))
								goto l12
							}
						l11:
							if v5 == 0 {
								goto l12
							}
							memory_copy(m.memory, uint32(v8+i32(8)), uint32(v8), uint32(v5))
						l12:
							t17 := int64(load64(m.memory[int64(uint32(v1))+8:]))
							v12 = t17
							t18 := int64(load64(m.memory[uint32(v1):]))
							v13 = t18
							v6 = i32(0)
						l20:
							{
								t19 := v8
								v3 = v6
								v10 = t19 + v3
								t20 := int32(m.memory[uint32(v10)])
								if t20 != i32(128) {
									goto l13
								}
								v6 = v8 - v3<<3
								v1 = v6 + i32(-4)
								v14 = v6 + i32(-8)
								v15 = v8 + (v3^i32(-1))<<3
							l19:
								{
									t21 := int32(load16(m.memory[uint32(v14):]))
									t22 := int32(load32(m.memory[uint32(v1):]))
									t23 := m.fn92(v13, v12, t21, t22)
									t24 := v4
									v9 = int32(t23)
									v6 = t24 & v9
									v5 = v6
									{
										t25 := int64(load64(m.memory[uint32(v8+v6):]))
										v11 = t25 & i64(-0x7f7f7f7f7f7f7f80)
										if v11 != i64(0) {
											goto l14
										}
										v16 = i32(8)
										v5 = v6
									l15:
										{
											v5 = v5 + v16
											v16 = v16 + i32(8)
											t26 := v8
											v5 = v5 & v4
											t27 := int64(load64(m.memory[uint32(t26+v5):]))
											v11 = t27 & i64(-0x7f7f7f7f7f7f7f80)
											if v11 == 0 {
												goto l15
											}
										}
									}
								l14:
									{
										t28 := v8
										v5 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v11))))>>3) + v5) & v4
										t29 := int32(int8(m.memory[uint32(t28+v5)]))
										if t29 < i32(0) {
											goto l16
										}
										t30 := int64(load64(m.memory[uint32(v8):]))
										v5 = int32(uint32(int64(bits.TrailingZeros64(uint64(t30&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
									}
								l16:
									{
										if uint32((v5-v6^(v3-v6))&v4) < uint32(i32(8)) {
											goto l17
										}
										v6 = v8 + v5
										t31 := int32(m.memory[uint32(v6)])
										v16 = t31
										t32 := v6
										v9 = int32(uint32(v9) >> 25)
										m.memory[uint32(t32)] = byte(v9)
										m.memory[uint32(v8+(v5+i32(-8))&v4+i32(8))] = byte(v9)
										v6 = v8 + (v5^i32(-1))<<3
										{
											if v16 != i32(255) {
												t34 := int64(load64(m.memory[uint32(v15):]))
												v11 = t34
												t35 := int64(load64(m.memory[uint32(v6):]))
												store64(m.memory[uint32(v15):], uint64(t35))
												store64(m.memory[uint32(v6):], uint64(v11))
												goto l19
											}
											m.memory[uint32(v10)] = byte(i32(255))
											m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(i32(255))
											t33 := int64(load64(m.memory[uint32(v15):]))
											store64(m.memory[uint32(v6):], uint64(t33))
											goto l13
										}
									}
								l17:
								}
								t36 := v10
								v6 = int32(uint32(v9) >> 25)
								m.memory[uint32(t36)] = byte(v6)
								m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(v6)
							}
						l13:
							v6 = v3 + i32(1)
							if v3 != v4 {
								goto l20
							}
							p37 := v7
							if uint32(v4) < uint32(i32(8)) {
								p37 = v4
							}
							v3 = p37
						}
					l7:
						store32(m.memory[int64(uint32(v0))+8:], uint32(v3-v2))
						goto l21
					}
					v8 = v8 + i32(1)
					p5 := v3
					if uint32(v8) > uint32(v3) {
						p5 = v8
					}
					v3 = p5
					if uint32(v3) < uint32(i32(15)) {
						goto l2
					}
					{
						if uint32(v3) > uint32(i32(0x1fffffff)) {
							m.fn28(i32(1271248), i32(57), i32(1271276))
							panic("unreachable")
						}
						t6 := int32(uint32(v3<<3) / uint32(i32(7)))
						v3 = i32_shr_u(i32(-1), int32(bits.LeadingZeros32(uint32(t6+i32(-1)))))
						if uint32(v3) > uint32(i32(0x1ffffffe)) {
							goto l4
						}
						v3 = v3 + i32(1)
						goto l5
					}
				}
			}
		l2:
			p38 := v3&i32(8) + i32(8)
			if uint32(v3) < uint32(i32(4)) {
				p38 = i32(4)
			}
			v3 = p38
		}
	l5:
		v8 = v3 + i32(8)
		t39 := v8
		v10 = v3 << 3
		v6 = t39 + v10
		if uint32(v6) < uint32(v8) {
			goto l4
		}
		if uint32(v6) > uint32(i32(0x7ffffff8)) {
			goto l4
		}
		{
			t40 := m.fn5(v6)
			v5 = t40
			if v5 != 0 {
				v6 = v5 + v10
				if v8 == 0 {
					goto l23
				}
				memory_fill(m.memory, uint32(v6), i32(255), uint32(v8))
			l23:
				v5 = v3 + i32(-1)
				p41 := int32(uint32(v3)>>3) * i32(7)
				if uint32(v3) < uint32(i32(9)) {
					p41 = v5
				}
				v7 = p41
				t42 := int32(load32(m.memory[uint32(v0):]))
				v9 = t42
				{
					if v2 == 0 {
						goto l24
					}
					t43 := int64(load64(m.memory[uint32(v9):]))
					v11 = (t43 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
					t44 := int64(load64(m.memory[int64(uint32(v1))+8:]))
					v17 = t44
					t45 := int64(load64(m.memory[uint32(v1):]))
					v18 = t45
					v8 = v9
					v1 = v2
					v3 = i32(0)
				l30:
					{
						if v11 != i64(0) {
							goto l25
						}
					l26:
						{
							v3 = v3 + i32(8)
							v8 = v8 + i32(8)
							t46 := int64(load64(m.memory[uint32(v8):]))
							v11 = t46 & i64(-0x7f7f7f7f7f7f7f80)
							if v11 == i64(-0x7f7f7f7f7f7f7f80) {
								goto l26
							}
						}
						v11 = v11 ^ i64(-0x7f7f7f7f7f7f7f80)
					l25:
						{
							t47 := v6
							t48 := v5
							t49 := v18
							t50 := v17
							t51 := v9
							v14 = int32(uint32(int64(bits.TrailingZeros64(uint64(v11))))>>3) + v3
							v10 = t51 - v14<<3
							t52 := int32(load16(m.memory[uint32(v10+i32(-8)):]))
							t53 := int32(load32(m.memory[uint32(v10+i32(-4)):]))
							t54 := m.fn92(t49, t50, t52, t53)
							v15 = int32(t54)
							v10 = t48 & v15
							t55 := int64(load64(m.memory[uint32(t47+v10):]))
							v12 = t55 & i64(-0x7f7f7f7f7f7f7f80)
							if v12 != i64(0) {
								goto l27
							}
							v16 = i32(8)
						l28:
							{
								v10 = v10 + v16
								v16 = v16 + i32(8)
								t56 := v6
								v10 = v10 & v5
								t57 := int64(load64(m.memory[uint32(t56+v10):]))
								v12 = t57 & i64(-0x7f7f7f7f7f7f7f80)
								if v12 == 0 {
									goto l28
								}
							}
						}
					l27:
						v13 = v11 + i64(-1)
						{
							t58 := v6
							v10 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v12))))>>3) + v10) & v5
							t59 := int32(int8(m.memory[uint32(t58+v10)]))
							if t59 < i32(0) {
								goto l29
							}
							t60 := int64(load64(m.memory[uint32(v6):]))
							v10 = int32(uint32(int64(bits.TrailingZeros64(uint64(t60&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
						}
					l29:
						v11 = v13 & v11
						t61 := v6 + v10
						v15 = int32(uint32(v15) >> 25)
						m.memory[uint32(t61)] = byte(v15)
						m.memory[uint32(v6+(v10+i32(-8))&v5+i32(8))] = byte(v15)
						t62 := int64(load64(m.memory[uint32(v9+(v14^i32(-1))<<3):]))
						store64(m.memory[uint32(v6+(v10^i32(-1))<<3):], uint64(t62))
						v1 = v1 + i32(-1)
						if v1 != 0 {
							goto l30
						}
					}
				}
			l24:
				store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
				store32(m.memory[uint32(v0):], uint32(v6))
				store32(m.memory[int64(uint32(v0))+8:], uint32(v7-v2))
				if v4 == 0 {
					goto l21
				}
				t63 := v4
				v8 = (v4<<3 + i32(15)) & i32(-8)
				v3 = t63 + v8 + i32(9)
				if v3 == 0 {
					goto l21
				}
				v4 = v9 - v8
				t64 := int32(load32(m.memory[uint32(v4+i32(-4)):]))
				v8 = t64
				v6 = v8 & i32(-8)
				t65 := v6
				v8 = v8 & i32(3)
				p66 := i32(8)
				if v8 != 0 {
					p66 = i32(4)
				}
				if uint32(t65) < uint32(p66+v3) {
					m.fn3(i32(1273840), i32(46), i32(1273888))
					panic("unreachable")
				}
				if v8 == 0 {
					goto l32
				}
				if uint32(v6) > uint32(v3+i32(39)) {
					m.fn3(i32(1273904), i32(46), i32(1273952))
					panic("unreachable")
				}
			l32:
				m.fn1(v4)
				return i32(-1)
			}
			m.fn24(i32(8), v6)
			panic("unreachable")
		}
	}
l21:
	return i32(-1)
l4:
	m.fn28(i32(1271248), i32(57), i32(1271276))
	panic("unreachable")
}
func (m *Module) fn92(v0, v1 int64, v2, v3 int32) int64 {
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
	store16(m.memory[int64(uint32(v4))+76:], uint16(v2))
	m.fn59(v4+i32(8), v4+i32(76), i32(2))
	store32(m.memory[int64(uint32(v4))+76:], uint32(v3))
	m.fn59(v4+i32(8), v4+i32(76), i32(4))
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
func (m *Module) fn93(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7, v8, v9, v10 int32
	var v11, v12, v13 int64
	var v14, v15 int32
	var v16, v17 int64
	var v18 int32
	{
		{
			{
				t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
				v2 = t0
				v3 = v2 + i32(1)
				if v3 == 0 {
					m.fn28(i32(1271248), i32(57), i32(1271276))
					panic("unreachable")
				}
				{
					t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
					t2 := v3
					v4 = t1
					t3 := v4
					v5 = v4 + i32(1)
					v6 = int32(uint32(v5) >> 3)
					v7 = v6 * i32(7)
					p4 := v7
					if uint32(v4) < uint32(i32(8)) {
						p4 = t3
					}
					v8 = p4
					if uint32(t2) <= uint32(int32(uint32(v8)>>1)) {
						{
							if v5 != 0 {
								goto l6
							}
							v3 = i32(0)
							goto l7
						l6:
							t7 := int32(load32(m.memory[uint32(v0):]))
							v8 = t7
							v3 = i32(0)
							{
								{
									t8 := v6
									var p9 int32
									if v5&i32(7) != i32(0) {
										p9 = 1
									}
									v6 = t8 + p9
									if v6 == i32(1) {
										goto l8
									}
									v9 = v6 & i32(1)
									v10 = v6 & i32(0x3ffffffe)
									v3 = i32(0)
								l9:
									{
										v6 = v8 + v3
										t10 := int64(load64(m.memory[uint32(v6):]))
										t11 := v6
										v11 = t10
										store64(m.memory[uint32(t11):], uint64(int64(uint64(v11^i64(-1))>>7)&i64(72340172838076673)+(v11|i64(0x7f7f7f7f7f7f7f7f))))
										v6 = v6 + i32(8)
										t12 := int64(load64(m.memory[uint32(v6):]))
										t13 := v6
										v11 = t12
										store64(m.memory[uint32(t13):], uint64(int64(uint64(v11^i64(-1))>>7)&i64(72340172838076673)+(v11|i64(0x7f7f7f7f7f7f7f7f))))
										v3 = v3 + i32(16)
										v10 = v10 + i32(-2)
										if v10 != 0 {
											goto l9
										}
									}
									if v9 == 0 {
										goto l10
									}
								}
							l8:
								v3 = v8 + v3
								t14 := int64(load64(m.memory[uint32(v3):]))
								t15 := v3
								v11 = t14
								store64(m.memory[uint32(t15):], uint64(int64(uint64(v11^i64(-1))>>7)&i64(72340172838076673)+(v11|i64(0x7f7f7f7f7f7f7f7f))))
							}
						l10:
							{
								if uint32(v5) < uint32(i32(8)) {
									goto l11
								}
								t16 := int64(load64(m.memory[uint32(v8):]))
								store64(m.memory[uint32(v8+v5):], uint64(t16))
								goto l12
							}
						l11:
							if v5 == 0 {
								goto l12
							}
							memory_copy(m.memory, uint32(v8+i32(8)), uint32(v8), uint32(v5))
						l12:
							t17 := int64(load64(m.memory[int64(uint32(v1))+8:]))
							v12 = t17
							t18 := int64(load64(m.memory[uint32(v1):]))
							v13 = t18
							v6 = i32(0)
						l20:
							{
								t19 := v8
								v3 = v6
								v10 = t19 + v3
								t20 := int32(m.memory[uint32(v10)])
								if t20 != i32(128) {
									goto l13
								}
								v1 = v8 - v3<<4 + i32(-16)
								v14 = v8 + (v3^i32(-1))<<4
							l19:
								{
									t21 := int32(load32(m.memory[uint32(v1):]))
									t22 := m.fn94(v13, v12, t21)
									t23 := v4
									v9 = int32(t22)
									v6 = t23 & v9
									v5 = v6
									{
										t24 := int64(load64(m.memory[uint32(v8+v6):]))
										v11 = t24 & i64(-0x7f7f7f7f7f7f7f80)
										if v11 != i64(0) {
											goto l14
										}
										v15 = i32(8)
										v5 = v6
									l15:
										{
											v5 = v5 + v15
											v15 = v15 + i32(8)
											t25 := v8
											v5 = v5 & v4
											t26 := int64(load64(m.memory[uint32(t25+v5):]))
											v11 = t26 & i64(-0x7f7f7f7f7f7f7f80)
											if v11 == 0 {
												goto l15
											}
										}
									}
								l14:
									{
										t27 := v8
										v5 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v11))))>>3) + v5) & v4
										t28 := int32(int8(m.memory[uint32(t27+v5)]))
										if t28 < i32(0) {
											goto l16
										}
										t29 := int64(load64(m.memory[uint32(v8):]))
										v5 = int32(uint32(int64(bits.TrailingZeros64(uint64(t29&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
									}
								l16:
									{
										if uint32((v5-v6^(v3-v6))&v4) < uint32(i32(8)) {
											goto l17
										}
										v6 = v8 + v5
										t30 := int32(m.memory[uint32(v6)])
										v15 = t30
										t31 := v6
										v9 = int32(uint32(v9) >> 25)
										m.memory[uint32(t31)] = byte(v9)
										m.memory[uint32(v8+(v5+i32(-8))&v4+i32(8))] = byte(v9)
										v6 = v8 + (v5^i32(-1))<<4
										{
											if v15 != i32(255) {
												t34 := int64(load64(m.memory[uint32(v14):]))
												v11 = t34
												t35 := int64(load64(m.memory[uint32(v6):]))
												store64(m.memory[uint32(v14):], uint64(t35))
												store64(m.memory[uint32(v6):], uint64(v11))
												t36 := int64(load64(m.memory[int64(uint32(v14))+8:]))
												v11 = t36
												t37 := int64(load64(m.memory[int64(uint32(v6))+8:]))
												store64(m.memory[int64(uint32(v14))+8:], uint64(t37))
												store64(m.memory[int64(uint32(v6))+8:], uint64(v11))
												goto l19
											}
											m.memory[uint32(v10)] = byte(i32(255))
											m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(i32(255))
											t32 := int64(load64(m.memory[int64(uint32(v14))+8:]))
											store64(m.memory[int64(uint32(v6))+8:], uint64(t32))
											t33 := int64(load64(m.memory[uint32(v14):]))
											store64(m.memory[uint32(v6):], uint64(t33))
											goto l13
										}
									}
								l17:
								}
								t38 := v10
								v6 = int32(uint32(v9) >> 25)
								m.memory[uint32(t38)] = byte(v6)
								m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(v6)
							}
						l13:
							v6 = v3 + i32(1)
							if v3 != v4 {
								goto l20
							}
							p39 := v7
							if uint32(v4) < uint32(i32(8)) {
								p39 = v4
							}
							v3 = p39
						}
					l7:
						store32(m.memory[int64(uint32(v0))+8:], uint32(v3-v2))
						goto l21
					}
					v8 = v8 + i32(1)
					p5 := v3
					if uint32(v8) > uint32(v3) {
						p5 = v8
					}
					v3 = p5
					if uint32(v3) < uint32(i32(15)) {
						goto l2
					}
					{
						if uint32(v3) > uint32(i32(0x1fffffff)) {
							m.fn28(i32(1271248), i32(57), i32(1271276))
							panic("unreachable")
						}
						t6 := int32(uint32(v3<<3) / uint32(i32(7)))
						v3 = i32_shr_u(i32(-1), int32(bits.LeadingZeros32(uint32(t6+i32(-1)))))
						if uint32(v3) > uint32(i32(0xffffffe)) {
							goto l4
						}
						v3 = v3 + i32(1)
						goto l5
					}
				}
			}
		l2:
			p40 := v3&i32(8) + i32(8)
			if uint32(v3) < uint32(i32(4)) {
				p40 = i32(4)
			}
			v3 = p40
		}
	l5:
		v8 = v3 + i32(8)
		t41 := v8
		v10 = v3 << 4
		v6 = t41 + v10
		if uint32(v6) < uint32(v8) {
			goto l4
		}
		if uint32(v6) > uint32(i32(0x7ffffff8)) {
			goto l4
		}
		{
			t42 := m.fn5(v6)
			v5 = t42
			if v5 != 0 {
				v6 = v5 + v10
				if v8 == 0 {
					goto l23
				}
				memory_fill(m.memory, uint32(v6), i32(255), uint32(v8))
			l23:
				v5 = v3 + i32(-1)
				p43 := int32(uint32(v3)>>3) * i32(7)
				if uint32(v3) < uint32(i32(9)) {
					p43 = v5
				}
				v15 = p43
				t44 := int32(load32(m.memory[uint32(v0):]))
				v9 = t44
				{
					if v2 == 0 {
						goto l24
					}
					t45 := int64(load64(m.memory[uint32(v9):]))
					v11 = (t45 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
					t46 := int64(load64(m.memory[int64(uint32(v1))+8:]))
					v16 = t46
					t47 := int64(load64(m.memory[uint32(v1):]))
					v17 = t47
					v8 = v9
					v1 = v2
					v3 = i32(0)
				l30:
					{
						if v11 != i64(0) {
							goto l25
						}
					l26:
						{
							v3 = v3 + i32(8)
							v8 = v8 + i32(8)
							t48 := int64(load64(m.memory[uint32(v8):]))
							v11 = t48 & i64(-0x7f7f7f7f7f7f7f80)
							if v11 == i64(-0x7f7f7f7f7f7f7f80) {
								goto l26
							}
						}
						v11 = v11 ^ i64(-0x7f7f7f7f7f7f7f80)
					l25:
						{
							t49 := v6
							t50 := v5
							t51 := v17
							t52 := v16
							t53 := v9
							v14 = int32(uint32(int64(bits.TrailingZeros64(uint64(v11))))>>3) + v3
							t54 := int32(load32(m.memory[uint32(t53-v14<<4+i32(-16)):]))
							t55 := m.fn94(t51, t52, t54)
							v7 = int32(t55)
							v10 = t50 & v7
							t56 := int64(load64(m.memory[uint32(t49+v10):]))
							v12 = t56 & i64(-0x7f7f7f7f7f7f7f80)
							if v12 != i64(0) {
								goto l27
							}
							v18 = i32(8)
						l28:
							{
								v10 = v10 + v18
								v18 = v18 + i32(8)
								t57 := v6
								v10 = v10 & v5
								t58 := int64(load64(m.memory[uint32(t57+v10):]))
								v12 = t58 & i64(-0x7f7f7f7f7f7f7f80)
								if v12 == 0 {
									goto l28
								}
							}
						}
					l27:
						v13 = v11 + i64(-1)
						{
							t59 := v6
							v10 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v12))))>>3) + v10) & v5
							t60 := int32(int8(m.memory[uint32(t59+v10)]))
							if t60 < i32(0) {
								goto l29
							}
							t61 := int64(load64(m.memory[uint32(v6):]))
							v10 = int32(uint32(int64(bits.TrailingZeros64(uint64(t61&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
						}
					l29:
						v11 = v13 & v11
						t62 := v6 + v10
						v7 = int32(uint32(v7) >> 25)
						m.memory[uint32(t62)] = byte(v7)
						m.memory[uint32(v6+(v10+i32(-8))&v5+i32(8))] = byte(v7)
						v10 = v6 + (v10^i32(-1))<<4
						t63 := v10
						v14 = v9 + (v14^i32(-1))<<4
						t64 := int64(load64(m.memory[int64(uint32(v14))+8:]))
						store64(m.memory[int64(uint32(t63))+8:], uint64(t64))
						t65 := int64(load64(m.memory[uint32(v14):]))
						store64(m.memory[uint32(v10):], uint64(t65))
						v1 = v1 + i32(-1)
						if v1 != 0 {
							goto l30
						}
					}
				}
			l24:
				store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
				store32(m.memory[uint32(v0):], uint32(v6))
				store32(m.memory[int64(uint32(v0))+8:], uint32(v15-v2))
				if v4 == 0 {
					goto l21
				}
				t66 := v4
				v8 = (v4<<4 + i32(23)) & i32(-16)
				v3 = t66 + v8 + i32(9)
				if v3 == 0 {
					goto l21
				}
				v4 = v9 - v8
				t67 := int32(load32(m.memory[uint32(v4+i32(-4)):]))
				v8 = t67
				v6 = v8 & i32(-8)
				t68 := v6
				v8 = v8 & i32(3)
				p69 := i32(8)
				if v8 != 0 {
					p69 = i32(4)
				}
				if uint32(t68) < uint32(p69+v3) {
					m.fn3(i32(1273840), i32(46), i32(1273888))
					panic("unreachable")
				}
				if v8 == 0 {
					goto l32
				}
				if uint32(v6) > uint32(v3+i32(39)) {
					m.fn3(i32(1273904), i32(46), i32(1273952))
					panic("unreachable")
				}
			l32:
				m.fn1(v4)
				return i32(-1)
			}
			m.fn24(i32(8), v6)
			panic("unreachable")
		}
	}
l21:
	return i32(-1)
l4:
	m.fn28(i32(1271248), i32(57), i32(1271276))
	panic("unreachable")
}
func (m *Module) fn94(v0, v1 int64, v2 int32) int64 {
	var v3 int32
	var v4, v5, v6, v7 int64
	t0 := m.g0
	v3 = t0 - i32(80)
	m.g0 = v3
	store64(m.memory[int64(uint32(v3))+56:], uint64(i64(0)))
	store64(m.memory[int64(uint32(v3))+64:], uint64(i64(0)))
	store64(m.memory[int64(uint32(v3))+48:], uint64(v1))
	store64(m.memory[int64(uint32(v3))+32:], uint64(v1^i64(8387220255154660723)))
	store64(m.memory[int64(uint32(v3))+24:], uint64(v1^i64(7237128888997146477)))
	store64(m.memory[int64(uint32(v3))+40:], uint64(v0))
	store64(m.memory[int64(uint32(v3))+16:], uint64(v0^i64(0x6c7967656e657261)))
	store64(m.memory[int64(uint32(v3))+8:], uint64(v0^i64(8317987319222330741)))
	store32(m.memory[int64(uint32(v3))+76:], uint32(v2))
	m.fn59(v3+i32(8), v3+i32(76), i32(4))
	t1 := int64(load64(m.memory[int64(uint32(v3))+8:]))
	v0 = t1
	t2 := int64(load64(m.memory[int64(uint32(v3))+24:]))
	v1 = t2
	t3 := int64(load32(m.memory[int64(uint32(v3))+64:]))
	v4 = t3
	t4 := int64(load64(m.memory[int64(uint32(v3))+56:]))
	v5 = t4
	t5 := int64(load64(m.memory[int64(uint32(v3))+32:]))
	v6 = t5
	t6 := int64(load64(m.memory[int64(uint32(v3))+16:]))
	v7 = t6
	m.g0 = v3 + i32(80)
	t7 := v6
	v4 = v5 | v4<<56
	v5 = t7 ^ v4
	t8 := i64_rotl(v5, i64(16))
	v5 = v5 + v7
	v6 = t8 ^ v5
	t9 := i64_rotl(v6, i64(21))
	t10 := v6
	v0 = v1 + v0
	v6 = t10 + i64_rotl(v0, i64(32))
	v7 = t9 ^ v6
	t11 := i64_rotl(v7, i64(16))
	t12 := v7
	t13 := v5
	v1 = i64_rotl(v1, i64(13)) ^ v0
	v0 = t13 + v1
	v5 = t12 + (i64_rotl(v0, i64(32)) ^ i64(255))
	v7 = t11 ^ v5
	t14 := i64_rotl(v7, i64(21))
	t15 := v7
	t16 := v6 ^ v4
	v1 = v0 ^ i64_rotl(v1, i64(17))
	v0 = t16 + v1
	v4 = t15 + i64_rotl(v0, i64(32))
	v6 = t14 ^ v4
	t17 := i64_rotl(v6, i64(16))
	t18 := v6
	v1 = v0 ^ i64_rotl(v1, i64(13))
	v0 = v1 + v5
	v5 = t18 + i64_rotl(v0, i64(32))
	v6 = t17 ^ v5
	t19 := i64_rotl(v6, i64(21))
	t20 := v6
	v1 = v0 ^ i64_rotl(v1, i64(17))
	v0 = v1 + v4
	v4 = t20 + i64_rotl(v0, i64(32))
	v6 = t19 ^ v4
	t21 := i64_rotl(v6, i64(16))
	t22 := v6
	v1 = i64_rotl(v1, i64(13)) ^ v0
	v0 = v1 + v5
	v5 = t22 + i64_rotl(v0, i64(32))
	t23 := i64_rotl(t21^v5, i64(21))
	v1 = i64_rotl(v1, i64(17)) ^ v0
	v1 = i64_rotl(v1, i64(13)) ^ (v1 + v4)
	t24 := t23 ^ i64_rotl(v1, i64(17))
	v1 = v1 + v5
	return t24 ^ i64_rotl(v1, i64(32)) ^ v1
}
func (m *Module) fn95(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7, v8, v9, v10 int32
	var v11, v12, v13 int64
	var v14, v15 int32
	var v16, v17 int64
	var v18 int32
	{
		{
			{
				t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
				v2 = t0
				v3 = v2 + i32(1)
				if v3 == 0 {
					m.fn28(i32(1271248), i32(57), i32(1271276))
					panic("unreachable")
				}
				{
					t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
					t2 := v3
					v4 = t1
					t3 := v4
					v5 = v4 + i32(1)
					v6 = int32(uint32(v5) >> 3)
					v7 = v6 * i32(7)
					p4 := v7
					if uint32(v4) < uint32(i32(8)) {
						p4 = t3
					}
					v8 = p4
					if uint32(t2) <= uint32(int32(uint32(v8)>>1)) {
						{
							if v5 != 0 {
								goto l6
							}
							v3 = i32(0)
							goto l7
						l6:
							t7 := int32(load32(m.memory[uint32(v0):]))
							v8 = t7
							v3 = i32(0)
							{
								{
									t8 := v6
									var p9 int32
									if v5&i32(7) != i32(0) {
										p9 = 1
									}
									v6 = t8 + p9
									if v6 == i32(1) {
										goto l8
									}
									v9 = v6 & i32(1)
									v10 = v6 & i32(0x3ffffffe)
									v3 = i32(0)
								l9:
									{
										v6 = v8 + v3
										t10 := int64(load64(m.memory[uint32(v6):]))
										t11 := v6
										v11 = t10
										store64(m.memory[uint32(t11):], uint64(int64(uint64(v11^i64(-1))>>7)&i64(72340172838076673)+(v11|i64(0x7f7f7f7f7f7f7f7f))))
										v6 = v6 + i32(8)
										t12 := int64(load64(m.memory[uint32(v6):]))
										t13 := v6
										v11 = t12
										store64(m.memory[uint32(t13):], uint64(int64(uint64(v11^i64(-1))>>7)&i64(72340172838076673)+(v11|i64(0x7f7f7f7f7f7f7f7f))))
										v3 = v3 + i32(16)
										v10 = v10 + i32(-2)
										if v10 != 0 {
											goto l9
										}
									}
									if v9 == 0 {
										goto l10
									}
								}
							l8:
								v3 = v8 + v3
								t14 := int64(load64(m.memory[uint32(v3):]))
								t15 := v3
								v11 = t14
								store64(m.memory[uint32(t15):], uint64(int64(uint64(v11^i64(-1))>>7)&i64(72340172838076673)+(v11|i64(0x7f7f7f7f7f7f7f7f))))
							}
						l10:
							{
								if uint32(v5) < uint32(i32(8)) {
									goto l11
								}
								t16 := int64(load64(m.memory[uint32(v8):]))
								store64(m.memory[uint32(v8+v5):], uint64(t16))
								goto l12
							}
						l11:
							if v5 == 0 {
								goto l12
							}
							memory_copy(m.memory, uint32(v8+i32(8)), uint32(v8), uint32(v5))
						l12:
							t17 := int64(load64(m.memory[int64(uint32(v1))+8:]))
							v12 = t17
							t18 := int64(load64(m.memory[uint32(v1):]))
							v13 = t18
							v6 = i32(0)
						l20:
							{
								t19 := v8
								v3 = v6
								v10 = t19 + v3
								t20 := int32(m.memory[uint32(v10)])
								if t20 != i32(128) {
									goto l13
								}
								v1 = v8 - v3<<4 + i32(-16)
								v14 = v8 + (v3^i32(-1))<<4
							l19:
								{
									t21 := int32(load32(m.memory[uint32(v1):]))
									t22 := m.fn94(v13, v12, t21)
									t23 := v4
									v9 = int32(t22)
									v6 = t23 & v9
									v5 = v6
									{
										t24 := int64(load64(m.memory[uint32(v8+v6):]))
										v11 = t24 & i64(-0x7f7f7f7f7f7f7f80)
										if v11 != i64(0) {
											goto l14
										}
										v15 = i32(8)
										v5 = v6
									l15:
										{
											v5 = v5 + v15
											v15 = v15 + i32(8)
											t25 := v8
											v5 = v5 & v4
											t26 := int64(load64(m.memory[uint32(t25+v5):]))
											v11 = t26 & i64(-0x7f7f7f7f7f7f7f80)
											if v11 == 0 {
												goto l15
											}
										}
									}
								l14:
									{
										t27 := v8
										v5 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v11))))>>3) + v5) & v4
										t28 := int32(int8(m.memory[uint32(t27+v5)]))
										if t28 < i32(0) {
											goto l16
										}
										t29 := int64(load64(m.memory[uint32(v8):]))
										v5 = int32(uint32(int64(bits.TrailingZeros64(uint64(t29&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
									}
								l16:
									{
										if uint32((v5-v6^(v3-v6))&v4) < uint32(i32(8)) {
											goto l17
										}
										v6 = v8 + v5
										t30 := int32(m.memory[uint32(v6)])
										v15 = t30
										t31 := v6
										v9 = int32(uint32(v9) >> 25)
										m.memory[uint32(t31)] = byte(v9)
										m.memory[uint32(v8+(v5+i32(-8))&v4+i32(8))] = byte(v9)
										v6 = v8 + (v5^i32(-1))<<4
										{
											if v15 != i32(255) {
												t34 := int64(load64(m.memory[uint32(v14):]))
												v11 = t34
												t35 := int64(load64(m.memory[uint32(v6):]))
												store64(m.memory[uint32(v14):], uint64(t35))
												store64(m.memory[uint32(v6):], uint64(v11))
												t36 := int64(load64(m.memory[int64(uint32(v14))+8:]))
												v11 = t36
												t37 := int64(load64(m.memory[int64(uint32(v6))+8:]))
												store64(m.memory[int64(uint32(v14))+8:], uint64(t37))
												store64(m.memory[int64(uint32(v6))+8:], uint64(v11))
												goto l19
											}
											m.memory[uint32(v10)] = byte(i32(255))
											m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(i32(255))
											t32 := int64(load64(m.memory[int64(uint32(v14))+8:]))
											store64(m.memory[int64(uint32(v6))+8:], uint64(t32))
											t33 := int64(load64(m.memory[uint32(v14):]))
											store64(m.memory[uint32(v6):], uint64(t33))
											goto l13
										}
									}
								l17:
								}
								t38 := v10
								v6 = int32(uint32(v9) >> 25)
								m.memory[uint32(t38)] = byte(v6)
								m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(v6)
							}
						l13:
							v6 = v3 + i32(1)
							if v3 != v4 {
								goto l20
							}
							p39 := v7
							if uint32(v4) < uint32(i32(8)) {
								p39 = v4
							}
							v3 = p39
						}
					l7:
						store32(m.memory[int64(uint32(v0))+8:], uint32(v3-v2))
						goto l21
					}
					v8 = v8 + i32(1)
					p5 := v3
					if uint32(v8) > uint32(v3) {
						p5 = v8
					}
					v3 = p5
					if uint32(v3) < uint32(i32(15)) {
						goto l2
					}
					{
						if uint32(v3) > uint32(i32(0x1fffffff)) {
							m.fn28(i32(1271248), i32(57), i32(1271276))
							panic("unreachable")
						}
						t6 := int32(uint32(v3<<3) / uint32(i32(7)))
						v3 = i32_shr_u(i32(-1), int32(bits.LeadingZeros32(uint32(t6+i32(-1)))))
						if uint32(v3) > uint32(i32(0xffffffe)) {
							goto l4
						}
						v3 = v3 + i32(1)
						goto l5
					}
				}
			}
		l2:
			p40 := v3&i32(8) + i32(8)
			if uint32(v3) < uint32(i32(4)) {
				p40 = i32(4)
			}
			v3 = p40
		}
	l5:
		v8 = v3 + i32(8)
		t41 := v8
		v10 = v3 << 4
		v6 = t41 + v10
		if uint32(v6) < uint32(v8) {
			goto l4
		}
		if uint32(v6) > uint32(i32(0x7ffffff8)) {
			goto l4
		}
		{
			t42 := m.fn5(v6)
			v5 = t42
			if v5 != 0 {
				v6 = v5 + v10
				if v8 == 0 {
					goto l23
				}
				memory_fill(m.memory, uint32(v6), i32(255), uint32(v8))
			l23:
				v5 = v3 + i32(-1)
				p43 := int32(uint32(v3)>>3) * i32(7)
				if uint32(v3) < uint32(i32(9)) {
					p43 = v5
				}
				v15 = p43
				t44 := int32(load32(m.memory[uint32(v0):]))
				v9 = t44
				{
					if v2 == 0 {
						goto l24
					}
					t45 := int64(load64(m.memory[uint32(v9):]))
					v11 = (t45 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
					t46 := int64(load64(m.memory[int64(uint32(v1))+8:]))
					v16 = t46
					t47 := int64(load64(m.memory[uint32(v1):]))
					v17 = t47
					v8 = v9
					v1 = v2
					v3 = i32(0)
				l30:
					{
						if v11 != i64(0) {
							goto l25
						}
					l26:
						{
							v3 = v3 + i32(8)
							v8 = v8 + i32(8)
							t48 := int64(load64(m.memory[uint32(v8):]))
							v11 = t48 & i64(-0x7f7f7f7f7f7f7f80)
							if v11 == i64(-0x7f7f7f7f7f7f7f80) {
								goto l26
							}
						}
						v11 = v11 ^ i64(-0x7f7f7f7f7f7f7f80)
					l25:
						{
							t49 := v6
							t50 := v5
							t51 := v17
							t52 := v16
							t53 := v9
							v14 = int32(uint32(int64(bits.TrailingZeros64(uint64(v11))))>>3) + v3
							t54 := int32(load32(m.memory[uint32(t53-v14<<4+i32(-16)):]))
							t55 := m.fn94(t51, t52, t54)
							v7 = int32(t55)
							v10 = t50 & v7
							t56 := int64(load64(m.memory[uint32(t49+v10):]))
							v12 = t56 & i64(-0x7f7f7f7f7f7f7f80)
							if v12 != i64(0) {
								goto l27
							}
							v18 = i32(8)
						l28:
							{
								v10 = v10 + v18
								v18 = v18 + i32(8)
								t57 := v6
								v10 = v10 & v5
								t58 := int64(load64(m.memory[uint32(t57+v10):]))
								v12 = t58 & i64(-0x7f7f7f7f7f7f7f80)
								if v12 == 0 {
									goto l28
								}
							}
						}
					l27:
						v13 = v11 + i64(-1)
						{
							t59 := v6
							v10 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v12))))>>3) + v10) & v5
							t60 := int32(int8(m.memory[uint32(t59+v10)]))
							if t60 < i32(0) {
								goto l29
							}
							t61 := int64(load64(m.memory[uint32(v6):]))
							v10 = int32(uint32(int64(bits.TrailingZeros64(uint64(t61&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
						}
					l29:
						v11 = v13 & v11
						t62 := v6 + v10
						v7 = int32(uint32(v7) >> 25)
						m.memory[uint32(t62)] = byte(v7)
						m.memory[uint32(v6+(v10+i32(-8))&v5+i32(8))] = byte(v7)
						v10 = v6 + (v10^i32(-1))<<4
						t63 := v10
						v14 = v9 + (v14^i32(-1))<<4
						t64 := int64(load64(m.memory[int64(uint32(v14))+8:]))
						store64(m.memory[int64(uint32(t63))+8:], uint64(t64))
						t65 := int64(load64(m.memory[uint32(v14):]))
						store64(m.memory[uint32(v10):], uint64(t65))
						v1 = v1 + i32(-1)
						if v1 != 0 {
							goto l30
						}
					}
				}
			l24:
				store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
				store32(m.memory[uint32(v0):], uint32(v6))
				store32(m.memory[int64(uint32(v0))+8:], uint32(v15-v2))
				if v4 == 0 {
					goto l21
				}
				t66 := v4
				v8 = (v4<<4 + i32(23)) & i32(-16)
				v3 = t66 + v8 + i32(9)
				if v3 == 0 {
					goto l21
				}
				v4 = v9 - v8
				t67 := int32(load32(m.memory[uint32(v4+i32(-4)):]))
				v8 = t67
				v6 = v8 & i32(-8)
				t68 := v6
				v8 = v8 & i32(3)
				p69 := i32(8)
				if v8 != 0 {
					p69 = i32(4)
				}
				if uint32(t68) < uint32(p69+v3) {
					m.fn3(i32(1273840), i32(46), i32(1273888))
					panic("unreachable")
				}
				if v8 == 0 {
					goto l32
				}
				if uint32(v6) > uint32(v3+i32(39)) {
					m.fn3(i32(1273904), i32(46), i32(1273952))
					panic("unreachable")
				}
			l32:
				m.fn1(v4)
				return i32(-1)
			}
			m.fn24(i32(8), v6)
			panic("unreachable")
		}
	}
l21:
	return i32(-1)
l4:
	m.fn28(i32(1271248), i32(57), i32(1271276))
	panic("unreachable")
}
func (m *Module) fn96(v0, v1, v2 int32) int32 {
	var v3, v4, v5, v6, v7, v8 int32
	var v9 int64
	var v10, v11, v12 int32
	var v13, v14 int64
	var v15 int32
	var v16 int64
	var v17 int32
	var v18 int64
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
		v3 = t0
		v1 = v3 + v1
		if uint32(v1) < uint32(v3) {
			m.fn28(i32(1271248), i32(57), i32(1271276))
			panic("unreachable")
		}
		{
			t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t2 := v1
			v4 = t1
			t3 := v4
			v5 = v4 + i32(1)
			v6 = int32(uint32(v5) >> 3)
			v7 = v6 * i32(7)
			p4 := v7
			if uint32(v4) < uint32(i32(8)) {
				p4 = t3
			}
			v8 = p4
			if uint32(t2) <= uint32(int32(uint32(v8)>>1)) {
				goto l1
			}
			{
				{
					{
						v8 = v8 + i32(1)
						p5 := v1
						if uint32(v8) > uint32(v1) {
							p5 = v8
						}
						v1 = p5
						if uint32(v1) < uint32(i32(15)) {
							goto l2
						}
						if uint32(v1) > uint32(i32(0x1fffffff)) {
							m.fn28(i32(1271248), i32(57), i32(1271276))
							panic("unreachable")
						}
						t6 := int32(uint32(v1<<3) / uint32(i32(7)))
						v1 = i32_shr_u(i32(-1), int32(bits.LeadingZeros32(uint32(t6+i32(-1))))) + i32(1)
						goto l4
					}
				l2:
					p7 := v1&i32(8) + i32(8)
					if uint32(v1) < uint32(i32(4)) {
						p7 = i32(4)
					}
					v1 = p7
				}
			l4:
				v9 = int64(uint32(v1)) * i64(12)
				if int32(int64(uint64(v9)>>32)) != 0 {
					goto l5
				}
				v6 = int32(v9)
				if uint32(v6) > uint32(i32(-8)) {
					goto l5
				}
				v8 = v1 + i32(8)
				t8 := v8
				v10 = (v6 + i32(7)) & i32(-8)
				v6 = t8 + v10
				if uint32(v6) < uint32(v8) {
					goto l5
				}
				if uint32(v6) > uint32(i32(0x7ffffff8)) {
					goto l5
				}
				t9 := m.fn5(v6)
				v5 = t9
				if v5 != 0 {
					v6 = v5 + v10
					if v8 == 0 {
						goto l7
					}
					memory_fill(m.memory, uint32(v6), i32(255), uint32(v8))
				l7:
					v5 = v1 + i32(-1)
					p10 := int32(uint32(v1)>>3) * i32(7)
					if uint32(v1) < uint32(i32(9)) {
						p10 = v5
					}
					v11 = p10
					t11 := int32(load32(m.memory[uint32(v0):]))
					v12 = t11
					{
						if v3 == 0 {
							goto l8
						}
						t12 := int64(load64(m.memory[uint32(v12):]))
						v9 = (t12 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
						t13 := int64(load64(m.memory[int64(uint32(v2))+8:]))
						v13 = t13
						t14 := int64(load64(m.memory[uint32(v2):]))
						v14 = t14
						v8 = v12
						v2 = v3
						v1 = i32(0)
					l14:
						{
							if v9 != i64(0) {
								goto l9
							}
						l10:
							{
								v1 = v1 + i32(8)
								v8 = v8 + i32(8)
								t15 := int64(load64(m.memory[uint32(v8):]))
								v9 = t15 & i64(-0x7f7f7f7f7f7f7f80)
								if v9 == i64(-0x7f7f7f7f7f7f7f80) {
									goto l10
								}
							}
							v9 = v9 ^ i64(-0x7f7f7f7f7f7f7f80)
						l9:
							{
								t16 := v6
								t17 := v5
								t18 := v14
								t19 := v13
								t20 := v12
								v15 = int32(uint32(int64(bits.TrailingZeros64(uint64(v9))))>>3) + v1
								t21 := int32(load32(m.memory[uint32(t20+(i32(0)-v15)*i32(12)+i32(-12)):]))
								t22 := m.fn94(t18, t19, t21)
								v7 = int32(t22)
								v10 = t17 & v7
								t23 := int64(load64(m.memory[uint32(t16+v10):]))
								v16 = t23 & i64(-0x7f7f7f7f7f7f7f80)
								if v16 != i64(0) {
									goto l11
								}
								v17 = i32(8)
							l12:
								{
									v10 = v10 + v17
									v17 = v17 + i32(8)
									t24 := v6
									v10 = v10 & v5
									t25 := int64(load64(m.memory[uint32(t24+v10):]))
									v16 = t25 & i64(-0x7f7f7f7f7f7f7f80)
									if v16 == 0 {
										goto l12
									}
								}
							}
						l11:
							v18 = v9 + i64(-1)
							{
								t26 := v6
								v10 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v16))))>>3) + v10) & v5
								t27 := int32(int8(m.memory[uint32(t26+v10)]))
								if t27 < i32(0) {
									goto l13
								}
								t28 := int64(load64(m.memory[uint32(v6):]))
								v10 = int32(uint32(int64(bits.TrailingZeros64(uint64(t28&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
							}
						l13:
							v9 = v18 & v9
							t29 := v6 + v10
							v7 = int32(uint32(v7) >> 25)
							m.memory[uint32(t29)] = byte(v7)
							m.memory[uint32(v6+(v10+i32(-8))&v5+i32(8))] = byte(v7)
							v10 = v6 + (v10^i32(-1))*i32(12)
							t30 := v10
							v15 = v12 + (v15^i32(-1))*i32(12)
							t31 := int32(load32(m.memory[int64(uint32(v15))+8:]))
							store32(m.memory[int64(uint32(t30))+8:], uint32(t31))
							t32 := int64(load64(m.memory[uint32(v15):]))
							store64(m.memory[uint32(v10):], uint64(t32))
							v2 = v2 + i32(-1)
							if v2 != 0 {
								goto l14
							}
						}
					}
				l8:
					store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
					store32(m.memory[uint32(v0):], uint32(v6))
					store32(m.memory[int64(uint32(v0))+8:], uint32(v11-v3))
					if v4 == 0 {
						goto l15
					}
					t33 := v4
					v8 = (v4*i32(12) + i32(19)) & i32(-8)
					v1 = t33 + v8 + i32(9)
					if v1 == 0 {
						goto l15
					}
					v4 = v12 - v8
					t34 := int32(load32(m.memory[uint32(v4+i32(-4)):]))
					v8 = t34
					v6 = v8 & i32(-8)
					t35 := v6
					v8 = v8 & i32(3)
					p36 := i32(8)
					if v8 != 0 {
						p36 = i32(4)
					}
					if uint32(t35) < uint32(p36+v1) {
						m.fn3(i32(1273840), i32(46), i32(1273888))
						panic("unreachable")
					}
					if v8 == 0 {
						goto l17
					}
					if uint32(v6) > uint32(v1+i32(39)) {
						m.fn3(i32(1273904), i32(46), i32(1273952))
						panic("unreachable")
					}
				l17:
					m.fn1(v4)
					return i32(-1)
				}
				m.fn24(i32(8), v6)
				panic("unreachable")
			}
		}
	l1:
		{
			if v5 != 0 {
				goto l19
			}
			v1 = i32(0)
			goto l20
		l19:
			t37 := int32(load32(m.memory[uint32(v0):]))
			v8 = t37
			v1 = i32(0)
			{
				{
					t38 := v6
					var p39 int32
					if v5&i32(7) != i32(0) {
						p39 = 1
					}
					v6 = t38 + p39
					if v6 == i32(1) {
						goto l21
					}
					v12 = v6 & i32(1)
					v10 = v6 & i32(0x3ffffffe)
					v1 = i32(0)
				l22:
					{
						v6 = v8 + v1
						t40 := int64(load64(m.memory[uint32(v6):]))
						t41 := v6
						v9 = t40
						store64(m.memory[uint32(t41):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
						v6 = v6 + i32(8)
						t42 := int64(load64(m.memory[uint32(v6):]))
						t43 := v6
						v9 = t42
						store64(m.memory[uint32(t43):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
						v1 = v1 + i32(16)
						v10 = v10 + i32(-2)
						if v10 != 0 {
							goto l22
						}
					}
					if v12 == 0 {
						goto l23
					}
				}
			l21:
				v1 = v8 + v1
				t44 := int64(load64(m.memory[uint32(v1):]))
				t45 := v1
				v9 = t44
				store64(m.memory[uint32(t45):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
			}
		l23:
			{
				if uint32(v5) < uint32(i32(8)) {
					goto l24
				}
				t46 := int64(load64(m.memory[uint32(v8):]))
				store64(m.memory[uint32(v8+v5):], uint64(t46))
				goto l25
			}
		l24:
			if v5 == 0 {
				goto l25
			}
			memory_copy(m.memory, uint32(v8+i32(8)), uint32(v8), uint32(v5))
		l25:
			t47 := int64(load64(m.memory[int64(uint32(v2))+8:]))
			v16 = t47
			t48 := int64(load64(m.memory[uint32(v2):]))
			v18 = t48
			v6 = i32(0)
		l33:
			{
				t49 := v8
				v1 = v6
				v10 = t49 + v1
				t50 := int32(m.memory[uint32(v10)])
				if t50 != i32(128) {
					goto l26
				}
				v15 = v8 + (v1^i32(-1))*i32(12)
				v2 = v8 + (i32(0)-v1)*i32(12) + i32(-12)
			l32:
				{
					t51 := int32(load32(m.memory[uint32(v2):]))
					t52 := m.fn94(v18, v16, t51)
					t53 := v4
					v12 = int32(t52)
					v6 = t53 & v12
					v5 = v6
					{
						t54 := int64(load64(m.memory[uint32(v8+v6):]))
						v9 = t54 & i64(-0x7f7f7f7f7f7f7f80)
						if v9 != i64(0) {
							goto l27
						}
						v11 = i32(8)
						v5 = v6
					l28:
						{
							v5 = v5 + v11
							v11 = v11 + i32(8)
							t55 := v8
							v5 = v5 & v4
							t56 := int64(load64(m.memory[uint32(t55+v5):]))
							v9 = t56 & i64(-0x7f7f7f7f7f7f7f80)
							if v9 == 0 {
								goto l28
							}
						}
					}
				l27:
					{
						t57 := v8
						v5 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v9))))>>3) + v5) & v4
						t58 := int32(int8(m.memory[uint32(t57+v5)]))
						if t58 < i32(0) {
							goto l29
						}
						t59 := int64(load64(m.memory[uint32(v8):]))
						v5 = int32(uint32(int64(bits.TrailingZeros64(uint64(t59&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
					}
				l29:
					{
						if uint32((v5-v6^(v1-v6))&v4) < uint32(i32(8)) {
							goto l30
						}
						v6 = v8 + v5
						t60 := int32(m.memory[uint32(v6)])
						v11 = t60
						t61 := v6
						v12 = int32(uint32(v12) >> 25)
						m.memory[uint32(t61)] = byte(v12)
						m.memory[uint32(v8+(v5+i32(-8))&v4+i32(8))] = byte(v12)
						v6 = v8 + (v5^i32(-1))*i32(12)
						{
							if v11 != i32(255) {
								t64 := int32(load32(m.memory[uint32(v15):]))
								v5 = t64
								t65 := int32(load32(m.memory[uint32(v6):]))
								store32(m.memory[uint32(v15):], uint32(t65))
								store32(m.memory[uint32(v6):], uint32(v5))
								t66 := int32(load32(m.memory[int64(uint32(v6))+4:]))
								v5 = t66
								t67 := int32(load32(m.memory[int64(uint32(v15))+4:]))
								store32(m.memory[int64(uint32(v6))+4:], uint32(t67))
								store32(m.memory[int64(uint32(v15))+4:], uint32(v5))
								t68 := int32(load32(m.memory[int64(uint32(v15))+8:]))
								v5 = t68
								t69 := int32(load32(m.memory[int64(uint32(v6))+8:]))
								store32(m.memory[int64(uint32(v15))+8:], uint32(t69))
								store32(m.memory[int64(uint32(v6))+8:], uint32(v5))
								goto l32
							}
							m.memory[uint32(v10)] = byte(i32(255))
							m.memory[uint32(v8+v4&(v1+i32(-8))+i32(8))] = byte(i32(255))
							t62 := int32(load32(m.memory[int64(uint32(v15))+8:]))
							store32(m.memory[int64(uint32(v6))+8:], uint32(t62))
							t63 := int64(load64(m.memory[uint32(v15):]))
							store64(m.memory[uint32(v6):], uint64(t63))
							goto l26
						}
					}
				l30:
				}
				t70 := v10
				v6 = int32(uint32(v12) >> 25)
				m.memory[uint32(t70)] = byte(v6)
				m.memory[uint32(v8+v4&(v1+i32(-8))+i32(8))] = byte(v6)
			}
		l26:
			v6 = v1 + i32(1)
			if v1 != v4 {
				goto l33
			}
			p71 := v7
			if uint32(v4) < uint32(i32(8)) {
				p71 = v4
			}
			v1 = p71
		}
	l20:
		store32(m.memory[int64(uint32(v0))+8:], uint32(v1-v3))
	l15:
		return i32(-1)
	}
l5:
	m.fn28(i32(1271248), i32(57), i32(1271276))
	panic("unreachable")
}
func (m *Module) fn97(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7, v8, v9, v10 int32
	var v11, v12, v13 int64
	var v14, v15 int32
	var v16, v17 int64
	var v18 int32
	{
		{
			{
				t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
				v2 = t0
				v3 = v2 + i32(1)
				if v3 == 0 {
					m.fn28(i32(1271248), i32(57), i32(1271276))
					panic("unreachable")
				}
				{
					t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
					t2 := v3
					v4 = t1
					t3 := v4
					v5 = v4 + i32(1)
					v6 = int32(uint32(v5) >> 3)
					v7 = v6 * i32(7)
					p4 := v7
					if uint32(v4) < uint32(i32(8)) {
						p4 = t3
					}
					v8 = p4
					if uint32(t2) <= uint32(int32(uint32(v8)>>1)) {
						{
							if v5 != 0 {
								goto l6
							}
							v3 = i32(0)
							goto l7
						l6:
							t7 := int32(load32(m.memory[uint32(v0):]))
							v8 = t7
							v3 = i32(0)
							{
								{
									t8 := v6
									var p9 int32
									if v5&i32(7) != i32(0) {
										p9 = 1
									}
									v6 = t8 + p9
									if v6 == i32(1) {
										goto l8
									}
									v9 = v6 & i32(1)
									v10 = v6 & i32(0x3ffffffe)
									v3 = i32(0)
								l9:
									{
										v6 = v8 + v3
										t10 := int64(load64(m.memory[uint32(v6):]))
										t11 := v6
										v11 = t10
										store64(m.memory[uint32(t11):], uint64(int64(uint64(v11^i64(-1))>>7)&i64(72340172838076673)+(v11|i64(0x7f7f7f7f7f7f7f7f))))
										v6 = v6 + i32(8)
										t12 := int64(load64(m.memory[uint32(v6):]))
										t13 := v6
										v11 = t12
										store64(m.memory[uint32(t13):], uint64(int64(uint64(v11^i64(-1))>>7)&i64(72340172838076673)+(v11|i64(0x7f7f7f7f7f7f7f7f))))
										v3 = v3 + i32(16)
										v10 = v10 + i32(-2)
										if v10 != 0 {
											goto l9
										}
									}
									if v9 == 0 {
										goto l10
									}
								}
							l8:
								v3 = v8 + v3
								t14 := int64(load64(m.memory[uint32(v3):]))
								t15 := v3
								v11 = t14
								store64(m.memory[uint32(t15):], uint64(int64(uint64(v11^i64(-1))>>7)&i64(72340172838076673)+(v11|i64(0x7f7f7f7f7f7f7f7f))))
							}
						l10:
							{
								if uint32(v5) < uint32(i32(8)) {
									goto l11
								}
								t16 := int64(load64(m.memory[uint32(v8):]))
								store64(m.memory[uint32(v8+v5):], uint64(t16))
								goto l12
							}
						l11:
							if v5 == 0 {
								goto l12
							}
							memory_copy(m.memory, uint32(v8+i32(8)), uint32(v8), uint32(v5))
						l12:
							t17 := int64(load64(m.memory[int64(uint32(v1))+8:]))
							v12 = t17
							t18 := int64(load64(m.memory[uint32(v1):]))
							v13 = t18
							v6 = i32(0)
						l20:
							{
								t19 := v8
								v3 = v6
								v10 = t19 + v3
								t20 := int32(m.memory[uint32(v10)])
								if t20 != i32(128) {
									goto l13
								}
								v1 = v8 - v3<<3 + i32(-8)
								v14 = v8 + (v3^i32(-1))<<3
							l19:
								{
									t21 := int32(load32(m.memory[uint32(v1):]))
									t22 := m.fn94(v13, v12, t21)
									t23 := v4
									v9 = int32(t22)
									v6 = t23 & v9
									v5 = v6
									{
										t24 := int64(load64(m.memory[uint32(v8+v6):]))
										v11 = t24 & i64(-0x7f7f7f7f7f7f7f80)
										if v11 != i64(0) {
											goto l14
										}
										v15 = i32(8)
										v5 = v6
									l15:
										{
											v5 = v5 + v15
											v15 = v15 + i32(8)
											t25 := v8
											v5 = v5 & v4
											t26 := int64(load64(m.memory[uint32(t25+v5):]))
											v11 = t26 & i64(-0x7f7f7f7f7f7f7f80)
											if v11 == 0 {
												goto l15
											}
										}
									}
								l14:
									{
										t27 := v8
										v5 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v11))))>>3) + v5) & v4
										t28 := int32(int8(m.memory[uint32(t27+v5)]))
										if t28 < i32(0) {
											goto l16
										}
										t29 := int64(load64(m.memory[uint32(v8):]))
										v5 = int32(uint32(int64(bits.TrailingZeros64(uint64(t29&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
									}
								l16:
									{
										if uint32((v5-v6^(v3-v6))&v4) < uint32(i32(8)) {
											goto l17
										}
										v6 = v8 + v5
										t30 := int32(m.memory[uint32(v6)])
										v15 = t30
										t31 := v6
										v9 = int32(uint32(v9) >> 25)
										m.memory[uint32(t31)] = byte(v9)
										m.memory[uint32(v8+(v5+i32(-8))&v4+i32(8))] = byte(v9)
										v6 = v8 + (v5^i32(-1))<<3
										{
											if v15 != i32(255) {
												t33 := int64(load64(m.memory[uint32(v14):]))
												v11 = t33
												t34 := int64(load64(m.memory[uint32(v6):]))
												store64(m.memory[uint32(v14):], uint64(t34))
												store64(m.memory[uint32(v6):], uint64(v11))
												goto l19
											}
											m.memory[uint32(v10)] = byte(i32(255))
											m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(i32(255))
											t32 := int64(load64(m.memory[uint32(v14):]))
											store64(m.memory[uint32(v6):], uint64(t32))
											goto l13
										}
									}
								l17:
								}
								t35 := v10
								v6 = int32(uint32(v9) >> 25)
								m.memory[uint32(t35)] = byte(v6)
								m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(v6)
							}
						l13:
							v6 = v3 + i32(1)
							if v3 != v4 {
								goto l20
							}
							p36 := v7
							if uint32(v4) < uint32(i32(8)) {
								p36 = v4
							}
							v3 = p36
						}
					l7:
						store32(m.memory[int64(uint32(v0))+8:], uint32(v3-v2))
						goto l21
					}
					v8 = v8 + i32(1)
					p5 := v3
					if uint32(v8) > uint32(v3) {
						p5 = v8
					}
					v3 = p5
					if uint32(v3) < uint32(i32(15)) {
						goto l2
					}
					{
						if uint32(v3) > uint32(i32(0x1fffffff)) {
							m.fn28(i32(1271248), i32(57), i32(1271276))
							panic("unreachable")
						}
						t6 := int32(uint32(v3<<3) / uint32(i32(7)))
						v3 = i32_shr_u(i32(-1), int32(bits.LeadingZeros32(uint32(t6+i32(-1)))))
						if uint32(v3) > uint32(i32(0x1ffffffe)) {
							goto l4
						}
						v3 = v3 + i32(1)
						goto l5
					}
				}
			}
		l2:
			p37 := v3&i32(8) + i32(8)
			if uint32(v3) < uint32(i32(4)) {
				p37 = i32(4)
			}
			v3 = p37
		}
	l5:
		v8 = v3 + i32(8)
		t38 := v8
		v10 = v3 << 3
		v6 = t38 + v10
		if uint32(v6) < uint32(v8) {
			goto l4
		}
		if uint32(v6) > uint32(i32(0x7ffffff8)) {
			goto l4
		}
		{
			t39 := m.fn5(v6)
			v5 = t39
			if v5 != 0 {
				v6 = v5 + v10
				if v8 == 0 {
					goto l23
				}
				memory_fill(m.memory, uint32(v6), i32(255), uint32(v8))
			l23:
				v5 = v3 + i32(-1)
				p40 := int32(uint32(v3)>>3) * i32(7)
				if uint32(v3) < uint32(i32(9)) {
					p40 = v5
				}
				v15 = p40
				t41 := int32(load32(m.memory[uint32(v0):]))
				v9 = t41
				{
					if v2 == 0 {
						goto l24
					}
					t42 := int64(load64(m.memory[uint32(v9):]))
					v11 = (t42 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
					t43 := int64(load64(m.memory[int64(uint32(v1))+8:]))
					v16 = t43
					t44 := int64(load64(m.memory[uint32(v1):]))
					v17 = t44
					v8 = v9
					v1 = v2
					v3 = i32(0)
				l30:
					{
						if v11 != i64(0) {
							goto l25
						}
					l26:
						{
							v3 = v3 + i32(8)
							v8 = v8 + i32(8)
							t45 := int64(load64(m.memory[uint32(v8):]))
							v11 = t45 & i64(-0x7f7f7f7f7f7f7f80)
							if v11 == i64(-0x7f7f7f7f7f7f7f80) {
								goto l26
							}
						}
						v11 = v11 ^ i64(-0x7f7f7f7f7f7f7f80)
					l25:
						{
							t46 := v6
							t47 := v5
							t48 := v17
							t49 := v16
							t50 := v9
							v14 = int32(uint32(int64(bits.TrailingZeros64(uint64(v11))))>>3) + v3
							t51 := int32(load32(m.memory[uint32(t50-v14<<3+i32(-8)):]))
							t52 := m.fn94(t48, t49, t51)
							v7 = int32(t52)
							v10 = t47 & v7
							t53 := int64(load64(m.memory[uint32(t46+v10):]))
							v12 = t53 & i64(-0x7f7f7f7f7f7f7f80)
							if v12 != i64(0) {
								goto l27
							}
							v18 = i32(8)
						l28:
							{
								v10 = v10 + v18
								v18 = v18 + i32(8)
								t54 := v6
								v10 = v10 & v5
								t55 := int64(load64(m.memory[uint32(t54+v10):]))
								v12 = t55 & i64(-0x7f7f7f7f7f7f7f80)
								if v12 == 0 {
									goto l28
								}
							}
						}
					l27:
						v13 = v11 + i64(-1)
						{
							t56 := v6
							v10 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v12))))>>3) + v10) & v5
							t57 := int32(int8(m.memory[uint32(t56+v10)]))
							if t57 < i32(0) {
								goto l29
							}
							t58 := int64(load64(m.memory[uint32(v6):]))
							v10 = int32(uint32(int64(bits.TrailingZeros64(uint64(t58&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
						}
					l29:
						v11 = v13 & v11
						t59 := v6 + v10
						v7 = int32(uint32(v7) >> 25)
						m.memory[uint32(t59)] = byte(v7)
						m.memory[uint32(v6+(v10+i32(-8))&v5+i32(8))] = byte(v7)
						t60 := int64(load64(m.memory[uint32(v9+(v14^i32(-1))<<3):]))
						store64(m.memory[uint32(v6+(v10^i32(-1))<<3):], uint64(t60))
						v1 = v1 + i32(-1)
						if v1 != 0 {
							goto l30
						}
					}
				}
			l24:
				store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
				store32(m.memory[uint32(v0):], uint32(v6))
				store32(m.memory[int64(uint32(v0))+8:], uint32(v15-v2))
				if v4 == 0 {
					goto l21
				}
				t61 := v4
				v8 = (v4<<3 + i32(15)) & i32(-8)
				v3 = t61 + v8 + i32(9)
				if v3 == 0 {
					goto l21
				}
				v4 = v9 - v8
				t62 := int32(load32(m.memory[uint32(v4+i32(-4)):]))
				v8 = t62
				v6 = v8 & i32(-8)
				t63 := v6
				v8 = v8 & i32(3)
				p64 := i32(8)
				if v8 != 0 {
					p64 = i32(4)
				}
				if uint32(t63) < uint32(p64+v3) {
					m.fn3(i32(1273840), i32(46), i32(1273888))
					panic("unreachable")
				}
				if v8 == 0 {
					goto l32
				}
				if uint32(v6) > uint32(v3+i32(39)) {
					m.fn3(i32(1273904), i32(46), i32(1273952))
					panic("unreachable")
				}
			l32:
				m.fn1(v4)
				return i32(-1)
			}
			m.fn24(i32(8), v6)
			panic("unreachable")
		}
	}
l21:
	return i32(-1)
l4:
	m.fn28(i32(1271248), i32(57), i32(1271276))
	panic("unreachable")
}
func (m *Module) fn98(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7, v8, v9, v10 int32
	var v11, v12, v13 int64
	var v14, v15 int32
	var v16, v17 int64
	var v18 int32
	{
		{
			{
				t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
				v2 = t0
				v3 = v2 + i32(1)
				if v3 == 0 {
					m.fn28(i32(1271248), i32(57), i32(1271276))
					panic("unreachable")
				}
				{
					t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
					t2 := v3
					v4 = t1
					t3 := v4
					v5 = v4 + i32(1)
					v6 = int32(uint32(v5) >> 3)
					v7 = v6 * i32(7)
					p4 := v7
					if uint32(v4) < uint32(i32(8)) {
						p4 = t3
					}
					v8 = p4
					if uint32(t2) <= uint32(int32(uint32(v8)>>1)) {
						{
							if v5 != 0 {
								goto l6
							}
							v3 = i32(0)
							goto l7
						l6:
							t7 := int32(load32(m.memory[uint32(v0):]))
							v8 = t7
							v3 = i32(0)
							{
								{
									t8 := v6
									var p9 int32
									if v5&i32(7) != i32(0) {
										p9 = 1
									}
									v6 = t8 + p9
									if v6 == i32(1) {
										goto l8
									}
									v9 = v6 & i32(1)
									v10 = v6 & i32(0x3ffffffe)
									v3 = i32(0)
								l9:
									{
										v6 = v8 + v3
										t10 := int64(load64(m.memory[uint32(v6):]))
										t11 := v6
										v11 = t10
										store64(m.memory[uint32(t11):], uint64(int64(uint64(v11^i64(-1))>>7)&i64(72340172838076673)+(v11|i64(0x7f7f7f7f7f7f7f7f))))
										v6 = v6 + i32(8)
										t12 := int64(load64(m.memory[uint32(v6):]))
										t13 := v6
										v11 = t12
										store64(m.memory[uint32(t13):], uint64(int64(uint64(v11^i64(-1))>>7)&i64(72340172838076673)+(v11|i64(0x7f7f7f7f7f7f7f7f))))
										v3 = v3 + i32(16)
										v10 = v10 + i32(-2)
										if v10 != 0 {
											goto l9
										}
									}
									if v9 == 0 {
										goto l10
									}
								}
							l8:
								v3 = v8 + v3
								t14 := int64(load64(m.memory[uint32(v3):]))
								t15 := v3
								v11 = t14
								store64(m.memory[uint32(t15):], uint64(int64(uint64(v11^i64(-1))>>7)&i64(72340172838076673)+(v11|i64(0x7f7f7f7f7f7f7f7f))))
							}
						l10:
							{
								if uint32(v5) < uint32(i32(8)) {
									goto l11
								}
								t16 := int64(load64(m.memory[uint32(v8):]))
								store64(m.memory[uint32(v8+v5):], uint64(t16))
								goto l12
							}
						l11:
							if v5 == 0 {
								goto l12
							}
							memory_copy(m.memory, uint32(v8+i32(8)), uint32(v8), uint32(v5))
						l12:
							t17 := int64(load64(m.memory[int64(uint32(v1))+8:]))
							v12 = t17
							t18 := int64(load64(m.memory[uint32(v1):]))
							v13 = t18
							v6 = i32(0)
						l20:
							{
								t19 := v8
								v3 = v6
								v10 = t19 + v3
								t20 := int32(m.memory[uint32(v10)])
								if t20 != i32(128) {
									goto l13
								}
								v1 = v8 - v3<<2 + i32(-4)
								v14 = v8 + (v3^i32(-1))<<2
							l19:
								{
									t21 := int32(load32(m.memory[uint32(v1):]))
									t22 := m.fn94(v13, v12, t21)
									t23 := v4
									v9 = int32(t22)
									v6 = t23 & v9
									v5 = v6
									{
										t24 := int64(load64(m.memory[uint32(v8+v6):]))
										v11 = t24 & i64(-0x7f7f7f7f7f7f7f80)
										if v11 != i64(0) {
											goto l14
										}
										v15 = i32(8)
										v5 = v6
									l15:
										{
											v5 = v5 + v15
											v15 = v15 + i32(8)
											t25 := v8
											v5 = v5 & v4
											t26 := int64(load64(m.memory[uint32(t25+v5):]))
											v11 = t26 & i64(-0x7f7f7f7f7f7f7f80)
											if v11 == 0 {
												goto l15
											}
										}
									}
								l14:
									{
										t27 := v8
										v5 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v11))))>>3) + v5) & v4
										t28 := int32(int8(m.memory[uint32(t27+v5)]))
										if t28 < i32(0) {
											goto l16
										}
										t29 := int64(load64(m.memory[uint32(v8):]))
										v5 = int32(uint32(int64(bits.TrailingZeros64(uint64(t29&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
									}
								l16:
									{
										if uint32((v5-v6^(v3-v6))&v4) < uint32(i32(8)) {
											goto l17
										}
										v6 = v8 + v5
										t30 := int32(m.memory[uint32(v6)])
										v15 = t30
										t31 := v6
										v9 = int32(uint32(v9) >> 25)
										m.memory[uint32(t31)] = byte(v9)
										m.memory[uint32(v8+(v5+i32(-8))&v4+i32(8))] = byte(v9)
										v6 = v8 + (v5^i32(-1))<<2
										{
											if v15 != i32(255) {
												t33 := int32(load32(m.memory[uint32(v14):]))
												v5 = t33
												t34 := int32(load32(m.memory[uint32(v6):]))
												store32(m.memory[uint32(v14):], uint32(t34))
												store32(m.memory[uint32(v6):], uint32(v5))
												goto l19
											}
											m.memory[uint32(v10)] = byte(i32(255))
											m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(i32(255))
											t32 := int32(load32(m.memory[uint32(v14):]))
											store32(m.memory[uint32(v6):], uint32(t32))
											goto l13
										}
									}
								l17:
								}
								t35 := v10
								v6 = int32(uint32(v9) >> 25)
								m.memory[uint32(t35)] = byte(v6)
								m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(v6)
							}
						l13:
							v6 = v3 + i32(1)
							if v3 != v4 {
								goto l20
							}
							p36 := v7
							if uint32(v4) < uint32(i32(8)) {
								p36 = v4
							}
							v3 = p36
						}
					l7:
						store32(m.memory[int64(uint32(v0))+8:], uint32(v3-v2))
						goto l21
					}
					v8 = v8 + i32(1)
					p5 := v3
					if uint32(v8) > uint32(v3) {
						p5 = v8
					}
					v3 = p5
					if uint32(v3) < uint32(i32(15)) {
						goto l2
					}
					{
						if uint32(v3) > uint32(i32(0x1fffffff)) {
							m.fn28(i32(1271248), i32(57), i32(1271276))
							panic("unreachable")
						}
						t6 := int32(uint32(v3<<3) / uint32(i32(7)))
						v3 = i32_shr_u(i32(-1), int32(bits.LeadingZeros32(uint32(t6+i32(-1)))))
						if uint32(v3) > uint32(i32(0x3ffffffe)) {
							goto l4
						}
						v3 = v3 + i32(1)
						goto l5
					}
				}
			}
		l2:
			p37 := v3&i32(8) + i32(8)
			if uint32(v3) < uint32(i32(4)) {
				p37 = i32(4)
			}
			v3 = p37
		}
	l5:
		v8 = v3 + i32(8)
		t38 := v8
		v10 = (v3<<2 + i32(7)) & i32(-8)
		v6 = t38 + v10
		if uint32(v6) < uint32(v8) {
			goto l4
		}
		if uint32(v6) > uint32(i32(0x7ffffff8)) {
			goto l4
		}
		{
			t39 := m.fn5(v6)
			v5 = t39
			if v5 != 0 {
				v6 = v5 + v10
				if v8 == 0 {
					goto l23
				}
				memory_fill(m.memory, uint32(v6), i32(255), uint32(v8))
			l23:
				v5 = v3 + i32(-1)
				p40 := int32(uint32(v3)>>3) * i32(7)
				if uint32(v3) < uint32(i32(9)) {
					p40 = v5
				}
				v15 = p40
				t41 := int32(load32(m.memory[uint32(v0):]))
				v9 = t41
				{
					if v2 == 0 {
						goto l24
					}
					t42 := int64(load64(m.memory[uint32(v9):]))
					v11 = (t42 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
					t43 := int64(load64(m.memory[int64(uint32(v1))+8:]))
					v16 = t43
					t44 := int64(load64(m.memory[uint32(v1):]))
					v17 = t44
					v8 = v9
					v1 = v2
					v3 = i32(0)
				l30:
					{
						if v11 != i64(0) {
							goto l25
						}
					l26:
						{
							v3 = v3 + i32(8)
							v8 = v8 + i32(8)
							t45 := int64(load64(m.memory[uint32(v8):]))
							v11 = t45 & i64(-0x7f7f7f7f7f7f7f80)
							if v11 == i64(-0x7f7f7f7f7f7f7f80) {
								goto l26
							}
						}
						v11 = v11 ^ i64(-0x7f7f7f7f7f7f7f80)
					l25:
						{
							t46 := v6
							t47 := v5
							t48 := v17
							t49 := v16
							t50 := v9
							v14 = int32(uint32(int64(bits.TrailingZeros64(uint64(v11))))>>3) + v3
							t51 := int32(load32(m.memory[uint32(t50-v14<<2+i32(-4)):]))
							t52 := m.fn94(t48, t49, t51)
							v7 = int32(t52)
							v10 = t47 & v7
							t53 := int64(load64(m.memory[uint32(t46+v10):]))
							v12 = t53 & i64(-0x7f7f7f7f7f7f7f80)
							if v12 != i64(0) {
								goto l27
							}
							v18 = i32(8)
						l28:
							{
								v10 = v10 + v18
								v18 = v18 + i32(8)
								t54 := v6
								v10 = v10 & v5
								t55 := int64(load64(m.memory[uint32(t54+v10):]))
								v12 = t55 & i64(-0x7f7f7f7f7f7f7f80)
								if v12 == 0 {
									goto l28
								}
							}
						}
					l27:
						v13 = v11 + i64(-1)
						{
							t56 := v6
							v10 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v12))))>>3) + v10) & v5
							t57 := int32(int8(m.memory[uint32(t56+v10)]))
							if t57 < i32(0) {
								goto l29
							}
							t58 := int64(load64(m.memory[uint32(v6):]))
							v10 = int32(uint32(int64(bits.TrailingZeros64(uint64(t58&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
						}
					l29:
						v11 = v13 & v11
						t59 := v6 + v10
						v7 = int32(uint32(v7) >> 25)
						m.memory[uint32(t59)] = byte(v7)
						m.memory[uint32(v6+(v10+i32(-8))&v5+i32(8))] = byte(v7)
						t60 := int32(load32(m.memory[uint32(v9+(v14^i32(-1))<<2):]))
						store32(m.memory[uint32(v6+(v10^i32(-1))<<2):], uint32(t60))
						v1 = v1 + i32(-1)
						if v1 != 0 {
							goto l30
						}
					}
				}
			l24:
				store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
				store32(m.memory[uint32(v0):], uint32(v6))
				store32(m.memory[int64(uint32(v0))+8:], uint32(v15-v2))
				if v4 == 0 {
					goto l21
				}
				t61 := v4
				v8 = (v4<<2 + i32(11)) & i32(-8)
				v3 = t61 + v8 + i32(9)
				if v3 == 0 {
					goto l21
				}
				v4 = v9 - v8
				t62 := int32(load32(m.memory[uint32(v4+i32(-4)):]))
				v8 = t62
				v6 = v8 & i32(-8)
				t63 := v6
				v8 = v8 & i32(3)
				p64 := i32(8)
				if v8 != 0 {
					p64 = i32(4)
				}
				if uint32(t63) < uint32(p64+v3) {
					m.fn3(i32(1273840), i32(46), i32(1273888))
					panic("unreachable")
				}
				if v8 == 0 {
					goto l32
				}
				if uint32(v6) > uint32(v3+i32(39)) {
					m.fn3(i32(1273904), i32(46), i32(1273952))
					panic("unreachable")
				}
			l32:
				m.fn1(v4)
				return i32(-1)
			}
			m.fn24(i32(8), v6)
			panic("unreachable")
		}
	}
l21:
	return i32(-1)
l4:
	m.fn28(i32(1271248), i32(57), i32(1271276))
	panic("unreachable")
}
func (m *Module) fn99(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7, v8 int32
	var v9 int64
	var v10, v11 int32
	var v12, v13 int64
	var v14, v15 int32
	var v16 int64
	var v17 int32
	var v18 int64
	var v19 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
		v2 = t0
		v3 = v2 + i32(1)
		if v3 == 0 {
			m.fn28(i32(1271248), i32(57), i32(1271276))
			panic("unreachable")
		}
		{
			t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t2 := v3
			v4 = t1
			t3 := v4
			v5 = v4 + i32(1)
			v6 = int32(uint32(v5) >> 3)
			v7 = v6 * i32(7)
			p4 := v7
			if uint32(v4) < uint32(i32(8)) {
				p4 = t3
			}
			v8 = p4
			if uint32(t2) <= uint32(int32(uint32(v8)>>1)) {
				goto l1
			}
			{
				{
					{
						v8 = v8 + i32(1)
						p5 := v3
						if uint32(v8) > uint32(v3) {
							p5 = v8
						}
						v3 = p5
						if uint32(v3) < uint32(i32(15)) {
							goto l2
						}
						if uint32(v3) > uint32(i32(0x1fffffff)) {
							m.fn28(i32(1271248), i32(57), i32(1271276))
							panic("unreachable")
						}
						t6 := int32(uint32(v3<<3) / uint32(i32(7)))
						v3 = i32_shr_u(i32(-1), int32(bits.LeadingZeros32(uint32(t6+i32(-1))))) + i32(1)
						goto l4
					}
				l2:
					p7 := v3&i32(8) + i32(8)
					if uint32(v3) < uint32(i32(4)) {
						p7 = i32(4)
					}
					v3 = p7
				}
			l4:
				v9 = int64(uint32(v3)) * i64(296)
				if int32(int64(uint64(v9)>>32)) != 0 {
					goto l5
				}
				v8 = v3 + i32(8)
				t8 := v8
				v10 = int32(v9)
				v6 = t8 + v10
				if uint32(v6) < uint32(v8) {
					goto l5
				}
				if uint32(v6) > uint32(i32(0x7ffffff8)) {
					goto l5
				}
				t9 := m.fn5(v6)
				v5 = t9
				if v5 != 0 {
					v6 = v5 + v10
					if v8 == 0 {
						goto l7
					}
					memory_fill(m.memory, uint32(v6), i32(255), uint32(v8))
				l7:
					v5 = v3 + i32(-1)
					p10 := int32(uint32(v3)>>3) * i32(7)
					if uint32(v3) < uint32(i32(9)) {
						p10 = v5
					}
					v7 = p10
					t11 := int32(load32(m.memory[uint32(v0):]))
					v11 = t11
					{
						if v2 == 0 {
							goto l8
						}
						t12 := int64(load64(m.memory[uint32(v11):]))
						v9 = (t12 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
						t13 := int64(load64(m.memory[int64(uint32(v1))+8:]))
						v12 = t13
						t14 := int64(load64(m.memory[uint32(v1):]))
						v13 = t14
						v8 = v11
						v1 = v2
						v3 = i32(0)
					l14:
						{
							if v9 != i64(0) {
								goto l9
							}
						l10:
							{
								v3 = v3 + i32(8)
								v8 = v8 + i32(8)
								t15 := int64(load64(m.memory[uint32(v8):]))
								v9 = t15 & i64(-0x7f7f7f7f7f7f7f80)
								if v9 == i64(-0x7f7f7f7f7f7f7f80) {
									goto l10
								}
							}
							v9 = v9 ^ i64(-0x7f7f7f7f7f7f7f80)
						l9:
							{
								t16 := v6
								t17 := v5
								t18 := v13
								t19 := v12
								t20 := v11
								v14 = int32(uint32(int64(bits.TrailingZeros64(uint64(v9))))>>3) + v3
								t21 := int32(load32(m.memory[uint32(t20+(i32(0)-v14)*i32(296)+i32(-296)):]))
								t22 := m.fn94(t18, t19, t21)
								v15 = int32(t22)
								v10 = t17 & v15
								t23 := int64(load64(m.memory[uint32(t16+v10):]))
								v16 = t23 & i64(-0x7f7f7f7f7f7f7f80)
								if v16 != i64(0) {
									goto l11
								}
								v17 = i32(8)
							l12:
								{
									v10 = v10 + v17
									v17 = v17 + i32(8)
									t24 := v6
									v10 = v10 & v5
									t25 := int64(load64(m.memory[uint32(t24+v10):]))
									v16 = t25 & i64(-0x7f7f7f7f7f7f7f80)
									if v16 == 0 {
										goto l12
									}
								}
							}
						l11:
							v18 = v9 + i64(-1)
							{
								t26 := v6
								v10 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v16))))>>3) + v10) & v5
								t27 := int32(int8(m.memory[uint32(t26+v10)]))
								if t27 < i32(0) {
									goto l13
								}
								t28 := int64(load64(m.memory[uint32(v6):]))
								v10 = int32(uint32(int64(bits.TrailingZeros64(uint64(t28&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
							}
						l13:
							v9 = v18 & v9
							t29 := v6 + v10
							v15 = int32(uint32(v15) >> 25)
							m.memory[uint32(t29)] = byte(v15)
							m.memory[uint32(v6+(v10+i32(-8))&v5+i32(8))] = byte(v15)
							memory_copy(m.memory, uint32(v6+(v10^i32(-1))*i32(296)), uint32(v11+(v14^i32(-1))*i32(296)), uint32(i32(296)))
							v1 = v1 + i32(-1)
							if v1 != 0 {
								goto l14
							}
						}
					}
				l8:
					store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
					store32(m.memory[uint32(v0):], uint32(v6))
					store32(m.memory[int64(uint32(v0))+8:], uint32(v7-v2))
					if v4 == 0 {
						goto l15
					}
					t30 := v4
					v8 = (v4*i32(296) + i32(303)) & i32(-8)
					v3 = t30 + v8 + i32(9)
					if v3 == 0 {
						goto l15
					}
					v4 = v11 - v8
					t31 := int32(load32(m.memory[uint32(v4+i32(-4)):]))
					v8 = t31
					v6 = v8 & i32(-8)
					t32 := v6
					v8 = v8 & i32(3)
					p33 := i32(8)
					if v8 != 0 {
						p33 = i32(4)
					}
					if uint32(t32) < uint32(p33+v3) {
						m.fn3(i32(1273840), i32(46), i32(1273888))
						panic("unreachable")
					}
					if v8 == 0 {
						goto l17
					}
					if uint32(v6) > uint32(v3+i32(39)) {
						m.fn3(i32(1273904), i32(46), i32(1273952))
						panic("unreachable")
					}
				l17:
					m.fn1(v4)
					return i32(-1)
				}
				m.fn24(i32(8), v6)
				panic("unreachable")
			}
		}
	l1:
		{
			if v5 != 0 {
				goto l19
			}
			v3 = i32(0)
			goto l20
		l19:
			t34 := int32(load32(m.memory[uint32(v0):]))
			v8 = t34
			v3 = i32(0)
			{
				{
					t35 := v6
					var p36 int32
					if v5&i32(7) != i32(0) {
						p36 = 1
					}
					v6 = t35 + p36
					if v6 == i32(1) {
						goto l21
					}
					v11 = v6 & i32(1)
					v10 = v6 & i32(0x3ffffffe)
					v3 = i32(0)
				l22:
					{
						v6 = v8 + v3
						t37 := int64(load64(m.memory[uint32(v6):]))
						t38 := v6
						v9 = t37
						store64(m.memory[uint32(t38):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
						v6 = v6 + i32(8)
						t39 := int64(load64(m.memory[uint32(v6):]))
						t40 := v6
						v9 = t39
						store64(m.memory[uint32(t40):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
						v3 = v3 + i32(16)
						v10 = v10 + i32(-2)
						if v10 != 0 {
							goto l22
						}
					}
					if v11 == 0 {
						goto l23
					}
				}
			l21:
				v3 = v8 + v3
				t41 := int64(load64(m.memory[uint32(v3):]))
				t42 := v3
				v9 = t41
				store64(m.memory[uint32(t42):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
			}
		l23:
			{
				if uint32(v5) < uint32(i32(8)) {
					goto l24
				}
				t43 := int64(load64(m.memory[uint32(v8):]))
				store64(m.memory[uint32(v8+v5):], uint64(t43))
				goto l25
			}
		l24:
			if v5 == 0 {
				goto l25
			}
			memory_copy(m.memory, uint32(v8+i32(8)), uint32(v8), uint32(v5))
		l25:
			t44 := int64(load64(m.memory[int64(uint32(v1))+8:]))
			v16 = t44
			t45 := int64(load64(m.memory[uint32(v1):]))
			v18 = t45
			v6 = v8
			v10 = i32(0)
		l34:
			{
				t46 := v8
				v3 = v10
				v5 = t46 + v3
				t47 := int32(m.memory[uint32(v5)])
				if t47 != i32(128) {
					goto l26
				}
				v15 = v8 + (v3^i32(-1))*i32(296)
				v14 = v8 + (i32(0)-v3)*i32(296) + i32(-296)
				{
				l33:
					{
						t48 := int32(load32(m.memory[uint32(v14):]))
						t49 := m.fn94(v18, v16, t48)
						t50 := v4
						v1 = int32(t49)
						v10 = t50 & v1
						v11 = v10
						{
							t51 := int64(load64(m.memory[uint32(v8+v10):]))
							v9 = t51 & i64(-0x7f7f7f7f7f7f7f80)
							if v9 != i64(0) {
								goto l27
							}
							v17 = i32(8)
							v11 = v10
						l28:
							{
								v11 = v11 + v17
								v17 = v17 + i32(8)
								t52 := v8
								v11 = v11 & v4
								t53 := int64(load64(m.memory[uint32(t52+v11):]))
								v9 = t53 & i64(-0x7f7f7f7f7f7f7f80)
								if v9 == 0 {
									goto l28
								}
							}
						}
					l27:
						{
							t54 := v8
							v11 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v9))))>>3) + v11) & v4
							t55 := int32(int8(m.memory[uint32(t54+v11)]))
							if t55 < i32(0) {
								goto l29
							}
							t56 := int64(load64(m.memory[uint32(v8):]))
							v11 = int32(uint32(int64(bits.TrailingZeros64(uint64(t56&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
						}
					l29:
						{
							if uint32((v11-v10^(v3-v10))&v4) < uint32(i32(8)) {
								goto l30
							}
							v10 = v8 + v11
							t57 := int32(m.memory[uint32(v10)])
							v17 = t57
							t58 := v10
							v1 = int32(uint32(v1) >> 25)
							m.memory[uint32(t58)] = byte(v1)
							m.memory[uint32(v8+(v11+i32(-8))&v4+i32(8))] = byte(v1)
							if v17 == i32(255) {
								goto l31
							}
							v10 = i32(-296)
							v19 = v8 + v11*i32(-296)
						l32:
							{
								v11 = v19 + v10
								t59 := int32(load32(m.memory[uint32(v11):]))
								v17 = t59
								t60 := v11
								v1 = v6 + v10
								t61 := int32(load32(m.memory[uint32(v1):]))
								store32(m.memory[uint32(t60):], uint32(t61))
								store32(m.memory[uint32(v1):], uint32(v17))
								v1 = v1 + i32(4)
								t62 := int32(load32(m.memory[uint32(v1):]))
								v17 = t62
								t63 := v1
								v11 = v11 + i32(4)
								t64 := int32(load32(m.memory[uint32(v11):]))
								store32(m.memory[uint32(t63):], uint32(t64))
								store32(m.memory[uint32(v11):], uint32(v17))
								v10 = v10 + i32(8)
								if v10 != 0 {
									goto l32
								}
								goto l33
							}
						}
					l30:
					}
					t65 := v5
					v10 = int32(uint32(v1) >> 25)
					m.memory[uint32(t65)] = byte(v10)
					m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(v10)
					goto l26
				}
			l31:
				m.memory[uint32(v5)] = byte(i32(255))
				m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(i32(255))
				memory_copy(m.memory, uint32(v8+(v11^i32(-1))*i32(296)), uint32(v15), uint32(i32(296)))
			}
		l26:
			v10 = v3 + i32(1)
			v6 = v6 + i32(-296)
			if v3 != v4 {
				goto l34
			}
			p66 := v7
			if uint32(v4) < uint32(i32(8)) {
				p66 = v4
			}
			v3 = p66
		}
	l20:
		store32(m.memory[int64(uint32(v0))+8:], uint32(v3-v2))
	l15:
		return i32(-1)
	l5:
		m.fn28(i32(1271248), i32(57), i32(1271276))
		panic("unreachable")
	}
}
func (m *Module) fn100(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7, v8 int32
	var v9 int64
	var v10, v11, v12 int32
	var v13, v14 int64
	var v15 int32
	var v16 int64
	var v17 int32
	var v18 int64
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
		v2 = t0
		v3 = v2 + i32(1)
		if v3 == 0 {
			m.fn28(i32(1271248), i32(57), i32(1271276))
			panic("unreachable")
		}
		{
			t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t2 := v3
			v4 = t1
			t3 := v4
			v5 = v4 + i32(1)
			v6 = int32(uint32(v5) >> 3)
			v7 = v6 * i32(7)
			p4 := v7
			if uint32(v4) < uint32(i32(8)) {
				p4 = t3
			}
			v8 = p4
			if uint32(t2) <= uint32(int32(uint32(v8)>>1)) {
				goto l1
			}
			{
				{
					{
						v8 = v8 + i32(1)
						p5 := v3
						if uint32(v8) > uint32(v3) {
							p5 = v8
						}
						v3 = p5
						if uint32(v3) < uint32(i32(15)) {
							goto l2
						}
						if uint32(v3) > uint32(i32(0x1fffffff)) {
							m.fn28(i32(1271248), i32(57), i32(1271276))
							panic("unreachable")
						}
						t6 := int32(uint32(v3<<3) / uint32(i32(7)))
						v3 = i32_shr_u(i32(-1), int32(bits.LeadingZeros32(uint32(t6+i32(-1))))) + i32(1)
						goto l4
					}
				l2:
					p7 := v3&i32(8) + i32(8)
					if uint32(v3) < uint32(i32(4)) {
						p7 = i32(4)
					}
					v3 = p7
				}
			l4:
				v9 = int64(uint32(v3)) * i64(12)
				if int32(int64(uint64(v9)>>32)) != 0 {
					goto l5
				}
				v6 = int32(v9)
				if uint32(v6) > uint32(i32(-8)) {
					goto l5
				}
				v8 = v3 + i32(8)
				t8 := v8
				v10 = (v6 + i32(7)) & i32(-8)
				v6 = t8 + v10
				if uint32(v6) < uint32(v8) {
					goto l5
				}
				if uint32(v6) > uint32(i32(0x7ffffff8)) {
					goto l5
				}
				t9 := m.fn5(v6)
				v5 = t9
				if v5 != 0 {
					v6 = v5 + v10
					if v8 == 0 {
						goto l7
					}
					memory_fill(m.memory, uint32(v6), i32(255), uint32(v8))
				l7:
					v5 = v3 + i32(-1)
					p10 := int32(uint32(v3)>>3) * i32(7)
					if uint32(v3) < uint32(i32(9)) {
						p10 = v5
					}
					v11 = p10
					t11 := int32(load32(m.memory[uint32(v0):]))
					v12 = t11
					{
						if v2 == 0 {
							goto l8
						}
						t12 := int64(load64(m.memory[uint32(v12):]))
						v9 = (t12 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
						t13 := int64(load64(m.memory[int64(uint32(v1))+8:]))
						v13 = t13
						t14 := int64(load64(m.memory[uint32(v1):]))
						v14 = t14
						v8 = v12
						v1 = v2
						v3 = i32(0)
					l14:
						{
							if v9 != i64(0) {
								goto l9
							}
						l10:
							{
								v3 = v3 + i32(8)
								v8 = v8 + i32(8)
								t15 := int64(load64(m.memory[uint32(v8):]))
								v9 = t15 & i64(-0x7f7f7f7f7f7f7f80)
								if v9 == i64(-0x7f7f7f7f7f7f7f80) {
									goto l10
								}
							}
							v9 = v9 ^ i64(-0x7f7f7f7f7f7f7f80)
						l9:
							{
								t16 := v6
								t17 := v5
								t18 := v14
								t19 := v13
								t20 := v12
								v15 = int32(uint32(int64(bits.TrailingZeros64(uint64(v9))))>>3) + v3
								t21 := int32(load32(m.memory[uint32(t20+(i32(0)-v15)*i32(12)+i32(-12)):]))
								t22 := m.fn94(t18, t19, t21)
								v7 = int32(t22)
								v10 = t17 & v7
								t23 := int64(load64(m.memory[uint32(t16+v10):]))
								v16 = t23 & i64(-0x7f7f7f7f7f7f7f80)
								if v16 != i64(0) {
									goto l11
								}
								v17 = i32(8)
							l12:
								{
									v10 = v10 + v17
									v17 = v17 + i32(8)
									t24 := v6
									v10 = v10 & v5
									t25 := int64(load64(m.memory[uint32(t24+v10):]))
									v16 = t25 & i64(-0x7f7f7f7f7f7f7f80)
									if v16 == 0 {
										goto l12
									}
								}
							}
						l11:
							v18 = v9 + i64(-1)
							{
								t26 := v6
								v10 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v16))))>>3) + v10) & v5
								t27 := int32(int8(m.memory[uint32(t26+v10)]))
								if t27 < i32(0) {
									goto l13
								}
								t28 := int64(load64(m.memory[uint32(v6):]))
								v10 = int32(uint32(int64(bits.TrailingZeros64(uint64(t28&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
							}
						l13:
							v9 = v18 & v9
							t29 := v6 + v10
							v7 = int32(uint32(v7) >> 25)
							m.memory[uint32(t29)] = byte(v7)
							m.memory[uint32(v6+(v10+i32(-8))&v5+i32(8))] = byte(v7)
							v10 = v6 + (v10^i32(-1))*i32(12)
							t30 := v10
							v15 = v12 + (v15^i32(-1))*i32(12)
							t31 := int32(load32(m.memory[int64(uint32(v15))+8:]))
							store32(m.memory[int64(uint32(t30))+8:], uint32(t31))
							t32 := int64(load64(m.memory[uint32(v15):]))
							store64(m.memory[uint32(v10):], uint64(t32))
							v1 = v1 + i32(-1)
							if v1 != 0 {
								goto l14
							}
						}
					}
				l8:
					store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
					store32(m.memory[uint32(v0):], uint32(v6))
					store32(m.memory[int64(uint32(v0))+8:], uint32(v11-v2))
					if v4 == 0 {
						goto l15
					}
					t33 := v4
					v8 = (v4*i32(12) + i32(19)) & i32(-8)
					v3 = t33 + v8 + i32(9)
					if v3 == 0 {
						goto l15
					}
					v4 = v12 - v8
					t34 := int32(load32(m.memory[uint32(v4+i32(-4)):]))
					v8 = t34
					v6 = v8 & i32(-8)
					t35 := v6
					v8 = v8 & i32(3)
					p36 := i32(8)
					if v8 != 0 {
						p36 = i32(4)
					}
					if uint32(t35) < uint32(p36+v3) {
						m.fn3(i32(1273840), i32(46), i32(1273888))
						panic("unreachable")
					}
					if v8 == 0 {
						goto l17
					}
					if uint32(v6) > uint32(v3+i32(39)) {
						m.fn3(i32(1273904), i32(46), i32(1273952))
						panic("unreachable")
					}
				l17:
					m.fn1(v4)
					return i32(-1)
				}
				m.fn24(i32(8), v6)
				panic("unreachable")
			}
		}
	l1:
		{
			if v5 != 0 {
				goto l19
			}
			v3 = i32(0)
			goto l20
		l19:
			t37 := int32(load32(m.memory[uint32(v0):]))
			v8 = t37
			v3 = i32(0)
			{
				{
					t38 := v6
					var p39 int32
					if v5&i32(7) != i32(0) {
						p39 = 1
					}
					v6 = t38 + p39
					if v6 == i32(1) {
						goto l21
					}
					v12 = v6 & i32(1)
					v10 = v6 & i32(0x3ffffffe)
					v3 = i32(0)
				l22:
					{
						v6 = v8 + v3
						t40 := int64(load64(m.memory[uint32(v6):]))
						t41 := v6
						v9 = t40
						store64(m.memory[uint32(t41):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
						v6 = v6 + i32(8)
						t42 := int64(load64(m.memory[uint32(v6):]))
						t43 := v6
						v9 = t42
						store64(m.memory[uint32(t43):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
						v3 = v3 + i32(16)
						v10 = v10 + i32(-2)
						if v10 != 0 {
							goto l22
						}
					}
					if v12 == 0 {
						goto l23
					}
				}
			l21:
				v3 = v8 + v3
				t44 := int64(load64(m.memory[uint32(v3):]))
				t45 := v3
				v9 = t44
				store64(m.memory[uint32(t45):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
			}
		l23:
			{
				if uint32(v5) < uint32(i32(8)) {
					goto l24
				}
				t46 := int64(load64(m.memory[uint32(v8):]))
				store64(m.memory[uint32(v8+v5):], uint64(t46))
				goto l25
			}
		l24:
			if v5 == 0 {
				goto l25
			}
			memory_copy(m.memory, uint32(v8+i32(8)), uint32(v8), uint32(v5))
		l25:
			t47 := int64(load64(m.memory[int64(uint32(v1))+8:]))
			v16 = t47
			t48 := int64(load64(m.memory[uint32(v1):]))
			v18 = t48
			v6 = i32(0)
		l33:
			{
				t49 := v8
				v3 = v6
				v10 = t49 + v3
				t50 := int32(m.memory[uint32(v10)])
				if t50 != i32(128) {
					goto l26
				}
				v15 = v8 + (v3^i32(-1))*i32(12)
				v1 = v8 + (i32(0)-v3)*i32(12) + i32(-12)
			l32:
				{
					t51 := int32(load32(m.memory[uint32(v1):]))
					t52 := m.fn94(v18, v16, t51)
					t53 := v4
					v12 = int32(t52)
					v6 = t53 & v12
					v5 = v6
					{
						t54 := int64(load64(m.memory[uint32(v8+v6):]))
						v9 = t54 & i64(-0x7f7f7f7f7f7f7f80)
						if v9 != i64(0) {
							goto l27
						}
						v11 = i32(8)
						v5 = v6
					l28:
						{
							v5 = v5 + v11
							v11 = v11 + i32(8)
							t55 := v8
							v5 = v5 & v4
							t56 := int64(load64(m.memory[uint32(t55+v5):]))
							v9 = t56 & i64(-0x7f7f7f7f7f7f7f80)
							if v9 == 0 {
								goto l28
							}
						}
					}
				l27:
					{
						t57 := v8
						v5 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v9))))>>3) + v5) & v4
						t58 := int32(int8(m.memory[uint32(t57+v5)]))
						if t58 < i32(0) {
							goto l29
						}
						t59 := int64(load64(m.memory[uint32(v8):]))
						v5 = int32(uint32(int64(bits.TrailingZeros64(uint64(t59&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
					}
				l29:
					{
						if uint32((v5-v6^(v3-v6))&v4) < uint32(i32(8)) {
							goto l30
						}
						v6 = v8 + v5
						t60 := int32(m.memory[uint32(v6)])
						v11 = t60
						t61 := v6
						v12 = int32(uint32(v12) >> 25)
						m.memory[uint32(t61)] = byte(v12)
						m.memory[uint32(v8+(v5+i32(-8))&v4+i32(8))] = byte(v12)
						v6 = v8 + (v5^i32(-1))*i32(12)
						{
							if v11 != i32(255) {
								t64 := int32(load32(m.memory[uint32(v15):]))
								v5 = t64
								t65 := int32(load32(m.memory[uint32(v6):]))
								store32(m.memory[uint32(v15):], uint32(t65))
								store32(m.memory[uint32(v6):], uint32(v5))
								t66 := int32(load32(m.memory[int64(uint32(v6))+4:]))
								v5 = t66
								t67 := int32(load32(m.memory[int64(uint32(v15))+4:]))
								store32(m.memory[int64(uint32(v6))+4:], uint32(t67))
								store32(m.memory[int64(uint32(v15))+4:], uint32(v5))
								t68 := int32(load32(m.memory[int64(uint32(v15))+8:]))
								v5 = t68
								t69 := int32(load32(m.memory[int64(uint32(v6))+8:]))
								store32(m.memory[int64(uint32(v15))+8:], uint32(t69))
								store32(m.memory[int64(uint32(v6))+8:], uint32(v5))
								goto l32
							}
							m.memory[uint32(v10)] = byte(i32(255))
							m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(i32(255))
							t62 := int32(load32(m.memory[int64(uint32(v15))+8:]))
							store32(m.memory[int64(uint32(v6))+8:], uint32(t62))
							t63 := int64(load64(m.memory[uint32(v15):]))
							store64(m.memory[uint32(v6):], uint64(t63))
							goto l26
						}
					}
				l30:
				}
				t70 := v10
				v6 = int32(uint32(v12) >> 25)
				m.memory[uint32(t70)] = byte(v6)
				m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(v6)
			}
		l26:
			v6 = v3 + i32(1)
			if v3 != v4 {
				goto l33
			}
			p71 := v7
			if uint32(v4) < uint32(i32(8)) {
				p71 = v4
			}
			v3 = p71
		}
	l20:
		store32(m.memory[int64(uint32(v0))+8:], uint32(v3-v2))
	l15:
		return i32(-1)
	}
l5:
	m.fn28(i32(1271248), i32(57), i32(1271276))
	panic("unreachable")
}
func (m *Module) fn101(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7, v8 int32
	var v9 int64
	var v10, v11, v12 int32
	var v13, v14 int64
	var v15 int32
	var v16 int64
	var v17 int32
	var v18 int64
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
		v2 = t0
		v3 = v2 + i32(1)
		if v3 == 0 {
			m.fn28(i32(1271248), i32(57), i32(1271276))
			panic("unreachable")
		}
		{
			t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t2 := v3
			v4 = t1
			t3 := v4
			v5 = v4 + i32(1)
			v6 = int32(uint32(v5) >> 3)
			v7 = v6 * i32(7)
			p4 := v7
			if uint32(v4) < uint32(i32(8)) {
				p4 = t3
			}
			v8 = p4
			if uint32(t2) <= uint32(int32(uint32(v8)>>1)) {
				goto l1
			}
			{
				{
					{
						v8 = v8 + i32(1)
						p5 := v3
						if uint32(v8) > uint32(v3) {
							p5 = v8
						}
						v3 = p5
						if uint32(v3) < uint32(i32(15)) {
							goto l2
						}
						if uint32(v3) > uint32(i32(0x1fffffff)) {
							m.fn28(i32(1271248), i32(57), i32(1271276))
							panic("unreachable")
						}
						t6 := int32(uint32(v3<<3) / uint32(i32(7)))
						v3 = i32_shr_u(i32(-1), int32(bits.LeadingZeros32(uint32(t6+i32(-1))))) + i32(1)
						goto l4
					}
				l2:
					p7 := v3&i32(8) + i32(8)
					if uint32(v3) < uint32(i32(4)) {
						p7 = i32(4)
					}
					v3 = p7
				}
			l4:
				v9 = int64(uint32(v3)) * i64(96)
				if int32(int64(uint64(v9)>>32)) != 0 {
					goto l5
				}
				v8 = v3 + i32(8)
				t8 := v8
				v10 = int32(v9)
				v6 = t8 + v10
				if uint32(v6) < uint32(v8) {
					goto l5
				}
				if uint32(v6) > uint32(i32(0x7ffffff8)) {
					goto l5
				}
				t9 := m.fn5(v6)
				v5 = t9
				if v5 != 0 {
					v6 = v5 + v10
					if v8 == 0 {
						goto l7
					}
					memory_fill(m.memory, uint32(v6), i32(255), uint32(v8))
				l7:
					v5 = v3 + i32(-1)
					p10 := int32(uint32(v3)>>3) * i32(7)
					if uint32(v3) < uint32(i32(9)) {
						p10 = v5
					}
					v11 = p10
					t11 := int32(load32(m.memory[uint32(v0):]))
					v12 = t11
					{
						if v2 == 0 {
							goto l8
						}
						t12 := int64(load64(m.memory[uint32(v12):]))
						v9 = (t12 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
						t13 := int64(load64(m.memory[int64(uint32(v1))+8:]))
						v13 = t13
						t14 := int64(load64(m.memory[uint32(v1):]))
						v14 = t14
						v8 = v12
						v1 = v2
						v3 = i32(0)
					l14:
						{
							if v9 != i64(0) {
								goto l9
							}
						l10:
							{
								v3 = v3 + i32(8)
								v8 = v8 + i32(8)
								t15 := int64(load64(m.memory[uint32(v8):]))
								v9 = t15 & i64(-0x7f7f7f7f7f7f7f80)
								if v9 == i64(-0x7f7f7f7f7f7f7f80) {
									goto l10
								}
							}
							v9 = v9 ^ i64(-0x7f7f7f7f7f7f7f80)
						l9:
							{
								t16 := v6
								t17 := v5
								t18 := v14
								t19 := v13
								t20 := v12
								v15 = int32(uint32(int64(bits.TrailingZeros64(uint64(v9))))>>3) + v3
								t21 := int32(load32(m.memory[uint32(t20+(i32(0)-v15)*i32(96)+i32(-96)):]))
								t22 := m.fn94(t18, t19, t21)
								v7 = int32(t22)
								v10 = t17 & v7
								t23 := int64(load64(m.memory[uint32(t16+v10):]))
								v16 = t23 & i64(-0x7f7f7f7f7f7f7f80)
								if v16 != i64(0) {
									goto l11
								}
								v17 = i32(8)
							l12:
								{
									v10 = v10 + v17
									v17 = v17 + i32(8)
									t24 := v6
									v10 = v10 & v5
									t25 := int64(load64(m.memory[uint32(t24+v10):]))
									v16 = t25 & i64(-0x7f7f7f7f7f7f7f80)
									if v16 == 0 {
										goto l12
									}
								}
							}
						l11:
							v18 = v9 + i64(-1)
							{
								t26 := v6
								v10 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v16))))>>3) + v10) & v5
								t27 := int32(int8(m.memory[uint32(t26+v10)]))
								if t27 < i32(0) {
									goto l13
								}
								t28 := int64(load64(m.memory[uint32(v6):]))
								v10 = int32(uint32(int64(bits.TrailingZeros64(uint64(t28&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
							}
						l13:
							v9 = v18 & v9
							t29 := v6 + v10
							v7 = int32(uint32(v7) >> 25)
							m.memory[uint32(t29)] = byte(v7)
							m.memory[uint32(v6+(v10+i32(-8))&v5+i32(8))] = byte(v7)
							memory_copy(m.memory, uint32(v6+(v10^i32(-1))*i32(96)), uint32(v12+(v15^i32(-1))*i32(96)), uint32(i32(96)))
							v1 = v1 + i32(-1)
							if v1 != 0 {
								goto l14
							}
						}
					}
				l8:
					store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
					store32(m.memory[uint32(v0):], uint32(v6))
					store32(m.memory[int64(uint32(v0))+8:], uint32(v11-v2))
					if v4 == 0 {
						goto l15
					}
					t30 := v4
					v8 = (v4*i32(96) + i32(103)) & i32(-32)
					v3 = t30 + v8 + i32(9)
					if v3 == 0 {
						goto l15
					}
					v4 = v12 - v8
					t31 := int32(load32(m.memory[uint32(v4+i32(-4)):]))
					v8 = t31
					v6 = v8 & i32(-8)
					t32 := v6
					v8 = v8 & i32(3)
					p33 := i32(8)
					if v8 != 0 {
						p33 = i32(4)
					}
					if uint32(t32) < uint32(p33+v3) {
						m.fn3(i32(1273840), i32(46), i32(1273888))
						panic("unreachable")
					}
					if v8 == 0 {
						goto l17
					}
					if uint32(v6) > uint32(v3+i32(39)) {
						m.fn3(i32(1273904), i32(46), i32(1273952))
						panic("unreachable")
					}
				l17:
					m.fn1(v4)
					return i32(-1)
				}
				m.fn24(i32(8), v6)
				panic("unreachable")
			}
		}
	l1:
		{
			if v5 != 0 {
				goto l19
			}
			v3 = i32(0)
			goto l20
		l19:
			t34 := int32(load32(m.memory[uint32(v0):]))
			v8 = t34
			v3 = i32(0)
			{
				{
					t35 := v6
					var p36 int32
					if v5&i32(7) != i32(0) {
						p36 = 1
					}
					v6 = t35 + p36
					if v6 == i32(1) {
						goto l21
					}
					v12 = v6 & i32(1)
					v10 = v6 & i32(0x3ffffffe)
					v3 = i32(0)
				l22:
					{
						v6 = v8 + v3
						t37 := int64(load64(m.memory[uint32(v6):]))
						t38 := v6
						v9 = t37
						store64(m.memory[uint32(t38):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
						v6 = v6 + i32(8)
						t39 := int64(load64(m.memory[uint32(v6):]))
						t40 := v6
						v9 = t39
						store64(m.memory[uint32(t40):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
						v3 = v3 + i32(16)
						v10 = v10 + i32(-2)
						if v10 != 0 {
							goto l22
						}
					}
					if v12 == 0 {
						goto l23
					}
				}
			l21:
				v3 = v8 + v3
				t41 := int64(load64(m.memory[uint32(v3):]))
				t42 := v3
				v9 = t41
				store64(m.memory[uint32(t42):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
			}
		l23:
			{
				if uint32(v5) < uint32(i32(8)) {
					goto l24
				}
				t43 := int64(load64(m.memory[uint32(v8):]))
				store64(m.memory[uint32(v8+v5):], uint64(t43))
				goto l25
			}
		l24:
			if v5 == 0 {
				goto l25
			}
			memory_copy(m.memory, uint32(v8+i32(8)), uint32(v8), uint32(v5))
		l25:
			t44 := int64(load64(m.memory[int64(uint32(v1))+8:]))
			v16 = t44
			t45 := int64(load64(m.memory[uint32(v1):]))
			v18 = t45
			v6 = i32(0)
		l33:
			{
				t46 := v8
				v3 = v6
				v10 = t46 + v3
				t47 := int32(m.memory[uint32(v10)])
				if t47 != i32(128) {
					goto l26
				}
				v15 = v8 + (v3^i32(-1))*i32(96)
				v1 = v8 + (i32(0)-v3)*i32(96) + i32(-96)
			l32:
				{
					t48 := int32(load32(m.memory[uint32(v1):]))
					t49 := m.fn94(v18, v16, t48)
					t50 := v4
					v12 = int32(t49)
					v6 = t50 & v12
					v5 = v6
					{
						t51 := int64(load64(m.memory[uint32(v8+v6):]))
						v9 = t51 & i64(-0x7f7f7f7f7f7f7f80)
						if v9 != i64(0) {
							goto l27
						}
						v11 = i32(8)
						v5 = v6
					l28:
						{
							v5 = v5 + v11
							v11 = v11 + i32(8)
							t52 := v8
							v5 = v5 & v4
							t53 := int64(load64(m.memory[uint32(t52+v5):]))
							v9 = t53 & i64(-0x7f7f7f7f7f7f7f80)
							if v9 == 0 {
								goto l28
							}
						}
					}
				l27:
					{
						t54 := v8
						v5 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v9))))>>3) + v5) & v4
						t55 := int32(int8(m.memory[uint32(t54+v5)]))
						if t55 < i32(0) {
							goto l29
						}
						t56 := int64(load64(m.memory[uint32(v8):]))
						v5 = int32(uint32(int64(bits.TrailingZeros64(uint64(t56&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
					}
				l29:
					{
						if uint32((v5-v6^(v3-v6))&v4) < uint32(i32(8)) {
							goto l30
						}
						v6 = v8 + v5
						t57 := int32(m.memory[uint32(v6)])
						v11 = t57
						t58 := v6
						v12 = int32(uint32(v12) >> 25)
						m.memory[uint32(t58)] = byte(v12)
						m.memory[uint32(v8+(v5+i32(-8))&v4+i32(8))] = byte(v12)
						v6 = v8 + (v5^i32(-1))*i32(96)
						if v11 != i32(255) {
							t59 := int32(load32(m.memory[int64(uint32(v15))+20:]))
							v5 = t59
							t60 := int32(load32(m.memory[int64(uint32(v15))+16:]))
							v12 = t60
							t61 := int64(load64(m.memory[int64(uint32(v6))+16:]))
							store64(m.memory[int64(uint32(v15))+16:], uint64(t61))
							t62 := int64(load64(m.memory[uint32(v15):]))
							v9 = t62
							t63 := int64(load64(m.memory[uint32(v6):]))
							store64(m.memory[uint32(v15):], uint64(t63))
							store64(m.memory[uint32(v6):], uint64(v9))
							t64 := int64(load64(m.memory[int64(uint32(v15))+8:]))
							v9 = t64
							t65 := int64(load64(m.memory[int64(uint32(v6))+8:]))
							store64(m.memory[int64(uint32(v15))+8:], uint64(t65))
							store64(m.memory[int64(uint32(v6))+8:], uint64(v9))
							store32(m.memory[int64(uint32(v6))+16:], uint32(v12))
							t66 := int32(load32(m.memory[int64(uint32(v15))+24:]))
							v12 = t66
							t67 := int32(load32(m.memory[int64(uint32(v6))+24:]))
							store32(m.memory[int64(uint32(v15))+24:], uint32(t67))
							store32(m.memory[int64(uint32(v6))+20:], uint32(v5))
							store32(m.memory[int64(uint32(v6))+24:], uint32(v12))
							t68 := int32(load32(m.memory[int64(uint32(v15))+28:]))
							v5 = t68
							t69 := int32(load32(m.memory[int64(uint32(v6))+28:]))
							store32(m.memory[int64(uint32(v15))+28:], uint32(t69))
							store32(m.memory[int64(uint32(v6))+28:], uint32(v5))
							t70 := int32(load32(m.memory[int64(uint32(v15))+32:]))
							v5 = t70
							t71 := int32(load32(m.memory[int64(uint32(v6))+32:]))
							store32(m.memory[int64(uint32(v15))+32:], uint32(t71))
							store32(m.memory[int64(uint32(v6))+32:], uint32(v5))
							t72 := int32(load32(m.memory[int64(uint32(v15))+36:]))
							v5 = t72
							t73 := int32(load32(m.memory[int64(uint32(v6))+36:]))
							store32(m.memory[int64(uint32(v15))+36:], uint32(t73))
							store32(m.memory[int64(uint32(v6))+36:], uint32(v5))
							t74 := int32(load32(m.memory[int64(uint32(v15))+40:]))
							v5 = t74
							t75 := int32(load32(m.memory[int64(uint32(v6))+40:]))
							store32(m.memory[int64(uint32(v15))+40:], uint32(t75))
							store32(m.memory[int64(uint32(v6))+40:], uint32(v5))
							t76 := int32(load32(m.memory[int64(uint32(v15))+44:]))
							v5 = t76
							t77 := int32(load32(m.memory[int64(uint32(v6))+44:]))
							store32(m.memory[int64(uint32(v15))+44:], uint32(t77))
							store32(m.memory[int64(uint32(v6))+44:], uint32(v5))
							t78 := int32(load32(m.memory[int64(uint32(v15))+48:]))
							v5 = t78
							t79 := int32(load32(m.memory[int64(uint32(v6))+48:]))
							store32(m.memory[int64(uint32(v15))+48:], uint32(t79))
							store32(m.memory[int64(uint32(v6))+48:], uint32(v5))
							t80 := int32(load32(m.memory[int64(uint32(v15))+52:]))
							v5 = t80
							t81 := int32(load32(m.memory[int64(uint32(v6))+52:]))
							store32(m.memory[int64(uint32(v15))+52:], uint32(t81))
							store32(m.memory[int64(uint32(v6))+52:], uint32(v5))
							t82 := int32(load32(m.memory[int64(uint32(v15))+56:]))
							v5 = t82
							t83 := int32(load32(m.memory[int64(uint32(v6))+56:]))
							store32(m.memory[int64(uint32(v15))+56:], uint32(t83))
							store32(m.memory[int64(uint32(v6))+56:], uint32(v5))
							t84 := int32(load32(m.memory[int64(uint32(v15))+60:]))
							v5 = t84
							t85 := int32(load32(m.memory[int64(uint32(v6))+60:]))
							store32(m.memory[int64(uint32(v15))+60:], uint32(t85))
							store32(m.memory[int64(uint32(v6))+60:], uint32(v5))
							t86 := int32(load32(m.memory[int64(uint32(v15))+64:]))
							v5 = t86
							t87 := int32(load32(m.memory[int64(uint32(v6))+64:]))
							store32(m.memory[int64(uint32(v15))+64:], uint32(t87))
							store32(m.memory[int64(uint32(v6))+64:], uint32(v5))
							t88 := int32(load32(m.memory[int64(uint32(v15))+68:]))
							v5 = t88
							t89 := int32(load32(m.memory[int64(uint32(v6))+68:]))
							store32(m.memory[int64(uint32(v15))+68:], uint32(t89))
							store32(m.memory[int64(uint32(v6))+68:], uint32(v5))
							t90 := int32(load32(m.memory[int64(uint32(v15))+72:]))
							v5 = t90
							t91 := int32(load32(m.memory[int64(uint32(v6))+72:]))
							store32(m.memory[int64(uint32(v15))+72:], uint32(t91))
							store32(m.memory[int64(uint32(v6))+72:], uint32(v5))
							t92 := int32(load32(m.memory[int64(uint32(v15))+76:]))
							v5 = t92
							t93 := int32(load32(m.memory[int64(uint32(v6))+76:]))
							store32(m.memory[int64(uint32(v15))+76:], uint32(t93))
							store32(m.memory[int64(uint32(v6))+76:], uint32(v5))
							t94 := int32(load32(m.memory[int64(uint32(v15))+80:]))
							v5 = t94
							t95 := int32(load32(m.memory[int64(uint32(v6))+80:]))
							store32(m.memory[int64(uint32(v15))+80:], uint32(t95))
							store32(m.memory[int64(uint32(v6))+80:], uint32(v5))
							t96 := int32(load32(m.memory[int64(uint32(v15))+84:]))
							v5 = t96
							t97 := int32(load32(m.memory[int64(uint32(v6))+84:]))
							store32(m.memory[int64(uint32(v15))+84:], uint32(t97))
							store32(m.memory[int64(uint32(v6))+84:], uint32(v5))
							t98 := int32(load32(m.memory[int64(uint32(v15))+88:]))
							v5 = t98
							t99 := int32(load32(m.memory[int64(uint32(v6))+88:]))
							store32(m.memory[int64(uint32(v15))+88:], uint32(t99))
							store32(m.memory[int64(uint32(v6))+88:], uint32(v5))
							t100 := int32(load32(m.memory[int64(uint32(v15))+92:]))
							v5 = t100
							t101 := int32(load32(m.memory[int64(uint32(v6))+92:]))
							store32(m.memory[int64(uint32(v15))+92:], uint32(t101))
							store32(m.memory[int64(uint32(v6))+92:], uint32(v5))
							goto l32
						}
						m.memory[uint32(v10)] = byte(i32(255))
						m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(i32(255))
						memory_copy(m.memory, uint32(v6), uint32(v15), uint32(i32(96)))
						goto l26
					}
				l30:
				}
				t102 := v10
				v6 = int32(uint32(v12) >> 25)
				m.memory[uint32(t102)] = byte(v6)
				m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(v6)
			}
		l26:
			v6 = v3 + i32(1)
			if v3 != v4 {
				goto l33
			}
			p103 := v7
			if uint32(v4) < uint32(i32(8)) {
				p103 = v4
			}
			v3 = p103
		}
	l20:
		store32(m.memory[int64(uint32(v0))+8:], uint32(v3-v2))
	l15:
		return i32(-1)
	}
l5:
	m.fn28(i32(1271248), i32(57), i32(1271276))
	panic("unreachable")
}
func (m *Module) fn102(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7, v8 int32
	var v9 int64
	var v10, v11, v12 int32
	var v13, v14 int64
	var v15 int32
	var v16 int64
	var v17 int32
	var v18 int64
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
		v2 = t0
		v3 = v2 + i32(1)
		if v3 == 0 {
			m.fn28(i32(1271248), i32(57), i32(1271276))
			panic("unreachable")
		}
		{
			t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t2 := v3
			v4 = t1
			t3 := v4
			v5 = v4 + i32(1)
			v6 = int32(uint32(v5) >> 3)
			v7 = v6 * i32(7)
			p4 := v7
			if uint32(v4) < uint32(i32(8)) {
				p4 = t3
			}
			v8 = p4
			if uint32(t2) <= uint32(int32(uint32(v8)>>1)) {
				goto l1
			}
			{
				{
					{
						v8 = v8 + i32(1)
						p5 := v3
						if uint32(v8) > uint32(v3) {
							p5 = v8
						}
						v3 = p5
						if uint32(v3) < uint32(i32(15)) {
							goto l2
						}
						if uint32(v3) > uint32(i32(0x1fffffff)) {
							m.fn28(i32(1271248), i32(57), i32(1271276))
							panic("unreachable")
						}
						t6 := int32(uint32(v3<<3) / uint32(i32(7)))
						v3 = i32_shr_u(i32(-1), int32(bits.LeadingZeros32(uint32(t6+i32(-1))))) + i32(1)
						goto l4
					}
				l2:
					p7 := v3&i32(8) + i32(8)
					if uint32(v3) < uint32(i32(4)) {
						p7 = i32(4)
					}
					v3 = p7
				}
			l4:
				v9 = int64(uint32(v3)) * i64(20)
				if int32(int64(uint64(v9)>>32)) != 0 {
					goto l5
				}
				v6 = int32(v9)
				if uint32(v6) > uint32(i32(-8)) {
					goto l5
				}
				v8 = v3 + i32(8)
				t8 := v8
				v10 = (v6 + i32(7)) & i32(-8)
				v6 = t8 + v10
				if uint32(v6) < uint32(v8) {
					goto l5
				}
				if uint32(v6) > uint32(i32(0x7ffffff8)) {
					goto l5
				}
				t9 := m.fn5(v6)
				v5 = t9
				if v5 != 0 {
					v6 = v5 + v10
					if v8 == 0 {
						goto l7
					}
					memory_fill(m.memory, uint32(v6), i32(255), uint32(v8))
				l7:
					v5 = v3 + i32(-1)
					p10 := int32(uint32(v3)>>3) * i32(7)
					if uint32(v3) < uint32(i32(9)) {
						p10 = v5
					}
					v11 = p10
					t11 := int32(load32(m.memory[uint32(v0):]))
					v12 = t11
					{
						if v2 == 0 {
							goto l8
						}
						t12 := int64(load64(m.memory[uint32(v12):]))
						v9 = (t12 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
						t13 := int64(load64(m.memory[int64(uint32(v1))+8:]))
						v13 = t13
						t14 := int64(load64(m.memory[uint32(v1):]))
						v14 = t14
						v8 = v12
						v15 = v2
						v3 = i32(0)
					l14:
						{
							if v9 != i64(0) {
								goto l9
							}
						l10:
							{
								v3 = v3 + i32(8)
								v8 = v8 + i32(8)
								t15 := int64(load64(m.memory[uint32(v8):]))
								v9 = t15 & i64(-0x7f7f7f7f7f7f7f80)
								if v9 == i64(-0x7f7f7f7f7f7f7f80) {
									goto l10
								}
							}
							v9 = v9 ^ i64(-0x7f7f7f7f7f7f7f80)
						l9:
							{
								t16 := v6
								t17 := v5
								t18 := v14
								t19 := v13
								t20 := v12
								v1 = int32(uint32(int64(bits.TrailingZeros64(uint64(v9))))>>3) + v3
								t21 := int32(load32(m.memory[uint32(t20+(i32(0)-v1)*i32(20)+i32(-20)):]))
								t22 := m.fn94(t18, t19, t21)
								v7 = int32(t22)
								v10 = t17 & v7
								t23 := int64(load64(m.memory[uint32(t16+v10):]))
								v16 = t23 & i64(-0x7f7f7f7f7f7f7f80)
								if v16 != i64(0) {
									goto l11
								}
								v17 = i32(8)
							l12:
								{
									v10 = v10 + v17
									v17 = v17 + i32(8)
									t24 := v6
									v10 = v10 & v5
									t25 := int64(load64(m.memory[uint32(t24+v10):]))
									v16 = t25 & i64(-0x7f7f7f7f7f7f7f80)
									if v16 == 0 {
										goto l12
									}
								}
							}
						l11:
							v18 = v9 + i64(-1)
							{
								t26 := v6
								v10 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v16))))>>3) + v10) & v5
								t27 := int32(int8(m.memory[uint32(t26+v10)]))
								if t27 < i32(0) {
									goto l13
								}
								t28 := int64(load64(m.memory[uint32(v6):]))
								v10 = int32(uint32(int64(bits.TrailingZeros64(uint64(t28&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
							}
						l13:
							v9 = v18 & v9
							t29 := v6 + v10
							v7 = int32(uint32(v7) >> 25)
							m.memory[uint32(t29)] = byte(v7)
							m.memory[uint32(v6+(v10+i32(-8))&v5+i32(8))] = byte(v7)
							v10 = v6 + (v10^i32(-1))*i32(20)
							t30 := v10
							v1 = v12 + (v1^i32(-1))*i32(20)
							t31 := int32(load32(m.memory[int64(uint32(v1))+16:]))
							store32(m.memory[int64(uint32(t30))+16:], uint32(t31))
							t32 := int64(load64(m.memory[int64(uint32(v1))+8:]))
							store64(m.memory[int64(uint32(v10))+8:], uint64(t32))
							t33 := int64(load64(m.memory[uint32(v1):]))
							store64(m.memory[uint32(v10):], uint64(t33))
							v15 = v15 + i32(-1)
							if v15 != 0 {
								goto l14
							}
						}
					}
				l8:
					store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
					store32(m.memory[uint32(v0):], uint32(v6))
					store32(m.memory[int64(uint32(v0))+8:], uint32(v11-v2))
					if v4 == 0 {
						goto l15
					}
					t34 := v4
					v8 = (v4*i32(20) + i32(27)) & i32(-8)
					v3 = t34 + v8 + i32(9)
					if v3 == 0 {
						goto l15
					}
					v4 = v12 - v8
					t35 := int32(load32(m.memory[uint32(v4+i32(-4)):]))
					v8 = t35
					v6 = v8 & i32(-8)
					t36 := v6
					v8 = v8 & i32(3)
					p37 := i32(8)
					if v8 != 0 {
						p37 = i32(4)
					}
					if uint32(t36) < uint32(p37+v3) {
						m.fn3(i32(1273840), i32(46), i32(1273888))
						panic("unreachable")
					}
					if v8 == 0 {
						goto l17
					}
					if uint32(v6) > uint32(v3+i32(39)) {
						m.fn3(i32(1273904), i32(46), i32(1273952))
						panic("unreachable")
					}
				l17:
					m.fn1(v4)
					return i32(-1)
				}
				m.fn24(i32(8), v6)
				panic("unreachable")
			}
		}
	l1:
		{
			if v5 != 0 {
				goto l19
			}
			v3 = i32(0)
			goto l20
		l19:
			t38 := int32(load32(m.memory[uint32(v0):]))
			v8 = t38
			v3 = i32(0)
			{
				{
					t39 := v6
					var p40 int32
					if v5&i32(7) != i32(0) {
						p40 = 1
					}
					v6 = t39 + p40
					if v6 == i32(1) {
						goto l21
					}
					v12 = v6 & i32(1)
					v10 = v6 & i32(0x3ffffffe)
					v3 = i32(0)
				l22:
					{
						v6 = v8 + v3
						t41 := int64(load64(m.memory[uint32(v6):]))
						t42 := v6
						v9 = t41
						store64(m.memory[uint32(t42):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
						v6 = v6 + i32(8)
						t43 := int64(load64(m.memory[uint32(v6):]))
						t44 := v6
						v9 = t43
						store64(m.memory[uint32(t44):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
						v3 = v3 + i32(16)
						v10 = v10 + i32(-2)
						if v10 != 0 {
							goto l22
						}
					}
					if v12 == 0 {
						goto l23
					}
				}
			l21:
				v3 = v8 + v3
				t45 := int64(load64(m.memory[uint32(v3):]))
				t46 := v3
				v9 = t45
				store64(m.memory[uint32(t46):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
			}
		l23:
			{
				if uint32(v5) < uint32(i32(8)) {
					goto l24
				}
				t47 := int64(load64(m.memory[uint32(v8):]))
				store64(m.memory[uint32(v8+v5):], uint64(t47))
				goto l25
			}
		l24:
			if v5 == 0 {
				goto l25
			}
			memory_copy(m.memory, uint32(v8+i32(8)), uint32(v8), uint32(v5))
		l25:
			t48 := int64(load64(m.memory[int64(uint32(v1))+8:]))
			v16 = t48
			t49 := int64(load64(m.memory[uint32(v1):]))
			v18 = t49
			v6 = i32(0)
		l33:
			{
				t50 := v8
				v3 = v6
				v10 = t50 + v3
				t51 := int32(m.memory[uint32(v10)])
				if t51 != i32(128) {
					goto l26
				}
				v15 = v8 + (v3^i32(-1))*i32(20)
				v12 = v8 + (i32(0)-v3)*i32(20) + i32(-20)
			l32:
				{
					t52 := int32(load32(m.memory[uint32(v12):]))
					t53 := m.fn94(v18, v16, t52)
					t54 := v4
					v1 = int32(t53)
					v6 = t54 & v1
					v5 = v6
					{
						t55 := int64(load64(m.memory[uint32(v8+v6):]))
						v9 = t55 & i64(-0x7f7f7f7f7f7f7f80)
						if v9 != i64(0) {
							goto l27
						}
						v11 = i32(8)
						v5 = v6
					l28:
						{
							v5 = v5 + v11
							v11 = v11 + i32(8)
							t56 := v8
							v5 = v5 & v4
							t57 := int64(load64(m.memory[uint32(t56+v5):]))
							v9 = t57 & i64(-0x7f7f7f7f7f7f7f80)
							if v9 == 0 {
								goto l28
							}
						}
					}
				l27:
					{
						t58 := v8
						v5 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v9))))>>3) + v5) & v4
						t59 := int32(int8(m.memory[uint32(t58+v5)]))
						if t59 < i32(0) {
							goto l29
						}
						t60 := int64(load64(m.memory[uint32(v8):]))
						v5 = int32(uint32(int64(bits.TrailingZeros64(uint64(t60&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
					}
				l29:
					{
						if uint32((v5-v6^(v3-v6))&v4) < uint32(i32(8)) {
							goto l30
						}
						v6 = v8 + v5
						t61 := int32(m.memory[uint32(v6)])
						v11 = t61
						t62 := v6
						v1 = int32(uint32(v1) >> 25)
						m.memory[uint32(t62)] = byte(v1)
						m.memory[uint32(v8+(v5+i32(-8))&v4+i32(8))] = byte(v1)
						v6 = v8 + (v5^i32(-1))*i32(20)
						{
							if v11 != i32(255) {
								t66 := int32(load32(m.memory[uint32(v15):]))
								v5 = t66
								t67 := int32(load32(m.memory[uint32(v6):]))
								store32(m.memory[uint32(v15):], uint32(t67))
								store32(m.memory[uint32(v6):], uint32(v5))
								t68 := int32(load32(m.memory[int64(uint32(v6))+4:]))
								v5 = t68
								t69 := int32(load32(m.memory[int64(uint32(v15))+4:]))
								store32(m.memory[int64(uint32(v6))+4:], uint32(t69))
								store32(m.memory[int64(uint32(v15))+4:], uint32(v5))
								t70 := int32(load32(m.memory[int64(uint32(v15))+8:]))
								v5 = t70
								t71 := int32(load32(m.memory[int64(uint32(v6))+8:]))
								store32(m.memory[int64(uint32(v15))+8:], uint32(t71))
								store32(m.memory[int64(uint32(v6))+8:], uint32(v5))
								t72 := int32(load32(m.memory[int64(uint32(v6))+12:]))
								v5 = t72
								t73 := int32(load32(m.memory[int64(uint32(v15))+12:]))
								store32(m.memory[int64(uint32(v6))+12:], uint32(t73))
								store32(m.memory[int64(uint32(v15))+12:], uint32(v5))
								t74 := int32(load32(m.memory[int64(uint32(v15))+16:]))
								v5 = t74
								t75 := int32(load32(m.memory[int64(uint32(v6))+16:]))
								store32(m.memory[int64(uint32(v15))+16:], uint32(t75))
								store32(m.memory[int64(uint32(v6))+16:], uint32(v5))
								goto l32
							}
							m.memory[uint32(v10)] = byte(i32(255))
							m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(i32(255))
							t63 := int32(load32(m.memory[int64(uint32(v15))+16:]))
							store32(m.memory[int64(uint32(v6))+16:], uint32(t63))
							t64 := int64(load64(m.memory[int64(uint32(v15))+8:]))
							store64(m.memory[int64(uint32(v6))+8:], uint64(t64))
							t65 := int64(load64(m.memory[uint32(v15):]))
							store64(m.memory[uint32(v6):], uint64(t65))
							goto l26
						}
					}
				l30:
				}
				t76 := v10
				v6 = int32(uint32(v1) >> 25)
				m.memory[uint32(t76)] = byte(v6)
				m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(v6)
			}
		l26:
			v6 = v3 + i32(1)
			if v3 != v4 {
				goto l33
			}
			p77 := v7
			if uint32(v4) < uint32(i32(8)) {
				p77 = v4
			}
			v3 = p77
		}
	l20:
		store32(m.memory[int64(uint32(v0))+8:], uint32(v3-v2))
	l15:
		return i32(-1)
	}
l5:
	m.fn28(i32(1271248), i32(57), i32(1271276))
	panic("unreachable")
}
func (m *Module) fn103(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7, v8 int32
	var v9 int64
	var v10, v11 int32
	var v12, v13 int64
	var v14, v15 int32
	var v16 int64
	var v17 int32
	var v18 int64
	var v19, v20, v21 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
		v2 = t0
		v3 = v2 + i32(1)
		if v3 == 0 {
			m.fn28(i32(1271248), i32(57), i32(1271276))
			panic("unreachable")
		}
		{
			t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t2 := v3
			v4 = t1
			t3 := v4
			v5 = v4 + i32(1)
			v6 = int32(uint32(v5) >> 3)
			v7 = v6 * i32(7)
			p4 := v7
			if uint32(v4) < uint32(i32(8)) {
				p4 = t3
			}
			v8 = p4
			if uint32(t2) <= uint32(int32(uint32(v8)>>1)) {
				goto l1
			}
			{
				{
					{
						v8 = v8 + i32(1)
						p5 := v3
						if uint32(v8) > uint32(v3) {
							p5 = v8
						}
						v3 = p5
						if uint32(v3) < uint32(i32(15)) {
							goto l2
						}
						if uint32(v3) > uint32(i32(0x1fffffff)) {
							m.fn28(i32(1271248), i32(57), i32(1271276))
							panic("unreachable")
						}
						t6 := int32(uint32(v3<<3) / uint32(i32(7)))
						v3 = i32_shr_u(i32(-1), int32(bits.LeadingZeros32(uint32(t6+i32(-1))))) + i32(1)
						goto l4
					}
				l2:
					p7 := v3&i32(8) + i32(8)
					if uint32(v3) < uint32(i32(4)) {
						p7 = i32(4)
					}
					v3 = p7
				}
			l4:
				v9 = int64(uint32(v3)) * i64(368)
				if int32(int64(uint64(v9)>>32)) != 0 {
					goto l5
				}
				v8 = v3 + i32(8)
				t8 := v8
				v10 = int32(v9)
				v6 = t8 + v10
				if uint32(v6) < uint32(v8) {
					goto l5
				}
				if uint32(v6) > uint32(i32(0x7ffffff8)) {
					goto l5
				}
				t9 := m.fn5(v6)
				v5 = t9
				if v5 != 0 {
					v6 = v5 + v10
					if v8 == 0 {
						goto l7
					}
					memory_fill(m.memory, uint32(v6), i32(255), uint32(v8))
				l7:
					v5 = v3 + i32(-1)
					p10 := int32(uint32(v3)>>3) * i32(7)
					if uint32(v3) < uint32(i32(9)) {
						p10 = v5
					}
					v7 = p10
					t11 := int32(load32(m.memory[uint32(v0):]))
					v11 = t11
					{
						if v2 == 0 {
							goto l8
						}
						t12 := int64(load64(m.memory[uint32(v11):]))
						v9 = (t12 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
						t13 := int64(load64(m.memory[int64(uint32(v1))+8:]))
						v12 = t13
						t14 := int64(load64(m.memory[uint32(v1):]))
						v13 = t14
						v8 = v11
						v1 = v2
						v3 = i32(0)
					l14:
						{
							if v9 != i64(0) {
								goto l9
							}
						l10:
							{
								v3 = v3 + i32(8)
								v8 = v8 + i32(8)
								t15 := int64(load64(m.memory[uint32(v8):]))
								v9 = t15 & i64(-0x7f7f7f7f7f7f7f80)
								if v9 == i64(-0x7f7f7f7f7f7f7f80) {
									goto l10
								}
							}
							v9 = v9 ^ i64(-0x7f7f7f7f7f7f7f80)
						l9:
							{
								t16 := v6
								t17 := v5
								t18 := v13
								t19 := v12
								t20 := v11
								v14 = int32(uint32(int64(bits.TrailingZeros64(uint64(v9))))>>3) + v3
								t21 := int32(load32(m.memory[uint32(t20+(i32(0)-v14)*i32(368)+i32(-368)):]))
								t22 := m.fn94(t18, t19, t21)
								v15 = int32(t22)
								v10 = t17 & v15
								t23 := int64(load64(m.memory[uint32(t16+v10):]))
								v16 = t23 & i64(-0x7f7f7f7f7f7f7f80)
								if v16 != i64(0) {
									goto l11
								}
								v17 = i32(8)
							l12:
								{
									v10 = v10 + v17
									v17 = v17 + i32(8)
									t24 := v6
									v10 = v10 & v5
									t25 := int64(load64(m.memory[uint32(t24+v10):]))
									v16 = t25 & i64(-0x7f7f7f7f7f7f7f80)
									if v16 == 0 {
										goto l12
									}
								}
							}
						l11:
							v18 = v9 + i64(-1)
							{
								t26 := v6
								v10 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v16))))>>3) + v10) & v5
								t27 := int32(int8(m.memory[uint32(t26+v10)]))
								if t27 < i32(0) {
									goto l13
								}
								t28 := int64(load64(m.memory[uint32(v6):]))
								v10 = int32(uint32(int64(bits.TrailingZeros64(uint64(t28&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
							}
						l13:
							v9 = v18 & v9
							t29 := v6 + v10
							v15 = int32(uint32(v15) >> 25)
							m.memory[uint32(t29)] = byte(v15)
							m.memory[uint32(v6+(v10+i32(-8))&v5+i32(8))] = byte(v15)
							memory_copy(m.memory, uint32(v6+(v10^i32(-1))*i32(368)), uint32(v11+(v14^i32(-1))*i32(368)), uint32(i32(368)))
							v1 = v1 + i32(-1)
							if v1 != 0 {
								goto l14
							}
						}
					}
				l8:
					store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
					store32(m.memory[uint32(v0):], uint32(v6))
					store32(m.memory[int64(uint32(v0))+8:], uint32(v7-v2))
					if v4 == 0 {
						goto l15
					}
					t30 := v4
					v8 = (v4*i32(368) + i32(375)) & i32(-16)
					v3 = t30 + v8 + i32(9)
					if v3 == 0 {
						goto l15
					}
					v4 = v11 - v8
					t31 := int32(load32(m.memory[uint32(v4+i32(-4)):]))
					v8 = t31
					v6 = v8 & i32(-8)
					t32 := v6
					v8 = v8 & i32(3)
					p33 := i32(8)
					if v8 != 0 {
						p33 = i32(4)
					}
					if uint32(t32) < uint32(p33+v3) {
						m.fn3(i32(1273840), i32(46), i32(1273888))
						panic("unreachable")
					}
					if v8 == 0 {
						goto l17
					}
					if uint32(v6) > uint32(v3+i32(39)) {
						m.fn3(i32(1273904), i32(46), i32(1273952))
						panic("unreachable")
					}
				l17:
					m.fn1(v4)
					return i32(-1)
				}
				m.fn24(i32(8), v6)
				panic("unreachable")
			}
		}
	l1:
		{
			if v5 != 0 {
				goto l19
			}
			v3 = i32(0)
			goto l20
		l19:
			t34 := int32(load32(m.memory[uint32(v0):]))
			v8 = t34
			v3 = i32(0)
			{
				{
					t35 := v6
					var p36 int32
					if v5&i32(7) != i32(0) {
						p36 = 1
					}
					v6 = t35 + p36
					if v6 == i32(1) {
						goto l21
					}
					v11 = v6 & i32(1)
					v10 = v6 & i32(0x3ffffffe)
					v3 = i32(0)
				l22:
					{
						v6 = v8 + v3
						t37 := int64(load64(m.memory[uint32(v6):]))
						t38 := v6
						v9 = t37
						store64(m.memory[uint32(t38):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
						v6 = v6 + i32(8)
						t39 := int64(load64(m.memory[uint32(v6):]))
						t40 := v6
						v9 = t39
						store64(m.memory[uint32(t40):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
						v3 = v3 + i32(16)
						v10 = v10 + i32(-2)
						if v10 != 0 {
							goto l22
						}
					}
					if v11 == 0 {
						goto l23
					}
				}
			l21:
				v3 = v8 + v3
				t41 := int64(load64(m.memory[uint32(v3):]))
				t42 := v3
				v9 = t41
				store64(m.memory[uint32(t42):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
			}
		l23:
			{
				if uint32(v5) < uint32(i32(8)) {
					goto l24
				}
				t43 := int64(load64(m.memory[uint32(v8):]))
				store64(m.memory[uint32(v8+v5):], uint64(t43))
				goto l25
			}
		l24:
			if v5 == 0 {
				goto l25
			}
			memory_copy(m.memory, uint32(v8+i32(8)), uint32(v8), uint32(v5))
		l25:
			t44 := int64(load64(m.memory[int64(uint32(v1))+8:]))
			v16 = t44
			t45 := int64(load64(m.memory[uint32(v1):]))
			v18 = t45
			v6 = v8
			v10 = i32(0)
		l34:
			{
				t46 := v8
				v3 = v10
				v5 = t46 + v3
				t47 := int32(m.memory[uint32(v5)])
				if t47 != i32(128) {
					goto l26
				}
				v15 = v8 + (v3^i32(-1))*i32(368)
				v14 = v8 + (i32(0)-v3)*i32(368) + i32(-368)
				{
				l33:
					{
						t48 := int32(load32(m.memory[uint32(v14):]))
						t49 := m.fn94(v18, v16, t48)
						t50 := v4
						v1 = int32(t49)
						v10 = t50 & v1
						v11 = v10
						{
							t51 := int64(load64(m.memory[uint32(v8+v10):]))
							v9 = t51 & i64(-0x7f7f7f7f7f7f7f80)
							if v9 != i64(0) {
								goto l27
							}
							v17 = i32(8)
							v11 = v10
						l28:
							{
								v11 = v11 + v17
								v17 = v17 + i32(8)
								t52 := v8
								v11 = v11 & v4
								t53 := int64(load64(m.memory[uint32(t52+v11):]))
								v9 = t53 & i64(-0x7f7f7f7f7f7f7f80)
								if v9 == 0 {
									goto l28
								}
							}
						}
					l27:
						{
							t54 := v8
							v11 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v9))))>>3) + v11) & v4
							t55 := int32(int8(m.memory[uint32(t54+v11)]))
							if t55 < i32(0) {
								goto l29
							}
							t56 := int64(load64(m.memory[uint32(v8):]))
							v11 = int32(uint32(int64(bits.TrailingZeros64(uint64(t56&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
						}
					l29:
						{
							if uint32((v11-v10^(v3-v10))&v4) < uint32(i32(8)) {
								goto l30
							}
							v10 = v8 + v11
							t57 := int32(m.memory[uint32(v10)])
							v17 = t57
							t58 := v10
							v1 = int32(uint32(v1) >> 25)
							m.memory[uint32(t58)] = byte(v1)
							m.memory[uint32(v8+(v11+i32(-8))&v4+i32(8))] = byte(v1)
							if v17 == i32(255) {
								goto l31
							}
							v1 = i32(-368)
							v19 = v8 + v11*i32(-368)
						l32:
							{
								v10 = v6 + v1
								t59 := int32(load32(m.memory[uint32(v10):]))
								v17 = t59
								t60 := v10
								v11 = v19 + v1
								t61 := int32(load32(m.memory[uint32(v11):]))
								store32(m.memory[uint32(t60):], uint32(t61))
								store32(m.memory[uint32(v11):], uint32(v17))
								v17 = v11 + i32(4)
								t62 := int32(load32(m.memory[uint32(v17):]))
								v20 = t62
								t63 := v17
								v21 = v10 + i32(4)
								t64 := int32(load32(m.memory[uint32(v21):]))
								store32(m.memory[uint32(t63):], uint32(t64))
								store32(m.memory[uint32(v21):], uint32(v20))
								v17 = v10 + i32(8)
								t65 := int32(load32(m.memory[uint32(v17):]))
								v20 = t65
								t66 := v17
								v21 = v11 + i32(8)
								t67 := int32(load32(m.memory[uint32(v21):]))
								store32(m.memory[uint32(t66):], uint32(t67))
								store32(m.memory[uint32(v21):], uint32(v20))
								v11 = v11 + i32(12)
								t68 := int32(load32(m.memory[uint32(v11):]))
								v17 = t68
								t69 := v11
								v10 = v10 + i32(12)
								t70 := int32(load32(m.memory[uint32(v10):]))
								store32(m.memory[uint32(t69):], uint32(t70))
								store32(m.memory[uint32(v10):], uint32(v17))
								v1 = v1 + i32(16)
								if v1 != 0 {
									goto l32
								}
								goto l33
							}
						}
					l30:
					}
					t71 := v5
					v10 = int32(uint32(v1) >> 25)
					m.memory[uint32(t71)] = byte(v10)
					m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(v10)
					goto l26
				}
			l31:
				m.memory[uint32(v5)] = byte(i32(255))
				m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(i32(255))
				memory_copy(m.memory, uint32(v8+(v11^i32(-1))*i32(368)), uint32(v15), uint32(i32(368)))
			}
		l26:
			v10 = v3 + i32(1)
			v6 = v6 + i32(-368)
			if v3 != v4 {
				goto l34
			}
			p72 := v7
			if uint32(v4) < uint32(i32(8)) {
				p72 = v4
			}
			v3 = p72
		}
	l20:
		store32(m.memory[int64(uint32(v0))+8:], uint32(v3-v2))
	l15:
		return i32(-1)
	l5:
		m.fn28(i32(1271248), i32(57), i32(1271276))
		panic("unreachable")
	}
}
func (m *Module) fn104(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7, v8, v9, v10 int32
	var v11 int64
	var v12, v13, v14 int32
	var v15 int64
	var v16 int32
	var v17 int64
	{
		{
			{
				t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
				v2 = t0
				v3 = v2 + i32(1)
				if v3 == 0 {
					m.fn28(i32(1271248), i32(57), i32(1271276))
					panic("unreachable")
				}
				{
					t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
					t2 := v3
					v4 = t1
					t3 := v4
					v5 = v4 + i32(1)
					v6 = int32(uint32(v5) >> 3)
					v7 = v6 * i32(7)
					p4 := v7
					if uint32(v4) < uint32(i32(8)) {
						p4 = t3
					}
					v8 = p4
					if uint32(t2) <= uint32(int32(uint32(v8)>>1)) {
						{
							if v5 != 0 {
								goto l6
							}
							v3 = i32(0)
							goto l7
						l6:
							t7 := int32(load32(m.memory[uint32(v0):]))
							v8 = t7
							v3 = i32(0)
							{
								{
									t8 := v6
									var p9 int32
									if v5&i32(7) != i32(0) {
										p9 = 1
									}
									v6 = t8 + p9
									if v6 == i32(1) {
										goto l8
									}
									v9 = v6 & i32(1)
									v10 = v6 & i32(0x3ffffffe)
									v3 = i32(0)
								l9:
									{
										v6 = v8 + v3
										t10 := int64(load64(m.memory[uint32(v6):]))
										t11 := v6
										v11 = t10
										store64(m.memory[uint32(t11):], uint64(int64(uint64(v11^i64(-1))>>7)&i64(72340172838076673)+(v11|i64(0x7f7f7f7f7f7f7f7f))))
										v6 = v6 + i32(8)
										t12 := int64(load64(m.memory[uint32(v6):]))
										t13 := v6
										v11 = t12
										store64(m.memory[uint32(t13):], uint64(int64(uint64(v11^i64(-1))>>7)&i64(72340172838076673)+(v11|i64(0x7f7f7f7f7f7f7f7f))))
										v3 = v3 + i32(16)
										v10 = v10 + i32(-2)
										if v10 != 0 {
											goto l9
										}
									}
									if v9 == 0 {
										goto l10
									}
								}
							l8:
								v3 = v8 + v3
								t14 := int64(load64(m.memory[uint32(v3):]))
								t15 := v3
								v11 = t14
								store64(m.memory[uint32(t15):], uint64(int64(uint64(v11^i64(-1))>>7)&i64(72340172838076673)+(v11|i64(0x7f7f7f7f7f7f7f7f))))
							}
						l10:
							{
								if uint32(v5) < uint32(i32(8)) {
									goto l11
								}
								t16 := int64(load64(m.memory[uint32(v8):]))
								store64(m.memory[uint32(v8+v5):], uint64(t16))
								goto l12
							}
						l11:
							if v5 == 0 {
								goto l12
							}
							memory_copy(m.memory, uint32(v8+i32(8)), uint32(v8), uint32(v5))
						l12:
							v6 = i32(0)
						l20:
							{
								t17 := v8
								v3 = v6
								v10 = t17 + v3
								t18 := int32(m.memory[uint32(v10)])
								if t18 != i32(128) {
									goto l13
								}
								v12 = v8 - v3<<2 + i32(-4)
								v13 = v8 + (v3^i32(-1))<<2
							l19:
								{
									t19 := int32(load32(m.memory[uint32(v12):]))
									t20 := v4
									v6 = t19
									v9 = int32(((((int64(uint32(v6&i32(255)))^i64(-0x340d631b7bdddcdb))*i64(0x100000001b3)^int64(uint32(int32(uint32(v6)>>8)&i32(255))))*i64(0x100000001b3)^int64(uint32(int32(uint32(v6)>>16)&i32(255))))*i64(0x100000001b3) ^ int64(uint32(int32(uint32(v6)>>24)))) * i64(0x100000001b3))
									v6 = t20 & v9
									v5 = v6
									{
										t21 := int64(load64(m.memory[uint32(v8+v6):]))
										v11 = t21 & i64(-0x7f7f7f7f7f7f7f80)
										if v11 != i64(0) {
											goto l14
										}
										v14 = i32(8)
										v5 = v6
									l15:
										{
											v5 = v5 + v14
											v14 = v14 + i32(8)
											t22 := v8
											v5 = v5 & v4
											t23 := int64(load64(m.memory[uint32(t22+v5):]))
											v11 = t23 & i64(-0x7f7f7f7f7f7f7f80)
											if v11 == 0 {
												goto l15
											}
										}
									}
								l14:
									{
										t24 := v8
										v5 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v11))))>>3) + v5) & v4
										t25 := int32(int8(m.memory[uint32(t24+v5)]))
										if t25 < i32(0) {
											goto l16
										}
										t26 := int64(load64(m.memory[uint32(v8):]))
										v5 = int32(uint32(int64(bits.TrailingZeros64(uint64(t26&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
									}
								l16:
									{
										if uint32((v5-v6^(v3-v6))&v4) < uint32(i32(8)) {
											goto l17
										}
										v6 = v8 + v5
										t27 := int32(m.memory[uint32(v6)])
										v14 = t27
										t28 := v6
										v9 = int32(uint32(v9) >> 25)
										m.memory[uint32(t28)] = byte(v9)
										m.memory[uint32(v8+(v5+i32(-8))&v4+i32(8))] = byte(v9)
										v6 = v8 + (v5^i32(-1))<<2
										{
											if v14 != i32(255) {
												t30 := int32(load32(m.memory[uint32(v13):]))
												v5 = t30
												t31 := int32(load32(m.memory[uint32(v6):]))
												store32(m.memory[uint32(v13):], uint32(t31))
												store32(m.memory[uint32(v6):], uint32(v5))
												goto l19
											}
											m.memory[uint32(v10)] = byte(i32(255))
											m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(i32(255))
											t29 := int32(load32(m.memory[uint32(v13):]))
											store32(m.memory[uint32(v6):], uint32(t29))
											goto l13
										}
									}
								l17:
								}
								t32 := v10
								v6 = int32(uint32(v9) >> 25)
								m.memory[uint32(t32)] = byte(v6)
								m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(v6)
							}
						l13:
							v6 = v3 + i32(1)
							if v3 != v4 {
								goto l20
							}
							p33 := v7
							if uint32(v4) < uint32(i32(8)) {
								p33 = v4
							}
							v3 = p33
						}
					l7:
						store32(m.memory[int64(uint32(v0))+8:], uint32(v3-v2))
						goto l21
					}
					v8 = v8 + i32(1)
					p5 := v3
					if uint32(v8) > uint32(v3) {
						p5 = v8
					}
					v3 = p5
					if uint32(v3) < uint32(i32(15)) {
						goto l2
					}
					{
						if uint32(v3) > uint32(i32(0x1fffffff)) {
							m.fn28(i32(1271248), i32(57), i32(1271276))
							panic("unreachable")
						}
						t6 := int32(uint32(v3<<3) / uint32(i32(7)))
						v3 = i32_shr_u(i32(-1), int32(bits.LeadingZeros32(uint32(t6+i32(-1)))))
						if uint32(v3) > uint32(i32(0x3ffffffe)) {
							goto l4
						}
						v3 = v3 + i32(1)
						goto l5
					}
				}
			}
		l2:
			p34 := v3&i32(8) + i32(8)
			if uint32(v3) < uint32(i32(4)) {
				p34 = i32(4)
			}
			v3 = p34
		}
	l5:
		v8 = v3 + i32(8)
		t35 := v8
		v10 = (v3<<2 + i32(7)) & i32(-8)
		v6 = t35 + v10
		if uint32(v6) < uint32(v8) {
			goto l4
		}
		if uint32(v6) > uint32(i32(0x7ffffff8)) {
			goto l4
		}
		{
			t36 := m.fn5(v6)
			v5 = t36
			if v5 != 0 {
				v6 = v5 + v10
				if v8 == 0 {
					goto l23
				}
				memory_fill(m.memory, uint32(v6), i32(255), uint32(v8))
			l23:
				v5 = v3 + i32(-1)
				p37 := int32(uint32(v3)>>3) * i32(7)
				if uint32(v3) < uint32(i32(9)) {
					p37 = v5
				}
				v14 = p37
				t38 := int32(load32(m.memory[uint32(v0):]))
				v9 = t38
				{
					if v2 == 0 {
						goto l24
					}
					t39 := int64(load64(m.memory[uint32(v9):]))
					v11 = (t39 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
					v8 = v9
					v12 = v2
					v3 = i32(0)
				l30:
					{
						if v11 != i64(0) {
							goto l25
						}
					l26:
						{
							v3 = v3 + i32(8)
							v8 = v8 + i32(8)
							t40 := int64(load64(m.memory[uint32(v8):]))
							v11 = t40 & i64(-0x7f7f7f7f7f7f7f80)
							if v11 == i64(-0x7f7f7f7f7f7f7f80) {
								goto l26
							}
						}
						v11 = v11 ^ i64(-0x7f7f7f7f7f7f7f80)
					l25:
						{
							t41 := v6
							t42 := v5
							t43 := v9
							v13 = int32(uint32(int64(bits.TrailingZeros64(uint64(v11))))>>3) + v3
							t44 := int32(load32(m.memory[uint32(t43-v13<<2+i32(-4)):]))
							v10 = t44
							v7 = int32(((((int64(uint32(v10&i32(255)))^i64(-0x340d631b7bdddcdb))*i64(0x100000001b3)^int64(uint32(int32(uint32(v10)>>8)&i32(255))))*i64(0x100000001b3)^int64(uint32(int32(uint32(v10)>>16)&i32(255))))*i64(0x100000001b3) ^ int64(uint32(int32(uint32(v10)>>24)))) * i64(0x100000001b3))
							v10 = t42 & v7
							t45 := int64(load64(m.memory[uint32(t41+v10):]))
							v15 = t45 & i64(-0x7f7f7f7f7f7f7f80)
							if v15 != i64(0) {
								goto l27
							}
							v16 = i32(8)
						l28:
							{
								v10 = v10 + v16
								v16 = v16 + i32(8)
								t46 := v6
								v10 = v10 & v5
								t47 := int64(load64(m.memory[uint32(t46+v10):]))
								v15 = t47 & i64(-0x7f7f7f7f7f7f7f80)
								if v15 == 0 {
									goto l28
								}
							}
						}
					l27:
						v17 = v11 + i64(-1)
						{
							t48 := v6
							v10 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v15))))>>3) + v10) & v5
							t49 := int32(int8(m.memory[uint32(t48+v10)]))
							if t49 < i32(0) {
								goto l29
							}
							t50 := int64(load64(m.memory[uint32(v6):]))
							v10 = int32(uint32(int64(bits.TrailingZeros64(uint64(t50&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
						}
					l29:
						v11 = v17 & v11
						t51 := v6 + v10
						v7 = int32(uint32(v7) >> 25)
						m.memory[uint32(t51)] = byte(v7)
						m.memory[uint32(v6+(v10+i32(-8))&v5+i32(8))] = byte(v7)
						t52 := int32(load32(m.memory[uint32(v9+(v13^i32(-1))<<2):]))
						store32(m.memory[uint32(v6+(v10^i32(-1))<<2):], uint32(t52))
						v12 = v12 + i32(-1)
						if v12 != 0 {
							goto l30
						}
					}
				}
			l24:
				store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
				store32(m.memory[uint32(v0):], uint32(v6))
				store32(m.memory[int64(uint32(v0))+8:], uint32(v14-v2))
				if v4 == 0 {
					goto l21
				}
				t53 := v4
				v8 = (v4<<2 + i32(11)) & i32(-8)
				v3 = t53 + v8 + i32(9)
				if v3 == 0 {
					goto l21
				}
				v4 = v9 - v8
				t54 := int32(load32(m.memory[uint32(v4+i32(-4)):]))
				v8 = t54
				v6 = v8 & i32(-8)
				t55 := v6
				v8 = v8 & i32(3)
				p56 := i32(8)
				if v8 != 0 {
					p56 = i32(4)
				}
				if uint32(t55) < uint32(p56+v3) {
					m.fn3(i32(1273840), i32(46), i32(1273888))
					panic("unreachable")
				}
				if v8 == 0 {
					goto l32
				}
				if uint32(v6) > uint32(v3+i32(39)) {
					m.fn3(i32(1273904), i32(46), i32(1273952))
					panic("unreachable")
				}
			l32:
				m.fn1(v4)
				return i32(-1)
			}
			m.fn24(i32(8), v6)
			panic("unreachable")
		}
	}
l21:
	return i32(-1)
l4:
	m.fn28(i32(1271248), i32(57), i32(1271276))
	panic("unreachable")
}
func (m *Module) fn105(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7, v8, v9, v10 int32
	var v11, v12, v13 int64
	var v14, v15 int32
	var v16, v17 int64
	var v18 int32
	{
		{
			{
				t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
				v2 = t0
				v3 = v2 + i32(1)
				if v3 == 0 {
					m.fn28(i32(1271248), i32(57), i32(1271276))
					panic("unreachable")
				}
				{
					t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
					t2 := v3
					v4 = t1
					t3 := v4
					v5 = v4 + i32(1)
					v6 = int32(uint32(v5) >> 3)
					v7 = v6 * i32(7)
					p4 := v7
					if uint32(v4) < uint32(i32(8)) {
						p4 = t3
					}
					v8 = p4
					if uint32(t2) <= uint32(int32(uint32(v8)>>1)) {
						{
							if v5 != 0 {
								goto l6
							}
							v3 = i32(0)
							goto l7
						l6:
							t7 := int32(load32(m.memory[uint32(v0):]))
							v8 = t7
							v3 = i32(0)
							{
								{
									t8 := v6
									var p9 int32
									if v5&i32(7) != i32(0) {
										p9 = 1
									}
									v6 = t8 + p9
									if v6 == i32(1) {
										goto l8
									}
									v9 = v6 & i32(1)
									v10 = v6 & i32(0x3ffffffe)
									v3 = i32(0)
								l9:
									{
										v6 = v8 + v3
										t10 := int64(load64(m.memory[uint32(v6):]))
										t11 := v6
										v11 = t10
										store64(m.memory[uint32(t11):], uint64(int64(uint64(v11^i64(-1))>>7)&i64(72340172838076673)+(v11|i64(0x7f7f7f7f7f7f7f7f))))
										v6 = v6 + i32(8)
										t12 := int64(load64(m.memory[uint32(v6):]))
										t13 := v6
										v11 = t12
										store64(m.memory[uint32(t13):], uint64(int64(uint64(v11^i64(-1))>>7)&i64(72340172838076673)+(v11|i64(0x7f7f7f7f7f7f7f7f))))
										v3 = v3 + i32(16)
										v10 = v10 + i32(-2)
										if v10 != 0 {
											goto l9
										}
									}
									if v9 == 0 {
										goto l10
									}
								}
							l8:
								v3 = v8 + v3
								t14 := int64(load64(m.memory[uint32(v3):]))
								t15 := v3
								v11 = t14
								store64(m.memory[uint32(t15):], uint64(int64(uint64(v11^i64(-1))>>7)&i64(72340172838076673)+(v11|i64(0x7f7f7f7f7f7f7f7f))))
							}
						l10:
							{
								if uint32(v5) < uint32(i32(8)) {
									goto l11
								}
								t16 := int64(load64(m.memory[uint32(v8):]))
								store64(m.memory[uint32(v8+v5):], uint64(t16))
								goto l12
							}
						l11:
							if v5 == 0 {
								goto l12
							}
							memory_copy(m.memory, uint32(v8+i32(8)), uint32(v8), uint32(v5))
						l12:
							t17 := int64(load64(m.memory[int64(uint32(v1))+8:]))
							v12 = t17
							t18 := int64(load64(m.memory[uint32(v1):]))
							v13 = t18
							v6 = i32(0)
						l20:
							{
								t19 := v8
								v3 = v6
								v10 = t19 + v3
								t20 := int32(m.memory[uint32(v10)])
								if t20 != i32(128) {
									goto l13
								}
								v1 = v8 - v3<<4 + i32(-16)
								v14 = v8 + (v3^i32(-1))<<4
							l19:
								{
									t21 := int32(load16(m.memory[uint32(v1):]))
									t22 := m.fn106(v13, v12, t21)
									t23 := v4
									v9 = int32(t22)
									v6 = t23 & v9
									v5 = v6
									{
										t24 := int64(load64(m.memory[uint32(v8+v6):]))
										v11 = t24 & i64(-0x7f7f7f7f7f7f7f80)
										if v11 != i64(0) {
											goto l14
										}
										v15 = i32(8)
										v5 = v6
									l15:
										{
											v5 = v5 + v15
											v15 = v15 + i32(8)
											t25 := v8
											v5 = v5 & v4
											t26 := int64(load64(m.memory[uint32(t25+v5):]))
											v11 = t26 & i64(-0x7f7f7f7f7f7f7f80)
											if v11 == 0 {
												goto l15
											}
										}
									}
								l14:
									{
										t27 := v8
										v5 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v11))))>>3) + v5) & v4
										t28 := int32(int8(m.memory[uint32(t27+v5)]))
										if t28 < i32(0) {
											goto l16
										}
										t29 := int64(load64(m.memory[uint32(v8):]))
										v5 = int32(uint32(int64(bits.TrailingZeros64(uint64(t29&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
									}
								l16:
									{
										if uint32((v5-v6^(v3-v6))&v4) < uint32(i32(8)) {
											goto l17
										}
										v6 = v8 + v5
										t30 := int32(m.memory[uint32(v6)])
										v15 = t30
										t31 := v6
										v9 = int32(uint32(v9) >> 25)
										m.memory[uint32(t31)] = byte(v9)
										m.memory[uint32(v8+(v5+i32(-8))&v4+i32(8))] = byte(v9)
										v6 = v8 + (v5^i32(-1))<<4
										{
											if v15 != i32(255) {
												t34 := int64(load64(m.memory[uint32(v14):]))
												v11 = t34
												t35 := int64(load64(m.memory[uint32(v6):]))
												store64(m.memory[uint32(v14):], uint64(t35))
												store64(m.memory[uint32(v6):], uint64(v11))
												t36 := int64(load64(m.memory[int64(uint32(v14))+8:]))
												v11 = t36
												t37 := int64(load64(m.memory[int64(uint32(v6))+8:]))
												store64(m.memory[int64(uint32(v14))+8:], uint64(t37))
												store64(m.memory[int64(uint32(v6))+8:], uint64(v11))
												goto l19
											}
											m.memory[uint32(v10)] = byte(i32(255))
											m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(i32(255))
											t32 := int64(load64(m.memory[int64(uint32(v14))+8:]))
											store64(m.memory[int64(uint32(v6))+8:], uint64(t32))
											t33 := int64(load64(m.memory[uint32(v14):]))
											store64(m.memory[uint32(v6):], uint64(t33))
											goto l13
										}
									}
								l17:
								}
								t38 := v10
								v6 = int32(uint32(v9) >> 25)
								m.memory[uint32(t38)] = byte(v6)
								m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(v6)
							}
						l13:
							v6 = v3 + i32(1)
							if v3 != v4 {
								goto l20
							}
							p39 := v7
							if uint32(v4) < uint32(i32(8)) {
								p39 = v4
							}
							v3 = p39
						}
					l7:
						store32(m.memory[int64(uint32(v0))+8:], uint32(v3-v2))
						goto l21
					}
					v8 = v8 + i32(1)
					p5 := v3
					if uint32(v8) > uint32(v3) {
						p5 = v8
					}
					v3 = p5
					if uint32(v3) < uint32(i32(15)) {
						goto l2
					}
					{
						if uint32(v3) > uint32(i32(0x1fffffff)) {
							m.fn28(i32(1271248), i32(57), i32(1271276))
							panic("unreachable")
						}
						t6 := int32(uint32(v3<<3) / uint32(i32(7)))
						v3 = i32_shr_u(i32(-1), int32(bits.LeadingZeros32(uint32(t6+i32(-1)))))
						if uint32(v3) > uint32(i32(0xffffffe)) {
							goto l4
						}
						v3 = v3 + i32(1)
						goto l5
					}
				}
			}
		l2:
			p40 := v3&i32(8) + i32(8)
			if uint32(v3) < uint32(i32(4)) {
				p40 = i32(4)
			}
			v3 = p40
		}
	l5:
		v8 = v3 + i32(8)
		t41 := v8
		v10 = v3 << 4
		v6 = t41 + v10
		if uint32(v6) < uint32(v8) {
			goto l4
		}
		if uint32(v6) > uint32(i32(0x7ffffff8)) {
			goto l4
		}
		{
			t42 := m.fn5(v6)
			v5 = t42
			if v5 != 0 {
				v6 = v5 + v10
				if v8 == 0 {
					goto l23
				}
				memory_fill(m.memory, uint32(v6), i32(255), uint32(v8))
			l23:
				v5 = v3 + i32(-1)
				p43 := int32(uint32(v3)>>3) * i32(7)
				if uint32(v3) < uint32(i32(9)) {
					p43 = v5
				}
				v15 = p43
				t44 := int32(load32(m.memory[uint32(v0):]))
				v9 = t44
				{
					if v2 == 0 {
						goto l24
					}
					t45 := int64(load64(m.memory[uint32(v9):]))
					v11 = (t45 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
					t46 := int64(load64(m.memory[int64(uint32(v1))+8:]))
					v16 = t46
					t47 := int64(load64(m.memory[uint32(v1):]))
					v17 = t47
					v8 = v9
					v1 = v2
					v3 = i32(0)
				l30:
					{
						if v11 != i64(0) {
							goto l25
						}
					l26:
						{
							v3 = v3 + i32(8)
							v8 = v8 + i32(8)
							t48 := int64(load64(m.memory[uint32(v8):]))
							v11 = t48 & i64(-0x7f7f7f7f7f7f7f80)
							if v11 == i64(-0x7f7f7f7f7f7f7f80) {
								goto l26
							}
						}
						v11 = v11 ^ i64(-0x7f7f7f7f7f7f7f80)
					l25:
						{
							t49 := v6
							t50 := v5
							t51 := v17
							t52 := v16
							t53 := v9
							v14 = int32(uint32(int64(bits.TrailingZeros64(uint64(v11))))>>3) + v3
							t54 := int32(load16(m.memory[uint32(t53-v14<<4+i32(-16)):]))
							t55 := m.fn106(t51, t52, t54)
							v7 = int32(t55)
							v10 = t50 & v7
							t56 := int64(load64(m.memory[uint32(t49+v10):]))
							v12 = t56 & i64(-0x7f7f7f7f7f7f7f80)
							if v12 != i64(0) {
								goto l27
							}
							v18 = i32(8)
						l28:
							{
								v10 = v10 + v18
								v18 = v18 + i32(8)
								t57 := v6
								v10 = v10 & v5
								t58 := int64(load64(m.memory[uint32(t57+v10):]))
								v12 = t58 & i64(-0x7f7f7f7f7f7f7f80)
								if v12 == 0 {
									goto l28
								}
							}
						}
					l27:
						v13 = v11 + i64(-1)
						{
							t59 := v6
							v10 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v12))))>>3) + v10) & v5
							t60 := int32(int8(m.memory[uint32(t59+v10)]))
							if t60 < i32(0) {
								goto l29
							}
							t61 := int64(load64(m.memory[uint32(v6):]))
							v10 = int32(uint32(int64(bits.TrailingZeros64(uint64(t61&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
						}
					l29:
						v11 = v13 & v11
						t62 := v6 + v10
						v7 = int32(uint32(v7) >> 25)
						m.memory[uint32(t62)] = byte(v7)
						m.memory[uint32(v6+(v10+i32(-8))&v5+i32(8))] = byte(v7)
						v10 = v6 + (v10^i32(-1))<<4
						t63 := v10
						v14 = v9 + (v14^i32(-1))<<4
						t64 := int64(load64(m.memory[int64(uint32(v14))+8:]))
						store64(m.memory[int64(uint32(t63))+8:], uint64(t64))
						t65 := int64(load64(m.memory[uint32(v14):]))
						store64(m.memory[uint32(v10):], uint64(t65))
						v1 = v1 + i32(-1)
						if v1 != 0 {
							goto l30
						}
					}
				}
			l24:
				store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
				store32(m.memory[uint32(v0):], uint32(v6))
				store32(m.memory[int64(uint32(v0))+8:], uint32(v15-v2))
				if v4 == 0 {
					goto l21
				}
				t66 := v4
				v8 = (v4<<4 + i32(23)) & i32(-16)
				v3 = t66 + v8 + i32(9)
				if v3 == 0 {
					goto l21
				}
				v4 = v9 - v8
				t67 := int32(load32(m.memory[uint32(v4+i32(-4)):]))
				v8 = t67
				v6 = v8 & i32(-8)
				t68 := v6
				v8 = v8 & i32(3)
				p69 := i32(8)
				if v8 != 0 {
					p69 = i32(4)
				}
				if uint32(t68) < uint32(p69+v3) {
					m.fn3(i32(1273840), i32(46), i32(1273888))
					panic("unreachable")
				}
				if v8 == 0 {
					goto l32
				}
				if uint32(v6) > uint32(v3+i32(39)) {
					m.fn3(i32(1273904), i32(46), i32(1273952))
					panic("unreachable")
				}
			l32:
				m.fn1(v4)
				return i32(-1)
			}
			m.fn24(i32(8), v6)
			panic("unreachable")
		}
	}
l21:
	return i32(-1)
l4:
	m.fn28(i32(1271248), i32(57), i32(1271276))
	panic("unreachable")
}
func (m *Module) fn106(v0, v1 int64, v2 int32) int64 {
	var v3 int32
	var v4, v5, v6, v7 int64
	t0 := m.g0
	v3 = t0 - i32(80)
	m.g0 = v3
	store64(m.memory[int64(uint32(v3))+56:], uint64(i64(0)))
	store64(m.memory[int64(uint32(v3))+64:], uint64(i64(0)))
	store64(m.memory[int64(uint32(v3))+48:], uint64(v1))
	store64(m.memory[int64(uint32(v3))+32:], uint64(v1^i64(8387220255154660723)))
	store64(m.memory[int64(uint32(v3))+24:], uint64(v1^i64(7237128888997146477)))
	store64(m.memory[int64(uint32(v3))+40:], uint64(v0))
	store64(m.memory[int64(uint32(v3))+16:], uint64(v0^i64(0x6c7967656e657261)))
	store64(m.memory[int64(uint32(v3))+8:], uint64(v0^i64(8317987319222330741)))
	store16(m.memory[int64(uint32(v3))+78:], uint16(v2))
	m.fn59(v3+i32(8), v3+i32(78), i32(2))
	t1 := int64(load64(m.memory[int64(uint32(v3))+8:]))
	v0 = t1
	t2 := int64(load64(m.memory[int64(uint32(v3))+24:]))
	v1 = t2
	t3 := int64(load32(m.memory[int64(uint32(v3))+64:]))
	v4 = t3
	t4 := int64(load64(m.memory[int64(uint32(v3))+56:]))
	v5 = t4
	t5 := int64(load64(m.memory[int64(uint32(v3))+32:]))
	v6 = t5
	t6 := int64(load64(m.memory[int64(uint32(v3))+16:]))
	v7 = t6
	m.g0 = v3 + i32(80)
	t7 := v6
	v4 = v5 | v4<<56
	v5 = t7 ^ v4
	t8 := i64_rotl(v5, i64(16))
	v5 = v5 + v7
	v6 = t8 ^ v5
	t9 := i64_rotl(v6, i64(21))
	t10 := v6
	v0 = v1 + v0
	v6 = t10 + i64_rotl(v0, i64(32))
	v7 = t9 ^ v6
	t11 := i64_rotl(v7, i64(16))
	t12 := v7
	t13 := v5
	v1 = i64_rotl(v1, i64(13)) ^ v0
	v0 = t13 + v1
	v5 = t12 + (i64_rotl(v0, i64(32)) ^ i64(255))
	v7 = t11 ^ v5
	t14 := i64_rotl(v7, i64(21))
	t15 := v7
	t16 := v6 ^ v4
	v1 = v0 ^ i64_rotl(v1, i64(17))
	v0 = t16 + v1
	v4 = t15 + i64_rotl(v0, i64(32))
	v6 = t14 ^ v4
	t17 := i64_rotl(v6, i64(16))
	t18 := v6
	v1 = v0 ^ i64_rotl(v1, i64(13))
	v0 = v1 + v5
	v5 = t18 + i64_rotl(v0, i64(32))
	v6 = t17 ^ v5
	t19 := i64_rotl(v6, i64(21))
	t20 := v6
	v1 = v0 ^ i64_rotl(v1, i64(17))
	v0 = v1 + v4
	v4 = t20 + i64_rotl(v0, i64(32))
	v6 = t19 ^ v4
	t21 := i64_rotl(v6, i64(16))
	t22 := v6
	v1 = i64_rotl(v1, i64(13)) ^ v0
	v0 = v1 + v5
	v5 = t22 + i64_rotl(v0, i64(32))
	t23 := i64_rotl(t21^v5, i64(21))
	v1 = i64_rotl(v1, i64(17)) ^ v0
	v1 = i64_rotl(v1, i64(13)) ^ (v1 + v4)
	t24 := t23 ^ i64_rotl(v1, i64(17))
	v1 = v1 + v5
	return t24 ^ i64_rotl(v1, i64(32)) ^ v1
}
func (m *Module) fn107(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7, v8, v9, v10 int32
	var v11, v12, v13 int64
	var v14, v15 int32
	var v16, v17 int64
	var v18 int32
	{
		{
			{
				t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
				v2 = t0
				v3 = v2 + i32(1)
				if v3 == 0 {
					m.fn28(i32(1271248), i32(57), i32(1271276))
					panic("unreachable")
				}
				{
					t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
					t2 := v3
					v4 = t1
					t3 := v4
					v5 = v4 + i32(1)
					v6 = int32(uint32(v5) >> 3)
					v7 = v6 * i32(7)
					p4 := v7
					if uint32(v4) < uint32(i32(8)) {
						p4 = t3
					}
					v8 = p4
					if uint32(t2) <= uint32(int32(uint32(v8)>>1)) {
						{
							if v5 != 0 {
								goto l6
							}
							v3 = i32(0)
							goto l7
						l6:
							t7 := int32(load32(m.memory[uint32(v0):]))
							v8 = t7
							v3 = i32(0)
							{
								{
									t8 := v6
									var p9 int32
									if v5&i32(7) != i32(0) {
										p9 = 1
									}
									v6 = t8 + p9
									if v6 == i32(1) {
										goto l8
									}
									v9 = v6 & i32(1)
									v10 = v6 & i32(0x3ffffffe)
									v3 = i32(0)
								l9:
									{
										v6 = v8 + v3
										t10 := int64(load64(m.memory[uint32(v6):]))
										t11 := v6
										v11 = t10
										store64(m.memory[uint32(t11):], uint64(int64(uint64(v11^i64(-1))>>7)&i64(72340172838076673)+(v11|i64(0x7f7f7f7f7f7f7f7f))))
										v6 = v6 + i32(8)
										t12 := int64(load64(m.memory[uint32(v6):]))
										t13 := v6
										v11 = t12
										store64(m.memory[uint32(t13):], uint64(int64(uint64(v11^i64(-1))>>7)&i64(72340172838076673)+(v11|i64(0x7f7f7f7f7f7f7f7f))))
										v3 = v3 + i32(16)
										v10 = v10 + i32(-2)
										if v10 != 0 {
											goto l9
										}
									}
									if v9 == 0 {
										goto l10
									}
								}
							l8:
								v3 = v8 + v3
								t14 := int64(load64(m.memory[uint32(v3):]))
								t15 := v3
								v11 = t14
								store64(m.memory[uint32(t15):], uint64(int64(uint64(v11^i64(-1))>>7)&i64(72340172838076673)+(v11|i64(0x7f7f7f7f7f7f7f7f))))
							}
						l10:
							{
								if uint32(v5) < uint32(i32(8)) {
									goto l11
								}
								t16 := int64(load64(m.memory[uint32(v8):]))
								store64(m.memory[uint32(v8+v5):], uint64(t16))
								goto l12
							}
						l11:
							if v5 == 0 {
								goto l12
							}
							memory_copy(m.memory, uint32(v8+i32(8)), uint32(v8), uint32(v5))
						l12:
							t17 := int64(load64(m.memory[int64(uint32(v1))+8:]))
							v12 = t17
							t18 := int64(load64(m.memory[uint32(v1):]))
							v13 = t18
							v6 = i32(0)
						l20:
							{
								t19 := v8
								v3 = v6
								v10 = t19 + v3
								t20 := int32(m.memory[uint32(v10)])
								if t20 != i32(128) {
									goto l13
								}
								v1 = v8 - v3<<2 + i32(-4)
								v14 = v8 + (v3^i32(-1))<<2
							l19:
								{
									t21 := int32(load16(m.memory[uint32(v1):]))
									t22 := m.fn106(v13, v12, t21)
									t23 := v4
									v9 = int32(t22)
									v6 = t23 & v9
									v5 = v6
									{
										t24 := int64(load64(m.memory[uint32(v8+v6):]))
										v11 = t24 & i64(-0x7f7f7f7f7f7f7f80)
										if v11 != i64(0) {
											goto l14
										}
										v15 = i32(8)
										v5 = v6
									l15:
										{
											v5 = v5 + v15
											v15 = v15 + i32(8)
											t25 := v8
											v5 = v5 & v4
											t26 := int64(load64(m.memory[uint32(t25+v5):]))
											v11 = t26 & i64(-0x7f7f7f7f7f7f7f80)
											if v11 == 0 {
												goto l15
											}
										}
									}
								l14:
									{
										t27 := v8
										v5 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v11))))>>3) + v5) & v4
										t28 := int32(int8(m.memory[uint32(t27+v5)]))
										if t28 < i32(0) {
											goto l16
										}
										t29 := int64(load64(m.memory[uint32(v8):]))
										v5 = int32(uint32(int64(bits.TrailingZeros64(uint64(t29&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
									}
								l16:
									{
										if uint32((v5-v6^(v3-v6))&v4) < uint32(i32(8)) {
											goto l17
										}
										v6 = v8 + v5
										t30 := int32(m.memory[uint32(v6)])
										v15 = t30
										t31 := v6
										v9 = int32(uint32(v9) >> 25)
										m.memory[uint32(t31)] = byte(v9)
										m.memory[uint32(v8+(v5+i32(-8))&v4+i32(8))] = byte(v9)
										v6 = v8 + (v5^i32(-1))<<2
										{
											if v15 != i32(255) {
												t33 := int32(load32(m.memory[uint32(v14):]))
												v5 = t33
												t34 := int32(load32(m.memory[uint32(v6):]))
												store32(m.memory[uint32(v14):], uint32(t34))
												store32(m.memory[uint32(v6):], uint32(v5))
												goto l19
											}
											m.memory[uint32(v10)] = byte(i32(255))
											m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(i32(255))
											t32 := int32(load32(m.memory[uint32(v14):]))
											store32(m.memory[uint32(v6):], uint32(t32))
											goto l13
										}
									}
								l17:
								}
								t35 := v10
								v6 = int32(uint32(v9) >> 25)
								m.memory[uint32(t35)] = byte(v6)
								m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(v6)
							}
						l13:
							v6 = v3 + i32(1)
							if v3 != v4 {
								goto l20
							}
							p36 := v7
							if uint32(v4) < uint32(i32(8)) {
								p36 = v4
							}
							v3 = p36
						}
					l7:
						store32(m.memory[int64(uint32(v0))+8:], uint32(v3-v2))
						goto l21
					}
					v8 = v8 + i32(1)
					p5 := v3
					if uint32(v8) > uint32(v3) {
						p5 = v8
					}
					v3 = p5
					if uint32(v3) < uint32(i32(15)) {
						goto l2
					}
					{
						if uint32(v3) > uint32(i32(0x1fffffff)) {
							m.fn28(i32(1271248), i32(57), i32(1271276))
							panic("unreachable")
						}
						t6 := int32(uint32(v3<<3) / uint32(i32(7)))
						v3 = i32_shr_u(i32(-1), int32(bits.LeadingZeros32(uint32(t6+i32(-1)))))
						if uint32(v3) > uint32(i32(0x3ffffffe)) {
							goto l4
						}
						v3 = v3 + i32(1)
						goto l5
					}
				}
			}
		l2:
			p37 := v3&i32(8) + i32(8)
			if uint32(v3) < uint32(i32(4)) {
				p37 = i32(4)
			}
			v3 = p37
		}
	l5:
		v8 = v3 + i32(8)
		t38 := v8
		v10 = (v3<<2 + i32(7)) & i32(-8)
		v6 = t38 + v10
		if uint32(v6) < uint32(v8) {
			goto l4
		}
		if uint32(v6) > uint32(i32(0x7ffffff8)) {
			goto l4
		}
		{
			t39 := m.fn5(v6)
			v5 = t39
			if v5 != 0 {
				v6 = v5 + v10
				if v8 == 0 {
					goto l23
				}
				memory_fill(m.memory, uint32(v6), i32(255), uint32(v8))
			l23:
				v5 = v3 + i32(-1)
				p40 := int32(uint32(v3)>>3) * i32(7)
				if uint32(v3) < uint32(i32(9)) {
					p40 = v5
				}
				v15 = p40
				t41 := int32(load32(m.memory[uint32(v0):]))
				v9 = t41
				{
					if v2 == 0 {
						goto l24
					}
					t42 := int64(load64(m.memory[uint32(v9):]))
					v11 = (t42 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
					t43 := int64(load64(m.memory[int64(uint32(v1))+8:]))
					v16 = t43
					t44 := int64(load64(m.memory[uint32(v1):]))
					v17 = t44
					v8 = v9
					v1 = v2
					v3 = i32(0)
				l30:
					{
						if v11 != i64(0) {
							goto l25
						}
					l26:
						{
							v3 = v3 + i32(8)
							v8 = v8 + i32(8)
							t45 := int64(load64(m.memory[uint32(v8):]))
							v11 = t45 & i64(-0x7f7f7f7f7f7f7f80)
							if v11 == i64(-0x7f7f7f7f7f7f7f80) {
								goto l26
							}
						}
						v11 = v11 ^ i64(-0x7f7f7f7f7f7f7f80)
					l25:
						{
							t46 := v6
							t47 := v5
							t48 := v17
							t49 := v16
							t50 := v9
							v14 = int32(uint32(int64(bits.TrailingZeros64(uint64(v11))))>>3) + v3
							t51 := int32(load16(m.memory[uint32(t50-v14<<2+i32(-4)):]))
							t52 := m.fn106(t48, t49, t51)
							v7 = int32(t52)
							v10 = t47 & v7
							t53 := int64(load64(m.memory[uint32(t46+v10):]))
							v12 = t53 & i64(-0x7f7f7f7f7f7f7f80)
							if v12 != i64(0) {
								goto l27
							}
							v18 = i32(8)
						l28:
							{
								v10 = v10 + v18
								v18 = v18 + i32(8)
								t54 := v6
								v10 = v10 & v5
								t55 := int64(load64(m.memory[uint32(t54+v10):]))
								v12 = t55 & i64(-0x7f7f7f7f7f7f7f80)
								if v12 == 0 {
									goto l28
								}
							}
						}
					l27:
						v13 = v11 + i64(-1)
						{
							t56 := v6
							v10 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v12))))>>3) + v10) & v5
							t57 := int32(int8(m.memory[uint32(t56+v10)]))
							if t57 < i32(0) {
								goto l29
							}
							t58 := int64(load64(m.memory[uint32(v6):]))
							v10 = int32(uint32(int64(bits.TrailingZeros64(uint64(t58&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
						}
					l29:
						v11 = v13 & v11
						t59 := v6 + v10
						v7 = int32(uint32(v7) >> 25)
						m.memory[uint32(t59)] = byte(v7)
						m.memory[uint32(v6+(v10+i32(-8))&v5+i32(8))] = byte(v7)
						t60 := int32(load32(m.memory[uint32(v9+(v14^i32(-1))<<2):]))
						store32(m.memory[uint32(v6+(v10^i32(-1))<<2):], uint32(t60))
						v1 = v1 + i32(-1)
						if v1 != 0 {
							goto l30
						}
					}
				}
			l24:
				store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
				store32(m.memory[uint32(v0):], uint32(v6))
				store32(m.memory[int64(uint32(v0))+8:], uint32(v15-v2))
				if v4 == 0 {
					goto l21
				}
				t61 := v4
				v8 = (v4<<2 + i32(11)) & i32(-8)
				v3 = t61 + v8 + i32(9)
				if v3 == 0 {
					goto l21
				}
				v4 = v9 - v8
				t62 := int32(load32(m.memory[uint32(v4+i32(-4)):]))
				v8 = t62
				v6 = v8 & i32(-8)
				t63 := v6
				v8 = v8 & i32(3)
				p64 := i32(8)
				if v8 != 0 {
					p64 = i32(4)
				}
				if uint32(t63) < uint32(p64+v3) {
					m.fn3(i32(1273840), i32(46), i32(1273888))
					panic("unreachable")
				}
				if v8 == 0 {
					goto l32
				}
				if uint32(v6) > uint32(v3+i32(39)) {
					m.fn3(i32(1273904), i32(46), i32(1273952))
					panic("unreachable")
				}
			l32:
				m.fn1(v4)
				return i32(-1)
			}
			m.fn24(i32(8), v6)
			panic("unreachable")
		}
	}
l21:
	return i32(-1)
l4:
	m.fn28(i32(1271248), i32(57), i32(1271276))
	panic("unreachable")
}
func (m *Module) fn108(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7, v8 int32
	var v9 int64
	var v10, v11, v12, v13 int32
	var v14, v15, v16 int64
	var v17 int32
	var v18 int64
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
		v2 = t0
		v3 = v2 + i32(1)
		if v3 == 0 {
			m.fn28(i32(1271248), i32(57), i32(1271276))
			panic("unreachable")
		}
		{
			t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t2 := v3
			v4 = t1
			t3 := v4
			v5 = v4 + i32(1)
			v6 = int32(uint32(v5) >> 3)
			v7 = v6 * i32(7)
			p4 := v7
			if uint32(v4) < uint32(i32(8)) {
				p4 = t3
			}
			v8 = p4
			if uint32(t2) <= uint32(int32(uint32(v8)>>1)) {
				goto l1
			}
			{
				{
					{
						v8 = v8 + i32(1)
						p5 := v3
						if uint32(v8) > uint32(v3) {
							p5 = v8
						}
						v3 = p5
						if uint32(v3) < uint32(i32(15)) {
							goto l2
						}
						if uint32(v3) > uint32(i32(0x1fffffff)) {
							m.fn28(i32(1271248), i32(57), i32(1271276))
							panic("unreachable")
						}
						t6 := int32(uint32(v3<<3) / uint32(i32(7)))
						v3 = i32_shr_u(i32(-1), int32(bits.LeadingZeros32(uint32(t6+i32(-1))))) + i32(1)
						goto l4
					}
				l2:
					p7 := v3&i32(8) + i32(8)
					if uint32(v3) < uint32(i32(4)) {
						p7 = i32(4)
					}
					v3 = p7
				}
			l4:
				v9 = int64(uint32(v3)) * i64(60)
				if int32(int64(uint64(v9)>>32)) != 0 {
					goto l5
				}
				v6 = int32(v9)
				if uint32(v6) > uint32(i32(-8)) {
					goto l5
				}
				v8 = v3 + i32(8)
				t8 := v8
				v10 = (v6 + i32(7)) & i32(-8)
				v6 = t8 + v10
				if uint32(v6) < uint32(v8) {
					goto l5
				}
				if uint32(v6) > uint32(i32(0x7ffffff8)) {
					goto l5
				}
				t9 := m.fn5(v6)
				v5 = t9
				if v5 != 0 {
					v5 = v5 + v10
					if v8 == 0 {
						goto l7
					}
					memory_fill(m.memory, uint32(v5), i32(255), uint32(v8))
				l7:
					v11 = v3 + i32(-1)
					p10 := int32(uint32(v3)>>3) * i32(7)
					if uint32(v3) < uint32(i32(9)) {
						p10 = v11
					}
					v12 = p10
					t11 := int32(load32(m.memory[uint32(v0):]))
					v13 = t11
					{
						if v2 == 0 {
							goto l8
						}
						t12 := int64(load64(m.memory[uint32(v13):]))
						v9 = (t12 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
						t13 := int64(load64(m.memory[int64(uint32(v1))+8:]))
						v14 = t13
						t14 := int64(load64(m.memory[uint32(v1):]))
						v15 = t14
						v8 = v13
						v1 = v2
						v3 = i32(0)
					l14:
						{
							if v9 != i64(0) {
								goto l9
							}
						l10:
							{
								v3 = v3 + i32(8)
								v8 = v8 + i32(8)
								t15 := int64(load64(m.memory[uint32(v8):]))
								v9 = t15 & i64(-0x7f7f7f7f7f7f7f80)
								if v9 == i64(-0x7f7f7f7f7f7f7f80) {
									goto l10
								}
							}
							v9 = v9 ^ i64(-0x7f7f7f7f7f7f7f80)
						l9:
							{
								t16 := v5
								t17 := v11
								t18 := v15
								t19 := v14
								t20 := v13
								v10 = int32(uint32(int64(bits.TrailingZeros64(uint64(v9))))>>3) + v3
								t21 := int32(load16(m.memory[uint32(t20+(i32(0)-v10)*i32(60)+i32(-60)):]))
								t22 := m.fn106(t18, t19, t21)
								v7 = int32(t22)
								v6 = t17 & v7
								t23 := int64(load64(m.memory[uint32(t16+v6):]))
								v16 = t23 & i64(-0x7f7f7f7f7f7f7f80)
								if v16 != i64(0) {
									goto l11
								}
								v17 = i32(8)
							l12:
								{
									v6 = v6 + v17
									v17 = v17 + i32(8)
									t24 := v5
									v6 = v6 & v11
									t25 := int64(load64(m.memory[uint32(t24+v6):]))
									v16 = t25 & i64(-0x7f7f7f7f7f7f7f80)
									if v16 == 0 {
										goto l12
									}
								}
							}
						l11:
							v18 = v9 + i64(-1)
							{
								t26 := v5
								v6 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v16))))>>3) + v6) & v11
								t27 := int32(int8(m.memory[uint32(t26+v6)]))
								if t27 < i32(0) {
									goto l13
								}
								t28 := int64(load64(m.memory[uint32(v5):]))
								v6 = int32(uint32(int64(bits.TrailingZeros64(uint64(t28&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
							}
						l13:
							v9 = v18 & v9
							t29 := v5 + v6
							v7 = int32(uint32(v7) >> 25)
							m.memory[uint32(t29)] = byte(v7)
							m.memory[uint32(v5+(v6+i32(-8))&v11+i32(8))] = byte(v7)
							v6 = v5 + (v6^i32(-1))*i32(60)
							t30 := v6
							v10 = v13 + (v10^i32(-1))*i32(60)
							t31 := int32(load32(m.memory[int64(uint32(v10))+56:]))
							store32(m.memory[int64(uint32(t30))+56:], uint32(t31))
							t32 := int64(load64(m.memory[int64(uint32(v10))+48:]))
							store64(m.memory[int64(uint32(v6))+48:], uint64(t32))
							t33 := int64(load64(m.memory[int64(uint32(v10))+40:]))
							store64(m.memory[int64(uint32(v6))+40:], uint64(t33))
							t34 := int64(load64(m.memory[int64(uint32(v10))+32:]))
							store64(m.memory[int64(uint32(v6))+32:], uint64(t34))
							t35 := int64(load64(m.memory[int64(uint32(v10))+24:]))
							store64(m.memory[int64(uint32(v6))+24:], uint64(t35))
							t36 := int64(load64(m.memory[int64(uint32(v10))+16:]))
							store64(m.memory[int64(uint32(v6))+16:], uint64(t36))
							t37 := int64(load64(m.memory[int64(uint32(v10))+8:]))
							store64(m.memory[int64(uint32(v6))+8:], uint64(t37))
							t38 := int64(load64(m.memory[uint32(v10):]))
							store64(m.memory[uint32(v6):], uint64(t38))
							v1 = v1 + i32(-1)
							if v1 != 0 {
								goto l14
							}
						}
					}
				l8:
					store32(m.memory[int64(uint32(v0))+4:], uint32(v11))
					store32(m.memory[uint32(v0):], uint32(v5))
					store32(m.memory[int64(uint32(v0))+8:], uint32(v12-v2))
					if v4 == 0 {
						goto l15
					}
					t39 := v4
					v8 = (v4*i32(60) + i32(67)) & i32(-8)
					v3 = t39 + v8 + i32(9)
					if v3 == 0 {
						goto l15
					}
					v6 = v13 - v8
					t40 := int32(load32(m.memory[uint32(v6+i32(-4)):]))
					v8 = t40
					v10 = v8 & i32(-8)
					t41 := v10
					v8 = v8 & i32(3)
					p42 := i32(8)
					if v8 != 0 {
						p42 = i32(4)
					}
					if uint32(t41) < uint32(p42+v3) {
						m.fn3(i32(1273840), i32(46), i32(1273888))
						panic("unreachable")
					}
					if v8 == 0 {
						goto l17
					}
					if uint32(v10) > uint32(v3+i32(39)) {
						m.fn3(i32(1273904), i32(46), i32(1273952))
						panic("unreachable")
					}
				l17:
					m.fn1(v6)
					return i32(-1)
				}
				m.fn24(i32(8), v6)
				panic("unreachable")
			}
		}
	l1:
		{
			if v5 != 0 {
				goto l19
			}
			v3 = i32(0)
			goto l20
		l19:
			t43 := int32(load32(m.memory[uint32(v0):]))
			v8 = t43
			v3 = i32(0)
			{
				{
					t44 := v6
					var p45 int32
					if v5&i32(7) != i32(0) {
						p45 = 1
					}
					v6 = t44 + p45
					if v6 == i32(1) {
						goto l21
					}
					v11 = v6 & i32(1)
					v10 = v6 & i32(0x3ffffffe)
					v3 = i32(0)
				l22:
					{
						v6 = v8 + v3
						t46 := int64(load64(m.memory[uint32(v6):]))
						t47 := v6
						v9 = t46
						store64(m.memory[uint32(t47):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
						v6 = v6 + i32(8)
						t48 := int64(load64(m.memory[uint32(v6):]))
						t49 := v6
						v9 = t48
						store64(m.memory[uint32(t49):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
						v3 = v3 + i32(16)
						v10 = v10 + i32(-2)
						if v10 != 0 {
							goto l22
						}
					}
					if v11 == 0 {
						goto l23
					}
				}
			l21:
				v3 = v8 + v3
				t50 := int64(load64(m.memory[uint32(v3):]))
				t51 := v3
				v9 = t50
				store64(m.memory[uint32(t51):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
			}
		l23:
			{
				if uint32(v5) < uint32(i32(8)) {
					goto l24
				}
				t52 := int64(load64(m.memory[uint32(v8):]))
				store64(m.memory[uint32(v8+v5):], uint64(t52))
				goto l25
			}
		l24:
			if v5 == 0 {
				goto l25
			}
			memory_copy(m.memory, uint32(v8+i32(8)), uint32(v8), uint32(v5))
		l25:
			t53 := int64(load64(m.memory[int64(uint32(v1))+8:]))
			v16 = t53
			t54 := int64(load64(m.memory[uint32(v1):]))
			v18 = t54
			v6 = i32(0)
		l33:
			{
				t55 := v8
				v3 = v6
				v10 = t55 + v3
				t56 := int32(m.memory[uint32(v10)])
				if t56 != i32(128) {
					goto l26
				}
				v1 = v8 + (v3^i32(-1))*i32(60)
				v13 = v8 + (i32(0)-v3)*i32(60) + i32(-60)
			l32:
				{
					t57 := int32(load16(m.memory[uint32(v13):]))
					t58 := m.fn106(v18, v16, t57)
					t59 := v4
					v11 = int32(t58)
					v6 = t59 & v11
					v5 = v6
					{
						t60 := int64(load64(m.memory[uint32(v8+v6):]))
						v9 = t60 & i64(-0x7f7f7f7f7f7f7f80)
						if v9 != i64(0) {
							goto l27
						}
						v12 = i32(8)
						v5 = v6
					l28:
						{
							v5 = v5 + v12
							v12 = v12 + i32(8)
							t61 := v8
							v5 = v5 & v4
							t62 := int64(load64(m.memory[uint32(t61+v5):]))
							v9 = t62 & i64(-0x7f7f7f7f7f7f7f80)
							if v9 == 0 {
								goto l28
							}
						}
					}
				l27:
					{
						t63 := v8
						v5 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v9))))>>3) + v5) & v4
						t64 := int32(int8(m.memory[uint32(t63+v5)]))
						if t64 < i32(0) {
							goto l29
						}
						t65 := int64(load64(m.memory[uint32(v8):]))
						v5 = int32(uint32(int64(bits.TrailingZeros64(uint64(t65&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
					}
				l29:
					{
						if uint32((v5-v6^(v3-v6))&v4) < uint32(i32(8)) {
							goto l30
						}
						v6 = v8 + v5
						t66 := int32(m.memory[uint32(v6)])
						v12 = t66
						t67 := v6
						v11 = int32(uint32(v11) >> 25)
						m.memory[uint32(t67)] = byte(v11)
						m.memory[uint32(v8+(v5+i32(-8))&v4+i32(8))] = byte(v11)
						v6 = v8 + (v5^i32(-1))*i32(60)
						{
							if v12 != i32(255) {
								t76 := int32(load32(m.memory[uint32(v1):]))
								v5 = t76
								t77 := int32(load32(m.memory[uint32(v6):]))
								store32(m.memory[uint32(v1):], uint32(t77))
								store32(m.memory[uint32(v6):], uint32(v5))
								t78 := int32(load32(m.memory[int64(uint32(v6))+4:]))
								v5 = t78
								t79 := int32(load32(m.memory[int64(uint32(v1))+4:]))
								store32(m.memory[int64(uint32(v6))+4:], uint32(t79))
								store32(m.memory[int64(uint32(v1))+4:], uint32(v5))
								t80 := int32(load32(m.memory[int64(uint32(v1))+8:]))
								v5 = t80
								t81 := int32(load32(m.memory[int64(uint32(v6))+8:]))
								store32(m.memory[int64(uint32(v1))+8:], uint32(t81))
								store32(m.memory[int64(uint32(v6))+8:], uint32(v5))
								t82 := int32(load32(m.memory[int64(uint32(v6))+12:]))
								v5 = t82
								t83 := int32(load32(m.memory[int64(uint32(v1))+12:]))
								store32(m.memory[int64(uint32(v6))+12:], uint32(t83))
								store32(m.memory[int64(uint32(v1))+12:], uint32(v5))
								t84 := int32(load32(m.memory[int64(uint32(v1))+16:]))
								v5 = t84
								t85 := int32(load32(m.memory[int64(uint32(v6))+16:]))
								store32(m.memory[int64(uint32(v1))+16:], uint32(t85))
								store32(m.memory[int64(uint32(v6))+16:], uint32(v5))
								t86 := int32(load32(m.memory[int64(uint32(v6))+20:]))
								v5 = t86
								t87 := int32(load32(m.memory[int64(uint32(v1))+20:]))
								store32(m.memory[int64(uint32(v6))+20:], uint32(t87))
								store32(m.memory[int64(uint32(v1))+20:], uint32(v5))
								t88 := int32(load32(m.memory[int64(uint32(v1))+24:]))
								v5 = t88
								t89 := int32(load32(m.memory[int64(uint32(v6))+24:]))
								store32(m.memory[int64(uint32(v1))+24:], uint32(t89))
								store32(m.memory[int64(uint32(v6))+24:], uint32(v5))
								t90 := int32(load32(m.memory[int64(uint32(v6))+28:]))
								v5 = t90
								t91 := int32(load32(m.memory[int64(uint32(v1))+28:]))
								store32(m.memory[int64(uint32(v6))+28:], uint32(t91))
								store32(m.memory[int64(uint32(v1))+28:], uint32(v5))
								t92 := int32(load32(m.memory[int64(uint32(v1))+32:]))
								v5 = t92
								t93 := int32(load32(m.memory[int64(uint32(v6))+32:]))
								store32(m.memory[int64(uint32(v1))+32:], uint32(t93))
								store32(m.memory[int64(uint32(v6))+32:], uint32(v5))
								t94 := int32(load32(m.memory[int64(uint32(v6))+36:]))
								v5 = t94
								t95 := int32(load32(m.memory[int64(uint32(v1))+36:]))
								store32(m.memory[int64(uint32(v6))+36:], uint32(t95))
								store32(m.memory[int64(uint32(v1))+36:], uint32(v5))
								t96 := int32(load32(m.memory[int64(uint32(v1))+40:]))
								v5 = t96
								t97 := int32(load32(m.memory[int64(uint32(v6))+40:]))
								store32(m.memory[int64(uint32(v1))+40:], uint32(t97))
								store32(m.memory[int64(uint32(v6))+40:], uint32(v5))
								t98 := int32(load32(m.memory[int64(uint32(v6))+44:]))
								v5 = t98
								t99 := int32(load32(m.memory[int64(uint32(v1))+44:]))
								store32(m.memory[int64(uint32(v6))+44:], uint32(t99))
								store32(m.memory[int64(uint32(v1))+44:], uint32(v5))
								t100 := int32(load32(m.memory[int64(uint32(v1))+48:]))
								v5 = t100
								t101 := int32(load32(m.memory[int64(uint32(v6))+48:]))
								store32(m.memory[int64(uint32(v1))+48:], uint32(t101))
								store32(m.memory[int64(uint32(v6))+48:], uint32(v5))
								t102 := int32(load32(m.memory[int64(uint32(v6))+52:]))
								v5 = t102
								t103 := int32(load32(m.memory[int64(uint32(v1))+52:]))
								store32(m.memory[int64(uint32(v6))+52:], uint32(t103))
								store32(m.memory[int64(uint32(v1))+52:], uint32(v5))
								t104 := int32(load32(m.memory[int64(uint32(v1))+56:]))
								v5 = t104
								t105 := int32(load32(m.memory[int64(uint32(v6))+56:]))
								store32(m.memory[int64(uint32(v1))+56:], uint32(t105))
								store32(m.memory[int64(uint32(v6))+56:], uint32(v5))
								goto l32
							}
							m.memory[uint32(v10)] = byte(i32(255))
							m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(i32(255))
							t68 := int32(load32(m.memory[int64(uint32(v1))+56:]))
							store32(m.memory[int64(uint32(v6))+56:], uint32(t68))
							t69 := int64(load64(m.memory[int64(uint32(v1))+48:]))
							store64(m.memory[int64(uint32(v6))+48:], uint64(t69))
							t70 := int64(load64(m.memory[int64(uint32(v1))+40:]))
							store64(m.memory[int64(uint32(v6))+40:], uint64(t70))
							t71 := int64(load64(m.memory[int64(uint32(v1))+32:]))
							store64(m.memory[int64(uint32(v6))+32:], uint64(t71))
							t72 := int64(load64(m.memory[int64(uint32(v1))+24:]))
							store64(m.memory[int64(uint32(v6))+24:], uint64(t72))
							t73 := int64(load64(m.memory[int64(uint32(v1))+16:]))
							store64(m.memory[int64(uint32(v6))+16:], uint64(t73))
							t74 := int64(load64(m.memory[int64(uint32(v1))+8:]))
							store64(m.memory[int64(uint32(v6))+8:], uint64(t74))
							t75 := int64(load64(m.memory[uint32(v1):]))
							store64(m.memory[uint32(v6):], uint64(t75))
							goto l26
						}
					}
				l30:
				}
				t106 := v10
				v6 = int32(uint32(v11) >> 25)
				m.memory[uint32(t106)] = byte(v6)
				m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(v6)
			}
		l26:
			v6 = v3 + i32(1)
			if v3 != v4 {
				goto l33
			}
			p107 := v7
			if uint32(v4) < uint32(i32(8)) {
				p107 = v4
			}
			v3 = p107
		}
	l20:
		store32(m.memory[int64(uint32(v0))+8:], uint32(v3-v2))
	l15:
		return i32(-1)
	}
l5:
	m.fn28(i32(1271248), i32(57), i32(1271276))
	panic("unreachable")
}
func (m *Module) fn109(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7, v8 int32
	var v9 int64
	var v10, v11, v12, v13 int32
	var v14, v15, v16 int64
	var v17 int32
	var v18 int64
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
		v2 = t0
		v3 = v2 + i32(1)
		if v3 == 0 {
			m.fn28(i32(1271248), i32(57), i32(1271276))
			panic("unreachable")
		}
		{
			t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t2 := v3
			v4 = t1
			t3 := v4
			v5 = v4 + i32(1)
			v6 = int32(uint32(v5) >> 3)
			v7 = v6 * i32(7)
			p4 := v7
			if uint32(v4) < uint32(i32(8)) {
				p4 = t3
			}
			v8 = p4
			if uint32(t2) <= uint32(int32(uint32(v8)>>1)) {
				goto l1
			}
			{
				{
					{
						v8 = v8 + i32(1)
						p5 := v3
						if uint32(v8) > uint32(v3) {
							p5 = v8
						}
						v3 = p5
						if uint32(v3) < uint32(i32(15)) {
							goto l2
						}
						if uint32(v3) > uint32(i32(0x1fffffff)) {
							m.fn28(i32(1271248), i32(57), i32(1271276))
							panic("unreachable")
						}
						t6 := int32(uint32(v3<<3) / uint32(i32(7)))
						v3 = i32_shr_u(i32(-1), int32(bits.LeadingZeros32(uint32(t6+i32(-1))))) + i32(1)
						goto l4
					}
				l2:
					p7 := v3&i32(8) + i32(8)
					if uint32(v3) < uint32(i32(4)) {
						p7 = i32(4)
					}
					v3 = p7
				}
			l4:
				v9 = int64(uint32(v3)) * i64(36)
				if int32(int64(uint64(v9)>>32)) != 0 {
					goto l5
				}
				v6 = int32(v9)
				if uint32(v6) > uint32(i32(-8)) {
					goto l5
				}
				v8 = v3 + i32(8)
				t8 := v8
				v10 = (v6 + i32(7)) & i32(-8)
				v6 = t8 + v10
				if uint32(v6) < uint32(v8) {
					goto l5
				}
				if uint32(v6) > uint32(i32(0x7ffffff8)) {
					goto l5
				}
				t9 := m.fn5(v6)
				v5 = t9
				if v5 != 0 {
					v6 = v5 + v10
					if v8 == 0 {
						goto l7
					}
					memory_fill(m.memory, uint32(v6), i32(255), uint32(v8))
				l7:
					v11 = v3 + i32(-1)
					p10 := int32(uint32(v3)>>3) * i32(7)
					if uint32(v3) < uint32(i32(9)) {
						p10 = v11
					}
					v12 = p10
					t11 := int32(load32(m.memory[uint32(v0):]))
					v13 = t11
					{
						if v2 == 0 {
							goto l8
						}
						t12 := int64(load64(m.memory[uint32(v13):]))
						v9 = (t12 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
						t13 := int64(load64(m.memory[int64(uint32(v1))+8:]))
						v14 = t13
						t14 := int64(load64(m.memory[uint32(v1):]))
						v15 = t14
						v8 = v13
						v1 = v2
						v3 = i32(0)
					l14:
						{
							if v9 != i64(0) {
								goto l9
							}
						l10:
							{
								v3 = v3 + i32(8)
								v8 = v8 + i32(8)
								t15 := int64(load64(m.memory[uint32(v8):]))
								v9 = t15 & i64(-0x7f7f7f7f7f7f7f80)
								if v9 == i64(-0x7f7f7f7f7f7f7f80) {
									goto l10
								}
							}
							v9 = v9 ^ i64(-0x7f7f7f7f7f7f7f80)
						l9:
							{
								t16 := v6
								t17 := v11
								t18 := v15
								t19 := v14
								t20 := v13
								v5 = int32(uint32(int64(bits.TrailingZeros64(uint64(v9))))>>3) + v3
								t21 := int32(load16(m.memory[uint32(t20+(i32(0)-v5)*i32(36)+i32(-36)):]))
								t22 := m.fn106(t18, t19, t21)
								v7 = int32(t22)
								v10 = t17 & v7
								t23 := int64(load64(m.memory[uint32(t16+v10):]))
								v16 = t23 & i64(-0x7f7f7f7f7f7f7f80)
								if v16 != i64(0) {
									goto l11
								}
								v17 = i32(8)
							l12:
								{
									v10 = v10 + v17
									v17 = v17 + i32(8)
									t24 := v6
									v10 = v10 & v11
									t25 := int64(load64(m.memory[uint32(t24+v10):]))
									v16 = t25 & i64(-0x7f7f7f7f7f7f7f80)
									if v16 == 0 {
										goto l12
									}
								}
							}
						l11:
							v18 = v9 + i64(-1)
							{
								t26 := v6
								v10 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v16))))>>3) + v10) & v11
								t27 := int32(int8(m.memory[uint32(t26+v10)]))
								if t27 < i32(0) {
									goto l13
								}
								t28 := int64(load64(m.memory[uint32(v6):]))
								v10 = int32(uint32(int64(bits.TrailingZeros64(uint64(t28&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
							}
						l13:
							v9 = v18 & v9
							t29 := v6 + v10
							v7 = int32(uint32(v7) >> 25)
							m.memory[uint32(t29)] = byte(v7)
							m.memory[uint32(v6+(v10+i32(-8))&v11+i32(8))] = byte(v7)
							v10 = v6 + (v10^i32(-1))*i32(36)
							t30 := v10
							v5 = v13 + (v5^i32(-1))*i32(36)
							t31 := int32(load32(m.memory[int64(uint32(v5))+32:]))
							store32(m.memory[int64(uint32(t30))+32:], uint32(t31))
							t32 := int64(load64(m.memory[int64(uint32(v5))+24:]))
							store64(m.memory[int64(uint32(v10))+24:], uint64(t32))
							t33 := int64(load64(m.memory[int64(uint32(v5))+16:]))
							store64(m.memory[int64(uint32(v10))+16:], uint64(t33))
							t34 := int64(load64(m.memory[int64(uint32(v5))+8:]))
							store64(m.memory[int64(uint32(v10))+8:], uint64(t34))
							t35 := int64(load64(m.memory[uint32(v5):]))
							store64(m.memory[uint32(v10):], uint64(t35))
							v1 = v1 + i32(-1)
							if v1 != 0 {
								goto l14
							}
						}
					}
				l8:
					store32(m.memory[int64(uint32(v0))+4:], uint32(v11))
					store32(m.memory[uint32(v0):], uint32(v6))
					store32(m.memory[int64(uint32(v0))+8:], uint32(v12-v2))
					if v4 == 0 {
						goto l15
					}
					t36 := v4
					v8 = (v4*i32(36) + i32(43)) & i32(-8)
					v3 = t36 + v8 + i32(9)
					if v3 == 0 {
						goto l15
					}
					v4 = v13 - v8
					t37 := int32(load32(m.memory[uint32(v4+i32(-4)):]))
					v8 = t37
					v6 = v8 & i32(-8)
					t38 := v6
					v8 = v8 & i32(3)
					p39 := i32(8)
					if v8 != 0 {
						p39 = i32(4)
					}
					if uint32(t38) < uint32(p39+v3) {
						m.fn3(i32(1273840), i32(46), i32(1273888))
						panic("unreachable")
					}
					if v8 == 0 {
						goto l17
					}
					if uint32(v6) > uint32(v3+i32(39)) {
						m.fn3(i32(1273904), i32(46), i32(1273952))
						panic("unreachable")
					}
				l17:
					m.fn1(v4)
					return i32(-1)
				}
				m.fn24(i32(8), v6)
				panic("unreachable")
			}
		}
	l1:
		{
			if v5 != 0 {
				goto l19
			}
			v3 = i32(0)
			goto l20
		l19:
			t40 := int32(load32(m.memory[uint32(v0):]))
			v8 = t40
			v3 = i32(0)
			{
				{
					t41 := v6
					var p42 int32
					if v5&i32(7) != i32(0) {
						p42 = 1
					}
					v6 = t41 + p42
					if v6 == i32(1) {
						goto l21
					}
					v11 = v6 & i32(1)
					v10 = v6 & i32(0x3ffffffe)
					v3 = i32(0)
				l22:
					{
						v6 = v8 + v3
						t43 := int64(load64(m.memory[uint32(v6):]))
						t44 := v6
						v9 = t43
						store64(m.memory[uint32(t44):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
						v6 = v6 + i32(8)
						t45 := int64(load64(m.memory[uint32(v6):]))
						t46 := v6
						v9 = t45
						store64(m.memory[uint32(t46):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
						v3 = v3 + i32(16)
						v10 = v10 + i32(-2)
						if v10 != 0 {
							goto l22
						}
					}
					if v11 == 0 {
						goto l23
					}
				}
			l21:
				v3 = v8 + v3
				t47 := int64(load64(m.memory[uint32(v3):]))
				t48 := v3
				v9 = t47
				store64(m.memory[uint32(t48):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
			}
		l23:
			{
				if uint32(v5) < uint32(i32(8)) {
					goto l24
				}
				t49 := int64(load64(m.memory[uint32(v8):]))
				store64(m.memory[uint32(v8+v5):], uint64(t49))
				goto l25
			}
		l24:
			if v5 == 0 {
				goto l25
			}
			memory_copy(m.memory, uint32(v8+i32(8)), uint32(v8), uint32(v5))
		l25:
			t50 := int64(load64(m.memory[int64(uint32(v1))+8:]))
			v16 = t50
			t51 := int64(load64(m.memory[uint32(v1):]))
			v18 = t51
			v6 = i32(0)
		l33:
			{
				t52 := v8
				v3 = v6
				v10 = t52 + v3
				t53 := int32(m.memory[uint32(v10)])
				if t53 != i32(128) {
					goto l26
				}
				v1 = v8 + (v3^i32(-1))*i32(36)
				v13 = v8 + (i32(0)-v3)*i32(36) + i32(-36)
			l32:
				{
					t54 := int32(load16(m.memory[uint32(v13):]))
					t55 := m.fn106(v18, v16, t54)
					t56 := v4
					v11 = int32(t55)
					v6 = t56 & v11
					v5 = v6
					{
						t57 := int64(load64(m.memory[uint32(v8+v6):]))
						v9 = t57 & i64(-0x7f7f7f7f7f7f7f80)
						if v9 != i64(0) {
							goto l27
						}
						v12 = i32(8)
						v5 = v6
					l28:
						{
							v5 = v5 + v12
							v12 = v12 + i32(8)
							t58 := v8
							v5 = v5 & v4
							t59 := int64(load64(m.memory[uint32(t58+v5):]))
							v9 = t59 & i64(-0x7f7f7f7f7f7f7f80)
							if v9 == 0 {
								goto l28
							}
						}
					}
				l27:
					{
						t60 := v8
						v5 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v9))))>>3) + v5) & v4
						t61 := int32(int8(m.memory[uint32(t60+v5)]))
						if t61 < i32(0) {
							goto l29
						}
						t62 := int64(load64(m.memory[uint32(v8):]))
						v5 = int32(uint32(int64(bits.TrailingZeros64(uint64(t62&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
					}
				l29:
					{
						if uint32((v5-v6^(v3-v6))&v4) < uint32(i32(8)) {
							goto l30
						}
						v6 = v8 + v5
						t63 := int32(m.memory[uint32(v6)])
						v12 = t63
						t64 := v6
						v11 = int32(uint32(v11) >> 25)
						m.memory[uint32(t64)] = byte(v11)
						m.memory[uint32(v8+(v5+i32(-8))&v4+i32(8))] = byte(v11)
						v6 = v8 + (v5^i32(-1))*i32(36)
						{
							if v12 != i32(255) {
								t70 := int32(load32(m.memory[uint32(v1):]))
								v5 = t70
								t71 := int32(load32(m.memory[uint32(v6):]))
								store32(m.memory[uint32(v1):], uint32(t71))
								store32(m.memory[uint32(v6):], uint32(v5))
								t72 := int32(load32(m.memory[int64(uint32(v6))+4:]))
								v5 = t72
								t73 := int32(load32(m.memory[int64(uint32(v1))+4:]))
								store32(m.memory[int64(uint32(v6))+4:], uint32(t73))
								store32(m.memory[int64(uint32(v1))+4:], uint32(v5))
								t74 := int32(load32(m.memory[int64(uint32(v1))+8:]))
								v5 = t74
								t75 := int32(load32(m.memory[int64(uint32(v6))+8:]))
								store32(m.memory[int64(uint32(v1))+8:], uint32(t75))
								store32(m.memory[int64(uint32(v6))+8:], uint32(v5))
								t76 := int32(load32(m.memory[int64(uint32(v6))+12:]))
								v5 = t76
								t77 := int32(load32(m.memory[int64(uint32(v1))+12:]))
								store32(m.memory[int64(uint32(v6))+12:], uint32(t77))
								store32(m.memory[int64(uint32(v1))+12:], uint32(v5))
								t78 := int32(load32(m.memory[int64(uint32(v1))+16:]))
								v5 = t78
								t79 := int32(load32(m.memory[int64(uint32(v6))+16:]))
								store32(m.memory[int64(uint32(v1))+16:], uint32(t79))
								store32(m.memory[int64(uint32(v6))+16:], uint32(v5))
								t80 := int32(load32(m.memory[int64(uint32(v6))+20:]))
								v5 = t80
								t81 := int32(load32(m.memory[int64(uint32(v1))+20:]))
								store32(m.memory[int64(uint32(v6))+20:], uint32(t81))
								store32(m.memory[int64(uint32(v1))+20:], uint32(v5))
								t82 := int32(load32(m.memory[int64(uint32(v1))+24:]))
								v5 = t82
								t83 := int32(load32(m.memory[int64(uint32(v6))+24:]))
								store32(m.memory[int64(uint32(v1))+24:], uint32(t83))
								store32(m.memory[int64(uint32(v6))+24:], uint32(v5))
								t84 := int32(load32(m.memory[int64(uint32(v6))+28:]))
								v5 = t84
								t85 := int32(load32(m.memory[int64(uint32(v1))+28:]))
								store32(m.memory[int64(uint32(v6))+28:], uint32(t85))
								store32(m.memory[int64(uint32(v1))+28:], uint32(v5))
								t86 := int32(load32(m.memory[int64(uint32(v1))+32:]))
								v5 = t86
								t87 := int32(load32(m.memory[int64(uint32(v6))+32:]))
								store32(m.memory[int64(uint32(v1))+32:], uint32(t87))
								store32(m.memory[int64(uint32(v6))+32:], uint32(v5))
								goto l32
							}
							m.memory[uint32(v10)] = byte(i32(255))
							m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(i32(255))
							t65 := int32(load32(m.memory[int64(uint32(v1))+32:]))
							store32(m.memory[int64(uint32(v6))+32:], uint32(t65))
							t66 := int64(load64(m.memory[int64(uint32(v1))+24:]))
							store64(m.memory[int64(uint32(v6))+24:], uint64(t66))
							t67 := int64(load64(m.memory[int64(uint32(v1))+16:]))
							store64(m.memory[int64(uint32(v6))+16:], uint64(t67))
							t68 := int64(load64(m.memory[int64(uint32(v1))+8:]))
							store64(m.memory[int64(uint32(v6))+8:], uint64(t68))
							t69 := int64(load64(m.memory[uint32(v1):]))
							store64(m.memory[uint32(v6):], uint64(t69))
							goto l26
						}
					}
				l30:
				}
				t88 := v10
				v6 = int32(uint32(v11) >> 25)
				m.memory[uint32(t88)] = byte(v6)
				m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(v6)
			}
		l26:
			v6 = v3 + i32(1)
			if v3 != v4 {
				goto l33
			}
			p89 := v7
			if uint32(v4) < uint32(i32(8)) {
				p89 = v4
			}
			v3 = p89
		}
	l20:
		store32(m.memory[int64(uint32(v0))+8:], uint32(v3-v2))
	l15:
		return i32(-1)
	}
l5:
	m.fn28(i32(1271248), i32(57), i32(1271276))
	panic("unreachable")
}
func (m *Module) fn110(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7, v8 int32
	var v9 int64
	var v10, v11 int32
	var v12, v13 int64
	var v14, v15 int32
	var v16 int64
	var v17 int32
	var v18 int64
	var v19 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
		v2 = t0
		v3 = v2 + i32(1)
		if v3 == 0 {
			m.fn28(i32(1271248), i32(57), i32(1271276))
			panic("unreachable")
		}
		{
			t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t2 := v3
			v4 = t1
			t3 := v4
			v5 = v4 + i32(1)
			v6 = int32(uint32(v5) >> 3)
			v7 = v6 * i32(7)
			p4 := v7
			if uint32(v4) < uint32(i32(8)) {
				p4 = t3
			}
			v8 = p4
			if uint32(t2) <= uint32(int32(uint32(v8)>>1)) {
				goto l1
			}
			{
				{
					{
						v8 = v8 + i32(1)
						p5 := v3
						if uint32(v8) > uint32(v3) {
							p5 = v8
						}
						v3 = p5
						if uint32(v3) < uint32(i32(15)) {
							goto l2
						}
						if uint32(v3) > uint32(i32(0x1fffffff)) {
							m.fn28(i32(1271248), i32(57), i32(1271276))
							panic("unreachable")
						}
						t6 := int32(uint32(v3<<3) / uint32(i32(7)))
						v3 = i32_shr_u(i32(-1), int32(bits.LeadingZeros32(uint32(t6+i32(-1))))) + i32(1)
						goto l4
					}
				l2:
					p7 := v3&i32(8) + i32(8)
					if uint32(v3) < uint32(i32(4)) {
						p7 = i32(4)
					}
					v3 = p7
				}
			l4:
				v9 = int64(uint32(v3)) * i64(520)
				if int32(int64(uint64(v9)>>32)) != 0 {
					goto l5
				}
				v8 = v3 + i32(8)
				t8 := v8
				v10 = int32(v9)
				v6 = t8 + v10
				if uint32(v6) < uint32(v8) {
					goto l5
				}
				if uint32(v6) > uint32(i32(0x7ffffff8)) {
					goto l5
				}
				t9 := m.fn5(v6)
				v5 = t9
				if v5 != 0 {
					v6 = v5 + v10
					if v8 == 0 {
						goto l7
					}
					memory_fill(m.memory, uint32(v6), i32(255), uint32(v8))
				l7:
					v5 = v3 + i32(-1)
					p10 := int32(uint32(v3)>>3) * i32(7)
					if uint32(v3) < uint32(i32(9)) {
						p10 = v5
					}
					v7 = p10
					t11 := int32(load32(m.memory[uint32(v0):]))
					v11 = t11
					{
						if v2 == 0 {
							goto l8
						}
						t12 := int64(load64(m.memory[uint32(v11):]))
						v9 = (t12 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
						t13 := int64(load64(m.memory[int64(uint32(v1))+8:]))
						v12 = t13
						t14 := int64(load64(m.memory[uint32(v1):]))
						v13 = t14
						v8 = v11
						v1 = v2
						v3 = i32(0)
					l14:
						{
							if v9 != i64(0) {
								goto l9
							}
						l10:
							{
								v3 = v3 + i32(8)
								v8 = v8 + i32(8)
								t15 := int64(load64(m.memory[uint32(v8):]))
								v9 = t15 & i64(-0x7f7f7f7f7f7f7f80)
								if v9 == i64(-0x7f7f7f7f7f7f7f80) {
									goto l10
								}
							}
							v9 = v9 ^ i64(-0x7f7f7f7f7f7f7f80)
						l9:
							{
								t16 := v6
								t17 := v5
								t18 := v13
								t19 := v12
								t20 := v11
								v14 = int32(uint32(int64(bits.TrailingZeros64(uint64(v9))))>>3) + v3
								t21 := int32(load16(m.memory[uint32(t20+(i32(0)-v14)*i32(520)+i32(-520)):]))
								t22 := m.fn106(t18, t19, t21)
								v15 = int32(t22)
								v10 = t17 & v15
								t23 := int64(load64(m.memory[uint32(t16+v10):]))
								v16 = t23 & i64(-0x7f7f7f7f7f7f7f80)
								if v16 != i64(0) {
									goto l11
								}
								v17 = i32(8)
							l12:
								{
									v10 = v10 + v17
									v17 = v17 + i32(8)
									t24 := v6
									v10 = v10 & v5
									t25 := int64(load64(m.memory[uint32(t24+v10):]))
									v16 = t25 & i64(-0x7f7f7f7f7f7f7f80)
									if v16 == 0 {
										goto l12
									}
								}
							}
						l11:
							v18 = v9 + i64(-1)
							{
								t26 := v6
								v10 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v16))))>>3) + v10) & v5
								t27 := int32(int8(m.memory[uint32(t26+v10)]))
								if t27 < i32(0) {
									goto l13
								}
								t28 := int64(load64(m.memory[uint32(v6):]))
								v10 = int32(uint32(int64(bits.TrailingZeros64(uint64(t28&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
							}
						l13:
							v9 = v18 & v9
							t29 := v6 + v10
							v15 = int32(uint32(v15) >> 25)
							m.memory[uint32(t29)] = byte(v15)
							m.memory[uint32(v6+(v10+i32(-8))&v5+i32(8))] = byte(v15)
							memory_copy(m.memory, uint32(v6+(v10^i32(-1))*i32(520)), uint32(v11+(v14^i32(-1))*i32(520)), uint32(i32(520)))
							v1 = v1 + i32(-1)
							if v1 != 0 {
								goto l14
							}
						}
					}
				l8:
					store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
					store32(m.memory[uint32(v0):], uint32(v6))
					store32(m.memory[int64(uint32(v0))+8:], uint32(v7-v2))
					if v4 == 0 {
						goto l15
					}
					t30 := v4
					v8 = (v4*i32(520) + i32(527)) & i32(-8)
					v3 = t30 + v8 + i32(9)
					if v3 == 0 {
						goto l15
					}
					v4 = v11 - v8
					t31 := int32(load32(m.memory[uint32(v4+i32(-4)):]))
					v8 = t31
					v6 = v8 & i32(-8)
					t32 := v6
					v8 = v8 & i32(3)
					p33 := i32(8)
					if v8 != 0 {
						p33 = i32(4)
					}
					if uint32(t32) < uint32(p33+v3) {
						m.fn3(i32(1273840), i32(46), i32(1273888))
						panic("unreachable")
					}
					if v8 == 0 {
						goto l17
					}
					if uint32(v6) > uint32(v3+i32(39)) {
						m.fn3(i32(1273904), i32(46), i32(1273952))
						panic("unreachable")
					}
				l17:
					m.fn1(v4)
					return i32(-1)
				}
				m.fn24(i32(8), v6)
				panic("unreachable")
			}
		}
	l1:
		{
			if v5 != 0 {
				goto l19
			}
			v3 = i32(0)
			goto l20
		l19:
			t34 := int32(load32(m.memory[uint32(v0):]))
			v8 = t34
			v3 = i32(0)
			{
				{
					t35 := v6
					var p36 int32
					if v5&i32(7) != i32(0) {
						p36 = 1
					}
					v6 = t35 + p36
					if v6 == i32(1) {
						goto l21
					}
					v11 = v6 & i32(1)
					v10 = v6 & i32(0x3ffffffe)
					v3 = i32(0)
				l22:
					{
						v6 = v8 + v3
						t37 := int64(load64(m.memory[uint32(v6):]))
						t38 := v6
						v9 = t37
						store64(m.memory[uint32(t38):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
						v6 = v6 + i32(8)
						t39 := int64(load64(m.memory[uint32(v6):]))
						t40 := v6
						v9 = t39
						store64(m.memory[uint32(t40):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
						v3 = v3 + i32(16)
						v10 = v10 + i32(-2)
						if v10 != 0 {
							goto l22
						}
					}
					if v11 == 0 {
						goto l23
					}
				}
			l21:
				v3 = v8 + v3
				t41 := int64(load64(m.memory[uint32(v3):]))
				t42 := v3
				v9 = t41
				store64(m.memory[uint32(t42):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
			}
		l23:
			{
				if uint32(v5) < uint32(i32(8)) {
					goto l24
				}
				t43 := int64(load64(m.memory[uint32(v8):]))
				store64(m.memory[uint32(v8+v5):], uint64(t43))
				goto l25
			}
		l24:
			if v5 == 0 {
				goto l25
			}
			memory_copy(m.memory, uint32(v8+i32(8)), uint32(v8), uint32(v5))
		l25:
			t44 := int64(load64(m.memory[int64(uint32(v1))+8:]))
			v16 = t44
			t45 := int64(load64(m.memory[uint32(v1):]))
			v18 = t45
			v6 = v8
			v10 = i32(0)
		l34:
			{
				t46 := v8
				v3 = v10
				v5 = t46 + v3
				t47 := int32(m.memory[uint32(v5)])
				if t47 != i32(128) {
					goto l26
				}
				v15 = v8 + (v3^i32(-1))*i32(520)
				v14 = v8 + (i32(0)-v3)*i32(520) + i32(-520)
				{
				l33:
					{
						t48 := int32(load16(m.memory[uint32(v14):]))
						t49 := m.fn106(v18, v16, t48)
						t50 := v4
						v1 = int32(t49)
						v10 = t50 & v1
						v11 = v10
						{
							t51 := int64(load64(m.memory[uint32(v8+v10):]))
							v9 = t51 & i64(-0x7f7f7f7f7f7f7f80)
							if v9 != i64(0) {
								goto l27
							}
							v17 = i32(8)
							v11 = v10
						l28:
							{
								v11 = v11 + v17
								v17 = v17 + i32(8)
								t52 := v8
								v11 = v11 & v4
								t53 := int64(load64(m.memory[uint32(t52+v11):]))
								v9 = t53 & i64(-0x7f7f7f7f7f7f7f80)
								if v9 == 0 {
									goto l28
								}
							}
						}
					l27:
						{
							t54 := v8
							v11 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v9))))>>3) + v11) & v4
							t55 := int32(int8(m.memory[uint32(t54+v11)]))
							if t55 < i32(0) {
								goto l29
							}
							t56 := int64(load64(m.memory[uint32(v8):]))
							v11 = int32(uint32(int64(bits.TrailingZeros64(uint64(t56&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
						}
					l29:
						{
							if uint32((v11-v10^(v3-v10))&v4) < uint32(i32(8)) {
								goto l30
							}
							v10 = v8 + v11
							t57 := int32(m.memory[uint32(v10)])
							v17 = t57
							t58 := v10
							v1 = int32(uint32(v1) >> 25)
							m.memory[uint32(t58)] = byte(v1)
							m.memory[uint32(v8+(v11+i32(-8))&v4+i32(8))] = byte(v1)
							if v17 == i32(255) {
								goto l31
							}
							v10 = i32(-520)
							v19 = v8 + v11*i32(-520)
						l32:
							{
								v11 = v19 + v10
								t59 := int32(load32(m.memory[uint32(v11):]))
								v17 = t59
								t60 := v11
								v1 = v6 + v10
								t61 := int32(load32(m.memory[uint32(v1):]))
								store32(m.memory[uint32(t60):], uint32(t61))
								store32(m.memory[uint32(v1):], uint32(v17))
								v1 = v1 + i32(4)
								t62 := int32(load32(m.memory[uint32(v1):]))
								v17 = t62
								t63 := v1
								v11 = v11 + i32(4)
								t64 := int32(load32(m.memory[uint32(v11):]))
								store32(m.memory[uint32(t63):], uint32(t64))
								store32(m.memory[uint32(v11):], uint32(v17))
								v10 = v10 + i32(8)
								if v10 != 0 {
									goto l32
								}
								goto l33
							}
						}
					l30:
					}
					t65 := v5
					v10 = int32(uint32(v1) >> 25)
					m.memory[uint32(t65)] = byte(v10)
					m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(v10)
					goto l26
				}
			l31:
				m.memory[uint32(v5)] = byte(i32(255))
				m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(i32(255))
				memory_copy(m.memory, uint32(v8+(v11^i32(-1))*i32(520)), uint32(v15), uint32(i32(520)))
			}
		l26:
			v10 = v3 + i32(1)
			v6 = v6 + i32(-520)
			if v3 != v4 {
				goto l34
			}
			p66 := v7
			if uint32(v4) < uint32(i32(8)) {
				p66 = v4
			}
			v3 = p66
		}
	l20:
		store32(m.memory[int64(uint32(v0))+8:], uint32(v3-v2))
	l15:
		return i32(-1)
	l5:
		m.fn28(i32(1271248), i32(57), i32(1271276))
		panic("unreachable")
	}
}
func (m *Module) fn111(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7, v8, v9, v10, v11 int32
	var v12, v13, v14 int64
	var v15 int32
	var v16 int64
	var v17 int32
	var v18 int64
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
		v2 = t0
		v3 = v2 + i32(1)
		if v3 == 0 {
			m.fn28(i32(1271248), i32(57), i32(1271276))
			panic("unreachable")
		}
		{
			t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t2 := v3
			v4 = t1
			t3 := v4
			v5 = v4 + i32(1)
			v6 = int32(uint32(v5) >> 3)
			v7 = v6 * i32(7)
			p4 := v7
			if uint32(v4) < uint32(i32(8)) {
				p4 = t3
			}
			v8 = p4
			if uint32(t2) <= uint32(int32(uint32(v8)>>1)) {
				goto l1
			}
			{
				{
					{
						v8 = v8 + i32(1)
						p5 := v3
						if uint32(v8) > uint32(v3) {
							p5 = v8
						}
						v3 = p5
						if uint32(v3) < uint32(i32(15)) {
							goto l2
						}
						if uint32(v3) > uint32(i32(0x1fffffff)) {
							m.fn28(i32(1271248), i32(57), i32(1271276))
							panic("unreachable")
						}
						t6 := int32(uint32(v3<<3) / uint32(i32(7)))
						v3 = i32_shr_u(i32(-1), int32(bits.LeadingZeros32(uint32(t6+i32(-1))))) + i32(1)
						goto l4
					}
				l2:
					p7 := v3&i32(8) + i32(8)
					if uint32(v3) < uint32(i32(4)) {
						p7 = i32(4)
					}
					v3 = p7
				}
			l4:
				v8 = v3 + i32(8)
				t8 := v8
				v9 = (v3<<1 + i32(7)) & i32(-8)
				v6 = t8 + v9
				if uint32(v6) > uint32(i32(0x7ffffff8)) {
					m.fn28(i32(1271248), i32(57), i32(1271276))
					panic("unreachable")
				}
				t9 := m.fn5(v6)
				v5 = t9
				if v5 != 0 {
					v6 = v5 + v9
					if v8 == 0 {
						goto l7
					}
					memory_fill(m.memory, uint32(v6), i32(255), uint32(v8))
				l7:
					v5 = v3 + i32(-1)
					p10 := int32(uint32(v3)>>3) * i32(7)
					if uint32(v3) < uint32(i32(9)) {
						p10 = v5
					}
					v10 = p10
					t11 := int32(load32(m.memory[uint32(v0):]))
					v11 = t11
					{
						if v2 == 0 {
							goto l8
						}
						t12 := int64(load64(m.memory[uint32(v11):]))
						v12 = (t12 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
						t13 := int64(load64(m.memory[int64(uint32(v1))+8:]))
						v13 = t13
						t14 := int64(load64(m.memory[uint32(v1):]))
						v14 = t14
						v8 = v11
						v1 = v2
						v3 = i32(0)
					l14:
						{
							if v12 != i64(0) {
								goto l9
							}
						l10:
							{
								v3 = v3 + i32(8)
								v8 = v8 + i32(8)
								t15 := int64(load64(m.memory[uint32(v8):]))
								v12 = t15 & i64(-0x7f7f7f7f7f7f7f80)
								if v12 == i64(-0x7f7f7f7f7f7f7f80) {
									goto l10
								}
							}
							v12 = v12 ^ i64(-0x7f7f7f7f7f7f7f80)
						l9:
							{
								t16 := v6
								t17 := v5
								t18 := v14
								t19 := v13
								t20 := v11
								v15 = int32(uint32(int64(bits.TrailingZeros64(uint64(v12))))>>3) + v3
								t21 := int32(load16(m.memory[uint32(t20-v15<<1+i32(-2)):]))
								t22 := m.fn106(t18, t19, t21)
								v7 = int32(t22)
								v9 = t17 & v7
								t23 := int64(load64(m.memory[uint32(t16+v9):]))
								v16 = t23 & i64(-0x7f7f7f7f7f7f7f80)
								if v16 != i64(0) {
									goto l11
								}
								v17 = i32(8)
							l12:
								{
									v9 = v9 + v17
									v17 = v17 + i32(8)
									t24 := v6
									v9 = v9 & v5
									t25 := int64(load64(m.memory[uint32(t24+v9):]))
									v16 = t25 & i64(-0x7f7f7f7f7f7f7f80)
									if v16 == 0 {
										goto l12
									}
								}
							}
						l11:
							v18 = v12 + i64(-1)
							{
								t26 := v6
								v9 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v16))))>>3) + v9) & v5
								t27 := int32(int8(m.memory[uint32(t26+v9)]))
								if t27 < i32(0) {
									goto l13
								}
								t28 := int64(load64(m.memory[uint32(v6):]))
								v9 = int32(uint32(int64(bits.TrailingZeros64(uint64(t28&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
							}
						l13:
							v12 = v18 & v12
							t29 := v6 + v9
							v7 = int32(uint32(v7) >> 25)
							m.memory[uint32(t29)] = byte(v7)
							m.memory[uint32(v6+(v9+i32(-8))&v5+i32(8))] = byte(v7)
							t30 := int32(load16(m.memory[uint32(v11+(v15^i32(-1))<<1):]))
							store16(m.memory[uint32(v6+(v9^i32(-1))<<1):], uint16(t30))
							v1 = v1 + i32(-1)
							if v1 != 0 {
								goto l14
							}
						}
					}
				l8:
					store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
					store32(m.memory[uint32(v0):], uint32(v6))
					store32(m.memory[int64(uint32(v0))+8:], uint32(v10-v2))
					if v4 == 0 {
						goto l15
					}
					t31 := v4
					v8 = (v4<<1 + i32(9)) & i32(-8)
					v3 = t31 + v8 + i32(9)
					if v3 == 0 {
						goto l15
					}
					v4 = v11 - v8
					t32 := int32(load32(m.memory[uint32(v4+i32(-4)):]))
					v8 = t32
					v6 = v8 & i32(-8)
					t33 := v6
					v8 = v8 & i32(3)
					p34 := i32(8)
					if v8 != 0 {
						p34 = i32(4)
					}
					if uint32(t33) < uint32(p34+v3) {
						m.fn3(i32(1273840), i32(46), i32(1273888))
						panic("unreachable")
					}
					if v8 == 0 {
						goto l17
					}
					if uint32(v6) > uint32(v3+i32(39)) {
						m.fn3(i32(1273904), i32(46), i32(1273952))
						panic("unreachable")
					}
				l17:
					m.fn1(v4)
					return i32(-1)
				}
				m.fn24(i32(8), v6)
				panic("unreachable")
			}
		}
	l1:
		{
			if v5 != 0 {
				goto l19
			}
			v3 = i32(0)
			goto l20
		l19:
			t35 := int32(load32(m.memory[uint32(v0):]))
			v8 = t35
			v3 = i32(0)
			{
				{
					t36 := v6
					var p37 int32
					if v5&i32(7) != i32(0) {
						p37 = 1
					}
					v6 = t36 + p37
					if v6 == i32(1) {
						goto l21
					}
					v11 = v6 & i32(1)
					v9 = v6 & i32(0x3ffffffe)
					v3 = i32(0)
				l22:
					{
						v6 = v8 + v3
						t38 := int64(load64(m.memory[uint32(v6):]))
						t39 := v6
						v12 = t38
						store64(m.memory[uint32(t39):], uint64(int64(uint64(v12^i64(-1))>>7)&i64(72340172838076673)+(v12|i64(0x7f7f7f7f7f7f7f7f))))
						v6 = v6 + i32(8)
						t40 := int64(load64(m.memory[uint32(v6):]))
						t41 := v6
						v12 = t40
						store64(m.memory[uint32(t41):], uint64(int64(uint64(v12^i64(-1))>>7)&i64(72340172838076673)+(v12|i64(0x7f7f7f7f7f7f7f7f))))
						v3 = v3 + i32(16)
						v9 = v9 + i32(-2)
						if v9 != 0 {
							goto l22
						}
					}
					if v11 == 0 {
						goto l23
					}
				}
			l21:
				v3 = v8 + v3
				t42 := int64(load64(m.memory[uint32(v3):]))
				t43 := v3
				v12 = t42
				store64(m.memory[uint32(t43):], uint64(int64(uint64(v12^i64(-1))>>7)&i64(72340172838076673)+(v12|i64(0x7f7f7f7f7f7f7f7f))))
			}
		l23:
			{
				if uint32(v5) < uint32(i32(8)) {
					goto l24
				}
				t44 := int64(load64(m.memory[uint32(v8):]))
				store64(m.memory[uint32(v8+v5):], uint64(t44))
				goto l25
			}
		l24:
			if v5 == 0 {
				goto l25
			}
			memory_copy(m.memory, uint32(v8+i32(8)), uint32(v8), uint32(v5))
		l25:
			t45 := int64(load64(m.memory[int64(uint32(v1))+8:]))
			v16 = t45
			t46 := int64(load64(m.memory[uint32(v1):]))
			v18 = t46
			v6 = i32(0)
		l33:
			{
				t47 := v8
				v3 = v6
				v9 = t47 + v3
				t48 := int32(m.memory[uint32(v9)])
				if t48 != i32(128) {
					goto l26
				}
				v1 = v8 - v3<<1 + i32(-2)
				v15 = v8 + (v3^i32(-1))<<1
			l32:
				{
					t49 := int32(load16(m.memory[uint32(v1):]))
					t50 := m.fn106(v18, v16, t49)
					t51 := v4
					v11 = int32(t50)
					v6 = t51 & v11
					v5 = v6
					{
						t52 := int64(load64(m.memory[uint32(v8+v6):]))
						v12 = t52 & i64(-0x7f7f7f7f7f7f7f80)
						if v12 != i64(0) {
							goto l27
						}
						v10 = i32(8)
						v5 = v6
					l28:
						{
							v5 = v5 + v10
							v10 = v10 + i32(8)
							t53 := v8
							v5 = v5 & v4
							t54 := int64(load64(m.memory[uint32(t53+v5):]))
							v12 = t54 & i64(-0x7f7f7f7f7f7f7f80)
							if v12 == 0 {
								goto l28
							}
						}
					}
				l27:
					{
						t55 := v8
						v5 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v12))))>>3) + v5) & v4
						t56 := int32(int8(m.memory[uint32(t55+v5)]))
						if t56 < i32(0) {
							goto l29
						}
						t57 := int64(load64(m.memory[uint32(v8):]))
						v5 = int32(uint32(int64(bits.TrailingZeros64(uint64(t57&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
					}
				l29:
					{
						if uint32((v5-v6^(v3-v6))&v4) < uint32(i32(8)) {
							goto l30
						}
						v6 = v8 + v5
						t58 := int32(m.memory[uint32(v6)])
						v10 = t58
						t59 := v6
						v11 = int32(uint32(v11) >> 25)
						m.memory[uint32(t59)] = byte(v11)
						m.memory[uint32(v8+(v5+i32(-8))&v4+i32(8))] = byte(v11)
						v6 = v8 + (v5^i32(-1))<<1
						{
							if v10 != i32(255) {
								t61 := int32(load16(m.memory[uint32(v15):]))
								v5 = t61
								t62 := int32(load16(m.memory[uint32(v6):]))
								store16(m.memory[uint32(v15):], uint16(t62))
								store16(m.memory[uint32(v6):], uint16(v5))
								goto l32
							}
							m.memory[uint32(v9)] = byte(i32(255))
							m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(i32(255))
							t60 := int32(load16(m.memory[uint32(v15):]))
							store16(m.memory[uint32(v6):], uint16(t60))
							goto l26
						}
					}
				l30:
				}
				t63 := v9
				v6 = int32(uint32(v11) >> 25)
				m.memory[uint32(t63)] = byte(v6)
				m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(v6)
			}
		l26:
			v6 = v3 + i32(1)
			if v3 != v4 {
				goto l33
			}
			p64 := v7
			if uint32(v4) < uint32(i32(8)) {
				p64 = v4
			}
			v3 = p64
		}
	l20:
		store32(m.memory[int64(uint32(v0))+8:], uint32(v3-v2))
	l15:
		return i32(-1)
	}
}
func (m *Module) fn112(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7, v8 int32
	var v9 int64
	var v10, v11, v12 int32
	var v13, v14 int64
	var v15 int32
	var v16 int64
	var v17 int32
	var v18 int64
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
		v2 = t0
		v3 = v2 + i32(1)
		if v3 == 0 {
			m.fn28(i32(1271248), i32(57), i32(1271276))
			panic("unreachable")
		}
		{
			t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t2 := v3
			v4 = t1
			t3 := v4
			v5 = v4 + i32(1)
			v6 = int32(uint32(v5) >> 3)
			v7 = v6 * i32(7)
			p4 := v7
			if uint32(v4) < uint32(i32(8)) {
				p4 = t3
			}
			v8 = p4
			if uint32(t2) <= uint32(int32(uint32(v8)>>1)) {
				goto l1
			}
			{
				{
					{
						v8 = v8 + i32(1)
						p5 := v3
						if uint32(v8) > uint32(v3) {
							p5 = v8
						}
						v3 = p5
						if uint32(v3) < uint32(i32(15)) {
							goto l2
						}
						if uint32(v3) > uint32(i32(0x1fffffff)) {
							m.fn28(i32(1271248), i32(57), i32(1271276))
							panic("unreachable")
						}
						t6 := int32(uint32(v3<<3) / uint32(i32(7)))
						v3 = i32_shr_u(i32(-1), int32(bits.LeadingZeros32(uint32(t6+i32(-1))))) + i32(1)
						goto l4
					}
				l2:
					p7 := v3&i32(8) + i32(8)
					if uint32(v3) < uint32(i32(4)) {
						p7 = i32(4)
					}
					v3 = p7
				}
			l4:
				v9 = int64(uint32(v3)) * i64(104)
				if int32(int64(uint64(v9)>>32)) != 0 {
					goto l5
				}
				v8 = v3 + i32(8)
				t8 := v8
				v10 = int32(v9)
				v6 = t8 + v10
				if uint32(v6) < uint32(v8) {
					goto l5
				}
				if uint32(v6) > uint32(i32(0x7ffffff8)) {
					goto l5
				}
				t9 := m.fn5(v6)
				v5 = t9
				if v5 != 0 {
					v6 = v5 + v10
					if v8 == 0 {
						goto l7
					}
					memory_fill(m.memory, uint32(v6), i32(255), uint32(v8))
				l7:
					v5 = v3 + i32(-1)
					p10 := int32(uint32(v3)>>3) * i32(7)
					if uint32(v3) < uint32(i32(9)) {
						p10 = v5
					}
					v11 = p10
					t11 := int32(load32(m.memory[uint32(v0):]))
					v12 = t11
					{
						if v2 == 0 {
							goto l8
						}
						t12 := int64(load64(m.memory[uint32(v12):]))
						v9 = (t12 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
						t13 := int64(load64(m.memory[int64(uint32(v1))+8:]))
						v13 = t13
						t14 := int64(load64(m.memory[uint32(v1):]))
						v14 = t14
						v8 = v12
						v1 = v2
						v3 = i32(0)
					l14:
						{
							if v9 != i64(0) {
								goto l9
							}
						l10:
							{
								v3 = v3 + i32(8)
								v8 = v8 + i32(8)
								t15 := int64(load64(m.memory[uint32(v8):]))
								v9 = t15 & i64(-0x7f7f7f7f7f7f7f80)
								if v9 == i64(-0x7f7f7f7f7f7f7f80) {
									goto l10
								}
							}
							v9 = v9 ^ i64(-0x7f7f7f7f7f7f7f80)
						l9:
							{
								t16 := v6
								t17 := v5
								t18 := v14
								t19 := v13
								t20 := v12
								v15 = int32(uint32(int64(bits.TrailingZeros64(uint64(v9))))>>3) + v3
								t21 := int64(load64(m.memory[uint32(t20+(i32(0)-v15)*i32(104)+i32(-104)):]))
								t22 := m.fn113(t18, t19, t21)
								v7 = int32(t22)
								v10 = t17 & v7
								t23 := int64(load64(m.memory[uint32(t16+v10):]))
								v16 = t23 & i64(-0x7f7f7f7f7f7f7f80)
								if v16 != i64(0) {
									goto l11
								}
								v17 = i32(8)
							l12:
								{
									v10 = v10 + v17
									v17 = v17 + i32(8)
									t24 := v6
									v10 = v10 & v5
									t25 := int64(load64(m.memory[uint32(t24+v10):]))
									v16 = t25 & i64(-0x7f7f7f7f7f7f7f80)
									if v16 == 0 {
										goto l12
									}
								}
							}
						l11:
							v18 = v9 + i64(-1)
							{
								t26 := v6
								v10 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v16))))>>3) + v10) & v5
								t27 := int32(int8(m.memory[uint32(t26+v10)]))
								if t27 < i32(0) {
									goto l13
								}
								t28 := int64(load64(m.memory[uint32(v6):]))
								v10 = int32(uint32(int64(bits.TrailingZeros64(uint64(t28&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
							}
						l13:
							v9 = v18 & v9
							t29 := v6 + v10
							v7 = int32(uint32(v7) >> 25)
							m.memory[uint32(t29)] = byte(v7)
							m.memory[uint32(v6+(v10+i32(-8))&v5+i32(8))] = byte(v7)
							memory_copy(m.memory, uint32(v6+(v10^i32(-1))*i32(104)), uint32(v12+(v15^i32(-1))*i32(104)), uint32(i32(104)))
							v1 = v1 + i32(-1)
							if v1 != 0 {
								goto l14
							}
						}
					}
				l8:
					store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
					store32(m.memory[uint32(v0):], uint32(v6))
					store32(m.memory[int64(uint32(v0))+8:], uint32(v11-v2))
					if v4 == 0 {
						goto l15
					}
					t30 := v4
					v8 = (v4*i32(104) + i32(111)) & i32(-8)
					v3 = t30 + v8 + i32(9)
					if v3 == 0 {
						goto l15
					}
					v4 = v12 - v8
					t31 := int32(load32(m.memory[uint32(v4+i32(-4)):]))
					v8 = t31
					v6 = v8 & i32(-8)
					t32 := v6
					v8 = v8 & i32(3)
					p33 := i32(8)
					if v8 != 0 {
						p33 = i32(4)
					}
					if uint32(t32) < uint32(p33+v3) {
						m.fn3(i32(1273840), i32(46), i32(1273888))
						panic("unreachable")
					}
					if v8 == 0 {
						goto l17
					}
					if uint32(v6) > uint32(v3+i32(39)) {
						m.fn3(i32(1273904), i32(46), i32(1273952))
						panic("unreachable")
					}
				l17:
					m.fn1(v4)
					return i32(-1)
				}
				m.fn24(i32(8), v6)
				panic("unreachable")
			}
		}
	l1:
		{
			if v5 != 0 {
				goto l19
			}
			v3 = i32(0)
			goto l20
		l19:
			t34 := int32(load32(m.memory[uint32(v0):]))
			v8 = t34
			v3 = i32(0)
			{
				{
					t35 := v6
					var p36 int32
					if v5&i32(7) != i32(0) {
						p36 = 1
					}
					v6 = t35 + p36
					if v6 == i32(1) {
						goto l21
					}
					v12 = v6 & i32(1)
					v10 = v6 & i32(0x3ffffffe)
					v3 = i32(0)
				l22:
					{
						v6 = v8 + v3
						t37 := int64(load64(m.memory[uint32(v6):]))
						t38 := v6
						v9 = t37
						store64(m.memory[uint32(t38):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
						v6 = v6 + i32(8)
						t39 := int64(load64(m.memory[uint32(v6):]))
						t40 := v6
						v9 = t39
						store64(m.memory[uint32(t40):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
						v3 = v3 + i32(16)
						v10 = v10 + i32(-2)
						if v10 != 0 {
							goto l22
						}
					}
					if v12 == 0 {
						goto l23
					}
				}
			l21:
				v3 = v8 + v3
				t41 := int64(load64(m.memory[uint32(v3):]))
				t42 := v3
				v9 = t41
				store64(m.memory[uint32(t42):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
			}
		l23:
			{
				if uint32(v5) < uint32(i32(8)) {
					goto l24
				}
				t43 := int64(load64(m.memory[uint32(v8):]))
				store64(m.memory[uint32(v8+v5):], uint64(t43))
				goto l25
			}
		l24:
			if v5 == 0 {
				goto l25
			}
			memory_copy(m.memory, uint32(v8+i32(8)), uint32(v8), uint32(v5))
		l25:
			t44 := int64(load64(m.memory[int64(uint32(v1))+8:]))
			v16 = t44
			t45 := int64(load64(m.memory[uint32(v1):]))
			v18 = t45
			v6 = i32(0)
		l33:
			{
				t46 := v8
				v3 = v6
				v10 = t46 + v3
				t47 := int32(m.memory[uint32(v10)])
				if t47 != i32(128) {
					goto l26
				}
				v15 = v8 + (v3^i32(-1))*i32(104)
				v1 = v8 + (i32(0)-v3)*i32(104) + i32(-104)
			l32:
				{
					t48 := int64(load64(m.memory[uint32(v1):]))
					t49 := m.fn113(v18, v16, t48)
					t50 := v4
					v12 = int32(t49)
					v6 = t50 & v12
					v5 = v6
					{
						t51 := int64(load64(m.memory[uint32(v8+v6):]))
						v9 = t51 & i64(-0x7f7f7f7f7f7f7f80)
						if v9 != i64(0) {
							goto l27
						}
						v11 = i32(8)
						v5 = v6
					l28:
						{
							v5 = v5 + v11
							v11 = v11 + i32(8)
							t52 := v8
							v5 = v5 & v4
							t53 := int64(load64(m.memory[uint32(t52+v5):]))
							v9 = t53 & i64(-0x7f7f7f7f7f7f7f80)
							if v9 == 0 {
								goto l28
							}
						}
					}
				l27:
					{
						t54 := v8
						v5 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v9))))>>3) + v5) & v4
						t55 := int32(int8(m.memory[uint32(t54+v5)]))
						if t55 < i32(0) {
							goto l29
						}
						t56 := int64(load64(m.memory[uint32(v8):]))
						v5 = int32(uint32(int64(bits.TrailingZeros64(uint64(t56&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
					}
				l29:
					{
						if uint32((v5-v6^(v3-v6))&v4) < uint32(i32(8)) {
							goto l30
						}
						v6 = v8 + v5
						t57 := int32(m.memory[uint32(v6)])
						v11 = t57
						t58 := v6
						v12 = int32(uint32(v12) >> 25)
						m.memory[uint32(t58)] = byte(v12)
						m.memory[uint32(v8+(v5+i32(-8))&v4+i32(8))] = byte(v12)
						v6 = v8 + (v5^i32(-1))*i32(104)
						if v11 != i32(255) {
							t59 := int64(load64(m.memory[uint32(v6):]))
							v9 = t59
							t60 := int64(load64(m.memory[uint32(v15):]))
							store64(m.memory[uint32(v6):], uint64(t60))
							store64(m.memory[uint32(v15):], uint64(v9))
							t61 := int64(load64(m.memory[int64(uint32(v15))+8:]))
							v9 = t61
							t62 := int64(load64(m.memory[int64(uint32(v6))+8:]))
							store64(m.memory[int64(uint32(v15))+8:], uint64(t62))
							store64(m.memory[int64(uint32(v6))+8:], uint64(v9))
							t63 := int32(load32(m.memory[int64(uint32(v6))+16:]))
							v5 = t63
							t64 := int32(load32(m.memory[int64(uint32(v15))+16:]))
							store32(m.memory[int64(uint32(v6))+16:], uint32(t64))
							store32(m.memory[int64(uint32(v15))+16:], uint32(v5))
							t65 := int32(load32(m.memory[int64(uint32(v15))+20:]))
							v5 = t65
							t66 := int32(load32(m.memory[int64(uint32(v6))+20:]))
							store32(m.memory[int64(uint32(v15))+20:], uint32(t66))
							store32(m.memory[int64(uint32(v6))+20:], uint32(v5))
							t67 := int32(load32(m.memory[int64(uint32(v15))+24:]))
							v5 = t67
							t68 := int32(load32(m.memory[int64(uint32(v6))+24:]))
							store32(m.memory[int64(uint32(v15))+24:], uint32(t68))
							store32(m.memory[int64(uint32(v6))+24:], uint32(v5))
							t69 := int32(load32(m.memory[int64(uint32(v6))+28:]))
							v5 = t69
							t70 := int32(load32(m.memory[int64(uint32(v15))+28:]))
							store32(m.memory[int64(uint32(v6))+28:], uint32(t70))
							store32(m.memory[int64(uint32(v15))+28:], uint32(v5))
							t71 := int32(load32(m.memory[int64(uint32(v6))+32:]))
							v5 = t71
							t72 := int32(load32(m.memory[int64(uint32(v15))+32:]))
							store32(m.memory[int64(uint32(v6))+32:], uint32(t72))
							store32(m.memory[int64(uint32(v15))+32:], uint32(v5))
							t73 := int32(load32(m.memory[int64(uint32(v15))+36:]))
							v5 = t73
							t74 := int32(load32(m.memory[int64(uint32(v6))+36:]))
							store32(m.memory[int64(uint32(v15))+36:], uint32(t74))
							store32(m.memory[int64(uint32(v6))+36:], uint32(v5))
							t75 := int32(load32(m.memory[int64(uint32(v15))+40:]))
							v5 = t75
							t76 := int32(load32(m.memory[int64(uint32(v6))+40:]))
							store32(m.memory[int64(uint32(v15))+40:], uint32(t76))
							store32(m.memory[int64(uint32(v6))+40:], uint32(v5))
							t77 := int32(load32(m.memory[int64(uint32(v6))+44:]))
							v5 = t77
							t78 := int32(load32(m.memory[int64(uint32(v15))+44:]))
							store32(m.memory[int64(uint32(v6))+44:], uint32(t78))
							store32(m.memory[int64(uint32(v15))+44:], uint32(v5))
							t79 := int32(load32(m.memory[int64(uint32(v6))+48:]))
							v5 = t79
							t80 := int32(load32(m.memory[int64(uint32(v15))+48:]))
							store32(m.memory[int64(uint32(v6))+48:], uint32(t80))
							store32(m.memory[int64(uint32(v15))+48:], uint32(v5))
							t81 := int32(load32(m.memory[int64(uint32(v15))+52:]))
							v5 = t81
							t82 := int32(load32(m.memory[int64(uint32(v6))+52:]))
							store32(m.memory[int64(uint32(v15))+52:], uint32(t82))
							store32(m.memory[int64(uint32(v6))+52:], uint32(v5))
							t83 := int32(load32(m.memory[int64(uint32(v15))+56:]))
							v5 = t83
							t84 := int32(load32(m.memory[int64(uint32(v6))+56:]))
							store32(m.memory[int64(uint32(v15))+56:], uint32(t84))
							store32(m.memory[int64(uint32(v6))+56:], uint32(v5))
							t85 := int32(load32(m.memory[int64(uint32(v6))+60:]))
							v5 = t85
							t86 := int32(load32(m.memory[int64(uint32(v15))+60:]))
							store32(m.memory[int64(uint32(v6))+60:], uint32(t86))
							store32(m.memory[int64(uint32(v15))+60:], uint32(v5))
							t87 := int32(load32(m.memory[int64(uint32(v6))+64:]))
							v5 = t87
							t88 := int32(load32(m.memory[int64(uint32(v15))+64:]))
							store32(m.memory[int64(uint32(v6))+64:], uint32(t88))
							store32(m.memory[int64(uint32(v15))+64:], uint32(v5))
							t89 := int32(load32(m.memory[int64(uint32(v15))+68:]))
							v5 = t89
							t90 := int32(load32(m.memory[int64(uint32(v6))+68:]))
							store32(m.memory[int64(uint32(v15))+68:], uint32(t90))
							store32(m.memory[int64(uint32(v6))+68:], uint32(v5))
							t91 := int32(load32(m.memory[int64(uint32(v15))+72:]))
							v5 = t91
							t92 := int32(load32(m.memory[int64(uint32(v6))+72:]))
							store32(m.memory[int64(uint32(v15))+72:], uint32(t92))
							store32(m.memory[int64(uint32(v6))+72:], uint32(v5))
							t93 := int32(load32(m.memory[int64(uint32(v6))+76:]))
							v5 = t93
							t94 := int32(load32(m.memory[int64(uint32(v15))+76:]))
							store32(m.memory[int64(uint32(v6))+76:], uint32(t94))
							store32(m.memory[int64(uint32(v15))+76:], uint32(v5))
							t95 := int32(load32(m.memory[int64(uint32(v6))+80:]))
							v5 = t95
							t96 := int32(load32(m.memory[int64(uint32(v15))+80:]))
							store32(m.memory[int64(uint32(v6))+80:], uint32(t96))
							store32(m.memory[int64(uint32(v15))+80:], uint32(v5))
							t97 := int32(load32(m.memory[int64(uint32(v15))+84:]))
							v5 = t97
							t98 := int32(load32(m.memory[int64(uint32(v6))+84:]))
							store32(m.memory[int64(uint32(v15))+84:], uint32(t98))
							store32(m.memory[int64(uint32(v6))+84:], uint32(v5))
							t99 := int32(load32(m.memory[int64(uint32(v15))+88:]))
							v5 = t99
							t100 := int32(load32(m.memory[int64(uint32(v6))+88:]))
							store32(m.memory[int64(uint32(v15))+88:], uint32(t100))
							store32(m.memory[int64(uint32(v6))+88:], uint32(v5))
							t101 := int32(load32(m.memory[int64(uint32(v6))+92:]))
							v5 = t101
							t102 := int32(load32(m.memory[int64(uint32(v15))+92:]))
							store32(m.memory[int64(uint32(v6))+92:], uint32(t102))
							store32(m.memory[int64(uint32(v15))+92:], uint32(v5))
							t103 := int32(load32(m.memory[int64(uint32(v6))+96:]))
							v5 = t103
							t104 := int32(load32(m.memory[int64(uint32(v15))+96:]))
							store32(m.memory[int64(uint32(v6))+96:], uint32(t104))
							store32(m.memory[int64(uint32(v15))+96:], uint32(v5))
							t105 := int32(load32(m.memory[int64(uint32(v15))+100:]))
							v5 = t105
							t106 := int32(load32(m.memory[int64(uint32(v6))+100:]))
							store32(m.memory[int64(uint32(v15))+100:], uint32(t106))
							store32(m.memory[int64(uint32(v6))+100:], uint32(v5))
							goto l32
						}
						m.memory[uint32(v10)] = byte(i32(255))
						m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(i32(255))
						memory_copy(m.memory, uint32(v6), uint32(v15), uint32(i32(104)))
						goto l26
					}
				l30:
				}
				t107 := v10
				v6 = int32(uint32(v12) >> 25)
				m.memory[uint32(t107)] = byte(v6)
				m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(v6)
			}
		l26:
			v6 = v3 + i32(1)
			if v3 != v4 {
				goto l33
			}
			p108 := v7
			if uint32(v4) < uint32(i32(8)) {
				p108 = v4
			}
			v3 = p108
		}
	l20:
		store32(m.memory[int64(uint32(v0))+8:], uint32(v3-v2))
	l15:
		return i32(-1)
	}
l5:
	m.fn28(i32(1271248), i32(57), i32(1271276))
	panic("unreachable")
}
func (m *Module) fn113(v0, v1, v2 int64) int64 {
	var v3 int32
	var v4, v5, v6 int64
	t0 := m.g0
	v3 = t0 - i32(80)
	m.g0 = v3
	store64(m.memory[int64(uint32(v3))+56:], uint64(i64(0)))
	store64(m.memory[int64(uint32(v3))+64:], uint64(i64(0)))
	store64(m.memory[int64(uint32(v3))+48:], uint64(v1))
	store64(m.memory[int64(uint32(v3))+32:], uint64(v1^i64(8387220255154660723)))
	store64(m.memory[int64(uint32(v3))+24:], uint64(v1^i64(7237128888997146477)))
	store64(m.memory[int64(uint32(v3))+40:], uint64(v0))
	store64(m.memory[int64(uint32(v3))+16:], uint64(v0^i64(0x6c7967656e657261)))
	store64(m.memory[int64(uint32(v3))+8:], uint64(v0^i64(8317987319222330741)))
	store64(m.memory[int64(uint32(v3))+72:], uint64(v2))
	m.fn59(v3+i32(8), v3+i32(72), i32(8))
	t1 := int64(load64(m.memory[int64(uint32(v3))+8:]))
	v0 = t1
	t2 := int64(load64(m.memory[int64(uint32(v3))+24:]))
	v1 = t2
	t3 := int64(load32(m.memory[int64(uint32(v3))+64:]))
	v2 = t3
	t4 := int64(load64(m.memory[int64(uint32(v3))+56:]))
	v4 = t4
	t5 := int64(load64(m.memory[int64(uint32(v3))+32:]))
	v5 = t5
	t6 := int64(load64(m.memory[int64(uint32(v3))+16:]))
	v6 = t6
	m.g0 = v3 + i32(80)
	t7 := v5
	v2 = v4 | v2<<56
	v4 = t7 ^ v2
	t8 := i64_rotl(v4, i64(16))
	v4 = v4 + v6
	v5 = t8 ^ v4
	t9 := i64_rotl(v5, i64(21))
	t10 := v5
	v0 = v1 + v0
	v5 = t10 + i64_rotl(v0, i64(32))
	v6 = t9 ^ v5
	t11 := i64_rotl(v6, i64(16))
	t12 := v6
	t13 := v4
	v1 = i64_rotl(v1, i64(13)) ^ v0
	v0 = t13 + v1
	v4 = t12 + (i64_rotl(v0, i64(32)) ^ i64(255))
	v6 = t11 ^ v4
	t14 := i64_rotl(v6, i64(21))
	t15 := v6
	t16 := v5 ^ v2
	v1 = v0 ^ i64_rotl(v1, i64(17))
	v0 = t16 + v1
	v2 = t15 + i64_rotl(v0, i64(32))
	v5 = t14 ^ v2
	t17 := i64_rotl(v5, i64(16))
	t18 := v5
	v1 = v0 ^ i64_rotl(v1, i64(13))
	v0 = v1 + v4
	v4 = t18 + i64_rotl(v0, i64(32))
	v5 = t17 ^ v4
	t19 := i64_rotl(v5, i64(21))
	t20 := v5
	v1 = v0 ^ i64_rotl(v1, i64(17))
	v0 = v1 + v2
	v2 = t20 + i64_rotl(v0, i64(32))
	v5 = t19 ^ v2
	t21 := i64_rotl(v5, i64(16))
	t22 := v5
	v1 = i64_rotl(v1, i64(13)) ^ v0
	v0 = v1 + v4
	v4 = t22 + i64_rotl(v0, i64(32))
	t23 := i64_rotl(t21^v4, i64(21))
	v1 = i64_rotl(v1, i64(17)) ^ v0
	v1 = i64_rotl(v1, i64(13)) ^ (v1 + v2)
	t24 := t23 ^ i64_rotl(v1, i64(17))
	v1 = v1 + v4
	return t24 ^ i64_rotl(v1, i64(32)) ^ v1
}
func (m *Module) fn114(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7, v8 int32
	var v9 int64
	var v10, v11 int32
	var v12, v13 int64
	var v14, v15 int32
	var v16 int64
	var v17 int32
	var v18 int64
	var v19, v20, v21 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
		v2 = t0
		v3 = v2 + i32(1)
		if v3 == 0 {
			m.fn28(i32(1271248), i32(57), i32(1271276))
			panic("unreachable")
		}
		{
			t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t2 := v3
			v4 = t1
			t3 := v4
			v5 = v4 + i32(1)
			v6 = int32(uint32(v5) >> 3)
			v7 = v6 * i32(7)
			p4 := v7
			if uint32(v4) < uint32(i32(8)) {
				p4 = t3
			}
			v8 = p4
			if uint32(t2) <= uint32(int32(uint32(v8)>>1)) {
				goto l1
			}
			{
				{
					{
						v8 = v8 + i32(1)
						p5 := v3
						if uint32(v8) > uint32(v3) {
							p5 = v8
						}
						v3 = p5
						if uint32(v3) < uint32(i32(15)) {
							goto l2
						}
						if uint32(v3) > uint32(i32(0x1fffffff)) {
							m.fn28(i32(1271248), i32(57), i32(1271276))
							panic("unreachable")
						}
						t6 := int32(uint32(v3<<3) / uint32(i32(7)))
						v3 = i32_shr_u(i32(-1), int32(bits.LeadingZeros32(uint32(t6+i32(-1))))) + i32(1)
						goto l4
					}
				l2:
					p7 := v3&i32(8) + i32(8)
					if uint32(v3) < uint32(i32(4)) {
						p7 = i32(4)
					}
					v3 = p7
				}
			l4:
				v9 = int64(uint32(v3)) * i64(480)
				if int32(int64(uint64(v9)>>32)) != 0 {
					goto l5
				}
				v8 = v3 + i32(8)
				t8 := v8
				v10 = int32(v9)
				v6 = t8 + v10
				if uint32(v6) < uint32(v8) {
					goto l5
				}
				if uint32(v6) > uint32(i32(0x7ffffff8)) {
					goto l5
				}
				t9 := m.fn5(v6)
				v5 = t9
				if v5 != 0 {
					v6 = v5 + v10
					if v8 == 0 {
						goto l7
					}
					memory_fill(m.memory, uint32(v6), i32(255), uint32(v8))
				l7:
					v5 = v3 + i32(-1)
					p10 := int32(uint32(v3)>>3) * i32(7)
					if uint32(v3) < uint32(i32(9)) {
						p10 = v5
					}
					v7 = p10
					t11 := int32(load32(m.memory[uint32(v0):]))
					v11 = t11
					{
						if v2 == 0 {
							goto l8
						}
						t12 := int64(load64(m.memory[uint32(v11):]))
						v9 = (t12 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
						t13 := int64(load64(m.memory[int64(uint32(v1))+8:]))
						v12 = t13
						t14 := int64(load64(m.memory[uint32(v1):]))
						v13 = t14
						v8 = v11
						v1 = v2
						v3 = i32(0)
					l14:
						{
							if v9 != i64(0) {
								goto l9
							}
						l10:
							{
								v3 = v3 + i32(8)
								v8 = v8 + i32(8)
								t15 := int64(load64(m.memory[uint32(v8):]))
								v9 = t15 & i64(-0x7f7f7f7f7f7f7f80)
								if v9 == i64(-0x7f7f7f7f7f7f7f80) {
									goto l10
								}
							}
							v9 = v9 ^ i64(-0x7f7f7f7f7f7f7f80)
						l9:
							{
								t16 := v6
								t17 := v5
								t18 := v13
								t19 := v12
								t20 := v11
								v14 = int32(uint32(int64(bits.TrailingZeros64(uint64(v9))))>>3) + v3
								t21 := int64(load64(m.memory[uint32(t20+(i32(0)-v14)*i32(480)+i32(-480)):]))
								t22 := m.fn113(t18, t19, t21)
								v15 = int32(t22)
								v10 = t17 & v15
								t23 := int64(load64(m.memory[uint32(t16+v10):]))
								v16 = t23 & i64(-0x7f7f7f7f7f7f7f80)
								if v16 != i64(0) {
									goto l11
								}
								v17 = i32(8)
							l12:
								{
									v10 = v10 + v17
									v17 = v17 + i32(8)
									t24 := v6
									v10 = v10 & v5
									t25 := int64(load64(m.memory[uint32(t24+v10):]))
									v16 = t25 & i64(-0x7f7f7f7f7f7f7f80)
									if v16 == 0 {
										goto l12
									}
								}
							}
						l11:
							v18 = v9 + i64(-1)
							{
								t26 := v6
								v10 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v16))))>>3) + v10) & v5
								t27 := int32(int8(m.memory[uint32(t26+v10)]))
								if t27 < i32(0) {
									goto l13
								}
								t28 := int64(load64(m.memory[uint32(v6):]))
								v10 = int32(uint32(int64(bits.TrailingZeros64(uint64(t28&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
							}
						l13:
							v9 = v18 & v9
							t29 := v6 + v10
							v15 = int32(uint32(v15) >> 25)
							m.memory[uint32(t29)] = byte(v15)
							m.memory[uint32(v6+(v10+i32(-8))&v5+i32(8))] = byte(v15)
							memory_copy(m.memory, uint32(v6+(v10^i32(-1))*i32(480)), uint32(v11+(v14^i32(-1))*i32(480)), uint32(i32(480)))
							v1 = v1 + i32(-1)
							if v1 != 0 {
								goto l14
							}
						}
					}
				l8:
					store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
					store32(m.memory[uint32(v0):], uint32(v6))
					store32(m.memory[int64(uint32(v0))+8:], uint32(v7-v2))
					if v4 == 0 {
						goto l15
					}
					t30 := v4
					v8 = (v4*i32(480) + i32(487)) & i32(-32)
					v3 = t30 + v8 + i32(9)
					if v3 == 0 {
						goto l15
					}
					v4 = v11 - v8
					t31 := int32(load32(m.memory[uint32(v4+i32(-4)):]))
					v8 = t31
					v6 = v8 & i32(-8)
					t32 := v6
					v8 = v8 & i32(3)
					p33 := i32(8)
					if v8 != 0 {
						p33 = i32(4)
					}
					if uint32(t32) < uint32(p33+v3) {
						m.fn3(i32(1273840), i32(46), i32(1273888))
						panic("unreachable")
					}
					if v8 == 0 {
						goto l17
					}
					if uint32(v6) > uint32(v3+i32(39)) {
						m.fn3(i32(1273904), i32(46), i32(1273952))
						panic("unreachable")
					}
				l17:
					m.fn1(v4)
					return i32(-1)
				}
				m.fn24(i32(8), v6)
				panic("unreachable")
			}
		}
	l1:
		{
			if v5 != 0 {
				goto l19
			}
			v3 = i32(0)
			goto l20
		l19:
			t34 := int32(load32(m.memory[uint32(v0):]))
			v8 = t34
			v3 = i32(0)
			{
				{
					t35 := v6
					var p36 int32
					if v5&i32(7) != i32(0) {
						p36 = 1
					}
					v6 = t35 + p36
					if v6 == i32(1) {
						goto l21
					}
					v11 = v6 & i32(1)
					v10 = v6 & i32(0x3ffffffe)
					v3 = i32(0)
				l22:
					{
						v6 = v8 + v3
						t37 := int64(load64(m.memory[uint32(v6):]))
						t38 := v6
						v9 = t37
						store64(m.memory[uint32(t38):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
						v6 = v6 + i32(8)
						t39 := int64(load64(m.memory[uint32(v6):]))
						t40 := v6
						v9 = t39
						store64(m.memory[uint32(t40):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
						v3 = v3 + i32(16)
						v10 = v10 + i32(-2)
						if v10 != 0 {
							goto l22
						}
					}
					if v11 == 0 {
						goto l23
					}
				}
			l21:
				v3 = v8 + v3
				t41 := int64(load64(m.memory[uint32(v3):]))
				t42 := v3
				v9 = t41
				store64(m.memory[uint32(t42):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
			}
		l23:
			{
				if uint32(v5) < uint32(i32(8)) {
					goto l24
				}
				t43 := int64(load64(m.memory[uint32(v8):]))
				store64(m.memory[uint32(v8+v5):], uint64(t43))
				goto l25
			}
		l24:
			if v5 == 0 {
				goto l25
			}
			memory_copy(m.memory, uint32(v8+i32(8)), uint32(v8), uint32(v5))
		l25:
			t44 := int64(load64(m.memory[int64(uint32(v1))+8:]))
			v16 = t44
			t45 := int64(load64(m.memory[uint32(v1):]))
			v18 = t45
			v6 = v8
			v10 = i32(0)
		l34:
			{
				t46 := v8
				v3 = v10
				v5 = t46 + v3
				t47 := int32(m.memory[uint32(v5)])
				if t47 != i32(128) {
					goto l26
				}
				v15 = v8 + (v3^i32(-1))*i32(480)
				v14 = v8 + (i32(0)-v3)*i32(480) + i32(-480)
				{
				l33:
					{
						t48 := int64(load64(m.memory[uint32(v14):]))
						t49 := m.fn113(v18, v16, t48)
						t50 := v4
						v1 = int32(t49)
						v10 = t50 & v1
						v11 = v10
						{
							t51 := int64(load64(m.memory[uint32(v8+v10):]))
							v9 = t51 & i64(-0x7f7f7f7f7f7f7f80)
							if v9 != i64(0) {
								goto l27
							}
							v17 = i32(8)
							v11 = v10
						l28:
							{
								v11 = v11 + v17
								v17 = v17 + i32(8)
								t52 := v8
								v11 = v11 & v4
								t53 := int64(load64(m.memory[uint32(t52+v11):]))
								v9 = t53 & i64(-0x7f7f7f7f7f7f7f80)
								if v9 == 0 {
									goto l28
								}
							}
						}
					l27:
						{
							t54 := v8
							v11 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v9))))>>3) + v11) & v4
							t55 := int32(int8(m.memory[uint32(t54+v11)]))
							if t55 < i32(0) {
								goto l29
							}
							t56 := int64(load64(m.memory[uint32(v8):]))
							v11 = int32(uint32(int64(bits.TrailingZeros64(uint64(t56&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
						}
					l29:
						{
							if uint32((v11-v10^(v3-v10))&v4) < uint32(i32(8)) {
								goto l30
							}
							v10 = v8 + v11
							t57 := int32(m.memory[uint32(v10)])
							v17 = t57
							t58 := v10
							v1 = int32(uint32(v1) >> 25)
							m.memory[uint32(t58)] = byte(v1)
							m.memory[uint32(v8+(v11+i32(-8))&v4+i32(8))] = byte(v1)
							if v17 == i32(255) {
								goto l31
							}
							v1 = i32(-480)
							v19 = v8 + v11*i32(-480)
						l32:
							{
								v10 = v6 + v1
								t59 := int32(load32(m.memory[uint32(v10):]))
								v17 = t59
								t60 := v10
								v11 = v19 + v1
								t61 := int32(load32(m.memory[uint32(v11):]))
								store32(m.memory[uint32(t60):], uint32(t61))
								store32(m.memory[uint32(v11):], uint32(v17))
								v17 = v11 + i32(4)
								t62 := int32(load32(m.memory[uint32(v17):]))
								v20 = t62
								t63 := v17
								v21 = v10 + i32(4)
								t64 := int32(load32(m.memory[uint32(v21):]))
								store32(m.memory[uint32(t63):], uint32(t64))
								store32(m.memory[uint32(v21):], uint32(v20))
								v17 = v10 + i32(8)
								t65 := int32(load32(m.memory[uint32(v17):]))
								v20 = t65
								t66 := v17
								v21 = v11 + i32(8)
								t67 := int32(load32(m.memory[uint32(v21):]))
								store32(m.memory[uint32(t66):], uint32(t67))
								store32(m.memory[uint32(v21):], uint32(v20))
								v11 = v11 + i32(12)
								t68 := int32(load32(m.memory[uint32(v11):]))
								v17 = t68
								t69 := v11
								v10 = v10 + i32(12)
								t70 := int32(load32(m.memory[uint32(v10):]))
								store32(m.memory[uint32(t69):], uint32(t70))
								store32(m.memory[uint32(v10):], uint32(v17))
								v1 = v1 + i32(16)
								if v1 != 0 {
									goto l32
								}
								goto l33
							}
						}
					l30:
					}
					t71 := v5
					v10 = int32(uint32(v1) >> 25)
					m.memory[uint32(t71)] = byte(v10)
					m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(v10)
					goto l26
				}
			l31:
				m.memory[uint32(v5)] = byte(i32(255))
				m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(i32(255))
				memory_copy(m.memory, uint32(v8+(v11^i32(-1))*i32(480)), uint32(v15), uint32(i32(480)))
			}
		l26:
			v10 = v3 + i32(1)
			v6 = v6 + i32(-480)
			if v3 != v4 {
				goto l34
			}
			p72 := v7
			if uint32(v4) < uint32(i32(8)) {
				p72 = v4
			}
			v3 = p72
		}
	l20:
		store32(m.memory[int64(uint32(v0))+8:], uint32(v3-v2))
	l15:
		return i32(-1)
	l5:
		m.fn28(i32(1271248), i32(57), i32(1271276))
		panic("unreachable")
	}
}
func (m *Module) fn115(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7, v8 int32
	var v9 int64
	var v10, v11, v12 int32
	var v13, v14 int64
	var v15 int32
	var v16 int64
	var v17 int32
	var v18 int64
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
		v2 = t0
		v3 = v2 + i32(1)
		if v3 == 0 {
			m.fn28(i32(1271248), i32(57), i32(1271276))
			panic("unreachable")
		}
		{
			t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t2 := v3
			v4 = t1
			t3 := v4
			v5 = v4 + i32(1)
			v6 = int32(uint32(v5) >> 3)
			v7 = v6 * i32(7)
			p4 := v7
			if uint32(v4) < uint32(i32(8)) {
				p4 = t3
			}
			v8 = p4
			if uint32(t2) <= uint32(int32(uint32(v8)>>1)) {
				goto l1
			}
			{
				{
					{
						v8 = v8 + i32(1)
						p5 := v3
						if uint32(v8) > uint32(v3) {
							p5 = v8
						}
						v3 = p5
						if uint32(v3) < uint32(i32(15)) {
							goto l2
						}
						if uint32(v3) > uint32(i32(0x1fffffff)) {
							m.fn28(i32(1271248), i32(57), i32(1271276))
							panic("unreachable")
						}
						t6 := int32(uint32(v3<<3) / uint32(i32(7)))
						v3 = i32_shr_u(i32(-1), int32(bits.LeadingZeros32(uint32(t6+i32(-1))))) + i32(1)
						goto l4
					}
				l2:
					p7 := v3&i32(8) + i32(8)
					if uint32(v3) < uint32(i32(4)) {
						p7 = i32(4)
					}
					v3 = p7
				}
			l4:
				v9 = int64(uint32(v3)) * i64(24)
				if int32(int64(uint64(v9)>>32)) != 0 {
					goto l5
				}
				v8 = v3 + i32(8)
				t8 := v8
				v10 = int32(v9)
				v6 = t8 + v10
				if uint32(v6) < uint32(v8) {
					goto l5
				}
				if uint32(v6) > uint32(i32(0x7ffffff8)) {
					goto l5
				}
				t9 := m.fn5(v6)
				v5 = t9
				if v5 != 0 {
					v6 = v5 + v10
					if v8 == 0 {
						goto l7
					}
					memory_fill(m.memory, uint32(v6), i32(255), uint32(v8))
				l7:
					v5 = v3 + i32(-1)
					p10 := int32(uint32(v3)>>3) * i32(7)
					if uint32(v3) < uint32(i32(9)) {
						p10 = v5
					}
					v11 = p10
					t11 := int32(load32(m.memory[uint32(v0):]))
					v12 = t11
					{
						if v2 == 0 {
							goto l8
						}
						t12 := int64(load64(m.memory[uint32(v12):]))
						v9 = (t12 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
						t13 := int64(load64(m.memory[int64(uint32(v1))+8:]))
						v13 = t13
						t14 := int64(load64(m.memory[uint32(v1):]))
						v14 = t14
						v8 = v12
						v15 = v2
						v3 = i32(0)
					l14:
						{
							if v9 != i64(0) {
								goto l9
							}
						l10:
							{
								v3 = v3 + i32(8)
								v8 = v8 + i32(8)
								t15 := int64(load64(m.memory[uint32(v8):]))
								v9 = t15 & i64(-0x7f7f7f7f7f7f7f80)
								if v9 == i64(-0x7f7f7f7f7f7f7f80) {
									goto l10
								}
							}
							v9 = v9 ^ i64(-0x7f7f7f7f7f7f7f80)
						l9:
							{
								t16 := v6
								t17 := v5
								t18 := v14
								t19 := v13
								t20 := v12
								v1 = int32(uint32(int64(bits.TrailingZeros64(uint64(v9))))>>3) + v3
								t21 := int64(load64(m.memory[uint32(t20+(i32(0)-v1)*i32(24)+i32(-24)):]))
								t22 := m.fn113(t18, t19, t21)
								v7 = int32(t22)
								v10 = t17 & v7
								t23 := int64(load64(m.memory[uint32(t16+v10):]))
								v16 = t23 & i64(-0x7f7f7f7f7f7f7f80)
								if v16 != i64(0) {
									goto l11
								}
								v17 = i32(8)
							l12:
								{
									v10 = v10 + v17
									v17 = v17 + i32(8)
									t24 := v6
									v10 = v10 & v5
									t25 := int64(load64(m.memory[uint32(t24+v10):]))
									v16 = t25 & i64(-0x7f7f7f7f7f7f7f80)
									if v16 == 0 {
										goto l12
									}
								}
							}
						l11:
							v18 = v9 + i64(-1)
							{
								t26 := v6
								v10 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v16))))>>3) + v10) & v5
								t27 := int32(int8(m.memory[uint32(t26+v10)]))
								if t27 < i32(0) {
									goto l13
								}
								t28 := int64(load64(m.memory[uint32(v6):]))
								v10 = int32(uint32(int64(bits.TrailingZeros64(uint64(t28&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
							}
						l13:
							v9 = v18 & v9
							t29 := v6 + v10
							v7 = int32(uint32(v7) >> 25)
							m.memory[uint32(t29)] = byte(v7)
							m.memory[uint32(v6+(v10+i32(-8))&v5+i32(8))] = byte(v7)
							v10 = v6 + (v10^i32(-1))*i32(24)
							t30 := v10
							v1 = v12 + (v1^i32(-1))*i32(24)
							t31 := int64(load64(m.memory[int64(uint32(v1))+16:]))
							store64(m.memory[int64(uint32(t30))+16:], uint64(t31))
							t32 := int64(load64(m.memory[int64(uint32(v1))+8:]))
							store64(m.memory[int64(uint32(v10))+8:], uint64(t32))
							t33 := int64(load64(m.memory[uint32(v1):]))
							store64(m.memory[uint32(v10):], uint64(t33))
							v15 = v15 + i32(-1)
							if v15 != 0 {
								goto l14
							}
						}
					}
				l8:
					store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
					store32(m.memory[uint32(v0):], uint32(v6))
					store32(m.memory[int64(uint32(v0))+8:], uint32(v11-v2))
					if v4 == 0 {
						goto l15
					}
					t34 := v4
					v8 = (v4*i32(24) + i32(31)) & i32(-8)
					v3 = t34 + v8 + i32(9)
					if v3 == 0 {
						goto l15
					}
					v4 = v12 - v8
					t35 := int32(load32(m.memory[uint32(v4+i32(-4)):]))
					v8 = t35
					v6 = v8 & i32(-8)
					t36 := v6
					v8 = v8 & i32(3)
					p37 := i32(8)
					if v8 != 0 {
						p37 = i32(4)
					}
					if uint32(t36) < uint32(p37+v3) {
						m.fn3(i32(1273840), i32(46), i32(1273888))
						panic("unreachable")
					}
					if v8 == 0 {
						goto l17
					}
					if uint32(v6) > uint32(v3+i32(39)) {
						m.fn3(i32(1273904), i32(46), i32(1273952))
						panic("unreachable")
					}
				l17:
					m.fn1(v4)
					return i32(-1)
				}
				m.fn24(i32(8), v6)
				panic("unreachable")
			}
		}
	l1:
		{
			if v5 != 0 {
				goto l19
			}
			v3 = i32(0)
			goto l20
		l19:
			t38 := int32(load32(m.memory[uint32(v0):]))
			v8 = t38
			v3 = i32(0)
			{
				{
					t39 := v6
					var p40 int32
					if v5&i32(7) != i32(0) {
						p40 = 1
					}
					v6 = t39 + p40
					if v6 == i32(1) {
						goto l21
					}
					v12 = v6 & i32(1)
					v10 = v6 & i32(0x3ffffffe)
					v3 = i32(0)
				l22:
					{
						v6 = v8 + v3
						t41 := int64(load64(m.memory[uint32(v6):]))
						t42 := v6
						v9 = t41
						store64(m.memory[uint32(t42):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
						v6 = v6 + i32(8)
						t43 := int64(load64(m.memory[uint32(v6):]))
						t44 := v6
						v9 = t43
						store64(m.memory[uint32(t44):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
						v3 = v3 + i32(16)
						v10 = v10 + i32(-2)
						if v10 != 0 {
							goto l22
						}
					}
					if v12 == 0 {
						goto l23
					}
				}
			l21:
				v3 = v8 + v3
				t45 := int64(load64(m.memory[uint32(v3):]))
				t46 := v3
				v9 = t45
				store64(m.memory[uint32(t46):], uint64(int64(uint64(v9^i64(-1))>>7)&i64(72340172838076673)+(v9|i64(0x7f7f7f7f7f7f7f7f))))
			}
		l23:
			{
				if uint32(v5) < uint32(i32(8)) {
					goto l24
				}
				t47 := int64(load64(m.memory[uint32(v8):]))
				store64(m.memory[uint32(v8+v5):], uint64(t47))
				goto l25
			}
		l24:
			if v5 == 0 {
				goto l25
			}
			memory_copy(m.memory, uint32(v8+i32(8)), uint32(v8), uint32(v5))
		l25:
			t48 := int64(load64(m.memory[int64(uint32(v1))+8:]))
			v16 = t48
			t49 := int64(load64(m.memory[uint32(v1):]))
			v18 = t49
			v6 = i32(0)
		l33:
			{
				t50 := v8
				v3 = v6
				v10 = t50 + v3
				t51 := int32(m.memory[uint32(v10)])
				if t51 != i32(128) {
					goto l26
				}
				v15 = v8 + (v3^i32(-1))*i32(24)
				v12 = v8 + (i32(0)-v3)*i32(24) + i32(-24)
			l32:
				{
					t52 := int64(load64(m.memory[uint32(v12):]))
					t53 := m.fn113(v18, v16, t52)
					t54 := v4
					v1 = int32(t53)
					v6 = t54 & v1
					v5 = v6
					{
						t55 := int64(load64(m.memory[uint32(v8+v6):]))
						v9 = t55 & i64(-0x7f7f7f7f7f7f7f80)
						if v9 != i64(0) {
							goto l27
						}
						v11 = i32(8)
						v5 = v6
					l28:
						{
							v5 = v5 + v11
							v11 = v11 + i32(8)
							t56 := v8
							v5 = v5 & v4
							t57 := int64(load64(m.memory[uint32(t56+v5):]))
							v9 = t57 & i64(-0x7f7f7f7f7f7f7f80)
							if v9 == 0 {
								goto l28
							}
						}
					}
				l27:
					{
						t58 := v8
						v5 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v9))))>>3) + v5) & v4
						t59 := int32(int8(m.memory[uint32(t58+v5)]))
						if t59 < i32(0) {
							goto l29
						}
						t60 := int64(load64(m.memory[uint32(v8):]))
						v5 = int32(uint32(int64(bits.TrailingZeros64(uint64(t60&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
					}
				l29:
					{
						if uint32((v5-v6^(v3-v6))&v4) < uint32(i32(8)) {
							goto l30
						}
						v6 = v8 + v5
						t61 := int32(m.memory[uint32(v6)])
						v11 = t61
						t62 := v6
						v1 = int32(uint32(v1) >> 25)
						m.memory[uint32(t62)] = byte(v1)
						m.memory[uint32(v8+(v5+i32(-8))&v4+i32(8))] = byte(v1)
						v6 = v8 + (v5^i32(-1))*i32(24)
						{
							if v11 != i32(255) {
								t66 := int64(load64(m.memory[uint32(v15):]))
								v9 = t66
								t67 := int64(load64(m.memory[uint32(v6):]))
								store64(m.memory[uint32(v15):], uint64(t67))
								store64(m.memory[uint32(v6):], uint64(v9))
								t68 := int64(load64(m.memory[int64(uint32(v6))+8:]))
								v9 = t68
								t69 := int64(load64(m.memory[int64(uint32(v15))+8:]))
								store64(m.memory[int64(uint32(v6))+8:], uint64(t69))
								store64(m.memory[int64(uint32(v15))+8:], uint64(v9))
								t70 := int32(load32(m.memory[int64(uint32(v15))+16:]))
								v5 = t70
								t71 := int32(load32(m.memory[int64(uint32(v6))+16:]))
								store32(m.memory[int64(uint32(v15))+16:], uint32(t71))
								t72 := int32(load32(m.memory[int64(uint32(v6))+20:]))
								v1 = t72
								t73 := int32(load32(m.memory[int64(uint32(v15))+20:]))
								store32(m.memory[int64(uint32(v6))+20:], uint32(t73))
								store32(m.memory[int64(uint32(v15))+20:], uint32(v1))
								store32(m.memory[int64(uint32(v6))+16:], uint32(v5))
								goto l32
							}
							m.memory[uint32(v10)] = byte(i32(255))
							m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(i32(255))
							t63 := int64(load64(m.memory[int64(uint32(v15))+16:]))
							store64(m.memory[int64(uint32(v6))+16:], uint64(t63))
							t64 := int64(load64(m.memory[int64(uint32(v15))+8:]))
							store64(m.memory[int64(uint32(v6))+8:], uint64(t64))
							t65 := int64(load64(m.memory[uint32(v15):]))
							store64(m.memory[uint32(v6):], uint64(t65))
							goto l26
						}
					}
				l30:
				}
				t74 := v10
				v6 = int32(uint32(v1) >> 25)
				m.memory[uint32(t74)] = byte(v6)
				m.memory[uint32(v8+v4&(v3+i32(-8))+i32(8))] = byte(v6)
			}
		l26:
			v6 = v3 + i32(1)
			if v3 != v4 {
				goto l33
			}
			p75 := v7
			if uint32(v4) < uint32(i32(8)) {
				p75 = v4
			}
			v3 = p75
		}
	l20:
		store32(m.memory[int64(uint32(v0))+8:], uint32(v3-v2))
	l15:
		return i32(-1)
	}
l5:
	m.fn28(i32(1271248), i32(57), i32(1271276))
	panic("unreachable")
}
func (m *Module) fn116(v0, v1 int32) {
	var v2, v3, v4, v5 int32
	t0 := m.g0
	v2 = t0 - i32(4096)
	m.g0 = v2
	{
		{
			p1 := i32(111111)
			if uint32(v1) < uint32(i32(111111)) {
				p1 = v1
			}
			v3 = p1
			t2 := v3
			v4 = v1 - int32(uint32(v1)>>1)
			p3 := v4
			if uint32(v3) > uint32(v4) {
				p3 = t2
			}
			v3 = p3
			if uint32(v3) < uint32(i32(57)) {
				goto l0
			}
			if uint32(v4) >= uint32(i32(29826162)) {
				m.fn9()
				panic("unreachable")
			}
			v5 = v3 * i32(72)
			t4 := m.fn5(v5)
			v4 = t4
			if v4 == 0 {
				m.fn10(i32(4), v5)
				panic("unreachable")
			}
			store32(m.memory[int64(uint32(v2))+8:], uint32(i32(0)))
			store32(m.memory[int64(uint32(v2))+4:], uint32(v4))
			store32(m.memory[uint32(v2):], uint32(v3))
			t5 := v0
			t6 := v1
			t7 := v4
			t8 := v3
			var p9 int32
			if uint32(v1) < uint32(i32(65)) {
				p9 = 1
			}
			m.fn117(t5, t6, t7, t8, p9)
			m.fn118(v2)
			goto l3
		}
	l0:
		t10 := v0
		t11 := v1
		t12 := v2
		var p13 int32
		if uint32(v1) < uint32(i32(65)) {
			p13 = 1
		}
		m.fn117(t10, t11, t12, i32(56), p13)
	}
l3:
	m.g0 = v2 + i32(4096)
}
func (m *Module) fn117(v0, v1, v2, v3, v4 int32) {
	var v5 int32
	var v6, v7 int64
	var v8, v9, v10, v11, v12, v13, v14, v15, v16, v17, v18, v19, v20, v21, v22, v23, v24, v25 int32
	t0 := m.g0
	v5 = t0 - i32(336)
	m.g0 = v5
	v6 = int64(uint32(v1))
	t1 := int64(uint64(i64(0x4000000000000000)) / uint64(v6))
	v7 = t1
	var p2 int32
	if v7*v6 != i64(0x4000000000000000) {
		p2 = 1
	}
	v6 = int64(uint32(p2))
	{
		if uint32(v1) < uint32(i32(4097)) {
			goto l0
		}
		v8 = int32(bits.LeadingZeros32(uint32(v1|i32(1)))) ^ i32(31)
		v8 = int32(uint32(v8)>>1) + v8&i32(1)
		v9 = int32(uint32(i32_shl(i32(1), v8)+i32_shr_u(v1, v8)) >> 1)
		goto l1
	l0:
		v8 = v1 - int32(uint32(v1)>>1)
		p3 := i32(64)
		if uint32(v8) < uint32(i32(64)) {
			p3 = v8
		}
		v9 = p3
	}
l1:
	v7 = v7 + v6
	v10 = v0 + i32(-72)
	v11 = v0 + i32(208)
	v8 = i32(1)
	v12 = i32(0)
	v13 = i32(0)
l32:
	v14 = i32(0)
	v15 = i32(1)
	{
		var p4 int32
		if uint32(v1) > uint32(v12) {
			p4 = 1
		}
		v16 = p4
		if v16 == 0 {
			goto l2
		}
		t5 := v0
		v17 = v12 * i32(72)
		v18 = t5 + v17
		{
			v19 = v1 - v12
			if uint32(v19) < uint32(v9) {
				goto l3
			}
			if uint32(v19) >= uint32(i32(2)) {
				goto l4
			}
			v20 = v19
			goto l5
		l4:
			{
				t6 := int32(load32(m.memory[uint32(v18+i32(136)):]))
				v21 = t6
				t7 := int32(load32(m.memory[uint32(v18+i32(64)):]))
				var p8 int32
				if uint32(v21) < uint32(t7) {
					p8 = 1
				}
				v14 = p8
				if v14 != 0 {
					goto l6
				}
				v20 = i32(2)
				if v19 == i32(2) {
					goto l5
				}
				v22 = v11 + v17
				v20 = i32(2)
			l8:
				{
					t9 := int32(load32(m.memory[uint32(v22):]))
					v23 = t9
					if uint32(v23) < uint32(v21) {
						goto l7
					}
					v22 = v22 + i32(72)
					v21 = v23
					t10 := v19
					v20 = v20 + i32(1)
					if t10 != v20 {
						goto l8
					}
					goto l9
				}
			}
		l6:
			v20 = i32(2)
			v24 = i32(1)
			if v19 == i32(2) {
				goto l10
			}
			v22 = v11 + v17
			v20 = i32(2)
		l11:
			{
				t11 := int32(load32(m.memory[uint32(v22):]))
				v23 = t11
				if uint32(v23) >= uint32(v21) {
					goto l7
				}
				v22 = v22 + i32(72)
				v21 = v23
				t12 := v19
				v20 = v20 + i32(1)
				if t12 != v20 {
					goto l11
				}
			}
		l9:
			v20 = v19
		l7:
			if uint32(v20) < uint32(v9) {
				goto l3
			}
			if v14 == 0 {
				goto l5
			}
			v24 = int32(uint32(v20) >> 1)
			if v24 == 0 {
				goto l5
			}
		l10:
			v22 = v10 + v20*i32(72)
			v23 = v0
		l12:
			{
				v19 = v23 + v17
				v18 = v19 + i32(8)
				t13 := int64(load64(m.memory[uint32(v18):]))
				v6 = t13
				t14 := v18
				v21 = v22 + v17
				v14 = v21 + i32(8)
				t15 := int64(load64(m.memory[uint32(v14):]))
				store64(m.memory[uint32(t14):], uint64(t15))
				store64(m.memory[uint32(v14):], uint64(v6))
				t16 := int32(load32(m.memory[uint32(v21+i32(20)):]))
				v18 = t16
				v14 = v21 + i32(16)
				t17 := int32(load32(m.memory[uint32(v14):]))
				v15 = t17
				t18 := v14
				v25 = v19 + i32(16)
				t19 := int64(load64(m.memory[uint32(v25):]))
				store64(m.memory[uint32(t18):], uint64(t19))
				t20 := int64(load64(m.memory[uint32(v19):]))
				v6 = t20
				t21 := int64(load64(m.memory[uint32(v21):]))
				store64(m.memory[uint32(v19):], uint64(t21))
				store64(m.memory[uint32(v21):], uint64(v6))
				store32(m.memory[uint32(v25):], uint32(v15))
				store32(m.memory[uint32(v19+i32(20)):], uint32(v18))
				v18 = v21 + i32(24)
				t22 := int32(load32(m.memory[uint32(v18):]))
				v14 = t22
				t23 := v18
				v15 = v19 + i32(24)
				t24 := int32(load32(m.memory[uint32(v15):]))
				store32(m.memory[uint32(t23):], uint32(t24))
				store32(m.memory[uint32(v15):], uint32(v14))
				v18 = v21 + i32(28)
				t25 := int32(load32(m.memory[uint32(v18):]))
				v14 = t25
				t26 := v18
				v15 = v19 + i32(28)
				t27 := int32(load32(m.memory[uint32(v15):]))
				store32(m.memory[uint32(t26):], uint32(t27))
				store32(m.memory[uint32(v15):], uint32(v14))
				v18 = v19 + i32(32)
				t28 := int32(load32(m.memory[uint32(v18):]))
				v14 = t28
				t29 := v18
				v15 = v21 + i32(32)
				t30 := int32(load32(m.memory[uint32(v15):]))
				store32(m.memory[uint32(t29):], uint32(t30))
				store32(m.memory[uint32(v15):], uint32(v14))
				v18 = v19 + i32(36)
				t31 := int32(load32(m.memory[uint32(v18):]))
				v14 = t31
				t32 := v18
				v15 = v21 + i32(36)
				t33 := int32(load32(m.memory[uint32(v15):]))
				store32(m.memory[uint32(t32):], uint32(t33))
				store32(m.memory[uint32(v15):], uint32(v14))
				v18 = v19 + i32(40)
				t34 := int32(load32(m.memory[uint32(v18):]))
				v14 = t34
				t35 := v18
				v15 = v21 + i32(40)
				t36 := int32(load32(m.memory[uint32(v15):]))
				store32(m.memory[uint32(t35):], uint32(t36))
				store32(m.memory[uint32(v15):], uint32(v14))
				v18 = v19 + i32(44)
				t37 := int32(load32(m.memory[uint32(v18):]))
				v14 = t37
				t38 := v18
				v15 = v21 + i32(44)
				t39 := int32(load32(m.memory[uint32(v15):]))
				store32(m.memory[uint32(t38):], uint32(t39))
				store32(m.memory[uint32(v15):], uint32(v14))
				v18 = v19 + i32(48)
				t40 := int32(load32(m.memory[uint32(v18):]))
				v14 = t40
				t41 := v18
				v15 = v21 + i32(48)
				t42 := int32(load32(m.memory[uint32(v15):]))
				store32(m.memory[uint32(t41):], uint32(t42))
				store32(m.memory[uint32(v15):], uint32(v14))
				v18 = v19 + i32(52)
				t43 := int32(load32(m.memory[uint32(v18):]))
				v14 = t43
				t44 := v18
				v15 = v21 + i32(52)
				t45 := int32(load32(m.memory[uint32(v15):]))
				store32(m.memory[uint32(t44):], uint32(t45))
				store32(m.memory[uint32(v15):], uint32(v14))
				v18 = v19 + i32(56)
				t46 := int32(load32(m.memory[uint32(v18):]))
				v14 = t46
				t47 := v18
				v15 = v21 + i32(56)
				t48 := int32(load32(m.memory[uint32(v15):]))
				store32(m.memory[uint32(t47):], uint32(t48))
				store32(m.memory[uint32(v15):], uint32(v14))
				v18 = v19 + i32(60)
				t49 := int32(load32(m.memory[uint32(v18):]))
				v14 = t49
				t50 := v18
				v15 = v21 + i32(60)
				t51 := int32(load32(m.memory[uint32(v15):]))
				store32(m.memory[uint32(t50):], uint32(t51))
				store32(m.memory[uint32(v15):], uint32(v14))
				v18 = v19 + i32(64)
				t52 := int32(load32(m.memory[uint32(v18):]))
				v14 = t52
				t53 := v18
				v15 = v21 + i32(64)
				t54 := int32(load32(m.memory[uint32(v15):]))
				store32(m.memory[uint32(t53):], uint32(t54))
				store32(m.memory[uint32(v15):], uint32(v14))
				v19 = v19 + i32(68)
				t55 := int32(load32(m.memory[uint32(v19):]))
				v18 = t55
				t56 := v19
				v21 = v21 + i32(68)
				t57 := int32(load32(m.memory[uint32(v21):]))
				store32(m.memory[uint32(t56):], uint32(t57))
				store32(m.memory[uint32(v21):], uint32(v18))
				v22 = v22 + i32(-72)
				v23 = v23 + i32(72)
				v24 = v24 + i32(-1)
				if v24 != 0 {
					goto l12
				}
			}
		l5:
			v15 = v20<<1 | i32(1)
			goto l13
		l3:
			{
				if v4 != 0 {
					goto l14
				}
				p58 := v9
				if uint32(v19) < uint32(v9) {
					p58 = v19
				}
				v15 = p58 << 1
				goto l13
			}
		l14:
			t60 := v18
			p59 := i32(32)
			if uint32(v19) < uint32(i32(32)) {
				p59 = v19
			}
			v19 = p59
			m.fn119(t60, v19, v2, v3, i32(0), i32(0))
			v15 = v19<<1 | i32(1)
		}
	l13:
		v14 = int32(int64(bits.LeadingZeros64(uint64(v7*int64(uint32(int32(uint32(v15)>>1)+v12<<1)) ^ (int64(uint32(v12-int32(uint32(v8)>>1)))+int64(uint32(v12)))*v7))))
	}
l2:
	{
		if uint32(v13) < uint32(i32(2)) {
			goto l15
		}
		t61 := v10
		v19 = v12 * i32(72)
		v25 = t61 + v19
		v18 = v0 + v19
	l29:
		{
			t62 := v5 + i32(270)
			v22 = v13 + i32(-1)
			t63 := int32(m.memory[uint32(t62+v22)])
			if uint32(t63) < uint32(v14) {
				goto l15
			}
			{
				t64 := int32(load32(m.memory[uint32(v5+i32(4)+v22<<2):]))
				v23 = t64
				v19 = int32(uint32(v23) >> 1)
				t65 := v19
				v13 = int32(uint32(v8) >> 1)
				v20 = t65 + v13
				if uint32(v20) > uint32(v3) {
					goto l16
				}
				if (v23|v8)&i32(1) == 0 {
					v8 = v20 << 1
					goto l20
				}
			}
		l16:
			v21 = v0 + (v12-v20)*i32(72)
			if v23&i32(1) == 0 {
				goto l18
			}
			goto l19
		l18:
			m.fn119(v21, v19, v2, v3, int32(bits.LeadingZeros32(uint32(v19|i32(1))))<<1^i32(62), i32(0))
		l19:
			if v8&i32(1) != 0 {
				goto l21
			}
			m.fn119(v21+v19*i32(72), v13, v2, v3, int32(bits.LeadingZeros32(uint32(v13|i32(1))))<<1^i32(62), i32(0))
		l21:
			{
				if v13 == 0 {
					goto l22
				}
				if v19 == 0 {
					goto l22
				}
				t66 := v3
				t67 := v13
				t68 := v19
				var p69 int32
				if uint32(v13) < uint32(v19) {
					p69 = 1
				}
				v23 = p69
				p70 := t68
				if v23 != 0 {
					p70 = t67
				}
				v13 = p70
				if uint32(t66) < uint32(v13) {
					goto l22
				}
				v8 = v21 + v19*i32(72)
				{
					v19 = v13 * i32(72)
					if v19 == 0 {
						goto l23
					}
					t72 := v2
					p71 := v21
					if v23 != 0 {
						p71 = v8
					}
					memory_copy(m.memory, uint32(t72), uint32(p71), uint32(v19))
				}
			l23:
				v19 = v2 + v19
				if v23 != 0 {
					goto l24
				}
				v13 = v2
			l26:
				{
					t73 := int32(load32(m.memory[uint32(v8+i32(64)):]))
					t74 := v21
					t75 := v8
					t76 := v13
					v23 = t73
					t77 := int32(load32(m.memory[uint32(v13+i32(64)):]))
					t78 := v23
					v17 = t77
					var p79 int32
					if uint32(t78) < uint32(v17) {
						p79 = 1
					}
					v24 = p79
					p80 := t76
					if v24 != 0 {
						p80 = t75
					}
					memory_copy(m.memory, uint32(t74), uint32(p80), uint32(i32(72)))
					v21 = v21 + i32(72)
					t81 := v13
					var p82 int32
					if uint32(v23) >= uint32(v17) {
						p82 = 1
					}
					v13 = t81 + p82*i32(72)
					if v13 == v19 {
						goto l25
					}
					v8 = v8 + v24*i32(72)
					if v8 != v18 {
						goto l26
					}
					goto l25
				}
			l24:
				v13 = v25
			l28:
				{
					t83 := v13
					v23 = v8 + i32(-72)
					t84 := v23
					v17 = v19 + i32(-72)
					t85 := int32(load32(m.memory[uint32(v19+i32(-8)):]))
					t86 := v17
					v24 = t85
					t87 := int32(load32(m.memory[uint32(v8+i32(-8)):]))
					t88 := v24
					v8 = t87
					var p89 int32
					if uint32(t88) < uint32(v8) {
						p89 = 1
					}
					v19 = p89
					p90 := t86
					if v19 != 0 {
						p90 = t84
					}
					memory_copy(m.memory, uint32(t83), uint32(p90), uint32(i32(72)))
					v19 = v17 + v19*i32(72)
					t91 := v23
					var p92 int32
					if uint32(v24) >= uint32(v8) {
						p92 = 1
					}
					v8 = t91 + p92*i32(72)
					if v8 == v21 {
						goto l27
					}
					v13 = v13 + i32(-72)
					if v19 != v2 {
						goto l28
					}
				}
			l27:
				v21 = v8
				v13 = v2
			l25:
				v8 = v19 - v13
				if v8 == 0 {
					goto l22
				}
				memory_copy(m.memory, uint32(v21), uint32(v13), uint32(v8))
			}
		l22:
			v8 = v20<<1 | i32(1)
		l20:
			v19 = i32(1)
			v13 = v22
			if uint32(v22) > uint32(i32(1)) {
				goto l29
			}
			goto l30
		}
	}
l15:
	v19 = v13
l30:
	m.memory[uint32(v5+i32(270)+v19)] = byte(v14)
	store32(m.memory[uint32(v5+i32(4)+v19<<2):], uint32(v8))
	if v16 == 0 {
		goto l31
	}
	v13 = v19 + i32(1)
	v12 = int32(uint32(v15)>>1) + v12
	v8 = v15
	goto l32
l31:
	if v8&i32(1) != 0 {
		goto l33
	}
	m.fn119(v0, v1, v2, v3, int32(bits.LeadingZeros32(uint32(v1|i32(1))))<<1^i32(62), i32(0))
l33:
	m.g0 = v5 + i32(336)
}
func (m *Module) fn118(v0 int32) {
	var v1, v2, v3, v4, v5, v6, v7 int32
	t0 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	v1 = t0
	{
		t1 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		v2 = t1
		if v2 == 0 {
			goto l0
		}
		v3 = v1 + i32(48)
	l13:
		{
			t2 := int32(load32(m.memory[uint32(v3):]))
			v4 = t2
			if v4 == 0 {
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
				m.fn3(i32(1273840), i32(46), i32(1273888))
				panic("unreachable")
			}
			if v6 == 0 {
				goto l3
			}
			if uint32(v7) > uint32(v4+i32(39)) {
				m.fn3(i32(1273904), i32(46), i32(1273952))
				panic("unreachable")
			}
		l3:
			m.fn1(v5)
		}
	l1:
		{
			t7 := int32(load32(m.memory[uint32(v3+i32(-40)):]))
			v4 = t7
			if v4 == i32(-1) {
				goto l5
			}
			{
				if v4 == 0 {
					goto l6
				}
				t8 := int32(load32(m.memory[uint32(v3+i32(-36)):]))
				v5 = t8
				t9 := int32(load32(m.memory[uint32(v5+i32(-4)):]))
				v6 = t9
				v7 = v6 & i32(-8)
				t10 := v7
				v6 = v6 & i32(3)
				p11 := i32(8)
				if v6 != 0 {
					p11 = i32(4)
				}
				v4 = v4 << 1
				if uint32(t10) < uint32(p11+v4) {
					m.fn3(i32(1273840), i32(46), i32(1273888))
					panic("unreachable")
				}
				if v6 == 0 {
					goto l8
				}
				if uint32(v7) > uint32(v4+i32(39)) {
					m.fn3(i32(1273904), i32(46), i32(1273952))
					panic("unreachable")
				}
			l8:
				m.fn1(v5)
			}
		l6:
			t12 := int32(load32(m.memory[uint32(v3+i32(-28)):]))
			v4 = t12
			if v4 == 0 {
				goto l5
			}
			t13 := int32(load32(m.memory[uint32(v3+i32(-24)):]))
			v5 = t13
			t14 := int32(load32(m.memory[uint32(v5+i32(-4)):]))
			v6 = t14
			v7 = v6 & i32(-8)
			t15 := v7
			v6 = v6 & i32(3)
			p16 := i32(8)
			if v6 != 0 {
				p16 = i32(4)
			}
			v4 = v4 << 2
			if uint32(t15) < uint32(p16+v4) {
				m.fn3(i32(1273840), i32(46), i32(1273888))
				panic("unreachable")
			}
			if v6 == 0 {
				goto l11
			}
			if uint32(v7) > uint32(v4+i32(39)) {
				m.fn3(i32(1273904), i32(46), i32(1273952))
				panic("unreachable")
			}
		l11:
			m.fn1(v5)
		}
	l5:
		v3 = v3 + i32(72)
		v2 = v2 + i32(-1)
		if v2 != 0 {
			goto l13
		}
	}
l0:
	{
		t17 := int32(load32(m.memory[uint32(v0):]))
		v3 = t17
		if v3 == 0 {
			return
		}
		t18 := int32(load32(m.memory[uint32(v1+i32(-4)):]))
		v2 = t18
		v4 = v2 & i32(-8)
		t19 := v4
		v2 = v2 & i32(3)
		p20 := i32(8)
		if v2 != 0 {
			p20 = i32(4)
		}
		v3 = v3 * i32(72)
		if uint32(t19) < uint32(p20+v3) {
			m.fn3(i32(1273840), i32(46), i32(1273888))
			panic("unreachable")
		}
		if v2 == 0 {
			goto l16
		}
		if uint32(v4) > uint32(v3+i32(39)) {
			m.fn3(i32(1273904), i32(46), i32(1273952))
			panic("unreachable")
		}
	l16:
		m.fn1(v1)
	}
}
func (m *Module) fn119(v0, v1, v2, v3, v4, v5 int32) {
	var v6, v7, v8, v9, v10, v11, v12, v13, v14, v15, v16, v17, v18, v19, v20 int32
	t0 := m.g0
	v6 = t0 - i32(80)
	m.g0 = v6
	{
		if uint32(v1) >= uint32(i32(33)) {
			v8 = v2 + i32(-72)
		l18:
			{
				if v4 != 0 {
					goto l2
				}
				m.fn117(v0, v1, v2, v3, i32(1))
				goto l3
			l2:
				t1 := v0
				v9 = int32(uint32(v1) >> 3)
				v7 = t1 + v9*i32(504)
				v10 = v0 + v9*i32(288)
				{
					{
						if uint32(v1) < uint32(i32(64)) {
							goto l4
						}
						t2 := m.fn120(v0, v10, v7, v9)
						v11 = t2
						goto l5
					}
				l4:
					t3 := int32(load32(m.memory[uint32(v0+i32(64)):]))
					t4 := v0
					t5 := v7
					t6 := v10
					v9 = t3
					t7 := int32(load32(m.memory[uint32(v10+i32(64)):]))
					t8 := v9
					v12 = t7
					var p9 int32
					if uint32(t8) < uint32(v12) {
						p9 = 1
					}
					v13 = p9
					t10 := int32(load32(m.memory[uint32(v7+i32(64)):]))
					t11 := v13
					t12 := v12
					v11 = t10
					var p13 int32
					if uint32(t12) < uint32(v11) {
						p13 = 1
					}
					p14 := t6
					if t11^p13 != 0 {
						p14 = t5
					}
					t15 := v13
					var p16 int32
					if uint32(v9) < uint32(v11) {
						p16 = 1
					}
					p17 := p14
					if t15^p16 != 0 {
						p17 = t4
					}
					v11 = p17
				}
			l5:
				v4 = v4 + i32(-1)
				memory_copy(m.memory, uint32(v6+i32(8)), uint32(v11), uint32(i32(72)))
				t18 := int32(uint32(v11-v0) / uint32(i32(72)))
				v14 = t18
				{
					{
						if v5 == 0 {
							goto l6
						}
						t19 := int32(load32(m.memory[uint32(v5+i32(64)):]))
						t20 := int32(load32(m.memory[uint32(v11+i32(64)):]))
						if uint32(t19) >= uint32(t20) {
							goto l7
						}
					}
				l6:
					if uint32(v3) < uint32(v1) {
						goto l8
					}
					v7 = i32(0)
					v9 = v0
					t21 := v2
					v15 = v1 * i32(72)
					v16 = t21 + v15
					v10 = v16
					v17 = v14
				l12:
					{
						t22 := v9
						v13 = v0 + v17*i32(72)
						if uint32(t22) >= uint32(v13) {
							goto l9
						}
					l10:
						{
							t23 := v2
							v10 = v10 + i32(-72)
							t24 := int32(load32(m.memory[uint32(v9+i32(64)):]))
							t25 := int32(load32(m.memory[uint32(v11+i32(64)):]))
							t26 := v10
							var p27 int32
							if uint32(t24) < uint32(t25) {
								p27 = 1
							}
							v12 = p27
							p28 := t26
							if v12 != 0 {
								p28 = t23
							}
							memory_copy(m.memory, uint32(p28+v7*i32(72)), uint32(v9), uint32(i32(72)))
							v7 = v7 + v12
							v9 = v9 + i32(72)
							if uint32(v9) < uint32(v13) {
								goto l10
							}
						}
					}
				l9:
					if v17 == v1 {
						v17 = v7 * i32(72)
						if v17 == 0 {
							goto l13
						}
						memory_copy(m.memory, uint32(v0), uint32(v2), uint32(v17))
					l13:
						v18 = v1 - v7
						if v1 == v7 {
							goto l14
						}
						v19 = v0 + v17
						v10 = i32(0)
						if v1 == v7+i32(1) {
							goto l15
						}
						v20 = v18 & i32(1)
						v13 = v18 & i32(-2)
						v12 = v8 + v15
						v10 = i32(0)
						v9 = v19
					l16:
						{
							memory_copy(m.memory, uint32(v9), uint32(v12), uint32(i32(72)))
							memory_copy(m.memory, uint32(v9+i32(72)), uint32(v16+(v10^i32(0x1ffffffe))*i32(72)), uint32(i32(72)))
							v12 = v12 + i32(-144)
							v9 = v9 + i32(144)
							t29 := v13
							v10 = v10 + i32(2)
							if t29 != v10 {
								goto l16
							}
						}
						if v20 == 0 {
							goto l14
						}
					l15:
						memory_copy(m.memory, uint32(v19+v10*i32(72)), uint32(v16+(v10^i32(-1))*i32(72)), uint32(i32(72)))
					l14:
						if v7 == 0 {
							goto l7
						}
						if uint32(v1) < uint32(v7) {
							goto l17
						}
						m.fn119(v0+v17, v18, v2, v3, v4, v6+i32(8))
						v1 = v7
						if uint32(v7) < uint32(i32(33)) {
							goto l1
						}
						goto l18
					}
					v10 = v10 + i32(-72)
					memory_copy(m.memory, uint32(v10+v7*i32(72)), uint32(v9), uint32(i32(72)))
					v9 = v9 + i32(72)
					v17 = v1
					goto l12
				}
			l7:
				if uint32(v3) < uint32(v1) {
					goto l8
				}
				v10 = i32(0)
				v9 = v0
				t30 := v2
				v16 = v1 * i32(72)
				v17 = t30 + v16
				v7 = v17
			l22:
				{
					t31 := v9
					v13 = v0 + v14*i32(72)
					if uint32(t31) >= uint32(v13) {
						goto l19
					}
				l20:
					{
						t32 := v2
						v7 = v7 + i32(-72)
						t33 := int32(load32(m.memory[uint32(v11+i32(64)):]))
						t34 := int32(load32(m.memory[uint32(v9+i32(64)):]))
						t35 := v7
						var p36 int32
						if uint32(t33) >= uint32(t34) {
							p36 = 1
						}
						v12 = p36
						p37 := t35
						if v12 != 0 {
							p37 = t32
						}
						memory_copy(m.memory, uint32(p37+v10*i32(72)), uint32(v9), uint32(i32(72)))
						v10 = v10 + v12
						v9 = v9 + i32(72)
						if uint32(v9) < uint32(v13) {
							goto l20
						}
					}
				}
			l19:
				if v14 == v1 {
					goto l21
				}
				memory_copy(m.memory, uint32(v2+v10*i32(72)), uint32(v9), uint32(i32(72)))
				v9 = v9 + i32(72)
				v10 = v10 + i32(1)
				v7 = v7 + i32(-72)
				v14 = v1
				goto l22
			l21:
				v9 = v10 * i32(72)
				if v9 == 0 {
					goto l23
				}
				memory_copy(m.memory, uint32(v0), uint32(v2), uint32(v9))
			l23:
				if v1 == v10 {
					goto l3
				}
				v7 = v1 - v10
				v0 = v0 + v9
				v12 = i32(0)
				if v1 == v10+i32(1) {
					goto l24
				}
				v14 = v7 & i32(1)
				v11 = v7 & i32(-2)
				v13 = v8 + v16
				v12 = i32(0)
				v9 = v0
			l25:
				{
					memory_copy(m.memory, uint32(v9), uint32(v13), uint32(i32(72)))
					memory_copy(m.memory, uint32(v9+i32(72)), uint32(v17+(v12^i32(0x1ffffffe))*i32(72)), uint32(i32(72)))
					v13 = v13 + i32(-144)
					v9 = v9 + i32(144)
					t38 := v11
					v12 = v12 + i32(2)
					if t38 != v12 {
						goto l25
					}
				}
				if v14 == 0 {
					goto l26
				}
			l24:
				memory_copy(m.memory, uint32(v0+v12*i32(72)), uint32(v17+(v12^i32(-1))*i32(72)), uint32(i32(72)))
			l26:
				if uint32(v1) < uint32(v10) {
					goto l27
				}
				v5 = i32(0)
				v1 = v7
				if uint32(v7) < uint32(i32(33)) {
					goto l1
				}
				goto l18
			l27:
			}
			m.fn121(v10, v1, v1, i32(1069572))
			panic("unreachable")
		l17:
			m.fn28(i32(1271784), i32(19), i32(1069556))
		l8:
			panic("unreachable")
		}
		v7 = v1
		goto l1
	l1:
		if uint32(v7) < uint32(i32(2)) {
			goto l3
		}
		v14 = i32(1)
		v1 = int32(uint32(v7) >> 1)
		{
			{
				if uint32(v7) < uint32(i32(8)) {
					goto l28
				}
				t39 := int32(load32(m.memory[uint32(v0+i32(136)):]))
				t40 := v0
				v10 = t39
				t41 := int32(load32(m.memory[uint32(v0+i32(64)):]))
				t42 := v10
				v13 = t41
				var p43 int32
				if uint32(t42) < uint32(v13) {
					p43 = 1
				}
				v12 = t40 + p43*i32(72)
				t44 := int32(load32(m.memory[uint32(v0+i32(280)):]))
				t45 := int32(load32(m.memory[uint32(v0+i32(208)):]))
				t46 := v12
				t47 := v0
				var p48 int32
				if uint32(t44) < uint32(t45) {
					p48 = 1
				}
				v11 = p48
				p49 := i32(144)
				if v11 != 0 {
					p49 = i32(216)
				}
				v9 = t47 + p49
				t50 := v9
				t51 := v0
				var p52 int32
				if uint32(v10) >= uint32(v13) {
					p52 = 1
				}
				v10 = t51 + p52*i32(72)
				t54 := v10
				t55 := v0
				p53 := i32(216)
				if v11 != 0 {
					p53 = i32(144)
				}
				v13 = t55 + p53
				t56 := int32(load32(m.memory[uint32(v13+i32(64)):]))
				t57 := int32(load32(m.memory[uint32(v10+i32(64)):]))
				var p58 int32
				if uint32(t56) < uint32(t57) {
					p58 = 1
				}
				v11 = p58
				p59 := t54
				if v11 != 0 {
					p59 = t50
				}
				t60 := int32(load32(m.memory[uint32(v9+i32(64)):]))
				t61 := int32(load32(m.memory[uint32(v12+i32(64)):]))
				var p62 int32
				if uint32(t60) < uint32(t61) {
					p62 = 1
				}
				v17 = p62
				p63 := p59
				if v17 != 0 {
					p63 = t46
				}
				v16 = p63
				t64 := int32(load32(m.memory[uint32(v16+i32(64)):]))
				v4 = t64
				t66 := v13
				p65 := v9
				if v17 != 0 {
					p65 = v10
				}
				p67 := p65
				if v11 != 0 {
					p67 = t66
				}
				v14 = p67
				t68 := int32(load32(m.memory[uint32(v14+i32(64)):]))
				v18 = t68
				t70 := v2
				p69 := v12
				if v17 != 0 {
					p69 = v9
				}
				memory_copy(m.memory, uint32(t70), uint32(p69), uint32(i32(72)))
				t71 := v2 + i32(72)
				t72 := v14
				t73 := v16
				var p74 int32
				if uint32(v18) < uint32(v4) {
					p74 = 1
				}
				v9 = p74
				p75 := t73
				if v9 != 0 {
					p75 = t72
				}
				memory_copy(m.memory, uint32(t71), uint32(p75), uint32(i32(72)))
				t77 := v2 + i32(144)
				p76 := v14
				if v9 != 0 {
					p76 = v16
				}
				memory_copy(m.memory, uint32(t77), uint32(p76), uint32(i32(72)))
				t79 := v2 + i32(216)
				p78 := v13
				if v11 != 0 {
					p78 = v10
				}
				memory_copy(m.memory, uint32(t79), uint32(p78), uint32(i32(72)))
				t80 := v0
				v18 = v1 * i32(72)
				v9 = t80 + v18
				t81 := int32(load32(m.memory[uint32(v9+i32(136)):]))
				t82 := v9
				v12 = t81
				t83 := int32(load32(m.memory[uint32(v9+i32(64)):]))
				t84 := v12
				v11 = t83
				var p85 int32
				if uint32(t84) < uint32(v11) {
					p85 = 1
				}
				v13 = t82 + p85*i32(72)
				t86 := int32(load32(m.memory[uint32(v9+i32(280)):]))
				t87 := int32(load32(m.memory[uint32(v9+i32(208)):]))
				t88 := v13
				t89 := v9
				var p90 int32
				if uint32(t86) < uint32(t87) {
					p90 = 1
				}
				v17 = p90
				p91 := i32(144)
				if v17 != 0 {
					p91 = i32(216)
				}
				v10 = t89 + p91
				t92 := v10
				t93 := v9
				var p94 int32
				if uint32(v12) >= uint32(v11) {
					p94 = 1
				}
				v12 = t93 + p94*i32(72)
				t96 := v12
				t97 := v9
				p95 := i32(216)
				if v17 != 0 {
					p95 = i32(144)
				}
				v11 = t97 + p95
				t98 := int32(load32(m.memory[uint32(v11+i32(64)):]))
				t99 := int32(load32(m.memory[uint32(v12+i32(64)):]))
				var p100 int32
				if uint32(t98) < uint32(t99) {
					p100 = 1
				}
				v17 = p100
				p101 := t96
				if v17 != 0 {
					p101 = t92
				}
				t102 := int32(load32(m.memory[uint32(v10+i32(64)):]))
				t103 := int32(load32(m.memory[uint32(v13+i32(64)):]))
				var p104 int32
				if uint32(t102) < uint32(t103) {
					p104 = 1
				}
				v16 = p104
				p105 := p101
				if v16 != 0 {
					p105 = t88
				}
				v14 = p105
				t106 := int32(load32(m.memory[uint32(v14+i32(64)):]))
				v3 = t106
				t108 := v11
				p107 := v10
				if v16 != 0 {
					p107 = v12
				}
				p109 := p107
				if v17 != 0 {
					p109 = t108
				}
				v4 = p109
				t110 := int32(load32(m.memory[uint32(v4+i32(64)):]))
				v5 = t110
				v9 = v2 + v18
				t112 := v9
				p111 := v13
				if v16 != 0 {
					p111 = v10
				}
				memory_copy(m.memory, uint32(t112), uint32(p111), uint32(i32(72)))
				t113 := v9 + i32(72)
				t114 := v4
				t115 := v14
				var p116 int32
				if uint32(v5) < uint32(v3) {
					p116 = 1
				}
				v10 = p116
				p117 := t115
				if v10 != 0 {
					p117 = t114
				}
				memory_copy(m.memory, uint32(t113), uint32(p117), uint32(i32(72)))
				t119 := v9 + i32(144)
				p118 := v4
				if v10 != 0 {
					p118 = v14
				}
				memory_copy(m.memory, uint32(t119), uint32(p118), uint32(i32(72)))
				t121 := v9 + i32(216)
				p120 := v11
				if v17 != 0 {
					p120 = v12
				}
				memory_copy(m.memory, uint32(t121), uint32(p120), uint32(i32(72)))
				v14 = i32(4)
				goto l29
			}
		l28:
			memory_copy(m.memory, uint32(v2), uint32(v0), uint32(i32(72)))
			t122 := v2
			v9 = v1 * i32(72)
			memory_copy(m.memory, uint32(t122+v9), uint32(v0+v9), uint32(i32(72)))
		}
	l29:
		v3 = v7 - v1
		if uint32(v14) >= uint32(v1) {
			goto l30
		}
		v17 = v14 * i32(72)
		v11 = v14
	l35:
		{
			t123 := v2
			v10 = v11 * i32(72)
			v9 = t123 + v10
			t124 := v9
			v13 = v0 + v10
			memory_copy(m.memory, uint32(t124), uint32(v13), uint32(i32(72)))
			{
				t125 := int32(load32(m.memory[uint32(v9+i32(64)):]))
				v12 = t125
				t126 := int32(load32(m.memory[uint32(v9+i32(-8)):]))
				if uint32(v12) >= uint32(t126) {
					goto l31
				}
				t127 := int32(load32(m.memory[int64(uint32(v9))+68:]))
				v16 = t127
				v9 = v17
			l33:
				{
					v10 = v2 + v9
					memory_copy(m.memory, uint32(v10), uint32(v10+i32(-72)), uint32(i32(72)))
					if v9 == i32(72) {
						goto l32
					}
					v9 = v9 + i32(-72)
					t128 := int32(load32(m.memory[uint32(v10+i32(-80)):]))
					if uint32(v12) < uint32(t128) {
						goto l33
					}
				}
				v9 = v2 + v9
				goto l34
			l32:
				v9 = v2
			l34:
				t129 := int64(load64(m.memory[int64(uint32(v13))+56:]))
				store64(m.memory[int64(uint32(v9))+56:], uint64(t129))
				t130 := int64(load64(m.memory[int64(uint32(v13))+48:]))
				store64(m.memory[int64(uint32(v9))+48:], uint64(t130))
				t131 := int64(load64(m.memory[int64(uint32(v13))+40:]))
				store64(m.memory[int64(uint32(v9))+40:], uint64(t131))
				t132 := int64(load64(m.memory[int64(uint32(v13))+32:]))
				store64(m.memory[int64(uint32(v9))+32:], uint64(t132))
				t133 := int64(load64(m.memory[int64(uint32(v13))+24:]))
				store64(m.memory[int64(uint32(v9))+24:], uint64(t133))
				t134 := int64(load64(m.memory[int64(uint32(v13))+16:]))
				store64(m.memory[int64(uint32(v9))+16:], uint64(t134))
				t135 := int64(load64(m.memory[int64(uint32(v13))+8:]))
				store64(m.memory[int64(uint32(v9))+8:], uint64(t135))
				t136 := int64(load64(m.memory[uint32(v13):]))
				store64(m.memory[uint32(v9):], uint64(t136))
				store32(m.memory[uint32(v10+i32(-4)):], uint32(v16))
				store32(m.memory[uint32(v10+i32(-8)):], uint32(v12))
			}
		l31:
			v17 = v17 + i32(72)
			v11 = v11 + i32(1)
			if v11 != v1 {
				goto l35
			}
		}
	l30:
		t137 := v2
		v9 = v1 * i32(72)
		v11 = t137 + v9
		if uint32(v14) >= uint32(v3) {
			goto l36
		}
		v5 = v0 + v9
		v13 = v14 * i32(72)
		v4 = i32(72)
		v18 = v11
	l41:
		{
			t138 := v11
			v10 = v14 * i32(72)
			v9 = t138 + v10
			t139 := v9
			v16 = v5 + v10
			memory_copy(m.memory, uint32(t139), uint32(v16), uint32(i32(72)))
			{
				t140 := int32(load32(m.memory[uint32(v9+i32(64)):]))
				v17 = t140
				t141 := int32(load32(m.memory[uint32(v9+i32(-8)):]))
				if uint32(v17) >= uint32(t141) {
					goto l37
				}
				t142 := int32(load32(m.memory[int64(uint32(v9))+68:]))
				v15 = t142
				v12 = v4
				v10 = v18
			l39:
				{
					v9 = v10 + v13
					memory_copy(m.memory, uint32(v9), uint32(v9+i32(-72)), uint32(i32(72)))
					if v13 == v12 {
						goto l38
					}
					v12 = v12 + i32(72)
					v10 = v10 + i32(-72)
					t143 := int32(load32(m.memory[uint32(v9+i32(-80)):]))
					if uint32(v17) < uint32(t143) {
						goto l39
					}
				}
				v10 = v10 + v13
				goto l40
			l38:
				v10 = v11
			l40:
				t144 := int64(load64(m.memory[int64(uint32(v16))+56:]))
				store64(m.memory[int64(uint32(v10))+56:], uint64(t144))
				t145 := int64(load64(m.memory[int64(uint32(v16))+48:]))
				store64(m.memory[int64(uint32(v10))+48:], uint64(t145))
				t146 := int64(load64(m.memory[int64(uint32(v16))+40:]))
				store64(m.memory[int64(uint32(v10))+40:], uint64(t146))
				t147 := int64(load64(m.memory[int64(uint32(v16))+32:]))
				store64(m.memory[int64(uint32(v10))+32:], uint64(t147))
				t148 := int64(load64(m.memory[int64(uint32(v16))+24:]))
				store64(m.memory[int64(uint32(v10))+24:], uint64(t148))
				t149 := int64(load64(m.memory[int64(uint32(v16))+16:]))
				store64(m.memory[int64(uint32(v10))+16:], uint64(t149))
				t150 := int64(load64(m.memory[int64(uint32(v16))+8:]))
				store64(m.memory[int64(uint32(v10))+8:], uint64(t150))
				t151 := int64(load64(m.memory[uint32(v16):]))
				store64(m.memory[uint32(v10):], uint64(t151))
				store32(m.memory[uint32(v9+i32(-4)):], uint32(v15))
				store32(m.memory[uint32(v9+i32(-8)):], uint32(v17))
			}
		l37:
			v4 = v4 + i32(-72)
			v18 = v18 + i32(72)
			v14 = v14 + i32(1)
			if v14 != v3 {
				goto l41
			}
		}
	l36:
		v9 = v11 + i32(-72)
		t152 := v2
		v12 = v7*i32(72) + i32(-72)
		v10 = t152 + v12
		v12 = v0 + v12
	l42:
		{
			t153 := int32(load32(m.memory[uint32(v11+i32(64)):]))
			t154 := v0
			t155 := v11
			t156 := v2
			v13 = t153
			t157 := int32(load32(m.memory[uint32(v2+i32(64)):]))
			t158 := v13
			v17 = t157
			var p159 int32
			if uint32(t158) < uint32(v17) {
				p159 = 1
			}
			v16 = p159
			p160 := t156
			if v16 != 0 {
				p160 = t155
			}
			memory_copy(m.memory, uint32(t154), uint32(p160), uint32(i32(72)))
			t161 := int32(load32(m.memory[uint32(v10+i32(64)):]))
			t162 := v12
			t163 := v9
			t164 := v10
			v14 = t161
			t165 := int32(load32(m.memory[uint32(v9+i32(64)):]))
			t166 := v14
			v4 = t165
			var p167 int32
			if uint32(t166) < uint32(v4) {
				p167 = 1
			}
			v18 = p167
			p168 := t164
			if v18 != 0 {
				p168 = t163
			}
			memory_copy(m.memory, uint32(t162), uint32(p168), uint32(i32(72)))
			t169 := v2
			var p170 int32
			if uint32(v13) >= uint32(v17) {
				p170 = 1
			}
			v2 = t169 + p170*i32(72)
			v11 = v11 + v16*i32(72)
			t172 := v9
			p171 := i32(0)
			if v18 != 0 {
				p171 = i32(-72)
			}
			v9 = t172 + p171
			t174 := v10
			p173 := i32(0)
			if uint32(v14) >= uint32(v4) {
				p173 = i32(-72)
			}
			v10 = t174 + p173
			v12 = v12 + i32(-72)
			v0 = v0 + i32(72)
			v1 = v1 + i32(-1)
			if v1 != 0 {
				goto l42
			}
		}
		v9 = v9 + i32(72)
		{
			if v7&i32(1) == 0 {
				goto l43
			}
			t175 := v0
			t176 := v2
			t177 := v11
			var p178 int32
			if uint32(v2) < uint32(v9) {
				p178 = 1
			}
			v7 = p178
			p179 := t177
			if v7 != 0 {
				p179 = t176
			}
			memory_copy(m.memory, uint32(t175), uint32(p179), uint32(i32(72)))
			t180 := v11
			var p181 int32
			if uint32(v2) >= uint32(v9) {
				p181 = 1
			}
			v11 = t180 + p181*i32(72)
			v2 = v2 + v7*i32(72)
		}
	l43:
		if v2 != v9 {
			goto l44
		}
		if v11 == v10+i32(72) {
			goto l3
		}
	l44:
		m.fn122()
		panic("unreachable")
	}
l3:
	m.g0 = v6 + i32(80)
}
func (m *Module) fn120(v0, v1, v2, v3 int32) int32 {
	var v4, v5, v6 int32
	{
		if uint32(v3) < uint32(i32(8)) {
			goto l0
		}
		t0 := v0
		t1 := v0
		v3 = int32(uint32(v3) >> 3)
		v4 = v3 * i32(288)
		t2 := t1 + v4
		t3 := v0
		v5 = v3 * i32(504)
		t4 := m.fn120(t0, t2, t3+v5, v3)
		v0 = t4
		t5 := m.fn120(v1, v1+v4, v1+v5, v3)
		v1 = t5
		t6 := m.fn120(v2, v2+v4, v2+v5, v3)
		v2 = t6
	}
l0:
	t7 := int32(load32(m.memory[uint32(v0+i32(64)):]))
	t8 := v0
	t9 := v2
	t10 := v1
	v3 = t7
	t11 := int32(load32(m.memory[uint32(v1+i32(64)):]))
	t12 := v3
	v4 = t11
	var p13 int32
	if uint32(t12) < uint32(v4) {
		p13 = 1
	}
	v5 = p13
	t14 := int32(load32(m.memory[uint32(v2+i32(64)):]))
	t15 := v5
	t16 := v4
	v6 = t14
	var p17 int32
	if uint32(t16) < uint32(v6) {
		p17 = 1
	}
	p18 := t10
	if t15^p17 != 0 {
		p18 = t9
	}
	t19 := v5
	var p20 int32
	if uint32(v3) < uint32(v6) {
		p20 = 1
	}
	p21 := p18
	if t19^p20 != 0 {
		p21 = t8
	}
	return p21
}
func (m *Module) fn121(v0, v1, v2, v3 int32) {
	var v4 int32
	var v5 int64
	t0 := m.g0
	v4 = t0 - i32(32)
	m.g0 = v4
	if uint32(v0) > uint32(v2) {
		store32(m.memory[int64(uint32(v4))+8:], uint32(v0))
		store32(m.memory[int64(uint32(v4))+12:], uint32(v2))
		t1 := v4
		v5 = int64(uint32(i32(2))) << 32
		store64(m.memory[int64(uint32(t1))+24:], uint64(v5|int64(uint32(v4+i32(12)))))
		store64(m.memory[int64(uint32(v4))+16:], uint64(v5|int64(uint32(v4+i32(8)))))
		m.fn28(i32(1050466), v4+i32(16), v3)
		panic("unreachable")
	}
	if uint32(v1) > uint32(v2) {
		store32(m.memory[int64(uint32(v4))+8:], uint32(v1))
		store32(m.memory[int64(uint32(v4))+12:], uint32(v2))
		t2 := v4
		v5 = int64(uint32(i32(2))) << 32
		store64(m.memory[int64(uint32(t2))+24:], uint64(v5|int64(uint32(v4+i32(12)))))
		store64(m.memory[int64(uint32(v4))+16:], uint64(v5|int64(uint32(v4+i32(8)))))
		m.fn28(i32(1050523), v4+i32(16), v3)
		panic("unreachable")
	}
	v5 = int64(uint32(i32(2))) << 32
	if uint32(v0) <= uint32(v1) {
		store32(m.memory[int64(uint32(v4))+8:], uint32(v1))
		store32(m.memory[int64(uint32(v4))+12:], uint32(v2))
		store64(m.memory[int64(uint32(v4))+24:], uint64(v5|int64(uint32(v4+i32(12)))))
		store64(m.memory[int64(uint32(v4))+16:], uint64(v5|int64(uint32(v4+i32(8)))))
		m.fn28(i32(1050523), v4+i32(16), v3)
		panic("unreachable")
	}
	store32(m.memory[int64(uint32(v4))+8:], uint32(v0))
	store32(m.memory[int64(uint32(v4))+12:], uint32(v1))
	store64(m.memory[int64(uint32(v4))+24:], uint64(v5|int64(uint32(v4+i32(12)))))
	store64(m.memory[int64(uint32(v4))+16:], uint64(v5|int64(uint32(v4+i32(8)))))
	m.fn28(i32(1049723), v4+i32(16), v3)
	panic("unreachable")
}
func (m *Module) fn122() {
	m.fn28(i32(1119608), i32(153), i32(1119684))
	panic("unreachable")
}
func (m *Module) fn123(v0, v1 int32) {
	var v2, v3, v4, v5 int32
	t0 := m.g0
	v2 = t0 - i32(4096)
	m.g0 = v2
	{
		{
			p1 := i32(1000000)
			if uint32(v1) < uint32(i32(1000000)) {
				p1 = v1
			}
			v3 = p1
			t2 := v3
			v4 = v1 - int32(uint32(v1)>>1)
			p3 := v4
			if uint32(v3) > uint32(v4) {
				p3 = t2
			}
			v3 = p3
			if uint32(v3) < uint32(i32(513)) {
				goto l0
			}
			if uint32(v4) > uint32(i32(0x1fffffff)) {
				goto l1
			}
			v4 = v3 << 3
			if uint32(v4) >= uint32(i32(0x7ffffffd)) {
				goto l1
			}
			t4 := m.fn5(v4)
			v5 = t4
			if v5 == 0 {
				m.fn10(i32(4), v4)
				panic("unreachable")
			}
			t5 := v0
			t6 := v1
			t7 := v5
			t8 := v3
			var p9 int32
			if uint32(v1) < uint32(i32(65)) {
				p9 = 1
			}
			m.fn124(t5, t6, t7, t8, p9)
			t10 := int32(load32(m.memory[uint32(v5+i32(-4)):]))
			v1 = t10
			v3 = v1 & i32(-8)
			t11 := v3
			v1 = v1 & i32(3)
			p12 := i32(8)
			if v1 != 0 {
				p12 = i32(4)
			}
			if uint32(t11) < uint32(p12+v4) {
				m.fn3(i32(1273840), i32(46), i32(1273888))
				panic("unreachable")
			}
			if v1 == 0 {
				goto l4
			}
			if uint32(v3) > uint32(v4+i32(39)) {
				m.fn3(i32(1273904), i32(46), i32(1273952))
				panic("unreachable")
			}
		l4:
			m.fn1(v5)
			goto l6
		}
	l0:
		t13 := v0
		t14 := v1
		t15 := v2
		var p16 int32
		if uint32(v1) < uint32(i32(65)) {
			p16 = 1
		}
		m.fn124(t13, t14, t15, i32(512), p16)
	}
l6:
	m.g0 = v2 + i32(4096)
	return
l1:
	m.fn9()
	panic("unreachable")
}
func (m *Module) fn124(v0, v1, v2, v3, v4 int32) {
	var v5 int32
	var v6, v7 int64
	var v8, v9, v10, v11, v12, v13, v14, v15, v16, v17, v18, v19, v20, v21, v22, v23, v24, v25 int32
	t0 := m.g0
	v5 = t0 - i32(336)
	m.g0 = v5
	v6 = int64(uint32(v1))
	t1 := int64(uint64(i64(0x4000000000000000)) / uint64(v6))
	v7 = t1
	var p2 int32
	if v7*v6 != i64(0x4000000000000000) {
		p2 = 1
	}
	v6 = int64(uint32(p2))
	{
		if uint32(v1) < uint32(i32(4097)) {
			goto l0
		}
		v8 = int32(bits.LeadingZeros32(uint32(v1|i32(1)))) ^ i32(31)
		v8 = int32(uint32(v8)>>1) + v8&i32(1)
		v9 = int32(uint32(i32_shl(i32(1), v8)+i32_shr_u(v1, v8)) >> 1)
		goto l1
	l0:
		v8 = v1 - int32(uint32(v1)>>1)
		p3 := i32(64)
		if uint32(v8) < uint32(i32(64)) {
			p3 = v8
		}
		v9 = p3
	}
l1:
	v6 = v7 + v6
	v10 = v0 + i32(-8)
	v11 = v0 + i32(20)
	v8 = i32(1)
	v12 = i32(0)
	v13 = i32(0)
l33:
	v14 = i32(0)
	v15 = i32(1)
	{
		var p4 int32
		if uint32(v1) > uint32(v12) {
			p4 = 1
		}
		v16 = p4
		if v16 == 0 {
			goto l2
		}
		t5 := v0
		v17 = v12 << 3
		v18 = t5 + v17
		{
			v19 = v1 - v12
			if uint32(v19) < uint32(v9) {
				goto l3
			}
			{
				if uint32(v19) >= uint32(i32(2)) {
					goto l4
				}
				v20 = v19
				goto l5
			l4:
				{
					{
						t6 := int32(load32(m.memory[uint32(v18+i32(12)):]))
						v21 = t6
						t7 := int32(load32(m.memory[uint32(v18+i32(4)):]))
						var p8 int32
						if uint32(v21) < uint32(t7) {
							p8 = 1
						}
						v22 = p8
						if v22 != 0 {
							goto l6
						}
						v20 = i32(2)
						if v19 == i32(2) {
							goto l5
						}
						v23 = v11 + v17
						v20 = i32(2)
					l8:
						{
							t9 := int32(load32(m.memory[uint32(v23):]))
							v24 = t9
							if uint32(v24) < uint32(v21) {
								goto l7
							}
							v23 = v23 + i32(8)
							v21 = v24
							t10 := v19
							v20 = v20 + i32(1)
							if t10 != v20 {
								goto l8
							}
							goto l9
						}
					}
				l6:
					if v19 == i32(2) {
						goto l10
					}
					v23 = v11 + v17
					v20 = i32(2)
				l11:
					{
						t11 := int32(load32(m.memory[uint32(v23):]))
						v24 = t11
						if uint32(v24) >= uint32(v21) {
							goto l7
						}
						v23 = v23 + i32(8)
						v21 = v24
						t12 := v19
						v20 = v20 + i32(1)
						if t12 != v20 {
							goto l11
						}
					}
				l9:
					v20 = v19
				l7:
					if uint32(v20) < uint32(v9) {
						goto l3
					}
					if v22 == 0 {
						goto l5
					}
					v21 = int32(uint32(v20) >> 1)
					if v21 == 0 {
						goto l5
					}
					t13 := v18
					v23 = v20 << 3
					v14 = t13 + v23
					v19 = i32(0)
					if v21 == i32(1) {
						goto l12
					}
					v25 = v21 & i32(1)
					v15 = v21 & i32(0x7ffffffe)
					v21 = v0 + v23
					v19 = i32(0)
					v23 = v0
				l13:
					{
						v22 = v21 + v17 + i32(-8)
						t14 := int64(load64(m.memory[uint32(v22):]))
						v7 = t14
						t15 := v22
						v24 = v23 + v17
						t16 := int64(load64(m.memory[uint32(v24):]))
						store64(m.memory[uint32(t15):], uint64(t16))
						store64(m.memory[uint32(v24):], uint64(v7))
						v24 = v24 + i32(8)
						t17 := int64(load64(m.memory[uint32(v24):]))
						v7 = t17
						t18 := v24
						v22 = v14 + (v19^i32(0x1ffffffe))<<3
						t19 := int64(load64(m.memory[uint32(v22):]))
						store64(m.memory[uint32(t18):], uint64(t19))
						store64(m.memory[uint32(v22):], uint64(v7))
						v21 = v21 + i32(-16)
						v23 = v23 + i32(16)
						t20 := v15
						v19 = v19 + i32(2)
						if t20 != v19 {
							goto l13
						}
					}
					if v25 == 0 {
						goto l5
					}
					goto l12
				}
			l10:
				v14 = v18 + i32(16)
				v19 = i32(0)
				v20 = i32(2)
			l12:
				v21 = v18 + v19<<3
				t21 := int64(load64(m.memory[uint32(v21):]))
				v7 = t21
				t22 := v21
				v19 = v14 + (v19^i32(-1))<<3
				t23 := int64(load64(m.memory[uint32(v19):]))
				store64(m.memory[uint32(t22):], uint64(t23))
				store64(m.memory[uint32(v19):], uint64(v7))
			}
		l5:
			v15 = v20<<1 | i32(1)
			goto l14
		l3:
			{
				if v4 != 0 {
					goto l15
				}
				p24 := v9
				if uint32(v19) < uint32(v9) {
					p24 = v19
				}
				v15 = p24 << 1
				goto l14
			}
		l15:
			t26 := v18
			p25 := i32(32)
			if uint32(v19) < uint32(i32(32)) {
				p25 = v19
			}
			v20 = p25
			m.fn125(t26, v20, v2, v3, i32(0), i32(0))
			v15 = v20<<1 | i32(1)
		}
	l14:
		v14 = int32(int64(bits.LeadingZeros64(uint64(v6*int64(uint32(int32(uint32(v15)>>1)+v12<<1)) ^ (int64(uint32(v12-int32(uint32(v8)>>1)))+int64(uint32(v12)))*v6))))
	}
l2:
	{
		if uint32(v13) < uint32(i32(2)) {
			goto l16
		}
		t27 := v10
		v20 = v12 << 3
		v25 = t27 + v20
		v22 = v0 + v20
	l30:
		{
			t28 := v5 + i32(270)
			v21 = v13 + i32(-1)
			t29 := int32(m.memory[uint32(t28+v21)])
			if uint32(t29) < uint32(v14) {
				goto l16
			}
			{
				t30 := int32(load32(m.memory[uint32(v5+i32(4)+v21<<2):]))
				v23 = t30
				v13 = int32(uint32(v23) >> 1)
				t31 := v13
				v19 = int32(uint32(v8) >> 1)
				v17 = t31 + v19
				if uint32(v17) > uint32(v3) {
					goto l17
				}
				if (v23|v8)&i32(1) == 0 {
					v8 = v17 << 1
					goto l21
				}
			}
		l17:
			v20 = v0 + (v12-v17)<<3
			if v23&i32(1) == 0 {
				goto l19
			}
			goto l20
		l19:
			m.fn125(v20, v13, v2, v3, int32(bits.LeadingZeros32(uint32(v13|i32(1))))<<1^i32(62), i32(0))
		l20:
			if v8&i32(1) != 0 {
				goto l22
			}
			m.fn125(v20+v13<<3, v19, v2, v3, int32(bits.LeadingZeros32(uint32(v19|i32(1))))<<1^i32(62), i32(0))
		l22:
			{
				if v19 == 0 {
					goto l23
				}
				if v13 == 0 {
					goto l23
				}
				t32 := v3
				t33 := v19
				t34 := v13
				var p35 int32
				if uint32(v19) < uint32(v13) {
					p35 = 1
				}
				v23 = p35
				p36 := t34
				if v23 != 0 {
					p36 = t33
				}
				v19 = p36
				if uint32(t32) < uint32(v19) {
					goto l23
				}
				v8 = v20 + v13<<3
				{
					v13 = v19 << 3
					if v13 == 0 {
						goto l24
					}
					t38 := v2
					p37 := v20
					if v23 != 0 {
						p37 = v8
					}
					memory_copy(m.memory, uint32(t38), uint32(p37), uint32(v13))
				}
			l24:
				v13 = v2 + v13
				if v23 != 0 {
					goto l25
				}
				v19 = v2
			l27:
				{
					t39 := int32(load32(m.memory[uint32(v8+i32(4)):]))
					t40 := v20
					t41 := v8
					t42 := v19
					v23 = t39
					t43 := int32(load32(m.memory[uint32(v19+i32(4)):]))
					t44 := v23
					v24 = t43
					var p45 int32
					if uint32(t44) < uint32(v24) {
						p45 = 1
					}
					v18 = p45
					p46 := t42
					if v18 != 0 {
						p46 = t41
					}
					t47 := int64(load64(m.memory[uint32(p46):]))
					store64(m.memory[uint32(t40):], uint64(t47))
					v20 = v20 + i32(8)
					t48 := v19
					var p49 int32
					if uint32(v23) >= uint32(v24) {
						p49 = 1
					}
					v19 = t48 + p49<<3
					if v19 == v13 {
						goto l26
					}
					v8 = v8 + v18<<3
					if v8 != v22 {
						goto l27
					}
					goto l26
				}
			l25:
				v19 = v25
			l29:
				{
					t50 := v19
					v23 = v8 + i32(-8)
					t51 := v23
					v24 = v13 + i32(-8)
					t52 := int32(load32(m.memory[uint32(v13+i32(-4)):]))
					t53 := v24
					v18 = t52
					t54 := int32(load32(m.memory[uint32(v8+i32(-4)):]))
					t55 := v18
					v8 = t54
					var p56 int32
					if uint32(t55) < uint32(v8) {
						p56 = 1
					}
					v13 = p56
					p57 := t53
					if v13 != 0 {
						p57 = t51
					}
					t58 := int64(load64(m.memory[uint32(p57):]))
					store64(m.memory[uint32(t50):], uint64(t58))
					v13 = v24 + v13<<3
					t59 := v23
					var p60 int32
					if uint32(v18) >= uint32(v8) {
						p60 = 1
					}
					v8 = t59 + p60<<3
					if v8 == v20 {
						goto l28
					}
					v19 = v19 + i32(-8)
					if v13 != v2 {
						goto l29
					}
				}
			l28:
				v20 = v8
				v19 = v2
			l26:
				v8 = v13 - v19
				if v8 == 0 {
					goto l23
				}
				memory_copy(m.memory, uint32(v20), uint32(v19), uint32(v8))
			}
		l23:
			v8 = v17<<1 | i32(1)
		l21:
			v20 = i32(1)
			v13 = v21
			if uint32(v21) > uint32(i32(1)) {
				goto l30
			}
			goto l31
		}
	}
l16:
	v20 = v13
l31:
	m.memory[uint32(v5+i32(270)+v20)] = byte(v14)
	store32(m.memory[uint32(v5+i32(4)+v20<<2):], uint32(v8))
	if v16 == 0 {
		goto l32
	}
	v13 = v20 + i32(1)
	v12 = int32(uint32(v15)>>1) + v12
	v8 = v15
	goto l33
l32:
	if v8&i32(1) != 0 {
		goto l34
	}
	m.fn125(v0, v1, v2, v3, int32(bits.LeadingZeros32(uint32(v1|i32(1))))<<1^i32(62), i32(0))
l34:
	m.g0 = v5 + i32(336)
}
func (m *Module) fn125(v0, v1, v2, v3, v4, v5 int32) {
	var v6, v7, v8, v9, v10, v11, v12, v13, v14, v15, v16, v17, v18, v19, v20, v21, v22, v23 int32
	var v24 int64
	t0 := m.g0
	v6 = t0 - i32(16)
	m.g0 = v6
	{
		if uint32(v1) >= uint32(i32(33)) {
			v9 = v2 + i32(-8)
		l22:
			{
				if v4 != 0 {
					goto l2
				}
				m.fn124(v0, v1, v2, v3, i32(1))
				goto l3
			l2:
				t1 := v0
				v10 = int32(uint32(v1) >> 3)
				v8 = t1 + v10*i32(56)
				v11 = v0 + v10<<5
				{
					{
						if uint32(v1) < uint32(i32(64)) {
							goto l4
						}
						t2 := m.fn126(v0, v11, v8, v10)
						v12 = t2
						goto l5
					}
				l4:
					t3 := int32(load32(m.memory[uint32(v0+i32(4)):]))
					t4 := v0
					t5 := v8
					t6 := v11
					v10 = t3
					t7 := int32(load32(m.memory[uint32(v11+i32(4)):]))
					t8 := v10
					v13 = t7
					var p9 int32
					if uint32(t8) < uint32(v13) {
						p9 = 1
					}
					v14 = p9
					t10 := int32(load32(m.memory[uint32(v8+i32(4)):]))
					t11 := v14
					t12 := v13
					v15 = t10
					var p13 int32
					if uint32(t12) < uint32(v15) {
						p13 = 1
					}
					p14 := t6
					if t11^p13 != 0 {
						p14 = t5
					}
					t15 := v14
					var p16 int32
					if uint32(v10) < uint32(v15) {
						p16 = 1
					}
					p17 := p14
					if t15^p16 != 0 {
						p17 = t4
					}
					v12 = p17
				}
			l5:
				v4 = v4 + i32(-1)
				t18 := int32(load32(m.memory[int64(uint32(v12))+4:]))
				t19 := v6
				v10 = t18
				store32(m.memory[int64(uint32(t19))+12:], uint32(v10))
				t20 := int32(load32(m.memory[uint32(v12):]))
				store32(m.memory[int64(uint32(v6))+8:], uint32(t20))
				v16 = int32(uint32(v12-v0) >> 3)
				{
					{
						if v5 == 0 {
							goto l6
						}
						t21 := int32(load32(m.memory[uint32(v5+i32(4)):]))
						if uint32(t21) >= uint32(v10) {
							goto l7
						}
					}
				l6:
					if uint32(v3) < uint32(v1) {
						goto l8
					}
					v8 = i32(0)
					v7 = v0
					t22 := v2
					v17 = v1 << 3
					v18 = t22 + v17
					v14 = v18
					v19 = v16
				l15:
					{
						{
							t23 := v7
							t24 := v0
							v10 = v19 + i32(-3)
							p25 := v10
							if uint32(v10) > uint32(v19) {
								p25 = i32(0)
							}
							v20 = t24 + p25<<3
							if uint32(t23) < uint32(v20) {
								goto l9
							}
							v10 = v7
							goto l10
						}
					l9:
						t26 := int32(load32(m.memory[int64(uint32(v12))+4:]))
						v11 = t26
						v21 = i32(0)
						v15 = i32(0)
					l11:
						{
							t27 := v2
							v13 = v14 + v21
							t28 := v13 + i32(-8)
							v10 = v7 + v15
							t29 := int32(load32(m.memory[uint32(v10+i32(4)):]))
							var p30 int32
							if uint32(t29) < uint32(v11) {
								p30 = 1
							}
							v22 = p30
							p31 := t28
							if v22 != 0 {
								p31 = t27
							}
							t32 := int64(load64(m.memory[uint32(v10):]))
							store64(m.memory[uint32(p31+v8<<3):], uint64(t32))
							t33 := int32(load32(m.memory[uint32(v10+i32(12)):]))
							t34 := v2
							t35 := v13 + i32(-16)
							var p36 int32
							if uint32(t33) < uint32(v11) {
								p36 = 1
							}
							v23 = p36
							p37 := t35
							if v23 != 0 {
								p37 = t34
							}
							v8 = v8 + v22
							t38 := int64(load64(m.memory[uint32(v10+i32(8)):]))
							store64(m.memory[uint32(p37+v8<<3):], uint64(t38))
							t39 := int32(load32(m.memory[uint32(v10+i32(20)):]))
							t40 := v2
							t41 := v13 + i32(-24)
							var p42 int32
							if uint32(t39) < uint32(v11) {
								p42 = 1
							}
							v22 = p42
							p43 := t41
							if v22 != 0 {
								p43 = t40
							}
							v8 = v8 + v23
							t44 := int64(load64(m.memory[uint32(v10+i32(16)):]))
							store64(m.memory[uint32(p43+v8<<3):], uint64(t44))
							t45 := int32(load32(m.memory[uint32(v10+i32(28)):]))
							t46 := v2
							t47 := v13 + i32(-32)
							var p48 int32
							if uint32(t45) < uint32(v11) {
								p48 = 1
							}
							v13 = p48
							p49 := t47
							if v13 != 0 {
								p49 = t46
							}
							v8 = v8 + v22
							t50 := int64(load64(m.memory[uint32(v10+i32(24)):]))
							store64(m.memory[uint32(p49+v8<<3):], uint64(t50))
							v8 = v8 + v13
							v21 = v21 + i32(-32)
							t51 := v7
							v15 = v15 + i32(32)
							v10 = t51 + v15
							if uint32(v10) < uint32(v20) {
								goto l11
							}
						}
						v14 = v14 - v15
					}
				l10:
					{
						t52 := v10
						v13 = v0 + v19<<3
						if uint32(t52) >= uint32(v13) {
							goto l12
						}
						t53 := int32(load32(m.memory[int64(uint32(v12))+4:]))
						v15 = t53
					l13:
						{
							t54 := v2
							v14 = v14 + i32(-8)
							t55 := int32(load32(m.memory[uint32(v10+i32(4)):]))
							t56 := v14
							var p57 int32
							if uint32(t55) < uint32(v15) {
								p57 = 1
							}
							v11 = p57
							p58 := t56
							if v11 != 0 {
								p58 = t54
							}
							t59 := int64(load64(m.memory[uint32(v10):]))
							store64(m.memory[uint32(p58+v8<<3):], uint64(t59))
							v8 = v8 + v11
							v10 = v10 + i32(8)
							if uint32(v10) < uint32(v13) {
								goto l13
							}
						}
					}
				l12:
					{
						if v19 == v1 {
							v21 = v8 << 3
							if v21 == 0 {
								goto l16
							}
							memory_copy(m.memory, uint32(v0), uint32(v2), uint32(v21))
						l16:
							v7 = v1 - v8
							{
								if v1 == v8 {
									goto l17
								}
								v13 = v7 & i32(3)
								v11 = i32(0)
								if uint32(v8-v1) > uint32(i32(-4)) {
									goto l18
								}
								v10 = v0 + v21
								v15 = v7 & i32(-4)
								v14 = v9 + v17
								v11 = i32(0)
							l19:
								{
									t61 := int64(load64(m.memory[uint32(v14):]))
									store64(m.memory[uint32(v10):], uint64(t61))
									t62 := int64(load64(m.memory[uint32(v18+(v11^i32(0x1ffffffe))<<3):]))
									store64(m.memory[uint32(v10+i32(8)):], uint64(t62))
									t63 := int64(load64(m.memory[uint32(v18+(v11^i32(0x1ffffffd))<<3):]))
									store64(m.memory[uint32(v10+i32(16)):], uint64(t63))
									t64 := int64(load64(m.memory[uint32(v18+(v11^i32(0x1ffffffc))<<3):]))
									store64(m.memory[uint32(v10+i32(24)):], uint64(t64))
									v14 = v14 + i32(-32)
									v10 = v10 + i32(32)
									t65 := v15
									v11 = v11 + i32(4)
									if t65 != v11 {
										goto l19
									}
								}
								if v13 == 0 {
									goto l17
								}
							l18:
								t66 := v9
								t67 := v17
								v11 = v11 << 3
								v10 = t66 + (t67 - v11)
								v11 = v0 + v11 + v21
							l20:
								{
									t68 := int64(load64(m.memory[uint32(v10):]))
									store64(m.memory[uint32(v11):], uint64(t68))
									v10 = v10 + i32(-8)
									v11 = v11 + i32(8)
									v13 = v13 + i32(-1)
									if v13 != 0 {
										goto l20
									}
								}
							}
						l17:
							if v8 == 0 {
								goto l7
							}
							if uint32(v8) > uint32(v1) {
								goto l21
							}
							m.fn125(v0+v21, v7, v2, v3, v4, v6+i32(8))
							v1 = v8
							if uint32(v8) >= uint32(i32(33)) {
								goto l22
							}
							v7 = v0
							goto l1
						}
						v14 = v14 + i32(-8)
						t60 := int64(load64(m.memory[uint32(v10):]))
						store64(m.memory[uint32(v14+v8<<3):], uint64(t60))
						v7 = v10 + i32(8)
						v19 = v1
						goto l15
					}
				}
			l7:
				if uint32(v3) < uint32(v1) {
					goto l8
				}
				v11 = i32(0)
				v7 = v0
				t69 := v2
				v18 = v1 << 3
				v19 = t69 + v18
				v14 = v19
			l29:
				{
					{
						t70 := v7
						t71 := v0
						v10 = v16 + i32(-3)
						p72 := v10
						if uint32(v10) > uint32(v16) {
							p72 = i32(0)
						}
						v20 = t71 + p72<<3
						if uint32(t70) < uint32(v20) {
							goto l23
						}
						v10 = v7
						goto l24
					}
				l23:
					t73 := int32(load32(m.memory[int64(uint32(v12))+4:]))
					v8 = t73
					v21 = i32(0)
					v15 = i32(0)
				l25:
					{
						t74 := v2
						v13 = v14 + v21
						t75 := v13 + i32(-8)
						t76 := v8
						v10 = v7 + v15
						t77 := int32(load32(m.memory[uint32(v10+i32(4)):]))
						var p78 int32
						if uint32(t76) >= uint32(t77) {
							p78 = 1
						}
						v22 = p78
						p79 := t75
						if v22 != 0 {
							p79 = t74
						}
						t80 := int64(load64(m.memory[uint32(v10):]))
						store64(m.memory[uint32(p79+v11<<3):], uint64(t80))
						t81 := int32(load32(m.memory[uint32(v10+i32(12)):]))
						t82 := v2
						t83 := v13 + i32(-16)
						var p84 int32
						if uint32(v8) >= uint32(t81) {
							p84 = 1
						}
						v23 = p84
						p85 := t83
						if v23 != 0 {
							p85 = t82
						}
						v11 = v11 + v22
						t86 := int64(load64(m.memory[uint32(v10+i32(8)):]))
						store64(m.memory[uint32(p85+v11<<3):], uint64(t86))
						t87 := int32(load32(m.memory[uint32(v10+i32(20)):]))
						t88 := v2
						t89 := v13 + i32(-24)
						var p90 int32
						if uint32(v8) >= uint32(t87) {
							p90 = 1
						}
						v22 = p90
						p91 := t89
						if v22 != 0 {
							p91 = t88
						}
						v11 = v11 + v23
						t92 := int64(load64(m.memory[uint32(v10+i32(16)):]))
						store64(m.memory[uint32(p91+v11<<3):], uint64(t92))
						t93 := int32(load32(m.memory[uint32(v10+i32(28)):]))
						t94 := v2
						t95 := v13 + i32(-32)
						var p96 int32
						if uint32(v8) >= uint32(t93) {
							p96 = 1
						}
						v13 = p96
						p97 := t95
						if v13 != 0 {
							p97 = t94
						}
						v11 = v11 + v22
						t98 := int64(load64(m.memory[uint32(v10+i32(24)):]))
						store64(m.memory[uint32(p97+v11<<3):], uint64(t98))
						v11 = v11 + v13
						v21 = v21 + i32(-32)
						t99 := v7
						v15 = v15 + i32(32)
						v10 = t99 + v15
						if uint32(v10) < uint32(v20) {
							goto l25
						}
					}
					v14 = v14 - v15
				}
			l24:
				{
					t100 := v10
					v13 = v0 + v16<<3
					if uint32(t100) >= uint32(v13) {
						goto l26
					}
					t101 := int32(load32(m.memory[int64(uint32(v12))+4:]))
					v15 = t101
				l27:
					{
						t102 := v2
						v14 = v14 + i32(-8)
						t103 := int32(load32(m.memory[uint32(v10+i32(4)):]))
						t104 := v14
						var p105 int32
						if uint32(v15) >= uint32(t103) {
							p105 = 1
						}
						v8 = p105
						p106 := t104
						if v8 != 0 {
							p106 = t102
						}
						t107 := int64(load64(m.memory[uint32(v10):]))
						store64(m.memory[uint32(p106+v11<<3):], uint64(t107))
						v11 = v11 + v8
						v10 = v10 + i32(8)
						if uint32(v10) < uint32(v13) {
							goto l27
						}
					}
				}
			l26:
				{
					if v16 == v1 {
						goto l28
					}
					t108 := int64(load64(m.memory[uint32(v10):]))
					store64(m.memory[uint32(v2+v11<<3):], uint64(t108))
					v7 = v10 + i32(8)
					v11 = v11 + i32(1)
					v14 = v14 + i32(-8)
					v16 = v1
					goto l29
				}
			l28:
				v22 = v11 << 3
				if v22 == 0 {
					goto l30
				}
				memory_copy(m.memory, uint32(v0), uint32(v2), uint32(v22))
			l30:
				if v1 == v11 {
					goto l3
				}
				v8 = v1 - v11
				v14 = v8 & i32(3)
				v7 = v0 + v22
				v13 = i32(0)
				{
					if uint32(v11-v1) > uint32(i32(-4)) {
						goto l31
					}
					v21 = v8 & i32(-4)
					v15 = v9 + v18
					v13 = i32(0)
					v10 = v7
				l32:
					{
						t109 := int64(load64(m.memory[uint32(v15):]))
						store64(m.memory[uint32(v10):], uint64(t109))
						t110 := int64(load64(m.memory[uint32(v19+(v13^i32(0x1ffffffe))<<3):]))
						store64(m.memory[uint32(v10+i32(8)):], uint64(t110))
						t111 := int64(load64(m.memory[uint32(v19+(v13^i32(0x1ffffffd))<<3):]))
						store64(m.memory[uint32(v10+i32(16)):], uint64(t111))
						t112 := int64(load64(m.memory[uint32(v19+(v13^i32(0x1ffffffc))<<3):]))
						store64(m.memory[uint32(v10+i32(24)):], uint64(t112))
						v15 = v15 + i32(-32)
						v10 = v10 + i32(32)
						t113 := v21
						v13 = v13 + i32(4)
						if t113 != v13 {
							goto l32
						}
					}
					if v14 == 0 {
						goto l33
					}
				l31:
					t114 := v9
					v13 = v13 << 3
					v10 = t114 - v13 + v18
					v13 = v0 + v13 + v22
				l34:
					{
						t115 := int64(load64(m.memory[uint32(v10):]))
						store64(m.memory[uint32(v13):], uint64(t115))
						v10 = v10 + i32(-8)
						v13 = v13 + i32(8)
						v14 = v14 + i32(-1)
						if v14 != 0 {
							goto l34
						}
					}
				}
			l33:
				if uint32(v11) > uint32(v1) {
					goto l35
				}
				v5 = i32(0)
				v0 = v7
				v1 = v8
				if uint32(v8) < uint32(i32(33)) {
					goto l1
				}
				goto l22
			l35:
			}
			m.fn121(v11, v1, v1, i32(1069572))
			panic("unreachable")
		l21:
			m.fn28(i32(1271784), i32(19), i32(1069556))
		l8:
			panic("unreachable")
		}
		v7 = v0
		v8 = v1
		goto l1
	l1:
		if uint32(v8) < uint32(i32(2)) {
			goto l3
		}
		v22 = int32(uint32(v8) >> 1)
		{
			if uint32(v8) > uint32(i32(15)) {
				goto l36
			}
			{
				if uint32(v8) <= uint32(i32(7)) {
					t209 := int64(load64(m.memory[uint32(v7):]))
					store64(m.memory[uint32(v2):], uint64(t209))
					t210 := v2
					v10 = v22 << 3
					t211 := int64(load64(m.memory[uint32(v7+v10):]))
					store64(m.memory[uint32(t210+v10):], uint64(t211))
					v23 = i32(1)
					goto l38
				}
				v23 = i32(4)
				t116 := int32(load32(m.memory[uint32(v7+i32(28)):]))
				t117 := int32(load32(m.memory[uint32(v7+i32(20)):]))
				t118 := v2
				t119 := v7
				var p120 int32
				if uint32(t116) < uint32(t117) {
					p120 = 1
				}
				v15 = p120
				p121 := i32(16)
				if v15 != 0 {
					p121 = i32(24)
				}
				v10 = t119 + p121
				t122 := int32(load32(m.memory[uint32(v7+i32(12)):]))
				t123 := v10
				t124 := v7
				v11 = t122
				t125 := int32(load32(m.memory[uint32(v7+i32(4)):]))
				t126 := v11
				v21 = t125
				var p127 int32
				if uint32(t126) < uint32(v21) {
					p127 = 1
				}
				v13 = t124 + p127<<3
				t128 := int32(load32(m.memory[uint32(v10+i32(4)):]))
				t129 := int32(load32(m.memory[uint32(v13+i32(4)):]))
				t130 := v13
				var p131 int32
				if uint32(t128) < uint32(t129) {
					p131 = 1
				}
				v14 = p131
				p132 := t130
				if v14 != 0 {
					p132 = t123
				}
				t133 := int64(load64(m.memory[uint32(p132):]))
				store64(m.memory[uint32(t118):], uint64(t133))
				t134 := v2
				t135 := v7
				var p136 int32
				if uint32(v11) >= uint32(v21) {
					p136 = 1
				}
				v11 = t135 + p136<<3
				t138 := v11
				t139 := v7
				p137 := i32(24)
				if v15 != 0 {
					p137 = i32(16)
				}
				v15 = t139 + p137
				t140 := int32(load32(m.memory[uint32(v15+i32(4)):]))
				t141 := int32(load32(m.memory[uint32(v11+i32(4)):]))
				t142 := v15
				var p143 int32
				if uint32(t140) < uint32(t141) {
					p143 = 1
				}
				v21 = p143
				p144 := t142
				if v21 != 0 {
					p144 = t138
				}
				t145 := int64(load64(m.memory[uint32(p144):]))
				store64(m.memory[int64(uint32(t134))+24:], uint64(t145))
				t147 := v2
				t148 := v15
				p146 := v10
				if v14 != 0 {
					p146 = v11
				}
				p149 := p146
				if v21 != 0 {
					p149 = t148
				}
				v15 = p149
				t151 := v15
				t152 := v13
				p150 := v11
				if v21 != 0 {
					p150 = v10
				}
				p153 := p150
				if v14 != 0 {
					p153 = t152
				}
				v10 = p153
				t154 := int32(load32(m.memory[uint32(v15+i32(4)):]))
				t155 := int32(load32(m.memory[uint32(v10+i32(4)):]))
				t156 := v10
				var p157 int32
				if uint32(t154) < uint32(t155) {
					p157 = 1
				}
				v11 = p157
				p158 := t156
				if v11 != 0 {
					p158 = t151
				}
				t159 := int64(load64(m.memory[uint32(p158):]))
				store64(m.memory[int64(uint32(t147))+8:], uint64(t159))
				t161 := v2
				p160 := v15
				if v11 != 0 {
					p160 = v10
				}
				t162 := int64(load64(m.memory[uint32(p160):]))
				store64(m.memory[int64(uint32(t161))+16:], uint64(t162))
				t163 := v2
				v10 = v22 << 3
				v11 = t163 + v10
				t164 := v11
				v10 = v7 + v10
				t165 := int32(load32(m.memory[uint32(v10+i32(28)):]))
				t166 := int32(load32(m.memory[uint32(v10+i32(20)):]))
				t167 := v10
				var p168 int32
				if uint32(t165) < uint32(t166) {
					p168 = 1
				}
				v21 = p168
				p169 := i32(16)
				if v21 != 0 {
					p169 = i32(24)
				}
				v13 = t167 + p169
				t170 := int32(load32(m.memory[uint32(v10+i32(12)):]))
				t171 := v13
				t172 := v10
				v20 = t170
				t173 := int32(load32(m.memory[uint32(v10+i32(4)):]))
				t174 := v20
				v19 = t173
				var p175 int32
				if uint32(t174) < uint32(v19) {
					p175 = 1
				}
				v14 = t172 + p175<<3
				t176 := int32(load32(m.memory[uint32(v13+i32(4)):]))
				t177 := int32(load32(m.memory[uint32(v14+i32(4)):]))
				t178 := v14
				var p179 int32
				if uint32(t176) < uint32(t177) {
					p179 = 1
				}
				v15 = p179
				p180 := t178
				if v15 != 0 {
					p180 = t171
				}
				t181 := int64(load64(m.memory[uint32(p180):]))
				store64(m.memory[uint32(t164):], uint64(t181))
				t183 := v11
				t184 := v10
				p182 := i32(24)
				if v21 != 0 {
					p182 = i32(16)
				}
				v21 = t184 + p182
				t185 := v21
				t186 := v10
				var p187 int32
				if uint32(v20) >= uint32(v19) {
					p187 = 1
				}
				v10 = t186 + p187<<3
				p188 := v13
				if v15 != 0 {
					p188 = v10
				}
				t189 := int32(load32(m.memory[uint32(v21+i32(4)):]))
				t190 := int32(load32(m.memory[uint32(v10+i32(4)):]))
				var p191 int32
				if uint32(t189) < uint32(t190) {
					p191 = 1
				}
				v20 = p191
				p192 := p188
				if v20 != 0 {
					p192 = t185
				}
				v19 = p192
				t194 := v19
				t195 := v14
				p193 := v10
				if v20 != 0 {
					p193 = v13
				}
				p196 := p193
				if v15 != 0 {
					p196 = t195
				}
				v13 = p196
				t197 := int32(load32(m.memory[uint32(v19+i32(4)):]))
				t198 := int32(load32(m.memory[uint32(v13+i32(4)):]))
				t199 := v13
				var p200 int32
				if uint32(t197) < uint32(t198) {
					p200 = 1
				}
				v14 = p200
				p201 := t199
				if v14 != 0 {
					p201 = t194
				}
				t202 := int64(load64(m.memory[uint32(p201):]))
				store64(m.memory[int64(uint32(t183))+8:], uint64(t202))
				t204 := v11
				p203 := v19
				if v14 != 0 {
					p203 = v13
				}
				t205 := int64(load64(m.memory[uint32(p203):]))
				store64(m.memory[int64(uint32(t204))+16:], uint64(t205))
				t207 := v11
				p206 := v21
				if v20 != 0 {
					p206 = v10
				}
				t208 := int64(load64(m.memory[uint32(p206):]))
				store64(m.memory[int64(uint32(t207))+24:], uint64(t208))
				goto l38
			}
		l36:
			t212 := v7
			t213 := v2
			v10 = v2 + v8<<3
			m.fn127(t212, t213, v10)
			t214 := v7
			v11 = v22 << 3
			m.fn127(t214+v11, v2+v11, v10+i32(64))
			v23 = i32(8)
		}
	l38:
		v0 = v8 - v22
		if uint32(v23) >= uint32(v22) {
			goto l39
		}
		v15 = v23 << 3
		v14 = v23
	l44:
		{
			t215 := v2
			v10 = v14 << 3
			v11 = t215 + v10
			t216 := int64(load64(m.memory[uint32(v7+v10):]))
			t217 := v11
			v24 = t216
			store64(m.memory[uint32(t217):], uint64(v24))
			{
				t218 := int32(load32(m.memory[uint32(v11+i32(-4)):]))
				v13 = int32(int64(uint64(v24) >> 32))
				if uint32(t218) <= uint32(v13) {
					goto l40
				}
				v21 = int32(v24)
				v10 = v15
			l42:
				{
					v11 = v2 + v10
					t219 := int64(load64(m.memory[uint32(v11+i32(-8)):]))
					store64(m.memory[uint32(v11):], uint64(t219))
					if v10 == i32(8) {
						goto l41
					}
					v10 = v10 + i32(-8)
					t220 := int32(load32(m.memory[uint32(v11+i32(-12)):]))
					if uint32(t220) > uint32(v13) {
						goto l42
					}
				}
				v10 = v2 + v10
				goto l43
			l41:
				v10 = v2
			l43:
				store32(m.memory[uint32(v10):], uint32(v21))
				store32(m.memory[uint32(v11+i32(-4)):], uint32(v13))
			}
		l40:
			v15 = v15 + i32(8)
			v14 = v14 + i32(1)
			if v14 != v22 {
				goto l44
			}
		}
	l39:
		t221 := v2
		v10 = v22 << 3
		v15 = t221 + v10
		if uint32(v23) >= uint32(v0) {
			goto l45
		}
		v1 = v7 + v10
		v14 = v23 << 3
		v20 = i32(8)
		v19 = v15
	l50:
		{
			t222 := v15
			v10 = v23 << 3
			v11 = t222 + v10
			t223 := int64(load64(m.memory[uint32(v1+v10):]))
			t224 := v11
			v24 = t223
			store64(m.memory[uint32(t224):], uint64(v24))
			{
				t225 := int32(load32(m.memory[uint32(v11+i32(-4)):]))
				v21 = int32(int64(uint64(v24) >> 32))
				if uint32(t225) <= uint32(v21) {
					goto l46
				}
				v12 = int32(v24)
				v13 = v20
				v11 = v19
			l48:
				{
					v10 = v11 + v14
					t226 := int64(load64(m.memory[uint32(v10+i32(-8)):]))
					store64(m.memory[uint32(v10):], uint64(t226))
					if v14 == v13 {
						goto l47
					}
					v13 = v13 + i32(8)
					v11 = v11 + i32(-8)
					t227 := int32(load32(m.memory[uint32(v10+i32(-12)):]))
					if uint32(t227) > uint32(v21) {
						goto l48
					}
				}
				v11 = v11 + v14
				goto l49
			l47:
				v11 = v15
			l49:
				store32(m.memory[uint32(v11):], uint32(v12))
				store32(m.memory[uint32(v10+i32(-4)):], uint32(v21))
			}
		l46:
			v20 = v20 + i32(-8)
			v19 = v19 + i32(8)
			v23 = v23 + i32(1)
			if v23 != v0 {
				goto l50
			}
		}
	l45:
		v10 = v15 + i32(-8)
		t228 := v2
		v13 = v8<<3 + i32(-8)
		v11 = t228 + v13
		v13 = v7 + v13
	l51:
		{
			t229 := int32(load32(m.memory[uint32(v15+i32(4)):]))
			t230 := v7
			t231 := v15
			t232 := v2
			v14 = t229
			t233 := int32(load32(m.memory[uint32(v2+i32(4)):]))
			t234 := v14
			v21 = t233
			var p235 int32
			if uint32(t234) < uint32(v21) {
				p235 = 1
			}
			v23 = p235
			p236 := t232
			if v23 != 0 {
				p236 = t231
			}
			t237 := int64(load64(m.memory[uint32(p236):]))
			store64(m.memory[uint32(t230):], uint64(t237))
			t238 := int32(load32(m.memory[uint32(v11+i32(4)):]))
			t239 := v13
			t240 := v10
			t241 := v11
			v20 = t238
			t242 := int32(load32(m.memory[uint32(v10+i32(4)):]))
			t243 := v20
			v19 = t242
			var p244 int32
			if uint32(t243) < uint32(v19) {
				p244 = 1
			}
			v0 = p244
			p245 := t241
			if v0 != 0 {
				p245 = t240
			}
			t246 := int64(load64(m.memory[uint32(p245):]))
			store64(m.memory[uint32(t239):], uint64(t246))
			t248 := v10
			p247 := i32(0)
			if v0 != 0 {
				p247 = i32(-8)
			}
			v10 = t248 + p247
			t250 := v11
			p249 := i32(0)
			if uint32(v20) >= uint32(v19) {
				p249 = i32(-8)
			}
			v11 = t250 + p249
			t251 := v2
			var p252 int32
			if uint32(v14) >= uint32(v21) {
				p252 = 1
			}
			v2 = t251 + p252<<3
			v15 = v15 + v23<<3
			v13 = v13 + i32(-8)
			v7 = v7 + i32(8)
			v22 = v22 + i32(-1)
			if v22 != 0 {
				goto l51
			}
		}
		v10 = v10 + i32(8)
		{
			if v8&i32(1) == 0 {
				goto l52
			}
			t253 := v7
			t254 := v2
			t255 := v15
			var p256 int32
			if uint32(v2) < uint32(v10) {
				p256 = 1
			}
			v8 = p256
			p257 := t255
			if v8 != 0 {
				p257 = t254
			}
			t258 := int64(load64(m.memory[uint32(p257):]))
			store64(m.memory[uint32(t253):], uint64(t258))
			t259 := v15
			var p260 int32
			if uint32(v2) >= uint32(v10) {
				p260 = 1
			}
			v15 = t259 + p260<<3
			v2 = v2 + v8<<3
		}
	l52:
		if v2 != v10 {
			goto l53
		}
		if v15 == v11+i32(8) {
			goto l3
		}
	l53:
		m.fn122()
		panic("unreachable")
	}
l3:
	m.g0 = v6 + i32(16)
}
func (m *Module) fn126(v0, v1, v2, v3 int32) int32 {
	var v4, v5, v6 int32
	{
		if uint32(v3) < uint32(i32(8)) {
			goto l0
		}
		t0 := v0
		t1 := v0
		v3 = int32(uint32(v3) >> 3)
		v4 = v3 << 5
		t2 := t1 + v4
		t3 := v0
		v5 = v3 * i32(56)
		t4 := m.fn126(t0, t2, t3+v5, v3)
		v0 = t4
		t5 := m.fn126(v1, v1+v4, v1+v5, v3)
		v1 = t5
		t6 := m.fn126(v2, v2+v4, v2+v5, v3)
		v2 = t6
	}
l0:
	t7 := int32(load32(m.memory[uint32(v0+i32(4)):]))
	t8 := v0
	t9 := v2
	t10 := v1
	v3 = t7
	t11 := int32(load32(m.memory[uint32(v1+i32(4)):]))
	t12 := v3
	v4 = t11
	var p13 int32
	if uint32(t12) < uint32(v4) {
		p13 = 1
	}
	v5 = p13
	t14 := int32(load32(m.memory[uint32(v2+i32(4)):]))
	t15 := v5
	t16 := v4
	v6 = t14
	var p17 int32
	if uint32(t16) < uint32(v6) {
		p17 = 1
	}
	p18 := t10
	if t15^p17 != 0 {
		p18 = t9
	}
	t19 := v5
	var p20 int32
	if uint32(v3) < uint32(v6) {
		p20 = 1
	}
	p21 := p18
	if t19^p20 != 0 {
		p21 = t8
	}
	return p21
}
func (m *Module) fn127(v0, v1, v2 int32) {
	var v3, v4, v5, v6, v7, v8, v9, v10, v11, v12 int32
	var v13, v14 int64
	var v15 int32
	var v16, v17 int64
	t0 := int32(load32(m.memory[uint32(v0+i32(12)):]))
	t1 := v0
	v3 = t0
	t2 := int32(load32(m.memory[uint32(v0+i32(4)):]))
	t3 := v3
	v4 = t2
	var p4 int32
	if uint32(t3) < uint32(v4) {
		p4 = 1
	}
	v5 = t1 + p4<<3
	t5 := int32(load32(m.memory[uint32(v0+i32(28)):]))
	t6 := int32(load32(m.memory[uint32(v0+i32(20)):]))
	t7 := v5
	t8 := v0
	var p9 int32
	if uint32(t5) < uint32(t6) {
		p9 = 1
	}
	v6 = p9
	p10 := i32(16)
	if v6 != 0 {
		p10 = i32(24)
	}
	v7 = t8 + p10
	t11 := v7
	t12 := v0
	var p13 int32
	if uint32(v3) >= uint32(v4) {
		p13 = 1
	}
	v3 = t12 + p13<<3
	t15 := v3
	t16 := v0
	p14 := i32(24)
	if v6 != 0 {
		p14 = i32(16)
	}
	v4 = t16 + p14
	t17 := int32(load32(m.memory[uint32(v4+i32(4)):]))
	t18 := int32(load32(m.memory[uint32(v3+i32(4)):]))
	var p19 int32
	if uint32(t17) < uint32(t18) {
		p19 = 1
	}
	v6 = p19
	p20 := t15
	if v6 != 0 {
		p20 = t11
	}
	t21 := int32(load32(m.memory[uint32(v7+i32(4)):]))
	t22 := int32(load32(m.memory[uint32(v5+i32(4)):]))
	var p23 int32
	if uint32(t21) < uint32(t22) {
		p23 = 1
	}
	v8 = p23
	p24 := p20
	if v8 != 0 {
		p24 = t7
	}
	v9 = p24
	t25 := int32(load32(m.memory[uint32(v9+i32(4)):]))
	v10 = t25
	t27 := v4
	p26 := v7
	if v8 != 0 {
		p26 = v3
	}
	p28 := p26
	if v6 != 0 {
		p28 = t27
	}
	v11 = p28
	t29 := int32(load32(m.memory[uint32(v11+i32(4)):]))
	v12 = t29
	t31 := v2
	p30 := v5
	if v8 != 0 {
		p30 = v7
	}
	t32 := int64(load64(m.memory[uint32(p30):]))
	v13 = t32
	store64(m.memory[uint32(t31):], uint64(v13))
	t33 := v2
	t34 := v11
	t35 := v9
	var p36 int32
	if uint32(v12) < uint32(v10) {
		p36 = 1
	}
	v7 = p36
	p37 := t35
	if v7 != 0 {
		p37 = t34
	}
	t38 := int64(load64(m.memory[uint32(p37):]))
	store64(m.memory[int64(uint32(t33))+8:], uint64(t38))
	t40 := v2
	p39 := v11
	if v7 != 0 {
		p39 = v9
	}
	t41 := int64(load64(m.memory[uint32(p39):]))
	store64(m.memory[int64(uint32(t40))+16:], uint64(t41))
	v10 = v2 + i32(24)
	t43 := v10
	p42 := v4
	if v6 != 0 {
		p42 = v3
	}
	t44 := int64(load64(m.memory[uint32(p42):]))
	v14 = t44
	store64(m.memory[uint32(t43):], uint64(v14))
	v7 = v0 + i32(32)
	t45 := int32(load32(m.memory[uint32(v0+i32(44)):]))
	t46 := v7
	v3 = t45
	t47 := int32(load32(m.memory[uint32(v0+i32(36)):]))
	t48 := v3
	v4 = t47
	var p49 int32
	if uint32(t48) < uint32(v4) {
		p49 = 1
	}
	v5 = t46 + p49<<3
	t50 := int32(load32(m.memory[uint32(v0+i32(60)):]))
	t51 := int32(load32(m.memory[uint32(v0+i32(52)):]))
	t52 := v5
	t53 := v7
	var p54 int32
	if uint32(t50) < uint32(t51) {
		p54 = 1
	}
	v6 = p54
	p55 := i32(16)
	if v6 != 0 {
		p55 = i32(24)
	}
	v0 = t53 + p55
	t56 := v0
	t57 := v7
	var p58 int32
	if uint32(v3) >= uint32(v4) {
		p58 = 1
	}
	v3 = t57 + p58<<3
	t60 := v3
	t61 := v7
	p59 := i32(24)
	if v6 != 0 {
		p59 = i32(16)
	}
	v7 = t61 + p59
	t62 := int32(load32(m.memory[uint32(v7+i32(4)):]))
	t63 := int32(load32(m.memory[uint32(v3+i32(4)):]))
	var p64 int32
	if uint32(t62) < uint32(t63) {
		p64 = 1
	}
	v4 = p64
	p65 := t60
	if v4 != 0 {
		p65 = t56
	}
	t66 := int32(load32(m.memory[uint32(v0+i32(4)):]))
	t67 := int32(load32(m.memory[uint32(v5+i32(4)):]))
	var p68 int32
	if uint32(t66) < uint32(t67) {
		p68 = 1
	}
	v6 = p68
	p69 := p65
	if v6 != 0 {
		p69 = t52
	}
	v8 = p69
	t70 := int32(load32(m.memory[uint32(v8+i32(4)):]))
	v12 = t70
	t72 := v7
	p71 := v0
	if v6 != 0 {
		p71 = v3
	}
	p73 := p71
	if v4 != 0 {
		p73 = t72
	}
	v9 = p73
	t74 := int32(load32(m.memory[uint32(v9+i32(4)):]))
	v15 = t74
	v11 = v2 + i32(32)
	t76 := v11
	p75 := v5
	if v6 != 0 {
		p75 = v0
	}
	t77 := int64(load64(m.memory[uint32(p75):]))
	v16 = t77
	store64(m.memory[uint32(t76):], uint64(v16))
	t78 := v2 + i32(40)
	t79 := v9
	t80 := v8
	var p81 int32
	if uint32(v15) < uint32(v12) {
		p81 = 1
	}
	v0 = p81
	p82 := t80
	if v0 != 0 {
		p82 = t79
	}
	t83 := int64(load64(m.memory[uint32(p82):]))
	store64(m.memory[uint32(t78):], uint64(t83))
	t85 := v2 + i32(48)
	p84 := v9
	if v0 != 0 {
		p84 = v8
	}
	t86 := int64(load64(m.memory[uint32(p84):]))
	store64(m.memory[uint32(t85):], uint64(t86))
	v5 = v2 + i32(56)
	t88 := v5
	p87 := v7
	if v4 != 0 {
		p87 = v3
	}
	t89 := int64(load64(m.memory[uint32(p87):]))
	v17 = t89
	store64(m.memory[uint32(t88):], uint64(v17))
	t90 := v1
	t91 := v11
	t92 := v2
	v16 = int64(uint64(v16) >> 32)
	t93 := v16
	v13 = int64(uint64(v13) >> 32)
	var p94 int32
	if uint64(t93) < uint64(v13) {
		p94 = 1
	}
	v0 = p94
	p95 := t92
	if v0 != 0 {
		p95 = t91
	}
	t96 := int64(load64(m.memory[uint32(p95):]))
	store64(m.memory[uint32(t90):], uint64(t96))
	t97 := v1
	t98 := v2
	v17 = int64(uint64(v17) >> 32)
	t99 := v17
	v14 = int64(uint64(v14) >> 32)
	var p100 int32
	if uint64(t99) < uint64(v14) {
		p100 = 1
	}
	v7 = p100
	p101 := i32(56)
	if v7 != 0 {
		p101 = i32(24)
	}
	t102 := int64(load64(m.memory[uint32(t98+p101):]))
	store64(m.memory[int64(uint32(t97))+56:], uint64(t102))
	t103 := v1
	v0 = v11 + v0<<3
	t104 := v0
	t105 := v2
	var p106 int32
	if uint64(v16) >= uint64(v13) {
		p106 = 1
	}
	v2 = t105 + p106<<3
	t107 := int32(load32(m.memory[uint32(v0+i32(4)):]))
	t108 := v2
	v4 = t107
	t109 := int32(load32(m.memory[uint32(v2+i32(4)):]))
	t110 := v4
	v6 = t109
	var p111 int32
	if uint32(t110) < uint32(v6) {
		p111 = 1
	}
	v8 = p111
	p112 := t108
	if v8 != 0 {
		p112 = t104
	}
	t113 := int64(load64(m.memory[uint32(p112):]))
	store64(m.memory[int64(uint32(t103))+8:], uint64(t113))
	t115 := v1
	t116 := v10
	p114 := i32(0)
	if v7 != 0 {
		p114 = i32(-8)
	}
	v7 = t116 + p114
	t118 := v7
	t119 := v5
	p117 := i32(0)
	if uint64(v17) >= uint64(v14) {
		p117 = i32(-8)
	}
	v3 = t119 + p117
	t120 := int32(load32(m.memory[uint32(v3+i32(4)):]))
	t121 := v3
	v5 = t120
	t122 := int32(load32(m.memory[uint32(v7+i32(4)):]))
	t123 := v5
	v9 = t122
	var p124 int32
	if uint32(t123) < uint32(v9) {
		p124 = 1
	}
	v11 = p124
	p125 := t121
	if v11 != 0 {
		p125 = t118
	}
	t126 := int64(load64(m.memory[uint32(p125):]))
	store64(m.memory[int64(uint32(t115))+48:], uint64(t126))
	t127 := v1
	v0 = v0 + v8<<3
	t128 := v0
	t129 := v2
	var p130 int32
	if uint32(v4) >= uint32(v6) {
		p130 = 1
	}
	v2 = t129 + p130<<3
	t131 := int32(load32(m.memory[uint32(v0+i32(4)):]))
	t132 := v2
	v4 = t131
	t133 := int32(load32(m.memory[uint32(v2+i32(4)):]))
	t134 := v4
	v6 = t133
	var p135 int32
	if uint32(t134) < uint32(v6) {
		p135 = 1
	}
	v8 = p135
	p136 := t132
	if v8 != 0 {
		p136 = t128
	}
	t137 := int64(load64(m.memory[uint32(p136):]))
	store64(m.memory[int64(uint32(t127))+16:], uint64(t137))
	t139 := v1
	t140 := v7
	p138 := i32(0)
	if v11 != 0 {
		p138 = i32(-8)
	}
	v7 = t140 + p138
	t142 := v7
	t143 := v3
	p141 := i32(0)
	if uint32(v5) >= uint32(v9) {
		p141 = i32(-8)
	}
	v3 = t143 + p141
	t144 := int32(load32(m.memory[uint32(v3+i32(4)):]))
	t145 := v3
	v9 = t144
	t146 := int32(load32(m.memory[uint32(v7+i32(4)):]))
	t147 := v9
	v11 = t146
	var p148 int32
	if uint32(t147) < uint32(v11) {
		p148 = 1
	}
	v10 = p148
	p149 := t145
	if v10 != 0 {
		p149 = t142
	}
	t150 := int64(load64(m.memory[uint32(p149):]))
	store64(m.memory[int64(uint32(t139))+40:], uint64(t150))
	t151 := v1
	v5 = v0 + v8<<3
	t152 := v5
	t153 := v2
	var p154 int32
	if uint32(v4) >= uint32(v6) {
		p154 = 1
	}
	v0 = t153 + p154<<3
	t155 := int32(load32(m.memory[uint32(v5+i32(4)):]))
	t156 := v0
	v4 = t155
	t157 := int32(load32(m.memory[uint32(v0+i32(4)):]))
	t158 := v4
	v6 = t157
	var p159 int32
	if uint32(t158) < uint32(v6) {
		p159 = 1
	}
	v8 = p159
	p160 := t156
	if v8 != 0 {
		p160 = t152
	}
	t161 := int64(load64(m.memory[uint32(p160):]))
	store64(m.memory[int64(uint32(t151))+24:], uint64(t161))
	t163 := v1
	t164 := v7
	p162 := i32(0)
	if v10 != 0 {
		p162 = i32(-8)
	}
	v2 = t164 + p162
	t166 := v2
	t167 := v3
	p165 := i32(0)
	if uint32(v9) >= uint32(v11) {
		p165 = i32(-8)
	}
	v7 = t167 + p165
	t168 := int32(load32(m.memory[uint32(v7+i32(4)):]))
	t169 := v7
	v9 = t168
	t170 := int32(load32(m.memory[uint32(v2+i32(4)):]))
	t171 := v9
	v11 = t170
	var p172 int32
	if uint32(t171) < uint32(v11) {
		p172 = 1
	}
	v3 = p172
	p173 := t169
	if v3 != 0 {
		p173 = t166
	}
	t174 := int64(load64(m.memory[uint32(p173):]))
	store64(m.memory[int64(uint32(t163))+32:], uint64(t174))
	{
		t175 := v0
		var p176 int32
		if uint32(v4) >= uint32(v6) {
			p176 = 1
		}
		t178 := t175 + p176<<3
		t179 := v2
		p177 := i32(0)
		if v3 != 0 {
			p177 = i32(-8)
		}
		if t178 != t179+p177+i32(8) {
			goto l0
		}
		t181 := v5 + v8<<3
		t182 := v7
		p180 := i32(0)
		if uint32(v9) >= uint32(v11) {
			p180 = i32(-8)
		}
		if t181 == t182+p180+i32(8) {
			return
		}
	}
l0:
	m.fn122()
	panic("unreachable")
}
func (m *Module) fn128(v0, v1 int32) {
	var v2, v3, v4, v5 int32
	t0 := m.g0
	v2 = t0 - i32(4096)
	m.g0 = v2
	{
		{
			p1 := i32(666666)
			if uint32(v1) < uint32(i32(666666)) {
				p1 = v1
			}
			v3 = p1
			t2 := v3
			v4 = v1 - int32(uint32(v1)>>1)
			p3 := v4
			if uint32(v3) > uint32(v4) {
				p3 = t2
			}
			v3 = p3
			if uint32(v3) < uint32(i32(342)) {
				goto l0
			}
			if uint32(v4) >= uint32(i32(0xaaaaaab)) {
				m.fn9()
				panic("unreachable")
			}
			v5 = v3 * i32(12)
			t4 := m.fn5(v5)
			v4 = t4
			if v4 == 0 {
				m.fn10(i32(4), v5)
				panic("unreachable")
			}
			t5 := v0
			t6 := v1
			t7 := v4
			t8 := v3
			var p9 int32
			if uint32(v1) < uint32(i32(65)) {
				p9 = 1
			}
			m.fn129(t5, t6, t7, t8, p9)
			t10 := int32(load32(m.memory[uint32(v4+i32(-4)):]))
			v1 = t10
			v3 = v1 & i32(-8)
			t11 := v3
			v1 = v1 & i32(3)
			p12 := i32(8)
			if v1 != 0 {
				p12 = i32(4)
			}
			if uint32(t11) < uint32(p12+v5) {
				m.fn3(i32(1273840), i32(46), i32(1273888))
				panic("unreachable")
			}
			if v1 == 0 {
				goto l4
			}
			if uint32(v3) > uint32(v5+i32(39)) {
				m.fn3(i32(1273904), i32(46), i32(1273952))
				panic("unreachable")
			}
		l4:
			m.fn1(v4)
			goto l6
		}
	l0:
		t13 := v0
		t14 := v1
		t15 := v2
		var p16 int32
		if uint32(v1) < uint32(i32(65)) {
			p16 = 1
		}
		m.fn129(t13, t14, t15, i32(341), p16)
	}
l6:
	m.g0 = v2 + i32(4096)
}
func (m *Module) fn129(v0, v1, v2, v3, v4 int32) {
	var v5 int32
	var v6, v7 int64
	var v8, v9, v10, v11, v12, v13, v14, v15, v16, v17, v18, v19, v20, v21, v22, v23, v24, v25, v26 int32
	t0 := m.g0
	v5 = t0 - i32(336)
	m.g0 = v5
	v6 = int64(uint32(v1))
	t1 := int64(uint64(i64(0x4000000000000000)) / uint64(v6))
	v7 = t1
	var p2 int32
	if v7*v6 != i64(0x4000000000000000) {
		p2 = 1
	}
	v6 = int64(uint32(p2))
	{
		if uint32(v1) < uint32(i32(4097)) {
			goto l0
		}
		v8 = int32(bits.LeadingZeros32(uint32(v1|i32(1)))) ^ i32(31)
		v8 = int32(uint32(v8)>>1) + v8&i32(1)
		v9 = int32(uint32(i32_shl(i32(1), v8)+i32_shr_u(v1, v8)) >> 1)
		goto l1
	l0:
		v8 = v1 - int32(uint32(v1)>>1)
		p3 := i32(64)
		if uint32(v8) < uint32(i32(64)) {
			p3 = v8
		}
		v9 = p3
	}
l1:
	v7 = v7 + v6
	v10 = v0 + i32(-12)
	v11 = v0 + i32(24)
	v8 = i32(1)
	v12 = i32(0)
	v13 = i32(0)
l32:
	v14 = i32(0)
	v15 = i32(1)
	{
		var p4 int32
		if uint32(v1) > uint32(v12) {
			p4 = 1
		}
		v16 = p4
		if v16 == 0 {
			goto l2
		}
		t5 := v0
		v17 = v12 * i32(12)
		v18 = t5 + v17
		{
			v19 = v1 - v12
			if uint32(v19) < uint32(v9) {
				goto l3
			}
			if uint32(v19) >= uint32(i32(2)) {
				goto l4
			}
			v20 = v19
			goto l5
		l4:
			{
				t6 := int32(load32(m.memory[int64(uint32(v18))+12:]))
				v21 = t6
				t7 := int32(load32(m.memory[uint32(v18):]))
				var p8 int32
				if uint32(v21) < uint32(t7) {
					p8 = 1
				}
				v22 = p8
				if v22 != 0 {
					goto l6
				}
				v20 = i32(2)
				if v19 == i32(2) {
					goto l5
				}
				v23 = v11 + v17
				v20 = i32(2)
			l8:
				{
					t9 := int32(load32(m.memory[uint32(v23):]))
					v24 = t9
					if uint32(v24) < uint32(v21) {
						goto l7
					}
					v23 = v23 + i32(12)
					v21 = v24
					t10 := v19
					v20 = v20 + i32(1)
					if t10 != v20 {
						goto l8
					}
					goto l9
				}
			}
		l6:
			v20 = i32(2)
			v23 = i32(1)
			if v19 == i32(2) {
				goto l10
			}
			v23 = v11 + v17
			v20 = i32(2)
		l11:
			{
				t11 := int32(load32(m.memory[uint32(v23):]))
				v24 = t11
				if uint32(v24) >= uint32(v21) {
					goto l7
				}
				v23 = v23 + i32(12)
				v21 = v24
				t12 := v19
				v20 = v20 + i32(1)
				if t12 != v20 {
					goto l11
				}
			}
		l9:
			v20 = v19
		l7:
			if uint32(v20) < uint32(v9) {
				goto l3
			}
			if v22 == 0 {
				goto l5
			}
			v23 = int32(uint32(v20) >> 1)
			if v23 == 0 {
				goto l5
			}
		l10:
			v19 = v10 + (v20*i32(12) + v17)
		l12:
			{
				t13 := int32(load32(m.memory[uint32(v18):]))
				v21 = t13
				t14 := int32(load32(m.memory[uint32(v19):]))
				store32(m.memory[uint32(v18):], uint32(t14))
				store32(m.memory[uint32(v19):], uint32(v21))
				v21 = v18 + i32(4)
				t15 := int64(load64(m.memory[uint32(v21):]))
				v6 = t15
				t16 := v21
				v24 = v19 + i32(4)
				t17 := int64(load64(m.memory[uint32(v24):]))
				store64(m.memory[uint32(t16):], uint64(t17))
				store64(m.memory[uint32(v24):], uint64(v6))
				v19 = v19 + i32(-12)
				v18 = v18 + i32(12)
				v23 = v23 + i32(-1)
				if v23 != 0 {
					goto l12
				}
			}
		l5:
			v15 = v20<<1 | i32(1)
			goto l13
		l3:
			{
				if v4 != 0 {
					goto l14
				}
				p18 := v9
				if uint32(v19) < uint32(v9) {
					p18 = v19
				}
				v15 = p18 << 1
				goto l13
			}
		l14:
			t20 := v18
			p19 := i32(32)
			if uint32(v19) < uint32(i32(32)) {
				p19 = v19
			}
			v19 = p19
			m.fn130(t20, v19, v2, v3, i32(0), i32(0))
			v15 = v19<<1 | i32(1)
		}
	l13:
		v14 = int32(int64(bits.LeadingZeros64(uint64(v7*int64(uint32(int32(uint32(v15)>>1)+v12<<1)) ^ (int64(uint32(v12-int32(uint32(v8)>>1)))+int64(uint32(v12)))*v7))))
	}
l2:
	{
		if uint32(v13) < uint32(i32(2)) {
			goto l15
		}
		t21 := v10
		v18 = v12 * i32(12)
		v25 = t21 + v18
		v26 = v0 + v18
	l29:
		{
			t22 := v5 + i32(270)
			v23 = v13 + i32(-1)
			t23 := int32(m.memory[uint32(t22+v23)])
			if uint32(t23) < uint32(v14) {
				goto l15
			}
			{
				t24 := int32(load32(m.memory[uint32(v5+i32(4)+v23<<2):]))
				v20 = t24
				v18 = int32(uint32(v20) >> 1)
				t25 := v18
				v19 = int32(uint32(v8) >> 1)
				v22 = t25 + v19
				if uint32(v22) > uint32(v3) {
					goto l16
				}
				if (v20|v8)&i32(1) == 0 {
					v8 = v22 << 1
					goto l20
				}
			}
		l16:
			v13 = v0 + (v12-v22)*i32(12)
			if v20&i32(1) == 0 {
				goto l18
			}
			goto l19
		l18:
			m.fn130(v13, v18, v2, v3, int32(bits.LeadingZeros32(uint32(v18|i32(1))))<<1^i32(62), i32(0))
		l19:
			if v8&i32(1) != 0 {
				goto l21
			}
			m.fn130(v13+v18*i32(12), v19, v2, v3, int32(bits.LeadingZeros32(uint32(v19|i32(1))))<<1^i32(62), i32(0))
		l21:
			{
				if v19 == 0 {
					goto l22
				}
				if v18 == 0 {
					goto l22
				}
				t26 := v3
				t27 := v19
				t28 := v18
				var p29 int32
				if uint32(v19) < uint32(v18) {
					p29 = 1
				}
				v20 = p29
				p30 := t28
				if v20 != 0 {
					p30 = t27
				}
				v19 = p30
				if uint32(t26) < uint32(v19) {
					goto l22
				}
				v8 = v13 + v18*i32(12)
				{
					v18 = v19 * i32(12)
					if v18 == 0 {
						goto l23
					}
					t32 := v2
					p31 := v13
					if v20 != 0 {
						p31 = v8
					}
					memory_copy(m.memory, uint32(t32), uint32(p31), uint32(v18))
				}
			l23:
				v19 = v2 + v18
				if v20 != 0 {
					goto l24
				}
				v18 = v2
			l26:
				{
					t33 := int32(load32(m.memory[uint32(v8):]))
					t34 := v13
					t35 := v8
					t36 := v18
					v20 = t33
					t37 := int32(load32(m.memory[uint32(v18):]))
					t38 := v20
					v21 = t37
					var p39 int32
					if uint32(t38) < uint32(v21) {
						p39 = 1
					}
					v17 = p39
					p40 := t36
					if v17 != 0 {
						p40 = t35
					}
					v24 = p40
					t41 := int64(load64(m.memory[uint32(v24):]))
					store64(m.memory[uint32(t34):], uint64(t41))
					t42 := int32(load32(m.memory[int64(uint32(v24))+8:]))
					store32(m.memory[int64(uint32(v13))+8:], uint32(t42))
					v13 = v13 + i32(12)
					t43 := v18
					var p44 int32
					if uint32(v20) >= uint32(v21) {
						p44 = 1
					}
					v18 = t43 + p44*i32(12)
					if v18 == v19 {
						goto l25
					}
					v8 = v8 + v17*i32(12)
					if v8 != v26 {
						goto l26
					}
					goto l25
				}
			l24:
				v18 = v25
			l28:
				{
					t45 := v18
					v8 = v8 + i32(-12)
					t46 := v8
					v19 = v19 + i32(-12)
					t47 := int32(load32(m.memory[uint32(v19):]))
					t48 := v19
					v20 = t47
					t49 := int32(load32(m.memory[uint32(v8):]))
					t50 := v20
					v21 = t49
					var p51 int32
					if uint32(t50) < uint32(v21) {
						p51 = 1
					}
					v24 = p51
					p52 := t48
					if v24 != 0 {
						p52 = t46
					}
					v17 = p52
					t53 := int64(load64(m.memory[uint32(v17):]))
					store64(m.memory[uint32(t45):], uint64(t53))
					t54 := int32(load32(m.memory[int64(uint32(v17))+8:]))
					store32(m.memory[int64(uint32(v18))+8:], uint32(t54))
					v19 = v19 + v24*i32(12)
					t55 := v8
					var p56 int32
					if uint32(v20) >= uint32(v21) {
						p56 = 1
					}
					v8 = t55 + p56*i32(12)
					if v8 == v13 {
						goto l27
					}
					v18 = v18 + i32(-12)
					if v19 != v2 {
						goto l28
					}
				}
			l27:
				v13 = v8
				v18 = v2
			l25:
				v8 = v19 - v18
				if v8 == 0 {
					goto l22
				}
				memory_copy(m.memory, uint32(v13), uint32(v18), uint32(v8))
			}
		l22:
			v8 = v22<<1 | i32(1)
		l20:
			v18 = i32(1)
			v13 = v23
			if uint32(v23) > uint32(i32(1)) {
				goto l29
			}
			goto l30
		}
	}
l15:
	v18 = v13
l30:
	m.memory[uint32(v5+i32(270)+v18)] = byte(v14)
	store32(m.memory[uint32(v5+i32(4)+v18<<2):], uint32(v8))
	if v16 == 0 {
		goto l31
	}
	v13 = v18 + i32(1)
	v12 = int32(uint32(v15)>>1) + v12
	v8 = v15
	goto l32
l31:
	if v8&i32(1) != 0 {
		goto l33
	}
	m.fn130(v0, v1, v2, v3, int32(bits.LeadingZeros32(uint32(v1|i32(1))))<<1^i32(62), i32(0))
l33:
	m.g0 = v5 + i32(336)
}
func (m *Module) fn130(v0, v1, v2, v3, v4, v5 int32) {
	var v6, v7, v8, v9, v10, v11, v12, v13, v14, v15, v16, v17, v18, v19, v20, v21 int32
	var v22 int64
	t0 := m.g0
	v6 = t0 - i32(16)
	m.g0 = v6
	{
		if uint32(v1) >= uint32(i32(33)) {
			v8 = v2 + i32(-12)
		l20:
			{
				if v4 != 0 {
					goto l2
				}
				m.fn129(v0, v1, v2, v3, i32(1))
				goto l3
			l2:
				t1 := v0
				v9 = int32(uint32(v1) >> 3)
				v10 = t1 + v9*i32(84)
				v7 = v0 + v9*i32(48)
				{
					{
						if uint32(v1) < uint32(i32(64)) {
							goto l4
						}
						t2 := m.fn131(v0, v7, v10, v9)
						v11 = t2
						goto l5
					}
				l4:
					t3 := int32(load32(m.memory[uint32(v0):]))
					t4 := v0
					t5 := v10
					t6 := v7
					v9 = t3
					t7 := int32(load32(m.memory[uint32(v7):]))
					t8 := v9
					v11 = t7
					var p9 int32
					if uint32(t8) < uint32(v11) {
						p9 = 1
					}
					v12 = p9
					t10 := int32(load32(m.memory[uint32(v10):]))
					t11 := v12
					t12 := v11
					v13 = t10
					var p13 int32
					if uint32(t12) < uint32(v13) {
						p13 = 1
					}
					p14 := t6
					if t11^p13 != 0 {
						p14 = t5
					}
					t15 := v12
					var p16 int32
					if uint32(v9) < uint32(v13) {
						p16 = 1
					}
					p17 := p14
					if t15^p16 != 0 {
						p17 = t4
					}
					v11 = p17
				}
			l5:
				v4 = v4 + i32(-1)
				t18 := int32(load32(m.memory[int64(uint32(v11))+8:]))
				store32(m.memory[int64(uint32(v6))+8:], uint32(t18))
				t19 := int64(load64(m.memory[uint32(v11):]))
				store64(m.memory[uint32(v6):], uint64(t19))
				t20 := int32(uint32(v11-v0) / uint32(i32(12)))
				v14 = t20
				{
					{
						if v5 == 0 {
							goto l6
						}
						t21 := int32(load32(m.memory[uint32(v5):]))
						t22 := int32(load32(m.memory[uint32(v11):]))
						if uint32(t21) >= uint32(t22) {
							goto l7
						}
					}
				l6:
					if uint32(v3) < uint32(v1) {
						goto l8
					}
					v7 = i32(0)
					v9 = v0
					t23 := v2
					v15 = v1 * i32(12)
					v16 = t23 + v15
					v10 = v16
					v17 = v14
				l14:
					{
						t24 := v9
						t25 := v0
						v12 = v17 + i32(-3)
						p26 := v12
						if uint32(v12) > uint32(v17) {
							p26 = i32(0)
						}
						v18 = t25 + p26*i32(12)
						if uint32(t24) >= uint32(v18) {
							goto l9
						}
					l10:
						{
							t27 := int32(load32(m.memory[uint32(v9):]))
							t28 := int32(load32(m.memory[uint32(v11):]))
							t29 := v2
							t30 := v10 + i32(-12)
							var p31 int32
							if uint32(t27) < uint32(t28) {
								p31 = 1
							}
							v12 = p31
							p32 := t30
							if v12 != 0 {
								p32 = t29
							}
							v13 = p32 + v7*i32(12)
							t33 := int32(load32(m.memory[int64(uint32(v9))+8:]))
							store32(m.memory[int64(uint32(v13))+8:], uint32(t33))
							t34 := int64(load64(m.memory[uint32(v9):]))
							store64(m.memory[uint32(v13):], uint64(t34))
							t35 := v2
							t36 := v10 + i32(-24)
							v13 = v9 + i32(12)
							t37 := int32(load32(m.memory[uint32(v13):]))
							t38 := int32(load32(m.memory[uint32(v11):]))
							var p39 int32
							if uint32(t37) < uint32(t38) {
								p39 = 1
							}
							v19 = p39
							p40 := t36
							if v19 != 0 {
								p40 = t35
							}
							v7 = v7 + v12
							v12 = p40 + v7*i32(12)
							t41 := int32(load32(m.memory[uint32(v9+i32(20)):]))
							store32(m.memory[int64(uint32(v12))+8:], uint32(t41))
							t42 := int64(load64(m.memory[uint32(v13):]))
							store64(m.memory[uint32(v12):], uint64(t42))
							t43 := v2
							t44 := v10 + i32(-36)
							v12 = v9 + i32(24)
							t45 := int32(load32(m.memory[uint32(v12):]))
							t46 := int32(load32(m.memory[uint32(v11):]))
							var p47 int32
							if uint32(t45) < uint32(t46) {
								p47 = 1
							}
							v13 = p47
							p48 := t44
							if v13 != 0 {
								p48 = t43
							}
							v7 = v7 + v19
							v19 = p48 + v7*i32(12)
							t49 := int32(load32(m.memory[uint32(v9+i32(32)):]))
							store32(m.memory[int64(uint32(v19))+8:], uint32(t49))
							t50 := int64(load64(m.memory[uint32(v12):]))
							store64(m.memory[uint32(v19):], uint64(t50))
							t51 := v2
							v10 = v10 + i32(-48)
							t52 := v10
							v12 = v9 + i32(36)
							t53 := int32(load32(m.memory[uint32(v12):]))
							t54 := int32(load32(m.memory[uint32(v11):]))
							var p55 int32
							if uint32(t53) < uint32(t54) {
								p55 = 1
							}
							v19 = p55
							p56 := t52
							if v19 != 0 {
								p56 = t51
							}
							v7 = v7 + v13
							v13 = p56 + v7*i32(12)
							t57 := int32(load32(m.memory[uint32(v9+i32(44)):]))
							store32(m.memory[int64(uint32(v13))+8:], uint32(t57))
							t58 := int64(load64(m.memory[uint32(v12):]))
							store64(m.memory[uint32(v13):], uint64(t58))
							v7 = v7 + v19
							v9 = v9 + i32(48)
							if uint32(v9) < uint32(v18) {
								goto l10
							}
						}
					}
				l9:
					{
						t59 := v9
						v19 = v0 + v17*i32(12)
						if uint32(t59) >= uint32(v19) {
							goto l11
						}
					l12:
						{
							t60 := v2
							v10 = v10 + i32(-12)
							t61 := int32(load32(m.memory[uint32(v9):]))
							t62 := int32(load32(m.memory[uint32(v11):]))
							t63 := v10
							var p64 int32
							if uint32(t61) < uint32(t62) {
								p64 = 1
							}
							v12 = p64
							p65 := t63
							if v12 != 0 {
								p65 = t60
							}
							v13 = p65 + v7*i32(12)
							t66 := int32(load32(m.memory[int64(uint32(v9))+8:]))
							store32(m.memory[int64(uint32(v13))+8:], uint32(t66))
							t67 := int64(load64(m.memory[uint32(v9):]))
							store64(m.memory[uint32(v13):], uint64(t67))
							v7 = v7 + v12
							v9 = v9 + i32(12)
							if uint32(v9) < uint32(v19) {
								goto l12
							}
						}
					}
				l11:
					{
						if v17 == v1 {
							v18 = v7 * i32(12)
							if v18 == 0 {
								goto l15
							}
							memory_copy(m.memory, uint32(v0), uint32(v2), uint32(v18))
						l15:
							v17 = v1 - v7
							{
								if v1 == v7 {
									goto l16
								}
								v20 = v0 + v18
								v12 = i32(0)
								if v1 == v7+i32(1) {
									goto l17
								}
								v21 = v17 & i32(1)
								v19 = v17 & i32(-2)
								v10 = v8 + v15
								v12 = i32(0)
								v9 = v20
							l18:
								{
									t70 := int32(load32(m.memory[int64(uint32(v10))+8:]))
									store32(m.memory[int64(uint32(v9))+8:], uint32(t70))
									t71 := int64(load64(m.memory[uint32(v10):]))
									store64(m.memory[uint32(v9):], uint64(t71))
									t72 := v9 + i32(12)
									v13 = v16 + (v12^i32(0x3ffffffe))*i32(12)
									t73 := int64(load64(m.memory[uint32(v13):]))
									store64(m.memory[uint32(t72):], uint64(t73))
									t74 := int32(load32(m.memory[uint32(v13+i32(8)):]))
									store32(m.memory[uint32(v9+i32(20)):], uint32(t74))
									v10 = v10 + i32(-24)
									v9 = v9 + i32(24)
									t75 := v19
									v12 = v12 + i32(2)
									if t75 != v12 {
										goto l18
									}
								}
								if v21 == 0 {
									goto l16
								}
							l17:
								v9 = v20 + v12*i32(12)
								t76 := v9
								v10 = v16 + (v12^i32(-1))*i32(12)
								t77 := int32(load32(m.memory[int64(uint32(v10))+8:]))
								store32(m.memory[int64(uint32(t76))+8:], uint32(t77))
								t78 := int64(load64(m.memory[uint32(v10):]))
								store64(m.memory[uint32(v9):], uint64(t78))
							}
						l16:
							if v7 == 0 {
								goto l7
							}
							if uint32(v1) < uint32(v7) {
								goto l19
							}
							m.fn130(v0+v18, v17, v2, v3, v4, v6)
							v1 = v7
							if uint32(v7) < uint32(i32(33)) {
								goto l1
							}
							goto l20
						}
						v10 = v10 + i32(-12)
						v12 = v10 + v7*i32(12)
						t68 := int32(load32(m.memory[int64(uint32(v9))+8:]))
						store32(m.memory[int64(uint32(v12))+8:], uint32(t68))
						t69 := int64(load64(m.memory[uint32(v9):]))
						store64(m.memory[uint32(v12):], uint64(t69))
						v9 = v9 + i32(12)
						v17 = v1
						goto l14
					}
				}
			l7:
				if uint32(v3) < uint32(v1) {
					goto l8
				}
				v12 = i32(0)
				v9 = v0
				t79 := v2
				v16 = v1 * i32(12)
				v17 = t79 + v16
				v10 = v17
			l26:
				{
					t80 := v9
					t81 := v0
					v7 = v14 + i32(-3)
					p82 := v7
					if uint32(v7) > uint32(v14) {
						p82 = i32(0)
					}
					v18 = t81 + p82*i32(12)
					if uint32(t80) >= uint32(v18) {
						goto l21
					}
				l22:
					{
						t83 := int32(load32(m.memory[uint32(v11):]))
						t84 := int32(load32(m.memory[uint32(v9):]))
						t85 := v2
						t86 := v10 + i32(-12)
						var p87 int32
						if uint32(t83) >= uint32(t84) {
							p87 = 1
						}
						v7 = p87
						p88 := t86
						if v7 != 0 {
							p88 = t85
						}
						v13 = p88 + v12*i32(12)
						t89 := int32(load32(m.memory[int64(uint32(v9))+8:]))
						store32(m.memory[int64(uint32(v13))+8:], uint32(t89))
						t90 := int64(load64(m.memory[uint32(v9):]))
						store64(m.memory[uint32(v13):], uint64(t90))
						t91 := int32(load32(m.memory[uint32(v11):]))
						t92 := v2
						t93 := v10 + i32(-24)
						v13 = v9 + i32(12)
						t94 := int32(load32(m.memory[uint32(v13):]))
						var p95 int32
						if uint32(t91) >= uint32(t94) {
							p95 = 1
						}
						v19 = p95
						p96 := t93
						if v19 != 0 {
							p96 = t92
						}
						v7 = v12 + v7
						v12 = p96 + v7*i32(12)
						t97 := int32(load32(m.memory[uint32(v9+i32(20)):]))
						store32(m.memory[int64(uint32(v12))+8:], uint32(t97))
						t98 := int64(load64(m.memory[uint32(v13):]))
						store64(m.memory[uint32(v12):], uint64(t98))
						t99 := int32(load32(m.memory[uint32(v11):]))
						t100 := v2
						t101 := v10 + i32(-36)
						v12 = v9 + i32(24)
						t102 := int32(load32(m.memory[uint32(v12):]))
						var p103 int32
						if uint32(t99) >= uint32(t102) {
							p103 = 1
						}
						v13 = p103
						p104 := t101
						if v13 != 0 {
							p104 = t100
						}
						v7 = v7 + v19
						v19 = p104 + v7*i32(12)
						t105 := int32(load32(m.memory[uint32(v9+i32(32)):]))
						store32(m.memory[int64(uint32(v19))+8:], uint32(t105))
						t106 := int64(load64(m.memory[uint32(v12):]))
						store64(m.memory[uint32(v19):], uint64(t106))
						t107 := v2
						v10 = v10 + i32(-48)
						t108 := int32(load32(m.memory[uint32(v11):]))
						t109 := v10
						v12 = v9 + i32(36)
						t110 := int32(load32(m.memory[uint32(v12):]))
						var p111 int32
						if uint32(t108) >= uint32(t110) {
							p111 = 1
						}
						v19 = p111
						p112 := t109
						if v19 != 0 {
							p112 = t107
						}
						v7 = v7 + v13
						v13 = p112 + v7*i32(12)
						t113 := int32(load32(m.memory[uint32(v9+i32(44)):]))
						store32(m.memory[int64(uint32(v13))+8:], uint32(t113))
						t114 := int64(load64(m.memory[uint32(v12):]))
						store64(m.memory[uint32(v13):], uint64(t114))
						v12 = v7 + v19
						v9 = v9 + i32(48)
						if uint32(v9) < uint32(v18) {
							goto l22
						}
					}
				}
			l21:
				{
					t115 := v9
					v19 = v0 + v14*i32(12)
					if uint32(t115) >= uint32(v19) {
						goto l23
					}
				l24:
					{
						t116 := v2
						v10 = v10 + i32(-12)
						t117 := int32(load32(m.memory[uint32(v11):]))
						t118 := int32(load32(m.memory[uint32(v9):]))
						t119 := v10
						var p120 int32
						if uint32(t117) >= uint32(t118) {
							p120 = 1
						}
						v7 = p120
						p121 := t119
						if v7 != 0 {
							p121 = t116
						}
						v13 = p121 + v12*i32(12)
						t122 := int32(load32(m.memory[int64(uint32(v9))+8:]))
						store32(m.memory[int64(uint32(v13))+8:], uint32(t122))
						t123 := int64(load64(m.memory[uint32(v9):]))
						store64(m.memory[uint32(v13):], uint64(t123))
						v12 = v12 + v7
						v9 = v9 + i32(12)
						if uint32(v9) < uint32(v19) {
							goto l24
						}
					}
				}
			l23:
				{
					if v14 == v1 {
						goto l25
					}
					v7 = v2 + v12*i32(12)
					t124 := int32(load32(m.memory[int64(uint32(v9))+8:]))
					store32(m.memory[int64(uint32(v7))+8:], uint32(t124))
					t125 := int64(load64(m.memory[uint32(v9):]))
					store64(m.memory[uint32(v7):], uint64(t125))
					v9 = v9 + i32(12)
					v12 = v12 + i32(1)
					v10 = v10 + i32(-12)
					v14 = v1
					goto l26
				}
			l25:
				v9 = v12 * i32(12)
				if v9 == 0 {
					goto l27
				}
				memory_copy(m.memory, uint32(v0), uint32(v2), uint32(v9))
			l27:
				if v1 == v12 {
					goto l3
				}
				v7 = v1 - v12
				v0 = v0 + v9
				v11 = i32(0)
				{
					if v1 == v12+i32(1) {
						goto l28
					}
					v18 = v7 & i32(1)
					v19 = v7 & i32(-2)
					v10 = v8 + v16
					v11 = i32(0)
					v9 = v0
				l29:
					{
						t126 := int32(load32(m.memory[int64(uint32(v10))+8:]))
						store32(m.memory[int64(uint32(v9))+8:], uint32(t126))
						t127 := int64(load64(m.memory[uint32(v10):]))
						store64(m.memory[uint32(v9):], uint64(t127))
						t128 := v9 + i32(12)
						v13 = v17 + (v11^i32(0x3ffffffe))*i32(12)
						t129 := int64(load64(m.memory[uint32(v13):]))
						store64(m.memory[uint32(t128):], uint64(t129))
						t130 := int32(load32(m.memory[uint32(v13+i32(8)):]))
						store32(m.memory[uint32(v9+i32(20)):], uint32(t130))
						v10 = v10 + i32(-24)
						v9 = v9 + i32(24)
						t131 := v19
						v11 = v11 + i32(2)
						if t131 != v11 {
							goto l29
						}
					}
					if v18 == 0 {
						goto l30
					}
				l28:
					v9 = v0 + v11*i32(12)
					t132 := v9
					v10 = v17 + (v11^i32(-1))*i32(12)
					t133 := int32(load32(m.memory[int64(uint32(v10))+8:]))
					store32(m.memory[int64(uint32(t132))+8:], uint32(t133))
					t134 := int64(load64(m.memory[uint32(v10):]))
					store64(m.memory[uint32(v9):], uint64(t134))
				}
			l30:
				if uint32(v1) < uint32(v12) {
					goto l31
				}
				v5 = i32(0)
				v1 = v7
				if uint32(v7) < uint32(i32(33)) {
					goto l1
				}
				goto l20
			l31:
			}
			m.fn121(v12, v1, v1, i32(1069572))
			panic("unreachable")
		l19:
			m.fn28(i32(1271784), i32(19), i32(1069556))
		l8:
			panic("unreachable")
		}
		v7 = v1
		goto l1
	l1:
		if uint32(v7) < uint32(i32(2)) {
			goto l3
		}
		v17 = int32(uint32(v7) >> 1)
		{
			if uint32(v7) > uint32(i32(15)) {
				goto l32
			}
			{
				if uint32(v7) <= uint32(i32(7)) {
					t234 := int32(load32(m.memory[int64(uint32(v0))+8:]))
					store32(m.memory[int64(uint32(v2))+8:], uint32(t234))
					t235 := int64(load64(m.memory[uint32(v0):]))
					store64(m.memory[uint32(v2):], uint64(t235))
					t236 := v2
					v9 = v17 * i32(12)
					v10 = t236 + v9
					t237 := v10
					v9 = v0 + v9
					t238 := int32(load32(m.memory[int64(uint32(v9))+8:]))
					store32(m.memory[int64(uint32(t237))+8:], uint32(t238))
					t239 := int64(load64(m.memory[uint32(v9):]))
					store64(m.memory[uint32(v10):], uint64(t239))
					v1 = i32(1)
					goto l34
				}
				t135 := int32(load32(m.memory[int64(uint32(v0))+12:]))
				t136 := v0
				v10 = t135
				t137 := int32(load32(m.memory[uint32(v0):]))
				t138 := v10
				v12 = t137
				var p139 int32
				if uint32(t138) < uint32(v12) {
					p139 = 1
				}
				v11 = t136 + p139*i32(12)
				t140 := int32(load32(m.memory[int64(uint32(v0))+36:]))
				t141 := int32(load32(m.memory[int64(uint32(v0))+24:]))
				t142 := v11
				t143 := v0
				var p144 int32
				if uint32(t140) < uint32(t141) {
					p144 = 1
				}
				v13 = p144
				p145 := i32(24)
				if v13 != 0 {
					p145 = i32(36)
				}
				v9 = t143 + p145
				t146 := v9
				t147 := v0
				var p148 int32
				if uint32(v10) >= uint32(v12) {
					p148 = 1
				}
				v10 = t147 + p148*i32(12)
				t150 := v10
				t151 := v0
				p149 := i32(36)
				if v13 != 0 {
					p149 = i32(24)
				}
				v12 = t151 + p149
				t152 := int32(load32(m.memory[uint32(v12):]))
				t153 := int32(load32(m.memory[uint32(v10):]))
				var p154 int32
				if uint32(t152) < uint32(t153) {
					p154 = 1
				}
				v13 = p154
				p155 := t150
				if v13 != 0 {
					p155 = t146
				}
				t156 := int32(load32(m.memory[uint32(v9):]))
				t157 := int32(load32(m.memory[uint32(v11):]))
				var p158 int32
				if uint32(t156) < uint32(t157) {
					p158 = 1
				}
				v19 = p158
				p159 := p155
				if v19 != 0 {
					p159 = t142
				}
				v18 = p159
				t160 := int32(load32(m.memory[uint32(v18):]))
				v14 = t160
				t162 := v12
				p161 := v9
				if v19 != 0 {
					p161 = v10
				}
				p163 := p161
				if v13 != 0 {
					p163 = t162
				}
				v1 = p163
				t164 := int32(load32(m.memory[uint32(v1):]))
				v16 = t164
				t166 := v2
				p165 := v11
				if v19 != 0 {
					p165 = v9
				}
				v9 = p165
				t167 := int32(load32(m.memory[int64(uint32(v9))+8:]))
				store32(m.memory[int64(uint32(t166))+8:], uint32(t167))
				t168 := int64(load64(m.memory[uint32(v9):]))
				store64(m.memory[uint32(v2):], uint64(t168))
				t169 := v2
				t170 := v1
				t171 := v18
				var p172 int32
				if uint32(v16) < uint32(v14) {
					p172 = 1
				}
				v9 = p172
				p173 := t171
				if v9 != 0 {
					p173 = t170
				}
				v11 = p173
				t174 := int32(load32(m.memory[int64(uint32(v11))+8:]))
				store32(m.memory[int64(uint32(t169))+20:], uint32(t174))
				t175 := int64(load64(m.memory[uint32(v11):]))
				store64(m.memory[int64(uint32(v2))+12:], uint64(t175))
				t177 := v2
				p176 := v1
				if v9 != 0 {
					p176 = v18
				}
				v9 = p176
				t178 := int32(load32(m.memory[int64(uint32(v9))+8:]))
				store32(m.memory[int64(uint32(t177))+32:], uint32(t178))
				t179 := int64(load64(m.memory[uint32(v9):]))
				store64(m.memory[int64(uint32(v2))+24:], uint64(t179))
				t181 := v2
				p180 := v12
				if v13 != 0 {
					p180 = v10
				}
				v9 = p180
				t182 := int32(load32(m.memory[int64(uint32(v9))+8:]))
				store32(m.memory[int64(uint32(t181))+44:], uint32(t182))
				t183 := int64(load64(m.memory[uint32(v9):]))
				store64(m.memory[int64(uint32(v2))+36:], uint64(t183))
				t184 := v0
				v16 = v17 * i32(12)
				v9 = t184 + v16
				t185 := int32(load32(m.memory[int64(uint32(v9))+36:]))
				t186 := int32(load32(m.memory[int64(uint32(v9))+24:]))
				t187 := v9
				var p188 int32
				if uint32(t185) < uint32(t186) {
					p188 = 1
				}
				v11 = p188
				p189 := i32(36)
				if v11 != 0 {
					p189 = i32(24)
				}
				v12 = t187 + p189
				t190 := int32(load32(m.memory[int64(uint32(v9))+12:]))
				t191 := v12
				t192 := v9
				v13 = t190
				t193 := int32(load32(m.memory[uint32(v9):]))
				t194 := v13
				v19 = t193
				var p195 int32
				if uint32(t194) >= uint32(v19) {
					p195 = 1
				}
				v10 = t192 + p195*i32(12)
				t197 := v10
				t198 := v9
				p196 := i32(24)
				if v11 != 0 {
					p196 = i32(36)
				}
				v11 = t198 + p196
				t199 := int32(load32(m.memory[uint32(v11):]))
				t200 := v11
				t201 := v9
				var p202 int32
				if uint32(v13) < uint32(v19) {
					p202 = 1
				}
				v13 = t201 + p202*i32(12)
				t203 := int32(load32(m.memory[uint32(v13):]))
				var p204 int32
				if uint32(t199) < uint32(t203) {
					p204 = 1
				}
				v19 = p204
				p205 := t200
				if v19 != 0 {
					p205 = t197
				}
				t206 := int32(load32(m.memory[uint32(v12):]))
				t207 := int32(load32(m.memory[uint32(v10):]))
				var p208 int32
				if uint32(t206) < uint32(t207) {
					p208 = 1
				}
				v18 = p208
				p209 := p205
				if v18 != 0 {
					p209 = t191
				}
				v1 = p209
				t210 := int32(load32(m.memory[uint32(v1):]))
				v4 = t210
				t212 := v13
				p211 := v10
				if v18 != 0 {
					p211 = v11
				}
				p213 := p211
				if v19 != 0 {
					p213 = t212
				}
				v14 = p213
				t214 := int32(load32(m.memory[uint32(v14):]))
				v3 = t214
				v9 = v2 + v16
				t216 := v9
				p215 := v13
				if v19 != 0 {
					p215 = v11
				}
				v11 = p215
				t217 := int32(load32(m.memory[int64(uint32(v11))+8:]))
				store32(m.memory[int64(uint32(t216))+8:], uint32(t217))
				t218 := int64(load64(m.memory[uint32(v11):]))
				store64(m.memory[uint32(v9):], uint64(t218))
				t219 := v9
				t220 := v1
				t221 := v14
				var p222 int32
				if uint32(v4) < uint32(v3) {
					p222 = 1
				}
				v11 = p222
				p223 := t221
				if v11 != 0 {
					p223 = t220
				}
				v13 = p223
				t224 := int64(load64(m.memory[uint32(v13):]))
				store64(m.memory[int64(uint32(t219))+12:], uint64(t224))
				t225 := int32(load32(m.memory[int64(uint32(v13))+8:]))
				store32(m.memory[int64(uint32(v9))+20:], uint32(t225))
				t227 := v9
				p226 := v1
				if v11 != 0 {
					p226 = v14
				}
				v11 = p226
				t228 := int64(load64(m.memory[uint32(v11):]))
				store64(m.memory[int64(uint32(t227))+24:], uint64(t228))
				t229 := int32(load32(m.memory[int64(uint32(v11))+8:]))
				store32(m.memory[int64(uint32(v9))+32:], uint32(t229))
				t231 := v9
				p230 := v12
				if v18 != 0 {
					p230 = v10
				}
				v10 = p230
				t232 := int64(load64(m.memory[uint32(v10):]))
				store64(m.memory[int64(uint32(t231))+36:], uint64(t232))
				t233 := int32(load32(m.memory[int64(uint32(v10))+8:]))
				store32(m.memory[int64(uint32(v9))+44:], uint32(t233))
				v1 = i32(4)
				goto l34
			}
		l32:
			t240 := v0
			t241 := v2
			v9 = v2 + v7*i32(12)
			m.fn132(t240, t241, v9)
			t242 := v0
			v10 = v17 * i32(12)
			m.fn132(t242+v10, v2+v10, v9+i32(96))
			v1 = i32(8)
		}
	l34:
		v4 = v7 - v17
		if uint32(v1) >= uint32(v17) {
			goto l35
		}
		v19 = v1 * i32(12)
		v13 = v1
	l40:
		{
			t243 := v2
			v10 = v13 * i32(12)
			v9 = t243 + v10
			t244 := v9
			v10 = v0 + v10
			t245 := int32(load32(m.memory[int64(uint32(v10))+8:]))
			store32(m.memory[int64(uint32(t244))+8:], uint32(t245))
			t246 := int64(load64(m.memory[uint32(v10):]))
			t247 := v9
			v22 = t246
			store64(m.memory[uint32(t247):], uint64(v22))
			{
				v12 = int32(v22)
				t248 := int32(load32(m.memory[uint32(v9+i32(-12)):]))
				if uint32(v12) >= uint32(t248) {
					goto l36
				}
				t249 := int64(load64(m.memory[int64(uint32(v9))+4:]))
				v22 = t249
				v9 = v19
			l39:
				{
					v10 = v2 + v9
					t250 := v10
					v11 = v10 + i32(-12)
					t251 := int32(load32(m.memory[int64(uint32(v11))+8:]))
					store32(m.memory[int64(uint32(t250))+8:], uint32(t251))
					t252 := int64(load64(m.memory[uint32(v11):]))
					store64(m.memory[uint32(v10):], uint64(t252))
					if v9 != i32(12) {
						goto l37
					}
					v9 = v2
					goto l38
				l37:
					v9 = v9 + i32(-12)
					t253 := int32(load32(m.memory[uint32(v10+i32(-24)):]))
					if uint32(v12) < uint32(t253) {
						goto l39
					}
				}
				v9 = v2 + v9
			l38:
				store64(m.memory[int64(uint32(v9))+4:], uint64(v22))
				store32(m.memory[uint32(v9):], uint32(v12))
			}
		l36:
			v19 = v19 + i32(12)
			v13 = v13 + i32(1)
			if v13 != v17 {
				goto l40
			}
		}
	l35:
		t254 := v2
		v9 = v17 * i32(12)
		v19 = t254 + v9
		if uint32(v1) >= uint32(v4) {
			goto l41
		}
		v3 = v0 + v9
		v12 = v1 * i32(12)
		v14 = i32(12)
		v16 = v19
	l46:
		{
			t255 := v19
			v10 = v1 * i32(12)
			v9 = t255 + v10
			t256 := v9
			v10 = v3 + v10
			t257 := int32(load32(m.memory[int64(uint32(v10))+8:]))
			store32(m.memory[int64(uint32(t256))+8:], uint32(t257))
			t258 := int64(load64(m.memory[uint32(v10):]))
			t259 := v9
			v22 = t258
			store64(m.memory[uint32(t259):], uint64(v22))
			{
				v18 = int32(v22)
				t260 := int32(load32(m.memory[uint32(v9+i32(-12)):]))
				if uint32(v18) >= uint32(t260) {
					goto l42
				}
				t261 := int64(load64(m.memory[int64(uint32(v9))+4:]))
				v22 = t261
				v11 = v14
				v10 = v16
			l45:
				{
					v9 = v10 + v12
					t262 := v9
					v13 = v9 + i32(-12)
					t263 := int32(load32(m.memory[int64(uint32(v13))+8:]))
					store32(m.memory[int64(uint32(t262))+8:], uint32(t263))
					t264 := int64(load64(m.memory[uint32(v13):]))
					store64(m.memory[uint32(v9):], uint64(t264))
					if v12 != v11 {
						goto l43
					}
					v9 = v19
					goto l44
				l43:
					v11 = v11 + i32(12)
					v10 = v10 + i32(-12)
					t265 := int32(load32(m.memory[uint32(v9+i32(-24)):]))
					if uint32(v18) < uint32(t265) {
						goto l45
					}
				}
				v9 = v10 + v12
			l44:
				store64(m.memory[int64(uint32(v9))+4:], uint64(v22))
				store32(m.memory[uint32(v9):], uint32(v18))
			}
		l42:
			v14 = v14 + i32(-12)
			v16 = v16 + i32(12)
			v1 = v1 + i32(1)
			if v1 != v4 {
				goto l46
			}
		}
	l41:
		v9 = v19 + i32(-12)
		t266 := v2
		v11 = v7*i32(12) + i32(-12)
		v10 = t266 + v11
		v11 = v0 + v11
	l47:
		{
			t267 := int32(load32(m.memory[uint32(v19):]))
			t268 := v0
			t269 := v19
			t270 := v2
			v12 = t267
			t271 := int32(load32(m.memory[uint32(v2):]))
			t272 := v12
			v13 = t271
			var p273 int32
			if uint32(t272) < uint32(v13) {
				p273 = 1
			}
			v18 = p273
			p274 := t270
			if v18 != 0 {
				p274 = t269
			}
			v1 = p274
			t275 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			store32(m.memory[int64(uint32(t268))+8:], uint32(t275))
			t276 := int64(load64(m.memory[uint32(v1):]))
			store64(m.memory[uint32(v0):], uint64(t276))
			t277 := int32(load32(m.memory[uint32(v10):]))
			t278 := v11
			t279 := v9
			t280 := v10
			v1 = t277
			t281 := int32(load32(m.memory[uint32(v9):]))
			t282 := v1
			v14 = t281
			var p283 int32
			if uint32(t282) < uint32(v14) {
				p283 = 1
			}
			v16 = p283
			p284 := t280
			if v16 != 0 {
				p284 = t279
			}
			v4 = p284
			t285 := int32(load32(m.memory[int64(uint32(v4))+8:]))
			store32(m.memory[int64(uint32(t278))+8:], uint32(t285))
			t286 := int64(load64(m.memory[uint32(v4):]))
			store64(m.memory[uint32(v11):], uint64(t286))
			v11 = v11 + i32(-12)
			v0 = v0 + i32(12)
			t287 := v2
			var p288 int32
			if uint32(v12) >= uint32(v13) {
				p288 = 1
			}
			v2 = t287 + p288*i32(12)
			v19 = v19 + v18*i32(12)
			t290 := v9
			p289 := i32(0)
			if v16 != 0 {
				p289 = i32(-12)
			}
			v9 = t290 + p289
			t292 := v10
			p291 := i32(0)
			if uint32(v1) >= uint32(v14) {
				p291 = i32(-12)
			}
			v10 = t292 + p291
			v17 = v17 + i32(-1)
			if v17 != 0 {
				goto l47
			}
		}
		v9 = v9 + i32(12)
		{
			if v7&i32(1) == 0 {
				goto l48
			}
			t293 := v0
			t294 := v2
			t295 := v19
			var p296 int32
			if uint32(v2) < uint32(v9) {
				p296 = 1
			}
			v7 = p296
			p297 := t295
			if v7 != 0 {
				p297 = t294
			}
			v11 = p297
			t298 := int32(load32(m.memory[int64(uint32(v11))+8:]))
			store32(m.memory[int64(uint32(t293))+8:], uint32(t298))
			t299 := int64(load64(m.memory[uint32(v11):]))
			store64(m.memory[uint32(v0):], uint64(t299))
			t300 := v19
			var p301 int32
			if uint32(v2) >= uint32(v9) {
				p301 = 1
			}
			v19 = t300 + p301*i32(12)
			v2 = v2 + v7*i32(12)
		}
	l48:
		if v2 != v9 {
			goto l49
		}
		if v19 == v10+i32(12) {
			goto l3
		}
	l49:
		m.fn122()
		panic("unreachable")
	}
l3:
	m.g0 = v6 + i32(16)
}
func (m *Module) fn131(v0, v1, v2, v3 int32) int32 {
	var v4, v5, v6 int32
	{
		if uint32(v3) < uint32(i32(8)) {
			goto l0
		}
		t0 := v0
		t1 := v0
		v3 = int32(uint32(v3) >> 3)
		v4 = v3 * i32(48)
		t2 := t1 + v4
		t3 := v0
		v5 = v3 * i32(84)
		t4 := m.fn131(t0, t2, t3+v5, v3)
		v0 = t4
		t5 := m.fn131(v1, v1+v4, v1+v5, v3)
		v1 = t5
		t6 := m.fn131(v2, v2+v4, v2+v5, v3)
		v2 = t6
	}
l0:
	t7 := int32(load32(m.memory[uint32(v0):]))
	t8 := v0
	t9 := v2
	t10 := v1
	v3 = t7
	t11 := int32(load32(m.memory[uint32(v1):]))
	t12 := v3
	v4 = t11
	var p13 int32
	if uint32(t12) < uint32(v4) {
		p13 = 1
	}
	v5 = p13
	t14 := int32(load32(m.memory[uint32(v2):]))
	t15 := v5
	t16 := v4
	v6 = t14
	var p17 int32
	if uint32(t16) < uint32(v6) {
		p17 = 1
	}
	p18 := t10
	if t15^p17 != 0 {
		p18 = t9
	}
	t19 := v5
	var p20 int32
	if uint32(v3) < uint32(v6) {
		p20 = 1
	}
	p21 := p18
	if t19^p20 != 0 {
		p21 = t8
	}
	return p21
}
