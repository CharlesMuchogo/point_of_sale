package databasehandler

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"main.go/mystructs"
	"os"
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

func AddProduct(product_name string,
	serial_number string,
	product_quantity int,
	product_price int,
	product_image string,
	category_id int) int {

	mainFoodQuery := "INSERT INTO product(product_name, serial_number, product_quantity , product_price, product_image, category_id ) VALUES($1, $2, $3, $4, $5, $6)"

	insertMainMeal, err := DbConnect().Exec(mainFoodQuery, product_name,
		serial_number,
		product_quantity,
		product_price,
		product_image,
		category_id)
	defer DbConnect().Close()
	CheckError(err)

	mainMealRowsaffected, err := insertMainMeal.RowsAffected()
	CheckError(err)
	fmt.Printf("Inserted data %v row affected \n", mainMealRowsaffected)

	return int(mainMealRowsaffected)
}

func AddCategories(category_name string) int {
	categoryUrl := "INSERT INTO category(name) VALUES($1)"
	insertCartegory, err := DbConnect().Exec(categoryUrl, category_name)
	defer DbConnect().Close()
	CheckError(err)

	rowsaffected, err := insertCartegory.RowsAffected()
	CheckError(err)
	fmt.Printf("Inserted data %v row affected \n", rowsaffected)

	return int(rowsaffected)

}

func GetProducts() []mystructs.Product {
	fmt.Printf("reaching here")
	query := "SELECT product_id, date_created, product_name, serial_number, product_quantity, product_price, product_image, category_id FROM product;"
	rows, err := DbConnect().Query(query)
	defer DbConnect().Close()
	CheckError(err)

	var current_product mystructs.Product

	var productsSlice []mystructs.Product

	for rows.Next() {

		err = rows.Scan(&current_product.Product_Id, &current_product.Date_Created, &current_product.Product_name, &current_product.Serial_number, &current_product.Product_quantity, &current_product.Product_price, &current_product.Product_image, &current_product.Product_Id)
		productsSlice = append(productsSlice, current_product)
		CheckError(err)

	}

	return productsSlice
}

func GetProductCategories() []mystructs.ProductCategories {

	query := "SELECT category_id, name , date_created  FROM category;"
	rows, err := DbConnect().Query(query)
	defer DbConnect().Close()
	CheckError(err)

	var current_product_category mystructs.ProductCategories

	var productCategorySlice []mystructs.ProductCategories

	for rows.Next() {
		fmt.Printf("reaching here")
		err = rows.Scan(&current_product_category.Category_Id, &current_product_category.Category_name, &current_product_category.Date_created)
		productCategorySlice = append(productCategorySlice, current_product_category)
		CheckError(err)

	}

	return productCategorySlice
}

func CheckError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func SaveUser(username string, password []byte) (int, error) {
	categoryUrl := "INSERT INTO users(username, password) VALUES($1, $2)"

	insertCartegory, err := DbConnect().Exec(categoryUrl, username, password)
	defer DbConnect().Close()
	if err != nil {
		return 0, err
	}

	rowsaffected, err := insertCartegory.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(rowsaffected), nil

}
