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

func CheckUser(username string) error {
	o := orm.NewOrm()
	o.Using("default")

	users := Users{Id: 1}

	err := o.Read(&users)

	if err == orm.ErrNoRows {
		fmt.Println("No result found.")
	} else if err == orm.ErrMissPK {
		fmt.Println("No primary key found.")
	} else {
		fmt.Printf("Results: %v\n", users.username)
	}
	return err
}
