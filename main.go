package main

import (
	"fmt"

	"main/Database"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {

	r := gin.Default()

	r.POST("/api/product", addProduct())
	r.POST("/api/category", addCategories())

	r.GET("/api/product", get())

	r.GET("/bloc", bloc())

	r.Run()
	fmt.Println("Server started on port 8080")

}

// type Product struct {
// 	product_name     string
// 	serial_number    string
// 	product_quantity int
// 	product_price    int
// 	product_image    string
// 	category_id      int
// }

func addCategories() gin.HandlerFunc {

	return func(c *gin.Context) {

		category_name := c.DefaultPostForm("category_name", "unknown")

		if category_name != "unknown" {
			fmt.Printf("category name to add is  - %v", category_name)

			status := Database.AddCategories(category_name)

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

func addProduct() gin.HandlerFunc {

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

			status := Database.AddProduct(product_name, serial_number, product_quantity, product_price, product_image, category_id)
			var message string

			if status < 1 {
				message = "Error adding food"
			} else {
				message = "Food added"
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

func get() gin.HandlerFunc {

	return func(c *gin.Context) {

		foodname, stew, image := Database.GetData()

		c.JSON(http.StatusOK, gin.H{
			"main food": foodname,
			"stew":      stew,
			"image":     image,
		})
	}

}

func bloc() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "200"})
	}
}
