package main

import
	"fmt"


func main() {
	var stock_code = "123"
	var qq = "2223"
	fmt.Printf("第二个函数我觉得能编译也能跑通" + "22")
	var url = fmt.Sprintf(stock_code, qq)
	fmt.Println(url)
}
