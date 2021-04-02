package models

import (
	"time"
)

type TRocket struct {
	RocketId       int64     `xorm:"not null pk comment('火箭id') BIGINT"`
	Img            string    `xorm:"comment('图片') VARCHAR(255)"`
	Chinesename    string    `xorm:"comment('中文名') VARCHAR(255)"`
	Englishname    string    `xorm:"comment('英文名') VARCHAR(255)"`
	Type           string    `xorm:"comment('类别') CHAR(50)"`
	Rdagency       int64     `xorm:"comment('研制单位') BIGINT"`
	Leo            string    `xorm:"comment('LEO运载能力') CHAR(100)"`
	Gto            string    `xorm:"comment('GTO运载能力') CHAR(100)"`
	Height         string    `xorm:"comment('箭体高度') CHAR(100)"`
	Liftoffthrust  string    `xorm:"comment('起飞推力') CHAR(100)"`
	Liftoffquality string    `xorm:"comment('起飞质量') CHAR(100)"`
	Status         string    `xorm:"comment('状态') CHAR(20)"`
	Istop          int       `xorm:"comment('是否置顶') TINYINT(1)"`
	Commentcount   int64     `xorm:"comment('评论计数') BIGINT"`
	Likecount      int64     `xorm:"comment('点赞计数') BIGINT"`
	Hatecount      int64     `xorm:"comment('踩计数') BIGINT"`
	Sharecount     int64     `xorm:"comment('分享计数') BIGINT"`
	Favoritecount  int64     `xorm:"comment('收藏计数') BIGINT"`
	Createdat      time.Time `xorm:"comment('创建日期') DATETIME"`
	Updatedat      time.Time `xorm:"comment('更新日期') DATETIME"`
	Commnetat      time.Time `xorm:"comment('最后评论日期') DATETIME"`
	Deletedat      time.Time `xorm:"comment('删除时间') DATETIME"`
	Pid            int64     `xorm:"comment('基本型火箭ID') BIGINT"`
}
