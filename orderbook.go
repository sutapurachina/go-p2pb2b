package p2pb2b

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type OrderBookResult struct {
	Result
	OrderBook OrderBook `json:"result"`
}

type OrderBook struct {
	Offset int64            `json:"offset"`
	Limit  int64            `json:"limit"`
	Total  int64            `json:"total"`
	Orders []OrderBookEntry `json:"orders"`
}

type OrderBookEntry struct {
	ID        int64   `json:"id"`
	Left      float64 `json:"left,string"`
	Market    string  `json:"market"`
	Amount    float64 `json:"amount,string"`
	Type      string  `json:"type"`
	Price     float64 `json:"price,string"`
	Timestamp float64 `json:"timestamp"`
	Side      string  `json:"side"`
	DealFee   float64 `json:"dealFee,string"`
	TakerFee  float64 `json:"takerFee,string"`
	MakerFee  float64 `json:"makerFee,string"`
	DealStock float64 `json:"dealStock,string"`
	DealMoney float64 `json:"dealMoney,string"`
}

func (c *client) GetOrderBook(market string, side string, offset int64, limit int64) (*OrderBookResult, error) {
	if market == "" {
		return nil, fmt.Errorf("parameter market must not be empty")
	}
	if side == "" {
		return nil, fmt.Errorf("parameter side must not be empty")
	}
	if offset < 0 {
		return nil, fmt.Errorf("parameter offset must not be < 0")
	}
	if limit <= 0 {
		return nil, fmt.Errorf("parameter limit must not be <= 0")
	}
	v := url.Values{}
	v.Set("market", market)
	v.Add("side", side)
	v.Add("offset", fmt.Sprintf("%d", offset))
	v.Add("limit", fmt.Sprintf("%d", limit))

	url := fmt.Sprintf("%s/public/book?%s", c.url, v.Encode())

	resp, err := c.sendGet(url, nil)
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

	var result OrderBookResult
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
