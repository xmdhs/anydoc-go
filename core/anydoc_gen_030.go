package core

import (
	"math"
	"math/bits"
)

func (m *Module) fn1302(v0 int32) {
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
	m.fn1353(v3)
	v3 = v3 + i32(56)
	goto l1
l0:
	t2 := int32(load32(m.memory[uint32(v0):]))
	m.fn1354(t2, v2)
}
func (m *Module) fn1303(v0 int32) {
	t0 := int32(load32(m.memory[uint32(v0):]))
	if t0 == i32(-1) {
		return
	}
	m.fn1355(v0)
}
func (m *Module) fn1304(v0 int32) {
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
	t3 := int64(load64(m.memory[int64(uint32(i32(0)))+1286264:]))
	store64(m.memory[int64(uint32(v0))+8:], uint64(t3))
	t4 := int64(load64(m.memory[int64(uint32(i32(0)))+1286256:]))
	store64(m.memory[uint32(v0):], uint64(t4))
	store64(m.memory[int64(uint32(v0))+24:], uint64(v3))
	store64(m.memory[int64(uint32(v0))+16:], uint64(v2))
	m.g0 = v1 + i32(16)
}
func (m *Module) fn1305(v0, v1, v2 int32) {
	var v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13 int32
	var v14, v15 int64
	var v16 int32
	t0 := m.g0
	v3 = t0 - i32(1344)
	m.g0 = v3
	store32(m.memory[int64(uint32(v3))+60:], uint32(v1))
	store32(m.memory[int64(uint32(v3))+64:], uint32(v1+v2*i32(44)))
	v4 = v0 + i32(448)
	v5 = v0 + i32(416)
	v6 = v0 + i32(432)
	v7 = v0 + i32(400)
	v8 = v0 + i32(464)
	v9 = v3 + i32(928) + i32(16)
l1:
	{
		t1 := m.fn904(v3 + i32(60))
		v1 = t1
		if v1 == 0 {
			goto l0
		}
		t2 := int32(load32(m.memory[int64(uint32(v1))+32:]))
		v2 = t2
		t3 := int32(load32(m.memory[int64(uint32(v1))+28:]))
		t4 := v3
		v1 = t3
		store32(m.memory[int64(uint32(t4))+68:], uint32(v1))
		store32(m.memory[int64(uint32(v3))+72:], uint32(v1+v2*i32(44)))
	l3:
		{
			t5 := m.fn904(v3 + i32(68))
			v1 = t5
			if v1 == 0 {
				goto l1
			}
			{
				t6 := m.fn847(v1, i32(1074169), i32(48), i32(1077344), i32(16))
				if t6 != 0 {
					goto l2
				}
				t7 := m.fn847(v1, i32(1074169), i32(48), i32(1074248), i32(6))
				if t7 == 0 {
					goto l3
				}
			}
		l2:
			t8 := int32(load32(m.memory[int64(uint32(v1))+32:]))
			v2 = t8
			t9 := int32(load32(m.memory[int64(uint32(v1))+28:]))
			t10 := v3
			v1 = t9
			store32(m.memory[int64(uint32(t10))+76:], uint32(v1))
			store32(m.memory[int64(uint32(v3))+80:], uint32(v1+v2*i32(44)))
		l14:
			{
				t11 := m.fn904(v3 + i32(76))
				v1 = t11
				if v1 == 0 {
					goto l3
				}
				{
					t12 := m.fn847(v1, i32(1077249), i32(47), i32(1077360), i32(13))
					if t12 == 0 {
						goto l4
					}
					t13 := int32(load32(m.memory[uint32(v1+i32(16)):]))
					t14 := int32(load32(m.memory[uint32(v1+i32(20)):]))
					m.fn1046(v3+i32(48), t13, t14, i32(1077249), i32(47), i32(1077373), i32(6))
					t15 := int32(load32(m.memory[int64(uint32(v3))+48:]))
					v2 = t15
					if v2 == 0 {
						goto l4
					}
					t16 := int32(load32(m.memory[int64(uint32(v3))+52:]))
					m.fn51(v3+i32(928), v2, t16)
					t17 := int32(load32(m.memory[uint32(v1+i32(28)):]))
					t18 := int32(load32(m.memory[uint32(v1+i32(32)):]))
					t19 := m.fn1327(t17, t18)
					m.fn1105(v8, v3+i32(928), t19)
				}
			l4:
				{
					t20 := m.fn847(v1, i32(1077249), i32(47), i32(1077144), i32(5))
					if t20 == 0 {
						goto l5
					}
					t21 := int32(load32(m.memory[uint32(v1+i32(16)):]))
					t22 := v3 + i32(40)
					v2 = t21
					t23 := int32(load32(m.memory[uint32(v1+i32(20)):]))
					t24 := v2
					v10 = t23
					m.fn1046(t22, t24, v10, i32(1077249), i32(47), i32(1073713), i32(4))
					t25 := int32(load32(m.memory[int64(uint32(v3))+40:]))
					v11 = t25
					if v11 == 0 {
						goto l5
					}
					t26 := int32(load32(m.memory[int64(uint32(v3))+44:]))
					v12 = t26
					m.fn1046(v3+i32(32), v2, v10, i32(1077249), i32(47), i32(1077373), i32(6))
					m.fn1046(v3+i32(24), v2, v10, i32(1077249), i32(47), i32(1077379), i32(17))
					t27 := int32(load32(m.memory[int64(uint32(v3))+36:]))
					t28 := int32(load32(m.memory[int64(uint32(v3))+32:]))
					v2 = t28
					p29 := i32(0)
					if v2 != 0 {
						p29 = t27
					}
					v10 = p29
					p30 := i32(1)
					if v2 != 0 {
						p30 = v2
					}
					v2 = p30
					{
						t31 := int32(load32(m.memory[int64(uint32(v3))+24:]))
						v13 = t31
						if v13 == 0 {
							goto l6
						}
						t32 := int32(load32(m.memory[int64(uint32(v3))+28:]))
						m.fn1328(v3+i32(84), v2, v10, v13, t32)
						goto l7
					}
				l6:
					store32(m.memory[int64(uint32(v3))+84:], uint32(i32(-1)))
				l7:
					m.fn1328(v3+i32(928), v2, v10, v11, v12)
					t33 := int64(load64(m.memory[int64(uint32(v0))+416:]))
					t34 := int64(load64(m.memory[int64(uint32(v0))+424:]))
					t35 := int32(load32(m.memory[int64(uint32(v3))+932:]))
					t36 := int32(load32(m.memory[int64(uint32(v3))+936:]))
					t37 := m.fn540(t33, t34, t35, t36)
					v14 = t37
					store32(m.memory[int64(uint32(v3))+96:], uint32(v3+i32(928)))
					{
						t38 := int32(load32(m.memory[int64(uint32(v0))+408:]))
						if t38 != 0 {
							goto l8
						}
						_ = m.fn679(v7, v5)
					}
				l8:
					store32(m.memory[int64(uint32(v3))+516:], uint32(v7))
					store32(m.memory[int64(uint32(v3))+512:], uint32(v3+i32(96)))
					t40 := int32(load32(m.memory[int64(uint32(v0))+400:]))
					t41 := int32(load32(m.memory[int64(uint32(v0))+404:]))
					m.fn69(v3+i32(16), t40, t41, v14, v3+i32(512), i32(184))
					t42 := int32(load32(m.memory[int64(uint32(v0))+400:]))
					v2 = t42
					t43 := int32(load32(m.memory[int64(uint32(v3))+20:]))
					v10 = t43
					{
						t44 := int32(load32(m.memory[int64(uint32(v3))+16:]))
						if t44 != i32(1) {
							goto l9
						}
						v11 = v2 + v10
						t45 := int32(m.memory[uint32(v11)])
						v12 = t45
						t46 := int32(load32(m.memory[int64(uint32(v3))+936:]))
						v13 = t46
						t47 := int64(load64(m.memory[int64(uint32(v3))+928:]))
						v15 = t47
						t48 := v11
						v16 = int32(uint32(int32(v14)) >> 25)
						m.memory[uint32(t48)] = byte(v16)
						t49 := int32(load32(m.memory[int64(uint32(v0))+404:]))
						m.memory[uint32(v2+t49&(v10+i32(-8))+i32(8))] = byte(v16)
						t50 := int32(load32(m.memory[int64(uint32(v0))+412:]))
						store32(m.memory[int64(uint32(v0))+412:], uint32(t50+i32(1)))
						t51 := int32(load32(m.memory[int64(uint32(v0))+408:]))
						store32(m.memory[int64(uint32(v0))+408:], uint32(t51-v12&i32(1)))
						v2 = v2 + (i32(0)-v10)*i32(28)
						v10 = v2 + i32(-28)
						store64(m.memory[uint32(v10):], uint64(v15))
						store32(m.memory[int64(uint32(v10))+8:], uint32(v13))
						store32(m.memory[uint32(v2+i32(-16)):], uint32(v1))
						v2 = v2 + i32(-12)
						t52 := int64(load64(m.memory[int64(uint32(v3))+84:]))
						store64(m.memory[uint32(v2):], uint64(t52))
						t53 := int32(load32(m.memory[int64(uint32(v3))+92:]))
						store32(m.memory[int64(uint32(v2))+8:], uint32(t53))
						goto l5
					}
				l9:
					v2 = v2 + (i32(0)-v10)*i32(28)
					store32(m.memory[uint32(v2+i32(-16)):], uint32(v1))
					v10 = v2 + i32(-12)
					t54 := int32(load32(m.memory[int64(uint32(v3))+92:]))
					store32(m.memory[int64(uint32(v10))+8:], uint32(t54))
					t55 := int32(load32(m.memory[uint32(v2+i32(-8)):]))
					v11 = t55
					t56 := int32(load32(m.memory[uint32(v10):]))
					v2 = t56
					t57 := int64(load64(m.memory[int64(uint32(v3))+84:]))
					store64(m.memory[uint32(v10):], uint64(t57))
					t58 := int32(load32(m.memory[int64(uint32(v3))+928:]))
					t59 := int32(load32(m.memory[int64(uint32(v3))+932:]))
					m.fn16(t58, t59)
					if v2 == i32(-2) {
						goto l5
					}
					m.fn134(v2, v11)
				}
			l5:
				{
					t60 := m.fn847(v1, i32(1074680), i32(46), i32(1077396), i32(10))
					if t60 == 0 {
						goto l10
					}
					t61 := int32(load32(m.memory[uint32(v1+i32(16)):]))
					t62 := int32(load32(m.memory[uint32(v1+i32(20)):]))
					m.fn1046(v3+i32(8), t61, t62, i32(1077249), i32(47), i32(1073713), i32(4))
					t63 := int32(load32(m.memory[int64(uint32(v3))+8:]))
					v2 = t63
					if v2 == 0 {
						goto l10
					}
					t64 := int32(load32(m.memory[int64(uint32(v3))+12:]))
					m.fn51(v3+i32(500), v2, t64)
					t65 := int32(load32(m.memory[uint32(v1+i32(28)):]))
					t66 := int32(load32(m.memory[uint32(v1+i32(32)):]))
					m.fn1329(v3+i32(512), t65, t66)
					t67 := int64(load64(m.memory[int64(uint32(v0))+448:]))
					t68 := int64(load64(m.memory[int64(uint32(v0))+456:]))
					t69 := int32(load32(m.memory[int64(uint32(v3))+504:]))
					t70 := int32(load32(m.memory[int64(uint32(v3))+508:]))
					t71 := m.fn540(t67, t68, t69, t70)
					v14 = t71
					store32(m.memory[int64(uint32(v3))+96:], uint32(v3+i32(500)))
					{
						t72 := int32(load32(m.memory[int64(uint32(v0))+440:]))
						if t72 != 0 {
							goto l11
						}
						_ = m.fn656(v6, v4)
					}
				l11:
					store32(m.memory[int64(uint32(v3))+932:], uint32(v6))
					store32(m.memory[int64(uint32(v3))+928:], uint32(v3+i32(96)))
					t74 := int32(load32(m.memory[int64(uint32(v0))+432:]))
					t75 := int32(load32(m.memory[int64(uint32(v0))+436:]))
					m.fn69(v3, t74, t75, v14, v3+i32(928), i32(185))
					t76 := int32(load32(m.memory[int64(uint32(v0))+432:]))
					v2 = t76
					t77 := int32(load32(m.memory[int64(uint32(v3))+4:]))
					v10 = t77
					{
						{
							t78 := int32(load32(m.memory[uint32(v3):]))
							if t78 != i32(1) {
								goto l12
							}
							t79 := int32(load32(m.memory[int64(uint32(v3))+508:]))
							store32(m.memory[int64(uint32(v3))+936:], uint32(t79))
							t80 := int64(load64(m.memory[int64(uint32(v3))+500:]))
							store64(m.memory[int64(uint32(v3))+928:], uint64(t80))
							memory_copy(m.memory, uint32(v9), uint32(v3+i32(512)), uint32(i32(400)))
							v11 = v2 + v10
							t81 := int32(m.memory[uint32(v11)])
							v12 = t81
							t82 := v11
							v13 = int32(uint32(int32(v14)) >> 25)
							m.memory[uint32(t82)] = byte(v13)
							t83 := int32(load32(m.memory[int64(uint32(v0))+444:]))
							store32(m.memory[int64(uint32(v0))+444:], uint32(t83+i32(1)))
							t84 := int32(load32(m.memory[int64(uint32(v0))+440:]))
							store32(m.memory[int64(uint32(v0))+440:], uint32(t84-v12&i32(1)))
							t85 := int32(load32(m.memory[int64(uint32(v0))+436:]))
							m.memory[uint32(v2+t85&(v10+i32(-8))+i32(8))] = byte(v13)
							memory_copy(m.memory, uint32(v2+(i32(0)-v10)*i32(416)+i32(-416)), uint32(v3+i32(928)), uint32(i32(416)))
							store32(m.memory[int64(uint32(v3))+104:], uint32(i32(-1)))
							goto l13
						}
					l12:
						t86 := v3 + i32(96)
						v2 = v2 + (i32(0)-v10)*i32(416) + i32(-400)
						memory_copy(m.memory, uint32(t86), uint32(v2), uint32(i32(400)))
						memory_copy(m.memory, uint32(v2), uint32(v3+i32(512)), uint32(i32(400)))
						t87 := int32(load32(m.memory[int64(uint32(v3))+500:]))
						t88 := int32(load32(m.memory[int64(uint32(v3))+504:]))
						m.fn16(t87, t88)
					}
				l13:
					m.fn1330(v3 + i32(96))
				}
			l10:
				t89 := m.fn847(v1, i32(1074680), i32(46), i32(1077406), i32(13))
				if t89 == 0 {
					goto l14
				}
				t90 := int32(load32(m.memory[uint32(v1+i32(28)):]))
				t91 := int32(load32(m.memory[uint32(v1+i32(32)):]))
				m.fn1329(v3+i32(928), t90, t91)
				m.fn1330(v0)
				memory_copy(m.memory, uint32(v0), uint32(v3+i32(928)), uint32(i32(400)))
				goto l14
			}
		}
	}
l0:
	m.g0 = v3 + i32(1344)
}
func (m *Module) fn1306(v0, v1, v2 int32) {
	var v3, v4 int32
	t0 := m.g0
	v3 = t0 - i32(64)
	m.g0 = v3
	store32(m.memory[int64(uint32(v3))+12:], uint32(i32(0)))
	store64(m.memory[int64(uint32(v3))+4:], uint64(i64(0x800000000)))
	store32(m.memory[int64(uint32(v3))+16:], uint32(i32(0)))
	t1 := int32(load32(m.memory[int64(uint32(v1))+32:]))
	v4 = t1
	t2 := int32(load32(m.memory[int64(uint32(v1))+28:]))
	t3 := v3
	v1 = t2
	store32(m.memory[int64(uint32(t3))+32:], uint32(v1))
	store32(m.memory[int64(uint32(v3))+36:], uint32(v1+v4*i32(44)))
	{
		{
		l1:
			{
				t4 := m.fn904(v3 + i32(32))
				v1 = t4
				if v1 == 0 {
					goto l0
				}
				m.fn1331(v3+i32(40), v1, v2, v3+i32(4), v3+i32(16))
				t5 := int32(load32(m.memory[int64(uint32(v3))+40:]))
				v1 = t5
				if v1 == i32(-1) {
					goto l1
				}
			}
			t6 := int32(load32(m.memory[int64(uint32(v3))+60:]))
			store32(m.memory[int64(uint32(v0))+20:], uint32(t6))
			t7 := int64(load64(m.memory[int64(uint32(v3))+52:]))
			store64(m.memory[int64(uint32(v0))+12:], uint64(t7))
			t8 := int64(load64(m.memory[int64(uint32(v3))+44:]))
			store64(m.memory[int64(uint32(v0))+4:], uint64(t8))
			store32(m.memory[uint32(v0):], uint32(v1))
			m.fn1332(v3 + i32(16))
			m.fn969(v3 + i32(4))
			goto l2
		}
	l0:
		m.fn1333(v3+i32(16), v3+i32(4))
		store32(m.memory[uint32(v0):], uint32(i32(-1)))
		t9 := int32(load32(m.memory[int64(uint32(v3))+12:]))
		store32(m.memory[int64(uint32(v0))+12:], uint32(t9))
		t10 := int64(load64(m.memory[int64(uint32(v3))+4:]))
		store64(m.memory[int64(uint32(v0))+4:], uint64(t10))
		m.fn1332(v3 + i32(16))
	}
l2:
	m.g0 = v3 + i32(64)
}
func (m *Module) fn1307(v0, v1, v2 int32) {
	var v3, v4 int32
	t0 := m.g0
	v3 = t0 - i32(176)
	m.g0 = v3
	v4 = v3 + i32(32)
	m.fn1165(v4)
	store64(m.memory[int64(uint32(v3))+8:], uint64(i64(0)))
	store64(m.memory[int64(uint32(v3))+88:], uint64(i64(0)))
	store64(m.memory[int64(uint32(v3))+16:], uint64(i64(0)))
	store64(m.memory[int64(uint32(v3))+24:], uint64(i64(0)))
	m.fn1335(v3+i32(120), v1, v2, v3+i32(8), i32(1))
	{
		t1 := int32(load32(m.memory[int64(uint32(v3))+120:]))
		v2 = t1
		if v2 == i32(-1) {
			goto l0
		}
		t2 := int32(load32(m.memory[int64(uint32(v3))+140:]))
		store32(m.memory[int64(uint32(v0))+20:], uint32(t2))
		t3 := int64(load64(m.memory[int64(uint32(v3))+132:]))
		store64(m.memory[int64(uint32(v0))+12:], uint64(t3))
		t4 := int64(load64(m.memory[int64(uint32(v3))+124:]))
		store64(m.memory[int64(uint32(v0))+4:], uint64(t4))
		store32(m.memory[uint32(v0):], uint32(v2))
		m.fn1259(v4)
		goto l1
	}
l0:
	memory_copy(m.memory, uint32(v3+i32(120)), uint32(v4), uint32(i32(56)))
	m.fn1168(v3+i32(100), v3+i32(120))
	{
		t5 := int32(load32(m.memory[int64(uint32(v3))+108:]))
		v2 = t5
		if v2 == 0 {
			goto l2
		}
		t6 := int32(load32(m.memory[int64(uint32(v3))+104:]))
		t7 := int32(load32(m.memory[int64(uint32(v3))+88:]))
		t8 := m.fn1234(t6, v2, t7)
		store32(m.memory[int64(uint32(v3))+112:], uint32(t8))
		t9 := m.fn113(i32(8), i32(32))
		v2 = t9
		store32(m.memory[uint32(v2):], uint32(i32(-0x7ffffffe)))
		store32(m.memory[int64(uint32(v0))+12:], uint32(i32(1)))
		store32(m.memory[int64(uint32(v0))+8:], uint32(v2))
		store64(m.memory[uint32(v0):], uint64(i64(0x1ffffffff)))
		t10 := int64(load64(m.memory[int64(uint32(v3))+100:]))
		store64(m.memory[int64(uint32(v2))+4:], uint64(t10))
		t11 := int64(load64(m.memory[int64(uint32(v3))+108:]))
		store64(m.memory[int64(uint32(v2))+12:], uint64(t11))
		t12 := int32(load32(m.memory[int64(uint32(v3))+116:]))
		store32(m.memory[int64(uint32(v2))+20:], uint32(t12))
		goto l1
	}
l2:
	store64(m.memory[int64(uint32(v0))+8:], uint64(i64(8)))
	store64(m.memory[uint32(v0):], uint64(i64(0xffffffff)))
	m.fn972(v3 + i32(100))
l1:
	m.g0 = v3 + i32(176)
}
func (m *Module) fn1308(v0, v1 int32) {
	m.fn44(v1, v0)
}
func (m *Module) fn1309(v0, v1, v2, v3, v4, v5 int32) {
	var v6, v7, v8, v9, v10, v11, v12, v13, v14 int32
	var v15 int64
	t0 := m.g0
	v6 = t0 - i32(192)
	m.g0 = v6
	t1 := int32(load32(m.memory[int64(uint32(v1))+32:]))
	v7 = t1
	t2 := int32(load32(m.memory[int64(uint32(v1))+28:]))
	t3 := v6
	v1 = t2
	store32(m.memory[int64(uint32(t3))+16:], uint32(v1))
	store32(m.memory[int64(uint32(v6))+20:], uint32(v1+v7*i32(44)))
	v7 = v6 + i32(132) + i32(4)
	v8 = v6 + i32(160) + i32(12)
	v9 = v6 + i32(160) | i32(4)
	v10 = v6 + i32(160) + i32(4)
l2:
	{
		{
			t4 := m.fn904(v6 + i32(16))
			v1 = t4
			if v1 == 0 {
				store32(m.memory[uint32(v0):], uint32(i32(-1)))
				goto l7
			}
			t5 := m.fn847(v1, i32(1081681), i32(54), i32(1081735), i32(5))
			if t5 != 0 {
				t65 := int32(load32(m.memory[uint32(v1+i32(28)):]))
				t66 := int32(load32(m.memory[uint32(v1+i32(32)):]))
				m.fn868(v6+i32(132), t65, t66)
				store32(m.memory[int64(uint32(v6))+152:], uint32(i32(1081740)))
				store32(m.memory[int64(uint32(v6))+148:], uint32(i32(49)))
				store32(m.memory[int64(uint32(v6))+144:], uint32(i32(1074120)))
				store32(m.memory[int64(uint32(v6))+156:], uint32(i32(5)))
				store32(m.memory[int64(uint32(v6))+184:], uint32(i32(5)))
				t67 := int64(load64(m.memory[int64(uint32(v6))+148:]))
				store64(m.memory[int64(uint32(v6))+176:], uint64(t67))
				t68 := int64(load64(m.memory[int64(uint32(v6))+140:]))
				store64(m.memory[int64(uint32(v6))+168:], uint64(t68))
				t69 := int64(load64(m.memory[int64(uint32(v6))+132:]))
				store64(m.memory[int64(uint32(v6))+160:], uint64(t69))
			l19:
				{
					t70 := m.fn863(v6 + i32(160))
					v1 = t70
					if v1 == 0 {
						t82 := int32(load32(m.memory[int64(uint32(v6))+160:]))
						t83 := int32(load32(m.memory[int64(uint32(v6))+164:]))
						m.fn44(t82, t83)
						goto l2
					}
					t71 := int32(load32(m.memory[uint32(v1+i32(28)):]))
					t72 := int32(load32(m.memory[uint32(v1+i32(32)):]))
					t73 := m.fn886(t71, t72, i32(1074120), i32(49), i32(1081813), i32(8))
					v1 = t73
					if v1 == 0 {
						goto l19
					}
					m.fn1306(v6+i32(132), v1, v2)
					t74 := int64(load64(m.memory[uint32(v7):]))
					store64(m.memory[int64(uint32(v6))+24:], uint64(t74))
					t75 := int32(load32(m.memory[int64(uint32(v7))+8:]))
					store32(m.memory[int64(uint32(v6))+32:], uint32(t75))
					{
						t76 := int32(load32(m.memory[int64(uint32(v6))+132:]))
						v1 = t76
						if v1 == i32(-1) {
							m.fn1271(v5, v6+i32(24))
							goto l19
						}
						t77 := int64(load64(m.memory[int64(uint32(v6))+148:]))
						v15 = t77
						t78 := int32(load32(m.memory[int64(uint32(v6))+32:]))
						store32(m.memory[int64(uint32(v0))+12:], uint32(t78))
						t79 := int64(load64(m.memory[int64(uint32(v6))+24:]))
						store64(m.memory[int64(uint32(v0))+4:], uint64(t79))
						store64(m.memory[int64(uint32(v0))+16:], uint64(v15))
						store32(m.memory[uint32(v0):], uint32(v1))
						t80 := int32(load32(m.memory[int64(uint32(v6))+160:]))
						t81 := int32(load32(m.memory[int64(uint32(v6))+164:]))
						m.fn44(t80, t81)
						goto l7
					}
				}
			}
			t6 := int32(load32(m.memory[int64(uint32(v1))+36:]))
			v11 = t6
			if v11 == 0 {
				goto l2
			}
			t7 := int32(load32(m.memory[int64(uint32(v1))+40:]))
			t8 := m.fn1337(v11+i32(8), t7, i32(1074120), i32(49))
			if t8 != 0 {
				goto l2
			}
			t9 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			v11 = t9
			t10 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			t11 := v11
			v12 = t10
			t12 := m.fn15(t11, v12, i32(1081740), i32(5))
			if t12 != 0 {
				t26 := int32(load32(m.memory[uint32(v1+i32(16)):]))
				t27 := int32(load32(m.memory[uint32(v1+i32(20)):]))
				m.fn1046(v6+i32(8), t26, t27, i32(1081681), i32(54), i32(1077139), i32(5))
				t28 := int32(load32(m.memory[int64(uint32(v6))+8:]))
				v11 = t28
				p29 := i32(1)
				if v11 != 0 {
					p29 = v11
				}
				v12 = p29
				t30 := int32(load32(m.memory[int64(uint32(v6))+12:]))
				t32 := v12
				p31 := i32(0)
				if v11 != 0 {
					p31 = t30
				}
				v13 = p31
				t33 := m.fn15(t32, v13, i32(1081793), i32(11))
				if t33 != 0 {
					goto l2
				}
				t34 := m.fn15(v12, v13, i32(1081804), i32(9))
				if t34 != 0 {
					goto l2
				}
				t35 := m.fn15(v12, v13, i32(1078281), i32(6))
				if t35 != 0 {
					goto l2
				}
				t36 := m.fn15(v12, v13, i32(1078275), i32(6))
				if t36 != 0 {
					goto l2
				}
				store32(m.memory[int64(uint32(v6))+44:], uint32(i32(0)))
				store64(m.memory[int64(uint32(v6))+36:], uint64(i64(0x800000000)))
				t37 := int32(load32(m.memory[int64(uint32(v1))+32:]))
				v11 = t37
				t38 := int32(load32(m.memory[int64(uint32(v1))+28:]))
				t39 := v6
				v14 = t38
				store32(m.memory[int64(uint32(t39))+92:], uint32(v14))
				store32(m.memory[int64(uint32(v6))+96:], uint32(v14+v11*i32(44)))
			l11:
				{
					t40 := m.fn904(v6 + i32(92))
					v11 = t40
					if v11 == 0 {
						goto l8
					}
					{
						t41 := m.fn847(v11, i32(1074120), i32(49), i32(1081813), i32(8))
						if t41 != 0 {
							m.fn1306(v6+i32(160), v11, v2)
							t59 := int64(load64(m.memory[uint32(v10):]))
							store64(m.memory[int64(uint32(v6))+48:], uint64(t59))
							t60 := int32(load32(m.memory[int64(uint32(v10))+8:]))
							store32(m.memory[int64(uint32(v6))+56:], uint32(t60))
							{
								t61 := int32(load32(m.memory[int64(uint32(v6))+160:]))
								v11 = t61
								if v11 == i32(-1) {
									v11 = v6 + i32(48)
									goto l16
								}
								t62 := int64(load64(m.memory[int64(uint32(v6))+176:]))
								v15 = t62
								t63 := int32(load32(m.memory[int64(uint32(v6))+56:]))
								store32(m.memory[int64(uint32(v0))+12:], uint32(t63))
								t64 := int64(load64(m.memory[int64(uint32(v6))+48:]))
								store64(m.memory[int64(uint32(v0))+4:], uint64(t64))
								store64(m.memory[int64(uint32(v0))+16:], uint64(v15))
								store32(m.memory[uint32(v0):], uint32(v11))
								goto l13
							}
						}
						{
							t42 := m.fn847(v11, i32(1074726), i32(47), i32(1074842), i32(5))
							if t42 != 0 {
								m.fn1307(v6+i32(160), v11, v2)
								t53 := int64(load64(m.memory[uint32(v10):]))
								store64(m.memory[int64(uint32(v6))+64:], uint64(t53))
								t54 := int32(load32(m.memory[int64(uint32(v10))+8:]))
								store32(m.memory[int64(uint32(v6))+72:], uint32(t54))
								{
									t55 := int32(load32(m.memory[int64(uint32(v6))+160:]))
									v11 = t55
									if v11 == i32(-1) {
										v11 = v6 + i32(64)
										goto l16
									}
									t56 := int64(load64(m.memory[int64(uint32(v6))+176:]))
									v15 = t56
									t57 := int32(load32(m.memory[int64(uint32(v6))+72:]))
									store32(m.memory[int64(uint32(v0))+12:], uint32(t57))
									t58 := int64(load64(m.memory[int64(uint32(v6))+64:]))
									store64(m.memory[int64(uint32(v0))+4:], uint64(t58))
									store64(m.memory[int64(uint32(v0))+16:], uint64(v15))
									store32(m.memory[uint32(v0):], uint32(v11))
									goto l13
								}
							}
							t43 := m.fn847(v11, i32(1074120), i32(49), i32(1077128), i32(5))
							if t43 == 0 {
								goto l11
							}
							store32(m.memory[int64(uint32(v6))+112:], uint32(i32(0)))
							store64(m.memory[int64(uint32(v6))+104:], uint64(i64(0x400000000)))
							store32(m.memory[int64(uint32(v6))+140:], uint32(i32(0)))
							store64(m.memory[int64(uint32(v6))+132:], uint64(i64(0x800000000)))
							m.fn1338(v6+i32(160), v1, v2, v6+i32(104), v6+i32(132))
							{
								t44 := int32(load32(m.memory[int64(uint32(v6))+160:]))
								v1 = t44
								if v1 == i32(-1) {
									t48 := int32(load32(m.memory[int64(uint32(v6))+108:]))
									t49 := int32(load32(m.memory[int64(uint32(v6))+112:]))
									t50 := m.fn23(t48, t49)
									if t50 != 0 {
										m.fn1271(v6+i32(36), v6+i32(132))
										m.fn894(v6 + i32(104))
										goto l8
									}
									t51 := int32(load32(m.memory[int64(uint32(v6))+112:]))
									store32(m.memory[int64(uint32(v9))+8:], uint32(t51))
									t52 := int64(load64(m.memory[int64(uint32(v6))+104:]))
									store64(m.memory[uint32(v9):], uint64(t52))
									store32(m.memory[int64(uint32(v6))+160:], uint32(i32(-0x80000000)))
									m.fn338(v6+i32(36), v6+i32(160))
									m.fn1271(v6+i32(36), v6+i32(132))
									goto l8
								}
								t45 := int32(load32(m.memory[int64(uint32(v6))+180:]))
								store32(m.memory[int64(uint32(v0))+20:], uint32(t45))
								t46 := int64(load64(m.memory[int64(uint32(v6))+172:]))
								store64(m.memory[int64(uint32(v0))+12:], uint64(t46))
								t47 := int64(load64(m.memory[int64(uint32(v6))+164:]))
								store64(m.memory[int64(uint32(v0))+4:], uint64(t47))
								store32(m.memory[uint32(v0):], uint32(v1))
								m.fn969(v6 + i32(132))
								m.fn894(v6 + i32(104))
								goto l13
							}
						}
					}
				l16:
					m.fn1271(v6+i32(36), v11)
					goto l11
				}
			}
			t13 := m.fn15(v11, v12, i32(1081745), i32(1))
			if t13 != 0 {
				m.fn1309(v6+i32(160), v1, v2, v3, v4, v5)
				t22 := int32(load32(m.memory[int64(uint32(v6))+160:]))
				v1 = t22
				if v1 == i32(-1) {
					goto l2
				}
				t23 := int32(load32(m.memory[int64(uint32(v6))+180:]))
				store32(m.memory[int64(uint32(v0))+20:], uint32(t23))
				t24 := int64(load64(m.memory[int64(uint32(v6))+172:]))
				store64(m.memory[int64(uint32(v0))+12:], uint64(t24))
				t25 := int64(load64(m.memory[int64(uint32(v6))+164:]))
				store64(m.memory[int64(uint32(v0))+4:], uint64(t25))
				store32(m.memory[uint32(v0):], uint32(v1))
				goto l7
			}
			t14 := m.fn15(v11, v12, i32(1081746), i32(12))
			if t14 != 0 {
				goto l5
			}
			t15 := m.fn15(v11, v12, i32(1081758), i32(4))
			if t15 == 0 {
				t16 := m.fn15(v11, v12, i32(1081762), i32(7))
				if t16 != 0 {
					goto l5
				}
				t17 := m.fn15(v11, v12, i32(1081769), i32(7))
				if t17 != 0 {
					goto l5
				}
				t18 := m.fn15(v11, v12, i32(1081776), i32(4))
				if t18 != 0 {
					goto l5
				}
				t19 := m.fn15(v11, v12, i32(1079233), i32(4))
				if t19 != 0 {
					goto l5
				}
				t20 := m.fn15(v11, v12, i32(1081780), i32(9))
				if t20 != 0 {
					goto l5
				}
				t21 := m.fn15(v11, v12, i32(1073494), i32(7))
				if t21 == 0 {
					goto l2
				}
				goto l5
			}
			goto l5
		}
	l8:
		{
			t84 := m.fn15(v12, v13, i32(1073751), i32(5))
			if t84 != 0 {
				t85 := int32(load32(m.memory[int64(uint32(v6))+36:]))
				v1 = t85
				t86 := int32(load32(m.memory[int64(uint32(v6))+40:]))
				v11 = t86
				t87 := int32(load32(m.memory[int64(uint32(v6))+44:]))
				v12 = t87
				store32(m.memory[int64(uint32(v6))+100:], uint32(i32(0)))
				store64(m.memory[int64(uint32(v6))+92:], uint64(i64(0x400000000)))
				t88 := v6
				t89 := v11
				v12 = v12 << 5
				v14 = t89 + v12
				store32(m.memory[int64(uint32(t88))+116:], uint32(v14))
				store32(m.memory[int64(uint32(v6))+112:], uint32(v1))
				store32(m.memory[int64(uint32(v6))+104:], uint32(v11))
			l29:
				{
					if v12 == 0 {
						goto l22
					}
					{
						t90 := int32(load32(m.memory[uint32(v11):]))
						v13 = t90
						if v13 == i32(-1) {
							goto l23
						}
						t91 := v9
						v1 = v11 + i32(4)
						t92 := int64(load64(m.memory[uint32(v1):]))
						store64(m.memory[uint32(t91):], uint64(t92))
						t93 := int64(load64(m.memory[int64(uint32(v1))+8:]))
						store64(m.memory[int64(uint32(v9))+8:], uint64(t93))
						t94 := int64(load64(m.memory[int64(uint32(v1))+16:]))
						store64(m.memory[int64(uint32(v9))+16:], uint64(t94))
						t95 := int32(load32(m.memory[int64(uint32(v1))+24:]))
						store32(m.memory[int64(uint32(v9))+24:], uint32(t95))
						store32(m.memory[int64(uint32(v6))+160:], uint32(v13))
						if v13 == i32(-0x80000000) {
							goto l24
						}
						m.fn970(v6 + i32(160))
						goto l25
					}
				l23:
					v14 = v11 + i32(32)
				l22:
					store32(m.memory[int64(uint32(v6))+108:], uint32(v14))
					m.fn1339(v6 + i32(104))
					{
						t96 := int32(load32(m.memory[int64(uint32(v6))+96:]))
						v1 = t96
						t97 := int32(load32(m.memory[int64(uint32(v6))+100:]))
						t98 := v1
						v11 = t97
						t99 := m.fn23(t98, v11)
						if t99 != 0 {
							m.fn894(v6 + i32(92))
							goto l2
						}
						m.fn45(v8, v1, v11)
						m.memory[int64(uint32(v6))+184] = byte(i32(2))
						t100 := int32(load32(m.memory[int64(uint32(v6))+100:]))
						store32(m.memory[int64(uint32(v6))+168:], uint32(t100))
						t101 := int64(load64(m.memory[int64(uint32(v6))+92:]))
						store64(m.memory[int64(uint32(v6))+160:], uint64(t101))
						m.fn338(v3, v6+i32(160))
						goto l2
					}
				l24:
					t102 := int64(load64(m.memory[uint32(v1):]))
					store64(m.memory[int64(uint32(v6))+120:], uint64(t102))
					t103 := int32(load32(m.memory[int64(uint32(v1))+8:]))
					t104 := v6
					v1 = t103
					store32(m.memory[int64(uint32(t104))+128:], uint32(v1))
					{
						t105 := int32(load32(m.memory[int64(uint32(v6))+124:]))
						t106 := m.fn23(t105, v1)
						if t106 != 0 {
							goto l27
						}
						{
							t107 := int32(load32(m.memory[int64(uint32(v6))+100:]))
							if t107 == 0 {
								goto l28
							}
							store32(m.memory[int64(uint32(v6))+132:], uint32(i32(8)))
							m.fn1340(v6+i32(92), v6+i32(132))
						}
					l28:
						m.fn1341(v6+i32(92), v6+i32(120))
						goto l25
					}
				l27:
					m.fn894(v6 + i32(120))
				}
			l25:
				v11 = v11 + i32(32)
				v12 = v12 + i32(-32)
				goto l29
			}
			m.fn1310(v4, v6+i32(36))
			m.fn969(v6 + i32(36))
			goto l2
		}
	l13:
		m.fn969(v6 + i32(36))
		goto l7
	l5:
		t108 := int32(load32(m.memory[int64(uint32(v1))+32:]))
		v11 = t108
		t109 := int32(load32(m.memory[int64(uint32(v1))+28:]))
		t110 := v6
		v12 = t109
		store32(m.memory[int64(uint32(t110))+132:], uint32(v12))
		store32(m.memory[int64(uint32(v6))+136:], uint32(v12+v11*i32(44)))
	l31:
		{
			t111 := m.fn904(v6 + i32(132))
			v11 = t111
			if v11 == 0 {
				goto l2
			}
			{
				t112 := m.fn847(v11, i32(1074680), i32(46), i32(1077161), i32(1))
				if t112 != 0 {
					goto l30
				}
				t113 := m.fn847(v11, i32(1074680), i32(46), i32(1081789), i32(4))
				if t113 == 0 {
					goto l31
				}
			}
		l30:
		}
		m.fn1306(v6+i32(160), v1, v2)
		t114 := int64(load64(m.memory[uint32(v10):]))
		store64(m.memory[int64(uint32(v6))+80:], uint64(t114))
		t115 := int32(load32(m.memory[int64(uint32(v10))+8:]))
		store32(m.memory[int64(uint32(v6))+88:], uint32(t115))
		t116 := int32(load32(m.memory[int64(uint32(v6))+160:]))
		v1 = t116
		if v1 == i32(-1) {
			m.fn1271(v4, v6+i32(80))
			goto l2
		}
		t117 := int64(load64(m.memory[int64(uint32(v6))+176:]))
		v15 = t117
		t118 := int32(load32(m.memory[int64(uint32(v6))+88:]))
		store32(m.memory[int64(uint32(v0))+12:], uint32(t118))
		t119 := int64(load64(m.memory[int64(uint32(v6))+80:]))
		store64(m.memory[int64(uint32(v0))+4:], uint64(t119))
		store64(m.memory[int64(uint32(v0))+16:], uint64(v15))
		store32(m.memory[uint32(v0):], uint32(v1))
	}
l7:
	m.g0 = v6 + i32(192)
}
func (m *Module) fn1310(v0, v1 int32) {
	t0 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	t1 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	m.fn1336(v0, t0, t1)
	store32(m.memory[int64(uint32(v1))+8:], uint32(i32(0)))
}
func (m *Module) fn1311(v0 int32) {
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
		t8 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		t9 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t10 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		m.fn40(v3-t8, t9, t10)
	}
