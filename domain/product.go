package domain

import (
	"encoding/json"

	"github.com/asaskevich/govalidator"
)

type Product struct {
	ID             string            `json:"id" valid:"uuid" gorm:"type:uuid;primary_key"`
	Name           string            `json:"name" valid:"notnull"`
	TotalStock     int               `json:"total_stock" valid:"notnull"`
	CuttingStock   int               `json:"cutting_stock" valid:"notnull"`
	AvailableStock int               `json:"available_stock" valid:"notnull"`
	PriceFrom      json.Number       `json:"price_from" valid:"notnull"`
	PriceTo        json.Number       `json:"price_to" valid:"notnull"`
	ProductHistory *[]ProductHistory `json:"-" valid:"-" gorm:"ForeingKey:ProductID"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func (product *Product) Validate() error {
	_, err := govalidator.ValidateStruct(product)

	if err != nil {
		return err
	}

	return nil
}

func NewProduct() *Product {
	return &Product{}
}
