package bnservice

import (
	"context"
	"net/http"
	bntradereq "tradethingqueryorder/app/bn/bn_request"
	bntraderes "tradethingqueryorder/app/bn/bn_response"

	bncaller "github.com/non26/tradepkg/pkg/bn/binance_caller"
	bnrequest "github.com/non26/tradepkg/pkg/bn/binance_request"
	bnresponse "github.com/non26/tradepkg/pkg/bn/binance_response"
)

func (b *binanceFutureTradeService) QueryPositionV3(
	ctx context.Context,
	request *bntradereq.PositionInformationRequest) (bntraderes.PositionsInFormationResponse, error) {
	c := bncaller.NewCallBinance(
		bnrequest.NewBinanceServiceHttpRequest[bntradereq.PositionInformationRequest](),
		bnresponse.NewBinanceServiceHttpResponse[bntraderes.PositionsInFormationResponse](),
		b.httpttransport,
		b.httpclient,
	)
	res, err := c.CallBinance(
		request,
		b.bnBaseUrl,
		b.bnPositionUrl,
		http.MethodPost,
		b.secretKey,
		b.apiKey,
		b.serviceName,
	)
	if err != nil {
		return bntraderes.PositionsInFormationResponse{}, err
	}
	return *res, nil
}
