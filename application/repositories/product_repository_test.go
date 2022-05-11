package repositories_test

import (
	"testing"
	"time"

	"github.com/alvessergio/pan-integrations/framework/database"

	"github.com/alvessergio/pan-integrations/domain"

	"github.com/alvessergio/pan-integrations/application/repositories"

	"github.com/stretchr/testify/require"
)

func TestProductRepositoryDbInsert(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	product := domain.NewProduct()

	product.Code = "A11"
	product.Name = "fake name"
	product.TotalStock = 5
	product.CuttingStock = 3
	product.AvailableStock = 2
	product.TotalStock = 5
	product.PriceFrom = 12.90
	product.PriceTo = 12.01
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()

	repo := repositories.NewProductRepository(db)
	repo.Insert(product)

	got := repo.Find(product.Code)

	require.NotEmpty(t, got.Code)
	require.NotNil(t, got)
	require.Equal(t, got.Code, product.Code)
}
