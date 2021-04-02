package apiv1

import (
	"github.com/gin-gonic/gin"
	"xcdh.space/space/controller"
)

// GetNotice godoc
// @Summary 获取通知
// @Description 根据用户名获取通知
// @ID get-notice-user-json
// @Accept  application/json
// @Produce application/json
// @Param  userid  path  string  true  "用户名"
// @Security ApiKeyAuth
// @Success 200 {object} utils.Message
// @Failure 400 {object} utils.HTTPError
// @Failure 404 {object} utils.HTTPError
// @Failure 500 {object} utils.HTTPError
// @Router /v1/notice/{userid} [get]
func GetNotice(c *gin.Context) {
	controller.GetNoticeList(c)
}

// SetNoticeRead godoc
// @Summary 设置通知已读
// @Description 根据用户名设置通知已读
// @ID set-notice-user-json
// @Accept  application/json
// @Produce application/json
// @Param  userid  path  string  true  "用户名"
// @Security ApiKeyAuth
// @Success 200 {object} utils.Message
// @Failure 400 {object} utils.HTTPError
// @Failure 404 {object} utils.HTTPError
// @Failure 500 {object} utils.HTTPError
// @Router /v1/noticeread/{userid} [get]
func SetNoticeRead(c *gin.Context) {
	controller.SetNoticeRead(c)
}
