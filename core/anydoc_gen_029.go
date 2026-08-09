package core

import (
	"math/bits"
)

func (m *Module) fn1257(v0, v1, v2, v3 int32) int32 {
	t0 := m.fn100(v0, v1, v2, v3)
	return t0
}
func (m *Module) fn1258(v0, v1, v2 int32) int32 {
	var v3, v4 int32
	v3 = i32(1)
	{
		v4 = v2 & i32(255)
		if uint32(v4) > uint32(i32(99)) {
			goto l0
		}
		t0 := int32(uint32(v4) / uint32(i32(10)))
		t1 := v0
		v4 = t0
		t2 := int32(load32(m.memory[int64(uint32(v1))+16:]))
		t3 := v4 | i32(48)
		v1 = t2
		t4 := m.t0[uint(v1)].(func(int32, int32) int32)(t1, t3)
		if t4 != 0 {
			goto l0
		}
		t5 := m.t0[uint(v1)].(func(int32, int32) int32)(v0, (v2-v4*i32(10)|i32(48))&i32(255))
		v3 = t5
	}
l0:
	return v3
}
func (m *Module) fn1259(v0 int32) {
	m.fn972(v0 + i32(40))
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	m.fn1174(t0, t1)
}
func (m *Module) fn1260(v0 int32) int32 {
	var v1, v2, v3 int32
	var v4 int64
	t0 := m.g0
	v1 = t0 - i32(48)
	m.g0 = v1
	t1 := m.fn1175(v0)
	v2 = t1
	t2 := int32(load32(m.memory[int64(uint32(v0))+44:]))
	t3 := int32(load32(m.memory[int64(uint32(v0))+48:]))
	t4 := m.fn857(t2, t3, v2, i32(1077000))
	t5 := int32(load32(m.memory[int64(uint32(t4))+8:]))
	store32(m.memory[int64(uint32(v1))+24:], uint32(t5))
	store32(m.memory[int64(uint32(v1))+20:], uint32(v2))
	m.fn653(v1+i32(8), v0, v1+i32(20))
	{
		{
			t6 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			v3 = t6
			if v3 != i32(1) {
				goto l0
			}
			t7 := int64(load64(m.memory[int64(uint32(v1))+12:]))
			v4 = t7
			t8 := int32(load32(m.memory[int64(uint32(v0))+44:]))
			t9 := int32(load32(m.memory[int64(uint32(v0))+48:]))
			t10 := m.fn857(t8, t9, v2, i32(1077032))
			v0 = t10
			store32(m.memory[int64(uint32(v1))+28:], uint32(i32(-1)))
			store64(m.memory[int64(uint32(v1))+32:], uint64(v4))
			m.fn1172(v0, v1+i32(28))
			goto l1
		}
	l0:
		t11 := int32(load32(m.memory[int64(uint32(v0))+44:]))
		t12 := int32(load32(m.memory[int64(uint32(v0))+48:]))
		t13 := m.fn857(t11, t12, v2, i32(1077016))
		v0 = t13
		store32(m.memory[int64(uint32(v1))+44:], uint32(i32(0)))
		store64(m.memory[int64(uint32(v1))+36:], uint64(i64(0)))
		store64(m.memory[int64(uint32(v1))+28:], uint64(i64(0x800000000)))
		m.fn1172(v0, v1+i32(28))
	}
l1:
	m.g0 = v1 + i32(48)
	return v3
}
func (m *Module) fn1261(v0 int32) {
	m.fn969(v0)
	m.fn1229(v0 + i32(12))
	m.fn1272(v0 + i32(24))
}
func (m *Module) fn1262(v0, v1, v2 int32) int32 {
	var v3, v4 int32
	t0 := m.g0
	v3 = t0 - i32(32)
	m.g0 = v3
	{
		t1 := int32(load32(m.memory[uint32(v0):]))
		v4 = t1
		p2 := i32(3)
		if uint32(v4) > uint32(i32(-0x7ffffff2)) {
			p2 = v4 + i32(0x7ffffff1)
		}
		switch p2 {
		case 5:
			panic("unreachable")
		default:
			store32(m.memory[int64(uint32(v3))+28:], uint32(v0+i32(4)))
			store32(m.memory[int64(uint32(v3))+16:], uint32(i32(24)))
			store32(m.memory[int64(uint32(v3))+12:], uint32(v3+i32(28)))
			t3 := m.fn284(v1, v2, i32(1051739), v3+i32(12))
			v0 = t3
			goto l29
		case 1:
			store32(m.memory[int64(uint32(v3))+28:], uint32(v0+i32(4)))
			store32(m.memory[int64(uint32(v3))+16:], uint32(i32(167)))
			store32(m.memory[int64(uint32(v3))+12:], uint32(v3+i32(28)))
			t4 := m.fn284(v1, v2, i32(0x100bb5), v3+i32(12))
			v0 = t4
			goto l29
		case 2:
			store32(m.memory[int64(uint32(v3))+28:], uint32(v0+i32(4)))
			store32(m.memory[int64(uint32(v3))+16:], uint32(i32(20)))
			store32(m.memory[int64(uint32(v3))+12:], uint32(v3+i32(28)))
			t5 := m.fn284(v1, v2, i32(1051725), v3+i32(12))
			v0 = t5
			goto l29
		case 3:
			store32(m.memory[int64(uint32(v3))+28:], uint32(v0))
			store32(m.memory[int64(uint32(v3))+16:], uint32(i32(168)))
			store32(m.memory[int64(uint32(v3))+12:], uint32(v3+i32(28)))
			t6 := m.fn284(v1, v2, i32(1051635), v3+i32(12))
			v0 = t6
			goto l29
		case 4:
			store32(m.memory[int64(uint32(v3))+28:], uint32(v0+i32(4)))
			store32(m.memory[int64(uint32(v3))+16:], uint32(i32(154)))
			store32(m.memory[int64(uint32(v3))+12:], uint32(v3+i32(28)))
			t7 := m.fn284(v1, v2, i32(1051672), v3+i32(12))
			v0 = t7
			goto l29
		case 6:
			store32(m.memory[int64(uint32(v3))+28:], uint32(v0+i32(4)))
			store32(m.memory[int64(uint32(v3))+16:], uint32(i32(178)))
			store32(m.memory[int64(uint32(v3))+12:], uint32(v3+i32(28)))
			t8 := m.fn284(v1, v2, i32(1051476), v3+i32(12))
			v0 = t8
			goto l29
		case 7:
			store32(m.memory[int64(uint32(v3))+28:], uint32(v0+i32(4)))
			store32(m.memory[int64(uint32(v3))+16:], uint32(i32(177)))
			store32(m.memory[int64(uint32(v3))+12:], uint32(v3+i32(28)))
			t9 := m.fn284(v1, v2, i32(1051549), v3+i32(12))
			v0 = t9
			goto l29
		case 8:
			store32(m.memory[int64(uint32(v3))+28:], uint32(v0+i32(4)))
			store32(m.memory[int64(uint32(v3))+16:], uint32(i32(22)))
			store32(m.memory[int64(uint32(v3))+12:], uint32(v3+i32(28)))
			t10 := m.fn284(v1, v2, i32(1069647), v3+i32(12))
			v0 = t10
			goto l29
		case 9:
			store32(m.memory[int64(uint32(v3))+28:], uint32(v0+i32(4)))
			store32(m.memory[int64(uint32(v3))+16:], uint32(i32(22)))
			store32(m.memory[int64(uint32(v3))+12:], uint32(v3+i32(28)))
			t11 := m.fn284(v1, v2, i32(1067883), v3+i32(12))
			v0 = t11
			goto l29
		case 10:
			store32(m.memory[int64(uint32(v3))+28:], uint32(v0+i32(4)))
			store32(m.memory[int64(uint32(v3))+16:], uint32(i32(36)))
			store32(m.memory[int64(uint32(v3))+12:], uint32(v3+i32(28)))
			t12 := m.fn284(v1, v2, i32(1069941), v3+i32(12))
			v0 = t12
			goto l29
		case 11:
			t13 := int32(load32(m.memory[int64(uint32(v2))+12:]))
			t14 := m.t0[uint(t13)].(func(int32, int32, int32) int32)(v1, i32(1100154), i32(22))
			v0 = t14
			goto l29
		case 12:
			store32(m.memory[int64(uint32(v3))+28:], uint32(v0+i32(4)))
			store32(m.memory[int64(uint32(v3))+16:], uint32(i32(170)))
			store32(m.memory[int64(uint32(v3))+12:], uint32(v3+i32(28)))
			t15 := m.fn284(v1, v2, i32(0x10033d), v3+i32(12))
			v0 = t15
			goto l29
		case 13:
			store32(m.memory[int64(uint32(v3))+28:], uint32(v0+i32(4)))
			store32(m.memory[int64(uint32(v3))+16:], uint32(i32(183)))
			store32(m.memory[int64(uint32(v3))+12:], uint32(v3+i32(28)))
			t16 := m.fn284(v1, v2, i32(1049446), v3+i32(12))
			v0 = t16
			goto l29
		case 14:
			t17 := int32(load32(m.memory[int64(uint32(v2))+12:]))
			t18 := m.t0[uint(t17)].(func(int32, int32, int32) int32)(v1, i32(1100176), i32(47))
			v0 = t18
			goto l29
		case 15:
			t19 := int32(load32(m.memory[int64(uint32(v2))+12:]))
			t20 := m.t0[uint(t19)].(func(int32, int32, int32) int32)(v1, i32(1100223), i32(44))
			v0 = t20
			goto l29
		case 16:
			t21 := int32(load32(m.memory[int64(uint32(v2))+12:]))
			t22 := m.t0[uint(t21)].(func(int32, int32, int32) int32)(v1, i32(1100267), i32(22))
			v0 = t22
			goto l29
		case 17:
			t23 := int32(load32(m.memory[int64(uint32(v2))+12:]))
			t24 := m.t0[uint(t23)].(func(int32, int32, int32) int32)(v1, i32(1100289), i32(19))
			v0 = t24
			goto l29
		case 18:
			store32(m.memory[int64(uint32(v3))+28:], uint32(v0+i32(4)))
			store32(m.memory[int64(uint32(v3))+16:], uint32(i32(141)))
			store32(m.memory[int64(uint32(v3))+12:], uint32(v3+i32(28)))
			t25 := m.fn284(v1, v2, i32(1049524), v3+i32(12))
			v0 = t25
			goto l29
		case 19:
			store32(m.memory[int64(uint32(v3))+28:], uint32(v0+i32(4)))
			store32(m.memory[int64(uint32(v3))+16:], uint32(i32(76)))
			store32(m.memory[int64(uint32(v3))+12:], uint32(v3+i32(28)))
			t26 := m.fn284(v1, v2, i32(1052203), v3+i32(12))
			v0 = t26
			goto l29
		case 20:
			store32(m.memory[int64(uint32(v3))+28:], uint32(v0+i32(4)))
			store32(m.memory[int64(uint32(v3))+16:], uint32(i32(22)))
			store32(m.memory[int64(uint32(v3))+12:], uint32(v3+i32(28)))
			t27 := m.fn284(v1, v2, i32(1052692), v3+i32(12))
			v0 = t27
			goto l29
		case 21:
			store32(m.memory[int64(uint32(v3))+8:], uint32(v0+i32(16)))
			store32(m.memory[int64(uint32(v3))+28:], uint32(v0+i32(4)))
			store32(m.memory[int64(uint32(v3))+24:], uint32(i32(36)))
			store32(m.memory[int64(uint32(v3))+16:], uint32(i32(22)))
			store32(m.memory[int64(uint32(v3))+20:], uint32(v3+i32(28)))
			store32(m.memory[int64(uint32(v3))+12:], uint32(v3+i32(8)))
			t28 := m.fn284(v1, v2, i32(1052669), v3+i32(12))
			v0 = t28
			goto l29
		case 22:
			store32(m.memory[int64(uint32(v3))+28:], uint32(v0+i32(4)))
			store32(m.memory[int64(uint32(v3))+16:], uint32(i32(36)))
			store32(m.memory[int64(uint32(v3))+12:], uint32(v3+i32(28)))
			t29 := m.fn284(v1, v2, i32(1069881), v3+i32(12))
			v0 = t29
			goto l29
		case 23:
			t30 := int32(load32(m.memory[int64(uint32(v2))+12:]))
			t31 := m.t0[uint(t30)].(func(int32, int32, int32) int32)(v1, i32(1099985), i32(30))
			v0 = t31
			goto l29
		case 24:
			store32(m.memory[int64(uint32(v3))+28:], uint32(v0+i32(4)))
			store32(m.memory[int64(uint32(v3))+16:], uint32(i32(36)))
			store32(m.memory[int64(uint32(v3))+12:], uint32(v3+i32(28)))
			t32 := m.fn284(v1, v2, i32(1068112), v3+i32(12))
			v0 = t32
			goto l29
		case 25:
			store32(m.memory[int64(uint32(v3))+28:], uint32(v0+i32(4)))
			store32(m.memory[int64(uint32(v3))+16:], uint32(i32(36)))
			store32(m.memory[int64(uint32(v3))+12:], uint32(v3+i32(28)))
			t33 := m.fn284(v1, v2, i32(1068138), v3+i32(12))
			v0 = t33
			goto l29
		case 26:
			store32(m.memory[int64(uint32(v3))+28:], uint32(v0+i32(4)))
			store32(m.memory[int64(uint32(v3))+16:], uint32(i32(36)))
			store32(m.memory[int64(uint32(v3))+12:], uint32(v3+i32(28)))
			t34 := m.fn284(v1, v2, i32(1049375), v3+i32(12))
			v0 = t34
			goto l29
		case 27:
			store32(m.memory[int64(uint32(v3))+28:], uint32(v0+i32(4)))
			store32(m.memory[int64(uint32(v3))+16:], uint32(i32(172)))
			store32(m.memory[int64(uint32(v3))+12:], uint32(v3+i32(28)))
			t35 := m.fn284(v1, v2, i32(1051649), v3+i32(12))
			v0 = t35
			goto l29
		case 28:
			store32(m.memory[int64(uint32(v3))+28:], uint32(v0+i32(4)))
			store32(m.memory[int64(uint32(v3))+16:], uint32(i32(36)))
			store32(m.memory[int64(uint32(v3))+12:], uint32(v3+i32(28)))
			t36 := m.fn284(v1, v2, i32(1053091), v3+i32(12))
			v0 = t36
		}
	}
l29:
	m.g0 = v3 + i32(32)
	return v0
}
func (m *Module) fn1263(v0, v1, v2, v3 int32) {
	var v4, v5, v6 int32
	t0 := m.g0
	v4 = t0 - i32(48)
	m.g0 = v4
	m.fn1050(v4+i32(24), v1, v2, v3)
	t1 := int32(load32(m.memory[int64(uint32(v4))+32:]))
	v5 = t1
	t2 := int32(load32(m.memory[int64(uint32(v4))+28:]))
	v1 = t2
	{
		{
			{
				t3 := int32(load32(m.memory[int64(uint32(v4))+24:]))
				v6 = t3
				if v6 == i32(-1) {
					goto l0
				}
				t4 := int32(load32(m.memory[int64(uint32(v4))+44:]))
				store32(m.memory[int64(uint32(v4))+16:], uint32(t4))
				t5 := int64(load64(m.memory[int64(uint32(v4))+36:]))
				store64(m.memory[int64(uint32(v4))+8:], uint64(t5))
				goto l1
			}
		l0:
			if v1 != 0 {
				goto l2
			}
			m.fn51(v4+i32(28), v2, v3)
			t6 := int64(load64(m.memory[int64(uint32(v4))+36:]))
			store64(m.memory[int64(uint32(v4))+8:], uint64(t6))
			t7 := int32(load32(m.memory[int64(uint32(v4))+44:]))
			store32(m.memory[int64(uint32(v4))+16:], uint32(t7))
			t8 := int32(load32(m.memory[int64(uint32(v4))+32:]))
			v5 = t8
			t9 := int32(load32(m.memory[int64(uint32(v4))+28:]))
			v1 = t9
			v6 = i32(-0x7ffffffc)
		}
	l1:
		store32(m.memory[int64(uint32(v0))+12:], uint32(v5))
		store32(m.memory[int64(uint32(v0))+8:], uint32(v1))
		store32(m.memory[int64(uint32(v0))+4:], uint32(v6))
		t10 := int64(load64(m.memory[int64(uint32(v4))+8:]))
		store64(m.memory[int64(uint32(v0))+16:], uint64(t10))
		t11 := int32(load32(m.memory[int64(uint32(v4))+16:]))
		store32(m.memory[int64(uint32(v0))+24:], uint32(t11))
		store32(m.memory[uint32(v0):], uint32(i32(-1)))
		goto l3
	}
l2:
	store32(m.memory[int64(uint32(v4))+28:], uint32(v5))
	store32(m.memory[int64(uint32(v4))+24:], uint32(v1))
	m.fn1053(v0, v1+i32(8), v5)
	m.fn754(v4 + i32(24))
l3:
	m.g0 = v4 + i32(48)
}
func (m *Module) fn1264(v0 int32) int32 {
	var v1 int32
l1:
	{
		{
			t0 := m.fn866(v0)
			v1 = t0
			if v1 != 0 {
				goto l0
			}
			return i32(0)
		}
	l0:
		t1 := int32(load32(m.memory[uint32(v1):]))
		if t1 == i32(-1) {
			goto l1
		}
		t2 := int32(load32(m.memory[uint32(v1+i32(4)):]))
		t3 := int32(load32(m.memory[uint32(v1+i32(8)):]))
		t4 := int32(load32(m.memory[int64(uint32(v0))+12:]))
		t5 := int32(load32(m.memory[int64(uint32(v0))+16:]))
		t6 := m.fn773(t2, t3, t4, t5)
		if t6 == 0 {
			goto l1
		}
	}
	return v1
}
func (m *Module) fn1265(v0, v1, v2, v3, v4 int32) {
	m.fn51(v0+i32(12), v1, v2)
	m.fn51(v0, v3, v4)
}
func (m *Module) fn1266(v0, v1, v2 int32) {
	var v3, v4, v5, v6, v7, v8 int32
	var v9, v10 int64
	var v11, v12, v13 int32
	var v14 int64
	var v15 int32
	var v16, v17, v18 int64
	var v19, v20, v21 int32
	t0 := m.g0
	v3 = t0 - i32(256)
	m.g0 = v3
	m.fn140(v3+i32(128), v2)
l2:
	{
		m.fn1055(v3+i32(64), v1, v2, i32(1081161), i32(2))
		t1 := int32(load32(m.memory[int64(uint32(v3))+64:]))
		if t1 != i32(1) {
			goto l0
		}
		t2 := int32(load32(m.memory[int64(uint32(v3))+68:]))
		t3 := v3 + i32(56)
		t4 := v1
		t5 := v2
		v4 = t2
		m.fn825(t3, t4, t5, v4, i32(1081164))
		t6 := int32(load32(m.memory[int64(uint32(v3))+56:]))
		t7 := int32(load32(m.memory[int64(uint32(v3))+60:]))
		m.fn75(v3+i32(128), t6, t7)
		m.fn826(v3+i32(48), v4, v1, v2, i32(1081180))
		t8 := int32(load32(m.memory[int64(uint32(v3))+48:]))
		t9 := int32(load32(m.memory[int64(uint32(v3))+52:]))
		m.fn1055(v3+i32(40), t8, t9, i32(1081196), i32(2))
		t10 := int32(load32(m.memory[int64(uint32(v3))+40:]))
		if t10&i32(1) == 0 {
			goto l1
		}
		t11 := int32(load32(m.memory[int64(uint32(v3))+44:]))
		m.fn826(v3+i32(8), v4+t11+i32(2), v1, v2, i32(1081200))
		t12 := int32(load32(m.memory[int64(uint32(v3))+12:]))
		v2 = t12
		t13 := int32(load32(m.memory[int64(uint32(v3))+8:]))
		v1 = t13
		goto l2
	}
l0:
	m.fn75(v3+i32(128), v1, v2)
l1:
	t14 := int32(load32(m.memory[int64(uint32(v3))+128:]))
	v5 = t14
	t15 := int32(load32(m.memory[int64(uint32(v3))+132:]))
	t16 := v3 + i32(72)
	v6 = t15
	t17 := int32(load32(m.memory[int64(uint32(v3))+136:]))
	m.fn513(t16, v6, t17, i32(125))
	v7 = v3 + i32(232) + i32(12)
	v4 = v3 + i32(112) | i32(5)
	v8 = v3 + i32(112) + i32(9)
l4:
	{
		m.fn515(v3+i32(32), v3+i32(72))
		t18 := int32(load32(m.memory[int64(uint32(v3))+32:]))
		v2 = t18
		if v2 == 0 {
			goto l3
		}
		t19 := int32(load32(m.memory[int64(uint32(v3))+36:]))
		m.fn775(v3+i32(128), v2, t19, i32(123))
		t20 := int32(load32(m.memory[int64(uint32(v3))+128:]))
		v2 = t20
		if v2 == 0 {
			goto l4
		}
		t21 := int32(load32(m.memory[int64(uint32(v3))+132:]))
		v1 = t21
		t22 := int32(load32(m.memory[int64(uint32(v3))+136:]))
		t23 := int32(load32(m.memory[int64(uint32(v3))+140:]))
		m.fn1404(v3+i32(112), t22, t23)
		{
			t24 := m.fn1405(v3 + i32(112))
			if t24 == 0 {
				goto l5
			}
			t25 := m.fn1405(v4)
			if t25 != 0 {
				goto l4
			}
		}
	l5:
		m.fn513(v3+i32(128), v2, v1, i32(44))
		t26 := int64(load32(m.memory[int64(uint32(v3))+117:]))
		t27 := int64(m.memory[uint32(v8)])
		v9 = t26 | t27<<32
		t28 := int64(load32(m.memory[int64(uint32(v3))+112:]))
		t29 := int64(m.memory[int64(uint32(v3))+116])
		v10 = t28 | t29<<32
	l6:
		{
			m.fn515(v3+i32(24), v3+i32(128))
			t30 := int32(load32(m.memory[int64(uint32(v3))+24:]))
			v2 = t30
			if v2 == 0 {
				goto l4
			}
			t31 := int32(load32(m.memory[int64(uint32(v3))+28:]))
			m.fn46(v3+i32(16), v2, t31)
			t32 := int32(load32(m.memory[int64(uint32(v3))+16:]))
			v2 = t32
			t33 := int32(load32(m.memory[int64(uint32(v3))+20:]))
			v1 = t33
			if v1 == 0 {
				goto l6
			}
			t34 := m.fn779(v2, v1, i32(32))
			if t34 != 0 {
				goto l6
			}
			t35 := m.fn779(v2, v1, i32(58))
			if t35 != 0 {
				goto l6
			}
			t36 := m.fn779(v2, v1, i32(91))
			if t36 != 0 {
				goto l6
			}
			m.fn775(v3+i32(192), v2, v1, i32(46))
			{
				{
					{
						t37 := int32(load32(m.memory[int64(uint32(v3))+192:]))
						v11 = t37
						if v11 == 0 {
							m.fn14(v3+i32(232), v2, v1)
							t41 := int64(load64(m.memory[int64(uint32(v3))+236:]))
							v14 = t41
							t42 := int32(load32(m.memory[int64(uint32(v3))+232:]))
							v13 = t42
							v15 = i32(-1)
							goto l10
						}
						t38 := int32(load32(m.memory[int64(uint32(v3))+204:]))
						v2 = t38
						t39 := int32(load32(m.memory[int64(uint32(v3))+200:]))
						v1 = t39
						t40 := int32(load32(m.memory[int64(uint32(v3))+196:]))
						v12 = t40
						if v12 != 0 {
							goto l8
						}
						v13 = i32(-1)
						goto l9
					}
				l8:
					m.fn14(v3+i32(232), v11, v12)
					t43 := int64(load64(m.memory[int64(uint32(v3))+236:]))
					v14 = t43
					t44 := int32(load32(m.memory[int64(uint32(v3))+232:]))
					v13 = t44
				}
			l9:
				m.fn51(v3+i32(232), v1, v2)
				t45 := int64(load64(m.memory[int64(uint32(v3))+236:]))
				v16 = t45
				t46 := int32(load32(m.memory[int64(uint32(v3))+232:]))
				v15 = t46
			}
		l10:
			store64(m.memory[int64(uint32(v3))+172:], uint64(v14))
			store32(m.memory[int64(uint32(v3))+168:], uint32(v13))
			store64(m.memory[int64(uint32(v3))+184:], uint64(v16))
			store32(m.memory[int64(uint32(v3))+180:], uint32(v15))
			store32(m.memory[int64(uint32(v3))+220:], uint32(i32(1000000)))
			t47 := v3
			v17 = v17&i64(-0x10000000000) | v9
			store64(m.memory[int64(uint32(t47))+212:], uint64(v17))
			v2 = i32(0)
			store32(m.memory[int64(uint32(v3))+208:], uint32(i32(0)))
			t48 := v3
			v18 = v18&i64(-0x10000000000) | v10
			store64(m.memory[int64(uint32(t48))+200:], uint64(v18))
			p49 := i32(10)
			if v15 == i32(-1) {
				p49 = i32(0)
			}
			var p50 int32
			if v13 != i32(-1) {
				p50 = 1
			}
			v19 = p49 | p50
			v20 = int32(v14)
			v21 = int32(v16)
		l14:
			{
				if v2 == i32(24) {
					goto l11
				}
				v1 = v3 + i32(192) + v2
				t51 := int32(m.memory[uint32(v1+i32(8))])
				v11 = t51
				if v11 == i32(255) {
					goto l11
				}
				t52 := int32(load32(m.memory[uint32(v1+i32(16)):]))
				v12 = t52
				m.memory[int64(uint32(v3))+224] = byte(v11)
				t53 := int32(load32(m.memory[uint32(v1+i32(9)):]))
				store32(m.memory[int64(uint32(v3))+225:], uint32(t53))
				{
					t54 := m.fn1405(v3 + i32(224))
					if t54 != 0 {
						goto l12
					}
					m.fn225(v3+i32(232), v3+i32(168))
					m.fn225(v7, v3+i32(180))
					v11 = v19 + v12
					{
						t55 := int32(load32(m.memory[int64(uint32(v0))+8:]))
						v1 = t55
						t56 := int32(load32(m.memory[uint32(v0):]))
						if v1 != t56 {
							goto l13
						}
						m.fn1146(v0)
					}
				l13:
					store32(m.memory[int64(uint32(v0))+8:], uint32(v1+i32(1)))
					t57 := int32(load32(m.memory[int64(uint32(v0))+4:]))
					v1 = t57 + v1*i32(36)
					t58 := int64(load64(m.memory[int64(uint32(v3))+232:]))
					store64(m.memory[uint32(v1):], uint64(t58))
					t59 := int64(load64(m.memory[int64(uint32(v3))+240:]))
					store64(m.memory[int64(uint32(v1))+8:], uint64(t59))
					t60 := int64(load64(m.memory[int64(uint32(v3))+248:]))
					store64(m.memory[int64(uint32(v1))+16:], uint64(t60))
					store32(m.memory[int64(uint32(v1))+24:], uint32(v11))
					t61 := int32(load32(m.memory[int64(uint32(v3))+224:]))
					store32(m.memory[int64(uint32(v1))+28:], uint32(t61))
					t62 := int32(m.memory[int64(uint32(v3))+228])
					m.memory[int64(uint32(v1))+32] = byte(t62)
				}
			l12:
				v2 = v2 + i32(12)
				goto l14
			}
		l11:
			m.fn134(v15, v21)
			m.fn134(v13, v20)
			goto l6
		}
	}
l3:
	m.fn16(v5, v6)
	m.g0 = v3 + i32(256)
}
func (m *Module) fn1267(v0, v1, v2, v3 int32) {
	var v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14, v15, v16, v17, v18, v19, v20, v21, v22, v23, v24, v25, v26, v27, v28, v29, v30, v31, v32, v33, v34, v35, v36 int32
	var v37, v38, v39, v40, v41 int64
	var v42, v43, v44, v45, v46, v47 int32
	t0 := m.g0
	v4 = t0 - i32(480)
	m.g0 = v4
	v5 = v1 + i32(12)
	v6 = int32(uint32(v3) >> 24)
	v7 = int32(uint32(v3) >> 16)
	v8 = int32(uint32(v3) >> 8)
	t1 := int32(load32(m.memory[int64(uint32(v2))+28:]))
	v9 = t1
	t2 := int32(load32(m.memory[int64(uint32(v2))+32:]))
	v10 = v9 + t2*i32(44)
	v11 = v4 + i32(336) + i32(8)
	v12 = v4 + i32(336) + i32(20)
	v13 = v4 + i32(248) + i32(12)
	v14 = v4 + i32(248) + i32(8)
	v15 = v4 + i32(336) + i32(12)
	v16 = v4 + i32(336) + i32(4)
	v17 = v4 + i32(336) | i32(4)
	v18 = v4 + i32(280) + i32(4)
	v19 = v4 + i32(248) + i32(16)
	v20 = v4 + i32(336) + i32(16)
	v21 = v4 + i32(280) + i32(5)
l7:
	v2 = v9
	{
		{
			{
			l133:
				{
					{
						{
							{
								{
									if v2 == v10 {
										store32(m.memory[uint32(v0):], uint32(i32(-1)))
										goto l5
									}
									v9 = v2 + i32(44)
									t3 := int32(load32(m.memory[uint32(v2):]))
									if t3 != i32(-1) {
										goto l1
									}
									t4 := int32(load32(m.memory[uint32(v2+i32(8)):]))
									t5 := int32(load32(m.memory[uint32(v2+i32(12)):]))
									m.fn865(v4+i32(336), t4, t5)
									t6 := int32(load32(m.memory[int64(uint32(v4))+340:]))
									t7 := v4 + i32(248)
									v2 = t6
									t8 := int32(load32(m.memory[int64(uint32(v4))+344:]))
									m.fn1406(t7, v2, t8)
									t9 := int32(load32(m.memory[int64(uint32(v4))+336:]))
									m.fn16(t9, v2)
									t10 := int32(load32(m.memory[int64(uint32(v4))+256:]))
									v2 = t10
									if v2 == 0 {
										goto l2
									}
									t11 := int32(load32(m.memory[int64(uint32(v1))+16:]))
									t12 := int32(load32(m.memory[int64(uint32(v1))+20:]))
									t13 := int32(m.memory[int64(uint32(v1))+36])
									t14 := m.fn1407(t11, t12, t13)
									v22 = t14
									if v22 != 0 {
										goto l3
									}
									t15 := int32(load32(m.memory[int64(uint32(v4))+256:]))
									store32(m.memory[int64(uint32(v4))+288:], uint32(t15))
									t16 := int64(load64(m.memory[int64(uint32(v4))+248:]))
									store64(m.memory[int64(uint32(v4))+280:], uint64(t16))
									goto l4
								}
							l3:
								t17 := int32(load32(m.memory[int64(uint32(v4))+252:]))
								m.fn631(v4, t17, v2, i32(32))
								t18 := int32(load32(m.memory[uint32(v4):]))
								t19 := int32(load32(m.memory[int64(uint32(v4))+4:]))
								m.fn51(v4+i32(280), t18, t19)
							}
						l4:
							{
								t20 := int32(load32(m.memory[int64(uint32(v4))+288:]))
								if t20 == 0 {
									goto l6
								}
								t21 := int32(load32(m.memory[int64(uint32(v4))+288:]))
								store32(m.memory[int64(uint32(v16))+8:], uint32(t21))
								t22 := int64(load64(m.memory[int64(uint32(v4))+280:]))
								store64(m.memory[uint32(v16):], uint64(t22))
								store32(m.memory[int64(uint32(v4))+336:], uint32(i32(3)))
								t23 := m.fn1188(v3)
								store32(m.memory[int64(uint32(v4))+352:], uint32(t23))
								m.fn1340(v5, v4+i32(336))
								if v22 == 0 {
									goto l7
								}
								goto l2
							}
						l6:
							t24 := int32(load32(m.memory[int64(uint32(v4))+280:]))
							t25 := int32(load32(m.memory[int64(uint32(v4))+284:]))
							m.fn16(t24, t25)
							if v22 == 0 {
								goto l7
							}
						}
					l2:
						t26 := int32(load32(m.memory[int64(uint32(v4))+248:]))
						t27 := int32(load32(m.memory[int64(uint32(v4))+252:]))
						m.fn16(t26, t27)
						goto l7
					}
				l1:
					t28 := v4 + i32(200)
					v23 = v2 + i32(16)
					t29 := int32(load32(m.memory[uint32(v23):]))
					v24 = v2 + i32(20)
					t30 := int32(load32(m.memory[uint32(v24):]))
					m.fn909(t28, t29, t30, i32(1077139), i32(5))
					v25 = i32(4)
					v26 = i32(0)
					{
						{
							t31 := int32(load32(m.memory[int64(uint32(v4))+200:]))
							v22 = t31
							if v22 == 0 {
								goto l8
							}
							t32 := int32(load32(m.memory[int64(uint32(v4))+204:]))
							v27 = t32
							store16(m.memory[int64(uint32(v4))+308:], uint16(i32(1)))
							store32(m.memory[int64(uint32(v4))+304:], uint32(i32(0)))
							store32(m.memory[int64(uint32(v4))+296:], uint32(v22))
							store32(m.memory[int64(uint32(v4))+288:], uint32(v22))
							store32(m.memory[int64(uint32(v4))+280:], uint32(i32(0)))
							store32(m.memory[int64(uint32(v4))+292:], uint32(v27))
							store32(m.memory[int64(uint32(v4))+284:], uint32(v27))
							store32(m.memory[int64(uint32(v4))+300:], uint32(v22+v27))
							m.fn875(v4+i32(192), v4+i32(280))
							t33 := int32(load32(m.memory[int64(uint32(v4))+192:]))
							v22 = t33
							if v22 == 0 {
								goto l8
							}
							t34 := int32(load32(m.memory[int64(uint32(v4))+196:]))
							v27 = t34
							m.fn59(v4+i32(184), i32(4), i32(4), i32(8))
							t35 := int32(load32(m.memory[int64(uint32(v4))+184:]))
							v28 = t35
							t36 := int32(load32(m.memory[int64(uint32(v4))+188:]))
							v29 = t36
							store32(m.memory[int64(uint32(v29))+4:], uint32(v27))
							store32(m.memory[uint32(v29):], uint32(v22))
							v22 = i32(1)
							store32(m.memory[int64(uint32(v4))+256:], uint32(i32(1)))
							store32(m.memory[int64(uint32(v4))+252:], uint32(v29))
							store32(m.memory[int64(uint32(v4))+248:], uint32(v28))
							t37 := int64(load64(m.memory[int64(uint32(v4))+304:]))
							store64(m.memory[int64(uint32(v4))+360:], uint64(t37))
							t38 := int64(load64(m.memory[int64(uint32(v4))+296:]))
							store64(m.memory[int64(uint32(v4))+352:], uint64(t38))
							t39 := int64(load64(m.memory[int64(uint32(v4))+288:]))
							store64(m.memory[int64(uint32(v4))+344:], uint64(t39))
							t40 := int64(load64(m.memory[int64(uint32(v4))+280:]))
							store64(m.memory[int64(uint32(v4))+336:], uint64(t40))
							v27 = i32(12)
						l11:
							{
								m.fn875(v4+i32(176), v4+i32(336))
								t41 := int32(load32(m.memory[int64(uint32(v4))+176:]))
								v28 = t41
								if v28 == 0 {
									goto l9
								}
								t42 := int32(load32(m.memory[int64(uint32(v4))+180:]))
								v30 = t42
								{
									t43 := int32(load32(m.memory[int64(uint32(v4))+248:]))
									if v22 != t43 {
										goto l10
									}
									m.fn797(v4 + i32(248))
									t44 := int32(load32(m.memory[int64(uint32(v4))+252:]))
									v29 = t44
								}
							l10:
								v31 = v29 + v27
								store32(m.memory[uint32(v31):], uint32(v30))
								store32(m.memory[uint32(v31+i32(-4)):], uint32(v28))
								t45 := v4
								v22 = v22 + i32(1)
								store32(m.memory[int64(uint32(t45))+256:], uint32(v22))
								v27 = v27 + i32(8)
								goto l11
							}
						}
					l8:
						v27 = i32(0)
						goto l12
					l9:
						v27 = i32(0)
						t46 := int32(load32(m.memory[int64(uint32(v4))+248:]))
						v28 = t46
						if v28 == i32(-1) {
							goto l12
						}
						t47 := int32(load32(m.memory[int64(uint32(v4))+252:]))
						v25 = t47
						v26 = v28
						v27 = v22
					}
				l12:
					t48 := int32(load32(m.memory[int64(uint32(v1))+24:]))
					v22 = t48
					t49 := v4
					v32 = v2 + i32(4)
					t50 := int64(load64(m.memory[uint32(v32):]))
					store64(m.memory[int64(uint32(t49))+212:], uint64(t50))
					t51 := int32(load32(m.memory[int64(uint32(v22))+8:]))
					v28 = t51
					t52 := int32(load32(m.memory[int64(uint32(v22))+4:]))
					v22 = t52
					store32(m.memory[int64(uint32(v4))+296:], uint32(v27))
					store32(m.memory[int64(uint32(v4))+292:], uint32(v25))
					store32(m.memory[int64(uint32(v4))+280:], uint32(v22))
					store32(m.memory[int64(uint32(v4))+284:], uint32(v22+v28*i32(36)))
					store32(m.memory[int64(uint32(v4))+288:], uint32(v4+i32(212)))
					m.fn913(v4+i32(436), v4+i32(280))
					{
						t53 := int32(m.memory[int64(uint32(v4))+440])
						if t53 == i32(255) {
							goto l13
						}
						v27 = i32(12)
						m.fn59(v4+i32(168), i32(4), i32(4), i32(12))
						t54 := int32(load32(m.memory[int64(uint32(v4))+168:]))
						v28 = t54
						t55 := int32(load32(m.memory[int64(uint32(v4))+172:]))
						v30 = t55
						t56 := int32(load32(m.memory[int64(uint32(v4))+444:]))
						store32(m.memory[int64(uint32(v30))+8:], uint32(t56))
						t57 := int64(load64(m.memory[int64(uint32(v4))+436:]))
						store64(m.memory[uint32(v30):], uint64(t57))
						v22 = i32(1)
						store32(m.memory[int64(uint32(v4))+472:], uint32(i32(1)))
						store32(m.memory[int64(uint32(v4))+468:], uint32(v30))
						store32(m.memory[int64(uint32(v4))+464:], uint32(v28))
						t58 := int32(load32(m.memory[int64(uint32(v4))+296:]))
						store32(m.memory[int64(uint32(v4))+352:], uint32(t58))
						t59 := int64(load64(m.memory[int64(uint32(v4))+288:]))
						store64(m.memory[int64(uint32(v4))+344:], uint64(t59))
						t60 := int64(load64(m.memory[int64(uint32(v4))+280:]))
						store64(m.memory[int64(uint32(v4))+336:], uint64(t60))
					l16:
						{
							m.fn913(v4+i32(248), v4+i32(336))
							t61 := int32(m.memory[int64(uint32(v4))+252])
							if t61 == i32(255) {
								t67 := int64(load64(m.memory[int64(uint32(v4))+464:]))
								store64(m.memory[int64(uint32(v4))+424:], uint64(t67))
								t68 := int32(load32(m.memory[int64(uint32(v4))+472:]))
								store32(m.memory[int64(uint32(v4))+432:], uint32(t68))
								goto l17
							}
							{
								t62 := int32(load32(m.memory[int64(uint32(v4))+464:]))
								if v22 != t62 {
									goto l15
								}
								m.fn62(v4+i32(464), v22, i32(1), i32(4), i32(12))
								t63 := int32(load32(m.memory[int64(uint32(v4))+468:]))
								v30 = t63
							}
						l15:
							v28 = v30 + v27
							t64 := int32(load32(m.memory[int64(uint32(v4))+256:]))
							store32(m.memory[int64(uint32(v28))+8:], uint32(t64))
							t65 := int64(load64(m.memory[int64(uint32(v4))+248:]))
							store64(m.memory[uint32(v28):], uint64(t65))
							t66 := v4
							v22 = v22 + i32(1)
							store32(m.memory[int64(uint32(t66))+472:], uint32(v22))
							v27 = v27 + i32(12)
							goto l16
						}
					}
				l13:
					store32(m.memory[int64(uint32(v4))+432:], uint32(i32(0)))
					store64(m.memory[int64(uint32(v4))+424:], uint64(i64(0x400000000)))
				l17:
					t69 := int32(load32(m.memory[uint32(v23):]))
					t70 := int32(load32(m.memory[uint32(v24):]))
					m.fn909(v4+i32(160), t69, t70, i32(1077144), i32(5))
					{
						t71 := int32(load32(m.memory[int64(uint32(v4))+160:]))
						v22 = t71
						if v22 == 0 {
							goto l18
						}
						t72 := int32(load32(m.memory[int64(uint32(v4))+164:]))
						m.fn1404(v4+i32(280), v22, t72)
						t73 := int32(m.memory[int64(uint32(v4))+284])
						m.memory[int64(uint32(v16))+4] = byte(t73)
						t74 := int32(load32(m.memory[int64(uint32(v4))+280:]))
						store32(m.memory[uint32(v16):], uint32(t74))
						store32(m.memory[int64(uint32(v4))+336:], uint32(i32(100000)))
						m.fn1408(v4+i32(424), v4+i32(336))
						t75 := int32(m.memory[int64(uint32(v21))+4])
						m.memory[int64(uint32(v16))+4] = byte(t75)
						t76 := int32(load32(m.memory[uint32(v21):]))
						store32(m.memory[uint32(v16):], uint32(t76))
						store32(m.memory[int64(uint32(v4))+336:], uint32(i32(1100000)))
						m.fn1408(v4+i32(424), v4+i32(336))
					}
				l18:
					t77 := int32(load32(m.memory[int64(uint32(v4))+428:]))
					v33 = t77
					{
						t78 := int32(load32(m.memory[int64(uint32(v4))+432:]))
						v22 = t78
						if uint32(v22) < uint32(i32(2)) {
							goto l19
						}
						if uint32(v22) < uint32(i32(21)) {
							v27 = i32(12)
							v28 = v22 * i32(12)
						l21:
							if v28 == v27 {
								goto l19
							}
							m.fn1005(v33, v33+v27)
							v27 = v27 + i32(12)
							goto l21
						}
						m.fn997(v33, v22)
						goto l19
					}
				l19:
					v27 = v22 * i32(12)
					v30 = i32(33686018)
					v28 = i32(2)
					t79 := int32(load32(m.memory[int64(uint32(v4))+424:]))
					v34 = t79
					v22 = v33
				l23:
					{
						if v27 == 0 {
							goto l22
						}
						t80 := int32(m.memory[int64(uint32(v22))+4])
						v31 = t80
						if v31 == i32(255) {
							goto l22
						}
						t81 := int32(load16(m.memory[int64(uint32(v22))+5:]))
						v29 = t81
						v35 = v22 + i32(7)
						t82 := int32(m.memory[int64(uint32(v22))+8])
						t83 := v28
						v36 = t82
						p84 := v36
						if v36 == i32(2) {
							p84 = t83
						}
						v28 = p84
						v27 = v27 + i32(-12)
						v22 = v22 + i32(12)
						t85 := int32(m.memory[uint32(v35)])
						t86 := fn1373(v30, t85<<24|v29<<8|v31)
						v30 = t86
						goto l23
					}
				l22:
					m.fn911(v34, v33)
					m.fn639(v26, v25)
					if v28&i32(255) == i32(1) {
						goto l7
					}
					v28 = i32(1)
					{
						{
							t87 := int32(load32(m.memory[uint32(v32):]))
							v22 = t87
							t88 := v22
							v36 = v2 + i32(8)
							t89 := int32(load32(m.memory[uint32(v36):]))
							v27 = t89
							t90 := m.fn15(t88, v27, i32(1073720), i32(1))
							if t90 == 0 {
								goto l24
							}
							v31 = v6
							v29 = v7
							v35 = v8
							goto l25
						}
					l24:
						v31 = v6
						v29 = v7
						v35 = v8
						t91 := m.fn15(v22, v27, i32(1080987), i32(6))
						if t91 != 0 {
							goto l25
						}
						v35 = i32(1)
						{
							t92 := m.fn15(v22, v27, i32(1073721), i32(1))
							if t92 != 0 {
								goto l26
							}
							t93 := m.fn15(v22, v27, i32(1080993), i32(2))
							if t93 != 0 {
								goto l26
							}
							t94 := m.fn15(v22, v27, i32(1080995), i32(4))
							if t94 != 0 {
								goto l26
							}
							t95 := m.fn15(v22, v27, i32(1080999), i32(3))
							if t95 != 0 {
								goto l26
							}
							v31 = v6
							v29 = v7
							v28 = v3
							t96 := m.fn15(v22, v27, i32(1081002), i32(3))
							if t96 != 0 {
								goto l25
							}
							v29 = i32(1)
							{
								t97 := m.fn15(v22, v27, i32(1079224), i32(1))
								if t97 != 0 {
									goto l27
								}
								t98 := m.fn15(v22, v27, i32(1081005), i32(3))
								if t98 != 0 {
									goto l27
								}
								v31 = v6
								v35 = v8
								v28 = v3
								t99 := m.fn15(v22, v27, i32(1073722), i32(6))
								if t99 != 0 {
									goto l25
								}
								{
									t100 := m.fn15(v22, v27, i32(1285222), i32(4))
									if t100 != 0 {
										goto l28
									}
									t101 := m.fn15(v22, v27, i32(1081008), i32(3))
									if t101 != 0 {
										goto l28
									}
									t102 := m.fn15(v22, v27, i32(1081011), i32(4))
									if t102 != 0 {
										goto l28
									}
									v31 = v6
									v29 = v7
									v35 = v8
									v28 = v3
									t103 := m.fn15(v22, v27, i32(1081015), i32(2))
									if t103 == 0 {
										goto l25
									}
								}
							l28:
								v31 = i32(1)
								v29 = v7
								goto l29
							}
						l27:
							v31 = v6
						l29:
							v35 = v8
							v28 = v3
							goto l25
						}
					l26:
						v31 = v6
						v29 = v7
						v28 = v3
					}
				l25:
					t104 := fn1373(v29<<16&i32(0xff0000)|v31<<24|v35<<8&i32(0xff00)|v28&i32(255), v30)
					v33 = t104
					t105 := m.fn15(v22, v27, i32(1077149), i32(2))
					if t105 != 0 {
						goto l30
					}
					t106 := m.fn15(v22, v27, i32(1077151), i32(2))
					if t106 != 0 {
						goto l30
					}
					t107 := m.fn15(v22, v27, i32(1077153), i32(2))
					if t107 != 0 {
						goto l30
					}
					t108 := m.fn15(v22, v27, i32(1077155), i32(2))
					if t108 != 0 {
						goto l30
					}
					t109 := m.fn15(v22, v27, i32(1077157), i32(2))
					if t109 != 0 {
						goto l30
					}
					t110 := m.fn15(v22, v27, i32(1077159), i32(2))
					if t110 != 0 {
						goto l30
					}
					t111 := m.fn15(v22, v27, i32(1077161), i32(1))
					if t111 != 0 {
						m.fn1409(v1)
						m.fn1417(v1, v2)
						t435 := int32(load32(m.memory[int64(uint32(v5))+8:]))
						v22 = t435
						store32(m.memory[int64(uint32(v1))+20:], uint32(i32(0)))
						t436 := int64(load64(m.memory[uint32(v5):]))
						v37 = t436
						store64(m.memory[int64(uint32(v1))+12:], uint64(i64(0x400000000)))
						store32(m.memory[int64(uint32(v4))+288:], uint32(v22))
						store64(m.memory[int64(uint32(v4))+280:], uint64(v37))
						m.fn1420(v4+i32(336), v1, v2, v33)
						t437 := int32(load32(m.memory[int64(uint32(v4))+348:]))
						v2 = t437
						t438 := int32(load32(m.memory[int64(uint32(v4))+344:]))
						v22 = t438
						t439 := int32(load32(m.memory[int64(uint32(v4))+340:]))
						v27 = t439
						{
							t440 := int32(load32(m.memory[int64(uint32(v4))+336:]))
							v30 = t440
							if v30 == i32(-1) {
								store32(m.memory[int64(uint32(v4))+232:], uint32(v2))
								store32(m.memory[int64(uint32(v4))+228:], uint32(v22))
								store32(m.memory[int64(uint32(v4))+224:], uint32(v27))
								m.fn1341(v4+i32(280), v4+i32(224))
								{
									t442 := int32(load32(m.memory[int64(uint32(v4))+284:]))
									t443 := int32(load32(m.memory[int64(uint32(v4))+288:]))
									t444 := m.fn1421(t442, t443)
									if t444 != 0 {
										t445 := int32(load32(m.memory[int64(uint32(v4))+288:]))
										store32(m.memory[int64(uint32(v17))+8:], uint32(t445))
										t446 := int64(load64(m.memory[int64(uint32(v4))+280:]))
										store64(m.memory[uint32(v17):], uint64(t446))
										store32(m.memory[int64(uint32(v4))+336:], uint32(i32(-0x80000000)))
										m.fn338(v1, v4+i32(336))
										goto l7
									}
									m.fn894(v4 + i32(280))
									goto l7
								}
							}
							t441 := int64(load64(m.memory[int64(uint32(v4))+352:]))
							v41 = t441
							m.fn894(v4 + i32(280))
							v42 = v2
							v43 = v22
							v44 = v27
							goto l89
						}
					}
					{
						{
							t112 := m.fn15(v22, v27, i32(1077162), i32(2))
							if t112 != 0 {
								goto l32
							}
							t113 := m.fn15(v22, v27, i32(1077048), i32(2))
							if t113 == 0 {
								goto l33
							}
						}
					l32:
						m.fn1409(v1)
						t114 := int32(load32(m.memory[uint32(v32):]))
						t115 := int32(load32(m.memory[uint32(v36):]))
						t116 := m.fn773(t114, t115, i32(1077048), i32(2))
						v30 = t116
						t117 := int32(load32(m.memory[int64(uint32(v2))+32:]))
						v22 = t117
						t118 := int32(load32(m.memory[int64(uint32(v2))+28:]))
						t119 := v4
						v2 = t118
						store32(m.memory[int64(uint32(t119))+248:], uint32(v2))
						store32(m.memory[int64(uint32(v4))+252:], uint32(v2+v22*i32(44)))
						{
							t120 := m.fn906(v4 + i32(248))
							v22 = t120
							if v22 == 0 {
								v30 = i32(-1)
								v2 = i32(8)
								v31 = i32(4)
								v29 = i32(0)
								v26 = i32(0)
								v35 = i32(0)
								goto l45
							}
							t121 := int64(load64(m.memory[int64(uint32(v4))+248:]))
							v37 = t121
							v2 = i32(4)
							m.fn59(v4+i32(80), i32(4), i32(4), i32(4))
							t122 := int32(load32(m.memory[int64(uint32(v4))+80:]))
							v27 = t122
							t123 := int32(load32(m.memory[int64(uint32(v4))+84:]))
							v28 = t123
							store32(m.memory[uint32(v28):], uint32(v22))
							v22 = i32(1)
							store32(m.memory[int64(uint32(v4))+344:], uint32(i32(1)))
							store32(m.memory[int64(uint32(v4))+340:], uint32(v28))
							store32(m.memory[int64(uint32(v4))+336:], uint32(v27))
							store64(m.memory[int64(uint32(v4))+280:], uint64(v37))
						l37:
							{
								t124 := m.fn906(v4 + i32(280))
								v27 = t124
								if v27 == 0 {
									t128 := int32(load32(m.memory[int64(uint32(v4))+336:]))
									v26 = t128
									t129 := int32(load32(m.memory[int64(uint32(v4))+340:]))
									v31 = t129
									{
										if v30 != 0 {
											t130 := int32(load32(m.memory[uint32(v23):]))
											t131 := v4 + i32(72)
											v27 = t130
											t132 := int32(load32(m.memory[uint32(v24):]))
											t133 := v27
											v28 = t132
											m.fn909(t131, t133, v28, i32(1074404), i32(4))
											t134 := int32(load32(m.memory[int64(uint32(v4))+72:]))
											v30 = t134
											if v30 != 0 {
												goto l40
											}
											v25 = i32(1)
											goto l41
										}
										m.fn1410(v4+i32(392), v22)
										v22 = i32(0)
										{
										l44:
											{
												if v2 == v22 {
													t432 := m.fn113(i32(8), i32(32))
													v2 = t432
													store64(m.memory[int64(uint32(v2))+8:], uint64(i64(1)))
													store32(m.memory[uint32(v2):], uint32(i32(-0x7fffffff)))
													m.memory[int64(uint32(v2))+28] = byte(i32(0))
													t433 := int64(load64(m.memory[int64(uint32(v4))+392:]))
													store64(m.memory[int64(uint32(v2))+16:], uint64(t433))
													t434 := int32(load32(m.memory[int64(uint32(v4))+400:]))
													store32(m.memory[int64(uint32(v2))+24:], uint32(t434))
													v30 = i32(-1)
													v29 = i32(1)
													v35 = i32(1)
													goto l45
												}
												t135 := int32(load32(m.memory[uint32(v31+v22):]))
												m.fn1411(v4+i32(280), v1, t135, v33)
												{
													t136 := int32(load32(m.memory[int64(uint32(v4))+280:]))
													v30 = t136
													if v30 != i32(-1) {
														goto l43
													}
													m.memory[int64(uint32(v4))+360] = byte(i32(2))
													t137 := int32(load32(m.memory[int64(uint32(v4))+292:]))
													store32(m.memory[int64(uint32(v4))+344:], uint32(t137))
													t138 := int64(load64(m.memory[int64(uint32(v4))+284:]))
													store64(m.memory[int64(uint32(v4))+336:], uint64(t138))
													store32(m.memory[int64(uint32(v4))+348:], uint32(i32(-1)))
													v22 = v22 + i32(4)
													m.fn1412(v4+i32(392), v4+i32(336))
													goto l44
												}
											l43:
											}
											t139 := int64(load64(m.memory[int64(uint32(v4))+296:]))
											v38 = t139
											t140 := int32(load32(m.memory[int64(uint32(v4))+292:]))
											v29 = t140
											t141 := int32(load32(m.memory[int64(uint32(v4))+288:]))
											v2 = t141
											t142 := int32(load32(m.memory[int64(uint32(v4))+284:]))
											v35 = t142
											m.fn971(v4 + i32(392))
											goto l45
										}
									l40:
										{
											t143 := int32(load32(m.memory[int64(uint32(v4))+76:]))
											t144 := v30
											v29 = t143
											t145 := m.fn15(t144, v29, i32(1077050), i32(1))
											if t145 == 0 {
												goto l46
											}
											v25 = i32(2)
											goto l41
										}
									l46:
										{
											t146 := m.fn15(v30, v29, i32(1077051), i32(1))
											if t146 == 0 {
												goto l47
											}
											v25 = i32(3)
											goto l41
										}
									l47:
										{
											t147 := m.fn15(v30, v29, i32(1073721), i32(1))
											if t147 == 0 {
												goto l48
											}
											v25 = i32(4)
											goto l41
										}
									l48:
										t148 := m.fn15(v30, v29, i32(1077052), i32(1))
										p149 := i32(1)
										if t148 != 0 {
											p149 = i32(5)
										}
										v25 = p149
									}
								l41:
									m.fn909(v4+i32(64), v27, v28, i32(1077053), i32(8))
									t150 := int32(load32(m.memory[int64(uint32(v4))+64:]))
									v30 = t150
									m.fn909(v4+i32(56), v27, v28, i32(1077061), i32(5))
									{
										{
											t151 := int32(load32(m.memory[int64(uint32(v4))+56:]))
											v27 = t151
											if v27 != 0 {
												goto l49
											}
											p152 := i64(1)
											if v30 != 0 {
												p152 = int64(uint32(v22))
											}
											v37 = p152
											goto l50
										}
									l49:
										t153 := int32(load32(m.memory[int64(uint32(v4))+60:]))
										m.fn1322(v4+i32(336), v27, t153)
										p154 := i64(1)
										if v30 != 0 {
											p154 = int64(uint32(v22))
										}
										t155 := int64(load64(m.memory[int64(uint32(v4))+344:]))
										t156 := int32(m.memory[int64(uint32(v4))+336])
										p157 := t155
										if t156 != 0 {
											p157 = p154
										}
										v37 = p157
									}
								l50:
									m.fn59(v4+i32(48), v22, i32(8), i32(8))
									v27 = i32(0)
									store32(m.memory[int64(uint32(v4))+472:], uint32(i32(0)))
									t158 := int64(load64(m.memory[int64(uint32(v4))+48:]))
									store64(m.memory[int64(uint32(v4))+464:], uint64(t158))
								l69:
									{
										{
											if v2 == v27 {
												t169 := int32(load32(m.memory[int64(uint32(v4))+472:]))
												v30 = t169
												v28 = v30 << 3
												v2 = i32(0)
												t170 := int32(load32(m.memory[int64(uint32(v4))+468:]))
												v36 = t170
												{
												l56:
													{
														if v28 == v2 {
															store32(m.memory[int64(uint32(v4))+444:], uint32(i32(0)))
															store64(m.memory[int64(uint32(v4))+436:], uint64(i64(0x800000000)))
															store32(m.memory[int64(uint32(v4))+256:], uint32(i32(-1)))
															p173 := v22
															if uint32(v30) < uint32(v22) {
																p173 = v30
															}
															v22 = p173
															v39 = i64(0)
															v27 = v36
															v28 = v31
														l65:
															{
																if v22 == 0 {
																	{
																		t181 := int32(load32(m.memory[int64(uint32(v4))+256:]))
																		if t181 == i32(-1) {
																			goto l61
																		}
																		t182 := int64(load64(m.memory[int64(uint32(v4))+264:]))
																		store64(m.memory[int64(uint32(v11))+16:], uint64(t182))
																		t183 := int64(load64(m.memory[int64(uint32(v4))+256:]))
																		store64(m.memory[int64(uint32(v11))+8:], uint64(t183))
																		t184 := int64(load64(m.memory[int64(uint32(v4))+248:]))
																		store64(m.memory[uint32(v11):], uint64(t184))
																		store32(m.memory[int64(uint32(v4))+336:], uint32(i32(-0x7fffffff)))
																		m.fn338(v4+i32(436), v4+i32(336))
																	}
																l61:
																	t185 := int32(load32(m.memory[int64(uint32(v4))+444:]))
																	v29 = t185
																	t186 := int32(load32(m.memory[int64(uint32(v4))+440:]))
																	v2 = t186
																	t187 := int32(load32(m.memory[int64(uint32(v4))+436:]))
																	v35 = t187
																	t188 := int32(load32(m.memory[int64(uint32(v4))+464:]))
																	m.fn1415(t188, v36)
																	v30 = i32(-1)
																	goto l45
																}
																t174 := int64(load64(m.memory[uint32(v27):]))
																v37 = t174
																t175 := int32(load32(m.memory[uint32(v28):]))
																m.fn1411(v4+i32(336), v1, t175, v33)
																t176 := int32(load32(m.memory[int64(uint32(v4))+348:]))
																v29 = t176
																t177 := int32(load32(m.memory[int64(uint32(v4))+344:]))
																v2 = t177
																t178 := int32(load32(m.memory[int64(uint32(v4))+340:]))
																v35 = t178
																t179 := int32(load32(m.memory[int64(uint32(v4))+336:]))
																v30 = t179
																if v30 == i32(-1) {
																	m.memory[int64(uint32(v4))+304] = byte(i32(2))
																	store32(m.memory[int64(uint32(v4))+288:], uint32(v29))
																	store32(m.memory[int64(uint32(v4))+284:], uint32(v2))
																	store32(m.memory[int64(uint32(v4))+280:], uint32(v35))
																	store32(m.memory[int64(uint32(v4))+292:], uint32(i32(-1)))
																	{
																		t189 := int32(load32(m.memory[int64(uint32(v4))+256:]))
																		v2 = t189
																		if v2 == i32(-1) {
																			goto l62
																		}
																		v40 = v39 + i64(1)
																		if v40 < v39 {
																			goto l63
																		}
																		if v40 == v37 {
																			goto l64
																		}
																	l63:
																		store32(m.memory[int64(uint32(v4))+256:], uint32(i32(-1)))
																		t190 := int64(load64(m.memory[int64(uint32(v4))+248:]))
																		v39 = t190
																		t191 := int32(load32(m.memory[int64(uint32(v13))+8:]))
																		store32(m.memory[int64(uint32(v12))+8:], uint32(t191))
																		t192 := int64(load64(m.memory[uint32(v13):]))
																		store64(m.memory[uint32(v12):], uint64(t192))
																		store32(m.memory[int64(uint32(v4))+352:], uint32(v2))
																		store64(m.memory[int64(uint32(v4))+344:], uint64(v39))
																		store32(m.memory[int64(uint32(v4))+336:], uint32(i32(-0x7fffffff)))
																		m.fn338(v4+i32(436), v4+i32(336))
																	}
																l62:
																	m.fn1414(v4 + i32(248))
																	m.memory[int64(uint32(v4))+268] = byte(v25)
																	store32(m.memory[int64(uint32(v4))+264:], uint32(i32(0)))
																	store64(m.memory[int64(uint32(v4))+256:], uint64(i64(0x400000000)))
																	store64(m.memory[int64(uint32(v4))+248:], uint64(v37))
																l64:
																	m.fn1412(v14, v4+i32(280))
																	v22 = v22 + i32(-1)
																	v27 = v27 + i32(8)
																	v28 = v28 + i32(4)
																	v39 = v37
																	goto l65
																}
																t180 := int64(load64(m.memory[int64(uint32(v4))+352:]))
																v38 = t180
																m.fn1414(v4 + i32(248))
																m.fn969(v4 + i32(436))
																goto l60
															}
														}
														v27 = v36 + v2
														v2 = v2 + i32(8)
														t171 := int64(load64(m.memory[uint32(v27):]))
														if t171 >= i64(1) {
															goto l56
														}
													}
													m.fn1410(v4+i32(408), v22)
													p172 := v22
													if uint32(v30) < uint32(v22) {
														p172 = v30
													}
													v2 = p172
													v22 = v36
													v27 = v31
													{
													l68:
														{
															if v2 == 0 {
																t429 := m.fn113(i32(8), i32(32))
																v2 = t429
																store64(m.memory[int64(uint32(v2))+8:], uint64(i64(1)))
																store32(m.memory[uint32(v2):], uint32(i32(-0x7fffffff)))
																m.memory[int64(uint32(v2))+28] = byte(v25)
																t430 := int64(load64(m.memory[int64(uint32(v4))+408:]))
																store64(m.memory[int64(uint32(v2))+16:], uint64(t430))
																t431 := int32(load32(m.memory[int64(uint32(v4))+416:]))
																store32(m.memory[int64(uint32(v2))+24:], uint32(t431))
																v30 = i32(-1)
																v29 = i32(1)
																v35 = i32(1)
																goto l60
															}
															t193 := int64(load64(m.memory[uint32(v22):]))
															store64(m.memory[int64(uint32(v4))+248:], uint64(t193))
															t194 := int32(load32(m.memory[uint32(v27):]))
															m.fn1411(v4+i32(280), v1, t194, v33)
															{
																t195 := int32(load32(m.memory[int64(uint32(v4))+280:]))
																v30 = t195
																if v30 != i32(-1) {
																	goto l67
																}
																t196 := int32(load32(m.memory[int64(uint32(v4))+292:]))
																store32(m.memory[int64(uint32(v4))+344:], uint32(t196))
																t197 := int64(load64(m.memory[int64(uint32(v4))+284:]))
																store64(m.memory[int64(uint32(v4))+336:], uint64(t197))
																store32(m.memory[int64(uint32(v4))+284:], uint32(i32(144)))
																store32(m.memory[int64(uint32(v4))+280:], uint32(v4+i32(248)))
																m.fn73(v15, i32(1068665), v4+i32(280))
																m.memory[int64(uint32(v4))+360] = byte(i32(2))
																v2 = v2 + i32(-1)
																v22 = v22 + i32(8)
																v27 = v27 + i32(4)
																m.fn1412(v4+i32(408), v4+i32(336))
																goto l68
															}
														l67:
														}
														t198 := int64(load64(m.memory[int64(uint32(v4))+296:]))
														v38 = t198
														t199 := int32(load32(m.memory[int64(uint32(v4))+292:]))
														v29 = t199
														t200 := int32(load32(m.memory[int64(uint32(v4))+288:]))
														v2 = t200
														t201 := int32(load32(m.memory[int64(uint32(v4))+284:]))
														v35 = t201
														m.fn971(v4 + i32(408))
														goto l60
													}
												}
											}
											t159 := int32(load32(m.memory[uint32(v31+v27):]))
											t160 := v4 + i32(40)
											v28 = t159
											t161 := int32(load32(m.memory[uint32(v28+i32(16)):]))
											t162 := int32(load32(m.memory[uint32(v28+i32(20)):]))
											m.fn909(t160, t161, t162, i32(1077066), i32(5))
											{
												t163 := int32(load32(m.memory[int64(uint32(v4))+40:]))
												v28 = t163
												if v28 == 0 {
													goto l52
												}
												t164 := int32(load32(m.memory[int64(uint32(v4))+44:]))
												m.fn1322(v4+i32(336), v28, t164)
												t165 := int64(load64(m.memory[int64(uint32(v4))+344:]))
												t166 := int32(m.memory[int64(uint32(v4))+336])
												p167 := t165
												if t166 != 0 {
													p167 = v37
												}
												v37 = p167
											}
										l52:
											m.fn1413(v4+i32(464), v37)
											if v30 != 0 {
												goto l53
											}
											v39 = v37 + i64(1)
											p168 := v39
											if v39 < v37 {
												p168 = i64(0x7fffffffffffffff)
											}
											v37 = p168
											goto l54
										}
									l53:
										v39 = v37 + i64(-1)
										p202 := v39
										if v39 >= v37 {
											p202 = i64(-0x8000000000000000)
										}
										v37 = p202
									}
								l54:
									v27 = v27 + i32(4)
									goto l69
								}
								{
									t125 := int32(load32(m.memory[int64(uint32(v4))+336:]))
									if v22 != t125 {
										goto l36
									}
									m.fn905(v4 + i32(336))
									t126 := int32(load32(m.memory[int64(uint32(v4))+340:]))
									v28 = t126
								}
							l36:
								store32(m.memory[uint32(v28+v2):], uint32(v27))
								t127 := v4
								v22 = v22 + i32(1)
								store32(m.memory[int64(uint32(t127))+344:], uint32(v22))
								v2 = v2 + i32(4)
								goto l37
							}
						}
					}
				l33:
					{
						{
							{
								{
									{
										t203 := m.fn15(v22, v27, i32(1074842), i32(5))
										if t203 != 0 {
											m.fn1409(v1)
											t232 := int32(load32(m.memory[int64(uint32(v2))+32:]))
											v31 = t232
											v30 = v31 * i32(44)
											t233 := int32(load32(m.memory[int64(uint32(v2))+28:]))
											v28 = t233
											v22 = i32(0)
										l83:
											if v30 == v22 {
												goto l80
											}
											{
												v27 = v28 + v22
												t234 := int32(load32(m.memory[uint32(v27):]))
												if t234 == i32(-1) {
													goto l81
												}
												t235 := int32(load32(m.memory[uint32(v27+i32(4)):]))
												t236 := int32(load32(m.memory[uint32(v27+i32(8)):]))
												t237 := m.fn773(t235, t236, i32(1073494), i32(7))
												if t237 != 0 {
													goto l82
												}
											}
										l81:
											v22 = v22 + i32(44)
											goto l83
										}
										t204 := m.fn15(v22, v27, i32(1077164), i32(10))
										if t204 != 0 {
											m.fn1409(v1)
											m.fn1411(v4+i32(336), v1, v2, v33)
											t314 := int32(load32(m.memory[int64(uint32(v4))+348:]))
											v2 = t314
											t315 := int32(load32(m.memory[int64(uint32(v4))+344:]))
											v22 = t315
											t316 := int32(load32(m.memory[int64(uint32(v4))+340:]))
											v27 = t316
											{
												t317 := int32(load32(m.memory[int64(uint32(v4))+336:]))
												v30 = t317
												if v30 == i32(-1) {
													store32(m.memory[int64(uint32(v4))+288:], uint32(v2))
													store32(m.memory[int64(uint32(v4))+284:], uint32(v22))
													store32(m.memory[int64(uint32(v4))+280:], uint32(v27))
													{
														if v2 == 0 {
															m.fn969(v4 + i32(280))
															goto l7
														}
														t319 := int32(load32(m.memory[int64(uint32(v4))+288:]))
														store32(m.memory[int64(uint32(v17))+8:], uint32(t319))
														t320 := int64(load64(m.memory[int64(uint32(v4))+280:]))
														store64(m.memory[uint32(v17):], uint64(t320))
														store32(m.memory[int64(uint32(v4))+336:], uint32(i32(-0x7ffffffd)))
														m.fn338(v1, v4+i32(336))
														goto l7
													}
												}
												t318 := int64(load64(m.memory[int64(uint32(v4))+352:]))
												v41 = t318
												v42 = v2
												v43 = v22
												v44 = v27
												goto l89
											}
										}
										t205 := m.fn15(v22, v27, i32(1077174), i32(3))
										if t205 != 0 {
											m.fn1409(v1)
											t305 := int32(load32(m.memory[uint32(v2+i32(28)):]))
											t306 := int32(load32(m.memory[uint32(v2+i32(32)):]))
											m.fn864(v4+i32(280), t305, t306)
											t307 := int32(load32(m.memory[int64(uint32(v4))+284:]))
											t308 := v4 + i32(112)
											v2 = t307
											t309 := int32(load32(m.memory[int64(uint32(v4))+288:]))
											m.fn46(t308, v2, t309)
											{
												t310 := int32(load32(m.memory[int64(uint32(v4))+116:]))
												if t310 == 0 {
													t313 := int32(load32(m.memory[int64(uint32(v4))+280:]))
													m.fn16(t313, v2)
													goto l7
												}
												t311 := int32(load32(m.memory[int64(uint32(v4))+288:]))
												store32(m.memory[int64(uint32(v17))+8:], uint32(t311))
												t312 := int64(load64(m.memory[int64(uint32(v4))+280:]))
												store64(m.memory[uint32(v17):], uint64(t312))
												store32(m.memory[int64(uint32(v4))+352:], uint32(i32(-1)))
												store32(m.memory[int64(uint32(v4))+336:], uint32(i32(-0x7ffffffc)))
												m.fn338(v1, v4+i32(336))
												goto l7
											}
										}
										t206 := m.fn15(v22, v27, i32(1077177), i32(2))
										if t206 != 0 {
											m.fn1409(v1)
											store32(m.memory[int64(uint32(v4))+336:], uint32(i32(-0x7ffffffb)))
											m.fn338(v1, v4+i32(336))
											goto l7
										}
										t207 := m.fn1416(v22, v27)
										if t207 != 0 {
											m.fn1417(v1, v2)
											t238 := int32(load32(m.memory[int64(uint32(v2))+32:]))
											v30 = t238 * i32(44)
											t239 := int32(load32(m.memory[int64(uint32(v2))+28:]))
											v28 = t239
										l87:
											if v30 == 0 {
												m.fn1267(v4+i32(336), v1, v2, v33)
												t303 := int32(load32(m.memory[int64(uint32(v4))+336:]))
												v30 = t303
												if v30 == i32(-1) {
													goto l7
												}
												goto l91
											}
											{
												t240 := int32(load32(m.memory[uint32(v28):]))
												if t240 == i32(-1) {
													goto l85
												}
												t241 := int32(load32(m.memory[uint32(v28+i32(4)):]))
												v22 = t241
												t242 := int32(load32(m.memory[uint32(v28+i32(8)):]))
												t243 := v22
												v27 = t242
												t244 := m.fn1416(t243, v27)
												if t244 != 0 {
													goto l86
												}
												t245 := m.fn15(v22, v27, i32(1077161), i32(1))
												if t245 != 0 {
													goto l86
												}
												t246 := m.fn15(v22, v27, i32(1077162), i32(2))
												if t246 != 0 {
													goto l86
												}
												t247 := m.fn15(v22, v27, i32(1077048), i32(2))
												if t247 != 0 {
													goto l86
												}
												t248 := m.fn15(v22, v27, i32(1074842), i32(5))
												if t248 != 0 {
													goto l86
												}
												t249 := m.fn15(v22, v27, i32(1077164), i32(10))
												if t249 != 0 {
													goto l86
												}
												t250 := m.fn15(v22, v27, i32(1077174), i32(3))
												if t250 != 0 {
													goto l86
												}
												t251 := m.fn15(v22, v27, i32(1077177), i32(2))
												if t251 != 0 {
													goto l86
												}
												t252 := m.fn15(v22, v27, i32(1077149), i32(2))
												if t252 != 0 {
													goto l86
												}
												t253 := m.fn15(v22, v27, i32(1077151), i32(2))
												if t253 != 0 {
													goto l86
												}
												t254 := m.fn15(v22, v27, i32(1077153), i32(2))
												if t254 != 0 {
													goto l86
												}
												t255 := m.fn15(v22, v27, i32(1077155), i32(2))
												if t255 != 0 {
													goto l86
												}
												t256 := m.fn15(v22, v27, i32(1077157), i32(2))
												if t256 != 0 {
													goto l86
												}
												t257 := m.fn15(v22, v27, i32(1077159), i32(2))
												if t257 != 0 {
													goto l86
												}
											}
										l85:
											v28 = v28 + i32(44)
											v30 = v30 + i32(-44)
											goto l87
										}
										t208 := m.fn15(v22, v27, i32(1077179), i32(6))
										if t208 != 0 {
											goto l7
										}
										t209 := m.fn15(v22, v27, i32(1077144), i32(5))
										if t209 != 0 {
											goto l7
										}
										t210 := m.fn15(v22, v27, i32(1077185), i32(4))
										if t210 != 0 {
											goto l7
										}
										t211 := m.fn15(v22, v27, i32(1077189), i32(8))
										if t211 != 0 {
											goto l7
										}
										t212 := m.fn15(v22, v27, i32(1077197), i32(8))
										if t212 != 0 {
											goto l7
										}
										m.fn1417(v1, v2)
										t213 := int32(load32(m.memory[uint32(v32):]))
										v22 = t213
										t214 := int32(load32(m.memory[uint32(v36):]))
										t215 := v22
										v27 = t214
										t216 := m.fn15(t215, v27, i32(1077123), i32(2))
										if t216 != 0 {
											goto l75
										}
										{
											t217 := m.fn15(v22, v27, i32(1077125), i32(3))
											if t217 != 0 {
												goto l76
											}
											t218 := m.fn15(v22, v27, i32(1077128), i32(5))
											if t218 == 0 {
												t273 := m.fn15(v22, v27, i32(1077050), i32(1))
												if t273 != 0 {
													t275 := int32(load32(m.memory[uint32(v23):]))
													t276 := int32(load32(m.memory[uint32(v24):]))
													m.fn909(v4+i32(152), t275, t276, i32(1073490), i32(4))
													{
														t277 := int32(load32(m.memory[int64(uint32(v4))+152:]))
														v22 = t277
														if v22 == 0 {
															goto l92
														}
														t278 := int32(load32(m.memory[int64(uint32(v1))+28:]))
														t279 := int32(load32(m.memory[int64(uint32(v4))+156:]))
														t280 := int32(load32(m.memory[int64(uint32(v1))+32:]))
														t281 := int32(load32(m.memory[int64(uint32(t280))+12:]))
														m.t0[uint(t281)].(func(int32, int32, int32, int32))(v4+i32(280), t278, v22, t279)
														goto l93
													}
												l92:
													store32(m.memory[int64(uint32(v4))+280:], uint32(i32(-1)))
												l93:
													t282 := int32(load32(m.memory[int64(uint32(v1))+16:]))
													t283 := int32(load32(m.memory[int64(uint32(v1))+20:]))
													t284 := int32(m.memory[int64(uint32(v1))+36])
													t285 := m.fn1407(t282, t283, t284)
													m.fn1418(v4+i32(336), v1, v2, v33, t285)
													t286 := int32(load32(m.memory[int64(uint32(v4))+348:]))
													v2 = t286
													t287 := int32(load32(m.memory[int64(uint32(v4))+344:]))
													v22 = t287
													t288 := int32(load32(m.memory[int64(uint32(v4))+340:]))
													v27 = t288
													{
														t289 := int32(load32(m.memory[int64(uint32(v4))+336:]))
														v30 = t289
														if v30 == i32(-1) {
															store32(m.memory[int64(uint32(v4))+256:], uint32(v2))
															store32(m.memory[int64(uint32(v4))+252:], uint32(v22))
															store32(m.memory[int64(uint32(v4))+248:], uint32(v27))
															{
																t291 := int32(load32(m.memory[int64(uint32(v4))+280:]))
																if t291 == i32(-1) {
																	m.fn1341(v5, v4+i32(248))
																	goto l7
																}
																t292 := int64(load64(m.memory[int64(uint32(v4))+248:]))
																store64(m.memory[uint32(v20):], uint64(t292))
																t293 := int32(load32(m.memory[int64(uint32(v4))+256:]))
																store32(m.memory[int64(uint32(v20))+8:], uint32(t293))
																t294 := int64(load64(m.memory[int64(uint32(v4))+288:]))
																store64(m.memory[int64(uint32(v4))+344:], uint64(t294))
																t295 := int64(load64(m.memory[int64(uint32(v4))+280:]))
																store64(m.memory[int64(uint32(v4))+336:], uint64(t295))
																m.fn1340(v5, v4+i32(336))
																goto l7
															}
														}
														t290 := int64(load64(m.memory[int64(uint32(v4))+352:]))
														v41 = t290
														m.fn1419(v4 + i32(280))
														v42 = v2
														v43 = v22
														v44 = v27
														goto l89
													}
												}
												m.fn1267(v4+i32(336), v1, v2, v33)
												t274 := int32(load32(m.memory[int64(uint32(v4))+336:]))
												v30 = t274
												if v30 == i32(-1) {
													goto l7
												}
												goto l91
											}
										}
									l76:
										t219 := int32(load32(m.memory[uint32(v23):]))
										t220 := int32(load32(m.memory[uint32(v24):]))
										m.fn909(v4+i32(144), t219, t220, i32(1077133), i32(3))
										t221 := int32(load32(m.memory[int64(uint32(v4))+144:]))
										t222 := v4 + i32(464)
										v2 = t221
										p223 := i32(1)
										if v2 != 0 {
											p223 = v2
										}
										t224 := int32(load32(m.memory[int64(uint32(v4))+148:]))
										p225 := i32(0)
										if v2 != 0 {
											p225 = t224
										}
										m.fn865(t222, p223, p225)
										t226 := int32(load32(m.memory[uint32(v23):]))
										t227 := v4 + i32(136)
										v22 = t226
										t228 := int32(load32(m.memory[uint32(v24):]))
										t229 := v22
										v27 = t228
										m.fn909(t227, t229, v27, i32(1077136), i32(3))
										t230 := int32(load32(m.memory[int64(uint32(v4))+136:]))
										v2 = t230
										if v2 == 0 {
											goto l78
										}
										t231 := int32(load32(m.memory[int64(uint32(v4))+140:]))
										v22 = t231
										goto l79
									}
								l78:
									m.fn909(v4+i32(128), v22, v27, i32(1073490), i32(4))
									t258 := int32(load32(m.memory[int64(uint32(v4))+132:]))
									v22 = t258
									t259 := int32(load32(m.memory[int64(uint32(v4))+128:]))
									v2 = t259
								}
							l79:
								t260 := int32(load32(m.memory[int64(uint32(v1))+28:]))
								t262 := v4 + i32(336)
								p261 := i32(1)
								if v2 != 0 {
									p261 = v2
								}
								p263 := i32(0)
								if v2 != 0 {
									p263 = v22
								}
								t264 := int32(load32(m.memory[int64(uint32(v1))+32:]))
								t265 := int32(load32(m.memory[int64(uint32(t264))+16:]))
								m.t0[uint(t265)].(func(int32, int32, int32, int32))(t262, t260, p261, p263)
								t266 := int32(load32(m.memory[int64(uint32(v4))+348:]))
								v22 = t266
								t267 := int32(load32(m.memory[int64(uint32(v4))+344:]))
								v27 = t267
								t268 := int32(load32(m.memory[int64(uint32(v4))+340:]))
								v2 = t268
								t269 := int32(load32(m.memory[int64(uint32(v4))+336:]))
								v30 = t269
								if v30 == i32(-1) {
									{
										if v2 != i32(-1) {
											goto l96
										}
										t296 := int32(load32(m.memory[int64(uint32(v4))+468:]))
										t297 := v4 + i32(120)
										v28 = t296
										t298 := int32(load32(m.memory[int64(uint32(v4))+472:]))
										m.fn46(t297, v28, t298)
										t299 := int32(load32(m.memory[int64(uint32(v4))+124:]))
										if t299 != 0 {
											goto l96
										}
										t300 := int32(load32(m.memory[int64(uint32(v4))+464:]))
										m.fn16(t300, v28)
										goto l7
									}
								l96:
									t301 := int32(load32(m.memory[int64(uint32(v4))+472:]))
									store32(m.memory[int64(uint32(v16))+8:], uint32(t301))
									t302 := int64(load64(m.memory[int64(uint32(v4))+464:]))
									store64(m.memory[uint32(v16):], uint64(t302))
									store32(m.memory[int64(uint32(v4))+444:], uint32(v22))
									store32(m.memory[int64(uint32(v4))+440:], uint32(v27))
									store32(m.memory[int64(uint32(v4))+436:], uint32(v2))
									store32(m.memory[int64(uint32(v4))+280:], uint32(i32(-0x7fffffff)))
									m.fn1360(v20, v4+i32(436), v4+i32(280))
									store32(m.memory[int64(uint32(v4))+336:], uint32(i32(5)))
									m.fn1340(v5, v4+i32(336))
									goto l7
								}
								t270 := int64(load64(m.memory[int64(uint32(v4))+352:]))
								v41 = t270
								t271 := int32(load32(m.memory[int64(uint32(v4))+464:]))
								t272 := int32(load32(m.memory[int64(uint32(v4))+468:]))
								m.fn16(t271, t272)
								v42 = v22
								v43 = v27
								v44 = v2
								goto l89
							}
						l86:
							m.fn1409(v1)
							m.fn1267(v4+i32(336), v1, v2, v33)
							t304 := int32(load32(m.memory[int64(uint32(v4))+336:]))
							v30 = t304
							if v30 != i32(-1) {
								goto l91
							}
							m.fn1409(v1)
							goto l7
						}
					l82:
						m.fn1420(v4+i32(336), v1, v27, v33)
						t321 := int32(load32(m.memory[int64(uint32(v4))+348:]))
						v22 = t321
						t322 := int32(load32(m.memory[int64(uint32(v4))+344:]))
						v27 = t322
						t323 := int32(load32(m.memory[int64(uint32(v4))+340:]))
						v28 = t323
						{
							t324 := int32(load32(m.memory[int64(uint32(v4))+336:]))
							v30 = t324
							if v30 == i32(-1) {
								goto l100
							}
							t325 := int64(load64(m.memory[int64(uint32(v4))+352:]))
							v41 = t325
							v42 = v22
							v43 = v27
							v44 = v28
							goto l89
						}
					l100:
						store32(m.memory[int64(uint32(v4))+288:], uint32(v22))
						store32(m.memory[int64(uint32(v4))+284:], uint32(v27))
						store32(m.memory[int64(uint32(v4))+280:], uint32(v28))
						{
							{
								t326 := m.fn1421(v27, v22)
								if t326 != 0 {
									goto l101
								}
								m.fn894(v4 + i32(280))
								goto l102
							}
						l101:
							t327 := int32(load32(m.memory[int64(uint32(v4))+288:]))
							store32(m.memory[int64(uint32(v17))+8:], uint32(t327))
							t328 := int64(load64(m.memory[int64(uint32(v4))+280:]))
							store64(m.memory[uint32(v17):], uint64(t328))
							store32(m.memory[int64(uint32(v4))+336:], uint32(i32(-0x80000000)))
							m.fn338(v1, v4+i32(336))
						}
					l102:
						t329 := int32(load32(m.memory[int64(uint32(v2))+32:]))
						v31 = t329
						t330 := int32(load32(m.memory[int64(uint32(v2))+28:]))
						v28 = t330
					}
				l80:
					v30 = i32(0)
					store32(m.memory[int64(uint32(v4))+472:], uint32(i32(0)))
					store64(m.memory[int64(uint32(v4))+464:], uint64(i64(0x400000000)))
					store32(m.memory[int64(uint32(v4))+280:], uint32(v28))
					store32(m.memory[int64(uint32(v4))+284:], uint32(v28+v31*i32(44)))
				l137:
					{
						v28 = i32(0)
						{
							{
							l131:
								{
									t331 := m.fn904(v4 + i32(280))
									v27 = t331
									if v27 == 0 {
										t337 := int32(load32(m.memory[int64(uint32(v4))+472:]))
										v2 = t337
										if v2 != 0 {
											m.fn27(v4 + i32(248))
											t339 := int32(load32(m.memory[int64(uint32(v4))+468:]))
											v34 = t339
											t340 := v34
											v27 = v2 * i32(12)
											v23 = t340 + v27
											v28 = i32(0)
											v30 = v34
										l110:
											{
												if v27 == 0 {
													m.fn1165(v4 + i32(280))
													v26 = i32(0)
													v32 = v34
													v24 = i32(0)
												l130:
													v25 = v26
													{
														v31 = v32
														if v31 == v23 {
															memory_copy(m.memory, uint32(v4+i32(336)), uint32(v4+i32(280)), uint32(i32(56)))
															m.fn1168(v4+i32(436), v4+i32(336))
															t360 := int32(load32(m.memory[int64(uint32(v4))+444:]))
															v2 = t360
															if v2 != 0 {
																t407 := int32(load32(m.memory[int64(uint32(v4))+440:]))
																v44 = t407
																t408 := m.fn1234(v44, v2, v24)
																v42 = t408
																t409 := int32(load32(m.memory[int64(uint32(v4))+436:]))
																v45 = t409
																t410 := int64(load32(m.memory[int64(uint32(v4))+452:]))
																v37 = t410
																t411 := int32(load32(m.memory[int64(uint32(v4))+248:]))
																t412 := int32(load32(m.memory[int64(uint32(v4))+252:]))
																m.fn56(t411, t412)
																t413 := int32(load32(m.memory[int64(uint32(v4))+464:]))
																m.fn911(t413, v34)
																t414 := v4
																v41 = v37 | v41&i64(-0x100000000)
																store64(m.memory[int64(uint32(t414))+356:], uint64(v41))
																store32(m.memory[int64(uint32(v4))+352:], uint32(v42))
																store32(m.memory[int64(uint32(v4))+348:], uint32(v2))
																store32(m.memory[int64(uint32(v4))+344:], uint32(v44))
																store32(m.memory[int64(uint32(v4))+340:], uint32(v45))
																store32(m.memory[int64(uint32(v4))+336:], uint32(i32(-0x7ffffffe)))
																m.fn338(v1, v4+i32(336))
																v43 = v2
																goto l7
															}
															m.fn972(v4 + i32(436))
															v30 = v45
															goto l114
														}
														m.fn1166(v4 + i32(280))
														t356 := int32(load32(m.memory[uint32(v31):]))
														v2 = t356
														t357 := int32(load32(m.memory[int64(uint32(v2))+32:]))
														v22 = t357
														t358 := int32(load32(m.memory[int64(uint32(v2))+28:]))
														t359 := v4
														v2 = t358
														store32(m.memory[int64(uint32(t359))+424:], uint32(v2))
														store32(m.memory[int64(uint32(v4))+428:], uint32(v2+v22*i32(44)))
														v28 = i32(1)
														v26 = v25 + i32(1)
														v32 = v31 + i32(12)
														v30 = i32(0)
														{
														l117:
															{
																t361 := m.fn904(v4 + i32(424))
																v2 = t361
																if v2 == 0 {
																	if v25 != v24 {
																		goto l130
																	}
																	t415 := int32(m.memory[int64(uint32(v31))+4])
																	v24 = v25 + (t415|v28&v30)&i32(1)
																	goto l130
																}
																{
																	t362 := int32(load32(m.memory[uint32(v2+i32(4)):]))
																	v22 = t362
																	t363 := int32(load32(m.memory[uint32(v2+i32(8)):]))
																	t364 := v22
																	v27 = t363
																	t365 := m.fn15(t364, v27, i32(1077071), i32(2))
																	if t365 != 0 {
																		goto l116
																	}
																	t366 := m.fn15(v22, v27, i32(1077073), i32(2))
																	if t366 == 0 {
																		goto l117
																	}
																}
															l116:
																t367 := m.fn607(v22, v27, i32(1077073), i32(2))
																v36 = t367
																t368 := int32(load32(m.memory[uint32(v2+i32(16)):]))
																t369 := v4 + i32(96)
																v22 = t368
																t370 := int32(load32(m.memory[uint32(v2+i32(20)):]))
																t371 := v22
																v30 = t370
																m.fn909(t369, t371, v30, i32(1077075), i32(7))
																v27 = i32(1)
																{
																	t372 := int32(load32(m.memory[int64(uint32(v4))+96:]))
																	v29 = t372
																	if v29 == 0 {
																		goto l118
																	}
																	t373 := int32(load32(m.memory[int64(uint32(v4))+100:]))
																	m.fn1071(v4+i32(336), v29, t373)
																	t374 := int32(load32(m.memory[int64(uint32(v4))+340:]))
																	v27 = t374
																	p375 := i32(1)
																	if uint32(v27) > uint32(i32(1)) {
																		p375 = v27
																	}
																	v27 = p375
																	p376 := i32(1000)
																	if uint32(v27) < uint32(i32(1000)) {
																		p376 = v27
																	}
																	t377 := int32(m.memory[int64(uint32(v4))+336])
																	p378 := p376
																	if t377 != 0 {
																		p378 = i32(1)
																	}
																	v27 = p378
																}
															l118:
																m.fn909(v4+i32(88), v22, v30, i32(1077082), i32(7))
																v22 = i32(1)
																{
																	t379 := int32(load32(m.memory[int64(uint32(v4))+88:]))
																	v30 = t379
																	if v30 == 0 {
																		goto l119
																	}
																	t380 := int32(load32(m.memory[int64(uint32(v4))+92:]))
																	m.fn1071(v4+i32(336), v30, t380)
																	v22 = i32(1)
																	t381 := int32(m.memory[int64(uint32(v4))+336])
																	if t381 != 0 {
																		goto l119
																	}
																	{
																		t382 := int32(load32(m.memory[int64(uint32(v4))+340:]))
																		v22 = t382
																		if v22 != 0 {
																			goto l120
																		}
																		t383 := int32(load32(m.memory[int64(uint32(v4))+260:]))
																		if t383 == 0 {
																			goto l121
																		}
																		t384 := int64(load64(m.memory[int64(uint32(v4))+264:]))
																		t385 := int64(load64(m.memory[int64(uint32(v4))+272:]))
																		t386 := int32(load32(m.memory[int64(uint32(v31))+8:]))
																		v35 = t386
																		t387 := m.fn66(t384, t385, v35)
																		v37 = t387
																		t388 := int32(load32(m.memory[int64(uint32(v4))+252:]))
																		v29 = t388
																		v22 = v29 & int32(v37)
																		v39 = int64(uint64(v37)>>25) & i64(127) * i64(72340172838076673)
																		v46 = i32(0)
																		t389 := int32(load32(m.memory[int64(uint32(v4))+248:]))
																		v30 = t389
																	l126:
																		{
																			t390 := int64(load64(m.memory[uint32(v30+v22):]))
																			v40 = t390
																			v37 = v40 ^ v39
																			v37 = (v37 ^ i64(-1)) & (v37 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
																		l124:
																			{
																				if v37 == 0 {
																					if v40&(v40<<1)&i64(-0x7f7f7f7f7f7f7f80) == 0 {
																						t395 := v22
																						v46 = v46 + i32(8)
																						v22 = (t395 + v46) & v29
																						goto l126
																					}
																					goto l121
																				}
																				t391 := v35
																				v47 = v30 - (int32(uint32(int64(bits.TrailingZeros64(uint64(v37))))>>3)+v22)&v29<<3
																				t392 := int32(load32(m.memory[uint32(v47+i32(-8)):]))
																				if t391 == t392 {
																					t393 := int32(load32(m.memory[uint32(v47+i32(-4)):]))
																					v22 = t393 - v25 + i32(1)
																					p394 := i32(1)
																					if uint32(v22) > uint32(i32(1)) {
																						p394 = v22
																					}
																					v22 = p394
																					goto l119
																				}
																				v37 = (v37 + i64(-1)) & v37
																				goto l124
																			}
																		}
																	}
																l120:
																	p396 := i32(65534)
																	if uint32(v22) < uint32(i32(65534)) {
																		p396 = v22
																	}
																	v22 = p396
																}
															l119:
																m.fn1411(v4+i32(336), v1, v2, v33)
																t397 := int32(load32(m.memory[int64(uint32(v4))+348:]))
																v2 = t397
																t398 := int32(load32(m.memory[int64(uint32(v4))+344:]))
																v29 = t398
																t399 := int32(load32(m.memory[int64(uint32(v4))+340:]))
																v35 = t399
																{
																	t400 := int32(load32(m.memory[int64(uint32(v4))+336:]))
																	v30 = t400
																	if v30 == i32(-1) {
																		goto l127
																	}
																	t401 := int64(load64(m.memory[int64(uint32(v4))+352:]))
																	v41 = t401
																	v44 = v35
																	v43 = v29
																	v42 = v2
																	goto l128
																}
															l127:
																store32(m.memory[int64(uint32(v4))+452:], uint32(v22))
																store32(m.memory[int64(uint32(v4))+448:], uint32(v27))
																store32(m.memory[int64(uint32(v4))+444:], uint32(v2))
																store32(m.memory[int64(uint32(v4))+440:], uint32(v29))
																store32(m.memory[int64(uint32(v4))+436:], uint32(v35))
																m.fn1167(v4+i32(336), v4+i32(280), v4+i32(436))
																{
																	t402 := int32(load32(m.memory[int64(uint32(v4))+336:]))
																	v30 = t402
																	if v30 != i32(-1) {
																		goto l129
																	}
																	v28 = (v36 ^ i32(1)) & v28
																	v30 = i32(1)
																	goto l117
																}
															l129:
															}
															t403 := int64(load64(m.memory[int64(uint32(v4))+352:]))
															v41 = t403
															t404 := int32(load32(m.memory[int64(uint32(v4))+348:]))
															v42 = t404
															t405 := int32(load32(m.memory[int64(uint32(v4))+344:]))
															v43 = t405
															t406 := int32(load32(m.memory[int64(uint32(v4))+340:]))
															v44 = t406
														}
													l128:
														m.fn1259(v4 + i32(280))
														goto l114
													}
												l121:
													m.fn633(i32(1087080), i32(22), i32(1077092))
													panic("unreachable")
												}
												t341 := int32(load32(m.memory[int64(uint32(v30))+8:]))
												t342 := v4
												v31 = t341
												store32(m.memory[int64(uint32(t342))+436:], uint32(v31))
												t343 := int64(load64(m.memory[int64(uint32(v4))+264:]))
												t344 := int64(load64(m.memory[int64(uint32(v4))+272:]))
												t345 := m.fn66(t343, t344, v31)
												v37 = t345
												store32(m.memory[int64(uint32(v4))+280:], uint32(v4+i32(436)))
												m.fn713(v4+i32(248), v19)
												store32(m.memory[int64(uint32(v4))+340:], uint32(v4+i32(248)))
												store32(m.memory[int64(uint32(v4))+336:], uint32(v4+i32(280)))
												t346 := int32(load32(m.memory[int64(uint32(v4))+248:]))
												t347 := int32(load32(m.memory[int64(uint32(v4))+252:]))
												m.fn69(v4+i32(104), t346, t347, v37, v4+i32(336), i32(163))
												t348 := int32(load32(m.memory[int64(uint32(v4))+108:]))
												v2 = t348
												t349 := int32(load32(m.memory[int64(uint32(v4))+248:]))
												v22 = t349
												{
													t350 := int32(load32(m.memory[int64(uint32(v4))+104:]))
													if t350 != i32(1) {
														goto l109
													}
													v29 = v22 + v2
													t351 := int32(m.memory[uint32(v29)])
													v35 = t351
													t352 := v29
													v36 = int32(uint32(int32(v37)) >> 25)
													m.memory[uint32(t352)] = byte(v36)
													t353 := int32(load32(m.memory[int64(uint32(v4))+252:]))
													m.memory[uint32(v22+t353&(v2+i32(-8))+i32(8))] = byte(v36)
													store32(m.memory[uint32(v22-v2<<3+i32(-8)):], uint32(v31))
													t354 := int32(load32(m.memory[int64(uint32(v4))+260:]))
													store32(m.memory[int64(uint32(v4))+260:], uint32(t354+i32(1)))
													t355 := int32(load32(m.memory[int64(uint32(v4))+256:]))
													store32(m.memory[int64(uint32(v4))+256:], uint32(t355-v35&i32(1)))
												}
											l109:
												v30 = v30 + i32(12)
												store32(m.memory[uint32(v22+(i32(0)-v2)<<3+i32(-4)):], uint32(v28))
												v28 = v28 + i32(1)
												v27 = v27 + i32(-12)
												goto l110
											}
										}
										v2 = i32(0)
										t338 := int32(load32(m.memory[int64(uint32(v4))+468:]))
										v34 = t338
										v30 = v45
										goto l107
									}
									t332 := int32(load32(m.memory[uint32(v27+i32(4)):]))
									v2 = t332
									t333 := int32(load32(m.memory[uint32(v27+i32(8)):]))
									t334 := v2
									v22 = t333
									t335 := m.fn15(t334, v22, i32(1077108), i32(5))
									if t335 != 0 {
										goto l104
									}
									t336 := m.fn15(v2, v22, i32(1077113), i32(5))
									if t336 == 0 {
										t416 := m.fn15(v2, v22, i32(1077118), i32(5))
										if t416 != 0 {
											goto l104
										}
										t417 := m.fn15(v2, v22, i32(1073488), i32(2))
										if t417 == 0 {
											goto l131
										}
										store32(m.memory[int64(uint32(v4))+344:], uint32(v30))
										m.memory[int64(uint32(v4))+340] = byte(i32(0))
										store32(m.memory[int64(uint32(v4))+336:], uint32(v27))
										m.fn1422(v4+i32(464), v4+i32(336))
										v28 = i32(1)
										goto l131
									}
									goto l104
								}
							l114:
								;
								var p418 int32
								if v31 != v23 {
									p418 = 1
								}
								v2 = p418
								t419 := int32(load32(m.memory[int64(uint32(v4))+248:]))
								t420 := int32(load32(m.memory[int64(uint32(v4))+252:]))
								m.fn56(t419, t420)
							}
						l107:
							t421 := int32(load32(m.memory[int64(uint32(v4))+464:]))
							m.fn911(t421, v34)
							if v2 != 0 {
								v45 = i32(-1)
								v2 = v9
								if v30 == i32(-1) {
									goto l133
								}
								goto l89
							}
							v45 = v30
							goto l7
						}
					l104:
						v28 = v30 + v28&i32(1)
						t422 := m.fn773(v2, v22, i32(1077108), i32(5))
						v30 = t422
						t423 := int32(load32(m.memory[int64(uint32(v27))+32:]))
						v22 = t423 * i32(44)
						t424 := int32(load32(m.memory[int64(uint32(v27))+28:]))
						v2 = t424
					l136:
						if v22 == 0 {
							v30 = v28 + i32(1)
							goto l137
						}
						{
							t425 := int32(load32(m.memory[uint32(v2):]))
							if t425 == i32(-1) {
								goto l135
							}
							t426 := int32(load32(m.memory[uint32(v2+i32(4)):]))
							t427 := int32(load32(m.memory[uint32(v2+i32(8)):]))
							t428 := m.fn773(t426, t427, i32(1073488), i32(2))
							if t428 == 0 {
								goto l135
							}
							store32(m.memory[int64(uint32(v4))+344:], uint32(v28))
							m.memory[int64(uint32(v4))+340] = byte(v30)
							store32(m.memory[int64(uint32(v4))+336:], uint32(v2))
							m.fn1422(v4+i32(464), v4+i32(336))
						}
					l135:
						v2 = v2 + i32(44)
						v22 = v22 + i32(-44)
						goto l136
					}
				l75:
				}
				store32(m.memory[int64(uint32(v4))+336:], uint32(i32(8)))
				m.fn1340(v5, v4+i32(336))
				goto l7
			l91:
				t447 := int64(load64(m.memory[int64(uint32(v4))+352:]))
				v41 = t447
				t448 := int32(load32(m.memory[int64(uint32(v4))+348:]))
				v42 = t448
				t449 := int32(load32(m.memory[int64(uint32(v4))+344:]))
				v43 = t449
				t450 := int32(load32(m.memory[int64(uint32(v4))+340:]))
				v44 = t450
				goto l89
			}
		l60:
			t451 := int32(load32(m.memory[int64(uint32(v4))+464:]))
			m.fn1415(t451, v36)
		}
	l45:
		m.fn44(v26, v31)
		if v30 == i32(-1) {
			store32(m.memory[int64(uint32(v4))+244:], uint32(v29))
			store32(m.memory[int64(uint32(v4))+240:], uint32(v2))
			store32(m.memory[int64(uint32(v4))+236:], uint32(v35))
			m.fn1271(v1, v4+i32(236))
			goto l7
		}
		v42 = v29
		v43 = v2
		v44 = v35
		v41 = v38
		goto l89
	l30:
		m.fn1409(v1)
		t452 := int32(load32(m.memory[uint32(v32):]))
		t453 := int32(load32(m.memory[uint32(v36):]))
		m.fn826(v4+i32(32), i32(1), t452, t453, i32(1077208))
		t454 := int32(load32(m.memory[int64(uint32(v4))+32:]))
		t455 := int32(load32(m.memory[int64(uint32(v4))+36:]))
		m.fn1423(v4+i32(24), t454, t455)
		t456 := int32(m.memory[int64(uint32(v4))+25])
		v31 = t456
		t457 := int32(m.memory[int64(uint32(v4))+24])
		v29 = t457
		m.fn1420(v4+i32(336), v1, v2, v33)
		t458 := int32(load32(m.memory[int64(uint32(v4))+348:]))
		v22 = t458
		t459 := int32(load32(m.memory[int64(uint32(v4))+344:]))
		v2 = t459
		t460 := int32(load32(m.memory[int64(uint32(v4))+340:]))
		v27 = t460
		{
			t461 := int32(load32(m.memory[int64(uint32(v4))+336:]))
			v30 = t461
			if v30 == i32(-1) {
				t463 := m.fn1188(v33)
				m.fn1368(v2, v22, t463)
				t464 := int32(load32(m.memory[uint32(v23):]))
				t465 := int32(load32(m.memory[uint32(v24):]))
				m.fn909(v4+i32(16), t464, t465, i32(1073226), i32(2))
				{
					t466 := int32(load32(m.memory[int64(uint32(v4))+16:]))
					v28 = t466
					if v28 == 0 {
						goto l142
					}
					t467 := int32(load32(m.memory[int64(uint32(v1))+28:]))
					t468 := int32(load32(m.memory[int64(uint32(v4))+20:]))
					t469 := int32(load32(m.memory[int64(uint32(v1))+32:]))
					t470 := int32(load32(m.memory[int64(uint32(t469))+20:]))
					m.t0[uint(t470)].(func(int32, int32, int32, int32))(v4+i32(212), t467, v28, t468)
					goto l143
				}
			l142:
				store32(m.memory[int64(uint32(v4))+212:], uint32(i32(-1)))
			l143:
				{
					t471 := m.fn23(v2, v22)
					if t471 != 0 {
						v28 = i32(-1)
						{
							t476 := int32(load32(m.memory[int64(uint32(v4))+212:]))
							if t476 == i32(-1) {
								goto l145
							}
							t477 := int32(load32(m.memory[int64(uint32(v4))+220:]))
							store32(m.memory[int64(uint32(v4))+344:], uint32(t477))
							t478 := int64(load64(m.memory[int64(uint32(v4))+212:]))
							store64(m.memory[int64(uint32(v4))+336:], uint64(t478))
							v28 = i32(6)
						}
					l145:
						t479 := int32(load32(m.memory[int64(uint32(v4))+344:]))
						store32(m.memory[int64(uint32(v18))+8:], uint32(t479))
						t480 := int64(load64(m.memory[int64(uint32(v4))+336:]))
						store64(m.memory[uint32(v18):], uint64(t480))
						store32(m.memory[int64(uint32(v4))+316:], uint32(v27))
						store32(m.memory[int64(uint32(v4))+312:], uint32(v2))
						store32(m.memory[int64(uint32(v4))+308:], uint32(v2))
						store32(m.memory[int64(uint32(v4))+280:], uint32(v28))
						store32(m.memory[int64(uint32(v4))+320:], uint32(v2+v22*i32(28)))
						m.fn896(v4+i32(436), v4+i32(280))
						{
							t481 := int32(load32(m.memory[int64(uint32(v4))+436:]))
							if t481 == i32(-1) {
								goto l146
							}
							m.fn901(v4+i32(336), v4+i32(280))
							v22 = i32(1)
							v27 = i32(28)
							t482 := int32(load32(m.memory[int64(uint32(v4))+336:]))
							t483 := v4 + i32(8)
							v2 = t482 + i32(1)
							p484 := i32(-1)
							if v2 != 0 {
								p484 = v2
							}
							v2 = p484
							p485 := i32(4)
							if uint32(v2) > uint32(i32(4)) {
								p485 = v2
							}
							m.fn59(t483, p485, i32(4), i32(28))
							t486 := int32(load32(m.memory[int64(uint32(v4))+8:]))
							v2 = t486
							t487 := int32(load32(m.memory[int64(uint32(v4))+12:]))
							v28 = t487
							t488 := int32(load32(m.memory[int64(uint32(v4))+460:]))
							store32(m.memory[int64(uint32(v28))+24:], uint32(t488))
							t489 := int64(load64(m.memory[int64(uint32(v4))+452:]))
							store64(m.memory[int64(uint32(v28))+16:], uint64(t489))
							t490 := int64(load64(m.memory[int64(uint32(v4))+444:]))
							store64(m.memory[int64(uint32(v28))+8:], uint64(t490))
							t491 := int64(load64(m.memory[int64(uint32(v4))+436:]))
							store64(m.memory[uint32(v28):], uint64(t491))
							store32(m.memory[int64(uint32(v4))+432:], uint32(i32(1)))
							store32(m.memory[int64(uint32(v4))+428:], uint32(v28))
							store32(m.memory[int64(uint32(v4))+424:], uint32(v2))
							memory_copy(m.memory, uint32(v4+i32(336)), uint32(v4+i32(280)), uint32(i32(44)))
						l149:
							{
								m.fn896(v4+i32(248), v4+i32(336))
								t492 := int32(load32(m.memory[int64(uint32(v4))+248:]))
								if t492 == i32(-1) {
									m.fn899(v4 + i32(336))
									t503 := int32(load32(m.memory[int64(uint32(v4))+432:]))
									t504 := v4
									v2 = t503
									store32(m.memory[int64(uint32(t504))+472:], uint32(v2))
									t505 := int64(load64(m.memory[int64(uint32(v4))+424:]))
									store64(m.memory[int64(uint32(v4))+464:], uint64(t505))
									if v2 == 0 {
										goto l150
									}
									t506 := int32(load32(m.memory[int64(uint32(v4))+472:]))
									store32(m.memory[int64(uint32(v17))+8:], uint32(t506))
									t507 := int64(load64(m.memory[int64(uint32(v4))+464:]))
									store64(m.memory[uint32(v17):], uint64(t507))
									store32(m.memory[int64(uint32(v4))+336:], uint32(i32(-0x80000000)))
									m.fn338(v1, v4+i32(336))
									goto l7
								}
								{
									t493 := int32(load32(m.memory[int64(uint32(v4))+424:]))
									if v22 != t493 {
										goto l148
									}
									m.fn901(v4+i32(464), v4+i32(336))
									t494 := int32(load32(m.memory[int64(uint32(v4))+464:]))
									t495 := v4 + i32(424)
									v2 = t494 + i32(1)
									p496 := i32(-1)
									if v2 != 0 {
										p496 = v2
									}
									m.fn892(t495, p496)
									t497 := int32(load32(m.memory[int64(uint32(v4))+428:]))
									v28 = t497
								}
							l148:
								v2 = v28 + v27
								t498 := int32(load32(m.memory[int64(uint32(v4))+272:]))
								store32(m.memory[int64(uint32(v2))+24:], uint32(t498))
								t499 := int64(load64(m.memory[int64(uint32(v4))+264:]))
								store64(m.memory[int64(uint32(v2))+16:], uint64(t499))
								t500 := int64(load64(m.memory[int64(uint32(v4))+256:]))
								store64(m.memory[int64(uint32(v2))+8:], uint64(t500))
								t501 := int64(load64(m.memory[int64(uint32(v4))+248:]))
								store64(m.memory[uint32(v2):], uint64(t501))
								t502 := v4
								v22 = v22 + i32(1)
								store32(m.memory[int64(uint32(t502))+432:], uint32(v22))
								v27 = v27 + i32(28)
								goto l149
							}
						}
					l146:
						store32(m.memory[int64(uint32(v4))+472:], uint32(i32(0)))
						store64(m.memory[int64(uint32(v4))+464:], uint64(i64(0x400000000)))
						m.fn899(v4 + i32(280))
					l150:
						m.fn894(v4 + i32(464))
						goto l7
					}
					t472 := int64(load64(m.memory[int64(uint32(v4))+212:]))
					store64(m.memory[uint32(v15):], uint64(t472))
					t473 := int32(load32(m.memory[int64(uint32(v4))+220:]))
					store32(m.memory[int64(uint32(v15))+8:], uint32(t473))
					t475 := v4
					p474 := v31
					if v29 != 0 {
						p474 = i32(1)
					}
					m.memory[int64(uint32(t475))+360] = byte(p474)
					store32(m.memory[int64(uint32(v4))+344:], uint32(v22))
					store32(m.memory[int64(uint32(v4))+340:], uint32(v2))
					store32(m.memory[int64(uint32(v4))+336:], uint32(v27))
					m.fn338(v1, v4+i32(336))
					goto l7
				}
			}
			t462 := int64(load64(m.memory[int64(uint32(v4))+352:]))
			v41 = t462
			v42 = v22
			v43 = v2
			v44 = v27
			goto l89
		}
	}
l89:
	store64(m.memory[int64(uint32(v0))+16:], uint64(v41))
	store32(m.memory[int64(uint32(v0))+12:], uint32(v42))
	store32(m.memory[int64(uint32(v0))+8:], uint32(v43))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v44))
	store32(m.memory[uint32(v0):], uint32(v30))
