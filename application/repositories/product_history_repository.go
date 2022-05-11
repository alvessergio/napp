package repositories

import (
	"github.com/alvessergio/pan-integrations/domain"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type PorductHistoryRepository interface {
	InsertProductHistory(product *domain.ProductHistory) (*domain.ProductHistory, error)
}

type ProductHistoryRepositoryDb struct {
	Db *gorm.DB
}

func NewProductHistoryRepository(db *gorm.DB) *ProductHistoryRepositoryDb {
	return &ProductHistoryRepositoryDb{Db: db}
}

func (repo ProductRepositoryDb) InsertProductHistory(product *domain.ProductHistory) (*domain.ProductHistory, error) {
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
