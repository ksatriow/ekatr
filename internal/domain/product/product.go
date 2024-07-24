package product

import "time"

type ProductID int

type Product struct {
    ID          ProductID
    Name        string
    Description string
    Price       float64
    Stock       int
    CreatedAt   time.Time
}

func NewProduct(name, description string, price float64, stock int) *Product {
    return &Product{
        Name:        name,
        Description: description,
        Price:       price,
        Stock:       stock,
        CreatedAt:   time.Now(),
    }
}
