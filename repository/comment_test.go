package repository_test

import (
	"testing"

	"xcdh.space/space/entity"
	"xcdh.space/space/repository"
)

func TestGetCommentList(t *testing.T) {
	var pageinfo entity.Pageinfo
	var commentList = make([]entity.CommentExtends, 0)
	err := repository.GetCommentList(&pageinfo, int64(1), &commentList)
	if err == nil {
		t.Log("GetCommentList ok!")
	} else {
		t.Error(err)
	}
}
