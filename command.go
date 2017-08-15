package main

import "command"

func main() {
	// this is a general, and publish command to a soldier
	command_run := command.Run{"tree", "grass_ground"}
	soldier_A := command.Soldier{"zhangsan"}
	soldier_A.ExecCommand(command_run) // I am zhangsan, I will execute general's command: run to grass_ground from tree
	soldier_A.UndoCommand(command_run) // I am zhangsan, I will execute general's command: run to tree from grass_ground

	command_attention := command.Attention{}
	soldier_A.ExecCommand(command_attention) // I am zhangsan, I will execute general's command: Attention!
	soldier_A.UndoCommand(command_attention) // I am zhangsan, I will execute general's command: At ease!
}
