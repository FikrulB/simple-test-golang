package services

import (
	"context"
	"errors"
	"fmt"
	"golang-test/domain"
	"golang-test/entities"
	"golang-test/repositories"
	"time"
)

type ProductService interface {
	GetProducts(map[string]interface{}, string, string) ([]domain.GetProductsRes, error)
	CreateProduct(domain.CreateProductReq) error
}

type productService struct {
	productRepo repositories.ProductRepository
	redisRepo   repositories.RedisRepository
}

func NewProductService(prr repositories.ProductRepository, rdr repositories.RedisRepository) ProductService {
	return &productService{productRepo: prr, redisRepo: rdr}
}

func (s *productService) GetProducts(filters map[string]interface{}, sortBy, sortOrder string) ([]domain.GetProductsRes, error) {
	ctx := context.Background()
	cacheKey := fmt.Sprintf("products:%v:%s:%s", filters, sortBy, sortOrder)
	fmt.Println("cacheKey => ", cacheKey)

	var cachedResponse []domain.GetProductsRes
	err := s.redisRepo.GetCache(ctx, cacheKey, &cachedResponse)
	if err == nil && len(cachedResponse) > 0 {
		return cachedResponse, nil
	}

	products, err := s.productRepo.FindAll(filters, sortBy, sortOrder)
	if err != nil {
		return nil, err
	}

	var response []domain.GetProductsRes
	for _, product := range products {
		response = append(response, domain.GetProductsRes{
			Code: product.Code,
			Name: product.Name,
			Category: domain.CategoryGetProductRes{
				ID:       product.Category.ID,
				Name:     product.Category.Name,
				CreateAt: product.Category.CreatedAt,
			},
			Price:     product.Price,
			CreatedAt: product.CreatedAt,
		})
	}

	err = s.redisRepo.SetCache(ctx, cacheKey, response, 10*time.Minute)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *productService) CreateProduct(payload domain.CreateProductReq) error {
	ctx := context.Background()
	if payload.Price < 0 {
		return errors.New("harga tidak boleh di bawah 0")
	}

	var category entities.ProductCategory
	if err := s.productRepo.FindCategoryByID(payload.CategoryID, &category); err != nil {
		return errors.New("kategori tidak ditemukan")
	}

	product := entities.Product{
		Code:       payload.Code,
		Name:       payload.Name,
		Price:      payload.Price,
		CategoryID: payload.CategoryID,
	}

	err := s.productRepo.Create(&product)
	if err != nil {
		return err
	}

	keys, err := s.redisRepo.GetKeys(ctx, "products:*")
	if err == nil {
		for _, key := range keys {
			_ = s.redisRepo.DeleteCache(ctx, key)
		}
	}

	return nil
}
