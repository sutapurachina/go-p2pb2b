package p2pb2b

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
		assert.Equal(t, "eyJyZXF1ZXN0Ijoie3tyZXF1ZXN0fX0iLCJub25jZSI6Int7bm9uY2V9fSIsIm1hcmtldCI6IkVUSF9CVEMiLCJzaWRlIjoiYnV5IiwiYW1vdW50IjoiMC4wMDEiLCJwcmljZSI6IjEwMDAifQ==", r.Header.Get("X-TXC-PAYLOAD"))
		assert.Equal(t, "b7cea8337772ecd0aa0af981037ea919b40645e8015becfdce36a3abdff6b440", r.Header.Get("X-TXC-SIGNATURE"))

		expectedReqBody := `{
			"market": "ETH_BTC",
			"side": "buy",
			"amount": "0.001",
			"price": "1000",
			"request": "{{request}}",
			"nonce": "{{nonce}}"
		}`
		reqBody, _ := ioutil.ReadAll(r.Body)
		equal, err := IsEqualJSON(expectedReqBody, string(reqBody))
		assert.Nil(t, err, err)
		assert.True(t, equal, fmt.Sprintf("%s is not equal to %s", expectedReqBody, string(reqBody)))

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
			Request: "{{request}}",
			Nonce:   "{{nonce}}",
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
	assert.True(t, equal, fmt.Sprintf("%s is not equal to %s", body, string(respBytes)))
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
		assert.Equal(t, "eyJyZXF1ZXN0Ijoie3tyZXF1ZXN0fX0iLCJub25jZSI6Int7bm9uY2V9fSIsIm1hcmtldCI6IkVUSF9CVEMiLCJvcmRlcklkIjoyNTc0OX0=", r.Header.Get("X-TXC-PAYLOAD"))
		assert.Equal(t, "f8883be6c0bacaae8c3e3aa4d3fc914cb2ee9303ec46648df2a265cc317e1cdd", r.Header.Get("X-TXC-SIGNATURE"))

		expectedReqBody := `{
			"market": "ETH_BTC",
			"orderId": 25749,
			"request": "{{request}}",
			"nonce": "{{nonce}}"
		}`
		reqBody, _ := ioutil.ReadAll(r.Body)
		equal, err := IsEqualJSON(expectedReqBody, string(reqBody))
		assert.Nil(t, err, err)
		assert.True(t, equal, fmt.Sprintf("%s is not equal to %s", expectedReqBody, string(reqBody)))

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
			Request: "{{request}}",
			Nonce:   "{{nonce}}",
		},
		Market:  "ETH_BTC",
		OrderID: 25749,
	}
	resp, err := client.CancelOrder(request)

	assert.NotNil(t, resp, fmt.Sprintf("error: %v", err))
	assert.Equal(t, true, resp.Success)

	respBytes, _ := json.Marshal(resp)
	equal, _ := IsEqualJSON(body, string(respBytes))
	assert.True(t, equal, fmt.Sprintf("%s is not equal to %s", body, string(respBytes)))
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
		assert.Equal(t, "eyJyZXF1ZXN0Ijoie3tyZXF1ZXN0fX0iLCJub25jZSI6Int7bm9uY2V9fSIsIm1hcmtldCI6IkVUSF9CVEMiLCJvZmZzZXQiOjAsImxpbWl0IjoxMDB9", r.Header.Get("X-TXC-PAYLOAD"))
		assert.Equal(t, "407a831bebf538cd01c179dd88117b9e5753986ba4e1f820afee16317d37a814", r.Header.Get("X-TXC-SIGNATURE"))

		expectedReqBody := `{
			"market": "ETH_BTC",
			"offset": 0,
			"limit": 100,
			"request": "{{request}}",
			"nonce": "{{nonce}}"
		}`
		reqBody, _ := ioutil.ReadAll(r.Body)
		equal, err := IsEqualJSON(expectedReqBody, string(reqBody))
		assert.Nil(t, err, err)
		assert.True(t, equal, fmt.Sprintf("%s is not equal to %s", expectedReqBody, string(reqBody)))

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
			Request: "{{request}}",
			Nonce:   "{{nonce}}",
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
	assert.True(t, equal, fmt.Sprintf("%s is not equal to %s", body, string(respBytes)))
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
		assert.Equal(t, "eyJyZXF1ZXN0Ijoie3tyZXF1ZXN0fX0iLCJub25jZSI6Int7bm9uY2V9fSIsIm9mZnNldCI6MCwibGltaXQiOjEwMH0=", r.Header.Get("X-TXC-PAYLOAD"))
		assert.Equal(t, "260d4db5d0cd1f8139b90f3f14b624c050eb6496837681e7c41f2638f8d565dc", r.Header.Get("X-TXC-SIGNATURE"))

		expectedReqBody := `{
			"offset": 0,
			"limit": 100,
			"request": "{{request}}",
			"nonce": "{{nonce}}"
		}`
		reqBody, _ := ioutil.ReadAll(r.Body)
		equal, err := IsEqualJSON(expectedReqBody, string(reqBody))
		assert.Nil(t, err, err)
		assert.True(t, equal, fmt.Sprintf("%s is not equal to %s", expectedReqBody, string(reqBody)))

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
			Request: "{{request}}",
			Nonce:   "{{nonce}}",
		},
		Offset: 0,
		Limit:  100,
	}
	resp, err := client.QueryExecuted(request)

	assert.NotNil(t, resp, fmt.Sprintf("error: %v", err))
	assert.Equal(t, true, resp.Success)

	respBytes, _ := json.Marshal(resp)
	equal, _ := IsEqualJSON(body, string(respBytes))
	assert.True(t, equal, fmt.Sprintf("%s is not equal to %s", body, string(respBytes)))
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
			  "deal": "170.6344677716224"
			}
		  ]
		}
	  }`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "/account/order", r.URL.String())
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

		assert.Equal(t, pseudoAPIKey.String(), r.Header.Get("X-TXC-APIKEY"))
		assert.Equal(t, "eyJyZXF1ZXN0Ijoie3tyZXF1ZXN0fX0iLCJub25jZSI6Int7bm9uY2V9fSIsIm9yZGVySWQiOjEyMzQsIm9mZnNldCI6MTAsImxpbWl0IjoxMDB9", r.Header.Get("X-TXC-PAYLOAD"))
		assert.Equal(t, "5fcc71442011d25e63749aff14995b8a539d9353c2ea12ae49e1fc77de49bb55", r.Header.Get("X-TXC-SIGNATURE"))

		expectedReqBody := `{
			"orderId": 1234,
			"offset": 10,
			"limit": 100,
			"request": "{{request}}",
			"nonce": "{{nonce}}"
		}`
		reqBody, _ := ioutil.ReadAll(r.Body)
		equal, err := IsEqualJSON(expectedReqBody, string(reqBody))
		assert.Nil(t, err, err)
		assert.True(t, equal, fmt.Sprintf("%s is not equal to %s", expectedReqBody, string(reqBody)))

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
			Request: "{{request}}",
			Nonce:   "{{nonce}}",
		},
		OrderID: 1234,
		Offset:  10,
		Limit:   100,
	}
	resp, err := client.QueryDeals(request)

	assert.NotNil(t, resp, fmt.Sprintf("error: %v", err))
	assert.Equal(t, true, resp.Success)

	respBytes, _ := json.Marshal(resp)
	equal, _ := IsEqualJSON(body, string(respBytes))
	assert.True(t, equal, fmt.Sprintf("%s is not equal to %s", body, string(respBytes)))
}
