package lib1

import "fmt"

// 当前lib1包的API
func Lib1Test() {
	fmt.Println("Lib1Test() ...")
}

func init() {
	fmt.Println("lib1.init() ...")
}
