package repositories

import (
	"fmt"

	"github.com/alvessergio/pan-integrations/domain"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type PorductRepository interface {
	Insert(product *domain.Product) (*domain.Product, error)
	Find(id string) (*domain.Product, error)
}

type ProductRepositoryDb struct {
	Db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepositoryDb {
	return &ProductRepositoryDb{Db: db}
}

func (repo ProductRepositoryDb) Insert(product *domain.Product) (*domain.Product, error) {
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

func (repo ProductRepositoryDb) Find(id string) (*domain.Product, error) {
	var product domain.Product

	repo.Db.Find(&product, "id = ?", id)

	if product.ID == "" {
		return nil, fmt.Errorf("product does not exist")
	}

	return &product, nil
}
