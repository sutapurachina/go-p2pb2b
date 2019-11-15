package p2pb2b

import (
	"math"
	"net/http"
	"time"
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
	GetOrderBook(market string, side string, offset int64, limit int64) (*OrderBookResult, error)
	GetProducts() (*ProductsResult, error)
	GetSymbols() (*SymbolsResult, error)
}

type Result struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func TimestampToTime(timestamp float64) time.Time {
	sec, dec := math.Modf(timestamp)
	return time.Unix(int64(sec), int64(dec*(1e9)))
}
