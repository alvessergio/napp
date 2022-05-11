package repositories

import (
	"fmt"

	"github.com/alvessergio/pan-integrations/domain"

	"github.com/jinzhu/gorm"
)

type PorductRepository interface {
	Insert(product *domain.Product) (*domain.Product, error)
	Find(code string) (*domain.Product, error)
	GetAll() []*domain.Product
	Update(product *domain.Product) (*domain.Product, error)
	Delete(code string) error
}

type ProductRepositoryDb struct {
	Db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepositoryDb {
	return &ProductRepositoryDb{Db: db}
}

func (repo ProductRepositoryDb) Insert(product *domain.Product) (*domain.Product, error) {
	if product.Code == "" {
		return nil, fmt.Errorf("code is empty")
	}

	err := repo.Db.Create(product).Error

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (repo *ProductRepositoryDb) Find(code string) (*domain.Product, error) {
	if code == "" {
		return nil, fmt.Errorf("code is empty")
	}
	var product domain.Product

	repo.Db.Find(&product, "code = ?", code)

	if product.Code == "" {
		return nil, fmt.Errorf("product does not exist")
	}

	return &product, nil
}

func (repo *ProductRepositoryDb) GetAll() []*domain.Product {
	var products []*domain.Product

	repo.Db.Find(&products).Order("name ASC")

	return products
}

func (repo *ProductRepositoryDb) Update(product *domain.Product) (*domain.Product, error) {
	if product.Code == "" {
		return nil, fmt.Errorf("code is empty")
	}

	var p domain.Product

	repo.Db.Find(&p, "code = ?", product.Code)

	if p.Code == "" {
		return nil, fmt.Errorf("product does not exist")
	}

	p.Name = product.Name
	p.TotalStock = product.TotalStock
	p.CuttingStock = product.CuttingStock
	p.PriceFrom = product.PriceFrom
	p.PriceTo = product.PriceTo

	repo.Db.Save(&p)

	return product, nil
}

func (repo *ProductRepositoryDb) Delete(code string) error {
	if code == "" {
		return fmt.Errorf("code is empty")
	}

	err := repo.Db.Delete(&domain.Product{Code: code}).Error

	return err
}
