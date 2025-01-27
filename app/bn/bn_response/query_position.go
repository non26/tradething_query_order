package bnresponse

import (
	"fmt"
	"strconv"

	dynamodbmodels "github.com/non26/tradepkg/pkg/bn/dynamodb_repository/models"
)

type PositionsInFormationResponse []PositionInformation

func (p *PositionsInFormationResponse) IsFound() bool {
	return len(*p) > 0
}

type PositionInformation struct {
	Symbol                 string `json:"symbol"`
	PositionSide           string `json:"positionSide"`
	PositionAmt            string `json:"positionAmt"`
	EntryPrice             string `json:"entryPrice"`
	BreakEvenPrice         string `json:"breakEvenPrice"`
	MarkPrice              string `json:"markPrice"`
	UnRealizedProfit       string `json:"unRealizedProfit"`
	LiquidationPrice       string `json:"liquidationPrice"`
	IsolatedMargin         string `json:"isolatedMargin"`
	Notional               string `json:"notional"`
	MarginAsset            string `json:"marginAsset"`
	IsolatedWallet         string `json:"isolatedWallet"`
	InitialMargin          string `json:"initialMargin"`
	MaintMargin            string `json:"maintMargin"`
	PositionInitialMargin  string `json:"positionInitialMargin"`
	OpenOrderInitialMargin string `json:"openOrderInitialMargin"`
	Adl                    int    `json:"adl"`
	BidNotional            string `json:"bidNotional"`
	AskNotional            string `json:"askNotional"`
	UpdateTime             int64  `json:"updateTime"`
}

func (p *PositionInformation) ToOpenPositionDynamodb(clientId string, side string) *dynamodbmodels.BnFtOpeningPosition {
	return &dynamodbmodels.BnFtOpeningPosition{
		Symbol:       p.Symbol,
		PositionSide: p.PositionSide,
		AmountQ:      p.PositionAmt,
		ClientId:     clientId,
		Side:         side,
	}
}

func (p *PositionInformation) GetUnSignedPositionAmt() string {
	amt, _ := strconv.ParseFloat(p.PositionAmt, 64)
	if amt < 0 {
		return fmt.Sprintf("%v", amt*-1)
	}
	return p.PositionAmt
}
