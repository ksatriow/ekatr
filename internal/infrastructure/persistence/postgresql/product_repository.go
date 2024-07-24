package postgresql

import (
    "database/sql"
    "ekatr/internal/domain/product"
)

type ProductRepository struct {
    db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
    return &ProductRepository{db: db}
}

func (r *ProductRepository) Save(p *product.Product) error {
    query := "INSERT INTO product (name, description, price, stock, created_at) VALUES ($1, $2, $3, $4, $5)"
    _, err := r.db.Exec(query, p.Name, p.Description, p.Price, p.Stock, p.CreatedAt)
    return err
}
