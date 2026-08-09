package core

import (
	"math/bits"
)

func (m *Module) fn312(v0, v1, v2, v3 int32) {
	var v4, v5 int32
	t0 := m.g0
	v4 = t0 - i32(16)
	m.g0 = v4
l6:
	{
		if v3 == 0 {
			m.memory[uint32(v0)] = byte(i32(255))
			goto l3
		}
		m.fn298(v4+i32(8), v1, v2, v3)
		{
			t1 := int32(m.memory[int64(uint32(v4))+8])
			v5 = t1
			if v5 == i32(255) {
				t4 := int32(load32(m.memory[int64(uint32(v4))+12:]))
				v5 = t4
				if v5 == 0 {
					goto l4
				}
				if uint32(v3) < uint32(v5) {
					m.fn151(v5, v3, v3, i32(1072408))
					panic("unreachable")
				}
				v2 = v2 + v5
				v3 = v3 - v5
				goto l6
			}
			t2 := m.fn313(v4 + i32(8))
			if t2 != 0 {
				t6 := int32(load32(m.memory[int64(uint32(v4))+12:]))
				m.fn119(v5, t6)
				goto l6
			}
			t3 := int64(load64(m.memory[int64(uint32(v4))+8:]))
			store64(m.memory[uint32(v0):], uint64(t3))
			goto l3
		}
	l4:
		t5 := int64(load64(m.memory[int64(uint32(i32(0)))+1287056:]))
		store64(m.memory[uint32(v0):], uint64(t5))
	}
l3:
	m.g0 = v4 + i32(16)
}
func (m *Module) fn313(v0 int32) int32 {
	var v1 int32
	v1 = i32(0)
	{
		t0 := int32(m.memory[uint32(v0)])
		switch t0 {
		case 1:
			t1 := int32(m.memory[int64(uint32(v0))+1])
			var p2 int32
			if t1 == i32(35) {
				p2 = 1
			}
			return p2
		case 2:
			t3 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t4 := int32(m.memory[int64(uint32(t3))+8])
			var p5 int32
			if t4 == i32(35) {
				p5 = 1
			}
			return p5
		case 3:
			t6 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t7 := int32(m.memory[int64(uint32(t6))+8])
			var p8 int32
			if t7 == i32(35) {
				p8 = 1
			}
			v1 = p8
			fallthrough
		default:
			return v1
		}
	}
}
func (m *Module) fn314(v0, v1 int64, v2, v3 int32) int64 {
	var v4 int32
	t0 := m.g0
	v4 = t0 - i32(64)
	m.g0 = v4
	store64(m.memory[int64(uint32(v4))+48:], uint64(i64(0)))
	store64(m.memory[int64(uint32(v4))+56:], uint64(i64(0)))
	store64(m.memory[int64(uint32(v4))+40:], uint64(v1))
	store64(m.memory[int64(uint32(v4))+24:], uint64(v1^i64(8387220255154660723)))
	store64(m.memory[int64(uint32(v4))+16:], uint64(v1^i64(7237128888997146477)))
	store64(m.memory[int64(uint32(v4))+32:], uint64(v0))
	store64(m.memory[int64(uint32(v4))+8:], uint64(v0^i64(0x6c7967656e657261)))
	store64(m.memory[uint32(v4):], uint64(v0^i64(8317987319222330741)))
	m.fn172(v4, v3)
	m.fn285(v4, v2, v3)
	t1 := m.fn174(v4)
	v1 = t1
	m.g0 = v4 + i32(64)
	return v1
}
func (m *Module) fn315(v0 int32) int64 {
	var v1 int64
	var v2 int32
	v1 = i64(0)
	v2 = i32(0)
l1:
	{
		if v2 == i32(4) {
			return v1
		}
		t0 := int64(m.memory[uint32(v0+v2)])
		v1 = i64_shl(i64(1), t0) | v1
		v2 = v2 + i32(1)
		goto l1
	}
}
func (m *Module) fn316(v0, v1 int32) {
	var v2, v3, v4, v5, v6, v7, v8 int32
	v2 = i32(3)
	v3 = i32(4)
l8:
	v4 = v3
	v3 = v2
	v5 = i32(1)
l7:
	v2 = i32(0)
l9:
	{
		if uint32(v2) >= uint32(v3) {
			goto l0
		}
		{
			t0 := v4
			v6 = v2 ^ i32(-1)
			v7 = t0 + v6
			if uint32(v7) > uint32(i32(3)) {
				m.fn158(v7, i32(4), i32(1280992))
				panic("unreachable")
			}
			{
				v6 = v3 + v6
				if uint32(v6) >= uint32(i32(4)) {
					m.fn158(v6, i32(4), i32(1281008))
					panic("unreachable")
				}
				t1 := int32(m.memory[int64(uint32(v6))+1071308])
				v8 = t1
				t2 := int32(m.memory[int64(uint32(v7))+1071308])
				v7 = t2
				{
					{
						if v1 == 0 {
							goto l3
						}
						v8 = v8 & i32(255)
						t3 := v8
						v7 = v7 & i32(255)
						if uint32(t3) > uint32(v7) {
							goto l4
						}
						if uint32(v8) < uint32(v7) {
							goto l5
						}
						goto l6
					}
				l3:
					v8 = v8 & i32(255)
					t4 := v8
					v7 = v7 & i32(255)
					if uint32(t4) < uint32(v7) {
						goto l4
					}
					if uint32(v8) <= uint32(v7) {
						goto l6
					}
				}
			l5:
				v5 = v4 - v6
				v3 = v6
				goto l7
			}
		l4:
			v2 = v3 + i32(-1)
			goto l8
		}
	l6:
		v2 = v2 + i32(1)
		t5 := v2
		var p6 int32
		if v2 == v5 {
			p6 = 1
		}
		v6 = p6
		p7 := t5
		if v6 != 0 {
			p7 = i32(0)
		}
		v2 = p7
		t9 := v3
		p8 := i32(0)
		if v6 != 0 {
			p8 = v5
		}
		v3 = t9 - p8
		goto l9
	}
l0:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
	store32(m.memory[uint32(v0):], uint32(v4))
}
func (m *Module) fn317(v0, v1, v2, v3 int32) {
	if uint32(v2) < uint32(i32(5)) {
		goto l0
	}
	m.fn91(i32(1301172), i32(19), v3)
	panic("unreachable")
l0:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v2))
	store32(m.memory[uint32(v0):], uint32(v1))
	store32(m.memory[int64(uint32(v0))+12:], uint32(i32(4)-v2))
	store32(m.memory[int64(uint32(v0))+8:], uint32(v1+v2))
}
func (m *Module) fn318(v0, v1, v2 int64, v3 int32) int64 {
	var v4 int32
	t0 := m.g0
	v4 = t0 - i32(32)
	m.g0 = v4
	{
		if uint64(v1) <= uint64(v2) {
			m.g0 = v4 + i32(32)
			t3 := v1
			p2 := v2
			if uint64(v0) < uint64(v2) {
				p2 = v0
			}
			p4 := p2
			if uint64(v0) < uint64(v1) {
				p4 = t3
			}
			return p4
		}
		store64(m.memory[uint32(v4):], uint64(v1))
		store64(m.memory[int64(uint32(v4))+8:], uint64(v2))
		t1 := v4
		v2 = int64(uint32(i32(58))) << 32
		store64(m.memory[int64(uint32(t1))+24:], uint64(v2|int64(uint32(v4+i32(8)))))
		store64(m.memory[int64(uint32(v4))+16:], uint64(v2|int64(uint32(v4))))
		m.fn91(i32(1051106), v4+i32(16), v3)
		panic("unreachable")
	}
}
func (m *Module) fn319(v0, v1, v2 int32) {
	var v3 int32
	t0 := m.g0
	v3 = t0 - i32(16)
	m.g0 = v3
	m.fn363(v3+i32(8), v2, v1, i32(1024), i32(1070512))
	t1 := int32(load32(m.memory[int64(uint32(v3))+12:]))
	v1 = t1
	t2 := int32(load32(m.memory[int64(uint32(v3))+8:]))
	store32(m.memory[uint32(v0):], uint32(t2))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
	m.g0 = v3 + i32(16)
}
func (m *Module) fn320(v0, v1, v2 int32) int32 {
	var v3, v4, v5 int32
	v3 = v2 & i32(3)
	t0 := v0
	v4 = v2 & i32(0x7ffffffc)
	v5 = t0 + v4
	v4 = v1 + v4
	var _ int32
l5:
	{
		{
			if uint32(v2) > uint32(i32(3)) {
				goto l0
			}
			v2 = i32(1)
			{
				if uint32(v3) <= uint32(i32(1)) {
					goto l1
				}
				t2 := int32(load16(m.memory[uint32(v5):]))
				t3 := int32(load16(m.memory[uint32(v4):]))
				if t2 != t3 {
					goto l2
				}
				v3 = v3 + i32(-2)
				v4 = v4 + i32(2)
				v5 = v5 + i32(2)
			}
		l1:
			if v3 == 0 {
				goto l3
			}
			t4 := int32(m.memory[uint32(v5)])
			t5 := int32(m.memory[uint32(v4)])
			var p6 int32
			if t4 == t5 {
				p6 = 1
			}
			return p6
		}
	l0:
		t7 := int32(load32(m.memory[uint32(v0):]))
		t8 := int32(load32(m.memory[uint32(v1):]))
		if t7 == t8 {
			v2 = v2 + i32(-4)
			v1 = v1 + i32(4)
			v0 = v0 + i32(4)
			goto l5
		}
	}
l2:
	v2 = i32(0)
l3:
	return v2
}
func (m *Module) fn321(v0, v1 int32) {
	var v2, v3, v4 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	m.fn1(v2+i32(4), v1, i32(1), i32(1), i32(1))
	t1 := int32(load32(m.memory[int64(uint32(v2))+8:]))
	v3 = t1
	{
		t2 := int32(load32(m.memory[int64(uint32(v2))+4:]))
		if t2 != i32(1) {
			goto l0
		}
		t3 := int32(load32(m.memory[int64(uint32(v2))+12:]))
		m.fn2(v3, t3)
		panic("unreachable")
	}
l0:
	t4 := int32(load32(m.memory[int64(uint32(v2))+12:]))
	v4 = t4
	store32(m.memory[int64(uint32(v0))+8:], uint32(v1))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v4))
	store32(m.memory[uint32(v0):], uint32(v3))
	m.g0 = v2 + i32(16)
}
func (m *Module) fn322(v0, v1 int32) {
	var v2, v3 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	{
		{
			t1 := int32(load32(m.memory[uint32(v1):]))
			t2 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			v3 = t2
			if uint32(t1) <= uint32(v3) {
				goto l0
			}
			m.fn477(v2+i32(8), v1, v3, i32(1), i32(1))
			t3 := int32(load32(m.memory[int64(uint32(v2))+8:]))
			v3 = t3
			if v3 != i32(-1) {
				t6 := int32(load32(m.memory[int64(uint32(v2))+12:]))
				m.fn2(v3, t6)
				panic("unreachable")
			}
			t4 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			v3 = t4
		}
	l0:
		store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
		t5 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		store32(m.memory[uint32(v0):], uint32(t5))
		m.g0 = v2 + i32(16)
		return
	}
}
func (m *Module) fn323(v0 int32) int32 {
	var v1 int32
	t0 := m.g0
	v1 = t0 - i32(1152)
	m.g0 = v1
	{
		t1 := int32(load32(m.memory[uint32(v0):]))
		if t1 != i32(2) {
			goto l0
		}
		memory_zero(m.memory, uint32(v1+i32(88)), uint32(i32(1024)))
		m.fn362(v1+i32(16), i32(1287584))
		store64(m.memory[int64(uint32(v1))+1112:], uint64(i64(0)))
		store32(m.memory[int64(uint32(v1))+8:], uint32(i32(0)))
		store64(m.memory[int64(uint32(v1))+1120:], uint64(i64(0)))
		store64(m.memory[int64(uint32(v1))+1128:], uint64(i64(0)))
		m.memory[int64(uint32(v1))+1144] = byte(i32(2))
		memory_copy(m.memory, uint32(v0), uint32(v1+i32(8)), uint32(i32(1144)))
	}
l0:
	m.g0 = v1 + i32(1152)
	return v0
}
func (m *Module) fn324(v0, v1 int32, v2, v3 int64) int32 {
	var v4 int32
	t0 := m.g0
	v4 = t0 - i32(80)
	m.g0 = v4
	m.fn362(v4+i32(8), v1)
	m.fn328(v0 + i32(64))
	memory_copy(m.memory, uint32(v0+i32(8)), uint32(v4+i32(8)), uint32(i32(72)))
	store64(m.memory[int64(uint32(v0))+1120:], uint64(v3))
	store64(m.memory[int64(uint32(v0))+1112:], uint64(v2))
	store64(m.memory[int64(uint32(v0))+1104:], uint64(v2))
	store32(m.memory[uint32(v0):], uint32(i32(0)))
	m.g0 = v4 + i32(80)
	return v0
}
func (m *Module) fn325(v0, v1, v2 int32) {
	var v3, v4 int32
	var v5, v6 int64
	var v7, v8, v9, v10 int32
	var v11 int64
	var v12 int32
	t0 := m.g0
	v3 = t0 - i32(48)
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
			store32(m.memory[int64(uint32(v3))+28:], uint32(i32(0)))
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
				m.fn117(v3+i32(32), v2, v3+i32(28), i32(4))
				{
					t6 := int32(m.memory[int64(uint32(v3))+32])
					if t6 == i32(255) {
						goto l2
					}
					t7 := int64(load64(m.memory[int64(uint32(v3))+32:]))
					t8 := v3
					v5 = t7
					store64(m.memory[int64(uint32(t8))+40:], uint64(v5))
					t9 := m.fn118(v3 + i32(40))
					if t9&i32(255) == i32(37) {
						goto l3
					}
					store64(m.memory[int64(uint32(v0))+8:], uint64(v5))
					store64(m.memory[uint32(v0):], uint64(i64(-0x7fffffffffffffff)))
					goto l4
				}
			l2:
				t10 := int32(load32(m.memory[int64(uint32(v1))+72:]))
				if t10 != i32(4) {
					goto l1
				}
				t11 := int32(load32(m.memory[int64(uint32(v1))+68:]))
				t12 := int32(load32(m.memory[uint32(t11):]))
				t13 := int32(load32(m.memory[int64(uint32(v3))+28:]))
				if t12 == t13 {
					store64(m.memory[int64(uint32(v2))+8:], uint64(v5))
					m.memory[int64(uint32(v1))+1136] = byte(i32(2))
					store64(m.memory[int64(uint32(v0))+16:], uint64(v5))
					store64(m.memory[int64(uint32(v0))+8:], uint64(i64(1)))
					store32(m.memory[uint32(v0):], uint32(i32(0)))
					goto l4
				}
			}
		l1:
			if v4&i32(1) == 0 {
				goto l6
			}
			store32(m.memory[uint32(v0):], uint32(i32(0)))
			store64(m.memory[int64(uint32(v0))+8:], uint64(i64(0)))
			goto l4
		l3:
			v7 = int32(int64(uint64(v5) >> 32))
			v8 = int32(v5)
			if v4&i32(1) != 0 {
				store32(m.memory[uint32(v0):], uint32(i32(0)))
				store64(m.memory[int64(uint32(v0))+8:], uint64(i64(0)))
				m.fn119(v8, v7)
				goto l4
			}
			m.fn119(v8, v7)
		l6:
			m.memory[int64(uint32(v1))+1136] = byte(i32(2))
			goto l0
		}
	l0:
		v9 = v1 + i32(8)
		v10 = v1 + i32(80)
		t14 := int64(load64(m.memory[int64(uint32(v1))+1104:]))
		v5 = t14
	l17:
		{
			t15 := int64(load64(m.memory[int64(uint32(v1))+1112:]))
			if uint64(v5) < uint64(t15) {
				goto l8
			}
			t16 := int64(load64(m.memory[int64(uint32(v1))+1120:]))
			t17 := v5
			v6 = t16
			if uint64(t17) >= uint64(v6) {
				goto l8
			}
			v11 = v5 + i64(1024)
			p18 := v11
			if uint64(v11) < uint64(v5) {
				p18 = i64(-1)
			}
			v11 = p18
			if uint64(v11) <= uint64(v5) {
				goto l8
			}
			t20 := v3 + i32(16)
			t21 := v10
			p19 := v11
			if uint64(v6) < uint64(v11) {
				p19 = v6
			}
			m.fn319(t20, t21, int32(p19-v5))
			t22 := int32(load32(m.memory[int64(uint32(v3))+20:]))
			v4 = t22
			t23 := int32(load32(m.memory[int64(uint32(v3))+16:]))
			v7 = t23
			{
				{
					t24 := int32(load32(m.memory[uint32(v1):]))
					if t24 != 0 {
						goto l9
					}
					store64(m.memory[int64(uint32(v2))+8:], uint64(v5))
					m.fn117(v3+i32(32), v2, v7, v4)
					{
						t25 := int32(m.memory[int64(uint32(v3))+32])
						if t25 == i32(255) {
							goto l10
						}
						t26 := int64(load64(m.memory[int64(uint32(v3))+32:]))
						t27 := v3
						v5 = t26
						store64(m.memory[int64(uint32(t27))+40:], uint64(v5))
						{
							t28 := m.fn118(v3 + i32(40))
							if t28&i32(255) == i32(37) {
								m.fn119(int32(v5), int32(int64(uint64(v5)>>32)))
								goto l8
							}
							store64(m.memory[int64(uint32(v0))+8:], uint64(v5))
							store64(m.memory[uint32(v0):], uint64(i64(-0x7fffffffffffffff)))
							goto l4
						}
					}
				l10:
					v8 = i32(0)
					t29 := int32(load32(m.memory[uint32(v1):]))
					if t29 == 0 {
						goto l12
					}
				}
			l9:
				t30 := int32(load32(m.memory[int64(uint32(v1))+4:]))
				t31 := v4
				v8 = t30
				if uint32(t31) < uint32(v8) {
					m.fn151(v8, v4, v4, i32(1286884))
					panic("unreachable")
				}
				v7 = v7 + v8
				v4 = v4 - v8
			}
		l12:
			store64(m.memory[int64(uint32(v3))+40:], uint64(i64(1)))
			{
				t32 := int32(load32(m.memory[int64(uint32(v1))+72:]))
				t33 := v4
				v12 = t32
				if uint32(t33) < uint32(v12) {
					goto l14
				}
				t34 := int32(load32(m.memory[int64(uint32(v1))+68:]))
				t35 := int32(load32(m.memory[int64(uint32(v1))+56:]))
				m.t0[uint(t35)].(func(int32, int32, int32, int32, int32, int32, int32))(v3+i32(8), v9, v3+i32(40), v7, v4, t34, v12)
				t36 := int32(load32(m.memory[int64(uint32(v3))+8:]))
				if t36&i32(1) != 0 {
					goto l15
				}
			}
		l14:
			store32(m.memory[uint32(v1):], uint32(i32(0)))
			{
				t37 := int64(load64(m.memory[int64(uint32(v1))+1104:]))
				v5 = t37
				v6 = v5 + i64(1021)
				p38 := v6
				if uint64(v6) < uint64(v5) {
					p38 = i64(-1)
				}
				v5 = p38
				t39 := int64(load64(m.memory[int64(uint32(v1))+1120:]))
				t40 := v5
				v6 = t39
				if uint64(t40) >= uint64(v6) {
					goto l16
				}
				store64(m.memory[int64(uint32(v1))+1104:], uint64(v5))
				goto l17
			}
		l16:
		}
		store64(m.memory[int64(uint32(v1))+1112:], uint64(v6))
	l8:
		store32(m.memory[uint32(v0):], uint32(i32(0)))
		store64(m.memory[int64(uint32(v0))+8:], uint64(i64(0)))
		goto l4
	l15:
		t41 := int32(load32(m.memory[int64(uint32(v3))+12:]))
		v4 = t41
		store32(m.memory[uint32(v1):], uint32(i32(1)))
		store64(m.memory[int64(uint32(v0))+8:], uint64(i64(1)))
		store32(m.memory[uint32(v0):], uint32(i32(0)))
		store32(m.memory[int64(uint32(v1))+4:], uint32(v8+v4+i32(4)))
		t42 := v2
		v5 = v5 + int64(uint32(v8)) + int64(uint32(v4))
		store64(m.memory[int64(uint32(t42))+8:], uint64(v5))
		store64(m.memory[int64(uint32(v0))+16:], uint64(v5))
	}
