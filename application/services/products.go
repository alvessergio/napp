package services

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/alvessergio/pan-integrations/domain"
	log "github.com/sirupsen/logrus"
)

type productsAPI interface {
	GetProducts(ctx context.Context, traceID string) []*domain.Product
	GetProductById(ctx context.Context, traceID, productID string) (*domain.Product, error)
	PutProduct(ctx context.Context, traceID string, product *domain.Product) (*domain.Product, error)
	PostProduct(ctx context.Context, traceID string, product *domain.Product) (*domain.Product, error)
	DeleteProduct(ctx context.Context, traceID, id string) error
}

type productsServer server

func (p *productsServer) GetProducts(ctx context.Context, traceID string) []*domain.Product {

	products := p.service.ProductRepository.GetAll()
	return products
}

func (p *productsServer) GetProductById(ctx context.Context, traceID, productID string) (*domain.Product, error) {
	product, err := p.service.ProductRepository.Find(productID)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *productsServer) PutProduct(ctx context.Context, traceID string, product *domain.Product) (*domain.Product, error) {

	product, err := p.service.ProductRepository.Update(product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *productsServer) PostProduct(ctx context.Context, traceID string, product *domain.Product) (*domain.Product, error) {

	product, err := p.service.ProductRepository.Insert(product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *productsServer) DeleteProduct(ctx context.Context, traceID, id string) error {
	err := p.service.ProductRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func getProductsHandler(p *productsServer) func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		traceID := ctx.Value("traceID").(string)
		l := log.WithFields(log.Fields{
			"trace_id": traceID,
		})

		products := p.GetProducts(ctx, traceID)

		resp, err := json.Marshal(products)
		if err != nil {
			l.WithFields(log.Fields{
				"event":  "products_serialize_failed",
				"reason": err,
			}).Error("error on serialize products")
			encodeErrorResponse(rw, traceID, err)
			return
		}

		rw.Header().Add("Content-Type", "application/json")
		rw.Write(resp)
	}
}

func getProductByIdHandler(p *productsServer) func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		ctx := r.Context()
		traceID := ctx.Value("traceID").(string)
		id := params.Get("id")
		l := log.WithFields(log.Fields{
			"trace_id": traceID,
		})
		if id == "" {
			l.WithFields(log.Fields{
				"event":  "get_product_by_id_failed_no_id",
				"reason": "id is empty",
			}).Error("error getting product by id, id is empty")
			encodeErrorResponse(rw, traceID, NewError(ErrEmptyParams, "id is empty"))
			return
		}

		product, err := p.GetProductById(ctx, traceID, id)
		if err != nil {
			encodeErrorResponse(rw, traceID, err)
			return
		}

		resp, err := json.Marshal(product)
		if err != nil {
			l.WithFields(log.Fields{
				"event":  "product_serialize_failed",
				"reason": err,
			}).Error("error on serialize product")
			encodeErrorResponse(rw, traceID, err)
			return
		}

		rw.Header().Add("Content-Type", "application/json")
		rw.Write(resp)
	}
}

func putProductHadler(p *productsServer) func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		traceID := ctx.Value("traceID").(string)
		l := log.WithFields(log.Fields{
			"trace_id": traceID,
		})

		defer r.Body.Close()
		var product *domain.Product

		err := json.NewDecoder(r.Body).Decode(&product)
		if err != nil {
			encodeErrorResponse(rw, traceID, ErrEmptyParams)
			return
		}

		if reflect.DeepEqual(product, &domain.Product{}) {
			encodeErrorResponse(rw, traceID, ErrEmptyParams)
			return
		}

		p, err := p.PutProduct(ctx, traceID, product)
		if err != nil {
			encodeErrorResponse(rw, traceID, err)
			return
		}

		resp, err := json.Marshal(p)
		if err != nil {
			l.WithFields(log.Fields{
				"event":  "product_serialize_failed",
				"reason": err,
			}).Error("error on serialize product")
			encodeErrorResponse(rw, traceID, err)
			return
		}

		rw.Header().Add("Content-Type", "application/json")
		rw.Write(resp)
	}
}

func postProductHandler(p *productsServer) func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
	}
}

func deleteProductHandler(p *productsServer) func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {}
}
