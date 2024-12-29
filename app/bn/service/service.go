package service

import (
	"context"
	handlerresponse "tradethingqueryorder/app/bn/handler_response"
	servicerequest "tradethingqueryorder/app/bn/service_request"

	bnservice "tradethingqueryorder/app/bn/bn_service"

	dynamodbrepository "github.com/non26/tradepkg/pkg/bn/dynamodb_repository"
)

type IService interface {
	QueryOrder(ctx context.Context, request *servicerequest.QueryOrder) (handlerresponse.QueryOrderResponse, error)
}

type service struct {
	bn_api_key      string
	bn_secret_key   string
	bn_position_url string
	base_url        string
	repository      dynamodbrepository.IRepository
	bnService       bnservice.IBinanceFutureTradeService
}

func NewService(
	bn_api_key string,
	bn_secret_key string,
	bn_position_url string,
	base_url string,
	repository dynamodbrepository.IRepository,
	bnService bnservice.IBinanceFutureTradeService,
) IService {
	return &service{
		bn_api_key:      bn_api_key,
		bn_secret_key:   bn_secret_key,
		bn_position_url: bn_position_url,
		base_url:        base_url,
		repository:      repository,
		bnService:       bnService,
	}
}
