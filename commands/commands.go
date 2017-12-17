package commands

import (
	"os"

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
	setFirstArg()

	switch {
	case firstArg == Project:
		parseProjectArg(os.Args)

	case firstArg == Build:
		parseBuildArg(os.Args)

	case firstArg == Release:
		parseReleaseArg(os.Args)

	case firstArg == Empty ||
		firstArg == ShHelp ||
		firstArg == DhHelp:
		printDefaultHelp()

	default:
		logger.Stderr(DefaultErrorText)
		os.Exit(1)
	}
}

func setFirstArg() {
	if len(os.Args) >= 2 {
		firstArg = os.Args[1]
	} else {
		firstArg = Empty
	}
}
