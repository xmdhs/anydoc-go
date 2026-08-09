package core

import (
	"encoding/binary"
	"math"
	"math/bits"
	_ "embed"
)

func (m *Module) fn1842(v0, v1, v2, v3, v4 int32) {
	var v5 int32
	t0 := m.g0
	v5 = t0 - i32(16)
	m.g0 = v5
	m.fn1844(v5+i32(8), v0, v1, v2, v3, v4)
	{
		t1 := int32(load32(m.memory[int64(uint32(v5))+8:]))
		v4 = t1
		if v4 == i32(-1) {
			goto l0
		}
		t2 := int32(load32(m.memory[int64(uint32(v5))+12:]))
		m.fn2(v4, t2)
		panic("unreachable")
	}
l0:
	m.g0 = v5 + i32(16)
}
func (m *Module) fn1843(v0, v1, v2, v3 int32) {
	if v1 != v3 {
		m.fn1668(v1, v3, i32(1300940))
		panic("unreachable")
	}
	if v1 == 0 {
		return
	}
	memory_copy(m.memory, uint32(v0), uint32(v2), uint32(v1))
}
func (m *Module) fn1844(v0, v1, v2, v3, v4, v5 int32) {
	var v6, v7, v8, v9 int32
	var v10 int64
	t0 := m.g0
	v6 = t0 - i32(32)
	m.g0 = v6
	v2 = v3 + v2
	if uint32(v2) >= uint32(v3) {
		goto l0
	}
	v3 = i32(0)
	goto l9
l0:
	v7 = i32(0)
	v8 = v6 + i32(20)
	{
		{
			t1 := int32(load32(m.memory[uint32(v1):]))
			t2 := int64(uint32(v5))
			t3 := v2
			v9 = t1
			v3 = v9 << 1
			p4 := v3
			if uint32(v2) > uint32(v3) {
				p4 = t3
			}
			v3 = p4
			t6 := v3
			p5 := i32(4)
			if v5 == i32(1) {
				p5 = i32(8)
			}
			v2 = p5
			p7 := v2
			if uint32(v3) > uint32(v2) {
				p7 = t6
			}
			v2 = p7
			v10 = t2 * int64(uint32(v2))
			if int32(int64(uint64(v10)>>32)) != 0 {
				goto l2
			}
			v3 = int32(v10)
			if uint32(v3) > uint32(i32(-0x80000000)-v4) {
				goto l2
			}
			{
				if v9 != 0 {
					goto l3
				}
				v5 = i32(0)
				v7 = v6 + i32(28)
				goto l4
			l3:
				t8 := int32(load32(m.memory[int64(uint32(v1))+4:]))
				v8 = t8
				store32(m.memory[int64(uint32(v6))+28:], uint32(v4))
				v5 = v9 * v5
				v7 = v6 + i32(24)
			}
		l4:
			store32(m.memory[uint32(v7):], uint32(v5))
			{
				{
					t9 := int32(load32(m.memory[int64(uint32(v6))+28:]))
					if t9 == 0 {
						goto l5
					}
					{
						t10 := int32(load32(m.memory[int64(uint32(v6))+24:]))
						v5 = t10
						if v5 != 0 {
							t12 := m.fn89(v8, v5, v4, v3)
							v5 = t12
							goto l7
						}
						m.fn1819(v6+i32(8), v4, v3, i32(0))
						t11 := int32(load32(m.memory[int64(uint32(v6))+8:]))
						v5 = t11
						goto l7
					}
				}
			l5:
				m.fn1819(v6, v4, v3, i32(0))
				t13 := int32(load32(m.memory[uint32(v6):]))
				v5 = t13
			}
		l7:
			if v5 != 0 {
				goto l8
			}
			store32(m.memory[int64(uint32(v6))+20:], uint32(v4))
			v8 = v6 + i32(16)
			v7 = v3
		}
	l2:
		store32(m.memory[uint32(v8):], uint32(v7))
		t14 := int32(load32(m.memory[int64(uint32(v6))+16:]))
		v5 = t14
		t15 := int32(load32(m.memory[int64(uint32(v6))+20:]))
		v3 = t15
		goto l9
	}
l8:
	store32(m.memory[uint32(v1):], uint32(v2))
	store32(m.memory[int64(uint32(v1))+4:], uint32(v5))
	v3 = i32(-1)
l9:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
	store32(m.memory[uint32(v0):], uint32(v3))
	m.g0 = v6 + i32(32)
}
func (m *Module) fn1845(v0, v1, v2, v3 int32) {
	var v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14, v15, v16, v17, v18, v19, v20, v21, v22, v23, v24, v25, v26, v27, v28, v29, v30, v31, v32 int32
	t0 := m.g0
	v4 = t0 - i32(144)
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
			v6 = p3
			p4 := t2
			if v6 != 0 {
				p4 = t1
			}
			v7 = p4
			v8 = v1 - v5
			v9 = v0 + v5<<2
			v10 = v0
		l5:
			{
				{
					{
						if uint32(v7) > uint32(i32(12)) {
							goto l3
						}
						v11 = i32(1)
						if uint32(v7) <= uint32(i32(8)) {
							goto l4
						}
						t5 := int32(load32(m.memory[int64(uint32(v10))+32:]))
						t6 := v10
						v11 = t5
						t7 := int32(load32(m.memory[int64(uint32(v10))+16:]))
						t8 := v11
						v3 = t7
						p9 := v3
						if uint32(v11) > uint32(v3) {
							p9 = t8
						}
						v2 = p9
						t10 := int32(load32(m.memory[int64(uint32(v10))+12:]))
						t11 := v2
						v12 = t10
						t12 := int32(load32(m.memory[uint32(v10):]))
						t13 := v12
						v13 = t12
						p14 := v13
						if uint32(v12) > uint32(v13) {
							p14 = t13
						}
						v14 = p14
						p15 := v14
						if uint32(v2) > uint32(v14) {
							p15 = t11
						}
						v15 = p15
						t16 := int32(load32(m.memory[int64(uint32(v10))+28:]))
						t17 := v15
						v16 = t16
						t18 := int32(load32(m.memory[int64(uint32(v10))+4:]))
						t19 := v16
						v17 = t18
						p20 := v17
						if uint32(v16) > uint32(v17) {
							p20 = t19
						}
						v18 = p20
						t22 := v18
						p21 := v13
						if uint32(v12) < uint32(v13) {
							p21 = v12
						}
						v12 = p21
						p23 := v12
						if uint32(v18) > uint32(v12) {
							p23 = t22
						}
						v13 = p23
						p24 := v13
						if uint32(v15) > uint32(v13) {
							p24 = t17
						}
						v19 = p24
						t25 := int32(load32(m.memory[int64(uint32(v10))+24:]))
						t26 := v19
						v20 = t25
						t27 := int32(load32(m.memory[int64(uint32(v10))+20:]))
						t28 := v20
						v21 = t27
						t29 := int32(load32(m.memory[int64(uint32(v10))+8:]))
						t30 := v21
						v22 = t29
						p31 := v22
						if uint32(v21) > uint32(v22) {
							p31 = t30
						}
						v23 = p31
						p32 := v23
						if uint32(v20) > uint32(v23) {
							p32 = t28
						}
						v24 = p32
						t34 := v24
						p33 := v14
						if uint32(v2) < uint32(v14) {
							p33 = v2
						}
						v2 = p33
						t36 := v2
						p35 := v17
						if uint32(v16) < uint32(v17) {
							p35 = v16
						}
						v14 = p35
						p37 := v14
						if uint32(v2) > uint32(v14) {
							p37 = t36
						}
						v16 = p37
						p38 := v16
						if uint32(v24) > uint32(v16) {
							p38 = t34
						}
						v17 = p38
						p39 := v17
						if uint32(v19) > uint32(v17) {
							p39 = t26
						}
						store32(m.memory[int64(uint32(t6))+32:], uint32(p39))
						t41 := v10
						p40 := v23
						if uint32(v20) < uint32(v23) {
							p40 = v20
						}
						v20 = p40
						t43 := v20
						p42 := v3
						if uint32(v11) < uint32(v3) {
							p42 = v11
						}
						v11 = p42
						t45 := v11
						p44 := v22
						if uint32(v21) < uint32(v22) {
							p44 = v21
						}
						v3 = p44
						p46 := v3
						if uint32(v11) > uint32(v3) {
							p46 = t45
						}
						v21 = p46
						p47 := v21
						if uint32(v20) < uint32(v21) {
							p47 = t43
						}
						v22 = p47
						t49 := v22
						p48 := v14
						if uint32(v2) < uint32(v14) {
							p48 = v2
						}
						v2 = p48
						p50 := v2
						if uint32(v22) < uint32(v2) {
							p50 = t49
						}
						v14 = p50
						t52 := v14
						p51 := v3
						if uint32(v11) < uint32(v3) {
							p51 = v11
						}
						v11 = p51
						t54 := v11
						p53 := v12
						if uint32(v18) < uint32(v12) {
							p53 = v18
						}
						v3 = p53
						p55 := v3
						if uint32(v11) < uint32(v3) {
							p55 = t54
						}
						v12 = p55
						p56 := v12
						if uint32(v14) < uint32(v12) {
							p56 = t52
						}
						store32(m.memory[uint32(t41):], uint32(p56))
						t58 := v10
						p57 := v13
						if uint32(v15) < uint32(v13) {
							p57 = v15
						}
						v13 = p57
						t60 := v13
						p59 := v21
						if uint32(v20) > uint32(v21) {
							p59 = v20
						}
						v15 = p59
						p61 := v15
						if uint32(v13) > uint32(v15) {
							p61 = t60
						}
						v18 = p61
						t63 := v18
						p62 := v17
						if uint32(v19) < uint32(v17) {
							p62 = v19
						}
						v17 = p62
						p64 := v17
						if uint32(v18) > uint32(v17) {
							p64 = t63
						}
						store32(m.memory[int64(uint32(t58))+28:], uint32(p64))
						t66 := v10
						p65 := v17
						if uint32(v18) < uint32(v17) {
							p65 = v18
						}
						v17 = p65
						t68 := v17
						p67 := v15
						if uint32(v13) < uint32(v15) {
							p67 = v13
						}
						v13 = p67
						t70 := v13
						p69 := v16
						if uint32(v24) < uint32(v16) {
							p69 = v24
						}
						v15 = p69
						p71 := v15
						if uint32(v13) > uint32(v15) {
							p71 = t70
						}
						v16 = p71
						t73 := v16
						p72 := v2
						if uint32(v22) > uint32(v2) {
							p72 = v22
						}
						v2 = p72
						t75 := v2
						p74 := v3
						if uint32(v11) > uint32(v3) {
							p74 = v11
						}
						v11 = p74
						p76 := v11
						if uint32(v2) > uint32(v11) {
							p76 = t75
						}
						v3 = p76
						p77 := v3
						if uint32(v16) > uint32(v3) {
							p77 = t73
						}
						v18 = p77
						p78 := v18
						if uint32(v17) > uint32(v18) {
							p78 = t68
						}
						store32(m.memory[int64(uint32(t66))+24:], uint32(p78))
						t80 := v10
						p79 := v18
						if uint32(v17) < uint32(v18) {
							p79 = v17
						}
						store32(m.memory[int64(uint32(t80))+20:], uint32(p79))
						t82 := v10
						p81 := v3
						if uint32(v16) < uint32(v3) {
							p81 = v16
						}
						v3 = p81
						t84 := v3
						p83 := v15
						if uint32(v13) < uint32(v15) {
							p83 = v13
						}
						v13 = p83
						t86 := v13
						p85 := v11
						if uint32(v2) < uint32(v11) {
							p85 = v2
						}
						v11 = p85
						p87 := v11
						if uint32(v13) > uint32(v11) {
							p87 = t86
						}
						v2 = p87
						p88 := v2
						if uint32(v3) > uint32(v2) {
							p88 = t84
						}
						store32(m.memory[int64(uint32(t82))+16:], uint32(p88))
						t90 := v10
						p89 := v2
						if uint32(v3) < uint32(v2) {
							p89 = v3
						}
						store32(m.memory[int64(uint32(t90))+12:], uint32(p89))
						t92 := v10
						p91 := v11
						if uint32(v13) < uint32(v11) {
							p91 = v13
						}
						v11 = p91
						t94 := v11
						p93 := v12
						if uint32(v14) > uint32(v12) {
							p93 = v14
						}
						v3 = p93
						p95 := v3
						if uint32(v11) > uint32(v3) {
							p95 = t94
						}
						store32(m.memory[int64(uint32(t92))+8:], uint32(p95))
						t97 := v10
						p96 := v3
						if uint32(v11) < uint32(v3) {
							p96 = v11
						}
						store32(m.memory[int64(uint32(t97))+4:], uint32(p96))
						v11 = i32(9)
						goto l4
					}
				l3:
					t98 := int32(load32(m.memory[int64(uint32(v10))+48:]))
					t99 := v10
					v11 = t98
					t100 := int32(load32(m.memory[uint32(v10):]))
					t101 := v11
					v3 = t100
					p102 := v3
					if uint32(v11) > uint32(v3) {
						p102 = t101
					}
					v2 = p102
					t103 := int32(load32(m.memory[int64(uint32(v10))+44:]))
					t104 := v2
					v12 = t103
					t105 := int32(load32(m.memory[int64(uint32(v10))+20:]))
					t106 := v12
					v13 = t105
					p107 := v13
					if uint32(v12) > uint32(v13) {
						p107 = t106
					}
					v14 = p107
					t108 := int32(load32(m.memory[int64(uint32(v10))+16:]))
					t109 := v14
					v15 = t108
					p110 := v15
					if uint32(v14) > uint32(v15) {
						p110 = t109
					}
					v16 = p110
					p111 := v16
					if uint32(v2) > uint32(v16) {
						p111 = t104
					}
					v17 = p111
					t112 := int32(load32(m.memory[int64(uint32(v10))+40:]))
					t113 := v17
					v18 = t112
					t114 := int32(load32(m.memory[int64(uint32(v10))+4:]))
					t115 := v18
					v19 = t114
					p116 := v19
					if uint32(v18) > uint32(v19) {
						p116 = t115
					}
					v20 = p116
					t117 := int32(load32(m.memory[int64(uint32(v10))+32:]))
					t118 := v20
					v21 = t117
					t119 := int32(load32(m.memory[int64(uint32(v10))+24:]))
					t120 := v21
					v22 = t119
					p121 := v22
					if uint32(v21) > uint32(v22) {
						p121 = t120
					}
					v23 = p121
					p122 := v23
					if uint32(v20) > uint32(v23) {
						p122 = t118
					}
					v24 = p122
					t123 := int32(load32(m.memory[int64(uint32(v10))+36:]))
					t124 := v24
					v25 = t123
					t125 := int32(load32(m.memory[int64(uint32(v10))+8:]))
					t126 := v25
					v26 = t125
					p127 := v26
					if uint32(v25) > uint32(v26) {
						p127 = t126
					}
					v27 = p127
					t128 := int32(load32(m.memory[int64(uint32(v10))+28:]))
					t129 := v27
					v28 = t128
					t130 := int32(load32(m.memory[int64(uint32(v10))+12:]))
					t131 := v28
					v29 = t130
					p132 := v29
					if uint32(v28) > uint32(v29) {
						p132 = t131
					}
					v30 = p132
					p133 := v30
					if uint32(v27) > uint32(v30) {
						p133 = t129
					}
					v31 = p133
					p134 := v31
					if uint32(v24) > uint32(v31) {
						p134 = t124
					}
					v32 = p134
					p135 := v32
					if uint32(v17) > uint32(v32) {
						p135 = t113
					}
					store32(m.memory[int64(uint32(t99))+48:], uint32(p135))
					t137 := v10
					p136 := v16
					if uint32(v2) < uint32(v16) {
						p136 = v2
					}
					v2 = p136
					t139 := v2
					p138 := v23
					if uint32(v20) < uint32(v23) {
						p138 = v20
					}
					v16 = p138
					t141 := v16
					p140 := v30
					if uint32(v27) < uint32(v30) {
						p140 = v27
					}
					v20 = p140
					p142 := v20
					if uint32(v16) > uint32(v20) {
						p142 = t141
					}
					v23 = p142
					p143 := v23
					if uint32(v2) > uint32(v23) {
						p143 = t139
					}
					v27 = p143
					t145 := v27
					p144 := v22
					if uint32(v21) < uint32(v22) {
						p144 = v21
					}
					v21 = p144
					t147 := v21
					p146 := v19
					if uint32(v18) < uint32(v19) {
						p146 = v18
					}
					v18 = p146
					p148 := v18
					if uint32(v21) > uint32(v18) {
						p148 = t147
					}
					v19 = p148
					t150 := v19
					p149 := v29
					if uint32(v28) < uint32(v29) {
						p149 = v28
					}
					v22 = p149
					t152 := v22
					p151 := v26
					if uint32(v25) < uint32(v26) {
						p151 = v25
					}
					v25 = p151
					p153 := v25
					if uint32(v22) > uint32(v25) {
						p153 = t152
					}
					v26 = p153
					p154 := v26
					if uint32(v19) > uint32(v26) {
						p154 = t150
					}
					v28 = p154
					t156 := v28
					p155 := v15
					if uint32(v14) < uint32(v15) {
						p155 = v14
					}
					v14 = p155
					t158 := v14
					p157 := v3
					if uint32(v11) < uint32(v3) {
						p157 = v11
					}
					v11 = p157
					p159 := v11
					if uint32(v14) > uint32(v11) {
						p159 = t158
					}
					v3 = p159
					p160 := v3
					if uint32(v28) > uint32(v3) {
						p160 = t156
					}
					v15 = p160
					p161 := v15
					if uint32(v27) > uint32(v15) {
						p161 = t145
					}
					v29 = p161
					t163 := v29
					p162 := v32
					if uint32(v17) < uint32(v32) {
						p162 = v17
					}
					v17 = p162
					t165 := v17
					p164 := v31
					if uint32(v24) < uint32(v31) {
						p164 = v24
					}
					v24 = p164
					t167 := v24
					p166 := v13
					if uint32(v12) < uint32(v13) {
						p166 = v12
					}
					v12 = p166
					p168 := v12
					if uint32(v24) > uint32(v12) {
						p168 = t167
					}
					v13 = p168
					p169 := v13
					if uint32(v17) > uint32(v13) {
						p169 = t165
					}
					v30 = p169
					p170 := v30
					if uint32(v29) > uint32(v30) {
						p170 = t163
					}
					store32(m.memory[int64(uint32(t137))+44:], uint32(p170))
					t172 := v10
					p171 := v25
					if uint32(v22) < uint32(v25) {
						p171 = v22
					}
					v22 = p171
					t174 := v22
					p173 := v18
					if uint32(v21) < uint32(v18) {
						p173 = v21
					}
					v18 = p173
					p175 := v18
					if uint32(v22) < uint32(v18) {
						p175 = t174
					}
					v21 = p175
					t177 := v21
					p176 := v12
					if uint32(v24) < uint32(v12) {
						p176 = v24
					}
					v12 = p176
					t179 := v12
					p178 := v11
					if uint32(v14) < uint32(v11) {
						p178 = v14
					}
					v11 = p178
					p180 := v11
					if uint32(v12) < uint32(v11) {
						p180 = t179
					}
					v14 = p180
					p181 := v14
					if uint32(v21) < uint32(v14) {
						p181 = t177
					}
					store32(m.memory[uint32(t172):], uint32(p181))
					t183 := v10
					p182 := v30
					if uint32(v29) < uint32(v30) {
						p182 = v29
					}
					v24 = p182
					t185 := v24
					p184 := v13
					if uint32(v17) < uint32(v13) {
						p184 = v17
					}
					v13 = p184
					t187 := v13
					p186 := v15
					if uint32(v27) < uint32(v15) {
						p186 = v27
					}
					v15 = p186
					p188 := v15
					if uint32(v13) > uint32(v15) {
						p188 = t187
					}
					v17 = p188
					p189 := v17
					if uint32(v24) > uint32(v17) {
						p189 = t185
					}
					store32(m.memory[int64(uint32(t183))+40:], uint32(p189))
					t191 := v10
					p190 := v20
					if uint32(v16) < uint32(v20) {
						p190 = v16
					}
					v16 = p190
					t193 := v16
					p192 := v3
					if uint32(v28) < uint32(v3) {
						p192 = v28
					}
					v3 = p192
					p194 := v3
					if uint32(v16) < uint32(v3) {
						p194 = t193
					}
					v20 = p194
					t196 := v20
					p195 := v11
					if uint32(v12) > uint32(v11) {
						p195 = v12
					}
					v11 = p195
					t198 := v11
					p197 := v18
					if uint32(v22) > uint32(v18) {
						p197 = v22
					}
					v12 = p197
					p199 := v12
					if uint32(v11) < uint32(v12) {
						p199 = t198
					}
					v18 = p199
					p200 := v18
					if uint32(v20) < uint32(v18) {
						p200 = t196
					}
					v22 = p200
					t202 := v22
					p201 := v23
					if uint32(v2) < uint32(v23) {
						p201 = v2
					}
					v2 = p201
					t204 := v2
					p203 := v26
					if uint32(v19) < uint32(v26) {
						p203 = v19
					}
					v19 = p203
					p205 := v19
					if uint32(v2) < uint32(v19) {
						p205 = t204
					}
					v23 = p205
					t207 := v23
					p206 := v14
					if uint32(v21) > uint32(v14) {
						p206 = v21
					}
					v14 = p206
					p208 := v14
					if uint32(v23) < uint32(v14) {
						p208 = t207
					}
					v21 = p208
					p209 := v21
					if uint32(v22) < uint32(v21) {
						p209 = t202
					}
					store32(m.memory[int64(uint32(t191))+4:], uint32(p209))
					t211 := v10
					p210 := v17
					if uint32(v24) < uint32(v17) {
						p210 = v24
					}
					v17 = p210
					t213 := v17
					p212 := v19
					if uint32(v2) > uint32(v19) {
						p212 = v2
					}
					v2 = p212
					t215 := v2
					p214 := v3
					if uint32(v16) > uint32(v3) {
						p214 = v16
					}
					v3 = p214
					p216 := v3
					if uint32(v2) > uint32(v3) {
						p216 = t215
					}
					v16 = p216
					t218 := v16
					p217 := v15
					if uint32(v13) < uint32(v15) {
						p217 = v13
					}
					v13 = p217
					t220 := v13
					p219 := v12
					if uint32(v11) > uint32(v12) {
						p219 = v11
					}
					v11 = p219
					p221 := v11
					if uint32(v13) > uint32(v11) {
						p221 = t220
					}
					v12 = p221
					p222 := v12
					if uint32(v16) > uint32(v12) {
						p222 = t218
					}
					v15 = p222
					p223 := v15
					if uint32(v17) > uint32(v15) {
						p223 = t213
					}
					store32(m.memory[int64(uint32(t211))+36:], uint32(p223))
					t225 := v10
					p224 := v15
					if uint32(v17) < uint32(v15) {
						p224 = v17
					}
					store32(m.memory[int64(uint32(t225))+32:], uint32(p224))
					t227 := v10
					p226 := v3
					if uint32(v2) < uint32(v3) {
						p226 = v2
					}
					v3 = p226
					t229 := v3
					p228 := v11
					if uint32(v13) < uint32(v11) {
						p228 = v13
					}
					v11 = p228
					p230 := v11
					if uint32(v3) > uint32(v11) {
						p230 = t229
					}
					v2 = p230
					t232 := v2
					p231 := v12
					if uint32(v16) < uint32(v12) {
						p231 = v16
					}
					v12 = p231
					p233 := v12
					if uint32(v2) > uint32(v12) {
						p233 = t232
					}
					store32(m.memory[int64(uint32(t227))+28:], uint32(p233))
					t235 := v10
					p234 := v18
					if uint32(v20) > uint32(v18) {
						p234 = v20
					}
					v13 = p234
					t237 := v13
					p236 := v14
					if uint32(v23) > uint32(v14) {
						p236 = v23
					}
					v14 = p236
					p238 := v14
					if uint32(v13) < uint32(v14) {
						p238 = t237
					}
					v15 = p238
					t240 := v15
					p239 := v21
					if uint32(v22) > uint32(v21) {
						p239 = v22
					}
					v16 = p239
					p241 := v16
					if uint32(v15) < uint32(v16) {
						p241 = t240
					}
					store32(m.memory[int64(uint32(t235))+8:], uint32(p241))
					t243 := v10
					p242 := v12
					if uint32(v2) < uint32(v12) {
						p242 = v2
					}
					v2 = p242
					t245 := v2
					p244 := v11
					if uint32(v3) < uint32(v11) {
						p244 = v3
					}
					v11 = p244
					t247 := v11
					p246 := v14
					if uint32(v13) > uint32(v14) {
						p246 = v13
					}
					v3 = p246
					p248 := v3
					if uint32(v11) > uint32(v3) {
						p248 = t247
					}
					v12 = p248
					p249 := v12
					if uint32(v2) > uint32(v12) {
						p249 = t245
					}
					store32(m.memory[int64(uint32(t243))+24:], uint32(p249))
					t251 := v10
					p250 := v12
					if uint32(v2) < uint32(v12) {
						p250 = v2
					}
					store32(m.memory[int64(uint32(t251))+20:], uint32(p250))
					t253 := v10
					p252 := v3
					if uint32(v11) < uint32(v3) {
						p252 = v11
					}
					v11 = p252
					t255 := v11
					p254 := v16
					if uint32(v15) > uint32(v16) {
						p254 = v15
					}
					v3 = p254
					p256 := v3
					if uint32(v11) > uint32(v3) {
						p256 = t255
					}
					store32(m.memory[int64(uint32(t253))+16:], uint32(p256))
					t258 := v10
					p257 := v3
					if uint32(v11) < uint32(v3) {
						p257 = v11
					}
					store32(m.memory[int64(uint32(t258))+12:], uint32(p257))
					v11 = i32(13)
				}
			l4:
				m.fn1179(v10, v7, v11)
				if v6 != 0 {
					goto l2
				}
				var p259 int32
				if v10 == v0 {
					p259 = 1
				}
				v11 = p259
				v7 = v8
				v10 = v9
				if v11 != 0 {
					goto l5
				}
			}
			v3 = v9 + i32(-4)
			t260 := v0
			v10 = v1<<2 + i32(-4)
			v2 = t260 + v10
			v13 = v4 + i32(12) + v10
			v12 = v4 + i32(12)
			v11 = v0
		l9:
			if v5 != 0 {
				t269 := int32(load32(m.memory[uint32(v9):]))
				t270 := v12
				v14 = t269
				t271 := int32(load32(m.memory[uint32(v11):]))
				t272 := v14
				v15 = t271
				t273 := v15
				var p274 int32
				if uint32(v14) < uint32(v15) {
					p274 = 1
				}
				v16 = p274
				p275 := t273
				if v16 != 0 {
					p275 = t272
				}
				store32(m.memory[uint32(t270):], uint32(p275))
				t276 := int32(load32(m.memory[uint32(v2):]))
				t277 := v13
				v10 = t276
				t278 := int32(load32(m.memory[uint32(v3):]))
				t279 := v10
				v7 = t278
				p280 := v7
				if uint32(v10) > uint32(v7) {
					p280 = t279
				}
				store32(m.memory[uint32(t277):], uint32(p280))
				v5 = v5 + i32(-1)
				v13 = v13 + i32(-4)
				v12 = v12 + i32(4)
				t282 := v3
				p281 := i32(0)
				if uint32(v10) < uint32(v7) {
					p281 = i32(-4)
				}
				v3 = t282 + p281
				t284 := v2
				p283 := i32(0)
				if uint32(v10) >= uint32(v7) {
					p283 = i32(-4)
				}
				v2 = t284 + p283
				t285 := v11
				var p286 int32
				if uint32(v14) >= uint32(v15) {
					p286 = 1
				}
				v11 = t285 + p286<<2
				v9 = v9 + v16<<2
				goto l9
			}
			v10 = v3 + i32(4)
			{
				if v1&i32(1) == 0 {
					goto l7
				}
				t261 := v12
				t262 := v11
				t263 := v9
				var p264 int32
				if uint32(v11) < uint32(v10) {
					p264 = 1
				}
				v7 = p264
				p265 := t263
				if v7 != 0 {
					p265 = t262
				}
				t266 := int32(load32(m.memory[uint32(p265):]))
				store32(m.memory[uint32(t261):], uint32(t266))
				t267 := v9
				var p268 int32
				if uint32(v11) >= uint32(v10) {
					p268 = 1
				}
				v9 = t267 + p268<<2
				v11 = v11 + v7<<2
			}
		l7:
			if v11 != v10 {
				goto l8
			}
			if v9 != v2+i32(4) {
				goto l8
			}
			v10 = v1 << 2
			if v10 == 0 {
				goto l2
			}
			memory_copy(m.memory, uint32(v0), uint32(v4+i32(12)), uint32(v10))
			goto l2
		l8:
			m.fn987()
			panic("unreachable")
		}
		if v3 != 0 {
			goto l1
		}
		m.fn1847(v0, v1)
		goto l2
	l1:
		t287 := v0
		v10 = int32(uint32(v1) >> 3)
		v9 = t287 + v10*i32(28)
		v7 = v0 + v10<<4
		{
			{
				if uint32(v1) < uint32(i32(64)) {
					goto l10
				}
				t288 := m.fn1848(v0, v7, v9, v10)
				v10 = t288
				goto l11
			}
		l10:
			t289 := int32(load32(m.memory[uint32(v0):]))
			t290 := v0
			t291 := v9
			t292 := v7
			v10 = t289
			t293 := int32(load32(m.memory[uint32(v7):]))
			t294 := v10
			v5 = t293
			var p295 int32
			if uint32(t294) < uint32(v5) {
				p295 = 1
			}
			v11 = p295
			t296 := int32(load32(m.memory[uint32(v9):]))
			t297 := v11
			t298 := v5
			v12 = t296
			var p299 int32
			if uint32(t298) < uint32(v12) {
				p299 = 1
			}
			p300 := t292
			if t297^p299 != 0 {
				p300 = t291
			}
			t301 := v11
			var p302 int32
			if uint32(v10) < uint32(v12) {
				p302 = 1
			}
			p303 := p300
			if t301^p302 != 0 {
				p303 = t290
			}
			v10 = p303
		}
	l11:
		v3 = v3 + i32(-1)
		v10 = v10 - v0
		{
			if v2 != 0 {
				t306 := int32(load32(m.memory[uint32(v0):]))
				v7 = t306
				t307 := int32(load32(m.memory[uint32(v2):]))
				v5 = v0 + v10
				t308 := int32(load32(m.memory[uint32(v5):]))
				v9 = t308
				if uint32(t307) < uint32(v9) {
					goto l13
				}
				store32(m.memory[uint32(v0):], uint32(v9))
				store32(m.memory[uint32(v5):], uint32(v7))
				m.fn1846(v4+i32(12), v0, v1, i32(1), i32(1301028))
				{
					t309 := int32(load32(m.memory[int64(uint32(v4))+16:]))
					if t309 == 0 {
						m.fn158(i32(0), i32(0), i32(1301044))
						panic("unreachable")
					}
					{
						{
							t310 := int32(load32(m.memory[int64(uint32(v4))+24:]))
							v5 = t310
							if v5 != 0 {
								goto l15
							}
							v10 = i32(0)
							goto l16
						}
					l15:
						t311 := int32(load32(m.memory[int64(uint32(v4))+12:]))
						v7 = t311
						t312 := int32(load32(m.memory[int64(uint32(v4))+20:]))
						t313 := v4
						v9 = t312
						t314 := int32(load32(m.memory[uint32(v9):]))
						store32(m.memory[int64(uint32(t313))+140:], uint32(t314))
						store32(m.memory[int64(uint32(v4))+24:], uint32(i32(0)))
						t315 := v4
						v10 = v9 + i32(4)
						store32(m.memory[int64(uint32(t315))+20:], uint32(v10))
						store32(m.memory[int64(uint32(v4))+12:], uint32(v9))
						v5 = v9 + v5<<2
						v11 = v5 + i32(-4)
						store32(m.memory[int64(uint32(v4))+16:], uint32(v4+i32(140)))
					l20:
						if uint32(v10) < uint32(v11) {
							t318 := int32(load32(m.memory[uint32(v7):]))
							m.fn1849(t318, v9, v4+i32(12))
							t319 := int32(load32(m.memory[uint32(v7):]))
							m.fn1849(t319, v9, v4+i32(12))
							t320 := int32(load32(m.memory[int64(uint32(v4))+20:]))
							v10 = t320
							goto l20
						}
					l19:
						{
							if v10 == v5 {
								goto l18
							}
							t316 := int32(load32(m.memory[uint32(v7):]))
							m.fn1849(t316, v9, v4+i32(12))
							t317 := int32(load32(m.memory[int64(uint32(v4))+20:]))
							v10 = t317
							goto l19
						}
					l18:
						t321 := int32(load32(m.memory[int64(uint32(v4))+16:]))
						store32(m.memory[int64(uint32(v4))+20:], uint32(t321))
						t322 := int32(load32(m.memory[uint32(v7):]))
						m.fn1849(t322, v9, v4+i32(12))
						t323 := int32(load32(m.memory[int64(uint32(v4))+24:]))
						v10 = t323
					}
				l16:
					if uint32(v10) >= uint32(v1) {
						goto l21
					}
					t324 := int32(load32(m.memory[uint32(v0):]))
					v9 = t324
					t325 := v0
					v7 = v0 + v10<<2
					t326 := int32(load32(m.memory[uint32(v7):]))
					store32(m.memory[uint32(t325):], uint32(t326))
					store32(m.memory[uint32(v7):], uint32(v9))
					t327 := v1
					v10 = v10 + i32(1)
					v1 = t327 - v10
					v0 = v0 + v10<<2
					v2 = i32(0)
					goto l22
				}
			}
			t304 := int32(load32(m.memory[uint32(v0+v10):]))
			v9 = t304
			t305 := int32(load32(m.memory[uint32(v0):]))
			v7 = t305
			goto l13
		}
	l13:
		store32(m.memory[uint32(v0):], uint32(v9))
		store32(m.memory[uint32(v0+v10):], uint32(v7))
		m.fn1846(v4+i32(12), v0, v1, i32(1), i32(1301028))
		{
			t328 := int32(load32(m.memory[int64(uint32(v4))+16:]))
			if t328 == 0 {
				goto l23
			}
			{
				{
					t329 := int32(load32(m.memory[int64(uint32(v4))+24:]))
					v5 = t329
					if v5 != 0 {
						goto l24
					}
					v10 = i32(0)
					goto l25
				}
			l24:
				t330 := int32(load32(m.memory[int64(uint32(v4))+12:]))
				v7 = t330
				t331 := int32(load32(m.memory[int64(uint32(v4))+20:]))
				t332 := v4
				v9 = t331
				t333 := int32(load32(m.memory[uint32(v9):]))
				store32(m.memory[int64(uint32(t332))+140:], uint32(t333))
				store32(m.memory[int64(uint32(v4))+24:], uint32(i32(0)))
				t334 := v4
				v10 = v9 + i32(4)
				store32(m.memory[int64(uint32(t334))+20:], uint32(v10))
				store32(m.memory[int64(uint32(v4))+12:], uint32(v9))
				v5 = v9 + v5<<2
				v11 = v5 + i32(-4)
				store32(m.memory[int64(uint32(v4))+16:], uint32(v4+i32(140)))
			l29:
				if uint32(v10) < uint32(v11) {
					t337 := int32(load32(m.memory[uint32(v7):]))
					m.fn1850(t337, v9, v4+i32(12))
					t338 := int32(load32(m.memory[uint32(v7):]))
					m.fn1850(t338, v9, v4+i32(12))
					t339 := int32(load32(m.memory[int64(uint32(v4))+20:]))
					v10 = t339
					goto l29
				}
			l28:
				{
					if v10 == v5 {
						goto l27
					}
					t335 := int32(load32(m.memory[uint32(v7):]))
					m.fn1850(t335, v9, v4+i32(12))
					t336 := int32(load32(m.memory[int64(uint32(v4))+20:]))
					v10 = t336
					goto l28
				}
			l27:
				t340 := int32(load32(m.memory[int64(uint32(v4))+16:]))
				store32(m.memory[int64(uint32(v4))+20:], uint32(t340))
				t341 := int32(load32(m.memory[uint32(v7):]))
				m.fn1850(t341, v9, v4+i32(12))
				t342 := int32(load32(m.memory[int64(uint32(v4))+24:]))
				v10 = t342
			}
		l25:
			if uint32(v10) >= uint32(v1) {
				goto l21
			}
			t343 := int32(load32(m.memory[uint32(v0):]))
			v9 = t343
			t344 := v0
			v7 = v0 + v10<<2
			t345 := int32(load32(m.memory[uint32(v7):]))
			store32(m.memory[uint32(t344):], uint32(t345))
			store32(m.memory[uint32(v7):], uint32(v9))
			m.fn1846(v4+i32(12), v0, v1, v10, i32(1301060))
			t346 := int32(load32(m.memory[int64(uint32(v4))+16:]))
			v10 = t346
			t347 := int32(load32(m.memory[int64(uint32(v4))+12:]))
			v9 = t347
			t348 := int32(load32(m.memory[int64(uint32(v4))+20:]))
			t349 := int32(load32(m.memory[int64(uint32(v4))+24:]))
			m.fn1846(v4+i32(12), t348, t349, i32(1), i32(1301076))
			t350 := int32(load32(m.memory[int64(uint32(v4))+16:]))
			if t350 == 0 {
				m.fn158(i32(0), i32(0), i32(1301092))
				panic("unreachable")
			}
			t351 := int32(load32(m.memory[int64(uint32(v4))+24:]))
			v1 = t351
			t352 := int32(load32(m.memory[int64(uint32(v4))+20:]))
			v0 = t352
			t353 := int32(load32(m.memory[int64(uint32(v4))+12:]))
			v7 = t353
			m.fn1845(v9, v10, v2, v3)
			v2 = v7
			goto l22
		}
	l23:
	}
	m.fn158(i32(0), i32(0), i32(1301044))
