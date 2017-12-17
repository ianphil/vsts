package main

import (
	"github.com/tripdubroot/vsts/commands"
)

func main() {
	commandToExecute := "cseblack" //args.ValidateArgs()
	commands.Exec(commandToExecute)
}
