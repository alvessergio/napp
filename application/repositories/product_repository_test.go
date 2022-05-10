package repositories_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/alvessergio/pan-integrations/framework/database"
	"github.com/jinzhu/gorm"

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

func TestProductRepositoryDbUpdate(t *testing.T) {
	db := prepare()
	defer db.Close()

	product := domain.NewProduct()

	product.ID = "132223e5-63ac-4805-b36f-db318d42aa75"
	product.Name = "fake name"
	product.TotalStock = 6
	product.CuttingStock = 7
	product.AvailableStock = 1
	product.PriceFrom = json.Number("10,00")
	product.PriceTo = json.Number("9,01")

	repo := repositories.ProductRepositoryDb{Db: db}
	got, err := repo.Update(product)

	require.NotEmpty(t, got.ID)
	require.Nil(t, err)
	require.Equal(t, got.ID, product.ID)

	got, err = repo.Find(product.ID)
	require.Nil(t, err)
	require.Equal(t, fmt.Sprint(got), fmt.Sprint(product))
}

func TestProductRepositoryDbDelete(t *testing.T) {
	db := prepare()
	defer db.Close()

	repo := repositories.ProductRepositoryDb{Db: db}

	id := "132223e5-63ac-4805-b36f-db318d42aa75"

	err := repo.Delete(id)

	require.Nil(t, err)

	got, err := repo.Find(id)

	require.Nil(t, got)
	require.Equal(t, fmt.Errorf("product does not exist"), err)
}

func TestProductRepositoryDbGetAll(t *testing.T) {
	db := prepare()
	defer db.Close()

	repo := repositories.ProductRepositoryDb{Db: db}
	products := repo.GetAll()

	require.Equal(t, 1, len(products))
}

func prepare() *gorm.DB {
	db := database.NewDbTest()

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

	return db
}
