package domain

import (
	"time"
)

type (
	CreateProductReq struct {
		Code       string  `json:"code" validate:"required"`
		Name       string  `json:"name" validate:"required"`
		CategoryID uint    `json:"category_id" validate:"required,gt=0"`
		Price      float64 `json:"price" validate:"required,gte=0"`
	}

	GetProductsReq struct {
		ID        string `query:"id"`
		Category  string `query:"category_id"`
		Name      string `query:"name"`
		SortBy    string `query:"sortBy" validate:"omitempty,oneof=name price created_at"`
		SortOrder string `query:"order" validate:"omitempty,oneof=asc desc"`
	}

	GetProductsRes struct {
		Code      string                `json:"code"`
		Name      string                `json:"name"`
		Category  CategoryGetProductRes `json:"category"`
		Price     float64               `json:"price"`
		CreatedAt time.Time             `json:"created_at"`
	}

	CategoryGetProductRes struct {
		ID       uint      `json:"id"`
		Name     string    `json:"name"`
		CreateAt time.Time `json:"create_at"`
	}
)
