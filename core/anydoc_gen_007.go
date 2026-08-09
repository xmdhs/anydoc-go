package core

import (
	"math/bits"
)

func (m *Module) fn267(v0 int32) {
	var v1, v2, v3 int32
	{
		t0 := int32(load32(m.memory[uint32(v0):]))
		v1 = t0
		if v1 == 0 {
			return
		}
		t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		v2 = t1
		t2 := int32(load32(m.memory[uint32(v2+i32(-4)):]))
		v0 = t2
		v3 = v0 & i32(-8)
		t3 := v3
		v0 = v0 & i32(3)
		p4 := i32(8)
		if v0 != 0 {
			p4 = i32(4)
		}
		if uint32(t3) < uint32(p4+v1) {
			m.fn2(i32(1273840), i32(46), i32(1273888))
			panic("unreachable")
		}
		if v0 == 0 {
			goto l2
		}
		if uint32(v3) > uint32(v1+i32(39)) {
			m.fn2(i32(1273904), i32(46), i32(1273952))
			panic("unreachable")
		}
	l2:
		m.fn1(v2)
	}
}
func (m *Module) fn268(v0, v1 int32) {
	t0 := int64(load64(m.memory[int64(uint32(i32(0)))+1276864:]))
	store64(m.memory[int64(uint32(v0))+8:], uint64(t0))
	t1 := int64(load64(m.memory[int64(uint32(i32(0)))+1276856:]))
	store64(m.memory[uint32(v0):], uint64(t1))
}
func (m *Module) fn269(v0, v1 int32) {
	store32(m.memory[int64(uint32(v0))+4:], uint32(i32(40)))
	store32(m.memory[uint32(v0):], uint32(i32(1276813)))
}
func (m *Module) fn270(v0, v1 int32) {
	store32(m.memory[uint32(v0):], uint32(i32(0)))
}
func fn271(v0, v1, v2 int32) {
}
func (m *Module) fn272(v0, v1, v2, v3 int32) {
	var v4 int64
	var v5, v6 int32
	var v7 int64
	var v8 int32
	var v9 int64
	var v10, v11, v12 int32
	{
		{
			t0 := int64(load64(m.memory[uint32(v1):]))
			if t0 != i64(1) {
				goto l0
			}
			{
				t1 := int64(load64(m.memory[int64(uint32(v1))+16:]))
				v4 = t1
				if !(v4 == 0) {
					t2 := int32(load32(m.memory[int64(uint32(v1))+24:]))
					v6 = t2
					t3 := int32(load32(m.memory[uint32(v6):]))
					t4 := int64(load64(m.memory[int64(uint32(v6))+8:]))
					v7 = t4
					t5 := int32(load32(m.memory[int64(uint32(v6))+4:]))
					t6 := v7
					v8 = t5
					v9 = int64(uint32(v8))
					p7 := v9
					if uint64(v7) < uint64(v9) {
						p7 = t6
					}
					v10 = t3 + int32(p7)
					{
						t9 := v8
						p8 := i64(0xffffffff)
						if uint64(v7) < uint64(i64(0xffffffff)) {
							p8 = v7
						}
						v11 = t9 - int32(p8)
						p10 := v11
						if uint32(v11) > uint32(v8) {
							p10 = i32(0)
						}
						v8 = p10
						t11 := v8
						t12 := v4
						v9 = int64(uint32(v3))
						p13 := v9
						if uint64(v4) < uint64(v9) {
							p13 = t12
						}
						v11 = int32(p13)
						p14 := v11
						if uint32(v8) < uint32(v11) {
							p14 = t11
						}
						v5 = p14
						if v5 != i32(1) {
							goto l3
						}
						t15 := int32(m.memory[uint32(v10)])
						m.memory[uint32(v2)] = byte(t15)
						goto l4
					}
				l3:
					if v5 == 0 {
						goto l4
					}
					memory_copy(m.memory, uint32(v2), uint32(v10), uint32(v5))
				l4:
					t16 := v1
					t17 := v4
					v9 = int64(uint32(v5))
					store64(m.memory[int64(uint32(t16))+16:], uint64(t17-v9))
					store64(m.memory[int64(uint32(v6))+8:], uint64(v7+v9))
					if v5 != 0 {
						v11 = v3 + i32(-1)
						v8 = i32(1) - v5
						t18 := int32(load32(m.memory[int64(uint32(v1))+36:]))
						v10 = t18
						t19 := int32(load32(m.memory[int64(uint32(v1))+32:]))
						v6 = t19
						t20 := int32(load32(m.memory[int64(uint32(v1))+40:]))
						v3 = t20
					l7:
						{
							t21 := int32(m.memory[uint32(v2)])
							t22 := v2
							v12 = v3 | i32(2)
							v12 = t21 ^ int32(uint32((v12^i32(1))*v12&i32(0xff00))>>8)
							m.memory[uint32(t22)] = byte(v12)
							t23 := int32(load32(m.memory[int64(uint32((v12^v6)&i32(255)<<2))+1284496:]))
							t24 := v10
							v6 = int32(uint32(v6)>>8) ^ t23
							v10 = (t24+v6&i32(255))*i32(134775813) + i32(1)
							t25 := int32(load32(m.memory[int64(uint32((int32(uint32(v10)>>24)^v3&i32(255))<<2))+1284496:]))
							v3 = t25 ^ int32(uint32(v3)>>8)
							if v8 == 0 {
								goto l6
							}
							v8 = v8 + i32(1)
							v2 = v2 + i32(1)
							v12 = v11
							v11 = v11 + i32(-1)
							if v12 != 0 {
								goto l7
							}
						}
					l6:
						store32(m.memory[int64(uint32(v1))+36:], uint32(v10))
						store32(m.memory[int64(uint32(v1))+32:], uint32(v6))
						store32(m.memory[int64(uint32(v1))+40:], uint32(v3))
						goto l2
					}
					v5 = i32(0)
					goto l2
				}
				v5 = i32(0)
				goto l2
			}
		}
	l0:
		{
			t26 := int64(load64(m.memory[int64(uint32(v1))+16:]))
			v4 = t26
			if !(v4 == 0) {
				goto l8
			}
			v5 = i32(0)
			goto l2
		}
	l8:
		t27 := int32(load32(m.memory[int64(uint32(v1))+24:]))
		v6 = t27
		t28 := int32(load32(m.memory[uint32(v6):]))
		t29 := int64(load64(m.memory[int64(uint32(v6))+8:]))
		v7 = t29
		t30 := int32(load32(m.memory[int64(uint32(v6))+4:]))
		t31 := v7
		v8 = t30
		v9 = int64(uint32(v8))
		p32 := v9
		if uint64(v7) < uint64(v9) {
			p32 = t31
		}
		v10 = t28 + int32(p32)
		{
			t34 := v8
			p33 := i64(0xffffffff)
			if uint64(v7) < uint64(i64(0xffffffff)) {
				p33 = v7
			}
			v11 = t34 - int32(p33)
			p35 := v11
			if uint32(v11) > uint32(v8) {
				p35 = i32(0)
			}
			v8 = p35
			t36 := v8
			t37 := v4
			v9 = int64(uint32(v3))
			p38 := v9
			if uint64(v4) < uint64(v9) {
				p38 = t37
			}
			v3 = int32(p38)
			p39 := v3
			if uint32(v8) < uint32(v3) {
				p39 = t36
			}
			v5 = p39
			if v5 != i32(1) {
				goto l9
			}
			t40 := int32(m.memory[uint32(v10)])
			m.memory[uint32(v2)] = byte(t40)
			goto l10
		}
	l9:
		if v5 == 0 {
			goto l10
		}
		memory_copy(m.memory, uint32(v2), uint32(v10), uint32(v5))
	l10:
		t41 := v1
		t42 := v4
		v9 = int64(uint32(v5))
		store64(m.memory[int64(uint32(t41))+16:], uint64(t42-v9))
		store64(m.memory[int64(uint32(v6))+8:], uint64(v7+v9))
	}
l2:
	m.memory[uint32(v0)] = byte(i32(255))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
}
func (m *Module) fn273(v0, v1, v2, v3 int32) {
	var v4 int32
	{
		t0 := m.fn11(v3)
		v4 = t0
		if v4 == 0 {
			m.fn7(i32(1), v3)
			panic("unreachable")
		}
		if v3 == 0 {
			goto l1
		}
		memory_copy(m.memory, uint32(v4), uint32(v2), uint32(v3))
	l1:
		t1 := m.fn11(i32(12))
		v2 = t1
		if v2 == 0 {
			m.fn30(i32(4), i32(12))
			panic("unreachable")
		}
		store32(m.memory[int64(uint32(v2))+8:], uint32(v3))
		store32(m.memory[int64(uint32(v2))+4:], uint32(v4))
		store32(m.memory[uint32(v2):], uint32(v3))
		t2 := m.fn11(i32(12))
		v3 = t2
		if v3 == 0 {
			m.fn30(i32(4), i32(12))
			panic("unreachable")
		}
		m.memory[int64(uint32(v3))+8] = byte(v1)
		store32(m.memory[int64(uint32(v3))+4:], uint32(i32(1092944)))
		store32(m.memory[uint32(v3):], uint32(v2))
		store64(m.memory[uint32(v0):], uint64(int64(uint32(v3))<<32|i64(3)))
		return
	}
}
func (m *Module) fn274(v0, v1, v2 int32) {
	var v3, v4, v5 int32
	t0 := int64(load64(m.memory[uint32(v0):]))
	store64(m.memory[uint32(v0):], uint64(t0+int64(uint32(v2))))
	t1 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	v3 = t1 ^ i32(-1)
	if uint32(v2) < uint32(i32(64)) {
		goto l0
	}
l1:
	{
		t2 := int32(m.memory[uint32(v1+i32(62))])
		t3 := int32(load32(m.memory[int64(uint32(t2<<2))+1123756:]))
		t4 := int32(m.memory[uint32(v1+i32(63))])
		t5 := int32(load32(m.memory[int64(uint32(t4<<2))+1122732:]))
		t6 := int32(m.memory[uint32(v1+i32(61))])
		t7 := int32(load32(m.memory[int64(uint32(t6<<2))+1124780:]))
		t8 := int32(m.memory[uint32(v1+i32(60))])
		t9 := int32(load32(m.memory[int64(uint32(t8<<2))+1125804:]))
		t10 := int32(m.memory[uint32(v1+i32(59))])
		t11 := int32(load32(m.memory[int64(uint32(t10<<2))+1126828:]))
		t12 := int32(m.memory[uint32(v1+i32(58))])
		t13 := int32(load32(m.memory[int64(uint32(t12<<2))+1127852:]))
		t14 := int32(m.memory[uint32(v1+i32(57))])
		t15 := int32(load32(m.memory[int64(uint32(t14<<2))+1128876:]))
		t16 := int32(m.memory[uint32(v1+i32(56))])
		t17 := int32(load32(m.memory[int64(uint32(t16<<2))+1129900:]))
		t18 := int32(m.memory[uint32(v1+i32(55))])
		t19 := int32(load32(m.memory[int64(uint32(t18<<2))+1130924:]))
		t20 := int32(m.memory[uint32(v1+i32(54))])
		t21 := int32(load32(m.memory[int64(uint32(t20<<2))+1131948:]))
		t22 := int32(m.memory[uint32(v1+i32(53))])
		t23 := int32(load32(m.memory[int64(uint32(t22<<2))+1132972:]))
		t24 := int32(m.memory[uint32(v1+i32(52))])
		t25 := int32(load32(m.memory[int64(uint32(t24<<2))+1133996:]))
		t26 := int32(m.memory[uint32(v1+i32(47))])
		t27 := int32(load32(m.memory[int64(uint32(t26<<2))+1122732:]))
		t28 := int32(m.memory[uint32(v1+i32(46))])
		t29 := int32(load32(m.memory[int64(uint32(t28<<2))+1123756:]))
		t30 := int32(m.memory[uint32(v1+i32(45))])
		t31 := int32(load32(m.memory[int64(uint32(t30<<2))+1124780:]))
		t32 := int32(m.memory[uint32(v1+i32(44))])
		t33 := int32(load32(m.memory[int64(uint32(t32<<2))+1125804:]))
		t34 := int32(m.memory[uint32(v1+i32(43))])
		t35 := int32(load32(m.memory[int64(uint32(t34<<2))+1126828:]))
		t36 := int32(m.memory[uint32(v1+i32(42))])
		t37 := int32(load32(m.memory[int64(uint32(t36<<2))+1127852:]))
		t38 := int32(m.memory[uint32(v1+i32(41))])
		t39 := int32(load32(m.memory[int64(uint32(t38<<2))+1128876:]))
		t40 := int32(m.memory[uint32(v1+i32(40))])
		t41 := int32(load32(m.memory[int64(uint32(t40<<2))+1129900:]))
		t42 := int32(m.memory[uint32(v1+i32(39))])
		t43 := int32(load32(m.memory[int64(uint32(t42<<2))+1130924:]))
		t44 := int32(m.memory[uint32(v1+i32(38))])
		t45 := int32(load32(m.memory[int64(uint32(t44<<2))+1131948:]))
		t46 := int32(m.memory[uint32(v1+i32(37))])
		t47 := int32(load32(m.memory[int64(uint32(t46<<2))+1132972:]))
		t48 := int32(m.memory[uint32(v1+i32(36))])
		t49 := int32(load32(m.memory[int64(uint32(t48<<2))+1133996:]))
		t50 := int32(m.memory[uint32(v1+i32(31))])
		t51 := int32(load32(m.memory[int64(uint32(t50<<2))+1122732:]))
		t52 := int32(m.memory[uint32(v1+i32(30))])
		t53 := int32(load32(m.memory[int64(uint32(t52<<2))+1123756:]))
		t54 := int32(m.memory[uint32(v1+i32(29))])
		t55 := int32(load32(m.memory[int64(uint32(t54<<2))+1124780:]))
		t56 := int32(m.memory[uint32(v1+i32(28))])
		t57 := int32(load32(m.memory[int64(uint32(t56<<2))+1125804:]))
		t58 := int32(m.memory[uint32(v1+i32(27))])
		t59 := int32(load32(m.memory[int64(uint32(t58<<2))+1126828:]))
		t60 := int32(m.memory[uint32(v1+i32(26))])
		t61 := int32(load32(m.memory[int64(uint32(t60<<2))+1127852:]))
		t62 := int32(m.memory[uint32(v1+i32(25))])
		t63 := int32(load32(m.memory[int64(uint32(t62<<2))+1128876:]))
		t64 := int32(m.memory[uint32(v1+i32(24))])
		t65 := int32(load32(m.memory[int64(uint32(t64<<2))+1129900:]))
		t66 := int32(m.memory[uint32(v1+i32(23))])
		t67 := int32(load32(m.memory[int64(uint32(t66<<2))+1130924:]))
		t68 := int32(m.memory[uint32(v1+i32(22))])
		t69 := int32(load32(m.memory[int64(uint32(t68<<2))+1131948:]))
		t70 := int32(m.memory[uint32(v1+i32(21))])
		t71 := int32(load32(m.memory[int64(uint32(t70<<2))+1132972:]))
		t72 := int32(m.memory[uint32(v1+i32(20))])
		t73 := int32(load32(m.memory[int64(uint32(t72<<2))+1133996:]))
		t74 := int32(m.memory[uint32(v1+i32(15))])
		t75 := int32(load32(m.memory[int64(uint32(t74<<2))+1122732:]))
		t76 := int32(m.memory[uint32(v1+i32(14))])
		t77 := int32(load32(m.memory[int64(uint32(t76<<2))+1123756:]))
		t78 := int32(m.memory[uint32(v1+i32(13))])
		t79 := int32(load32(m.memory[int64(uint32(t78<<2))+1124780:]))
		t80 := int32(m.memory[uint32(v1+i32(12))])
		t81 := int32(load32(m.memory[int64(uint32(t80<<2))+1125804:]))
		t82 := int32(m.memory[uint32(v1+i32(11))])
		t83 := int32(load32(m.memory[int64(uint32(t82<<2))+1126828:]))
		t84 := int32(m.memory[uint32(v1+i32(10))])
		t85 := int32(load32(m.memory[int64(uint32(t84<<2))+1127852:]))
		t86 := int32(m.memory[uint32(v1+i32(9))])
		t87 := int32(load32(m.memory[int64(uint32(t86<<2))+1128876:]))
		t88 := int32(m.memory[uint32(v1+i32(8))])
		t89 := int32(load32(m.memory[int64(uint32(t88<<2))+1129900:]))
		t90 := int32(m.memory[uint32(v1+i32(7))])
		t91 := int32(load32(m.memory[int64(uint32(t90<<2))+1130924:]))
		t92 := int32(m.memory[uint32(v1+i32(6))])
		t93 := int32(load32(m.memory[int64(uint32(t92<<2))+1131948:]))
		t94 := int32(m.memory[uint32(v1+i32(5))])
		t95 := int32(load32(m.memory[int64(uint32(t94<<2))+1132972:]))
		t96 := int32(m.memory[uint32(v1+i32(4))])
		t97 := int32(load32(m.memory[int64(uint32(t96<<2))+1133996:]))
		t98 := int32(m.memory[uint32(v1+i32(3))])
		t99 := int32(load32(m.memory[int64(uint32((int32(uint32(v3)>>24)^t98)<<2))+1135020:]))
		t100 := int32(m.memory[uint32(v1+i32(2))])
		t101 := int32(load32(m.memory[int64(uint32((int32(uint32(v3)>>16)&i32(255)^t100)<<2))+1136044:]))
		t102 := int32(m.memory[uint32(v1+i32(1))])
		t103 := int32(load32(m.memory[int64(uint32((int32(uint32(v3)>>8)&i32(255)^t102)<<2))+1137068:]))
		t104 := int32(m.memory[uint32(v1)])
		t105 := int32(load32(m.memory[int64(uint32((v3&i32(255)^t104)<<2))+1138092:]))
		t106 := t3 ^ t5 ^ t7 ^ t9 ^ t11 ^ t13 ^ t15 ^ t17 ^ t19 ^ t21 ^ t23 ^ t25
		t107 := t27 ^ t29 ^ t31 ^ t33 ^ t35 ^ t37 ^ t39 ^ t41 ^ t43 ^ t45 ^ t47 ^ t49
		t108 := t51 ^ t53 ^ t55 ^ t57 ^ t59 ^ t61 ^ t63 ^ t65 ^ t67 ^ t69 ^ t71 ^ t73
		v3 = t75 ^ t77 ^ t79 ^ t81 ^ t83 ^ t85 ^ t87 ^ t89 ^ t91 ^ t93 ^ t95 ^ t97 ^ t99 ^ t101 ^ t103 ^ t105
		t109 := int32(m.memory[uint32(v1+i32(19))])
		t110 := int32(load32(m.memory[int64(uint32((int32(uint32(v3)>>24)^t109)<<2))+1135020:]))
		t111 := int32(m.memory[uint32(v1+i32(18))])
		t112 := int32(load32(m.memory[int64(uint32((int32(uint32(v3)>>16)&i32(255)^t111)<<2))+1136044:]))
		t113 := int32(m.memory[uint32(v1+i32(17))])
		t114 := int32(load32(m.memory[int64(uint32((int32(uint32(v3)>>8)&i32(255)^t113)<<2))+1137068:]))
		t115 := int32(m.memory[uint32(v1+i32(16))])
		t116 := int32(load32(m.memory[int64(uint32((v3&i32(255)^t115)<<2))+1138092:]))
		v3 = t108 ^ t110 ^ t112 ^ t114 ^ t116
		t117 := int32(m.memory[uint32(v1+i32(35))])
		t118 := int32(load32(m.memory[int64(uint32((int32(uint32(v3)>>24)^t117)<<2))+1135020:]))
		t119 := int32(m.memory[uint32(v1+i32(34))])
		t120 := int32(load32(m.memory[int64(uint32((int32(uint32(v3)>>16)&i32(255)^t119)<<2))+1136044:]))
		t121 := int32(m.memory[uint32(v1+i32(33))])
		t122 := int32(load32(m.memory[int64(uint32((int32(uint32(v3)>>8)&i32(255)^t121)<<2))+1137068:]))
		t123 := int32(m.memory[uint32(v1+i32(32))])
		t124 := int32(load32(m.memory[int64(uint32((v3&i32(255)^t123)<<2))+1138092:]))
		v3 = t107 ^ t118 ^ t120 ^ t122 ^ t124
		t125 := int32(m.memory[uint32(v1+i32(51))])
		t126 := int32(load32(m.memory[int64(uint32((int32(uint32(v3)>>24)^t125)<<2))+1135020:]))
		t127 := int32(m.memory[uint32(v1+i32(50))])
		t128 := int32(load32(m.memory[int64(uint32((int32(uint32(v3)>>16)&i32(255)^t127)<<2))+1136044:]))
		t129 := int32(m.memory[uint32(v1+i32(49))])
		t130 := int32(load32(m.memory[int64(uint32((int32(uint32(v3)>>8)&i32(255)^t129)<<2))+1137068:]))
		t131 := int32(m.memory[uint32(v1+i32(48))])
		t132 := int32(load32(m.memory[int64(uint32((v3&i32(255)^t131)<<2))+1138092:]))
		v3 = t106 ^ t126 ^ t128 ^ t130 ^ t132
		v1 = v1 + i32(64)
		v2 = v2 + i32(-64)
		if uint32(v2) > uint32(i32(63)) {
			goto l1
		}
	}
l0:
	if v2 == 0 {
		goto l2
	}
	v4 = v2 & i32(3)
	if v4 != 0 {
		goto l3
	}
	v5 = v1
	goto l4
l3:
	v5 = v1
l5:
	{
		t133 := int32(m.memory[uint32(v5)])
		t134 := int32(load32(m.memory[int64(uint32((t133^v3)&i32(255)<<2))+1122732:]))
		v3 = t134 ^ int32(uint32(v3)>>8)
		v5 = v5 + i32(1)
		v4 = v4 + i32(-1)
		if v4 != 0 {
			goto l5
		}
	}
l4:
	if uint32(v2) < uint32(i32(4)) {
		goto l2
	}
	v2 = v1 + v2
l6:
	{
		t135 := int32(m.memory[uint32(v5+i32(3))])
		t136 := int32(m.memory[uint32(v5+i32(2))])
		t137 := int32(m.memory[uint32(v5+i32(1))])
		t138 := int32(m.memory[uint32(v5)])
		t139 := int32(load32(m.memory[int64(uint32((t138^v3)&i32(255)<<2))+1122732:]))
		v1 = t139 ^ int32(uint32(v3)>>8)
		t140 := int32(load32(m.memory[int64(uint32((t137^v1)&i32(255)<<2))+1122732:]))
		v1 = t140 ^ int32(uint32(v1)>>8)
		t141 := int32(load32(m.memory[int64(uint32((t136^v1)&i32(255)<<2))+1122732:]))
		v1 = t141 ^ int32(uint32(v1)>>8)
		t142 := int32(load32(m.memory[int64(uint32((t135^v1)&i32(255)<<2))+1122732:]))
		v3 = t142 ^ int32(uint32(v1)>>8)
		v5 = v5 + i32(4)
		if v5 != v2 {
			goto l6
		}
	}
l2:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v3^i32(-1)))
}
func (m *Module) fn275(v0, v1, v2, v3, v4, v5, v6 int32) {
	var v7, v8, v9 int32
	var v10, v11 int64
	var v12, v13, v14, v15, v16, v17, v18, v19, v20, v21, v22, v23, v24 int32
	var v25, v26 int64
	var v27, v28, v29 int32
	var v30, v31 int64
	var v32, v33, v34 int32
	t0 := m.g0
	v7 = t0 - i32(32)
	m.g0 = v7
	store32(m.memory[int64(uint32(v1))+48:], uint32(v5))
	store32(m.memory[int64(uint32(v1))+36:], uint32(v3))
	store32(m.memory[int64(uint32(v1))+44:], uint32(v4))
	store32(m.memory[int64(uint32(v1))+32:], uint32(v2))
	{
		t1 := int32(load32(m.memory[int64(uint32(v1))+60:]))
		v8 = t1
		t2 := int32(m.memory[uint32(v8)])
		v9 = t2
		if v9 != i32(12) {
			goto l0
		}
		v9 = i32(13)
		m.memory[uint32(v8)] = byte(i32(13))
	}
l0:
	t3 := int64(load64(m.memory[int64(uint32(v1))+24:]))
	v10 = t3
	t4 := int64(load64(m.memory[int64(uint32(v1))+16:]))
	v11 = t4
	store32(m.memory[int64(uint32(v8))+116:], uint32(v5))
	store32(m.memory[int64(uint32(v8))+112:], uint32(v3))
	store32(m.memory[int64(uint32(v8))+80:], uint32(i32(0)))
	store32(m.memory[int64(uint32(v8))+76:], uint32(v5))
	store32(m.memory[int64(uint32(v8))+72:], uint32(v4))
	store32(m.memory[int64(uint32(v8))+56:], uint32(v2))
	m.memory[int64(uint32(v8))+4] = byte(v6)
	store32(m.memory[int64(uint32(v8))+60:], uint32(v2+v3))
	v12 = v8 + i32(65)
	v13 = v8 + i32(8)
	v14 = v8 + i32(13828)
	v15 = v8 + i32(13188)
	v16 = v8 + i32(164)
	v17 = v8 + i32(5492)
	v18 = v8 + i32(10820)
	v19 = v8 + i32(72)
l18:
	v5 = v9 & i32(255)
	v9 = i32(18)
	v20 = i32(29)
	v3 = i32(1)
	switch v5 {
	case 1:
		t629 := int64(load64(m.memory[int64(uint32(v8))+48:]))
		v25 = t629
		{
			t630 := int32(m.memory[int64(uint32(v8))+64])
			v9 = t630
			if uint32(v9) > uint32(i32(15)) {
				goto l264
			}
			v3 = i32(0)
			{
				t631 := int32(load32(m.memory[int64(uint32(v8))+56:]))
				v5 = t631
				t632 := int32(load32(m.memory[int64(uint32(v8))+60:]))
				t633 := v5
				v29 = t632
				if t633 != v29 {
					goto l265
				}
				v20 = i32(1)
				goto l30
			}
		l265:
			v20 = i32(1)
			t634 := v8
			v27 = v5 + i32(1)
			store32(m.memory[int64(uint32(t634))+56:], uint32(v27))
			t635 := v8
			v30 = int64(uint32(v9))
			v26 = v30 + i64(8)
			m.memory[int64(uint32(t635))+64] = byte(v26)
			t636 := int64(m.memory[uint32(v5)])
			t637 := v8
			v25 = i64_shl(t636, v30) | v25
			store64(m.memory[int64(uint32(t637))+48:], uint64(v25))
			if uint32(v9) > uint32(i32(7)) {
				goto l264
			}
			if v27 == v29 {
				goto l30
			}
			store32(m.memory[int64(uint32(v8))+56:], uint32(v5+i32(2)))
			m.memory[int64(uint32(v8))+64] = byte(v9 | i32(16))
			t638 := int64(m.memory[int64(uint32(v5))+1])
			t639 := v8
			v25 = i64_shl(t638, v26) | v25
			store64(m.memory[int64(uint32(t639))+48:], uint64(v25))
		}
	l264:
		t640 := v8
		v9 = int32(v25)
		store32(m.memory[int64(uint32(t640))+120:], uint32(v9))
		{
			if v9&i32(255) != i32(8) {
				store32(m.memory[int64(uint32(v8))+136:], uint32(i32(27)))
				store32(m.memory[int64(uint32(v8))+132:], uint32(i32(1065362)))
				goto l57
			}
			if v9&i32(57344) != 0 {
				store32(m.memory[int64(uint32(v8))+136:], uint32(i32(25)))
				store32(m.memory[int64(uint32(v8))+132:], uint32(i32(1052858)))
				goto l57
			}
			{
				t641 := int32(load32(m.memory[int64(uint32(v8))+140:]))
				v5 = t641
				if v5 == 0 {
					goto l268
				}
				store32(m.memory[uint32(v5):], uint32(int32(uint32(v9)>>8)&i32(1)))
				t642 := int32(load32(m.memory[int64(uint32(v8))+120:]))
				v9 = t642
			}
		l268:
			if v9&i32(512) == 0 {
				goto l269
			}
			t643 := int32(m.memory[int64(uint32(v8))+3])
			if t643&i32(4) == 0 {
				goto l269
			}
			t644 := int32(load32(m.memory[int64(uint32(v8))+124:]))
			v9 = t644
			t645 := int64(load64(m.memory[int64(uint32(v8))+48:]))
			store16(m.memory[int64(uint32(v7))+16:], uint16(t645))
			t646 := m.fn910(v9, v7+i32(16), i32(2))
			store32(m.memory[int64(uint32(v8))+124:], uint32(t646))
			goto l269
		}
	l269:
		m.memory[int64(uint32(v8))+64] = byte(i32(0))
		store64(m.memory[int64(uint32(v8))+48:], uint64(i64(0)))
		v9 = i32(2)
		goto l18
	case 2:
		{
			t604 := int32(m.memory[int64(uint32(v8))+64])
			v9 = t604
			if uint32(v9) > uint32(i32(31)) {
				goto l261
			}
			v3 = i32(0)
			v20 = i32(2)
			t605 := int32(load32(m.memory[int64(uint32(v8))+56:]))
			v5 = t605
			t606 := int32(load32(m.memory[int64(uint32(v8))+60:]))
			t607 := v5
			v29 = t606
			if t607 == v29 {
				goto l30
			}
			t608 := v8
			v27 = v5 + i32(1)
			store32(m.memory[int64(uint32(t608))+56:], uint32(v27))
			t609 := v8
			v25 = int64(uint32(v9))
			v30 = v25 + i64(8)
			m.memory[int64(uint32(t609))+64] = byte(v30)
			t610 := int64(m.memory[uint32(v5)])
			t611 := int64(load64(m.memory[int64(uint32(v8))+48:]))
			t612 := v8
			v26 = i64_shl(t610, v25) | t611
			store64(m.memory[int64(uint32(t612))+48:], uint64(v26))
			if uint32(v9) > uint32(i32(23)) {
				goto l261
			}
			if v27 == v29 {
				goto l30
			}
			v20 = i32(2)
			t613 := v8
			v27 = v5 + i32(2)
			store32(m.memory[int64(uint32(t613))+56:], uint32(v27))
			t614 := v8
			v31 = v25 + i64(16)
			m.memory[int64(uint32(t614))+64] = byte(v31)
			t615 := int64(m.memory[int64(uint32(v5))+1])
			t616 := v8
			v30 = i64_shl(t615, v30) | v26
			store64(m.memory[int64(uint32(t616))+48:], uint64(v30))
			if uint32(v9) > uint32(i32(15)) {
				goto l261
			}
			if v27 == v29 {
				goto l30
			}
			t617 := v8
			v27 = v5 + i32(3)
			store32(m.memory[int64(uint32(t617))+56:], uint32(v27))
			t618 := v8
			v25 = v25 + i64(24)
			m.memory[int64(uint32(t618))+64] = byte(v25)
			t619 := int64(m.memory[int64(uint32(v5))+2])
			t620 := v8
			v30 = i64_shl(t619, v31) | v30
			store64(m.memory[int64(uint32(t620))+48:], uint64(v30))
			if uint32(v9) > uint32(i32(7)) {
				goto l261
			}
			if v27 == v29 {
				goto l30
			}
			store32(m.memory[int64(uint32(v8))+56:], uint32(v5+i32(4)))
			m.memory[int64(uint32(v8))+64] = byte(v9 | i32(32))
			t621 := int64(m.memory[int64(uint32(v5))+3])
			store64(m.memory[int64(uint32(v8))+48:], uint64(i64_shl(t621, v25)|v30))
		}
	l261:
		{
			t622 := int32(load32(m.memory[int64(uint32(v8))+140:]))
			v9 = t622
			if v9 == 0 {
				goto l262
			}
			t623 := int64(load64(m.memory[int64(uint32(v8))+48:]))
			store32(m.memory[int64(uint32(v9))+4:], uint32(t623))
		}
	l262:
		{
			t624 := int32(m.memory[int64(uint32(v8))+121])
			if t624&i32(2) == 0 {
				goto l263
			}
			t625 := int32(m.memory[int64(uint32(v8))+3])
			if t625&i32(4) == 0 {
				goto l263
			}
			t626 := int64(load64(m.memory[int64(uint32(v8))+48:]))
			store32(m.memory[int64(uint32(v7))+16:], uint32(t626))
			t627 := int32(load32(m.memory[int64(uint32(v8))+124:]))
			t628 := m.fn910(t627, v7+i32(16), i32(4))
			store32(m.memory[int64(uint32(v8))+124:], uint32(t628))
		}
	l263:
		m.memory[int64(uint32(v8))+64] = byte(i32(0))
		store64(m.memory[int64(uint32(v8))+48:], uint64(i64(0)))
		v9 = i32(3)
		goto l18
	case 3:
		{
			t585 := int32(m.memory[int64(uint32(v8))+64])
			v9 = t585
			if uint32(v9) > uint32(i32(15)) {
				goto l258
			}
			v3 = i32(0)
			v20 = i32(3)
			t586 := int32(load32(m.memory[int64(uint32(v8))+56:]))
			v5 = t586
			t587 := int32(load32(m.memory[int64(uint32(v8))+60:]))
			t588 := v5
			v29 = t587
			if t588 == v29 {
				goto l30
			}
			t589 := v8
			v27 = v5 + i32(1)
			store32(m.memory[int64(uint32(t589))+56:], uint32(v27))
			t590 := v8
			v25 = int64(uint32(v9))
			v30 = v25 + i64(8)
			m.memory[int64(uint32(t590))+64] = byte(v30)
			t591 := int64(m.memory[uint32(v5)])
			t592 := int64(load64(m.memory[int64(uint32(v8))+48:]))
			t593 := v8
			v25 = i64_shl(t591, v25) | t592
			store64(m.memory[int64(uint32(t593))+48:], uint64(v25))
			if uint32(v9) > uint32(i32(7)) {
				goto l258
			}
			if v27 == v29 {
				goto l30
			}
			store32(m.memory[int64(uint32(v8))+56:], uint32(v5+i32(2)))
			m.memory[int64(uint32(v8))+64] = byte(v9 | i32(16))
			t594 := int64(m.memory[int64(uint32(v5))+1])
			store64(m.memory[int64(uint32(v8))+48:], uint64(i64_shl(t594, v30)|v25))
		}
	l258:
		{
			t595 := int32(load32(m.memory[int64(uint32(v8))+140:]))
			v9 = t595
			if v9 == 0 {
				goto l259
			}
			t596 := int32(m.memory[int64(uint32(v8))+48])
			store32(m.memory[int64(uint32(v9))+8:], uint32(t596))
			t597 := int32(load32(m.memory[int64(uint32(v8))+140:]))
			t598 := int64(load64(m.memory[int64(uint32(v8))+48:]))
			store32(m.memory[int64(uint32(t597))+12:], uint32(int64(uint64(t598)>>8)))
		}
	l259:
		{
			t599 := int32(m.memory[int64(uint32(v8))+121])
			if t599&i32(2) == 0 {
				goto l260
			}
			t600 := int32(m.memory[int64(uint32(v8))+3])
			if t600&i32(4) == 0 {
				goto l260
			}
			t601 := int64(load64(m.memory[int64(uint32(v8))+48:]))
			store16(m.memory[int64(uint32(v7))+16:], uint16(t601))
			t602 := int32(load32(m.memory[int64(uint32(v8))+124:]))
			t603 := m.fn910(t602, v7+i32(16), i32(2))
			store32(m.memory[int64(uint32(v8))+124:], uint32(t603))
		}
	l260:
		m.memory[int64(uint32(v8))+64] = byte(i32(0))
		store64(m.memory[int64(uint32(v8))+48:], uint64(i64(0)))
		v9 = i32(4)
		goto l18
	case 4:
		{
			{
				t257 := int32(load32(m.memory[int64(uint32(v8))+120:]))
				v9 = t257
				if v9&i32(1024) == 0 {
					v9 = i32(5)
					t273 := int32(load32(m.memory[int64(uint32(v8))+140:]))
					v5 = t273
					if v5 == 0 {
						goto l18
					}
					store32(m.memory[int64(uint32(v5))+16:], uint32(i32(0)))
					goto l18
				}
				t258 := int64(load64(m.memory[int64(uint32(v8))+48:]))
				v25 = t258
				{
					t259 := int32(m.memory[int64(uint32(v8))+64])
					v5 = t259
					if uint32(v5) > uint32(i32(15)) {
						goto l154
					}
					v3 = i32(0)
					v20 = i32(4)
					t260 := int32(load32(m.memory[int64(uint32(v8))+56:]))
					v29 = t260
					t261 := int32(load32(m.memory[int64(uint32(v8))+60:]))
					t262 := v29
					v27 = t261
					if t262 == v27 {
						goto l30
					}
					t263 := v8
					v28 = v29 + i32(1)
					store32(m.memory[int64(uint32(t263))+56:], uint32(v28))
					t264 := v8
					v30 = int64(uint32(v5))
					v26 = v30 + i64(8)
					m.memory[int64(uint32(t264))+64] = byte(v26)
					t265 := int64(m.memory[uint32(v29)])
					t266 := v8
					v25 = i64_shl(t265, v30) | v25
					store64(m.memory[int64(uint32(t266))+48:], uint64(v25))
					if uint32(v5) > uint32(i32(7)) {
						goto l154
					}
					if v28 == v27 {
						goto l30
					}
					store32(m.memory[int64(uint32(v8))+56:], uint32(v29+i32(2)))
					m.memory[int64(uint32(v8))+64] = byte(v5 | i32(16))
					t267 := int64(m.memory[int64(uint32(v29))+1])
					t268 := v8
					v25 = i64_shl(t267, v26) | v25
					store64(m.memory[int64(uint32(t268))+48:], uint64(v25))
				}
			l154:
				t269 := v8
				v5 = int32(v25)
				store32(m.memory[int64(uint32(t269))+88:], uint32(v5))
				{
					t270 := int32(load32(m.memory[int64(uint32(v8))+140:]))
					v3 = t270
					if v3 == 0 {
						goto l155
					}
					store32(m.memory[int64(uint32(v3))+20:], uint32(v5))
					t271 := int32(load32(m.memory[int64(uint32(v8))+120:]))
					v9 = t271
				}
			l155:
				if v9&i32(512) == 0 {
					goto l156
				}
				t272 := int32(m.memory[int64(uint32(v8))+3])
				if t272&i32(4) != 0 {
					goto l157
				}
				goto l156
			}
		l157:
			t274 := int64(load64(m.memory[int64(uint32(v8))+48:]))
			store16(m.memory[int64(uint32(v7))+16:], uint16(t274))
			t275 := int32(load32(m.memory[int64(uint32(v8))+124:]))
			t276 := m.fn910(t275, v7+i32(16), i32(2))
			store32(m.memory[int64(uint32(v8))+124:], uint32(t276))
		}
	l156:
		m.memory[int64(uint32(v8))+64] = byte(i32(0))
		store64(m.memory[int64(uint32(v8))+48:], uint64(i64(0)))
		v9 = i32(5)
		goto l18
	case 5:
		{
			t226 := int32(load32(m.memory[int64(uint32(v8))+120:]))
			v3 = t226
			if v3&i32(1024) == 0 {
				goto l147
			}
			{
				t227 := int32(load32(m.memory[int64(uint32(v8))+60:]))
				t228 := int32(load32(m.memory[int64(uint32(v8))+56:]))
				v29 = t228
				v5 = t227 - v29
				t229 := int32(load32(m.memory[int64(uint32(v8))+88:]))
				t230 := v5
				v9 = t229
				p231 := v9
				if uint32(v5) < uint32(v9) {
					p231 = t230
				}
				v5 = p231
				if v5 == 0 {
					goto l148
				}
				{
					t232 := int32(load32(m.memory[int64(uint32(v8))+140:]))
					v20 = t232
					if v20 == 0 {
						goto l149
					}
					t233 := int32(load32(m.memory[int64(uint32(v20))+16:]))
					v27 = t233
					if v27 == 0 {
						goto l149
					}
					{
						t234 := int32(load32(m.memory[int64(uint32(v20))+24:]))
						t235 := v5
						v3 = t234
						t236 := int32(load32(m.memory[int64(uint32(v20))+20:]))
						t237 := v3
						v9 = t236 - v9
						v20 = t237 - v9
						p238 := v20
						if uint32(v20) > uint32(v3) {
							p238 = i32(0)
						}
						v20 = p238
						p239 := v20
						if uint32(v5) < uint32(v20) {
							p239 = t235
						}
						v20 = p239
						if v20 == 0 {
							goto l150
						}
						t241 := v27
						p240 := v9
						if uint32(v3) < uint32(v9) {
							p240 = v3
						}
						memory_copy(m.memory, uint32(t241+p240), uint32(v29), uint32(v20))
					}
				l150:
					t242 := int32(load32(m.memory[int64(uint32(v8))+120:]))
					v3 = t242
				}
			l149:
				{
					if v3&i32(512) == 0 {
						goto l151
					}
					t243 := int32(m.memory[int64(uint32(v8))+3])
					if t243&i32(4) == 0 {
						goto l151
					}
					{
						t244 := int32(load32(m.memory[int64(uint32(v8))+60:]))
						t245 := int32(load32(m.memory[int64(uint32(v8))+56:]))
						t246 := v5
						v9 = t245
						v3 = t244 - v9
						if uint32(t246) > uint32(v3) {
							m.fn127(i32(0), v5, v3, i32(1284004))
							panic("unreachable")
						}
						t247 := int32(load32(m.memory[int64(uint32(v8))+124:]))
						t248 := m.fn910(t247, v9, v5)
						store32(m.memory[int64(uint32(v8))+124:], uint32(t248))
						goto l151
					}
				}
			l151:
				t249 := int32(load32(m.memory[int64(uint32(v8))+112:]))
				store32(m.memory[int64(uint32(v8))+112:], uint32(t249-v5))
				t250 := int32(load32(m.memory[int64(uint32(v8))+88:]))
				t251 := v8
				v9 = t250 - v5
				store32(m.memory[int64(uint32(t251))+88:], uint32(v9))
				t252 := int32(load32(m.memory[int64(uint32(v8))+60:]))
				t253 := v8
				v3 = t252
				t254 := int32(load32(m.memory[int64(uint32(v8))+56:]))
				t255 := v3
				v5 = t254 + v5
				p256 := v5
				if uint32(v3) < uint32(v5) {
					p256 = t255
				}
				store32(m.memory[int64(uint32(t253))+56:], uint32(p256))
			}
		l148:
			if v9 == 0 {
				goto l147
			}
			v3 = i32(0)
			v20 = i32(5)
			goto l30
		}
	l147:
		store32(m.memory[int64(uint32(v8))+88:], uint32(i32(0)))
		v9 = i32(6)
		goto l18
	case 6:
		{
			t196 := int32(load32(m.memory[int64(uint32(v8))+120:]))
			v27 = t196
			if v27&i32(2048) == 0 {
				t198 := int32(load32(m.memory[int64(uint32(v8))+140:]))
				v9 = t198
				if v9 == 0 {
					goto l136
				}
				store32(m.memory[int64(uint32(v9))+28:], uint32(i32(0)))
				goto l136
			}
			v20 = i32(6)
			t197 := int32(load32(m.memory[int64(uint32(v8))+112:]))
			if t197 != 0 {
				{
					t199 := int32(load32(m.memory[int64(uint32(v8))+60:]))
					v3 = t199
					t200 := int32(load32(m.memory[int64(uint32(v8))+56:]))
					t201 := v3
					v5 = t200
					if t201 != v5 {
						goto l137
					}
					v29 = i32(0)
					goto l138
				}
			l137:
				v29 = v3 - v5
				v9 = i32(0)
			l140:
				{
					t202 := int32(m.memory[uint32(v5+v9)])
					if t202 == 0 {
						goto l139
					}
					t203 := v5
					v9 = v9 + i32(1)
					if t203+v9 != v3 {
						goto l140
					}
					goto l138
				}
			l139:
				v29 = v9 + i32(1)
			l138:
				{
					t204 := int32(load32(m.memory[int64(uint32(v8))+140:]))
					v9 = t204
					if v9 == 0 {
						goto l141
					}
					t205 := int32(load32(m.memory[int64(uint32(v9))+28:]))
					v28 = t205
					if v28 == 0 {
						goto l141
					}
					{
						t206 := int32(load32(m.memory[int64(uint32(v9))+32:]))
						v3 = t206
						t207 := int32(load32(m.memory[int64(uint32(v8))+88:]))
						t208 := v3
						v9 = t207
						if uint32(t208) < uint32(v9) {
							m.fn146(i32(1284020), i32(18), i32(1284040))
							panic("unreachable")
						}
						{
							v3 = v3 - v9
							p209 := v29
							if uint32(v3) < uint32(v29) {
								p209 = v3
							}
							v3 = p209
							if v3 == 0 {
								goto l143
							}
							memory_copy(m.memory, uint32(v28+v9), uint32(v5), uint32(v3))
						}
					l143:
						t210 := int32(load32(m.memory[int64(uint32(v8))+88:]))
						store32(m.memory[int64(uint32(v8))+88:], uint32(t210+v3))
						t211 := int32(load32(m.memory[int64(uint32(v8))+120:]))
						v27 = t211
						goto l141
					}
				}
			l141:
				{
					if v27&i32(512) == 0 {
						goto l144
					}
					t212 := int32(m.memory[int64(uint32(v8))+3])
					if t212&i32(4) == 0 {
						goto l144
					}
					t213 := int32(load32(m.memory[int64(uint32(v8))+124:]))
					t214 := m.fn910(t213, v5, v29)
					store32(m.memory[int64(uint32(v8))+124:], uint32(t214))
				}
			l144:
				{
					{
						if v29 != 0 {
							goto l145
						}
						t215 := int32(load32(m.memory[int64(uint32(v8))+60:]))
						t216 := v8
						v9 = t215
						t217 := int32(load32(m.memory[int64(uint32(v8))+56:]))
						t218 := v9
						v5 = t217
						p219 := v5
						if uint32(v9) < uint32(v5) {
							p219 = t218
						}
						v5 = p219
						store32(m.memory[int64(uint32(t216))+56:], uint32(v5))
						goto l146
					}
				l145:
					t220 := int32(m.memory[uint32(v5+v29+i32(-1))])
					v3 = t220
					t221 := int32(load32(m.memory[int64(uint32(v8))+60:]))
					t222 := v8
					v9 = t221
					t223 := int32(load32(m.memory[int64(uint32(v8))+56:]))
					t224 := v9
					v5 = t223 + v29
					p225 := v5
					if uint32(v9) < uint32(v5) {
						p225 = t224
					}
					v5 = p225
					store32(m.memory[int64(uint32(t222))+56:], uint32(v5))
					if v3 == 0 {
						goto l136
					}
				}
			l146:
				if v9 != v5 {
					goto l136
				}
				v3 = i32(0)
				goto l30
			}
			v3 = i32(0)
			goto l30
		}
	l136:
		store32(m.memory[int64(uint32(v8))+88:], uint32(i32(0)))
		v9 = i32(7)
		goto l18
	case 7:
		t166 := int32(load32(m.memory[int64(uint32(v8))+120:]))
		v27 = t166
		if v27&i32(4096) == 0 {
			v9 = i32(8)
			t168 := int32(load32(m.memory[int64(uint32(v8))+140:]))
			v5 = t168
			if v5 == 0 {
				goto l18
			}
			store32(m.memory[int64(uint32(v5))+36:], uint32(i32(0)))
			goto l18
		}
		v20 = i32(7)
		t167 := int32(load32(m.memory[int64(uint32(v8))+112:]))
		if t167 != 0 {
			{
				t169 := int32(load32(m.memory[int64(uint32(v8))+60:]))
				v3 = t169
				t170 := int32(load32(m.memory[int64(uint32(v8))+56:]))
				t171 := v3
				v5 = t170
				if t171 != v5 {
					goto l124
				}
				v29 = i32(0)
				goto l125
			}
		l124:
			v29 = v3 - v5
			v9 = i32(0)
		l127:
			{
				t172 := int32(m.memory[uint32(v5+v9)])
				if t172 == 0 {
					goto l126
				}
				t173 := v5
				v9 = v9 + i32(1)
				if t173+v9 != v3 {
					goto l127
				}
				goto l125
			}
		l126:
			v29 = v9 + i32(1)
		l125:
			{
				t174 := int32(load32(m.memory[int64(uint32(v8))+140:]))
				v9 = t174
				if v9 == 0 {
					goto l128
				}
				t175 := int32(load32(m.memory[int64(uint32(v9))+36:]))
				v28 = t175
				if v28 == 0 {
					goto l128
				}
				{
					t176 := int32(load32(m.memory[int64(uint32(v9))+40:]))
					v3 = t176
					t177 := int32(load32(m.memory[int64(uint32(v8))+88:]))
					t178 := v3
					v9 = t177
					if uint32(t178) < uint32(v9) {
						m.fn146(i32(1284056), i32(18), i32(1284076))
						panic("unreachable")
					}
					{
						v3 = v3 - v9
						p179 := v29
						if uint32(v3) < uint32(v29) {
							p179 = v3
						}
						v3 = p179
						if v3 == 0 {
							goto l130
						}
						memory_copy(m.memory, uint32(v28+v9), uint32(v5), uint32(v3))
					}
				l130:
					t180 := int32(load32(m.memory[int64(uint32(v8))+88:]))
					store32(m.memory[int64(uint32(v8))+88:], uint32(t180+v3))
					t181 := int32(load32(m.memory[int64(uint32(v8))+120:]))
					v27 = t181
					goto l128
				}
			}
		l128:
			{
				if v27&i32(512) == 0 {
					goto l131
				}
				t182 := int32(m.memory[int64(uint32(v8))+3])
				if t182&i32(4) == 0 {
					goto l131
				}
				t183 := int32(load32(m.memory[int64(uint32(v8))+124:]))
				t184 := m.fn910(t183, v5, v29)
				store32(m.memory[int64(uint32(v8))+124:], uint32(t184))
			}
		l131:
			{
				{
					if v29 != 0 {
						goto l132
					}
					t185 := int32(load32(m.memory[int64(uint32(v8))+60:]))
					t186 := v8
					v5 = t185
					t187 := int32(load32(m.memory[int64(uint32(v8))+56:]))
					t188 := v5
					v9 = t187
					p189 := v9
					if uint32(v5) < uint32(v9) {
						p189 = t188
					}
					v3 = p189
					store32(m.memory[int64(uint32(t186))+56:], uint32(v3))
					goto l133
				}
			l132:
				t190 := int32(m.memory[uint32(v5+v29+i32(-1))])
				v27 = t190
				t191 := int32(load32(m.memory[int64(uint32(v8))+60:]))
				t192 := v8
				v5 = t191
				t193 := int32(load32(m.memory[int64(uint32(v8))+56:]))
				t194 := v5
				v9 = t193 + v29
				p195 := v9
				if uint32(v5) < uint32(v9) {
					p195 = t194
				}
				v3 = p195
				store32(m.memory[int64(uint32(t192))+56:], uint32(v3))
				v9 = i32(8)
				if v27 == 0 {
					goto l18
				}
			}
		l133:
			v9 = i32(8)
			if v5 != v3 {
				goto l18
			}
			v3 = i32(0)
			goto l30
		}
		v3 = i32(0)
		goto l30
	case 8:
		{
			t148 := int32(load32(m.memory[int64(uint32(v8))+120:]))
			v5 = t148
			if v5&i32(512) == 0 {
				goto l117
			}
			{
				t149 := int32(m.memory[int64(uint32(v8))+64])
				v9 = t149
				if uint32(v9) > uint32(i32(15)) {
					goto l118
				}
				v3 = i32(0)
				v20 = i32(8)
				t150 := int32(load32(m.memory[int64(uint32(v8))+56:]))
				v29 = t150
				t151 := int32(load32(m.memory[int64(uint32(v8))+60:]))
				t152 := v29
				v27 = t151
				if t152 == v27 {
					goto l30
				}
				t153 := v8
				v28 = v29 + i32(1)
				store32(m.memory[int64(uint32(t153))+56:], uint32(v28))
				t154 := v8
				v25 = int64(uint32(v9))
				v30 = v25 + i64(8)
				m.memory[int64(uint32(t154))+64] = byte(v30)
				t155 := int64(m.memory[uint32(v29)])
				t156 := int64(load64(m.memory[int64(uint32(v8))+48:]))
				t157 := v8
				v25 = i64_shl(t155, v25) | t156
				store64(m.memory[int64(uint32(t157))+48:], uint64(v25))
				if uint32(v9) > uint32(i32(7)) {
					goto l118
				}
				if v28 == v27 {
					goto l30
				}
				store32(m.memory[int64(uint32(v8))+56:], uint32(v29+i32(2)))
				m.memory[int64(uint32(v8))+64] = byte(v9 | i32(16))
				t158 := int64(m.memory[int64(uint32(v29))+1])
				store64(m.memory[int64(uint32(v8))+48:], uint64(i64_shl(t158, v30)|v25))
			}
		l118:
			{
				t159 := int32(m.memory[int64(uint32(v8))+3])
				if t159&i32(4) == 0 {
					goto l119
				}
				t160 := int32(load16(m.memory[int64(uint32(v8))+124:]))
				t161 := int32(load32(m.memory[int64(uint32(v8))+48:]))
				if t160 != t161 {
					store32(m.memory[int64(uint32(v8))+136:], uint32(i32(20)))
					store32(m.memory[int64(uint32(v8))+132:], uint32(i32(1064797)))
					goto l57
				}
			}
		l119:
			m.memory[int64(uint32(v8))+64] = byte(i32(0))
			store64(m.memory[int64(uint32(v8))+48:], uint64(i64(0)))
		}
	l117:
		{
			t162 := int32(load32(m.memory[int64(uint32(v8))+140:]))
			v9 = t162
			if v9 == 0 {
				goto l121
			}
			store32(m.memory[int64(uint32(v9))+44:], uint32(int32(uint32(v5)>>9)&i32(1)))
			t163 := int32(load32(m.memory[int64(uint32(v8))+140:]))
			store32(m.memory[int64(uint32(t163))+48:], uint32(i32(1)))
			t164 := int32(load32(m.memory[int64(uint32(v8))+120:]))
			v5 = t164
		}
	l121:
		v9 = i32(12)
		t165 := int32(m.memory[int64(uint32(v8))+3])
		if t165&i32(4) == 0 {
			goto l18
		}
		if v5 == 0 {
			goto l18
		}
		store64(m.memory[int64(uint32(v8))+124:], uint64(i64(0)))
		goto l18
	case 11:
		{
			t563 := int32(m.memory[int64(uint32(v8))+3])
			v9 = t563
			if v9 == 0 {
				goto l254
			}
			t564 := int32(load32(m.memory[int64(uint32(v8))+120:]))
			if t564 == 0 {
				goto l254
			}
			{
				t565 := int32(m.memory[int64(uint32(v8))+64])
				v5 = t565
				if uint32(v5) > uint32(i32(31)) {
					goto l255
				}
				v3 = i32(0)
				v20 = i32(11)
				t566 := int32(load32(m.memory[int64(uint32(v8))+56:]))
				v15 = t566
				t567 := int32(load32(m.memory[int64(uint32(v8))+60:]))
				t568 := v15
				v29 = t567
				if t568 == v29 {
					goto l30
				}
				t569 := v8
				v27 = v15 + i32(1)
				store32(m.memory[int64(uint32(t569))+56:], uint32(v27))
				t570 := v8
				v25 = int64(uint32(v5))
				v30 = v25 + i64(8)
				m.memory[int64(uint32(t570))+64] = byte(v30)
				t571 := int64(m.memory[uint32(v15)])
				t572 := int64(load64(m.memory[int64(uint32(v8))+48:]))
				t573 := v8
				v26 = i64_shl(t571, v25) | t572
				store64(m.memory[int64(uint32(t573))+48:], uint64(v26))
				if uint32(v5) > uint32(i32(23)) {
					goto l255
				}
				if v27 == v29 {
					goto l30
				}
				t574 := v8
				v27 = v15 + i32(2)
				store32(m.memory[int64(uint32(t574))+56:], uint32(v27))
				t575 := v8
				v31 = v25 + i64(16)
				m.memory[int64(uint32(t575))+64] = byte(v31)
				t576 := int64(m.memory[int64(uint32(v15))+1])
				t577 := v8
				v30 = i64_shl(t576, v30) | v26
				store64(m.memory[int64(uint32(t577))+48:], uint64(v30))
				if uint32(v5) > uint32(i32(15)) {
					goto l255
				}
				if v27 == v29 {
					goto l30
				}
				t578 := v8
				v27 = v15 + i32(3)
				store32(m.memory[int64(uint32(t578))+56:], uint32(v27))
				t579 := v8
				v25 = v25 + i64(24)
				m.memory[int64(uint32(t579))+64] = byte(v25)
				t580 := int64(m.memory[int64(uint32(v15))+2])
				t581 := v8
				v30 = i64_shl(t580, v31) | v30
				store64(m.memory[int64(uint32(t581))+48:], uint64(v30))
				if uint32(v5) > uint32(i32(7)) {
					goto l255
				}
				if v27 == v29 {
					goto l30
				}
				store32(m.memory[int64(uint32(v8))+56:], uint32(v15+i32(4)))
				m.memory[int64(uint32(v8))+64] = byte(v5 | i32(32))
				t582 := int64(m.memory[int64(uint32(v15))+3])
				store64(m.memory[int64(uint32(v8))+48:], uint64(i64_shl(t582, v25)|v30))
			}
		l255:
			{
				if v9&i32(4) == 0 {
					goto l256
				}
				t583 := int32(load32(m.memory[int64(uint32(v8))+84:]))
				t584 := int32(load32(m.memory[int64(uint32(v8))+48:]))
				if t583 != t584 {
					store32(m.memory[int64(uint32(v8))+136:], uint32(i32(23)))
					store32(m.memory[int64(uint32(v8))+132:], uint32(i32(1064723)))
					goto l57
				}
			}
		l256:
			m.memory[int64(uint32(v8))+64] = byte(i32(0))
			store64(m.memory[int64(uint32(v8))+48:], uint64(i64(0)))
			goto l254
		}
	l254:
		v20 = i32(29)
		v3 = i32(1)
		goto l30
	case 12:
		v9 = i32(13)
		t562 := int32(m.memory[int64(uint32(v8))+4])
		if uint32(t562) <= uint32(i32(4)) {
			goto l18
		}
		v3 = i32(0)
		v20 = i32(12)
		goto l30
	case 13:
		t137 := int32(m.memory[int64(uint32(v8))+1])
		v5 = t137
		if v5&i32(1) == 0 {
			t141 := int64(load64(m.memory[int64(uint32(v8))+48:]))
			v25 = t141
			{
				t142 := int32(m.memory[int64(uint32(v8))+64])
				v9 = t142
				if uint32(v9) > uint32(i32(2)) {
					goto l111
				}
				{
					t143 := int32(load32(m.memory[int64(uint32(v8))+56:]))
					v3 = t143
					t144 := int32(load32(m.memory[int64(uint32(v8))+60:]))
					if v3 != t144 {
						goto l112
					}
					v3 = i32(0)
					v20 = i32(13)
					goto l30
				}
			l112:
				v30 = int64(uint32(v9))
				store32(m.memory[int64(uint32(v8))+56:], uint32(v3+i32(1)))
				v9 = v9 | i32(8)
				t145 := int64(m.memory[uint32(v3)])
				v25 = i64_shl(t145, v30) | v25
			}
		l111:
			m.memory[int64(uint32(v8))+64] = byte(v9 + i32(-1))
			t146 := v8
			v30 = int64(uint64(v25) >> 1)
			store64(m.memory[int64(uint32(t146))+48:], uint64(v30))
			m.memory[int64(uint32(v8))+1] = byte(int32(v25)&i32(1) | v5)
			switch int32(v30) & i32(3) {
			default:
				m.memory[int64(uint32(v8))+64] = byte(v9 + i32(-3))
				store64(m.memory[int64(uint32(v8))+48:], uint64(int64(uint64(v25)>>3)))
				v9 = i32(14)
				goto l18
			case 1:
				v3 = i32(0)
				m.memory[int64(uint32(v8))+160] = byte(i32(0))
				store32(m.memory[int64(uint32(v8))+156:], uint32(i32(5)))
				m.memory[int64(uint32(v8))+152] = byte(i32(0))
				store32(m.memory[int64(uint32(v8))+148:], uint32(i32(9)))
				m.memory[int64(uint32(v8))+64] = byte(v9 + i32(-3))
				store64(m.memory[int64(uint32(v8))+48:], uint64(int64(uint64(v25)>>3)))
				v9 = i32(17)
				v20 = i32(17)
				t147 := int32(m.memory[int64(uint32(v8))+4])
				if t147 != i32(6) {
					goto l18
				}
				goto l30
			case 2:
				m.memory[int64(uint32(v8))+64] = byte(v9 + i32(-3))
				store64(m.memory[int64(uint32(v8))+48:], uint64(int64(uint64(v25)>>3)))
				v9 = i32(24)
				goto l18
			case 3:
				store32(m.memory[int64(uint32(v8))+136:], uint32(i32(19)))
				store32(m.memory[int64(uint32(v8))+132:], uint32(i32(1065062)))
				v3 = i32(-3)
				m.memory[int64(uint32(v8))+64] = byte(v9 + i32(-3))
				store64(m.memory[int64(uint32(v8))+48:], uint64(int64(uint64(v25)>>3)))
				v20 = i32(30)
				goto l30
			}
		}
		t138 := int32(m.memory[int64(uint32(v8))+64])
		t139 := v8
		v9 = t138
		m.memory[int64(uint32(t139))+64] = byte(v9 & i32(248))
		t140 := int64(load64(m.memory[int64(uint32(v8))+48:]))
		store64(m.memory[int64(uint32(v8))+48:], uint64(i64_shr_u(t140, int64(uint32(v9&i32(7))))))
		v9 = i32(16)
		goto l18
	case 14:
		t540 := int32(m.memory[int64(uint32(v8))+64])
		t541 := v8
		v9 = t540
		m.memory[int64(uint32(t541))+64] = byte(v9 & i32(248))
		t542 := int64(load64(m.memory[int64(uint32(v8))+48:]))
		t543 := v8
		v25 = i64_shr_u(t542, int64(uint32(v9&i32(7))))
		store64(m.memory[int64(uint32(t543))+48:], uint64(v25))
		{
			if uint32(v9) > uint32(i32(31)) {
				goto l252
			}
			v3 = i32(0)
			v20 = i32(14)
			t544 := int32(load32(m.memory[int64(uint32(v8))+56:]))
			v5 = t544
			t545 := int32(load32(m.memory[int64(uint32(v8))+60:]))
			t546 := v5
			v29 = t545
			if t546 == v29 {
				goto l30
			}
			t547 := v8
			v27 = v5 + i32(1)
			store32(m.memory[int64(uint32(t547))+56:], uint32(v27))
			t548 := v8
			v30 = int64(uint32(v9)) & i64(24)
			v26 = v30 + i64(8)
			m.memory[int64(uint32(t548))+64] = byte(v26)
			t549 := int64(m.memory[uint32(v5)])
			t550 := v8
			v25 = i64_shl(t549, v30) | v25
			store64(m.memory[int64(uint32(t550))+48:], uint64(v25))
			v9 = v9 & i32(24)
			if v9 == i32(24) {
				goto l252
			}
			if v27 == v29 {
				goto l30
			}
			t551 := v8
			v27 = v5 + i32(2)
			store32(m.memory[int64(uint32(t551))+56:], uint32(v27))
			t552 := v8
			v31 = v30 + i64(16)
			m.memory[int64(uint32(t552))+64] = byte(v31)
			t553 := int64(m.memory[int64(uint32(v5))+1])
			t554 := v8
			v25 = i64_shl(t553, v26) | v25
			store64(m.memory[int64(uint32(t554))+48:], uint64(v25))
			if uint32(v9) > uint32(i32(15)) {
				goto l252
			}
			if v27 == v29 {
				goto l30
			}
			t555 := v8
			v27 = v5 + i32(3)
			store32(m.memory[int64(uint32(t555))+56:], uint32(v27))
			t556 := v8
			v30 = v30 + i64(24)
			m.memory[int64(uint32(t556))+64] = byte(v30)
			t557 := int64(m.memory[int64(uint32(v5))+2])
			t558 := v8
			v25 = i64_shl(t557, v31) | v25
			store64(m.memory[int64(uint32(t558))+48:], uint64(v25))
			if v9 != 0 {
				goto l252
			}
			if v27 == v29 {
				goto l30
			}
			m.memory[int64(uint32(v8))+64] = byte(i32(32))
			store32(m.memory[int64(uint32(v8))+56:], uint32(v5+i32(4)))
			t559 := int64(m.memory[int64(uint32(v5))+3])
			t560 := v8
			v25 = i64_shl(t559, v30) | v25
			store64(m.memory[int64(uint32(t560))+48:], uint64(v25))
		}
	l252:
		{
			if (int64(uint64(v25)>>16)^v25)&i64(0xffff) != i64(0xffff) {
				store32(m.memory[int64(uint32(v8))+136:], uint32(i32(29)))
				store32(m.memory[int64(uint32(v8))+132:], uint32(i32(1064016)))
				goto l57
			}
			v3 = i32(0)
			m.memory[int64(uint32(v8))+64] = byte(i32(0))
			store64(m.memory[int64(uint32(v8))+48:], uint64(i64(0)))
			store32(m.memory[int64(uint32(v8))+88:], uint32(int32(v25)&i32(0xffff)))
			v9 = i32(15)
			t561 := int32(m.memory[int64(uint32(v8))+4])
			if t561 != i32(6) {
				goto l18
			}
			v20 = i32(14)
			goto l30
		}
	case 15:
		v9 = i32(12)
		t647 := int32(load32(m.memory[int64(uint32(v8))+88:]))
		v20 = t647
		if v20 == 0 {
			goto l18
		}
		t648 := int32(load32(m.memory[int64(uint32(v8))+56:]))
		v27 = t648
		t649 := int32(load32(m.memory[int64(uint32(v8))+60:]))
		v29 = t649
		t650 := int32(load32(m.memory[int64(uint32(v8))+80:]))
		v3 = t650
		t651 := int32(load32(m.memory[int64(uint32(v8))+72:]))
		v21 = t651
		t652 := int32(load32(m.memory[int64(uint32(v8))+76:]))
		v28 = t652
	l274:
		{
			v5 = v29 - v27
			t653 := int32(load32(m.memory[int64(uint32(v8))+76:]))
			t654 := v5
			v29 = t653 - v3
			p655 := v20
			if uint32(v29) < uint32(v20) {
				p655 = v29
			}
			v20 = p655
			p656 := v20
			if uint32(v5) < uint32(v20) {
				p656 = t654
			}
			v5 = p656
			if v5 != 0 {
				if uint32(v28) < uint32(v3) {
					m.fn127(v3, v28, v28, i32(1283800))
					panic("unreachable")
				}
				t657 := v5
				v20 = v28 - v3
				if uint32(t657) > uint32(v20) {
					m.fn127(i32(0), v5, v20, i32(1283784))
					panic("unreachable")
				}
				if v5 == 0 {
					goto l273
				}
				memory_copy(m.memory, uint32(v21+v3), uint32(v27), uint32(v5))
			l273:
				t658 := v8
				v3 = v5 + v3
				store32(m.memory[int64(uint32(t658))+80:], uint32(v3))
				t659 := int32(load32(m.memory[int64(uint32(v8))+88:]))
				t660 := v8
				v20 = t659 - v5
				store32(m.memory[int64(uint32(t660))+88:], uint32(v20))
				t661 := int32(load32(m.memory[int64(uint32(v8))+60:]))
				t662 := v8
				v29 = t661
				t663 := int32(load32(m.memory[int64(uint32(v8))+56:]))
				t664 := v29
				v5 = t663 + v5
				p665 := v5
				if uint32(v29) < uint32(v5) {
					p665 = t664
				}
				v27 = p665
				store32(m.memory[int64(uint32(t662))+56:], uint32(v27))
				if v20 == 0 {
					goto l18
				}
				goto l274
			}
			v3 = i32(0)
			v20 = i32(15)
			goto l30
		}
	case 16:
		v9 = i32(11)
		t100 := int32(m.memory[int64(uint32(v8))+3])
		v5 = t100
		if v5 == 0 {
			goto l18
		}
		{
			t101 := int32(m.memory[int64(uint32(v8))+64])
			v29 = t101
			if uint32(v29) > uint32(i32(31)) {
				goto l102
			}
			v3 = i32(0)
			v20 = i32(16)
			t102 := int32(load32(m.memory[int64(uint32(v8))+56:]))
			v27 = t102
			t103 := int32(load32(m.memory[int64(uint32(v8))+60:]))
			t104 := v27
			v28 = t103
			if t104 == v28 {
				goto l30
			}
			t105 := v8
			v21 = v27 + i32(1)
			store32(m.memory[int64(uint32(t105))+56:], uint32(v21))
			t106 := v8
			v25 = int64(uint32(v29))
			v30 = v25 + i64(8)
			m.memory[int64(uint32(t106))+64] = byte(v30)
			t107 := int64(m.memory[uint32(v27)])
			t108 := int64(load64(m.memory[int64(uint32(v8))+48:]))
			t109 := v8
			v26 = i64_shl(t107, v25) | t108
			store64(m.memory[int64(uint32(t109))+48:], uint64(v26))
			if uint32(v29) > uint32(i32(23)) {
				goto l102
			}
			if v21 == v28 {
				goto l30
			}
			t110 := v8
			v21 = v27 + i32(2)
			store32(m.memory[int64(uint32(t110))+56:], uint32(v21))
			t111 := v8
			v31 = v25 + i64(16)
			m.memory[int64(uint32(t111))+64] = byte(v31)
			t112 := int64(m.memory[int64(uint32(v27))+1])
			t113 := v8
			v30 = i64_shl(t112, v30) | v26
			store64(m.memory[int64(uint32(t113))+48:], uint64(v30))
			if uint32(v29) > uint32(i32(15)) {
				goto l102
			}
			if v21 == v28 {
				goto l30
			}
			t114 := v8
			v21 = v27 + i32(3)
			store32(m.memory[int64(uint32(t114))+56:], uint32(v21))
			t115 := v8
			v25 = v25 + i64(24)
			m.memory[int64(uint32(t115))+64] = byte(v25)
			t116 := int64(m.memory[int64(uint32(v27))+2])
			t117 := v8
			v30 = i64_shl(t116, v31) | v30
			store64(m.memory[int64(uint32(t117))+48:], uint64(v30))
			if uint32(v29) > uint32(i32(7)) {
				goto l102
			}
			if v21 == v28 {
				goto l30
			}
			store32(m.memory[int64(uint32(v8))+56:], uint32(v27+i32(4)))
			m.memory[int64(uint32(v8))+64] = byte(v29 | i32(32))
			t118 := int64(m.memory[int64(uint32(v27))+3])
			store64(m.memory[int64(uint32(v8))+48:], uint64(i64_shl(t118, v25)|v30))
		}
	l102:
		t119 := int32(load32(m.memory[int64(uint32(v8))+84:]))
		t120 := int32(load32(m.memory[int64(uint32(v8))+80:]))
		t121 := v8
		v20 = t120
		store32(m.memory[int64(uint32(t121))+84:], uint32(t119+v20))
		t122 := int32(load32(m.memory[int64(uint32(v8))+120:]))
		v3 = t122
		{
			{
				if v5&i32(4) == 0 {
					goto l103
				}
				{
					if v3 != 0 {
						goto l104
					}
					t123 := int32(load32(m.memory[int64(uint32(v8))+124:]))
					t124 := int32(load32(m.memory[int64(uint32(v8))+72:]))
					t125 := m.fn909(t123, t124, v20)
					store32(m.memory[int64(uint32(v8))+124:], uint32(t125))
					t126 := int64(load64(m.memory[int64(uint32(v8))+48:]))
					v25 = t126
					goto l105
				}
			l104:
				t127 := int32(load32(m.memory[int64(uint32(v8))+128:]))
				t128 := int32(load32(m.memory[int64(uint32(v8))+72:]))
				t129 := m.fn910(t127, t128, v20)
				t130 := v8
				v5 = t129
				store32(m.memory[int64(uint32(t130))+124:], uint32(v5))
				store32(m.memory[int64(uint32(v8))+128:], uint32(v5))
				t131 := int32(load32(m.memory[int64(uint32(v8))+120:]))
				v3 = t131
			}
		l103:
			t132 := int64(load64(m.memory[int64(uint32(v8))+48:]))
			v25 = t132
			if v3 != 0 {
				goto l106
			}
		}
	l105:
		v5 = int32(v25)
		v5 = i32_rotr(v5&i32(0xff00ff), i32(8)) | i32_rotr(v5, i32(24))&i32(0xff00ff)
		goto l107
	l106:
		v5 = int32(v25)
	l107:
		t133 := int32(load32(m.memory[int64(uint32(v8))+76:]))
		t134 := int32(load32(m.memory[int64(uint32(v8))+80:]))
		store32(m.memory[int64(uint32(v8))+116:], uint32(t133-t134))
		{
			t135 := int32(m.memory[int64(uint32(v8))+3])
			if t135&i32(4) == 0 {
				goto l108
			}
			t136 := int32(load32(m.memory[int64(uint32(v8))+124:]))
			if v5 != t136 {
				store32(m.memory[int64(uint32(v8))+136:], uint32(i32(21)))
				store32(m.memory[int64(uint32(v8))+132:], uint32(i32(1064746)))
				goto l57
			}
		}
	l108:
		m.memory[int64(uint32(v8))+64] = byte(i32(0))
		store64(m.memory[int64(uint32(v8))+48:], uint64(i64(0)))
		goto l18
	case 17:
		goto l18
	case 18:
		m.memory[uint32(v8)] = byte(i32(18))
		{
			t387 := int32(load32(m.memory[int64(uint32(v8))+60:]))
			v3 = t387
			t388 := int32(load32(m.memory[int64(uint32(v8))+56:]))
			t389 := v3
			v24 = t388
			if uint32(t389-v24) < uint32(i32(15)) {
				goto l189
			}
			t390 := int32(load32(m.memory[int64(uint32(v8))+76:]))
			t391 := int32(load32(m.memory[int64(uint32(v8))+80:]))
			if uint32(t390-t391) <= uint32(i32(259)) {
				goto l189
			}
			m.fn913(v8)
			t392 := int32(m.memory[uint32(v8)])
			v9 = t392
			if v9 != i32(18) {
				goto l18
			}
			t393 := int32(load32(m.memory[int64(uint32(v8))+60:]))
			v3 = t393
			t394 := int32(load32(m.memory[int64(uint32(v8))+56:]))
			v24 = t394
		}
	l189:
		t395 := int32(load32(m.memory[int64(uint32(v19))+8:]))
		v9 = t395
		store32(m.memory[int64(uint32(v8))+80:], uint32(i32(0)))
		t396 := int64(load64(m.memory[uint32(v19):]))
		v25 = t396
		store64(m.memory[int64(uint32(v8))+72:], uint64(i64(1)))
		store32(m.memory[int64(uint32(v7))+24:], uint32(v9))
		store64(m.memory[int64(uint32(v7))+16:], uint64(v25))
		t397 := int32(load32(m.memory[uint32(v12):]))
		store32(m.memory[int64(uint32(v7))+8:], uint32(t397))
		t398 := int32(load32(m.memory[int64(uint32(v12))+3:]))
		store32(m.memory[int64(uint32(v7))+11:], uint32(t398))
		v28 = i32(512)
		v21 = i32(1277088)
		{
			t399 := int32(m.memory[int64(uint32(v8))+152])
			switch t399 {
			default:
				goto l190
			case 1:
				v28 = i32(1332)
				v21 = v16
				goto l190
			case 2:
				v28 = i32(1332)
				v21 = v17
				goto l190
			case 3:
				v28 = i32(592)
				v21 = v18
			}
		}
	l190:
		v29 = i32(32)
		v27 = i32(1279136)
		{
			t400 := int32(m.memory[int64(uint32(v8))+160])
			switch t400 {
			default:
				goto l194
			case 1:
				v29 = i32(1332)
				v27 = v16
				goto l194
			case 2:
				v29 = i32(1332)
				v27 = v17
				goto l194
			case 3:
				v29 = i32(592)
				v27 = v18
			}
		}
	l194:
		t401 := int32(m.memory[int64(uint32(v8))+64])
		v22 = t401
		t402 := int64(load64(m.memory[int64(uint32(v8))+48:]))
		v25 = t402
		{
		l220:
			v20 = i32(18)
		l212:
			{
				{
					{
						{
							{
								switch v20&i32(255) + i32(-18) {
								case 1:
									t418 := int32(load32(m.memory[int64(uint32(v7))+24:]))
									v9 = t418
									t419 := int32(load32(m.memory[int64(uint32(v7))+20:]))
									t420 := v9
									v5 = t419
									if t420 == v5 {
										v20 = i32(19)
										m.memory[uint32(v8)] = byte(i32(19))
										t483 := int64(load64(m.memory[int64(uint32(v7))+16:]))
										store64(m.memory[uint32(v19):], uint64(t483))
										t484 := int32(load32(m.memory[int64(uint32(v7))+24:]))
										store32(m.memory[int64(uint32(v19))+8:], uint32(t484))
										m.memory[int64(uint32(v8))+64] = byte(v22)
										store32(m.memory[int64(uint32(v8))+60:], uint32(v3))
										store32(m.memory[int64(uint32(v8))+56:], uint32(v24))
										store64(m.memory[int64(uint32(v8))+48:], uint64(v25))
										t485 := int32(load32(m.memory[int64(uint32(v7))+8:]))
										store32(m.memory[uint32(v12):], uint32(t485))
										t486 := int32(load32(m.memory[int64(uint32(v7))+11:]))
										store32(m.memory[int64(uint32(v12))+3:], uint32(t486))
										v3 = i32(0)
										goto l30
									}
									{
										if uint32(v9) >= uint32(v5) {
											m.fn39(v9, v5, i32(1279296))
											panic("unreachable")
										}
										t421 := int32(load32(m.memory[int64(uint32(v7))+16:]))
										t422 := int32(load32(m.memory[int64(uint32(v8))+88:]))
										m.memory[uint32(t421+v9)] = byte(t422)
										store32(m.memory[int64(uint32(v7))+24:], uint32(v9+i32(1)))
										goto l220
									}
								case 3:
									v5 = v24
									v9 = v22
									t429 := int64(load32(m.memory[int64(uint32(v8))+156:]))
									t430 := v29
									t431 := v25
									v31 = i64_shl(i64(-1), t429) ^ i64(-1)
									v20 = int32(t431 & v31)
									if uint32(t430) <= uint32(v20) {
										goto l223
									}
								l226:
									{
										v23 = v27 + v20<<2
										t432 := int32(m.memory[int64(uint32(v23))+3])
										v20 = t432
										if uint32(v20) <= uint32(v9&i32(255)) {
											goto l224
										}
										{
											if v5 == v3 {
												goto l225
											}
											v30 = int64(uint32(v9))
											t433 := int64(m.memory[uint32(v5)])
											v26 = t433
											v5 = v5 + i32(1)
											v9 = v9 + i32(8)
											t434 := v29
											v25 = i64_shl(v26, v30) | v25
											v20 = int32(v25 & v31)
											if uint32(t434) <= uint32(v20) {
												goto l223
											}
											goto l226
										}
									l225:
									}
									v20 = i32(21)
									m.memory[uint32(v8)] = byte(i32(21))
									t435 := int64(load64(m.memory[int64(uint32(v7))+16:]))
									store64(m.memory[uint32(v19):], uint64(t435))
									t436 := int32(load32(m.memory[int64(uint32(v7))+24:]))
									store32(m.memory[int64(uint32(v19))+8:], uint32(t436))
									store32(m.memory[int64(uint32(v8))+60:], uint32(v3))
									store32(m.memory[int64(uint32(v8))+56:], uint32(v3))
									store64(m.memory[int64(uint32(v8))+48:], uint64(v25))
									t437 := int32(load32(m.memory[int64(uint32(v7))+8:]))
									store32(m.memory[uint32(v12):], uint32(t437))
									t438 := int32(load32(m.memory[int64(uint32(v7))+11:]))
									store32(m.memory[int64(uint32(v12))+3:], uint32(t438))
									m.memory[int64(uint32(v8))+64] = byte(v22 + v3<<3 - v24<<3)
									v3 = i32(0)
									goto l30
								case 5:
									t445 := int32(load32(m.memory[int64(uint32(v7))+20:]))
									v5 = t445
									t446 := int32(load32(m.memory[int64(uint32(v7))+24:]))
									t447 := v5
									v9 = t446
									if t447 == v9 {
										v20 = i32(23)
										m.memory[uint32(v8)] = byte(i32(23))
										t495 := int64(load64(m.memory[int64(uint32(v7))+16:]))
										store64(m.memory[uint32(v19):], uint64(t495))
										t496 := int32(load32(m.memory[int64(uint32(v7))+24:]))
										store32(m.memory[int64(uint32(v19))+8:], uint32(t496))
										m.memory[int64(uint32(v8))+64] = byte(v22)
										store32(m.memory[int64(uint32(v8))+60:], uint32(v3))
										store32(m.memory[int64(uint32(v8))+56:], uint32(v24))
										store64(m.memory[int64(uint32(v8))+48:], uint64(v25))
										t497 := int32(load32(m.memory[int64(uint32(v7))+8:]))
										store32(m.memory[uint32(v12):], uint32(t497))
										t498 := int32(load32(m.memory[int64(uint32(v7))+11:]))
										store32(m.memory[int64(uint32(v12))+3:], uint32(t498))
										v3 = i32(0)
										goto l30
									}
									v5 = v5 - v9
									{
										t448 := int32(load32(m.memory[int64(uint32(v8))+92:]))
										v20 = t448
										if uint32(v20) > uint32(v9) {
											v9 = v20 - v9
											t454 := int32(load32(m.memory[int64(uint32(v8))+16:]))
											if uint32(v9) > uint32(t454) {
												t473 := int32(m.memory[int64(uint32(v8))+1])
												if t473&i32(4) == 0 {
													m.fn34(i32(1279392), i32(85), i32(1283988))
													panic("unreachable")
												}
												t474 := int32(load32(m.memory[int64(uint32(v7))+24:]))
												store32(m.memory[int64(uint32(v19))+8:], uint32(t474))
												t475 := int64(load64(m.memory[int64(uint32(v7))+16:]))
												store64(m.memory[uint32(v19):], uint64(t475))
												m.memory[int64(uint32(v8))+64] = byte(v22)
												store32(m.memory[int64(uint32(v8))+60:], uint32(v3))
												store32(m.memory[int64(uint32(v8))+56:], uint32(v24))
												store64(m.memory[int64(uint32(v8))+48:], uint64(v25))
												t476 := int32(load32(m.memory[int64(uint32(v7))+8:]))
												store32(m.memory[uint32(v12):], uint32(t476))
												t477 := int32(load32(m.memory[int64(uint32(v7))+11:]))
												store32(m.memory[int64(uint32(v12))+3:], uint32(t477))
												v20 = i32(30)
												store32(m.memory[int64(uint32(v8))+136:], uint32(i32(30)))
												store32(m.memory[int64(uint32(v8))+132:], uint32(i32(1064767)))
												goto l235
											}
											t455 := int32(load32(m.memory[int64(uint32(v8))+12:]))
											v23 = t455
											v20 = v23 + i32(-64)
											if uint32(v20) >= uint32(i32(-63)) {
												m.fn2(i32(1283612), i32(74), i32(1283688))
												panic("unreachable")
											}
											t457 := v7 + i32(16)
											t458 := v13
											p456 := v20
											if uint32(v20) > uint32(v23) {
												p456 = i32(0)
											}
											t459 := int32(load32(m.memory[int64(uint32(v8))+20:]))
											t460 := v9
											v20 = t459
											v23 = t460 - v20
											t461 := p456 - v23
											t462 := v20 - v9
											var p463 int32
											if uint32(v9) > uint32(v20) {
												p463 = 1
											}
											v20 = p463
											p464 := t462
											if v20 != 0 {
												p464 = t461
											}
											v32 = p464
											t465 := int32(load32(m.memory[int64(uint32(v8))+88:]))
											t466 := v32
											t467 := v32
											t468 := v5
											v33 = t465
											t470 := v33
											p469 := v9
											if v20 != 0 {
												p469 = v23
											}
											v9 = p469
											p471 := v9
											if uint32(v33) < uint32(v9) {
												p471 = t470
											}
											v9 = p471
											p472 := v9
											if uint32(v5) < uint32(v9) {
												p472 = t468
											}
											v9 = p472
											m.fn912(t457, t458, t466, t467+v9)
											goto l231
										}
										t449 := int32(load32(m.memory[int64(uint32(v8))+88:]))
										t450 := v7 + i32(16)
										t451 := v20
										t452 := v5
										v9 = t449
										p453 := v9
										if uint32(v5) < uint32(v9) {
											p453 = t452
										}
										v9 = p453
										m.fn911(t450, t451, v9)
										goto l231
									}
								default:
									{
										if uint32(v3-v24) < uint32(i32(15)) {
											goto l204
										}
										t403 := int32(load32(m.memory[int64(uint32(v7))+20:]))
										t404 := int32(load32(m.memory[int64(uint32(v7))+24:]))
										if uint32(t403-t404) > uint32(i32(259)) {
											m.memory[uint32(v8)] = byte(i32(18))
											t478 := int64(load64(m.memory[int64(uint32(v7))+16:]))
											store64(m.memory[uint32(v19):], uint64(t478))
											t479 := int32(load32(m.memory[int64(uint32(v7))+24:]))
											store32(m.memory[int64(uint32(v19))+8:], uint32(t479))
											m.memory[int64(uint32(v8))+64] = byte(v22)
											store32(m.memory[int64(uint32(v8))+60:], uint32(v3))
											store32(m.memory[int64(uint32(v8))+56:], uint32(v24))
											store64(m.memory[int64(uint32(v8))+48:], uint64(v25))
											t480 := int32(load32(m.memory[int64(uint32(v7))+8:]))
											store32(m.memory[uint32(v12):], uint32(t480))
											t481 := int32(load32(m.memory[int64(uint32(v7))+11:]))
											store32(m.memory[int64(uint32(v12))+3:], uint32(t481))
											m.fn913(v8)
											t482 := int32(m.memory[uint32(v8)])
											v9 = t482
											goto l18
										}
									}
								l204:
									store32(m.memory[int64(uint32(v8))+100:], uint32(i32(0)))
									v5 = v24
									v9 = v22
									t405 := int64(load32(m.memory[int64(uint32(v8))+148:]))
									t406 := v28
									t407 := v25
									v31 = i64_shl(i64(-1), t405) ^ i64(-1)
									v20 = int32(t407 & v31)
									if uint32(t406) > uint32(v20) {
									l217:
										{
											v23 = v21 + v20<<2
											t411 := int32(m.memory[int64(uint32(v23))+3])
											v20 = t411
											if uint32(v20) <= uint32(v9&i32(255)) {
												goto l215
											}
											{
												if v5 == v3 {
													goto l216
												}
												v30 = int64(uint32(v9))
												t412 := int64(m.memory[uint32(v5)])
												v26 = t412
												v5 = v5 + i32(1)
												v9 = v9 + i32(8)
												t413 := v28
												v25 = i64_shl(v26, v30) | v25
												v20 = int32(v25 & v31)
												if uint32(t413) <= uint32(v20) {
													goto l207
												}
												goto l217
											}
										l216:
										}
										v20 = i32(18)
										m.memory[uint32(v8)] = byte(i32(18))
										t414 := int64(load64(m.memory[int64(uint32(v7))+16:]))
										store64(m.memory[uint32(v19):], uint64(t414))
										t415 := int32(load32(m.memory[int64(uint32(v7))+24:]))
										store32(m.memory[int64(uint32(v19))+8:], uint32(t415))
										store32(m.memory[int64(uint32(v8))+60:], uint32(v3))
										store32(m.memory[int64(uint32(v8))+56:], uint32(v3))
										store64(m.memory[int64(uint32(v8))+48:], uint64(v25))
										t416 := int32(load32(m.memory[int64(uint32(v7))+8:]))
										store32(m.memory[uint32(v12):], uint32(t416))
										t417 := int32(load32(m.memory[int64(uint32(v7))+11:]))
										store32(m.memory[int64(uint32(v12))+3:], uint32(t417))
										m.memory[int64(uint32(v8))+64] = byte(v22 + v3<<3 - v24<<3)
										v3 = i32(0)
										goto l30
									}
									goto l207
								case 2:
									t408 := int32(load32(m.memory[int64(uint32(v8))+96:]))
									v20 = t408
									if v20 != 0 {
										v9 = v24
										v5 = v22
										if uint32(v20) > uint32(v22&i32(255)) {
											goto l222
										}
										v5 = v22
										v23 = v24
										goto l211
									}
									t409 := int32(load32(m.memory[int64(uint32(v8))+88:]))
									v9 = t409
									goto l209
								case 4:
									v20 = i32(23)
									t410 := int32(load32(m.memory[int64(uint32(v8))+96:]))
									v23 = t410
									if v23 == 0 {
										goto l212
									}
									v9 = v24
									v5 = v22
									if uint32(v23) > uint32(v22&i32(255)) {
										goto l228
									}
									v5 = v22
									v32 = v24
									goto l214
								}
							l222:
								{
									if v9 == v3 {
										v20 = i32(20)
										m.memory[uint32(v8)] = byte(i32(20))
										t487 := int64(load64(m.memory[int64(uint32(v7))+16:]))
										store64(m.memory[uint32(v19):], uint64(t487))
										t488 := int32(load32(m.memory[int64(uint32(v7))+24:]))
										store32(m.memory[int64(uint32(v19))+8:], uint32(t488))
										store32(m.memory[int64(uint32(v8))+60:], uint32(v3))
										store32(m.memory[int64(uint32(v8))+56:], uint32(v3))
										store64(m.memory[int64(uint32(v8))+48:], uint64(v25))
										t489 := int32(load32(m.memory[int64(uint32(v7))+8:]))
										store32(m.memory[uint32(v12):], uint32(t489))
										t490 := int32(load32(m.memory[int64(uint32(v7))+11:]))
										store32(m.memory[int64(uint32(v12))+3:], uint32(t490))
										m.memory[int64(uint32(v8))+64] = byte(v22 + v3<<3 - v24<<3)
										v3 = i32(0)
										goto l30
									}
									t423 := int64(m.memory[uint32(v9)])
									v25 = i64_shl(t423, int64(uint32(v5))) | v25
									v23 = v9 + i32(1)
									v9 = v23
									t424 := v20
									v5 = v5 + i32(8)
									if uint32(t424) > uint32(v5&i32(255)) {
										goto l222
									}
								}
							l211:
								t425 := int32(load32(m.memory[int64(uint32(v8))+100:]))
								store32(m.memory[int64(uint32(v8))+100:], uint32(t425+v20))
								t426 := int32(load32(m.memory[int64(uint32(v8))+88:]))
								t427 := v8
								t428 := v25
								v30 = int64(uint32(v20))
								v9 = t426 + int32(t428&(i64_shl(i64(-1), v30)^i64(-1)))
								store32(m.memory[int64(uint32(t427))+88:], uint32(v9))
								v22 = v5 - v20
								v25 = i64_shr_u(v25, v30)
								v24 = v23
							}
						l209:
							store32(m.memory[int64(uint32(v8))+104:], uint32(v9))
							v20 = i32(21)
							goto l212
						l228:
							{
								if v9 == v3 {
									v20 = i32(22)
									m.memory[uint32(v8)] = byte(i32(22))
									t491 := int64(load64(m.memory[int64(uint32(v7))+16:]))
									store64(m.memory[uint32(v19):], uint64(t491))
									t492 := int32(load32(m.memory[int64(uint32(v7))+24:]))
									store32(m.memory[int64(uint32(v19))+8:], uint32(t492))
									store32(m.memory[int64(uint32(v8))+60:], uint32(v3))
									store32(m.memory[int64(uint32(v8))+56:], uint32(v3))
									store64(m.memory[int64(uint32(v8))+48:], uint64(v25))
									t493 := int32(load32(m.memory[int64(uint32(v7))+8:]))
									store32(m.memory[uint32(v12):], uint32(t493))
									t494 := int32(load32(m.memory[int64(uint32(v7))+11:]))
									store32(m.memory[int64(uint32(v12))+3:], uint32(t494))
									m.memory[int64(uint32(v8))+64] = byte(v22 + v3<<3 - v24<<3)
									v3 = i32(0)
									goto l30
								}
								t439 := int64(m.memory[uint32(v9)])
								v25 = i64_shl(t439, int64(uint32(v5))) | v25
								v32 = v9 + i32(1)
								v9 = v32
								t440 := v23
								v5 = v5 + i32(8)
								if uint32(t440) > uint32(v5&i32(255)) {
									goto l228
								}
							}
						l214:
							t441 := int32(load32(m.memory[int64(uint32(v8))+100:]))
							store32(m.memory[int64(uint32(v8))+100:], uint32(t441+v23))
							t442 := int32(load32(m.memory[int64(uint32(v8))+92:]))
							t443 := v8
							t444 := v25
							v30 = int64(uint32(v23))
							store32(m.memory[int64(uint32(t443))+92:], uint32(t442+int32(t444&(i64_shl(i64(-1), v30)^i64(-1)))))
							v22 = v5 - v23
							v25 = i64_shr_u(v25, v30)
							v24 = v32
							goto l212
						}
					l231:
						t499 := int32(load32(m.memory[int64(uint32(v8))+88:]))
						t500 := v8
						v5 = t499
						store32(m.memory[int64(uint32(t500))+88:], uint32(v5-v9))
						p501 := i32(23)
						if v5 == v9 {
							p501 = i32(18)
						}
						v20 = p501
						goto l212
					}
				l224:
					t502 := int32(load16(m.memory[uint32(v23):]))
					v32 = t502
					{
						{
							t503 := int32(m.memory[int64(uint32(v23))+2])
							v34 = t503
							if uint32(v34) <= uint32(i32(15)) {
								goto l236
							}
							v23 = v20
							goto l237
						}
					l236:
						t504 := v25
						v31 = i64_shl(i64(-1), int64(uint32(v34+v20))) ^ i64(-1)
						v23 = i32_shr_u(int32(t504&v31), v20) + v32
						if uint32(v23) >= uint32(v29) {
							goto l238
						}
						v34 = v20 & i32(31)
						{
						l241:
							{
								v33 = v27 + v23<<2
								t505 := int32(m.memory[int64(uint32(v33))+3])
								v23 = t505
								if uint32((v23+v20)&i32(255)) <= uint32(v9&i32(255)) {
									goto l239
								}
								{
									if v5 == v3 {
										goto l240
									}
									v30 = int64(uint32(v9))
									t506 := int64(m.memory[uint32(v5)])
									v26 = t506
									v5 = v5 + i32(1)
									v9 = v9 + i32(8)
									v25 = i64_shl(v26, v30) | v25
									v23 = i32_shr_u(int32(v25&v31), v34) + v32
									if uint32(v23) >= uint32(v29) {
										goto l238
									}
									goto l241
								}
							l240:
							}
							v20 = i32(21)
							m.memory[uint32(v8)] = byte(i32(21))
							t507 := int64(load64(m.memory[int64(uint32(v7))+16:]))
							store64(m.memory[uint32(v19):], uint64(t507))
							t508 := int32(load32(m.memory[int64(uint32(v7))+24:]))
							store32(m.memory[int64(uint32(v19))+8:], uint32(t508))
							store32(m.memory[int64(uint32(v8))+60:], uint32(v3))
							store32(m.memory[int64(uint32(v8))+56:], uint32(v3))
							store64(m.memory[int64(uint32(v8))+48:], uint64(v25))
							t509 := int32(load32(m.memory[int64(uint32(v7))+8:]))
							store32(m.memory[uint32(v12):], uint32(t509))
							t510 := int32(load32(m.memory[int64(uint32(v7))+11:]))
							store32(m.memory[int64(uint32(v12))+3:], uint32(t510))
							m.memory[int64(uint32(v8))+64] = byte(v22 + v3<<3 - v24<<3)
							v3 = i32(0)
							goto l30
						}
					l239:
						t511 := int32(m.memory[int64(uint32(v33))+2])
						v34 = t511
						t512 := int32(load16(m.memory[uint32(v33):]))
						v32 = t512
						t513 := int32(load32(m.memory[int64(uint32(v8))+100:]))
						store32(m.memory[int64(uint32(v8))+100:], uint32(t513+v20))
						v9 = v9 - v20
						v25 = i64_shr_u(v25, int64(uint32(v20)))
					}
				l237:
					v22 = v9 - v23
					v25 = i64_shr_u(v25, int64(uint32(v23)))
					if v34&i32(64) != 0 {
						goto l242
					}
					store32(m.memory[int64(uint32(v8))+96:], uint32(v34&i32(15)))
					store32(m.memory[int64(uint32(v8))+92:], uint32(v32&i32(0xffff)))
					v20 = i32(22)
					v24 = v5
					goto l212
				l242:
					t514 := int32(load32(m.memory[int64(uint32(v7))+24:]))
					store32(m.memory[int64(uint32(v19))+8:], uint32(t514))
					t515 := int64(load64(m.memory[int64(uint32(v7))+16:]))
					store64(m.memory[uint32(v19):], uint64(t515))
					m.memory[int64(uint32(v8))+64] = byte(v22)
					store32(m.memory[int64(uint32(v8))+60:], uint32(v3))
					store32(m.memory[int64(uint32(v8))+56:], uint32(v5))
					store64(m.memory[int64(uint32(v8))+48:], uint64(v25))
					t516 := int32(load32(m.memory[int64(uint32(v7))+8:]))
					store32(m.memory[uint32(v12):], uint32(t516))
					t517 := int32(load32(m.memory[int64(uint32(v7))+11:]))
					store32(m.memory[int64(uint32(v12))+3:], uint32(t517))
					store32(m.memory[int64(uint32(v8))+136:], uint32(i32(22)))
					store32(m.memory[int64(uint32(v8))+132:], uint32(i32(1065209)))
					v20 = i32(30)
				}
			l235:
				m.memory[uint32(v8)] = byte(v20)
				v3 = i32(-3)
				goto l30
			l238:
				m.fn39(v23, v29, i32(1283972))
				panic("unreachable")
			l223:
				m.fn39(v20, v29, i32(1283956))
				panic("unreachable")
			l215:
				t518 := int32(load16(m.memory[uint32(v23):]))
				v32 = t518
				{
					t519 := int32(m.memory[int64(uint32(v23))+2])
					v34 = t519
					if v34 != 0 {
						v23 = i32(0)
						if uint32(v34) <= uint32(i32(15)) {
							{
								t520 := v28
								t521 := v25
								v31 = i64_shl(i64(-1), int64(uint32(v34+v20))) ^ i64(-1)
								t522 := int32(t521&v31) & i32(0xffff)
								v34 = v20 & i32(15)
								v23 = (i32_shr_u(t522, v34) + v32) & i32(0xffff)
								if uint32(t520) <= uint32(v23) {
									goto l246
								}
								{
								l249:
									{
										v23 = v21 + v23<<2
										t523 := int32(m.memory[int64(uint32(v23))+3])
										v33 = t523
										if uint32((v33+v20)&i32(255)) <= uint32(v9&i32(255)) {
											v9 = v9 - v20
											v25 = i64_shr_u(v25, int64(uint32(v20)))
											t530 := int32(m.memory[int64(uint32(v23))+2])
											v34 = t530
											t531 := int32(load16(m.memory[uint32(v23):]))
											v32 = t531
											v23 = v20
											v24 = v5
											v20 = v33
											goto l244
										}
										{
											if v5 == v3 {
												goto l248
											}
											v30 = int64(uint32(v9))
											t524 := int64(m.memory[uint32(v5)])
											v26 = t524
											v5 = v5 + i32(1)
											v9 = v9 + i32(8)
											t525 := v28
											v25 = i64_shl(v26, v30) | v25
											v23 = (i32_shr_u(int32(v25&v31)&i32(0xffff), v34) + v32) & i32(0xffff)
											if uint32(t525) <= uint32(v23) {
												goto l246
											}
											goto l249
										}
									l248:
									}
									v20 = i32(18)
									m.memory[uint32(v8)] = byte(i32(18))
									t526 := int64(load64(m.memory[int64(uint32(v7))+16:]))
									store64(m.memory[uint32(v19):], uint64(t526))
									t527 := int32(load32(m.memory[int64(uint32(v7))+24:]))
									store32(m.memory[int64(uint32(v19))+8:], uint32(t527))
									store32(m.memory[int64(uint32(v8))+60:], uint32(v3))
									store32(m.memory[int64(uint32(v8))+56:], uint32(v3))
									store64(m.memory[int64(uint32(v8))+48:], uint64(v25))
									t528 := int32(load32(m.memory[int64(uint32(v7))+8:]))
									store32(m.memory[uint32(v12):], uint32(t528))
									t529 := int32(load32(m.memory[int64(uint32(v7))+11:]))
									store32(m.memory[int64(uint32(v12))+3:], uint32(t529))
									m.memory[int64(uint32(v8))+64] = byte(v22 + v3<<3 - v24<<3)
									v3 = i32(0)
									goto l30
								}
							}
						l246:
							m.fn39(v23, v28, i32(1283940))
							panic("unreachable")
						}
						v24 = v5
						goto l244
					}
					v23 = i32(0)
					v24 = v5
					v34 = i32(0)
					goto l244
				}
			l244:
				store32(m.memory[int64(uint32(v8))+88:], uint32(v32&i32(0xffff)))
				store32(m.memory[int64(uint32(v8))+100:], uint32(v23+v20&i32(255)))
				v22 = v9 - v20
				v25 = i64_shr_u(v25, int64(uint32(v20)))
				v20 = i32(19)
				if v34&i32(255) == 0 {
					goto l212
				}
				if v34&i32(32) != 0 {
					v9 = i32(12)
					m.memory[uint32(v8)] = byte(i32(12))
					store32(m.memory[int64(uint32(v8))+100:], uint32(i32(-1)))
					t536 := int64(load64(m.memory[int64(uint32(v7))+16:]))
					store64(m.memory[uint32(v19):], uint64(t536))
					t537 := int32(load32(m.memory[int64(uint32(v7))+24:]))
					store32(m.memory[int64(uint32(v19))+8:], uint32(t537))
					m.memory[int64(uint32(v8))+64] = byte(v22)
					store32(m.memory[int64(uint32(v8))+60:], uint32(v3))
					store32(m.memory[int64(uint32(v8))+56:], uint32(v24))
					store64(m.memory[int64(uint32(v8))+48:], uint64(v25))
					t538 := int32(load32(m.memory[int64(uint32(v7))+8:]))
					store32(m.memory[uint32(v12):], uint32(t538))
					t539 := int32(load32(m.memory[int64(uint32(v7))+11:]))
					store32(m.memory[int64(uint32(v12))+3:], uint32(t539))
					goto l18
				}
				if v34&i32(64) != 0 {
					goto l251
				}
				store32(m.memory[int64(uint32(v8))+96:], uint32(v34&i32(15)))
				v20 = i32(20)
				goto l212
			l251:
			}
			v20 = i32(30)
			m.memory[uint32(v8)] = byte(i32(30))
			t532 := int64(load64(m.memory[int64(uint32(v7))+16:]))
			store64(m.memory[uint32(v19):], uint64(t532))
			t533 := int32(load32(m.memory[int64(uint32(v7))+24:]))
			store32(m.memory[int64(uint32(v19))+8:], uint32(t533))
			m.memory[int64(uint32(v8))+64] = byte(v22)
			store32(m.memory[int64(uint32(v8))+60:], uint32(v3))
			store32(m.memory[int64(uint32(v8))+56:], uint32(v24))
			store64(m.memory[int64(uint32(v8))+48:], uint64(v25))
			t534 := int32(load32(m.memory[int64(uint32(v7))+8:]))
			store32(m.memory[uint32(v12):], uint32(t534))
			t535 := int32(load32(m.memory[int64(uint32(v7))+11:]))
			store32(m.memory[int64(uint32(v12))+3:], uint32(t535))
			store32(m.memory[int64(uint32(v8))+136:], uint32(i32(28)))
			store32(m.memory[int64(uint32(v8))+132:], uint32(i32(1065181)))
			v3 = i32(-3)
			goto l30
		}
	l207:
		m.fn39(v20, v28, i32(1283924))
		panic("unreachable")
	case 19:
		t381 := int32(load32(m.memory[int64(uint32(v8))+80:]))
		v5 = t381
		t382 := int32(load32(m.memory[int64(uint32(v8))+76:]))
		t383 := v5
		v3 = t382
		if t383 != v3 {
			if uint32(v5) >= uint32(v3) {
				m.fn39(v5, v3, i32(1279296))
				panic("unreachable")
			}
			t384 := int32(load32(m.memory[int64(uint32(v8))+72:]))
			t385 := int32(load32(m.memory[int64(uint32(v8))+88:]))
			m.memory[uint32(t384+v5)] = byte(t385)
			t386 := int32(load32(m.memory[int64(uint32(v8))+80:]))
			store32(m.memory[int64(uint32(v8))+80:], uint32(t386+i32(1)))
			goto l18
		}
		v3 = i32(0)
		v20 = i32(19)
		goto l30
	case 20:
		{
			{
				t365 := int32(load32(m.memory[int64(uint32(v8))+96:]))
				v29 = t365
				if v29 != 0 {
					goto l181
				}
				t366 := int32(load32(m.memory[int64(uint32(v8))+88:]))
				v9 = t366
				goto l182
			}
		l181:
			t367 := int64(load64(m.memory[int64(uint32(v8))+48:]))
			v25 = t367
			{
				{
					t368 := int32(m.memory[int64(uint32(v8))+64])
					t369 := v29
					v5 = t368
					if uint32(t369) > uint32(v5) {
						goto l183
					}
					v3 = v5
					goto l184
				}
			l183:
				t370 := int32(load32(m.memory[int64(uint32(v8))+56:]))
				v9 = t370
				t371 := int32(load32(m.memory[int64(uint32(v8))+60:]))
				v27 = t371
			l186:
				{
					if v9 != v27 {
						goto l185
					}
					v3 = i32(0)
					v20 = i32(20)
					goto l30
				l185:
					t372 := v8
					v20 = v9 + i32(1)
					store32(m.memory[int64(uint32(t372))+56:], uint32(v20))
					t373 := v8
					v3 = v5 + i32(8)
					m.memory[int64(uint32(t373))+64] = byte(v3)
					t374 := int64(m.memory[uint32(v9)])
					t375 := v8
					v25 = i64_shl(t374, int64(uint32(v5))) | v25
					store64(m.memory[int64(uint32(t375))+48:], uint64(v25))
					v9 = v20
					v5 = v3
					if uint32(v29) > uint32(v3&i32(255)) {
						goto l186
					}
				}
			}
		l184:
			m.memory[int64(uint32(v8))+64] = byte(v3 - v29)
			t376 := v8
			t377 := v25
			v30 = int64(uint32(v29))
			store64(m.memory[int64(uint32(t376))+48:], uint64(i64_shr_u(t377, v30)))
			t378 := int32(load32(m.memory[int64(uint32(v8))+100:]))
			store32(m.memory[int64(uint32(v8))+100:], uint32(t378+v29))
			t379 := int32(load32(m.memory[int64(uint32(v8))+88:]))
			t380 := v8
			v9 = t379 + int32(v25&(i64_shl(i64(-1), v30)^i64(-1)))
			store32(m.memory[int64(uint32(t380))+88:], uint32(v9))
		}
	l182:
		store32(m.memory[int64(uint32(v8))+104:], uint32(v9))
		v9 = i32(21)
		goto l18
	case 21:
		t79 := int64(load32(m.memory[int64(uint32(v8))+156:]))
		v26 = i64_shl(i64(-1), t79) ^ i64(-1)
		t80 := int32(load32(m.memory[int64(uint32(v8))+56:]))
		v5 = t80
		t81 := int32(m.memory[int64(uint32(v8))+64])
		v9 = t81
		t82 := int32(load32(m.memory[int64(uint32(v8))+60:]))
		v28 = t82
		t83 := int64(load64(m.memory[int64(uint32(v8))+48:]))
		v25 = t83
		t84 := int32(m.memory[int64(uint32(v8))+160])
		v27 = t84
		{
		l93:
			v3 = int32(v25 & v26)
			v20 = i32(32)
			v29 = i32(1279136)
			switch v27 {
			default:
				goto l83
			case 1:
				v20 = i32(1332)
				v29 = v16
				goto l83
			case 2:
				v20 = i32(1332)
				v29 = v17
				goto l83
			case 3:
				v20 = i32(592)
				v29 = v18
			}
		l83:
			if uint32(v20) <= uint32(v3) {
				m.fn39(v3, v20, i32(1284092))
				panic("unreachable")
			}
			{
				v20 = v29 + v3<<2
				t85 := int32(m.memory[int64(uint32(v20))+3])
				v3 = t85
				if uint32(v3) <= uint32(v9&i32(255)) {
					t86 := int32(load16(m.memory[uint32(v20):]))
					v23 = t86
					t87 := int32(m.memory[int64(uint32(v20))+2])
					v21 = t87
					if uint32(v21&i32(255)) <= uint32(i32(15)) {
						goto l91
					}
					v20 = v3
					goto l92
				}
				if v5 != v28 {
					t88 := int64(m.memory[uint32(v5)])
					v30 = t88
					t89 := v8
					v3 = v9 + i32(8)
					m.memory[int64(uint32(t89))+64] = byte(v3)
					t90 := v8
					v5 = v5 + i32(1)
					store32(m.memory[int64(uint32(t90))+56:], uint32(v5))
					t91 := v8
					v25 = i64_shl(v30, int64(uint32(v9))) | v25
					store64(m.memory[int64(uint32(t91))+48:], uint64(v25))
					v9 = v3
					goto l93
				}
				goto l90
			}
		l91:
			v24 = v3 & i32(31)
			v26 = i64_shl(i64(-1), int64(uint32(v3+v21))) ^ i64(-1)
		l100:
			v20 = i32_shr_u(int32(v25&v26), v24) + v23
			v29 = i32(32)
			v21 = i32(1279136)
			switch v27 {
			default:
				goto l94
			case 1:
				v29 = i32(1332)
				v21 = v16
				goto l94
			case 2:
				v29 = i32(1332)
				v21 = v17
				goto l94
			case 3:
				v29 = i32(592)
				v21 = v18
			}
		l94:
			if uint32(v20) >= uint32(v29) {
				m.fn39(v20, v29, i32(1284092))
				panic("unreachable")
			}
			{
				v29 = v21 + v20<<2
				t92 := int32(m.memory[int64(uint32(v29))+3])
				v20 = t92
				if uint32((v20+v3)&i32(255)) <= uint32(v9&i32(255)) {
					goto l99
				}
				if v5 == v28 {
					goto l90
				}
				t93 := int64(m.memory[uint32(v5)])
				v30 = t93
				t94 := v8
				v20 = v9 + i32(8)
				m.memory[int64(uint32(t94))+64] = byte(v20)
				t95 := v8
				v5 = v5 + i32(1)
				store32(m.memory[int64(uint32(t95))+56:], uint32(v5))
				t96 := v8
				v25 = i64_shl(v30, int64(uint32(v9))) | v25
				store64(m.memory[int64(uint32(t96))+48:], uint64(v25))
				v9 = v20
				goto l100
			}
		l99:
			t97 := int32(m.memory[int64(uint32(v29))+2])
			v21 = t97
			t98 := int32(load16(m.memory[uint32(v29):]))
			v23 = t98
			t99 := int32(load32(m.memory[int64(uint32(v8))+100:]))
			store32(m.memory[int64(uint32(v8))+100:], uint32(t99+v3))
			v9 = v9 - v3
			v25 = i64_shr_u(v25, int64(uint32(v3)))
		}
	l92:
		m.memory[int64(uint32(v8))+64] = byte(v9 - v20)
		store64(m.memory[int64(uint32(v8))+48:], uint64(i64_shr_u(v25, int64(uint32(v20)))))
		if v21&i32(64) != 0 {
			store32(m.memory[int64(uint32(v8))+136:], uint32(i32(22)))
			store32(m.memory[int64(uint32(v8))+132:], uint32(i32(1065209)))
			goto l57
		}
		store32(m.memory[int64(uint32(v8))+96:], uint32(v21&i32(15)))
		store32(m.memory[int64(uint32(v8))+92:], uint32(v23&i32(0xffff)))
		v9 = i32(22)
		goto l18
	l90:
		v3 = i32(0)
		v20 = i32(21)
		goto l30
	case 22:
		v9 = i32(23)
		t65 := int32(load32(m.memory[int64(uint32(v8))+96:]))
		v27 = t65
		if v27 == 0 {
			goto l18
		}
		t66 := int64(load64(m.memory[int64(uint32(v8))+48:]))
		v25 = t66
		{
			{
				t67 := int32(m.memory[int64(uint32(v8))+64])
				t68 := v27
				v3 = t67
				if uint32(t68) > uint32(v3) {
					goto l79
				}
				v20 = v3
				goto l80
			}
		l79:
			t69 := int32(load32(m.memory[int64(uint32(v8))+56:]))
			v5 = t69
			t70 := int32(load32(m.memory[int64(uint32(v8))+60:]))
			v28 = t70
		l82:
			{
				if v5 != v28 {
					goto l81
				}
				v3 = i32(0)
				v20 = i32(22)
				goto l30
			l81:
				t71 := v8
				v29 = v5 + i32(1)
				store32(m.memory[int64(uint32(t71))+56:], uint32(v29))
				t72 := v8
				v20 = v3 + i32(8)
				m.memory[int64(uint32(t72))+64] = byte(v20)
				t73 := int64(m.memory[uint32(v5)])
				t74 := v8
				v25 = i64_shl(t73, int64(uint32(v3))) | v25
				store64(m.memory[int64(uint32(t74))+48:], uint64(v25))
				v5 = v29
				v3 = v20
				if uint32(v27) > uint32(v20&i32(255)) {
					goto l82
				}
			}
		}
	l80:
		m.memory[int64(uint32(v8))+64] = byte(v20 - v27)
		t75 := v8
		t76 := v25
		v30 = int64(uint32(v27))
		store64(m.memory[int64(uint32(t75))+48:], uint64(i64_shr_u(t76, v30)))
		t77 := int32(load32(m.memory[int64(uint32(v8))+100:]))
		store32(m.memory[int64(uint32(v8))+100:], uint32(t77+v27))
		t78 := int32(load32(m.memory[int64(uint32(v8))+92:]))
		store32(m.memory[int64(uint32(v8))+92:], uint32(t78+int32(v25&(i64_shl(i64(-1), v30)^i64(-1)))))
		goto l18
	case 23:
	l180:
		{
			t334 := int32(load32(m.memory[int64(uint32(v8))+76:]))
			v3 = t334
			t335 := int32(load32(m.memory[int64(uint32(v8))+80:]))
			t336 := v3
			v5 = t335
			if t336 != v5 {
				v3 = v3 - v5
				{
					t337 := int32(load32(m.memory[int64(uint32(v8))+92:]))
					v20 = t337
					if uint32(v20) > uint32(v5) {
						v5 = v20 - v5
						t343 := int32(load32(m.memory[int64(uint32(v8))+16:]))
						if uint32(v5) > uint32(t343) {
							t362 := int32(m.memory[int64(uint32(v8))+1])
							if t362&i32(4) == 0 {
								m.fn34(i32(1279392), i32(85), i32(1284108))
								panic("unreachable")
							}
							v20 = i32(30)
							store32(m.memory[int64(uint32(v8))+136:], uint32(i32(30)))
							store32(m.memory[int64(uint32(v8))+132:], uint32(i32(1064767)))
							v3 = i32(-3)
							goto l30
						}
						t344 := int32(load32(m.memory[int64(uint32(v8))+12:]))
						v29 = t344
						v20 = v29 + i32(-64)
						if uint32(v20) >= uint32(i32(-63)) {
							m.fn2(i32(1283612), i32(74), i32(1283688))
							panic("unreachable")
						}
						t346 := v19
						t347 := v13
						p345 := v20
						if uint32(v20) > uint32(v29) {
							p345 = i32(0)
						}
						t348 := int32(load32(m.memory[int64(uint32(v8))+20:]))
						t349 := v5
						v20 = t348
						v29 = t349 - v20
						t350 := p345 - v29
						t351 := v20 - v5
						var p352 int32
						if uint32(v5) > uint32(v20) {
							p352 = 1
						}
						v20 = p352
						p353 := t351
						if v20 != 0 {
							p353 = t350
						}
						v27 = p353
						t354 := int32(load32(m.memory[int64(uint32(v8))+88:]))
						t355 := v27
						t356 := v27
						t357 := v3
						v28 = t354
						t359 := v28
						p358 := v5
						if v20 != 0 {
							p358 = v29
						}
						v5 = p358
						p360 := v5
						if uint32(v28) < uint32(v5) {
							p360 = t359
						}
						v5 = p360
						p361 := v5
						if uint32(v3) < uint32(v5) {
							p361 = t357
						}
						v5 = p361
						m.fn912(t346, t347, t355, t356+v5)
						goto l176
					}
					t338 := int32(load32(m.memory[int64(uint32(v8))+88:]))
					t339 := v19
					t340 := v20
					t341 := v3
					v5 = t338
					p342 := v5
					if uint32(v3) < uint32(v5) {
						p342 = t341
					}
					v5 = p342
					m.fn911(t339, t340, v5)
					goto l176
				}
			l176:
				t363 := int32(load32(m.memory[int64(uint32(v8))+88:]))
				t364 := v8
				v3 = t363
				store32(m.memory[int64(uint32(t364))+88:], uint32(v3-v5))
				if v3 == v5 {
					goto l18
				}
				goto l180
			}
			v3 = i32(0)
			v20 = i32(23)
			goto l30
		}
	case 24:
		t320 := int64(load64(m.memory[int64(uint32(v8))+48:]))
		v25 = t320
		{
			t321 := int32(m.memory[int64(uint32(v8))+64])
			v9 = t321
			if uint32(v9) > uint32(i32(13)) {
				goto l171
			}
			v3 = i32(0)
			v20 = i32(24)
			t322 := int32(load32(m.memory[int64(uint32(v8))+56:]))
			v5 = t322
			t323 := int32(load32(m.memory[int64(uint32(v8))+60:]))
			t324 := v5
			v29 = t323
			if t324 == v29 {
				goto l30
			}
			t325 := v8
			v27 = v5 + i32(1)
			store32(m.memory[int64(uint32(t325))+56:], uint32(v27))
			t326 := v8
			v30 = int64(uint32(v9))
			v26 = v30 + i64(8)
			m.memory[int64(uint32(t326))+64] = byte(v26)
			t327 := int64(m.memory[uint32(v5)])
			t328 := v8
			v25 = i64_shl(t327, v30) | v25
			store64(m.memory[int64(uint32(t328))+48:], uint64(v25))
			if uint32(v9) <= uint32(i32(5)) {
				goto l172
			}
			v9 = int32(v26)
			goto l171
		l172:
			if v27 == v29 {
				goto l30
			}
			store32(m.memory[int64(uint32(v8))+56:], uint32(v5+i32(2)))
			t329 := v8
			v9 = v9 | i32(16)
			m.memory[int64(uint32(t329))+64] = byte(v9)
			t330 := int64(m.memory[int64(uint32(v5))+1])
			t331 := v8
			v25 = i64_shl(t330, v26) | v25
			store64(m.memory[int64(uint32(t331))+48:], uint64(v25))
		}
	l171:
		m.memory[int64(uint32(v8))+64] = byte(v9 + i32(-14))
		store64(m.memory[int64(uint32(v8))+48:], uint64(int64(uint64(v25)>>14)))
		t332 := v8
		v9 = int32(v25)
		v5 = v9 & i32(31)
		store32(m.memory[int64(uint32(t332))+28:], uint32(v5+i32(257)))
		t333 := v8
		v3 = int32(uint32(v9)>>5) & i32(31)
		store32(m.memory[int64(uint32(t333))+32:], uint32(v3+i32(1)))
		store32(m.memory[int64(uint32(v8))+24:], uint32(int32(uint32(v9)>>10)&i32(15)+i32(4)))
		if uint32(v5) > uint32(i32(29)) {
			goto l173
		}
		if uint32(v3) > uint32(i32(29)) {
			goto l173
		}
		store32(m.memory[int64(uint32(v8))+36:], uint32(i32(0)))
		v9 = i32(25)
		goto l18
	l173:
		store32(m.memory[int64(uint32(v8))+136:], uint32(i32(36)))
		store32(m.memory[int64(uint32(v8))+132:], uint32(i32(1063922)))
		goto l57
	case 25:
		goto l26
	case 27:
		t300 := int64(load64(m.memory[int64(uint32(v8))+48:]))
		v25 = t300
		{
			t301 := int32(m.memory[int64(uint32(v8))+64])
			v9 = t301
			if uint32(v9) > uint32(i32(31)) {
				goto l170
			}
			v3 = i32(0)
			v20 = i32(27)
			t302 := int32(load32(m.memory[int64(uint32(v8))+56:]))
			v5 = t302
			t303 := int32(load32(m.memory[int64(uint32(v8))+60:]))
			t304 := v5
			v29 = t303
			if t304 == v29 {
				goto l30
			}
			t305 := v8
			v27 = v5 + i32(1)
			store32(m.memory[int64(uint32(t305))+56:], uint32(v27))
			t306 := v8
			v30 = int64(uint32(v9))
			v26 = v30 + i64(8)
			m.memory[int64(uint32(t306))+64] = byte(v26)
			t307 := int64(m.memory[uint32(v5)])
			t308 := v8
			v25 = i64_shl(t307, v30) | v25
			store64(m.memory[int64(uint32(t308))+48:], uint64(v25))
			if uint32(v9) > uint32(i32(23)) {
				goto l170
			}
			if v27 == v29 {
				goto l30
			}
			t309 := v8
			v27 = v5 + i32(2)
			store32(m.memory[int64(uint32(t309))+56:], uint32(v27))
			t310 := v8
			v31 = v30 + i64(16)
			m.memory[int64(uint32(t310))+64] = byte(v31)
			t311 := int64(m.memory[int64(uint32(v5))+1])
			t312 := v8
			v25 = i64_shl(t311, v26) | v25
			store64(m.memory[int64(uint32(t312))+48:], uint64(v25))
			if uint32(v9) > uint32(i32(15)) {
				goto l170
			}
			if v27 == v29 {
				goto l30
			}
			t313 := v8
			v27 = v5 + i32(3)
			store32(m.memory[int64(uint32(t313))+56:], uint32(v27))
			t314 := v8
			v30 = v30 + i64(24)
			m.memory[int64(uint32(t314))+64] = byte(v30)
			t315 := int64(m.memory[int64(uint32(v5))+2])
			t316 := v8
			v25 = i64_shl(t315, v31) | v25
			store64(m.memory[int64(uint32(t316))+48:], uint64(v25))
			if uint32(v9) > uint32(i32(7)) {
				goto l170
			}
			if v27 == v29 {
				goto l30
			}
			store32(m.memory[int64(uint32(v8))+56:], uint32(v5+i32(4)))
			m.memory[int64(uint32(v8))+64] = byte(v9 | i32(32))
			t317 := int64(m.memory[int64(uint32(v5))+3])
			t318 := v8
			v25 = i64_shl(t317, v30) | v25
			store64(m.memory[int64(uint32(t318))+48:], uint64(v25))
		}
	l170:
		m.memory[int64(uint32(v8))+64] = byte(i32(0))
		store64(m.memory[int64(uint32(v8))+48:], uint64(i64(0)))
		t319 := v8
		v9 = int32(v25)
		store32(m.memory[int64(uint32(t319))+124:], uint32(i32_rotr(v9&i32(0xff00ff), i32(8))|i32_rotr(v9, i32(24))&i32(0xff00ff)))
		v9 = i32(28)
		goto l18
	case 28:
		v3 = i32(2)
		{
			t299 := int32(m.memory[int64(uint32(v8))+1])
			if t299&i32(2) != 0 {
				store32(m.memory[int64(uint32(v8))+124:], uint32(i32(1)))
				v9 = i32(12)
				goto l18
			}
			v20 = i32(28)
			goto l30
		}
	case 29:
		goto l30
	case 30:
		store32(m.memory[int64(uint32(v8))+136:], uint32(i32(29)))
		store32(m.memory[int64(uint32(v8))+132:], uint32(i32(1065033)))
		goto l57
	default:
		v9 = i32(13)
		t277 := int32(m.memory[int64(uint32(v8))+3])
		v5 = t277
		if v5 == 0 {
			goto l18
		}
		t278 := int64(load64(m.memory[int64(uint32(v8))+48:]))
		v25 = t278
		{
			t279 := int32(m.memory[int64(uint32(v8))+64])
			v9 = t279
			if uint32(v9) > uint32(i32(15)) {
				goto l158
			}
			t280 := int32(load32(m.memory[int64(uint32(v8))+56:]))
			v3 = t280
			t281 := int32(load32(m.memory[int64(uint32(v8))+60:]))
			t282 := v3
			v20 = t281
			if t282 == v20 {
				goto l159
			}
			t283 := v8
			v29 = v3 + i32(1)
			store32(m.memory[int64(uint32(t283))+56:], uint32(v29))
			t284 := v8
			v30 = int64(uint32(v9))
			v26 = v30 + i64(8)
			m.memory[int64(uint32(t284))+64] = byte(v26)
			t285 := int64(m.memory[uint32(v3)])
			t286 := v8
			v25 = i64_shl(t285, v30) | v25
			store64(m.memory[int64(uint32(t286))+48:], uint64(v25))
			if uint32(v9) > uint32(i32(7)) {
				goto l158
			}
			if v29 == v20 {
				goto l159
			}
			store32(m.memory[int64(uint32(v8))+56:], uint32(v3+i32(2)))
			m.memory[int64(uint32(v8))+64] = byte(v9 | i32(16))
			t287 := int64(m.memory[int64(uint32(v3))+1])
			t288 := v8
			v25 = i64_shl(t287, v26) | v25
			store64(m.memory[int64(uint32(t288))+48:], uint64(v25))
		}
	l158:
		if v5&i32(2) == 0 {
			goto l160
		}
		if v25 == i64(35615) {
			{
				t293 := int32(m.memory[int64(uint32(v8))+2])
				if t293 != 0 {
					goto l165
				}
				m.memory[int64(uint32(v8))+2] = byte(i32(15))
			}
		l165:
			store16(m.memory[int64(uint32(v7))+16:], uint16(i32(35615)))
			t294 := m.fn910(i32(0), v7+i32(16), i32(2))
			store32(m.memory[int64(uint32(v8))+124:], uint32(t294))
			m.memory[int64(uint32(v8))+64] = byte(i32(0))
			store64(m.memory[int64(uint32(v8))+48:], uint64(i64(0)))
			v9 = i32(1)
			goto l18
		}
	l160:
		{
			t289 := int32(load32(m.memory[int64(uint32(v8))+140:]))
			v9 = t289
			if v9 == 0 {
				goto l162
			}
			store32(m.memory[int64(uint32(v9))+48:], uint32(i32(-1)))
			t290 := int32(m.memory[int64(uint32(v8))+3])
			v5 = t290
		}
	l162:
		{
			if v5&i32(1) == 0 {
				goto l163
			}
			t291 := int64(load64(m.memory[int64(uint32(v8))+48:]))
			v25 = t291
			t292 := int64(uint64(v25<<8&i64(0xff00)+int64(uint64(v25)>>8)) % uint64(i64(31)))
			if t292 == 0 {
				if v25&i64(15) != i64(8) {
					store32(m.memory[int64(uint32(v8))+136:], uint32(i32(27)))
					store32(m.memory[int64(uint32(v8))+132:], uint32(i32(1065362)))
					goto l57
				}
				t295 := v8
				v30 = int64(uint64(v25) >> 4)
				store64(m.memory[int64(uint32(t295))+48:], uint64(v30))
				t296 := int32(m.memory[int64(uint32(v8))+64])
				m.memory[int64(uint32(v8))+64] = byte(t296 + i32(-4))
				v3 = int32(v30) & i32(15)
				v9 = v3 + i32(8)
				{
					t297 := int32(m.memory[int64(uint32(v8))+2])
					v5 = t297
					if v5 != 0 {
						goto l167
					}
					m.memory[int64(uint32(v8))+2] = byte(v9)
					v5 = v9
				}
			l167:
				{
					if uint32(v3) > uint32(i32(7)) {
						goto l168
					}
					if uint32(v9) > uint32(v5) {
						goto l168
					}
					store64(m.memory[int64(uint32(v8))+120:], uint64(i64(0x100000000)))
					m.memory[int64(uint32(v8))+64] = byte(i32(0))
					store64(m.memory[int64(uint32(v8))+48:], uint64(i64(0)))
					store32(m.memory[int64(uint32(v8))+144:], uint32(i32_shl(i32(1), v9)))
					p298 := i32(27)
					if v25&i64(8192) == 0 {
						p298 = i32(12)
					}
					v9 = p298
					goto l18
				}
			l168:
				store32(m.memory[int64(uint32(v8))+136:], uint32(i32(20)))
				store32(m.memory[int64(uint32(v8))+132:], uint32(i32(1064864)))
				goto l57
			}
		}
	l163:
		store32(m.memory[int64(uint32(v8))+136:], uint32(i32(23)))
		store32(m.memory[int64(uint32(v8))+132:], uint32(i32(1064700)))
		goto l57
	l159:
		v20 = i32(0)
		v3 = i32(0)
		goto l30
	case 26:
		t5 := int32(load32(m.memory[int64(uint32(v8))+36:]))
		v21 = t5
		t6 := int32(load32(m.memory[int64(uint32(v8))+32:]))
		t7 := v21
		v22 = t6
		t8 := int32(load32(m.memory[int64(uint32(v8))+28:]))
		t9 := v22
		v23 = t8
		v24 = t9 + v23
		if uint32(t7) >= uint32(v24) {
			goto l32
		}
		t10 := int32(load32(m.memory[int64(uint32(v8))+56:]))
		v5 = t10
		t11 := int32(m.memory[int64(uint32(v8))+64])
		v9 = t11
		t12 := int64(load64(m.memory[int64(uint32(v8))+48:]))
		v25 = t12
	l72:
		{
			t13 := int64(load32(m.memory[int64(uint32(v8))+148:]))
			v26 = i64_shl(i64(-1), t13) ^ i64(-1)
			t14 := int32(load32(m.memory[int64(uint32(v8))+60:]))
			v27 = t14
			t15 := int32(m.memory[int64(uint32(v8))+152])
			v28 = t15
		l40:
			v3 = int32(v25 & v26)
			v20 = i32(512)
			v29 = i32(1277088)
			switch v28 {
			default:
				goto l33
			case 1:
				v20 = i32(1332)
				v29 = v16
				goto l33
			case 2:
				v20 = i32(1332)
				v29 = v17
				goto l33
			case 3:
				v20 = i32(592)
				v29 = v18
			}
		l33:
			if uint32(v20) <= uint32(v3) {
				m.fn39(v3, v20, i32(1284208))
				panic("unreachable")
			}
			{
				t16 := v9 & i32(255)
				v3 = v29 + v3<<2
				t17 := int32(m.memory[int64(uint32(v3))+3])
				v20 = t17
				if uint32(t16) >= uint32(v20) {
					{
						t22 := int32(load16(m.memory[uint32(v3):]))
						v3 = t22
						if uint32(v3) < uint32(i32(16)) {
							t50 := v8
							v9 = v9 - v20
							m.memory[int64(uint32(t50))+64] = byte(v9)
							t51 := v8
							v25 = i64_shr_u(v25, int64(uint32(v20)))
							store64(m.memory[int64(uint32(t51))+48:], uint64(v25))
							if uint32(v21) < uint32(i32(320)) {
								store16(m.memory[uint32(v15+v21<<1):], uint16(v3))
								v21 = v21 + i32(1)
								goto l70
							}
							m.fn39(v21, i32(320), i32(1284320))
							panic("unreachable")
						}
						v29 = v9 & i32(255)
						switch v3 + i32(-16) {
						default:
							v28 = v20 + i32(7)
							if uint32(v28) > uint32(v29) {
								goto l47
							}
							v29 = v5
							v3 = v9
							goto l46
						l47:
							{
								if v5 == v27 {
									goto l39
								}
								t23 := v8
								v29 = v5 + i32(1)
								store32(m.memory[int64(uint32(t23))+56:], uint32(v29))
								t24 := v8
								v3 = v9 + i32(8)
								m.memory[int64(uint32(t24))+64] = byte(v3)
								t25 := int64(m.memory[uint32(v5)])
								t26 := v8
								v25 = i64_shl(t25, int64(uint32(v9))) | v25
								store64(m.memory[int64(uint32(t26))+48:], uint64(v25))
								v5 = v29
								v9 = v3
								if uint32(v28) > uint32(v3&i32(255)) {
									goto l47
								}
							}
						l46:
							t27 := v8
							v9 = v3 - v20 + i32(-7)
							m.memory[int64(uint32(t27))+64] = byte(v9)
							t28 := v8
							v30 = i64_shr_u(v25, int64(uint32(v20)))
							v25 = int64(uint64(v30) >> 7)
							store64(m.memory[int64(uint32(t28))+48:], uint64(v25))
							v5 = int32(v30)&i32(127) + i32(11)
							v3 = v5 + v21
							if uint32(v3) > uint32(v24) {
								store32(m.memory[int64(uint32(v8))+136:], uint32(i32(26)))
								store32(m.memory[int64(uint32(v8))+132:], uint32(i32(1053006)))
								goto l57
							}
							if uint32(v21) > uint32(i32(320)) {
								m.fn127(v21, i32(320), i32(320), i32(1284304))
								panic("unreachable")
							}
							t29 := v5
							v20 = i32(320) - v21
							if uint32(t29) > uint32(v20) {
								m.fn127(i32(0), v5, v20, i32(1284288))
								panic("unreachable")
							}
							v5 = v5 << 1
							if v5 != 0 {
								memory_zero(m.memory, uint32(v15+v21<<1), uint32(v5))
								goto l52
							}
							goto l52
						case 0:
							v28 = v20 + i32(2)
							if uint32(v28) > uint32(v29) {
								goto l53
							}
							v29 = v9
							goto l54
						l53:
							v3 = v5
						l55:
							{
								if v3 == v27 {
									goto l39
								}
								t30 := v8
								v5 = v3 + i32(1)
								store32(m.memory[int64(uint32(t30))+56:], uint32(v5))
								t31 := v8
								v29 = v9 + i32(8)
								m.memory[int64(uint32(t31))+64] = byte(v29)
								t32 := int64(m.memory[uint32(v3)])
								t33 := v8
								v25 = i64_shl(t32, int64(uint32(v9))) | v25
								store64(m.memory[int64(uint32(t33))+48:], uint64(v25))
								v3 = v5
								v9 = v29
								if uint32(v28) > uint32(v29&i32(255)) {
									goto l55
								}
							}
						l54:
							t34 := v8
							v9 = v29 - v20
							m.memory[int64(uint32(t34))+64] = byte(v9)
							t35 := v8
							v30 = i64_shr_u(v25, int64(uint32(v20)))
							store64(m.memory[int64(uint32(t35))+48:], uint64(v30))
							if v21 != 0 {
								v3 = v21 + i32(-1)
								if uint32(v21) > uint32(i32(320)) {
									m.fn39(v3, i32(320), i32(1284224))
									panic("unreachable")
								}
								t43 := int32(load16(m.memory[uint32(v15+v3<<1):]))
								v3 = t43
								t44 := v8
								v9 = v9 + i32(-2)
								m.memory[int64(uint32(t44))+64] = byte(v9)
								t45 := v8
								v25 = int64(uint64(v30) >> 2)
								store64(m.memory[int64(uint32(t45))+48:], uint64(v25))
								v27 = int32(v30) & i32(3)
								v29 = v27 + i32(3)
								if uint32(v29+v21) <= uint32(v24) {
									t46 := v29
									v20 = i32(320) - v21
									if uint32(t46) > uint32(v20) {
										m.fn127(i32(0), v29, v20, i32(1284240))
										panic("unreachable")
									}
									v20 = v15 + v21<<1
									store16(m.memory[int64(uint32(v20))+4:], uint16(v3))
									store16(m.memory[int64(uint32(v20))+2:], uint16(v3))
									store16(m.memory[uint32(v20):], uint16(v3))
									if v27 == 0 {
										goto l69
									}
									store16(m.memory[int64(uint32(v20))+6:], uint16(v3))
									if v29 == i32(4) {
										goto l69
									}
									store16(m.memory[int64(uint32(v20))+8:], uint16(v3))
									if v29 == i32(5) {
										goto l69
									}
									store16(m.memory[int64(uint32(v20))+10:], uint16(v3))
								l69:
									t47 := int32(load32(m.memory[int64(uint32(v8))+36:]))
									v21 = t47 + v29
									t48 := int32(load32(m.memory[int64(uint32(v8))+32:]))
									v22 = t48
									t49 := int32(load32(m.memory[int64(uint32(v8))+28:]))
									v23 = t49
									goto l70
								}
								store32(m.memory[int64(uint32(v8))+136:], uint32(i32(26)))
								store32(m.memory[int64(uint32(v8))+132:], uint32(i32(1053006)))
								goto l57
							}
							store32(m.memory[int64(uint32(v8))+136:], uint32(i32(26)))
							store32(m.memory[int64(uint32(v8))+132:], uint32(i32(1053006)))
							goto l57
						case 1:
							v3 = v20 + i32(3)
							if uint32(v3) > uint32(v29) {
								goto l60
							}
							v28 = v5
							v29 = v9
							goto l59
						l60:
							{
								if v5 == v27 {
									goto l39
								}
								t36 := v8
								v28 = v5 + i32(1)
								store32(m.memory[int64(uint32(t36))+56:], uint32(v28))
								t37 := v8
								v29 = v9 + i32(8)
								m.memory[int64(uint32(t37))+64] = byte(v29)
								t38 := int64(m.memory[uint32(v5)])
								t39 := v8
								v25 = i64_shl(t38, int64(uint32(v9))) | v25
								store64(m.memory[int64(uint32(t39))+48:], uint64(v25))
								v5 = v28
								v9 = v29
								if uint32(v3) > uint32(v29&i32(255)) {
									goto l60
								}
							}
						l59:
							v3 = i32(-3)
							t40 := v8
							v9 = v29 - v20 + i32(-3)
							m.memory[int64(uint32(t40))+64] = byte(v9)
							t41 := v8
							v30 = i64_shr_u(v25, int64(uint32(v20)))
							v25 = int64(uint64(v30) >> 3)
							store64(m.memory[int64(uint32(t41))+48:], uint64(v25))
							v5 = int32(v30)&i32(7) + i32(3)
							v20 = v5 + v21
							if uint32(v20) > uint32(v24) {
								store32(m.memory[int64(uint32(v8))+136:], uint32(i32(26)))
								store32(m.memory[int64(uint32(v8))+132:], uint32(i32(1053006)))
								v20 = i32(30)
								goto l30
							}
							if uint32(v21) > uint32(i32(320)) {
								m.fn127(v21, i32(320), i32(320), i32(1284272))
								panic("unreachable")
							}
							t42 := v5
							v3 = i32(320) - v21
							if uint32(t42) > uint32(v3) {
								m.fn127(i32(0), v5, v3, i32(1284256))
								panic("unreachable")
							}
							v5 = v5 << 1
							if v5 != 0 {
								memory_zero(m.memory, uint32(v15+v21<<1), uint32(v5))
								goto l65
							}
							goto l65
						}
					}
				l65:
					v5 = v28
					v21 = v20
					goto l70
				l52:
					v5 = v29
					v21 = v3
				l70:
					store32(m.memory[int64(uint32(v8))+36:], uint32(v21))
					t52 := v21
					v24 = v23 + v22
					if uint32(t52) < uint32(v24) {
						goto l72
					}
					goto l32
				}
				if v5 == v27 {
					goto l39
				}
				t18 := int64(m.memory[uint32(v5)])
				v30 = t18
				t19 := v8
				v3 = v9 + i32(8)
				m.memory[int64(uint32(t19))+64] = byte(v3)
				t20 := v8
				v5 = v5 + i32(1)
				store32(m.memory[int64(uint32(t20))+56:], uint32(v5))
				t21 := v8
				v25 = i64_shl(v30, int64(uint32(v9))) | v25
				store64(m.memory[int64(uint32(t21))+48:], uint64(v25))
				v9 = v3
				goto l40
			}
		}
	case 10:
		v3 = i32(-4)
		v20 = i32(10)
		goto l30
	case 9:
		v3 = i32(-2)
		v20 = i32(9)
		goto l30
	}
l39:
	v3 = i32(0)
	v20 = i32(26)
	goto l30
l32:
	{
		t53 := int32(load16(m.memory[int64(uint32(v8))+13700:]))
		if t53 != 0 {
			if uint32(v23) >= uint32(i32(321)) {
				m.fn127(i32(0), v23, i32(320), i32(1284192))
				panic("unreachable")
			}
			m.fn908(v7+i32(16), i32(1), v15, v23, v17, i32(1332), i32(10), v14)
			t54 := int32(load32(m.memory[int64(uint32(v7))+16:]))
			if t54 != 0 {
				store32(m.memory[int64(uint32(v8))+136:], uint32(i32(28)))
				store32(m.memory[int64(uint32(v8))+132:], uint32(i32(1052805)))
				goto l57
			}
			t55 := int32(load32(m.memory[int64(uint32(v7))+24:]))
			v9 = t55
			t56 := int32(load32(m.memory[int64(uint32(v7))+20:]))
			store32(m.memory[int64(uint32(v8))+148:], uint32(t56))
			m.memory[int64(uint32(v8))+152] = byte(i32(2))
			store32(m.memory[int64(uint32(v8))+40:], uint32(v9))
			t57 := int32(load32(m.memory[int64(uint32(v8))+28:]))
			v9 = t57
			if uint32(v9) > uint32(i32(320)) {
				m.fn127(v9, i32(320), i32(320), i32(1284176))
				panic("unreachable")
			}
			t58 := int32(load32(m.memory[int64(uint32(v8))+32:]))
			v5 = t58
			t59 := v5
			v3 = i32(320) - v9
			if uint32(t59) > uint32(v3) {
				m.fn127(i32(0), v5, v3, i32(1284160))
				panic("unreachable")
			}
			m.fn908(v7+i32(16), i32(2), v15+v9<<1, v5, v18, i32(592), i32(9), v14)
			{
				t60 := int32(load32(m.memory[int64(uint32(v7))+16:]))
				if t60 != 0 {
					store32(m.memory[int64(uint32(v8))+136:], uint32(i32(22)))
					store32(m.memory[int64(uint32(v8))+132:], uint32(i32(1052883)))
					goto l57
				}
				t61 := int32(load32(m.memory[int64(uint32(v7))+24:]))
				v9 = t61
				t62 := int32(load32(m.memory[int64(uint32(v7))+20:]))
				v5 = t62
				m.memory[int64(uint32(v8))+160] = byte(i32(3))
				store32(m.memory[int64(uint32(v8))+156:], uint32(v5))
				t63 := int32(load32(m.memory[int64(uint32(v8))+40:]))
				store32(m.memory[int64(uint32(v8))+40:], uint32(v9+t63))
				v9 = i32(17)
				t64 := int32(m.memory[int64(uint32(v8))+4])
				if t64 != i32(6) {
					goto l18
				}
				v3 = i32(0)
				v20 = i32(17)
				goto l30
			}
		}
		store32(m.memory[int64(uint32(v8))+136:], uint32(i32(37)))
		store32(m.memory[int64(uint32(v8))+132:], uint32(i32(1064663)))
		goto l57
	}
l26:
	{
		t666 := int32(load32(m.memory[int64(uint32(v8))+36:]))
		v9 = t666
		t667 := int32(load32(m.memory[int64(uint32(v8))+24:]))
		t668 := v9
		v29 = t667
		if uint32(t668) >= uint32(v29) {
			goto l275
		}
		p669 := i32(19)
		if uint32(v9) > uint32(i32(19)) {
			p669 = v9
		}
		v27 = p669
		t670 := int32(load32(m.memory[int64(uint32(v8))+56:]))
		v20 = t670
		t671 := int64(load64(m.memory[int64(uint32(v8))+48:]))
		v30 = t671
		t672 := int32(m.memory[int64(uint32(v8))+64])
		v5 = t672
		t673 := int32(load32(m.memory[int64(uint32(v8))+60:]))
		v21 = t673
	l280:
		{
			{
				if uint32(v5&i32(255)) <= uint32(i32(2)) {
					goto l276
				}
				v25 = v30
				v3 = v5
				goto l277
			l276:
				if v20 != v21 {
					goto l278
				}
				v3 = i32(0)
				v20 = i32(25)
				goto l30
			l278:
				t674 := v8
				v28 = v20 + i32(1)
				store32(m.memory[int64(uint32(t674))+56:], uint32(v28))
				t675 := v8
				v3 = v5 | i32(8)
				m.memory[int64(uint32(t675))+64] = byte(v3)
				t676 := int64(m.memory[uint32(v20)])
				t677 := v8
				v25 = i64_shl(t676, int64(uint32(v5))&i64(255)) | v30
				store64(m.memory[int64(uint32(t677))+48:], uint64(v25))
				v20 = v28
			}
		l277:
			if v27 == v9 {
				m.fn39(v27, i32(19), i32(1284144))
				panic("unreachable")
			}
			t678 := v8
			v5 = v3 + i32(-3)
			m.memory[int64(uint32(t678))+64] = byte(v5)
			t679 := v8
			v30 = int64(uint64(v25) >> 3)
			store64(m.memory[int64(uint32(t679))+48:], uint64(v30))
			t680 := v8
			v3 = v9 + i32(1)
			store32(m.memory[int64(uint32(t680))+36:], uint32(v3))
			t681 := int32(m.memory[uint32(v9+i32(1284124))])
			store16(m.memory[uint32(v15+t681<<1):], uint16(int32(v25)&i32(7)))
			v9 = v3
			if v29 != v3 {
				goto l280
			}
		}
		v9 = v29
		goto l275
	}
l275:
	if uint32(v9) > uint32(i32(18)) {
		goto l281
	}
	v5 = v9
	v3 = v9 & i32(3)
	if v3 == i32(3) {
		goto l282
	}
	v5 = i32(0)
l283:
	{
		t682 := int32(m.memory[uint32(v9+v5+i32(1284124))])
		store16(m.memory[uint32(v15+t682<<1):], uint16(i32(0)))
		t683 := v3
		v5 = v5 + i32(1)
		if t683^v5 != i32(3) {
			goto l283
		}
	}
	v5 = v9 + v5
l282:
	if uint32(v9) > uint32(i32(15)) {
		goto l284
	}
l285:
	{
		t684 := int32(m.memory[uint32(v5+i32(1284124))])
		store16(m.memory[uint32(v15+t684<<1):], uint16(i32(0)))
		t685 := int32(m.memory[uint32(v5+i32(1284125))])
		store16(m.memory[uint32(v15+t685<<1):], uint16(i32(0)))
		t686 := int32(m.memory[uint32(v5+i32(1284126))])
		store16(m.memory[uint32(v15+t686<<1):], uint16(i32(0)))
		t687 := int32(m.memory[uint32(v5+i32(1284127))])
		store16(m.memory[uint32(v15+t687<<1):], uint16(i32(0)))
		v5 = v5 + i32(4)
		if v5 != i32(19) {
			goto l285
		}
	}
l284:
	store32(m.memory[int64(uint32(v8))+36:], uint32(i32(19)))
l281:
	m.fn908(v7+i32(16), i32(0), v15, i32(19), v16, i32(1332), i32(7), v14)
	{
		t688 := int32(load32(m.memory[int64(uint32(v7))+16:]))
		if t688 != 0 {
			goto l286
		}
		t689 := int32(load32(m.memory[int64(uint32(v7))+20:]))
		v9 = t689
		t690 := int32(load32(m.memory[int64(uint32(v7))+24:]))
		v5 = t690
		m.memory[int64(uint32(v8))+152] = byte(i32(1))
		store32(m.memory[int64(uint32(v8))+40:], uint32(v5))
		store32(m.memory[int64(uint32(v8))+148:], uint32(v9))
		store32(m.memory[int64(uint32(v8))+36:], uint32(i32(0)))
		v9 = i32(26)
		goto l18
	}
l286:
	store32(m.memory[int64(uint32(v8))+136:], uint32(i32(25)))
	store32(m.memory[int64(uint32(v8))+132:], uint32(i32(1052833)))
l57:
	v3 = i32(-3)
	v20 = i32(30)
l30:
	m.memory[uint32(v8)] = byte(v20)
	t691 := int32(load32(m.memory[int64(uint32(v1))+32:]))
	v29 = t691
	t692 := int32(load32(m.memory[int64(uint32(v1))+60:]))
	t693 := v1
	v8 = t692
	t694 := int32(load32(m.memory[int64(uint32(v8))+56:]))
	v5 = t694
	store32(m.memory[int64(uint32(t693))+32:], uint32(v5))
	t695 := int32(load32(m.memory[int64(uint32(v8))+60:]))
	store32(m.memory[int64(uint32(v1))+36:], uint32(t695-v5))
	t696 := int32(load32(m.memory[int64(uint32(v8))+76:]))
	t697 := v1
	v9 = t696
	t698 := int32(load32(m.memory[int64(uint32(v8))+80:]))
	t699 := v9
	v20 = t698
	store32(m.memory[int64(uint32(t697))+48:], uint32(t699-v20))
	t700 := int32(load32(m.memory[int64(uint32(v8))+72:]))
	t701 := v1
	t702 := v20
	v27 = t700
	v28 = t702 + v27
	store32(m.memory[int64(uint32(t701))+44:], uint32(v28))
	t703 := int32(load32(m.memory[int64(uint32(v1))+40:]))
	store32(m.memory[int64(uint32(v1))+40:], uint32(v5-v29+t703))
	t704 := int32(load32(m.memory[int64(uint32(v8))+84:]))
	t705 := int32(load32(m.memory[int64(uint32(v8))+116:]))
	t706 := v8
	v9 = v20 - v9 + t705
	v15 = t704 + v9
	store32(m.memory[int64(uint32(t706))+84:], uint32(v15))
	store32(m.memory[int64(uint32(v1))+52:], uint32(v15))
	t707 := int32(load32(m.memory[int64(uint32(v8))+124:]))
	t708 := v1
	v16 = t707
	store32(m.memory[int64(uint32(t708))+80:], uint32(v16))
	{
		{
			t709 := int32(load32(m.memory[uint32(v8+i32(12)):]))
			v15 = t709
			if uint32(v15+i32(-64)) >= uint32(i32(-63)) {
				m.fn2(i32(1283612), i32(74), i32(1283688))
				panic("unreachable")
			}
			{
				if uint32(v15) > uint32(i32(64)) {
					goto l288
				}
				if v9 == 0 {
					goto l289
				}
				t710 := int32(m.memory[uint32(v8)])
				v19 = t710
				if uint32(v19) > uint32(i32(30)) {
					goto l288
				}
				v19 = i32_shl(i32(1), v19)
				if v19&i32(0x40000600) != 0 {
					goto l289
				}
				if v19&i32(67584) == 0 {
					goto l288
				}
				t711 := int32(m.memory[int64(uint32(v8))+4])
				if t711 == i32(4) {
					goto l289
				}
			}
		l288:
			if uint32(v9) > uint32(v20) {
				m.fn127(i32(0), v9, v20, i32(1284432))
				panic("unreachable")
			}
			t712 := int32(m.memory[int64(uint32(v8))+3])
			v18 = t712 & i32(4)
			t713 := int32(load32(m.memory[int64(uint32(v8))+120:]))
			v21 = t713
			{
				t714 := v9
				v20 = v15 + i32(-64)
				p715 := v20
				if uint32(v20) > uint32(v15) {
					p715 = i32(0)
				}
				v20 = p715
				if uint32(t714) >= uint32(v20) {
					t732 := v27
					v19 = v9 - v20
					v15 = t732 + v19
					{
						if v18 == 0 {
							if v20 == 0 {
								goto l307
							}
							t736 := int32(load32(m.memory[int64(uint32(v8))+8:]))
							memory_copy(m.memory, uint32(t736), uint32(v15), uint32(v20))
							goto l307
						}
						if v21 != 0 {
							t737 := int32(load32(m.memory[int64(uint32(v8))+128:]))
							t738 := m.fn910(t737, v27, v19)
							v27 = t738
							t739 := int32(load32(m.memory[int64(uint32(v8))+8:]))
							v19 = t739
							t740 := m.fn910(v27, v15, v20)
							store32(m.memory[int64(uint32(v8))+128:], uint32(t740))
							if v20 == 0 {
								goto l307
							}
							memory_copy(m.memory, uint32(v19), uint32(v15), uint32(v20))
							goto l307
						}
						t733 := m.fn909(v16, v27, v19)
						v27 = t733
						{
							if v20 == 0 {
								goto l306
							}
							t734 := int32(load32(m.memory[int64(uint32(v8))+8:]))
							memory_copy(m.memory, uint32(t734), uint32(v15), uint32(v20))
						}
					l306:
						t735 := m.fn909(v27, v15, v20)
						store32(m.memory[int64(uint32(v8))+124:], uint32(t735))
						goto l307
					}
				}
				t716 := int32(load32(m.memory[int64(uint32(v8))+20:]))
				t717 := v9
				t718 := v9
				t719 := v20
				v17 = t716
				v13 = t719 - v17
				p720 := v13
				if uint32(v9) < uint32(v13) {
					p720 = t718
				}
				v19 = p720
				v23 = t717 - v19
				v12 = v27 + v19
				{
					if v18 != 0 {
						if uint32(v15) < uint32(v17) {
							m.fn127(v17, v15, v15, i32(1283752))
							panic("unreachable")
						}
						t723 := v19
						v15 = v15 - v17
						if uint32(t723) > uint32(v15) {
							m.fn127(i32(0), v19, v15, i32(1283736))
							panic("unreachable")
						}
						t724 := int32(load32(m.memory[int64(uint32(v8))+8:]))
						v18 = t724
						v15 = v18 + v17
						{
							if v21 != 0 {
								t728 := int32(load32(m.memory[int64(uint32(v8))+128:]))
								t729 := m.fn910(t728, v27, v19)
								store32(m.memory[int64(uint32(v8))+128:], uint32(t729))
								if v19 == 0 {
									goto l303
								}
								memory_copy(m.memory, uint32(v15), uint32(v27), uint32(v19))
							l303:
								if uint32(v9) <= uint32(v13) {
									goto l296
								}
								t730 := int32(load32(m.memory[int64(uint32(v8))+128:]))
								t731 := m.fn910(t730, v12, v23)
								store32(m.memory[int64(uint32(v8))+128:], uint32(t731))
								if v23 == 0 {
									goto l297
								}
								memory_copy(m.memory, uint32(v18), uint32(v12), uint32(v23))
								goto l297
							}
							if v19 == 0 {
								goto l301
							}
							memory_copy(m.memory, uint32(v15), uint32(v27), uint32(v19))
						l301:
							t725 := m.fn909(v16, v27, v19)
							t726 := v8
							v15 = t725
							store32(m.memory[int64(uint32(t726))+124:], uint32(v15))
							if uint32(v9) <= uint32(v13) {
								goto l296
							}
							if v23 == 0 {
								goto l302
							}
							memory_copy(m.memory, uint32(v18), uint32(v12), uint32(v23))
						l302:
							t727 := m.fn909(v15, v12, v23)
							store32(m.memory[int64(uint32(v8))+124:], uint32(t727))
							goto l297
						}
					}
					if uint32(v15) < uint32(v17) {
						m.fn127(v17, v15, v15, i32(1283720))
						panic("unreachable")
					}
					t721 := v19
					v15 = v15 - v17
					if uint32(t721) > uint32(v15) {
						m.fn127(i32(0), v19, v15, i32(1283704))
						panic("unreachable")
					}
					t722 := int32(load32(m.memory[int64(uint32(v8))+8:]))
					v15 = t722
					if v19 == 0 {
						goto l295
					}
					memory_copy(m.memory, uint32(v15+v17), uint32(v27), uint32(v19))
				l295:
					if uint32(v9) <= uint32(v13) {
						goto l296
					}
					if v23 == 0 {
						goto l297
					}
					memory_copy(m.memory, uint32(v15), uint32(v12), uint32(v23))
					goto l297
				}
			}
		}
	l307:
		store32(m.memory[int64(uint32(v8))+16:], uint32(v20))
		store32(m.memory[int64(uint32(v8))+20:], uint32(i32(0)))
		goto l289
	l297:
		store32(m.memory[int64(uint32(v8))+16:], uint32(v20))
		store32(m.memory[int64(uint32(v8))+20:], uint32(v23))
		goto l289
	l296:
		t741 := v8
		v15 = v19 + v17
		p742 := v15
		if v15 == v20 {
			p742 = i32(0)
		}
		store32(m.memory[int64(uint32(t741))+20:], uint32(p742))
		t743 := int32(load32(m.memory[int64(uint32(v8))+16:]))
		v15 = t743
		if uint32(v15) >= uint32(v20) {
			goto l289
		}
		store32(m.memory[int64(uint32(v8))+16:], uint32(v15+v19))
	}
l289:
	{
		t744 := int32(load32(m.memory[int64(uint32(v8))+132:]))
		v15 = t744
		if v15 == 0 {
			goto l308
		}
		{
			t745 := int32(load32(m.memory[int64(uint32(v8))+136:]))
			v20 = t745
			if v20 == 0 {
				goto l309
			}
			t746 := int32(m.memory[uint32(v15+v20+i32(-1))])
			if t746 == 0 {
				goto l310
			}
		}
	l309:
		m.fn2(i32(1284379), i32(37), i32(1284416))
		panic("unreachable")
	l310:
		store32(m.memory[int64(uint32(v1))+56:], uint32(v15))
	}
l308:
	t747 := int32(m.memory[int64(uint32(v8))+1])
	t748 := int32(m.memory[int64(uint32(v8))+64])
	v15 = t747<<6&i32(64) | t748
	{
		t749 := int32(m.memory[uint32(v8)])
		v8 = (t749 + i32(-12)) & i32(255)
		if uint32(v8) >= uint32(i32(6)) {
			goto l311
		}
		t750 := int32(load32(m.memory[int64(uint32(v8<<2))+1290720:]))
		v15 = t750 | v15
	}
l311:
	store32(m.memory[int64(uint32(v1))+76:], uint32(v15))
	t751 := int64(load64(m.memory[int64(uint32(v1))+16:]))
	t752 := v1
	v25 = t751 + int64(uint32(v5-v2))
	store64(m.memory[int64(uint32(t752))+16:], uint64(v25))
	t753 := int64(load64(m.memory[int64(uint32(v1))+24:]))
	t754 := v1
	v30 = t753 + int64(uint32(v28-v4))
	store64(m.memory[int64(uint32(t754))+24:], uint64(v30))
	{
		{
			if v3 != 0 {
				goto l312
			}
			var p755 int32
			if v6&i32(255) == i32(4) {
				p755 = 1
			}
			var p756 int32
			if v5 == v29 {
				p756 = 1
			}
			var p757 int32
			if v9 == 0 {
				p757 = 1
			}
			if p755|p756&p757 != 0 {
				goto l313
			}
		}
	l312:
		switch v3 + i32(5) {
		case 1, 2, 3:
			t762 := int64(load64(m.memory[uint32(v1):]))
			store64(m.memory[uint32(v1):], uint64(v25-v11+t762))
			t763 := int64(load64(m.memory[int64(uint32(v1))+8:]))
			store64(m.memory[int64(uint32(v1))+8:], uint64(v30-v10+t763))
			v8 = v3 + i32(-2)
			switch v8 {
			case 0:
				goto l320
			case 1:
				goto l319
			default:
				{
					{
						t764 := int32(load32(m.memory[int64(uint32(v1))+56:]))
						v8 = t764
						if v8 != 0 {
							goto l322
						}
						v1 = i32(0)
						goto l323
					}
				l322:
					t765 := m.fn981(v8)
					m.fn10(v7+i32(16), v8, t765)
					t766 := int32(load32(m.memory[int64(uint32(v7))+20:]))
					t767 := int32(load32(m.memory[int64(uint32(v7))+16:]))
					p768 := t766
					if t767 != 0 {
						p768 = i32(0)
					}
					v1 = p768
					t769 := int32(load32(m.memory[int64(uint32(v7))+24:]))
					v8 = t769
				}
			l323:
				store32(m.memory[int64(uint32(v0))+8:], uint32(v8))
				store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
				store32(m.memory[uint32(v0):], uint32(i32(0)))
				goto l324
			}
		case 5:
			goto l316
		case 7:
			t760 := int64(load64(m.memory[uint32(v1):]))
			store64(m.memory[uint32(v1):], uint64(v25-v11+t760))
			t761 := int64(load64(m.memory[int64(uint32(v1))+8:]))
			store64(m.memory[int64(uint32(v1))+8:], uint64(v30-v10+t761))
			goto l320
		default:
			goto l313
		case 6:
			v3 = i32(2)
			goto l316
		case 4:
			m.fn34(i32(1283816), i32(147), i32(1283892))
			panic("unreachable")
		}
	l313:
		v3 = i32(1)
	l316:
		t758 := int64(load64(m.memory[uint32(v1):]))
		store64(m.memory[uint32(v1):], uint64(v25-v11+t758))
		t759 := int64(load64(m.memory[int64(uint32(v1))+8:]))
		store64(m.memory[int64(uint32(v1))+8:], uint64(v30-v10+t759))
		goto l319
	}
l320:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v16))
	store32(m.memory[uint32(v0):], uint32(i32(1)))
	goto l324
