package p2pb2b

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type OrdersCreateResult struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Order   Order  `json:"result"`
}

type Order struct {
	OrderID   int64   `json:"orderId"`
	Market    string  `json:"market"`
	Price     float64 `json:"price,string"`
	Side      string  `json:"side"`
	Type      string  `json:"type"`
	Timestamp float64 `json:"timestamp"`
	DealMoney float64 `json:"dealMoney,string"`
	DealStock float64 `json:"dealStock,string"`
	Amount    float64 `json:"amount,string"`
	TakerFee  float64 `json:"takerFee,string"`
	MakerFee  float64 `json:"makerFee,string"`
	Left      float64 `json:"left,string"`
	DealFee   float64 `json:"dealFee,string"`
}

type OrdersCreateRequest struct {
	Market  string  `json:"market"`
	Side    string  `json:"side"`
	Amount  float64 `json:"amount,string"`
	Price   float64 `json:"price,string"`
	Request string  `json:"request"`
	Nonce   string  `json:"nonce"`
}

func (c *client) PostCreateOrder(request *OrdersCreateRequest) (*OrdersCreateResult, error) {
	url := fmt.Sprintf("%s/order/new", c.url)
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

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if http.StatusOK != resp.StatusCode {
		return nil, fmt.Errorf("status: %s, body: %s", resp.Status, bodyBytes)
	}

	var result OrdersCreateResult
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
