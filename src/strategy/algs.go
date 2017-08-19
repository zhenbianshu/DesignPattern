package strategy

import "fmt"

// 算法接口 有 "被使用" 方法
type IAlg interface {
	used()
}

/*****************
	算法 冒泡
 *****************/
type BubbleSort struct {
	name string
}

func NewBubbleSearch() BubbleSort {
	binary_search := BubbleSort{"bubble sort"}

	return binary_search
}

func (alg BubbleSort) used() {
	fmt.Println("I am using util " + alg.name)
}

/*****************
	算法 快速排序法
 *****************/
type quickSort struct {
	name string
}

func NewQuickSort() quickSort {
	quick_sort := quickSort{"quick sort"}

	return quick_sort
}

func (alg quickSort) used() {
	fmt.Println("I am using util " + alg.name)
}
