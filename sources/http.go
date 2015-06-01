package sources

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HttpProgressChecker struct {
	client *http.Client
}

type HookResponse struct {
	Count int64 `json:"process_count"`
}

func NewHttpProgressChecker() *HttpProgressChecker {
	return &HttpProgressChecker{client: &http.Client{}}
}

func (self *HttpProgressChecker) CheckProgress() (bool, error) {
	url := fmt.Sprintf("http://%s:%d%s", *host, *port, *context)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false, err
	}

	resp := &HookResponse{}

	err = doPostRequestAndGetValue(self.client, req, resp)
	if err != nil {
		return false, err
	}

	return (resp.Count == 0), nil
}

func doPostRequestAndGetValue(client *http.Client, req *http.Request, value interface{}) error {
	response, err := client.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	dec := json.NewDecoder(response.Body)
	dec.UseNumber()
	err = dec.Decode(value)
	if err != nil {
		return err
	}
	return nil
}
