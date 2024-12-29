package bnrequest

import (
	"strconv"

	bnutils "github.com/non26/tradepkg/pkg/bn/utils"
)

type PositionInformationRequest struct {
	RecvWindow string `json:"recvWindow"`
	Timestamp  string `json:"timestamp"`
}

func (p *PositionInformationRequest) setTimestamp() {
	p.Timestamp = strconv.FormatInt(bnutils.GetBinanceTimestamp(), 10)
}

func (r *PositionInformationRequest) PrepareRequest() {
	r.setTimestamp()
	r.RecvWindow = "5000"
}

func (r *PositionInformationRequest) GetData() interface{} {
	return r
}
