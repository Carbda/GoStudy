package main

import "fmt"

//声明一种数据类型，是int的别名
type myint int

type Book struct {
	title string
	auth  string
}

//只传了book的副本
func changeBook1(book Book) {
	book.auth = "lisi"
}

//传指针
func changeBook2(book *Book) {
	book.auth = "lisi"
}

func main1() {
	//var a myint = 10
	//fmt.Println(a)
	//fmt.Printf("type of a : %T\n", a)

	var book1 Book
	book1.title = "Golang"
	book1.auth = "zhangsan"

	fmt.Println(book1)

	changeBook1(book1)
	fmt.Println(book1)

	changeBook2(&book1)
	fmt.Println(book1)

}
