package main

import (
	"decorator"
	"fmt"
)

func main() {
	gift := decorator.NewGift("flower")
	gift_packed_first := decorator.NewPack(gift, "red ribbon")
	gift_packed_second := decorator.NewPack(gift_packed_first, "box")

	fmt.Println("a good sold, its detail:" + gift_packed_second.GetDesc()) // a good sold, its detail:flower, with red ribbon packed, with box packed
}
