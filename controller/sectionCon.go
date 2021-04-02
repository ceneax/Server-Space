package controller

import (
	"github.com/gin-gonic/gin"
	"xcdh.space/space/service"
	"xcdh.space/space/utils"
)

func GetAllSection(c *gin.Context) {
	sectionList := service.GetAllSection()
	if sectionList != nil {
		utils.NewMessageWithData(c, 1, "查询成功", sectionList)
	} else {
		utils.NewMessage(c, 0, "查询无数据")
	}
}

func AddSection(c *gin.Context) {
	var section service.Section
	section.Name = c.Query("name")
	b, err := section.AddSetion()
	if b && err == nil {
		utils.NewMessage(c, 1, "创建成功！")
	} else {
		utils.NewMessageWithError(c, err)
	}
}
