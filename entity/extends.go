package entity

import "xcdh.space/space/models"

type UserInfoExtends struct {
	UserInfo models.TUserInfo `xorm:"extends"`
	Username string           `xorm:"comment('用户名') CHAR(50)"`
	Avatar   string           `xorm:"VARCHAR(255)"`
}

func (UserInfoExtends) TableName() string {
	return "t_user_info"
}

type TopicExtends struct {
	Topic    models.TTopic `xorm:"extends"`
	Username string        `xorm:"comment('用户名') CHAR(50)"`
	Avatar   string        `xorm:"VARCHAR(255)"`
}

func (TopicExtends) TableName() string {
	return "t_topic"
}

type CommentExtends struct {
	Comment  models.TComment `xorm:"extends"`
	Username string          `xorm:"comment('用户名') CHAR(50)"`
	Avatar   string          `xorm:"VARCHAR(255)"`
}

func (CommentExtends) TableName() string {
	return "t_comment"
}

type ReplyExtends struct {
	Reply    models.TReply `xorm:"extends"`
	Username string        `xorm:"comment('用户名') CHAR(50)"`
	Avatar   string        `xorm:"VARCHAR(255)"`
}

func (ReplyExtends) TableName() string {
	return "t_reply"
}

type NoticeExtends struct {
	Notice   models.TNotice `xorm:"extends"`
	Username string         `xorm:"comment('用户名') CHAR(50)"`
	Avatar   string         `xorm:"VARCHAR(255)"`
}

func (NoticeExtends) TableName() string {
	return "t_notice"
}
