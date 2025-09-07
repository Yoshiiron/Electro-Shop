package usecases

import (
	"backend/internal/entity"
)

type ProductService struct {
	repo entity.ProductRepository
}

func NewProductService(repo entity.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}
