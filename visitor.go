package main

import (
	"fmt"
	"visitor"
)

func main() {
	cart := visitor.NewShoppingCart()
	water := visitor.NewGoods("nofushanquan", 2)
	water.AddToCart(cart)

	milk := visitor.NewGoods("wangzai", 5)
	milk.AddToCart(cart)

	ballPen := visitor.NewGoods("morning light", 1)
	ballPen.AddToCart(cart)

	scanner := visitor.NewScanner()
	fmt.Println("goods's prices are:" + scanner.VisitGoods(cart)) // goods's prices are:2元 5元 1元

	printer := visitor.NewPrinter()
	fmt.Println("goods's names are:" + printer.VisitGoods(cart)) // goods's names are:nofushanquan wangzai morning light

}
