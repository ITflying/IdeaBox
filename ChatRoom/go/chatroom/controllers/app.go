package controllers

import (
	"chatroom/dao"
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "index.html"
}


func (c *MainController) Login() {
	c.Data["key"] = dao.GetFrontUserPubKey()
	c.TplName = "login.html"
}