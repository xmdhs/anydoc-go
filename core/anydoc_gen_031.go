package core

import (
	"math/bits"
)

func (m *Module) fn1347(v0, v1, v2, v3, v4, v5, v6, v7 int32) {
	var v8, v9, v10 int32
	t0 := m.g0
	v8 = t0 - i32(48)
	m.g0 = v8
	{
		{
			t1 := int32(load32(m.memory[uint32(v2):]))
			if t1 != i32(1) {
				store64(m.memory[uint32(v0):], uint64(i64(0xffffffff)))
				goto l6
			}
			t2 := int32(load32(m.memory[int64(uint32(v2))+8:]))
			t3 := v7 ^ i32(-1)
			v9 = t2
			var p4 int32
			if v9 != i32(0) {
				p4 = 1
			}
			v10 = t3 & p4
			t5 := int32(load32(m.memory[int64(uint32(v2))+4:]))
			t6 := m.fn1286(v3, t5)
			v2 = t6
			if v2 == 0 {
				goto l1
			}
			t7 := int32(load32(m.memory[uint32(v2):]))
			m.fn1285(v8, v4, v5, t7)
			t8 := int32(load32(m.memory[int64(uint32(v8))+4:]))
			v3 = t8
			if v3 == 0 {
				goto l1
			}
			t9 := int32(load16(m.memory[int64(uint32(v8))+2:]))
			if t9 != v6&i32(0xffff) {
				goto l1
			}
			t10 := int32(load32(m.memory[int64(uint32(v8))+8:]))
			v5 = t10
			if v7 == 0 {
				goto l2
			}
			store32(m.memory[int64(uint32(v8))+20:], uint32(i32(0)))
			store32(m.memory[int64(uint32(v8))+16:], uint32(v5))
			store32(m.memory[int64(uint32(v8))+12:], uint32(v3))
			{
			l4:
				{
					m.fn1287(v8+i32(24), v8+i32(12))
					t11 := int32(load32(m.memory[int64(uint32(v8))+28:]))
					v2 = t11
					if v2 == 0 {
						goto l3
					}
					t12 := int32(load16(m.memory[int64(uint32(v8))+26:]))
					if t12<<16 != i32(0x3f10000) {
						goto l4
					}
				}
				t13 := int32(load32(m.memory[int64(uint32(v8))+32:]))
				if uint32(t13) >= uint32(i32(4)) {
					t14 := int32(load32(m.memory[uint32(v2):]))
					v9 = t14
					var p15 int32
					if v9 != i32(0) {
						p15 = 1
					}
					v10 = p15
					goto l2
				}
			}
		l3:
			v10 = i32(0)
			goto l2
		}
	l2:
		m.fn1291(v8+i32(24), v1, v3, v5)
		t16 := int32(load32(m.memory[int64(uint32(v8))+24:]))
		v2 = t16
		if v2 == i32(-1) {
			goto l1
		}
		t17 := int32(load32(m.memory[int64(uint32(v8))+44:]))
		store32(m.memory[int64(uint32(v0))+20:], uint32(t17))
		t18 := int64(load64(m.memory[int64(uint32(v8))+36:]))
		store64(m.memory[int64(uint32(v0))+12:], uint64(t18))
		t19 := int64(load64(m.memory[int64(uint32(v8))+28:]))
		store64(m.memory[int64(uint32(v0))+4:], uint64(t19))
		store32(m.memory[uint32(v0):], uint32(v2))
		goto l6
	}
l1:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v9))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v10))
	store32(m.memory[uint32(v0):], uint32(i32(-1)))
l6:
	m.g0 = v8 + i32(48)
}
func (m *Module) fn1348(v0, v1 int32) {
	var v2 int32
	var v3 int64
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	t1 := int64(load64(m.memory[int64(uint32(v1))+8:]))
	t2 := v1
	v3 = t1 + i64(1)
	store64(m.memory[int64(uint32(t2))+8:], uint64(v3))
	v1 = i32(-1)
	if uint64(v3) < uint64(i64(16000001)) {
		goto l0
	}
	store32(m.memory[int64(uint32(v2))+12:], uint32(i32(28)))
	store32(m.memory[int64(uint32(v2))+8:], uint32(i32(1075168)))
	m.fn73(v0+i32(4), i32(1067030), v2+i32(8))
	store32(m.memory[int64(uint32(v0))+20:], uint32(i32(11)))
	store32(m.memory[int64(uint32(v0))+16:], uint32(i32(1075176)))
	v1 = i32(-0x7ffffffd)
l0:
	store32(m.memory[uint32(v0):], uint32(v1))
	m.g0 = v2 + i32(16)
}
func (m *Module) fn1349(v0, v1, v2, v3 int32) {
	var v4, v5, v6, v7, v8, v9, v10, v11, v12 int32
	var v13, v14 int64
	t0 := m.g0
	v4 = t0 - i32(96)
	m.g0 = v4
	switch v1&i32(0xffff) + i32(-3999) {
	default:
		goto l3
	case 0:
		m.fn1350(v0)
		{
			if v3 != 0 {
				goto l5
			}
			v2 = i32(1)
			goto l6
		l5:
			t1 := int32(m.memory[uint32(v2)])
			v2 = t1
		}
	l6:
		m.fn1303(v0 + i32(16))
		m.memory[int64(uint32(v0))+52] = byte(v2)
		store64(m.memory[int64(uint32(v0))+24:], uint64(i64(-0x100000000)))
		store64(m.memory[int64(uint32(v0))+16:], uint64(i64(0x100000000)))
		goto l3
	case 1:
		store32(m.memory[int64(uint32(v4))+56:], uint32(i32(0)))
		store64(m.memory[int64(uint32(v4))+48:], uint64(i64(0x100000000)))
		t2 := v4 + i32(60)
		v1 = v3 & i32(0x7ffffffe)
		m.fn492(t2, v1, i32(2))
		t3 := int32(load32(m.memory[int64(uint32(v4))+60:]))
		t4 := v4 + i32(48)
		v5 = t3
		m.fn47(t4, v5-int32(uint32(v5)>>1))
		store32(m.memory[int64(uint32(v4))+80:], uint32(i32(2)))
		store32(m.memory[int64(uint32(v4))+76:], uint32(v3&i32(1)))
		store32(m.memory[int64(uint32(v4))+68:], uint32(v1))
		store32(m.memory[int64(uint32(v4))+64:], uint32(v2))
		store32(m.memory[int64(uint32(v4))+72:], uint32(v2+v1))
		v1 = v4 + i32(64)
		v3 = i32(0)
	l12:
		{
			store16(m.memory[int64(uint32(v4))+60:], uint16(i32(0)))
			{
				if v3&i32(1) != 0 {
					goto l7
				}
				m.fn1022(v4+i32(8), v1)
				t5 := int32(load16(m.memory[int64(uint32(v4))+8:]))
				if t5&i32(1) == 0 {
					m.fn1502(v0, v4+i32(48))
					goto l3
				}
				t6 := int32(load16(m.memory[int64(uint32(v4))+10:]))
				v2 = t6
			}
		l7:
			{
				if v2&i32(63488) == i32(55296) {
					goto l9
				}
				v3 = v2 & i32(0xffff)
				goto l10
			l9:
				v3 = i32(65533)
				if uint32(v2&i32(0xffff)) > uint32(i32(56319)) {
					goto l10
				}
				m.fn1022(v4, v1)
				t7 := int32(load16(m.memory[uint32(v4):]))
				if t7&i32(1) == 0 {
					goto l10
				}
				{
					t8 := int32(load16(m.memory[int64(uint32(v4))+2:]))
					v5 = t8
					if uint32((v5+i32(8192))&i32(0xffff)) > uint32(i32(64511)) {
						goto l11
					}
					store16(m.memory[int64(uint32(v4))+62:], uint16(v5))
					store16(m.memory[int64(uint32(v4))+60:], uint16(i32(1)))
					goto l10
				}
			l11:
				v3 = v2&i32(1023)<<10 | v5&i32(1023) + i32(65536)
			}
		l10:
			m.fn74(v4+i32(48), v3)
			t9 := int32(load16(m.memory[int64(uint32(v4))+60:]))
			v3 = t9
			t10 := int32(load16(m.memory[int64(uint32(v4))+62:]))
			v2 = t10
			goto l12
		}
	case 9:
		store32(m.memory[int64(uint32(v4))+92:], uint32(i32(0)))
		store64(m.memory[int64(uint32(v4))+84:], uint64(i64(0x100000000)))
		m.fn47(v4+i32(84), v3)
		if v3 == 0 {
			goto l13
		}
	l14:
		{
			t11 := int32(m.memory[uint32(v2)])
			m.fn74(v4+i32(84), t11)
			v2 = v2 + i32(1)
			v3 = v3 + i32(-1)
			if v3 != 0 {
				goto l14
			}
		}
	l13:
		m.fn1502(v0, v4+i32(84))
		goto l3
	case 2:
		t12 := int32(load32(m.memory[int64(uint32(v0))+16:]))
		if t12 == i32(-1) {
			goto l3
		}
		t13 := int32(load32(m.memory[int64(uint32(v0))+20:]))
		t14 := v4
		v1 = t13
		t15 := int32(load32(m.memory[int64(uint32(v0))+24:]))
		store32(m.memory[int64(uint32(t14))+64:], uint32(v1+t15))
		store32(m.memory[int64(uint32(v4))+60:], uint32(v1))
		v5 = i32(0)
	l16:
		{
			t16 := m.fn48(v4 + i32(60))
			v1 = t16
			if v1 == i32(-1) {
				goto l15
			}
			p17 := i32(2)
			if uint32(v1) < uint32(i32(65536)) {
				p17 = i32(1)
			}
			v5 = p17 + v5
			goto l16
		}
	l15:
		store64(m.memory[int64(uint32(v4))+40:], uint64(i64(4)))
		store64(m.memory[int64(uint32(v4))+32:], uint64(i64(0)))
		store64(m.memory[int64(uint32(v4))+24:], uint64(i64(0x400000000)))
		v6 = v4 + i32(36)
		v7 = i32(0)
		v8 = i32(1)
		v9 = i32(4)
		v1 = i32(0)
		v10 = i32(0)
	l20:
		{
			if uint32(v3) < uint32(v1) {
				goto l17
			}
			if uint32(v3-v1) < uint32(i32(4)) {
				goto l17
			}
			t18 := v3
			v11 = v1 + i32(4)
			if uint32(t18-v11) <= uint32(i32(1)) {
				goto l17
			}
			t19 := int32(load32(m.memory[uint32(v2+v1):]))
			v12 = t19
			t20 := int64(load16(m.memory[uint32(v2+v11):]))
			v13 = t20
			m.fn1345(v4+i32(16), v2, v3, v1+i32(6))
			t21 := int32(m.memory[int64(uint32(v4))+16])
			v11 = t21
			if v11 == i32(255) {
				goto l18
			}
			t22 := int32(load32(m.memory[int64(uint32(v4))+20:]))
			v1 = t22
			{
				t23 := int32(load32(m.memory[int64(uint32(v4))+24:]))
				if v8+i32(-1) != t23 {
					goto l19
				}
				m.fn625(v4 + i32(24))
				t24 := int32(load32(m.memory[int64(uint32(v4))+28:]))
				v9 = t24
			}
		l19:
			store64(m.memory[uint32(v9+v7):], uint64(int64(uint32(v11))<<48|v13<<32|int64(uint32(v12))))
			store32(m.memory[int64(uint32(v4))+32:], uint32(v8))
			if v12 == 0 {
				goto l17
			}
			v7 = v7 + i32(8)
			v8 = v8 + i32(1)
			v10 = v12 + v10
			if uint32(v10) <= uint32(v5) {
				goto l20
			}
		}
	l17:
		v7 = i32(0)
	l22:
		{
			if uint32(v3) < uint32(v1) {
				goto l18
			}
			if uint32(v3-v1) < uint32(i32(4)) {
				goto l18
			}
			t25 := int32(load32(m.memory[uint32(v2+v1):]))
			v8 = t25
			m.fn1346(v4+i32(60), v2, v3, v1+i32(4))
			t26 := int64(m.memory[int64(uint32(v4))+60])
			v13 = t26
			if v13 == i64(255) {
				goto l18
			}
			t27 := int32(load32(m.memory[int64(uint32(v4))+64:]))
			v1 = t27
			t28 := int64(m.memory[int64(uint32(v4))+61])
			v14 = t28
			{
				t29 := int32(load32(m.memory[int64(uint32(v4))+44:]))
				v12 = t29
				t30 := int32(load32(m.memory[int64(uint32(v4))+36:]))
				if v12 != t30 {
					goto l21
				}
				m.fn625(v6)
			}
		l21:
			t31 := int32(load32(m.memory[int64(uint32(v4))+40:]))
			store64(m.memory[uint32(t31+v12<<3):], uint64(v13<<32|int64(uint32(v8))|v14<<40))
			store32(m.memory[int64(uint32(v4))+44:], uint32(v12+i32(1)))
			if v8 == 0 {
				goto l18
			}
			v7 = v8 + v7
			if uint32(v7) <= uint32(v5) {
				goto l22
			}
		}
	l18:
		m.fn1498(v0 + i32(28))
		t32 := int64(load64(m.memory[int64(uint32(v4))+40:]))
		store64(m.memory[int64(uint32(v0))+44:], uint64(t32))
		t33 := int64(load64(m.memory[int64(uint32(v4))+32:]))
		store64(m.memory[int64(uint32(v0))+36:], uint64(t33))
		t34 := int64(load64(m.memory[int64(uint32(v4))+24:]))
		store64(m.memory[int64(uint32(v0))+28:], uint64(t34))
	}
l3:
	m.g0 = v4 + i32(96)
}
func (m *Module) fn1350(v0 int32) {
	var v1, v2, v3 int32
	var v4 int64
	var v5, v6, v7 int32
	var v8 int64
	var v9 int32
	var v10 int64
	var v11 int32
	var v12 int64
	var v13, v14, v15, v16, v17, v18, v19, v20, v21, v22, v23, v24, v25, v26 int32
	t0 := m.g0
	v1 = t0 - i32(208)
	m.g0 = v1
	t1 := int32(load32(m.memory[int64(uint32(v0))+16:]))
	v2 = t1
	store32(m.memory[int64(uint32(v0))+16:], uint32(i32(-1)))
	if v2 == i32(-1) {
		goto l0
	}
	store32(m.memory[int64(uint32(v1))+12:], uint32(v2))
	memory_copy(m.memory, uint32(v1+i32(12)+i32(4)), uint32(v0+i32(20)), uint32(i32(36)))
	{
		t2 := int32(load32(m.memory[int64(uint32(v1))+20:]))
		v3 = t2
		if v3 == 0 {
			m.fn1355(v1 + i32(12))
			goto l0
		}
		t3 := int64(load64(m.memory[uint32(v0):]))
		t4 := v0
		v4 = t3 + i64(1)
		store64(m.memory[uint32(t4):], uint64(v4))
		{
			t5 := int32(load32(m.memory[int64(uint32(v1))+24:]))
			v5 = t5
			if v5 == i32(-1) {
				goto l2
			}
			v6 = v1 + i32(52) + i32(4)
			t6 := int32(load32(m.memory[int64(uint32(v0))+48:]))
			store32(m.memory[int64(uint32(v6))+16:], uint32(t6))
			t7 := int64(load64(m.memory[int64(uint32(v0))+40:]))
			store64(m.memory[int64(uint32(v6))+8:], uint64(t7))
			t8 := int64(load64(m.memory[int64(uint32(v0))+32:]))
			store64(m.memory[uint32(v6):], uint64(t8))
			goto l3
		}
	l2:
		v5 = i32(0)
		store32(m.memory[int64(uint32(v1))+72:], uint32(i32(0)))
		store64(m.memory[int64(uint32(v1))+64:], uint64(i64(0x400000000)))
		store64(m.memory[int64(uint32(v1))+56:], uint64(i64(4)))
	l3:
		t9 := int32(m.memory[int64(uint32(v1))+48])
		v7 = t9
		store32(m.memory[int64(uint32(v1))+52:], uint32(v5))
		{
			t10 := int32(load32(m.memory[int64(uint32(v0))+56:]))
			v5 = t10
			t11 := int32(load32(m.memory[int64(uint32(v0))+104:]))
			if uint32(v5) >= uint32(t11) {
				goto l4
			}
			t12 := int32(load32(m.memory[int64(uint32(v0))+100:]))
			v5 = t12 + v5*i32(40)
			t13 := int32(load32(m.memory[int64(uint32(v5))+20:]))
			if t13 == 0 {
				goto l4
			}
			t14 := int64(load64(m.memory[int64(uint32(v5))+24:]))
			t15 := int64(load64(m.memory[uint32(v5+i32(32)):]))
			t16 := m.fn529(t14, t15, v7)
			v8 = t16
			t17 := int32(load32(m.memory[int64(uint32(v5))+12:]))
			v9 = t17
			v6 = v9 & int32(v8)
			v10 = int64(uint64(v8)>>25) & i64(127) * i64(72340172838076673)
			t18 := int32(load32(m.memory[int64(uint32(v5))+8:]))
			v5 = t18
			v11 = i32(0)
		l8:
			{
				t19 := int64(load64(m.memory[uint32(v5+v6):]))
				v12 = t19
				v8 = v12 ^ v10
				v8 = (v8 ^ i64(-1)) & (v8 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
			l7:
				{
					if v8 == 0 {
						if !(v12&(v12<<1)&i64(-0x7f7f7f7f7f7f7f80) == 0) {
							goto l4
						}
						t22 := v6
						v11 = v11 + i32(8)
						v6 = (t22 + v11) & v9
						goto l8
					}
					t20 := v5
					v13 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v8))))>>3) + v6) & v9
					t21 := int32(load16(m.memory[uint32(t20-v13<<4+i32(-16)):]))
					if t21 == v7 {
						goto l6
					}
					v8 = (v8 + i64(-1)) & v8
					goto l7
				}
			l6:
			}
			v5 = v5 - v13<<4
			t23 := int32(load32(m.memory[uint32(v5+i32(-8)):]))
			v6 = t23
			t24 := int32(load32(m.memory[uint32(v5+i32(-4)):]))
			t25 := v1
			v14 = t24
			m.fn59(t25, v14, i32(1), i32(3))
			t26 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			v15 = t26
			t27 := int32(load32(m.memory[uint32(v1):]))
			v16 = t27
			if v14 != 0 {
				goto l9
			}
			v14 = i32(0)
			goto l10
		l9:
			v5 = v14 * i32(3)
			if v5 == 0 {
				goto l10
			}
			memory_copy(m.memory, uint32(v15), uint32(v6), uint32(v5))
		l10:
			if v16 != i32(-1) {
				goto l11
			}
		}
	l4:
		v15 = i32(1)
		v14 = i32(0)
		v16 = i32(0)
	l11:
		v9 = i32(0)
		v13 = i32(-1)
		v11 = i32(-1)
		v17 = i32(0)
		t28 := int32(load32(m.memory[int64(uint32(v1))+68:]))
		v5 = t28
		v18 = v5
		{
			t29 := int32(load32(m.memory[int64(uint32(v1))+72:]))
			v6 = t29
			if v6 == 0 {
				goto l12
			}
			v18 = v5 + i32(8)
			t30 := int32(load32(m.memory[uint32(v5):]))
			v11 = t30
			v17 = v5
		}
	l12:
		v19 = v6 << 3
		t31 := int32(load32(m.memory[int64(uint32(v1))+60:]))
		v20 = t31
		v21 = v20 << 3
		t32 := int32(load32(m.memory[int64(uint32(v1))+56:]))
		v6 = t32
		v22 = v6
		{
			if v20 == 0 {
				goto l13
			}
			v22 = v6 + i32(8)
			t33 := int32(load32(m.memory[uint32(v6):]))
			v13 = t33
			v9 = v6
		}
	l13:
		v23 = v5 + v19
		v21 = v6 + v21
		store32(m.memory[int64(uint32(v1))+84:], uint32(i32(0)))
		store64(m.memory[int64(uint32(v1))+76:], uint64(i64(0x400000000)))
		store32(m.memory[int64(uint32(v1))+96:], uint32(i32(0)))
		store64(m.memory[int64(uint32(v1))+88:], uint64(i64(0x400000000)))
		store32(m.memory[int64(uint32(v1))+108:], uint32(i32(0)))
		store64(m.memory[int64(uint32(v1))+100:], uint64(i64(0x100000000)))
		store32(m.memory[int64(uint32(v1))+112:], uint32(i32(0)))
		t34 := int32(load32(m.memory[int64(uint32(v1))+16:]))
		t35 := v1
		v24 = t34
		store32(m.memory[int64(uint32(t35))+120:], uint32(v24+v3))
		store32(m.memory[int64(uint32(v1))+116:], uint32(v24))
		v20 = v1 + i32(156)
		v25 = i32(0)
	l54:
		{
			{
				t36 := m.fn48(v1 + i32(116))
				v6 = t36
				if v6 == i32(-1) {
					{
						t45 := int32(load32(m.memory[int64(uint32(v1))+108:]))
						v5 = t45
						if v5 == 0 {
							goto l23
						}
						t46 := int32(load32(m.memory[int64(uint32(v1))+104:]))
						m.fn865(v1+i32(128), t46, v5)
						{
							t47 := int32(load32(m.memory[int64(uint32(v1))+136:]))
							if t47 == 0 {
								goto l24
							}
							t48 := int32(load32(m.memory[int64(uint32(v1))+136:]))
							store32(m.memory[int64(uint32(v1))+164:], uint32(t48))
							t49 := int64(load64(m.memory[int64(uint32(v1))+128:]))
							store64(m.memory[int64(uint32(v1))+156:], uint64(t49))
							store32(m.memory[int64(uint32(v1))+168:], uint32(v25))
							store32(m.memory[int64(uint32(v1))+152:], uint32(i32(3)))
							m.fn1340(v1+i32(88), v1+i32(152))
							goto l23
						}
					l24:
						t50 := int32(load32(m.memory[int64(uint32(v1))+128:]))
						t51 := int32(load32(m.memory[int64(uint32(v1))+132:]))
						m.fn16(t50, t51)
					}
				l23:
					{
						t52 := int32(load32(m.memory[int64(uint32(v1))+96:]))
						v26 = t52
						if v26 == 0 {
							goto l25
						}
						{
							if v9 != 0 {
								goto l26
							}
							v5 = i32(2)
							v6 = i32(0)
							goto l27
						l26:
							t53 := int32(m.memory[int64(uint32(v9))+6])
							v5 = t53
							t54 := int32(load16(m.memory[int64(uint32(v9))+4:]))
							v6 = t54
						}
					l27:
						t55 := int32(load32(m.memory[int64(uint32(v1))+96:]))
						store32(m.memory[int64(uint32(v1))+160:], uint32(t55))
						t56 := int64(load64(m.memory[int64(uint32(v1))+88:]))
						store64(m.memory[int64(uint32(v1))+152:], uint64(t56))
						m.memory[int64(uint32(v1))+166] = byte(v5)
						store16(m.memory[int64(uint32(v1))+164:], uint16(v6))
						m.fn1501(v1+i32(76), v1+i32(152))
					}
				l25:
					v3 = v0 + i32(84)
					v18 = v0 + i32(72)
					t57 := int32(load32(m.memory[int64(uint32(v1))+80:]))
					v21 = t57
					t58 := int32(load32(m.memory[int64(uint32(v1))+84:]))
					t59 := v21
					v17 = t58 << 4
					v25 = t59 + v17
					v23 = v1 + i32(152) | i32(4)
					v19 = v1 + i32(164)
					t60 := int32(load32(m.memory[int64(uint32(v1))+76:]))
					v0 = t60
					var p61 int32
					if v7 == i32(6) {
						p61 = 1
					}
					v20 = p61
					v6 = i32(0)
					v5 = v21
				l38:
					{
						{
							if v17 != v6 {
								goto l28
							}
							v5 = v25
							goto l29
						l28:
							{
								t62 := int32(load32(m.memory[uint32(v5):]))
								v9 = t62
								if v9 == i32(-1) {
									goto l30
								}
								t63 := int32(m.memory[int64(uint32(v5))+14])
								v22 = t63
								t64 := int32(load16(m.memory[int64(uint32(v5))+12:]))
								v11 = t64
								store32(m.memory[int64(uint32(v1))+140:], uint32(v9))
								t65 := int64(load64(m.memory[int64(uint32(v5))+4:]))
								t66 := v1
								v8 = t65
								store64(m.memory[int64(uint32(t66))+144:], uint64(v8))
								v9 = int32(v8)
								t67 := v9
								v13 = int32(int64(uint64(v8) >> 32))
								t68 := m.fn23(t67, v13)
								if t68 != 0 {
									m.fn1351(v18, v3)
									m.fn894(v1 + i32(140))
									goto l34
								}
								if v7 == 0 {
									goto l32
								}
								if v20 != 0 {
									goto l32
								}
								t69 := m.fn1500(v15, v14, v11)
								p70 := v22
								if v22&i32(255) == i32(2) {
									p70 = t69
								}
								if p70&i32(1) != 0 {
									goto l33
								}
								m.fn1351(v18, v3)
								t71 := int32(load32(m.memory[int64(uint32(v1))+148:]))
								store32(m.memory[int64(uint32(v23))+8:], uint32(t71))
								t72 := int64(load64(m.memory[int64(uint32(v1))+140:]))
								store64(m.memory[uint32(v23):], uint64(t72))
								store32(m.memory[int64(uint32(v1))+152:], uint32(i32(-0x80000000)))
								m.fn338(v18, v1+i32(152))
								goto l34
							}
						l30:
							v5 = v21 + v6 + i32(16)
						l29:
							v6 = int32(uint32(v25-v5) >> 4)
						l36:
							if v6 == 0 {
								m.fn136(v0, v21, i32(4), i32(16))
								t73 := int32(load32(m.memory[int64(uint32(v1))+100:]))
								t74 := int32(load32(m.memory[int64(uint32(v1))+104:]))
								m.fn16(t73, t74)
								if v26 != 0 {
									goto l37
								}
								m.fn894(v1 + i32(88))
								goto l37
							}
							v6 = v6 + i32(-1)
							m.fn894(v5)
							v5 = v5 + i32(16)
							goto l36
						l32:
							m.fn1351(v18, v3)
							t75 := m.fn1500(v15, v14, v11)
							t76 := m.fn1188(int32(uint32(t75&i32(0xffff00))>>8) | i32(33685504))
							m.fn1368(v9, v13, t76)
							m.fn45(v19, v9, v13)
							m.memory[int64(uint32(v1))+176] = byte(i32(2))
							t77 := int32(load32(m.memory[int64(uint32(v1))+148:]))
							store32(m.memory[int64(uint32(v1))+160:], uint32(t77))
							t78 := int64(load64(m.memory[int64(uint32(v1))+140:]))
							store64(m.memory[int64(uint32(v1))+152:], uint64(t78))
							m.fn338(v18, v1+i32(152))
							goto l34
						}
					l33:
						t79 := m.fn113(i32(8), i32(32))
						v9 = t79
						store32(m.memory[uint32(v9):], uint32(i32(-0x80000000)))
						t80 := int64(load64(m.memory[int64(uint32(v1))+140:]))
						store64(m.memory[int64(uint32(v9))+4:], uint64(t80))
						t81 := int32(load32(m.memory[int64(uint32(v1))+148:]))
						store32(m.memory[int64(uint32(v9))+12:], uint32(t81))
						store32(m.memory[int64(uint32(v1))+176:], uint32(v11))
						store32(m.memory[int64(uint32(v1))+200:], uint32(i32(1)))
						store32(m.memory[int64(uint32(v1))+196:], uint32(v9))
						store32(m.memory[int64(uint32(v1))+192:], uint32(i32(1)))
						store32(m.memory[int64(uint32(v1))+180:], uint32(i32(-1)))
						store64(m.memory[int64(uint32(v1))+168:], uint64(i64(0)))
						m.memory[int64(uint32(v1))+160] = byte(i32(0))
						store64(m.memory[int64(uint32(v1))+152:], uint64(v4))
						m.fn1369(v3, v1+i32(152))
					}
				l34:
					v5 = v5 + i32(16)
					v6 = v6 + i32(16)
					goto l38
				}
				{
					if v9 != 0 {
						goto l15
					}
					v5 = i32(0)
					goto l16
				l15:
					t37 := int32(load16(m.memory[int64(uint32(v9))+4:]))
					v5 = t37
				}
			l16:
				t38 := m.fn1500(v15, v14, v5)
				v3 = t38 & i32(0xffffff)
				v5 = int32(uint32(v3) >> 16)
				v3 = int32(uint32(v3) >> 8)
				{
					{
						if v17 != 0 {
							goto l17
						}
						var p39 int32
						if v3&i32(255) != i32(2) {
							p39 = 1
						}
						v19 = p39 & v3
						goto l18
					}
				l17:
					{
						t40 := int32(m.memory[int64(uint32(v17))+4])
						v19 = t40
						if v19 != i32(2) {
							goto l19
						}
						var p41 int32
						if v3&i32(255) != i32(2) {
							p41 = 1
						}
						v19 = p41 & v3
					}
				l19:
					t42 := int32(m.memory[int64(uint32(v17))+5])
					t43 := v5
					v3 = t42
					p44 := v3
					if v3 == i32(2) {
						p44 = t43
					}
					v5 = p44
				}
			l18:
				store16(m.memory[int64(uint32(v1))+126:], uint16(i32(0)))
				m.memory[int64(uint32(v1))+125] = byte(v5 & i32(1))
				m.memory[int64(uint32(v1))+124] = byte(v19 & i32(1))
				switch v6 + i32(-11) {
				case 0:
					{
						t94 := int32(load32(m.memory[int64(uint32(v1))+108:]))
						v3 = t94
						if v3 == 0 {
							goto l44
						}
						t95 := int32(load32(m.memory[int64(uint32(v1))+100:]))
						v19 = t95
						t96 := int32(load32(m.memory[int64(uint32(v1))+104:]))
						v5 = t96
						store64(m.memory[int64(uint32(v1))+100:], uint64(i64(0x100000000)))
						store32(m.memory[int64(uint32(v1))+108:], uint32(i32(0)))
						m.fn865(v1+i32(128), v5, v3)
						m.fn16(v19, v5)
						{
							t97 := int32(load32(m.memory[int64(uint32(v1))+136:]))
							if t97 == 0 {
								goto l45
							}
							t98 := int32(load32(m.memory[int64(uint32(v1))+136:]))
							store32(m.memory[int64(uint32(v20))+8:], uint32(t98))
							t99 := int64(load64(m.memory[int64(uint32(v1))+128:]))
							store64(m.memory[uint32(v20):], uint64(t99))
							store32(m.memory[int64(uint32(v1))+168:], uint32(v25))
							store32(m.memory[int64(uint32(v1))+152:], uint32(i32(3)))
							m.fn1340(v1+i32(88), v1+i32(152))
							goto l44
						}
					l45:
						t100 := int32(load32(m.memory[int64(uint32(v1))+128:]))
						t101 := int32(load32(m.memory[int64(uint32(v1))+132:]))
						m.fn16(t100, t101)
					}
				l44:
					store32(m.memory[int64(uint32(v1))+152:], uint32(i32(8)))
					m.fn1340(v1+i32(88), v1+i32(152))
					goto l43
				case 2:
					{
						t82 := int32(load32(m.memory[int64(uint32(v1))+108:]))
						v3 = t82
						if v3 == 0 {
							goto l39
						}
						t83 := int32(load32(m.memory[int64(uint32(v1))+100:]))
						v19 = t83
						t84 := int32(load32(m.memory[int64(uint32(v1))+104:]))
						v5 = t84
						store64(m.memory[int64(uint32(v1))+100:], uint64(i64(0x100000000)))
						store32(m.memory[int64(uint32(v1))+108:], uint32(i32(0)))
						m.fn865(v1+i32(128), v5, v3)
						m.fn16(v19, v5)
						{
							t85 := int32(load32(m.memory[int64(uint32(v1))+136:]))
							if t85 == 0 {
								goto l40
							}
							t86 := int32(load32(m.memory[int64(uint32(v1))+136:]))
							store32(m.memory[int64(uint32(v20))+8:], uint32(t86))
							t87 := int64(load64(m.memory[int64(uint32(v1))+128:]))
							store64(m.memory[uint32(v20):], uint64(t87))
							store32(m.memory[int64(uint32(v1))+168:], uint32(v25))
							store32(m.memory[int64(uint32(v1))+152:], uint32(i32(3)))
							m.fn1340(v1+i32(88), v1+i32(152))
							goto l39
						}
					l40:
						t88 := int32(load32(m.memory[int64(uint32(v1))+128:]))
						t89 := int32(load32(m.memory[int64(uint32(v1))+132:]))
						m.fn16(t88, t89)
					}
				l39:
					{
						if v9 != 0 {
							goto l41
						}
						v5 = i32(2)
						v3 = i32(0)
						goto l42
					l41:
						t90 := int32(m.memory[int64(uint32(v9))+6])
						v5 = t90
						t91 := int32(load16(m.memory[int64(uint32(v9))+4:]))
						v3 = t91
					}
				l42:
					t92 := int64(load64(m.memory[int64(uint32(v1))+88:]))
					v8 = t92
					store64(m.memory[int64(uint32(v1))+88:], uint64(i64(0x400000000)))
					t93 := int32(load32(m.memory[int64(uint32(v1))+96:]))
					v19 = t93
					store32(m.memory[int64(uint32(v1))+96:], uint32(i32(0)))
					store32(m.memory[int64(uint32(v1))+160:], uint32(v19))
					store64(m.memory[int64(uint32(v1))+152:], uint64(v8))
					m.memory[int64(uint32(v1))+166] = byte(v5)
					store16(m.memory[int64(uint32(v1))+164:], uint16(v3))
					m.fn1501(v1+i32(76), v1+i32(152))
					goto l43
				default:
					goto l21
				}
			}
		l21:
			{
				t102 := m.fn822(v1+i32(124), v1+i32(112))
				if t102 == 0 {
					goto l46
				}
				t103 := int32(load32(m.memory[int64(uint32(v1))+108:]))
				v3 = t103
				if v3 == 0 {
					goto l46
				}
				t104 := int32(load32(m.memory[int64(uint32(v1))+100:]))
				v19 = t104
				t105 := int32(load32(m.memory[int64(uint32(v1))+104:]))
				v5 = t105
				store64(m.memory[int64(uint32(v1))+100:], uint64(i64(0x100000000)))
				store32(m.memory[int64(uint32(v1))+108:], uint32(i32(0)))
				m.fn865(v1+i32(128), v5, v3)
				m.fn16(v19, v5)
				{
					t106 := int32(load32(m.memory[int64(uint32(v1))+136:]))
					if t106 == 0 {
						goto l47
					}
					t107 := int32(load32(m.memory[int64(uint32(v1))+136:]))
					store32(m.memory[int64(uint32(v20))+8:], uint32(t107))
					t108 := int64(load64(m.memory[int64(uint32(v1))+128:]))
					store64(m.memory[uint32(v20):], uint64(t108))
					store32(m.memory[int64(uint32(v1))+168:], uint32(v25))
					store32(m.memory[int64(uint32(v1))+152:], uint32(i32(3)))
					m.fn1340(v1+i32(88), v1+i32(152))
					goto l46
				}
			l47:
				t109 := int32(load32(m.memory[int64(uint32(v1))+128:]))
				t110 := int32(load32(m.memory[int64(uint32(v1))+132:]))
				m.fn16(t109, t110)
			}
		l46:
			t111 := int32(load32(m.memory[int64(uint32(v1))+124:]))
			t112 := v1
			v25 = t111
			store32(m.memory[int64(uint32(t112))+112:], uint32(v25))
			m.fn74(v1+i32(100), v6)
		}
	l43:
		{
			{
				t114 := v11
				p113 := i32(2)
				if uint32(v6) < uint32(i32(65536)) {
					p113 = i32(1)
				}
				v5 = p113
				if uint32(t114) <= uint32(v5) {
					goto l48
				}
				v6 = v11 - v5
				p115 := v6
				if uint32(v6) > uint32(v11) {
					p115 = i32(0)
				}
				v11 = p115
				v6 = v18
				goto l49
			}
		l48:
			if v18 != v23 {
				goto l50
			}
			v11 = i32(-1)
			v17 = i32(0)
			v6 = v23
			goto l49
		l50:
			v6 = v18 + i32(8)
			t116 := int32(load32(m.memory[uint32(v18):]))
			v11 = t116
			v17 = v18
		}
	l49:
		{
			{
				if uint32(v13) <= uint32(v5) {
					goto l51
				}
				v5 = v13 - v5
				p117 := v5
				if uint32(v5) > uint32(v13) {
					p117 = i32(0)
				}
				v13 = p117
				v5 = v22
				goto l52
			}
		l51:
			if v22 != v21 {
				goto l53
			}
			v13 = i32(-1)
			v9 = i32(0)
			v5 = v21
			goto l52
		l53:
			v5 = v22 + i32(8)
			t118 := int32(load32(m.memory[uint32(v22):]))
			v13 = t118
			v9 = v22
		}
	l52:
		v22 = v5
		v18 = v6
		goto l54
	}
