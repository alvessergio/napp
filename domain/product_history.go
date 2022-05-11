package domain

import (
	"time"

	"github.com/asaskevich/govalidator"
)

type ProductHistory struct {
	ID             string    `json:"id" valid:"uuid" gorm:"type:uuid;primary_key"`
	Name           string    `json:"name" valid:"notnull"`
	TotalStock     int       `json:"total_stock" gorm:"type:int" valid:"notnull"`
	CuttingStock   int       `json:"cutting_stock" gorm:"type:int" valid:"notnull"`
	AvailableStock int       `json:"available_stock" gorm:"type:int" valid:"notnull"`
	PriceFrom      float64   `json:"price_from" gorm:"type:decimal(10,2)" valid:"notnull"`
	PriceTo        float64   `json:"price_to" gorm:"type:decimal(10,2)" valid:"notnull"`
	UpdatedAt      time.Time `json:"updated_at" valid:"-"`
	CreatedAt      time.Time `json:"created_at" valid:"-"`
	ProductCode    string    `json:"-" valid:"-" gorm:"column:product_code;notnull"`
	Product        *Product  `json:"product" valid:"-"`
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
