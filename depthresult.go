package p2pb2b

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type DepthResultResp struct {
	Response
	Result      DepthResultResult `json:"result"`
	CacheTime   float64           `json:"cache_time"`
	CurrentTime float64           `json:"current_time"`
}

type DepthResultResult struct {
	Asks []Float64Pair `json:"asks"`
	Bids []Float64Pair `json:"bids"`
}

type Float64Pair [2]float64

func (c *client) GetDepthResult(market string, limit int64) (*DepthResultResp, error) {
	if market == "" {
		return nil, fmt.Errorf("parameter market must not be empty")
	}
	if limit <= 0 {
		return nil, fmt.Errorf("parameter limit must not be <= 0")
	}

	url := fmt.Sprintf("%s/public/depth/result?market=%s&limit=%d", c.url, market, limit)

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

	var result DepthResultResp
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (pair *Float64Pair) UnmarshalJSON(b []byte) error {
	tmp := []json.Number{}
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}

	left, err := tmp[0].Float64()
	if err != nil {
		return err
	}
	right, err := tmp[1].Float64()
	if err != nil {
		return err
	}
	*pair = Float64Pair{left, right}

	return nil
}
