package services

type PostProductRequest struct {
	Code         string  `json:"code"`
	Name         string  `json:"name"`
	TotalStock   int     `json:"total_stock"`
	CuttingStock int     `json:"cutting_stock"`
	PriceFrom    float64 `json:"price_from"`
	PriceTo      float64 `json:"price_to"`
}

type PutProductRequest PostProductRequest
