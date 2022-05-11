package services

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	// productsPath represents products endpoint
	productsPath = "/products"
	// productsPath represents products endpoint search by id
	productsByCodePath = productsPath + "/{code}"
	// HealthPath represents health endpoint
	healthPath = "/health"
)

type API interface {
	productsAPI
}

type server struct {
	service *Service
}

func (s *Service) NewServer() *mux.Router {
	router := mux.NewRouter().
		PathPrefix("/api/v1/"). // add prefix for v1 api `/api/v1/`
		Subrouter()

	svr := &server{
		service: s,
	}

	productSvr := (*productsServer)(svr)

	router.Use(traceIDMiddleware)
	router.HandleFunc(productsPath, getProductsHandler(productSvr)).Methods(http.MethodGet)
	router.HandleFunc(productsByCodePath, getProductByCodeHandler(productSvr)).Methods(http.MethodGet)
	router.HandleFunc(productsPath, postProductHandler(productSvr)).Methods(http.MethodPost)
	router.HandleFunc(productsByCodePath, putProductHadler(productSvr)).Methods(http.MethodPut)
	router.HandleFunc(productsByCodePath, deleteProductHandler(productSvr)).Methods(http.MethodDelete)
	router.HandleFunc(healthPath, healthHandler).Methods(http.MethodGet)

	return router
}

func healthHandler(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode("Healthy")
}
