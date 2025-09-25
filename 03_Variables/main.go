package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var age string = "salam salam salam salam"
	fmt.Println(age)

	t, t1 := test(2, 4)
	fmt.Println(t)
	fmt.Println(t1)

	// count := 5
	fmt.Println(unsafe.Sizeof(age))
}

func test(a int, b int) (int, int) {
	return a + b, 2
}
