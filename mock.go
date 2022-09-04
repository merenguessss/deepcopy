package main

import (
	"fmt"
	"reflect"
)

const LEN = 10

type MockStruct struct {
	x       int
	Integer int
	String  string
	Double  float64

	P          *int
	IntegerPtr *int
	StringPtr  *string
	DoublePtr  *float64

	//array
	Array    [LEN]int
	PtrArray [LEN]*InnerStruct

	Map         map[int]int
	InnerMap    map[int]InnerStruct
	InnerPtrMap map[int]*InnerStruct

	Slice    []int
	PtrSlice []*InnerStruct

	InnerStruct    InnerStruct
	InnerStructPtr *InnerStruct
}

func (m *MockStruct) DeepEqual(ms *MockStruct) bool {
	if !reflect.DeepEqual(ms, m) {
		return false
	}

	// compare basic type
	if &m.x == &ms.x || &m.Integer == &ms.Integer || &m.String == &ms.String || &m.Double == &ms.Double {
		return false
	}

	// compare pointer
	if m.P == ms.P || m.IntegerPtr == ms.IntegerPtr ||
		m.StringPtr == ms.StringPtr || m.DoublePtr == ms.DoublePtr {
		return false
	}

	// compare map
	if ptrEqual(m.Map, ms.Map) || ptrEqual(m.InnerMap, ms.InnerMap) ||
		ptrEqual(m.InnerPtrMap, ms.InnerPtrMap) {
		return false
	}

	for k := range m.InnerMap {
		if m.InnerMap[k] == ms.InnerMap[k] {
			return false
		}
	}

	for k := range m.InnerPtrMap {
		if m.InnerPtrMap[k] == ms.InnerPtrMap[k] {
			return false
		}
	}

	// compare slice
	if ptrEqual(m.Slice, ms.Slice) || ptrEqual(m.PtrSlice, ms.PtrSlice) {
		return false
	}

	for i := range m.Slice {
		if &m.Slice[i] == &ms.Slice[i] {
			return false
		}
	}

	for i := range m.PtrSlice {
		if m.PtrSlice[i] == ms.PtrSlice[i] {
			return false
		}
	}

	// compare struct
	if m.InnerStruct == ms.InnerStruct || m.InnerStructPtr == ms.InnerStructPtr {
		return false
	}
	return true
}

func ptrEqual(a, b interface{}) bool {
	p1 := fmt.Sprintf("%p", a)
	p2 := fmt.Sprintf("%p", b)
	return p1 == p2
}

func NewMockStruct() *MockStruct {
	x := new(int)
	*x = 1
	i := new(int)
	*i = 1
	s := new(string)
	*s = "b"
	d := new(float64)
	*d = 1.23

	m := make(map[int]int, LEN)
	for i := 0; i < LEN; i++ {
		m[i] = i
	}

	var array [LEN]int
	for i := 0; i < LEN; i++ {
		array[i] = i
	}

	var ptrArray [LEN]*InnerStruct
	for i := 0; i < LEN; i++ {
		ptrArray[i] = NewInnerStructPtr()
	}

	innerMap := make(map[int]InnerStruct, LEN)
	for i := 0; i < LEN; i++ {
		innerMap[i] = NewInnerStruct()
	}

	innerPtrMap := make(map[int]*InnerStruct, LEN)
	for i := 0; i < LEN; i++ {
		innerPtrMap[i] = NewInnerStructPtr()
	}

	slice := make([]int, LEN)
	for i := 0; i < LEN; i++ {
		slice[i] = i
	}

	ptrSlice := make([]*InnerStruct, LEN)
	for i := 0; i < LEN; i++ {
		ptrSlice[i] = NewInnerStructPtr()
	}

	innerStruct := NewInnerStruct()
	innerStructPtr := NewInnerStructPtr()

	return &MockStruct{
		x:              1,
		Integer:        2,
		String:         "a",
		Double:         1.5,
		P:              x,
		IntegerPtr:     i,
		StringPtr:      s,
		DoublePtr:      d,
		Array:          array,
		PtrArray:       ptrArray,
		Map:            m,
		InnerMap:       innerMap,
		InnerPtrMap:    innerPtrMap,
		Slice:          slice,
		PtrSlice:       ptrSlice,
		InnerStruct:    innerStruct,
		InnerStructPtr: innerStructPtr,
	}
}

type InnerStruct struct {
	String    string
	StringPtr *string
}

func NewInnerStruct() InnerStruct {
	str := "b"
	return InnerStruct{
		String:    "a",
		StringPtr: &str,
	}
}

func NewInnerStructPtr() *InnerStruct {
	s := NewInnerStruct()
	return &s
}
