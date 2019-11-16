package p2pb2b

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type HistoryResult struct {
	Result
	HistoryEntries []HistoryEntry `json:"result"`
}

type HistoryEntry struct {
	ID     int64   `json:"id"`
	Type   string  `json:"type"`
	Time   float64 `json:"time"`
	Amount string  `json:"amount"`
	Price  string  `json:"price"`
}

func (c *client) GetHistory(market string, lastID int64, limit int64) (*HistoryResult, error) {
	if market == "" {
		return nil, fmt.Errorf("parameter market must not be empty")
	}
	if lastID < 0 {
		return nil, fmt.Errorf("parameter offset must not be < 0")
	}
	if limit <= 0 {
		return nil, fmt.Errorf("parameter limit must not be <= 0")
	}
	v := url.Values{}
	v.Set("market", market)
	v.Add("lastId", fmt.Sprintf("%d", lastID))
	v.Add("limit", fmt.Sprintf("%d", limit))

	url := fmt.Sprintf("%s/public/history?%s", c.url, v.Encode())

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

	var result HistoryResult
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