l37:
	m.fn765(v16, v15)
	m.fn1499(v1 + i32(52))
	m.fn16(v2, v24)
l0:
	m.g0 = v1 + i32(208)
}
func (m *Module) fn1351(v0, v1 int32) {
	var v2, v3 int32
	var v4 int64
	t0 := m.g0
	v2 = t0 - i32(32)
	m.g0 = v2
	t1 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	v3 = t1
	store32(m.memory[int64(uint32(v1))+8:], uint32(i32(0)))
	t2 := int64(load64(m.memory[uint32(v1):]))
	v4 = t2
	store64(m.memory[uint32(v1):], uint64(i64(0x800000000)))
	store32(m.memory[int64(uint32(v2))+16:], uint32(v3))
	store64(m.memory[int64(uint32(v2))+8:], uint64(v4))
	if v3 == 0 {
		goto l0
	}
	m.fn1453(v2+i32(20), v2+i32(8))
	m.fn1271(v0, v2+i32(20))
	goto l1
l0:
	m.fn1302(v2 + i32(8))
l1:
	m.g0 = v2 + i32(32)
}
func (m *Module) fn1352(v0 int32) {
	t0 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	t1 := int32(load32(m.memory[int64(uint32(v0))+12:]))
	m.fn128(t0, t1)
	m.fn127(v0 + i32(40))
}
func (m *Module) fn1353(v0 int32) {
	t0 := int32(load32(m.memory[int64(uint32(v0))+28:]))
	t1 := int32(load32(m.memory[int64(uint32(v0))+32:]))
	m.fn134(t0, t1)
	m.fn969(v0 + i32(40))
}
func (m *Module) fn1354(v0, v1 int32) {
	m.fn136(v0, v1, i32(8), i32(56))
}
func (m *Module) fn1355(v0 int32) {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	m.fn16(t0, t1)
	m.fn1498(v0 + i32(12))
}
func (m *Module) fn1356(v0, v1, v2 int32) {
	var v3, v4, v5, v6, v7, v8, v9 int32
	var v10 int64
	t0 := m.g0
	v3 = t0 - i32(32)
	m.g0 = v3
	{
		{
			t1 := m.fn886(v1, v2, i32(1074411), i32(53), i32(1086315), i32(6))
			if t1 == 0 {
				goto l0
			}
			v4 = i32(1)
			goto l1
		}
	l0:
		{
			t2 := m.fn886(v1, v2, i32(1074411), i32(53), i32(1086321), i32(9))
			v5 = t2
			if v5 == 0 {
				goto l2
			}
			t3 := int32(load32(m.memory[uint32(v5+i32(16)):]))
			t4 := v3 + i32(8)
			v6 = t3
			t5 := int32(load32(m.memory[uint32(v5+i32(20)):]))
			t6 := v6
			v7 = t5
			m.fn1046(t4, t6, v7, i32(1074411), i32(53), i32(1074404), i32(4))
			v8 = i32(2)
			v9 = i32(2)
			{
				t7 := int32(load32(m.memory[int64(uint32(v3))+8:]))
				v4 = t7
				p8 := i32(1)
				if v4 != 0 {
					p8 = v4
				}
				v5 = p8
				t9 := int32(load32(m.memory[int64(uint32(v3))+12:]))
				t11 := v5
				p10 := i32(0)
				if v4 != 0 {
					p10 = t9
				}
				v4 = p10
				t12 := m.fn159(t11, v4, i32(1086330), i32(7))
				if t12 != 0 {
					goto l3
				}
				{
					t13 := m.fn159(v5, v4, i32(1086337), i32(7))
					if t13 == 0 {
						goto l4
					}
					v9 = i32(3)
					goto l3
				}
			l4:
				{
					t14 := m.fn159(v5, v4, i32(1086344), i32(7))
					if t14 == 0 {
						goto l5
					}
					v9 = i32(4)
					goto l3
				}
			l5:
				t15 := m.fn159(v5, v4, i32(1086351), i32(7))
				p16 := i32(1)
				if t15 != 0 {
					p16 = i32(5)
				}
				v9 = p16
			}
		l3:
			{
				t17 := m.fn1061(v5, v4, i32(1086358), i32(9))
				if t17 != 0 {
					goto l6
				}
				v8 = i32(1)
				t18 := m.fn1061(v5, v4, i32(1086367), i32(6))
				if t18 != 0 {
					goto l6
				}
				t19 := m.fn1061(v5, v4, i32(1086373), i32(5))
				p20 := i32(0)
				if t19 != 0 {
					p20 = i32(3)
				}
				v8 = p20
			}
		l6:
			m.fn1046(v3, v6, v7, i32(1074411), i32(53), i32(1086378), i32(7))
			v4 = i32(3)
			v10 = i64(1)
			t21 := int32(load32(m.memory[uint32(v3):]))
			v5 = t21
			if v5 == 0 {
				goto l1
			}
			t22 := int32(load32(m.memory[int64(uint32(v3))+4:]))
			m.fn1322(v3+i32(16), v5, t22)
			t23 := int32(m.memory[int64(uint32(v3))+16])
			if t23 == i32(1) {
				goto l1
			}
			t24 := int64(load64(m.memory[int64(uint32(v3))+24:]))
			v10 = t24
			p25 := i64(1)
			if v10 > i64(1) {
				p25 = v10
			}
			v10 = p25
			p26 := i64(0x7fff)
			if v10 < i64(0x7fff) {
				p26 = v10
			}
			v10 = p26
			goto l1
		}
	l2:
		t27 := m.fn886(v1, v2, i32(1074411), i32(53), i32(1086385), i32(6))
		p28 := i32(0)
		if t27 != 0 {
			p28 = i32(2)
		}
		v4 = p28
	}
l1:
	v5 = i32(255)
	{
		t29 := m.fn886(v1, v2, i32(1074411), i32(53), i32(1086391), i32(6))
		v2 = t29
		if v2 == 0 {
			goto l7
		}
		t30 := int32(load32(m.memory[uint32(v2+i32(16)):]))
		t31 := int32(load32(m.memory[uint32(v2+i32(20)):]))
		t32 := m.fn1493(t30, t31)
		v5 = t32
	}
l7:
	store64(m.memory[int64(uint32(v0))+8:], uint64(v10))
	m.memory[int64(uint32(v0))+2] = byte(v8)
	m.memory[int64(uint32(v0))+1] = byte(v9)
	m.memory[uint32(v0)] = byte(v4)
	t34 := v0
	p33 := v5
	if v5&i32(255) == i32(255) {
		p33 = i32(33686018)
	}
	store32(m.memory[int64(uint32(t34))+16:], uint32(p33))
	m.g0 = v3 + i32(32)
}
func (m *Module) fn1357(v0, v1, v2, v3 int32) {
	t0 := int32(load32(m.memory[uint32(v1):]))
	t1 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	t2 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	t3 := int32(load32(m.memory[int64(uint32(v1))+12:]))
	m.fn1496(v0, t0, t1, t2, t3, v2, v3)
}
func (m *Module) fn1358(v0, v1, v2 int32) {
	if v1 == 0 {
		goto l0
	}
	m.fn865(v0, v1, v2)
	return
l0:
	store32(m.memory[uint32(v0):], uint32(i32(-1)))
}
func (m *Module) fn1359(v0, v1, v2, v3, v4, v5, v6, v7 int32) {
	var v8, v9 int32
	var v10 int64
	t0 := m.g0
	v8 = t0 - i32(64)
	m.g0 = v8
	{
		t1 := m.fn846(v2, v6, v7)
		v9 = t1
		if v9 == 0 {
			store64(m.memory[uint32(v0):], uint64(i64(-1)))
			goto l3
		}
		{
			t2 := int32(m.memory[int64(uint32(v9))+24])
			if t2 != 0 {
				{
					{
						t8 := int32(load32(m.memory[uint32(v9+i32(8)):]))
						v7 = t8
						if v7 != 0 {
							goto l4
						}
						store32(m.memory[int64(uint32(v8))+28:], uint32(i32(-1)))
						goto l5
					}
				l4:
					t9 := int32(load32(m.memory[uint32(v9+i32(4)):]))
					m.fn31(v8+i32(28), t9, v7)
				}
			l5:
				t10 := int32(load32(m.memory[int64(uint32(v8))+36:]))
				store32(m.memory[int64(uint32(v0))+12:], uint32(t10))
				t11 := int64(load64(m.memory[int64(uint32(v8))+28:]))
				store64(m.memory[int64(uint32(v0))+4:], uint64(t11))
				store32(m.memory[uint32(v0):], uint32(i32(-1)))
				goto l3
			}
			m.fn1496(v8+i32(28), v1, v2, v3, v4, v6, v7)
			t3 := int32(load32(m.memory[int64(uint32(v8))+48:]))
			v7 = t3
			t4 := int32(load32(m.memory[int64(uint32(v8))+44:]))
			v6 = t4
			t5 := int64(load64(m.memory[int64(uint32(v8))+36:]))
			v10 = t5
			t6 := int32(load32(m.memory[int64(uint32(v8))+32:]))
			v9 = t6
			t7 := int32(load32(m.memory[int64(uint32(v8))+28:]))
			v2 = t7
			if v2 == i32(-1) {
				if v9 == i32(-1) {
					goto l6
				}
				store32(m.memory[int64(uint32(v8))+12:], uint32(v7))
				store32(m.memory[int64(uint32(v8))+8:], uint32(v6))
				m.fn1476(v8+i32(16), int32(v10), int32(int64(uint64(v10)>>32)))
				m.fn1182(v8, v5, i32(1081364))
				t12 := int32(load32(m.memory[int64(uint32(v8))+4:]))
				v2 = t12
				t13 := int32(load32(m.memory[uint32(v8):]))
				v4 = t13
				store64(m.memory[int64(uint32(v8))+56:], uint64(v10))
				store32(m.memory[int64(uint32(v8))+52:], uint32(v9))
				m.fn1296(v8+i32(28), v4, v8+i32(16), v8+i32(52), v6+i32(8), v7)
				t14 := int32(load32(m.memory[int64(uint32(v8))+32:]))
				v9 = t14
				{
					t15 := int32(load32(m.memory[int64(uint32(v8))+28:]))
					v7 = t15
					if v7 == i32(-1) {
						store32(m.memory[int64(uint32(v0))+8:], uint32(v9))
						store64(m.memory[uint32(v0):], uint64(i64(-0x7fffffff00000001)))
						t19 := int32(load32(m.memory[uint32(v2):]))
						store32(m.memory[uint32(v2):], uint32(t19+i32(1)))
						m.fn754(v8 + i32(8))
						goto l3
					}
					t16 := int64(load64(m.memory[int64(uint32(v8))+44:]))
					store64(m.memory[int64(uint32(v0))+16:], uint64(t16))
					t17 := int64(load64(m.memory[int64(uint32(v8))+36:]))
					store64(m.memory[int64(uint32(v0))+8:], uint64(t17))
					store32(m.memory[int64(uint32(v0))+4:], uint32(v9))
					store32(m.memory[uint32(v0):], uint32(v7))
					t18 := int32(load32(m.memory[uint32(v2):]))
					store32(m.memory[uint32(v2):], uint32(t18+i32(1)))
					m.fn754(v8 + i32(8))
					goto l3
				}
			}
			store32(m.memory[int64(uint32(v0))+20:], uint32(v7))
			store32(m.memory[int64(uint32(v0))+16:], uint32(v6))
			store64(m.memory[int64(uint32(v0))+8:], uint64(v10))
			store32(m.memory[int64(uint32(v0))+4:], uint32(v9))
			store32(m.memory[uint32(v0):], uint32(v2))
			goto l3
		}
	}
l6:
	store64(m.memory[uint32(v0):], uint64(i64(-1)))
l3:
	m.g0 = v8 + i32(64)
}
func (m *Module) fn1360(v0, v1, v2 int32) {
	{
		t0 := int32(load32(m.memory[uint32(v1):]))
		if t0 == i32(-1) {
			goto l0
		}
		t1 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		store32(m.memory[int64(uint32(v0))+8:], uint32(t1))
		t2 := int64(load64(m.memory[uint32(v1):]))
		store64(m.memory[uint32(v0):], uint64(t2))
		t3 := int32(load32(m.memory[uint32(v2):]))
		t4 := int32(load32(m.memory[int64(uint32(v2))+4:]))
		m.fn895(t3, t4)
		return
	}
l0:
	t5 := int32(load32(m.memory[int64(uint32(v2))+8:]))
	store32(m.memory[int64(uint32(v0))+8:], uint32(t5))
	t6 := int64(load64(m.memory[uint32(v2):]))
	store64(m.memory[uint32(v0):], uint64(t6))
}
func (m *Module) fn1361(v0, v1, v2, v3 int32) int32 {
	var v4, v5, v6, v7, v8, v9, v10 int32
	t0 := m.g0
	v4 = t0 - i32(80)
	m.g0 = v4
	store32(m.memory[int64(uint32(v4))+44:], uint32(i32(6)))
	store32(m.memory[int64(uint32(v4))+40:], uint32(i32(1080683)))
	store32(m.memory[int64(uint32(v4))+36:], uint32(i32(59)))
	store32(m.memory[int64(uint32(v4))+32:], uint32(i32(1073848)))
	store32(m.memory[int64(uint32(v4))+24:], uint32(v0))
	t1 := v4
	v5 = v0 + v1*i32(44)
	store32(m.memory[int64(uint32(t1))+28:], uint32(v5))
	v6 = v4 + i32(24) + i32(8)
	v7 = v0
	{
	l5:
		if v7 == v5 {
			goto l0
		}
		v8 = i32(0)
		{
			t2 := int32(load32(m.memory[uint32(v7):]))
			var p3 int32
			if t2 == i32(-1) {
				p3 = 1
			}
			v9 = p3
			if v9 != 0 {
				goto l1
			}
			t4 := m.fn844(v6, v7)
			if t4 == 0 {
				goto l1
			}
			t5 := int32(load32(m.memory[uint32(v7+i32(16)):]))
			t6 := int32(load32(m.memory[uint32(v7+i32(20)):]))
			m.fn1046(v4+i32(16), t5, t6, i32(1073848), i32(59), i32(1074595), i32(8))
			{
				t7 := int32(load32(m.memory[int64(uint32(v4))+16:]))
				v10 = t7
				if v10 != 0 {
					goto l2
				}
				v8 = v7
				goto l3
			}
		l2:
			p8 := v7
			if v9 != 0 {
				p8 = i32(0)
			}
			v8 = p8
			t9 := int32(load32(m.memory[int64(uint32(v4))+20:]))
			v9 = t9
			store16(m.memory[int64(uint32(v4))+76:], uint16(i32(1)))
			store32(m.memory[int64(uint32(v4))+72:], uint32(i32(0)))
			store32(m.memory[int64(uint32(v4))+64:], uint32(v10))
			store32(m.memory[int64(uint32(v4))+56:], uint32(v10))
			store32(m.memory[int64(uint32(v4))+48:], uint32(i32(0)))
			store32(m.memory[int64(uint32(v4))+60:], uint32(v9))
			store32(m.memory[int64(uint32(v4))+52:], uint32(v9))
			store32(m.memory[int64(uint32(v4))+68:], uint32(v10+v9))
		l4:
			{
				m.fn875(v4+i32(8), v4+i32(48))
				t10 := int32(load32(m.memory[int64(uint32(v4))+8:]))
				v9 = t10
				if v9 == 0 {
					goto l1
				}
				t11 := int32(load32(m.memory[int64(uint32(v4))+12:]))
				t12 := m.fn914(v9, t11, v2, v3)
				if t12 != 0 {
					goto l4
				}
			}
			v8 = i32(0)
		}
	l1:
		v7 = v7 + i32(44)
		if v8 == 0 {
			goto l5
		}
		goto l3
	l0:
		t13 := m.fn886(v0, v1, i32(1073848), i32(59), i32(1073907), i32(8))
		v8 = t13
	}
l3:
	m.g0 = v4 + i32(80)
	return v8
}
func (m *Module) fn1362(v0, v1, v2 int32) {
	var v3, v4, v5 int32
	var v6 int64
	var v7, v8, v9, v10, v11, v12, v13, v14, v15, v16, v17, v18 int32
	t0 := m.g0
	v3 = t0 - i32(224)
	m.g0 = v3
	store32(m.memory[int64(uint32(v3))+36:], uint32(i32(0)))
	store64(m.memory[int64(uint32(v3))+28:], uint64(i64(0x800000000)))
	{
		t1 := m.fn1097(v1, v2, i32(1073932), i32(54), i32(1073751), i32(5))
		v4 = t1
		if v4 == 0 {
			goto l0
		}
		t2 := int32(load32(m.memory[uint32(v4+i32(28)):]))
		t3 := int32(load32(m.memory[uint32(v4+i32(32)):]))
		m.fn1497(v3+i32(128), t2, t3)
		t4 := int32(load32(m.memory[int64(uint32(v3))+132:]))
		t5 := v3 + i32(160)
		v4 = t4
		t6 := int32(load32(m.memory[int64(uint32(v3))+136:]))
		m.fn865(t5, v4, t6)
		t7 := int32(load32(m.memory[int64(uint32(v3))+128:]))
		m.fn16(t7, v4)
		t8 := int32(load32(m.memory[int64(uint32(v3))+160:]))
		v5 = t8
		if v5 == i32(-1) {
			goto l0
		}
		t9 := int32(load32(m.memory[int64(uint32(v3))+164:]))
		t10 := v3 + i32(16)
		v4 = t9
		t11 := int32(load32(m.memory[int64(uint32(v3))+168:]))
		m.fn46(t10, v4, t11)
		{
			t12 := int32(load32(m.memory[int64(uint32(v3))+20:]))
			if t12 != 0 {
				goto l1
			}
			m.fn16(v5, v4)
			goto l0
		}
	l1:
		t13 := int64(load64(m.memory[int64(uint32(v3))+164:]))
		v6 = t13
		t14 := m.fn113(i32(4), i32(28))
		v4 = t14
		store32(m.memory[int64(uint32(v4))+16:], uint32(i32(1)))
		store64(m.memory[int64(uint32(v4))+8:], uint64(v6))
		store32(m.memory[int64(uint32(v4))+4:], uint32(v5))
		store32(m.memory[uint32(v4):], uint32(i32(3)))
		store32(m.memory[int64(uint32(v3))+140:], uint32(i32(1)))
		store32(m.memory[int64(uint32(v3))+136:], uint32(v4))
		store64(m.memory[int64(uint32(v3))+128:], uint64(i64(0x180000000)))
		m.fn338(v3+i32(28), v3+i32(128))
	}
l0:
	v7 = i32(0)
	store32(m.memory[int64(uint32(v3))+48:], uint32(i32(0)))
	store64(m.memory[int64(uint32(v3))+40:], uint64(i64(0x400000000)))
	store32(m.memory[int64(uint32(v3))+60:], uint32(i32(0)))
	store64(m.memory[int64(uint32(v3))+52:], uint64(i64(0x400000000)))
	m.fn868(v3+i32(128), v1, v2)
	store32(m.memory[int64(uint32(v3))+148:], uint32(i32(1081466)))
	store32(m.memory[int64(uint32(v3))+144:], uint32(i32(54)))
	store32(m.memory[int64(uint32(v3))+140:], uint32(i32(1073932)))
	store32(m.memory[int64(uint32(v3))+152:], uint32(i32(3)))
	store32(m.memory[int64(uint32(v3))+88:], uint32(i32(3)))
	t15 := int64(load64(m.memory[int64(uint32(v3))+144:]))
	store64(m.memory[int64(uint32(v3))+80:], uint64(t15))
	t16 := int64(load64(m.memory[int64(uint32(v3))+136:]))
	store64(m.memory[int64(uint32(v3))+72:], uint64(t16))
	t17 := int64(load64(m.memory[int64(uint32(v3))+128:]))
	store64(m.memory[int64(uint32(v3))+64:], uint64(t17))
	v8 = i32(4)
l33:
	{
		{
			{
				t18 := m.fn863(v3 + i32(64))
				v4 = t18
				if v4 == 0 {
					t66 := int32(load32(m.memory[int64(uint32(v3))+64:]))
					t67 := int32(load32(m.memory[int64(uint32(v3))+68:]))
					m.fn44(t66, t67)
					if v7 == 0 {
						goto l16
					}
					t68 := int32(load32(m.memory[int64(uint32(v3))+48:]))
					v11 = t68
					if v11 == 0 {
						goto l16
					}
					v14 = i32(0)
					v15 = i32(1)
					v10 = i32(0)
					{
						t69 := m.fn1097(v1, v2, i32(1073932), i32(54), i32(1081469), i32(5))
						v4 = t69
						if v4 == 0 {
							goto l17
						}
						v14 = i32(0)
						v15 = i32(1)
						v10 = i32(0)
						t70 := int32(load32(m.memory[uint32(v4+i32(28)):]))
						t71 := int32(load32(m.memory[uint32(v4+i32(32)):]))
						t72 := m.fn886(t70, t71, i32(1073932), i32(54), i32(1073751), i32(5))
						v4 = t72
						if v4 == 0 {
							goto l17
						}
						t73 := int32(load32(m.memory[uint32(v4+i32(28)):]))
						t74 := int32(load32(m.memory[uint32(v4+i32(32)):]))
						m.fn1497(v3+i32(128), t73, t74)
						t75 := int32(load32(m.memory[int64(uint32(v3))+132:]))
						t76 := v3 + i32(160)
						v4 = t75
						t77 := int32(load32(m.memory[int64(uint32(v3))+136:]))
						m.fn865(t76, v4, t77)
						t78 := int32(load32(m.memory[int64(uint32(v3))+128:]))
						m.fn16(t78, v4)
						v14 = i32(0)
						v15 = i32(1)
						v10 = i32(0)
						t79 := int32(load32(m.memory[int64(uint32(v3))+160:]))
						v4 = t79
						if v4 == i32(-1) {
							goto l17
						}
						t80 := int32(load32(m.memory[int64(uint32(v3))+168:]))
						v10 = t80
						t81 := int32(load32(m.memory[int64(uint32(v3))+164:]))
						v15 = t81
						v14 = v4
					}
				l17:
					t82 := m.fn113(i32(4), i32(20))
					v5 = t82
					t83 := m.fn113(i32(4), i32(28))
					v4 = t83
					store32(m.memory[int64(uint32(v4))+16:], uint32(i32(0)))
					store32(m.memory[int64(uint32(v4))+12:], uint32(v10))
					store32(m.memory[int64(uint32(v4))+8:], uint32(v15))
					store32(m.memory[int64(uint32(v4))+4:], uint32(v14))
					store32(m.memory[uint32(v4):], uint32(i32(3)))
					store32(m.memory[int64(uint32(v3))+168:], uint32(i32(1)))
					store32(m.memory[int64(uint32(v3))+164:], uint32(v4))
					store32(m.memory[int64(uint32(v3))+160:], uint32(i32(1)))
					m.fn888(v3+i32(128), v3+i32(160))
					t84 := int32(load32(m.memory[int64(uint32(v3))+144:]))
					store32(m.memory[int64(uint32(v5))+16:], uint32(t84))
					t85 := int64(load64(m.memory[int64(uint32(v3))+136:]))
					store64(m.memory[int64(uint32(v5))+8:], uint64(t85))
					t86 := int64(load64(m.memory[int64(uint32(v3))+128:]))
					store64(m.memory[uint32(v5):], uint64(t86))
					store32(m.memory[int64(uint32(v3))+72:], uint32(i32(1)))
					store32(m.memory[int64(uint32(v3))+68:], uint32(v5))
					store32(m.memory[int64(uint32(v3))+64:], uint32(i32(1)))
					v14 = v8 + i32(8)
					m.fn887(v3+i32(64), v7)
					t87 := int32(load32(m.memory[int64(uint32(v3))+72:]))
					v4 = t87
					v10 = v4 + v7
					t88 := int32(load32(m.memory[int64(uint32(v3))+68:]))
					v5 = t88 + v4*i32(20)
					v15 = v7
				l18:
					{
						t89 := m.fn113(i32(4), i32(28))
						v4 = t89
						t90 := int32(load32(m.memory[uint32(v14+i32(-4)):]))
						t91 := int32(load32(m.memory[uint32(v14):]))
						m.fn31(v3+i32(128), t90, t91)
						store32(m.memory[uint32(v4):], uint32(i32(3)))
						store32(m.memory[int64(uint32(v4))+16:], uint32(i32(0)))
						t92 := int64(load64(m.memory[int64(uint32(v3))+128:]))
						store64(m.memory[int64(uint32(v4))+4:], uint64(t92))
						t93 := int32(load32(m.memory[int64(uint32(v3))+136:]))
						store32(m.memory[int64(uint32(v4))+12:], uint32(t93))
						store32(m.memory[int64(uint32(v3))+168:], uint32(i32(1)))
						store32(m.memory[int64(uint32(v3))+164:], uint32(v4))
						store32(m.memory[int64(uint32(v3))+160:], uint32(i32(1)))
						m.fn888(v3+i32(128), v3+i32(160))
						t94 := int32(load32(m.memory[int64(uint32(v3))+144:]))
						store32(m.memory[int64(uint32(v5))+16:], uint32(t94))
						t95 := int64(load64(m.memory[int64(uint32(v3))+136:]))
						store64(m.memory[int64(uint32(v5))+8:], uint64(t95))
						t96 := int64(load64(m.memory[int64(uint32(v3))+128:]))
						store64(m.memory[uint32(v5):], uint64(t96))
						v14 = v14 + i32(24)
						v5 = v5 + i32(20)
						v15 = v15 + i32(-1)
						if v15 != 0 {
							goto l18
						}
					}
					store32(m.memory[int64(uint32(v3))+72:], uint32(v10))
					t97 := m.fn113(i32(4), i32(12))
					v4 = t97
					t98 := int32(load32(m.memory[int64(uint32(v3))+72:]))
					store32(m.memory[int64(uint32(v4))+8:], uint32(t98))
					t99 := int64(load64(m.memory[int64(uint32(v3))+64:]))
					store64(m.memory[uint32(v4):], uint64(t99))
					store32(m.memory[int64(uint32(v3))+112:], uint32(i32(1)))
					store32(m.memory[int64(uint32(v3))+108:], uint32(v4))
					store32(m.memory[int64(uint32(v3))+104:], uint32(i32(1)))
					v1 = v7 * i32(24)
					t100 := int32(load32(m.memory[int64(uint32(v3))+44:]))
					v16 = t100
					v2 = v16 + v11*i32(12)
					v18 = i32(0)
				l24:
					{
						if v16 == v2 {
							m.fn1164(v3+i32(128)|i32(4), v3+i32(104), i32(1))
							store32(m.memory[int64(uint32(v3))+128:], uint32(i32(-0x7ffffffe)))
							m.fn338(v3+i32(28), v3+i32(128))
							goto l16
						}
						v10 = i32(20)
						t101 := m.fn113(i32(4), i32(20))
						v17 = t101
						t102 := m.fn113(i32(4), i32(28))
						v4 = t102
						t103 := int32(load32(m.memory[uint32(v16+i32(4)):]))
						t104 := int32(load32(m.memory[uint32(v16+i32(8)):]))
						m.fn31(v3+i32(128), t103, t104)
						store32(m.memory[uint32(v4):], uint32(i32(3)))
						store32(m.memory[int64(uint32(v4))+16:], uint32(i32(0)))
						t105 := int64(load64(m.memory[int64(uint32(v3))+128:]))
						store64(m.memory[int64(uint32(v4))+4:], uint64(t105))
						t106 := int32(load32(m.memory[int64(uint32(v3))+136:]))
						store32(m.memory[int64(uint32(v4))+12:], uint32(t106))
						v11 = i32(1)
						store32(m.memory[int64(uint32(v3))+168:], uint32(i32(1)))
						store32(m.memory[int64(uint32(v3))+164:], uint32(v4))
						store32(m.memory[int64(uint32(v3))+160:], uint32(i32(1)))
						m.fn888(v3+i32(128), v3+i32(160))
						t107 := int32(load32(m.memory[int64(uint32(v3))+144:]))
						store32(m.memory[int64(uint32(v17))+16:], uint32(t107))
						t108 := int64(load64(m.memory[int64(uint32(v3))+136:]))
						store64(m.memory[int64(uint32(v17))+8:], uint64(t108))
						t109 := int64(load64(m.memory[int64(uint32(v3))+128:]))
						store64(m.memory[uint32(v17):], uint64(t109))
						store32(m.memory[int64(uint32(v3))+124:], uint32(i32(1)))
						store32(m.memory[int64(uint32(v3))+120:], uint32(v17))
						store32(m.memory[int64(uint32(v3))+116:], uint32(i32(1)))
						v15 = v1
						v5 = v8
					l23:
						{
							if v15 == 0 {
								v18 = v18 + i32(1)
								v16 = v16 + i32(12)
								m.fn1169(v3+i32(104), v3+i32(116))
								goto l24
							}
							v14 = i32(-1)
							{
								t110 := int32(load32(m.memory[int64(uint32(v5))+20:]))
								if uint32(v18) >= uint32(t110) {
									goto l21
								}
								t111 := int32(load32(m.memory[int64(uint32(v5))+16:]))
								t112 := v3 + i32(128)
								v4 = t111 + v18*i32(12)
								t113 := int32(load32(m.memory[uint32(v4+i32(4)):]))
								t114 := int32(load32(m.memory[uint32(v4+i32(8)):]))
								m.fn31(t112, t113, t114)
								t115 := int32(load32(m.memory[int64(uint32(v3))+128:]))
								v14 = t115
							}
						l21:
							t116 := int32(load32(m.memory[int64(uint32(v3))+132:]))
							v12 = t116
							t117 := int32(load32(m.memory[int64(uint32(v3))+136:]))
							v13 = t117
							t118 := m.fn113(i32(4), i32(28))
							v4 = t118
							store32(m.memory[int64(uint32(v4))+16:], uint32(i32(0)))
							t119 := v4
							t120 := v13
							var p121 int32
							if v14 == i32(-1) {
								p121 = 1
							}
							v9 = p121
							p122 := t120
							if v9 != 0 {
								p122 = i32(0)
							}
							store32(m.memory[int64(uint32(t119))+12:], uint32(p122))
							t124 := v4
							p123 := v12
							if v9 != 0 {
								p123 = i32(1)
							}
							store32(m.memory[int64(uint32(t124))+8:], uint32(p123))
							t126 := v4
							p125 := v14
							if v9 != 0 {
								p125 = i32(0)
							}
							store32(m.memory[int64(uint32(t126))+4:], uint32(p125))
							store32(m.memory[uint32(v4):], uint32(i32(3)))
							store32(m.memory[int64(uint32(v3))+168:], uint32(i32(1)))
							store32(m.memory[int64(uint32(v3))+164:], uint32(v4))
							store32(m.memory[int64(uint32(v3))+160:], uint32(i32(1)))
							m.fn888(v3+i32(128), v3+i32(160))
							{
								t127 := int32(load32(m.memory[int64(uint32(v3))+116:]))
								if v11 != t127 {
									goto l22
								}
								m.fn418(v3 + i32(116))
								t128 := int32(load32(m.memory[int64(uint32(v3))+120:]))
								v17 = t128
							}
						l22:
							v5 = v5 + i32(24)
							v4 = v17 + v10
							t129 := int64(load64(m.memory[int64(uint32(v3))+128:]))
							store64(m.memory[uint32(v4):], uint64(t129))
							t130 := int32(load32(m.memory[int64(uint32(v3))+144:]))
							store32(m.memory[int64(uint32(v4))+16:], uint32(t130))
							t131 := int64(load64(m.memory[int64(uint32(v3))+136:]))
							store64(m.memory[int64(uint32(v4))+8:], uint64(t131))
							t132 := v3
							v11 = v11 + i32(1)
							store32(m.memory[int64(uint32(t132))+124:], uint32(v11))
							v15 = v15 + i32(-24)
							v10 = v10 + i32(20)
							goto l23
						}
					}
				}
				v9 = i32(0)
				{
					{
						{
							v10 = v4 + i32(28)
							t19 := int32(load32(m.memory[uint32(v10):]))
							v11 = v4 + i32(32)
							t20 := int32(load32(m.memory[uint32(v11):]))
							t21 := m.fn886(t19, t20, i32(1073932), i32(54), i32(1081474), i32(2))
							v4 = t21
							if v4 != 0 {
								goto l3
							}
							v12 = i32(1)
							goto l4
						}
					l3:
						v12 = i32(1)
						t22 := int32(load32(m.memory[uint32(v4+i32(28)):]))
						t23 := int32(load32(m.memory[uint32(v4+i32(32)):]))
						t24 := m.fn1097(t22, t23, i32(1073932), i32(54), i32(1072447), i32(1))
						v4 = t24
						if v4 != 0 {
							goto l5
						}
					}
				l4:
					v13 = i32(0)
					goto l6
				l5:
					t25 := int32(load32(m.memory[uint32(v4+i32(28)):]))
					t26 := int32(load32(m.memory[uint32(v4+i32(32)):]))
					m.fn864(v3+i32(128), t25, t26)
					t27 := int32(load32(m.memory[int64(uint32(v3))+132:]))
					t28 := v3 + i32(160)
					v4 = t27
					t29 := int32(load32(m.memory[int64(uint32(v3))+136:]))
					m.fn865(t28, v4, t29)
					t30 := int32(load32(m.memory[int64(uint32(v3))+128:]))
					m.fn16(t30, v4)
					v13 = i32(0)
					t31 := int32(load32(m.memory[int64(uint32(v3))+160:]))
					v4 = t31
					if v4 == i32(-1) {
						goto l6
					}
					t32 := int32(load32(m.memory[int64(uint32(v3))+168:]))
					v9 = t32
					t33 := int32(load32(m.memory[int64(uint32(v3))+164:]))
					v12 = t33
					v13 = v4
				}
			l6:
				{
					t34 := int32(load32(m.memory[uint32(v10):]))
					t35 := int32(load32(m.memory[uint32(v11):]))
					t36 := m.fn886(t34, t35, i32(1073932), i32(54), i32(1081476), i32(3))
					v4 = t36
					if v4 == 0 {
						goto l7
					}
					t37 := int32(load32(m.memory[uint32(v4+i32(28)):]))
					t38 := int32(load32(m.memory[uint32(v4+i32(32)):]))
					m.fn868(v3+i32(160), t37, t38)
					store32(m.memory[int64(uint32(v3))+184:], uint32(i32(1)))
					store32(m.memory[int64(uint32(v3))+180:], uint32(i32(1072447)))
					store32(m.memory[int64(uint32(v3))+176:], uint32(i32(54)))
					store32(m.memory[int64(uint32(v3))+172:], uint32(i32(1073932)))
					m.fn862(v3+i32(200), v3+i32(160))
					{
						t39 := int32(load32(m.memory[int64(uint32(v3))+200:]))
						if t39 == i32(-1) {
							goto l8
						}
						v5 = i32(12)
						m.fn59(v3+i32(8), i32(4), i32(4), i32(12))
						t40 := int32(load32(m.memory[int64(uint32(v3))+8:]))
						v14 = t40
						t41 := int32(load32(m.memory[int64(uint32(v3))+12:]))
						v15 = t41
						t42 := int32(load32(m.memory[int64(uint32(v3))+208:]))
						store32(m.memory[int64(uint32(v15))+8:], uint32(t42))
						t43 := int64(load64(m.memory[int64(uint32(v3))+200:]))
						store64(m.memory[uint32(v15):], uint64(t43))
						v4 = i32(1)
						store32(m.memory[int64(uint32(v3))+196:], uint32(i32(1)))
						store32(m.memory[int64(uint32(v3))+192:], uint32(v15))
						store32(m.memory[int64(uint32(v3))+188:], uint32(v14))
						t44 := int32(load32(m.memory[int64(uint32(v3))+184:]))
						store32(m.memory[int64(uint32(v3))+152:], uint32(t44))
						t45 := int64(load64(m.memory[int64(uint32(v3))+176:]))
						store64(m.memory[int64(uint32(v3))+144:], uint64(t45))
						t46 := int64(load64(m.memory[int64(uint32(v3))+168:]))
						store64(m.memory[int64(uint32(v3))+136:], uint64(t46))
						t47 := int64(load64(m.memory[int64(uint32(v3))+160:]))
						store64(m.memory[int64(uint32(v3))+128:], uint64(t47))
					l11:
						{
							m.fn862(v3+i32(212), v3+i32(128))
							t48 := int32(load32(m.memory[int64(uint32(v3))+212:]))
							if t48 == i32(-1) {
								t54 := int32(load32(m.memory[int64(uint32(v3))+128:]))
								t55 := int32(load32(m.memory[int64(uint32(v3))+132:]))
								m.fn44(t54, t55)
								t56 := int32(load32(m.memory[int64(uint32(v3))+188:]))
								v5 = t56
								if v5 == i32(-1) {
									goto l7
								}
								t57 := int32(load32(m.memory[int64(uint32(v3))+192:]))
								v14 = t57
								goto l12
							}
							{
								t49 := int32(load32(m.memory[int64(uint32(v3))+188:]))
								if v4 != t49 {
									goto l10
								}
								m.fn60(v3+i32(188), i32(1))
								t50 := int32(load32(m.memory[int64(uint32(v3))+192:]))
								v15 = t50
							}
						l10:
							v14 = v15 + v5
							t51 := int32(load32(m.memory[int64(uint32(v3))+220:]))
							store32(m.memory[int64(uint32(v14))+8:], uint32(t51))
							t52 := int64(load64(m.memory[int64(uint32(v3))+212:]))
							store64(m.memory[uint32(v14):], uint64(t52))
							t53 := v3
							v4 = v4 + i32(1)
							store32(m.memory[int64(uint32(t53))+196:], uint32(v4))
							v5 = v5 + i32(12)
							goto l11
						}
					}
				l8:
					t58 := int32(load32(m.memory[int64(uint32(v3))+160:]))
					t59 := int32(load32(m.memory[int64(uint32(v3))+164:]))
					m.fn44(t58, t59)
				}
			l7:
				v14 = i32(4)
				v5 = i32(0)
				v4 = i32(0)
			l12:
				store32(m.memory[int64(uint32(v3))+100:], uint32(v4))
				store32(m.memory[int64(uint32(v3))+96:], uint32(v14))
				store32(m.memory[int64(uint32(v3))+92:], uint32(v5))
				{
					t60 := int32(load32(m.memory[int64(uint32(v3))+48:]))
					v16 = t60
					if v16 != 0 {
						goto l13
					}
					m.fn78(v3 + i32(40))
					t61 := int32(load32(m.memory[int64(uint32(v3))+100:]))
					store32(m.memory[int64(uint32(v3))+48:], uint32(t61))
					t62 := int64(load64(m.memory[int64(uint32(v3))+92:]))
					store64(m.memory[int64(uint32(v3))+40:], uint64(t62))
				}
			l13:
				v17 = i32(4)
				v18 = i32(0)
				t63 := int32(load32(m.memory[uint32(v10):]))
				t64 := int32(load32(m.memory[uint32(v11):]))
				t65 := m.fn886(t63, t64, i32(1073932), i32(54), i32(1073156), i32(3))
				v4 = t65
				if v4 != 0 {
					t133 := int32(load32(m.memory[uint32(v4+i32(28)):]))
					t134 := int32(load32(m.memory[uint32(v4+i32(32)):]))
					m.fn868(v3+i32(160), t133, t134)
					store32(m.memory[int64(uint32(v3))+184:], uint32(i32(1)))
					store32(m.memory[int64(uint32(v3))+180:], uint32(i32(1072447)))
					store32(m.memory[int64(uint32(v3))+176:], uint32(i32(54)))
					store32(m.memory[int64(uint32(v3))+172:], uint32(i32(1073932)))
					m.fn862(v3+i32(200), v3+i32(160))
					{
						t135 := int32(load32(m.memory[int64(uint32(v3))+200:]))
						if t135 == i32(-1) {
							t150 := int32(load32(m.memory[int64(uint32(v3))+160:]))
							t151 := int32(load32(m.memory[int64(uint32(v3))+164:]))
							m.fn44(t150, t151)
							goto l15
						}
						v5 = i32(12)
						m.fn59(v3, i32(4), i32(4), i32(12))
						t136 := int32(load32(m.memory[uint32(v3):]))
						v14 = t136
						t137 := int32(load32(m.memory[int64(uint32(v3))+4:]))
						v15 = t137
						t138 := int32(load32(m.memory[int64(uint32(v3))+208:]))
						store32(m.memory[int64(uint32(v15))+8:], uint32(t138))
						t139 := int64(load64(m.memory[int64(uint32(v3))+200:]))
						store64(m.memory[uint32(v15):], uint64(t139))
						v4 = i32(1)
						store32(m.memory[int64(uint32(v3))+196:], uint32(i32(1)))
						store32(m.memory[int64(uint32(v3))+192:], uint32(v15))
						store32(m.memory[int64(uint32(v3))+188:], uint32(v14))
						t140 := int32(load32(m.memory[int64(uint32(v3))+184:]))
						store32(m.memory[int64(uint32(v3))+152:], uint32(t140))
						t141 := int64(load64(m.memory[int64(uint32(v3))+176:]))
						store64(m.memory[int64(uint32(v3))+144:], uint64(t141))
						t142 := int64(load64(m.memory[int64(uint32(v3))+168:]))
						store64(m.memory[int64(uint32(v3))+136:], uint64(t142))
						t143 := int64(load64(m.memory[int64(uint32(v3))+160:]))
						store64(m.memory[int64(uint32(v3))+128:], uint64(t143))
					l28:
						{
							m.fn862(v3+i32(212), v3+i32(128))
							t144 := int32(load32(m.memory[int64(uint32(v3))+212:]))
							if t144 == i32(-1) {
								t152 := int32(load32(m.memory[int64(uint32(v3))+128:]))
								t153 := int32(load32(m.memory[int64(uint32(v3))+132:]))
								m.fn44(t152, t153)
								v14 = i32(0)
								t154 := int32(load32(m.memory[int64(uint32(v3))+188:]))
								v5 = t154
								if v5 == i32(-1) {
									goto l29
								}
								t155 := int32(load32(m.memory[int64(uint32(v3))+192:]))
								v17 = t155
								v18 = v4
								v14 = v5
								goto l29
							}
							{
								t145 := int32(load32(m.memory[int64(uint32(v3))+188:]))
								if v4 != t145 {
									goto l27
								}
								m.fn60(v3+i32(188), i32(1))
								t146 := int32(load32(m.memory[int64(uint32(v3))+192:]))
								v15 = t146
							}
						l27:
							v14 = v15 + v5
							t147 := int32(load32(m.memory[int64(uint32(v3))+220:]))
							store32(m.memory[int64(uint32(v14))+8:], uint32(t147))
							t148 := int64(load64(m.memory[int64(uint32(v3))+212:]))
							store64(m.memory[uint32(v14):], uint64(t148))
							t149 := v3
							v4 = v4 + i32(1)
							store32(m.memory[int64(uint32(t149))+196:], uint32(v4))
							v5 = v5 + i32(12)
							goto l28
						}
					}
				}
				goto l15
			}
		l16:
			t156 := int32(load32(m.memory[int64(uint32(v3))+36:]))
			store32(m.memory[int64(uint32(v0))+8:], uint32(t156))
			t157 := int64(load64(m.memory[int64(uint32(v3))+28:]))
			store64(m.memory[uint32(v0):], uint64(t157))
			v4 = v8
		l31:
			{
				if v7 == 0 {
					t160 := int32(load32(m.memory[int64(uint32(v3))+52:]))
					m.fn136(t160, v8, i32(4), i32(24))
					m.fn78(v3 + i32(40))
					m.g0 = v3 + i32(224)
					return
				}
				t158 := int32(load32(m.memory[uint32(v4):]))
				t159 := int32(load32(m.memory[uint32(v4+i32(4)):]))
				m.fn16(t158, t159)
				m.fn78(v4 + i32(12))
				v7 = v7 + i32(-1)
				v4 = v4 + i32(24)
				goto l31
			}
		}
	l15:
		v14 = i32(0)
	l29:
		{
			t161 := int32(load32(m.memory[int64(uint32(v3))+60:]))
			v5 = t161
			t162 := int32(load32(m.memory[int64(uint32(v3))+52:]))
			if v5 != t162 {
				goto l32
			}
			m.fn289(v3 + i32(52))
		}
	l32:
		t163 := int32(load32(m.memory[int64(uint32(v3))+56:]))
		v8 = t163
		v4 = v8 + v5*i32(24)
		store32(m.memory[int64(uint32(v4))+20:], uint32(v18))
		store32(m.memory[int64(uint32(v4))+16:], uint32(v17))
		store32(m.memory[int64(uint32(v4))+12:], uint32(v14))
		store32(m.memory[int64(uint32(v4))+8:], uint32(v9))
		store32(m.memory[int64(uint32(v4))+4:], uint32(v12))
		store32(m.memory[uint32(v4):], uint32(v13))
		t164 := v3
		v7 = v5 + i32(1)
		store32(m.memory[int64(uint32(t164))+60:], uint32(v7))
		if v16 == 0 {
			goto l33
		}
		m.fn78(v3 + i32(92))
		goto l33
	}
}
func (m *Module) fn1363(v0, v1, v2 int32) {
	var v3, v4, v5 int32
	t0 := m.g0
	v3 = t0 - i32(144)
	m.g0 = v3
	m.fn868(v3+i32(20), v1, v2)
	store32(m.memory[int64(uint32(v3))+44:], uint32(i32(2)))
	store32(m.memory[int64(uint32(v3))+40:], uint32(i32(1081479)))
	store32(m.memory[int64(uint32(v3))+36:], uint32(i32(56)))
	store32(m.memory[int64(uint32(v3))+32:], uint32(i32(1073986)))
	m.fn885(v3+i32(60), v3+i32(20))
	{
		{
			t1 := int32(load32(m.memory[int64(uint32(v3))+60:]))
			if t1 == i32(-1) {
				goto l0
			}
			m.fn59(v3+i32(8), i32(4), i32(4), i32(28))
			t2 := int32(load32(m.memory[int64(uint32(v3))+8:]))
			v2 = t2
			t3 := int32(load32(m.memory[int64(uint32(v3))+12:]))
			v4 = t3
			t4 := int32(load32(m.memory[int64(uint32(v3))+84:]))
			store32(m.memory[int64(uint32(v4))+24:], uint32(t4))
			t5 := int64(load64(m.memory[int64(uint32(v3))+76:]))
			store64(m.memory[int64(uint32(v4))+16:], uint64(t5))
			t6 := int64(load64(m.memory[int64(uint32(v3))+68:]))
			store64(m.memory[int64(uint32(v4))+8:], uint64(t6))
			t7 := int64(load64(m.memory[int64(uint32(v3))+60:]))
			store64(m.memory[uint32(v4):], uint64(t7))
			store32(m.memory[int64(uint32(v3))+56:], uint32(i32(1)))
			store32(m.memory[int64(uint32(v3))+52:], uint32(v4))
			store32(m.memory[int64(uint32(v3))+48:], uint32(v2))
			t8 := int32(load32(m.memory[int64(uint32(v3))+44:]))
			store32(m.memory[int64(uint32(v3))+112:], uint32(t8))
			t9 := int64(load64(m.memory[int64(uint32(v3))+36:]))
			store64(m.memory[int64(uint32(v3))+104:], uint64(t9))
			t10 := int64(load64(m.memory[int64(uint32(v3))+28:]))
			store64(m.memory[int64(uint32(v3))+96:], uint64(t10))
			t11 := int64(load64(m.memory[int64(uint32(v3))+20:]))
			store64(m.memory[int64(uint32(v3))+88:], uint64(t11))
			v5 = i32(28)
			v1 = i32(1)
		l3:
			{
				m.fn885(v3+i32(116), v3+i32(88))
				t12 := int32(load32(m.memory[int64(uint32(v3))+116:]))
				if t12 == i32(-1) {
					t20 := int32(load32(m.memory[int64(uint32(v3))+88:]))
					t21 := int32(load32(m.memory[int64(uint32(v3))+92:]))
					m.fn44(t20, t21)
					t22 := int32(load32(m.memory[int64(uint32(v3))+56:]))
					if t22 == 0 {
						goto l4
					}
					t23 := m.fn113(i32(8), i32(32))
					v2 = t23
					store64(m.memory[int64(uint32(v2))+8:], uint64(i64(1)))
					store32(m.memory[uint32(v2):], uint32(i32(-0x7fffffff)))
					m.memory[int64(uint32(v2))+28] = byte(i32(0))
					t24 := int64(load64(m.memory[int64(uint32(v3))+48:]))
					store64(m.memory[int64(uint32(v2))+16:], uint64(t24))
					t25 := int32(load32(m.memory[int64(uint32(v3))+56:]))
					store32(m.memory[int64(uint32(v2))+24:], uint32(t25))
					v1 = i32(1)
					goto l5
				}
				{
					t13 := int32(load32(m.memory[int64(uint32(v3))+48:]))
					if v1 != t13 {
						goto l2
					}
					m.fn62(v3+i32(48), v1, i32(1), i32(4), i32(28))
					t14 := int32(load32(m.memory[int64(uint32(v3))+52:]))
					v4 = t14
				}
			l2:
				v2 = v4 + v5
				t15 := int64(load64(m.memory[int64(uint32(v3))+116:]))
				store64(m.memory[uint32(v2):], uint64(t15))
				t16 := int32(load32(m.memory[int64(uint32(v3))+140:]))
				store32(m.memory[int64(uint32(v2))+24:], uint32(t16))
				t17 := int64(load64(m.memory[int64(uint32(v3))+132:]))
				store64(m.memory[int64(uint32(v2))+16:], uint64(t17))
				t18 := int64(load64(m.memory[int64(uint32(v3))+124:]))
				store64(m.memory[int64(uint32(v2))+8:], uint64(t18))
				t19 := v3
				v1 = v1 + i32(1)
				store32(m.memory[int64(uint32(t19))+56:], uint32(v1))
				v5 = v5 + i32(28)
				goto l3
			}
		}
	l0:
		store32(m.memory[int64(uint32(v3))+56:], uint32(i32(0)))
		store64(m.memory[int64(uint32(v3))+48:], uint64(i64(0x400000000)))
		t26 := int32(load32(m.memory[int64(uint32(v3))+20:]))
		t27 := int32(load32(m.memory[int64(uint32(v3))+24:]))
		m.fn44(t26, t27)
	}
l4:
	m.fn971(v3 + i32(48))
	v2 = i32(8)
	v1 = i32(0)
l5:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v1))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v2))
	store32(m.memory[uint32(v0):], uint32(v1))
	m.g0 = v3 + i32(144)
}
func (m *Module) fn1364(v0, v1 int32) int32 {
	var v2 int32
	v2 = i32(0)
	{
		t0 := m.fn15(v0, v1, i32(1073751), i32(5))
		if t0 != 0 {
			goto l0
		}
		t1 := m.fn15(v0, v1, i32(1086307), i32(8))
		v2 = t1 ^ i32(1)
	}
l0:
	return v2
}
func (m *Module) fn1365(v0, v1, v2, v3, v4 int32) {
	var v5, v6, v7, v8, v9, v10 int32
	t0 := m.g0
	v5 = t0 - i32(96)
	m.g0 = v5
	t1 := int32(load32(m.memory[int64(uint32(v1))+20:]))
	t3 := v5
	p2 := i32(8)
	if uint32(v4) < uint32(i32(8)) {
		p2 = v4
	}
	v6 = p2
	v4 = t1 + v6*i32(24)
	t4 := int64(load64(m.memory[uint32(v4):]))
	store64(m.memory[int64(uint32(t3))+24:], uint64(t4))
	t5 := int64(load64(m.memory[int64(uint32(v4))+8:]))
	store64(m.memory[int64(uint32(v5))+32:], uint64(t5))
	t6 := int64(load64(m.memory[int64(uint32(v4))+16:]))
	store64(m.memory[int64(uint32(v5))+40:], uint64(t6))
	{
		t7 := int32(load32(m.memory[int64(uint32(v1))+36:]))
		v7 = t7
		if v7 == 0 {
			goto l0
		}
		{
			if v2 == 0 {
				goto l1
			}
			t8 := int32(load32(m.memory[int64(uint32(v2))+4:]))
			t9 := v5 + i32(72)
			t10 := v4
			t11 := v7
			v8 = t8
			t12 := int32(load32(m.memory[int64(uint32(v2))+8:]))
			t13 := v8
			v9 = t12
			t14 := m.fn1364(t13, v9)
			p15 := i32(0)
			if t14 != 0 {
				p15 = i32(216)
			}
			t16 := t11 + p15
			v10 = v6 * i32(24)
			m.fn1366(t9, t10, t16+v10)
			t17 := int64(load64(m.memory[int64(uint32(v5))+88:]))
			store64(m.memory[int64(uint32(v5))+40:], uint64(t17))
			t18 := int64(load64(m.memory[int64(uint32(v5))+80:]))
			store64(m.memory[int64(uint32(v5))+32:], uint64(t18))
			t19 := int64(load64(m.memory[int64(uint32(v5))+72:]))
			store64(m.memory[int64(uint32(v5))+24:], uint64(t19))
			m.fn855(v5+i32(16), v2+i32(12))
			t20 := int32(load32(m.memory[int64(uint32(v7))+652:]))
			t21 := int32(load32(m.memory[int64(uint32(v7))+656:]))
			t22 := int32(load32(m.memory[int64(uint32(v5))+16:]))
			t23 := int32(load32(m.memory[int64(uint32(v5))+20:]))
			t24 := m.fn1492(t20, t21, v8, v9, t22, t23)
			v4 = t24
			if v4 == 0 {
				goto l0
			}
			t25 := v5
			v4 = v4 + v10
			t26 := int64(load64(m.memory[int64(uint32(v4))+16:]))
			store64(m.memory[int64(uint32(t25))+88:], uint64(t26))
			t27 := int64(load64(m.memory[int64(uint32(v4))+8:]))
			store64(m.memory[int64(uint32(v5))+80:], uint64(t27))
			t28 := int64(load64(m.memory[uint32(v4):]))
			store64(m.memory[int64(uint32(v5))+72:], uint64(t28))
			m.fn1366(v5+i32(48), v5+i32(24), v5+i32(72))
			t29 := int64(load64(m.memory[int64(uint32(v5))+48:]))
			store64(m.memory[int64(uint32(v5))+24:], uint64(t29))
			t30 := int64(load64(m.memory[int64(uint32(v5))+56:]))
			store64(m.memory[int64(uint32(v5))+32:], uint64(t30))
			t31 := int64(load64(m.memory[int64(uint32(v5))+64:]))
			store64(m.memory[int64(uint32(v5))+40:], uint64(t31))
			goto l0
		}
	l1:
		m.fn1366(v5+i32(72), v4, v7+v6*i32(24)+i32(432))
		t32 := int64(load64(m.memory[int64(uint32(v5))+88:]))
		store64(m.memory[int64(uint32(v5))+40:], uint64(t32))
		t33 := int64(load64(m.memory[int64(uint32(v5))+80:]))
		store64(m.memory[int64(uint32(v5))+32:], uint64(t33))
		t34 := int64(load64(m.memory[int64(uint32(v5))+72:]))
		store64(m.memory[int64(uint32(v5))+24:], uint64(t34))
	}
l0:
	{
		if v2 == 0 {
			goto l2
		}
		t35 := int32(load32(m.memory[int64(uint32(v1))+32:]))
		v1 = t35
		if v1 == 0 {
			goto l2
		}
		m.fn855(v5+i32(8), v2+i32(12))
		t36 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t37 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		t38 := int32(load32(m.memory[int64(uint32(v2))+4:]))
		t39 := int32(load32(m.memory[int64(uint32(v2))+8:]))
		t40 := int32(load32(m.memory[int64(uint32(v5))+8:]))
		t41 := int32(load32(m.memory[int64(uint32(v5))+12:]))
		t42 := m.fn1492(t36, t37, t38, t39, t40, t41)
		v2 = t42
		if v2 == 0 {
			goto l2
		}
		t43 := v5
		v2 = v2 + v6*i32(24)
		t44 := int64(load64(m.memory[int64(uint32(v2))+16:]))
		store64(m.memory[int64(uint32(t43))+88:], uint64(t44))
		t45 := int64(load64(m.memory[int64(uint32(v2))+8:]))
		store64(m.memory[int64(uint32(v5))+80:], uint64(t45))
		t46 := int64(load64(m.memory[uint32(v2):]))
		store64(m.memory[int64(uint32(v5))+72:], uint64(t46))
		m.fn1366(v5+i32(48), v5+i32(24), v5+i32(72))
		t47 := int64(load64(m.memory[int64(uint32(v5))+48:]))
		store64(m.memory[int64(uint32(v5))+24:], uint64(t47))
		t48 := int64(load64(m.memory[int64(uint32(v5))+56:]))
		store64(m.memory[int64(uint32(v5))+32:], uint64(t48))
		t49 := int64(load64(m.memory[int64(uint32(v5))+64:]))
		store64(m.memory[int64(uint32(v5))+40:], uint64(t49))
	}
l2:
	m.fn1366(v0, v5+i32(24), v3+v6*i32(24))
	m.g0 = v5 + i32(96)
}
func (m *Module) fn1366(v0, v1, v2 int32) {
	var v3 int32
	t0 := int32(m.memory[uint32(v2)])
	t2 := v0
	p1 := v1
	if t0 != 0 {
		p1 = v2
	}
	v3 = p1
	t3 := int64(load64(m.memory[uint32(v3):]))
	store64(m.memory[uint32(t2):], uint64(t3))
	t4 := int64(load64(m.memory[int64(uint32(v3))+8:]))
	store64(m.memory[int64(uint32(v0))+8:], uint64(t4))
	t5 := int32(load32(m.memory[int64(uint32(v1))+16:]))
	t6 := int32(load32(m.memory[int64(uint32(v2))+16:]))
	t7 := fn1373(t5, t6)
	store32(m.memory[int64(uint32(v0))+16:], uint32(t7))
}
func (m *Module) fn1367(v0, v1, v2, v3, v4 int32) {
	var v5, v6, v7, v8, v9, v10, v11, v12, v13, v14 int32
	t0 := m.g0
	v5 = t0 - i32(128)
	m.g0 = v5
	store32(m.memory[int64(uint32(v5))+20:], uint32(i32(0)))
	store64(m.memory[int64(uint32(v5))+12:], uint64(i64(0x400000000)))
	store32(m.memory[int64(uint32(v5))+24:], uint32(v1))
	store32(m.memory[int64(uint32(v5))+28:], uint32(v1+v2*i32(44)))
	v6 = v5 + i32(100) + i32(4)
	v7 = v5 + i32(84) + i32(4)
	v8 = v5 + i32(44) + i32(4)
	t1 := int32(load32(m.memory[int64(uint32(v3))+28:]))
	v9 = t1
	t2 := int32(load32(m.memory[int64(uint32(v3))+12:]))
	v10 = t2
	t3 := int32(load32(m.memory[int64(uint32(v3))+8:]))
	v11 = t3
	t4 := int32(load32(m.memory[int64(uint32(v3))+4:]))
	v12 = t4
l1:
	{
		{
			{
				t5 := m.fn904(v5 + i32(24))
				v3 = t5
				if v3 == 0 {
					t59 := int32(load32(m.memory[int64(uint32(v5))+20:]))
					store32(m.memory[int64(uint32(v0))+8:], uint32(t59))
					t60 := int64(load64(m.memory[int64(uint32(v5))+12:]))
					store64(m.memory[uint32(v0):], uint64(t60))
					m.g0 = v5 + i32(128)
					return
				}
				t6 := int32(load32(m.memory[int64(uint32(v3))+36:]))
				v1 = t6
				if v1 == 0 {
					goto l1
				}
				t7 := int32(load32(m.memory[int64(uint32(v3))+40:]))
				t8 := m.fn1337(v1+i32(8), t7, i32(1074411), i32(53))
				if t8 != 0 {
					goto l1
				}
				{
					{
						t9 := int32(load32(m.memory[int64(uint32(v3))+4:]))
						v1 = t9
						t10 := int32(load32(m.memory[int64(uint32(v3))+8:]))
						t11 := v1
						v2 = t10
						t12 := m.fn15(t11, v2, i32(1072195), i32(1))
						if t12 != 0 {
							goto l2
						}
						t13 := m.fn15(v1, v2, i32(1083168), i32(3))
						if t13 == 0 {
							t42 := m.fn15(v1, v2, i32(1077123), i32(2))
							if t42 == 0 {
								goto l1
							}
							store32(m.memory[int64(uint32(v5))+100:], uint32(i32(8)))
							m.fn1340(v5+i32(12), v5+i32(100))
							goto l1
						}
					}
				l2:
					t14 := int32(load32(m.memory[uint32(v3+i32(28)):]))
					v13 = t14
					t15 := int32(load32(m.memory[uint32(v3+i32(32)):]))
					t16 := v13
					v14 = t15
					t17 := m.fn886(t16, v14, i32(1074411), i32(53), i32(1073717), i32(3))
					v1 = t17
					v2 = i32(0)
					v3 = i32(1)
					{
						{
							t18 := m.fn886(v13, v14, i32(1074411), i32(53), i32(1072196), i32(1))
							v13 = t18
							if v13 != 0 {
								goto l4
							}
							v13 = i32(0)
							goto l5
						}
					l4:
						t19 := int32(load32(m.memory[uint32(v13+i32(28)):]))
						t20 := int32(load32(m.memory[uint32(v13+i32(32)):]))
						m.fn864(v5+i32(100), t19, t20)
						v13 = i32(0)
						t21 := int32(load32(m.memory[int64(uint32(v5))+100:]))
						v14 = t21
						if v14 == i32(-1) {
							goto l5
						}
						t22 := int32(load32(m.memory[int64(uint32(v5))+108:]))
						v13 = t22
						t23 := int32(load32(m.memory[int64(uint32(v5))+104:]))
						v3 = t23
						v2 = v14
					}
				l5:
					m.fn865(v5+i32(32), v3, v13)
					m.fn16(v2, v3)
					t24 := int32(load32(m.memory[int64(uint32(v5))+40:]))
					if t24 == 0 {
						t43 := int32(load32(m.memory[int64(uint32(v5))+32:]))
						t44 := int32(load32(m.memory[int64(uint32(v5))+36:]))
						m.fn16(t43, t44)
						goto l1
					}
					if v1 == 0 {
						t57 := int32(load32(m.memory[int64(uint32(v5))+40:]))
						store32(m.memory[int64(uint32(v8))+8:], uint32(t57))
						t58 := int64(load64(m.memory[int64(uint32(v5))+32:]))
						store64(m.memory[uint32(v8):], uint64(t58))
						store32(m.memory[int64(uint32(v5))+60:], uint32(v4))
						store32(m.memory[int64(uint32(v5))+44:], uint32(i32(3)))
						goto l8
					}
					t25 := int32(load32(m.memory[uint32(v1+i32(16)):]))
					t26 := int32(load32(m.memory[uint32(v1+i32(20)):]))
					t27 := m.fn1493(t25, t26)
					v3 = t27
					t28 := int32(load32(m.memory[int64(uint32(v5))+40:]))
					store32(m.memory[int64(uint32(v8))+8:], uint32(t28))
					t29 := int64(load64(m.memory[int64(uint32(v5))+32:]))
					store64(m.memory[uint32(v8):], uint64(t29))
					t30 := fn1319(v3, v4)
					store32(m.memory[int64(uint32(v5))+60:], uint32(t30))
					store32(m.memory[int64(uint32(v5))+44:], uint32(i32(3)))
					t31 := int32(load32(m.memory[uint32(v1+i32(28)):]))
					t32 := int32(load32(m.memory[uint32(v1+i32(32)):]))
					t33 := m.fn886(t31, t32, i32(1074411), i32(53), i32(1074472), i32(10))
					v3 = t33
					if v3 == 0 {
						goto l8
					}
					t34 := int32(load32(m.memory[uint32(v3+i32(16)):]))
					t35 := int32(load32(m.memory[uint32(v3+i32(20)):]))
					m.fn845(v5, t34, t35, i32(1073159), i32(67), i32(1073226), i32(2))
					t36 := int32(load32(m.memory[uint32(v5):]))
					v3 = t36
					if v3 == 0 {
						goto l8
					}
					t37 := int32(load32(m.memory[int64(uint32(v5))+4:]))
					t38 := m.fn846(v12, v3, t37)
					v3 = t38
					if v3 == 0 {
						goto l8
					}
					t39 := int32(m.memory[int64(uint32(v3))+24])
					v1 = t39
					if v1 == 0 {
						t45 := int32(load32(m.memory[int64(uint32(v3))+4:]))
						t46 := v5 + i32(100)
						t47 := v11
						t48 := v10
						v13 = t45
						t49 := int32(load32(m.memory[int64(uint32(v3))+8:]))
						t50 := v13
						v2 = t49
						m.fn774(t46, t47, t48, t50, v2)
						{
							t51 := int32(load32(m.memory[int64(uint32(v5))+100:]))
							if t51 != 0 {
								m.fn781(v5 + i32(100))
								goto l10
							}
							t52 := m.fn1282(v9, v6)
							v3 = t52
							if v3 == 0 {
								goto l12
							}
							t53 := int32(load32(m.memory[uint32(v3+i32(4)):]))
							t54 := int32(load32(m.memory[uint32(v3+i32(8)):]))
							m.fn31(v7, t53, t54)
							m.fn784(v6)
							t55 := int64(load64(m.memory[uint32(v7):]))
							store64(m.memory[int64(uint32(v5))+72:], uint64(t55))
							t56 := int32(load32(m.memory[int64(uint32(v7))+8:]))
							store32(m.memory[int64(uint32(v5))+80:], uint32(t56))
							v1 = i32(2)
							goto l13
						}
					}
					t40 := int32(load32(m.memory[int64(uint32(v3))+8:]))
					v2 = t40
					t41 := int32(load32(m.memory[int64(uint32(v3))+4:]))
					v13 = t41
					goto l10
				}
			}
		l12:
			m.fn784(v6)
		l10:
			m.fn1494(v5+i32(84), v1, v13, v2)
			t61 := int64(load64(m.memory[uint32(v7):]))
			store64(m.memory[int64(uint32(v5))+72:], uint64(t61))
			t62 := int32(load32(m.memory[int64(uint32(v7))+8:]))
			store32(m.memory[int64(uint32(v5))+80:], uint32(t62))
			t63 := int32(load32(m.memory[int64(uint32(v5))+84:]))
			v1 = t63
			if v1 == i32(-1) {
				goto l8
			}
		}
	l13:
		t64 := int32(load32(m.memory[int64(uint32(v5))+80:]))
		store32(m.memory[int64(uint32(v6))+8:], uint32(t64))
		t65 := int64(load64(m.memory[int64(uint32(v5))+72:]))
		store64(m.memory[uint32(v6):], uint64(t65))
		t66 := m.fn113(i32(4), i32(28))
		v3 = t66
		t67 := int32(load32(m.memory[int64(uint32(v5))+68:]))
		store32(m.memory[int64(uint32(v3))+24:], uint32(t67))
		t68 := int64(load64(m.memory[int64(uint32(v5))+60:]))
		store64(m.memory[int64(uint32(v3))+16:], uint64(t68))
		t69 := int64(load64(m.memory[int64(uint32(v5))+52:]))
		store64(m.memory[int64(uint32(v3))+8:], uint64(t69))
		t70 := int64(load64(m.memory[int64(uint32(v5))+44:]))
		store64(m.memory[uint32(v3):], uint64(t70))
		store32(m.memory[int64(uint32(v5))+124:], uint32(i32(1)))
		store32(m.memory[int64(uint32(v5))+120:], uint32(v3))
		store32(m.memory[int64(uint32(v5))+116:], uint32(i32(1)))
		store32(m.memory[int64(uint32(v5))+100:], uint32(v1))
		m.fn1340(v5+i32(12), v5+i32(100))
		goto l1
	}
l8:
	m.fn1340(v5+i32(12), v5+i32(44))
	goto l1
}
func (m *Module) fn1368(v0, v1, v2 int32) {
	var v3, v4, v5, v6, v7, v8, v9 int32
	t0 := m.g0
	v3 = t0 - i32(16)
	m.g0 = v3
	store32(m.memory[int64(uint32(v3))+12:], uint32(v2))
	{
		t1 := m.fn823(v3+i32(12), i32(1287584))
		if t1 != 0 {
			goto l0
		}
		v4 = v1 * i32(28)
		v5 = v2 ^ i32(1)
		v1 = i32(0)
		var p2 int32
		if v2&i32(256) == 0 {
			p2 = 1
		}
		v6 = p2
		var p3 int32
		if v2&i32(65536) == 0 {
			p3 = 1
		}
		v7 = p3
	l4:
		if v4 == v1 {
			goto l0
		}
		{
			v8 = v0 + v1
			t4 := int32(load32(m.memory[uint32(v8):]))
			v9 = t4
			p5 := i32(1)
			if uint32(v9) > uint32(i32(2)) {
				p5 = v9 + i32(-3)
			}
			switch p5 {
			default:
				goto l3
			case 0:
				v9 = v8 + i32(16)
				t6 := int32(m.memory[uint32(v9)])
				m.memory[uint32(v9)] = byte(v5 & t6)
				v9 = v8 + i32(17)
				t7 := int32(m.memory[uint32(v9)])
				m.memory[uint32(v9)] = byte(v6 & t7)
				v8 = v8 + i32(18)
				t8 := int32(m.memory[uint32(v8)])
				m.memory[uint32(v8)] = byte(v7 & t8)
				goto l3
			case 1:
				t9 := int32(load32(m.memory[uint32(v8+i32(20)):]))
				t10 := int32(load32(m.memory[uint32(v8+i32(24)):]))
				m.fn1368(t9, t10, v2)
			}
		}
	l3:
		v1 = v1 + i32(28)
		goto l4
	}
l0:
	m.g0 = v3 + i32(16)
}
func (m *Module) fn1369(v0, v1 int32) {
	var v2 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		v2 = t0
		t1 := int32(load32(m.memory[uint32(v0):]))
		if v2 != t1 {
			goto l0
		}
		m.fn1147(v0)
	}
