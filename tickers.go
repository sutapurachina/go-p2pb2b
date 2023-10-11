package p2pb2b

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type TickersResp struct {
	Response
	Result      map[string]TickersResult `json:"result"`
	CacheTime   float64                  `json:"cache_time"`
	CurrentTime float64                  `json:"current_time"`
}

type TickersResult struct {
	At     int          `json:"at"`
	Ticker TickersEntry `json:"ticker"`
}

type TickersEntry struct {
	Bid    float64 `json:"bid,string"`
	Ask    float64 `json:"ask,string"`
	Low    float64 `json:"low,string"`
	High   float64 `json:"high,string"`
	Last   float64 `json:"last,string"`
	Volume float64 `json:"vol,string"`
	Change float64 `json:"change,string"`
}

func (c *client) GetTickers() (*TickersResp, error) {
	url := fmt.Sprintf("%s/public/tickers", c.url)
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
	var result TickersResp
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
