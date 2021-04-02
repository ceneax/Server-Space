package models

import (
	"time"
)

type TInform struct {
	InformId  int64     `xorm:"not null pk autoincr comment('公告ID') BIGINT"`
	Ishome    int       `xorm:"comment('是否首页提示') TINYINT(1)"`
	Type      int       `xorm:"comment('公告类型') INT"`
	Content   string    `xorm:"comment('内容') LONGTEXT"`
	Createdat time.Time `xorm:"comment('创建日期') DATETIME"`
}
