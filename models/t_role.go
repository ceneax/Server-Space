package models

type TRole struct {
	RoleId   int64  `xorm:"not null pk autoincr comment('角色ID') BIGINT"`
	Rolename string `xorm:"comment('角色名') CHAR(20)"`
}