l319:
	store32(m.memory[uint32(v0):], uint32(i32(2)))
	m.memory[int64(uint32(v0))+4] = byte(v3)
l324:
	m.g0 = v7 + i32(32)
}
func (m *Module) fn276(v0, v1 int32) int32 {
	var v2, v3, v4 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	t1 := int32(m.memory[uint32(v0)])
	t2 := v2
	v3 = t1
	t3 := int32(m.memory[int64(uint32(v3&i32(1)))+1122353])
	m.memory[int64(uint32(t2))+15] = byte(t3)
	v4 = int32(uint32(v3) >> 1)
	v0 = int32(bits.LeadingZeros32(uint32(v4))) + i32(-24)
	{
		if v4 != 0 {
			goto l0
		}
		v4 = v2 + i32(8) + i32(7)
		goto l1
	l0:
		t4 := int32(m.memory[uint32(v4&i32(1)+i32(1122353))])
		m.memory[int64(uint32(v2))+14] = byte(t4)
		if v0 != i32(7) {
			goto l2
		}
		v4 = v2 + i32(8) + i32(6)
		goto l1
	l2:
		t5 := int32(m.memory[uint32(int32(uint32(v3)>>2)&i32(1)+i32(1122353))])
		m.memory[int64(uint32(v2))+13] = byte(t5)
		if v0 != i32(6) {
			goto l3
		}
		v4 = v2 + i32(8) + i32(5)
		goto l1
	l3:
		t6 := int32(m.memory[uint32(int32(uint32(v3)>>3)&i32(1)+i32(1122353))])
		m.memory[int64(uint32(v2))+12] = byte(t6)
		if v0 != i32(5) {
			goto l4
		}
		v4 = v2 + i32(8) + i32(4)
		goto l1
	l4:
		t7 := int32(m.memory[uint32(int32(uint32(v3)>>4)&i32(1)+i32(1122353))])
		m.memory[int64(uint32(v2))+11] = byte(t7)
		if v0 != i32(4) {
			goto l5
		}
		v4 = v2 + i32(8) + i32(3)
		goto l1
	l5:
		t8 := int32(m.memory[uint32(int32(uint32(v3)>>5)&i32(1)+i32(1122353))])
		m.memory[int64(uint32(v2))+10] = byte(t8)
		if v0 != i32(3) {
			goto l6
		}
		v4 = v2 + i32(8) + i32(2)
		goto l1
	l6:
		t9 := int32(m.memory[uint32(int32(uint32(v3)>>6)&i32(1)+i32(1122353))])
		m.memory[int64(uint32(v2))+9] = byte(t9)
		if v0 != i32(2) {
			goto l7
		}
		v4 = v2 + i32(8) + i32(1)
		goto l1
	l7:
		t10 := int32(m.memory[int64(uint32(int32(uint32(v3)>>7)))+1122353])
		m.memory[int64(uint32(v2))+8] = byte(t10)
		v4 = v2 + i32(8)
	}
l1:
	t11 := m.fn312(v1, i32(1), i32(1122355), i32(2), v4, (i32(9)-v0)&i32(255))
	v0 = t11
	m.g0 = v2 + i32(16)
	return v0
}
func (m *Module) fn277(v0, v1 int32) int32 {
	var v2 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	t1 := int32(load16(m.memory[uint32(v0):]))
	store16(m.memory[int64(uint32(v2))+6:], uint16(t1))
	store64(m.memory[int64(uint32(v2))+8:], uint64(int64(uint32(i32(13)))<<32|int64(uint32(v2+i32(6)))))
	t2 := int32(load32(m.memory[uint32(v1):]))
	t3 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	t4 := m.fn51(t2, t3, i32(1276564), v2+i32(8))
	v1 = t4
	m.g0 = v2 + i32(16)
	return v1
}
func (m *Module) fn278(v0, v1, v2, v3 int32) {
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
							t20 := m.fn915(v6, v8, v4)
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
										m.fn39(v6, v3, i32(0x137888))
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
											m.fn39(v17, v4, i32(1276088))
											panic("unreachable")
										}
										if uint32(v7) >= uint32(v3) {
											m.fn39(v7, v3, i32(1276104))
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
									m.fn39(v5, v4, i32(1276056))
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
								m.fn39(v5, v3, i32(1276072))
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
								m.fn39(v6, v3, i32(1275976))
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
							m.fn39(v23, v4, i32(1275992))
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
								m.fn39(v5, v3, i32(1276008))
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
							m.fn39(v7, v3, i32(1276024))
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
func (m *Module) fn279(v0, v1 int32) int32 {
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
					t4 := int32(m.memory[int64(uint32(int32(v4)&i32(15)))+1098832])
					m.memory[uint32(v2+v0+i32(-2))] = byte(t4)
					v0 = v0 + i32(-1)
					v4 = int64(uint64(v4) >> 4)
					if v4 != i64(0) {
						goto l3
					}
				}
				t5 := m.fn312(v1, i32(1), i32(1122566), i32(2), v2+v0+i32(-1), i32(17)-v0)
				v0 = t5
				goto l2
			}
			if v3&i32(0x4000000) != 0 {
				goto l1
			}
			t2 := m.fn168(v0, v1)
			v0 = t2
			goto l2
		}
	l1:
		t6 := int64(load64(m.memory[uint32(v0):]))
		v4 = t6
		v0 = i32(17)
	l4:
		{
			t7 := int32(m.memory[int64(uint32(int32(v4)&i32(15)))+1122568])
			m.memory[uint32(v2+v0+i32(-2))] = byte(t7)
			v0 = v0 + i32(-1)
			v4 = int64(uint64(v4) >> 4)
			if v4 != i64(0) {
				goto l4
			}
		}
		t8 := m.fn312(v1, i32(1), i32(1122566), i32(2), v2+v0+i32(-1), i32(17)-v0)
		v0 = t8
	}
