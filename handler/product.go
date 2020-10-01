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

type ProductHandler struct {
	ProductSrv service.ProductSrv
}

func (h *ProductHandler) GetEntity(result model.Product) resp.Product {
	return resp.Product{
		Id:                   result.ProductId,
		Key:                  result.ProductId,
		ProductId:            result.ProductId,
		ProductName:          result.ProductName,
		ProductIntro:         result.ProductIntro,
		CategoryId:           result.CategoryId,
		ProductCoverImg:      result.ProductCoverImg,
		ProductBanner:        result.ProductBanner,
		OriginalPrice:        result.OriginalPrice,
		SellingPrice:         result.SellingPrice,
		StockNum:             result.StockNum,
		Tag:                  result.Tag,
		SellStatus:           result.SellStatus,
		ProductDetailContent: result.ProductDetailContent,
		IsDeleted:            result.IsDeleted,
	}
}

func (h *ProductHandler) ProductInfoHandler(c *gin.Context) {
	entity := resp.Entity{
		Code:      int(enum.OperateFail),
		Msg:       enum.OperateFail.String(),
		Total:     0,
		TotalPage: 1,
		Data:      nil,
	}
	productId := c.Param("id")
	if productId == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	u := model.Product{
		ProductId: productId,
	}
	result, err := h.ProductSrv.Get(u)

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

func (h *ProductHandler) ProductListHandler(c *gin.Context) {
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
	list, err := h.ProductSrv.List(&q)
	total, err := h.ProductSrv.GetTotal(&q)

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
	var newList []*resp.Product
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

func (h *ProductHandler) AddProductHandler(c *gin.Context) {
	entity := resp.Entity{
		Code:  int(enum.OperateFail),
		Msg:   enum.OperateFail.String(),
		Total: 0,
		Data:  nil,
	}
	p := model.Product{}
	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		return
	}

	r, err := h.ProductSrv.Add(p)
	if err != nil {
		entity.Msg = err.Error()
		return
	}
	if r.ProductId == "" {
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		return
	}
	entity.Code = int(enum.OperateOk)
	entity.Msg = enum.OperateOk.String()
	c.JSON(http.StatusOK, gin.H{"entity": entity})

}

func (h *ProductHandler) EditProductHandler(c *gin.Context) {
	p := model.Product{}
	entity := resp.Entity{
		Code:  int(enum.OperateFail),
		Msg:   enum.OperateFail.String(),
		Total: 0,
		Data:  nil,
	}
	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		return
	}
	b, err := h.ProductSrv.Edit(p)
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

func (h *ProductHandler) DeleteProductHandler(c *gin.Context) {
	id := c.Param("id")
	b, err := h.ProductSrv.Delete(id)
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
