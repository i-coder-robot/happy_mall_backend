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

type BannerHandler struct {
	BannerSrv service.BannerSrv
}

func (h *BannerHandler) GetEntity(result model.Banner) resp.Banner {
	return resp.Banner{
		Id:          result.BannerID,
		Key:         result.BannerID,
		BannerID:    result.BannerID,
		Url:         result.Url,
		RedirectUrl: result.RedirectUrl,
		OrderBy:     result.Order,
	}
}

func (h *BannerHandler) BannerInfoHandler(c *gin.Context) {
	entity := resp.Entity{
		Code:      int(enum.OperateFail),
		Msg:       enum.OperateFail.String(),
		Total:     0,
		TotalPage: 1,
		Data:      nil,
	}
	bannerId := c.Param("id")
	if bannerId == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	b := model.Banner{
		BannerID: bannerId,
	}
	result, err := h.BannerSrv.Get(b)

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

func (h *BannerHandler) BannerListHandler(c *gin.Context) {
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
	list, err := h.BannerSrv.List(&q)
	total, err := h.BannerSrv.GetTotal(&q)

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
	var newList []*resp.Banner
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

func (h *BannerHandler) AddBannerHandler(c *gin.Context) {
	entity := resp.Entity{
		Code:  int(enum.OperateFail),
		Msg:   enum.OperateFail.String(),
		Total: 0,
		Data:  nil,
	}
	banner := model.Banner{}
	err := c.ShouldBindJSON(&banner)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		return
	}

	r, err := h.BannerSrv.Add(banner)
	if err != nil {
		entity.Msg = err.Error()
		return
	}
	if r.BannerID == "" {
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		return
	}
	entity.Code = int(enum.OperateOk)
	entity.Msg = enum.OperateOk.String()
	c.JSON(http.StatusOK, gin.H{"entity": entity})
}

func (h *BannerHandler) EditBannerHandler(c *gin.Context) {
	banner := model.Banner{}
	entity := resp.Entity{
		Code:  int(enum.OperateFail),
		Msg:   enum.OperateFail.String(),
		Total: 0,
		Data:  nil,
	}
	err := c.ShouldBindJSON(&banner)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		return
	}
	b, err := h.BannerSrv.Edit(banner)
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

func (h *BannerHandler) DeleteBannerHandler(c *gin.Context) {
	id := c.Param("id")

	b, err := h.BannerSrv.Delete(id)
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
