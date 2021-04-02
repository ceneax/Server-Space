package models

import (
	"time"
)

type TUserInfo struct {
	UserId        string    `xorm:"not null pk comment('用户ID') CHAR(20)"`
	Sex           int       `xorm:"comment('性别，1男，2女') TINYINT(1)"`
	Emailverified int       `xorm:"comment('是否验证邮箱') TINYINT(1)"`
	Email         string    `xorm:"comment('邮箱') VARCHAR(255)"`
	Createdat     time.Time `xorm:"comment('创建于') DATETIME"`
	Updatedat     time.Time `xorm:"comment('更新于') DATETIME"`
}
