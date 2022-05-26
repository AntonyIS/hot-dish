package postgres

import (
	"fmt"
	"log"
	"os"

	"github.com/AntonyIS/go-foods/internal/core/domain"
	"github.com/AntonyIS/go-foods/internal/core/ports"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

type PostgresRepository struct {
	DB        *gorm.DB
	TableName string
}

func PostgresDB() *gorm.DB {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")

	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", host, port, user, dbname, password))

	if err != nil {
		log.Fatal(err)
	}

	db.LogMode(false)
	db.AutoMigrate(domain.Courier{})
	db.AutoMigrate(domain.Customer{})
	db.AutoMigrate(domain.Restaurant{})

	return db

}

func NewRepostory() ports.ItemRepository {
	return PostgresRepository{
		DB: PostgresDB(),
	}
}
