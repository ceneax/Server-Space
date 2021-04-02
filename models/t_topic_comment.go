package models

import (
	"time"
)

type TTopicComment struct {
	TopicId   int64     `xorm:"not null pk comment('主题ID') BIGINT"`
	CommentId int64     `xorm:"not null pk comment('评论ID') BIGINT"`
	Deletedat time.Time `xorm:"comment('删除时间') DATETIME deleted"`
}
