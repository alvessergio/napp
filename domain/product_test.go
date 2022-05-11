package domain_test

import (
	"testing"
	"time"

	"github.com/alvessergio/pan-integrations/domain"

	"github.com/stretchr/testify/require"
)

func TestValidateIfProductIsEmpty(t *testing.T) {
	product := domain.NewProduct()

	err := product.Validate()

	require.Error(t, err)
}

func TestValidateIDIsEmpty(t *testing.T) {
	product := domain.NewProduct()

	product.Code = ""
	product.Name = "fake name"
	product.TotalStock = 5
	product.CuttingStock = 3
	product.PriceFrom = 12.90
	product.PriceTo = 12.01
	product.AvailableStock = 2
	product.UpdatedAt = time.Now()
	product.CreatedAt = time.Now()

	err := product.Validate()

	require.Error(t, err)
}

func TestValidation(t *testing.T) {
	product := domain.NewProduct()

	product.Code = "abc"
	product.Name = "fake name"
	product.TotalStock = 5
	product.CuttingStock = 3
	product.PriceFrom = 12.90
	product.PriceTo = 12.01
	product.AvailableStock = 2
	product.UpdatedAt = time.Now()
	product.CreatedAt = time.Now()

	err := product.Validate()

	require.Nil(t, err)
}
