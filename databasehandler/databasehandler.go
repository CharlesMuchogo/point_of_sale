package databasehandler

import (
	"main.go/mystructs"
)

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
