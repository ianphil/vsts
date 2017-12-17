package main

import (
	"github.com/tripdubroot/vsts/args"
	"github.com/tripdubroot/vsts/commands"
)

func main() {
	commandToExecute := args.ParseArgs()
	commands.Exec(commandToExecute)
}
