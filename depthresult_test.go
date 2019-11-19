package p2pb2b

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetDepthResult(t *testing.T) {
	pseudoAPIKey := uuid.NewV4()
	pseudoAPISecret := "4a894c5c-8a7e-4337-bb6b-9fde16e3dddd"
	body := `{
		"success": true,
		"message": "",
		"result": {
			"asks": [
				[
					"0.021486",
					"0.654"
				],
				[
					"0.021488",
					"0.836"
				]
			],
			"bids": [
				[
					"0.021462",
					"0.665"
				],
				[
					"0.021455",
					"0.032"
				]
			]
		},
		"cache_time": 1574193389.575688,
		"current_time": 1574193389.5758
	}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "/public/depth/result?market=ETH_BTC&limit=2", r.URL.String())
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
	resp, err := client.GetDepthResult("ETH_BTC", 2)
	if err != nil {
		t.Error(err.Error())
	}
	assert.NotNil(t, resp, fmt.Sprintf("error: %v", err))
	assert.Equal(t, true, resp.Success)
	assert.Equal(t, 1574193389.575688, resp.CacheTime)
	assert.Equal(t, 1574193389.5758, resp.CurrentTime)
	assert.Equal(t, "", resp.Message)

	assert.Equal(t, 2, len(resp.Result.Asks))
	assert.Equal(t, 0.021486, resp.Result.Asks[0][0])
	assert.Equal(t, 0.654, resp.Result.Asks[0][1])
	assert.Equal(t, 0.021488, resp.Result.Asks[1][0])
	assert.Equal(t, 0.836, resp.Result.Asks[1][1])
	assert.Equal(t, 0.021462, resp.Result.Bids[0][0])
	assert.Equal(t, 0.665, resp.Result.Bids[0][1])
	assert.Equal(t, 0.021455, resp.Result.Bids[1][0])
	assert.Equal(t, 0.032, resp.Result.Bids[1][1])
}

func TestGetDepthResultNegative(t *testing.T) {
	client, err := NewClient(uuid.NewV4().String(), uuid.NewV4().String())
	if err != nil {
		t.Error(err.Error())
	}

	depthResult, err := client.GetDepthResult("ETH_BTC", 0)
	assert.Nil(t, depthResult)
	assert.NotNil(t, err)

	depthResult, err = client.GetDepthResult("blubb", 1)
	assert.Nil(t, depthResult)
	assert.NotNil(t, err)

	depthResult, err = client.GetDepthResult("ETH_BTC", -5)
	assert.Nil(t, depthResult)
	assert.NotNil(t, err)
}
