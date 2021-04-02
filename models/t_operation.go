package models

type TOperation struct {
	OperId   int64  `xorm:"not null pk autoincr comment('操作ID') BIGINT"`
	Opername string `xorm:"comment('操作名') CHAR(50)"`
	Code     string `xorm:"comment('编码') VARCHAR(255)"`
	Method   string `xorm:"comment('方法') VARCHAR(255)"`
	Url      string `xorm:"comment('URL') VARCHAR(255)"`
	Pid      int64  `xorm:"comment('父操作ID') BIGINT"`
}
