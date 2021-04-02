package models

type TPermission struct {
	PerId   int64  `xorm:"not null pk autoincr comment('权限ID') BIGINT"`
	Pername string `xorm:"comment('权限名') CHAR(50)"`
}
