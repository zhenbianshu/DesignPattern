package main

import (
	"fmt"
	"singleton"
)

func main() {
	object1 := singleton.GetInstance()
	fmt.Println(object1.GetName()) // test1

	object2 := singleton.GetInstance()
	object2.SetName("test2")
	fmt.Println(object2.GetName()) // test2
}
