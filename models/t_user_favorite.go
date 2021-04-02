package models

type TUserFavorite struct {
	Id        int64  `xorm:"pk comment('主键ID') BIGINT"`
	UserId    string `xorm:"comment('用户ID') CHAR(20)"`
	SectionId int64  `xorm:"comment('板块ID') BIGINT"`
	TypeId    int64  `xorm:"comment('收藏对应的ID') BIGINT"`
	Type      string `xorm:"comment('收藏类型，1-主题收藏') CHAR(1)"`
	Status    int    `xorm:"comment('收藏状态，-1-取消收藏，1-收藏') TINYINT(1)"`
}
