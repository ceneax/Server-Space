package models

import (
	"time"
)

type TNotice struct {
	NoticeId    int64     `xorm:"not null pk autoincr comment('通知ID') BIGINT"`
	ResultId    int64     `xorm:"comment('推送ID') BIGINT"`
	Type        string    `xorm:"comment('通知类型') CHAR(1)"`
	Isread      int       `xorm:"comment('是否已读') TINYINT(1)"`
	UserId      string    `xorm:"comment('推送用户ID') CHAR(20)"`
	Createdat   time.Time `xorm:"comment('创建日期') DATETIME"`
	CauseId     int64     `xorm:"comment('引起通知的ID') BIGINT"`
	CauseuserId string    `xorm:"comment('引起通知的用户ID') CHAR(20)"`
}
