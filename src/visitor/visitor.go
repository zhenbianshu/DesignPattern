package visitor

type visitor interface {
	visitGoods(cart *shoppingCart) string
}
