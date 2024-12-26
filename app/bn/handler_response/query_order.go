package handlerresponse

type QueryOrderResponse struct {
	Data []QueryOrderResponseData `json:"data"`
}

type QueryOrderResponseData struct {
	Symbol       string `json:"symbol"`
	PositionSide string `json:"positionSide"`
}
