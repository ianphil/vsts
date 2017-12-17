package commands

import (
	"os"
	"fmt"

	"github.com/tripdubroot/vsts/commands"
	"github.com/tripdubroot/vsts/utils/logger"
)

// // Exec executes the command passed by args
// func Exec(cmd string) {
// 	// var result models.ProjectList
// 	// result = execListProjects(cmd)

// 	// for i := 1; i < result.Count; i++ {
// 	// 	fmt.Println(result.Value[i].Name)
// 	// }

// 	fmt.Println(cmd)
// }

var firstArg string

// ParseArgs is an exported function
func ParseAndExecute() {
	var result string

	setFirstArg()

	switch {
	case firstArg == commands.Project:
		result = parseProjectArg(os.Args)

	case firstArg == commands.Build:
		result = parseBuildArg(os.Args)

	case firstArg == commands.Release:
		result = parseReleaseArg(os.Args)

	case firstArg == commands.Empty ||
		firstArg == commands.ShHelp ||
		firstArg == commands.DhHelp:
		printDefaultHelp()

	default:
		logger.Stderr(DefaultErrorText)
		os.Exit(1)
	}

	return result
}

func setFirstArg() {
	if len(os.Args) >= 2 {
		firstArg = os.Args[1]
	} else {
		firstArg = commands.Empty
	}
}
