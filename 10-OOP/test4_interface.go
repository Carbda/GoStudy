package main

import "fmt"

//interface本质是指针
type AnimalIF interface {
	Sleep()
	GetColor() string
	GetType() string
}

type Cat struct {
	color string
}

func (cat *Cat) Sleep() {
	fmt.Println("Cat.Sleep()...")
}

func (cat *Cat) GetColor() string {
	return cat.color
}

func (cat *Cat) GetType() string {
	return "Cat"
}

type Dog struct {
	color string
}

func (dog *Dog) Sleep() {
	fmt.Println("Dog.Sleep()...")
}

func (dog *Dog) GetColor() string {
	return dog.color
}

func (dog *Dog) GetType() string {
	return "Dog"
}

func showAnimal(animal AnimalIF) {
	animal.Sleep()
	fmt.Println("color = ", animal.GetColor())
	fmt.Println("type = ", animal.GetType())
}

func main4() {
	var animal AnimalIF //接口的数据类型，父类指针

	animal = &Cat{"Green"}
	animal.Sleep() //调用Cat的Sleep()

	animal = &Dog{"Yellow"}
	animal.Sleep()

	fmt.Println("----------------------")

	cat := Cat{"Yellow"}
	dog := Dog{"Green"}

	showAnimal(&cat)
	showAnimal(&dog)
}
