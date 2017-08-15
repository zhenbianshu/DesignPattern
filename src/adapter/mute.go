package adapter

import "fmt"

type Mute struct {
	Word string
}

type NoteBook struct {
	User Mute
}

func (notebook NoteBook) Express() {
	fmt.Println(notebook.User.Word)
}
