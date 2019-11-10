package p2pb2b

import (
	"testing"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetTicker(t *testing.T) {
	client, err := NewClient(uuid.NewV4().String(), uuid.NewV4().String())
	if err != nil {
		t.Error(err.Error())
	}
	ticker, err := client.GetTicker("ETH_BTC")
	if err != nil {
		t.Error(err.Error())
	}
	t.Logf("sample ticker %+v", ticker.Ticker)

	assert.NotNil(t, ticker)
	assert.True(t, ticker.Success)
	assert.NotEmpty(t, ticker.Ticker)
	assert.True(t, ticker.Ticker.Bid >= 0)
}