l5:
	m.g0 = v4 + i32(480)
}
func (m *Module) fn1268(v0 int32) {
	m.fn969(v0)
	m.fn894(v0 + i32(12))
}
func (m *Module) fn1269(v0 int32) {
	var v1, v2, v3 int32
	t0 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	v1 = t0
	t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	v2 = t1
	v3 = v2
l1:
	{
		if v1 == 0 {
			goto l0
		}
		t2 := int32(load32(m.memory[uint32(v3):]))
		t3 := int32(load32(m.memory[uint32(v3+i32(4)):]))
		m.fn134(t2, t3)
		t4 := int32(load32(m.memory[uint32(v3+i32(12)):]))
		t5 := int32(load32(m.memory[uint32(v3+i32(16)):]))
		m.fn134(t4, t5)
		v1 = v1 + i32(-1)
		v3 = v3 + i32(36)
		goto l1
	}
l0:
	t6 := int32(load32(m.memory[uint32(v0):]))
	m.fn136(t6, v2, i32(4), i32(36))
}
func (m *Module) fn1270(v0, v1 int32) {
	m.fn1409(v1)
	t0 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	store32(m.memory[int64(uint32(v0))+8:], uint32(t0))
	t1 := int64(load64(m.memory[uint32(v1):]))
	store64(m.memory[uint32(v0):], uint64(t1))
	m.fn894(v1 + i32(12))
}
func (m *Module) fn1271(v0, v1 int32) {
	var v2, v3 int32
	t0 := int32(load32(m.memory[uint32(v1):]))
	v2 = t0
	t1 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	t2 := v0
	v3 = t1
	t3 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	m.fn1336(t2, v3, t3)
	m.fn80(v2, v3)
}
func (m *Module) fn1272(v0 int32) {
	var v1 int32
	t0 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	v1 = t0
	t1 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	m.fn83(v1, t1)
	t2 := int32(load32(m.memory[uint32(v0):]))
	m.fn84(t2, v1)
}
func (m *Module) fn1273(v0 int32) {
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
					m.fn134(t8, t9)
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
func (m *Module) fn1274(v0 int32) {
	m.fn1272(v0 + i32(36))
	m.fn57(v0)
}
func (m *Module) fn1275(v0 int32) {
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
					v7 = v6 + (i32(0)-int32(uint32(int64(bits.TrailingZeros64(uint64(v5))))>>3))*i32(36)
					t6 := int32(load32(m.memory[uint32(v7+i32(-36)):]))
					t7 := int32(load32(m.memory[uint32(v7+i32(-32)):]))
					m.fn16(t6, t7)
					m.fn169(v7 + i32(-24))
					v4 = v4 + i32(-1)
					v5 = (v5 + i64(-1)) & v5
					goto l4
				}
				v6 = v6 + i32(-288)
				t5 := int64(load64(m.memory[uint32(v0):]))
				v5 = (t5 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
				v0 = v0 + i32(8)
				goto l3
			}
		}
	l1:
		m.fn39(v1+i32(4), i32(36), i32(8), v2+i32(1))
		t8 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		t9 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t10 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		m.fn40(v3-t8, t9, t10)
	}
