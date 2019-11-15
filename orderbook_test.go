package p2pb2b

import (
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetOrderBook(t *testing.T) {
	client, err := NewClient(uuid.NewV4().String(), uuid.NewV4().String())
	if err != nil {
		t.Error(err.Error())
	}
	orderBook, err := client.GetOrderBook("ETH_BTC", "sell", 0, 10)
	if err != nil {
		t.Error(err.Error())
	}
	assert.NotNil(t, orderBook)
	assert.True(t, orderBook.Success)
	assert.True(t, len(orderBook.OrderBook.Orders) == 10)
}

func TestGetOrderBookNegative(t *testing.T) {
	client, err := NewClient(uuid.NewV4().String(), uuid.NewV4().String())
	if err != nil {
		t.Error(err.Error())
	}

	orderBook, err := client.GetOrderBook("ETH_BTC", "sell", -1, 10)
	assert.Nil(t, orderBook)
	assert.NotNil(t, err)

	orderBook, err = client.GetOrderBook("ETH_BTC", "sell", 2, 0)
	assert.Nil(t, orderBook)
	assert.NotNil(t, err)

	orderBook, err = client.GetOrderBook("ETH_BTC", "sell", 2, -5)
	assert.Nil(t, orderBook)
	assert.NotNil(t, err)

	orderBook, err = client.GetOrderBook("blubb", "sell", 0, 10)
	assert.Nil(t, orderBook)
	assert.NotNil(t, err)

	orderBook, err = client.GetOrderBook("ETH_BTC", "blubb", 0, 10)
	assert.Nil(t, orderBook)
	assert.NotNil(t, err)
}
