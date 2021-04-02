package apiv1

import (
	"github.com/gin-gonic/gin"
	"xcdh.space/space/controller"
)

// GetComment godoc
// @Summary 获取评论
// @Description 根据页码等数据获取指定主题的评论
// @ID get-comment-page-json
// @Accept  application/json
// @Produce application/json
// @Param  page query  integer  true  "页码"
// @Param  pagesize  query  integer  false  "页大小"
// @Param  total  query  integer  false  "总数"
// @Param  sort  query  string  false  "排序方式"
// @Param  topicid  query  integer  true  "主题ID"
// @Success 200 {object} utils.Message
// @Failure 400 {object} utils.HTTPError
// @Failure 404 {object} utils.HTTPError
// @Failure 500 {object} utils.HTTPError
// @Router /v1/comment/get [get]
func GetComment(c *gin.Context) {
	controller.GetCommentList(c)
}

// GetComment godoc
// @Summary 获取评论-带用户
// @Description 根据页码等数据获取指定主题的评论-带用户权限
// @ID get-comment-page-user-json
// @Accept  application/json
// @Produce application/json
// @Param  page query  integer  true  "页码"
// @Param  pagesize  query  integer  false  "页大小"
// @Param  total  query  integer  false  "总数"
// @Param  sort  query  string  false  "排序方式"
// @Param  topicid  query  integer  true  "主题ID"
// @Param  userid  path  string  true  "用户ID"
// @Security ApiKeyAuth
// @Success 200 {object} utils.Message
// @Failure 400 {object} utils.HTTPError
// @Failure 404 {object} utils.HTTPError
// @Failure 500 {object} utils.HTTPError
// @Router /v1/comment/get/{userid} [get]
func GetCommentUser(c *gin.Context) {
	controller.GetCommentList(c)
}

// AddComment godoc
// @Summary 添加评论
// @Description 添加评论
// @ID post-comment-add-json
// @Accept  application/json
// @Produce application/json
// @Param  comment  body  service.Comment  true  "主题"
// @Security ApiKeyAuth
// @Success 200 {object} utils.Message
// @Failure 400 {object} utils.HTTPError
// @Failure 404 {object} utils.HTTPError
// @Failure 500 {object} utils.HTTPError
// @Router /v1/comment/add [post]
func AddComment(c *gin.Context) {
	controller.AddComment(c)
}

// UpdateComment godoc
// @Summary 更新评论
// @Description 获取评论内容
// @ID get-topic-content-json
// @Accept  application/json
// @Produce application/json
// @Param  topic  body  service.UpdateComment  true  "更新评论传入json"
// @Security ApiKeyAuth
// @Success 200 {object} utils.Message
// @Failure 400 {object} utils.HTTPError
// @Failure 404 {object} utils.HTTPError
// @Failure 500 {object} utils.HTTPError
// @Router /v1/comment/update [post]
func UpdateComment(c *gin.Context) {
	controller.UpdateComment(c)
}

// DeleteComment  godoc
// @Summary 删除评论
// @Description 删除评论
// @ID get-topic-content-json
// @Accept  application/json
// @Produce application/json
// @Param  topic  body  service.DeleteComment true  "删除评论传入json"
// @Security ApiKeyAuth
// @Success 200 {object} utils.Message
// @Failure 400 {object} utils.HTTPError
// @Failure 404 {object} utils.HTTPError
// @Failure 500 {object} utils.HTTPError
// @Router /v1/comment/delete [post]
func DeleteComment(c *gin.Context) {
	controller.DeleteComment(c)
}
