package api

// https://privacy.com/developer/docs

import (
	"errors"
	"io"
	"net/http"
	"strconv"
)

const BaseURL = "https://api.privacy.com/v1/"

type PrivacyClient struct {
	Key string
}

func (pc PrivacyClient) GET(url string) (*[]byte, error) {

	client := http.Client{}
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Authorization", "api-key "+pc.Key)
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusMultipleChoices {
		return nil, errors.New("Failed, status code is " + strconv.Itoa(res.StatusCode))
	}

	if bBody, err := io.ReadAll(res.Body); err == nil {
		return &bBody, nil
	} else {
		return nil, err
	}

}
