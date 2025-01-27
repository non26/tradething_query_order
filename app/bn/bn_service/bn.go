package bnservice

import (
	"context"
	bntradereq "tradethingqueryorder/app/bn/bn_request"
	bntraderes "tradethingqueryorder/app/bn/bn_response"

	bnclient "github.com/non26/tradepkg/pkg/bn/bn_client"
	bntransport "github.com/non26/tradepkg/pkg/bn/bn_transport"
)

type IBinanceFutureTradeService interface {
	// return map[string]*bntraderes.PositionInformation
	// where key is symbol + positionSide
	QueryPositionV3(ctx context.Context, request *bntradereq.PositionInformationRequest) (map[string]*bntraderes.PositionInformation, error)
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
