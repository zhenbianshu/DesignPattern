package composite

import "fmt"

type Military interface {
	Fight()
}

type General struct {
	Name      string
	Followers []Captain
}

func (general General) Fight() {
	fmt.Println(general.Name + " is fighting!")
	for _, captain := range general.Followers {
		captain.Fight()
	}
}

type Captain struct {
	Name      string
	Followers []Soldier
}

func (captain Captain) Fight() {
	fmt.Println(captain.Name + " is fighting")
	for _, soldier := range captain.Followers {
		soldier.Fight()
	}
}

type Soldier struct {
	Name string
}

func (soldier Soldier) Fight() {
	fmt.Println(soldier.Name + " is fighting!")
}
