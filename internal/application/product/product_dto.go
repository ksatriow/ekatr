package product

import (
    "errors"
)

type CreateProductDTO struct {
    Name         string  `json:"name"`
    Description  string  `json:"description"`
    Price        float64 `json:"price"`
    Stock        int     `json:"stock"`
    Category     string  `json:"category"`
    ProductImage string  `json:"product_image"`
}

func (dto *CreateProductDTO) Validate() error {
    if dto.Category == "" {
        return errors.New("category is required")
    }
    return nil
}
