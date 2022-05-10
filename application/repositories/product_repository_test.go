package repositories_test

import (
	"encoding/json"
	"testing"

	"github.com/alvessergio/pan-integrations/framework/database"

	"github.com/alvessergio/pan-integrations/domain"

	"github.com/alvessergio/pan-integrations/application/repositories"

	"github.com/stretchr/testify/require"
)

func TestProductRepositoryDbInsert(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	product := domain.NewProduct()

	product.ID = "132223e5-63ac-4805-b36f-db318d42aa75"
	product.Name = "fake name"
	product.TotalStock = 5
	product.CuttingStock = 3
	product.AvailableStock = 2
	product.PriceFrom = json.Number("12,90")
	product.PriceTo = json.Number("12,01")

	repo := repositories.ProductRepositoryDb{Db: db}
	repo.Insert(product)

	got, err := repo.Find(product.ID)

	require.NotEmpty(t, got.ID)
	require.Nil(t, err)
	require.Equal(t, got.ID, product.ID)
}
