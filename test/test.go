package test

import "fmt"

var TestVar1 = 1
var TestVar2 = 1 + TestVar1

func Print() {
	fmt.Println(TestVar1)
	fmt.Println(TestVar2)
}
