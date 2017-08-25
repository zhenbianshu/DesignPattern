package main

import "flyweight"

func main() {
	treeA := map[string]int{"height": 20}
	treeB := map[string]int{"height": 14}
	treeC := map[string]int{"height": 25}
	treeD := map[string]int{"height": 7}
	treeE := map[string]int{"height": 40}
	trees := []map[string]int{treeA, treeB, treeC, treeD, treeE}

	forest := flyweight.NewForest(trees)

	forest.Fell()
	/*
		a 20 m tree cutted!
		a 14 m tree cutted!
		a 25 m tree cutted!
		a 7 m tree cutted!
		a 40 m tree cutted!
	*/
}
