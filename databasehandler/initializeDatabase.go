package databasehandler

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"main.go/mystructs"
	"os"
)

var DB *gorm.DB

func InitializeDatabase(connectionString string) {
	var err error
	DB, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

}

func Migrate() {
	err := DB.AutoMigrate(&mystructs.Product{})
	err = DB.AutoMigrate(&mystructs.ProductCategories{})
	if err != nil {
		return
	}
	log.Println("Database Migration Completed!")
}

func GetPostgresConnectionString() string {

	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
}
