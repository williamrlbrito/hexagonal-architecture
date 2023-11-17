package aplication_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/williamrlbrito/hexagonal-architecture/aplication"
)

func TestProductEnable(t *testing.T) {
	product := aplication.Product{}
	product.Name = "Product Test"
	product.Status = aplication.DISABLED
	product.Price = 10

	err := product.Enable()
	assert.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	assert.Equal(t, "the price must be greater than zero to enable the product", err.Error())
}
