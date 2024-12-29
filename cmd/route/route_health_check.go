package route

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func HealthCheck(app *echo.Echo) {
	app.GET("/", func(c echo.Context) error {
		type HealthCheck struct {
			Message string `json:"message"`
		}
		fmt.Println("health check passed")
		return c.JSON(
			http.StatusOK,
			&HealthCheck{
				Message: "success",
			},
		)
	})
}
