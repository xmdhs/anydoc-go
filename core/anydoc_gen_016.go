package core

import (
	"math"
	"math/bits"
)

func (m *Module) fn672(v0, v1, v2 int32) int32 {
	var v3, v4 int32
	v3 = i32(1)
	{
		t0 := int32(load32(m.memory[int64(uint32(v2))+12:]))
		t1 := v1
		v4 = t0
		t2 := m.t0[uint(v4)].(func(int32, int32, int32) int32)(t1, i32(1273626), i32(21))
		if t2 != 0 {
			goto l0
		}
		{
			{
				t3 := int32(load32(m.memory[uint32(v0):]))
				if t3 != i32(-1) {
					goto l1
				}
				t4 := m.t0[uint(v4)].(func(int32, int32, int32) int32)(v1, i32(1272313), i32(9))
				if t4 != 0 {
					goto l0
				}
				t5 := int32(load32(m.memory[int64(uint32(v0))+4:]))
				t6 := int32(load32(m.memory[int64(uint32(v0))+8:]))
				t7 := m.fn685(v1, v2, t5, t6)
				if t7 == 0 {
					goto l2
				}
				goto l0
			}
		l1:
			t8 := m.t0[uint(v4)].(func(int32, int32, int32) int32)(v1, i32(1272322), i32(6))
			if t8 != 0 {
				goto l0
			}
			t9 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t10 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			t11 := m.fn685(v1, v2, t9, t10)
			if t11 != 0 {
				goto l0
			}
		}
	l2:
		v3 = i32(1)
		t12 := m.t0[uint(v4)].(func(int32, int32, int32) int32)(v1, i32(1272328), i32(1))
		if t12 != 0 {
			goto l0
		}
		t13 := m.t0[uint(v4)].(func(int32, int32, int32) int32)(v1, i32(1273624), i32(2))
		v3 = t13
	}
l0:
	return v3
}
func (m *Module) fn673(v0, v1, v2 int32) int32 {
	var v3, v4 int32
	v3 = i32(1)
	{
		t0 := int32(load32(m.memory[int64(uint32(v2))+12:]))
		t1 := v1
		v4 = t0
		t2 := m.t0[uint(v4)].(func(int32, int32, int32) int32)(t1, i32(1273649), i32(22))
		if t2 != 0 {
			goto l0
		}
		{
			{
				t3 := int32(load32(m.memory[uint32(v0):]))
				if t3 != i32(-1) {
					goto l1
				}
				t4 := m.t0[uint(v4)].(func(int32, int32, int32) int32)(v1, i32(1272313), i32(9))
				if t4 != 0 {
					goto l0
				}
				t5 := int32(load32(m.memory[int64(uint32(v0))+4:]))
				t6 := int32(load32(m.memory[int64(uint32(v0))+8:]))
				t7 := m.fn685(v1, v2, t5, t6)
				if t7 == 0 {
					goto l2
				}
				goto l0
			}
		l1:
			t8 := m.t0[uint(v4)].(func(int32, int32, int32) int32)(v1, i32(1272322), i32(6))
			if t8 != 0 {
				goto l0
			}
			t9 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t10 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			t11 := m.fn685(v1, v2, t9, t10)
			if t11 != 0 {
				goto l0
			}
		}
	l2:
		v3 = i32(1)
		t12 := m.t0[uint(v4)].(func(int32, int32, int32) int32)(v1, i32(1272328), i32(1))
		if t12 != 0 {
			goto l0
		}
		t13 := m.t0[uint(v4)].(func(int32, int32, int32) int32)(v1, i32(1273624), i32(2))
		v3 = t13
	}
l0:
	return v3
}
func (m *Module) fn674(v0, v1 int32) int32 {
	var v2, v3, v4 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	store32(m.memory[int64(uint32(v2))+4:], uint32(v0))
	t1 := int32(load32(m.memory[uint32(v1):]))
	t2 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	t3 := int32(load32(m.memory[int64(uint32(t2))+12:]))
	t4 := m.t0[uint(t3)].(func(int32, int32, int32) int32)(t1, i32(1079860), i32(9))
	v0 = t4
	m.memory[int64(uint32(v2))+13] = byte(i32(0))
	m.memory[int64(uint32(v2))+12] = byte(v0)
	store32(m.memory[int64(uint32(v2))+8:], uint32(v1))
	t5 := m.fn350(v2+i32(8), i32(1079869), i32(7), v2+i32(4), i32(85))
	v3 = t5
	t6 := int32(m.memory[int64(uint32(v2))+13])
	v0 = t6
	t7 := int32(m.memory[int64(uint32(v2))+12])
	t8 := v0
	v4 = t7
	v1 = t8 | v4
	{
		if v0 != i32(1) {
			goto l0
		}
		if v4&i32(1) != 0 {
			goto l0
		}
		{
			t9 := int32(load32(m.memory[uint32(v3):]))
			v1 = t9
			t10 := int32(m.memory[int64(uint32(v1))+10])
			if t10&i32(128) != 0 {
				goto l1
			}
			t11 := int32(load32(m.memory[uint32(v1):]))
			t12 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t13 := int32(load32(m.memory[int64(uint32(t12))+12:]))
			t14 := m.t0[uint(t13)].(func(int32, int32, int32) int32)(t11, i32(1273624), i32(2))
			v1 = t14
			goto l0
		}
	l1:
		t15 := int32(load32(m.memory[uint32(v1):]))
		t16 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t17 := int32(load32(m.memory[int64(uint32(t16))+12:]))
		t18 := m.t0[uint(t17)].(func(int32, int32, int32) int32)(t15, i32(1099063), i32(1))
		v1 = t18
	}
l0:
	m.g0 = v2 + i32(16)
	return v1 & i32(1)
}
func (m *Module) fn675(v0, v1, v2 int32) int32 {
	var v3, v4 int32
	v3 = i32(1)
	{
		t0 := int32(load32(m.memory[int64(uint32(v2))+12:]))
		t1 := v1
		v4 = t0
		t2 := m.t0[uint(v4)].(func(int32, int32, int32) int32)(t1, i32(1273671), i32(19))
		if t2 != 0 {
			goto l0
		}
		{
			{
				t3 := int32(load32(m.memory[uint32(v0):]))
				if t3 != i32(-1) {
					goto l1
				}
				t4 := m.t0[uint(v4)].(func(int32, int32, int32) int32)(v1, i32(1272313), i32(9))
				if t4 != 0 {
					goto l0
				}
				t5 := int32(load32(m.memory[int64(uint32(v0))+4:]))
				t6 := int32(load32(m.memory[int64(uint32(v0))+8:]))
				t7 := m.fn685(v1, v2, t5, t6)
				if t7 == 0 {
					goto l2
				}
				goto l0
			}
		l1:
			t8 := m.t0[uint(v4)].(func(int32, int32, int32) int32)(v1, i32(1272322), i32(6))
			if t8 != 0 {
				goto l0
			}
			t9 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t10 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			t11 := m.fn685(v1, v2, t9, t10)
			if t11 != 0 {
				goto l0
			}
		}
	l2:
		v3 = i32(1)
		t12 := m.t0[uint(v4)].(func(int32, int32, int32) int32)(v1, i32(1272328), i32(1))
		if t12 != 0 {
			goto l0
		}
		t13 := m.t0[uint(v4)].(func(int32, int32, int32) int32)(v1, i32(1273624), i32(2))
		v3 = t13
	}
l0:
	return v3
}
func (m *Module) fn676(v0, v1, v2 int32) int32 {
	var v3, v4 int32
	v3 = i32(1)
	{
		t0 := int32(load32(m.memory[int64(uint32(v2))+12:]))
		t1 := v1
		v4 = t0
		t2 := m.t0[uint(v4)].(func(int32, int32, int32) int32)(t1, i32(1273690), i32(20))
		if t2 != 0 {
			goto l0
		}
		{
			{
				t3 := int32(load32(m.memory[uint32(v0):]))
				if t3 != i32(-1) {
					goto l1
				}
				t4 := m.t0[uint(v4)].(func(int32, int32, int32) int32)(v1, i32(1272313), i32(9))
				if t4 != 0 {
					goto l0
				}
				t5 := int32(load32(m.memory[int64(uint32(v0))+4:]))
				t6 := int32(load32(m.memory[int64(uint32(v0))+8:]))
				t7 := m.fn685(v1, v2, t5, t6)
				if t7 == 0 {
					goto l2
				}
				goto l0
			}
		l1:
			t8 := m.t0[uint(v4)].(func(int32, int32, int32) int32)(v1, i32(1272322), i32(6))
			if t8 != 0 {
				goto l0
			}
			t9 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			t10 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			t11 := m.fn685(v1, v2, t9, t10)
			if t11 != 0 {
				goto l0
			}
		}
	l2:
		v3 = i32(1)
		t12 := m.t0[uint(v4)].(func(int32, int32, int32) int32)(v1, i32(1272328), i32(1))
		if t12 != 0 {
			goto l0
		}
		t13 := m.t0[uint(v4)].(func(int32, int32, int32) int32)(v1, i32(1273624), i32(2))
		v3 = t13
	}
l0:
	return v3
}
func (m *Module) fn677(v0, v1, v2 int32) {
	var v3, v4, v5, v6, v7, v8 int32
	var v9 int64
	var v10, v11, v12 int32
	t0 := m.g0
	v3 = t0 - i32(16)
	m.g0 = v3
	{
		t1 := int32(uint32(v2-v1) / uint32(i32(24)))
		v4 = t1
		t2 := int32(load32(m.memory[uint32(v0):]))
		t3 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		t4 := v4
		v5 = t3
		if uint32(t4) <= uint32(t2-v5) {
			goto l0
		}
		m.fn203(v0, v5, v4, i32(8), i32(24))
		t5 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		v5 = t5
	}
l0:
	{
		if v1 == v2 {
			goto l1
		}
		t6 := int32(load32(m.memory[int64(uint32(v0))+4:]))
		v6 = t6 + v5*i32(24)
		v7 = i32(0)
	l18:
		{
			{
				v2 = v1 + v7
				t7 := int32(m.memory[uint32(v2)])
				v8 = t7
				switch v8 {
				case 8:
					goto l10
				default:
					t8 := int32(m.memory[uint32(v2+i32(3))])
					m.memory[int64(uint32(v3))+14] = byte(t8)
					t9 := int32(load16(m.memory[uint32(v2+i32(1)):]))
					store16(m.memory[int64(uint32(v3))+12:], uint16(t9))
					t10 := int64(load64(m.memory[uint32(v2+i32(16)):]))
					v9 = t10
					t11 := int32(load32(m.memory[uint32(v2+i32(12)):]))
					v10 = t11
					t12 := int32(load32(m.memory[uint32(v2+i32(8)):]))
					v11 = t12
					t13 := int32(load32(m.memory[uint32(v2+i32(4)):]))
					v12 = t13
					goto l10
				case 1:
					t14 := int32(m.memory[uint32(v2+i32(3))])
					m.memory[int64(uint32(v3))+14] = byte(t14)
					t15 := int32(load16(m.memory[uint32(v2+i32(1)):]))
					store16(m.memory[int64(uint32(v3))+12:], uint16(t15))
					t16 := int64(load64(m.memory[uint32(v2+i32(16)):]))
					v9 = t16
					t17 := int32(load32(m.memory[uint32(v2+i32(12)):]))
					v10 = t17
					t18 := int32(load32(m.memory[uint32(v2+i32(8)):]))
					v11 = t18
					t19 := int32(load32(m.memory[uint32(v2+i32(4)):]))
					v12 = t19
					goto l10
				case 2:
					t20 := int32(load32(m.memory[uint32(v2+i32(12)):]))
					v10 = t20
					if v10 == 0 {
						goto l11
					}
					t21 := int32(load32(m.memory[uint32(v2+i32(8)):]))
					v2 = t21
					t22 := m.fn11(v10)
					v11 = t22
					if v11 == 0 {
						m.fn7(i32(1), v10)
						panic("unreachable")
					}
					if v10 != 0 {
						memory_copy(m.memory, uint32(v11), uint32(v2), uint32(v10))
						v12 = v10
						goto l10
					}
					v12 = v10
					goto l10
				case 3:
					t23 := int32(m.memory[uint32(v2+i32(3))])
					m.memory[int64(uint32(v3))+14] = byte(t23)
					t24 := int32(load16(m.memory[uint32(v2+i32(1)):]))
					store16(m.memory[int64(uint32(v3))+12:], uint16(t24))
					t25 := int64(load64(m.memory[uint32(v2+i32(16)):]))
					v9 = t25
					t26 := int32(load32(m.memory[uint32(v2+i32(12)):]))
					v10 = t26
					t27 := int32(load32(m.memory[uint32(v2+i32(8)):]))
					v11 = t27
					t28 := int32(load32(m.memory[uint32(v2+i32(4)):]))
					v12 = t28
					goto l10
				case 4:
					t29 := int32(m.memory[uint32(v2+i32(3))])
					m.memory[int64(uint32(v3))+14] = byte(t29)
					t30 := int32(load16(m.memory[uint32(v2+i32(1)):]))
					store16(m.memory[int64(uint32(v3))+12:], uint16(t30))
					t31 := int64(load64(m.memory[uint32(v2+i32(16)):]))
					v9 = t31
					t32 := int32(load32(m.memory[uint32(v2+i32(12)):]))
					v10 = t32
					t33 := int32(load32(m.memory[uint32(v2+i32(8)):]))
					v11 = t33
					t34 := int32(load32(m.memory[uint32(v2+i32(4)):]))
					v12 = t34
					goto l10
				case 5:
					t35 := int32(load32(m.memory[uint32(v2+i32(12)):]))
					v10 = t35
					if v10 == 0 {
						goto l11
					}
					t36 := int32(load32(m.memory[uint32(v2+i32(8)):]))
					v2 = t36
					t37 := m.fn11(v10)
					v11 = t37
					if v11 == 0 {
						m.fn7(i32(1), v10)
						panic("unreachable")
					}
					if v10 != 0 {
						memory_copy(m.memory, uint32(v11), uint32(v2), uint32(v10))
						v12 = v10
						goto l10
					}
					v12 = v10
					goto l10
				case 6:
					t38 := int32(load32(m.memory[uint32(v2+i32(12)):]))
					v10 = t38
					if v10 == 0 {
						goto l11
					}
					t39 := int32(load32(m.memory[uint32(v2+i32(8)):]))
					v2 = t39
					t40 := m.fn11(v10)
					v11 = t40
					if v11 == 0 {
						m.fn7(i32(1), v10)
						panic("unreachable")
					}
					if v10 != 0 {
						memory_copy(m.memory, uint32(v11), uint32(v2), uint32(v10))
						v12 = v10
						goto l10
					}
					v12 = v10
					goto l10
				case 7:
					t41 := int32(m.memory[uint32(v2+i32(3))])
					m.memory[int64(uint32(v3))+14] = byte(t41)
					t42 := int32(load16(m.memory[uint32(v2+i32(1)):]))
					store16(m.memory[int64(uint32(v3))+12:], uint16(t42))
					t43 := int64(load64(m.memory[uint32(v2+i32(16)):]))
					v9 = t43
					t44 := int32(load32(m.memory[uint32(v2+i32(12)):]))
					v10 = t44
					t45 := int32(load32(m.memory[uint32(v2+i32(8)):]))
					v11 = t45
					t46 := int32(load32(m.memory[uint32(v2+i32(4)):]))
					v12 = t46
					goto l10
				}
			}
		l11:
			v11 = i32(1)
			v10 = i32(0)
			v12 = i32(0)
		l10:
			v2 = v6 + v7
			m.memory[uint32(v2)] = byte(v8)
			t47 := int32(load16(m.memory[int64(uint32(v3))+12:]))
			store16(m.memory[uint32(v2+i32(1)):], uint16(t47))
			t48 := int32(m.memory[int64(uint32(v3))+14])
			m.memory[uint32(v2+i32(3))] = byte(t48)
			store64(m.memory[uint32(v2+i32(16)):], uint64(v9))
			store32(m.memory[uint32(v2+i32(12)):], uint32(v10))
			store32(m.memory[uint32(v2+i32(8)):], uint32(v11))
			store32(m.memory[uint32(v2+i32(4)):], uint32(v12))
			v7 = v7 + i32(24)
			v5 = v5 + i32(1)
			v4 = v4 + i32(-1)
			if v4 != 0 {
				goto l18
			}
		}
	}
l1:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v5))
	m.g0 = v3 + i32(16)
}
func (m *Module) fn678(v0, v1, v2 int32) {
	var v3 int32
	var v4 int64
	var v5, v6, v7, v8, v9, v10, v11 int32
	var v12 int64
	var v13, v14 int32
	var v15 int64
	var v16 int32
	var v17 float64
	var v18, v19, v20, v21, v22, v23 int64
	t0 := m.g0
	v3 = t0 - i32(1648)
	m.g0 = v3
	v4 = i64(0)
	if v2 == 0 {
		goto l0
	}
	{
		t1 := int32(m.memory[uint32(v1)])
		v5 = t1
		if v5 != i32(45) {
			goto l1
		}
		if v2 == i32(1) {
			goto l0
		}
		v6 = v1 + i32(1)
		goto l2
	}
l1:
	v6 = v1
	if v5 != i32(43) {
		goto l2
	}
	if v2 == i32(1) {
		goto l0
	}
	v6 = v1 + i32(1)
l2:
	v7 = i32(0)
	{
		t2 := v6
		v8 = v1 + v2
		var p3 int32
		if t2 == v8 {
			p3 = 1
		}
		v9 = p3
		if v9 == 0 {
			goto l3
		}
		v4 = i64(0)
		v10 = v8
		v11 = v8
		v12 = i64(0)
		goto l4
	}
l3:
	v13 = v1 + v2
	v4 = i64(0)
	v10 = v6
l6:
	{
		t4 := int32(m.memory[uint32(v10)])
		v11 = t4 + i32(-48)
		if uint32(v11&i32(255)) > uint32(i32(9)) {
			goto l5
		}
		v4 = v4*i64(10) + int64(uint32(v11))&i64(255)
		v10 = v10 + i32(1)
		if v10 != v8 {
			goto l6
		}
	}
	v10 = v13
l5:
	v12 = i64(0)
	if v10 != v8 {
		goto l7
	}
	v10 = v8
	v11 = v8
	goto l4
l7:
	{
		t5 := int32(m.memory[uint32(v10)])
		if t5 == i32(46) {
			goto l8
		}
		v11 = v10
		goto l4
	}
l8:
	{
		{
			t6 := v8
			v14 = v10 + i32(1)
			v7 = t6 - v14
			if v7 >= i32(8) {
				goto l9
			}
			v11 = v14
			goto l10
		}
	l9:
		{
			t7 := int64(load64(m.memory[uint32(v14):]))
			v15 = t7
			t8 := v15 + i64(5063812098665367110)
			v15 = v15 + i64(-3472328296227680304)
			if (t8|v15)&i64(-0x7f7f7f7f7f7f7f80) == i64(0) {
				goto l11
			}
			v11 = v14
			goto l10
		}
	l11:
		v15 = v15*i64(10) + int64(uint64(v15)>>8)
		v4 = int64(uint64(int64(uint64(v15)>>16)&i64(0xff000000ff)*i64(0x271000000001)+v15&i64(0xff000000ff)*i64(0xf424000000064))>>32) + v4*i64(100000000)
		t9 := v8
		v11 = v10 + i32(9)
		if t9-v11 < i32(8) {
			goto l10
		}
		t10 := int64(load64(m.memory[uint32(v11):]))
		v15 = t10
		t11 := v15 + i64(5063812098665367110)
		v15 = v15 + i64(-3472328296227680304)
		if (t11|v15)&i64(-0x7f7f7f7f7f7f7f80) != i64(0) {
			goto l10
		}
		v15 = v15*i64(10) + int64(uint64(v15)>>8)
		v4 = int64(uint64(int64(uint64(v15)>>16)&i64(0xff000000ff)*i64(0x271000000001)+v15&i64(0xff000000ff)*i64(0xf424000000064))>>32) + v4*i64(100000000)
		v11 = v10 + i32(17)
	}
l10:
	if v11 != v8 {
		goto l12
	}
	v11 = v8
	goto l13
l12:
	v13 = v11 + (v13 - v11)
l15:
	{
		t12 := int32(m.memory[uint32(v11)])
		v7 = t12 + i32(-48)
		if uint32(v7&i32(255)) > uint32(i32(9)) {
			goto l14
		}
		v4 = v4*i64(10) + int64(uint32(v7))&i64(255)
		v11 = v11 + i32(1)
		if v11 != v8 {
			goto l15
		}
	}
	v11 = v13
l14:
	v7 = v11 - v14
l13:
	v12 = int64(i32(0) - v7)
l4:
	{
		{
			{
				v14 = v7 + (v10 - v6)
				if v14 == 0 {
					v16 = i32(3)
					v4 = i64(0)
					if uint32(v2) < uint32(i32(3)) {
						goto l0
					}
					v17 = math.Float64frombits(0x7ff8000000000000)
					t29 := int32(m.memory[int64(uint32(v1))+1])
					v8 = t29
					t30 := int32(m.memory[int64(uint32(v1))+2])
					t31 := v8 ^ i32(65) | (v5 ^ i32(78))
					v10 = t30
					v11 = v10 ^ i32(78)
					if (t31|v11)&i32(223) == 0 {
						goto l40
					}
					t32 := v5 ^ i32(73) | (v10 ^ i32(70))
					v7 = v8 ^ i32(110)
					if (t32|v7)&i32(223) == 0 {
						goto l41
					}
					if v2 == i32(3) {
						goto l0
					}
					v6 = v1 + i32(1)
					switch v5 + i32(-43) {
					default:
						goto l0
					case 0:
						t33 := int32(m.memory[int64(uint32(v1))+3])
						t34 := v10 ^ i32(65)
						v10 = t33
						if (t34|(v10^i32(78))|v7)&i32(223) != 0 {
							if (v8^i32(73)|(v10^i32(70))|v11)&i32(223) != 0 {
								goto l0
							}
							t35 := m.fn681(v6, v2+i32(-1))
							v16 = t35 + i32(1)
							v17 = math.Float64frombits(0x7ff0000000000000)
							goto l40
						}
						v16 = i32(4)
						goto l40
					case 2:
						t36 := int32(m.memory[int64(uint32(v1))+3])
						t37 := v10 ^ i32(65)
						v10 = t36
						if (t37|(v10^i32(78))|v7)&i32(223) != 0 {
							if (v8^i32(73)|(v10^i32(70))|v11)&i32(223) != 0 {
								goto l0
							}
							t38 := m.fn681(v6, v2+i32(-1))
							v16 = t38 + i32(1)
							v17 = math.Float64frombits(0xfff0000000000000)
							goto l40
						}
						v16 = i32(4)
						v17 = math.Float64frombits(0xfff8000000000000)
						goto l40
					}
				}
				v15 = i64(0)
				{
					if v11 != v8 {
						goto l17
					}
					v11 = v8
					goto l18
				l17:
					t13 := int32(m.memory[uint32(v11)])
					if t13|i32(32) != i32(101) {
						goto l18
					}
					v16 = i32(0)
					v7 = v11 + i32(1)
					if v7 == v8 {
						goto l19
					}
					{
						t14 := int32(m.memory[uint32(v7)])
						v13 = t14
						switch v13 + i32(-43) {
						default:
							goto l19
						case 0, 2:
							v7 = v11 + i32(2)
							var p15 int32
							if v13 == i32(45) {
								p15 = 1
							}
							v16 = p15
						}
					}
				l19:
					if v7 == v8 {
						goto l18
					}
					t16 := int32(m.memory[uint32(v7)])
					if uint32((t16+i32(-48))&i32(255)) > uint32(i32(9)) {
						goto l18
					}
					v11 = v1 + v2
					v15 = i64(0)
				l23:
					{
						{
							t17 := int32(m.memory[uint32(v7)])
							v13 = t17 + i32(-48)
							if uint32(v13&i32(255)) <= uint32(i32(9)) {
								goto l21
							}
							v11 = v7
							goto l22
						}
					l21:
						p18 := v15
						if v15 < i64(65536) {
							p18 = v15*i64(10) + int64(uint32(v13))&i64(255)
						}
						v15 = p18
						v7 = v7 + i32(1)
						if v7 != v8 {
							goto l23
						}
					}
				l22:
					p19 := v15
					if v16 != 0 {
						p19 = i64(0) - v15
					}
					v15 = p19
				}
			l18:
				v16 = v11 - v1
				if v14 < i32(20) {
					goto l24
				}
				if v9 != 0 {
					goto l25
				}
				v13 = v14 + i32(-19)
				v11 = v6
			l28:
				{
					t20 := int32(m.memory[uint32(v11)])
					v7 = t20
					switch v7 + i32(-46) {
					default:
						goto l27
					case 0, 2:
						t21 := v13
						v14 = v7 + i32(-47)
						p22 := v14
						if uint32(v14) > uint32(v7) {
							p22 = i32(0)
						}
						v13 = t21 - p22
						v11 = v11 + i32(1)
						if v11 != v8 {
							goto l28
						}
					}
				}
			l27:
				if v13 < i32(1) {
					goto l24
				}
			l25:
				v7 = v1 + v2
				v4 = i64(0)
			l33:
				{
					if v6 == v8 {
						goto l29
					}
					t23 := int32(m.memory[uint32(v6)])
					v11 = t23 + i32(-48)
					if uint32(v11&i32(255)) <= uint32(i32(9)) {
						goto l30
					}
					v7 = v6
				}
			l29:
				v7 = v7 + i32(1)
				if v7 != v8 {
					v10 = v7
				l35:
					{
						t24 := int32(m.memory[uint32(v10)])
						v11 = t24 + i32(-48)
						if uint32(v11&i32(255)) >= uint32(i32(10)) {
							goto l34
						}
						v10 = v10 + i32(1)
						v4 = v4*i64(10) + int64(uint32(v11))&i64(255)
						if uint64(v4) > uint64(i64(999999999999999999)) {
							goto l34
						}
						if v10 != v8 {
							goto l35
						}
						goto l34
					}
				}
				v10 = v7 - v8
				goto l32
			l30:
				v6 = v6 + i32(1)
				v4 = v4*i64(10) + int64(uint32(v11))&i64(255)
				if uint64(v4) < uint64(i64(1000000000000000000)) {
					goto l33
				}
				v10 = v10 - v6
				goto l32
			l24:
				v10 = i32(0)
				v15 = v15 + v12
				if uint64(v15+i64(-38)) < uint64(i64(-60)) {
					goto l36
				}
				if uint64(v4) > uint64(i64(0x20000000000000)) {
					goto l36
				}
				{
					if v15 > i64(22) {
						t26 := int64(load64(m.memory[uint32(int32(v15)<<3+i32(1098528)):]))
						m.fn982(v3+i32(64), v4, i64(0), t26, i64(0))
						t27 := int64(load64(m.memory[int64(uint32(v3))+72:]))
						if t27 != i64(0) {
							goto l36
						}
						t28 := int64(load64(m.memory[int64(uint32(v3))+64:]))
						v12 = t28
						if uint64(v12) > uint64(i64(0x20000000000000)) {
							goto l36
						}
						v17 = float64(float64(uint64(v12)) * float64(1e+22))
						goto l39
					}
					v10 = int32(v15)
					v17 = float64(uint64(v4))
					if v15 < i64(0) {
						goto l38
					}
					t25 := math.Float64frombits(load64(m.memory[int64(uint32(v10<<3))+1122056:]))
					v17 = float64(t25 * v17)
					goto l39
				}
			l38:
				t39 := math.Float64frombits(load64(m.memory[uint32(i32(1122056)-v10<<3):]))
				v17 = float64(v17 / t39)
			}
		l39:
			t41 := v0
			p40 := v17
			if v5 == i32(45) {
				p40 = -v17
			}
			store64(m.memory[int64(uint32(t41))+8:], math.Float64bits(p40))
			goto l46
		}
	l41:
		v17 = math.Float64frombits(0x7ff0000000000000)
		t42 := m.fn681(v1, v2)
		v16 = t42
	}
l40:
	store64(m.memory[int64(uint32(v0))+8:], math.Float64bits(v17))
	goto l46
l34:
	v10 = v7 - v10
l32:
	v15 = v15 + int64(v10)
	v10 = i32(1)
l36:
	v8 = i32(0)
	v12 = i64(0)
	{
		{
			{
				if v4 == 0 {
					goto l47
				}
				if v15 < i64(-342) {
					goto l47
				}
				v8 = i32(2047)
				if v15 > i64(308) {
					goto l47
				}
				t43 := v3 + i32(48)
				v11 = int32(v15)
				v7 = v11 << 4
				t44 := int64(load64(m.memory[uint32(v7+i32(1114680)):]))
				t45 := v4
				v18 = int64(bits.LeadingZeros64(uint64(v4)))
				v19 = i64_shl(t45, v18)
				m.fn982(t43, t44, i64(0), v19, i64(0))
				t46 := int64(load64(m.memory[int64(uint32(v3))+48:]))
				v12 = t46
				{
					t47 := int64(load64(m.memory[int64(uint32(v3))+56:]))
					v20 = t47
					if v20&i64(511) != i64(511) {
						goto l48
					}
					t48 := int64(load64(m.memory[uint32(v7+i32(1109208)+i32(5480)):]))
					m.fn982(v3+i32(32), t48, i64(0), v19, i64(0))
					t49 := int64(load64(m.memory[int64(uint32(v3))+40:]))
					v19 = t49
					v12 = v19 + v12
					var p50 int32
					if uint64(v12) < uint64(v19) {
						p50 = 1
					}
					v20 = int64(uint32(p50)) + v20
				}
			l48:
				if v12 != i64(-1) {
					goto l49
				}
				if uint64(v15+i64(27)) <= uint64(i64(82)) {
					goto l49
				}
				if v10 == 0 {
					goto l50
				}
				v12 = i64(0)
				v8 = i32(-1)
				goto l51
			l49:
				t51 := v20
				v21 = int64(uint64(v20) >> 63)
				v22 = v21 + i64(9)
				v19 = i64_shr_u(t51, v22)
				{
					v11 = v11*i32(217706)>>16 - int32(v18) + int32(v21) + i32(63)
					if v11 < i32(-1022) {
						if uint32(v11) >= uint32(i32(-1085)) {
							goto l54
						}
						v12 = i64(0)
						v8 = i32(0)
						goto l47
					}
					p52 := v19
					if i64_shl(v19, v22) == v20 {
						p52 = v19 & i64(0xfffffffffffffc)
					}
					p53 := v19
					if v19&i64(3) == i64(1) {
						p53 = p52
					}
					p54 := v19
					if uint64(v12) < uint64(i64(2)) {
						p54 = p53
					}
					p55 := v19
					if uint64(v15+i64(4)) < uint64(i64(28)) {
						p55 = p54
					}
					v12 = p55
					v12 = v12&i64(1) + v12
					var p56 int32
					if uint64(v12) > uint64(i64(0x3fffffffffffff)) {
						p56 = 1
					}
					v7 = p56
					p57 := i32(1023)
					if v7 != 0 {
						p57 = i32(1024)
					}
					v11 = p57 + v11
					if uint32(v11) <= uint32(i32(2046)) {
						p58 := int64(uint64(v12)>>1) & i64(0x7fefffffffffffff)
						if v7 != 0 {
							p58 = i64(0)
						}
						v12 = p58
						v8 = v11
						goto l47
					}
					v12 = i64(0)
					goto l47
				}
			l54:
				v12 = i64_shr_u(v19, int64(uint32(i32(-1022)-v11)))
				v12 = v12&i64(1) + v12
				var p59 int32
				if uint64(v12) > uint64(i64(0x1fffffffffffff)) {
					p59 = 1
				}
				v8 = p59
				v12 = int64(uint64(v12) >> 1)
			}
		l47:
			if v10 == 0 {
				goto l55
			}
		l51:
			v10 = i32(0)
			v20 = i64(0)
			{
				if v15 < i64(-342) {
					goto l56
				}
				v4 = v4 + i64(1)
				if v4 == 0 {
					goto l56
				}
				v10 = i32(2047)
				if v15 > i64(308) {
					goto l56
				}
				v20 = i64(0)
				t60 := v3 + i32(16)
				v11 = int32(v15)
				v7 = v11 << 4
				t61 := int64(load64(m.memory[uint32(v7+i32(1114680)):]))
				t62 := v4
				v21 = int64(bits.LeadingZeros64(uint64(v4)))
				v18 = i64_shl(t62, v21)
				m.fn982(t60, t61, i64(0), v18, i64(0))
				t63 := int64(load64(m.memory[int64(uint32(v3))+16:]))
				v4 = t63
				{
					t64 := int64(load64(m.memory[int64(uint32(v3))+24:]))
					v19 = t64
					if v19&i64(511) != i64(511) {
						goto l57
					}
					t65 := int64(load64(m.memory[uint32(v7+i32(1109208)+i32(5480)):]))
					m.fn982(v3, t65, i64(0), v18, i64(0))
					t66 := int64(load64(m.memory[int64(uint32(v3))+8:]))
					v18 = t66
					v4 = v18 + v4
					var p67 int32
					if uint64(v4) < uint64(v18) {
						p67 = 1
					}
					v19 = int64(uint32(p67)) + v19
				}
			l57:
				if v4 != i64(-1) {
					goto l58
				}
				if uint64(v15+i64(27)) <= uint64(i64(82)) {
					goto l58
				}
				v10 = i32(-1)
				goto l56
			l58:
				t68 := v19
				v22 = int64(uint64(v19) >> 63)
				v23 = v22 + i64(9)
				v18 = i64_shr_u(t68, v23)
				{
					v11 = v11*i32(217706)>>16 - int32(v21) + int32(v22) + i32(63)
					if v11 < i32(-1022) {
						goto l59
					}
					p69 := v18
					if i64_shl(v18, v23) == v19 {
						p69 = v18 & i64(0xfffffffffffffc)
					}
					p70 := v18
					if v18&i64(3) == i64(1) {
						p70 = p69
					}
					p71 := v18
					if uint64(v4) < uint64(i64(2)) {
						p71 = p70
					}
					p72 := v18
					if uint64(v15+i64(4)) < uint64(i64(28)) {
						p72 = p71
					}
					v4 = p72
					v4 = v4&i64(1) + v4
					var p73 int32
					if uint64(v4) > uint64(i64(0x3fffffffffffff)) {
						p73 = 1
					}
					v7 = p73
					p74 := i32(1023)
					if v7 != 0 {
						p74 = i32(1024)
					}
					v11 = p74 + v11
					if uint32(v11) > uint32(i32(2046)) {
						goto l56
					}
					p75 := int64(uint64(v4)>>1) & i64(0x7fefffffffffffff)
					if v7 != 0 {
						p75 = i64(0)
					}
					v20 = p75
					v10 = v11
					goto l56
				}
			l59:
				v10 = i32(0)
				if uint32(v11) < uint32(i32(-1085)) {
					goto l56
				}
				v4 = i64_shr_u(v18, int64(uint32(i32(-1022)-v11)))
				v4 = v4&i64(1) + v4
				var p76 int32
				if uint64(v4) > uint64(i64(0x1fffffffffffff)) {
					p76 = 1
				}
				v10 = p76
				v20 = int64(uint64(v4) >> 1)
			}
		l56:
			if v12 != v20 {
				goto l50
			}
			if v8 < i32(0) {
				goto l50
			}
			if v8 == v10 {
				goto l55
			}
		l50:
			v10 = i32(0)
			memory_zero(m.memory, uint32(v3+i32(868)), uint32(i32(778)))
			t77 := v3
			var p78 int32
			if v5 == i32(45) {
				p78 = 1
			}
			m.memory[int64(uint32(t77))+1644] = byte(p78)
			v7 = v1
			v14 = v2
			switch v5 + i32(-43) {
			case 0, 2:
				v13 = v1 + i32(1)
				v14 = v2 + i32(-1)
				if v14 == 0 {
					goto l62
				}
				v7 = v13
				fallthrough
			default:
				v13 = v7 + v14
				v8 = v14
			l64:
				{
					v6 = v7 + v10
					t79 := int32(m.memory[uint32(v6)])
					v11 = t79
					if v11 != i32(48) {
						goto l63
					}
					v10 = v10 + i32(1)
					v8 = v8 + i32(-1)
					if v8 != 0 {
						goto l64
					}
				}
			}
		l62:
			v11 = i32(0)
			goto l65
		l63:
			v13 = v11 + i32(-48)
			if uint32(v13&i32(255)) > uint32(i32(9)) {
				if v11 == i32(46) {
					v13 = v6 + i32(1)
					v7 = v8 + i32(-1)
					goto l75
				}
				v14 = i32(0)
				goto l74
			}
			v14 = v14 + i32(-1)
			v11 = i32(0)
		l70:
			{
				{
					if uint32(v11) > uint32(i32(767)) {
						goto l67
					}
					m.memory[uint32(v3+i32(868)+v11)] = byte(v13)
					t80 := int32(load32(m.memory[int64(uint32(v3))+1636:]))
					v11 = t80
				}
			l67:
				t81 := v3
				v11 = v11 + i32(1)
				store32(m.memory[int64(uint32(t81))+1636:], uint32(v11))
				v6 = v7 + v10
				{
					if v14 == v10 {
						goto l68
					}
					v8 = v8 + i32(-1)
					v10 = v10 + i32(1)
					t82 := int32(m.memory[uint32(v6+i32(1))])
					v9 = t82
					v13 = v9 + i32(-48)
					if uint32(v13&i32(255)) > uint32(i32(9)) {
						v6 = v7 + v10
						if v9&i32(255) == i32(46) {
							goto l72
						}
						v13 = v6
						goto l71
					}
					goto l70
				}
			l68:
			}
			v13 = v6 + i32(1)
		l65:
			v8 = i32(0)
			goto l71
		l72:
			v13 = v6 + i32(-1) + i32(2)
			v7 = v8 + i32(1) + i32(-2)
			v10 = v7
			if v11 != 0 {
				goto l76
			}
		l75:
			if v7 != 0 {
				goto l77
			}
			v7 = i32(0)
			v11 = i32(0)
			goto l78
		l77:
			v6 = v6 + v8
			v10 = i32(0)
		l80:
			{
				v8 = v13 + v10
				t83 := int32(m.memory[uint32(v8)])
				if t83 != i32(48) {
					goto l79
				}
				t84 := v7
				v10 = v10 + i32(1)
				if t84 != v10 {
					goto l80
				}
			}
			v11 = i32(0)
			v8 = i32(0)
			v13 = v6
			goto l81
		l79:
			v10 = v7 - v10
			v11 = i32(0)
			v13 = v8
		l76:
			if uint32(v10) < uint32(i32(8)) {
				goto l82
			}
		l85:
			{
				if uint32(v11+i32(8)) >= uint32(i32(768)) {
					goto l88
				}
				t85 := int64(load64(m.memory[uint32(v13):]))
				v4 = t85
				t86 := v4 + i64(5063812098665367110)
				v4 = v4 + i64(-3472328296227680304)
				if (t86|v4)&i64(-0x7f7f7f7f7f7f7f80) != i64(0) {
					goto l88
				}
				if uint32(v11) >= uint32(i32(769)) {
					m.fn127(v11, i32(768), i32(768), i32(1090952))
					panic("unreachable")
				}
				store64(m.memory[uint32(v3+i32(868)+v11):], uint64(v4))
				t87 := int32(load32(m.memory[int64(uint32(v3))+1636:]))
				t88 := v3
				v11 = t87 + i32(8)
				store32(m.memory[int64(uint32(t88))+1636:], uint32(v11))
				v13 = v13 + i32(8)
				v10 = v10 + i32(-8)
				if uint32(v10) > uint32(i32(7)) {
					goto l85
				}
			}
		l82:
			if v10 == 0 {
				goto l78
			}
		l88:
			{
				{
					t89 := int32(m.memory[uint32(v13)])
					v8 = t89 + i32(-48)
					if uint32(v8&i32(255)) <= uint32(i32(9)) {
						goto l86
					}
					v8 = v10
					goto l81
				}
			l86:
				{
					if uint32(v11) > uint32(i32(767)) {
						goto l87
					}
					m.memory[uint32(v3+i32(868)+v11)] = byte(v8)
					t90 := int32(load32(m.memory[int64(uint32(v3))+1636:]))
					v11 = t90
				}
			l87:
				t91 := v3
				v11 = v11 + i32(1)
				store32(m.memory[int64(uint32(t91))+1636:], uint32(v11))
				v13 = v13 + i32(1)
				v10 = v10 + i32(-1)
				if v10 != 0 {
					goto l88
				}
			}
		l78:
			v8 = i32(0)
		l81:
			store32(m.memory[int64(uint32(v3))+1640:], uint32(v8-v7))
			goto l71
		l71:
			{
				if v11 != 0 {
					goto l89
				}
				v14 = i32(0)
				goto l90
			l89:
				v10 = v2 - v8
				if uint32(v2) < uint32(v8) {
					m.fn127(i32(0), v10, v2, i32(1090968))
					panic("unreachable")
				}
				v7 = i32(0)
				if v2 == v8 {
					goto l92
				}
				v6 = v1 + i32(-1)
				v7 = i32(0)
			l95:
				{
					t92 := int32(m.memory[uint32(v6+v10)])
					switch t92 + i32(-46) {
					default:
						goto l92
					case 2:
						v7 = v7 + i32(1)
						fallthrough
					case 0:
						v10 = v10 + i32(-1)
						if v10 != 0 {
							goto l95
						}
					}
				}
			l92:
				t93 := int32(load32(m.memory[int64(uint32(v3))+1640:]))
				store32(m.memory[int64(uint32(v3))+1640:], uint32(t93+v11))
				t94 := v3
				v14 = v11 - v7
				store32(m.memory[int64(uint32(t94))+1636:], uint32(v14))
				if uint32(v14) < uint32(i32(769)) {
					goto l90
				}
				v14 = i32(768)
				store32(m.memory[int64(uint32(v3))+1636:], uint32(i32(768)))
				m.memory[int64(uint32(v3))+1645] = byte(i32(1))
			}
		l90:
			v6 = v13
		l74:
			{
				if v8 == 0 {
					goto l96
				}
				t95 := int32(m.memory[uint32(v6)])
				if t95|i32(32) != i32(101) {
					goto l96
				}
				v7 = i32(0)
				v13 = v8 + i32(-1)
				if v13 == 0 {
					goto l97
				}
				v11 = v6 + i32(1)
				{
					t96 := int32(m.memory[int64(uint32(v6))+1])
					switch t96 + i32(-43) {
					case 0:
						v13 = v8 + i32(-2)
						if v13 == 0 {
							goto l97
						}
						v11 = v6 + i32(2)
						fallthrough
					default:
						v7 = i32(0)
						v10 = i32(0)
					l101:
						{
							t97 := int32(m.memory[uint32(v11)])
							v8 = (t97 + i32(-48)) & i32(255)
							if uint32(v8) > uint32(i32(9)) {
								goto l97
							}
							v8 = v10*i32(10) + v8
							t98 := v8
							t99 := v10
							var p100 int32
							if v10 < i32(65536) {
								p100 = 1
							}
							v6 = p100
							p101 := t99
							if v6 != 0 {
								p101 = t98
							}
							v10 = p101
							p102 := v7
							if v6 != 0 {
								p102 = v8
							}
							v7 = p102
							v11 = v11 + i32(1)
							v13 = v13 + i32(-1)
							if v13 != 0 {
								goto l101
							}
							goto l97
						}
					case 2:
						v11 = i32(0)
						v7 = v8 + i32(-2)
						if v7 == 0 {
							goto l102
						}
						v8 = v6 + i32(2)
						v11 = i32(0)
						v10 = i32(0)
					l103:
						{
							t103 := int32(m.memory[uint32(v8)])
							v6 = (t103 + i32(-48)) & i32(255)
							if uint32(v6) > uint32(i32(9)) {
								goto l102
							}
							v6 = v10*i32(10) + v6
							t104 := v6
							t105 := v10
							var p106 int32
							if v10 < i32(65536) {
								p106 = 1
							}
							v13 = p106
							p107 := t105
							if v13 != 0 {
								p107 = t104
							}
							v10 = p107
							p108 := v11
							if v13 != 0 {
								p108 = v6
							}
							v11 = p108
							v8 = v8 + i32(1)
							v7 = v7 + i32(-1)
							if v7 != 0 {
								goto l103
							}
						}
					l102:
						v7 = i32(0) - v11
					}
				}
			l97:
				t109 := int32(load32(m.memory[int64(uint32(v3))+1640:]))
				store32(m.memory[int64(uint32(v3))+1640:], uint32(t109+v7))
			}
		l96:
			if uint32(v14) > uint32(i32(18)) {
				goto l104
			}
			v10 = i32(19) - v14
			if v10 == 0 {
				goto l104
			}
			memory_zero(m.memory, uint32(v3+i32(868)+v14), uint32(v10))
		l104:
			memory_copy(m.memory, uint32(v3+i32(88)), uint32(v3+i32(868)), uint32(i32(780)))
			v8 = i32(0)
			v12 = i64(0)
			t110 := int32(load32(m.memory[int64(uint32(v3))+856:]))
			if t110 == 0 {
				goto l55
			}
			t111 := int32(load32(m.memory[int64(uint32(v3))+860:]))
			v10 = t111
			if v10 < i32(-324) {
				goto l55
			}
			v8 = i32(2047)
			if v10 > i32(309) {
				goto l55
			}
			if v10 >= i32(1) {
				v11 = i32(0)
			l109:
				v7 = i32(60)
				{
					if uint32(v10) >= uint32(i32(19)) {
						goto l107
					}
					t112 := int32(m.memory[int64(uint32(v10))+1099028])
					v7 = t112
				}
			l107:
				m.fn682(v3+i32(88), v7)
				{
					t113 := int32(load32(m.memory[int64(uint32(v3))+860:]))
					v10 = t113
					if v10 <= i32(-2048) {
						v8 = i32(0)
						goto l55
					}
					v11 = v7 + v11
					if v10 < i32(1) {
						goto l114
					}
					goto l109
				}
			}
			v11 = i32(0)
			goto l114
		l114:
			{
				{
					if v10 != 0 {
						goto l110
					}
					t114 := int32(m.memory[int64(uint32(v3))+88])
					v10 = t114
					if uint32(v10) > uint32(i32(4)) {
						goto l111
					}
					p115 := i32(1)
					if uint32(v10) < uint32(i32(2)) {
						p115 = i32(2)
					}
					v7 = p115
					goto l112
				}
			l110:
				v7 = i32(60)
				v10 = i32(0) - v10
				if uint32(v10) >= uint32(i32(19)) {
					goto l112
				}
				t116 := int32(m.memory[int64(uint32(v10))+1099028])
				v7 = t116
			}
		l112:
			m.fn683(v3+i32(88), v7)
			{
				t117 := int32(load32(m.memory[int64(uint32(v3))+860:]))
				v10 = t117
				if v10 <= i32(2047) {
					goto l113
				}
				v8 = i32(2047)
				goto l55
			}
		l113:
			v11 = v11 - v7
			if v10 < i32(1) {
				goto l114
			}
		l111:
			v10 = v11 + i32(-1)
			if v10 > i32(-1023) {
				goto l115
			}
		l116:
			{
				t118 := v3 + i32(88)
				v11 = i32(-1022) - v10
				p119 := i32(60)
				if uint32(v11) < uint32(i32(60)) {
					p119 = v11
				}
				v11 = p119
				m.fn682(t118, v11)
				v10 = v11 + v10
				if uint32(v10) < uint32(i32(-1022)) {
					goto l116
				}
			}
		l115:
			if v10+i32(1023) > i32(2046) {
				goto l55
			}
			m.fn683(v3+i32(88), i32(53))
			{
				{
					{
						t120 := int32(load32(m.memory[int64(uint32(v3))+856:]))
						v7 = t120
						if v7 == 0 {
							goto l117
						}
						t121 := int32(load32(m.memory[int64(uint32(v3))+860:]))
						v14 = t121
						if v14 < i32(0) {
							goto l117
						}
						if uint32(v14) > uint32(i32(18)) {
							goto l118
						}
						if v14 != 0 {
							if v14 != i32(1) {
								v2 = v14 & i32(1)
								v13 = v14 & i32(30)
								v6 = i32(0)
								v4 = i64(0)
							l126:
								v4 = v4 * i64(10)
								{
									v11 = v6
									if uint32(v11) >= uint32(v7) {
										goto l123
									}
									t122 := int64(m.memory[uint32(v3+i32(88)+v11)])
									v4 = v4 + t122
								}
							l123:
								v4 = v4 * i64(10)
								{
									v6 = v11 + i32(1)
									if uint32(v6) >= uint32(v7) {
										goto l124
									}
									t123 := int64(m.memory[uint32(v3+i32(88)+v11+i32(1))])
									v4 = v4 + t123
								}
							l124:
								v6 = v6 + i32(1)
								if v6 == v13 {
									goto l125
								}
								goto l126
							}
							v11 = i32(0)
							v4 = i64(0)
							goto l122
						}
						v4 = i64(0)
						goto l120
					}
				l117:
					v8 = v10 + i32(1022)
					goto l55
				l125:
					if v2 == 0 {
						goto l120
					}
					v11 = v11 + i32(2)
				l122:
					v4 = v4 * i64(10)
					if uint32(v11) >= uint32(v7) {
						goto l120
					}
					t124 := int64(m.memory[uint32(v3+i32(88)+v11)])
					v4 = v4 + t124
				}
			l120:
				{
					if uint32(v14) >= uint32(v7) {
						goto l127
					}
					v6 = v3 + i32(88) + v14
					t125 := int32(m.memory[uint32(v6)])
					v11 = t125
					{
						if v14+i32(1) != v7 {
							goto l128
						}
						if v11&i32(255) == i32(5) {
							goto l129
						}
					l128:
						if uint32(v11&i32(255)) > uint32(i32(4)) {
							goto l130
						}
						goto l127
					l129:
						t126 := int32(m.memory[int64(uint32(v3))+865])
						if t126 != 0 {
							goto l130
						}
						if v14 == 0 {
							goto l127
						}
						t127 := int32(m.memory[uint32(v6+i32(-1))])
						if t127&i32(1) == 0 {
							goto l127
						}
					}
				l130:
					v4 = v4 + i64(1)
				}
			l127:
				if uint64(v4) < uint64(i64(0x20000000000000)) {
					goto l131
				}
			l118:
				m.fn682(v3+i32(88), i32(1))
				t128 := m.fn684(v3 + i32(88))
				v4 = t128
				if v10+i32(1024) > i32(2046) {
					goto l55
				}
				v10 = v10 + i32(1)
			}
		l131:
			v12 = v4 & i64(0xfffffffffffff)
			p129 := i32(1023)
			if uint64(v4) < uint64(i64(0x10000000000000)) {
				p129 = i32(1022)
			}
			v8 = p129 + v10
		}
	l55:
		t130 := v0
		v4 = int64(uint32(v8))<<52 | v12
		p131 := v4
		if v5 == i32(45) {
			p131 = v4 | i64(-0x8000000000000000)
		}
		store64(m.memory[int64(uint32(t130))+8:], uint64(p131))
		goto l46
	}
l46:
	store32(m.memory[int64(uint32(v0))+16:], uint32(v16))
	v4 = i64(1)
l0:
	store64(m.memory[uint32(v0):], uint64(v4))
	m.g0 = v3 + i32(1648)
}
func (m *Module) fn679(v0, v1, v2 int32) {
	var v3 int32
	t0 := m.g0
	v3 = t0 - i32(32)
	m.g0 = v3
	m.fn584(v3+i32(8), v1, v2)
	{
		t1 := int32(m.memory[int64(uint32(v3))+8])
		if t1 != i32(1) {
			t3 := math.Float64frombits(load64(m.memory[int64(uint32(v3))+16:]))
			store64(m.memory[int64(uint32(v3))+24:], math.Float64bits(t3))
			m.fn48(i32(1080252), i32(46), v3+i32(24), i32(1080236), i32(1070064))
			panic("unreachable")
		}
		t2 := int32(m.memory[int64(uint32(v3))+9])
		v2 = t2
		store32(m.memory[uint32(v0):], uint32(i32(-0x7fffffeb)))
		m.memory[int64(uint32(v0))+4] = byte(v2)
		m.g0 = v3 + i32(32)
		return
	}
}
func (m *Module) fn680(v0 int32) {
	var v1, v2, v3 int32
	{
		t0 := int32(m.memory[uint32(v0)])
		switch t0 + i32(-2) {
		default:
			return
		case 0, 3, 4:
			t1 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			v1 = t1
			if v1 == 0 {
				return
			}
			t2 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			v2 = t2
			t3 := int32(load32(m.memory[uint32(v2+i32(-4)):]))
			v0 = t3
			v3 = v0 & i32(-8)
			t4 := v3
			v0 = v0 & i32(3)
			p5 := i32(8)
			if v0 != 0 {
				p5 = i32(4)
			}
			if uint32(t4) < uint32(p5+v1) {
				m.fn2(i32(1273840), i32(46), i32(1273888))
				panic("unreachable")
			}
			if v0 == 0 {
				goto l3
			}
			if uint32(v3) > uint32(v1+i32(39)) {
				m.fn2(i32(1273904), i32(46), i32(1273952))
				panic("unreachable")
			}
		l3:
			m.fn1(v2)
		}
	}
}
func (m *Module) fn681(v0, v1 int32) int32 {
	var v2 int32
	v2 = i32(3)
	{
		if uint32(v1) < uint32(i32(8)) {
			goto l0
		}
		t0 := int32(m.memory[int64(uint32(v0))+4])
		t1 := int32(m.memory[int64(uint32(v0))+3])
		t2 := int32(m.memory[int64(uint32(v0))+5])
		t3 := int32(m.memory[int64(uint32(v0))+6])
		t4 := int32(m.memory[int64(uint32(v0))+7])
		p5 := i32(8)
		if (t0^i32(78)|(t1^i32(73))|(t2^i32(73))|(t3^i32(84))|(t4^i32(89)))&i32(223) != 0 {
			p5 = i32(3)
		}
		v2 = p5
	}
l0:
	return v2
}
func (m *Module) fn682(v0, v1 int32) {
	var v2, v3, v4 int32
	var v5, v6 int64
	var v7, v8 int32
	var v9, v10 int64
	v2 = v0 + i32(768)
	t0 := int32(load32(m.memory[int64(uint32(v0))+768:]))
	v3 = t0
	v4 = i32(0) - v3
	v5 = int64(uint32(v1 & i32(63)))
	v1 = i32(-768)
	v6 = i64(0)
	{
	l3:
		{
			v7 = v4 + v1
			if v7 == i32(-768) {
				goto l0
			}
			if v1 == 0 {
				m.fn39(i32(768), i32(768), i32(1080536))
				panic("unreachable")
			}
			t1 := v6 * i64(10)
			v8 = v0 + v1
			t2 := int64(m.memory[uint32(v8+i32(768))])
			v6 = t1 + t2
			if i64_shr_u(v6, v5) != i64(0) {
				v7 = v1 + i32(769)
				goto l4
			}
			if v7 == i32(-769) {
				goto l0
			}
			v1 = v1 + i32(2)
			t3 := int64(m.memory[uint32(v8+i32(769))])
			v6 = v6*i64(10) + t3
			if i64_shr_u(v6, v5) == 0 {
				goto l3
			}
		}
		v7 = v1 + i32(768)
		goto l4
	l0:
		if v6 == 0 {
			return
		}
		if i64_shr_u(v6, v5) == i64(0) {
			goto l6
		}
		v7 = v3
		goto l4
	l6:
		v7 = v3
	l7:
		v7 = v7 + i32(1)
		v6 = v6 * i64(10)
		if i64_shr_u(v6, v5) == 0 {
			goto l7
		}
	l4:
		t4 := int32(load32(m.memory[int64(uint32(v0))+772:]))
		t5 := v0
		v1 = t4 - v7 + i32(1)
		store32(m.memory[int64(uint32(t5))+772:], uint32(v1))
		{
			if v1 < i32(-2047) {
				goto l8
			}
			v9 = i64_shl(i64(-1), v5) ^ i64(-1)
			v1 = i32(0)
			{
				if uint32(v7) >= uint32(v3) {
					goto l9
				}
				v1 = i32(0)
				v8 = i32(768) - v7
				p6 := v8
				if uint32(v8) > uint32(i32(768)) {
					p6 = i32(0)
				}
				v8 = p6
				v4 = v0 + v7
			l11:
				{
					if v8 != v1 {
						goto l10
					}
					m.fn39(v7+v1, i32(768), i32(1080552))
					panic("unreachable")
				l10:
					t7 := int64(m.memory[uint32(v4+v1)])
					v10 = t7
					m.memory[uint32(v0+v1)] = byte(i64_shr_u(v6, v5))
					v6 = v10 + v6&v9*i64(10)
					t8 := v7
					v1 = v1 + i32(1)
					t9 := int32(load32(m.memory[int64(uint32(v0))+768:]))
					if uint32(t8+v1) < uint32(t9) {
						goto l11
					}
				}
			}
		l9:
			if v6 == 0 {
				goto l12
			}
		l15:
			v10 = v6
			v6 = v10 & v9 * i64(10)
			v7 = int32(i64_shr_u(v10, v5))
			if uint32(v1) < uint32(i32(768)) {
				goto l13
			}
			if v7&i32(255) == 0 {
				goto l14
			}
			m.memory[int64(uint32(v0))+777] = byte(i32(1))
			goto l14
		l13:
			m.memory[uint32(v0+v1)] = byte(v7)
			v1 = v1 + i32(1)
		l14:
			if !(v6 == 0) {
				goto l15
			}
		l12:
			v4 = v0 + i32(-1)
			var p10 int32
			if uint32(v1) > uint32(i32(768)) {
				p10 = 1
			}
			v8 = p10
		l17:
			store32(m.memory[uint32(v2):], uint32(v1))
			if v1 == 0 {
				return
			}
			v0 = v1 + i32(-1)
			{
				if v8 != 0 {
					m.fn39(v0, i32(768), i32(1080504))
					panic("unreachable")
				}
				v7 = v4 + v1
				v1 = v0
				t11 := int32(m.memory[uint32(v7)])
				if t11 == 0 {
					goto l17
				}
				return
			}
		}
	l8:
		store16(m.memory[int64(uint32(v2))+8:], uint16(i32(0)))
		store64(m.memory[uint32(v2):], uint64(i64(0)))
	}
}
func (m *Module) fn683(v0, v1 int32) {
	var v2, v3, v4, v5, v6, v7, v8, v9 int32
	var v10, v11, v12, v13 int64
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+768:]))
		v2 = t0
		if v2 == 0 {
			return
		}
		v3 = v1 & i32(63)
		v1 = v3 << 1
		t1 := int32(load16(m.memory[int64(uint32(v1))+1107736:]))
		v4 = t1
		v5 = v4 & i32(2047)
		t2 := int32(load16(m.memory[int64(uint32(v1))+1107738:]))
		v6 = v5 - t2&i32(2047)
		v7 = i32(0) - v2
		v8 = int32(uint32(v4) >> 11)
		v1 = i32(-1308)
		{
		l4:
			{
				if v6+v1 == i32(-1308) {
					goto l1
				}
				v4 = v5 + v1
				if v4 == 0 {
					goto l1
				}
				if v7+v1 == i32(-1308) {
					v8 = v8 + i32(-1)
					goto l1
				}
				if v1 == i32(-540) {
					m.fn39(i32(768), i32(768), i32(1090984))
					panic("unreachable")
				}
				v9 = v0 + v1
				v1 = v1 + i32(1)
				t3 := int32(m.memory[uint32(v9+i32(1308))])
				v9 = t3
				t4 := int32(m.memory[uint32(v4+i32(1109174))])
				t5 := v9
				v4 = t4
				if t5 == v4&i32(255) {
					goto l4
				}
			}
			t6 := v8
			var p7 int32
			if uint32(v9) < uint32(v4&i32(255)) {
				p7 = 1
			}
			v8 = t6 - p7
			goto l1
		}
	l1:
		v4 = v0 + i32(-1)
		t8 := v0
		v9 = v8 + i32(-1)
		v6 = t8 + v9
		v10 = int64(uint32(v3))
		v11 = i64(0)
	l8:
		v1 = v2
		v2 = v1 + i32(-1)
		{
			if uint32(v1) >= uint32(i32(769)) {
				m.fn39(v2, i32(768), i32(1080520))
				panic("unreachable")
			}
			t9 := int64(m.memory[uint32(v4+v1)])
			v12 = i64_shl(t9, v10) + v11
			t10 := int64(uint64(v12) / uint64(i64(10)))
			t11 := v12
			v11 = t10
			v13 = t11 + v11*i64(-10)
			if uint32(v9+v1) < uint32(i32(768)) {
				goto l6
			}
			if v13 == 0 {
				goto l7
			}
			m.memory[int64(uint32(v0))+777] = byte(i32(1))
			goto l7
		}
	l6:
		m.memory[uint32(v6+v1)] = byte(v13)
	l7:
		if v2 != 0 {
			goto l8
		}
		if uint64(v12) < uint64(i64(10)) {
			goto l9
		}
		v1 = v8 + i32(-1)
	l12:
		{
			v12 = v11
			t12 := int64(uint64(v12) / uint64(i64(10)))
			t13 := v12
			v11 = t12
			v13 = t13 + v11*i64(-10)
			if uint32(v1) < uint32(i32(768)) {
				goto l10
			}
			if v13 == 0 {
				goto l11
			}
			m.memory[int64(uint32(v0))+777] = byte(i32(1))
			goto l11
		l10:
			m.memory[uint32(v0+v1)] = byte(v13)
		l11:
			v1 = v1 + i32(-1)
			if uint64(v12) >= uint64(i64(10)) {
				goto l12
			}
		}
	l9:
		t14 := int32(load32(m.memory[int64(uint32(v0))+772:]))
		store32(m.memory[int64(uint32(v0))+772:], uint32(t14+v8))
		t15 := int32(load32(m.memory[int64(uint32(v0))+768:]))
		t16 := v0
		v2 = t15 + v8
		p17 := i32(768)
		if uint32(v2) < uint32(i32(768)) {
			p17 = v2
		}
		v1 = p17
		store32(m.memory[int64(uint32(t16))+768:], uint32(v1))
		if v2 == 0 {
			return
		}
		v4 = v0 + i32(-1)
	l15:
		v2 = v1 + i32(-1)
		{
			if uint32(v1) > uint32(i32(768)) {
				m.fn39(v2, i32(768), i32(1080504))
				panic("unreachable")
			}
			t18 := int32(m.memory[uint32(v4+v1)])
			if t18 == 0 {
				goto l14
			}
			return
		}
	l14:
		store32(m.memory[int64(uint32(v0))+768:], uint32(v2))
		v1 = v2
		if v2 != 0 {
			goto l15
		}
	}
}
func (m *Module) fn684(v0 int32) int64 {
	var v1 int64
	var v2, v3, v4, v5, v6, v7 int32
	v1 = i64(0)
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+768:]))
		v2 = t0
		if v2 == 0 {
			goto l0
		}
		t1 := int32(load32(m.memory[int64(uint32(v0))+772:]))
		v3 = t1
		if v3 < i32(0) {
			goto l0
		}
		v1 = i64(-1)
		if uint32(v3) > uint32(i32(18)) {
			goto l0
		}
		{
			if v3 != 0 {
				goto l1
			}
			v1 = i64(0)
			goto l2
		l1:
			if v3 != i32(1) {
				goto l3
			}
			v4 = i32(0)
			v1 = i64(0)
			goto l4
		l3:
			v5 = v3 & i32(1)
			v6 = v3 & i32(30)
			v7 = i32(0)
			v1 = i64(0)
		l7:
			v1 = v1 * i64(10)
			{
				v4 = v7
				if uint32(v4) >= uint32(v2) {
					goto l5
				}
				t2 := int64(m.memory[uint32(v0+v4)])
				v1 = v1 + t2
			}
		l5:
			v1 = v1 * i64(10)
			{
				v7 = v4 + i32(1)
				if uint32(v7) >= uint32(v2) {
					goto l6
				}
				t3 := int64(m.memory[uint32(v0+v4+i32(1))])
				v1 = v1 + t3
			}
		l6:
			v7 = v7 + i32(1)
			if v7 != v6 {
				goto l7
			}
			if v5 == 0 {
				goto l2
			}
			v4 = v4 + i32(2)
		l4:
			v1 = v1 * i64(10)
			if uint32(v4) >= uint32(v2) {
				goto l2
			}
			t4 := int64(m.memory[uint32(v0+v4)])
			v1 = v1 + t4
		}
	l2:
		if uint32(v3) >= uint32(v2) {
			goto l0
		}
		v7 = v0 + v3
		t5 := int32(m.memory[uint32(v7)])
		v4 = t5
		{
			if v3+i32(1) != v2 {
				goto l8
			}
			if v4&i32(255) == i32(5) {
				goto l9
			}
		l8:
			if uint32(v4&i32(255)) > uint32(i32(4)) {
				goto l10
			}
			goto l0
		l9:
			t6 := int32(m.memory[int64(uint32(v0))+777])
			if t6 != 0 {
				goto l10
			}
			if v3 == 0 {
				goto l0
			}
			t7 := int32(m.memory[uint32(v7+i32(-1))])
			if t7&i32(1) == 0 {
				goto l0
			}
		}
	l10:
		v1 = v1 + i64(1)
	}
