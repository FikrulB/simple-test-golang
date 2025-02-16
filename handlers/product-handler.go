package handlers

import (
	"github.com/labstack/echo/v4"
	"golang-test/domain"
	"golang-test/services"
	"golang-test/utils"
	"net/http"
	"strconv"
)

type ProductHandler interface {
	CreateProduct(echo.Context) error
	GetProducts(echo.Context) error
}

type productHandler struct {
	productService services.ProductService
}

func NewProductHandler(prs services.ProductService) ProductHandler {
	return &productHandler{productService: prs}
}

func (h *productHandler) CreateProduct(ctx echo.Context) error {
	var payload domain.CreateProductReq

	if err := ctx.Bind(&payload); err != nil {
		return utils.JSONResponse(ctx, http.StatusBadRequest, "Invalid request", nil, err)
	}

	if err := ctx.Validate(&payload); err != nil {
		return utils.JSONResponse(ctx, http.StatusBadRequest, "Validation failed", nil, err)
	}

	err := h.productService.CreateProduct(payload)
	if err != nil {
		return utils.JSONResponse(ctx, http.StatusInternalServerError, "Failed to add product", nil, err)
	}

	return utils.JSONResponse(ctx, http.StatusCreated, "Product added successfully", nil, nil)
}

func (h *productHandler) GetProducts(ctx echo.Context) error {
	var query domain.GetProductsReq

	if err := ctx.Bind(&query); err != nil {
		return utils.JSONResponse(ctx, http.StatusBadRequest, "Invalid query parameters", nil, err)
	}

	if err := ctx.Validate(&query); err != nil {
		return utils.JSONResponse(ctx, http.StatusBadRequest, "Validation failed", nil, err)
	}

	filters := map[string]interface{}{}
	if query.Category != "" {
		categoryID, err := strconv.Atoi(query.Category)
		if err == nil {
			filters["category_id"] = categoryID
		}
	}

	if query.Name != "" {
		filters["name"] = query.Name
	}

	if query.ID != "" {
		productID, err := strconv.Atoi(query.ID)
		if err == nil {
			filters["id"] = productID
		}
	}

	products, err := h.productService.GetProducts(filters, query.SortBy, query.SortOrder)
	if err != nil {
		return utils.JSONResponse(ctx, http.StatusInternalServerError, "Failed to get products", nil, err)
	}

	return utils.JSONResponse(ctx, http.StatusOK, "Products retrieved successfully", products, nil)
}
