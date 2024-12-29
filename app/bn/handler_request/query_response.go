package handlerrequest

import svcmodel "tradethingqueryorder/app/bn/service_request"

type QueryRequest struct {
	Symbol string `json:"symbol"`
}

func (q *QueryRequest) ToQueryOrder() *svcmodel.QueryOrder {
	return &svcmodel.QueryOrder{
		Symbol: q.Symbol,
	}
}
