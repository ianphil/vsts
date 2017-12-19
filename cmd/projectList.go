package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tripdubroot/vsts/models"
	"github.com/tripdubroot/vsts/utils/comms"
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

		fmt.Println(" ")
		fmt.Printf(" %-45s| %-20s\n", "ID", "Name")
		fmt.Println("------------------------------------------------------------------")

		for i := 0; i < projList.Count; i++ {
			fmt.Printf(" %-45s  %-20s\n", projList.Value[i].ID, projList.Value[i].Name)
		}

		fmt.Println(" ")
	},
}

func init() {
	projectCmd.AddCommand(projectListCmd)
	projectListCmd.Flags().StringVarP(&instance, "instance", "i", "", "Name of the VSTS instance")
}
