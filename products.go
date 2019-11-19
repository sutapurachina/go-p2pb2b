package p2pb2b

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ProductsResp struct {
	Response
	Result      []Product `json:"result"`
	CacheTime   float64   `json:"cache_time"`
	CurrentTime float64   `json:"current_time"`
}

type Product struct {
	ID         string `json:"id"`
	FromSymbol string `json:"fromSymbol"`
	ToSymbol   string `json:"toSymbol"`
}

func (c *client) GetProducts() (*ProductsResp, error) {
	url := fmt.Sprintf("%s/public/products", c.url)
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

	var result ProductsResp
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