l0:
	return v1
}
func (m *Module) fn685(v0, v1, v2, v3 int32) int32 {
	var v4, v5, v6 int32
	var v7, v8 int64
	t0 := m.g0
	v4 = t0 - i32(16)
	m.g0 = v4
	v5 = i32(1)
	{
		t1 := int32(load32(m.memory[int64(uint32(v1))+12:]))
		t2 := v0
		v6 = t1
		t3 := m.t0[uint(v6)].(func(int32, int32, int32) int32)(t2, i32(1272329), i32(1))
		if t3 != 0 {
			goto l0
		}
		{
			if v3 == 0 {
				goto l1
			}
			v7 = int64(uint32(i32(45)))<<32 | int64(uint32(v4+i32(4)))
			v8 = int64(uint32(i32(62)))<<32 | int64(uint32(v4))
		l7:
			store32(m.memory[uint32(v4):], uint32(v2))
			{
				{
					{
						t4 := int32(m.memory[uint32(v2)])
						v5 = t4
						if v5&i32(254) == i32(32) {
							goto l2
						}
						if uint32((v5+i32(-35))&i32(255)) > uint32(i32(91)) {
							goto l3
						}
					}
				l2:
					store32(m.memory[int64(uint32(v4))+4:], uint32(v5))
					store64(m.memory[int64(uint32(v4))+8:], uint64(v7))
					t5 := m.fn51(v0, v1, i32(1052562), v4+i32(8))
					if t5 != 0 {
						goto l4
					}
					goto l5
				}
			l3:
				{
					if v5 == i32(34) {
						goto l6
					}
					store64(m.memory[int64(uint32(v4))+8:], uint64(v8))
					t6 := m.fn51(v0, v1, i32(1272332), v4+i32(8))
					if t6 == 0 {
						goto l5
					}
					goto l4
				}
			l6:
				t7 := m.t0[uint(v6)].(func(int32, int32, int32) int32)(v0, i32(1272330), i32(2))
				if t7 != 0 {
					goto l4
				}
			}
		l5:
			v2 = v2 + i32(1)
			v3 = v3 + i32(-1)
			if v3 != 0 {
				goto l7
			}
		l1:
			t8 := m.t0[uint(v6)].(func(int32, int32, int32) int32)(v0, i32(1272329), i32(1))
			v5 = t8
			goto l0
		}
	l4:
		v5 = i32(1)
	}
l0:
	m.g0 = v4 + i32(16)
	return v5
}
func (m *Module) fn686(v0, v1 int32) {
	var v2, v3, v4 int32
	t0 := m.g0
	v2 = t0 - i32(32)
	m.g0 = v2
	v3 = i32(0)
	{
		{
			t1 := int32(m.memory[int64(uint32(v1))+37])
			if t1 == 0 {
				goto l0
			}
			goto l1
		}
	l0:
		{
			t2 := int32(m.memory[int64(uint32(v1))+36])
			if t2 != 0 {
				goto l2
			}
			m.memory[int64(uint32(v1))+36] = byte(i32(1))
			m.fn686(v2+i32(8), v1)
			{
				t3 := int32(load32(m.memory[int64(uint32(v2))+8:]))
				v3 = t3
				if v3 == 0 {
					goto l3
				}
				t4 := int32(load32(m.memory[int64(uint32(v2))+12:]))
				v4 = t4
				if v4 != 0 {
					goto l1
				}
			}
		l3:
			v3 = i32(0)
			t5 := int32(m.memory[int64(uint32(v1))+37])
			if t5 == i32(1) {
				goto l1
			}
		}
	l2:
		t6 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		v3 = t6
		m.fn158(v2+i32(20), v1)
		{
			t7 := int32(load32(m.memory[int64(uint32(v2))+20:]))
			if t7 != 0 {
				goto l4
			}
			m.memory[int64(uint32(v1))+37] = byte(i32(1))
			t8 := int32(load32(m.memory[int64(uint32(v1))+32:]))
			t9 := int32(load32(m.memory[int64(uint32(v1))+28:]))
			v1 = t9
			v4 = t8 - v1
			v3 = v3 + v1
			goto l1
		}
	l4:
		t10 := int32(load32(m.memory[int64(uint32(v1))+32:]))
		v4 = t10
		t11 := int32(load32(m.memory[int64(uint32(v2))+24:]))
		store32(m.memory[int64(uint32(v1))+32:], uint32(t11))
		t12 := int32(load32(m.memory[int64(uint32(v2))+28:]))
		t13 := v4
		v1 = t12
		v4 = t13 - v1
		v3 = v3 + v1
	}
l1:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v4))
	store32(m.memory[uint32(v0):], uint32(v3))
	m.g0 = v2 + i32(32)
}
func (m *Module) fn687(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6 int32
	t0 := m.g0
	v2 = t0 - i32(48)
	m.g0 = v2
	{
		{
			{
				t1 := int32(m.memory[uint32(v0)])
				switch t1 {
				default:
					v3 = i32(1)
					t2 := int32(load32(m.memory[uint32(v1):]))
					v4 = t2
					t3 := int32(load32(m.memory[int64(uint32(v1))+4:]))
					t4 := v4
					v5 = t3
					t5 := int32(load32(m.memory[int64(uint32(v5))+12:]))
					v6 = t5
					t6 := m.t0[uint(v6)].(func(int32, int32, int32) int32)(t4, i32(1091476), i32(2))
					if t6 != 0 {
						goto l7
					}
					v0 = v0 + i32(4)
					{
						{
							t7 := int32(m.memory[int64(uint32(v1))+10])
							if t7&i32(128) != 0 {
								goto l8
							}
							v3 = i32(1)
							t8 := m.t0[uint(v6)].(func(int32, int32, int32) int32)(v4, i32(1099059), i32(1))
							if t8 != 0 {
								goto l7
							}
							t9 := m.fn351(v0, v1)
							if t9 == 0 {
								goto l9
							}
							goto l7
						}
					l8:
						t10 := m.t0[uint(v6)].(func(int32, int32, int32) int32)(v4, i32(1099060), i32(2))
						if t10 != 0 {
							goto l7
						}
						v3 = i32(1)
						m.memory[int64(uint32(v2))+12] = byte(i32(1))
						store32(m.memory[int64(uint32(v2))+20:], uint32(v5))
						store32(m.memory[int64(uint32(v2))+16:], uint32(v4))
						store32(m.memory[int64(uint32(v2))+32:], uint32(i32(1099936)))
						t11 := int64(load64(m.memory[int64(uint32(v1))+8:]))
						store64(m.memory[int64(uint32(v2))+36:], uint64(t11))
						store32(m.memory[int64(uint32(v2))+24:], uint32(v2+i32(12)))
						store32(m.memory[int64(uint32(v2))+28:], uint32(v2+i32(16)))
						t12 := m.fn351(v0, v2+i32(28))
						if t12 != 0 {
							goto l7
						}
						t13 := int32(load32(m.memory[int64(uint32(v2))+28:]))
						t14 := int32(load32(m.memory[int64(uint32(v2))+32:]))
						t15 := int32(load32(m.memory[int64(uint32(t14))+12:]))
						t16 := m.t0[uint(t15)].(func(int32, int32, int32) int32)(t13, i32(1099057), i32(2))
						if t16 != 0 {
							goto l7
						}
					}
				l9:
					t17 := int32(load32(m.memory[uint32(v1):]))
					t18 := int32(load32(m.memory[int64(uint32(v1))+4:]))
					t19 := int32(load32(m.memory[int64(uint32(t18))+12:]))
					t20 := m.t0[uint(t19)].(func(int32, int32, int32) int32)(t17, i32(1272328), i32(1))
					v3 = t20
					goto l7
				case 1:
					store32(m.memory[int64(uint32(v2))+16:], uint32(v0+i32(1)))
					t21 := int32(load32(m.memory[uint32(v1):]))
					t22 := int32(load32(m.memory[int64(uint32(v1))+4:]))
					t23 := int32(load32(m.memory[int64(uint32(t22))+12:]))
					t24 := m.t0[uint(t23)].(func(int32, int32, int32) int32)(t21, i32(1079556), i32(3))
					v3 = t24
					m.memory[int64(uint32(v2))+33] = byte(i32(0))
					m.memory[int64(uint32(v2))+32] = byte(v3)
					store32(m.memory[int64(uint32(v2))+28:], uint32(v1))
					t25 := m.fn350(v2+i32(28), i32(1079559), i32(3), v0+i32(12), i32(82))
					t26 := m.fn350(t25, i32(1079562), i32(9), v2+i32(16), i32(86))
					v4 = t26
					t27 := int32(m.memory[int64(uint32(v2))+33])
					v1 = t27
					t28 := int32(m.memory[int64(uint32(v2))+32])
					t29 := v1
					v0 = t28
					v3 = t29 | v0
					if v1 != i32(1) {
						goto l7
					}
					if v0&i32(1) != 0 {
						goto l7
					}
					{
						t30 := int32(load32(m.memory[uint32(v4):]))
						v1 = t30
						t31 := int32(m.memory[int64(uint32(v1))+10])
						if t31&i32(128) != 0 {
							t36 := int32(load32(m.memory[uint32(v1):]))
							t37 := int32(load32(m.memory[int64(uint32(v1))+4:]))
							t38 := int32(load32(m.memory[int64(uint32(t37))+12:]))
							t39 := m.t0[uint(t38)].(func(int32, int32, int32) int32)(t36, i32(1099063), i32(1))
							v3 = t39
							goto l7
						}
						t32 := int32(load32(m.memory[uint32(v1):]))
						t33 := int32(load32(m.memory[int64(uint32(v1))+4:]))
						t34 := int32(load32(m.memory[int64(uint32(t33))+12:]))
						t35 := m.t0[uint(t34)].(func(int32, int32, int32) int32)(t32, i32(1273624), i32(2))
						v3 = t35
						goto l7
					}
				case 2:
					t40 := int32(load32(m.memory[uint32(v1):]))
					t41 := int32(load32(m.memory[int64(uint32(v1))+4:]))
					t42 := int32(load32(m.memory[int64(uint32(t41))+12:]))
					t43 := m.t0[uint(t42)].(func(int32, int32, int32) int32)(t40, i32(1079571), i32(12))
					v3 = t43
					goto l7
				case 3:
					v3 = i32(1)
					t44 := int32(load32(m.memory[uint32(v1):]))
					v4 = t44
					t45 := int32(load32(m.memory[int64(uint32(v1))+4:]))
					t46 := v4
					v5 = t45
					t47 := int32(load32(m.memory[int64(uint32(v5))+12:]))
					v6 = t47
					t48 := m.t0[uint(v6)].(func(int32, int32, int32) int32)(t46, i32(1079583), i32(14))
					if t48 != 0 {
						goto l7
					}
					{
						{
							t49 := int32(m.memory[int64(uint32(v1))+10])
							if t49&i32(128) != 0 {
								goto l11
							}
							v3 = i32(1)
							t50 := m.t0[uint(v6)].(func(int32, int32, int32) int32)(v4, i32(1099059), i32(1))
							if t50 != 0 {
								goto l7
							}
							t51 := int32(load32(m.memory[int64(uint32(v0))+8:]))
							t52 := int32(load32(m.memory[int64(uint32(v0))+12:]))
							t53 := int32(load32(m.memory[uint32(v1):]))
							t54 := int32(load32(m.memory[int64(uint32(v1))+4:]))
							t55 := m.fn58(t51, t52, t53, t54)
							if t55 == 0 {
								goto l12
							}
							goto l7
						}
					l11:
						t56 := m.t0[uint(v6)].(func(int32, int32, int32) int32)(v4, i32(1099060), i32(2))
						if t56 != 0 {
							goto l7
						}
						store32(m.memory[int64(uint32(v2))+32:], uint32(v5))
						store32(m.memory[int64(uint32(v2))+28:], uint32(v4))
						v3 = i32(1)
						m.memory[int64(uint32(v2))+16] = byte(i32(1))
						store32(m.memory[int64(uint32(v2))+36:], uint32(v2+i32(16)))
						t57 := int32(load32(m.memory[int64(uint32(v0))+8:]))
						t58 := int32(load32(m.memory[int64(uint32(v0))+12:]))
						t59 := m.fn58(t57, t58, v2+i32(28), i32(1099936))
						if t59 != 0 {
							goto l7
						}
						t60 := m.fn348(v2+i32(28), i32(1099057), i32(2))
						if t60 != 0 {
							goto l7
						}
					}
				l12:
					t61 := int32(load32(m.memory[uint32(v1):]))
					t62 := int32(load32(m.memory[int64(uint32(v1))+4:]))
					t63 := int32(load32(m.memory[int64(uint32(t62))+12:]))
					t64 := m.t0[uint(t63)].(func(int32, int32, int32) int32)(t61, i32(1272328), i32(1))
					v3 = t64
					goto l7
				case 4:
					store32(m.memory[int64(uint32(v2))+12:], uint32(v0+i32(2)))
					t65 := int32(load32(m.memory[uint32(v1):]))
					t66 := int32(load32(m.memory[int64(uint32(v1))+4:]))
					t67 := int32(load32(m.memory[int64(uint32(t66))+12:]))
					t68 := m.t0[uint(t67)].(func(int32, int32, int32) int32)(t65, i32(1079597), i32(7))
					v4 = t68
					store32(m.memory[int64(uint32(v2))+16:], uint32(v1))
					v3 = i32(1)
					if v4 != 0 {
						goto l13
					}
					{
						{
							t69 := int32(m.memory[int64(uint32(v1))+10])
							if t69&i32(128) != 0 {
								goto l14
							}
							v3 = i32(1)
							t70 := int32(load32(m.memory[uint32(v1):]))
							t71 := int32(load32(m.memory[int64(uint32(v1))+4:]))
							t72 := int32(load32(m.memory[int64(uint32(t71))+12:]))
							t73 := m.t0[uint(t72)].(func(int32, int32, int32) int32)(t70, i32(1099047), i32(3))
							if t73 != 0 {
								goto l13
							}
							v3 = i32(1)
							t74 := int32(load32(m.memory[uint32(v1):]))
							t75 := int32(load32(m.memory[int64(uint32(v1))+4:]))
							t76 := int32(load32(m.memory[int64(uint32(t75))+12:]))
							t77 := m.t0[uint(t76)].(func(int32, int32, int32) int32)(t74, i32(1070584), i32(4))
							if t77 != 0 {
								goto l13
							}
							v3 = i32(1)
							t78 := int32(load32(m.memory[uint32(v1):]))
							t79 := int32(load32(m.memory[int64(uint32(v1))+4:]))
							t80 := int32(load32(m.memory[int64(uint32(t79))+12:]))
							t81 := m.t0[uint(t80)].(func(int32, int32, int32) int32)(t78, i32(1099052), i32(2))
							if t81 != 0 {
								goto l13
							}
							v3 = i32(1)
							t82 := int32(load32(m.memory[int64(uint32(v0))+4:]))
							t83 := int32(load32(m.memory[int64(uint32(v0))+8:]))
							t84 := int32(load32(m.memory[uint32(v1):]))
							t85 := int32(load32(m.memory[int64(uint32(v1))+4:]))
							t86 := m.fn58(t82, t83, t84, t85)
							if t86 == 0 {
								goto l15
							}
							goto l13
						}
					l14:
						v3 = i32(1)
						t87 := int32(load32(m.memory[uint32(v1):]))
						t88 := int32(load32(m.memory[int64(uint32(v1))+4:]))
						t89 := int32(load32(m.memory[int64(uint32(t88))+12:]))
						t90 := m.t0[uint(t89)].(func(int32, int32, int32) int32)(t87, i32(1099054), i32(3))
						if t90 != 0 {
							goto l13
						}
						v3 = i32(1)
						m.memory[int64(uint32(v2))+47] = byte(i32(1))
						t91 := int64(load64(m.memory[uint32(v1):]))
						store64(m.memory[int64(uint32(v2))+28:], uint64(t91))
						store32(m.memory[int64(uint32(v2))+36:], uint32(v2+i32(47)))
						t92 := m.fn348(v2+i32(28), i32(1070584), i32(4))
						if t92 != 0 {
							goto l13
						}
						t93 := m.fn348(v2+i32(28), i32(1099052), i32(2))
						if t93 != 0 {
							goto l13
						}
						t94 := int32(load32(m.memory[int64(uint32(v0))+4:]))
						t95 := int32(load32(m.memory[int64(uint32(v0))+8:]))
						t96 := m.fn58(t94, t95, v2+i32(28), i32(1099936))
						if t96 != 0 {
							goto l13
						}
						v3 = i32(1)
						t97 := m.fn348(v2+i32(28), i32(1099057), i32(2))
						v4 = t97
						m.memory[int64(uint32(v2))+21] = byte(i32(1))
						m.memory[int64(uint32(v2))+20] = byte(v4)
						if v4 != 0 {
							goto l13
						}
					}
				l15:
					{
						t98 := int32(m.memory[int64(uint32(v1))+10])
						if t98&i32(128) != 0 {
							t116 := int64(load64(m.memory[uint32(v1):]))
							store64(m.memory[int64(uint32(v2))+28:], uint64(t116))
							v3 = i32(1)
							m.memory[int64(uint32(v2))+47] = byte(i32(1))
							store32(m.memory[int64(uint32(v2))+36:], uint32(v2+i32(47)))
							t117 := m.fn348(v2+i32(28), i32(1079412), i32(8))
							if t117 == 0 {
								goto l17
							}
							goto l13
						}
						v3 = i32(1)
						t99 := int32(load32(m.memory[uint32(v1):]))
						t100 := int32(load32(m.memory[int64(uint32(v1))+4:]))
						t101 := int32(load32(m.memory[int64(uint32(t100))+12:]))
						t102 := m.t0[uint(t101)].(func(int32, int32, int32) int32)(t99, i32(1099050), i32(2))
						if t102 != 0 {
							goto l13
						}
						v3 = i32(1)
						t103 := int32(load32(m.memory[uint32(v1):]))
						t104 := int32(load32(m.memory[int64(uint32(v1))+4:]))
						t105 := int32(load32(m.memory[int64(uint32(t104))+12:]))
						t106 := m.t0[uint(t105)].(func(int32, int32, int32) int32)(t103, i32(1079412), i32(8))
						if t106 != 0 {
							goto l13
						}
						v3 = i32(1)
						t107 := int32(load32(m.memory[uint32(v1):]))
						t108 := int32(load32(m.memory[int64(uint32(v1))+4:]))
						t109 := int32(load32(m.memory[int64(uint32(t108))+12:]))
						t110 := m.t0[uint(t109)].(func(int32, int32, int32) int32)(t107, i32(1099052), i32(2))
						if t110 != 0 {
							goto l13
						}
						t111 := int32(load32(m.memory[int64(uint32(v0))+12:]))
						t112 := int32(load32(m.memory[int64(uint32(v0))+16:]))
						t113 := int32(load32(m.memory[uint32(v1):]))
						t114 := int32(load32(m.memory[int64(uint32(v1))+4:]))
						t115 := m.fn58(t111, t112, t113, t114)
						v3 = t115
						goto l13
					}
				case 5:
					store32(m.memory[int64(uint32(v2))+12:], uint32(v0+i32(2)))
					v3 = i32(1)
					t118 := int32(load32(m.memory[uint32(v1):]))
					v0 = t118
					t119 := int32(load32(m.memory[int64(uint32(v1))+4:]))
					t120 := v0
					v6 = t119
					t121 := int32(load32(m.memory[int64(uint32(v6))+12:]))
					v4 = t121
					t122 := m.t0[uint(v4)].(func(int32, int32, int32) int32)(t120, i32(1079604), i32(16))
					if t122 != 0 {
						goto l7
					}
					{
						{
							t123 := int32(m.memory[int64(uint32(v1))+10])
							if t123&i32(128) != 0 {
								goto l18
							}
							v3 = i32(1)
							t124 := m.t0[uint(v4)].(func(int32, int32, int32) int32)(v0, i32(1099059), i32(1))
							if t124 != 0 {
								goto l7
							}
							t125 := m.fn691(v2+i32(12), v1)
							if t125 == 0 {
								goto l19
							}
							goto l7
						}
					l18:
						t126 := m.t0[uint(v4)].(func(int32, int32, int32) int32)(v0, i32(1099060), i32(2))
						if t126 != 0 {
							goto l7
						}
						v3 = i32(1)
						m.memory[int64(uint32(v2))+47] = byte(i32(1))
						store32(m.memory[int64(uint32(v2))+20:], uint32(v6))
						store32(m.memory[int64(uint32(v2))+16:], uint32(v0))
						store32(m.memory[int64(uint32(v2))+32:], uint32(i32(1099936)))
						t127 := int64(load64(m.memory[int64(uint32(v1))+8:]))
						store64(m.memory[int64(uint32(v2))+36:], uint64(t127))
						store32(m.memory[int64(uint32(v2))+24:], uint32(v2+i32(47)))
						store32(m.memory[int64(uint32(v2))+28:], uint32(v2+i32(16)))
						t128 := m.fn691(v2+i32(12), v2+i32(28))
						if t128 != 0 {
							goto l7
						}
						t129 := int32(load32(m.memory[int64(uint32(v2))+28:]))
						t130 := int32(load32(m.memory[int64(uint32(v2))+32:]))
						t131 := int32(load32(m.memory[int64(uint32(t130))+12:]))
						t132 := m.t0[uint(t131)].(func(int32, int32, int32) int32)(t129, i32(1099057), i32(2))
						if t132 != 0 {
							goto l7
						}
					}
				l19:
					t133 := int32(load32(m.memory[uint32(v1):]))
					t134 := int32(load32(m.memory[int64(uint32(v1))+4:]))
					t135 := int32(load32(m.memory[int64(uint32(t134))+12:]))
					t136 := m.t0[uint(t135)].(func(int32, int32, int32) int32)(t133, i32(1272328), i32(1))
					v3 = t136
					goto l7
				case 6:
					v3 = i32(1)
					t137 := int32(load32(m.memory[uint32(v1):]))
					v4 = t137
					t138 := int32(load32(m.memory[int64(uint32(v1))+4:]))
					t139 := v4
					v5 = t138
					t140 := int32(load32(m.memory[int64(uint32(v5))+12:]))
					v6 = t140
					t141 := m.t0[uint(v6)].(func(int32, int32, int32) int32)(t139, i32(1079620), i32(15))
					if t141 != 0 {
						goto l7
					}
					v0 = v0 + i32(4)
					{
						{
							t142 := int32(m.memory[int64(uint32(v1))+10])
							if t142&i32(128) != 0 {
								goto l20
							}
							v3 = i32(1)
							t143 := m.t0[uint(v6)].(func(int32, int32, int32) int32)(v4, i32(1099059), i32(1))
							if t143 != 0 {
								goto l7
							}
							t144 := m.fn694(v0, v1)
							if t144 == 0 {
								goto l21
							}
							goto l7
						}
					l20:
						t145 := m.t0[uint(v6)].(func(int32, int32, int32) int32)(v4, i32(1099060), i32(2))
						if t145 != 0 {
							goto l7
						}
						v3 = i32(1)
						m.memory[int64(uint32(v2))+12] = byte(i32(1))
						store32(m.memory[int64(uint32(v2))+20:], uint32(v5))
						store32(m.memory[int64(uint32(v2))+16:], uint32(v4))
						store32(m.memory[int64(uint32(v2))+32:], uint32(i32(1099936)))
						t146 := int64(load64(m.memory[int64(uint32(v1))+8:]))
						store64(m.memory[int64(uint32(v2))+36:], uint64(t146))
						store32(m.memory[int64(uint32(v2))+24:], uint32(v2+i32(12)))
						store32(m.memory[int64(uint32(v2))+28:], uint32(v2+i32(16)))
						t147 := m.fn694(v0, v2+i32(28))
						if t147 != 0 {
							goto l7
						}
						t148 := int32(load32(m.memory[int64(uint32(v2))+28:]))
						t149 := int32(load32(m.memory[int64(uint32(v2))+32:]))
						t150 := int32(load32(m.memory[int64(uint32(t149))+12:]))
						t151 := m.t0[uint(t150)].(func(int32, int32, int32) int32)(t148, i32(1099057), i32(2))
						if t151 != 0 {
							goto l7
						}
					}
				l21:
					t152 := int32(load32(m.memory[uint32(v1):]))
					t153 := int32(load32(m.memory[int64(uint32(v1))+4:]))
					t154 := int32(load32(m.memory[int64(uint32(t153))+12:]))
					t155 := m.t0[uint(t154)].(func(int32, int32, int32) int32)(t152, i32(1272328), i32(1))
					v3 = t155
					goto l7
				}
			}
		l17:
			t156 := m.fn348(v2+i32(28), i32(1099052), i32(2))
			if t156 != 0 {
				goto l13
			}
			{
				t157 := int32(load32(m.memory[int64(uint32(v0))+12:]))
				t158 := int32(load32(m.memory[int64(uint32(v0))+16:]))
				t159 := m.fn58(t157, t158, v2+i32(28), i32(1099936))
				if t159 == 0 {
					goto l22
				}
				v3 = i32(1)
				goto l13
			}
		l22:
			t160 := m.fn348(v2+i32(28), i32(1099057), i32(2))
			v3 = t160
		}
	l13:
		m.memory[int64(uint32(v2))+20] = byte(v3)
		m.memory[int64(uint32(v2))+21] = byte(i32(1))
		t161 := m.fn350(v2+i32(16), i32(1079420), i32(5), v2+i32(12), i32(84))
		v4 = t161
		t162 := int32(m.memory[int64(uint32(v2))+21])
		v1 = t162
		t163 := int32(m.memory[int64(uint32(v2))+20])
		t164 := v1
		v0 = t163
		v3 = t164 | v0
		if v1 != i32(1) {
			goto l7
		}
		if v0&i32(1) != 0 {
			goto l7
		}
		{
			t165 := int32(load32(m.memory[uint32(v4):]))
			v1 = t165
			t166 := int32(m.memory[int64(uint32(v1))+10])
			if t166&i32(128) != 0 {
				goto l23
			}
			t167 := int32(load32(m.memory[uint32(v1):]))
			t168 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t169 := int32(load32(m.memory[int64(uint32(t168))+12:]))
			t170 := m.t0[uint(t169)].(func(int32, int32, int32) int32)(t167, i32(1273624), i32(2))
			v3 = t170
			goto l7
		}
	l23:
		t171 := int32(load32(m.memory[uint32(v1):]))
		t172 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t173 := int32(load32(m.memory[int64(uint32(t172))+12:]))
		t174 := m.t0[uint(t173)].(func(int32, int32, int32) int32)(t171, i32(1099063), i32(1))
		v3 = t174
	}
l7:
	m.g0 = v2 + i32(48)
	return v3 & i32(1)
}
func (m *Module) fn688(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6 int32
	t0 := m.g0
	v2 = t0 - i32(48)
	m.g0 = v2
	{
		t1 := int32(m.memory[uint32(v0)])
		v3 = t1
		p2 := i32(0)
		if uint32(v3) > uint32(i32(6)) {
			p2 = v3 + i32(-6)
		}
		switch p2 {
		default:
			v3 = i32(1)
			t3 := int32(load32(m.memory[uint32(v1):]))
			v4 = t3
			t4 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t5 := v4
			v5 = t4
			t6 := int32(load32(m.memory[int64(uint32(v5))+12:]))
			v6 = t6
			t7 := m.t0[uint(v6)].(func(int32, int32, int32) int32)(t5, i32(1079365), i32(3))
			if t7 != 0 {
				goto l6
			}
			{
				{
					t8 := int32(m.memory[int64(uint32(v1))+10])
					if t8&i32(128) != 0 {
						goto l7
					}
					v3 = i32(1)
					t9 := m.t0[uint(v6)].(func(int32, int32, int32) int32)(v4, i32(1099059), i32(1))
					if t9 != 0 {
						goto l6
					}
					t10 := m.fn687(v0, v1)
					if t10 == 0 {
						goto l8
					}
					goto l6
				}
			l7:
				t11 := m.t0[uint(v6)].(func(int32, int32, int32) int32)(v4, i32(1099060), i32(2))
				if t11 != 0 {
					goto l6
				}
				v3 = i32(1)
				m.memory[int64(uint32(v2))+12] = byte(i32(1))
				store32(m.memory[int64(uint32(v2))+20:], uint32(v5))
				store32(m.memory[int64(uint32(v2))+16:], uint32(v4))
				store32(m.memory[int64(uint32(v2))+32:], uint32(i32(1099936)))
				t12 := int64(load64(m.memory[int64(uint32(v1))+8:]))
				store64(m.memory[int64(uint32(v2))+36:], uint64(t12))
				store32(m.memory[int64(uint32(v2))+24:], uint32(v2+i32(12)))
				store32(m.memory[int64(uint32(v2))+28:], uint32(v2+i32(16)))
				t13 := m.fn687(v0, v2+i32(28))
				if t13 != 0 {
					goto l6
				}
				t14 := int32(load32(m.memory[int64(uint32(v2))+28:]))
				t15 := int32(load32(m.memory[int64(uint32(v2))+32:]))
				t16 := int32(load32(m.memory[int64(uint32(t15))+12:]))
				t17 := m.t0[uint(t16)].(func(int32, int32, int32) int32)(t14, i32(1099057), i32(2))
				if t17 != 0 {
					goto l6
				}
			}
		l8:
			t18 := int32(load32(m.memory[uint32(v1):]))
			t19 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t20 := int32(load32(m.memory[int64(uint32(t19))+12:]))
			t21 := m.t0[uint(t20)].(func(int32, int32, int32) int32)(t18, i32(1272328), i32(1))
			v3 = t21
			goto l6
		case 1:
			v3 = i32(1)
			t22 := int32(load32(m.memory[uint32(v1):]))
			v4 = t22
			t23 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t24 := v4
			v5 = t23
			t25 := int32(load32(m.memory[int64(uint32(v5))+12:]))
			v6 = t25
			t26 := m.t0[uint(v6)].(func(int32, int32, int32) int32)(t24, i32(1091476), i32(2))
			if t26 != 0 {
				goto l6
			}
			v0 = v0 + i32(4)
			{
				{
					t27 := int32(m.memory[int64(uint32(v1))+10])
					if t27&i32(128) != 0 {
						goto l9
					}
					v3 = i32(1)
					t28 := m.t0[uint(v6)].(func(int32, int32, int32) int32)(v4, i32(1099059), i32(1))
					if t28 != 0 {
						goto l6
					}
					t29 := m.fn351(v0, v1)
					if t29 == 0 {
						goto l10
					}
					goto l6
				}
			l9:
				t30 := m.t0[uint(v6)].(func(int32, int32, int32) int32)(v4, i32(1099060), i32(2))
				if t30 != 0 {
					goto l6
				}
				v3 = i32(1)
				m.memory[int64(uint32(v2))+12] = byte(i32(1))
				store32(m.memory[int64(uint32(v2))+20:], uint32(v5))
				store32(m.memory[int64(uint32(v2))+16:], uint32(v4))
				store32(m.memory[int64(uint32(v2))+32:], uint32(i32(1099936)))
				t31 := int64(load64(m.memory[int64(uint32(v1))+8:]))
				store64(m.memory[int64(uint32(v2))+36:], uint64(t31))
				store32(m.memory[int64(uint32(v2))+24:], uint32(v2+i32(12)))
				store32(m.memory[int64(uint32(v2))+28:], uint32(v2+i32(16)))
				t32 := m.fn351(v0, v2+i32(28))
				if t32 != 0 {
					goto l6
				}
				t33 := int32(load32(m.memory[int64(uint32(v2))+28:]))
				t34 := int32(load32(m.memory[int64(uint32(v2))+32:]))
				t35 := int32(load32(m.memory[int64(uint32(t34))+12:]))
				t36 := m.t0[uint(t35)].(func(int32, int32, int32) int32)(t33, i32(1099057), i32(2))
				if t36 != 0 {
					goto l6
				}
			}
		l10:
			t37 := int32(load32(m.memory[uint32(v1):]))
			t38 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t39 := int32(load32(m.memory[int64(uint32(t38))+12:]))
			t40 := m.t0[uint(t39)].(func(int32, int32, int32) int32)(t37, i32(1272328), i32(1))
			v3 = t40
			goto l6
		case 2:
			v3 = i32(1)
			t41 := int32(load32(m.memory[uint32(v1):]))
			v4 = t41
			t42 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t43 := v4
			v5 = t42
			t44 := int32(load32(m.memory[int64(uint32(v5))+12:]))
			v6 = t44
			t45 := m.t0[uint(v6)].(func(int32, int32, int32) int32)(t43, i32(1079368), i32(14))
			if t45 != 0 {
				goto l6
			}
			{
				{
					t46 := int32(m.memory[int64(uint32(v1))+10])
					if t46&i32(128) != 0 {
						goto l11
					}
					v3 = i32(1)
					t47 := m.t0[uint(v6)].(func(int32, int32, int32) int32)(v4, i32(1099059), i32(1))
					if t47 != 0 {
						goto l6
					}
					t48 := int32(load32(m.memory[int64(uint32(v0))+8:]))
					t49 := int32(load32(m.memory[int64(uint32(v0))+12:]))
					t50 := int32(load32(m.memory[uint32(v1):]))
					t51 := int32(load32(m.memory[int64(uint32(v1))+4:]))
					t52 := m.fn58(t48, t49, t50, t51)
					if t52 == 0 {
						goto l12
					}
					goto l6
				}
			l11:
				t53 := m.t0[uint(v6)].(func(int32, int32, int32) int32)(v4, i32(1099060), i32(2))
				if t53 != 0 {
					goto l6
				}
				store32(m.memory[int64(uint32(v2))+32:], uint32(v5))
				store32(m.memory[int64(uint32(v2))+28:], uint32(v4))
				v3 = i32(1)
				m.memory[int64(uint32(v2))+16] = byte(i32(1))
				store32(m.memory[int64(uint32(v2))+36:], uint32(v2+i32(16)))
				t54 := int32(load32(m.memory[int64(uint32(v0))+8:]))
				t55 := int32(load32(m.memory[int64(uint32(v0))+12:]))
				t56 := m.fn58(t54, t55, v2+i32(28), i32(1099936))
				if t56 != 0 {
					goto l6
				}
				t57 := m.fn348(v2+i32(28), i32(1099057), i32(2))
				if t57 != 0 {
					goto l6
				}
			}
		l12:
			t58 := int32(load32(m.memory[uint32(v1):]))
			t59 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t60 := int32(load32(m.memory[int64(uint32(t59))+12:]))
			t61 := m.t0[uint(t60)].(func(int32, int32, int32) int32)(t58, i32(1272328), i32(1))
			v3 = t61
			goto l6
		case 3:
			store32(m.memory[int64(uint32(v2))+12:], uint32(v0+i32(12)))
			t62 := int32(load32(m.memory[uint32(v1):]))
			t63 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t64 := int32(load32(m.memory[int64(uint32(t63))+12:]))
			t65 := m.t0[uint(t64)].(func(int32, int32, int32) int32)(t62, i32(1079382), i32(7))
			v4 = t65
			store32(m.memory[int64(uint32(v2))+16:], uint32(v1))
			v3 = i32(1)
			{
				if v4 != 0 {
					goto l13
				}
				{
					t66 := int32(m.memory[int64(uint32(v1))+10])
					if t66&i32(128) != 0 {
						goto l14
					}
					v3 = i32(1)
					t67 := int32(load32(m.memory[uint32(v1):]))
					t68 := int32(load32(m.memory[int64(uint32(v1))+4:]))
					t69 := int32(load32(m.memory[int64(uint32(t68))+12:]))
					t70 := m.t0[uint(t69)].(func(int32, int32, int32) int32)(t67, i32(1099047), i32(3))
					if t70 != 0 {
						goto l13
					}
					v3 = i32(1)
					t71 := int32(load32(m.memory[uint32(v1):]))
					t72 := int32(load32(m.memory[int64(uint32(v1))+4:]))
					t73 := int32(load32(m.memory[int64(uint32(t72))+12:]))
					t74 := m.t0[uint(t73)].(func(int32, int32, int32) int32)(t71, i32(1079389), i32(3))
					if t74 != 0 {
						goto l13
					}
					v3 = i32(1)
					t75 := int32(load32(m.memory[uint32(v1):]))
					t76 := int32(load32(m.memory[int64(uint32(v1))+4:]))
					t77 := int32(load32(m.memory[int64(uint32(t76))+12:]))
					t78 := m.t0[uint(t77)].(func(int32, int32, int32) int32)(t75, i32(1099052), i32(2))
					if t78 != 0 {
						goto l13
					}
					t79 := int32(load32(m.memory[int64(uint32(v0))+4:]))
					t80 := int32(load32(m.memory[int64(uint32(v0))+8:]))
					t81 := int32(load32(m.memory[uint32(v1):]))
					t82 := int32(load32(m.memory[int64(uint32(v1))+4:]))
					t83 := m.fn58(t79, t80, t81, t82)
					v3 = t83
					goto l13
				}
			l14:
				v3 = i32(1)
				t84 := int32(load32(m.memory[uint32(v1):]))
				t85 := int32(load32(m.memory[int64(uint32(v1))+4:]))
				t86 := int32(load32(m.memory[int64(uint32(t85))+12:]))
				t87 := m.t0[uint(t86)].(func(int32, int32, int32) int32)(t84, i32(1099054), i32(3))
				if t87 != 0 {
					goto l13
				}
				v3 = i32(1)
				m.memory[int64(uint32(v2))+47] = byte(i32(1))
				t88 := int64(load64(m.memory[uint32(v1):]))
				store64(m.memory[int64(uint32(v2))+28:], uint64(t88))
				store32(m.memory[int64(uint32(v2))+36:], uint32(v2+i32(47)))
				t89 := m.fn348(v2+i32(28), i32(1079389), i32(3))
				if t89 != 0 {
					goto l13
				}
				t90 := m.fn348(v2+i32(28), i32(1099052), i32(2))
				if t90 != 0 {
					goto l13
				}
				{
					t91 := int32(load32(m.memory[int64(uint32(v0))+4:]))
					t92 := int32(load32(m.memory[int64(uint32(v0))+8:]))
					t93 := m.fn58(t91, t92, v2+i32(28), i32(1099936))
					if t93 == 0 {
						goto l15
					}
					v3 = i32(1)
					goto l13
				}
			l15:
				t94 := m.fn348(v2+i32(28), i32(1099057), i32(2))
				v3 = t94
			}
		l13:
			m.memory[int64(uint32(v2))+20] = byte(v3)
			m.memory[int64(uint32(v2))+21] = byte(i32(1))
			t95 := m.fn350(v2+i32(16), i32(1069495), i32(3), v2+i32(12), i32(84))
			v4 = t95
			t96 := int32(m.memory[int64(uint32(v2))+21])
			v1 = t96
			t97 := int32(m.memory[int64(uint32(v2))+20])
			t98 := v1
			v0 = t97
			v3 = t98 | v0
			if v1 != i32(1) {
				goto l6
			}
			if v0&i32(1) != 0 {
				goto l6
			}
			{
				t99 := int32(load32(m.memory[uint32(v4):]))
				v1 = t99
				t100 := int32(m.memory[int64(uint32(v1))+10])
				if t100&i32(128) != 0 {
					t105 := int32(load32(m.memory[uint32(v1):]))
					t106 := int32(load32(m.memory[int64(uint32(v1))+4:]))
					t107 := int32(load32(m.memory[int64(uint32(t106))+12:]))
					t108 := m.t0[uint(t107)].(func(int32, int32, int32) int32)(t105, i32(1099063), i32(1))
					v3 = t108
					goto l6
				}
				t101 := int32(load32(m.memory[uint32(v1):]))
				t102 := int32(load32(m.memory[int64(uint32(v1))+4:]))
				t103 := int32(load32(m.memory[int64(uint32(t102))+12:]))
				t104 := m.t0[uint(t103)].(func(int32, int32, int32) int32)(t101, i32(1273624), i32(2))
				v3 = t104
				goto l6
			}
		case 4:
			t109 := int32(load32(m.memory[uint32(v1):]))
			t110 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t111 := int32(load32(m.memory[int64(uint32(t110))+12:]))
			t112 := m.t0[uint(t111)].(func(int32, int32, int32) int32)(t109, i32(1079392), i32(5))
			v3 = t112
			goto l6
		case 5:
			store32(m.memory[int64(uint32(v2))+16:], uint32(v0+i32(4)))
			t113 := int32(load32(m.memory[uint32(v1):]))
			t114 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t115 := int32(load32(m.memory[int64(uint32(t114))+12:]))
			t116 := m.t0[uint(t115)].(func(int32, int32, int32) int32)(t113, i32(1079397), i32(15))
			v3 = t116
			m.memory[int64(uint32(v2))+33] = byte(i32(0))
			m.memory[int64(uint32(v2))+32] = byte(v3)
			store32(m.memory[int64(uint32(v2))+28:], uint32(v1))
			t117 := m.fn350(v2+i32(28), i32(1079412), i32(8), v0+i32(2), i32(87))
			t118 := m.fn350(t117, i32(1079420), i32(5), v2+i32(16), i32(84))
			v4 = t118
			t119 := int32(m.memory[int64(uint32(v2))+33])
			v1 = t119
			t120 := int32(m.memory[int64(uint32(v2))+32])
			t121 := v1
			v0 = t120
			v3 = t121 | v0
			if v1 != i32(1) {
				goto l6
			}
			if v0&i32(1) != 0 {
				goto l6
			}
			{
				t122 := int32(load32(m.memory[uint32(v4):]))
				v1 = t122
				t123 := int32(m.memory[int64(uint32(v1))+10])
				if t123&i32(128) != 0 {
					goto l17
				}
				t124 := int32(load32(m.memory[uint32(v1):]))
				t125 := int32(load32(m.memory[int64(uint32(v1))+4:]))
				t126 := int32(load32(m.memory[int64(uint32(t125))+12:]))
				t127 := m.t0[uint(t126)].(func(int32, int32, int32) int32)(t124, i32(1273624), i32(2))
				v3 = t127
				goto l6
			}
		l17:
			t128 := int32(load32(m.memory[uint32(v1):]))
			t129 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t130 := int32(load32(m.memory[int64(uint32(t129))+12:]))
			t131 := m.t0[uint(t130)].(func(int32, int32, int32) int32)(t128, i32(1099063), i32(1))
			v3 = t131
		}
	}
l6:
	m.g0 = v2 + i32(48)
	return v3 & i32(1)
}
func (m *Module) fn689(v0, v1 int32) int32 {
	var v2, v3 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	{
		{
			t1 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			v3 = t1
			if v3&i32(0x2000000) != 0 {
				t3 := int32(load32(m.memory[uint32(v0):]))
				v3 = t3
				v0 = i32(9)
			l3:
				{
					t4 := int32(m.memory[int64(uint32(v3&i32(15)))+1098832])
					m.memory[uint32(v2+i32(8)+v0+i32(-2))] = byte(t4)
					v0 = v0 + i32(-1)
					v3 = int32(uint32(v3) >> 4)
					if v3 != 0 {
						goto l3
					}
				}
				t5 := m.fn312(v1, i32(1), i32(1122566), i32(2), v2+i32(8)+v0+i32(-1), i32(9)-v0)
				v0 = t5
				goto l2
			}
			if v3&i32(0x4000000) != 0 {
				goto l1
			}
			t2 := m.fn24(v0, v1)
			v0 = t2
			goto l2
		}
	l1:
		t6 := int32(load32(m.memory[uint32(v0):]))
		v3 = t6
		v0 = i32(9)
	l4:
		{
			t7 := int32(m.memory[int64(uint32(v3&i32(15)))+1122568])
			m.memory[uint32(v2+i32(8)+v0+i32(-2))] = byte(t7)
			v0 = v0 + i32(-1)
			v3 = int32(uint32(v3) >> 4)
			if v3 != 0 {
				goto l4
			}
		}
		t8 := m.fn312(v1, i32(1), i32(1122566), i32(2), v2+i32(8)+v0+i32(-1), i32(9)-v0)
		v0 = t8
	}
l2:
	m.g0 = v2 + i32(16)
	return v0
}
func (m *Module) fn690(v0, v1 int32) int32 {
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
					t5 := int32(m.memory[int64(uint32(v3&i32(15)))+1098832])
					m.memory[uint32(v2+i32(8)+v0+i32(-2))] = byte(t5)
					v0 = v0 + i32(-1)
					v3 = int32(uint32(v3) >> 4)
					if v3 != 0 {
						goto l3
					}
				}
				t6 := m.fn312(v1, i32(1), i32(1122566), i32(2), v2+i32(8)+v0+i32(-1), i32(9)-v0)
				v0 = t6
				goto l2
			}
			if v3&i32(0x4000000) != 0 {
				goto l1
			}
			t3 := m.fn24(v0, v1)
			v0 = t3
			goto l2
		}
	l1:
		t7 := int32(load32(m.memory[uint32(v0):]))
		v3 = t7
		v0 = i32(9)
	l4:
		{
			t8 := int32(m.memory[int64(uint32(v3&i32(15)))+1122568])
			m.memory[uint32(v2+i32(8)+v0+i32(-2))] = byte(t8)
			v0 = v0 + i32(-1)
			v3 = int32(uint32(v3) >> 4)
			if v3 != 0 {
				goto l4
			}
		}
		t9 := m.fn312(v1, i32(1), i32(1122566), i32(2), v2+i32(8)+v0+i32(-1), i32(9)-v0)
		v0 = t9
	}
