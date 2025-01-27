package service

import (
	"context"
	"fmt"
	"time"

	handlerresponse "tradethingqueryorder/app/bn/handler_response"
	servicerequest "tradethingqueryorder/app/bn/service_request"

	bnconstant "github.com/non26/tradepkg/pkg/bn/bn_constant"
	dynamodbrepository "github.com/non26/tradepkg/pkg/bn/dynamodb_future/models"
	"github.com/non26/tradepkg/pkg/bn/utils"
	"github.com/shopspring/decimal"
)

func (s *service) QueryOrder(ctx context.Context, request *servicerequest.QueryOrder) (handlerresponse.QueryOrderResponse, error) {
	response := handlerresponse.QueryOrderResponse{
		Data: []handlerresponse.QueryOrderResponseData{},
	}
	bnPositions, err := s.bnService.QueryPositionV3(ctx, request.ToBnServiceRequest())
	if err != nil {
		fmt.Println("error query position", err)
		return handlerresponse.QueryOrderResponse{}, err
	}

	dbPositions, err := s.bnFtOpeningPositionTable.GetAll(ctx)
	if err != nil {
		fmt.Println("error get all open orders", err)
		return handlerresponse.QueryOrderResponse{}, err
	}
	if bnPositions != nil {
		// found position in Binance but not position in DynamoDB then create new position
		if len(dbPositions) == 0 && len(bnPositions) > 0 {
			for _, position := range bnPositions {
				qoute, err := s.bnFtQouteUsdtTable.Get(ctx, position.Symbol)
				if !qoute.IsFound() || err != nil {
					qoute.SetSymbol(position.Symbol)
					if utils.IsLongPosition(position.PositionSide) {
						qoute.SetCountingLong(1)
						qoute.SetCountingShort(0)
					} else {
						qoute.SetCountingLong(0)
						qoute.SetCountingShort(1)
					}
					err = s.bnFtQouteUsdtTable.Insert(ctx, qoute)
					if err != nil {
						fmt.Println("error insert new symbol qoute usdt", err)
					}
				}
				var side string
				var clientId string
				if utils.IsLongPosition(position.PositionSide) {
					side = bnconstant.BUY
					clientId = createDefaultClientId(position.Symbol, position.PositionSide, qoute.CountingLong)
				} else {
					side = bnconstant.SELL
					clientId = createDefaultClientId(position.Symbol, position.PositionSide, qoute.CountingShort)
				}

				err = s.bnFtOpeningPositionTable.Insert(ctx, position.ToOpenPositionDynamodb(clientId, side))
				if err != nil {
					fmt.Println("error insert new open order", err)
				}

				response.Data = append(response.Data, handlerresponse.QueryOrderResponseData{
					Symbol:       position.Symbol,
					PositionSide: position.PositionSide,
				})
			}

			return response, nil
		}

		// found position in Binance and position in DynamoDB then update position
		mutaulKey, inBnNotDb, inDbNotBn := compareMapKey(bnPositions, dbPositions)
		for _, key := range mutaulKey {
			need_update := false
			bnPosition := bnPositions[key]
			dbPosition := dbPositions[key]
			if bnPosition.PositionSide != dbPosition.PositionSide {
				dbPosition.PositionSide = bnPosition.PositionSide
				need_update = true
			}
			if bnPosition.Symbol != dbPosition.Symbol {
				dbPosition.Symbol = bnPosition.Symbol
				need_update = true
			}

			dbAmountq, _ := decimal.NewFromString(dbPosition.AmountB)
			bnAmountq, _ := decimal.NewFromString(bnPosition.PositionAmt)
			if dbAmountq.Cmp(bnAmountq) != 0 {
				dbPosition.AmountB = bnPosition.PositionAmt
				need_update = true
			}
			if need_update {
				err = s.bnFtOpeningPositionTable.UpdateAmountB(ctx, dbPosition)
				if err != nil {
					fmt.Println("error update open order", err)
				}
			}
			response.Data = append(response.Data, handlerresponse.QueryOrderResponseData{
				Symbol:       bnPosition.Symbol,
				PositionSide: bnPosition.PositionSide,
			})
		}

		// found position in Binance but not position in DynamoDB then create new position in db
		for _, key := range inBnNotDb {
			bnPosition := bnPositions[key]
			qoute, err := s.bnFtQouteUsdtTable.Get(ctx, bnPosition.Symbol)
			if !qoute.IsFound() || err != nil {
				qoute.SetSymbol(bnPosition.Symbol)
				if bnPosition.PositionSide == bnconstant.LONG {
					qoute.SetCountingLong(1)
					qoute.SetCountingShort(0)
				} else {
					qoute.SetCountingLong(0)
					qoute.SetCountingShort(1)
				}
				err = s.bnFtQouteUsdtTable.Insert(ctx, qoute)
				if err != nil {
					fmt.Println("error insert new symbol qoute usdt", err)
				}
			}
			var side string
			var clientId string
			if bnPosition.PositionSide == bnconstant.LONG {
				side = bnconstant.BUY
				clientId = createDefaultClientId(bnPosition.Symbol, bnPosition.PositionSide, qoute.CountingLong)
			} else {
				side = bnconstant.SELL
				clientId = createDefaultClientId(bnPosition.Symbol, bnPosition.PositionSide, qoute.CountingShort)
			}
			err = s.bnFtOpeningPositionTable.Insert(ctx, bnPosition.ToOpenPositionDynamodb(clientId, side))
			if err != nil {
				fmt.Println("error insert new open order", err)
			}
			response.Data = append(response.Data, handlerresponse.QueryOrderResponseData{
				Symbol:       bnPosition.Symbol,
				PositionSide: bnPosition.PositionSide,
			})
		}

		// found position in DynamoDB but not position in Binance then delete position
		for _, key := range inDbNotBn {
			dbPosition := dbPositions[key]
			err = s.bnFtOpeningPositionTable.Delete(ctx, dbPosition)
			if err != nil {
				fmt.Println("error delete open order", err)
			}

			s.bnFtHistoryTable.Insert(ctx, &dynamodbrepository.BnFtHistory{
				ClientId:     dbPosition.ClientId,
				Symbol:       dbPosition.Symbol,
				PositionSide: dbPosition.PositionSide,
			})
			response.Data = append(response.Data, handlerresponse.QueryOrderResponseData{
				Symbol:       dbPosition.Symbol,
				PositionSide: dbPosition.PositionSide,
			})
		}
	}

	return response, nil
}

func createDefaultClientId(symbol string, positionSide string, counting int) string {
	loc, _ := time.LoadLocation("Asia/Bangkok")
	return fmt.Sprintf("default_%s_%s_%s-%d", symbol, positionSide, time.Now().In(loc).Format("20060102150405"), counting)
}

func compareMapKey[T any, K any](a map[string]*T, b map[string]*K) (mutualKey []string, in_a_not_b []string, in_b_not_a []string) {
	mutualKey = []string{}
	in_a_not_b = []string{}
	in_b_not_a = []string{}

	for akey := range a {
		if _, ok := b[akey]; ok {
			mutualKey = append(mutualKey, akey)
		} else {
			in_a_not_b = append(in_a_not_b, akey)
		}
	}

	for bkey := range b {
		if _, ok := a[bkey]; !ok {
			in_b_not_a = append(in_b_not_a, bkey)
		}
	}

	return mutualKey, in_a_not_b, in_b_not_a
}
