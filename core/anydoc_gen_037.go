package core

import (
	"math/bits"
)

func (m *Module) fn1617(v0, v1, v2 int32) {
	{
		if v2 == 0 {
			goto l0
		}
		t0 := m.fn248(v2, v1)
		v1 = t0
	}
l0:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v2))
	store32(m.memory[uint32(v0):], uint32(v1))
}
func (m *Module) fn1618(v0, v1, v2 int32) int64 {
	t0 := int64(load64(m.memory[uint32(v0):]))
	t1 := int64(load64(m.memory[uint32(v0+i32(8)):]))
	t2 := int32(load32(m.memory[uint32(v1-v2<<3+i32(-8)):]))
	t3 := fn1615(t0, t1, t2)
	return t3
}
func (m *Module) fn1619(v0, v1, v2 int32) {
	var v3 int32
l1:
	if v2 != 0 {
		t0 := int32(load32(m.memory[uint32(v0):]))
		v3 = t0
		t1 := int32(load32(m.memory[uint32(v1):]))
		store32(m.memory[uint32(v0):], uint32(t1))
		store32(m.memory[uint32(v1):], uint32(v3))
		v2 = v2 + i32(-1)
		v1 = v1 + i32(4)
		v0 = v0 + i32(4)
		goto l1
	}
}
func (m *Module) fn1620() int32 {
	var v0, v1 int32
	t0 := m.g0
	v0 = t0 - i32(16)
	m.g0 = v0
	m.fn1617(v0+i32(8), i32(4), i32(12))
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
func (m *Module) fn1621(v0 int32) {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	m.fn608(t0, t1, i32(1), i32(1))
}
func (m *Module) fn1622(v0, v1, v2 int32) {
	var v3 int32
	t0 := m.g0
	v3 = t0 - i32(16)
	m.g0 = v3
	m.fn1623(v3+i32(4), v1, v2)
	t1 := m.fn1620()
	v2 = t1
	t2 := int32(load32(m.memory[int64(uint32(v3))+12:]))
	store32(m.memory[int64(uint32(v2))+8:], uint32(t2))
	t3 := int64(load64(m.memory[int64(uint32(v3))+4:]))
	store64(m.memory[uint32(v2):], uint64(t3))
	m.fn343(v0, i32(20), v2, i32(1102000))
	m.g0 = v3 + i32(16)
}
func (m *Module) fn1623(v0, v1, v2 int32) {
	var v3, v4, v5 int32
	t0 := m.g0
	v3 = t0 - i32(16)
	m.g0 = v3
	m.fn599(v3+i32(8), v2, i32(1), i32(1))
	t1 := int32(load32(m.memory[int64(uint32(v3))+8:]))
	v4 = t1
	t2 := int32(load32(m.memory[int64(uint32(v3))+12:]))
	v5 = t2
	store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v5))
	store32(m.memory[uint32(v0):], uint32(v4))
	if v2 == 0 {
		goto l0
	}
	if v2 == 0 {
		goto l1
	}
	memory_copy(m.memory, uint32(v5), uint32(v1), uint32(v2))
