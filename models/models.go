package models

import (
	"github.com/astaxie/beego/orm"
)

func AddUser(username string) error {

	o := orm.NewOrm()
	o.Using("default")

	err := o.Raw("INSERT INTO users SET username = ?", username).QueryRow()

	return err

}
