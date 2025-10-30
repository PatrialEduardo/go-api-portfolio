package main

import (
	"go-api/controller"
	"go-api/usecase"
	"github.com/gin-gonic/gin"
	"go-api/repository"
	"go-api/db"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	//camada de repository
	ProductRepository := repository.NewProductRepository(dbConnection)
	//camada de usecase
	ProductUseCase := usecase.NewProductUsecase(ProductRepository)
	//camada de controllers
	ProductController := controller.NewProductController(ProductUseCase)
	server.GET("/products", ProductController.GetProducts)
	server.POST("/product", ProductController.CreateProduct)
	server.GET("/product/:productId", ProductController.GetProductsById)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.Run(":8080")
}

