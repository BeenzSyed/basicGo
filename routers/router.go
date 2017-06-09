package routers

import (
	"github.com/BeenzSyed/goWebApp/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	//beego.Router("/user/register/:username", &controllers.MainController{}, "get:UserRegister")
	//beego.Router("/signup/:username", &controllers.MainController{}, "get:UserSignUp")
	beego.Router("/signup/:username", controllers.SignupPage)
	//beego.Router("/login", &controllers.MainController{}, "get:UserLogin")

}
