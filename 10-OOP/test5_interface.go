package main

import "fmt"

//interface{}是万能数据类型，我感觉相当于Java的Object类型
func myFunc(arg interface{}) {
	fmt.Println("myFunc is called..")
	fmt.Println(arg)

	//interface{}如何区分，此时底层的数据类型是什么？

	//给 interface{} 提供“类型断言”机制
	value, ok := arg.(string)
	if !ok {
		fmt.Println("arg is not string")
	} else {
		fmt.Println("arg is string, value = ", value)
	}
	fmt.Println("--------------------")
}

type Book2 struct {
	auth string
}

func main() {
	book := Book2{"Golang"}

	myFunc(book)
	myFunc(100)
	myFunc("abc")
	myFunc(3.14)
}
