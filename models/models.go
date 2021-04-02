package models

import (
	"log"

	_ "github.com/go-sql-driver/mysql" //myqsl驱动
	"github.com/go-xorm/xorm"
	"xcdh.space/space/config"
)

//X xorm引擎
var X *xorm.Engine

func init() {
	var url string
	url = config.User + ":" + config.Password + "@(" + config.Host + ":" + config.Port + ")/" + config.Name + "?charset=utf8mb4"
	var err error
	X, err = xorm.NewEngine(config.Type, url)
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}
}
