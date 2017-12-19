package tests

import (
	"testing"

	"github.com/tripdubroot/vsts/utils/comms"
)

const (
	instance               string = "cseblack"
	expectedInstanceURI    string = "https://cseblack.visualstudio.com"
	expectedListProjectURI string = "https://cseblack.visualstudio.com/_apis/projects?api-version=1.0"
)

func TestInstanceURI(t *testing.T) {
	result := comms.InstanceURI(instance)

	if result != expectedInstanceURI {
		t.Errorf("Test Error: TestInstanceURI")
	}
}

func TestListProjectsURI(t *testing.T) {
	result := comms.ListProjectsURI(instance)

	if result != expectedListProjectURI {
		t.Errorf("Test Error: TestListProjectsURI")
	}
}
