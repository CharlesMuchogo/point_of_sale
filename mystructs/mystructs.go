package mystructs

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Product_name        string `json:"name"`
	Serial_number       string `json:"serial_number"`
	Product_quantity    int    `json:"quantity"`
	Product_price       int    `json:"price"`
	Product_image       string `json:"image"`
	Category_id         int    `json:"category"`
	Product_Description string `json:"description"`
}

type ProductCategories struct {
	gorm.Model
	Category_name  string `json:"name"`
	Category_image string `json:"image"`
}
