package apiv1

import (
	"github.com/gin-gonic/gin"
	"xcdh.space/space/controller"
)

// Register godoc
// @Summary 注册
// @Description 注册用户
// @ID register-main-json
// @Accept  application/json
// @Produce application/json
// @Param  user body  service.User  true  "用户注册信息"
// @Success 200 {string} utils.Message
// @Failure 400 {object} utils.HTTPError
// @Failure 404 {object} utils.HTTPError
// @Failure 500 {object} utils.HTTPError
// @Router /v1/register [post]
func Register(c *gin.Context) {
	controller.Register(c)
}

// GetUserInfo godoc
// @Summary 获取用户信息
// @Description 根据用户id获取用户信息
// @ID get-user-info-json
// @Accept  application/json
// @Produce application/json
// @Param  userid  path  string  true  "用户名"
// @Security ApiKeyAuth
// @Success 200 {string} utils.Message
// @Failure 400 {object} utils.HTTPError
// @Failure 404 {object} utils.HTTPError
// @Failure 500 {object} utils.HTTPError
// @Router /v1/user/info/{userid} [get]
func GetUserInfo(c *gin.Context) {
	controller.GetUserInfo(c)
}

// UpdateUserInfo godoc
// @Summary 修改用户信息
// @Description 根据传过来的值修改用户信息
// @ID update-user-info-json
// @Accept  application/json
// @Produce application/json
// @Param  user  body  service.User  true  "用户注册信息"
// @Security ApiKeyAuth
// @Success 200 {string} utils.Message
// @Failure 400 {object} utils.HTTPError
// @Failure 404 {object} utils.HTTPError
// @Failure 500 {object} utils.HTTPError
// @Router /v1/user/updateinfo [post]
func UpdateUserInfo(c *gin.Context) {
	controller.UpdateUserInfo(c)
}

// UpdateUserPassword godoc
// @Summary 修改用户密码
// @Description 根据传过来的值修改用户密码
// @ID update-user-password-json
// @Accept  application/json
// @Produce application/json
// @Param  user  body  entity.ChangePassword  true  "用户密码修改"
// @Security ApiKeyAuth
// @Success 200 {string} utils.Message
// @Failure 400 {object} utils.HTTPError
// @Failure 404 {object} utils.HTTPError
// @Failure 500 {object} utils.HTTPError
// @Router /v1/user/updatepass [post]
func UpdateUserPassword(c *gin.Context) {
	controller.UpdateUserPassword(c)
}
