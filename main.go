package main

import (
	"landtick/database"
	"landtick/pkg/mysql"
	"landtick/routes"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mysql.DatabaseInit()
	database.RunMigration()

	e := echo.New()

	e.GET("/api/v1", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	routes.RouteInit(e.Group("/api/v1"))

	e.Logger.Fatal(e.Start(":5000"))
}
