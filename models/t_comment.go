package models

import (
	"time"
)

type TComment struct {
	CommentId int64     `xorm:"not null pk autoincr comment('评论ID') BIGINT"`
	Content   string    `xorm:"comment('内容') TINYTEXT"`
	Createdat time.Time `xorm:"comment('创建日期') DATETIME"`
	Updatedat time.Time `xorm:"comment('更新日期') DATETIME"`
	Replyedat time.Time `xorm:"comment('回复日期') DATETIME"`
	Ip        int64     `xorm:"comment('IP地址') BIGINT"`
	Havereply int       `xorm:"comment('是否有回复') TINYINT"`
	UserId    string    `xorm:"comment('评论人ID') CHAR(20)"`
	Likecount int64     `xorm:"comment('点赞计数') BIGINT"`
	Hatecount int64     `xorm:"comment('踩计数') BIGINT"`
}
