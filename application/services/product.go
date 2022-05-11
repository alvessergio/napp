package services

import "encoding/json"

type PostProductRequest struct {
	Code         string
	Name         string
	TotalStock   int
	CuttingStock int
	PriceFrom    json.Number
	PriceTo      json.Number
}

type PutProductRequest PostProductRequest
