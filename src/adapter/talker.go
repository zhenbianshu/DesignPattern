package adapter

import "fmt"

type Talker interface {
	Express(word string)
}

type Person struct {
	Word string
}

func (person Person) Express() {
	fmt.Println(person.Word)
}
