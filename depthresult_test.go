package p2pb2b

import (
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetDepthResult(t *testing.T) {
	client, err := NewClient(uuid.NewV4().String(), uuid.NewV4().String())
	if err != nil {
		t.Error(err.Error())
	}
	depthResult, err := client.GetDepthResult("ETH_BTC", 10)
	if err != nil {
		t.Error(err.Error())
	}
	assert.NotNil(t, depthResult)
	assert.True(t, depthResult.Success)
	assert.True(t, len(depthResult.DepthResult.Asks) == 10)
	assert.True(t, len(depthResult.DepthResult.Bids) == 10)
}

func TestGetDepthResultNegative(t *testing.T) {
	client, err := NewClient(uuid.NewV4().String(), uuid.NewV4().String())
	if err != nil {
		t.Error(err.Error())
	}

	depthResult, err := client.GetDepthResult("ETH_BTC", 0)
	assert.Nil(t, depthResult)
	assert.NotNil(t, err)

	depthResult, err = client.GetDepthResult("blubb", 1)
	assert.Nil(t, depthResult)
	assert.NotNil(t, err)

	depthResult, err = client.GetDepthResult("ETH_BTC", -5)
	assert.Nil(t, depthResult)
	assert.NotNil(t, err)
}
