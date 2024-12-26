package bnservice

import (
	"context"
	bnrequest "tradethingqueryorder/app/bn/bn_request"
	bnresponse "tradethingqueryorder/app/bn/bn_response"

	bnclient "github.com/non26/tradepkg/pkg/bn/binance_client"
	bntransport "github.com/non26/tradepkg/pkg/bn/binance_transport"
)

type IBinanceFutureTradeService interface {
	QueryPositionV3(ctx context.Context, request *bnrequest.PositionInformationRequest) (bnresponse.PositionsInFormationResponse, error)
}

type binanceFutureTradeService struct {
	bnBaseUrl      string
	bnPositionUrl  string
	apiKey         string
	secretKey      string
	httpttransport bntransport.IBinanceServiceHttpTransport
	httpclient     bnclient.IBinanceSerivceHttpClient
	serviceName    string
}

func NewBinanceFutureTradeService(
	bnBaseUrl string,
	bnPositionUrl string,
	apiKey string,
	secretKey string,
	httpttransport bntransport.IBinanceServiceHttpTransport,
	httpclient bnclient.IBinanceSerivceHttpClient,
	serviceName string,
) IBinanceFutureTradeService {
	return &binanceFutureTradeService{
		bnBaseUrl:      bnBaseUrl,
		bnPositionUrl:  bnPositionUrl,
		apiKey:         apiKey,
		secretKey:      secretKey,
		httpttransport: httpttransport,
		httpclient:     httpclient,
		serviceName:    serviceName,
	}
}