l0:
	m.g0 = v1 + i32(16)
}
func (m *Module) fn1276(v0, v1 int32) {
	var v2, v3, v4, v5, v6, v7, v8, v9, v10, v11 int32
	t0 := m.g0
	v2 = t0 - i32(256)
	m.g0 = v2
	store32(m.memory[int64(uint32(v2))+208:], uint32(i32(33686018)))
	v3 = i32(0)
	m.memory[int64(uint32(v2))+192] = byte(i32(0))
	store32(m.memory[int64(uint32(v2))+184:], uint32(i32(33686018)))
	m.memory[int64(uint32(v2))+168] = byte(i32(0))
	store32(m.memory[int64(uint32(v2))+160:], uint32(i32(33686018)))
	m.memory[int64(uint32(v2))+144] = byte(i32(0))
	store32(m.memory[int64(uint32(v2))+136:], uint32(i32(33686018)))
	m.memory[int64(uint32(v2))+120] = byte(i32(0))
	store32(m.memory[int64(uint32(v2))+112:], uint32(i32(33686018)))
	m.memory[int64(uint32(v2))+96] = byte(i32(0))
	store32(m.memory[int64(uint32(v2))+88:], uint32(i32(33686018)))
	m.memory[int64(uint32(v2))+72] = byte(i32(0))
	store32(m.memory[int64(uint32(v2))+64:], uint32(i32(33686018)))
	m.memory[int64(uint32(v2))+48] = byte(i32(0))
	store32(m.memory[int64(uint32(v2))+40:], uint32(i32(33686018)))
	m.memory[int64(uint32(v2))+24] = byte(i32(0))
	store32(m.memory[int64(uint32(v2))+16:], uint32(i32(33686018)))
	m.memory[uint32(v2)] = byte(i32(0))
	{
		if v1 == 0 {
			goto l0
		}
		t1 := int32(load32(m.memory[int64(uint32(v1))+32:]))
		v4 = t1 * i32(44)
		t2 := int32(load32(m.memory[int64(uint32(v1))+28:]))
		v5 = t2
		v6 = i32(0)
	l5:
		{
			if v3 == i32(216) {
				goto l0
			}
			v7 = v2 + v3
			t3 := v2
			v6 = v6 + i32(1)
			store32(m.memory[int64(uint32(t3))+220:], uint32(v6))
			store32(m.memory[int64(uint32(v2))+228:], uint32(i32(5)))
			store32(m.memory[int64(uint32(v2))+224:], uint32(v2+i32(220)))
			m.fn73(v2+i32(232), i32(1067219), v2+i32(224))
			v3 = v3 + i32(24)
			t4 := int32(load32(m.memory[int64(uint32(v2))+232:]))
			v8 = t4
			t5 := int32(load32(m.memory[int64(uint32(v2))+236:]))
			v9 = t5
			t6 := int32(load32(m.memory[int64(uint32(v2))+240:]))
			v10 = t6
			v11 = v4
			v1 = v5
			{
			l4:
				if v11 == 0 {
					goto l1
				}
				{
					t7 := int32(load32(m.memory[uint32(v1):]))
					if t7 == i32(-1) {
						goto l2
					}
					t8 := m.fn847(v1, i32(1074411), i32(53), v9, v10)
					if t8 != 0 {
						goto l3
					}
				}
			l2:
				v1 = v1 + i32(44)
				v11 = v11 + i32(-44)
				goto l4
			l3:
				t9 := int32(load32(m.memory[uint32(v1+i32(28)):]))
				t10 := int32(load32(m.memory[uint32(v1+i32(32)):]))
				m.fn1356(v2+i32(232), t9, t10)
				t11 := int64(load64(m.memory[int64(uint32(v2))+248:]))
				store64(m.memory[int64(uint32(v7))+16:], uint64(t11))
				t12 := int64(load64(m.memory[int64(uint32(v2))+240:]))
				store64(m.memory[int64(uint32(v7))+8:], uint64(t12))
				t13 := int64(load64(m.memory[int64(uint32(v2))+232:]))
				store64(m.memory[uint32(v7):], uint64(t13))
			}
		l1:
			m.fn16(v8, v9)
			goto l5
		}
	}
l0:
	memory_copy(m.memory, uint32(v0), uint32(v2), uint32(i32(216)))
	m.g0 = v2 + i32(256)
}
func (m *Module) fn1277(v0 int32) {
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
	m.fn1043(v3)
	v3 = v3 + i32(32)
	goto l1
l0:
	t2 := int32(load32(m.memory[uint32(v0):]))
	m.fn136(t2, v2, i32(8), i32(32))
}
func (m *Module) fn1278(v0 int32) {
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
					v7 = v6 + (i32(0)-int32(uint32(int64(bits.TrailingZeros64(uint64(v5))))>>3))*i32(680)
					t6 := int32(load32(m.memory[uint32(v7+i32(-680)):]))
					t7 := int32(load32(m.memory[uint32(v7+i32(-676)):]))
					m.fn16(t6, t7)
					m.fn757(v7 + i32(-16))
					v4 = v4 + i32(-1)
					v5 = (v5 + i64(-1)) & v5
					goto l4
				}
				v6 = v6 + i32(-5440)
				t5 := int64(load64(m.memory[uint32(v0):]))
				v5 = (t5 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
				v0 = v0 + i32(8)
				goto l3
			}
		}
	l1:
		m.fn39(v1+i32(4), i32(680), i32(8), v2+i32(1))
		t8 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		t9 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t10 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		m.fn40(v3-t8, t9, t10)
	}
