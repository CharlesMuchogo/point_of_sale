package mystructs

import (
	"time"
)

type Product struct {
	Product_Id       int       `json:"product_id"`
	Date_Created     time.Time `json:"date_created"`
	Product_name     string    `json:"product_name"`
	Serial_number    string    `json:"serial_number"`
	Product_quantity int       `json:"quantity"`
	Product_price    int       `json:"price"`
	Product_image    string    `json:"image"`
	Category_id      int       `json:"category_id"`
}

type ProductCategories struct {
	Category_Id   int
	Category_name string
	Date_created  time.Time
}
