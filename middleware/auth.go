package middleware

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func NotFoundHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)

		if err != nil {
			if he, ok := err.(*echo.HTTPError); ok {
				if he.Code == http.StatusNotFound {
					// Penanganan khusus jika endpoint tidak ditemukan
					errorMessage := "Endpoint tidak ditemukan."
					return c.JSON(http.StatusNotFound, map[string]interface{}{
						"message": errorMessage,
					})
				}
			}

			// Menangani kesalahan lainnya
			fmt.Println("Terjadi kesalahan:", err)
		}

		return err
	}
}
