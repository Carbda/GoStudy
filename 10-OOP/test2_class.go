package main

import "fmt"

//类或方法首字母大写，表示其他包也可以访问
type Hero struct {
	//如果类的属性首字母大写，表示属性是可以对外访问的，否则只能类内部访问
	Name  string
	Ad    int
	Level int
}

func (hero *Hero) Show() {
	fmt.Println("Name = ", hero.Name)
	fmt.Println("Ad = ", hero.Ad)
	fmt.Println("Level = ", hero.Level)
}

func (hero *Hero) GetName() {
	fmt.Println("name = ", hero.Name)
}

func (hero *Hero) SetName(newName string) {
	//hero是当前对象的一个拷贝
	hero.Name = newName
}

func main2() {
	//创建一个对象
	hero := Hero{Name: "zhangsan", Ad: 100, Level: 1}
	hero.Show()
	hero.SetName("lisi")
	hero.Show()
}
