package visitor

type goods struct {
	name  string
	price int
}

func NewGoods(name string, price int) *goods {
	goods := &goods{name, price}
	return goods
}

func (goods *goods) AddToCart(cart *shoppingCart) {
	cart.bought = append(cart.bought, goods)
}

type shoppingCart struct {
	bought []*goods
}

func NewShoppingCart() *shoppingCart {
	bought := make([]*goods, 0)
	cart := &shoppingCart{bought}
	return cart
}
