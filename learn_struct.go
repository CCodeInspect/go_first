/*
@author: pauline
@date 2021年09月09日3:20 下午
*/
package main

import "fmt"

func main() {
	var a int
	a = get_bigger(1, 2)
	fmt.Println(a)
}
func get_bigger(n_0, n_1 int) int {
	var result int
	if n_0 > n_1 {
		fmt.Println("nb >")
	} else {
		result = n_1
	}
	return result
}
