package mystructs

import (
	"time"
)

type Product struct {
	Product_Id       int       `json:"id"`
	Date_Created     time.Time `json:"date_created"`
	Product_name     string    `json:"name"`
	Serial_number    string    `json:"serial_number"`
	Product_quantity int       `json:"quantity"`
	Product_price    int       `json:"price"`
	Product_image    string    `json:"image"`
	Category_id      int       `json:"category"`
	description      string    `json:"description"`
}

type ProductCategories struct {
	Category_Id    int       `json:"id"`
	Category_name  string    `json:"name"`
	Date_created   time.Time `json:"date_created"`
	Category_image string    `json:"image"`
}
