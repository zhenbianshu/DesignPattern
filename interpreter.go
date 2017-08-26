package main

import (
	"fmt"
	"interpreter"
)

func main() {
	hello := interpreter.NewWords("hello")
	greeting := interpreter.NewVariable("greeting")
	context := interpreter.NewContext()
	context.AddItem("$greeting", hello)
	context.AddItem("$test", greeting)

	var_test := interpreter.NewVariable("test")
	fmt.Println(var_test.Interpret(context)) // $greeting = "hello"; $test = "greeting"; $$test = "hello"
}
