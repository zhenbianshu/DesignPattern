package main

import (
	"fmt"
	"memento"
)

func main() {
	time_controller := memento.NewTimeController()
	xiaoming := memento.NewPerson("xiaoming", 18, 175, 70)
	time_controller.Save(xiaoming.GetAge(), xiaoming.Bak())

	xiaoming.GrowUp(80, 170, 65)
	fmt.Println(xiaoming) // &{xiaoming 80 170 65} 小明变成老明了
	xiaoming.Renewal(time_controller.GetMemo(18))
	fmt.Println(xiaoming) // &{xiaoming 18 175 70} 小明又成为小明了
}
