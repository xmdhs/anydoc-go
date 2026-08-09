package core

import (
	"math/bits"
)

func (m *Module) fn987() {
	m.fn91(i32(1128728), i32(153), i32(1128804))
	panic("unreachable")
}
func (m *Module) fn988(v0, v1, v2, v3 int32) int32 {
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
		t4 := m.fn988(t0, t2, t3+v5, v3)
		v0 = t4
		t5 := m.fn988(v1, v1+v4, v1+v5, v3)
		v1 = t5
		t6 := m.fn988(v2, v2+v4, v2+v5, v3)
		v2 = t6
	}
l0:
	t7 := int32(load32(m.memory[int64(uint32(v0))+64:]))
	t8 := v0
	t9 := v2
	t10 := v1
	v3 = t7
	t11 := int32(load32(m.memory[int64(uint32(v1))+64:]))
	t12 := v3
	v4 = t11
	var p13 int32
	if uint32(t12) < uint32(v4) {
		p13 = 1
	}
	v5 = p13
	t14 := int32(load32(m.memory[int64(uint32(v2))+64:]))
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
func (m *Module) fn989(v0, v1, v2, v3, v4 int32) {
	var v5 int32
	var v6, v7 int64
	var v8, v9, v10, v11, v12, v13, v14, v15, v16, v17, v18, v19, v20, v21, v22, v23, v24, v25 int32
	t0 := m.g0
	v5 = t0 - i32(352)
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
		{
			if uint32(v1) < uint32(i32(4097)) {
				goto l0
			}
			t3 := fn980(v1)
			v8 = t3
			goto l1
		}
	l0:
		v9 = v1 - int32(uint32(v1)>>1)
		p4 := i32(64)
		if uint32(v9) < uint32(i32(64)) {
			p4 = v9
		}
		v8 = p4
	}
l1:
	v7 = v7 + v6
	v10 = v0 + i32(-8)
	v11 = v0 + i32(12)
	v9 = i32(1)
	v12 = i32(0)
	v13 = i32(0)
l18:
	{
		v14 = i32(0)
		v15 = i32(1)
		{
			var p5 int32
			if uint32(v1) > uint32(v12) {
				p5 = 1
			}
			v16 = p5
			if v16 == 0 {
				goto l2
			}
			t6 := v0
			v17 = v12 << 3
			v18 = t6 + v17
			{
				v19 = v1 - v12
				if uint32(v19) < uint32(v8) {
					goto l3
				}
				v20 = i32(0)
				if uint32(v19) < uint32(i32(2)) {
					goto l4
				}
				{
					t7 := int32(load32(m.memory[int64(uint32(v18))+12:]))
					t8 := int32(load32(m.memory[int64(uint32(v18))+4:]))
					if uint32(t7) < uint32(t8) {
						v17 = v11 + v17
						v21 = i32(2)
					l8:
						{
							v20 = i32(1)
							if v19 == v21 {
								goto l4
							}
							v22 = v17 + i32(8)
							t11 := int32(load32(m.memory[uint32(v22):]))
							t12 := int32(load32(m.memory[uint32(v17):]))
							if uint32(t11) >= uint32(t12) {
								goto l6
							}
							v21 = v21 + i32(1)
							v17 = v22
							goto l8
						}
					}
					v17 = v11 + v17
					v21 = i32(2)
				l7:
					{
						if v19 == v21 {
							goto l4
						}
						v22 = v17 + i32(8)
						t9 := int32(load32(m.memory[uint32(v22):]))
						t10 := int32(load32(m.memory[uint32(v17):]))
						if uint32(t9) < uint32(t10) {
							goto l6
						}
						v21 = v21 + i32(1)
						v17 = v22
						goto l7
					}
				}
			l4:
				v21 = v19
			l6:
				if uint32(v21) < uint32(v8) {
					goto l3
				}
				{
					if v20 == 0 {
						goto l9
					}
					t13 := v5 + i32(336)
					t14 := v18
					v20 = int32(uint32(v21) >> 1)
					m.fn990(t13, t14, v20, v20, i32(1301108))
					t15 := int32(load32(m.memory[int64(uint32(v5))+336:]))
					v17 = t15
					t16 := int32(load32(m.memory[int64(uint32(v5))+340:]))
					v14 = t16
					t17 := v5 + i32(336)
					t18 := v18 + v21<<3
					v22 = v20 << 3
					m.fn990(t17, t18-v22, v20, v20, i32(1301124))
					v19 = i32(0)
					v18 = i32(0) - v14
					t19 := int32(load32(m.memory[int64(uint32(v5))+336:]))
					v22 = t19 + v22 + i32(-8)
					t20 := int32(load32(m.memory[int64(uint32(v5))+340:]))
					t21 := v20 + i32(-1)
					v15 = t20
					var p22 int32
					if uint32(t21) < uint32(v15) {
						p22 = 1
					}
					v23 = p22
				l12:
					v24 = v20 + v19
					if v24 == 0 {
						goto l9
					}
					if v18 == v19 {
						m.fn158(v14, v14, i32(1301140))
						panic("unreachable")
					}
					{
						if v23 == 0 {
							m.fn158(v24+i32(-1), v15, i32(1301156))
							panic("unreachable")
						}
						t23 := int64(load64(m.memory[uint32(v17):]))
						v6 = t23
						t24 := int64(load64(m.memory[uint32(v22):]))
						store64(m.memory[uint32(v17):], uint64(t24))
						store64(m.memory[uint32(v22):], uint64(v6))
						v17 = v17 + i32(8)
						v22 = v22 + i32(-8)
						v19 = v19 + i32(-1)
						goto l12
					}
				}
			l9:
				v15 = v21<<1 | i32(1)
				goto l13
			l3:
				{
					if v4 != 0 {
						goto l14
					}
					p25 := v8
					if uint32(v19) < uint32(v8) {
						p25 = v19
					}
					v15 = p25 << 1
					goto l13
				}
			l14:
				t27 := v18
				p26 := i32(32)
				if uint32(v19) < uint32(i32(32)) {
					p26 = v19
				}
				v19 = p26
				m.fn991(t27, v19, v2, v3, i32(0), i32(0))
				v15 = v19<<1 | i32(1)
			}
		l13:
			v14 = int32(int64(bits.LeadingZeros64(uint64(v7*int64(uint32(int32(uint32(v15)>>1)+v12<<1)) ^ (int64(uint32(v12-int32(uint32(v9)>>1)))+int64(uint32(v12)))*v7))))
		}
	l2:
		t28 := v10
		v19 = v12 << 3
		v25 = t28 + v19
		v24 = v0 + v19
	l24:
		{
			if uint32(v13) < uint32(i32(2)) {
				goto l15
			}
			t29 := v5 + i32(270)
			v18 = v13 + i32(-1)
			t30 := int32(m.memory[uint32(t29+v18)])
			if uint32(t30) >= uint32(v14) {
				{
					t31 := int32(load32(m.memory[uint32(v5+i32(4)+v18<<2):]))
					v19 = t31
					v13 = int32(uint32(v19) >> 1)
					t32 := v13
					v17 = int32(uint32(v9) >> 1)
					v23 = t32 + v17
					if uint32(v23) > uint32(v3) {
						goto l19
					}
					if (v19|v9)&i32(1) == 0 {
						v9 = v23 << 1
						v13 = v18
						goto l24
					}
				}
			l19:
				v20 = v0 + (v12-v23)<<3
				if v19&i32(1) == 0 {
					goto l21
				}
				goto l22
			}
		}
	l15:
		m.memory[uint32(v5+i32(270)+v13)] = byte(v14)
		store32(m.memory[uint32(v5+i32(4)+v13<<2):], uint32(v9))
		if v16 == 0 {
			if v9&i32(1) != 0 {
				goto l23
			}
			m.fn992(v0, v1, v2, v3)
		l23:
			m.g0 = v5 + i32(352)
			return
		}
		v13 = v13 + i32(1)
		v12 = int32(uint32(v15)>>1) + v12
		v9 = v15
		goto l18
	l21:
		m.fn992(v20, v13, v2, v3)
	l22:
		if v9&i32(1) != 0 {
			goto l25
		}
		m.fn992(v20+v13<<3, v17, v2, v3)
	l25:
		{
			if v13 == 0 {
				goto l26
			}
			if v17 == 0 {
				goto l26
			}
			t33 := v3
			t34 := v17
			t35 := v13
			var p36 int32
			if uint32(v17) < uint32(v13) {
				p36 = 1
			}
			v19 = p36
			p37 := t35
			if v19 != 0 {
				p37 = t34
			}
			v22 = p37
			if uint32(t33) < uint32(v22) {
				goto l26
			}
			v17 = v20 + v13<<3
			p38 := v20
			if v19 != 0 {
				p38 = v17
			}
			v9 = p38
			v13 = v22 << 3
			if v13 == 0 {
				goto l27
			}
			memory_copy(m.memory, uint32(v2), uint32(v9), uint32(v13))
		l27:
			v13 = v2 + v13
			if v19 != 0 {
				goto l28
			}
			v19 = v2
		l30:
			{
				if v19 == v13 {
					goto l29
				}
				if v17 == v24 {
					goto l29
				}
				t39 := int32(load32(m.memory[int64(uint32(v17))+4:]))
				t40 := v9
				t41 := v17
				t42 := v19
				v22 = t39
				t43 := int32(load32(m.memory[int64(uint32(v19))+4:]))
				t44 := v22
				v21 = t43
				var p45 int32
				if uint32(t44) < uint32(v21) {
					p45 = 1
				}
				v20 = p45
				p46 := t42
				if v20 != 0 {
					p46 = t41
				}
				t47 := int64(load64(m.memory[uint32(p46):]))
				store64(m.memory[uint32(t40):], uint64(t47))
				v9 = v9 + i32(8)
				v17 = v17 + v20<<3
				t48 := v19
				var p49 int32
				if uint32(v22) >= uint32(v21) {
					p49 = 1
				}
				v19 = t48 + p49<<3
				goto l30
			}
		l28:
			v19 = v25
		l32:
			{
				t50 := v19
				v17 = v9 + i32(-8)
				t51 := v17
				v22 = v13 + i32(-8)
				t52 := int32(load32(m.memory[uint32(v13+i32(-4)):]))
				t53 := v22
				v21 = t52
				t54 := int32(load32(m.memory[uint32(v9+i32(-4)):]))
				t55 := v21
				v9 = t54
				var p56 int32
				if uint32(t55) < uint32(v9) {
					p56 = 1
				}
				v13 = p56
				p57 := t53
				if v13 != 0 {
					p57 = t51
				}
				t58 := int64(load64(m.memory[uint32(p57):]))
				store64(m.memory[uint32(t50):], uint64(t58))
				v13 = v22 + v13<<3
				t59 := v17
				var p60 int32
				if uint32(v21) >= uint32(v9) {
					p60 = 1
				}
				v9 = t59 + p60<<3
				if v9 == v20 {
					goto l31
				}
				v19 = v19 + i32(-8)
				if v13 != v2 {
					goto l32
				}
			}
		l31:
			v19 = v2
		l29:
			v13 = v13 - v19
			if v13 == 0 {
				goto l26
			}
			memory_copy(m.memory, uint32(v9), uint32(v19), uint32(v13))
		}
	l26:
		v9 = v23<<1 | i32(1)
		v13 = v18
		goto l24
	}
}
func (m *Module) fn990(v0, v1, v2, v3, v4 int32) {
	if uint32(v2) >= uint32(v3) {
		goto l0
	}
	m.fn91(i32(1301172), i32(19), v4)
	panic("unreachable")
l0:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
	store32(m.memory[uint32(v0):], uint32(v1))
	store32(m.memory[int64(uint32(v0))+12:], uint32(v2-v3))
	store32(m.memory[int64(uint32(v0))+8:], uint32(v1+v3<<3))
}
func (m *Module) fn991(v0, v1, v2, v3, v4, v5 int32) {
	var v6, v7, v8, v9, v10, v11, v12, v13, v14, v15, v16, v17, v18 int32
	t0 := m.g0
	v6 = t0 - i32(32)
	m.g0 = v6
	v7 = v2 + i32(-8)
	{
	l19:
		{
			if uint32(v1) < uint32(i32(33)) {
				if uint32(v1) < uint32(i32(2)) {
					goto l2
				}
				if uint32(v3) < uint32(v1+i32(16)) {
					goto l3
				}
				v8 = int32(uint32(v1) >> 1)
				if uint32(v1) > uint32(i32(15)) {
					t67 := v0
					t68 := v2
					v9 = v2 + v1<<3
					m.fn995(t67, t68, v9)
					t69 := v0
					v11 = v8 << 3
					m.fn995(t69+v11, v2+v11, v9+i32(64))
					v10 = i32(8)
					goto l6
				}
				{
					if uint32(v1) <= uint32(i32(7)) {
						t2 := int64(load64(m.memory[uint32(v0):]))
						store64(m.memory[uint32(v2):], uint64(t2))
						t3 := v2
						v9 = v8 << 3
						t4 := int64(load64(m.memory[uint32(v0+v9):]))
						store64(m.memory[uint32(t3+v9):], uint64(t4))
						v10 = i32(1)
						goto l6
					}
					m.fn993(v0, v2)
					t1 := v0
					v9 = v8 << 3
					m.fn993(t1+v9, v2+v9)
					v10 = i32(4)
					goto l6
				}
			}
			if v4 != 0 {
				t5 := v0
				v9 = int32(uint32(v1) >> 3)
				v11 = t5 + v9*i32(56)
				v12 = v0 + v9<<5
				{
					{
						if uint32(v1) < uint32(i32(64)) {
							goto l7
						}
						t6 := m.fn994(v0, v12, v11, v9)
						v10 = t6
						goto l8
					}
				l7:
					t7 := int32(load32(m.memory[int64(uint32(v0))+4:]))
					t8 := v0
					t9 := v11
					t10 := v12
					v9 = t7
					t11 := int32(load32(m.memory[int64(uint32(v12))+4:]))
					t12 := v9
					v13 = t11
					var p13 int32
					if uint32(t12) < uint32(v13) {
						p13 = 1
					}
					v10 = p13
					t14 := int32(load32(m.memory[int64(uint32(v11))+4:]))
					t15 := v10
					t16 := v13
					v8 = t14
					var p17 int32
					if uint32(t16) < uint32(v8) {
						p17 = 1
					}
					p18 := t10
					if t15^p17 != 0 {
						p18 = t9
					}
					t19 := v10
					var p20 int32
					if uint32(v9) < uint32(v8) {
						p20 = 1
					}
					p21 := p18
					if t19^p20 != 0 {
						p21 = t8
					}
					v10 = p21
				}
			l8:
				v4 = v4 + i32(-1)
				t22 := int32(load32(m.memory[uint32(v10):]))
				v9 = t22
				t23 := int32(load32(m.memory[int64(uint32(v10))+4:]))
				t24 := v6
				v11 = t23
				store32(m.memory[int64(uint32(t24))+12:], uint32(v11))
				store32(m.memory[int64(uint32(v6))+8:], uint32(v9))
				v14 = int32(uint32(v10-v0) >> 3)
				{
					if v5 == 0 {
						goto l9
					}
					t25 := int32(load32(m.memory[int64(uint32(v5))+4:]))
					if uint32(t25) >= uint32(v11) {
						goto l10
					}
				}
			l9:
				if uint32(v3) < uint32(v1) {
					goto l3
				}
				t26 := v2
				v15 = v1 << 3
				v11 = t26 + v15
				v12 = i32(0)
				v9 = v0
				v16 = v14
			l15:
				{
					t27 := v0
					v13 = v16 + i32(-3)
					p28 := v13
					if uint32(v13) > uint32(v16) {
						p28 = i32(0)
					}
					v17 = t27 + p28<<3
				l12:
					{
						if uint32(v9) >= uint32(v17) {
							v8 = v0 + v16<<3
						l20:
							if uint32(v9) < uint32(v8) {
								t60 := v2
								v11 = v11 + i32(-8)
								t61 := int32(load32(m.memory[uint32(v9+i32(4)):]))
								t62 := int32(load32(m.memory[int64(uint32(v10))+4:]))
								t63 := v11
								var p64 int32
								if uint32(t61) < uint32(t62) {
									p64 = 1
								}
								v13 = p64
								p65 := t63
								if v13 != 0 {
									p65 = t60
								}
								t66 := int64(load64(m.memory[uint32(v9):]))
								store64(m.memory[uint32(p65+v12<<3):], uint64(t66))
								v9 = v9 + i32(8)
								v12 = v12 + v13
								goto l20
							}
							{
								if v16 == v1 {
									v9 = v12 << 3
									if v9 == 0 {
										goto l16
									}
									memory_copy(m.memory, uint32(v0), uint32(v2), uint32(v9))
								l16:
									v11 = v7 + v15
									v13 = v0 + v9
									v9 = v12
								l18:
									{
										if v1 == v9 {
											if v12 == 0 {
												goto l10
											}
											m.fn990(v6+i32(16), v0, v1, v12, i32(1072736))
											t56 := int32(load32(m.memory[int64(uint32(v6))+20:]))
											v1 = t56
											t57 := int32(load32(m.memory[int64(uint32(v6))+16:]))
											v0 = t57
											t58 := int32(load32(m.memory[int64(uint32(v6))+24:]))
											t59 := int32(load32(m.memory[int64(uint32(v6))+28:]))
											m.fn991(t58, t59, v2, v3, v4, v6+i32(8))
											goto l19
										}
										t55 := int64(load64(m.memory[uint32(v11):]))
										store64(m.memory[uint32(v13):], uint64(t55))
										v9 = v9 + i32(1)
										v13 = v13 + i32(8)
										v11 = v11 + i32(-8)
										goto l18
									}
								}
								v11 = v11 + i32(-8)
								t54 := int64(load64(m.memory[uint32(v9):]))
								store64(m.memory[uint32(v11+v12<<3):], uint64(t54))
								v9 = v9 + i32(8)
								v16 = v1
								goto l15
							}
						}
						t29 := int32(load32(m.memory[uint32(v9+i32(4)):]))
						t30 := int32(load32(m.memory[int64(uint32(v10))+4:]))
						t31 := v2
						t32 := v11 + i32(-8)
						v13 = t30
						var p33 int32
						if uint32(t29) < uint32(v13) {
							p33 = 1
						}
						v8 = p33
						p34 := t32
						if v8 != 0 {
							p34 = t31
						}
						t35 := int64(load64(m.memory[uint32(v9):]))
						store64(m.memory[uint32(p34+v12<<3):], uint64(t35))
						t36 := int32(load32(m.memory[uint32(v9+i32(12)):]))
						t37 := v2
						t38 := v11 + i32(-16)
						var p39 int32
						if uint32(t36) < uint32(v13) {
							p39 = 1
						}
						v18 = p39
						p40 := t38
						if v18 != 0 {
							p40 = t37
						}
						v12 = v12 + v8
						t41 := int64(load64(m.memory[uint32(v9+i32(8)):]))
						store64(m.memory[uint32(p40+v12<<3):], uint64(t41))
						t42 := int32(load32(m.memory[uint32(v9+i32(20)):]))
						t43 := v2
						t44 := v11 + i32(-24)
						var p45 int32
						if uint32(t42) < uint32(v13) {
							p45 = 1
						}
						v8 = p45
						p46 := t44
						if v8 != 0 {
							p46 = t43
						}
						v12 = v12 + v18
						t47 := int64(load64(m.memory[uint32(v9+i32(16)):]))
						store64(m.memory[uint32(p46+v12<<3):], uint64(t47))
						t48 := v2
						v11 = v11 + i32(-32)
						t49 := int32(load32(m.memory[uint32(v9+i32(28)):]))
						t50 := v11
						var p51 int32
						if uint32(t49) < uint32(v13) {
							p51 = 1
						}
						v13 = p51
						p52 := t50
						if v13 != 0 {
							p52 = t48
						}
						v12 = v12 + v8
						t53 := int64(load64(m.memory[uint32(v9+i32(24)):]))
						store64(m.memory[uint32(p52+v12<<3):], uint64(t53))
						v12 = v12 + v13
						v9 = v9 + i32(32)
						goto l12
					}
				}
			}
			m.fn989(v0, v1, v2, v3, i32(1))
			goto l2
		l10:
			if uint32(v3) < uint32(v1) {
				goto l3
			}
			t70 := v2
			v16 = v1 << 3
			v11 = t70 + v16
			v12 = i32(0)
			v9 = v0
		l25:
			{
				t71 := v0
				v13 = v14 + i32(-3)
				p72 := v13
				if uint32(v13) > uint32(v14) {
					p72 = i32(0)
				}
				v17 = t71 + p72<<3
			l22:
				{
					if uint32(v9) >= uint32(v17) {
						v8 = v0 + v14<<3
					l30:
						if uint32(v9) < uint32(v8) {
							t100 := v2
							v11 = v11 + i32(-8)
							t101 := int32(load32(m.memory[int64(uint32(v10))+4:]))
							t102 := int32(load32(m.memory[uint32(v9+i32(4)):]))
							t103 := v11
							var p104 int32
							if uint32(t101) >= uint32(t102) {
								p104 = 1
							}
							v13 = p104
							p105 := t103
							if v13 != 0 {
								p105 = t100
							}
							t106 := int64(load64(m.memory[uint32(v9):]))
							store64(m.memory[uint32(p105+v12<<3):], uint64(t106))
							v9 = v9 + i32(8)
							v12 = v12 + v13
							goto l30
						}
						{
							if v14 == v1 {
								v13 = v12 << 3
								if v13 == 0 {
									goto l26
								}
								memory_copy(m.memory, uint32(v0), uint32(v2), uint32(v13))
							l26:
								v11 = v7 + v16
								v10 = v1 - v12
								v9 = v10
								v0 = v0 + v13
								v13 = v0
							l28:
								{
									if v9 == 0 {
										if uint32(v1) < uint32(v12) {
											m.fn151(v12, v1, v1, i32(1072752))
											panic("unreachable")
										}
										v5 = i32(0)
										v1 = v10
										goto l19
									}
									t99 := int64(load64(m.memory[uint32(v11):]))
									store64(m.memory[uint32(v13):], uint64(t99))
									v9 = v9 + i32(-1)
									v13 = v13 + i32(8)
									v11 = v11 + i32(-8)
									goto l28
								}
							}
							t98 := int64(load64(m.memory[uint32(v9):]))
							store64(m.memory[uint32(v2+v12<<3):], uint64(t98))
							v9 = v9 + i32(8)
							v12 = v12 + i32(1)
							v11 = v11 + i32(-8)
							v14 = v1
							goto l25
						}
					}
					t73 := int32(load32(m.memory[int64(uint32(v10))+4:]))
					t74 := v2
					t75 := v11 + i32(-8)
					v13 = t73
					t76 := int32(load32(m.memory[uint32(v9+i32(4)):]))
					var p77 int32
					if uint32(v13) >= uint32(t76) {
						p77 = 1
					}
					v8 = p77
					p78 := t75
					if v8 != 0 {
						p78 = t74
					}
					t79 := int64(load64(m.memory[uint32(v9):]))
					store64(m.memory[uint32(p78+v12<<3):], uint64(t79))
					t80 := int32(load32(m.memory[uint32(v9+i32(12)):]))
					t81 := v2
					t82 := v11 + i32(-16)
					var p83 int32
					if uint32(v13) >= uint32(t80) {
						p83 = 1
					}
					v18 = p83
					p84 := t82
					if v18 != 0 {
						p84 = t81
					}
					v12 = v12 + v8
					t85 := int64(load64(m.memory[uint32(v9+i32(8)):]))
					store64(m.memory[uint32(p84+v12<<3):], uint64(t85))
					t86 := int32(load32(m.memory[uint32(v9+i32(20)):]))
					t87 := v2
					t88 := v11 + i32(-24)
					var p89 int32
					if uint32(v13) >= uint32(t86) {
						p89 = 1
					}
					v8 = p89
					p90 := t88
					if v8 != 0 {
						p90 = t87
					}
					v12 = v12 + v18
					t91 := int64(load64(m.memory[uint32(v9+i32(16)):]))
					store64(m.memory[uint32(p90+v12<<3):], uint64(t91))
					t92 := v2
					v11 = v11 + i32(-32)
					t93 := int32(load32(m.memory[uint32(v9+i32(28)):]))
					t94 := v11
					var p95 int32
					if uint32(v13) >= uint32(t93) {
						p95 = 1
					}
					v13 = p95
					p96 := t94
					if v13 != 0 {
						p96 = t92
					}
					v12 = v12 + v8
					t97 := int64(load64(m.memory[uint32(v9+i32(24)):]))
					store64(m.memory[uint32(p96+v12<<3):], uint64(t97))
					v12 = v12 + v13
					v9 = v9 + i32(32)
					goto l22
				}
			}
		}
	l3:
		panic("unreachable")
	l6:
		store64(m.memory[int64(uint32(v6))+16:], uint64(i64(0x200000000)))
		store32(m.memory[int64(uint32(v6))+24:], uint32(i32(0)))
		v18 = i32(0) - v10
		t107 := v2
		v9 = v10 << 3
		v17 = t107 + v9
		v16 = v0 + v9
		store32(m.memory[int64(uint32(v6))+28:], uint32(v8))
		v14 = v1 - v8
	l32:
		{
			m.fn985(v6, v6+i32(16))
			t108 := int32(load32(m.memory[uint32(v6):]))
			if t108 != i32(1) {
				goto l31
			}
			t109 := int32(load32(m.memory[int64(uint32(v6))+4:]))
			t110 := v18
			t111 := v14
			t112 := v8
			v9 = t109
			p113 := t112
			if v9 != 0 {
				p113 = t111
			}
			v11 = p113
			p114 := v10
			if uint32(v11) > uint32(v10) {
				p114 = v11
			}
			v11 = t110 + p114
			t115 := v17
			v13 = v9 << 3
			v9 = t115 + v13
			v12 = v16 + v13
			v13 = v2 + v13
		l33:
			{
				if v11 == 0 {
					goto l32
				}
				t116 := int64(load64(m.memory[uint32(v12):]))
				store64(m.memory[uint32(v9):], uint64(t116))
				m.fn64(v13, v9)
				v11 = v11 + i32(-1)
				v9 = v9 + i32(8)
				v12 = v12 + i32(8)
				goto l33
			}
		}
	l31:
		m.fn996(v2, v1, v0)
	}
l2:
	m.g0 = v6 + i32(32)
}
func (m *Module) fn992(v0, v1, v2, v3 int32) {
	m.fn991(v0, v1, v2, v3, int32(bits.LeadingZeros32(uint32(v1|i32(1))))<<1^i32(62), i32(0))
}
func (m *Module) fn993(v0, v1 int32) {
	var v2, v3, v4, v5, v6, v7, v8, v9, v10 int32
	t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
	t1 := v0
	v2 = t0
	t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t3 := v2
	v3 = t2
	var p4 int32
	if uint32(t3) < uint32(v3) {
		p4 = 1
	}
	v4 = t1 + p4<<3
	t5 := int32(load32(m.memory[int64(uint32(v0))+28:]))
	t6 := int32(load32(m.memory[int64(uint32(v0))+20:]))
	t7 := v4
	t8 := v0
	var p9 int32
	if uint32(t5) < uint32(t6) {
		p9 = 1
	}
	v5 = p9
	p10 := i32(16)
	if v5 != 0 {
		p10 = i32(24)
	}
	v6 = t8 + p10
	t11 := v6
	t12 := v0
	var p13 int32
	if uint32(v2) >= uint32(v3) {
		p13 = 1
	}
	v2 = t12 + p13<<3
	t15 := v2
	t16 := v0
	p14 := i32(24)
	if v5 != 0 {
		p14 = i32(16)
	}
	v0 = t16 + p14
	t17 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t18 := int32(load32(m.memory[int64(uint32(v2))+4:]))
	var p19 int32
	if uint32(t17) < uint32(t18) {
		p19 = 1
	}
	v3 = p19
	p20 := t15
	if v3 != 0 {
		p20 = t11
	}
	t21 := int32(load32(m.memory[int64(uint32(v6))+4:]))
	t22 := int32(load32(m.memory[int64(uint32(v4))+4:]))
	var p23 int32
	if uint32(t21) < uint32(t22) {
		p23 = 1
	}
	v5 = p23
	p24 := p20
	if v5 != 0 {
		p24 = t7
	}
	v7 = p24
	t25 := int32(load32(m.memory[int64(uint32(v7))+4:]))
	v8 = t25
	t27 := v0
	p26 := v6
	if v5 != 0 {
		p26 = v2
	}
	p28 := p26
	if v3 != 0 {
		p28 = t27
	}
	v9 = p28
	t29 := int32(load32(m.memory[int64(uint32(v9))+4:]))
	v10 = t29
	t31 := v1
	p30 := v4
	if v5 != 0 {
		p30 = v6
	}
	t32 := int64(load64(m.memory[uint32(p30):]))
	store64(m.memory[uint32(t31):], uint64(t32))
	t33 := v1
	t34 := v9
	t35 := v7
	var p36 int32
	if uint32(v10) < uint32(v8) {
		p36 = 1
	}
	v6 = p36
	p37 := t35
	if v6 != 0 {
		p37 = t34
	}
	t38 := int64(load64(m.memory[uint32(p37):]))
	store64(m.memory[int64(uint32(t33))+8:], uint64(t38))
	t40 := v1
	p39 := v9
	if v6 != 0 {
		p39 = v7
	}
	t41 := int64(load64(m.memory[uint32(p39):]))
	store64(m.memory[int64(uint32(t40))+16:], uint64(t41))
	t43 := v1
	p42 := v0
	if v3 != 0 {
		p42 = v2
	}
	t44 := int64(load64(m.memory[uint32(p42):]))
	store64(m.memory[int64(uint32(t43))+24:], uint64(t44))
}
func (m *Module) fn994(v0, v1, v2, v3 int32) int32 {
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
		t4 := m.fn994(t0, t2, t3+v5, v3)
		v0 = t4
		t5 := m.fn994(v1, v1+v4, v1+v5, v3)
		v1 = t5
		t6 := m.fn994(v2, v2+v4, v2+v5, v3)
		v2 = t6
	}
