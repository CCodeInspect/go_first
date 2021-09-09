/*
@author: pauline
@date 2021年09月09日2:46 下午
*/
package main

import "fmt"

func main() {
	/*	sum := 0
		for i := 0; i <= 10; i++ {
			sum += 1
		}
		fmt.Println(sum)*/
	strings := []string{"str", "goo"}
	for i, s := range strings {
		fmt.Println(i, s)
	}
	numbers := [6]int{1, 2, 3, 4, 5, 6}
	for i, x := range numbers {
		fmt.Println(i, x)
	}
	floats := [4]float64{1.2, 1.1, 1.222}
	for i, x := range floats {
		fmt.Println(i, x)
	}
	nums := [5]bool{true, false}
	for i, x := range nums {
		fmt.Println(i, x)
	}
}
