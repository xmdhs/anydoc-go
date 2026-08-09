package core

import (
	"math/bits"
)

func (m *Module) fn1707(v0, v1, v2, v3, v4, v5, v6, v7 int32) {
	var v8, v9, v10 int32
	t0 := m.g0
	v8 = t0 - i32(32)
	m.g0 = v8
	m.memory[int64(uint32(v1))+24] = byte(i32(9))
	{
		{
			if v6 != 0 {
				m.fn1706(v0, v1, v2, v3, v4, v5)
				goto l4
			}
			m.memory[int64(uint32(v8))+19] = byte(v7)
			v7 = i32(0)
			m.fn1710(v8+i32(20), v1, v8+i32(19), i32(1), v4, v5, i32(0))
			t1 := int32(load32(m.memory[int64(uint32(v8))+28:]))
			v6 = t1
			t2 := int32(load16(m.memory[int64(uint32(v8))+25:]))
			v9 = t2
			v10 = i32(2)
			t3 := int32(m.memory[int64(uint32(v8))+24])
			switch t3 {
			case 1:
				m.fn256(i32(1155344), i32(39), i32(1155296))
				panic("unreachable")
			case 2:
				goto l3
			default:
				goto l1
			}
		}
	l1:
		m.fn212(v8+i32(8), v6, v4, v5, i32(1155296))
		t4 := int32(load32(m.memory[int64(uint32(v8))+8:]))
		t5 := int32(load32(m.memory[int64(uint32(v8))+12:]))
		m.fn1706(v8+i32(20), v1, v2, v3, t4, t5)
		t6 := int32(load32(m.memory[int64(uint32(v8))+28:]))
		v6 = t6 + v6
		t7 := int32(load32(m.memory[int64(uint32(v8))+20:]))
		v7 = t7
		t8 := int32(load16(m.memory[int64(uint32(v8))+25:]))
		v9 = t8
		t9 := int32(m.memory[int64(uint32(v8))+24])
		v10 = t9
	}
l3:
	store16(m.memory[int64(uint32(v0))+5:], uint16(v9))
	m.memory[int64(uint32(v0))+4] = byte(v10)
	store32(m.memory[int64(uint32(v0))+8:], uint32(v6))
	store32(m.memory[uint32(v0):], uint32(v7))
l4:
	m.g0 = v8 + i32(32)
}
func (m *Module) fn1708(v0, v1, v2, v3, v4, v5, v6 int32) {
	var v7 int32
	t0 := m.g0
	v7 = t0 - i32(32)
	m.g0 = v7
	m.fn148(v7+i32(8), v6, v2, v3, i32(1155296))
	t1 := int32(load32(m.memory[int64(uint32(v7))+8:]))
	t2 := int32(load32(m.memory[int64(uint32(v7))+12:]))
	m.fn1706(v7+i32(20), v1, t1, t2, v4, v5)
	t3 := int32(m.memory[int64(uint32(v7))+26])
	m.memory[int64(uint32(v0))+6] = byte(t3)
	t4 := int32(load16(m.memory[int64(uint32(v7))+24:]))
	store16(m.memory[int64(uint32(v0))+4:], uint16(t4))
	t5 := int32(load32(m.memory[int64(uint32(v7))+20:]))
	v3 = t5
	t6 := int32(load32(m.memory[int64(uint32(v7))+28:]))
	store32(m.memory[int64(uint32(v0))+8:], uint32(t6))
	store32(m.memory[uint32(v0):], uint32(v3+v6))
	m.g0 = v7 + i32(32)
}
func (m *Module) fn1709(v0, v1, v2, v3, v4, v5, v6 int32) {
	var v7, v8, v9 int32
	t0 := m.g0
	v7 = t0 - i32(32)
	m.g0 = v7
	m.memory[int64(uint32(v1))+24] = byte(i32(9))
	switch v6 {
	case 0:
		store16(m.memory[int64(uint32(v7))+18:], uint16(i32(48111)))
		v8 = i32(0)
		v9 = i32(2)
		m.fn1710(v7+i32(20), v1, v7+i32(18), i32(2), v4, v5, i32(0))
		t1 := int32(load32(m.memory[int64(uint32(v7))+28:]))
		v6 = t1
		{
			t2 := int32(m.memory[int64(uint32(v7))+24])
			switch t2 {
			default:
				m.fn212(v7+i32(8), v6, v4, v5, i32(1155296))
				t3 := int32(load32(m.memory[int64(uint32(v7))+8:]))
				t4 := int32(load32(m.memory[int64(uint32(v7))+12:]))
				m.fn1706(v7+i32(20), v1, v2, v3, t3, t4)
				t5 := int32(load32(m.memory[int64(uint32(v7))+28:]))
				v6 = t5 + v6
				t6 := int32(load32(m.memory[int64(uint32(v7))+20:]))
				v8 = t6
				t7 := int32(load16(m.memory[int64(uint32(v7))+25:]))
				v5 = t7
				t8 := int32(m.memory[int64(uint32(v7))+24])
				v9 = t8
				goto l6
			case 1:
				m.fn256(i32(1155344), i32(39), i32(1155296))
				panic("unreachable")
			case 2:
				t9 := int32(load16(m.memory[int64(uint32(v7))+25:]))
				v5 = t9
				t10 := int32(load32(m.memory[int64(uint32(v7))+20:]))
				if t10 == i32(1) {
					goto l7
				}
				goto l6
			}
		}
	case 1:
		m.fn1707(v0, v1, v2, v3, v4, v5, i32(0), i32(239))
		goto l8
	default:
		m.fn1706(v0, v1, v2, v3, v4, v5)
		goto l8
	}
l7:
	m.memory[int64(uint32(v1))+24] = byte(i32(8))
l6:
	store16(m.memory[int64(uint32(v0))+5:], uint16(v5))
	m.memory[int64(uint32(v0))+4] = byte(v9)
	store32(m.memory[int64(uint32(v0))+8:], uint32(v6))
	store32(m.memory[uint32(v0):], uint32(v8))
l8:
	m.g0 = v7 + i32(32)
}
func (m *Module) fn1710(v0, v1, v2, v3, v4, v5, v6 int32) {
	var v7, v8, v9, v10, v11, v12, v13, v14, v15, v16, v17, v18, v19, v20, v21, v22, v23, v24, v25, v26 int32
	t0 := m.g0
	v7 = t0 - i32(352)
	m.g0 = v7
	{
		t1 := int32(m.memory[uint32(v1)])
		switch t1 {
		default:
			store32(m.memory[int64(uint32(v7))+344:], uint32(v5))
			store32(m.memory[int64(uint32(v7))+340:], uint32(v4))
			v6 = v3 + i32(-1)
			t2 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			v8 = t2
			v9 = i32(0)
			v10 = i32(0)
		l36:
			{
				m.fn148(v7+i32(16), v10, v2, v3, i32(1155464))
				t3 := int32(load32(m.memory[int64(uint32(v7))+16:]))
				v11 = t3
				t4 := int32(load32(m.memory[int64(uint32(v7))+20:]))
				v1 = t4
				m.fn212(v7+i32(8), v9, v4, v5, i32(1155480))
				t5 := int32(load32(m.memory[int64(uint32(v7))+12:]))
				v12 = t5
				t6 := v12
				t7 := v1
				var p8 int32
				if uint32(v12) < uint32(v1) {
					p8 = 1
				}
				v13 = p8
				p9 := t7
				if v13 != 0 {
					p9 = t6
				}
				v14 = p9
				v1 = i32(0)
				{
					{
						t10 := int32(load32(m.memory[int64(uint32(v7))+8:]))
						t11 := v11
						v15 = t10
						if (t11^v15)&i32(3) != 0 {
							goto l11
						}
						v1 = i32(0)
						v16 = (i32(0) - v11) & i32(3)
						if uint32(v16|i32(8)) > uint32(v14) {
							goto l11
						}
						v1 = i32(0)
					l18:
						if v16 != v1 {
							t18 := int32(int8(m.memory[uint32(v11+v1)]))
							v12 = t18
							if v12 < i32(0) {
								goto l16
							}
							m.memory[uint32(v15+v1)] = byte(v12)
							v1 = v1 + i32(1)
							goto l18
						}
						v17 = v14 + i32(-8)
					l17:
						{
							v16 = v11 + v1
							t12 := int32(load32(m.memory[uint32(v16):]))
							v12 = t12
							v18 = v15 + v1
							t13 := int32(load32(m.memory[uint32(v16+i32(4)):]))
							t14 := v18 + i32(4)
							v16 = t13
							store32(m.memory[uint32(t14):], uint32(v16))
							store32(m.memory[uint32(v18):], uint32(v12))
							{
								v16 = v16 & i32(-2139062144)
								t15 := v16
								v12 = v12 & i32(-2139062144)
								if t15|v12 == 0 {
									v1 = v1 + i32(8)
									if uint32(v1) <= uint32(v17) {
										goto l17
									}
									goto l11
								}
								if v12 != 0 {
									goto l14
								}
								v12 = int32(uint32(int32(bits.TrailingZeros32(uint32(v16))))>>3) + i32(4)
								goto l15
							l14:
								v12 = int32(uint32(int32(bits.TrailingZeros32(uint32(v12)))) >> 3)
							l15:
								t16 := v11
								v1 = v12 + v1
								t17 := int32(m.memory[uint32(t16+v1)])
								v12 = t17
								goto l16
							}
						}
					}
				l11:
					p19 := v14
					if uint32(v1) > uint32(v14) {
						p19 = v1
					}
					v16 = p19
				l20:
					{
						if v16 == v1 {
							v18 = v14 + v9
							v11 = v14 + v10
							goto l21
						}
						t20 := int32(int8(m.memory[uint32(v11+v1)]))
						v12 = t20
						if v12 < i32(0) {
							goto l16
						}
						m.memory[uint32(v15+v1)] = byte(v12)
						v1 = v1 + i32(1)
						goto l20
					}
				}
			l16:
				v11 = v1 + v10
				v18 = v1 + v9
				if uint32(v18+i32(2)) < uint32(v5) {
					v9 = v11 + i32(1)
				l34:
					{
						t21 := int32(load16(m.memory[uint32(v8+v12&i32(255)<<1+i32(-256)):]))
						v1 = t21
						if v1 != 0 {
							v5 = v4 + v18
							if uint32(v1) < uint32(i32(2048)) {
								goto l26
							}
							m.memory[uint32(v5+i32(1))] = byte(int32(uint32(v1)>>6)&i32(63) | i32(128))
							v4 = int32(uint32(v1)>>12) | i32(-32)
							v17 = i32(3)
							v12 = i32(2)
							goto l27
						l26:
							v4 = int32(uint32(v1)>>6) | i32(-64)
							v17 = i32(2)
							v12 = i32(1)
						l27:
							m.memory[uint32(v5)] = byte(v4)
							m.memory[uint32(v5+v12)] = byte(v1&i32(63) | i32(128))
							v11 = v17 + v18
							if uint32(v9) < uint32(v3) {
								t22 := int32(load32(m.memory[int64(uint32(v7))+344:]))
								t23 := v11 + i32(2)
								v5 = t22
								if uint32(t23) < uint32(v5) {
									v15 = v2 + v9
									v14 = v6 - v9
									t24 := int32(load32(m.memory[int64(uint32(v7))+340:]))
									v4 = t24
									v16 = v4 + v11
									v1 = i32(0)
								l33:
									{
										t25 := int32(int8(m.memory[uint32(v15+v1)]))
										v12 = t25
										if v12 < i32(0) {
											v18 = v11 + v1
											v9 = v9 + v1 + i32(1)
											goto l34
										}
										m.memory[uint32(v16+v1)] = byte(v12)
										if uint32(v12) > uint32(i32(59)) {
											v10 = v9 + v1 + i32(1)
											v9 = v17 + v18 + v1 + i32(1)
											goto l36
										}
										if v14 != v1 {
											goto l32
										}
										v18 = v11 + v1 + i32(1)
										store32(m.memory[uint32(v0):], uint32(v3))
										m.memory[int64(uint32(v0))+4] = byte(i32(0))
										goto l24
									l32:
										t26 := v11
										v1 = v1 + i32(1)
										v12 = t26 + v1
										if uint32(v12+i32(2)) < uint32(v5) {
											goto l33
										}
									}
									store32(m.memory[uint32(v0):], uint32(v9+v1))
									m.memory[int64(uint32(v0))+4] = byte(i32(1))
									v18 = v12
									goto l24
								}
								store32(m.memory[uint32(v0):], uint32(v9))
								m.memory[int64(uint32(v0))+4] = byte(i32(1))
								v18 = v11
								goto l24
							}
							store32(m.memory[uint32(v0):], uint32(v9))
							m.memory[int64(uint32(v0))+4] = byte(i32(0))
							v18 = v11
							goto l24
						}
						m.memory[int64(uint32(v0))+6] = byte(i32(0))
						store16(m.memory[int64(uint32(v0))+4:], uint16(i32(258)))
						store32(m.memory[uint32(v0):], uint32(v9))
						goto l24
					}
				}
				v13 = i32(1)
				goto l21
			l21:
				store32(m.memory[uint32(v0):], uint32(v11))
				store32(m.memory[int64(uint32(v0))+4:], uint32(v13|(v7+i32(340))&i32(-256)))
				goto l24
			l24:
				store32(m.memory[int64(uint32(v0))+8:], uint32(v18))
				goto l35
			}
		case 10:
			{
				{
					t27 := int32(m.memory[int64(uint32(v1))+7])
					if t27 != 0 {
						goto l37
					}
					t28 := int32(load16(m.memory[int64(uint32(v1))+4:]))
					v19 = t28
					v9 = i32(0)
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
					t29 := int32(load16(m.memory[int64(uint32(v1))+4:]))
					v12 = t29
					if uint32(v12) < uint32(i32(128)) {
						goto l41
					}
					if uint32(v12) < uint32(i32(2048)) {
						m.memory[int64(uint32(v4))+1] = byte(v12&i32(63) | i32(128))
						m.memory[uint32(v4)] = byte(int32(uint32(v12)>>6) | i32(192))
						v9 = i32(2)
						goto l43
					}
					m.memory[int64(uint32(v4))+2] = byte(v12&i32(63) | i32(128))
					m.memory[uint32(v4)] = byte(int32(uint32(v12)>>12) | i32(224))
					m.memory[int64(uint32(v4))+1] = byte(int32(uint32(v12)>>6)&i32(63) | i32(128))
					v9 = i32(3)
					goto l43
				}
			l41:
				m.memory[uint32(v4)] = byte(v12)
				v9 = i32(1)
			l43:
				v19 = i32(0)
				store16(m.memory[int64(uint32(v1))+4:], uint16(i32(0)))
				m.memory[int64(uint32(v1))+7] = byte(i32(0))
			l38:
				t30 := int32(m.memory[int64(uint32(v1))+2])
				v10 = t30
				t31 := int32(m.memory[int64(uint32(v1))+3])
				v20 = t31
				t32 := int32(m.memory[int64(uint32(v1))+6])
				v21 = t32
				v17 = i32(0)
			l106:
				{
					v13 = v10 & i32(1)
					if v13 != 0 {
						goto l44
					}
					if v19&i32(0xffff) != 0 {
						goto l44
					}
					m.fn148(v7+i32(328), v17, v2, v3, i32(1148760))
					t33 := int32(load32(m.memory[int64(uint32(v7))+328:]))
					v8 = t33
					t34 := int32(load32(m.memory[int64(uint32(v7))+332:]))
					v12 = t34
					m.fn212(v7+i32(320), v9, v4, v5, i32(1148776))
					v15 = int32(uint32(v12) >> 1)
					t35 := int32(load32(m.memory[int64(uint32(v7))+324:]))
					v14 = t35
					t36 := int32(load32(m.memory[int64(uint32(v7))+320:]))
					v18 = t36
					{
						if v21&i32(1) != 0 {
							if v15 == 0 {
								goto l44
							}
							v11 = i32(0)
							v12 = i32(0)
							if uint32(v14) < uint32(i32(4)) {
								goto l68
							}
							v12 = v15 + i32(-1)
							t47 := int32(load16(m.memory[uint32(v8+v12<<1):]))
							p48 := v15
							if t47&i32(252) == i32(216) {
								p48 = v12
							}
							v22 = p48
							v24 = v14 + i32(-3)
							v11 = i32(0)
							v23 = i32(0)
						l87:
							{
								if uint32(v22) < uint32(v11) {
									m.fn256(i32(1155088), i32(34), i32(1155124))
									panic("unreachable")
								}
								m.fn212(v7+i32(304), v23, v18, v14, i32(1148792))
								t49 := int32(load32(m.memory[int64(uint32(v7))+308:]))
								v12 = t49
								t50 := v12
								v15 = v22 - v11
								p51 := v15
								if uint32(v12) < uint32(v15) {
									p51 = t50
								}
								v25 = p51
								v16 = v8 + v11<<1
								v12 = i32(0)
								t52 := int32(load32(m.memory[int64(uint32(v7))+304:]))
								v26 = t52
							l72:
								if v25 == v12 {
									goto l70
								}
								{
									t53 := int32(load16(m.memory[uint32(v16):]))
									v15 = t53
									v15 = v15<<8 | int32(uint32(v15)>>8)
									if uint32(v15&i32(0xffff)) > uint32(i32(127)) {
										goto l71
									}
									m.memory[uint32(v26+v12)] = byte(v15)
									v16 = v16 + i32(2)
									v12 = v12 + i32(1)
									goto l72
								}
							l71:
								v11 = v12 + v11
								v12 = v12 + v23
								if uint32(v12) >= uint32(v24) {
									goto l68
								}
								v11 = v11 + i32(1)
							l86:
								{
									{
										v16 = (v15 + i32(10240)) & i32(0xffff)
										if uint32(v16) > uint32(i32(2047)) {
											if uint32(v15&i32(0xffff)) < uint32(i32(2048)) {
												if uint32(v12) < uint32(v14) {
													m.memory[uint32(v18+v12)] = byte(int32(uint32(v15)>>6) | i32(192))
													v25 = v12 + i32(1)
													if uint32(v25) >= uint32(v14) {
														m.fn158(v25, v14, i32(1148936))
														panic("unreachable")
													}
													v26 = i32(2)
													v16 = v15
													goto l78
												}
												m.fn158(v12, v14, i32(1148920))
												panic("unreachable")
											}
											if uint32(v12) >= uint32(v14) {
												m.fn158(v12, v14, i32(1148872))
												panic("unreachable")
											}
											m.memory[uint32(v18+v12)] = byte(int32(uint32(v15&i32(61440))>>12) | i32(224))
											v16 = v12 + i32(1)
											if uint32(v16) >= uint32(v14) {
												m.fn158(v16, v14, i32(1148888))
												panic("unreachable")
											}
											m.memory[uint32(v18+v16)] = byte(int32(uint32(v15)>>6)&i32(63) | i32(128))
											v25 = v12 + i32(2)
											if uint32(v25) >= uint32(v14) {
												m.fn158(v25, v14, i32(1148904))
												panic("unreachable")
											}
											v26 = i32(3)
											v16 = v15
											goto l78
										}
										if uint32(v16) > uint32(i32(1023)) {
											goto l52
										}
										if uint32(v11) >= uint32(v22) {
											goto l52
										}
										t54 := int32(load16(m.memory[uint32(v8+v11<<1):]))
										v16 = t54
										v16 = v16<<8 | int32(uint32(v16)>>8)
										if v16&i32(64512) != i32(56320) {
											goto l52
										}
										if uint32(v12) >= uint32(v14) {
											m.fn158(v12, v14, i32(1148808))
											panic("unreachable")
										}
										t55 := v18 + v12
										v15 = v15&i32(0xffff)<<10 + v16&i32(0xffff) + i32(-56613888)
										m.memory[uint32(t55)] = byte(int32(uint32(v15)>>18) | i32(240))
										v25 = v12 + i32(1)
										if uint32(v25) >= uint32(v14) {
											m.fn158(v25, v14, i32(1148824))
											panic("unreachable")
										}
										m.memory[uint32(v18+v25)] = byte(int32(uint32(v15)>>12)&i32(63) | i32(128))
										v25 = v12 + i32(2)
										if uint32(v25) >= uint32(v14) {
											m.fn158(v25, v14, i32(1148840))
											panic("unreachable")
										}
										m.memory[uint32(v18+v25)] = byte(int32(uint32(v15)>>6)&i32(63) | i32(128))
										v25 = v12 + i32(3)
										if uint32(v25) >= uint32(v14) {
											m.fn158(v25, v14, i32(1148856))
											panic("unreachable")
										}
										v11 = v11 + i32(1)
										v26 = i32(4)
										goto l78
									}
								l78:
									m.memory[uint32(v18+v25)] = byte(v16&i32(63) | i32(128))
									v12 = v26 + v12
									if uint32(v12) >= uint32(v24) {
										goto l68
									}
									if v11 == v22 {
										goto l68
									}
									if uint32(v11) >= uint32(v22) {
										goto l85
									}
									v15 = v11 << 1
									v11 = v11 + i32(1)
									t56 := int32(load16(m.memory[uint32(v8+v15):]))
									v15 = t56
									v15 = v15<<8 | int32(uint32(v15)>>8)
									if uint32(v15&i32(0xffff)) > uint32(i32(127)) {
										goto l86
									}
								}
								m.memory[uint32(v18+v12)] = byte(v15)
								v23 = v12 + i32(1)
								goto l87
							l85:
							}
							m.fn256(i32(1155040), i32(30), i32(1155072))
							panic("unreachable")
						}
						if v15 == 0 {
							goto l44
						}
						v12 = i32(0)
						v16 = i32(0)
						if uint32(v14) < uint32(i32(4)) {
							goto l46
						}
						v12 = v15 + i32(-1)
						t37 := int32(load16(m.memory[uint32(v8+v12<<1):]))
						p38 := v15
						if t37&i32(64512) == i32(55296) {
							p38 = v12
						}
						v22 = p38
						v23 = v14 + i32(-3)
						v11 = i32(0)
						v24 = i32(0)
					l67:
						{
							if uint32(v22) < uint32(v11) {
								m.fn256(i32(1155088), i32(34), i32(1155124))
								panic("unreachable")
							}
							m.fn212(v7+i32(312), v24, v18, v14, i32(1148792))
							t39 := int32(load32(m.memory[int64(uint32(v7))+316:]))
							v12 = t39
							t40 := v12
							v15 = v22 - v11
							p41 := v15
							if uint32(v12) < uint32(v15) {
								p41 = t40
							}
							v25 = p41
							v16 = v8 + v11<<1
							v12 = i32(0)
							t42 := int32(load32(m.memory[int64(uint32(v7))+312:]))
							v26 = t42
						l50:
							if v25 == v12 {
								goto l48
							}
							{
								t43 := int32(load16(m.memory[uint32(v16):]))
								v15 = t43
								if uint32(v15) > uint32(i32(127)) {
									goto l49
								}
								m.memory[uint32(v26+v12)] = byte(v15)
								v16 = v16 + i32(2)
								v12 = v12 + i32(1)
								goto l50
							}
						l49:
							v16 = v12 + v11
							v12 = v12 + v24
							if uint32(v12) >= uint32(v23) {
								goto l46
							}
							v11 = v16 + i32(1)
						l66:
							{
								{
									v16 = (v15 + i32(10240)) & i32(0xffff)
									if uint32(v16) > uint32(i32(2047)) {
										if uint32(v15&i32(0xffff)) < uint32(i32(2048)) {
											if uint32(v12) < uint32(v14) {
												m.memory[uint32(v18+v12)] = byte(int32(uint32(v15)>>6) | i32(192))
												v26 = v12 + i32(1)
												if uint32(v26) >= uint32(v14) {
													m.fn158(v26, v14, i32(1148936))
													panic("unreachable")
												}
												v24 = i32(2)
												goto l62
											}
											m.fn158(v12, v14, i32(1148920))
											panic("unreachable")
										}
										if uint32(v12) >= uint32(v14) {
											m.fn158(v12, v14, i32(1148872))
											panic("unreachable")
										}
										m.memory[uint32(v18+v12)] = byte(int32(uint32(v15&i32(61440))>>12) | i32(224))
										v16 = v12 + i32(1)
										if uint32(v16) >= uint32(v14) {
											m.fn158(v16, v14, i32(1148888))
											panic("unreachable")
										}
										m.memory[uint32(v18+v16)] = byte(int32(uint32(v15)>>6)&i32(63) | i32(128))
										v26 = v12 + i32(2)
										if uint32(v26) >= uint32(v14) {
											m.fn158(v26, v14, i32(1148904))
											panic("unreachable")
										}
										v24 = i32(3)
										goto l62
									}
									if uint32(v16) > uint32(i32(1023)) {
										goto l52
									}
									if uint32(v11) >= uint32(v22) {
										goto l52
									}
									t44 := int32(load16(m.memory[uint32(v8+v11<<1):]))
									v25 = t44
									if v25&i32(64512) != i32(56320) {
										goto l52
									}
									if uint32(v12) >= uint32(v14) {
										m.fn158(v12, v14, i32(1148808))
										panic("unreachable")
									}
									t45 := v18 + v12
									v15 = v15&i32(0xffff)<<10 + v25 + i32(-56613888)
									m.memory[uint32(t45)] = byte(int32(uint32(v15)>>18) | i32(240))
									v16 = v12 + i32(1)
									if uint32(v16) >= uint32(v14) {
										m.fn158(v16, v14, i32(1148824))
										panic("unreachable")
									}
									m.memory[uint32(v18+v16)] = byte(int32(uint32(v15)>>12)&i32(63) | i32(128))
									v16 = v12 + i32(2)
									if uint32(v16) >= uint32(v14) {
										m.fn158(v16, v14, i32(1148840))
										panic("unreachable")
									}
									m.memory[uint32(v18+v16)] = byte(int32(uint32(v15)>>6)&i32(63) | i32(128))
									v26 = v12 + i32(3)
									if uint32(v26) >= uint32(v14) {
										m.fn158(v26, v14, i32(1148856))
										panic("unreachable")
									}
									v16 = v11 + i32(1)
									v24 = i32(4)
									goto l57
								}
							l62:
								v25 = v15
								v16 = v11
							l57:
								m.memory[uint32(v18+v26)] = byte(v25&i32(63) | i32(128))
								v12 = v24 + v12
								if uint32(v12) >= uint32(v23) {
									goto l46
								}
								if v16 == v22 {
									goto l46
								}
								if uint32(v16) >= uint32(v22) {
									goto l65
								}
								v11 = v16 + i32(1)
								t46 := int32(load16(m.memory[uint32(v8+v16<<1):]))
								v15 = t46
								if uint32(v15) > uint32(i32(127)) {
									goto l66
								}
							}
							m.memory[uint32(v18+v12)] = byte(v15)
							v24 = v12 + i32(1)
							goto l67
						l65:
						}
						m.fn256(i32(1155040), i32(30), i32(1155072))
						panic("unreachable")
					}
				l52:
					m.memory[int64(uint32(v0))+6] = byte(i32(0))
					store16(m.memory[int64(uint32(v0))+4:], uint16(i32(514)))
					v9 = v12 + v9
					v12 = v11<<1 + v17
					goto l88
				l70:
					v12 = v25 + v23
					v11 = v25 + v11
				l68:
					v9 = v12 + v9
					v12 = v11<<1 + v17
					goto l89
				l48:
					v12 = v25 + v24
					v16 = v25 + v11
				l46:
					v9 = v12 + v9
					v12 = v16<<1 + v17
					goto l89
				}
			l44:
				v12 = v17
			l89:
				if uint32(v12) < uint32(v3) {
					v11 = v9 + i32(3)
					if uint32(v11) < uint32(v5) {
						v17 = v12 + i32(1)
						t59 := int32(m.memory[uint32(v2+v12)])
						v12 = t59
						{
							if v13 == 0 {
								m.memory[int64(uint32(v1))+3] = byte(v12)
								m.memory[int64(uint32(v1))+2] = byte(i32(1))
								v20 = v12
								goto l98
							}
							m.memory[int64(uint32(v1))+2] = byte(i32(0))
							v12 = v20<<8 | v12&i32(255)
							p60 := v12<<8 | v20&i32(255)
							if v21&i32(1) != 0 {
								p60 = v12
							}
							v12 = p60
							v15 = v12 & i32(64512)
							if v15 == i32(55296) {
								store16(m.memory[int64(uint32(v1))+4:], uint16(v12))
								if v19&i32(0xffff) != 0 {
									m.memory[int64(uint32(v0))+6] = byte(i32(2))
									store16(m.memory[int64(uint32(v0))+4:], uint16(i32(514)))
									v12 = v17
									goto l88
								}
								v19 = v12
								goto l98
							}
							if v15 != i32(56320) {
								if v19&i32(0xffff) != 0 {
									m.memory[int64(uint32(v1))+7] = byte(i32(1))
									store16(m.memory[int64(uint32(v1))+4:], uint16(v12))
									m.memory[int64(uint32(v0))+6] = byte(i32(2))
									store16(m.memory[int64(uint32(v0))+4:], uint16(i32(514)))
									v12 = v17
									goto l88
								}
								v15 = v4 + v9
								v16 = v12 & i32(0xffff)
								if uint32(v16) < uint32(i32(128)) {
									goto l101
								}
								if uint32(v16) < uint32(i32(2048)) {
									m.memory[uint32(v15+i32(1))] = byte(v12&i32(63) | i32(128))
									m.memory[uint32(v15)] = byte(int32(uint32(v12)>>6) | i32(192))
									v9 = v9 + i32(2)
									v19 = i32(0)
									goto l98
								}
								m.memory[uint32(v15+i32(2))] = byte(v12&i32(63) | i32(128))
								m.memory[uint32(v15+i32(1))] = byte(int32(uint32(v12)>>6)&i32(63) | i32(128))
								m.memory[uint32(v15)] = byte(int32(uint32(v12&i32(61440))>>12) | i32(224))
								v19 = i32(0)
								v9 = v11
								goto l98
							}
							v15 = v19 & i32(0xffff)
							if v15 != 0 {
								v19 = i32(0)
								store16(m.memory[int64(uint32(v1))+4:], uint16(i32(0)))
								v11 = v4 + v9
								m.memory[uint32(v11+i32(3))] = byte(v12&i32(63) | i32(128))
								t61 := v11
								v12 = v15<<10 + v12&i32(0xffff) + i32(-56613888)
								m.memory[uint32(t61)] = byte(int32(uint32(v12)>>18) | i32(240))
								m.memory[uint32(v11+i32(2))] = byte(int32(uint32(v12)>>6)&i32(63) | i32(128))
								m.memory[uint32(v11+i32(1))] = byte(int32(uint32(v12)>>12)&i32(63) | i32(128))
								v9 = v9 + i32(4)
								goto l98
							}
							m.memory[int64(uint32(v0))+6] = byte(i32(0))
							store16(m.memory[int64(uint32(v0))+4:], uint16(i32(514)))
							v12 = v17
							goto l88
						}
					}
					m.memory[int64(uint32(v0))+4] = byte(i32(1))
					goto l88
				}
				{
					if v6 == 0 {
						goto l91
					}
					t57 := v10
					var p58 int32
					if v19&i32(0xffff) != i32(0) {
						p58 = 1
					}
					if (t57|p58)&i32(1) != 0 {
						if uint32(v9+i32(2)) < uint32(v5) {
							if v19&i32(0xffff) != 0 {
								store16(m.memory[int64(uint32(v1))+4:], uint16(i32(0)))
								if v10&i32(1) == 0 {
									m.memory[int64(uint32(v0))+6] = byte(i32(0))
									store16(m.memory[int64(uint32(v0))+4:], uint16(i32(514)))
									goto l88
								}
								m.memory[int64(uint32(v0))+6] = byte(i32(0))
								store16(m.memory[int64(uint32(v0))+4:], uint16(i32(770)))
								m.memory[int64(uint32(v1))+2] = byte(i32(0))
								goto l88
							}
							m.memory[int64(uint32(v0))+6] = byte(i32(0))
							store16(m.memory[int64(uint32(v0))+4:], uint16(i32(258)))
							m.memory[int64(uint32(v1))+2] = byte(i32(0))
							goto l88
						}
						m.memory[int64(uint32(v0))+4] = byte(i32(1))
						goto l40
					}
				}
			l91:
				m.memory[int64(uint32(v0))+4] = byte(i32(0))
				goto l88
			l101:
				m.memory[uint32(v15)] = byte(v12)
				v9 = v9 + i32(1)
				v19 = i32(0)
			l98:
				v10 = v10 ^ i32(1)
				goto l106
			}
		l40:
			v12 = i32(0)
			v9 = i32(0)
		l88:
			store32(m.memory[int64(uint32(v0))+8:], uint32(v9))
			store32(m.memory[uint32(v0):], uint32(v12))
			goto l35
		case 9:
			v16 = i32(0)
			v1 = i32(0)
			v12 = i32(0)
		l111:
			if v3 != v12 {
				goto l107
			}
			v12 = v3
			goto l108
		l107:
			if uint32(v1+i32(2)) < uint32(v5) {
				v15 = i32(1)
				{
					t62 := int32(int8(m.memory[uint32(v2+v12)]))
					v11 = t62
					if v11 > i32(-1) {
						goto l110
					}
					v15 = v4 + v1
					m.memory[uint32(v15+i32(2))] = byte(v11 & i32(191))
					m.memory[uint32(v15+i32(1))] = byte(int32(uint32(v11&i32(192))>>6) | i32(156))
					v11 = i32(239)
					v15 = i32(3)
				}
			l110:
				m.memory[uint32(v4+v1)] = byte(v11)
				v12 = v12 + i32(1)
				v1 = v15 + v1
				goto l111
			}
			v16 = i32(1)
		l108:
			store32(m.memory[int64(uint32(v0))+8:], uint32(v1))
			store32(m.memory[uint32(v0):], uint32(v12))
			m.memory[int64(uint32(v0))+4] = byte(v16)
			goto l35
		case 8:
			{
				if v3 == 0 {
					goto l112
				}
				t63 := int32(m.memory[int64(uint32(v1))+1])
				if t63&i32(1) == 0 {
					goto l113
				}
			}
		l112:
			m.memory[int64(uint32(v0))+4] = byte(i32(0))
			goto l114
		l113:
			if uint32(v5) < uint32(i32(3)) {
				goto l115
			}
			m.memory[int64(uint32(v0))+6] = byte(i32(0))
			store16(m.memory[int64(uint32(v0))+4:], uint16(i32(258)))
			v3 = i32(1)
			m.memory[int64(uint32(v1))+1] = byte(i32(1))
			goto l114
		l115:
			m.memory[int64(uint32(v0))+4] = byte(i32(1))
			v3 = i32(0)
		l114:
			store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
			store32(m.memory[uint32(v0):], uint32(v3))
			goto l35
		case 7:
			store32(m.memory[int64(uint32(v7))+348:], uint32(i32(0)))
			store32(m.memory[int64(uint32(v7))+344:], uint32(v5))
			store32(m.memory[int64(uint32(v7))+340:], uint32(v4))
			v18 = i32(0)
			v9 = i32(0)
			{
				t64 := int32(m.memory[int64(uint32(v1))+1])
				if t64 == 0 {
					goto l116
				}
				m.memory[int64(uint32(v1))+1] = byte(i32(0))
				{
					if v3 == 0 {
						if v6 != 0 {
							m.memory[int64(uint32(v0))+6] = byte(i32(0))
							store16(m.memory[int64(uint32(v0))+4:], uint16(i32(258)))
							store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
							store32(m.memory[uint32(v0):], uint32(i32(0)))
							goto l35
						}
						store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
						store32(m.memory[uint32(v0):], uint32(i32(0)))
						m.memory[int64(uint32(v0))+4] = byte(i32(0))
						goto l35
					}
					if uint32(v5) > uint32(i32(2)) {
						t65 := int32(int8(m.memory[uint32(v2)]))
						v12 = t65
						{
							t66 := int32(m.memory[int64(uint32(v1))+2])
							v11 = t66
							if uint32(v11&i32(255)) > uint32(i32(31)) {
								v15 = v11 + i32(-32)
								v16 = v12 + i32(95)
								if uint32(v16&i32(255)) < uint32(i32(94)) {
									t67 := v15 & i32(255) * i32(94)
									v14 = v16 & i32(255)
									v12 = t67 + v14
									v15 = v12 + i32(-1410)
									if uint32(v15) < uint32(i32(2350)) {
										t82 := int32(load16(m.memory[int64(uint32(v15<<1))+1228728:]))
										t83 := v4
										v12 = t82
										m.memory[uint32(t83)] = byte(int32(uint32(v12)>>12) | i32(224))
										m.memory[int64(uint32(v4))+1] = byte(int32(uint32(v12)>>6)&i32(63) | i32(128))
										goto l140
									}
									if uint32(v12) < uint32(i32(165)) {
										v11 = i32(1)
										{
											t68 := int32(load16(m.memory[int64(uint32(v12<<1))+1236264:]))
											v12 = t68
											if uint32(v12) < uint32(i32(2048)) {
												m.memory[uint32(v4)] = byte(int32(uint32(v12)>>6) | i32(192))
												v18 = i32(2)
												goto l141
											}
											m.memory[uint32(v4)] = byte(int32(uint32(v12)>>12) | i32(224))
											m.memory[int64(uint32(v4))+1] = byte(int32(uint32(v12)>>6)&i32(63) | i32(128))
											goto l140
										}
									}
									v15 = v12 + i32(-3854)
									if uint32(v15) < uint32(i32(4888)) {
										t69 := int32(load16(m.memory[int64(uint32(v15<<1))+1218764:]))
										t70 := v4
										v12 = t69
										m.memory[uint32(t70)] = byte(int32(uint32(v12)>>12) | i32(224))
										m.memory[int64(uint32(v4))+1] = byte(int32(uint32(v12)>>6)&i32(63) | i32(128))
										goto l140
									}
									v15 = v11 & i32(255)
									if v15 != i32(39) {
										goto l132
									}
									if uint32(v16&i32(255)) < uint32(i32(15)) {
										if v16&i32(13) != i32(4) {
											v11 = i32(1)
											t71 := int32(load16(m.memory[int64(uint32(v14<<1))+1244128:]))
											t72 := v4
											v12 = t71
											m.memory[uint32(t72)] = byte(int32(uint32(v12)>>6) | i32(192))
											v18 = i32(2)
											goto l141
										}
										m.memory[int64(uint32(v0))+6] = byte(i32(0))
										store16(m.memory[int64(uint32(v0))+4:], uint16(i32(514)))
										store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
										store32(m.memory[uint32(v0):], uint32(i32(1)))
										goto l35
									}
								l132:
									if v15 != i32(40) {
										goto l134
									}
									if uint32(v16&i32(255)) < uint32(i32(16)) {
										v11 = i32(1)
										t73 := int32(load16(m.memory[int64(uint32(v14<<1))+1244096:]))
										t74 := v4
										v12 = t73
										m.memory[uint32(t74)] = byte(int32(uint32(v12)>>6) | i32(192))
										v18 = i32(2)
										goto l141
									}
								l134:
									if v11&i32(255) != i32(37) {
										goto l136
									}
									if uint32(v16&i32(255)) < uint32(i32(68)) {
										t75 := int32(load16(m.memory[int64(uint32(v14<<1))+1155746:]))
										t76 := v4
										v12 = t75
										m.memory[uint32(t76)] = byte(int32(uint32(v12)>>12) | i32(224))
										m.memory[int64(uint32(v4))+1] = byte(int32(uint32(v12)>>6)&i32(63) | i32(128))
										goto l140
									}
								l136:
									v12 = v12 + i32(-188)
									if uint32(v12) < uint32(i32(927)) {
										m.fn1711(v7+i32(280), i32(1251350), i32(77), v12)
										t77 := int32(load32(m.memory[int64(uint32(v7))+284:]))
										v11 = t77
										{
											{
												t78 := int32(load32(m.memory[int64(uint32(v7))+280:]))
												if t78 != i32(1) {
													if uint32(v11) > uint32(i32(76)) {
														m.fn158(v11, i32(77), i32(1236596))
														panic("unreachable")
													}
													t79 := int32(load16(m.memory[int64(uint32(v11<<1))+1272488:]))
													v12 = t79
													goto l146
												}
												v11 = v11 + i32(-1)
												if uint32(v11) < uint32(i32(77)) {
													goto l144
												}
												m.fn158(v11, i32(77), i32(1236612))
												panic("unreachable")
											}
										l144:
											v11 = v11 << 1
											t80 := int32(load16(m.memory[int64(uint32(v11))+1272488:]))
											t81 := int32(load16(m.memory[int64(uint32(v11))+1251350:]))
											v12 = t80 + v12 - t81
										}
									l146:
										v11 = v12 & i32(0xffff)
										if uint32(v11) < uint32(i32(128)) {
											m.memory[int64(uint32(v0))+6] = byte(i32(0))
											store16(m.memory[int64(uint32(v0))+4:], uint16(i32(514)))
											store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
											store32(m.memory[uint32(v0):], uint32(i32(1)))
											goto l35
										}
										if uint32(v11) < uint32(i32(2048)) {
											m.memory[uint32(v4)] = byte(int32(uint32(v12)>>6) | i32(192))
											v18 = i32(2)
											v11 = i32(1)
											goto l141
										}
										m.memory[int64(uint32(v4))+1] = byte(int32(uint32(v12)>>6)&i32(63) | i32(128))
										m.memory[uint32(v4)] = byte(int32(uint32(v12&i32(61440))>>12) | i32(224))
										goto l140
									}
									m.memory[int64(uint32(v0))+6] = byte(i32(0))
									store16(m.memory[int64(uint32(v0))+4:], uint16(i32(514)))
									store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
									store32(m.memory[uint32(v0):], uint32(i32(1)))
									goto l35
								}
								if uint32((v12+i32(127))&i32(255)) < uint32(i32(32)) {
									v11 = v12 + i32(-77)
									goto l127
								}
								if uint32((v12+i32(-97))&i32(255)) < uint32(i32(26)) {
									v11 = v12 + i32(-71)
									goto l127
								}
								v11 = v12 + i32(-65)
								if uint32(v11&i32(255)) < uint32(i32(26)) {
									goto l127
								}
								m.memory[int64(uint32(v0))+4] = byte(i32(2))
								if v12 > i32(-1) {
									store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
									store32(m.memory[uint32(v0):], uint32(i32(0)))
									store16(m.memory[int64(uint32(v0))+5:], uint16(i32(1)))
									goto l35
								}
								store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
								store32(m.memory[uint32(v0):], uint32(i32(1)))
								store16(m.memory[int64(uint32(v0))+5:], uint16(i32(2)))
								goto l35
							}
							if uint32((v12+i32(127))&i32(255)) < uint32(i32(126)) {
								v15 = v12 + i32(-77)
								goto l122
							}
							if uint32((v12+i32(-97))&i32(255)) < uint32(i32(26)) {
								v15 = v12 + i32(-71)
								goto l122
							}
							v15 = v12 + i32(-65)
							if uint32(v15&i32(255)) < uint32(i32(26)) {
								goto l122
							}
							m.memory[int64(uint32(v0))+4] = byte(i32(2))
							if v12 > i32(-1) {
								store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
								store32(m.memory[uint32(v0):], uint32(i32(0)))
								store16(m.memory[int64(uint32(v0))+5:], uint16(i32(1)))
								goto l35
							}
							store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
							store32(m.memory[uint32(v0):], uint32(i32(1)))
							store16(m.memory[int64(uint32(v0))+5:], uint16(i32(2)))
							goto l35
						}
					}
					store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
					store32(m.memory[uint32(v0):], uint32(i32(0)))
					m.memory[int64(uint32(v0))+4] = byte(i32(1))
					goto l35
				l127:
					v11 = v15&i32(255)*i32(84) + v11&i32(255)
					if uint32(v11) < uint32(i32(3126)) {
						m.fn1711(v7+i32(288), i32(1260824), i32(535), v11)
						t84 := int32(load32(m.memory[int64(uint32(v7))+292:]))
						v12 = t84
						{
							{
								t85 := int32(load32(m.memory[int64(uint32(v7))+288:]))
								if t85 != i32(1) {
									if uint32(v12) > uint32(i32(534)) {
										m.fn158(v12, i32(535), i32(1236596))
										panic("unreachable")
									}
									t86 := int32(load16(m.memory[int64(uint32(v12<<1))+1253764:]))
									v12 = t86
									v11 = int32(uint32(v12) >> 12)
									v15 = int32(uint32(v12) >> 6)
									goto l155
								}
								v12 = v12 + i32(-1)
								if uint32(v12) < uint32(i32(535)) {
									goto l153
								}
								m.fn158(v12, i32(535), i32(1236612))
								panic("unreachable")
							}
						l153:
							t87 := v11
							v12 = v12 << 1
							t88 := int32(load16(m.memory[int64(uint32(v12))+1260824:]))
							t89 := int32(load16(m.memory[int64(uint32(v12))+1253764:]))
							v12 = t87 - t88 + t89
							v15 = v12 & i32(0xffff)
							v11 = int32(uint32(v15) >> 12)
							v15 = int32(uint32(v15) >> 6)
						}
					l155:
						m.memory[uint32(v4)] = byte(v11 | i32(224))
						m.memory[int64(uint32(v4))+1] = byte(v15&i32(63) | i32(128))
						goto l140
					}
					m.memory[int64(uint32(v0))+4] = byte(i32(2))
					if v12 > i32(-1) {
						store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
						store32(m.memory[uint32(v0):], uint32(i32(0)))
						store16(m.memory[int64(uint32(v0))+5:], uint16(i32(1)))
						goto l35
					}
					store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
					store32(m.memory[uint32(v0):], uint32(i32(1)))
					store16(m.memory[int64(uint32(v0))+5:], uint16(i32(2)))
					goto l35
				l122:
					t90 := v7 + i32(296)
					v11 = v11&i32(255)*i32(178) + v15&i32(255)
					m.fn1711(t90, i32(1254834), i32(1079), v11)
					t91 := int32(load32(m.memory[int64(uint32(v7))+300:]))
					v12 = t91
					{
						{
							t92 := int32(load32(m.memory[int64(uint32(v7))+296:]))
							if t92 != i32(1) {
								if uint32(v12) > uint32(i32(1078)) {
									m.fn158(v12, i32(1079), i32(1236596))
									panic("unreachable")
								}
								t93 := int32(load16(m.memory[int64(uint32(v12<<1))+1251540:]))
								v12 = t93
								v11 = int32(uint32(v12) >> 12)
								v15 = int32(uint32(v12) >> 6)
								goto l159
							}
							v12 = v12 + i32(-1)
							if uint32(v12) < uint32(i32(1079)) {
								goto l157
							}
							m.fn158(v12, i32(1079), i32(1236612))
							panic("unreachable")
						}
					l157:
						t94 := v11
						v12 = v12 << 1
						t95 := int32(load16(m.memory[int64(uint32(v12))+1254834:]))
						t96 := int32(load16(m.memory[int64(uint32(v12))+1251540:]))
						v12 = t94 - t95 + t96
						v15 = v12 & i32(0xffff)
						v11 = int32(uint32(v15) >> 12)
						v15 = int32(uint32(v15) >> 6)
					}
				l159:
					m.memory[uint32(v4)] = byte(v11 | i32(224))
					m.memory[int64(uint32(v4))+1] = byte(v15&i32(63) | i32(128))
				}
			l140:
				v18 = i32(3)
				v11 = i32(2)
			l141:
				m.memory[uint32(v4+v11)] = byte(v12&i32(63) | i32(128))
				v9 = i32(1)
			}
		l116:
			v19 = i32(2) - v3
			v25 = v2 + i32(1)
		l221:
			{
				m.fn148(v7+i32(272), v9, v2, v3, i32(1155464))
				t97 := int32(load32(m.memory[int64(uint32(v7))+272:]))
				v15 = t97
				t98 := int32(load32(m.memory[int64(uint32(v7))+276:]))
				v12 = t98
				m.fn212(v7+i32(264), v18, v4, v5, i32(1155480))
				t99 := int32(load32(m.memory[int64(uint32(v7))+268:]))
				v11 = t99
				t100 := v11
				t101 := v12
				var p102 int32
				if uint32(v11) < uint32(v12) {
					p102 = 1
				}
				v13 = p102
				p103 := t101
				if v13 != 0 {
					p103 = t100
				}
				v10 = p103
				v12 = i32(0)
				{
					{
						t104 := int32(load32(m.memory[int64(uint32(v7))+264:]))
						t105 := v15
						v16 = t104
						if (t105^v16)&i32(3) != 0 {
							goto l160
						}
						v12 = i32(0)
						v14 = (i32(0) - v15) & i32(3)
						if uint32(v14|i32(8)) > uint32(v10) {
							goto l160
						}
						v12 = i32(0)
					l167:
						if v14 != v12 {
							t112 := int32(int8(m.memory[uint32(v15+v12)]))
							v11 = t112
							if v11 < i32(0) {
								goto l165
							}
							m.memory[uint32(v16+v12)] = byte(v11)
							v12 = v12 + i32(1)
							goto l167
						}
						v8 = v10 + i32(-8)
					l166:
						{
							v14 = v15 + v12
							t106 := int32(load32(m.memory[uint32(v14):]))
							v11 = t106
							v17 = v16 + v12
							t107 := int32(load32(m.memory[uint32(v14+i32(4)):]))
							t108 := v17 + i32(4)
							v14 = t107
							store32(m.memory[uint32(t108):], uint32(v14))
							store32(m.memory[uint32(v17):], uint32(v11))
							{
								v14 = v14 & i32(-2139062144)
								t109 := v14
								v11 = v11 & i32(-2139062144)
								if t109|v11 == 0 {
									v12 = v12 + i32(8)
									if uint32(v12) <= uint32(v8) {
										goto l166
									}
									goto l160
								}
								if v11 != 0 {
									goto l163
								}
								v11 = int32(uint32(int32(bits.TrailingZeros32(uint32(v14))))>>3) + i32(4)
								goto l164
							l163:
								v11 = int32(uint32(int32(bits.TrailingZeros32(uint32(v11)))) >> 3)
							l164:
								t110 := v15
								v12 = v11 + v12
								t111 := int32(m.memory[uint32(t110+v12)])
								v11 = t111
								goto l165
							}
						}
					}
				l160:
					p113 := v10
					if uint32(v12) > uint32(v10) {
						p113 = v12
					}
					v14 = p113
				l169:
					{
						if v14 == v12 {
							v12 = v10 + v18
							v15 = v10 + v9
							goto l170
						}
						t114 := int32(int8(m.memory[uint32(v15+v12)]))
						v11 = t114
						if v11 < i32(0) {
							goto l165
						}
						m.memory[uint32(v16+v12)] = byte(v11)
						v12 = v12 + i32(1)
						goto l169
					}
				}
			l165:
				v15 = v12 + v9
				v12 = v12 + v18
				if uint32(v12+i32(2)) < uint32(v5) {
					v16 = v15 + i32(1)
				l225:
					{
						v5 = v11 + i32(127)
						v15 = v5 & i32(255)
						if uint32(v15) > uint32(i32(125)) {
							m.memory[int64(uint32(v0))+6] = byte(i32(0))
							store16(m.memory[int64(uint32(v0))+4:], uint16(i32(258)))
							goto l176
						}
						if uint32(v16) < uint32(v3) {
							v18 = v16 + i32(1)
							t115 := int32(int8(m.memory[uint32(v2+v16)]))
							v5 = t115
							if uint32(v15) > uint32(i32(31)) {
								v11 = v11 + i32(95)
								v9 = v5 + i32(95)
								v14 = v9 & i32(255)
								if uint32(v14) < uint32(i32(94)) {
									v5 = v11&i32(255)*i32(94) + v14
									v11 = v5 + i32(-1410)
									if uint32(v11) < uint32(i32(2350)) {
										v5 = v4 + v12
										t132 := int32(load16(m.memory[int64(uint32(v11<<1))+1228728:]))
										t133 := v5 + i32(2)
										v4 = t132
										m.memory[uint32(t133)] = byte(v4&i32(63) | i32(128))
										m.memory[uint32(v5)] = byte(int32(uint32(v4)>>12) | i32(224))
										m.memory[uint32(v5+i32(1))] = byte(int32(uint32(v4)>>6)&i32(63) | i32(128))
										goto l199
									}
									if uint32(v5) < uint32(i32(165)) {
										v4 = v4 + v12
										{
											t116 := int32(load16(m.memory[int64(uint32(v5<<1))+1236264:]))
											v5 = t116
											if uint32(v5) < uint32(i32(2048)) {
												m.memory[uint32(v4+i32(1))] = byte(v5&i32(63) | i32(128))
												m.memory[uint32(v4)] = byte(int32(uint32(v5)>>6) | i32(192))
												v5 = i32(2)
												goto l200
											}
											m.memory[uint32(v4+i32(2))] = byte(v5&i32(63) | i32(128))
											m.memory[uint32(v4)] = byte(int32(uint32(v5)>>12) | i32(224))
											m.memory[uint32(v4+i32(1))] = byte(int32(uint32(v5)>>6)&i32(63) | i32(128))
											goto l199
										}
									}
									v11 = v5 + i32(-3854)
									if uint32(v11) < uint32(i32(4888)) {
										v5 = v4 + v12
										t117 := int32(load16(m.memory[int64(uint32(v11<<1))+1218764:]))
										t118 := v5 + i32(2)
										v4 = t117
										m.memory[uint32(t118)] = byte(v4&i32(63) | i32(128))
										m.memory[uint32(v5)] = byte(int32(uint32(v4)>>12) | i32(224))
										m.memory[uint32(v5+i32(1))] = byte(int32(uint32(v4)>>6)&i32(63) | i32(128))
										goto l199
									}
									if v15 != i32(39) {
										goto l190
									}
									if uint32(v14) < uint32(i32(15)) {
										if v9&i32(13) != i32(4) {
											t119 := int32(load16(m.memory[int64(uint32(v14<<1))+1244128:]))
											t120 := v4 + v12
											v5 = t119
											m.memory[uint32(t120)] = byte(int32(uint32(v5)>>6) | i32(192))
											t121 := int32(load32(m.memory[int64(uint32(v7))+340:]))
											m.memory[uint32(t121+v12+i32(1))] = byte(v5&i32(63) | i32(128))
											v5 = i32(2)
											goto l200
										}
										m.memory[int64(uint32(v0))+6] = byte(i32(0))
										store16(m.memory[int64(uint32(v0))+4:], uint16(i32(514)))
										goto l197
									}
								l190:
									if v15 != i32(40) {
										goto l192
									}
									if uint32(v14) < uint32(i32(16)) {
										t122 := int32(load16(m.memory[int64(uint32(v14<<1))+1244096:]))
										t123 := v4 + v12
										v5 = t122
										m.memory[uint32(t123)] = byte(int32(uint32(v5)>>6) | i32(192))
										t124 := int32(load32(m.memory[int64(uint32(v7))+340:]))
										m.memory[uint32(t124+v12+i32(1))] = byte(v5&i32(63) | i32(128))
										v5 = i32(2)
										goto l200
									}
								l192:
									if v15 != i32(37) {
										goto l194
									}
									if uint32(v14) < uint32(i32(68)) {
										v5 = v4 + v12
										t125 := int32(load16(m.memory[int64(uint32(v14<<1))+1155746:]))
										t126 := v5 + i32(2)
										v4 = t125
										m.memory[uint32(t126)] = byte(v4&i32(63) | i32(128))
										m.memory[uint32(v5)] = byte(int32(uint32(v4)>>12) | i32(224))
										m.memory[uint32(v5+i32(1))] = byte(int32(uint32(v4)>>6)&i32(63) | i32(128))
										goto l199
									}
								l194:
									v5 = v5 + i32(-188)
									if uint32(v5) < uint32(i32(927)) {
										m.fn1711(v7+i32(240), i32(1251350), i32(77), v5)
										t127 := int32(load32(m.memory[int64(uint32(v7))+244:]))
										v11 = t127
										{
											{
												t128 := int32(load32(m.memory[int64(uint32(v7))+240:]))
												if t128 != i32(1) {
													if uint32(v11) > uint32(i32(76)) {
														m.fn158(v11, i32(77), i32(1236596))
														panic("unreachable")
													}
													t129 := int32(load16(m.memory[int64(uint32(v11<<1))+1272488:]))
													v5 = t129
													goto l205
												}
												v11 = v11 + i32(-1)
												if uint32(v11) < uint32(i32(77)) {
													goto l203
												}
												m.fn158(v11, i32(77), i32(1236612))
												panic("unreachable")
											}
										l203:
											v11 = v11 << 1
											t130 := int32(load16(m.memory[int64(uint32(v11))+1272488:]))
											t131 := int32(load16(m.memory[int64(uint32(v11))+1251350:]))
											v5 = t130 + v5 - t131
										}
									l205:
										v11 = v5 & i32(0xffff)
										if uint32(v11) < uint32(i32(128)) {
											m.memory[int64(uint32(v0))+6] = byte(i32(0))
											store16(m.memory[int64(uint32(v0))+4:], uint16(i32(514)))
											goto l197
										}
										v4 = v4 + v12
										if uint32(v11) < uint32(i32(2048)) {
											m.memory[uint32(v4+i32(1))] = byte(v5&i32(63) | i32(128))
											m.memory[uint32(v4)] = byte(int32(uint32(v5)>>6) | i32(192))
											v5 = i32(2)
											goto l200
										}
										m.memory[uint32(v4+i32(2))] = byte(v5&i32(63) | i32(128))
										m.memory[uint32(v4+i32(1))] = byte(int32(uint32(v5)>>6)&i32(63) | i32(128))
										m.memory[uint32(v4)] = byte(int32(uint32(v5&i32(61440))>>12) | i32(224))
										goto l199
									}
									m.memory[int64(uint32(v0))+6] = byte(i32(0))
									store16(m.memory[int64(uint32(v0))+4:], uint16(i32(514)))
									goto l197
								}
								if uint32((v5+i32(127))&i32(255)) < uint32(i32(32)) {
									v15 = v5 + i32(-77)
									goto l185
								}
								if uint32((v5+i32(-97))&i32(255)) < uint32(i32(26)) {
									v15 = v5 + i32(-71)
									goto l185
								}
								v15 = v5 + i32(-65)
								if uint32(v15&i32(255)) < uint32(i32(26)) {
									goto l185
								}
								m.memory[int64(uint32(v0))+4] = byte(i32(2))
								if v5 > i32(-1) {
									store32(m.memory[int64(uint32(v0))+8:], uint32(v12))
									store32(m.memory[uint32(v0):], uint32(v16))
									store16(m.memory[int64(uint32(v0))+5:], uint16(i32(1)))
									goto l35
								}
								store32(m.memory[int64(uint32(v0))+8:], uint32(v12))
								store32(m.memory[uint32(v0):], uint32(v18))
								store16(m.memory[int64(uint32(v0))+5:], uint16(i32(2)))
								goto l35
							}
							if uint32((v5+i32(127))&i32(255)) < uint32(i32(126)) {
								v11 = v5 + i32(-77)
								goto l180
							}
							if uint32((v5+i32(-97))&i32(255)) < uint32(i32(26)) {
								v11 = v5 + i32(-71)
								goto l180
							}
							v11 = v5 + i32(-65)
							if uint32(v11&i32(255)) < uint32(i32(26)) {
								goto l180
							}
							m.memory[int64(uint32(v0))+4] = byte(i32(2))
							if v5 > i32(-1) {
								store32(m.memory[int64(uint32(v0))+8:], uint32(v12))
								store32(m.memory[uint32(v0):], uint32(v16))
								store16(m.memory[int64(uint32(v0))+5:], uint16(i32(1)))
								goto l35
							}
							store32(m.memory[int64(uint32(v0))+8:], uint32(v12))
							store32(m.memory[uint32(v0):], uint32(v18))
							store16(m.memory[int64(uint32(v0))+5:], uint16(i32(2)))
							goto l35
						}
						if v6 != 0 {
							m.memory[int64(uint32(v0))+6] = byte(i32(0))
							store16(m.memory[int64(uint32(v0))+4:], uint16(i32(258)))
							goto l176
						}
						m.memory[int64(uint32(v1))+2] = byte(v5)
						m.memory[int64(uint32(v1))+1] = byte(i32(1))
						m.memory[int64(uint32(v0))+4] = byte(i32(0))
						goto l176
					l197:
						store32(m.memory[int64(uint32(v0))+8:], uint32(v12))
						store32(m.memory[uint32(v0):], uint32(v18))
						goto l35
					l185:
						v11 = v11&i32(255)*i32(84) + v15&i32(255)
						if uint32(v11) < uint32(i32(3126)) {
							m.fn1711(v7+i32(248), i32(1260824), i32(535), v11)
							t134 := int32(load32(m.memory[int64(uint32(v7))+252:]))
							v5 = t134
							{
								{
									t135 := int32(load32(m.memory[int64(uint32(v7))+248:]))
									if t135 != i32(1) {
										if uint32(v5) > uint32(i32(534)) {
											m.fn158(v5, i32(535), i32(1236596))
											panic("unreachable")
										}
										t136 := int32(load16(m.memory[int64(uint32(v5<<1))+1253764:]))
										v11 = t136
										v15 = int32(uint32(v11) >> 12)
										v14 = int32(uint32(v11) >> 6)
										goto l213
									}
									v5 = v5 + i32(-1)
									if uint32(v5) < uint32(i32(535)) {
										goto l211
									}
									m.fn158(v5, i32(535), i32(1236612))
									panic("unreachable")
								}
							l211:
								t137 := v11
								v5 = v5 << 1
								t138 := int32(load16(m.memory[int64(uint32(v5))+1260824:]))
								t139 := int32(load16(m.memory[int64(uint32(v5))+1253764:]))
								v11 = t137 - t138 + t139
								v5 = v11 & i32(0xffff)
								v15 = int32(uint32(v5) >> 12)
								v14 = int32(uint32(v5) >> 6)
							}
						l213:
							v5 = v4 + v12
							m.memory[uint32(v5)] = byte(v15 | i32(224))
							m.memory[uint32(v5+i32(2))] = byte(v11&i32(63) | i32(128))
							m.memory[uint32(v5+i32(1))] = byte(v14&i32(63) | i32(128))
							goto l199
						}
						m.memory[int64(uint32(v0))+4] = byte(i32(2))
						if v5 > i32(-1) {
							store32(m.memory[int64(uint32(v0))+8:], uint32(v12))
							store32(m.memory[uint32(v0):], uint32(v16))
							store16(m.memory[int64(uint32(v0))+5:], uint16(i32(1)))
							goto l35
						}
						store32(m.memory[int64(uint32(v0))+8:], uint32(v12))
						store32(m.memory[uint32(v0):], uint32(v18))
						store16(m.memory[int64(uint32(v0))+5:], uint16(i32(2)))
						goto l35
					l180:
						t140 := v7 + i32(256)
						v11 = v15*i32(178) + v11&i32(255)
						m.fn1711(t140, i32(1254834), i32(1079), v11)
						t141 := int32(load32(m.memory[int64(uint32(v7))+260:]))
						v5 = t141
						{
							{
								t142 := int32(load32(m.memory[int64(uint32(v7))+256:]))
								if t142 != i32(1) {
									if uint32(v5) > uint32(i32(1078)) {
										m.fn158(v5, i32(1079), i32(1236596))
										panic("unreachable")
									}
									t143 := int32(load16(m.memory[int64(uint32(v5<<1))+1251540:]))
									v11 = t143
									v15 = int32(uint32(v11) >> 12)
									v14 = int32(uint32(v11) >> 6)
									goto l217
								}
								v5 = v5 + i32(-1)
								if uint32(v5) < uint32(i32(1079)) {
									goto l215
								}
								m.fn158(v5, i32(1079), i32(1236612))
								panic("unreachable")
							}
						l215:
							t144 := v11
							v5 = v5 << 1
							t145 := int32(load16(m.memory[int64(uint32(v5))+1254834:]))
							t146 := int32(load16(m.memory[int64(uint32(v5))+1251540:]))
							v11 = t144 - t145 + t146
							v5 = v11 & i32(0xffff)
							v15 = int32(uint32(v5) >> 12)
							v14 = int32(uint32(v5) >> 6)
						}
					l217:
						v5 = v4 + v12
						m.memory[uint32(v5)] = byte(v15 | i32(224))
						m.memory[uint32(v5+i32(2))] = byte(v11&i32(63) | i32(128))
						m.memory[uint32(v5+i32(1))] = byte(v14&i32(63) | i32(128))
					}
				l199:
					v5 = i32(3)
				l200:
					v14 = v5 + v12
					if uint32(v18) < uint32(v3) {
						t147 := int32(load32(m.memory[int64(uint32(v7))+344:]))
						t148 := v14 + i32(2)
						v5 = t147
						if uint32(t148) < uint32(v5) {
							v8 = v19 + v16
							v10 = v25 + v16
							t149 := int32(load32(m.memory[int64(uint32(v7))+340:]))
							v4 = t149
							v17 = v4 + v14
							v12 = i32(0)
						l224:
							{
								v15 = v16 + v12
								t150 := int32(int8(m.memory[uint32(v10+v12)]))
								v11 = t150
								if v11 < i32(0) {
									v12 = v14 + v12
									v16 = v15 + i32(2)
									goto l225
								}
								v9 = v15 + i32(2)
								m.memory[uint32(v17+v12)] = byte(v11)
								v15 = v14 + v12
								v18 = v15 + i32(1)
								if uint32(v11) > uint32(i32(59)) {
									goto l221
								}
								if v8+v12 != 0 {
									if uint32(v15+i32(3)) < uint32(v5) {
										v12 = v12 + i32(1)
										goto l224
									}
									store32(m.memory[int64(uint32(v0))+8:], uint32(v18))
									store32(m.memory[uint32(v0):], uint32(v9))
									m.memory[int64(uint32(v0))+4] = byte(i32(1))
									goto l35
								}
								store32(m.memory[int64(uint32(v0))+8:], uint32(v18))
								store32(m.memory[uint32(v0):], uint32(v3))
								m.memory[int64(uint32(v0))+4] = byte(i32(0))
								goto l35
							}
						}
						store32(m.memory[int64(uint32(v0))+8:], uint32(v14))
						store32(m.memory[uint32(v0):], uint32(v18))
						m.memory[int64(uint32(v0))+4] = byte(i32(1))
						goto l35
					}
					store32(m.memory[int64(uint32(v0))+8:], uint32(v14))
					store32(m.memory[uint32(v0):], uint32(v18))
					m.memory[int64(uint32(v0))+4] = byte(i32(0))
					goto l35
				l176:
					store32(m.memory[int64(uint32(v0))+8:], uint32(v12))
					store32(m.memory[uint32(v0):], uint32(v16))
					goto l35
				}
				v13 = i32(1)
				goto l170
			l170:
				store32(m.memory[int64(uint32(v0))+8:], uint32(v12))
				store32(m.memory[uint32(v0):], uint32(v15))
				store32(m.memory[int64(uint32(v0))+4:], uint32(v13|(v7+i32(340))&i32(-256)))
				goto l35
			}
		case 6:
			store32(m.memory[int64(uint32(v7))+348:], uint32(i32(0)))
			store32(m.memory[int64(uint32(v7))+344:], uint32(v5))
			store32(m.memory[int64(uint32(v7))+340:], uint32(v4))
			v10 = i32(0)
			v15 = i32(0)
			{
				t151 := int32(m.memory[int64(uint32(v1))+1])
				if t151 == 0 {
					goto l292
				}
				m.memory[int64(uint32(v1))+1] = byte(i32(0))
				if v3 == 0 {
					if v6 != 0 {
						m.memory[int64(uint32(v0))+6] = byte(i32(0))
						store16(m.memory[int64(uint32(v0))+4:], uint16(i32(258)))
						store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
						store32(m.memory[uint32(v0):], uint32(i32(0)))
						goto l35
					}
					store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
					store32(m.memory[uint32(v0):], uint32(i32(0)))
					m.memory[int64(uint32(v0))+4] = byte(i32(0))
					goto l35
				}
				if uint32(v5) > uint32(i32(2)) {
					t152 := int32(int8(m.memory[uint32(v2)]))
					v12 = t152
					{
						t153 := int32(m.memory[int64(uint32(v1))+2])
						v11 = t153
						if v11 != i32(1) {
							goto l229
						}
						if uint32((v12+i32(97))&i32(255)) < uint32(i32(83)) {
							m.memory[uint32(v4)] = byte(i32(227))
							t169 := v4
							v12 = v12 + i32(-94)
							m.memory[int64(uint32(t169))+1] = byte(int32(uint32(v12&i32(192))>>6) | i32(128))
							goto l240
						}
					}
				l229:
					v15 = v12 + i32(-64)
					if uint32(v15&i32(255)) <= uint32(i32(62)) {
						goto l231
					}
					if v12 > i32(-4) {
						m.memory[int64(uint32(v0))+4] = byte(i32(2))
						if v12 > i32(-1) {
							store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
							store32(m.memory[uint32(v0):], uint32(i32(0)))
							store16(m.memory[int64(uint32(v0))+5:], uint16(i32(1)))
							goto l35
						}
						store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
						store32(m.memory[uint32(v0):], uint32(i32(1)))
						store16(m.memory[int64(uint32(v0))+5:], uint16(i32(2)))
						goto l35
					}
					v15 = v12 + i32(-65)
				l231:
					{
						if v11 != i32(2) {
							goto l233
						}
						v16 = v15 & i32(255)
						if uint32(v16) < uint32(i32(86)) {
							m.memory[uint32(v4)] = byte(i32(227))
							t155 := v4
							v12 = v16 + i32(12449)
							m.memory[int64(uint32(t155))+1] = byte(int32(uint32(v12)>>6) & i32(135))
							goto l240
						}
					l233:
						v11 = v11*i32(188) + v15&i32(255)
						v15 = v11 + i32(-1410)
						if uint32(v15) < uint32(i32(2965)) {
							t156 := int32(load16(m.memory[int64(uint32(v15<<1))+1244278:]))
							t157 := v4
							v12 = t156
							m.memory[uint32(t157)] = byte(int32(uint32(v12)>>12) | i32(224))
							m.memory[int64(uint32(v4))+1] = byte(int32(uint32(v12)>>6)&i32(63) | i32(128))
							goto l240
						}
						v15 = v11 + i32(-4418)
						if uint32(v15) < uint32(i32(3390)) {
							t158 := int32(load16(m.memory[int64(uint32(v15<<1))+1272642:]))
							t159 := v4
							v12 = t158
							m.memory[uint32(t159)] = byte(int32(uint32(v12)>>12) | i32(224))
							m.memory[int64(uint32(v4))+1] = byte(int32(uint32(v12)>>6)&i32(63) | i32(128))
							goto l240
						}
						v15 = v11 + i32(-10744)
						if uint32(v15) < uint32(i32(360)) {
							t160 := int32(load16(m.memory[int64(uint32(v15<<1))+1279422:]))
							t161 := v4
							v12 = t160
							m.memory[uint32(t161)] = byte(int32(uint32(v12)>>12) | i32(224))
							m.memory[int64(uint32(v4))+1] = byte(int32(uint32(v12)>>6)&i32(63) | i32(128))
							goto l240
						}
						v15 = v11 + i32(-8272)
						if uint32(v15) < uint32(i32(360)) {
							t162 := int32(load16(m.memory[int64(uint32(v15<<1))+1279422:]))
							t163 := v4
							v12 = t162
							m.memory[uint32(t163)] = byte(int32(uint32(v12)>>12) | i32(224))
							m.memory[int64(uint32(v4))+1] = byte(int32(uint32(v12)>>6)&i32(63) | i32(128))
							goto l240
						}
						if uint32(v11+i32(-8836)) >= uint32(i32(1880)) {
							m.fn1712(v7+i32(232), v11)
							{
								t164 := int32(load16(m.memory[int64(uint32(v7))+232:]))
								if t164&i32(1) != 0 {
									t167 := int32(load16(m.memory[int64(uint32(v7))+234:]))
									v12 = t167
									if uint32(v12&i32(0xffff)) < uint32(i32(2048)) {
										m.memory[uint32(v4)] = byte(int32(uint32(v12)>>6) | i32(192))
										v10 = i32(2)
										v11 = i32(1)
										goto l246
									}
									m.memory[int64(uint32(v4))+1] = byte(int32(uint32(v12)>>6)&i32(63) | i32(128))
									m.memory[uint32(v4)] = byte(int32(uint32(v12&i32(61440))>>12) | i32(224))
									goto l240
								}
								v16 = i32(0)
								v15 = i32(1250730)
							l244:
								{
									if uint32(v16) > uint32(i32(53)) {
										m.memory[int64(uint32(v0))+4] = byte(i32(2))
										if v12 > i32(-1) {
											store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
											store32(m.memory[uint32(v0):], uint32(i32(0)))
											store16(m.memory[int64(uint32(v0))+5:], uint16(i32(1)))
											goto l35
										}
										store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
										store32(m.memory[uint32(v0):], uint32(i32(1)))
										store16(m.memory[int64(uint32(v0))+5:], uint16(i32(2)))
										goto l35
									}
									t165 := int32(load16(m.memory[uint32(v15):]))
									v14 = v11 - t165
									t166 := int32(load16(m.memory[uint32(v15+i32(2)):]))
									if uint32(v14) < uint32(t166) {
										v12 = v16 + i32(2)
										if uint32(v16) >= uint32(i32(52)) {
											m.fn158(v12, i32(54), i32(1250208))
											panic("unreachable")
										}
										v11 = i32(1)
										{
											t168 := int32(load16(m.memory[int64(uint32(v12<<1))+1250730:]))
											v12 = t168 + v14
											if uint32(v12&i32(0xffff)) < uint32(i32(2048)) {
												m.memory[uint32(v4)] = byte(int32(uint32(v12)>>6) | i32(192))
												v10 = i32(2)
												goto l246
											}
											m.memory[int64(uint32(v4))+1] = byte(int32(uint32(v12)>>6)&i32(63) | i32(128))
											m.memory[uint32(v4)] = byte(int32(uint32(v12&i32(61440))>>12) | i32(224))
											goto l240
										}
									}
									v15 = v15 + i32(6)
									v16 = v16 + i32(3)
									goto l244
								}
							}
						}
						t154 := v4
						v12 = v11 + i32(-17028)
						m.memory[int64(uint32(t154))+1] = byte(int32(uint32(v12)>>6)&i32(63) | i32(128))
						m.memory[uint32(v4)] = byte(int32(uint32(v12&i32(61440))>>12) | i32(224))
						goto l240
					}
				}
				store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
				store32(m.memory[uint32(v0):], uint32(i32(0)))
				m.memory[int64(uint32(v0))+4] = byte(i32(1))
				goto l35
			l240:
				v10 = i32(3)
				v11 = i32(2)
			l246:
				m.memory[uint32(v4+v11)] = byte(v12&i32(63) | i32(128))
				v15 = i32(1)
			}
		l292:
			{
				m.fn148(v7+i32(224), v15, v2, v3, i32(1155464))
				t170 := int32(load32(m.memory[int64(uint32(v7))+224:]))
				v16 = t170
				t171 := int32(load32(m.memory[int64(uint32(v7))+228:]))
				v12 = t171
				m.fn212(v7+i32(216), v10, v4, v5, i32(1155480))
				t172 := int32(load32(m.memory[int64(uint32(v7))+220:]))
				v11 = t172
				t173 := v11
				t174 := v12
				var p175 int32
				if uint32(v11) < uint32(v12) {
					p175 = 1
				}
				v13 = p175
				p176 := t174
				if v13 != 0 {
					p176 = t173
				}
				v9 = p176
				v12 = i32(0)
				{
					{
						t177 := int32(load32(m.memory[int64(uint32(v7))+216:]))
						t178 := v16
						v14 = t177
						if (t178^v14)&i32(3) != 0 {
							goto l252
						}
						v12 = i32(0)
						v18 = (i32(0) - v16) & i32(3)
						if uint32(v18|i32(8)) > uint32(v9) {
							goto l252
						}
						v12 = i32(0)
					l259:
						if v18 != v12 {
							t185 := int32(int8(m.memory[uint32(v16+v12)]))
							v11 = t185
							if v11 < i32(0) {
								goto l257
							}
							m.memory[uint32(v14+v12)] = byte(v11)
							v12 = v12 + i32(1)
							goto l259
						}
						v8 = v9 + i32(-8)
					l258:
						{
							v18 = v16 + v12
							t179 := int32(load32(m.memory[uint32(v18):]))
							v11 = t179
							v17 = v14 + v12
							t180 := int32(load32(m.memory[uint32(v18+i32(4)):]))
							t181 := v17 + i32(4)
							v18 = t180
							store32(m.memory[uint32(t181):], uint32(v18))
							store32(m.memory[uint32(v17):], uint32(v11))
							{
								v18 = v18 & i32(-2139062144)
								t182 := v18
								v11 = v11 & i32(-2139062144)
								if t182|v11 == 0 {
									v12 = v12 + i32(8)
									if uint32(v12) <= uint32(v8) {
										goto l258
									}
									goto l252
								}
								if v11 != 0 {
									goto l255
								}
								v11 = int32(uint32(int32(bits.TrailingZeros32(uint32(v18))))>>3) + i32(4)
								goto l256
							l255:
								v11 = int32(uint32(int32(bits.TrailingZeros32(uint32(v11)))) >> 3)
							l256:
								t183 := v16
								v12 = v11 + v12
								t184 := int32(m.memory[uint32(t183+v12)])
								v11 = t184
								goto l257
							}
						}
					}
				l252:
					p186 := v9
					if uint32(v12) > uint32(v9) {
						p186 = v12
					}
					v18 = p186
				l261:
					{
						if v18 == v12 {
							v12 = v9 + v10
							v15 = v9 + v15
							goto l262
						}
						t187 := int32(int8(m.memory[uint32(v16+v12)]))
						v11 = t187
						if v11 < i32(0) {
							goto l257
						}
						m.memory[uint32(v14+v12)] = byte(v11)
						v12 = v12 + i32(1)
						goto l261
					}
				}
			l257:
				v15 = v12 + v15
				v12 = v12 + v10
				if uint32(v12+i32(2)) < uint32(v5) {
					v15 = v15 + i32(1)
				l296:
					{
						v16 = v11 + i32(127)
						if uint32(v16&i32(255)) < uint32(i32(31)) {
							goto l264
						}
						if uint32((v11+i32(3))&i32(255)) < uint32(i32(227)) {
							v16 = (v11 + i32(95)) & i32(255)
							if uint32(v16) > uint32(i32(62)) {
								if v11&i32(255) != i32(128) {
									m.memory[int64(uint32(v0))+6] = byte(i32(0))
									store16(m.memory[int64(uint32(v0))+4:], uint16(i32(258)))
									goto l268
								}
								t202 := int32(load32(m.memory[int64(uint32(v7))+340:]))
								v4 = t202
								store16(m.memory[uint32(v4+v12):], uint16(i32(32962)))
								v10 = v12 + i32(2)
								goto l292
							}
							t200 := int32(load32(m.memory[int64(uint32(v7))+340:]))
							v4 = t200
							v11 = v4 + v12
							m.memory[uint32(v11)] = byte(i32(239))
							t201 := v11 + i32(2)
							v16 = v16 + i32(-159)
							m.memory[uint32(t201)] = byte(v16&i32(63) | i32(128))
							m.memory[uint32(v11+i32(1))] = byte(int32(uint32(v16)>>6) & i32(191))
							v10 = v12 + i32(3)
							goto l292
						}
						v16 = v11 + i32(63)
					l264:
						if uint32(v15) < uint32(v3) {
							v11 = v15 + i32(1)
							t188 := int32(int8(m.memory[uint32(v2+v15)]))
							v5 = t188
							{
								v16 = v16 & i32(255)
								if v16 != i32(1) {
									goto l269
								}
								if uint32((v5+i32(97))&i32(255)) < uint32(i32(83)) {
									t192 := int32(load32(m.memory[int64(uint32(v7))+340:]))
									v4 = t192
									v16 = v4 + v12
									m.memory[uint32(v16)] = byte(i32(227))
									t193 := v16 + i32(2)
									v5 = v5 + i32(-94)
									m.memory[uint32(t193)] = byte(v5&i32(63) | i32(128))
									m.memory[uint32(v16+i32(1))] = byte(int32(uint32(v5&i32(192))>>6) | i32(128))
									goto l284
								}
							l269:
								v14 = v5 + i32(-64)
								if uint32(v14&i32(255)) <= uint32(i32(62)) {
									goto l271
								}
								if v5 > i32(-4) {
									m.memory[int64(uint32(v0))+4] = byte(i32(2))
									if v5 > i32(-1) {
										store32(m.memory[int64(uint32(v0))+8:], uint32(v12))
										store32(m.memory[uint32(v0):], uint32(v15))
										store16(m.memory[int64(uint32(v0))+5:], uint16(i32(1)))
										goto l35
									}
									store32(m.memory[int64(uint32(v0))+8:], uint32(v12))
									store32(m.memory[uint32(v0):], uint32(v11))
									store16(m.memory[int64(uint32(v0))+5:], uint16(i32(2)))
									goto l35
								}
								v14 = v5 + i32(-65)
							l271:
								if v16 != i32(2) {
									goto l273
								}
								v18 = v14 & i32(255)
								if uint32(v18) < uint32(i32(86)) {
									t203 := int32(load32(m.memory[int64(uint32(v7))+340:]))
									v4 = t203
									v5 = v4 + v12
									m.memory[uint32(v5)] = byte(i32(227))
									t204 := v5 + i32(2)
									v16 = v18 + i32(12449)
									m.memory[uint32(t204)] = byte(v16&i32(63) | i32(128))
									m.memory[uint32(v5+i32(1))] = byte(int32(uint32(v16)>>6) & i32(135))
									goto l284
								}
							l273:
								v16 = v16*i32(188) + v14&i32(255)
								v14 = v16 + i32(-1410)
								if uint32(v14) < uint32(i32(2965)) {
									t205 := int32(load32(m.memory[int64(uint32(v7))+340:]))
									v4 = t205
									v5 = v4 + v12
									t206 := int32(load16(m.memory[int64(uint32(v14<<1))+1244278:]))
									t207 := v5 + i32(2)
									v16 = t206
									m.memory[uint32(t207)] = byte(v16&i32(63) | i32(128))
									m.memory[uint32(v5)] = byte(int32(uint32(v16)>>12) | i32(224))
									m.memory[uint32(v5+i32(1))] = byte(int32(uint32(v16)>>6)&i32(63) | i32(128))
									goto l284
								}
								v14 = v16 + i32(-4418)
								if uint32(v14) < uint32(i32(3390)) {
									t208 := int32(load32(m.memory[int64(uint32(v7))+340:]))
									v4 = t208
									v5 = v4 + v12
									t209 := int32(load16(m.memory[int64(uint32(v14<<1))+1272642:]))
									t210 := v5 + i32(2)
									v16 = t209
									m.memory[uint32(t210)] = byte(v16&i32(63) | i32(128))
									m.memory[uint32(v5)] = byte(int32(uint32(v16)>>12) | i32(224))
									m.memory[uint32(v5+i32(1))] = byte(int32(uint32(v16)>>6)&i32(63) | i32(128))
									goto l284
								}
								v14 = v16 + i32(-10744)
								if uint32(v14) < uint32(i32(360)) {
									t211 := int32(load32(m.memory[int64(uint32(v7))+340:]))
									v4 = t211
									v5 = v4 + v12
									t212 := int32(load16(m.memory[int64(uint32(v14<<1))+1279422:]))
									t213 := v5 + i32(2)
									v16 = t212
									m.memory[uint32(t213)] = byte(v16&i32(63) | i32(128))
									m.memory[uint32(v5)] = byte(int32(uint32(v16)>>12) | i32(224))
									m.memory[uint32(v5+i32(1))] = byte(int32(uint32(v16)>>6)&i32(63) | i32(128))
									goto l284
								}
								v14 = v16 + i32(-8272)
								if uint32(v14) < uint32(i32(360)) {
									t214 := int32(load32(m.memory[int64(uint32(v7))+340:]))
									v4 = t214
									v5 = v4 + v12
									t215 := int32(load16(m.memory[int64(uint32(v14<<1))+1279422:]))
									t216 := v5 + i32(2)
									v16 = t215
									m.memory[uint32(t216)] = byte(v16&i32(63) | i32(128))
									m.memory[uint32(v5)] = byte(int32(uint32(v16)>>12) | i32(224))
									m.memory[uint32(v5+i32(1))] = byte(int32(uint32(v16)>>6)&i32(63) | i32(128))
									goto l284
								}
								if uint32(v16+i32(-8836)) < uint32(i32(1880)) {
									t194 := int32(load32(m.memory[int64(uint32(v7))+340:]))
									v4 = t194
									v5 = v4 + v12
									t195 := v5 + i32(2)
									v16 = v16 + i32(-17028)
									m.memory[uint32(t195)] = byte(v16&i32(63) | i32(128))
									m.memory[uint32(v5+i32(1))] = byte(int32(uint32(v16)>>6)&i32(63) | i32(128))
									m.memory[uint32(v5)] = byte(int32(uint32(v16&i32(61440))>>12) | i32(224))
									goto l284
								}
								m.fn1712(v7+i32(208), v16)
								t189 := int32(load16(m.memory[int64(uint32(v7))+208:]))
								if t189&i32(1) != 0 {
									t196 := int32(load16(m.memory[int64(uint32(v7))+210:]))
									v5 = t196
									if uint32(v5&i32(0xffff)) < uint32(i32(2048)) {
										t198 := int32(load32(m.memory[int64(uint32(v7))+340:]))
										v4 = t198
										v16 = v4 + v12
										m.memory[uint32(v16+i32(1))] = byte(v5&i32(63) | i32(128))
										m.memory[uint32(v16)] = byte(int32(uint32(v5)>>6) | i32(192))
										v5 = i32(2)
										goto l286
									}
									t197 := int32(load32(m.memory[int64(uint32(v7))+340:]))
									v4 = t197
									v16 = v4 + v12
									m.memory[uint32(v16+i32(2))] = byte(v5&i32(63) | i32(128))
									m.memory[uint32(v16+i32(1))] = byte(int32(uint32(v5)>>6)&i32(63) | i32(128))
									m.memory[uint32(v16)] = byte(int32(uint32(v5&i32(61440))>>12) | i32(224))
									goto l284
								}
								v18 = i32(0)
								v14 = i32(1250730)
							l283:
								{
									if uint32(v18) > uint32(i32(53)) {
										m.memory[int64(uint32(v0))+4] = byte(i32(2))
										if v5 > i32(-1) {
											store32(m.memory[int64(uint32(v0))+8:], uint32(v12))
											store32(m.memory[uint32(v0):], uint32(v15))
											store16(m.memory[int64(uint32(v0))+5:], uint16(i32(1)))
											goto l35
										}
										store32(m.memory[int64(uint32(v0))+8:], uint32(v12))
										store32(m.memory[uint32(v0):], uint32(v11))
										store16(m.memory[int64(uint32(v0))+5:], uint16(i32(2)))
										goto l35
									}
									t190 := int32(load16(m.memory[uint32(v14):]))
									v9 = v16 - t190
									t191 := int32(load16(m.memory[uint32(v14+i32(2)):]))
									if uint32(v9) < uint32(t191) {
										v5 = v18 + i32(2)
										if uint32(v18) >= uint32(i32(52)) {
											m.fn158(v5, i32(54), i32(1250208))
											panic("unreachable")
										}
										{
											t199 := int32(load16(m.memory[int64(uint32(v5<<1))+1250730:]))
											v5 = t199 + v9
											if uint32(v5&i32(0xffff)) < uint32(i32(2048)) {
												v16 = v4 + v12
												m.memory[uint32(v16+i32(1))] = byte(v5&i32(63) | i32(128))
												m.memory[uint32(v16)] = byte(int32(uint32(v5)>>6) | i32(192))
												v5 = i32(2)
												goto l286
											}
											v16 = v4 + v12
											m.memory[uint32(v16+i32(2))] = byte(v5&i32(63) | i32(128))
											m.memory[uint32(v16+i32(1))] = byte(int32(uint32(v5)>>6)&i32(63) | i32(128))
											m.memory[uint32(v16)] = byte(int32(uint32(v5&i32(61440))>>12) | i32(224))
											goto l284
										}
									}
									v14 = v14 + i32(6)
									v18 = v18 + i32(3)
									goto l283
								}
							}
						}
						if v6 != 0 {
							m.memory[int64(uint32(v0))+6] = byte(i32(0))
							store16(m.memory[int64(uint32(v0))+4:], uint16(i32(258)))
							goto l268
						}
						m.memory[int64(uint32(v1))+2] = byte(v16)
						m.memory[int64(uint32(v1))+1] = byte(i32(1))
						m.memory[int64(uint32(v0))+4] = byte(i32(0))
						goto l268
					l268:
						store32(m.memory[int64(uint32(v0))+8:], uint32(v12))
						store32(m.memory[uint32(v0):], uint32(v15))
						goto l35
					l284:
						v5 = i32(3)
					l286:
						v12 = v5 + v12
						if uint32(v11) < uint32(v3) {
							goto l294
						}
						store32(m.memory[int64(uint32(v0))+8:], uint32(v12))
						store32(m.memory[uint32(v0):], uint32(v11))
						m.memory[int64(uint32(v0))+4] = byte(i32(0))
						goto l35
					l294:
						{
							t217 := int32(load32(m.memory[int64(uint32(v7))+344:]))
							t218 := v12 + i32(2)
							v5 = t217
							if uint32(t218) < uint32(v5) {
								goto l295
							}
							store32(m.memory[int64(uint32(v0))+8:], uint32(v12))
							store32(m.memory[uint32(v0):], uint32(v11))
							m.memory[int64(uint32(v0))+4] = byte(i32(1))
							goto l35
						}
					l295:
						v15 = v15 + i32(2)
						t219 := int32(int8(m.memory[uint32(v2+v11)]))
						v11 = t219
						if v11 < i32(0) {
							goto l296
						}
					}
					t220 := int32(load32(m.memory[int64(uint32(v7))+340:]))
					v4 = t220
					m.memory[uint32(v4+v12)] = byte(v11)
					v10 = v12 + i32(1)
					goto l292
				}
				v13 = i32(1)
			l262:
				store32(m.memory[int64(uint32(v0))+8:], uint32(v12))
				store32(m.memory[uint32(v0):], uint32(v15))
				store32(m.memory[int64(uint32(v0))+4:], uint32(v13|(v7+i32(340))&i32(-256)))
				goto l35
			}
		case 5:
			{
				{
					t221 := int32(m.memory[int64(uint32(v1))+2])
					if t221 != 0 {
						goto l297
					}
					t222 := int32(m.memory[int64(uint32(v1))+1])
					v17 = t222
					t223 := int32(m.memory[int64(uint32(v1))+3])
					v15 = t223
					v16 = i32(0)
					goto l298
				}
			l297:
				if uint32(v5) > uint32(i32(2)) {
					goto l299
				}
				m.memory[int64(uint32(v0))+4] = byte(i32(1))
				v14 = i32(0)
				v16 = i32(0)
				goto l300
			l299:
				v16 = i32(0)
				store16(m.memory[int64(uint32(v1))+1:], uint16(i32(0)))
				{
					t224 := int32(m.memory[int64(uint32(v1))+3])
					v15 = t224
					switch v15 {
					default:
						m.fn256(i32(1286542), i32(40), i32(1155024))
						panic("unreachable")
					case 0, 1:
						t225 := int32(m.memory[int64(uint32(v1))+5])
						m.memory[uint32(v4)] = byte(t225)
						v17 = i32(0)
						m.memory[int64(uint32(v1))+5] = byte(i32(0))
						v16 = i32(1)
						goto l298
					case 2:
						t226 := int32(m.memory[int64(uint32(v1))+5])
						v12 = t226
						v17 = i32(0)
						m.memory[int64(uint32(v1))+5] = byte(i32(0))
						m.memory[int64(uint32(v4))+2] = byte(v12&i32(63) | i32(128))
						t227 := v4
						v12 = v12 + i32(-192)
						m.memory[int64(uint32(t227))+1] = byte(int32(uint32(v12)>>6)&i32(63) | i32(128))
						m.memory[uint32(v4)] = byte(int32(uint32(v12&i32(61440))>>12) | i32(224))
						v16 = i32(3)
						v15 = i32(2)
						goto l298
					case 3:
						v15 = i32(4)
						m.memory[int64(uint32(v1))+3] = byte(i32(4))
						v17 = i32(0)
					}
				}
			l298:
				t228 := int32(m.memory[int64(uint32(v1))+5])
				v10 = t228
				v12 = i32(0)
			l322:
				if v12 != v3 {
					v18 = v16 + i32(2)
					if uint32(v18) < uint32(v5) {
						v14 = v12 + i32(1)
						t229 := int32(m.memory[uint32(v2+v12)])
						v9 = t229
						v11 = int32(int8(v9))
						switch v15 & i32(255) {
						case 5:
							switch v9 + i32(-36) {
							case 0, 4:
								goto l316
							default:
								m.memory[int64(uint32(v1))+1] = byte(i32(0))
								m.memory[int64(uint32(v0))+6] = byte(i32(0))
								store16(m.memory[int64(uint32(v0))+4:], uint16(i32(258)))
								t238 := int32(m.memory[int64(uint32(v1))+4])
								m.memory[int64(uint32(v1))+3] = byte(t238)
								v14 = v12
								goto l300
							}
						case 6:
							v18 = v10 & i32(255)
							var p230 int32
							if v18 != i32(40) {
								p230 = 1
							}
							v15 = p230
							if v15 != 0 {
								goto l318
							}
							if v11 != i32(66) {
								goto l318
							}
							v15 = i32(0)
							goto l319
						default:
							if v11 == i32(27) {
								goto l320
							}
							m.memory[int64(uint32(v1))+1] = byte(i32(0))
							if v11 < i32(0) {
								goto l321
							}
							if v11&i32(254) == i32(14) {
								goto l321
							}
							m.memory[uint32(v4+v16)] = byte(v11)
							v16 = v16 + i32(1)
							v15 = i32(0)
							v12 = v14
							v17 = i32(0)
							goto l322
						l321:
							m.memory[int64(uint32(v0))+6] = byte(i32(0))
							store16(m.memory[int64(uint32(v0))+4:], uint16(i32(258)))
							goto l300
						case 1:
							if v11 == i32(27) {
								goto l320
							}
							m.memory[int64(uint32(v1))+1] = byte(i32(0))
							if v9 == i32(126) {
								m.memory[uint32(v4+v18)] = byte(i32(190))
								store16(m.memory[uint32(v4+v16):], uint16(i32(32994)))
								v16 = v16 + i32(3)
								v17 = i32(0)
								v15 = i32(1)
								v12 = v14
								goto l322
							}
							if v9 != i32(92) {
								if v11 < i32(0) {
									goto l326
								}
								if v11&i32(254) != i32(14) {
									m.memory[uint32(v4+v16)] = byte(v11)
									v15 = i32(1)
									v16 = v16 + i32(1)
									v17 = i32(0)
									v12 = v14
									goto l322
								}
							l326:
								m.memory[int64(uint32(v0))+6] = byte(i32(0))
								store16(m.memory[int64(uint32(v0))+4:], uint16(i32(258)))
								goto l300
							}
							store16(m.memory[uint32(v4+v16):], uint16(i32(42434)))
							v17 = i32(0)
							v15 = i32(1)
							goto l325
						case 2:
							if v11 == i32(27) {
								goto l320
							}
							m.memory[int64(uint32(v1))+1] = byte(i32(0))
							if uint32((v11+i32(-33))&i32(255)) < uint32(i32(63)) {
								v12 = v4 + v16
								m.memory[uint32(v12)] = byte(i32(239))
								m.memory[uint32(v4+v18)] = byte(v11&i32(63) | i32(128))
								m.memory[uint32(v12+i32(1))] = byte(int32(uint32(v11+i32(16192))>>6) & i32(191))
								v16 = v16 + i32(3)
								v17 = i32(0)
								v15 = i32(2)
								v12 = v14
								goto l322
							}
							m.memory[int64(uint32(v0))+6] = byte(i32(0))
							store16(m.memory[int64(uint32(v0))+4:], uint16(i32(258)))
							goto l300
						case 3:
							if v11 == i32(27) {
								goto l320
							}
							m.memory[int64(uint32(v1))+1] = byte(i32(0))
							if uint32((v11+i32(-33))&i32(255)) < uint32(i32(94)) {
								v15 = i32(4)
								m.memory[int64(uint32(v1))+3] = byte(i32(4))
								m.memory[int64(uint32(v1))+5] = byte(v11)
								v17 = i32(0)
								goto l330
							}
							m.memory[int64(uint32(v0))+6] = byte(i32(0))
							store16(m.memory[int64(uint32(v0))+4:], uint16(i32(258)))
							goto l300
						case 4:
							if v11 == i32(27) {
								m.memory[int64(uint32(v0))+6] = byte(i32(1))
								store16(m.memory[int64(uint32(v0))+4:], uint16(i32(258)))
								m.memory[int64(uint32(v1))+3] = byte(i32(5))
								goto l300
							}
							m.memory[int64(uint32(v1))+3] = byte(i32(3))
							v12 = v11 + i32(-33)
							v15 = (v10 + i32(-33)) & i32(255)
							if v15 != i32(3) {
								goto l332
							}
							if uint32(v12&i32(255)) < uint32(i32(83)) {
								v12 = v4 + v16
								m.memory[uint32(v12)] = byte(i32(227))
								t242 := v4 + v18
								v11 = v11 + i32(32)
								m.memory[uint32(t242)] = byte(v11&i32(63) | i32(128))
								m.memory[uint32(v12+i32(1))] = byte(int32(uint32(v11&i32(192))>>6) | i32(128))
								v15 = i32(3)
								v16 = v16 + i32(3)
								v10 = i32(36)
								v12 = v14
								goto l322
							}
						l332:
							if v15 != i32(4) {
								goto l334
							}
							v11 = v12 & i32(255)
							if uint32(v11) < uint32(i32(86)) {
								v12 = v4 + v16
								m.memory[uint32(v12)] = byte(i32(227))
								t243 := v12 + i32(1)
								v12 = v11 + i32(12449)
								m.memory[uint32(t243)] = byte(int32(uint32(v12)>>6) & i32(135))
								m.memory[uint32(v4+v18)] = byte(v12&i32(63) | i32(128))
								v15 = i32(3)
								v16 = v16 + i32(3)
								v10 = i32(37)
								v12 = v14
								goto l322
							}
						l334:
							v12 = v12 & i32(255)
							if uint32(v12) > uint32(i32(93)) {
								m.memory[int64(uint32(v0))+6] = byte(i32(0))
								store16(m.memory[int64(uint32(v0))+4:], uint16(i32(514)))
								goto l300
							}
							v15 = v15*i32(94) + v12
							v12 = v15 + i32(-1410)
							if uint32(v12) < uint32(i32(2965)) {
								v11 = v4 + v16
								t244 := int32(load16(m.memory[int64(uint32(v12<<1))+1244278:]))
								t245 := v11
								v12 = t244
								m.memory[uint32(t245)] = byte(int32(uint32(v12)>>12) | i32(224))
								m.memory[uint32(v4+v18)] = byte(v12&i32(63) | i32(128))
								m.memory[uint32(v11+i32(1))] = byte(int32(uint32(v12)>>6)&i32(63) | i32(128))
								goto l344
							}
							v12 = v15 + i32(-4418)
							if uint32(v12) < uint32(i32(3390)) {
								v11 = v4 + v16
								t246 := int32(load16(m.memory[int64(uint32(v12<<1))+1272642:]))
								t247 := v11
								v12 = t246
								m.memory[uint32(t247)] = byte(int32(uint32(v12)>>12) | i32(224))
								m.memory[uint32(v4+v18)] = byte(v12&i32(63) | i32(128))
								m.memory[uint32(v11+i32(1))] = byte(int32(uint32(v12)>>6)&i32(63) | i32(128))
								goto l344
							}
							{
								v12 = v15 + i32(-8272)
								if uint32(v12) < uint32(i32(360)) {
									v11 = v4 + v16
									t234 := int32(load16(m.memory[int64(uint32(v12<<1))+1279422:]))
									t235 := v11
									v12 = t234
									m.memory[uint32(t235)] = byte(int32(uint32(v12)>>12) | i32(224))
									m.memory[uint32(v4+v18)] = byte(v12&i32(63) | i32(128))
									m.memory[uint32(v11+i32(1))] = byte(int32(uint32(v12)>>6)&i32(63) | i32(128))
									goto l344
								}
								m.fn1712(v7+i32(200), v15)
								t231 := int32(load16(m.memory[int64(uint32(v7))+200:]))
								if t231&i32(1) != 0 {
									v11 = v4 + v16
									{
										t236 := int32(load16(m.memory[int64(uint32(v7))+202:]))
										v12 = t236
										if uint32(v12&i32(0xffff)) < uint32(i32(2048)) {
											m.memory[uint32(v11+i32(1))] = byte(v12&i32(63) | i32(128))
											m.memory[uint32(v11)] = byte(int32(uint32(v12)>>6) | i32(192))
											goto l346
										}
										m.memory[uint32(v4+v18)] = byte(v12&i32(63) | i32(128))
										m.memory[uint32(v11+i32(1))] = byte(int32(uint32(v12)>>6)&i32(63) | i32(128))
										m.memory[uint32(v11)] = byte(int32(uint32(v12&i32(61440))>>12) | i32(224))
										goto l344
									}
								}
								v11 = i32(0)
								v12 = i32(1250730)
							l343:
								{
									if uint32(v11) > uint32(i32(53)) {
										m.memory[int64(uint32(v0))+6] = byte(i32(0))
										store16(m.memory[int64(uint32(v0))+4:], uint16(i32(514)))
										goto l300
									}
									t232 := int32(load16(m.memory[uint32(v12):]))
									v9 = v15 - t232
									t233 := int32(load16(m.memory[uint32(v12+i32(2)):]))
									if uint32(v9) < uint32(t233) {
										v12 = v11 + i32(2)
										if uint32(v11) >= uint32(i32(52)) {
											m.fn158(v12, i32(54), i32(1250208))
											panic("unreachable")
										}
										v11 = v4 + v16
										{
											t237 := int32(load16(m.memory[int64(uint32(v12<<1))+1250730:]))
											v12 = t237 + v9
											if uint32(v12&i32(0xffff)) < uint32(i32(2048)) {
												m.memory[uint32(v11+i32(1))] = byte(v12&i32(63) | i32(128))
												m.memory[uint32(v11)] = byte(int32(uint32(v12)>>6) | i32(192))
												goto l346
											}
											m.memory[uint32(v4+v18)] = byte(v12&i32(63) | i32(128))
											m.memory[uint32(v11+i32(1))] = byte(int32(uint32(v12)>>6)&i32(63) | i32(128))
											m.memory[uint32(v11)] = byte(int32(uint32(v12&i32(61440))>>12) | i32(224))
											goto l344
										}
									}
									v12 = v12 + i32(6)
									v11 = v11 + i32(3)
									goto l343
								}
							}
						}
					l316:
						v15 = i32(6)
						m.memory[int64(uint32(v1))+3] = byte(i32(6))
						m.memory[int64(uint32(v1))+5] = byte(v11)
					l330:
						v10 = v11
						v12 = v14
						goto l322
					l318:
						if v15 != 0 {
							goto l349
						}
						if v11 != i32(74) {
							goto l349
						}
						v15 = i32(1)
						goto l319
					l349:
						if v15 != 0 {
							goto l350
						}
						if v11 != i32(73) {
							goto l350
						}
						v15 = i32(2)
						goto l319
					l350:
						if v18 != i32(36) {
							goto l351
						}
						v15 = i32(3)
						switch v9 + i32(-64) {
						case 0, 2:
							goto l319
						default:
							goto l351
						}
					l351:
						store16(m.memory[int64(uint32(v1))+1:], uint16(i32(256)))
						m.memory[int64(uint32(v0))+6] = byte(i32(1))
						store16(m.memory[int64(uint32(v0))+4:], uint16(i32(258)))
						t239 := int32(m.memory[int64(uint32(v1))+4])
						m.memory[int64(uint32(v1))+3] = byte(t239)
						v14 = v12
						goto l300
					}
					m.memory[int64(uint32(v0))+4] = byte(i32(1))
					v14 = v12
					goto l300
				}
				if v6 != 0 {
					switch v15&i32(255) + i32(-4) {
					default:
						goto l307
					case 0, 1:
						m.memory[int64(uint32(v0))+6] = byte(i32(0))
						store16(m.memory[int64(uint32(v0))+4:], uint16(i32(258)))
						t240 := int32(m.memory[int64(uint32(v1))+4])
						m.memory[int64(uint32(v1))+3] = byte(t240)
						goto l354
					case 2:
						m.memory[int64(uint32(v1))+2] = byte(i32(1))
						m.memory[int64(uint32(v0))+6] = byte(i32(1))
						store16(m.memory[int64(uint32(v0))+4:], uint16(i32(258)))
						t241 := int32(m.memory[int64(uint32(v1))+4])
						m.memory[int64(uint32(v1))+3] = byte(t241)
						goto l354
					}
				}
				goto l307
			l319:
				m.memory[int64(uint32(v1))+4] = byte(v15)
				m.memory[int64(uint32(v1))+3] = byte(v15)
				v10 = i32(0)
				m.memory[int64(uint32(v1))+5] = byte(i32(0))
				m.memory[int64(uint32(v1))+1] = byte(i32(1))
				v11 = v17 & i32(1)
				v12 = v14
				v17 = i32(1)
				if v11 == 0 {
					goto l322
				}
				m.memory[int64(uint32(v0))+6] = byte(i32(3))
				store16(m.memory[int64(uint32(v0))+4:], uint16(i32(770)))
				goto l300
			l346:
				v15 = i32(3)
			l325:
				v16 = v18
				v12 = v14
				goto l322
			l344:
				v15 = i32(3)
				v16 = v16 + i32(3)
				v12 = v14
				goto l322
			l320:
				v15 = i32(5)
				m.memory[int64(uint32(v1))+3] = byte(i32(5))
				v12 = v14
				goto l322
			l307:
				m.memory[int64(uint32(v0))+4] = byte(i32(0))
			l354:
				v14 = v3
			}
		l300:
			store32(m.memory[int64(uint32(v0))+8:], uint32(v16))
			store32(m.memory[uint32(v0):], uint32(v14))
			goto l35
		case 4:
			store32(m.memory[int64(uint32(v7))+348:], uint32(i32(0)))
			store32(m.memory[int64(uint32(v7))+344:], uint32(v5))
			store32(m.memory[int64(uint32(v7))+340:], uint32(v4))
			t248 := int32(m.memory[int64(uint32(v1))+2])
			v10 = t248
			t249 := int32(m.memory[int64(uint32(v1))+1])
			v16 = t249
			v11 = i32(0)
			v12 = i32(0)
		l443:
			{
				v16 = v16 & i32(255)
				if v16 != 0 {
					if v12 != v3 {
						v18 = v11 + i32(2)
						if uint32(v18) < uint32(v5) {
							m.memory[int64(uint32(v1))+1] = byte(i32(0))
							v14 = v12 + i32(1)
							t297 := int32(int8(m.memory[uint32(v2+v12)]))
							v9 = t297
							v15 = v9 + i32(95)
							switch v16 + i32(-1) {
							default:
								v16 = v10 & i32(255)
								if v16 != i32(3) {
									goto l423
								}
								if uint32(v15&i32(255)) < uint32(i32(83)) {
									v12 = v4 + v11
									m.memory[uint32(v12)] = byte(i32(227))
									t310 := v4 + v18
									v15 = v9 + i32(-96)
									m.memory[uint32(t310)] = byte(v15&i32(63) | i32(128))
									m.memory[uint32(v12+i32(1))] = byte(int32(uint32(v15&i32(192))>>6) | i32(128))
									v10 = i32(3)
									v11 = v11 + i32(3)
									goto l454
								}
							l423:
								if v16 != i32(4) {
									goto l425
								}
								v17 = v15 & i32(255)
								if uint32(v17) < uint32(i32(86)) {
									v12 = v4 + v11
									m.memory[uint32(v12)] = byte(i32(227))
									t311 := v12 + i32(1)
									v12 = v17 + i32(12449)
									m.memory[uint32(t311)] = byte(int32(uint32(v12)>>6) & i32(135))
									m.memory[uint32(v4+v18)] = byte(v12&i32(63) | i32(128))
									t312 := v7
									v11 = v11 + i32(3)
									store32(m.memory[int64(uint32(t312))+348:], uint32(v11))
									v16 = i32(0)
									v10 = i32(4)
									v12 = v14
									goto l443
								}
							l425:
								v15 = v15 & i32(255)
								if uint32(v15) > uint32(i32(93)) {
									m.memory[int64(uint32(v0))+4] = byte(i32(2))
									if v9 > i32(-1) {
										store32(m.memory[int64(uint32(v0))+8:], uint32(v11))
										store32(m.memory[uint32(v0):], uint32(v12))
										store16(m.memory[int64(uint32(v0))+5:], uint16(i32(1)))
										goto l35
									}
									store32(m.memory[int64(uint32(v0))+8:], uint32(v11))
									store32(m.memory[uint32(v0):], uint32(v14))
									store16(m.memory[int64(uint32(v0))+5:], uint16(i32(2)))
									goto l35
								}
								v16 = v16*i32(94) + v15
								v12 = v16 + i32(-1410)
								if uint32(v12) < uint32(i32(2965)) {
									v15 = v4 + v11
									t313 := int32(load16(m.memory[int64(uint32(v12<<1))+1244278:]))
									t314 := v15
									v12 = t313
									m.memory[uint32(t314)] = byte(int32(uint32(v12)>>12) | i32(224))
									m.memory[uint32(v4+v18)] = byte(v12&i32(63) | i32(128))
									m.memory[uint32(v15+i32(1))] = byte(int32(uint32(v12)>>6)&i32(63) | i32(128))
									goto l435
								}
								v12 = v16 + i32(-4418)
								if uint32(v12) < uint32(i32(3390)) {
									v15 = v4 + v11
									t315 := int32(load16(m.memory[int64(uint32(v12<<1))+1272642:]))
									t316 := v15
									v12 = t315
									m.memory[uint32(t316)] = byte(int32(uint32(v12)>>12) | i32(224))
									m.memory[uint32(v4+v18)] = byte(v12&i32(63) | i32(128))
									m.memory[uint32(v15+i32(1))] = byte(int32(uint32(v12)>>6)&i32(63) | i32(128))
									goto l435
								}
								{
									v12 = v16 + i32(-8272)
									if uint32(v12) < uint32(i32(360)) {
										v15 = v4 + v11
										t301 := int32(load16(m.memory[int64(uint32(v12<<1))+1279422:]))
										t302 := v15
										v12 = t301
										m.memory[uint32(t302)] = byte(int32(uint32(v12)>>12) | i32(224))
										m.memory[uint32(v4+v18)] = byte(v12&i32(63) | i32(128))
										m.memory[uint32(v15+i32(1))] = byte(int32(uint32(v12)>>6)&i32(63) | i32(128))
										goto l435
									}
									m.fn1712(v7+i32(184), v16)
									t298 := int32(load16(m.memory[int64(uint32(v7))+184:]))
									if t298&i32(1) != 0 {
										t303 := int32(load16(m.memory[int64(uint32(v7))+186:]))
										v12 = t303
										if uint32(v12&i32(0xffff)) < uint32(i32(2048)) {
											v11 = v4 + v11
											m.memory[uint32(v11+i32(1))] = byte(v12&i32(63) | i32(128))
											m.memory[uint32(v11)] = byte(int32(uint32(v12)>>6) | i32(192))
											goto l437
										}
										m.memory[uint32(v4+v18)] = byte(v12&i32(63) | i32(128))
										v15 = v4 + v11
										m.memory[uint32(v15+i32(1))] = byte(int32(uint32(v12)>>6)&i32(63) | i32(128))
										m.memory[uint32(v15)] = byte(int32(uint32(v12&i32(61440))>>12) | i32(224))
										goto l435
									}
									v15 = i32(0)
									v12 = i32(1250730)
								l434:
									{
										if uint32(v15) > uint32(i32(53)) {
											m.memory[int64(uint32(v0))+6] = byte(i32(0))
											store16(m.memory[int64(uint32(v0))+4:], uint16(i32(514)))
											goto l440
										}
										t299 := int32(load16(m.memory[uint32(v12):]))
										v9 = v16 - t299
										t300 := int32(load16(m.memory[uint32(v12+i32(2)):]))
										if uint32(v9) < uint32(t300) {
											v12 = v15 + i32(2)
											if uint32(v15) >= uint32(i32(52)) {
												m.fn158(v12, i32(54), i32(1250208))
												panic("unreachable")
											}
											{
												t304 := int32(load16(m.memory[int64(uint32(v12<<1))+1250730:]))
												v12 = t304 + v9
												if uint32(v12&i32(0xffff)) < uint32(i32(2048)) {
													v11 = v4 + v11
													m.memory[uint32(v11+i32(1))] = byte(v12&i32(63) | i32(128))
													m.memory[uint32(v11)] = byte(int32(uint32(v12)>>6) | i32(192))
													goto l437
												}
												m.memory[uint32(v4+v18)] = byte(v12&i32(63) | i32(128))
												v15 = v4 + v11
												m.memory[uint32(v15+i32(1))] = byte(int32(uint32(v12)>>6)&i32(63) | i32(128))
												m.memory[uint32(v15)] = byte(int32(uint32(v12&i32(61440))>>12) | i32(224))
												goto l435
											}
										}
										v12 = v12 + i32(6)
										v15 = v15 + i32(3)
										goto l434
									}
								}
							case 1:
								if uint32(v15&i32(255)) > uint32(i32(93)) {
									m.memory[int64(uint32(v0))+4] = byte(i32(2))
									if v9 > i32(-1) {
										store32(m.memory[int64(uint32(v0))+8:], uint32(v11))
										store32(m.memory[uint32(v0):], uint32(v12))
										store16(m.memory[int64(uint32(v0))+5:], uint16(i32(1)))
										goto l35
									}
									store32(m.memory[int64(uint32(v0))+8:], uint32(v11))
									store32(m.memory[uint32(v0):], uint32(v14))
									store16(m.memory[int64(uint32(v0))+5:], uint16(i32(2)))
									goto l35
								}
								m.memory[int64(uint32(v1))+2] = byte(v15)
								v16 = i32(3)
								m.memory[int64(uint32(v1))+1] = byte(i32(3))
								v10 = v15
								v12 = v14
								goto l443
							case 2:
								v15 = v15 & i32(255)
								if uint32(v15) > uint32(i32(93)) {
									m.memory[int64(uint32(v0))+4] = byte(i32(2))
									if v9 > i32(-1) {
										store32(m.memory[int64(uint32(v0))+8:], uint32(v11))
										store32(m.memory[uint32(v0):], uint32(v12))
										store16(m.memory[int64(uint32(v0))+5:], uint16(i32(2)))
										goto l35
									}
									store32(m.memory[int64(uint32(v0))+8:], uint32(v11))
									store32(m.memory[uint32(v0):], uint32(v14))
									store16(m.memory[int64(uint32(v0))+5:], uint16(i32(3)))
									goto l35
								}
								v12 = v10&i32(255)*i32(94) + v15
								v15 = v12 + i32(-1410)
								if uint32(v15) < uint32(i32(5801)) {
									v16 = v4 + v11
									t317 := int32(load16(m.memory[int64(uint32(v15<<1))+1207162:]))
									t318 := v16
									v12 = t317
									m.memory[uint32(t318)] = byte(int32(uint32(v12)>>12) | i32(224))
									m.memory[uint32(v4+v18)] = byte(v12&i32(63) | i32(128))
									m.memory[uint32(v16+i32(1))] = byte(int32(uint32(v12)>>6)&i32(63) | i32(128))
									goto l435
								}
								m.fn1713(v7+i32(192), v12)
								{
									t305 := int32(load16(m.memory[int64(uint32(v7))+192:]))
									if t305&i32(1) == 0 {
										v15 = v12 + i32(-597)
										if uint32(v15) < uint32(i32(11)) {
											m.memory[uint32(v4+v11)] = byte(i32(208))
											t319 := int32(load32(m.memory[int64(uint32(v7))+340:]))
											v4 = t319
											m.memory[uint32(v4+v11+i32(1))] = byte(v15 + i32(-126))
											goto l437
										}
										v12 = v12 + i32(-645)
										if uint32(v12) < uint32(i32(11)) {
											m.memory[uint32(v4+v11)] = byte(i32(209))
											t307 := int32(load32(m.memory[int64(uint32(v7))+340:]))
											v4 = t307
											m.memory[uint32(v4+v11+i32(1))] = byte(v12 + i32(-110))
											goto l437
										}
										m.memory[int64(uint32(v0))+6] = byte(i32(0))
										store16(m.memory[int64(uint32(v0))+4:], uint16(i32(770)))
										goto l440
									}
									{
										t306 := int32(load16(m.memory[int64(uint32(v7))+194:]))
										v12 = t306
										if uint32(v12&i32(0xffff)) < uint32(i32(2048)) {
											v11 = v4 + v11
											m.memory[uint32(v11+i32(1))] = byte(v12&i32(63) | i32(128))
											m.memory[uint32(v11)] = byte(int32(uint32(v12)>>6) | i32(192))
											goto l437
										}
										m.memory[uint32(v4+v18)] = byte(v12&i32(63) | i32(128))
										v15 = v4 + v11
										m.memory[uint32(v15+i32(1))] = byte(int32(uint32(v12)>>6)&i32(63) | i32(128))
										m.memory[uint32(v15)] = byte(int32(uint32(v12&i32(61440))>>12) | i32(224))
										goto l435
									}
								}
							case 3:
								v15 = v15 & i32(255)
								if uint32(v15) > uint32(i32(62)) {
									m.memory[int64(uint32(v0))+4] = byte(i32(2))
									if v9 > i32(-1) {
										store32(m.memory[int64(uint32(v0))+8:], uint32(v11))
										store32(m.memory[uint32(v0):], uint32(v12))
										store16(m.memory[int64(uint32(v0))+5:], uint16(i32(1)))
										goto l35
									}
									store32(m.memory[int64(uint32(v0))+8:], uint32(v11))
									store32(m.memory[uint32(v0):], uint32(v14))
									store16(m.memory[int64(uint32(v0))+5:], uint16(i32(2)))
									goto l35
								}
								v12 = v4 + v11
								m.memory[uint32(v12)] = byte(i32(239))
								t308 := v12 + i32(1)
								v12 = v15 + i32(-159)
								m.memory[uint32(t308)] = byte(int32(uint32(v12)>>6) & i32(191))
								m.memory[uint32(v4+v18)] = byte(v12&i32(63) | i32(128))
								goto l435
							}
						}
						store32(m.memory[int64(uint32(v0))+8:], uint32(v11))
						store32(m.memory[uint32(v0):], uint32(v12))
						m.memory[int64(uint32(v0))+4] = byte(i32(1))
						goto l35
					}
					if v6 != 0 {
						m.memory[int64(uint32(v0))+6] = byte(i32(0))
						m.memory[int64(uint32(v0))+4] = byte(i32(2))
						store32(m.memory[int64(uint32(v0))+8:], uint32(v11))
						store32(m.memory[uint32(v0):], uint32(v3))
						t309 := int32(m.memory[int64(uint32(v1))+1])
						v3 = t309
						m.memory[int64(uint32(v1))+1] = byte(i32(0))
						m.memory[int64(uint32(v0))+5] = byte(i64_shr_u(i64(4328587520), int64(uint32(v3<<3))))
						goto l35
					}
					store32(m.memory[int64(uint32(v0))+8:], uint32(v11))
					store32(m.memory[uint32(v0):], uint32(v3))
					m.memory[int64(uint32(v0))+4] = byte(i32(0))
					goto l35
				l435:
					v11 = v11 + i32(3)
				l454:
					store32(m.memory[int64(uint32(v7))+348:], uint32(v11))
					v16 = i32(0)
					v12 = v14
					goto l443
				l440:
					store32(m.memory[int64(uint32(v0))+8:], uint32(v11))
					store32(m.memory[uint32(v0):], uint32(v14))
					goto l35
				l437:
					store32(m.memory[int64(uint32(v7))+348:], uint32(v18))
					v16 = i32(0)
					v11 = v18
					v12 = v14
					goto l443
				}
				t250 := int32(load32(m.memory[int64(uint32(v7))+344:]))
				v9 = t250
			l415:
				{
					m.fn148(v7+i32(176), v12, v2, v3, i32(1155464))
					t251 := int32(load32(m.memory[int64(uint32(v7))+176:]))
					v16 = t251
					t252 := int32(load32(m.memory[int64(uint32(v7))+180:]))
					v5 = t252
					m.fn212(v7+i32(168), v11, v4, v9, i32(1155480))
					t253 := int32(load32(m.memory[int64(uint32(v7))+172:]))
					v15 = t253
					t254 := v15
					t255 := v5
					var p256 int32
					if uint32(v15) < uint32(v5) {
						p256 = 1
					}
					v13 = p256
					p257 := t255
					if v13 != 0 {
						p257 = t254
					}
					v10 = p257
					v5 = i32(0)
					{
						{
							t258 := int32(load32(m.memory[int64(uint32(v7))+168:]))
							t259 := v16
							v14 = t258
							if (t259^v14)&i32(3) != 0 {
								goto l356
							}
							v5 = i32(0)
							v18 = (i32(0) - v16) & i32(3)
							if uint32(v18|i32(8)) > uint32(v10) {
								goto l356
							}
							v5 = i32(0)
						l363:
							if v18 != v5 {
								t266 := int32(int8(m.memory[uint32(v16+v5)]))
								v15 = t266
								if v15 < i32(0) {
									goto l361
								}
								m.memory[uint32(v14+v5)] = byte(v15)
								v5 = v5 + i32(1)
								goto l363
							}
							v8 = v10 + i32(-8)
						l362:
							{
								v18 = v16 + v5
								t260 := int32(load32(m.memory[uint32(v18):]))
								v15 = t260
								v17 = v14 + v5
								t261 := int32(load32(m.memory[uint32(v18+i32(4)):]))
								t262 := v17 + i32(4)
								v18 = t261
								store32(m.memory[uint32(t262):], uint32(v18))
								store32(m.memory[uint32(v17):], uint32(v15))
								{
									v18 = v18 & i32(-2139062144)
									t263 := v18
									v15 = v15 & i32(-2139062144)
									if t263|v15 == 0 {
										v5 = v5 + i32(8)
										if uint32(v5) <= uint32(v8) {
											goto l362
										}
										goto l356
									}
									if v15 != 0 {
										goto l359
									}
									v15 = int32(uint32(int32(bits.TrailingZeros32(uint32(v18))))>>3) + i32(4)
									goto l360
								l359:
									v15 = int32(uint32(int32(bits.TrailingZeros32(uint32(v15)))) >> 3)
								l360:
									t264 := v16
									v5 = v15 + v5
									t265 := int32(m.memory[uint32(t264+v5)])
									v15 = t265
									goto l361
								}
							}
						}
					l356:
						p267 := v10
						if uint32(v5) > uint32(v10) {
							p267 = v5
						}
						v18 = p267
					l365:
						{
							if v18 == v5 {
								v5 = v10 + v11
								v12 = v10 + v12
								goto l366
							}
							t268 := int32(int8(m.memory[uint32(v16+v5)]))
							v15 = t268
							if v15 < i32(0) {
								goto l361
							}
							m.memory[uint32(v14+v5)] = byte(v15)
							v5 = v5 + i32(1)
							goto l365
						}
					}
				l361:
					v12 = v5 + v12
					v5 = v5 + v11
					if uint32(v5+i32(2)) < uint32(v9) {
						goto l367
					}
					v13 = i32(1)
				l366:
					store32(m.memory[int64(uint32(v0))+8:], uint32(v5))
					store32(m.memory[uint32(v0):], uint32(v12))
					store32(m.memory[int64(uint32(v0))+4:], uint32(v13|(v7+i32(340))&i32(-256)))
					goto l35
				l367:
					v12 = v12 + i32(1)
					{
					l414:
						{
							v11 = v15 + i32(95)
							v16 = v11 & i32(255)
							if uint32(v16) < uint32(i32(94)) {
								if uint32(v12) < uint32(v3) {
									v11 = v12 + i32(1)
									t278 := int32(int8(m.memory[uint32(v2+v12)]))
									v14 = t278
									v15 = v14 + i32(95)
									if v16 != i32(3) {
										goto l396
									}
									if uint32(v15&i32(255)) < uint32(i32(83)) {
										v12 = v4 + v5
										m.memory[uint32(v12)] = byte(i32(227))
										t287 := v12 + i32(2)
										v15 = v14 + i32(-96)
										m.memory[uint32(t287)] = byte(v15&i32(63) | i32(128))
										m.memory[uint32(v12+i32(1))] = byte(int32(uint32(v15&i32(192))>>6) | i32(128))
										goto l383
									}
								l396:
									if v16 != i32(4) {
										goto l398
									}
									v18 = v15 & i32(255)
									if uint32(v18) < uint32(i32(86)) {
										v12 = v4 + v5
										m.memory[uint32(v12)] = byte(i32(227))
										t288 := v12 + i32(2)
										v15 = v18 + i32(12449)
										m.memory[uint32(t288)] = byte(v15&i32(63) | i32(128))
										m.memory[uint32(v12+i32(1))] = byte(int32(uint32(v15)>>6) & i32(135))
										goto l383
									}
								l398:
									v15 = v15 & i32(255)
									if uint32(v15) > uint32(i32(93)) {
										m.memory[int64(uint32(v0))+4] = byte(i32(2))
										if v14 > i32(-1) {
											store32(m.memory[int64(uint32(v0))+8:], uint32(v5))
											store32(m.memory[uint32(v0):], uint32(v12))
											store16(m.memory[int64(uint32(v0))+5:], uint16(i32(1)))
											goto l35
										}
										store32(m.memory[int64(uint32(v0))+8:], uint32(v5))
										store32(m.memory[uint32(v0):], uint32(v11))
										store16(m.memory[int64(uint32(v0))+5:], uint16(i32(2)))
										goto l35
									}
									v16 = v16*i32(94) + v15
									v12 = v16 + i32(-1410)
									if uint32(v12) < uint32(i32(2965)) {
										v15 = v4 + v5
										t289 := int32(load16(m.memory[int64(uint32(v12<<1))+1244278:]))
										t290 := v15 + i32(2)
										v12 = t289
										m.memory[uint32(t290)] = byte(v12&i32(63) | i32(128))
										m.memory[uint32(v15)] = byte(int32(uint32(v12)>>12) | i32(224))
										m.memory[uint32(v15+i32(1))] = byte(int32(uint32(v12)>>6)&i32(63) | i32(128))
										goto l383
									}
									v12 = v16 + i32(-4418)
									if uint32(v12) < uint32(i32(3390)) {
										v15 = v4 + v5
										t291 := int32(load16(m.memory[int64(uint32(v12<<1))+1272642:]))
										t292 := v15 + i32(2)
										v12 = t291
										m.memory[uint32(t292)] = byte(v12&i32(63) | i32(128))
										m.memory[uint32(v15)] = byte(int32(uint32(v12)>>12) | i32(224))
										m.memory[uint32(v15+i32(1))] = byte(int32(uint32(v12)>>6)&i32(63) | i32(128))
										goto l383
									}
									{
										v12 = v16 + i32(-8272)
										if uint32(v12) < uint32(i32(360)) {
											v15 = v4 + v5
											t282 := int32(load16(m.memory[int64(uint32(v12<<1))+1279422:]))
											t283 := v15 + i32(2)
											v12 = t282
											m.memory[uint32(t283)] = byte(v12&i32(63) | i32(128))
											m.memory[uint32(v15)] = byte(int32(uint32(v12)>>12) | i32(224))
											m.memory[uint32(v15+i32(1))] = byte(int32(uint32(v12)>>6)&i32(63) | i32(128))
											goto l383
										}
										m.fn1712(v7+i32(152), v16)
										t279 := int32(load16(m.memory[int64(uint32(v7))+152:]))
										if t279&i32(1) != 0 {
											t284 := int32(load16(m.memory[int64(uint32(v7))+154:]))
											v12 = t284
											if uint32(v12&i32(0xffff)) < uint32(i32(2048)) {
												v15 = v4 + v5
												m.memory[uint32(v15+i32(1))] = byte(v12&i32(63) | i32(128))
												m.memory[uint32(v15)] = byte(int32(uint32(v12)>>6) | i32(192))
												v12 = i32(2)
												goto l384
											}
											v15 = v4 + v5
											m.memory[uint32(v15+i32(2))] = byte(v12&i32(63) | i32(128))
											m.memory[uint32(v15+i32(1))] = byte(int32(uint32(v12)>>6)&i32(63) | i32(128))
											m.memory[uint32(v15)] = byte(int32(uint32(v12&i32(61440))>>12) | i32(224))
											goto l383
										}
										v15 = i32(0)
										v12 = i32(1250730)
									l407:
										{
											if uint32(v15) > uint32(i32(53)) {
												m.memory[int64(uint32(v0))+6] = byte(i32(0))
												store16(m.memory[int64(uint32(v0))+4:], uint16(i32(514)))
												goto l387
											}
											t280 := int32(load16(m.memory[uint32(v12):]))
											v14 = v16 - t280
											t281 := int32(load16(m.memory[uint32(v12+i32(2)):]))
											if uint32(v14) < uint32(t281) {
												v12 = v15 + i32(2)
												if uint32(v15) >= uint32(i32(52)) {
													m.fn158(v12, i32(54), i32(1250208))
													panic("unreachable")
												}
												{
													t285 := int32(load16(m.memory[int64(uint32(v12<<1))+1250730:]))
													v12 = t285 + v14
													if uint32(v12&i32(0xffff)) < uint32(i32(2048)) {
														v15 = v4 + v5
														m.memory[uint32(v15+i32(1))] = byte(v12&i32(63) | i32(128))
														m.memory[uint32(v15)] = byte(int32(uint32(v12)>>6) | i32(192))
														v12 = i32(2)
														goto l384
													}
													v15 = v4 + v5
													m.memory[uint32(v15+i32(2))] = byte(v12&i32(63) | i32(128))
													m.memory[uint32(v15+i32(1))] = byte(int32(uint32(v12)>>6)&i32(63) | i32(128))
													m.memory[uint32(v15)] = byte(int32(uint32(v12&i32(61440))>>12) | i32(224))
													goto l383
												}
											}
											v12 = v12 + i32(6)
											v15 = v15 + i32(3)
											goto l407
										}
									}
								}
								if v6 != 0 {
									m.memory[int64(uint32(v0))+6] = byte(i32(0))
									store16(m.memory[int64(uint32(v0))+4:], uint16(i32(258)))
									goto l374
								}
								m.memory[int64(uint32(v1))+2] = byte(v11)
								m.memory[int64(uint32(v1))+1] = byte(i32(1))
								m.memory[int64(uint32(v0))+4] = byte(i32(0))
								goto l374
							}
							switch v15&i32(255) + i32(-142) {
							case 1:
								if uint32(v12) < uint32(v3) {
									v15 = v12 + i32(1)
									{
										t269 := int32(int8(m.memory[uint32(v2+v12)]))
										v11 = t269
										v14 = v11 + i32(95)
										v16 = v14 & i32(255)
										if uint32(v16) > uint32(i32(93)) {
											m.memory[int64(uint32(v0))+4] = byte(i32(2))
											if v11 > i32(-1) {
												store32(m.memory[int64(uint32(v0))+8:], uint32(v5))
												store32(m.memory[uint32(v0):], uint32(v12))
												store16(m.memory[int64(uint32(v0))+5:], uint16(i32(1)))
												goto l35
											}
											store32(m.memory[int64(uint32(v0))+8:], uint32(v5))
											store32(m.memory[uint32(v0):], uint32(v15))
											store16(m.memory[int64(uint32(v0))+5:], uint16(i32(2)))
											goto l35
										}
										if uint32(v15) < uint32(v3) {
											v11 = v12 + i32(2)
											{
												t270 := int32(int8(m.memory[uint32(v2+v15)]))
												v14 = t270
												v12 = (v14 + i32(95)) & i32(255)
												if uint32(v12) > uint32(i32(93)) {
													m.memory[int64(uint32(v0))+4] = byte(i32(2))
													if v14 > i32(-1) {
														store32(m.memory[int64(uint32(v0))+8:], uint32(v5))
														store32(m.memory[uint32(v0):], uint32(v15))
														store16(m.memory[int64(uint32(v0))+5:], uint16(i32(2)))
														goto l35
													}
													store32(m.memory[int64(uint32(v0))+8:], uint32(v5))
													store32(m.memory[uint32(v0):], uint32(v11))
													store16(m.memory[int64(uint32(v0))+5:], uint16(i32(3)))
													goto l35
												}
												v12 = v16*i32(94) + v12
												v15 = v12 + i32(-1410)
												if uint32(v15) < uint32(i32(5801)) {
													v12 = v4 + v5
													t276 := int32(load16(m.memory[int64(uint32(v15<<1))+1207162:]))
													t277 := v12 + i32(2)
													v15 = t276
													m.memory[uint32(t277)] = byte(v15&i32(63) | i32(128))
													m.memory[uint32(v12)] = byte(int32(uint32(v15)>>12) | i32(224))
													m.memory[uint32(v12+i32(1))] = byte(int32(uint32(v15)>>6)&i32(63) | i32(128))
													goto l383
												}
												m.fn1713(v7+i32(160), v12)
												{
													t271 := int32(load16(m.memory[int64(uint32(v7))+160:]))
													if t271&i32(1) == 0 {
														v15 = v12 + i32(-597)
														if uint32(v15) < uint32(i32(11)) {
															m.memory[uint32(v4+v5)] = byte(i32(208))
															t286 := int32(load32(m.memory[int64(uint32(v7))+340:]))
															v4 = t286
															m.memory[uint32(v4+v5+i32(1))] = byte(v15 + i32(-126))
															v12 = i32(2)
															goto l384
														}
														v12 = v12 + i32(-645)
														if uint32(v12) < uint32(i32(11)) {
															m.memory[uint32(v4+v5)] = byte(i32(209))
															t273 := int32(load32(m.memory[int64(uint32(v7))+340:]))
															v4 = t273
															m.memory[uint32(v4+v5+i32(1))] = byte(v12 + i32(-110))
															v12 = i32(2)
															goto l384
														}
														m.memory[int64(uint32(v0))+6] = byte(i32(0))
														store16(m.memory[int64(uint32(v0))+4:], uint16(i32(770)))
														goto l387
													}
													{
														t272 := int32(load16(m.memory[int64(uint32(v7))+162:]))
														v12 = t272
														if uint32(v12&i32(0xffff)) < uint32(i32(2048)) {
															v15 = v4 + v5
															m.memory[uint32(v15+i32(1))] = byte(v12&i32(63) | i32(128))
															m.memory[uint32(v15)] = byte(int32(uint32(v12)>>6) | i32(192))
															v12 = i32(2)
															goto l384
														}
														v15 = v4 + v5
														m.memory[uint32(v15+i32(2))] = byte(v12&i32(63) | i32(128))
														m.memory[uint32(v15+i32(1))] = byte(int32(uint32(v12)>>6)&i32(63) | i32(128))
														m.memory[uint32(v15)] = byte(int32(uint32(v12&i32(61440))>>12) | i32(224))
														goto l383
													}
												}
											}
										}
										if v6 != 0 {
											goto l377
										}
										m.memory[int64(uint32(v1))+2] = byte(v14)
										m.memory[int64(uint32(v1))+1] = byte(i32(3))
										m.memory[int64(uint32(v0))+4] = byte(i32(0))
										goto l378
									l377:
										m.memory[int64(uint32(v0))+6] = byte(i32(0))
										store16(m.memory[int64(uint32(v0))+4:], uint16(i32(514)))
									l378:
										store32(m.memory[int64(uint32(v0))+8:], uint32(v5))
										store32(m.memory[uint32(v0):], uint32(v15))
										goto l35
									}
								}
								if v6 != 0 {
									m.memory[int64(uint32(v0))+6] = byte(i32(0))
									store16(m.memory[int64(uint32(v0))+4:], uint16(i32(258)))
									goto l374
								}
								m.memory[int64(uint32(v0))+4] = byte(i32(0))
								m.memory[int64(uint32(v1))+1] = byte(i32(2))
								goto l374
							default:
								m.memory[int64(uint32(v0))+6] = byte(i32(0))
								store16(m.memory[int64(uint32(v0))+4:], uint16(i32(258)))
								goto l374
							case 0:
								if uint32(v12) < uint32(v3) {
									v11 = v12 + i32(1)
									t274 := int32(int8(m.memory[uint32(v2+v12)]))
									v16 = t274
									v15 = (v16 + i32(95)) & i32(255)
									if uint32(v15) > uint32(i32(62)) {
										m.memory[int64(uint32(v0))+4] = byte(i32(2))
										if v16 > i32(-1) {
											store32(m.memory[int64(uint32(v0))+8:], uint32(v5))
											store32(m.memory[uint32(v0):], uint32(v12))
											store16(m.memory[int64(uint32(v0))+5:], uint16(i32(1)))
											goto l35
										}
										store32(m.memory[int64(uint32(v0))+8:], uint32(v5))
										store32(m.memory[uint32(v0):], uint32(v11))
										store16(m.memory[int64(uint32(v0))+5:], uint16(i32(2)))
										goto l35
									}
									v12 = v4 + v5
									m.memory[uint32(v12)] = byte(i32(239))
									t275 := v12 + i32(2)
									v15 = v15 + i32(-159)
									m.memory[uint32(t275)] = byte(v15&i32(63) | i32(128))
									m.memory[uint32(v12+i32(1))] = byte(int32(uint32(v15)>>6) & i32(191))
									goto l383
								}
								if v6 != 0 {
									m.memory[int64(uint32(v0))+6] = byte(i32(0))
									store16(m.memory[int64(uint32(v0))+4:], uint16(i32(258)))
									goto l374
								}
								m.memory[int64(uint32(v0))+4] = byte(i32(0))
								m.memory[int64(uint32(v1))+1] = byte(i32(4))
								goto l374
							}
						l383:
							v12 = i32(3)
						l384:
							v5 = v12 + v5
							if uint32(v11) < uint32(v3) {
								goto l412
							}
							store32(m.memory[int64(uint32(v0))+8:], uint32(v5))
							store32(m.memory[uint32(v0):], uint32(v11))
							m.memory[int64(uint32(v0))+4] = byte(i32(0))
							goto l35
						l412:
							{
								t293 := int32(load32(m.memory[int64(uint32(v7))+344:]))
								t294 := v5 + i32(2)
								v9 = t293
								if uint32(t294) < uint32(v9) {
									goto l413
								}
								store32(m.memory[int64(uint32(v0))+8:], uint32(v5))
								store32(m.memory[uint32(v0):], uint32(v11))
								m.memory[int64(uint32(v0))+4] = byte(i32(1))
								goto l35
							}
						l413:
							v12 = v11 + i32(1)
							t295 := int32(int8(m.memory[uint32(v2+v11)]))
							v15 = t295
							if v15 < i32(0) {
								goto l414
							}
						}
						m.memory[uint32(v4+v5)] = byte(v15)
						t296 := v7
						v11 = v5 + i32(1)
						store32(m.memory[int64(uint32(t296))+348:], uint32(v11))
						goto l415
					}
				l387:
				}
				store32(m.memory[int64(uint32(v0))+8:], uint32(v5))
				store32(m.memory[uint32(v0):], uint32(v11))
				goto l35
			l374:
				store32(m.memory[int64(uint32(v0))+8:], uint32(v5))
				store32(m.memory[uint32(v0):], uint32(v12))
				goto l35
			}
		case 3:
			store32(m.memory[int64(uint32(v7))+344:], uint32(v5))
			store32(m.memory[int64(uint32(v7))+340:], uint32(v4))
			v10 = i32(0)
			v16 = i32(0)
			{
				t320 := int32(m.memory[int64(uint32(v1))+1])
				if t320 == 0 {
					goto l511
				}
				m.memory[int64(uint32(v1))+1] = byte(i32(0))
				if v3 == 0 {
					if v6 != 0 {
						m.memory[int64(uint32(v0))+6] = byte(i32(0))
						store16(m.memory[int64(uint32(v0))+4:], uint16(i32(258)))
						store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
						store32(m.memory[uint32(v0):], uint32(i32(0)))
						goto l35
					}
					store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
					store32(m.memory[uint32(v0):], uint32(i32(0)))
					m.memory[int64(uint32(v0))+4] = byte(i32(0))
					goto l35
				}
				if uint32(v5) > uint32(i32(3)) {
					t321 := int32(m.memory[int64(uint32(v1))+2])
					v15 = t321
					{
						t322 := int32(int8(m.memory[uint32(v2)]))
						v12 = t322
						v11 = v12 + i32(-64)
						if uint32(v11&i32(255)) < uint32(i32(63)) {
							goto l458
						}
						if uint32((v12+i32(1))&i32(255)) < uint32(i32(162)) {
							m.memory[int64(uint32(v0))+4] = byte(i32(2))
							if v12 > i32(-1) {
								store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
								store32(m.memory[uint32(v0):], uint32(i32(0)))
								store16(m.memory[int64(uint32(v0))+5:], uint16(i32(1)))
								goto l35
							}
							store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
							store32(m.memory[uint32(v0):], uint32(i32(1)))
							store16(m.memory[int64(uint32(v0))+5:], uint16(i32(2)))
							goto l35
						}
						v11 = v12 + i32(-98)
					}
				l458:
					{
						v15 = v15*i32(157) + v11&i32(255)
						v11 = v15 + i32(-942)
						if uint32(v11) > uint32(i32(18839)) {
							goto l460
						}
						t323 := int32(load16(m.memory[int64(uint32(v11<<1))+1169418:]))
						v16 = t323
						if v16 != 0 {
							v12 = i32(3)
							t324 := int32(load32(m.memory[int64(uint32(int32(uint32(v11)>>3)&i32(0x1ffffffc)))+1233428:]))
							if i32_shr_u(t324, v11)&i32(1) == 0 {
								if uint32(v16) < uint32(i32(2048)) {
									m.memory[uint32(v4)] = byte(int32(uint32(v16)>>6) | i32(192))
									v11 = v16&i32(63) | i32(-128)
									v10 = i32(2)
									v12 = i32(1)
									goto l472
								}
								m.memory[uint32(v4)] = byte(int32(uint32(v16)>>12) | i32(224))
								m.memory[int64(uint32(v4))+1] = byte(int32(uint32(v16)>>6)&i32(63) | i32(128))
								v11 = v16&i32(63) | i32(-128)
								v10 = i32(3)
								v12 = i32(2)
								goto l472
							}
							m.memory[uint32(v4)] = byte(i32(240))
							m.memory[int64(uint32(v4))+1] = byte(int32(uint32(v16)>>12) | i32(160))
							m.memory[int64(uint32(v4))+2] = byte(int32(uint32(v16)>>6)&i32(63) | i32(128))
							v11 = v16&i32(63) | i32(-128)
							v10 = i32(4)
							goto l472
						}
					}
				l460:
					switch v15 + i32(-1133) {
					case 1:
						goto l463
					default:
						switch v15 + i32(-1164) {
						case 0:
							m.memory[int64(uint32(v4))+2] = byte(i32(204))
							store16(m.memory[uint32(v4):], uint16(i32(43715)))
							goto l468
						case 2:
							m.memory[int64(uint32(v4))+2] = byte(i32(204))
							store16(m.memory[uint32(v4):], uint16(i32(43715)))
							goto l469
						default:
							goto l463
						}
					case 0:
						m.memory[int64(uint32(v4))+2] = byte(i32(204))
						store16(m.memory[uint32(v4):], uint16(i32(35523)))
						goto l468
					case 2:
						m.memory[int64(uint32(v4))+2] = byte(i32(204))
						store16(m.memory[uint32(v4):], uint16(i32(35523)))
						goto l469
					}
				l463:
					m.memory[int64(uint32(v0))+4] = byte(i32(2))
					if v12 > i32(-1) {
						store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
						store32(m.memory[uint32(v0):], uint32(i32(0)))
						store16(m.memory[int64(uint32(v0))+5:], uint16(i32(1)))
						goto l35
					}
					store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
					store32(m.memory[uint32(v0):], uint32(i32(1)))
					store16(m.memory[int64(uint32(v0))+5:], uint16(i32(2)))
					goto l35
				}
				store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
				store32(m.memory[uint32(v0):], uint32(i32(0)))
				m.memory[int64(uint32(v0))+4] = byte(i32(1))
				goto l35
			l469:
				v10 = i32(4)
				v11 = i32(140)
				v12 = i32(3)
				goto l472
			l468:
				v10 = i32(4)
				v11 = i32(132)
				v12 = i32(3)
			l472:
				m.memory[uint32(v4+v12)] = byte(v11)
				v16 = i32(1)
			}
		l511:
			{
				m.fn148(v7+i32(144), v16, v2, v3, i32(1155496))
				t325 := int32(load32(m.memory[int64(uint32(v7))+144:]))
				v15 = t325
				t326 := int32(load32(m.memory[int64(uint32(v7))+148:]))
				v12 = t326
				m.fn212(v7+i32(136), v10, v4, v5, i32(1155512))
				t327 := int32(load32(m.memory[int64(uint32(v7))+140:]))
				v11 = t327
				t328 := v11
				t329 := v12
				var p330 int32
				if uint32(v11) < uint32(v12) {
					p330 = 1
				}
				v13 = p330
				p331 := t329
				if v13 != 0 {
					p331 = t328
				}
				v9 = p331
				v12 = i32(0)
				{
					{
						t332 := int32(load32(m.memory[int64(uint32(v7))+136:]))
						t333 := v15
						v14 = t332
						if (t333^v14)&i32(3) != 0 {
							goto l476
						}
						v12 = i32(0)
						v18 = (i32(0) - v15) & i32(3)
						if uint32(v18|i32(8)) > uint32(v9) {
							goto l476
						}
						v12 = i32(0)
					l483:
						if v18 != v12 {
							t340 := int32(int8(m.memory[uint32(v15+v12)]))
							v11 = t340
							if v11 < i32(0) {
								goto l481
							}
							m.memory[uint32(v14+v12)] = byte(v11)
							v12 = v12 + i32(1)
							goto l483
						}
						v8 = v9 + i32(-8)
					l482:
						{
							v18 = v15 + v12
							t334 := int32(load32(m.memory[uint32(v18):]))
							v11 = t334
							v17 = v14 + v12
							t335 := int32(load32(m.memory[uint32(v18+i32(4)):]))
							t336 := v17 + i32(4)
							v18 = t335
							store32(m.memory[uint32(t336):], uint32(v18))
							store32(m.memory[uint32(v17):], uint32(v11))
							{
								v18 = v18 & i32(-2139062144)
								t337 := v18
								v11 = v11 & i32(-2139062144)
								if t337|v11 == 0 {
									v12 = v12 + i32(8)
									if uint32(v12) <= uint32(v8) {
										goto l482
									}
									goto l476
								}
								if v11 != 0 {
									goto l479
								}
								v11 = int32(uint32(int32(bits.TrailingZeros32(uint32(v18))))>>3) + i32(4)
								goto l480
							l479:
								v11 = int32(uint32(int32(bits.TrailingZeros32(uint32(v11)))) >> 3)
							l480:
								t338 := v15
								v12 = v11 + v12
								t339 := int32(m.memory[uint32(t338+v12)])
								v11 = t339
								goto l481
							}
						}
					}
				l476:
					p341 := v9
					if uint32(v12) > uint32(v9) {
						p341 = v12
					}
					v18 = p341
				l485:
					{
						if v18 == v12 {
							v12 = v9 + v10
							v15 = v9 + v16
							goto l486
						}
						t342 := int32(int8(m.memory[uint32(v15+v12)]))
						v11 = t342
						if v11 < i32(0) {
							goto l481
						}
						m.memory[uint32(v14+v12)] = byte(v11)
						v12 = v12 + i32(1)
						goto l485
					}
				}
			l481:
				v15 = v12 + v16
				v12 = v12 + v10
				if uint32(v12+i32(3)) < uint32(v5) {
					goto l487
				}
				v13 = i32(1)
			l486:
				store32(m.memory[int64(uint32(v0))+8:], uint32(v12))
				store32(m.memory[uint32(v0):], uint32(v15))
				store32(m.memory[int64(uint32(v0))+4:], uint32(v13|(v7+i32(340))&i32(-256)))
				goto l35
			l487:
				v16 = v15 + i32(1)
				{
				l510:
					{
						v5 = v11 + i32(127)
						v11 = v5 & i32(255)
						if uint32(v11) > uint32(i32(125)) {
							m.memory[int64(uint32(v0))+6] = byte(i32(0))
							store16(m.memory[int64(uint32(v0))+4:], uint16(i32(258)))
							goto l491
						}
						if uint32(v16) < uint32(v3) {
							v15 = v16 + i32(1)
							{
								v18 = v2 + v16
								t343 := int32(int8(m.memory[uint32(v18)]))
								v14 = t343
								v5 = v14 + i32(-64)
								if uint32(v5&i32(255)) < uint32(i32(63)) {
									goto l492
								}
								if uint32((v14+i32(1))&i32(255)) < uint32(i32(162)) {
									m.memory[int64(uint32(v0))+4] = byte(i32(2))
									if v14 > i32(-1) {
										store32(m.memory[int64(uint32(v0))+8:], uint32(v12))
										store32(m.memory[uint32(v0):], uint32(v16))
										store16(m.memory[int64(uint32(v0))+5:], uint16(i32(1)))
										goto l35
									}
									store32(m.memory[int64(uint32(v0))+8:], uint32(v12))
									store32(m.memory[uint32(v0):], uint32(v15))
									store16(m.memory[int64(uint32(v0))+5:], uint16(i32(2)))
									goto l35
								}
								v5 = v14 + i32(-98)
							}
						l492:
							{
								v11 = v11*i32(157) + v5&i32(255)
								v5 = v11 + i32(-942)
								if uint32(v5) > uint32(i32(18839)) {
									goto l494
								}
								t344 := int32(load16(m.memory[int64(uint32(v5<<1))+1169418:]))
								v9 = t344
								if v9 != 0 {
									t345 := int32(load32(m.memory[int64(uint32(int32(uint32(v5)>>3)&i32(0x1ffffffc)))+1233428:]))
									if i32_shr_u(t345, v5)&i32(1) == 0 {
										if uint32(v9) < uint32(i32(2048)) {
											v5 = v4 + v12
											m.memory[uint32(v5+i32(1))] = byte(v9&i32(63) | i32(128))
											m.memory[uint32(v5)] = byte(int32(uint32(v9)>>6) | i32(192))
											v5 = i32(2)
											goto l507
										}
										v5 = v4 + v12
										m.memory[uint32(v5+i32(2))] = byte(v9&i32(63) | i32(128))
										m.memory[uint32(v5)] = byte(int32(uint32(v9)>>12) | i32(224))
										m.memory[uint32(v5+i32(1))] = byte(int32(uint32(v9)>>6)&i32(63) | i32(128))
										v5 = i32(3)
										goto l507
									}
									m.memory[uint32(v4+v12)] = byte(i32(240))
									t346 := int32(load32(m.memory[int64(uint32(v7))+340:]))
									v4 = t346
									v5 = v4 + v12
									m.memory[uint32(v5+i32(3))] = byte(v9&i32(63) | i32(128))
									m.memory[uint32(v5+i32(2))] = byte(int32(uint32(v9)>>6)&i32(63) | i32(128))
									m.memory[uint32(v5+i32(1))] = byte(int32(uint32(v9)>>12) | i32(160))
									goto l502
								}
							}
						l494:
							switch v11 + i32(-1133) {
							case 1:
								goto l497
							default:
								switch v11 + i32(-1164) {
								case 0:
									store32(m.memory[uint32(v4+v12):], uint32(i32(-0x7b33553d)))
									goto l502
								case 2:
									store32(m.memory[uint32(v4+v12):], uint32(i32(-0x7333553d)))
									goto l502
								default:
									goto l497
								}
							case 0:
								store32(m.memory[uint32(v4+v12):], uint32(i32(-2066969917)))
								goto l502
							case 2:
								store32(m.memory[uint32(v4+v12):], uint32(i32(-0x7333753d)))
								goto l502
							}
						l497:
							m.memory[int64(uint32(v0))+4] = byte(i32(2))
							if v14 > i32(-1) {
								store32(m.memory[int64(uint32(v0))+8:], uint32(v12))
								store32(m.memory[uint32(v0):], uint32(v16))
								store16(m.memory[int64(uint32(v0))+5:], uint16(i32(1)))
								goto l35
							}
							store32(m.memory[int64(uint32(v0))+8:], uint32(v12))
							store32(m.memory[uint32(v0):], uint32(v15))
							store16(m.memory[int64(uint32(v0))+5:], uint16(i32(2)))
							goto l35
						}
						if v6 != 0 {
							m.memory[int64(uint32(v0))+6] = byte(i32(0))
							store16(m.memory[int64(uint32(v0))+4:], uint16(i32(258)))
							goto l491
						}
						m.memory[int64(uint32(v1))+2] = byte(v5)
						m.memory[int64(uint32(v1))+1] = byte(i32(1))
						m.memory[int64(uint32(v0))+4] = byte(i32(0))
						goto l491
					l502:
						v5 = i32(4)
					l507:
						v12 = v5 + v12
						if uint32(v15) < uint32(v3) {
							goto l508
						}
						store32(m.memory[int64(uint32(v0))+8:], uint32(v12))
						store32(m.memory[uint32(v0):], uint32(v15))
						m.memory[int64(uint32(v0))+4] = byte(i32(0))
						goto l35
					l508:
						{
							t347 := int32(load32(m.memory[int64(uint32(v7))+344:]))
							if uint32(v12+i32(3)) < uint32(t347) {
								goto l509
							}
							store32(m.memory[int64(uint32(v0))+8:], uint32(v12))
							store32(m.memory[uint32(v0):], uint32(v15))
							m.memory[int64(uint32(v0))+4] = byte(i32(1))
							goto l35
						}
					l509:
						v16 = v16 + i32(2)
						t348 := int32(int8(m.memory[uint32(v18+i32(1))]))
						v11 = t348
						if v11 < i32(0) {
							goto l510
						}
					}
					m.memory[uint32(v4+v12)] = byte(v11)
					v10 = v12 + i32(1)
					t349 := int32(load32(m.memory[int64(uint32(v7))+340:]))
					v4 = t349
					t350 := int32(load32(m.memory[int64(uint32(v7))+344:]))
					v5 = t350
					goto l511
				}
			l491:
			}
			store32(m.memory[int64(uint32(v0))+8:], uint32(v12))
			store32(m.memory[uint32(v0):], uint32(v16))
			goto l35
		case 2:
			store32(m.memory[int64(uint32(v7))+348:], uint32(i32(0)))
			store32(m.memory[int64(uint32(v7))+344:], uint32(v5))
			store32(m.memory[int64(uint32(v7))+340:], uint32(v4))
			v11 = i32(0)
			{
				{
					t351 := int32(m.memory[int64(uint32(v1))+7])
					if t351 != i32(1) {
						goto l512
					}
					if uint32(v5) > uint32(i32(2)) {
						goto l513
					}
					v11 = i32(0)
					store32(m.memory[uint32(v0):], uint32(i32(0)))
					m.memory[int64(uint32(v0))+4] = byte(i32(1))
					goto l514
				l513:
					m.memory[int64(uint32(v1))+7] = byte(i32(0))
					t352 := int32(m.memory[int64(uint32(v1))+8])
					m.memory[uint32(v4)] = byte(t352)
					v11 = i32(1)
					store32(m.memory[int64(uint32(v7))+348:], uint32(i32(1)))
				}
			l512:
				t353 := int32(m.memory[int64(uint32(v1))+12])
				v9 = t353
				t354 := int32(m.memory[int64(uint32(v1))+11])
				v10 = t354
				t355 := int32(m.memory[int64(uint32(v1))+9])
				v15 = t355
				v16 = i32(0)
			l601:
				v12 = v16
				{
					v14 = v15 & i32(255)
					if v14 != 0 {
						if v3 != v12 {
							v18 = v11 + i32(3)
							if uint32(v18) < uint32(v5) {
								v16 = v12 + i32(1)
								t423 := int32(int8(m.memory[uint32(v2+v12)]))
								v15 = t423
								switch v14 + i32(-1) {
								default:
									m.memory[int64(uint32(v1))+9] = byte(i32(0))
									v14 = v15 + i32(-48)
									if uint32(v14&i32(255)) > uint32(i32(9)) {
										t424 := int32(m.memory[int64(uint32(v1))+10])
										v14 = t424
										if uint32(v14) > uint32(i32(31)) {
											v17 = (v15 + i32(95)) & i32(255)
											if uint32(v17) < uint32(i32(94)) {
												{
													v12 = (v14 + i32(-47)) & i32(255)
													if uint32(v12) < uint32(i32(72)) {
														t449 := int32(load32(m.memory[int64(uint32(v7))+340:]))
														v4 = t449
														v11 = v4 + v11
														t450 := int32(load16(m.memory[int64(uint32(v12*i32(188)+v17<<1))+1155882:]))
														t451 := v11 + i32(2)
														v12 = t450
														m.memory[uint32(t451)] = byte(v12&i32(63) | i32(128))
														m.memory[uint32(v11)] = byte(int32(uint32(v12)>>12) | i32(224))
														m.memory[uint32(v11+i32(1))] = byte(int32(uint32(v12)>>6)&i32(63) | i32(128))
														goto l609
													}
													switch v14 + i32(-37) {
													case 0:
														v12 = (v15 + i32(32)) & i32(255)
														if uint32(v12) > uint32(i32(21)) {
															goto l633
														}
														t452 := int32(load32(m.memory[int64(uint32(v7))+340:]))
														v4 = t452
														v11 = v4 + v11
														t453 := int32(load16(m.memory[int64(uint32(v12<<1))+1261894:]))
														t454 := v11 + i32(2)
														v12 = t453
														m.memory[uint32(t454)] = byte(v12&i32(63) | i32(128))
														m.memory[uint32(v11)] = byte(int32(uint32(v12)>>12) | i32(224))
														m.memory[uint32(v11+i32(1))] = byte(int32(uint32(v12)>>6)&i32(63) | i32(128))
														goto l609
													case 1:
														goto l630
													case 2:
														goto l631
													default:
														goto l632
													}
												l632:
													if v14 == i32(32) {
														t479 := int32(load16(m.memory[int64(uint32(v17<<1))+1228540:]))
														v12 = t479
														if uint32(v12) < uint32(i32(2048)) {
															t481 := int32(load32(m.memory[int64(uint32(v7))+340:]))
															v4 = t481
															v15 = v4 + v11
															m.memory[uint32(v15+i32(1))] = byte(v12&i32(63) | i32(128))
															m.memory[uint32(v15)] = byte(int32(uint32(v12)>>6) | i32(192))
															goto l626
														}
														t480 := int32(load32(m.memory[int64(uint32(v7))+340:]))
														v4 = t480
														v11 = v4 + v11
														m.memory[uint32(v11+i32(2))] = byte(v12&i32(63) | i32(128))
														m.memory[uint32(v11)] = byte(int32(uint32(v12)>>12) | i32(224))
														m.memory[uint32(v11+i32(1))] = byte(int32(uint32(v12)>>6)&i32(63) | i32(128))
														goto l609
													}
												l630:
													if uint32(v14) <= uint32(i32(118)) {
														goto l633
													}
													t455 := int32(load32(m.memory[int64(uint32(v7))+340:]))
													v4 = t455
													v12 = v4 + v11
													t456 := v12 + i32(2)
													v11 = v17 + (v14+i32(-119))&i32(255)*i32(94) + i32(-7628)
													m.memory[uint32(t456)] = byte(v11&i32(63) | i32(128))
													m.memory[uint32(v12+i32(1))] = byte(int32(uint32(v11)>>6)&i32(63) | i32(128))
													m.memory[uint32(v12)] = byte(int32(uint32(v11&i32(61440))>>12) | i32(224))
													goto l609
												}
											l631:
												if uint32(v17) < uint32(i32(32)) {
													t476 := int32(load16(m.memory[int64(uint32(v17<<1))+1207098:]))
													v12 = t476
													{
														if v17 != i32(27) {
															t478 := int32(load32(m.memory[int64(uint32(v7))+340:]))
															v4 = t478
															v15 = v4 + v11
															m.memory[uint32(v15+i32(1))] = byte(v12&i32(63) | i32(128))
															m.memory[uint32(v15)] = byte(int32(uint32(v12)>>6) | i32(192))
															goto l626
														}
														t477 := int32(load32(m.memory[int64(uint32(v7))+340:]))
														v4 = t477
														v11 = v4 + v11
														m.memory[uint32(v11+i32(2))] = byte(v12&i32(63) | i32(128))
														m.memory[uint32(v11+i32(1))] = byte(int32(uint32(v12)>>6)&i32(63) | i32(128))
														m.memory[uint32(v11)] = byte(int32(uint32(v12&i32(61440))>>12) | i32(224))
														goto l609
													}
												}
											l633:
												t457 := v7 + i32(104)
												v12 = (v14+i32(-33))&i32(255)*i32(94) + v17
												m.fn1711(t457, i32(1250636), i32(46), v12)
												t458 := int32(load32(m.memory[int64(uint32(v7))+108:]))
												v4 = t458
												{
													{
														t459 := int32(load32(m.memory[int64(uint32(v7))+104:]))
														if t459 != i32(1) {
															if uint32(v4) > uint32(i32(45)) {
																m.fn158(v4, i32(46), i32(1236596))
																panic("unreachable")
															}
															t460 := int32(load16(m.memory[int64(uint32(v4<<1))+1272396:]))
															v12 = t460
															goto l639
														}
														v4 = v4 + i32(-1)
														if uint32(v4) < uint32(i32(46)) {
															goto l637
														}
														m.fn158(v4, i32(46), i32(1236612))
														panic("unreachable")
													}
												l637:
													v4 = v4 << 1
													t461 := int32(load16(m.memory[int64(uint32(v4))+1272396:]))
													t462 := int32(load16(m.memory[int64(uint32(v4))+1250636:]))
													v12 = t461 + v12 - t462
												}
											l639:
												{
													if uint32(v12&i32(0xffff)) < uint32(i32(2048)) {
														t464 := int32(load32(m.memory[int64(uint32(v7))+340:]))
														v4 = t464
														v15 = v4 + v11
														m.memory[uint32(v15+i32(1))] = byte(v12&i32(63) | i32(128))
														m.memory[uint32(v15)] = byte(int32(uint32(v12)>>6) | i32(192))
														goto l626
													}
													t463 := int32(load32(m.memory[int64(uint32(v7))+340:]))
													v4 = t463
													v11 = v4 + v11
													m.memory[uint32(v11+i32(2))] = byte(v12&i32(63) | i32(128))
													m.memory[uint32(v11+i32(1))] = byte(int32(uint32(v12)>>6)&i32(63) | i32(128))
													m.memory[uint32(v11)] = byte(int32(uint32(v12&i32(61440))>>12) | i32(224))
													goto l609
												}
											}
											v4 = v15 + i32(-64)
											if uint32(v4&i32(255)) <= uint32(i32(62)) {
												goto l612
											}
											if v15 > i32(-96) {
												m.memory[int64(uint32(v0))+4] = byte(i32(2))
												if v15 > i32(-1) {
													store32(m.memory[uint32(v0):], uint32(v12))
													store16(m.memory[int64(uint32(v0))+5:], uint16(i32(1)))
													goto l514
												}
												store32(m.memory[uint32(v0):], uint32(v16))
												store16(m.memory[int64(uint32(v0))+5:], uint16(i32(2)))
												goto l514
											}
											v4 = v15 + i32(-65)
										l612:
											{
												v4 = (v14+i32(-32))&i32(255)*i32(96) + v4&i32(255)
												v12 = v4 + i32(-864)
												if uint32(v12) < uint32(i32(8059)) {
													m.fn1711(v7+i32(112), i32(1269142), i32(1627), v12)
													t435 := int32(load32(m.memory[int64(uint32(v7))+116:]))
													v4 = t435
													{
														{
															t436 := int32(load32(m.memory[int64(uint32(v7))+112:]))
															if t436 != i32(1) {
																if uint32(v4) > uint32(i32(1626)) {
																	m.fn158(v4, i32(1627), i32(1236596))
																	panic("unreachable")
																}
																t437 := int32(load16(m.memory[int64(uint32(v4<<1))+1261938:]))
																v15 = t437
																v14 = int32(uint32(v15) >> 12)
																v17 = int32(uint32(v15) >> 6)
																goto l620
															}
															v4 = v4 + i32(-1)
															if uint32(v4) < uint32(i32(1627)) {
																goto l618
															}
															m.fn158(v4, i32(1627), i32(1236612))
															panic("unreachable")
														}
													l618:
														t438 := v12
														v4 = v4 << 1
														t439 := int32(load16(m.memory[int64(uint32(v4))+1269142:]))
														t440 := int32(load16(m.memory[int64(uint32(v4))+1261938:]))
														v15 = t438 - t439 + t440
														v4 = v15 & i32(0xffff)
														v14 = int32(uint32(v4) >> 12)
														v17 = int32(uint32(v4) >> 6)
													}
												l620:
													t441 := int32(load32(m.memory[int64(uint32(v7))+340:]))
													v4 = t441
													v12 = v4 + v11
													m.memory[uint32(v12)] = byte(v14 | i32(224))
													m.memory[uint32(v12+i32(2))] = byte(v15&i32(63) | i32(128))
													m.memory[uint32(v12+i32(1))] = byte(v17&i32(63) | i32(128))
													goto l609
												}
												if uint32(v4) < uint32(i32(864)) {
													m.fn1711(v7+i32(120), i32(1244158), i32(59), v4)
													t442 := int32(load32(m.memory[int64(uint32(v7))+124:]))
													v12 = t442
													{
														{
															t443 := int32(load32(m.memory[int64(uint32(v7))+120:]))
															if t443 != i32(1) {
																if uint32(v12) > uint32(i32(58)) {
																	m.fn158(v12, i32(59), i32(1236596))
																	panic("unreachable")
																}
																t444 := int32(load16(m.memory[int64(uint32(v12<<1))+1265192:]))
																v12 = t444
																goto l624
															}
															v12 = v12 + i32(-1)
															if uint32(v12) < uint32(i32(59)) {
																goto l622
															}
															m.fn158(v12, i32(59), i32(1236612))
															panic("unreachable")
														}
													l622:
														v12 = v12 << 1
														t445 := int32(load16(m.memory[int64(uint32(v12))+1265192:]))
														t446 := int32(load16(m.memory[int64(uint32(v12))+1244158:]))
														v12 = t445 + v4 - t446
													}
												l624:
													{
														if uint32(v12&i32(0xffff)) < uint32(i32(2048)) {
															t448 := int32(load32(m.memory[int64(uint32(v7))+340:]))
															v4 = t448
															v15 = v4 + v11
															m.memory[uint32(v15+i32(1))] = byte(v12&i32(63) | i32(128))
															m.memory[uint32(v15)] = byte(int32(uint32(v12)>>6) | i32(192))
															goto l626
														}
														t447 := int32(load32(m.memory[int64(uint32(v7))+340:]))
														v4 = t447
														v11 = v4 + v11
														m.memory[uint32(v11+i32(2))] = byte(v12&i32(63) | i32(128))
														m.memory[uint32(v11+i32(1))] = byte(int32(uint32(v12)>>6)&i32(63) | i32(128))
														m.memory[uint32(v11)] = byte(int32(uint32(v12&i32(61440))>>12) | i32(224))
														goto l609
													}
												}
												v15 = v4 + i32(-8923)
												if uint32(v15) >= uint32(i32(101)) {
													m.fn158(v15, i32(101), i32(1155528))
													panic("unreachable")
												}
												t432 := int32(load32(m.memory[int64(uint32(v7))+340:]))
												v4 = t432
												v12 = v4 + v11
												t433 := int32(load16(m.memory[int64(uint32(v15<<1))+1155544:]))
												t434 := v12 + i32(2)
												v11 = t433
												m.memory[uint32(t434)] = byte(v11&i32(63) | i32(128))
												m.memory[uint32(v12)] = byte(int32(uint32(v11)>>12) | i32(224))
												m.memory[uint32(v12+i32(1))] = byte(int32(uint32(v11)>>6)&i32(63) | i32(128))
												goto l609
											}
										}
										{
											v17 = v15 + i32(-64)
											if uint32(v17&i32(255)) <= uint32(i32(62)) {
												goto l603
											}
											if v15 > i32(-2) {
												m.memory[int64(uint32(v0))+4] = byte(i32(2))
												if v15 > i32(-1) {
													store32(m.memory[uint32(v0):], uint32(v12))
													store16(m.memory[int64(uint32(v0))+5:], uint16(i32(1)))
													goto l514
												}
												store32(m.memory[uint32(v0):], uint32(v16))
												store16(m.memory[int64(uint32(v0))+5:], uint16(i32(2)))
												goto l514
											}
											v17 = v15 + i32(-65)
										l603:
											t425 := v7 + i32(128)
											v15 = v14*i32(190) + v17&i32(255)
											m.fn1711(t425, i32(1265310), i32(1916), v15)
											t426 := int32(load32(m.memory[int64(uint32(v7))+132:]))
											v12 = t426
											{
												{
													t427 := int32(load32(m.memory[int64(uint32(v7))+128:]))
													if t427 != i32(1) {
														if uint32(v12) > uint32(i32(1915)) {
															m.fn158(i32(1916), i32(1916), i32(1236596))
															panic("unreachable")
														}
														t428 := int32(load16(m.memory[int64(uint32(v12<<1))+1256992:]))
														v15 = t428
														v14 = int32(uint32(v15) >> 12)
														v17 = int32(uint32(v15) >> 6)
														goto l608
													}
													v12 = v12 + i32(-1)
													if uint32(v12) < uint32(i32(1916)) {
														goto l606
													}
													m.fn158(i32(-1), i32(1916), i32(1236612))
													panic("unreachable")
												}
											l606:
												t429 := v15
												v12 = v12 << 1
												t430 := int32(load16(m.memory[int64(uint32(v12))+1265310:]))
												t431 := int32(load16(m.memory[int64(uint32(v12))+1256992:]))
												v15 = t429 - t430 + t431
												v12 = v15 & i32(0xffff)
												v14 = int32(uint32(v12) >> 12)
												v17 = int32(uint32(v12) >> 6)
											}
										l608:
											v12 = v4 + v11
											m.memory[uint32(v12)] = byte(v14 | i32(224))
											m.memory[uint32(v12+i32(2))] = byte(v15&i32(63) | i32(128))
											m.memory[uint32(v12+i32(1))] = byte(v17&i32(63) | i32(128))
											goto l609
										}
									}
									m.memory[int64(uint32(v1))+11] = byte(v14)
									v15 = i32(2)
									m.memory[int64(uint32(v1))+9] = byte(i32(2))
									v10 = v14
									goto l601
								case 1:
									m.memory[int64(uint32(v1))+9] = byte(i32(0))
									v9 = v15 + i32(127)
									if uint32(v9&i32(255)) > uint32(i32(125)) {
										m.memory[int64(uint32(v1))+7] = byte(i32(1))
										m.memory[int64(uint32(v0))+6] = byte(i32(1))
										store16(m.memory[int64(uint32(v0))+4:], uint16(i32(258)))
										store32(m.memory[uint32(v0):], uint32(v12))
										m.memory[int64(uint32(v1))+8] = byte(v10 + i32(48))
										goto l514
									}
									m.memory[int64(uint32(v1))+12] = byte(v9)
									v15 = i32(3)
									m.memory[int64(uint32(v1))+9] = byte(i32(3))
									goto l601
								case 2:
									m.memory[int64(uint32(v1))+9] = byte(i32(0))
									v4 = (v15 + i32(-48)) & i32(255)
									if uint32(v4) > uint32(i32(9)) {
										m.memory[int64(uint32(v1))+10] = byte(v9)
										m.memory[int64(uint32(v1))+9] = byte(i32(1))
										m.memory[int64(uint32(v1))+7] = byte(i32(1))
										m.memory[int64(uint32(v0))+6] = byte(i32(2))
										store16(m.memory[int64(uint32(v0))+4:], uint16(i32(258)))
										store32(m.memory[uint32(v0):], uint32(v12))
										m.memory[int64(uint32(v1))+8] = byte(v10 + i32(48))
										goto l514
									}
									{
										t465 := int32(m.memory[int64(uint32(v1))+10])
										v12 = v10&i32(255)*i32(1260) + v9&i32(255)*i32(10) + v4 + t465*i32(12600)
										if uint32(v12) < uint32(i32(39420)) {
											if v12 == i32(7457) {
												t475 := int32(load32(m.memory[int64(uint32(v7))+340:]))
												v4 = t475
												v12 = v4 + v11
												store16(m.memory[uint32(v12):], uint16(i32(40942)))
												m.memory[uint32(v12+i32(2))] = byte(i32(135))
												goto l609
											}
											m.fn1711(v7+i32(96), i32(1250872), i32(206), v12)
											t468 := int32(load32(m.memory[int64(uint32(v7))+100:]))
											v4 = t468
											{
												{
													t469 := int32(load32(m.memory[int64(uint32(v7))+96:]))
													if t469 != i32(1) {
														if uint32(v4) > uint32(i32(205)) {
															m.fn158(v4, i32(206), i32(1236596))
															panic("unreachable")
														}
														t470 := int32(load16(m.memory[int64(uint32(v4<<1))+1250224:]))
														v12 = t470
														goto l650
													}
													v4 = v4 + i32(-1)
													if uint32(v4) < uint32(i32(206)) {
														goto l648
													}
													m.fn158(v4, i32(206), i32(1236612))
													panic("unreachable")
												}
											l648:
												v4 = v4 << 1
												t471 := int32(load16(m.memory[int64(uint32(v4))+1250224:]))
												t472 := int32(load16(m.memory[int64(uint32(v4))+1250872:]))
												v12 = t471 + v12 - t472
											}
										l650:
											{
												if uint32(v12&i32(0xffff)) < uint32(i32(2048)) {
													t474 := int32(load32(m.memory[int64(uint32(v7))+340:]))
													v4 = t474
													v15 = v4 + v11
													m.memory[uint32(v15+i32(1))] = byte(v12&i32(63) | i32(128))
													m.memory[uint32(v15)] = byte(int32(uint32(v12)>>6) | i32(192))
													goto l626
												}
												t473 := int32(load32(m.memory[int64(uint32(v7))+340:]))
												v4 = t473
												v11 = v4 + v11
												m.memory[uint32(v11+i32(2))] = byte(v12&i32(63) | i32(128))
												m.memory[uint32(v11+i32(1))] = byte(int32(uint32(v12)>>6)&i32(63) | i32(128))
												m.memory[uint32(v11)] = byte(int32(uint32(v12&i32(61440))>>12) | i32(224))
												goto l609
											}
										}
										if uint32(v12+i32(-189000)) < uint32(i32(0x100000)) {
											t466 := int32(load32(m.memory[int64(uint32(v7))+340:]))
											v4 = t466
											v15 = v4 + v11
											t467 := v15
											v12 = v12 + i32(-123464)
											m.memory[uint32(t467)] = byte(int32(uint32(v12)>>18) | i32(240))
											m.memory[uint32(v15+i32(2))] = byte(int32(uint32(v12)>>6)&i32(63) | i32(128))
											m.memory[uint32(v15+i32(1))] = byte(int32(uint32(v12)>>12)&i32(63) | i32(128))
											m.memory[uint32(v4+v18)] = byte(v12&i32(63) | i32(128))
											v11 = v11 + i32(4)
											goto l645
										}
										m.memory[int64(uint32(v0))+6] = byte(i32(0))
										store16(m.memory[int64(uint32(v0))+4:], uint16(i32(1026)))
										store32(m.memory[uint32(v0):], uint32(v16))
										goto l514
									}
								}
							}
							store32(m.memory[uint32(v0):], uint32(v12))
							m.memory[int64(uint32(v0))+4] = byte(i32(1))
							goto l514
						}
						if v6 != 0 {
							m.memory[int64(uint32(v0))+6] = byte(i32(0))
							m.memory[int64(uint32(v0))+5] = byte(v15)
							m.memory[int64(uint32(v0))+4] = byte(i32(2))
							m.memory[int64(uint32(v1))+9] = byte(i32(0))
							store32(m.memory[uint32(v0):], uint32(v3))
							goto l514
						}
						store32(m.memory[uint32(v0):], uint32(v3))
						m.memory[int64(uint32(v0))+4] = byte(i32(0))
						goto l514
					l609:
						store32(m.memory[int64(uint32(v7))+348:], uint32(v18))
						v15 = i32(0)
						v11 = v18
						goto l601
					l626:
						v11 = v11 + i32(2)
					l645:
						store32(m.memory[int64(uint32(v7))+348:], uint32(v11))
						v15 = i32(0)
						goto l601
					}
					t356 := int32(load32(m.memory[int64(uint32(v7))+340:]))
					v18 = t356
					t357 := int32(load32(m.memory[int64(uint32(v7))+344:]))
					v9 = t357
				l593:
					{
						m.fn148(v7+i32(88), v12, v2, v3, i32(1155496))
						t358 := int32(load32(m.memory[int64(uint32(v7))+88:]))
						v15 = t358
						t359 := int32(load32(m.memory[int64(uint32(v7))+92:]))
						v5 = t359
						m.fn212(v7+i32(80), v11, v18, v9, i32(1155512))
						t360 := int32(load32(m.memory[int64(uint32(v7))+84:]))
						v4 = t360
						t361 := v4
						t362 := v5
						var p363 int32
						if uint32(v4) < uint32(v5) {
							p363 = 1
						}
						v13 = p363
						p364 := t362
						if v13 != 0 {
							p364 = t361
						}
						v10 = p364
						v5 = i32(0)
						{
							{
								{
									t365 := int32(load32(m.memory[int64(uint32(v7))+80:]))
									t366 := v15
									v16 = t365
									if (t366^v16)&i32(3) != 0 {
										goto l516
									}
									v5 = i32(0)
									v14 = (i32(0) - v15) & i32(3)
									if uint32(v14|i32(8)) > uint32(v10) {
										goto l516
									}
									v5 = i32(0)
								l523:
									if v14 != v5 {
										t373 := int32(int8(m.memory[uint32(v15+v5)]))
										v4 = t373
										if v4 < i32(0) {
											goto l521
										}
										m.memory[uint32(v16+v5)] = byte(v4)
										v5 = v5 + i32(1)
										goto l523
									}
									v8 = v10 + i32(-8)
								l522:
									{
										v14 = v15 + v5
										t367 := int32(load32(m.memory[uint32(v14):]))
										v4 = t367
										v17 = v16 + v5
										t368 := int32(load32(m.memory[uint32(v14+i32(4)):]))
										t369 := v17 + i32(4)
										v14 = t368
										store32(m.memory[uint32(t369):], uint32(v14))
										store32(m.memory[uint32(v17):], uint32(v4))
										{
											v14 = v14 & i32(-2139062144)
											t370 := v14
											v4 = v4 & i32(-2139062144)
											if t370|v4 == 0 {
												v5 = v5 + i32(8)
												if uint32(v5) <= uint32(v8) {
													goto l522
												}
												goto l516
											}
											if v4 != 0 {
												goto l519
											}
											v4 = int32(uint32(int32(bits.TrailingZeros32(uint32(v14))))>>3) + i32(4)
											goto l520
										l519:
											v4 = int32(uint32(int32(bits.TrailingZeros32(uint32(v4)))) >> 3)
										l520:
											t371 := v15
											v5 = v4 + v5
											t372 := int32(m.memory[uint32(t371+v5)])
											v4 = t372
											goto l521
										}
									}
								}
							l516:
								p374 := v10
								if uint32(v5) > uint32(v10) {
									p374 = v5
								}
								v14 = p374
							l525:
								{
									if v14 == v5 {
										v11 = v10 + v11
										v5 = v10 + v12
										goto l526
									}
									t375 := int32(int8(m.memory[uint32(v15+v5)]))
									v4 = t375
									if v4 < i32(0) {
										goto l521
									}
									m.memory[uint32(v16+v5)] = byte(v4)
									v5 = v5 + i32(1)
									goto l525
								}
							}
						l521:
							t376 := v7
							v11 = v5 + v11
							store32(m.memory[int64(uint32(t376))+348:], uint32(v11))
							v5 = v5 + v12
							if uint32(v11+i32(3)) < uint32(v9) {
								v12 = v5 + i32(1)
							l592:
								{
									v10 = v4 + i32(127)
									v15 = v10 & i32(255)
									if uint32(v15) > uint32(i32(125)) {
										goto l528
									}
									if uint32(v12) < uint32(v3) {
										v5 = v12 + i32(1)
										{
											t377 := int32(int8(m.memory[uint32(v2+v12)]))
											v16 = t377
											v9 = v16 + i32(-48)
											v14 = v9 & i32(255)
											if uint32(v14) > uint32(i32(9)) {
												if uint32(v15) > uint32(i32(31)) {
													v14 = (v16 + i32(95)) & i32(255)
													if uint32(v14) < uint32(i32(94)) {
														{
															v12 = (v4 + i32(80)) & i32(255)
															if uint32(v12) < uint32(i32(72)) {
																v4 = v18 + v11
																t407 := int32(load16(m.memory[int64(uint32(v12*i32(188)+v14<<1))+1155882:]))
																t408 := v4 + i32(2)
																v12 = t407
																m.memory[uint32(t408)] = byte(v12&i32(63) | i32(128))
																m.memory[uint32(v4)] = byte(int32(uint32(v12)>>12) | i32(224))
																m.memory[uint32(v4+i32(1))] = byte(int32(uint32(v12)>>6)&i32(63) | i32(128))
																v4 = i32(3)
																goto l541
															}
															v12 = v4 & i32(255)
															switch v12 + i32(-166) {
															case 0:
																v12 = (v16 + i32(32)) & i32(255)
																if uint32(v12) > uint32(i32(21)) {
																	goto l578
																}
																v4 = v18 + v11
																t409 := int32(load16(m.memory[int64(uint32(v12<<1))+1261894:]))
																t410 := v4 + i32(2)
																v12 = t409
																m.memory[uint32(t410)] = byte(v12&i32(63) | i32(128))
																m.memory[uint32(v4)] = byte(int32(uint32(v12)>>12) | i32(224))
																m.memory[uint32(v4+i32(1))] = byte(int32(uint32(v12)>>6)&i32(63) | i32(128))
																v4 = i32(3)
																goto l541
															case 1:
																goto l575
															case 2:
																goto l576
															default:
																goto l577
															}
														l577:
															if v12 == i32(161) {
																goto l579
															}
														l575:
															if uint32(v15) <= uint32(i32(118)) {
																goto l578
															}
															v12 = v18 + v11
															t411 := v12 + i32(2)
															v4 = (v4+i32(8))&i32(255)*i32(94) + v14 + i32(-7628)
															m.memory[uint32(t411)] = byte(v4&i32(63) | i32(128))
															m.memory[uint32(v12+i32(1))] = byte(int32(uint32(v4)>>6) & i32(191))
															m.memory[uint32(v12)] = byte(int32(uint32(v4&i32(61440))>>12) | i32(224))
															v4 = i32(3)
															goto l541
														}
													l576:
														if uint32(v14) < uint32(i32(32)) {
															t418 := int32(load16(m.memory[int64(uint32(v14<<1))+1207098:]))
															v4 = t418
															if v14 != i32(27) {
																v12 = v18 + v11
																m.memory[uint32(v12+i32(1))] = byte(v4&i32(63) | i32(128))
																m.memory[uint32(v12)] = byte(int32(uint32(v4)>>6) | i32(192))
																goto l548
															}
															v12 = v18 + v11
															m.memory[uint32(v12+i32(2))] = byte(v4&i32(63) | i32(128))
															m.memory[uint32(v12+i32(1))] = byte(int32(uint32(v4)>>6)&i32(63) | i32(128))
															m.memory[uint32(v12)] = byte(int32(uint32(v4&i32(61440))>>12) | i32(224))
															v4 = i32(3)
															goto l541
														}
													l578:
														t412 := v7 + i32(40)
														v12 = (v4+i32(94))&i32(255)*i32(94) + v14
														m.fn1711(t412, i32(1250636), i32(46), v12)
														t413 := int32(load32(m.memory[int64(uint32(v7))+44:]))
														v4 = t413
														{
															{
																t414 := int32(load32(m.memory[int64(uint32(v7))+40:]))
																if t414 != i32(1) {
																	if uint32(v4) > uint32(i32(45)) {
																		m.fn158(v4, i32(46), i32(1236596))
																		panic("unreachable")
																	}
																	t415 := int32(load16(m.memory[int64(uint32(v4<<1))+1272396:]))
																	v4 = t415
																	goto l584
																}
																v4 = v4 + i32(-1)
																if uint32(v4) < uint32(i32(46)) {
																	goto l582
																}
																m.fn158(v4, i32(46), i32(1236612))
																panic("unreachable")
															}
														l582:
															v4 = v4 << 1
															t416 := int32(load16(m.memory[int64(uint32(v4))+1272396:]))
															t417 := int32(load16(m.memory[int64(uint32(v4))+1250636:]))
															v4 = t416 + v12 - t417
														}
													l584:
														if uint32(v4&i32(0xffff)) < uint32(i32(2048)) {
															v12 = v18 + v11
															m.memory[uint32(v12+i32(1))] = byte(v4&i32(63) | i32(128))
															m.memory[uint32(v12)] = byte(int32(uint32(v4)>>6) | i32(192))
															goto l548
														}
														v12 = v18 + v11
														m.memory[uint32(v12+i32(2))] = byte(v4&i32(63) | i32(128))
														m.memory[uint32(v12+i32(1))] = byte(int32(uint32(v4)>>6)&i32(63) | i32(128))
														m.memory[uint32(v12)] = byte(int32(uint32(v4&i32(61440))>>12) | i32(224))
														v4 = i32(3)
														goto l541
													}
													v15 = v16 + i32(-64)
													if uint32(v15&i32(255)) <= uint32(i32(62)) {
														goto l558
													}
													if v16 > i32(-96) {
														m.memory[int64(uint32(v0))+4] = byte(i32(2))
														if v16 > i32(-1) {
															store32(m.memory[uint32(v0):], uint32(v12))
															store16(m.memory[int64(uint32(v0))+5:], uint16(i32(1)))
															goto l514
														}
														store32(m.memory[uint32(v0):], uint32(v5))
														store16(m.memory[int64(uint32(v0))+5:], uint16(i32(2)))
														goto l514
													}
													v15 = v16 + i32(-65)
												l558:
													{
														v4 = (v4+i32(95))&i32(255)*i32(96) + v15&i32(255)
														v12 = v4 + i32(-864)
														if uint32(v12) < uint32(i32(8059)) {
															m.fn1711(v7+i32(48), i32(1269142), i32(1627), v12)
															t396 := int32(load32(m.memory[int64(uint32(v7))+52:]))
															v4 = t396
															{
																{
																	t397 := int32(load32(m.memory[int64(uint32(v7))+48:]))
																	if t397 != i32(1) {
																		if uint32(v4) > uint32(i32(1626)) {
																			m.fn158(v4, i32(1627), i32(1236596))
																			panic("unreachable")
																		}
																		t398 := int32(load16(m.memory[int64(uint32(v4<<1))+1261938:]))
																		v12 = t398
																		v15 = int32(uint32(v12) >> 12)
																		v16 = int32(uint32(v12) >> 6)
																		goto l566
																	}
																	v4 = v4 + i32(-1)
																	if uint32(v4) < uint32(i32(1627)) {
																		goto l564
																	}
																	m.fn158(v4, i32(1627), i32(1236612))
																	panic("unreachable")
																}
															l564:
																t399 := v12
																v4 = v4 << 1
																t400 := int32(load16(m.memory[int64(uint32(v4))+1269142:]))
																t401 := int32(load16(m.memory[int64(uint32(v4))+1261938:]))
																v12 = t399 - t400 + t401
																v4 = v12 & i32(0xffff)
																v15 = int32(uint32(v4) >> 12)
																v16 = int32(uint32(v4) >> 6)
															}
														l566:
															v4 = v18 + v11
															m.memory[uint32(v4)] = byte(v15 | i32(224))
															m.memory[uint32(v4+i32(2))] = byte(v12&i32(63) | i32(128))
															m.memory[uint32(v4+i32(1))] = byte(v16&i32(63) | i32(128))
															v4 = i32(3)
															goto l541
														}
														if uint32(v4) < uint32(i32(864)) {
															m.fn1711(v7+i32(56), i32(1244158), i32(59), v4)
															t402 := int32(load32(m.memory[int64(uint32(v7))+60:]))
															v12 = t402
															{
																{
																	t403 := int32(load32(m.memory[int64(uint32(v7))+56:]))
																	if t403 != i32(1) {
																		if uint32(v12) > uint32(i32(58)) {
																			m.fn158(v12, i32(59), i32(1236596))
																			panic("unreachable")
																		}
																		t404 := int32(load16(m.memory[int64(uint32(v12<<1))+1265192:]))
																		v4 = t404
																		goto l570
																	}
																	v12 = v12 + i32(-1)
																	if uint32(v12) < uint32(i32(59)) {
																		goto l568
																	}
																	m.fn158(v12, i32(59), i32(1236612))
																	panic("unreachable")
																}
															l568:
																v12 = v12 << 1
																t405 := int32(load16(m.memory[int64(uint32(v12))+1265192:]))
																t406 := int32(load16(m.memory[int64(uint32(v12))+1244158:]))
																v4 = t405 + v4 - t406
															}
														l570:
															if uint32(v4&i32(0xffff)) < uint32(i32(2048)) {
																v12 = v18 + v11
																m.memory[uint32(v12+i32(1))] = byte(v4&i32(63) | i32(128))
																m.memory[uint32(v12)] = byte(int32(uint32(v4)>>6) | i32(192))
																goto l548
															}
															v12 = v18 + v11
															m.memory[uint32(v12+i32(2))] = byte(v4&i32(63) | i32(128))
															m.memory[uint32(v12+i32(1))] = byte(int32(uint32(v4)>>6)&i32(63) | i32(128))
															m.memory[uint32(v12)] = byte(int32(uint32(v4&i32(61440))>>12) | i32(224))
															v4 = i32(3)
															goto l541
														}
														v12 = v4 + i32(-8923)
														if uint32(v12) >= uint32(i32(101)) {
															m.fn158(v12, i32(101), i32(1155528))
															panic("unreachable")
														}
														v4 = v18 + v11
														t394 := int32(load16(m.memory[int64(uint32(v12<<1))+1155544:]))
														t395 := v4 + i32(2)
														v12 = t394
														m.memory[uint32(t395)] = byte(v12&i32(63) | i32(128))
														m.memory[uint32(v4)] = byte(int32(uint32(v12)>>12) | i32(224))
														m.memory[uint32(v4+i32(1))] = byte(int32(uint32(v12)>>6)&i32(63) | i32(128))
														v4 = i32(3)
														goto l541
													}
												}
												{
													v4 = v16 + i32(-64)
													if uint32(v4&i32(255)) <= uint32(i32(62)) {
														goto l550
													}
													if v16 > i32(-2) {
														m.memory[int64(uint32(v0))+4] = byte(i32(2))
														if v16 > i32(-1) {
															store32(m.memory[uint32(v0):], uint32(v12))
															store16(m.memory[int64(uint32(v0))+5:], uint16(i32(1)))
															goto l514
														}
														store32(m.memory[uint32(v0):], uint32(v5))
														store16(m.memory[int64(uint32(v0))+5:], uint16(i32(2)))
														goto l514
													}
													v4 = v16 + i32(-65)
												l550:
													t387 := v7 + i32(64)
													v12 = v15*i32(190) + v4&i32(255)
													m.fn1711(t387, i32(1265310), i32(1916), v12)
													t388 := int32(load32(m.memory[int64(uint32(v7))+68:]))
													v4 = t388
													{
														{
															t389 := int32(load32(m.memory[int64(uint32(v7))+64:]))
															if t389 != i32(1) {
																if uint32(v4) > uint32(i32(1915)) {
																	m.fn158(i32(1916), i32(1916), i32(1236596))
																	panic("unreachable")
																}
																t390 := int32(load16(m.memory[int64(uint32(v4<<1))+1256992:]))
																v12 = t390
																v15 = int32(uint32(v12) >> 12)
																v16 = int32(uint32(v12) >> 6)
																goto l555
															}
															v4 = v4 + i32(-1)
															if uint32(v4) < uint32(i32(1916)) {
																goto l553
															}
															m.fn158(i32(-1), i32(1916), i32(1236612))
															panic("unreachable")
														}
													l553:
														t391 := v12
														v4 = v4 << 1
														t392 := int32(load16(m.memory[int64(uint32(v4))+1265310:]))
														t393 := int32(load16(m.memory[int64(uint32(v4))+1256992:]))
														v12 = t391 - t392 + t393
														v4 = v12 & i32(0xffff)
														v15 = int32(uint32(v4) >> 12)
														v16 = int32(uint32(v4) >> 6)
													}
												l555:
													v4 = v18 + v11
													m.memory[uint32(v4)] = byte(v15 | i32(224))
													m.memory[uint32(v4+i32(2))] = byte(v12&i32(63) | i32(128))
													m.memory[uint32(v4+i32(1))] = byte(v16&i32(63) | i32(128))
													v4 = i32(3)
													goto l541
												}
											}
											if uint32(v5) < uint32(v3) {
												t378 := int32(m.memory[uint32(v2+v5)])
												v17 = t378 + i32(127)
												v4 = v17 & i32(255)
												if uint32(v4) > uint32(i32(125)) {
													m.memory[int64(uint32(v1))+7] = byte(i32(1))
													m.memory[int64(uint32(v0))+6] = byte(i32(1))
													store16(m.memory[int64(uint32(v0))+4:], uint16(i32(258)))
													store32(m.memory[uint32(v0):], uint32(v5))
													m.memory[int64(uint32(v1))+8] = byte(v9 | i32(48))
													goto l514
												}
												v5 = v12 + i32(2)
												if uint32(v5) < uint32(v3) {
													t379 := int32(m.memory[uint32(v2+v5)])
													v16 = (t379 + i32(-48)) & i32(255)
													if uint32(v16) > uint32(i32(9)) {
														m.memory[int64(uint32(v1))+10] = byte(v17)
														m.memory[int64(uint32(v1))+9] = byte(i32(1))
														m.memory[int64(uint32(v1))+7] = byte(i32(1))
														m.memory[int64(uint32(v0))+6] = byte(i32(2))
														store16(m.memory[int64(uint32(v0))+4:], uint16(i32(258)))
														store32(m.memory[uint32(v0):], uint32(v5))
														m.memory[int64(uint32(v1))+8] = byte(v9 | i32(48))
														goto l514
													}
													v5 = v12 + i32(3)
													v4 = v14*i32(1260) + v15*i32(12600) + v4*i32(10) + v16
													if uint32(v4) < uint32(i32(39420)) {
														if v4 == i32(7457) {
															goto l542
														}
														m.fn1711(v7+i32(72), i32(1250872), i32(206), v4)
														t382 := int32(load32(m.memory[int64(uint32(v7))+76:]))
														v12 = t382
														{
															{
																t383 := int32(load32(m.memory[int64(uint32(v7))+72:]))
																if t383 != i32(1) {
																	if uint32(v12) > uint32(i32(205)) {
																		m.fn158(v12, i32(206), i32(1236596))
																		panic("unreachable")
																	}
																	t384 := int32(load16(m.memory[int64(uint32(v12<<1))+1250224:]))
																	v4 = t384
																	goto l546
																}
																v12 = v12 + i32(-1)
																if uint32(v12) < uint32(i32(206)) {
																	goto l544
																}
																m.fn158(v12, i32(206), i32(1236612))
																panic("unreachable")
															}
														l544:
															v12 = v12 << 1
															t385 := int32(load16(m.memory[int64(uint32(v12))+1250224:]))
															t386 := int32(load16(m.memory[int64(uint32(v12))+1250872:]))
															v4 = t385 + v4 - t386
														}
													l546:
														if uint32(v4&i32(0xffff)) < uint32(i32(2048)) {
															v12 = v18 + v11
															m.memory[uint32(v12+i32(1))] = byte(v4&i32(63) | i32(128))
															m.memory[uint32(v12)] = byte(int32(uint32(v4)>>6) | i32(192))
															goto l548
														}
														v12 = v18 + v11
														m.memory[uint32(v12+i32(2))] = byte(v4&i32(63) | i32(128))
														m.memory[uint32(v12+i32(1))] = byte(int32(uint32(v4)>>6)&i32(63) | i32(128))
														m.memory[uint32(v12)] = byte(int32(uint32(v4&i32(61440))>>12) | i32(224))
														v4 = i32(3)
														goto l541
													}
													if uint32(v4+i32(-189000)) < uint32(i32(0x100000)) {
														t380 := v18 + v11
														v4 = v4 + i32(-123464)
														m.memory[uint32(t380)] = byte(int32(uint32(v4)>>18) | i32(240))
														t381 := int32(load32(m.memory[int64(uint32(v7))+340:]))
														v18 = t381
														v12 = v18 + v11
														m.memory[uint32(v12+i32(3))] = byte(v4&i32(63) | i32(128))
														m.memory[uint32(v12+i32(2))] = byte(int32(uint32(v4)>>6)&i32(63) | i32(128))
														m.memory[uint32(v12+i32(1))] = byte(int32(uint32(v4)>>12)&i32(63) | i32(128))
														v4 = i32(4)
														goto l541
													}
													m.memory[int64(uint32(v0))+6] = byte(i32(0))
													store16(m.memory[int64(uint32(v0))+4:], uint16(i32(1026)))
													store32(m.memory[uint32(v0):], uint32(v5))
													goto l514
												}
												if v6 != 0 {
													m.memory[int64(uint32(v0))+6] = byte(i32(0))
													store16(m.memory[int64(uint32(v0))+4:], uint16(i32(770)))
													store32(m.memory[uint32(v0):], uint32(v5))
													goto l514
												}
												m.memory[int64(uint32(v1))+12] = byte(v17)
												m.memory[int64(uint32(v1))+11] = byte(v9)
												m.memory[int64(uint32(v1))+10] = byte(v10)
												m.memory[int64(uint32(v1))+9] = byte(i32(3))
												m.memory[int64(uint32(v0))+4] = byte(i32(0))
												store32(m.memory[uint32(v0):], uint32(v5))
												goto l514
											}
											if v6 != 0 {
												m.memory[int64(uint32(v0))+6] = byte(i32(0))
												store16(m.memory[int64(uint32(v0))+4:], uint16(i32(514)))
												store32(m.memory[uint32(v0):], uint32(v5))
												goto l514
											}
											m.memory[int64(uint32(v1))+11] = byte(v9)
											m.memory[int64(uint32(v1))+10] = byte(v10)
											m.memory[int64(uint32(v1))+9] = byte(i32(2))
											m.memory[int64(uint32(v0))+4] = byte(i32(0))
											store32(m.memory[uint32(v0):], uint32(v5))
											goto l514
										}
									}
									if v6 != 0 {
										m.memory[int64(uint32(v0))+6] = byte(i32(0))
										store16(m.memory[int64(uint32(v0))+4:], uint16(i32(258)))
										goto l531
									}
									m.memory[int64(uint32(v1))+10] = byte(v10)
									m.memory[int64(uint32(v1))+9] = byte(i32(1))
									m.memory[int64(uint32(v0))+4] = byte(i32(0))
									goto l531
								l528:
									if v4&i32(255) == i32(128) {
										v5 = v18 + v11
										store16(m.memory[uint32(v5):], uint16(i32(33506)))
										m.memory[uint32(v5+i32(2))] = byte(i32(172))
										v11 = v11 + i32(3)
										goto l587
									}
									m.memory[int64(uint32(v0))+6] = byte(i32(0))
									store16(m.memory[int64(uint32(v0))+4:], uint16(i32(258)))
								l531:
									store32(m.memory[uint32(v0):], uint32(v12))
									goto l514
								l579:
									{
										t419 := int32(load16(m.memory[int64(uint32(v14<<1))+1228540:]))
										v4 = t419
										if uint32(v4) < uint32(i32(2048)) {
											goto l589
										}
										v12 = v18 + v11
										m.memory[uint32(v12+i32(2))] = byte(v4&i32(63) | i32(128))
										m.memory[uint32(v12)] = byte(int32(uint32(v4)>>12) | i32(224))
										m.memory[uint32(v12+i32(1))] = byte(int32(uint32(v4)>>6)&i32(63) | i32(128))
										v4 = i32(3)
										goto l541
									}
								l589:
									v12 = v18 + v11
									m.memory[uint32(v12+i32(1))] = byte(v4&i32(63) | i32(128))
									m.memory[uint32(v12)] = byte(int32(uint32(v4)>>6) | i32(192))
								l548:
									v4 = i32(2)
									goto l541
								l542:
									v4 = v18 + v11
									store16(m.memory[uint32(v4):], uint16(i32(40942)))
									m.memory[uint32(v4+i32(2))] = byte(i32(135))
									v4 = i32(3)
								l541:
									v11 = v4 + v11
									if uint32(v5) < uint32(v3) {
										goto l590
									}
									store32(m.memory[uint32(v0):], uint32(v5))
									m.memory[int64(uint32(v0))+4] = byte(i32(0))
									goto l514
								l590:
									{
										t420 := int32(load32(m.memory[int64(uint32(v7))+344:]))
										t421 := v11 + i32(3)
										v9 = t420
										if uint32(t421) < uint32(v9) {
											goto l591
										}
										store32(m.memory[uint32(v0):], uint32(v5))
										m.memory[int64(uint32(v0))+4] = byte(i32(1))
										goto l514
									}
								l591:
									v12 = v5 + i32(1)
									t422 := int32(int8(m.memory[uint32(v2+v5)]))
									v4 = t422
									if v4 < i32(0) {
										goto l592
									}
								}
								m.memory[uint32(v18+v11)] = byte(v4)
								v11 = v11 + i32(1)
							l587:
								store32(m.memory[int64(uint32(v7))+348:], uint32(v11))
								goto l593
							}
							v13 = i32(1)
						}
					l526:
						store32(m.memory[uint32(v0):], uint32(v5))
						store32(m.memory[int64(uint32(v0))+4:], uint32(v13|(v7+i32(340))&i32(-256)))
						goto l514
					}
				}
			}
		l514:
			store32(m.memory[int64(uint32(v0))+8:], uint32(v11))
			goto l35
		case 1:
			t482 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			v17 = t482
			t483 := int32(m.memory[int64(uint32(v1))+17])
			v19 = t483
			t484 := int32(m.memory[int64(uint32(v1))+16])
			v13 = t484
			t485 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			v11 = t485
			t486 := int32(load32(m.memory[int64(uint32(v1))+12:]))
			v15 = t486
			v12 = i32(0)
			v14 = i32(0)
		l677:
			v8 = v11
		l671:
			{
				{
					if v15 == 0 {
						goto l654
					}
					v11 = v14
					goto l655
				l654:
					m.fn148(v7+i32(32), v14, v2, v3, i32(1155384))
					t487 := int32(load32(m.memory[int64(uint32(v7))+32:]))
					v9 = t487
					t488 := int32(load32(m.memory[int64(uint32(v7))+36:]))
					v16 = t488
					m.fn212(v7+i32(24), v12, v4, v5, i32(1155400))
					t489 := int32(load32(m.memory[int64(uint32(v7))+24:]))
					v10 = t489
					t490 := int32(load32(m.memory[int64(uint32(v7))+28:]))
					t491 := v9
					v18 = t490
					p492 := v16
					if uint32(v18) < uint32(v16) {
						p492 = v18
					}
					t493 := m.fn1692(t491, p492)
					v11 = t493
					if uint32(v11) > uint32(v18) {
						m.fn151(i32(0), v11, v18, i32(1155416))
						panic("unreachable")
					}
					if uint32(v11) > uint32(v16) {
						m.fn151(i32(0), v11, v16, i32(1155432))
						panic("unreachable")
					}
					m.fn310(v10, v11, v9, v11, i32(1155448))
					v12 = v11 + v12
					v11 = v11 + v14
				}
			l655:
				if uint32(v11) < uint32(v3) {
					v18 = v12 + i32(3)
					if uint32(v18) < uint32(v5) {
						v14 = v11 + i32(1)
						t494 := int32(m.memory[uint32(v2+v11)])
						v16 = t494
						{
							if v15 != 0 {
								if uint32(v16) < uint32(v13&i32(255)) {
									goto l669
								}
								if uint32(v16) <= uint32(v19&i32(255)) {
									store16(m.memory[int64(uint32(v1))+16:], uint16(i32(49024)))
									t497 := v1
									v17 = v17 + i32(1)
									store32(m.memory[int64(uint32(t497))+8:], uint32(v17))
									t498 := v1
									v9 = v8 << 6
									t499 := v9
									v16 = v16 & i32(63)
									v11 = t499 | v16
									store32(m.memory[int64(uint32(t498))+4:], uint32(v11))
									v13 = i32(128)
									v19 = i32(191)
									if v17 != v15 {
										goto l677
									}
									if v15 != i32(3) {
										goto l678
									}
									v11 = v4 + v12
									m.memory[uint32(v11+i32(3))] = byte(v16 | i32(128))
									m.memory[uint32(v11+i32(2))] = byte(v8&i32(63) | i32(128))
									m.memory[uint32(v11)] = byte(int32(uint32(v9)>>18) | i32(240))
									m.memory[uint32(v11+i32(1))] = byte(int32(uint32(v9)>>12)&i32(63) | i32(128))
									v12 = v12 + i32(4)
									goto l679
								l678:
									if uint32(v11&i32(0xffff)) < uint32(i32(2048)) {
										goto l680
									}
									v12 = v4 + v12
									m.memory[uint32(v12+i32(2))] = byte(v16 | i32(128))
									m.memory[uint32(v12+i32(1))] = byte(v8&i32(63) | i32(128))
									m.memory[uint32(v12)] = byte(int32(uint32(v11&i32(61440))>>12) | i32(224))
									v12 = v18
									goto l679
								l680:
									v11 = v4 + v12
									m.memory[uint32(v11)] = byte(v8 | i32(192))
									m.memory[uint32(v11+i32(1))] = byte(v16 | i32(128))
									v12 = v12 + i32(2)
								l679:
									v17 = i32(0)
									store32(m.memory[int64(uint32(v1))+12:], uint32(i32(0)))
									store64(m.memory[int64(uint32(v1))+4:], uint64(i64(0)))
									v11 = i32(0)
									v15 = i32(0)
									goto l677
								}
							l669:
								store32(m.memory[int64(uint32(v1))+12:], uint32(i32(0)))
								store64(m.memory[int64(uint32(v1))+4:], uint64(i64(0)))
								store16(m.memory[int64(uint32(v1))+16:], uint16(i32(49024)))
								m.memory[int64(uint32(v0))+6] = byte(i32(0))
								m.memory[int64(uint32(v0))+4] = byte(i32(2))
								m.memory[int64(uint32(v0))+5] = byte(v17 + i32(1))
								goto l661
							}
							if int32(int8(v16)) > i32(-1) {
								m.memory[uint32(v4+v12)] = byte(v16)
								v12 = v12 + i32(1)
								v15 = i32(0)
								goto l671
							}
							if uint32(v16) < uint32(i32(194)) {
								m.memory[int64(uint32(v0))+6] = byte(i32(0))
								store16(m.memory[int64(uint32(v0))+4:], uint16(i32(258)))
								v11 = v14
								goto l661
							}
							if uint32(v16) < uint32(i32(224)) {
								v15 = i32(1)
								store32(m.memory[int64(uint32(v1))+12:], uint32(i32(1)))
								t495 := v1
								v8 = v16 & i32(31)
								store32(m.memory[int64(uint32(t495))+4:], uint32(v8))
								goto l671
							}
							if uint32(v16) < uint32(i32(240)) {
								if v16 == i32(224) {
									goto l672
								}
								if v16 != i32(237) {
									goto l673
								}
								v19 = i32(159)
								m.memory[int64(uint32(v1))+17] = byte(i32(159))
								goto l673
							}
							if uint32(v16) < uint32(i32(245)) {
								switch v16 + i32(-240) {
								default:
									goto l675
								case 0:
									v13 = i32(144)
									m.memory[int64(uint32(v1))+16] = byte(i32(144))
									goto l675
								case 4:
									v19 = i32(143)
									m.memory[int64(uint32(v1))+17] = byte(i32(143))
									goto l675
								}
							}
							m.memory[int64(uint32(v0))+6] = byte(i32(0))
							store16(m.memory[int64(uint32(v0))+4:], uint16(i32(258)))
							v11 = v14
							goto l661
						l672:
							v13 = i32(160)
							m.memory[int64(uint32(v1))+16] = byte(i32(160))
						l673:
							v15 = i32(2)
							store32(m.memory[int64(uint32(v1))+12:], uint32(i32(2)))
							t496 := v1
							v8 = v16 & i32(15)
							store32(m.memory[int64(uint32(t496))+4:], uint32(v8))
							goto l671
						}
					}
					m.memory[int64(uint32(v0))+4] = byte(i32(1))
					goto l661
				}
				if v6 == 0 {
					goto l659
				}
				if v15 != 0 {
					goto l660
				}
			l659:
				m.memory[int64(uint32(v0))+4] = byte(i32(0))
				goto l661
			l660:
				store32(m.memory[int64(uint32(v1))+12:], uint32(i32(0)))
				store64(m.memory[int64(uint32(v1))+4:], uint64(i64(0)))
				m.memory[int64(uint32(v0))+6] = byte(i32(0))
				m.memory[int64(uint32(v0))+4] = byte(i32(2))
				m.memory[int64(uint32(v0))+5] = byte(v17 + i32(1))
			l661:
				store32(m.memory[int64(uint32(v0))+8:], uint32(v12))
				store32(m.memory[uint32(v0):], uint32(v11))
				goto l35
			l675:
				v15 = i32(3)
				store32(m.memory[int64(uint32(v1))+12:], uint32(i32(3)))
				t500 := v1
				v8 = v16 & i32(7)
				store32(m.memory[int64(uint32(t500))+4:], uint32(v8))
				goto l671
			}
		}
	}
l35:
	m.g0 = v7 + i32(352)
}
func (m *Module) fn1711(v0, v1, v2, v3 int32) {
	var v4, v5, v6, v7 int32
	v4 = i32(0)
	v5 = v3 & i32(0xffff)
	{
	l3:
		{
			if uint32(v2) > uint32(i32(1)) {
				t2 := v4
				v6 = int32(uint32(v2) >> 1)
				v7 = v6 + v4
				t3 := int32(load16(m.memory[uint32(v1+v7<<1):]))
				p4 := v7
				if uint32(t3) > uint32(v5) {
					p4 = t2
				}
				v4 = p4
				v2 = v2 - v6
				goto l3
			}
			v2 = i32(1)
			t0 := int32(load16(m.memory[uint32(v1+v4<<1):]))
			v6 = t0
			t1 := v6
			v7 = v3 & i32(0xffff)
			if t1 != v7 {
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
func (m *Module) fn1712(v0, v1 int32) {
	var v2, v3, v4, v5 int32
	v2 = i32(0)
	v3 = i32(1251284)
	v4 = i32(0)
l3:
	if uint32(v4) < uint32(i32(33)) {
		goto l0
	}
	goto l1
l0:
	{
		t0 := int32(load16(m.memory[uint32(v3):]))
		v5 = v1 - t0
		t1 := int32(load16(m.memory[uint32(v3+i32(2)):]))
		if uint32(v5) < uint32(t1) {
			goto l2
		}
		v3 = v3 + i32(6)
		v4 = v4 + i32(3)
		goto l3
	}
l2:
	v3 = v4 + i32(2)
	if uint32(v4) > uint32(i32(30)) {
		m.fn158(v3, i32(33), i32(1250840))
		panic("unreachable")
	}
	v2 = i32(1)
	{
		t2 := int32(load16(m.memory[int64(uint32(v3<<1))+1251284:]))
		v3 = v5 + t2
		if uint32(v3) >= uint32(i32(240)) {
			m.fn158(v3, i32(240), i32(1250856))
			panic("unreachable")
		}
		t3 := int32(load16(m.memory[int64(uint32(v3<<1))+1235784:]))
		v3 = t3
		goto l1
	}
l1:
	store16(m.memory[int64(uint32(v0))+2:], uint16(v3))
	store16(m.memory[uint32(v0):], uint16(v2))
}
func (m *Module) fn1713(v0, v1 int32) {
	var v2, v3, v4, v5 int32
	v2 = i32(0)
	v3 = i32(1253698)
	v4 = i32(0)
l3:
	if uint32(v4) < uint32(i32(33)) {
		goto l0
	}
	goto l1
l0:
	{
		t0 := int32(load16(m.memory[uint32(v3):]))
		v5 = v1 - t0
		t1 := int32(load16(m.memory[uint32(v3+i32(2)):]))
		if uint32(v5) < uint32(t1) {
			goto l2
		}
		v3 = v3 + i32(6)
		v4 = v4 + i32(3)
		goto l3
	}
l2:
	v3 = v4 + i32(2)
	if uint32(v4) > uint32(i32(30)) {
		m.fn158(v3, i32(33), i32(1251508))
		panic("unreachable")
	}
	{
		t2 := int32(load16(m.memory[int64(uint32(v3<<1))+1253698:]))
		v3 = v5 + t2
		if uint32(v3) >= uint32(i32(255)) {
			m.fn158(v3, i32(255), i32(1251524))
			panic("unreachable")
		}
		t3 := int32(load16(m.memory[int64(uint32(v3<<1))+1236628:]))
		v3 = t3
		var p4 int32
		if v3 != i32(0) {
			p4 = 1
		}
		v2 = p4
		goto l1
	}
l1:
	store16(m.memory[int64(uint32(v0))+2:], uint16(v3))
	store16(m.memory[uint32(v0):], uint16(v2))
}
func (m *Module) fn1714(v0, v1, v2 int32) {
	var v3, v4 int32
	var v5 int64
	v3 = i32(0)
	{
		{
			t0 := int32(m.memory[uint32(v1)])
			t1 := int32(m.memory[int64(uint32(v1))+8])
			t2 := int32(m.memory[int64(uint32(v1))+2])
			t3 := int32(m.memory[int64(uint32(v1))+4])
			t4 := int32(m.memory[int64(uint32(v1))+6])
			v4 = v2 + (t0+t1+t2+t3+t4)&i32(63)
			if uint32(v4) >= uint32(v2) {
				goto l0
			}
			goto l1
		}
	l0:
		v5 = int64(uint32(v4)) * i64(3)
		if int32(int64(uint64(v5)>>32)) != 0 {
			goto l1
		}
		v1 = int32(v5)
		var p5 int32
		if v1 != i32(-1) {
			p5 = 1
		}
		v3 = p5
		v1 = v1 + i32(1)
	}
l1:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
	store32(m.memory[uint32(v0):], uint32(v3))
}
func (m *Module) fn1715(v0, v1, v2 int32) {
	var v3, v4, v5, v6 int32
	var v7 int64
	var v8 int32
	v3 = i32(0)
	{
		t0 := int32(m.memory[int64(uint32(v1))+1])
		t1 := v2
		v4 = t0
		t2 := int32(m.memory[int64(uint32(v1))+4])
		t3 := v4 ^ i32(-1)
		v5 = t2
		p4 := i32(0)
		if v5 != 0 {
			p4 = t3
		}
		t5 := int32(m.memory[int64(uint32(v1))+2])
		t6 := t1 + p4&i32(1)
		var p7 int32
		if uint32(t5) > uint32(i32(4)) {
			p7 = 1
		}
		v6 = t6 + p7
		if uint32(v6) < uint32(v2) {
			goto l0
		}
		v3 = i32(0)
		t8 := int32(m.memory[uint32(v1)])
		t9 := v6
		var p10 int32
		if v5 != i32(0) {
			p10 = 1
		}
		v1 = t9 + (t8+p10&v4)&i32(3)
		if uint32(v1) < uint32(v6) {
			goto l0
		}
		v7 = int64(uint32(v1)) * i64(3)
		v8 = int32(v7)
		var p11 int32
		if int32(int64(uint64(v7)>>32)) == 0 {
			p11 = 1
		}
		v3 = p11
	}
l0:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v8))
	store32(m.memory[uint32(v0):], uint32(v3))
}
func (m *Module) fn1716(v0, v1, v2 int32) {
	var v3 int64
	{
		v1 = v2 + v1&i32(255)
		if uint32(v1) >= uint32(v2) {
			goto l0
		}
		v2 = i32(0)
		goto l1
	l0:
		v3 = int64(uint32(v1)) * i64(3)
		v1 = int32(v3)
		var p0 int32
		if int32(int64(uint64(v3)>>32)) == 0 {
			p0 = 1
		}
		v2 = p0
	}
l1:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
	store32(m.memory[uint32(v0):], uint32(v2))
}
func (m *Module) fn1717(v0, v1, v2, v3 int32) int32 {
	var v4, v5, v6 int32
	{
		if v3 == v1 {
			v2 = v2 + i32(-1)
			{
			l3:
				{
					v1 = v3
					if v1 != 0 {
						goto l1
					}
					v3 = i32(1)
					v0 = i32(0)
					goto l2
				l1:
					t2 := int32(m.memory[uint32(v2+v1)])
					v4 = t2
					t3 := v4
					t4 := v0
					v3 = v1 + i32(-1)
					v5 = t4 + v3
					t5 := int32(m.memory[uint32(v5)])
					v6 = t5
					if t3 == v6 {
						goto l3
					}
				}
				var p6 int32
				if uint32(v4) > uint32(v6) {
					p6 = 1
				}
				var p7 int32
				if uint32(v4) < uint32(v6) {
					p7 = 1
				}
				v3 = p6 - p7
				var p8 int32
				if v0 != v5 {
					p8 = 1
				}
				v0 = p8
			}
		l2:
			p9 := i32(0) - v0
			if v1 != 0 {
				p9 = v3
			}
			return p9
		}
		var p0 int32
		if uint32(v3) > uint32(v1) {
			p0 = 1
		}
		var p1 int32
		if uint32(v3) < uint32(v1) {
			p1 = 1
		}
		return p0 - p1
	}
}
func (m *Module) fn1718(v0, v1 int32) {
	store32(m.memory[uint32(v0):], uint32(i32(0)))
}
func (m *Module) fn1719(v0, v1, v2, v3, v4, v5 int32) {
	var v6, v7, v8 int32
	var v9 int64
	t0 := m.g0
	v6 = t0 - i32(32)
	m.g0 = v6
	v7 = i32(1)
	v8 = i32(4)
	v9 = int64(uint32(v5)) * int64(uint32(v3))
	if int32(int64(uint64(v9)>>32)) != 0 {
		goto l0
	}
	v3 = int32(v9)
	if uint32(v3) > uint32(i32(-0x80000000)-v4) {
		goto l0
	}
	if v1 != 0 {
		goto l1
	}
	v8 = i32(0)
	v5 = v6 + i32(28)
	goto l2
l1:
	store32(m.memory[int64(uint32(v6))+28:], uint32(v4))
	v8 = v1 * v5
	v5 = v6 + i32(24)
l2:
	store32(m.memory[uint32(v5):], uint32(v8))
	{
		{
			t1 := int32(load32(m.memory[int64(uint32(v6))+28:]))
			if t1 == 0 {
				goto l3
			}
			{
				t2 := int32(load32(m.memory[int64(uint32(v6))+24:]))
				v8 = t2
				if v8 != 0 {
					t5 := m.fn89(v2, v8, v4, v3)
					v8 = t5
					v5 = v3
					goto l5
				}
				m.fn1680(v6+i32(16), v4, v3, i32(0))
				t3 := int32(load32(m.memory[int64(uint32(v6))+20:]))
				v5 = t3
				t4 := int32(load32(m.memory[int64(uint32(v6))+16:]))
				v8 = t4
				goto l5
			}
		}
	l3:
		m.fn1680(v6+i32(8), v4, v3, i32(0))
		t6 := int32(load32(m.memory[int64(uint32(v6))+12:]))
		v5 = t6
		t7 := int32(load32(m.memory[int64(uint32(v6))+8:]))
		v8 = t7
	}
l5:
	if v8 != 0 {
		store32(m.memory[int64(uint32(v0))+4:], uint32(v8))
		v7 = i32(0)
		v8 = i32(8)
		v3 = v5
		goto l7
	}
	store32(m.memory[int64(uint32(v0))+4:], uint32(v4))
	v8 = i32(8)
	goto l7
l0:
	v3 = i32(0)
l7:
	store32(m.memory[uint32(v0+v8):], uint32(v3))
	store32(m.memory[uint32(v0):], uint32(v7))
	m.g0 = v6 + i32(32)
}
func (m *Module) fn1720(v0, v1 int32) {
	var v2 int32
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
			m.fn97(i32(1291936), i32(43), v2+i32(15), i32(1291980), i32(1295240))
			panic("unreachable")
		}
		m.fn10(v1, v0, i32(64))
	}
l0:
	m.g0 = v2 + i32(16)
}
func (m *Module) fn1721(v0, v1, v2 int32) int32 {
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
		if uint32(v1) >= uint32(i32(0x7fffffc1)) {
			m.fn97(i32(1291936), i32(43), v3+i32(15), i32(1291980), i32(1295256))
			panic("unreachable")
		}
		t1 := m.fn1557(i32(64), v1)
		v1 = t1
	}
