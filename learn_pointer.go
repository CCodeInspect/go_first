/*
@author: pauline
@date 2021年09月09日5:07 下午
*/
package main

import (
	"fmt"
)

func main() {
	var a = 1
	var b = a
	fmt.Println("地址为", &a)
	fmt.Println("第二次地址变成这样了：", &a, &b)
	//
	var ip *int //声明指针变量
	ip = &a     //取到的地址存到指针变量中
	fmt.Println("指针变量：", ip)
	//
	var nil_pointer *float32
	fmt.Println("声明一个空指针，他存到nil_pointer中了,不信你打印看看", nil_pointer)
	//
	arr := []int{1, 2, 3, 4}
	const (
		max = 3
	)
	var ptr [max]*int
	var i int
	for i = 0; i < max; i++ {
		fmt.Println(i, arr[i], "laallalal")
	}
	for i = 0; i < max; i++ {
		ptr[i] = &a[i]
		fmt.Println(i, ptr[i])
	}
}
