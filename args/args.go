package args

import (
	"os"

	"github.com/tripdubroot/vsts/utils/logger"
)

// ValidateArgs is an exported function
func ValidateArgs() string {
	var firstArg string
	var result string

	if len(os.Args) >= 2 {
		firstArg = os.Args[1]
	} else {
		firstArg = "empty"
	}

	switch {
	case firstArg == "build":
		result = validateBuildArg(os.Args)
	case firstArg == "release":
		result = validateReleaseArg(os.Args)
	case firstArg == "empty" || firstArg == "-h" || firstArg == "--help":
		printDefaultHelp()
	default:
		logger.LogMessage("Please supply a valid command, -h, or --help switch")
		os.Exit(1)
	}

	return result
}
