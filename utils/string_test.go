package utils_test

import (
	"testing"

	"xcdh.space/space/utils"
)

func TestSubstring(t *testing.T) {
	var str = `测试1234%62143，。,。测试用..，nmi![描述]( http://baidu.com/images/1.jpg)测试用测试用
![描述](/upload/images/1.jpg)`

	substr := utils.Substring(str, 100)
	if substr == "测试1234%62143，。,。测试用..，nmi![描述]( http://baidu.com/images/1.jpg)" {
		t.Log("测试ok")
	} else {
		t.Error("测试失败")
	}
}
