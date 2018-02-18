package client

import (
	"net/http"
	"io/ioutil"
	"bytes"
)

type WebClient struct {

}

func NewWebClient()*WebClient {
	return &WebClient{}
}

func (webC *WebClient) POST(url string, body []byte) (int, []byte, error) {
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, nil, err
	}

	return resp.StatusCode, data, nil
}