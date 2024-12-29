package handler

import (
	"net/http"
	handlerrequest "tradethingqueryorder/app/bn/handler_request"
	"tradethingqueryorder/app/bn/service"

	"github.com/labstack/echo/v4"
)

type QueryOrderHandler struct {
	service service.IService
}

func NewQueryOrderHandler(service service.IService) *QueryOrderHandler {
	return &QueryOrderHandler{
		service: service,
	}
}

func (q *QueryOrderHandler) Handler(c echo.Context) error {
	request := new(handlerrequest.QueryRequest)
	err := c.Bind(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	ctx := c.Request().Context()
	res, err := q.service.QueryOrder(ctx, request.ToQueryOrder())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, res)

}
