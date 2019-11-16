package p2pb2b

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

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
		Request: "doesnt",
		Nonce:   "matter",
	}
	_, err = client.PostCurrencyBalance(request)
	assert.True(t, err != nil)
}

func TestPostCurrencyBalance(t *testing.T) {
	pseudoAPIKey := uuid.NewV4()
	pseudoAPISecret := "4a894c5c-8a7e-4337-bb6b-9fde16e3dddd"

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "/account/balance", r.URL.String())
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

		assert.Equal(t, pseudoAPIKey.String(), r.Header.Get("X-TXC-APIKEY"))
		assert.Equal(t, "eyJjdXJyZW5jeSI6IiIsInJlcXVlc3QiOiJkb2VzbnQiLCJub25jZSI6Im1hdHRlciJ9", r.Header.Get("X-TXC-PAYLOAD"))
		assert.Equal(t, "0d2b7d975581a595add02931b1de04cc99c09b4e2b9efba850071442e3275717", r.Header.Get("X-TXC-SIGNATURE"))

		w.WriteHeader(http.StatusOK)

		balanceMap := map[string]AccountCurrencyBalance{}
		balanceMap["blubb"] = AccountCurrencyBalance{
			Available: 5.0,
			Freeze:    1.0,
		}
		result := &AccountCurrencyBalanceResult{
			CurrencyBalances: balanceMap,
		}
		result.Success = true
		result.Message = ""
		asJSON, err := json.Marshal(result)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		w.Write(asJSON)
	}))
	defer ts.Close()

	client, err := newClientWithURL(ts.URL, pseudoAPIKey.String(), pseudoAPISecret)
	if err != nil {
		t.Error(err.Error())
	}
	request := &AccountCurrencyBalanceRequest{
		Request: "doesnt",
		Nonce:   "matter",
	}
	resp, err := client.PostCurrencyBalance(request)
	assert.Equal(t, true, resp.Success)
}
