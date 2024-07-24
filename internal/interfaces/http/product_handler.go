package http

import (
    "encoding/json"
    "net/http"
    "ekatr/internal/application/product"
    "ekatr/internal/logger"
)

type ProductHandler struct {
    service *product.ProductService
}

func NewProductHandler(service *product.ProductService) *ProductHandler {
    return &ProductHandler{service: service}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
    var dto product.CreateProductDTO
    if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
        logger.ErrorLogger.Printf("Error decoding request body: %v", err)
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Set default category if not provided
    if dto.Category == "" {
        dto.Category = "Uncategorized" // Or any default value you prefer
    }

    prod, err := h.service.CreateProduct(dto.Name, dto.Description, dto.Price, dto.Stock, dto.Category)
    if err != nil {
        logger.ErrorLogger.Printf("Error creating product: %v", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(prod)
}
