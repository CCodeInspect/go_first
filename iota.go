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
	const (
		a = iota      //0
		b = iota      //1
		c = iota      //2
		d = "hh"      //hh
		e             //hh
		f = iota      //5
		g = iota << 6 //6*2^6 =128
		h = iota >> 2 //7*2^2 =28
	)

	fmt.Println(a, b, c, d, e, f, g, h)
}
