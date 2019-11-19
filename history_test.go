package p2pb2b

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetHistory(t *testing.T) {
	pseudoAPIKey := uuid.NewV4()
	pseudoAPISecret := "4a894c5c-8a7e-4337-bb6b-9fde16e3dddd"
	body := `{
		"success": true,
		"message": "",
		"result": [
			{
				"id": 160354369,
				"type": "sell",
				"time": 1574195085.511277,
				"amount": "0.174",
				"price": "0.021427"
			},
			{
				"id": 160354368,
				"type": "sell",
				"time": 1574195085.511159,
				"amount": "0.501",
				"price": "0.021427"
			}
		],
		"cache_time": 1574195086.649298,
		"current_time": 1574195086.650654
	}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "/public/history?market=ETH_BTC&lastId=1&limit=2", r.URL.String())
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

	resp, err := client.GetHistory("ETH_BTC", 1, 2)
	if err != nil {
		t.Error(err.Error())
	}

	assert.NotNil(t, resp, fmt.Sprintf("error: %v", err))
	assert.Equal(t, true, resp.Success)
	assert.Equal(t, 1574195086.649298, resp.CacheTime)
	assert.Equal(t, 1574195086.650654, resp.CurrentTime)
	assert.Equal(t, "", resp.Message)

	assert.Equal(t, 2, len(resp.Result))
	assert.Equal(t, 160354369, resp.Result[0].ID)
	assert.Equal(t, "sell", resp.Result[0].Type)
	assert.Equal(t, 1574195085.511277, resp.Result[0].Time)
	assert.Equal(t, 0.174, resp.Result[0].Amount)
	assert.Equal(t, 0.021427, resp.Result[0].Price)
	assert.Equal(t, 160354368, resp.Result[1].ID)
	assert.Equal(t, "sell", resp.Result[1].Type)
	assert.Equal(t, 1574195085.511159, resp.Result[1].Time)
	assert.Equal(t, 0.501, resp.Result[1].Amount)
	assert.Equal(t, 0.021427, resp.Result[1].Price)
}

func TestGetHistoryNegative(t *testing.T) {
	client, err := NewClient(uuid.NewV4().String(), uuid.NewV4().String())
	if err != nil {
		t.Error(err.Error())
	}

	history, err := client.GetHistory("ETH_BTC", -1, 10)
	assert.Nil(t, history)
	assert.NotNil(t, err)

	history, err = client.GetHistory("ETH_BTC", 2, 0)
	assert.Nil(t, history)
	assert.NotNil(t, err)

	history, err = client.GetHistory("ETH_BTC", 2, -5)
	assert.Nil(t, history)
	assert.NotNil(t, err)

	history, err = client.GetHistory("blubb", 0, 10)
	assert.Nil(t, history)
	assert.NotNil(t, err)
}
