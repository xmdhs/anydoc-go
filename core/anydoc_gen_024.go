package core

import (
	"math/bits"
)

func (m *Module) fn1032(v0, v1, v2, v3 int32) int32 {
	var v4, v5 int32
	v4 = i32(0)
	if v1 != v3 {
		goto l0
	}
l2:
	{
		if v1 != 0 {
			goto l1
		}
		return i32(1)
	l1:
		t0 := int32(m.memory[uint32(v2)])
		v3 = t0
		v4 = i32(0)
		t1 := int32(m.memory[uint32(v0)])
		v5 = t1
		v2 = v2 + i32(1)
		v1 = v1 + i32(-1)
		v0 = v0 + i32(1)
		t3 := v5
		p2 := i32(0)
		if uint32((v5+i32(-65))&i32(255)) < uint32(i32(26)) {
			p2 = i32(32)
		}
		t5 := (t3 | p2) & i32(255)
		t6 := v3
		p4 := i32(0)
		if uint32((v3+i32(-65))&i32(255)) < uint32(i32(26)) {
			p4 = i32(32)
		}
		if t5 == (t6|p4)&i32(255) {
			goto l2
		}
	}
l0:
	return v4
}
func (m *Module) fn1033(v0 int32) {
	t0 := int32(load32(m.memory[int64(uint32(v0))+44:]))
	t1 := int32(load32(m.memory[int64(uint32(v0))+48:]))
	m.fn16(t0, t1)
	t2 := int32(load32(m.memory[int64(uint32(v0))+56:]))
	t3 := int32(load32(m.memory[int64(uint32(v0))+60:]))
	m.fn16(t2, t3)
}
func (m *Module) fn1034(v0, v1, v2 int32) {
	var v3 int32
	var v4, v5, v6 int64
	t0 := m.g0
	v3 = t0 - i32(144)
	m.g0 = v3
	store64(m.memory[int64(uint32(v3))+104:], uint64(i64(0)))
	store32(m.memory[int64(uint32(v3))+100:], uint32(v2))
	store32(m.memory[int64(uint32(v3))+96:], uint32(v1))
	m.fn111(v3+i32(32), v3+i32(96))
	{
		{
			t1 := int32(load32(m.memory[int64(uint32(v3))+32:]))
			if t1 != 0 {
				goto l0
			}
			t2 := int32(load32(m.memory[int64(uint32(v3))+44:]))
			store32(m.memory[int64(uint32(v3))+128:], uint32(t2))
			t3 := int64(load64(m.memory[int64(uint32(v3))+36:]))
			store64(m.memory[int64(uint32(v3))+120:], uint64(t3))
			store32(m.memory[int64(uint32(v3))+140:], uint32(i32(145)))
			store32(m.memory[int64(uint32(v3))+136:], uint32(v3+i32(120)))
			m.fn73(v3+i32(96), i32(1052086), v3+i32(136))
			store32(m.memory[int64(uint32(v3))+108:], uint32(i32(-1)))
			m.fn116(v3 + i32(120))
			t4 := int64(load64(m.memory[int64(uint32(v3))+96:]))
			t5 := v3
			v4 = t4
			store64(m.memory[int64(uint32(t5))+56:], uint64(v4))
			t6 := int64(load64(m.memory[int64(uint32(v3))+112:]))
			t7 := v3
			v5 = t6
			store64(m.memory[int64(uint32(t7))+72:], uint64(v5))
			t8 := int64(load64(m.memory[int64(uint32(v3))+104:]))
			t9 := v3
			v6 = t8
			store64(m.memory[int64(uint32(t9))+64:], uint64(v6))
			store64(m.memory[int64(uint32(v3))+16:], uint64(v5))
			store64(m.memory[int64(uint32(v3))+8:], uint64(v6))
			store64(m.memory[uint32(v3):], uint64(v4))
			store64(m.memory[int64(uint32(v0))+20:], uint64(v5))
			store64(m.memory[int64(uint32(v0))+12:], uint64(v6))
			store64(m.memory[int64(uint32(v0))+4:], uint64(v4))
			store32(m.memory[uint32(v0):], uint32(i32(0)))
			goto l1
		}
	l0:
		t10 := int64(load64(m.memory[int64(uint32(v3))+32:]))
		store64(m.memory[int64(uint32(v3))+96:], uint64(t10))
		t11 := int64(load64(m.memory[int64(uint32(v3))+40:]))
		store64(m.memory[int64(uint32(v3))+104:], uint64(t11))
		t12 := int64(load64(m.memory[int64(uint32(v3))+48:]))
		t13 := v3
		v4 = t12
		store64(m.memory[int64(uint32(t13))+112:], uint64(v4))
		{
			t14 := int32(load32(m.memory[int64(uint32(int32(v4)))+64:]))
			v2 = t14
			if uint32(v2) > uint32(i32(100000)) {
				goto l2
			}
			t15 := int64(load64(m.memory[int64(uint32(v3))+112:]))
			store64(m.memory[int64(uint32(v3))+72:], uint64(t15))
			t16 := int64(load64(m.memory[int64(uint32(v3))+104:]))
			store64(m.memory[int64(uint32(v3))+64:], uint64(t16))
			t17 := int64(load64(m.memory[int64(uint32(v3))+96:]))
			store64(m.memory[int64(uint32(v3))+56:], uint64(t17))
			m.fn22(v3, i32(3))
			t18 := int64(load64(m.memory[int64(uint32(i32(0)))+1286256:]))
			store64(m.memory[int64(uint32(v3))+80:], uint64(t18))
			t19 := int64(load64(m.memory[int64(uint32(i32(0)))+1286264:]))
			store64(m.memory[int64(uint32(v3))+88:], uint64(t19))
			t20 := int64(load64(m.memory[uint32(v3):]))
			v4 = t20
			t21 := int64(load64(m.memory[int64(uint32(v3))+8:]))
			v5 = t21
			memory_copy(m.memory, uint32(v0), uint32(v3+i32(56)), uint32(i32(40)))
			store64(m.memory[int64(uint32(v0))+56:], uint64(i64(0)))
			store64(m.memory[int64(uint32(v0))+48:], uint64(v5))
			store64(m.memory[int64(uint32(v0))+40:], uint64(v4))
			goto l1
		}
	l2:
		store32(m.memory[uint32(v3):], uint32(v2))
		store32(m.memory[int64(uint32(v3))+60:], uint32(i32(5)))
		store32(m.memory[int64(uint32(v3))+56:], uint32(v3))
		m.fn73(v0+i32(8), i32(1066974), v3+i32(56))
		store32(m.memory[int64(uint32(v0))+24:], uint32(i32(15)))
		store32(m.memory[int64(uint32(v0))+20:], uint32(i32(1075208)))
		store64(m.memory[uint32(v0):], uint64(i64(-0x7ffffffd00000000)))
		m.fn132(v3 + i32(96))
	}
l1:
	m.g0 = v3 + i32(144)
}
func (m *Module) fn1035(v0, v1, v2, v3 int32) {
	var v4 int32
	t0 := m.g0
	v4 = t0 - i32(32)
	m.g0 = v4
	m.fn1050(v4+i32(8), v1, v2, v3)
	{
		t1 := int32(load32(m.memory[int64(uint32(v4))+8:]))
		v3 = t1
		if v3 == i32(-1) {
			t5 := int64(load64(m.memory[int64(uint32(v4))+24:]))
			store64(m.memory[int64(uint32(v0))+16:], uint64(t5))
			t6 := int64(load64(m.memory[int64(uint32(v4))+16:]))
			store64(m.memory[int64(uint32(v0))+8:], uint64(t6))
			t7 := int64(load64(m.memory[int64(uint32(v4))+8:]))
			store64(m.memory[uint32(v0):], uint64(t7))
			goto l2
		}
		if v3 != i32(-0x7ffffffd) {
			goto l1
		}
		t2 := int64(load64(m.memory[int64(uint32(v4))+24:]))
		store64(m.memory[int64(uint32(v0))+16:], uint64(t2))
		t3 := int64(load64(m.memory[int64(uint32(v4))+16:]))
		store64(m.memory[int64(uint32(v0))+8:], uint64(t3))
		t4 := int64(load64(m.memory[int64(uint32(v4))+8:]))
		store64(m.memory[uint32(v0):], uint64(t4))
		goto l2
	}
l1:
	store64(m.memory[uint32(v0):], uint64(i64(0xffffffff)))
	m.fn785(v4 + i32(8))
l2:
	m.g0 = v4 + i32(32)
}
func (m *Module) fn1036(v0, v1 int32) int32 {
	var v2, v3 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	m.fn626(v2+i32(8), i32(1083910), i32(9), v0, v1)
	{
		{
			t1 := int32(load32(m.memory[int64(uint32(v2))+8:]))
			v3 = t1
			p2 := v0
			if v3 != 0 {
				p2 = v3
			}
			v0 = p2
			t3 := int32(load32(m.memory[int64(uint32(v2))+12:]))
			t5 := v0
			p4 := v1
			if v3 != 0 {
				p4 = t3
			}
			v1 = p4
			t6 := m.fn15(t5, v1, i32(1083919), i32(20))
			if t6 == 0 {
				goto l0
			}
			v1 = i32(7)
			goto l1
		}
	l0:
		{
			t7 := m.fn15(v0, v1, i32(1083939), i32(39))
			if t7 == 0 {
				goto l2
			}
			v1 = i32(2)
			goto l1
		}
	l2:
		{
			t8 := m.fn15(v0, v1, i32(1083978), i32(46))
			if t8 == 0 {
				goto l3
			}
			v1 = i32(9)
			goto l1
		}
	l3:
		t9 := m.fn15(v0, v1, i32(1084024), i32(47))
		p10 := i32(-1)
		if t9 != 0 {
			p10 = i32(10)
		}
		v1 = p10
	}
l1:
	m.g0 = v2 + i32(16)
	return v1
}
func (m *Module) fn1037(v0 int32) {
	{
		t0 := int32(load32(m.memory[uint32(v0):]))
		if t0 != i32(-1) {
			goto l0
		}
		m.fn1051(v0 + i32(4))
		return
	}
l0:
	m.fn785(v0)
}
func (m *Module) fn1038(v0, v1, v2, v3 int32) {
	var v4 int32
	var v5, v6, v7 int64
	var v8, v9, v10, v11, v12, v13, v14 int32
	t0 := m.g0
	v4 = t0 - i32(304)
	m.g0 = v4
	m.fn1040(v4+i32(88), v1, v2, v3)
	v3 = v4 + i32(92)
	{
		{
			t1 := int32(load32(m.memory[int64(uint32(v4))+88:]))
			v2 = t1
			if v2 != i32(-2) {
				goto l0
			}
			t2 := int64(load64(m.memory[int64(uint32(v3))+16:]))
			t3 := v4
			v5 = t2
			store64(m.memory[int64(uint32(t3))+64:], uint64(v5))
			t4 := int64(load64(m.memory[int64(uint32(v3))+8:]))
			t5 := v4
			v6 = t4
			store64(m.memory[int64(uint32(t5))+56:], uint64(v6))
			t6 := int64(load64(m.memory[uint32(v3):]))
			t7 := v4
			v7 = t6
			store64(m.memory[int64(uint32(t7))+48:], uint64(v7))
			store64(m.memory[int64(uint32(v0))+20:], uint64(v5))
			store64(m.memory[int64(uint32(v0))+12:], uint64(v6))
			store64(m.memory[int64(uint32(v0))+4:], uint64(v7))
			store32(m.memory[uint32(v0):], uint32(i32(0)))
			goto l1
		}
	l0:
		memory_copy(m.memory, uint32(v4+i32(48)), uint32(v3), uint32(i32(40)))
		{
			if v2 == i32(-1) {
				goto l2
			}
			store32(m.memory[int64(uint32(v4))+88:], uint32(v2))
			memory_copy(m.memory, uint32(v4+i32(88)+i32(4)), uint32(v4+i32(48)), uint32(i32(40)))
			m.fn22(v4+i32(256), i32(3))
			t8 := int64(load64(m.memory[int64(uint32(i32(0)))+1286256:]))
			store64(m.memory[int64(uint32(v4))+136:], uint64(t8))
			t9 := int64(load64(m.memory[int64(uint32(i32(0)))+1286264:]))
			store64(m.memory[int64(uint32(v4))+144:], uint64(t9))
			t10 := int64(load64(m.memory[int64(uint32(v4))+264:]))
			store64(m.memory[int64(uint32(v4))+160:], uint64(t10))
			t11 := int64(load64(m.memory[int64(uint32(v4))+256:]))
			store64(m.memory[int64(uint32(v4))+152:], uint64(t11))
			t12 := int32(load32(m.memory[int64(uint32(v4))+116:]))
			t13 := int32(load32(m.memory[int64(uint32(v4))+120:]))
			m.fn868(v4+i32(172), t12, t13)
			store32(m.memory[int64(uint32(v4))+192:], uint32(i32(1076594)))
			store32(m.memory[int64(uint32(v4))+188:], uint32(i32(60)))
			store32(m.memory[int64(uint32(v4))+184:], uint32(i32(1084156)))
			v8 = v4 + i32(136) + i32(16)
			store32(m.memory[int64(uint32(v4))+196:], uint32(i32(12)))
			v9 = v4 + i32(256) + i32(12)
			v10 = v4 + i32(228) + i32(12)
		l4:
			{
				{
					{
						t14 := m.fn863(v4 + i32(172))
						v3 = t14
						if v3 == 0 {
							t24 := int32(load32(m.memory[int64(uint32(v4))+172:]))
							t25 := int32(load32(m.memory[int64(uint32(v4))+176:]))
							m.fn44(t24, t25)
							t26 := int64(load64(m.memory[int64(uint32(v4))+160:]))
							store64(m.memory[int64(uint32(v0))+24:], uint64(t26))
							t27 := int64(load64(m.memory[int64(uint32(v4))+152:]))
							store64(m.memory[int64(uint32(v0))+16:], uint64(t27))
							t28 := int64(load64(m.memory[int64(uint32(v4))+144:]))
							store64(m.memory[int64(uint32(v0))+8:], uint64(t28))
							t29 := int64(load64(m.memory[int64(uint32(v4))+136:]))
							store64(m.memory[uint32(v0):], uint64(t29))
							m.fn1042(v4 + i32(88))
							goto l1
						}
						t15 := int32(load32(m.memory[uint32(v3+i32(16)):]))
						t16 := v4 + i32(40)
						v2 = t15
						t17 := int32(load32(m.memory[uint32(v3+i32(20)):]))
						t18 := v2
						v3 = t17
						m.fn909(t16, t18, v3, i32(1084216), i32(2))
						t19 := int32(load32(m.memory[int64(uint32(v4))+44:]))
						v11 = t19
						m.fn909(v4+i32(32), v2, v3, i32(1084218), i32(6))
						t20 := int32(load32(m.memory[int64(uint32(v4))+36:]))
						v12 = t20
						t21 := int32(load32(m.memory[int64(uint32(v4))+32:]))
						v1 = t21
						t22 := int32(load32(m.memory[int64(uint32(v4))+40:]))
						v13 = t22
						if v13 == 0 {
							goto l4
						}
						if v1 == 0 {
							goto l4
						}
						m.fn909(v4+i32(24), v2, v3, i32(1084224), i32(10))
						t23 := int32(load32(m.memory[int64(uint32(v4))+24:]))
						v14 = t23
						if v14 != 0 {
							goto l5
						}
						v14 = i32(0)
						goto l6
					}
				l5:
					t30 := int32(load32(m.memory[int64(uint32(v4))+28:]))
					t31 := m.fn1032(v14, t30, i32(1084234), i32(8))
					v14 = t31
				}
			l6:
				m.fn909(v4+i32(16), v2, v3, i32(1084242), i32(4))
				t32 := int32(load32(m.memory[int64(uint32(v4))+16:]))
				t33 := v4 + i32(256)
				v3 = t32
				p34 := i32(1)
				if v3 != 0 {
					p34 = v3
				}
				v2 = p34
				t35 := int32(load32(m.memory[int64(uint32(v4))+20:]))
				t37 := v2
				p36 := i32(0)
				if v3 != 0 {
					p36 = t35
				}
				v3 = p36
				m.fn877(t33, t37, v3)
				{
					t38 := int32(load32(m.memory[int64(uint32(v4))+256:]))
					if t38 == i32(-1) {
						goto l7
					}
					t39 := int32(load32(m.memory[int64(uint32(v4))+264:]))
					store32(m.memory[int64(uint32(v4))+208:], uint32(t39))
					t40 := int64(load64(m.memory[int64(uint32(v4))+256:]))
					store64(m.memory[int64(uint32(v4))+200:], uint64(t40))
					goto l8
				}
			l7:
				m.fn51(v4+i32(200), v2, v3)
			l8:
				m.fn51(v4+i32(216), v13, v11)
				m.fn51(v4+i32(228), v1, v12)
				t41 := int32(load32(m.memory[int64(uint32(v4))+208:]))
				store32(m.memory[int64(uint32(v10))+8:], uint32(t41))
				t42 := int64(load64(m.memory[int64(uint32(v4))+200:]))
				store64(m.memory[uint32(v10):], uint64(t42))
				m.memory[int64(uint32(v4))+252] = byte(v14)
				t43 := int64(load64(m.memory[int64(uint32(v4))+152:]))
				t44 := int64(load64(m.memory[int64(uint32(v4))+160:]))
				t45 := int32(load32(m.memory[int64(uint32(v4))+220:]))
				t46 := int32(load32(m.memory[int64(uint32(v4))+224:]))
				t47 := m.fn540(t43, t44, t45, t46)
				v5 = t47
				store32(m.memory[int64(uint32(v4))+300:], uint32(v4+i32(216)))
				{
					t48 := int32(load32(m.memory[int64(uint32(v4))+144:]))
					if t48 != 0 {
						goto l9
					}
					_ = m.fn672(v4+i32(136), v8)
				}
			l9:
				store32(m.memory[int64(uint32(v4))+260:], uint32(v4+i32(136)))
				store32(m.memory[int64(uint32(v4))+256:], uint32(v4+i32(300)))
				t50 := int32(load32(m.memory[int64(uint32(v4))+136:]))
				t51 := int32(load32(m.memory[int64(uint32(v4))+140:]))
				m.fn69(v4+i32(8), t50, t51, v5, v4+i32(256), i32(146))
				t52 := int32(load32(m.memory[int64(uint32(v4))+12:]))
				v3 = t52
				t53 := int32(load32(m.memory[int64(uint32(v4))+136:]))
				v2 = t53
				{
					t54 := int32(load32(m.memory[int64(uint32(v4))+8:]))
					if t54 != i32(1) {
						v3 = v2 + (i32(0)-v3)*i32(40) + i32(-28)
						t66 := int64(load64(m.memory[uint32(v3):]))
						v5 = t66
						t67 := int64(load64(m.memory[int64(uint32(v4))+228:]))
						store64(m.memory[uint32(v3):], uint64(t67))
						t68 := int64(load64(m.memory[int64(uint32(v3))+8:]))
						v6 = t68
						t69 := int64(load64(m.memory[int64(uint32(v4))+236:]))
						store64(m.memory[int64(uint32(v3))+8:], uint64(t69))
						t70 := int64(load64(m.memory[int64(uint32(v3))+16:]))
						v7 = t70
						t71 := int64(load64(m.memory[int64(uint32(v4))+244:]))
						store64(m.memory[int64(uint32(v3))+16:], uint64(t71))
						t72 := int32(load32(m.memory[int64(uint32(v3))+24:]))
						v2 = t72
						t73 := int32(load32(m.memory[int64(uint32(v4))+252:]))
						store32(m.memory[int64(uint32(v3))+24:], uint32(t73))
						store32(m.memory[int64(uint32(v4))+280:], uint32(v2))
						store64(m.memory[int64(uint32(v4))+272:], uint64(v7))
						store64(m.memory[int64(uint32(v4))+264:], uint64(v6))
						store64(m.memory[int64(uint32(v4))+256:], uint64(v5))
						t74 := int32(load32(m.memory[int64(uint32(v4))+216:]))
						t75 := int32(load32(m.memory[int64(uint32(v4))+220:]))
						m.fn16(t74, t75)
						t76 := int32(load32(m.memory[int64(uint32(v4))+256:]))
						if t76 == i32(-1) {
							goto l4
						}
						m.fn759(v4 + i32(256))
						goto l4
					}
					v1 = v2 + v3
					t55 := int32(m.memory[uint32(v1)])
					v11 = t55
					t56 := int64(load64(m.memory[int64(uint32(v4))+216:]))
					v6 = t56
					t57 := int32(load32(m.memory[int64(uint32(v4))+224:]))
					v12 = t57
					t58 := v1
					v13 = int32(uint32(int32(v5)) >> 25)
					m.memory[uint32(t58)] = byte(v13)
					t59 := int64(load64(m.memory[int64(uint32(v4))+228:]))
					store64(m.memory[uint32(v9):], uint64(t59))
					t60 := int64(load64(m.memory[int64(uint32(v4))+236:]))
					store64(m.memory[int64(uint32(v9))+8:], uint64(t60))
					t61 := int64(load64(m.memory[int64(uint32(v4))+244:]))
					store64(m.memory[int64(uint32(v9))+16:], uint64(t61))
					t62 := int32(load32(m.memory[int64(uint32(v4))+252:]))
					store32(m.memory[int64(uint32(v9))+24:], uint32(t62))
					t63 := int32(load32(m.memory[int64(uint32(v4))+140:]))
					m.memory[uint32(v2+t63&(v3+i32(-8))+i32(8))] = byte(v13)
					store32(m.memory[int64(uint32(v4))+264:], uint32(v12))
					store64(m.memory[int64(uint32(v4))+256:], uint64(v6))
					t64 := int32(load32(m.memory[int64(uint32(v4))+144:]))
					store32(m.memory[int64(uint32(v4))+144:], uint32(t64-v11&i32(1)))
					t65 := int32(load32(m.memory[int64(uint32(v4))+148:]))
					store32(m.memory[int64(uint32(v4))+148:], uint32(t65+i32(1)))
					memory_copy(m.memory, uint32(v2+(i32(0)-v3)*i32(40)+i32(-40)), uint32(v4+i32(256)), uint32(i32(40)))
					goto l4
				}
			}
		}
	l2:
		m.fn34(v4 + i32(88))
		t77 := int64(load64(m.memory[int64(uint32(v4))+88:]))
		v5 = t77
		t78 := int64(load64(m.memory[int64(uint32(v4))+96:]))
		v6 = t78
		t79 := int64(load64(m.memory[int64(uint32(i32(0)))+1286264:]))
		store64(m.memory[int64(uint32(v0))+8:], uint64(t79))
		t80 := int64(load64(m.memory[int64(uint32(i32(0)))+1286256:]))
		store64(m.memory[uint32(v0):], uint64(t80))
		store64(m.memory[int64(uint32(v0))+24:], uint64(v6))
		store64(m.memory[int64(uint32(v0))+16:], uint64(v5))
	}
l1:
	m.g0 = v4 + i32(304)
}
func (m *Module) fn1039(v0, v1, v2 int32) int32 {
	var v3, v4, v5, v6, v7 int32
	var v8 int64
	var v9, v10 int32
	t0 := m.g0
	v3 = t0 - i32(48)
	m.g0 = v3
	store32(m.memory[int64(uint32(v3))+4:], uint32(v2))
	store32(m.memory[uint32(v3):], uint32(v1))
	t1 := int32(load32(m.memory[int64(uint32(v0))+12:]))
	t2 := v3
	v2 = t1
	store32(m.memory[int64(uint32(t2))+32:], uint32(v2))
	t3 := int32(load32(m.memory[uint32(v0):]))
	t4 := v3
	v1 = t3
	store32(m.memory[int64(uint32(t4))+24:], uint32(v1))
	store32(m.memory[int64(uint32(v3))+16:], uint32(v1+i32(8)))
	t5 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	store32(m.memory[int64(uint32(v3))+20:], uint32(v1+t5+i32(1)))
	t6 := int64(load64(m.memory[uint32(v1):]))
	store64(m.memory[int64(uint32(v3))+8:], uint64((t6^i64(-1))&i64(-0x7f7f7f7f7f7f7f80)))
	store32(m.memory[int64(uint32(v3))+40:], uint32(v3))
	{
	l2:
		{
			if v2 != 0 {
				goto l0
			}
			v2 = i32(0)
			goto l1
		l0:
			t7 := m.fn758(v3 + i32(8))
			v0 = t7
			t8 := int32(load32(m.memory[int64(uint32(v3))+32:]))
			t9 := v3
			v2 = t8 + i32(-1)
			store32(m.memory[int64(uint32(t9))+32:], uint32(v2))
			t10 := int32(load32(m.memory[int64(uint32(v3))+40:]))
			v1 = t10
			t11 := int32(load32(m.memory[uint32(v1):]))
			v4 = t11
			t12 := int32(load32(m.memory[uint32(v1+i32(4)):]))
			t13 := v4
			v5 = t12
			t14 := v5
			v6 = v0 + i32(-28)
			t15 := m.fn772(t13, t14, v6)
			if t15 == 0 {
				goto l2
			}
		}
		v7 = v0 + i32(-40)
		t16 := int32(load32(m.memory[int64(uint32(v3))+24:]))
		v1 = t16
		t17 := int32(load32(m.memory[int64(uint32(v3))+16:]))
		v0 = t17
		t18 := int64(load64(m.memory[int64(uint32(v3))+8:]))
		v8 = t18
	l5:
		if v8 == 0 {
			goto l3
		}
		{
			t19 := v4
			t20 := v5
			v9 = v1 + (i32(0)-int32(uint32(int64(bits.TrailingZeros64(uint64(v8))))>>3))*i32(40)
			v10 = v9 + i32(-28)
			t21 := m.fn772(t19, t20, v10)
			if t21 == 0 {
				goto l4
			}
			t22 := int32(load32(m.memory[uint32(v7+i32(4)):]))
			t23 := int32(load32(m.memory[uint32(v7+i32(8)):]))
			t24 := int32(load32(m.memory[uint32(v9+i32(-36)):]))
			t25 := int32(load32(m.memory[uint32(v9+i32(-32)):]))
			t26 := m.fn259(t22, t23, t24, t25)
			t27 := v7
			t28 := v9 + i32(-40)
			var p29 int32
			if int32(int8(t26)) < i32(1) {
				p29 = 1
			}
			v9 = p29
			p30 := t28
			if v9 != 0 {
				p30 = t27
			}
			v7 = p30
			p31 := v10
			if v9 != 0 {
				p31 = v6
			}
			v6 = p31
		}
	l4:
		v8 = (v8 + i64(-1)) & v8
		v2 = v2 + i32(-1)
		goto l5
	l3:
		{
			if v2 == 0 {
				goto l6
			}
			v1 = v1 + i32(-320)
			t32 := int64(load64(m.memory[uint32(v0):]))
			v8 = (t32 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
			v0 = v0 + i32(8)
			goto l5
		}
	l6:
		p33 := i32(0)
		if v7 != 0 {
			p33 = v6
		}
		v2 = p33
	}
l1:
	m.g0 = v3 + i32(48)
	return v2
}
func (m *Module) fn1040(v0, v1, v2, v3 int32) {
	var v4 int32
	t0 := m.g0
	v4 = t0 - i32(64)
	m.g0 = v4
	m.fn1035(v4+i32(20), v1, v2, v3)
	t1 := int32(load32(m.memory[int64(uint32(v4))+28:]))
	v2 = t1
	t2 := int32(load32(m.memory[int64(uint32(v4))+24:]))
	v3 = t2
	{
		t3 := int32(load32(m.memory[int64(uint32(v4))+20:]))
		v1 = t3
		if v1 == i32(-1) {
			goto l0
		}
		t4 := int32(load32(m.memory[int64(uint32(v4))+40:]))
		store32(m.memory[int64(uint32(v0))+24:], uint32(t4))
		t5 := int64(load64(m.memory[int64(uint32(v4))+32:]))
		store64(m.memory[int64(uint32(v0))+16:], uint64(t5))
		store32(m.memory[int64(uint32(v0))+12:], uint32(v2))
		store32(m.memory[int64(uint32(v0))+8:], uint32(v3))
		store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
		store32(m.memory[uint32(v0):], uint32(i32(-2)))
		goto l1
	}
l0:
	if v3 == 0 {
		store32(m.memory[uint32(v0):], uint32(i32(-1)))
		goto l1
	}
	store32(m.memory[int64(uint32(v4))+16:], uint32(v2))
	store32(m.memory[int64(uint32(v4))+12:], uint32(v3))
	m.fn1053(v4+i32(20), v3+i32(8), v2)
	{
		t6 := int32(load32(m.memory[int64(uint32(v4))+20:]))
		if t6 != i32(-1) {
			memory_copy(m.memory, uint32(v0), uint32(v4+i32(20)), uint32(i32(44)))
			goto l5
		}
		v3 = v4 + i32(24)
		t7 := int32(load32(m.memory[int64(uint32(v4))+24:]))
		if t7 != i32(-0x7ffffffd) {
			store32(m.memory[uint32(v0):], uint32(i32(-1)))
			m.fn785(v3)
			goto l5
		}
		store32(m.memory[uint32(v0):], uint32(i32(-2)))
		t8 := int64(load64(m.memory[int64(uint32(v3))+16:]))
		store64(m.memory[int64(uint32(v0))+20:], uint64(t8))
		t9 := int64(load64(m.memory[int64(uint32(v3))+8:]))
		store64(m.memory[int64(uint32(v0))+12:], uint64(t9))
		t10 := int64(load64(m.memory[uint32(v3):]))
		store64(m.memory[int64(uint32(v0))+4:], uint64(t10))
		goto l5
	}
l5:
	m.fn754(v4 + i32(12))
l1:
	m.g0 = v4 + i32(64)
}
func (m *Module) fn1041(v0, v1, v2 int32) {
	if v1 == 0 {
		goto l0
	}
	m.fn51(v0, v1, v2)
	return
l0:
	store32(m.memory[uint32(v0):], uint32(i32(-1)))
}
func (m *Module) fn1042(v0 int32) {
	var v1, v2, v3 int32
	m.fn1051(v0 + i32(36))
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	m.fn16(t0, t1)
	t2 := int32(load32(m.memory[int64(uint32(v0))+20:]))
	v1 = t2
	t3 := int32(load32(m.memory[int64(uint32(v0))+16:]))
	v2 = t3
	v3 = v2
l1:
	{
		if v1 == 0 {
			goto l0
		}
		m.fn1051(v3 + i32(24))
		t4 := int32(load32(m.memory[uint32(v3):]))
		t5 := int32(load32(m.memory[uint32(v3+i32(4)):]))
		m.fn16(t4, t5)
		t6 := int32(load32(m.memory[uint32(v3+i32(12)):]))
		t7 := int32(load32(m.memory[uint32(v3+i32(16)):]))
		m.fn16(t6, t7)
		v1 = v1 + i32(-1)
		v3 = v3 + i32(32)
		goto l1
	}
l0:
	t8 := int32(load32(m.memory[int64(uint32(v0))+12:]))
	m.fn136(t8, v2, i32(4), i32(32))
	t9 := int32(load32(m.memory[int64(uint32(v0))+32:]))
	v1 = t9
	t10 := int32(load32(m.memory[int64(uint32(v0))+28:]))
	v2 = t10
	v3 = v2
l5:
	if v1 == 0 {
		goto l2
	}
	{
		{
			t11 := int32(load32(m.memory[uint32(v3):]))
			if t11 == i32(-1) {
				goto l3
			}
			m.fn1042(v3)
			goto l4
		}
	l3:
		t12 := int32(load32(m.memory[uint32(v3+i32(4)):]))
		t13 := int32(load32(m.memory[uint32(v3+i32(8)):]))
		m.fn16(t12, t13)
	}
l4:
	v1 = v1 + i32(-1)
	v3 = v3 + i32(44)
	goto l5
l2:
	t14 := int32(load32(m.memory[int64(uint32(v0))+24:]))
	m.fn136(t14, v2, i32(4), i32(44))
}
func (m *Module) fn1043(v0 int32) {
	var v1, v2, v3 int32
	var v4 int64
	var v5 int32
	t0 := m.g0
	v1 = t0 - i32(32)
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
			v0 = t3
			if v0 == 0 {
				goto l1
			}
			t4 := int64(load64(m.memory[uint32(v3):]))
			v4 = t4
			store32(m.memory[int64(uint32(v1))+24:], uint32(v0))
			store32(m.memory[int64(uint32(v1))+16:], uint32(v3))
			v5 = i32(1)
			store32(m.memory[int64(uint32(v1))+12:], uint32(v3+v2+i32(1)))
			store32(m.memory[int64(uint32(v1))+8:], uint32(v3+i32(8)))
			store64(m.memory[uint32(v1):], uint64((v4^i64(-1))&i64(-0x7f7f7f7f7f7f7f80)))
		l2:
			{
				if v5 == 0 {
					goto l1
				}
				t5 := m.fn758(v1)
				v0 = t5
				t6 := int32(load32(m.memory[int64(uint32(v1))+24:]))
				t7 := v1
				v5 = t6 + i32(-1)
				store32(m.memory[int64(uint32(t7))+24:], uint32(v5))
				t8 := int32(load32(m.memory[uint32(v0+i32(-40)):]))
				t9 := int32(load32(m.memory[uint32(v0+i32(-36)):]))
				m.fn16(t8, t9)
				m.fn759(v0 + i32(-28))
				goto l2
			}
		}
	l1:
		m.fn39(v1, i32(40), i32(8), v2+i32(1))
		t10 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		t11 := int32(load32(m.memory[uint32(v1):]))
		t12 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		m.fn40(v3-t10, t11, t12)
	}
