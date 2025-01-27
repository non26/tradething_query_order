package service

import (
	"context"
	handlerresponse "tradethingqueryorder/app/bn/handler_response"
	servicerequest "tradethingqueryorder/app/bn/service_request"

	bnservice "tradethingqueryorder/app/bn/bn_service"

	bndynamodb "github.com/non26/tradepkg/pkg/bn/dynamodb_future"
)

type IService interface {
	QueryOrder(ctx context.Context, request *servicerequest.QueryOrder) (handlerresponse.QueryOrderResponse, error)
}

type service struct {
	bn_api_key               string
	bn_secret_key            string
	bn_position_url          string
	base_url                 string
	bnFtOpeningPositionTable bndynamodb.IBnFtOpeningPositionRepository
	bnFtQouteUsdtTable       bndynamodb.IBnFtQouteUSDTRepository
	bnFtHistoryTable         bndynamodb.IBnFtHistoryRepository
	bnService                bnservice.IBinanceFutureTradeService
}

func NewService(
	bn_api_key string,
	bn_secret_key string,
	bn_position_url string,
	base_url string,
	bnFtOpeningPositionTable bndynamodb.IBnFtOpeningPositionRepository,
	bnFtQouteUsdtTable bndynamodb.IBnFtQouteUSDTRepository,
	bnFtHistoryTable bndynamodb.IBnFtHistoryRepository,
	bnService bnservice.IBinanceFutureTradeService,
) IService {
	return &service{
		bn_api_key:               bn_api_key,
		bn_secret_key:            bn_secret_key,
		bn_position_url:          bn_position_url,
		base_url:                 base_url,
		bnFtOpeningPositionTable: bnFtOpeningPositionTable,
		bnFtQouteUsdtTable:       bnFtQouteUsdtTable,
		bnFtHistoryTable:         bnFtHistoryTable,
		bnService:                bnService,
	}
}
