package routers

import (
	"chatroom/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/Login", &controllers.MainController{}, "*:Login")

    // 登录相关API
    beego.Router("/login", &controllers.UserController{}, "post:Login")
    beego.Router("/regUser", &controllers.UserController{}, "post:RegUser")
}
