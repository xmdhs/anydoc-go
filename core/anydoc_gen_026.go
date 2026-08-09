package core

import (
	"math"
)

func (m *Module) fn1122(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(t0):]))
	t2 := int32(load32(m.memory[uint32(t1):]))
	t3 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t4 := int32(load32(m.memory[uint32(t3):]))
	t5 := int32(load32(m.memory[uint32(t4-v1<<4+i32(-16)):]))
	var p6 int32
	if t2 == t5 {
		p6 = 1
	}
	return p6
}
func (m *Module) fn1123(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(t0):]))
	t2 := int32(load32(m.memory[uint32(t1):]))
	t3 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t4 := int32(load32(m.memory[uint32(t3):]))
	t5 := int32(load32(m.memory[uint32(t4-v1<<3+i32(-8)):]))
	var p6 int32
	if t2 == t5 {
		p6 = 1
	}
	return p6
}
func (m *Module) fn1124(v0, v1, v2, v3 int32) {
	var v4 int32
	var v5 int64
	var v6, v7, v8, v9 int32
	t0 := m.g0
	v4 = t0 - i32(320)
	m.g0 = v4
	store32(m.memory[int64(uint32(v4))+20:], uint32(v2))
	t1 := int64(load64(m.memory[int64(uint32(v1))+16:]))
	t2 := int64(load64(m.memory[int64(uint32(v1))+24:]))
	t3 := m.fn66(t1, t2, v2)
	v5 = t3
	store32(m.memory[int64(uint32(v4))+316:], uint32(v4+i32(20)))
	{
		t4 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		if t4 != 0 {
			goto l0
		}
		_ = m.fn715(v1, v1+i32(16))
	}
l0:
	store32(m.memory[int64(uint32(v4))+24:], uint32(v4+i32(316)))
	store32(m.memory[int64(uint32(v4))+28:], uint32(v1))
	t6 := int32(load32(m.memory[uint32(v1):]))
	t7 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	m.fn69(v4+i32(8), t6, t7, v5, v4+i32(24), i32(160))
	t8 := int32(load32(m.memory[uint32(v1):]))
	v6 = t8
	t9 := int32(load32(m.memory[int64(uint32(v4))+12:]))
	v7 = t9
	{
		{
			t10 := int32(load32(m.memory[int64(uint32(v4))+8:]))
			if t10 != i32(1) {
				goto l1
			}
			memory_copy(m.memory, uint32(v4+i32(28)), uint32(v3), uint32(i32(288)))
			v3 = v6 + v7
			t11 := int32(m.memory[uint32(v3)])
			v8 = t11
			t12 := v3
			v9 = int32(uint32(int32(v5)) >> 25)
			m.memory[uint32(t12)] = byte(v9)
			t13 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			m.memory[uint32(v6+t13&(v7+i32(-8))+i32(8))] = byte(v9)
			t14 := int32(load32(m.memory[int64(uint32(v1))+12:]))
			store32(m.memory[int64(uint32(v1))+12:], uint32(t14+i32(1)))
			t15 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			store32(m.memory[int64(uint32(v1))+8:], uint32(t15-v8&i32(1)))
			v1 = v6 + (i32(0)-v7)*i32(296)
			store32(m.memory[uint32(v1+i32(-296)):], uint32(v2))
			memory_copy(m.memory, uint32(v1+i32(-292)), uint32(v4+i32(24)), uint32(i32(292)))
			store32(m.memory[uint32(v0):], uint32(i32(-1)))
			goto l2
		}
	l1:
		t16 := v0
		v1 = v6 + (i32(0)-v7)*i32(296) + i32(-288)
		memory_copy(m.memory, uint32(t16), uint32(v1), uint32(i32(288)))
		memory_copy(m.memory, uint32(v1), uint32(v3), uint32(i32(288)))
	}
l2:
	m.g0 = v4 + i32(320)
}
func (m *Module) fn1125(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(t0):]))
	t2 := int32(load32(m.memory[uint32(t1):]))
	t3 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t4 := int32(load32(m.memory[uint32(t3):]))
	t5 := int32(load32(m.memory[uint32(t4+(i32(0)-v1)*i32(296)+i32(-296)):]))
	var p6 int32
	if t2 == t5 {
		p6 = 1
	}
	return p6
}
func (m *Module) fn1126(v0, v1 int32) int32 {
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
func (m *Module) fn1127(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(t0):]))
	t2 := int32(load32(m.memory[uint32(t1):]))
	t3 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t4 := int32(load32(m.memory[uint32(t3):]))
	t5 := int32(load32(m.memory[uint32(t4+(i32(0)-v1)*i32(20)+i32(-20)):]))
	var p6 int32
	if t2 == t5 {
		p6 = 1
	}
	return p6
}
func (m *Module) fn1128(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(t0):]))
	t2 := int32(load32(m.memory[uint32(t1):]))
	t3 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t4 := int32(load32(m.memory[uint32(t3):]))
	t5 := int32(load32(m.memory[uint32(t4+(i32(0)-v1)*i32(368)+i32(-368)):]))
	var p6 int32
	if t2 == t5 {
		p6 = 1
	}
	return p6
}
func (m *Module) fn1129(v0, v1, v2, v3 int32) {
	var v4 int32
	var v5 int64
	var v6, v7, v8, v9 int32
	t0 := m.g0
	v4 = t0 - i32(80)
	m.g0 = v4
	store16(m.memory[int64(uint32(v4))+14:], uint16(v2))
	t1 := int64(load64(m.memory[int64(uint32(v1))+16:]))
	t2 := int64(load64(m.memory[int64(uint32(v1))+24:]))
	t3 := m.fn529(t1, t2, v2)
	v5 = t3
	store32(m.memory[int64(uint32(v4))+76:], uint32(v4+i32(14)))
	{
		t4 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		if t4 != 0 {
			goto l0
		}
		_ = m.fn731(v1, v1+i32(16))
	}
l0:
	store32(m.memory[int64(uint32(v4))+16:], uint32(v4+i32(76)))
	store32(m.memory[int64(uint32(v4))+20:], uint32(v1))
	t6 := int32(load32(m.memory[uint32(v1):]))
	t7 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	m.fn69(v4, t6, t7, v5, v4+i32(16), i32(161))
	t8 := int32(load32(m.memory[uint32(v1):]))
	v6 = t8
	t9 := int32(load32(m.memory[int64(uint32(v4))+4:]))
	v7 = t9
	{
		{
			t10 := int32(load32(m.memory[uint32(v4):]))
			if t10 != i32(1) {
				goto l1
			}
			memory_copy(m.memory, uint32(v4+i32(16)+i32(2)), uint32(v3), uint32(i32(56)))
			v3 = v6 + v7
			t11 := int32(m.memory[uint32(v3)])
			v8 = t11
			t12 := v3
			v9 = int32(uint32(int32(v5)) >> 25)
			m.memory[uint32(t12)] = byte(v9)
			t13 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			m.memory[uint32(v6+t13&(v7+i32(-8))+i32(8))] = byte(v9)
			t14 := int32(load32(m.memory[int64(uint32(v1))+12:]))
			store32(m.memory[int64(uint32(v1))+12:], uint32(t14+i32(1)))
			t15 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			store32(m.memory[int64(uint32(v1))+8:], uint32(t15-v8&i32(1)))
			v1 = v6 + (i32(0)-v7)*i32(60)
			store16(m.memory[uint32(v1+i32(-60)):], uint16(v2))
			memory_copy(m.memory, uint32(v1+i32(-58)), uint32(v4+i32(16)), uint32(i32(58)))
			store32(m.memory[uint32(v0):], uint32(i32(2)))
			goto l2
		}
	l1:
		t16 := v0
		v1 = v6 + (i32(0)-v7)*i32(60) + i32(-56)
		memory_copy(m.memory, uint32(t16), uint32(v1), uint32(i32(56)))
		memory_copy(m.memory, uint32(v1), uint32(v3), uint32(i32(56)))
	}
l2:
	m.g0 = v4 + i32(80)
}
func (m *Module) fn1130(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(t0):]))
	t2 := int32(load16(m.memory[uint32(t1):]))
	t3 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t4 := int32(load32(m.memory[uint32(t3):]))
	t5 := int32(load16(m.memory[uint32(t4+(i32(0)-v1)*i32(60)+i32(-60)):]))
	var p6 int32
	if t2 == t5 {
		p6 = 1
	}
	return p6
}
func (m *Module) fn1131(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(t0):]))
	t2 := int32(load16(m.memory[uint32(t1):]))
	t3 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t4 := int32(load32(m.memory[uint32(t3):]))
	t5 := int32(load16(m.memory[uint32(t4+(i32(0)-v1)*i32(36)+i32(-36)):]))
	var p6 int32
	if t2 == t5 {
		p6 = 1
	}
	return p6
}
func (m *Module) fn1132(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(t0):]))
	t2 := int32(load16(m.memory[uint32(t1):]))
	t3 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t4 := int32(load32(m.memory[uint32(t3):]))
	t5 := int32(load16(m.memory[uint32(t4+(i32(0)-v1)*i32(520)+i32(-520)):]))
	var p6 int32
	if t2 == t5 {
		p6 = 1
	}
	return p6
}
func (m *Module) fn1133(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(t0):]))
	t2 := int32(load16(m.memory[uint32(t1):]))
	t3 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t4 := int32(load32(m.memory[uint32(t3):]))
	t5 := int32(load16(m.memory[uint32(t4-v1<<1+i32(-2)):]))
	var p6 int32
	if t2 == t5 {
		p6 = 1
	}
	return p6
}
func (m *Module) fn1134(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(t0):]))
	t2 := int64(load64(m.memory[uint32(t1):]))
	t3 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t4 := int32(load32(m.memory[uint32(t3):]))
	t5 := int64(load64(m.memory[uint32(t4+(i32(0)-v1)*i32(480)+i32(-480)):]))
	var p6 int32
	if t2 == t5 {
		p6 = 1
	}
	return p6
}
func (m *Module) fn1135(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(t0):]))
	t2 := int64(load64(m.memory[uint32(t1):]))
	t3 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t4 := int32(load32(m.memory[uint32(t3):]))
	t5 := int64(load64(m.memory[uint32(t4+(i32(0)-v1)*i32(24)+i32(-24)):]))
	var p6 int32
	if t2 == t5 {
		p6 = 1
	}
	return p6
}
func (m *Module) fn1136(v0, v1 int32) {
	t0 := m.fn934()
	store32(m.memory[int64(uint32(v0))+4:], uint32(t0))
	store32(m.memory[uint32(v0):], uint32(v1))
}
func (m *Module) fn1137(v0, v1 int32) {
	var v2, v3, v4, v5 int32
	var v6 int64
	t0 := m.g0
	v2 = t0 - i32(352)
	m.g0 = v2
	v3 = i32(0)
l1:
	{
		if v3 == i32(288) {
			goto l0
		}
		v4 = v1 + v3
		t1 := int32(m.memory[uint32(v4+i32(24))])
		v5 = t1
		t2 := int64(load64(m.memory[uint32(v4+i32(16)):]))
		v6 = t2
		m.fn1138(v2+i32(320), v4)
		v4 = v2 + v3
		store64(m.memory[int64(uint32(v4))+16:], uint64(v6))
		t3 := int64(load64(m.memory[int64(uint32(v2))+320:]))
		store64(m.memory[uint32(v4):], uint64(t3))
		t4 := int64(load64(m.memory[int64(uint32(v2))+328:]))
		store64(m.memory[int64(uint32(v4))+8:], uint64(t4))
		m.memory[int64(uint32(v2))+344] = byte(v5)
		t5 := int64(load64(m.memory[int64(uint32(v2))+344:]))
		store64(m.memory[int64(uint32(v4))+24:], uint64(t5))
		v3 = v3 + i32(32)
		goto l1
	}
l0:
	memory_copy(m.memory, uint32(v0), uint32(v2), uint32(i32(288)))
	m.g0 = v2 + i32(352)
}
func (m *Module) fn1138(v0, v1 int32) {
	var v2, v3, v4, v5, v6, v7, v8, v9, v10 int32
	t0 := m.g0
	v2 = t0 - i32(32)
	m.g0 = v2
	t1 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	v3 = t1
	t2 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	t3 := v2 + i32(8)
	v4 = t2
	m.fn59(t3, v4, i32(4), i32(12))
	v5 = v4 * i32(12)
	v6 = i32(0)
	t4 := int32(load32(m.memory[int64(uint32(v2))+12:]))
	v7 = t4
	t5 := int32(load32(m.memory[int64(uint32(v2))+8:]))
	v8 = t5
	v9 = v8
l3:
	{
		if v9 == 0 {
			goto l0
		}
		if v5 == v6 {
			goto l0
		}
		{
			{
				v10 = v3 + v6
				t6 := int32(load32(m.memory[uint32(v10):]))
				if t6 != i32(-1) {
					goto l1
				}
				t7 := int32(load32(m.memory[int64(uint32(v10))+8:]))
				store32(m.memory[int64(uint32(v2))+24:], uint32(t7))
				t8 := int64(load64(m.memory[uint32(v10):]))
				store64(m.memory[int64(uint32(v2))+16:], uint64(t8))
				goto l2
			}
		l1:
			t9 := int32(load32(m.memory[uint32(v10+i32(4)):]))
			t10 := int32(load32(m.memory[uint32(v10+i32(8)):]))
			m.fn31(v2+i32(16), t9, t10)
		}
	l2:
		v10 = v7 + v6
		t11 := int32(load32(m.memory[int64(uint32(v2))+24:]))
		store32(m.memory[int64(uint32(v10))+8:], uint32(t11))
		t12 := int64(load64(m.memory[int64(uint32(v2))+16:]))
		store64(m.memory[uint32(v10):], uint64(t12))
		v9 = v9 + i32(-1)
		v6 = v6 + i32(12)
		goto l3
	}
l0:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v4))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v7))
	store32(m.memory[uint32(v0):], uint32(v8))
	t13 := int32(m.memory[int64(uint32(v1))+12])
	m.memory[int64(uint32(v0))+12] = byte(t13)
	m.g0 = v2 + i32(32)
}
func (m *Module) fn1139(v0, v1 int32) {
	var v2, v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14, v15, v16, v17, v18, v19, v20 int32
	var v21, v22 int64
	var v23, v24, v25, v26, v27, v28, v29, v30, v31, v32, v33, v34, v35, v36, v37, v38, v39, v40, v41, v42 int32
	t0 := m.g0
	v2 = t0 - i32(336)
	m.g0 = v2
	v3 = v1 + i32(328)
	v4 = v2 + i32(304) + i32(8)
	v5 = v2 + i32(108) + i32(8)
	v6 = v2 + i32(112)
	{
		{
		l25:
			{
				store32(m.memory[int64(uint32(v1))+336:], uint32(i32(0)))
				m.fn141(v2+i32(108), v1, v3)
				{
					t1 := int32(load32(m.memory[int64(uint32(v2))+108:]))
					if t1 != i32(1) {
						goto l0
					}
					m.memory[uint32(v0)] = byte(i32(254))
					t2 := int64(load64(m.memory[int64(uint32(v6))+16:]))
					store64(m.memory[int64(uint32(v0))+20:], uint64(t2))
					t3 := int64(load64(m.memory[int64(uint32(v6))+8:]))
					store64(m.memory[int64(uint32(v0))+12:], uint64(t3))
					t4 := int64(load64(m.memory[uint32(v6):]))
					store64(m.memory[int64(uint32(v0))+4:], uint64(t4))
					goto l1
				}
			l0:
				{
					{
						{
							{
								t5 := int32(load32(m.memory[int64(uint32(v2))+112:]))
								v7 = t5
								switch v7 {
								case 1:
									goto l3
								case 0:
									m.fn551(v2+i32(80), v5)
									{
										{
											t6 := int32(load32(m.memory[int64(uint32(v2))+84:]))
											if t6 != i32(3) {
												goto l5
											}
											t7 := int32(load32(m.memory[int64(uint32(v2))+80:]))
											v7 = t7
											t8 := int32(load16(m.memory[uint32(v7):]))
											t9 := int32(m.memory[uint32(v7+i32(2))])
											if t8|t9<<16 == i32(7827314) {
												t26 := int32(load32(m.memory[int64(uint32(v2))+120:]))
												v7 = t26
												t27 := int32(load32(m.memory[int64(uint32(v2))+116:]))
												v16 = t27
												m.fn165(v2+i32(304), v5, i32(1072195), i32(1))
												{
													{
														t28 := int32(m.memory[int64(uint32(v2))+304])
														v15 = t28
														if v15 == i32(255) {
															goto l21
														}
														t29 := int32(m.memory[int64(uint32(v2))+307])
														t30 := v2
														v1 = t29
														m.memory[int64(uint32(t30))+162] = byte(v1)
														t31 := int32(load16(m.memory[int64(uint32(v2))+305:]))
														t32 := v2
														v3 = t31
														store16(m.memory[int64(uint32(t32))+160:], uint16(v3))
														t33 := int64(load64(m.memory[int64(uint32(v2))+308:]))
														v21 = t33
														m.memory[int64(uint32(v0))+8] = byte(v15)
														store16(m.memory[int64(uint32(v0))+9:], uint16(v3))
														m.memory[int64(uint32(v0))+11] = byte(v1)
														store64(m.memory[int64(uint32(v0))+12:], uint64(v21))
														store32(m.memory[int64(uint32(v0))+4:], uint32(i32(-0x7fffffed)))
														m.memory[uint32(v0)] = byte(i32(254))
														goto l22
													}
												l21:
													{
														t34 := int32(load32(m.memory[int64(uint32(v2))+308:]))
														v15 = t34
														if v15 == 0 {
															goto l23
														}
														t35 := int32(load32(m.memory[int64(uint32(v2))+312:]))
														m.fn1141(v2+i32(304), v15, t35)
														t36 := int32(load32(m.memory[int64(uint32(v2))+308:]))
														v15 = t36
														t37 := int32(load32(m.memory[int64(uint32(v2))+304:]))
														v13 = t37
														if v13 != i32(-1) {
															goto l24
														}
														store32(m.memory[int64(uint32(v1))+400:], uint32(v15))
													}
												l23:
													m.fn134(v16, v7)
													goto l25
												l24:
													t38 := int64(load64(m.memory[int64(uint32(v4))+8:]))
													t39 := v2
													v21 = t38
													store64(m.memory[int64(uint32(t39))+168:], uint64(v21))
													t40 := int64(load64(m.memory[uint32(v4):]))
													t41 := v2
													v22 = t40
													store64(m.memory[int64(uint32(t41))+160:], uint64(v22))
													store64(m.memory[int64(uint32(v0))+20:], uint64(v21))
													store64(m.memory[int64(uint32(v0))+12:], uint64(v22))
													store32(m.memory[int64(uint32(v0))+8:], uint32(v15))
													store32(m.memory[int64(uint32(v0))+4:], uint32(v13))
													m.memory[uint32(v0)] = byte(i32(254))
												}
											l22:
												m.fn134(v16, v7)
												v1 = i32(0)
												goto l26
											}
										}
									l5:
										m.fn551(v2+i32(72), v5)
										t10 := int32(load32(m.memory[int64(uint32(v2))+76:]))
										if t10 != i32(1) {
											goto l7
										}
										t11 := int32(load32(m.memory[int64(uint32(v2))+72:]))
										t12 := int32(m.memory[uint32(t11)])
										if t12 != i32(99) {
											goto l7
										}
										t13 := int32(load32(m.memory[int64(uint32(v2))+120:]))
										v8 = t13
										t14 := int32(load32(m.memory[int64(uint32(v2))+116:]))
										v9 = t14
										m.fn166(v2+i32(160), v5)
										v10 = i32(0)
										v11 = i32(0)
										v12 = i32(0)
										v13 = i32(0)
										{
										l11:
											{
												m.fn167(v2+i32(304), v2+i32(160))
												{
													t15 := int32(load32(m.memory[int64(uint32(v2))+304:]))
													if t15 == i32(1) {
														goto l8
													}
													v4 = v14
													v15 = v10
													goto l9
												}
											l8:
												t16 := int32(load32(m.memory[int64(uint32(v2))+320:]))
												v3 = t16
												t17 := int32(load32(m.memory[int64(uint32(v2))+316:]))
												v5 = t17
												t18 := int32(load32(m.memory[int64(uint32(v2))+312:]))
												v7 = t18
												{
													t19 := int32(load32(m.memory[int64(uint32(v2))+308:]))
													v16 = t19
													if v16 == 0 {
														goto l10
													}
													if v7 != i32(1) {
														goto l11
													}
													v4 = v3
													v15 = v5
													{
														t20 := int32(m.memory[uint32(v16)])
														switch t20 + i32(-114) {
														case 0:
															goto l12
														default:
															goto l11
														case 1:
															v4 = v14
															v15 = v10
															v17 = v3
															v11 = v5
															goto l12
														case 2:
															v4 = v14
															v15 = v10
															v18 = v3
															v12 = v5
														}
													}
												l12:
													v10 = v15
													v14 = v4
													v13 = v13 + i32(1)
													if v13&i32(255) != i32(3) {
														goto l11
													}
													goto l9
												}
											l10:
											}
											v15 = v10
											v4 = v14
											if v7&i32(255) != i32(255) {
												goto l15
											}
										l9:
											if v15 == 0 {
												goto l16
											}
											m.fn1140(v2+i32(304), v15, v4)
											t21 := int32(load32(m.memory[int64(uint32(v2))+312:]))
											v19 = t21
											t22 := int32(load32(m.memory[int64(uint32(v2))+308:]))
											v20 = t22
											{
												t23 := int32(load32(m.memory[int64(uint32(v2))+304:]))
												v7 = t23
												if v7 == i32(-1) {
													store32(m.memory[int64(uint32(v1))+404:], uint32(v19))
													goto l19
												}
												t24 := int32(load32(m.memory[int64(uint32(v2))+324:]))
												store32(m.memory[int64(uint32(v0))+24:], uint32(t24))
												t25 := int64(load64(m.memory[int64(uint32(v2))+316:]))
												store64(m.memory[int64(uint32(v0))+16:], uint64(t25))
												store32(m.memory[int64(uint32(v0))+12:], uint32(v19))
												store32(m.memory[int64(uint32(v0))+8:], uint32(v20))
												store32(m.memory[int64(uint32(v0))+4:], uint32(v7))
												m.memory[uint32(v0)] = byte(i32(254))
												goto l18
											}
										}
									l15:
										store32(m.memory[int64(uint32(v0))+16:], uint32(v3))
										store32(m.memory[int64(uint32(v0))+12:], uint32(v5))
										store32(m.memory[int64(uint32(v0))+8:], uint32(v7))
										store32(m.memory[int64(uint32(v0))+4:], uint32(i32(-0x7fffffed)))
										m.memory[uint32(v0)] = byte(i32(254))
									l18:
										v1 = i32(0)
										goto l20
									}
								default:
									if v7 == i32(10) {
										store32(m.memory[int64(uint32(v0))+12:], uint32(i32(9)))
										store32(m.memory[int64(uint32(v0))+8:], uint32(i32(1077933)))
										store32(m.memory[int64(uint32(v0))+4:], uint32(i32(-0x7fffffe9)))
										m.memory[uint32(v0)] = byte(i32(254))
										v1 = i32(0)
										goto l26
									}
									m.fn200(v6)
									goto l25
								}
							}
						l7:
							t42 := int32(load32(m.memory[int64(uint32(v2))+116:]))
							t43 := int32(load32(m.memory[int64(uint32(v2))+120:]))
							m.fn134(t42, t43)
							goto l25
						}
					l16:
						t44 := int32(load32(m.memory[int64(uint32(v1))+404:]))
						v19 = t44
						t45 := int32(load32(m.memory[int64(uint32(v1))+400:]))
						v20 = t45
					}
				l19:
					m.memory[int64(uint32(v2))+136] = byte(i32(9))
					v23 = int32(uint32(i32(1072448)) >> 8)
					v24 = v1 + i32(376)
					v25 = int32(uint32(i32(1072447)) >> 8)
					v26 = v1 + i32(364)
					v27 = v1 + i32(352)
					v4 = v1 + i32(340)
					v10 = v2 + i32(136) + i32(8)
					v28 = v2 + i32(304) + i32(16)
					v29 = v2 + i32(240) + i32(8)
					v30 = v2 + i32(304) + i32(12)
					v31 = v2 + i32(264) + i32(12)
					v32 = v2 + i32(160) + i32(8)
					v33 = v2 + i32(164)
					v34 = v2 + i32(315)
					{
					l77:
						store32(m.memory[int64(uint32(v1))+348:], uint32(i32(0)))
						m.fn141(v2+i32(160), v1, v4)
						{
							t46 := int32(load32(m.memory[int64(uint32(v2))+160:]))
							if t46 != i32(1) {
								goto l28
							}
							m.memory[uint32(v0)] = byte(i32(254))
							t47 := int64(load64(m.memory[int64(uint32(v33))+16:]))
							store64(m.memory[int64(uint32(v0))+20:], uint64(t47))
							t48 := int64(load64(m.memory[int64(uint32(v33))+8:]))
							store64(m.memory[int64(uint32(v0))+12:], uint64(t48))
							t49 := int64(load64(m.memory[uint32(v33):]))
							store64(m.memory[int64(uint32(v0))+4:], uint64(t49))
							goto l29
						}
					l28:
						{
							{
								{
									{
										t50 := int32(load32(m.memory[int64(uint32(v2))+164:]))
										v5 = t50
										switch v5 {
										case 0:
											t51 := int32(load32(m.memory[int64(uint32(v2))+172:]))
											v14 = t51
											t52 := int32(load32(m.memory[int64(uint32(v2))+168:]))
											v35 = t52
											t53 := int32(m.memory[int64(uint32(v1))+408])
											m.memory[int64(uint32(v2))+204] = byte(t53)
											t54 := int64(load64(m.memory[int64(uint32(v1))+304:]))
											store64(m.memory[int64(uint32(v2))+196:], uint64(t54))
											t55 := int64(load64(m.memory[int64(uint32(v1))+296:]))
											store64(m.memory[int64(uint32(v2))+188:], uint64(t55))
											m.fn551(v2+i32(56), v32)
											v13 = i32(1)
											v7 = i32(11)
											v3 = i32(-0x7fffffe8)
											v16 = i32(1072448)
											t56 := int32(load32(m.memory[int64(uint32(v2))+56:]))
											v36 = t56
											v15 = v23
											{
												t57 := int32(load32(m.memory[int64(uint32(v2))+60:]))
												switch t57 + i32(-1) {
												default:
													goto l35
												case 1:
													t58 := int32(m.memory[uint32(v36)])
													if t58 != i32(105) {
														goto l36
													}
													t59 := int32(m.memory[int64(uint32(v36))+1])
													if t59 != i32(115) {
														goto l36
													}
													m.fn164(v2+i32(8), v32)
													t60 := int32(load32(m.memory[int64(uint32(v2))+8:]))
													t61 := int32(load32(m.memory[int64(uint32(v2))+12:]))
													m.fn557(v2+i32(304), v1, t60, t61, v27, v24)
													t62 := int32(load32(m.memory[int64(uint32(v2))+308:]))
													v7 = t62
													{
														t63 := int32(load32(m.memory[int64(uint32(v2))+304:]))
														v3 = t63
														if v3 == i32(-1) {
															t67 := int64(load64(m.memory[int64(uint32(v2))+312:]))
															v21 = t67
															v16 = i32(9)
															m.memory[int64(uint32(v2))+304] = byte(i32(9))
															if v7 == i32(-1) {
																goto l38
															}
															m.fn964(v2 + i32(304))
															v16 = i32(2)
														l38:
															v15 = v37
															goto l57
														}
														t64 := int64(load64(m.memory[int64(uint32(v2))+320:]))
														store64(m.memory[int64(uint32(v2))+228:], uint64(t64))
														t65 := int32(load32(m.memory[int64(uint32(v2))+316:]))
														store32(m.memory[int64(uint32(v2))+224:], uint32(t65))
														v15 = int32(uint32(v7) >> 8)
														v16 = v7
														t66 := int32(load32(m.memory[int64(uint32(v2))+312:]))
														v7 = t66
														goto l35
													}
												case 0:
													{
														t68 := int32(m.memory[uint32(v36)])
														v15 = t68
														if v15 == i32(102) {
															store32(m.memory[int64(uint32(v1))+360:], uint32(i32(0)))
															m.fn164(v2+i32(48), v32)
															t152 := int32(load32(m.memory[int64(uint32(v2))+48:]))
															t153 := int32(load32(m.memory[int64(uint32(v2))+52:]))
															m.fn220(v2+i32(304), v1, t152, t153, v27)
															t154 := int32(load32(m.memory[int64(uint32(v2))+304:]))
															v3 = t154
															if v3 == i32(-1) {
																goto l47
															}
															t155 := int64(load64(m.memory[int64(uint32(v2))+320:]))
															store64(m.memory[int64(uint32(v2))+228:], uint64(t155))
															t156 := int64(load64(m.memory[int64(uint32(v2))+312:]))
															t157 := v2
															v21 = t156
															store32(m.memory[int64(uint32(t157))+224:], uint32(int64(uint64(v21)>>32)))
															t158 := int32(load32(m.memory[int64(uint32(v2))+308:]))
															v16 = t158
															v15 = int32(uint32(v16) >> 8)
															v7 = int32(v21)
															goto l35
														}
														if v15 != i32(118) {
															goto l36
														}
														if v12 == 0 {
															goto l41
														}
														{
															switch v18 + i32(-1) {
															case 0:
																goto l42
															default:
																if v18 != i32(9) {
																	goto l45
																}
																t69 := int32(m.memory[uint32(v12)])
																if t69 != i32(105) {
																	goto l45
																}
																t70 := int32(m.memory[int64(uint32(v12))+1])
																if t70 != i32(110) {
																	goto l45
																}
																t71 := int32(m.memory[int64(uint32(v12))+2])
																if t71 != i32(108) {
																	goto l45
																}
																t72 := int32(m.memory[int64(uint32(v12))+3])
																if t72 != i32(105) {
																	goto l45
																}
																t73 := int32(m.memory[int64(uint32(v12))+4])
																if t73 != i32(110) {
																	goto l45
																}
																t74 := int32(m.memory[int64(uint32(v12))+5])
																if t74 != i32(101) {
																	goto l45
																}
																t75 := int32(m.memory[int64(uint32(v12))+6])
																if t75 != i32(83) {
																	goto l45
																}
																t76 := int32(m.memory[int64(uint32(v12))+7])
																if t76 != i32(116) {
																	goto l45
																}
																t77 := int32(m.memory[int64(uint32(v12))+8])
																if t77 == i32(114) {
																	goto l46
																}
																goto l45
															case 1:
																t78 := int32(m.memory[uint32(v12)])
																if t78 != i32(105) {
																	goto l45
																}
																t79 := int32(m.memory[int64(uint32(v12))+1])
																if t79 != i32(115) {
																	goto l45
																}
															}
														l46:
															store32(m.memory[int64(uint32(v1))+360:], uint32(i32(0)))
															m.fn164(v2+i32(40), v32)
															t80 := int32(load32(m.memory[int64(uint32(v2))+40:]))
															t81 := int32(load32(m.memory[int64(uint32(v2))+44:]))
															m.fn220(v2+i32(304), v1, t80, t81, v27)
															t82 := int32(load32(m.memory[int64(uint32(v2))+304:]))
															v3 = t82
															if v3 == i32(-1) {
																goto l47
															}
															t83 := int64(load64(m.memory[int64(uint32(v2))+320:]))
															store64(m.memory[int64(uint32(v2))+228:], uint64(t83))
															t84 := int64(load64(m.memory[int64(uint32(v2))+312:]))
															t85 := v2
															v21 = t84
															store32(m.memory[int64(uint32(t85))+224:], uint32(int64(uint64(v21)>>32)))
															t86 := int32(load32(m.memory[int64(uint32(v2))+308:]))
															v16 = t86
															v15 = int32(uint32(v16) >> 8)
															v7 = int32(v21)
															goto l35
														}
													l42:
														{
															t87 := int32(m.memory[uint32(v12)])
															v7 = t87
															switch v7 + i32(-98) {
															case 0, 3:
																goto l41
															case 1, 2:
																goto l45
															default:
																if v7 == i32(110) {
																	goto l41
																}
																if v7 == i32(115) {
																	goto l41
																}
															}
														}
													l45:
														store32(m.memory[int64(uint32(v1))+372:], uint32(i32(0)))
													l60:
														{
															store32(m.memory[int64(uint32(v1))+360:], uint32(i32(0)))
															m.fn141(v2+i32(304), v1, v27)
															t88 := int64(load64(m.memory[uint32(v28):]))
															store64(m.memory[int64(uint32(v2))+288:], uint64(t88))
															t89 := int32(load32(m.memory[int64(uint32(v28))+8:]))
															store32(m.memory[int64(uint32(v2))+296:], uint32(t89))
															t90 := int32(load32(m.memory[int64(uint32(v2))+316:]))
															v7 = t90
															t91 := int32(load32(m.memory[int64(uint32(v2))+312:]))
															v16 = t91
															t92 := int32(load32(m.memory[int64(uint32(v2))+308:]))
															v3 = t92
															{
																t93 := int32(load32(m.memory[int64(uint32(v2))+304:]))
																if t93 != i32(1) {
																	t96 := int64(load64(m.memory[int64(uint32(v2))+288:]))
																	store64(m.memory[uint32(v31):], uint64(t96))
																	t97 := int32(load32(m.memory[int64(uint32(v2))+296:]))
																	store32(m.memory[int64(uint32(v31))+8:], uint32(t97))
																	store32(m.memory[int64(uint32(v2))+272:], uint32(v7))
																	store32(m.memory[int64(uint32(v2))+268:], uint32(v16))
																	store32(m.memory[int64(uint32(v2))+264:], uint32(v3))
																	switch v3 + i32(-1) {
																	case 9:
																		m.fn200(v2 + i32(264))
																		v7 = i32(1)
																		v3 = i32(-0x7fffffe9)
																		v16 = i32(1072447)
																		v15 = v25
																		goto l35
																	default:
																		m.fn200(v2 + i32(264))
																		goto l60
																	case 0:
																		t98 := int32(load32(m.memory[int64(uint32(v2))+276:]))
																		v3 = t98
																		m.fn164(v2+i32(32), v32)
																		t99 := int32(load32(m.memory[int64(uint32(v2))+32:]))
																		t100 := int32(load32(m.memory[int64(uint32(v2))+36:]))
																		t101 := m.fn123(v7, v3, t99, t100)
																		if t101 == 0 {
																			goto l55
																		}
																		m.fn134(v16, v7)
																		t102 := int32(load32(m.memory[int64(uint32(v1))+368:]))
																		t103 := int32(load32(m.memory[int64(uint32(v1))+372:]))
																		m.fn1142(v2+i32(304), v2+i32(188), t102, t103, v11, v17, v12, v18)
																		{
																			t104 := int32(load32(m.memory[int64(uint32(v2))+304:]))
																			if t104 != i32(1) {
																				t112 := int32(load16(m.memory[int64(uint32(v2))+313:]))
																				t113 := int32(m.memory[uint32(v34)])
																				v15 = t112 | t113<<16
																				t114 := int64(load64(m.memory[int64(uint32(v2))+328:]))
																				v22 = t114
																				t115 := int64(load64(m.memory[int64(uint32(v2))+320:]))
																				v21 = t115
																				t116 := int32(load32(m.memory[int64(uint32(v2))+316:]))
																				v7 = t116
																				t117 := int32(m.memory[int64(uint32(v2))+312])
																				v16 = t117
																				goto l57
																			}
																			t105 := int32(load32(m.memory[int64(uint32(v2))+328:]))
																			store32(m.memory[int64(uint32(v2))+232:], uint32(t105))
																			t106 := int64(load64(m.memory[int64(uint32(v2))+320:]))
																			store64(m.memory[int64(uint32(v2))+224:], uint64(t106))
																			t107 := int32(load16(m.memory[int64(uint32(v2))+313:]))
																			t108 := int32(m.memory[uint32(v34)])
																			v15 = t107 | t108<<16
																			t109 := int32(load32(m.memory[int64(uint32(v2))+316:]))
																			v7 = t109
																			t110 := int32(m.memory[int64(uint32(v2))+312])
																			v16 = t110
																			t111 := int32(load32(m.memory[int64(uint32(v2))+308:]))
																			v3 = t111
																			goto l35
																		}
																	case 2:
																		store32(m.memory[int64(uint32(v2))+308:], uint32(v7))
																		store32(m.memory[int64(uint32(v2))+304:], uint32(v16))
																		t118 := int64(load64(m.memory[int64(uint32(v2))+288:]))
																		store64(m.memory[int64(uint32(v2))+312:], uint64(t118))
																		m.fn201(v2+i32(240), v2+i32(304))
																		t119 := int32(load32(m.memory[int64(uint32(v2))+248:]))
																		v38 = t119
																		t120 := int32(load32(m.memory[int64(uint32(v2))+244:]))
																		v36 = t120
																		t121 := int32(load32(m.memory[int64(uint32(v2))+240:]))
																		v3 = t121
																		if v3 == i32(-2) {
																			m.fn134(v16, v7)
																			v15 = int32(uint32(v36) >> 8)
																			v3 = i32(-0x7fffffd6)
																			v16 = v36
																			v7 = v38
																			goto l35
																		}
																		m.fn75(v26, v36, v38)
																		m.fn134(v3, v36)
																		goto l55
																	case 8:
																		store32(m.memory[int64(uint32(v2))+244:], uint32(v7))
																		store32(m.memory[int64(uint32(v2))+240:], uint32(v16))
																		t122 := int64(load64(m.memory[int64(uint32(v2))+288:]))
																		store64(m.memory[int64(uint32(v2))+248:], uint64(t122))
																		m.fn202(v2+i32(304), v2+i32(240), v26)
																		t123 := int32(load32(m.memory[int64(uint32(v2))+304:]))
																		v3 = t123
																		if v3 != i32(-1) {
																			t175 := int64(load64(m.memory[uint32(v30):]))
																			store64(m.memory[int64(uint32(v2))+224:], uint64(t175))
																			t176 := int32(load32(m.memory[int64(uint32(v30))+8:]))
																			store32(m.memory[int64(uint32(v2))+232:], uint32(t176))
																			t177 := int32(load32(m.memory[int64(uint32(v2))+312:]))
																			v40 = t177
																			t178 := int32(load32(m.memory[int64(uint32(v2))+308:]))
																			v36 = t178
																			m.fn134(v16, v7)
																			v15 = int32(uint32(v36) >> 8)
																			v16 = v36
																			v7 = v40
																			goto l35
																		}
																	}
																l55:
																	m.fn134(v16, v7)
																	goto l60
																}
																t94 := int32(load32(m.memory[int64(uint32(v2))+296:]))
																store32(m.memory[int64(uint32(v2))+232:], uint32(t94))
																t95 := int64(load64(m.memory[int64(uint32(v2))+288:]))
																store64(m.memory[int64(uint32(v2))+224:], uint64(t95))
																v15 = int32(uint32(v16) >> 8)
																goto l35
															}
														}
													l41:
														store32(m.memory[int64(uint32(v1))+360:], uint32(i32(0)))
														m.fn141(v2+i32(304), v1, v27)
														t124 := int64(load64(m.memory[int64(uint32(v2))+324:]))
														v21 = t124
														t125 := int32(load32(m.memory[int64(uint32(v2))+320:]))
														v7 = t125
														t126 := int32(load32(m.memory[int64(uint32(v2))+316:]))
														v39 = t126
														t127 := int32(load32(m.memory[int64(uint32(v2))+312:]))
														v36 = t127
														t128 := int32(load32(m.memory[int64(uint32(v2))+308:]))
														v3 = t128
														t129 := int32(load32(m.memory[int64(uint32(v2))+304:]))
														if t129 != 0 {
															store64(m.memory[int64(uint32(v2))+228:], uint64(v21))
															store32(m.memory[int64(uint32(v2))+224:], uint32(v7))
															v15 = int32(uint32(v36) >> 8)
															v16 = v36
															v7 = v39
															goto l35
														}
														store64(m.memory[int64(uint32(v2))+280:], uint64(v21))
														store32(m.memory[int64(uint32(v2))+276:], uint32(v7))
														store32(m.memory[int64(uint32(v2))+272:], uint32(v39))
														store32(m.memory[int64(uint32(v2))+268:], uint32(v36))
														store32(m.memory[int64(uint32(v2))+264:], uint32(v3))
														switch v3 + i32(-1) {
														case 0:
															m.fn164(v2+i32(24), v32)
															t130 := int32(load32(m.memory[int64(uint32(v2))+24:]))
															t131 := int32(load32(m.memory[int64(uint32(v2))+28:]))
															t132 := m.fn123(v39, v7, t130, t131)
															if t132 == 0 {
																goto l66
															}
															v13 = i32(0)
															v16 = i32(9)
															v15 = v38
															v3 = v40
															v7 = v41
															goto l67
														case 2:
															m.fn1142(v2+i32(304), v2+i32(188), v39, v7, v11, v17, v12, v18)
															t133 := int32(load32(m.memory[int64(uint32(v2))+304:]))
															if t133 != 0 {
																goto l68
															}
															t134 := int64(load64(m.memory[uint32(v28):]))
															t135 := v2
															v21 = t134
															store64(m.memory[int64(uint32(t135))+288:], uint64(v21))
															t136 := int32(load32(m.memory[int64(uint32(v28))+8:]))
															t137 := v2
															v7 = t136
															store32(m.memory[int64(uint32(t137))+296:], uint32(v7))
															t138 := int32(load32(m.memory[int64(uint32(v2))+332:]))
															v3 = t138
															t139 := int32(load32(m.memory[int64(uint32(v2))+312:]))
															v16 = t139
															t140 := int32(load32(m.memory[int64(uint32(v2))+316:]))
															v42 = t140
															store64(m.memory[uint32(v29):], uint64(v21))
															store32(m.memory[int64(uint32(v29))+8:], uint32(v7))
															store32(m.memory[int64(uint32(v2))+244:], uint32(v42))
															store32(m.memory[int64(uint32(v2))+240:], uint32(v16))
															store32(m.memory[int64(uint32(v2))+260:], uint32(v3))
															goto l69
														default:
															if v3 == i32(10) {
																m.fn200(v2 + i32(264))
																v7 = i32(1)
																v3 = i32(-0x7fffffe9)
																v16 = i32(1072447)
																v15 = v25
																v13 = i32(1)
																goto l35
															}
															fallthrough
														case 1:
															v16 = i32(9)
															m.memory[int64(uint32(v2))+240] = byte(i32(9))
															m.fn200(v2 + i32(264))
															goto l71
														}
													l66:
														v16 = i32(9)
														m.memory[int64(uint32(v2))+240] = byte(i32(9))
													l69:
														m.fn134(v36, v39)
													l71:
														store32(m.memory[int64(uint32(v1))+360:], uint32(i32(0)))
														m.fn164(v2+i32(16), v32)
														t141 := int32(load32(m.memory[int64(uint32(v2))+16:]))
														t142 := int32(load32(m.memory[int64(uint32(v2))+20:]))
														m.fn220(v2+i32(304), v1, t141, t142, v27)
														{
															t143 := int32(load32(m.memory[int64(uint32(v2))+304:]))
															v3 = t143
															if v3 == i32(-1) {
																t148 := int32(load16(m.memory[int64(uint32(v2))+241:]))
																t149 := int32(m.memory[int64(uint32(v2))+243])
																v15 = t148 | t149<<16
																t150 := int64(load64(m.memory[int64(uint32(v2))+248:]))
																v21 = t150
																t151 := int64(load64(m.memory[int64(uint32(v2))+256:]))
																v22 = t151
																v7 = v42
																goto l57
															}
															t144 := int64(load64(m.memory[int64(uint32(v2))+320:]))
															store64(m.memory[int64(uint32(v2))+228:], uint64(t144))
															t145 := int64(load64(m.memory[int64(uint32(v2))+312:]))
															t146 := v2
															v21 = t145
															store32(m.memory[int64(uint32(t146))+224:], uint32(int64(uint64(v21)>>32)))
															t147 := int32(load32(m.memory[int64(uint32(v2))+308:]))
															v16 = t147
															v15 = int32(uint32(v16) >> 8)
															v7 = int32(v21)
															m.fn964(v2 + i32(240))
															goto l35
														}
													}
												l47:
													v16 = i32(9)
													v15 = v37
													goto l57
												}
											}
										case 1:
											t159 := int32(load32(m.memory[int64(uint32(v2))+172:]))
											t160 := v2 + i32(64)
											v7 = t159
											t161 := int32(load32(m.memory[int64(uint32(v2))+176:]))
											m.fn553(t160, v7, t161)
											t162 := int32(load32(m.memory[int64(uint32(v2))+68:]))
											if t162 != i32(1) {
												goto l73
											}
											t163 := int32(load32(m.memory[int64(uint32(v2))+64:]))
											t164 := int32(m.memory[uint32(t163)])
											if t164 != i32(99) {
												goto l73
											}
											t165 := int32(load32(m.memory[int64(uint32(v2))+168:]))
											m.fn134(t165, v7)
											t166 := int32(load32(m.memory[int64(uint32(v1))+404:]))
											store32(m.memory[int64(uint32(v1))+404:], uint32(t166+i32(1)))
											t167 := int64(load64(m.memory[int64(uint32(v2))+136:]))
											store64(m.memory[uint32(v0):], uint64(t167))
											t168 := int64(load64(m.memory[int64(uint32(v2))+144:]))
											store64(m.memory[int64(uint32(v0))+8:], uint64(t168))
											t169 := int64(load64(m.memory[int64(uint32(v2))+152:]))
											store64(m.memory[int64(uint32(v0))+16:], uint64(t169))
											store32(m.memory[int64(uint32(v0))+28:], uint32(v19))
											store32(m.memory[int64(uint32(v0))+24:], uint32(v20))
											goto l74
										default:
											if v5 != i32(10) {
												goto l75
											}
											store32(m.memory[int64(uint32(v0))+12:], uint32(i32(1)))
											store32(m.memory[int64(uint32(v0))+8:], uint32(i32(1077932)))
											store32(m.memory[int64(uint32(v0))+4:], uint32(i32(-0x7fffffe9)))
											m.memory[uint32(v0)] = byte(i32(254))
											m.fn200(v33)
											goto l29
										}
									}
								l68:
									t170 := int64(load64(m.memory[uint32(v28):]))
									store64(m.memory[int64(uint32(v2))+224:], uint64(t170))
									t171 := int32(load32(m.memory[int64(uint32(v28))+8:]))
									store32(m.memory[int64(uint32(v2))+232:], uint32(t171))
									t172 := int32(load32(m.memory[int64(uint32(v2))+312:]))
									v16 = t172
									v15 = int32(uint32(v16) >> 8)
									t173 := int32(load32(m.memory[int64(uint32(v2))+316:]))
									v7 = t173
									t174 := int32(load32(m.memory[int64(uint32(v2))+308:]))
									v3 = t174
									v13 = i32(1)
								}
							l67:
								m.fn134(v36, v39)
								goto l35
							l57:
								store64(m.memory[int64(uint32(v2))+232:], uint64(v22))
								store64(m.memory[int64(uint32(v2))+224:], uint64(v21))
								v13 = i32(0)
								v3 = v40
								v37 = v15
								goto l35
							l36:
								v15 = v23
							l35:
								v16 = v15<<8 | v16&i32(255)
								if v13 != 0 {
									goto l76
								}
								t179 := int64(load64(m.memory[int64(uint32(v2))+232:]))
								store64(m.memory[int64(uint32(v2))+216:], uint64(t179))
								t180 := int64(load64(m.memory[int64(uint32(v2))+224:]))
								store64(m.memory[int64(uint32(v2))+208:], uint64(t180))
								m.fn964(v2 + i32(136))
								t181 := int64(load64(m.memory[int64(uint32(v2))+208:]))
								store64(m.memory[uint32(v10):], uint64(t181))
								t182 := int64(load64(m.memory[int64(uint32(v2))+216:]))
								store64(m.memory[int64(uint32(v10))+8:], uint64(t182))
								store32(m.memory[int64(uint32(v2))+140:], uint32(v7))
								store32(m.memory[int64(uint32(v2))+136:], uint32(v16))
								m.fn134(v35, v14)
								v38 = v15
								v40 = v3
								v41 = v7
							}
						l75:
							switch v5 {
							case 0:
								goto l77
							case 1:
								goto l73
							default:
								m.fn200(v33)
								goto l77
							}
						l73:
							t183 := int32(load32(m.memory[int64(uint32(v2))+168:]))
							t184 := int32(load32(m.memory[int64(uint32(v2))+172:]))
							m.fn134(t183, t184)
							goto l77
						}
					l76:
						t185 := int32(load32(m.memory[int64(uint32(v2))+232:]))
						t186 := v2
						v1 = t185
						store32(m.memory[int64(uint32(t186))+216:], uint32(v1))
						t187 := int64(load64(m.memory[int64(uint32(v2))+224:]))
						t188 := v2
						v21 = t187
						store64(m.memory[int64(uint32(t188))+208:], uint64(v21))
						store32(m.memory[int64(uint32(v0))+12:], uint32(v7))
						store32(m.memory[int64(uint32(v0))+8:], uint32(v16))
						store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
						store64(m.memory[int64(uint32(v0))+16:], uint64(v21))
						store32(m.memory[int64(uint32(v0))+24:], uint32(v1))
						m.memory[uint32(v0)] = byte(i32(254))
						m.fn134(v35, v14)
					}
				l29:
					m.fn964(v2 + i32(136))
				l74:
					t189 := int32(load32(m.memory[int64(uint32(v2))+108:]))
					v1 = t189
				}
			l20:
				m.fn134(v9, v8)
				goto l26
			l3:
				t190 := int32(load32(m.memory[int64(uint32(v2))+120:]))
				t191 := v2 + i32(96)
				v7 = t190
				t192 := int32(load32(m.memory[int64(uint32(v2))+124:]))
				t193 := v7
				v16 = t192
				m.fn553(t191, t193, v16)
				{
					{
						t194 := int32(load32(m.memory[int64(uint32(v2))+100:]))
						if t194 != i32(3) {
							goto l79
						}
						t195 := int32(load32(m.memory[int64(uint32(v2))+96:]))
						v15 = t195
						t196 := int32(load16(m.memory[uint32(v15):]))
						t197 := int32(m.memory[uint32(v15+i32(2))])
						if t196|t197<<16 == i32(7827314) {
							store32(m.memory[int64(uint32(v1))+404:], uint32(i32(0)))
							t203 := int32(load32(m.memory[int64(uint32(v1))+400:]))
							store32(m.memory[int64(uint32(v1))+400:], uint32(t203+i32(1)))
							t204 := int32(load32(m.memory[int64(uint32(v2))+116:]))
							m.fn134(t204, v7)
							goto l25
						}
					}
				l79:
					m.fn553(v2+i32(88), v7, v16)
					t198 := int32(load32(m.memory[int64(uint32(v2))+88:]))
					t199 := int32(load32(m.memory[int64(uint32(v2))+92:]))
					t200 := m.fn949(t198, t199, i32(1077933))
					if t200 != 0 {
						goto l81
					}
					t201 := int32(load32(m.memory[int64(uint32(v2))+116:]))
					t202 := int32(load32(m.memory[int64(uint32(v2))+120:]))
					m.fn134(t201, t202)
					goto l25
				}
			l81:
			}
			m.memory[uint32(v0)] = byte(i32(255))
			t205 := int32(load32(m.memory[int64(uint32(v2))+116:]))
			m.fn134(t205, v7)
			v1 = i32(0)
		}
	l26:
		if v1 != 0 {
			goto l1
		}
		t206 := int32(load32(m.memory[int64(uint32(v2))+112:]))
		if uint32(t206) < uint32(i32(2)) {
			goto l1
		}
		m.fn200(v6)
	}
l1:
	m.g0 = v2 + i32(336)
}
func (m *Module) fn1140(v0, v1, v2 int32) {
	var v3, v4, v5 int32
	t0 := m.g0
	v3 = t0 - i32(32)
	m.g0 = v3
	m.fn1141(v3+i32(8), v1, v2)
	t1 := int32(load32(m.memory[int64(uint32(v3))+20:]))
	v1 = t1
	t2 := int32(load32(m.memory[int64(uint32(v3))+16:]))
	v2 = t2
	t3 := int32(load32(m.memory[int64(uint32(v3))+12:]))
	v4 = t3
	{
		t4 := int32(load32(m.memory[int64(uint32(v3))+8:]))
		v5 = t4
		if v5 == i32(-1) {
			goto l0
		}
		t5 := int64(load64(m.memory[int64(uint32(v3))+24:]))
		store64(m.memory[int64(uint32(v0))+16:], uint64(t5))
		store32(m.memory[int64(uint32(v0))+12:], uint32(v1))
		store32(m.memory[int64(uint32(v0))+8:], uint32(v2))
		store32(m.memory[int64(uint32(v0))+4:], uint32(v4))
		store32(m.memory[uint32(v0):], uint32(v5))
		goto l1
	}
l0:
	store32(m.memory[int64(uint32(v3))+8:], uint32(i32(-0x7fffffe3)))
	if v2&i32(1) != 0 {
		goto l2
	}
	store32(m.memory[uint32(v0):], uint32(i32(-0x7fffffe3)))
	goto l1
l2:
	m.fn1564(v3 + i32(8))
	store32(m.memory[int64(uint32(v0))+8:], uint32(v1))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v4))
	store32(m.memory[uint32(v0):], uint32(i32(-1)))
