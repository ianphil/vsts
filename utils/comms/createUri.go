package comms

import "fmt"

const (
	projectVer string = "?api-version=1.0"
)

// InstanceURI generates URI host/project
func InstanceURI(instance string) string {
	uri := fmt.Sprintf("http://%s.visualstudio.com", instance)

	return uri
}

// ListProjectsURI does this:
// https://www.visualstudio.com/en-us/docs/integrate/api/tfs/projects#get-a-list-of-team-projects
func ListProjectsURI(instance string) string {
	instance = InstanceURI(instance)
	uri := fmt.Sprintf("%s/_apis/projects%s", instance, projectVer)

	return uri
}
