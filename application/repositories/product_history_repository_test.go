package repositories_test

import (
	"testing"

	"github.com/alvessergio/pan-integrations/framework/database"

	"github.com/alvessergio/pan-integrations/domain"

	"github.com/alvessergio/pan-integrations/application/repositories"

	"github.com/stretchr/testify/require"
)

func TestProductHistoryRepositoryDbInsert(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	product := domain.NewProductHistory()

	product.ID = "132223e5-63ac-4805-b36f-db318d42aa75"
	product.Name = "fake name"
	product.TotalStock = 5
	product.CuttingStock = 3
	product.AvailableStock = 2
	product.PriceFrom = 12.90
	product.PriceTo = 12.01

	repo := repositories.NewProductHistoryRepository(db)
	repo.InsertProductHistory(product)

	got, err := repo.FindProductHistory(product.ID)

	require.NotEmpty(t, got.ID)
	require.Nil(t, err)
	require.Equal(t, got.ID, product.ID)
}
