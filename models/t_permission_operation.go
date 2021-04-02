package models

type TPermissionOperation struct {
	PerId  int64 `xorm:"not null pk comment('权限ID') BIGINT"`
	OperId int64 `xorm:"not null pk comment('操作ID') index BIGINT"`
}
