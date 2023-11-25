package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/xynurm/golang-ws-chat/routes"
)

func main() {
	errEnv := godotenv.Load("../.env")
	if errEnv != nil {
		panic("Failed to load environment")
	}

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Init routes
	routes.RouteInit(e.Group("/api/v1"))

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	// Start server
	fmt.Println("Server starting at :5000")
	e.Logger.Fatal(e.Start(":" + port))
}
