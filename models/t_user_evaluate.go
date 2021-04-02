package models

type TUserEvaluate struct {
	Id        int64  `xorm:"pk autoincr comment('记录ID') BIGINT"`
	UserId    string `xorm:"comment('用户ID') CHAR(20)"`
	SectionId int64  `xorm:"comment('板块ID') BIGINT"`
	TypeId    int64  `xorm:"comment('主题或评论的ID') BIGINT"`
	Type      string `xorm:"comment('类型，1主题点赞，2，评论点赞，3，回复点赞') CHAR(1)"`
	Status    int    `xorm:"comment('状态，-1取消评价，1点赞，点踩') TINYINT(1)"`
}
