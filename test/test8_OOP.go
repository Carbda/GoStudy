package main

import (
	"fmt"
	"sync"
)

type OopHandler interface {
	Run()
}

type Oop struct {
	name string
}

var oopOnce sync.Once
var oopIns *Oop

func GetOopIns() *Oop {
	oopOnce.Do(func() {
		oopIns = &Oop{
			name: "OOP",
		}
	})

	return oopIns
}

func (p *Oop) Run() {
	fmt.Printf("%s.Run start", p.name)
}
