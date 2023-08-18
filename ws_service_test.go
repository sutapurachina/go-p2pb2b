package p2pb2b

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLastPriceStream(t *testing.T) {
	lastPriceC, _, err := LastPriceStream("ETH_USDT")
	assert.NoError(t, err)
	for {
		select {
		case res := <-lastPriceC:
			fmt.Println(res)
		}
	}
}
