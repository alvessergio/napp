package services

import (
	"time"

	"github.com/alvessergio/pan-integrations/domain"
)

func isValueMinusThan(value, comparingValue float64) bool {
	return value < comparingValue
}

func castPOSTRequestToProduct(req PostProductRequest) *domain.Product {

	pro := &domain.Product{
		Code:         req.Code,
		Name:         req.Name,
		CuttingStock: req.CuttingStock,
		TotalStock:   req.TotalStock,
		PriceFrom:    req.PriceFrom,
		PriceTo:      req.PriceTo,
	}

	pro.AvailableStock = pro.TotalStock - pro.CuttingStock
	pro.CreatedAt = time.Now()
	pro.UpdatedAt = pro.CreatedAt

	return pro
}

func castPUTRequestToProduct(req PutProductRequest) *domain.Product {

	pro := &domain.Product{
		Code:         req.Code,
		Name:         req.Name,
		CuttingStock: req.CuttingStock,
		TotalStock:   req.TotalStock,
		PriceFrom:    req.PriceFrom,
		PriceTo:      req.PriceTo,
	}

	pro.AvailableStock = pro.TotalStock - pro.CuttingStock
	pro.UpdatedAt = time.Now()

	return pro
}

func castRequestToProductHistory(pro *domain.Product, action string) *domain.ProductHistory {
	product := &domain.ProductHistory{
		ProductCode:    pro.Code,
		Name:           pro.Name,
		CuttingStock:   pro.CuttingStock,
		TotalStock:     pro.TotalStock,
		PriceFrom:      pro.PriceFrom,
		PriceTo:        pro.PriceTo,
		UpdatedAt:      pro.UpdatedAt,
		CreatedAt:      pro.CreatedAt,
		AvailableStock: pro.AvailableStock,
		ActionPoint:    action,
	}

	return product
}
