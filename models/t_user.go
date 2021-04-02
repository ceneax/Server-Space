package models

type TUser struct {
	UserId   string `xorm:"not null pk comment('用户ID') CHAR(20)"`
	Username string `xorm:"comment('用户名') CHAR(50)"`
	Password string `xorm:"comment('密码') CHAR(20)"`
	Avatar   string `xorm:"comment('头像') VARCHAR(255)"`
}
