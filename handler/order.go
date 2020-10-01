package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/i-coder-robot/gin-demo/enum"
	"github.com/i-coder-robot/gin-demo/model"
	"github.com/i-coder-robot/gin-demo/query"
	"github.com/i-coder-robot/gin-demo/resp"
	"github.com/i-coder-robot/gin-demo/service"
)

type OrderHandler struct {
	OrderSrv service.OrderSrv
}

func (h *OrderHandler) GetEntity(result model.Order) resp.Order {
	return resp.Order{
		Key:         result.OrderId,
		Id:          result.OrderId,
		OrderId:     result.OrderId,
		NickName:    result.NickName,
		Mobile:      result.Mobile,
		TotalPrice:  result.TotalPrice,
		PayStatus:   result.PayStatus,
		PayType:     result.PayType,
		PayTime:     result.PayTime,
		OrderStatus: result.OrderStatus,
		ExtraInfo:   result.ExtraInfo,
		UserAddress: result.UserAddress,
		IsDeleted:   result.IsDeleted,
	}
}

func (h *OrderHandler) OrderInfoHandler(c *gin.Context) {
	entity := resp.Entity{
		Code:      int(enum.OperateFail),
		Msg:       enum.OperateFail.String(),
		Total:     0,
		TotalPage: 1,
		Data:      nil,
	}
	orderId := c.Param("id")
	if orderId == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	u := model.Order{
		OrderId: orderId,
	}
	result, err := h.OrderSrv.Get(u)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}

	r := h.GetEntity(*result)

	entity = resp.Entity{
		Code:      http.StatusOK,
		Msg:       "OK",
		Total:     0,
		TotalPage: 0,
		Data:      r,
	}
	c.JSON(http.StatusOK, gin.H{"entity": entity})
}

func (h *OrderHandler) OrderListHandler(c *gin.Context) {
	var q query.ListQuery
	entity := resp.Entity{
		Code:      int(enum.OperateFail),
		Msg:       enum.OperateFail.String(),
		Total:     0,
		TotalPage: 1,
		Data:      nil,
	}
	err := c.ShouldBindQuery(&q)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	list, err := h.OrderSrv.List(&q)
	total, err := h.OrderSrv.GetTotal(&q)

	if err != nil {
		panic(err)
	}
	if q.PageSize == 0 {
		q.PageSize = 5
	}
	ret := int(total % q.PageSize)
	ret2 := int(total / q.PageSize)
	totalPage := 0
	if ret == 0 {
		totalPage = ret2
	} else {
		totalPage = ret2 + 1
	}
	var newList []*resp.Order
	for _, item := range list {
		r := h.GetEntity(*item)
		newList = append(newList, &r)
	}

	entity = resp.Entity{
		Code:      http.StatusOK,
		Msg:       "OK",
		Total:     total,
		TotalPage: totalPage,
		Data:      newList,
	}
	c.JSON(http.StatusOK, gin.H{"entity": entity})
}

func (h *OrderHandler) AddOrderHandler(c *gin.Context) {
	entity := resp.Entity{
		Code:  int(enum.OperateFail),
		Msg:   enum.OperateFail.String(),
		Total: 0,
		Data:  nil,
	}
	O := model.Order{}
	err := c.ShouldBindJSON(&O)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		return
	}

	r, err := h.OrderSrv.Add(O)
	if err != nil {
		entity.Msg = err.Error()
		return
	}
	if r.OrderId == "" {
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		return
	}
	entity.Code = int(enum.OperateOk)
	entity.Msg = enum.OperateOk.String()
	c.JSON(http.StatusOK, gin.H{"entity": entity})
}

func (h *OrderHandler) EditOrderHandler(c *gin.Context) {
	o := model.Order{}
	entity := resp.Entity{
		Code:  int(enum.OperateFail),
		Msg:   enum.OperateFail.String(),
		Total: 0,
		Data:  nil,
	}
	err := c.ShouldBindJSON(&o)
	if err != nil || o.OrderId == "" {
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		return
	}
	b, err := h.OrderSrv.Edit(o)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		return
	}
	if b {
		entity.Code = int(enum.OperateOk)
		entity.Msg = enum.OperateOk.String()
		c.JSON(http.StatusOK, gin.H{"entity": entity})
	}

}

func (h *OrderHandler) DeleteOrderHandler(c *gin.Context) {
	id := c.Param("id")
	r := h.OrderSrv.ExistByOrderID(id)
	b, err := h.OrderSrv.Delete(*r)
	entity := resp.Entity{
		Code:  int(enum.OperateFail),
		Msg:   enum.OperateFail.String(),
		Total: 0,
		Data:  nil,
	}
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		return
	}
	if b {
		entity.Code = int(enum.OperateOk)
		entity.Msg = enum.OperateOk.String()
		c.JSON(http.StatusOK, gin.H{"entity": entity})
	}
}
