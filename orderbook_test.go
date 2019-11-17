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
	assert.NotEmpty(t, orderBook.OrderBook.Orders[0].ID)
	assert.NotEmpty(t, orderBook.OrderBook.Orders[0].Market)
	assert.NotEmpty(t, orderBook.OrderBook.Orders[0].Side)
	assert.NotEmpty(t, orderBook.OrderBook.Orders[0].Type)
	assert.NotEmpty(t, orderBook.OrderBook.Orders[0].Left)
	assert.NotEmpty(t, orderBook.OrderBook.Orders[0].Price)
	assert.NotEmpty(t, orderBook.OrderBook.Orders[0].TakerFee)
	assert.NotEmpty(t, orderBook.OrderBook.Orders[0].MakerFee)
	assert.True(t, orderBook.OrderBook.Orders[0].DealFee >= 0)
	assert.True(t, orderBook.OrderBook.Orders[0].DealMoney >= 0)
	assert.True(t, orderBook.OrderBook.Orders[0].DealStock >= 0)
	assert.NotEmpty(t, orderBook.OrderBook.Orders[0].Timestamp)
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
