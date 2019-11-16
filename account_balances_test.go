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
		assert.Equal(t, "eyJyZXF1ZXN0IjoiZG9lc250Iiwibm9uY2UiOiJtYXR0ZXIifQ==", r.Header.Get("X-TXC-PAYLOAD"))
		assert.Equal(t, "dbc97a56c2949b099865f0cb11b3e4f01b9b00a519e8f7a995b2b6d9b7fedef4", r.Header.Get("X-TXC-SIGNATURE"))

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
			Request: "doesnt",
			Nonce:   "matter",
		},
	}
	resp, err := client.PostBalances(request)

	assert.NotNil(t, resp, fmt.Sprintf("error: %v", err))
	assert.Equal(t, true, resp.Success)

	respBytes, _ := json.Marshal(resp)
	equal, _ := IsEqualJSON(body, string(respBytes))
	assert.True(t, equal, fmt.Sprintf("%v is not equal to %v", body, resp))
}
