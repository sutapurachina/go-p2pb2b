package p2pb2b

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetTicker(t *testing.T) {
	pseudoAPIKey := uuid.NewV4()
	pseudoAPISecret := "4a894c5c-8a7e-4337-bb6b-9fde16e3dddd"
	body := `{
		"success": true,
		"message": "",
		"result": {
			"bid": "0.021475",
			"ask": "0.0215",
			"open": "0.021763",
			"high": "0.021855",
			"low": "0.021422",
			"last": "0.021489",
			"volume": "160302.558",
			"deal": "3469.740614926",
			"change": "-0.95"
		},
		"cache_time": 1574197469.668056,
		"current_time": 1574197469.668141
	}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "/public/ticker?market=ETH_BTC", r.URL.String())
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

	resp, err := client.GetTicker("ETH_BTC")
	if err != nil {
		t.Error(err.Error())
	}

	assert.NotNil(t, resp, fmt.Sprintf("error: %v", err))
	assert.Equal(t, true, resp.Success)
	assert.Equal(t, 1574197469.668056, resp.CacheTime)
	assert.Equal(t, 1574197469.668141, resp.CurrentTime)

	assert.Equal(t, 0.021475, resp.Result.Bid)
	assert.Equal(t, 0.0215, resp.Result.Ask)
	assert.Equal(t, 0.021763, resp.Result.Open)
	assert.Equal(t, 0.021855, resp.Result.High)
	assert.Equal(t, 0.021422, resp.Result.Low)
	assert.Equal(t, 0.021489, resp.Result.Last)
	assert.Equal(t, 160302.558, resp.Result.Volume)
	assert.Equal(t, 3469.740614926, resp.Result.Deal)
	assert.Equal(t, -0.95, resp.Result.Change)

}
