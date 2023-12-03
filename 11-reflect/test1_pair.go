package main

import "fmt"

func main1() {
	var a string

	//pair<statictype:string,value:"aceld">
	a = "aceld"

	//pair<type: value: >
	var allType interface{}
	allType = a

	str, _ := allType.(string)

	fmt.Println(str)
}