l0:
	m.g0 = v1 + i32(16)
}
func (m *Module) fn1279(v0 int32) {
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
					v7 = v6 + (i32(0)-int32(uint32(int64(bits.TrailingZeros64(uint64(v5))))>>3))*i32(36)
					t6 := int32(load32(m.memory[uint32(v7+i32(-36)):]))
					t7 := int32(load32(m.memory[uint32(v7+i32(-32)):]))
					m.fn16(t6, t7)
					m.fn756(v7 + i32(-24))
					v4 = v4 + i32(-1)
					v5 = (v5 + i64(-1)) & v5
					goto l4
				}
				v6 = v6 + i32(-288)
				t5 := int64(load64(m.memory[uint32(v0):]))
				v5 = (t5 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
				v0 = v0 + i32(8)
				goto l3
			}
		}
	l1:
		m.fn39(v1+i32(4), i32(36), i32(8), v2+i32(1))
		t8 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		t9 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t10 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		m.fn40(v3-t8, t9, t10)
	}
l0:
	m.g0 = v1 + i32(16)
}
func (m *Module) fn1280(v0, v1, v2, v3, v4, v5 int32) {
	var v6 int32
	t0 := m.g0
	v6 = t0 - i32(64)
	m.g0 = v6
	{
		t1 := m.fn1039(v1, v4, v5)
		v5 = t1
		if v5 == 0 {
			goto l0
		}
		t2 := int32(load32(m.memory[int64(uint32(v5))+4:]))
		t3 := int32(load32(m.memory[int64(uint32(v5))+8:]))
		m.fn774(v6+i32(36), v2, v3, t2, t3)
		m.fn780(v6+i32(12), v6+i32(36))
		t4 := int32(load32(m.memory[int64(uint32(v6))+12:]))
		if t4 == i32(-1) {
			goto l0
		}
		t5 := int64(load64(m.memory[int64(uint32(v6))+12:]))
		store64(m.memory[uint32(v0):], uint64(t5))
		t6 := int32(load32(m.memory[int64(uint32(v6))+20:]))
		store32(m.memory[int64(uint32(v0))+8:], uint32(t6))
		t7 := int32(load32(m.memory[int64(uint32(v6))+24:]))
		t8 := int32(load32(m.memory[int64(uint32(v6))+28:]))
		m.fn134(t7, t8)
		goto l1
	}
l0:
	store32(m.memory[uint32(v0):], uint32(i32(-1)))
l1:
	m.g0 = v6 + i32(64)
}
func (m *Module) fn1281(v0, v1, v2 int32) {
	var v3, v4, v5, v6, v7, v8 int32
	t0 := m.g0
	v3 = t0 - i32(512)
	m.g0 = v3
	v4 = i32(0)
	store32(m.memory[int64(uint32(v3))+24:], uint32(i32(0)))
	store64(m.memory[int64(uint32(v3))+16:], uint64(i64(0x800000000)))
	m.fn868(v3+i32(28), v1, v2)
	store32(m.memory[int64(uint32(v3))+52:], uint32(i32(2)))
	store32(m.memory[int64(uint32(v3))+48:], uint32(i32(1078581)))
	store32(m.memory[int64(uint32(v3))+44:], uint32(i32(58)))
	store32(m.memory[int64(uint32(v3))+40:], uint32(i32(1074346)))
	v5 = v3 + i32(272) + i32(216)
	v6 = v3 + i32(500)
	v7 = i32(8)
l1:
	{
		t1 := m.fn863(v3 + i32(28))
		v2 = t1
		if v2 == 0 {
			goto l0
		}
		v1 = v2 + i32(28)
		t2 := int32(load32(m.memory[uint32(v1):]))
		v8 = v2 + i32(32)
		t3 := int32(load32(m.memory[uint32(v8):]))
		t4 := m.fn1097(t2, t3, i32(1074346), i32(58), i32(1083073), i32(2))
		v2 = t4
		if v2 == 0 {
			goto l1
		}
		{
			{
				t5 := int32(load32(m.memory[uint32(v1):]))
				t6 := int32(load32(m.memory[uint32(v8):]))
				t7 := m.fn886(t5, t6, i32(1074346), i32(58), i32(1083084), i32(6))
				v1 = t7
				if v1 != 0 {
					goto l2
				}
				v1 = i32(0)
				goto l3
			}
		l2:
			t8 := int32(load32(m.memory[uint32(v1+i32(28)):]))
			t9 := int32(load32(m.memory[uint32(v1+i32(32)):]))
			t10 := m.fn886(t8, t9, i32(1074411), i32(53), i32(1074905), i32(8))
			v1 = t10
		}
	l3:
		m.fn1276(v3+i32(56), v1)
		t11 := v3 + i32(8)
		v1 = v2 + i32(16)
		t12 := int32(load32(m.memory[uint32(v1):]))
		v8 = v2 + i32(20)
		t13 := int32(load32(m.memory[uint32(v8):]))
		m.fn1046(t11, t12, t13, i32(1074346), i32(58), i32(1074404), i32(4))
		t14 := int32(load32(m.memory[int64(uint32(v3))+8:]))
		t15 := v6
		v2 = t14
		p16 := i32(1073232)
		if v2 != 0 {
			p16 = v2
		}
		t17 := int32(load32(m.memory[int64(uint32(v3))+12:]))
		p18 := i32(4)
		if v2 != 0 {
			p18 = t17
		}
		m.fn51(t15, p16, p18)
		t19 := int32(load32(m.memory[uint32(v1):]))
		t20 := int32(load32(m.memory[uint32(v8):]))
		m.fn1046(v3, t19, t20, i32(1074346), i32(58), i32(1074408), i32(3))
		t21 := int32(load32(m.memory[uint32(v3):]))
		t22 := int32(load32(m.memory[int64(uint32(v3))+4:]))
		m.fn1041(v5, t21, t22)
		memory_copy(m.memory, uint32(v3+i32(272)), uint32(v3+i32(56)), uint32(i32(216)))
		{
			t23 := int32(load32(m.memory[int64(uint32(v3))+16:]))
			if v4 != t23 {
				goto l4
			}
			m.fn1152(v3 + i32(16))
			t24 := int32(load32(m.memory[int64(uint32(v3))+20:]))
			v7 = t24
		}
	l4:
		memory_copy(m.memory, uint32(v7+v4*i32(240)), uint32(v3+i32(272)), uint32(i32(240)))
		t25 := v3
		v4 = v4 + i32(1)
		store32(m.memory[int64(uint32(t25))+24:], uint32(v4))
		goto l1
	}
l0:
	t26 := int32(load32(m.memory[int64(uint32(v3))+28:]))
	t27 := int32(load32(m.memory[int64(uint32(v3))+32:]))
	m.fn44(t26, t27)
	t28 := int32(load32(m.memory[int64(uint32(v3))+24:]))
	store32(m.memory[int64(uint32(v0))+8:], uint32(t28))
	t29 := int64(load64(m.memory[int64(uint32(v3))+16:]))
	store64(m.memory[uint32(v0):], uint64(t29))
	m.g0 = v3 + i32(512)
}
func (m *Module) fn1282(v0, v1 int32) int32 {
	var v2 int64
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
		if t0 != 0 {
			t1 := int64(load64(m.memory[int64(uint32(v0))+16:]))
			t2 := int64(load64(m.memory[int64(uint32(v0))+24:]))
			t3 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t4 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			t5 := m.fn540(t1, t2, t3, t4)
			v2 = t5
			t6 := int32(load32(m.memory[uint32(v0):]))
			t7 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t8 := m.fn644(t6, t7, v2, v1)
			v0 = t8
			p9 := i32(0)
			if v0 != 0 {
				p9 = v0 + i32(-12)
			}
			return p9
		}
		return i32(0)
	}
}
func (m *Module) fn1283(v0, v1, v2, v3 int32) {
	var v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14, v15, v16, v17, v18, v19 int32
	var v20 int64
	var v21, v22, v23, v24 int32
	var v25 int64
	var v26, v27, v28 int32
	t0 := m.g0
	v4 = t0 - i32(672)
	m.g0 = v4
	t1 := int32(load32(m.memory[int64(uint32(v1))+32:]))
	v5 = t1
	t2 := int32(load32(m.memory[int64(uint32(v1))+28:]))
	t3 := v4
	v1 = t2
	store32(m.memory[int64(uint32(t3))+152:], uint32(v1))
	store32(m.memory[int64(uint32(v4))+156:], uint32(v1+v5*i32(44)))
	v6 = v4 + i32(208) + i32(4)
	v7 = v4 + i32(568) + i32(12)
	v8 = v4 + i32(208) + i32(12)
	v9 = v4 + i32(524) + i32(4)
	v10 = v4 + i32(480) + i32(4)
	v11 = v4 + i32(208) | i32(4)
	t4 := int32(load32(m.memory[int64(uint32(v2))+16:]))
	v12 = t4
	t5 := int32(load32(m.memory[int64(uint32(v2))+12:]))
	v13 = t5
	t6 := int32(load32(m.memory[int64(uint32(v2))+8:]))
	v14 = t6
	t7 := int32(load32(m.memory[int64(uint32(v2))+4:]))
	v15 = t7
	t8 := int32(load32(m.memory[uint32(v2):]))
	v16 = t8
	{
	l2:
		{
			{
				{
					{
						{
							t9 := m.fn904(v4 + i32(152))
							v1 = t9
							if v1 == 0 {
								store32(m.memory[uint32(v0):], uint32(i32(-1)))
								goto l7
							}
							t10 := m.fn847(v1, i32(1073848), i32(59), i32(1077491), i32(16))
							if t10 != 0 {
								t212 := int32(load32(m.memory[uint32(v1+i32(28)):]))
								t213 := int32(load32(m.memory[uint32(v1+i32(32)):]))
								t214 := m.fn1361(t212, t213, i32(1083136), i32(4))
								v1 = t214
								if v1 == 0 {
									goto l2
								}
								m.fn1283(v4+i32(208), v1, v2, v3)
								t215 := int32(load32(m.memory[int64(uint32(v4))+208:]))
								v1 = t215
								if v1 == i32(-1) {
									goto l2
								}
								t216 := int32(load32(m.memory[int64(uint32(v4))+228:]))
								store32(m.memory[int64(uint32(v0))+20:], uint32(t216))
								t217 := int64(load64(m.memory[int64(uint32(v4))+220:]))
								store64(m.memory[int64(uint32(v0))+12:], uint64(t217))
								t218 := int64(load64(m.memory[int64(uint32(v4))+212:]))
								store64(m.memory[int64(uint32(v0))+4:], uint64(t218))
								store32(m.memory[uint32(v0):], uint32(v1))
								goto l7
							}
							t11 := int32(load32(m.memory[int64(uint32(v1))+36:]))
							v5 = t11
							if v5 == 0 {
								goto l2
							}
							t12 := int32(load32(m.memory[int64(uint32(v1))+40:]))
							t13 := m.fn1337(v5+i32(8), t12, i32(1074346), i32(58))
							if t13 != 0 {
								goto l2
							}
							{
								t14 := int32(load32(m.memory[int64(uint32(v1))+4:]))
								v5 = t14
								t15 := int32(load32(m.memory[int64(uint32(v1))+8:]))
								t16 := v5
								v17 = t15
								t17 := m.fn15(t16, v17, i32(1078581), i32(2))
								if t17 != 0 {
									goto l3
								}
								t18 := m.fn15(v5, v17, i32(1083104), i32(5))
								if t18 == 0 {
									t44 := m.fn15(v5, v17, i32(1083109), i32(5))
									if t44 != 0 {
										m.fn1283(v4+i32(208), v1, v2, v3)
										t208 := int32(load32(m.memory[int64(uint32(v4))+208:]))
										v1 = t208
										if v1 == i32(-1) {
											goto l2
										}
										t209 := int32(load32(m.memory[int64(uint32(v4))+228:]))
										store32(m.memory[int64(uint32(v0))+20:], uint32(t209))
										t210 := int64(load64(m.memory[int64(uint32(v4))+220:]))
										store64(m.memory[int64(uint32(v0))+12:], uint64(t210))
										t211 := int64(load64(m.memory[int64(uint32(v4))+212:]))
										store64(m.memory[int64(uint32(v0))+4:], uint64(t211))
										store32(m.memory[uint32(v0):], uint32(v1))
										goto l7
									}
									{
										{
											t45 := m.fn15(v5, v17, i32(1083114), i32(12))
											if t45 != 0 {
												v5 = v1 + i32(28)
												t50 := int32(load32(m.memory[uint32(v5):]))
												v17 = v1 + i32(32)
												t51 := int32(load32(m.memory[uint32(v17):]))
												t52 := m.fn1097(t50, t51, i32(1074411), i32(53), i32(1083171), i32(3))
												v1 = t52
												if v1 == 0 {
													t118 := int32(load32(m.memory[uint32(v5):]))
													t119 := int32(load32(m.memory[uint32(v17):]))
													t120 := m.fn1097(t118, t119, i32(1074346), i32(58), i32(1083174), i32(6))
													v1 = t120
													if v1 == 0 {
														t158 := int32(load32(m.memory[uint32(v5):]))
														t159 := int32(load32(m.memory[uint32(v17):]))
														t160 := m.fn1097(t158, t159, i32(1073932), i32(54), i32(1077424), i32(5))
														v1 = t160
														if v1 == 0 {
															goto l38
														}
														t161 := int32(load32(m.memory[uint32(v1+i32(16)):]))
														t162 := int32(load32(m.memory[uint32(v1+i32(20)):]))
														m.fn845(v4+i32(32), t161, t162, i32(1073159), i32(67), i32(1073226), i32(2))
														t163 := int32(load32(m.memory[int64(uint32(v4))+32:]))
														v1 = t163
														if v1 == 0 {
															goto l38
														}
														t164 := int32(load32(m.memory[int64(uint32(v4))+36:]))
														m.fn1357(v4+i32(208), v2, v1, t164)
														t165 := int32(load32(m.memory[int64(uint32(v4))+228:]))
														v18 = t165
														t166 := int32(load32(m.memory[int64(uint32(v4))+224:]))
														v21 = t166
														t167 := int32(load32(m.memory[int64(uint32(v4))+212:]))
														v19 = t167
														{
															t168 := int32(load32(m.memory[int64(uint32(v4))+208:]))
															v1 = t168
															if v1 == i32(-1) {
																if v19 == i32(-1) {
																	goto l38
																}
																t170 := int32(load32(m.memory[int64(uint32(v4))+216:]))
																v5 = t170
																store32(m.memory[int64(uint32(v4))+572:], uint32(v18))
																store32(m.memory[int64(uint32(v4))+568:], uint32(v21))
																m.fn1053(v4+i32(480), v21+i32(8), v18)
																t171 := int32(load32(m.memory[int64(uint32(v4))+480:]))
																if t171 != i32(-1) {
																	t220 := int32(load32(m.memory[int64(uint32(v4))+508:]))
																	t221 := int32(load32(m.memory[int64(uint32(v4))+512:]))
																	m.fn1362(v4+i32(208), t220, t221)
																	m.fn1271(v3, v4+i32(208))
																	m.fn1042(v4 + i32(480))
																	goto l43
																}
																{
																	t172 := int32(load32(m.memory[int64(uint32(v4))+484:]))
																	if t172 != i32(-0x7ffffffd) {
																		m.fn785(v10)
																		goto l43
																	}
																	v1 = i32(-0x7ffffffd)
																	t173 := int32(load32(m.memory[int64(uint32(v4))+504:]))
																	v23 = t173
																	t174 := int32(load32(m.memory[int64(uint32(v4))+500:]))
																	v24 = t174
																	t175 := int64(load64(m.memory[int64(uint32(v4))+492:]))
																	v25 = t175
																	t176 := int32(load32(m.memory[int64(uint32(v4))+488:]))
																	v26 = t176
																	goto l42
																}
															}
															t169 := int64(load64(m.memory[int64(uint32(v4))+216:]))
															v25 = t169
															v23 = v18
															v24 = v21
															v26 = v19
															goto l29
														}
													}
													t121 := v4 + i32(72)
													v5 = v1 + i32(16)
													t122 := int32(load32(m.memory[uint32(v5):]))
													v1 = v1 + i32(20)
													t123 := int32(load32(m.memory[uint32(v1):]))
													m.fn1046(t121, t122, t123, i32(1074346), i32(58), i32(1083180), i32(6))
													t124 := int32(load32(m.memory[int64(uint32(v4))+76:]))
													t125 := int32(load32(m.memory[int64(uint32(v4))+72:]))
													t126 := v4
													v17 = t125
													p127 := i32(6)
													if v17 != 0 {
														p127 = t124
													}
													store32(m.memory[int64(uint32(t126))+188:], uint32(p127))
													t129 := v4
													p128 := i32(1077452)
													if v17 != 0 {
														p128 = v17
													}
													store32(m.memory[int64(uint32(t129))+184:], uint32(p128))
													t130 := int32(load32(m.memory[uint32(v5):]))
													t131 := int32(load32(m.memory[uint32(v1):]))
													m.fn1046(v4+i32(64), t130, t131, i32(1074346), i32(58), i32(1073713), i32(4))
													t132 := int32(load32(m.memory[int64(uint32(v4))+64:]))
													t133 := v4 + i32(56)
													v17 = t132
													p134 := i32(1)
													if v17 != 0 {
														p134 = v17
													}
													t135 := int32(load32(m.memory[int64(uint32(v4))+68:]))
													p136 := i32(0)
													if v17 != 0 {
														p136 = t135
													}
													m.fn46(t133, p134, p136)
													{
														t137 := int32(load32(m.memory[int64(uint32(v4))+60:]))
														v17 = t137
														if v17 == 0 {
															goto l31
														}
														t138 := int32(load32(m.memory[int64(uint32(v4))+56:]))
														m.fn51(v4+i32(624), t138, v17)
														goto l32
													}
												l31:
													store32(m.memory[int64(uint32(v4))+212:], uint32(i32(1)))
													store32(m.memory[int64(uint32(v4))+208:], uint32(v4+i32(184)))
													m.fn73(v4+i32(624), i32(1051400), v4+i32(208))
												l32:
													t139 := int32(load32(m.memory[uint32(v5):]))
													t140 := int32(load32(m.memory[uint32(v1):]))
													m.fn845(v4+i32(48), t139, t140, i32(1073159), i32(67), i32(1073226), i32(2))
													t141 := int32(load32(m.memory[int64(uint32(v4))+48:]))
													v1 = t141
													if v1 == 0 {
														goto l33
													}
													t142 := int32(load32(m.memory[int64(uint32(v4))+52:]))
													m.fn1357(v4+i32(208), v2, v1, t142)
													t143 := int32(load32(m.memory[int64(uint32(v4))+228:]))
													v17 = t143
													t144 := int32(load32(m.memory[int64(uint32(v4))+224:]))
													v19 = t144
													t145 := int64(load64(m.memory[int64(uint32(v4))+216:]))
													v20 = t145
													t146 := int32(load32(m.memory[int64(uint32(v4))+212:]))
													v5 = t146
													{
														{
															t147 := int32(load32(m.memory[int64(uint32(v4))+208:]))
															v1 = t147
															if v1 == i32(-1) {
																goto l34
															}
															v23 = v17
															v24 = v19
															v25 = v20
															v26 = v5
															goto l35
														}
													l34:
														if v5 == i32(-1) {
															goto l33
														}
														store64(m.memory[int64(uint32(v4))+652:], uint64(v20))
														store32(m.memory[int64(uint32(v4))+648:], uint32(v5))
														store32(m.memory[int64(uint32(v4))+444:], uint32(v17))
														store32(m.memory[int64(uint32(v4))+440:], uint32(v19))
														m.fn1182(v4+i32(40), v12, i32(1083188))
														t148 := int32(load32(m.memory[int64(uint32(v4))+44:]))
														v5 = t148
														t149 := int32(load32(m.memory[int64(uint32(v4))+40:]))
														v1 = t149
														m.fn51(v4+i32(568), i32(1077458), i32(29))
														m.fn1296(v4+i32(208), v1, v4+i32(568), v4+i32(648), v19+i32(8), v17)
														t150 := int32(load32(m.memory[int64(uint32(v4))+212:]))
														v17 = t150
														t151 := int32(load32(m.memory[int64(uint32(v4))+208:]))
														v1 = t151
														if v1 == i32(-1) {
															t219 := int32(load32(m.memory[uint32(v5):]))
															store32(m.memory[uint32(v5):], uint32(t219+i32(1)))
															store32(m.memory[int64(uint32(v4))+164:], uint32(v17))
															store32(m.memory[int64(uint32(v4))+160:], uint32(i32(-0x80000000)))
															m.fn754(v4 + i32(440))
															goto l49
														}
														t152 := int32(load32(m.memory[int64(uint32(v4))+228:]))
														v23 = t152
														t153 := int32(load32(m.memory[int64(uint32(v4))+224:]))
														v24 = t153
														t154 := int64(load64(m.memory[int64(uint32(v4))+216:]))
														v25 = t154
														t155 := int32(load32(m.memory[uint32(v5):]))
														store32(m.memory[uint32(v5):], uint32(t155+i32(1)))
														m.fn754(v4 + i32(440))
														v26 = v17
													}
												l35:
													t156 := int32(load32(m.memory[int64(uint32(v4))+624:]))
													t157 := int32(load32(m.memory[int64(uint32(v4))+628:]))
													m.fn16(t156, t157)
													goto l37
												}
												v22 = i32(0)
												{
													t53 := int32(load32(m.memory[uint32(v1+i32(28)):]))
													v5 = t53
													t54 := int32(load32(m.memory[uint32(v1+i32(32)):]))
													t55 := v5
													v17 = t54
													t56 := m.fn886(t55, v17, i32(1074411), i32(53), i32(1083090), i32(5))
													v1 = t56
													if v1 == 0 {
														goto l16
													}
													t57 := int32(load32(m.memory[uint32(v1+i32(16)):]))
													t58 := int32(load32(m.memory[uint32(v1+i32(20)):]))
													m.fn1046(v4+i32(112), t57, t58, i32(1074411), i32(53), i32(1074464), i32(8))
													t59 := int32(load32(m.memory[int64(uint32(v4))+112:]))
													v1 = t59
													if v1 == 0 {
														goto l16
													}
													v22 = i32(1)
													t60 := int32(load32(m.memory[int64(uint32(v4))+116:]))
													t61 := v1
													v19 = t60
													t62 := m.fn15(t61, v19, i32(1073318), i32(1))
													if t62 != 0 {
														goto l16
													}
													t63 := m.fn15(v1, v19, i32(1071691), i32(4))
													v22 = t63
												}
											l16:
												m.fn1165(v4 + i32(568))
												store32(m.memory[int64(uint32(v4))+644:], uint32(i32(2)))
												store32(m.memory[int64(uint32(v4))+640:], uint32(i32(1073488)))
												store32(m.memory[int64(uint32(v4))+636:], uint32(i32(53)))
												store32(m.memory[int64(uint32(v4))+632:], uint32(i32(1074411)))
												store32(m.memory[int64(uint32(v4))+624:], uint32(v5))
												store32(m.memory[int64(uint32(v4))+628:], uint32(v5+v17*i32(44)))
												{
												l20:
													{
														t64 := m.fn1186(v4 + i32(624))
														v1 = t64
														if v1 == 0 {
															memory_copy(m.memory, uint32(v4+i32(208)), uint32(v4+i32(568)), uint32(i32(56)))
															m.fn1168(v4+i32(648), v4+i32(208))
															t67 := int32(load32(m.memory[int64(uint32(v4))+656:]))
															v1 = t67
															if v1 == 0 {
																m.fn972(v4 + i32(648))
																goto l2
															}
															t68 := int32(load32(m.memory[int64(uint32(v4))+652:]))
															t69 := m.fn1234(t68, v1, v22)
															store32(m.memory[int64(uint32(v4))+660:], uint32(t69))
															t70 := int32(load32(m.memory[int64(uint32(v4))+664:]))
															store32(m.memory[int64(uint32(v11))+16:], uint32(t70))
															t71 := int64(load64(m.memory[int64(uint32(v4))+656:]))
															store64(m.memory[int64(uint32(v11))+8:], uint64(t71))
															t72 := int64(load64(m.memory[int64(uint32(v4))+648:]))
															store64(m.memory[uint32(v11):], uint64(t72))
															store32(m.memory[int64(uint32(v4))+208:], uint32(i32(-0x7ffffffe)))
															m.fn338(v3, v4+i32(208))
															goto l2
														}
														m.fn1166(v4 + i32(568))
														t65 := int32(load32(m.memory[int64(uint32(v1))+32:]))
														v5 = t65
														t66 := int32(load32(m.memory[int64(uint32(v1))+28:]))
														v1 = t66
														store32(m.memory[int64(uint32(v4))+668:], uint32(i32(2)))
														store32(m.memory[int64(uint32(v4))+664:], uint32(i32(1083095)))
														store32(m.memory[int64(uint32(v4))+660:], uint32(i32(53)))
														store32(m.memory[int64(uint32(v4))+656:], uint32(i32(1074411)))
														store32(m.memory[int64(uint32(v4))+648:], uint32(v1))
														store32(m.memory[int64(uint32(v4))+652:], uint32(v1+v5*i32(44)))
														goto l24
													}
												l24:
													{
														t73 := m.fn1186(v4 + i32(648))
														v1 = t73
														if v1 == 0 {
															goto l20
														}
														t74 := int32(load32(m.memory[uint32(v1+i32(16)):]))
														t75 := v4 + i32(104)
														v5 = t74
														t76 := int32(load32(m.memory[uint32(v1+i32(20)):]))
														t77 := v5
														v17 = t76
														m.fn1046(t75, t77, v17, i32(1074411), i32(53), i32(1074893), i32(6))
														{
															{
																t78 := int32(load32(m.memory[int64(uint32(v4))+104:]))
																v19 = t78
																if v19 == 0 {
																	goto l21
																}
																t79 := int32(load32(m.memory[int64(uint32(v4))+108:]))
																t80 := v19
																v18 = t79
																t81 := m.fn15(t80, v18, i32(1073318), i32(1))
																if t81 != 0 {
																	goto l22
																}
																t82 := m.fn15(v19, v18, i32(1071691), i32(4))
																if t82 != 0 {
																	goto l22
																}
															}
														l21:
															m.fn1046(v4+i32(96), v5, v17, i32(1074411), i32(53), i32(1074899), i32(6))
															t83 := int32(load32(m.memory[int64(uint32(v4))+96:]))
															v19 = t83
															if v19 == 0 {
																goto l23
															}
															t84 := int32(load32(m.memory[int64(uint32(v4))+100:]))
															t85 := v19
															v18 = t84
															t86 := m.fn15(t85, v18, i32(1073318), i32(1))
															if t86 != 0 {
																goto l22
															}
															t87 := m.fn15(v19, v18, i32(1071691), i32(4))
															if t87 == 0 {
																goto l23
															}
														}
													l22:
														_ = m.fn1260(v4 + i32(568))
														goto l24
													l23:
														m.fn1046(v4+i32(88), v5, v17, i32(1074411), i32(53), i32(1074885), i32(8))
														v19 = i32(1)
														v18 = i32(1)
														{
															t89 := int32(load32(m.memory[int64(uint32(v4))+88:]))
															v21 = t89
															if v21 == 0 {
																goto l25
															}
															t90 := int32(load32(m.memory[int64(uint32(v4))+92:]))
															m.fn1071(v4+i32(208), v21, t90)
															t91 := int32(load32(m.memory[int64(uint32(v4))+212:]))
															v18 = t91
															p92 := i32(1)
															if uint32(v18) > uint32(i32(1)) {
																p92 = v18
															}
															t93 := int32(m.memory[int64(uint32(v4))+208])
															p94 := p92
															if t93 != 0 {
																p94 = i32(1)
															}
															v18 = p94
														}
													l25:
														m.fn1046(v4+i32(80), v5, v17, i32(1074411), i32(53), i32(1083097), i32(7))
														{
															t95 := int32(load32(m.memory[int64(uint32(v4))+80:]))
															v5 = t95
															if v5 == 0 {
																goto l26
															}
															t96 := int32(load32(m.memory[int64(uint32(v4))+84:]))
															m.fn1071(v4+i32(208), v5, t96)
															t97 := int32(load32(m.memory[int64(uint32(v4))+212:]))
															v5 = t97
															p98 := i32(1)
															if uint32(v5) > uint32(i32(1)) {
																p98 = v5
															}
															t99 := int32(m.memory[int64(uint32(v4))+208])
															p100 := p98
															if t99 != 0 {
																p100 = i32(1)
															}
															v19 = p100
														}
													l26:
														store32(m.memory[int64(uint32(v4))+192:], uint32(i32(0)))
														store64(m.memory[int64(uint32(v4))+184:], uint64(i64(0x800000000)))
														{
															t101 := int32(load32(m.memory[uint32(v1+i32(28)):]))
															t102 := int32(load32(m.memory[uint32(v1+i32(32)):]))
															t103 := m.fn886(t101, t102, i32(1074411), i32(53), i32(1083084), i32(6))
															v1 = t103
															if v1 == 0 {
																goto l27
															}
															t104 := int32(load32(m.memory[uint32(v1+i32(28)):]))
															t105 := int32(load32(m.memory[uint32(v1+i32(32)):]))
															m.fn1284(v4+i32(208), t104, t105, v2, i32(0), v4+i32(184))
															t106 := int32(load32(m.memory[int64(uint32(v4))+208:]))
															v1 = t106
															if v1 == i32(-1) {
																goto l27
															}
															t107 := int32(load32(m.memory[int64(uint32(v4))+228:]))
															v23 = t107
															t108 := int32(load32(m.memory[int64(uint32(v4))+224:]))
															v24 = t108
															t109 := int64(load64(m.memory[int64(uint32(v4))+216:]))
															v25 = t109
															t110 := int32(load32(m.memory[int64(uint32(v4))+212:]))
															v26 = t110
															m.fn969(v4 + i32(184))
															goto l28
														}
													l27:
														t111 := int32(load32(m.memory[int64(uint32(v4))+192:]))
														store32(m.memory[int64(uint32(v4))+448:], uint32(t111))
														t112 := int64(load64(m.memory[int64(uint32(v4))+184:]))
														store64(m.memory[int64(uint32(v4))+440:], uint64(t112))
														store32(m.memory[int64(uint32(v4))+456:], uint32(v19))
														store32(m.memory[int64(uint32(v4))+452:], uint32(v18))
														m.fn1167(v4+i32(208), v4+i32(568), v4+i32(440))
														t113 := int32(load32(m.memory[int64(uint32(v4))+208:]))
														v1 = t113
														if v1 == i32(-1) {
															goto l24
														}
													}
													t114 := int32(load32(m.memory[int64(uint32(v4))+228:]))
													v23 = t114
													t115 := int32(load32(m.memory[int64(uint32(v4))+224:]))
													v24 = t115
													t116 := int64(load64(m.memory[int64(uint32(v4))+216:]))
													v25 = t116
													t117 := int32(load32(m.memory[int64(uint32(v4))+212:]))
													v26 = t117
												}
											l28:
												m.fn1259(v4 + i32(568))
												goto l29
											}
											t46 := m.fn15(v5, v17, i32(1083126), i32(3))
											if t46 == 0 {
												goto l2
											}
											v17 = v1 + i32(28)
											t47 := int32(load32(m.memory[uint32(v17):]))
											v18 = v1 + i32(32)
											t48 := int32(load32(m.memory[uint32(v18):]))
											t49 := m.fn1097(t47, t48, i32(1074346), i32(58), i32(1083129), i32(5))
											v1 = t49
											if v1 != 0 {
												goto l13
											}
											v1 = i32(0)
											goto l14
										}
									l13:
										t177 := int32(load32(m.memory[uint32(v1+i32(16)):]))
										t178 := int32(load32(m.memory[uint32(v1+i32(20)):]))
										m.fn1046(v4+i32(144), t177, t178, i32(1074346), i32(58), i32(1073571), i32(5))
										t179 := int32(load32(m.memory[int64(uint32(v4))+148:]))
										v5 = t179
										t180 := int32(load32(m.memory[int64(uint32(v4))+144:]))
										v1 = t180
									}
								l14:
									m.fn1358(v4+i32(208), v1, v5)
									t181 := int32(load32(m.memory[int64(uint32(v4))+208:]))
									v1 = t181
									t182 := v1
									var p183 int32
									if v1 == i32(-1) {
										p183 = 1
									}
									v1 = p183
									p184 := t182
									if v1 != 0 {
										p184 = i32(0)
									}
									v21 = p184
									t185 := int32(load32(m.memory[int64(uint32(v4))+212:]))
									p186 := t185
									if v1 != 0 {
										p186 = i32(1)
									}
									v5 = p186
									t187 := int32(load32(m.memory[int64(uint32(v4))+216:]))
									p188 := t187
									if v1 != 0 {
										p188 = i32(0)
									}
									v19 = p188
									{
										{
											t189 := int32(load32(m.memory[uint32(v17):]))
											t190 := int32(load32(m.memory[uint32(v18):]))
											t191 := m.fn1097(t189, t190, i32(1074411), i32(53), i32(1077487), i32(4))
											v1 = t191
											if v1 == 0 {
												goto l44
											}
											t192 := int32(load32(m.memory[uint32(v1+i32(16)):]))
											t193 := v4 + i32(136)
											v17 = t192
											t194 := int32(load32(m.memory[uint32(v1+i32(20)):]))
											t195 := v17
											v18 = t194
											m.fn845(t193, t195, v18, i32(1073159), i32(67), i32(1073614), i32(5))
											{
												{
													t196 := int32(load32(m.memory[int64(uint32(v4))+136:]))
													v1 = t196
													if v1 == 0 {
														goto l45
													}
													t197 := int32(load32(m.memory[int64(uint32(v4))+140:]))
													v17 = t197
													goto l46
												}
											l45:
												m.fn845(v4+i32(128), v17, v18, i32(1073159), i32(67), i32(1073228), i32(4))
												t198 := int32(load32(m.memory[int64(uint32(v4))+128:]))
												v1 = t198
												if v1 == 0 {
													goto l44
												}
												t199 := int32(load32(m.memory[int64(uint32(v4))+132:]))
												v17 = t199
											}
										l46:
											m.fn1359(v4+i32(208), v16, v15, v14, v13, v12, v1, v17)
											t200 := int32(load32(m.memory[int64(uint32(v4))+220:]))
											v27 = t200
											t201 := int32(load32(m.memory[int64(uint32(v4))+216:]))
											v28 = t201
											t202 := int32(load32(m.memory[int64(uint32(v4))+212:]))
											v17 = t202
											t203 := int32(load32(m.memory[int64(uint32(v4))+208:]))
											v1 = t203
											if v1 != i32(-1) {
												t222 := int64(load64(m.memory[int64(uint32(v4))+224:]))
												v25 = t222
												store32(m.memory[int64(uint32(v0))+12:], uint32(v27))
												store32(m.memory[int64(uint32(v0))+8:], uint32(v28))
												store32(m.memory[int64(uint32(v0))+4:], uint32(v17))
												store64(m.memory[int64(uint32(v0))+16:], uint64(v25))
												store32(m.memory[uint32(v0):], uint32(v1))
												m.fn16(v21, v5)
												goto l7
											}
											if v17 != i32(-1) {
												goto l48
											}
										}
									l44:
										m.fn46(v4+i32(120), v5, v19)
										v17 = i32(-1)
										t204 := int32(load32(m.memory[int64(uint32(v4))+124:]))
										if t204 != 0 {
											goto l48
										}
										m.fn16(v21, v5)
										goto l2
									}
								l48:
									t205 := m.fn113(i32(4), i32(28))
									v1 = t205
									store32(m.memory[int64(uint32(v4))+576:], uint32(v27))
									store32(m.memory[int64(uint32(v4))+572:], uint32(v28))
									store32(m.memory[int64(uint32(v4))+568:], uint32(v17))
									store32(m.memory[int64(uint32(v4))+208:], uint32(i32(-0x7fffffff)))
									m.fn1360(v4+i32(648), v4+i32(568), v4+i32(208))
									store32(m.memory[int64(uint32(v1))+12:], uint32(v19))
									store32(m.memory[int64(uint32(v1))+8:], uint32(v5))
									store32(m.memory[int64(uint32(v1))+4:], uint32(v21))
									store32(m.memory[uint32(v1):], uint32(i32(5)))
									t206 := int64(load64(m.memory[int64(uint32(v4))+648:]))
									store64(m.memory[int64(uint32(v1))+16:], uint64(t206))
									t207 := int32(load32(m.memory[int64(uint32(v4))+656:]))
									store32(m.memory[int64(uint32(v1))+24:], uint32(t207))
									store32(m.memory[int64(uint32(v4))+220:], uint32(i32(1)))
									store32(m.memory[int64(uint32(v4))+216:], uint32(v1))
									store64(m.memory[int64(uint32(v4))+208:], uint64(i64(0x180000000)))
									m.fn338(v3, v4+i32(208))
									goto l2
								}
							}
						l3:
							v5 = v1 + i32(28)
							t19 := int32(load32(m.memory[uint32(v5):]))
							v17 = v1 + i32(32)
							t20 := int32(load32(m.memory[uint32(v17):]))
							t21 := m.fn1097(t19, t20, i32(1074346), i32(58), i32(1083073), i32(2))
							v1 = t21
							if v1 != 0 {
								goto l5
							}
							v18 = i32(-1)
							store32(m.memory[int64(uint32(v4))+184:], uint32(i32(-1)))
							goto l6
						}
					l5:
						t22 := v4 + i32(16)
						v19 = v1 + i32(16)
						t23 := int32(load32(m.memory[uint32(v19):]))
						v18 = v1 + i32(20)
						t24 := int32(load32(m.memory[uint32(v18):]))
						m.fn1046(t22, t23, t24, i32(1074346), i32(58), i32(1074404), i32(4))
						t25 := int32(load32(m.memory[int64(uint32(v4))+16:]))
						t26 := v4 + i32(208)
						v1 = t25
						p27 := i32(1073232)
						if v1 != 0 {
							p27 = v1
						}
						t28 := int32(load32(m.memory[int64(uint32(v4))+20:]))
						p29 := i32(4)
						if v1 != 0 {
							p29 = t28
						}
						m.fn51(t26, p27, p29)
						t30 := int32(load32(m.memory[uint32(v19):]))
						t31 := int32(load32(m.memory[uint32(v18):]))
						m.fn1046(v4+i32(8), t30, t31, i32(1074346), i32(58), i32(1074408), i32(3))
						t32 := int32(load32(m.memory[int64(uint32(v4))+8:]))
						t33 := int32(load32(m.memory[int64(uint32(v4))+12:]))
						m.fn1041(v8, t32, t33)
						t34 := int64(load64(m.memory[int64(uint32(v4))+224:]))
						store64(m.memory[int64(uint32(v4))+200:], uint64(t34))
						t35 := int64(load64(m.memory[int64(uint32(v4))+216:]))
						store64(m.memory[int64(uint32(v4))+192:], uint64(t35))
						t36 := int64(load64(m.memory[int64(uint32(v4))+208:]))
						t37 := v4
						v20 = t36
						store64(m.memory[int64(uint32(t37))+184:], uint64(v20))
						v18 = i32(-1)
						v1 = int32(v20)
						if v1 != i32(-1) {
							v21 = i32(-1)
							t38 := int32(load32(m.memory[int64(uint32(v4))+188:]))
							v19 = t38
							t39 := int32(load32(m.memory[int64(uint32(v4))+192:]))
							t40 := v19
							v18 = t39
							t41 := m.fn15(t40, v18, i32(1083075), i32(6))
							if t41 != 0 {
								goto l10
							}
							t42 := m.fn15(v19, v18, i32(1080983), i32(2))
							if t42 != 0 {
								goto l10
							}
							t43 := m.fn15(v19, v18, i32(1083081), i32(3))
							if t43 != 0 {
								goto l10
							}
							v19 = i32(0)
							v18 = v1
							goto l9
						}
					}
				l6:
					v19 = i32(1)
					goto l9
				l43:
					v1 = i32(-1)
				l42:
					m.fn754(v4 + i32(568))
					m.fn16(v19, v5)
					goto l37
				l38:
					t223 := int32(load32(m.memory[uint32(v5):]))
					t224 := int32(load32(m.memory[uint32(v17):]))
					t225 := m.fn1097(t223, t224, i32(1073986), i32(56), i32(1077429), i32(6))
					v1 = t225
					if v1 == 0 {
						goto l2
					}
					t226 := int32(load32(m.memory[uint32(v1+i32(16)):]))
					t227 := int32(load32(m.memory[uint32(v1+i32(20)):]))
					m.fn845(v4+i32(24), t226, t227, i32(1073159), i32(67), i32(1077435), i32(2))
					t228 := int32(load32(m.memory[int64(uint32(v4))+24:]))
					v1 = t228
					if v1 == 0 {
						goto l2
					}
					t229 := int32(load32(m.memory[int64(uint32(v4))+28:]))
					m.fn1357(v4+i32(208), v2, v1, t229)
					t230 := int32(load32(m.memory[int64(uint32(v4))+228:]))
					v17 = t230
					t231 := int32(load32(m.memory[int64(uint32(v4))+224:]))
					v19 = t231
					t232 := int32(load32(m.memory[int64(uint32(v4))+212:]))
					v5 = t232
					{
						t233 := int32(load32(m.memory[int64(uint32(v4))+208:]))
						v1 = t233
						if v1 == i32(-1) {
							if v5 == i32(-1) {
								goto l2
							}
							t235 := int32(load32(m.memory[int64(uint32(v4))+216:]))
							v1 = t235
							store32(m.memory[int64(uint32(v4))+572:], uint32(v17))
							store32(m.memory[int64(uint32(v4))+568:], uint32(v19))
							m.fn1053(v4+i32(524), v19+i32(8), v17)
							{
								t236 := int32(load32(m.memory[int64(uint32(v4))+524:]))
								if t236 != i32(-1) {
									t242 := int32(load32(m.memory[int64(uint32(v4))+552:]))
									t243 := int32(load32(m.memory[int64(uint32(v4))+556:]))
									m.fn1363(v4+i32(208), t242, t243)
									m.fn1271(v3, v4+i32(208))
									m.fn1042(v4 + i32(524))
									goto l53
								}
								t237 := int32(load32(m.memory[int64(uint32(v4))+528:]))
								if t237 != i32(-0x7ffffffd) {
									goto l52
								}
								t238 := int32(load32(m.memory[int64(uint32(v4))+548:]))
								v23 = t238
								t239 := int32(load32(m.memory[int64(uint32(v4))+544:]))
								v24 = t239
								t240 := int64(load64(m.memory[int64(uint32(v4))+536:]))
								v25 = t240
								t241 := int32(load32(m.memory[int64(uint32(v4))+532:]))
								v26 = t241
								m.fn754(v4 + i32(568))
								m.fn16(v5, v1)
								v1 = i32(-0x7ffffffd)
								goto l29
							}
						l52:
							m.fn785(v9)
						l53:
							m.fn754(v4 + i32(568))
							m.fn16(v5, v1)
							goto l2
						}
						t234 := int64(load64(m.memory[int64(uint32(v4))+216:]))
						v25 = t234
						v23 = v17
						v24 = v19
						v26 = v5
						goto l29
					}
				}
			l33:
				store32(m.memory[int64(uint32(v4))+160:], uint32(i32(-1)))
			l49:
				t244 := m.fn113(i32(4), i32(28))
				v1 = t244
				store32(m.memory[int64(uint32(v4))+208:], uint32(i32(-0x7fffffff)))
				m.fn1360(v4+i32(568), v4+i32(160), v4+i32(208))
				store32(m.memory[uint32(v1):], uint32(i32(5)))
				t245 := int64(load64(m.memory[int64(uint32(v4))+624:]))
				store64(m.memory[int64(uint32(v1))+4:], uint64(t245))
				t246 := int32(load32(m.memory[int64(uint32(v4))+632:]))
				store32(m.memory[int64(uint32(v1))+12:], uint32(t246))
				t247 := int64(load64(m.memory[int64(uint32(v4))+568:]))
				store64(m.memory[int64(uint32(v1))+16:], uint64(t247))
				t248 := int32(load32(m.memory[int64(uint32(v4))+576:]))
				store32(m.memory[int64(uint32(v1))+24:], uint32(t248))
				store32(m.memory[int64(uint32(v4))+220:], uint32(i32(1)))
				store32(m.memory[int64(uint32(v4))+216:], uint32(v1))
				store64(m.memory[int64(uint32(v4))+208:], uint64(i64(0x180000000)))
				m.fn338(v3, v4+i32(208))
				v1 = i32(-1)
			}
		l37:
			if v1 == i32(-1) {
				goto l2
			}
		l29:
			store32(m.memory[int64(uint32(v0))+20:], uint32(v23))
			store32(m.memory[int64(uint32(v0))+16:], uint32(v24))
			store64(m.memory[int64(uint32(v0))+8:], uint64(v25))
			store32(m.memory[int64(uint32(v0))+4:], uint32(v26))
			store32(m.memory[uint32(v0):], uint32(v1))
			goto l7
		l10:
			v18 = v1
			goto l54
		l9:
			v21 = i32(-1)
			{
				t249 := int32(load32(m.memory[uint32(v5):]))
				t250 := int32(load32(m.memory[uint32(v17):]))
				t251 := m.fn886(t249, t250, i32(1074346), i32(58), i32(1083084), i32(6))
				v1 = t251
				if v1 == 0 {
					goto l55
				}
				{
					{
						if v19 == 0 {
							goto l56
						}
						v5 = i32(0)
						goto l57
					l56:
						t252 := int32(load32(m.memory[int64(uint32(v4))+188:]))
						t253 := int32(load32(m.memory[int64(uint32(v4))+192:]))
						t254 := m.fn1364(t252, t253)
						if t254 == 0 {
							t261 := int32(load32(m.memory[uint32(v1+i32(28)):]))
							t262 := v4 + i32(208)
							v5 = t261
							t263 := int32(load32(m.memory[uint32(v1+i32(32)):]))
							t264 := v5
							v1 = t263
							t265 := m.fn886(t264, v1, i32(1074411), i32(53), i32(1074905), i32(8))
							m.fn1276(t262, t265)
							store32(m.memory[int64(uint32(v4))+436:], uint32(i32(0)))
							store64(m.memory[int64(uint32(v4))+428:], uint64(i64(0x400000000)))
							store32(m.memory[int64(uint32(v4))+460:], uint32(i32(1)))
							store32(m.memory[int64(uint32(v4))+456:], uint32(i32(1077161)))
							store32(m.memory[int64(uint32(v4))+452:], uint32(i32(53)))
							store32(m.memory[int64(uint32(v4))+448:], uint32(i32(1074411)))
							store32(m.memory[int64(uint32(v4))+444:], uint32(v5+v1*i32(44)))
							store32(m.memory[int64(uint32(v4))+440:], uint32(v5))
						l67:
							{
								{
									t266 := m.fn1186(v4 + i32(440))
									v1 = t266
									if v1 == 0 {
										t282 := int32(load32(m.memory[int64(uint32(v4))+432:]))
										v1 = t282
										t283 := int32(load32(m.memory[int64(uint32(v4))+436:]))
										t284 := v1
										v5 = t283
										t285 := m.fn23(t284, v5)
										if t285 != 0 {
											m.fn894(v4 + i32(428))
											goto l55
										}
										m.fn45(v7, v1, v5)
										m.memory[int64(uint32(v4))+592] = byte(i32(2))
										t286 := int32(load32(m.memory[int64(uint32(v4))+436:]))
										store32(m.memory[int64(uint32(v4))+576:], uint32(t286))
										t287 := int64(load64(m.memory[int64(uint32(v4))+428:]))
										store64(m.memory[int64(uint32(v4))+568:], uint64(t287))
										m.fn338(v3, v4+i32(568))
										goto l55
									}
									v5 = v1 + i32(28)
									t267 := int32(load32(m.memory[uint32(v5):]))
									v17 = v1 + i32(32)
									t268 := int32(load32(m.memory[uint32(v17):]))
									t269 := m.fn886(t267, t268, i32(1074411), i32(53), i32(1073735), i32(3))
									v1 = t269
									if v1 == 0 {
										goto l60
									}
									t270 := int32(load32(m.memory[uint32(v1+i32(16)):]))
									t271 := int32(load32(m.memory[uint32(v1+i32(20)):]))
									m.fn1046(v4, t270, t271, i32(1074411), i32(53), i32(1072633), i32(3))
									{
										{
											t272 := int32(load32(m.memory[uint32(v4):]))
											v19 = t272
											if v19 != 0 {
												goto l61
											}
											v19 = i32(0)
											goto l62
										}
									l61:
										t273 := int32(load32(m.memory[int64(uint32(v4))+4:]))
										m.fn197(v4+i32(568), v19, t273)
										t274 := int32(load32(m.memory[int64(uint32(v4))+572:]))
										t275 := int32(m.memory[int64(uint32(v4))+568])
										p276 := t274
										if t275 != 0 {
											p276 = i32(0)
										}
										v19 = p276
									}
								l62:
									m.fn1365(v4+i32(624), v2, v4+i32(184), v4+i32(208), v19)
									t277 := int32(load32(m.memory[uint32(v1+i32(28)):]))
									t278 := int32(load32(m.memory[uint32(v1+i32(32)):]))
									m.fn1356(v4+i32(568), t277, t278)
									m.fn1366(v4+i32(648), v4+i32(624), v4+i32(568))
									t279 := int64(load64(m.memory[int64(uint32(v4))+648:]))
									store64(m.memory[int64(uint32(v4))+624:], uint64(t279))
									t280 := int64(load64(m.memory[int64(uint32(v4))+656:]))
									store64(m.memory[int64(uint32(v4))+632:], uint64(t280))
									t281 := int64(load64(m.memory[int64(uint32(v4))+664:]))
									store64(m.memory[int64(uint32(v4))+640:], uint64(t281))
									goto l63
								}
							l60:
								m.fn1365(v4+i32(624), v2, v4+i32(184), v4+i32(208), i32(0))
							l63:
								t288 := int32(load32(m.memory[uint32(v5):]))
								t289 := int32(load32(m.memory[uint32(v17):]))
								t290 := int32(load32(m.memory[int64(uint32(v4))+640:]))
								t291 := m.fn1188(t290)
								t292 := v4 + i32(468)
								t293 := v2
								v1 = t291
								m.fn1367(t292, t288, t289, t293, v1)
								t294 := int32(load32(m.memory[int64(uint32(v4))+472:]))
								v5 = t294
								t295 := int32(load32(m.memory[int64(uint32(v4))+476:]))
								t296 := v5
								v17 = t295
								t297 := m.fn23(t296, v17)
								if t297 != 0 {
									m.fn894(v4 + i32(468))
									goto l67
								}
								m.fn1368(v5, v17, v1)
								{
									t298 := int32(load32(m.memory[int64(uint32(v4))+436:]))
									if t298 == 0 {
										goto l66
									}
									store32(m.memory[int64(uint32(v4))+568:], uint32(i32(8)))
									m.fn1340(v4+i32(428), v4+i32(568))
								}
							l66:
								m.fn1341(v4+i32(428), v4+i32(468))
								goto l67
							}
						}
						v5 = v4 + i32(184)
					}
				l57:
					t255 := int32(load32(m.memory[uint32(v1+i32(28)):]))
					t256 := int32(load32(m.memory[uint32(v1+i32(32)):]))
					m.fn1284(v4+i32(208), t255, t256, v2, v5, v3)
					t257 := int32(load32(m.memory[int64(uint32(v4))+208:]))
					v1 = t257
					if v1 == i32(-1) {
						goto l55
					}
					t258 := int32(load32(m.memory[int64(uint32(v6))+16:]))
					store32(m.memory[int64(uint32(v4))+176:], uint32(t258))
					t259 := int64(load64(m.memory[int64(uint32(v6))+8:]))
					store64(m.memory[int64(uint32(v4))+168:], uint64(t259))
					t260 := int64(load64(m.memory[uint32(v6):]))
					store64(m.memory[int64(uint32(v4))+160:], uint64(t260))
					v21 = v1
					goto l55
				}
			}
		l55:
			if v18 == i32(-1) {
				goto l68
			}
		l54:
			t299 := int32(load32(m.memory[int64(uint32(v4))+188:]))
			m.fn16(v18, t299)
			t300 := int32(load32(m.memory[int64(uint32(v4))+196:]))
			t301 := int32(load32(m.memory[int64(uint32(v4))+200:]))
			m.fn134(t300, t301)
		}
	l68:
		if v21 == i32(-1) {
			goto l2
		}
		t302 := int32(load32(m.memory[int64(uint32(v4))+176:]))
		store32(m.memory[int64(uint32(v0))+20:], uint32(t302))
		t303 := int64(load64(m.memory[int64(uint32(v4))+168:]))
		store64(m.memory[int64(uint32(v0))+12:], uint64(t303))
		t304 := int64(load64(m.memory[int64(uint32(v4))+160:]))
		store64(m.memory[int64(uint32(v0))+4:], uint64(t304))
		store32(m.memory[uint32(v0):], uint32(v21))
	}
l7:
	m.g0 = v4 + i32(672)
}
func (m *Module) fn1284(v0, v1, v2, v3, v4, v5 int32) {
	var v6, v7 int32
	var v8 int64
	var v9, v10, v11, v12, v13 int32
	var v14 int64
	t0 := m.g0
	v6 = t0 - i32(480)
	m.g0 = v6
	t1 := m.fn886(v1, v2, i32(1074411), i32(53), i32(1074905), i32(8))
	m.fn1276(v6+i32(8), t1)
	t2 := int32(load32(m.memory[int64(uint32(v3))+24:]))
	v7 = t2
	t3 := int64(load64(m.memory[uint32(v7):]))
	t4 := v7
	v8 = t3 + i64(1)
	store64(m.memory[uint32(t4):], uint64(v8))
	memory_zero(m.memory, uint32(v6+i32(224)), uint32(i32(72)))
	m.memory[int64(uint32(v6))+304] = byte(i32(0))
	store64(m.memory[int64(uint32(v6))+296:], uint64(i64(0)))
	store32(m.memory[int64(uint32(v6))+316:], uint32(i32(0)))
	store64(m.memory[int64(uint32(v6))+308:], uint64(i64(0x800000000)))
	store32(m.memory[int64(uint32(v6))+340:], uint32(i32(1)))
	store32(m.memory[int64(uint32(v6))+336:], uint32(i32(1077161)))
	store32(m.memory[int64(uint32(v6))+332:], uint32(i32(53)))
	store32(m.memory[int64(uint32(v6))+328:], uint32(i32(1074411)))
	store32(m.memory[int64(uint32(v6))+324:], uint32(v1+v2*i32(44)))
	store32(m.memory[int64(uint32(v6))+320:], uint32(v1))
	v9 = v6 + i32(416) | i32(4)
	v10 = v6 + i32(416) + i32(28)
	v11 = v6 + i32(305)
l7:
	{
		t5 := m.fn1186(v6 + i32(320))
		v1 = t5
		if v1 == 0 {
			goto l0
		}
		v2 = i32(0)
		{
			v7 = v1 + i32(28)
			t6 := int32(load32(m.memory[uint32(v7):]))
			v12 = v1 + i32(32)
			t7 := int32(load32(m.memory[uint32(v12):]))
			t8 := m.fn886(t6, t7, i32(1074411), i32(53), i32(1073735), i32(3))
			v1 = t8
			if v1 == 0 {
				goto l1
			}
			t9 := int32(load32(m.memory[uint32(v1+i32(16)):]))
			t10 := int32(load32(m.memory[uint32(v1+i32(20)):]))
			m.fn1046(v6, t9, t10, i32(1074411), i32(53), i32(1072633), i32(3))
			t11 := int32(load32(m.memory[uint32(v6):]))
			v13 = t11
			if v13 == 0 {
				goto l1
			}
			t12 := int32(load32(m.memory[int64(uint32(v6))+4:]))
			m.fn197(v6+i32(416), v13, t12)
			t13 := int32(load32(m.memory[int64(uint32(v6))+420:]))
			t14 := int32(m.memory[int64(uint32(v6))+416])
			p15 := t13
			if t14 != 0 {
				p15 = i32(0)
			}
			v2 = p15
		}
	l1:
		m.fn1365(v6+i32(344), v3, v4, v6+i32(8), v2)
		{
			if v1 == 0 {
				goto l2
			}
			t16 := int32(load32(m.memory[uint32(v1+i32(28)):]))
			t17 := int32(load32(m.memory[uint32(v1+i32(32)):]))
			m.fn1356(v6+i32(416), t16, t17)
			m.fn1366(v6+i32(368), v6+i32(344), v6+i32(416))
			t18 := int64(load64(m.memory[int64(uint32(v6))+368:]))
			store64(m.memory[int64(uint32(v6))+344:], uint64(t18))
			t19 := int64(load64(m.memory[int64(uint32(v6))+376:]))
			store64(m.memory[int64(uint32(v6))+352:], uint64(t19))
			t20 := int64(load64(m.memory[int64(uint32(v6))+384:]))
			store64(m.memory[int64(uint32(v6))+360:], uint64(t20))
		}
	l2:
		t21 := int32(load32(m.memory[uint32(v7):]))
		t22 := int32(load32(m.memory[uint32(v12):]))
		t23 := int32(load32(m.memory[int64(uint32(v6))+360:]))
		t24 := m.fn1188(t23)
		m.fn1367(v6+i32(392), t21, t22, v3, t24)
		{
			t25 := int32(load32(m.memory[int64(uint32(v6))+396:]))
			t26 := int32(load32(m.memory[int64(uint32(v6))+400:]))
			t27 := m.fn23(t25, t26)
			if t27 != 0 {
				m.fn1351(v5, v6+i32(308))
				m.fn894(v6 + i32(392))
				goto l7
			}
			t28 := int32(m.memory[int64(uint32(v6))+344])
			switch t28 {
			case 2:
				t31 := m.fn113(i32(8), i32(32))
				v1 = t31
				store32(m.memory[uint32(v1):], uint32(i32(-0x80000000)))
				t32 := int64(load64(m.memory[int64(uint32(v6))+392:]))
				store64(m.memory[int64(uint32(v1))+4:], uint64(t32))
				t33 := int32(load32(m.memory[int64(uint32(v6))+400:]))
				store32(m.memory[int64(uint32(v1))+12:], uint32(t33))
				store32(m.memory[int64(uint32(v6))+440:], uint32(v2))
				store32(m.memory[int64(uint32(v6))+464:], uint32(i32(1)))
				store32(m.memory[int64(uint32(v6))+460:], uint32(v1))
				store32(m.memory[int64(uint32(v6))+456:], uint32(i32(1)))
				store32(m.memory[int64(uint32(v6))+444:], uint32(i32(-1)))
				store64(m.memory[int64(uint32(v6))+432:], uint64(i64(0)))
				m.memory[int64(uint32(v6))+424] = byte(i32(0))
				store64(m.memory[int64(uint32(v6))+416:], uint64(v8))
				m.fn1369(v6+i32(308), v6+i32(416))
				goto l7
			case 3:
				{
					{
						t35 := v6 + i32(296)
						p34 := i32(8)
						if uint32(v2) < uint32(i32(8)) {
							p34 = v2
						}
						v1 = p34
						v7 = t35 + v1
						t36 := int32(m.memory[uint32(v7)])
						if t36 != 0 {
							goto l8
						}
						m.memory[uint32(v7)] = byte(i32(1))
						t37 := int64(load64(m.memory[int64(uint32(v6))+352:]))
						v14 = t37
						goto l9
					}
				l8:
					t38 := int64(load64(m.memory[uint32(v6+i32(224)+v1<<3):]))
					v14 = t38 + i64(1)
					p39 := v14
					if v14 == 0 {
						p39 = i64(-1)
					}
					v14 = p39
				}
			l9:
				t40 := int32(m.memory[int64(uint32(v6))+346])
				v12 = t40
				t41 := int32(m.memory[int64(uint32(v6))+345])
				v7 = t41
				store64(m.memory[uint32(v6+i32(224)+v1<<3):], uint64(v14))
				store32(m.memory[int64(uint32(v6))+424:], uint32(v1+i32(1)))
				store32(m.memory[int64(uint32(v6))+420:], uint32(v11))
				store32(m.memory[int64(uint32(v6))+416:], uint32(v6+i32(296)))
			l11:
				{
					t42 := m.fn1370(v6 + i32(416))
					v1 = t42
					if v1 == 0 {
						switch v12 {
						default:
							store32(m.memory[int64(uint32(v6))+404:], uint32(i32(-1)))
							goto l16
						case 1:
							m.fn804(v6+i32(368), v7, v14)
							store32(m.memory[int64(uint32(v6))+476:], uint32(i32(25)))
							store32(m.memory[int64(uint32(v6))+472:], uint32(v6+i32(368)))
							m.fn73(v6+i32(404), i32(1069643), v6+i32(472))
							t43 := int32(load32(m.memory[int64(uint32(v6))+368:]))
							t44 := int32(load32(m.memory[int64(uint32(v6))+372:]))
							m.fn16(t43, t44)
							goto l16
						case 2:
							m.fn804(v6+i32(368), v7, v14)
							store32(m.memory[int64(uint32(v6))+476:], uint32(i32(25)))
							store32(m.memory[int64(uint32(v6))+472:], uint32(v6+i32(368)))
							m.fn73(v6+i32(404), i32(1069101), v6+i32(472))
							t45 := int32(load32(m.memory[int64(uint32(v6))+368:]))
							t46 := int32(load32(m.memory[int64(uint32(v6))+372:]))
							m.fn16(t45, t46)
							goto l16
						case 3:
							m.fn804(v6+i32(404), v7, v14)
						}
					l16:
						t47 := m.fn113(i32(8), i32(32))
						v1 = t47
						store32(m.memory[uint32(v1):], uint32(i32(-0x80000000)))
						t48 := int64(load64(m.memory[int64(uint32(v6))+392:]))
						store64(m.memory[int64(uint32(v1))+4:], uint64(t48))
						t49 := int32(load32(m.memory[int64(uint32(v6))+400:]))
						store32(m.memory[int64(uint32(v1))+12:], uint32(t49))
						t50 := int64(load64(m.memory[int64(uint32(v6))+404:]))
						store64(m.memory[uint32(v10):], uint64(t50))
						t51 := int32(load32(m.memory[int64(uint32(v6))+412:]))
						store32(m.memory[int64(uint32(v10))+8:], uint32(t51))
						store32(m.memory[int64(uint32(v6))+440:], uint32(v2))
						store64(m.memory[int64(uint32(v6))+432:], uint64(v14))
						m.memory[int64(uint32(v6))+424] = byte(v7)
						store64(m.memory[int64(uint32(v6))+416:], uint64(v8))
						store32(m.memory[int64(uint32(v6))+464:], uint32(i32(1)))
						store32(m.memory[int64(uint32(v6))+460:], uint32(v1))
						store32(m.memory[int64(uint32(v6))+456:], uint32(i32(1)))
						m.fn1369(v6+i32(308), v6+i32(416))
						goto l7
					}
					m.memory[uint32(v1)] = byte(i32(0))
					goto l11
				}
			default:
				m.fn1351(v5, v6+i32(308))
				t29 := int32(load32(m.memory[int64(uint32(v6))+400:]))
				store32(m.memory[int64(uint32(v9))+8:], uint32(t29))
				t30 := int64(load64(m.memory[int64(uint32(v6))+392:]))
				store64(m.memory[uint32(v9):], uint64(t30))
				store32(m.memory[int64(uint32(v6))+416:], uint32(i32(-0x80000000)))
				m.fn338(v5, v6+i32(416))
				goto l7
			}
		}
	}
l0:
	m.fn1351(v5, v6+i32(308))
	store32(m.memory[uint32(v0):], uint32(i32(-1)))
	m.fn1302(v6 + i32(308))
	m.g0 = v6 + i32(480)
}
func (m *Module) fn1285(v0, v1, v2, v3 int32) {
	if uint32(v2) >= uint32(v3) {
		goto l0
	}
	store32(m.memory[int64(uint32(v0))+4:], uint32(i32(0)))
	return
l0:
	v2 = v2 - v3
	if uint32(v2) > uint32(i32(7)) {
		goto l1
	}
	store32(m.memory[int64(uint32(v0))+4:], uint32(i32(0)))
	return
l1:
	{
		v3 = v1 + v3
		t0 := int32(load32(m.memory[int64(uint32(v3))+4:]))
		v1 = t0
		if uint32(v1) <= uint32(v2+i32(-8)) {
			goto l2
		}
		store32(m.memory[int64(uint32(v0))+4:], uint32(i32(0)))
		return
	}
l2:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v1))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v3+i32(8)))
	t1 := int32(load32(m.memory[uint32(v3):]))
	store32(m.memory[uint32(v0):], uint32(t1))
}
func (m *Module) fn1286(v0, v1 int32) int32 {
	var v2 int64
	var v3, v4 int32
	var v5 int64
	var v6 int32
	var v7 int64
	var v8, v9 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
		if t0 != 0 {
			t1 := int64(load64(m.memory[int64(uint32(v0))+16:]))
			t2 := int64(load64(m.memory[int64(uint32(v0))+24:]))
			t3 := m.fn66(t1, t2, v1)
			v2 = t3
			t4 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			v3 = t4
			v4 = v3 & int32(v2)
			v5 = int64(uint64(v2)>>25) & i64(127) * i64(72340172838076673)
			t5 := int32(load32(m.memory[uint32(v0):]))
			v0 = t5
			v6 = i32(0)
			var _ int32
		l5:
			{
				t7 := int64(load64(m.memory[uint32(v0+v4):]))
				v7 = t7
				v2 = v7 ^ v5
				v2 = (v2 ^ i64(-1)) & (v2 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
				{
				l3:
					{
						var p8 int32
						if v2 == 0 {
							p8 = 1
						}
						v8 = p8
						if v8 != 0 {
							goto l1
						}
						t9 := v1
						t10 := v0
						v9 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v2))))>>3) + v4) & v3
						t11 := int32(load32(m.memory[uint32(t10-v9<<3+i32(-8)):]))
						if t9 == t11 {
							goto l2
						}
						v2 = (v2 + i64(-1)) & v2
						goto l3
					}
				l1:
					if v7&(v7<<1)&i64(-0x7f7f7f7f7f7f7f80) == 0 {
						t14 := v4
						v6 = v6 + i32(8)
						v4 = (t14 + v6) & v3
						goto l5
					}
				l2:
					p12 := v0 - v9<<3
					if v8 != 0 {
						p12 = i32(0)
					}
					p13 := p12 + i32(-4)
					if v8 != 0 {
						p13 = i32(0)
					}
					return p13
				}
			}
		}
		return i32(0)
	}
}
func (m *Module) fn1287(v0, v1 int32) {
	var v2, v3, v4, v5, v6 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	t1 := int32(load32(m.memory[uint32(v1):]))
	t2 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	t3 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	t4 := v2 + i32(4)
	v3 = t3
	m.fn1285(t4, t1, t2, v3)
	{
		t5 := int32(load32(m.memory[int64(uint32(v2))+8:]))
		v4 = t5
		if v4 == 0 {
			goto l0
		}
		t6 := int32(load32(m.memory[int64(uint32(v2))+4:]))
		v5 = t6
		t7 := int32(load32(m.memory[int64(uint32(v2))+12:]))
		t8 := v0
		v6 = t7
		store32(m.memory[int64(uint32(t8))+8:], uint32(v6))
		store32(m.memory[uint32(v0):], uint32(v5))
		store32(m.memory[int64(uint32(v1))+8:], uint32(v3+v6+i32(8)))
	}
l0:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v4))
	m.g0 = v2 + i32(16)
}
func (m *Module) fn1288(v0, v1, v2 int32) {
	var v3, v4, v5, v6, v7, v8 int32
	var v9 int64
	var v10, v11 int32
	var v12 int64
	var v13 int32
	var v14, v15 int64
	var v16, v17 int32
	t0 := m.g0
	v3 = t0 - i32(96)
	m.g0 = v3
	m.fn22(v3+i32(72), i32(3))
	t1 := int64(load64(m.memory[int64(uint32(i32(0)))+1286256:]))
	store64(m.memory[int64(uint32(v3))+16:], uint64(t1))
	t2 := int64(load64(m.memory[int64(uint32(i32(0)))+1286264:]))
	store64(m.memory[int64(uint32(v3))+24:], uint64(t2))
	t3 := int64(load64(m.memory[int64(uint32(v3))+80:]))
	store64(m.memory[int64(uint32(v3))+40:], uint64(t3))
	t4 := int64(load64(m.memory[int64(uint32(v3))+72:]))
	store64(m.memory[int64(uint32(v3))+32:], uint64(t4))
	store32(m.memory[int64(uint32(v3))+56:], uint32(i32(0)))
	store32(m.memory[int64(uint32(v3))+52:], uint32(v2))
	store32(m.memory[int64(uint32(v3))+48:], uint32(v1))
	v4 = v3 + i32(16) + i32(16)
l1:
	{
		m.fn1287(v3+i32(60), v3+i32(48))
		t5 := int32(load32(m.memory[int64(uint32(v3))+64:]))
		v5 = t5
		if v5 == 0 {
			goto l0
		}
		t6 := int32(load16(m.memory[int64(uint32(v3))+62:]))
		if t6 != i32(4003) {
			goto l1
		}
		t7 := int32(load32(m.memory[int64(uint32(v3))+68:]))
		v6 = t7
		t8 := int64(load64(m.memory[int64(uint32(v3))+32:]))
		t9 := int64(load64(m.memory[int64(uint32(v3))+40:]))
		t10 := int32(load16(m.memory[int64(uint32(v3))+60:]))
		v7 = t10
		v8 = int32(uint32(v7) >> 4)
		t11 := m.fn529(t8, t9, v8)
		v9 = t11
		t12 := int32(load32(m.memory[int64(uint32(v3))+20:]))
		v10 = t12
		t13 := v10
		v11 = int32(v9)
		v2 = t13 & v11
		v12 = int64(uint64(v9)>>25) & i64(127) * i64(72340172838076673)
		v13 = i32(0)
		t14 := int32(load32(m.memory[int64(uint32(v3))+16:]))
		v1 = t14
	l12:
		{
			t15 := int64(load64(m.memory[uint32(v1+v2):]))
			v14 = t15
			v15 = v14 ^ v12
			v15 = (v15 ^ i64(-1)) & (v15 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
		l3:
			{
				if v15 == 0 {
					if v14&(v14<<1)&i64(-0x7f7f7f7f7f7f7f80) == 0 {
						t44 := v2
						v13 = v13 + i32(8)
						v2 = (t44 + v13) & v10
						goto l12
					}
					{
						t17 := int32(load32(m.memory[int64(uint32(v3))+24:]))
						if t17 != 0 {
							goto l5
						}
						_ = m.fn729(v3+i32(16), v4)
					}
				l5:
					{
						if uint32(v6) >= uint32(i32(2)) {
							goto l6
						}
						v10 = i32(1)
						v2 = i32(0)
						v6 = i32(0)
						goto l7
					l6:
						t19 := int32(load16(m.memory[uint32(v5):]))
						v2 = t19
						store64(m.memory[int64(uint32(v3))+72:], uint64(i64(0x100000000)))
						p20 := i32(10)
						if uint32(v2) < uint32(i32(10)) {
							p20 = v2
						}
						v16 = p20
						v1 = i32(2)
						v2 = i32(0)
						v13 = i32(1)
						var p21 int32
						if uint32(v7) > uint32(i32(79)) {
							p21 = 1
						}
						v17 = p21
					l11:
						store32(m.memory[int64(uint32(v3))+80:], uint32(v2))
						if v2 != v16 {
							t23 := v3 + i32(8)
							t24 := v5
							t25 := v6
							p22 := v1
							if v17 != 0 {
								p22 = v1 + i32(2)
							}
							m.fn1345(t23, t24, t25, p22)
							t26 := int32(m.memory[int64(uint32(v3))+8])
							v1 = t26
							if v1 == i32(255) {
								goto l9
							}
							t27 := int32(load32(m.memory[int64(uint32(v3))+12:]))
							m.fn1346(v3+i32(88), v5, v6, t27)
							t28 := int32(m.memory[int64(uint32(v3))+88])
							v10 = t28
							if v10 == i32(255) {
								goto l9
							}
							t29 := int32(m.memory[int64(uint32(v3))+89])
							v10 = t29<<16 | (v10<<8 | v1)
							t30 := int32(load32(m.memory[int64(uint32(v3))+92:]))
							v1 = t30
							{
								t31 := int32(load32(m.memory[int64(uint32(v3))+72:]))
								if v2 != t31 {
									goto l10
								}
								m.fn1150(v3 + i32(72))
								t32 := int32(load32(m.memory[int64(uint32(v3))+76:]))
								v13 = t32
							}
						l10:
							v7 = v13 + v2*i32(3)
							store16(m.memory[uint32(v7):], uint16(v10))
							m.memory[uint32(v7+i32(2))] = byte(int32(uint32(v10) >> 16))
							v2 = v2 + i32(1)
							goto l11
						}
						v2 = v16
						goto l9
					l9:
						t33 := int32(load32(m.memory[int64(uint32(v3))+76:]))
						v10 = t33
						t34 := int32(load32(m.memory[int64(uint32(v3))+72:]))
						v6 = t34
					}
				l7:
					t35 := int32(load32(m.memory[int64(uint32(v3))+16:]))
					v1 = t35
					t36 := int32(load32(m.memory[int64(uint32(v3))+20:]))
					t37 := v1
					t38 := v1
					v7 = t36
					t39 := m.fn26(t38, v7, v9)
					v5 = t39
					v13 = t37 + v5
					t40 := int32(m.memory[uint32(v13)])
					v16 = t40
					t41 := v13
					v11 = int32(uint32(v11) >> 25)
					m.memory[uint32(t41)] = byte(v11)
					m.memory[uint32(v1+v7&(v5+i32(-8))+i32(8))] = byte(v11)
					v1 = v1 - v5<<4
					store32(m.memory[uint32(v1+i32(-4)):], uint32(v2))
					store32(m.memory[uint32(v1+i32(-8)):], uint32(v10))
					store32(m.memory[uint32(v1+i32(-12)):], uint32(v6))
					store16(m.memory[uint32(v1+i32(-16)):], uint16(v8))
					t42 := int32(load32(m.memory[int64(uint32(v3))+28:]))
					store32(m.memory[int64(uint32(v3))+28:], uint32(t42+i32(1)))
					t43 := int32(load32(m.memory[int64(uint32(v3))+24:]))
					store32(m.memory[int64(uint32(v3))+24:], uint32(t43-v16&i32(1)))
					goto l1
				}
				t16 := int32(load16(m.memory[uint32(v1-(int32(uint32(int64(bits.TrailingZeros64(uint64(v15))))>>3)+v2)&v10<<4+i32(-16)):]))
				if t16 == v8 {
					goto l1
				}
				v15 = (v15 + i64(-1)) & v15
				goto l3
			}
		}
	}
l0:
	t45 := int64(load64(m.memory[int64(uint32(v3))+40:]))
	store64(m.memory[int64(uint32(v0))+24:], uint64(t45))
	t46 := int64(load64(m.memory[int64(uint32(v3))+32:]))
	store64(m.memory[int64(uint32(v0))+16:], uint64(t46))
	t47 := int64(load64(m.memory[int64(uint32(v3))+24:]))
	store64(m.memory[int64(uint32(v0))+8:], uint64(t47))
	t48 := int64(load64(m.memory[int64(uint32(v3))+16:]))
	store64(m.memory[uint32(v0):], uint64(t48))
	m.g0 = v3 + i32(96)
}
func (m *Module) fn1289(v0, v1 int32) {
	var v2 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		v2 = t0
		t1 := int32(load32(m.memory[uint32(v0):]))
		if v2 != t1 {
			goto l0
		}
		m.fn1149(v0)
	}
l0:
	t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	memory_copy(m.memory, uint32(t2+v2*i32(40)), uint32(v1), uint32(i32(40)))
	store32(m.memory[int64(uint32(v0))+8:], uint32(v2+i32(1)))
}
func (m *Module) fn1290(v0 int32) {
	var v1, v2, v3 int32
	t0 := int32(load32(m.memory[int64(uint32(v0))+64:]))
	v1 = t0
	v2 = v1 + i32(8)
	t1 := int32(load32(m.memory[int64(uint32(v0))+68:]))
	v3 = t1
l1:
	if v3 == 0 {
		goto l0
	}
	v3 = v3 + i32(-1)
	m.fn969(v2)
	v2 = v2 + i32(24)
	goto l1
l0:
	t2 := int32(load32(m.memory[int64(uint32(v0))+60:]))
	m.fn1201(t2, v1)
	m.fn969(v0 + i32(72))
	m.fn1302(v0 + i32(84))
	m.fn1303(v0 + i32(16))
	m.fn1293(v0 + i32(96))
}
func (m *Module) fn1291(v0, v1, v2, v3 int32) {
	var v4, v5, v6, v7, v8, v9, v10 int32
	t0 := m.g0
	v4 = t0 - i32(48)
	m.g0 = v4
	t1 := m.fn113(i32(4), i32(12))
	v5 = t1
	store32(m.memory[int64(uint32(v5))+8:], uint32(i32(0)))
	store32(m.memory[int64(uint32(v5))+4:], uint32(v3))
	store32(m.memory[uint32(v5):], uint32(v2))
	store32(m.memory[int64(uint32(v4))+8:], uint32(i32(1)))
	store32(m.memory[int64(uint32(v4))+4:], uint32(v5))
	store32(m.memory[uint32(v4):], uint32(i32(1)))
	v3 = i32(1)
l6:
	{
		if v3 == 0 {
			goto l0
		}
		t2 := v5
		v6 = v3 * i32(12)
		v2 = t2 + v6
		v7 = v2 + i32(-12)
		if v7 == 0 {
			goto l0
		}
		t3 := int32(load32(m.memory[uint32(v7):]))
		t4 := int32(load32(m.memory[uint32(v2+i32(-8)):]))
		t5 := v4 + i32(24)
		v2 = v2 + i32(-4)
		t6 := int32(load32(m.memory[uint32(v2):]))
		m.fn1285(t5, t3, t4, t6)
		t7 := int32(load32(m.memory[int64(uint32(v4))+28:]))
		v8 = t7
		if v8 == 0 {
			t26 := v4
			v3 = v3 + i32(-1)
			store32(m.memory[int64(uint32(t26))+8:], uint32(v3))
			goto l6
		}
		t8 := int32(load16(m.memory[int64(uint32(v4))+26:]))
		v7 = t8
		t9 := int32(load16(m.memory[int64(uint32(v4))+24:]))
		v9 = t9
		t10 := int32(load32(m.memory[int64(uint32(v4))+32:]))
		t11 := v2
		v10 = t10
		t12 := int32(load32(m.memory[uint32(v2):]))
		store32(m.memory[uint32(t11):], uint32(v10+t12+i32(8)))
		m.fn1348(v4+i32(24), v1)
		{
			t13 := int32(load32(m.memory[int64(uint32(v4))+24:]))
			v2 = t13
			if v2 == i32(-1) {
				if v9&i32(15) != i32(15) {
					m.fn1349(v1, v7, v8, v10)
					goto l6
				}
				if v7 == i32(1008) {
					t17 := int32(m.memory[int64(uint32(v1))+110])
					if t17 == 0 {
						goto l6
					}
					store32(m.memory[int64(uint32(v4))+20:], uint32(i32(0)))
					store32(m.memory[int64(uint32(v4))+16:], uint32(v10))
					store32(m.memory[int64(uint32(v4))+12:], uint32(v8))
					{
					l10:
						{
							m.fn1287(v4+i32(24), v4+i32(12))
							t18 := int32(load32(m.memory[int64(uint32(v4))+28:]))
							v2 = t18
							if v2 == 0 {
								goto l9
							}
							t19 := int32(load16(m.memory[int64(uint32(v4))+26:]))
							if t19<<16 != i32(0x3f10000) {
								goto l10
							}
						}
						t20 := int32(load32(m.memory[int64(uint32(v4))+32:]))
						if uint32(t20) < uint32(i32(4)) {
							goto l9
						}
						t21 := int32(load32(m.memory[uint32(v2):]))
						if t21 < i32(0) {
							goto l6
						}
					}
				l9:
					m.fn1292(v1, i32(0), v4)
					m.memory[int64(uint32(v1))+108] = byte(i32(1))
					m.fn1291(v4+i32(24), v1, v8, v10)
					{
						t22 := int32(load32(m.memory[int64(uint32(v4))+24:]))
						v2 = t22
						if v2 == i32(-1) {
							m.fn1292(v1, i32(0), v4)
							m.memory[int64(uint32(v1))+108] = byte(i32(0))
							goto l6
						}
						t23 := int32(load32(m.memory[int64(uint32(v4))+44:]))
						store32(m.memory[int64(uint32(v0))+20:], uint32(t23))
						t24 := int64(load64(m.memory[int64(uint32(v4))+36:]))
						store64(m.memory[int64(uint32(v0))+12:], uint64(t24))
						t25 := int64(load64(m.memory[int64(uint32(v4))+28:]))
						store64(m.memory[int64(uint32(v0))+4:], uint64(t25))
						store32(m.memory[uint32(v0):], uint32(v2))
						goto l3
					}
				}
				if v7 == i32(1016) {
					goto l6
				}
				if v7 == i32(4041) {
					goto l6
				}
				if v7 == i32(4080) {
					if uint32(v9&i32(0xffff)) >= uint32(i32(16)) {
						goto l6
					}
					goto l8
				}
				if v7 != i32(12052) {
					goto l8
				}
				m.memory[int64(uint32(v1))+109] = byte(i32(1))
				goto l6
			}
			t14 := int32(load32(m.memory[int64(uint32(v4))+44:]))
			store32(m.memory[int64(uint32(v0))+20:], uint32(t14))
			t15 := int64(load64(m.memory[int64(uint32(v4))+36:]))
			store64(m.memory[int64(uint32(v0))+12:], uint64(t15))
			t16 := int64(load64(m.memory[int64(uint32(v4))+28:]))
			store64(m.memory[int64(uint32(v0))+4:], uint64(t16))
			store32(m.memory[uint32(v0):], uint32(v2))
			goto l3
		}
	}
l0:
	store32(m.memory[uint32(v0):], uint32(i32(-1)))
	goto l3
l8:
	{
		if uint32(v3) > uint32(i32(63)) {
			goto l12
		}
		{
			t27 := int32(load32(m.memory[uint32(v4):]))
			if v3 != t27 {
				goto l13
			}
			m.fn272(v4)
			t28 := int32(load32(m.memory[int64(uint32(v4))+4:]))
			v5 = t28
		}
	l13:
		v2 = v5 + v6
		store32(m.memory[int64(uint32(v2))+8:], uint32(i32(0)))
		store32(m.memory[int64(uint32(v2))+4:], uint32(v10))
		store32(m.memory[uint32(v2):], uint32(v8))
		t29 := v4
		v3 = v3 + i32(1)
		store32(m.memory[int64(uint32(t29))+8:], uint32(v3))
		t30 := int32(load32(m.memory[int64(uint32(v4))+4:]))
		v5 = t30
		goto l6
	}
l12:
	store32(m.memory[int64(uint32(v4))+28:], uint32(i32(5)))
	store32(m.memory[int64(uint32(v4))+24:], uint32(i32(1075188)))
	m.fn73(v0+i32(4), i32(1050111), v4+i32(24))
	store32(m.memory[int64(uint32(v0))+20:], uint32(i32(16)))
	store32(m.memory[int64(uint32(v0))+16:], uint32(i32(1075192)))
	store32(m.memory[uint32(v0):], uint32(i32(-0x7ffffffd)))
l3:
	t31 := int32(load32(m.memory[uint32(v4):]))
	t32 := int32(load32(m.memory[int64(uint32(v4))+4:]))
	m.fn136(t31, t32, i32(4), i32(12))
	m.g0 = v4 + i32(48)
}
func (m *Module) fn1292(v0, v1, v2 int32) {
	var v3, v4, v5 int32
	var v6 int64
	t0 := m.g0
	v3 = t0 - i32(16)
	m.g0 = v3
	m.fn1350(v0)
	v4 = v0 + i32(72)
	m.fn1351(v4, v0+i32(84))
	{
		t1 := int32(load32(m.memory[int64(uint32(v0))+80:]))
		if t1 == 0 {
			goto l0
		}
		t2 := int32(load32(m.memory[int64(uint32(v4))+8:]))
		v5 = t2
		store32(m.memory[int64(uint32(v0))+80:], uint32(i32(0)))
		t3 := int64(load64(m.memory[uint32(v4):]))
		v6 = t3
		store64(m.memory[int64(uint32(v0))+72:], uint64(i64(0x800000000)))
		store32(m.memory[int64(uint32(v3))+8:], uint32(v5))
		store64(m.memory[uint32(v3):], uint64(v6))
		t4 := int32(m.memory[int64(uint32(v0))+108])
		v5 = t4
		{
			t5 := int32(load32(m.memory[int64(uint32(v0))+68:]))
			v4 = t5
			t6 := int32(load32(m.memory[int64(uint32(v0))+60:]))
			if v4 != t6 {
				goto l1
			}
			m.fn289(v0 + i32(60))
		}
	l1:
		store32(m.memory[int64(uint32(v0))+68:], uint32(v4+i32(1)))
		t7 := int32(load32(m.memory[int64(uint32(v0))+64:]))
		v0 = t7 + v4*i32(24)
		store32(m.memory[int64(uint32(v0))+4:], uint32(v2))
		store32(m.memory[uint32(v0):], uint32(v1))
		t8 := int64(load64(m.memory[uint32(v3):]))
		store64(m.memory[int64(uint32(v0))+8:], uint64(t8))
		t9 := int32(load32(m.memory[int64(uint32(v3))+8:]))
		store32(m.memory[int64(uint32(v0))+16:], uint32(t9))
		m.memory[int64(uint32(v0))+20] = byte(v5)
	}
l0:
	m.g0 = v3 + i32(16)
}
func (m *Module) fn1293(v0 int32) {
	var v1, v2, v3, v4, v5, v6, v7, v8, v9 int32
	var v10 int64
	var v11 int32
	t0 := m.g0
	v1 = t0 - i32(16)
	m.g0 = v1
	t1 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	v2 = t1
	t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	v3 = t2
	v4 = i32(0)
l1:
	{
		if v4 == v2 {
			goto l0
		}
		v5 = v4 * i32(40)
		v6 = v4 + i32(1)
		v4 = v6
		v5 = v3 + v5
		t3 := int32(load32(m.memory[int64(uint32(v5))+12:]))
		v7 = t3
		if v7 == 0 {
			goto l1
		}
		t4 := int32(load32(m.memory[int64(uint32(v5))+8:]))
		v8 = t4
		{
			t5 := int32(load32(m.memory[int64(uint32(v5))+20:]))
			v9 = t5
			if v9 == 0 {
				goto l2
			}
			v4 = v8 + i32(8)
			t6 := int64(load64(m.memory[uint32(v8):]))
			v10 = (t6 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
			v5 = v8
		l5:
			if v9 == 0 {
				goto l2
			}
		l4:
			{
				if v10 != i64(0) {
					v11 = v5 - int32(int64(bits.TrailingZeros64(uint64(v10))))<<1&i32(240)
					t8 := int32(load32(m.memory[uint32(v11+i32(-12)):]))
					t9 := int32(load32(m.memory[uint32(v11+i32(-8)):]))
					m.fn765(t8, t9)
					v9 = v9 + i32(-1)
					v10 = (v10 + i64(-1)) & v10
					goto l5
				}
				v5 = v5 + i32(-128)
				t7 := int64(load64(m.memory[uint32(v4):]))
				v10 = (t7 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
				v4 = v4 + i32(8)
				goto l4
			}
		}
	l2:
		m.fn39(v1+i32(4), i32(16), i32(8), v7+i32(1))
		t10 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		t11 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t12 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		m.fn40(v8-t10, t11, t12)
		v4 = v6
		goto l1
	}
l0:
	t13 := int32(load32(m.memory[uint32(v0):]))
	m.fn136(t13, v3, i32(8), i32(40))
	m.g0 = v1 + i32(16)
}
func (m *Module) fn1294(v0, v1, v2, v3, v4, v5, v6, v7, v8 int32) {
	var v9, v10, v11 int32
	var v12 int64
	var v13, v14 int32
	var v15 int64
	t0 := m.g0
	v9 = t0 - i32(80)
	m.g0 = v9
	store32(m.memory[int64(uint32(v9))+8:], uint32(i32(0)))
	store32(m.memory[int64(uint32(v9))+4:], uint32(v3))
	store32(m.memory[uint32(v9):], uint32(v2))
	v10 = i32(0)
	{
		{
		l8:
			{
				m.fn1287(v9+i32(12), v9)
				{
					t1 := int32(load32(m.memory[int64(uint32(v9))+16:]))
					v11 = t1
					if v11 == 0 {
						store64(m.memory[int64(uint32(v9))+72:], uint64(v12))
						store32(m.memory[int64(uint32(v9))+68:], uint32(v10))
						m.fn1347(v9+i32(44), v1, v9+i32(68), v4, v5, v6, v8, v7)
						t9 := int32(load32(m.memory[int64(uint32(v9))+52:]))
						v3 = t9
						t10 := int32(load32(m.memory[int64(uint32(v9))+48:]))
						v2 = t10
						{
							t11 := int32(load32(m.memory[int64(uint32(v9))+44:]))
							v11 = t11
							if v11 == i32(-1) {
								m.fn1292(v1, v2, v3)
								store32(m.memory[uint32(v0):], uint32(i32(-1)))
								goto l3
							}
							t12 := int32(load32(m.memory[int64(uint32(v9))+64:]))
							store32(m.memory[int64(uint32(v0))+20:], uint32(t12))
							t13 := int64(load64(m.memory[int64(uint32(v9))+56:]))
							store64(m.memory[int64(uint32(v0))+12:], uint64(t13))
							store32(m.memory[int64(uint32(v0))+8:], uint32(v3))
							store32(m.memory[int64(uint32(v0))+4:], uint32(v2))
							store32(m.memory[uint32(v0):], uint32(v11))
							goto l3
						}
					}
					t2 := int32(load32(m.memory[int64(uint32(v9))+20:]))
					v2 = t2
					t3 := int32(load16(m.memory[int64(uint32(v9))+14:]))
					v3 = t3
					if v3 != i32(1011) {
						t14 := int32(load16(m.memory[int64(uint32(v9))+12:]))
						v13 = t14
						m.fn1348(v9+i32(44), v1)
						{
							t15 := int32(load32(m.memory[int64(uint32(v9))+44:]))
							v14 = t15
							if v14 == i32(-1) {
								if v13&i32(15) != i32(15) {
									m.fn1349(v1, v3, v11, v2)
									goto l8
								}
								if v3 == i32(1008) {
									goto l8
								}
								if v3 == i32(1016) {
									goto l8
								}
								if v3 == i32(4041) {
									goto l8
								}
								if v3 == i32(4080) {
									if uint32(v13&i32(0xffff)) >= uint32(i32(16)) {
										goto l8
									}
									goto l10
								}
								if v3 != i32(12052) {
									goto l10
								}
								m.memory[int64(uint32(v1))+109] = byte(i32(1))
								goto l8
							}
							t16 := int32(load32(m.memory[int64(uint32(v9))+64:]))
							store32(m.memory[int64(uint32(v9))+40:], uint32(t16))
							t17 := int64(load64(m.memory[int64(uint32(v9))+56:]))
							store64(m.memory[int64(uint32(v9))+32:], uint64(t17))
							t18 := int64(load64(m.memory[int64(uint32(v9))+48:]))
							store64(m.memory[int64(uint32(v9))+24:], uint64(t18))
							goto l6
						}
					}
					store64(m.memory[int64(uint32(v9))+72:], uint64(v12))
					store32(m.memory[int64(uint32(v9))+68:], uint32(v10))
					m.fn1347(v9+i32(44), v1, v9+i32(68), v4, v5, v6, v8, v7)
					t4 := int32(load32(m.memory[int64(uint32(v9))+52:]))
					v3 = t4
					t5 := int32(load32(m.memory[int64(uint32(v9))+48:]))
					v10 = t5
					t6 := int32(load32(m.memory[int64(uint32(v9))+44:]))
					v13 = t6
					if v13 == i32(-1) {
						m.fn1292(v1, v10, v3)
						m.memory[int64(uint32(v1))+108] = byte(v7)
						v10 = i32(0)
						{
							if uint32(v2) < uint32(i32(4)) {
								goto l11
							}
							t19 := int64(load32(m.memory[uint32(v11):]))
							v15 = t19
							v12 = i64(0)
							{
								if v2 < i32(16) {
									goto l12
								}
								t20 := int64(load32(m.memory[int64(uint32(v11))+12:]))
								v12 = t20 << 32
							}
						l12:
							v12 = v12 | v15
							v10 = i32(1)
						}
					l11:
						if v7 != 0 {
							goto l8
						}
						if v10 == 0 {
							goto l13
						}
						t21 := m.fn1286(v4, int32(v12))
						v3 = t21
						if v3 == 0 {
							goto l13
						}
						t22 := int32(load32(m.memory[uint32(v3):]))
						m.fn1285(v9+i32(44), v5, v6, t22)
						t23 := int32(load32(m.memory[int64(uint32(v9))+48:]))
						v3 = t23
						if v3 == 0 {
							goto l13
						}
						t24 := int32(load16(m.memory[int64(uint32(v9))+46:]))
						if t24&i32(0xffff) != i32(1006) {
							goto l13
						}
						t25 := int32(load32(m.memory[int64(uint32(v9))+52:]))
						v2 = t25
						store32(m.memory[int64(uint32(v9))+76:], uint32(i32(0)))
						store32(m.memory[int64(uint32(v9))+72:], uint32(v2))
						store32(m.memory[int64(uint32(v9))+68:], uint32(v3))
					l14:
						{
							m.fn1287(v9+i32(44), v9+i32(68))
							t26 := int32(load32(m.memory[int64(uint32(v9))+48:]))
							v3 = t26
							if v3 == 0 {
								goto l13
							}
							t27 := int32(load16(m.memory[int64(uint32(v9))+46:]))
							if t27<<16 != i32(0x3ef0000) {
								goto l14
							}
						}
						v2 = i32(0)
						t28 := int32(load32(m.memory[int64(uint32(v9))+52:]))
						if t28 < i32(16) {
							goto l15
						}
						t29 := int32(load32(m.memory[int64(uint32(v1))+104:]))
						v11 = t29 * i32(40)
						t30 := int32(load32(m.memory[int64(uint32(v1))+100:]))
						v13 = t30
						t31 := int32(load32(m.memory[int64(uint32(v3))+12:]))
						v14 = t31
						v2 = i32(0)
						v3 = i32(0)
					l17:
						if v11 == 0 {
							goto l15
						}
						{
							t32 := int32(load32(m.memory[uint32(v13):]))
							if t32 != v14 {
								v11 = v11 + i32(-40)
								v3 = v3 + i32(1)
								v13 = v13 + i32(40)
								goto l17
							}
							v2 = v3
							goto l15
						}
					}
					t7 := int32(load32(m.memory[int64(uint32(v9))+64:]))
					store32(m.memory[int64(uint32(v0))+20:], uint32(t7))
					t8 := int64(load64(m.memory[int64(uint32(v9))+56:]))
					store64(m.memory[int64(uint32(v0))+12:], uint64(t8))
					store32(m.memory[int64(uint32(v0))+8:], uint32(v3))
					store32(m.memory[int64(uint32(v0))+4:], uint32(v10))
					store32(m.memory[uint32(v0):], uint32(v13))
					goto l3
				}
			l13:
				v2 = i32(0)
			l15:
				store32(m.memory[int64(uint32(v1))+56:], uint32(v2))
				goto l8
			l10:
				m.fn1291(v9+i32(44), v1, v11, v2)
				t33 := int32(load32(m.memory[int64(uint32(v9))+44:]))
				v14 = t33
				if v14 == i32(-1) {
					goto l8
				}
			}
			t34 := int32(load32(m.memory[int64(uint32(v9))+64:]))
			store32(m.memory[int64(uint32(v9))+40:], uint32(t34))
			t35 := int64(load64(m.memory[int64(uint32(v9))+56:]))
			store64(m.memory[int64(uint32(v9))+32:], uint64(t35))
			t36 := int64(load64(m.memory[int64(uint32(v9))+48:]))
			store64(m.memory[int64(uint32(v9))+24:], uint64(t36))
		}
	l6:
		store32(m.memory[uint32(v0):], uint32(v14))
		t37 := int64(load64(m.memory[int64(uint32(v9))+24:]))
		store64(m.memory[int64(uint32(v0))+4:], uint64(t37))
		t38 := int64(load64(m.memory[int64(uint32(v9))+32:]))
		store64(m.memory[int64(uint32(v0))+12:], uint64(t38))
		t39 := int32(load32(m.memory[int64(uint32(v9))+40:]))
		store32(m.memory[int64(uint32(v0))+20:], uint32(t39))
	}
l3:
	m.g0 = v9 + i32(80)
}
func (m *Module) fn1295(v0, v1, v2, v3, v4 int32) {
	var v5, v6 int32
	var v7 int64
	var v8, v9, v10, v11, v12, v13 int32
	t0 := m.g0
	v5 = t0 - i32(192)
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
		if uint32(v4) >= uint32(v6) {
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
			p3 := i32(1079563)
			if v1 != 0 {
				p3 = i32(1079551)
			}
			store32(m.memory[int64(uint32(t1))+20:], uint32(p3))
			t5 := v0
			p4 := i32(9)
			if v1 != 0 {
				p4 = i32(10)
			}
			store32(m.memory[int64(uint32(t5))+16:], uint32(p4))
			t7 := v0
			p6 := i32(1079554)
			if v1 != 0 {
				p6 = i32(1079541)
			}
			store32(m.memory[int64(uint32(t7))+12:], uint32(p6))
			goto l3
		}
		store32(m.memory[uint32(v0):], uint32(i32(-2)))
		goto l3
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
				goto l8
			}
			{
				{
					{
						v4 = v4 - v1
						if uint32(v4) < uint32(i32(4)) {
							store32(m.memory[uint32(v0):], uint32(i32(-2)))
							goto l3
						}
						if uint32(v4) > uint32(i32(32)) {
							if v4 != i32(33) {
								v1 = v3 + v1
								v6 = v1 + i32(34)
								v4 = v4 + i32(-34)
								v3 = i32(-1)
								t11 := int32(m.memory[int64(uint32(v1))+32])
								if t11 != 0 {
									goto l12
								}
								t12 := int32(load32(m.memory[uint32(v1):]))
								v1 = t12
								store32(m.memory[int64(uint32(v5))+20:], uint32(i32(0)))
								store64(m.memory[int64(uint32(v5))+12:], uint64(i64(0x100000000)))
								m.fn321(v5+i32(176), i32(0x8000))
								m.fn322(v5, v5+i32(176))
								t13 := int64(load64(m.memory[uint32(v5):]))
								v7 = t13
								m.fn121(v5 + i32(48))
								store64(m.memory[int64(uint32(v5))+40:], uint64(i64(0)))
								store64(m.memory[int64(uint32(v5))+32:], uint64(v7))
								store32(m.memory[int64(uint32(v5))+28:], uint32(v4))
								store32(m.memory[int64(uint32(v5))+24:], uint32(v6))
								t15 := v5
								p14 := i32(0x8000000)
								if uint32(v1) < uint32(i32(0x8000000)) {
									p14 = v1
								}
								v7 = int64(uint32(p14))
								store64(m.memory[int64(uint32(t15))+144:], uint64(v7))
								store64(m.memory[int64(uint32(v5))+136:], uint64(v7))
								m.fn960(v5+i32(176), v5+i32(24), v5+i32(12))
								t16 := int32(m.memory[int64(uint32(v5))+176])
								if t16 != i32(255) {
									goto l13
								}
								t17 := int32(load32(m.memory[int64(uint32(v5))+180:]))
								if t17 == 0 {
									goto l14
								}
								v8 = i32(8192)
								t18 := int32(load32(m.memory[int64(uint32(v5))+12:]))
								v9 = t18
								t19 := int32(load32(m.memory[int64(uint32(v5))+20:]))
								v3 = t19
							l29:
								{
									if v9|v3 != 0 {
										goto l15
									}
									m.fn960(v5+i32(176), v5+i32(24), v5+i32(12))
									t20 := int32(m.memory[int64(uint32(v5))+176])
									if t20 != i32(255) {
										goto l13
									}
									t21 := int32(load32(m.memory[int64(uint32(v5))+180:]))
									if t21 == 0 {
										goto l14
									}
									t22 := int32(load32(m.memory[int64(uint32(v5))+12:]))
									v9 = t22
									t23 := int32(load32(m.memory[int64(uint32(v5))+20:]))
									v3 = t23
								}
							l15:
								{
									{
										if v3 != v9 {
											goto l16
										}
										t24 := m.fn351(v5+i32(12), v9, i32(32))
										if t24 != i32(-1) {
											v3 = i32(0)
											v1 = i32(1)
											goto l31
										}
										t25 := int32(load32(m.memory[int64(uint32(v5))+12:]))
										v9 = t25
										t26 := int32(load32(m.memory[int64(uint32(v5))+20:]))
										v3 = t26
									}
								l16:
									t27 := int32(load32(m.memory[int64(uint32(v5))+16:]))
									v4 = t27
									v1 = i32(0)
									m.memory[int64(uint32(v5))+164] = byte(i32(0))
									store32(m.memory[int64(uint32(v5))+160:], uint32(i32(0)))
									store32(m.memory[int64(uint32(v5))+152:], uint32(v4+v3))
									t28 := v5
									t29 := v8
									v10 = v9 - v3
									p30 := v10
									if uint32(v8) < uint32(v10) {
										p30 = t29
									}
									v11 = p30
									store32(m.memory[int64(uint32(t28))+156:], uint32(v11))
									t31 := int64(load64(m.memory[int64(uint32(v5))+144:]))
									v7 = t31
								l30:
									{
										{
											if v7 != i64(0) {
												goto l18
											}
											m.memory[int64(uint32(v5))+168] = byte(i32(255))
											goto l19
										l18:
											{
												{
													t32 := int32(load32(m.memory[int64(uint32(v5))+156:]))
													t33 := v7
													v4 = t32 - v1
													if uint64(t33) < uint64(uint32(v4)) {
														goto l20
													}
													m.fn950(v5+i32(168), v5+i32(24), v5+i32(152))
													t34 := int64(load64(m.memory[int64(uint32(v5))+144:]))
													t35 := int32(load32(m.memory[int64(uint32(v5))+160:]))
													v4 = t35
													v7 = t34 - int64(uint32(v4-v1))
													goto l21
												}
											l20:
												t36 := int32(m.memory[int64(uint32(v5))+164])
												v6 = t36
												m.memory[int64(uint32(v5))+188] = byte(i32(0))
												store32(m.memory[int64(uint32(v5))+184:], uint32(i32(0)))
												t37 := v5
												v12 = int32(v7)
												store32(m.memory[int64(uint32(t37))+180:], uint32(v12))
												t38 := int32(load32(m.memory[int64(uint32(v5))+152:]))
												t39 := v5
												v13 = t38 + v1
												store32(m.memory[int64(uint32(t39))+176:], uint32(v13))
												{
													{
														if v6 != 0 {
															goto l22
														}
														m.fn950(v5+i32(168), v5+i32(24), v5+i32(176))
														t40 := int32(load32(m.memory[int64(uint32(v5))+184:]))
														v6 = t40
														t41 := int32(m.memory[int64(uint32(v5))+188])
														if t41 == 0 {
															goto l23
														}
														m.fn1094(v13+v12, v4-v12)
														m.memory[int64(uint32(v5))+164] = byte(i32(1))
														goto l23
													}
												l22:
													m.memory[int64(uint32(v5))+188] = byte(i32(1))
													m.fn950(v5+i32(168), v5+i32(24), v5+i32(176))
													t42 := int32(load32(m.memory[int64(uint32(v5))+184:]))
													v6 = t42
												}
											l23:
												t43 := v5
												v4 = v1 + v6
												store32(m.memory[int64(uint32(t43))+160:], uint32(v4))
												t44 := int64(load64(m.memory[int64(uint32(v5))+144:]))
												v7 = t44 - int64(uint32(v6))
											}
										l21:
											store64(m.memory[int64(uint32(v5))+144:], uint64(v7))
											t45 := int32(m.memory[int64(uint32(v5))+168])
											if t45 != i32(255) {
												t47 := m.fn313(v5 + i32(168))
												if t47 != 0 {
													t51 := int32(load32(m.memory[int64(uint32(v5))+168:]))
													t52 := int32(load32(m.memory[int64(uint32(v5))+172:]))
													m.fn119(t51, t52)
													v1 = v4
													goto l30
												}
												store32(m.memory[int64(uint32(v5))+20:], uint32(v3+v4))
												t48 := int32(load32(m.memory[int64(uint32(v5))+168:]))
												v1 = t48
												t49 := int32(load32(m.memory[int64(uint32(v5))+172:]))
												v3 = t49
												goto l26
											}
											v1 = v4
										}
									l19:
										t46 := v5
										v3 = v3 + v1
										store32(m.memory[int64(uint32(t46))+20:], uint32(v3))
										if v1 != 0 {
											{
												t50 := int32(m.memory[int64(uint32(v5))+164])
												if t50&i32(1) == 0 {
													goto l28
												}
												if uint32(v10) < uint32(v8) {
													goto l29
												}
												if v1 != v11 {
													goto l29
												}
												if v8 <= i32(-1) {
													goto l28
												}
												v8 = v8 << 1
												goto l29
											}
										l28:
											v8 = i32(-1)
											goto l29
										}
										v1 = i32(255)
										goto l26
									}
								}
							}
							store32(m.memory[uint32(v0):], uint32(i32(-2)))
							goto l3
						}
						store32(m.memory[uint32(v0):], uint32(i32(-2)))
						goto l3
					l13:
						t53 := int32(load32(m.memory[int64(uint32(v5))+180:]))
						v3 = t53
						t54 := int32(load32(m.memory[int64(uint32(v5))+176:]))
						v1 = t54
					}
				l26:
					if v1&i32(255) != i32(255) {
						goto l31
					}
				l14:
					t55 := int32(load32(m.memory[int64(uint32(v5))+20:]))
					v4 = t55
					t56 := int32(load32(m.memory[int64(uint32(v5))+16:]))
					v6 = t56
					t57 := int32(load32(m.memory[int64(uint32(v5))+12:]))
					v3 = t57
					m.fn1352(v5 + i32(24))
				}
			l12:
				store32(m.memory[int64(uint32(v0))+24:], uint32(i32(3)))
				t58 := v0
				var p59 int32
				if v2&i32(0xffff) == i32(61466) {
					p59 = 1
				}
				v1 = p59
				p60 := i32(1079526)
				if v1 != 0 {
					p60 = i32(1079538)
				}
				store32(m.memory[int64(uint32(t58))+20:], uint32(p60))
				store32(m.memory[int64(uint32(v0))+16:], uint32(i32(9)))
				t62 := v0
				p61 := i32(1079517)
				if v1 != 0 {
					p61 = i32(1079529)
				}
				store32(m.memory[int64(uint32(t62))+12:], uint32(p61))
				store32(m.memory[int64(uint32(v0))+8:], uint32(v4))
				store32(m.memory[int64(uint32(v0))+4:], uint32(v6))
				store32(m.memory[uint32(v0):], uint32(v3))
				goto l3
			}
		l31:
			m.fn961(v1, v3)
			store32(m.memory[uint32(v0):], uint32(i32(-2)))
			m.fn1352(v5 + i32(24))
			t63 := int32(load32(m.memory[int64(uint32(v5))+12:]))
			t64 := int32(load32(m.memory[int64(uint32(v5))+16:]))
			m.fn16(t63, t64)
			goto l3
		}
	l8:
		store32(m.memory[uint32(v0):], uint32(i32(-2)))
	}
