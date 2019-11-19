package p2pb2b

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetTickers(t *testing.T) {
	pseudoAPIKey := uuid.NewV4()
	pseudoAPISecret := "4a894c5c-8a7e-4337-bb6b-9fde16e3dddd"
	body := `{
		"success": true,
		"message": "",
		"result": {
			"ETH_BTC": {
				"at": 1574197772,
				"ticker": {
					"bid": "0.021477",
					"ask": "0.021491",
					"low": "0.021422",
					"high": "0.021855",
					"last": "0.021477",
					"vol": "159706.449",
					"change": "-1.05"
				}
			},
			"BTC_USD": {
				"at": 1574197772,
				"ticker": {
					"bid": "8067.06",
					"ask": "8149.88",
					"low": "8003.5",
					"high": "8348.5",
					"last": "8107.72",
					"vol": "34976.482947",
					"change": "-1.28"
				}
			}
		},
		"cache_time": 1574197772.016465,
		"current_time": 1574197772.018502
	}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "/public/tickers", r.URL.String())
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

	resp, err := client.GetTickers()
	if err != nil {
		t.Error(err.Error())
	}

	assert.NotNil(t, resp, fmt.Sprintf("error: %v", err))
	assert.Equal(t, true, resp.Success)
	assert.Equal(t, 1574197772.016465, resp.CacheTime)
	assert.Equal(t, 1574197772.018502, resp.CurrentTime)

	assert.NotEmpty(t, resp.Result["ETH_BTC"])
	assert.Equal(t, 1574197772, resp.Result["ETH_BTC"].At)
	assert.Equal(t, 0.021477, resp.Result["ETH_BTC"].Ticker.Bid)
	assert.Equal(t, 0.021491, resp.Result["ETH_BTC"].Ticker.Ask)
	assert.Equal(t, 0.021422, resp.Result["ETH_BTC"].Ticker.Low)
	assert.Equal(t, 0.021855, resp.Result["ETH_BTC"].Ticker.High)
	assert.Equal(t, 0.021477, resp.Result["ETH_BTC"].Ticker.Last)
	assert.Equal(t, 159706.449, resp.Result["ETH_BTC"].Ticker.Volume)
	assert.Equal(t, -1.05, resp.Result["ETH_BTC"].Ticker.Change)

	assert.NotEmpty(t, resp.Result["BTC_USD"])
	assert.Equal(t, 1574197772, resp.Result["BTC_USD"].At)
	assert.Equal(t, 8067.06, resp.Result["BTC_USD"].Ticker.Bid)
	assert.Equal(t, 8149.88, resp.Result["BTC_USD"].Ticker.Ask)
	assert.Equal(t, 8003.5, resp.Result["BTC_USD"].Ticker.Low)
	assert.Equal(t, 8348.5, resp.Result["BTC_USD"].Ticker.High)
	assert.Equal(t, 8107.72, resp.Result["BTC_USD"].Ticker.Last)
	assert.Equal(t, 34976.482947, resp.Result["BTC_USD"].Ticker.Volume)
	assert.Equal(t, -1.28, resp.Result["BTC_USD"].Ticker.Change)
}
