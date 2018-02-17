package services

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const host = "http://165.227.202.198:5000"

type Email struct {
	Key string
}

type EmailVerifyResponse struct {
	Status   string `json:"status"`
	Info     string `json:"info"`
	Valid    bool   `json:"valid"`
	CatchAll bool   `json:"catchAll"`
}

func NewService(key string) *Email {
	return &Email{Key: key}
}

func (service *Email) VerifyEmail(email string) (bool, error) {

	values := map[string]string{
		"email": email,
		"key":   service.Key,
	}

	res, err := json.Marshal(values)
	if err != nil {
		return false, err
	}

	resp, err := http.Post(host, "application/json", bytes.NewBuffer(res))
	if err != nil {
		return false, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	response := EmailVerifyResponse{}

	// unmarshal response
	if err := json.Unmarshal(body, &response); err != nil {
		return false, err
	}

	return response.Valid, nil
}
