package main

import (
	"go-api/controller"
	"go-api/usecase"
	"github.com/gin-gonic/gin"
	"go-api/repository"
	"go-api/db"
    "os"
    "log"

)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

		sqlBytes, err := os.ReadFile("db/init.sql")
	if err != nil {
		log.Println("⚠️  init.sql não encontrado, ignorando migração")
	} else {
		_, err := dbConnection.Exec(string(sqlBytes))
		if err != nil {
			log.Println("⚠️  Erro ao executar init.sql:", err)
		} else {
			log.Println("✅ Migração SQL executada com sucesso")
		}
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

