package cli_test

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/williamrlbrito/hexagonal-architecture/adapters/cli"
	"github.com/williamrlbrito/hexagonal-architecture/application"
	mock_application "github.com/williamrlbrito/hexagonal-architecture/application/mocks"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := application.NewProduct()
	product.Name = "Product Test"
	product.Price = 25

	productMock := mock_application.NewMockProductInterface(ctrl)
	productMock.EXPECT().GetID().Return(product.GetID()).AnyTimes()
	productMock.EXPECT().GetName().Return(product.GetName()).AnyTimes()
	productMock.EXPECT().GetPrice().Return(product.GetPrice()).AnyTimes()
	productMock.EXPECT().GetStatus().Return(product.GetStatus()).AnyTimes()

	service := mock_application.NewMockProductServiceInterface(ctrl)
	service.EXPECT().Create(product.GetName(), product.GetPrice()).Return(productMock, nil).AnyTimes()
	service.EXPECT().Get(product.GetID()).Return(productMock, nil).AnyTimes()
	service.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	service.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	resultExpected := fmt.Sprintf("Product ID %s with the name %s has been created with the price %f and status %s",
		product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())

	result, err := cli.Run(service, "create", "", product.GetName(), product.GetPrice())

	assert.Nil(t, err)
	assert.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product %s has been enabled", product.GetName())
	result, err = cli.Run(service, "enable", product.GetID(), "", 0)

	assert.Nil(t, err)
	assert.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product %s has been disable", product.GetName())
	result, err = cli.Run(service, "disable", product.GetID(), "", 0)

	assert.Nil(t, err)
	assert.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product ID: %s\nName: %s\nPrice: %f\nStatus: %s",
		product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())

	result, err = cli.Run(service, "get", product.GetID(), "", 0)

	assert.Nil(t, err)
	assert.Equal(t, resultExpected, result)
}
