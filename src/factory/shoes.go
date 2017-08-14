package factory

import "fmt"

type Shoe interface {
	Wear() // todo tooooo many implements!!! How to Optimize???
}

type Function struct {
	name  string
	price int
}

type liNingPlasticShoe struct {
	name  string
	price int
}

func (shoe liNingPlasticShoe) Wear() {
	fmt.Printf("a %s, it's price is %d元;\n", shoe.name, shoe.price)
}

type liNingCottonShoe struct {
	name  string
	price int
}

func (shoe liNingCottonShoe) Wear() {
	fmt.Printf("a %s, it's price is %d元;\n", shoe.name, shoe.price)
}

type liNingSportsShoe struct {
	name  string
	price int
}

func (shoe liNingSportsShoe) Wear() {
	fmt.Printf("a %s, it's price is %d元;\n", shoe.name, shoe.price)
}

type adidasPlasticShoe struct {
	name  string
	price int
}

func (shoe adidasPlasticShoe) Wear() {
	fmt.Printf("a %s, it's price is %d元;\n", shoe.name, shoe.price)
}

type adidasCottonShoe struct {
	name  string
	price int
}

func (shoe adidasCottonShoe) Wear() {
	fmt.Printf("a %s, it's price is %d元;\n", shoe.name, shoe.price)
}

type adidasSportsShoe struct {
	name  string
	price int
}

func (shoe adidasSportsShoe) Wear() {
	fmt.Printf("a %s, it's price is %d元;\n", shoe.name, shoe.price)
}
