package service

import (
	"context"

	handlerresponse "tradethingqueryorder/app/bn/handler_response"
	servicerequest "tradethingqueryorder/app/bn/service_request"
)

func (s *service) QueryOrder(ctx context.Context, request *servicerequest.QueryOrder) (handlerresponse.QueryOrderResponse, error) {
	return handlerresponse.QueryOrderResponse{}, nil
}