l1:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v2))
l0:
	m.g0 = v3 + i32(16)
}
func (m *Module) fn1624(v0, v1 int32) {
	var v2, v3 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	{
		{
			t1 := int32(load16(m.memory[int64(uint32(v1))+8:]))
			v3 = t1
			if v3 != 0 {
				store16(m.memory[int64(uint32(v1))+8:], uint16(i32(0)))
				goto l3
			}
			m.fn374(v2+i32(8), v1)
			t2 := int32(load32(m.memory[int64(uint32(v2))+8:]))
			if t2&i32(1) != 0 {
				goto l1
			}
			v1 = i32(0)
			goto l2
		}
	l1:
		t3 := int32(load32(m.memory[int64(uint32(v2))+12:]))
		v3 = t3
		if uint32(v3) <= uint32(i32(0xffff)) {
			goto l3
		}
		store16(m.memory[int64(uint32(v1))+8:], uint16(v3&i32(1023)|i32(56320)))
		v3 = int32(uint32(v3+i32(0xff0000))>>10) | i32(-10240)
	}
l3:
	v1 = i32(1)
l2:
	store16(m.memory[int64(uint32(v0))+2:], uint16(v3))
	store16(m.memory[uint32(v0):], uint16(v1))
	m.g0 = v2 + i32(16)
}
func (m *Module) fn1625(v0 int32) int32 {
	var v1 int32
	var v2, v3 int64
	var v4, v5, v6, v7 int32
	var v8 int64
	var v9, v10, v11, v12, v13 int32
	var v14 int64
	var v15, v16 int32
	t0 := m.g0
	v1 = t0 - i32(64)
	m.g0 = v1
	{
		t1 := int32(m.memory[int64(uint32(i32(0)))+1303128])
		switch t1 {
		case 2:
			m.fn91(i32(1100728), i32(113), i32(1100712))
			panic("unreachable")
		default:
			m.memory[int64(uint32(i32(0)))+1303128] = byte(i32(2))
			{
				{
					t2 := int32(m.memory[int64(uint32(i32(0)))+1303632])
					if t2 != i32(1) {
						goto l3
					}
					t3 := int64(load64(m.memory[int64(uint32(i32(0)))+1303624:]))
					v2 = t3
					t4 := int64(load64(m.memory[int64(uint32(i32(0)))+1303616:]))
					v3 = t4
					goto l4
				}
			l3:
				m.fn1626(v1 + i32(32))
				t5 := int32(m.memory[int64(uint32(i32(0)))+1303632])
				if t5 == i32(2) {
					m.fn91(i32(1286164), i32(125), i32(1286228))
					panic("unreachable")
				}
				t6 := int64(load64(m.memory[int64(uint32(v1))+40:]))
				v2 = t6
				t7 := int64(load64(m.memory[int64(uint32(v1))+32:]))
				v3 = t7
				m.memory[int64(uint32(i32(0)))+1303632] = byte(i32(1))
				store64(m.memory[int64(uint32(i32(0)))+1303624:], uint64(v2))
			}
		l4:
			store64(m.memory[int64(uint32(v1))+48:], uint64(v3))
			store64(m.memory[int64(uint32(i32(0)))+1303616:], uint64(v3+i64(1)))
			store64(m.memory[int64(uint32(v1))+56:], uint64(v2))
			t8 := int64(load64(m.memory[int64(uint32(i32(0)))+1286256:]))
			store64(m.memory[int64(uint32(v1))+32:], uint64(t8))
			t9 := int64(load64(m.memory[int64(uint32(i32(0)))+1286264:]))
			store64(m.memory[int64(uint32(v1))+40:], uint64(t9))
			t10 := v1 + i32(32)
			v4 = v1 + i32(48)
			m.fn1613(t10, i32(129), v4)
			v5 = i32(0)
		l16:
			{
				t11 := int64(load64(m.memory[int64(uint32(v1))+48:]))
				t12 := int64(load64(m.memory[int64(uint32(v1))+56:]))
				v6 = v5 << 3
				t13 := int32(load32(m.memory[int64(uint32(v6))+1100832:]))
				v7 = t13
				t14 := fn1615(t11, t12, v7)
				v3 = t14
				m.fn1613(v1+i32(32), i32(1), v4)
				v8 = int64(uint64(v3) >> 25)
				v2 = v8 & i64(127) * i64(72340172838076673)
				t15 := int32(load32(m.memory[int64(uint32(v1))+36:]))
				v9 = t15
				v10 = v9 & int32(v3)
				t16 := int32(load32(m.memory[int64(uint32(v6))+1100836:]))
				v11 = t16
				t17 := int32(load32(m.memory[int64(uint32(v1))+32:]))
				v6 = t17
				v12 = i32(0)
				v13 = i32(0)
				{
				l14:
					{
						t18 := int64(load64(m.memory[uint32(v6+v10):]))
						v14 = t18
						v3 = v14 ^ v2
						v3 = (v3 ^ i64(-1)) & (v3 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
					l9:
						if v3 == 0 {
							goto l6
						}
						{
							t19 := v7
							t20 := v6
							v15 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v3))))>>3) + v10) & v9
							t21 := int32(load32(m.memory[uint32(t20-v15<<3+i32(-8)):]))
							if t19 != t21 {
								v3 = (v3 + i64(-1)) & v3
								goto l9
							}
							v10 = i32(0) - v15
							goto l8
						}
					l6:
						v3 = v14 & i64(-0x7f7f7f7f7f7f7f80)
						{
							if v12 == i32(1) {
								goto l10
							}
							if !(v3 == 0) {
								goto l11
							}
							v12 = i32(0)
							goto l12
						l11:
							v16 = (v10 + int32(uint32(int64(bits.TrailingZeros64(uint64(v3))))>>3)) & v9
						l10:
							if v3&(v14<<1) != i64(0) {
								goto l13
							}
							v12 = i32(1)
						l12:
							t22 := v10
							v13 = v13 + i32(8)
							v10 = (t22 + v13) & v9
							goto l14
						}
					l13:
					}
					{
						t23 := int32(int8(m.memory[uint32(v6+v16)]))
						v10 = t23
						if v10 < i32(0) {
							goto l15
						}
						t24 := int64(load64(m.memory[uint32(v6):]))
						t25 := v6
						v16 = int32(uint32(int64(bits.TrailingZeros64(uint64(t24&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
						t26 := int32(m.memory[uint32(t25+v16)])
						v10 = t26
					}
				l15:
					t27 := v6 + v16
					v15 = int32(v8) & i32(127)
					m.memory[uint32(t27)] = byte(v15)
					m.memory[uint32(v6+(v16+i32(-8))&v9+i32(8))] = byte(v15)
					store32(m.memory[uint32(v6-v16<<3+i32(-8)):], uint32(v7))
					t28 := int32(load32(m.memory[int64(uint32(v1))+44:]))
					store32(m.memory[int64(uint32(v1))+44:], uint32(t28+i32(1)))
					t29 := int32(load32(m.memory[int64(uint32(v1))+40:]))
					store32(m.memory[int64(uint32(v1))+40:], uint32(t29-v10&i32(1)))
					v10 = i32(0) - v16
				}
			l8:
				store32(m.memory[uint32(v6+v10<<3+i32(-4)):], uint32(v11))
				v5 = v5 + i32(1)
				if v5 != i32(129) {
					goto l16
				}
			}
			t30 := int64(load64(m.memory[int64(uint32(v1))+32:]))
			store64(m.memory[int64(uint32(i32(0)))+1303096:], uint64(t30))
			t31 := int64(load64(m.memory[int64(uint32(v1))+40:]))
			store64(m.memory[int64(uint32(i32(0)))+1303104:], uint64(t31))
			t32 := int64(load64(m.memory[int64(uint32(v1))+48:]))
			store64(m.memory[int64(uint32(i32(0)))+1303112:], uint64(t32))
			t33 := int64(load64(m.memory[int64(uint32(v1))+56:]))
			store64(m.memory[int64(uint32(i32(0)))+1303120:], uint64(t33))
			m.memory[int64(uint32(i32(0)))+1303128] = byte(i32(3))
			fallthrough
		case 3:
			v6 = i32(-1)
			{
				t34 := int32(load32(m.memory[int64(uint32(i32(0)))+1303108:]))
				if t34 == 0 {
					goto l17
				}
				v11 = i32(0)
				t35 := int32(load32(m.memory[int64(uint32(i32(0)))+1303100:]))
				v9 = t35
				t36 := int64(load64(m.memory[int64(uint32(i32(0)))+1303112:]))
				t37 := int64(load64(m.memory[int64(uint32(i32(0)))+1303120:]))
				t38 := fn1615(t36, t37, v0)
				t39 := v9
				v3 = t38
				v5 = t39 & int32(v3)
				v2 = int64(uint64(v3)>>25) & i64(127) * i64(72340172838076673)
				t40 := int32(load32(m.memory[int64(uint32(i32(0)))+1303096:]))
				v10 = t40
			l21:
				{
					t41 := int64(load64(m.memory[uint32(v10+v5):]))
					v14 = t41
					v3 = v14 ^ v2
					v3 = (v3 ^ i64(-1)) & (v3 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
				l20:
					{
						if v3 == 0 {
							if !(v14&(v14<<1)&i64(-0x7f7f7f7f7f7f7f80) == 0) {
								goto l17
							}
							t44 := v5
							v11 = v11 + i32(8)
							v5 = (t44 + v11) & v9
							goto l21
						}
						t42 := v0
						v7 = v10 - (int32(uint32(int64(bits.TrailingZeros64(uint64(v3))))>>3)+v5)&v9<<3
						t43 := int32(load32(m.memory[uint32(v7+i32(-8)):]))
						if t42 == t43 {
							goto l19
						}
						v3 = (v3 + i64(-1)) & v3
						goto l20
					}
				l19:
				}
				t45 := int32(load32(m.memory[uint32(v7+i32(-4)):]))
				v6 = t45
			}
		l17:
			{
				{
					if uint32(v0) < uint32(i32(181)) {
						goto l22
					}
					m.fn1627(v1+i32(32), v0, i32(1114812))
					t46 := int32(load32(m.memory[int64(uint32(v1))+32:]))
					t47 := v0
					v5 = t46
					p48 := v5
					if v5 == i32(-1) {
						p48 = t47
					}
					v0 = p48
					goto l23
				}
			l22:
				p49 := v0
				if uint32(v0+i32(-97)) < uint32(i32(26)) {
					p49 = v0 & i32(95)
				}
				v0 = p49
			}
		l23:
			m.g0 = v1 + i32(64)
			p50 := v6
			if v6 == i32(-1) {
				p50 = v0
			}
			return p50
		}
	}
}
func (m *Module) fn1626(v0 int32) {
	var v1, v2 int32
	t0 := m.g0
	v1 = t0 - i32(16)
	m.g0 = v1
	m.memory[int64(uint32(v1))+15] = byte(i32(0))
	{
		t1 := m.fn4(i32(1))
		v2 = t1
		if v2 != 0 {
			goto l0
		}
		m.fn85(i32(1), i32(1))
		panic("unreachable")
	}
l0:
	store64(m.memory[uint32(v0):], uint64(uint32(v1+i32(15))))
	store64(m.memory[int64(uint32(v0))+8:], uint64(uint32(v2)))
	m.fn10(v2, i32(1), i32(1))
	m.g0 = v1 + i32(16)
}
func (m *Module) fn1627(v0, v1, v2 int32) {
	var v3, v4, v5, v6, v7, v8 int32
	if uint32(v1) < uint32(i32(0x20000)) {
		goto l0
	}
	store32(m.memory[uint32(v0):], uint32(i32(-1)))
	return
l0:
	v3 = v2 + int32(uint32(v1)>>12)&i32(0xffff0)
	t0 := int32(load32(m.memory[uint32(v3):]))
	v4 = t0
	v2 = i32(0)
	{
		t1 := int32(load32(m.memory[int64(uint32(v3))+4:]))
		v5 = t1
		switch v5 {
		case 0:
			goto l1
		default:
			v2 = i32(0)
			v6 = v1 & i32(0xffff)
		l4:
			{
				t2 := v2
				v7 = int32(uint32(v5) >> 1)
				v8 = v7 + v2
				t3 := int32(load16(m.memory[uint32(v4+v8*i32(6)):]))
				p4 := v8
				if uint32(t3) > uint32(v6) {
					p4 = t2
				}
				v2 = p4
				v5 = v5 - v7
				if uint32(v5) > uint32(i32(1)) {
					goto l4
				}
			}
			fallthrough
		case 1:
			v2 = v4 + v2*i32(6)
			t5 := int32(load16(m.memory[uint32(v2):]))
			v5 = t5
			t6 := v5
			v7 = v1 & i32(0xffff)
			if uint32(t6) > uint32(v7) {
				goto l1
			}
			t7 := int32(m.memory[uint32(v2+i32(2))])
			if uint32((v5+t7)&i32(0xffff)) < uint32(v7) {
				goto l1
			}
			t8 := int32(m.memory[int64(uint32(v2))+3])
			if (v5^v1)&t8&i32(1) != 0 {
				goto l1
			}
			store64(m.memory[int64(uint32(v0))+4:], uint64(i64(0)))
			t9 := int32(load16(m.memory[int64(uint32(v2))+4:]))
			store32(m.memory[uint32(v0):], uint32(v1&i32(65536)|(t9+v1)&i32(0xffff)))
			return
		}
	}
l1:
	t10 := int32(load32(m.memory[int64(uint32(v3))+8:]))
	v4 = t10
	v2 = i32(0)
	{
		t11 := int32(load32(m.memory[int64(uint32(v3))+12:]))
		v5 = t11
		switch v5 {
		default:
			v2 = i32(0)
			v6 = v1 & i32(0xffff)
		l8:
			{
				t12 := v2
				v7 = int32(uint32(v5) >> 1)
				v8 = v7 + v2
				t13 := int32(load16(m.memory[uint32(v4+v8<<3):]))
				p14 := v8
				if uint32(t13) > uint32(v6) {
					p14 = t12
				}
				v2 = p14
				v5 = v5 - v7
				if uint32(v5) > uint32(i32(1)) {
					goto l8
				}
			}
			fallthrough
		case 1:
			v2 = v4 + v2<<3
			t15 := int32(load16(m.memory[uint32(v2):]))
			if t15 == v1&i32(0xffff) {
				goto l9
			}
			fallthrough
		case 0:
			store32(m.memory[uint32(v0):], uint32(i32(-1)))
			return
		}
	}
l9:
	t16 := v0
	v5 = v1 & i32(65536)
	t17 := int32(load16(m.memory[int64(uint32(v2))+6:]))
	store32(m.memory[int64(uint32(t16))+8:], uint32(v5|t17))
	t18 := int32(load16(m.memory[int64(uint32(v2))+4:]))
	store32(m.memory[int64(uint32(v0))+4:], uint32(v5|t18))
	t19 := int32(load16(m.memory[int64(uint32(v2))+2:]))
	store32(m.memory[uint32(v0):], uint32(v5|t19))
}
func (m *Module) fn1628(v0 int32) int32 {
	var v1 int32
	v1 = i32(-1)
	{
		t0 := m.fn48(v0)
		v0 = t0
		if v0 == i32(-1) {
			goto l0
		}
		t1 := m.fn1625(v0)
		v1 = t1
	}
l0:
	return v1
}
func (m *Module) fn1629(v0, v1, v2, v3, v4, v5 int32) {
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
func (m *Module) fn1630(v0, v1 int32) int32 {
	var v2, v3 int32
	if uint32(v1) > uint32(i32(3)) {
		v3 = i32(0)
		t1 := int32(load32(m.memory[uint32(v0):]))
		if t1&i32(-2139062144) != 0 {
			goto l3
		}
		v2 = (v0 + i32(3)) & i32(-4)
		p2 := v2 - v0
		if v2 == v0 {
			p2 = i32(4)
		}
		v2 = p2
		v1 = v1 + i32(-4)
	l5:
		{
			if uint32(v2) < uint32(v1) {
				t5 := int32(load32(m.memory[uint32(v0+v2):]))
				if t5&i32(-2139062144) != 0 {
					goto l3
				}
				v2 = v2 + i32(4)
				goto l5
			}
			t3 := int32(load32(m.memory[uint32(v0+v1):]))
			var p4 int32
			if t3&i32(-2139062144) == 0 {
				p4 = 1
			}
			return p4
		}
	}
	v0 = v0 + i32(-1)
l2:
	{
		if v1 != 0 {
			goto l1
		}
		return i32(1)
	l1:
		v2 = v0 + v1
		v1 = v1 + i32(-1)
		t0 := int32(int8(m.memory[uint32(v2)]))
		if t0 > i32(-1) {
			goto l2
		}
	}
	v3 = i32(0)
	goto l3
l3:
	return v3
}
func (m *Module) fn1631(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t2 := m.fn110(v1, t0, t1)
	return t2
}
func (m *Module) fn1632(v0, v1, v2, v3, v4, v5, v6, v7 int32) {
	var v8 int32
	var v9 int64
	t0 := m.g0
	v8 = t0 - i32(64)
	m.g0 = v8
	store32(m.memory[int64(uint32(v8))+4:], uint32(v2))
	store32(m.memory[uint32(v8):], uint32(v1))
	store32(m.memory[int64(uint32(v8))+12:], uint32(v4))
	store32(m.memory[int64(uint32(v8))+8:], uint32(v3))
	store32(m.memory[int64(uint32(v8))+20:], uint32(i32(2)))
	t2 := v8
	p1 := i32(1109315)
	if v0&i32(1) != 0 {
		p1 = i32(1109317)
	}
	store32(m.memory[int64(uint32(t2))+16:], uint32(p1))
	{
		if v5 == 0 {
			t4 := v8
			v9 = int64(uint32(i32(40))) << 32
			store64(m.memory[int64(uint32(t4))+48:], uint64(v9|int64(uint32(v8+i32(8)))))
			store64(m.memory[int64(uint32(v8))+40:], uint64(v9|int64(uint32(v8))))
			store64(m.memory[int64(uint32(v8))+32:], uint64(int64(uint32(i32(41)))<<32|int64(uint32(v8+i32(16)))))
			m.fn91(i32(1051286), v8+i32(32), v7)
			panic("unreachable")
		}
		store32(m.memory[int64(uint32(v8))+28:], uint32(v6))
		store32(m.memory[int64(uint32(v8))+24:], uint32(v5))
		t3 := v8
		v9 = int64(uint32(i32(40))) << 32
		store64(m.memory[int64(uint32(t3))+56:], uint64(v9|int64(uint32(v8+i32(8)))))
		store64(m.memory[int64(uint32(v8))+48:], uint64(v9|int64(uint32(v8))))
		store64(m.memory[int64(uint32(v8))+40:], uint64(int64(uint32(i32(198)))<<32|int64(uint32(v8+i32(24)))))
		store64(m.memory[int64(uint32(v8))+32:], uint64(int64(uint32(i32(41)))<<32|int64(uint32(v8+i32(16)))))
		m.fn91(i32(1051341), v8+i32(32), v7)
		panic("unreachable")
	}
}
func (m *Module) fn1633(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t2 := int32(load32(m.memory[int64(uint32(t1))+12:]))
	t3 := m.t0[uint(t2)].(func(int32, int32) int32)(t0, v1)
	return t3
}
func (m *Module) fn1634(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v1):]))
	t1 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	t2 := int32(load32(m.memory[uint32(v0):]))
	t3 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t4 := m.fn100(t0, t1, t2, t3)
	return t4
}
func (m *Module) fn1635(v0, v1 int32) int32 {
	t0 := int32(load32(m.memory[uint32(v0):]))
	t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	t2 := m.fn110(v1, t0, t1)
	return t2
}
func (m *Module) fn1636(v0 int32) {
	var v1 int32
	var v2 int64
	t0 := m.g0
	v1 = t0 - i32(16)
	m.g0 = v1
	t1 := int64(load64(m.memory[uint32(v0):]))
	v2 = t1
	store32(m.memory[int64(uint32(v1))+12:], uint32(v0))
	store64(m.memory[int64(uint32(v1))+4:], uint64(v2))
	m.fn1793(v1 + i32(4))
	panic("unreachable")
}
func (m *Module) fn1637(v0, v1 int32) int32 {
	var v2, v3 int32
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
				t4 := int32(load32(m.memory[uint32(v0):]))
				v3 = t4
				v0 = i32(9)
			l3:
				{
					t5 := int32(m.memory[int64(uint32(v3&i32(15)))+1107936])
					m.memory[uint32(v2+i32(8)+v0+i32(-2))] = byte(t5)
					v0 = v0 + i32(-1)
					v3 = int32(uint32(v3) >> 4)
					if v3 != 0 {
						goto l3
					}
				}
				t6 := m.fn1638(v1, i32(1), i32(1131670), i32(2), v2+i32(8)+v0+i32(-1), i32(9)-v0)
				v0 = t6
				goto l2
			}
			if v3&i32(0x4000000) != 0 {
				goto l1
			}
			t3 := m.fn72(v0, v1)
			v0 = t3
			goto l2
		}
	l1:
		t7 := int32(load32(m.memory[uint32(v0):]))
		v3 = t7
		v0 = i32(9)
	l4:
		{
			t8 := int32(m.memory[int64(uint32(v3&i32(15)))+1131672])
			m.memory[uint32(v2+i32(8)+v0+i32(-2))] = byte(t8)
			v0 = v0 + i32(-1)
			v3 = int32(uint32(v3) >> 4)
			if v3 != 0 {
				goto l4
			}
		}
		t9 := m.fn1638(v1, i32(1), i32(1131670), i32(2), v2+i32(8)+v0+i32(-1), i32(9)-v0)
		v0 = t9
	}
