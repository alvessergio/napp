package services

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/alvessergio/pan-integrations/domain"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type productsAPI interface {
	GetProducts(ctx context.Context, traceID string) []*domain.Product
	GetProductByCode(ctx context.Context, traceID, code string) (*domain.Product, error)
	PutProduct(ctx context.Context, traceID string, product *domain.Product) (*domain.Product, error)
	PostProduct(ctx context.Context, traceID string, product *domain.Product) (*domain.Product, error)
	DeleteProduct(ctx context.Context, traceID, code string) error
	PostProductHistory(ctx context.Context, traceID string, product *domain.ProductHistory) (*domain.ProductHistory, error)
}

type productsServer server

func (p *productsServer) GetProducts(ctx context.Context, traceID string) []*domain.Product {
	products := p.service.ProductRepository.GetAll()
	return products
}

func (p *productsServer) GetProductByCode(ctx context.Context, traceID, code string) (*domain.Product, error) {
	product, err := p.service.ProductRepository.Find(code)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *productsServer) PutProduct(ctx context.Context, traceID string, product *domain.Product) (*domain.Product, error) {
	pro, err := p.service.ProductRepository.Update(product)
	if err != nil {
		return nil, err
	}

	return pro, nil
}

func (p *productsServer) PostProduct(ctx context.Context, traceID string, product *domain.Product) (*domain.Product, error) {
	pro, err := p.service.ProductRepository.Insert(product)
	if err != nil {
		return nil, err
	}

	return pro, nil
}

func (p *productsServer) DeleteProduct(ctx context.Context, traceID, code string) error {
	err := p.service.ProductRepository.Delete(code)
	if err != nil {
		return err
	}

	return nil
}

func (p *productsServer) PostProductHistory(ctx context.Context, traceID string, product *domain.ProductHistory) (*domain.ProductHistory, error) {
	pro, err := p.service.ProductHistoryRepository.InsertProductHistory(product)
	if err != nil {
		return nil, err
	}

	return pro, nil
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

		rw.Write(resp)
	}
}

func getProductByCodeHandler(p *productsServer) func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		ctx := r.Context()
		traceID := ctx.Value("traceID").(string)
		code := vars["code"]
		l := log.WithFields(log.Fields{
			"trace_id": traceID,
		})
		if code == "" {
			l.WithFields(log.Fields{
				"event":  "get_product_by_code_failed_no_code",
				"reason": "code is empty",
			}).Error("error getting product by code, code is empty")
			encodeErrorResponse(rw, traceID, NewError(ErrEmptyParams, "code is empty"))
			return
		}

		product, err := p.GetProductByCode(ctx, traceID, code)
		if err != nil {
			l.WithFields(log.Fields{
				"event":  "put_product_failed",
				"reason": "internal error",
			}).Error(err.Error())
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

		rw.Write(resp)
	}
}

func putProductHadler(p *productsServer) func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		ctx := r.Context()
		traceID := ctx.Value("traceID").(string)
		code := vars["code"]
		l := log.WithFields(log.Fields{
			"trace_id": traceID,
		})
		if code == "" {
			l.WithFields(log.Fields{
				"event":  "put_product_failed_no_code",
				"reason": "code is empty",
			}).Error("error update product by code, code is empty")
			encodeErrorResponse(rw, traceID, NewError(ErrEmptyParams, "code is empty"))
			return
		}

		defer r.Body.Close()
		var productReq PutProductRequest

		err := json.NewDecoder(r.Body).Decode(&productReq)
		if err != nil {
			l.WithFields(log.Fields{
				"event":  "put_product_failed_incorrect_request",
				"reason": "incorrect request",
			}).Error("error update product by code, incorrect request")
			encodeErrorResponse(rw, traceID, NewError(ErrEmptyParams, "incorrect request"))
			return
		}

		if reflect.DeepEqual(productReq, PutProductRequest{}) {
			l.WithFields(log.Fields{
				"event":  "put_product_failed_empty_request",
				"reason": "request is empty",
			}).Error("error update product by code, request is empty")
			encodeErrorResponse(rw, traceID, NewError(ErrEmptyParams, "empty request"))
			return
		}

		if isValueMinusThan(productReq.PriceFrom, productReq.PriceTo) {
			l.WithFields(log.Fields{
				"event":  "post_product_failed_price_validation",
				"reason": "price from can not be minus than price to",
			}).Error("error create a product, price validation")
			encodeErrorResponse(rw, traceID, NewError(ErrValidation, "price from can not be minus than price to"))
		}

		product := castPUTRequestToProduct(productReq)

		product.Code = code

		pro, err := p.PutProduct(ctx, traceID, product)
		if err != nil {
			l.WithFields(log.Fields{
				"event":  "put_product_failed",
				"reason": "internal error",
			}).Error("error update product by code, internal error")
			encodeErrorResponse(rw, traceID, err)
			return
		}

		productHistory := castRequestToProductHistory(pro)

		proHistory, err := p.PostProductHistory(ctx, traceID, productHistory)

		if err != nil {
			l.WithFields(log.Fields{
				"event":   "post_producthostory_failed",
				"reason":  err.Error(),
				"history": proHistory,
			}).Warn("error create product history")
		}

		resp, err := json.Marshal(pro)
		if err != nil {
			l.WithFields(log.Fields{
				"event":  "product_serialize_failed",
				"reason": err,
			}).Error("error on serialize product")
			encodeErrorResponse(rw, traceID, err)
			return
		}

		rw.Write(resp)
	}
}

