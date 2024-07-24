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
    query := "INSERT INTO product (name, description, price, stock, created_at, category) VALUES ($1, $2, $3, $4, $5, $6) RETURNING product_id"
    return r.db.QueryRow(query, p.Name, p.Description, p.Price, p.Stock, p.CreatedAt, p.Category).Scan(&p.ID)
}

func (r *ProductRepository) FindByID(id int) (*product.Product, error) {
    query := "SELECT product_id, name, description, price, stock, created_at, category FROM product WHERE product_id = $1"
    row := r.db.QueryRow(query, id)

    var p product.Product
    if err := row.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Stock, &p.CreatedAt, &p.Category); err != nil {
        if err == sql.ErrNoRows {
            return nil, nil
        }
        return nil, err
    }
    return &p, nil
}