l0:
	m.g0 = v1 + i32(16)
}
func (m *Module) fn1312(v0 int32) {
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
		t8 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		t9 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t10 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		m.fn40(v3-t8, t9, t10)
	}
l0:
	m.g0 = v1 + i32(16)
}
func (m *Module) fn1313(v0 int32) {
	var v1, v2, v3, v4, v5 int32
	var v6 int64
	var v7, v8 int32
	t0 := m.g0
	v1 = t0 - i32(16)
	m.g0 = v1
	{
		t1 := int32(load32(m.memory[int64(uint32(v0))+404:]))
		v2 = t1
		if v2 == 0 {
			goto l0
		}
		t2 := int32(load32(m.memory[int64(uint32(v0))+400:]))
		v3 = t2
		{
			t3 := int32(load32(m.memory[int64(uint32(v0))+412:]))
			v4 = t3
			if v4 == 0 {
				goto l1
			}
			v5 = v3 + i32(8)
			t4 := int64(load64(m.memory[uint32(v3):]))
			v6 = (t4 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
			v7 = v3
		l4:
			if v4 == 0 {
				goto l1
			}
		l3:
			{
				if v6 != i64(0) {
					v8 = v7 + (i32(0)-int32(uint32(int64(bits.TrailingZeros64(uint64(v6))))>>3))*i32(28)
					t6 := int32(load32(m.memory[uint32(v8+i32(-28)):]))
					t7 := int32(load32(m.memory[uint32(v8+i32(-24)):]))
					m.fn16(t6, t7)
					t8 := int32(load32(m.memory[uint32(v8+i32(-12)):]))
					t9 := int32(load32(m.memory[uint32(v8+i32(-8)):]))
					m.fn134(t8, t9)
					v4 = v4 + i32(-1)
					v6 = (v6 + i64(-1)) & v6
					goto l4
				}
				v7 = v7 + i32(-224)
				t5 := int64(load64(m.memory[uint32(v5):]))
				v6 = (t5 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
				v5 = v5 + i32(8)
				goto l3
			}
		}
	l1:
		m.fn39(v1+i32(4), i32(28), i32(8), v2+i32(1))
		t10 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		t11 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t12 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		m.fn40(v3-t10, t11, t12)
	}
l0:
	m.fn1334(v0 + i32(504))
	{
		t13 := int32(load32(m.memory[int64(uint32(v0))+436:]))
		v2 = t13
		if v2 == 0 {
			goto l5
		}
		t14 := int32(load32(m.memory[int64(uint32(v0))+432:]))
		v3 = t14
		{
			t15 := int32(load32(m.memory[int64(uint32(v0))+444:]))
			v8 = t15
			if v8 == 0 {
				goto l6
			}
			v5 = v3 + i32(8)
			t16 := int64(load64(m.memory[uint32(v3):]))
			v6 = (t16 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
			v7 = v3
		l9:
			if v8 == 0 {
				goto l6
			}
		l8:
			{
				if v6 != i64(0) {
					v4 = v7 + (i32(0)-int32(uint32(int64(bits.TrailingZeros64(uint64(v6))))>>3))*i32(416)
					t18 := int32(load32(m.memory[uint32(v4+i32(-416)):]))
					t19 := int32(load32(m.memory[uint32(v4+i32(-412)):]))
					m.fn16(t18, t19)
					m.fn752(v4 + i32(-400))
					v8 = v8 + i32(-1)
					v6 = (v6 + i64(-1)) & v6
					goto l9
				}
				v7 = v7 + i32(-3328)
				t17 := int64(load64(m.memory[uint32(v5):]))
				v6 = (t17 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
				v5 = v5 + i32(8)
				goto l8
			}
		}
	l6:
		m.fn39(v1+i32(4), i32(416), i32(8), v2+i32(1))
		t20 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		t21 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t22 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		m.fn40(v3-t20, t21, t22)
	}
l5:
	m.fn1330(v0)
	m.fn1334(v0 + i32(464))
	m.g0 = v1 + i32(16)
}
func (m *Module) fn1314(v0, v1, v2 int32) {
	var v3 int32
	t0 := m.g0
	v3 = t0 - i32(64)
	m.g0 = v3
	store32(m.memory[int64(uint32(v3))+8:], uint32(i32(0)))
	store64(m.memory[uint32(v3):], uint64(i64(0x800000000)))
	store32(m.memory[int64(uint32(v3))+36:], uint32(i32(0)))
	store64(m.memory[int64(uint32(v3))+28:], uint64(i64(0x800000000)))
	store32(m.memory[int64(uint32(v3))+12:], uint32(i32(0)))
	m.fn1323(v3+i32(40), v1, v2, v3, v3+i32(12))
	{
		{
			t1 := int32(load32(m.memory[int64(uint32(v3))+40:]))
			v2 = t1
			if v2 == i32(-1) {
				goto l0
			}
			t2 := int32(load32(m.memory[int64(uint32(v3))+60:]))
			store32(m.memory[int64(uint32(v0))+20:], uint32(t2))
			t3 := int64(load64(m.memory[int64(uint32(v3))+52:]))
			store64(m.memory[int64(uint32(v0))+12:], uint64(t3))
			t4 := int64(load64(m.memory[int64(uint32(v3))+44:]))
			store64(m.memory[int64(uint32(v0))+4:], uint64(t4))
			m.fn1324(v3 + i32(12))
			m.fn969(v3)
			goto l1
		}
	l0:
		m.fn1325(v3+i32(12), v3)
		t5 := int32(load32(m.memory[int64(uint32(v3))+8:]))
		store32(m.memory[int64(uint32(v0))+12:], uint32(t5))
		t6 := int64(load64(m.memory[uint32(v3):]))
		store64(m.memory[int64(uint32(v0))+4:], uint64(t6))
		m.fn1324(v3 + i32(12))
	}
l1:
	store32(m.memory[uint32(v0):], uint32(v2))
	m.g0 = v3 + i32(64)
}
func (m *Module) fn1315(v0 int32) {
	m.fn1043(v0)
	t0 := int32(load32(m.memory[int64(uint32(v0))+52:]))
	t1 := int32(load32(m.memory[int64(uint32(v0))+56:]))
	m.fn16(t0, t1)
}
func (m *Module) fn1316(v0 int32) {
	var v1, v2 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		v1 = t0
		t1 := int32(load32(m.memory[uint32(v0):]))
		t2 := v1
		v2 = t1
		if t2 == v2 {
			return
		}
		v1 = v1 - v2
		v0 = v2*i32(40) + v0 + i32(12)
	l1:
		{
			if v1 == 0 {
				return
			}
			t3 := int32(load32(m.memory[uint32(v0+i32(-4)):]))
			t4 := int32(load32(m.memory[uint32(v0):]))
			m.fn16(t3, t4)
			v1 = v1 + i32(-1)
			v0 = v0 + i32(40)
			goto l1
		}
	}
}
func (m *Module) fn1317(v0, v1 int32) {
	var v2 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	{
		if v1 == 0 {
			goto l0
		}
		m.fn39(v2+i32(4), i32(104), i32(8), v1+i32(1))
		t1 := int32(load32(m.memory[int64(uint32(v2))+12:]))
		t2 := int32(load32(m.memory[int64(uint32(v2))+4:]))
		t3 := int32(load32(m.memory[int64(uint32(v2))+8:]))
		m.fn40(v0-t1, t2, t3)
	}
l0:
	m.g0 = v2 + i32(16)
}
func (m *Module) fn1318(v0, v1, v2, v3 int32) int32 {
	var v4 int32
	t0 := m.g0
	v4 = t0 - i32(16)
	m.g0 = v4
	{
		{
			t1 := m.fn886(v0, v1, i32(1072544), i32(60), v2, v3)
			v3 = t1
			if v3 != 0 {
				goto l0
			}
			v3 = i32(2)
			goto l1
		}
	l0:
		t2 := int32(load32(m.memory[uint32(v3+i32(16)):]))
		t3 := int32(load32(m.memory[uint32(v3+i32(20)):]))
		m.fn1046(v4+i32(8), t2, t3, i32(1072544), i32(60), i32(1073156), i32(3))
		{
			t4 := int32(load32(m.memory[int64(uint32(v4))+8:]))
			v2 = t4
			if v2 != 0 {
				goto l2
			}
			v3 = i32(1)
			goto l1
		}
	l2:
		v3 = i32(0)
		t5 := int32(load32(m.memory[int64(uint32(v4))+12:]))
		t6 := v2
		v1 = t5
		t7 := m.fn15(t6, v1, i32(1108008), i32(1))
		if t7 != 0 {
			goto l1
		}
		t8 := m.fn15(v2, v1, i32(1081456), i32(5))
		if t8 != 0 {
			goto l1
		}
		t9 := m.fn15(v2, v1, i32(1086043), i32(3))
		if t9 != 0 {
			goto l1
		}
		t10 := m.fn15(v2, v1, i32(1074851), i32(4))
		v3 = t10 ^ i32(1)
	}
l1:
	m.g0 = v4 + i32(16)
	return v3
}
func fn1319(v0, v1 int32) int32 {
	p0 := v0
	if v0&i32(0xff00) == i32(512) {
		p0 = v1
	}
	t2 := p0 & i32(256)
	p1 := v0
	if v0&i32(0xff0000) == i32(0x20000) {
		p1 = v1
	}
	t4 := t2 | p1&i32(65536)
	p3 := v0
	if v0&i32(255) == i32(2) {
		p3 = v1
	}
	t5 := t4 | p3&i32(1)
	t6 := int32(uint32(v1) >> 24)
	v0 = int32(uint32(v0) >> 24)
	p7 := v0
	if v0 == i32(2) {
		p7 = t6
	}
	p8 := i32(0)
	if p7&i32(1) != 0 {
		p8 = i32(0x1000000)
	}
	return t5 | p8
}
func (m *Module) fn1320(v0, v1 int32, v2 int64, v3, v4 int32) int32 {
	var v5 int32
	var v6 int64
	var v7 int32
	var v8 int64
	var v9, v10, v11 int32
	v5 = v1 & int32(v2)
	v6 = int64(uint64(v2)>>25) & i64(127) * i64(72340172838076673)
	v7 = i32(0)
	var _ int32
l4:
	{
		t1 := int64(load64(m.memory[uint32(v0+v5):]))
		v8 = t1
		v2 = v8 ^ v6
		v2 = (v2 ^ i64(-1)) & (v2 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
		{
		l2:
			{
				var p2 int32
				if v2 == 0 {
					p2 = 1
				}
				v9 = p2
				if v9 != 0 {
					goto l0
				}
				t3 := v3
				t4 := v4
				t5 := v0
				v10 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v2))))>>3) + v5) & v1
				v11 = t5 + (i32(0)-v10)*i32(20)
				t6 := int32(load32(m.memory[uint32(v11+i32(-20)):]))
				t7 := int32(load32(m.memory[uint32(v11+i32(-16)):]))
				t8 := m.fn15(t3, t4, t6, t7)
				if t8 != 0 {
					goto l1
				}
				v2 = (v2 + i64(-1)) & v2
				goto l2
			}
		l0:
			if v8&(v8<<1)&i64(-0x7f7f7f7f7f7f7f80) == 0 {
				t10 := v5
				v7 = v7 + i32(8)
				v5 = (t10 + v7) & v1
				goto l4
			}
		l1:
			p9 := v0 + (i32(0)-v10)*i32(20)
			if v9 != 0 {
				p9 = i32(0)
			}
			return p9
		}
	}
}
func (m *Module) fn1321(v0, v1 int32) {
	var v2 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		v2 = t0
		t1 := int32(load32(m.memory[uint32(v0):]))
		if v2 != t1 {
			goto l0
		}
		m.fn272(v0)
	}
