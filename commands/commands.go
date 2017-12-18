package commands

import (
	"os"

	"github.com/tripdubroot/vsts/commands/builds"
	"github.com/tripdubroot/vsts/commands/projects"
	"github.com/tripdubroot/vsts/commands/releases"
	"github.com/tripdubroot/vsts/utils/logger"
)

const (
	Project string = "project"
	Build   string = "build"
	Release string = "release"
	Empty   string = "empty"
	ShHelp  string = "-h"
	DhHelp  string = "--help"
)

var (
	cmd string
)

// ParseAndExecute is an exported function
func ParseAndExecute() {
	setCmd()

	switch {
	case cmd == Project:
		projects.Exec(os.Args)

	case cmd == Build:
		builds.Exec(os.Args)

	case cmd == Release:
		releases.Exec(os.Args)

	case cmd == Empty ||
		cmd == ShHelp ||
		cmd == DhHelp:
		printDefaultHelp()

	default:
		logger.Stderr(DefaultErrorText)
		os.Exit(1)
	}
}

func setCmd() {
	if len(os.Args) >= 2 {
		cmd = os.Args[1]
	} else {
		cmd = Empty
	}
}

func printDefaultHelp() {
	logger.Stdout(DefaultHelpText)
	os.Exit(0)
}
