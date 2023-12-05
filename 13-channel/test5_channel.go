package main

import "fmt"

func fibonacci(c chan int, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x: //若c可写，则会进这个case
			x = y
			y = x + y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

/*
select：如果有多个 case 都可以运行，select 会随机公平地选出一个执行，其他不会执行
否则：	如果有default，则执行
		如果没有default，select会阻塞，直到某个通道可以运行
*/

func main() {
	//定义channel
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}

		quit <- 0
	}()

	fibonacci(c, quit)

	str := "abc"
	fmt.Printf("type: %T\n", str[1])
	fmt.Println(str[1])
}
