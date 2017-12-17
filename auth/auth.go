package auth

import (
	"encoding/base64"
	"os"
)

// GetTokenFromPAT gets barer token from VSTS_PAT env var
func GetTokenFromPAT() string {
	data := []byte(os.Getenv("VSTS_PAT"))
	token := base64.StdEncoding.EncodeToString(data)

	return token
}
