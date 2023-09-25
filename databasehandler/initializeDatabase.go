package databasehandler

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"main.go/mystructs"
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
