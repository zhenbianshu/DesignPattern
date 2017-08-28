package main

import (
	"fmt"
	"prototype"
)

func main() {
	guard_sample := prototype.NewhonorGuard(18, 180, 72, "m")

	zhangsan := guard_sample.Clone()
	zhangsan.SetName("zhangsan")
	fmt.Println(zhangsan) // &{18 180 72 m zhangsan}

	lisi := guard_sample.Clone()
	lisi.SetName("lisi")
	fmt.Println(lisi) // &{18 180 72 m lisi}
}
