package main

import "facade"

func main() {
	coffee := facade.NewCoffee()
	coffee.Brew()
	// water boiled;
	// coffee bean grind;
	// coffee is done!
}
