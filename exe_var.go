package main

import "fmt"

func main() {
	_, numb, _ := numbers()
	fmt.Println(numb)

}
func numbers() (int, int, string) {
	a, b, c := 1, 2, "str"
	return a, b, c

}

//func main() {
//	/*var b bool = true
//	var a int = 1
//	var c float64 = 2.33
//	var str_0 string = "str_0"
//	var blank_int int
//	var true_bool bool = true
//	var blank_bool bool
//	var blank_str string
//	fmt.Println(str_0)
//	fmt.Println(a, b, c)
//	fmt.Println(blank_int, "blank int")
//	fmt.Println(true_bool, "true bool")
//	fmt.Println(blank_bool, "blank bool = false ")
//	fmt.Println(blank_str, "''")*/
//
//	/*<nil> [] map[] <nil> <nil> <nil>?*/
//	//var a *int
//	//var b []int            /*这个啥意思？*/
//	//var c map[string]int   /*这个啥意思？*/
//	//var d chan int         /*这个啥意思？*/
//	//var e func(string) int /*这个啥意思？*/
//	//var f error            /*这是个接口？？*/
//	//fmt.Println(a, b, c, d, e, f)
//	//var ffloatt float64
//	//ffloatt = 2223.3
//	//ffloattp := 1.200000123010230102301023012030123010230102301023012030
//	//ffloatt := 23.4
//	//str_1 := "123"
//
//	//name0, name1, name2 := 1, "qq", 3.2
//	//fmt.Println(name0, name1, name2)
//	//fmt.Println(ffloatt, str_1)
//
//}
//func main() {
//	var x, y int
//	var (
//		a int
//		b int
//	)
//	name := 1
//	sex := "female"
//	bo_0 := false
//	fmt.Println(x, y, a, b, name, sex, bo_0)
//	//undefin_0 := 11
//	a_0, a_1, a_2 := 1, 2, "str_1"
//	a_0 = a_1
//	fmt.Println(a_0, a_1, a_2)
//
//}