l0:
	t7 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t8 := v0
	t9 := v2
	t10 := v1
	v3 = t7
	t11 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	t12 := v3
	v4 = t11
	var p13 int32
	if uint32(t12) < uint32(v4) {
		p13 = 1
	}
	v5 = p13
	t14 := int32(load32(m.memory[int64(uint32(v2))+4:]))
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
func (m *Module) fn995(v0, v1, v2 int32) {
	m.fn993(v0, v2)
	m.fn993(v0+i32(32), v2+i32(32))
	m.fn996(v2, i32(8), v1)
}
func (m *Module) fn996(v0, v1, v2 int32) {
	var v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13 int32
	t0 := v2
	v3 = v1<<3 + i32(-8)
	v4 = t0 + v3
	v5 = v0 + v3
	t1 := v0
	v6 = int32(uint32(v1) >> 1)
	v3 = t1 + v6<<3
	v7 = v3 + i32(-8)
l4:
	if v6 != 0 {
		t10 := int32(load32(m.memory[int64(uint32(v3))+4:]))
		t11 := v2
		t12 := v3
		t13 := v0
		v8 = t10
		t14 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		t15 := v8
		v9 = t14
		var p16 int32
		if uint32(t15) < uint32(v9) {
			p16 = 1
		}
		v10 = p16
		p17 := t13
		if v10 != 0 {
			p17 = t12
		}
		t18 := int64(load64(m.memory[uint32(p17):]))
		store64(m.memory[uint32(t11):], uint64(t18))
		t19 := int32(load32(m.memory[int64(uint32(v5))+4:]))
		t20 := v4
		t21 := v7
		t22 := v5
		v11 = t19
		t23 := int32(load32(m.memory[int64(uint32(v7))+4:]))
		t24 := v11
		v12 = t23
		var p25 int32
		if uint32(t24) < uint32(v12) {
			p25 = 1
		}
		v13 = p25
		p26 := t22
		if v13 != 0 {
			p26 = t21
		}
		t27 := int64(load64(m.memory[uint32(p26):]))
		store64(m.memory[uint32(t20):], uint64(t27))
		v6 = v6 + i32(-1)
		v4 = v4 + i32(-8)
		v2 = v2 + i32(8)
		t29 := v7
		p28 := i32(0)
		if v13 != 0 {
			p28 = i32(-8)
		}
		v7 = t29 + p28
		t31 := v5
		p30 := i32(0)
		if uint32(v11) >= uint32(v12) {
			p30 = i32(-8)
		}
		v5 = t31 + p30
		t32 := v0
		var p33 int32
		if uint32(v8) >= uint32(v9) {
			p33 = 1
		}
		v0 = t32 + p33<<3
		v3 = v3 + v10<<3
		goto l4
	}
	v7 = v7 + i32(8)
	{
		if v1&i32(1) == 0 {
			goto l1
		}
		t2 := v2
		t3 := v0
		t4 := v3
		var p5 int32
		if uint32(v0) < uint32(v7) {
			p5 = 1
		}
		v6 = p5
		p6 := t4
		if v6 != 0 {
			p6 = t3
		}
		t7 := int64(load64(m.memory[uint32(p6):]))
		store64(m.memory[uint32(t2):], uint64(t7))
		t8 := v3
		var p9 int32
		if uint32(v0) >= uint32(v7) {
			p9 = 1
		}
		v3 = t8 + p9<<3
		v0 = v0 + v6<<3
	}
l1:
	if v0 != v7 {
		goto l2
	}
	if v3 == v5+i32(8) {
		return
	}
l2:
	m.fn987()
	panic("unreachable")
}
func (m *Module) fn997(v0, v1 int32) {
	var v2, v3, v4 int32
	t0 := m.g0
	v2 = t0 - i32(4112)
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
			m.fn59(v2+i32(8), v3, i32(4), i32(12))
			t4 := int32(load32(m.memory[int64(uint32(v2))+12:]))
			t5 := v0
			t6 := v1
			v3 = t4
			t7 := int32(load32(m.memory[int64(uint32(v2))+8:]))
			t8 := v3
			v4 = t7
			t9 := v4
			var p10 int32
			if uint32(v1) < uint32(i32(65)) {
				p10 = 1
			}
			m.fn998(t5, t6, t8, t9, p10)
			m.fn911(v4, v3)
			goto l1
		}
	l0:
		t11 := v0
		t12 := v1
		t13 := v2 + i32(16)
		var p14 int32
		if uint32(v1) < uint32(i32(65)) {
			p14 = 1
		}
		m.fn998(t11, t12, t13, i32(341), p14)
	}
