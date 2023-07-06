package routes

import (
	"landtick/handlers"

	"github.com/labstack/echo/v4"
)

func TicketRoutes(e *echo.Group) {
	e.GET("/tickets", handlers.GetTicket)
}
