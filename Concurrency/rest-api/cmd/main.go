package main

import (
	"rest-api/config"
	"rest-api/handler"
	"rest-api/repository"
	"rest-api/routes"
	"rest-api/service"

	"github.com/gin-gonic/gin"
)

func main() {

	config.ConnectDB()

	router := gin.Default()

	productRepo := &repository.ProductRepository{
		DB: config.DB,
	}

	productService := &service.ProductService{
		Repo: productRepo,
	}

	productHandler := &handler.ProductHandler{
		Service: productService,
	}

	routes.SetUpRoutes(router, productHandler)

	router.Run(":8080")
}
