/*
@author: pauline
@date 2021年09月09日4:02 下午
*/
package main

import "fmt"

func main() {
	var arr_0 = [5]int{1, 2, 3, 4, 5}
	var arr_1 = [10]string{"s1", "s2", "s3", "s4"}
	var arr_2 = [...]float64{1.1, 2.4, 3.3}
	fmt.Println(arr_0, arr_1, arr_2)
	var define_arr = []int{0: 1, 2: 0}
	fmt.Println(define_arr)
	fmt.Println(define_arr[0])
	//
	var i, j, k int
	arr := [5]float32{1.1, 2.1, 3.1, 4.1, 5.1}
	/*输出数组元素*/
	for i, x := range arr {
		fmt.Println(i, x)
	}
	for i = 0; i < 5; i++ {
		fmt.Println(i, arr[i])
	}
	arr_p := [...]int{1, 2, 3, 4, 5}
	for j = 0; j < 5; j++ {
		fmt.Println(j, arr_p[j])
	}
	arr_q := [3]int{0: 1, 1: 1, 2: 2}
	for k = 0; k < 3; k++ {
		fmt.Println(k, arr_q[k])
	}

}
