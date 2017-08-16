package facade

import "fmt"

type water struct {
}

func (water) boil() {
	fmt.Println("water boiled;")
}

type coffeeBean struct {
}

func (coffeeBean) grind() {
	fmt.Println("coffee bean grind;")
}

type coffeePowder struct {
}

func (coffeePowder) brew() {
	fmt.Println("coffee is done!")
}