l0:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v2+i32(1)))
	t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	v0 = t2 + v2*i32(12)
	t3 := int64(load64(m.memory[uint32(v1):]))
	store64(m.memory[uint32(v0):], uint64(t3))
	t4 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	store32(m.memory[int64(uint32(v0))+8:], uint32(t4))
}
func (m *Module) fn1322(v0, v1, v2 int32) {
	var v3, v4 int32
	var v5, v6 int64
	t0 := m.g0
	v3 = t0 - i32(80)
	m.g0 = v3
	switch v2 {
	case 0:
		m.memory[int64(uint32(v0))+1] = byte(i32(0))
		v1 = i32(1)
		goto l3
	case 1:
		t1 := int32(m.memory[uint32(v1)])
		v4 = t1
		switch v4 + i32(-43) {
		case 0, 2:
			goto l4
		default:
			goto l5
		}
	default:
		t2 := int32(m.memory[uint32(v1)])
		v4 = t2
	}
l5:
	switch v4&i32(255) + i32(-43) {
	case 2:
		v4 = v2 + i32(-1)
		v1 = v1 + i32(1)
		if uint32(v2) < uint32(i32(17)) {
			v5 = i64(0)
		l14:
			{
				if v4 == 0 {
					goto l10
				}
				t10 := int32(m.memory[uint32(v1)])
				m.fn199(v3+i32(40), t10, i32(10))
				t11 := int32(load32(m.memory[int64(uint32(v3))+40:]))
				if t11 != i32(1) {
					goto l4
				}
				v1 = v1 + i32(1)
				v4 = v4 + i32(-1)
				t12 := int32(load32(m.memory[int64(uint32(v3))+44:]))
				v5 = v5*i64(10) - int64(uint32(t12))
				goto l14
			}
		}
		v5 = i64(0)
	l13:
		{
			if v4 == 0 {
				goto l10
			}
			m.fn1853(v3+i32(48), v5, v5>>63, i64(10), i64(0))
			t3 := int32(m.memory[uint32(v1)])
			m.fn199(v3+i32(72), t3, i32(10))
			t4 := int32(load32(m.memory[int64(uint32(v3))+72:]))
			v2 = t4
			{
				t5 := int64(load64(m.memory[int64(uint32(v3))+56:]))
				t6 := int64(load64(m.memory[int64(uint32(v3))+48:]))
				v6 = t6
				if t5 != v6>>63 {
					goto l11
				}
				if v2&i32(1) == 0 {
					goto l4
				}
				t7 := int32(load32(m.memory[int64(uint32(v3))+76:]))
				v5 = int64(uint32(t7))
				var p8 int32
				if v5 > i64(0) {
					p8 = 1
				}
				v5 = v6 - v5
				var p9 int32
				if v5 < v6 {
					p9 = 1
				}
				if p8^p9 != 0 {
					m.memory[int64(uint32(v0))+1] = byte(i32(3))
					v1 = i32(1)
					goto l3
				}
				v1 = v1 + i32(1)
				v4 = v4 + i32(-1)
				goto l13
			}
		l11:
		}
		v1 = i32(1)
		if v2&i32(1) == 0 {
			goto l4
		}
		m.memory[int64(uint32(v0))+1] = byte(i32(3))
		goto l3
	case 0:
		v2 = v2 + i32(-1)
		v1 = v1 + i32(1)
		fallthrough
	default:
		if uint32(v2) < uint32(i32(16)) {
			v5 = i64(0)
		l19:
			{
				if v2 == 0 {
					goto l10
				}
				t20 := int32(m.memory[uint32(v1)])
				m.fn199(v3+i32(8), t20, i32(10))
				t21 := int32(load32(m.memory[int64(uint32(v3))+8:]))
				if t21 != i32(1) {
					goto l4
				}
				v1 = v1 + i32(1)
				v2 = v2 + i32(-1)
				t22 := int32(load32(m.memory[int64(uint32(v3))+12:]))
				v5 = v5*i64(10) + int64(uint32(t22))
				goto l19
			}
		}
		v5 = i64(0)
	l18:
		{
			if v2 == 0 {
				goto l10
			}
			m.fn1853(v3+i32(16), v5, v5>>63, i64(10), i64(0))
			t13 := int32(m.memory[uint32(v1)])
			m.fn199(v3+i32(32), t13, i32(10))
			t14 := int32(load32(m.memory[int64(uint32(v3))+32:]))
			v4 = t14
			{
				t15 := int64(load64(m.memory[int64(uint32(v3))+24:]))
				t16 := int64(load64(m.memory[int64(uint32(v3))+16:]))
				v6 = t16
				if t15 != v6>>63 {
					goto l16
				}
				if v4&i32(1) == 0 {
					goto l4
				}
				t17 := int32(load32(m.memory[int64(uint32(v3))+36:]))
				v5 = int64(uint32(t17))
				var p18 int32
				if v5 < i64(0) {
					p18 = 1
				}
				v5 = v6 + v5
				var p19 int32
				if v5 < v6 {
					p19 = 1
				}
				if p18^p19 != 0 {
					m.memory[int64(uint32(v0))+1] = byte(i32(2))
					v1 = i32(1)
					goto l3
				}
				v1 = v1 + i32(1)
				v2 = v2 + i32(-1)
				goto l18
			}
		l16:
		}
		v1 = i32(1)
		if v4&i32(1) == 0 {
			goto l4
		}
		m.memory[int64(uint32(v0))+1] = byte(i32(2))
		goto l3
	}
l4:
	v1 = i32(1)
	m.memory[int64(uint32(v0))+1] = byte(i32(1))
	goto l3
l10:
	store64(m.memory[int64(uint32(v0))+8:], uint64(v5))
	v1 = i32(0)
l3:
	m.memory[uint32(v0)] = byte(v1)
	m.g0 = v3 + i32(80)
}
func (m *Module) fn1323(v0, v1, v2, v3, v4 int32) {
	var v5, v6, v7, v8, v9, v10, v11, v12, v13, v14, v15, v16, v17, v18, v19, v20, v21, v22, v23, v24, v25, v26, v27, v28, v29, v30 int32
	var v31, v32 int64
	var v33, v34 int32
	var v35, v36, v37 int64
	var v38, v39 int32
	var v40 int64
	t0 := m.g0
	v5 = t0 - i32(400)
	m.g0 = v5
	t1 := int32(load32(m.memory[int64(uint32(v1))+32:]))
	v6 = t1
	t2 := int32(load32(m.memory[int64(uint32(v1))+28:]))
	t3 := v5
	v1 = t2
	store32(m.memory[int64(uint32(t3))+120:], uint32(v1))
	store32(m.memory[int64(uint32(v5))+124:], uint32(v1+v6*i32(44)))
	v7 = v4 + i32(16)
	v8 = v5 + i32(336) + i32(16)
	v9 = v5 + i32(336) | i32(4)
	v10 = v5 + i32(248) + i32(4)
	v11 = v5 + i32(376)
	v12 = v5 + i32(192) + i32(28)
	v13 = v5 + i32(152) + i32(28)
	v14 = v5 + i32(336) + i32(8)
	t4 := int32(load32(m.memory[int64(uint32(v2))+36:]))
	v15 = t4
	v16 = v5 + i32(336) + i32(11)
	v17 = v5 + i32(351)
	v18 = v5 + i32(336) + i32(6)
	{
	l2:
		{
			t5 := m.fn904(v5 + i32(120))
			v19 = t5
			if v19 == 0 {
				store32(m.memory[uint32(v0):], uint32(i32(-1)))
				goto l6
			}
			t6 := m.fn847(v19, i32(1073848), i32(59), i32(1077491), i32(16))
			if t6 != 0 {
				t128 := int32(load32(m.memory[uint32(v19+i32(28)):]))
				t129 := int32(load32(m.memory[uint32(v19+i32(32)):]))
				t130 := m.fn1531(t128, t129)
				v1 = t130
				if v1 == 0 {
					goto l2
				}
				m.fn1323(v5+i32(336), v1, v2, v3, v4)
				t131 := int32(load32(m.memory[int64(uint32(v5))+336:]))
				v1 = t131
				if v1 == i32(-1) {
					goto l2
				}
				t132 := int32(load32(m.memory[int64(uint32(v5))+356:]))
				store32(m.memory[int64(uint32(v0))+20:], uint32(t132))
				t133 := int64(load64(m.memory[int64(uint32(v5))+348:]))
				store64(m.memory[int64(uint32(v0))+12:], uint64(t133))
				t134 := int64(load64(m.memory[int64(uint32(v5))+340:]))
				store64(m.memory[int64(uint32(v0))+4:], uint64(t134))
				store32(m.memory[uint32(v0):], uint32(v1))
				goto l6
			}
			t7 := int32(load32(m.memory[int64(uint32(v19))+36:]))
			v1 = t7
			if v1 == 0 {
				goto l2
			}
			t8 := int32(load32(m.memory[int64(uint32(v19))+40:]))
			t9 := m.fn1337(v1+i32(8), t8, i32(1072544), i32(60))
			if t9 != 0 {
				goto l2
			}
			t10 := int32(load32(m.memory[int64(uint32(v19))+4:]))
			v1 = t10
			t11 := int32(load32(m.memory[int64(uint32(v19))+8:]))
			t12 := v1
			v6 = t11
			t13 := m.fn15(t12, v6, i32(1077161), i32(1))
			if t13 != 0 {
				t21 := int32(load32(m.memory[uint32(v19+i32(28)):]))
				t22 := int32(load32(m.memory[uint32(v19+i32(32)):]))
				t23 := m.fn886(t21, t22, i32(1072544), i32(60), i32(1073735), i32(3))
				v20 = t23
				if v20 != 0 {
					v22 = i32(0)
					v21 = i32(0)
					{
						t24 := int32(load32(m.memory[uint32(v20+i32(28)):]))
						v6 = t24
						t25 := int32(load32(m.memory[uint32(v20+i32(32)):]))
						t26 := v6
						v23 = t25
						t27 := m.fn886(t26, v23, i32(1072544), i32(60), i32(1074872), i32(6))
						v1 = t27
						if v1 == 0 {
							goto l9
						}
						t28 := int32(load32(m.memory[uint32(v1+i32(16)):]))
						t29 := int32(load32(m.memory[uint32(v1+i32(20)):]))
						m.fn1046(v5+i32(104), t28, t29, i32(1072544), i32(60), i32(1073156), i32(3))
						t30 := int32(load32(m.memory[int64(uint32(v5))+108:]))
						v24 = t30
						t31 := int32(load32(m.memory[int64(uint32(v5))+104:]))
						v21 = t31
					}
				l9:
					t32 := m.fn886(v6, v23, i32(1072544), i32(60), i32(1073756), i32(10))
					v1 = t32
					if v1 == 0 {
						goto l8
					}
					t33 := int32(load32(m.memory[uint32(v1+i32(16)):]))
					t34 := int32(load32(m.memory[uint32(v1+i32(20)):]))
					m.fn1046(v5+i32(96), t33, t34, i32(1072544), i32(60), i32(1073156), i32(3))
					t35 := int32(load32(m.memory[int64(uint32(v5))+96:]))
					v1 = t35
					if v1 == 0 {
						goto l8
					}
					t36 := int32(load32(m.memory[int64(uint32(v5))+100:]))
					m.fn1423(v5+i32(88), v1, t36)
					t37 := int32(m.memory[int64(uint32(v5))+88])
					v22 = t37 ^ i32(1)
					t38 := int32(m.memory[int64(uint32(v5))+89])
					v25 = t38
					goto l8
				}
				v21 = i32(0)
				v22 = i32(0)
				goto l8
			}
			t14 := m.fn15(v1, v6, i32(1083171), i32(3))
			if t14 != 0 {
				m.fn1325(v4, v3)
				v26 = i32(0)
				store32(m.memory[int64(uint32(v5))+244:], uint32(i32(0)))
				store64(m.memory[int64(uint32(v5))+236:], uint64(i64(0x400000000)))
				t39 := int32(load32(m.memory[int64(uint32(v19))+32:]))
				v6 = t39
				t40 := int32(load32(m.memory[int64(uint32(v19))+28:]))
				v1 = t40
				store32(m.memory[int64(uint32(v5))+356:], uint32(i32(2)))
				store32(m.memory[int64(uint32(v5))+352:], uint32(i32(1073488)))
				store32(m.memory[int64(uint32(v5))+348:], uint32(i32(60)))
				store32(m.memory[int64(uint32(v5))+344:], uint32(i32(1072544)))
				store32(m.memory[int64(uint32(v5))+336:], uint32(v1))
				store32(m.memory[int64(uint32(v5))+340:], uint32(v1+v6*i32(44)))
				v21 = i32(4)
			l17:
				{
					t41 := m.fn1186(v5 + i32(336))
					v27 = t41
					if v27 == 0 {
						m.fn22(v5+i32(336), i32(3))
						t61 := int64(load64(m.memory[int64(uint32(v5))+344:]))
						v31 = t61
						t62 := int64(load64(m.memory[int64(uint32(v5))+336:]))
						v32 = t62
						v33 = i32(1286248)
						v34 = i32(0)
						t63 := int32(load32(m.memory[int64(uint32(v5))+240:]))
						v28 = t63
						v20 = i32(0)
						v23 = i32(0)
					l36:
						{
							if v23 == v26 {
								v24 = i32(0)
								m.memory[int64(uint32(v5))+360] = byte(i32(0))
								store32(m.memory[int64(uint32(v5))+356:], uint32(i32(2)))
								store32(m.memory[int64(uint32(v5))+352:], uint32(i32(1073488)))
								store32(m.memory[int64(uint32(v5))+348:], uint32(i32(60)))
								store32(m.memory[int64(uint32(v5))+344:], uint32(i32(1072544)))
								t142 := int32(load32(m.memory[int64(uint32(v19))+28:]))
								t143 := v5
								v6 = t142
								store32(m.memory[int64(uint32(t143))+336:], uint32(v6))
								t144 := int32(load32(m.memory[int64(uint32(v19))+32:]))
								t145 := v5
								v23 = v6 + t144*i32(44)
								store32(m.memory[int64(uint32(t145))+340:], uint32(v23))
							l38:
								{
									v1 = v6
									if v1 == v23 {
										goto l37
									}
									v6 = v1 + i32(44)
									t146 := int32(load32(m.memory[uint32(v1):]))
									if t146 == i32(-1) {
										goto l38
									}
									t147 := m.fn844(v14, v1)
									if t147 == 0 {
										goto l38
									}
									t148 := int32(load32(m.memory[uint32(v1+i32(28)):]))
									t149 := int32(load32(m.memory[uint32(v1+i32(32)):]))
									t150 := m.fn886(t148, t149, i32(1072544), i32(60), i32(1074868), i32(4))
									v1 = t150
									if v1 == 0 {
										goto l37
									}
									t151 := int32(load32(m.memory[uint32(v1+i32(28)):]))
									t152 := int32(load32(m.memory[uint32(v1+i32(32)):]))
									t153 := m.fn1318(t151, t152, i32(1073309), i32(9))
									if t153&i32(255) != i32(1) {
										goto l37
									}
									v24 = v24 + i32(1)
									goto l38
								}
							l37:
								v30 = v28 + v26*i32(12)
								m.fn1165(v5 + i32(248))
							l41:
								{
									if v28 == v30 {
										goto l39
									}
									v21 = v28 + i32(12)
									m.fn1166(v5 + i32(248))
									t154 := int32(load32(m.memory[int64(uint32(v28))+4:]))
									v29 = t154
									t155 := int32(load32(m.memory[int64(uint32(v28))+8:]))
									v28 = v29 + t155*i32(28)
								l52:
									if v29 != v28 {
										v27 = v29 + i32(28)
										{
											{
												t156 := int32(m.memory[int64(uint32(v29))+24])
												if t156 != 0 {
													t163 := int32(load32(m.memory[int64(uint32(v29))+16:]))
													v1 = t163
												l47:
													if v1 == 0 {
														goto l46
													}
													v1 = v1 + i32(-1)
													_ = m.fn1260(v5 + i32(248))
													goto l47
												}
												{
													t157 := int32(load32(m.memory[int64(uint32(v29))+12:]))
													v1 = t157
													if v1 != 0 {
														m.fn1314(v5+i32(336), v1, v2)
														t158 := int32(load32(m.memory[int64(uint32(v5))+348:]))
														v6 = t158
														t159 := int32(load32(m.memory[int64(uint32(v5))+344:]))
														v1 = t159
														t160 := int32(load32(m.memory[int64(uint32(v5))+340:]))
														v26 = t160
														t161 := int32(load32(m.memory[int64(uint32(v5))+336:]))
														v23 = t161
														if v23 == i32(-1) {
															goto l44
														}
														t162 := int64(load64(m.memory[int64(uint32(v5))+352:]))
														v40 = t162
														goto l45
													}
													v1 = i32(8)
													v26 = i32(0)
													v6 = i32(0)
													goto l44
												}
											}
										l44:
											store32(m.memory[int64(uint32(v5))+312:], uint32(v6))
											store32(m.memory[int64(uint32(v5))+308:], uint32(v1))
											store32(m.memory[int64(uint32(v5))+304:], uint32(v26))
											t165 := int32(load32(m.memory[int64(uint32(v29))+8:]))
											v1 = t165 << 2
											t166 := int32(load32(m.memory[int64(uint32(v29))+4:]))
											v6 = t166
											{
											l50:
												{
													if v1 == 0 {
														goto l48
													}
													t167 := int32(load32(m.memory[uint32(v6):]))
													m.fn1314(v5+i32(336), t167, v2)
													{
														t168 := int32(load32(m.memory[int64(uint32(v5))+336:]))
														v23 = t168
														if v23 != i32(-1) {
															goto l49
														}
														t169 := int32(load32(m.memory[int64(uint32(v5))+348:]))
														store32(m.memory[int64(uint32(v5))+160:], uint32(t169))
														t170 := int64(load64(m.memory[int64(uint32(v5))+340:]))
														store64(m.memory[int64(uint32(v5))+152:], uint64(t170))
														v1 = v1 + i32(-4)
														v6 = v6 + i32(4)
														m.fn1271(v5+i32(304), v5+i32(152))
														goto l50
													}
												l49:
												}
												t171 := int64(load64(m.memory[int64(uint32(v5))+352:]))
												v40 = t171
												t172 := int32(load32(m.memory[int64(uint32(v5))+348:]))
												v6 = t172
												t173 := int32(load32(m.memory[int64(uint32(v5))+344:]))
												v1 = t173
												t174 := int32(load32(m.memory[int64(uint32(v5))+340:]))
												v26 = t174
												m.fn969(v5 + i32(304))
												goto l45
											}
										l48:
											t175 := int32(load32(m.memory[int64(uint32(v5))+312:]))
											store32(m.memory[int64(uint32(v5))+200:], uint32(t175))
											t176 := int64(load64(m.memory[int64(uint32(v5))+304:]))
											store64(m.memory[int64(uint32(v5))+192:], uint64(t176))
											t177 := int32(load32(m.memory[int64(uint32(v29))+20:]))
											t178 := v5
											v1 = t177
											p179 := i32(1)
											if uint32(v1) > uint32(i32(1)) {
												p179 = v1
											}
											store32(m.memory[int64(uint32(t178))+208:], uint32(p179))
											t180 := int32(load32(m.memory[int64(uint32(v29))+16:]))
											t181 := v5
											v1 = t180
											p182 := i32(1)
											if uint32(v1) > uint32(i32(1)) {
												p182 = v1
											}
											store32(m.memory[int64(uint32(t181))+204:], uint32(p182))
											m.fn1167(v5+i32(336), v5+i32(248), v5+i32(192))
											t183 := int32(load32(m.memory[int64(uint32(v5))+336:]))
											v23 = t183
											if v23 == i32(-1) {
												goto l46
											}
											t184 := int64(load64(m.memory[int64(uint32(v5))+352:]))
											v40 = t184
											t185 := int32(load32(m.memory[int64(uint32(v5))+348:]))
											v6 = t185
											t186 := int32(load32(m.memory[int64(uint32(v5))+344:]))
											v1 = t186
											t187 := int32(load32(m.memory[int64(uint32(v5))+340:]))
											v26 = t187
										}
									l45:
										m.fn1259(v5 + i32(248))
										goto l51
									l46:
										v29 = v27
										goto l52
									}
									v28 = v21
									goto l41
								}
							l39:
								memory_copy(m.memory, uint32(v5+i32(336)), uint32(v5+i32(248)), uint32(i32(56)))
								m.fn1168(v5+i32(316), v5+i32(336))
								{
									t188 := int32(load32(m.memory[int64(uint32(v5))+324:]))
									v1 = t188
									if v1 == 0 {
										goto l53
									}
									t189 := int32(load32(m.memory[int64(uint32(v5))+320:]))
									t190 := m.fn1234(t189, v1, v24)
									store32(m.memory[int64(uint32(v5))+328:], uint32(t190))
									t191 := m.fn113(i32(8), i32(32))
									v1 = t191
									store32(m.memory[uint32(v1):], uint32(i32(-0x7ffffffe)))
									t192 := int64(load64(m.memory[int64(uint32(v5))+316:]))
									store64(m.memory[int64(uint32(v1))+4:], uint64(t192))
									t193 := int64(load64(m.memory[int64(uint32(v5))+324:]))
									store64(m.memory[int64(uint32(v1))+12:], uint64(t193))
									t194 := int32(load32(m.memory[int64(uint32(v5))+332:]))
									store32(m.memory[int64(uint32(v1))+20:], uint32(t194))
									m.fn1530(v33, v34)
									m.fn1532(v5 + i32(236))
									v6 = i32(1)
									v26 = i32(1)
									goto l54
								}
							l53:
								m.fn972(v5 + i32(316))
								v23 = i32(-1)
								v1 = i32(8)
								v26 = i32(0)
								v6 = i32(0)
							l51:
								m.fn1530(v33, v34)
								m.fn1532(v5 + i32(236))
								if v23 == i32(-1) {
									goto l54
								}
								store64(m.memory[int64(uint32(v0))+16:], uint64(v40))
								store32(m.memory[int64(uint32(v0))+12:], uint32(v6))
								store32(m.memory[int64(uint32(v0))+8:], uint32(v1))
								store32(m.memory[int64(uint32(v0))+4:], uint32(v26))
								store32(m.memory[uint32(v0):], uint32(v23))
								goto l6
							l54:
								store32(m.memory[int64(uint32(v5))+136:], uint32(v6))
								store32(m.memory[int64(uint32(v5))+132:], uint32(v1))
								store32(m.memory[int64(uint32(v5))+128:], uint32(v26))
								m.fn1271(v3, v5+i32(128))
								goto l2
							}
							v6 = i32(0)
							store32(m.memory[int64(uint32(v5))+200:], uint32(i32(0)))
							store64(m.memory[int64(uint32(v5))+192:], uint64(i64(0x400000000)))
							t64 := m.fn857(v28, v26, v23, i32(1086048))
							t65 := int32(load32(m.memory[int64(uint32(t64))+8:]))
							v24 = t65
							v1 = i32(0)
						l22:
							{
								if v6 == v24 {
									t99 := int32(load32(m.memory[int64(uint32(v5))+196:]))
									v38 = t99
									t100 := int32(load32(m.memory[int64(uint32(v5))+192:]))
									v39 = t100
									t101 := int32(load32(m.memory[int64(uint32(v5))+200:]))
									v1 = t101
									m.fn34(v5 + i32(248))
									t102 := int64(load64(m.memory[int64(uint32(i32(0)))+1286256:]))
									store64(m.memory[int64(uint32(v5))+336:], uint64(t102))
									t103 := int64(load64(m.memory[int64(uint32(i32(0)))+1286264:]))
									store64(m.memory[int64(uint32(v5))+344:], uint64(t103))
									t104 := int64(load64(m.memory[int64(uint32(v5))+256:]))
									store64(m.memory[int64(uint32(v5))+360:], uint64(t104))
									t105 := int64(load64(m.memory[int64(uint32(v5))+248:]))
									store64(m.memory[int64(uint32(v5))+352:], uint64(t105))
									m.fn710(v5+i32(336), v1, v8)
									v6 = v1 * i32(12)
									v1 = v38
								l35:
									{
										if v6 == 0 {
											m.fn136(v39, v38, i32(4), i32(12))
											t123 := int64(load64(m.memory[int64(uint32(v5))+360:]))
											v31 = t123
											t124 := int64(load64(m.memory[int64(uint32(v5))+352:]))
											v32 = t124
											t125 := int32(load32(m.memory[int64(uint32(v5))+348:]))
											v20 = t125
											t126 := int32(load32(m.memory[int64(uint32(v5))+340:]))
											v1 = t126
											t127 := int32(load32(m.memory[int64(uint32(v5))+336:]))
											v6 = t127
											m.fn1530(v33, v34)
											v23 = v23 + i32(1)
											v34 = v1
											v33 = v6
											goto l36
										}
										t106 := int32(load32(m.memory[int64(uint32(v1))+8:]))
										v21 = t106
										t107 := int32(load32(m.memory[int64(uint32(v1))+4:]))
										v24 = t107
										t108 := int32(load32(m.memory[uint32(v1):]))
										t109 := v5
										v30 = t108
										store32(m.memory[int64(uint32(t109))+396:], uint32(v30))
										t110 := int64(load64(m.memory[int64(uint32(v5))+352:]))
										t111 := int64(load64(m.memory[int64(uint32(v5))+360:]))
										t112 := m.fn66(t110, t111, v30)
										v35 = t112
										store32(m.memory[int64(uint32(v5))+304:], uint32(v5+i32(396)))
										m.fn710(v5+i32(336), i32(1), v8)
										store32(m.memory[int64(uint32(v5))+252:], uint32(v5+i32(336)))
										store32(m.memory[int64(uint32(v5))+248:], uint32(v5+i32(304)))
										t113 := int32(load32(m.memory[int64(uint32(v5))+336:]))
										t114 := int32(load32(m.memory[int64(uint32(v5))+340:]))
										m.fn69(v5+i32(112), t113, t114, v35, v5+i32(248), i32(186))
										t115 := int32(load32(m.memory[int64(uint32(v5))+116:]))
										v29 = t115
										t116 := int32(load32(m.memory[int64(uint32(v5))+336:]))
										v27 = t116
										{
											t117 := int32(load32(m.memory[int64(uint32(v5))+112:]))
											if t117 != i32(1) {
												goto l33
											}
											v20 = v27 + v29
											t118 := int32(m.memory[uint32(v20)])
											v25 = t118
											t119 := v20
											v22 = int32(uint32(int32(v35)) >> 25)
											m.memory[uint32(t119)] = byte(v22)
											t120 := int32(load32(m.memory[int64(uint32(v5))+340:]))
											m.memory[uint32(v27+t120&(v29+i32(-8))+i32(8))] = byte(v22)
											v29 = v27 + (i32(0)-v29)*i32(12)
											store32(m.memory[uint32(v29+i32(-4)):], uint32(v21))
											store32(m.memory[uint32(v29+i32(-8)):], uint32(v24))
											store32(m.memory[uint32(v29+i32(-12)):], uint32(v30))
											t121 := int32(load32(m.memory[int64(uint32(v5))+348:]))
											store32(m.memory[int64(uint32(v5))+348:], uint32(t121+i32(1)))
											t122 := int32(load32(m.memory[int64(uint32(v5))+344:]))
											store32(m.memory[int64(uint32(v5))+344:], uint32(t122-v25&i32(1)))
											goto l34
										}
									l33:
										v29 = v27 + (i32(0)-v29)*i32(12)
										store32(m.memory[uint32(v29+i32(-4)):], uint32(v21))
										store32(m.memory[uint32(v29+i32(-8)):], uint32(v24))
									l34:
										v1 = v1 + i32(12)
										v6 = v6 + i32(-12)
										goto l35
									}
								}
								t66 := m.fn857(v28, v26, v23, i32(1086064))
								v29 = t66
								t67 := int32(load32(m.memory[uint32(v29+i32(4)):]))
								t68 := int32(load32(m.memory[uint32(v29+i32(8)):]))
								t69 := m.fn1528(t67, t68, v6, i32(1086080))
								t70 := int32(m.memory[int64(uint32(t69))+24])
								v29 = t70
								t71 := m.fn857(v28, v26, v23, i32(1086096))
								v27 = t71
								t72 := int32(load32(m.memory[uint32(v27+i32(4)):]))
								t73 := int32(load32(m.memory[uint32(v27+i32(8)):]))
								t74 := m.fn1528(t72, t73, v6, i32(1086112))
								t75 := int32(load32(m.memory[int64(uint32(t74))+16:]))
								v30 = t75
								v21 = v6 + i32(1)
								{
									if v29 != 0 {
										{
											if v20 == 0 {
												goto l24
											}
											t78 := m.fn66(v32, v31, v1)
											t79 := v34
											v35 = t78
											v29 = t79 & int32(v35)
											v36 = int64(uint64(v35)>>25) & i64(127) * i64(72340172838076673)
											v25 = i32(0)
										l31:
											{
												t80 := int64(load64(m.memory[uint32(v33+v29):]))
												v37 = t80
												v35 = v37 ^ v36
												v35 = (v35 ^ i64(-1)) & (v35 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
											l27:
												{
													if v35 == 0 {
														if v37&(v37<<1)&i64(-0x7f7f7f7f7f7f7f80) == 0 {
															t94 := v29
															v25 = v25 + i32(8)
															v29 = (t94 + v25) & v34
															goto l31
														}
														goto l24
													}
													t81 := v1
													v27 = v33 + (i32(0)-(int32(uint32(int64(bits.TrailingZeros64(uint64(v35))))>>3)+v29)&v34)*i32(12)
													t82 := int32(load32(m.memory[uint32(v27+i32(-12)):]))
													if t81 == t82 {
														t83 := int32(load32(m.memory[uint32(v27+i32(-4)):]))
														v6 = t83
														t84 := int32(load32(m.memory[uint32(v27+i32(-8)):]))
														t85 := v28
														t86 := v26
														v27 = t84
														t87 := m.fn857(t85, t86, v27, i32(1086128))
														v29 = t87
														t88 := int32(load32(m.memory[uint32(v29+i32(4)):]))
														t89 := int32(load32(m.memory[uint32(v29+i32(8)):]))
														t90 := m.fn1528(t88, t89, v6, i32(1086144))
														v29 = t90
														t91 := int32(load32(m.memory[int64(uint32(v29))+20:]))
														store32(m.memory[int64(uint32(v29))+20:], uint32(t91+i32(1)))
														t92 := v1
														v30 = v30 + v1
														p93 := v30
														if uint32(v1) > uint32(v30) {
															p93 = t92
														}
														v29 = p93
													l30:
														if v29 != v1 {
															store32(m.memory[int64(uint32(v5))+344:], uint32(v6))
															store32(m.memory[int64(uint32(v5))+340:], uint32(v27))
															store32(m.memory[int64(uint32(v5))+336:], uint32(v1))
															v1 = v1 + i32(1)
															m.fn1529(v5+i32(192), v5+i32(336))
															goto l30
														}
														v6 = v21
														v1 = v30
														goto l22
													}
													v35 = (v35 + i64(-1)) & v35
													goto l27
												}
											}
										}
									l24:
										t95 := m.fn857(v28, v26, v23, i32(1086160))
										v29 = t95
										t96 := int32(load32(m.memory[uint32(v29+i32(4)):]))
										t97 := int32(load32(m.memory[uint32(v29+i32(8)):]))
										t98 := m.fn1528(t96, t97, v6, i32(1086176))
										m.memory[int64(uint32(t98))+24] = byte(i32(0))
										v1 = v30 + v1
										v6 = v21
										goto l22
									}
									t76 := v1
									v27 = v30 + v1
									p77 := v27
									if uint32(v1) > uint32(v27) {
										p77 = t76
									}
									v29 = p77
								l23:
									if v29 != v1 {
										store32(m.memory[int64(uint32(v5))+344:], uint32(v6))
										store32(m.memory[int64(uint32(v5))+340:], uint32(v23))
										store32(m.memory[int64(uint32(v5))+336:], uint32(v1))
										v1 = v1 + i32(1)
										m.fn1529(v5+i32(192), v5+i32(336))
										goto l23
									}
									v6 = v21
									v1 = v27
									goto l22
								}
							}
						}
					}
					t42 := int32(load32(m.memory[uint32(v27+i32(28)):]))
					t43 := int32(load32(m.memory[uint32(v27+i32(32)):]))
					t44 := m.fn886(t42, t43, i32(1072544), i32(60), i32(1074868), i32(4))
					v28 = t44
					v1 = i32(0)
					store32(m.memory[int64(uint32(v5))+200:], uint32(i32(0)))
					store64(m.memory[int64(uint32(v5))+192:], uint64(i64(0x400000000)))
					t45 := m.fn1526(v28, i32(1086192), i32(10))
					t46 := v5 + i32(192)
					v6 = t45
					m.fn902(t46, v6)
					v23 = v6 * i32(28)
					t47 := int32(load32(m.memory[int64(uint32(v5))+200:]))
					t48 := v6
					v29 = t47
					v30 = t48 + v29
					t49 := int32(load32(m.memory[int64(uint32(v5))+196:]))
					v29 = t49 + v29*i32(28)
				l13:
					if v23 == v1 {
						store32(m.memory[int64(uint32(v5))+200:], uint32(v30))
						m.fn1527(v27, v5+i32(192))
						t50 := m.fn1526(v28, i32(1086202), i32(9))
						t51 := v5 + i32(192)
						v1 = t50
						m.fn902(t51, v1)
						v23 = v1 * i32(28)
						t52 := int32(load32(m.memory[int64(uint32(v5))+200:]))
						t53 := v1
						v6 = t52
						v27 = t53 + v6
						t54 := int32(load32(m.memory[int64(uint32(v5))+196:]))
						v29 = t54 + v6*i32(28)
						v1 = i32(0)
					l15:
						if v23 == v1 {
							store32(m.memory[int64(uint32(v5))+200:], uint32(v27))
							store32(m.memory[int64(uint32(v5))+256:], uint32(v27))
							t55 := int64(load64(m.memory[int64(uint32(v5))+192:]))
							store64(m.memory[int64(uint32(v5))+248:], uint64(t55))
							{
								t56 := int32(load32(m.memory[int64(uint32(v5))+236:]))
								if v26 != t56 {
									goto l16
								}
								m.fn272(v5 + i32(236))
								t57 := int32(load32(m.memory[int64(uint32(v5))+240:]))
								v21 = t57
							}
						l16:
							v1 = v21 + v26*i32(12)
							t58 := int32(load32(m.memory[int64(uint32(v5))+256:]))
							store32(m.memory[int64(uint32(v1))+8:], uint32(t58))
							t59 := int64(load64(m.memory[int64(uint32(v5))+248:]))
							store64(m.memory[uint32(v1):], uint64(t59))
							t60 := v5
							v26 = v26 + i32(1)
							store32(m.memory[int64(uint32(t60))+244:], uint32(v26))
							goto l17
						}
						v6 = v29 + v1
						store64(m.memory[uint32(v6):], uint64(i64(0x400000000)))
						m.memory[uint32(v6+i32(24))] = byte(i32(0))
						store64(m.memory[uint32(v6+i32(16)):], uint64(i64(0x100000001)))
						store64(m.memory[uint32(v6+i32(8)):], uint64(i64(0)))
						v1 = v1 + i32(28)
						goto l15
					}
					v6 = v29 + v1
					store64(m.memory[uint32(v6):], uint64(i64(0x400000000)))
					m.memory[uint32(v6+i32(24))] = byte(i32(0))
					store64(m.memory[uint32(v6+i32(16)):], uint64(i64(0x100000001)))
					store64(m.memory[uint32(v6+i32(8)):], uint64(i64(0)))
					v1 = v1 + i32(28)
					goto l13
				}
			}
			t15 := m.fn15(v1, v6, i32(1077645), i32(3))
			if t15 != 0 {
				t135 := int32(load32(m.memory[uint32(v19+i32(28)):]))
				t136 := int32(load32(m.memory[uint32(v19+i32(32)):]))
				t137 := m.fn886(t135, t136, i32(1072544), i32(60), i32(1077680), i32(10))
				v1 = t137
				if v1 == 0 {
					goto l2
				}
				m.fn1323(v5+i32(336), v1, v2, v3, v4)
				t138 := int32(load32(m.memory[int64(uint32(v5))+336:]))
				v1 = t138
				if v1 == i32(-1) {
					goto l2
				}
				t139 := int32(load32(m.memory[int64(uint32(v5))+356:]))
				store32(m.memory[int64(uint32(v0))+20:], uint32(t139))
				t140 := int64(load64(m.memory[int64(uint32(v5))+348:]))
				store64(m.memory[int64(uint32(v0))+12:], uint64(t140))
				t141 := int64(load64(m.memory[int64(uint32(v5))+340:]))
				store64(m.memory[int64(uint32(v0))+4:], uint64(t141))
				store32(m.memory[uint32(v0):], uint32(v1))
				goto l6
			}
			t16 := m.fn15(v1, v6, i32(1077671), i32(9))
			if t16 == 0 {
				goto l2
			}
			m.fn1323(v5+i32(336), v19, v2, v3, v4)
			t17 := int32(load32(m.memory[int64(uint32(v5))+336:]))
			v1 = t17
			if v1 == i32(-1) {
				goto l2
			}
			t18 := int32(load32(m.memory[int64(uint32(v5))+356:]))
			store32(m.memory[int64(uint32(v0))+20:], uint32(t18))
			t19 := int64(load64(m.memory[int64(uint32(v5))+348:]))
			store64(m.memory[int64(uint32(v0))+12:], uint64(t19))
			t20 := int64(load64(m.memory[int64(uint32(v5))+340:]))
			store64(m.memory[int64(uint32(v0))+4:], uint64(t20))
			store32(m.memory[uint32(v0):], uint32(v1))
			goto l6
		}
	l8:
		{
			if v21 != 0 {
				goto l55
			}
			v29 = i32(2)
			goto l56
		l55:
			m.fn27(v5 + i32(336))
			m.fn1533(v5+i32(80), v15, v21, v24)
			{
				{
					t195 := int32(load32(m.memory[int64(uint32(v5))+80:]))
					v1 = t195
					if v1 != 0 {
						goto l57
					}
					v1 = i32(0)
					goto l71
				}
			l57:
				t196 := int32(load32(m.memory[int64(uint32(v1))+4:]))
				v6 = t196
				t197 := int32(load32(m.memory[uint32(v1):]))
				v1 = t197
			}
		l71:
			{
				if v1 == 0 {
					m.memory[int64(uint32(v5))+252] = byte(i32(2))
					goto l68
				}
				store32(m.memory[int64(uint32(v5))+156:], uint32(v6))
				store32(m.memory[int64(uint32(v5))+152:], uint32(v1))
				{
					t198 := m.fn1521(v5+i32(336), v1, v6)
					if t198 != 0 {
						t200 := m.fn1534(v15, v1, v6)
						v1 = t200
						t201 := int32(load32(m.memory[int64(uint32(v1))+8:]))
						v27 = t201
						t202 := int32(load32(m.memory[int64(uint32(v1))+4:]))
						v6 = t202
						{
							{
								t203 := int32(load32(m.memory[uint32(v1):]))
								v1 = t203
								v28 = v1 + i32(28)
								t204 := int32(load32(m.memory[uint32(v28):]))
								v30 = v1 + i32(32)
								t205 := int32(load32(m.memory[uint32(v30):]))
								t206 := m.fn886(t204, t205, i32(1072544), i32(60), i32(1073713), i32(4))
								v1 = t206
								if v1 != 0 {
									goto l62
								}
								v1 = i32(0)
								goto l63
							}
						l62:
							t207 := int32(load32(m.memory[uint32(v1+i32(16)):]))
							t208 := int32(load32(m.memory[uint32(v1+i32(20)):]))
							m.fn1046(v5+i32(72), t207, t208, i32(1072544), i32(60), i32(1073156), i32(3))
							t209 := int32(load32(m.memory[int64(uint32(v5))+76:]))
							v23 = t209
							t210 := int32(load32(m.memory[int64(uint32(v5))+72:]))
							v1 = t210
						}
					l63:
						t212 := v5 + i32(192)
						p211 := i32(1)
						if v1 != 0 {
							p211 = v1
						}
						p213 := i32(0)
						if v1 != 0 {
							p213 = v23
						}
						m.fn14(t212, p211, p213)
						t214 := int32(load32(m.memory[int64(uint32(v5))+196:]))
						t215 := v5 + i32(64)
						v1 = t214
						t216 := int32(load32(m.memory[int64(uint32(v5))+200:]))
						t217 := v1
						v29 = t216
						m.fn565(t215, t217, v29, i32(1073743), i32(8))
						{
							t218 := int32(load32(m.memory[int64(uint32(v5))+64:]))
							v23 = t218
							if v23 == 0 {
								goto l64
							}
							t219 := int32(load32(m.memory[int64(uint32(v5))+68:]))
							m.fn46(v5+i32(56), v23, t219)
							t220 := int32(load32(m.memory[int64(uint32(v5))+56:]))
							t221 := int32(load32(m.memory[int64(uint32(v5))+60:]))
							m.fn1423(v5+i32(48), t220, t221)
							t222 := int32(m.memory[int64(uint32(v5))+48])
							if t222 == 0 {
								t237 := int32(m.memory[int64(uint32(v5))+49])
								v26 = t237
								v23 = i32(1)
								goto l66
							}
						}
					l64:
						v26 = i32(1)
						v23 = i32(1)
						t223 := m.fn773(v1, v29, i32(1073751), i32(5))
						if t223 != 0 {
							goto l66
						}
						v23 = i32(2)
						t224 := int32(load32(m.memory[uint32(v28):]))
						t225 := int32(load32(m.memory[uint32(v30):]))
						t226 := m.fn886(t224, t225, i32(1072544), i32(60), i32(1073735), i32(3))
						v29 = t226
						if v29 == 0 {
							goto l66
						}
						t227 := int32(load32(m.memory[uint32(v29+i32(28)):]))
						t228 := int32(load32(m.memory[uint32(v29+i32(32)):]))
						t229 := m.fn886(t227, t228, i32(1072544), i32(60), i32(1073756), i32(10))
						v29 = t229
						if v29 == 0 {
							goto l66
						}
						t230 := int32(load32(m.memory[uint32(v29+i32(16)):]))
						t231 := int32(load32(m.memory[uint32(v29+i32(20)):]))
						m.fn1046(v5+i32(40), t230, t231, i32(1072544), i32(60), i32(1073156), i32(3))
						t232 := int32(load32(m.memory[int64(uint32(v5))+40:]))
						v29 = t232
						if v29 == 0 {
							goto l66
						}
						t233 := int32(load32(m.memory[int64(uint32(v5))+44:]))
						m.fn1423(v5+i32(32), v29, t233)
						t234 := int32(m.memory[int64(uint32(v5))+32])
						if t234 != 0 {
							goto l66
						}
						t235 := int32(m.memory[int64(uint32(v5))+33])
						v23 = t235
						v26 = v23 + i32(1)
						var p236 int32
						if uint32(v23&i32(255)) < uint32(i32(9)) {
							p236 = 1
						}
						v23 = p236
						goto l66
					}
					store32(m.memory[int64(uint32(v5))+196:], uint32(i32(71)))
					store32(m.memory[int64(uint32(v5))+192:], uint32(v5+i32(152)))
					m.fn73(v5+i32(248), i32(1049807), v5+i32(192))
					store32(m.memory[int64(uint32(v5))+260:], uint32(i32(-1)))
					t199 := int32(load32(m.memory[int64(uint32(v5))+248:]))
					v1 = t199
					goto l61
				}
			l66:
				t238 := int32(load32(m.memory[int64(uint32(v5))+192:]))
				m.fn16(t238, v1)
				if v23 == i32(2) {
					v1 = i32(0)
					if v6 != 0 {
						m.fn1533(v5+i32(24), v15, v6, v27)
						t239 := int32(load32(m.memory[int64(uint32(v5))+24:]))
						v23 = t239
						if v23 == 0 {
							goto l71
						}
						t240 := int32(load32(m.memory[int64(uint32(v23))+4:]))
						v6 = t240
						t241 := int32(load32(m.memory[uint32(v23):]))
						v1 = t241
						goto l71
					}
					goto l71
				}
				m.memory[int64(uint32(v5))+253] = byte(v26)
				m.memory[int64(uint32(v5))+252] = byte(v23)
				goto l68
			}
		l68:
			v1 = i32(-1)
		l61:
			t242 := int32(load32(m.memory[int64(uint32(v5))+336:]))
			t243 := int32(load32(m.memory[int64(uint32(v5))+340:]))
			m.fn56(t242, t243)
			{
				if v1 == i32(-1) {
					goto l72
				}
				t244 := int64(load32(m.memory[int64(uint32(v5))+252:]))
				v36 = t244<<32 | int64(uint32(v1))
				t245 := int32(load16(m.memory[int64(uint32(v5))+257:]))
				t246 := int32(m.memory[uint32(v5+i32(259))])
				v27 = t245 | t246<<16
				goto l73
			}
		l72:
			t247 := int32(m.memory[int64(uint32(v5))+253])
			v30 = t247
			t248 := int32(m.memory[int64(uint32(v5))+252])
			v29 = t248
		}
	l56:
		{
			{
				{
					if v21 == 0 {
						m.fn1535(v5+i32(336), v20, i32(0), v24, v2)
						t274 := int32(load16(m.memory[int64(uint32(v5))+345:]))
						t275 := int32(m.memory[uint32(v16)])
						v27 = t274 | t275<<16
						t276 := int64(load64(m.memory[int64(uint32(v5))+352:]))
						v35 = t276
						t277 := int32(load32(m.memory[int64(uint32(v5))+348:]))
						v28 = t277
						t278 := int32(m.memory[int64(uint32(v5))+344])
						v26 = t278
						t279 := int64(load64(m.memory[int64(uint32(v5))+336:]))
						v36 = t279
						t280 := int32(load32(m.memory[int64(uint32(v5))+364:]))
						v1 = t280
						if v1 == i32(-3) {
							goto l85
						}
						store16(m.memory[int64(uint32(v5))+161:], uint16(v27))
						m.memory[uint32(v5+i32(152)+i32(11))] = byte(int32(uint32(v27) >> 16))
						t281 := int64(load64(m.memory[int64(uint32(v5))+368:]))
						t282 := v5
						v31 = t281
						store64(m.memory[int64(uint32(t282))+184:], uint64(v31))
						store32(m.memory[int64(uint32(v5))+180:], uint32(v1))
						t283 := int32(load32(m.memory[int64(uint32(v5))+360:]))
						t284 := v5
						v27 = t283
						store32(m.memory[int64(uint32(t284))+176:], uint32(v27))
						store64(m.memory[int64(uint32(v5))+168:], uint64(v35))
						store32(m.memory[int64(uint32(v5))+164:], uint32(v28))
						m.memory[int64(uint32(v5))+160] = byte(v26)
						store64(m.memory[int64(uint32(v5))+152:], uint64(v36))
						v24 = int32(v31)
						v6 = i32(0)
						v33 = i32(2)
						goto l86
					}
					m.fn27(v5 + i32(336))
					m.fn1533(v5+i32(16), v15, v21, v24)
					{
						{
							t249 := int32(load32(m.memory[int64(uint32(v5))+16:]))
							v1 = t249
							if v1 != 0 {
								goto l75
							}
							v1 = i32(0)
							goto l83
						}
					l75:
						t250 := int32(load32(m.memory[int64(uint32(v1))+4:]))
						v6 = t250
						t251 := int32(load32(m.memory[uint32(v1):]))
						v1 = t251
					}
				l83:
					if v1 == 0 {
						goto l77
					}
					store32(m.memory[int64(uint32(v5))+156:], uint32(v6))
					store32(m.memory[int64(uint32(v5))+152:], uint32(v1))
					{
						t252 := m.fn1521(v5+i32(336), v1, v6)
						if t252 != 0 {
							t254 := m.fn1534(v15, v1, v6)
							v1 = t254
							t255 := int32(load32(m.memory[int64(uint32(v1))+8:]))
							v23 = t255
							t256 := int32(load32(m.memory[int64(uint32(v1))+4:]))
							v6 = t256
							{
								t257 := int32(load32(m.memory[uint32(v1):]))
								v1 = t257
								t258 := int32(load32(m.memory[uint32(v1+i32(28)):]))
								t259 := int32(load32(m.memory[uint32(v1+i32(32)):]))
								t260 := m.fn886(t258, t259, i32(1072544), i32(60), i32(1073713), i32(4))
								v1 = t260
								if v1 == 0 {
									goto l80
								}
								t261 := int32(load32(m.memory[uint32(v1+i32(16)):]))
								t262 := int32(load32(m.memory[uint32(v1+i32(20)):]))
								m.fn1046(v5+i32(8), t261, t262, i32(1072544), i32(60), i32(1073156), i32(3))
								t263 := int32(load32(m.memory[int64(uint32(v5))+8:]))
								v1 = t263
								if v1 == 0 {
									goto l80
								}
								t264 := int32(load32(m.memory[int64(uint32(v5))+12:]))
								t265 := m.fn1208(v1, t264)
								v1 = t265 & i32(255)
								if v1 == i32(2) {
									goto l80
								}
								m.memory[int64(uint32(v5))+252] = byte(v1)
								goto l81
							}
						l80:
							v1 = i32(0)
							if v6 != 0 {
								m.fn1533(v5, v15, v6, v23)
								t266 := int32(load32(m.memory[uint32(v5):]))
								v23 = t266
								if v23 == 0 {
									goto l83
								}
								t267 := int32(load32(m.memory[int64(uint32(v23))+4:]))
								v6 = t267
								t268 := int32(load32(m.memory[uint32(v23):]))
								v1 = t268
								goto l83
							}
							goto l83
						}
						store32(m.memory[int64(uint32(v5))+196:], uint32(i32(71)))
						store32(m.memory[int64(uint32(v5))+192:], uint32(v5+i32(152)))
						m.fn73(v5+i32(248), i32(1049807), v5+i32(192))
						store32(m.memory[int64(uint32(v5))+260:], uint32(i32(-1)))
						t253 := int32(load32(m.memory[int64(uint32(v5))+248:]))
						v1 = t253
						goto l79
					}
				l77:
					m.memory[int64(uint32(v5))+252] = byte(i32(2))
				l81:
					v1 = i32(-1)
				l79:
					t269 := int32(load32(m.memory[int64(uint32(v5))+336:]))
					t270 := int32(load32(m.memory[int64(uint32(v5))+340:]))
					m.fn56(t269, t270)
					if v1 == i32(-1) {
						goto l84
					}
					t271 := int64(load32(m.memory[int64(uint32(v5))+252:]))
					v36 = t271<<32 | int64(uint32(v1))
					t272 := int32(load16(m.memory[int64(uint32(v5))+257:]))
					t273 := int32(m.memory[uint32(v5+i32(259))])
					v27 = t272 | t273<<16
					goto l73
				}
			l84:
				t285 := int32(m.memory[int64(uint32(v5))+252])
				v33 = t285
				m.fn1535(v5+i32(336), v20, v21, v24, v2)
				t286 := int32(load16(m.memory[int64(uint32(v5))+345:]))
				t287 := int32(m.memory[uint32(v16)])
				v27 = t286 | t287<<16
				t288 := int64(load64(m.memory[int64(uint32(v5))+352:]))
				v35 = t288
				t289 := int32(load32(m.memory[int64(uint32(v5))+348:]))
				v28 = t289
				t290 := int32(m.memory[int64(uint32(v5))+344])
				v26 = t290
				t291 := int64(load64(m.memory[int64(uint32(v5))+336:]))
				v36 = t291
				t292 := int32(load32(m.memory[int64(uint32(v5))+364:]))
				v1 = t292
				if v1 == i32(-3) {
					goto l85
				}
				store16(m.memory[int64(uint32(v5))+161:], uint16(v27))
				m.memory[uint32(v5+i32(152)+i32(11))] = byte(int32(uint32(v27) >> 16))
				t293 := int64(load64(m.memory[int64(uint32(v5))+368:]))
				t294 := v5
				v31 = t293
				store64(m.memory[int64(uint32(t294))+184:], uint64(v31))
				store32(m.memory[int64(uint32(v5))+180:], uint32(v1))
				t295 := int32(load32(m.memory[int64(uint32(v5))+360:]))
				t296 := v5
				v27 = t295
				store32(m.memory[int64(uint32(t296))+176:], uint32(v27))
				store64(m.memory[int64(uint32(v5))+168:], uint64(v35))
				store32(m.memory[int64(uint32(v5))+164:], uint32(v28))
				m.memory[int64(uint32(v5))+160] = byte(v26)
				store64(m.memory[int64(uint32(v5))+152:], uint64(v36))
				m.fn1536(v5+i32(336), v15, v21, v24)
				v24 = int32(v31)
				{
					t297 := int32(load32(m.memory[int64(uint32(v5))+336:]))
					v6 = t297
					if v6 == i32(-1) {
						goto l87
					}
					t298 := int64(load32(m.memory[int64(uint32(v5))+340:]))
					v36 = t298<<32 | int64(uint32(v6))
					t299 := int32(load16(m.memory[int64(uint32(v5))+345:]))
					t300 := int32(m.memory[uint32(v5+i32(347))])
					v27 = t299 | t300<<16
					t301 := int64(load64(m.memory[int64(uint32(v5))+352:]))
					v35 = t301
					t302 := int32(load32(m.memory[int64(uint32(v5))+348:]))
					v28 = t302
					t303 := int32(m.memory[int64(uint32(v5))+344])
					v26 = t303
					if v1 == i32(-2) {
						goto l85
					}
					goto l88
				}
			l87:
				t304 := int32(load16(m.memory[int64(uint32(v5))+340:]))
				t305 := int32(m.memory[uint32(v18)])
				v6 = t304 | t305<<16
			}
		l86:
			t306 := int32(load32(m.memory[int64(uint32(v15))+32:]))
			t307 := fn1537(v6, t306)
			v6 = t307
			{
				{
					var p308 int32
					if uint32(v25&i32(255)) < uint32(i32(9)) {
						p308 = 1
					}
					t309 := v29
					v23 = v22 & i32(1)
					p310 := t309
					if v23 != 0 {
						p310 = p308
					}
					if p310&i32(1) == 0 {
						if v1 == i32(-2) {
							if v33&i32(255) == i32(2) {
								store32(m.memory[int64(uint32(v5))+220:], uint32(i32(-0x7ffffffd)))
								goto l95
							}
							store32(m.memory[int64(uint32(v5))+220:], uint32(i32(-0x7ffffffe)))
							m.memory[int64(uint32(v5))+192] = byte(v33)
							goto l95
						}
						t312 := int32(load32(m.memory[int64(uint32(v13))+8:]))
						store32(m.memory[int64(uint32(v12))+8:], uint32(t312))
						t313 := int64(load64(m.memory[uint32(v13):]))
						store64(m.memory[uint32(v12):], uint64(t313))
						store32(m.memory[int64(uint32(v5))+216:], uint32(v27))
						store64(m.memory[int64(uint32(v5))+208:], uint64(v35))
						m.memory[int64(uint32(v5))+200] = byte(v26)
						store64(m.memory[int64(uint32(v5))+192:], uint64(v36))
						v23 = i32(0)
						goto l93
					}
					p311 := v30
					if v23 != 0 {
						p311 = v25 + i32(1)
					}
					v29 = p311
					v23 = i32(-1)
					if v1 != i32(-2) {
						goto l90
					}
					goto l91
				}
			l90:
				if v26&i32(255) == 0 {
					goto l91
				}
				m.fn225(v5+i32(336), v13)
				{
					t314 := int32(load32(m.memory[int64(uint32(v5))+336:]))
					if t314 == i32(-1) {
						goto l96
					}
					t315 := int32(load32(m.memory[int64(uint32(v5))+344:]))
					store32(m.memory[int64(uint32(v5))+256:], uint32(t315))
					t316 := int64(load64(m.memory[int64(uint32(v5))+336:]))
					store64(m.memory[int64(uint32(v5))+248:], uint64(t316))
					goto l97
				}
			l96:
				m.fn800(v5+i32(248), v26, v35)
			l97:
				store32(m.memory[int64(uint32(v5))+308:], uint32(i32(25)))
				store32(m.memory[int64(uint32(v5))+304:], uint32(v5+i32(248)))
				m.fn73(v5+i32(336), i32(1070105), v5+i32(304))
				t317 := int32(load32(m.memory[int64(uint32(v5))+248:]))
				t318 := int32(load32(m.memory[int64(uint32(v5))+252:]))
				m.fn16(t317, t318)
				t319 := int32(load32(m.memory[int64(uint32(v5))+336:]))
				v23 = t319
				t320 := int64(load64(m.memory[int64(uint32(v5))+340:]))
				v36 = t320
			}
		l91:
			store32(m.memory[int64(uint32(v5))+220:], uint32(i32(-0x80000000)))
			store32(m.memory[int64(uint32(v5))+204:], uint32(v6))
			store64(m.memory[int64(uint32(v5))+196:], uint64(v36))
			store32(m.memory[int64(uint32(v5))+192:], uint32(v23))
			m.memory[int64(uint32(v5))+208] = byte(v29)
		l95:
			v23 = i32(1)
		l93:
			store32(m.memory[int64(uint32(v5))+284:], uint32(v6))
			store32(m.memory[int64(uint32(v5))+288:], uint32(v2))
			store32(m.memory[int64(uint32(v5))+280:], uint32(i32(0)))
			store64(m.memory[int64(uint32(v5))+272:], uint64(i64(0x400000000)))
			store64(m.memory[int64(uint32(v5))+264:], uint64(i64(4)))
			store64(m.memory[int64(uint32(v5))+256:], uint64(i64(0)))
			store64(m.memory[int64(uint32(v5))+248:], uint64(i64(0x400000000)))
			m.fn1538(v5+i32(336), v5+i32(248), v19)
			{
				t321 := int32(load32(m.memory[int64(uint32(v5))+336:]))
				v6 = t321
				if v6 == i32(-1) {
					t334 := int32(m.memory[uint32(v5+i32(192)+i32(11))])
					v29 = t334
					t335 := int32(load32(m.memory[int64(uint32(v5))+192:]))
					v30 = t335
					t336 := int32(load32(m.memory[int64(uint32(v5))+196:]))
					v19 = t336
					t337 := int64(load64(m.memory[int64(uint32(v5))+192:]))
					v36 = t337
					t338 := int32(m.memory[int64(uint32(v5))+200])
					v26 = t338
					t339 := int32(load32(m.memory[int64(uint32(v5))+204:]))
					v28 = t339
					t340 := int64(load64(m.memory[int64(uint32(v5))+208:]))
					v35 = t340
					t341 := int32(load32(m.memory[int64(uint32(v5))+216:]))
					v20 = t341
					t342 := int32(load32(m.memory[int64(uint32(v5))+220:]))
					v6 = t342
					t343 := int64(load64(m.memory[int64(uint32(v5))+224:]))
					v32 = t343
					t344 := int32(load16(m.memory[int64(uint32(v5))+201:]))
					v27 = t344
					memory_copy(m.memory, uint32(v5+i32(336)), uint32(v5+i32(248)), uint32(i32(44)))
					v21 = v29 << 16
					m.fn1540(v5+i32(304), v5+i32(336))
					t345 := int32(load32(m.memory[int64(uint32(v5))+304:]))
					v29 = t345
					t346 := int64(load64(m.memory[int64(uint32(v5))+308:]))
					v31 = t346
					{
						t347 := v23 ^ i32(1)
						var p348 int32
						if v1 == i32(-2) {
							p348 = 1
						}
						if t347|p348 != 0 {
							goto l102
						}
						m.fn134(v1, v24)
					}
				l102:
					v27 = v27 | v21
					if v29 == i32(-1) {
						goto l85
					}
					store32(m.memory[int64(uint32(v5))+140:], uint32(v29))
					store64(m.memory[int64(uint32(v5))+144:], uint64(v31))
					v21 = int32(int64(uint64(v31) >> 32))
					v1 = int32(v31)
					{
						v23 = v6 ^ i32(-0x80000000)
						p349 := i32(1)
						if uint32(v23) < uint32(i32(4)) {
							p349 = v23
						}
						switch p349 {
						case 1:
							store64(m.memory[int64(uint32(v5))+368:], uint64(v32))
							store32(m.memory[int64(uint32(v5))+364:], uint32(v6))
							m.fn1333(v4, v3)
							m.fn1541(v11, v5+i32(140))
							store64(m.memory[int64(uint32(v5))+352:], uint64(v35))
							m.memory[int64(uint32(v5))+344] = byte(v26)
							store64(m.memory[int64(uint32(v5))+336:], uint64(v36))
							store32(m.memory[int64(uint32(v5))+360:], uint32(v20))
							m.fn1369(v7, v5+i32(336))
							goto l2
						case 3:
							m.fn1325(v4, v3)
							m.fn1541(v5+i32(336), v5+i32(140))
							m.fn1271(v3, v5+i32(336))
							goto l2
						default:
							v24 = int32(v35)
							m.fn1325(v4, v3)
							t350 := v5
							t351 := v1
							v23 = v21 << 4
							v20 = t351 + v23
							store32(m.memory[int64(uint32(t350))+204:], uint32(v20))
							store32(m.memory[int64(uint32(v5))+200:], uint32(v29))
							store32(m.memory[int64(uint32(v5))+192:], uint32(v1))
							v21 = i32(0)
						l115:
							if v23 == 0 {
								goto l107
							}
							{
								t352 := int32(load32(m.memory[uint32(v1):]))
								v29 = t352
								if v29 == i32(2) {
									goto l108
								}
								t353 := v10
								v6 = v1 + i32(4)
								t354 := int64(load64(m.memory[uint32(v6):]))
								store64(m.memory[uint32(t353):], uint64(t354))
								t355 := int32(load32(m.memory[int64(uint32(v6))+8:]))
								store32(m.memory[int64(uint32(v10))+8:], uint32(t355))
								store32(m.memory[int64(uint32(v5))+248:], uint32(v29))
								if v29&i32(1) == 0 {
									goto l109
								}
								t356 := int32(load32(m.memory[int64(uint32(v6))+8:]))
								store32(m.memory[int64(uint32(v5))+344:], uint32(t356))
								t357 := int64(load64(m.memory[uint32(v6):]))
								store64(m.memory[int64(uint32(v5))+336:], uint64(t357))
								m.fn1271(v3, v5+i32(336))
								goto l110
							}
						l108:
							v20 = v1 + i32(16)
						l107:
							store32(m.memory[int64(uint32(v5))+196:], uint32(v20))
							m.fn1542(v5 + i32(192))
							m.fn134(v30, v19)
							goto l2
						l109:
							{
								{
									t358 := int32(load32(m.memory[int64(uint32(v5))+256:]))
									t359 := int32(load32(m.memory[int64(uint32(v5))+260:]))
									t360 := m.fn23(t358, t359)
									if t360 != 0 {
										m.fn894(v10)
										goto l110
									}
									t361 := int64(load64(m.memory[uint32(v6):]))
									store64(m.memory[int64(uint32(v5))+152:], uint64(t361))
									t362 := int32(load32(m.memory[int64(uint32(v6))+8:]))
									t363 := v5
									v6 = t362
									store32(m.memory[int64(uint32(t363))+160:], uint32(v6))
									t364 := int32(load32(m.memory[int64(uint32(v5))+156:]))
									m.fn1368(t364, v6, v28)
									if v21&i32(1) != 0 {
										goto l112
									}
									if v30 == i32(-1) {
										goto l113
									}
									store16(m.memory[int64(uint32(v5))+349:], uint16(v27))
									m.memory[uint32(v17)] = byte(int32(uint32(v27) >> 16))
									store32(m.memory[int64(uint32(v5))+352:], uint32(i32(0)))
									m.memory[int64(uint32(v5))+348] = byte(v26)
									store32(m.memory[int64(uint32(v5))+344:], uint32(v19))
									store32(m.memory[int64(uint32(v5))+340:], uint32(v30))
									store32(m.memory[int64(uint32(v5))+336:], uint32(i32(3)))
									m.fn1163(v5+i32(152), v5+i32(336))
								l113:
									t365 := int32(load32(m.memory[int64(uint32(v5))+160:]))
									store32(m.memory[int64(uint32(v5))+344:], uint32(t365))
									t366 := int64(load64(m.memory[int64(uint32(v5))+152:]))
									store64(m.memory[int64(uint32(v5))+336:], uint64(t366))
									m.memory[int64(uint32(v5))+360] = byte(v24)
									v30 = i32(-1)
									store32(m.memory[int64(uint32(v5))+348:], uint32(i32(-1)))
									m.fn338(v3, v5+i32(336))
									goto l114
								}
							l112:
								t367 := int32(load32(m.memory[int64(uint32(v5))+160:]))
								store32(m.memory[int64(uint32(v9))+8:], uint32(t367))
								t368 := int64(load64(m.memory[int64(uint32(v5))+152:]))
								store64(m.memory[uint32(v9):], uint64(t368))
								store32(m.memory[int64(uint32(v5))+336:], uint32(i32(-0x80000000)))
								m.fn338(v3, v5+i32(336))
							}
						l114:
							v21 = i32(1)
						l110:
							v1 = v1 + i32(16)
							v23 = v23 + i32(-16)
							goto l115
						case 2:
							m.fn1351(v3, v7)
							t369 := v5
							t370 := v1
							v6 = v21 << 4
							v27 = t370 + v6
							store32(m.memory[int64(uint32(t369))+348:], uint32(v27))
							store32(m.memory[int64(uint32(v5))+344:], uint32(v29))
							store32(m.memory[int64(uint32(v5))+336:], uint32(v1))
							v29 = int32(v36) & i32(1)
						l121:
							{
								if v6 == 0 {
									goto l116
								}
								v23 = v1 + i32(4)
								{
									t371 := int32(load32(m.memory[uint32(v1):]))
									switch t371 {
									case 0:
										goto l117
									default:
										t372 := int32(load32(m.memory[int64(uint32(v23))+8:]))
										store32(m.memory[int64(uint32(v5))+256:], uint32(t372))
										t373 := int64(load64(m.memory[uint32(v23):]))
										store64(m.memory[int64(uint32(v5))+248:], uint64(t373))
										m.fn1333(v4, v3)
										m.fn1271(v3, v5+i32(248))
										goto l120
									case 2:
										v27 = v1 + i32(16)
									}
								}
							l116:
								store32(m.memory[int64(uint32(v5))+340:], uint32(v27))
								m.fn1542(v5 + i32(336))
								goto l2
							l117:
								t374 := int32(load32(m.memory[int64(uint32(v23))+8:]))
								store32(m.memory[int64(uint32(v5))+256:], uint32(t374))
								t375 := int64(load64(m.memory[uint32(v23):]))
								store64(m.memory[int64(uint32(v5))+248:], uint64(t375))
								m.fn1445(v4, v29, v5+i32(248), v3)
							}
						l120:
							v1 = v1 + i32(16)
							v6 = v6 + i32(-16)
							goto l121
						}
					}
				}
				t322 := int32(m.memory[uint32(v5+i32(347))])
				v29 = t322
				t323 := int64(load64(m.memory[int64(uint32(v5))+352:]))
				v35 = t323
				t324 := int32(load32(m.memory[int64(uint32(v5))+348:]))
				v28 = t324
				t325 := int32(m.memory[int64(uint32(v5))+344])
				v26 = t325
				t326 := int32(load16(m.memory[int64(uint32(v5))+345:]))
				v27 = t326
				t327 := int64(load32(m.memory[int64(uint32(v5))+340:]))
				v36 = t327
				m.fn1539(v5 + i32(248))
				v36 = v36 << 32
				v31 = int64(uint32(v6))
				v29 = v29 << 16
				v6 = v5 + i32(192)
				{
					t328 := int32(load32(m.memory[int64(uint32(v5))+220:]))
					v30 = t328 ^ i32(-0x80000000)
					p329 := i32(1)
					if uint32(v30) < uint32(i32(4)) {
						p329 = v30
					}
					switch p329 {
					case 1:
						v6 = v12
						fallthrough
					case 0:
						t330 := int32(load32(m.memory[uint32(v6):]))
						t331 := int32(load32(m.memory[uint32(v6+i32(4)):]))
						m.fn134(t330, t331)
						fallthrough
					default:
						v36 = v36 | v31
						v27 = v27 | v29
						t332 := v23
						var p333 int32
						if v1 != i32(-2) {
							p333 = 1
						}
						if t332&p333 == 0 {
							goto l85
						}
						goto l88
					}
				}
			}
		}
	l88:
		m.fn134(v1, v24)
		goto l85
	l73:
		t376 := int64(load64(m.memory[int64(uint32(v5))+264:]))
		v35 = t376
		t377 := int32(load32(m.memory[int64(uint32(v5))+260:]))
		v28 = t377
		t378 := int32(m.memory[int64(uint32(v5))+256])
		v26 = t378
	}
l85:
	store16(m.memory[int64(uint32(v0))+9:], uint16(v27))
	store64(m.memory[int64(uint32(v0))+16:], uint64(v35))
	store32(m.memory[int64(uint32(v0))+12:], uint32(v28))
	m.memory[int64(uint32(v0))+8] = byte(v26)
	store64(m.memory[uint32(v0):], uint64(v36))
	m.memory[uint32(v0+i32(11))] = byte(int32(uint32(v27) >> 16))
l6:
	m.g0 = v5 + i32(400)
}
func (m *Module) fn1324(v0 int32) {
	m.fn1302(v0 + i32(16))
	m.fn1332(v0)
}
func (m *Module) fn1325(v0, v1 int32) {
	m.fn1333(v0, v1)
	m.fn1351(v1, v0+i32(16))
}
func (m *Module) fn1326(v0 int32) {
	var v1 int32
	t0 := m.g0
	v1 = t0 - i32(16)
	m.g0 = v1
	store64(m.memory[uint32(v1):], uint64(int64(uint32(i32(187)))<<32|int64(uint32(v1+i32(15)))))
	m.fn91(i32(1052692), v1, v0)
	panic("unreachable")
}
func (m *Module) fn1327(v0, v1 int32) int32 {
	var v2, v3, v4, v5 int32
	t0 := m.g0
	v2 = t0 - i32(32)
	m.g0 = v2
	{
		{
			t1 := m.fn886(v0, v1, i32(1077249), i32(47), i32(1085814), i32(15))
			v1 = t1
			if v1 != 0 {
				goto l0
			}
			v1 = i32(33686018)
			goto l1
		}
	l0:
		t2 := int32(load32(m.memory[uint32(v1+i32(16)):]))
		t3 := v2 + i32(16)
		v0 = t2
		t4 := int32(load32(m.memory[uint32(v1+i32(20)):]))
		t5 := v0
		v3 = t4
		m.fn1046(t3, t5, v3, i32(1085829), i32(59), i32(1081080), i32(11))
		{
			{
				t6 := int32(load32(m.memory[int64(uint32(v2))+16:]))
				v4 = t6
				if v4 != 0 {
					goto l2
				}
				v4 = i32(33554434)
				goto l3
			}
		l2:
			{
				t7 := int32(load32(m.memory[int64(uint32(v2))+20:]))
				t8 := v4
				v5 = t7
				t9 := m.fn15(t8, v5, i32(1074847), i32(4))
				if t9 == 0 {
					goto l4
				}
				v4 = i32(33554433)
				goto l3
			}
		l4:
			m.fn1071(v2+i32(24), v4, v5)
			t10 := int32(m.memory[int64(uint32(v2))+24])
			t11 := int32(load32(m.memory[int64(uint32(v2))+28:]))
			t12 := t10 ^ i32(-1)
			var p13 int32
			if uint32(t11) > uint32(i32(599)) {
				p13 = 1
			}
			v4 = t12&p13 | i32(0x2000000)
			t14 := int32(load32(m.memory[uint32(v1+i32(20)):]))
			v3 = t14
			t15 := int32(load32(m.memory[uint32(v1+i32(16)):]))
			v0 = t15
		}
	l3:
		m.fn1046(v2+i32(8), v0, v3, i32(1085829), i32(59), i32(1081091), i32(10))
		{
			{
				t16 := int32(load32(m.memory[int64(uint32(v2))+8:]))
				v1 = t16
				if v1 != 0 {
					goto l5
				}
				v1 = i32(512)
				goto l6
			}
		l5:
			{
				t17 := int32(load32(m.memory[int64(uint32(v2))+12:]))
				t18 := v1
				v5 = t17
				t19 := m.fn15(t18, v5, i32(1074855), i32(6))
				if t19 == 0 {
					goto l7
				}
				v1 = i32(256)
				goto l6
			}
		l7:
			t20 := m.fn15(v1, v5, i32(1074861), i32(7))
			p21 := i32(0)
			if t20 != 0 {
				p21 = i32(256)
			}
			v1 = p21
		}
	l6:
		m.fn1046(v2, v0, v3, i32(1077249), i32(47), i32(1085888), i32(23))
		{
			{
				t22 := int32(load32(m.memory[uint32(v2):]))
				v0 = t22
				if v0 != 0 {
					goto l8
				}
				v0 = i32(0x20000)
				goto l9
			}
		l8:
			t23 := int32(load32(m.memory[int64(uint32(v2))+4:]))
			t24 := m.fn1337(v0, t23, i32(1074851), i32(4))
			p25 := i32(0)
			if t24 != 0 {
				p25 = i32(65536)
			}
			v0 = p25
		}
	l9:
		v1 = v4 | v1 | v0
	}
l1:
	m.g0 = v2 + i32(32)
	return v1
}
func (m *Module) fn1328(v0, v1, v2, v3, v4 int32) {
	var v5 int32
	t0 := m.g0
	v5 = t0 - i32(32)
	m.g0 = v5
	store32(m.memory[int64(uint32(v5))+4:], uint32(v2))
	store32(m.memory[uint32(v5):], uint32(v1))
	store32(m.memory[int64(uint32(v5))+12:], uint32(v4))
	store32(m.memory[int64(uint32(v5))+8:], uint32(v3))
	store32(m.memory[int64(uint32(v5))+28:], uint32(i32(1)))
	store32(m.memory[int64(uint32(v5))+20:], uint32(i32(1)))
	store32(m.memory[int64(uint32(v5))+24:], uint32(v5+i32(8)))
	store32(m.memory[int64(uint32(v5))+16:], uint32(v5))
	m.fn73(v0, i32(1085911), v5+i32(16))
	m.g0 = v5 + i32(32)
}
func (m *Module) fn1329(v0, v1, v2 int32) {
	var v3, v4, v5, v6, v7, v8, v9 int32
	t0 := m.g0
	v3 = t0 - i32(480)
	m.g0 = v3
	v4 = i32(0)
l1:
	if v4 == i32(400) {
		store32(m.memory[int64(uint32(v3))+56:], uint32(v1))
		store32(m.memory[int64(uint32(v3))+60:], uint32(v1+v2*i32(44)))
	l3:
		{
			t1 := m.fn904(v3 + i32(56))
			v5 = t1
			if v5 == 0 {
				memory_copy(m.memory, uint32(v0), uint32(v3+i32(80)), uint32(i32(400)))
				m.g0 = v3 + i32(480)
				return
			}
			t2 := v3 + i32(48)
			v6 = v5 + i32(16)
			t3 := int32(load32(m.memory[uint32(v6):]))
			v2 = t3
			t4 := v2
			v7 = v5 + i32(20)
			t5 := int32(load32(m.memory[uint32(v7):]))
			v8 = t5
			m.fn1046(t2, t4, v8, i32(1074680), i32(46), i32(1085692), i32(5))
			t6 := int32(load32(m.memory[int64(uint32(v3))+52:]))
			v4 = t6
			t7 := int32(load32(m.memory[int64(uint32(v3))+48:]))
			v1 = t7
			if v1 == 0 {
				goto l3
			}
			m.fn197(v3+i32(64), v1, v4)
			t8 := int32(load32(m.memory[int64(uint32(v3))+68:]))
			v4 = t8
			t9 := int32(m.memory[int64(uint32(v3))+64])
			if t9 != 0 {
				goto l3
			}
			t10 := m.fn1468(i32(1085700), v4)
			if t10 == 0 {
				goto l3
			}
			v4 = v4 + i32(-1)
			if uint32(v4) > uint32(i32(9)) {
				m.fn158(v4, i32(10), i32(1085712))
				panic("unreachable")
			}
			{
				t11 := m.fn847(v5, i32(1074680), i32(46), i32(1085728), i32(23))
				if t11 != 0 {
					goto l5
				}
				t12 := m.fn847(v5, i32(1074680), i32(46), i32(1085751), i32(19))
				if t12 == 0 {
					goto l3
				}
			}
		l5:
			v5 = v3 + i32(80) + v4*i32(40)
			m.fn1046(v3+i32(40), v2, v8, i32(1077249), i32(47), i32(1085770), i32(10))
			v1 = i32(1)
			v4 = i32(1)
			{
				t13 := int32(load32(m.memory[int64(uint32(v3))+40:]))
				v9 = t13
				if v9 == 0 {
					goto l6
				}
				{
					t14 := int32(load32(m.memory[int64(uint32(v3))+44:]))
					t15 := v9
					v4 = t14
					t16 := m.fn15(t15, v4, i32(1077050), i32(1))
					if t16 == 0 {
						goto l7
					}
					v4 = i32(2)
					goto l6
				}
			l7:
				{
					t17 := m.fn15(v9, v4, i32(1077051), i32(1))
					if t17 == 0 {
						goto l8
					}
					v4 = i32(3)
					goto l6
				}
			l8:
				{
					t18 := m.fn15(v9, v4, i32(1073721), i32(1))
					if t18 == 0 {
						goto l9
					}
					v4 = i32(4)
					goto l6
				}
			l9:
				{
					t19 := m.fn15(v9, v4, i32(1077052), i32(1))
					if t19 == 0 {
						goto l10
					}
					v4 = i32(5)
					goto l6
				}
			l10:
				t20 := m.fn15(v9, v4, i32(1), i32(0))
				v4 = t20 ^ i32(1)
			}
		l6:
			m.memory[int64(uint32(v5))+36] = byte(v4)
			m.fn1046(v3+i32(32), v2, v8, i32(1074680), i32(46), i32(1085180), i32(11))
			t21 := int32(load32(m.memory[int64(uint32(v3))+32:]))
			t22 := int32(load32(m.memory[int64(uint32(v3))+36:]))
			m.fn1514(v3+i32(64), t21, t22)
			t23 := int64(load64(m.memory[int64(uint32(v3))+72:]))
			t24 := int32(load32(m.memory[int64(uint32(v3))+64:]))
			t26 := v5
			p25 := i64(1)
			if t24 != 0 {
				p25 = t23
			}
			store64(m.memory[uint32(t26):], uint64(p25))
			m.fn1046(v3+i32(24), v2, v8, i32(1077249), i32(47), i32(1085780), i32(10))
			t27 := int32(load32(m.memory[int64(uint32(v3))+24:]))
			t28 := v3 + i32(64)
			v4 = t27
			p29 := i32(1)
			if v4 != 0 {
				p29 = v4
			}
			t30 := int32(load32(m.memory[int64(uint32(v3))+28:]))
			p31 := i32(0)
			if v4 != 0 {
				p31 = t30
			}
			m.fn51(t28, p29, p31)
			t32 := int32(load32(m.memory[int64(uint32(v5))+8:]))
			t33 := int32(load32(m.memory[uint32(v5+i32(12)):]))
			m.fn16(t32, t33)
			t34 := int32(load32(m.memory[int64(uint32(v3))+72:]))
			store32(m.memory[int64(uint32(v5))+16:], uint32(t34))
			t35 := int64(load64(m.memory[int64(uint32(v3))+64:]))
			store64(m.memory[int64(uint32(v5))+8:], uint64(t35))
			t36 := int32(load32(m.memory[uint32(v6):]))
			t37 := int32(load32(m.memory[uint32(v7):]))
			m.fn1046(v3+i32(16), t36, t37, i32(1077249), i32(47), i32(1085790), i32(10))
			t38 := int32(load32(m.memory[int64(uint32(v3))+16:]))
			t39 := int32(load32(m.memory[int64(uint32(v3))+20:]))
			m.fn1041(v3+i32(64), t38, t39)
			t40 := int32(load32(m.memory[int64(uint32(v5))+20:]))
			t41 := int32(load32(m.memory[uint32(v5+i32(24)):]))
			m.fn134(t40, t41)
			t42 := int32(load32(m.memory[int64(uint32(v3))+72:]))
			store32(m.memory[int64(uint32(v5))+28:], uint32(t42))
			t43 := int64(load64(m.memory[int64(uint32(v3))+64:]))
			store64(m.memory[int64(uint32(v5))+20:], uint64(t43))
			t44 := int32(load32(m.memory[uint32(v6):]))
			t45 := int32(load32(m.memory[uint32(v7):]))
			m.fn1046(v3+i32(8), t44, t45, i32(1074680), i32(46), i32(1085800), i32(14))
			{
				t46 := int32(load32(m.memory[int64(uint32(v3))+8:]))
				v4 = t46
				if v4 == 0 {
					goto l11
				}
				t47 := int32(load32(m.memory[int64(uint32(v3))+12:]))
				m.fn197(v3+i32(64), v4, t47)
				t48 := int32(load32(m.memory[int64(uint32(v3))+68:]))
				t49 := int32(m.memory[int64(uint32(v3))+64])
				p50 := t48
				if t49 != 0 {
					p50 = i32(1)
				}
				v1 = p50
			}
		l11:
			store32(m.memory[int64(uint32(v5))+32:], uint32(v1))
			goto l3
		}
	}
	v5 = v3 + i32(80) + v4
	store64(m.memory[uint32(v5):], uint64(i64(1)))
	m.memory[uint32(v5+i32(36))] = byte(i32(0))
	store32(m.memory[uint32(v5+i32(32)):], uint32(i32(1)))
	store64(m.memory[uint32(v5+i32(16)):], uint64(i64(-0x100000000)))
	store64(m.memory[uint32(v5+i32(8)):], uint64(i64(0x100000000)))
	v4 = v4 + i32(40)
	goto l1
}
func (m *Module) fn1330(v0 int32) {
	t0 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	if t0 == i32(-1) {
		return
	}
	m.fn752(v0)
}
func (m *Module) fn1331(v0, v1, v2, v3, v4 int32) {
	var v5, v6 int32
	var v7 int64
	var v8, v9, v10, v11, v12, v13, v14, v15, v16 int32
	var v17 int64
	t0 := m.g0
	v5 = t0 - i32(272)
	m.g0 = v5
	{
		{
			t1 := int32(load32(m.memory[int64(uint32(v1))+36:]))
			v6 = t1
			if v6 == 0 {
				goto l0
			}
			t2 := int32(load32(m.memory[int64(uint32(v1))+40:]))
			t3 := m.fn15(v6+i32(8), t2, i32(1074680), i32(46))
			if t3 != 0 {
				goto l1
			}
		}
	l0:
		{
			t4 := m.fn847(v1, i32(1074726), i32(47), i32(1074842), i32(5))
			if t4 == 0 {
				goto l2
			}
			m.fn1333(v4, v3)
			m.fn1307(v5+i32(216), v1, v2)
			t5 := int64(load64(m.memory[int64(uint32(v5))+220:]))
			store64(m.memory[int64(uint32(v5))+184:], uint64(t5))
			t6 := int32(load32(m.memory[int64(uint32(v5))+228:]))
			store32(m.memory[int64(uint32(v5))+192:], uint32(t6))
			{
				t7 := int32(load32(m.memory[int64(uint32(v5))+216:]))
				v1 = t7
				if v1 == i32(-1) {
					goto l3
				}
				t8 := int64(load64(m.memory[int64(uint32(v5))+232:]))
				v7 = t8
				t9 := int32(load32(m.memory[int64(uint32(v5))+192:]))
				store32(m.memory[int64(uint32(v0))+12:], uint32(t9))
				t10 := int64(load64(m.memory[int64(uint32(v5))+184:]))
				store64(m.memory[int64(uint32(v0))+4:], uint64(t10))
				store64(m.memory[int64(uint32(v0))+16:], uint64(v7))
				store32(m.memory[uint32(v0):], uint32(v1))
				goto l4
			}
		l3:
			m.fn1271(v3, v5+i32(184))
		}
	l2:
		store32(m.memory[uint32(v0):], uint32(i32(-1)))
		goto l4
	l1:
		{
			t11 := m.fn847(v1, i32(1074680), i32(46), i32(1077161), i32(1))
			if t11 != 0 {
				goto l5
			}
			m.fn1333(v4, v3)
		}
	l5:
		{
			{
				{
					t12 := int32(load32(m.memory[int64(uint32(v1))+4:]))
					v6 = t12
					t13 := int32(load32(m.memory[int64(uint32(v1))+8:]))
					t14 := v6
					v8 = t13
					t15 := m.fn15(t14, v8, i32(1085393), i32(1))
					if t15 != 0 {
						t20 := int32(load32(m.memory[int64(uint32(v1))+16:]))
						t21 := v5 + i32(32)
						v4 = t20
						t22 := int32(load32(m.memory[int64(uint32(v1))+20:]))
						t23 := v4
						v8 = t22
						m.fn1046(t21, t23, v8, i32(1074680), i32(46), i32(1085479), i32(13))
						{
							{
								t24 := int32(load32(m.memory[int64(uint32(v5))+32:]))
								v6 = t24
								if v6 != 0 {
									goto l11
								}
								v6 = i32(1)
								goto l12
							}
						l11:
							t25 := int32(load32(m.memory[int64(uint32(v5))+36:]))
							m.fn1423(v5+i32(24), v6, t25)
							t26 := int32(m.memory[int64(uint32(v5))+25])
							t27 := int32(m.memory[int64(uint32(v5))+24])
							p28 := t26
							if t27&i32(1) != 0 {
								p28 = i32(1)
							}
							v6 = p28
						}
					l12:
						m.fn1510(v5+i32(216), v1, v2)
						t29 := int64(load64(m.memory[int64(uint32(v5))+220:]))
						store64(m.memory[int64(uint32(v5))+112:], uint64(t29))
						t30 := int64(load64(m.memory[int64(uint32(v5))+228:]))
						store64(m.memory[int64(uint32(v5))+120:], uint64(t30))
						t31 := int64(load64(m.memory[int64(uint32(v5))+236:]))
						store64(m.memory[int64(uint32(v5))+128:], uint64(t31))
						t32 := int32(load32(m.memory[int64(uint32(v5))+216:]))
						if t32 == 0 {
							t80 := int32(load32(m.memory[int64(uint32(v5))+120:]))
							t81 := v5
							v1 = t80
							store32(m.memory[int64(uint32(t81))+64:], uint32(v1))
							t82 := int64(load64(m.memory[int64(uint32(v5))+112:]))
							store64(m.memory[int64(uint32(v5))+56:], uint64(t82))
							t83 := int32(load32(m.memory[int64(uint32(v5))+132:]))
							store32(m.memory[int64(uint32(v5))+80:], uint32(t83))
							t84 := int64(load64(m.memory[int64(uint32(v5))+124:]))
							store64(m.memory[int64(uint32(v5))+72:], uint64(t84))
							{
								{
									t85 := int32(load32(m.memory[int64(uint32(v5))+60:]))
									t86 := m.fn23(t85, v1)
									if t86 != 0 {
										m.fn1271(v3, v5+i32(72))
										store32(m.memory[uint32(v0):], uint32(i32(-1)))
										m.fn894(v5 + i32(56))
										goto l4
									}
									t87 := int64(load64(m.memory[int64(uint32(v5))+56:]))
									store64(m.memory[int64(uint32(v5))+88:], uint64(t87))
									t88 := int32(load32(m.memory[int64(uint32(v5))+64:]))
									t89 := v5
									v9 = t88
									store32(m.memory[int64(uint32(t89))+96:], uint32(v9))
									t90 := int32(load32(m.memory[int64(uint32(v5))+92:]))
									v10 = t90
									t91 := v5 + i32(216)
									t92 := v4
									t93 := v8
									v12 = v2 + i32(200)
									t94 := int32(load32(m.memory[uint32(v12):]))
									m.fn1513(t91, t92, t93, t94)
									t95 := int32(load32(m.memory[int64(uint32(v5))+220:]))
									v1 = t95
									{
										t96 := int32(load32(m.memory[int64(uint32(v5))+216:]))
										v11 = t96
										if v11 == i32(-1) {
											t99 := m.fn1188(v1)
											m.fn1368(v10, v9, t99)
											m.fn45(v5+i32(100), v10, v9)
											t100 := int32(load32(m.memory[uint32(v12):]))
											v10 = t100
											t101 := int32(load32(m.memory[int64(uint32(v10))+8:]))
											var p102 int32
											if t101 == i32(-1) {
												p102 = 1
											}
											v13 = p102
											if v13 != 0 {
												goto l26
											}
											t103 := v10
											v1 = v6 & i32(255)
											p104 := i32(1)
											if uint32(v1) > uint32(i32(1)) {
												p104 = v1
											}
											v1 = p104 + i32(-1)
											p105 := i32(9)
											if uint32(v1) < uint32(i32(9)) {
												p105 = v1
											}
											v9 = p105
											v14 = t103 + v9*i32(40)
											t106 := int32(m.memory[int64(uint32(v14))+36])
											v15 = t106
											if v15 == 0 {
												goto l26
											}
											m.fn1046(v5+i32(16), v4, v8, i32(1074680), i32(46), i32(1085344), i32(14))
											t107 := int32(load32(m.memory[int64(uint32(v5))+16:]))
											t108 := int32(load32(m.memory[int64(uint32(v5))+20:]))
											t109 := m.fn848(t107, t108, i32(1071691), i32(4))
											if t109 != 0 {
												goto l26
											}
											t110 := int32(load32(m.memory[int64(uint32(v2))+96:]))
											if t110 != 0 {
												m.fn1326(i32(1085360))
												panic("unreachable")
											}
											store32(m.memory[int64(uint32(v2))+96:], uint32(i32(-1)))
											m.fn1046(v5+i32(8), v4, v8, i32(1074680), i32(46), i32(1085376), i32(17))
											t111 := int32(load32(m.memory[int64(uint32(v5))+8:]))
											t112 := int32(load32(m.memory[int64(uint32(v5))+12:]))
											t113 := m.fn848(t111, t112, i32(1071691), i32(4))
											v1 = t113
											m.fn1046(v5, v4, v8, i32(1074680), i32(46), i32(1085180), i32(11))
											t114 := int32(load32(m.memory[uint32(v5):]))
											t115 := int32(load32(m.memory[int64(uint32(v5))+4:]))
											m.fn1514(v5+i32(112), t114, t115)
											v11 = v2 + i32(184)
											v12 = v2 + i32(104)
											{
												{
													t116 := int64(load64(m.memory[int64(uint32(v5))+112:]))
													if t116 != i64(1) {
														goto l28
													}
													t117 := int64(load64(m.memory[int64(uint32(v5))+120:]))
													v7 = t117
													goto l29
												}
											l28:
												{
													t118 := int32(m.memory[uint32(v11+v9)])
													if (v1|(t118^i32(-1)))&i32(1) == 0 {
														goto l30
													}
													t119 := int64(load64(m.memory[uint32(v14):]))
													v7 = t119
													goto l29
												}
											l30:
												t120 := int64(load64(m.memory[uint32(v12+v9<<3):]))
												v7 = t120 + i64(1)
												p121 := v7
												if v7 == 0 {
													p121 = i64(-1)
												}
												v7 = p121
											}
										l29:
											p122 := v10
											if v13 != 0 {
												p122 = i32(0)
											}
											v13 = p122
											m.memory[uint32(v11+v9)] = byte(i32(1))
											store64(m.memory[uint32(v12+v9<<3):], uint64(v7))
											store32(m.memory[int64(uint32(v5))+224:], uint32(v9+i32(1)))
											store32(m.memory[int64(uint32(v5))+220:], uint32(v2+i32(194)))
											store32(m.memory[int64(uint32(v5))+216:], uint32(v11))
										l32:
											{
												t123 := m.fn1370(v5 + i32(216))
												v1 = t123
												if v1 == 0 {
													m.fn1515(v5+i32(216), v14, v9)
													t124 := int32(load32(m.memory[int64(uint32(v5))+224:]))
													v1 = t124
													if v1 == 0 {
														goto l33
													}
													store32(m.memory[int64(uint32(v5))+256:], uint32(i32(0)))
													store64(m.memory[int64(uint32(v5))+248:], uint64(i64(0x100000000)))
													v4 = v1 * i32(12)
													t125 := int32(m.memory[int64(uint32(v5))+228])
													v16 = t125
													t126 := int32(load32(m.memory[int64(uint32(v5))+220:]))
													v1 = t126
												l40:
													{
														{
															if v4 == 0 {
																m.fn800(v5+i32(260), v15, v7)
																t140 := int32(load32(m.memory[int64(uint32(v5))+252:]))
																v8 = t140
																t141 := int32(load32(m.memory[int64(uint32(v5))+256:]))
																t142 := int32(load32(m.memory[int64(uint32(v5))+264:]))
																t143 := v8
																v1 = t142
																t144 := int32(load32(m.memory[int64(uint32(v5))+268:]))
																t145 := m.fn191(t143, t141, v1, t144)
																v4 = t145
																t146 := int32(load32(m.memory[int64(uint32(v5))+260:]))
																m.fn16(t146, v1)
																t147 := int32(load32(m.memory[int64(uint32(v5))+248:]))
																v1 = t147
																{
																	if v4 != 0 {
																		m.fn16(v1, v8)
																		goto l33
																	}
																	t148 := int64(load64(m.memory[int64(uint32(v5))+252:]))
																	v17 = t148
																	goto l39
																}
															}
															t127 := int32(load32(m.memory[uint32(v1):]))
															if t127 != i32(-1) {
																goto l35
															}
															t128 := int32(m.memory[uint32(v1+i32(4))])
															v8 = t128
															p129 := i32(9)
															if uint32(v8) < uint32(i32(9)) {
																p129 = v8
															}
															v8 = p129
															v9 = i32(1)
															{
																if v16&i32(1) != 0 {
																	goto l36
																}
																t130 := int32(m.memory[int64(uint32(v10+v8*i32(40)))+36])
																v9 = t130
															}
														l36:
															t131 := int32(m.memory[uint32(v11+v8)])
															t133 := v5 + i32(260)
															t134 := v9
															p132 := v13 + v8*i32(40)
															if t131 != 0 {
																p132 = v12 + v8<<3
															}
															t135 := int64(load64(m.memory[uint32(p132):]))
															m.fn804(t133, t134, t135)
															t136 := int32(load32(m.memory[int64(uint32(v5))+264:]))
															t137 := v5 + i32(248)
															v8 = t136
															t138 := int32(load32(m.memory[int64(uint32(v5))+268:]))
															m.fn75(t137, v8, t138)
															t139 := int32(load32(m.memory[int64(uint32(v5))+260:]))
															m.fn16(t139, v8)
															goto l37
														}
													l35:
														t149 := int32(load32(m.memory[uint32(v1+i32(4)):]))
														t150 := int32(load32(m.memory[uint32(v1+i32(8)):]))
														m.fn75(v5+i32(248), t149, t150)
													}
												l37:
													v1 = v1 + i32(12)
													v4 = v4 + i32(-12)
													goto l40
												}
												m.memory[uint32(v1)] = byte(i32(0))
												goto l32
											}
										}
										t97 := int64(load64(m.memory[int64(uint32(v5))+232:]))
										store64(m.memory[int64(uint32(v0))+16:], uint64(t97))
										t98 := int64(load64(m.memory[int64(uint32(v5))+224:]))
										store64(m.memory[int64(uint32(v0))+8:], uint64(t98))
										store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
										store32(m.memory[uint32(v0):], uint32(v11))
										m.fn894(v5 + i32(56))
										m.fn969(v5 + i32(72))
										goto l4
									}
								}
							l33:
								v1 = i32(-1)
							l39:
								m.fn763(v5 + i32(216))
								{
									if v1 == i32(-1) {
										goto l41
									}
									store64(m.memory[int64(uint32(v5))+220:], uint64(v17))
									store32(m.memory[int64(uint32(v5))+216:], uint32(v1))
									goto l42
								l41:
									t151 := int32(m.memory[int64(uint32(v14))+36])
									m.fn800(v5+i32(216), t151, v7)
								}
							l42:
								store32(m.memory[int64(uint32(v5))+252:], uint32(i32(25)))
								store32(m.memory[int64(uint32(v5))+248:], uint32(v5+i32(216)))
								m.fn73(v5+i32(260), i32(1070105), v5+i32(248))
								t152 := int32(load32(m.memory[int64(uint32(v5))+216:]))
								t153 := int32(load32(m.memory[int64(uint32(v5))+220:]))
								m.fn16(t152, t153)
								t154 := int32(load32(m.memory[int64(uint32(v2))+96:]))
								store32(m.memory[int64(uint32(v2))+96:], uint32(t154+i32(1)))
								t155 := int32(load32(m.memory[int64(uint32(v5))+260:]))
								if t155 == i32(-1) {
									goto l26
								}
								t156 := int32(load32(m.memory[int64(uint32(v5))+268:]))
								store32(m.memory[int64(uint32(v5))+228:], uint32(t156))
								t157 := int64(load64(m.memory[int64(uint32(v5))+260:]))
								store64(m.memory[int64(uint32(v5))+220:], uint64(t157))
								store32(m.memory[int64(uint32(v5))+232:], uint32(i32(0)))
								store32(m.memory[int64(uint32(v5))+216:], uint32(i32(3)))
								m.fn1163(v5+i32(88), v5+i32(216))
							}
						l26:
							t158 := int32(load32(m.memory[int64(uint32(v5))+96:]))
							store32(m.memory[int64(uint32(v5))+224:], uint32(t158))
							t159 := int64(load64(m.memory[int64(uint32(v5))+88:]))
							store64(m.memory[int64(uint32(v5))+216:], uint64(t159))
							m.memory[int64(uint32(v5))+240] = byte(v6)
							t160 := int64(load64(m.memory[int64(uint32(v5))+100:]))
							store64(m.memory[int64(uint32(v5))+228:], uint64(t160))
							t161 := int32(load32(m.memory[int64(uint32(v5))+108:]))
							store32(m.memory[int64(uint32(v5))+236:], uint32(t161))
							m.fn338(v3, v5+i32(216))
							m.fn1271(v3, v5+i32(72))
							store32(m.memory[uint32(v0):], uint32(i32(-1)))
							goto l4
						}
						t33 := int64(load64(m.memory[int64(uint32(v5))+128:]))
						store64(m.memory[int64(uint32(v0))+16:], uint64(t33))
						t34 := int64(load64(m.memory[int64(uint32(v5))+120:]))
						store64(m.memory[int64(uint32(v0))+8:], uint64(t34))
						t35 := int64(load64(m.memory[int64(uint32(v5))+112:]))
						store64(m.memory[uint32(v0):], uint64(t35))
						goto l4
					}
					t16 := m.fn15(v6, v8, i32(1077161), i32(1))
					if t16 != 0 {
						m.fn1510(v5+i32(216), v1, v2)
						t47 := int64(load64(m.memory[int64(uint32(v5))+220:]))
						store64(m.memory[int64(uint32(v5))+112:], uint64(t47))
						t48 := int64(load64(m.memory[int64(uint32(v5))+228:]))
						store64(m.memory[int64(uint32(v5))+120:], uint64(t48))
						t49 := int64(load64(m.memory[int64(uint32(v5))+236:]))
						store64(m.memory[int64(uint32(v5))+128:], uint64(t49))
						t50 := int32(load32(m.memory[int64(uint32(v5))+216:]))
						if t50 != 0 {
							t165 := int64(load64(m.memory[int64(uint32(v5))+128:]))
							store64(m.memory[int64(uint32(v0))+16:], uint64(t165))
							t166 := int64(load64(m.memory[int64(uint32(v5))+120:]))
							store64(m.memory[int64(uint32(v0))+8:], uint64(t166))
							t167 := int64(load64(m.memory[int64(uint32(v5))+112:]))
							store64(m.memory[uint32(v0):], uint64(t167))
							goto l4
						}
						t51 := int32(load32(m.memory[int64(uint32(v5))+120:]))
						store32(m.memory[int64(uint32(v5))+208:], uint32(t51))
						t52 := int64(load64(m.memory[int64(uint32(v5))+112:]))
						store64(m.memory[int64(uint32(v5))+200:], uint64(t52))
						t53 := int32(load32(m.memory[int64(uint32(v5))+132:]))
						store32(m.memory[int64(uint32(v5))+256:], uint32(t53))
						t54 := int64(load64(m.memory[int64(uint32(v5))+124:]))
						store64(m.memory[int64(uint32(v5))+248:], uint64(t54))
						t55 := int32(load32(m.memory[int64(uint32(v1))+16:]))
						t56 := int32(load32(m.memory[int64(uint32(v1))+20:]))
						m.fn1046(v5+i32(48), t55, t56, i32(1074680), i32(46), i32(1085052), i32(10))
						t57 := int32(load32(m.memory[int64(uint32(v5))+48:]))
						v1 = t57
						if v1 == 0 {
							goto l17
						}
						t58 := int32(load32(m.memory[int64(uint32(v2))+200:]))
						v2 = t58
						t59 := int32(load32(m.memory[int64(uint32(v5))+52:]))
						m.fn1328(v5+i32(260), i32(1077240), i32(9), v1, t59)
						m.fn27(v5 + i32(216))
						t60 := int32(load32(m.memory[int64(uint32(v5))+268:]))
						v6 = t60
						t61 := int32(load32(m.memory[int64(uint32(v5))+264:]))
						v1 = t61
						t62 := int32(load32(m.memory[int64(uint32(v5))+260:]))
						v9 = t62
						v10 = v2 + i32(404)
						v11 = v2 + i32(424)
					l23:
						m.fn31(v5+i32(112), v1, v6)
						{
							t63 := m.fn32(v5+i32(216), v5+i32(112))
							if t63 == 0 {
								goto l18
							}
							t64 := int32(load32(m.memory[int64(uint32(v2))+412:]))
							if t64 == 0 {
								goto l18
							}
							t65 := int64(load64(m.memory[int64(uint32(v2))+416:]))
							t66 := int64(load64(m.memory[uint32(v11):]))
							t67 := m.fn540(t65, t66, v1, v6)
							v7 = t67
							t68 := int32(load32(m.memory[int64(uint32(v2))+400:]))
							t69 := int32(load32(m.memory[uint32(v10):]))
							t70 := m.fn1512(t68, t69, v7, v5+i32(260))
							v6 = t70
							if v6 == 0 {
								goto l18
							}
							t71 := int32(load32(m.memory[uint32(v6+i32(-16)):]))
							t72 := v5 + i32(40)
							v8 = t71
							t73 := int32(load32(m.memory[uint32(v8+i32(16)):]))
							t74 := int32(load32(m.memory[uint32(v8+i32(20)):]))
							m.fn1046(t72, t73, t74, i32(1077249), i32(47), i32(1073713), i32(4))
							{
								t75 := int32(load32(m.memory[int64(uint32(v5))+40:]))
								v8 = t75
								if v8 == 0 {
									goto l19
								}
								t76 := int32(load32(m.memory[int64(uint32(v5))+44:]))
								t77 := m.fn1208(v8, t76)
								v8 = t77 & i32(255)
								if v8 != i32(2) {
									goto l20
								}
							}
						l19:
							m.fn225(v5+i32(112), v6+i32(-12))
							t78 := int32(load32(m.memory[int64(uint32(v5))+112:]))
							v8 = t78
							if v8 != i32(-1) {
								t79 := int64(load64(m.memory[int64(uint32(v5))+116:]))
								v7 = t79
								m.fn16(v9, v1)
								store64(m.memory[int64(uint32(v5))+264:], uint64(v7))
								store32(m.memory[int64(uint32(v5))+260:], uint32(v8))
								v6 = int32(int64(uint64(v7) >> 32))
								v1 = int32(v7)
								v9 = v8
								goto l23
							}
						}
					l18:
						v8 = i32(2)
					l20:
						m.fn38(v5 + i32(216))
						m.fn16(v9, v1)
						if v8 == i32(2) {
							goto l17
						}
						m.fn1445(v4, v8&i32(1), v5+i32(200), v3)
						goto l22
					}
					t17 := m.fn15(v6, v8, i32(1081789), i32(4))
					if t17 != 0 {
						m.fn1511(v5+i32(216), v1, v2, i32(0), i32(0), v5, i32(8), i32(0))
						t41 := int64(load64(m.memory[int64(uint32(v5))+220:]))
						store64(m.memory[int64(uint32(v5))+136:], uint64(t41))
						t42 := int32(load32(m.memory[int64(uint32(v5))+228:]))
						store32(m.memory[int64(uint32(v5))+144:], uint32(t42))
						{
							t43 := int32(load32(m.memory[int64(uint32(v5))+216:]))
							v1 = t43
							if v1 == i32(-1) {
								m.fn1271(v3, v5+i32(136))
								store32(m.memory[uint32(v0):], uint32(i32(-1)))
								goto l4
							}
							t44 := int64(load64(m.memory[int64(uint32(v5))+232:]))
							v7 = t44
							t45 := int32(load32(m.memory[int64(uint32(v5))+144:]))
							store32(m.memory[int64(uint32(v0))+12:], uint32(t45))
							t46 := int64(load64(m.memory[int64(uint32(v5))+136:]))
							store64(m.memory[int64(uint32(v0))+4:], uint64(t46))
							store64(m.memory[int64(uint32(v0))+16:], uint64(v7))
							store32(m.memory[uint32(v0):], uint32(v1))
							goto l4
						}
					}
					t18 := m.fn15(v6, v8, i32(1080919), i32(7))
					if t18 != 0 {
						goto l9
					}
					t19 := m.fn15(v6, v8, i32(1085394), i32(10))
					if t19 == 0 {
						t36 := m.fn15(v6, v8, i32(1085404), i32(11))
						if t36 != 0 {
							goto l9
						}
						t37 := m.fn15(v6, v8, i32(1085415), i32(16))
						if t37 != 0 {
							goto l14
						}
						t38 := m.fn15(v6, v8, i32(1085431), i32(18))
						if t38 != 0 {
							goto l14
						}
						t39 := m.fn15(v6, v8, i32(1085449), i32(12))
						if t39 != 0 {
							goto l14
						}
						t40 := m.fn15(v6, v8, i32(1085461), i32(18))
						if t40 != 0 {
							goto l14
						}
						store32(m.memory[uint32(v0):], uint32(i32(-1)))
						goto l4
					}
					goto l9
				}
			l17:
				m.fn1333(v4, v3)
				t162 := int32(load32(m.memory[int64(uint32(v5))+208:]))
				store32(m.memory[int64(uint32(v5))+228:], uint32(t162))
				t163 := int64(load64(m.memory[int64(uint32(v5))+200:]))
				store64(m.memory[int64(uint32(v5))+220:], uint64(t163))
				store32(m.memory[int64(uint32(v5))+216:], uint32(i32(-0x80000000)))
				m.fn338(v3, v5+i32(216))
			}
		l22:
			{
				t164 := int32(load32(m.memory[int64(uint32(v5))+256:]))
				if t164 == 0 {
					goto l43
				}
				m.fn1333(v4, v3)
				m.fn1271(v3, v5+i32(248))
				goto l44
			}
		l43:
			m.fn969(v5 + i32(248))
		l44:
			store32(m.memory[uint32(v0):], uint32(i32(-1)))
			goto l4
		l14:
			m.fn1306(v5+i32(216), v1, v2)
			t168 := int64(load64(m.memory[int64(uint32(v5))+220:]))
			store64(m.memory[int64(uint32(v5))+168:], uint64(t168))
			t169 := int32(load32(m.memory[int64(uint32(v5))+228:]))
			store32(m.memory[int64(uint32(v5))+176:], uint32(t169))
			{
				t170 := int32(load32(m.memory[int64(uint32(v5))+216:]))
				v1 = t170
				if v1 == i32(-1) {
					m.fn1271(v3, v5+i32(168))
					store32(m.memory[uint32(v0):], uint32(i32(-1)))
					goto l4
				}
				t171 := int64(load64(m.memory[int64(uint32(v5))+232:]))
				v7 = t171
				t172 := int32(load32(m.memory[int64(uint32(v5))+176:]))
				store32(m.memory[int64(uint32(v0))+12:], uint32(t172))
				t173 := int64(load64(m.memory[int64(uint32(v5))+168:]))
				store64(m.memory[int64(uint32(v0))+4:], uint64(t173))
				store64(m.memory[int64(uint32(v0))+16:], uint64(v7))
				store32(m.memory[uint32(v0):], uint32(v1))
				goto l4
			}
		}
	l9:
		m.fn1306(v5+i32(216), v1, v2)
		t174 := int64(load64(m.memory[int64(uint32(v5))+220:]))
		store64(m.memory[int64(uint32(v5))+152:], uint64(t174))
		t175 := int32(load32(m.memory[int64(uint32(v5))+228:]))
		store32(m.memory[int64(uint32(v5))+160:], uint32(t175))
		{
			t176 := int32(load32(m.memory[int64(uint32(v5))+216:]))
			v1 = t176
			if v1 == i32(-1) {
				goto l46
			}
			t177 := int64(load64(m.memory[int64(uint32(v5))+232:]))
			v7 = t177
			t178 := int32(load32(m.memory[int64(uint32(v5))+160:]))
			store32(m.memory[int64(uint32(v0))+12:], uint32(t178))
			t179 := int64(load64(m.memory[int64(uint32(v5))+152:]))
			store64(m.memory[int64(uint32(v0))+4:], uint64(t179))
			store64(m.memory[int64(uint32(v0))+16:], uint64(v7))
			store32(m.memory[uint32(v0):], uint32(v1))
			goto l4
		}
	l46:
		m.fn1271(v3, v5+i32(152))
		store32(m.memory[uint32(v0):], uint32(i32(-1)))
	}
l4:
	m.g0 = v5 + i32(272)
}
func (m *Module) fn1332(v0 int32) {
	t0 := int32(load32(m.memory[uint32(v0):]))
	switch t0 {
	default:
		m.fn78(v0 + i32(4))
		fallthrough
	case 0:
		return
	case 1:
		m.fn969(v0 + i32(4))
	}
}
func (m *Module) fn1333(v0, v1 int32) {
	var v2, v3, v4, v5, v6, v7, v8, v9 int32
	t0 := m.g0
	v2 = t0 - i32(64)
	m.g0 = v2
	t1 := int32(load32(m.memory[uint32(v0):]))
	v3 = t1
	v4 = i32(0)
	store32(m.memory[uint32(v0):], uint32(i32(0)))
	v5 = v0 + i32(4)
	{
		switch v3 {
		default:
			goto l0
		case 2:
			t2 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			v6 = t2
			t3 := int32(load32(m.memory[int64(uint32(v0))+12:]))
			t4 := v6
			v7 = t3
			v8 = v7 * i32(12)
			v0 = t4 + v8
			v3 = v6
		l4:
			{
				if v8 == 0 {
					goto l3
				}
				t5 := int32(load32(m.memory[int64(uint32(v3))+4:]))
				t6 := int32(load32(m.memory[int64(uint32(v3))+8:]))
				m.fn46(v2+i32(24), t5, t6)
				t7 := int32(load32(m.memory[int64(uint32(v2))+28:]))
				if t7 != 0 {
					goto l3
				}
				v8 = v8 + i32(-12)
				v4 = v4 + i32(1)
				v3 = v3 + i32(12)
				goto l4
			}
		l3:
			v3 = v7 * i32(-12)
			v9 = v7
		l6:
			{
				if v3 == 0 {
					goto l5
				}
				t8 := int32(load32(m.memory[uint32(v0+i32(-8)):]))
				t9 := int32(load32(m.memory[uint32(v0+i32(-4)):]))
				m.fn46(v2+i32(16), t8, t9)
				v3 = v3 + i32(12)
				v9 = v9 + i32(-1)
				v0 = v0 + i32(-12)
				t10 := int32(load32(m.memory[int64(uint32(v2))+20:]))
				if t10 == 0 {
					goto l6
				}
			}
			if v8 != 0 {
				goto l7
			}
			goto l5
		case 1:
			t11 := int32(load32(m.memory[int64(uint32(v0))+12:]))
			if t11 == 0 {
				m.fn969(v5)
				goto l0
			}
			store32(m.memory[int64(uint32(v2))+32:], uint32(i32(-0x7ffffffd)))
			t12 := int32(load32(m.memory[int64(uint32(v5))+8:]))
			store32(m.memory[int64(uint32(v2))+44:], uint32(t12))
			t13 := int64(load64(m.memory[uint32(v5):]))
			store64(m.memory[int64(uint32(v2))+36:], uint64(t13))
			m.fn338(v1, v2+i32(32))
			goto l0
		}
	l7:
		m.memory[int64(uint32(v2))+40] = byte(i32(0))
		store32(m.memory[int64(uint32(v2))+36:], uint32(v9))
		store32(m.memory[int64(uint32(v2))+32:], uint32(v4))
		m.fn194(v2+i32(8), v2+i32(32), v6, v7, i32(1075080))
		t14 := int32(load32(m.memory[int64(uint32(v2))+8:]))
		t15 := int32(load32(m.memory[int64(uint32(v2))+12:]))
		m.fn77(v2+i32(32)|i32(4), t14, t15, i32(1108166), i32(1))
		store32(m.memory[int64(uint32(v2))+32:], uint32(i32(-0x7ffffffc)))
		store32(m.memory[int64(uint32(v2))+48:], uint32(i32(-1)))
		m.fn338(v1, v2+i32(32))
	}
l5:
	m.fn78(v5)
l0:
	m.g0 = v2 + i32(64)
}
func (m *Module) fn1334(v0 int32) {
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
					v7 = v6 - int32(int64(bits.TrailingZeros64(uint64(v5))))<<1&i32(240)
					t6 := int32(load32(m.memory[uint32(v7+i32(-16)):]))
					t7 := int32(load32(m.memory[uint32(v7+i32(-12)):]))
					m.fn16(t6, t7)
					v4 = v4 + i32(-1)
					v5 = (v5 + i64(-1)) & v5
					goto l4
				}
				v6 = v6 + i32(-128)
				t5 := int64(load64(m.memory[uint32(v0):]))
				v5 = (t5 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
				v0 = v0 + i32(8)
				goto l3
			}
		}
	l1:
		m.fn39(v1+i32(4), i32(16), i32(8), v2+i32(1))
		t8 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		t9 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t10 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		m.fn40(v3-t8, t9, t10)
	}
