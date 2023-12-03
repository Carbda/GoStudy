package main

import "fmt"

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

type Book struct {
}

func (book *Book) ReadBook() {
	fmt.Println("Read a book")
}

func (book *Book) WriteBook() {
	fmt.Println("Write a book")
}

func main3() {
	//b: pair<type:Book value:Book{}地址>
	b := &Book{}

	//r: pair<type: value:>
	var r Reader
	//r: pair<type:Book value:Book{}地址>
	r = b

	r.ReadBook()

	var w Writer
	w = r.(Writer)
	w.WriteBook()

	fmt.Println(&b, &r, &w)
}
