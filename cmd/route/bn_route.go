package route

import (
	"net/http"
	"tradethingqueryorder/config"

	"tradethingqueryorder/app/bn/handler"
	svc "tradethingqueryorder/app/bn/service"

	bntradesvc "tradethingqueryorder/app/bn/bn_service"

	bnclient "github.com/non26/tradepkg/pkg/bn/binance_client"
	bntransport "github.com/non26/tradepkg/pkg/bn/binance_transport"
	bndynamodb "github.com/non26/tradepkg/pkg/bn/dynamodb_repository"

	"github.com/labstack/echo/v4"
)

func BnRoute(app_echo *echo.Echo, config *config.Config) {
	serviceName := "BNF"
	bnroute := app_echo.Group("/" + serviceName)

	dynamodbconfig := bndynamodb.NewDynamodbConfig()
	dynamodbendpoint := bndynamodb.NewEndPointResolver(config.Dynamodb.Region, config.Dynamodb.Endpoint)
	dynamodbcredential := bndynamodb.NewCredential(config.Dynamodb.Ak, config.Dynamodb.Sk)
	dynamodbclient := bndynamodb.DynamoDB(dynamodbendpoint, dynamodbcredential, dynamodbconfig.LoadConfig()).NewLocal()
	svcrepository := bndynamodb.NewDynamoDBRepository(dynamodbclient)

	httptransport := bntransport.NewBinanceTransport(&http.Transport{})
	httpclient := bnclient.NewBinanceSerivceHttpClient()

	bnsvc := bntradesvc.NewBinanceFutureTradeService(
		config.Bn.BaseURL,
		config.Bn.EndPoint.PositionInformation,
		config.BnCredentials.APIKey,
		config.BnCredentials.SecretKey,
		httptransport,
		httpclient,
		"BNF",
	)

	service := svc.NewService(
		config.BnCredentials.APIKey,
		config.BnCredentials.SecretKey,
		config.BnCredentials.APIKey,
		config.BnCredentials.SecretKey,
		svcrepository,
		bnsvc,
	)

	handler := handler.NewQueryOrderHandler(service)
	bnroute.POST("/get-positions", handler.Handler)
}