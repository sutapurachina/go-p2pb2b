package p2pb2b

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type ProductsResult struct {
	Result
	Products []Product `json:"result"`
}

type Product struct {
	ID         string `json:"id"`
	FromSymbol string `json:"fromSymbols"`
	ToSymbol   string `json:"toSymbol"`
}

func (c *client) GetProducts() (*ProductsResult, error) {
	url := fmt.Sprintf("%s/public/products", c.url)
	resp, err := c.sendGet(url, nil)
	if err != nil {
		return nil, err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result ProductsResult
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
