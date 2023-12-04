package main

import (
	"fmt"
	"time"
)

// 主goroutine
func main() {
	////用go创建一个形参为空，返回为空的函数
	//go func() {
	//	defer fmt.Println("A.defer")
	//
	//	func() {
	//		defer fmt.Println("B.defer")
	//		fmt.Println("B")
	//	}()
	//
	//	fmt.Println("A")
	//}()

	go func(a int, b int) bool {
		fmt.Println("a:", a, "b:", b)
		return true
	}(10, 20)

	//死循环
	for {
		time.Sleep(1 * time.Second)
	}
}
