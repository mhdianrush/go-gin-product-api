package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mhdianrush/go-gin-product-api/config"
	"github.com/mhdianrush/go-gin-product-api/controllers"
	"github.com/sirupsen/logrus"
)

func main() {
	config.ConnectDB()
	routes := gin.Default()

	routes.GET("/api/products", controllers.Index)
	routes.GET("/api/product/:id", controllers.Find)
	routes.POST("/api/product", controllers.Create)
	routes.PUT("/api/product/:id", controllers.Update)
	routes.DELETE("/api/product", controllers.Delete)

	routes.Run()

	logger := logrus.New()
	
	file, err := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		logger.Printf("failed create log file %s", err.Error())
	}
	logger.SetOutput(file)

	if err := godotenv.Load(); err != nil {
		logger.Printf("failed load env file %s", err.Error())
	}

	server := http.Server{
		Addr:    ":" + os.Getenv("SERVER_PORT"),
		Handler: routes,
	}
	if err = server.ListenAndServe(); err != nil {
		logger.Printf("failed connect to server %s", err.Error())
	}

	logger.Printf("Server Running on Port %s", os.Getenv("SERVER_PORT"))
}