l2:
	m.g0 = v2 + i32(16)
	return v0
}
func (m *Module) fn280(v0, v1 int32) {
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
		t5 := int32(m.memory[int64(uint32(v5))+1276308])
		t6 := int32(m.memory[int64(uint32(v6))+1276308])
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
				m.fn971(v2+i32(8), v2+i32(24), v11)
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
			t15 := int32(m.memory[int64(uint32(v11))+1276308])
			v13 = t15
			t16 := v13
			v14 = v10 & i32(255)
			t17 := int32(m.memory[int64(uint32(v14))+1276308])
			if uint32(t16) < uint32(t17) {
				goto l4
			}
			{
				if v11 == v14 {
					goto l5
				}
				t18 := int32(m.memory[int64(uint32(v9&i32(255)))+1276308])
				if uint32(v13) >= uint32(t18) {
					goto l5
				}
				if uint32(v12) > uint32(i32(255)) {
					m.memory[int64(uint32(v2))+47] = byte(i32(2))
					m.fn48(i32(1284336), i32(43), v2+i32(47), i32(1275608), i32(1275624))
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
						m.fn39(v11, i32(4), i32(1275380))
						panic("unreachable")
					}
					t21 := int32(m.memory[uint32(v1+v11)])
					v15 = t21
					t22 := int32(m.memory[int64(uint32(v15))+1276308])
					if uint32(t22) <= uint32(i32(250)) {
						goto l11
					}
					v16 = i32(41)
					goto l12
				}
			l11:
				v12 = v7 & i32(255)
				if uint32(v12) >= uint32(i32(4)) {
					m.fn39(v12, i32(4), i32(1275396))
					panic("unreachable")
				}
				t23 := int32(m.memory[uint32(v1+v12)])
				v17 = v15<<16 | t23<<24 | v12<<8 | v11
				v16 = i32(42)
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
				m.fn39(v12, i32(4), i32(1275912))
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
				m.fn39(v12, i32(4), i32(1275912))
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
				m.fn34(i32(1271784), i32(19), i32(1271484))
				panic("unreachable")
			}
			p43 := v18
			if v12 != 0 {
				p43 = v7
			}
			v12 = p43
			if uint32(v12) > uint32(v10) {
				m.fn127(i32(0), v12, v10, i32(1271500))
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
		m.fn972(v2+i32(22), v2+i32(23))
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
		m.fn48(i32(1284336), i32(43), v2+i32(47), i32(1275608), i32(1275640))
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
	store32(m.memory[int64(uint32(v0))+24:], uint32(i32(43)))
	store32(m.memory[int64(uint32(v0))+16:], uint32(v11))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v12))
	store32(m.memory[uint32(v0):], uint32(v10))
	store32(m.memory[int64(uint32(v0))+40:], uint32((v5<<1+v6<<2+v4)<<1+v3))
	store64(m.memory[int64(uint32(v0))+8:], uint64(i64_shl(i64(1), int64(uint32(v3)))|(i64_shl(i64(1), int64(uint32(v4)))|(i64_shl(i64(1), int64(uint32(v5)))|i64_shl(i64(1), int64(uint32(v6)))))))
	m.g0 = v2 + i32(48)
}
func (m *Module) fn281(v0, v1, v2 int32) {
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
					t13 := int64(load64(m.memory[int64(uint32(i32(0)))+1276648:]))
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
						m.fn282(v8, v7)
						goto l12
					}
					m.fn282(v8, v7)
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
				t24 := int64(load64(m.memory[int64(uint32(i32(0)))+1276648:]))
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
						m.fn127(v8, v4, v4, i32(1276292))
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
			m.fn127(i32(0), v4, i32(1024), i32(1067880))
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
func (m *Module) fn282(v0, v1 int32) {
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
				m.fn2(i32(1273840), i32(46), i32(1273888))
				panic("unreachable")
			}
			if v3 == 0 {
				goto l4
			}
			if uint32(v4) > uint32(v2+i32(39)) {
				m.fn2(i32(1273904), i32(46), i32(1273952))
				panic("unreachable")
			}
		l4:
			m.fn1(v0)
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
			m.fn2(i32(1273840), i32(46), i32(1273888))
			panic("unreachable")
		}
		if v0 == 0 {
			goto l7
		}
		if uint32(v2) >= uint32(i32(52)) {
			m.fn2(i32(1273904), i32(46), i32(1273952))
			panic("unreachable")
		}
	l7:
		m.fn1(v1)
	}
}
func (m *Module) fn283(v0, v1, v2 int32) {
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
				m.fn7(i32(1), v2)
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
				t11 := int64(load64(m.memory[int64(uint32(i32(0)))+1276648:]))
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
			store32(m.memory[int64(uint32(v0))+4:], uint32(i32(1068848)))
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
		store32(m.memory[int64(uint32(v0))+4:], uint32(i32(1068848)))
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
				m.fn2(i32(1273840), i32(46), i32(1273888))
				panic("unreachable")
			}
			if v7 == 0 {
				goto l18
			}
			if uint32(v11) > uint32(v4+i32(39)) {
				m.fn2(i32(1273904), i32(46), i32(1273952))
				panic("unreachable")
			}
		l18:
			m.fn1(v0)
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
			m.fn2(i32(1273840), i32(46), i32(1273888))
			panic("unreachable")
		}
		if v0 == 0 {
			goto l21
		}
		if uint32(v4) >= uint32(i32(52)) {
			m.fn2(i32(1273904), i32(46), i32(1273952))
			panic("unreachable")
		}
	l21:
		m.fn1(v1)
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
			m.fn2(i32(1273840), i32(46), i32(1273888))
			panic("unreachable")
		}
		if v1 == 0 {
			goto l25
		}
		if uint32(v0) > uint32(v2+i32(39)) {
			m.fn2(i32(1273904), i32(46), i32(1273952))
			panic("unreachable")
		}
	l25:
		m.fn1(v3)
	}
}
func (m *Module) fn284(v0, v1, v2 int32) {
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
				m.fn295(v3+i32(28), i32(0), v2)
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
						t4 := int32(load32(m.memory[int64(uint32(v4))+1292744:]))
						v5 = t4
						t5 := int32(load32(m.memory[int64(uint32(v4))+1292232:]))
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
						m.fn295(v3+i32(28), v7, v5)
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
			m.fn10(v3+i32(4), v1, v2)
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
		store64(m.memory[int64(uint32(v3))+40:], uint64(int64(uint32(i32(44)))<<32|int64(uint32(v3+i32(16)))))
		m.fn14(v3+i32(28), i32(1052366), v3+i32(40))
		m.fn969(v3+i32(40), v3+i32(28))
		t16 := int64(load64(m.memory[int64(uint32(v3))+40:]))
		v11 = t16
		store32(m.memory[uint32(v0):], uint32(i32(-2)))
		store64(m.memory[int64(uint32(v0))+4:], uint64(v11))
	}
