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
	product.TotalStock = 5
	product.PriceFrom = 12.90
	product.PriceTo = 12.01
	product.ProductCode = "A11"
	product.ActionPoint = "update"

	repo := repositories.NewProductHistoryRepository(db)
	repo.InsertProductHistory(product)

	got := repo.FindProductHistory("A11")

	require.Equal(t, 1, len(got))
}
