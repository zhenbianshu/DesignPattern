package visitor

import "strconv"

type scanner struct {
}

func (scanner scanner) VisitGoods(cart *shoppingCart) string {
	princes := ""
	for _, goods := range cart.bought {
		princes += strconv.Itoa(goods.price) + "元 "
	}

	return princes
}

func NewScanner() *scanner {
	scanner := &scanner{}

	return scanner
}
