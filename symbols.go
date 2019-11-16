package p2pb2b

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type SymbolsResp struct {
	Response
	Symbols []string `json:"result"`
}

func (c *client) GetSymbols() (*SymbolsResp, error) {
	url := fmt.Sprintf("%s/public/symbols", c.url)
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

	var result SymbolsResp
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
