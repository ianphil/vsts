package main

import (
	"vsts/args"
	"vsts/models/user"
	"vsts/utils/comms"
)

func main() {
	args.ValidateArgs()

	usr := user.User{ID: "Ian", Balance: 8}
	comms.PostRequest(usr)
}
