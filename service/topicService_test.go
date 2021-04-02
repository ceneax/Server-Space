package service_test

import (
	"testing"
	"time"

	"xcdh.space/space/entity"
	"xcdh.space/space/service"
)

func TestAddTopic(t *testing.T) {
	timeStr := time.Now().Format("20060102150405")
	var title string = "测试主题" + timeStr
	topic := &service.Topic{
		SectionId: 1,
		Title:     title,
		UserId:    "test",
		Content:   "测试主题的内容测试主题的内容测试主题的内容测试主题的内容测试主题的内容测试主题的内容",
	}
	b, err := topic.AddTopic()
	if err == nil && b {
		t.Log("添加主题OK")
	} else {
		t.Error(err)
	}
}

func TestGetTopic(t *testing.T) {
	var pageinfo entity.Pageinfo
	tc := service.GetTopic(&pageinfo, 1)
	if tc != nil {
		t.Log("获取成功")
	}
}

func TestGetTopicContent(t *testing.T) {
	topic := &service.Topic{
		TopicId:   10,
		SectionId: 1,
	}
	b, err := topic.GetTopicContent()
	if err == nil && b {
		t.Log("获取主题内容OK")
	} else {
		t.Error(err)
	}
}
