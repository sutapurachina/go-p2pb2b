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
			Request: "doesnt",
			Nonce:   "matter",
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
		assert.Equal(t, "eyJyZXF1ZXN0IjoiZG9lc250Iiwibm9uY2UiOiJtYXR0ZXIiLCJjdXJyZW5jeSI6IkVUSCJ9", r.Header.Get("X-TXC-PAYLOAD"))
		assert.Equal(t, "f4967e290d9b86df09db3551c4b930c35288d77911fe97819e663af35a97c6a5", r.Header.Get("X-TXC-SIGNATURE"))

		expectedReqBody := `{
			"currency": "ETH",
			"request": "{{request}}",
			"nonce": "{{nonce}}"
		}`
		reqBody, _ := ioutil.ReadAll(r.Body)
		fmt.Println(string(reqBody))
		fmt.Println(expectedReqBody)
		equal, err := IsEqualJSON(expectedReqBody, string(reqBody))
		assert.Nil(t, err, err)
		assert.True(t, equal, fmt.Sprintf("%v is not equal to %v", reqBody, expectedReqBody))
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
			Request: "doesnt",
			Nonce:   "matter",
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
