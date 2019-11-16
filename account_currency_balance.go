package p2pb2b

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type AccountCurrencyBalanceResult struct {
	Result
	CurrencyBalances map[string]AccountCurrencyBalance `json:"result"`
}

type AccountCurrencyBalance struct {
	Available float64 `json:"available,string"`
	Freeze    float64 `json:"freeze,string"`
}

type AccountCurrencyBalanceRequest struct {
	Currency string `json:"currency"`
	Request  string `json:"request"`
	Nonce    string `json:"nonce"`
}

func (c *client) PostCurrencyBalance(request *AccountCurrencyBalanceRequest) (*AccountCurrencyBalanceResult, error) {
	url := fmt.Sprintf("%s/account/balance", c.url)
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

	var result AccountCurrencyBalanceResult
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
