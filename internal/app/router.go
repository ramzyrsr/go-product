package app

import (
	"log"
	"net/http"
	"product/internal/infrastructure/handlers"
	"product/internal/infrastructure/middleware"

	"github.com/gorilla/mux"
)

func SetupRouter() {
	handler := handlers.InitHandlers()

	r := mux.NewRouter()
	v1 := r.PathPrefix("/api/v1").Subrouter()
	// auth := v1
	v1.Use(middleware.CORS)

	v1.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		middleware.Response(w, http.StatusOK, "Welcome to API", nil)
	}).Methods("GET")

	v1.HandleFunc("/products", handler.ProductHandler.CreateProduct).Methods("POST")
	v1.HandleFunc("/products", handler.ProductHandler.GetProducts).Methods("GET")

	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