l0:
	t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	memory_copy(m.memory, uint32(t2+v2*i32(56)), uint32(v1), uint32(i32(56)))
	store32(m.memory[int64(uint32(v0))+8:], uint32(v2+i32(1)))
}
func (m *Module) fn1370(v0 int32) int32 {
	var v1, v2, v3, v4 int32
	{
		{
			t0 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			v1 = t0
			if v1 != 0 {
				goto l0
			}
			{
				t1 := int32(load32(m.memory[uint32(v0):]))
				v1 = t1
				t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
				if v1 != t2 {
					v2 = v1 + i32(1)
					goto l2
				}
				return i32(0)
			}
		}
	l0:
		store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
		t3 := int32(load32(m.memory[uint32(v0):]))
		v2 = t3
		v3 = v2 + v1
		t4 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		t5 := v3
		t6 := v1
		v4 = t4
		var p7 int32
		if uint32(t6) < uint32(v4-v2) {
			p7 = 1
		}
		v2 = p7
		p8 := i32(0)
		if v2 != 0 {
			p8 = t5
		}
		v1 = p8
		p9 := v4
		if v2 != 0 {
			p9 = v3 + i32(1)
		}
		v2 = p9
	}
l2:
	store32(m.memory[uint32(v0):], uint32(v2))
	return v1
}
func (m *Module) fn1371(v0, v1 int32) {
	var v2, v3, v4, v5, v6, v7, v8, v9, v10, v11 int32
	var v12 int64
	t0 := m.g0
	v2 = t0 - i32(48)
	m.g0 = v2
	t1 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	v3 = t1
	t2 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	t3 := v3
	v4 = t2
	p4 := v4
	if uint32(v3) > uint32(v4) {
		p4 = t3
	}
	v5 = p4
	t5 := int32(load32(m.memory[uint32(v1):]))
	v6 = t5
	{
		{
		l2:
			{
				{
					t6 := v5
					v7 = v3
					if t6 != v7 {
						goto l0
					}
					store32(m.memory[uint32(v0):], uint32(i32(-1)))
					goto l1
				}
			l0:
				t7 := v1
				v3 = v7 + i32(1)
				store32(m.memory[int64(uint32(t7))+8:], uint32(v3))
				{
					t8 := int32(m.memory[uint32(v6+v7)])
					v8 = t8
					switch v8 + i32(-10) {
					case 0, 3:
						goto l2
					case 1, 2:
						goto l3
					default:
						if v8 == 0 {
							goto l2
						}
					}
				}
			}
			switch v8 + i32(-123) {
			case 1:
				goto l3
			default:
				if v8 != i32(92) {
					goto l3
				}
				if uint32(v3) >= uint32(v4) {
					goto l8
				}
				v5 = v6 + v3
				t9 := int32(m.memory[uint32(v5)])
				v8 = t9
				if uint32((v8&i32(223)+i32(-65))&i32(255)) > uint32(i32(25)) {
					goto l9
				}
				v9 = v6 + v4
				v7 = v3
			l16:
				{
					if v4 != v7 {
						goto l10
					}
					v7 = v4
					v8 = v9
					goto l11
				l10:
					v8 = v6 + v7
					t10 := int32(m.memory[uint32(v8)])
					if uint32((t10&i32(223)+i32(-65))&i32(255)) < uint32(i32(26)) {
						t16 := v1
						v7 = v7 + i32(1)
						store32(m.memory[int64(uint32(t16))+8:], uint32(v7))
						goto l16
					}
				}
			l11:
				{
					if uint32(v7) < uint32(v3) {
						goto l13
					}
					if uint32(v7) > uint32(v4) {
						goto l13
					}
					m.fn12(v2+i32(16), v5, v7-v3)
					t11 := int32(load32(m.memory[int64(uint32(v2))+16:]))
					v3 = t11
					t12 := int32(load32(m.memory[int64(uint32(v2))+20:]))
					v5 = t12
					t13 := int32(load32(m.memory[int64(uint32(v2))+24:]))
					v10 = t13
					if uint32(v7) >= uint32(v4) {
						goto l14
					}
					if v8 == 0 {
						goto l14
					}
					t14 := int32(m.memory[uint32(v8)])
					if t14 != i32(45) {
						goto l14
					}
					v8 = i32(1)
					t15 := v1
					v7 = v7 + i32(1)
					store32(m.memory[int64(uint32(t15))+8:], uint32(v7))
					goto l15
				}
			l13:
				m.fn151(v3, v7, v4, i32(1075796))
				panic("unreachable")
			case 2:
				v10 = i32(3)
				goto l17
			case 0:
				v10 = i32(2)
			}
		l17:
			goto l18
		l3:
			v10 = i32(7)
			goto l18
		l14:
			v8 = i32(0)
		l15:
			p17 := v5
			if v3 != 0 {
				p17 = i32(1)
			}
			v9 = p17
			p18 := v10
			if v3 != 0 {
				p18 = i32(0)
			}
			v5 = p18
			p19 := v4
			if uint32(v7) > uint32(v4) {
				p19 = v7
			}
			v11 = p19
			v3 = v7
		l26:
			{
				if v11 == v3 {
					goto l19
				}
				t20 := int32(m.memory[uint32(v6+v3)])
				if uint32((t20+i32(-48))&i32(255)) < uint32(i32(10)) {
					t33 := v1
					v3 = v3 + i32(1)
					store32(m.memory[int64(uint32(t33))+8:], uint32(v3))
					goto l26
				}
				v11 = v3
			}
		l19:
			{
				if uint32(v11) > uint32(v7) {
					if uint32(v11) > uint32(v4) {
						m.fn151(v7, v11, v4, i32(1075812))
						panic("unreachable")
					}
					m.fn12(v2+i32(36), v6+v7, v11-v7)
					v10 = i32(1)
					t22 := int32(load32(m.memory[int64(uint32(v2))+40:]))
					t23 := int32(load32(m.memory[int64(uint32(v2))+36:]))
					t24 := v2 + i32(16)
					v7 = t23
					p25 := t22
					if v7 != 0 {
						p25 = i32(1108008)
					}
					t26 := int32(load32(m.memory[int64(uint32(v2))+44:]))
					p27 := t26
					if v7 != 0 {
						p27 = i32(1)
					}
					m.fn1322(t24, p25, p27)
					{
						t28 := int32(m.memory[int64(uint32(v2))+16])
						if t28 == 0 {
							t29 := int64(load64(m.memory[int64(uint32(v2))+24:]))
							v12 = t29
							p30 := v12
							if v8 != 0 {
								p30 = i64(0) - v12
							}
							v12 = p30
							p31 := i64(-0x80000000)
							if v12 > i64(-0x80000000) {
								p31 = v12
							}
							v12 = p31
							p32 := i64(0x7fffffff)
							if v12 < i64(0x7fffffff) {
								p32 = v12
							}
							v8 = int32(p32)
							goto l25
						}
						v10 = i32(0)
						goto l25
					}
				}
				v10 = i32(0)
				if v8 == 0 {
					goto l25
				}
				t21 := v1
				v11 = v11 + i32(-1)
				store32(m.memory[int64(uint32(t21))+8:], uint32(v11))
				goto l25
			}
		l25:
			{
				if uint32(v11) >= uint32(v4) {
					goto l27
				}
				t34 := int32(m.memory[uint32(v6+v11)])
				if t34 != i32(32) {
					goto l27
				}
				t35 := v1
				v11 = v11 + i32(1)
				store32(m.memory[int64(uint32(t35))+8:], uint32(v11))
			}
		l27:
			{
				t36 := m.fn15(v9, v5, i32(1075828), i32(3))
				if t36 != 0 {
					t38 := v4
					t39 := v11
					p37 := i32(0)
					if v8 > i32(0) {
						p37 = v8
					}
					v7 = t39 + p37
					p40 := v7
					if uint32(v7) < uint32(v11) {
						p40 = i32(-1)
					}
					p41 := v11
					if v10 != 0 {
						p41 = p40
					}
					v7 = p41
					p42 := v7
					if uint32(v4) < uint32(v7) {
						p42 = t38
					}
					v7 = p42
					if uint32(v4) < uint32(v11) {
						m.fn151(v11, v7, v4, i32(1075832))
						panic("unreachable")
					}
					store32(m.memory[int64(uint32(v1))+8:], uint32(v7))
					v9 = v7 - v11
					v10 = i32(8)
					v8 = v6 + v11
					v3 = int32(uint32(v8) >> 8)
					goto l18
				}
				v3 = int32(uint32(v8) >> 8)
				goto l18
			}
		}
	l9:
		t43 := v1
		v11 = v7 + i32(2)
		store32(m.memory[int64(uint32(t43))+8:], uint32(v11))
		v5 = i32(3)
		v10 = i32(0)
		v9 = i32(1075793)
		switch v8 + i32(-10) {
		case 0, 3:
			goto l18
		default:
			if v8 == i32(39) {
				goto l32
			}
			fallthrough
		case 1, 2:
			v10 = i32(5)
			goto l18
		}
	l32:
		if uint32(v4-v11) < uint32(i32(2)) {
			goto l8
		}
		t44 := v2 + i32(8)
		v3 = v6 + v11
		t45 := int32(m.memory[uint32(v3)])
		m.fn199(t44, t45, i32(16))
		t46 := int32(load32(m.memory[int64(uint32(v2))+12:]))
		v6 = t46
		t47 := int32(m.memory[int64(uint32(v3))+1])
		m.fn199(v2, t47, i32(16))
		v8 = i32(39)
		v10 = i32(7)
		{
			t48 := int32(load32(m.memory[int64(uint32(v2))+8:]))
			if t48 == i32(1) {
				goto l33
			}
			goto l18
		}
	l33:
		t49 := int32(load32(m.memory[uint32(v2):]))
		if t49&i32(1) == 0 {
			goto l18
		}
		t50 := int32(load32(m.memory[int64(uint32(v2))+4:]))
		v3 = t50
		store32(m.memory[int64(uint32(v1))+8:], uint32(v7+i32(4)))
		v8 = v3 + v6<<4
		v10 = i32(6)
	}
l18:
	store16(m.memory[int64(uint32(v0))+5:], uint16(v3))
	store32(m.memory[int64(uint32(v0))+12:], uint32(v5))
	store32(m.memory[int64(uint32(v0))+8:], uint32(v9))
	m.memory[int64(uint32(v0))+4] = byte(v8)
	store32(m.memory[uint32(v0):], uint32(v10))
	m.memory[uint32(v0+i32(7))] = byte(int32(uint32(v3) >> 16))
	goto l1
l8:
	store32(m.memory[uint32(v0):], uint32(i32(-1)))
l1:
	m.g0 = v2 + i32(48)
}
func (m *Module) fn1372(v0, v1, v2, v3, v4 int32) {
	var v5, v6, v7, v8, v9, v10, v11, v12, v13, v14, v15 int32
	t0 := m.g0
	v5 = t0 - i32(48)
	m.g0 = v5
	v6 = i32(0)
	store32(m.memory[int64(uint32(v5))+16:], uint32(i32(0)))
	store64(m.memory[int64(uint32(v5))+8:], uint64(i64(0x400000000)))
	store32(m.memory[int64(uint32(v5))+28:], uint32(i32(0)))
	store32(m.memory[int64(uint32(v5))+24:], uint32(v2))
	store32(m.memory[int64(uint32(v5))+20:], uint32(v1))
	v7 = i32(0)
	v8 = i32(0)
	v9 = i32(0)
	v10 = i32(0)
l13:
	{
		v11 = v6
		m.fn1371(v5+i32(32), v5+i32(20))
		{
			{
				t1 := int32(load32(m.memory[int64(uint32(v5))+32:]))
				v12 = t1
				if v12 == i32(-1) {
					t3 := int32(load32(m.memory[int64(uint32(v5))+16:]))
					store32(m.memory[int64(uint32(v0))+8:], uint32(t3))
					t4 := int64(load64(m.memory[int64(uint32(v5))+8:]))
					store64(m.memory[uint32(v0):], uint64(t4))
					m.g0 = v5 + i32(48)
					return
				}
				v6 = i32(0)
				p2 := v12 + i32(-2)
				if uint32(v12) < uint32(i32(2)) {
					p2 = i32(2)
				}
				switch p2 {
				case 0:
					v6 = i32(1)
					v13 = v9 + i32(1)
					goto l6
				case 1:
					v6 = i32(0)
					if v10&i32(1) != 0 {
						v10 = i32(1)
						if v9 != v14 {
							goto l11
						}
						if uint32(v7) < uint32(v15) {
							goto l12
						}
						if uint32(v7) > uint32(v2) {
							goto l12
						}
						m.fn1207(v5+i32(8), v1+v15, v7-v15)
						v8 = i32(0)
						goto l8
					}
					goto l8
				case 2:
					if v11 != i32(1) {
						goto l5
					}
					if v13 != v9 {
						goto l5
					}
					t5 := int32(load32(m.memory[int64(uint32(v5))+40:]))
					t6 := int32(load32(m.memory[int64(uint32(v5))+44:]))
					t7 := m.fn15(t5, t6, v3, v4)
					if t7 != 0 {
						if v8 != 0 {
							goto l6
						}
						v8 = i32(1)
						t9 := int32(load32(m.memory[int64(uint32(v5))+28:]))
						v15 = t9
						v14 = v13
						v9 = v13
						v10 = i32(1)
						goto l5
					}
					goto l6
				case 3:
					t8 := int32(m.memory[int64(uint32(v5))+36])
					if t8&i32(255) != i32(42) {
						goto l5
					}
					if v11&i32(1) != 0 {
						var p10 int32
						if v13 == v9 {
							p10 = 1
						}
						v6 = p10
						goto l5
					}
					goto l5
				default:
					goto l5
				}
			}
		l12:
			m.fn151(v15, v7, v2, i32(1085916))
			panic("unreachable")
		l6:
			v9 = v13
			goto l5
		l8:
			v10 = i32(0)
		l11:
			t11 := v9
			var p12 int32
			if v9 != i32(0) {
				p12 = 1
			}
			v9 = t11 - p12
		}
	l5:
		t13 := int32(load32(m.memory[int64(uint32(v5))+28:]))
		v7 = t13
		goto l13
	}
}
func fn1373(v0, v1 int32) int32 {
	var v2 int32
	t0 := int32(uint32(v0) >> 24)
	v2 = int32(uint32(v1) >> 24)
	p1 := v2
	if v2 == i32(2) {
		p1 = t0
	}
	t3 := p1 << 24
	p2 := v1
	if v1&i32(0xff0000) == i32(0x20000) {
		p2 = v0
	}
	t5 := t3 | p2&i32(0xff0000)
	p4 := v1
	if v1&i32(0xff00) == i32(512) {
		p4 = v0
	}
	t7 := t5 | p4&i32(0xff00)
	p6 := v1
	if v1&i32(255) == i32(2) {
		p6 = v0
	}
	return t7 | p6&i32(255)
}
func (m *Module) fn1374(v0, v1, v2, v3, v4, v5 int32) {
	var v6, v7, v8 int32
	t0 := m.g0
	v6 = t0 - i32(48)
	m.g0 = v6
	{
		if v2 == 0 {
			goto l0
		}
		t1 := int32(m.memory[uint32(v1)])
		v7 = t1
		v8 = i32(0)
		store32(m.memory[int64(uint32(v6))+20:], uint32(i32(0)))
		store64(m.memory[int64(uint32(v6))+12:], uint64(i64(0x400000000)))
		store32(m.memory[int64(uint32(v6))+32:], uint32(i32(0)))
		store64(m.memory[int64(uint32(v6))+24:], uint64(i64(0x100000000)))
		t2 := v7
		v2 = v2 + i32(-1)
		p3 := v2
		if uint32(v7) < uint32(v2) {
			p3 = t2
		}
		v7 = p3
		v1 = v1 + i32(1)
	l6:
		{
			if v7 == v8 {
				m.fn1478(v5, v6+i32(24), v6+i32(12))
				t5 := int32(load32(m.memory[int64(uint32(v6))+20:]))
				store32(m.memory[int64(uint32(v0))+8:], uint32(t5))
				t6 := int64(load64(m.memory[int64(uint32(v6))+12:]))
				store64(m.memory[uint32(v0):], uint64(t6))
				t7 := int32(load32(m.memory[int64(uint32(v6))+24:]))
				t8 := int32(load32(m.memory[int64(uint32(v6))+28:]))
				m.fn16(t7, t8)
				goto l4
			}
			t4 := int32(m.memory[uint32(v1+v8)])
			v2 = t4
			if uint32(v2) < uint32(i32(9)) {
				t9 := m.fn1432(v3, v4, v8+i32(1))
				if t9 == 0 {
					goto l3
				}
				m.fn1478(v5, v6+i32(24), v6+i32(12))
				store32(m.memory[int64(uint32(v6))+36:], uint32(i32(-1)))
				m.memory[int64(uint32(v6))+40] = byte(v2)
				m.fn1321(v6+i32(12), v6+i32(36))
				goto l5
			}
			goto l3
		}
	l3:
		m.fn145(v6+i32(24), v2)
	l5:
		v8 = v8 + i32(1)
		goto l6
	}
l0:
	store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
	store64(m.memory[uint32(v0):], uint64(i64(0x400000000)))
l4:
	m.g0 = v6 + i32(48)
}
func fn1375(v0 int32) int32 {
	var v1 int32
	switch v0 + i32(-1) {
	default:
		if v0 == i32(23) {
			goto l5
		}
		v1 = i32(1)
		if v0 != i32(255) {
			goto l6
		}
		return i32(255)
	case 0:
		return i32(5)
	case 1:
		return i32(4)
	case 2:
		return i32(3)
	case 3:
		return i32(2)
	}
l5:
	v1 = i32(0)
l6:
	return v1
}
func (m *Module) fn1376(v0, v1 int32) {
	t0 := int32(m.memory[int64(uint32(v0))+25])
	switch t0 {
	case 0:
		v0 = v0 + i32(12)
		fallthrough
	default:
		m.fn145(v0, v1)
		fallthrough
	case 2:
	}
}
func (m *Module) fn1377(v0 int32) {
	t0 := int32(load32(m.memory[uint32(v0):]))
	if t0 == i32(-1) {
		return
	}
	m.fn764(v0)
}
func (m *Module) fn1378(v0 int32) {
	var v1 int32
	{
		t0 := int32(load32(m.memory[uint32(v0):]))
		v1 = t0
		if v1 == i32(-1) {
			return
		}
		t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		m.fn16(v1, t1)
		t2 := int32(load32(m.memory[int64(uint32(v0))+12:]))
		t3 := int32(load32(m.memory[int64(uint32(v0))+16:]))
		m.fn16(t2, t3)
	}
}
func (m *Module) fn1379(v0, v1, v2, v3, v4, v5 int32) {
	var v6, v7, v8, v9, v10, v11, v12, v13 int32
	t0 := m.g0
	v6 = t0 - i32(832)
	m.g0 = v6
	t1 := int32(load32(m.memory[uint32(v1):]))
	v7 = t1
	v8 = i32(0)
	store32(m.memory[uint32(v1):], uint32(i32(0)))
	t2 := int32(load32(m.memory[uint32(v2):]))
	v9 = t2
	store32(m.memory[uint32(v2):], uint32(i32(0)))
	{
		if v9&v7 == 0 {
			goto l11
		}
		t3 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		v10 = t3
		{
			t4 := int32(load32(m.memory[uint32(v0):]))
			t5 := int32(load32(m.memory[int64(uint32(v2))+4:]))
			t6 := m.fn1479(t4, t5)
			v1 = t6
			if v1 == 0 {
				goto l1
			}
			m.fn1137(v6, v1)
			t7 := int32(load32(m.memory[uint32(v6):]))
			if t7 != i32(-1) {
				goto l2
			}
		}
	l1:
		v2 = i32(-288)
	l4:
		if v2 == 0 {
			goto l3
		}
		v1 = v6 + i32(288) + v2
		m.memory[uint32(v1+i32(312))] = byte(i32(0))
		store64(m.memory[uint32(v1+i32(304)):], uint64(i64(1)))
		m.memory[uint32(v1+i32(300))] = byte(i32(0))
		store32(m.memory[uint32(v1+i32(296)):], uint32(i32(0)))
		store64(m.memory[uint32(v1+i32(288)):], uint64(i64(0x400000000)))
		v2 = v2 + i32(32)
		goto l4
	l2:
		memory_copy(m.memory, uint32(v6+i32(288)), uint32(v6), uint32(i32(288)))
	l3:
		t8 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		v11 = t8
		v2 = i32(0)
		v1 = i32(-252)
		v7 = v3
		v9 = v4
	l9:
		{
			if v1 != 0 {
				{
					t10 := int32(m.memory[uint32(v9)])
					v12 = t10
					if v12 == i32(254) {
						goto l6
					}
					m.memory[uint32(v6+i32(288)+v2+i32(24))] = byte(v12)
				}
			l6:
				{
					t11 := int64(load64(m.memory[uint32(v7):]))
					if t11 != i64(1) {
						goto l7
					}
					t12 := int64(load64(m.memory[uint32(v7+i32(8)):]))
					store64(m.memory[uint32(v6+i32(288)+v2+i32(16)):], uint64(t12))
				}
			l7:
				{
					v12 = v5 + v1
					t13 := int32(load32(m.memory[uint32(v12+i32(252)):]))
					if t13 == i32(-1) {
						goto l8
					}
					v13 = v6 + i32(288) + v2
					t14 := int32(m.memory[uint32(v13+i32(24))])
					if uint32((t14+i32(-1))&i32(255)) > uint32(i32(253)) {
						goto l8
					}
					t15 := int32(load32(m.memory[uint32(v12+i32(256)):]))
					t16 := int32(load32(m.memory[uint32(v12+i32(260)):]))
					t17 := int32(load32(m.memory[uint32(v12+i32(268)):]))
					t18 := int32(load32(m.memory[uint32(v12+i32(272)):]))
					t19 := int32(load32(m.memory[uint32(v11):]))
					m.fn1374(v6, t15, t16, t17, t18, t19)
					t20 := int32(m.memory[uint32(v12+i32(276))])
					v12 = t20
					m.fn763(v13)
					m.memory[uint32(v13+i32(12))] = byte(v12)
					t21 := int32(load32(m.memory[int64(uint32(v6))+8:]))
					store32(m.memory[int64(uint32(v13))+8:], uint32(t21))
					t22 := int64(load64(m.memory[uint32(v6):]))
					store64(m.memory[uint32(v13):], uint64(t22))
				}
			l8:
				v1 = v1 + i32(28)
				v2 = v2 + i32(32)
				v7 = v7 + i32(16)
				v9 = v9 + i32(1)
				goto l9
			}
			t9 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			m.fn1124(v6, t9, v10, v6+i32(288))
			m.fn1377(v6)
			goto l11
		}
	}
l11:
	if v8 == i32(144) {
		goto l10
	}
	store64(m.memory[uint32(v3+v8):], uint64(i64(0)))
	v8 = v8 + i32(16)
	goto l11
l10:
	m.memory[int64(uint32(v4))+8] = byte(i32(-2))
	store64(m.memory[uint32(v4):], uint64(i64(-72340172838076674)))
	v8 = i32(0)
l13:
	if v8 == i32(252) {
		goto l12
	}
	store32(m.memory[uint32(v6+i32(580)+v8):], uint32(i32(-1)))
	v8 = v8 + i32(28)
	goto l13
l12:
	m.fn1380(v5)
	memory_copy(m.memory, uint32(v5), uint32(v6+i32(580)), uint32(i32(252)))
	m.g0 = v6 + i32(832)
}
func (m *Module) fn1380(v0 int32) {
	var v1 int32
	v1 = i32(0)
l1:
	if v1 == i32(252) {
		return
	}
	m.fn1378(v0 + v1)
	v1 = v1 + i32(28)
	goto l1
}
func (m *Module) fn1381(v0 int32) {
	var v1, v2, v3, v4 int32
	var v5 int64
	var v6 int32
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
					m.fn764(v6 + (i32(0)-int32(uint32(int64(bits.TrailingZeros64(uint64(v5))))>>3))*i32(296) + i32(-288))
					v4 = v4 + i32(-1)
					v5 = (v5 + i64(-1)) & v5
					goto l4
				}
				v6 = v6 + i32(-2368)
				t5 := int64(load64(m.memory[uint32(v0):]))
				v5 = (t5 ^ i64(-1)) & i64(-0x7f7f7f7f7f7f7f80)
				v0 = v0 + i32(8)
				goto l3
			}
		}
	l1:
		m.fn39(v1+i32(4), i32(296), i32(8), v2+i32(1))
		t6 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		t7 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t8 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		m.fn40(v3-t6, t7, t8)
	}
