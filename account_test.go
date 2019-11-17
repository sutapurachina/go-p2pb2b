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

func TestPostBalancesNoKeyProvided(t *testing.T) {
	pseudoAPIKey := uuid.NewV4()
	pseudoAPISecret := "4a894c5c-8a7e-4337-bb6b-9fde16e3dddd"

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "POST")
		assert.Equal(t, r.URL.String(), "/account/balances")
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"success\":false,\"message\":[[\"Key not provided.\"]],\"result\":[]}"))
	}))
	defer ts.Close()

	client, err := newClientWithURL(ts.URL, pseudoAPIKey.String(), pseudoAPISecret)
	if err != nil {
		t.Error(err.Error())
	}
	request := &AccountBalancesRequest{
		Request: Request{
			Request: "doesnt",
			Nonce:   "matter",
		},
	}
	_, err = client.PostBalances(request)
	assert.True(t, err != nil)
}

func TestPostBalances(t *testing.T) {
	pseudoAPIKey := uuid.NewV4()
	pseudoAPISecret := "4a894c5c-8a7e-4337-bb6b-9fde16e3dddd"
	body := `{
		"success": true,
		"message": "",
		"result": {
		"ETH": {
			"available": "0.1",
			"freeze": "0.4"
		}
		}
	}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "/account/balances", r.URL.String())
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

		assert.Equal(t, pseudoAPIKey.String(), r.Header.Get("X-TXC-APIKEY"))
		assert.Equal(t, "eyJyZXF1ZXN0Ijoie3tyZXF1ZXN0fX0iLCJub25jZSI6Int7bm9uY2V9fSJ9", r.Header.Get("X-TXC-PAYLOAD"))
		assert.Equal(t, "fc545d1ebb38029d0d339b7545ee6252ae9ea2af9689b1a970775c3b74d9326c", r.Header.Get("X-TXC-SIGNATURE"))

		expectedReqBody := `{
			"request": "{{request}}",
			"nonce": "{{nonce}}"
		}`
		reqBody, _ := ioutil.ReadAll(r.Body)
		equal, err := IsEqualJSON(expectedReqBody, string(reqBody))
		assert.Nil(t, err, err)
		assert.True(t, equal, fmt.Sprintf("%s is not equal to %v", reqBody, expectedReqBody))

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(body))
	}))
	defer ts.Close()

	client, err := newClientWithURL(ts.URL, pseudoAPIKey.String(), pseudoAPISecret)
	if err != nil {
		t.Error(err.Error())
	}
	request := &AccountBalancesRequest{
		Request: Request{
			Request: "{{request}}",
			Nonce:   "{{nonce}}",
		},
	}
	resp, err := client.PostBalances(request)

	assert.NotNil(t, resp, fmt.Sprintf("error: %v", err))
	assert.Equal(t, true, resp.Success)

	respBytes, _ := json.Marshal(resp)
	equal, _ := IsEqualJSON(body, string(respBytes))
	assert.True(t, equal, fmt.Sprintf("%v is not equal to %v", body, resp))
}

func TestPostCurrencyBalanceNoKeyProvided(t *testing.T) {
	pseudoAPIKey := uuid.NewV4()
	pseudoAPISecret := "4a894c5c-8a7e-4337-bb6b-9fde16e3dddd"

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "POST")
		assert.Equal(t, r.URL.String(), "/account/balance")
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"success\":false,\"message\":[[\"Key not provided.\"]],\"result\":[]}"))
	}))
	defer ts.Close()

	client, err := newClientWithURL(ts.URL, pseudoAPIKey.String(), pseudoAPISecret)
	if err != nil {
		t.Error(err.Error())
	}
	request := &AccountCurrencyBalanceRequest{
		Request: Request{
			Request: "{{request}}",
			Nonce:   "{{nonce}}",
		},
		Currency: "ETH",
	}
	_, err = client.PostCurrencyBalance(request)
	assert.True(t, err != nil)
}

func TestPostCurrencyBalance(t *testing.T) {
	pseudoAPIKey := uuid.NewV4()
	pseudoAPISecret := "4a894c5c-8a7e-4337-bb6b-9fde16e3dddd"

	body := `{
		"success": true,
		"message": "",
		"result": {
		  "ETH": {
			"available": "0.63",
			"freeze": "0"
		  }
		}
	  }`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "/account/balance", r.URL.String())
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

		assert.Equal(t, pseudoAPIKey.String(), r.Header.Get("X-TXC-APIKEY"))
		assert.Equal(t, "eyJyZXF1ZXN0Ijoie3tyZXF1ZXN0fX0iLCJub25jZSI6Int7bm9uY2V9fSIsImN1cnJlbmN5IjoiRVRIIn0=", r.Header.Get("X-TXC-PAYLOAD"))
		assert.Equal(t, "018092f426a985c45d8b0c027ea361ddd71e84324bb35d2f33533fc1548bae08", r.Header.Get("X-TXC-SIGNATURE"))

		expectedReqBody := `{
			"currency": "ETH",
			"request": "{{request}}",
			"nonce": "{{nonce}}"
		}`
		reqBody, _ := ioutil.ReadAll(r.Body)
		equal, err := IsEqualJSON(expectedReqBody, string(reqBody))
		assert.Nil(t, err, err)
		assert.True(t, equal, fmt.Sprintf("%s is not equal to %v", reqBody, expectedReqBody))
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(body))
	}))
	defer ts.Close()

	client, err := newClientWithURL(ts.URL, pseudoAPIKey.String(), pseudoAPISecret)
	if err != nil {
		t.Error(err.Error())
	}
	request := &AccountCurrencyBalanceRequest{
		Request: Request{
			Request: "{{request}}",
			Nonce:   "{{nonce}}",
		},
		Currency: "ETH",
	}
	resp, err := client.PostCurrencyBalance(request)
	assert.NotNil(t, resp, fmt.Sprintf("erro: %v", err))
	assert.Equal(t, true, resp.Success)

	respBytes, _ := json.Marshal(resp)
	equal, _ := IsEqualJSON(body, string(respBytes))
	assert.True(t, equal, fmt.Sprintf("%v is not equal to %v", body, resp))
}
