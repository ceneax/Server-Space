package apiv1

import (
	"github.com/gin-gonic/gin"
	"xcdh.space/space/controller"
)

// SetEvaluate godoc
// @Summary 提交评价
// @Description 提交评价
// @ID set-evaluate-user-json
// @Accept  application/json
// @Produce application/json
// @Param  evaluate  body  service.Evaluate  true  "评价传入数据"
// @Security ApiKeyAuth
// @Success 200 {object} utils.Message
// @Failure 400 {object} utils.HTTPError
// @Failure 404 {object} utils.HTTPError
// @Failure 500 {object} utils.HTTPError
// @Router /v1/user/evaluate [post]
func SetEvaluate(c *gin.Context) {
	controller.SetEvaluate(c)
}
