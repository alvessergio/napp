package services

import "github.com/alvessergio/pan-integrations/application/repositories"

type Service struct {
	ProductRepository repositories.PorductRepository
}

func NewService(productRepository repositories.PorductRepository) *Service {

	return &Service{
		ProductRepository: productRepository,
	}
}
