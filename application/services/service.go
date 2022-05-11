package services

import "github.com/alvessergio/pan-integrations/application/repositories"

type Service struct {
	ProductRepository        repositories.PorductRepository
	ProductHistoryRepository repositories.PorductHistoryRepository
}

func NewService(productRepository repositories.PorductRepository, productHistoryService repositories.PorductHistoryRepository) *Service {

	return &Service{
		ProductRepository:        productRepository,
		ProductHistoryRepository: productHistoryService,
	}
}
