package p2pb2b

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type HistoryResp struct {
	Response
	Result      []HistoryEntry `json:"result"`
	CacheTime   float64        `json:"cache_time"`
	CurrentTime float64        `json:"current_time"`
}

type HistoryEntry struct {
	ID     int     `json:"id"`
	Type   string  `json:"type"`
	Time   float64 `json:"time"`
	Amount float64 `json:"amount,string"`
	Price  float64 `json:"price,string"`
}

func (c *client) GetHistory(market string, lastID int64, limit int64) (*HistoryResp, error) {
	if market == "" {
		return nil, fmt.Errorf("parameter market must not be empty")
	}
	if lastID < 0 {
		return nil, fmt.Errorf("parameter offset must not be < 0")
	}
	if limit <= 0 {
		return nil, fmt.Errorf("parameter limit must not be <= 0")
	}

	url := fmt.Sprintf("%s/public/history?market=%s&lastId=%d&limit=%d", c.url, market, lastID, limit)

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

	var result HistoryResp
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
