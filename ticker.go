package p2pb2b

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type TickerResp struct {
	Response
	Result      Ticker  `json:"result"`
	CacheTime   float64 `json:"cache_time"`
	CurrentTime float64 `json:"current_time"`
}

type Ticker struct {
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

func (c *client) GetTicker(market string) (*TickerResp, error) {
	url := fmt.Sprintf("%s/public/ticker?market=%s", c.url, market)
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

	var result TickerResp
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
