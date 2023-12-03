package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) Call() {
	fmt.Println("user is called...")
	fmt.Println("%v\n", u)
}

func (u *User) SetName(str string) {
	u.Name = str
}

func main5() {
	user := User{1, "abc", 22}
	user.SetName("cba")
	DoFiledAndMethod(user)
}

func DoFiledAndMethod(input interface{}) {
	//获取input的type
	inputType := reflect.TypeOf(input)
	fmt.Println("type is:", inputType.Name())
	//获取input的value
	inputValue := reflect.ValueOf(input)
	fmt.Println("value is:", inputValue)

	//通过type获取字段
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()
		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}
	//通过type获取方法
	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}
}
