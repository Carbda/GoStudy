package main

import (
	"fmt"
	"reflect"
)

type Resume struct {
	Name string `info:"name" doc:"我的名字"`
	Sex  string `info:"Sex"`
}

func findTag(arg interface{}) {
	t := reflect.TypeOf(arg).Elem()

	for i := 0; i < t.NumField(); i++ {
		tagInfo := t.Field(i).Tag.Get("info")
		tagDoc := t.Field(i).Tag.Get("doc")
		fmt.Println("info :", tagInfo, "doc :", tagDoc)
	}
}

func main6() {
	var re Resume

	findTag(&re)
}
