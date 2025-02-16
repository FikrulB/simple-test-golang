package repositories

import (
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
	"golang-test/entities"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(product *entities.Product) error
	FindAll(map[string]interface{}, string, string) ([]entities.Product, error)
	FindByID(uint) (*entities.Product, error)
	FindByName(string) (*entities.Product, error)
	FindCategoryByID(uint, *entities.ProductCategory) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) Create(product *entities.Product) error {
	if err := r.db.Create(&product).Error; err != nil {
		return handlePgErrorProduct(err, product)
	}
	return nil
}

func (r *productRepository) FindAll(filters map[string]interface{}, sortBy string, sortOrder string) ([]entities.Product, error) {
	var products []entities.Product
	query := r.db

	if categoryID, exists := filters["category_id"]; exists {
		query = query.Where("products.category_id = ?", categoryID.(int))
	}

	if name, exists := filters["name"]; exists {
		query = query.Where("products.name LIKE ?", "%"+name.(string)+"%")
	}

	if productID, exists := filters["id"]; exists {
		query = query.Where("products.id = ?", productID)
	}

	// Sorting
	if sortBy != "" {
		order := "ASC"
		if sortOrder == "desc" {
			order = "DESC"
		}
		query = query.Order(sortBy + " " + order)
	}

	err := query.Preload("Category").Find(&products).Error
	return products, err
}

func (r *productRepository) FindByID(id uint) (*entities.Product, error) {
	var product entities.Product
	err := r.db.First(&product, id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) FindByName(name string) (*entities.Product, error) {
	var product entities.Product
	err := r.db.Where("name = ?", name).First(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) FindCategoryByID(id uint, category *entities.ProductCategory) error {
	return r.db.First(category, id).Error
}

func handlePgErrorProduct(err error, product *entities.Product) error {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == "23505" {
		switch pgErr.ConstraintName {
		case "uni_products_code":
			return fmt.Errorf("produk dengan kode '%s' sudah ada", product.Code)
		case "uni_products_name_category":
			return fmt.Errorf("produk dengan nama '%s' dan kategori ID %d sudah ada", product.Name, product.CategoryID)
		}
	}
	return err
}