func postProductHandler(p *productsServer) func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		traceID := ctx.Value("traceID").(string)
		l := log.WithFields(log.Fields{
			"trace_id": traceID,
		})

		defer r.Body.Close()
		var productReq *PostProductRequest

		err := json.NewDecoder(r.Body).Decode(&productReq)
		if err != nil {
			l.WithFields(log.Fields{
				"event":  "post_product_failed_incorrect_request",
				"reason": "incorrect request",
			}).Error("error create a product, incorrect request")
			encodeErrorResponse(rw, traceID, NewError(ErrEmptyParams, "incorrect body"))
			return
		}

		if reflect.DeepEqual(productReq, PostProductRequest{}) {
			l.WithFields(log.Fields{
				"event":  "post_product_failed_empty_request",
				"reason": "empty request",
			}).Error("error create a product, empty request")
			encodeErrorResponse(rw, traceID, NewError(ErrEmptyParams, "empty body"))
			return
		}

		if isValueMinusThan(productReq.PriceFrom, productReq.PriceTo) {
			l.WithFields(log.Fields{
				"event":  "post_product_failed_price_validation",
				"reason": "price from can not be minus than price to",
			}).Error("error create a product, price validation")
			encodeErrorResponse(rw, traceID, NewError(ErrValidation, "price from can not be minus than price to"))
		}

		product := castPOSTRequestToProduct(*productReq)

		pro, err := p.PostProduct(ctx, traceID, product)
		if err != nil {
			l.WithFields(log.Fields{
				"event":  "post_product_failed",
				"reason": "internal error",
			}).Error("error create product, internal error")
			encodeErrorResponse(rw, traceID, err)
			return
		}

		productHistory := castRequestToProductHistory(pro)

		proHistory, err := p.PostProductHistory(ctx, traceID, productHistory)

		if err != nil {
			l.WithFields(log.Fields{
				"event":   "post_producthostory_failed",
				"reason":  err.Error(),
				"history": proHistory,
			}).Warn("error create product history")
		}

		resp, err := json.Marshal(pro)
		if err != nil {
			l.WithFields(log.Fields{
				"event":  "product_serialize_failed",
				"reason": err,
			}).Error("error on serialize product")
			encodeErrorResponse(rw, traceID, err)
			return
		}

		rw.Write(resp)
	}
}

func deleteProductHandler(p *productsServer) func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		ctx := r.Context()
		traceID := ctx.Value("traceID").(string)
		code := vars["code"]
		l := log.WithFields(log.Fields{
			"trace_id": traceID,
		})
		if code == "" {
			l.WithFields(log.Fields{
				"event":  "delete_product_failed_no_code",
				"reason": "code is empty",
			}).Error("error delete product by code, code is empty")
			encodeErrorResponse(rw, traceID, NewError(ErrEmptyParams, "code is empty"))
			return
		}

		err := p.DeleteProduct(ctx, traceID, code)
		if err != nil {
			l.WithFields(log.Fields{
				"event":  "delete_product_failed",
				"reason": "internal error",
			}).Error("error delete product by code, internal error")
			encodeErrorResponse(rw, traceID, err)
			return
		}

		resp, err := json.Marshal("{}")
		if err != nil {
			l.WithFields(log.Fields{
				"event":  "product_serialize_failed",
				"reason": err,
			}).Error("error on serialize response")
			encodeErrorResponse(rw, traceID, err)
			return
		}

		rw.Write(resp)
	}
}
