package product

import (
    "ekatr/internal/domain/product"
    "errors"
)

type ProductRepository interface {
    Save(product *product.Product) error
    FindByID(id int) (*product.Product, error)
    FindAll() ([]*product.Product, error)
    Update(product *product.Product) error
}

type ProductService struct {
    repo ProductRepository
}

func NewProductService(repo ProductRepository) *ProductService {
    return &ProductService{repo: repo}
}

func (s *ProductService) CreateProduct(name, description string, price float64, stock int, category, productImage string) (*product.Product, error) {
    if category == "" {
        category = "Uncategorized" // Or any default value you prefer
    }
    product := product.NewProduct(name, description, price, stock, category, productImage)
    if err := s.repo.Save(product); err != nil {
        return nil, err
    }
    return product, nil
}

func (s *ProductService) GetProductByID(id int) (*product.Product, error) {
    return s.repo.FindByID(id)
}

func (s *ProductService) GetAllProducts() ([]*product.Product, error) {
    return s.repo.FindAll()
}

func (s *ProductService) UpdateProduct(id int, name, description string, price float64, stock int, category, productImage string) (*product.Product, error) {
    product, err := s.repo.FindByID(id)
    if err != nil {
        return nil, err
    }
    if product == nil {
        return nil, errors.New("product not found")
    }

    // Update fields only if they are provided
    if name != "" {
        product.Name = name
    }
    if description != "" {
        product.Description = description
    }
    if price != 0 {
        product.Price = price
    }
    if stock != 0 {
        product.Stock = stock
    }
    if category != "" {
        product.Category = category
    }
    if productImage != "" {
        product.ProductImage = productImage
    }

    if err := s.repo.Update(product); err != nil {
        return nil, err
    }
    return product, nil
}
