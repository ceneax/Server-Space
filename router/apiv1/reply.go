package apiv1

import (
	"github.com/gin-gonic/gin"
	"xcdh.space/space/controller"
)

// GetReply godoc
// @Summary 获取回复
// @Description 根据页码等数据获取指定评论的回复
// @ID get-reply-page-json
// @Accept  application/json
// @Produce application/json
// @Param  page query  integer  true  "页码"
// @Param  pagesize  query  integer  false  "页大小"
// @Param  total  query  integer  false  "总数"
// @Param  sort  query  string  false  "排序方式"
// @Param  commentid  query  integer  true  "评论ID"
// @Success 200 {object} utils.Message
// @Failure 400 {object} utils.HTTPError
// @Failure 404 {object} utils.HTTPError
// @Failure 500 {object} utils.HTTPError
// @Router /v1/reply/get [get]
func GetReply(c *gin.Context) {
	controller.GetReplyList(c)
}

// GetReplyUser godoc
// @Summary 获取回复-带用户
// @Description 根据页码等数据获取指定评论的回复-带有用户权限
// @ID get-reply-page-user-json
// @Accept  application/json
// @Produce application/json
// @Param  page query  integer  true  "页码"
// @Param  pagesize  query  integer  false  "页大小"
// @Param  total  query  integer  false  "总数"
// @Param  sort  query  string  false  "排序方式"
// @Param  commentid  query  integer  true  "评论ID"
// @Param  userid  path  string  true  "用户ID"
// @Security ApiKeyAuth
// @Success 200 {object} utils.Message
// @Failure 400 {object} utils.HTTPError
// @Failure 404 {object} utils.HTTPError
// @Failure 500 {object} utils.HTTPError
// @Router /v1/reply/get/{userid} [get]
func GetReplyUser(c *gin.Context) {
	controller.GetReplyList(c)
}

// AddReply godoc
// @Summary 添加回复
// @Description 添加回复
// @ID post-reply-add-json
// @Accept  application/json
// @Produce application/json
// @Param  reply  body  service.Reply  true  "回复"
// @Security ApiKeyAuth
// @Success 200 {object} utils.Message
// @Failure 400 {object} utils.HTTPError
// @Failure 404 {object} utils.HTTPError
// @Failure 500 {object} utils.HTTPError
// @Router /v1/reply/add [post]
func AddReply(c *gin.Context) {
	controller.AddReply(c)
}

// UpdateReply godoc
// @Summary 更新回复
// @Description 获取回复内容
// @ID get-topic-content-json
// @Accept  application/json
// @Produce application/json
// @Param  topic  body  service.UpdateReply  true  "更新回复传入json"
// @Security ApiKeyAuth
// @Success 200 {object} utils.Message
// @Failure 400 {object} utils.HTTPError
// @Failure 404 {object} utils.HTTPError
// @Failure 500 {object} utils.HTTPError
// @Router /v1/reply/update [post]
func UpdateReply(c *gin.Context) {
	controller.UpdateReply(c)
}

// DeleteReply  godoc
// @Summary 删除回复
// @Description 删除回复
// @ID get-topic-content-json
// @Accept  application/json
// @Produce application/json
// @Param  topic  body  service.DeleteReply true  "删除回复传入json"
// @Security ApiKeyAuth
// @Success 200 {object} utils.Message
// @Failure 400 {object} utils.HTTPError
// @Failure 404 {object} utils.HTTPError
// @Failure 500 {object} utils.HTTPError
// @Router /v1/reply/delete [post]
func DeleteReply(c *gin.Context) {
	controller.DeleteReply(c)
}