l2:
	m.g0 = v2 + i32(16)
	return v0
}
func (m *Module) fn1638(v0, v1, v2, v3, v4, v5 int32) int32 {
	var v6, v7, v8, v9, v10, v11, v12, v13 int32
	var v14 int64
	t0 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	v6 = t0
	v7 = v6 & i32(0x200000)
	p1 := i32(-1)
	if v7 != 0 {
		p1 = i32(43)
	}
	v8 = p1
	p2 := i32(1)
	if v1 != 0 {
		p2 = int32(uint32(v7) >> 21)
	}
	v9 = p2 + v5
	if v6&i32(0x800000) != 0 {
		goto l0
	}
	v2 = i32(0)
	goto l1
l0:
	v10 = i32(0)
	if v3 == 0 {
		goto l2
	}
	v7 = v2
	v11 = v3
l3:
	{
		t3 := int32(int8(m.memory[uint32(v7)]))
		t4 := v10
		var p5 int32
		if t3 > i32(-65) {
			p5 = 1
		}
		v10 = t4 + p5
		v7 = v7 + i32(1)
		v11 = v11 + i32(-1)
		if v11 != 0 {
			goto l3
		}
	}
l2:
	v9 = v10 + v9
l1:
	p6 := i32(45)
	if v1 != 0 {
		p6 = v8
	}
	v12 = p6
	{
		{
			t7 := int32(load16(m.memory[int64(uint32(v0))+12:]))
			t8 := v9
			v8 = t7
			if uint32(t8) >= uint32(v8) {
				goto l4
			}
			{
				if v6&i32(0x1000000) != 0 {
					t13 := int64(load64(m.memory[int64(uint32(v0))+8:]))
					t14 := v0
					v14 = t13
					store32(m.memory[int64(uint32(t14))+8:], uint32(int32(v14)&i32(-0x60200000)|i32(0x20000030)))
					v10 = i32(1)
					t15 := int32(load32(m.memory[uint32(v0):]))
					v11 = t15
					t16 := int32(load32(m.memory[int64(uint32(v0))+4:]))
					t17 := v11
					v1 = t16
					t18 := m.fn1642(t17, v1, v12, v2, v3)
					if t18 != 0 {
						goto l11
					}
					v7 = i32(0)
					v9 = (v8 - v9) & i32(0xffff)
				l13:
					{
						if uint32(v7&i32(0xffff)) >= uint32(v9) {
							v10 = i32(1)
							t26 := int32(load32(m.memory[int64(uint32(v1))+12:]))
							t27 := m.t0[uint(t26)].(func(int32, int32, int32) int32)(v11, v4, v5)
							if t27 != 0 {
								goto l11
							}
							store64(m.memory[int64(uint32(v0))+8:], uint64(v14))
							return i32(0)
						}
						v10 = i32(1)
						v7 = v7 + i32(1)
						t19 := int32(load32(m.memory[int64(uint32(v1))+16:]))
						t20 := m.t0[uint(t19)].(func(int32, int32) int32)(v11, i32(48))
						if t20 == 0 {
							goto l13
						}
						goto l11
					}
				}
				v13 = v8 - v9
				v7 = i32(0)
				v8 = i32(0)
				switch int32(uint32(v6)>>29) & i32(3) {
				default:
					goto l6
				case 1, 3:
					v8 = v13
					goto l6
				case 2:
					v8 = int32(uint32(v13&i32(65534)) >> 1)
				}
			l6:
				v1 = v6 & i32(0x1fffff)
				t9 := int32(load32(m.memory[int64(uint32(v0))+4:]))
				v9 = t9
				t10 := int32(load32(m.memory[uint32(v0):]))
				v11 = t10
			l10:
				{
					if uint32(v7&i32(0xffff)) >= uint32(v8&i32(0xffff)) {
						v10 = i32(1)
						t21 := m.fn1642(v11, v9, v12, v2, v3)
						if t21 != 0 {
							goto l11
						}
						t22 := int32(load32(m.memory[int64(uint32(v9))+12:]))
						t23 := m.t0[uint(t22)].(func(int32, int32, int32) int32)(v11, v4, v5)
						if t23 != 0 {
							goto l11
						}
						v0 = (v13 - v8) & i32(0xffff)
						v7 = i32(0)
					l15:
						if uint32(v7&i32(0xffff)) < uint32(v0) {
							v10 = i32(1)
							v7 = v7 + i32(1)
							t24 := int32(load32(m.memory[int64(uint32(v9))+16:]))
							t25 := m.t0[uint(t24)].(func(int32, int32) int32)(v11, v1)
							if t25 == 0 {
								goto l15
							}
							goto l11
						}
						return i32(0)
					}
					v10 = i32(1)
					v7 = v7 + i32(1)
					t11 := int32(load32(m.memory[int64(uint32(v9))+16:]))
					t12 := m.t0[uint(t11)].(func(int32, int32) int32)(v11, v1)
					if t12 == 0 {
						goto l10
					}
					goto l11
				}
			}
		}
	l4:
		v10 = i32(1)
		t28 := int32(load32(m.memory[uint32(v0):]))
		v7 = t28
		t29 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		t30 := v7
		v11 = t29
		t31 := m.fn1642(t30, v11, v12, v2, v3)
		if t31 != 0 {
			goto l11
		}
		t32 := int32(load32(m.memory[int64(uint32(v11))+12:]))
		t33 := m.t0[uint(t32)].(func(int32, int32, int32) int32)(v7, v4, v5)
		v10 = t33
	}
l11:
	return v10
}
func (m *Module) fn1639(v0, v1, v2 int32) int32 {
	var v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13 int32
	t0 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	v3 = t0
	t1 := int32(load32(m.memory[uint32(v0):]))
	v4 = t1
	t2 := int32(load32(m.memory[int64(uint32(v0))+8:]))
	v5 = t2
	v6 = i32(0)
	v7 = i32(0)
	v8 = i32(0)
	v9 = i32(0)
l20:
	if v9&i32(1) != 0 {
		goto l0
	}
	if uint32(v2) < uint32(v8) {
		goto l1
	}
l16:
	v9 = v1 + v8
	v10 = v2 - v8
	if uint32(v10) > uint32(i32(7)) {
		v0 = (v9 + i32(3)) & i32(-4)
		if v0 == v9 {
			goto l4
		}
		v11 = v0 - v9
		v0 = i32(0)
	l6:
		{
			t3 := int32(m.memory[uint32(v9+v0)])
			if t3 == i32(10) {
				goto l5
			}
			t4 := v11
			v0 = v0 + i32(1)
			if t4 != v0 {
				goto l6
			}
		}
		t5 := v11
		v12 = v10 + i32(-8)
		if uint32(t5) > uint32(v12) {
			goto l7
		}
		goto l10
	}
	if v2 != v8 {
		v0 = i32(0)
	l9:
		{
			t6 := int32(m.memory[uint32(v9+v0)])
			if t6 == i32(10) {
				goto l5
			}
			t7 := v10
			v0 = v0 + i32(1)
			if t7 != v0 {
				goto l9
			}
		}
		v8 = v2
		goto l1
	}
	v8 = v2
	goto l1
l4:
	v12 = v10 + i32(-8)
	v11 = i32(0)
l10:
	{
		v0 = v9 + v11
		t8 := int32(load32(m.memory[uint32(v0):]))
		v13 = t8
		t9 := int32(load32(m.memory[uint32(v0+i32(4)):]))
		t10 := i32(16843008) - (v13 ^ i32(168430090)) | v13
		v0 = t9
		if t10&(i32(16843008)-(v0^i32(168430090))|v0)&i32(-2139062144) != i32(-2139062144) {
			goto l7
		}
		v11 = v11 + i32(8)
		if uint32(v11) <= uint32(v12) {
			goto l10
		}
	}
l7:
	if v10 != v11 {
		goto l11
	}
	v8 = v2
	goto l1
l11:
	v13 = v9 + v11
	v10 = v2 - v11 - v8
	v0 = i32(0)
l13:
	{
		t11 := int32(m.memory[uint32(v13+v0)])
		if t11 == i32(10) {
			goto l12
		}
		t12 := v10
		v0 = v0 + i32(1)
		if t12 != v0 {
			goto l13
		}
	}
	v8 = v2
	goto l1
l12:
	v0 = v0 + v11
l5:
	v11 = v8 + v0
	v8 = v11 + i32(1)
	{
		if uint32(v11) >= uint32(v2) {
			goto l14
		}
		t13 := int32(m.memory[uint32(v9+v0)])
		if t13 != i32(10) {
			goto l14
		}
		v9 = i32(0)
		v13 = v8
		v0 = v8
		goto l15
	}
l14:
	if uint32(v2) >= uint32(v8) {
		goto l16
	}
l1:
	if v2 == v7 {
		goto l0
	}
	v9 = i32(1)
	v13 = v7
	v0 = v2
l15:
	{
		{
			t14 := int32(m.memory[uint32(v5)])
			if t14 == 0 {
				goto l17
			}
			t15 := int32(load32(m.memory[int64(uint32(v3))+12:]))
			t16 := m.t0[uint(t15)].(func(int32, int32, int32) int32)(v4, i32(1131108), i32(4))
			if t16 != 0 {
				goto l18
			}
		}
	l17:
		v10 = v0 - v7
		v11 = i32(0)
		{
			if v0 == v7 {
				goto l19
			}
			t17 := int32(m.memory[uint32(v1+v0+i32(-1))])
			var p18 int32
			if t17 == i32(10) {
				p18 = 1
			}
			v11 = p18
		}
	l19:
		v0 = v1 + v7
		m.memory[uint32(v5)] = byte(v11)
		v7 = v13
		t19 := int32(load32(m.memory[int64(uint32(v3))+12:]))
		t20 := m.t0[uint(t19)].(func(int32, int32, int32) int32)(v4, v0, v10)
		if t20 == 0 {
			goto l20
		}
	}
l18:
	v6 = i32(1)
l0:
	return v6
}
func (m *Module) fn1640(v0, v1 int32) int32 {
	var v2, v3 int32
	t0 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	v2 = t0
	t1 := int32(load32(m.memory[uint32(v0):]))
	v3 = t1
	{
		t2 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		v0 = t2
		t3 := int32(m.memory[uint32(v0)])
		if t3 == 0 {
			goto l0
		}
		t4 := int32(load32(m.memory[int64(uint32(v2))+12:]))
		t5 := m.t0[uint(t4)].(func(int32, int32, int32) int32)(v3, i32(1131108), i32(4))
		if t5 == 0 {
			goto l0
		}
		return i32(1)
	}
l0:
	t6 := v0
	var p7 int32
	if v1 == i32(10) {
		p7 = 1
	}
	m.memory[uint32(t6)] = byte(p7)
	t8 := int32(load32(m.memory[int64(uint32(v2))+16:]))
	t9 := m.t0[uint(t8)].(func(int32, int32) int32)(v3, v1)
	return t9
}
func (m *Module) fn1641(v0, v1, v2 int32) int32 {
	t0 := m.fn100(v0, i32(1109040), v1, v2)
	return t0
}
func (m *Module) fn1642(v0, v1, v2, v3, v4 int32) int32 {
	{
		if v2 == i32(-1) {
			goto l0
		}
		t0 := int32(load32(m.memory[int64(uint32(v1))+16:]))
		t1 := m.t0[uint(t0)].(func(int32, int32) int32)(v0, v2)
		if t1 == 0 {
			goto l0
		}
		return i32(1)
	}
l0:
	if v3 != 0 {
		t2 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		t3 := m.t0[uint(t2)].(func(int32, int32, int32) int32)(v0, v3, v4)
		return t3
	}
	return i32(0)
}
func (m *Module) fn1643(v0, v1 int32) {
	var v2 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	store32(m.memory[int64(uint32(v2))+12:], uint32(v1))
	store32(m.memory[int64(uint32(v2))+8:], uint32(v0))
	m.fn1632(i32(0), v2+i32(8), i32(1107968), v2+i32(12), i32(1107968), i32(0), v2, i32(1108920))
	panic("unreachable")
}
func (m *Module) fn1644(v0, v1 int32) int32 {
	var v2, v3 int32
	var v4 int64
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
				t4 := int64(load64(m.memory[uint32(v0):]))
				v4 = t4
				v0 = i32(17)
			l3:
				{
					t5 := int32(m.memory[int64(uint32(int32(v4)&i32(15)))+1107936])
					m.memory[uint32(v2+v0+i32(-2))] = byte(t5)
					v0 = v0 + i32(-1)
					v4 = int64(uint64(v4) >> 4)
					if v4 != i64(0) {
						goto l3
					}
				}
				t6 := m.fn1638(v1, i32(1), i32(1131670), i32(2), v2+v0+i32(-1), i32(17)-v0)
				v0 = t6
				goto l2
			}
			if v3&i32(0x4000000) != 0 {
				goto l1
			}
			t3 := m.fn579(v0, v1)
			v0 = t3
			goto l2
		}
	l1:
		t7 := int64(load64(m.memory[uint32(v0):]))
		v4 = t7
		v0 = i32(17)
	l4:
		{
			t8 := int32(m.memory[int64(uint32(int32(v4)&i32(15)))+1131672])
			m.memory[uint32(v2+v0+i32(-2))] = byte(t8)
			v0 = v0 + i32(-1)
			v4 = int64(uint64(v4) >> 4)
			if v4 != i64(0) {
				goto l4
			}
		}
		t9 := m.fn1638(v1, i32(1), i32(1131670), i32(2), v2+v0+i32(-1), i32(17)-v0)
		v0 = t9
	}