l0:
	m.g0 = v1 + i32(16)
}
func (m *Module) fn1335(v0, v1, v2, v3, v4 int32) {
	var v5, v6, v7, v8 int32
	var v9, v10 int64
	var v11, v12, v13, v14, v15, v16, v17, v18 int32
	var v19 float64
	var v20, v21 int64
	t0 := m.g0
	v5 = t0 - i32(352)
	m.g0 = v5
	t1 := int32(load32(m.memory[int64(uint32(v1))+32:]))
	v6 = t1
	t2 := int32(load32(m.memory[int64(uint32(v1))+28:]))
	t3 := v5
	v1 = t2
	store32(m.memory[int64(uint32(t3))+180:], uint32(v1))
	store32(m.memory[int64(uint32(v5))+184:], uint32(v1+v6*i32(44)))
	v7 = v3 + i32(24)
l1:
	{
		t4 := m.fn904(v5 + i32(180))
		v1 = t4
		if v1 == 0 {
			store32(m.memory[uint32(v0):], uint32(i32(-1)))
			goto l5
		}
		t5 := int32(load32(m.memory[int64(uint32(v1))+36:]))
		v6 = t5
		if v6 == 0 {
			goto l1
		}
		t6 := int32(load32(m.memory[int64(uint32(v1))+40:]))
		t7 := m.fn1337(v6+i32(8), t6, i32(1074726), i32(47))
		if t7 != 0 {
			goto l1
		}
		t8 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		v6 = t8
		t9 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		t10 := v6
		v8 = t9
		t11 := m.fn15(t10, v8, i32(1085621), i32(17))
		if t11 != 0 {
			t53 := int32(load32(m.memory[int64(uint32(v3))+84:]))
			v6 = t53
			m.fn1335(v5+i32(240), v1, v2, v3, i32(0))
			{
				t54 := int32(load32(m.memory[int64(uint32(v5))+240:]))
				v1 = t54
				if v1 == i32(-1) {
					if v4 == 0 {
						goto l1
					}
					t58 := int32(load32(m.memory[int64(uint32(v3))+80:]))
					if t58 != 0 {
						goto l1
					}
					t59 := int32(load32(m.memory[int64(uint32(v3))+84:]))
					store32(m.memory[int64(uint32(v3))+80:], uint32(t59-v6))
					goto l1
				}
				t55 := int32(load32(m.memory[int64(uint32(v5))+260:]))
				store32(m.memory[int64(uint32(v0))+20:], uint32(t55))
				t56 := int64(load64(m.memory[int64(uint32(v5))+252:]))
				store64(m.memory[int64(uint32(v0))+12:], uint64(t56))
				t57 := int64(load64(m.memory[int64(uint32(v5))+244:]))
				store64(m.memory[int64(uint32(v0))+4:], uint64(t57))
				store32(m.memory[uint32(v0):], uint32(v1))
				goto l5
			}
		}
		{
			t12 := m.fn15(v6, v8, i32(1085638), i32(10))
			if t12 != 0 {
				goto l3
			}
			t13 := m.fn15(v6, v8, i32(1085648), i32(15))
			if t13 == 0 {
				t18 := m.fn15(v6, v8, i32(1085663), i32(9))
				if t18 == 0 {
					goto l1
				}
				t19 := int32(load32(m.memory[uint32(v1+i32(16)):]))
				t20 := int32(load32(m.memory[uint32(v1+i32(20)):]))
				m.fn1046(v5+i32(168), t19, t20, i32(1074726), i32(47), i32(1085672), i32(20))
				{
					{
						t21 := int32(load32(m.memory[int64(uint32(v5))+168:]))
						v6 = t21
						if v6 != 0 {
							goto l6
						}
						v9 = i64(1)
						goto l7
					}
				l6:
					t22 := int32(load32(m.memory[int64(uint32(v5))+172:]))
					m.fn1190(v5+i32(240), v6, t22)
					t23 := int64(load64(m.memory[int64(uint32(v5))+248:]))
					v10 = t23
					p24 := i64(1)
					if uint64(v10) > uint64(i64(1)) {
						p24 = v10
					}
					t25 := int32(m.memory[int64(uint32(v5))+240])
					p26 := p24
					if t25 != 0 {
						p26 = i64(1)
					}
					v9 = p26
				}
			l7:
				t27 := int32(load32(m.memory[int64(uint32(v1))+32:]))
				v11 = t27
				v12 = v11 * i32(44)
				t28 := int32(load32(m.memory[int64(uint32(v1))+28:]))
				v13 = t28
				v6 = i32(0)
			l11:
				if v12 == v6 {
					t40 := int64(load64(m.memory[int64(uint32(v3))+16:]))
					t41 := v3
					v10 = t40
					v9 = v10 + v9
					p42 := v9
					if uint64(v9) < uint64(v10) {
						p42 = i64(-1)
					}
					store64(m.memory[int64(uint32(t41))+16:], uint64(p42))
					goto l1
				}
				{
					v8 = v13 + v6
					t29 := int32(load32(m.memory[uint32(v8):]))
					if t29 == i32(-1) {
						goto l9
					}
					t30 := m.fn847(v8, i32(1074726), i32(47), i32(1074773), i32(18))
					if t30 != 0 {
						goto l10
					}
					t31 := m.fn847(v8, i32(1074726), i32(47), i32(1071943), i32(10))
					if t31 == 0 {
						goto l9
					}
					t32 := int32(load32(m.memory[uint32(v8+i32(16)):]))
					t33 := v5 + i32(160)
					v14 = t32
					t34 := int32(load32(m.memory[uint32(v8+i32(20)):]))
					t35 := v14
					v15 = t34
					m.fn1046(t33, t35, v15, i32(1074726), i32(47), i32(0x106667), i32(22))
					t36 := int32(load32(m.memory[int64(uint32(v5))+160:]))
					if t36 != 0 {
						goto l10
					}
					m.fn1046(v5+i32(152), v14, v15, i32(1074726), i32(47), i32(1074813), i32(19))
					t37 := int32(load32(m.memory[int64(uint32(v5))+152:]))
					if t37 != 0 {
						goto l10
					}
					m.fn1046(v5+i32(144), v14, v15, i32(1074169), i32(48), i32(1074832), i32(10))
					t38 := int32(load32(m.memory[int64(uint32(v5))+144:]))
					if t38 != 0 {
						goto l10
					}
					t39 := int32(load32(m.memory[uint32(v8+i32(32)):]))
					if t39 != 0 {
						goto l10
					}
				}
			l9:
				v6 = v6 + i32(44)
				goto l11
			l10:
				t43 := int64(load64(m.memory[int64(uint32(v3))+16:]))
				v10 = t43
				if v10 == 0 {
					goto l12
				}
				m.fn1503(v5+i32(240), v3, v10)
				{
					t44 := int32(load32(m.memory[int64(uint32(v5))+240:]))
					v8 = t44
					if v8 == i32(-1) {
						t49 := int64(load64(m.memory[int64(uint32(v3))+16:]))
						v10 = t49
					l16:
						{
							if v10 == 0 {
								store64(m.memory[int64(uint32(v3))+16:], uint64(i64(0)))
								t51 := int32(load32(m.memory[int64(uint32(v1))+32:]))
								v11 = t51
								t52 := int32(load32(m.memory[int64(uint32(v1))+28:]))
								v13 = t52
								goto l12
							}
							m.fn1166(v7)
							t50 := int32(load32(m.memory[int64(uint32(v3))+84:]))
							store32(m.memory[int64(uint32(v3))+84:], uint32(t50+i32(1)))
							v10 = v10 + i64(-1)
							goto l16
						}
					}
					t45 := int64(load64(m.memory[int64(uint32(v5))+256:]))
					v10 = t45
					t46 := int32(load32(m.memory[int64(uint32(v5))+252:]))
					v12 = t46
					t47 := int32(load32(m.memory[int64(uint32(v5))+248:]))
					v14 = t47
					t48 := int32(load32(m.memory[int64(uint32(v5))+244:]))
					v16 = t48
					goto l14
				}
			}
		}
	l3:
		m.fn1335(v5+i32(240), v1, v2, v3, i32(0))
		t14 := int32(load32(m.memory[int64(uint32(v5))+240:]))
		v1 = t14
		if v1 == i32(-1) {
			goto l1
		}
		t15 := int32(load32(m.memory[int64(uint32(v5))+260:]))
		store32(m.memory[int64(uint32(v0))+20:], uint32(t15))
		t16 := int64(load64(m.memory[int64(uint32(v5))+252:]))
		store64(m.memory[int64(uint32(v0))+12:], uint64(t16))
		t17 := int64(load64(m.memory[int64(uint32(v5))+244:]))
		store64(m.memory[int64(uint32(v0))+4:], uint64(t17))
		store32(m.memory[uint32(v0):], uint32(v1))
		goto l5
	}
l12:
	store32(m.memory[int64(uint32(v5))+228:], uint32(i32(0)))
	store64(m.memory[int64(uint32(v5))+220:], uint64(i64(0x800000000)))
	store32(m.memory[int64(uint32(v5))+232:], uint32(v13))
	store32(m.memory[int64(uint32(v5))+236:], uint32(v13+v11*i32(44)))
l22:
	{
		{
			{
				{
					t60 := m.fn904(v5 + i32(232))
					v1 = t60
					if v1 == 0 {
						t149 := int32(load32(m.memory[int64(uint32(v5))+228:]))
						t150 := v5
						v1 = t149
						store32(m.memory[int64(uint32(t150))+196:], uint32(v1))
						t151 := int32(load32(m.memory[int64(uint32(v5))+224:]))
						t152 := v5
						v11 = t151
						store32(m.memory[int64(uint32(t152))+192:], uint32(v11))
						t153 := int32(load32(m.memory[int64(uint32(v5))+220:]))
						store32(m.memory[int64(uint32(v5))+188:], uint32(t153))
						t154 := v5 + i32(240)
						t155 := v3
						t156 := v9
						var p157 int32
						if v9 != i64(0) {
							p157 = 1
						}
						m.fn1503(t154, t155, t156-int64(uint32(p157)))
						{
							{
								t158 := int32(load32(m.memory[int64(uint32(v5))+240:]))
								v8 = t158
								if v8 != i32(-1) {
									goto l61
								}
								v13 = v11 + v1*i32(40)
								t159 := int64(load64(m.memory[int64(uint32(v3))+8:]))
								v10 = t159
								v6 = v11
								{
								l63:
									{
										v1 = v6
										if v1 == v13 {
											v20 = i64(0)
										l66:
											{
												if v20 == v9 {
													m.fn1505(v5 + i32(188))
													goto l1
												}
												m.fn1166(v7)
												t178 := int32(load32(m.memory[int64(uint32(v3))+84:]))
												store32(m.memory[int64(uint32(v3))+84:], uint32(t178+i32(1)))
												v20 = v20 + i64(1)
												v10 = i64(0)
												v6 = v11
											l72:
												store64(m.memory[int64(uint32(v5))+280:], uint64(v10))
											l74:
												if v6 == v13 {
													goto l66
												}
												v15 = v6 + i32(40)
												{
													{
														t179 := int32(load32(m.memory[uint32(v6):]))
														if t179 != i32(1) {
															m.fn1506(v5+i32(240), v3, v5+i32(280))
															t181 := int32(load32(m.memory[int64(uint32(v5))+240:]))
															v8 = t181
															if v8 != i32(-1) {
																goto l61
															}
															t182 := int64(load64(m.memory[int64(uint32(v6))+8:]))
															t183 := v5 + i32(240)
															t184 := v3
															v10 = t182
															m.fn1503(t183, t184, v10)
															t185 := int32(load32(m.memory[int64(uint32(v5))+240:]))
															v8 = t185
															if v8 != i32(-1) {
																goto l61
															}
														l71:
															if v10 == 0 {
																goto l70
															}
															_ = m.fn1260(v7)
															v10 = v10 + i64(-1)
															goto l71
														}
														t180 := int32(load32(m.memory[int64(uint32(v6))+20:]))
														if t180 == 0 {
															t187 := int32(load32(m.memory[int64(uint32(v6))+4:]))
															if t187 != i32(1) {
																goto l69
															}
															t188 := int32(load32(m.memory[int64(uint32(v6))+8:]))
															if t188 != i32(1) {
																goto l69
															}
															t189 := int64(load64(m.memory[int64(uint32(v5))+280:]))
															v10 = t189
															t190 := int64(load64(m.memory[int64(uint32(v6))+24:]))
															v21 = v10 + t190
															p191 := v21
															if uint64(v21) < uint64(v10) {
																p191 = i64(-1)
															}
															v10 = p191
															v6 = v15
															goto l72
														}
														goto l69
													}
												l69:
													m.fn1506(v5+i32(240), v3, v5+i32(280))
													t192 := int32(load32(m.memory[int64(uint32(v5))+240:]))
													v8 = t192
													if v8 != i32(-1) {
														goto l61
													}
													t193 := int64(load64(m.memory[int64(uint32(v6))+24:]))
													t194 := v5
													v10 = t193
													t195 := int32(load32(m.memory[int64(uint32(v6))+4:]))
													t196 := v10
													v1 = t195
													m.fn1853(t194, t196, i64(0), int64(uint32(v1)), i64(0))
													t197 := int64(load64(m.memory[uint32(v5):]))
													t198 := int64(load64(m.memory[int64(uint32(v5))+8:]))
													t200 := v5 + i32(240)
													t201 := v3
													p199 := t197
													if t198 != i64(0) {
														p199 = i64(-1)
													}
													m.fn1503(t200, t201, p199)
													t202 := int32(load32(m.memory[int64(uint32(v5))+240:]))
													v8 = t202
													if v8 != i32(-1) {
														goto l61
													}
													v12 = v6 + i32(12)
													p203 := i32(1)
													if uint32(v1) > uint32(i32(1)) {
														p203 = v1
													}
													v14 = p203
												l73:
													{
														if v10 == 0 {
															goto l70
														}
														m.fn1507(v5+i32(308), v12)
														store32(m.memory[int64(uint32(v5))+212:], uint32(v14))
														t204 := int64(load64(m.memory[int64(uint32(v5))+308:]))
														store64(m.memory[int64(uint32(v5))+200:], uint64(t204))
														t205 := int32(load32(m.memory[int64(uint32(v5))+316:]))
														store32(m.memory[int64(uint32(v5))+208:], uint32(t205))
														t206 := int32(load32(m.memory[int64(uint32(v6))+8:]))
														t207 := v5
														v1 = t206
														p208 := i32(1)
														if uint32(v1) > uint32(i32(1)) {
															p208 = v1
														}
														store32(m.memory[int64(uint32(t207))+216:], uint32(p208))
														m.fn1167(v5+i32(240), v7, v5+i32(200))
														t209 := int32(load32(m.memory[int64(uint32(v5))+240:]))
														v8 = t209
														if v8 != i32(-1) {
															goto l61
														}
														v10 = v10 + i64(-1)
														goto l73
													}
												}
											l70:
												v6 = v15
												goto l74
											}
										}
										v6 = v1 + i32(40)
										t160 := int32(load32(m.memory[uint32(v1):]))
										if t160 == 0 {
											goto l63
										}
										t161 := int64(load64(m.memory[int64(uint32(v1))+24:]))
										m.fn1853(v5+i32(16), v9, i64(0), t161, i64(0))
										t162 := int64(load64(m.memory[int64(uint32(v1))+32:]))
										t163 := int64(load64(m.memory[int64(uint32(v5))+16:]))
										t164 := v5 + i32(32)
										v20 = t163
										t165 := v20
										var p166 int32
										if v20 != i64(0) {
											p166 = 1
										}
										t167 := int64(load64(m.memory[int64(uint32(v5))+24:]))
										p168 := t165 - int64(uint32(p166))
										if t167 != i64(0) {
											p168 = i64(-2)
										}
										m.fn1853(t164, t162, i64(0), p168, i64(0))
										t169 := int64(load64(m.memory[int64(uint32(v5))+32:]))
										t170 := v3
										v20 = v10 + t169
										p171 := v20
										if uint64(v20) < uint64(v10) {
											p171 = i64(-1)
										}
										v10 = p171
										t172 := int64(load64(m.memory[int64(uint32(v5))+40:]))
										p173 := v10
										if t172 != i64(0) {
											p173 = i64(-1)
										}
										v20 = p173
										store64(m.memory[int64(uint32(t170))+8:], uint64(v20))
										if uint64(v20) < uint64(i64(0x4000001)) {
											goto l63
										}
									}
									m.fn51(v5+i32(244), i32(1075648), i32(59))
									store32(m.memory[int64(uint32(v5))+260:], uint32(i32(24)))
									store32(m.memory[int64(uint32(v5))+256:], uint32(i32(1075707)))
									t174 := int32(load32(m.memory[int64(uint32(v5))+244:]))
									v16 = t174
									t175 := int32(load32(m.memory[int64(uint32(v5))+248:]))
									v14 = t175
									t176 := int32(load32(m.memory[int64(uint32(v5))+252:]))
									v12 = t176
									t177 := int64(load64(m.memory[int64(uint32(v5))+256:]))
									v10 = t177
									v8 = i32(-0x7ffffffd)
									goto l64
								}
							}
						l61:
							t210 := int64(load64(m.memory[int64(uint32(v5))+256:]))
							v10 = t210
							t211 := int32(load32(m.memory[int64(uint32(v5))+252:]))
							v12 = t211
							t212 := int32(load32(m.memory[int64(uint32(v5))+248:]))
							v14 = t212
							t213 := int32(load32(m.memory[int64(uint32(v5))+244:]))
							v16 = t213
						}
					l64:
						m.fn1505(v5 + i32(188))
						goto l14
					}
					t61 := v5 + i32(136)
					v15 = v1 + i32(16)
					t62 := int32(load32(m.memory[uint32(v15):]))
					v8 = t62
					t63 := v8
					v11 = v1 + i32(20)
					t64 := int32(load32(m.memory[uint32(v11):]))
					v13 = t64
					m.fn1046(t61, t63, v13, i32(1074726), i32(47), i32(1085598), i32(23))
					{
						{
							t65 := int32(load32(m.memory[int64(uint32(v5))+136:]))
							v6 = t65
							if v6 != 0 {
								goto l19
							}
							v10 = i64(1)
							goto l20
						}
					l19:
						t66 := int32(load32(m.memory[int64(uint32(v5))+140:]))
						m.fn1190(v5+i32(240), v6, t66)
						t67 := int64(load64(m.memory[int64(uint32(v5))+248:]))
						v10 = t67
						p68 := i64(1)
						if uint64(v10) > uint64(i64(1)) {
							p68 = v10
						}
						t69 := int32(m.memory[int64(uint32(v5))+240])
						p70 := p68
						if t69 != 0 {
							p70 = i64(1)
						}
						v10 = p70
					}
				l20:
					{
						t71 := m.fn847(v1, i32(1074726), i32(47), i32(1074773), i32(18))
						if t71 != 0 {
							store32(m.memory[int64(uint32(v5))+240:], uint32(i32(0)))
							store64(m.memory[int64(uint32(v5))+248:], uint64(v10))
							m.fn1504(v5+i32(220), v5+i32(240))
							goto l22
						}
						t72 := m.fn847(v1, i32(1074726), i32(47), i32(1071943), i32(10))
						if t72 == 0 {
							goto l22
						}
						m.fn1046(v5+i32(128), v8, v13, i32(1074726), i32(47), i32(0x106667), i32(22))
						v17 = i32(1)
						v18 = i32(1)
						{
							t73 := int32(load32(m.memory[int64(uint32(v5))+128:]))
							v6 = t73
							if v6 == 0 {
								goto l23
							}
							t74 := int32(load32(m.memory[int64(uint32(v5))+132:]))
							m.fn1071(v5+i32(240), v6, t74)
							t75 := int32(load32(m.memory[int64(uint32(v5))+244:]))
							v6 = t75
							p76 := i32(1)
							if uint32(v6) > uint32(i32(1)) {
								p76 = v6
							}
							t77 := int32(m.memory[int64(uint32(v5))+240])
							p78 := p76
							if t77 != 0 {
								p78 = i32(1)
							}
							v18 = p78
						}
					l23:
						m.fn1046(v5+i32(120), v8, v13, i32(1074726), i32(47), i32(1074813), i32(19))
						{
							t79 := int32(load32(m.memory[int64(uint32(v5))+120:]))
							v6 = t79
							if v6 == 0 {
								goto l24
							}
							t80 := int32(load32(m.memory[int64(uint32(v5))+124:]))
							m.fn1071(v5+i32(240), v6, t80)
							t81 := int32(load32(m.memory[int64(uint32(v5))+244:]))
							v6 = t81
							p82 := i32(1)
							if uint32(v6) > uint32(i32(1)) {
								p82 = v6
							}
							t83 := int32(m.memory[int64(uint32(v5))+240])
							p84 := p82
							if t83 != 0 {
								p84 = i32(1)
							}
							v17 = p84
						}
					l24:
						m.fn1306(v5+i32(240), v1, v2)
						t85 := int32(load32(m.memory[int64(uint32(v5))+252:]))
						v12 = t85
						t86 := int32(load32(m.memory[int64(uint32(v5))+248:]))
						v14 = t86
						t87 := int32(load32(m.memory[int64(uint32(v5))+244:]))
						v16 = t87
						t88 := int32(load32(m.memory[int64(uint32(v5))+240:]))
						v8 = t88
						if v8 != i32(-1) {
							t148 := int64(load64(m.memory[int64(uint32(v5))+256:]))
							v10 = t148
							m.fn1505(v5 + i32(220))
							goto l14
						}
						store32(m.memory[int64(uint32(v5))+284:], uint32(v14))
						store32(m.memory[int64(uint32(v5))+280:], uint32(v16))
						store32(m.memory[int64(uint32(v5))+288:], uint32(v12))
						v6 = v12 << 5
						v1 = v14
					l28:
						{
							if v6 == 0 {
								t93 := int32(load32(m.memory[uint32(v15):]))
								t94 := v5 + i32(112)
								v6 = t93
								t95 := int32(load32(m.memory[uint32(v11):]))
								t96 := v6
								v8 = t95
								m.fn1046(t94, t96, v8, i32(1074169), i32(48), i32(1074832), i32(10))
								t97 := int32(load32(m.memory[int64(uint32(v5))+112:]))
								v1 = t97
								if v1 == 0 {
									goto l29
								}
								t98 := int32(load32(m.memory[int64(uint32(v5))+116:]))
								t99 := v1
								v13 = t98
								t100 := m.fn15(t99, v13, i32(1085492), i32(10))
								if t100 != 0 {
									m.fn1046(v5+i32(48), v6, v8, i32(1074169), i32(48), i32(1077066), i32(5))
									t214 := int32(load32(m.memory[int64(uint32(v5))+48:]))
									v1 = t214
									if v1 == 0 {
										goto l29
									}
									t215 := int32(load32(m.memory[int64(uint32(v5))+52:]))
									m.fn217(v5+i32(240), v1, t215)
									t216 := int32(m.memory[int64(uint32(v5))+240])
									if t216 == i32(1) {
										goto l29
									}
									t217 := math.Float64frombits(load64(m.memory[int64(uint32(v5))+248:]))
									m.fn1508(v5+i32(240), float64(t217*float64(100)))
									store32(m.memory[int64(uint32(v5))+204:], uint32(i32(25)))
									store32(m.memory[int64(uint32(v5))+200:], uint32(v5+i32(240)))
									m.fn73(v5+i32(296), i32(1070086), v5+i32(200))
									t218 := int32(load32(m.memory[int64(uint32(v5))+240:]))
									t219 := int32(load32(m.memory[int64(uint32(v5))+244:]))
									m.fn16(t218, t219)
									goto l36
								}
								t101 := m.fn15(v1, v13, i32(1085502), i32(8))
								if t101 != 0 {
									m.fn1046(v5+i32(64), v6, v8, i32(1074169), i32(48), i32(1077066), i32(5))
									t220 := int32(load32(m.memory[int64(uint32(v5))+64:]))
									v13 = t220
									if v13 == 0 {
										goto l29
									}
									t221 := int32(load32(m.memory[int64(uint32(v5))+68:]))
									v12 = t221
									m.fn1046(v5+i32(56), v6, v8, i32(1074169), i32(48), i32(1085502), i32(8))
									t222 := int32(load32(m.memory[int64(uint32(v5))+60:]))
									t223 := int32(load32(m.memory[int64(uint32(v5))+56:]))
									t224 := v5
									v1 = t223
									p225 := i32(0)
									if v1 != 0 {
										p225 = t222
									}
									v6 = p225
									store32(m.memory[int64(uint32(t224))+348:], uint32(v6))
									t227 := v5
									p226 := i32(1)
									if v1 != 0 {
										p226 = v1
									}
									store32(m.memory[int64(uint32(t227))+344:], uint32(p226))
									{
										{
											{
												if v6 == 0 {
													goto l75
												}
												m.fn217(v5+i32(240), v13, v12)
												t228 := int32(m.memory[int64(uint32(v5))+240])
												if t228 != 0 {
													goto l76
												}
												t229 := math.Float64frombits(load64(m.memory[int64(uint32(v5))+248:]))
												m.fn1508(v5+i32(200), t229)
												store32(m.memory[int64(uint32(v5))+252:], uint32(i32(1)))
												store32(m.memory[int64(uint32(v5))+244:], uint32(i32(25)))
												store32(m.memory[int64(uint32(v5))+248:], uint32(v5+i32(344)))
												store32(m.memory[int64(uint32(v5))+240:], uint32(v5+i32(200)))
												m.fn73(v5+i32(308), i32(1052689), v5+i32(240))
												t230 := int32(load32(m.memory[int64(uint32(v5))+200:]))
												t231 := int32(load32(m.memory[int64(uint32(v5))+204:]))
												m.fn16(t230, t231)
												goto l77
											}
										l75:
											m.fn217(v5+i32(240), v13, v12)
											t232 := int32(m.memory[int64(uint32(v5))+240])
											if t232 != 0 {
												goto l76
											}
											t233 := math.Float64frombits(load64(m.memory[int64(uint32(v5))+248:]))
											m.fn1508(v5+i32(308), t233)
										}
									l77:
										t234 := int32(load32(m.memory[int64(uint32(v5))+316:]))
										store32(m.memory[int64(uint32(v5))+304:], uint32(t234))
										t235 := int64(load64(m.memory[int64(uint32(v5))+308:]))
										store64(m.memory[int64(uint32(v5))+296:], uint64(t235))
										goto l36
									}
								l76:
									store32(m.memory[int64(uint32(v5))+296:], uint32(i32(-1)))
									goto l36
								}
								t102 := m.fn15(v1, v13, i32(1085510), i32(5))
								if t102 != 0 {
									m.fn1046(v5+i32(72), v6, v8, i32(1074169), i32(48), i32(1077066), i32(5))
									t236 := int32(load32(m.memory[int64(uint32(v5))+72:]))
									v1 = t236
									if v1 == 0 {
										goto l29
									}
									t237 := int32(load32(m.memory[int64(uint32(v5))+76:]))
									m.fn217(v5+i32(240), v1, t237)
									t238 := int32(m.memory[int64(uint32(v5))+240])
									if t238 == i32(1) {
										goto l29
									}
									t239 := math.Float64frombits(load64(m.memory[int64(uint32(v5))+248:]))
									m.fn1508(v5+i32(296), t239)
									goto l36
								}
								{
									t103 := m.fn15(v1, v13, i32(1085515), i32(4))
									if t103 != 0 {
										m.fn1046(v5+i32(80), v6, v8, i32(1074169), i32(48), i32(1085565), i32(10))
										t146 := int32(load32(m.memory[int64(uint32(v5))+80:]))
										t147 := int32(load32(m.memory[int64(uint32(v5))+84:]))
										m.fn1041(v5+i32(296), t146, t147)
										goto l36
									}
									{
										t104 := m.fn15(v1, v13, i32(1085519), i32(4))
										if t104 != 0 {
											m.fn1046(v5+i32(88), v6, v8, i32(1074169), i32(48), i32(1085555), i32(10))
											t109 := int32(load32(m.memory[int64(uint32(v5))+88:]))
											v12 = t109
											if v12 == 0 {
												goto l29
											}
											t110 := int32(load32(m.memory[int64(uint32(v5))+92:]))
											v14 = t110
											v6 = i32(0)
											store32(m.memory[int64(uint32(v5))+208:], uint32(i32(0)))
											store64(m.memory[int64(uint32(v5))+200:], uint64(i64(0x100000000)))
											store32(m.memory[int64(uint32(v5))+344:], uint32(v12))
											store32(m.memory[int64(uint32(v5))+348:], uint32(v12+v14))
											v19 = float64(0)
											v13 = i32(0)
										l39:
											v8 = v6
											{
												{
													{
														{
															{
																t111 := m.fn48(v5 + i32(344))
																v1 = t111
																switch v1 + i32(-77) {
																case 0:
																	goto l37
																case 1, 2, 4, 5, 8, 9, 11:
																	goto l38
																case 3:
																	goto l39
																case 6:
																	if v13&i32(1) == 0 {
																		goto l38
																	}
																	goto l56
																case 10:
																	goto l42
																case 12:
																	if v13&i32(1) != 0 {
																		goto l38
																	}
																	goto l57
																default:
																	switch v1 + i32(-68) {
																	case 0:
																		if v13&i32(1) != 0 {
																			goto l38
																		}
																		goto l55
																	case 1, 2, 3:
																		goto l38
																	case 4:
																		if v13&i32(1) == 0 {
																			goto l38
																		}
																		goto l56
																	default:
																		if v1 == i32(45) {
																			if v19 == float64(0) {
																				v6 = i32(1)
																				t122 := int32(load32(m.memory[int64(uint32(v5))+208:]))
																				if t122 == 0 {
																					goto l39
																				}
																				goto l38
																			}
																			goto l38
																		}
																		if v1 != i32(-1) {
																			goto l38
																		}
																		{
																			t112 := fn1854(float64(v19 * float64(1000)))
																			v19 = t112
																			if !(v19 >= float64(0)) {
																				goto l49
																			}
																			if v19 < float64(1e+18) {
																				t113 := v5
																				v1 = v8 & i32(1)
																				store32(m.memory[int64(uint32(t113))+324:], uint32(v1))
																				t115 := v5
																				p114 := i32(1)
																				if v1 != 0 {
																					p114 = i32(1108000)
																				}
																				store32(m.memory[int64(uint32(t115))+320:], uint32(p114))
																				t116 := v5
																				v20 = i64_trunc_sat_f64_u(v19)
																				t117 := int64(uint64(v20) / uint64(i64(3600000)))
																				v21 = t117
																				store64(m.memory[int64(uint32(t116))+328:], uint64(v21))
																				t118 := v5
																				v1 = int32(v20 - v21*i64(3600000))
																				t119 := int32(uint32(v1) / uint32(i32(60000)))
																				v6 = t119
																				store64(m.memory[int64(uint32(t118))+336:], uint64(uint32(v6)))
																				v1 = v1 - v6*i32(60000)
																				t120 := int32(uint32(v1&i32(0xffff)) / uint32(i32(1000)))
																				t121 := v1
																				v6 = t120
																				if (t121-v6*i32(1000))&i32(0xffff) != 0 {
																					store64(m.memory[int64(uint32(v5))+344:], math.Float64bits(float64(float64(uint32(v1))/float64(1000))))
																					store32(m.memory[int64(uint32(v5))+268:], uint32(i32(66)))
																					store32(m.memory[int64(uint32(v5))+260:], uint32(i32(28)))
																					store32(m.memory[int64(uint32(v5))+252:], uint32(i32(28)))
																					store32(m.memory[int64(uint32(v5))+244:], uint32(i32(1)))
																					store32(m.memory[int64(uint32(v5))+264:], uint32(v5+i32(344)))
																					store32(m.memory[int64(uint32(v5))+256:], uint32(v5+i32(336)))
																					store32(m.memory[int64(uint32(v5))+248:], uint32(v5+i32(328)))
																					store32(m.memory[int64(uint32(v5))+240:], uint32(v5+i32(320)))
																					m.fn73(v5+i32(308), i32(1085575), v5+i32(240))
																					goto l51
																				}
																				store64(m.memory[int64(uint32(v5))+344:], uint64(uint32(v6)))
																				store32(m.memory[int64(uint32(v5))+268:], uint32(i32(28)))
																				store32(m.memory[int64(uint32(v5))+260:], uint32(i32(28)))
																				store32(m.memory[int64(uint32(v5))+252:], uint32(i32(28)))
																				store32(m.memory[int64(uint32(v5))+244:], uint32(i32(1)))
																				store32(m.memory[int64(uint32(v5))+264:], uint32(v5+i32(344)))
																				store32(m.memory[int64(uint32(v5))+256:], uint32(v5+i32(336)))
																				store32(m.memory[int64(uint32(v5))+248:], uint32(v5+i32(328)))
																				store32(m.memory[int64(uint32(v5))+240:], uint32(v5+i32(320)))
																				m.fn73(v5+i32(308), i32(1083673), v5+i32(240))
																				goto l51
																			}
																		}
																	l49:
																		m.fn51(v5+i32(308), v12, v14)
																		goto l51
																	}
																case 7:
																	store32(m.memory[int64(uint32(v5))+208:], uint32(i32(0)))
																	goto l54
																}
															}
														l37:
															if v13&i32(1) == 0 {
																goto l57
															}
														l56:
															t123 := int32(load32(m.memory[int64(uint32(v5))+204:]))
															t124 := int32(load32(m.memory[int64(uint32(v5))+208:]))
															m.fn217(v5+i32(240), t123, t124)
															store32(m.memory[int64(uint32(v5))+208:], uint32(i32(0)))
															t126 := v19
															p125 := float64(1)
															if v1 == i32(77) {
																p125 = float64(60)
															}
															p127 := p125
															if v1 == i32(72) {
																p127 = float64(3600)
															}
															t128 := math.Float64frombits(load64(m.memory[int64(uint32(v5))+248:]))
															t129 := int32(m.memory[int64(uint32(v5))+240])
															p130 := t128
															if t129 != 0 {
																p130 = float64(0)
															}
															v19 = float64(t126 + float64(p127*p130))
															goto l54
														}
													l42:
														if v13&i32(1) != 0 {
															goto l38
														}
													l55:
														t131 := int32(load32(m.memory[int64(uint32(v5))+204:]))
														t132 := int32(load32(m.memory[int64(uint32(v5))+208:]))
														m.fn217(v5+i32(240), t131, t132)
														v13 = i32(0)
														store32(m.memory[int64(uint32(v5))+208:], uint32(i32(0)))
														t134 := v19
														p133 := float64(86400)
														if v1 == i32(87) {
															p133 = float64(604800)
														}
														t135 := math.Float64frombits(load64(m.memory[int64(uint32(v5))+248:]))
														t136 := int32(m.memory[int64(uint32(v5))+240])
														p137 := t135
														if t136 != 0 {
															p137 = float64(0)
														}
														v19 = float64(t134 + float64(p133*p137))
														v6 = v8
														goto l39
													}
												l57:
													t138 := int32(load32(m.memory[int64(uint32(v5))+204:]))
													t139 := int32(load32(m.memory[int64(uint32(v5))+208:]))
													m.fn217(v5+i32(240), t138, t139)
													{
														t140 := int32(m.memory[int64(uint32(v5))+240])
														if t140 != 0 {
															goto l58
														}
														t141 := math.Float64frombits(load64(m.memory[int64(uint32(v5))+248:]))
														if t141 != float64(0) {
															goto l59
														}
													}
												l58:
													v13 = i32(0)
													store32(m.memory[int64(uint32(v5))+208:], uint32(i32(0)))
													v6 = v8
													goto l39
												l59:
													m.fn51(v5+i32(308), v12, v14)
												}
											l51:
												t142 := int32(load32(m.memory[int64(uint32(v5))+200:]))
												t143 := int32(load32(m.memory[int64(uint32(v5))+204:]))
												m.fn16(t142, t143)
												t144 := int64(load64(m.memory[int64(uint32(v5))+308:]))
												store64(m.memory[int64(uint32(v5))+296:], uint64(t144))
												t145 := int32(load32(m.memory[int64(uint32(v5))+316:]))
												store32(m.memory[int64(uint32(v5))+304:], uint32(t145))
												goto l36
											}
										l54:
											v13 = i32(1)
											v6 = v8
											goto l39
										l38:
											if v1 == i32(46) {
												goto l60
											}
											if uint32(v1+i32(-48)) < uint32(i32(10)) {
												goto l60
											}
											store32(m.memory[int64(uint32(v5))+208:], uint32(i32(0)))
											v6 = v8
											goto l39
										l60:
											m.fn74(v5+i32(200), v1)
											v6 = v8
											goto l39
										}
										t105 := m.fn15(v1, v13, i32(1085523), i32(7))
										if t105 != 0 {
											goto l35
										}
										t106 := m.fn15(v1, v13, i32(1088388), i32(6))
										if t106 == 0 {
											goto l29
										}
										m.fn1046(v5+i32(104), v6, v8, i32(1074169), i32(48), i32(1085530), i32(12))
										t107 := int32(load32(m.memory[int64(uint32(v5))+104:]))
										t108 := int32(load32(m.memory[int64(uint32(v5))+108:]))
										m.fn1041(v5+i32(296), t107, t108)
										goto l36
									}
								}
							}
							t89 := int32(load32(m.memory[uint32(v1):]))
							if t89 != i32(-0x80000000) {
								goto l27
							}
							v6 = v6 + i32(-32)
							t90 := int32(load32(m.memory[int64(uint32(v1))+12:]))
							v8 = t90
							t91 := int32(load32(m.memory[int64(uint32(v1))+8:]))
							v13 = t91
							v1 = v1 + i32(32)
							t92 := m.fn23(v13, v8)
							if t92 != 0 {
								goto l28
							}
							goto l27
						}
					}
				}
			l35:
				m.fn1046(v5+i32(96), v6, v8, i32(1074169), i32(48), i32(1085542), i32(13))
				t240 := int32(load32(m.memory[int64(uint32(v5))+96:]))
				v1 = t240
				if v1 == 0 {
					goto l29
				}
				t241 := int32(load32(m.memory[int64(uint32(v5))+100:]))
				t242 := m.fn15(v1, t241, i32(1071691), i32(4))
				t243 := v5 + i32(296)
				v1 = t242
				p244 := i32(1089116)
				if v1 != 0 {
					p244 = i32(1089121)
				}
				p245 := i32(5)
				if v1 != 0 {
					p245 = i32(4)
				}
				m.fn51(t243, p244, p245)
			}
		l36:
			t246 := int32(load32(m.memory[int64(uint32(v5))+296:]))
			if t246 == i32(-1) {
				goto l78
			}
			t247 := m.fn113(i32(8), i32(32))
			v14 = t247
			t248 := m.fn113(i32(4), i32(28))
			v1 = t248
			store32(m.memory[uint32(v1):], uint32(i32(3)))
			store32(m.memory[int64(uint32(v1))+16:], uint32(i32(0)))
			t249 := int64(load64(m.memory[int64(uint32(v5))+296:]))
			store64(m.memory[int64(uint32(v1))+4:], uint64(t249))
			t250 := int32(load32(m.memory[int64(uint32(v5))+304:]))
			store32(m.memory[int64(uint32(v1))+12:], uint32(t250))
			store64(m.memory[uint32(v14):], uint64(i64(0x180000000)))
			v12 = i32(1)
			store32(m.memory[int64(uint32(v14))+12:], uint32(i32(1)))
			store32(m.memory[int64(uint32(v14))+8:], uint32(v1))
			goto l79
		}
	l29:
		store32(m.memory[int64(uint32(v5))+296:], uint32(i32(-1)))
	l78:
		v14 = i32(8)
		v12 = i32(0)
	l79:
		m.fn969(v5 + i32(280))
		v16 = v12
	l27:
		t251 := m.fn1019(v14, v12)
		store64(m.memory[int64(uint32(v5))+272:], uint64(t251))
		store32(m.memory[int64(uint32(v5))+260:], uint32(v12))
		store32(m.memory[int64(uint32(v5))+256:], uint32(v14))
		store32(m.memory[int64(uint32(v5))+252:], uint32(v16))
		store32(m.memory[int64(uint32(v5))+248:], uint32(v17))
		store32(m.memory[int64(uint32(v5))+244:], uint32(v18))
		store64(m.memory[int64(uint32(v5))+264:], uint64(v10))
		store32(m.memory[int64(uint32(v5))+240:], uint32(i32(1)))
		m.fn1504(v5+i32(220), v5+i32(240))
		goto l22
	}
l14:
	store64(m.memory[int64(uint32(v0))+16:], uint64(v10))
	store32(m.memory[int64(uint32(v0))+12:], uint32(v12))
	store32(m.memory[int64(uint32(v0))+8:], uint32(v14))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v16))
	store32(m.memory[uint32(v0):], uint32(v8))
