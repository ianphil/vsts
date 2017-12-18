package cmd

import (
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/tripdubroot/vsts/models"
	"github.com/tripdubroot/vsts/utils/comms"
	"github.com/tripdubroot/vsts/utils/logger"
)

var instance string

// projectListCmd represents the projectList command
var projectListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all projects in VSTS instance",
	Long:  "List all projects in VSTS instance longer",
	Run: func(cmd *cobra.Command, args []string) {
		lstProjURI := comms.ListProjectsURI(instance)
		data := comms.GetRequest(lstProjURI)

		var projList models.ProjectList
		err := json.Unmarshal([]byte(data), &projList)
		if err != nil {
			panic(err)
		}

		logger.Stdout("Project Name: " + projList.Value[0].Name)
	},
}

func init() {
	projectCmd.AddCommand(projectListCmd)
	projectListCmd.Flags().StringVarP(&instance, "instance", "i", "", "Name of the VSTS instance")
}
