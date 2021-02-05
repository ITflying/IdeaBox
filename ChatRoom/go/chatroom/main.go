package main

import (
	"chatroom/models"
	_ "chatroom/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func init()  {
	models.InitDb()
}

func main() {
	beego.Run()
}

