package visitor

type printer struct {
}

func (printer printer) VisitGoods(cart *shoppingCart) string {
	names := ""
	for _, goods := range cart.bought {
		names += goods.name + " "
	}

	return names
}

func NewPrinter() *printer {
	printer := &printer{}

	return printer
}
