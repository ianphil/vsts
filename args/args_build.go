package args

import (
	"os"
	"vsts/utils/logger"
)

// validateBuildArg parses the build command
func validateBuildArg(args []string) {
	logger.LogMessage("Build arg")
	os.Exit(0)
}
