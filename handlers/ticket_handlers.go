package handlers

import (
	"landtick/models"
	"landtick/pkg/mysql"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetTicket(c echo.Context) error {
	var ticket []models.Ticket
	err := mysql.DB.Preload("users").Find(&ticket).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, ticket)
}
