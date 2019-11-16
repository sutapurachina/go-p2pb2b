package p2pb2b

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateOrder(t *testing.T) {
	pseudoAPIKey := uuid.NewV4()
	pseudoAPISecret := "4a894c5c-8a7e-4337-bb6b-9fde16e3dddd"
	body := `{
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

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "/order/new", r.URL.String())
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

		assert.Equal(t, pseudoAPIKey.String(), r.Header.Get("X-TXC-APIKEY"))
		assert.Equal(t, "eyJyZXF1ZXN0IjoiZG9lc250Iiwibm9uY2UiOiJtYXR0ZXIiLCJtYXJrZXQiOiJFVEhfQlRDIiwic2lkZSI6ImJ1eSIsImFtb3VudCI6IjAuMDAxIiwicHJpY2UiOiIxMDAwIn0=", r.Header.Get("X-TXC-PAYLOAD"))
		assert.Equal(t, "43af45d5df5e81bd4e8e562b35c5320193d989902babb00452419c6775e31cc0", r.Header.Get("X-TXC-SIGNATURE"))

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(body))
	}))
	defer ts.Close()

	client, err := newClientWithURL(ts.URL, pseudoAPIKey.String(), pseudoAPISecret)
	if err != nil {
		t.Error(err.Error())
	}
	request := &CreateOrderRequest{
		Request: Request{
			Request: "doesnt",
			Nonce:   "matter",
		},
		Market: "ETH_BTC",
		Side:   "buy",
		Amount: 0.001,
		Price:  1000.0,
	}
	resp, err := client.CreateOrder(request)

	assert.NotNil(t, resp, fmt.Sprintf("error: %v", err))
	assert.Equal(t, true, resp.Success)

	respBytes, _ := json.Marshal(resp)
	equal, _ := IsEqualJSON(body, string(respBytes))
	assert.True(t, equal, fmt.Sprintf("%v is not equal to %v", body, resp))
}

func TestCancelOrder(t *testing.T) {
	pseudoAPIKey := uuid.NewV4()
	pseudoAPISecret := "4a894c5c-8a7e-4337-bb6b-9fde16e3dddd"
	body := `{
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

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "/order/cancel", r.URL.String())
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

		assert.Equal(t, pseudoAPIKey.String(), r.Header.Get("X-TXC-APIKEY"))
		assert.Equal(t, "eyJyZXF1ZXN0IjoiZG9lc250Iiwibm9uY2UiOiJtYXR0ZXIiLCJtYXJrZXQiOiJFVEhfQlRDIiwib3JkZXJJZCI6MTIzfQ==", r.Header.Get("X-TXC-PAYLOAD"))
		assert.Equal(t, "ada019a050f66d2311029191870b5c0d91b720f40836f09ad8414f3ce15ed232", r.Header.Get("X-TXC-SIGNATURE"))

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(body))
	}))
	defer ts.Close()

	client, err := newClientWithURL(ts.URL, pseudoAPIKey.String(), pseudoAPISecret)
	if err != nil {
		t.Error(err.Error())
	}
	request := &CancelOrderRequest{
		Request: Request{
			Request: "doesnt",
			Nonce:   "matter",
		},
		Market:  "ETH_BTC",
		OrderID: 123,
	}
	resp, err := client.CancelOrder(request)

	assert.NotNil(t, resp, fmt.Sprintf("error: %v", err))
	assert.Equal(t, true, resp.Success)

	respBytes, _ := json.Marshal(resp)
	equal, _ := IsEqualJSON(body, string(respBytes))
	assert.True(t, equal, fmt.Sprintf("%v is not equal to %v", body, resp))
}

func TestQueryUnexecuted(t *testing.T) {
	pseudoAPIKey := uuid.NewV4()
	pseudoAPISecret := "4a894c5c-8a7e-4337-bb6b-9fde16e3dddd"

	body := `{
		"success": true,
		"message": "",
		"result": {
		  "limit": 100,
		  "offset": 0,
		  "total": 1,
		  "result": [
			{
			  "id": 3900714,
			  "left": "1",
			  "market": "ETH_BTC",
			  "amount": "1",
			  "type": "limit",
			  "price": "0.008",
			  "timestamp": 1546459568.376407,
			  "side": "buy",
			  "dealFee": "0",
			  "takerFee": "0.001",
			  "makerFee": "0.001",
			  "dealStock": "0",
			  "dealMoney": "0"
			}
		  ]
		}
	  }`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "/orders", r.URL.String())
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

		assert.Equal(t, pseudoAPIKey.String(), r.Header.Get("X-TXC-APIKEY"))
		assert.Equal(t, "eyJyZXF1ZXN0IjoiZG9lc250Iiwibm9uY2UiOiJtYXR0ZXIiLCJtYXJrZXQiOiJFVEhfQlRDIiwib2Zmc2V0IjowLCJsaW1pdCI6MTAwfQ==", r.Header.Get("X-TXC-PAYLOAD"))
		assert.Equal(t, "1b56e05299d8809d5f937f58483f70ca8951a9c0b6eb81498212f1236362b5bc", r.Header.Get("X-TXC-SIGNATURE"))

		w.WriteHeader(http.StatusOK)

		w.Write([]byte(body))
	}))
	defer ts.Close()

	client, err := newClientWithURL(ts.URL, pseudoAPIKey.String(), pseudoAPISecret)
	if err != nil {
		t.Error(err.Error())
	}
	request := &QueryUnexecutedRequest{
		Request: Request{
			Request: "doesnt",
			Nonce:   "matter",
		},
		Market: "ETH_BTC",
		Offset: 0,
		Limit:  100,
	}
	resp, err := client.QueryUnexecuted(request)

	assert.NotNil(t, resp, fmt.Sprintf("error: %v", err))
	assert.Equal(t, true, resp.Success)

	respBytes, _ := json.Marshal(resp)
	equal, _ := IsEqualJSON(body, string(respBytes))
	assert.True(t, equal, fmt.Sprintf("%v is not equal to %v", body, resp))
}

