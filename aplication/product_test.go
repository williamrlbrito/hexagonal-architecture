package aplication_test

import (
	"testing"

	uuid "github.com/satori/go.uuid"
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

func TestProductDisable(t *testing.T) {
	product := aplication.Product{}
	product.Name = "Product Test"
	product.Status = aplication.ENABLED
	product.Price = 0

	err := product.Disable()
	assert.Nil(t, err)

	product.Price = 10
	err = product.Disable()
	assert.Equal(t, "the price must be zero to disable the product", err.Error())
}

func TestProductIsValid(t *testing.T) {
	product := aplication.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "Product Test"
	product.Status = aplication.ENABLED
	product.Price = 10

	_, err := product.IsValid()
	assert.Nil(t, err)

	product.Price = -10
	_, err = product.IsValid()
	assert.Equal(t, "the price must be greater than zero", err.Error())

	product.Price = 10
	product.Status = "invalid"
	_, err = product.IsValid()
	assert.Equal(t, "the status must be enabled or disabled", err.Error())

	product.Status = aplication.ENABLED
	_, err = product.IsValid()
	assert.Nil(t, err)
}
