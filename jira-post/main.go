package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var data = `{
    "fields": {
        "project": {
            "key": "ASX"
        },
        "summary": "TEST ABC s",
        "description": "Creating of an issue using project keys and issue type names using the REST API",
        "issuetype": {
            "name": "Bug"
        }
    }
}`

func main() {
	_main()
}

func _main() {
	/*body := bytes.NewBuffer([]byte(data))
	err := request("POST", "https://leevmo.atlassian.net/rest/api/2/issue", body)
	fmt.Println(err)*/
	res, err := request("GET", "https://leevmo.atlassian.net/rest/api/2/project", nil)
	var issue IssueDetail
	body,_ := io.ReadAll(res.Body)
	fmt.Println(string(body))
	err = json.Unmarshal(body, &issue)
	fmt.Println(issue, err, issue.ID, issue.Self, issue.Key)
}

func request(method, url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	req.Header.Set("Authorization","Basic dHJvbmcubGVAdm1vZGV2LmNvbTpvMGFobk13Vm9PNldHTDFHdTF5bzVGQUY=")
	req.Header.Set("content-type", "application/json")
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	return client.Do(req)
}
