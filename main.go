package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	r := gin.Default()

	logger := logrus.New()
	file, err := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	logger.SetOutput(file)

	logger.Println("Server Running on Port 8080")

	http.ListenAndServe(":8080", r)
}
