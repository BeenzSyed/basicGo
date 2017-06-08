package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type Users struct {
	Id       int
	username string
}

func init() {
	orm.RegisterModel(new(Users))
}

func AddUser(username string) error {
	o := orm.NewOrm()
	o.Using("default")

	err := o.Raw("INSERT INTO users SET username = ?", username).QueryRow()

	return err
}

func CheckUser(username string) int {
	o := orm.NewOrm()
	o.Using("default")

	var exist int
	o.Raw("SELECT EXISTS(SELECT * FROM users WHERE username = ?) as exist;",
		username).QueryRow(&exist)

	fmt.Println(exist)

	return exist
}
