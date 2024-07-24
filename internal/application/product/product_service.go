package product

import "ekatr/internal/domain/product"

type ProductRepository interface {
    Save(product *product.Product) error
    FindByID(id int) (*product.Product, error)
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