func TestQueryExecuted(t *testing.T) {
	pseudoAPIKey := uuid.NewV4()
	pseudoAPISecret := "4a894c5c-8a7e-4337-bb6b-9fde16e3dddd"

	body := `{
		"success": true,
		"message": "",
		"result": {
		  "ETH_BTC": [
			{
			  "amount": "1",
			  "price": "0.01",
			  "type": "limit",
			  "id": 9740,
			  "side": "sell",
			  "ctime": 1533568890.583023,
			  "takerFee": "0.002",
			  "ftime": 1533630652.62185,
			  "market": "ETH_BTC",
			  "makerFee": "0.002",
			  "dealFee": "0.002",
			  "dealStock": "1",
			  "dealMoney": "0.01",
			  "marketName": "ETH_BTC"
			}
		  ],
		  "ATB_USD": [
			{
			  "amount": "0.3",
			  "price": "0.06296168",
			  "type": "market",
			  "id": 11669,
			  "source": "",
			  "side": "buy",
			  "ctime": 1533626329.696647,
			  "takerFee": "0.002",
			  "ftime": 1533626329.696659,
			  "market": "ATB_USD",
			  "makerFee": "0.002",
			  "dealFee": "0.000037777008",
			  "dealStock": "0.3",
			  "dealMoney": "0.018888504",
			  "marketName": "ATB_USD"
			}
		  ]
		}
	  }`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "/account/order_history", r.URL.String())
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

		assert.Equal(t, pseudoAPIKey.String(), r.Header.Get("X-TXC-APIKEY"))
		assert.Equal(t, "eyJyZXF1ZXN0IjoiZG9lc250Iiwibm9uY2UiOiJtYXR0ZXIiLCJvZmZzZXQiOjAsImxpbWl0IjoxMDB9", r.Header.Get("X-TXC-PAYLOAD"))
		assert.Equal(t, "78d32292f175514604561c9e3c6dd2b4f6e95eba8615a144b0589adc1bb22577", r.Header.Get("X-TXC-SIGNATURE"))

		w.WriteHeader(http.StatusOK)

		w.Write([]byte(body))
	}))
	defer ts.Close()

	client, err := newClientWithURL(ts.URL, pseudoAPIKey.String(), pseudoAPISecret)
	if err != nil {
		t.Error(err.Error())
	}
	request := &QueryExecutedRequest{
		Request: Request{
			Request: "doesnt",
			Nonce:   "matter",
		},
		Offset: 0,
		Limit:  100,
	}
	resp, err := client.QueryExecuted(request)

	assert.NotNil(t, resp, fmt.Sprintf("error: %v", err))
	assert.Equal(t, true, resp.Success)

	respBytes, _ := json.Marshal(resp)
	equal, _ := IsEqualJSON(body, string(respBytes))
	assert.True(t, equal, fmt.Sprintf("%v is not equal to %v", body, resp))
}

func TestQueryDeals(t *testing.T) {
	pseudoAPIKey := uuid.NewV4()
	pseudoAPISecret := "4a894c5c-8a7e-4337-bb6b-9fde16e3dddd"
	body := `{
		"success": true,
		"message": "",
		"result": {
		  "offset": 0,
		  "limit": 50,
		  "records": [
			{
			  "time": 1533310924.935978,
			  "fee": "0",
			  "price": "80.22761599",
			  "amount": "2.12687945",
			  "id": 548,
			  "dealOrderId": 1237,
			  "role": 1,
			  "deal": "170.6344677716224055"
			}
		  ]
		}
	  }`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "/account/order", r.URL.String())
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

		assert.Equal(t, pseudoAPIKey.String(), r.Header.Get("X-TXC-APIKEY"))
		assert.Equal(t, "eyJyZXF1ZXN0IjoiZG9lc250Iiwibm9uY2UiOiJtYXR0ZXIiLCJvcmRlcklkIjoxMjMsIm9mZnNldCI6MCwibGltaXQiOjEwMH0=", r.Header.Get("X-TXC-PAYLOAD"))
		assert.Equal(t, "7db3e70852db48e776d35614e80c0ad205e2908ee839367ccad3a5b98c8e6b58", r.Header.Get("X-TXC-SIGNATURE"))

		w.WriteHeader(http.StatusOK)

		w.Write([]byte(body))
	}))
	defer ts.Close()

	client, err := newClientWithURL(ts.URL, pseudoAPIKey.String(), pseudoAPISecret)
	if err != nil {
		t.Error(err.Error())
	}
	request := &QueryDealsRequest{
		Request: Request{
			Request: "doesnt",
			Nonce:   "matter",
		},
		OrderID: 123,
		Offset:  0,
		Limit:   100,
	}
	resp, err := client.QueryDeals(request)

	assert.NotNil(t, resp, fmt.Sprintf("error: %v", err))
	assert.Equal(t, true, resp.Success)

	respBytes, _ := json.Marshal(resp)
	equal, _ := IsEqualJSON(body, string(respBytes))
	assert.True(t, equal, fmt.Sprintf("%v is not equal to %v", body, resp))
}
