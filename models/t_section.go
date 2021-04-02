package models

type TSection struct {
	SectionId     int64  `xorm:"not null pk autoincr comment('板块ID') BIGINT"`
	Name          string `xorm:"comment('板块名称') CHAR(50)"`
	Clickrate     int64  `xorm:"comment('点击率') BIGINT"`
	Posts         int64  `xorm:"comment('发帖数') BIGINT"`
	Delflag       int    `xorm:"comment('删除标记') TINYINT(1)"`
	Classfication string `xorm:"comment('分类') CHAR(1)"`
}
