package facade

type coffee struct {
}

func (coffee) Brew() {
	water := water{}
	water.boil()

	coffee_bean := coffeeBean{}
	coffee_bean.grind()

	coffee_powder := coffeePowder{}
	coffee_powder.brew()
}

func NewCoffee() coffee {
	coffee := coffee{}
	return coffee
}
