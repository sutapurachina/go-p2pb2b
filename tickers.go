package p2pb2b

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type TickersResult struct {
	Result
	Tickers map[string]TickersSnapshot `json:"result"`
}

type TickersSnapshot struct {
	At     int64        `json:"at"`
	Ticker TickersEntry `json:"ticker"`
}

type TickersEntry struct {
	Bid  float64 `json:"bid,string"`
	Ask  float64 `json:"ask,string"`
	Low  float64 `json:"low,string"`
	High float64 `json:"high,string"`
	Last float64 `json:"last,string"`
	Vol  float64 `json:"vol,string"`
}

func (c *client) GetTickers() (*TickersResult, error) {
	url := fmt.Sprintf("%s/public/tickers", c.url)
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

	var result TickersResult
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
