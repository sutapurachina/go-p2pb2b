package p2pb2b

import (
	"net/http"
)

func NewClient() (Client, error) {
	return &client{
		http: &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		},
		auth: nil,
	}, nil
}

type Client interface {
	GetTickers() (*Tickers, error)
}