l4:
	m.g0 = v3 + i32(48)
}
func (m *Module) fn326(v0 int32) {
	t0 := int32(load32(m.memory[uint32(v0):]))
	if t0 == i32(-2) {
		return
	}
	m.fn116(v0)
}
func (m *Module) fn327(v0, v1, v2 int32) {
	{
		t0 := int32(load32(m.memory[uint32(v1):]))
		if t0 == i32(-2) {
			goto l0
		}
		t1 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		store32(m.memory[int64(uint32(v0))+8:], uint32(t1))
		t2 := int64(load64(m.memory[uint32(v1):]))
		store64(m.memory[uint32(v0):], uint64(t2))
		m.fn116(v2)
		return
	}
l0:
	t3 := int32(load32(m.memory[int64(uint32(v2))+8:]))
	store32(m.memory[int64(uint32(v0))+8:], uint32(t3))
	t4 := int64(load64(m.memory[uint32(v2):]))
	store64(m.memory[uint32(v0):], uint64(t4))
}
func (m *Module) fn328(v0 int32) {
	t0 := int32(load32(m.memory[uint32(v0):]))
	if t0 == 0 {
		return
	}
	t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t2 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	m.fn128(t1, t2)
}
func (m *Module) fn329(v0 int32) {
	var v1, v2 int32
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
			t3 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			m.fn344(t2, t3)
			return
		case 1:
			if v1 == i32(-1) {
				return
			}
			t4 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			m.fn16(v1, t4)
		}
	}
}
func (m *Module) fn330(v0, v1, v2 int32) {
	var v3, v4 int32
	var v5 int64
	t0 := m.g0
	v3 = t0 - i32(32)
	m.g0 = v3
	m.fn321(v3+i32(8), v2)
	m.fn322(v3, v3+i32(8))
	t1 := int32(load32(m.memory[uint32(v3):]))
	t2 := v3 + i32(24)
	t3 := v1
	v2 = t1
	t4 := int32(load32(m.memory[int64(uint32(v3))+4:]))
	t5 := v2
	v4 = t4
	m.fn117(t2, t3, t5, v4)
	{
		t6 := int32(m.memory[int64(uint32(v3))+24])
		if t6 == i32(255) {
			goto l0
		}
		t7 := int64(load64(m.memory[int64(uint32(v3))+24:]))
		t8 := v3
		v5 = t7
		store64(m.memory[int64(uint32(t8))+8:], uint64(v5))
		{
			t9 := m.fn118(v3 + i32(8))
			if t9&i32(255) == i32(37) {
				goto l1
			}
			store64(m.memory[int64(uint32(v0))+4:], uint64(v5))
			v1 = i32(-0x80000000)
			goto l2
		}
	l1:
		store32(m.memory[int64(uint32(v0))+8:], uint32(i32(50)))
		store32(m.memory[int64(uint32(v0))+4:], uint32(i32(1071188)))
		m.fn119(int32(v5), int32(int64(uint64(v5)>>32)))
		v1 = i32(-1)
	l2:
		store32(m.memory[uint32(v0):], uint32(v1))
		m.fn128(v2, v4)
		goto l3
	}
l0:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v4))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v2))
	store32(m.memory[uint32(v0):], uint32(i32(-2)))
