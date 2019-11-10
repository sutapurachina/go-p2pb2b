package p2pb2b

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type SymbolsResult struct {
	Result
	Symbols []string `json:"result"`
}

func (c *client) GetSymbols() (*SymbolsResult, error) {
	url := fmt.Sprintf("%s/public/symbols", c.url)
	resp, err := c.sendGet(url, nil)
	if err != nil {
		return nil, err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result SymbolsResult
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
