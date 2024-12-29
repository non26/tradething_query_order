package bnservice

import (
	"context"
	"fmt"
	"net/http"
	bntradereq "tradethingqueryorder/app/bn/bn_request"
	bntraderes "tradethingqueryorder/app/bn/bn_response"

	bnrequest "github.com/non26/tradepkg/pkg/bn/binance_request"
	bnresponse "github.com/non26/tradepkg/pkg/bn/binance_response"

	bncaller "github.com/non26/tradepkg/pkg/bn/binance_caller"
)

func (b *binanceFutureTradeService) QueryPositionV3(
	ctx context.Context,
	request *bntradereq.PositionInformationRequest) (map[string]*bntraderes.PositionInformation, error) {
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
		http.MethodGet,
		b.secretKey,
		b.apiKey,
		b.serviceName,
	)
	if err != nil {
		return nil, err
	}
	positionMap := make(map[string]*bntraderes.PositionInformation)
	if len(*res) > 0 {
		positionMap = make(map[string]*bntraderes.PositionInformation)
		for _, position := range *res {
			key := fmt.Sprintf("%s%s", position.Symbol, position.PositionSide)
			positionMap[key] = &position
		}
	}
	return positionMap, nil

	// bnsign := sign.NewSignHMACSHA256[bntradereq.PositionInformationRequest]("", b.secretKey)
	// signature, err := bnsign.GetQueryStringBinanceSignature(request)
	// if err != nil {
	// 	return nil, err
	// }
	// httpreq, err := http.NewRequest(http.MethodGet, b.bnPositionUrl, nil)
	// if err != nil {
	// 	return nil, err
	// }
	// httpreq.Header.Set("X-MBX-APIKEY", b.apiKey)
	// httpreq.Header.Set("Content-Type", "application/json")

}
