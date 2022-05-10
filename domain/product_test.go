package domain_test

import (
	"encoding/json"
	"testing"

	"github.com/alvessergio/pan-integrations/domain"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestValidateIfProductIsEmpty(t *testing.T) {
	product := domain.NewProduct()

	err := product.Validate()

	require.Error(t, err)
}

func TestValidateIDIsNotAUUID(t *testing.T) {
	product := domain.NewProduct()

	product.ID = "123"
	product.Name = "fake name"
	product.TotalStock = 5
	product.CuttingStock = 3
	product.AvailableStock = 2
	product.PriceFrom = json.Number("12,90")
	product.PriceTo = json.Number("12,01")

	err := product.Validate()

	require.Error(t, err)
}

func TestValidation(t *testing.T) {
	product := domain.NewProduct()

	product.ID = uuid.NewV4().String()
	product.Name = "fake name"
	product.TotalStock = 5
	product.CuttingStock = 3
	product.AvailableStock = 2
	product.PriceFrom = json.Number("12,90")
	product.PriceTo = json.Number("12,01")

	err := product.Validate()

	require.Nil(t, err)
}
