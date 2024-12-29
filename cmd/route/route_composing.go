package route

import (
	"tradethingqueryorder/config"

	"github.com/labstack/echo/v4"
)

func Compose(app_echo *echo.Echo, config *config.Config) {
	HealthCheck(app_echo)
	BnRoute(app_echo, config)
}
