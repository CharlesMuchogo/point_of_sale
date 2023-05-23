package product

import (
	"fmt"
	"main.go/databasehandler"
	"main.go/mystructs"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddProductCategories() gin.HandlerFunc {

	return func(c *gin.Context) {

		category_name := c.DefaultPostForm("category_name", "unknown")

		if category_name != "unknown" {
			fmt.Printf("category name to add is  - %v", category_name)

			status := databasehandler.AddCategories(category_name)

			message := "Error adding category"
			if status < 1 {
				message = "Error adding category"
			} else {
				message = "category added"
			}

			c.JSON(http.StatusOK, gin.H{
				"message": message,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": "From data not complete",
			})
		}

	}

}

func AddProduct() gin.HandlerFunc {

	return func(c *gin.Context) {

		product_name := c.DefaultPostForm("product_name", "unknown")
		serial_number := c.DefaultPostForm("serial_number", "unknown")
		product_quantity, _ := strconv.Atoi(c.DefaultPostForm("product_quantity", "unknown"))
		product_price, _ := strconv.Atoi(c.DefaultPostForm("product_price", "unknown"))
		product_image := c.DefaultPostForm("product_image", "unknown")
		category_id, err := strconv.Atoi(c.DefaultPostForm("category_id", "4"))
		product_description := c.DefaultPostForm("product_description", "unknown")

		if err != nil {
			fmt.Printf("Error during conversion %s", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "product data not complete",
			})
			return
		}
		if serial_number != "unknown" || product_name != "unknown" {
			var product_to_add mystructs.Product
			product_to_add.Product_name = product_name
			product_to_add.Serial_number = serial_number
			product_to_add.Product_quantity = product_quantity
			product_to_add.Product_price = product_price
			product_to_add.Product_image = product_image
			product_to_add.Category_id = category_id
			product_to_add.Product_Description = product_description

			status := databasehandler.AddProduct(product_to_add)
			var message string

			if status < 1 {
				message = "Error adding product"
			} else {
				message = "product added successfully"
			}

			c.JSON(http.StatusOK, gin.H{
				"message": message,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "product data not complete",
			})
		}

	}

}

func GetProducts() gin.HandlerFunc {

	return func(c *gin.Context) {

		var products = databasehandler.GetProducts()

		c.JSON(http.StatusOK, gin.H{
			"products": products,
			"message":  "fetched products successfully",
		})

	}

}

func GetProductsCategories() gin.HandlerFunc {
	fmt.Printf("reaching here")
	return func(c *gin.Context) {

		var productCategories = databasehandler.GetProductCategories()

		c.JSON(http.StatusOK, gin.H{
			"product_categories": productCategories,
		})

	}

}
