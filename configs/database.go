package configs

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func NewDB(env *Config) *gorm.DB {
	dbUrl := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", env.DBHost, env.DBUser, env.DBPass, env.DBName, env.DBPort)
	fmt.Println(dbUrl)
	database, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database connected and migrated successfully")
	db = database

	return db
}
