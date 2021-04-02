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

func GetCommentList(c *gin.Context) {
	var pageinfo entity.Pageinfo
	if err := c.BindQuery(&pageinfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	topicid, _ := strconv.ParseInt(c.Query("topicid"), 10, 64)
	commentList := service.GetCommentList(&pageinfo, topicid)
	if len(commentList) != 0 {
		commentBackList := make([]entity.CommentBackList, len(commentList))
		v, ok := c.Get(config.IdentityKey)
		if ok {
			userid := v.(string)
			//构造SQL语句中的In条件
			var buffer bytes.Buffer
			buffer.WriteString("(")
			for index := len(commentList) - 1; index > 0; index-- {
				buffer.WriteString(strconv.FormatInt(commentList[index].Comment.CommentId, 10))
				buffer.WriteString(",")
			}
			buffer.WriteString(strconv.FormatInt(commentList[0].Comment.CommentId, 10))
			buffer.WriteString(")")
			evaluateMap := service.GetEvaluateByUseridAndTypeid("2", userid, buffer.String())
			for index, value := range commentList {
				comment := entity.CommentBackList{
					CommentId:      value.Comment.CommentId,
					Content:        value.Comment.Content,
					Createdat:      value.Comment.Createdat,
					Updatedat:      value.Comment.Updatedat,
					Replyedat:      value.Comment.Replyedat,
					Ip:             value.Comment.Ip,
					Havereply:      value.Comment.Havereply,
					UserId:         value.Comment.UserId,
					TopicId:        topicid,
					Likecount:      value.Comment.Likecount,
					Hatecount:      value.Comment.Hatecount,
					Username:       value.Username,
					Avatar:         value.Avatar,
					EvaluateStatus: evaluateMap[value.Comment.CommentId],
					// FavoriteStatus: value.Topic.TopicId,
				}
				commentBackList[index] = comment
			}
		} else {
			for index, value := range commentList {
				comment := entity.CommentBackList{
					CommentId:      value.Comment.CommentId,
					Content:        value.Comment.Content,
					Createdat:      value.Comment.Createdat,
					Updatedat:      value.Comment.Updatedat,
					Replyedat:      value.Comment.Replyedat,
					Ip:             value.Comment.Ip,
					Havereply:      value.Comment.Havereply,
					UserId:         value.Comment.UserId,
					TopicId:        topicid,
					Likecount:      value.Comment.Likecount,
					Hatecount:      value.Comment.Hatecount,
					Username:       value.Username,
					Avatar:         value.Avatar,
					EvaluateStatus: 0,
					FavoriteStatus: 0,
				}
				commentBackList[index] = comment
			}
		}

		commentBack := new(entity.CommentBack)
		commentBack.Pageinfo = pageinfo
		commentBack.CommentBackList = commentBackList
		utils.NewMessageWithData(c, 1, "查询成功", commentBack)
	} else {
		utils.NewMessage(c, 0, "查询无数据")
	}
}

func AddComment(c *gin.Context) {
	var Comment service.Comment
	if err := c.ShouldBindJSON(&Comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ip := c.ClientIP()
	Comment.Ip = utils.IptoInt(ip)
	b, err := Comment.AddComment()
	if b && err == nil {
		utils.NewMessage(c, 1, "创建成功！")
	} else {
		utils.NewMessageWithError(c, err)
	}
}

func UpdateComment(c *gin.Context) {
	var ucomment service.UpdateComment
	if err := c.ShouldBindJSON(&ucomment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ip := c.ClientIP()
	ucomment.Ip = utils.IptoInt(ip)
	b, err := ucomment.UpdateComment()
	if b && err == nil {
		utils.NewMessage(c, 1, "更新成功！")
	} else {
		utils.NewMessageWithError(c, err)
	}
}

func DeleteComment(c *gin.Context) {
	var dcomment service.DeleteComment
	if err := c.ShouldBindJSON(&dcomment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	b, err := dcomment.DeleteComment()
	if b && err == nil {
		utils.NewMessage(c, 1, "删除成功！")
	} else {
		utils.NewMessageWithError(c, err)
	}
}
