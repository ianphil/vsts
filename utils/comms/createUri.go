package comms

import "fmt"

const (
	projectVer string = "?api-version=1.0"
	Host       string = "https://%s.visualstudio.com"
)

// InstanceURI generates URI host/project
func InstanceURI(instance string) string {
	uri := fmt.Sprintf(Host, instance)

	return uri
}

// ListProjectsURI does this:
// https://www.visualstudio.com/en-us/docs/integrate/api/tfs/projects#get-a-list-of-team-projects
func ListProjectsURI(instance string) string {
	instance = InstanceURI(instance)
	uri := fmt.Sprintf("%s/_apis/projects%s", instance, projectVer)

	return uri
}
