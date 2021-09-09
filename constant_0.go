package main

import (
	"fmt"
	"unsafe"
)

func main() {
	const leng int = 2
	const wid int = 9
	var area int
	//const a, b, c = 1, false, "str"
	const (
		a = "abc"
		b = len(a)
		c = unsafe.Sizeof(a)
	)
	area = leng * wid
	fmt.Println("面积为", area)
	fmt.Println(a, b, c)

}