l1:
	m.g0 = v3 + i32(32)
}
func (m *Module) fn1141(v0, v1, v2 int32) {
	var v3, v4, v5, v6, v7, v8 int32
	t0 := m.g0
	v3 = t0 - i32(32)
	m.g0 = v3
	v4 = i32(0)
	v5 = i32(0)
	v6 = i32(0)
	{
	l5:
		if v2 != v6 {
			{
				t1 := int32(m.memory[uint32(v1+v6)])
				v8 = t1
				v7 = v8 + i32(-65)
				if uint32(v7&i32(255)) < uint32(i32(26)) {
					goto l2
				}
				v7 = v8 + i32(-97)
				if uint32(v7&i32(255)) < uint32(i32(26)) {
					goto l2
				}
				v4 = (v8 + i32(-48)) & i32(255)
				if uint32(v4) < uint32(i32(10)) {
					v7 = v6 + i32(1)
					goto l1
				}
				m.memory[int64(uint32(v0))+4] = byte(v8)
				store32(m.memory[uint32(v0):], uint32(i32(-0x7fffffe5)))
				goto l4
			}
		l2:
			v5 = v5*i32(26) + v7&i32(255) + i32(1)
			v6 = v6 + i32(1)
			goto l5
		}
		v7 = v2
		goto l1
	l1:
		v6 = v2 - v7
		p2 := v6
		if uint32(v6) > uint32(v2) {
			p2 = i32(0)
		}
		v6 = p2
		v7 = v1 + v7
	l9:
		if v6 != 0 {
			goto l6
		}
		store32(m.memory[int64(uint32(v3))+8:], uint32(i32(-0x7fffffe2)))
		if v4 != 0 {
			m.fn1564(v3 + i32(8))
			store32(m.memory[int64(uint32(v0))+12:], uint32(v5+i32(-1)))
			t3 := v0
			var p4 int32
			if v5 != i32(0) {
				p4 = 1
			}
			store32(m.memory[int64(uint32(t3))+8:], uint32(p4))
			store32(m.memory[int64(uint32(v0))+4:], uint32(v4+i32(-1)))
			store32(m.memory[uint32(v0):], uint32(i32(-1)))
			goto l4
		}
		store32(m.memory[uint32(v0):], uint32(i32(-0x7fffffe2)))
		goto l4
	l6:
		{
			t5 := int32(m.memory[uint32(v7)])
			v2 = t5
			v8 = (v2 + i32(-48)) & i32(255)
			if uint32(v8) >= uint32(i32(10)) {
				goto l8
			}
			v4 = v4*i32(10) + v8
			v6 = v6 + i32(-1)
			v7 = v7 + i32(1)
			goto l9
		}
	l8:
		m.memory[int64(uint32(v0))+4] = byte(v2)
		store32(m.memory[uint32(v0):], uint32(i32(-0x7fffffe5)))
	}
l4:
	m.g0 = v3 + i32(32)
}
func (m *Module) fn1142(v0, v1, v2, v3, v4, v5, v6, v7 int32) {
	var v8, v9, v10, v11, v12 int32
	var v13 int64
	var v14, v15 int32
	var v16 int64
	var v17 int32
	var v18 float64
	var v19 int64
	t0 := m.g0
	v8 = t0 - i32(1728)
	m.g0 = v8
	{
		if v4 != 0 {
			goto l0
		}
		v9 = i32(1070108)
		goto l1
	l0:
		m.fn1587(v8+i32(944), v4, v5)
		t1 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		t2 := int32(load32(m.memory[int64(uint32(v8))+948:]))
		t3 := int32(load32(m.memory[int64(uint32(v8))+944:]))
		p4 := i32(0)
		if t3 == i32(-1) {
			p4 = t2
		}
		v4 = p4
		t5 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		p6 := i32(0)
		if uint32(v4) < uint32(t5) {
			p6 = t1 + v4
		}
		v9 = p6
	}
l1:
	{
		{
			if v6 == 0 {
				goto l2
			}
			switch v7 + i32(-1) {
			case 2:
				t44 := int32(m.memory[uint32(v6)])
				if t44 != i32(115) {
					goto l4
				}
				t45 := int32(m.memory[int64(uint32(v6))+1])
				if t45 != i32(116) {
					goto l4
				}
				t46 := int32(m.memory[int64(uint32(v6))+2])
				if t46 != i32(114) {
					goto l4
				}
				m.fn1588(v8+i32(944), v2, v3)
				t47 := int32(load32(m.memory[int64(uint32(v8))+952:]))
				v4 = t47
				t48 := int32(load32(m.memory[int64(uint32(v8))+948:]))
				v7 = t48
				{
					t49 := int32(load32(m.memory[int64(uint32(v8))+944:]))
					v6 = t49
					if v6 == i32(-1) {
						m.fn377(v8+i32(112), v7, v4)
						t52 := int32(load32(m.memory[int64(uint32(v8))+120:]))
						store32(m.memory[int64(uint32(v8))+163:], uint32(t52))
						t53 := int64(load64(m.memory[int64(uint32(v8))+112:]))
						store64(m.memory[int64(uint32(v8))+155:], uint64(t53))
						m.memory[int64(uint32(v0))+8] = byte(i32(2))
						t54 := int64(load64(m.memory[int64(uint32(v8))+152:]))
						store64(m.memory[int64(uint32(v0))+9:], uint64(t54))
						t55 := int64(load64(m.memory[int64(uint32(v8))+159:]))
						store64(m.memory[int64(uint32(v0))+16:], uint64(t55))
						store32(m.memory[uint32(v0):], uint32(i32(0)))
						goto l12
					}
					t50 := int32(load32(m.memory[int64(uint32(v8))+964:]))
					store32(m.memory[int64(uint32(v0))+24:], uint32(t50))
					t51 := int64(load64(m.memory[int64(uint32(v8))+956:]))
					store64(m.memory[int64(uint32(v0))+16:], uint64(t51))
					store32(m.memory[int64(uint32(v0))+12:], uint32(v4))
					store32(m.memory[int64(uint32(v0))+8:], uint32(v7))
					store32(m.memory[int64(uint32(v0))+4:], uint32(v6))
					store32(m.memory[uint32(v0):], uint32(i32(1)))
					goto l12
				}
			default:
				goto l4
			case 0:
				t7 := int32(m.memory[uint32(v6)])
				v4 = t7
				switch v4 + i32(-98) {
				case 1:
					goto l4
				case 2:
					m.fn1588(v8+i32(944), v2, v3)
					t19 := int32(load32(m.memory[int64(uint32(v8))+952:]))
					v4 = t19
					t20 := int32(load32(m.memory[int64(uint32(v8))+948:]))
					v7 = t20
					{
						t21 := int32(load32(m.memory[int64(uint32(v8))+944:]))
						v6 = t21
						if v6 == i32(-1) {
							m.fn377(v8+i32(112), v7, v4)
							t24 := int32(load32(m.memory[int64(uint32(v8))+120:]))
							store32(m.memory[int64(uint32(v8))+163:], uint32(t24))
							t25 := int64(load64(m.memory[int64(uint32(v8))+112:]))
							store64(m.memory[int64(uint32(v8))+155:], uint64(t25))
							m.memory[int64(uint32(v0))+8] = byte(i32(6))
							t26 := int64(load64(m.memory[int64(uint32(v8))+152:]))
							store64(m.memory[int64(uint32(v0))+9:], uint64(t26))
							t27 := int64(load64(m.memory[int64(uint32(v8))+159:]))
							store64(m.memory[int64(uint32(v0))+16:], uint64(t27))
							store32(m.memory[uint32(v0):], uint32(i32(0)))
							goto l12
						}
						t22 := int32(load32(m.memory[int64(uint32(v8))+964:]))
						store32(m.memory[int64(uint32(v0))+24:], uint32(t22))
						t23 := int64(load64(m.memory[int64(uint32(v8))+956:]))
						store64(m.memory[int64(uint32(v0))+16:], uint64(t23))
						store32(m.memory[int64(uint32(v0))+12:], uint32(v4))
						store32(m.memory[int64(uint32(v0))+8:], uint32(v7))
						store32(m.memory[int64(uint32(v0))+4:], uint32(v6))
						store32(m.memory[uint32(v0):], uint32(i32(1)))
						goto l12
					}
				case 3:
					m.fn1588(v8+i32(944), v2, v3)
					t28 := int32(load32(m.memory[int64(uint32(v8))+952:]))
					v4 = t28
					t29 := int32(load32(m.memory[int64(uint32(v8))+948:]))
					v7 = t29
					{
						t30 := int32(load32(m.memory[int64(uint32(v8))+944:]))
						v6 = t30
						if v6 == i32(-1) {
							v6 = i32(0)
							{
								t33 := m.fn1562(v7, v4, i32(1089054), i32(7))
								if t33 != 0 {
									goto l18
								}
								{
									t34 := m.fn1562(v7, v4, i32(1089079), i32(4))
									if t34 == 0 {
										goto l19
									}
									v6 = i32(1)
									goto l18
								}
							l19:
								{
									t35 := m.fn1562(v7, v4, i32(1089068), i32(6))
									if t35 == 0 {
										goto l20
									}
									v6 = i32(2)
									goto l18
								}
							l20:
								{
									t36 := m.fn1562(v7, v4, i32(1089048), i32(6))
									if t36 == 0 {
										goto l21
									}
									v6 = i32(3)
									goto l18
								}
							l21:
								v6 = i32(5)
								{
									t37 := m.fn1562(v7, v4, i32(1089074), i32(5))
									if t37 == 0 {
										goto l22
									}
									v6 = i32(4)
									goto l18
								}
							l22:
								t38 := m.fn1562(v7, v4, i32(1088624), i32(5))
								if t38 != 0 {
									goto l18
								}
								t39 := m.fn1562(v7, v4, i32(1089061), i32(7))
								if t39 == 0 {
									m.fn377(v8+i32(156), v7, v4)
									t56 := int64(load64(m.memory[int64(uint32(v8))+157:]))
									store64(m.memory[int64(uint32(v0))+9:], uint64(t56))
									t57 := int64(load64(m.memory[int64(uint32(v8))+165:]))
									store64(m.memory[int64(uint32(v0))+17:], uint64(t57))
									t58 := int32(load32(m.memory[int64(uint32(v8))+172:]))
									store32(m.memory[int64(uint32(v0))+24:], uint32(t58))
									t59 := int32(m.memory[int64(uint32(v8))+156])
									m.memory[int64(uint32(v0))+8] = byte(t59)
									store64(m.memory[uint32(v0):], uint64(i64(-0x7fffffdaffffffff)))
									goto l12
								}
								v6 = i32(6)
							}
						l18:
							m.memory[int64(uint32(v0))+9] = byte(v6)
							m.memory[int64(uint32(v0))+8] = byte(i32(8))
							store32(m.memory[uint32(v0):], uint32(i32(0)))
							goto l12
						}
						t31 := int32(load32(m.memory[int64(uint32(v8))+964:]))
						store32(m.memory[int64(uint32(v0))+24:], uint32(t31))
						t32 := int64(load64(m.memory[int64(uint32(v8))+956:]))
						store64(m.memory[int64(uint32(v0))+16:], uint64(t32))
						store32(m.memory[int64(uint32(v0))+12:], uint32(v4))
						store32(m.memory[int64(uint32(v0))+8:], uint32(v7))
						store32(m.memory[int64(uint32(v0))+4:], uint32(v6))
						store32(m.memory[uint32(v0):], uint32(i32(1)))
						goto l12
					}
				default:
					if v4 == i32(110) {
						goto l2
					}
					if v4 == i32(115) {
						if v3 == 0 {
							store32(m.memory[uint32(v0):], uint32(i32(0)))
							m.memory[int64(uint32(v0))+8] = byte(i32(9))
							goto l12
						}
						m.fn1587(v8+i32(944), v2, v3)
						v4 = i32(0)
						{
							{
								t10 := int32(load32(m.memory[int64(uint32(v8))+948:]))
								t11 := int32(load32(m.memory[int64(uint32(v8))+944:]))
								p12 := i32(0)
								if t11 == i32(-1) {
									p12 = t10
								}
								v7 = p12
								t13 := int32(load32(m.memory[int64(uint32(v1))+4:]))
								if uint32(v7) >= uint32(t13) {
									goto l14
								}
								t14 := int32(load32(m.memory[uint32(v1):]))
								t15 := int64(load64(m.memory[int64(uint32(t14+v7*i32(12)))+4:]))
								store64(m.memory[int64(uint32(v0))+12:], uint64(t15))
								m.memory[int64(uint32(v0))+8] = byte(i32(3))
								store32(m.memory[int64(uint32(v8))+952:], uint32(i32(51)))
								store32(m.memory[int64(uint32(v8))+948:], uint32(i32(1099828)))
								store32(m.memory[int64(uint32(v8))+944:], uint32(i32(-0x7fffffdd)))
								m.fn1564(v8 + i32(944))
								goto l15
							}
						l14:
							t16 := int64(load64(m.memory[int64(uint32(v8))+960:]))
							store64(m.memory[int64(uint32(v0))+20:], uint64(t16))
							store32(m.memory[int64(uint32(v8))+944:], uint32(i32(-0x7fffffdd)))
							store32(m.memory[int64(uint32(v8))+948:], uint32(i32(1099828)))
							t17 := int64(load64(m.memory[int64(uint32(v8))+944:]))
							store64(m.memory[int64(uint32(v0))+4:], uint64(t17))
							store32(m.memory[int64(uint32(v8))+952:], uint32(i32(51)))
							t18 := int64(load64(m.memory[int64(uint32(v8))+952:]))
							store64(m.memory[int64(uint32(v0))+12:], uint64(t18))
							v4 = i32(1)
						}
					l15:
						store32(m.memory[uint32(v0):], uint32(v4))
						goto l12
					}
					goto l4
				case 0:
					v4 = i32(1)
					{
						if v3 != i32(1) {
							goto l11
						}
						t8 := int32(m.memory[uint32(v2)])
						var p9 int32
						if t8 != i32(48) {
							p9 = 1
						}
						v4 = p9
					}
				l11:
					m.memory[int64(uint32(v0))+9] = byte(v4)
					m.memory[int64(uint32(v0))+8] = byte(i32(4))
					store32(m.memory[uint32(v0):], uint32(i32(0)))
					goto l12
				}
			}
		l2:
			{
				if v3 == 0 {
					store32(m.memory[uint32(v0):], uint32(i32(0)))
					m.memory[int64(uint32(v0))+8] = byte(i32(9))
					goto l12
				}
				store32(m.memory[int64(uint32(v8))+136:], uint32(v2))
				t40 := v8
				v10 = v2 + v3
				store32(m.memory[int64(uint32(t40))+140:], uint32(v10))
				{
					t41 := int32(m.memory[uint32(v2)])
					v11 = t41
					if v11 == i32(45) {
						t43 := v8
						v12 = v2 + i32(1)
						store32(m.memory[int64(uint32(t43))+136:], uint32(v12))
						if v3 != i32(1) {
							goto l26
						}
						goto l27
					}
					v12 = v2
					if v11 != i32(43) {
						goto l26
					}
					t42 := v8
					v12 = v2 + i32(1)
					store32(m.memory[int64(uint32(t42))+136:], uint32(v12))
					if v3 == i32(1) {
						goto l27
					}
					goto l26
				}
			}
		l26:
			v13 = i64(0)
			store64(m.memory[int64(uint32(v8))+152:], uint64(i64(0)))
			m.fn203(v8+i32(136), v8+i32(152))
			t60 := int32(load32(m.memory[int64(uint32(v8))+136:]))
			v14 = t60
			v7 = v14 - v12
			v4 = i32(0)
			{
				{
					t61 := int32(load32(m.memory[int64(uint32(v8))+140:]))
					t62 := v14
					v5 = t61
					if t62 != v5 {
						goto l29
					}
					v15 = v14
					goto l30
				}
			l29:
				v15 = v14
				t63 := int32(m.memory[uint32(v14)])
				if t63 != i32(46) {
					goto l30
				}
				t64 := v8
				v4 = v14 + i32(1)
				store32(m.memory[int64(uint32(t64))+136:], uint32(v4))
				{
					if v5-v4 < i32(8) {
						goto l31
					}
					t65 := int64(load64(m.memory[uint32(v4):]))
					v13 = t65
					if (v13+i64(5063812098665367110)|(v13+i64(-3472328296227680304)))&i64(-0x7f7f7f7f7f7f7f80) != i64(0) {
						goto l31
					}
					t66 := int64(load64(m.memory[int64(uint32(v8))+152:]))
					v16 = t66
					t67 := v8
					v17 = v14 + i32(9)
					store32(m.memory[int64(uint32(t67))+136:], uint32(v17))
					t68 := fn204(v13)
					t69 := v8
					v13 = v16*i64(100000000) + t68
					store64(m.memory[int64(uint32(t69))+152:], uint64(v13))
					if v5-v17 < i32(8) {
						goto l31
					}
					t70 := int64(load64(m.memory[uint32(v17):]))
					v16 = t70
					if (v16+i64(5063812098665367110)|(v16+i64(-3472328296227680304)))&i64(-0x7f7f7f7f7f7f7f80) != i64(0) {
						goto l31
					}
					store32(m.memory[int64(uint32(v8))+136:], uint32(v14+i32(17)))
					t71 := fn204(v16)
					store64(m.memory[int64(uint32(v8))+152:], uint64(t71+v13*i64(100000000)))
				}
			l31:
				m.fn203(v8+i32(136), v8+i32(152))
				t72 := int32(load32(m.memory[int64(uint32(v8))+136:]))
				v15 = t72
				v4 = v15 - v4
				v13 = int64(i32(0) - v4)
			}
		l30:
			{
				{
					{
						{
							v7 = v4 + v7
							if v7 == 0 {
								v14 = i32(3)
								if uint32(v3) < uint32(i32(3)) {
									goto l27
								}
								v18 = math.Float64frombits(0x7ff8000000000000)
								t81 := m.fn206(v2, v3, i32(1087560))
								if t81 != 0 {
									goto l40
								}
								{
									t82 := m.fn206(v2, v3, i32(1108005))
									if t82 != 0 {
										v18 = math.Float64frombits(0x7ff0000000000000)
										t95 := m.fn208(v2, v3)
										v14 = t95
										goto l40
									}
									if v3 == i32(3) {
										goto l27
									}
									switch v11 + i32(-43) {
									default:
										goto l27
									case 0:
										m.fn207(v8, v2, v3, i32(1))
										{
											t83 := int32(load32(m.memory[uint32(v8):]))
											v4 = t83
											t84 := int32(load32(m.memory[int64(uint32(v8))+4:]))
											t85 := v4
											v7 = t84
											t86 := m.fn206(t85, v7, i32(1087560))
											if t86 == 0 {
												t87 := m.fn206(v4, v7, i32(1108005))
												if t87 == 0 {
													goto l27
												}
												t88 := m.fn208(v4, v7)
												v14 = t88 + i32(1)
												v18 = math.Float64frombits(0x7ff0000000000000)
												goto l40
											}
											v14 = i32(4)
											goto l40
										}
									case 2:
										m.fn207(v8+i32(8), v2, v3, i32(1))
										{
											t89 := int32(load32(m.memory[int64(uint32(v8))+8:]))
											v4 = t89
											t90 := int32(load32(m.memory[int64(uint32(v8))+12:]))
											t91 := v4
											v7 = t90
											t92 := m.fn206(t91, v7, i32(1087560))
											if t92 == 0 {
												t93 := m.fn206(v4, v7, i32(1108005))
												if t93 == 0 {
													goto l27
												}
												t94 := m.fn208(v4, v7)
												v14 = t94 + i32(1)
												v18 = math.Float64frombits(0xfff0000000000000)
												goto l40
											}
											v14 = i32(4)
											v18 = math.Float64frombits(0xfff8000000000000)
											goto l40
										}
									}
								}
							}
							v16 = i64(0)
							t73 := int32(load32(m.memory[int64(uint32(v8))+140:]))
							t74 := v15
							v4 = t73
							if t74 == v4 {
								goto l33
							}
							t75 := int32(m.memory[uint32(v15)])
							if t75|i32(32) != i32(101) {
								goto l33
							}
							t76 := v8
							v5 = v15 + i32(1)
							store32(m.memory[int64(uint32(t76))+136:], uint32(v5))
							v16 = i64(0)
							store64(m.memory[int64(uint32(v8))+944:], uint64(i64(0)))
							if v5 == v4 {
								goto l34
							}
							{
								t77 := int32(m.memory[uint32(v5)])
								v17 = t77
								switch v17 + i32(-43) {
								case 0, 2:
									t78 := v8
									v5 = v15 + i32(2)
									store32(m.memory[int64(uint32(t78))+136:], uint32(v5))
									if v5 == v4 {
										goto l34
									}
									t79 := int32(m.memory[uint32(v5)])
									if uint32((t79+i32(-48))&i32(255)) < uint32(i32(10)) {
										m.fn205(v8+i32(136), v8+i32(944))
										if v17 != i32(45) {
											goto l38
										}
										t80 := int64(load64(m.memory[int64(uint32(v8))+944:]))
										v16 = i64(0) - t80
										goto l39
									}
									goto l34
								default:
									if uint32((v17+i32(-48))&i32(255)) >= uint32(i32(10)) {
										goto l34
									}
									m.fn205(v8+i32(136), v8+i32(944))
									goto l38
								}
							}
						}
					l38:
						t96 := int64(load64(m.memory[int64(uint32(v8))+944:]))
						v16 = t96
						goto l39
					}
				l34:
					store32(m.memory[int64(uint32(v8))+136:], uint32(v15))
				l39:
					t97 := int32(load32(m.memory[int64(uint32(v8))+136:]))
					v15 = t97
				}
			l33:
				v4 = i32(0)
				{
					if v7 < i32(20) {
						goto l46
					}
					v5 = v7 + i32(-19)
					v4 = v12
				l49:
					if v4 == v10 {
						goto l47
					}
					{
						t98 := int32(m.memory[uint32(v4)])
						v7 = t98
						switch v7 + i32(-46) {
						default:
							goto l47
						case 0, 2:
							t99 := v5
							v17 = v7 + i32(-47)
							p100 := v17
							if uint32(v17) > uint32(v7) {
								p100 = i32(0)
							}
							v5 = t99 - p100
							v4 = v4 + i32(1)
							goto l49
						}
					}
				l47:
					v4 = i32(0)
					if v5 <= i32(0) {
						goto l46
					}
					store64(m.memory[int64(uint32(v8))+152:], uint64(i64(0)))
					store32(m.memory[int64(uint32(v8))+948:], uint32(v10))
					store32(m.memory[int64(uint32(v8))+944:], uint32(v12))
					m.fn209(v8+i32(944), v8+i32(152))
					t101 := int32(load32(m.memory[int64(uint32(v8))+944:]))
					v4 = t101
					{
						t102 := int64(load64(m.memory[int64(uint32(v8))+152:]))
						if uint64(t102) > uint64(i64(999999999999999999)) {
							goto l50
						}
						t103 := v8
						v4 = v4 + i32(1)
						store32(m.memory[int64(uint32(t103))+944:], uint32(v4))
						m.fn209(v8+i32(944), v8+i32(152))
						t104 := int32(load32(m.memory[int64(uint32(v8))+944:]))
						v4 = v4 - t104
						goto l51
					}
				l50:
					v4 = v14 - v4
				l51:
					v13 = int64(v4)
					v4 = i32(1)
				}
			l46:
				v14 = v15 - v2
				v13 = v13 + v16
				t105 := int64(load64(m.memory[int64(uint32(v8))+152:]))
				v16 = t105
				{
					if v4 != 0 {
						goto l52
					}
					if uint64(v13+i64(-38)) < uint64(i64(-60)) {
						goto l52
					}
					if uint64(v16) > uint64(i64(0x20000000000000)) {
						goto l52
					}
					{
						{
							if v13 < i64(23) {
								goto l53
							}
							t106 := int64(load64(m.memory[uint32(int32(v13)<<3+i32(1107632)):]))
							m.fn1853(v8+i32(96), v16, i64(0), t106, i64(0))
							t107 := int64(load64(m.memory[int64(uint32(v8))+104:]))
							if t107 != i64(0) {
								goto l52
							}
							t108 := int64(load64(m.memory[int64(uint32(v8))+96:]))
							v19 = t108
							if uint64(v19) > uint64(i64(0x20000000000000)) {
								goto l52
							}
							v18 = float64(float64(uint64(v19)) * float64(1e+22))
							goto l54
						}
					l53:
						v4 = int32(v13)
						v18 = float64(uint64(v16))
						{
							if v13 < i64(0) {
								goto l55
							}
							t109 := math.Float64frombits(load64(m.memory[int64(uint32(v4<<3))+1131160:]))
							v18 = float64(t109 * v18)
							goto l54
						}
					l55:
						t110 := math.Float64frombits(load64(m.memory[uint32(i32(1131160)-v4<<3):]))
						v18 = float64(v18 / t110)
					}
				l54:
					p111 := v18
					if v11 == i32(45) {
						p111 = -v18
					}
					v18 = p111
					goto l40
				}
			l52:
				m.fn210(v8+i32(136), v13, v16)
				{
					{
						{
							if v4 != 0 {
								m.fn210(v8+i32(944), v13, v16+i64(1))
								{
									t113 := int64(load64(m.memory[int64(uint32(v8))+136:]))
									t114 := int64(load64(m.memory[int64(uint32(v8))+944:]))
									if t113 != t114 {
										goto l58
									}
									t115 := int32(load32(m.memory[int64(uint32(v8))+144:]))
									v17 = t115
									t116 := int32(load32(m.memory[int64(uint32(v8))+952:]))
									if v17 == t116 {
										goto l57
									}
								}
							l58:
								store32(m.memory[int64(uint32(v8))+144:], uint32(i32(-1)))
								goto l59
							}
							t112 := int32(load32(m.memory[int64(uint32(v8))+144:]))
							v17 = t112
							goto l57
						}
					l57:
						if v17 < i32(0) {
							goto l59
						}
						t117 := int64(load64(m.memory[int64(uint32(v8))+136:]))
						v13 = t117
						goto l60
					}
				l59:
					v4 = i32(0)
					memory_zero(m.memory, uint32(v8+i32(944)), uint32(i32(778)))
					t118 := v8
					var p119 int32
					if v11 == i32(45) {
						p119 = 1
					}
					m.memory[int64(uint32(t118))+1720] = byte(p119)
					v7 = v3
					v5 = v2
					switch v11 + i32(-43) {
					case 0, 2:
						m.fn207(v8+i32(88), v2, v3, i32(1))
						t120 := int32(load32(m.memory[int64(uint32(v8))+92:]))
						v7 = t120
						t121 := int32(load32(m.memory[int64(uint32(v8))+88:]))
						v5 = t121
						fallthrough
					default:
						m.fn1579(v8+i32(80), v5, v7)
						t122 := int32(load32(m.memory[int64(uint32(v8))+84:]))
						v7 = t122
						t123 := int32(load32(m.memory[int64(uint32(v8))+80:]))
						v5 = t123
						{
						l67:
							if v7 != 0 {
								goto l63
							}
							v7 = i32(0)
							store32(m.memory[int64(uint32(v8))+940:], uint32(i32(0)))
							store32(m.memory[int64(uint32(v8))+936:], uint32(v5))
							goto l64
						l63:
							{
								t124 := int32(m.memory[uint32(v5)])
								v10 = t124
								v17 = v10 + i32(-48)
								if uint32(v17&i32(255)) > uint32(i32(9)) {
									goto l65
								}
								{
									if uint32(v4) > uint32(i32(767)) {
										goto l66
									}
									m.memory[uint32(v8+i32(944)+v4)] = byte(v17)
									t125 := int32(load32(m.memory[int64(uint32(v8))+1712:]))
									v4 = t125
								}
							l66:
								t126 := v8
								v4 = v4 + i32(1)
								store32(m.memory[int64(uint32(t126))+1712:], uint32(v4))
								m.fn207(v8+i32(40), v5, v7, i32(1))
								t127 := int32(load32(m.memory[int64(uint32(v8))+44:]))
								v7 = t127
								t128 := int32(load32(m.memory[int64(uint32(v8))+40:]))
								v5 = t128
								goto l67
							}
						l65:
							store32(m.memory[int64(uint32(v8))+940:], uint32(v7))
							store32(m.memory[int64(uint32(v8))+936:], uint32(v5))
							if v10&i32(255) != i32(46) {
								goto l64
							}
							m.fn207(v8+i32(72), v5, v7, i32(1))
							t129 := int32(load32(m.memory[int64(uint32(v8))+76:]))
							t130 := v8
							v10 = t129
							store32(m.memory[int64(uint32(t130))+940:], uint32(v10))
							t131 := int32(load32(m.memory[int64(uint32(v8))+72:]))
							t132 := v8
							v5 = t131
							store32(m.memory[int64(uint32(t132))+936:], uint32(v5))
							{
								if v4 == 0 {
									goto l68
								}
								v7 = v10
								goto l72
							l68:
								m.fn1579(v8+i32(64), v5, v10)
								t133 := int32(load32(m.memory[int64(uint32(v8))+68:]))
								t134 := v8
								v7 = t133
								store32(m.memory[int64(uint32(t134))+940:], uint32(v7))
								t135 := int32(load32(m.memory[int64(uint32(v8))+64:]))
								t136 := v8
								v5 = t135
								store32(m.memory[int64(uint32(t136))+936:], uint32(v5))
								v4 = i32(0)
							}
						l72:
							{
								if uint32(v7) < uint32(i32(8)) {
									goto l70
								}
								if uint32(v4+i32(8)) >= uint32(i32(768)) {
									goto l70
								}
								t137 := int64(load64(m.memory[uint32(v5):]))
								v13 = t137
								t138 := v13 + i64(5063812098665367110)
								v13 = v13 + i64(-3472328296227680304)
								if (t138|v13)&i64(-0x7f7f7f7f7f7f7f80) != i64(0) {
									goto l70
								}
								{
									if uint32(v4) >= uint32(i32(769)) {
										goto l71
									}
									store64(m.memory[uint32(v8+i32(944)+v4):], uint64(v13))
									t139 := int32(load32(m.memory[int64(uint32(v8))+1712:]))
									t140 := v8
									v4 = t139 + i32(8)
									store32(m.memory[int64(uint32(t140))+1712:], uint32(v4))
									m.fn207(v8+i32(48), v5, v7, i32(8))
									t141 := int32(load32(m.memory[int64(uint32(v8))+52:]))
									v7 = t141
									t142 := int32(load32(m.memory[int64(uint32(v8))+48:]))
									v5 = t142
									goto l72
								}
							l71:
							}
							m.fn151(v4, i32(768), i32(768), i32(1088152))
							panic("unreachable")
						l70:
							store32(m.memory[int64(uint32(v8))+940:], uint32(v7))
							store32(m.memory[int64(uint32(v8))+936:], uint32(v5))
						l76:
							if v7 != 0 {
								t143 := int32(m.memory[uint32(v5)])
								v17 = t143 + i32(-48)
								if uint32(v17&i32(255)) > uint32(i32(9)) {
									goto l74
								}
								{
									if uint32(v4) > uint32(i32(767)) {
										goto l75
									}
									m.memory[uint32(v8+i32(944)+v4)] = byte(v17)
									t144 := int32(load32(m.memory[int64(uint32(v8))+1712:]))
									v4 = t144
								}
							l75:
								t145 := v8
								v4 = v4 + i32(1)
								store32(m.memory[int64(uint32(t145))+1712:], uint32(v4))
								m.fn207(v8+i32(56), v5, v7, i32(1))
								t146 := int32(load32(m.memory[int64(uint32(v8))+60:]))
								v7 = t146
								t147 := int32(load32(m.memory[int64(uint32(v8))+56:]))
								v5 = t147
								goto l76
							}
							v7 = i32(0)
							goto l74
						l74:
							store32(m.memory[int64(uint32(v8))+936:], uint32(v5))
							store32(m.memory[int64(uint32(v8))+940:], uint32(v7))
							store32(m.memory[int64(uint32(v8))+1716:], uint32(v7-v10))
						}
					l64:
						if v4 != 0 {
							v17 = v3 - v7
							{
								if uint32(v3) < uint32(v7) {
									m.fn151(i32(0), v17, v3, i32(1088136))
									panic("unreachable")
								}
								v17 = v2 + v17
								v10 = i32(0)
							l81:
								if v17 == v2 {
									goto l80
								}
								{
									v17 = v17 + i32(-1)
									t148 := int32(m.memory[uint32(v17)])
									switch t148 + i32(-46) {
									case 0:
										goto l81
									default:
										goto l80
									case 2:
										v10 = v10 + i32(1)
										goto l81
									}
								}
							l80:
								t149 := int32(load32(m.memory[int64(uint32(v8))+1716:]))
								store32(m.memory[int64(uint32(v8))+1716:], uint32(t149+v4))
								t150 := v8
								v4 = v4 - v10
								store32(m.memory[int64(uint32(t150))+1712:], uint32(v4))
								if uint32(v4) <= uint32(i32(768)) {
									goto l78
								}
								store32(m.memory[int64(uint32(v8))+1712:], uint32(i32(768)))
								m.memory[int64(uint32(v8))+1721] = byte(i32(1))
								v4 = i32(768)
								goto l78
							}
						}
						v4 = i32(0)
						goto l78
					l78:
						{
							if v7 == 0 {
								goto l83
							}
							t151 := int32(m.memory[uint32(v5)])
							if t151&i32(223) != i32(69) {
								goto l83
							}
							m.fn207(v8+i32(32), v5, v7, i32(1))
							t152 := int32(load32(m.memory[int64(uint32(v8))+36:]))
							t153 := v8
							v7 = t152
							store32(m.memory[int64(uint32(t153))+940:], uint32(v7))
							t154 := int32(load32(m.memory[int64(uint32(v8))+32:]))
							t155 := v8
							v5 = t154
							store32(m.memory[int64(uint32(t155))+936:], uint32(v5))
							{
								if v7 == 0 {
									goto l84
								}
								{
									t156 := int32(m.memory[uint32(v5)])
									switch t156 + i32(-43) {
									default:
										goto l84
									case 2:
										m.fn207(v8+i32(16), v5, v7, i32(1))
										t157 := int64(load64(m.memory[int64(uint32(v8))+16:]))
										store64(m.memory[int64(uint32(v8))+936:], uint64(t157))
										store32(m.memory[int64(uint32(v8))+152:], uint32(i32(0)))
										m.fn213(v8+i32(936), v8+i32(152))
										t158 := int32(load32(m.memory[int64(uint32(v8))+152:]))
										v7 = i32(0) - t158
										goto l87
									case 0:
										m.fn207(v8+i32(24), v5, v7, i32(1))
										t159 := int64(load64(m.memory[int64(uint32(v8))+24:]))
										store64(m.memory[int64(uint32(v8))+936:], uint64(t159))
									}
								}
							l84:
								store32(m.memory[int64(uint32(v8))+152:], uint32(i32(0)))
								m.fn213(v8+i32(936), v8+i32(152))
								t160 := int32(load32(m.memory[int64(uint32(v8))+152:]))
								v7 = t160
							}
						l87:
							t161 := int32(load32(m.memory[int64(uint32(v8))+1716:]))
							store32(m.memory[int64(uint32(v8))+1716:], uint32(t161+v7))
						}
					l83:
						p162 := i32(19)
						if uint32(v4) > uint32(i32(19)) {
							p162 = v4
						}
						v7 = p162
					l89:
						if v7 == v4 {
							goto l88
						}
						m.memory[uint32(v8+i32(944)+v4)] = byte(i32(0))
						v4 = v4 + i32(1)
						goto l89
					l88:
						memory_copy(m.memory, uint32(v8+i32(152)), uint32(v8+i32(944)), uint32(i32(780)))
						v17 = i32(0)
						v13 = i64(0)
						{
							t163 := int32(load32(m.memory[int64(uint32(v8))+920:]))
							if t163 == 0 {
								goto l90
							}
							t164 := int32(load32(m.memory[int64(uint32(v8))+924:]))
							v4 = t164
							if v4 < i32(-324) {
								goto l90
							}
							v17 = i32(2047)
							if v4 > i32(309) {
								goto l90
							}
							v5 = i32(0)
						l103:
							{
								if v4 > i32(0) {
									goto l91
								}
							l101:
								{
									{
										{
											if v4 > i32(0) {
												goto l92
											}
											if v4 != 0 {
												v7 = i32(60)
												v4 = i32(0) - v4
												if uint32(v4) > uint32(i32(18)) {
													goto l96
												}
												t166 := int32(m.memory[int64(uint32(v4))+1108132])
												v7 = t166
												goto l96
											}
											t165 := int32(m.memory[int64(uint32(v8))+152])
											v4 = t165
											if uint32(v4) <= uint32(i32(4)) {
												goto l94
											}
										}
									l92:
										v4 = v5 + i32(-1)
									l99:
										{
											if v4 >= i32(-1022) {
												if v4+i32(1023) > i32(2046) {
													goto l90
												}
												m.fn215(v8+i32(152), i32(53))
												{
													t171 := m.fn216(v8 + i32(152))
													v16 = t171
													if uint64(v16) < uint64(i64(0x20000000000000)) {
														goto l100
													}
													m.fn214(v8+i32(152), i32(1))
													t172 := m.fn216(v8 + i32(152))
													v16 = t172
													if v4+i32(1024) > i32(2046) {
														goto l90
													}
													v4 = v4 + i32(1)
												}
											l100:
												v13 = v16 & i64(0xfffffffffffff)
												p173 := i32(1023)
												if uint64(v16) < uint64(i64(0x10000000000000)) {
													p173 = i32(1022)
												}
												v17 = p173 + v4
												goto l90
											}
											t169 := v8 + i32(152)
											v7 = i32(-1022) - v4
											p170 := i32(60)
											if uint32(v7) < uint32(i32(60)) {
												p170 = v7
											}
											v7 = p170
											m.fn214(t169, v7)
											v4 = v7 + v4
											goto l99
										}
									l94:
										p167 := i32(1)
										if uint32(v4) < uint32(i32(2)) {
											p167 = i32(2)
										}
										v7 = p167
									}
								l96:
									m.fn215(v8+i32(152), v7)
									t168 := int32(load32(m.memory[int64(uint32(v8))+924:]))
									v4 = t168
									if v4 <= i32(2047) {
										v5 = v5 - v7
										goto l101
									}
									v17 = i32(2047)
									goto l90
								}
							l91:
								v7 = i32(60)
								{
									if uint32(v4) > uint32(i32(18)) {
										goto l102
									}
									t174 := int32(m.memory[int64(uint32(v4))+1108132])
									v7 = t174
								}
							l102:
								m.fn214(v8+i32(152), v7)
								v5 = v7 + v5
								t175 := int32(load32(m.memory[int64(uint32(v8))+924:]))
								v4 = t175
								if v4 >= i32(-2047) {
									goto l103
								}
							}
							v17 = i32(0)
						}
					l90:
						store32(m.memory[int64(uint32(v8))+144:], uint32(v17))
					}
				}
			l60:
				v13 = int64(uint32(v17))<<52 | v13
				p176 := v13
				if v11 == i32(45) {
					p176 = v13 | i64(-0x8000000000000000)
				}
				v18 = math.Float64frombits(uint64(p176))
			}
		l40:
			if v14 != v3 {
				goto l27
			}
			t177 := int32(m.memory[int64(uint32(v1))+16])
			m.fn1100(v8+i32(112), v18, v9, t177)
			t178 := int32(m.memory[int64(uint32(v8))+112])
			if t178 != i32(255) {
				t195 := int64(load64(m.memory[int64(uint32(v8))+128:]))
				store64(m.memory[int64(uint32(v0))+24:], uint64(t195))
				t196 := int64(load64(m.memory[int64(uint32(v8))+120:]))
				store64(m.memory[int64(uint32(v0))+16:], uint64(t196))
				t197 := int64(load64(m.memory[int64(uint32(v8))+112:]))
				store64(m.memory[int64(uint32(v0))+8:], uint64(t197))
				store32(m.memory[uint32(v0):], uint32(i32(0)))
				goto l12
			}
		}
	l27:
		{
			if v6 == 0 {
				m.fn1588(v8+i32(944), v2, v3)
				t186 := int32(load32(m.memory[int64(uint32(v8))+952:]))
				v4 = t186
				t187 := int32(load32(m.memory[int64(uint32(v8))+948:]))
				v7 = t187
				{
					t188 := int32(load32(m.memory[int64(uint32(v8))+944:]))
					v6 = t188
					if v6 == i32(-1) {
						m.fn377(v8+i32(136), v7, v4)
						t191 := int32(load32(m.memory[int64(uint32(v8))+144:]))
						store32(m.memory[int64(uint32(v8))+163:], uint32(t191))
						t192 := int64(load64(m.memory[int64(uint32(v8))+136:]))
						store64(m.memory[int64(uint32(v8))+155:], uint64(t192))
						m.memory[int64(uint32(v0))+8] = byte(i32(2))
						t193 := int64(load64(m.memory[int64(uint32(v8))+152:]))
						store64(m.memory[int64(uint32(v0))+9:], uint64(t193))
						t194 := int64(load64(m.memory[int64(uint32(v8))+159:]))
						store64(m.memory[int64(uint32(v0))+16:], uint64(t194))
						store32(m.memory[uint32(v0):], uint32(i32(0)))
						goto l12
					}
					t189 := int32(load32(m.memory[int64(uint32(v8))+964:]))
					store32(m.memory[int64(uint32(v0))+24:], uint32(t189))
					t190 := int64(load64(m.memory[int64(uint32(v8))+956:]))
					store64(m.memory[int64(uint32(v0))+16:], uint64(t190))
					store32(m.memory[int64(uint32(v0))+12:], uint32(v4))
					store32(m.memory[int64(uint32(v0))+8:], uint32(v7))
					store32(m.memory[int64(uint32(v0))+4:], uint32(v6))
					store32(m.memory[uint32(v0):], uint32(i32(1)))
					goto l12
				}
			}
			m.fn1588(v8+i32(944), v2, v3)
			t179 := int32(load32(m.memory[int64(uint32(v8))+952:]))
			v4 = t179
			t180 := int32(load32(m.memory[int64(uint32(v8))+948:]))
			v7 = t180
			{
				t181 := int32(load32(m.memory[int64(uint32(v8))+944:]))
				v6 = t181
				if v6 == i32(-1) {
					m.fn217(v8+i32(152), v7, v4)
					t184 := int32(m.memory[int64(uint32(v8))+152])
					if t184 == 0 {
						t198 := math.Float64frombits(load64(m.memory[int64(uint32(v8))+160:]))
						store64(m.memory[int64(uint32(v8))+944:], math.Float64bits(t198))
						m.fn97(i32(1087712), i32(46), v8+i32(944), i32(1087696), i32(1087596))
						panic("unreachable")
					}
					t185 := int32(m.memory[int64(uint32(v8))+153])
					m.memory[int64(uint32(v0))+8] = byte(t185)
					store64(m.memory[uint32(v0):], uint64(i64(-0x7fffffeaffffffff)))
					goto l12
				}
				t182 := int32(load32(m.memory[int64(uint32(v8))+964:]))
				store32(m.memory[int64(uint32(v0))+24:], uint32(t182))
				t183 := int64(load64(m.memory[int64(uint32(v8))+956:]))
				store64(m.memory[int64(uint32(v0))+16:], uint64(t183))
				store32(m.memory[int64(uint32(v0))+12:], uint32(v4))
				store32(m.memory[int64(uint32(v0))+8:], uint32(v7))
				store32(m.memory[int64(uint32(v0))+4:], uint32(v6))
				store32(m.memory[uint32(v0):], uint32(i32(1)))
				goto l12
			}
		}
	l4:
		m.fn12(v8+i32(944), v6, v7)
		t199 := int32(load32(m.memory[int64(uint32(v8))+948:]))
		t200 := int32(load32(m.memory[int64(uint32(v8))+944:]))
		t201 := v8 + i32(944)
		v4 = t200
		p202 := t199
		if v4 != 0 {
			p202 = i32(1099879)
		}
		t203 := int32(load32(m.memory[int64(uint32(v8))+952:]))
		p204 := t203
		if v4 != 0 {
			p204 = i32(12)
		}
		m.fn377(t201, p202, p204)
		store64(m.memory[uint32(v0):], uint64(i64(-0x7fffffddffffffff)))
		t205 := int64(load64(m.memory[int64(uint32(v8))+944:]))
		store64(m.memory[int64(uint32(v0))+8:], uint64(t205))
		t206 := int32(load32(m.memory[int64(uint32(v8))+952:]))
		store32(m.memory[int64(uint32(v0))+16:], uint32(t206))
	}
l12:
	m.g0 = v8 + i32(1728)
}
func (m *Module) fn1143(v0 int32) {
	var v1 int32
	t0 := m.g0
	v1 = t0 - i32(16)
	m.g0 = v1
	t1 := int32(load32(m.memory[uint32(v0):]))
	m.fn273(v1+i32(8), v0, t1, i32(1), i32(4), i32(28))
	{
		t2 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		v0 = t2
		if v0 == i32(-1) {
			goto l0
		}
		t3 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		m.fn2(v0, t3)
		panic("unreachable")
	}
l0:
	m.g0 = v1 + i32(16)
}
func (m *Module) fn1144(v0 int32) {
	var v1 int32
	t0 := m.g0
	v1 = t0 - i32(16)
	m.g0 = v1
	t1 := int32(load32(m.memory[uint32(v0):]))
	m.fn273(v1+i32(8), v0, t1, i32(1), i32(4), i32(40))
	{
		t2 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		v0 = t2
		if v0 == i32(-1) {
			goto l0
		}
		t3 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		m.fn2(v0, t3)
		panic("unreachable")
	}
l0:
	m.g0 = v1 + i32(16)
}
func (m *Module) fn1145(v0 int32) {
	var v1 int32
	t0 := m.g0
	v1 = t0 - i32(16)
	m.g0 = v1
	t1 := int32(load32(m.memory[uint32(v0):]))
	m.fn273(v1+i32(8), v0, t1, i32(1), i32(8), i32(16))
	{
		t2 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		v0 = t2
		if v0 == i32(-1) {
			goto l0
		}
		t3 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		m.fn2(v0, t3)
		panic("unreachable")
	}
l0:
	m.g0 = v1 + i32(16)
}
func (m *Module) fn1146(v0 int32) {
	var v1 int32
	t0 := m.g0
	v1 = t0 - i32(16)
	m.g0 = v1
	t1 := int32(load32(m.memory[uint32(v0):]))
	m.fn273(v1+i32(8), v0, t1, i32(1), i32(4), i32(36))
	{
		t2 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		v0 = t2
		if v0 == i32(-1) {
			goto l0
		}
		t3 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		m.fn2(v0, t3)
		panic("unreachable")
	}
l0:
	m.g0 = v1 + i32(16)
}
func (m *Module) fn1147(v0 int32) {
	var v1 int32
	t0 := m.g0
	v1 = t0 - i32(16)
	m.g0 = v1
	t1 := int32(load32(m.memory[uint32(v0):]))
	m.fn273(v1+i32(8), v0, t1, i32(1), i32(8), i32(56))
	{
		t2 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		v0 = t2
		if v0 == i32(-1) {
			goto l0
		}
		t3 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		m.fn2(v0, t3)
		panic("unreachable")
	}
l0:
	m.g0 = v1 + i32(16)
}
func (m *Module) fn1148(v0 int32) {
	var v1 int32
	t0 := m.g0
	v1 = t0 - i32(16)
	m.g0 = v1
	t1 := int32(load32(m.memory[uint32(v0):]))
	m.fn273(v1+i32(8), v0, t1, i32(1), i32(4), i32(72))
	{
		t2 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		v0 = t2
		if v0 == i32(-1) {
			goto l0
		}
		t3 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		m.fn2(v0, t3)
		panic("unreachable")
	}
l0:
	m.g0 = v1 + i32(16)
}
func (m *Module) fn1149(v0 int32) {
	var v1 int32
	t0 := m.g0
	v1 = t0 - i32(16)
	m.g0 = v1
	t1 := int32(load32(m.memory[uint32(v0):]))
	m.fn273(v1+i32(8), v0, t1, i32(1), i32(8), i32(40))
	{
		t2 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		v0 = t2
		if v0 == i32(-1) {
			goto l0
		}
		t3 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		m.fn2(v0, t3)
		panic("unreachable")
	}
l0:
	m.g0 = v1 + i32(16)
}
func (m *Module) fn1150(v0 int32) {
	var v1 int32
	t0 := m.g0
	v1 = t0 - i32(16)
	m.g0 = v1
	t1 := int32(load32(m.memory[uint32(v0):]))
	m.fn273(v1+i32(8), v0, t1, i32(1), i32(1), i32(3))
	{
		t2 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		v0 = t2
		if v0 == i32(-1) {
			goto l0
		}
		t3 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		m.fn2(v0, t3)
		panic("unreachable")
	}
l0:
	m.g0 = v1 + i32(16)
}
func (m *Module) fn1151(v0 int32) {
	var v1 int32
	t0 := m.g0
	v1 = t0 - i32(16)
	m.g0 = v1
	t1 := int32(load32(m.memory[uint32(v0):]))
	m.fn273(v1+i32(8), v0, t1, i32(1), i32(8), i32(64))
	{
		t2 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		v0 = t2
		if v0 == i32(-1) {
			goto l0
		}
		t3 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		m.fn2(v0, t3)
		panic("unreachable")
	}
l0:
	m.g0 = v1 + i32(16)
}
func (m *Module) fn1152(v0 int32) {
	var v1 int32
	t0 := m.g0
	v1 = t0 - i32(16)
	m.g0 = v1
	t1 := int32(load32(m.memory[uint32(v0):]))
	m.fn273(v1+i32(8), v0, t1, i32(1), i32(8), i32(240))
	{
		t2 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		v0 = t2
		if v0 == i32(-1) {
			goto l0
		}
		t3 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		m.fn2(v0, t3)
		panic("unreachable")
	}
l0:
	m.g0 = v1 + i32(16)
}
func (m *Module) fn1153(v0 int32) {
	var v1 int32
	t0 := m.g0
	v1 = t0 - i32(16)
	m.g0 = v1
	t1 := int32(load32(m.memory[uint32(v0):]))
	m.fn273(v1+i32(8), v0, t1, i32(1), i32(2), i32(2))
	{
		t2 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		v0 = t2
		if v0 == i32(-1) {
			goto l0
		}
		t3 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		m.fn2(v0, t3)
		panic("unreachable")
	}
l0:
	m.g0 = v1 + i32(16)
}
func (m *Module) fn1154(v0 int32) {
	var v1 int32
	t0 := m.g0
	v1 = t0 - i32(16)
	m.g0 = v1
	t1 := int32(load32(m.memory[uint32(v0):]))
	m.fn273(v1+i32(8), v0, t1, i32(1), i32(8), i32(8))
	{
		t2 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		v0 = t2
		if v0 == i32(-1) {
			goto l0
		}
		t3 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		m.fn2(v0, t3)
		panic("unreachable")
	}
l0:
	m.g0 = v1 + i32(16)
}
func (m *Module) fn1155(v0, v1, v2, v3, v4, v5 int32) {
	var v6, v7, v8, v9, v10, v11, v12, v13, v14, v15 int32
	var v16 int64
	t0 := m.g0
	v6 = t0 - i32(128)
	m.g0 = v6
	{
		if v4 == v2 {
			goto l0
		}
		if uint32(v4) >= uint32(v2) {
			goto l1
		}
		goto l2
	l0:
		if uint32(v5) < uint32(v3) {
			goto l2
		}
	l1:
		m.memory[int64(uint32(v6))+56] = byte(i32(8))
		t1 := v6 + i32(28)
		t2 := v6 + i32(56)
		v7 = v5 - v3 + i32(1)
		m.fn178(t1, t2, v7*(v4-v2+i32(1)))
		store32(m.memory[int64(uint32(v6))+52:], uint32(v5))
		store32(m.memory[int64(uint32(v6))+48:], uint32(v4))
		store32(m.memory[int64(uint32(v6))+44:], uint32(v3))
		store32(m.memory[int64(uint32(v6))+40:], uint32(v2))
		{
			t3 := int32(load32(m.memory[int64(uint32(v1))+12:]))
			t4 := v2
			v8 = t3
			p5 := v8
			if uint32(v2) > uint32(v8) {
				p5 = t4
			}
			v9 = p5
			t6 := int32(load32(m.memory[int64(uint32(v1))+20:]))
			t7 := v9
			t8 := v4
			v10 = t6
			p9 := v10
			if uint32(v4) < uint32(v10) {
				p9 = t8
			}
			v11 = p9
			if uint32(t7) > uint32(v11) {
				goto l3
			}
			t10 := int32(load32(m.memory[int64(uint32(v1))+16:]))
			t11 := v3
			v4 = t10
			p12 := v4
			if uint32(v3) > uint32(v4) {
				p12 = t11
			}
			v12 = p12
			t13 := int32(load32(m.memory[int64(uint32(v1))+24:]))
			t14 := v12
			t15 := v5
			v10 = t13
			p16 := v10
			if uint32(v5) < uint32(v10) {
				p16 = t15
			}
			v5 = p16
			if uint32(t14) > uint32(v5) {
				goto l3
			}
			{
				v10 = v10 + i32(1)
				if v10 == v4 {
					goto l4
				}
				t17 := int32(load32(m.memory[int64(uint32(v1))+8:]))
				v13 = t17
				if v13 == 0 {
					goto l4
				}
				t18 := int32(load32(m.memory[int64(uint32(v6))+36:]))
				t19 := v7
				v14 = t18
				p20 := i32(0)
				if v14 != 0 {
					p20 = t19
				}
				v7 = p20
				if v7 == 0 {
					m.fn91(i32(1075037), i32(55), i32(1079636))
					panic("unreachable")
				}
				t21 := int32(load32(m.memory[int64(uint32(v1))+4:]))
				v1 = t21
				t22 := int32(load32(m.memory[int64(uint32(v6))+32:]))
				v15 = t22
				store64(m.memory[int64(uint32(v6))+96:], uint64(i64(0)))
				store32(m.memory[int64(uint32(v6))+92:], uint32(v9-v2))
				t23 := v6
				v11 = v11 + i32(1)
				store32(m.memory[int64(uint32(t23))+88:], uint32(v11-v2))
				store32(m.memory[int64(uint32(v6))+84:], uint32(v7))
				store32(m.memory[int64(uint32(v6))+80:], uint32(v14))
				store32(m.memory[int64(uint32(v6))+76:], uint32(v15))
				t24 := v6
				v2 = v9 - v8
				store32(m.memory[int64(uint32(t24))+72:], uint32(v2))
				store32(m.memory[int64(uint32(v6))+68:], uint32(v11-v8))
				store32(m.memory[int64(uint32(v6))+64:], uint32(v10-v4))
				store32(m.memory[int64(uint32(v6))+60:], uint32(v13))
				store32(m.memory[int64(uint32(v6))+56:], uint32(v1))
				v1 = v5 + i32(1)
				v8 = v1 - v3
				t25 := v8
				v10 = v12 - v3
				t26 := t25 - v10
				v1 = v1 - v4
				t27 := v1
				v5 = v12 - v4
				v13 = t27 - v5
				var p28 int32
				if t26 != v13 {
					p28 = 1
				}
				v11 = p28
			l16:
				{
					{
						if v2 != 0 {
							goto l6
						}
						t29 := int32(load32(m.memory[int64(uint32(v6))+68:]))
						v2 = t29
						if v2 == 0 {
							goto l3
						}
						store32(m.memory[int64(uint32(v6))+68:], uint32(v2+i32(-1)))
						m.fn1156(v6, v6+i32(56))
						t30 := int32(load32(m.memory[int64(uint32(v6))+4:]))
						v9 = t30
						t31 := int32(load32(m.memory[uint32(v6):]))
						v4 = t31
						goto l7
					}
				l6:
					store32(m.memory[int64(uint32(v6))+72:], uint32(i32(0)))
					{
						t32 := int32(load32(m.memory[int64(uint32(v6))+68:]))
						v4 = t32
						if uint32(v4) > uint32(v2) {
							goto l8
						}
						if v4 == 0 {
							goto l3
						}
						m.fn1157(v6+i32(16), v6+i32(56), v4+i32(-1))
						goto l3
					}
				l8:
					store32(m.memory[int64(uint32(v6))+68:], uint32(v4+(v2^i32(-1))))
					m.fn1157(v6+i32(8), v6+i32(56), v2)
					t33 := int32(load32(m.memory[int64(uint32(v6))+12:]))
					v9 = t33
					t34 := int32(load32(m.memory[int64(uint32(v6))+8:]))
					v4 = t34
				}
			l7:
				if v4 == 0 {
					goto l3
				}
				{
					{
						t35 := int32(load32(m.memory[int64(uint32(v6))+92:]))
						v2 = t35
						if v2 != 0 {
							goto l9
						}
						t36 := int32(load32(m.memory[int64(uint32(v6))+88:]))
						v2 = t36
						if v2 == 0 {
							goto l3
						}
						store32(m.memory[int64(uint32(v6))+88:], uint32(v2+i32(-1)))
						t37 := int32(load32(m.memory[int64(uint32(v6))+80:]))
						v2 = t37
						if v2 == 0 {
							goto l3
						}
						t38 := int32(load32(m.memory[int64(uint32(v6))+84:]))
						t39 := v2
						v3 = t38
						p40 := v2
						if uint32(v3) < uint32(v2) {
							p40 = v3
						}
						v3 = p40
						v12 = t39 - v3
						t41 := int32(load32(m.memory[int64(uint32(v6))+76:]))
						v2 = t41
						v7 = v2 + v3*i32(24)
						goto l10
					}
				l9:
					store32(m.memory[int64(uint32(v6))+92:], uint32(i32(0)))
					t42 := int32(load32(m.memory[int64(uint32(v6))+88:]))
					v3 = t42
					if uint32(v3) <= uint32(v2) {
						goto l3
					}
					store32(m.memory[int64(uint32(v6))+88:], uint32(v3+(v2^i32(-1))))
					t43 := int32(load32(m.memory[int64(uint32(v6))+84:]))
					t44 := int64(uint32(v2))
					v2 = t43
					v16 = t44 * int64(uint32(v2))
					if int32(int64(uint64(v16)>>32)) != 0 {
						goto l3
					}
					t45 := int32(load32(m.memory[int64(uint32(v6))+80:]))
					v3 = t45
					t46 := v3
					v7 = int32(v16)
					if uint32(t46) <= uint32(v7) {
						goto l3
					}
					v3 = v3 - v7
					t48 := v3
					p47 := v2
					if uint32(v3) < uint32(v2) {
						p47 = v3
					}
					v3 = p47
					v12 = t48 - v3
					t49 := int32(load32(m.memory[int64(uint32(v6))+76:]))
					v2 = t49 + v7*i32(24)
					v7 = v2 + v3*i32(24)
				}
			l10:
				store32(m.memory[int64(uint32(v6))+80:], uint32(v12))
				store32(m.memory[int64(uint32(v6))+76:], uint32(v7))
				if v2 == 0 {
					goto l3
				}
				if uint32(v1) < uint32(v5) {
					goto l11
				}
				if uint32(v1) > uint32(v9) {
					goto l11
				}
				if uint32(v8) < uint32(v10) {
					goto l12
				}
				if uint32(v8) > uint32(v3) {
					goto l12
				}
				if v11 != 0 {
					m.fn91(i32(1086764), i32(105), i32(1079684))
					panic("unreachable")
				}
				v2 = v2 + v10*i32(24)
				v3 = v4 + v5*i32(24)
				v4 = v13
			l15:
				{
					if v4 == 0 {
						t53 := int32(load32(m.memory[int64(uint32(v6))+72:]))
						v2 = t53
						goto l16
					}
					m.fn219(v6+i32(104), v3)
					m.fn182(v2)
					t50 := int64(load64(m.memory[int64(uint32(v6))+120:]))
					store64(m.memory[int64(uint32(v2))+16:], uint64(t50))
					t51 := int64(load64(m.memory[int64(uint32(v6))+112:]))
					store64(m.memory[int64(uint32(v2))+8:], uint64(t51))
					t52 := int64(load64(m.memory[int64(uint32(v6))+104:]))
					store64(m.memory[uint32(v2):], uint64(t52))
					v4 = v4 + i32(-1)
					v2 = v2 + i32(24)
					v3 = v3 + i32(24)
					goto l15
				}
			}
		l4:
			m.fn91(i32(1075037), i32(55), i32(1079620))
			panic("unreachable")
		l11:
			m.fn151(v5, v1, v9, i32(1079652))
			panic("unreachable")
		l12:
			m.fn151(v10, v8, v3, i32(1079668))
			panic("unreachable")
		}
	l3:
		t54 := int32(load32(m.memory[int64(uint32(v6))+52:]))
		store32(m.memory[int64(uint32(v0))+24:], uint32(t54))
		t55 := int64(load64(m.memory[int64(uint32(v6))+44:]))
		store64(m.memory[int64(uint32(v0))+16:], uint64(t55))
		t56 := int64(load64(m.memory[int64(uint32(v6))+36:]))
		store64(m.memory[int64(uint32(v0))+8:], uint64(t56))
		t57 := int64(load64(m.memory[int64(uint32(v6))+28:]))
		store64(m.memory[uint32(v0):], uint64(t57))
		m.g0 = v6 + i32(128)
		return
	}
l2:
	m.fn91(i32(1079566), i32(41), i32(1079588))
	panic("unreachable")
}
func (m *Module) fn1156(v0, v1 int32) {
	var v2, v3, v4 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	{
		{
			t1 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			v3 = t1
			if v3 != 0 {
				goto l0
			}
			v3 = i32(0)
			goto l1
		}
	l0:
		t2 := int32(load32(m.memory[uint32(v1):]))
		t3 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		t4 := v2
		t5 := v3
		v4 = t3
		p6 := v3
		if uint32(v4) < uint32(v3) {
			p6 = v4
		}
		m.fn1158(t4, t2, t5, p6, i32(1100308))
		t7 := int32(load32(m.memory[int64(uint32(v2))+4:]))
		v4 = t7
		t8 := int32(load32(m.memory[uint32(v2):]))
		v3 = t8
		t9 := int64(load64(m.memory[int64(uint32(v2))+8:]))
		store64(m.memory[uint32(v1):], uint64(t9))
	}
l1:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v4))
	store32(m.memory[uint32(v0):], uint32(v3))
	m.g0 = v2 + i32(16)
}
func (m *Module) fn1157(v0, v1, v2 int32) {
	var v3 int32
	var v4 int64
	var v5, v6 int32
	t0 := m.g0
	v3 = t0 - i32(16)
	m.g0 = v3
	{
		t1 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		t2 := int64(uint32(v2))
		v2 = t1
		v4 = t2 * int64(uint32(v2))
		if int32(int64(uint64(v4)>>32)) != 0 {
			goto l0
		}
		t3 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		v5 = t3
		t4 := v5
		v6 = int32(v4)
		if uint32(t4) <= uint32(v6) {
			goto l0
		}
		t5 := int32(load32(m.memory[uint32(v1):]))
		t6 := v3
		t7 := t5 + v6*i32(24)
		v6 = v5 - v6
		t9 := v6
		p8 := v2
		if uint32(v6) < uint32(v2) {
			p8 = v6
		}
		m.fn1158(t6, t7, t9, p8, i32(1086488))
		t10 := int32(load32(m.memory[int64(uint32(v3))+4:]))
		v2 = t10
		t11 := int32(load32(m.memory[uint32(v3):]))
		v6 = t11
		t12 := int64(load64(m.memory[int64(uint32(v3))+8:]))
		store64(m.memory[uint32(v1):], uint64(t12))
		goto l1
	}
l0:
	v6 = i32(0)
	store32(m.memory[int64(uint32(v1))+4:], uint32(i32(0)))
l1:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v2))
	store32(m.memory[uint32(v0):], uint32(v6))
	m.g0 = v3 + i32(16)
}
func (m *Module) fn1158(v0, v1, v2, v3, v4 int32) {
	if uint32(v2) >= uint32(v3) {
		goto l0
	}
	m.fn91(i32(1301172), i32(19), v4)
	panic("unreachable")
l0:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
	store32(m.memory[uint32(v0):], uint32(v1))
	store32(m.memory[int64(uint32(v0))+12:], uint32(v2-v3))
	store32(m.memory[int64(uint32(v0))+8:], uint32(v1+v3*i32(24)))
}
func (m *Module) fn1159(v0, v1 int32) {
	var v2, v3, v4, v5, v6, v7, v8, v9, v10, v11, v12 int32
	var v13, v14 int64
	var v15, v16, v17 int32
	t0 := m.g0
	v2 = t0 - i32(112)
	m.g0 = v2
	{
		t1 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		v3 = t1
		if v3 == 0 {
			goto l0
		}
		t2 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		v4 = t2
		t3 := v4
		v5 = v3 << 5
		v6 = t3 + v5
		v7 = i32(-1)
		v8 = i32(0)
		v9 = v4
		v10 = i32(0)
		v11 = i32(-1)
	l2:
		{
			if v5 == 0 {
				m.memory[int64(uint32(v2))+40] = byte(i32(9))
				t12 := v2 + i32(16)
				v13 = int64(uint32(v8 - v7 + i32(1)))
				v14 = v13 * int64(uint32(v10-v11+i32(1)))
				p13 := int32(v14)
				if int32(int64(uint64(v14)>>32)) != 0 {
					p13 = i32(-1)
				}
				v5 = p13
				m.fn59(t12, v5, i32(8), i32(24))
				v15 = i32(0)
				store32(m.memory[int64(uint32(v2))+108:], uint32(i32(0)))
				t14 := int32(load32(m.memory[int64(uint32(v2))+20:]))
				t15 := v2
				v9 = t14
				store32(m.memory[int64(uint32(t15))+104:], uint32(v9))
				t16 := int32(load32(m.memory[int64(uint32(v2))+16:]))
				t17 := v2
				v12 = t16
				store32(m.memory[int64(uint32(t17))+100:], uint32(v12))
				{
					if uint32(v5) <= uint32(v12) {
						goto l3
					}
					m.fn62(v2+i32(100), i32(0), v5, i32(8), i32(24))
					t18 := int32(load32(m.memory[int64(uint32(v2))+108:]))
					v15 = t18
					t19 := int32(load32(m.memory[int64(uint32(v2))+104:]))
					v9 = t19
				}
			l3:
				v9 = v9 + v15*i32(24)
				p20 := i32(1)
				if uint32(v5) > uint32(i32(1)) {
					p20 = v5
				}
				v16 = p20
				v12 = v16 + i32(-1)
				v16 = v15 + v16
				v15 = v16 + i32(-1)
			l18:
				{
					if v12 != 0 {
						m.memory[int64(uint32(v2))+64] = byte(i32(9))
						t43 := int64(load64(m.memory[int64(uint32(v2))+80:]))
						store64(m.memory[int64(uint32(v9))+16:], uint64(t43))
						t44 := int64(load64(m.memory[int64(uint32(v2))+72:]))
						store64(m.memory[int64(uint32(v9))+8:], uint64(t44))
						t45 := int64(load64(m.memory[int64(uint32(v2))+64:]))
						store64(m.memory[uint32(v9):], uint64(t45))
						v12 = v12 + i32(-1)
						v9 = v9 + i32(24)
						goto l18
					}
					{
						if v5 != 0 {
							goto l5
						}
						m.fn964(v2 + i32(40))
						goto l6
					l5:
						t21 := int64(load64(m.memory[int64(uint32(v2))+56:]))
						store64(m.memory[int64(uint32(v9))+16:], uint64(t21))
						t22 := int64(load64(m.memory[int64(uint32(v2))+48:]))
						store64(m.memory[int64(uint32(v9))+8:], uint64(t22))
						t23 := int64(load64(m.memory[int64(uint32(v2))+40:]))
						store64(m.memory[uint32(v9):], uint64(t23))
						v15 = v16
					}
				l6:
					t24 := int64(load64(m.memory[int64(uint32(v2))+100:]))
					t25 := v2
					v14 = t24
					store64(m.memory[int64(uint32(t25))+24:], uint64(v14))
					store32(m.memory[int64(uint32(v2))+32:], uint32(v15))
					{
						if uint32(int32(v14)) <= uint32(v15) {
							goto l7
						}
						m.fn425(v2+i32(8), v2+i32(24), v15, i32(8), i32(24))
						t26 := int32(load32(m.memory[int64(uint32(v2))+8:]))
						v9 = t26
						if v9 != i32(-1) {
							t46 := int32(load32(m.memory[int64(uint32(v2))+12:]))
							m.fn2(v9, t46)
							panic("unreachable")
						}
						t27 := int32(load32(m.memory[int64(uint32(v2))+32:]))
						v15 = t27
					}
				l7:
					v3 = v3 << 5
					v12 = v2 + i32(64) | i32(1)
					t28 := int32(load32(m.memory[uint32(v1):]))
					v17 = t28
					v5 = i32(0)
					t29 := int32(load32(m.memory[int64(uint32(v2))+28:]))
					v16 = t29
					v9 = v4
				l17:
					if v3 != v5 {
						goto l9
					}
					v9 = v6
					goto l10
				l9:
					{
						t30 := int32(m.memory[uint32(v9)])
						v1 = t30
						if v1 == i32(255) {
							goto l11
						}
						t31 := int64(load64(m.memory[int64(uint32(v9))+24:]))
						store64(m.memory[int64(uint32(v12))+23:], uint64(t31))
						t32 := int64(load64(m.memory[int64(uint32(v9))+17:]))
						store64(m.memory[int64(uint32(v12))+16:], uint64(t32))
						t33 := int64(load64(m.memory[int64(uint32(v9))+1:]))
						store64(m.memory[uint32(v12):], uint64(t33))
						t34 := int64(load64(m.memory[int64(uint32(v9))+9:]))
						store64(m.memory[int64(uint32(v12))+8:], uint64(t34))
						m.memory[int64(uint32(v2))+64] = byte(v1)
						t35 := int32(load32(m.memory[int64(uint32(v2))+88:]))
						v14 = int64(uint32(t35-v11)) * v13
						p36 := int32(v14)
						if int32(int64(uint64(v14)>>32)) != 0 {
							p36 = i32(-1)
						}
						t37 := int32(load32(m.memory[int64(uint32(v2))+92:]))
						v1 = p36 + (t37 - v7)
						if uint32(v1) >= uint32(v15) {
							goto l12
						}
						v1 = v16 + v1*i32(24)
						m.fn964(v1)
						t38 := int64(load64(m.memory[int64(uint32(v2))+80:]))
						store64(m.memory[int64(uint32(v1))+16:], uint64(t38))
						t39 := int64(load64(m.memory[int64(uint32(v2))+72:]))
						store64(m.memory[int64(uint32(v1))+8:], uint64(t39))
						t40 := int64(load64(m.memory[int64(uint32(v2))+64:]))
						store64(m.memory[uint32(v1):], uint64(t40))
						goto l13
					}
				l11:
					v9 = v4 + v5 + i32(32)
				l10:
					v12 = int32(uint32(v6-v9) >> 5)
				l15:
					if v12 == 0 {
						m.fn80(v17, v4)
						store32(m.memory[int64(uint32(v0))+24:], uint32(v8))
						store32(m.memory[int64(uint32(v0))+20:], uint32(v10))
						store32(m.memory[int64(uint32(v0))+16:], uint32(v7))
						store32(m.memory[int64(uint32(v0))+12:], uint32(v11))
						t41 := int32(load32(m.memory[int64(uint32(v2))+32:]))
						store32(m.memory[int64(uint32(v0))+8:], uint32(t41))
						t42 := int64(load64(m.memory[int64(uint32(v2))+24:]))
						store64(m.memory[uint32(v0):], uint64(t42))
						goto l16
					}
					v12 = v12 + i32(-1)
					m.fn964(v9)
					v9 = v9 + i32(32)
					goto l15
				l12:
					m.fn964(v2 + i32(64))
				l13:
					v9 = v9 + i32(32)
					v5 = v5 + i32(32)
					goto l17
				}
			}
			t4 := int32(load32(m.memory[int64(uint32(v9))+28:]))
			t5 := v8
			v12 = t4
			p6 := v12
			if uint32(v8) > uint32(v12) {
				p6 = t5
			}
			v8 = p6
			p7 := v12
			if uint32(v7) < uint32(v12) {
				p7 = v7
			}
			v7 = p7
			t8 := int32(load32(m.memory[int64(uint32(v9))+24:]))
			t9 := v10
			v12 = t8
			p10 := v12
			if uint32(v10) > uint32(v12) {
				p10 = t9
			}
			v10 = p10
			p11 := v12
			if uint32(v11) < uint32(v12) {
				p11 = v11
			}
			v11 = p11
			v5 = v5 + i32(-32)
			v9 = v9 + i32(32)
			goto l2
		}
	}
l0:
	store64(m.memory[int64(uint32(v0))+12:], uint64(i64(0)))
	store64(m.memory[uint32(v0):], uint64(i64(0x800000000)))
	store64(m.memory[int64(uint32(v0))+20:], uint64(i64(0)))
	store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
	m.fn1160(v1)
l16:
	m.g0 = v2 + i32(112)
}
func (m *Module) fn1160(v0 int32) {
	var v1, v2, v3 int32
	t0 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	v1 = t0
	t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	v2 = t1
	v3 = v2
l1:
	if v1 == 0 {
		goto l0
	}
	v1 = v1 + i32(-1)
	m.fn964(v3)
	v3 = v3 + i32(32)
	goto l1
l0:
	t2 := int32(load32(m.memory[uint32(v0):]))
	m.fn80(t2, v2)
}
func (m *Module) fn1161(v0, v1 int32) {
	var v2, v3, v4 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		v2 = t0
		t1 := int32(load32(m.memory[uint32(v0):]))
		if v2 != t1 {
			goto l0
		}
		m.fn396(v0)
	}