l2:
	m.g0 = v2 + i32(16)
	return v0
}
func (m *Module) fn691(v0, v1 int32) int32 {
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
				t4 := int32(load16(m.memory[uint32(v0):]))
				v3 = t4
				v0 = i32(5)
			l3:
				{
					t5 := int32(m.memory[uint32(v3&i32(15)+i32(1098832))])
					m.memory[uint32(v2+i32(8)+v0+i32(-2))] = byte(t5)
					v0 = v0 + i32(-1)
					v3 = int32(uint32(v3)>>4) & i32(0xfff)
					if v3 != 0 {
						goto l3
					}
				}
				t6 := m.fn312(v1, i32(1), i32(1122566), i32(2), v2+i32(8)+v0+i32(-1), i32(5)-v0)
				v0 = t6
				goto l2
			}
			if v3&i32(0x4000000) != 0 {
				goto l1
			}
			t3 := m.fn172(v0, v1)
			v0 = t3
			goto l2
		}
	l1:
		t7 := int32(load16(m.memory[uint32(v0):]))
		v3 = t7
		v0 = i32(5)
	l4:
		{
			t8 := int32(m.memory[uint32(v3&i32(15)+i32(1122568))])
			m.memory[uint32(v2+i32(12)+v0+i32(-2))] = byte(t8)
			v0 = v0 + i32(-1)
			v3 = int32(uint32(v3)>>4) & i32(0xfff)
			if v3 != 0 {
				goto l4
			}
		}
		t9 := m.fn312(v1, i32(1), i32(1122566), i32(2), v2+i32(12)+v0+i32(-1), i32(5)-v0)
		v0 = t9
	}
