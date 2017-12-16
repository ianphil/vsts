package comms

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	u "vsts/models/user"
)

// PostRequest is an exported type
func PostRequest(usr u.User) {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(usr)
	res, _ := http.Post("https://requestb.in/19r9tq41", "application/json; charset=utf-8", b)
	io.Copy(os.Stdout, res.Body)
	fmt.Println(" ")
}
