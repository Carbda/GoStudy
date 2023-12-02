package main

import "fmt"

func main3() {
	//slice1 := []int{1, 2, 3} // 声明slice1是一个切片，并初始化

	//var slice1 []int        // 声明slice1是一个切片，还没分配空间
	//slice1 = make([]int, 3) // 开辟3个空间
	//slice1[0] = 100

	//var slice1 []int = make([]int, 3) // 声明slice1是一个切片，同时分配空间

	slice1 := make([]int, 3)

	fmt.Printf("len = %d slice = %v\n", len(slice1), slice1)

	if slice1 == nil {
		fmt.Println("空切片")
	} else {
		fmt.Println("不是空切片")
	}
}
