package strategy

import "fmt"

// 算法接口 有 "被使用" 方法
type IAlg interface {
	Used()
}

/*****************
	算法 二分查询
 *****************/
type binarySearch struct {
	name string
}

func NewBinarySearch() binarySearch {
	binary_search := binarySearch{"binary search"}

	return binary_search
}

func (alg binarySearch) Used() {
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

func (alg quickSort) Used() {
	fmt.Println("I am using util " + alg.name)
}
