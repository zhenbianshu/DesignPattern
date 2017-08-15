package command

import "fmt"

type Soldier struct {
	Name string
}

func (soldier Soldier) ExecCommand(command Command) {
	fmt.Print("I am " + soldier.Name + ", I will execute general's command: ")
	command.exec()
}

func (soldier Soldier) UndoCommand(command Command) {
	fmt.Print("I am " + soldier.Name + ", I will execute general's command: ")
	command.undo()
}
