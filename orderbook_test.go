package p2pb2b

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetOrderBook(t *testing.T) {
	pseudoAPIKey := uuid.NewV4()
	pseudoAPISecret := "4a894c5c-8a7e-4337-bb6b-9fde16e3dddd"
	body := `{
		"success": true,
		"message": "",
		"result": {
			"offset": 0,
			"limit": 2,
			"total": 126,
			"orders": [
				{
					"id": 1456837208,
					"left": "1.7",
					"market": "ETH_BTC",
					"amount": "1.7",
					"type": "limit",
					"price": "0.021443",
					"timestamp": 1574195955.326248,
					"side": "sell",
					"dealFee": "0.000080124",
					"takerFee": "0.0001",
					"makerFee": "0.0001",
					"dealStock": "1.821",
					"dealMoney": "0.040062"
				},
				{
					"id": 1456837207,
					"left": "2.986",
					"market": "ETH_BTC",
					"amount": "2.986",
					"type": "limit",
					"price": "0.021446",
					"timestamp": 1574195955.322718,
					"side": "sell",
					"dealFee": "0.000080124",
					"takerFee": "0.0001",
					"makerFee": "0.0001",
					"dealStock": "1.821",
					"dealMoney": "0.040062"
				}
			]
		},
		"cache_time": 1574195955.999738,
		"current_time": 1574195955.999887
	}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "/public/book?market=ETH_BTC&side=sell&offset=0&limit=2", r.URL.String())
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

		assert.Equal(t, pseudoAPIKey.String(), r.Header.Get("X-TXC-APIKEY"))
		assert.Empty(t, r.Header.Get("X-TXC-PAYLOAD"))
		assert.Empty(t, r.Header.Get("X-TXC-SIGNATURE"))

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(body))
	}))
	defer ts.Close()

	client, err := newClientWithURL(ts.URL, pseudoAPIKey.String(), pseudoAPISecret)
	if err != nil {
		t.Error(err.Error())
	}

	resp, err := client.GetOrderBook("ETH_BTC", "sell", 0, 2)
	if err != nil {
		t.Error(err.Error())
	}

	assert.NotNil(t, resp, fmt.Sprintf("error: %v", err))
	assert.Equal(t, true, resp.Success)
	assert.Equal(t, 1574195955.999738, resp.CacheTime)
	assert.Equal(t, 1574195955.999887, resp.CurrentTime)
	assert.Equal(t, "", resp.Message)
	assert.Equal(t, 0, resp.Result.Offset)
	assert.Equal(t, 2, resp.Result.Limit)
	assert.Equal(t, 126, resp.Result.Total)
	assert.Equal(t, 2, len(resp.Result.Orders))

	assert.Equal(t, 1456837208, resp.Result.Orders[0].ID)
	assert.Equal(t, 1.7, resp.Result.Orders[0].Left)
	assert.Equal(t, "ETH_BTC", resp.Result.Orders[0].Market)
	assert.Equal(t, 1.7, resp.Result.Orders[0].Amount)
	assert.Equal(t, "limit", resp.Result.Orders[0].Type)
	assert.Equal(t, 0.021443, resp.Result.Orders[0].Price)
	assert.Equal(t, 1574195955.326248, resp.Result.Orders[0].Timestamp)
	assert.Equal(t, "sell", resp.Result.Orders[0].Side)
	assert.Equal(t, 0.000080124, resp.Result.Orders[0].DealFee)
	assert.Equal(t, 0.0001, resp.Result.Orders[0].TakerFee)
	assert.Equal(t, 0.0001, resp.Result.Orders[0].MakerFee)
	assert.Equal(t, 1.821, resp.Result.Orders[0].DealStock)
	assert.Equal(t, 0.040062, resp.Result.Orders[0].DealMoney)

	assert.Equal(t, 1456837207, resp.Result.Orders[1].ID)
	assert.Equal(t, 2.986, resp.Result.Orders[1].Left)
	assert.Equal(t, "ETH_BTC", resp.Result.Orders[1].Market)
	assert.Equal(t, 2.986, resp.Result.Orders[1].Amount)
	assert.Equal(t, "limit", resp.Result.Orders[1].Type)
	assert.Equal(t, 0.021446, resp.Result.Orders[1].Price)
	assert.Equal(t, 1574195955.322718, resp.Result.Orders[1].Timestamp)
	assert.Equal(t, "sell", resp.Result.Orders[1].Side)
	assert.Equal(t, 0.000080124, resp.Result.Orders[1].DealFee)
	assert.Equal(t, 0.0001, resp.Result.Orders[1].TakerFee)
	assert.Equal(t, 0.0001, resp.Result.Orders[1].MakerFee)
	assert.Equal(t, 1.821, resp.Result.Orders[1].DealStock)
	assert.Equal(t, 0.040062, resp.Result.Orders[1].DealMoney)
}

func TestGetOrderBookNegative(t *testing.T) {
	client, err := NewClient(uuid.NewV4().String(), uuid.NewV4().String())
	if err != nil {
		t.Error(err.Error())
	}

	orderBook, err := client.GetOrderBook("ETH_BTC", "sell", -1, 10)
	assert.Nil(t, orderBook)
	assert.NotNil(t, err)

	orderBook, err = client.GetOrderBook("ETH_BTC", "sell", 2, 0)
	assert.Nil(t, orderBook)
	assert.NotNil(t, err)

	orderBook, err = client.GetOrderBook("ETH_BTC", "sell", 2, -5)
	assert.Nil(t, orderBook)
	assert.NotNil(t, err)

	orderBook, err = client.GetOrderBook("blubb", "sell", 0, 10)
	assert.Nil(t, orderBook)
	assert.NotNil(t, err)

	orderBook, err = client.GetOrderBook("ETH_BTC", "blubb", 0, 10)
	assert.Nil(t, orderBook)
	assert.NotNil(t, err)
}
