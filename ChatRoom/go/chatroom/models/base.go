package models

import (
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

func InitDb()  {
	// 注册驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)

	// 注册数据库
	username,_ := beego.AppConfig.String("mysql::username")
	password,_ := beego.AppConfig.String("mysql::password")
	url,_ := beego.AppConfig.String("mysql::url")
	dataSource := username + ":" + password + url
	orm.RegisterDataBase("default","mysql", dataSource)

	// 自动建表
	orm.RunSyncdb("default", false, true)
}