package controller

import (
	"bytes"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"xcdh.space/space/config"
	"xcdh.space/space/entity"
	"xcdh.space/space/service"
	"xcdh.space/space/utils"
)

func GetReplyList(c *gin.Context) {
	var pageinfo entity.Pageinfo
	if err := c.BindQuery(&pageinfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	commentid, _ := strconv.ParseInt(c.Query("commentid"), 10, 64)
	replyList := service.GetReplyList(&pageinfo, commentid)
	if len(replyList) != 0 {
		replyBackList := make([]entity.ReplyBackList, len(replyList))
		v, ok := c.Get(config.IdentityKey)
		if ok {
			userid := v.(string)
			//构造SQL语句中的In条件
			var buffer bytes.Buffer
			buffer.WriteString("(")
			for index := len(replyList) - 1; index > 0; index-- {
				buffer.WriteString(strconv.FormatInt(replyList[index].Reply.ReplyId, 10))
				buffer.WriteString(",")
			}
			buffer.WriteString(strconv.FormatInt(replyList[0].Reply.ReplyId, 10))
			buffer.WriteString(")")
			evaluateMap := service.GetEvaluateByUseridAndTypeid("3", userid, buffer.String())
			for index, value := range replyList {
				reply := entity.ReplyBackList{
					ReplyId:        value.Reply.ReplyId,
					CommentId:      value.Reply.CommentId,
					Content:        value.Reply.Content,
					Createdat:      value.Reply.Createdat,
					Updatedat:      value.Reply.Updatedat,
					Ip:             value.Reply.Ip,
					Pid:            value.Reply.Pid,
					UserId:         value.Reply.UserId,
					Likecount:      value.Reply.Likecount,
					Hatecount:      value.Reply.Hatecount,
					Username:       value.Username,
					Avatar:         value.Avatar,
					EvaluateStatus: evaluateMap[value.Reply.ReplyId],
					// FavoriteStatus: value.Topic.TopicId,
				}
				replyBackList[index] = reply
			}
		} else {
			for index, value := range replyList {
				reply := entity.ReplyBackList{
					ReplyId:        value.Reply.ReplyId,
					CommentId:      value.Reply.CommentId,
					Content:        value.Reply.Content,
					Createdat:      value.Reply.Createdat,
					Updatedat:      value.Reply.Updatedat,
					Ip:             value.Reply.Ip,
					Pid:            value.Reply.Pid,
					UserId:         value.Reply.UserId,
					Likecount:      value.Reply.Likecount,
					Hatecount:      value.Reply.Hatecount,
					Username:       value.Username,
					Avatar:         value.Avatar,
					EvaluateStatus: 0,
					// FavoriteStatus: value.Topic.TopicId,
				}
				replyBackList[index] = reply
			}
		}

		replyBack := new(entity.ReplyBack)
		replyBack.Pageinfo = pageinfo
		replyBack.ReplyBackList = replyBackList

		utils.NewMessageWithData(c, 1, "查询成功", replyBack)
	} else {
		utils.NewMessage(c, 0, "查询无数据")
	}
}

func AddReply(c *gin.Context) {
	var Reply service.Reply
	if err := c.ShouldBindJSON(&Reply); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ip := c.ClientIP()
	Reply.Ip = utils.IptoInt(ip)
	b, err := Reply.AddReply()
	if b && err == nil {
		utils.NewMessage(c, 1, "创建成功！")
	} else {
		utils.NewMessageWithError(c, err)
	}
}

func UpdateReply(c *gin.Context) {
	var utopic service.UpdateReply
	if err := c.ShouldBindJSON(&utopic); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	b, err := utopic.UpdateReply()
	if b && err == nil {
		utils.NewMessage(c, 1, "更新成功！")
	} else {
		utils.NewMessageWithError(c, err)
	}
}

func DeleteReply(c *gin.Context) {
	var dtopic service.DeleteReply
	if err := c.ShouldBindJSON(&dtopic); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	b, err := dtopic.DeleteReply()
	if b && err == nil {
		utils.NewMessage(c, 1, "删除成功！")
	} else {
		utils.NewMessageWithError(c, err)
	}
}
