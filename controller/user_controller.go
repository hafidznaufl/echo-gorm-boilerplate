package controller

import (
	"app/config"
	"app/model"
	"app/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	var users []model.User

	err := config.DB.Find(&users).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve user"))
	}

	if len(users) == 0 {
		return c.JSON(http.StatusNotFound, utils.ErrorResponse("Empty data"))
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse("User data successfully retrieved", users))
}

func Show(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid ID"))
	}

	var user model.User
	result := config.DB.First(&user, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve user"))
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse("User data successfully retrieved", user))
}


func Store(c echo.Context) error {
	var user model.User

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body"))
	}

	if err := config.DB.Create(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to store user data"))
	}

	return c.JSON(http.StatusCreated, utils.SuccessResponse("Success Created Data", user))
}

func Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid ID"))
	}

	var updatedUser model.User

	if err := c.Bind(&updatedUser); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body"))
	}

	var existingUser model.User
	result := config.DB.First(&existingUser, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve user"))
	}

	// Memperbarui data pengguna dengan data yang baru
	config.DB.Model(&existingUser).Updates(updatedUser)

	return c.JSON(http.StatusOK, utils.SuccessResponse("User data successfully updated", existingUser))
}

func Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid ID"))
	}

	var existingUser model.User
	result := config.DB.First(&existingUser, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve user"))
	}

	// Menghapus data pengguna dari database
	config.DB.Delete(&existingUser)

	return c.JSON(http.StatusOK, utils.SuccessResponse("User data successfully deleted", nil))
}


