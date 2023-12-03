package main

import (
	"fmt"
	"io"
	"os"
)

func main2() {
	tty, err := os.OpenFile("tty", os.O_RDWR, 0)

	if err != nil {
		fmt.Println("open file error", err)
		return
	}

	var r io.Reader
	r = tty

	var w io.Writer
	w = r.(io.Writer)

	w.Write([]byte("hello\n"))
}
