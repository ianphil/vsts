package args

import (
	"os"

	"github.com/tripdubroot/vsts/utils/logger"
)

// ValidateArgs is an exported function
func ValidateArgs() {
	var firstArg string

	if len(os.Args) >= 2 {
		firstArg = os.Args[1]
	} else {
		firstArg = "empty"
	}

	switch {
	case firstArg == "build":
		validateBuildArg(os.Args)
	case firstArg == "release":
		validateReleaseArg(os.Args)
	case firstArg == "empty" || firstArg == "-h" || firstArg == "--help":
		printDefaultHelp()
	default:
		logger.LogMessage("Please supply a valid command, -h, or --help switch")
		os.Exit(1)
	}

	// if len(args) == 1 || args[1] == "-h" || args[1] == "--help" {
	// 	printDefaultHelp()
	// }
}
