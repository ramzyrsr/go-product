package service

import (
	"errors"
	"product/internal/domain/models"
	"product/internal/domain/repository"
)

type ProductService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) CreateProduct(product *models.Product) (*models.Product, error) {
	return s.repo.Create(product)
}

func (s *ProductService) GetAllProducts(price, limit, offset int) (map[string]interface{}, error) {
	if limit <= 0 || offset < 0 {
		return nil, errors.New("invalid limit or offset values")
	}

	return s.repo.GetAllProducts(price, limit, offset)
}
