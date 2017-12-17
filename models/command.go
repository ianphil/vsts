package models

// Command represents arguments passed to vsts_cli
type Command struct {
	Command     string
	Subcommands []string
}
