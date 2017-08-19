package state

import "fmt"

type Stage interface {
	water()
	harvest()
}

type seedling struct {
}

func (seedling) water() {
	fmt.Println("The plant is growing!")
}

func (seedling) harvest() {
	fmt.Println("The plant is just planted！")
}

type bloom struct {
}

func (bloom) water() {
	fmt.Println("The plant is ripping!")
}

func (bloom) harvest() {
	fmt.Println("The plant is blooming, don't do that！")
}

type maturity struct {
}

func (maturity) water() {
	fmt.Println("The plant don't need water!")
}

func (maturity) harvest() {
	fmt.Println("You got lots of fruits！")
}
