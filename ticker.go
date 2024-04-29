package p2pb2b

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type TickerResp struct {
	Response
	Result      Ticker  `json:"result"`
	CacheTime   float64 `json:"cache_time"`
	CurrentTime float64 `json:"current_time"`
}

type Ticker struct {
	Bid    float64 `json:"bid,string"`
	Ask    float64 `json:"ask,string"`
	Open   float64 `json:"open,string"`
	High   float64 `json:"high,string"`
	Low    float64 `json:"low,string"`
	Last   float64 `json:"last,string"`
	Volume float64 `json:"volume,string"`
	Deal   float64 `json:"deal,string"`
	Change float64 `json:"change,string"`
}

type KlineResponse struct {
	Response
	Result      [][]interface{} `json:"result"`
	CacheTime   float64         `json:"cache_time"`
	CurrentTime float64         `json:"current_time"`
}

type Kline struct {
	KlineOpenTime          int64
	OpenPrice              float64
	ClosePrice             float64
	HighestPrice           float64
	LowestPrice            float64
	VolumeForStockCurrency float64
	VolumeForMoney         float64
	MarketName             string
}

func (c *client) GetTicker(market string) (*TickerResp, error) {
	url := fmt.Sprintf("%s/public/ticker?market=%s", c.url, market)
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

	var result TickerResp
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *client) Klines(market, interval string, limit, offset int) ([]*Kline, error) {
	url := fmt.Sprintf("%s/public/market/kline?market=%s&interval=%s&limit=%d&offset=%d", c.url, market, interval, limit, offset)
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

	var result KlineResponse
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return nil, err
	}
	res := make([]*Kline, 0, 1)
	for _, klineArr := range result.Result {
		kline := Kline{}
		for idx, par := range klineArr {
			switch idx {
			case 0:
				kline.KlineOpenTime = int64(par.(float64))
			case 1:
				kline.OpenPrice, err = strconv.ParseFloat(par.(string), 64)
				if err != nil {
					return nil, err
				}
			case 2:
				kline.ClosePrice, err = strconv.ParseFloat(par.(string), 64)
				if err != nil {
					return nil, err
				}
			case 3:
				kline.HighestPrice, err = strconv.ParseFloat(par.(string), 64)
				if err != nil {
					return nil, err
				}
			case 4:
				kline.LowestPrice, err = strconv.ParseFloat(par.(string), 64)
				if err != nil {
					return nil, err
				}
			case 5:
				kline.VolumeForStockCurrency, err = strconv.ParseFloat(par.(string), 64)
				if err != nil {
					return nil, err
				}
			case 6:
				kline.VolumeForMoney, err = strconv.ParseFloat(par.(string), 64)
				if err != nil {
					return nil, err
				}
			case 7:
				kline.MarketName = par.(string)

			}
			res = append(res, &kline)
		}
	}
	return res, nil

}
