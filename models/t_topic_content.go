package models

type TTopicContent struct {
	TopicId   int64  `xorm:"not null pk comment('主题ID') BIGINT"`
	SectionId int64  `xorm:"not null pk comment('板块ID') BIGINT"`
	Content   string `xorm:"comment('内容') LONGTEXT"`
}
