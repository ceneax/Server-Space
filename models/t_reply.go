package models

import (
	"time"
)

type TReply struct {
	ReplyId   int64     `xorm:"not null pk autoincr comment('回复ID') BIGINT"`
	CommentId int64     `xorm:"not null comment('评论ID') BIGINT"`
	Content   string    `xorm:"comment('内容') TINYTEXT"`
	Createdat time.Time `xorm:"comment('创建日期') DATETIME"`
	Updatedat time.Time `xorm:"comment('更新日期') DATETIME"`
	Ip        int64     `xorm:"comment('IP地址') BIGINT"`
	Pid       int64     `xorm:"comment('父回复') BIGINT"`
	UserId    string    `xorm:"comment('回复人员ID') CHAR(20)"`
	Likecount int64     `xorm:"comment('点赞计数') BIGINT"`
	Hatecount int64     `xorm:"comment('踩计数') BIGINT"`
	Deletedat time.Time `xorm:"comment('删除时间') DATETIME deleted"`
}