l3:
	m.g0 = v3 + i32(32)
}
func (m *Module) fn331(v0, v1, v2 int32) {
	var v3, v4, v5, v6, v7, v8, v9, v10, v11 int32
	var v12 int64
	t0 := m.g0
	v3 = t0 - i32(64)
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
				store32(m.memory[int64(uint32(v3))+60:], uint32(i32(0)))
				store64(m.memory[int64(uint32(v3))+52:], uint64(i64(0x100000000)))
				m.fn47(v3+i32(52), v2)
			l6:
				{
					{
						t2 := int32(int8(m.memory[uint32(v1)]))
						v4 = t2
						if v4 > i32(-1) {
							goto l2
						}
						t3 := int32(load32(m.memory[int64(uint32(v4&i32(127)<<2))+1302492:]))
						v4 = t3
					}
				l2:
					t4 := int32(load32(m.memory[int64(uint32(v3))+60:]))
					v5 = t4
					t5 := v3 + i32(52)
					var p6 int32
					if uint32(v4) < uint32(i32(2048)) {
						p6 = 1
					}
					v6 = p6
					p7 := i32(3)
					if v6 != 0 {
						p7 = i32(2)
					}
					var p8 int32
					if uint32(v4) < uint32(i32(128)) {
						p8 = 1
					}
					v7 = p8
					p9 := p7
					if v7 != 0 {
						p9 = i32(1)
					}
					v8 = p9
					m.fn47(t5, v8)
					t10 := int32(load32(m.memory[int64(uint32(v3))+56:]))
					t11 := v5
					v9 = t10
					v10 = t11 + v9
					if v7 != 0 {
						goto l3
					}
					v7 = int32(uint32(v4) >> 6)
					v11 = v4&i32(63) | i32(-128)
					if v6 == 0 {
						m.memory[int64(uint32(v10))+2] = byte(v11)
						m.memory[int64(uint32(v10))+1] = byte(v7 | i32(128))
						m.memory[uint32(v10)] = byte(int32(uint32(v4)>>12) | i32(224))
						goto l5
					}
					m.memory[int64(uint32(v10))+1] = byte(v11)
					m.memory[uint32(v10)] = byte(v7 | i32(192))
					goto l5
				l3:
					m.memory[uint32(v10)] = byte(v4)
				l5:
					t12 := v3
					v4 = v8 + v5
					store32(m.memory[int64(uint32(t12))+60:], uint32(v4))
					v1 = v1 + i32(1)
					v2 = v2 + i32(-1)
					if v2 != 0 {
						goto l6
					}
				}
				t13 := int32(load32(m.memory[int64(uint32(v3))+52:]))
				v1 = t13
				goto l7
			}
		l0:
			m.fn12(v3+i32(20), v1, v2)
			t14 := int32(load32(m.memory[int64(uint32(v3))+20:]))
			if t14 != 0 {
				goto l8
			}
			t15 := int32(load32(m.memory[int64(uint32(v3))+28:]))
			v4 = t15
			t16 := int32(load32(m.memory[int64(uint32(v3))+24:]))
			v9 = t16
			v1 = i32(-1)
		}
	l7:
		store32(m.memory[int64(uint32(v0))+8:], uint32(v4))
		store32(m.memory[int64(uint32(v0))+4:], uint32(v9))
		store32(m.memory[uint32(v0):], uint32(v1))
		goto l9
	l8:
		t17 := int64(load64(m.memory[int64(uint32(v3))+24:]))
		store64(m.memory[int64(uint32(v3))+32:], uint64(t17))
		store32(m.memory[int64(uint32(v3))+48:], uint32(i32(59)))
		store32(m.memory[int64(uint32(v3))+44:], uint32(v3+i32(32)))
		m.fn341(v3+i32(52), i32(1052496), v3+i32(44))
		m.fn580(v3+i32(12), i32(21), v3+i32(52))
		t18 := int64(load64(m.memory[int64(uint32(v3))+12:]))
		v12 = t18
		store32(m.memory[uint32(v0):], uint32(i32(-2)))
		store64(m.memory[int64(uint32(v0))+4:], uint64(v12))
	}
