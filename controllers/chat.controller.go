package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rfauzi44/chat-api/models"
)

func AddChat(c echo.Context) error {
	var chat models.Chat
	err := c.Bind(&chat)
	if err != nil {
		return err
	}
	result, err := models.AddChat(chat.SenderID, chat.ReceiverID, chat.MessageText)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}

func GetChatWith(c echo.Context) error {
	var chat models.Chat
	err := c.Bind(&chat)
	if err != nil {
		return err
	}

	result, err := models.GetChatWith(chat.SenderID, chat.ReceiverID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)

}