l11:
	m.g0 = v3 + i32(48)
}
func (m *Module) fn285(v0, v1, v2, v3, v4, v5 int32) {
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
			t0 := m.fn15(v2, v5*v1, v4, v3)
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
		t1 := m.fn27(v3, v4)
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
func (m *Module) fn286(v0, v1 int32) {
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
		t8 := int64(load64(m.memory[int64(uint32(i32(0)))+1276648:]))
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
func (m *Module) fn287(v0 int32) {
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
	m.fn285(t2, t4, t3, v2, i32(8), i32(32))
	{
		t6 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		if t6 != i32(1) {
			goto l0
		}
		t7 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		t8 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		m.fn7(t7, t8)
		panic("unreachable")
	}
l0:
	t9 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	v3 = t9
	store32(m.memory[uint32(v0):], uint32(v2))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
	m.g0 = v1 + i32(16)
}
func (m *Module) fn288(v0, v1, v2 int32) {
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
		t5 := int64(load64(m.memory[int64(uint32(i32(0)))+1276648:]))
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
			t10 := int64(load64(m.memory[int64(uint32(i32(0)))+1276648:]))
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
					store32(m.memory[int64(uint32(v0))+8:], uint32(i32(1275412)))
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
						m.fn7(i32(1), v6)
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
							t18 := int64(load64(m.memory[int64(uint32(i32(0)))+1276648:]))
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
						m.fn2(i32(1273840), i32(46), i32(1273888))
						panic("unreachable")
					}
					if v1 == 0 {
						goto l20
					}
					if uint32(v0) > uint32(v12+i32(34)) {
						m.fn2(i32(1273904), i32(46), i32(1273952))
						panic("unreachable")
					}
				l20:
					m.fn1(v2)
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
func (m *Module) fn289(v0, v1 int32) {
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
		t8 := int64(load64(m.memory[int64(uint32(i32(0)))+1276648:]))
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
func (m *Module) fn290(v0, v1, v2, v3 int32) {
	var v4 int32
	var v5 int64
	t0 := m.g0
	v4 = t0 - i32(32)
	m.g0 = v4
	store32(m.memory[int64(uint32(v4))+24:], uint32(i32(0)))
	store64(m.memory[int64(uint32(v4))+16:], uint64(i64(0)))
	m.fn274(v4+i32(16), v2, v3)
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
		v5 = int64(uint32(i32(15))) << 32
		store64(m.memory[int64(uint32(t5))+24:], uint64(v5|int64(uint32(v4+i32(12)))))
		store64(m.memory[int64(uint32(v4))+16:], uint64(v5|int64(uint32(v1+i32(8)))))
		m.fn14(v0, i32(1275820), v4+i32(16))
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
			m.fn2(i32(1273840), i32(46), i32(1273888))
			panic("unreachable")
		}
		if v1 == 0 {
			goto l3
		}
		if uint32(v2) > uint32(v0+i32(39)) {
			m.fn2(i32(1273904), i32(46), i32(1273952))
			panic("unreachable")
		}
	l3:
		m.fn1(v3)
	}
l1:
	m.g0 = v4 + i32(32)
}
func (m *Module) fn291(v0, v1, v2 int32) {
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
			m.fn7(i32(1), v2)
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
func (m *Module) fn292(v0, v1 int32) {
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
					t8 := m.fn15(v1, v2, i32(1), v3)
					v1 = t8
					if v1 != 0 {
						goto l1
					}
					m.fn7(i32(1), v3)
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
					m.fn2(i32(1273840), i32(46), i32(1273888))
					panic("unreachable")
				}
				if v4 == 0 {
					goto l4
				}
				if uint32(v5) > uint32(v2+i32(39)) {
					m.fn2(i32(1273904), i32(46), i32(1273952))
					panic("unreachable")
				}
			l4:
				m.fn1(v1)
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
func (m *Module) fn293(v0 int32) {
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
					m.fn2(i32(1273840), i32(46), i32(1273888))
					panic("unreachable")
				}
				if v3 == 0 {
					goto l6
				}
				if uint32(v4) > uint32(v1+i32(39)) {
					m.fn2(i32(1273904), i32(46), i32(1273952))
					panic("unreachable")
				}
			l6:
				m.fn1(v2)
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
				m.fn2(i32(1273840), i32(46), i32(1273888))
				panic("unreachable")
			}
			if v2 == 0 {
				goto l9
			}
			if uint32(v1) < uint32(i32(52)) {
				goto l9
			}
			m.fn2(i32(1273904), i32(46), i32(1273952))
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
				m.fn2(i32(1273840), i32(46), i32(1273888))
				panic("unreachable")
			}
			if v2 == 0 {
				goto l9
			}
			if uint32(v3) <= uint32(v1+i32(39)) {
				goto l9
			}
			m.fn2(i32(1273904), i32(46), i32(1273952))
			panic("unreachable")
		}
	}