l9:
	m.g0 = v3 + i32(64)
}
func (m *Module) fn332(v0, v1 int32) {
	var v2, v3, v4 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	{
		{
			t1 := int32(load32(m.memory[uint32(v1):]))
			if t1 == i32(-1) {
				goto l0
			}
			m.fn108(v2+i32(8), v1)
			t2 := int32(load32(m.memory[int64(uint32(v2))+12:]))
			v1 = t2
			t3 := int32(load32(m.memory[int64(uint32(v2))+8:]))
			v3 = t3
			goto l1
		}
	l0:
		t4 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		v4 = t4
		{
			{
				t5 := int32(load32(m.memory[int64(uint32(v1))+8:]))
				v1 = t5
				if v1 != 0 {
					goto l2
				}
				v3 = i32(1)
				goto l3
			}
		l2:
			m.fn247(v2, i32(1), v1)
			t6 := int32(load32(m.memory[uint32(v2):]))
			v3 = t6
			if v3 == 0 {
				m.fn85(i32(1), v1)
				panic("unreachable")
			}
		}
	l3:
		if v1 == 0 {
			goto l1
		}
		memory_copy(m.memory, uint32(v3), uint32(v4), uint32(v1))
	}
l1:
	store32(m.memory[uint32(v0):], uint32(v3))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
	m.g0 = v2 + i32(16)
}
func (m *Module) fn333(v0, v1 int32) int32 {
	var v2 int32
	v2 = i32(0)
	{
		t0 := int32(m.memory[int64(uint32(v0))+1])
		if uint32(t0) > uint32(v1&i32(255)) {
			goto l0
		}
		t1 := int32(m.memory[int64(uint32(v0))+2])
		v2 = t1
		{
			t2 := int32(m.memory[uint32(v0)])
			if t2 != 0 {
				goto l1
			}
			var p3 int32
			if uint32(v1&i32(255)) <= uint32(v2&i32(255)) {
				p3 = 1
			}
			return p3
		}
	l1:
		;
		var p4 int32
		if uint32(v1&i32(255)) < uint32(v2&i32(255)) {
			p4 = 1
		}
		v2 = p4
	}
l0:
	return v2
}
func (m *Module) fn334(v0, v1, v2 int32) {
	var v3, v4, v5 int32
	t0 := m.g0
	v3 = t0 - i32(16)
	m.g0 = v3
	{
		if uint32(v2) >= uint32(i32(0x7ffffff5)) {
			m.fn97(i32(1291936), i32(43), v3+i32(15), i32(1070180), i32(1070244))
			panic("unreachable")
		}
		t1 := v3
		v4 = (v2 + i32(11)) & i32(0x7ffffffc)
		m.fn1824(t1, v4)
		t2 := int32(load32(m.memory[uint32(v3):]))
		v5 = t2
		if v5 == 0 {
			m.fn85(i32(4), v4)
			panic("unreachable")
		}
		store64(m.memory[uint32(v5):], uint64(i64(0x100000001)))
		if v2 == 0 {
			goto l2
		}
		memory_copy(m.memory, uint32(v5+i32(8)), uint32(v1), uint32(v2))
	l2:
		if v2 == 0 {
			goto l3
		}
		m.fn1825(v1, i32(1), v2)
	l3:
		store32(m.memory[int64(uint32(v0))+4:], uint32(v2))
		store32(m.memory[uint32(v0):], uint32(v5))
		m.g0 = v3 + i32(16)
		return
	}
}
func (m *Module) fn335(v0, v1 int32) {
	var v2 int32
	var v3 int64
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	store16(m.memory[int64(uint32(v2))+6:], uint16(i32(0)))
	m.fn117(v2+i32(8), v1, v2+i32(6), i32(2))
	{
		{
			t1 := int32(m.memory[int64(uint32(v2))+8])
			if t1 == i32(255) {
				goto l0
			}
			t2 := int64(load64(m.memory[int64(uint32(v2))+8:]))
			v3 = t2
			if v3&i64(255) == i64(255) {
				goto l0
			}
			store64(m.memory[uint32(v0):], uint64(v3))
			goto l1
		}
	l0:
		t3 := int32(load16(m.memory[int64(uint32(v2))+6:]))
		store16(m.memory[int64(uint32(v0))+2:], uint16(t3))
		m.memory[uint32(v0)] = byte(i32(255))
	}
l1:
	m.g0 = v2 + i32(16)
}
func (m *Module) fn336(v0, v1 int32) {
	var v2 int32
	var v3 int64
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	store32(m.memory[int64(uint32(v2))+4:], uint32(i32(0)))
	m.fn117(v2+i32(8), v1, v2+i32(4), i32(4))
	{
		{
			t1 := int32(m.memory[int64(uint32(v2))+8])
			if t1 == i32(255) {
				goto l0
			}
			t2 := int64(load64(m.memory[int64(uint32(v2))+8:]))
			v3 = t2
			if v3&i64(255) == i64(255) {
				goto l0
			}
			store64(m.memory[uint32(v0):], uint64(v3))
			goto l1
		}
	l0:
		t3 := int32(load32(m.memory[int64(uint32(v2))+4:]))
		store32(m.memory[int64(uint32(v0))+4:], uint32(t3))
		m.memory[uint32(v0)] = byte(i32(255))
	}
l1:
	m.g0 = v2 + i32(16)
}
func (m *Module) fn337(v0, v1 int32) {
	var v2 int32
	var v3 int64
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	store64(m.memory[uint32(v2):], uint64(i64(0)))
	m.fn117(v2+i32(8), v1, v2, i32(8))
	{
		{
			t1 := int32(m.memory[int64(uint32(v2))+8])
			if t1 == i32(255) {
				goto l0
			}
			t2 := int64(load64(m.memory[int64(uint32(v2))+8:]))
			v3 = t2
			if v3&i64(255) == i64(255) {
				goto l0
			}
			store64(m.memory[int64(uint32(v0))+4:], uint64(v3))
			v1 = i32(1)
			goto l1
		}
	l0:
		t3 := int64(load64(m.memory[uint32(v2):]))
		store64(m.memory[int64(uint32(v0))+8:], uint64(t3))
		v1 = i32(0)
	}
l1:
	store32(m.memory[uint32(v0):], uint32(v1))
	m.g0 = v2 + i32(16)
}
func (m *Module) fn338(v0, v1 int32) {
	var v2 int32
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
	store32(m.memory[int64(uint32(v0))+8:], uint32(v2+i32(1)))
	t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	v0 = t2 + v2<<5
	t3 := int64(load64(m.memory[uint32(v1):]))
	store64(m.memory[uint32(v0):], uint64(t3))
	t4 := int64(load64(m.memory[int64(uint32(v1))+8:]))
	store64(m.memory[int64(uint32(v0))+8:], uint64(t4))
	t5 := int64(load64(m.memory[int64(uint32(v1))+16:]))
	store64(m.memory[int64(uint32(v0))+16:], uint64(t5))
	t6 := int64(load64(m.memory[int64(uint32(v1))+24:]))
	store64(m.memory[int64(uint32(v0))+24:], uint64(t6))
}
func (m *Module) fn339(v0, v1 int32) int32 {
	var v2, v3, v4, v5 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	v3 = v2 + i32(16)
	t1 := int32(m.memory[uint32(v0)])
	v0 = t1
	v4 = i32(33) - int32(bits.LeadingZeros32(uint32(int32(uint32(v0)>>1))))
	v5 = v4
l0:
	{
		v3 = v3 + i32(-1)
		t2 := int32(m.memory[uint32(v0&i32(1)+i32(1131457))])
		m.memory[uint32(v3)] = byte(t2)
		v0 = int32(uint32(v0&i32(254)) >> 1)
		v5 = v5 + i32(-1)
		if v5&i32(255) != 0 {
			goto l0
		}
	}
	t3 := m.fn1638(v1, i32(1), i32(1131459), i32(2), v3, v4)
	v0 = t3
	m.g0 = v2 + i32(16)
	return v0
}
func (m *Module) fn340(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	{
		{
			t1 := int32(load16(m.memory[uint32(v0):]))
			v0 = t1
			if uint32(v0) < uint32(i32(1000)) {
				goto l0
			}
			v3 = i32(1)
			t2 := int32(uint32(v0) / uint32(i32(10000)))
			t3 := v2
			t4 := v0
			v4 = t2
			v5 = t4 - v4*i32(10000)
			t5 := int32(uint32(v5&i32(0xffff)) / uint32(i32(100)))
			v6 = t5
			t6 := int32(load16(m.memory[int64(uint32(v6<<1))+1109319:]))
			store16(m.memory[int64(uint32(t3))+12:], uint16(t6))
			t7 := int32(load16(m.memory[int64(uint32((v5-v6*i32(100))&i32(0xffff)<<1))+1109319:]))
			store16(m.memory[int64(uint32(v2))+14:], uint16(t7))
			goto l1
		}
	l0:
		v3 = i32(5)
		v4 = v0
		if uint32(v0) < uint32(i32(10)) {
			goto l1
		}
		t8 := int32(uint32(v0) / uint32(i32(100)))
		t9 := v2
		t10 := v0
		v4 = t8
		t11 := int32(load16(m.memory[int64(uint32((t10-v4*i32(100))&i32(0xffff)<<1))+1109319:]))
		store16(m.memory[int64(uint32(t9))+14:], uint16(t11))
		v3 = i32(3)
	}
l1:
	{
		if v0 == 0 {
			goto l2
		}
		if v4 == 0 {
			goto l3
		}
	l2:
		t12 := v2 + i32(11)
		v3 = v3 + i32(-1)
		t13 := int32(m.memory[int64(uint32(v4<<1))+1109320])
		m.memory[uint32(t12+v3)] = byte(t13)
	}
l3:
	t14 := m.fn1638(v1, i32(1), i32(1), i32(0), v2+i32(11)+v3, i32(5)-v3)
	v0 = t14
	m.g0 = v2 + i32(16)
	return v0
}
func (m *Module) fn341(v0, v1, v2 int32) {
	if v2&i32(1) == 0 {
		goto l0
	}
	m.fn1820(v0, v1, int32(uint32(v2)>>1))
	return
l0:
	m.fn6(v0, v1, v2)
}
func (m *Module) fn342() int32 {
	var v0, v1 int32
	t0 := m.g0
	v0 = t0 - i32(16)
	m.g0 = v0
	m.fn1824(v0+i32(8), i32(12))
	{
		t1 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		v1 = t1
		if v1 != 0 {
			m.g0 = v0 + i32(16)
			return v1
		}
		m.fn85(i32(4), i32(12))
		panic("unreachable")
	}
}
func (m *Module) fn343(v0, v1, v2, v3 int32) {
	var v4 int32
	{
		t0 := m.fn4(i32(12))
		v4 = t0
		if v4 != 0 {
			goto l0
		}
		m.fn85(i32(4), i32(12))
		panic("unreachable")
	}
l0:
	m.memory[int64(uint32(v4))+8] = byte(v1)
	store32(m.memory[int64(uint32(v4))+4:], uint32(v3))
	store32(m.memory[uint32(v4):], uint32(v2))
	store64(m.memory[uint32(v0):], uint64(int64(uint32(v4))<<32|i64(3)))
}
func (m *Module) fn344(v0, v1 int32) {
	var v2 int32
	if v0&i32(255) != i32(3) {
		return
	}
	{
		t0 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		v0 = t0
		t1 := int32(load32(m.memory[uint32(v0):]))
		v2 = t1
		if v2 == 0 {
			goto l1
		}
		t2 := int32(load32(m.memory[uint32(v1):]))
		m.t0[uint(v2)].(func(int32))(t2)
	}
l1:
	{
		t3 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		v2 = t3
		if v2 == 0 {
			goto l2
		}
		t4 := int32(load32(m.memory[uint32(v1):]))
		t5 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		m.fn1825(t4, t5, v2)
	}
l2:
	m.fn10(v1, i32(12), i32(4))
}
func (m *Module) fn345(v0, v1, v2 int32) {
	var v3 int32
	var v4 int64
	var v5, v6 int32
	t0 := m.g0
	v3 = t0 - i32(32)
	m.g0 = v3
	m.fn117(v3+i32(16), v1, v3+i32(15), i32(1))
	{
		t1 := int32(m.memory[int64(uint32(v3))+16])
		if t1 == i32(255) {
			goto l0
		}
		t2 := int64(load64(m.memory[int64(uint32(v3))+16:]))
		v4 = t2
		if v4&i64(255) == i64(255) {
			goto l0
		}
		store64(m.memory[int64(uint32(v0))+8:], uint64(v4))
		store64(m.memory[uint32(v0):], uint64(i64(-0x7fffffffffffffff)))
		goto l1
	}
l0:
	m.fn336(v3+i32(16), v1)
	{
		{
			t3 := int32(m.memory[int64(uint32(v3))+16])
			if t3 != i32(255) {
				goto l2
			}
			t4 := int32(load32(m.memory[int64(uint32(v3))+20:]))
			v5 = t4
			goto l3
		}
	l2:
		t5 := int64(load64(m.memory[int64(uint32(v3))+16:]))
		v4 = t5
		if v4&i64(255) != i64(255) {
			goto l4
		}
		v5 = int32(int64(uint64(v4) >> 32))
	}
l3:
	store32(m.memory[int64(uint32(v3))+24:], uint32(i32(32)))
	store32(m.memory[int64(uint32(v3))+20:], uint32(i32(1285992)))
	store32(m.memory[int64(uint32(v3))+16:], uint32(i32(-1)))
	v2 = v2 & i32(0xffff)
	if uint32(v2) > uint32(i32(4)) {
		m.fn329(v3 + i32(16))
		m.fn321(v3+i32(16), v2+i32(-5))
		m.fn322(v3, v3+i32(16))
		t6 := int32(load32(m.memory[uint32(v3):]))
		t7 := v3 + i32(16)
		t8 := v1
		v2 = t6
		t9 := int32(load32(m.memory[int64(uint32(v3))+4:]))
		t10 := v2
		v6 = t9
		m.fn117(t7, t8, t10, v6)
		{
			t11 := int32(m.memory[int64(uint32(v3))+16])
			if t11 == i32(255) {
				goto l6
			}
			t12 := int64(load64(m.memory[int64(uint32(v3))+16:]))
			v4 = t12
			if v4&i64(255) == i64(255) {
				goto l6
			}
			store64(m.memory[int64(uint32(v0))+8:], uint64(v4))
			store64(m.memory[uint32(v0):], uint64(i64(-0x7fffffffffffffff)))
			m.fn348(v2, v6)
			goto l1
		}
	l6:
		store32(m.memory[int64(uint32(v0))+12:], uint32(v5))
		store32(m.memory[int64(uint32(v0))+8:], uint32(v6))
		store32(m.memory[int64(uint32(v0))+4:], uint32(v2))
		store32(m.memory[uint32(v0):], uint32(i32(0)))
		goto l1
	}
	store32(m.memory[int64(uint32(v0))+12:], uint32(i32(32)))
	store32(m.memory[int64(uint32(v0))+8:], uint32(i32(1285992)))
	store64(m.memory[uint32(v0):], uint64(i64(-0xffffffff)))
	goto l1
l4:
	store64(m.memory[int64(uint32(v0))+8:], uint64(v4))
	store64(m.memory[uint32(v0):], uint64(i64(-0x7fffffffffffffff)))
l1:
	m.g0 = v3 + i32(32)
}
func (m *Module) fn346(v0, v1, v2, v3 int32) {
	var v4 int32
	t0 := m.g0
	v4 = t0 - i32(32)
	m.g0 = v4
	store32(m.memory[int64(uint32(v4))+24:], uint32(i32(0)))
	store64(m.memory[int64(uint32(v4))+16:], uint64(i64(0)))
	m.fn304(v4+i32(16), v2, v3)
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
		store32(m.memory[int64(uint32(v4))+28:], uint32(i32(60)))
		store32(m.memory[int64(uint32(v4))+20:], uint32(i32(60)))
		store32(m.memory[int64(uint32(v4))+16:], uint32(v1+i32(8)))
		store32(m.memory[int64(uint32(v4))+24:], uint32(v4+i32(12)))
		m.fn341(v0, i32(1286352), v4+i32(16))
		t5 := int32(load32(m.memory[uint32(v1):]))
		t6 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		m.fn348(t5, t6)
	}
