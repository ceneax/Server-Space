package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"xcdh.space/space/config"
	"xcdh.space/space/entity"
	"xcdh.space/space/service"
	"xcdh.space/space/utils"
)

func Register(c *gin.Context) {
	var user service.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	b, err := user.Register()
	if b {
		utils.NewMessage(c, 1, "注册成功，请登陆！")
	} else {
		utils.NewMessageWithError(c, err)
	}

}

func UpdateUserInfo(c *gin.Context) {
	var user service.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	v, ok := c.Get(config.IdentityKey)
	if v.(string) != user.UserId && !ok {
		utils.NewMessage(c, 0, "权限出错！")
		return
	}
	b, err := user.UpdateUserInfo()
	if b {
		utils.NewMessage(c, 1, "修改成功")
	} else {
		utils.NewMessageWithError(c, err)
	}
}

func UpdateUserPassword(c *gin.Context) {
	var user entity.ChangePassword
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	v, ok := c.Get(config.IdentityKey)
	if v.(string) != user.UserId && !ok {
		utils.NewMessage(c, 0, "权限出错！")
		return
	}
	b, err := service.UpdateUserPassword(user)
	if b {
		utils.NewMessage(c, 1, "修改成功")
	} else {
		utils.NewMessageWithError(c, err)
	}
}

func GetUserInfo(c *gin.Context) {
	userid := c.Param("userid")
	v, ok := c.Get(config.IdentityKey)
	if v.(string) != userid && !ok {
		utils.NewMessage(c, 0, "权限出错！")
		return
	}
	userinfo := service.GetUserInfo(userid)
	if userinfo != nil {
		utils.NewMessageWithData(c, 1, "查询成功", userinfo)
	} else {
		utils.NewMessage(c, 0, "查询无数据")
	}
}
