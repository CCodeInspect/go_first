/*
@author: pauline
@date 2021年09月09日3:51 下午
*/
package main

import "fmt"

func main() {
	var a int
	a = add(1, 2)
	fmt.Println(a)

}
func add(x, y int) int {
	return x + y
}
