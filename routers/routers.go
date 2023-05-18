package routers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rfauzi44/chat/controllers"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		link := "<a href='https://documenter.getpostman.com/view/25042327/2s93eVYaBS'>here</a>"
		response := "Hello World! This is rfauzi44/chat API. You can check Postman Documentation " + link
		return c.HTML(http.StatusOK, response)
	})

	//User
	e.GET("/user", controllers.GetAllUser)
	e.POST("/user", controllers.AddUser)

	//Chat
	e.POST("/chat", controllers.AddChat)
	e.GET("/chat", controllers.GetChatWith)

	return e
}
