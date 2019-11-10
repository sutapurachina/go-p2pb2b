package p2pb2b

import (
	"net/http"
)

const base_api = "https://api.p2pb2b.io/api/v1"

// for testing purposes only
func newClientWithURL(url string, apiKey string, apiSecret string) (Client, error) {
	return &client{
		http: &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		},
		auth: &auth{
			APIKey:    apiKey,
			APISecret: apiSecret,
		},
		url: url,
	}, nil
}

func NewClient(apiKey string, apiSecret string) (Client, error) {
	return newClientWithURL(base_api, apiKey, apiSecret)
}

type Client interface {
	PostCurrencyBalance(request *AccountBalanceRequest) (*AccountBalanceResult, error)
	GetMarkets() (*MarketsResult, error)
	GetTickers() (*TickersResult, error)
	GetTicker(market string) (*Ticker, error)
	GetProducts() (*ProductsResult, error)
	GetSymbols() (*SymbolsResult, error)
}

func (c *client) withURL(url string) {
	c.url = url
}
