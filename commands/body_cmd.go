package commands

import (
	"encoding/json"

	"github.com/tripdubroot/vsts/models"
	"github.com/tripdubroot/vsts/utils/comms"
)

func execListProjects(cmd string) models.ProjectList {
	lstProjURI := comms.ListProjectsURI(cmd)
	data := comms.GetRequest(lstProjURI)

	var projList models.ProjectList
	err := json.Unmarshal([]byte(data), &projList)

	if err != nil {
		panic(err)
	}

	return projList
}
