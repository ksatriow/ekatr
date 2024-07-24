package product

import "time"

type ProductID int

type Product struct {
    ID           ProductID
    Name         string
    Description  string
    Price        float64
    Stock        int
    Category     string
    ProductImage string
    CreatedAt    time.Time
}

func NewProduct(name, description string, price float64, stock int, category, productImage string) *Product {
    return &Product{
        Name:         name,
        Description:  description,
        Price:        price,
        Stock:        stock,
        Category:     category,
        ProductImage: productImage,
        CreatedAt:    time.Now(),
    }
}
