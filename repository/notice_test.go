package repository_test

import (
	"testing"

	"xcdh.space/space/entity"
	"xcdh.space/space/repository"
)

func TestGetNoticeList(t *testing.T) {
	var noticeList = make([]entity.NoticeExtends, 0)
	err := repository.GetNoticeList("test", &noticeList)
	if err == nil {
		t.Log("GetNoticeList ok!")
	} else {
		t.Error(err)
	}
}
