package product

import (
	"fmt"
	"main.go/databasehandler"

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
				"Error": "Form data not complete",
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
		category_id, err := strconv.Atoi(c.DefaultPostForm("category_id", "unknown"))

		if err != nil {
			fmt.Println("Error during conversion")
			return
		}
		if serial_number != "unknown" || product_name != "unknown" {
			fmt.Printf("food to add is  - %v", product_price)
			fmt.Printf("image to add is  - %v", product_quantity)

			status := databasehandler.AddProduct(product_name, serial_number, product_quantity, product_price, product_image, category_id)
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
		})

	}

}

func GetProductsCategories() gin.HandlerFunc {

	return func(c *gin.Context) {

		var productCategories = databasehandler.GetProductCategories()

		c.JSON(http.StatusOK, gin.H{
			"product_categories": productCategories,
		})

	}

}