l2:
	m.g0 = v2 + i32(16)
	return v0
}
func (m *Module) fn692(v0, v1 int32) int32 {
	var v2, v3 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	{
		{
			t1 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			v3 = t1
			if v3&i32(0x2000000) != 0 {
				t3 := int32(load16(m.memory[uint32(v0):]))
				v3 = t3
				v0 = i32(5)
			l3:
				{
					t4 := int32(m.memory[uint32(v3&i32(15)+i32(1098832))])
					m.memory[uint32(v2+i32(8)+v0+i32(-2))] = byte(t4)
					v0 = v0 + i32(-1)
					v3 = int32(uint32(v3)>>4) & i32(0xfff)
					if v3 != 0 {
						goto l3
					}
				}
				t5 := m.fn312(v1, i32(1), i32(1122566), i32(2), v2+i32(8)+v0+i32(-1), i32(5)-v0)
				v0 = t5
				goto l2
			}
			if v3&i32(0x4000000) != 0 {
				goto l1
			}
			t2 := m.fn172(v0, v1)
			v0 = t2
			goto l2
		}
	l1:
		t6 := int32(load16(m.memory[uint32(v0):]))
		v3 = t6
		v0 = i32(5)
	l4:
		{
			t7 := int32(m.memory[uint32(v3&i32(15)+i32(1122568))])
			m.memory[uint32(v2+i32(12)+v0+i32(-2))] = byte(t7)
			v0 = v0 + i32(-1)
			v3 = int32(uint32(v3)>>4) & i32(0xfff)
			if v3 != 0 {
				goto l4
			}
		}
		t8 := m.fn312(v1, i32(1), i32(1122566), i32(2), v2+i32(12)+v0+i32(-1), i32(5)-v0)
		v0 = t8
	}
l2:
	m.g0 = v2 + i32(16)
	return v0
}
func (m *Module) fn693(v0, v1 int32) int32 {
	var v2, v3, v4, v5, v6 int32
	t0 := m.g0
	v2 = t0 - i32(32)
	m.g0 = v2
	v3 = i32(1)
	{
		{
			t1 := int32(load32(m.memory[uint32(v0):]))
			v0 = t1
			t2 := int32(m.memory[uint32(v0)])
			if t2 != i32(1) {
				goto l0
			}
			t3 := int32(load32(m.memory[uint32(v1):]))
			v4 = t3
			t4 := int32(load32(m.memory[int64(uint32(v1))+4:]))
			t5 := v4
			v5 = t4
			t6 := int32(load32(m.memory[int64(uint32(v5))+12:]))
			v6 = t6
			t7 := m.t0[uint(v6)].(func(int32, int32, int32) int32)(t5, i32(1079880), i32(4))
			if t7 != 0 {
				goto l1
			}
			v0 = v0 + i32(1)
			{
				{
					t8 := int32(m.memory[int64(uint32(v1))+10])
					if t8&i32(128) != 0 {
						goto l2
					}
					v3 = i32(1)
					t9 := m.t0[uint(v6)].(func(int32, int32, int32) int32)(v4, i32(1099059), i32(1))
					if t9 != 0 {
						goto l1
					}
					t10 := m.fn170(v0, v1)
					if t10 != 0 {
						goto l1
					}
					t11 := int32(load32(m.memory[uint32(v1):]))
					v4 = t11
					t12 := int32(load32(m.memory[int64(uint32(v1))+4:]))
					t13 := int32(load32(m.memory[int64(uint32(t12))+12:]))
					v6 = t13
					goto l3
				}
			l2:
				t14 := m.t0[uint(v6)].(func(int32, int32, int32) int32)(v4, i32(1099060), i32(2))
				if t14 != 0 {
					goto l1
				}
				v3 = i32(1)
				m.memory[int64(uint32(v2))+15] = byte(i32(1))
				store32(m.memory[int64(uint32(v2))+4:], uint32(v5))
				store32(m.memory[uint32(v2):], uint32(v4))
				store32(m.memory[int64(uint32(v2))+20:], uint32(i32(1099936)))
				t15 := int64(load64(m.memory[int64(uint32(v1))+8:]))
				store64(m.memory[int64(uint32(v2))+24:], uint64(t15))
				store32(m.memory[int64(uint32(v2))+8:], uint32(v2+i32(15)))
				store32(m.memory[int64(uint32(v2))+16:], uint32(v2))
				t16 := m.fn170(v0, v2+i32(16))
				if t16 != 0 {
					goto l1
				}
				t17 := int32(load32(m.memory[int64(uint32(v2))+16:]))
				t18 := int32(load32(m.memory[int64(uint32(v2))+20:]))
				t19 := int32(load32(m.memory[int64(uint32(t18))+12:]))
				t20 := m.t0[uint(t19)].(func(int32, int32, int32) int32)(t17, i32(1099057), i32(2))
				if t20 != 0 {
					goto l1
				}
			}
		l3:
			t21 := m.t0[uint(v6)].(func(int32, int32, int32) int32)(v4, i32(1272328), i32(1))
			v3 = t21
			goto l1
		}
	l0:
		t22 := int32(load32(m.memory[uint32(v1):]))
		t23 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t24 := int32(load32(m.memory[int64(uint32(t23))+12:]))
		t25 := m.t0[uint(t24)].(func(int32, int32, int32) int32)(t22, i32(1079876), i32(4))
		v3 = t25
	}
l1:
	m.g0 = v2 + i32(32)
	return v3
}
func (m *Module) fn694(v0, v1 int32) int32 {
	var v2, v3 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	{
		{
			t1 := int32(load32(m.memory[int64(uint32(v1))+8:]))
			v3 = t1
			if v3&i32(0x2000000) != 0 {
				t3 := int32(load32(m.memory[uint32(v0):]))
				v3 = t3
				v0 = i32(9)
			l3:
				{
					t4 := int32(m.memory[int64(uint32(v3&i32(15)))+1098832])
					m.memory[uint32(v2+i32(8)+v0+i32(-2))] = byte(t4)
					v0 = v0 + i32(-1)
					v3 = int32(uint32(v3) >> 4)
					if v3 != 0 {
						goto l3
					}
				}
				t5 := m.fn312(v1, i32(1), i32(1122566), i32(2), v2+i32(8)+v0+i32(-1), i32(9)-v0)
				v0 = t5
				goto l2
			}
			if v3&i32(0x4000000) != 0 {
				goto l1
			}
			t2 := m.fn24(v0, v1)
			v0 = t2
			goto l2
		}
	l1:
		t6 := int32(load32(m.memory[uint32(v0):]))
		v3 = t6
		v0 = i32(9)
	l4:
		{
			t7 := int32(m.memory[int64(uint32(v3&i32(15)))+1122568])
			m.memory[uint32(v2+i32(8)+v0+i32(-2))] = byte(t7)
			v0 = v0 + i32(-1)
			v3 = int32(uint32(v3) >> 4)
			if v3 != 0 {
				goto l4
			}
		}
		t8 := m.fn312(v1, i32(1), i32(1122566), i32(2), v2+i32(8)+v0+i32(-1), i32(9)-v0)
		v0 = t8
	}
l2:
	m.g0 = v2 + i32(16)
	return v0
}
func (m *Module) fn695(v0 int32) {
	m.fn34(i32(1102200), i32(51), v0)
	panic("unreachable")
}
func (m *Module) fn696(v0, v1, v2 int32) {
	var v3, v4, v5, v6, v7, v8, v9, v10, v11, v12 int32
	var v13 int64
	t0 := m.g0
	v3 = t0 - i32(160)
	m.g0 = v3
	store16(m.memory[int64(uint32(v3))+40:], uint16(i32(514)))
	store64(m.memory[int64(uint32(v3))+32:], uint64(i64(144680345676153346)))
	store16(m.memory[int64(uint32(v3))+80:], uint16(i32(1)))
	store32(m.memory[int64(uint32(v3))+76:], uint32(v2))
	store32(m.memory[int64(uint32(v3))+72:], uint32(i32(0)))
	m.memory[int64(uint32(v3))+68] = byte(i32(1))
	store32(m.memory[int64(uint32(v3))+64:], uint32(i32(59)))
	store32(m.memory[int64(uint32(v3))+60:], uint32(v2))
	store32(m.memory[int64(uint32(v3))+56:], uint32(i32(0)))
	store32(m.memory[int64(uint32(v3))+52:], uint32(v2))
	store32(m.memory[int64(uint32(v3))+48:], uint32(v1))
	store32(m.memory[int64(uint32(v3))+44:], uint32(i32(59)))
l10:
	{
		t1 := int32(load32(m.memory[int64(uint32(v3))+48:]))
		v2 = t1
		m.fn205(v3+i32(96), v3+i32(44))
		{
			{
				t2 := int32(load32(m.memory[int64(uint32(v3))+96:]))
				if t2 != i32(1) {
					goto l0
				}
				t3 := int32(load32(m.memory[int64(uint32(v3))+72:]))
				v4 = t3
				t4 := int32(load32(m.memory[int64(uint32(v3))+104:]))
				store32(m.memory[int64(uint32(v3))+72:], uint32(t4))
				v1 = v2 + v4
				t5 := int32(load32(m.memory[int64(uint32(v3))+100:]))
				v2 = t5 - v4
				goto l1
			}
		l0:
			t6 := int32(m.memory[int64(uint32(v3))+81])
			if t6 != 0 {
				goto l2
			}
			m.memory[int64(uint32(v3))+81] = byte(i32(1))
			{
				{
					t7 := int32(m.memory[int64(uint32(v3))+80])
					if t7 != i32(1) {
						goto l3
					}
					t8 := int32(load32(m.memory[int64(uint32(v3))+76:]))
					v4 = t8
					t9 := int32(load32(m.memory[int64(uint32(v3))+72:]))
					v2 = t9
					goto l4
				}
			l3:
				t10 := int32(load32(m.memory[int64(uint32(v3))+76:]))
				v4 = t10
				t11 := int32(load32(m.memory[int64(uint32(v3))+72:]))
				t12 := v4
				v2 = t11
				if t12 == v2 {
					goto l2
				}
			}
		l4:
			t13 := int32(load32(m.memory[int64(uint32(v3))+48:]))
			v1 = t13 + v2
			v2 = v4 - v2
		}
	l1:
		store32(m.memory[int64(uint32(v3))+112:], uint32(v2))
		store32(m.memory[int64(uint32(v3))+108:], uint32(i32(0)))
		store32(m.memory[int64(uint32(v3))+104:], uint32(v2))
		store32(m.memory[int64(uint32(v3))+100:], uint32(v1))
		m.memory[int64(uint32(v3))+120] = byte(i32(1))
		store32(m.memory[int64(uint32(v3))+96:], uint32(i32(58)))
		store32(m.memory[int64(uint32(v3))+116:], uint32(i32(58)))
		m.fn205(v3+i32(84), v3+i32(96))
		{
			{
				{
					{
						t14 := int32(load32(m.memory[int64(uint32(v3))+84:]))
						if t14 == 0 {
							t20 := int32(m.memory[int64(uint32(v3))+81])
							if t20 == 0 {
								goto l10
							}
							goto l2
						}
						t15 := int32(load32(m.memory[int64(uint32(v3))+92:]))
						v5 = t15
						t16 := int32(load32(m.memory[int64(uint32(v3))+88:]))
						m.fn150(v3+i32(24), v1, t16)
						t17 := int32(load32(m.memory[int64(uint32(v3))+28:]))
						v6 = t17
						if v6 <= i32(-1) {
							goto l6
						}
						if v6 != 0 {
							t18 := int32(load32(m.memory[int64(uint32(v3))+24:]))
							v4 = t18
							t19 := m.fn11(v6)
							v7 = t19
							if v7 != 0 {
								goto l9
							}
							m.fn7(i32(1), v6)
							panic("unreachable")
						}
						v7 = i32(1)
						goto l8
					}
				l9:
					if v6 == 0 {
						goto l11
					}
					memory_copy(m.memory, uint32(v7), uint32(v4), uint32(v6))
				l11:
					v4 = i32(0)
					if v6 == i32(1) {
						goto l12
					}
					v8 = v6 & i32(1)
					v9 = v6 & i32(0x7ffffffe)
					v4 = i32(0)
				l13:
					{
						v10 = v7 + v4
						t21 := int32(m.memory[uint32(v10)])
						t22 := v10
						v11 = t21
						p23 := i32(0)
						if uint32((v11+i32(-65))&i32(255)) < uint32(i32(26)) {
							p23 = i32(32)
						}
						m.memory[uint32(t22)] = byte(p23 | v11)
						v10 = v10 + i32(1)
						t24 := int32(m.memory[uint32(v10)])
						t25 := v10
						v10 = t24
						p26 := i32(0)
						if uint32((v10+i32(-65))&i32(255)) < uint32(i32(26)) {
							p26 = i32(32)
						}
						m.memory[uint32(t25)] = byte(p26 | v10)
						t27 := v9
						v4 = v4 + i32(2)
						if t27 != v4 {
							goto l13
						}
					}
					if v8 == 0 {
						goto l8
					}
				l12:
					v4 = v7 + v4
					t28 := int32(m.memory[uint32(v4)])
					t29 := v4
					v4 = t28
					p30 := i32(0)
					if uint32((v4+i32(-65))&i32(255)) < uint32(i32(26)) {
						p30 = i32(32)
					}
					m.memory[uint32(t29)] = byte(p30 | v4)
				}
			l8:
				m.fn150(v3+i32(16), v1+v5, v2-v5)
				t31 := int32(load32(m.memory[int64(uint32(v3))+20:]))
				v11 = t31
				if v11 <= i32(-1) {
					goto l6
				}
				{
					if v11 != 0 {
						goto l14
					}
					v10 = i32(1)
					goto l15
				l14:
					t32 := int32(load32(m.memory[int64(uint32(v3))+16:]))
					v2 = t32
					t33 := m.fn11(v11)
					v10 = t33
					if v10 == 0 {
						m.fn7(i32(1), v11)
						panic("unreachable")
					}
					if v11 == 0 {
						goto l17
					}
					memory_copy(m.memory, uint32(v10), uint32(v2), uint32(v11))
				l17:
					v2 = i32(0)
					if v11 == i32(1) {
						goto l18
					}
					v5 = v11 & i32(1)
					v9 = v11 & i32(0x7ffffffe)
					v2 = i32(0)
				l19:
					{
						v1 = v10 + v2
						t34 := int32(m.memory[uint32(v1)])
						t35 := v1
						v4 = t34
						p36 := i32(0)
						if uint32((v4+i32(-65))&i32(255)) < uint32(i32(26)) {
							p36 = i32(32)
						}
						m.memory[uint32(t35)] = byte(p36 | v4)
						v1 = v1 + i32(1)
						t37 := int32(m.memory[uint32(v1)])
						t38 := v1
						v1 = t37
						p39 := i32(0)
						if uint32((v1+i32(-65))&i32(255)) < uint32(i32(26)) {
							p39 = i32(32)
						}
						m.memory[uint32(t38)] = byte(p39 | v1)
						t40 := v9
						v2 = v2 + i32(2)
						if t40 != v2 {
							goto l19
						}
					}
					if v5 == 0 {
						goto l15
					}
				l18:
					v2 = v10 + v2
					t41 := int32(m.memory[uint32(v2)])
					t42 := v2
					v2 = t41
					p43 := i32(0)
					if uint32((v2+i32(-65))&i32(255)) < uint32(i32(26)) {
						p43 = i32(32)
					}
					m.memory[uint32(t42)] = byte(p43 | v2)
				}
			l15:
				store32(m.memory[int64(uint32(v3))+112:], uint32(v11))
				store32(m.memory[int64(uint32(v3))+108:], uint32(i32(0)))
				store32(m.memory[int64(uint32(v3))+104:], uint32(v11))
				store32(m.memory[int64(uint32(v3))+100:], uint32(v10))
				store32(m.memory[int64(uint32(v3))+96:], uint32(i32(33)))
				store32(m.memory[int64(uint32(v3))+116:], uint32(i32(33)))
				m.memory[int64(uint32(v3))+120] = byte(i32(1))
				m.fn205(v3+i32(84), v3+i32(96))
				{
					t44 := int32(load32(m.memory[int64(uint32(v3))+84:]))
					if t44 != i32(1) {
						v9 = v3 + i32(32)
						v2 = v11
						goto l24
					}
					t45 := int32(load32(m.memory[int64(uint32(v3))+88:]))
					v4 = t45
					v2 = v4 + i32(1)
					if v2 == 0 {
						goto l21
					}
					{
						if uint32(v11) > uint32(v2) {
							goto l22
						}
						if v11 != v2 {
							goto l23
						}
						goto l21
					l22:
						t46 := int32(int8(m.memory[uint32(v10+v2)]))
						if t46 > i32(-65) {
							goto l21
						}
					}
				l23:
					m.fn44(v10, v11, v2, v11, i32(1075780))
					panic("unreachable")
				}
			}
		l6:
			m.fn12()
			panic("unreachable")
		l21:
			m.fn150(v3+i32(8), v10+v2, v11-v2)
			v9 = v3 + i32(32)
			{
				t47 := int32(load32(m.memory[int64(uint32(v3))+12:]))
				if t47 != i32(9) {
					goto l25
				}
				t48 := int32(load32(m.memory[int64(uint32(v3))+8:]))
				t49 := v3 + i32(32)
				v2 = t48
				t50 := int64(load64(m.memory[uint32(v2):]))
				t51 := int64(m.memory[uint32(v2+i32(8))])
				p52 := i32(0)
				if t50^i64(7953766451757739369)|(t51^i64(116)) == 0 {
					p52 = i32(5)
				}
				v9 = t49 | p52
			}
		l25:
			v1 = v11
			{
				if uint32(v4) > uint32(v11) {
					goto l26
				}
				if v4 != 0 {
					goto l27
				}
				v1 = i32(0)
				goto l26
			l27:
				if uint32(v4) < uint32(v11) {
					goto l28
				}
				v1 = v4
				goto l26
			l28:
				v1 = v4
				t53 := int32(int8(m.memory[uint32(v10+v4)]))
				if t53 <= i32(-65) {
					m.fn2(i32(1080413), i32(48), i32(1075796))
					panic("unreachable")
				}
			}
		l26:
			m.fn705(v3, v10, v1)
			{
				t54 := int32(load32(m.memory[int64(uint32(v3))+4:]))
				v2 = t54
				if uint32(v2) <= uint32(v1) {
					if v2 != 0 {
						if uint32(v2) >= uint32(v1) {
							goto l24
						}
						t55 := int32(int8(m.memory[uint32(v10+v2)]))
						if t55 > i32(-65) {
							goto l24
						}
						m.fn2(i32(1080413), i32(48), i32(1075812))
						panic("unreachable")
					}
					v2 = i32(0)
					goto l24
				}
				v2 = v1
				goto l24
			}
		l24:
			{
				{
					{
						switch v6 + i32(-7) {
						default:
							goto l33
						case 4:
							t56 := int64(load64(m.memory[uint32(v7):]))
							t57 := int64(load64(m.memory[uint32(v7+i32(3)):]))
							if t56^i64(0x6965772d746e6f66)|(t57^i64(8388068008349085044)) != i64(0) {
								goto l33
							}
							v1 = i32(0)
							switch v2 {
							case 0:
								goto l38
							case 1:
								t79 := int32(m.memory[uint32(v10)])
								v4 = t79
								switch v4 + i32(-43) {
								case 0, 2:
									goto l38
								default:
									goto l48
								}
							case 4:
								t76 := int32(load32(m.memory[uint32(v10):]))
								if t76 != i32(1684828002) {
									goto l40
								}
								m.memory[uint32(v9)] = byte(i32(1))
								goto l33
							case 6:
								t77 := int32(load32(m.memory[uint32(v10):]))
								t78 := int32(load16(m.memory[uint32(v10+i32(4)):]))
								if t77^i32(1684828002)|(t78^i32(29285)) != 0 {
									goto l40
								}
								m.memory[uint32(v9)] = byte(i32(1))
								goto l33
							default:
								goto l40
							}
						case 3:
							t58 := int64(load64(m.memory[uint32(v7):]))
							t59 := int64(load16(m.memory[uint32(v7+i32(8)):]))
							if t58^i64(8751746614951833446)|(t59^i64(25964)) != i64(0) {
								goto l33
							}
							v1 = i32(0)
							switch v2 + i32(-6) {
							case 0:
								t69 := int32(load32(m.memory[uint32(v10):]))
								t70 := int32(load16(m.memory[uint32(v10+i32(4)):]))
								t71 := v9
								var p72 int32
								if t69^i32(1818326121)|(t70^i32(25449)) == 0 {
									p72 = 1
								}
								m.memory[int64(uint32(t71))+1] = byte(p72)
								goto l33
							case 1:
								goto l44
							default:
								goto l45
							}
						case 8:
							t60 := int64(load64(m.memory[uint32(v7):]))
							t61 := int64(load64(m.memory[uint32(v7+i32(7)):]))
							if t60^i64(7162240928792995188)|(t61^i64(0x6e6f697461726f63)) != i64(0) {
								goto l33
							}
							goto l46
						case 13:
							t62 := int64(load64(m.memory[uint32(v7):]))
							t63 := int64(load64(m.memory[uint32(v7+i32(8)):]))
							t64 := int64(load32(m.memory[uint32(v7+i32(16)):]))
							if t62^i64(7162240928792995188)|(t63^i64(3273676477859721839))|(t64^i64(1701734764)) == 0 {
								goto l46
							}
							goto l33
						case 0:
							t65 := int32(load32(m.memory[uint32(v7):]))
							t66 := int32(load32(m.memory[uint32(v7+i32(3)):]))
							if t65^i32(1886611812)|(t66^i32(2036427888)) != 0 {
								goto l33
							}
							v1 = i32(0)
							{
								if v2 != i32(4) {
									goto l47
								}
								t67 := int32(load32(m.memory[uint32(v10):]))
								var p68 int32
								if t67 == i32(1701736302) {
									p68 = 1
								}
								v1 = p68
							}
						l47:
							m.memory[int64(uint32(v9))+4] = byte(v1)
							goto l33
						}
					l44:
						t73 := int32(load32(m.memory[uint32(v10):]))
						t74 := int32(load32(m.memory[uint32(v10+i32(3)):]))
						var p75 int32
						if t73^i32(1768710767)|(t74^i32(1702195561)) == 0 {
							p75 = 1
						}
						v1 = p75
					}
				l45:
					m.memory[int64(uint32(v9))+1] = byte(v1)
					goto l33
				l40:
					t80 := int32(m.memory[uint32(v10)])
					v4 = t80
				}
			l48:
				t81 := v10
				var p82 int32
				if v4&i32(255) == i32(43) {
					p82 = 1
				}
				v1 = p82
				v4 = t81 + v1
				v2 = v2 - v1
				if uint32(v2) < uint32(i32(9)) {
					goto l49
				}
				v8 = i32(0)
			l51:
				if v2 != 0 {
					v1 = i32(0)
					v13 = int64(uint32(v8)) * i64(10)
					if int32(int64(uint64(v13)>>32)) != 0 {
						goto l38
					}
					t83 := int32(m.memory[uint32(v4)])
					v5 = t83 + i32(-48)
					if uint32(v5) > uint32(i32(9)) {
						goto l38
					}
					v4 = v4 + i32(1)
					v2 = v2 + i32(-1)
					v8 = v5 + int32(v13)
					if uint32(v8) >= uint32(v5) {
						goto l51
					}
					goto l38
				}
				v1 = i32(1)
				v12 = v8
				goto l38
			l49:
				if v2 != 0 {
					goto l52
				}
				v1 = i32(1)
				v12 = i32(0)
				goto l38
			l52:
				v1 = i32(0)
				t84 := int32(m.memory[uint32(v4)])
				v5 = t84 + i32(-48)
				if uint32(v5) > uint32(i32(9)) {
					goto l38
				}
				{
					if v2 != i32(1) {
						goto l53
					}
					v12 = v5
					goto l54
				l53:
					t85 := int32(m.memory[int64(uint32(v4))+1])
					v8 = t85 + i32(-48)
					if uint32(v8) > uint32(i32(9)) {
						goto l38
					}
					v5 = v8 + v5*i32(10)
					if v2 != i32(2) {
						goto l55
					}
					v12 = v5
					goto l54
				l55:
					t86 := int32(m.memory[int64(uint32(v4))+2])
					v8 = t86 + i32(-48)
					if uint32(v8) > uint32(i32(9)) {
						goto l38
					}
					v5 = v8 + v5*i32(10)
					if v2 != i32(3) {
						goto l56
					}
					v12 = v5
					goto l54
				l56:
					t87 := int32(m.memory[int64(uint32(v4))+3])
					v8 = t87 + i32(-48)
					if uint32(v8) > uint32(i32(9)) {
						goto l38
					}
					v5 = v8 + v5*i32(10)
					if v2 != i32(4) {
						goto l57
					}
					v12 = v5
					goto l54
				l57:
					t88 := int32(m.memory[int64(uint32(v4))+4])
					v8 = t88 + i32(-48)
					if uint32(v8) > uint32(i32(9)) {
						goto l38
					}
					v5 = v8 + v5*i32(10)
					if v2 != i32(5) {
						goto l58
					}
					v12 = v5
					goto l54
				l58:
					t89 := int32(m.memory[int64(uint32(v4))+5])
					v8 = t89 + i32(-48)
					if uint32(v8) > uint32(i32(9)) {
						goto l38
					}
					v5 = v8 + v5*i32(10)
					if v2 != i32(6) {
						goto l59
					}
					v12 = v5
					goto l54
				l59:
					t90 := int32(m.memory[int64(uint32(v4))+6])
					v8 = t90 + i32(-48)
					if uint32(v8) > uint32(i32(9)) {
						goto l38
					}
					v5 = v8 + v5*i32(10)
					if v2 != i32(7) {
						goto l60
					}
					v12 = v5
					goto l54
				l60:
					t91 := int32(m.memory[int64(uint32(v4))+7])
					v2 = t91 + i32(-48)
					if uint32(v2) > uint32(i32(9)) {
						goto l38
					}
					v12 = v2 + v5*i32(10)
				}
			l54:
				v1 = i32(1)
			}
		l38:
			t92 := v9
			t93 := v1
			var p94 int32
			if uint32(v12) > uint32(i32(599)) {
				p94 = 1
			}
			m.memory[uint32(t92)] = byte(t93 & p94)
			goto l33
		}
	l46:
		if uint32(v2) > uint32(i32(12)) {
			m.fn164(v3+i32(96), v10, v2, i32(1075849), i32(12))
			m.fn165(v3+i32(84), v3+i32(96))
			t97 := int32(load32(m.memory[int64(uint32(v3))+84:]))
			if t97 != 0 {
				goto l64
			}
			goto l33
		}
		switch v2 + i32(-4) {
		case 0:
			t98 := int32(load32(m.memory[uint32(v10):]))
			if t98 != i32(1701736302) {
				goto l33
			}
			m.memory[int64(uint32(v9))+2] = byte(i32(0))
			goto l33
		default:
			goto l33
		case 8:
			t95 := int64(load64(m.memory[uint32(v10):]))
			t96 := int64(load32(m.memory[uint32(v10+i32(8)):]))
			if !(t95^i64(8243966856225778028)|(t96^i64(0x6867756f)) == 0) {
				goto l33
			}
			goto l64
		}
	l64:
		m.memory[int64(uint32(v9))+2] = byte(i32(1))
	l33:
		{
			{
				if v11 == 0 {
					goto l65
				}
				t99 := int32(load32(m.memory[uint32(v10+i32(-4)):]))
				v2 = t99
				v1 = v2 & i32(-8)
				t100 := v1
				v2 = v2 & i32(3)
				p101 := i32(8)
				if v2 != 0 {
					p101 = i32(4)
				}
				if uint32(t100) < uint32(p101+v11) {
					goto l66
				}
				if v2 == 0 {
					goto l67
				}
				if uint32(v1) > uint32(v11+i32(39)) {
					m.fn2(i32(1273904), i32(46), i32(1273952))
					panic("unreachable")
				}
			l67:
				m.fn1(v10)
			}
		l65:
			{
				if v6 == 0 {
					goto l69
				}
				t102 := int32(load32(m.memory[uint32(v7+i32(-4)):]))
				v2 = t102
				v1 = v2 & i32(-8)
				t103 := v1
				v2 = v2 & i32(3)
				p104 := i32(8)
				if v2 != 0 {
					p104 = i32(4)
				}
				if uint32(t103) < uint32(p104+v6) {
					m.fn2(i32(1273840), i32(46), i32(1273888))
					panic("unreachable")
				}
				if v2 == 0 {
					goto l71
				}
				if uint32(v1) > uint32(v6+i32(39)) {
					m.fn2(i32(1273904), i32(46), i32(1273952))
					panic("unreachable")
				}
			l71:
				m.fn1(v7)
			}
		l69:
			t105 := int32(m.memory[int64(uint32(v3))+81])
			if t105 == 0 {
				goto l10
			}
			goto l2
		}
	l66:
	}
	m.fn2(i32(1273840), i32(46), i32(1273888))
	panic("unreachable")
l2:
	t106 := int32(load16(m.memory[int64(uint32(v3))+40:]))
	store16(m.memory[int64(uint32(v0))+8:], uint16(t106))
	t107 := int64(load64(m.memory[int64(uint32(v3))+32:]))
	store64(m.memory[uint32(v0):], uint64(t107))
	m.g0 = v3 + i32(160)
}
func (m *Module) fn697(v0, v1 int32) int32 {
	var v2, v3 int32
	var v4, v5, v6 int64
	var v7, v8 int32
	t0 := m.g0
	v2 = t0 - i32(32)
	m.g0 = v2
	v3 = i32(20)
	{
		t1 := int64(load64(m.memory[uint32(v0):]))
		v4 = t1
		t2 := v4
		v5 = v4 >> 63
		v6 = t2 ^ v5 - v5
		if uint64(v6) < uint64(i64(1000)) {
			goto l0
		}
		v3 = i32(20)
	l1:
		{
			v0 = v2 + i32(12) + v3
			t3 := v0 + i32(-4)
			v5 = v6
			t4 := int64(uint64(v5) / uint64(i64(10000)))
			t5 := v5
			v6 = t4
			v7 = int32(t5 - v6*i64(10000))
			t6 := int32(uint32(v7&i32(0xffff)) / uint32(i32(100)))
			v8 = t6
			t7 := int32(load16(m.memory[int64(uint32(v8<<1))+1100215:]))
			store16(m.memory[uint32(t3):], uint16(t7))
			t8 := int32(load16(m.memory[int64(uint32((v7-v8*i32(100))&i32(0xffff)<<1))+1100215:]))
			store16(m.memory[uint32(v0+i32(-2)):], uint16(t8))
			v3 = v3 + i32(-4)
			if uint64(v5) > uint64(i64(9999999)) {
				goto l1
			}
		}
	}
l0:
	{
		if uint64(v6) <= uint64(i64(9)) {
			goto l2
		}
		t9 := v2 + i32(12)
		v3 = v3 + i32(-2)
		t10 := t9 + v3
		v0 = int32(v6)
		t11 := int32(uint32(v0&i32(0xffff)) / uint32(i32(100)))
		t12 := v0
		v0 = t11
		t13 := int32(load16(m.memory[int64(uint32((t12-v0*i32(100))&i32(0xffff)<<1))+1100215:]))
		store16(m.memory[uint32(t10):], uint16(t13))
		v6 = int64(uint32(v0))
	}
l2:
	{
		if v4 == 0 {
			goto l3
		}
		if v6 == 0 {
			goto l4
		}
	l3:
		t14 := v2 + i32(12)
		v3 = v3 + i32(-1)
		t15 := int32(m.memory[int64(uint32(int32(v6)<<1))+1100216])
		m.memory[uint32(t14+v3)] = byte(t15)
	}
l4:
	t16 := v1
	var p17 int32
	if v4 > i64(-1) {
		p17 = 1
	}
	t18 := m.fn312(t16, p17, i32(1), i32(0), v2+i32(12)+v3, i32(20)-v3)
	v3 = t18
	m.g0 = v2 + i32(32)
	return v3
}
func (m *Module) fn698(v0, v1, v2 int32) {
	var v3, v4, v5, v6, v7, v8, v9, v10, v11 int32
	t0 := m.g0
	v3 = t0 - i32(16)
	m.g0 = v3
	if v2 <= i32(-1) {
		m.fn12()
		panic("unreachable")
	}
	{
		{
			if v2 != 0 {
				goto l1
			}
			store32(m.memory[int64(uint32(v3))+12:], uint32(i32(0)))
			store64(m.memory[int64(uint32(v3))+4:], uint64(i64(0x100000000)))
			goto l2
		l1:
			t1 := m.fn11(v2)
			v4 = t1
			if v4 == 0 {
				m.fn7(i32(1), v2)
				panic("unreachable")
			}
			v5 = i32(0)
			store32(m.memory[int64(uint32(v3))+12:], uint32(i32(0)))
			store32(m.memory[int64(uint32(v3))+8:], uint32(v4))
			store32(m.memory[int64(uint32(v3))+4:], uint32(v2))
			v6 = v1 + v2
			v7 = i32(0)
		l25:
			{
				{
					t2 := int32(int8(m.memory[uint32(v1)]))
					v2 = t2
					if v2 <= i32(-1) {
						goto l4
					}
					v1 = v1 + i32(1)
					v2 = v2 & i32(255)
					goto l5
				}
			l4:
				t3 := int32(m.memory[int64(uint32(v1))+1])
				v8 = t3 & i32(63)
				v9 = v2 & i32(31)
				if uint32(v2) > uint32(i32(-33)) {
					goto l6
				}
				v2 = v9<<6 | v8
				v1 = v1 + i32(2)
				goto l5
			l6:
				t4 := int32(m.memory[int64(uint32(v1))+2])
				v8 = v8<<6 | t4&i32(63)
				if uint32(v2) >= uint32(i32(-16)) {
					goto l7
				}
				v2 = v8 | v9<<12
				v1 = v1 + i32(3)
				goto l5
			l7:
				t5 := int32(m.memory[int64(uint32(v1))+3])
				v2 = v8<<6 | t5&i32(63) | v9<<18&i32(0x1c0000)
				v1 = v1 + i32(4)
			}
		l5:
			{
				{
					if uint32(v2+i32(-9)) < uint32(i32(5)) {
						goto l8
					}
					if v2 == i32(32) {
						goto l8
					}
					{
						{
							if uint32(v2) < uint32(i32(133)) {
								var p7 int32
								if uint32(v2) < uint32(i32(128)) {
									p7 = 1
								}
								v8 = p7
								p8 := i32(2)
								if v8 != 0 {
									p8 = i32(1)
								}
								v9 = p8
								goto l15
							}
							v9 = int32(uint32(v2) >> 8)
							switch v9 + i32(-22) {
							case 0:
								if v2 != i32(5760) {
									goto l11
								}
								goto l8
							case 1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25:
								goto l11
							default:
								if v9 == 0 {
									goto l16
								}
								goto l11
							case 26:
								if v2 != i32(12288) {
									goto l11
								}
								goto l8
							case 10:
								t6 := int32(m.memory[int64(uint32(v2&i32(255)))+1139180])
								if t6&i32(2) == 0 {
									goto l11
								}
								goto l8
							}
						l16:
							t9 := int32(m.memory[int64(uint32(v2&i32(255)))+1139180])
							if t9&i32(1) != 0 {
								goto l8
							}
						}
					l11:
						v9 = i32(2)
						v8 = i32(0)
						if uint32(v2) < uint32(i32(2048)) {
							goto l15
						}
						p10 := i32(4)
						if uint32(v2) < uint32(i32(65536)) {
							p10 = i32(3)
						}
						v9 = p10
					}
				l15:
					{
						t11 := int32(load32(m.memory[int64(uint32(v3))+4:]))
						if uint32(v9) <= uint32(t11-v7) {
							goto l17
						}
						m.fn203(v3+i32(4), v7, v9, i32(1), i32(1))
					}
				l17:
					t12 := int32(load32(m.memory[int64(uint32(v3))+8:]))
					v4 = t12
					v5 = v4 + v7
					if v8 != 0 {
						goto l18
					}
					v8 = v2&i32(63) | i32(-128)
					v10 = int32(uint32(v2) >> 6)
					if uint32(v2) >= uint32(i32(2048)) {
						v11 = int32(uint32(v2) >> 12)
						v10 = v10&i32(63) | i32(-128)
						if uint32(v2) > uint32(i32(0xffff)) {
							m.memory[int64(uint32(v5))+3] = byte(v8)
							m.memory[int64(uint32(v5))+2] = byte(v10)
							m.memory[int64(uint32(v5))+1] = byte(v11&i32(63) | i32(-128))
							m.memory[uint32(v5)] = byte(int32(uint32(v2)>>18) | i32(-16))
							goto l20
						}
						m.memory[int64(uint32(v5))+2] = byte(v8)
						m.memory[int64(uint32(v5))+1] = byte(v10)
						m.memory[uint32(v5)] = byte(v11 | i32(224))
						goto l20
					}
					m.memory[int64(uint32(v5))+1] = byte(v8)
					m.memory[uint32(v5)] = byte(v10 | i32(192))
					goto l20
				}
			l8:
				if v5&i32(1) == 0 {
					{
						t13 := int32(load32(m.memory[int64(uint32(v3))+4:]))
						if t13 != v7 {
							goto l24
						}
						m.fn203(v3+i32(4), v7, i32(1), i32(1), i32(1))
						t14 := int32(load32(m.memory[int64(uint32(v3))+8:]))
						v4 = t14
					}
				l24:
					m.memory[uint32(v4+v7)] = byte(i32(32))
					v5 = i32(1)
					t15 := v3
					v7 = v7 + i32(1)
					store32(m.memory[int64(uint32(t15))+12:], uint32(v7))
					goto l23
				}
				v5 = i32(1)
				goto l23
			l18:
				m.memory[uint32(v5)] = byte(v2)
			l20:
				v5 = i32(0)
				t16 := v3
				v7 = v9 + v7
				store32(m.memory[int64(uint32(t16))+12:], uint32(v7))
			}
		l23:
			if v1 != v6 {
				goto l25
			}
		}
	l2:
		t17 := int32(load32(m.memory[int64(uint32(v3))+12:]))
		store32(m.memory[int64(uint32(v0))+8:], uint32(t17))
		t18 := int64(load64(m.memory[int64(uint32(v3))+4:]))
		store64(m.memory[uint32(v0):], uint64(t18))
		m.g0 = v3 + i32(16)
		return
	}
}
func (m *Module) fn699(v0, v1, v2 int32) int32 {
	var v3, v4, v5, v6, v7 int32
	if v1 == 0 {
		goto l0
	}
	v3 = v0 + v1*i32(28)
l20:
	v4 = v0
	v5 = v2
l8:
	v2 = i32(1)
	{
		v6 = v3 + i32(-28)
		t0 := int32(load32(m.memory[uint32(v6):]))
		v1 = t0
		p1 := i32(1)
		if uint32(v1) > uint32(i32(2)) {
			p1 = v1 + i32(-3)
		}
		switch p1 {
		case 2, 4:
			v2 = i32(0)
			goto l0
		case 3:
			goto l4
		case 5:
			goto l0
		case 1:
			t2 := int32(load32(m.memory[uint32(v3+i32(-4)):]))
			v7 = t2
			v1 = v7 * i32(28)
			t3 := int32(load32(m.memory[uint32(v3+i32(-8)):]))
			v0 = t3
			v3 = v0 + i32(-28)
		l6:
			{
				if v1 == 0 {
					goto l4
				}
				v1 = v1 + i32(-28)
				v3 = v3 + i32(28)
				t4 := m.fn317(v3)
				if t4 == 0 {
					goto l5
				}
				goto l6
			}
		default:
			t5 := int32(load32(m.memory[uint32(v3+i32(-16)):]))
			v1 = t5
			if v1 != 0 {
				{
					t6 := int32(load32(m.memory[uint32(v3+i32(-20)):]))
					v3 = t6 + v1
					t7 := int32(int8(m.memory[uint32(v3+i32(-1))]))
					v1 = t7
					if v1 > i32(-1) {
						goto l9
					}
					{
						t8 := int32(m.memory[uint32(v3+i32(-2))])
						v6 = t8
						v2 = int32(int8(v6))
						if v2 < i32(-64) {
							goto l10
						}
						v3 = v6 & i32(31)
						goto l11
					}
				l10:
					{
						{
							t9 := int32(m.memory[uint32(v3+i32(-3))])
							v6 = t9
							v4 = int32(int8(v6))
							if v4 < i32(-64) {
								goto l12
							}
							v3 = v6 & i32(15)
							goto l13
						}
					l12:
						t10 := int32(m.memory[uint32(v3+i32(-4))])
						v3 = t10&i32(7)<<6 | v4&i32(63)
					}
				l13:
					v3 = v3<<6 | v2&i32(63)
				l11:
					v1 = v3<<6 | v1&i32(63)
				}
			l9:
				v3 = v1 + i32(-9)
				if uint32(v3) > uint32(i32(23)) {
					goto l14
				}
				if i32_shr_u(i32(8388639), v3)&i32(1) == 0 {
					goto l14
				}
				v2 = i32(1)
				goto l0
			l14:
				v2 = i32(0)
				if uint32(v1) < uint32(i32(133)) {
					goto l0
				}
				v3 = int32(uint32(v1) >> 8)
				switch v3 + i32(-22) {
				case 1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25:
					goto l0
				default:
					if v3 != 0 {
						goto l0
					}
					t11 := int32(m.memory[int64(uint32(v1&i32(255)))+1139180])
					if t11&i32(1) != 0 {
						goto l19
					}
					goto l0
				case 0:
					if v1 == i32(5760) {
						goto l19
					}
					goto l0
				case 26:
					if v1 == i32(12288) {
						goto l19
					}
					goto l0
				case 10:
					t12 := int32(m.memory[int64(uint32(v1&i32(255)))+1139180])
					if t12&i32(2) == 0 {
						goto l0
					}
				}
			l19:
				v2 = i32(1)
				goto l0
			}
		}
	}
l4:
	v3 = v6
	if v4 != v6 {
		goto l8
	}
	v2 = v5
	goto l0
l5:
	v3 = v0 + v7*i32(28)
	v2 = i32(0)
	if v7 != 0 {
		goto l20
	}
l0:
	return v2 & i32(1)
}
func (m *Module) fn700(v0, v1 int32) {
	var v2, v3, v4, v5, v6, v7 int32
	t0 := m.g0
	v2 = t0 - i32(16)
	m.g0 = v2
	t1 := int32(load32(m.memory[int64(uint32(v1))+20:]))
	v3 = t1
	v4 = v3 << 5
	t2 := int32(load32(m.memory[int64(uint32(v1))+16:]))
	v5 = t2
	{
		if v3 == 0 {
			goto l0
		}
		v6 = v4
		v7 = v5
	l3:
		{
			t3 := int32(load32(m.memory[uint32(v7+i32(8)):]))
			if t3 != i32(2) {
				goto l1
			}
			t4 := int32(load32(m.memory[uint32(v7+i32(4)):]))
			t5 := int32(load16(m.memory[uint32(t4):]))
			if t5 == i32(25705) {
				goto l2
			}
		}
	l1:
		v7 = v7 + i32(32)
		v6 = v6 + i32(-32)
		if v6 != 0 {
			goto l3
		}
		goto l0
	l2:
		t6 := int32(load32(m.memory[int64(uint32(v7))+20:]))
		v6 = t6
		if v6 == 0 {
			goto l0
		}
		t7 := int32(load32(m.memory[int64(uint32(v0))+28:]))
		t8 := int32(load32(m.memory[int64(uint32(v7))+16:]))
		t9 := int32(load32(m.memory[int64(uint32(v0))+32:]))
		t10 := int32(load32(m.memory[int64(uint32(t9))+20:]))
		m.t0[uint(t10)].(func(int32, int32, int32, int32))(v2+i32(4), t7, t8, v6)
		{
			t11 := int32(load32(m.memory[int64(uint32(v0))+20:]))
			v7 = t11
			t12 := int32(load32(m.memory[int64(uint32(v0))+12:]))
			if v7 != t12 {
				goto l4
			}
			m.fn324(v0 + i32(12))
		}
	l4:
		store32(m.memory[int64(uint32(v0))+20:], uint32(v7+i32(1)))
		t13 := int32(load32(m.memory[int64(uint32(v0))+16:]))
		v7 = t13 + v7*i32(28)
		store32(m.memory[uint32(v7):], uint32(i32(6)))
		t14 := int64(load64(m.memory[int64(uint32(v2))+4:]))
		store64(m.memory[int64(uint32(v7))+4:], uint64(t14))
		t15 := int32(load32(m.memory[int64(uint32(v2))+12:]))
		store32(m.memory[int64(uint32(v7))+12:], uint32(t15))
	}
l0:
	{
		t16 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		if t16 != i32(1) {
			goto l5
		}
		if v3 == 0 {
			goto l5
		}
		t17 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t18 := int32(m.memory[uint32(t17)])
		if t18&i32(255) != i32(97) {
			goto l5
		}
	l8:
		{
			t19 := int32(load32(m.memory[uint32(v5+i32(8)):]))
			if t19 != i32(4) {
				goto l6
			}
			t20 := int32(load32(m.memory[uint32(v5+i32(4)):]))
			t21 := int32(load32(m.memory[uint32(t20):]))
			if t21 == i32(1701667182) {
				goto l7
			}
		}
	l6:
		v5 = v5 + i32(32)
		v4 = v4 + i32(-32)
		if v4 != 0 {
			goto l8
		}
		goto l5
	l7:
		t22 := int32(load32(m.memory[int64(uint32(v5))+20:]))
		v7 = t22
		if v7 == 0 {
			goto l5
		}
		t23 := int32(load32(m.memory[int64(uint32(v0))+28:]))
		t24 := int32(load32(m.memory[int64(uint32(v5))+16:]))
		t25 := int32(load32(m.memory[int64(uint32(v0))+32:]))
		t26 := int32(load32(m.memory[int64(uint32(t25))+20:]))
		m.t0[uint(t26)].(func(int32, int32, int32, int32))(v2+i32(4), t23, t24, v7)
		{
			t27 := int32(load32(m.memory[int64(uint32(v0))+20:]))
			v7 = t27
			t28 := int32(load32(m.memory[int64(uint32(v0))+12:]))
			if v7 != t28 {
				goto l9
			}
			m.fn324(v0 + i32(12))
		}
	l9:
		store32(m.memory[int64(uint32(v0))+20:], uint32(v7+i32(1)))
		t29 := int32(load32(m.memory[int64(uint32(v0))+16:]))
		v7 = t29 + v7*i32(28)
		store32(m.memory[uint32(v7):], uint32(i32(6)))
		t30 := int64(load64(m.memory[int64(uint32(v2))+4:]))
		store64(m.memory[int64(uint32(v7))+4:], uint64(t30))
		t31 := int32(load32(m.memory[int64(uint32(v2))+12:]))
		store32(m.memory[int64(uint32(v7))+12:], uint32(t31))
	}
l5:
	m.g0 = v2 + i32(16)
}
func (m *Module) fn701(v0, v1, v2, v3, v4 int32) {
	var v5 int32
	var v6 int64
	var v7, v8, v9, v10, v11, v12, v13, v14, v15, v16, v17 int32
	t0 := m.g0
	v5 = t0 - i32(144)
	m.g0 = v5
	m.memory[int64(uint32(v5))+100] = byte(v4)
	store64(m.memory[int64(uint32(v5))+80:], uint64(i64(4)))
	store64(m.memory[int64(uint32(v5))+72:], uint64(i64(0)))
	store64(m.memory[int64(uint32(v5))+64:], uint64(i64(0x800000000)))
	t1 := int32(load32(m.memory[int64(uint32(v1))+32:]))
	store32(m.memory[int64(uint32(v5))+96:], uint32(t1))
	t2 := int64(load64(m.memory[int64(uint32(v1))+24:]))
	store64(m.memory[int64(uint32(v5))+88:], uint64(t2))
	m.fn498(v5+i32(104), v5+i32(64), v2, v3)
	{
		{
			t3 := int32(load32(m.memory[int64(uint32(v5))+104:]))
			v2 = t3
			if v2 == i32(-1) {
				t20 := int64(load64(m.memory[int64(uint32(v5))+96:]))
				store64(m.memory[int64(uint32(v5))+136:], uint64(t20))
				t21 := int64(load64(m.memory[int64(uint32(v5))+88:]))
				store64(m.memory[int64(uint32(v5))+128:], uint64(t21))
				t22 := int64(load64(m.memory[int64(uint32(v5))+80:]))
				store64(m.memory[int64(uint32(v5))+120:], uint64(t22))
				t23 := int64(load64(m.memory[int64(uint32(v5))+72:]))
				store64(m.memory[int64(uint32(v5))+112:], uint64(t23))
				t24 := int64(load64(m.memory[int64(uint32(v5))+64:]))
				store64(m.memory[int64(uint32(v5))+104:], uint64(t24))
				m.fn499(v5 + i32(104))
				t25 := int32(load32(m.memory[int64(uint32(v5))+120:]))
				v3 = t25
				t26 := int32(load32(m.memory[int64(uint32(v5))+112:]))
				v2 = t26
				t27 := int32(load32(m.memory[int64(uint32(v5))+108:]))
				v7 = t27
				t28 := int32(load32(m.memory[int64(uint32(v5))+104:]))
				v11 = t28
				{
					t29 := int32(load32(m.memory[int64(uint32(v5))+124:]))
					v4 = t29
					if v4 == 0 {
						goto l14
					}
					v1 = v3
				l15:
					m.fn343(v1)
					v1 = v1 + i32(28)
					v4 = v4 + i32(-1)
					if v4 != 0 {
						goto l15
					}
				}
			l14:
				{
					t30 := int32(load32(m.memory[int64(uint32(v5))+116:]))
					v1 = t30
					if v1 == 0 {
						goto l16
					}
					t31 := int32(load32(m.memory[uint32(v3+i32(-4)):]))
					v4 = t31
					v8 = v4 & i32(-8)
					t32 := v8
					v4 = v4 & i32(3)
					p33 := i32(8)
					if v4 != 0 {
						p33 = i32(4)
					}
					v1 = v1 * i32(28)
					if uint32(t32) < uint32(p33+v1) {
						m.fn2(i32(1273840), i32(46), i32(1273888))
						panic("unreachable")
					}
					if v4 == 0 {
						goto l18
					}
					if uint32(v8) > uint32(v1+i32(39)) {
						m.fn2(i32(1273904), i32(46), i32(1273952))
						panic("unreachable")
					}
				l18:
					m.fn1(v3)
				}
			l16:
				{
					{
						if v2 != i32(1) {
							goto l20
						}
						t34 := int32(load32(m.memory[uint32(v7):]))
						if t34 != i32(-0x80000000) {
							goto l20
						}
						store32(m.memory[uint32(v0):], uint32(i32(-1)))
						t35 := int32(load32(m.memory[int64(uint32(v7))+12:]))
						store32(m.memory[int64(uint32(v0))+12:], uint32(t35))
						t36 := int64(load64(m.memory[int64(uint32(v7))+4:]))
						store64(m.memory[int64(uint32(v0))+4:], uint64(t36))
						store32(m.memory[int64(uint32(v7))+12:], uint32(i32(0)))
						store64(m.memory[int64(uint32(v7))+4:], uint64(i64(0x400000000)))
						m.fn341(v7)
						if v11 == 0 {
							goto l13
						}
						t37 := int32(load32(m.memory[uint32(v7+i32(-4)):]))
						v1 = t37
						v4 = v1 & i32(-8)
						t38 := v4
						v1 = v1 & i32(3)
						p39 := i32(8)
						if v1 != 0 {
							p39 = i32(4)
						}
						v0 = v11 << 5
						if uint32(t38) < uint32(p39|v0) {
							m.fn2(i32(1273840), i32(46), i32(1273888))
							panic("unreachable")
						}
						if v1 == 0 {
							goto l22
						}
						if uint32(v4) > uint32(v0+i32(39)) {
							m.fn2(i32(1273904), i32(46), i32(1273952))
							panic("unreachable")
						}
					l22:
						m.fn1(v7)
						goto l13
					}
				l20:
					store32(m.memory[int64(uint32(v5))+20:], uint32(i32(0)))
					store64(m.memory[int64(uint32(v5))+12:], uint64(i64(0x400000000)))
					v12 = v7 + v2<<5
					v1 = v7
					if v2 == 0 {
						goto l24
					}
					v1 = v7 + i32(32)
					t40 := int32(load32(m.memory[uint32(v7):]))
					v4 = t40
					if v4 == i32(-1) {
						goto l24
					}
					t41 := int32(load32(m.memory[int64(uint32(v7))+16:]))
					v13 = t41
					t42 := int32(load32(m.memory[int64(uint32(v7))+12:]))
					v8 = t42
					t43 := int32(load32(m.memory[int64(uint32(v7))+8:]))
					v9 = t43
					t44 := int32(load32(m.memory[int64(uint32(v7))+4:]))
					v14 = t44
					t45 := int32(load32(m.memory[int64(uint32(v7))+28:]))
					store32(m.memory[int64(uint32(v5))+32:], uint32(t45))
					t46 := int64(load64(m.memory[int64(uint32(v7))+20:]))
					store64(m.memory[int64(uint32(v5))+24:], uint64(t46))
					v3 = v8
					v10 = v14
					v15 = v4 >> 31 & (v4 + i32(-0x7fffffff))
					switch v15 {
					case 0:
						v3 = v9
						v10 = v4
						v9 = v14
						fallthrough
					case 1:
						{
							if v3 != 0 {
								goto l28
							}
							v4 = i32(0)
							v16 = i32(4)
							goto l29
						l28:
							m.fn203(v5+i32(12), i32(0), v3, i32(4), i32(28))
							t47 := int32(load32(m.memory[int64(uint32(v5))+20:]))
							v4 = t47
							t48 := int32(load32(m.memory[int64(uint32(v5))+16:]))
							v16 = t48
							v14 = v3 * i32(28)
							if v14 == 0 {
								goto l29
							}
							memory_copy(m.memory, uint32(v16+v4*i32(28)), uint32(v9), uint32(v14))
						}
					l29:
						t49 := v5
						v3 = v4 + v3
						store32(m.memory[int64(uint32(t49))+20:], uint32(v3))
						{
							if v10 == 0 {
								goto l30
							}
							t50 := int32(load32(m.memory[uint32(v9+i32(-4)):]))
							v4 = t50
							v14 = v4 & i32(-8)
							t51 := v14
							v4 = v4 & i32(3)
							p52 := i32(8)
							if v4 != 0 {
								p52 = i32(4)
							}
							v10 = v10 * i32(28)
							if uint32(t51) < uint32(p52+v10) {
								m.fn2(i32(1273840), i32(46), i32(1273888))
								panic("unreachable")
							}
							if v4 == 0 {
								goto l32
							}
							if uint32(v14) > uint32(v10+i32(39)) {
								m.fn2(i32(1273904), i32(46), i32(1273952))
								panic("unreachable")
							}
						l32:
							m.fn1(v9)
						}
					l30:
						if v15 != 0 {
							goto l34
						}
						if uint32(v8+i32(-1)) > uint32(i32(-3)) {
							goto l34
						}
						t53 := int32(load32(m.memory[uint32(v13+i32(-4)):]))
						v4 = t53
						v9 = v4 & i32(-8)
						t54 := v9
						v4 = v4 & i32(3)
						p55 := i32(8)
						if v4 != 0 {
							p55 = i32(4)
						}
						if uint32(t54) < uint32(p55+v8) {
							m.fn2(i32(1273840), i32(46), i32(1273888))
							panic("unreachable")
						}
						if v4 == 0 {
							goto l36
						}
						if uint32(v9) > uint32(v8+i32(39)) {
							m.fn2(i32(1273904), i32(46), i32(1273952))
							panic("unreachable")
						}
					l36:
						m.fn1(v13)
						goto l34
					default:
						store32(m.memory[int64(uint32(v5))+120:], uint32(v13))
						store32(m.memory[int64(uint32(v5))+116:], uint32(v8))
						store32(m.memory[int64(uint32(v5))+112:], uint32(v9))
						store32(m.memory[int64(uint32(v5))+108:], uint32(v14))
						store32(m.memory[int64(uint32(v5))+104:], uint32(v4))
						t56 := v5
						v4 = v7 + i32(20)
						t57 := int64(load64(m.memory[uint32(v4):]))
						store64(m.memory[int64(uint32(t56))+124:], uint64(t57))
						t58 := int32(load32(m.memory[int64(uint32(v4))+8:]))
						store32(m.memory[int64(uint32(v5))+132:], uint32(t58))
						m.fn706(v5+i32(64), v5+i32(104))
						t59 := int32(load32(m.memory[int64(uint32(v5))+68:]))
						t60 := v5 + i32(52)
						v4 = t59
						t61 := int32(load32(m.memory[int64(uint32(v5))+72:]))
						m.fn698(t60, v4, t61)
						t62 := int32(load32(m.memory[int64(uint32(v5))+60:]))
						store32(m.memory[int64(uint32(v5))+48:], uint32(t62))
						t63 := int64(load64(m.memory[int64(uint32(v5))+52:]))
						store64(m.memory[int64(uint32(v5))+40:], uint64(t63))
						m.fn324(v5 + i32(12))
						t64 := int32(load32(m.memory[int64(uint32(v5))+16:]))
						v16 = t64
						store32(m.memory[uint32(v16):], uint32(i32(3)))
						t65 := int64(load64(m.memory[int64(uint32(v5))+40:]))
						store64(m.memory[int64(uint32(v16))+4:], uint64(t65))
						t66 := int32(load32(m.memory[int64(uint32(v5))+48:]))
						store32(m.memory[int64(uint32(v16))+12:], uint32(t66))
						store32(m.memory[int64(uint32(v16))+16:], uint32(i32(0)))
						v3 = i32(1)
						store32(m.memory[int64(uint32(v5))+20:], uint32(i32(1)))
						{
							t67 := int32(load32(m.memory[int64(uint32(v5))+64:]))
							v8 = t67
							if v8 == 0 {
								goto l38
							}
							m.fn21(v4, v8, i32(1))
						}
					l38:
						m.fn341(v5 + i32(104))
					}
				l34:
					if v2 == i32(1) {
						goto l39
					}
					v17 = v5 + i32(104) + i32(20)
				l61:
					{
						t68 := int32(load32(m.memory[uint32(v1):]))
						v4 = t68
						if v4 != i32(-1) {
							t69 := int32(load32(m.memory[uint32(v1+i32(16)):]))
							v13 = t69
							t70 := int32(load32(m.memory[uint32(v1+i32(12)):]))
							v10 = t70
							t71 := int32(load32(m.memory[uint32(v1+i32(8)):]))
							v14 = t71
							t72 := int32(load32(m.memory[uint32(v1+i32(4)):]))
							v9 = t72
							t73 := int32(load32(m.memory[uint32(v1+i32(28)):]))
							store32(m.memory[int64(uint32(v5))+32:], uint32(t73))
							t74 := int64(load64(m.memory[uint32(v1+i32(20)):]))
							store64(m.memory[int64(uint32(v5))+24:], uint64(t74))
							{
								t75 := int32(load32(m.memory[int64(uint32(v5))+12:]))
								if v3 != t75 {
									goto l41
								}
								m.fn324(v5 + i32(12))
								t76 := int32(load32(m.memory[int64(uint32(v5))+16:]))
								v16 = t76
							}
						l41:
							store32(m.memory[uint32(v16+v3*i32(28)):], uint32(i32(8)))
							t77 := v5
							v8 = v3 + i32(1)
							store32(m.memory[int64(uint32(t77))+20:], uint32(v8))
							v2 = v14
							v15 = v4 >> 31 & (v4 + i32(-0x7fffffff))
							switch v15 {
							case 1:
								v2 = v10
								v4 = v9
								v9 = v14
								fallthrough
							case 0:
								{
									{
										t78 := int32(load32(m.memory[int64(uint32(v5))+12:]))
										if uint32(v2) <= uint32(t78-v8) {
											goto l45
										}
										m.fn203(v5+i32(12), v8, v2, i32(4), i32(28))
										t79 := int32(load32(m.memory[int64(uint32(v5))+20:]))
										v8 = t79
										goto l46
									}
								l45:
									if v2 == 0 {
										goto l47
									}
								l46:
									t80 := int32(load32(m.memory[int64(uint32(v5))+16:]))
									v16 = t80
									v3 = v2 * i32(28)
									if v3 == 0 {
										goto l47
									}
									memory_copy(m.memory, uint32(v16+v8*i32(28)), uint32(v9), uint32(v3))
								}
							l47:
								t81 := v5
								v3 = v8 + v2
								store32(m.memory[int64(uint32(t81))+20:], uint32(v3))
								{
									if v4 == 0 {
										goto l48
									}
									t82 := int32(load32(m.memory[uint32(v9+i32(-4)):]))
									v2 = t82
									v8 = v2 & i32(-8)
									t83 := v8
									v2 = v2 & i32(3)
									p84 := i32(8)
									if v2 != 0 {
										p84 = i32(4)
									}
									v4 = v4 * i32(28)
									if uint32(t83) < uint32(p84+v4) {
										m.fn2(i32(1273840), i32(46), i32(1273888))
										panic("unreachable")
									}
									if v2 == 0 {
										goto l50
									}
									if uint32(v8) > uint32(v4+i32(39)) {
										m.fn2(i32(1273904), i32(46), i32(1273952))
										panic("unreachable")
									}
								l50:
									m.fn1(v9)
								}
							l48:
								if v15 != 0 {
									goto l52
								}
								if uint32(v10+i32(-1)) > uint32(i32(-3)) {
									goto l52
								}
								t85 := int32(load32(m.memory[uint32(v13+i32(-4)):]))
								v4 = t85
								v2 = v4 & i32(-8)
								t86 := v2
								v4 = v4 & i32(3)
								p87 := i32(8)
								if v4 != 0 {
									p87 = i32(4)
								}
								if uint32(t86) < uint32(p87+v10) {
									m.fn2(i32(1273840), i32(46), i32(1273888))
									panic("unreachable")
								}
								if v4 == 0 {
									goto l54
								}
								if uint32(v2) > uint32(v10+i32(39)) {
									m.fn2(i32(1273904), i32(46), i32(1273952))
									panic("unreachable")
								}
							l54:
								m.fn1(v13)
								goto l52
							default:
								t88 := int64(load64(m.memory[int64(uint32(v5))+24:]))
								store64(m.memory[uint32(v17):], uint64(t88))
								t89 := int32(load32(m.memory[int64(uint32(v5))+32:]))
								store32(m.memory[int64(uint32(v17))+8:], uint32(t89))
								store32(m.memory[int64(uint32(v5))+120:], uint32(v13))
								store32(m.memory[int64(uint32(v5))+116:], uint32(v10))
								store32(m.memory[int64(uint32(v5))+112:], uint32(v14))
								store32(m.memory[int64(uint32(v5))+108:], uint32(v9))
								store32(m.memory[int64(uint32(v5))+104:], uint32(v4))
								m.fn706(v5+i32(64), v5+i32(104))
								t90 := int32(load32(m.memory[int64(uint32(v5))+68:]))
								t91 := v5 + i32(52)
								v2 = t90
								t92 := int32(load32(m.memory[int64(uint32(v5))+72:]))
								m.fn698(t91, v2, t92)
								t93 := int32(load32(m.memory[int64(uint32(v5))+60:]))
								store32(m.memory[int64(uint32(v5))+48:], uint32(t93))
								t94 := int64(load64(m.memory[int64(uint32(v5))+52:]))
								store64(m.memory[int64(uint32(v5))+40:], uint64(t94))
								{
									t95 := int32(load32(m.memory[int64(uint32(v5))+12:]))
									if v8 != t95 {
										goto l56
									}
									m.fn324(v5 + i32(12))
								}
							l56:
								t96 := int32(load32(m.memory[int64(uint32(v5))+16:]))
								v16 = t96
								v4 = v16 + v8*i32(28)
								store32(m.memory[uint32(v4):], uint32(i32(3)))
								t97 := int64(load64(m.memory[int64(uint32(v5))+40:]))
								store64(m.memory[int64(uint32(v4))+4:], uint64(t97))
								t98 := int32(load32(m.memory[int64(uint32(v5))+48:]))
								store32(m.memory[int64(uint32(v4))+12:], uint32(t98))
								store32(m.memory[int64(uint32(v4))+16:], uint32(i32(0)))
								t99 := v5
								v3 = v3 + i32(2)
								store32(m.memory[int64(uint32(t99))+20:], uint32(v3))
								{
									t100 := int32(load32(m.memory[int64(uint32(v5))+64:]))
									v4 = t100
									if v4 == 0 {
										goto l57
									}
									t101 := int32(load32(m.memory[uint32(v2+i32(-4)):]))
									v8 = t101
									v9 = v8 & i32(-8)
									t102 := v9
									v8 = v8 & i32(3)
									p103 := i32(8)
									if v8 != 0 {
										p103 = i32(4)
									}
									if uint32(t102) < uint32(p103+v4) {
										m.fn2(i32(1273840), i32(46), i32(1273888))
										panic("unreachable")
									}
									if v8 == 0 {
										goto l59
									}
									if uint32(v9) > uint32(v4+i32(39)) {
										m.fn2(i32(1273904), i32(46), i32(1273952))
										panic("unreachable")
									}
								l59:
									m.fn1(v2)
								}
							l57:
								m.fn341(v5 + i32(104))
								goto l52
							}
						l52:
							v1 = v1 + i32(32)
							if v1 != v12 {
								goto l61
							}
							goto l39
						}
						v1 = v1 + i32(32)
						goto l24
					}
				}
			}
			t4 := int64(load64(m.memory[int64(uint32(v5))+120:]))
			v6 = t4
			t5 := int32(load32(m.memory[int64(uint32(v5))+116:]))
			v7 = t5
			t6 := int32(load32(m.memory[int64(uint32(v5))+112:]))
			v8 = t6
			t7 := int32(load32(m.memory[int64(uint32(v5))+108:]))
			v9 = t7
			t8 := int32(load32(m.memory[int64(uint32(v5))+68:]))
			v3 = t8
			{
				t9 := int32(load32(m.memory[int64(uint32(v5))+72:]))
				v4 = t9
				if v4 == 0 {
					goto l1
				}
				v1 = v3
			l2:
				m.fn341(v1)
				v1 = v1 + i32(32)
				v4 = v4 + i32(-1)
				if v4 != 0 {
					goto l2
				}
			}
		l1:
			{
				t10 := int32(load32(m.memory[int64(uint32(v5))+64:]))
				v1 = t10
				if v1 == 0 {
					goto l3
				}
				t11 := int32(load32(m.memory[uint32(v3+i32(-4)):]))
				v4 = t11
				v10 = v4 & i32(-8)
				t12 := v10
				v4 = v4 & i32(3)
				p13 := i32(8)
				if v4 != 0 {
					p13 = i32(4)
				}
				v1 = v1 << 5
				if uint32(t12) < uint32(p13|v1) {
					m.fn2(i32(1273840), i32(46), i32(1273888))
					panic("unreachable")
				}
				if v4 == 0 {
					goto l5
				}
				if uint32(v10) > uint32(v1+i32(39)) {
					m.fn2(i32(1273904), i32(46), i32(1273952))
					panic("unreachable")
				}
			l5:
				m.fn1(v3)
			}
		l3:
			t14 := int32(load32(m.memory[int64(uint32(v5))+80:]))
			v3 = t14
			{
				t15 := int32(load32(m.memory[int64(uint32(v5))+84:]))
				v4 = t15
				if v4 == 0 {
					goto l7
				}
				v1 = v3
			l8:
				m.fn343(v1)
				v1 = v1 + i32(28)
				v4 = v4 + i32(-1)
				if v4 != 0 {
					goto l8
				}
			}
		l7:
			{
				t16 := int32(load32(m.memory[int64(uint32(v5))+76:]))
				v1 = t16
				if v1 == 0 {
					goto l9
				}
				t17 := int32(load32(m.memory[uint32(v3+i32(-4)):]))
				v4 = t17
				v10 = v4 & i32(-8)
				t18 := v10
				v4 = v4 & i32(3)
				p19 := i32(8)
				if v4 != 0 {
					p19 = i32(4)
				}
				v1 = v1 * i32(28)
				if uint32(t18) < uint32(p19+v1) {
					m.fn2(i32(1273840), i32(46), i32(1273888))
					panic("unreachable")
				}
				if v4 == 0 {
					goto l11
				}
				if uint32(v10) > uint32(v1+i32(39)) {
					m.fn2(i32(1273904), i32(46), i32(1273952))
					panic("unreachable")
				}
			l11:
				m.fn1(v3)
			}
		l9:
			store64(m.memory[int64(uint32(v0))+16:], uint64(v6))
			store32(m.memory[int64(uint32(v0))+12:], uint32(v7))
			store32(m.memory[int64(uint32(v0))+8:], uint32(v8))
			store32(m.memory[int64(uint32(v0))+4:], uint32(v9))
			store32(m.memory[uint32(v0):], uint32(v2))
			goto l13
		}
	l24:
		if v12 == v1 {
			goto l39
		}
		v4 = int32(uint32(v12-v1) >> 5)
	l62:
		m.fn341(v1)
		v1 = v1 + i32(32)
		v4 = v4 + i32(-1)
		if v4 != 0 {
			goto l62
		}
	l39:
		{
			if v11 == 0 {
				goto l63
			}
			t104 := int32(load32(m.memory[uint32(v7+i32(-4)):]))
			v1 = t104
			v4 = v1 & i32(-8)
			t105 := v4
			v1 = v1 & i32(3)
			p106 := i32(8)
			if v1 != 0 {
				p106 = i32(4)
			}
			v3 = v11 << 5
			if uint32(t105) < uint32(p106|v3) {
				m.fn2(i32(1273840), i32(46), i32(1273888))
				panic("unreachable")
			}
			if v1 == 0 {
				goto l65
			}
			if uint32(v4) > uint32(v3+i32(39)) {
				m.fn2(i32(1273904), i32(46), i32(1273952))
				panic("unreachable")
			}
		l65:
			m.fn1(v7)
		}
	l63:
		t107 := int32(load32(m.memory[int64(uint32(v5))+20:]))
		store32(m.memory[int64(uint32(v0))+12:], uint32(t107))
		t108 := int64(load64(m.memory[int64(uint32(v5))+12:]))
		store64(m.memory[int64(uint32(v0))+4:], uint64(t108))
		store32(m.memory[uint32(v0):], uint32(i32(-1)))
	}
l13:
	m.g0 = v5 + i32(144)
}
func (m *Module) fn702(v0, v1, v2 int32) {
	var v3, v4, v5, v6, v7 int32
	t0 := m.g0
	v3 = t0 - i32(16)
	m.g0 = v3
	{
		if v1 == v2 {
			goto l0
		}
	l3:
		v4 = v1
		v1 = v4 + i32(44)
		{
			t1 := int32(load32(m.memory[uint32(v4):]))
			if t1 == i32(-1) {
				goto l1
			}
			t2 := int32(load32(m.memory[uint32(v4+i32(8)):]))
			if t2 != i32(2) {
				goto l1
			}
			t3 := int32(load32(m.memory[uint32(v4+i32(4)):]))
			t4 := int32(load16(m.memory[uint32(t3):]))
			v5 = t4
			if (v5<<8|int32(uint32(v5)>>8))&i32(0xffff) == i32(27753) {
				goto l2
			}
		}
	l1:
		if v1 != v2 {
			goto l3
		}
	l0:
		store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
		store64(m.memory[uint32(v0):], uint64(i64(0x400000000)))
		goto l4
	l2:
		t5 := m.fn11(i32(16))
		v6 = t5
		if v6 == 0 {
			m.fn7(i32(4), i32(16))
			panic("unreachable")
		}
		store32(m.memory[uint32(v6):], uint32(v4))
		store32(m.memory[int64(uint32(v3))+12:], uint32(i32(1)))
		store32(m.memory[int64(uint32(v3))+8:], uint32(v6))
		store32(m.memory[int64(uint32(v3))+4:], uint32(i32(4)))
		if v1 == v2 {
			goto l6
		}
		v7 = i32(1)
	l9:
		{
			v4 = v1
			v1 = v4 + i32(44)
			{
				t6 := int32(load32(m.memory[uint32(v4):]))
				if t6 == i32(-1) {
					goto l7
				}
				t7 := int32(load32(m.memory[uint32(v4+i32(8)):]))
				if t7 != i32(2) {
					goto l7
				}
				t8 := int32(load32(m.memory[uint32(v4+i32(4)):]))
				t9 := int32(load16(m.memory[uint32(t8):]))
				v5 = t9
				if (v5<<8|int32(uint32(v5)>>8))&i32(0xffff) == i32(27753) {
					goto l8
				}
			}
		l7:
			if v1 != v2 {
				goto l9
			}
			goto l6
		l8:
			{
				t10 := int32(load32(m.memory[int64(uint32(v3))+4:]))
				if v7 != t10 {
					goto l10
				}
				m.fn203(v3+i32(4), v7, i32(1), i32(4), i32(4))
				t11 := int32(load32(m.memory[int64(uint32(v3))+8:]))
				v6 = t11
			}
		l10:
			store32(m.memory[uint32(v6+v7<<2):], uint32(v4))
			t12 := v3
			v7 = v7 + i32(1)
			store32(m.memory[int64(uint32(t12))+12:], uint32(v7))
			if v1 != v2 {
				goto l9
			}
		}
	l6:
		t13 := int32(load32(m.memory[int64(uint32(v3))+12:]))
		store32(m.memory[int64(uint32(v0))+8:], uint32(t13))
		t14 := int64(load64(m.memory[int64(uint32(v3))+4:]))
		store64(m.memory[uint32(v0):], uint64(t14))
	}
l4:
	m.g0 = v3 + i32(16)
}
func (m *Module) fn703(v0, v1, v2, v3 int32) {
	var v4, v5 int32
	t0 := m.g0
	v4 = t0 - i32(96)
	m.g0 = v4
	m.memory[int64(uint32(v4))+36] = byte(i32(1))
	store64(m.memory[int64(uint32(v4))+16:], uint64(i64(4)))
	store64(m.memory[int64(uint32(v4))+8:], uint64(i64(0)))
	store64(m.memory[uint32(v4):], uint64(i64(0x800000000)))
	t1 := int32(load32(m.memory[int64(uint32(v1))+32:]))
	store32(m.memory[int64(uint32(v4))+32:], uint32(t1))
	t2 := int64(load64(m.memory[int64(uint32(v1))+24:]))
	store64(m.memory[int64(uint32(v4))+24:], uint64(t2))
	m.fn498(v4+i32(56), v4, v2, v3)
	{
		{
			t3 := int32(load32(m.memory[int64(uint32(v4))+56:]))
			if t3 == i32(-1) {
				goto l0
			}
			t4 := int64(load64(m.memory[int64(uint32(v4))+72:]))
			store64(m.memory[int64(uint32(v0))+16:], uint64(t4))
			t5 := int64(load64(m.memory[int64(uint32(v4))+64:]))
			store64(m.memory[int64(uint32(v0))+8:], uint64(t5))
			t6 := int64(load64(m.memory[int64(uint32(v4))+56:]))
			store64(m.memory[uint32(v0):], uint64(t6))
			t7 := int32(load32(m.memory[int64(uint32(v4))+4:]))
			v0 = t7
			{
				t8 := int32(load32(m.memory[int64(uint32(v4))+8:]))
				v3 = t8
				if v3 == 0 {
					goto l1
				}
				v1 = v0
			l2:
				m.fn341(v1)
				v1 = v1 + i32(32)
				v3 = v3 + i32(-1)
				if v3 != 0 {
					goto l2
				}
			}
		l1:
			{
				t9 := int32(load32(m.memory[uint32(v4):]))
				v1 = t9
				if v1 == 0 {
					goto l3
				}
				t10 := int32(load32(m.memory[uint32(v0+i32(-4)):]))
				v3 = t10
				v2 = v3 & i32(-8)
				t11 := v2
				v3 = v3 & i32(3)
				p12 := i32(8)
				if v3 != 0 {
					p12 = i32(4)
				}
				v1 = v1 << 5
				if uint32(t11) < uint32(p12|v1) {
					m.fn2(i32(1273840), i32(46), i32(1273888))
					panic("unreachable")
				}
				if v3 == 0 {
					goto l5
				}
				if uint32(v2) > uint32(v1+i32(39)) {
					m.fn2(i32(1273904), i32(46), i32(1273952))
					panic("unreachable")
				}
			l5:
				m.fn1(v0)
			}
		l3:
			t13 := int32(load32(m.memory[int64(uint32(v4))+16:]))
			v0 = t13
			{
				t14 := int32(load32(m.memory[int64(uint32(v4))+20:]))
				v3 = t14
				if v3 == 0 {
					goto l7
				}
				v1 = v0
			l8:
				m.fn343(v1)
				v1 = v1 + i32(28)
				v3 = v3 + i32(-1)
				if v3 != 0 {
					goto l8
				}
			}
		l7:
			t15 := int32(load32(m.memory[int64(uint32(v4))+12:]))
			v1 = t15
			if v1 == 0 {
				goto l9
			}
			t16 := int32(load32(m.memory[uint32(v0+i32(-4)):]))
			v3 = t16
			v2 = v3 & i32(-8)
			t17 := v2
			v3 = v3 & i32(3)
			p18 := i32(8)
			if v3 != 0 {
				p18 = i32(4)
			}
			v1 = v1 * i32(28)
			if uint32(t17) < uint32(p18+v1) {
				m.fn2(i32(1273840), i32(46), i32(1273888))
				panic("unreachable")
			}
			if v3 == 0 {
				goto l11
			}
			if uint32(v2) > uint32(v1+i32(39)) {
				m.fn2(i32(1273904), i32(46), i32(1273952))
				panic("unreachable")
			}
		l11:
			m.fn1(v0)
			goto l9
		}
	l0:
		t19 := int64(load64(m.memory[int64(uint32(v4))+32:]))
		store64(m.memory[int64(uint32(v4))+88:], uint64(t19))
		t20 := int64(load64(m.memory[int64(uint32(v4))+24:]))
		store64(m.memory[int64(uint32(v4))+80:], uint64(t20))
		t21 := int64(load64(m.memory[int64(uint32(v4))+16:]))
		store64(m.memory[int64(uint32(v4))+72:], uint64(t21))
		t22 := int64(load64(m.memory[int64(uint32(v4))+8:]))
		store64(m.memory[int64(uint32(v4))+64:], uint64(t22))
		t23 := int64(load64(m.memory[uint32(v4):]))
		store64(m.memory[int64(uint32(v4))+56:], uint64(t23))
		m.fn499(v4 + i32(56))
		t24 := int32(load32(m.memory[int64(uint32(v4))+64:]))
		store32(m.memory[int64(uint32(v4))+48:], uint32(t24))
		t25 := int64(load64(m.memory[int64(uint32(v4))+56:]))
		store64(m.memory[int64(uint32(v4))+40:], uint64(t25))
		t26 := int32(load32(m.memory[int64(uint32(v4))+72:]))
		v2 = t26
		{
			t27 := int32(load32(m.memory[int64(uint32(v4))+76:]))
			v3 = t27
			if v3 == 0 {
				goto l13
			}
			v1 = v2
		l14:
			m.fn343(v1)
			v1 = v1 + i32(28)
			v3 = v3 + i32(-1)
			if v3 != 0 {
				goto l14
			}
		}
	l13:
		{
			t28 := int32(load32(m.memory[int64(uint32(v4))+68:]))
			v1 = t28
			if v1 == 0 {
				goto l15
			}
			t29 := int32(load32(m.memory[uint32(v2+i32(-4)):]))
			v3 = t29
			v5 = v3 & i32(-8)
			t30 := v5
			v3 = v3 & i32(3)
			p31 := i32(8)
			if v3 != 0 {
				p31 = i32(4)
			}
			v1 = v1 * i32(28)
			if uint32(t30) < uint32(p31+v1) {
				m.fn2(i32(1273840), i32(46), i32(1273888))
				panic("unreachable")
			}
			if v3 == 0 {
				goto l17
			}
			if uint32(v5) > uint32(v1+i32(39)) {
				m.fn2(i32(1273904), i32(46), i32(1273952))
				panic("unreachable")
			}
		l17:
			m.fn1(v2)
		}
	l15:
		t32 := int32(load32(m.memory[int64(uint32(v4))+48:]))
		store32(m.memory[int64(uint32(v0))+12:], uint32(t32))
		t33 := int64(load64(m.memory[int64(uint32(v4))+40:]))
		store64(m.memory[int64(uint32(v0))+4:], uint64(t33))
		store32(m.memory[uint32(v0):], uint32(i32(-1)))
	}
l9:
	m.g0 = v4 + i32(96)
}
func (m *Module) fn704(v0, v1 int32) int32 {
	var v2, v3 int32
	v2 = i32(0)
	{
		switch v1 + i32(-2) {
		default:
			goto l6
		case 1:
			v1 = i32(1)
			t0 := int32(load16(m.memory[uint32(v0):]))
			t1 := t0 ^ i32(26980)
			v3 = v0 + i32(2)
			t2 := int32(m.memory[uint32(v3)])
			if (t1|(t2^i32(118)))&i32(0xffff) == 0 {
				goto l8
			}
			t3 := int32(load16(m.memory[uint32(v0):]))
			t4 := int32(m.memory[uint32(v3)])
			if (t3^i32(24942)|(t4^i32(118)))&i32(0xffff) != 0 {
				goto l6
			}
			goto l8
		case 5:
			v1 = i32(1)
			t5 := int32(load32(m.memory[uint32(v0):]))
			t6 := t5 ^ i32(1952671091)
			v3 = v0 + i32(3)
			t7 := int32(load32(m.memory[uint32(v3):]))
			if t6|(t7^i32(1852795252)) == 0 {
				goto l8
			}
			t8 := int32(load32(m.memory[uint32(v0):]))
			t9 := int32(load32(m.memory[uint32(v3):]))
			if t8^i32(1769239137)|(t9^i32(1701602153)) == 0 {
				goto l8
			}
			t10 := int32(load32(m.memory[uint32(v0):]))
			t11 := t10 ^ i32(1635018084)
			v3 = v0 + i32(3)
			t12 := int32(load32(m.memory[uint32(v3):]))
			if t11|(t12^i32(1936484705)) == 0 {
				goto l8
			}
			t13 := int32(load32(m.memory[uint32(v0):]))
			t14 := int32(load32(m.memory[uint32(v3):]))
			if t13^i32(1835890035)|(t14^i32(2037539181)) != 0 {
				goto l6
			}
			goto l8
		case 3:
			t15 := int32(load32(m.memory[uint32(v0):]))
			t16 := int32(m.memory[uint32(v0+i32(4))])
			if t15^i32(1684632417)|(t16^i32(101)) != 0 {
				goto l6
			}
			return i32(1)
		case 2:
			t17 := int32(load32(m.memory[uint32(v0):]))
			if t17 != i32(1852399981) {
				goto l9
			}
			return i32(1)
		case 4:
			v1 = i32(1)
			t18 := int32(load32(m.memory[uint32(v0):]))
			t19 := t18 ^ i32(1684104552)
			v3 = v0 + i32(4)
			t20 := int32(load16(m.memory[uint32(v3):]))
			if t19|(t20^i32(29285)) == 0 {
				goto l8
			}
			t21 := int32(load32(m.memory[uint32(v0):]))
			t22 := int32(load16(m.memory[uint32(v3):]))
			if t21^i32(1953460070)|(t22^i32(29285)) == 0 {
				goto l8
			}
			t23 := int32(load32(m.memory[uint32(v0):]))
			t24 := t23 ^ i32(1969711462)
			v3 = v0 + i32(4)
			t25 := int32(load16(m.memory[uint32(v3):]))
			if t24|(t25^i32(25970)) == 0 {
				goto l8
			}
			t26 := int32(load32(m.memory[uint32(v0):]))
			t27 := int32(load16(m.memory[uint32(v3):]))
			if t26^i32(1953391971)|(t27^i32(29285)) != 0 {
				goto l6
			}
			goto l8
		case 8:
			t28 := int64(load64(m.memory[uint32(v0):]))
			t29 := int64(load16(m.memory[uint32(v0+i32(8)):]))
			if !(t28^i64(7598822034862729574)|(t29^i64(28271)) == 0) {
				goto l6
			}
			return i32(1)
		case 0:
			v1 = i32(1)
			t30 := int32(load16(m.memory[uint32(v0):]))
			if t30 == i32(26988) {
				goto l8
			}
			t31 := int32(load16(m.memory[uint32(v0):]))
			if t31 == i32(27748) {
				goto l8
			}
			t32 := int32(load16(m.memory[uint32(v0):]))
			if t32 == i32(29796) {
				goto l8
			}
			t33 := int32(load16(m.memory[uint32(v0):]))
			if t33 != i32(25700) {
				goto l6
			}
			goto l8
		}
	l9:
		t34 := int32(load32(m.memory[uint32(v0):]))
		var p35 int32
		if t34 == i32(2036625250) {
			p35 = 1
		}
		v2 = p35
	}
l6:
	v1 = v2
l8:
	return v1
}
func (m *Module) fn705(v0, v1, v2 int32) {
	var v3, v4, v5, v6, v7, v8 int32
	v3 = i32(0)
	if v2 == 0 {
		goto l0
	}
	v2 = v1 + v2
l12:
	{
		v4 = v2
		v2 = v4 + i32(-1)
		t0 := int32(int8(m.memory[uint32(v2)]))
		v5 = t0
		if v5 > i32(-1) {
			goto l1
		}
		{
			v2 = v4 + i32(-2)
			t1 := int32(m.memory[uint32(v2)])
			v6 = t1
			v7 = int32(int8(v6))
			if v7 < i32(-64) {
				goto l2
			}
			v6 = v6 & i32(31)
			goto l3
		}
	l2:
		{
			{
				v2 = v4 + i32(-3)
				t2 := int32(m.memory[uint32(v2)])
				v6 = t2
				v8 = int32(int8(v6))
				if v8 < i32(-64) {
					goto l4
				}
				v6 = v6 & i32(15)
				goto l5
			}
		l4:
			v2 = v4 + i32(-4)
			t3 := int32(m.memory[uint32(v2)])
			v6 = t3&i32(7)<<6 | v8&i32(63)
		}
	l5:
		v6 = v6<<6 | v7&i32(63)
	l3:
		v5 = v6<<6 | v5&i32(63)
	}
l1:
	if uint32(v5+i32(-9)) < uint32(i32(5)) {
		goto l6
	}
	if v5 == i32(32) {
		goto l6
	}
	if uint32(v5) < uint32(i32(133)) {
		goto l7
	}
	v6 = int32(uint32(v5) >> 8)
	switch v6 + i32(-22) {
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25:
		goto l7
	case 0:
		if v5 == i32(5760) {
			goto l6
		}
		goto l7
	case 26:
		if v5 == i32(12288) {
			goto l6
		}
		goto l7
	case 10:
		t4 := int32(m.memory[int64(uint32(v5&i32(255)))+1139180])
		if t4&i32(2) != 0 {
			goto l6
		}
		goto l7
	default:
		if v6 != 0 {
			goto l7
		}
		t5 := int32(m.memory[int64(uint32(v5&i32(255)))+1139180])
		if t5&i32(1) == 0 {
			goto l7
		}
	}
l6:
	if v1 != v2 {
		goto l12
	}
	goto l0
l7:
	v3 = v4 - v1
l0:
	store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
	store32(m.memory[uint32(v0):], uint32(v1))
}
func (m *Module) fn706(v0, v1 int32) {
	var v2, v3, v4, v5, v6, v7 int32
	var v8 int64
	var v9, v10, v11, v12, v13, v14, v15, v16 int32
	t0 := m.g0
	v2 = t0 - i32(48)
	m.g0 = v2
	{
		{
			t1 := int32(load32(m.memory[uint32(v1):]))
			v3 = t1
			switch v3 >> 31 & (v3 + i32(-0x7fffffff)) {
			case 2:
				t2 := int32(load32(m.memory[int64(uint32(v1))+20:]))
				v3 = t2
				t3 := int32(load32(m.memory[int64(uint32(v1))+24:]))
				t4 := v3
				v4 = t3 * i32(28)
				v5 = t4 + v4
				v1 = i32(0)
			l11:
				{
					if v1 == 0 {
						goto l7
					}
					if v1 == v6 {
						goto l7
					}
					m.fn706(v2+i32(36), v1)
					t5 := int32(load32(m.memory[int64(uint32(v2))+36:]))
					v7 = t5
					if v7 != i32(-1) {
						t8 := int64(load64(m.memory[int64(uint32(v2))+40:]))
						v8 = t8
						t9 := v6
						v1 = v1 + i32(32)
						v4 = int32(uint32(t9-v1) >> 5)
						p10 := i32(3)
						if uint32(v4) > uint32(i32(3)) {
							p10 = v4
						}
						v4 = p10 + i32(1)
						v9 = v4 * i32(12)
						t11 := m.fn11(v9)
						v10 = t11
						if v10 == 0 {
							m.fn7(i32(4), v9)
							panic("unreachable")
						}
						store64(m.memory[int64(uint32(v10))+4:], uint64(v8))
						store32(m.memory[uint32(v10):], uint32(v7))
						store32(m.memory[int64(uint32(v2))+32:], uint32(i32(1)))
						store32(m.memory[int64(uint32(v2))+28:], uint32(v10))
						store32(m.memory[int64(uint32(v2))+24:], uint32(v4))
						v4 = i32(1)
					l16:
						{
							{
								if v1 == v6 {
									goto l13
								}
								m.fn706(v2+i32(36), v1)
								t12 := int32(load32(m.memory[int64(uint32(v2))+36:]))
								v7 = t12
								if v7 != i32(-1) {
									t15 := int64(load64(m.memory[int64(uint32(v2))+40:]))
									v8 = t15
									v1 = v1 + i32(32)
									{
										t16 := int32(load32(m.memory[int64(uint32(v2))+24:]))
										if v4 != t16 {
											goto l17
										}
										m.fn203(v2+i32(24), v4, int32(uint32(v6-v1)>>5)+i32(1), i32(4), i32(12))
										t17 := int32(load32(m.memory[int64(uint32(v2))+28:]))
										v10 = t17
									}
								l17:
									v9 = v10 + v4*i32(12)
									store64(m.memory[int64(uint32(v9))+4:], uint64(v8))
									store32(m.memory[uint32(v9):], uint32(v7))
									t18 := v2
									v4 = v4 + i32(1)
									store32(m.memory[int64(uint32(t18))+32:], uint32(v4))
									goto l16
								}
							}
						l13:
							if v3 == v5 {
								t19 := int32(load32(m.memory[int64(uint32(v2))+24:]))
								v9 = t19
								t20 := int32(load32(m.memory[int64(uint32(v2))+28:]))
								t21 := v0
								v10 = t20
								m.fn209(t21, v10, v4, i32(1089413), i32(1))
								v3 = v10
							l22:
								{
									t22 := int32(load32(m.memory[uint32(v3):]))
									v1 = t22
									if v1 == 0 {
										goto l18
									}
									t23 := int32(load32(m.memory[uint32(v3+i32(4)):]))
									v5 = t23
									t24 := int32(load32(m.memory[uint32(v5+i32(-4)):]))
									v6 = t24
									v7 = v6 & i32(-8)
									t25 := v7
									v6 = v6 & i32(3)
									p26 := i32(8)
									if v6 != 0 {
										p26 = i32(4)
									}
									if uint32(t25) < uint32(p26+v1) {
										m.fn2(i32(1273840), i32(46), i32(1273888))
										panic("unreachable")
									}
									if v6 == 0 {
										goto l20
									}
									if uint32(v7) > uint32(v1+i32(39)) {
										m.fn2(i32(1273904), i32(46), i32(1273952))
										panic("unreachable")
									}
								l20:
									m.fn1(v5)
								}
							l18:
								v3 = v3 + i32(12)
								v4 = v4 + i32(-1)
								if v4 != 0 {
									goto l22
								}
								if v9 == 0 {
									goto l10
								}
								t27 := int32(load32(m.memory[uint32(v10+i32(-4)):]))
								v3 = t27
								v1 = v3 & i32(-8)
								t28 := v1
								v3 = v3 & i32(3)
								p29 := i32(8)
								if v3 != 0 {
									p29 = i32(4)
								}
								v4 = v9 * i32(12)
								if uint32(t28) < uint32(p29+v4) {
									m.fn2(i32(1273840), i32(46), i32(1273888))
									panic("unreachable")
								}
								if v3 == 0 {
									goto l24
								}
								if uint32(v1) > uint32(v4+i32(39)) {
									m.fn2(i32(1273904), i32(46), i32(1273952))
									panic("unreachable")
								}
							l24:
								m.fn1(v10)
								goto l10
							}
							t13 := int32(load32(m.memory[uint32(v3+i32(4)):]))
							v1 = t13
							t14 := int32(load32(m.memory[uint32(v3+i32(8)):]))
							v6 = v1 + t14<<5
							v3 = v3 + i32(28)
							goto l16
						}
					}
				}
			l7:
				if v4 != 0 {
					t6 := int32(load32(m.memory[uint32(v3+i32(4)):]))
					v1 = t6
					t7 := int32(load32(m.memory[uint32(v3+i32(8)):]))
					v6 = v1 + t7<<5
					v4 = v4 + i32(-28)
					v3 = v3 + i32(28)
					goto l11
				}
				store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
				store64(m.memory[uint32(v0):], uint64(i64(0x100000000)))
				goto l10
			case 3:
				t30 := int32(load32(m.memory[int64(uint32(v1))+8:]))
				v11 = t30
				t31 := int32(load32(m.memory[int64(uint32(v1))+12:]))
				v12 = v11 + t31*i32(12)
				v3 = i32(0)
				v1 = i32(1)
			l43:
				if v1&i32(1) != 0 {
					goto l26
				}
			l41:
				if v3 == v13 {
					goto l26
				}
				v7 = v3 + i32(20)
				{
					{
						t32 := int32(load32(m.memory[uint32(v3):]))
						if t32 != i32(-1) {
							goto l27
						}
						v1 = i32(-1)
						store32(m.memory[int64(uint32(v2))+24:], uint32(i32(-1)))
						goto l28
					}
				l27:
					{
						t33 := int32(load32(m.memory[int64(uint32(v3))+8:]))
						v4 = t33
						if v4 != 0 {
							goto l29
						}
						v1 = i32(0)
						store32(m.memory[int64(uint32(v2))+32:], uint32(i32(0)))
						store64(m.memory[int64(uint32(v2))+24:], uint64(i64(0x100000000)))
						goto l30
					}
				l29:
					t34 := int32(load32(m.memory[int64(uint32(v3))+4:]))
					v1 = t34
					v14 = v4 * i32(12)
					t35 := m.fn11(v14)
					v9 = t35
					v3 = v9
					v6 = v4
					if v9 == 0 {
						m.fn7(i32(4), v14)
						panic("unreachable")
					}
				l32:
					{
						m.fn706(v2+i32(36), v1)
						t36 := int32(load32(m.memory[int64(uint32(v2))+44:]))
						store32(m.memory[int64(uint32(v3))+8:], uint32(t36))
						t37 := int64(load64(m.memory[int64(uint32(v2))+36:]))
						store64(m.memory[uint32(v3):], uint64(t37))
						v3 = v3 + i32(12)
						v1 = v1 + i32(32)
						v6 = v6 + i32(-1)
						if v6 != 0 {
							goto l32
						}
					}
					m.fn209(v2+i32(24), v9, v4, i32(1089413), i32(1))
					v3 = v9
				l37:
					{
						t38 := int32(load32(m.memory[uint32(v3):]))
						v1 = t38
						if v1 == 0 {
							goto l33
						}
						t39 := int32(load32(m.memory[uint32(v3+i32(4)):]))
						v5 = t39
						t40 := int32(load32(m.memory[uint32(v5+i32(-4)):]))
						v6 = t40
						v10 = v6 & i32(-8)
						t41 := v10
						v6 = v6 & i32(3)
						p42 := i32(8)
						if v6 != 0 {
							p42 = i32(4)
						}
						if uint32(t41) < uint32(p42+v1) {
							m.fn2(i32(1273840), i32(46), i32(1273888))
							panic("unreachable")
						}
						if v6 == 0 {
							goto l35
						}
						if uint32(v10) > uint32(v1+i32(39)) {
							m.fn2(i32(1273904), i32(46), i32(1273952))
							panic("unreachable")
						}
					l35:
						m.fn1(v5)
					}
				l33:
					v3 = v3 + i32(12)
					v4 = v4 + i32(-1)
					if v4 != 0 {
						goto l37
					}
					t43 := int32(load32(m.memory[uint32(v9+i32(-4)):]))
					v3 = t43
					v1 = v3 & i32(-8)
					t44 := v1
					v3 = v3 & i32(3)
					p45 := i32(8)
					if v3 != 0 {
						p45 = i32(4)
					}
					if uint32(t44) < uint32(p45+v14) {
						m.fn2(i32(1273840), i32(46), i32(1273888))
						panic("unreachable")
					}
					if v3 == 0 {
						goto l39
					}
					if uint32(v1) > uint32(v14+i32(39)) {
						m.fn2(i32(1273904), i32(46), i32(1273952))
						panic("unreachable")
					}
				l39:
					m.fn1(v9)
					t46 := int32(load32(m.memory[int64(uint32(v2))+24:]))
					v1 = t46
				}
			l28:
				v3 = v7
				if v1 == i32(-1) {
					goto l41
				}
				goto l30
			l26:
				{
					if v11 == v12 {
						store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
						store64(m.memory[uint32(v0):], uint64(i64(0x100000000)))
						goto l10
					}
					t47 := int32(load32(m.memory[uint32(v11+i32(4)):]))
					v3 = t47
					t48 := int32(load32(m.memory[uint32(v11+i32(8)):]))
					v13 = v3 + t48*i32(20)
					v11 = v11 + i32(12)
					v1 = i32(0)
					goto l43
				}
			case 4:
				t49 := int32(load32(m.memory[int64(uint32(v1))+12:]))
				v4 = t49
				if v4 != 0 {
					t50 := int32(load32(m.memory[int64(uint32(v1))+8:]))
					v1 = t50
					v9 = v4 * i32(12)
					t51 := m.fn11(v9)
					v10 = t51
					if v10 == 0 {
						m.fn7(i32(4), v9)
						panic("unreachable")
					}
					v3 = v10
					v6 = v4
				l46:
					{
						m.fn706(v2+i32(36), v1)
						t52 := int32(load32(m.memory[int64(uint32(v2))+44:]))
						store32(m.memory[int64(uint32(v3))+8:], uint32(t52))
						t53 := int64(load64(m.memory[int64(uint32(v2))+36:]))
						store64(m.memory[uint32(v3):], uint64(t53))
						v3 = v3 + i32(12)
						v1 = v1 + i32(32)
						v6 = v6 + i32(-1)
						if v6 != 0 {
							goto l46
						}
					}
					m.fn209(v0, v10, v4, i32(1089413), i32(1))
					v3 = v10
				l51:
					{
						t54 := int32(load32(m.memory[uint32(v3):]))
						v1 = t54
						if v1 == 0 {
							goto l47
						}
						t55 := int32(load32(m.memory[uint32(v3+i32(4)):]))
						v5 = t55
						t56 := int32(load32(m.memory[uint32(v5+i32(-4)):]))
						v6 = t56
						v7 = v6 & i32(-8)
						t57 := v7
						v6 = v6 & i32(3)
						p58 := i32(8)
						if v6 != 0 {
							p58 = i32(4)
						}
						if uint32(t57) < uint32(p58+v1) {
							m.fn2(i32(1273840), i32(46), i32(1273888))
							panic("unreachable")
						}
						if v6 == 0 {
							goto l49
						}
						if uint32(v7) > uint32(v1+i32(39)) {
							m.fn2(i32(1273904), i32(46), i32(1273952))
							panic("unreachable")
						}
					l49:
						m.fn1(v5)
					}
				l47:
					v3 = v3 + i32(12)
					v4 = v4 + i32(-1)
					if v4 != 0 {
						goto l51
					}
					t59 := int32(load32(m.memory[uint32(v10+i32(-4)):]))
					v3 = t59
					v1 = v3 & i32(-8)
					t60 := v1
					v3 = v3 & i32(3)
					p61 := i32(8)
					if v3 != 0 {
						p61 = i32(4)
					}
					if uint32(t60) < uint32(p61+v9) {
						m.fn2(i32(1273840), i32(46), i32(1273888))
						panic("unreachable")
					}
					if v3 == 0 {
						goto l53
					}
					if uint32(v1) > uint32(v9+i32(39)) {
						m.fn2(i32(1273904), i32(46), i32(1273952))
						panic("unreachable")
					}
				l53:
					m.fn1(v10)
					goto l10
				}
				store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
				store64(m.memory[uint32(v0):], uint64(i64(0x100000000)))
				goto l10
			case 5:
				{
					{
						t62 := int32(load32(m.memory[int64(uint32(v1))+12:]))
						v3 = t62
						if v3 != 0 {
							goto l55
						}
						v1 = i32(1)
						goto l56
					}
				l55:
					t63 := int32(load32(m.memory[int64(uint32(v1))+8:]))
					v4 = t63
					t64 := m.fn11(v3)
					v1 = t64
					if v1 == 0 {
						m.fn7(i32(1), v3)
						panic("unreachable")
					}
					if v3 == 0 {
						goto l56
					}
					memory_copy(m.memory, uint32(v1), uint32(v4), uint32(v3))
				}
			l56:
				store32(m.memory[int64(uint32(v0))+8:], uint32(v3))
				store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
				store32(m.memory[uint32(v0):], uint32(v3))
				goto l10
			case 6:
				store32(m.memory[int64(uint32(v0))+8:], uint32(i32(0)))
				store64(m.memory[uint32(v0):], uint64(i64(0x100000000)))
				goto l10
			case 1:
				v1 = v1 + i32(4)
				fallthrough
			default:
				t65 := int32(load32(m.memory[int64(uint32(v1))+8:]))
				v3 = t65
				t66 := int32(load32(m.memory[int64(uint32(v1))+4:]))
				v1 = t66
				store32(m.memory[int64(uint32(v2))+44:], uint32(i32(0)))
				store64(m.memory[int64(uint32(v2))+36:], uint64(i64(0x100000000)))
				m.fn466(v1, v3, v2+i32(36))
				t67 := int32(load32(m.memory[int64(uint32(v2))+44:]))
				store32(m.memory[int64(uint32(v0))+8:], uint32(t67))
				t68 := int64(load64(m.memory[int64(uint32(v2))+36:]))
				store64(m.memory[uint32(v0):], uint64(t68))
				goto l10
			}
		}
	l30:
		t69 := int64(load64(m.memory[int64(uint32(v2))+28:]))
		v8 = t69
		{
			t70 := m.fn11(i32(48))
			v15 = t70
			if v15 == 0 {
				m.fn7(i32(4), i32(48))
				panic("unreachable")
			}
			store64(m.memory[int64(uint32(v15))+4:], uint64(v8))
			store32(m.memory[uint32(v15):], uint32(v1))
			store32(m.memory[int64(uint32(v2))+20:], uint32(i32(1)))
			store32(m.memory[int64(uint32(v2))+16:], uint32(v15))
			store32(m.memory[int64(uint32(v2))+12:], uint32(i32(4)))
			v16 = i32(1)
		l76:
			{
				if v7 == v13 {
					goto l59
				}
			l74:
				v9 = v7 + i32(20)
				{
					{
						t71 := int32(load32(m.memory[uint32(v7):]))
						if t71 != i32(-1) {
							goto l60
						}
						store32(m.memory[int64(uint32(v2))+24:], uint32(i32(-1)))
						goto l61
					}
				l60:
					{
						t72 := int32(load32(m.memory[int64(uint32(v7))+8:]))
						v4 = t72
						if v4 != 0 {
							goto l62
						}
						v3 = i32(0)
						store32(m.memory[int64(uint32(v2))+32:], uint32(i32(0)))
						store64(m.memory[int64(uint32(v2))+24:], uint64(i64(0x100000000)))
						goto l63
					}
				l62:
					t73 := int32(load32(m.memory[int64(uint32(v7))+4:]))
					v1 = t73
					v14 = v4 * i32(12)
					t74 := m.fn11(v14)
					v10 = t74
					v3 = v10
					v6 = v4
					if v10 == 0 {
						m.fn7(i32(4), v14)
						panic("unreachable")
					}
				l65:
					{
						m.fn706(v2+i32(36), v1)
						t75 := int32(load32(m.memory[int64(uint32(v2))+44:]))
						store32(m.memory[int64(uint32(v3))+8:], uint32(t75))
						t76 := int64(load64(m.memory[int64(uint32(v2))+36:]))
						store64(m.memory[uint32(v3):], uint64(t76))
						v3 = v3 + i32(12)
						v1 = v1 + i32(32)
						v6 = v6 + i32(-1)
						if v6 != 0 {
							goto l65
						}
					}
					m.fn209(v2+i32(24), v10, v4, i32(1089413), i32(1))
					v3 = v10
				l70:
					{
						t77 := int32(load32(m.memory[uint32(v3):]))
						v1 = t77
						if v1 == 0 {
							goto l66
						}
						t78 := int32(load32(m.memory[uint32(v3+i32(4)):]))
						v5 = t78
						t79 := int32(load32(m.memory[uint32(v5+i32(-4)):]))
						v6 = t79
						v7 = v6 & i32(-8)
						t80 := v7
						v6 = v6 & i32(3)
						p81 := i32(8)
						if v6 != 0 {
							p81 = i32(4)
						}
						if uint32(t80) < uint32(p81+v1) {
							m.fn2(i32(1273840), i32(46), i32(1273888))
							panic("unreachable")
						}
						if v6 == 0 {
							goto l68
						}
						if uint32(v7) > uint32(v1+i32(39)) {
							m.fn2(i32(1273904), i32(46), i32(1273952))
							panic("unreachable")
						}
					l68:
						m.fn1(v5)
					}
				l66:
					v3 = v3 + i32(12)
					v4 = v4 + i32(-1)
					if v4 != 0 {
						goto l70
					}
					t82 := int32(load32(m.memory[uint32(v10+i32(-4)):]))
					v3 = t82
					v1 = v3 & i32(-8)
					t83 := v1
					v3 = v3 & i32(3)
					p84 := i32(8)
					if v3 != 0 {
						p84 = i32(4)
					}
					if uint32(t83) < uint32(p84+v14) {
						m.fn2(i32(1273840), i32(46), i32(1273888))
						panic("unreachable")
					}
					if v3 == 0 {
						goto l72
					}
					if uint32(v1) > uint32(v14+i32(39)) {
						m.fn2(i32(1273904), i32(46), i32(1273952))
						panic("unreachable")
					}
				l72:
					m.fn1(v10)
					t85 := int32(load32(m.memory[int64(uint32(v2))+24:]))
					v3 = t85
					if v3 != i32(-1) {
						goto l63
					}
				}
			l61:
				v7 = v9
				if v9 != v13 {
					goto l74
				}
			l59:
				{
					if v11 == v12 {
						t88 := int32(load32(m.memory[int64(uint32(v2))+12:]))
						v10 = t88
						t89 := int32(load32(m.memory[int64(uint32(v2))+16:]))
						t90 := v0
						v7 = t89
						m.fn209(t90, v7, v16, i32(1089413), i32(1))
						v3 = v7
					l81:
						{
							t91 := int32(load32(m.memory[uint32(v3):]))
							v1 = t91
							if v1 == 0 {
								goto l77
							}
							t92 := int32(load32(m.memory[uint32(v3+i32(4)):]))
							v6 = t92
							t93 := int32(load32(m.memory[uint32(v6+i32(-4)):]))
							v4 = t93
							v5 = v4 & i32(-8)
							t94 := v5
							v4 = v4 & i32(3)
							p95 := i32(8)
							if v4 != 0 {
								p95 = i32(4)
							}
							if uint32(t94) < uint32(p95+v1) {
								m.fn2(i32(1273840), i32(46), i32(1273888))
								panic("unreachable")
							}
							if v4 == 0 {
								goto l79
							}
							if uint32(v5) > uint32(v1+i32(39)) {
								m.fn2(i32(1273904), i32(46), i32(1273952))
								panic("unreachable")
							}
						l79:
							m.fn1(v6)
						}
					l77:
						v3 = v3 + i32(12)
						v16 = v16 + i32(-1)
						if v16 != 0 {
							goto l81
						}
						if v10 == 0 {
							goto l10
						}
						t96 := int32(load32(m.memory[uint32(v7+i32(-4)):]))
						v3 = t96
						v1 = v3 & i32(-8)
						t97 := v1
						v3 = v3 & i32(3)
						p98 := i32(8)
						if v3 != 0 {
							p98 = i32(4)
						}
						v4 = v10 * i32(12)
						if uint32(t97) < uint32(p98+v4) {
							m.fn2(i32(1273840), i32(46), i32(1273888))
							panic("unreachable")
						}
						if v3 == 0 {
							goto l83
						}
						if uint32(v1) > uint32(v4+i32(39)) {
							m.fn2(i32(1273904), i32(46), i32(1273952))
							panic("unreachable")
						}
					l83:
						m.fn1(v7)
						goto l10
					}
					t86 := int32(load32(m.memory[uint32(v11+i32(4)):]))
					v7 = t86
					t87 := int32(load32(m.memory[uint32(v11+i32(8)):]))
					v13 = v7 + t87*i32(20)
					v11 = v11 + i32(12)
					goto l76
				}
			l63:
				t99 := int64(load64(m.memory[int64(uint32(v2))+28:]))
				v8 = t99
				{
					t100 := int32(load32(m.memory[int64(uint32(v2))+12:]))
					if v16 != t100 {
						goto l85
					}
					m.fn203(v2+i32(12), v16, i32(1), i32(4), i32(12))
					t101 := int32(load32(m.memory[int64(uint32(v2))+16:]))
					v15 = t101
				}
			l85:
				v1 = v15 + v16*i32(12)
				store64(m.memory[int64(uint32(v1))+4:], uint64(v8))
				store32(m.memory[uint32(v1):], uint32(v3))
				t102 := v2
				v16 = v16 + i32(1)
				store32(m.memory[int64(uint32(t102))+20:], uint32(v16))
				v7 = v9
				goto l76
			}
		}
	}
l10:
	m.g0 = v2 + i32(48)
}
func (m *Module) fn707(v0 int32) {
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
func (m *Module) fn708(v0, v1, v2, v3 int32) {
	var v4, v5, v6, v7 int32
	var v8 int64
	t0 := m.g0
	v4 = t0 - i32(96)
	m.g0 = v4
	if v3 != 0 {
		goto l0
	}
	store32(m.memory[uint32(v0):], uint32(i32(-1)))
	goto l1
l0:
	{
		{
			{
				t1 := int32(m.memory[uint32(v2)])
				if t1 == i32(35) {
					m.fn206(v4+i32(80), v2+i32(1), v3+i32(-1))
					t16 := int32(load32(m.memory[int64(uint32(v1))+4:]))
					v5 = t16
					t17 := int32(load32(m.memory[int64(uint32(v4))+88:]))
					v2 = t17
					t18 := int32(load32(m.memory[int64(uint32(v4))+84:]))
					v3 = t18
					t19 := int32(load32(m.memory[int64(uint32(v1))+8:]))
					t20 := v4
					v1 = t19
					store32(m.memory[int64(uint32(t20))+68:], uint32(v1))
					store32(m.memory[int64(uint32(v4))+64:], uint32(v5))
					{
						if v2 != 0 {
							goto l7
						}
						if v1 <= i32(-1) {
							goto l8
						}
						{
							if v1 != 0 {
								goto l9
							}
							v2 = i32(1)
							goto l10
						l9:
							t21 := m.fn11(v1)
							v2 = t21
							if v2 == 0 {
								m.fn7(i32(1), v1)
								panic("unreachable")
							}
							if v1 == 0 {
								goto l10
							}
							memory_copy(m.memory, uint32(v2), uint32(v5), uint32(v1))
						}
					l10:
						store32(m.memory[int64(uint32(v4))+20:], uint32(v1))
						store32(m.memory[int64(uint32(v4))+16:], uint32(v2))
						store32(m.memory[int64(uint32(v4))+12:], uint32(v1))
						goto l12
					l7:
						store32(m.memory[int64(uint32(v4))+76:], uint32(v2))
						store32(m.memory[int64(uint32(v4))+72:], uint32(v3))
						t22 := v4
						v8 = int64(uint32(i32(1))) << 32
						store64(m.memory[int64(uint32(t22))+32:], uint64(v8|int64(uint32(v4+i32(72)))))
						store64(m.memory[int64(uint32(v4))+24:], uint64(v8|int64(uint32(v4+i32(64)))))
						m.fn14(v4+i32(12), i32(0x1000a1), v4+i32(24))
					}
				l12:
					store32(m.memory[uint32(v0):], uint32(i32(2)))
					t23 := int64(load64(m.memory[int64(uint32(v4))+12:]))
					store64(m.memory[int64(uint32(v0))+4:], uint64(t23))
					t24 := int32(load32(m.memory[int64(uint32(v4))+20:]))
					store32(m.memory[int64(uint32(v0))+12:], uint32(t24))
					t25 := int32(load32(m.memory[int64(uint32(v4))+80:]))
					v0 = t25
					if v0 == 0 {
						goto l1
					}
					t26 := int32(load32(m.memory[uint32(v3+i32(-4)):]))
					v2 = t26
					v1 = v2 & i32(-8)
					t27 := v1
					v2 = v2 & i32(3)
					p28 := i32(8)
					if v2 != 0 {
						p28 = i32(4)
					}
					if uint32(t27) < uint32(p28+v0) {
						m.fn2(i32(1273840), i32(46), i32(1273888))
						panic("unreachable")
					}
					if v2 == 0 {
						goto l14
					}
					if uint32(v1) > uint32(v0+i32(39)) {
						m.fn2(i32(1273904), i32(46), i32(1273952))
						panic("unreachable")
					}
				l14:
					m.fn1(v3)
					goto l1
				}
				t2 := m.fn583(v2, v3)
				if t2 != 0 {
					if v3 <= i32(-1) {
						goto l8
					}
					t29 := m.fn11(v3)
					v1 = t29
					if v1 != 0 {
						if v3 == 0 {
							goto l25
						}
						memory_copy(m.memory, uint32(v1), uint32(v2), uint32(v3))
					l25:
						store32(m.memory[int64(uint32(v0))+12:], uint32(v3))
						store32(m.memory[int64(uint32(v0))+8:], uint32(v1))
						store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
						store32(m.memory[uint32(v0):], uint32(i32(0)))
						goto l1
					}
					m.fn7(i32(1), v3)
					panic("unreachable")
				}
				t3 := int32(load32(m.memory[int64(uint32(v1))+4:]))
				t4 := int32(load32(m.memory[int64(uint32(v1))+8:]))
				m.fn155(v4+i32(24), t3, t4, v2, v3)
				t5 := int32(load32(m.memory[int64(uint32(v4))+24:]))
				v5 = t5
				if v5 != 0 {
					goto l4
				}
				t6 := int32(load32(m.memory[int64(uint32(v1))+20:]))
				t7 := int32(load32(m.memory[int64(uint32(v4))+32:]))
				v6 = t7
				t8 := int32(load32(m.memory[int64(uint32(v4))+36:]))
				t9 := v6
				v1 = t8
				t10 := m.fn385(t6, t9, v1)
				if t10 == 0 {
					goto l4
				}
				t11 := int32(load32(m.memory[int64(uint32(v4))+48:]))
				v2 = t11
				t12 := int32(load32(m.memory[int64(uint32(v4))+44:]))
				v7 = t12
				t13 := int32(load32(m.memory[int64(uint32(v4))+28:]))
				v5 = t13
				t14 := int32(load32(m.memory[int64(uint32(v4))+40:]))
				v3 = t14
				store32(m.memory[int64(uint32(v4))+68:], uint32(v1))
				store32(m.memory[int64(uint32(v4))+64:], uint32(v6))
				if v3 == i32(-1) {
					goto l5
				}
				if v2 == 0 {
					goto l5
				}
				store32(m.memory[int64(uint32(v4))+76:], uint32(v2))
				store32(m.memory[int64(uint32(v4))+72:], uint32(v7))
				t15 := v4
				v8 = int64(uint32(i32(1))) << 32
				store64(m.memory[int64(uint32(t15))+88:], uint64(v8|int64(uint32(v4+i32(72)))))
				store64(m.memory[int64(uint32(v4))+80:], uint64(v8|int64(uint32(v4+i32(64)))))
				m.fn14(v4+i32(52), i32(0x1000a1), v4+i32(80))
				goto l6
			}
		l4:
			if v3 <= i32(-1) {
				goto l8
			}
			t30 := m.fn11(v3)
			v1 = t30
			if v1 != 0 {
				if v3 == 0 {
					goto l22
				}
				memory_copy(m.memory, uint32(v1), uint32(v2), uint32(v3))
			l22:
				store32(m.memory[int64(uint32(v0))+12:], uint32(v3))
				store32(m.memory[int64(uint32(v0))+8:], uint32(v1))
				store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
				store32(m.memory[uint32(v0):], uint32(i32(1)))
				if v5 != 0 {
					goto l23
				}
				{
					t34 := int32(load32(m.memory[int64(uint32(v4))+28:]))
					v3 = t34
					if v3 == 0 {
						goto l24
					}
					t35 := int32(load32(m.memory[int64(uint32(v4))+32:]))
					m.fn21(t35, v3, i32(1))
				}
			l24:
				t36 := int32(load32(m.memory[int64(uint32(v4))+40:]))
				v3 = t36
				if v3 == i32(-1) {
					goto l1
				}
				if v3 == 0 {
					goto l1
				}
				t37 := int32(load32(m.memory[int64(uint32(v4))+44:]))
				m.fn21(t37, v3, i32(1))
				goto l1
			}
			m.fn7(i32(1), v3)
			panic("unreachable")
		}
	l5:
		if v1 <= i32(-1) {
			goto l8
		}
		{
			if v1 != 0 {
				goto l18
			}
			v2 = i32(1)
			goto l19
		l18:
			t31 := m.fn11(v1)
			v2 = t31
			if v2 == 0 {
				m.fn7(i32(1), v1)
				panic("unreachable")
			}
			if v1 == 0 {
				goto l19
			}
			memory_copy(m.memory, uint32(v2), uint32(v6), uint32(v1))
		}
	l19:
		store32(m.memory[int64(uint32(v4))+60:], uint32(v1))
		store32(m.memory[int64(uint32(v4))+56:], uint32(v2))
		store32(m.memory[int64(uint32(v4))+52:], uint32(v1))
	l6:
		store32(m.memory[uint32(v0):], uint32(i32(2)))
		t32 := int64(load64(m.memory[int64(uint32(v4))+52:]))
		store64(m.memory[int64(uint32(v0))+4:], uint64(t32))
		t33 := int32(load32(m.memory[int64(uint32(v4))+60:]))
		store32(m.memory[int64(uint32(v0))+12:], uint32(t33))
		if v5 == 0 {
			goto l21
		}
		m.fn21(v6, v5, i32(1))
	l21:
		if uint32(v3+i32(-1)) > uint32(i32(-3)) {
			goto l1
		}
		m.fn21(v7, v3, i32(1))
		goto l1
	}
l8:
	m.fn12()
	panic("unreachable")
l23:
	m.fn149(v4 + i32(28))
l1:
	m.g0 = v4 + i32(96)
}
func (m *Module) fn709(v0, v1, v2, v3 int32) {
	var v4, v5, v6, v7, v8, v9 int32
	var v10 int64
	var v11 int32
	t0 := m.g0
	v4 = t0 - i32(64)
	m.g0 = v4
	if v3 != 0 {
		goto l0
	}
	store64(m.memory[uint32(v0):], uint64(i64(-1)))
	goto l1
l0:
	{
		t1 := m.fn583(v2, v3)
		if t1 != 0 {
			if v3 <= i32(-1) {
				m.fn12()
				panic("unreachable")
			}
			t5 := m.fn11(v3)
			v1 = t5
			if v1 != 0 {
				if v3 == 0 {
					goto l22
				}
				memory_copy(m.memory, uint32(v1), uint32(v2), uint32(v3))
			l22:
				store32(m.memory[int64(uint32(v0))+12:], uint32(v3))
				store32(m.memory[int64(uint32(v0))+8:], uint32(v1))
				store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
				store32(m.memory[uint32(v0):], uint32(i32(-1)))
				goto l1
			}
			m.fn7(i32(1), v3)
			panic("unreachable")
		}
		t2 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		t3 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		m.fn155(v4, t2, t3, v2, v3)
		t4 := int32(load32(m.memory[uint32(v4):]))
		if t4 == 0 {
			{
				{
					t6 := int32(load32(m.memory[int64(uint32(v1))+12:]))
					v3 = t6
					t7 := int32(load32(m.memory[uint32(v3):]))
					if t7 != 0 {
						m.fn361(i32(1079216))
						panic("unreachable")
					}
					t8 := int32(load32(m.memory[int64(uint32(v4))+20:]))
					v5 = t8
					t9 := int32(load32(m.memory[int64(uint32(v4))+16:]))
					v6 = t9
					t10 := int32(load32(m.memory[int64(uint32(v4))+12:]))
					v7 = t10
					t11 := int32(load32(m.memory[int64(uint32(v4))+8:]))
					v2 = t11
					t12 := int32(load32(m.memory[int64(uint32(v4))+4:]))
					v8 = t12
					store32(m.memory[uint32(v3):], uint32(i32(-1)))
					m.fn148(v4+i32(28), v3+i32(8), v2, v7)
					{
						t13 := int32(load32(m.memory[int64(uint32(v4))+28:]))
						v9 = t13
						if v9 == i32(-1) {
							goto l7
						}
						if v9 == i32(-0x7ffffffd) {
							goto l8
						}
						t14 := int64(load64(m.memory[int64(uint32(v4))+28:]))
						v10 = t14
						store32(m.memory[int64(uint32(v4))+32:], uint32(i32(0)))
						t15 := int64(load64(m.memory[int64(uint32(v4))+44:]))
						store64(m.memory[int64(uint32(v4))+16:], uint64(t15))
						t16 := int64(load64(m.memory[int64(uint32(v4))+36:]))
						store64(m.memory[int64(uint32(v4))+8:], uint64(t16))
						store64(m.memory[uint32(v4):], uint64(v10))
						m.fn149(v4)
					}
				l7:
					t17 := int32(load32(m.memory[int64(uint32(v4))+32:]))
					v9 = t17
					if v9 == 0 {
						store64(m.memory[uint32(v0):], uint64(i64(-1)))
						t39 := int32(load32(m.memory[uint32(v3):]))
						store32(m.memory[uint32(v3):], uint32(t39+i32(1)))
						if v8 == 0 {
							goto l20
						}
						m.fn21(v2, v8, i32(1))
						goto l20
					}
					t18 := int32(load32(m.memory[int64(uint32(v4))+36:]))
					v11 = t18
					m.fn710(v4+i32(52), v2, v7)
					t19 := int32(load32(m.memory[int64(uint32(v1))+16:]))
					v1 = t19
					t20 := int32(load32(m.memory[uint32(v1):]))
					if t20 != 0 {
						m.fn361(i32(1079200))
						panic("unreachable")
					}
					store32(m.memory[uint32(v1):], uint32(i32(-1)))
					store32(m.memory[int64(uint32(v4))+36:], uint32(v7))
					store32(m.memory[int64(uint32(v4))+32:], uint32(v2))
					store32(m.memory[int64(uint32(v4))+28:], uint32(v8))
					m.fn451(v4, v1+i32(8), v4+i32(52), v4+i32(28), v9+i32(8), v11)
					t21 := int32(load32(m.memory[int64(uint32(v4))+4:]))
					v2 = t21
					t22 := int32(load32(m.memory[uint32(v4):]))
					v7 = t22
					if v7 == i32(-1) {
						store32(m.memory[int64(uint32(v0))+8:], uint32(v2))
						store64(m.memory[uint32(v0):], uint64(i64(-0x7fffffff00000001)))
						t40 := int32(load32(m.memory[uint32(v1):]))
						store32(m.memory[uint32(v1):], uint32(t40+i32(1)))
						t41 := int32(load32(m.memory[uint32(v9):]))
						t42 := v9
						v0 = t41 + i32(-1)
						store32(m.memory[uint32(t42):], uint32(v0))
						if v0 != 0 {
							goto l21
						}
						m.fn152(v9, v11)
					l21:
						t43 := int32(load32(m.memory[uint32(v3):]))
						store32(m.memory[uint32(v3):], uint32(t43+i32(1)))
						goto l20
					}
					t23 := int64(load64(m.memory[int64(uint32(v4))+16:]))
					store64(m.memory[int64(uint32(v0))+16:], uint64(t23))
					t24 := int64(load64(m.memory[int64(uint32(v4))+8:]))
					store64(m.memory[int64(uint32(v0))+8:], uint64(t24))
					store32(m.memory[int64(uint32(v0))+4:], uint32(v2))
					store32(m.memory[uint32(v0):], uint32(v7))
					t25 := int32(load32(m.memory[uint32(v1):]))
					store32(m.memory[uint32(v1):], uint32(t25+i32(1)))
					t26 := int32(load32(m.memory[uint32(v9):]))
					t27 := v9
					v0 = t26 + i32(-1)
					store32(m.memory[uint32(t27):], uint32(v0))
					if v0 != 0 {
						goto l12
					}
					m.fn152(v9, v11)
				l12:
					t28 := int32(load32(m.memory[uint32(v3):]))
					store32(m.memory[uint32(v3):], uint32(t28+i32(1)))
					goto l13
				}
			l8:
				t29 := int64(load64(m.memory[int64(uint32(v4))+40:]))
				store64(m.memory[int64(uint32(v0))+12:], uint64(t29))
				t30 := int32(load32(m.memory[int64(uint32(v4))+48:]))
				store32(m.memory[int64(uint32(v0))+20:], uint32(t30))
				t31 := int64(load64(m.memory[int64(uint32(v4))+32:]))
				store64(m.memory[int64(uint32(v0))+4:], uint64(t31))
				store32(m.memory[uint32(v0):], uint32(i32(-0x7ffffffd)))
				t32 := int32(load32(m.memory[uint32(v3):]))
				store32(m.memory[uint32(v3):], uint32(t32+i32(1)))
				if v8 == 0 {
					goto l13
				}
				t33 := int32(load32(m.memory[uint32(v2+i32(-4)):]))
				v3 = t33
				v0 = v3 & i32(-8)
				t34 := v0
				v3 = v3 & i32(3)
				p35 := i32(8)
				if v3 != 0 {
					p35 = i32(4)
				}
				if uint32(t34) < uint32(p35+v8) {
					m.fn2(i32(1273840), i32(46), i32(1273888))
					panic("unreachable")
				}
				if v3 == 0 {
					goto l15
				}
				if uint32(v0) > uint32(v8+i32(39)) {
					m.fn2(i32(1273904), i32(46), i32(1273952))
					panic("unreachable")
				}
			l15:
				m.fn1(v2)
			}
		l13:
			if uint32(v6+i32(-1)) > uint32(i32(-3)) {
				goto l1
			}
			t36 := int32(load32(m.memory[uint32(v5+i32(-4)):]))
			v3 = t36
			v0 = v3 & i32(-8)
			t37 := v0
			v3 = v3 & i32(3)
			p38 := i32(8)
			if v3 != 0 {
				p38 = i32(4)
			}
			if uint32(t37) < uint32(p38+v6) {
				m.fn2(i32(1273840), i32(46), i32(1273888))
				panic("unreachable")
			}
			if v3 == 0 {
				goto l18
			}
			if uint32(v0) > uint32(v6+i32(39)) {
				m.fn2(i32(1273904), i32(46), i32(1273952))
				panic("unreachable")
			}
		l18:
			m.fn1(v5)
			goto l1
		}
		m.fn149(v4 + i32(4))
		store64(m.memory[uint32(v0):], uint64(i64(-1)))
		goto l1
	}
l20:
	if uint32(v6+i32(-1)) > uint32(i32(-3)) {
		goto l1
	}
	m.fn21(v5, v6, i32(1))
l1:
	m.g0 = v4 + i32(64)
}
func (m *Module) fn710(v0, v1, v2 int32) {
	var v3, v4, v5, v6, v7, v8 int32
	t0 := m.g0
	v3 = t0 - i32(48)
	m.g0 = v3
	store32(m.memory[int64(uint32(v3))+36:], uint32(v2))
	v4 = i32(0)
	store32(m.memory[int64(uint32(v3))+32:], uint32(i32(0)))
	store32(m.memory[int64(uint32(v3))+28:], uint32(v2))
	store32(m.memory[int64(uint32(v3))+24:], uint32(v1))
	store32(m.memory[int64(uint32(v3))+20:], uint32(i32(46)))
	store32(m.memory[int64(uint32(v3))+40:], uint32(i32(46)))
	v5 = i32(1)
	m.memory[int64(uint32(v3))+44] = byte(i32(1))
	m.fn158(v3+i32(8), v3+i32(20))
	{
		{
			t1 := int32(load32(m.memory[int64(uint32(v3))+8:]))
			if t1 != i32(1) {
				goto l0
			}
			t2 := int32(load32(m.memory[int64(uint32(v3))+16:]))
			t3 := v2
			v6 = t2
			v7 = t3 - v6
			if v7 <= i32(-1) {
				m.fn12()
				panic("unreachable")
			}
			if v7 == 0 {
				goto l0
			}
			t4 := m.fn11(v7)
			v5 = t4
			if v5 == 0 {
				m.fn7(i32(1), v7)
				panic("unreachable")
			}
			if v7 == 0 {
				goto l3
			}
			memory_copy(m.memory, uint32(v5), uint32(v1+v6), uint32(v7))
		l3:
			v1 = i32(0)
			{
				if v2 == v6+i32(1) {
					goto l4
				}
				v8 = v7 & i32(1)
				v6 = v7 & i32(0x7ffffffe)
				v1 = i32(0)
			l5:
				{
					v2 = v5 + v1
					t5 := int32(m.memory[uint32(v2)])
					t6 := v2
					v4 = t5
					p7 := i32(0)
					if uint32((v4+i32(-65))&i32(255)) < uint32(i32(26)) {
						p7 = i32(32)
					}
					m.memory[uint32(t6)] = byte(p7 | v4)
					v2 = v2 + i32(1)
					t8 := int32(m.memory[uint32(v2)])
					t9 := v2
					v2 = t8
					p10 := i32(0)
					if uint32((v2+i32(-65))&i32(255)) < uint32(i32(26)) {
						p10 = i32(32)
					}
					m.memory[uint32(t9)] = byte(p10 | v2)
					t11 := v6
					v1 = v1 + i32(2)
					if t11 != v1 {
						goto l5
					}
				}
				if v8 == 0 {
					goto l6
				}
			l4:
				v1 = v5 + v1
				t12 := int32(m.memory[uint32(v1)])
				t13 := v1
				v1 = t12
				p14 := i32(0)
				if uint32((v1+i32(-65))&i32(255)) < uint32(i32(26)) {
					p14 = i32(32)
				}
				m.memory[uint32(t13)] = byte(p14 | v1)
			}
		l6:
			v4 = v7
			switch v7 + i32(-3) {
			default:
				goto l0
			case 0:
				v1 = i32(9)
				v4 = i32(3)
				{
					t15 := int32(load16(m.memory[uint32(v5):]))
					t16 := t15 ^ i32(28272)
					v2 = v5 + i32(2)
					t17 := int32(m.memory[uint32(v2)])
					if (t16|(t17^i32(103)))&i32(0xffff) != 0 {
						t18 := int32(load16(m.memory[uint32(v5):]))
						t19 := int32(m.memory[uint32(v2)])
						if (t18^i32(28778)|(t19^i32(103)))&i32(0xffff) != 0 {
							t21 := int32(load16(m.memory[uint32(v5):]))
							t22 := t21 ^ i32(26983)
							v2 = v5 + i32(2)
							t23 := int32(m.memory[uint32(v2)])
							if (t22|(t23^i32(102)))&i32(0xffff) != 0 {
								t24 := int32(load16(m.memory[uint32(v5):]))
								t25 := int32(m.memory[uint32(v2)])
								if (t24^i32(28002)|(t25^i32(112)))&i32(0xffff) != 0 {
									t26 := int32(load16(m.memory[uint32(v5):]))
									t27 := t26 ^ i32(26996)
									v2 = v5 + i32(2)
									t28 := int32(m.memory[uint32(v2)])
									if (t27|(t28^i32(102)))&i32(0xffff) != 0 {
										t30 := int32(load16(m.memory[uint32(v5):]))
										t31 := int32(m.memory[uint32(v2)])
										if (t30^i32(30323)|(t31^i32(103)))&i32(0xffff) != 0 {
											t32 := int32(load16(m.memory[uint32(v5):]))
											t33 := t32 ^ i32(28005)
											v2 = v5 + i32(2)
											t34 := int32(m.memory[uint32(v2)])
											if (t33|(t34^i32(102)))&i32(0xffff) != 0 {
												v4 = i32(3)
												t35 := int32(load16(m.memory[uint32(v5):]))
												t36 := int32(m.memory[uint32(v2)])
												if (t35^i32(28023)|(t36^i32(102)))&i32(0xffff) != 0 {
													goto l0
												}
												v6 = i32(1074814)
												goto l10
											}
											v6 = i32(1074826)
											goto l10
										}
										v6 = i32(1075950)
										v1 = i32(13)
										goto l10
									}
									v6 = i32(1075963)
									v1 = i32(10)
									goto l10
								}
								v6 = i32(1075973)
								goto l10
							}
							v6 = i32(1075982)
							goto l10
						}
						v6 = i32(1074838)
						v1 = i32(10)
						goto l10
					}
					v6 = i32(1074851)
					goto l10
				}
			case 1:
				v1 = i32(10)
				v4 = i32(4)
				t20 := int32(load32(m.memory[uint32(v5):]))
				if t20 != i32(1734701162) {
					t29 := int32(load32(m.memory[uint32(v5):]))
					if t29 != i32(0x66666974) {
						v4 = i32(4)
						t37 := int32(load32(m.memory[uint32(v5):]))
						if t37 != i32(1885496695) {
							goto l0
						}
						v6 = i32(1075940)
						goto l10
					}
					v6 = i32(1075963)
					goto l10
				}
				v6 = i32(1074838)
				goto l10
			}
		}
	l0:
		v6 = i32(1075916)
		v1 = i32(24)
	l10:
		t38 := m.fn11(v1)
		v2 = t38
		if v2 == 0 {
			m.fn7(i32(1), v1)
			panic("unreachable")
		}
		if v1 == 0 {
			goto l20
		}
		memory_copy(m.memory, uint32(v2), uint32(v6), uint32(v1))
	l20:
		store32(m.memory[int64(uint32(v0))+8:], uint32(v1))
		store32(m.memory[int64(uint32(v0))+4:], uint32(v2))
		store32(m.memory[uint32(v0):], uint32(v1))
		{
			if v4 == 0 {
				goto l21
			}
			t39 := int32(load32(m.memory[uint32(v5+i32(-4)):]))
			v1 = t39
			v2 = v1 & i32(-8)
			t40 := v2
			v1 = v1 & i32(3)
			p41 := i32(8)
			if v1 != 0 {
				p41 = i32(4)
			}
			if uint32(t40) < uint32(p41+v4) {
				m.fn2(i32(1273840), i32(46), i32(1273888))
				panic("unreachable")
			}
			if v1 == 0 {
				goto l23
			}
			if uint32(v2) > uint32(v4+i32(39)) {
				m.fn2(i32(1273904), i32(46), i32(1273952))
				panic("unreachable")
			}
		l23:
			m.fn1(v5)
		}
	l21:
		m.g0 = v3 + i32(48)
		return
	}
}
func (m *Module) fn711(v0, v1, v2, v3 int32) {
	var v4, v5 int32
	var v6 int64
	t0 := m.g0
	v4 = t0 - i32(32)
	m.g0 = v4
	t1 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	v5 = t1
	t2 := int32(load32(m.memory[int64(uint32(v1))+8:]))
	t3 := v4
	v1 = t2
	store32(m.memory[int64(uint32(t3))+4:], uint32(v1))
	store32(m.memory[uint32(v4):], uint32(v5))
	{
		if v3 != 0 {
			goto l0
		}
		if v1 <= i32(-1) {
			m.fn12()
			panic("unreachable")
		}
		{
			if v1 != 0 {
				goto l2
			}
			v3 = i32(1)
			goto l3
		l2:
			t4 := m.fn11(v1)
			v3 = t4
			if v3 == 0 {
				m.fn7(i32(1), v1)
				panic("unreachable")
			}
			if v1 == 0 {
				goto l3
			}
			memory_copy(m.memory, uint32(v3), uint32(v5), uint32(v1))
		}
	l3:
		store32(m.memory[int64(uint32(v0))+8:], uint32(v1))
		store32(m.memory[int64(uint32(v0))+4:], uint32(v3))
		store32(m.memory[uint32(v0):], uint32(v1))
		goto l5
	l0:
		store32(m.memory[int64(uint32(v4))+12:], uint32(v3))
		store32(m.memory[int64(uint32(v4))+8:], uint32(v2))
		t5 := v4
		v6 = int64(uint32(i32(1))) << 32
		store64(m.memory[int64(uint32(t5))+24:], uint64(v6|int64(uint32(v4+i32(8)))))
		store64(m.memory[int64(uint32(v4))+16:], uint64(v6|int64(uint32(v4))))
		m.fn14(v0, i32(0x1000a1), v4+i32(16))
	}
l5:
	m.g0 = v4 + i32(32)
}
func (m *Module) fn712(v0, v1, v2 int32) {
	var v3, v4, v5, v6, v7 int32
	t0 := m.g0
	v3 = t0 - i32(16)
	m.g0 = v3
	{
		t1 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		v4 = t1
		if v4 == 0 {
			goto l0
		}
		v5 = i32(3)
		t2 := int32(load32(m.memory[int64(uint32(v1))+4:]))
		v6 = t2
		{
			{
				{
					if uint32(v4) < uint32(i32(3)) {
						if v4 == i32(2) {
							goto l2
						}
						v4 = i32(1)
						goto l4
					}
					t3 := int32(load16(m.memory[uint32(v6):]))
					t4 := int32(m.memory[uint32(v6+i32(2))])
					if (t3^i32(48111)|(t4^i32(191)))&i32(0xffff) != 0 {
						goto l2
					}
					v0 = i32(1271548)
					goto l3
				}
			l2:
				v5 = i32(2)
				{
					t5 := int32(load16(m.memory[uint32(v6):]))
					if t5 != i32(65279) {
						goto l5
					}
					v0 = i32(1271552)
					goto l3
				}
			l5:
				t6 := int32(load16(m.memory[uint32(v6):]))
				v7 = t6
				if (v7<<8|int32(uint32(v7)>>8))&i32(0xffff) != i32(65279) {
					goto l4
				}
				v0 = i32(1271556)
			}
		l3:
			if uint32(v4) < uint32(v5) {
				m.fn127(v5, v4, v4, i32(1080316))
				panic("unreachable")
			}
			v6 = v6 + v5
			v4 = v4 - v5
			t7 := int32(load32(m.memory[uint32(v0):]))
			v0 = t7
		}
	l4:
		m.fn215(v3, v0, v6, v4)
		t8 := int32(load32(m.memory[int64(uint32(v3))+8:]))
		v4 = t8
		t9 := int32(load32(m.memory[int64(uint32(v3))+4:]))
		v6 = t9
		{
			{
				t10 := int32(load32(m.memory[uint32(v3):]))
				v5 = t10
				if v5 == i32(-1) {
					goto l7
				}
				v0 = v6
				goto l8
			}
		l7:
			if v4 <= i32(-1) {
				m.fn12()
				panic("unreachable")
			}
			if v4 != 0 {
				goto l10
			}
			v5 = i32(0)
			v0 = i32(1)
			v4 = i32(0)
			goto l8
		l10:
			t11 := m.fn11(v4)
			v0 = t11
			if v0 == 0 {
				m.fn7(i32(1), v4)
				panic("unreachable")
			}
			if v4 == 0 {
				goto l12
			}
			memory_copy(m.memory, uint32(v0), uint32(v6), uint32(v4))
		l12:
			v5 = v4
		}
	l8:
		{
			t12 := int32(load32(m.memory[int64(uint32(v2))+8:]))
			v6 = t12
			t13 := int32(load32(m.memory[uint32(v2):]))
			if v6 != t13 {
				goto l13
			}
			m.fn322(v2)
		}
	l13:
		store32(m.memory[int64(uint32(v1))+8:], uint32(i32(0)))
		store32(m.memory[int64(uint32(v2))+8:], uint32(v6+i32(1)))
		t14 := int32(load32(m.memory[int64(uint32(v2))+4:]))
		v1 = t14 + v6*i32(12)
		store32(m.memory[int64(uint32(v1))+8:], uint32(v4))
		store32(m.memory[int64(uint32(v1))+4:], uint32(v0))
		store32(m.memory[uint32(v1):], uint32(v5))
	}
l0:
	m.g0 = v3 + i32(16)
}
func (m *Module) fn713(v0, v1, v2, v3 int32) {
	var v4, v5, v6 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+20:]))
		v4 = t0
		if uint32(v4) >= uint32(v1) {
			goto l0
		}
		v5 = v4 * i32(12)
	l2:
		{
			{
				t1 := int32(load32(m.memory[int64(uint32(v0))+12:]))
				if v4 != t1 {
					goto l1
				}
				m.fn322(v0 + i32(12))
			}
		l1:
			t2 := v0
			v4 = v4 + i32(1)
			store32(m.memory[int64(uint32(t2))+20:], uint32(v4))
			t3 := int32(load32(m.memory[int64(uint32(v0))+16:]))
			v6 = t3 + v5
			store64(m.memory[uint32(v6):], uint64(i64(0x800000000)))
			store32(m.memory[uint32(v6+i32(8)):], uint32(i32(0)))
			v5 = v5 + i32(12)
			if v4 != v1 {
				goto l2
			}
		}
		v4 = v1
	}
