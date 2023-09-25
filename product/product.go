package product

import (
	"main.go/databasehandler"
	"main.go/mystructs"

	"github.com/gin-gonic/gin"
	"net/http"
)

func AddProductCategories() gin.HandlerFunc {

	return func(c *gin.Context) {
		var category mystructs.ProductCategories
		if err := c.ShouldBindJSON(&category); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			c.Abort()
			return
		}

		databasehandler.AddCategories(category)
		c.JSON(http.StatusCreated, gin.H{
			"message": "Product category added",
		})

	}

}

func AddProduct() gin.HandlerFunc {

	return func(c *gin.Context) {
		var product mystructs.Product
		if err := c.ShouldBindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			c.Abort()
			return
		}

		databasehandler.AddProduct(product)

		c.JSON(http.StatusCreated, gin.H{
			"message": "Product added successfully",
		})

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
			"Categories": productCategories,
		})

	}

}
