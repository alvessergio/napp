package repositories

import (
	"github.com/alvessergio/pan-integrations/domain"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type PorductHistoryRepository interface {
	InsertProductHistory(product *domain.ProductHistory) (*domain.ProductHistory, error)
	FindProductHistory(productCode string) []*domain.ProductHistory
}

type ProductHistoryRepositoryDb struct {
	Db *gorm.DB
}

func NewProductHistoryRepository(db *gorm.DB) *ProductHistoryRepositoryDb {
	return &ProductHistoryRepositoryDb{Db: db}
}

func (repo *ProductHistoryRepositoryDb) InsertProductHistory(product *domain.ProductHistory) (*domain.ProductHistory, error) {
	if product.ID == "" {
		id := uuid.NewV4().String()
		product.ID = id
	}

	err := repo.Db.Create(product).Error

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (repo *ProductHistoryRepositoryDb) FindProductHistory(productCode string) []*domain.ProductHistory {
	var audities []*domain.ProductHistory
	repo.Db.Find(&audities, "product_code = ?", productCode).Order("updated_at DESC")
	return audities
}
