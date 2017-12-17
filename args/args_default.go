package args

import (
	"os"

	"github.com/tripdubroot/vsts/utils/logger"
)

// printDefaultHelp displays the default help command
func printDefaultHelp() {
	logger.Stdout(DefaultHelpText)
	os.Exit(0)
}
