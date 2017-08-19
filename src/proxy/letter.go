package proxy

import "fmt"

type letterContent interface {
	get()
}

type realContent struct {
}

func (realContent) get() {
	fmt.Println("The content is:I love you, too.")
}

type tempContent struct {
}

func (tempContent) get() {
	fmt.Println("I am try my best to get the reply letter.")
}
