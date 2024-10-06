package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/nuwasdzarrin/golang-gin-jwt-auth-crud/api/router"
	"github.com/nuwasdzarrin/golang-gin-jwt-auth-crud/config"
	"github.com/nuwasdzarrin/golang-gin-jwt-auth-crud/db/initializers"
)

func init() {
	config.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	fmt.Println("Hello auth")

	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "" {
		ginMode = "debug"
	}
	gin.SetMode(ginMode)

	r := gin.Default()
	router.GetRoute(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default to 8080 if PORT is not set
	}

	r.Run(":" + port)
}
