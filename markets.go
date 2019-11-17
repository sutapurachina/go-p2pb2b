package p2pb2b

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type MarketsResp struct {
	Response
	Markets []Market `json:"result"`
}

type Market struct {
	Name      string  `json:"name"`
	MoneyPrec int64   `json:"moneyPrec,string"`
	Stock     string  `json:"stock"`
	Money     string  `json:"money"`
	StockPrec int64   `json:"stockPrec,string"`
	FeePrec   int64   `json:"feePrec,string"`
	MinAmount float64 `json:"minAmount,string"`
}

func (c *client) GetMarkets() (*MarketsResp, error) {
	url := fmt.Sprintf("%s/public/markets", c.url)
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

	var result MarketsResp
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