l0:
	m.g0 = v1 + i32(32)
}
func (m *Module) fn1044(v0 int32) {
	{
		t0 := int32(load32(m.memory[uint32(v0):]))
		if t0 == 0 {
			goto l0
		}
		m.fn1043(v0)
		return
	}
l0:
	m.fn785(v0 + i32(4))
}
func (m *Module) fn1045(v0, v1, v2 int32) int32 {
	var v3 int32
	t0 := m.g0
	v3 = t0 - i32(16)
	m.g0 = v3
	m.fn631(v3+i32(8), v1, v2, i32(47))
	t1 := int32(load32(m.memory[int64(uint32(v3))+8:]))
	t2 := int32(load32(m.memory[int64(uint32(v3))+12:]))
	m.fn115(v3, v0+i32(24), t1, t2)
	t3 := int32(load32(m.memory[uint32(v3):]))
	v2 = t3
	m.g0 = v3 + i32(16)
	var p4 int32
	if v2 != i32(0) {
		p4 = 1
	}
	return p4
}
func (m *Module) fn1046(v0, v1, v2, v3, v4, v5, v6 int32) {
	var v7 int32
	t0 := m.g0
	v7 = t0 - i32(16)
	m.g0 = v7
	m.fn845(v7+i32(8), v1, v2, v3, v4, v5, v6)
	{
		{
			t1 := int32(load32(m.memory[int64(uint32(v7))+8:]))
			v4 = t1
			if v4 == 0 {
				goto l0
			}
			t2 := int32(load32(m.memory[int64(uint32(v7))+12:]))
			v2 = t2
			goto l1
		}
	l0:
		v2 = v2 << 5
	l5:
		if v2 != 0 {
			goto l2
		}
		v4 = i32(0)
		goto l1
	l2:
		{
			t3 := int32(load32(m.memory[uint32(v1+i32(4)):]))
			t4 := int32(load32(m.memory[uint32(v1+i32(8)):]))
			t5 := m.fn773(t3, t4, v5, v6)
			if t5 == 0 {
				goto l3
			}
			t6 := int32(load32(m.memory[uint32(v1+i32(24)):]))
			if t6 == 0 {
				goto l4
			}
		}
	l3:
		v1 = v1 + i32(32)
		v2 = v2 + i32(-32)
		goto l5
	l4:
		t7 := int32(load32(m.memory[int64(uint32(v1))+20:]))
		v2 = t7
		t8 := int32(load32(m.memory[int64(uint32(v1))+16:]))
		v4 = t8
	}
l1:
	store32(m.memory[uint32(v0):], uint32(v4))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v2))
	m.g0 = v7 + i32(16)
}
func (m *Module) fn1047(v0 int32) {
	{
		t0 := int32(load32(m.memory[uint32(v0):]))
		if t0 == i32(-2) {
			goto l0
		}
		m.fn1054(v0)
		return
	}
l0:
	m.fn785(v0 + i32(4))
}
func (m *Module) fn1048(v0 int32) {
	var v1, v2, v3, v4 int32
	var v5 int64
	var v6, v7 int32
	t0 := m.g0
	v1 = t0 - i32(16)
	m.g0 = v1
	m.fn132(v0)
	{
		t1 := int32(load32(m.memory[int64(uint32(v0))+28:]))
		v2 = t1
		if v2 == 0 {
			goto l0
		}
		t2 := int32(load32(m.memory[int64(uint32(v0))+24:]))
		v3 = t2
		{
			t3 := int32(load32(m.memory[int64(uint32(v0))+36:]))
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
					v7 = v6 + (i32(0)-int32(uint32(int64(bits.TrailingZeros64(uint64(v5))))>>3))*i32(20)
					t6 := int32(load32(m.memory[uint32(v7+i32(-20)):]))
					t7 := int32(load32(m.memory[uint32(v7+i32(-16)):]))
					m.fn16(t6, t7)
					m.fn754(v7 + i32(-8))
					v4 = v4 + i32(-1)
					v5 = (v5 + i64(-1)) & v5
					goto l4
				}
				v6 = v6 + i32(-160)
				t5 := int64(load64(m.memory[uint32(v0):]))
				v5 = (t5 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
				v0 = v0 + i32(8)
				goto l3
			}
		}
	l1:
		m.fn39(v1+i32(4), i32(20), i32(8), v2+i32(1))
		t8 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		t9 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t10 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		m.fn40(v3-t8, t9, t10)
	}
