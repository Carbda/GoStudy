package main

import (
	"fmt"
)

func main1() {
	//定义channel
	c := make(chan int)

	go func() {
		defer fmt.Println("goroutine finish")
		fmt.Println("goroutine go...")
		c <- 666 //将666发送给c
	}()

	num := <-c //从c中接收数据，并赋值给num，这时候是阻塞等待直到接收

	fmt.Println("num =", num)
	fmt.Println("main goroutine finish")
}
