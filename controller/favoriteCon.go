package controller

import (
	"github.com/gin-gonic/gin"
	"xcdh.space/space/config"
	"xcdh.space/space/service"
	"xcdh.space/space/utils"
)

func SetFavorite(c *gin.Context) {
	var favorite service.Favorite
	if err := c.ShouldBindJSON(&favorite); err != nil {
		utils.NewMessageWithError(c, err)
		return
	}
	v, ok := c.Get(config.IdentityKey)
	if v.(string) != favorite.UserId && !ok {
		utils.NewMessage(c, 0, "权限出错！")
		return
	}

	b, err := favorite.SetFavorite()
	if b && err == nil {
		utils.NewMessage(c, 1, "收藏/取消收藏成功")
	} else {
		utils.NewMessageWithError(c, err)
	}
}
