package lib2

import "fmt"

// 当前lib2包的API
func Lib2Test() {
	fmt.Println("Lib2Test() ...")
}

func init() {
	fmt.Println("lib2.init() ...")
}