l5:
	m.g0 = v5 + i32(352)
}
func (m *Module) fn1336(v0, v1, v2 int32) {
	var v3, v4 int32
	{
		{
			t0 := int32(load32(m.memory[uint32(v0):]))
			t1 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			t2 := v2
			v3 = t1
			if uint32(t2) <= uint32(t0-v3) {
				goto l0
			}
			m.fn62(v0, v3, v2, i32(8), i32(32))
			t3 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			v3 = t3
			goto l1
		}
	l0:
		if v2 == 0 {
			goto l2
		}
	l1:
		v4 = v2 << 5
		if v4 == 0 {
			goto l2
		}
		t4 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		memory_copy(m.memory, uint32(t4+v3<<5), uint32(v1), uint32(v4))
	}
l2:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v3+v2))
}
func (m *Module) fn1337(v0, v1, v2, v3 int32) int32 {
	t0 := m.fn15(v0, v1, v2, v3)
	return t0 ^ i32(1)
}
func (m *Module) fn1338(v0, v1, v2, v3, v4 int32) {
	var v5, v6, v7 int32
	var v8 int64
	var v9, v10, v11, v12, v13, v14, v15 int32
	var v16 int64
	t0 := m.g0
	v5 = t0 - i32(176)
	m.g0 = v5
	{
		{
			{
				{
					t1 := int32(load32(m.memory[int64(uint32(v1))+28:]))
					v6 = t1
					t2 := int32(load32(m.memory[int64(uint32(v1))+32:]))
					t3 := v6
					v1 = t2
					t4 := m.fn886(t3, v1, i32(1074120), i32(49), i32(1081813), i32(8))
					v7 = t4
					if v7 == 0 {
						{
							{
								t11 := m.fn1097(v6, v1, i32(1074620), i32(56), i32(1073751), i32(5))
								v7 = t11
								if v7 == 0 {
									goto l3
								}
								t12 := int32(load32(m.memory[uint32(v7+i32(28)):]))
								t13 := int32(load32(m.memory[uint32(v7+i32(32)):]))
								m.fn864(v5+i32(116), t12, t13)
								t14 := int32(load32(m.memory[int64(uint32(v5))+116:]))
								if t14 == i32(-1) {
									goto l3
								}
								t15 := int32(load32(m.memory[int64(uint32(v5))+124:]))
								store32(m.memory[int64(uint32(v5))+96:], uint32(t15))
								t16 := int64(load64(m.memory[int64(uint32(v5))+116:]))
								store64(m.memory[int64(uint32(v5))+88:], uint64(t16))
								goto l4
							}
						l3:
							{
								t17 := m.fn1097(v6, v1, i32(1074620), i32(56), i32(1074676), i32(4))
								v7 = t17
								if v7 == 0 {
									goto l5
								}
								t18 := int32(load32(m.memory[uint32(v7+i32(28)):]))
								t19 := int32(load32(m.memory[uint32(v7+i32(32)):]))
								m.fn864(v5+i32(88), t18, t19)
								goto l6
							}
						l5:
							store32(m.memory[int64(uint32(v5))+88:], uint32(i32(-1)))
						l6:
							t20 := int32(load32(m.memory[int64(uint32(v5))+120:]))
							m.fn134(i32(-1), t20)
						}
					l4:
						t21 := int32(load32(m.memory[int64(uint32(v5))+92:]))
						t22 := int32(load32(m.memory[int64(uint32(v5))+88:]))
						t23 := v5 + i32(24)
						v4 = t22
						var p24 int32
						if v4 == i32(-1) {
							p24 = 1
						}
						v7 = p24
						p25 := t21
						if v7 != 0 {
							p25 = i32(1)
						}
						v9 = p25
						t26 := int32(load32(m.memory[int64(uint32(v5))+96:]))
						t28 := v9
						p27 := t26
						if v7 != 0 {
							p27 = i32(0)
						}
						m.fn46(t23, t28, p27)
						t29 := int32(load32(m.memory[int64(uint32(v5))+24:]))
						t30 := int32(load32(m.memory[int64(uint32(v5))+28:]))
						m.fn865(v5+i32(44), t29, t30)
						p31 := v4
						if v7 != 0 {
							p31 = i32(0)
						}
						v7 = p31
						t32 := m.fn1097(v6, v1, i32(1074120), i32(49), i32(1077128), i32(5))
						v1 = t32
						if v1 == 0 {
							{
								{
									t66 := int32(load32(m.memory[int64(uint32(v5))+52:]))
									if t66 == 0 {
										goto l17
									}
									t67 := int32(load32(m.memory[int64(uint32(v5))+52:]))
									store32(m.memory[int64(uint32(v5))+128:], uint32(t67))
									t68 := int64(load64(m.memory[int64(uint32(v5))+44:]))
									store64(m.memory[int64(uint32(v5))+120:], uint64(t68))
									store32(m.memory[int64(uint32(v5))+132:], uint32(i32(-0x7fffffff)))
									store32(m.memory[int64(uint32(v5))+116:], uint32(i32(5)))
									m.fn1340(v3, v5+i32(116))
									goto l18
								}
							l17:
								t69 := int32(load32(m.memory[int64(uint32(v5))+44:]))
								t70 := int32(load32(m.memory[int64(uint32(v5))+48:]))
								m.fn16(t69, t70)
							}
						l18:
							store32(m.memory[uint32(v0):], uint32(i32(-1)))
							goto l19
						}
						t33 := int32(load32(m.memory[uint32(v1+i32(16)):]))
						t34 := int32(load32(m.memory[uint32(v1+i32(20)):]))
						m.fn1046(v5+i32(16), t33, t34, i32(1085191), i32(28), i32(1073490), i32(4))
						t35 := int32(load32(m.memory[int64(uint32(v5))+16:]))
						v1 = t35
						if v1 == 0 {
							goto l8
						}
						t36 := int32(load32(m.memory[int64(uint32(v5))+20:]))
						v6 = t36
						if v6 == 0 {
							goto l8
						}
						p37 := i32(1)
						if v1 != 0 {
							p37 = v1
						}
						v1 = p37
						t38 := m.fn1455(v1, v6)
						if t38 != 0 {
							m.fn51(v5+i32(92), v1, v6)
							goto l11
						}
						m.fn774(v5+i32(116), i32(1071695), i32(11), v1, v6)
						t39 := int32(load32(m.memory[int64(uint32(v5))+116:]))
						if t39 == 0 {
							t40 := int32(load32(m.memory[int64(uint32(v5))+136:]))
							v10 = t40
							t41 := int32(load32(m.memory[int64(uint32(v5))+132:]))
							v11 = t41
							t42 := int32(load32(m.memory[int64(uint32(v5))+120:]))
							v12 = t42
							t43 := int32(load32(m.memory[int64(uint32(v5))+128:]))
							v4 = t43
							t44 := int32(load32(m.memory[int64(uint32(v5))+124:]))
							v1 = t44
							t45 := int32(load32(m.memory[int64(uint32(v2))+204:]))
							m.fn1182(v5+i32(8), t45, i32(1085020))
							t46 := int32(load32(m.memory[int64(uint32(v5))+12:]))
							v6 = t46
							t47 := int32(load32(m.memory[int64(uint32(v5))+8:]))
							m.fn1035(v5+i32(116), t47, v1, v4)
							t48 := int32(load32(m.memory[int64(uint32(v5))+124:]))
							v13 = t48
							t49 := int32(load32(m.memory[int64(uint32(v5))+120:]))
							v14 = t49
							{
								t50 := int32(load32(m.memory[int64(uint32(v5))+116:]))
								v15 = t50
								if v15 == i32(-1) {
									if v14 == 0 {
										t63 := int32(load32(m.memory[uint32(v6):]))
										store32(m.memory[uint32(v6):], uint32(t63+i32(1)))
										store32(m.memory[int64(uint32(v5))+92:], uint32(i32(-1)))
										m.fn16(v12, v1)
										goto l16
									}
									store32(m.memory[int64(uint32(v5))+148:], uint32(v13))
									store32(m.memory[int64(uint32(v5))+144:], uint32(v14))
									m.fn1476(v5+i32(152), v1, v4)
									t54 := int32(load32(m.memory[int64(uint32(v2))+208:]))
									m.fn1182(v5, t54, i32(1085036))
									t55 := int32(load32(m.memory[int64(uint32(v5))+4:]))
									v2 = t55
									t56 := int32(load32(m.memory[uint32(v5):]))
									v15 = t56
									store32(m.memory[int64(uint32(v5))+172:], uint32(v4))
									store32(m.memory[int64(uint32(v5))+168:], uint32(v1))
									store32(m.memory[int64(uint32(v5))+164:], uint32(v12))
									m.fn1296(v5+i32(116), v15, v5+i32(152), v5+i32(164), v14+i32(8), v13)
									t57 := int32(load32(m.memory[int64(uint32(v5))+120:]))
									v1 = t57
									t58 := int32(load32(m.memory[int64(uint32(v5))+116:]))
									v15 = t58
									if v15 == i32(-1) {
										t64 := int32(load32(m.memory[uint32(v2):]))
										store32(m.memory[uint32(v2):], uint32(t64+i32(1)))
										store32(m.memory[int64(uint32(v5))+96:], uint32(v1))
										store32(m.memory[int64(uint32(v5))+92:], uint32(i32(-0x80000000)))
										m.fn754(v5 + i32(144))
										t65 := int32(load32(m.memory[uint32(v6):]))
										store32(m.memory[uint32(v6):], uint32(t65+i32(1)))
										goto l16
									}
									t59 := int64(load64(m.memory[int64(uint32(v5))+132:]))
									store64(m.memory[int64(uint32(v5))+104:], uint64(t59))
									t60 := int64(load64(m.memory[int64(uint32(v5))+124:]))
									store64(m.memory[int64(uint32(v5))+96:], uint64(t60))
									t61 := int32(load32(m.memory[uint32(v2):]))
									store32(m.memory[uint32(v2):], uint32(t61+i32(1)))
									store32(m.memory[int64(uint32(v5))+92:], uint32(v1))
									m.fn754(v5 + i32(144))
									t62 := int32(load32(m.memory[uint32(v6):]))
									store32(m.memory[uint32(v6):], uint32(t62+i32(1)))
									goto l13
								}
								t51 := int32(load32(m.memory[int64(uint32(v5))+136:]))
								store32(m.memory[int64(uint32(v5))+108:], uint32(t51))
								t52 := int64(load64(m.memory[int64(uint32(v5))+128:]))
								store64(m.memory[int64(uint32(v5))+100:], uint64(t52))
								t53 := int32(load32(m.memory[uint32(v6):]))
								store32(m.memory[uint32(v6):], uint32(t53+i32(1)))
								store32(m.memory[int64(uint32(v5))+96:], uint32(v13))
								store32(m.memory[int64(uint32(v5))+92:], uint32(v14))
								m.fn16(v12, v1)
								goto l13
							}
						}
						store32(m.memory[int64(uint32(v5))+92:], uint32(i32(-1)))
						m.fn785(v5 + i32(120))
						goto l11
					}
					m.fn1306(v5+i32(116), v7, v2)
					t5 := int64(load64(m.memory[int64(uint32(v5))+120:]))
					store64(m.memory[int64(uint32(v5))+32:], uint64(t5))
					t6 := int32(load32(m.memory[int64(uint32(v5))+128:]))
					store32(m.memory[int64(uint32(v5))+40:], uint32(t6))
					t7 := int32(load32(m.memory[int64(uint32(v5))+116:]))
					v1 = t7
					if v1 == i32(-1) {
						m.fn1271(v4, v5+i32(32))
						store32(m.memory[uint32(v0):], uint32(i32(-1)))
						goto l2
					}
					t8 := int64(load64(m.memory[int64(uint32(v5))+132:]))
					v8 = t8
					t9 := int32(load32(m.memory[int64(uint32(v5))+40:]))
					store32(m.memory[int64(uint32(v0))+12:], uint32(t9))
					t10 := int64(load64(m.memory[int64(uint32(v5))+32:]))
					store64(m.memory[int64(uint32(v0))+4:], uint64(t10))
					store64(m.memory[int64(uint32(v0))+16:], uint64(v8))
					store32(m.memory[uint32(v0):], uint32(v1))
					goto l2
				}
			l8:
				store32(m.memory[int64(uint32(v5))+92:], uint32(i32(-1)))
				goto l11
			l16:
				m.fn134(v11, v10)
			l11:
				t71 := int64(load64(m.memory[int64(uint32(v5))+92:]))
				t72 := v5
				v8 = t71
				store64(m.memory[int64(uint32(t72))+56:], uint64(v8))
				t73 := int32(load32(m.memory[int64(uint32(v5))+100:]))
				store32(m.memory[int64(uint32(v5))+64:], uint32(t73))
				{
					if int32(v8) != i32(-1) {
						goto l20
					}
					t74 := int32(load32(m.memory[int64(uint32(v5))+52:]))
					if t74 != 0 {
						goto l20
					}
					store32(m.memory[uint32(v0):], uint32(i32(-1)))
					t75 := int32(load32(m.memory[int64(uint32(v5))+60:]))
					m.fn1471(i32(-1), t75)
					goto l21
				}
			l20:
				t76 := int32(load32(m.memory[int64(uint32(v5))+52:]))
				store32(m.memory[int64(uint32(v5))+128:], uint32(t76))
				t77 := int64(load64(m.memory[int64(uint32(v5))+44:]))
				store64(m.memory[int64(uint32(v5))+120:], uint64(t77))
				store32(m.memory[int64(uint32(v5))+88:], uint32(i32(-0x7fffffff)))
				m.fn1360(v5+i32(132), v5+i32(56), v5+i32(88))
				store32(m.memory[int64(uint32(v5))+116:], uint32(i32(5)))
				m.fn1340(v3, v5+i32(116))
				store32(m.memory[uint32(v0):], uint32(i32(-1)))
				goto l19
			}
		l13:
			m.fn134(v11, v10)
			t78 := int32(load32(m.memory[int64(uint32(v5))+100:]))
			t79 := v5
			v1 = t78
			store32(m.memory[int64(uint32(t79))+80:], uint32(v1))
			t80 := int64(load64(m.memory[int64(uint32(v5))+92:]))
			t81 := v5
			v8 = t80
			store64(m.memory[int64(uint32(t81))+72:], uint64(v8))
			t82 := int64(load64(m.memory[int64(uint32(v5))+104:]))
			v16 = t82
			store32(m.memory[int64(uint32(v0))+12:], uint32(v1))
			store64(m.memory[int64(uint32(v0))+4:], uint64(v8))
			store64(m.memory[int64(uint32(v0))+16:], uint64(v16))
			store32(m.memory[uint32(v0):], uint32(v15))
		}
	l21:
		t83 := int32(load32(m.memory[int64(uint32(v5))+44:]))
		t84 := int32(load32(m.memory[int64(uint32(v5))+48:]))
		m.fn16(t83, t84)
	}
l19:
	m.fn16(v7, v9)
l2:
	m.g0 = v5 + i32(176)
}
func (m *Module) fn1339(v0 int32) {
	var v1, v2 int32
	t0 := int32(load32(m.memory[int64(uint32(v0))+12:]))
	t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	v1 = t1
	v2 = int32(uint32(t0-v1) >> 5)
l1:
	if v2 == 0 {
		goto l0
	}
	v2 = v2 + i32(-1)
	m.fn970(v1)
	v1 = v1 + i32(32)
	goto l1
l0:
	t2 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	t3 := int32(load32(m.memory[uint32(v0):]))
	m.fn80(t2, t3)
}
func (m *Module) fn1340(v0, v1 int32) {
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
func (m *Module) fn1341(v0, v1 int32) {
	var v2, v3 int32
	t0 := int32(load32(m.memory[uint32(v1):]))
	v2 = t0
	t1 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	t2 := v0
	v3 = t1
	t3 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	m.fn1470(t2, v3, t3)
	m.fn82(v2, v3)
}
func (m *Module) fn1342(v0, v1, v2, v3, v4 int32) {
	var v5 int32
	var v6 int64
	t0 := m.g0
	v5 = t0 - i32(16)
	m.g0 = v5
	v6 = int64(uint32(v4)) * int64(uint32(v1))
	if int32(int64(uint64(v6)>>32)) != 0 {
		goto l0
	}
	v4 = int32(v6)
	if uint32(v4) > uint32(i32(-0x80000000)-v3) {
		goto l0
	}
	if v4 != 0 {
		goto l1
	}
	store32(m.memory[int64(uint32(v0))+8:], uint32(v3))
	v3 = i32(0)
	store32(m.memory[int64(uint32(v0))+4:], uint32(i32(0)))
	goto l2
l0:
	store32(m.memory[int64(uint32(v0))+4:], uint32(i32(0)))
	goto l3
l1:
	m.fn1819(v5+i32(8), v3, v4, v2)
	{
		t1 := int32(load32(m.memory[int64(uint32(v5))+8:]))
		v2 = t1
		if v2 == 0 {
			goto l4
		}
		store32(m.memory[int64(uint32(v0))+8:], uint32(v2))
		store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
		v3 = i32(0)
		goto l2
	}
l4:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v4))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
l3:
	v3 = i32(1)
