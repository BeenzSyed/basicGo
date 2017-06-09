package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/BeenzSyed/goWebApp/models"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

/*  Local method activeContent defines an active page layout.
 *  For example the home page as described by landing-layout.tpl
 * 	{{.Header}}
 * 	{{.LayoutContent}}
 * 	{{.Home}}
 * 	{{.Footer}}
 * Note {{.LayoutContent}} is replaced by template content specified by view
 */
func (c *MainController) activeContent(view string) {
	c.Layout = "landing-layout.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Header"] = "header.tpl"
	c.LayoutSections["Home"] = "home.tpl"
	c.LayoutSections["Footer"] = "footer.tpl"
	c.TplNames = view + ".tpl"

	c.Data["Website"] = SiteTitle
	c.Data["xsrftoken"] = template.HTML(hc.XsrfFormHtml())
	c.Data["IsHome"] = true
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["Test"] = "Test string"
	c.TplName = "index.tpl"
}

func SignupPage(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.ServeFile(res, req, "signup.html")
		return
	}

	username := req.FormValue("username")
	password := req.FormValue("password")

	fmt.Print(username)
	fmt.Print(password)
}

func (c *MainController) UserSignUp() {
	c.activeContent("register")

	username := c.Ctx.Input.Param(":username")
	c.Data["Username"] = username

	//check to see if record exists, only add User if record does not exist
	check := models.CheckUser(username)
	c.Data["UserExists"] = check
	if check == 0 {
		err := models.AddUser(username)
		c.Data["Result"] = false
		if err == nil {
			c.Data["Result"] = true
		}
	}

	c.TplName = "register.tpl"
}
