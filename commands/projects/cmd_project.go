package projects

import (
	"encoding/json"
	"os"

	"github.com/tripdubroot/vsts/utils/comms"

	"github.com/tripdubroot/vsts/models"
	"github.com/tripdubroot/vsts/utils/logger"
)

const (
	List             string = "list"
	Empty            string = "empty"
	ShHelp           string = "-h"
	DhHelp           string = "--help"
	DefaultErrorText string = "Project Error Text"
)

var (
	subCmd string
)

// Exec parses cmds and executes
func Exec(args []string) {
	setSubCmd()

	switch {
	case subCmd == List:
		listProjects(args)

	case subCmd == Empty ||
		subCmd == ShHelp ||
		subCmd == DhHelp:
		printProjectHelp()

	default:
		logger.Stderr(DefaultErrorText)
		os.Exit(1)
	}
}

func setSubCmd() {
	if len(os.Args) >= 3 {
		subCmd = os.Args[2]
	} else {
		subCmd = Empty
	}
}

func printProjectHelp() {
	logger.Stdout("Project Help Text")
}

func listProjects(args []string) {
	cmd := args[4]

	lstProjURI := comms.ListProjectsURI(cmd)
	data := comms.GetRequest(lstProjURI)

	var projList models.ProjectList
	err := json.Unmarshal([]byte(data), &projList)
	if err != nil {
		panic(err)
	}

	logger.Stdout("Project Name: " + projList.Value[0].Name)

	os.Exit(0)
}
