package main

import (
	_ "github.com/BeenzSyed/goWebApp/models"
	_ "github.com/BeenzSyed/goWebApp/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:beego@tcp(127.0.0.1:3306)/beegomysql?charset=utf8")
}

func main() {
	orm.Debug = true
	beego.Run()
}
