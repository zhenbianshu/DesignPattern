package main

import (
	"fmt"
	"iterator"
)

func main() {
	horses := []string{"white", "black", "brown"}
	horse_queue := iterator.NewHorses(horses)

	soldiers :=
		map[string]string{
			"zhangsan": "zhangsan",
			"lisi":     "lisi",
			"wangwu":   "wangwu",
		}
	soldier_queue := iterator.NewSoldiers(soldiers)

	fmt.Println(horse_queue.Count(), soldier_queue.Count())
	fmt.Println(horse_queue.Next(), soldier_queue.Next())
	horse_queue.Remove()
	soldier_queue.Remove()
}
