package repository_test

import (
	"testing"

	"xcdh.space/space/entity"
	"xcdh.space/space/repository"
)

func TestGetTopicList(t *testing.T) {
	var pageinfo entity.Pageinfo
	var topicList = make([]entity.TopicExtends, 0)
	err := repository.GetTopicList(&pageinfo, int64(1), &topicList)
	if err == nil {
		t.Log("GetTopicList ok!")
	} else {
		t.Error(err)
	}
}
