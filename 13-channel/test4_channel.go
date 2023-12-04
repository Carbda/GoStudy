package main

import "fmt"

func main4() {
	//定义channel
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		//close可以关闭一个channel
		close(c)
	}()

	//for {
	//	//ok如果为true表示channel没有关闭，为false表示channel已经关闭
	//	if data, ok := <-c; ok {
	//		fmt.Println(data)
	//	} else {
	//		break
	//	}
	//}

	//可以用range来迭代不断操作channel，若c未关闭但没数据，会阻塞，若c已经关闭，则跳出循环
	for data := range c {
		fmt.Println(data)
	}

	fmt.Println("main finish.")
}
