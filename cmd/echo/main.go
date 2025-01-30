package main

import (
	"fmt"
	"tradethingqueryorder/cmd/route"
	"tradethingqueryorder/config"

	"github.com/labstack/echo/v4"
)

func main() {
	config, err := config.ReadConfig()
	if err != nil {
		fmt.Println("error read config", err)
		return
	}

	app_echo := echo.New()
	route.Compose(app_echo, config)
	app_echo.Start(fmt.Sprintf(":%v", "8081"))
}
