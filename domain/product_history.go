package domain

import (
	"encoding/json"
	"time"

	"github.com/asaskevich/govalidator"
)

type ProductHistory struct {
	ID             string      `json:"job_id" valid:"uuid" gorm:"type:uuid;primary_key"`
	Name           string      `json:"name" valid:"notnull"`
	TotalStock     int         `json:"total_stock" valid:"notnull"`
	CuttingStock   int         `json:"cutting_stock" valid:"notnull"`
	AvailableStock int         `json:"available_stock" valid:"notnull"`
	PriceFrom      json.Number `json:"price_from" valid:"notnull"`
	PriceTo        json.Number `json:"price_to" valid:"notnull"`
	CreatedAt      time.Time   `json:"created_at" valid:"-"`
	UpdatedAt      time.Time   `json:"updated_at" valid:"-"`
	ProductID      string      `json:"-" valid:"-" gorm:"column:product_id;type:uuid;notnull"`
	Product        *Product    `json:"product" valid:"-"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func NewProductHistory() *Product {
	return &Product{}
}

func (product *ProductHistory) Validate() error {
	_, err := govalidator.ValidateStruct(product)

	if err != nil {
		return err
	}

	return nil
}
