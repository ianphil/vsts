package tests

import (
	"testing"

	"github.com/tripdubroot/vsts/auth"
)

// Create env var by:
// export VSTS_PAT="your vsts PAT"
func TestGetTokenFromPat(t *testing.T) {
	result := auth.GetTokenFromPAT()
	if result == "" {
		t.Errorf("Returned data was incorrect")
	}
}
