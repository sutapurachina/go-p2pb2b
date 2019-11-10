package p2pb2b

import (
	"testing"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetTickers(t *testing.T) {
	client, err := NewClient(uuid.NewV4().String(), uuid.NewV4().String())
	if err != nil {
		t.Error(err.Error())
	}
	tickers, err := client.GetTickers()
	if err != nil {
		t.Error(err.Error())
	}
	t.Logf("found %d tickers", len(tickers.Tickers))
	t.Logf("sample ticker %+v", tickers.Tickers["ETH_BTC"])

	assert.NotNil(t, tickers)
	assert.True(t, tickers.Success)
	assert.True(t, len(tickers.Tickers) > 0)
}
