package repository_test

import (
	"testing"

	"xcdh.space/space/repository"
)

func TestGetByUserId(t *testing.T) {
	user, _, err := repository.GetByUserId("test")
	if err == nil && user.Username == "测试" {
		t.Log("GetByUserId ok!")
	} else {
		t.Error(err)
	}
}
