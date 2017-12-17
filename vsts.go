package main

import (
	"fmt"

	"github.com/tripdubroot/vsts/args"
	"github.com/tripdubroot/vsts/auth"
	"github.com/tripdubroot/vsts/commands"
)

func main() {
	auth := auth.GetTokenFromPAT()

	fmt.Println(auth)

	commandToExecute := args.ValidateArgs()
	commands.Exec(commandToExecute)
}
