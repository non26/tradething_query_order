package servicerequest

import bntradereq "tradethingqueryorder/app/bn/bn_request"

type QueryOrder struct {
	Symbol string
}

func NewQueryOrder(symbol string) *QueryOrder {
	return &QueryOrder{Symbol: symbol}
}

func (q *QueryOrder) ToBnServiceRequest() *bntradereq.PositionInformationRequest {
	return &bntradereq.PositionInformationRequest{
		// Symbol: q.Symbol,
	}
}
