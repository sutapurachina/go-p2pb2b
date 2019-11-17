package p2pb2b

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type AccountBalancesResp struct {
	Response
	Balances map[string]AccountBalance `json:"result"`
}

type AccountBalance struct {
	Available float64 `json:"available,string"`
	Freeze    float64 `json:"freeze,string"`
}

type AccountBalancesRequest struct {
	Request
}

type AccountCurrencyBalanceResp struct {
	Response
	CurrencyBalances map[string]AccountCurrencyBalance `json:"result"`
}

type AccountCurrencyBalance struct {
	Available float64 `json:"available,string"`
	Freeze    float64 `json:"freeze,string"`
}

type AccountCurrencyBalanceRequest struct {
	Request
	Currency string `json:"currency"`
}

func (c *client) PostBalances(request *AccountBalancesRequest) (*AccountBalancesResp, error) {
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

	var result AccountBalancesResp
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *client) PostCurrencyBalance(request *AccountCurrencyBalanceRequest) (*AccountCurrencyBalanceResp, error) {
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

	var result AccountCurrencyBalanceResp
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
