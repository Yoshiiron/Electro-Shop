package memorydb

import "backend/internal/entity"

var products = []entity.Product{
	{ID: 1, Name: "Laptop", Description: "A high-performance laptop", Price: 999.99, Stock: 10},
	{ID: 2, Name: "Smartphone", Description: "A latest model smartphone", Price: 699.99, Stock: 25},
	{ID: 3, Name: "Headphones", Description: "Noise-cancelling headphones", Price: 199.99, Stock: 15},
}

type productRepository struct {
	products []entity.Product
}

func NewProductRepository() *productRepository {
	return &productRepository{products: products}
}
