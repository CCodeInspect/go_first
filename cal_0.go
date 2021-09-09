/*
@author: pauline
@date 2021年09月09日2:27 下午
*/
package main

import "fmt"

func main() {
	/*var a int = 4
	var b int32
	var c float32
	var ptr *int
	fmt.Println(a, b, c, ptr)
	ptr = &a
	fmt.Println(ptr, "a的地址")*/
	var a int = 20
	var b int = 10
	var c int = 15
	var d int = 5
	var e int
	e = (a + b) * c / d //30*15/5 =90
	fmt.Println(e)
	e = ((a + b) * c) / d //30*15/5=90
	fmt.Println(e)
	e = (a + b) * (c / d) //30 *3=90
	fmt.Println(e)
	e = a + (b*c)/d //20+150/5 -> 20+30 =50
	fmt.Println(e)
}