l9:
	m.fn1(v0)
}
func (m *Module) fn294(v0, v1 int32) {
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
		t8 := int64(load64(m.memory[int64(uint32(i32(0)))+1276648:]))
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
func (m *Module) fn295(v0, v1, v2 int32) {
	var v3 int32
	t0 := m.g0
	v3 = t0 - i32(16)
	m.g0 = v3
	v1 = v2 + v1
	if uint32(v1) >= uint32(v2) {
		goto l0
	}
	m.fn7(i32(0), i32(0))
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
	m.fn285(t2, t4, t3, v2, i32(1), i32(1))
	{
		t8 := int32(load32(m.memory[int64(uint32(v3))+4:]))
		if t8 != i32(1) {
			goto l1
		}
		t9 := int32(load32(m.memory[int64(uint32(v3))+8:]))
		t10 := int32(load32(m.memory[int64(uint32(v3))+12:]))
		m.fn7(t9, t10)
		panic("unreachable")
	}
l1:
	t11 := int32(load32(m.memory[int64(uint32(v3))+8:]))
	v1 = t11
	store32(m.memory[uint32(v0):], uint32(v2))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
	m.g0 = v3 + i32(16)
}
func (m *Module) fn296(v0 int32) {
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
	m.fn214(t2, t4, t3, v2, i32(8), i32(176))
	{
		t6 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		if t6 != i32(1) {
			goto l0
		}
		t7 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		t8 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		m.fn7(t7, t8)
		panic("unreachable")
	}
l0:
	t9 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	v3 = t9
	store32(m.memory[uint32(v0):], uint32(v2))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
	m.g0 = v1 + i32(16)
}
func (m *Module) fn297(v0, v1, v2 int32) int32 {
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
					m.fn34(i32(1271248), i32(57), i32(1271232))
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
								m.fn39(v7, v2, i32(1275728))
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
							m.fn34(i32(1271248), i32(57), i32(1271232))
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
							m.fn39(v12, v2, i32(1275728))
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
					m.fn2(i32(1273840), i32(46), i32(1273888))
					panic("unreachable")
				}
				if v9 == 0 {
					goto l33
				}
				if uint32(v12) > uint32(v4+i32(39)) {
					m.fn2(i32(1273904), i32(46), i32(1273952))
					panic("unreachable")
				}
			l33:
				m.fn1(v7)
				return i32(-1)
			}
			m.fn30(i32(8), v9)
			panic("unreachable")
		}
	}
