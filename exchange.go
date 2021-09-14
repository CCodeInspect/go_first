/*
@author: pauline
@date 2021年09月13日2:37 下午
*/
package main

import "fmt"

func main() {
	var a int = 10
	var b int = 100
	//fmt.Println(&a, &b)
	swap(&a, &b)
	fmt.Println(&a, &b)
	fmt.Println(a, b)
	swap(&a, &b)
	fmt.Println(a, b)
}

func swap(x *int, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp

}
