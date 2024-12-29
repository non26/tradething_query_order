package route

import (
	"net/http"
	"tradethingqueryorder/config"

	"github.com/labstack/echo/v4"
)

func UpdateAWSAppConfig(app *echo.Echo, _config *config.Config) {
	app.GET("/update-aws-config", func(c echo.Context) error {
		var err error
		_config, err = config.ReadAWSAppConfig()
		if err != nil {
			return err
		}
		type Res struct {
			Message string `json:"message"`
		}
		m := Res{}
		m.Message = "success"
		return c.JSON(http.StatusOK, &m)
	})
}
