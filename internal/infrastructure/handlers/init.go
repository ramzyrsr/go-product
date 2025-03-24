package handlers

import (
	"log"
	"product/internal/domain/repository"
	"product/internal/domain/service"
	"product/internal/infrastructure/db"
)

type Handlers struct {
	ProductHandler *ProductHandler
}

func InitHandlers() *Handlers {
	dbConn, err := db.Connect()
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}

	productRepository := repository.NewProductRepository(dbConn)
	productService := service.NewProductService(productRepository)
	productHandler := NewProductHandler(productService)

	handlers := &Handlers{
		ProductHandler: productHandler,
	}

	return handlers
}
