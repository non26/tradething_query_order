package main

import (
	"context"
	"fmt"
	"tradethingqueryorder/cmd/route"
	"tradethingqueryorder/config"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"
	"github.com/labstack/echo/v4"
)

var echoLambda *echoadapter.EchoLambda
var _config *config.Config

func init() {
	var err error
	_config, err = config.ReadAWSAppConfig()
	if err != nil {
		fmt.Println("error read config", err)
		panic(err)
	}
	app_echo := echo.New()
	route.UpdateAWSAppConfig(app_echo, _config)
	route.Compose(app_echo, _config)
	echoLambda = echoadapter.New(app_echo)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return echoLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
