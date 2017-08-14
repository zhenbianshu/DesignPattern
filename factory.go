package main

import (
	"factory"
)

func main() {
	lining_shoes_factory := factory.LiNingFactory{}
	lining_shoe := lining_shoes_factory.Product("plastic")
	lining_shoe.Wear() // a lining plastic shoe, it's price is 120元;

	adidas_shoe_factory := factory.AdidasFactory{}
	adidas_shoe := adidas_shoe_factory.Product("sports")
	adidas_shoe.Wear() // a adidas sports shoe, it's price is 70元;
}
