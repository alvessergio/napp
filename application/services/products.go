package services

import (
	"context"
	"net/http"

	"github.com/alvessergio/pan-integrations/domain"
)

type productsAPI interface {
	GetProducts(ctx context.Context, productID string) ([]*domain.Product, error)
	GetProductById(ctx context.Context, productID string) (*domain.Product, error)
	PutProduct(ctx context.Context, product *domain.Product) (*domain.Product, error)
	PostProduct(ctx context.Context, product *domain.Product) (*domain.Product, error)
	DeleteProduct(ctx context.Context, id string) error
}

type productsServer server

func (p *productsServer) GetProducts(ctx context.Context, productID string) ([]*domain.Product, error) {

	products, err := p.service.ProductRepository.GetAll()
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (p *productsServer) GetProductById(ctx context.Context, productID string) (*domain.Product, error) {
	product, err := p.service.ProductRepository.Find(productID)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *productsServer) PutProduct(ctx context.Context, product *domain.Product) (*domain.Product, error) {
	return nil, nil
}

func (p *productsServer) PostProduct(ctx context.Context, product *domain.Product) (*domain.Product, error) {

	product, err := p.service.ProductRepository.Insert(product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *productsServer) DeleteProduct(ctx context.Context, id string) error {
	return nil
}

func getProductsHandler(p *productsServer) func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {}
}

func getProductByIdHandler(p *productsServer) func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {}
}

func putProductHadler(p *productsServer) func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {}
}

func postProductHandler(p *productsServer) func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {}
}

func deleteProductHandler(p *productsServer) func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {}
}
