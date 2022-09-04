package main

import "fmt"

func main() {
	x := NewMockStruct()
	y := Copy(x).(*MockStruct)
	fmt.Println(x.DeepEqual(y))
	fmt.Println(x)
	fmt.Println(y)
}
