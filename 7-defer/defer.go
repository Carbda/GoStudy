package main

import "fmt"

func main() {
	defer fmt.Println("main end1") // 最后执行，有点像java的finally，defer会在return之后执行
	defer fmt.Println("main end2") // end2先执行，因为是压栈的形式，先入栈的后出栈

	fmt.Println("main::hello go 1")
	fmt.Println("main::hello go 2")
}
