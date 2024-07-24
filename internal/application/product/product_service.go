package product

import "ekatr/internal/domain/product"

type ProductRepository interface {
    Save(product *product.Product) error
}

type ProductService struct {
    repo ProductRepository
}

func NewProductService(repo ProductRepository) *ProductService {
    return &ProductService{repo: repo}
}

func (s *ProductService) CreateProduct(name, description string, price float64, stock int, category string) (*product.Product, error) {
    if category == "" {
        category = "Uncategorized" // Or any default value you prefer
    }
    product := product.NewProduct(name, description, price, stock, category)
    if err := s.repo.Save(product); err != nil {
        return nil, err
    }
    return product, nil
}
