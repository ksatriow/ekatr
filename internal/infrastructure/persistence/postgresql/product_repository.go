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
    query := "INSERT INTO product (name, description, price, stock, created_at, category, product_image) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING product_id"
    return r.db.QueryRow(query, p.Name, p.Description, p.Price, p.Stock, p.CreatedAt, p.Category, p.ProductImage).Scan(&p.ID)
}

func (r *ProductRepository) FindByID(id int) (*product.Product, error) {
    query := "SELECT product_id, name, description, price, stock, created_at, category, product_image FROM product WHERE product_id = $1"
    row := r.db.QueryRow(query, id)

    var p product.Product
    if err := row.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Stock, &p.CreatedAt, &p.Category, &p.ProductImage); err != nil {
        if err == sql.ErrNoRows {
            return nil, nil
        }
        return nil, err
    }
    return &p, nil
}

func (r *ProductRepository) FindAll() ([]*product.Product, error) {
    query := "SELECT product_id, name, description, price, stock, created_at, category, product_image FROM product"
    rows, err := r.db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var products []*product.Product
    for rows.Next() {
        var p product.Product
        if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Stock, &p.CreatedAt, &p.Category, &p.ProductImage); err != nil {
            return nil, err
        }
        products = append(products, &p)
    }

    return products, nil
}

func (r *ProductRepository) Update(p *product.Product) error {
    query := "UPDATE product SET name = $1, description = $2, price = $3, stock = $4, category = $5, product_image = $6 WHERE product_id = $7"
    _, err := r.db.Exec(query, p.Name, p.Description, p.Price, p.Stock, p.Category, p.ProductImage, p.ID)
    return err
}