l21:
	return i32(-1)
l4:
	m.fn34(i32(1271248), i32(57), i32(1271232))
	panic("unreachable")
}
func (m *Module) fn298(v0 int32) {
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
	m.fn285(t2, t4, t3, v2, i32(8), i32(192))
	{
		t6 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		if t6 != i32(1) {
			goto l0
		}
		t7 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		t8 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		m.fn7(t7, t8)
		panic("unreachable")
	}
l0:
	t9 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	v3 = t9
	store32(m.memory[uint32(v0):], uint32(v2))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
	m.g0 = v1 + i32(16)
}
func (m *Module) fn299(v0, v1 int32) int32 {
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
					t12 := int32(m.memory[uint32(v3&i32(15)+i32(1098832))])
					m.memory[uint32(v2+i32(9)+v0+i32(-2))] = byte(t12)
					v0 = v0 + i32(-1)
					v3 = int32(uint32(v3)>>4) & i32(15)
					if v3 != 0 {
						goto l6
					}
				}
				t13 := m.fn312(v1, i32(1), i32(1122566), i32(2), v2+i32(9)+v0+i32(-1), i32(3)-v0)
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
				t7 := int32(load16(m.memory[int64(uint32((t6-v4*i32(100))&i32(255)<<1))+1100215:]))
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
				t9 := int32(m.memory[int64(uint32(v4<<1))+1100216])
				m.memory[uint32(t8+v3)] = byte(t9)
			}
		l4:
			t10 := m.fn312(v1, i32(1), i32(1), i32(0), v2+i32(11)+v3, i32(3)-v3)
			v0 = t10
			goto l5
		}
	l1:
		t14 := int32(m.memory[uint32(v0)])
		v3 = t14
		v0 = i32(3)
	l7:
		{
			t15 := int32(m.memory[uint32(v3&i32(15)+i32(1122568))])
			m.memory[uint32(v2+i32(14)+v0+i32(-2))] = byte(t15)
			v0 = v0 + i32(-1)
			v3 = int32(uint32(v3)>>4) & i32(15)
			if v3 != 0 {
				goto l7
			}
		}
		t16 := m.fn312(v1, i32(1), i32(1122566), i32(2), v2+i32(14)+v0+i32(-1), i32(3)-v0)
		v0 = t16
	}
