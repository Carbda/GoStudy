package main

import "fmt"

func main3() {
	//定义channel
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		//close可以关闭一个channel
		close(c)
	}()

	for {
		//ok如果为true表示channel没有关闭，为false表示channel已经关闭
		if data, ok := <-c; ok {
			fmt.Println(data)
		} else {
			break
		}
	}

	fmt.Println("main finish.")
}
