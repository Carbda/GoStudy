package main

import (
	"fmt"
	"strconv"
)

func float2int() {
	var sum int = 17
	var count int = 5
	var mean float32

	mean = float32(sum) / float32(count)
	fmt.Printf("mean 的值为: %f\n", mean)
}

func str2float() {
	str := "3.14"
	num, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Println("转换错误:", err)
	} else {
		fmt.Printf("字符串 '%s' 转为浮点型为：%f\n", str, num)
	}
}

func interface2string() {
	var i interface{} = "Hello, World"
	str, ok := i.(string)
	if ok {
		fmt.Printf("'%s' is a string\n", str)
	} else {
		fmt.Println("conversion failed")
	}
}

func main2() {
	float2int()
	fmt.Println("---------")
	str2float()
	fmt.Println("---------")
	interface2string()
	fmt.Println("---------")
}
