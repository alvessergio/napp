package domain

import (
	"time"

	"github.com/asaskevich/govalidator"
)

type Product struct {
	Code           string            `json:"code" valid:"notnull" gorm:"primary_key"`
	Name           string            `json:"name" valid:"notnull"`
	TotalStock     int               `json:"total_stock" gorm:"type:int" valid:"notnull"`
	CuttingStock   int               `json:"cutting_stock" gorm:"type:int" valid:"notnull"`
	AvailableStock int               `json:"available_stock" gorm:"type:int" valid:"notnull"`
	PriceFrom      float64           `json:"price_from" gorm:"type:decimal(10,2)" valid:"notnull"`
	PriceTo        float64           `json:"price_to" gorm:"type:decimal(10,2)" valid:"notnull"`
	ProductHistory *[]ProductHistory `json:"product_history,omitempty" valid:"-" gorm:"ForeingKey:ProductID"`
	CreatedAt      time.Time         `json:"created_at" valid:"-"`
	UpdatedAt      time.Time         `json:"updated_at" valid:"-"`
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
