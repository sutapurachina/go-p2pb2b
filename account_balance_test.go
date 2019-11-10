package p2pb2b

import (
	"net/http"
	"net/http/httptest"
	"encoding/json"
	"testing"

	"github.com/satori/go.uuid"
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
	request := &AccountBalanceRequest{
		Currency: "doesnt",
		Request:  "really",
		Nonce:    "matter",
	}
	_, err = client.PostCurrencyBalance(request)
	assert.True(t, err != nil)
}

func TestPostCurrencyBalance(t *testing.T) {
	pseudoAPIKey := uuid.NewV4()
	pseudoAPISecret := "4a894c5c-8a7e-4337-bb6b-9fde16e3dddd"

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "POST")
		assert.Equal(t, r.URL.String(), "/account/balance")
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

		assert.Equal(t, pseudoAPIKey.String(), r.Header.Get("X-TXC-APIKEY"))
		assert.Equal(t, r.Header.Get("X-TXC-PAYLOAD"), "eyJjdXJyZW5jeSI6ImRvZXNudCIsInJlcXVlc3QiOiJyZWFsbHkiLCJub25jZSI6Im1hdHRlciJ9")
		assert.Equal(t, r.Header.Get("X-TXC-SIGNATURE"), "5dc72166666873eeb2ee9f8c535d79d79191aa34d0d36d711108455b60c81f6f")

		w.WriteHeader(http.StatusOK)

		balanceMap := map[string]AccountBalance{}
		balanceMap["blubb"] = AccountBalance{
			Available: 5.0,
			Freeze: 1.0,
		}
		result := &AccountBalanceResult{
			Balance: balanceMap,
		}
		result.Success = true
		result.Message = ""
		asJSON, err := json.Marshal(result)
		if err!=nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		w.Write(asJSON)
	}))
	defer ts.Close()

	client, err := newClientWithURL(ts.URL, pseudoAPIKey.String(), pseudoAPISecret)
	if err != nil {
		t.Error(err.Error())
	}
	request := &AccountBalanceRequest{
		Currency: "doesnt",
		Request:  "really",
		Nonce:    "matter",
	}
	resp, err := client.PostCurrencyBalance(request)
	assert.Equal(t, true, resp.Success)
}
