package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rfauzi44/chat/models"
)

func GetAllUser(c echo.Context) error {
	result, err := models.GetAllUser()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)

}

func AddUser(c echo.Context) error {
	var user models.User
	err := c.Bind(&user)
	if err != nil {
		return err
	}
	result, err := models.AddUser(user.Username)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}
