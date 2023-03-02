package main

import (
	"fmt"
	"main.go/product"
	"main.go/user"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	public := r.Group("/api")
	public.POST("/register", user.Register)

	r.POST("/api/product", product.AddProduct())
	r.POST("/api/category", product.AddProductCategories())

	r.GET("/api/product", product.GetProducts())
	r.GET("/api/category", product.GetProductsCategories())

	r.Run(":8080")
	fmt.Println("Server started on port 8080")

}