l3:
	m.g0 = v5 + i32(192)
}
func (m *Module) fn1296(v0, v1, v2, v3, v4, v5 int32) {
	var v6, v7, v8 int32
	var v9 int64
	var v10, v11 int32
	var v12 int64
	var v13, v14 int32
	var v15 int64
	var v16 int32
	t0 := m.g0
	v6 = t0 - i32(80)
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
				t7 := m.fn540(t2, t3, t6, v8)
				v9 = t7
				t8 := int32(load32(m.memory[int64(uint32(v1))+4:]))
				v10 = t8
				v11 = v10 & int32(v9)
				v12 = int64(uint64(v9)>>25) & i64(127) * i64(72340172838076673)
				t9 := int32(load32(m.memory[uint32(v1):]))
				v13 = t9
				v14 = i32(0)
			l6:
				{
					t10 := int64(load64(m.memory[uint32(v13+v11):]))
					v15 = t10
					v9 = v15 ^ v12
					v9 = (v9 ^ i64(-1)) & (v9 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
				l3:
					{
						if v9 == 0 {
							if v15&(v15<<1)&i64(-0x7f7f7f7f7f7f7f80) == 0 {
								t17 := v11
								v14 = v14 + i32(8)
								v11 = (t17 + v14) & v10
								goto l6
							}
							goto l0
						}
						t11 := v7
						t12 := v8
						v16 = v13 - (int32(uint32(int64(bits.TrailingZeros64(uint64(v9))))>>3)+v11)&v10<<4
						t13 := int32(load32(m.memory[uint32(v16+i32(-12)):]))
						t14 := int32(load32(m.memory[uint32(v16+i32(-8)):]))
						t15 := m.fn544(t11, t12, t13, t14)
						if t15 != 0 {
							store32(m.memory[uint32(v0):], uint32(i32(-1)))
							t16 := int32(load32(m.memory[uint32(v16+i32(-4)):]))
							store32(m.memory[int64(uint32(v0))+4:], uint32(t16))
							goto l5
						}
						v9 = (v9 + i64(-1)) & v9
						goto l3
					}
				}
			}
		l0:
			t18 := int32(load32(m.memory[int64(uint32(v1))+32:]))
			t19 := v1
			v7 = t18 + v5
			store32(m.memory[int64(uint32(t19))+32:], uint32(v7))
			{
				if uint32(v7) > uint32(i32(0x8000000)) {
					goto l7
				}
				t20 := int32(load32(m.memory[int64(uint32(v1))+44:]))
				v7 = t20
				t21 := int32(load32(m.memory[int64(uint32(v3))+4:]))
				t22 := int32(load32(m.memory[int64(uint32(v3))+8:]))
				m.fn31(v6+i32(16), t21, t22)
				t23 := int64(load64(m.memory[int64(uint32(v1))+16:]))
				t24 := int64(load64(m.memory[int64(uint32(v1))+24:]))
				t25 := int32(load32(m.memory[int64(uint32(v6))+20:]))
				t26 := int32(load32(m.memory[int64(uint32(v6))+24:]))
				t27 := m.fn540(t23, t24, t25, t26)
				v9 = t27
				store32(m.memory[int64(uint32(v6))+68:], uint32(v6+i32(16)))
				{
					t28 := int32(load32(m.memory[int64(uint32(v1))+8:]))
					if t28 != 0 {
						goto l8
					}
					_ = m.fn664(v1, v1+i32(16))
				}
			l8:
				store32(m.memory[int64(uint32(v6))+72:], uint32(v6+i32(68)))
				store32(m.memory[int64(uint32(v6))+76:], uint32(v1))
				t30 := int32(load32(m.memory[uint32(v1):]))
				t31 := int32(load32(m.memory[int64(uint32(v1))+4:]))
				m.fn69(v6+i32(8), t30, t31, v9, v6+i32(72), i32(38))
				t32 := int32(load32(m.memory[uint32(v1):]))
				v11 = t32
				t33 := int32(load32(m.memory[int64(uint32(v6))+12:]))
				v13 = t33
				{
					{
						t34 := int32(load32(m.memory[int64(uint32(v6))+8:]))
						if t34 != i32(1) {
							goto l9
						}
						v16 = v11 + v13
						t35 := int32(m.memory[uint32(v16)])
						v10 = t35
						t36 := int32(load32(m.memory[int64(uint32(v6))+24:]))
						v8 = t36
						t37 := int64(load64(m.memory[int64(uint32(v6))+16:]))
						v12 = t37
						t38 := v16
						v14 = int32(uint32(int32(v9)) >> 25)
						m.memory[uint32(t38)] = byte(v14)
						t39 := int32(load32(m.memory[int64(uint32(v1))+4:]))
						m.memory[uint32(v11+t39&(v13+i32(-8))+i32(8))] = byte(v14)
						t40 := int32(load32(m.memory[int64(uint32(v1))+12:]))
						store32(m.memory[int64(uint32(v1))+12:], uint32(t40+i32(1)))
						t41 := int32(load32(m.memory[int64(uint32(v1))+8:]))
						store32(m.memory[int64(uint32(v1))+8:], uint32(t41-v10&i32(1)))
						v11 = v11 - v13<<4
						v13 = v11 + i32(-16)
						store64(m.memory[uint32(v13):], uint64(v12))
						store32(m.memory[int64(uint32(v13))+8:], uint32(v8))
						store32(m.memory[uint32(v11+i32(-4)):], uint32(v7))
						goto l10
					}
				l9:
					store32(m.memory[uint32(v11-v13<<4+i32(-4)):], uint32(v7))
					t42 := int32(load32(m.memory[int64(uint32(v6))+16:]))
					t43 := int32(load32(m.memory[int64(uint32(v6))+20:]))
					m.fn16(t42, t43)
				}
			l10:
				t44 := int32(load32(m.memory[int64(uint32(v2))+8:]))
				store32(m.memory[int64(uint32(v6))+24:], uint32(t44))
				t45 := int64(load64(m.memory[uint32(v2):]))
				store64(m.memory[int64(uint32(v6))+16:], uint64(t45))
				t46 := int64(load64(m.memory[uint32(v3):]))
				store64(m.memory[int64(uint32(v6))+28:], uint64(t46))
				t47 := int32(load32(m.memory[int64(uint32(v3))+8:]))
				store32(m.memory[int64(uint32(v6))+36:], uint32(t47))
				m.fn51(v6+i32(40), v4, v5)
				{
					t48 := int32(load32(m.memory[int64(uint32(v1))+44:]))
					v3 = t48
					t49 := int32(load32(m.memory[int64(uint32(v1))+36:]))
					if v3 != t49 {
						goto l11
					}
					m.fn1144(v1 + i32(36))
				}
			l11:
				t50 := int32(load32(m.memory[int64(uint32(v1))+40:]))
				v2 = t50 + v3*i32(40)
				memory_copy(m.memory, uint32(v2), uint32(v6+i32(16)), uint32(i32(36)))
				store32(m.memory[int64(uint32(v2))+36:], uint32(v7))
				store32(m.memory[int64(uint32(v1))+44:], uint32(v3+i32(1)))
				store32(m.memory[uint32(v0):], uint32(i32(-1)))
				store32(m.memory[int64(uint32(v0))+4:], uint32(v7))
				goto l12
			}
		l7:
			m.fn51(v0+i32(4), i32(1075096), i32(45))
			store32(m.memory[int64(uint32(v0))+20:], uint32(i32(21)))
			store32(m.memory[int64(uint32(v0))+16:], uint32(i32(1075141)))
			store32(m.memory[uint32(v0):], uint32(i32(-0x7ffffffd)))
			t51 := int32(load32(m.memory[int64(uint32(v3))+4:]))
			v7 = t51
		}
	l5:
		t52 := int32(load32(m.memory[uint32(v3):]))
		m.fn16(t52, v7)
		t53 := int32(load32(m.memory[uint32(v2):]))
		t54 := int32(load32(m.memory[int64(uint32(v2))+4:]))
		m.fn16(t53, t54)
	}
