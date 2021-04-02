package apiv1

import (
	"github.com/gin-gonic/gin"
	"xcdh.space/space/controller"
)

// SetFavorite godoc
// @Summary 提交收藏
// @Description 提交收藏或取消收藏
// @ID set-favorite-user-json
// @Accept  application/json
// @Produce application/json
// @Param  favorite  body  service.Favorite  true  "收藏传入数据"
// @Security ApiKeyAuth
// @Success 200 {object} utils.Message
// @Failure 400 {object} utils.HTTPError
// @Failure 404 {object} utils.HTTPError
// @Failure 500 {object} utils.HTTPError
// @Router /v1/user/favorite [post]
func SetFavorite(c *gin.Context) {
	controller.SetFavorite(c)
}
