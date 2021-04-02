package controller

import (
	"github.com/gin-gonic/gin"
	"xcdh.space/space/config"
	"xcdh.space/space/service"
	"xcdh.space/space/utils"
)

func SetEvaluate(c *gin.Context) {
	var evaluate service.Evaluate
	if err := c.ShouldBindJSON(&evaluate); err != nil {
		utils.NewMessageWithError(c, err)
		return
	}
	v, ok := c.Get(config.IdentityKey)
	if v.(string) != evaluate.UserId && !ok {
		utils.NewMessage(c, 0, "权限出错！")
		return
	}

	b, err := evaluate.SetEvaluate()
	if b && err == nil {
		utils.NewMessage(c, 1, "评价成功")
	} else {
		utils.NewMessageWithError(c, err)
	}
}