l0:
	m.g0 = v1 + i32(16)
}
func (m *Module) fn1382(v0 int32) {
	var v1, v2, v3 int32
	var v4 int64
	var v5, v6 int32
	var v7 int64
	var v8, v9 int32
	var v10 int64
	t0 := m.g0
	v1 = t0 - i32(48)
	m.g0 = v1
	v2 = i32(0)
	{
		t1 := int32(load32(m.memory[int64(uint32(v0))+416:]))
		if t1 != i32(1) {
			goto l0
		}
		t2 := int32(load32(m.memory[int64(uint32(v0))+12:]))
		if t2 == 0 {
			goto l0
		}
		t3 := int64(load64(m.memory[int64(uint32(v0))+16:]))
		t4 := int64(load64(m.memory[int64(uint32(v0))+24:]))
		t5 := int32(load32(m.memory[int64(uint32(v0))+420:]))
		v3 = t5
		t6 := m.fn66(t3, t4, v3)
		v4 = t6
		t7 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		v5 = t7
		v6 = v5 & int32(v4)
		v7 = int64(uint64(v4)>>25) & i64(127) * i64(72340172838076673)
		t8 := int32(load32(m.memory[uint32(v0):]))
		v8 = t8
		v9 = i32(0)
	l4:
		{
			t9 := int64(load64(m.memory[uint32(v8+v6):]))
			v10 = t9
			v4 = v10 ^ v7
			v4 = (v4 ^ i64(-1)) & (v4 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
		l3:
			{
				if v4 == 0 {
					v2 = i32(0)
					if !(v10&(v10<<1)&i64(-0x7f7f7f7f7f7f7f80) == 0) {
						goto l0
					}
					t12 := v6
					v9 = v9 + i32(8)
					v6 = (t12 + v9) & v5
					goto l4
				}
				t10 := v3
				v2 = v8 - (int32(uint32(int64(bits.TrailingZeros64(uint64(v4))))>>3)+v6)&v5<<3
				t11 := int32(load32(m.memory[uint32(v2+i32(-8)):]))
				if t10 == t11 {
					goto l2
				}
				v4 = (v4 + i64(-1)) & v4
				goto l3
			}
		l2:
		}
		t13 := int32(load32(m.memory[uint32(v2+i32(-4)):]))
		v2 = t13
	}
l0:
	{
		t14 := int32(load32(m.memory[int64(uint32(v0))+304:]))
		v8 = t14
		if v8 == 0 {
			goto l5
		}
		store32(m.memory[int64(uint32(v0))+304:], uint32(i32(0)))
		t15 := int32(load32(m.memory[int64(uint32(v0))+296:]))
		v5 = t15
		t16 := int32(load32(m.memory[int64(uint32(v0))+300:]))
		v6 = t16
		store64(m.memory[int64(uint32(v0))+296:], uint64(i64(0x100000000)))
		t17 := int32(load32(m.memory[int64(uint32(v0))+312:]))
		t19 := v1 + i32(28)
		p18 := t17
		if v2 != 0 {
			p18 = v2
		}
		m.fn510(t19, p18, v6, v8)
		t20 := int32(load32(m.memory[int64(uint32(v1))+36:]))
		store32(m.memory[int64(uint32(v1))+24:], uint32(t20))
		t21 := int64(load64(m.memory[int64(uint32(v1))+28:]))
		store64(m.memory[int64(uint32(v1))+16:], uint64(t21))
		m.fn490(v1+i32(4), v1+i32(16))
		m.fn16(v5, v6)
		t22 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		if t22 == i32(-1) {
			goto l5
		}
		m.fn1384(v0, v1+i32(4))
	}
l5:
	m.g0 = v1 + i32(48)
}
func (m *Module) fn1383(v0, v1 int32) {
	var v2, v3 int32
	var v4 int64
	var v5, v6, v7, v8, v9, v10, v11, v12, v13, v14, v15, v16 int32
	var v17 int64
	t0 := m.g0
	v2 = t0 - i32(160)
	m.g0 = v2
	t1 := int32(load32(m.memory[int64(uint32(v1))+388:]))
	v3 = t1
	store32(m.memory[int64(uint32(v1))+388:], uint32(i32(0)))
	t2 := int64(load64(m.memory[int64(uint32(v1))+380:]))
	v4 = t2
	store64(m.memory[int64(uint32(v1))+380:], uint64(i64(0x400000000)))
	t3 := int32(load32(m.memory[int64(uint32(v1))+240:]))
	v5 = t3
	store32(m.memory[int64(uint32(v1))+240:], uint32(i32(-1)))
	store32(m.memory[int64(uint32(v2))+16:], uint32(v3))
	store64(m.memory[int64(uint32(v2))+8:], uint64(v4))
	t4 := int32(load32(m.memory[int64(uint32(v1))+244:]))
	v3 = t4
	{
		{
			t5 := int32(m.memory[int64(uint32(v1))+454])
			if t5 != 0 {
				t12 := int32(load32(m.memory[int64(uint32(v1))+444:]))
				t13 := v1 + i32(332)
				v7 = t12
				p14 := i32(1)
				if uint32(v7) > uint32(i32(1)) {
					p14 = v7
				}
				t15 := int32(m.memory[int64(uint32(v1))+456])
				m.fn1480(t13, p14, t15, v2+i32(8))
				store32(m.memory[uint32(v0):], uint32(i32(-1)))
				goto l4
			}
			t6 := int32(load32(m.memory[int64(uint32(v1))+248:]))
			v6 = t6
			m.fn1394(v2+i32(88), v1)
			{
				t7 := int32(load32(m.memory[int64(uint32(v2))+88:]))
				v7 = t7
				if v7 == i32(-1) {
					t11 := int32(m.memory[int64(uint32(v1))+456])
					v7 = t11
					if v7 == i32(2) {
						goto l3
					}
					v6 = v1 + i32(392)
					m.fn1351(v6, v1+i32(404))
					m.fn1445(v1+i32(176), v7&i32(1), v2+i32(8), v6)
					store32(m.memory[uint32(v0):], uint32(i32(-1)))
					goto l4
				}
				t8 := int32(load32(m.memory[int64(uint32(v2))+108:]))
				store32(m.memory[int64(uint32(v0))+20:], uint32(t8))
				t9 := int64(load64(m.memory[int64(uint32(v2))+100:]))
				store64(m.memory[int64(uint32(v0))+12:], uint64(t9))
				t10 := int64(load64(m.memory[int64(uint32(v2))+92:]))
				store64(m.memory[int64(uint32(v0))+4:], uint64(t10))
				goto l2
			}
		}
	l3:
		{
			t16 := int32(load32(m.memory[int64(uint32(v2))+12:]))
			t17 := int32(load32(m.memory[int64(uint32(v2))+16:]))
			t18 := m.fn23(t16, t17)
			if t18 != 0 {
				goto l5
			}
			{
				{
					t19 := int32(load32(m.memory[int64(uint32(v1))+424:]))
					if t19 != i32(1) {
						t26 := int32(m.memory[int64(uint32(v1))+459])
						v11 = t26
						if v11 != 0 {
							goto l9
						}
						v4 = i64(0)
						goto l10
					}
					t20 := int32(load32(m.memory[int64(uint32(v1))+448:]))
					v8 = t20
					t21 := int32(load32(m.memory[int64(uint32(v1))+428:]))
					t22 := v1 + i32(64)
					v9 = t21
					t23 := m.fn1479(t22, v9)
					v10 = t23
					if v10 != 0 {
						v7 = i32(-1)
						{
							t28 := v10
							p27 := i32(8)
							if uint32(v8) < uint32(i32(8)) {
								p27 = v8
							}
							v6 = p27
							v12 = t28 + v6<<5
							t29 := int32(m.memory[int64(uint32(v12))+24])
							v11 = t29
							if v11 != 0 {
								if v11 != i32(255) {
									v13 = v1 + i32(96)
									t30 := int64(load64(m.memory[int64(uint32(v12))+16:]))
									t31 := m.fn1482(v13, v9, v6, t30)
									v4 = t31
									m.fn1483(v2+i32(88), v13, v9)
									t32 := m.fn1484(v2 + i32(88))
									v14 = t32
									{
										t33 := int32(load32(m.memory[int64(uint32(v12))+8:]))
										v6 = t33
										if v6 != 0 {
											v15 = v14 + i32(72)
											store32(m.memory[int64(uint32(v2))+152:], uint32(i32(0)))
											store64(m.memory[int64(uint32(v2))+144:], uint64(i64(0x100000000)))
											v6 = v6 * i32(12)
											t34 := int32(m.memory[int64(uint32(v12))+12])
											v16 = t34
											t35 := int32(load32(m.memory[int64(uint32(v12))+4:]))
											v7 = t35
										l21:
											{
												{
													if v6 == 0 {
														m.fn800(v2+i32(88), v11, v4)
														t50 := int32(load32(m.memory[int64(uint32(v2))+148:]))
														v12 = t50
														t51 := int32(load32(m.memory[int64(uint32(v2))+152:]))
														t52 := int32(load32(m.memory[int64(uint32(v2))+92:]))
														t53 := v12
														v7 = t52
														t54 := int32(load32(m.memory[int64(uint32(v2))+96:]))
														t55 := m.fn191(t53, t51, v7, t54)
														v6 = t55
														t56 := int32(load32(m.memory[int64(uint32(v2))+88:]))
														m.fn16(t56, v7)
														t57 := int32(load32(m.memory[int64(uint32(v2))+144:]))
														v7 = t57
														{
															if v6 != 0 {
																m.fn16(v7, v12)
																v7 = i32(-1)
																goto l12
															}
															t58 := int64(load64(m.memory[int64(uint32(v2))+148:]))
															v17 = t58
															goto l12
														}
													}
													t36 := int32(load32(m.memory[uint32(v7):]))
													if t36 != i32(-1) {
														goto l17
													}
													t37 := int32(m.memory[uint32(v7+i32(4))])
													v12 = t37
													p38 := i32(8)
													if uint32(v12) < uint32(i32(8)) {
														p38 = v12
													}
													v12 = p38
													v13 = i32(1)
													{
														if v16&i32(1) != 0 {
															goto l18
														}
														t39 := int32(m.memory[int64(uint32(v10+v12<<5))+24])
														v13 = t39
														p40 := v13
														if v13 == i32(255) {
															p40 = i32(1)
														}
														v13 = p40
													}
												l18:
													t41 := int32(m.memory[uint32(v15+v12)])
													t43 := v2 + i32(88)
													t44 := v13
													p42 := v10 + v12<<5 + i32(16)
													if t41 != 0 {
														p42 = v14 + v12<<3
													}
													t45 := int64(load64(m.memory[uint32(p42):]))
													m.fn804(t43, t44, t45)
													t46 := int32(load32(m.memory[int64(uint32(v2))+92:]))
													t47 := v2 + i32(144)
													v12 = t46
													t48 := int32(load32(m.memory[int64(uint32(v2))+96:]))
													m.fn75(t47, v12, t48)
													t49 := int32(load32(m.memory[int64(uint32(v2))+88:]))
													m.fn16(t49, v12)
													goto l19
												}
											l17:
												t59 := int32(load32(m.memory[uint32(v7+i32(4)):]))
												t60 := int32(load32(m.memory[uint32(v7+i32(8)):]))
												m.fn75(v2+i32(144), t59, t60)
											}
										l19:
											v7 = v7 + i32(12)
											v6 = v6 + i32(-12)
											goto l21
										}
										goto l12
									}
								}
								store32(m.memory[int64(uint32(v2))+52:], uint32(i32(-2)))
								goto l14
							}
							v4 = i64(0)
							goto l12
						}
					}
					t25 := v2 + i32(52)
					p24 := v3
					if v5 == i32(-1) {
						p24 = i32(0)
					}
					m.fn1481(t25, p24, v6)
					store32(m.memory[int64(uint32(v2))+48:], uint32(v8))
					v4 = i64(0)
					store64(m.memory[int64(uint32(v2))+40:], uint64(i64(0)))
					v11 = i32(0)
					goto l8
				}
			l9:
				if v11 == i32(255) {
					{
						if v5 == i32(-1) {
							goto l30
						}
						m.fn46(v2, v3, v6)
						t78 := int32(load32(m.memory[int64(uint32(v2))+4:]))
						if t78 != 0 {
							t79 := int32(load32(m.memory[int64(uint32(v1))+448:]))
							v8 = t79
							m.fn1481(v2+i32(52), v3, v6)
							store32(m.memory[int64(uint32(v2))+48:], uint32(v8))
							v4 = i64(0)
							store64(m.memory[int64(uint32(v2))+40:], uint64(i64(0)))
							v11 = i32(0)
							m.memory[int64(uint32(v2))+32] = byte(i32(0))
							v17 = i64(-2)
							goto l32
						}
					}
				l30:
					store32(m.memory[int64(uint32(v2))+52:], uint32(i32(-2)))
					goto l14
				}
				if v5 == i32(-1) {
					goto l23
				}
				store32(m.memory[int64(uint32(v2))+152:], uint32(i32(0)))
				store64(m.memory[int64(uint32(v2))+144:], uint64(i64(0x100000000)))
				m.fn47(v2+i32(144), i32(0))
				m.memory[int64(uint32(v2))+96] = byte(i32(0))
				store32(m.memory[int64(uint32(v2))+88:], uint32(v3))
				store32(m.memory[int64(uint32(v2))+92:], uint32(v3+v6))
			l25:
				{
					t61 := m.fn48(v2 + i32(88))
					v7 = t61
					if uint32(v7+i32(-58)) < uint32(i32(-10)) {
						t62 := int32(load32(m.memory[int64(uint32(v2))+144:]))
						v7 = t62
						t63 := int32(load32(m.memory[int64(uint32(v2))+148:]))
						t64 := v2 + i32(88)
						v6 = t63
						t65 := int32(load32(m.memory[int64(uint32(v2))+152:]))
						m.fn1190(t64, v6, t65)
						{
							{
								t66 := int32(m.memory[int64(uint32(v2))+88])
								v12 = t66
								if v12 != i32(1) {
									goto l26
								}
								goto l27
							}
						l26:
							t67 := int64(load64(m.memory[int64(uint32(v2))+96:]))
							v4 = t67
							p68 := i64(0xffffffff)
							if uint64(v4) < uint64(i64(0xffffffff)) {
								p68 = v4
							}
							v4 = p68
						}
					l27:
						m.fn16(v7, v6)
						if v12 != 0 {
							goto l23
						}
						t69 := int32(load32(m.memory[int64(uint32(v1))+448:]))
						v7 = t69
						m.fn1483(v2+i32(88), v1+i32(96), i32(0x7fffffff))
						t70 := m.fn1484(v2 + i32(88))
						v6 = t70
						t72 := v6
						p71 := i32(8)
						if uint32(v7) < uint32(i32(8)) {
							p71 = v7
						}
						v7 = p71
						store64(m.memory[uint32(t72+v7<<3):], uint64(v4))
						v12 = v6 + i32(72)
						m.memory[uint32(v12+v7)] = byte(i32(1))
						store32(m.memory[int64(uint32(v2))+96:], uint32(v7+i32(1)))
						store32(m.memory[int64(uint32(v2))+92:], uint32(v6+i32(81)))
						store32(m.memory[int64(uint32(v2))+88:], uint32(v12))
					l28:
						{
							t73 := m.fn1370(v2 + i32(88))
							v7 = t73
							if v7 == 0 {
								goto l10
							}
							m.memory[uint32(v7)] = byte(i32(0))
							goto l28
						}
					}
					m.fn74(v2+i32(144), v7)
					goto l25
				}
			l23:
				t74 := int32(load32(m.memory[int64(uint32(v1))+448:]))
				t75 := m.fn1482(v1+i32(96), i32(0x7fffffff), t74, i64(1))
				v4 = t75
			}
		l10:
			store32(m.memory[int64(uint32(v2))+52:], uint32(i32(-1)))
			store64(m.memory[int64(uint32(v2))+40:], uint64(v4))
			m.memory[int64(uint32(v2))+32] = byte(v11)
			v17 = i64(-1)
			store64(m.memory[int64(uint32(v2))+24:], uint64(i64(-1)))
			t76 := int32(load32(m.memory[int64(uint32(v1))+448:]))
			t77 := v2
			v8 = t76
			store32(m.memory[int64(uint32(t77))+48:], uint32(v8))
			goto l14
		}
	l5:
		m.fn1396(v1)
		v7 = i32(-1)
	l2:
		store32(m.memory[uint32(v0):], uint32(v7))
		m.fn134(v5, v3)
		m.fn894(v2 + i32(8))
		goto l29
	l12:
		store64(m.memory[int64(uint32(v2))+56:], uint64(v17))
		store32(m.memory[int64(uint32(v2))+52:], uint32(v7))
		store32(m.memory[int64(uint32(v2))+48:], uint32(v8))
		store64(m.memory[int64(uint32(v2))+40:], uint64(v4))
	l8:
		m.memory[int64(uint32(v2))+32] = byte(v11)
		v17 = int64(v9)
	l32:
		store64(m.memory[int64(uint32(v2))+24:], uint64(v17))
	l14:
		{
			t80 := int32(m.memory[int64(uint32(v1))+452])
			if t80 == 0 {
				t94 := int32(load32(m.memory[int64(uint32(v2))+52:]))
				if t94 == i32(-2) {
					m.fn1396(v1)
					t101 := int32(load32(m.memory[int64(uint32(v2))+16:]))
					store32(m.memory[int64(uint32(v2))+100:], uint32(t101))
					t102 := int64(load64(m.memory[int64(uint32(v2))+8:]))
					store64(m.memory[int64(uint32(v2))+92:], uint64(t102))
					store32(m.memory[int64(uint32(v2))+88:], uint32(i32(-0x80000000)))
					m.fn338(v1+i32(392), v2+i32(88))
					store32(m.memory[uint32(v0):], uint32(i32(-1)))
					goto l4
				}
				t95 := v2
				v7 = v2 + i32(52)
				t96 := int32(load32(m.memory[int64(uint32(v7))+8:]))
				store32(m.memory[int64(uint32(t95))+124:], uint32(t96))
				t97 := int64(load64(m.memory[uint32(v7):]))
				store64(m.memory[int64(uint32(v2))+116:], uint64(t97))
				t98 := m.fn113(i32(8), i32(32))
				v7 = t98
				store32(m.memory[uint32(v7):], uint32(i32(-0x80000000)))
				t99 := int64(load64(m.memory[int64(uint32(v2))+8:]))
				store64(m.memory[int64(uint32(v7))+4:], uint64(t99))
				t100 := int32(load32(m.memory[int64(uint32(v2))+16:]))
				store32(m.memory[int64(uint32(v7))+12:], uint32(t100))
				store32(m.memory[int64(uint32(v2))+112:], uint32(v8))
				store32(m.memory[int64(uint32(v2))+136:], uint32(i32(1)))
				store32(m.memory[int64(uint32(v2))+132:], uint32(v7))
				store32(m.memory[int64(uint32(v2))+128:], uint32(i32(1)))
				store64(m.memory[int64(uint32(v2))+104:], uint64(v4))
				m.memory[int64(uint32(v2))+96] = byte(v11)
				store64(m.memory[int64(uint32(v2))+88:], uint64(v17))
				m.fn1369(v1+i32(404), v2+i32(88))
				store32(m.memory[uint32(v0):], uint32(i32(-1)))
				goto l4
			}
			t81 := int32(m.memory[int64(uint32(v1))+453])
			v7 = t81
			m.fn1396(v1)
			t82 := int64(load64(m.memory[int64(uint32(v2))+8:]))
			store64(m.memory[int64(uint32(v2))+64:], uint64(t82))
			t83 := int32(load32(m.memory[int64(uint32(v2))+16:]))
			t84 := v2
			v6 = t83
			store32(m.memory[int64(uint32(t84))+72:], uint32(v6))
			t85 := int32(load32(m.memory[int64(uint32(v2))+68:]))
			t86 := int32(load32(m.memory[int64(uint32(v1))+436:]))
			m.fn1368(t85, v6, t86)
			t87 := int32(load32(m.memory[int64(uint32(v2))+52:]))
			v12 = t87
			var p88 int32
			if v12 == i32(-2) {
				p88 = 1
			}
			v6 = p88
			if v6 != 0 {
				goto l34
			}
			if v11&i32(255) == 0 {
				goto l34
			}
			m.fn225(v2+i32(88), v2+i32(52))
			{
				t89 := int32(load32(m.memory[int64(uint32(v2))+88:]))
				if t89 == i32(-1) {
					goto l35
				}
				t90 := int32(load32(m.memory[int64(uint32(v2))+96:]))
				store32(m.memory[int64(uint32(v2))+152:], uint32(t90))
				t91 := int64(load64(m.memory[int64(uint32(v2))+88:]))
				store64(m.memory[int64(uint32(v2))+144:], uint64(t91))
				goto l36
			}
		l35:
			m.fn800(v2+i32(144), v11, v4)
		l36:
			store32(m.memory[int64(uint32(v2))+84:], uint32(i32(25)))
			store32(m.memory[int64(uint32(v2))+80:], uint32(v2+i32(144)))
			m.fn73(v2+i32(92), i32(1070105), v2+i32(80))
			t92 := int32(load32(m.memory[int64(uint32(v2))+144:]))
			t93 := int32(load32(m.memory[int64(uint32(v2))+148:]))
			m.fn16(t92, t93)
			store32(m.memory[int64(uint32(v2))+104:], uint32(i32(0)))
			store32(m.memory[int64(uint32(v2))+88:], uint32(i32(3)))
			m.fn1163(v2+i32(64), v2+i32(88))
			goto l34
		}
	l34:
		t103 := int32(load32(m.memory[int64(uint32(v2))+72:]))
		store32(m.memory[int64(uint32(v2))+96:], uint32(t103))
		t104 := int64(load64(m.memory[int64(uint32(v2))+64:]))
		store64(m.memory[int64(uint32(v2))+88:], uint64(t104))
		m.memory[int64(uint32(v2))+112] = byte(v7)
		store32(m.memory[int64(uint32(v2))+100:], uint32(i32(-1)))
		m.fn338(v1+i32(392), v2+i32(88))
		store32(m.memory[uint32(v0):], uint32(i32(-1)))
		if v6 != 0 {
			goto l4
		}
		t105 := int32(load32(m.memory[int64(uint32(v2))+56:]))
		m.fn134(v12, t105)
	}
l4:
	m.fn134(v5, v3)
l29:
	m.g0 = v2 + i32(160)
}
func (m *Module) fn1384(v0, v1 int32) {
	var v2, v3, v4, v5 int32
	t0 := m.g0
	v2 = t0 - i32(48)
	m.g0 = v2
	t1 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	t2 := v2 + i32(8)
	v3 = t1
	t3 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	m.fn865(t2, v3, t3)
	{
		{
			{
				t4 := int32(load32(m.memory[int64(uint32(v2))+16:]))
				v4 = t4
				if v4 == 0 {
					t6 := int32(load32(m.memory[int64(uint32(v2))+8:]))
					t7 := int32(load32(m.memory[int64(uint32(v2))+12:]))
					m.fn16(t6, t7)
					goto l6
				}
				t5 := int32(m.memory[int64(uint32(v0))+458])
				switch t5 {
				case 1:
					t8 := int32(load32(m.memory[int64(uint32(v0))+240:]))
					if t8 == i32(-1) {
						goto l5
					}
					v0 = v0 + i32(240)
					goto l7
				case 2:
					t9 := int32(load32(m.memory[int64(uint32(v0))+200:]))
					v5 = t9
					if v5 == 0 {
						goto l5
					}
					t10 := int32(load32(m.memory[int64(uint32(v0))+196:]))
					v0 = t10 + v5*i32(20) + i32(-20)
					if v0 != 0 {
						goto l7
					}
					goto l5
				case 3:
					goto l4
				case 4:
					goto l5
				default:
					t12 := int32(m.memory[int64(uint32(v0))+455])
					if t12 != 0 {
						goto l5
					}
					t13 := int32(load32(m.memory[int64(uint32(v2))+16:]))
					store32(m.memory[int64(uint32(v2))+32:], uint32(t13))
					t14 := int64(load64(m.memory[int64(uint32(v2))+8:]))
					store64(m.memory[int64(uint32(v2))+24:], uint64(t14))
					store32(m.memory[int64(uint32(v2))+20:], uint32(i32(3)))
					t15 := int32(load32(m.memory[int64(uint32(v0))+432:]))
					store32(m.memory[int64(uint32(v2))+36:], uint32(t15))
					m.fn1340(v0+i32(380), v2+i32(20))
					goto l6
				}
			}
		l4:
			v0 = v0 + i32(228)
		l7:
			t11 := int32(load32(m.memory[int64(uint32(v2))+12:]))
			m.fn75(v0, t11, v4)
			goto l5
		}
	l5:
		t16 := int32(load32(m.memory[int64(uint32(v2))+8:]))
		t17 := int32(load32(m.memory[int64(uint32(v2))+12:]))
		m.fn16(t16, t17)
	}
l6:
	t18 := int32(load32(m.memory[uint32(v1):]))
	m.fn16(t18, v3)
	m.g0 = v2 + i32(48)
}
func (m *Module) fn1385(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	t1 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	v3 = t1
	p2 := v1
	if uint32(v3) > uint32(v1) {
		p2 = v3
	}
	v4 = p2
	v5 = v3 << 6
	v6 = v2 + i32(4)
l2:
	{
		if v4 == v3 {
			m.fn1485(v0, v1)
			t8 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t9 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			t10 := m.fn1486(t8, t9, v1+i32(-1), i32(1075960))
			v7 = t10
			m.g0 = v2 + i32(16)
			return v7
		}
		store32(m.memory[int64(uint32(v6))+8:], uint32(i32(0)))
		store64(m.memory[uint32(v6):], uint64(i64(0)))
		{
			t3 := int32(load32(m.memory[uint32(v0):]))
			if v3 != t3 {
				goto l1
			}
			m.fn1151(v0)
		}
	l1:
		t4 := v0
		v3 = v3 + i32(1)
		store32(m.memory[int64(uint32(t4))+8:], uint32(v3))
		t5 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		v7 = t5 + v5
		store64(m.memory[uint32(v7):], uint64(i64(0x800000000)))
		m.memory[uint32(v7+i32(12))] = byte(i32(0))
		store32(m.memory[uint32(v7+i32(8)):], uint32(i32(0)))
		t6 := int64(load64(m.memory[int64(uint32(v2))+1:]))
		store64(m.memory[uint32(v7+i32(13)):], uint64(t6))
		t7 := int64(load64(m.memory[int64(uint32(v2))+8:]))
		store64(m.memory[uint32(v7+i32(20)):], uint64(t7))
		store64(m.memory[uint32(v7+i32(52)):], uint64(i64(0)))
		store32(m.memory[uint32(v7+i32(48)):], uint32(i32(8)))
		store64(m.memory[uint32(v7+i32(40)):], uint64(i64(0)))
		store64(m.memory[uint32(v7+i32(32)):], uint64(i64(0x400000000)))
		m.memory[uint32(v7+i32(60))] = byte(i32(0))
		v5 = v5 + i32(64)
		goto l2
	}
}
func (m *Module) fn1386(v0, v1, v2 int32) {
	var v3, v4 int32
	var v5 int64
	var v6, v7 int32
	t0 := m.g0
	v3 = t0 - i32(80)
	m.g0 = v3
	t1 := int32(load32(m.memory[int64(uint32(v1))+388:]))
	v4 = t1
	store32(m.memory[int64(uint32(v1))+388:], uint32(i32(0)))
	t2 := int64(load64(m.memory[int64(uint32(v1))+380:]))
	v5 = t2
	store64(m.memory[int64(uint32(v1))+380:], uint64(i64(0x400000000)))
	store32(m.memory[int64(uint32(v3))+8:], uint32(v4))
	store64(m.memory[uint32(v3):], uint64(v5))
	t3 := int32(load32(m.memory[int64(uint32(v1))+240:]))
	t4 := int32(load32(m.memory[int64(uint32(v1))+244:]))
	m.fn134(t3, t4)
	store32(m.memory[int64(uint32(v1))+240:], uint32(i32(-1)))
	v4 = v1 + i32(332)
	t5 := int32(m.memory[int64(uint32(v1))+456])
	m.fn1480(v4, v2, t5, v3)
	m.fn1487(v4, v2)
	m.fn1393(v3+i32(48), v4, v2+i32(1), v2)
	{
		{
			t6 := int32(load32(m.memory[int64(uint32(v3))+48:]))
			v1 = t6
			if v1 == i32(-1) {
				goto l0
			}
			t7 := int32(load32(m.memory[int64(uint32(v3))+68:]))
			store32(m.memory[int64(uint32(v0))+20:], uint32(t7))
			t8 := int64(load64(m.memory[int64(uint32(v3))+60:]))
			store64(m.memory[int64(uint32(v0))+12:], uint64(t8))
			t9 := int64(load64(m.memory[int64(uint32(v3))+52:]))
			store64(m.memory[int64(uint32(v0))+4:], uint64(t9))
			store32(m.memory[uint32(v0):], uint32(v1))
			goto l1
		}
	l0:
		t10 := m.fn1488(v4, v2)
		v1 = t10
		t11 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		v6 = t11
		store32(m.memory[int64(uint32(v1))+8:], uint32(i32(0)))
		t12 := int64(load64(m.memory[uint32(v1):]))
		v5 = t12
		store64(m.memory[uint32(v1):], uint64(i64(0x800000000)))
		store32(m.memory[int64(uint32(v3))+24:], uint32(v6))
		store64(m.memory[int64(uint32(v3))+16:], uint64(v5))
		{
			t13 := m.fn1385(v4, v2)
			v1 = t13
			t14 := int32(load32(m.memory[int64(uint32(v1))+56:]))
			v2 = t14
			t15 := int32(load32(m.memory[int64(uint32(v1))+52:]))
			if uint32(v2) >= uint32(t15) {
				goto l2
			}
			t16 := int32(load32(m.memory[int64(uint32(v1))+48:]))
			v4 = t16 + v2<<4
			t17 := int32(m.memory[int64(uint32(v4))+11])
			v6 = t17
			if v6 == i32(2) {
				goto l2
			}
			t18 := int64(load64(m.memory[uint32(v4):]))
			store64(m.memory[int64(uint32(v3))+32:], uint64(t18))
			t19 := int32(load32(m.memory[int64(uint32(v4))+7:]))
			store32(m.memory[int64(uint32(v3))+39:], uint32(t19))
			t20 := int32(load32(m.memory[int64(uint32(v4))+12:]))
			v7 = t20
			goto l3
		}
	l2:
		v6 = i32(0)
		store32(m.memory[int64(uint32(v3))+39:], uint32(i32(0)))
		store64(m.memory[int64(uint32(v3))+32:], uint64(i64(0)))
	l3:
		store32(m.memory[int64(uint32(v1))+56:], uint32(v2+i32(1)))
		t21 := int32(load32(m.memory[int64(uint32(v3))+24:]))
		store32(m.memory[int64(uint32(v3))+56:], uint32(t21))
		t22 := int64(load64(m.memory[int64(uint32(v3))+16:]))
		store64(m.memory[int64(uint32(v3))+48:], uint64(t22))
		t23 := int64(load64(m.memory[int64(uint32(v3))+32:]))
		store64(m.memory[int64(uint32(v3))+64:], uint64(t23))
		t24 := int32(load32(m.memory[int64(uint32(v3))+39:]))
		store32(m.memory[int64(uint32(v3))+71:], uint32(t24))
		{
			t25 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			v4 = t25
			t26 := int32(load32(m.memory[uint32(v1):]))
			if v4 != t26 {
				goto l4
			}
			m.fn396(v1)
		}
	l4:
		t27 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		v2 = t27 + v4<<5
		t28 := int32(load32(m.memory[int64(uint32(v3))+71:]))
		store32(m.memory[int64(uint32(v2))+23:], uint32(t28))
		t29 := int64(load64(m.memory[int64(uint32(v3))+64:]))
		store64(m.memory[int64(uint32(v2))+16:], uint64(t29))
		t30 := int64(load64(m.memory[int64(uint32(v3))+56:]))
		store64(m.memory[int64(uint32(v2))+8:], uint64(t30))
		t31 := int64(load64(m.memory[int64(uint32(v3))+48:]))
		store64(m.memory[uint32(v2):], uint64(t31))
		store32(m.memory[int64(uint32(v2))+28:], uint32(v7))
		m.memory[int64(uint32(v2))+27] = byte(v6)
		store32(m.memory[int64(uint32(v1))+8:], uint32(v4+i32(1)))
		store32(m.memory[uint32(v0):], uint32(i32(-1)))
	}
l1:
	m.g0 = v3 + i32(80)
}
func (m *Module) fn1387(v0, v1, v2 int32) {
	var v3, v4 int32
	var v5, v6 int64
	t0 := m.g0
	v3 = t0 - i32(32)
	m.g0 = v3
	{
		{
			{
				v4 = v1 + i32(332)
				t1 := m.fn1392(v4, v2)
				if t1 != 0 {
					goto l0
				}
				t2 := int32(load32(m.memory[int64(uint32(v1))+384:]))
				t3 := int32(load32(m.memory[int64(uint32(v1))+388:]))
				t4 := m.fn23(t2, t3)
				if t4 != 0 {
					goto l1
				}
			}
		l0:
			m.fn1386(v3+i32(8), v1, v2)
			t5 := int32(load32(m.memory[int64(uint32(v3))+8:]))
			v1 = t5
			if v1 == i32(-1) {
				goto l1
			}
			t6 := int32(load32(m.memory[int64(uint32(v3))+28:]))
			store32(m.memory[int64(uint32(v0))+20:], uint32(t6))
			t7 := int64(load64(m.memory[int64(uint32(v3))+20:]))
			store64(m.memory[int64(uint32(v0))+12:], uint64(t7))
			t8 := int64(load64(m.memory[int64(uint32(v3))+12:]))
			store64(m.memory[int64(uint32(v0))+4:], uint64(t8))
			goto l2
		}
	l1:
		t9 := m.fn1385(v4, v2)
		v1 = t9
		store32(m.memory[int64(uint32(v1))+56:], uint32(i32(0)))
		t10 := int64(load64(m.memory[int64(uint32(v1))+8:]))
		v5 = t10
		m.memory[int64(uint32(v1))+12] = byte(i32(0))
		store32(m.memory[int64(uint32(v1))+8:], uint32(i32(0)))
		t11 := int64(load64(m.memory[uint32(v1):]))
		v6 = t11
		store64(m.memory[uint32(v1):], uint64(i64(0x800000000)))
		store64(m.memory[int64(uint32(v3))+16:], uint64(v5))
		store64(m.memory[int64(uint32(v3))+8:], uint64(v6))
		t12 := int32(m.memory[int64(uint32(v1))+60])
		m.memory[int64(uint32(v3))+20] = byte(t12)
		{
			if int32(v5) == 0 {
				goto l3
			}
			{
				t13 := int32(load32(m.memory[int64(uint32(v1))+40:]))
				v2 = t13
				t14 := int32(load32(m.memory[int64(uint32(v1))+32:]))
				if v2 != t14 {
					goto l4
				}
				m.fn223(v1 + i32(32))
			}
		l4:
			t15 := int32(load32(m.memory[int64(uint32(v1))+36:]))
			v4 = t15 + v2<<4
			t16 := int64(load64(m.memory[int64(uint32(v3))+16:]))
			store64(m.memory[int64(uint32(v4))+8:], uint64(t16))
			t17 := int64(load64(m.memory[int64(uint32(v3))+8:]))
			store64(m.memory[uint32(v4):], uint64(t17))
			store32(m.memory[int64(uint32(v1))+40:], uint32(v2+i32(1)))
			goto l5
		}
	l3:
		m.fn968(v3 + i32(8))
	l5:
		v1 = i32(-1)
	}
l2:
	store32(m.memory[uint32(v0):], uint32(v1))
	m.g0 = v3 + i32(32)
}
func (m *Module) fn1388(v0 int32) {
	t0 := int32(load32(m.memory[uint32(v0):]))
	if t0 == i32(-1) {
		return
	}
	m.fn1402(v0)
}
func (m *Module) fn1389(v0, v1 int32) {
	var v2, v3 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	{
		t1 := int32(m.memory[int64(uint32(v0))+458])
		if t1 != 0 {
			goto l0
		}
		t2 := int32(m.memory[int64(uint32(v0))+455])
		if t2&i32(1) != 0 {
			goto l1
		}
	}
l0:
	{
		t3 := int32(load32(m.memory[int64(uint32(v0))+316:]))
		v3 = t3
		if v3 == 0 {
			goto l2
		}
		store32(m.memory[int64(uint32(v0))+316:], uint32(v3+i32(-1)))
		goto l1
	}
l2:
	m.fn1382(v0)
	m.fn1072(v2+i32(4), v1)
	m.fn1384(v0, v2+i32(4))
l1:
	m.g0 = v2 + i32(16)
}
func (m *Module) fn1390(v0, v1 int32) {
	if v0 == i32(-1) {
		return
	}
	m.fn16(v0, v1)
}
func (m *Module) fn1391(v0, v1 int32) {
	var v2, v3, v4 int32
	v2 = i32(0)
	{
		{
			t0 := int32(m.memory[int64(uint32(v1))+8])
			if t0 == 0 {
				goto l0
			}
			goto l1
		}
	l0:
		t1 := int32(load32(m.memory[uint32(v1):]))
		v3 = t1
		t2 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t3 := v3
		v4 = t2
		if uint32(t3) > uint32(v4) {
			goto l1
		}
		if uint32(v3) < uint32(v4) {
			goto l2
		}
		v2 = i32(1)
		m.memory[int64(uint32(v1))+8] = byte(i32(1))
		goto l1
	l2:
		store32(m.memory[int64(uint32(v1))+4:], uint32(v4+i32(-1)))
		v2 = i32(1)
	}
l1:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v4))
	store32(m.memory[uint32(v0):], uint32(v2))
}
