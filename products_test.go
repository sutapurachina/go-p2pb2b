package p2pb2b

import (
	"testing"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetProducts(t *testing.T) {
	client, err := NewClient(uuid.NewV4().String(), uuid.NewV4().String())
	if err != nil {
		t.Error(err.Error())
	}
	products, err := client.GetProducts()
	if err != nil {
		t.Error(err.Error())
	}
	t.Logf("found %d products", len(products.Products))
	t.Logf("sample product %+v", products.Products[0])

	assert.NotNil(t, products)
	assert.True(t, products.Success)
	assert.True(t, len(products.Products) > 0)
}
