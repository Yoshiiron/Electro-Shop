package httphandler

import "backend/internal/usecases"

type ProductHandler struct {
	service *usecases.ProductService
}

func NewProductHandler(service *usecases.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}