l5:
	m.g0 = v2 + i32(16)
	return v0
}
func (m *Module) fn300(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v1):]))
	t1 := int32(m.memory[uint32(v0)])
	v0 = t1 << 2
	t2 := int32(load32(m.memory[int64(uint32(v0))+1290172:]))
	t3 := int32(load32(m.memory[int64(uint32(v0))+1290156:]))
	t4 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	t5 := int32(load32(m.memory[int64(uint32(t4))+12:]))
	t6 := m.t0[uint(t5)].(func(int32, int32, int32) int32)(t0, t2, t3)
	return t6
}
func (m *Module) fn301(v0, v1, v2, v3 int32) int32 {
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
												_ = m.fn845(v4<<6 | t63&i32(63) | v5<<18&i32(0x1c0000))
												return i32(255)
											}
											_ = m.fn845(v4 | v5<<12)
											return i32(255)
										}
										_ = m.fn845(v5<<6 | v4)
										return i32(255)
									}
									_ = m.fn845(v1 & i32(255))
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
							t48 := m.fn845(v1)
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
							t53 := m.fn845(v1)
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
func (m *Module) fn302(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t1 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	t2 := int32(load32(m.memory[uint32(v1):]))
	t3 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	t4 := m.fn58(t0, t1, t2, t3)
	return t4
}
func (m *Module) fn303(v0 int32) {
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
	m.fn214(t2, t4, t3, v2, i32(4), i32(8))
	{
		t6 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		if t6 != i32(1) {
			goto l0
		}
		t7 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		t8 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		m.fn7(t7, t8)
		panic("unreachable")
	}
l0:
	t9 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	v3 = t9
	store32(m.memory[uint32(v0):], uint32(v2))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
	m.g0 = v1 + i32(16)
}
func (m *Module) fn304(v0, v1, v2, v3 int32) {
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
					m.fn39(v17, v12, i32(1079444))
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
							m.fn127(i32(0), v14, v3, i32(1079460))
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
					m.fn14(v4+i32(20), i32(1048938), v4+i32(32))
					m.fn169(v10, i32(21), v4+i32(20))
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
					m.fn127(v17, v3, v3, i32(1068832))
					panic("unreachable")
				}
			}
		l0:
			t36 := int64(load64(m.memory[int64(uint32(i32(0)))+1276648:]))
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
				m.fn2(i32(1273840), i32(46), i32(1273888))
				panic("unreachable")
			}
			if v20 == 0 {
				goto l19
			}
			if uint32(v22) > uint32(v19+i32(39)) {
				m.fn2(i32(1273904), i32(46), i32(1273952))
				panic("unreachable")
			}
		l19:
			m.fn1(v14)
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
			m.fn2(i32(1273840), i32(46), i32(1273888))
			panic("unreachable")
		}
		if v14 == 0 {
			goto l22
		}
		if uint32(v19) >= uint32(i32(52)) {
			m.fn2(i32(1273904), i32(46), i32(1273952))
			panic("unreachable")
		}
	l22:
		m.fn1(v17)
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
func (m *Module) fn305(v0 int32) {
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
	m.fn979(t2, t4, t3, v2, i32(2), i32(2))
	{
		t6 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		if t6 != i32(1) {
			goto l0
		}
		t7 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		t8 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		m.fn7(t7, t8)
		panic("unreachable")
	}
l0:
	t9 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	v3 = t9
	store32(m.memory[uint32(v0):], uint32(v2))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
	m.g0 = v1 + i32(16)
}
func (m *Module) fn306(v0, v1 int32) int32 {
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
		t5 := int32(load16(m.memory[int64(uint32((t4-v4*i32(100))&i32(255)<<1))+1100215:]))
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
		t7 := int32(m.memory[int64(uint32(v4<<1))+1100216])
		m.memory[uint32(t6+v3)] = byte(t7)
	}
