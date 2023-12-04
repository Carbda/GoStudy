package main

import (
	"fmt"
	"time"
)

func main2() {
	//定义带缓冲的channel
	c := make(chan int, 3)

	fmt.Println("len(c) =", len(c), "cap =", cap(c))

	go func() {
		defer fmt.Println("goroutine finish")

		for i := 0; i < 4; i++ {
			c <- i
			fmt.Println("goroutine running: element =", i, " len(c) =", len(c), "cap =", cap(c))
		}
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 4; i++ {
		num := <-c //从c接收数据
		fmt.Println("num =", num)
	}

	time.Sleep(1 * time.Second)

	fmt.Println("main finish")
}