l1:
	m.g0 = v3 + i32(16)
	return v1
}
func (m *Module) fn1722(v0 int32) {
	var v1 int32
	t0 := m.g0
	v1 = t0 - i32(16)
	m.g0 = v1
	store32(m.memory[int64(uint32(v1))+12:], uint32(i32(1287584)))
	store32(m.memory[int64(uint32(v1))+8:], uint32(v0))
	m.fn1632(i32(0), v1+i32(8), i32(1287776), v1+i32(12), i32(1287776), i32(0), v1, i32(1292452))
	panic("unreachable")
}
func (m *Module) fn1723(v0, v1, v2, v3, v4 int32) {
	if uint32(v3) < uint32(v1) {
		m.fn151(v1, v3, v3, v4)
		panic("unreachable")
	}
	store32(m.memory[int64(uint32(v0))+4:], uint32(v3-v1))
	store32(m.memory[uint32(v0):], uint32(v2+v1))
}
func (m *Module) fn1724(v0, v1, v2, v3, v4 int32) {
	if uint32(v1) <= uint32(v3) {
		goto l0
	}
	m.fn151(i32(0), v1, v3, v4)
	panic("unreachable")
l0:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
	store32(m.memory[uint32(v0):], uint32(v2))
}
func (m *Module) fn1725(v0, v1, v2 int32) int32 {
	var v3, v4, v5, v6, v7, v8 int32
	v3 = v0 & i32(0xffff)
	v4 = int32(uint32(v0) >> 16)
	switch v2 {
	default:
		if uint32(v2) < uint32(i32(16)) {
			t71 := m.fn1839(v3, v1, v2, v4)
			return t71
		}
		t0 := int32(uint32(v2) % uint32(i32(5552)))
		t1 := v2
		v5 = t0
		v6 = t1 - v5
		v7 = v6
		v8 = v1
	l9:
		if uint32(v7) < uint32(i32(5552)) {
			v8 = v5 & i32(8176)
			v2 = i32(0) - v8
			v7 = v1 + v6
			v0 = v7
		l8:
			{
				if v2 == 0 {
					t65 := m.fn1839(v3, v7+v8, v5&i32(15), v4)
					return t65
				}
				t34 := int32(m.memory[uint32(v0)])
				v3 = v3 + t34
				t35 := int32(m.memory[uint32(v0+i32(1))])
				t36 := v3 + v4
				v3 = v3 + t35
				t37 := int32(m.memory[uint32(v0+i32(2))])
				t38 := t36 + v3
				v3 = v3 + t37
				t39 := int32(m.memory[uint32(v0+i32(3))])
				t40 := t38 + v3
				v3 = v3 + t39
				t41 := int32(m.memory[uint32(v0+i32(4))])
				t42 := t40 + v3
				v3 = v3 + t41
				t43 := int32(m.memory[uint32(v0+i32(5))])
				t44 := t42 + v3
				v3 = v3 + t43
				t45 := int32(m.memory[uint32(v0+i32(6))])
				t46 := t44 + v3
				v3 = v3 + t45
				t47 := int32(m.memory[uint32(v0+i32(7))])
				t48 := t46 + v3
				v3 = v3 + t47
				t49 := int32(m.memory[uint32(v0+i32(8))])
				t50 := t48 + v3
				v3 = v3 + t49
				t51 := int32(m.memory[uint32(v0+i32(9))])
				t52 := t50 + v3
				v3 = v3 + t51
				t53 := int32(m.memory[uint32(v0+i32(10))])
				t54 := t52 + v3
				v3 = v3 + t53
				t55 := int32(m.memory[uint32(v0+i32(11))])
				t56 := t54 + v3
				v3 = v3 + t55
				t57 := int32(m.memory[uint32(v0+i32(12))])
				t58 := t56 + v3
				v3 = v3 + t57
				t59 := int32(m.memory[uint32(v0+i32(13))])
				t60 := t58 + v3
				v3 = v3 + t59
				t61 := int32(m.memory[uint32(v0+i32(14))])
				t62 := t60 + v3
				v3 = v3 + t61
				t63 := int32(m.memory[uint32(v0+i32(15))])
				t64 := t62 + v3
				v3 = v3 + t63
				v4 = t64 + v3
				v0 = v0 + i32(16)
				v2 = v2 + i32(16)
				goto l8
			}
		}
		v2 = i32(0)
	l6:
		{
			if v2 == i32(5552) {
				t66 := int32(uint32(v4) % uint32(i32(65521)))
				v4 = t66
				t67 := int32(uint32(v3) % uint32(i32(65521)))
				v3 = t67
				v7 = v7 + i32(-5552)
				v8 = v8 + i32(5552)
				goto l9
			}
			t2 := v3
			v0 = v8 + v2
			t3 := int32(m.memory[uint32(v0)])
			v3 = t2 + t3
			t4 := int32(m.memory[uint32(v0+i32(1))])
			t5 := v3 + v4
			v3 = v3 + t4
			t6 := int32(m.memory[uint32(v0+i32(2))])
			t7 := t5 + v3
			v3 = v3 + t6
			t8 := int32(m.memory[uint32(v0+i32(3))])
			t9 := t7 + v3
			v3 = v3 + t8
			t10 := int32(m.memory[uint32(v0+i32(4))])
			t11 := t9 + v3
			v3 = v3 + t10
			t12 := int32(m.memory[uint32(v0+i32(5))])
			t13 := t11 + v3
			v3 = v3 + t12
			t14 := int32(m.memory[uint32(v0+i32(6))])
			t15 := t13 + v3
			v3 = v3 + t14
			t16 := int32(m.memory[uint32(v0+i32(7))])
			t17 := t15 + v3
			v3 = v3 + t16
			t18 := int32(m.memory[uint32(v0+i32(8))])
			t19 := t17 + v3
			v3 = v3 + t18
			t20 := int32(m.memory[uint32(v0+i32(9))])
			t21 := t19 + v3
			v3 = v3 + t20
			t22 := int32(m.memory[uint32(v0+i32(10))])
			t23 := t21 + v3
			v3 = v3 + t22
			t24 := int32(m.memory[uint32(v0+i32(11))])
			t25 := t23 + v3
			v3 = v3 + t24
			t26 := int32(m.memory[uint32(v0+i32(12))])
			t27 := t25 + v3
			v3 = v3 + t26
			t28 := int32(m.memory[uint32(v0+i32(13))])
			t29 := t27 + v3
			v3 = v3 + t28
			t30 := int32(m.memory[uint32(v0+i32(14))])
			t31 := t29 + v3
			v3 = v3 + t30
			t32 := int32(m.memory[uint32(v0+i32(15))])
			t33 := t31 + v3
			v3 = v3 + t32
			v4 = t33 + v3
			v2 = v2 + i32(16)
			goto l6
		}
	case 1:
		t68 := int32(m.memory[uint32(v1)])
		v0 = v3 + t68
		p69 := v0 + i32(-65521)
		if uint32(v0) < uint32(i32(65521)) {
			p69 = v0
		}
		v0 = p69
		t70 := int32(uint32(v0+v4) % uint32(i32(65521)))
		v0 = t70<<16 + v0
		fallthrough
	case 0:
		return v0
	}
}
func (m *Module) fn1726(v0, v1, v2 int32) {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := m.fn1729(t0, v1, v2)
	store32(m.memory[uint32(v0):], uint32(t1))
}
func (m *Module) fn1727(v0, v1 int32) int32 {
	var v2 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	m.fn1837(v2+i32(8), v0)
	{
		t1 := int32(load32(m.memory[int64(uint32(v2))+12:]))
		t2 := v1
		v0 = t1
		if uint32(t2) >= uint32(v0) {
			m.fn158(v1, v0, i32(1294532))
			panic("unreachable")
		}
		t3 := int32(load32(m.memory[int64(uint32(v2))+8:]))
		t4 := int32(load32(m.memory[uint32(t3+v1<<2):]))
		v1 = t4
		m.g0 = v2 + i32(16)
		return v1
	}
}
func (m *Module) fn1728(v0, v1, v2, v3, v4 int32) {
	if uint32(v1) < uint32(v3) {
		goto l0
	}
	m.fn151(i32(0), v1, v3, v4)
	panic("unreachable")
l0:
	store32(m.memory[uint32(v0):], uint32(v2))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v1+i32(1)))
}
func (m *Module) fn1729(v0, v1, v2 int32) int32 {
	var v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13 int32
	t0 := m.g0
	v3 = t0 - i32(5168)
	m.g0 = v3
	v4 = i32(0)
	{
		v5 = (v1+i32(3))&i32(-4) - v1
		if uint32(v5) <= uint32(v2) {
			goto l0
		}
		v6 = i32(4)
		v7 = i32(0)
		v8 = i32(1)
		v9 = i32(0)
		goto l1
	l0:
		m.fn1739(v3+i32(20), v1, v2, v5, i32(1287364))
		t1 := int32(load32(m.memory[int64(uint32(v3))+32:]))
		v1 = t1
		v9 = v1 & i32(3)
		v7 = int32(uint32(v1) >> 2)
		t2 := int32(load32(m.memory[int64(uint32(v3))+28:]))
		v6 = t2
		v8 = v6 + v1&i32(-4)
		t3 := int32(load32(m.memory[int64(uint32(v3))+24:]))
		v2 = t3
		t4 := int32(load32(m.memory[int64(uint32(v3))+20:]))
		v1 = t4
	}
l1:
	t5 := m.fn1838(v1, v2, v0^i32(-1))
	v1 = t5
	store64(m.memory[int64(uint32(v3))+12:], uint64(i64(0)))
	store64(m.memory[int64(uint32(v3))+4:], uint64(i64(0)))
	store32(m.memory[uint32(v3):], uint32(v1))
	t6 := int32(uint32(v7) / uint32(i32(5)))
	v1 = t6
	t7 := v1
	var p8 int32
	if v1 != i32(0) {
		p8 = 1
	}
	v10 = t7 - p8
	v11 = v6
	v12 = i32(0)
l15:
	{
		if v12 == v10 {
			t15 := int32(load32(m.memory[uint32(v3):]))
			v0 = t15
			store32(m.memory[uint32(v3):], uint32(i32(0)))
			t16 := v7
			v1 = v10 * i32(5)
			if uint32(t16) < uint32(v1) {
				m.fn151(v1, v7, v7, i32(1291904))
				panic("unreachable")
			}
			if v7 == v1 {
				goto l8
			}
			v11 = v7 - v1
			v4 = v6 + v1<<2
			v13 = v3 + i32(28)
			v12 = v3 + i32(4132)
			v7 = i32(0)
		l11:
			{
				t17 := v3
				t18 := v4
				v1 = v7 << 2
				t19 := int32(load32(m.memory[uint32(t18+v1):]))
				t21 := t19 ^ v0
				p20 := i32(1287584)
				if uint32(v7) < uint32(i32(5)) {
					p20 = v3 + v1
				}
				t22 := int32(load32(m.memory[uint32(p20):]))
				store32(m.memory[int64(uint32(t17))+4132:], uint32(t21^t22))
				store64(m.memory[int64(uint32(v3))+4124:], uint64(i64(0x400000000)))
				memory_copy(m.memory, uint32(v13), uint32(i32(1296296)), uint32(i32(4096)))
				store64(m.memory[int64(uint32(v3))+4136:], uint64(i64(0)))
				v1 = i32(0)
				v2 = v12
				v0 = i32(0)
			l10:
				{
					if v1 == i32(4096) {
						goto l9
					}
					t23 := int32(m.memory[uint32(v2)])
					v5 = t23
					memory_copy(m.memory, uint32(v3+i32(4144)), uint32(v13+v1), uint32(i32(1024)))
					v1 = v1 + i32(1024)
					v2 = v2 + i32(1)
					t24 := int32(load32(m.memory[uint32(v3+i32(4144)+v5<<2):]))
					v0 = t24 ^ v0
					goto l10
				}
			l9:
				v7 = v7 + i32(1)
				if v7 != v11 {
					goto l11
				}
			}
		l8:
			t25 := m.fn1838(v8, v9, v0)
			v1 = t25
			m.g0 = v3 + i32(5168)
			return v1 ^ i32(-1)
		}
		p9 := v4
		if uint32(v7) < uint32(v4) {
			p9 = v7
		}
		v2 = p9
		v12 = v12 + i32(1)
		v1 = i32(0)
		v0 = v4
	l5:
		{
			if v1 == i32(20) {
				t12 := int32(load32(m.memory[int64(uint32(v3))+36:]))
				store32(m.memory[int64(uint32(v3))+4160:], uint32(t12))
				t13 := int64(load64(m.memory[int64(uint32(v3))+28:]))
				store64(m.memory[int64(uint32(v3))+4152:], uint64(t13))
				t14 := int64(load64(m.memory[int64(uint32(v3))+20:]))
				store64(m.memory[int64(uint32(v3))+4144:], uint64(t14))
				store64(m.memory[uint32(v3):], uint64(i64(0)))
				store64(m.memory[int64(uint32(v3))+8:], uint64(i64(0)))
				v13 = i32(0)
				store32(m.memory[int64(uint32(v3))+16:], uint32(i32(0)))
			l16:
				if v13 == i32(4) {
					v11 = v11 + i32(20)
					v4 = v4 + i32(5)
					goto l15
				}
				v5 = v13<<10 + i32(1287808)
				v1 = i32(0)
			l14:
				{
					if v1 == i32(20) {
						v13 = v13 + i32(1)
						goto l16
					}
					v2 = v3 + i32(4144) + v1
					t26 := int32(load32(m.memory[uint32(v2):]))
					t27 := v2
					v2 = t26
					store32(m.memory[uint32(t27):], uint32(int32(uint32(v2)>>8)))
					v0 = v3 + v1
					t28 := int32(load32(m.memory[uint32(v0):]))
					t29 := int32(load32(m.memory[uint32(v5+v2&i32(255)<<2):]))
					store32(m.memory[uint32(v0):], uint32(t28^t29))
					v1 = v1 + i32(4)
					goto l14
				}
			}
			if v7 == v2 {
				m.fn158(v0, v7, i32(1291920))
				panic("unreachable")
			}
			t10 := int32(load32(m.memory[uint32(v3+v1):]))
			t11 := int32(load32(m.memory[uint32(v11+v1):]))
			store32(m.memory[uint32(v3+i32(20)+v1):], uint32(t10^t11))
			v2 = v2 + i32(1)
			v0 = v0 + i32(1)
			v1 = v1 + i32(4)
			goto l5
		}
	}
}
func (m *Module) fn1730(v0, v1, v2, v3, v4, v5, v6, v7 int32) {
	var v8, v9, v10, v11, v12, v13, v14, v15, v16, v17, v18, v19, v20, v21, v22, v23, v24, v25, v26, v27, v28, v29, v30, v31, v32, v33 int32
	var v34 int64
	t0 := m.g0
	v8 = t0 - i32(80)
	m.g0 = v8
	store64(m.memory[int64(uint32(v8))+32:], uint64(i64(0)))
	store64(m.memory[int64(uint32(v8))+24:], uint64(i64(0)))
	store64(m.memory[int64(uint32(v8))+16:], uint64(i64(0)))
	store64(m.memory[int64(uint32(v8))+8:], uint64(i64(0)))
	v9 = v2 + v3<<1
	v10 = i32(0)
	v11 = v2
	v12 = i32(15)
	{
	l1:
		{
			if v11 == v9 {
				if v10 == 0 {
					store32(m.memory[int64(uint32(v0))+8:], uint32(i32(2)))
					store64(m.memory[uint32(v4):], uint64(i64(0x140000001400000)))
					store64(m.memory[uint32(v0):], uint64(i64(0x100000000)))
					goto l7
				}
				{
					if uint32(v12) > uint32(v10) {
						store32(m.memory[int64(uint32(v8))+72:], uint32(v12))
						store32(m.memory[int64(uint32(v8))+76:], uint32(v10))
						t34 := v8
						v34 = int64(uint32(i32(73))) << 32
						store64(m.memory[int64(uint32(t34))+48:], uint64(v34|int64(uint32(v8+i32(76)))))
						store64(m.memory[int64(uint32(v8))+40:], uint64(v34|int64(uint32(v8+i32(72)))))
						m.fn91(i32(1051106), v8+i32(40), i32(1300392))
						panic("unreachable")
					}
					t6 := v12
					p5 := v10
					if uint32(v6) < uint32(v10) {
						p5 = v6
					}
					p7 := p5
					if uint32(v6) < uint32(v12) {
						p7 = t6
					}
					v15 = p7
					v11 = i32(2)
					v13 = i32(1)
				l6:
					{
						if v11 == i32(32) {
							if v13 == 0 {
								goto l8
							}
							if v1&i32(255) == 0 {
								goto l9
							}
							if v10 != i32(1) {
								goto l9
							}
						l8:
							store64(m.memory[int64(uint32(v8))+64:], uint64(i64(0)))
							store64(m.memory[int64(uint32(v8))+56:], uint64(i64(0)))
							store64(m.memory[int64(uint32(v8))+48:], uint64(i64(0)))
							store64(m.memory[int64(uint32(v8))+40:], uint64(i64(0)))
							v9 = v8 + i32(8) | i32(2)
							v11 = i32(0)
						l43:
							if v11 != i32(28) {
								v13 = v8 + i32(40) + v11
								t32 := int32(load16(m.memory[uint32(v13+i32(2)):]))
								t33 := int32(load16(m.memory[uint32(v9+v11):]))
								store16(m.memory[uint32(v13+i32(4)):], uint16(t32+t33))
								v11 = v11 + i32(2)
								goto l43
							}
							v11 = v3 << 1
							v9 = i32(0)
							v13 = v2
						l15:
							if v11 == 0 {
								v16 = i32(20)
								v11 = v1 & i32(255)
								v17 = v11
								v18 = i32(1)
								v19 = i32(2)
								switch v11 {
								default:
									goto l16
								case 1:
									if uint32(v15) > uint32(i32(10)) {
										goto l19
									}
									v16 = i32(257)
									v19 = i32(1300408)
									v18 = i32(1300470)
									v17 = i32(31)
									goto l16
								case 2:
									if uint32(v15) > uint32(i32(9)) {
										goto l19
									}
									v16 = i32(0)
									v19 = i32(1300502)
									v18 = i32(1300566)
									v17 = i32(32)
								}
							l16:
								v20 = i32_shl(i32(1), v15)
								v21 = v20 + i32(-1)
								v22 = (v16 + i32(-1)) & i32(0xffff)
								v23 = v11 + i32(-1)
								v24 = i32(0)
								v25 = i32(0)
								v26 = i32(0)
								v11 = i32(0)
								v27 = v15
								v28 = i32(0)
								v29 = i32(-1)
							l42:
								{
									t13 := v24
									v30 = i32_shl(i32(1), v27)
									v31 = t13 + v30
									{
									l31:
										{
											if v25 == i32(288) {
												m.fn158(i32(288), i32(288), i32(1300600))
												panic("unreachable")
											}
											{
												t14 := int32(load16(m.memory[uint32(v7+v25<<1):]))
												v13 = t14
												if uint32(v13) >= uint32(v16) {
													t18 := v17
													v13 = (v13 - v16) & i32(0xffff)
													if uint32(t18) <= uint32(v13) {
														m.fn158(v13, v17, i32(1300616))
														panic("unreachable")
													}
													t19 := int32(m.memory[uint32(v18+v13)])
													v1 = t19
													t20 := int32(load16(m.memory[uint32(v19+v13<<1):]))
													v27 = t20
													goto l22
												}
												var p15 int32
												if uint32(v13) < uint32(v22) {
													p15 = 1
												}
												v9 = p15
												p16 := i32(96)
												if v9 != 0 {
													p16 = i32(0)
												}
												v1 = p16
												p17 := i32(0)
												if v9 != 0 {
													p17 = v13
												}
												v27 = p17
												goto l22
											}
										l22:
											t21 := v24
											v11 = i32_shr_u(v11, v28)
											t22 := t21 + v11
											v14 = v12 - v28
											v9 = i32_shl(i32(-1), v14)
											v6 = t22 + v9
											v11 = v4 + (v31+v11+v9)<<2
											v32 = v9 << 2
											v13 = v30
										l25:
											v33 = v6 + v13
											if uint32(v33) >= uint32(v5) {
												m.fn158(v33, v5, i32(1300632))
												panic("unreachable")
											}
											store16(m.memory[uint32(v11):], uint16(v27))
											m.memory[uint32(v11+i32(3))] = byte(v14)
											m.memory[uint32(v11+i32(2))] = byte(v1)
											v11 = v11 + v32
											v13 = v13 + v9
											if v13 != 0 {
												goto l25
											}
											if uint32(v12) > uint32(i32(15)) {
												m.fn158(v12, i32(16), i32(1300648))
												panic("unreachable")
											}
											v26 = i32_shr_u(i32(-0x80000000), v12+i32(-1)) + v26
											v11 = i32_rotr(v26&i32(0xff00ff), i32(8)) | i32_rotr(v26, i32(24))&i32(0xff00ff)
											v11 = int32(uint32(v11)>>4)&i32(252645135) | v11&i32(252645135)<<4
											v11 = int32(uint32(v11)>>2)&i32(0x33333333) | v11&i32(0x33333333)<<2
											v11 = int32(uint32(v11)>>1)&i32(0x55555555) | v11&i32(0x55555555)<<1
											v13 = v25 + i32(1)
											v9 = v8 + i32(8) + v12<<1
											t23 := int32(load16(m.memory[uint32(v9):]))
											t24 := v9
											v9 = t23 + i32(-1)
											store16(m.memory[uint32(t24):], uint16(v9))
											{
												if v9&i32(0xffff) != 0 {
													goto l27
												}
												if v12 == v10 {
													if v26 != 0 {
														goto l36
													}
													goto l37
												}
												if v25 == i32(287) {
													m.fn158(i32(288), i32(288), i32(1300664))
													panic("unreachable")
												}
												t25 := int32(load16(m.memory[uint32(v7+v13<<1):]))
												t26 := v3
												v9 = t25
												if uint32(t26) <= uint32(v9) {
													m.fn158(v9, v3, i32(1300680))
													panic("unreachable")
												}
												t27 := int32(load16(m.memory[uint32(v2+v9<<1):]))
												v12 = t27
											}
										l27:
											v14 = v11 & v21
											v25 = v13
											if uint32(v12) <= uint32(v15) {
												goto l31
											}
											v25 = v13
											if v14 == v29 {
												goto l31
											}
										}
										p28 := v10
										if uint32(v12) > uint32(v10) {
											p28 = v12
										}
										p29 := v15
										if v28 != 0 {
											p29 = v28
										}
										v28 = p29
										v27 = p28 - v28
										v1 = i32_shl(i32(1), v12-v28)
										v6 = v8 + i32(8) + v12<<1
										v9 = v12
									l35:
										{
											if uint32(v9) >= uint32(v10) {
												goto l32
											}
											if uint32(v9) >= uint32(i32(16)) {
												p31 := i32(16)
												if uint32(v12) > uint32(i32(16)) {
													p31 = v12
												}
												m.fn158(p31, i32(16), i32(1300728))
												panic("unreachable")
											}
											t30 := int32(load16(m.memory[uint32(v6):]))
											v1 = v1 - t30
											if v1 < i32(1) {
												goto l34
											}
											v6 = v6 + i32(2)
											v9 = v9 + i32(1)
											v1 = v1 << 1
											goto l35
										}
									}
								l34:
									v27 = v9 - v28
								l32:
									v20 = i32_shl(i32(1), v27) + v20
									switch v23 {
									case 0:
										if uint32(v20) <= uint32(i32(1332)) {
											goto l40
										}
										goto l19
									case 1:
										goto l39
									default:
										goto l40
									}
								l39:
									if uint32(v20) > uint32(i32(592)) {
										goto l19
									}
								l40:
									if uint32(v14) >= uint32(v5) {
										goto l41
									}
									v9 = v4 + v14<<2
									m.memory[int64(uint32(v9))+3] = byte(v15)
									m.memory[int64(uint32(v9))+2] = byte(v27)
									store16(m.memory[uint32(v9):], uint16(v31))
									v24 = v31
									v25 = v13
									v29 = v14
									goto l42
								l41:
								}
								m.fn158(v14, v5, i32(1300744))
								panic("unreachable")
							}
							{
								t10 := int32(load16(m.memory[uint32(v13):]))
								v14 = t10
								if v14 == 0 {
									goto l12
								}
								if uint32(v14) > uint32(i32(15)) {
									m.fn158(v14, i32(16), i32(1300760))
									panic("unreachable")
								}
								v14 = v8 + i32(40) + v14<<1
								t11 := int32(load16(m.memory[uint32(v14):]))
								t12 := v14
								v14 = t11
								store16(m.memory[uint32(t12):], uint16(v14+i32(1)))
								if uint32(v14) >= uint32(i32(288)) {
									m.fn158(v14, i32(288), i32(1300776))
									panic("unreachable")
								}
								store16(m.memory[uint32(v7+v14<<1):], uint16(v9))
							}
						l12:
							v13 = v13 + i32(2)
							v9 = v9 + i32(1)
							v11 = v11 + i32(-2)
							goto l15
						l9:
							store32(m.memory[uint32(v0):], uint32(i32(-1)))
							goto l7
						}
						v9 = v13 << 1
						t8 := int32(load16(m.memory[uint32(v8+i32(8)+v11):]))
						t9 := v9
						v14 = t8
						v13 = t9 - v14
						v11 = v11 + i32(2)
						if uint32(v9) >= uint32(v14) {
							goto l6
						}
					}
					store32(m.memory[uint32(v0):], uint32(i32(-1)))
					goto l7
				}
			}
			t1 := int32(load16(m.memory[uint32(v11):]))
			v13 = t1
			v14 = v11 + i32(2)
			v11 = v14
			if v13 == 0 {
				goto l1
			}
			if uint32(v13) > uint32(i32(15)) {
				m.fn158(v13, i32(16), i32(1300792))
				panic("unreachable")
			}
			v11 = v8 + i32(8) + v13<<1
			t2 := int32(load16(m.memory[uint32(v11):]))
			store16(m.memory[uint32(v11):], uint16(t2+i32(1)))
			p3 := v12
			if uint32(v13) < uint32(v12) {
				p3 = v13
			}
			v12 = p3
			p4 := v10
			if uint32(v13) > uint32(v10) {
				p4 = v13
			}
			v10 = p4
			v11 = v14
			goto l1
		}
	l19:
		store32(m.memory[uint32(v0):], uint32(i32(1)))
		goto l7
	l36:
		if uint32(v5) < uint32(v24) {
			m.fn151(v24, v5, v5, i32(1300696))
			panic("unreachable")
		}
		t35 := v11
		v13 = v5 - v24
		if uint32(t35) >= uint32(v13) {
			m.fn158(v11, v13, i32(1300712))
			panic("unreachable")
		}
		v11 = v4 + v24<<2 + v11<<2
		m.memory[int64(uint32(v11))+3] = byte(v14)
		m.memory[int64(uint32(v11))+2] = byte(i32(64))
		store16(m.memory[uint32(v11):], uint16(i32(0)))
	}
l37:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v20))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v15))
	store32(m.memory[uint32(v0):], uint32(i32(0)))
	goto l7
