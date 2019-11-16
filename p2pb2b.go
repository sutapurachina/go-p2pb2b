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
	PostCurrencyBalance(request *AccountCurrencyBalanceRequest) (*AccountCurrencyBalanceResp, error)
	PostBalances(request *AccountBalancesRequest) (*AccountBalancesResp, error)
	CreateOrder(request *CreateOrderRequest) (*CreateOrderResp, error)
	CancelOrder(request *CancelOrderRequest) (*CancelOrderResp, error)
	QueryUnexecuted(request *QueryUnexecutedRequest) (*QueryUnexecutedResp, error)
	QueryExecuted(request *QueryExecutedRequest) (*QueryExecutedResp, error)
	QueryDeals(request *QueryDealsRequest) (*QueryDealsResp, error)
	GetMarkets() (*MarketsResp, error)
	GetTickers() (*TickersResp, error)
	GetTicker(market string) (*TickerResp, error)
	GetOrderBook(market string, side string, offset int64, limit int64) (*OrderBookResp, error)
	GetHistory(market string, lastID int64, limit int64) (*HistoryResp, error)
	GetDepthResult(market string, limit int64) (*DepthResultResp, error)
	GetProducts() (*ProductsResp, error)
	GetSymbols() (*SymbolsResp, error)
}

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type Request struct {
	Request string `json:"request"`
	Nonce   string `json:"nonce"`
}

func TimestampToTime(timestamp float64) time.Time {
	sec, dec := math.Modf(timestamp)
	return time.Unix(int64(sec), int64(dec*(1e9)))
}
