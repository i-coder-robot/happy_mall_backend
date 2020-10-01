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

type CategoryHandler struct {
	CategorySrv service.CategorySrv
}

func (h *CategoryHandler) CategoryList4BackendHandler(c *gin.Context) {
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

	if q.PageSize == 0 {
		q.PageSize = 5
	}

	list, err := h.CategorySrv.List(&q)
	total, err := h.CategorySrv.GetTotal(&q)

	for _, item := range list {
		item.Key = item.C3CategoryID
		item.Id = item.C3CategoryID
	}

	pageTotal := 0
	if total%q.PageSize == 0 {
		pageTotal = int(total / q.PageSize)
	} else {
		pageTotal = int(total/q.PageSize) + 1
	}

	entity = resp.Entity{
		Code:      http.StatusOK,
		Msg:       "OK",
		Total:     total,
		TotalPage: pageTotal,
		Data:      list,
	}
	c.JSON(http.StatusOK, gin.H{"entity": entity})
}

func (h *CategoryHandler) CategoryListHandler(c *gin.Context) {
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

	if q.PageSize == 0 {
		q.PageSize = 5
	}

	list, err := h.CategorySrv.List(&q)
	total, err := h.CategorySrv.GetTotal(&q)

	newList := h.GetEntity(list)

	pageTotal := 0
	if total%q.PageSize == 0 {
		pageTotal = int(total / q.PageSize)
	} else {
		pageTotal = int(total/q.PageSize) + 1
	}

	entity = resp.Entity{
		Code:      http.StatusOK,
		Msg:       "OK",
		Total:     total,
		TotalPage: pageTotal,
		Data:      newList,
	}
	c.JSON(http.StatusOK, gin.H{"entity": entity})
}

func (h *CategoryHandler) GetEntity(result []*model.CategoryResult) map[string]*resp.Category {
	c3map := make(map[string]*resp.Category3)
	for _, item := range result {
		thirdCategory := &resp.Category3{
			Id:         item.C3CategoryID,
			Key:        item.C3CategoryID,
			CategoryID: item.C3CategoryID,
			Name:       item.C3Name,
			Order:      item.C3Order,
			ParentID:   item.C3ParentId,
			IsDeleted:  item.IsDeleted,
		}
		c3map[item.C3CategoryID] = thirdCategory
	}

	c2map := make(map[string]*resp.Category2)
	for _, item := range result {
		secondCategory := &resp.Category2{
			CategoryID: item.C2CategoryID,
			Name:       item.C2Name,
			Order:      item.C2Order,
			ParentID:   item.C2ParentId,
			Children:   c3map,
		}
		c2map[item.C2CategoryID] = secondCategory
	}

	cmap := make(map[string]*resp.Category)
	for _, item := range result {
		firstCategory := &resp.Category{
			CategoryID: item.C1CategoryID,
			Name:       item.C1Name,
			Desc:       item.C1Desc,
			Order:      item.C1Order,
			ParentID:   item.C1ParentId,
			Children:   c2map,
		}
		cmap[item.C1CategoryID] = firstCategory
	}
	return cmap
}

func (h *CategoryHandler) CategoryInfoHandler(c *gin.Context) {
	//这个info传入的是第三级categorydeId
	entity := resp.Entity{
		Code:      int(enum.OperateFail),
		Msg:       enum.OperateFail.String(),
		Total:     0,
		TotalPage: 1,
		Data:      nil,
	}
	categoryId := c.Param("id")
	if categoryId == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}

	result, err := h.CategorySrv.Get(categoryId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}

	listCategories := h.GetEntity(result)

	entity = resp.Entity{
		Code:      http.StatusOK,
		Msg:       "OK",
		Total:     0,
		TotalPage: 0,
		Data:      listCategories,
	}
	c.JSON(http.StatusOK, gin.H{"entity": entity})
}

func (h *CategoryHandler) AddCategoryHandler(c *gin.Context) {
	entity := resp.Entity{
		Code:  int(enum.OperateFail),
		Msg:   enum.OperateFail.String(),
		Total: 0,
		Data:  nil,
	}
	category := model.CategoryResult{}
	err := c.ShouldBindJSON(&category)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		return
	}

	b, err := h.CategorySrv.Add(category)
	if err != nil {
		entity.Msg = err.Error()
		return
	}
	if !b {
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		return
	}
	entity.Code = int(enum.OperateOk)
	entity.Msg = enum.OperateOk.String()
	c.JSON(http.StatusOK, gin.H{"entity": entity})

}

func (h *CategoryHandler) EditCategoryHandler(c *gin.Context) {
	category := model.Category{}
	entity := resp.Entity{
		Code:  int(enum.OperateFail),
		Msg:   enum.OperateFail.String(),
		Total: 0,
		Data:  nil,
	}
	err := c.ShouldBindJSON(&category)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		return
	}
	b, err := h.CategorySrv.Edit(category)
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

func (h *CategoryHandler) DeleteCategoryHandler(c *gin.Context) {
	id := c.Param("id")
	r := h.CategorySrv.ExistByCategoryID(id)
	b, err := h.CategorySrv.Delete(*r)
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
