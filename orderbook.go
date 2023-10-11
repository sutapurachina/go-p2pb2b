package p2pb2b

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type OrderBookResp struct {
	Response
	Result      OrderBook `json:"result"`
	CacheTime   float64   `json:"cache_time"`
	CurrentTime float64   `json:"current_time"`
}

type OrderBook struct {
	Offset int              `json:"offset"`
	Limit  int              `json:"limit"`
	Total  int              `json:"total"`
	Orders []OrderBookEntry `json:"orders"`
}

type OrderBookEntry struct {
	ID        int     `json:"id"`
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

func (c *client) GetOrderBook(market string, side string, offset int64, limit int64) (*OrderBookResp, error) {
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

	url := fmt.Sprintf("%s/public/book?market=%s&side=%s&offset=%d&limit=%d", c.url, market, side, offset, limit)

	resp, err := c.sendGet(url, nil)
	if err != nil {
		return nil, err
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = checkHTTPStatus(*resp, http.StatusOK)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s: %s\n", err.Error(), string(bodyBytes)))
	}

	var result OrderBookResp
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
