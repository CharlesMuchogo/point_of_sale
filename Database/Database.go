package Database

import (
	"database/sql"
	"fmt"

	"os"

	"github.com/joho/godotenv"
)

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

func GetData() (string, string, string) {

	mainFoodQuery := "SELECT meal_name FROM main_meal ORDER BY RANDOM() LIMIT 1;"
	stewQuery := "SELECT stew FROM stew ORDER BY RANDOM() LIMIT 1;"
	mainMeal := fetchFromDb(mainFoodQuery)
	stew := fetchFromDb(stewQuery)
	imageQueryParameter := fmt.Sprintf("%v and %v", mainMeal, stew)

	imageQuery := fmt.Sprintf("SELECT food_image FROM food_image where food_name = '%v' ORDER BY RANDOM() LIMIT 1;", imageQueryParameter)
	image := fetchFromDb(imageQuery)

	return mainMeal, stew, image
}

func fetchFromDb(query string) string {

	var data string
	fmt.Printf("Query == %s \n", query)
	rows, err := DbConnect().Query(query)
	defer DbConnect().Close()

	CheckError(err)

	for rows.Next() {
		err = rows.Scan(&data)

		CheckError(err)

	}

	return data
}

func CheckError(err error) {
	if err != nil {
		panic(err.Error())
	}
}
