package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"main.go/databasehandler"
	"main.go/product"
	"main.go/user"
	"main.go/utils"
)

func main() {
	databasehandler.InitializeDatabase(utils.GoDotEnvVariable("DATABASEURL"))
	databasehandler.Migrate()
	//gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	public := r.Group("/api")
	public.POST("/register", user.Register)

	r.POST("/api/products", product.AddProduct())
	r.POST("/api/categories", product.AddProductCategories())

	r.GET("/api/products", product.GetProducts())
	r.GET("/api/categories", product.GetProductsCategories())

	r.Run(":8005")
	fmt.Println("Server started on port 8005")

}