l0:
	m.g0 = v1 + i32(16)
}
func (m *Module) fn1049(v0, v1 int32) int32 {
	var v2, v3 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	{
		t1 := int32(load32(m.memory[uint32(v0):]))
		v3 = t1 ^ i32(-0x80000000)
		p2 := i32(1)
		if uint32(v3) < uint32(i32(6)) {
			p2 = v3
		}
		switch p2 {
		default:
			store32(m.memory[int64(uint32(v2))+4:], uint32(v0+i32(4)))
			store32(m.memory[int64(uint32(v2))+12:], uint32(i32(24)))
			store32(m.memory[int64(uint32(v2))+8:], uint32(v2+i32(4)))
			t3 := int32(load32(m.memory[uint32(v1):]))
			t4 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t5 := m.fn284(t3, t4, i32(1051600), v2+i32(8))
			v1 = t5
			goto l6
		case 1:
			store32(m.memory[int64(uint32(v2))+4:], uint32(v0))
			store32(m.memory[int64(uint32(v2))+12:], uint32(i32(147)))
			store32(m.memory[int64(uint32(v2))+8:], uint32(v2+i32(4)))
			t6 := int32(load32(m.memory[uint32(v1):]))
			t7 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t8 := m.fn284(t6, t7, i32(1052117), v2+i32(8))
			v1 = t8
			goto l6
		case 2:
			store32(m.memory[int64(uint32(v2))+4:], uint32(v0+i32(4)))
			store32(m.memory[int64(uint32(v2))+12:], uint32(i32(22)))
			store32(m.memory[int64(uint32(v2))+8:], uint32(v2+i32(4)))
			t9 := int32(load32(m.memory[uint32(v1):]))
			t10 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t11 := m.fn284(t9, t10, i32(1052141), v2+i32(8))
			v1 = t11
			goto l6
		case 3:
			t12 := int32(load32(m.memory[uint32(v1):]))
			t13 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t14 := int32(load32(m.memory[int64(uint32(t13))+12:]))
			t15 := m.t0[uint(t14)].(func(int32, int32, int32) int32)(t12, i32(1287111), i32(35))
			v1 = t15
			goto l6
		case 4:
			t16 := int32(load32(m.memory[uint32(v1):]))
			t17 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t18 := int32(load32(m.memory[int64(uint32(t17))+12:]))
			t19 := m.t0[uint(t18)].(func(int32, int32, int32) int32)(t16, i32(1287146), i32(30))
			v1 = t19
			goto l6
		case 5:
			store32(m.memory[int64(uint32(v2))+4:], uint32(v0+i32(4)))
			store32(m.memory[int64(uint32(v2))+12:], uint32(i32(148)))
			store32(m.memory[int64(uint32(v2))+8:], uint32(v2+i32(4)))
			t20 := int32(load32(m.memory[uint32(v1):]))
			t21 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t22 := m.fn284(t20, t21, i32(1052413), v2+i32(8))
			v1 = t22
		}
	}
l6:
	m.g0 = v2 + i32(16)
	return v1
}
func (m *Module) fn1050(v0, v1, v2, v3 int32) {
	var v4, v5, v6 int32
	var v7 int64
	var v8 int32
	var v9 int64
	var v10, v11 int32
	var v12 int64
	var v13, v14, v15, v16, v17, v18 int32
	t0 := m.g0
	v4 = t0 - i32(544)
	m.g0 = v4
	m.fn631(v4+i32(16), v2, v3, i32(47))
	t1 := int32(load32(m.memory[int64(uint32(v4))+20:]))
	t2 := v4
	v5 = t1
	store32(m.memory[int64(uint32(t2))+28:], uint32(v5))
	t3 := int32(load32(m.memory[int64(uint32(v4))+16:]))
	t4 := v4
	v6 = t3
	store32(m.memory[int64(uint32(t4))+24:], uint32(v6))
	{
		{
			t5 := int32(load32(m.memory[int64(uint32(v1))+36:]))
			if t5 == 0 {
				goto l0
			}
			t6 := int64(load64(m.memory[int64(uint32(v1))+40:]))
			t7 := int64(load64(m.memory[int64(uint32(v1))+48:]))
			t8 := m.fn29(t6, t7, v6, v5)
			v7 = t8
			t9 := int32(load32(m.memory[int64(uint32(v1))+28:]))
			v8 = t9
			v2 = v8 & int32(v7)
			v9 = int64(uint64(v7)>>25) & i64(127) * i64(72340172838076673)
			t10 := int32(load32(m.memory[int64(uint32(v1))+24:]))
			v10 = t10
			v11 = i32(0)
		l6:
			{
				t11 := int64(load64(m.memory[uint32(v10+v2):]))
				v12 = t11
				v7 = v12 ^ v9
				v7 = (v7 ^ i64(-1)) & (v7 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
				{
				l3:
					if v7 == 0 {
						if !(v12&(v12<<1)&i64(-0x7f7f7f7f7f7f7f80) == 0) {
							goto l0
						}
						t21 := v2
						v11 = v11 + i32(8)
						v2 = (t21 + v11) & v8
						goto l6
					}
					{
						t12 := v6
						t13 := v5
						v3 = v10 + (i32(0)-(int32(uint32(int64(bits.TrailingZeros64(uint64(v7))))>>3)+v2)&v8)*i32(20)
						t14 := int32(load32(m.memory[uint32(v3+i32(-16)):]))
						t15 := int32(load32(m.memory[uint32(v3+i32(-12)):]))
						t16 := m.fn15(t12, t13, t14, t15)
						if t16 != 0 {
							t17 := int32(load32(m.memory[uint32(v3+i32(-4)):]))
							v10 = t17
							t18 := int32(load32(m.memory[uint32(v3+i32(-8)):]))
							v3 = t18
							t19 := int32(load32(m.memory[uint32(v3):]))
							t20 := v3
							v2 = t19 + i32(1)
							store32(m.memory[uint32(t20):], uint32(v2))
							if v2 == 0 {
								goto l4
							}
							store32(m.memory[int64(uint32(v0))+8:], uint32(v10))
							store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
							store32(m.memory[uint32(v0):], uint32(i32(-1)))
							goto l5
						}
						v7 = (v7 + i64(-1)) & v7
						goto l3
					}
				}
			}
		}
	l0:
		m.fn114(v4+i32(240), v1, v6, v5)
		{
			{
				{
					t22 := int64(load64(m.memory[int64(uint32(v4))+240:]))
					if t22 != i64(-1) {
						memory_copy(m.memory, uint32(v4+i32(32)), uint32(v4+i32(240)), uint32(i32(208)))
						{
							{
								t29 := int64(load64(m.memory[int64(uint32(v4))+32:]))
								if t29 != i64(2) {
									goto l9
								}
								t30 := int32(load32(m.memory[int64(uint32(v4))+40:]))
								t31 := int64(load64(m.memory[int64(uint32(t30))+72:]))
								v7 = t31
								if uint64(v7) <= uint64(i64(0x8000000)) {
									goto l10
								}
								goto l11
							}
						l9:
							t32 := int64(load64(m.memory[int64(uint32(v4))+104:]))
							v7 = t32
							if uint64(v7) > uint64(i64(0x8000000)) {
								goto l11
							}
						}
					l10:
						t33 := int64(load64(m.memory[int64(uint32(v1))+56:]))
						v7 = t33
						v11 = i32(0)
						store32(m.memory[int64(uint32(v4))+460:], uint32(i32(0)))
						store64(m.memory[int64(uint32(v4))+452:], uint64(i64(0x100000000)))
						t34 := v4
						v7 = i64(0x20000000) - v7
						p35 := v7
						if uint64(v7) > uint64(i64(0x20000000)) {
							p35 = i64(0)
						}
						v9 = p35
						p36 := i64(0x8000000)
						if uint64(v9) < uint64(i64(0x8000000)) {
							p36 = v9
						}
						v12 = p36
						v7 = v12 + i64(1)
						store64(m.memory[int64(uint32(t34))+488:], uint64(v7))
						store64(m.memory[int64(uint32(v4))+480:], uint64(v7))
						store32(m.memory[int64(uint32(v4))+496:], uint32(v4+i32(32)))
						m.fn963(v4+i32(240), v4+i32(480), v4+i32(452))
						{
							t37 := int32(m.memory[int64(uint32(v4))+240])
							if t37 == i32(255) {
								t39 := int32(load32(m.memory[int64(uint32(v4))+244:]))
								if t39 == 0 {
									goto l14
								}
								v13 = i32(8192)
								t40 := int32(load32(m.memory[int64(uint32(v4))+452:]))
								v14 = t40
								t41 := int32(load32(m.memory[int64(uint32(v4))+460:]))
								v11 = t41
							l31:
								{
									if v14|v11 != 0 {
										goto l15
									}
									m.fn963(v4+i32(240), v4+i32(480), v4+i32(452))
									{
										t42 := int32(m.memory[int64(uint32(v4))+240])
										if t42 == i32(255) {
											goto l16
										}
										t43 := int64(load64(m.memory[int64(uint32(v4))+240:]))
										v7 = t43
										v11 = int32(int64(uint64(v7) >> 32))
										goto l13
									}
								l16:
									t44 := int32(load32(m.memory[int64(uint32(v4))+460:]))
									v11 = t44
									t45 := int32(load32(m.memory[int64(uint32(v4))+244:]))
									if t45 == 0 {
										goto l14
									}
									t46 := int32(load32(m.memory[int64(uint32(v4))+452:]))
									v14 = t46
								}
							l15:
								{
									{
										if v11 != v14 {
											goto l17
										}
										t47 := m.fn351(v4+i32(452), v14, i32(32))
										if t47 != i32(-1) {
											v3 = i32(1)
											v11 = i32(0)
											v7 = i64(9728)
											goto l33
										}
										t48 := int32(load32(m.memory[int64(uint32(v4))+452:]))
										v14 = t48
										t49 := int32(load32(m.memory[int64(uint32(v4))+460:]))
										v11 = t49
									}
								l17:
									t50 := int32(load32(m.memory[int64(uint32(v4))+456:]))
									v2 = t50
									v3 = i32(0)
									m.memory[int64(uint32(v4))+516] = byte(i32(0))
									store32(m.memory[int64(uint32(v4))+512:], uint32(i32(0)))
									store32(m.memory[int64(uint32(v4))+504:], uint32(v2+v11))
									t51 := v4
									t52 := v13
									v15 = v14 - v11
									p53 := v15
									if uint32(v13) < uint32(v15) {
										p53 = t52
									}
									v16 = p53
									store32(m.memory[int64(uint32(t51))+508:], uint32(v16))
									t54 := int64(load64(m.memory[int64(uint32(v4))+488:]))
									v7 = t54
									t55 := int32(load32(m.memory[int64(uint32(v4))+496:]))
									v8 = t55
								l32:
									{
										if v7 != i64(0) {
											goto l19
										}
										m.memory[int64(uint32(v4))+520] = byte(i32(255))
										store64(m.memory[int64(uint32(v4))+488:], uint64(i64(0)))
										v2 = i32(1)
										v10 = v3
										goto l20
									l19:
										{
											{
												t56 := int32(load32(m.memory[int64(uint32(v4))+508:]))
												t57 := v7
												v10 = t56 - v3
												if uint64(t57) < uint64(uint32(v10)) {
													goto l21
												}
												m.fn295(v4+i32(520), v8, v4+i32(504))
												t58 := int32(load32(m.memory[int64(uint32(v4))+512:]))
												v10 = t58
												v2 = v10 - v3
												goto l22
											}
										l21:
											t59 := int32(m.memory[int64(uint32(v4))+516])
											v2 = t59
											m.memory[int64(uint32(v4))+252] = byte(i32(0))
											store32(m.memory[int64(uint32(v4))+248:], uint32(i32(0)))
											t60 := v4
											v17 = int32(v7)
											store32(m.memory[int64(uint32(t60))+244:], uint32(v17))
											t61 := int32(load32(m.memory[int64(uint32(v4))+504:]))
											t62 := v4
											v18 = t61 + v3
											store32(m.memory[int64(uint32(t62))+240:], uint32(v18))
											{
												{
													if v2 != 0 {
														goto l23
													}
													m.fn295(v4+i32(520), v8, v4+i32(240))
													t63 := int32(load32(m.memory[int64(uint32(v4))+248:]))
													v2 = t63
													t64 := int32(m.memory[int64(uint32(v4))+252])
													if t64 == 0 {
														goto l24
													}
													m.fn1094(v18+v17, v10-v17)
													m.memory[int64(uint32(v4))+516] = byte(i32(1))
													goto l24
												}
											l23:
												m.memory[int64(uint32(v4))+252] = byte(i32(1))
												m.fn295(v4+i32(520), v8, v4+i32(240))
												t65 := int32(load32(m.memory[int64(uint32(v4))+248:]))
												v2 = t65
											}
										l24:
											t66 := v4
											v10 = v3 + v2
											store32(m.memory[int64(uint32(t66))+512:], uint32(v10))
										}
									l22:
										v7 = v7 - int64(uint32(v2))
										{
											t67 := int32(m.memory[int64(uint32(v4))+520])
											var p68 int32
											if t67 == i32(255) {
												p68 = 1
											}
											v2 = p68
											if v2 != 0 {
												goto l25
											}
											t69 := m.fn313(v4 + i32(520))
											if t69 != 0 {
												t73 := int32(load32(m.memory[int64(uint32(v4))+520:]))
												t74 := int32(load32(m.memory[int64(uint32(v4))+524:]))
												m.fn119(t73, t74)
												v3 = v10
												goto l32
											}
										}
									l25:
										store64(m.memory[int64(uint32(v4))+488:], uint64(v7))
									l20:
										t70 := v4
										v11 = v11 + v10
										store32(m.memory[int64(uint32(t70))+460:], uint32(v11))
										{
											if v2 != 0 {
												if v10 != 0 {
													{
														t72 := int32(m.memory[int64(uint32(v4))+516])
														if t72&i32(1) == 0 {
															goto l30
														}
														if uint32(v15) < uint32(v13) {
															goto l31
														}
														if v10 != v16 {
															goto l31
														}
														if v13 <= i32(-1) {
															goto l30
														}
														v13 = v13 << 1
														goto l31
													}
												l30:
													v13 = i32(-1)
													goto l31
												}
												v3 = i32(255)
												v7 = i64(0)
												goto l29
											}
											t71 := int64(load64(m.memory[int64(uint32(v4))+520:]))
											v7 = t71
											v11 = int32(int64(uint64(v7) >> 32))
											goto l13
										}
									}
								}
							}
							t38 := int64(load64(m.memory[int64(uint32(v4))+240:]))
							v7 = t38
							v11 = int32(int64(uint64(v7) >> 32))
							goto l13
						}
					}
					v3 = v4 + i32(248)
					t23 := int32(load32(m.memory[int64(uint32(v4))+248:]))
					if t23 == i32(-0x7ffffffd) {
						store64(m.memory[uint32(v0):], uint64(i64(0xffffffff)))
						m.fn116(v3)
						goto l5
					}
					t24 := int32(load32(m.memory[int64(uint32(v3))+8:]))
					store32(m.memory[int64(uint32(v4))+512:], uint32(t24))
					t25 := int64(load64(m.memory[uint32(v3):]))
					store64(m.memory[int64(uint32(v4))+504:], uint64(t25))
					m.fn51(v4+i32(492), v6, v5)
					store32(m.memory[int64(uint32(v4))+468:], uint32(i32(145)))
					store32(m.memory[int64(uint32(v4))+464:], uint32(v4+i32(504)))
					m.fn73(v4+i32(480), i32(1051162), v4+i32(464))
					t26 := int64(load64(m.memory[int64(uint32(v4))+496:]))
					store64(m.memory[int64(uint32(v0))+16:], uint64(t26))
					t27 := int64(load64(m.memory[int64(uint32(v4))+488:]))
					store64(m.memory[int64(uint32(v0))+8:], uint64(t27))
					t28 := int64(load64(m.memory[int64(uint32(v4))+480:]))
					store64(m.memory[uint32(v0):], uint64(t28))
					m.fn116(v4 + i32(504))
					goto l5
				}
			l13:
				v3 = int32(v7)
			l29:
				if v3&i32(255) == i32(255) {
					goto l14
				}
			l33:
				store64(m.memory[int64(uint32(v4))+520:], uint64(int64(uint32(v11))<<32|v7&i64(0xffffff00)|int64(uint32(v3))&i64(255)))
				m.fn51(v4+i32(252), v6, v5)
				store32(m.memory[int64(uint32(v4))+508:], uint32(i32(11)))
				store32(m.memory[int64(uint32(v4))+504:], uint32(v4+i32(520)))
				m.fn73(v4+i32(240), i32(1051136), v4+i32(504))
				t75 := int32(m.memory[int64(uint32(v4))+520])
				t76 := int32(load32(m.memory[int64(uint32(v4))+524:]))
				m.fn119(t75, t76)
				t77 := int64(load64(m.memory[int64(uint32(v4))+248:]))
				store64(m.memory[int64(uint32(v4))+464:], uint64(t77))
				t78 := int64(load64(m.memory[int64(uint32(v4))+256:]))
				store64(m.memory[int64(uint32(v4))+472:], uint64(t78))
				t79 := int32(load32(m.memory[int64(uint32(v4))+244:]))
				v11 = t79
				t80 := int32(load32(m.memory[int64(uint32(v4))+240:]))
				v3 = t80
				if v3 == i32(-1) {
					goto l14
				}
				t81 := int64(load64(m.memory[int64(uint32(v4))+472:]))
				store64(m.memory[int64(uint32(v0))+16:], uint64(t81))
				t82 := int64(load64(m.memory[int64(uint32(v4))+464:]))
				store64(m.memory[int64(uint32(v0))+8:], uint64(t82))
				store32(m.memory[int64(uint32(v0))+4:], uint32(v11))
				store32(m.memory[uint32(v0):], uint32(v3))
				goto l34
			}
		l14:
			{
				t83 := v12
				v7 = int64(uint32(v11))
				if uint64(t83) < uint64(v7) {
					goto l35
				}
				t84 := int64(load64(m.memory[int64(uint32(v1))+56:]))
				store64(m.memory[int64(uint32(v1))+56:], uint64(t84+v7))
				t85 := int32(load32(m.memory[int64(uint32(v4))+452:]))
				v10 = t85
				t86 := int32(load32(m.memory[int64(uint32(v4))+456:]))
				v8 = t86
				t87 := int32(load32(m.memory[int64(uint32(v4))+460:]))
				v2 = t87
				t88 := m.fn96(v2)
				v11 = t88
				t89 := m.fn96(v2)
				m.fn247(v4+i32(8), i32(4), t89)
				t90 := int32(load32(m.memory[int64(uint32(v4))+8:]))
				v3 = t90
				if v3 == 0 {
					m.fn85(i32(4), v11)
					panic("unreachable")
				}
				store64(m.memory[uint32(v3):], uint64(i64(0x100000001)))
				if v2 == 0 {
					goto l37
				}
				memory_copy(m.memory, uint32(v3+i32(8)), uint32(v8), uint32(v2))
			l37:
				if v10 != 0 {
					goto l38
				}
				v10 = i32(0)
				v11 = v4 + i32(240)
				goto l39
			l38:
				store32(m.memory[int64(uint32(v4))+240:], uint32(i32(1)))
				v11 = v4 + i32(480)
			l39:
				store32(m.memory[uint32(v11):], uint32(v10))
				{
					t91 := int32(load32(m.memory[int64(uint32(v4))+240:]))
					v10 = t91
					if v10 == 0 {
						goto l40
					}
					t92 := int32(load32(m.memory[int64(uint32(v4))+480:]))
					m.fn40(v8, v10, t92)
				}
			l40:
				m.fn51(v4+i32(240), v6, v5)
				t93 := int32(load32(m.memory[uint32(v3):]))
				t94 := v3
				v10 = t93 + i32(1)
				store32(m.memory[uint32(t94):], uint32(v10))
				if v10 == 0 {
					goto l4
				}
				v10 = v1 + i32(24)
				t95 := int64(load64(m.memory[int64(uint32(v1))+40:]))
				t96 := int64(load64(m.memory[int64(uint32(v1))+48:]))
				t97 := int32(load32(m.memory[int64(uint32(v4))+244:]))
				t98 := int32(load32(m.memory[int64(uint32(v4))+248:]))
				t99 := m.fn540(t95, t96, t97, t98)
				v7 = t99
				store32(m.memory[int64(uint32(v4))+504:], uint32(v4+i32(240)))
				{
					t100 := int32(load32(m.memory[int64(uint32(v1))+32:]))
					if t100 != 0 {
						goto l41
					}
					_ = m.fn658(v10, v1+i32(40))
				}
			l41:
				store32(m.memory[int64(uint32(v4))+484:], uint32(v10))
				store32(m.memory[int64(uint32(v4))+480:], uint32(v4+i32(504)))
				t102 := int32(load32(m.memory[int64(uint32(v1))+24:]))
				t103 := int32(load32(m.memory[int64(uint32(v1))+28:]))
				m.fn69(v4, t102, t103, v7, v4+i32(480), i32(149))
				t104 := int32(load32(m.memory[int64(uint32(v1))+24:]))
				v10 = t104
				t105 := int32(load32(m.memory[int64(uint32(v4))+4:]))
				v8 = t105
				{
					{
						t106 := int32(load32(m.memory[uint32(v4):]))
						if t106 != i32(1) {
							goto l42
						}
						v5 = v10 + v8
						t107 := int32(m.memory[uint32(v5)])
						v6 = t107
						t108 := int32(load32(m.memory[int64(uint32(v4))+248:]))
						v11 = t108
						t109 := int64(load64(m.memory[int64(uint32(v4))+240:]))
						v9 = t109
						t110 := v5
						v17 = int32(uint32(int32(v7)) >> 25)
						m.memory[uint32(t110)] = byte(v17)
						t111 := int32(load32(m.memory[int64(uint32(v1))+28:]))
						m.memory[uint32(v10+t111&(v8+i32(-8))+i32(8))] = byte(v17)
						t112 := int32(load32(m.memory[int64(uint32(v1))+36:]))
						store32(m.memory[int64(uint32(v1))+36:], uint32(t112+i32(1)))
						t113 := int32(load32(m.memory[int64(uint32(v1))+32:]))
						store32(m.memory[int64(uint32(v1))+32:], uint32(t113-v6&i32(1)))
						v1 = i32(0)
						v10 = v10 + (i32(0)-v8)*i32(20)
						v8 = v10 + i32(-20)
						store64(m.memory[uint32(v8):], uint64(v9))
						store32(m.memory[int64(uint32(v8))+8:], uint32(v11))
						store32(m.memory[uint32(v10+i32(-4)):], uint32(v2))
						store32(m.memory[uint32(v10+i32(-8)):], uint32(v3))
						goto l43
					}
				l42:
					v1 = v10 + (i32(0)-v8)*i32(20)
					v8 = v1 + i32(-4)
					t114 := int32(load32(m.memory[uint32(v8):]))
					v10 = t114
					store32(m.memory[uint32(v8):], uint32(v2))
					v8 = v1 + i32(-8)
					t115 := int32(load32(m.memory[uint32(v8):]))
					v1 = t115
					store32(m.memory[uint32(v8):], uint32(v3))
					t116 := int32(load32(m.memory[int64(uint32(v4))+240:]))
					t117 := int32(load32(m.memory[int64(uint32(v4))+244:]))
					m.fn16(t116, t117)
				}
			l43:
				store32(m.memory[int64(uint32(v4))+484:], uint32(v10))
				store32(m.memory[int64(uint32(v4))+480:], uint32(v1))
				m.fn1051(v4 + i32(480))
				store32(m.memory[int64(uint32(v0))+8:], uint32(v2))
				store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
				store32(m.memory[uint32(v0):], uint32(i32(-1)))
				m.fn124(v4 + i32(32))
				goto l5
			}
		l35:
			if uint64(v9) < uint64(i64(0x8000000)) {
				goto l44
			}
			store32(m.memory[int64(uint32(v4))+484:], uint32(i32(1)))
			store32(m.memory[int64(uint32(v4))+480:], uint32(v4+i32(24)))
			m.fn73(v4+i32(240), i32(1067229), v4+i32(480))
			v3 = i32(1072424)
			goto l45
		l44:
			store32(m.memory[int64(uint32(v4))+484:], uint32(i32(1)))
			store32(m.memory[int64(uint32(v4))+480:], uint32(v4+i32(24)))
			m.fn73(v4+i32(240), i32(1053035), v4+i32(480))
			v3 = i32(1075223)
		l45:
			store32(m.memory[uint32(v0):], uint32(i32(-0x7ffffffd)))
			t118 := int64(load64(m.memory[int64(uint32(v4))+240:]))
			store64(m.memory[int64(uint32(v0))+4:], uint64(t118))
			t119 := int32(load32(m.memory[int64(uint32(v4))+248:]))
			store32(m.memory[int64(uint32(v0))+12:], uint32(t119))
			store32(m.memory[int64(uint32(v0))+20:], uint32(i32(15)))
			store32(m.memory[int64(uint32(v0))+16:], uint32(v3))
		}
	l34:
		t120 := int32(load32(m.memory[int64(uint32(v4))+452:]))
		t121 := int32(load32(m.memory[int64(uint32(v4))+456:]))
		m.fn16(t120, t121)
		goto l46
	}
l4:
	panic("unreachable")
l11:
	store64(m.memory[int64(uint32(v4))+480:], uint64(v7))
	store32(m.memory[int64(uint32(v4))+252:], uint32(i32(28)))
	store32(m.memory[int64(uint32(v4))+244:], uint32(i32(1)))
	store32(m.memory[int64(uint32(v4))+248:], uint32(v4+i32(480)))
	store32(m.memory[int64(uint32(v4))+240:], uint32(v4+i32(24)))
	m.fn73(v0+i32(4), i32(1066768), v4+i32(240))
	store32(m.memory[int64(uint32(v0))+20:], uint32(i32(15)))
	store32(m.memory[int64(uint32(v0))+16:], uint32(i32(1072424)))
	store32(m.memory[uint32(v0):], uint32(i32(-0x7ffffffd)))
l46:
	m.fn124(v4 + i32(32))
l5:
	m.g0 = v4 + i32(544)
}
func (m *Module) fn1051(v0 int32) {
	t0 := int32(load32(m.memory[uint32(v0):]))
	if t0 == 0 {
		return
	}
	m.fn754(v0)
}
func (m *Module) fn1052(v0, v1 int32) int32 {
	var v2 int32
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[uint32(t0):]))
	v2 = t1
	t2 := int32(load32(m.memory[uint32(v2+i32(4)):]))
	t3 := int32(load32(m.memory[uint32(v2+i32(8)):]))
	t4 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t5 := int32(load32(m.memory[uint32(t4):]))
	v0 = t5 + (i32(0)-v1)*i32(40)
	t6 := int32(load32(m.memory[uint32(v0+i32(-36)):]))
	t7 := int32(load32(m.memory[uint32(v0+i32(-32)):]))
	t8 := m.fn545(t2, t3, t6, t7)
	return t8
}
func (m *Module) fn1053(v0, v1, v2 int32) {
	var v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14, v15, v16, v17, v18, v19, v20, v21, v22, v23 int32
	var v24, v25 int64
	var v26 int32
	t0 := m.g0
	v3 = t0 - i32(576)
	m.g0 = v3
	{
		if uint32(v2) <= uint32(i32(1)) {
			goto l0
		}
		{
			t1 := int32(m.memory[uint32(v1)])
			v4 = t1
			switch v4 + i32(-254) {
			default:
				if v2 == i32(2) {
					goto l0
				}
				if v4 != i32(239) {
					goto l0
				}
				t2 := int32(m.memory[int64(uint32(v1))+1])
				if t2 != i32(187) {
					goto l0
				}
				t3 := int32(m.memory[int64(uint32(v1))+2])
				if t3 != i32(191) {
					goto l0
				}
				store32(m.memory[int64(uint32(v3))+180:], uint32(i32(-1)))
				store32(m.memory[int64(uint32(v3))+188:], uint32(v2+i32(-3)))
				store32(m.memory[int64(uint32(v3))+184:], uint32(v1+i32(3)))
				goto l4
			case 1:
				t4 := int32(m.memory[int64(uint32(v1))+1])
				if t4 != i32(254) {
					goto l0
				}
				m.fn510(v3+i32(192), i32(1153092), v1, v2)
				t5 := int32(load32(m.memory[int64(uint32(v3))+200:]))
				store32(m.memory[int64(uint32(v3))+504:], uint32(t5))
				t6 := int64(load64(m.memory[int64(uint32(v3))+192:]))
				store64(m.memory[int64(uint32(v3))+496:], uint64(t6))
				m.fn490(v3+i32(180), v3+i32(496))
				goto l4
			case 0:
				t7 := int32(m.memory[int64(uint32(v1))+1])
				if t7 != i32(255) {
					goto l0
				}
				m.fn510(v3+i32(192), i32(1153064), v1, v2)
				t8 := int32(load32(m.memory[int64(uint32(v3))+200:]))
				store32(m.memory[int64(uint32(v3))+504:], uint32(t8))
				t9 := int64(load64(m.memory[int64(uint32(v3))+192:]))
				store64(m.memory[int64(uint32(v3))+496:], uint64(t9))
				m.fn490(v3+i32(180), v3+i32(496))
				goto l4
			}
		}
	l0:
		t11 := v3 + i32(192)
		t12 := v1
		p10 := i32(200)
		if uint32(v2) < uint32(i32(200)) {
			p10 = v2
		}
		m.fn12(t11, t12, p10)
		{
			t13 := int32(load32(m.memory[int64(uint32(v3))+192:]))
			if t13 != 0 {
				goto l5
			}
			t14 := int32(load32(m.memory[int64(uint32(v3))+196:]))
			t15 := v3 + i32(168)
			v4 = t14
			t16 := int32(load32(m.memory[int64(uint32(v3))+200:]))
			t17 := v4
			v5 = t16
			m.fn1055(t15, t17, v5, i32(1282576), i32(8))
			t18 := int32(load32(m.memory[int64(uint32(v3))+168:]))
			if t18 != i32(1) {
				goto l5
			}
			t19 := int32(load32(m.memory[int64(uint32(v3))+172:]))
			m.fn826(v3+i32(160), t19+i32(8), v4, v5, i32(1084464))
			t20 := int32(load32(m.memory[int64(uint32(v3))+160:]))
			t21 := int32(load32(m.memory[int64(uint32(v3))+164:]))
			m.fn824(v3+i32(152), t20, t21)
			t22 := int32(load32(m.memory[int64(uint32(v3))+152:]))
			t23 := int32(load32(m.memory[int64(uint32(v3))+156:]))
			m.fn13(v3+i32(144), t22, t23, i32(61))
			t24 := int32(load32(m.memory[int64(uint32(v3))+144:]))
			v4 = t24
			if v4 == 0 {
				goto l5
			}
			t25 := int32(load32(m.memory[int64(uint32(v3))+148:]))
			m.fn824(v3+i32(136), v4, t25)
			t26 := int32(load32(m.memory[int64(uint32(v3))+136:]))
			t27 := v3
			v4 = t26
			t28 := int32(load32(m.memory[int64(uint32(v3))+140:]))
			t29 := v4
			v6 = t28
			store32(m.memory[int64(uint32(t27))+196:], uint32(t29+v6))
			store32(m.memory[int64(uint32(v3))+192:], uint32(v4))
			{
				t30 := m.fn48(v3 + i32(192))
				v5 = t30
				if v5 == i32(39) {
					goto l6
				}
				if v5 != i32(34) {
					goto l5
				}
			}
		l6:
			m.fn826(v3+i32(128), i32(1), v4, v6, i32(1084480))
			t31 := int32(load32(m.memory[int64(uint32(v3))+128:]))
			t32 := v3 + i32(120)
			v4 = t31
			t33 := int32(load32(m.memory[int64(uint32(v3))+132:]))
			t34 := v4
			v6 = t33
			m.fn1056(t32, t34, v6, v5)
			t35 := int32(load32(m.memory[int64(uint32(v3))+120:]))
			if t35 != i32(1) {
				goto l5
			}
			t36 := int32(load32(m.memory[int64(uint32(v3))+124:]))
			m.fn825(v3+i32(112), v4, v6, t36, i32(1084496))
			t37 := int32(load32(m.memory[int64(uint32(v3))+112:]))
			t38 := int32(load32(m.memory[int64(uint32(v3))+116:]))
			m.fn51(v3+i32(340), t37, t38)
			t39 := int32(load32(m.memory[int64(uint32(v3))+340:]))
			v4 = t39
			if v4 == i32(-1) {
				goto l5
			}
			{
				t40 := int32(load32(m.memory[int64(uint32(v3))+344:]))
				v6 = t40
				t41 := int32(load32(m.memory[int64(uint32(v3))+348:]))
				t42 := m.fn1057(v6, t41)
				v5 = t42
				if v5 == 0 {
					goto l7
				}
				if v5 != i32(1148960) {
					goto l8
				}
			}
		l7:
			m.fn16(v4, v6)
		}
	l5:
		store32(m.memory[int64(uint32(v3))+188:], uint32(v2))
		store32(m.memory[int64(uint32(v3))+184:], uint32(v1))
		store32(m.memory[int64(uint32(v3))+180:], uint32(i32(-1)))
		goto l4
	l8:
		m.fn510(v3+i32(192), v5, v1, v2)
		t43 := int32(load32(m.memory[int64(uint32(v3))+200:]))
		store32(m.memory[int64(uint32(v3))+504:], uint32(t43))
		t44 := int64(load64(m.memory[int64(uint32(v3))+192:]))
		store64(m.memory[int64(uint32(v3))+496:], uint64(t44))
		m.fn490(v3+i32(180), v3+i32(496))
		m.fn16(v4, v6)
	}
l4:
	t45 := int32(load32(m.memory[int64(uint32(v3))+184:]))
	v6 = t45
	t46 := int32(load32(m.memory[int64(uint32(v3))+188:]))
	v7 = t46
	store32(m.memory[int64(uint32(v3))+348:], uint32(i32(0)))
	store64(m.memory[int64(uint32(v3))+340:], uint64(i64(0x100000000)))
	store32(m.memory[int64(uint32(v3))+504:], uint32(i32(0)))
	store64(m.memory[int64(uint32(v3))+496:], uint64(i64(0x400000000)))
	v2 = i32(-32)
l10:
	{
		if v2 == 0 {
			t54 := int32(load32(m.memory[int64(uint32(v3))+348:]))
			store32(m.memory[int64(uint32(v3))+272:], uint32(t54))
			t55 := int64(load64(m.memory[int64(uint32(v3))+340:]))
			store64(m.memory[int64(uint32(v3))+264:], uint64(t55))
			t56 := int64(load64(m.memory[int64(uint32(v3))+496:]))
			store64(m.memory[int64(uint32(v3))+276:], uint64(t56))
			t57 := int32(load32(m.memory[int64(uint32(v3))+504:]))
			store32(m.memory[int64(uint32(v3))+284:], uint32(t57))
			store16(m.memory[int64(uint32(v3))+292:], uint16(i32(0)))
			store32(m.memory[int64(uint32(v3))+288:], uint32(i32(256)))
			store64(m.memory[int64(uint32(v3))+208:], uint64(i64(0)))
			store64(m.memory[int64(uint32(v3))+200:], uint64(i64(0x10000000000)))
			store32(m.memory[int64(uint32(v3))+196:], uint32(i32(1148960)))
			store32(m.memory[int64(uint32(v3))+192:], uint32(i32(0)))
			store64(m.memory[int64(uint32(v3))+216:], uint64(i64(0)))
			store32(m.memory[int64(uint32(v3))+224:], uint32(i32(0)))
			m.memory[int64(uint32(v3))+296] = byte(i32(0))
			store32(m.memory[int64(uint32(v3))+260:], uint32(v7))
			store32(m.memory[int64(uint32(v3))+256:], uint32(v6))
			m.memory[int64(uint32(v3))+248] = byte(i32(0))
			store32(m.memory[int64(uint32(v3))+244:], uint32(i32(0)))
			store64(m.memory[int64(uint32(v3))+236:], uint64(i64(0x400000000)))
			store64(m.memory[int64(uint32(v3))+228:], uint64(i64(1)))
			m.fn22(v3+i32(496), i32(3))
			t58 := int64(load64(m.memory[int64(uint32(i32(0)))+1286256:]))
			store64(m.memory[int64(uint32(v3))+304:], uint64(t58))
			t59 := int64(load64(m.memory[int64(uint32(i32(0)))+1286264:]))
			store64(m.memory[int64(uint32(v3))+312:], uint64(t59))
			t60 := int64(load64(m.memory[int64(uint32(v3))+504:]))
			store64(m.memory[int64(uint32(v3))+328:], uint64(t60))
			t61 := int64(load64(m.memory[int64(uint32(v3))+496:]))
			store64(m.memory[int64(uint32(v3))+320:], uint64(t61))
			store64(m.memory[int64(uint32(v3))+364:], uint64(i64(0x400000000)))
			store64(m.memory[int64(uint32(v3))+356:], uint64(i64(4)))
			store64(m.memory[int64(uint32(v3))+348:], uint64(i64(0)))
			store64(m.memory[int64(uint32(v3))+340:], uint64(i64(0x100000000)))
			store64(m.memory[int64(uint32(v3))+372:], uint64(i64(0)))
			store32(m.memory[int64(uint32(v3))+392:], uint32(i32(0)))
			store64(m.memory[int64(uint32(v3))+384:], uint64(i64(0x400000000)))
			store32(m.memory[int64(uint32(v3))+396:], uint32(i32(0)))
			v8 = v3 + i32(508)
			v9 = v3 + i32(400) + i32(8)
			v10 = v3 + i32(544) + i32(4)
			v11 = v3 + i32(496) + i32(8)
			v12 = v3 + i32(496) + i32(4)
			v13 = v3 + i32(400) + i32(4)
			v14 = v3 + i32(256)
			v15 = v3 + i32(192) + i32(16)
			v16 = v3 + i32(264)
			v17 = i32(4)
			v18 = i32(0)
			v19 = i32(0)
			v20 = i32(0)
			v21 = i32(0)
		l173:
			{
				{
					t62 := int32(m.memory[int64(uint32(v3))+296])
					if t62 != i32(1) {
						goto l11
					}
					t63 := int32(load16(m.memory[int64(uint32(v3))+292:]))
					t64 := v3
					v2 = t63
					t65 := v2
					var p66 int32
					if v2 != i32(0) {
						p66 = 1
					}
					v4 = t65 - p66
					store16(m.memory[int64(uint32(t64))+292:], uint16(v4))
					t67 := int32(load32(m.memory[int64(uint32(v3))+284:]))
					v22 = t67
					v1 = v22 << 4
					v2 = i32(0) - v1
					t68 := int32(load32(m.memory[int64(uint32(v3))+280:]))
					v23 = t68
					v1 = v23 + v1
					v7 = v4 & i32(0xffff)
					v5 = v22
				l13:
					{
						v4 = v5
						if v2 == 0 {
							goto l12
						}
						v6 = v1 + i32(-4)
						v2 = v2 + i32(16)
						v5 = v4 + i32(-1)
						v1 = v1 + i32(-16)
						t69 := int32(load16(m.memory[uint32(v6):]))
						if uint32(t69) > uint32(v7) {
							goto l13
						}
					}
					if uint32(v4) >= uint32(v22) {
						goto l14
					}
					{
						t70 := int32(load32(m.memory[uint32(v23+v4<<4):]))
						v2 = t70
						t71 := int32(load32(m.memory[int64(uint32(v3))+272:]))
						if uint32(v2) > uint32(t71) {
							goto l15
						}
						store32(m.memory[int64(uint32(v3))+272:], uint32(v2))
					}
				l15:
					if uint32(v4) <= uint32(v22) {
						goto l16
					}
					goto l14
				l12:
					v4 = i32(0)
					store32(m.memory[int64(uint32(v3))+272:], uint32(i32(0)))
				l16:
					store32(m.memory[int64(uint32(v3))+284:], uint32(v4))
				l14:
					m.memory[int64(uint32(v3))+296] = byte(i32(0))
				}
			l11:
				t72 := int32(m.memory[int64(uint32(v3))+248])
				v2 = t72
				{
					{
						{
							{
								{
									{
										{
											{
												{
												l74:
													{
														{
															{
																{
																	switch v2 & i32(255) {
																	case 4:
																		m.memory[int64(uint32(v3))+248] = byte(i32(3))
																		{
																			t156 := int32(load32(m.memory[int64(uint32(v3))+244:]))
																			v2 = t156
																			if v2 == 0 {
																				m.fn153(i32(1282108))
																				panic("unreachable")
																			}
																			t157 := v3
																			v2 = v2 + i32(-1)
																			store32(m.memory[int64(uint32(t157))+244:], uint32(v2))
																			t158 := int32(load32(m.memory[int64(uint32(v3))+232:]))
																			v1 = t158
																			t159 := int32(load32(m.memory[int64(uint32(v3))+240:]))
																			t160 := int32(load32(m.memory[uint32(t159+v2<<2):]))
																			t161 := v1
																			v2 = t160
																			if uint32(t161) < uint32(v2) {
																				m.fn99(v2, v1, i32(1282124))
																				panic("unreachable")
																			}
																			t162 := v3 + i32(104)
																			v1 = v1 - v2
																			m.fn1064(t162, v1)
																			store32(m.memory[int64(uint32(v3))+232:], uint32(v2))
																			t163 := int32(load32(m.memory[int64(uint32(v3))+104:]))
																			v5 = t163
																			t164 := int32(load32(m.memory[int64(uint32(v3))+108:]))
																			v4 = t164
																			{
																				if v1 == 0 {
																					goto l66
																				}
																				t165 := int32(load32(m.memory[int64(uint32(v3))+228:]))
																				memory_copy(m.memory, uint32(v4), uint32(t165+v2), uint32(v1))
																			}
																		l66:
																			store32(m.memory[int64(uint32(v3))+512:], uint32(v1))
																			store32(m.memory[int64(uint32(v3))+508:], uint32(v4))
																			store32(m.memory[int64(uint32(v3))+504:], uint32(v5))
																			store64(m.memory[int64(uint32(v3))+496:], uint64(i64(0x100000000)))
																			v2 = i32(1)
																			goto l67
																		}
																	default:
																		v2 = i32(3)
																		t73 := int32(load32(m.memory[int64(uint32(v3))+256:]))
																		v5 = t73
																		t74 := int32(load32(m.memory[int64(uint32(v3))+260:]))
																		t75 := v5
																		v6 = t74
																		t76 := m.fn144(t75, v6)
																		v1 = t76 & i32(255)
																		if uint32(v1) > uint32(i32(5)) {
																			goto l23
																		}
																		v4 = i32(0)
																		v7 = i32_shl(i32(1), v1)
																		if v7&i32(21) != 0 {
																			goto l24
																		}
																		if v7&i32(40) != 0 {
																			goto l25
																		}
																		v4 = i32(3)
																		goto l24
																	case 1:
																		t77 := int64(load64(m.memory[int64(uint32(v3))+208:]))
																		v24 = t77
																		t78 := int32(load32(m.memory[int64(uint32(v3))+256:]))
																		t79 := v3 + i32(48)
																		v4 = t78
																		t80 := int32(load32(m.memory[int64(uint32(v3))+260:]))
																		t81 := v4
																		v2 = t80
																		m.fn148(t79, i32(1), t81, v2, i32(1284056))
																		t82 := int32(load32(m.memory[int64(uint32(v3))+48:]))
																		t83 := v3 + i32(40)
																		v1 = t82
																		t84 := int32(load32(m.memory[int64(uint32(v3))+52:]))
																		m.fn146(t83, i32(59), i32(38), i32(60), v1, v1+t84)
																		t85 := int32(load32(m.memory[int64(uint32(v3))+40:]))
																		if t85 != i32(1) {
																			goto l26
																		}
																		t86 := int32(load32(m.memory[int64(uint32(v3))+44:]))
																		v5 = t86 - v1
																		v1 = v5 + i32(1)
																		if uint32(v1) >= uint32(v2) {
																			m.fn158(v1, v2, i32(1284072))
																			panic("unreachable")
																		}
																		t87 := int32(m.memory[uint32(v4+v1)])
																		v6 = t87
																		if v6 == i32(59) {
																			t145 := v3 + i32(544)
																			t146 := v4
																			t147 := v2
																			v1 = v5 + i32(2)
																			m.fn309(t145, t146, t147, v1, i32(1284088))
																			m.memory[int64(uint32(v3))+248] = byte(i32(3))
																			t148 := int64(load64(m.memory[int64(uint32(v3))+552:]))
																			store64(m.memory[int64(uint32(v3))+256:], uint64(t148))
																			store64(m.memory[int64(uint32(v3))+208:], uint64(v24+int64(uint32(v1))))
																			{
																				t149 := int32(load32(m.memory[int64(uint32(v3))+548:]))
																				v2 = t149
																				if uint32(v2) <= uint32(i32(1)) {
																					m.fn151(i32(1), v2+i32(-1), v2, i32(1281456))
																					panic("unreachable")
																				}
																				t150 := int32(load32(m.memory[int64(uint32(v3))+544:]))
																				v1 = t150
																				t151 := int32(load32(m.memory[int64(uint32(v3))+196:]))
																				store32(m.memory[int64(uint32(v3))+516:], uint32(t151))
																				store32(m.memory[int64(uint32(v3))+504:], uint32(i32(-1)))
																				store64(m.memory[int64(uint32(v3))+496:], uint64(i64(0x900000000)))
																				store32(m.memory[int64(uint32(v3))+512:], uint32(v2+i32(-2)))
																				store32(m.memory[int64(uint32(v3))+508:], uint32(v1+i32(1)))
																				v2 = i32(9)
																				goto l63
																			}
																		}
																		m.fn309(v3+i32(544), v4, v2, v1, i32(1284104))
																		t88 := int64(load64(m.memory[int64(uint32(v3))+552:]))
																		store64(m.memory[int64(uint32(v3))+256:], uint64(t88))
																		store64(m.memory[int64(uint32(v3))+208:], uint64(v24+int64(uint32(v5))+i64(1)))
																		t89 := int32(load32(m.memory[int64(uint32(v3))+548:]))
																		v1 = t89
																		t90 := int32(load32(m.memory[int64(uint32(v3))+544:]))
																		v4 = t90
																		t91 := int32(m.memory[int64(uint32(v3))+200])
																		v2 = t91
																		if v6 != i32(38) {
																			m.memory[int64(uint32(v3))+248] = byte(i32(2))
																			if v2&i32(1) == 0 {
																				goto l30
																			}
																			t140 := int32(load32(m.memory[int64(uint32(v3))+196:]))
																			t141 := int32(m.memory[int64(uint32(v3))+207])
																			m.fn149(v11, t140, t141, v4, v1)
																			goto l31
																		}
																		if v2&i32(1) == 0 {
																			goto l30
																		}
																		t92 := int32(load32(m.memory[int64(uint32(v3))+196:]))
																		t93 := int32(m.memory[int64(uint32(v3))+207])
																		m.fn149(v11, t92, t93, v4, v1)
																		goto l31
																	case 2:
																		m.memory[int64(uint32(v3))+248] = byte(i32(3))
																		t94 := int64(load64(m.memory[int64(uint32(v3))+208:]))
																		v24 = t94
																		{
																			{
																				t95 := int32(load32(m.memory[int64(uint32(v3))+260:]))
																				v2 = t95
																				if uint32(v2) < uint32(i32(2)) {
																					m.memory[int64(uint32(v3))+504] = byte(i32(6))
																					store64(m.memory[int64(uint32(v3))+216:], uint64(v24))
																					goto l38
																				}
																				t96 := int32(load32(m.memory[int64(uint32(v3))+256:]))
																				v1 = t96
																				t97 := int32(m.memory[int64(uint32(v1))+1])
																				v4 = t97
																				if v4 == i32(33) {
																					v25 = i64(0)
																					if v2 == i32(2) {
																						goto l39
																					}
																					{
																						{
																							{
																								t102 := int32(m.memory[int64(uint32(v1))+2])
																								v4 = t102
																								if v4 == i32(45) {
																									store64(m.memory[int64(uint32(v3))+568:], uint64(i64(10)))
																									m.memory[int64(uint32(v3))+556] = byte(i32(62))
																									store32(m.memory[int64(uint32(v3))+552:], uint32(v1+v2))
																									store32(m.memory[int64(uint32(v3))+548:], uint32(v1))
																									store32(m.memory[int64(uint32(v3))+544:], uint32(v1))
																								l47:
																									{
																										m.fn155(v3+i32(64), v3+i32(544))
																										t106 := int32(load32(m.memory[int64(uint32(v3))+64:]))
																										if t106 != i32(1) {
																											v1 = i32(10)
																											goto l50
																										}
																										t107 := int32(load32(m.memory[int64(uint32(v3))+68:]))
																										v4 = t107
																										if uint32(v4) <= uint32(i32(5)) {
																											goto l47
																										}
																										if uint32(v4) > uint32(v2) {
																											m.fn151(i32(0), v4, v2, i32(1282428))
																											panic("unreachable")
																										}
																										t108 := m.fn1061(v1, v4, i32(1282444), i32(2))
																										if t108 == 0 {
																											goto l47
																										}
																									}
																									v5 = i32(10)
																									goto l45
																								}
																								if v4 == i32(68) {
																									goto l41
																								}
																								if v4 == i32(100) {
																									goto l41
																								}
																								if v4 != i32(91) {
																									goto l39
																								}
																								m.memory[int64(uint32(v3))+556] = byte(i32(62))
																								store32(m.memory[int64(uint32(v3))+548:], uint32(v1))
																								store32(m.memory[int64(uint32(v3))+544:], uint32(v1))
																								store32(m.memory[int64(uint32(v3))+552:], uint32(v1+v2))
																							l44:
																								{
																									m.fn155(v3+i32(56), v3+i32(544))
																									t103 := int32(load32(m.memory[int64(uint32(v3))+56:]))
																									if t103 != i32(1) {
																										v1 = i32(9)
																										goto l50
																									}
																									t104 := int32(load32(m.memory[int64(uint32(v3))+60:]))
																									v4 = t104
																									if uint32(v4) > uint32(v2) {
																										m.fn151(i32(0), v4, v2, i32(1282408))
																										panic("unreachable")
																									}
																									t105 := m.fn1061(v1, v4, i32(1282424), i32(2))
																									if t105 == 0 {
																										goto l44
																									}
																								}
																								v5 = i32(9)
																								goto l45
																							}
																						l41:
																							store64(m.memory[int64(uint32(v3))+568:], uint64(i64(0)))
																							m.fn157(v3+i32(72), v3+i32(568), i32(1), i32(0), v1, v2)
																							t109 := int32(load32(m.memory[int64(uint32(v3))+72:]))
																							if t109&i32(1) != 0 {
																								goto l49
																							}
																							t110 := int32(m.memory[int64(uint32(v3))+568])
																							v1 = t110
																							goto l50
																						}
																					l49:
																						t111 := int32(load32(m.memory[int64(uint32(v3))+76:]))
																						v4 = t111
																						t112 := int32(m.memory[int64(uint32(v3))+568])
																						v5 = t112
																					}
																				l45:
																					t113 := v3
																					t114 := v24
																					v4 = v4 + i32(1)
																					store64(m.memory[int64(uint32(t113))+208:], uint64(t114+int64(uint32(v4))))
																					m.fn309(v3+i32(544), v1, v2, v4, i32(1284040))
																					t115 := int64(load64(m.memory[int64(uint32(v3))+552:]))
																					store64(m.memory[int64(uint32(v3))+256:], uint64(t115))
																					t116 := int32(load32(m.memory[int64(uint32(v3))+544:]))
																					t117 := int32(load32(m.memory[int64(uint32(v3))+548:]))
																					m.fn163(v3+i32(496), v3+i32(192), v5, t116, t117)
																					goto l51
																				}
																				if v4 == i32(47) {
																					m.fn1060(v3+i32(544), v14, v15)
																					{
																						t118 := int32(load32(m.memory[int64(uint32(v3))+544:]))
																						if t118 == i32(-1) {
																							t122 := int32(load32(m.memory[int64(uint32(v3))+548:]))
																							t123 := int32(load32(m.memory[int64(uint32(v3))+552:]))
																							m.fn160(v3+i32(496), v3+i32(192), t122, t123)
																							goto l51
																						}
																						t119 := int64(load64(m.memory[int64(uint32(v3))+560:]))
																						store64(m.memory[int64(uint32(v12))+16:], uint64(t119))
																						t120 := int64(load64(m.memory[int64(uint32(v3))+552:]))
																						store64(m.memory[int64(uint32(v12))+8:], uint64(t120))
																						t121 := int64(load64(m.memory[int64(uint32(v3))+544:]))
																						store64(m.memory[uint32(v12):], uint64(t121))
																						store64(m.memory[int64(uint32(v3))+216:], uint64(v24))
																						store32(m.memory[int64(uint32(v3))+496:], uint32(i32(1)))
																						goto l51
																					}
																				}
																				if v4 == i32(63) {
																					m.memory[int64(uint32(v3))+568] = byte(i32(0))
																					m.fn1062(v3+i32(80), v3+i32(568), v1, v2)
																					{
																						t124 := int32(load32(m.memory[int64(uint32(v3))+80:]))
																						if t124&i32(1) != 0 {
																							t126 := int32(load32(m.memory[int64(uint32(v3))+84:]))
																							t127 := v3 + i32(544)
																							t128 := v1
																							t129 := v2
																							v4 = t126 + i32(1)
																							m.fn309(t127, t128, t129, v4, i32(1281600))
																							t130 := int64(load64(m.memory[int64(uint32(v3))+552:]))
																							store64(m.memory[int64(uint32(v3))+256:], uint64(t130))
																							store64(m.memory[int64(uint32(v3))+208:], uint64(v24+int64(uint32(v4))))
																							t131 := int32(load32(m.memory[int64(uint32(v3))+544:]))
																							t132 := int32(load32(m.memory[int64(uint32(v3))+548:]))
																							m.fn162(v3+i32(496), v3+i32(192), t131, t132)
																							goto l51
																						}
																						t125 := m.fn1063(v1, v2)
																						v1 = t125
																						store64(m.memory[int64(uint32(v3))+216:], uint64(v24))
																						store64(m.memory[int64(uint32(v3))+208:], uint64(v24+int64(uint32(v2))))
																						store64(m.memory[int64(uint32(v3))+496:], uint64(i64(-0x7ffffff6ffffffff)))
																						store32(m.memory[int64(uint32(v3))+504:], uint32(v1&i32(255)))
																						v2 = i32(-0x7ffffff7)
																						goto l54
																					}
																				}
																				m.fn1060(v3+i32(544), v14, v15)
																				t98 := int32(load32(m.memory[int64(uint32(v3))+544:]))
																				if t98 == i32(-1) {
																					goto l36
																				}
																				t99 := int64(load64(m.memory[int64(uint32(v3))+560:]))
																				store64(m.memory[int64(uint32(v12))+16:], uint64(t99))
																				t100 := int64(load64(m.memory[int64(uint32(v3))+552:]))
																				store64(m.memory[int64(uint32(v12))+8:], uint64(t100))
																				t101 := int64(load64(m.memory[int64(uint32(v3))+544:]))
																				store64(m.memory[uint32(v12):], uint64(t101))
																				store64(m.memory[int64(uint32(v3))+216:], uint64(v24))
																				v2 = i32(1)
																				goto l37
																			}
																		l36:
																			t133 := int32(load32(m.memory[int64(uint32(v3))+548:]))
																			t134 := int32(load32(m.memory[int64(uint32(v3))+552:]))
																			m.fn161(v12, v3+i32(192), t133, t134)
																			v2 = i32(0)
																		}
																	l37:
																		store32(m.memory[int64(uint32(v3))+496:], uint32(v2))
																	l51:
																		t135 := int32(load32(m.memory[int64(uint32(v3))+500:]))
																		v2 = t135
																		goto l54
																	case 3:
																		t136 := int32(load32(m.memory[int64(uint32(v3))+260:]))
																		v4 = t136
																		t137 := int32(load32(m.memory[int64(uint32(v3))+256:]))
																		v5 = t137
																		t138 := int32(m.memory[int64(uint32(v3))+206])
																		if t138 == 0 {
																			goto l55
																		}
																		v2 = i32(0)
																	l58:
																		if v4 != v2 {
																			t139 := int32(m.memory[uint32(v5+v2)])
																			v1 = t139 + i32(-9)
																			if uint32(v1) > uint32(i32(23)) {
																				goto l57
																			}
																			if i32_shl(i32(1), v1)&i32(8388627) == 0 {
																				goto l57
																			}
																			v2 = v2 + i32(1)
																			goto l58
																		}
																		v2 = v4
																		goto l57
																	case 5:
																		store64(m.memory[int64(uint32(v3))+496:], uint64(i64(0xa00000000)))
																		goto l59
																	}
																l26:
																	store64(m.memory[int64(uint32(v3))+256:], uint64(i64(1)))
																	m.memory[int64(uint32(v3))+248] = byte(i32(5))
																	store64(m.memory[int64(uint32(v3))+208:], uint64(v24+int64(uint32(v2))))
																	t142 := int32(m.memory[int64(uint32(v3))+200])
																	if t142 != 0 {
																		t143 := int32(load32(m.memory[int64(uint32(v3))+196:]))
																		t144 := int32(m.memory[int64(uint32(v3))+207])
																		m.fn149(v11, t143, t144, v4, v2)
																		goto l31
																	}
																}
															l30:
																store64(m.memory[int64(uint32(v3))+216:], uint64(v24))
																store64(m.memory[int64(uint32(v3))+496:], uint64(i64(-0x7ffffff8ffffffff)))
																v2 = i32(-0x7ffffff9)
																goto l61
															l25:
																v4 = i32(2)
															l24:
																m.fn148(v3, v4, v5, v6, i32(1284008))
																t152 := int64(load64(m.memory[uint32(v3):]))
																store64(m.memory[int64(uint32(v3))+256:], uint64(t152))
																t153 := int32(load32(m.memory[int64(uint32(v3))+192:]))
																if t153&i32(1) != 0 {
																	goto l23
																}
																store32(m.memory[int64(uint32(v3))+192:], uint32(i32(2)))
																t154 := int32(load32(m.memory[int64(uint32(v1<<2))+1301492:]))
																t155 := int32(load32(m.memory[uint32(t154):]))
																store32(m.memory[int64(uint32(v3))+196:], uint32(t155))
																goto l23
															}
														l57:
															t166 := int64(load64(m.memory[int64(uint32(v3))+208:]))
															store64(m.memory[int64(uint32(v3))+208:], uint64(t166+int64(uint32(v2))))
															m.fn148(v3+i32(96), v2, v5, v4, i32(1284024))
															t167 := int32(load32(m.memory[int64(uint32(v3))+100:]))
															t168 := v3
															v4 = t167
															store32(m.memory[int64(uint32(t168))+260:], uint32(v4))
															t169 := int32(load32(m.memory[int64(uint32(v3))+96:]))
															t170 := v3
															v5 = t169
															store32(m.memory[int64(uint32(t170))+256:], uint32(v5))
														}
													l55:
														m.fn1065(v3+i32(88), i32(60), i32(38), v5, v4)
														{
															t171 := int32(load32(m.memory[int64(uint32(v3))+88:]))
															if t171 != i32(1) {
																store64(m.memory[int64(uint32(v3))+256:], uint64(i64(1)))
																m.memory[int64(uint32(v3))+248] = byte(i32(5))
																t182 := int64(load64(m.memory[int64(uint32(v3))+208:]))
																store64(m.memory[int64(uint32(v3))+208:], uint64(t182+int64(uint32(v4))))
																t183 := int32(load32(m.memory[int64(uint32(v3))+196:]))
																t184 := int32(m.memory[int64(uint32(v3))+207])
																m.fn149(v3+i32(544), t183, t184, v5, v4)
																{
																	t185 := int32(load32(m.memory[int64(uint32(v3))+552:]))
																	if t185 == 0 {
																		store64(m.memory[int64(uint32(v3))+496:], uint64(i64(0xa00000000)))
																		t188 := int32(load32(m.memory[int64(uint32(v3))+544:]))
																		t189 := int32(load32(m.memory[int64(uint32(v3))+548:]))
																		m.fn1066(t188, t189)
																		goto l59
																	}
																	t186 := int64(load64(m.memory[int64(uint32(v3))+552:]))
																	store64(m.memory[int64(uint32(v11))+8:], uint64(t186))
																	t187 := int64(load64(m.memory[int64(uint32(v3))+544:]))
																	store64(m.memory[uint32(v11):], uint64(t187))
																	goto l31
																}
															}
															{
																t172 := int32(load32(m.memory[int64(uint32(v3))+92:]))
																v2 = t172
																if v2 != 0 {
																	if uint32(v2) < uint32(v4) {
																		v24 = int64(uint32(v2))
																		t173 := int32(m.memory[uint32(v5+v2)])
																		if t173 == i32(60) {
																			goto l72
																		}
																		m.fn309(v3+i32(544), v5, v4, v2, i32(1284168))
																		m.memory[int64(uint32(v3))+248] = byte(i32(1))
																		t174 := int64(load64(m.memory[int64(uint32(v3))+552:]))
																		store64(m.memory[int64(uint32(v3))+256:], uint64(t174))
																		t175 := int64(load64(m.memory[int64(uint32(v3))+208:]))
																		store64(m.memory[int64(uint32(v3))+208:], uint64(t175+v24))
																		t176 := int32(load32(m.memory[int64(uint32(v3))+196:]))
																		t177 := int32(m.memory[int64(uint32(v3))+207])
																		t178 := int32(load32(m.memory[int64(uint32(v3))+544:]))
																		t179 := int32(load32(m.memory[int64(uint32(v3))+548:]))
																		m.fn149(v11, t176, t177, t178, t179)
																		goto l31
																	}
																	m.fn158(v2, v4, i32(1284136))
																	panic("unreachable")
																}
																if v4 != 0 {
																	t180 := int32(m.memory[uint32(v5)])
																	p181 := i32(1)
																	if t180 == i32(60) {
																		p181 = i32(2)
																	}
																	v2 = p181
																	goto l23
																}
																m.fn158(i32(0), i32(0), i32(1284120))
																panic("unreachable")
															}
														}
													l72:
														m.fn309(v3+i32(544), v5, v4, v2, i32(1284152))
														m.memory[int64(uint32(v3))+248] = byte(i32(2))
														t190 := int64(load64(m.memory[int64(uint32(v3))+552:]))
														store64(m.memory[int64(uint32(v3))+256:], uint64(t190))
														t191 := int64(load64(m.memory[int64(uint32(v3))+208:]))
														store64(m.memory[int64(uint32(v3))+208:], uint64(t191+v24))
														t192 := int32(load32(m.memory[int64(uint32(v3))+196:]))
														t193 := int32(m.memory[int64(uint32(v3))+207])
														t194 := int32(load32(m.memory[int64(uint32(v3))+544:]))
														t195 := int32(load32(m.memory[int64(uint32(v3))+548:]))
														m.fn149(v11, t192, t193, t194, t195)
													}
												l31:
													store64(m.memory[int64(uint32(v3))+496:], uint64(i64(0x300000000)))
													v2 = i32(3)
													goto l63
												l23:
													m.memory[int64(uint32(v3))+248] = byte(v2)
													goto l74
												l50:
													store64(m.memory[int64(uint32(v3))+208:], uint64(v24+int64(uint32(v2))))
													t196 := fn1067(v1)
													v25 = int64(uint32(t196)) & i64(255)
												}
											l39:
												store64(m.memory[int64(uint32(v3))+216:], uint64(v24))
												store64(m.memory[int64(uint32(v3))+504:], uint64(v25))
											l38:
												store64(m.memory[int64(uint32(v3))+496:], uint64(i64(-0x7ffffff6ffffffff)))
												v2 = i32(-0x7ffffff7)
											l54:
												t197 := int32(load32(m.memory[int64(uint32(v3))+496:]))
												if t197 == 0 {
													if v2 == i32(10) {
														goto l59
													}
													v4 = i32(0)
													switch v2 {
													case 0:
														t198 := int32(load32(m.memory[int64(uint32(v3))+520:]))
														v4 = t198
														t199 := int32(load32(m.memory[int64(uint32(v3))+508:]))
														v5 = t199
														t200 := int32(load32(m.memory[int64(uint32(v3))+504:]))
														v1 = t200
														m.fn1068(v3+i32(544), v16, v11)
														t201 := int32(load32(m.memory[int64(uint32(v3))+544:]))
														v6 = t201
														if v6 == i32(-1) {
															t210 := int64(load64(m.memory[uint32(v8):]))
															store64(m.memory[int64(uint32(v3))+456:], uint64(t210))
															t211 := int32(load32(m.memory[int64(uint32(v8))+8:]))
															store32(m.memory[int64(uint32(v3))+464:], uint32(t211))
															v26 = v4
															goto l83
														}
														t202 := int32(load32(m.memory[int64(uint32(v10))+8:]))
														store32(m.memory[int64(uint32(v3))+464:], uint32(t202))
														t203 := int64(load64(m.memory[uint32(v10):]))
														store64(m.memory[int64(uint32(v3))+456:], uint64(t203))
														m.fn134(v1, v5)
														v1 = v6
														goto l80
													case 1:
														goto l67
													case 2:
														goto l78
													default:
														goto l76
													}
												}
												if uint32(v2) < uint32(i32(-0x7ffffff8)) {
													goto l61
												}
												m.memory[int64(uint32(v3))+248] = byte(i32(5))
											}
										l61:
											v4 = i32(1)
											goto l76
										l59:
											m.memory[int64(uint32(v3))+248] = byte(i32(5))
											v2 = i32(10)
										l63:
											v4 = i32(0)
											goto l76
										l78:
											t204 := int32(load32(m.memory[int64(uint32(v3))+520:]))
											v4 = t204
											t205 := int32(load32(m.memory[int64(uint32(v3))+508:]))
											v5 = t205
											t206 := int32(load32(m.memory[int64(uint32(v3))+504:]))
											v1 = t206
											m.fn1068(v3+i32(544), v16, v11)
											t207 := int32(load32(m.memory[int64(uint32(v3))+544:]))
											v6 = t207
											if v6 == i32(-1) {
												m.memory[int64(uint32(v3))+296] = byte(i32(1))
												t212 := int64(load64(m.memory[uint32(v8):]))
												store64(m.memory[int64(uint32(v3))+456:], uint64(t212))
												t213 := int32(load32(m.memory[int64(uint32(v8))+8:]))
												store32(m.memory[int64(uint32(v3))+464:], uint32(t213))
												v26 = v4
												goto l83
											}
											t208 := int32(load32(m.memory[int64(uint32(v10))+8:]))
											store32(m.memory[int64(uint32(v3))+464:], uint32(t208))
											t209 := int64(load64(m.memory[uint32(v10):]))
											store64(m.memory[int64(uint32(v3))+456:], uint64(t209))
											m.fn134(v1, v5)
											v1 = v6
										}
									l80:
										v2 = i32(-0x7ffffff2)
										goto l82
									l76:
										t214 := int64(load64(m.memory[uint32(v8):]))
										store64(m.memory[int64(uint32(v3))+456:], uint64(t214))
										t215 := int32(load32(m.memory[int64(uint32(v8))+8:]))
										store32(m.memory[int64(uint32(v3))+464:], uint32(t215))
										t216 := int32(load32(m.memory[int64(uint32(v3))+504:]))
										v1 = t216
										t217 := int32(load32(m.memory[int64(uint32(v3))+520:]))
										v26 = t217
										if v4 == 0 {
											goto l83
										}
									}
								l82:
									store32(m.memory[int64(uint32(v3))+500:], uint32(v1))
									store32(m.memory[int64(uint32(v3))+496:], uint32(v2))
									t218 := int64(load64(m.memory[int64(uint32(v3))+456:]))
									store64(m.memory[int64(uint32(v3))+504:], uint64(t218))
									t219 := int32(load32(m.memory[int64(uint32(v3))+464:]))
									store32(m.memory[int64(uint32(v3))+512:], uint32(t219))
									store32(m.memory[int64(uint32(v3))+516:], uint32(v26))
									store32(m.memory[int64(uint32(v3))+572:], uint32(i32(150)))
									store32(m.memory[int64(uint32(v3))+568:], uint32(v3+i32(496)))
									m.fn73(v3+i32(544), i32(1051951), v3+i32(568))
									store32(m.memory[int64(uint32(v3))+556:], uint32(i32(-1)))
									m.fn535(v3 + i32(496))
									t220 := int32(load32(m.memory[int64(uint32(v3))+560:]))
									t221 := v3
									v2 = t220
									store32(m.memory[int64(uint32(t221))+448:], uint32(v2))
									store32(m.memory[int64(uint32(v3))+432:], uint32(v2))
									t222 := int64(load64(m.memory[int64(uint32(v3))+552:]))
									t223 := v3
									v24 = t222
									store64(m.memory[int64(uint32(t223))+424:], uint64(v24))
									t224 := int32(load32(m.memory[int64(uint32(v3))+564:]))
									v1 = t224
									t225 := int64(load64(m.memory[int64(uint32(v3))+544:]))
									store64(m.memory[int64(uint32(v0))+4:], uint64(t225))
									store64(m.memory[int64(uint32(v0))+12:], uint64(v24))
									store32(m.memory[int64(uint32(v0))+20:], uint32(v2))
									store32(m.memory[int64(uint32(v0))+24:], uint32(v1))
									store32(m.memory[uint32(v0):], uint32(i32(-1)))
									goto l84
								}
							l67:
								t226 := int64(load64(m.memory[int64(uint32(v3))+508:]))
								store64(m.memory[int64(uint32(v3))+456:], uint64(t226))
								m.memory[int64(uint32(v3))+296] = byte(i32(1))
								t227 := int32(load32(m.memory[int64(uint32(v3))+504:]))
								v1 = t227
							}
						l83:
							t228 := int64(load64(m.memory[int64(uint32(v3))+456:]))
							t229 := v3
							v24 = t228
							store64(m.memory[int64(uint32(t229))+424:], uint64(v24))
							t230 := int32(load32(m.memory[int64(uint32(v3))+464:]))
							t231 := v3
							v4 = t230
							store32(m.memory[int64(uint32(t231))+432:], uint32(v4))
							store64(m.memory[uint32(v9):], uint64(v24))
							store32(m.memory[int64(uint32(v9))+8:], uint32(v4))
							store32(m.memory[int64(uint32(v3))+404:], uint32(v1))
							store32(m.memory[int64(uint32(v3))+400:], uint32(v2))
							store32(m.memory[int64(uint32(v3))+420:], uint32(v26))
							{
								{
									switch v2 {
									case 2:
										t314 := int32(load32(m.memory[int64(uint32(v3))+408:]))
										v4 = t314
										m.fn1073(v3+i32(496), v3+i32(396))
										{
											t315 := int32(load32(m.memory[int64(uint32(v3))+496:]))
											v5 = t315
											if v5 == i32(-1) {
												m.fn1074(v3+i32(496), v13, v3+i32(192), v3+i32(304))
												m.fn1076(v17, v19, v3+i32(340), v3+i32(496))
												goto l155
											}
											t316 := int32(load32(m.memory[int64(uint32(v3))+516:]))
											store32(m.memory[int64(uint32(v0))+24:], uint32(t316))
											t317 := int64(load64(m.memory[int64(uint32(v3))+508:]))
											store64(m.memory[int64(uint32(v0))+16:], uint64(t317))
											t318 := int64(load64(m.memory[int64(uint32(v3))+500:]))
											store64(m.memory[int64(uint32(v0))+8:], uint64(t318))
											store32(m.memory[uint32(v0):], uint32(i32(-1)))
											store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
											goto l153
										}
									case 3:
										t319 := int32(load32(m.memory[int64(uint32(v3))+412:]))
										v5 = t319
										t320 := int32(load32(m.memory[int64(uint32(v3))+408:]))
										v4 = t320
										t321 := int32(load32(m.memory[int64(uint32(v3))+416:]))
										m.fn1077(v3+i32(468), t321, v13)
										t322 := int32(load32(m.memory[int64(uint32(v3))+468:]))
										if t322 == i32(-2) {
											goto l156
										}
										m.fn490(v3+i32(544), v3+i32(468))
										goto l157
									case 4:
										t323 := int32(load32(m.memory[int64(uint32(v3))+408:]))
										t324 := v3 + i32(496)
										v4 = t323
										t325 := int32(load32(m.memory[int64(uint32(v3))+412:]))
										m.fn92(t324, v4, t325)
										m.fn490(v3+i32(544), v3+i32(496))
										m.fn1073(v3+i32(496), v3+i32(396))
										{
											t326 := int32(load32(m.memory[int64(uint32(v3))+496:]))
											v5 = t326
											if v5 == i32(-1) {
												t332 := int32(load32(m.memory[int64(uint32(v3))+552:]))
												store32(m.memory[int64(uint32(v3))+504:], uint32(t332))
												t333 := int64(load64(m.memory[int64(uint32(v3))+544:]))
												store64(m.memory[int64(uint32(v3))+496:], uint64(t333))
												m.fn1078(v17, v20, v3+i32(340), v3+i32(496))
												v21 = v20
												goto l94
											}
											t327 := int32(load32(m.memory[int64(uint32(v3))+516:]))
											store32(m.memory[int64(uint32(v0))+24:], uint32(t327))
											t328 := int64(load64(m.memory[int64(uint32(v3))+508:]))
											store64(m.memory[int64(uint32(v0))+16:], uint64(t328))
											t329 := int64(load64(m.memory[int64(uint32(v3))+500:]))
											store64(m.memory[int64(uint32(v0))+8:], uint64(t329))
											store32(m.memory[uint32(v0):], uint32(i32(-1)))
											store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
											t330 := int32(load32(m.memory[int64(uint32(v3))+544:]))
											t331 := int32(load32(m.memory[int64(uint32(v3))+548:]))
											m.fn16(t330, t331)
											goto l153
										}
									case 10:
										m.fn200(v3 + i32(400))
										if v21 == 0 {
											goto l159
										}
										v2 = v21 + i32(-1)
										v1 = v17 + v21*i32(44) + i32(-44)
										v5 = v3 + i32(496) + i32(4)
									l162:
										if v2 != i32(-1) {
											t334 := int32(load32(m.memory[uint32(v1):]))
											v4 = t334
											if v4 == i32(-1) {
												goto l161
											}
											memory_copy(m.memory, uint32(v5), uint32(v1+i32(4)), uint32(i32(40)))
											store32(m.memory[int64(uint32(v3))+496:], uint32(v4))
											m.fn1076(v17, v2, v3+i32(340), v3+i32(496))
											v2 = v2 + i32(-1)
											v1 = v1 + i32(-44)
											goto l162
										}
										v2 = i32(0)
										goto l161
									l161:
										store32(m.memory[int64(uint32(v3))+392:], uint32(v2))
									l159:
										memory_copy(m.memory, uint32(v0), uint32(v3+i32(340)), uint32(i32(44)))
										m.fn1079(v3 + i32(384))
										goto l163
									default:
										goto l90
									case 1:
										t232 := int32(load32(m.memory[int64(uint32(v3))+408:]))
										v4 = t232
										if v19 != 0 {
											t233 := int32(load32(m.memory[int64(uint32(v3))+412:]))
											v6 = t233
											t234 := v3
											v18 = v19 + i32(-1)
											store32(m.memory[int64(uint32(t234))+392:], uint32(v18))
											v5 = v17 + v18*i32(44)
											t235 := int32(load32(m.memory[uint32(v5):]))
											v7 = t235
											if v7 != i32(-1) {
												t309 := int64(load64(m.memory[int64(uint32(v5))+4:]))
												v24 = t309
												t310 := int64(load64(m.memory[int64(uint32(v5))+36:]))
												store64(m.memory[int64(uint32(v8))+24:], uint64(t310))
												t311 := int64(load64(m.memory[int64(uint32(v5))+28:]))
												store64(m.memory[int64(uint32(v8))+16:], uint64(t311))
												t312 := int64(load64(m.memory[int64(uint32(v5))+20:]))
												store64(m.memory[int64(uint32(v8))+8:], uint64(t312))
												t313 := int64(load64(m.memory[int64(uint32(v5))+12:]))
												store64(m.memory[uint32(v8):], uint64(t313))
												m.fn553(v3+i32(8), v4, v6)
												store64(m.memory[int64(uint32(v3))+500:], uint64(v24))
												store32(m.memory[int64(uint32(v3))+496:], uint32(v7))
												m.fn1076(v17, v18, v3+i32(340), v3+i32(496))
												goto l96
											}
											goto l96
										}
										v19 = i32(0)
										v20 = i32(0)
										v21 = i32(0)
										goto l94
									case 9:
										t236 := int32(load32(m.memory[int64(uint32(v3))+408:]))
										t237 := v3 + i32(496)
										v4 = t236
										t238 := int32(load32(m.memory[int64(uint32(v3))+412:]))
										m.fn92(t237, v4, t238)
										m.fn490(v3+i32(544), v3+i32(496))
										t239 := int32(load32(m.memory[int64(uint32(v3))+548:]))
										t240 := v3 + i32(32)
										v6 = t239
										t241 := int32(load32(m.memory[int64(uint32(v3))+552:]))
										t242 := v6
										v7 = t241
										m.fn13(t240, t242, v7, i32(35))
										{
											{
												{
													{
														t243 := int32(load32(m.memory[int64(uint32(v3))+32:]))
														v5 = t243
														if v5 == 0 {
															t252 := m.fn15(v6, v7, i32(1281596), i32(3))
															if t252 == 0 {
																goto l103
															}
															v5 = i32(38)
															goto l104
														}
														t244 := int32(load32(m.memory[int64(uint32(v3))+36:]))
														t245 := v3 + i32(24)
														t246 := v5
														v7 = t244
														m.fn13(t245, t246, v7, i32(120))
														{
															{
																t247 := int32(load32(m.memory[int64(uint32(v3))+24:]))
																v6 = t247
																if v6 == 0 {
																	goto l98
																}
																t248 := int32(load32(m.memory[int64(uint32(v3))+28:]))
																v18 = t248
																goto l99
															}
														l98:
															m.fn13(v3+i32(16), v5, v7, i32(88))
															t249 := int32(load32(m.memory[int64(uint32(v3))+20:]))
															v18 = t249
															t250 := int32(load32(m.memory[int64(uint32(v3))+16:]))
															v6 = t250
														}
													l99:
														if v6 == 0 {
															goto l100
														}
														m.fn1070(v3+i32(496), v6, v18, i32(16))
														t251 := int32(m.memory[int64(uint32(v3))+496])
														if t251 == 0 {
															goto l101
														}
														goto l102
													}
												l100:
													m.fn1071(v3+i32(496), v5, v7)
													t253 := int32(m.memory[int64(uint32(v3))+496])
													if t253 != 0 {
														goto l102
													}
												}
											l101:
												t254 := int32(load32(m.memory[int64(uint32(v3))+500:]))
												v5 = t254
												if uint32(v5^i32(55296)+i32(-1114112)) < uint32(i32(-1112064)) {
													goto l102
												}
												goto l104
											}
										l103:
											{
												t255 := m.fn15(v6, v7, i32(1084265), i32(2))
												if t255 == 0 {
													goto l105
												}
												v5 = i32(60)
												goto l104
											}
										l105:
											{
												t256 := m.fn15(v6, v7, i32(1084267), i32(2))
												if t256 == 0 {
													goto l106
												}
												v5 = i32(62)
												goto l104
											}
										l106:
											{
												t257 := m.fn15(v6, v7, i32(1084269), i32(4))
												if t257 == 0 {
													goto l107
												}
												v5 = i32(39)
												goto l104
											}
										l107:
											{
												t258 := m.fn15(v6, v7, i32(1084273), i32(4))
												if t258 == 0 {
													goto l108
												}
												v5 = i32(34)
												goto l104
											}
										l108:
											{
												t259 := m.fn15(v6, v7, i32(1084277), i32(4))
												if t259 == 0 {
													goto l109
												}
												v5 = i32(160)
												goto l104
											}
										l109:
											{
												t260 := m.fn15(v6, v7, i32(1084281), i32(3))
												if t260 == 0 {
													goto l110
												}
												v5 = i32(173)
												goto l104
											}
										l110:
											{
												t261 := m.fn15(v6, v7, i32(1084284), i32(5))
												if t261 == 0 {
													goto l111
												}
												v5 = i32(8212)
												goto l104
											}
										l111:
											{
												t262 := m.fn15(v6, v7, i32(1084289), i32(5))
												if t262 == 0 {
													goto l112
												}
												v5 = i32(8211)
												goto l104
											}
										l112:
											{
												t263 := m.fn15(v6, v7, i32(1084294), i32(5))
												if t263 == 0 {
													goto l113
												}
												v5 = i32(8216)
												goto l104
											}
										l113:
											{
												t264 := m.fn15(v6, v7, i32(1084299), i32(5))
												if t264 == 0 {
													goto l114
												}
												v5 = i32(8217)
												goto l104
											}
										l114:
											{
												t265 := m.fn15(v6, v7, i32(1084304), i32(5))
												if t265 == 0 {
													goto l115
												}
												v5 = i32(8220)
												goto l104
											}
										l115:
											{
												t266 := m.fn15(v6, v7, i32(1084309), i32(5))
												if t266 == 0 {
													goto l116
												}
												v5 = i32(8221)
												goto l104
											}
										l116:
											{
												t267 := m.fn15(v6, v7, i32(1084314), i32(6))
												if t267 == 0 {
													goto l117
												}
												v5 = i32(8230)
												goto l104
											}
										l117:
											{
												t268 := m.fn15(v6, v7, i32(1084320), i32(4))
												if t268 == 0 {
													goto l118
												}
												v5 = i32(169)
												goto l104
											}
										l118:
											{
												t269 := m.fn15(v6, v7, i32(1084324), i32(3))
												if t269 == 0 {
													goto l119
												}
												v5 = i32(174)
												goto l104
											}
										l119:
											{
												t270 := m.fn15(v6, v7, i32(1084327), i32(5))
												if t270 == 0 {
													goto l120
												}
												v5 = i32(8482)
												goto l104
											}
										l120:
											{
												t271 := m.fn15(v6, v7, i32(1084332), i32(3))
												if t271 == 0 {
													goto l121
												}
												v5 = i32(176)
												goto l104
											}
										l121:
											{
												t272 := m.fn15(v6, v7, i32(1084335), i32(6))
												if t272 == 0 {
													goto l122
												}
												v5 = i32(183)
												goto l104
											}
										l122:
											{
												t273 := m.fn15(v6, v7, i32(1084341), i32(4))
												if t273 == 0 {
													goto l123
												}
												v5 = i32(8226)
												goto l104
											}
										l123:
											{
												t274 := m.fn15(v6, v7, i32(1079225), i32(4))
												if t274 == 0 {
													goto l124
												}
												v5 = i32(167)
												goto l104
											}
										l124:
											{
												t275 := m.fn15(v6, v7, i32(1084345), i32(4))
												if t275 == 0 {
													goto l125
												}
												v5 = i32(182)
												goto l104
											}
										l125:
											{
												t276 := m.fn15(v6, v7, i32(1084349), i32(5))
												if t276 == 0 {
													goto l126
												}
												v5 = i32(171)
												goto l104
											}
										l126:
											{
												t277 := m.fn15(v6, v7, i32(1084354), i32(5))
												if t277 == 0 {
													goto l127
												}
												v5 = i32(187)
												goto l104
											}
										l127:
											{
												t278 := m.fn15(v6, v7, i32(1084359), i32(5))
												if t278 == 0 {
													goto l128
												}
												v5 = i32(215)
												goto l104
											}
										l128:
											{
												t279 := m.fn15(v6, v7, i32(1084364), i32(6))
												if t279 == 0 {
													goto l129
												}
												v5 = i32(247)
												goto l104
											}
										l129:
											{
												t280 := m.fn15(v6, v7, i32(1084370), i32(6))
												if t280 == 0 {
													goto l130
												}
												v5 = i32(177)
												goto l104
											}
										l130:
											{
												t281 := m.fn15(v6, v7, i32(1084376), i32(6))
												if t281 == 0 {
													goto l131
												}
												v5 = i32(189)
												goto l104
											}
										l131:
											{
												t282 := m.fn15(v6, v7, i32(1084382), i32(6))
												if t282 == 0 {
													goto l132
												}
												v5 = i32(188)
												goto l104
											}
										l132:
											{
												t283 := m.fn15(v6, v7, i32(1084388), i32(6))
												if t283 == 0 {
													goto l133
												}
												v5 = i32(233)
												goto l104
											}
										l133:
											{
												t284 := m.fn15(v6, v7, i32(1084394), i32(6))
												if t284 == 0 {
													goto l134
												}
												v5 = i32(232)
												goto l104
											}
										l134:
											{
												t285 := m.fn15(v6, v7, i32(1084400), i32(6))
												if t285 == 0 {
													goto l135
												}
												v5 = i32(224)
												goto l104
											}
										l135:
											{
												t286 := m.fn15(v6, v7, i32(1084406), i32(6))
												if t286 == 0 {
													goto l136
												}
												v5 = i32(231)
												goto l104
											}
										l136:
											{
												t287 := m.fn15(v6, v7, i32(1084412), i32(4))
												if t287 == 0 {
													goto l137
												}
												v5 = i32(252)
												goto l104
											}
										l137:
											{
												t288 := m.fn15(v6, v7, i32(1084416), i32(4))
												if t288 == 0 {
													goto l138
												}
												v5 = i32(246)
												goto l104
											}
										l138:
											{
												t289 := m.fn15(v6, v7, i32(1084420), i32(4))
												if t289 == 0 {
													goto l139
												}
												v5 = i32(228)
												goto l104
											}
										l139:
											{
												t290 := m.fn15(v6, v7, i32(1084424), i32(5))
												if t290 == 0 {
													goto l140
												}
												v5 = i32(223)
												goto l104
											}
										l140:
											{
												t291 := m.fn15(v6, v7, i32(1084429), i32(5))
												if t291 == 0 {
													goto l141
												}
												v5 = i32(229)
												goto l104
											}
										l141:
											{
												t292 := m.fn15(v6, v7, i32(1084434), i32(6))
												if t292 == 0 {
													goto l142
												}
												v5 = i32(248)
												goto l104
											}
										l142:
											{
												t293 := m.fn15(v6, v7, i32(1084440), i32(5))
												if t293 == 0 {
													goto l143
												}
												v5 = i32(230)
												goto l104
											}
										l143:
											{
												t294 := m.fn15(v6, v7, i32(1084445), i32(4))
												if t294 == 0 {
													goto l144
												}
												v5 = i32(8364)
												goto l104
											}
										l144:
											{
												t295 := m.fn15(v6, v7, i32(1084449), i32(5))
												if t295 == 0 {
													goto l145
												}
												v5 = i32(163)
												goto l104
											}
										l145:
											{
												t296 := m.fn15(v6, v7, i32(1084454), i32(3))
												if t296 == 0 {
													goto l146
												}
												v5 = i32(165)
												goto l104
											}
										l146:
											t297 := m.fn15(v6, v7, i32(1084457), i32(4))
											if t297 == 0 {
												goto l102
											}
											v5 = i32(162)
										}
									l104:
										m.fn1072(v3+i32(496), v5)
										t298 := int32(load32(m.memory[int64(uint32(v3))+496:]))
										if t298 == i32(-1) {
											goto l147
										}
										t299 := int32(load32(m.memory[int64(uint32(v3))+504:]))
										store32(m.memory[int64(uint32(v3))+488:], uint32(t299))
										t300 := int64(load64(m.memory[int64(uint32(v3))+496:]))
										store64(m.memory[int64(uint32(v3))+480:], uint64(t300))
										goto l148
									case 0:
										t301 := int32(load32(m.memory[int64(uint32(v3))+408:]))
										v4 = t301
										if uint32(v18) > uint32(i32(255)) {
											goto l149
										}
										m.fn1073(v3+i32(496), v3+i32(396))
										{
											t302 := int32(load32(m.memory[int64(uint32(v3))+496:]))
											v5 = t302
											if v5 == i32(-1) {
												m.fn1074(v3+i32(496), v13, v3+i32(192), v3+i32(304))
												{
													t306 := int32(load32(m.memory[int64(uint32(v3))+384:]))
													if v18 != t306 {
														goto l152
													}
													m.fn1075(v3 + i32(384))
													t307 := int32(load32(m.memory[int64(uint32(v3))+388:]))
													v17 = t307
												}
											l152:
												memory_copy(m.memory, uint32(v17+v18*i32(44)), uint32(v3+i32(496)), uint32(i32(44)))
												t308 := v3
												v18 = v18 + i32(1)
												store32(m.memory[int64(uint32(t308))+392:], uint32(v18))
												goto l96
											}
											t303 := int32(load32(m.memory[int64(uint32(v3))+516:]))
											store32(m.memory[int64(uint32(v0))+24:], uint32(t303))
											t304 := int64(load64(m.memory[int64(uint32(v3))+508:]))
											store64(m.memory[int64(uint32(v0))+16:], uint64(t304))
											t305 := int64(load64(m.memory[int64(uint32(v3))+500:]))
											store64(m.memory[int64(uint32(v0))+8:], uint64(t305))
											v2 = i32(4)
											goto l151
										}
									l149:
										store32(m.memory[int64(uint32(v3))+500:], uint32(i32(5)))
										store32(m.memory[int64(uint32(v3))+496:], uint32(i32(1084540)))
										m.fn73(v0+i32(8), i32(1050084), v3+i32(496))
										store32(m.memory[int64(uint32(v0))+20:], uint32(i32(1084544)))
										store32(m.memory[int64(uint32(v0))+4:], uint32(i32(-0x7ffffffd)))
										v5 = i32(13)
										v2 = i32(24)
									l151:
										store32(m.memory[uint32(v0+v2):], uint32(v5))
										store32(m.memory[uint32(v0):], uint32(i32(-1)))
										goto l153
									}
								l156:
									m.fn92(v3+i32(496), v4, v5)
									m.fn490(v3+i32(544), v3+i32(496))
								l157:
									t335 := int32(load32(m.memory[int64(uint32(v3))+552:]))
									if t335 == 0 {
										goto l164
									}
									m.fn1073(v3+i32(496), v3+i32(396))
									{
										t336 := int32(load32(m.memory[int64(uint32(v3))+496:]))
										v5 = t336
										if v5 == i32(-1) {
											goto l165
										}
										t337 := int32(load32(m.memory[int64(uint32(v3))+516:]))
										store32(m.memory[int64(uint32(v0))+24:], uint32(t337))
										t338 := int64(load64(m.memory[int64(uint32(v3))+508:]))
										store64(m.memory[int64(uint32(v0))+16:], uint64(t338))
										t339 := int64(load64(m.memory[int64(uint32(v3))+500:]))
										store64(m.memory[int64(uint32(v0))+8:], uint64(t339))
										store32(m.memory[uint32(v0):], uint32(i32(-1)))
										store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
										t340 := int32(load32(m.memory[int64(uint32(v3))+544:]))
										t341 := int32(load32(m.memory[int64(uint32(v3))+548:]))
										m.fn16(t340, t341)
										goto l153
									}
								l165:
									t342 := int32(load32(m.memory[int64(uint32(v3))+552:]))
									store32(m.memory[int64(uint32(v3))+504:], uint32(t342))
									t343 := int64(load64(m.memory[int64(uint32(v3))+544:]))
									store64(m.memory[int64(uint32(v3))+496:], uint64(t343))
									m.fn1078(v17, v19, v3+i32(340), v3+i32(496))
								}
							l155:
								v20 = v19
								v21 = v19
								goto l94
							l102:
								store32(m.memory[int64(uint32(v3))+496:], uint32(i32(-1)))
							l147:
								store32(m.memory[int64(uint32(v3))+460:], uint32(i32(25)))
								store32(m.memory[int64(uint32(v3))+456:], uint32(v3+i32(544)))
								m.fn73(v3+i32(480), i32(1068659), v3+i32(456))
							l148:
								m.fn1073(v3+i32(496), v3+i32(396))
								t344 := int32(load32(m.memory[int64(uint32(v3))+496:]))
								v5 = t344
								if v5 == i32(-1) {
									t352 := int32(load32(m.memory[int64(uint32(v3))+388:]))
									v17 = t352
									t353 := int32(load32(m.memory[int64(uint32(v3))+392:]))
									t354 := v17
									v18 = t353
									m.fn1078(t354, v18, v3+i32(340), v3+i32(480))
									t355 := int32(load32(m.memory[int64(uint32(v3))+544:]))
									t356 := int32(load32(m.memory[int64(uint32(v3))+548:]))
									m.fn16(t355, t356)
									goto l96
								}
								t345 := int32(load32(m.memory[int64(uint32(v3))+516:]))
								store32(m.memory[int64(uint32(v0))+24:], uint32(t345))
								t346 := int64(load64(m.memory[int64(uint32(v3))+508:]))
								store64(m.memory[int64(uint32(v0))+16:], uint64(t346))
								t347 := int64(load64(m.memory[int64(uint32(v3))+500:]))
								store64(m.memory[int64(uint32(v0))+8:], uint64(t347))
								store32(m.memory[uint32(v0):], uint32(i32(-1)))
								store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
								t348 := int32(load32(m.memory[int64(uint32(v3))+480:]))
								t349 := int32(load32(m.memory[int64(uint32(v3))+484:]))
								m.fn16(t348, t349)
								t350 := int32(load32(m.memory[int64(uint32(v3))+544:]))
								t351 := int32(load32(m.memory[int64(uint32(v3))+548:]))
								m.fn16(t350, t351)
							}
						l153:
							m.fn134(v1, v4)
						}
					l84:
						m.fn1079(v3 + i32(384))
						m.fn1042(v3 + i32(340))
						goto l163
					l163:
						{
							t357 := int32(load32(m.memory[int64(uint32(v3))+308:]))
							v7 = t357
							if v7 == 0 {
								goto l167
							}
							t358 := int32(load32(m.memory[int64(uint32(v3))+304:]))
							v6 = t358
							{
								t359 := int32(load32(m.memory[int64(uint32(v3))+316:]))
								v4 = t359
								if v4 == 0 {
									goto l168
								}
								v2 = v6 + i32(8)
								t360 := int64(load64(m.memory[uint32(v6):]))
								v24 = (t360 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
								v1 = v6
							l171:
								if v4 == 0 {
									goto l168
								}
							l170:
								{
									if v24 != i64(0) {
										v5 = v1 + (i32(0)-int32(uint32(int64(bits.TrailingZeros64(uint64(v24))))>>3))*i32(20)
										t362 := int32(load32(m.memory[uint32(v5+i32(-20)):]))
										t363 := int32(load32(m.memory[uint32(v5+i32(-16)):]))
										m.fn16(t362, t363)
										m.fn754(v5 + i32(-8))
										v4 = v4 + i32(-1)
										v24 = (v24 + i64(-1)) & v24
										goto l171
									}
									v1 = v1 + i32(-160)
									t361 := int64(load64(m.memory[uint32(v2):]))
									v24 = (t361 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
									v2 = v2 + i32(8)
									goto l170
								}
							}
						l168:
							m.fn39(v3+i32(496), i32(20), i32(8), v7+i32(1))
							t364 := int32(load32(m.memory[int64(uint32(v3))+504:]))
							t365 := int32(load32(m.memory[int64(uint32(v3))+496:]))
							t366 := int32(load32(m.memory[int64(uint32(v3))+500:]))
							m.fn40(v6-t364, t365, t366)
						}
					l167:
						m.fn229(v3 + i32(192))
						t367 := int32(load32(m.memory[int64(uint32(v3))+264:]))
						t368 := int32(load32(m.memory[int64(uint32(v3))+268:]))
						m.fn16(t367, t368)
						t369 := int32(load32(m.memory[int64(uint32(v3))+276:]))
						t370 := int32(load32(m.memory[int64(uint32(v3))+280:]))
						m.fn1080(t369, t370, i32(4), i32(16))
						t371 := int32(load32(m.memory[int64(uint32(v3))+180:]))
						t372 := int32(load32(m.memory[int64(uint32(v3))+184:]))
						m.fn134(t371, t372)
						m.g0 = v3 + i32(576)
						return
					}
				l96:
					v19 = v18
					v20 = v18
					v21 = v18
					goto l94
				l164:
					t373 := int32(load32(m.memory[int64(uint32(v3))+544:]))
					t374 := int32(load32(m.memory[int64(uint32(v3))+548:]))
					m.fn16(t373, t374)
				}
			l94:
				m.fn134(v1, v4)
			l90:
				if uint32(v2) > uint32(i32(9)) {
					goto l172
				}
				if i32_shl(i32(1), v2)&i32(543) != 0 {
					goto l173
				}
			l172:
				m.fn200(v3 + i32(400))
				goto l173
			}
		}
		store16(m.memory[int64(uint32(v3))+204:], uint16(i32(0)))
		t47 := int32(load32(m.memory[int64(uint32(v3))+348:]))
		store32(m.memory[int64(uint32(v3))+192:], uint32(t47))
		t48 := int32(load32(m.memory[uint32(v2+i32(1284272)):]))
		t49 := v3
		v1 = t48
		store32(m.memory[int64(uint32(t49))+200:], uint32(v1))
		t50 := int32(load32(m.memory[uint32(v2+i32(1284264)):]))
		t51 := v3
		v4 = t50
		store32(m.memory[int64(uint32(t51))+196:], uint32(v4))
		t52 := int32(load32(m.memory[uint32(v2+i32(1284260)):]))
		v5 = t52
		m.fn1058(v3+i32(496), v3+i32(192))
		m.fn1059(v3+i32(340), v5, v4)
		t53 := int32(load32(m.memory[uint32(v2+i32(1284268)):]))
		m.fn1059(v3+i32(340), t53, v1)
		v2 = v2 + i32(16)
		goto l10
	}
}
func (m *Module) fn1054(v0 int32) {
	t0 := int32(load32(m.memory[uint32(v0):]))
	if t0 == i32(-1) {
		return
	}
	m.fn1042(v0)
}
func (m *Module) fn1055(v0, v1, v2, v3, v4 int32) {
	var v5 int32
	t0 := m.g0
	v5 = t0 - i32(80)
	m.g0 = v5
	m.fn601(v5+i32(16), v1, v2, v3, v4)
	m.fn790(v5+i32(4), v5+i32(16))
	t1 := int32(load32(m.memory[int64(uint32(v5))+4:]))
	v4 = t1
	t2 := int32(load32(m.memory[int64(uint32(v5))+8:]))
	store32(m.memory[int64(uint32(v0))+4:], uint32(t2))
	store32(m.memory[uint32(v0):], uint32(v4))
	m.g0 = v5 + i32(80)
}
func (m *Module) fn1056(v0, v1, v2, v3 int32) {
	var v4 int32
	t0 := m.g0
	v4 = t0 - i32(48)
	m.g0 = v4
	m.fn514(v4+i32(20), v3, v1, v2)
	m.fn516(v4+i32(8), v4+i32(20))
	t1 := int32(load32(m.memory[int64(uint32(v4))+8:]))
	v2 = t1
	t2 := int32(load32(m.memory[int64(uint32(v4))+12:]))
	store32(m.memory[int64(uint32(v0))+4:], uint32(t2))
	store32(m.memory[uint32(v0):], uint32(v2))
	m.g0 = v4 + i32(48)
}
func (m *Module) fn1057(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7 int32
	t0 := m.g0
	v2 = t0 - i32(32)
	m.g0 = v2
	v3 = i32(0)
	store16(m.memory[int64(uint32(v2))+30:], uint16(i32(0)))
	store64(m.memory[int64(uint32(v2))+22:], uint64(i64(0)))
	store64(m.memory[int64(uint32(v2))+14:], uint64(i64(0)))
	v4 = v0 + v1
l2:
	if v1 == 0 {
		goto l0
	}
	v5 = v0 + i32(1)
	{
		t1 := int32(m.memory[uint32(v0)])
		v6 = t1
		v7 = v6 + i32(-9)
		if uint32(v7) > uint32(i32(23)) {
			goto l1
		}
		if i32_shl(i32(1), v7)&i32(8388635) == 0 {
			goto l1
		}
		v1 = v1 + i32(-1)
		v0 = v5
		goto l2
	}
l1:
	v1 = v6 + i32(-45)
	if uint32(v1) > uint32(i32(13)) {
		goto l3
	}
	if i32_shl(i32(1), v1)&i32(8195) != 0 {
		goto l4
	}
l3:
	if v6 == i32(95) {
		goto l4
	}
	if uint32((v6+i32(-65))&i32(255)) < uint32(i32(26)) {
		goto l5
	}
	if uint32((v6+i32(-97))&i32(255)) < uint32(i32(26)) {
		goto l6
	}
	if uint32((v6+i32(-48))&i32(255)) < uint32(i32(10)) {
		goto l6
	}
	goto l0
l4:
	v5 = v0 + i32(1)
	goto l6
l5:
	v6 = v6 | i32(32)
l6:
	m.memory[int64(uint32(v2))+13] = byte(v6)
	v0 = i32(0)
l26:
	v1 = v5 + v0
	if v1 != v4 {
		goto l7
	}
	v3 = v0 + i32(1)
	v0 = v4
	goto l22
l7:
	v7 = v0 + i32(1)
	{
		t2 := int32(m.memory[uint32(v1)])
		v1 = t2
		v6 = v1 + i32(-45)
		if uint32(v6) > uint32(i32(13)) {
			goto l9
		}
		if i32_shl(i32(1), v6)&i32(8195) != 0 {
			goto l10
		}
	}
l9:
	if v1 == i32(95) {
		goto l10
	}
	v6 = v1 + i32(-9)
	if uint32(v6) > uint32(i32(23)) {
		goto l11
	}
	if i32_shl(i32(1), v6)&i32(8388635) != 0 {
		goto l12
	}
l11:
	if uint32((v1+i32(-65))&i32(255)) < uint32(i32(26)) {
		if v0 == i32(18) {
			goto l15
		}
		if uint32(v7) > uint32(i32(18)) {
			m.fn158(v7, i32(19), i32(1155008))
			panic("unreachable")
		}
		m.memory[uint32(v2+i32(13)+v0+i32(1))] = byte(v1 | i32(32))
		goto l24
	}
	if uint32((v1+i32(-97))&i32(255)) < uint32(i32(26)) {
		goto l10
	}
	v3 = i32(0)
	if uint32((v1+i32(-58))&i32(255)) < uint32(i32(246)) {
		goto l0
	}
	if v0 == i32(18) {
		goto l0
	}
	goto l14
l10:
	if v0 != i32(18) {
		goto l14
	}
	goto l15
l12:
	v3 = v0 + i32(1)
	v0 = v3 + v5
l22:
	{
		if v0 == v4 {
			if uint32(v3) > uint32(i32(19)) {
				m.fn151(i32(0), v3, i32(19), i32(1154960))
				panic("unreachable")
			}
			v0 = i32(228)
			v1 = i32(0)
		l20:
			{
				if uint32(v0) < uint32(i32(2)) {
					t12 := v2 + i32(13)
					t13 := v3
					v0 = v1 << 3
					t14 := int32(load32(m.memory[int64(uint32(v0))+1151232:]))
					t15 := int32(load32(m.memory[uint32(v0+i32(1151236)):]))
					t16 := m.fn1717(t12, t13, t14, t15)
					v0 = t16 & i32(255)
					if v0 != 0 {
						goto l15
					}
					{
						t17 := v1
						var p18 int32
						if v0 == i32(255) {
							p18 = 1
						}
						v0 = t17 + p18
						if uint32(v0) > uint32(i32(227)) {
							m.fn158(i32(228), i32(228), i32(1154976))
							panic("unreachable")
						}
						t19 := int32(load32(m.memory[int64(uint32(v0<<2))+1153872:]))
						v3 = t19
						goto l0
					}
				}
				v7 = int32(uint32(v0) >> 1)
				v5 = v7 + v1
				t4 := v1
				t5 := v5
				t6 := v2 + i32(13)
				t7 := v3
				v6 = v5 << 3
				t8 := int32(load32(m.memory[int64(uint32(v6))+1151232:]))
				t9 := int32(load32(m.memory[uint32(v6+i32(1151236)):]))
				t10 := m.fn1717(t6, t7, t8, t9)
				p11 := t5
				if t10&i32(255) == i32(1) {
					p11 = t4
				}
				v1 = p11
				v0 = v0 - v7
				goto l20
			}
		}
		t3 := int32(m.memory[uint32(v0)])
		v1 = t3 + i32(-9)
		if uint32(v1) > uint32(i32(23)) {
			goto l15
		}
		if i32_shl(i32(1), v1)&i32(8388635) != 0 {
			v0 = v0 + i32(1)
			goto l22
		}
		goto l15
	}
l15:
	v3 = i32(0)
	goto l0
l14:
	if uint32(v7) > uint32(i32(18)) {
		m.fn158(v7, i32(19), i32(1154992))
		panic("unreachable")
	}
	m.memory[uint32(v2+i32(13)+v0+i32(1))] = byte(v1)
l24:
	v0 = v7
	goto l26
l0:
	m.g0 = v2 + i32(32)
	return v3
}
func (m *Module) fn1058(v0, v1 int32) {
	var v2 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		v2 = t0
		t1 := int32(load32(m.memory[uint32(v0):]))
		if v2 != t1 {
			goto l0
		}
		m.fn1778(v0)
	}
l0:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v2+i32(1)))
	t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	v0 = t2 + v2<<4
	t3 := int64(load64(m.memory[uint32(v1):]))
	store64(m.memory[uint32(v0):], uint64(t3))
	t4 := int64(load64(m.memory[int64(uint32(v1))+8:]))
	store64(m.memory[int64(uint32(v0))+8:], uint64(t4))
}
func (m *Module) fn1059(v0, v1, v2 int32) {
	m.fn1771(v0, v1, v1+v2)
}
func (m *Module) fn1060(v0, v1, v2 int32) {
	var v3, v4, v5 int32
	var v6 int64
	var v7 int32
	var v8 int64
	t0 := m.g0
	v3 = t0 - i32(32)
	m.g0 = v3
	m.memory[int64(uint32(v3))+15] = byte(i32(0))
	t1 := int32(load32(m.memory[uint32(v1):]))
	t2 := v3
	t3 := v3 + i32(15)
	v4 = t1
	t4 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	t5 := v4
	v5 = t4
	m.fn1762(t2, t3, t5, v5)
	{
		{
			t6 := int32(load32(m.memory[uint32(v3):]))
			if t6 != i32(1) {
				goto l0
			}
			t7 := int64(load64(m.memory[uint32(v2):]))
			v6 = t7
			t8 := int32(load32(m.memory[int64(uint32(v3))+4:]))
			t9 := v3 + i32(16)
			t10 := v4
			t11 := v5
			v7 = t8 + i32(1)
			m.fn309(t9, t10, t11, v7, i32(1281600))
			t12 := int64(load64(m.memory[int64(uint32(v3))+24:]))
			v8 = t12
			t13 := int64(load64(m.memory[int64(uint32(v3))+16:]))
			store64(m.memory[int64(uint32(v0))+4:], uint64(t13))
			store64(m.memory[uint32(v1):], uint64(v8))
			v6 = v6 + int64(uint32(v7))
			v1 = i32(-1)
			goto l1
		}
	l0:
		t14 := int32(m.memory[int64(uint32(v3))+15])
		m.memory[int64(uint32(v0))+4] = byte(t14 + i32(6))
		t15 := int64(load64(m.memory[uint32(v2):]))
		v6 = t15 + int64(uint32(v5))
		v1 = i32(-0x7ffffff7)
	}
l1:
	store32(m.memory[uint32(v0):], uint32(v1))
	store64(m.memory[uint32(v2):], uint64(v6))
	m.g0 = v3 + i32(32)
}
func (m *Module) fn1061(v0, v1, v2, v3 int32) int32 {
	var v4 int32
	v4 = i32(0)
	{
		if uint32(v1) < uint32(v3) {
			goto l0
		}
		t0 := m.fn1755(v2, v0+(v1-v3), v3)
		v4 = t0
	}
l0:
	return v4
}
func (m *Module) fn1062(v0, v1, v2, v3 int32) {
	var v4, v5, v6, v7, v8 int32
	t0 := m.g0
	v4 = t0 - i32(32)
	m.g0 = v4
	m.memory[int64(uint32(v4))+28] = byte(i32(62))
	t1 := v4
	v5 = v2 + v3
	store32(m.memory[int64(uint32(t1))+24:], uint32(v5))
	store32(m.memory[int64(uint32(v4))+20:], uint32(v2))
	store32(m.memory[int64(uint32(v4))+16:], uint32(v2))
	t2 := int32(m.memory[uint32(v1)])
	v6 = t2 & i32(1)
l6:
	{
		m.fn155(v4+i32(8), v4+i32(16))
		{
			t3 := int32(load32(m.memory[int64(uint32(v4))+8:]))
			if t3 != i32(1) {
				{
					v7 = v5 + i32(-1)
					if v7 != 0 {
						goto l3
					}
					goto l4
				l3:
					t5 := int32(m.memory[uint32(v7)])
					var p6 int32
					if t5 == i32(63) {
						p6 = 1
					}
					v2 = p6
				}
			l4:
				v8 = i32(0)
				t7 := v1
				var p8 int32
				if v7 != i32(0) {
					p8 = 1
				}
				m.memory[uint32(t7)] = byte(p8 & v2)
				goto l5
			}
			t4 := int32(load32(m.memory[int64(uint32(v4))+12:]))
			v7 = t4
			if v7 != 0 {
				goto l1
			}
			if v6 == 0 {
				goto l1
			}
			v7 = i32(0)
			goto l2
		}
	l1:
		if v7 == 0 {
			goto l6
		}
		v8 = v7 + i32(-1)
		if uint32(v8) >= uint32(v3) {
			m.fn158(v8, v3, i32(1283572))
			panic("unreachable")
		}
		t9 := int32(m.memory[uint32(v2+v8)])
		if t9 != i32(63) {
			goto l6
		}
	}
l2:
	v8 = i32(1)
	goto l5
l5:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v7))
	store32(m.memory[uint32(v0):], uint32(v8))
	m.g0 = v4 + i32(32)
}
func (m *Module) fn1063(v0, v1 int32) int32 {
	var v2 int32
	v2 = i32(1)
	{
		t0 := m.fn159(v0, v1, i32(1283588), i32(5))
		if t0 == 0 {
			goto l0
		}
		if uint32(v1) < uint32(i32(6)) {
			goto l1
		}
		{
			t1 := int32(m.memory[int64(uint32(v0))+5])
			v0 = t1
			v1 = v0 + i32(-9)
			if uint32(v1) > uint32(i32(23)) {
				goto l2
			}
			if i32_shl(i32(1), v1)&i32(8388627) != 0 {
				goto l1
			}
		}
	l2:
		if v0 != i32(63) {
			goto l0
		}
	l1:
		v2 = i32(2)
	}
l0:
	return v2
}
func (m *Module) fn1064(v0, v1 int32) {
	var v2 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	m.fn1705(v2+i32(4), v1)
	t1 := int32(load32(m.memory[int64(uint32(v2))+8:]))
	v1 = t1
	{
		t2 := int32(load32(m.memory[int64(uint32(v2))+4:]))
		if t2 != i32(1) {
			goto l0
		}
		t3 := int32(load32(m.memory[int64(uint32(v2))+12:]))
		m.fn2(v1, t3)
		panic("unreachable")
	}
l0:
	t4 := int32(load32(m.memory[int64(uint32(v2))+12:]))
	store32(m.memory[int64(uint32(v0))+4:], uint32(t4))
	store32(m.memory[uint32(v0):], uint32(v1))
	m.g0 = v2 + i32(16)
}
func (m *Module) fn1065(v0, v1, v2, v3, v4 int32) {
	var v5 int32
	t0 := m.g0
	v5 = t0 - i32(16)
	m.g0 = v5
	m.fn1761(v5+i32(8), v1, v2, v3, v3+v4)
	v4 = i32(1)
	{
		{
			t1 := int32(load32(m.memory[int64(uint32(v5))+8:]))
			if t1 == i32(1) {
				goto l0
			}
			v4 = i32(0)
			goto l1
		}
	l0:
		t2 := int32(load32(m.memory[int64(uint32(v5))+12:]))
		v3 = t2 - v3
	}
l1:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
	store32(m.memory[uint32(v0):], uint32(v4))
	m.g0 = v5 + i32(16)
}
func (m *Module) fn1066(v0, v1 int32) {
	if v0 == i32(-1) {
		return
	}
	m.fn1091(v0, v1)
}
func fn1067(v0 int32) int32 {
	p0 := i32(2)
	if uint32(v0&i32(255)) > uint32(i32(8)) {
		p0 = v0 + i32(-9)
	}
	return i32_shr_u(i32(262917), p0&i32(255)<<3)
}
func (m *Module) fn1068(v0, v1, v2 int32) {
	var v3, v4, v5, v6, v7, v8, v9, v10, v11 int32
	t0 := m.g0
	v3 = t0 - i32(112)
	m.g0 = v3
	t1 := int32(load16(m.memory[int64(uint32(v1))+28:]))
	store16(m.memory[int64(uint32(v1))+28:], uint16(t1+i32(1)))
	store64(m.memory[int64(uint32(v3))+24:], uint64(i64(0)))
	store64(m.memory[int64(uint32(v3))+16:], uint64(i64(0x400000000)))
	store32(m.memory[int64(uint32(v3))+8:], uint32(i32(1)))
	store16(m.memory[int64(uint32(v3))+44:], uint16(i32(0)))
	t2 := int64(load64(m.memory[int64(uint32(v2))+4:]))
	store64(m.memory[int64(uint32(v3))+48:], uint64(t2))
	t3 := int32(load32(m.memory[int64(uint32(v2))+12:]))
	store32(m.memory[int64(uint32(v3))+56:], uint32(t3))
	t4 := int32(load32(m.memory[int64(uint32(v2))+16:]))
	store32(m.memory[int64(uint32(v3))+12:], uint32(t4))
	v4 = v1 + i32(12)
	v5 = i32(0)
l14:
	{
		m.fn1767(v3+i32(60), v3+i32(8))
		{
			t5 := int32(load32(m.memory[int64(uint32(v3))+60:]))
			v6 = t5
			if v6 > i32(-2) {
				goto l0
			}
			m.fn1768(v3 + i32(8))
			store32(m.memory[uint32(v0):], uint32(i32(-1)))
			goto l1
		}
	l0:
		t6 := int32(load32(m.memory[int64(uint32(v3))+68:]))
		v7 = t6
		t7 := int32(load32(m.memory[int64(uint32(v3))+64:]))
		v2 = t7
		{
			t8 := int32(load32(m.memory[int64(uint32(v3))+72:]))
			v8 = t8
			t9 := int32(load32(m.memory[int64(uint32(v3))+76:]))
			t10 := v8
			v9 = t9
			t11 := m.fn159(t10, v9, i32(1282344), i32(5))
			if t11 == 0 {
				goto l2
			}
			v10 = i32(0)
			{
				if uint32(v9) < uint32(i32(6)) {
					goto l3
				}
				t12 := int32(m.memory[int64(uint32(v8))+5])
				if t12 != i32(58) {
					goto l2
				}
				m.fn148(v3, i32(6), v8, v9, i32(1282352))
				t13 := int32(load32(m.memory[int64(uint32(v3))+4:]))
				v11 = t13
				t14 := int32(load32(m.memory[uint32(v3):]))
				v10 = t14
			}
		l3:
			{
				{
					t15 := int32(load32(m.memory[int64(uint32(v1))+24:]))
					t16 := v5
					v9 = t15
					if uint32(t16) >= uint32(v9) {
						store32(m.memory[int64(uint32(v0))+4:], uint32(v9))
						store32(m.memory[uint32(v0):], uint32(i32(5)))
						goto l11
					}
					t17 := int32(load16(m.memory[int64(uint32(v1))+28:]))
					v9 = t17
					if v10 == 0 {
						t27 := int32(load32(m.memory[int64(uint32(v1))+8:]))
						v10 = t27
						m.fn1059(v1, v2, v7)
						store16(m.memory[int64(uint32(v3))+108:], uint16(v9))
						store32(m.memory[int64(uint32(v3))+104:], uint32(v7))
						store32(m.memory[int64(uint32(v3))+100:], uint32(i32(0)))
						store32(m.memory[int64(uint32(v3))+96:], uint32(v10))
						m.fn1058(v4, v3+i32(96))
						goto l9
					}
					switch v11 + i32(-3) {
					default:
						goto l7
					case 0:
						t18 := int32(m.memory[uint32(v10)])
						if t18 != i32(120) {
							goto l7
						}
						t19 := int32(m.memory[int64(uint32(v10))+1])
						if t19 != i32(109) {
							goto l7
						}
						t20 := int32(m.memory[int64(uint32(v10))+2])
						if t20 != i32(108) {
							goto l7
						}
						t21 := m.fn882(v2, v7, i32(1282584), i32(36))
						if t21 != 0 {
							goto l9
						}
						m.fn884(v3+i32(84), v2, v7)
						v7 = i32(1)
						goto l10
					case 2:
						t22 := int32(m.memory[uint32(v10)])
						if t22 != i32(120) {
							goto l7
						}
						t23 := int32(m.memory[int64(uint32(v10))+1])
						if t23 != i32(109) {
							goto l7
						}
						t24 := int32(m.memory[int64(uint32(v10))+2])
						if t24 != i32(108) {
							goto l7
						}
						t25 := int32(m.memory[int64(uint32(v10))+3])
						if t25 != i32(110) {
							goto l7
						}
						t26 := int32(m.memory[int64(uint32(v10))+4])
						if t26 != i32(115) {
							goto l7
						}
						m.fn884(v3+i32(84), v2, v7)
						v7 = i32(2)
						goto l10
					}
				}
			l7:
				{
					t28 := m.fn882(v2, v7, i32(1282584), i32(36))
					if t28 != 0 {
						goto l12
					}
					{
						t29 := m.fn882(v2, v7, i32(1282620), i32(29))
						if t29 != 0 {
							v7 = i32(4)
							m.fn884(v3+i32(80)+i32(4), v10, v11)
							goto l10
						}
						t30 := int32(load32(m.memory[int64(uint32(v1))+8:]))
						v8 = t30
						m.fn1059(v1, v10, v11)
						m.fn1059(v1, v2, v7)
						store16(m.memory[int64(uint32(v3))+108:], uint16(v9))
						store32(m.memory[int64(uint32(v3))+104:], uint32(v7))
						store32(m.memory[int64(uint32(v3))+100:], uint32(v11))
						store32(m.memory[int64(uint32(v3))+96:], uint32(v8))
						m.fn1058(v4, v3+i32(96))
						goto l9
					}
				}
			l12:
				m.fn884(v3+i32(84), v10, v11)
				v7 = i32(3)
			l10:
				t31 := int32(load32(m.memory[int64(uint32(v3))+92:]))
				store32(m.memory[int64(uint32(v0))+12:], uint32(t31))
				t32 := int64(load64(m.memory[int64(uint32(v3))+84:]))
				store64(m.memory[int64(uint32(v0))+4:], uint64(t32))
				store32(m.memory[uint32(v0):], uint32(v7))
				goto l11
			}
		l9:
			v5 = v5 + i32(1)
		}
	l2:
		m.fn1066(v6, v2)
		goto l14
	l11:
	}
	m.fn1066(v6, v2)
	m.fn1768(v3 + i32(8))
l1:
	m.g0 = v3 + i32(112)
}
func (m *Module) fn1069(v0, v1 int32) int32 {
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
			store32(m.memory[int64(uint32(v2))+28:], uint32(v0+i32(4)))
			store32(m.memory[int64(uint32(v2))+16:], uint32(i32(151)))
			store32(m.memory[int64(uint32(v2))+12:], uint32(v2+i32(28)))
			t3 := int32(load32(m.memory[uint32(v1):]))
			t4 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t5 := m.fn284(t3, t4, i32(1051739), v2+i32(12))
			v1 = t5
			goto l7
		case 1:
			store32(m.memory[int64(uint32(v2))+28:], uint32(v0+i32(4)))
			store32(m.memory[int64(uint32(v2))+16:], uint32(i32(152)))
			store32(m.memory[int64(uint32(v2))+12:], uint32(v2+i32(28)))
			t6 := int32(load32(m.memory[uint32(v1):]))
			t7 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t8 := m.fn284(t6, t7, i32(1051459), v2+i32(12))
			v1 = t8
			goto l7
		case 2:
			store32(m.memory[int64(uint32(v2))+28:], uint32(v0))
			store32(m.memory[int64(uint32(v2))+16:], uint32(i32(153)))
			store32(m.memory[int64(uint32(v2))+12:], uint32(v2+i32(28)))
			t9 := int32(load32(m.memory[uint32(v1):]))
			t10 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t11 := m.fn284(t9, t10, i32(1051262), v2+i32(12))
			v1 = t11
			goto l7
		case 3:
			store32(m.memory[int64(uint32(v2))+28:], uint32(v0+i32(4)))
			store32(m.memory[int64(uint32(v2))+16:], uint32(i32(154)))
			store32(m.memory[int64(uint32(v2))+12:], uint32(v2+i32(28)))
			t12 := int32(load32(m.memory[uint32(v1):]))
			t13 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t14 := m.fn284(t12, t13, i32(1052169), v2+i32(12))
			v1 = t14
			goto l7
		case 4:
			t15 := int32(load32(m.memory[uint32(v1):]))
			t16 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t17 := m.fn1604(v0+i32(4), t15, t16)
			v1 = t17
			goto l7
		case 5:
			t18 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			v3 = t18
			t19 := int32(load32(m.memory[uint32(v1):]))
			v1 = t19
			{
				t20 := int32(load32(m.memory[int64(uint32(v0))+4:]))
				v4 = t20
				switch v4 >> 31 & (v4 + i32(-0x7fffffff)) {
				default:
					store32(m.memory[int64(uint32(v2))+8:], uint32(v0+i32(16)))
					store32(m.memory[int64(uint32(v2))+28:], uint32(v0+i32(4)))
					store32(m.memory[int64(uint32(v2))+24:], uint32(i32(36)))
					store32(m.memory[int64(uint32(v2))+16:], uint32(i32(155)))
					store32(m.memory[int64(uint32(v2))+20:], uint32(v2+i32(28)))
					store32(m.memory[int64(uint32(v2))+12:], uint32(v2+i32(8)))
					t21 := m.fn284(v1, v3, i32(1068475), v2+i32(12))
					v1 = t21
					goto l7
				case 1:
					store32(m.memory[int64(uint32(v2))+28:], uint32(v0+i32(8)))
					store32(m.memory[int64(uint32(v2))+16:], uint32(i32(155)))
					store32(m.memory[int64(uint32(v2))+12:], uint32(v2+i32(28)))
					t22 := m.fn284(v1, v3, i32(1069689), v2+i32(12))
					v1 = t22
					goto l7
				case 2:
					store32(m.memory[int64(uint32(v2))+28:], uint32(v0+i32(8)))
					store32(m.memory[int64(uint32(v2))+16:], uint32(i32(156)))
					store32(m.memory[int64(uint32(v2))+12:], uint32(v2+i32(28)))
					t23 := m.fn284(v1, v3, i32(1052336), v2+i32(12))
					v1 = t23
					goto l7
				case 3:
					t24 := int32(load32(m.memory[int64(uint32(v3))+12:]))
					t25 := m.t0[uint(t24)].(func(int32, int32, int32) int32)(v1, i32(1283903), i32(46))
					v1 = t25
					goto l7
				}
			}
		case 6:
			t26 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			v4 = t26
			t27 := int32(load32(m.memory[uint32(v1):]))
			v3 = t27
			{
				t28 := int32(load32(m.memory[int64(uint32(v0))+4:]))
				switch t28 {
				default:
					v1 = i32(1)
					t29 := int32(load32(m.memory[int64(uint32(v4))+12:]))
					t30 := v3
					v5 = t29
					t31 := m.t0[uint(v5)].(func(int32, int32, int32) int32)(t30, i32(1282798), i32(26))
					if t31 != 0 {
						goto l7
					}
					t32 := int32(load32(m.memory[int64(uint32(v0))+12:]))
					t33 := int32(load32(m.memory[int64(uint32(v0))+16:]))
					t34 := m.fn1780(v3, v4, t32, t33)
					if t34 != 0 {
						goto l7
					}
					t35 := m.t0[uint(v5)].(func(int32, int32, int32) int32)(v3, i32(1282679), i32(1))
					v1 = t35
					goto l7
				case 1:
					v1 = i32(1)
					t36 := int32(load32(m.memory[int64(uint32(v4))+12:]))
					t37 := v3
					v5 = t36
					t38 := m.t0[uint(v5)].(func(int32, int32, int32) int32)(t37, i32(1282824), i32(47))
					if t38 != 0 {
						goto l7
					}
					t39 := int32(load32(m.memory[int64(uint32(v0))+12:]))
					t40 := int32(load32(m.memory[int64(uint32(v0))+16:]))
					t41 := m.fn1780(v3, v4, t39, t40)
					if t41 != 0 {
						goto l7
					}
					t42 := m.t0[uint(v5)].(func(int32, int32, int32) int32)(v3, i32(1282679), i32(1))
					v1 = t42
					goto l7
				case 2:
					v1 = i32(1)
					t43 := int32(load32(m.memory[int64(uint32(v4))+12:]))
					t44 := v3
					v5 = t43
					t45 := m.t0[uint(v5)].(func(int32, int32, int32) int32)(t44, i32(1282871), i32(49))
					if t45 != 0 {
						goto l7
					}
					t46 := int32(load32(m.memory[int64(uint32(v0))+12:]))
					t47 := int32(load32(m.memory[int64(uint32(v0))+16:]))
					t48 := m.fn1780(v3, v4, t46, t47)
					if t48 != 0 {
						goto l7
					}
					t49 := m.t0[uint(v5)].(func(int32, int32, int32) int32)(v3, i32(1282679), i32(1))
					v1 = t49
					goto l7
				case 3:
					v1 = i32(1)
					t50 := int32(load32(m.memory[int64(uint32(v4))+12:]))
					t51 := v3
					v5 = t50
					t52 := m.t0[uint(v5)].(func(int32, int32, int32) int32)(t51, i32(1282920), i32(22))
					if t52 != 0 {
						goto l7
					}
					t53 := int32(load32(m.memory[int64(uint32(v0))+12:]))
					t54 := int32(load32(m.memory[int64(uint32(v0))+16:]))
					t55 := m.fn1780(v3, v4, t53, t54)
					if t55 != 0 {
						goto l7
					}
					t56 := m.t0[uint(v5)].(func(int32, int32, int32) int32)(v3, i32(1282942), i32(59))
					v1 = t56
					goto l7
				case 4:
					v1 = i32(1)
					t57 := int32(load32(m.memory[int64(uint32(v4))+12:]))
					t58 := v3
					v5 = t57
					t59 := m.t0[uint(v5)].(func(int32, int32, int32) int32)(t58, i32(1282920), i32(22))
					if t59 != 0 {
						goto l7
					}
					t60 := int32(load32(m.memory[int64(uint32(v0))+12:]))
					t61 := int32(load32(m.memory[int64(uint32(v0))+16:]))
					t62 := m.fn1780(v3, v4, t60, t61)
					if t62 != 0 {
						goto l7
					}
					t63 := m.t0[uint(v5)].(func(int32, int32, int32) int32)(v3, i32(1283001), i32(52))
					v1 = t63
					goto l7
				case 5:
					store32(m.memory[int64(uint32(v2))+28:], uint32(v0+i32(8)))
					store32(m.memory[int64(uint32(v2))+16:], uint32(i32(141)))
					store32(m.memory[int64(uint32(v2))+12:], uint32(v2+i32(28)))
					t64 := m.fn284(v3, v4, i32(1052809), v2+i32(12))
					v1 = t64
				}
			}
		}
	}
