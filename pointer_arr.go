/*
@author: pauline
@date 2021年09月13日10:44 上午
*/
package main

import "fmt"

//func main() {
//	const a = 4
//	var i int
//	arr := [...]int{1, 2, 3, 4, 5, 6} //定义一个arr的数组
//	for i = 0; i < a; i++ {
//		fmt.Println(i, arr[i])
//	}
//
//}

/*保存一个数组到一个指针数组中*/
//func main() {
//	a := [...]int{1, 2, 3, 4}
//	//var po []*int
//	//var ps []*string
//	const b = 4
//	var ptr [b]*int //声明了 整型 + 指针 数组
//	var i int
//	for i = 0; i < b; i++ {
//		ptr[i] = &a[i] //整型数组 赋给 指针数组
//		fmt.Println(ptr, ptr[i])
//	}
//	//fmt.Println(ptr, ptr[i])
//}
//
func main() {
	var abc int
	var ptr *int
	var ll_p **int
	var qqq ***int
	abc = 1234
	ptr = &abc
	ll_p = &ptr
	qqq = &ll_p
	fmt.Println(ptr, ll_p, abc)
	fmt.Println(*ptr, **ll_p)
	fmt.Println(***qqq)

}
