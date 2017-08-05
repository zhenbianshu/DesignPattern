package strategy

import "fmt"

// 算法接口 有 "被使用" 方法
type IAlg interface {
	Used()
}

/*****************
	算法 二分查询
 *****************/
type BinarySearch struct {
	name string
}

func NewBinarySearch() BinarySearch {
	binary_search := BinarySearch{"binary search"}

	return binary_search
}

func (alg BinarySearch) Used() {
	fmt.Println("I am using util " + alg.name)
}

/*****************
	算法 快速排序法
 *****************/
type QuickSort struct {
	name string
}

func NewQuickSort() QuickSort {
	quick_sort := QuickSort{"quick sort"}

	return quick_sort
}

func (alg QuickSort) Used() {
	fmt.Println("I am using util " + alg.name)
}
