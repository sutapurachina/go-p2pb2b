package p2pb2b

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTickers(t *testing.T) {
	client, err := NewClient()
	if err != nil {
		t.Error(err.Error())
	}
	tickers, err := client.GetTickers()
	if err != nil {
		t.Error(err.Error())
	}
	t.Logf("tickers %v", tickers)

	assert.NotNil(t, tickers)
	assert.True(t, tickers.Success)
	assert.True(t, len(tickers.Result) > 0)
}
