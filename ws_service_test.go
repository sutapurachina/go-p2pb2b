package p2pb2b

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLastPriceStream(t *testing.T) {
	lastPriceC, stopC, doneC, err := LastPriceStream("ETH_USDT")
	assert.NoError(t, err)
	for i := 0; i < 5; i++ {
		select {
		case res := <-lastPriceC:
			fmt.Println(res)
		}
	}
	stopC <- struct{}{}
	<-doneC
}
