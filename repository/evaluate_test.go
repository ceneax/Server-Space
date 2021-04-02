package repository_test

import (
	"testing"

	"xcdh.space/space/repository"
)

func TestGetByUseridAndTypeid(t *testing.T) {
	repository.GetEvaluateByUseridAndTypeid("1", "test", "( 1 , 2 )")
	t.Log("GetCommentList ok!")
}
