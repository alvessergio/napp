package domain

import (
	"encoding/json"
	"time"

	"github.com/asaskevich/govalidator"
)

type ProductHistory struct {
	ID             string      `json:"id" valid:"uuid" gorm:"type:uuid;primary_key"`
	Name           string      `json:"name" valid:"notnull"`
	TotalStock     int         `json:"total_stock" valid:"notnull"`
	CuttingStock   int         `json:"cutting_stock" valid:"notnull"`
	AvailableStock int         `json:"available_stock" valid:"notnull"`
	UpdatedAt      time.Time   `json:"updated_at" valid:"-"`
	CreatedAt      time.Time   `json:"created_at" valid:"-"`
	PriceFrom      json.Number `json:"price_from" valid:"notnull"`
	PriceTo        json.Number `json:"price_to" valid:"notnull"`
	ProductCode    string      `json:"-" valid:"-" gorm:"column:product_id;notnull"`
	Product        *Product    `json:"product" valid:"-"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func NewProductHistory() *ProductHistory {
	return &ProductHistory{}
}

func (product *ProductHistory) Validate() error {
	_, err := govalidator.ValidateStruct(product)

	if err != nil {
		return err
	}

	return nil
}