l0:
	t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	v3 = t2
	if v2 == 0 {
		goto l1
	}
	v4 = v2 << 5
	if v4 == 0 {
		goto l1
	}
	memory_copy(m.memory, uint32(v3+i32(32)), uint32(v3), uint32(v4))
l1:
	t3 := int64(load64(m.memory[int64(uint32(v1))+24:]))
	store64(m.memory[int64(uint32(v3))+24:], uint64(t3))
	t4 := int64(load64(m.memory[int64(uint32(v1))+16:]))
	store64(m.memory[int64(uint32(v3))+16:], uint64(t4))
	t5 := int64(load64(m.memory[int64(uint32(v1))+8:]))
	store64(m.memory[int64(uint32(v3))+8:], uint64(t5))
	t6 := int64(load64(m.memory[uint32(v1):]))
	store64(m.memory[uint32(v3):], uint64(t6))
	store32(m.memory[int64(uint32(v0))+8:], uint32(v2+i32(1)))
}
func (m *Module) fn1162(v0, v1 int32) {
	var v2 int32
	{
		t0 := int32(load32(m.memory[uint32(v0):]))
		t1 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		t2 := v1
		v2 = t1
		if uint32(t2) <= uint32(t0-v2) {
			return
		}
		m.fn62(v0, v2, v1, i32(8), i32(32))
	}
}
func (m *Module) fn1163(v0, v1 int32) {
	var v2, v3, v4 int32
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
	t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	v3 = t2
	if v2 == 0 {
		goto l1
	}
	v4 = v2 * i32(28)
	if v4 == 0 {
		goto l1
	}
	memory_copy(m.memory, uint32(v3+i32(28)), uint32(v3), uint32(v4))
l1:
	t3 := int32(load32(m.memory[int64(uint32(v1))+24:]))
	store32(m.memory[int64(uint32(v3))+24:], uint32(t3))
	t4 := int64(load64(m.memory[int64(uint32(v1))+16:]))
	store64(m.memory[int64(uint32(v3))+16:], uint64(t4))
	t5 := int64(load64(m.memory[int64(uint32(v1))+8:]))
	store64(m.memory[int64(uint32(v3))+8:], uint64(t5))
	t6 := int64(load64(m.memory[uint32(v1):]))
	store64(m.memory[uint32(v3):], uint64(t6))
	store32(m.memory[int64(uint32(v0))+8:], uint32(v2+i32(1)))
}
func (m *Module) fn1164(v0, v1, v2 int32) {
	var v3, v4, v5, v6, v7, v8 int32
	var v9 int64
	var v10, v11, v12, v13, v14 int32
	t0 := m.g0
	v3 = t0 - i32(144)
	m.g0 = v3
	m.fn1165(v3 + i32(8))
	t1 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	v4 = t1
	t2 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	v5 = v4 + t2*i32(12)
	t3 := int32(load32(m.memory[uint32(v1):]))
	v6 = t3
	v1 = v4
	{
	l9:
		{
			if v1 != v5 {
				goto l0
			}
			v7 = v5
			goto l1
		l0:
			v7 = v1 + i32(12)
			t4 := int32(load32(m.memory[uint32(v1):]))
			v8 = t4
			if v8 == i32(-1) {
				goto l1
			}
			t5 := int64(load64(m.memory[int64(uint32(v1))+4:]))
			v9 = t5
			v10 = int32(v9)
			t6 := v10
			v11 = int32(int64(uint64(v9)>>32)) * i32(20)
			v12 = t6 + v11
			m.fn1166(v3 + i32(8))
			v13 = i32(0)
			v1 = v10
			{
			l6:
				if v11 != v13 {
					goto l2
				}
				v1 = v12
				goto l3
			l2:
				{
					t7 := int32(load32(m.memory[uint32(v1):]))
					v14 = t7
					if v14 == i32(-1) {
						goto l4
					}
					t8 := int64(load64(m.memory[int64(uint32(v1))+4:]))
					v9 = t8
					store64(m.memory[int64(uint32(v3))+100:], uint64(i64(0x100000001)))
					store64(m.memory[int64(uint32(v3))+92:], uint64(v9))
					store32(m.memory[int64(uint32(v3))+88:], uint32(v14))
					m.fn1167(v3+i32(64), v3+i32(8), v3+i32(88))
					t9 := int32(load32(m.memory[int64(uint32(v3))+64:]))
					if t9 != i32(-1) {
						goto l5
					}
					v1 = v1 + i32(20)
					v13 = v13 + i32(20)
					goto l6
				}
			l4:
				v1 = v10 + v13 + i32(20)
			l3:
				t10 := int32(uint32(v12-v1) / uint32(i32(20)))
				v13 = t10
			l8:
				if v13 == 0 {
					m.fn426(v8, v10)
					v1 = v7
					goto l9
				}
				v13 = v13 + i32(-1)
				m.fn969(v1)
				v1 = v1 + i32(20)
				goto l8
			}
		l5:
		}
		t11 := int64(load64(m.memory[int64(uint32(v3))+80:]))
		store64(m.memory[int64(uint32(v3))+104:], uint64(t11))
		t12 := int64(load64(m.memory[int64(uint32(v3))+72:]))
		store64(m.memory[int64(uint32(v3))+96:], uint64(t12))
		t13 := int64(load64(m.memory[int64(uint32(v3))+64:]))
		store64(m.memory[int64(uint32(v3))+88:], uint64(t13))
		m.fn97(i32(1079860), i32(54), v3+i32(88), i32(1075064), i32(1079916))
		panic("unreachable")
	}
l1:
	t14 := int32(uint32(v5-v7) / uint32(i32(12)))
	v12 = t14
	v14 = i32(0)
l13:
	{
		if v14 == v12 {
			m.fn136(v6, v4, i32(4), i32(12))
			memory_copy(m.memory, uint32(v3+i32(88)), uint32(v3+i32(8)), uint32(i32(56)))
			m.fn1168(v0, v3+i32(88))
			store32(m.memory[int64(uint32(v0))+12:], uint32(v2))
			m.g0 = v3 + i32(144)
			return
		}
		v11 = v7 + v14*i32(12)
		t15 := int32(load32(m.memory[uint32(v11+i32(8)):]))
		v1 = t15
		v10 = v11 + i32(4)
		t16 := int32(load32(m.memory[uint32(v10):]))
		v13 = t16
	l12:
		if v1 == 0 {
			t17 := int32(load32(m.memory[uint32(v11):]))
			t18 := int32(load32(m.memory[uint32(v10):]))
			m.fn426(t17, t18)
			v14 = v14 + i32(1)
			goto l13
		}
		v1 = v1 + i32(-1)
		m.fn969(v13)
		v13 = v13 + i32(20)
		goto l12
	}
}
func (m *Module) fn1165(v0 int32) {
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
	store32(m.memory[int64(uint32(v0))+48:], uint32(i32(0)))
	store64(m.memory[int64(uint32(v0))+40:], uint64(i64(0x400000000)))
	t3 := int64(load64(m.memory[int64(uint32(i32(0)))+1286256:]))
	store64(m.memory[uint32(v0):], uint64(t3))
	t4 := int64(load64(m.memory[int64(uint32(i32(0)))+1286264:]))
	store64(m.memory[int64(uint32(v0))+8:], uint64(t4))
	store64(m.memory[int64(uint32(v0))+32:], uint64(i64(0)))
	store64(m.memory[int64(uint32(v0))+24:], uint64(v3))
	store64(m.memory[int64(uint32(v0))+16:], uint64(v2))
	m.g0 = v1 + i32(16)
}
func (m *Module) fn1166(v0 int32) {
	var v1 int32
	t0 := m.g0
	v1 = t0 - i32(16)
	m.g0 = v1
	store32(m.memory[int64(uint32(v1))+12:], uint32(i32(0)))
	store64(m.memory[int64(uint32(v1))+4:], uint64(i64(0x400000000)))
	m.fn1169(v0+i32(40), v1+i32(4))
	m.g0 = v1 + i32(16)
}
