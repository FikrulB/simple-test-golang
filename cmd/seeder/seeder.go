package seeder

import (
	"fmt"
	"golang-test/configs"
	"golang-test/entities"
	"time"
)

func Run() {
	db := configs.NewDB(configs.NewEnv())

	products := []entities.ProductCategory{
		{Name: "Sayuran"},
		{Name: "Protein"},
		{Name: "Buah"},
		{Name: "Snack"},
	}

	for _, product := range products {
		var existing entities.ProductCategory

		if err := db.Where("name = ?", product.Name).First(&existing).Error; err != nil {
			product.CreatedAt = time.Now()
			db.Create(&product)
		}
	}

	fmt.Println("Seeding completed!")
}