l7:
	m.g0 = v2 + i32(32)
	return v1
}
func (m *Module) fn1070(v0, v1, v2, v3 int32) {
	var v4, v5 int32
	var v6 int64
	var v7 int32
	var v8 int64
	t0 := m.g0
	v4 = t0 - i32(16)
	m.g0 = v4
	{
		switch v2 {
		case 0:
			m.memory[int64(uint32(v0))+1] = byte(i32(0))
			goto l3
		case 1:
			t1 := int32(m.memory[uint32(v1)])
			v5 = t1
			switch v5 + i32(-43) {
			case 0, 2:
				goto l4
			default:
				goto l5
			}
		default:
			t2 := int32(m.memory[uint32(v1)])
			v5 = t2
		}
	l5:
		t3 := v1
		var p4 int32
		if v5&i32(255) == i32(43) {
			p4 = 1
		}
		v5 = p4
		v1 = t3 + v5
		v2 = v2 - v5
		if uint32(v2) < uint32(i32(9)) {
			v5 = i32(0)
		l11:
			{
				if v2 == 0 {
					goto l7
				}
				t8 := int32(m.memory[uint32(v1)])
				m.fn199(v4, t8, v3)
				t9 := int32(load32(m.memory[uint32(v4):]))
				if t9 != i32(1) {
					goto l4
				}
				v1 = v1 + i32(1)
				v2 = v2 + i32(-1)
				t10 := int32(load32(m.memory[int64(uint32(v4))+4:]))
				v5 = t10 + v5*v3
				goto l11
			}
		}
		v5 = i32(0)
		v6 = int64(uint32(v3))
	l9:
		{
			if v2 == 0 {
				goto l7
			}
			t5 := int32(m.memory[uint32(v1)])
			m.fn199(v4+i32(8), t5, v3)
			t6 := int32(load32(m.memory[int64(uint32(v4))+8:]))
			v7 = t6
			v8 = int64(uint32(v5)) * v6
			if int32(int64(uint64(v8)>>32)) != 0 {
				v2 = i32(1)
				if v7&i32(1) == 0 {
					goto l4
				}
				m.memory[int64(uint32(v0))+1] = byte(i32(2))
				goto l10
			}
			if v7&i32(1) == 0 {
				goto l4
			}
			v1 = v1 + i32(1)
			v2 = v2 + i32(-1)
			t7 := int32(load32(m.memory[int64(uint32(v4))+12:]))
			v7 = t7
			v5 = v7 + int32(v8)
			if uint32(v5) >= uint32(v7) {
				goto l9
			}
		}
		m.memory[int64(uint32(v0))+1] = byte(i32(2))
		goto l3
	}
l4:
	v2 = i32(1)
	m.memory[int64(uint32(v0))+1] = byte(i32(1))
	goto l10
l7:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
	v2 = i32(0)
	goto l10
l3:
	v2 = i32(1)
l10:
	m.memory[uint32(v0)] = byte(v2)
	m.g0 = v4 + i32(16)
}
func (m *Module) fn1071(v0, v1, v2 int32) {
	m.fn1070(v0, v1, v2, i32(10))
}
func (m *Module) fn1072(v0, v1 int32) {
	var v2 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	store32(m.memory[int64(uint32(v2))+12:], uint32(i32(0)))
	m.fn522(v2, v1, v2+i32(12))
	t1 := int32(load32(m.memory[uint32(v2):]))
	t2 := int32(load32(m.memory[int64(uint32(v2))+4:]))
	m.fn51(v0, t1, t2)
	m.g0 = v2 + i32(16)
}
func (m *Module) fn1073(v0, v1 int32) {
	var v2, v3 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	t1 := int32(load32(m.memory[uint32(v1):]))
	t2 := v1
	v3 = t1 + i32(1)
	store32(m.memory[uint32(t2):], uint32(v3))
	v1 = i32(-1)
	if uint32(v3) < uint32(i32(2000001)) {
		goto l0
	}
	store32(m.memory[int64(uint32(v2))+12:], uint32(i32(5)))
	store32(m.memory[int64(uint32(v2))+8:], uint32(i32(1084248)))
	m.fn73(v0+i32(4), i32(1067003), v2+i32(8))
	store32(m.memory[int64(uint32(v0))+20:], uint32(i32(13)))
	store32(m.memory[int64(uint32(v0))+16:], uint32(i32(1084252)))
	v1 = i32(-0x7ffffffd)
l0:
	store32(m.memory[uint32(v0):], uint32(v1))
	m.g0 = v2 + i32(16)
}
func (m *Module) fn1074(v0, v1, v2, v3 int32) {
	var v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14, v15, v16, v17, v18, v19, v20, v21, v22 int32
	var v23 int64
	var v24, v25, v26, v27, v28 int32
	var v29 int64
	t0 := m.g0
	v4 = t0 - i32(272)
	m.g0 = v4
	m.fn164(v4+i32(56), v1)
	t1 := v4 + i32(92)
	v5 = v2 + i32(72)
	t2 := int32(load32(m.memory[int64(uint32(v4))+56:]))
	t3 := int32(load32(m.memory[int64(uint32(v4))+60:]))
	m.fn876(t1, v5, t2, t3)
	t4 := int32(load32(m.memory[int64(uint32(v4))+108:]))
	v6 = t4
	t5 := int32(load32(m.memory[int64(uint32(v4))+104:]))
	v7 = t5
	t6 := int32(load32(m.memory[int64(uint32(v4))+96:]))
	v8 = t6
	v9 = i32(0)
	{
		t7 := int32(load32(m.memory[int64(uint32(v4))+92:]))
		v10 = t7
		if v10 != i32(-0x7fffffff) {
			goto l0
		}
		t8 := int32(load32(m.memory[int64(uint32(v4))+100:]))
		m.fn1083(v4+i32(48), v3, v8, t8)
		t9 := int32(load32(m.memory[int64(uint32(v4))+52:]))
		v11 = t9
		t10 := int32(load32(m.memory[int64(uint32(v4))+48:]))
		v9 = t10
	}
l0:
	m.fn92(v4+i32(92), v7, v6)
	m.fn490(v4+i32(68), v4+i32(92))
	store32(m.memory[int64(uint32(v4))+88:], uint32(i32(0)))
	store64(m.memory[int64(uint32(v4))+80:], uint64(i64(0x400000000)))
	t11 := int32(load32(m.memory[int64(uint32(v1))+16:]))
	v12 = t11
	t12 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	v6 = t12
	t13 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	v7 = t13
	t14 := int32(load32(m.memory[int64(uint32(v1))+12:]))
	store32(m.memory[int64(uint32(v4))+140:], uint32(t14))
	store32(m.memory[int64(uint32(v4))+136:], uint32(v7))
	store32(m.memory[int64(uint32(v4))+132:], uint32(v6))
	store16(m.memory[int64(uint32(v4))+128:], uint16(i32(256)))
	store64(m.memory[int64(uint32(v4))+108:], uint64(i64(0)))
	store64(m.memory[int64(uint32(v4))+100:], uint64(i64(0x400000000)))
	store32(m.memory[int64(uint32(v4))+96:], uint32(v12))
	store32(m.memory[int64(uint32(v4))+92:], uint32(i32(1)))
	v13 = v4 + i32(184)
	v14 = i32(4)
	v15 = v4 + i32(144) + i32(4)
	v16 = v4 + i32(172) + i32(4)
	t15 := int32(load32(m.memory[int64(uint32(v2))+4:]))
	v17 = t15
	v18 = i32(0)
l27:
	{
		m.fn1084(v4+i32(172), v4+i32(92), v6, v7)
		{
			{
				{
					t16 := int32(load32(m.memory[int64(uint32(v4))+172:]))
					v6 = t16
					switch v6 + i32(2) {
					case 1:
						goto l2
					default:
						t17 := int32(load32(m.memory[int64(uint32(v4))+188:]))
						v12 = t17
						t18 := int32(load32(m.memory[int64(uint32(v4))+184:]))
						v19 = t18
						t19 := int32(load32(m.memory[int64(uint32(v4))+132:]))
						t20 := v4 + i32(32)
						v20 = t19
						t21 := int32(load32(m.memory[int64(uint32(v4))+136:]))
						t22 := v20
						v21 = t21
						t23 := int32(load32(m.memory[int64(uint32(v4))+176:]))
						t24 := int32(load32(m.memory[int64(uint32(v4))+180:]))
						m.fn1085(t20, t22, v21, t23, t24)
						t25 := int32(load32(m.memory[int64(uint32(v4))+36:]))
						v7 = t25
						t26 := int32(load32(m.memory[int64(uint32(v4))+32:]))
						v1 = t26
						if v6 != i32(3) {
							goto l4
						}
						v6 = i32(1)
						v20 = i32(0)
						goto l5
					case 0:
						t27 := int32(load32(m.memory[int64(uint32(v4))+100:]))
						t28 := int32(load32(m.memory[int64(uint32(v4))+104:]))
						m.fn1086(t27, t28)
						{
							t29 := int32(load32(m.memory[int64(uint32(v4))+112:]))
							v6 = t29
							if v6 == 0 {
								goto l6
							}
							t30 := int32(load32(m.memory[int64(uint32(v4))+116:]))
							m.fn1087(v6, t30)
						}
					l6:
						{
							t31 := int32(load32(m.memory[int64(uint32(v4))+72:]))
							t32 := int32(load32(m.memory[int64(uint32(v4))+76:]))
							t33 := m.fn773(t31, t32, i32(1080683), i32(6))
							if t33 == 0 {
								goto l7
							}
							p34 := i32(0)
							if v9 != 0 {
								p34 = v9 + i32(8)
							}
							t35 := m.fn848(p34, v11, i32(1073848), i32(59))
							if t35 == 0 {
								goto l7
							}
							v7 = v18 << 5
							v1 = v14 + i32(8)
						l8:
							{
								v6 = v1
								if v7 == 0 {
									goto l7
								}
								v7 = v7 + i32(-32)
								v1 = v6 + i32(32)
								t36 := int32(load32(m.memory[uint32(v6+i32(-4)):]))
								t37 := int32(load32(m.memory[uint32(v6):]))
								t38 := m.fn773(t36, t37, i32(1074595), i32(8))
								if t38 == 0 {
									goto l8
								}
							}
							v19 = v6 + i32(-8)
							t39 := int32(load32(m.memory[uint32(v19+i32(16)):]))
							v6 = t39
							t40 := int32(load32(m.memory[int64(uint32(v19))+20:]))
							v7 = t40
							store16(m.memory[int64(uint32(v4))+204:], uint16(i32(1)))
							v1 = i32(0)
							store32(m.memory[int64(uint32(v4))+200:], uint32(i32(0)))
							store32(m.memory[int64(uint32(v4))+196:], uint32(v6+v7))
							store32(m.memory[int64(uint32(v4))+192:], uint32(v6))
							store32(m.memory[int64(uint32(v4))+188:], uint32(v7))
							store32(m.memory[int64(uint32(v4))+184:], uint32(v6))
							store32(m.memory[int64(uint32(v4))+180:], uint32(v7))
							store32(m.memory[int64(uint32(v4))+176:], uint32(i32(0)))
							store32(m.memory[int64(uint32(v4))+172:], uint32(v2))
							v20 = v19 + i32(12)
							m.fn874(v4+i32(240), v4+i32(172))
							{
								t41 := int32(load32(m.memory[int64(uint32(v4))+240:]))
								if t41 == i32(-1) {
									goto l9
								}
								m.fn59(v4+i32(40), i32(4), i32(4), i32(12))
								t42 := int32(load32(m.memory[int64(uint32(v4))+40:]))
								v6 = t42
								t43 := int32(load32(m.memory[int64(uint32(v4))+44:]))
								v12 = t43
								t44 := int32(load32(m.memory[int64(uint32(v4))+248:]))
								store32(m.memory[int64(uint32(v12))+8:], uint32(t44))
								t45 := int64(load64(m.memory[int64(uint32(v4))+240:]))
								store64(m.memory[uint32(v12):], uint64(t45))
								store32(m.memory[int64(uint32(v4))+216:], uint32(i32(1)))
								store32(m.memory[int64(uint32(v4))+212:], uint32(v12))
								store32(m.memory[int64(uint32(v4))+208:], uint32(v6))
								memory_copy(m.memory, uint32(v4+i32(92)), uint32(v4+i32(172)), uint32(i32(36)))
								v6 = i32(12)
								v7 = i32(1)
							l12:
								{
									m.fn874(v4+i32(252), v4+i32(92))
									t46 := int32(load32(m.memory[int64(uint32(v4))+252:]))
									if t46 == i32(-1) {
										t52 := int32(load32(m.memory[int64(uint32(v4))+216:]))
										v1 = t52
										t53 := int32(load32(m.memory[int64(uint32(v4))+212:]))
										v6 = t53
										goto l13
									}
									{
										t47 := int32(load32(m.memory[int64(uint32(v4))+208:]))
										if v7 != t47 {
											goto l11
										}
										m.fn60(v4+i32(208), i32(1))
										t48 := int32(load32(m.memory[int64(uint32(v4))+212:]))
										v12 = t48
									}
								l11:
									v1 = v12 + v6
									t49 := int32(load32(m.memory[int64(uint32(v4))+260:]))
									store32(m.memory[int64(uint32(v1))+8:], uint32(t49))
									t50 := int64(load64(m.memory[int64(uint32(v4))+252:]))
									store64(m.memory[uint32(v1):], uint64(t50))
									t51 := v4
									v7 = v7 + i32(1)
									store32(m.memory[int64(uint32(t51))+216:], uint32(v7))
									v6 = v6 + i32(12)
									goto l12
								}
							}
						l9:
							store32(m.memory[int64(uint32(v4))+216:], uint32(i32(0)))
							store64(m.memory[int64(uint32(v4))+208:], uint64(i64(0x400000000)))
							v6 = i32(4)
						l13:
							m.fn77(v4+i32(92), v6, v1, i32(1097368), i32(1))
							t54 := int32(load32(m.memory[int64(uint32(v19))+12:]))
							t55 := int32(load32(m.memory[uint32(v19+i32(16)):]))
							m.fn16(t54, t55)
							t56 := int32(load32(m.memory[int64(uint32(v4))+100:]))
							store32(m.memory[int64(uint32(v20))+8:], uint32(t56))
							t57 := int64(load64(m.memory[int64(uint32(v4))+92:]))
							store64(m.memory[uint32(v20):], uint64(t57))
							m.fn78(v4 + i32(208))
						}
					l7:
						t58 := int32(load32(m.memory[int64(uint32(v4))+88:]))
						store32(m.memory[int64(uint32(v0))+20:], uint32(t58))
						t59 := int64(load64(m.memory[int64(uint32(v4))+80:]))
						store64(m.memory[int64(uint32(v0))+12:], uint64(t59))
						store32(m.memory[int64(uint32(v0))+40:], uint32(v11))
						store32(m.memory[int64(uint32(v0))+36:], uint32(v9))
						t60 := int64(load64(m.memory[int64(uint32(v4))+68:]))
						store64(m.memory[uint32(v0):], uint64(t60))
						t61 := int32(load32(m.memory[int64(uint32(v4))+76:]))
						store32(m.memory[int64(uint32(v0))+8:], uint32(t61))
						store32(m.memory[int64(uint32(v0))+32:], uint32(i32(0)))
						store64(m.memory[int64(uint32(v0))+24:], uint64(i64(0x400000000)))
						m.fn879(v10, v8)
						m.g0 = v4 + i32(272)
						return
					}
				}
			l4:
				m.fn1085(v4+i32(24), v20, v21, v19, v12)
				t62 := int32(load32(m.memory[int64(uint32(v4))+24:]))
				v6 = t62
				t63 := int32(load32(m.memory[int64(uint32(v4))+28:]))
				v20 = t63
			}
		l5:
			{
				if uint32(v7) < uint32(i32(5)) {
					goto l14
				}
				m.fn309(v4+i32(172), v1, v7, i32(5), i32(0x105554))
				t64 := int32(load32(m.memory[int64(uint32(v4))+176:]))
				if t64 != i32(5) {
					goto l14
				}
				t65 := int32(load32(m.memory[int64(uint32(v4))+184:]))
				v12 = t65
				t66 := int32(load32(m.memory[int64(uint32(v4))+180:]))
				v19 = t66
				t67 := int32(load32(m.memory[int64(uint32(v4))+172:]))
				t68 := m.fn235(t67, i32(1282344), i32(5))
				if t68 == 0 {
					goto l14
				}
				if v12 == 0 {
					goto l2
				}
				t69 := int32(m.memory[uint32(v19)])
				if t69 == i32(58) {
					goto l2
				}
			}
		l14:
			v12 = i32(0)
			m.fn880(v4+i32(172), v5, v1, v7, i32(0))
			t70 := int32(load32(m.memory[int64(uint32(v4))+188:]))
			v7 = t70
			t71 := int32(load32(m.memory[int64(uint32(v4))+184:]))
			v21 = t71
			t72 := int32(load32(m.memory[int64(uint32(v4))+176:]))
			v1 = t72
			{
				t73 := int32(load32(m.memory[int64(uint32(v4))+172:]))
				v22 = t73
				if v22 != i32(-0x7fffffff) {
					goto l15
				}
				t74 := int32(load32(m.memory[int64(uint32(v4))+180:]))
				m.fn1083(v4+i32(16), v3, v1, t74)
				t75 := int32(load32(m.memory[int64(uint32(v4))+20:]))
				v19 = t75
				t76 := int32(load32(m.memory[int64(uint32(v4))+16:]))
				v12 = t76
			}
		l15:
			m.fn198(v4+i32(252), v6, v20, v17)
			t77 := int64(load64(m.memory[int64(uint32(v4))+256:]))
			v23 = t77
			{
				{
					{
						{
							t78 := int32(load32(m.memory[int64(uint32(v4))+252:]))
							v14 = t78
							if v14 != i32(-2) {
								goto l16
							}
							store64(m.memory[int64(uint32(v4))+176:], uint64(v23))
							store32(m.memory[int64(uint32(v4))+172:], uint32(i32(-0x7ffffff4)))
							goto l17
						}
					l16:
						store32(m.memory[int64(uint32(v4))+220:], uint32(v14))
						store64(m.memory[int64(uint32(v4))+224:], uint64(v23))
						t79 := v4
						v18 = int32(v23)
						store32(m.memory[int64(uint32(t79))+232:], uint32(v18))
						t80 := v4
						t81 := v18
						v24 = int32(int64(uint64(v23) >> 32))
						store32(m.memory[int64(uint32(t80))+236:], uint32(t81+v24))
						m.fn1088(v4+i32(8), v4+i32(232))
						{
							{
								t82 := int32(load32(m.memory[int64(uint32(v4))+8:]))
								if t82&i32(1) == 0 {
									goto l18
								}
								t83 := int32(load32(m.memory[int64(uint32(v4))+12:]))
								v25 = t83
								m.fn1064(v4, v24)
								store32(m.memory[int64(uint32(v4))+248:], uint32(i32(0)))
								t84 := int64(load64(m.memory[uint32(v4):]))
								store64(m.memory[int64(uint32(v4))+240:], uint64(t84))
								m.fn1089(v4+i32(252), v4+i32(240), v4+i32(232), v18, v24, i32(0), v25, i32(1))
								t85 := int32(load32(m.memory[int64(uint32(v4))+256:]))
								v25 = t85
								{
									t86 := int32(load32(m.memory[int64(uint32(v4))+252:]))
									v26 = t86
									if v26 == i32(-1) {
										m.fn1090(v4+i32(252), v4+i32(240), v4+i32(232), v18, v24, v25, i32(1))
										{
											t90 := int32(load32(m.memory[int64(uint32(v4))+252:]))
											v26 = t90
											if v26 == i32(-1) {
												t95 := int32(load32(m.memory[int64(uint32(v4))+240:]))
												v6 = t95
												if v6 == i32(-1) {
													goto l18
												}
												t96 := int64(load32(m.memory[int64(uint32(v4))+244:]))
												v23 = t96
												t97 := int64(load32(m.memory[int64(uint32(v4))+248:]))
												v29 = t97
												store32(m.memory[int64(uint32(v4))+176:], uint32(v6))
												store64(m.memory[int64(uint32(v4))+180:], uint64(v23|v29<<32))
												m.fn277(v14, v18)
												goto l22
											}
											t91 := int32(load32(m.memory[int64(uint32(v4))+268:]))
											v24 = t91
											t92 := int32(load32(m.memory[int64(uint32(v4))+264:]))
											v27 = t92
											t93 := int32(load32(m.memory[int64(uint32(v4))+260:]))
											v28 = t93
											t94 := int32(load32(m.memory[int64(uint32(v4))+256:]))
											v25 = t94
											goto l20
										}
									}
									t87 := int32(load32(m.memory[int64(uint32(v4))+268:]))
									v24 = t87
									t88 := int32(load32(m.memory[int64(uint32(v4))+264:]))
									v27 = t88
									t89 := int32(load32(m.memory[int64(uint32(v4))+260:]))
									v28 = t89
									goto l20
								}
							}
						l18:
							t98 := int32(load32(m.memory[int64(uint32(v4))+228:]))
							store32(m.memory[int64(uint32(v16))+8:], uint32(t98))
							t99 := int64(load64(m.memory[int64(uint32(v4))+220:]))
							store64(m.memory[uint32(v16):], uint64(t99))
						}
					l22:
						store32(m.memory[int64(uint32(v4))+172:], uint32(i32(-1)))
						goto l23
					l20:
						t100 := int32(load32(m.memory[int64(uint32(v4))+240:]))
						t101 := int32(load32(m.memory[int64(uint32(v4))+244:]))
						m.fn1091(t100, t101)
						store32(m.memory[int64(uint32(v4))+192:], uint32(v24))
						store32(m.memory[int64(uint32(v4))+188:], uint32(v27))
						store32(m.memory[int64(uint32(v4))+184:], uint32(v28))
						store32(m.memory[int64(uint32(v4))+180:], uint32(v25))
						store32(m.memory[int64(uint32(v4))+176:], uint32(v26))
						store32(m.memory[int64(uint32(v4))+172:], uint32(i32(-0x7ffffff3)))
						m.fn277(v14, v18)
						t102 := int32(load32(m.memory[int64(uint32(v4))+172:]))
						if t102 == i32(-1) {
							goto l23
						}
					}
				l17:
					t103 := int64(load64(m.memory[int64(uint32(v4))+188:]))
					store64(m.memory[int64(uint32(v4))+160:], uint64(t103))
					t104 := int64(load64(m.memory[int64(uint32(v4))+180:]))
					store64(m.memory[int64(uint32(v4))+152:], uint64(t104))
					t105 := int64(load64(m.memory[int64(uint32(v4))+172:]))
					t106 := v4
					v23 = t105
					store64(m.memory[int64(uint32(t106))+144:], uint64(v23))
					if int32(v23) == i32(-1) {
						goto l24
					}
					m.fn92(v4+i32(172), v6, v20)
					m.fn490(v4+i32(240), v4+i32(172))
					m.fn535(v4 + i32(144))
					goto l25
				}
			l23:
				m.fn490(v15, v16)
			l24:
				t107 := int32(load32(m.memory[int64(uint32(v15))+8:]))
				store32(m.memory[int64(uint32(v4))+248:], uint32(t107))
				t108 := int64(load64(m.memory[uint32(v15):]))
				store64(m.memory[int64(uint32(v4))+240:], uint64(t108))
			}
		l25:
			m.fn92(v4+i32(252), v21, v7)
			m.fn490(v4+i32(172), v4+i32(252))
			t109 := int32(load32(m.memory[int64(uint32(v4))+248:]))
			store32(m.memory[int64(uint32(v13))+8:], uint32(t109))
			t110 := int64(load64(m.memory[int64(uint32(v4))+240:]))
			store64(m.memory[uint32(v13):], uint64(t110))
			{
				t111 := int32(load32(m.memory[int64(uint32(v4))+88:]))
				v7 = t111
				t112 := int32(load32(m.memory[int64(uint32(v4))+80:]))
				if v7 != t112 {
					goto l26
				}
				m.fn434(v4 + i32(80))
			}
		l26:
			t113 := int32(load32(m.memory[int64(uint32(v4))+84:]))
			v14 = t113
			v6 = v14 + v7<<5
			t114 := int64(load64(m.memory[int64(uint32(v4))+172:]))
			store64(m.memory[uint32(v6):], uint64(t114))
			t115 := int64(load64(m.memory[int64(uint32(v4))+180:]))
			store64(m.memory[int64(uint32(v6))+8:], uint64(t115))
			t116 := int64(load64(m.memory[int64(uint32(v4))+188:]))
			store64(m.memory[int64(uint32(v6))+16:], uint64(t116))
			store32(m.memory[int64(uint32(v6))+28:], uint32(v19))
			store32(m.memory[int64(uint32(v6))+24:], uint32(v12))
			t117 := v4
			v18 = v7 + i32(1)
			store32(m.memory[int64(uint32(t117))+88:], uint32(v18))
			m.fn879(v22, v1)
		}
	l2:
		t118 := int32(load32(m.memory[int64(uint32(v4))+136:]))
		v7 = t118
		t119 := int32(load32(m.memory[int64(uint32(v4))+132:]))
		v6 = t119
		goto l27
	}
}
func (m *Module) fn1075(v0 int32) {
	var v1 int32
	t0 := m.g0
	v1 = t0 - i32(16)
	m.g0 = v1
	t1 := int32(load32(m.memory[uint32(v0):]))
	m.fn273(v1+i32(8), v0, t1, i32(1), i32(4), i32(44))
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
func (m *Module) fn1076(v0, v1, v2, v3 int32) {
	v2 = v2 + i32(24)
	t0 := v2
	v0 = v0 + v1*i32(44)
	p1 := v0 + i32(-20)
	if v0 == i32(44) {
		p1 = t0
	}
	p2 := v2
	if v1 != 0 {
		p2 = p1
	}
	m.fn1092(p2, v3)
}
