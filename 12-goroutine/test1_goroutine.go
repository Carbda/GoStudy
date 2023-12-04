package main

import (
	"fmt"
	"time"
)

// 子goroutine
func newTask() {
	i := 0
	for {
		i++
		fmt.Printf("newTask: i = %d\n", i)
		time.Sleep(1 * time.Second)
	}
}

// 主goroutine
func main1() {
	//创建go程，去执行newTask()流程
	go newTask()

	fmt.Println("main: exit")
	//i := 0
	//
	//for {
	//	i++
	//	fmt.Printf("main: i = %d\n", i)
	//	time.Sleep(1 * time.Second)
	//}
}
