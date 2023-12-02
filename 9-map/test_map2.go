package main

import (
	"fmt"
)

func printMap(cityMap map[string]string) {
	//cityMap是引用传递
	for key, value := range cityMap {
		fmt.Printf("key = %s, value = %s\n", key, value)
	}
}

//map是引用传递
func changeMap(cityMap map[string]string) {
	cityMap["England"] = "London"
}

func main() {
	cityMap := make(map[string]string)
	//添加
	cityMap["China"] = "Beijing"
	cityMap["Japan"] = "Tokyo"
	cityMap["USA"] = "NewYork"
	//遍历
	printMap(cityMap)
	fmt.Println("-----------------")
	//删除
	delete(cityMap, "China")
	//修改
	cityMap["USA"] = "DC"
	changeMap(cityMap)
	//遍历
	printMap(cityMap)
}
