package repository_test

import (
	"testing"

	"xcdh.space/space/entity"
	"xcdh.space/space/repository"
)

func TestGetReplyList(t *testing.T) {
	var pageinfo entity.Pageinfo
	var replyList = make([]entity.ReplyExtends, 0)
	err := repository.GetReplyList(&pageinfo, int64(1), &replyList)
	if err == nil {
		t.Log("GetReplyList ok!")
	} else {
		t.Error(err)
	}
}

func TestAddReply(t *testing.T) {

}
