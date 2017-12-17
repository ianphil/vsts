package comms

import (
	"io/ioutil"
	"net/http"

	"github.com/tripdubroot/vsts/auth"
)

// GetRequest gets data
func GetRequest() string {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://cseblack.visualstudio.com/DefaultCollection/_apis/projects?api-version=3.0", nil)
	req.Header.Set("Authorization", "Basic "+auth.GetTokenFromPAT())
	res, _ := client.Do(req)

	defer res.Body.Close()

	var body string
	if res.StatusCode == http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(res.Body)
		body = string(bodyBytes)
	} else {
		body = "Didn't work"
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
