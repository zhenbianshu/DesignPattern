package command

import "fmt"

type Command interface {
	exec()
	undo()
}

type Run struct {
	Start       string
	Destination string
}

func (command Run) exec() {
	fmt.Println("run to " + command.Destination + " from " + command.Start)
}

func (command Run) undo() {
	fmt.Println("run to " + command.Start + " from " + command.Destination)
}

type Attention struct {
}

func (command Attention) exec() {
	fmt.Println("Attention!")
}

func (command Attention) undo() {
	fmt.Println("At ease!")
}
