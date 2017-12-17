package main

import (
	"fmt"

	"github.com/tripdubroot/vsts/args"
	"github.com/tripdubroot/vsts/commands"
	"github.com/tripdubroot/vsts/utils/comms"
)

func main() {
	result := comms.GetRequest()

	fmt.Println(result)

	commandToExecute := args.ValidateArgs()
	commands.Exec(commandToExecute)
}
