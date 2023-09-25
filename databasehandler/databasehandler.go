package databasehandler

import (
	"database/sql"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"main.go/mystructs"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null;" json:"password"`
}

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	CheckError(err)

	return os.Getenv(key)
}

func DbConnect() *sql.DB {
	dsn := goDotEnvVariable("DATABASEURL")
	db, err := sql.Open("postgres", dsn)
	CheckError(err)
	return db
}

func AddProduct(
	product mystructs.Product,
) {
	DB.Create(&product)
}

func AddCategories(category mystructs.ProductCategories) {
	DB.Create(&category)
}

func GetProducts() []mystructs.Product {
	var products []mystructs.Product
	DB.Find(&products)
	return products
}

func GetProductCategories() []mystructs.ProductCategories {
	var categories []mystructs.ProductCategories
	DB.Find(&categories)
	return categories
}

func CheckError(err error) {
	if err != nil {
		panic(err.Error())
	}
}
