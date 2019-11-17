package p2pb2b

import (
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetMarkets(t *testing.T) {
	client, err := NewClient(uuid.NewV4().String(), uuid.NewV4().String())
	if err != nil {
		t.Error(err.Error())
	}
	markets, err := client.GetMarkets()
	if err != nil {
		t.Error(err.Error())
	}

	assert.NotNil(t, markets)
	assert.True(t, markets.Success)
	assert.True(t, len(markets.Markets) > 0)
	assert.NotEmpty(t, markets.Markets[0].Name)
	assert.NotEmpty(t, markets.Markets[0].Stock)
	assert.NotEmpty(t, markets.Markets[0].StockPrec)
	assert.NotEmpty(t, markets.Markets[0].Money)
	assert.NotEmpty(t, markets.Markets[0].MoneyPrec)
	assert.NotEmpty(t, markets.Markets[0].MinAmount)
	assert.NotEmpty(t, markets.Markets[0].FeePrec)
}
