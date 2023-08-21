package p2pb2b

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

type CreateOrderResp struct {
	Success   bool   `json:"success"`
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
	Result    Order  `json:"result"`
}

type Order struct {
	Amount    float64 `json:"amount,string"`
	DealFee   float64 `json:"dealFee,string"`
	DealMoney float64 `json:"dealMoney,string"`
	DealStock float64 `json:"dealStock,string"`
	Left      float64 `json:"left,string"`
	MakerFee  float64 `json:"makerFee,string"`
	Market    string  `json:"market"`
	OrderID   int64   `json:"orderId"`
	Price     float64 `json:"price,string"`
	Side      string  `json:"side"`
	TakerFee  float64 `json:"takerFee,string"`
	Timestamp float64 `json:"timestamp"`
	Type      string  `json:"type"`
}

type CreateOrderRequest struct {
	Request
	Market string  `json:"market"`
	Side   string  `json:"side"`
	Amount float64 `json:"amount,string"`
	Price  float64 `json:"price,string"`
}

type CancelOrderResp struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Result  Order  `json:"result"`
}

type CancelOrderRequest struct {
	Request
	Market  string `json:"market"`
	OrderID int64  `json:"orderId"`
}

type QueryUnexecutedRequest struct {
	Request
	Market string `json:"market"`
	Offset int64  `json:"offset"`
	Limit  int64  `json:"limit"`
}

type QueryUnexecutedResp struct {
	Success   bool              `json:"success"`
	Message   string            `json:"message"`
	ErrorCode string            `json:"errorCode"`
	Result    []UnexecutedOrder `json:"result"`
}

type UnexecutedOrder struct {
	Id        int     `json:"id"`
	Left      float64 `json:"left"`
	Market    string  `json:"market"`
	Amount    float64 `json:"amount"`
	Type      string  `json:"type"`
	Price     string  `json:"price"`
	Timestamp string  `json:"timestamp"`
	Side      string  `json:"side"`
	TakerFee  float64 `json:"takerFee"`
	MakerFee  float64 `json:"makerFee"`
	DealStock float64 `json:"dealStock"`
	DealMoney float64 `json:"dealMoney"`
}

type QueryExecutedRequest struct {
	Request
	Offset int64 `json:"offset"`
	Limit  int64 `json:"limit"`
}

type QueryExecutedResp struct {
	Response
	Result map[string][]AltOrder `json:"result"`
}

type AltOrder struct {
	Amount     float64 `json:"amount,string"`
	Price      float64 `json:"price,string"`
	Type       string  `json:"type"`
	ID         int64   `json:"id"`
	Source     string  `json:"source,omitempty"`
	Side       string  `json:"side"`
	Ctime      float64 `json:"ctime"`
	TakerFee   float64 `json:"takerFee,string"`
	Ftime      float64 `json:"ftime"`
	Market     string  `json:"market"`
	MakerFee   float64 `json:"makerFee,string"`
	DealFee    float64 `json:"dealFee,string"`
	DealStock  float64 `json:"dealStock,string"`
	DealMoney  float64 `json:"dealMoney,string"`
	MarketName string  `json:"marketName"`
}

type QueryDealsRequest struct {
	Request
	OrderID int64 `json:"orderId"`
	Offset  int64 `json:"offset"`
	Limit   int64 `json:"limit"`
}

type QueryDealsResp struct {
	Response
	Result QueryDealsResult `json:"result"`
}

type QueryDealsResult struct {
	Offset  int64    `json:"offset"`
	Limit   int64    `json:"limit"`
	Records []Record `json:"records"`
}

type Record struct {
	Time        float64 `json:"time"`
	Fee         float64 `json:"fee,string"`
	Price       float64 `json:"price,string"`
	Amount      float64 `json:"amount,string"`
	ID          int64   `json:"id"`
	DealOrderID int64   `json:"dealOrderId"`
	Role        int64   `json:"role"`
	Deal        float64 `json:"deal,string"`
}

func (c *client) CreateOrder(request *CreateOrderRequest) (*CreateOrderResp, error) {
	url := fmt.Sprintf("%s/order/new", c.url)
	request.Request.Nonce = strconv.FormatInt(time.Now().UnixMilli(), 10)
	request.Request.Request = "/api/v2/order/new"
	asJSON, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	resp, err := c.sendPost(url, nil, bytes.NewReader(asJSON))
	if err != nil {
		return nil, err
	}
	err = checkHTTPStatus(*resp, http.StatusOK)
	if err != nil {
		return nil, err
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result CreateOrderResp
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *client) CancelOrder(request *CancelOrderRequest) (*CancelOrderResp, error) {
	url := fmt.Sprintf("%s/order/cancel", c.url)
	request.Request.Nonce = strconv.FormatInt(time.Now().UnixMilli(), 10)
	request.Request.Request = "/api/v2/order/cancel"
	asJSON, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	resp, err := c.sendPost(url, nil, bytes.NewReader(asJSON))
	if err != nil {
		return nil, err
	}
	err = checkHTTPStatus(*resp, http.StatusOK)
	if err != nil {
		return nil, err
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result CancelOrderResp
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *client) QueryUnexecuted(request *QueryUnexecutedRequest) (*QueryUnexecutedResp, error) {
	url := fmt.Sprintf("%s/orders", c.url)
	request.Request.Nonce = strconv.FormatInt(time.Now().UnixMilli(), 10)
	request.Request.Request = "/api/v2/orders"
	asJSON, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	resp, err := c.sendPost(url, nil, bytes.NewReader(asJSON))
	if err != nil {
		return nil, err
	}
	err = checkHTTPStatus(*resp, http.StatusOK)
	if err != nil {
		return nil, err
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result QueryUnexecutedResp
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *client) QueryExecuted(request *QueryExecutedRequest) (*QueryExecutedResp, error) {
	url := fmt.Sprintf("%s/account/order_history", c.url)
	request.Request.Nonce = strconv.FormatInt(time.Now().UnixMilli(), 10)
	request.Request.Request = "/api/v2/account/order_history"
	asJSON, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	resp, err := c.sendPost(url, nil, bytes.NewReader(asJSON))
	if err != nil {
		return nil, err
	}
	err = checkHTTPStatus(*resp, http.StatusOK)
	if err != nil {
		return nil, err
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result QueryExecutedResp
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *client) QueryDeals(request *QueryDealsRequest) (*QueryDealsResp, error) {
	url := fmt.Sprintf("%s/account/order", c.url)
	request.Request.Nonce = strconv.FormatInt(time.Now().UnixMilli(), 10)
	request.Request.Request = "/api/v2/account/order"
	asJSON, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	resp, err := c.sendPost(url, nil, bytes.NewReader(asJSON))
	if err != nil {
		return nil, err
	}
	err = checkHTTPStatus(*resp, http.StatusOK)
	if err != nil {
		return nil, err
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result QueryDealsResp
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
