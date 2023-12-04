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

}