l1:
	m.g0 = v4 + i32(32)
}
func (m *Module) fn347(v0, v1 int32) {
	var v2 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	t1 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	t2 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	m.fn12(v2+i32(4), t1, t2)
	{
		{
			t3 := int32(load32(m.memory[int64(uint32(v2))+4:]))
			if t3 != i32(1) {
				goto l0
			}
			t4 := int64(load64(m.memory[int64(uint32(v2))+8:]))
			store64(m.memory[int64(uint32(v0))+12:], uint64(t4))
			t5 := int64(load64(m.memory[uint32(v1):]))
			store64(m.memory[uint32(v0):], uint64(t5))
			t6 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			store32(m.memory[int64(uint32(v0))+8:], uint32(t6))
			goto l1
		}
	l0:
		store32(m.memory[uint32(v0):], uint32(i32(-1)))
		t7 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		store32(m.memory[int64(uint32(v0))+12:], uint32(t7))
		t8 := int64(load64(m.memory[uint32(v1):]))
		store64(m.memory[int64(uint32(v0))+4:], uint64(t8))
	}
l1:
	m.g0 = v2 + i32(16)
}
func (m *Module) fn348(v0, v1 int32) {
	if v1 == 0 {
		return
	}
	m.fn1825(v0, i32(1), v1)
}
func (m *Module) fn349(v0, v1, v2 int32) {
	var v3 int32
	t0 := m.g0
	v3 = t0 - i32(32)
	m.g0 = v3
	m.fn1820(v3+i32(20), v1, v2)
	m.fn322(v3+i32(8), v3+i32(20))
	t1 := int64(load64(m.memory[int64(uint32(v3))+8:]))
	store64(m.memory[uint32(v0):], uint64(t1))
	m.g0 = v3 + i32(32)
}
func (m *Module) fn350(v0, v1 int32) int32 {
	var v2 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	t1 := int32(load16(m.memory[uint32(v0):]))
	store16(m.memory[int64(uint32(v2))+6:], uint16(t1))
	store32(m.memory[int64(uint32(v2))+12:], uint32(i32(61)))
	store32(m.memory[int64(uint32(v2))+8:], uint32(v2+i32(6)))
	t2 := int32(load32(m.memory[uint32(v1):]))
	t3 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	t4 := m.fn284(t2, t3, i32(1286900), v2+i32(8))
	v1 = t4
	m.g0 = v2 + i32(16)
	return v1
}
func (m *Module) fn351(v0, v1, v2 int32) int32 {
	var v3 int32
	t0 := m.g0
	v3 = t0 - i32(16)
	m.g0 = v3
	{
		t1 := int32(load32(m.memory[uint32(v0):]))
		if uint32(v2) <= uint32(t1-v1) {
			goto l0
		}
		m.fn273(v3+i32(8), v0, v1, v2, i32(1), i32(1))
		t2 := int32(load32(m.memory[int64(uint32(v3))+8:]))
		v0 = t2
		if v0 != i32(-1) {
			goto l1
		}
	}
l0:
	v0 = i32(-1)
l1:
	m.g0 = v3 + i32(16)
	return v0
}
func (m *Module) fn352(v0 int32) {
	var v1 int32
	t0 := m.g0
	v1 = t0 - i32(16)
	m.g0 = v1
	t1 := int32(load32(m.memory[uint32(v0):]))
	m.fn273(v1+i32(8), v0, t1, i32(1), i32(8), i32(176))
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
func (m *Module) fn353(v0, v1 int32) {
	m.fn136(v0, v1, i32(8), i32(176))
}
func (m *Module) fn354(v0, v1 int32) {
	var v2, v3, v4, v5, v6 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	{
		{
			{
				if uint32(v1) < uint32(i32(15)) {
					goto l0
				}
				if uint32(v1) > uint32(i32(0x1fffffff)) {
					goto l1
				}
				t1 := int32(uint32(v1<<3) / uint32(i32(7)))
				v1 = i32_shr_u(i32(-1), int32(bits.LeadingZeros32(uint32(t1+i32(-1))))) + i32(1)
				goto l2
			}
		l0:
			p2 := v1&i32(8) + i32(8)
			if uint32(v1) < uint32(i32(4)) {
				p2 = i32(4)
			}
			v1 = p2
		}
	l2:
		m.fn1829(v2+i32(4), i32(4), i32(8), v1)
		t3 := int32(load32(m.memory[int64(uint32(v2))+4:]))
		v3 = t3
		if v3 == 0 {
			goto l1
		}
		t4 := int32(load32(m.memory[int64(uint32(v2))+12:]))
		v4 = t4
		t5 := int32(load32(m.memory[int64(uint32(v2))+8:]))
		v5 = t5
		t6 := m.fn248(v5, v3)
		v6 = t6
		if v6 == 0 {
			m.fn85(v3, v5)
			panic("unreachable")
		}
		v4 = v6 + v4
		v3 = v1 + i32(8)
		if v3 == 0 {
			goto l4
		}
		memory_fill(m.memory, uint32(v4), i32(255), uint32(v3))
	l4:
		store32(m.memory[int64(uint32(v0))+12:], uint32(i32(0)))
		t7 := v0
		v3 = v1 + i32(-1)
		store32(m.memory[int64(uint32(t7))+4:], uint32(v3))
		store32(m.memory[uint32(v0):], uint32(v4))
		t9 := v0
		p8 := int32(uint32(v1)>>3) * i32(7)
		if uint32(v3) < uint32(i32(8)) {
			p8 = v3
		}
		store32(m.memory[int64(uint32(t9))+8:], uint32(p8))
		m.g0 = v2 + i32(16)
		return
	}
l1:
	m.fn1743()
	panic("unreachable")
}
func (m *Module) fn355(v0, v1, v2, v3 int32) {
	var v4 int32
	t0 := m.g0
	v4 = t0 - i32(16)
	m.g0 = v4
	m.fn1(v4+i32(4), v1, i32(0), v2, v3)
	t1 := int32(load32(m.memory[int64(uint32(v4))+8:]))
	v3 = t1
	{
		t2 := int32(load32(m.memory[int64(uint32(v4))+4:]))
		if t2 != i32(1) {
			goto l0
		}
		t3 := int32(load32(m.memory[int64(uint32(v4))+12:]))
		m.fn2(v3, t3)
		panic("unreachable")
	}
l0:
	t4 := int32(load32(m.memory[int64(uint32(v4))+12:]))
	store32(m.memory[int64(uint32(v0))+4:], uint32(t4))
	store32(m.memory[uint32(v0):], uint32(v3))
	m.g0 = v4 + i32(16)
}
func (m *Module) fn356(v0, v1, v2 int32) {
	var v3, v4 int32
	var v5 int64
	var v6 int32
	t0 := int32(load32(m.memory[int64(uint32(v0))+56:]))
	store32(m.memory[int64(uint32(v0))+56:], uint32(t0+v2))
	{
		{
			t1 := int32(load32(m.memory[int64(uint32(v0))+60:]))
			v3 = t1
			if v3 != 0 {
				t2 := v1
				t3 := v2
				v4 = i32(8) - v3
				p4 := v2
				if uint32(v4) < uint32(v2) {
					p4 = v4
				}
				t5 := m.fn287(t2, t3, i32(0), p4)
				v5 = t5
				t6 := int64(load64(m.memory[int64(uint32(v0))+48:]))
				t7 := v0
				v5 = t6 | i64_shl(v5, int64(uint32(v3<<3)))
				store64(m.memory[int64(uint32(t7))+48:], uint64(v5))
				{
					if uint32(v2) < uint32(v4) {
						v3 = v3 + v2
						goto l3
					}
					t8 := int64(load64(m.memory[int64(uint32(v0))+24:]))
					store64(m.memory[int64(uint32(v0))+24:], uint64(t8^v5))
					m.fn286(v0)
					store32(m.memory[int64(uint32(v0))+60:], uint32(i32(0)))
					t9 := int64(load64(m.memory[uint32(v0):]))
					t10 := int64(load64(m.memory[int64(uint32(v0))+48:]))
					store64(m.memory[uint32(v0):], uint64(t9^t10))
					goto l1
				}
			}
			v4 = i32(0)
			goto l1
		}
	l1:
		v6 = v2 - v4
		v3 = v6 & i32(-8)
	l5:
		{
			if uint32(v4) >= uint32(v3) {
				goto l4
			}
			t11 := int64(load64(m.memory[int64(uint32(v0))+24:]))
			t12 := int64(load64(m.memory[uint32(v1+v4):]))
			t13 := v0
			v5 = t12
			store64(m.memory[int64(uint32(t13))+24:], uint64(t11^v5))
			m.fn286(v0)
			t14 := int64(load64(m.memory[uint32(v0):]))
			store64(m.memory[uint32(v0):], uint64(v5^t14))
			v4 = v4 + i32(8)
			goto l5
		}
	l4:
		t15 := v0
		t16 := v1
		t17 := v2
		t18 := v4
		v3 = v6 & i32(7)
		t19 := m.fn287(t16, t17, t18, v3)
		store64(m.memory[int64(uint32(t15))+48:], uint64(t19))
	}
l3:
	store32(m.memory[int64(uint32(v0))+60:], uint32(v3))
}
