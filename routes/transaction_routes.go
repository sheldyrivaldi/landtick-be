package routes

import (
	"landtick/handlers"
	"landtick/pkg/middlewares"
	"landtick/pkg/mysql"
	"landtick/repositories"

	"github.com/labstack/echo/v4"
)

func TransactionRoutes(e *echo.Group) {
	transactionRepository := repositories.RepositoryStation(mysql.DB)
	h := handlers.HandlerTransaction(transactionRepository)

	e.GET("/transactions", middlewares.Auth(h.FindTransactions))
	e.GET("/transaction/:id", middlewares.Auth(h.GetTransaction))
	e.POST("/transaction", middlewares.Auth(middlewares.UploadFile(h.CreateTransaction)))

}
