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

func GetTopicList(c *gin.Context) {
	var pageinfo entity.Pageinfo
	if err := c.BindQuery(&pageinfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sectionid, _ := strconv.ParseInt(c.Query("sectionid"), 10, 64)
	topicList := service.GetTopic(&pageinfo, sectionid)
	topicBackList := make([]entity.TopicBackList, len(topicList))
	if len(topicList) != 0 {
		v, ok := c.Get(config.IdentityKey)
		if ok {
			userid := v.(string)
			//构造SQL语句中的In条件
			var buffer bytes.Buffer
			buffer.WriteString("(")
			for index := len(topicList) - 1; index > 0; index-- {
				buffer.WriteString(strconv.FormatInt(topicList[index].Topic.TopicId, 10))
				buffer.WriteString(",")
			}
			buffer.WriteString(strconv.FormatInt(topicList[0].Topic.TopicId, 10))
			buffer.WriteString(")")
			evaluateMap := service.GetEvaluateByUseridAndTypeid("1", userid, buffer.String())
			for index, value := range topicList {
				topic := entity.TopicBackList{
					TopicId:        value.Topic.TopicId,
					SectionId:      value.Topic.SectionId,
					Title:          value.Topic.Title,
					Shortcontent:   value.Topic.Shortcontent,
					UserId:         value.Topic.UserId,
					Istop:          value.Topic.Istop,
					Commentcount:   value.Topic.Commentcount,
					Likecount:      value.Topic.Likecount,
					Hatecount:      value.Topic.Hatecount,
					Sharecount:     value.Topic.Sharecount,
					Favoritecount:  value.Topic.Favoritecount,
					Createdat:      value.Topic.Createdat,
					Updatedat:      value.Topic.Updatedat,
					Commnetat:      value.Topic.Commnetat,
					Username:       value.Username,
					Avatar:         value.Avatar,
					EvaluateStatus: evaluateMap[value.Topic.TopicId],
					// FavoriteStatus: value.Topic.TopicId,
				}
				topicBackList[index] = topic
			}
		} else {
			for index, value := range topicList {
				topic := entity.TopicBackList{
					TopicId:        value.Topic.TopicId,
					SectionId:      value.Topic.SectionId,
					Title:          value.Topic.Title,
					Shortcontent:   value.Topic.Shortcontent,
					UserId:         value.Topic.UserId,
					Istop:          value.Topic.Istop,
					Commentcount:   value.Topic.Commentcount,
					Likecount:      value.Topic.Likecount,
					Hatecount:      value.Topic.Hatecount,
					Sharecount:     value.Topic.Sharecount,
					Favoritecount:  value.Topic.Favoritecount,
					Createdat:      value.Topic.Createdat,
					Updatedat:      value.Topic.Updatedat,
					Commnetat:      value.Topic.Commnetat,
					Username:       value.Username,
					Avatar:         value.Avatar,
					EvaluateStatus: 0,
					FavoriteStatus: 0,
				}
				topicBackList[index] = topic
			}
		}

		topicBack := new(entity.TopicBack)
		topicBack.Pageinfo = pageinfo
		topicBack.TopicBackList = topicBackList
		utils.NewMessageWithData(c, 1, "查询成功", topicBack)
	} else {
		utils.NewMessage(c, 0, "查询无数据")
	}
}

func AddTopic(c *gin.Context) {
	var topic service.Topic
	if err := c.ShouldBindJSON(&topic); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	b, err := topic.AddTopic()
	if b && err == nil {
		utils.NewMessage(c, 1, "创建成功！")
	} else {
		utils.NewMessageWithError(c, err)
	}
}

func GetTopicContent(c *gin.Context) {
	var topic service.Topic
	if err := c.BindQuery(&topic); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//处理登陆用户的情况
	v, ok := c.Get(config.IdentityKey)
	if ok {
		userid := v.(string)
		topic.EvaluateStatus = service.GetEvaluateIntByUseridAndTypeid("1", userid, topic.TopicId)
		topic.FavoriteStatus = service.GetFavoriteIntByUseridAndTypeid("1", userid, topic.TopicId)
	}
	b, err := topic.GetTopicContent()
	if b && err == nil {
		utils.NewMessageWithData(c, 1, "查询成功", topic)
	} else {
		utils.NewMessageWithError(c, err)
	}
}

func UpdateTopic(c *gin.Context) {
	var utopic service.UpdateTopic
	if err := c.ShouldBindJSON(&utopic); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	b, err := utopic.UpdateTopic()
	if b && err == nil {
		utils.NewMessage(c, 1, "更新成功！")
	} else {
		utils.NewMessageWithError(c, err)
	}
}

func DeleteTopic(c *gin.Context) {
	var dtopic service.DeleteTopic
	if err := c.ShouldBindJSON(&dtopic); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	b, err := dtopic.DeleteTopic()
	if b && err == nil {
		utils.NewMessage(c, 1, "删除成功！")
	} else {
		utils.NewMessageWithError(c, err)
	}
}
