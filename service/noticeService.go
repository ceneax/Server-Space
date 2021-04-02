package service

import (
	"xcdh.space/space/entity"
	"xcdh.space/space/repository"
)

//GetNoticet 或许所有通知
func GetNoticeList(userid string) []entity.NoticeExtends {
	no := make([]entity.NoticeExtends, 0)
	if repository.GetNoticeList(userid, &no) == nil {
		return no
	}
	return nil
}

//SetNoticeRead 设置通知为已读
func SetNoticeRead(userid string) (bool, error) {
	return repository.SetNoticeRead(userid)
}
