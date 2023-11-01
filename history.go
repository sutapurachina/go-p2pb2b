package p2pb2b

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
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

type DealHistoryResp struct {
	Response
	ErrorCode string         `json:"errorCode"`
	Result    DealHistoryRes `json:"result"`
}

type DealHistoryEntry struct {
	DealId          int64   `json:"deal_id"`
	DealTime        float64 `json:"deal_time"`
	DealOrderId     int64   `json:"deal_order_id"`
	OppositeOrderId int64   `json:"opposite_order_id"`
	Side            string  `json:"side"`
	Price           string  `json:"price"`
	Amount          string  `json:"amount"`
	Deal            string  `json:"deal"`
	DealFee         string  `json:"deal_fee"`
	Role            string  `json:"role"`
	IsSelfTrade     bool    `json:"isSelfTrade"`
}

type DealHistoryRes struct {
	Total int                `json:"total"`
	Deals []DealHistoryEntry `json:"deals"`
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
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = checkHTTPStatus(*resp, http.StatusOK)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s: %s\n", err.Error(), string(bodyBytes)))
	}

	var result HistoryResp
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *client) DealsHistoryByMarket(request *DealsHistoryByMarketRequest) (*DealHistoryResp, error) {

	url := fmt.Sprintf("%s/account/market_deal_history", c.url)
	request.Request.Nonce = strconv.FormatInt(time.Now().UnixMilli(), 10)
	request.Request.Request = "/api/v2/account/market_deal_history"
	asJSON, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	resp, err := c.sendPost(url, nil, bytes.NewReader(asJSON))
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

	var result DealHistoryResp
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
