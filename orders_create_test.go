package p2pb2b

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestPostCreateOrder(t *testing.T) {
	pseudoAPIKey := uuid.NewV4()
	pseudoAPISecret := "4a894c5c-8a7e-4337-bb6b-9fde16e3dddd"

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "/order/new", r.URL.String())
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

		assert.Equal(t, pseudoAPIKey.String(), r.Header.Get("X-TXC-APIKEY"))
		assert.Equal(t, "eyJtYXJrZXQiOiJFVEhfQlRDIiwic2lkZSI6ImJ1eSIsImFtb3VudCI6IjAuMDAxIiwicHJpY2UiOiIxMDAwIiwicmVxdWVzdCI6ImRvZXNudCIsIm5vbmNlIjoibWF0dGVyIn0=", r.Header.Get("X-TXC-PAYLOAD"))
		assert.Equal(t, "ff4b840c26f6247e5fe6de7dff1568002c108a915f09e6d96e9c10f7ed6076ec", r.Header.Get("X-TXC-SIGNATURE"))

		w.WriteHeader(http.StatusOK)

		resp := `{
			"success": true,
			"message": "",
			"result": {
				"orderId": 25749,
				"market": "ETH_BTC",
				"price": "0.1",
				"side": "sell",
				"type": "limit",
				"timestamp": 1537535284.828868,
				"dealMoney": "0",
				"dealStock": "0",
				"amount": "0.1",
				"takerFee": "0.002",
				"makerFee": "0.002",
				"left": "0.1",
				"dealFee": "0"
			}
		}`

		w.Write([]byte(resp))
	}))
	defer ts.Close()

	client, err := newClientWithURL(ts.URL, pseudoAPIKey.String(), pseudoAPISecret)
	if err != nil {
		t.Error(err.Error())
	}
	request := &OrdersCreateRequest{
		Market:  "ETH_BTC",
		Side:    "buy",
		Amount:  0.001,
		Price:   1000.0,
		Request: "doesnt",
		Nonce:   "matter",
	}
	resp, err := client.PostCreateOrder(request)

	assert.NotNil(t, resp, fmt.Sprintf("error: %v", err))
	assert.Equal(t, true, resp.Success)
	assert.Equal(t, 0.002, resp.Order.TakerFee)
}
