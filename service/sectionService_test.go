package service_test

import (
	"testing"
	"time"

	"xcdh.space/space/service"
)

func TestAddSection(t *testing.T) {
	timeStr := time.Now().Format("20060102150405")
	var name string = "测试" + timeStr
	section := &service.Section{Name: name}
	b, err := section.AddSetion()
	if err == nil && b {
		t.Log("添加板块OK")
	} else {
		t.Error(err)
	}
}

func TestGetAllSection(t *testing.T) {
	ts := service.GetAllSection()
	if ts != nil {
		t.Log("获取成功")
	}
}
