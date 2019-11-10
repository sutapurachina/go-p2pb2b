package p2pb2b

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Ticker struct {
	Result
	Ticker TickerEntry `json:"result"`
}

type TickerEntry struct {
	Bid    float64 `json:"bid,string"`
	Ask    float64 `json:"ask,string"`
	Open   float64 `json:"open,string"`
	High   float64 `json:"high,string"`
	Low    float64 `json:"low,string"`
	Last   float64 `json:"last,string"`
	Volume float64 `json:"volume,string"`
	Deal   float64 `json:"deal,string"`
	Change float64 `json:"change,string"`
}

func (c *client) GetTicker(market string) (*Ticker, error) {
	url := fmt.Sprintf("%s/public/ticker?market=%s", c.url, market)
	resp, err := c.sendGet(url, nil)
	if err != nil {
		return nil, err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result Ticker
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
