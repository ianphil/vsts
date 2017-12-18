package comms

import (
	"io/ioutil"
	"net/http"

	"github.com/tripdubroot/vsts/utils/logger"

	"github.com/tripdubroot/vsts/auth"
)

// GetRequest gets data
func GetRequest(uri string) string {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", uri, nil)
	req.Header.Set("Authorization", "Basic "+auth.GetTokenFromPAT())
	res, _ := client.Do(req)

	defer res.Body.Close()

	var body string
	if res.StatusCode == http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(res.Body)
		body = string(bodyBytes)
	} else {
		logger.Stderr("Not Authorized")
	}

	return body
}

// PostRequest is an exported type
// func PostRequest(usr u.User) {
// 	b := new(bytes.Buffer)
// 	json.NewEncoder(b).Encode(usr)
// 	res, _ := http.Post("https://requestb.in/19r9tq41", "application/json; charset=utf-8", b)
// 	io.Copy(os.Stdout, res.Body)
// 	fmt.Println(" ")
// }
