package main

import (
	"go-api/controller"
	"go-api/database"
	"go-api/repository"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, err := database.GetConnection()
	if err != nil {
		panic(err)
	}

	ProductRepository := repository.NewProductRepository(dbConnection)

	ProductUsecase := usecase.NewProductUsecase(ProductRepository)

	ProductController := controller.NewProductController(ProductUsecase)

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	server.GET("/products", ProductController.GetProducts)
	server.GET("/product/:id", ProductController.GetProductById)

	server.POST("/product", ProductController.CreateProduct)
	
	server.DELETE("/product/:id", ProductController.DeleteProduct)

	server.PUT("/product/:id", ProductController.UpdateProduct)
	
	server.Run(":8080")
}