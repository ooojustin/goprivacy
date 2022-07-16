package goprivacy

// https://privacy.com/developer/docs

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
)

const BaseURL = "https://api.privacy.com/v1/"

type Client struct {
	Key string
}

func (pc Client) Request(method string, url string, params *map[string]interface{}) (*[]byte, error) {

	// initialize http client and request
	client := http.Client{}
	var req *http.Request
	if params == nil {
		req, _ = http.NewRequest(method, url, nil)
	} else {
		if bb, err := json.Marshal(params); err == nil {
			body := bytes.NewBuffer(bb)
			req, _ = http.NewRequest(method, url, body)
		}
	}

	// send request with api key
	req.Header.Add("Authorization", "api-key "+pc.Key)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	// handle failed status code
	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusMultipleChoices {
		return nil, errors.New("Failed, status code is " + strconv.Itoa(res.StatusCode))
	}

	// handle success by returning pointer to response body
	if bBody, err := io.ReadAll(res.Body); err == nil {
		return &bBody, nil
	} else {
		return nil, err
	}

}

func (pc Client) GET(url string) (*[]byte, error) {
	return pc.Request("GET", url, nil)
}

func (pc Client) POST(url string, params map[string]interface{}) (*[]byte, error) {
	return pc.Request("POST", url, &params)
}
