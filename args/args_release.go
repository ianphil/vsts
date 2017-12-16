package args

import (
	"os"
	"vsts/utils/logger"
)

// validateReleaseArg parses the release command
func validateReleaseArg(args []string) {
	logger.LogMessage("release arg")
	os.Exit(0)
}