l7:
	m.g0 = v8 + i32(80)
}
func (m *Module) fn1731(v0, v1 int32) {
	var v2, v3 int32
	v2 = i32(1292484)
	v3 = i32(512)
	{
		t0 := int32(m.memory[int64(uint32(v1))+152])
		switch t0 {
		default:
			goto l0
		case 1:
			v2 = v1 + i32(164)
			v3 = i32(1332)
			goto l0
		case 2:
			v2 = v1 + i32(5492)
			v3 = i32(1332)
			goto l0
		case 3:
			v2 = v1 + i32(10820)
			v3 = i32(592)
		}
	}
l0:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
	store32(m.memory[uint32(v0):], uint32(v2))
}
func (m *Module) fn1732(v0, v1, v2, v3 int32) {
	if uint32(v2) < uint32(i32(321)) {
		goto l0
	}
	m.fn151(v2, i32(320), i32(320), v3)
	panic("unreachable")
l0:
	store32(m.memory[int64(uint32(v0))+4:], uint32(i32(320)-v2))
	store32(m.memory[uint32(v0):], uint32(v1+v2<<1))
}
func (m *Module) fn1733(v0, v1, v2, v3, v4 int32) {
	if uint32(v1) <= uint32(v3) {
		goto l0
	}
	m.fn151(i32(0), v1, v3, v4)
	panic("unreachable")
l0:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
	store32(m.memory[uint32(v0):], uint32(v2))
}
func (m *Module) fn1734(v0, v1, v2 int32) {
	var v3, v4, v5, v6, v7, v8, v9 int32
	t0 := m.g0
	v3 = t0 - i32(32)
	m.g0 = v3
	t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t2 := v3 + i32(24)
	v4 = t1
	t3 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	t4 := v4
	v5 = t3
	v6 = v5 + v2
	v7 = v6 + i32(8)
	p5 := v7
	if uint32(v4) < uint32(v7) {
		p5 = t4
	}
	t6 := int32(load32(m.memory[uint32(v0):]))
	m.fn1724(t2, p5, t6, v4, i32(1287380))
	t7 := int32(load32(m.memory[int64(uint32(v3))+28:]))
	v8 = t7
	t8 := int32(load32(m.memory[int64(uint32(v3))+24:]))
	v9 = t8
	store32(m.memory[int64(uint32(v0))+8:], uint32(v6))
	if uint32(v2) > uint32(v1) {
		if v1 == i32(1) {
			v0 = v5 + i32(-1)
			if uint32(v0) >= uint32(v8) {
				m.fn158(v0, v8, i32(1287396))
				panic("unreachable")
			}
			t14 := int32(m.memory[uint32(v9+v0)])
			v0 = t14
			m.fn1723(v3+i32(16), v5, v9, v8, i32(1287412))
			t15 := int32(load32(m.memory[int64(uint32(v3))+16:]))
			t16 := int32(load32(m.memory[int64(uint32(v3))+20:]))
			m.fn1724(v3+i32(8), v2, t15, t16, i32(1287428))
			t17 := int32(load32(m.memory[int64(uint32(v3))+12:]))
			v5 = t17
			t18 := int32(load32(m.memory[int64(uint32(v3))+8:]))
			v2 = t18
		l13:
			if v5 == 0 {
				goto l6
			}
			m.memory[uint32(v2)] = byte(v0)
			v5 = v5 + i32(-1)
			v2 = v2 + i32(1)
			goto l13
		}
		p12 := v5
		if uint32(v8) < uint32(v5) {
			p12 = v8
		}
		v0 = p12
		v4 = v9 - v1
		v1 = i32(0) - v1
	l11:
		if v2 == 0 {
			goto l6
		}
		v7 = v1 + v5
		if uint32(v7) >= uint32(v8) {
			m.fn158(v7, v8, i32(1287444))
			panic("unreachable")
		}
		{
			if v8 == v0 {
				m.fn158(v5, v8, i32(1287460))
				panic("unreachable")
			}
			t13 := int32(m.memory[uint32(v4+v5)])
			m.memory[uint32(v9+v5)] = byte(t13)
			v2 = v2 + i32(-1)
			v0 = v0 + i32(1)
			v5 = v5 + i32(1)
			goto l11
		}
	}
	if uint32(v5) < uint32(v1) {
		m.fn633(i32(1287476), i32(9), i32(1287488))
		panic("unreachable")
	}
	v0 = v5 - v1
	if uint32(v7) < uint32(v4) {
		if v2 == 0 {
			goto l6
		}
		v5 = v9 + v5
		t9 := v5
		v8 = v9 + v0
		t10 := int64(load64(m.memory[uint32(v8):]))
		store64(m.memory[uint32(t9):], uint64(t10))
		v0 = i32(0) - v1
		v5 = v5 + i32(8)
		v8 = v8 + v2
	l7:
		{
			v2 = v5 + v0
			if uint32(v2) >= uint32(v8) {
				goto l6
			}
			t11 := int64(load64(m.memory[uint32(v2):]))
			store64(m.memory[uint32(v5):], uint64(t11))
			v5 = v5 + i32(8)
			goto l7
		}
	}
	v1 = v0 + v2
	if uint32(v1) > uint32(v8) {
		m.fn151(i32(0), v1, v8, i32(1300808))
		panic("unreachable")
	}
	if uint32(v0) > uint32(v1) {
		m.fn151(v0, v1, v8, i32(1300824))
		panic("unreachable")
	}
	if uint32(v5) > uint32(v8-v2) {
		m.fn91(i32(1287340), i32(43), i32(1287504))
		panic("unreachable")
	}
	if v2 == 0 {
		goto l6
	}
	memory_copy(m.memory, uint32(v9+v5), uint32(v9+v0), uint32(v2))
	goto l6
l6:
	m.g0 = v3 + i32(32)
}
func (m *Module) fn1735(v0 int32) int32 {
	var v1 int32
	v1 = v0 + i32(-64)
	if uint32(v1) < uint32(i32(-63)) {
		p0 := v1
		if uint32(v1) > uint32(v0) {
			p0 = i32(0)
		}
		return p0
	}
	m.fn256(i32(1291996), i32(74), i32(1292072))
	panic("unreachable")
}
func (m *Module) fn1736(v0, v1, v2, v3 int32) {
	var v4, v5, v6, v7, v8, v9 int32
	t0 := m.g0
	v4 = t0 - i32(16)
	m.g0 = v4
	{
		t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		v5 = t1
		t2 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		t3 := v5
		v6 = t2
		t4 := t3 - v6
		v7 = v3 - v2
		if uint32(t4) >= uint32(v7+i32(8)) {
			if v3 == v2 {
				goto l3
			}
			t14 := int32(load32(m.memory[uint32(v0):]))
			v5 = t14 + v6
			t15 := int32(load32(m.memory[uint32(v1):]))
			t16 := v5
			v1 = t15
			v2 = v1 + v2
			t17 := int64(load64(m.memory[uint32(v2):]))
			store64(m.memory[uint32(t16):], uint64(t17))
			v1 = v1 + v3
			v3 = v2 + i32(8)
			v2 = v5 + i32(8)
		l4:
			{
				if uint32(v3) >= uint32(v1) {
					goto l3
				}
				t18 := int64(load64(m.memory[uint32(v3):]))
				store64(m.memory[uint32(v2):], uint64(t18))
				v3 = v3 + i32(8)
				v2 = v2 + i32(8)
				goto l4
			}
		}
		t5 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		v8 = t5
		t6 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t7 := v8
		v9 = t6
		if uint32(t7) > uint32(v9) {
			m.fn151(i32(0), v8, v9, i32(1292280))
			panic("unreachable")
		}
		if uint32(v3) < uint32(v2) {
			goto l2
		}
		if uint32(v3) > uint32(v8) {
			goto l2
		}
		t8 := int32(load32(m.memory[uint32(v1):]))
		v3 = t8
		t9 := int32(load32(m.memory[uint32(v0):]))
		m.fn1723(v4+i32(8), v6, t9, v5, i32(1287536))
		t10 := int32(load32(m.memory[int64(uint32(v4))+8:]))
		t11 := int32(load32(m.memory[int64(uint32(v4))+12:]))
		m.fn1724(v4, v7, t10, t11, i32(1287552))
		t12 := int32(load32(m.memory[uint32(v4):]))
		t13 := int32(load32(m.memory[int64(uint32(v4))+4:]))
		m.fn310(t12, t13, v3+v2, v7, i32(1287568))
		goto l3
	}
l2:
	m.fn151(v2, v3, v8, i32(1287520))
	panic("unreachable")
l3:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v6+v7))
	m.g0 = v4 + i32(16)
}
func (m *Module) fn1737(v0, v1 int32) {
	var v2, v3 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		v2 = t0
		t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		t2 := v2
		v3 = t1
		if uint32(t2) >= uint32(v3) {
			m.fn158(v2, v3, i32(1292296))
			panic("unreachable")
		}
		store32(m.memory[int64(uint32(v0))+8:], uint32(v2+i32(1)))
		t3 := int32(load32(m.memory[uint32(v0):]))
		m.memory[uint32(t3+v2)] = byte(v1)
		return
	}
}
func (m *Module) fn1738(v0 int32) {
	var v1, v2, v3, v4, v5 int32
	var v6 int64
	var v7 int32
	var v8 int64
	var v9, v10, v11, v12 int32
	var v13 int64
	var v14, v15, v16, v17 int32
	var v18 int64
	var v19, v20, v21, v22, v23, v24 int32
	var v25 int64
	var v26, v27 int32
	t0 := m.g0
	v1 = t0 - i32(64)
	m.g0 = v1
	m.memory[int64(uint32(v1))+40] = byte(i32(0))
	store64(m.memory[int64(uint32(v1))+24:], uint64(i64(0)))
	store64(m.memory[int64(uint32(v1))+32:], uint64(i64(0x100000001)))
	t1 := v1 + i32(24)
	v2 = v0 + i32(48)
	m.fn244(t1, v2, i32(6))
	store32(m.memory[int64(uint32(v1))+60:], uint32(i32(0)))
	store64(m.memory[int64(uint32(v1))+52:], uint64(i64(1)))
	t2 := v1 + i32(52)
	v3 = v0 + i32(72)
	m.fn244(t2, v3, i32(3))
	m.fn1731(v1+i32(16), v0)
	t3 := int32(load32(m.memory[int64(uint32(v1))+20:]))
	v4 = t3
	t4 := int32(load32(m.memory[int64(uint32(v1))+16:]))
	v5 = t4
	m.fn1837(v1+i32(8), v0)
	t5 := int64(load32(m.memory[int64(uint32(v0))+156:]))
	v6 = i64_shl(i64(-1), t5)
	t6 := int32(load32(m.memory[int64(uint32(v0))+148:]))
	v7 = t6
	v8 = i64_shl(i64(-1), int64(uint32(v7)))
	t7 := int32(load32(m.memory[int64(uint32(v1))+12:]))
	v9 = t7
	t8 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	v10 = t8
	t9 := int32(load32(m.memory[int64(uint32(v0))+12:]))
	t10 := m.fn1735(t9)
	v11 = t10
	t11 := int32(load32(m.memory[int64(uint32(v1))+32:]))
	v12 = t11
	t12 := int64(load64(m.memory[int64(uint32(v1))+24:]))
	v13 = t12
	{
		{
			t13 := int32(m.memory[int64(uint32(v1))+40])
			v14 = t13
			if uint32(v14) <= uint32(i32(9)) {
				goto l0
			}
			v15 = v12
			v16 = v14
			goto l1
		}
	l0:
		t14 := v1
		v16 = v14 | i32(56)
		m.memory[int64(uint32(t14))+40] = byte(v16)
		t15 := v1
		v15 = v12 + (int32(uint32(v14)>>3) ^ i32(7))
		store32(m.memory[int64(uint32(t15))+32:], uint32(v15))
		t16 := int64(load64(m.memory[uint32(v12):]))
		t17 := v1
		v13 = i64_shl(t16, int64(uint32(v14))) | v13
		store64(m.memory[int64(uint32(t17))+24:], uint64(v13))
	}
l1:
	v17 = v0 + i32(8)
	v18 = v6 ^ i64(-1)
	v8 = v8 ^ i64(-1)
	t18 := int32(load32(m.memory[int64(uint32(v0))+20:]))
	v19 = t18
	t19 := int32(load32(m.memory[int64(uint32(v0))+16:]))
	v20 = t19
	t20 := int32(load32(m.memory[int64(uint32(v1))+36:]))
	v21 = t20
	{
	l36:
		{
			t21 := int64(load64(m.memory[uint32(v15):]))
			v6 = i64_shl(t21, int64(uint32(v16))) | v13
			{
				{
					{
						t22 := v7
						v12 = v16 & i32(255)
						if uint32(t22) <= uint32(v12) {
							goto l2
						}
						t23 := v4
						v14 = int32(v6 & v8)
						if uint32(t23) > uint32(v14) {
							goto l3
						}
						m.fn158(v14, v4, i32(1287604))
						panic("unreachable")
					}
				l2:
					t24 := v4
					v14 = int32(v13 & v8)
					if uint32(t24) <= uint32(v14) {
						m.fn158(v14, v4, i32(1287620))
						panic("unreachable")
					}
				}
			l3:
				v16 = v16 | i32(56)
				v15 = v15 + (int32(uint32(v12)>>3) ^ i32(7))
				v14 = v5 + v14<<2
				t25 := int32(load16(m.memory[uint32(v14):]))
				v22 = t25
				t26 := int32(m.memory[int64(uint32(v14))+3])
				v12 = t26
				t27 := int32(m.memory[int64(uint32(v14))+2])
				v14 = t27
				if v14 == 0 {
					m.fn1737(v1+i32(52), v22)
					{
						t28 := v4
						v13 = i64_shr_u(v6, int64(uint32(v12)))
						v14 = int32(v13 & v8)
						if uint32(t28) <= uint32(v14) {
							m.fn158(v14, v4, i32(1287636))
							panic("unreachable")
						}
						v16 = v16 - v12
						v14 = v5 + v14<<2
						t29 := int32(load16(m.memory[uint32(v14):]))
						v22 = t29
						t30 := int32(m.memory[int64(uint32(v14))+3])
						v12 = t30
						t31 := int32(m.memory[int64(uint32(v14))+2])
						v14 = t31
						if v14 != 0 {
							goto l6
						}
						m.fn1737(v1+i32(52), v22)
						t32 := v4
						v13 = i64_shr_u(v13, int64(uint32(v12)))
						v23 = int32(v13 & v8)
						if uint32(t32) <= uint32(v23) {
							m.fn158(v23, v4, i32(1287652))
							panic("unreachable")
						}
						v16 = v16 - v12
						v24 = i32(0)
						goto l9
					}
				}
				v13 = v6
				goto l6
			}
		l6:
			v24 = i32(1)
		l9:
			{
				{
					{
						{
						l12:
							switch v24 {
							case 0:
								v14 = v5 + v23<<2
								t33 := int32(load16(m.memory[uint32(v14):]))
								v22 = t33
								t34 := int32(m.memory[int64(uint32(v14))+3])
								v12 = t34
								t35 := int32(m.memory[int64(uint32(v14))+2])
								v14 = t35
								v24 = i32(1)
								goto l12
							default:
								v16 = v16 - v12
								v13 = i64_shr_u(v13, int64(uint32(v12)))
								if v14&i32(255) != 0 {
									goto l13
								}
								m.fn1737(v1+i32(52), v22)
								goto l14
							l13:
								{
									{
										{
											if v14&i32(16) == 0 {
												if v14&i32(64) != 0 {
													if v14&i32(32) != 0 {
														m.memory[uint32(v0)] = byte(i32(12))
														v14 = int32(uint32(v16&i32(248)) >> 3)
														goto l23
													}
													v14 = i32(1067904)
													v12 = i32(28)
													goto l22
												}
												t38 := v4
												v23 = (v22 + int32(v13&(i64_shl(i64(-1), int64(uint32(v14))&i64(47))^i64(-1)))) & i32(0xffff)
												if uint32(t38) > uint32(v23) {
													v24 = i32(0)
													goto l12
												}
												m.fn158(v23, v4, i32(1287668))
												panic("unreachable")
											}
											t36 := v9
											t37 := v13
											v25 = int64(uint32(v14)) & i64(15)
											v6 = i64_shr_u(t37, v25)
											v12 = int32(v6 & v18)
											if uint32(t36) <= uint32(v12) {
												m.fn158(v12, v9, i32(1287684))
												panic("unreachable")
											}
											v24 = int32(v13 & (i64_shl(i64(-1), v25) ^ i64(-1)))
											v26 = v10 + v12<<2
											v14 = v16 - v14&i32(15)
											if uint32(v14&i32(255)) < uint32(i32(28)) {
												goto l17
											}
											v27 = v15
											v16 = v14
											goto l18
										}
									l17:
										v16 = v14 | i32(56)
										v27 = v15 + (int32(uint32(v14&i32(248))>>3) ^ i32(7))
										t39 := int64(load64(m.memory[uint32(v15):]))
										v6 = i64_shl(t39, int64(uint32(v14))&i64(255)) | v6
									}
								l18:
									v15 = v22 + v24
									t40 := int32(load16(m.memory[uint32(v26):]))
									v24 = t40
									t41 := int32(m.memory[int64(uint32(v26))+3])
									v12 = t41
									t42 := int32(m.memory[int64(uint32(v26))+2])
									v14 = t42
								l27:
									v16 = v16 - v12
									v6 = i64_shr_u(v6, int64(uint32(v12)))
									if v14&i32(16) != 0 {
										v16 = v16 - v14&i32(15)
										t47 := v6
										v25 = int64(uint32(v14)) & i64(15)
										v13 = i64_shr_u(t47, v25)
										{
											v12 = (v24 + int32(v6&(i64_shl(i64(-1), v25)^i64(-1)))) & i32(0xffff)
											t48 := int32(load32(m.memory[int64(uint32(v1))+60:]))
											t49 := v12
											v14 = t48
											if uint32(t49) > uint32(v14) {
												v14 = v12 - v14
												if uint32(v14) > uint32(v20) {
													t50 := int32(m.memory[int64(uint32(v0))+1])
													if t50&i32(4) == 0 {
														m.fn91(i32(1287716), i32(85), i32(1287760))
														panic("unreachable")
													}
													v14 = i32(1067490)
													v12 = i32(30)
													v15 = v27
													goto l22
												}
												if v19 != 0 {
													if uint32(v19) < uint32(v14) {
														goto l34
													}
													v22 = v19 - v14
													goto l32
												}
												v22 = v11 - v14
												goto l32
											}
											m.fn1734(v1+i32(52), v12, v15&i32(0xffff))
											goto l29
										}
									}
									if v14&i32(64) == 0 {
										t43 := v9
										v14 = (v24 + int32(v6&(i64_shl(i64(-1), int64(uint32(v14))&i64(47))^i64(-1)))) & i32(0xffff)
										if uint32(t43) <= uint32(v14) {
											m.fn158(v14, v9, i32(1287700))
											panic("unreachable")
										}
										v14 = v10 + v14<<2
										t44 := int32(load16(m.memory[uint32(v14):]))
										v24 = t44
										t45 := int32(m.memory[int64(uint32(v14))+3])
										v12 = t45
										t46 := int32(m.memory[int64(uint32(v14))+2])
										v14 = t46
										goto l27
									}
									v14 = i32(1067932)
									v12 = i32(22)
									v13 = v6
									v15 = v27
									goto l22
								}
							l34:
							}
							t51 := v11
							v14 = v14 - v19
							v22 = t51 - v14
							if uint32(v14) >= uint32(v15&i32(0xffff)) {
								goto l32
							}
							m.fn1736(v1+i32(52), v17, v22, v11)
							v15 = v15 - v14
							v22 = i32(0)
							v14 = v19
						}
					l32:
						t52 := v1 + i32(52)
						t53 := v17
						t54 := v22
						v24 = v15 & i32(0xffff)
						p55 := v14
						if uint32(v24) < uint32(v14) {
							p55 = v24
						}
						m.fn1736(t52, t53, t54, p55+v22)
						if uint32(v24) <= uint32(v14) {
							goto l29
						}
						m.fn1734(v1+i32(52), v12, v24-v14)
					}
				l29:
					v15 = v27
					goto l14
				l22:
					m.memory[uint32(v0)] = byte(i32(30))
					t56 := int64(load64(m.memory[int64(uint32(v1))+52:]))
					store64(m.memory[uint32(v3):], uint64(t56))
					t57 := int32(load32(m.memory[int64(uint32(v1))+60:]))
					store32(m.memory[int64(uint32(v3))+8:], uint32(t57))
					store32(m.memory[int64(uint32(v0))+136:], uint32(v12))
					store32(m.memory[int64(uint32(v0))+132:], uint32(v14))
					m.memory[int64(uint32(v1))+40] = byte(v16 & i32(7))
					t58 := int64(load64(m.memory[int64(uint32(v1))+40:]))
					store64(m.memory[int64(uint32(v2))+16:], uint64(t58))
					t59 := v2
					v6 = v13 & (i64_shl(i64(-1), int64(uint32(v16))&i64(7)) ^ i64(-1))
					store64(m.memory[uint32(t59):], uint64(v6))
					store32(m.memory[int64(uint32(v1))+32:], uint32(v15-int32(uint32(v16&i32(248))>>3)))
					t60 := int64(load64(m.memory[int64(uint32(v1))+32:]))
					store64(m.memory[int64(uint32(v2))+8:], uint64(t60))
					store64(m.memory[int64(uint32(v1))+24:], uint64(v6))
					goto l35
				}
			l14:
				t61 := v21 - v15
				v14 = int32(uint32(v16)>>3) & i32(31)
				if uint32(t61+v14) <= uint32(i32(14)) {
					goto l23
				}
				t62 := int32(load32(m.memory[int64(uint32(v1))+56:]))
				t63 := int32(load32(m.memory[int64(uint32(v1))+60:]))
				if uint32(t62-t63) > uint32(i32(259)) {
					goto l36
				}
			}
		l23:
		}
		t64 := int64(load64(m.memory[int64(uint32(v1))+52:]))
		store64(m.memory[uint32(v3):], uint64(t64))
		t65 := int32(load32(m.memory[int64(uint32(v1))+60:]))
		store32(m.memory[int64(uint32(v3))+8:], uint32(t65))
		store32(m.memory[int64(uint32(v1))+32:], uint32(v15-v14))
		t66 := int64(load64(m.memory[int64(uint32(v1))+32:]))
		store64(m.memory[int64(uint32(v2))+8:], uint64(t66))
		m.memory[int64(uint32(v1))+40] = byte(v16 & i32(7))
		t67 := int64(load64(m.memory[int64(uint32(v1))+40:]))
		store64(m.memory[int64(uint32(v2))+16:], uint64(t67))
		t68 := v2
		v6 = v13 & (i64_shl(i64(-1), int64(uint32(v16))&i64(7)) ^ i64(-1))
		store64(m.memory[uint32(t68):], uint64(v6))
		store64(m.memory[int64(uint32(v1))+24:], uint64(v6))
	}
l35:
	m.g0 = v1 + i32(64)
}
func (m *Module) fn1739(v0, v1, v2, v3, v4 int32) {
	if uint32(v2) >= uint32(v3) {
		goto l0
	}
	m.fn91(i32(1301172), i32(19), v4)
	panic("unreachable")
l0:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
	store32(m.memory[uint32(v0):], uint32(v1))
	store32(m.memory[int64(uint32(v0))+12:], uint32(v2-v3))
	store32(m.memory[int64(uint32(v0))+8:], uint32(v1+v3))
}
func (m *Module) fn1740(v0, v1, v2, v3, v4 int32) {
	if uint32(v1) <= uint32(v3) {
		goto l0
	}
	m.fn151(i32(0), v1, v3, v4)
	panic("unreachable")
l0:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
	store32(m.memory[uint32(v0):], uint32(v2))
}
func (m *Module) fn1741(v0, v1, v2, v3, v4 int32) {
	var v5 int32
	t0 := m.g0
	v5 = t0 - i32(16)
	m.g0 = v5
	m.fn1726(v0, v3, v4)
	m.fn1740(v5+i32(8), v4, v1, v2, i32(1295104))
	t1 := int32(load32(m.memory[int64(uint32(v5))+8:]))
	t2 := int32(load32(m.memory[int64(uint32(v5))+12:]))
	m.fn1689(t1, t2, v3, v4, i32(1295120))
	m.g0 = v5 + i32(16)
}
func (m *Module) fn1742(v0, v1, v2, v3, v4 int32) int32 {
	var v5 int32
	t0 := m.g0
	v5 = t0 - i32(16)
	m.g0 = v5
	m.fn1740(v5+i32(8), v4, v1, v2, i32(1295136))
	t1 := int32(load32(m.memory[int64(uint32(v5))+8:]))
	t2 := int32(load32(m.memory[int64(uint32(v5))+12:]))
	m.fn1689(t1, t2, v3, v4, i32(1295152))
	t3 := m.fn1725(v0, v3, v4)
	v4 = t3
	m.g0 = v5 + i32(16)
	return v4
}
func (m *Module) fn1743() {
	m.fn91(i32(1280660), i32(57), i32(1280644))
	panic("unreachable")
}
func (m *Module) fn1744(v0, v1, v2 int32) {
	var v3, v4, v5, v6, v7, v8, v9 int32
	v3 = i32(1)
	v4 = i32(0)
l7:
	v5 = v3
	v6 = i32(1)
l6:
	v3 = i32(0)
l8:
	{
		{
			v7 = v5 + v3
			if uint32(v7) >= uint32(i32(4)) {
				store32(m.memory[int64(uint32(v0))+4:], uint32(v6))
				store32(m.memory[uint32(v0):], uint32(v4))
				return
			}
			v8 = v4 + v3
			if uint32(v8) >= uint32(i32(4)) {
				m.fn158(v8, i32(4), i32(1280976))
				panic("unreachable")
			}
			t0 := int32(m.memory[uint32(v1+v7)])
			v9 = t0
			t1 := int32(m.memory[uint32(v1+v8)])
			v8 = t1
			{
				{
					if v2 == 0 {
						goto l2
					}
					v9 = v9 & i32(255)
					t2 := v9
					v8 = v8 & i32(255)
					if uint32(t2) > uint32(v8) {
						goto l3
					}
					if uint32(v9) < uint32(v8) {
						goto l4
					}
					goto l5
				}
			l2:
				v9 = v9 & i32(255)
				t3 := v9
				v8 = v8 & i32(255)
				if uint32(t3) < uint32(v8) {
					goto l3
				}
				if uint32(v9) <= uint32(v8) {
					goto l5
				}
			}
		l4:
			v5 = v7 + i32(1)
			v6 = v5 - v4
			goto l6
		}
	l3:
		v3 = v5 + i32(1)
		v4 = v5
		goto l7
	l5:
		v3 = v3 + i32(1)
		t4 := v3
		var p5 int32
		if v3 == v6 {
			p5 = 1
		}
		v7 = p5
		p6 := t4
		if v7 != 0 {
			p6 = i32(0)
		}
		v3 = p6
		p7 := i32(0)
		if v7 != 0 {
			p7 = v6
		}
		v5 = p7 + v5
		goto l8
	}
}
func (m *Module) fn1745(v0, v1, v2, v3, v4, v5, v6 int32) {
	var v7, v8, v9, v10, v11, v12 int32
	var v13 int64
	var v14, v15, v16, v17 int32
	t0 := m.g0
	v7 = t0 - i32(16)
	m.g0 = v7
	{
		if uint32(v4) < uint32(i32(16)) {
			t14 := int32(load32(m.memory[int64(uint32(v1))+40:]))
			t15 := int32(load32(m.memory[int64(uint32(v1))+44:]))
			m.fn1746(v7+i32(8), t14, t15, v3, v4, v5, v6)
			t16 := int32(load32(m.memory[int64(uint32(v7))+12:]))
			v15 = t16
			t17 := int32(load32(m.memory[int64(uint32(v7))+8:]))
			v9 = t17
			goto l4
		}
		v8 = v6 + i32(-1)
		v9 = i32(1)
		t1 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		v10 = t1
		t2 := int32(load32(m.memory[uint32(v1):]))
		if t2 != i32(1) {
			if v6 == 0 {
				goto l2
			}
			v17 = v6 - v10
			t18 := int64(load64(m.memory[int64(uint32(v1))+8:]))
			v13 = t18
			t19 := int32(load32(m.memory[int64(uint32(v1))+16:]))
			v11 = t19
			v9 = i32(0)
			v16 = i32(0)
			v1 = i32(0)
		l16:
			v12 = v16
			v15 = v1
			v1 = v15 + v6
			if uint32(v1) > uint32(v4) {
				goto l4
			}
			{
				v14 = v15 + v8
				if uint32(v14) >= uint32(v4) {
					m.fn158(v14, v4, i32(1280816))
					panic("unreachable")
				}
				v16 = i32(0)
				t20 := int64(m.memory[uint32(v3+v14)])
				if i64_shr_u(v13, t20)&i64(1) == 0 {
					goto l16
				}
				v16 = v3 + v15
				p21 := v11
				if uint32(v12) > uint32(v11) {
					p21 = v12
				}
				v14 = p21
				v1 = v14
			l21:
				if uint32(v1) < uint32(v6) {
					if uint32(v15+v1) >= uint32(v4) {
						t28 := v4
						v1 = v15 + v14
						p29 := v1
						if uint32(v4) > uint32(v1) {
							p29 = t28
						}
						m.fn158(p29, v4, i32(1280832))
						panic("unreachable")
					}
					t22 := int32(m.memory[uint32(v5+v1)])
					t23 := int32(m.memory[uint32(v16+v1)])
					if t22 != t23 {
						goto l20
					}
					v1 = v1 + i32(1)
					goto l21
				}
				v1 = v11
			l26:
				if uint32(v1) <= uint32(v12) {
					if uint32(v12) >= uint32(v6) {
						m.fn158(v12, v6, i32(1280880))
						panic("unreachable")
					}
					{
						v14 = v15 + v12
						if uint32(v14) >= uint32(v4) {
							m.fn158(v14, v4, i32(1280896))
							panic("unreachable")
						}
						v16 = v17
						v1 = v10
						t26 := int32(m.memory[uint32(v5+v12)])
						t27 := int32(m.memory[uint32(v3+v14)])
						if t26 == t27 {
							goto l8
						}
						goto l27
					}
				}
				if uint32(v1) >= uint32(v6) {
					m.fn158(v1, v6, i32(1280848))
					panic("unreachable")
				}
				v14 = v15 + v1
				if uint32(v14) >= uint32(v4) {
					m.fn158(v14, v4, i32(1280864))
					panic("unreachable")
				}
				{
					t24 := int32(m.memory[uint32(v5+v1)])
					t25 := int32(m.memory[uint32(v16+v1)])
					if t24 != t25 {
						goto l25
					}
					v1 = v1 + i32(-1)
					goto l26
				}
			l25:
				v16 = v17
				v1 = v10
				goto l27
			l20:
				v1 = v1 - v11 + i32(1)
				v16 = i32(0)
			l27:
				v1 = v1 + v15
				goto l16
			}
		}
		if v6 == 0 {
			goto l2
		}
		t3 := int32(load32(m.memory[int64(uint32(v1))+16:]))
		v11 = t3
		p4 := v6
		if uint32(v11) > uint32(v6) {
			p4 = v11
		}
		v12 = p4
		t5 := int64(load64(m.memory[int64(uint32(v1))+8:]))
		v13 = t5
		v9 = i32(0)
		var p6 int32
		if uint32(v11+i32(-1)) >= uint32(v6) {
			p6 = 1
		}
		v14 = p6
		v1 = i32(0)
	l6:
		{
			v15 = v1
			v1 = v15 + v6
			if uint32(v1) <= uint32(v4) {
				goto l3
			}
			goto l4
		l3:
			v16 = v15 + v8
			if uint32(v16) >= uint32(v4) {
				m.fn158(v16, v4, i32(1280736))
				panic("unreachable")
			}
			t7 := int64(m.memory[uint32(v3+v16)])
			if i64_shr_u(v13, t7)&i64(1) == 0 {
				goto l6
			}
			v16 = v3 + v15
			v1 = v11
		l14:
			if v12 != v1 {
				goto l7
			}
			v1 = v11
		l11:
			{
				if v1 == 0 {
					goto l8
				}
				v1 = v1 + i32(-1)
				if v14 != 0 {
					m.fn158(v1, v6, i32(1280768))
					panic("unreachable")
				}
				v16 = v1 + v15
				if uint32(v16) >= uint32(v4) {
					m.fn158(v16, v4, i32(1280784))
					panic("unreachable")
				}
				t8 := int32(m.memory[uint32(v5+v1)])
				t9 := int32(m.memory[uint32(v3+v16)])
				if t8 == t9 {
					goto l11
				}
			}
			v1 = v15 + v10
			goto l6
		l7:
			if uint32(v15+v1) >= uint32(v4) {
				goto l12
			}
			{
				t10 := int32(m.memory[uint32(v5+v1)])
				t11 := int32(m.memory[uint32(v16+v1)])
				if t10 != t11 {
					goto l13
				}
				v1 = v1 + i32(1)
				goto l14
			}
		l13:
			v1 = v15 - v11 + v1 + i32(1)
			goto l6
		l12:
		}
		t12 := v4
		v1 = v15 + v11
		p13 := v1
		if uint32(v4) > uint32(v1) {
			p13 = t12
		}
		m.fn158(p13, v4, i32(1280752))
		panic("unreachable")
	}
l8:
	v9 = i32(1)
	goto l4
l2:
	v15 = i32(0)
l4:
	store32(m.memory[uint32(v0):], uint32(v9))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v15))
	m.g0 = v7 + i32(16)
}
func (m *Module) fn1746(v0, v1, v2, v3, v4, v5, v6 int32) {
	var v7, v8, v9 int32
	if uint32(v6) > uint32(v4) {
		goto l0
	}
	v7 = v3 + v4
	v8 = v3 + v6
	v4 = i32(0)
	v9 = v3
l6:
	if uint32(v9) < uint32(v8) {
		t3 := int32(m.memory[uint32(v9)])
		v4 = v4<<1 + t3
		v9 = v9 + i32(1)
		goto l6
	}
	v8 = v7 - v6
	v9 = v3
l4:
	{
		{
			if v1 != v4 {
				goto l2
			}
			t0 := m.fn320(v9, v5, v6)
			if t0 != 0 {
				v9 = v9 - v3
				v4 = i32(1)
				goto l5
			}
		}
	l2:
		if uint32(v9) >= uint32(v8) {
			goto l0
		}
		t1 := int32(m.memory[uint32(v9)])
		t2 := int32(m.memory[uint32(v9+v6)])
		v4 = (v4-v2*t1)<<1 + t2
		v9 = v9 + i32(1)
		goto l4
	}
l0:
	v4 = i32(0)
l5:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v9))
	store32(m.memory[uint32(v0):], uint32(v4))
}
func (m *Module) fn1747(v0, v1, v2, v3 int32) {
	var v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14, v15 int32
	t0 := m.g0
	v4 = t0 - i32(16)
	m.g0 = v4
	t1 := int32(m.memory[int64(uint32(v1))+6])
	v5 = t1
	v6 = v5 * i32(16843009)
	t2 := int32(m.memory[int64(uint32(v1))+5])
	v7 = t2
	t3 := int32(m.memory[int64(uint32(v1))+4])
	v8 = t3
	v9 = i32(0)
	t4 := int32(m.memory[int64(uint32(v1))+7])
	v10 = t4 & i32(255)
	v11 = i32(0)
l11:
	{
		m.fn148(v4+i32(8), v11, v2, v3, i32(1280704))
		t5 := int32(load32(m.memory[int64(uint32(v4))+12:]))
		v12 = t5
		if v12 == 0 {
			goto l12
		}
		t6 := int32(load32(m.memory[int64(uint32(v4))+8:]))
		v13 = t6
		v14 = v13 + v12
		v1 = v13
		if uint32(v12) <= uint32(i32(3)) {
		l9:
			{
				if uint32(v1) >= uint32(v14) {
					goto l12
				}
				t12 := int32(m.memory[uint32(v1)])
				if v5 == t12 {
					goto l6
				}
				v1 = v1 + i32(1)
				goto l9
			}
		}
		v1 = v13
		{
			t7 := int32(load32(m.memory[uint32(v13):]))
			v15 = t7 ^ v6
			if (i32(16843008)-v15|v15)&i32(-2139062144) != i32(-2139062144) {
			l8:
				{
					if uint32(v1) >= uint32(v14) {
						goto l12
					}
					t11 := int32(m.memory[uint32(v1)])
					if v5 == t11 {
						goto l6
					}
					v1 = v1 + i32(1)
					goto l8
				}
			}
			v1 = v13&i32(-4) + i32(4)
			if uint32(v12) <= uint32(i32(8)) {
			l7:
				{
					if uint32(v1) >= uint32(v14) {
						goto l12
					}
					t10 := int32(m.memory[uint32(v1)])
					if v5 == t10 {
						goto l6
					}
					v1 = v1 + i32(1)
					goto l7
				}
			}
			v15 = v14 + i32(-8)
		l5:
			{
				if uint32(v1) > uint32(v15) {
					goto l10
				}
				t8 := int32(load32(m.memory[uint32(v1):]))
				v12 = t8 ^ v6
				if (i32(16843008)-v12|v12)&i32(-2139062144) != i32(-2139062144) {
					goto l10
				}
				t9 := int32(load32(m.memory[uint32(v1+i32(4)):]))
				v12 = t9 ^ v6
				if (i32(16843008)-v12|v12)&i32(-2139062144) != i32(-2139062144) {
					goto l10
				}
				v1 = v1 + i32(8)
				goto l5
			}
		}
	l10:
		{
			if uint32(v1) >= uint32(v14) {
				goto l12
			}
			t13 := int32(m.memory[uint32(v1)])
			if v5 == t13 {
				goto l6
			}
			v1 = v1 + i32(1)
			goto l10
		}
	l6:
		v14 = v1 - v13 + v11
		v1 = v14 - v8
		v11 = v14 + i32(1)
		if uint32(v14) < uint32(v8) {
			goto l11
		}
		v14 = v1 + v7
		if uint32(v14) < uint32(v1) {
			goto l11
		}
		if uint32(v14) >= uint32(v3) {
			goto l11
		}
		t14 := int32(m.memory[uint32(v2+v14)])
		if t14 != v10 {
			goto l11
		}
	}
	v9 = i32(1)
	goto l12
l12:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
	store32(m.memory[uint32(v0):], uint32(v9))
	m.g0 = v4 + i32(16)
}
func (m *Module) fn1748(v0, v1, v2, v3, v4, v5, v6 int32) {
	var v7, v8, v9, v10, v11, v12, v13 int32
	var v14 int64
	var v15, v16, v17, v18 int32
	t0 := m.g0
	v7 = t0 - i32(48)
	m.g0 = v7
	{
		if uint32(v4) < uint32(i32(16)) {
			t19 := int32(load32(m.memory[int64(uint32(v1))+40:]))
			t20 := int32(load32(m.memory[int64(uint32(v1))+44:]))
			m.fn1746(v7+i32(8), t19, t20, v3, v4, v5, v6)
			t21 := int32(load32(m.memory[int64(uint32(v7))+12:]))
			v16 = t21
			t22 := int32(load32(m.memory[int64(uint32(v7))+8:]))
			v10 = t22
			goto l17
		}
		v8 = v6 + i32(-1)
		v9 = v1 + i32(24)
		v10 = i32(1)
		t1 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		v11 = t1
		t2 := int32(load32(m.memory[uint32(v1):]))
		if t2 != i32(1) {
			goto l1
		}
		if v6 != 0 {
			t3 := int32(load32(m.memory[int64(uint32(v1))+16:]))
			v12 = t3
			p4 := v6
			if uint32(v12) > uint32(v6) {
				p4 = v12
			}
			v13 = p4
			t5 := int64(load64(m.memory[int64(uint32(v1))+8:]))
			v14 = t5
			v10 = i32(0)
			var p6 int32
			if uint32(v12+i32(-1)) >= uint32(v6) {
				p6 = 1
			}
			v15 = p6
			v16 = i32(0)
		l16:
			if uint32(v16+v6) > uint32(v4) {
				goto l17
			}
			{
				t7 := m.fn1749(v2)
				if t7 == 0 {
					goto l5
				}
				m.fn148(v7+i32(24), v16, v3, v4, i32(1280720))
				t8 := int32(load32(m.memory[int64(uint32(v7))+24:]))
				t9 := int32(load32(m.memory[int64(uint32(v7))+28:]))
				m.fn1750(v7+i32(16), v2, v9, t8, t9)
				t10 := int32(load32(m.memory[int64(uint32(v7))+16:]))
				if t10&i32(1) == 0 {
					goto l17
				}
				t11 := int32(load32(m.memory[int64(uint32(v7))+20:]))
				v16 = t11 + v16
				if uint32(v16+v6) > uint32(v4) {
					goto l17
				}
			}
		l5:
			v1 = v16 + v8
			if uint32(v1) >= uint32(v4) {
				m.fn158(v1, v4, i32(1280736))
				panic("unreachable")
			}
			v17 = v6
			{
				t12 := int64(m.memory[uint32(v3+v1)])
				if i64_shr_u(v14, t12)&i64(1) == 0 {
					goto l7
				}
				v17 = v3 + v16
				v1 = v12
			l15:
				if v13 != v1 {
					if uint32(v16+v1) >= uint32(v4) {
						t17 := v4
						v1 = v12 + v16
						p18 := v1
						if uint32(v4) > uint32(v1) {
							p18 = t17
						}
						m.fn158(p18, v4, i32(1280752))
						panic("unreachable")
					}
					{
						t15 := int32(m.memory[uint32(v5+v1)])
						t16 := int32(m.memory[uint32(v17+v1)])
						if t15 != t16 {
							v16 = v16 - v12 + v1 + i32(1)
							goto l16
						}
						v1 = v1 + i32(1)
						goto l15
					}
				}
				v1 = v12
			l12:
				{
					if v1 == 0 {
						goto l9
					}
					v1 = v1 + i32(-1)
					if v15 != 0 {
						m.fn158(v1, v6, i32(1280768))
						panic("unreachable")
					}
					v17 = v1 + v16
					if uint32(v17) >= uint32(v4) {
						m.fn158(v17, v4, i32(1280784))
						panic("unreachable")
					}
					t13 := int32(m.memory[uint32(v5+v1)])
					t14 := int32(m.memory[uint32(v3+v17)])
					if t13 == t14 {
						goto l12
					}
				}
				v17 = v11
				goto l7
			}
		l7:
			v16 = v17 + v16
			goto l16
		}
		v10 = i32(1)
		goto l3
	}
l1:
	if v6 != 0 {
		v18 = v6 - v11
		t23 := int64(load64(m.memory[int64(uint32(v1))+8:]))
		v14 = t23
		t24 := int32(load32(m.memory[int64(uint32(v1))+16:]))
		v12 = t24
		v10 = i32(0)
		v16 = i32(0)
		{
		l37:
			if uint32(v16+v6) > uint32(v4) {
				goto l19
			}
			{
				{
					t25 := m.fn1749(v2)
					if t25 != 0 {
						goto l20
					}
					p26 := v12
					if uint32(v10) > uint32(v12) {
						p26 = v10
					}
					v15 = p26
					goto l21
				}
			l20:
				m.fn148(v7+i32(40), v16, v3, v4, i32(1280800))
				t27 := int32(load32(m.memory[int64(uint32(v7))+40:]))
				t28 := int32(load32(m.memory[int64(uint32(v7))+44:]))
				m.fn1750(v7+i32(32), v2, v9, t27, t28)
				t29 := int32(load32(m.memory[int64(uint32(v7))+32:]))
				if t29&i32(1) == 0 {
					goto l19
				}
				v10 = i32(0)
				v15 = v12
				t30 := int32(load32(m.memory[int64(uint32(v7))+36:]))
				v16 = t30 + v16
				if uint32(v16+v6) > uint32(v4) {
					goto l17
				}
			}
		l21:
			v1 = v16 + v8
			if uint32(v1) < uint32(v4) {
				t31 := int64(m.memory[uint32(v3+v1)])
				if i64_shr_u(v14, t31)&i64(1) == 0 {
					v16 = v16 + v6
					v10 = i32(0)
					goto l37
				}
				p32 := v6
				if uint32(v15) > uint32(v6) {
					p32 = v15
				}
				v13 = p32
				v17 = v3 + v16
				v1 = v15
			l28:
				if v13 != v1 {
					if uint32(v16+v1) >= uint32(v4) {
						t39 := v4
						v1 = v15 + v16
						p40 := v1
						if uint32(v4) > uint32(v1) {
							p40 = t39
						}
						m.fn158(p40, v4, i32(1280832))
						panic("unreachable")
					}
					t33 := int32(m.memory[uint32(v5+v1)])
					t34 := int32(m.memory[uint32(v17+v1)])
					if t33 != t34 {
						goto l27
					}
					v1 = v1 + i32(1)
					goto l28
				}
				v1 = v12
			l33:
				if uint32(v1) <= uint32(v10) {
					if uint32(v10) >= uint32(v6) {
						m.fn158(v10, v6, i32(1280880))
						panic("unreachable")
					}
					{
						v17 = v16 + v10
						if uint32(v17) >= uint32(v4) {
							m.fn158(v17, v4, i32(1280896))
							panic("unreachable")
						}
						v13 = v5 + v10
						v10 = v18
						v1 = v11
						t37 := int32(m.memory[uint32(v13)])
						t38 := int32(m.memory[uint32(v3+v17)])
						if t37 == t38 {
							goto l9
						}
						goto l34
					}
				}
				if uint32(v1) >= uint32(v6) {
					m.fn158(v1, v6, i32(1280848))
					panic("unreachable")
				}
				v13 = v16 + v1
				if uint32(v13) >= uint32(v4) {
					m.fn158(v13, v4, i32(1280864))
					panic("unreachable")
				}
				{
					t35 := int32(m.memory[uint32(v5+v1)])
					t36 := int32(m.memory[uint32(v17+v1)])
					if t35 != t36 {
						goto l32
					}
					v1 = v1 + i32(-1)
					goto l33
				}
			l32:
				v10 = v18
				v1 = v11
				goto l34
			l27:
				v1 = v1 - v12 + i32(1)
				v10 = i32(0)
			l34:
				v16 = v1 + v16
				goto l37
			}
			m.fn158(v1, v4, i32(1280816))
			panic("unreachable")
		}
	l19:
		v10 = i32(0)
		goto l17
	}
l3:
	v16 = i32(0)
	goto l17
l9:
	v10 = i32(1)
	goto l17
l17:
	store32(m.memory[uint32(v0):], uint32(v10))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v16))
	m.g0 = v7 + i32(48)
}
func (m *Module) fn1749(v0 int32) int32 {
	var v1, v2 int32
	{
		t0 := int32(load32(m.memory[uint32(v0):]))
		v1 = t0
		if v1 != 0 {
			v2 = i32(1)
			{
				if uint32(v1) < uint32(i32(51)) {
					goto l1
				}
				t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
				if uint32(t1) >= uint32(v1<<3+i32(-8)) {
					goto l1
				}
				v2 = i32(0)
				store32(m.memory[uint32(v0):], uint32(i32(0)))
			}
		l1:
			return v2
		}
		return i32(0)
	}
}
func (m *Module) fn1750(v0, v1, v2, v3, v4 int32) {
	var v5, v6 int32
	t0 := m.g0
	v5 = t0 - i32(16)
	m.g0 = v5
	t1 := int32(load32(m.memory[uint32(v2):]))
	m.t0[uint(t1)].(func(int32, int32, int32, int32))(v5+i32(8), v2, v3, v4)
	t2 := int32(load32(m.memory[int64(uint32(v5))+12:]))
	v2 = t2
	t3 := int32(load32(m.memory[int64(uint32(v5))+8:]))
	v3 = t3
	t4 := int32(load32(m.memory[uint32(v1):]))
	t5 := v1
	v6 = t4 + i32(1)
	p6 := i32(-1)
	if v6 != 0 {
		p6 = v6
	}
	store32(m.memory[uint32(t5):], uint32(p6))
	t7 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	t8 := v1
	v6 = t7
	t10 := v6
	p9 := v4
	if v3&i32(1) != 0 {
		p9 = v2
	}
	v4 = t10 + p9
	p11 := v4
	if uint32(v4) < uint32(v6) {
		p11 = i32(-1)
	}
	store32(m.memory[int64(uint32(t8))+4:], uint32(p11))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v2))
	store32(m.memory[uint32(v0):], uint32(v3))
	m.g0 = v5 + i32(16)
}
func (m *Module) fn1751(v0, v1, v2, v3, v4 int32) int32 {
	var v5, v6 int32
	t0 := m.g0
	v5 = t0 - i32(48)
	m.g0 = v5
	store32(m.memory[int64(uint32(v5))+4:], uint32(v3))
	{
		if uint32(v3) >= uint32(v2) {
			m.fn158(v3, v2, i32(1282728))
			panic("unreachable")
		}
		t1 := int32(m.memory[uint32(v1+v3)])
		v6 = t1
		switch v6 + i32(-10) {
		case 0:
			goto l1
		case 3:
			m.fn1752(v0, v4)
			v4 = v3 + i32(1)
			if uint32(v4) >= uint32(v2) {
				goto l4
			}
			t2 := int32(m.memory[uint32(v1+v4)])
			p3 := v4
			if t2 == i32(10) {
				p3 = v3 + i32(2)
			}
			v4 = p3
			goto l4
		default:
			m.memory[int64(uint32(v5))+11] = byte(v6)
			store32(m.memory[int64(uint32(v5))+12:], uint32(v6))
			store32(m.memory[int64(uint32(v5))+44:], uint32(i32(188)))
			store32(m.memory[int64(uint32(v5))+36:], uint32(i32(96)))
			store32(m.memory[int64(uint32(v5))+28:], uint32(i32(97)))
			store32(m.memory[int64(uint32(v5))+20:], uint32(i32(5)))
			store32(m.memory[int64(uint32(v5))+40:], uint32(v5+i32(11)))
			store32(m.memory[int64(uint32(v5))+32:], uint32(v5+i32(11)))
			store32(m.memory[int64(uint32(v5))+24:], uint32(v5+i32(12)))
			store32(m.memory[int64(uint32(v5))+16:], uint32(v5+i32(4)))
			m.fn91(i32(1068376), v5+i32(16), i32(1282744))
			panic("unreachable")
		}
	}
l1:
	m.fn1752(v0, v4)
	v4 = v3 + i32(1)
l4:
	m.g0 = v5 + i32(48)
	return v4
}
