package main

import "fmt"

func main1() {
	// 固定长度的数组
	var myArray1 [10]int

	myArray2 := [10]int{1, 2, 3, 4}

	for i := 0; i < len(myArray1); i++ {
		fmt.Print(myArray1[i], " ")
	}
	fmt.Print("\n")
	for i := range myArray2 {
		fmt.Print(myArray2[i], " ")
	}
}
