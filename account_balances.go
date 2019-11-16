package p2pb2b

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type AccountBalancesResult struct {
	Result
	Balances map[string]AccountBalance `json:"result"`
}

type AccountBalance struct {
	Available float64 `json:"available"`
	Freeze    float64 `json:"freeze"`
}

type AccountBalancesRequest struct {
	Request string `json:"request"`
	Nonce   string `json:"nonce"`
}

func (c *client) PostBalances(request *AccountBalancesRequest) (*AccountBalancesResult, error) {
	url := fmt.Sprintf("%s/account/balances", c.url)
	asJSON, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	resp, err := c.sendPost(url, nil, bytes.NewReader(asJSON))
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

	if http.StatusOK != resp.StatusCode {
		return nil, fmt.Errorf("status: %s, body: %s", resp.Status, bodyBytes)
	}

	var result AccountBalancesResult
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}