l12:
	m.g0 = v6 + i32(80)
}
func (m *Module) fn1297(v0 int32) {
	{
		t0 := int32(load32(m.memory[uint32(v0):]))
		if t0 != i32(-1) {
			goto l0
		}
		t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		t2 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		m.fn16(t1, t2)
		return
	}
l0:
	m.fn785(v0)
}
func (m *Module) fn1298(v0, v1 int32) {
	var v2 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		v2 = t0
		t1 := int32(load32(m.memory[uint32(v0):]))
		if v2 != t1 {
			goto l0
		}
		m.fn418(v0)
	}
l0:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v2+i32(1)))
	t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	v0 = t2 + v2*i32(20)
	t3 := int64(load64(m.memory[uint32(v1):]))
	store64(m.memory[uint32(v0):], uint64(t3))
	t4 := int64(load64(m.memory[int64(uint32(v1))+8:]))
	store64(m.memory[int64(uint32(v0))+8:], uint64(t4))
	t5 := int32(load32(m.memory[int64(uint32(v1))+16:]))
	store32(m.memory[int64(uint32(v0))+16:], uint32(t5))
}
func (m *Module) fn1299(v0, v1, v2 int32) {
	var v3, v4 int32
	t0 := m.g0
	v3 = t0 - i32(16)
	m.g0 = v3
	if v1 < i32(0) {
		store32(m.memory[int64(uint32(v0))+4:], uint32(i32(0)))
		v4 = i32(1)
		goto l2
	}
	if v1 != 0 {
		goto l1
	}
	store64(m.memory[int64(uint32(v0))+4:], uint64(i64(0x100000000)))
	v4 = i32(0)
	goto l2
l1:
	v4 = i32(1)
	m.fn1819(v3+i32(8), i32(1), v1, v2)
	{
		t1 := int32(load32(m.memory[int64(uint32(v3))+8:]))
		v2 = t1
		if v2 == 0 {
			goto l3
		}
		store32(m.memory[int64(uint32(v0))+8:], uint32(v2))
		store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
		v4 = i32(0)
		goto l2
	}
l3:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v1))
	store32(m.memory[int64(uint32(v0))+4:], uint32(i32(1)))
