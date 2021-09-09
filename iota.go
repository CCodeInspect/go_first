package main

import "fmt"

func main() {
	//const (
	//	a = iota //0
	//	b
	//	c
	//	d = "hah"
	//	e
	//	f = iota
	//)
	//const (
	//	a = iota      //0
	//	b = iota      //1
	//	c = iota      //2
	//	d = "hh"      //hh
	//	e             //hh
	//	f = iota      //5
	//	g = iota << 6 //6*2^6 =384
	//	h = iota >> 2 //7/2^2 =7/4
	//)
	const (
		i = 1 << iota // 1*0^2=1
		j = 3 << iota //3*1^2=6
		k = 3 << iota //n*2^2=12
		l = 3 << iota //n*2^3=24
	)
	fmt.Println(i, j, k, l)

	//fmt.Println(a, b, c, d, e, f, g, h)
}
