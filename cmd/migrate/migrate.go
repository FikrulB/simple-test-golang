package migrate

import (
	"golang-test/configs"
	"golang-test/entities"
	"log"
)

func Run() {
	db := configs.NewDB(configs.NewEnv())

	if err := db.AutoMigrate(&entities.Product{}, &entities.ProductCategory{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
}