l2:
	store32(m.memory[uint32(v0):], uint32(v4))
	m.g0 = v3 + i32(16)
}
func (m *Module) fn1300(v0 int32) {
	var v1, v2 int32
	t0 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	v1 = t0
	v2 = v1 + i32(8)
	t1 := int32(load32(m.memory[int64(uint32(v0))+12:]))
	t2 := int32(uint32(t1-v1) / uint32(i32(20)))
	v1 = t2
l1:
	if v1 == 0 {
		goto l0
	}
	v1 = v1 + i32(-1)
	m.fn969(v2)
	v2 = v2 + i32(20)
	goto l1
l0:
	t3 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	t4 := int32(load32(m.memory[uint32(v0):]))
	m.fn136(t3, t4, i32(4), i32(20))
}
func (m *Module) fn1301(v0, v1, v2, v3 int32) {
	var v4 int32
	t0 := m.g0
	v4 = t0 - i32(16)
	m.g0 = v4
	if v0 != 0 {
		goto l0
	}
	v0 = i32(0)
	v3 = v4 + i32(12)
	goto l1
l0:
	store32(m.memory[int64(uint32(v4))+12:], uint32(v2))
	v0 = v0 * v3
	v3 = v4 + i32(8)
l1:
	store32(m.memory[uint32(v3):], uint32(v0))
	{
		t1 := int32(load32(m.memory[int64(uint32(v4))+12:]))
		v0 = t1
		if v0 == 0 {
			goto l2
		}
		t2 := int32(load32(m.memory[int64(uint32(v4))+8:]))
		m.fn40(v1, v0, t2)
	}
l2:
	m.g0 = v4 + i32(16)
}
