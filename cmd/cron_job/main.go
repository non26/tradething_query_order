package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
	"tradethingqueryorder/config"

	svc "tradethingqueryorder/app/bn/service"

	"github.com/aws/aws-lambda-go/lambda"
	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"

	bntradesvc "tradethingqueryorder/app/bn/bn_service"

	svcrequest "tradethingqueryorder/app/bn/service_request"

	bnclient "github.com/non26/tradepkg/pkg/bn/binance_client"
	bntransport "github.com/non26/tradepkg/pkg/bn/binance_transport"
	bndynamodb "github.com/non26/tradepkg/pkg/bn/dynamodb_repository"
)

var echoLambda *echoadapter.EchoLambda
var _config *config.Config

func ConJob() error {
	var err error
	_config, err = config.ReadAWSAppConfig()
	if err != nil {
		fmt.Println("error read config", err)
		return err
	} else {
		fmt.Println("config load success ", _config.Bn.BaseURL)
	}

	dynamodbconfig := bndynamodb.NewDynamodbConfig()
	dynamodbendpoint := bndynamodb.NewEndPointResolver(_config.Dynamodb.Region, _config.Dynamodb.Endpoint)
	dynamodbcredential := bndynamodb.NewCredential(_config.Dynamodb.Ak, _config.Dynamodb.Sk)
	dynamodbclient := bndynamodb.DynamoDB(dynamodbendpoint, dynamodbcredential, dynamodbconfig.LoadConfig()).NewPrd()
	svcrepository := bndynamodb.NewDynamoDBRepository(dynamodbclient)

	httptransport := bntransport.NewBinanceTransport(&http.Transport{})
	httpclient := bnclient.NewBinanceSerivceHttpClient()

	bnsvc := bntradesvc.NewBinanceFutureTradeService(
		_config.Bn.BaseURL,
		_config.Bn.EndPoint.PositionInformation,
		_config.BnCredentials.APIKey,
		_config.BnCredentials.SecretKey,
		httptransport,
		httpclient,
		"BNF",
	)

	service := svc.NewService(
		_config.BnCredentials.APIKey,
		_config.BnCredentials.SecretKey,
		_config.BnCredentials.APIKey,
		_config.BnCredentials.SecretKey,
		svcrepository,
		bnsvc,
	)
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	fmt.Println("start query bn position")
	bnres, err := service.QueryOrder(ctx, &svcrequest.QueryOrder{
		Symbol: "",
	})
	if err != nil {
		fmt.Println("error query bn position", err)
	} else {
		fmt.Println("success query bn position", bnres)
	}
	fmt.Println("end query bn position")

	return nil
}

func main() {
	fmt.Println("start cron job")
	lambda.Start(ConJob)
	fmt.Println("end cron job")
}
