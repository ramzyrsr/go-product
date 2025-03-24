package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"product/internal/domain/models"
	"product/internal/domain/service"
	"product/internal/infrastructure/middleware"
	"strconv"
)

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(service *service.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		middleware.Response(w, http.StatusBadRequest, "Invalid request payload", nil)
		return
	}

	if err := middleware.Validate.Struct(product); err != nil {
		middleware.Response(w, http.StatusBadRequest, fmt.Sprintf("Validation error: %v", err), nil)
		return
	}

	createdProduct, err := h.service.CreateProduct(&product)
	if err != nil {
		middleware.Response(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	middleware.Response(w, http.StatusOK, "Successfully create product", map[string]interface{}{
		"id": createdProduct.UUID,
	})
}

func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	priceStr := r.URL.Query().Get("price")
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")

	price, err := strconv.Atoi(priceStr)
	// if err != nil || price <= 0 {
	// 	// middleware.Response(w, http.StatusBadRequest, "Price must more than or equal 0. Please check your data", nil)
	// 	return
	// }

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		offset = 0
	}

	products, err := h.service.GetAllProducts(price, limit, offset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	middleware.Response(w, http.StatusOK, "Success", products)
}
