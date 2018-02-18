package domain

type IWebClient interface {
	POST(url string, body []byte) (int, []byte, error)
}