package p2pb2b

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Tickers struct {
	Success bool              `json:"success"`
	Message string            `json:"message"`
	Result  map[string]Result `json:"result"`
}

type Result struct {
	At     int64  `json:"at"`
	Ticker Ticker `json:"ticker"`
}

type Ticker struct {
	Bid  float64 `json:"bid,string"`
	Ask  float64 `json:"ask,string"`
	Low  float64 `json:"low,string"`
	High float64 `json:"high,string"`
	Last float64 `json:"last,string"`
	Vol  float64 `json:"vol,string"`
}

const TICKERS_BASE_API = "https://api.p2pb2b.io/api/v1"

func (c *client) GetTickers() (*Tickers, error) {
	url := fmt.Sprintf("%s/public/tickers", TICKERS_BASE_API)
	Infof("calling tickers url %s", url)
	resp, err := c.sendGet(url, nil)
	if err != nil {
		return nil, err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(bodyBytes))

	var tickers Tickers
	err = json.Unmarshal(bodyBytes, &tickers)
	if err != nil {
		return nil, err
	}
	return &tickers, nil
}