l2:
	m.g0 = v2 + i32(16)
	return v0
}
func (m *Module) fn1645(v0, v1, v2 int32) {
	var v3, v4, v5, v6, v7, v8 int32
	t0 := m.g0
	v3 = t0 - i32(32)
	m.g0 = v3
	{
		switch v1 {
		case 34:
			goto l5
		case 39:
			if v2&i32(256) == 0 {
				goto l12
			}
			store64(m.memory[int64(uint32(v0))+2:], uint64(i64(0)))
			store16(m.memory[uint32(v0):], uint16(i32(10076)))
			goto l14
		default:
			if v1 == i32(92) {
				store64(m.memory[int64(uint32(v0))+2:], uint64(i64(0)))
				store16(m.memory[uint32(v0):], uint16(i32(23644)))
				goto l14
			}
			fallthrough
		case 1, 2, 3, 4, 5, 6, 7, 8, 11, 12, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 35, 36, 37, 38:
			if uint32(v1) < uint32(i32(768)) {
				goto l9
			}
			if v2&i32(1) != 0 {
				t1 := m.fn1646(v1)
				if t1 == 0 {
					goto l13
				}
				m.memory[int64(uint32(v3))+14] = byte(i32(0))
				store16(m.memory[int64(uint32(v3))+12:], uint16(i32(0)))
				t2 := int32(m.memory[int64(uint32(int32(uint32(v1)>>20)))+1107936])
				m.memory[int64(uint32(v3))+15] = byte(t2)
				t3 := int32(m.memory[int64(uint32(int32(uint32(v1)>>4)&i32(15)))+1107936])
				m.memory[int64(uint32(v3))+19] = byte(t3)
				t4 := int32(m.memory[int64(uint32(int32(uint32(v1)>>8)&i32(15)))+1107936])
				m.memory[int64(uint32(v3))+18] = byte(t4)
				t5 := int32(m.memory[int64(uint32(int32(uint32(v1)>>12)&i32(15)))+1107936])
				m.memory[int64(uint32(v3))+17] = byte(t5)
				t6 := int32(m.memory[int64(uint32(int32(uint32(v1)>>16)&i32(15)))+1107936])
				m.memory[int64(uint32(v3))+16] = byte(t6)
				t7 := v3 + i32(12)
				v2 = int32(uint32(int32(bits.LeadingZeros32(uint32(v1|i32(1))))) >> 2)
				v4 = t7 + v2
				m.memory[uint32(v4)] = byte(i32(123))
				m.memory[uint32(v4+i32(-1))] = byte(i32(117))
				t8 := v3 + i32(12)
				v2 = v2 + i32(-2)
				m.memory[uint32(t8+v2)] = byte(i32(92))
				t9 := int64(load64(m.memory[int64(uint32(v3))+12:]))
				store64(m.memory[uint32(v0):], uint64(t9))
				m.memory[int64(uint32(v3))+21] = byte(i32(125))
				t10 := int32(m.memory[int64(uint32(v1&i32(15)))+1107936])
				m.memory[int64(uint32(v3))+20] = byte(t10)
				t11 := int32(load16(m.memory[int64(uint32(v3))+20:]))
				store16(m.memory[int64(uint32(v0))+8:], uint16(t11))
				goto l15
			}
		l9:
			if uint32(v1) < uint32(i32(32)) {
				goto l11
			}
			if uint32(v1) < uint32(i32(127)) {
				goto l12
			}
			goto l13
		case 0:
			store64(m.memory[int64(uint32(v0))+2:], uint64(i64(0)))
			store16(m.memory[uint32(v0):], uint16(i32(12380)))
			goto l14
		case 9:
			store64(m.memory[int64(uint32(v0))+2:], uint64(i64(0)))
			store16(m.memory[uint32(v0):], uint16(i32(29788)))
			goto l14
		case 13:
			store64(m.memory[int64(uint32(v0))+2:], uint64(i64(0)))
			store16(m.memory[uint32(v0):], uint16(i32(29276)))
			goto l14
		case 10:
			store64(m.memory[int64(uint32(v0))+2:], uint64(i64(0)))
			store16(m.memory[uint32(v0):], uint16(i32(28252)))
			goto l14
		}
	l5:
		if uint32(v2&i32(0xffffff)) < uint32(i32(65536)) {
			goto l12
		}
		store64(m.memory[int64(uint32(v0))+2:], uint64(i64(0)))
		store16(m.memory[uint32(v0):], uint16(i32(8796)))
	l14:
		v1 = i32(2)
		v2 = i32(0)
		goto l16
	l13:
		if uint32(v1) < uint32(i32(65536)) {
			v5 = i32(0)
			v6 = int32(uint32(v1)>>8) & i32(255)
			v2 = i32(0)
		l21:
			{
				v7 = v2 + i32(2)
				t12 := int32(m.memory[int64(uint32(v2))+1110621])
				t13 := v5
				v4 = t12
				v8 = t13 + v4
				{
					t14 := int32(m.memory[int64(uint32(v2))+1110620])
					v2 = t14
					if v2 == v6 {
						if uint32(v8) < uint32(v5) {
							goto l22
						}
						if uint32(v8) > uint32(i32(284)) {
							goto l22
						}
						if v4 == 0 {
							goto l23
						}
						v2 = v5 + i32(1110696)
						goto l25
					l22:
						m.fn151(v5, v8, i32(284), i32(1111272))
						panic("unreachable")
					l25:
						{
							t15 := int32(m.memory[uint32(v2)])
							if t15 == v1&i32(255) {
								goto l11
							}
							v2 = v2 + i32(1)
							v4 = v4 + i32(-1)
							if v4 != 0 {
								goto l25
							}
						}
					l23:
						v5 = v8
						v2 = v7
						if v7 != i32(76) {
							goto l21
						}
						goto l20
					}
					if uint32(v2) > uint32(v6) {
						goto l20
					}
					v5 = v8
					v2 = v7
					if v7 != i32(76) {
						goto l21
					}
					goto l20
				}
			}
		}
		if uint32(v1) < uint32(i32(0x20000)) {
			v5 = i32(0)
			v6 = int32(uint32(v1)>>8) & i32(255)
			v2 = i32(0)
		l28:
			{
				v7 = v2 + i32(2)
				t16 := int32(m.memory[int64(uint32(v2))+1109813])
				t17 := v5
				v4 = t16
				v8 = t17 + v4
				{
					t18 := int32(m.memory[int64(uint32(v2))+1109812])
					v2 = t18
					if v2 == v6 {
						goto l26
					}
					if uint32(v2) > uint32(v6) {
						goto l27
					}
					v5 = v8
					v2 = v7
					if v7 != i32(92) {
						goto l28
					}
					goto l27
				}
			l26:
				if uint32(v8) < uint32(v5) {
					goto l29
				}
				if uint32(v8) > uint32(i32(212)) {
					goto l29
				}
				if v4 == 0 {
					goto l30
				}
				v2 = v5 + i32(1109904)
				goto l32
			l29:
				m.fn151(v5, v8, i32(212), i32(1111272))
				panic("unreachable")
			l32:
				{
					t19 := int32(m.memory[uint32(v2)])
					if t19 == v1&i32(255) {
						goto l11
					}
					v2 = v2 + i32(1)
					v4 = v4 + i32(-1)
					if v4 != 0 {
						goto l32
					}
				}
			l30:
				v5 = v8
				v2 = v7
				if v7 != i32(92) {
					goto l28
				}
			}
		l27:
			v5 = v1 & i32(0xffff)
			v4 = i32(1)
			v2 = i32(0)
		l37:
			v7 = v2 + i32(1)
			{
				t20 := int32(int8(m.memory[int64(uint32(v2))+1110116]))
				v8 = t20
				if v8 < i32(0) {
					if v7 == i32(504) {
						m.fn153(i32(1111288))
						panic("unreachable")
					}
					t21 := int32(m.memory[uint32(v2+i32(1110117))])
					v8 = v8&i32(127)<<8 | t21
					v2 = v2 + i32(2)
					goto l34
				}
				v2 = v7
				goto l34
			}
		l34:
			v5 = v5 - v8
			if v5 < i32(0) {
				goto l36
			}
			v4 = v4 ^ i32(1)
			if v2 != i32(504) {
				goto l37
			}
			goto l36
		}
		v2 = v1 & i32(0x1ffffe)
		if v2 == i32(183982) {
			goto l11
		}
		if v1&i32(0x1fffe0) == i32(173792) {
			goto l11
		}
		if v2 == i32(178206) {
			goto l11
		}
		if uint32(v1+i32(-191472)) > uint32(i32(-16)) {
			goto l11
		}
		if uint32(v1+i32(-194560)) > uint32(i32(-2467)) {
			goto l11
		}
		if uint32(v1+i32(-0x30000)) > uint32(i32(-1507)) {
			goto l11
		}
		if uint32(v1+i32(-201552)) > uint32(i32(-6)) {
			goto l11
		}
		if uint32(v1+i32(-917760)) > uint32(i32(-707719)) {
			goto l11
		}
		if uint32(v1) < uint32(i32(918000)) {
			goto l12
		}
		goto l11
	l20:
		v4 = i32(1)
		v5 = v1
		v2 = i32(0)
	l41:
		v7 = v2 + i32(1)
		{
			t22 := int32(int8(m.memory[int64(uint32(v2))+1110980]))
			v8 = t22
			if v8 < i32(0) {
				if v7 == i32(292) {
					m.fn153(i32(1111288))
					panic("unreachable")
				}
				t23 := int32(m.memory[uint32(v2+i32(1110981))])
				v8 = v8&i32(127)<<8 | t23
				v2 = v2 + i32(2)
				goto l39
			}
			v2 = v7
			goto l39
		}
	l39:
		v5 = v5 - v8
		if v5 < i32(0) {
			goto l36
		}
		v4 = v4 ^ i32(1)
		if v2 != i32(292) {
			goto l41
		}
	l36:
		if v4&i32(1) == 0 {
			goto l11
		}
	l12:
		store32(m.memory[uint32(v0):], uint32(v1))
		v1 = i32(129)
		v2 = i32(128)
		goto l16
	l11:
		m.memory[int64(uint32(v3))+24] = byte(i32(0))
		store16(m.memory[int64(uint32(v3))+22:], uint16(i32(0)))
		t24 := int32(m.memory[int64(uint32(int32(uint32(v1)>>20)))+1107936])
		m.memory[int64(uint32(v3))+25] = byte(t24)
		t25 := int32(m.memory[int64(uint32(int32(uint32(v1)>>4)&i32(15)))+1107936])
		m.memory[int64(uint32(v3))+29] = byte(t25)
		t26 := int32(m.memory[int64(uint32(int32(uint32(v1)>>8)&i32(15)))+1107936])
		m.memory[int64(uint32(v3))+28] = byte(t26)
		t27 := int32(m.memory[int64(uint32(int32(uint32(v1)>>12)&i32(15)))+1107936])
		m.memory[int64(uint32(v3))+27] = byte(t27)
		t28 := int32(m.memory[int64(uint32(int32(uint32(v1)>>16)&i32(15)))+1107936])
		m.memory[int64(uint32(v3))+26] = byte(t28)
		t29 := v3 + i32(22)
		v2 = int32(uint32(int32(bits.LeadingZeros32(uint32(v1|i32(1))))) >> 2)
		v4 = t29 + v2
		m.memory[uint32(v4)] = byte(i32(123))
		m.memory[uint32(v4+i32(-1))] = byte(i32(117))
		t30 := v3 + i32(22)
		v2 = v2 + i32(-2)
		m.memory[uint32(t30+v2)] = byte(i32(92))
		t31 := int64(load64(m.memory[int64(uint32(v3))+22:]))
		store64(m.memory[uint32(v0):], uint64(t31))
		m.memory[int64(uint32(v3))+31] = byte(i32(125))
		t32 := int32(m.memory[int64(uint32(v1&i32(15)))+1107936])
		m.memory[int64(uint32(v3))+30] = byte(t32)
		t33 := int32(load16(m.memory[int64(uint32(v3))+30:]))
		store16(m.memory[int64(uint32(v0))+8:], uint16(t33))
	}
l15:
	v1 = i32(10)
l16:
	m.memory[int64(uint32(v0))+13] = byte(v1)
	m.memory[int64(uint32(v0))+12] = byte(v2)
	m.g0 = v3 + i32(32)
}
func (m *Module) fn1646(v0 int32) int32 {
	var v1, v2, v3, v4, v5 int32
	v1 = i32(0)
	p0 := i32(16)
	if uint32(v0) < uint32(i32(69291)) {
		p0 = i32(0)
	}
	v2 = p0
	t1 := v2
	v2 = v2 | i32(8)
	t2 := int32(load32(m.memory[int64(uint32(v2<<2))+1114988:]))
	t3 := v2
	t4 := t2 << 11
	v2 = v0 << 11
	p5 := t3
	if uint32(t4) > uint32(v2) {
		p5 = t1
	}
	v3 = p5
	t6 := v3
	v3 = v3 | i32(4)
	t7 := int32(load32(m.memory[int64(uint32(v3<<2))+1114988:]))
	p8 := v3
	if uint32(t7<<11) > uint32(v2) {
		p8 = t6
	}
	v3 = p8
	t9 := v3
	v3 = v3 | i32(2)
	t10 := int32(load32(m.memory[int64(uint32(v3<<2))+1114988:]))
	p11 := v3
	if uint32(t10<<11) > uint32(v2) {
		p11 = t9
	}
	v3 = p11
	t12 := v3
	v3 = v3 + i32(1)
	t13 := int32(load32(m.memory[int64(uint32(v3<<2))+1114988:]))
	p14 := v3
	if uint32(t13<<11) > uint32(v2) {
		p14 = t12
	}
	v3 = p14
	t15 := v3
	v3 = v3 + i32(1)
	t16 := int32(load32(m.memory[int64(uint32(v3<<2))+1114988:]))
	p17 := v3
	if uint32(t16<<11) > uint32(v2) {
		p17 = t15
	}
	v3 = p17
	t18 := int32(load32(m.memory[int64(uint32(v3<<2))+1114988:]))
	v4 = t18 << 11
	var p19 int32
	if v4 == v2 {
		p19 = 1
	}
	var p20 int32
	if uint32(v4) < uint32(v2) {
		p20 = 1
	}
	v3 = p19 + p20 + v3
	v2 = v3 << 2
	v5 = v2 + i32(1114988)
	t21 := int32(load32(m.memory[int64(uint32(v2))+1114988:]))
	v2 = int32(uint32(t21) >> 21)
	v4 = i32(767)
	{
		{
			if uint32(v3) > uint32(i32(31)) {
				goto l0
			}
			t22 := int32(load32(m.memory[int64(uint32(v5))+4:]))
			v4 = int32(uint32(t22) >> 21)
			if v3 == 0 {
				goto l1
			}
		}
	l0:
		t23 := int32(load32(m.memory[uint32(v5+i32(-4)):]))
		v1 = t23 & i32(0x1fffff)
	}
l1:
	if v4+(v2^i32(-1)) == 0 {
		goto l2
	}
	v3 = v0 - v1
	v4 = v4 + i32(-1)
	v0 = i32(0)
l3:
	{
		t24 := int32(m.memory[uint32(v2+i32(1106478))])
		v0 = v0 + t24
		if uint32(v0) > uint32(v3) {
			goto l2
		}
		t25 := v4
		v2 = v2 + i32(1)
		if t25 != v2 {
			goto l3
		}
	}
l2:
	return v2 & i32(1)
}
func (m *Module) fn1647(v0, v1, v2, v3, v4 int32) int32 {
	var v5, v6, v7, v8 int32
	t0 := m.g0
	v5 = t0 - i32(32)
	m.g0 = v5
	v6 = i32(1)
	{
		t1 := int32(m.memory[int64(uint32(v0))+4])
		if t1 != 0 {
			goto l0
		}
		t2 := int32(m.memory[int64(uint32(v0))+5])
		v7 = t2
		{
			t3 := int32(load32(m.memory[uint32(v0):]))
			v8 = t3
			t4 := int32(m.memory[int64(uint32(v8))+10])
			if t4&i32(128) != 0 {
				goto l1
			}
			v6 = i32(1)
			t5 := int32(load32(m.memory[uint32(v8):]))
			v7 = v7 & i32(1)
			p6 := i32(1108151)
			if v7 != 0 {
				p6 = i32(1108154)
			}
			p7 := i32(3)
			if v7 != 0 {
				p7 = i32(2)
			}
			t8 := int32(load32(m.memory[int64(uint32(v8))+4:]))
			t9 := int32(load32(m.memory[int64(uint32(t8))+12:]))
			t10 := m.t0[uint(t9)].(func(int32, int32, int32) int32)(t5, p6, p7)
			if t10 != 0 {
				goto l0
			}
			t11 := int32(load32(m.memory[uint32(v8):]))
			t12 := int32(load32(m.memory[int64(uint32(v8))+4:]))
			t13 := int32(load32(m.memory[int64(uint32(t12))+12:]))
			t14 := m.t0[uint(t13)].(func(int32, int32, int32) int32)(t11, v1, v2)
			if t14 != 0 {
				goto l0
			}
			t15 := int32(load32(m.memory[uint32(v8):]))
			t16 := int32(load32(m.memory[int64(uint32(v8))+4:]))
			t17 := int32(load32(m.memory[int64(uint32(t16))+12:]))
			t18 := m.t0[uint(t17)].(func(int32, int32, int32) int32)(t15, i32(1108156), i32(2))
			if t18 != 0 {
				goto l0
			}
			t19 := m.t0[uint(v4)].(func(int32, int32) int32)(v3, v8)
			v6 = t19
			goto l0
		}
	l1:
		v6 = i32(1)
		{
			if v7&i32(1) != 0 {
				goto l2
			}
			t20 := int32(load32(m.memory[uint32(v8):]))
			t21 := int32(load32(m.memory[int64(uint32(v8))+4:]))
			t22 := int32(load32(m.memory[int64(uint32(t21))+12:]))
			t23 := m.t0[uint(t22)].(func(int32, int32, int32) int32)(t20, i32(1108158), i32(3))
			if t23 != 0 {
				goto l0
			}
		}
	l2:
		v6 = i32(1)
		m.memory[int64(uint32(v5))+15] = byte(i32(1))
		store32(m.memory[int64(uint32(v5))+20:], uint32(i32(1109040)))
		t24 := int64(load64(m.memory[uint32(v8):]))
		store64(m.memory[uint32(v5):], uint64(t24))
		t25 := int64(load64(m.memory[int64(uint32(v8))+8:]))
		store64(m.memory[int64(uint32(v5))+24:], uint64(t25))
		store32(m.memory[int64(uint32(v5))+8:], uint32(v5+i32(15)))
		store32(m.memory[int64(uint32(v5))+16:], uint32(v5))
		t26 := m.fn1639(v5, v1, v2)
		if t26 != 0 {
			goto l0
		}
		t27 := m.fn1639(v5, i32(1108156), i32(2))
		if t27 != 0 {
			goto l0
		}
		{
			t28 := m.t0[uint(v4)].(func(int32, int32) int32)(v3, v5+i32(16))
			if t28 == 0 {
				goto l3
			}
			v6 = i32(1)
			goto l0
		}
	l3:
		t29 := int32(load32(m.memory[int64(uint32(v5))+16:]))
		t30 := int32(load32(m.memory[int64(uint32(v5))+20:]))
		t31 := int32(load32(m.memory[int64(uint32(t30))+12:]))
		t32 := m.t0[uint(t31)].(func(int32, int32, int32) int32)(t29, i32(1108161), i32(2))
		v6 = t32
	}
l0:
	m.memory[int64(uint32(v0))+5] = byte(i32(1))
	m.memory[int64(uint32(v0))+4] = byte(v6)
	m.g0 = v5 + i32(32)
	return v0
}
func (m *Module) fn1648(v0 int32) int32 {
	var v1, v2 int32
	t0 := int32(m.memory[int64(uint32(v0))+4])
	v1 = t0
	v2 = v1
	{
		t1 := int32(m.memory[int64(uint32(v0))+5])
		if t1 == 0 {
			goto l0
		}
		v2 = i32(1)
		{
			if v1&i32(1) != 0 {
				goto l1
			}
			{
				t2 := int32(load32(m.memory[uint32(v0):]))
				v2 = t2
				t3 := int32(m.memory[int64(uint32(v2))+10])
				if t3&i32(128) != 0 {
					goto l2
				}
				t4 := int32(load32(m.memory[uint32(v2):]))
				t5 := int32(load32(m.memory[int64(uint32(v2))+4:]))
				t6 := int32(load32(m.memory[int64(uint32(t5))+12:]))
				t7 := m.t0[uint(t6)].(func(int32, int32, int32) int32)(t4, i32(1283984), i32(2))
				v2 = t7
				goto l1
			}
		l2:
			t8 := int32(load32(m.memory[uint32(v2):]))
			t9 := int32(load32(m.memory[int64(uint32(v2))+4:]))
			t10 := int32(load32(m.memory[int64(uint32(t9))+12:]))
			t11 := m.t0[uint(t10)].(func(int32, int32, int32) int32)(t8, i32(1108167), i32(1))
			v2 = t11
		}
	l1:
		m.memory[int64(uint32(v0))+4] = byte(v2)
	}
l0:
	return v2 & i32(1)
}
func (m *Module) fn1649(v0, v1, v2 int32) int32 {
	var v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14 int32
	var v15, v16 int64
	t0 := m.g0
	v3 = t0 - i32(160)
	m.g0 = v3
	v4 = i32(0)
	memory_zero(m.memory, uint32(v3), uint32(i32(160)))
	{
		t1 := int32(load32(m.memory[int64(uint32(v0))+160:]))
		v5 = t1
		if uint32(v5) < uint32(v2) {
			v6 = v2 + i32(1)
			v7 = v2 << 2
			v14 = v0 + v5<<2
			v11 = v0
			v10 = i32(0)
		l14:
			v9 = v3 + v4<<2
		l9:
			{
				v13 = v4
				v12 = v9
				if v11 == v14 {
					goto l2
				}
				v9 = v12 + i32(4)
				v4 = v13 + i32(1)
				t8 := int32(load32(m.memory[uint32(v11):]))
				v8 = t8
				v5 = v11 + i32(4)
				v11 = v5
				if v8 == 0 {
					goto l9
				}
			}
			v15 = int64(uint32(v8))
			v16 = i64(0)
			v8 = v7
			v11 = v13
			v9 = v1
		l11:
			{
				if uint32(v11) >= uint32(i32(40)) {
					m.fn158(v11, i32(40), i32(1108992))
					panic("unreachable")
				}
				t9 := int64(load32(m.memory[uint32(v12):]))
				t10 := int64(load32(m.memory[uint32(v9):]))
				t11 := v12
				v16 = v16 + t9 + t10*v15
				store32(m.memory[uint32(t11):], uint32(v16))
				v16 = int64(uint64(v16) >> 32)
				v12 = v12 + i32(4)
				v11 = v11 + i32(1)
				v9 = v9 + i32(4)
				v8 = v8 + i32(-4)
				if v8 != 0 {
					goto l11
				}
			}
			v12 = v2
			{
				if v16 == 0 {
					goto l12
				}
				v12 = v13 + v2
				if uint32(v12) >= uint32(i32(40)) {
					m.fn158(v12, i32(40), i32(1108992))
					panic("unreachable")
				}
				store32(m.memory[uint32(v3+v12<<2):], uint32(int32(v16)))
				v12 = v6
			l12:
				t12 := v10
				v12 = v12 + v13
				p13 := v12
				if uint32(v10) > uint32(v12) {
					p13 = t12
				}
				v10 = p13
				v11 = v5
				goto l14
			}
		}
		if uint32(v5) >= uint32(i32(41)) {
			m.fn151(i32(0), v5, i32(40), i32(1108992))
			panic("unreachable")
		}
		v6 = v5 + i32(1)
		v7 = v5 << 2
		v8 = v1 + v2<<2
		v9 = i32(0)
		v10 = i32(0)
	l8:
		v11 = v3 + v9<<2
	l3:
		{
			v4 = v9
			v12 = v11
			if v1 == v8 {
				goto l2
			}
			v11 = v12 + i32(4)
			v9 = v4 + i32(1)
			t2 := int32(load32(m.memory[uint32(v1):]))
			v13 = t2
			v14 = v1 + i32(4)
			v1 = v14
			if v13 == 0 {
				goto l3
			}
		}
		v15 = int64(uint32(v13))
		v16 = i64(0)
		v13 = v7
		v1 = v4
		v11 = v0
	l5:
		{
			if uint32(v1) >= uint32(i32(40)) {
				m.fn158(v1, i32(40), i32(1108992))
				panic("unreachable")
			}
			t3 := int64(load32(m.memory[uint32(v12):]))
			t4 := int64(load32(m.memory[uint32(v11):]))
			t5 := v12
			v16 = v16 + t3 + t4*v15
			store32(m.memory[uint32(t5):], uint32(v16))
			v16 = int64(uint64(v16) >> 32)
			v12 = v12 + i32(4)
			v1 = v1 + i32(1)
			v11 = v11 + i32(4)
			v13 = v13 + i32(-4)
			if v13 != 0 {
				goto l5
			}
		}
		v12 = v5
		{
			if v16 == 0 {
				goto l6
			}
			v12 = v4 + v5
			if uint32(v12) >= uint32(i32(40)) {
				m.fn158(v12, i32(40), i32(1108992))
				panic("unreachable")
			}
			store32(m.memory[uint32(v3+v12<<2):], uint32(int32(v16)))
			v12 = v6
		l6:
			t6 := v10
			v12 = v12 + v4
			p7 := v12
			if uint32(v10) > uint32(v12) {
				p7 = t6
			}
			v10 = p7
			v1 = v14
			goto l8
		}
	}
l2:
	memory_copy(m.memory, uint32(v0), uint32(v3), uint32(i32(160)))
	store32(m.memory[int64(uint32(v0))+160:], uint32(v10))
	m.g0 = v3 + i32(160)
	return v0
}
func (m *Module) fn1650(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6, v7, v8 int32
	if uint32(v1) >= uint32(i32(1280)) {
		m.fn256(i32(1109008), i32(29), i32(1108992))
		panic("unreachable")
	}
	v2 = int32(uint32(v1) >> 5)
	{
		{
			t0 := int32(load32(m.memory[int64(uint32(v0))+160:]))
			v3 = t0
			if v3 == 0 {
				goto l1
			}
			v4 = v3 + i32(-1)
			v5 = v3<<2 + v0 + i32(-4)
			v6 = (v3+v2)<<2 + v0 + i32(-4)
			var p1 int32
			if uint32(v3) < uint32(i32(41)) {
				p1 = 1
			}
			v3 = p1
		l4:
			{
				if v3 == 0 {
					m.fn158(v4, i32(40), i32(1108992))
					panic("unreachable")
				}
				v7 = v2 + v4
				if uint32(v7) >= uint32(i32(40)) {
					m.fn158(v7, i32(40), i32(1108992))
					panic("unreachable")
				}
				t2 := int32(load32(m.memory[uint32(v5):]))
				store32(m.memory[uint32(v6):], uint32(t2))
				v6 = v6 + i32(-4)
				v5 = v5 + i32(-4)
				v4 = v4 + i32(-1)
				if v4 != i32(-1) {
					goto l4
				}
			}
		}
	l1:
		v3 = v1 & i32(31)
		if v2 == 0 {
			goto l5
		}
		v4 = v2 << 2
		if v4 == 0 {
			goto l5
		}
		memory_zero(m.memory, uint32(v0), uint32(v4))
	l5:
		t3 := int32(load32(m.memory[int64(uint32(v0))+160:]))
		v5 = t3 + v2
		if v3 != 0 {
			v4 = v5 + i32(-1)
			if uint32(v4) > uint32(i32(39)) {
				m.fn158(v4, i32(40), i32(1108992))
				panic("unreachable")
			}
			v8 = v5
			t4 := int32(load32(m.memory[uint32(v0+v4<<2):]))
			v7 = i32(32) - v3
			v4 = i32_shr_u(t4, v7)
			if v4 == 0 {
				goto l8
			}
			if uint32(v5) > uint32(i32(39)) {
				m.fn158(v5, i32(40), i32(1108992))
				panic("unreachable")
			}
			store32(m.memory[uint32(v0+v5<<2):], uint32(v4))
			v8 = v5 + i32(1)
			goto l8
		}
		store32(m.memory[int64(uint32(v0))+160:], uint32(v5))
		return v0
	}
l8:
	v1 = v2 + i32(1)
	if uint32(v1) >= uint32(v5) {
		goto l10
	}
	v4 = v5<<2 + v0 + i32(-8)
l11:
	{
		v6 = v4 + i32(4)
		t5 := int32(load32(m.memory[uint32(v4):]))
		t6 := int32(load32(m.memory[uint32(v6):]))
		store32(m.memory[uint32(v6):], uint32(i32_shr_u(t5, v7)|i32_shl(t6, v3)))
		v4 = v4 + i32(-4)
		t7 := v1
		v5 = v5 + i32(-1)
		if uint32(t7) < uint32(v5) {
			goto l11
		}
	}
l10:
	v4 = v0 + v2<<2
	t8 := int32(load32(m.memory[uint32(v4):]))
	store32(m.memory[uint32(v4):], uint32(i32_shl(t8, v3)))
	store32(m.memory[int64(uint32(v0))+160:], uint32(v8))
	return v0
}
func (m *Module) fn1651(v0, v1, v2 int32) int32 {
	var v3, v4, v5, v6 int32
	t0 := m.g0
	v3 = t0 - i32(32)
	m.g0 = v3
	v4 = i32(1)
	{
		t1 := int32(m.memory[int64(uint32(v0))+4])
		if t1 != 0 {
			goto l0
		}
		t2 := int32(m.memory[int64(uint32(v0))+5])
		v5 = t2
		{
			t3 := int32(load32(m.memory[uint32(v0):]))
			v6 = t3
			t4 := int32(m.memory[int64(uint32(v6))+10])
			if t4&i32(128) != 0 {
				v4 = i32(1)
				{
					if v5&i32(1) != 0 {
						goto l3
					}
					t9 := int32(load32(m.memory[uint32(v6):]))
					t10 := int32(load32(m.memory[int64(uint32(v6))+4:]))
					t11 := int32(load32(m.memory[int64(uint32(t10))+12:]))
					t12 := m.t0[uint(t11)].(func(int32, int32, int32) int32)(t9, i32(1108166), i32(1))
					if t12 != 0 {
						goto l0
					}
				}
			l3:
				v4 = i32(1)
				m.memory[int64(uint32(v3))+15] = byte(i32(1))
				store32(m.memory[int64(uint32(v3))+20:], uint32(i32(1109040)))
				t13 := int64(load64(m.memory[uint32(v6):]))
				store64(m.memory[uint32(v3):], uint64(t13))
				t14 := int64(load64(m.memory[int64(uint32(v6))+8:]))
				store64(m.memory[int64(uint32(v3))+24:], uint64(t14))
				store32(m.memory[int64(uint32(v3))+8:], uint32(v3+i32(15)))
				store32(m.memory[int64(uint32(v3))+16:], uint32(v3))
				t15 := m.t0[uint(v2)].(func(int32, int32) int32)(v1, v3+i32(16))
				if t15 != 0 {
					goto l0
				}
				t16 := int32(load32(m.memory[int64(uint32(v3))+16:]))
				t17 := int32(load32(m.memory[int64(uint32(v3))+20:]))
				t18 := int32(load32(m.memory[int64(uint32(t17))+12:]))
				t19 := m.t0[uint(t18)].(func(int32, int32, int32) int32)(t16, i32(1108161), i32(2))
				v4 = t19
				goto l0
			}
			v4 = i32(1)
			if v5&i32(1) == 0 {
				goto l2
			}
			t5 := int32(load32(m.memory[uint32(v6):]))
			t6 := int32(load32(m.memory[int64(uint32(v6))+4:]))
			t7 := int32(load32(m.memory[int64(uint32(t6))+12:]))
			t8 := m.t0[uint(t7)].(func(int32, int32, int32) int32)(t5, i32(1108154), i32(2))
			if t8 == 0 {
				goto l2
			}
			goto l0
		}
	l2:
		t20 := m.t0[uint(v2)].(func(int32, int32) int32)(v1, v6)
		v4 = t20
	}
l0:
	m.memory[int64(uint32(v0))+5] = byte(i32(1))
	m.memory[int64(uint32(v0))+4] = byte(v4)
	m.g0 = v3 + i32(32)
	return v0
}
func (m *Module) fn1652(v0 int32) int32 {
	var v1 int32
	v1 = i32(1)
	{
		t0 := int32(m.memory[int64(uint32(v0))+4])
		if t0 != 0 {
			goto l0
		}
		t1 := int32(load32(m.memory[uint32(v0):]))
		v1 = t1
		t2 := int32(load32(m.memory[uint32(v1):]))
		t3 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t4 := int32(load32(m.memory[int64(uint32(t3))+12:]))
		t5 := m.t0[uint(t4)].(func(int32, int32, int32) int32)(t2, i32(1108169), i32(1))
		v1 = t5
	}
l0:
	m.memory[int64(uint32(v0))+4] = byte(v1)
	return v1
}
func (m *Module) fn1653(v0, v1 int32) {
	var v2, v3, v4, v5, v6, v7, v8, v9, v10 int32
	var v11, v12, v13, v14 int64
	{
		t0 := int32(load32(m.memory[uint32(v0):]))
		v2 = t0
		if v2 == 0 {
			return
		}
		v3 = v1 & i32(63)
		v1 = v3 << 1
		t1 := int32(load16(m.memory[int64(uint32(v1))+1116840:]))
		v4 = t1
		v5 = v4 & i32(2047)
		t2 := int32(load16(m.memory[int64(uint32(v1))+1116842:]))
		v6 = v5 - t2&i32(2047)
		v7 = i32(0) - v2
		v8 = v0 + i32(8)
		v4 = int32(uint32(v4) >> 11)
		v1 = i32(-1308)
		{
		l4:
			{
				if v6+v1 == i32(-1308) {
					goto l1
				}
				v9 = v5 + v1
				if v9 == 0 {
					goto l1
				}
				if v7+v1 == i32(-1308) {
					v4 = v4 + i32(-1)
					goto l1
				}
				if v1 == i32(-540) {
					m.fn158(i32(768), i32(768), i32(1118280))
					panic("unreachable")
				}
				v10 = v8 + v1
				v1 = v1 + i32(1)
				t3 := int32(m.memory[uint32(v10+i32(1308))])
				v10 = t3
				t4 := int32(m.memory[uint32(v9+i32(1118278))])
				t5 := v10
				v9 = t4
				if t5 == v9&i32(255) {
					goto l4
				}
			}
			t6 := v4
			var p7 int32
			if uint32(v10) < uint32(v9&i32(255)) {
				p7 = 1
			}
			v4 = t6 - p7
			goto l1
		}
	l1:
		v9 = v0 + i32(7)
		v10 = v9 + v4
		v11 = int64(uint32(v3))
		v12 = i64(0)
	l8:
		v1 = v2
		v2 = v1 + i32(-1)
		{
			if uint32(v1) >= uint32(i32(769)) {
				m.fn158(v2, i32(768), i32(1109080))
				panic("unreachable")
			}
			t8 := int64(m.memory[uint32(v9+v1)])
			v13 = i64_shl(t8, v11) + v12
			t9 := int64(uint64(v13) / uint64(i64(10)))
			t10 := v13
			v12 = t9
			v14 = t10 + v12*i64(-10)
			if uint32(v2+v4) < uint32(i32(768)) {
				goto l6
			}
			if v14 == 0 {
				goto l7
			}
			m.memory[int64(uint32(v0))+776] = byte(i32(1))
			goto l7
		}
	l6:
		m.memory[uint32(v10+v1)] = byte(v14)
	l7:
		if v2 != 0 {
			goto l8
		}
		if uint64(v13) < uint64(i64(10)) {
			goto l9
		}
		v2 = v4 + i32(7)
	l12:
		{
			v13 = v12
			t11 := int64(uint64(v13) / uint64(i64(10)))
			t12 := v13
			v12 = t11
			v14 = t12 + v12*i64(-10)
			if uint32(v2+i32(-8)) < uint32(i32(768)) {
				goto l10
			}
			if v14 == 0 {
				goto l11
			}
			m.memory[int64(uint32(v0))+776] = byte(i32(1))
			goto l11
		l10:
			m.memory[uint32(v0+v2)] = byte(v14)
		l11:
			v2 = v2 + i32(-1)
			if uint64(v13) >= uint64(i64(10)) {
				goto l12
			}
		}
	l9:
		t13 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		store32(m.memory[int64(uint32(v0))+4:], uint32(t13+v4))
		t14 := int32(load32(m.memory[uint32(v0):]))
		t15 := v0
		v1 = t14 + v4
		p16 := i32(768)
		if uint32(v1) < uint32(i32(768)) {
			p16 = v1
		}
		v2 = p16
		store32(m.memory[uint32(t15):], uint32(v2))
		if v1 == 0 {
			return
		}
	l15:
		v1 = v2 + i32(-1)
		{
			if uint32(v2) > uint32(i32(768)) {
				m.fn158(v1, i32(768), i32(1109064))
				panic("unreachable")
			}
			t17 := int32(m.memory[uint32(v0+v2+i32(7))])
			if t17 == 0 {
				goto l14
			}
			return
		}
	l14:
		store32(m.memory[uint32(v0):], uint32(v1))
		v2 = v1
		if v1 != 0 {
			goto l15
		}
	}
}
func (m *Module) fn1654(v0, v1 int32) {
	var v2 int32
	var v3 int64
	var v4 int32
	var v5 int64
	var v6 int32
	var v7 int64
	var v8, v9, v10, v11 int32
	var v12 int64
	v2 = v0 + i32(8)
	v3 = int64(uint32(v1 & i32(63)))
	t0 := int32(load32(m.memory[uint32(v0):]))
	v4 = t0
	v5 = i64(0)
	v1 = i32(0)
	{
	l6:
		if v4 != v1 {
			goto l0
		}
		if v5 == 0 {
			return
		}
		if i64_shr_u(v5, v3) == i64(0) {
			v1 = v4
		l4:
			v1 = v1 + i32(1)
			v5 = v5 * i64(10)
			if i64_shr_u(v5, v3) == 0 {
				goto l4
			}
			goto l3
		}
		v1 = v4
		goto l3
	l0:
		{
			if v1 == i32(768) {
				goto l5
			}
			v6 = v0 + v1
			v1 = v1 + i32(1)
			t1 := int64(m.memory[uint32(v6+i32(8))])
			v5 = v5*i64(10) + t1
			if i64_shr_u(v5, v3) == 0 {
				goto l6
			}
			goto l3
		}
	l5:
		m.fn158(i32(768), i32(768), i32(1109096))
		panic("unreachable")
	l3:
		t2 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		t3 := v0
		v6 = t2 - v1 + i32(1)
		store32(m.memory[int64(uint32(t3))+4:], uint32(v6))
		{
			if v6 < i32(-2047) {
				goto l7
			}
			v7 = i64_shl(i64(-1), v3) ^ i64(-1)
			v6 = i32(0)
			{
				if uint32(v4) <= uint32(v1) {
					goto l8
				}
				v8 = i32(0)
				v6 = i32(768) - v1
				p4 := v6
				if uint32(v6) > uint32(i32(768)) {
					p4 = i32(0)
				}
				v9 = p4
				v10 = v1 - v4
				v11 = v2 + v1
				v6 = v4 - v1
			l10:
				{
					if v9 != v8 {
						goto l9
					}
					m.fn158(v1+v8, i32(768), i32(1109112))
					panic("unreachable")
				l9:
					t5 := int64(m.memory[uint32(v11+v8)])
					v12 = t5
					m.memory[uint32(v2+v8)] = byte(i64_shr_u(v5, v3))
					v5 = v12 + v5&v7*i64(10)
					t6 := v10
					v8 = v8 + i32(1)
					if t6+v8 != 0 {
						goto l10
					}
				}
			}
		l8:
			if v5 == 0 {
				goto l11
			}
		l14:
			v12 = v5
			v5 = v12 & v7 * i64(10)
			v1 = int32(i64_shr_u(v12, v3))
			if uint32(v6) < uint32(i32(768)) {
				goto l12
			}
			if v1&i32(255) == 0 {
				goto l13
			}
			m.memory[int64(uint32(v0))+776] = byte(i32(1))
			goto l13
		l12:
			m.memory[uint32(v2+v6)] = byte(v1)
			v6 = v6 + i32(1)
		l13:
			if !(v5 == 0) {
				goto l14
			}
		l11:
			;
			var p7 int32
			if uint32(v6) > uint32(i32(768)) {
				p7 = 1
			}
			v4 = p7
		l16:
			store32(m.memory[uint32(v0):], uint32(v6))
			if v6 == 0 {
				return
			}
			v1 = v6 + i32(-1)
			{
				if v4 != 0 {
					m.fn158(v1, i32(768), i32(1109064))
					panic("unreachable")
				}
				v8 = v0 + v6
				v6 = v1
				t8 := int32(m.memory[uint32(v8+i32(7))])
				if t8 == 0 {
					goto l16
				}
				return
			}
		}
	l7:
		m.memory[int64(uint32(v0))+776] = byte(i32(0))
		store64(m.memory[uint32(v0):], uint64(i64(0)))
	}
}
func (m *Module) fn1655(v0 int32) int64 {
	var v1 int64
	var v2, v3, v4, v5 int32
	v1 = i64(0)
	{
		t0 := int32(load32(m.memory[uint32(v0):]))
		v2 = t0
		if v2 == 0 {
			goto l0
		}
		t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		v3 = t1
		if v3 < i32(0) {
			goto l0
		}
		v1 = i64(-1)
		if uint32(v3) > uint32(i32(18)) {
			goto l0
		}
		if v3 != 0 {
			goto l1
		}
		v1 = i64(0)
		goto l2
	l1:
		v4 = v0 + i32(8)
		v5 = i32(0)
		v1 = i64(0)
	l4:
		{
			v1 = v1 * i64(10)
			{
				if uint32(v5) >= uint32(v2) {
					goto l3
				}
				t2 := int64(m.memory[uint32(v4+v5)])
				v1 = v1 + t2
			}
		l3:
			t3 := v3
			v5 = v5 + i32(1)
			if t3 != v5 {
				goto l4
			}
		}
	l2:
		if uint32(v3) >= uint32(v2) {
			goto l0
		}
		v4 = v0 + v3
		t4 := int32(m.memory[int64(uint32(v4))+8])
		v5 = t4
		{
			if v3+i32(1) != v2 {
				goto l5
			}
			if v5&i32(255) == i32(5) {
				goto l6
			}
		l5:
			if uint32(v5&i32(255)) > uint32(i32(4)) {
				goto l7
			}
			goto l0
		l6:
			t5 := int32(m.memory[int64(uint32(v0))+776])
			if t5 != 0 {
				goto l7
			}
			if v3 == 0 {
				goto l0
			}
			t6 := int32(m.memory[uint32(v4+i32(8)+i32(-1))])
			if t6&i32(1) == 0 {
				goto l0
			}
		}
	l7:
		v1 = v1 + i64(1)
	}
l0:
	return v1
}
func (m *Module) fn1656(v0, v1, v2, v3, v4 int32) {
	var v5, v6 int32
	var v7 int64
	t0 := m.g0
	v5 = t0 - i32(48)
	m.g0 = v5
	store32(m.memory[int64(uint32(v5))+4:], uint32(v3))
	store32(m.memory[uint32(v5):], uint32(v2))
	store32(m.memory[int64(uint32(v5))+8:], uint32(v1))
	{
		{
			if uint32(v2) > uint32(v1) {
				t5 := v5
				v7 = int64(uint32(i32(5))) << 32
				store64(m.memory[int64(uint32(t5))+32:], uint64(v7|int64(uint32(v5+i32(8)))))
				store64(m.memory[int64(uint32(v5))+24:], uint64(v7|int64(uint32(v5))))
				m.fn91(i32(1050351), v5+i32(24), v4)
				panic("unreachable")
			}
			if uint32(v3) > uint32(v1) {
				t6 := v5
				v7 = int64(uint32(i32(5))) << 32
				store64(m.memory[int64(uint32(t6))+32:], uint64(v7|int64(uint32(v5+i32(8)))))
				store64(m.memory[int64(uint32(v5))+24:], uint64(v7|int64(uint32(v5+i32(4)))))
				m.fn91(i32(1050412), v5+i32(24), v4)
				panic("unreachable")
			}
			if uint32(v2) > uint32(v3) {
				t7 := v5
				v7 = int64(uint32(i32(5))) << 32
				store64(m.memory[int64(uint32(t7))+32:], uint64(v7|int64(uint32(v5+i32(4)))))
				store64(m.memory[int64(uint32(v5))+24:], uint64(v7|int64(uint32(v5))))
				m.fn91(i32(1049768), v5+i32(24), v4)
				panic("unreachable")
			}
			if v2 == 0 {
				goto l3
			}
			if uint32(v2) >= uint32(v1) {
				goto l3
			}
			t1 := int32(int8(m.memory[uint32(v0+v2)]))
			if t1 > i32(-65) {
				goto l3
			}
			v6 = v2
		l5:
			{
				t2 := int32(int8(m.memory[uint32(v0+v6)]))
				if t2 > i32(-65) {
					goto l7
				}
				v6 = v6 + i32(-1)
				if v6 != 0 {
					goto l5
				}
			}
			v6 = i32(0)
		l7:
			{
				t3 := int32(int8(m.memory[uint32(v0+v2)]))
				if t3 > i32(-65) {
					goto l6
				}
				t4 := v1
				v2 = v2 + i32(1)
				if t4 != v2 {
					goto l7
				}
			}
			v2 = v1
			goto l6
		}
	l6:
		store32(m.memory[int64(uint32(v5))+12:], uint32(v6))
		store32(m.memory[int64(uint32(v5))+16:], uint32(v2))
		if uint32(v6) > uint32(v2) {
			goto l8
		}
		{
			if v6 == 0 {
				goto l9
			}
			if uint32(v6) < uint32(v1) {
				goto l10
			}
			if v6 == v1 {
				goto l9
			}
			goto l8
		l10:
			t8 := int32(int8(m.memory[uint32(v0+v6)]))
			if t8 < i32(-64) {
				goto l8
			}
		}
	l9:
		{
			if uint32(v2) < uint32(v1) {
				goto l11
			}
			if v2 != v1 {
				goto l8
			}
			goto l12
		l11:
			t9 := int32(int8(m.memory[uint32(v0+v2)]))
			if t9 <= i32(-65) {
				goto l8
			}
		}
	l12:
		if v6 == v2 {
			goto l13
		}
		{
			{
				v0 = v0 + v6
				t10 := int32(int8(m.memory[uint32(v0)]))
				v6 = t10
				if v6 <= i32(-1) {
					goto l14
				}
				v6 = v6 & i32(255)
				goto l15
			}
		l14:
			t11 := int32(m.memory[int64(uint32(v0))+1])
			v3 = t11 & i32(63)
			v2 = v6 & i32(31)
			if uint32(v6) > uint32(i32(-33)) {
				goto l16
			}
			v6 = v2<<6 | v3
			goto l15
		l16:
			t12 := int32(m.memory[int64(uint32(v0))+2])
			v3 = v3<<6 | t12&i32(63)
			if uint32(v6) >= uint32(i32(-16)) {
				goto l17
			}
			v6 = v3 | v2<<12
			goto l15
		l17:
			t13 := int32(m.memory[int64(uint32(v0))+3])
			v6 = v3<<6 | t13&i32(63) | v2<<18&i32(0x1c0000)
		}
	l15:
		store32(m.memory[int64(uint32(v5))+20:], uint32(v6))
		store64(m.memory[int64(uint32(v5))+40:], uint64(int64(uint32(i32(199)))<<32|int64(uint32(v5+i32(12)))))
		store64(m.memory[int64(uint32(v5))+32:], uint64(int64(uint32(i32(200)))<<32|int64(uint32(v5+i32(20)))))
		store64(m.memory[int64(uint32(v5))+24:], uint64(int64(uint32(i32(5)))<<32|int64(uint32(v5))))
		m.fn91(i32(1068719), v5+i32(24), v4)
		panic("unreachable")
	l8:
		m.fn556(v0, v1, v6, v2, v4)
		panic("unreachable")
	l3:
		if v3 == 0 {
			goto l18
		}
		if uint32(v3) >= uint32(v1) {
			goto l18
		}
		t14 := int32(int8(m.memory[uint32(v0+v3)]))
		if t14 > i32(-65) {
			goto l18
		}
		v6 = v3
	l20:
		{
			t15 := int32(int8(m.memory[uint32(v0+v6)]))
			if t15 > i32(-65) {
				goto l22
			}
			v6 = v6 + i32(-1)
			if v6 != 0 {
				goto l20
			}
		}
		v6 = i32(0)
	l22:
		{
			t16 := int32(int8(m.memory[uint32(v0+v3)]))
			if t16 > i32(-65) {
				goto l21
			}
			t17 := v1
			v3 = v3 + i32(1)
			if t17 != v3 {
				goto l22
			}
		}
		v3 = v1
	l21:
		store32(m.memory[int64(uint32(v5))+12:], uint32(v6))
		store32(m.memory[int64(uint32(v5))+16:], uint32(v3))
		if uint32(v6) > uint32(v3) {
			goto l23
		}
		{
			if v6 == 0 {
				goto l24
			}
			if uint32(v6) < uint32(v1) {
				goto l25
			}
			if v6 == v1 {
				goto l24
			}
			goto l23
		l25:
			t18 := int32(int8(m.memory[uint32(v0+v6)]))
			if t18 < i32(-64) {
				goto l23
			}
		}
	l24:
		{
			if uint32(v3) < uint32(v1) {
				goto l26
			}
			if v3 != v1 {
				goto l23
			}
			goto l27
		l26:
			t19 := int32(int8(m.memory[uint32(v0+v3)]))
			if t19 <= i32(-65) {
				goto l23
			}
		}
	l27:
		if v6 == v3 {
			goto l13
		}
		{
			{
				v0 = v0 + v6
				t20 := int32(int8(m.memory[uint32(v0)]))
				v6 = t20
				if v6 <= i32(-1) {
					goto l28
				}
				v6 = v6 & i32(255)
				goto l29
			}
		l28:
			t21 := int32(m.memory[int64(uint32(v0))+1])
			v3 = t21 & i32(63)
			v2 = v6 & i32(31)
			if uint32(v6) > uint32(i32(-33)) {
				goto l30
			}
			v6 = v2<<6 | v3
			goto l29
		l30:
			t22 := int32(m.memory[int64(uint32(v0))+2])
			v3 = v3<<6 | t22&i32(63)
			if uint32(v6) >= uint32(i32(-16)) {
				goto l31
			}
			v6 = v3 | v2<<12
			goto l29
		l31:
			t23 := int32(m.memory[int64(uint32(v0))+3])
			v6 = v3<<6 | t23&i32(63) | v2<<18&i32(0x1c0000)
		}
	l29:
		store32(m.memory[int64(uint32(v5))+20:], uint32(v6))
		store64(m.memory[int64(uint32(v5))+40:], uint64(int64(uint32(i32(199)))<<32|int64(uint32(v5+i32(12)))))
		store64(m.memory[int64(uint32(v5))+32:], uint64(int64(uint32(i32(200)))<<32|int64(uint32(v5+i32(20)))))
		store64(m.memory[int64(uint32(v5))+24:], uint64(int64(uint32(i32(5)))<<32|int64(uint32(v5+i32(4)))))
		m.fn91(i32(1068801), v5+i32(24), v4)
		panic("unreachable")
	}
l13:
	m.fn153(v4)
	panic("unreachable")
l23:
	m.fn556(v0, v1, v6, v3, v4)
	panic("unreachable")
l18:
	t24 := v5
	v7 = int64(uint32(i32(5))) << 32
	store64(m.memory[int64(uint32(t24))+32:], uint64(v7|int64(uint32(v5+i32(8)))))
	store64(m.memory[int64(uint32(v5))+24:], uint64(v7|int64(uint32(v5+i32(4)))))
	m.fn91(i32(1050412), v5+i32(24), v4)
	panic("unreachable")
}
func (m *Module) fn1657(v0, v1 int32) int32 {
	var v2, v3, v4 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	{
		{
			{
				t1 := int32(load32(m.memory[int64(uint32(v1))+8:]))
				v3 = t1
				if v3&i32(0x2000000) != 0 {
					t3 := int32(load32(m.memory[uint32(v0):]))
					v4 = t3
					v3 = i32(9)
				l4:
					{
						t4 := int32(m.memory[int64(uint32(v4&i32(15)))+1107936])
						m.memory[uint32(v2+i32(8)+v3+i32(-2))] = byte(t4)
						v3 = v3 + i32(-1)
						v4 = int32(uint32(v4) >> 4)
						if v4 != 0 {
							goto l4
						}
					}
					v4 = i32(1)
					t5 := m.fn1638(v1, i32(1), i32(1131670), i32(2), v2+i32(8)+v3+i32(-1), i32(9)-v3)
					if t5 == 0 {
						goto l2
					}
					goto l3
				}
				if v3&i32(0x4000000) != 0 {
					goto l1
				}
				t2 := m.fn72(v0, v1)
				if t2 == 0 {
					goto l2
				}
				v4 = i32(1)
				goto l3
			}
		l1:
			t6 := int32(load32(m.memory[uint32(v0):]))
			v4 = t6
			v3 = i32(9)
		l5:
			{
				t7 := int32(m.memory[int64(uint32(v4&i32(15)))+1131672])
				m.memory[uint32(v2+i32(8)+v3+i32(-2))] = byte(t7)
				v3 = v3 + i32(-1)
				v4 = int32(uint32(v4) >> 4)
				if v4 != 0 {
					goto l5
				}
			}
			v4 = i32(1)
			t8 := m.fn1638(v1, i32(1), i32(1131670), i32(2), v2+i32(8)+v3+i32(-1), i32(9)-v3)
			if t8 != 0 {
				goto l3
			}
		}
	l2:
		{
			t9 := int32(load32(m.memory[uint32(v1):]))
			t10 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t11 := int32(load32(m.memory[int64(uint32(t10))+12:]))
			t12 := m.t0[uint(t11)].(func(int32, int32, int32) int32)(t9, i32(1284184), i32(2))
			if t12 == 0 {
				goto l6
			}
			v4 = i32(1)
			goto l3
		}
	l6:
		v3 = v0 + i32(4)
		{
			t13 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			v4 = t13
			if v4&i32(0x2000000) != 0 {
				t15 := int32(load32(m.memory[uint32(v3):]))
				v4 = t15
				v3 = i32(9)
			l9:
				{
					t16 := int32(m.memory[int64(uint32(v4&i32(15)))+1107936])
					m.memory[uint32(v2+i32(8)+v3+i32(-2))] = byte(t16)
					v3 = v3 + i32(-1)
					v4 = int32(uint32(v4) >> 4)
					if v4 != 0 {
						goto l9
					}
				}
				t17 := m.fn1638(v1, i32(1), i32(1131670), i32(2), v2+i32(8)+v3+i32(-1), i32(9)-v3)
				v4 = t17
				goto l3
			}
			if v4&i32(0x4000000) != 0 {
				goto l8
			}
			t14 := m.fn72(v3, v1)
			v4 = t14
			goto l3
		}
	l8:
		t18 := int32(load32(m.memory[uint32(v3):]))
		v4 = t18
		v3 = i32(9)
	l10:
		{
			t19 := int32(m.memory[int64(uint32(v4&i32(15)))+1131672])
			m.memory[uint32(v2+i32(8)+v3+i32(-2))] = byte(t19)
			v3 = v3 + i32(-1)
			v4 = int32(uint32(v4) >> 4)
			if v4 != 0 {
				goto l10
			}
		}
		t20 := m.fn1638(v1, i32(1), i32(1131670), i32(2), v2+i32(8)+v3+i32(-1), i32(9)-v3)
		v4 = t20
	}
l3:
	m.g0 = v2 + i32(16)
	return v4
}
func (m *Module) fn1658(v0, v1 int32) int32 {
	var v2, v3, v4, v5 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	v3 = i32(1)
	{
		t1 := int32(load32(m.memory[uint32(v1):]))
		v4 = t1
		t2 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t3 := v4
		v5 = t2
		t4 := int32(load32(m.memory[int64(uint32(v5))+16:]))
		v1 = t4
		t5 := m.t0[uint(v1)].(func(int32, int32) int32)(t3, i32(39))
		if t5 != 0 {
			goto l0
		}
		t6 := int32(load32(m.memory[uint32(v0):]))
		m.fn1645(v2, t6, i32(257))
		{
			t7 := int32(m.memory[int64(uint32(v2))+13])
			v3 = t7
			if uint32(v3) < uint32(i32(129)) {
				t10 := int32(m.memory[int64(uint32(v2))+12])
				t11 := v4
				t12 := v2
				v0 = t10
				t13 := int32(load32(m.memory[int64(uint32(v5))+12:]))
				t14 := m.t0[uint(t13)].(func(int32, int32, int32) int32)(t11, t12+v0, v3-v0)
				if t14 == 0 {
					goto l2
				}
				v3 = i32(1)
				goto l0
			}
			t8 := int32(load32(m.memory[uint32(v2):]))
			t9 := m.t0[uint(v1)].(func(int32, int32) int32)(v4, t8)
			if t9 == 0 {
				goto l2
			}
			v3 = i32(1)
			goto l0
		}
	l2:
		t15 := m.t0[uint(v1)].(func(int32, int32) int32)(v4, i32(39))
		v3 = t15
	}
l0:
	m.g0 = v2 + i32(16)
	return v3
}
func (m *Module) fn1659(v0, v1 int32) int32 {
	t0 := m.fn110(v1, i32(1131812), i32(24))
	return t0
}
func (m *Module) fn1660(v0, v1 int32) int32 {
	t0 := m.fn110(v1, i32(1131780), i32(32))
	return t0
}
func (m *Module) fn1661(v0, v1, v2, v3, v4, v5 int32) {
	var v6 int32
	{
		if v2 == 0 {
			m.fn256(i32(1111329), i32(33), i32(1111364))
			panic("unreachable")
		}
		t0 := int32(m.memory[uint32(v1)])
		if uint32(t0) <= uint32(i32(48)) {
			m.fn256(i32(1111380), i32(31), i32(1111412))
			panic("unreachable")
		}
		store16(m.memory[uint32(v5):], uint16(i32(2)))
		{
			{
				v6 = int32(int16(v3))
				if v6 < i32(1) {
					store32(m.memory[int64(uint32(v5))+32:], uint32(v2))
					store32(m.memory[int64(uint32(v5))+28:], uint32(v1))
					store16(m.memory[int64(uint32(v5))+24:], uint16(i32(2)))
					store16(m.memory[int64(uint32(v5))+12:], uint16(i32(0)))
					store32(m.memory[int64(uint32(v5))+8:], uint32(i32(2)))
					store32(m.memory[int64(uint32(v5))+4:], uint32(i32(1108009)))
					t2 := v5
					v3 = i32(0) - v6
					store32(m.memory[int64(uint32(t2))+16:], uint32(v3))
					v1 = i32(3)
					if uint32(v4) <= uint32(v2) {
						goto l5
					}
					v2 = v4 - v2
					if uint32(v2) <= uint32(v3) {
						goto l5
					}
					v4 = v2 + v6
					goto l6
				}
				store32(m.memory[int64(uint32(v5))+4:], uint32(v1))
				t1 := v2
				v3 = v3 & i32(0xffff)
				if uint32(t1) > uint32(v3) {
					goto l3
				}
				store16(m.memory[int64(uint32(v5))+12:], uint16(i32(0)))
				store32(m.memory[int64(uint32(v5))+8:], uint32(v2))
				store32(m.memory[int64(uint32(v5))+16:], uint32(v3-v2))
				if v4 != 0 {
					store32(m.memory[int64(uint32(v5))+32:], uint32(i32(1)))
					store32(m.memory[int64(uint32(v5))+28:], uint32(i32(1109519)))
					store16(m.memory[int64(uint32(v5))+24:], uint16(i32(2)))
					goto l6
				}
				v1 = i32(2)
				goto l5
			}
		l3:
			store16(m.memory[int64(uint32(v5))+24:], uint16(i32(2)))
			store32(m.memory[int64(uint32(v5))+20:], uint32(i32(1)))
			store32(m.memory[int64(uint32(v5))+16:], uint32(i32(1109519)))
			store16(m.memory[int64(uint32(v5))+12:], uint16(i32(2)))
			store32(m.memory[int64(uint32(v5))+8:], uint32(v3))
			t3 := v5
			v2 = v2 - v3
			store32(m.memory[int64(uint32(t3))+32:], uint32(v2))
			store32(m.memory[int64(uint32(v5))+28:], uint32(v1+v3))
			if uint32(v4) > uint32(v2) {
				goto l7
			}
			v1 = i32(3)
			goto l5
		l7:
			v4 = v4 - v2
		}
	l6:
		store32(m.memory[int64(uint32(v5))+40:], uint32(v4))
		store16(m.memory[int64(uint32(v5))+36:], uint16(i32(0)))
		v1 = i32(4)
	l5:
		store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
		store32(m.memory[uint32(v0):], uint32(v5))
		return
	}
}
