package controller

import (
	"app/config"
	"app/model"
	"app/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	var users []model.User

	err := config.DB.Find(&users).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Internal Server Error"))
	}

	if len(users) == 0 {
		return c.JSON(http.StatusNotFound, utils.ErrorResponse("Not Found"))
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse("Success Get Data", users))
}