l2:
	t8 := m.fn312(v1, i32(1), i32(1), i32(0), v2+i32(13)+v3, i32(3)-v3)
	v3 = t8
	m.g0 = v2 + i32(16)
	return v3
}
func (m *Module) fn307(v0, v1, v2 int32) {
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
						m.fn7(i32(2), v12)
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
							m.fn725(v3+i32(48), v6, v5, i32(2), i32(2))
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
					t35 := int32(load32(m.memory[int64(uint32(v9))+1092824:]))
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
								m.fn164(v3+i32(48), v1, v2, v3+i32(32), v8)
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
										m.fn212(v3+i32(36), v18, v8, v10, v11, v12, i32(0))
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
							t37 := m.fn980(v3+i32(32), v1, v2)
							if t37 == 0 {
								goto l25
							}
							goto l24
						}
					l43:
						m.fn212(v3+i32(36), v18, v8, v10, v11, v12, i32(1))
					l42:
						t70 := int32(load32(m.memory[int64(uint32(v3))+36:]))
						if t70 == 0 {
							goto l24
						}
					}
				l25:
					store64(m.memory[int64(uint32(v3))+48:], uint64(int64(uint32(i32(45)))<<32|int64(uint32(v3+i32(16)))))
					m.fn14(v3+i32(20), i32(1064455), v3+i32(48))
					m.fn169(v0+i32(4), i32(20), v3+i32(20))
					store32(m.memory[uint32(v0):], uint32(i32(-1)))
					goto l50
				l30:
					m.fn44(v12, v11, v8, v11, i32(1092988))
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
			store64(m.memory[int64(uint32(v3))+48:], uint64(v28|int64(uint32(i32(1092840)))))
			m.fn14(v3+i32(4), i32(1066612), v3+i32(48))
			m.fn169(v0+i32(4), i32(20), v3+i32(4))
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
			m.fn2(i32(1273840), i32(46), i32(1273888))
			panic("unreachable")
		}
		if v8 == 0 {
			goto l62
		}
		if uint32(v10) > uint32(v1+i32(39)) {
			m.fn2(i32(1273904), i32(46), i32(1273952))
			panic("unreachable")
		}
	l62:
		m.fn1(v4)
	}
l52:
	m.g0 = v3 + i32(112)
}
func (m *Module) fn308(v0, v1, v2 int32) {
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
	t12 := int64(load64(m.memory[int64(uint32(i32(0)))+1276648:]))
	store64(m.memory[uint32(v0):], uint64(t12))
}
func (m *Module) fn309(v0, v1 int32) {
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
		t17 := int64(load64(m.memory[int64(uint32(i32(0)))+1276648:]))
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
		t32 := int64(load64(m.memory[int64(uint32(i32(0)))+1276648:]))
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
		t47 := int64(load64(m.memory[int64(uint32(i32(0)))+1276648:]))
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
		t62 := int64(load64(m.memory[int64(uint32(i32(0)))+1276648:]))
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
func (m *Module) fn310(v0, v1 int32) {
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
		t17 := int64(load64(m.memory[int64(uint32(i32(0)))+1276648:]))
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
func (m *Module) fn311(v0, v1 int32) {
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
		t17 := int64(load64(m.memory[int64(uint32(i32(0)))+1276648:]))
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
