package p2pb2b

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetMarkets(t *testing.T) {
	pseudoAPIKey := uuid.NewV4()
	pseudoAPISecret := "4a894c5c-8a7e-4337-bb6b-9fde16e3dddd"
	body := `{
		"success": true,
		"message": "",
		"result": [
			{
				"name": "ETH_BTC",
				"stock": "ETH",
				"money": "BTC",
				"moneyPrec": "6",
				"stockPrec": "3",
				"feePrec": "4",
				"minAmount": "0.001"
			},
			{
				"name": "BTC_USD",
				"stock": "BTC",
				"money": "USD",
				"moneyPrec": "2",
				"stockPrec": "6",
				"feePrec": "4",
				"minAmount": "0.001"
			}
		],
		"cache_time": 1574195556.499349,
		"current_time": 1574195556.501042
	}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "/public/markets", r.URL.String())
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

	resp, err := client.GetMarkets()
	if err != nil {
		t.Error(err.Error())
	}

	assert.NotNil(t, resp, fmt.Sprintf("error: %v", err))
	assert.Equal(t, true, resp.Success)
	assert.Equal(t, 1574195556.499349, resp.CacheTime)
	assert.Equal(t, 1574195556.501042, resp.CurrentTime)
	assert.Equal(t, "", resp.Message)

	assert.Equal(t, 2, len(resp.Result))
	assert.Equal(t, "ETH_BTC", resp.Result[0].Name)
	assert.Equal(t, "ETH", resp.Result[0].Stock)
	assert.Equal(t, "BTC", resp.Result[0].Money)
	assert.Equal(t, 6, resp.Result[0].MoneyPrec)
	assert.Equal(t, 3, resp.Result[0].StockPrec)
	assert.Equal(t, 4, resp.Result[0].FeePrec)
	assert.Equal(t, 0.001, resp.Result[0].MinAmount)

	assert.Equal(t, "BTC_USD", resp.Result[1].Name)
	assert.Equal(t, "BTC", resp.Result[1].Stock)
	assert.Equal(t, "USD", resp.Result[1].Money)
	assert.Equal(t, 2, resp.Result[1].MoneyPrec)
	assert.Equal(t, 6, resp.Result[1].StockPrec)
	assert.Equal(t, 4, resp.Result[1].FeePrec)
	assert.Equal(t, 0.001, resp.Result[1].MinAmount)
}
