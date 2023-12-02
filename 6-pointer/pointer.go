package main

import "fmt"

func main() {
	a := 10
	b := 20
	//swap
	swap(&a, &b)
	fmt.Println("a = ", a, " b = ", b)

	var p *int
	p = &a
	fmt.Println(p, &a)

	var pp **int // 二级指针
	pp = &p
	fmt.Println(pp, &p)
}

func swap(a *int, b *int) {
	var tmp int
	tmp = *a
	*a = *b
	*b = tmp
}