l2:
	store32(m.memory[uint32(v0):], uint32(v3))
	m.g0 = v5 + i32(16)
}
func (m *Module) fn1343(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6 int32
	var v7 int64
	var v8, v9, v10, v11, v12, v13, v14, v15, v16 int32
	var v17 int64
	t0 := m.g0
	v2 = t0 - i32(64)
	m.g0 = v2
	t1 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	v3 = t1
	t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	v4 = t2
	v0 = i32(1)
	{
		t3 := int32(load32(m.memory[uint32(v1):]))
		v5 = t3
		t4 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t5 := v5
		v6 = t4
		t6 := int32(load32(m.memory[int64(uint32(v6))+16:]))
		v1 = t6
		t7 := m.t0[uint(v1)].(func(int32, int32) int32)(t5, i32(34))
		if t7 != 0 {
			goto l0
		}
		store32(m.memory[int64(uint32(v2))+12:], uint32(v3))
		store32(m.memory[int64(uint32(v2))+8:], uint32(v4))
		v7 = int64(uint32(i32(188)))<<32 | int64(uint32(v2+i32(32)))
	l27:
		{
			m.fn93(v2+i32(16), v2+i32(8))
			{
				{
					t8 := int32(load32(m.memory[int64(uint32(v2))+16:]))
					v8 = t8
					if v8 == 0 {
						t32 := m.t0[uint(v1)].(func(int32, int32) int32)(v5, i32(34))
						v0 = t32
						goto l0
					}
					t9 := int32(load32(m.memory[int64(uint32(v2))+28:]))
					v9 = t9
					t10 := int32(load32(m.memory[int64(uint32(v2))+24:]))
					v4 = t10
					{
						t11 := int32(load32(m.memory[int64(uint32(v2))+20:]))
						v10 = t11
						if v10 != 0 {
							v11 = v8 + v10
							v0 = i32(0)
							v12 = v8
							v13 = i32(0)
						l23:
							{
								{
									t12 := int32(int8(m.memory[uint32(v12)]))
									v3 = t12
									if v3 <= i32(-1) {
										goto l4
									}
									v14 = v12 + i32(1)
									v15 = v3 & i32(255)
									goto l5
								}
							l4:
								t13 := int32(m.memory[int64(uint32(v12))+1])
								v16 = t13 & i32(63)
								v14 = v3 & i32(31)
								if uint32(v3) > uint32(i32(-33)) {
									goto l6
								}
								v15 = v14<<6 | v16
								v14 = v12 + i32(2)
								goto l5
							l6:
								t14 := int32(m.memory[int64(uint32(v12))+2])
								v16 = v16<<6 | t14&i32(63)
								if uint32(v3) >= uint32(i32(-16)) {
									goto l7
								}
								v15 = v16 | v14<<12
								v14 = v12 + i32(3)
								goto l5
							l7:
								t15 := int32(m.memory[int64(uint32(v12))+3])
								v15 = v16<<6 | t15&i32(63) | v14<<18&i32(0x1c0000)
								v14 = v12 + i32(4)
							}
						l5:
							m.fn1645(v2+i32(32), v15, i32(65537))
							{
								t16 := int32(m.memory[int64(uint32(v2))+45])
								t17 := int32(m.memory[int64(uint32(v2))+44])
								if (t16-t17)&i32(255) == i32(1) {
									goto l8
								}
								{
									if uint32(v13) < uint32(v0) {
										goto l9
									}
									{
										if v0 == 0 {
											goto l10
										}
										if uint32(v0) < uint32(v10) {
											goto l11
										}
										if v0 != v10 {
											goto l9
										}
										goto l10
									l11:
										t18 := int32(int8(m.memory[uint32(v8+v0)]))
										if t18 <= i32(-65) {
											goto l9
										}
									}
								l10:
									{
										if v13 == 0 {
											goto l12
										}
										if uint32(v13) < uint32(v10) {
											goto l13
										}
										if v13 == v10 {
											goto l12
										}
										goto l9
									l13:
										t19 := int32(int8(m.memory[uint32(v8+v13)]))
										if t19 <= i32(-65) {
											goto l9
										}
									}
								l12:
									t20 := int32(load32(m.memory[int64(uint32(v6))+12:]))
									t21 := m.t0[uint(t20)].(func(int32, int32, int32) int32)(v5, v8+v0, v13-v0)
									if t21 != 0 {
										goto l14
									}
									t22 := int64(load64(m.memory[int64(uint32(v2))+40:]))
									store64(m.memory[int64(uint32(v2))+56:], uint64(t22))
									t23 := int64(load64(m.memory[int64(uint32(v2))+32:]))
									t24 := v2
									v17 = t23
									store64(m.memory[int64(uint32(t24))+48:], uint64(v17))
									t25 := int32(m.memory[int64(uint32(v2))+60])
									v0 = t25
									t26 := int32(m.memory[int64(uint32(v2))+61])
									v3 = t26
									if uint32(v3) < uint32(i32(129)) {
										p28 := v3
										if uint32(v0) > uint32(v3) {
											p28 = v0
										}
										v16 = p28
									l18:
										{
											if v16 == v0 {
												goto l16
											}
											v3 = v2 + i32(48) + v0
											v0 = v0 + i32(1)
											t29 := int32(m.memory[uint32(v3)])
											t30 := m.t0[uint(v1)].(func(int32, int32) int32)(v5, t29)
											if t30 == 0 {
												goto l18
											}
											goto l14
										}
									}
									v16 = int32(v17)
								l17:
									{
										if uint32(v0&i32(255)) >= uint32(v3) {
											goto l16
										}
										v0 = v0 + i32(1)
										t27 := m.t0[uint(v1)].(func(int32, int32) int32)(v5, v16)
										if t27 == 0 {
											goto l17
										}
										goto l14
									}
								}
							l9:
								m.fn556(v8, v10, v0, v13, i32(1131140))
								panic("unreachable")
							l16:
								{
									if uint32(v15) >= uint32(i32(128)) {
										goto l19
									}
									v0 = i32(1)
									goto l20
								l19:
									if uint32(v15) >= uint32(i32(2048)) {
										goto l21
									}
									v0 = i32(2)
									goto l20
								l21:
									p31 := i32(4)
									if uint32(v15) < uint32(i32(65536)) {
										p31 = i32(3)
									}
									v0 = p31
								}
							l20:
								v0 = v0 + v13
							}
						l8:
							v13 = v13 - v12 + v14
							v12 = v14
							if v14 == v11 {
								goto l22
							}
							goto l23
						}
						v3 = i32(0)
						goto l3
					}
				}
			l22:
				if v0 != 0 {
					goto l24
				}
				v3 = i32(0)
				goto l3
			l24:
				if uint32(v0) < uint32(v10) {
					goto l25
				}
				v3 = v10
				if v0 == v10 {
					goto l3
				}
				goto l26
			l25:
				t33 := int32(int8(m.memory[uint32(v8+v0)]))
				if t33 < i32(-64) {
					goto l26
				}
				v3 = v0
			}
		l3:
			t34 := int32(load32(m.memory[int64(uint32(v6))+12:]))
			t35 := m.t0[uint(t34)].(func(int32, int32, int32) int32)(v5, v8+v3, v10-v3)
			if t35 != 0 {
				goto l14
			}
			if v9 == 0 {
				goto l27
			}
		l28:
			{
				t36 := int32(m.memory[uint32(v4)])
				m.memory[int64(uint32(v2))+32] = byte(t36)
				store64(m.memory[int64(uint32(v2))+48:], uint64(v7))
				t37 := m.fn100(v5, v6, i32(1131128), v2+i32(48))
				if t37 != 0 {
					goto l14
				}
				v4 = v4 + i32(1)
				v9 = v9 + i32(-1)
				if v9 == 0 {
					goto l27
				}
				goto l28
			}
		}
	l14:
		v0 = i32(1)
	}
l0:
	m.g0 = v2 + i32(64)
	return v0
l26:
	m.fn556(v8, v10, v0, v10, i32(1131112))
	panic("unreachable")
}
func (m *Module) fn1344(v0 int32) {
	var v1, v2, v3 int32
	t0 := m.g0
	v1 = t0 - i32(16)
	m.g0 = v1
	m.fn954(v1+i32(8), v0)
	{
		t1 := int32(m.memory[int64(uint32(v1))+8])
		v2 = t1
		if v2 == i32(255) {
			goto l0
		}
		t2 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		m.fn119(v2, t2)
	}
l0:
	{
		t3 := int32(load32(m.memory[int64(uint32(v0))+48:]))
		v2 = t3
		if v2 == i32(-1) {
			goto l1
		}
		t4 := int32(load32(m.memory[int64(uint32(v2))+4:]))
		t5 := v2
		v3 = t4
		store32(m.memory[int64(uint32(t5))+4:], uint32(v3+i32(-1)))
		if v3 != i32(1) {
			goto l1
		}
		m.fn10(v2, i32(136), i32(8))
	}
l1:
	t6 := int32(load32(m.memory[uint32(v0):]))
	t7 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	m.fn16(t6, t7)
	{
		t8 := int32(load32(m.memory[int64(uint32(v0))+40:]))
		v2 = t8
		if v2 == 0 {
			goto l2
		}
		t9 := int32(load32(m.memory[int64(uint32(v0))+44:]))
		m.fn958(v2, t9)
	}
l2:
	m.g0 = v1 + i32(16)
}
func (m *Module) fn1345(v0, v1, v2, v3 int32) {
	var v4, v5, v6, v7 int32
	v4 = i32(255)
	{
		if uint32(v2) < uint32(v3) {
			goto l4
		}
		if uint32(v2-v3) < uint32(i32(4)) {
			goto l4
		}
		v5 = v3 + i32(4)
		{
			{
				t0 := int32(load32(m.memory[uint32(v1+v3):]))
				v6 = t0
				if v6&i32(15) != 0 {
					goto l1
				}
				v7 = i32(2)
				goto l2
			}
		l1:
			if uint32(v2-v5) < uint32(i32(2)) {
				goto l4
			}
			t1 := int32(m.memory[uint32(v1+v5)])
			p2 := i32(2)
			if v6&i32(1) != 0 {
				p2 = t1 & i32(1)
			}
			v7 = p2
			v5 = v3 + i32(6)
		}
	l2:
		v3 = int32(uint32(v6) >> 3)
		v5 = v3&i32(2) + int32(uint32(v6)>>6)&i32(2) + int32(uint32(v6)>>5)&i32(2) + v3&i32(4) + int32(uint32(v6)>>10)&i32(2) + int32(uint32(v6)>>11)&i32(2) + int32(uint32(v6)>>12)&i32(2) + int32(uint32(v6)>>13)&i32(2) + int32(uint32(v6)>>7)&i32(2) + int32(uint32(v6)>>9)&i32(2) + int32(uint32(v6)>>14)&i32(2) + v5
		{
			if v6&i32(0x100000) == 0 {
				goto l3
			}
			if uint32(v2) < uint32(v5) {
				goto l4
			}
			if uint32(v2-v5) < uint32(i32(2)) {
				goto l4
			}
			t3 := int32(load16(m.memory[uint32(v1+v5):]))
			v5 = v5 + t3<<2 + i32(2)
		}
	l3:
		v3 = v5 + int32(uint32(v6)>>15)&i32(2)
		p4 := v3
		if v6&i32(0xe0000) != 0 {
			p4 = v3 + i32(2)
		}
		v6 = p4 + int32(uint32(v6)>>20)&i32(2)
		t5 := v6
		var p6 int32
		if uint32(v6) > uint32(v2) {
			p6 = 1
		}
		v6 = p6
		p7 := t5
		if v6 != 0 {
			p7 = i32(0)
		}
		v3 = p7
		p8 := v7
		if v6 != 0 {
			p8 = i32(-1)
		}
		v4 = p8
		goto l4
	}
l4:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
	m.memory[uint32(v0)] = byte(v4)
}
func (m *Module) fn1346(v0, v1, v2, v3 int32) {
	var v4, v5, v6 int32
	{
		{
			if uint32(v2) < uint32(v3) {
				goto l0
			}
			if uint32(v2-v3) < uint32(i32(4)) {
				goto l0
			}
			v4 = v3 + i32(4)
			t0 := int32(load32(m.memory[uint32(v1+v3):]))
			v5 = t0
			if v5&i32(0xffff) != 0 {
				goto l1
			}
			v1 = i32(2)
			v6 = i32(2)
			goto l2
		}
	l0:
		m.memory[uint32(v0)] = byte(i32(255))
		return
	l1:
		if uint32(v2-v4) < uint32(i32(2)) {
			goto l3
		}
		t1 := int32(m.memory[uint32(v1+v4)])
		v4 = t1
		p2 := i32(2)
		if v5&i32(1) != 0 {
			p2 = v4 & i32(1)
		}
		v1 = p2
		p3 := i32(2)
		if v5&i32(2) != 0 {
			p3 = int32(uint32(v4)>>1) & i32(1)
		}
		v6 = p3
		v4 = v3 + i32(6)
	}
l2:
	{
		t4 := int32(uint32(v5)>>20)&i32(2) + int32(uint32(v5)>>15)&i32(2) + int32(uint32(v5)>>21)&i32(2) + int32(uint32(v5)>>22)&i32(2)
		v3 = int32(uint32(v5) >> 16)
		v3 = t4 + v3&i32(2) + v3&i32(4) + int32(uint32(v5)>>18)&i32(2) + v4
		if uint32(v3) > uint32(v2) {
			m.memory[uint32(v0)] = byte(i32(255))
			return
		}
		store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
		m.memory[int64(uint32(v0))+1] = byte(v6)
		m.memory[uint32(v0)] = byte(v1)
		return
	}
l3:
	m.memory[uint32(v0)] = byte(i32(255))
}
