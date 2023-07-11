package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mhdianrush/go-gin-product-api/config"
	"github.com/mhdianrush/go-gin-product-api/controllers"
	"github.com/sirupsen/logrus"
)

func main() {
	config.ConnectDB()
	r := gin.Default()

	r.GET("/api/products", controllers.Index)
	r.GET("/api/product/:id", controllers.Find)
	r.POST("/api/product", controllers.Create)
	r.PUT("/api/product/:id", controllers.Update)
	r.DELETE("/api/product", controllers.Delete)

	r.Run()

	logger := logrus.New()
	file, err := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	logger.SetOutput(file)

	logger.Println("Server Running on Port 8080")

	http.ListenAndServe(":8080", r)
}
