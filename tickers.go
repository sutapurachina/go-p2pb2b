package p2pb2b

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type TickersResp struct {
	Response
	Tickers map[string]TickersResult `json:"result"`
}

type TickersResult struct {
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

func (c *client) GetTickers() (*TickersResp, error) {
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

	var result TickersResp
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
