package main

import "fmt"

//继承

type Human struct {
	name string
	sex  string
}

func (man *Human) Eat() {
	fmt.Println("Human.Eat()...")
}

func (man *Human) Walk() {
	fmt.Println("Human.Walk()...")
}

//SuperMan继承Human
type SuperMan struct {
	Human //表示该类继承了Human类

	level int
}

//重定义父类方法
func (man *SuperMan) Eat() {
	fmt.Println("SuperMan.Eat()...")
}

func (man *SuperMan) Fly() {
	fmt.Println("SuperMan.Fly()...")
}

func (man *SuperMan) print() {
	fmt.Println("name = ", man.name)
	fmt.Println("sex = ", man.sex)
	fmt.Println("level = ", man.level)

}

func main3() {
	h := Human{"zhangsan", "male"}
	h.Eat()
	h.Walk()

	//定义一个子类对象
	//s := SuperMan{Human{"lisi", "female"}, 88}
	var s SuperMan
	s.name = "lisi"
	s.sex = "female"
	s.level = 88

	s.Eat()
	s.Walk()
	s.Fly()

	s.print()
}
