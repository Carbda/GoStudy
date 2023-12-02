package main

import "fmt"

// 切片是引用传递，而数组是值传递
func printSlice(array []int) {
	for _, value := range array {
		println("value = ", value)
	}
}

func main2() {
	myArray := []int{1, 2, 3, 4}

	fmt.Printf("type of myArray : %T\n", myArray)

	printSlice(myArray)
}