l1:
	m.g0 = v2 + i32(4112)
}
func (m *Module) fn998(v0, v1, v2, v3, v4 int32) {
	var v5 int32
	var v6, v7 int64
	var v8, v9, v10, v11, v12, v13, v14, v15, v16, v17, v18, v19, v20, v21, v22, v23, v24, v25, v26 int32
	t0 := m.g0
	v5 = t0 - i32(352)
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
		{
			if uint32(v1) < uint32(i32(4097)) {
				goto l0
			}
			t3 := fn980(v1)
			v8 = t3
			goto l1
		}
	l0:
		v9 = v1 - int32(uint32(v1)>>1)
		p4 := i32(64)
		if uint32(v9) < uint32(i32(64)) {
			p4 = v9
		}
		v8 = p4
	}
l1:
	v6 = v7 + v6
	v10 = v0 + i32(-12)
	v11 = v0 + i32(12)
	v9 = i32(1)
	v12 = i32(0)
	v13 = i32(0)
l18:
	{
		v14 = i32(0)
		v15 = i32(1)
		{
			var p5 int32
			if uint32(v1) > uint32(v12) {
				p5 = 1
			}
			v16 = p5
			if v16 == 0 {
				goto l2
			}
			t6 := v0
			v17 = v12 * i32(12)
			v18 = t6 + v17
			{
				v19 = v1 - v12
				if uint32(v19) < uint32(v8) {
					goto l3
				}
				v20 = i32(0)
				if uint32(v19) < uint32(i32(2)) {
					goto l4
				}
				{
					t7 := int32(load32(m.memory[int64(uint32(v18))+12:]))
					t8 := int32(load32(m.memory[uint32(v18):]))
					if uint32(t7) < uint32(t8) {
						v21 = v11 + v17
						v17 = i32(2)
					l8:
						{
							v20 = i32(1)
							if v19 == v17 {
								goto l4
							}
							v22 = v21 + i32(12)
							t11 := int32(load32(m.memory[uint32(v22):]))
							t12 := int32(load32(m.memory[uint32(v21):]))
							if uint32(t11) >= uint32(t12) {
								goto l6
							}
							v17 = v17 + i32(1)
							v21 = v22
							goto l8
						}
					}
					v21 = v11 + v17
					v17 = i32(2)
				l7:
					{
						if v19 == v17 {
							goto l4
						}
						v22 = v21 + i32(12)
						t9 := int32(load32(m.memory[uint32(v22):]))
						t10 := int32(load32(m.memory[uint32(v21):]))
						if uint32(t9) < uint32(t10) {
							goto l6
						}
						v17 = v17 + i32(1)
						v21 = v22
						goto l7
					}
				}
			l4:
				v17 = v19
			l6:
				if uint32(v17) < uint32(v8) {
					goto l3
				}
				{
					if v20 == 0 {
						goto l9
					}
					t13 := v5 + i32(336)
					t14 := v18
					v19 = int32(uint32(v17) >> 1)
					m.fn999(t13, t14, v19, v19, i32(1301108))
					t15 := int32(load32(m.memory[int64(uint32(v5))+340:]))
					v23 = t15
					t16 := int32(load32(m.memory[int64(uint32(v5))+336:]))
					v22 = t16
					m.fn999(v5+i32(336), v18+v17*i32(12)+(i32(0)-v19)*i32(12), v19, v19, i32(1301124))
					t17 := int32(load32(m.memory[int64(uint32(v5))+336:]))
					v20 = t17 + v19*i32(12) + i32(-12)
					t18 := int32(load32(m.memory[int64(uint32(v5))+340:]))
					v18 = t18
					v21 = v23
					v24 = v19 + i32(-1)
					v19 = v24
				l12:
					if v19 == i32(-1) {
						goto l9
					}
					if v21 == 0 {
						m.fn158(v23, v23, i32(1301140))
						panic("unreachable")
					}
					if uint32(v24) >= uint32(v18) {
						m.fn158(v19, v18, i32(1301156))
						panic("unreachable")
					}
					m.fn244(v22, v20, i32(3))
					v21 = v21 + i32(-1)
					v22 = v22 + i32(12)
					v20 = v20 + i32(-12)
					v19 = v19 + i32(-1)
					goto l12
				}
			l9:
				v15 = v17<<1 | i32(1)
				goto l13
			l3:
				{
					if v4 != 0 {
						goto l14
					}
					p19 := v8
					if uint32(v19) < uint32(v8) {
						p19 = v19
					}
					v15 = p19 << 1
					goto l13
				}
			l14:
				t21 := v18
				p20 := i32(32)
				if uint32(v19) < uint32(i32(32)) {
					p20 = v19
				}
				v17 = p20
				m.fn1000(t21, v17, v2, v3, i32(0), i32(0))
				v15 = v17<<1 | i32(1)
			}
		l13:
			v14 = int32(int64(bits.LeadingZeros64(uint64(v6*int64(uint32(int32(uint32(v15)>>1)+v12<<1)) ^ (int64(uint32(v12-int32(uint32(v9)>>1)))+int64(uint32(v12)))*v6))))
		}
	l2:
		t22 := v10
		v17 = v12 * i32(12)
		v25 = t22 + v17
		v26 = v0 + v17
	l24:
		{
			if uint32(v13) < uint32(i32(2)) {
				goto l15
			}
			t23 := v5 + i32(270)
			v24 = v13 + i32(-1)
			t24 := int32(m.memory[uint32(t23+v24)])
			if uint32(t24) >= uint32(v14) {
				{
					t25 := int32(load32(m.memory[uint32(v5+i32(4)+v24<<2):]))
					v17 = t25
					v13 = int32(uint32(v17) >> 1)
					t26 := v13
					v19 = int32(uint32(v9) >> 1)
					v23 = t26 + v19
					if uint32(v23) > uint32(v3) {
						goto l19
					}
					if (v17|v9)&i32(1) == 0 {
						v9 = v23 << 1
						v13 = v24
						goto l24
					}
				}
			l19:
				v18 = v0 + (v12-v23)*i32(12)
				if v17&i32(1) == 0 {
					goto l21
				}
				goto l22
			}
		}
	l15:
		m.memory[uint32(v5+i32(270)+v13)] = byte(v14)
		store32(m.memory[uint32(v5+i32(4)+v13<<2):], uint32(v9))
		if v16 == 0 {
			if v9&i32(1) != 0 {
				goto l23
			}
			m.fn1001(v0, v1, v2, v3)
		l23:
			m.g0 = v5 + i32(352)
			return
		}
		v13 = v13 + i32(1)
		v12 = int32(uint32(v15)>>1) + v12
		v9 = v15
		goto l18
	l21:
		m.fn1001(v18, v13, v2, v3)
	l22:
		if v9&i32(1) != 0 {
			goto l25
		}
		m.fn1001(v18+v13*i32(12), v19, v2, v3)
	l25:
		{
			if v13 == 0 {
				goto l26
			}
			if v19 == 0 {
				goto l26
			}
			t27 := v3
			t28 := v19
			t29 := v13
			var p30 int32
			if uint32(v19) < uint32(v13) {
				p30 = 1
			}
			v21 = p30
			p31 := t29
			if v21 != 0 {
				p31 = t28
			}
			v19 = p31
			if uint32(t27) < uint32(v19) {
				goto l26
			}
			v17 = v18 + v13*i32(12)
			p32 := v18
			if v21 != 0 {
				p32 = v17
			}
			v9 = p32
			v13 = v19 * i32(12)
			if v13 == 0 {
				goto l27
			}
			memory_copy(m.memory, uint32(v2), uint32(v9), uint32(v13))
		l27:
			v19 = v2 + v13
			if v21 != 0 {
				goto l28
			}
			v13 = v2
		l30:
			{
				if v13 == v19 {
					goto l29
				}
				if v17 == v26 {
					goto l29
				}
				t33 := int32(load32(m.memory[uint32(v17):]))
				t34 := v9
				t35 := v17
				t36 := v13
				v21 = t33
				t37 := int32(load32(m.memory[uint32(v13):]))
				t38 := v21
				v22 = t37
				var p39 int32
				if uint32(t38) < uint32(v22) {
					p39 = 1
				}
				v20 = p39
				p40 := t36
				if v20 != 0 {
					p40 = t35
				}
				v18 = p40
				t41 := int64(load64(m.memory[uint32(v18):]))
				store64(m.memory[uint32(t34):], uint64(t41))
				t42 := int32(load32(m.memory[int64(uint32(v18))+8:]))
				store32(m.memory[int64(uint32(v9))+8:], uint32(t42))
				v9 = v9 + i32(12)
				v17 = v17 + v20*i32(12)
				t43 := v13
				var p44 int32
				if uint32(v21) >= uint32(v22) {
					p44 = 1
				}
				v13 = t43 + p44*i32(12)
				goto l30
			}
		l28:
			v13 = v25
		l32:
			{
				t45 := v13
				v9 = v9 + i32(-12)
				t46 := v9
				v17 = v19 + i32(-12)
				t47 := int32(load32(m.memory[uint32(v17):]))
				t48 := v17
				v21 = t47
				t49 := int32(load32(m.memory[uint32(v9):]))
				t50 := v21
				v22 = t49
				var p51 int32
				if uint32(t50) < uint32(v22) {
					p51 = 1
				}
				v19 = p51
				p52 := t48
				if v19 != 0 {
					p52 = t46
				}
				v20 = p52
				t53 := int64(load64(m.memory[uint32(v20):]))
				store64(m.memory[uint32(t45):], uint64(t53))
				t54 := int32(load32(m.memory[int64(uint32(v20))+8:]))
				store32(m.memory[int64(uint32(v13))+8:], uint32(t54))
				v19 = v17 + v19*i32(12)
				t55 := v9
				var p56 int32
				if uint32(v21) >= uint32(v22) {
					p56 = 1
				}
				v9 = t55 + p56*i32(12)
				if v9 == v18 {
					goto l31
				}
				v13 = v13 + i32(-12)
				if v19 != v2 {
					goto l32
				}
			}
		l31:
			v13 = v2
		l29:
			v17 = v19 - v13
			if v17 == 0 {
				goto l26
			}
			memory_copy(m.memory, uint32(v9), uint32(v13), uint32(v17))
		}
	l26:
		v9 = v23<<1 | i32(1)
		v13 = v24
		goto l24
	}
}
func (m *Module) fn999(v0, v1, v2, v3, v4 int32) {
	if uint32(v2) >= uint32(v3) {
		goto l0
	}
	m.fn91(i32(1301172), i32(19), v4)
	panic("unreachable")
l0:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
	store32(m.memory[uint32(v0):], uint32(v1))
	store32(m.memory[int64(uint32(v0))+12:], uint32(v2-v3))
	store32(m.memory[int64(uint32(v0))+8:], uint32(v1+v3*i32(12)))
}
func (m *Module) fn1000(v0, v1, v2, v3, v4, v5 int32) {
	var v6, v7, v8, v9, v10, v11, v12, v13, v14, v15, v16, v17, v18 int32
	t0 := m.g0
	v6 = t0 - i32(48)
	m.g0 = v6
	v7 = v2 + i32(-12)
	{
	l19:
		{
			if uint32(v1) < uint32(i32(33)) {
				if uint32(v1) < uint32(i32(2)) {
					goto l2
				}
				if uint32(v3) < uint32(v1+i32(16)) {
					goto l3
				}
				v8 = int32(uint32(v1) >> 1)
				if uint32(v1) > uint32(i32(15)) {
					t79 := v0
					t80 := v2
					v10 = v2 + v1*i32(12)
					m.fn1004(t79, t80, v10)
					t81 := v0
					v9 = v8 * i32(12)
					m.fn1004(t81+v9, v2+v9, v10+i32(96))
					v11 = i32(8)
					goto l6
				}
				t1 := v2
				v9 = v8 * i32(12)
				v10 = t1 + v9
				v9 = v0 + v9
				if uint32(v1) <= uint32(i32(7)) {
					t2 := int32(load32(m.memory[int64(uint32(v0))+8:]))
					store32(m.memory[int64(uint32(v2))+8:], uint32(t2))
					t3 := int64(load64(m.memory[uint32(v0):]))
					store64(m.memory[uint32(v2):], uint64(t3))
					t4 := int32(load32(m.memory[int64(uint32(v9))+8:]))
					store32(m.memory[int64(uint32(v10))+8:], uint32(t4))
					t5 := int64(load64(m.memory[uint32(v9):]))
					store64(m.memory[uint32(v10):], uint64(t5))
					v11 = i32(1)
					goto l6
				}
				m.fn1002(v0, v2)
				m.fn1002(v9, v10)
				v11 = i32(4)
				goto l6
			}
			if v4 != 0 {
				t6 := v0
				v10 = int32(uint32(v1) >> 3)
				v9 = t6 + v10*i32(84)
				v12 = v0 + v10*i32(48)
				{
					{
						if uint32(v1) < uint32(i32(64)) {
							goto l7
						}
						t7 := m.fn1003(v0, v12, v9, v10)
						v13 = t7
						goto l8
					}
				l7:
					t8 := int32(load32(m.memory[uint32(v0):]))
					t9 := v0
					t10 := v9
					t11 := v12
					v10 = t8
					t12 := int32(load32(m.memory[uint32(v12):]))
					t13 := v10
					v13 = t12
					var p14 int32
					if uint32(t13) < uint32(v13) {
						p14 = 1
					}
					v11 = p14
					t15 := int32(load32(m.memory[uint32(v9):]))
					t16 := v11
					t17 := v13
					v8 = t15
					var p18 int32
					if uint32(t17) < uint32(v8) {
						p18 = 1
					}
					p19 := t11
					if t16^p18 != 0 {
						p19 = t10
					}
					t20 := v11
					var p21 int32
					if uint32(v10) < uint32(v8) {
						p21 = 1
					}
					p22 := p19
					if t20^p21 != 0 {
						p22 = t9
					}
					v13 = p22
				}
			l8:
				v4 = v4 + i32(-1)
				t23 := int32(load32(m.memory[int64(uint32(v13))+8:]))
				store32(m.memory[int64(uint32(v6))+24:], uint32(t23))
				t24 := int64(load64(m.memory[uint32(v13):]))
				store64(m.memory[int64(uint32(v6))+16:], uint64(t24))
				t25 := int32(uint32(v13-v0) / uint32(i32(12)))
				v14 = t25
				{
					if v5 == 0 {
						goto l9
					}
					t26 := int32(load32(m.memory[uint32(v5):]))
					t27 := int32(load32(m.memory[uint32(v13):]))
					if uint32(t26) >= uint32(t27) {
						goto l10
					}
				}
			l9:
				if uint32(v3) < uint32(v1) {
					goto l3
				}
				t28 := v2
				v15 = v1 * i32(12)
				v9 = t28 + v15
				v12 = i32(0)
				v10 = v0
				v16 = v14
			l15:
				{
					t29 := v0
					v11 = v16 + i32(-3)
					p30 := v11
					if uint32(v11) > uint32(v16) {
						p30 = i32(0)
					}
					v17 = t29 + p30*i32(12)
				l12:
					{
						if uint32(v10) >= uint32(v17) {
							v18 = v0 + v16*i32(12)
						l20:
							if uint32(v10) < uint32(v18) {
								t71 := v2
								v9 = v9 + i32(-12)
								t72 := int32(load32(m.memory[uint32(v10):]))
								t73 := int32(load32(m.memory[uint32(v13):]))
								t74 := v9
								var p75 int32
								if uint32(t72) < uint32(t73) {
									p75 = 1
								}
								v11 = p75
								p76 := t74
								if v11 != 0 {
									p76 = t71
								}
								v8 = p76 + v12*i32(12)
								t77 := int32(load32(m.memory[int64(uint32(v10))+8:]))
								store32(m.memory[int64(uint32(v8))+8:], uint32(t77))
								t78 := int64(load64(m.memory[uint32(v10):]))
								store64(m.memory[uint32(v8):], uint64(t78))
								v10 = v10 + i32(12)
								v12 = v12 + v11
								goto l20
							}
							{
								if v16 == v1 {
									v9 = v12 * i32(12)
									if v9 == 0 {
										goto l16
									}
									memory_copy(m.memory, uint32(v0), uint32(v2), uint32(v9))
								l16:
									v10 = v7 + v15
									v9 = v0 + v9
									v11 = v12
								l18:
									{
										if v1 == v11 {
											if v12 == 0 {
												goto l10
											}
											m.fn999(v6+i32(32), v0, v1, v12, i32(1072736))
											t67 := int32(load32(m.memory[int64(uint32(v6))+36:]))
											v1 = t67
											t68 := int32(load32(m.memory[int64(uint32(v6))+32:]))
											v0 = t68
											t69 := int32(load32(m.memory[int64(uint32(v6))+40:]))
											t70 := int32(load32(m.memory[int64(uint32(v6))+44:]))
											m.fn1000(t69, t70, v2, v3, v4, v6+i32(16))
											goto l19
										}
										t65 := int32(load32(m.memory[int64(uint32(v10))+8:]))
										store32(m.memory[int64(uint32(v9))+8:], uint32(t65))
										t66 := int64(load64(m.memory[uint32(v10):]))
										store64(m.memory[uint32(v9):], uint64(t66))
										v11 = v11 + i32(1)
										v9 = v9 + i32(12)
										v10 = v10 + i32(-12)
										goto l18
									}
								}
								v9 = v9 + i32(-12)
								v11 = v9 + v12*i32(12)
								t63 := int32(load32(m.memory[int64(uint32(v10))+8:]))
								store32(m.memory[int64(uint32(v11))+8:], uint32(t63))
								t64 := int64(load64(m.memory[uint32(v10):]))
								store64(m.memory[uint32(v11):], uint64(t64))
								v10 = v10 + i32(12)
								v16 = v1
								goto l15
							}
						}
						t31 := int32(load32(m.memory[uint32(v10):]))
						t32 := int32(load32(m.memory[uint32(v13):]))
						t33 := v2
						t34 := v9 + i32(-12)
						var p35 int32
						if uint32(t31) < uint32(t32) {
							p35 = 1
						}
						v11 = p35
						p36 := t34
						if v11 != 0 {
							p36 = t33
						}
						v8 = p36 + v12*i32(12)
						t37 := int32(load32(m.memory[int64(uint32(v10))+8:]))
						store32(m.memory[int64(uint32(v8))+8:], uint32(t37))
						t38 := int64(load64(m.memory[uint32(v10):]))
						store64(m.memory[uint32(v8):], uint64(t38))
						t39 := v2
						t40 := v9 + i32(-24)
						v8 = v10 + i32(12)
						t41 := int32(load32(m.memory[uint32(v8):]))
						t42 := int32(load32(m.memory[uint32(v13):]))
						var p43 int32
						if uint32(t41) < uint32(t42) {
							p43 = 1
						}
						v18 = p43
						p44 := t40
						if v18 != 0 {
							p44 = t39
						}
						v12 = v12 + v11
						v11 = p44 + v12*i32(12)
						t45 := int32(load32(m.memory[uint32(v10+i32(20)):]))
						store32(m.memory[int64(uint32(v11))+8:], uint32(t45))
						t46 := int64(load64(m.memory[uint32(v8):]))
						store64(m.memory[uint32(v11):], uint64(t46))
						t47 := v2
						t48 := v9 + i32(-36)
						v11 = v10 + i32(24)
						t49 := int32(load32(m.memory[uint32(v11):]))
						t50 := int32(load32(m.memory[uint32(v13):]))
						var p51 int32
						if uint32(t49) < uint32(t50) {
							p51 = 1
						}
						v8 = p51
						p52 := t48
						if v8 != 0 {
							p52 = t47
						}
						v12 = v12 + v18
						v18 = p52 + v12*i32(12)
						t53 := int32(load32(m.memory[uint32(v10+i32(32)):]))
						store32(m.memory[int64(uint32(v18))+8:], uint32(t53))
						t54 := int64(load64(m.memory[uint32(v11):]))
						store64(m.memory[uint32(v18):], uint64(t54))
						t55 := v2
						v9 = v9 + i32(-48)
						t56 := v9
						v11 = v10 + i32(36)
						t57 := int32(load32(m.memory[uint32(v11):]))
						t58 := int32(load32(m.memory[uint32(v13):]))
						var p59 int32
						if uint32(t57) < uint32(t58) {
							p59 = 1
						}
						v18 = p59
						p60 := t56
						if v18 != 0 {
							p60 = t55
						}
						v12 = v12 + v8
						v8 = p60 + v12*i32(12)
						t61 := int32(load32(m.memory[uint32(v10+i32(44)):]))
						store32(m.memory[int64(uint32(v8))+8:], uint32(t61))
						t62 := int64(load64(m.memory[uint32(v11):]))
						store64(m.memory[uint32(v8):], uint64(t62))
						v12 = v12 + v18
						v10 = v10 + i32(48)
						goto l12
					}
				}
			}
			m.fn998(v0, v1, v2, v3, i32(1))
			goto l2
		l10:
			if uint32(v3) < uint32(v1) {
				goto l3
			}
			t82 := v2
			v16 = v1 * i32(12)
			v9 = t82 + v16
			v12 = i32(0)
			v10 = v0
		l25:
			{
				t83 := v0
				v11 = v14 + i32(-3)
				p84 := v11
				if uint32(v11) > uint32(v14) {
					p84 = i32(0)
				}
				v17 = t83 + p84*i32(12)
			l22:
				{
					if uint32(v10) >= uint32(v17) {
						v18 = v0 + v14*i32(12)
					l30:
						if uint32(v10) < uint32(v18) {
							t121 := v2
							v9 = v9 + i32(-12)
							t122 := int32(load32(m.memory[uint32(v13):]))
							t123 := int32(load32(m.memory[uint32(v10):]))
							t124 := v9
							var p125 int32
							if uint32(t122) >= uint32(t123) {
								p125 = 1
							}
							v11 = p125
							p126 := t124
							if v11 != 0 {
								p126 = t121
							}
							v8 = p126 + v12*i32(12)
							t127 := int32(load32(m.memory[int64(uint32(v10))+8:]))
							store32(m.memory[int64(uint32(v8))+8:], uint32(t127))
							t128 := int64(load64(m.memory[uint32(v10):]))
							store64(m.memory[uint32(v8):], uint64(t128))
							v10 = v10 + i32(12)
							v12 = v12 + v11
							goto l30
						}
						{
							if v14 == v1 {
								v9 = v12 * i32(12)
								if v9 == 0 {
									goto l26
								}
								memory_copy(m.memory, uint32(v0), uint32(v2), uint32(v9))
							l26:
								v10 = v7 + v16
								v11 = v1 - v12
								v13 = v11
								v0 = v0 + v9
								v9 = v0
							l28:
								{
									if v13 == 0 {
										if uint32(v1) < uint32(v12) {
											m.fn151(v12, v1, v1, i32(1072752))
											panic("unreachable")
										}
										v5 = i32(0)
										v1 = v11
										goto l19
									}
									t119 := int32(load32(m.memory[int64(uint32(v10))+8:]))
									store32(m.memory[int64(uint32(v9))+8:], uint32(t119))
									t120 := int64(load64(m.memory[uint32(v10):]))
									store64(m.memory[uint32(v9):], uint64(t120))
									v13 = v13 + i32(-1)
									v9 = v9 + i32(12)
									v10 = v10 + i32(-12)
									goto l28
								}
							}
							v11 = v2 + v12*i32(12)
							t117 := int32(load32(m.memory[int64(uint32(v10))+8:]))
							store32(m.memory[int64(uint32(v11))+8:], uint32(t117))
							t118 := int64(load64(m.memory[uint32(v10):]))
							store64(m.memory[uint32(v11):], uint64(t118))
							v10 = v10 + i32(12)
							v12 = v12 + i32(1)
							v9 = v9 + i32(-12)
							v14 = v1
							goto l25
						}
					}
					t85 := int32(load32(m.memory[uint32(v13):]))
					t86 := int32(load32(m.memory[uint32(v10):]))
					t87 := v2
					t88 := v9 + i32(-12)
					var p89 int32
					if uint32(t85) >= uint32(t86) {
						p89 = 1
					}
					v11 = p89
					p90 := t88
					if v11 != 0 {
						p90 = t87
					}
					v8 = p90 + v12*i32(12)
					t91 := int32(load32(m.memory[int64(uint32(v10))+8:]))
					store32(m.memory[int64(uint32(v8))+8:], uint32(t91))
					t92 := int64(load64(m.memory[uint32(v10):]))
					store64(m.memory[uint32(v8):], uint64(t92))
					t93 := int32(load32(m.memory[uint32(v13):]))
					t94 := v2
					t95 := v9 + i32(-24)
					v8 = v10 + i32(12)
					t96 := int32(load32(m.memory[uint32(v8):]))
					var p97 int32
					if uint32(t93) >= uint32(t96) {
						p97 = 1
					}
					v18 = p97
					p98 := t95
					if v18 != 0 {
						p98 = t94
					}
					v12 = v12 + v11
					v11 = p98 + v12*i32(12)
					t99 := int32(load32(m.memory[uint32(v10+i32(20)):]))
					store32(m.memory[int64(uint32(v11))+8:], uint32(t99))
					t100 := int64(load64(m.memory[uint32(v8):]))
					store64(m.memory[uint32(v11):], uint64(t100))
					t101 := int32(load32(m.memory[uint32(v13):]))
					t102 := v2
					t103 := v9 + i32(-36)
					v11 = v10 + i32(24)
					t104 := int32(load32(m.memory[uint32(v11):]))
					var p105 int32
					if uint32(t101) >= uint32(t104) {
						p105 = 1
					}
					v8 = p105
					p106 := t103
					if v8 != 0 {
						p106 = t102
					}
					v12 = v12 + v18
					v18 = p106 + v12*i32(12)
					t107 := int32(load32(m.memory[uint32(v10+i32(32)):]))
					store32(m.memory[int64(uint32(v18))+8:], uint32(t107))
					t108 := int64(load64(m.memory[uint32(v11):]))
					store64(m.memory[uint32(v18):], uint64(t108))
					t109 := v2
					v9 = v9 + i32(-48)
					t110 := int32(load32(m.memory[uint32(v13):]))
					t111 := v9
					v11 = v10 + i32(36)
					t112 := int32(load32(m.memory[uint32(v11):]))
					var p113 int32
					if uint32(t110) >= uint32(t112) {
						p113 = 1
					}
					v18 = p113
					p114 := t111
					if v18 != 0 {
						p114 = t109
					}
					v12 = v12 + v8
					v8 = p114 + v12*i32(12)
					t115 := int32(load32(m.memory[uint32(v10+i32(44)):]))
					store32(m.memory[int64(uint32(v8))+8:], uint32(t115))
					t116 := int64(load64(m.memory[uint32(v11):]))
					store64(m.memory[uint32(v8):], uint64(t116))
					v12 = v12 + v18
					v10 = v10 + i32(48)
					goto l22
				}
			}
		}
	l3:
		panic("unreachable")
	l6:
		store64(m.memory[int64(uint32(v6))+32:], uint64(i64(0x200000000)))
		store32(m.memory[int64(uint32(v6))+40:], uint32(i32(0)))
		v18 = i32(0) - v11
		t129 := v2
		v10 = v11 * i32(12)
		v17 = t129 + v10
		v16 = v0 + v10
		store32(m.memory[int64(uint32(v6))+44:], uint32(v8))
		v14 = v1 - v8
	l32:
		{
			m.fn985(v6+i32(8), v6+i32(32))
			t130 := int32(load32(m.memory[int64(uint32(v6))+8:]))
			if t130 != i32(1) {
				goto l31
			}
			t131 := int32(load32(m.memory[int64(uint32(v6))+12:]))
			t132 := v18
			t133 := v14
			t134 := v8
			v10 = t131
			p135 := t134
			if v10 != 0 {
				p135 = t133
			}
			v9 = p135
			p136 := v11
			if uint32(v9) > uint32(v11) {
				p136 = v9
			}
			v12 = t132 + p136
			t137 := v17
			v13 = v10 * i32(12)
			v10 = t137 + v13
			v9 = v16 + v13
			v13 = v2 + v13
		l33:
			{
				if v12 == 0 {
					goto l32
				}
				t138 := int32(load32(m.memory[int64(uint32(v9))+8:]))
				store32(m.memory[int64(uint32(v10))+8:], uint32(t138))
				t139 := int64(load64(m.memory[uint32(v9):]))
				store64(m.memory[uint32(v10):], uint64(t139))
				m.fn1005(v13, v10)
				v12 = v12 + i32(-1)
				v10 = v10 + i32(12)
				v9 = v9 + i32(12)
				goto l33
			}
		}
	l31:
		m.fn1006(v2, v1, v0)
	}
l2:
	m.g0 = v6 + i32(48)
}
func (m *Module) fn1001(v0, v1, v2, v3 int32) {
	m.fn1000(v0, v1, v2, v3, int32(bits.LeadingZeros32(uint32(v1|i32(1))))<<1^i32(62), i32(0))
}
func (m *Module) fn1002(v0, v1 int32) {
	var v2, v3, v4, v5, v6, v7, v8, v9, v10 int32
	t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
	t1 := v0
	v2 = t0
	t2 := int32(load32(m.memory[uint32(v0):]))
	t3 := v2
	v3 = t2
	var p4 int32
	if uint32(t3) < uint32(v3) {
		p4 = 1
	}
	v4 = t1 + p4*i32(12)
	t5 := int32(load32(m.memory[int64(uint32(v0))+36:]))
	t6 := int32(load32(m.memory[int64(uint32(v0))+24:]))
	t7 := v4
	t8 := v0
	var p9 int32
	if uint32(t5) < uint32(t6) {
		p9 = 1
	}
	v5 = p9
	p10 := i32(24)
	if v5 != 0 {
		p10 = i32(36)
	}
	v6 = t8 + p10
	t11 := v6
	t12 := v0
	var p13 int32
	if uint32(v2) >= uint32(v3) {
		p13 = 1
	}
	v2 = t12 + p13*i32(12)
	t15 := v2
	t16 := v0
	p14 := i32(36)
	if v5 != 0 {
		p14 = i32(24)
	}
	v0 = t16 + p14
	t17 := int32(load32(m.memory[uint32(v0):]))
	t18 := int32(load32(m.memory[uint32(v2):]))
	var p19 int32
	if uint32(t17) < uint32(t18) {
		p19 = 1
	}
	v3 = p19
	p20 := t15
	if v3 != 0 {
		p20 = t11
	}
	t21 := int32(load32(m.memory[uint32(v6):]))
	t22 := int32(load32(m.memory[uint32(v4):]))
	var p23 int32
	if uint32(t21) < uint32(t22) {
		p23 = 1
	}
	v5 = p23
	p24 := p20
	if v5 != 0 {
		p24 = t7
	}
	v7 = p24
	t25 := int32(load32(m.memory[uint32(v7):]))
	v8 = t25
	t27 := v0
	p26 := v6
	if v5 != 0 {
		p26 = v2
	}
	p28 := p26
	if v3 != 0 {
		p28 = t27
	}
	v9 = p28
	t29 := int32(load32(m.memory[uint32(v9):]))
	v10 = t29
	t31 := v1
	p30 := v4
	if v5 != 0 {
		p30 = v6
	}
	v6 = p30
	t32 := int32(load32(m.memory[int64(uint32(v6))+8:]))
	store32(m.memory[int64(uint32(t31))+8:], uint32(t32))
	t33 := int64(load64(m.memory[uint32(v6):]))
	store64(m.memory[uint32(v1):], uint64(t33))
	t34 := v1
	t35 := v9
	t36 := v7
	var p37 int32
	if uint32(v10) < uint32(v8) {
		p37 = 1
	}
	v6 = p37
	p38 := t36
	if v6 != 0 {
		p38 = t35
	}
	v4 = p38
	t39 := int32(load32(m.memory[int64(uint32(v4))+8:]))
	store32(m.memory[int64(uint32(t34))+20:], uint32(t39))
	t40 := int64(load64(m.memory[uint32(v4):]))
	store64(m.memory[int64(uint32(v1))+12:], uint64(t40))
	t42 := v1
	p41 := v9
	if v6 != 0 {
		p41 = v7
	}
	v6 = p41
	t43 := int32(load32(m.memory[int64(uint32(v6))+8:]))
	store32(m.memory[int64(uint32(t42))+32:], uint32(t43))
	t44 := int64(load64(m.memory[uint32(v6):]))
	store64(m.memory[int64(uint32(v1))+24:], uint64(t44))
	t46 := v1
	p45 := v0
	if v3 != 0 {
		p45 = v2
	}
	v0 = p45
	t47 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	store32(m.memory[int64(uint32(t46))+44:], uint32(t47))
	t48 := int64(load64(m.memory[uint32(v0):]))
	store64(m.memory[int64(uint32(v1))+36:], uint64(t48))
}
func (m *Module) fn1003(v0, v1, v2, v3 int32) int32 {
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
		t4 := m.fn1003(t0, t2, t3+v5, v3)
		v0 = t4
		t5 := m.fn1003(v1, v1+v4, v1+v5, v3)
		v1 = t5
		t6 := m.fn1003(v2, v2+v4, v2+v5, v3)
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
func (m *Module) fn1004(v0, v1, v2 int32) {
	m.fn1002(v0, v2)
	m.fn1002(v0+i32(48), v2+i32(48))
	m.fn1006(v2, i32(8), v1)
}
func (m *Module) fn1005(v0, v1 int32) {
	var v2, v3 int32
	var v4 int64
	{
		t0 := int32(load32(m.memory[uint32(v1):]))
		v2 = t0
		t1 := v2
		v3 = v1 + i32(-12)
		t2 := int32(load32(m.memory[uint32(v3):]))
		if uint32(t1) >= uint32(t2) {
			return
		}
		t3 := int64(load64(m.memory[int64(uint32(v1))+4:]))
		v4 = t3
	l2:
		{
			v1 = v3
			t4 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			store32(m.memory[uint32(v1+i32(20)):], uint32(t4))
			t5 := int64(load64(m.memory[uint32(v1):]))
			store64(m.memory[uint32(v1+i32(12)):], uint64(t5))
			if v1 == v0 {
				goto l1
			}
			t6 := v2
			v3 = v1 + i32(-12)
			t7 := int32(load32(m.memory[uint32(v3):]))
			if uint32(t6) < uint32(t7) {
				goto l2
			}
		}
	l1:
		store64(m.memory[int64(uint32(v1))+4:], uint64(v4))
		store32(m.memory[uint32(v1):], uint32(v2))
	}
}
func (m *Module) fn1006(v0, v1, v2 int32) {
	var v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14 int32
	t0 := v2
	v3 = v1*i32(12) + i32(-12)
	v4 = t0 + v3
	v5 = v0 + v3
	t1 := v0
	v6 = int32(uint32(v1) >> 1)
	v3 = t1 + v6*i32(12)
	v7 = v3 + i32(-12)
l4:
	if v6 != 0 {
		t11 := int32(load32(m.memory[uint32(v3):]))
		t12 := v2
		t13 := v3
		t14 := v0
		v8 = t11
		t15 := int32(load32(m.memory[uint32(v0):]))
		t16 := v8
		v9 = t15
		var p17 int32
		if uint32(t16) < uint32(v9) {
			p17 = 1
		}
		v10 = p17
		p18 := t14
		if v10 != 0 {
			p18 = t13
		}
		v11 = p18
		t19 := int64(load64(m.memory[uint32(v11):]))
		store64(m.memory[uint32(t12):], uint64(t19))
		t20 := int32(load32(m.memory[int64(uint32(v11))+8:]))
		store32(m.memory[int64(uint32(v2))+8:], uint32(t20))
		t21 := int32(load32(m.memory[uint32(v5):]))
		t22 := v4
		t23 := v7
		t24 := v5
		v11 = t21
		t25 := int32(load32(m.memory[uint32(v7):]))
		t26 := v11
		v12 = t25
		var p27 int32
		if uint32(t26) < uint32(v12) {
			p27 = 1
		}
		v13 = p27
		p28 := t24
		if v13 != 0 {
			p28 = t23
		}
		v14 = p28
		t29 := int64(load64(m.memory[uint32(v14):]))
		store64(m.memory[uint32(t22):], uint64(t29))
		t30 := int32(load32(m.memory[int64(uint32(v14))+8:]))
		store32(m.memory[int64(uint32(v4))+8:], uint32(t30))
		v6 = v6 + i32(-1)
		v4 = v4 + i32(-12)
		v2 = v2 + i32(12)
		t32 := v7
		p31 := i32(0)
		if v13 != 0 {
			p31 = i32(-12)
		}
		v7 = t32 + p31
		t34 := v5
		p33 := i32(0)
		if uint32(v11) >= uint32(v12) {
			p33 = i32(-12)
		}
		v5 = t34 + p33
		t35 := v0
		var p36 int32
		if uint32(v8) >= uint32(v9) {
			p36 = 1
		}
		v0 = t35 + p36*i32(12)
		v3 = v3 + v10*i32(12)
		goto l4
	}
	v7 = v7 + i32(12)
	{
		if v1&i32(1) == 0 {
			goto l1
		}
		t2 := v2
		t3 := v0
		t4 := v3
		var p5 int32
		if uint32(v0) < uint32(v7) {
			p5 = 1
		}
		v4 = p5
		p6 := t4
		if v4 != 0 {
			p6 = t3
		}
		v6 = p6
		t7 := int32(load32(m.memory[int64(uint32(v6))+8:]))
		store32(m.memory[int64(uint32(t2))+8:], uint32(t7))
		t8 := int64(load64(m.memory[uint32(v6):]))
		store64(m.memory[uint32(v2):], uint64(t8))
		t9 := v3
		var p10 int32
		if uint32(v0) >= uint32(v7) {
			p10 = 1
		}
		v3 = t9 + p10*i32(12)
		v0 = v0 + v4*i32(12)
	}
l1:
	if v0 != v7 {
		goto l2
	}
	if v3 == v5+i32(12) {
		return
	}
l2:
	m.fn987()
	panic("unreachable")
}
func (m *Module) fn1007(v0, v1 int32) {
	var v2, v3, v4, v5, v6, v7, v8 int32
	var v9 int64
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	{
		t1 := int64(load64(m.memory[int64(uint32(v0))+8:]))
		t2 := int64(load64(m.memory[uint32(v0):]))
		if t1 < t2 {
			v3 = v0 + i32(8)
			v4 = i32(2)
		l5:
			{
				if v1 == v4 {
					t7 := v2
					t8 := v0
					v5 = int32(uint32(v1) >> 1)
					m.fn1009(t7, t8, v5, v5, i32(1301108))
					t9 := int32(load32(m.memory[int64(uint32(v2))+4:]))
					v6 = t9
					t10 := int32(load32(m.memory[uint32(v2):]))
					v3 = t10
					t11 := v2
					t12 := v0 + v1<<3
					v4 = v5 << 3
					m.fn1009(t11, t12-v4, v5, v5, i32(1301124))
					t13 := int32(load32(m.memory[uint32(v2):]))
					v1 = v4 + t13 + i32(-8)
					v4 = i32(0)
					t14 := int32(load32(m.memory[int64(uint32(v2))+4:]))
					t15 := v5 + i32(-1)
					v7 = t14
					var p16 int32
					if uint32(t15) < uint32(v7) {
						p16 = 1
					}
					v0 = p16
				l8:
					v8 = v5 + v4
					if v8 == 0 {
						goto l1
					}
					if v6+v4 == 0 {
						m.fn158(v6, v6, i32(1301140))
						panic("unreachable")
					}
					{
						if v0 == 0 {
							m.fn158(v8+i32(-1), v7, i32(1301156))
							panic("unreachable")
						}
						t17 := int64(load64(m.memory[uint32(v3):]))
						v9 = t17
						t18 := int64(load64(m.memory[uint32(v1):]))
						store64(m.memory[uint32(v3):], uint64(t18))
						store64(m.memory[uint32(v1):], uint64(v9))
						v3 = v3 + i32(8)
						v1 = v1 + i32(-8)
						v4 = v4 + i32(-1)
						goto l8
					}
				}
				v5 = v3 + i32(8)
				t5 := int64(load64(m.memory[uint32(v5):]))
				t6 := int64(load64(m.memory[uint32(v3):]))
				if t5 >= t6 {
					goto l2
				}
				v4 = v4 + i32(1)
				v3 = v5
				goto l5
			}
		}
		v3 = v0 + i32(8)
		v4 = i32(2)
	l3:
		{
			if v1 == v4 {
				goto l1
			}
			v5 = v3 + i32(8)
			t3 := int64(load64(m.memory[uint32(v5):]))
			t4 := int64(load64(m.memory[uint32(v3):]))
			if t3 < t4 {
				goto l2
			}
			v4 = v4 + i32(1)
			v3 = v5
			goto l3
		}
	}
l2:
	m.fn1008(v0, v1, i32(0), int32(bits.LeadingZeros32(uint32(v1|i32(1))))<<1^i32(62))
	goto l1
l1:
	m.g0 = v2 + i32(16)
}
func (m *Module) fn1008(v0, v1, v2, v3 int32) {
	var v4, v5, v6, v7, v8, v9 int32
	var v10, v11, v12, v13, v14, v15, v16, v17, v18, v19, v20, v21, v22, v23, v24, v25, v26, v27, v28, v29, v30, v31, v32, v33 int64
	var v34 int32
	t0 := m.g0
	v4 = t0 - i32(272)
	m.g0 = v4
l22:
	{
		if uint32(v1) < uint32(i32(33)) {
			if uint32(v1) < uint32(i32(2)) {
				goto l2
			}
			t1 := v1
			v5 = int32(uint32(v1) >> 1)
			t2 := v5
			var p3 int32
			if uint32(v1) < uint32(i32(18)) {
				p3 = 1
			}
			v3 = p3
			p4 := t2
			if v3 != 0 {
				p4 = t1
			}
			v6 = p4
			v2 = v1 - v5
			v7 = v0 + v5<<3
			v8 = v0
		l5:
			{
				{
					{
						if uint32(v6) > uint32(i32(12)) {
							goto l3
						}
						v9 = i32(1)
						if uint32(v6) <= uint32(i32(8)) {
							goto l4
						}
						t5 := int64(load64(m.memory[int64(uint32(v8))+64:]))
						t6 := v8
						v10 = t5
						t7 := int64(load64(m.memory[int64(uint32(v8))+32:]))
						t8 := v10
						v11 = t7
						p9 := v11
						if v10 > v11 {
							p9 = t8
						}
						v12 = p9
						t10 := int64(load64(m.memory[int64(uint32(v8))+24:]))
						t11 := v12
						v13 = t10
						t12 := int64(load64(m.memory[uint32(v8):]))
						t13 := v13
						v14 = t12
						p14 := v14
						if v13 > v14 {
							p14 = t13
						}
						v15 = p14
						p15 := v15
						if v12 > v15 {
							p15 = t11
						}
						v16 = p15
						t16 := int64(load64(m.memory[int64(uint32(v8))+56:]))
						t17 := v16
						v17 = t16
						t18 := int64(load64(m.memory[int64(uint32(v8))+8:]))
						t19 := v17
						v18 = t18
						p20 := v18
						if v17 > v18 {
							p20 = t19
						}
						v19 = p20
						t22 := v19
						p21 := v14
						if v13 < v14 {
							p21 = v13
						}
						v13 = p21
						p23 := v13
						if v19 > v13 {
							p23 = t22
						}
						v14 = p23
						p24 := v14
						if v16 > v14 {
							p24 = t17
						}
						v20 = p24
						t25 := int64(load64(m.memory[int64(uint32(v8))+48:]))
						t26 := v20
						v21 = t25
						t27 := int64(load64(m.memory[int64(uint32(v8))+40:]))
						t28 := v21
						v22 = t27
						t29 := int64(load64(m.memory[int64(uint32(v8))+16:]))
						t30 := v22
						v23 = t29
						p31 := v23
						if v22 > v23 {
							p31 = t30
						}
						v24 = p31
						p32 := v24
						if v21 > v24 {
							p32 = t28
						}
						v25 = p32
						t34 := v25
						p33 := v15
						if v12 < v15 {
							p33 = v12
						}
						v12 = p33
						t36 := v12
						p35 := v18
						if v17 < v18 {
							p35 = v17
						}
						v15 = p35
						p37 := v15
						if v12 > v15 {
							p37 = t36
						}
						v17 = p37
						p38 := v17
						if v25 > v17 {
							p38 = t34
						}
						v18 = p38
						p39 := v18
						if v20 > v18 {
							p39 = t26
						}
						store64(m.memory[int64(uint32(t6))+64:], uint64(p39))
						t41 := v8
						p40 := v24
						if v21 < v24 {
							p40 = v21
						}
						v21 = p40
						t43 := v21
						p42 := v11
						if v10 < v11 {
							p42 = v10
						}
						v10 = p42
						t45 := v10
						p44 := v23
						if v22 < v23 {
							p44 = v22
						}
						v11 = p44
						p46 := v11
						if v10 > v11 {
							p46 = t45
						}
						v22 = p46
						p47 := v22
						if v21 < v22 {
							p47 = t43
						}
						v23 = p47
						t49 := v23
						p48 := v15
						if v12 < v15 {
							p48 = v12
						}
						v12 = p48
						p50 := v12
						if v23 < v12 {
							p50 = t49
						}
						v15 = p50
						t52 := v15
						p51 := v11
						if v10 < v11 {
							p51 = v10
						}
						v10 = p51
						t54 := v10
						p53 := v13
						if v19 < v13 {
							p53 = v19
						}
						v11 = p53
						p55 := v11
						if v10 < v11 {
							p55 = t54
						}
						v13 = p55
						p56 := v13
						if v15 < v13 {
							p56 = t52
						}
						store64(m.memory[uint32(t41):], uint64(p56))
						t58 := v8
						p57 := v14
						if v16 < v14 {
							p57 = v16
						}
						v14 = p57
						t60 := v14
						p59 := v22
						if v21 > v22 {
							p59 = v21
						}
						v16 = p59
						p61 := v16
						if v14 > v16 {
							p61 = t60
						}
						v19 = p61
						t63 := v19
						p62 := v18
						if v20 < v18 {
							p62 = v20
						}
						v18 = p62
						p64 := v18
						if v19 > v18 {
							p64 = t63
						}
						store64(m.memory[int64(uint32(t58))+56:], uint64(p64))
						t66 := v8
						p65 := v18
						if v19 < v18 {
							p65 = v19
						}
						v18 = p65
						t68 := v18
						p67 := v16
						if v14 < v16 {
							p67 = v14
						}
						v14 = p67
						t70 := v14
						p69 := v17
						if v25 < v17 {
							p69 = v25
						}
						v16 = p69
						p71 := v16
						if v14 > v16 {
							p71 = t70
						}
						v17 = p71
						t73 := v17
						p72 := v12
						if v23 > v12 {
							p72 = v23
						}
						v12 = p72
						t75 := v12
						p74 := v11
						if v10 > v11 {
							p74 = v10
						}
						v10 = p74
						p76 := v10
						if v12 > v10 {
							p76 = t75
						}
						v11 = p76
						p77 := v11
						if v17 > v11 {
							p77 = t73
						}
						v19 = p77
						p78 := v19
						if v18 > v19 {
							p78 = t68
						}
						store64(m.memory[int64(uint32(t66))+48:], uint64(p78))
						t80 := v8
						p79 := v19
						if v18 < v19 {
							p79 = v18
						}
						store64(m.memory[int64(uint32(t80))+40:], uint64(p79))
						t82 := v8
						p81 := v11
						if v17 < v11 {
							p81 = v17
						}
						v11 = p81
						t84 := v11
						p83 := v16
						if v14 < v16 {
							p83 = v14
						}
						v14 = p83
						t86 := v14
						p85 := v10
						if v12 < v10 {
							p85 = v12
						}
						v10 = p85
						p87 := v10
						if v14 > v10 {
							p87 = t86
						}
						v12 = p87
						p88 := v12
						if v11 > v12 {
							p88 = t84
						}
						store64(m.memory[int64(uint32(t82))+32:], uint64(p88))
						t90 := v8
						p89 := v12
						if v11 < v12 {
							p89 = v11
						}
						store64(m.memory[int64(uint32(t90))+24:], uint64(p89))
						t92 := v8
						p91 := v10
						if v14 < v10 {
							p91 = v14
						}
						v10 = p91
						t94 := v10
						p93 := v13
						if v15 > v13 {
							p93 = v15
						}
						v11 = p93
						p95 := v11
						if v10 > v11 {
							p95 = t94
						}
						store64(m.memory[int64(uint32(t92))+16:], uint64(p95))
						t97 := v8
						p96 := v11
						if v10 < v11 {
							p96 = v10
						}
						store64(m.memory[int64(uint32(t97))+8:], uint64(p96))
						v9 = i32(9)
						goto l4
					}
				l3:
					t98 := int64(load64(m.memory[int64(uint32(v8))+96:]))
					t99 := v8
					v10 = t98
					t100 := int64(load64(m.memory[uint32(v8):]))
					t101 := v10
					v11 = t100
					p102 := v11
					if v10 > v11 {
						p102 = t101
					}
					v12 = p102
					t103 := int64(load64(m.memory[int64(uint32(v8))+88:]))
					t104 := v12
					v13 = t103
					t105 := int64(load64(m.memory[int64(uint32(v8))+40:]))
					t106 := v13
					v14 = t105
					p107 := v14
					if v13 > v14 {
						p107 = t106
					}
					v15 = p107
					t108 := int64(load64(m.memory[int64(uint32(v8))+32:]))
					t109 := v15
					v16 = t108
					p110 := v16
					if v15 > v16 {
						p110 = t109
					}
					v17 = p110
					p111 := v17
					if v12 > v17 {
						p111 = t104
					}
					v18 = p111
					t112 := int64(load64(m.memory[int64(uint32(v8))+80:]))
					t113 := v18
					v19 = t112
					t114 := int64(load64(m.memory[int64(uint32(v8))+8:]))
					t115 := v19
					v20 = t114
					p116 := v20
					if v19 > v20 {
						p116 = t115
					}
					v21 = p116
					t117 := int64(load64(m.memory[int64(uint32(v8))+64:]))
					t118 := v21
					v22 = t117
					t119 := int64(load64(m.memory[int64(uint32(v8))+48:]))
					t120 := v22
					v23 = t119
					p121 := v23
					if v22 > v23 {
						p121 = t120
					}
					v24 = p121
					p122 := v24
					if v21 > v24 {
						p122 = t118
					}
					v25 = p122
					t123 := int64(load64(m.memory[int64(uint32(v8))+72:]))
					t124 := v25
					v26 = t123
					t125 := int64(load64(m.memory[int64(uint32(v8))+16:]))
					t126 := v26
					v27 = t125
					p127 := v27
					if v26 > v27 {
						p127 = t126
					}
					v28 = p127
					t128 := int64(load64(m.memory[int64(uint32(v8))+56:]))
					t129 := v28
					v29 = t128
					t130 := int64(load64(m.memory[int64(uint32(v8))+24:]))
					t131 := v29
					v30 = t130
					p132 := v30
					if v29 > v30 {
						p132 = t131
					}
					v31 = p132
					p133 := v31
					if v28 > v31 {
						p133 = t129
					}
					v32 = p133
					p134 := v32
					if v25 > v32 {
						p134 = t124
					}
					v33 = p134
					p135 := v33
					if v18 > v33 {
						p135 = t113
					}
					store64(m.memory[int64(uint32(t99))+96:], uint64(p135))
					t137 := v8
					p136 := v17
					if v12 < v17 {
						p136 = v12
					}
					v12 = p136
					t139 := v12
					p138 := v24
					if v21 < v24 {
						p138 = v21
					}
					v17 = p138
					t141 := v17
					p140 := v31
					if v28 < v31 {
						p140 = v28
					}
					v21 = p140
					p142 := v21
					if v17 > v21 {
						p142 = t141
					}
					v24 = p142
					p143 := v24
					if v12 > v24 {
						p143 = t139
					}
					v28 = p143
					t145 := v28
					p144 := v23
					if v22 < v23 {
						p144 = v22
					}
					v22 = p144
					t147 := v22
					p146 := v20
					if v19 < v20 {
						p146 = v19
					}
					v19 = p146
					p148 := v19
					if v22 > v19 {
						p148 = t147
					}
					v20 = p148
					t150 := v20
					p149 := v30
					if v29 < v30 {
						p149 = v29
					}
					v23 = p149
					t152 := v23
					p151 := v27
					if v26 < v27 {
						p151 = v26
					}
					v26 = p151
					p153 := v26
					if v23 > v26 {
						p153 = t152
					}
					v27 = p153
					p154 := v27
					if v20 > v27 {
						p154 = t150
					}
					v29 = p154
					t156 := v29
					p155 := v16
					if v15 < v16 {
						p155 = v15
					}
					v15 = p155
					t158 := v15
					p157 := v11
					if v10 < v11 {
						p157 = v10
					}
					v10 = p157
					p159 := v10
					if v15 > v10 {
						p159 = t158
					}
					v11 = p159
					p160 := v11
					if v29 > v11 {
						p160 = t156
					}
					v16 = p160
					p161 := v16
					if v28 > v16 {
						p161 = t145
					}
					v30 = p161
					t163 := v30
					p162 := v33
					if v18 < v33 {
						p162 = v18
					}
					v18 = p162
					t165 := v18
					p164 := v32
					if v25 < v32 {
						p164 = v25
					}
					v25 = p164
					t167 := v25
					p166 := v14
					if v13 < v14 {
						p166 = v13
					}
					v13 = p166
					p168 := v13
					if v25 > v13 {
						p168 = t167
					}
					v14 = p168
					p169 := v14
					if v18 > v14 {
						p169 = t165
					}
					v31 = p169
					p170 := v31
					if v30 > v31 {
						p170 = t163
					}
					store64(m.memory[int64(uint32(t137))+88:], uint64(p170))
					t172 := v8
					p171 := v26
					if v23 < v26 {
						p171 = v23
					}
					v23 = p171
					t174 := v23
					p173 := v19
					if v22 < v19 {
						p173 = v22
					}
					v19 = p173
					p175 := v19
					if v23 < v19 {
						p175 = t174
					}
					v22 = p175
					t177 := v22
					p176 := v13
					if v25 < v13 {
						p176 = v25
					}
					v13 = p176
					t179 := v13
					p178 := v10
					if v15 < v10 {
						p178 = v15
					}
					v10 = p178
					p180 := v10
					if v13 < v10 {
						p180 = t179
					}
					v15 = p180
					p181 := v15
					if v22 < v15 {
						p181 = t177
					}
					store64(m.memory[uint32(t172):], uint64(p181))
					t183 := v8
					p182 := v31
					if v30 < v31 {
						p182 = v30
					}
					v25 = p182
					t185 := v25
					p184 := v14
					if v18 < v14 {
						p184 = v18
					}
					v14 = p184
					t187 := v14
					p186 := v16
					if v28 < v16 {
						p186 = v28
					}
					v16 = p186
					p188 := v16
					if v14 > v16 {
						p188 = t187
					}
					v18 = p188
					p189 := v18
					if v25 > v18 {
						p189 = t185
					}
					store64(m.memory[int64(uint32(t183))+80:], uint64(p189))
					t191 := v8
					p190 := v21
					if v17 < v21 {
						p190 = v17
					}
					v17 = p190
					t193 := v17
					p192 := v11
					if v29 < v11 {
						p192 = v29
					}
					v11 = p192
					p194 := v11
					if v17 < v11 {
						p194 = t193
					}
					v21 = p194
					t196 := v21
					p195 := v10
					if v13 > v10 {
						p195 = v13
					}
					v10 = p195
					t198 := v10
					p197 := v19
					if v23 > v19 {
						p197 = v23
					}
					v13 = p197
					p199 := v13
					if v10 < v13 {
						p199 = t198
					}
					v19 = p199
					p200 := v19
					if v21 < v19 {
						p200 = t196
					}
					v23 = p200
					t202 := v23
					p201 := v24
					if v12 < v24 {
						p201 = v12
					}
					v12 = p201
					t204 := v12
					p203 := v27
					if v20 < v27 {
						p203 = v20
					}
					v20 = p203
					p205 := v20
					if v12 < v20 {
						p205 = t204
					}
					v24 = p205
					t207 := v24
					p206 := v15
					if v22 > v15 {
						p206 = v22
					}
					v15 = p206
					p208 := v15
					if v24 < v15 {
						p208 = t207
					}
					v22 = p208
					p209 := v22
					if v23 < v22 {
						p209 = t202
					}
					store64(m.memory[int64(uint32(t191))+8:], uint64(p209))
					t211 := v8
					p210 := v18
					if v25 < v18 {
						p210 = v25
					}
					v18 = p210
					t213 := v18
					p212 := v20
					if v12 > v20 {
						p212 = v12
					}
					v12 = p212
					t215 := v12
					p214 := v11
					if v17 > v11 {
						p214 = v17
					}
					v11 = p214
					p216 := v11
					if v12 > v11 {
						p216 = t215
					}
					v17 = p216
					t218 := v17
					p217 := v16
					if v14 < v16 {
						p217 = v14
					}
					v14 = p217
					t220 := v14
					p219 := v13
					if v10 > v13 {
						p219 = v10
					}
					v10 = p219
					p221 := v10
					if v14 > v10 {
						p221 = t220
					}
					v13 = p221
					p222 := v13
					if v17 > v13 {
						p222 = t218
					}
					v16 = p222
					p223 := v16
					if v18 > v16 {
						p223 = t213
					}
					store64(m.memory[int64(uint32(t211))+72:], uint64(p223))
					t225 := v8
					p224 := v16
					if v18 < v16 {
						p224 = v18
					}
					store64(m.memory[int64(uint32(t225))+64:], uint64(p224))
					t227 := v8
					p226 := v11
					if v12 < v11 {
						p226 = v12
					}
					v11 = p226
					t229 := v11
					p228 := v10
					if v14 < v10 {
						p228 = v14
					}
					v10 = p228
					p230 := v10
					if v11 > v10 {
						p230 = t229
					}
					v12 = p230
					t232 := v12
					p231 := v13
					if v17 < v13 {
						p231 = v17
					}
					v13 = p231
					p233 := v13
					if v12 > v13 {
						p233 = t232
					}
					store64(m.memory[int64(uint32(t227))+56:], uint64(p233))
					t235 := v8
					p234 := v19
					if v21 > v19 {
						p234 = v21
					}
					v14 = p234
					t237 := v14
					p236 := v15
					if v24 > v15 {
						p236 = v24
					}
					v15 = p236
					p238 := v15
					if v14 < v15 {
						p238 = t237
					}
					v16 = p238
					t240 := v16
					p239 := v22
					if v23 > v22 {
						p239 = v23
					}
					v17 = p239
					p241 := v17
					if v16 < v17 {
						p241 = t240
					}
					store64(m.memory[int64(uint32(t235))+16:], uint64(p241))
					t243 := v8
					p242 := v13
					if v12 < v13 {
						p242 = v12
					}
					v12 = p242
					t245 := v12
					p244 := v10
					if v11 < v10 {
						p244 = v11
					}
					v10 = p244
					t247 := v10
					p246 := v15
					if v14 > v15 {
						p246 = v14
					}
					v11 = p246
					p248 := v11
					if v10 > v11 {
						p248 = t247
					}
					v13 = p248
					p249 := v13
					if v12 > v13 {
						p249 = t245
					}
					store64(m.memory[int64(uint32(t243))+48:], uint64(p249))
					t251 := v8
					p250 := v13
					if v12 < v13 {
						p250 = v12
					}
					store64(m.memory[int64(uint32(t251))+40:], uint64(p250))
					t253 := v8
					p252 := v11
					if v10 < v11 {
						p252 = v10
					}
					v10 = p252
					t255 := v10
					p254 := v17
					if v16 > v17 {
						p254 = v16
					}
					v11 = p254
					p256 := v11
					if v10 > v11 {
						p256 = t255
					}
					store64(m.memory[int64(uint32(t253))+32:], uint64(p256))
					t258 := v8
					p257 := v11
					if v10 < v11 {
						p257 = v10
					}
					store64(m.memory[int64(uint32(t258))+24:], uint64(p257))
					v9 = i32(13)
				}
			l4:
				m.fn1011(v8, v6, v9)
				if v3 != 0 {
					goto l2
				}
				var p259 int32
				if v8 == v0 {
					p259 = 1
				}
				v9 = p259
				v6 = v2
				v8 = v7
				if v9 != 0 {
					goto l5
				}
			}
			v6 = v7 + i32(-8)
			t260 := v0
			v8 = v1<<3 + i32(-8)
			v9 = t260 + v8
			v2 = v4 + i32(8) + v8
			v3 = v4 + i32(8)
			v8 = v0
		l9:
			if v5 != 0 {
				t269 := int64(load64(m.memory[uint32(v7):]))
				t270 := v3
				v12 = t269
				t271 := int64(load64(m.memory[uint32(v8):]))
				t272 := v12
				v13 = t271
				t273 := v13
				var p274 int32
				if v12 < v13 {
					p274 = 1
				}
				v34 = p274
				p275 := t273
				if v34 != 0 {
					p275 = t272
				}
				store64(m.memory[uint32(t270):], uint64(p275))
				t276 := int64(load64(m.memory[uint32(v9):]))
				t277 := v2
				v10 = t276
				t278 := int64(load64(m.memory[uint32(v6):]))
				t279 := v10
				v11 = t278
				p280 := v11
				if v10 > v11 {
					p280 = t279
				}
				store64(m.memory[uint32(t277):], uint64(p280))
				v5 = v5 + i32(-1)
				v2 = v2 + i32(-8)
				v3 = v3 + i32(8)
				t282 := v6
				p281 := i32(0)
				if v10 < v11 {
					p281 = i32(-8)
				}
				v6 = t282 + p281
				t284 := v9
				p283 := i32(0)
				if v10 >= v11 {
					p283 = i32(-8)
				}
				v9 = t284 + p283
				t285 := v8
				var p286 int32
				if v12 >= v13 {
					p286 = 1
				}
				v8 = t285 + p286<<3
				v7 = v7 + v34<<3
				goto l9
			}
			v6 = v6 + i32(8)
			{
				if v1&i32(1) == 0 {
					goto l7
				}
				t261 := v3
				t262 := v8
				t263 := v7
				var p264 int32
				if uint32(v8) < uint32(v6) {
					p264 = 1
				}
				v5 = p264
				p265 := t263
				if v5 != 0 {
					p265 = t262
				}
				t266 := int64(load64(m.memory[uint32(p265):]))
				store64(m.memory[uint32(t261):], uint64(t266))
				t267 := v7
				var p268 int32
				if uint32(v8) >= uint32(v6) {
					p268 = 1
				}
				v7 = t267 + p268<<3
				v8 = v8 + v5<<3
			}
		l7:
			if v8 != v6 {
				goto l8
			}
			if v7 != v9+i32(8) {
				goto l8
			}
			v8 = v1 << 3
			if v8 == 0 {
				goto l2
			}
			memory_copy(m.memory, uint32(v0), uint32(v4+i32(8)), uint32(v8))
			goto l2
		l8:
			m.fn987()
			panic("unreachable")
		}
		if v3 != 0 {
			goto l1
		}
		m.fn1010(v0, v1)
		goto l2
	l1:
		t287 := v0
		v8 = int32(uint32(v1) >> 3)
		v7 = t287 + v8*i32(56)
		v6 = v0 + v8<<5
		{
			{
				if uint32(v1) < uint32(i32(64)) {
					goto l10
				}
				t288 := m.fn1012(v0, v6, v7, v8)
				v8 = t288
				goto l11
			}
		l10:
			t289 := int64(load64(m.memory[uint32(v0):]))
			t290 := v0
			t291 := v7
			t292 := v6
			v10 = t289
			t293 := int64(load64(m.memory[uint32(v6):]))
			t294 := v10
			v11 = t293
			var p295 int32
			if t294 < v11 {
				p295 = 1
			}
			v8 = p295
			t296 := int64(load64(m.memory[uint32(v7):]))
			t297 := v8
			t298 := v11
			v12 = t296
			var p299 int32
			if t298 < v12 {
				p299 = 1
			}
			p300 := t292
			if t297^p299 != 0 {
				p300 = t291
			}
			t301 := v8
			var p302 int32
			if v10 < v12 {
				p302 = 1
			}
			p303 := p300
			if t301^p302 != 0 {
				p303 = t290
			}
			v8 = p303
		}
	l11:
		v3 = v3 + i32(-1)
		v8 = v8 - v0
		{
			if v2 != 0 {
				t306 := int64(load64(m.memory[uint32(v0):]))
				v11 = t306
				t307 := int64(load64(m.memory[uint32(v2):]))
				v7 = v0 + v8
				t308 := int64(load64(m.memory[uint32(v7):]))
				v10 = t308
				if t307 < v10 {
					goto l13
				}
				store64(m.memory[uint32(v0):], uint64(v10))
				store64(m.memory[uint32(v7):], uint64(v11))
				m.fn1009(v4+i32(8), v0, v1, i32(1), i32(1301028))
				{
					t309 := int32(load32(m.memory[int64(uint32(v4))+12:]))
					if t309 == 0 {
						m.fn158(i32(0), i32(0), i32(1301044))
						panic("unreachable")
					}
					{
						{
							t310 := int32(load32(m.memory[int64(uint32(v4))+20:]))
							v5 = t310
							if v5 != 0 {
								goto l15
							}
							v8 = i32(0)
							goto l16
						}
					l15:
						t311 := int32(load32(m.memory[int64(uint32(v4))+8:]))
						v6 = t311
						t312 := int32(load32(m.memory[int64(uint32(v4))+16:]))
						t313 := v4
						v7 = t312
						t314 := int64(load64(m.memory[uint32(v7):]))
						store64(m.memory[int64(uint32(t313))+264:], uint64(t314))
						store32(m.memory[int64(uint32(v4))+20:], uint32(i32(0)))
						t315 := v4
						v8 = v7 + i32(8)
						store32(m.memory[int64(uint32(t315))+16:], uint32(v8))
						store32(m.memory[int64(uint32(v4))+8:], uint32(v7))
						v5 = v7 + v5<<3
						v9 = v5 + i32(-8)
						store32(m.memory[int64(uint32(v4))+12:], uint32(v4+i32(264)))
					l20:
						if uint32(v8) < uint32(v9) {
							t318 := int64(load64(m.memory[uint32(v6):]))
							m.fn1013(t318, v7, v4+i32(8))
							t319 := int64(load64(m.memory[uint32(v6):]))
							m.fn1013(t319, v7, v4+i32(8))
							t320 := int32(load32(m.memory[int64(uint32(v4))+16:]))
							v8 = t320
							goto l20
						}
					l19:
						{
							if v8 == v5 {
								goto l18
							}
							t316 := int64(load64(m.memory[uint32(v6):]))
							m.fn1013(t316, v7, v4+i32(8))
							t317 := int32(load32(m.memory[int64(uint32(v4))+16:]))
							v8 = t317
							goto l19
						}
					l18:
						t321 := int32(load32(m.memory[int64(uint32(v4))+12:]))
						store32(m.memory[int64(uint32(v4))+16:], uint32(t321))
						t322 := int64(load64(m.memory[uint32(v6):]))
						m.fn1013(t322, v7, v4+i32(8))
						t323 := int32(load32(m.memory[int64(uint32(v4))+20:]))
						v8 = t323
					}
				l16:
					if uint32(v8) >= uint32(v1) {
						goto l21
					}
					t324 := int64(load64(m.memory[uint32(v0):]))
					v10 = t324
					t325 := v0
					v7 = v0 + v8<<3
					t326 := int64(load64(m.memory[uint32(v7):]))
					store64(m.memory[uint32(t325):], uint64(t326))
					store64(m.memory[uint32(v7):], uint64(v10))
					t327 := v1
					v8 = v8 + i32(1)
					v1 = t327 - v8
					v0 = v0 + v8<<3
					v2 = i32(0)
					goto l22
				}
			}
			t304 := int64(load64(m.memory[uint32(v0+v8):]))
			v10 = t304
			t305 := int64(load64(m.memory[uint32(v0):]))
			v11 = t305
			goto l13
		}
	l13:
		store64(m.memory[uint32(v0):], uint64(v10))
		store64(m.memory[uint32(v0+v8):], uint64(v11))
		m.fn1009(v4+i32(8), v0, v1, i32(1), i32(1301028))
		{
			t328 := int32(load32(m.memory[int64(uint32(v4))+12:]))
			if t328 == 0 {
				goto l23
			}
			{
				{
					t329 := int32(load32(m.memory[int64(uint32(v4))+20:]))
					v5 = t329
					if v5 != 0 {
						goto l24
					}
					v8 = i32(0)
					goto l25
				}
			l24:
				t330 := int32(load32(m.memory[int64(uint32(v4))+8:]))
				v6 = t330
				t331 := int32(load32(m.memory[int64(uint32(v4))+16:]))
				t332 := v4
				v7 = t331
				t333 := int64(load64(m.memory[uint32(v7):]))
				store64(m.memory[int64(uint32(t332))+264:], uint64(t333))
				store32(m.memory[int64(uint32(v4))+20:], uint32(i32(0)))
				t334 := v4
				v8 = v7 + i32(8)
				store32(m.memory[int64(uint32(t334))+16:], uint32(v8))
				store32(m.memory[int64(uint32(v4))+8:], uint32(v7))
				v5 = v7 + v5<<3
				v9 = v5 + i32(-8)
				store32(m.memory[int64(uint32(v4))+12:], uint32(v4+i32(264)))
			l29:
				if uint32(v8) < uint32(v9) {
					t337 := int64(load64(m.memory[uint32(v6):]))
					m.fn1014(t337, v7, v4+i32(8))
					t338 := int64(load64(m.memory[uint32(v6):]))
					m.fn1014(t338, v7, v4+i32(8))
					t339 := int32(load32(m.memory[int64(uint32(v4))+16:]))
					v8 = t339
					goto l29
				}
			l28:
				{
					if v8 == v5 {
						goto l27
					}
					t335 := int64(load64(m.memory[uint32(v6):]))
					m.fn1014(t335, v7, v4+i32(8))
					t336 := int32(load32(m.memory[int64(uint32(v4))+16:]))
					v8 = t336
					goto l28
				}
			l27:
				t340 := int32(load32(m.memory[int64(uint32(v4))+12:]))
				store32(m.memory[int64(uint32(v4))+16:], uint32(t340))
				t341 := int64(load64(m.memory[uint32(v6):]))
				m.fn1014(t341, v7, v4+i32(8))
				t342 := int32(load32(m.memory[int64(uint32(v4))+20:]))
				v8 = t342
			}
		l25:
			if uint32(v8) >= uint32(v1) {
				goto l21
			}
			t343 := int64(load64(m.memory[uint32(v0):]))
			v10 = t343
			t344 := v0
			v7 = v0 + v8<<3
			t345 := int64(load64(m.memory[uint32(v7):]))
			store64(m.memory[uint32(t344):], uint64(t345))
			store64(m.memory[uint32(v7):], uint64(v10))
			m.fn1009(v4+i32(8), v0, v1, v8, i32(1301060))
			t346 := int32(load32(m.memory[int64(uint32(v4))+12:]))
			v8 = t346
			t347 := int32(load32(m.memory[int64(uint32(v4))+8:]))
			v7 = t347
			t348 := int32(load32(m.memory[int64(uint32(v4))+16:]))
			t349 := int32(load32(m.memory[int64(uint32(v4))+20:]))
			m.fn1009(v4+i32(8), t348, t349, i32(1), i32(1301076))
			t350 := int32(load32(m.memory[int64(uint32(v4))+12:]))
			if t350 == 0 {
				m.fn158(i32(0), i32(0), i32(1301092))
				panic("unreachable")
			}
			t351 := int32(load32(m.memory[int64(uint32(v4))+20:]))
			v1 = t351
			t352 := int32(load32(m.memory[int64(uint32(v4))+16:]))
			v0 = t352
			t353 := int32(load32(m.memory[int64(uint32(v4))+8:]))
			v6 = t353
			m.fn1008(v7, v8, v2, v3)
			v2 = v6
			goto l22
		}
	l23:
	}
	m.fn158(i32(0), i32(0), i32(1301044))
l21:
	panic("unreachable")
l2:
	m.g0 = v4 + i32(272)
}
func (m *Module) fn1009(v0, v1, v2, v3, v4 int32) {
	if uint32(v2) >= uint32(v3) {
		goto l0
	}
	m.fn91(i32(1301172), i32(19), v4)
	panic("unreachable")
l0:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
	store32(m.memory[uint32(v0):], uint32(v1))
	store32(m.memory[int64(uint32(v0))+12:], uint32(v2-v3))
	store32(m.memory[int64(uint32(v0))+8:], uint32(v1+v3<<3))
}
func (m *Module) fn1010(v0, v1 int32) {
	var v2, v3 int32
	var v4 int64
	var v5, v6, v7 int32
	v2 = int32(uint32(v1)>>1) + v1
l3:
	{
		if v2 == 0 {
			return
		}
		{
			v2 = v2 + i32(-1)
			if uint32(v2) < uint32(v1) {
				goto l1
			}
			v3 = v2 - v1
			goto l2
		l1:
			t0 := int64(load64(m.memory[uint32(v0):]))
			v4 = t0
			t1 := v0
			v5 = v0 + v2<<3
			t2 := int64(load64(m.memory[uint32(v5):]))
			store64(m.memory[uint32(t1):], uint64(t2))
			store64(m.memory[uint32(v5):], uint64(v4))
			v3 = i32(0)
		}
	l2:
		p3 := v2
		if uint32(v1) < uint32(v2) {
			p3 = v1
		}
		v6 = p3
	l5:
		{
			v7 = v3 << 1
			v5 = v7 | i32(1)
			if uint32(v5) >= uint32(v6) {
				goto l3
			}
			{
				v7 = v7 + i32(2)
				if uint32(v7) >= uint32(v6) {
					goto l4
				}
				t4 := int64(load64(m.memory[uint32(v0+v5<<3):]))
				t5 := int64(load64(m.memory[uint32(v0+v7<<3):]))
				t6 := v5
				var p7 int32
				if t4 < t5 {
					p7 = 1
				}
				v5 = t6 + p7
			}
		l4:
			v3 = v0 + v3<<3
			t8 := int64(load64(m.memory[uint32(v3):]))
			v7 = v0 + v5<<3
			t9 := int64(load64(m.memory[uint32(v7):]))
			if t8 >= t9 {
				goto l3
			}
			m.fn244(v3, v7, i32(2))
			v3 = v5
			goto l5
		}
	}
}
func (m *Module) fn1011(v0, v1, v2 int32) {
	var v3, v4, v5 int32
	var v6, v7 int64
	{
		if uint32(v2+i32(-1)) >= uint32(v1) {
			panic("unreachable")
		}
		t0 := v0
		v3 = v2 << 3
		v4 = t0 + v3
		v5 = v0 + v1<<3
	l6:
		if v4 == v5 {
			return
		}
		{
			t1 := int64(load64(m.memory[uint32(v4):]))
			v6 = t1
			t2 := int64(load64(m.memory[uint32(v4+i32(-8)):]))
			t3 := v6
			v7 = t2
			if t3 >= v7 {
				goto l2
			}
			v2 = v3
		l5:
			{
				v1 = v0 + v2
				store64(m.memory[uint32(v1):], uint64(v7))
				if v2 != i32(8) {
					goto l3
				}
				v2 = v0
				goto l4
			l3:
				v2 = v2 + i32(-8)
				t4 := int64(load64(m.memory[uint32(v1+i32(-16)):]))
				t5 := v6
				v7 = t4
				if t5 < v7 {
					goto l5
				}
			}
			v2 = v0 + v2
		l4:
			store64(m.memory[uint32(v2):], uint64(v6))
		}
	l2:
		v3 = v3 + i32(8)
		v4 = v4 + i32(8)
		goto l6
	}
}
func (m *Module) fn1012(v0, v1, v2, v3 int32) int32 {
	var v4, v5 int32
	var v6, v7, v8 int64
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
		t4 := m.fn1012(t0, t2, t3+v5, v3)
		v0 = t4
		t5 := m.fn1012(v1, v1+v4, v1+v5, v3)
		v1 = t5
		t6 := m.fn1012(v2, v2+v4, v2+v5, v3)
		v2 = t6
	}
l0:
	t7 := int64(load64(m.memory[uint32(v0):]))
	t8 := v0
	t9 := v2
	t10 := v1
	v6 = t7
	t11 := int64(load64(m.memory[uint32(v1):]))
	t12 := v6
	v7 = t11
	var p13 int32
	if t12 < v7 {
		p13 = 1
	}
	v3 = p13
	t14 := int64(load64(m.memory[uint32(v2):]))
	t15 := v3
	t16 := v7
	v8 = t14
	var p17 int32
	if t16 < v8 {
		p17 = 1
	}
	p18 := t10
	if t15^p17 != 0 {
		p18 = t9
	}
	t19 := v3
	var p20 int32
	if v6 < v8 {
		p20 = 1
	}
	p21 := p18
	if t19^p20 != 0 {
		p21 = t8
	}
	return p21
}
func (m *Module) fn1013(v0 int64, v1, v2 int32) {
	var v3 int32
	var v4 int64
	var v5 int32
	t0 := int32(load32(m.memory[int64(uint32(v2))+8:]))
	v3 = t0
	t1 := int64(load64(m.memory[uint32(v3):]))
	v4 = t1
	t2 := int32(load32(m.memory[uint32(v2):]))
	t3 := int32(load32(m.memory[int64(uint32(v2))+12:]))
	t4 := v1
	v5 = t3
	v1 = t4 + v5<<3
	t5 := int64(load64(m.memory[uint32(v1):]))
	store64(m.memory[uint32(t2):], uint64(t5))
	store32(m.memory[uint32(v2):], uint32(v3))
	store32(m.memory[int64(uint32(v2))+8:], uint32(v3+i32(8)))
	t6 := v2
	t7 := v5
	var p8 int32
	if v0 >= v4 {
		p8 = 1
	}
	store32(m.memory[int64(uint32(t6))+12:], uint32(t7+p8))
	t9 := int64(load64(m.memory[uint32(v3):]))
	store64(m.memory[uint32(v1):], uint64(t9))
}
func (m *Module) fn1014(v0 int64, v1, v2 int32) {
	var v3 int32
	var v4 int64
	var v5 int32
	t0 := int32(load32(m.memory[int64(uint32(v2))+8:]))
	v3 = t0
	t1 := int64(load64(m.memory[uint32(v3):]))
	v4 = t1
	t2 := int32(load32(m.memory[uint32(v2):]))
	t3 := int32(load32(m.memory[int64(uint32(v2))+12:]))
	t4 := v1
	v5 = t3
	v1 = t4 + v5<<3
	t5 := int64(load64(m.memory[uint32(v1):]))
	store64(m.memory[uint32(t2):], uint64(t5))
	store32(m.memory[uint32(v2):], uint32(v3))
	store32(m.memory[int64(uint32(v2))+8:], uint32(v3+i32(8)))
	t6 := v2
	t7 := v5
	var p8 int32
	if v4 < v0 {
		p8 = 1
	}
	store32(m.memory[int64(uint32(t6))+12:], uint32(t7+p8))
	t9 := int64(load64(m.memory[uint32(v3):]))
	store64(m.memory[uint32(v1):], uint64(t9))
}
func (m *Module) fn1015(v0, v1, v2 int32) {
	var v3, v4, v5, v6 int32
	var v7 int64
	var v8, v9 int32
	var v10 int64
	var v11 int32
	var v12 int64
	t0 := m.g0
	v3 = t0 - i32(64)
	m.g0 = v3
	t1 := int32(load32(m.memory[int64(uint32(v2))+24:]))
	v4 = t1
	t2 := int32(load32(m.memory[int64(uint32(v2))+16:]))
	v5 = t2
	t3 := int32(load32(m.memory[int64(uint32(v2))+8:]))
	v6 = t3
	t4 := int64(load64(m.memory[uint32(v2):]))
	v7 = t4
	t5 := int32(load32(m.memory[int64(uint32(v2))+32:]))
	v8 = t5
	v9 = v0 + i32(4)
l4:
	if v7 == 0 {
		goto l0
	}
	{
		v2 = v5 + (i32(0)-int32(uint32(int64(bits.TrailingZeros64(uint64(v7))))>>3))*i32(40)
		t6 := int32(load32(m.memory[uint32(v2+i32(-12)):]))
		t7 := int32(load32(m.memory[uint32(v2+i32(-8)):]))
		t8 := m.fn773(t6, t7, i32(1073236), i32(73))
		if t8 == 0 {
			goto l1
		}
		t9 := int32(m.memory[uint32(v2+i32(-4))])
		if t9&i32(255) != 0 {
			goto l1
		}
		t10 := int32(load32(m.memory[int64(uint32(v8))+4:]))
		t11 := int32(load32(m.memory[int64(uint32(v8))+8:]))
		t12 := int32(load32(m.memory[uint32(v2+i32(-24)):]))
		t13 := int32(load32(m.memory[uint32(v2+i32(-20)):]))
		m.fn774(v3+i32(36), t10, t11, t12, t13)
		m.fn780(v3+i32(12), v3+i32(36))
		t14 := int32(load32(m.memory[int64(uint32(v3))+12:]))
		v2 = t14
		if v2 == i32(-1) {
			goto l1
		}
		t15 := int64(load64(m.memory[int64(uint32(v3))+16:]))
		v10 = t15
		t16 := int32(load32(m.memory[int64(uint32(v3))+24:]))
		t17 := int32(load32(m.memory[int64(uint32(v3))+28:]))
		m.fn134(t16, t17)
		store64(m.memory[int64(uint32(v3))+16:], uint64(v10))
		store32(m.memory[int64(uint32(v3))+12:], uint32(v2))
		v11 = int32(v10)
		{
			t18 := int32(load32(m.memory[int64(uint32(v0))+12:]))
			if t18 == 0 {
				goto l2
			}
			t19 := int64(load64(m.memory[int64(uint32(v0))+16:]))
			t20 := int64(load64(m.memory[int64(uint32(v0))+24:]))
			t21 := m.fn540(t19, t20, v11, int32(int64(uint64(v10)>>32)))
			v12 = t21
			t22 := int32(load32(m.memory[uint32(v0):]))
			t23 := int32(load32(m.memory[uint32(v9):]))
			t24 := m.fn644(t22, t23, v12, v3+i32(12))
			if t24 != 0 {
				goto l3
			}
		}
	l2:
		m.fn16(v2, v11)
		goto l1
	l3:
		store64(m.memory[int64(uint32(v3))+40:], uint64(v10))
		store32(m.memory[int64(uint32(v3))+36:], uint32(v2))
		_ = m.fn782(v1, v3+i32(36))
	}
l1:
	v7 = (v7 + i64(-1)) & v7
	v4 = v4 + i32(-1)
	goto l4
l0:
	{
		if v4 == 0 {
			goto l5
		}
		v5 = v5 + i32(-320)
		t26 := int64(load64(m.memory[uint32(v6):]))
		v7 = (t26 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
		v6 = v6 + i32(8)
		goto l4
	}
l5:
	m.g0 = v3 + i32(64)
}
func (m *Module) fn1016(v0, v1, v2 int32) int32 {
	t0 := m.fn1017(v1, v2)
	t1 := v0
	v2 = t0
	p2 := v2
	if uint32(v0) > uint32(v2) {
		p2 = t1
	}
	return p2
}
func (m *Module) fn1017(v0, v1 int32) int32 {
	var v2, v3, v4 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	v3 = v1 * i32(-12)
	v0 = v0 + v1*i32(12)
l2:
	{
		v4 = v1
		if v3 != 0 {
			goto l0
		}
		v4 = i32(0)
		goto l1
	l0:
		t1 := int32(load32(m.memory[uint32(v0+i32(-8)):]))
		t2 := int32(load32(m.memory[uint32(v0+i32(-4)):]))
		m.fn46(v2+i32(8), t1, t2)
		v3 = v3 + i32(12)
		v1 = v4 + i32(-1)
		v0 = v0 + i32(-12)
		t3 := int32(load32(m.memory[int64(uint32(v2))+12:]))
		if t3 == 0 {
			goto l2
		}
	}
l1:
	m.g0 = v2 + i32(16)
	return v4
}
func (m *Module) fn1018(v0 int64, v1, v2 int32) int64 {
	var v3 int64
	{
		if v1 == v2 {
			goto l0
		}
		t0 := int32(uint32(v2-v1) / uint32(i32(20)))
		v2 = t0
	l3:
		{
			{
				t1 := int32(load32(m.memory[uint32(v1):]))
				if t1 != i32(-1) {
					goto l1
				}
				v3 = i64(0)
				goto l2
			}
		l1:
			t2 := int32(load32(m.memory[uint32(v1+i32(4)):]))
			t3 := int32(load32(m.memory[uint32(v1+i32(8)):]))
			t4 := m.fn1019(t2, t3)
			v3 = t4
		}
	l2:
		v1 = v1 + i32(20)
		v0 = v3 + v0
		v2 = v2 + i32(-1)
		if v2 != 0 {
			goto l3
		}
	}
l0:
	return v0
}
func (m *Module) fn1019(v0, v1 int32) int64 {
	var v2 int64
	var v3 int32
	var v4 int64
	var v5, v6, v7, v8 int32
	v2 = i64(0)
	if v1 == 0 {
		goto l0
	}
	v3 = i32(0)
l10:
	v4 = i64(0)
	{
		v5 = v0 + v3<<5
		t0 := int32(load32(m.memory[uint32(v5):]))
		v6 = t0
		switch v6 >> 31 & (v6 + i32(-0x7fffffff)) {
		case 6:
			goto l7
		case 2:
			v4 = i64(0)
			t1 := int32(load32(m.memory[int64(uint32(v5))+24:]))
			v6 = t1
			if v6 == 0 {
				goto l7
			}
			t2 := int32(load32(m.memory[int64(uint32(v5))+20:]))
			v5 = t2 + i32(8)
			v4 = i64(0)
		l8:
			{
				t3 := int32(load32(m.memory[uint32(v5+i32(-4)):]))
				t4 := int32(load32(m.memory[uint32(v5):]))
				t5 := m.fn1019(t3, t4)
				v4 = t5 + v4
				v5 = v5 + i32(28)
				v6 = v6 + i32(-1)
				if v6 != 0 {
					goto l8
				}
				goto l7
			}
		case 3:
			v4 = i64(0)
			t6 := int32(load32(m.memory[int64(uint32(v5))+12:]))
			v6 = t6
			if v6 == 0 {
				goto l7
			}
			t7 := int32(load32(m.memory[int64(uint32(v5))+8:]))
			v5 = t7 + i32(8)
			v4 = i64(0)
		l9:
			{
				t8 := int32(load32(m.memory[uint32(v5):]))
				v7 = t8
				v8 = v5 + i32(-4)
				v5 = v5 + i32(12)
				t9 := int32(load32(m.memory[uint32(v8):]))
				t10 := v4
				v8 = t9
				t11 := m.fn1018(t10, v8, v8+v7*i32(20))
				v4 = t11
				v6 = v6 + i32(-1)
				if v6 != 0 {
					goto l9
				}
				goto l7
			}
		case 4:
			t12 := int32(load32(m.memory[int64(uint32(v5))+8:]))
			t13 := int32(load32(m.memory[int64(uint32(v5))+12:]))
			t14 := m.fn1019(t12, t13)
			v4 = t14
			goto l7
		case 5:
			t15 := int64(load32(m.memory[int64(uint32(v5))+12:]))
			v4 = t15
			goto l7
		case 1:
			v5 = v5 + i32(4)
			fallthrough
		default:
			t16 := int32(load32(m.memory[int64(uint32(v5))+4:]))
			t17 := int32(load32(m.memory[int64(uint32(v5))+8:]))
			t18 := m.fn1020(t16, t17)
			v4 = t18
		}
	}
l7:
	v2 = v4 + v2
	v3 = v3 + i32(1)
	if v3 != v1 {
		goto l10
	}
l0:
	return v2
}
func (m *Module) fn1020(v0, v1 int32) int64 {
	var v2, v3 int64
	var v4 int32
	v2 = i64(0)
	if v1 == 0 {
		goto l0
	}
l6:
	v3 = i64(1)
	{
		t0 := int32(load32(m.memory[uint32(v0):]))
		v4 = t0
		p1 := i32(1)
		if uint32(v4) > uint32(i32(2)) {
			p1 = v4 + i32(-3)
		}
		switch p1 {
		case 5:
			goto l5
		default:
			t2 := int64(load32(m.memory[uint32(v0+i32(12)):]))
			v3 = t2
			goto l5
		case 1:
			t3 := int64(load32(m.memory[uint32(v0+i32(12)):]))
			t4 := int32(load32(m.memory[uint32(v0+i32(20)):]))
			t5 := int32(load32(m.memory[uint32(v0+i32(24)):]))
			t6 := m.fn1020(t4, t5)
			v3 = t3 + t6
			goto l5
		case 2:
			t7 := int64(load32(m.memory[uint32(v0+i32(12)):]))
			v3 = t7
			goto l5
		case 3, 4:
			t8 := int64(load32(m.memory[uint32(v0+i32(12)):]))
			v3 = t8
		}
	}
l5:
	v0 = v0 + i32(28)
	v2 = v3 + v2
	v1 = v1 + i32(-1)
	if v1 != 0 {
		goto l6
	}
l0:
	return v2
}
func (m *Module) fn1021(v0, v1 int32) int32 {
	var v2 int32
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(t0):]))
	v2 = t1
	t2 := int32(load32(m.memory[uint32(v2+i32(4)):]))
	t3 := int32(load32(m.memory[uint32(v2+i32(8)):]))
	t4 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t5 := int32(load32(m.memory[uint32(t4):]))
	v0 = t5 - v1<<4
	t6 := int32(load32(m.memory[uint32(v0+i32(-12)):]))
	t7 := int32(load32(m.memory[uint32(v0+i32(-8)):]))
	t8 := m.fn544(t2, t3, t6, t7)
	return t8
}
func (m *Module) fn1022(v0, v1 int32) {
	var v2, v3 int32
	{
		{
			t0 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			v2 = t0
			t1 := int32(load32(m.memory[int64(uint32(v1))+16:]))
			t2 := v2
			v3 = t1
			if uint32(t2) >= uint32(v3) {
				goto l0
			}
			v1 = i32(0)
			goto l1
		}
	l0:
		store32(m.memory[int64(uint32(v1))+4:], uint32(v2-v3))
		t3 := int32(load32(m.memory[uint32(v1):]))
		t4 := v1
		v2 = t3
		store32(m.memory[uint32(t4):], uint32(v2+v3))
		switch v3 {
		case 0:
			m.fn158(i32(0), i32(0), i32(1073376))
			panic("unreachable")
		case 1:
			m.fn158(i32(1), i32(1), i32(1073392))
			panic("unreachable")
		default:
			t5 := int32(load16(m.memory[uint32(v2):]))
			v3 = t5
			v1 = i32(1)
		}
	}
l1:
	store16(m.memory[int64(uint32(v0))+2:], uint16(v3))
	store16(m.memory[uint32(v0):], uint16(v1))
}
func (m *Module) fn1023(v0, v1 int32) {
	var v2, v3, v4 int32
	var v5 int64
	var v6 int32
	{
		{
			t0 := int32(load32(m.memory[int64(uint32(v1))+24:]))
			v2 = t0
			if v2 != 0 {
				goto l0
			}
			v1 = i32(0)
			goto l1
		}
	l0:
		t1 := int32(load32(m.memory[int64(uint32(v1))+16:]))
		v3 = t1
		t2 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		v4 = t2
		t3 := int64(load64(m.memory[uint32(v1):]))
		v5 = t3
	l3:
		{
			if v5 != i64(0) {
				goto l2
			}
			t4 := v1
			v3 = v3 + i32(-64)
			store32(m.memory[int64(uint32(t4))+16:], uint32(v3))
			t5 := v1
			v6 = v4 + i32(8)
			store32(m.memory[int64(uint32(t5))+8:], uint32(v6))
			t6 := int64(load64(m.memory[uint32(v4):]))
			t7 := v1
			v5 = (t6 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
			store64(m.memory[uint32(t7):], uint64(v5))
			v4 = v6
			goto l3
		}
	l2:
		store32(m.memory[int64(uint32(v1))+24:], uint32(v2+i32(-1)))
		store64(m.memory[uint32(v1):], uint64((v5+i64(-1))&v5))
		t8 := int64(load64(m.memory[uint32(v3-int32(int64(bits.TrailingZeros64(uint64(v5))))&i32(120)+i32(-8)):]))
		store64(m.memory[int64(uint32(v0))+4:], uint64(t8))
		v1 = i32(1)
	}
l1:
	store32(m.memory[uint32(v0):], uint32(v1))
}
func (m *Module) fn1024(v0 int32) {
	var v1 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+32:]))
		v1 = t0
		if v1 == 0 {
			return
		}
		t1 := int32(load32(m.memory[int64(uint32(v0))+40:]))
		t2 := int32(load32(m.memory[int64(uint32(v0))+36:]))
		m.fn40(t1, v1, t2)
	}
}
func (m *Module) fn1025(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(t0):]))
	t2 := int32(load32(m.memory[uint32(t1):]))
	t3 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t4 := int32(load32(m.memory[uint32(t3):]))
	t5 := int32(load32(m.memory[uint32(t4+(i32(0)-v1)*i32(12)+i32(-12)):]))
	var p6 int32
	if t2 == t5 {
		p6 = 1
	}
	return p6
}
func (m *Module) fn1026(v0, v1 int32) {
	var v2, v3, v4, v5, v6, v7, v8 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	t1 := int32(load32(m.memory[int64(uint32(v1))+96:]))
	v3 = t1
	t2 := int32(load32(m.memory[int64(uint32(v1))+88:]))
	v4 = t2
	v5 = i32(0)
	t3 := int32(load32(m.memory[int64(uint32(v1))+56:]))
	v6 = t3
	t4 := int32(load32(m.memory[int64(uint32(v1))+48:]))
	v7 = t4
	{
		{
			t5 := int32(load32(m.memory[uint32(v1):]))
			if t5 == 0 {
				goto l0
			}
			v8 = v2 + i32(12)
			t6 := int32(load32(m.memory[int64(uint32(v1))+20:]))
			t7 := int32(load32(m.memory[int64(uint32(v1))+16:]))
			if t6 != t7 {
				goto l1
			}
		}
	l0:
		t9 := v2
		p8 := i32(0)
		if v3 != 0 {
			p8 = v4
		}
		p10 := i32(0)
		if v6 != 0 {
			p10 = v7
		}
		v1 = p10
		v5 = p8 + v1
		var p11 int32
		if uint32(v5) >= uint32(v1) {
			p11 = 1
		}
		store32(m.memory[int64(uint32(t9))+12:], uint32(p11))
		v8 = v2 + i32(8)
	}
l1:
	store32(m.memory[uint32(v8):], uint32(v5))
	t12 := int32(load32(m.memory[int64(uint32(v2))+12:]))
	v1 = t12
	t13 := int32(load32(m.memory[int64(uint32(v2))+8:]))
	store32(m.memory[int64(uint32(v0))+8:], uint32(t13))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
	store32(m.memory[uint32(v0):], uint32(i32(0)))
}
func (m *Module) fn1027(v0 int32) {
	{
		t0 := int32(load32(m.memory[uint32(v0):]))
		if t0 != 0 {
			goto l0
		}
		m.fn956(v0 + i32(4))
		return
	}
l0:
	t1 := int32(m.memory[int64(uint32(v0))+4])
	t2 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	m.fn119(t1, t2)
}
func (m *Module) fn1028(v0, v1, v2 int32) {
	var v3 int32
	t0 := m.g0
	v3 = t0 - i32(16)
	m.g0 = v3
	m.fn1623(v3+i32(4), i32(1101983), i32(1))
	v2 = v2 << 3
l1:
	{
		if v2 == 0 {
			goto l0
		}
		t1 := int32(load32(m.memory[uint32(v1):]))
		t2 := int32(load32(m.memory[uint32(v1+i32(4)):]))
		m.fn1030(v3+i32(4), t1, t2)
		v2 = v2 + i32(-8)
		v1 = v1 + i32(8)
		goto l1
	}
l0:
	t3 := int32(load32(m.memory[int64(uint32(v3))+12:]))
	store32(m.memory[int64(uint32(v0))+8:], uint32(t3))
	t4 := int64(load64(m.memory[int64(uint32(v3))+4:]))
	store64(m.memory[uint32(v0):], uint64(t4))
	m.g0 = v3 + i32(16)
}
func (m *Module) fn1029(v0, v1, v2, v3 int32) {
	var v4, v5, v6, v7, v8 int32
	t0 := m.g0
	v4 = t0 - i32(32)
	m.g0 = v4
	t1 := int32(load32(m.memory[int64(uint32(v0))+12:]))
	t2 := int32(load32(m.memory[uint32(t1):]))
	m.fn642(v4+i32(12), t2+i32(8))
	t3 := int32(load32(m.memory[int64(uint32(v4))+20:]))
	v5 = t3
	t4 := int32(load32(m.memory[int64(uint32(v4))+16:]))
	v6 = t4
	v7 = v6 + i32(84)
	v8 = v6 + i32(80)
l2:
	{
		if v3 == i32(-1) {
			goto l0
		}
		m.fn1031(v4+i32(12), v1, v2)
		m.memory[int64(uint32(v4))+28] = byte(i32(1))
		store32(m.memory[int64(uint32(v4))+24:], uint32(v3))
		{
			t5 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			v6 = t5
			t6 := int32(load32(m.memory[uint32(v0):]))
			if v6 != t6 {
				goto l1
			}
			m.fn418(v0)
		}
	l1:
		store32(m.memory[int64(uint32(v0))+8:], uint32(v6+i32(1)))
		t7 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		v6 = t7 + v6*i32(20)
		t8 := int64(load64(m.memory[int64(uint32(v4))+12:]))
		store64(m.memory[uint32(v6):], uint64(t8))
		t9 := int64(load64(m.memory[int64(uint32(v4))+20:]))
		store64(m.memory[int64(uint32(v6))+8:], uint64(t9))
		t10 := int32(load32(m.memory[int64(uint32(v4))+28:]))
		store32(m.memory[int64(uint32(v6))+16:], uint32(t10))
		t11 := int32(load32(m.memory[uint32(v8):]))
		t12 := int32(load32(m.memory[uint32(v7):]))
		t13 := m.fn590(t11, t12, v3)
		t14 := int32(load32(m.memory[int64(uint32(t13))+40:]))
		v3 = t14
		goto l2
	}
l0:
	m.fn641(v5)
	m.g0 = v4 + i32(32)
}
func (m *Module) fn1030(v0, v1, v2 int32) {
	var v3, v4, v5 int32
	{
		{
			t0 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			v3 = t0
			if v3 != 0 {
				goto l0
			}
			v4 = i32(0)
			goto l1
		}
	l0:
		t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		t2 := int32(m.memory[uint32(t1+v3+i32(-1))])
		var p3 int32
		if t2 != i32(47) {
			p3 = 1
		}
		v4 = p3
	}
l1:
	{
		{
			{
				{
					{
						if v2 == 0 {
							goto l2
						}
						t4 := int32(m.memory[uint32(v1)])
						if t4 == i32(47) {
							goto l3
						}
					}
				l2:
					t5 := int32(load32(m.memory[uint32(v0):]))
					v5 = t5
					{
						if v4 != 0 {
							goto l4
						}
						v4 = v3
						goto l5
					l4:
						{
							if v5 != v3 {
								goto l6
							}
							m.fn1799(v0, v3, i32(1))
							t6 := int32(load32(m.memory[uint32(v0):]))
							v5 = t6
							t7 := int32(load32(m.memory[int64(uint32(v0))+8:]))
							v3 = t7
						}
					l6:
						t8 := v0
						v4 = v3 + i32(1)
						store32(m.memory[int64(uint32(t8))+8:], uint32(v4))
						t9 := int32(load32(m.memory[int64(uint32(v0))+4:]))
						m.memory[uint32(t9+v3)] = byte(i32(47))
					}
				l5:
					if uint32(v2) > uint32(v5-v4) {
						goto l7
					}
					if v2 != 0 {
						goto l8
					}
					goto l9
				}
			l3:
				v4 = i32(0)
				store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
				t10 := int32(load32(m.memory[uint32(v0):]))
				if uint32(v2) <= uint32(t10) {
					goto l8
				}
			}
		l7:
			m.fn1799(v0, v4, v2)
			t11 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			v4 = t11
		}
	l8:
		if v2 == 0 {
			goto l9
		}
		t12 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		memory_copy(m.memory, uint32(t12+v4), uint32(v1), uint32(v2))
	}
l9:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v4+v2))
}
func (m *Module) fn1031(v0, v1, v2 int32) {
	var v3 int32
	{
		if v2 != 0 {
			goto l0
		}
		v3 = i32(1)
		goto l1
	l0:
		t0 := m.fn4(v2)
		v3 = t0
		if v3 == 0 {
			m.fn2(i32(1), v2)
			panic("unreachable")
		}
		if v2 == 0 {
			goto l1
		}
		memory_copy(m.memory, uint32(v3), uint32(v1), uint32(v2))
	}
l1:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v2))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
	store32(m.memory[uint32(v0):], uint32(v2))
}
