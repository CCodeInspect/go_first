/*
@author: pauline
@date 2021年09月09日2:11 下午
*/
package main

import "fmt"

func main() {
	var a bool = true
	var b bool = false
	if a && b {
		fmt.Println("a&&b")
	}
	if a || b {
		fmt.Println("a||b")
	}
	fmt.Println("restart")
	a = false
	b = true
	if a && b {
		fmt.Println("a&&b")
	} else {
		fmt.Println("false")
	}
	if !(a && b) {
		fmt.Println("dfd")
	}
}
