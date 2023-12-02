package main

import "fmt"

func main1() {
	//第一种声明
	var myMap1 map[string]string
	if myMap1 == nil {
		fmt.Println("map为空")
	}
	//使用map前，要先给map分配数据空间
	myMap1 = make(map[string]string, 10)

	myMap1["one"] = "1"
	myMap1["two"] = "2"
	myMap1["three"] = "3"

	fmt.Println(myMap1)

	//第二种声明
	myMap2 := make(map[int]string)
	myMap2[1] = "one"
	myMap2[2] = "two"
	myMap2[3] = "three"
	fmt.Println(myMap2)

	//第三种声明
	myMap3 := map[string]string{
		"1": "one",
		"2": "two",
	}
	fmt.Println(myMap3)
}
