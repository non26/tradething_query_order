package bnrequest

type PositionInformationRequest struct {
	Symbol    string  `json:"symbol"`
	Timestamp float64 `json:"timestamp"`
}

func (r *PositionInformationRequest) PrepareRequest() {

}

func (r *PositionInformationRequest) GetData() interface{} {
	return r
}
