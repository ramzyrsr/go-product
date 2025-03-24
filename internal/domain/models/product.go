package models

import "time"

type Product struct {
	ID          int       `json:"id"`
	UUID        string    `json:"uuid"`
	Name        string    `json:"name" validate:"required,min=3,max=100"`
	Price       int       `json:"price" validate:"omitempty,min=0"`
	Description string    `json:"description" validate:"omitempty,max=1000"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ProductResponse struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Price       int       `json:"price"`
	Description string    `json:"description"`
	UpdatedAt   time.Time `json:"updated_at"`
}
