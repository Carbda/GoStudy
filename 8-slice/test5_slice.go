package main

import "fmt"

func main() {
	s := []int{1, 2, 3}

	s1 := s[0:2]
	fmt.Println(s1)
	fmt.Printf("type of s1 : %T\n", s1)

	s[0] = 100
	fmt.Println(s[0], s1[0]) // s和s1指向的地址是一样的

	// copy可以深拷贝
	s2 := make([]int, 3)
	// 将s中的值拷贝到s2
	copy(s2, s)
	fmt.Println(s2) // s和s1指向的地址是一样的
}
