package models

type TUserRole struct {
	UserId string `xorm:"not null pk comment('用户ID') CHAR(20)"`
	RoleId int64  `xorm:"not null pk comment('角色ID') BIGINT"`
}
