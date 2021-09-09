/*
@author: pauline
@date 2021年09月09日11:22 上午
*/
package main

import "fmt"

func main() {
	var a int = 1
	var b int = 2

	if a == b {
		fmt.Println("a==b")
	} else {
		fmt.Println("a!=b")
	}
	if a < b {
		fmt.Println("a<b")
	} else {
		fmt.Println("a不小于b")
	}
	if a > b {
		fmt.Println("a>b")
	} else {
		fmt.Println("a 不大于b")
	}

	a = 5
	b = 6

	if a <= b {
		fmt.Println("a小于等于b", a)
	}
	if b >= a {
		fmt.Println("b小于等于a", b)
	}

	/*var c int
	var d int
	var e int
	var f int
	var g int*/
	/*加减乘除取余自增自减*/
	/*c = a + b
	d = a - b
	e = a * b
	f = a / b
	g = a % b

	fmt.Println(c, "加")
	fmt.Println(d, "减")
	fmt.Println(e, "乘")
	fmt.Println(f, "除")
	fmt.Println(g, "取余")
	a = 1
	a++
	fmt.Println(a, "a++")
	a = 21
	a--
	fmt.Println(a, "a--")*/

}
