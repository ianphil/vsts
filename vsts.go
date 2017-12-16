package main

import (
	"github.com/tripdubroot/vsts/args"
	"github.com/tripdubroot/vsts/models/user"
	"github.com/tripdubroot/vsts/utils/comms"
)

func main() {
	args.ValidateArgs()

	usr := user.User{ID: "Ian", Balance: 8}
	comms.PostRequest(usr)
}
