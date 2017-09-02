package interpreter

import (
	"strings"
)

type expression interface {
	Interpret(*context) string
}

type words struct {
	val string
}

func NewWords(val string) *words {
	words := &words{val}
	return words
}

func (words words) Interpret(context *context) string {
	return words.val
}

type variable struct {
	val string
}

func NewVariable(val string) *variable {
	variable := &variable{val}
	return variable
}

func (variable variable) Interpret(context *context) string {
	//index := strings.Index(variable.val, "$")

	var index int
	for {
		index = strings.Index(variable.val, "$")
		if index < 0 {
			break
		}

		variable.val = string([]rune(variable.val)[1:])
	}

	return context.args[variable.val].Interpret(context)
}
