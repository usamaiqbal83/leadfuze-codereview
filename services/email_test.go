package services

import (
	"testing"
	"net/http"
	"encoding/json"
	"github.com/pkg/errors"
)

type MockWebClient struct {
	mockPost func(url string, body []byte) (int, []byte, error)
}

func (mock *MockWebClient) POST(url string, body []byte) (int, []byte, error) {
	if mock.mockPost != nil {
		return mock.mockPost(url, body)
	}
	return http.StatusBadRequest, nil, nil
}

func TestVerifyEmailWhenEmailIsValid(t *testing.T) {
	mockClient := MockWebClient{mockPost: func(url string, body []byte) (int, []byte, error) {
		respObj := EmailVerifyResponse{Valid: true}
		res, err := json.Marshal(respObj)
		if err != nil {
			return http.StatusBadRequest, nil, err
		}

		return http.StatusOK, res, nil
	}}

	service := NewService(&Options{Key:"1234", WebClient: &mockClient})

	res, err := service.VerifyEmail("some@valid.com")
	if err != nil {
		t.Fatalf("expected no error but got an error")
	}

	if res != true {
		t.Fatalf("expected true got false")
	}
}

func TestVerifyEmailWhenEmailIsNotValid(t *testing.T) {
	mockClient := MockWebClient{mockPost: func(url string, body []byte) (int, []byte, error) {
		respObj := EmailVerifyResponse{Valid: false}
		res, err := json.Marshal(respObj)
		if err != nil {
			return http.StatusBadRequest, nil, err
		}

		return http.StatusOK, res, nil
	}}

	service := NewService(&Options{Key:"1234", WebClient: &mockClient})

	res, err := service.VerifyEmail("some@invalid.com")
	if err != nil {
		t.Fatalf("expected no error but got an error")
	}

	if res == true {
		t.Fatalf("expected true got false")
	}
}

func TestVerifyEmailWhenServiceIsNotAccessible(t *testing.T) {
	mockClient := MockWebClient{mockPost: func(url string, body []byte) (int, []byte, error) {
		return http.StatusBadRequest, nil, errors.New("some error")
	}}

	service := NewService(&Options{Key:"1234", WebClient: &mockClient})

	_, err := service.VerifyEmail("some@email.com")
	if err == nil {
		t.Fatalf("expected error but got no error")
	}
}

