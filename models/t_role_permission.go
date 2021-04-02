package models

type TRolePermission struct {
	RoleId int64 `xorm:"not null pk comment('角色ID') BIGINT"`
	PerId  int64 `xorm:"not null pk comment('权限ID') index BIGINT"`
}
