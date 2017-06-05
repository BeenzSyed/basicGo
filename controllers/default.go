package controllers

import (
	"github.com/BeenzSyed/goWebApp/models"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["Test"] = "Test string"
	c.TplName = "index.tpl"
}

func (c *MainController) UserRegister() {
	username := c.Ctx.Input.Param(":username")

	err := models.AddUser(username)

	c.Data["Username"] = username
	c.Data["Result"] = false

	if err == nil {
		c.Data["Result"] = true
	}

	c.TplName = "register.tpl"
}
