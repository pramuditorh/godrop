package helpers

import (
	"bytes"
	"net/http"
	"os"
)

var (
	client *http.Client
)

func init() {
	client = &http.Client{}
}

func DoReq(method string, url string, body []byte) (*http.Response, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	var bearer = "Bearer " + os.Getenv("DIGITALOCEAN_TOKEN")
	req.Header.Set("Authorization", bearer)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil

}