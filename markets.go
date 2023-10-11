package p2pb2b

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type MarketsResp struct {
	Response
	Result      []Market `json:"result"`
	CacheTime   float64  `json:"cache_time"`
	CurrentTime float64  `json:"current_time"`
}

type Market struct {
	Name      string  `json:"name"`
	MoneyPrec int     `json:"moneyPrec,string"`
	Stock     string  `json:"stock"`
	Money     string  `json:"money"`
	StockPrec int     `json:"stockPrec,string"`
	FeePrec   int     `json:"feePrec,string"`
	MinAmount float64 `json:"minAmount,string"`
}

func (c *client) GetMarkets() (*MarketsResp, error) {
	url := fmt.Sprintf("%s/public/markets", c.url)
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

	var result MarketsResp
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