l0:
	{
		t4 := int32(load32(m.memory[int64(uint32(v0))+32:]))
		v5 = t4
		if uint32(v5) >= uint32(v1) {
			goto l3
		}
		v4 = v5 << 4
		v6 = v0 + i32(24)
	l5:
		{
			{
				t5 := int32(load32(m.memory[int64(uint32(v0))+24:]))
				if v5 != t5 {
					goto l4
				}
				m.fn323(v6)
			}
		l4:
			t6 := v0
			v5 = v5 + i32(1)
			store32(m.memory[int64(uint32(t6))+32:], uint32(v5))
			t7 := int32(load32(m.memory[int64(uint32(v0))+28:]))
			store32(m.memory[uint32(t7+v4):], uint32(i32(0)))
			v4 = v4 + i32(16)
			if v5 != v1 {
				goto l5
			}
		}
		t8 := int32(load32(m.memory[int64(uint32(v0))+20:]))
		v4 = t8
		v5 = v1
	}
l3:
	{
		v1 = v1 + i32(-1)
		if uint32(v1) >= uint32(v4) {
			m.fn39(v1, v4, i32(1072224))
			panic("unreachable")
		}
		if uint32(v1) >= uint32(v5) {
			m.fn39(v1, v5, i32(1072240))
			panic("unreachable")
		}
		t9 := int32(load32(m.memory[int64(uint32(v0))+16:]))
		v6 = t9 + v1*i32(12)
		t10 := int32(load32(m.memory[int64(uint32(v0))+28:]))
		v0 = t10 + v1<<4
		if v2&i32(255) == i32(2) {
			m.fn433(v0, v6)
			t11 := int32(load32(m.memory[int64(uint32(v3))+8:]))
			v4 = t11
			v1 = v4 * i32(28)
			t12 := int32(load32(m.memory[int64(uint32(v3))+4:]))
			v5 = t12
			v0 = i32(0)
			{
			l10:
				{
					if v1 == v0 {
						if v4 == 0 {
							goto l12
						}
						v0 = v5
					l13:
						m.fn343(v0)
						v0 = v0 + i32(28)
						v4 = v4 + i32(-1)
						if v4 != 0 {
							goto l13
						}
					l12:
						{
							t20 := int32(load32(m.memory[uint32(v3):]))
							v0 = t20
							if v0 == 0 {
								return
							}
							t21 := int32(load32(m.memory[uint32(v5+i32(-4)):]))
							v4 = t21
							v1 = v4 & i32(-8)
							t22 := v1
							v4 = v4 & i32(3)
							p23 := i32(8)
							if v4 != 0 {
								p23 = i32(4)
							}
							v0 = v0 * i32(28)
							if uint32(t22) < uint32(p23+v0) {
								m.fn2(i32(1273840), i32(46), i32(1273888))
								panic("unreachable")
							}
							if v4 == 0 {
								goto l16
							}
							if uint32(v1) > uint32(v0+i32(39)) {
								m.fn2(i32(1273904), i32(46), i32(1273952))
								panic("unreachable")
							}
						l16:
							m.fn1(v5)
						}
						return
					}
					t13 := v5
					v0 = v0 + i32(28)
					t14 := m.fn317(t13 + v0 + i32(-28))
					if t14 != 0 {
						goto l10
					}
				}
				{
					t15 := int32(load32(m.memory[int64(uint32(v6))+8:]))
					v0 = t15
					t16 := int32(load32(m.memory[uint32(v6):]))
					if v0 != t16 {
						goto l11
					}
					m.fn321(v6)
				}
			l11:
				t17 := int32(load32(m.memory[int64(uint32(v6))+4:]))
				v4 = t17 + v0<<5
				t18 := int64(load64(m.memory[uint32(v3):]))
				store64(m.memory[int64(uint32(v4))+4:], uint64(t18))
				store32(m.memory[uint32(v4):], uint32(i32(-0x80000000)))
				t19 := int32(load32(m.memory[int64(uint32(v3))+8:]))
				store32(m.memory[int64(uint32(v4))+12:], uint32(t19))
				store32(m.memory[int64(uint32(v6))+8:], uint32(v0+i32(1)))
				return
			}
		}
		m.fn569(v0, v2&i32(1), v3, v6)
		return
	}
}
func (m *Module) fn714(v0, v1, v2 int32, v3 int64) int64 {
	var v4 int32
	var v5, v6 int64
	var v7, v8, v9, v10, v11, v12, v13 int32
	var v14 int64
	var v15 int32
	t0 := m.g0
	v4 = t0 - i32(96)
	m.g0 = v4
	t1 := int64(load64(m.memory[int64(uint32(v0))+16:]))
	t2 := int64(load64(m.memory[int64(uint32(v0))+24:]))
	t3 := m.fn100(t1, t2, v1)
	v5 = t3
	v6 = int64(uint64(v5)>>25) & i64(127) * i64(72340172838076673)
	p4 := i32(8)
	if uint32(v2) < uint32(i32(8)) {
		p4 = v2
	}
	v7 = p4
	t5 := int32(load32(m.memory[uint32(v0):]))
	v8 = t5
	v9 = i32(0)
	t6 := int32(load32(m.memory[int64(uint32(v0))+4:]))
	v10 = t6
	t7 := v10
	v11 = int32(v5)
	v12 = t7 & v11
	v13 = v12
	{
	l4:
		{
			t8 := int64(load64(m.memory[uint32(v8+v13):]))
			v14 = t8
			v5 = v14 ^ v6
			v5 = (v5 ^ i64(-1)) & (v5 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
			if v5 == 0 {
				goto l0
			}
		l2:
			{
				v15 = v8 + (i32(0)-(int32(uint32(int64(bits.TrailingZeros64(uint64(v5))))>>3)+v13)&v10)*i32(96)
				t9 := int32(load32(m.memory[uint32(v15+i32(-96)):]))
				if t9 == v1 {
					goto l1
				}
				v5 = (v5 + i64(-1)) & v5
				if !(v5 == 0) {
					goto l2
				}
			}
		}
	l0:
		{
			if !(v14&(v14<<1)&i64(-0x7f7f7f7f7f7f7f80) == 0) {
				goto l3
			}
			t10 := v13
			v9 = v9 + i32(8)
			v13 = (t10 + v9) & v10
			goto l4
		}
	l3:
		{
			t11 := int32(load32(m.memory[int64(uint32(v0))+8:]))
			if t11 != 0 {
				goto l5
			}
			_ = m.fn107(v0, v0+i32(16))
			t13 := int32(load32(m.memory[int64(uint32(v0))+4:]))
			v10 = t13
			v12 = v10 & v11
			t14 := int32(load32(m.memory[uint32(v0):]))
			v8 = t14
		}
	l5:
		memory_zero(m.memory, uint32(v4+i32(8)), uint32(i32(81)))
		{
			t15 := int64(load64(m.memory[uint32(v8+v12):]))
			v5 = t15 & i64(-0x7f7f7f7f7f7f7f80)
			if v5 != i64(0) {
				goto l6
			}
			v15 = i32(8)
		l7:
			{
				v13 = v12 + v15
				v15 = v15 + i32(8)
				t16 := v8
				v12 = v13 & v10
				t17 := int64(load64(m.memory[uint32(t16+v12):]))
				v5 = t17 & i64(-0x7f7f7f7f7f7f7f80)
				if v5 == 0 {
					goto l7
				}
			}
		}
	l6:
		{
			t18 := v8
			v15 = (int32(uint32(int64(bits.TrailingZeros64(uint64(v5))))>>3) + v12) & v10
			t19 := int32(int8(m.memory[uint32(t18+v15)]))
			v13 = t19
			if v13 < i32(0) {
				goto l8
			}
			t20 := int64(load64(m.memory[uint32(v8):]))
			t21 := v8
			v15 = int32(uint32(int64(bits.TrailingZeros64(uint64(t20&i64(-0x7f7f7f7f7f7f7f80))))) >> 3)
			t22 := int32(m.memory[uint32(t21+v15)])
			v13 = t22
		}
	l8:
		t23 := v8 + v15
		v12 = int32(uint32(v11) >> 25)
		m.memory[uint32(t23)] = byte(v12)
		m.memory[uint32(v8+(v15+i32(-8))&v10+i32(8))] = byte(v12)
		t24 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		store32(m.memory[int64(uint32(v0))+8:], uint32(t24-v13&i32(1)))
		v15 = v8 + (i32(0)-v15)*i32(96)
		store32(m.memory[uint32(v15+i32(-96)):], uint32(v1))
		memory_copy(m.memory, uint32(v15+i32(-92)), uint32(v4+i32(4)), uint32(i32(92)))
		t25 := int32(load32(m.memory[int64(uint32(v0))+12:]))
		store32(m.memory[int64(uint32(v0))+12:], uint32(t25+i32(1)))
	}
l1:
	v0 = v15 + i32(-88)
	{
		v1 = v15 + v7
		v8 = v1 + i32(-16)
		t26 := int32(m.memory[uint32(v8)])
		if t26 != i32(1) {
			goto l9
		}
		t27 := int64(load64(m.memory[uint32(v0+v7<<3):]))
		v3 = t27 + i64(1)
		p28 := v3
		if v3 == 0 {
			p28 = i64(-1)
		}
		v3 = p28
	}
l9:
	m.memory[uint32(v8)] = byte(i32(1))
	store64(m.memory[uint32(v0+v7<<3):], uint64(v3))
	if uint32(v2) > uint32(i32(7)) {
		goto l10
	}
	m.memory[uint32(v1+i32(-15))] = byte(i32(0))
	v0 = v7 + i32(-14)
	if v0 == i32(-7) {
		goto l10
	}
	m.memory[uint32(v15+v0)] = byte(i32(0))
	v0 = v7 + i32(-13)
	if v0 == i32(-7) {
		goto l10
	}
	m.memory[uint32(v15+v0)] = byte(i32(0))
	v0 = v7 + i32(-12)
	if v0 == i32(-7) {
		goto l10
	}
	m.memory[uint32(v15+v0)] = byte(i32(0))
	v0 = v7 + i32(-11)
	if v0 == i32(-7) {
		goto l10
	}
	m.memory[uint32(v15+v0)] = byte(i32(0))
	v0 = v7 + i32(-10)
	if v0 == i32(-7) {
		goto l10
	}
	m.memory[uint32(v15+v0)] = byte(i32(0))
	v0 = v2 + i32(-9)
	if v0 == i32(-7) {
		goto l10
	}
	m.memory[uint32(v15+v0)] = byte(i32(0))
	if v2 == i32(1) {
		goto l10
	}
	m.memory[uint32(v1+i32(-8))] = byte(i32(0))
l10:
	m.g0 = v4 + i32(96)
	return v3
}
func (m *Module) fn715(v0, v1, v2 int32) {
	var v3 int64
	var v4, v5 int32
	var v6 int64
	var v7, v8 int32
	var v9, v10 int64
	var v11 int32
	t0 := int64(load64(m.memory[int64(uint32(v1))+16:]))
	t1 := int64(load64(m.memory[int64(uint32(v1))+24:]))
	t2 := m.fn100(t0, t1, v2)
	v3 = t2
	t3 := int32(load32(m.memory[int64(uint32(v1))+4:]))
	v4 = t3
	v5 = v4 & int32(v3)
	v6 = int64(uint64(v3)>>25) & i64(127) * i64(72340172838076673)
	t4 := int32(load32(m.memory[uint32(v1):]))
	v7 = t4
	v8 = i32(0)
l4:
	{
		{
			t5 := int64(load64(m.memory[uint32(v7+v5):]))
			v9 = t5
			v10 = v9 ^ v6
			v10 = (v10 ^ i64(-1)) & (v10 + i64(-72340172838076673)) & i64(-0x7f7f7f7f7f7f7f80)
			if v10 == 0 {
				goto l0
			}
		l2:
			{
				v11 = v7 + (i32(0)-(int32(uint32(int64(bits.TrailingZeros64(uint64(v10))))>>3)+v5)&v4)*i32(96)
				t6 := int32(load32(m.memory[uint32(v11+i32(-96)):]))
				if t6 == v2 {
					store32(m.memory[int64(uint32(v0))+4:], uint32(v1))
					store32(m.memory[uint32(v0):], uint32(v11))
					store32(m.memory[int64(uint32(v0))+12:], uint32(i32(0)))
					return
				}
				v10 = (v10 + i64(-1)) & v10
				if !(v10 == 0) {
					goto l2
				}
			}
		}
	l0:
		if !(v9&(v9<<1)&i64(-0x7f7f7f7f7f7f7f80) == 0) {
			goto l3
		}
		t7 := v5
		v8 = v8 + i32(8)
		v5 = (t7 + v8) & v4
		goto l4
	}
l3:
	{
		t8 := int32(load32(m.memory[int64(uint32(v1))+8:]))
		if t8 != 0 {
			goto l5
		}
		_ = m.fn107(v1, v1+i32(16))
	}
l5:
	store32(m.memory[int64(uint32(v0))+8:], uint32(v2))
	store64(m.memory[uint32(v0):], uint64(v3))
	store32(m.memory[int64(uint32(v0))+12:], uint32(v1))
}
func (m *Module) fn716(v0, v1 int32) {
	var v2, v3, v4 int32
	{
		t0 := int32(load32(m.memory[int64(uint32(v0))+8:]))
		v2 = t0
		t1 := int32(load32(m.memory[uint32(v0):]))
		if v2 != t1 {
			goto l0
		}
		m.fn324(v0)
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
