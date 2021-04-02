package apiv1

import (
	"github.com/gin-gonic/gin"
	"xcdh.space/space/controller"
)

// GetTopic godoc
// @Summary 获取主题
// @Description 根据页码等数据指定板块获取主题
// @ID get-topic-page-json
// @Accept  application/json
// @Produce application/json
// @Param  page query  integer  true  "页码"
// @Param  pagesize  query  integer  false  "页大小"
// @Param  total  query  integer  false  "总数"
// @Param  sectionid  query  integer  true  "板块ID"
// @Success 200 {object} utils.Message
// @Failure 400 {object} utils.HTTPError
// @Failure 404 {object} utils.HTTPError
// @Failure 500 {object} utils.HTTPError
// @Router /v1/topic/get [get]
func GetTopic(c *gin.Context) {
	controller.GetTopicList(c)
}

// GetTopicInUser godoc
// @Summary 获取主题-带用户
// @Description 根据页码等数据指定板块获取主题-带有用户权限
// @ID get-topic-page-user-json
// @Accept  application/json
// @Produce application/json
// @Param  page query  integer  true  "页码"
// @Param  pagesize  query  integer  false  "页大小"
// @Param  total  query  integer  false  "总数"
// @Param  sectionid  query  integer  true  "板块ID"
// @Param  userid  path  string  true  "用户ID"
// @Security ApiKeyAuth
// @Success 200 {object} utils.Message
// @Failure 400 {object} utils.HTTPError
// @Failure 404 {object} utils.HTTPError
// @Failure 500 {object} utils.HTTPError
// @Router /v1/topic/get/{userid} [get]
func GetTopicInUser(c *gin.Context) {
	controller.GetTopicList(c)
}

// AddTopic godoc
// @Summary 添加主题
// @Description 添加主题
// @ID post-topic-add-json
// @Accept  application/json
// @Produce application/json
// @Param  topic  body  service.Topic  true  "主题"
// @Security ApiKeyAuth
// @Success 200 {object} utils.Message
// @Failure 400 {object} utils.HTTPError
// @Failure 404 {object} utils.HTTPError
// @Failure 500 {object} utils.HTTPError
// @Router /v1/topic/add [post]
func AddTopic(c *gin.Context) {
	controller.AddTopic(c)
}

// GetTopicContent godoc
// @Summary 获取主题内容
// @Description 获取主题内容
// @ID get-topic-content-json
// @Accept  application/json
// @Produce application/json
// @Param  topicid  query  integer  true  "主题ID"
// @Param  sectionid  query  integer true  "板块ID"
// @Success 200 {object} utils.Message
// @Failure 400 {object} utils.HTTPError
// @Failure 404 {object} utils.HTTPError
// @Failure 500 {object} utils.HTTPError
// @Router /v1/topic/content/get [get]
func GetTopicContent(c *gin.Context) {
	controller.GetTopicContent(c)
}

// UpdateTopic godoc
// @Summary 更新主题
// @Description 获取主题内容
// @ID get-topic-content-json
// @Accept  application/json
// @Produce application/json
// @Param  topic  body  service.UpdateTopic  true  "更新主题传入json"
// @Security ApiKeyAuth
// @Success 200 {object} utils.Message
// @Failure 400 {object} utils.HTTPError
// @Failure 404 {object} utils.HTTPError
// @Failure 500 {object} utils.HTTPError
// @Router /v1/topic/update [post]
func UpdateTopic(c *gin.Context) {
	controller.UpdateTopic(c)
}

// DeleteTopic godoc
// @Summary 删除主题
// @Description 删除主题
// @ID get-topic-content-json
// @Accept  application/json
// @Produce application/json
// @Param  topic  body  service.DeleteTopic  true  "更新主题传入json"
// @Security ApiKeyAuth
// @Success 200 {object} utils.Message
// @Failure 400 {object} utils.HTTPError
// @Failure 404 {object} utils.HTTPError
// @Failure 500 {object} utils.HTTPError
// @Router /v1/topic/delete [post]
func DeleteTopic(c *gin.Context) {
	controller.DeleteTopic(c)
}
