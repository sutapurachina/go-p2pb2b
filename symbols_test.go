package p2pb2b

import (
	"testing"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetSymbols(t *testing.T) {
	client, err := NewClient(uuid.NewV4().String(), uuid.NewV4().String())
	if err != nil {
		t.Error(err.Error())
	}
	symbols, err := client.GetSymbols()
	if err != nil {
		t.Error(err.Error())
	}
	t.Logf("found %d symbols", len(symbols.Symbols))
	t.Logf("sample symbol %+v", symbols.Symbols[0])

	assert.NotNil(t, symbols)
	assert.True(t, symbols.Success)
	assert.True(t, len(symbols.Symbols) > 0)
}