l21:
	panic("unreachable")
l2:
	m.g0 = v4 + i32(144)
}
func (m *Module) fn1846(v0, v1, v2, v3, v4 int32) {
	if uint32(v2) >= uint32(v3) {
		goto l0
	}
	m.fn91(i32(1301172), i32(19), v4)
	panic("unreachable")
l0:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
	store32(m.memory[uint32(v0):], uint32(v1))
	store32(m.memory[int64(uint32(v0))+12:], uint32(v2-v3))
	store32(m.memory[int64(uint32(v0))+8:], uint32(v1+v3<<2))
}
func (m *Module) fn1847(v0, v1 int32) {
	var v2, v3, v4, v5, v6, v7 int32
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
			t0 := int32(load32(m.memory[uint32(v0):]))
			v4 = t0
			t1 := v0
			v3 = v0 + v2<<2
			t2 := int32(load32(m.memory[uint32(v3):]))
			store32(m.memory[uint32(t1):], uint32(t2))
			store32(m.memory[uint32(v3):], uint32(v4))
			v3 = i32(0)
		}
	l2:
		p3 := v2
		if uint32(v1) < uint32(v2) {
			p3 = v1
		}
		v5 = p3
	l5:
		{
			v6 = v3 << 1
			v4 = v6 | i32(1)
			if uint32(v4) >= uint32(v5) {
				goto l3
			}
			{
				v6 = v6 + i32(2)
				if uint32(v6) >= uint32(v5) {
					goto l4
				}
				t4 := int32(load32(m.memory[uint32(v0+v4<<2):]))
				t5 := int32(load32(m.memory[uint32(v0+v6<<2):]))
				t6 := v4
				var p7 int32
				if uint32(t4) < uint32(t5) {
					p7 = 1
				}
				v4 = t6 + p7
			}
		l4:
			v3 = v0 + v3<<2
			t8 := int32(load32(m.memory[uint32(v3):]))
			v6 = v0 + v4<<2
			t9 := int32(load32(m.memory[uint32(v6):]))
			v7 = t9
			if uint32(t8) >= uint32(v7) {
				goto l3
			}
			t10 := int32(load32(m.memory[uint32(v3):]))
			store32(m.memory[uint32(v6):], uint32(t10))
			store32(m.memory[uint32(v3):], uint32(v7))
			v3 = v4
			goto l5
		}
	}
}
func (m *Module) fn1848(v0, v1, v2, v3 int32) int32 {
	var v4, v5, v6 int32
	{
		if uint32(v3) < uint32(i32(8)) {
			goto l0
		}
		t0 := v0
		t1 := v0
		v3 = int32(uint32(v3) >> 3)
		v4 = v3 << 4
		t2 := t1 + v4
		t3 := v0
		v5 = v3 * i32(28)
		t4 := m.fn1848(t0, t2, t3+v5, v3)
		v0 = t4
		t5 := m.fn1848(v1, v1+v4, v1+v5, v3)
		v1 = t5
		t6 := m.fn1848(v2, v2+v4, v2+v5, v3)
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
func (m *Module) fn1849(v0, v1, v2 int32) {
	var v3, v4, v5 int32
	t0 := int32(load32(m.memory[int64(uint32(v2))+8:]))
	v3 = t0
	t1 := int32(load32(m.memory[uint32(v3):]))
	v4 = t1
	t2 := int32(load32(m.memory[uint32(v2):]))
	t3 := int32(load32(m.memory[int64(uint32(v2))+12:]))
	t4 := v1
	v5 = t3
	v1 = t4 + v5<<2
	t5 := int32(load32(m.memory[uint32(v1):]))
	store32(m.memory[uint32(t2):], uint32(t5))
	store32(m.memory[uint32(v2):], uint32(v3))
	store32(m.memory[int64(uint32(v2))+8:], uint32(v3+i32(4)))
	t6 := v2
	t7 := v5
	var p8 int32
	if uint32(v0) >= uint32(v4) {
		p8 = 1
	}
	store32(m.memory[int64(uint32(t6))+12:], uint32(t7+p8))
	t9 := int32(load32(m.memory[uint32(v3):]))
	store32(m.memory[uint32(v1):], uint32(t9))
}
func (m *Module) fn1850(v0, v1, v2 int32) {
	var v3, v4, v5 int32
	t0 := int32(load32(m.memory[int64(uint32(v2))+8:]))
	v3 = t0
	t1 := int32(load32(m.memory[uint32(v3):]))
	v4 = t1
	t2 := int32(load32(m.memory[uint32(v2):]))
	t3 := int32(load32(m.memory[int64(uint32(v2))+12:]))
	t4 := v1
	v5 = t3
	v1 = t4 + v5<<2
	t5 := int32(load32(m.memory[uint32(v1):]))
	store32(m.memory[uint32(t2):], uint32(t5))
	store32(m.memory[uint32(v2):], uint32(v3))
	store32(m.memory[int64(uint32(v2))+8:], uint32(v3+i32(4)))
	t6 := v2
	t7 := v5
	var p8 int32
	if uint32(v4) < uint32(v0) {
		p8 = 1
	}
	store32(m.memory[int64(uint32(t6))+12:], uint32(t7+p8))
	t9 := int32(load32(m.memory[uint32(v3):]))
	store32(m.memory[uint32(v1):], uint32(t9))
}
func (m *Module) fn1851(v0, v1, v2 int32) int32 {
	var v3, v4, v5 int32
	v3 = i32(0)
	if v2 == 0 {
		goto l0
	}
l2:
	{
		t0 := int32(m.memory[uint32(v0)])
		v4 = t0
		t1 := int32(m.memory[uint32(v1)])
		t2 := v4
		v5 = t1
		if t2 != v5 {
			goto l1
		}
		v0 = v0 + i32(1)
		v1 = v1 + i32(1)
		v2 = v2 + i32(-1)
		if v2 == 0 {
			goto l0
		}
		goto l2
	}
l1:
	v3 = v4 - v5
l0:
	return v3
}
func (m *Module) fn1852(v0 int32) int32 {
	var v1, v2 int32
	{
		t0 := int32(m.memory[uint32(v0)])
		if t0 != 0 {
			goto l0
		}
		v0 = i32(0)
		goto l1
	}
l0:
	v1 = v0 + i32(1)
	v0 = i32(0)
l2:
	{
		v2 = v0
		v0 = v2 + i32(1)
		t1 := int32(m.memory[uint32(v1+v2)])
		if t1 != 0 {
			goto l2
		}
	}
l1:
	return v0
}
func (m *Module) fn1853(v0 int32, v1, v2, v3, v4 int64) {
	var v5, v6, v7, v8, v9, v10 int64
	t0 := v0
	v5 = v3 & i64(0xffffffff)
	t1 := v5
	v6 = v1 & i64(0xffffffff)
	v7 = t1 * v6
	t2 := v7
	v8 = int64(uint64(v3) >> 32)
	v6 = v8 * v6
	t3 := v6
	t4 := v5
	v9 = int64(uint64(v1) >> 32)
	v5 = t3 + t4*v9
	v10 = t2 + v5<<32
	store64(m.memory[uint32(t0):], uint64(v10))
	t5 := v0
	t6 := v8 * v9
	var p7 int32
	if uint64(v5) < uint64(v6) {
		p7 = 1
	}
	t8 := t6 + (int64(uint32(p7))<<32 | int64(uint64(v5)>>32))
	var p9 int32
	if uint64(v10) < uint64(v7) {
		p9 = 1
	}
	store64(m.memory[int64(uint32(t5))+8:], uint64(t8+int64(uint32(p9))+(v4*v1+v3*v2)))
}
func fn1854(v0 float64) float64 {
	var v1 int64
	var v2 int32
	{
		v0 = float64(v0 + math.Copysign(float64(0.49999999999999994), v0))
		v1 = int64(math.Float64bits(v0))
		v2 = int32(int64(uint64(v1)>>52)) & i32(2047)
		if uint32(v2) > uint32(i32(1074)) {
			goto l0
		}
		p0 := i64_shr_s(i64(-0x10000000000000), int64(uint32(v2+i32(-1023))))
		if uint32(v2) < uint32(i32(1023)) {
			p0 = i64(-0x8000000000000000)
		}
		v0 = math.Float64frombits(uint64(p0 & v1))
	}
l0:
	return v0
}
func (m *Module) Xmemory() Memory {
	return (*wasmMemory)(&m.memory)
}

//go:nosplit
func i32(x int32) int32 { return x }

//go:nosplit
func i64(x int64) int64 { return x }

//go:nosplit
func i32_div_s(x, y int32) int32 {
	if y == -1 && x == math.MinInt32 {
		panic("integer overflow")
	}
	return x / y
}

//go:nosplit
func i32_shl(x, y int32) int32 {
	return x << (y & 31)
}

//go:nosplit
func i32_shr_u(x, y int32) int32 {
	return int32(uint32(x) >> (y & 31))
}

//go:nosplit
func i64_shl(x, y int64) int64 {
	return x << (y & 63)
}

//go:nosplit
func i64_shr_s(x, y int64) int64 {
	return x >> (y & 63)
}

//go:nosplit
func i64_shr_u(x, y int64) int64 {
	return int64(uint64(x) >> (y & 63))
}

//go:nosplit
func i32_rotl(x, y int32) int32 {
	return int32(bits.RotateLeft32(uint32(x), int(y)))
}

//go:nosplit
func i32_rotr(x, y int32) int32 {
	return int32(bits.RotateLeft32(uint32(x), -int(y)))
}

//go:nosplit
func i64_rotl(x, y int64) int64 {
	return int64(bits.RotateLeft64(uint64(x), int(y)))
}

//go:nosplit
func i64_trunc_sat_f64_s(f float64) int64 {
	switch {
	case f < math.MinInt64:
		return math.MinInt64
	case f >= math.MaxInt64:
		return math.MaxInt64
	case f != f:
		return 0
	}
	return int64(f)
}

//go:nosplit
func i64_trunc_sat_f64_u(f float64) int64 {
	var i uint64
	switch {
	case f <= 0 || f != f:
		i = 0
	case f >= math.MaxUint64:
		i = math.MaxUint64
	default:
		i = uint64(f)
	}
	return int64(i)
}

//go:nosplit
func load16(b []byte) uint16 {
	return binary.LittleEndian.Uint16(b)
}

//go:nosplit
func store16(b []byte, v uint16) {
	binary.LittleEndian.PutUint16(b, v)
}

//go:nosplit
func load32(b []byte) uint32 {
	return binary.LittleEndian.Uint32(b)
}

//go:nosplit
func store32(b []byte, v uint32) {
	binary.LittleEndian.PutUint32(b, v)
}

//go:nosplit
func load64(b []byte) uint64 {
	return binary.LittleEndian.Uint64(b)
}

//go:nosplit
func store64(b []byte, v uint64) {
	binary.LittleEndian.PutUint64(b, v)
}

func memory_grow(mem *[]byte, delta, max int64) int64 {
	buf := *mem
	len := int64(len(buf))
	old := len >> 16
	if delta == 0 {
		return old
	}
	new := old + delta
	add := new<<16 - len
	max = min(max, int64(math.MaxInt)>>16)
	if new > max || new < old || add < 0 {
		return -1
	}
	*mem = append(buf, make([]byte, add)...)
	return old
}

func memory_init[T1, T2 int | uint32 | uint64](mem []byte, data string, dest T1, src, n T2) {
	x := uint64(dest)
	z := uint64(src)
	y := x + uint64(n)
	w := z + uint64(n)
	copy(mem[x:y], data[z:w])
}

func memory_copy[T uint32 | uint64](mem []byte, dest, src, n T) {
	x := uint64(dest)
	z := uint64(src)
	y := x + uint64(n)
	w := z + uint64(n)
	copy(mem[x:y], mem[z:w])
}

func memory_fill[T uint32 | uint64](mem []byte, dest T, val int32, n T) {
	x := uint64(dest)
	y := x + uint64(n)
	buf := mem[x:y]
	if len(buf) > 0 {
		buf[0] = byte(val)
		for i := 1; i < len(buf); {
			chunk := min(i, 8192)
			i += copy(buf[i:], buf[:chunk])
		}
	}
}

func memory_zero[T uint32 | uint64](mem []byte, dest, n T) {
	x := uint64(dest)
	y := x + uint64(n)
	clear(mem[x:y])
}

func table_init[T1, T2, T3 int | int32 | int64](tab, elems []any, dest T1, src T2, n T3) {
	x := uint64(dest)
	z := uint64(src)
	y := x + uint64(n)
	w := z + uint64(n)
	copy(tab[x:y], elems[z:w])
}
//go:embed anydoc.wasm.dat
var data string

