package p2pb2b

import (
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetHistory(t *testing.T) {
	client, err := NewClient(uuid.NewV4().String(), uuid.NewV4().String())
	if err != nil {
		t.Error(err.Error())
	}
	history, err := client.GetHistory("ETH_BTC", 0, 10)
	if err != nil {
		t.Error(err.Error())
	}
	assert.NotNil(t, history)
	assert.True(t, history.Success)
	assert.True(t, len(history.HistoryEntries) == 10)
}

func TestGetHistoryNegative(t *testing.T) {
	client, err := NewClient(uuid.NewV4().String(), uuid.NewV4().String())
	if err != nil {
		t.Error(err.Error())
	}

	history, err := client.GetHistory("ETH_BTC", -1, 10)
	assert.Nil(t, history)
	assert.NotNil(t, err)

	history, err = client.GetHistory("ETH_BTC", 2, 0)
	assert.Nil(t, history)
	assert.NotNil(t, err)

	history, err = client.GetHistory("ETH_BTC", 2, -5)
	assert.Nil(t, history)
	assert.NotNil(t, err)

	history, err = client.GetHistory("blubb", 0, 10)
	assert.Nil(t, history)
	assert.NotNil(t, err)
}
