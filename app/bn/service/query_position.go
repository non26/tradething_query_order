package service

import (
	"context"
	"fmt"
	"strconv"
	"time"

	handlerresponse "tradethingqueryorder/app/bn/handler_response"
	servicerequest "tradethingqueryorder/app/bn/service_request"

	dynamodbrepository "github.com/non26/tradepkg/pkg/bn/dynamodb_repository/models"
	positionconstant "github.com/non26/tradepkg/pkg/bn/position_constant"
)

func (s *service) QueryOrder(ctx context.Context, request *servicerequest.QueryOrder) (handlerresponse.QueryOrderResponse, error) {
	bnPositions, err := s.bnService.QueryPositionV3(ctx, request.ToBnServiceRequest())
	if err != nil {
		return handlerresponse.QueryOrderResponse{}, err
	}

	dbPositions, err := s.repository.GetAllOpenOrders(ctx)
	if err != nil {
		return handlerresponse.QueryOrderResponse{}, err
	}
	if bnPositions != nil {
		// found position in Binance but not position in DynamoDB then create new position
		if len(dbPositions) == 0 && len(bnPositions) > 0 {
			for _, position := range bnPositions {
				qoute, err := s.repository.GetQouteUSDT(ctx, position.Symbol)
				if !qoute.IsFound() || err != nil {
					qoute.SetSymbol(position.Symbol)
					if position.PositionSide == positionconstant.LONG {
						qoute.SetCountingLong(1)
						qoute.SetCountingShort(0)
					} else {
						qoute.SetCountingLong(0)
						qoute.SetCountingShort(1)
					}
					err = s.repository.InsertNewSymbolQouteUSDT(ctx, qoute)
					if err != nil {
						fmt.Println("error insert new symbol qoute usdt", err)
					}
				}
				var side string
				var clientId string
				if position.PositionSide == positionconstant.LONG {
					side = positionconstant.BUY
					clientId = createDefaultClientId(position.Symbol, position.PositionSide, qoute.CountingLong)
				} else {
					side = positionconstant.SELL
					clientId = createDefaultClientId(position.Symbol, position.PositionSide, qoute.CountingShort)
				}

				err = s.repository.InsertNewOpenOrder(ctx, position.ToOpenPositionDynamodb(clientId, side))
				if err != nil {
					fmt.Println("error insert new open order", err)
				}
			}

			return handlerresponse.QueryOrderResponse{}, nil
		}

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

			dbAmountq, _ := strconv.ParseFloat(dbPosition.AmountQ, 64)
			bnAmountq, _ := strconv.ParseFloat(bnPosition.PositionAmt, 64)
			if dbAmountq != bnAmountq {
				dbPosition.AmountQ = bnPosition.PositionAmt
				need_update = true
			}
			if need_update {
				err = s.repository.UpdateOpenOrder(ctx, dbPosition)
				if err != nil {
					fmt.Println("error update open order", err)
				}
			}
		}

		for _, key := range inBnNotDb {
			bnPosition := bnPositions[key]
			qoute, err := s.repository.GetQouteUSDT(ctx, bnPosition.Symbol)
			if !qoute.IsFound() || err != nil {
				qoute.SetSymbol(bnPosition.Symbol)
				if bnPosition.PositionSide == positionconstant.LONG {
					qoute.SetCountingLong(1)
					qoute.SetCountingShort(0)
				} else {
					qoute.SetCountingLong(0)
					qoute.SetCountingShort(1)
				}
				err = s.repository.InsertNewSymbolQouteUSDT(ctx, qoute)
				if err != nil {
					fmt.Println("error insert new symbol qoute usdt", err)
				}
			}
			var side string
			var clientId string
			if bnPosition.PositionSide == positionconstant.LONG {
				side = positionconstant.BUY
				clientId = createDefaultClientId(bnPosition.Symbol, bnPosition.PositionSide, qoute.CountingLong)
			} else {
				side = positionconstant.SELL
				clientId = createDefaultClientId(bnPosition.Symbol, bnPosition.PositionSide, qoute.CountingShort)
			}
			err = s.repository.InsertNewOpenOrder(ctx, bnPosition.ToOpenPositionDynamodb(clientId, side))
			if err != nil {
				fmt.Println("error insert new open order", err)
			}
		}

		for _, key := range inDbNotBn {
			dbPosition := dbPositions[key]
			err = s.repository.DeleteOpenOrderBySymbolAndPositionSide(ctx, dbPosition)
			if err != nil {
				fmt.Println("error delete open order", err)
			}

			s.repository.InsertHistory(ctx, &dynamodbrepository.BnFtHistory{
				ClientId:     dbPosition.ClientId,
				Symbol:       dbPosition.Symbol,
				PositionSide: dbPosition.PositionSide,
			})
		}

	}

	return handlerresponse.QueryOrderResponse{}, nil
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
