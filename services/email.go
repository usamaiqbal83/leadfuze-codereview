package services

import (
	"encoding/json"
	"net/http"
	"github.com/usamaiqbal83/leadfuze-codereview/domain"
	"github.com/pkg/errors"
)

const host = "http://165.227.202.198:5000"

type Options struct {
	WebClient domain.IWebClient
	Key string
}

type Email struct {
	WebClient domain.IWebClient
	Key string
}

type EmailVerifyResponse struct {
	Status   string `json:"status"`
	Info     string `json:"info"`
	Valid    bool   `json:"valid"`
	CatchAll bool   `json:"catchAll"`
}

func NewService(options *Options) *Email {
	return &Email{WebClient: options.WebClient, Key: options.Key}
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

	code, data, err := service.WebClient.POST(host, res)
	if err != nil || code != http.StatusOK {
		return false, errors.New("not verified")
	}

	response := EmailVerifyResponse{}

	// unmarshal response
	if err := json.Unmarshal(data, &response); err != nil {
		return false, err
	}

	return response.Valid, nil
}
