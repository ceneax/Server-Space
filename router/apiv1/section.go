package apiv1

import (
	"github.com/gin-gonic/gin"
	"xcdh.space/space/controller"
)

// GetAllSection godoc
// @Summary 获取板块
// @Description 获取所有板块
// @ID get-section-all-json
// @Produce application/json
// @Success 200 {string} utils.Message
// @Failure 400 {object} utils.HTTPError
// @Failure 404 {object} utils.HTTPError
// @Failure 500 {object} utils.HTTPError
// @Router /v1/section/all [get]
func GetAllSection(c *gin.Context) {
	controller.GetAllSection(c)
}

// AddSection godoc
// @Summary 添加板块
// @Description 添加板块
// @ID get-section-add-json
// @Produce application/json
// @Param  name  query  string  true  "板块名称"
// @Security ApiKeyAuth
// @Success 200 {string} utils.Message
// @Failure 400 {object} utils.HTTPError
// @Failure 404 {object} utils.HTTPError
// @Failure 500 {object} utils.HTTPError
// @Router /v1/section/add [get]
func AddSection(c *gin.Context) {
	controller.AddSection(c)
}
