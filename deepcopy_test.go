package main

import (
	"encoding/json"
	"testing"
)

//go test -bench="." -benchmem

func BenchmarkCopy(b *testing.B) {
	x := NewMockStruct()
	for i := 0; i < b.N; i++ {
		Copy(x)
	}
}

func BenchmarkMarshal(b *testing.B) {
	x := NewMockStruct()
	for i := 0; i < b.N; i++ {
		CopyMockStruct(x)
	}
}

func CopyMockStruct(v *MockStruct) {
	y := &MockStruct{}
	bytes, _ := json.Marshal(v)
	json.Unmarshal(bytes, y)
}
