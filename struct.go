/*
@author: pauline
@date 2021年09月13日3:12 下午
*/
package main

import "fmt"

type book struct {
	year   string
	name   string
	author string
}

func main() {
	//fmt.Println(book{author: "aauthor", year: "ttitle", name: "nname"})
	var bk_1 book
	var bk_2 book
	var a *book
	var b *book
	bk_1.name = "《格列佛游记》"
	bk_1.author = "乔纳森·斯威夫特\n"
	bk_1.year = "19990101"
	bk_2.name = "《格列佛游记》"
	bk_2.author = "乔纳森·斯威夫特\n"
	bk_2.year = "19990101"

	fmt.Println(bk_1, bk_2)
	//print_b(bk_1)
	//print_b(bk_2)
	a = &bk_1
	b = &bk_2
	fmt.Println(&a, &b)
}
//func print_b(boo book) {
//	fmt.Println(boo)
//
//}
