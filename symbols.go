package p2pb2b

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type SymbolsResp struct {
	Response
	Result      []string `json:"result"`
	CacheTime   float64  `json:"cache_time"`
	CurrentTime float64  `json:"current_time"`
}

func (c *client) GetSymbols() (*SymbolsResp, error) {
	url := fmt.Sprintf("%s/public/symbols", c.url)
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
	var result SymbolsResp
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
