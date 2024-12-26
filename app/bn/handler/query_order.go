package handler

import (
	"tradethingqueryorder/app/bn/service"

	"github.com/gin-gonic/gin"
)

type QueryOrderHandler struct {
	proxy service.IService
}

func NewQueryOrderHandler(proxy service.IService) *QueryOrderHandler {
	return &QueryOrderHandler{
		proxy: proxy,
	}
}

func (q *QueryOrderHandler) Handler(c *gin.Context) {
	// var req handlerrequest.QueryRequest
	// if err := c.ShouldBind(&req); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	// ctx := c.Request.Context()
	// q.proxy.QueryOrder(ctx)

}
