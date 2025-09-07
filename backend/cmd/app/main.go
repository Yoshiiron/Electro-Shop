package main

import (
	handler "backend/internal/controller/http-handler"
	memorydb "backend/internal/repo/memory-db"
	"backend/internal/usecases"
)

func main() {
	repo := memorydb.NewProductRepository()
	usecase := usecases.NewProductService(repo)
	//For now it will stay like this, later we will add router and server
	handler.NewProductHandler(usecase)
}
