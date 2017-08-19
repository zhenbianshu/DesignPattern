package main

import "strategy"

func main() {
	calculator := strategy.NewCalculator()
	calculator.Cal() // I am using util bubble sort

	quick_sort := strategy.NewQuickSort()
	calculator.ChangeAlg(quick_sort)
	calculator.Cal() // I am using util quick sort
}
