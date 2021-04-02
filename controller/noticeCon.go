package controller

import (
	"github.com/gin-gonic/gin"
	"xcdh.space/space/config"
	"xcdh.space/space/entity"
	"xcdh.space/space/service"
	"xcdh.space/space/utils"
)

func GetNoticeList(c *gin.Context) {
	userid := c.Param("userid")
	v, ok := c.Get(config.IdentityKey)
	if v.(string) != userid && !ok {
		utils.NewMessage(c, 0, "权限出错！")
		return
	}
	user, b, err := service.GetByUserId(userid)
	if b != true || err != nil {
		utils.NewMessageWithError(c, err)
		return
	}
	noticeList := service.GetNoticeList(userid)
	if len(noticeList) != 0 {
		noticBackList := make([]entity.NoticeBackList, len(noticeList))
		var content, causeContent string
		for index, value := range noticeList {
			switch value.Notice.Type {
			case "T":
				topic, b, err := service.GetTopicById(value.Notice.ResultId)
				if b != true || err != nil {
					utils.NewMessageWithError(c, err)
					return
				}
				content = topic.Title
				comment, b, err := service.GetCommentById(value.Notice.CauseId)
				if b != true || err != nil {
					utils.NewMessageWithError(c, err)
					return
				}
				causeContent = comment.Content
			case "C":
				comment, b, err := service.GetCommentById(value.Notice.ResultId)
				if b != true || err != nil {
					utils.NewMessageWithError(c, err)
					return
				}
				content = comment.Content
				reply, b, err := service.GetReplyById(value.Notice.CauseId)
				if b != true || err != nil {
					utils.NewMessageWithError(c, err)
					return
				}
				causeContent = reply.Content
			case "R":
				reply, b, err := service.GetReplyById(value.Notice.CauseId)
				if b != true || err != nil {
					utils.NewMessageWithError(c, err)
					return
				}
				content = reply.Content
				resultreply, b, err := service.GetReplyById(value.Notice.ResultId)
				if b != true || err != nil {
					utils.NewMessageWithError(c, err)
					return
				}
				causeContent = resultreply.Content
			}
			noticback := entity.NoticeBackList{
				NoticeId:      value.Notice.NoticeId,
				ResultId:      value.Notice.ResultId,
				Type:          value.Notice.Type,
				Isread:        value.Notice.Isread,
				UserId:        value.Notice.UserId,
				Username:      user.Username,
				Avatar:        user.Avatar,
				Content:       content,
				Createdat:     value.Notice.Createdat,
				CauseId:       value.Notice.CauseId,
				CauseuserId:   value.Notice.CauseuserId,
				CauseUsername: value.Username,
				CauseAvatar:   value.Avatar,
				CauseContent:  causeContent,
			}
			noticBackList[index] = noticback
			causeContent = ""
			content = ""
		}
		utils.NewMessageWithData(c, 1, "查询成功", noticBackList)
	} else {
		utils.NewMessage(c, 0, "查询无数据")
	}
}

func SetNoticeRead(c *gin.Context) {
	userid := c.Param("userid")
	v, ok := c.Get(config.IdentityKey)
	if v.(string) != userid && !ok {
		utils.NewMessage(c, 0, "权限出错！")
		return
	}
	b, err := service.SetNoticeRead(userid)
	if b && err == nil {
		utils.NewMessage(c, 1, "查询成功")
	} else {
		utils.NewMessageWithError(c, err)
	}
}
