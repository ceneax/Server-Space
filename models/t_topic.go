package models

import (
	"time"
)

type TTopic struct {
	TopicId       int64     `xorm:"not null pk autoincr comment('主题ID') BIGINT"`
	SectionId     int64     `xorm:"not null comment('板块ID') BIGINT"`
	Title         string    `xorm:"comment('标题') VARCHAR(255)"`
	Shortcontent  string    `xorm:"comment('短内容') VARCHAR(255)"`
	UserId        string    `xorm:"comment('用户ID') CHAR(20)"`
	Istop         int       `xorm:"comment('是否置顶') TINYINT(1)"`
	Commentcount  int64     `xorm:"comment('评论计数') BIGINT"`
	Likecount     int64     `xorm:"comment('点赞计数') BIGINT"`
	Hatecount     int64     `xorm:"comment('踩计数') BIGINT"`
	Sharecount    int64     `xorm:"comment('分享计数') BIGINT"`
	Favoritecount int64     `xorm:"comment('收藏计数') BIGINT"`
	Createdat     time.Time `xorm:"comment('创建日期') DATETIME"`
	Updatedat     time.Time `xorm:"comment('更新日期') DATETIME"`
	Commnetat     time.Time `xorm:"comment('最后评论日期') DATETIME"`
	Deletedat     time.Time `xorm:"comment('删除时间') DATETIME deleted"`
}
