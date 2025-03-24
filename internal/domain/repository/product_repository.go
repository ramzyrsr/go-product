package repository

import (
	"database/sql"
	"product/internal/domain/models"
)

type ProductRepository interface {
	Create(product *models.Product) (*models.Product, error)
	GetAllProducts(price, limit, offset int) (map[string]interface{}, error)
}

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) Create(product *models.Product) (*models.Product, error) {
	query := "INSERT INTO products(name, price) VALUES($1, $2) RETURNING uuid"
	err := r.db.QueryRow(query, product.Name, product.Price).Scan(&product.UUID)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *productRepository) GetAllProducts(price, limit, offset int) (map[string]interface{}, error) {
	var products []*models.ProductResponse
	var totalProducts int

	query := `SELECT
		uuid AS id,
		name,
		price,
		description,
		COALESCE(updated_at, created_at) AS updated_at
	FROM
		products`

	if price > 0 {
		query += ` WHERE price = $3`
	}
	query += ` LIMIT $1 OFFSET $2`

	var rows *sql.Rows
	var err error

	if price > 0 {
		rows, err = r.db.Query(query, limit, offset, price)
	} else {
		rows, err = r.db.Query(query, limit, offset)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		product := &models.ProductResponse{}

		var description sql.NullString
		var price sql.NullInt32

		err := rows.Scan(&product.ID, &product.Name, &price, &description, &product.UpdatedAt)
		if err != nil {
			return nil, err
		}

		if description.Valid {
			product.Description = description.String
		} else {
			product.Description = ""
		}

		if price.Valid {
			product.Price = int(price.Int32)
		} else {
			product.Price = 0
		}

		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	countQuery := `SELECT COUNT(id) FROM products`
	if price > 0 {
		countQuery += ` WHERE price = $1`
		err = r.db.QueryRow(countQuery, price).Scan(&totalProducts)
	} else {
		err = r.db.QueryRow(countQuery).Scan(&totalProducts)
	}
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"data":              products,
		"totalDataFiltered": len(products),
		"totalData":         totalProducts,
	}, nil
}
