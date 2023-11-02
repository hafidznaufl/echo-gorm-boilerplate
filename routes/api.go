package routes

import (
	"app/controller"
	"app/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {

	e := echo.New()

	e.Use(middleware.NotFoundHandler)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to RESTful API Services")
	})

	e.GET("/users", controller.GetAllUsers)
	e.GET("/users/:id", controller.GetUserById)
	e.POST("/users/register", controller.Register)
	e.POST("/users/login", controller.Login)
	e.PUT("/users/:id", controller.Update)
	e.DELETE("/users/:id", controller.Delete)

	return e

}
