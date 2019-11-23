package p2pb2b

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeHeaders(t *testing.T) {
	firstHeaders := map[string]string{}
	firstHeaders["Tintin"] = "Reporter"
	firstHeaders["Haddock"] = "Captain"

	secondsHeaders := map[string]string{}
	secondsHeaders["Snowy"] = "Dog"
	secondsHeaders["Haddock"] = "Sailor"

	resultHeaders := mergeHeaders(firstHeaders, secondsHeaders)

	assert.Equal(t, "Reporter", resultHeaders["Tintin"], "Tintin is a reporter")
	assert.Equal(t, "Captain", resultHeaders["Haddock"], "Haddock is a captain")
	assert.Equal(t, "Dog", resultHeaders["Snowy"], "Snowy is a dog")

	firstHeaders = map[string]string{}
	firstHeaders["Tintin"] = "Reporter"
	firstHeaders["Haddock"] = "Captain"

	secondsHeaders = nil

	resultHeaders = mergeHeaders(firstHeaders, secondsHeaders)

	assert.Equal(t, "Reporter", resultHeaders["Tintin"], "Tintin is a reporter")
	assert.Equal(t, "Captain", resultHeaders["Haddock"], "Haddock is a captain")
}

func TestSendGet(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "GET")
		assert.Equal(t, r.URL.String(), "/somePath")
		assert.Equal(t, r.Header.Get(HeaderXTxcAPIKey), "Ireallydontcare")
	}))
	defer ts.Close()

	auth := &auth{
		APIKey: "Ireallydontcare",
	}
	client := &client{
		http: &http.Client{},
		auth: auth,
	}
	headers := map[string]string{}
	_, err := client.sendGet(fmt.Sprintf("%s/%s", ts.URL, "somePath"), headers)
	if err != nil {
		t.Errorf("error in SendGet, %v", err)
	}
}

func TestSendGetWithAdditionalHeaders(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "GET")
		assert.Equal(t, r.URL.String(), "/somePath")
		assert.Equal(t, r.Header.Get(HeaderXTxcAPIKey), "Ireallydontcare")
	}))
	defer ts.Close()

	auth := &auth{
		APIKey: "Ireallydontcare",
	}
	client := &client{
		http: &http.Client{},
		auth: auth,
	}
	_, err := client.sendGet(fmt.Sprintf("%s/%s", ts.URL, "somePath"), nil)
	if err != nil {
		t.Errorf("error in SendGet, %v", err)
	}
}
