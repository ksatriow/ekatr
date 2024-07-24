package http

import (
    "encoding/json"
    "net/http"
    "strconv"
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

    prod, err := h.service.CreateProduct(dto.Name, dto.Description, dto.Price, dto.Stock, dto.Category, dto.ProductImage)
    if err != nil {
        logger.ErrorLogger.Printf("Error creating product: %v", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(prod)
}

func (h *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
    idStr := r.URL.Query().Get("id")
    if idStr == "" {
        http.Error(w, "missing product ID", http.StatusBadRequest)
        return
    }

    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "invalid product ID", http.StatusBadRequest)
        return
    }

    product, err := h.service.GetProductByID(id)
    if err != nil {
        logger.ErrorLogger.Printf("Error getting product by ID: %v", err)
        if err.Error() == "product not found" {
            http.Error(w, err.Error(), http.StatusNotFound)
        } else {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
        return
    }

    if product == nil {
        http.Error(w, "product not found", http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
    products, err := h.service.GetAllProducts()
    if err != nil {
        logger.ErrorLogger.Printf("Error getting all products: %v", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(products)
}
