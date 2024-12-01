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

	//Camada de Repository
	ProductRepository := repository.NewProductRepository(dbConnection)

	//Camada de UseCase
	ProductUsecase := usecase.NewProductUsecase(ProductRepository)

	//Camada de Controllers
	ProductController := controller.NewProductController(ProductUsecase)

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", ProductController.GetProducts)

	server.Run(":8080")
}