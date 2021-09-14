/*
@author: pauline
@date 2021年09月09日5:07 下午
*/
package main

import (
	"fmt"
	"unsafe"
)

//func main() {
/*a := [4]int{1, 2, 3, 4}
var i int
for i = 0; i < 3; i++ {
	//fmt.Println(i, a[i])
	fmt.Println(&i, &a) //数组怎么取地址
}*/
/*	a := 2
	p := &a
	fmt.Println(p)
	*p = 3
	fmt.Println(a)
*/
//声明一个int值的指针变量
/*var ptr *int
var str_0 *string //声明一个string值的指针变量
var str_1 int
str_2 := &str_1
fmt.Println(&ptr)
fmt.Println(&str_0)
fmt.Println(str_2)*/
//var ip **int
//var t *time.Time
//if ip == nil {
//	fmt.Println(ip, "error")
//}
//if t == nil {
//	fmt.Println(t, "error")
//}
//if ip == t {
//	fmt.Println("error")
//} // ./main.go:11:8: invalid operation: ip == sp (mismatched types *int and *string)
//fmt.Printf("ip:%#v\n\n", ip)
//fmt.Printf("t:%#v\n\n", t)
//}

func main() {
	var i uint = 10
	var p2 *int

	p1 := &i //10
	p2 = (*int)(unsafe.Pointer(p1))
	fmt.Println(*p2)
}
