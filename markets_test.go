package p2pb2b

import (
	"testing"
	"github.com/satori/go.uuid"
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
	t.Logf("found %d markets", len(markets.Markets))
	t.Logf("sample market %+v", markets.Markets[0])

	assert.NotNil(t, markets)
	assert.True(t, markets.Success)
	assert.True(t, len(markets.Markets) > 0)
}
