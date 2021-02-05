package models

import "github.com/beego/beego/v2/client/orm"

type User struct {
	Id       int    `orm:"column(id);pk"`
	UserID   string `orm:"column(user_id)"`
	Account  string `orm:"column(account)"`
	Pwd      string `orm:"column(pwd)"`
	NickName string `orm:"column(nick_name)"`
	Tel      string `orm:"column(tel)"`
	Email    string `orm:"column(email)"`
	Salt     string `orm:"column(salt)"`
	Status   string `orm:"column(status)"`
	IsDelete string `orm:"column(is_delete)"`
}

func init() {
	// 注册模型
	orm.RegisterModel(new(User))
}

