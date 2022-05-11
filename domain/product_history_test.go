package domain_test

import (
	"testing"

	"github.com/alvessergio/pan-integrations/domain"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestValidateIfProductHistoyIsEmpty(t *testing.T) {
	product := domain.NewProductHistory()

	err := product.Validate()

	require.Error(t, err)
}

func TestValidateProductHistoryIDIsNotAUUID(t *testing.T) {
	product := domain.NewProductHistory()

	product.ID = "123"
	product.Name = "fake name"
	product.TotalStock = 5
	product.CuttingStock = 3
	product.AvailableStock = 2
	product.PriceFrom = 12.90
	product.PriceTo = 12.01

	err := product.Validate()

	require.Error(t, err)
}

func TestProductHistoryValidation(t *testing.T) {
	product := domain.NewProductHistory()

	product.ID = uuid.NewV4().String()
	product.Name = "fake name"
	product.TotalStock = 5
	product.CuttingStock = 3
	product.AvailableStock = 2
	product.PriceFrom = 12.90
	product.PriceTo = 12.01
	product.ProductCode = "A11"
	product.ActionPoint = "update"

	err := product.Validate()

	require.Nil(t, err)
}
