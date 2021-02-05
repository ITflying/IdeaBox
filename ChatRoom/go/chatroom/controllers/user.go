package controllers

import (
	"chatroom/dao"
	"chatroom/models"
	"chatroom/util"
	_ "crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
}

/*
1、验证用户是否存在(用户名/邮箱/手机号)
2、
*/
func (c *UserController) Login() {
	// 1. 获取参数
	c.Data["json"] = models.InitResponseResultMsg("10000", "服务器异常")

	req := make(map[string]string)
	reqBody := c.Ctx.Input.RequestBody
	err := json.Unmarshal(reqBody, &req)
	if err != nil {
		return
	}

	// 2. 判断用户是否存在
	c.Data["json"] = models.InitResponseResultMsg("10002", "用户不存在")
	var account64 []byte
	if account64, err = base64.StdEncoding.DecodeString(req["account"]); err != nil {
		return
	}
	account := string(account64)
	var userId string
	if util.IsTel(account) {
		if userId, err = dao.CheckTel(account); err != nil {
			c.ServeJSON()
			return
		}
	} else if util.IsEmail(account) {
		if userId, err = dao.CheckEmail(account); err != nil {
			c.ServeJSON()
			return
		}
	} else {
		if userId, err = dao.CheckAccount(account); err != nil {
			c.ServeJSON()
			return
		}
	}

	user, err := dao.GetUserByUserId(userId)
	if err != nil {
		c.ServeJSON()
		return
	}
	fmt.Println(user)

	// 3. 校验密码
	if err = dao.CheckUserPwd(req["pwd"], user.Pwd, user.Salt); err != nil {
		c.Data["json"] = models.InitResponseResultMsg("10003", "账号或密码不正确")
		c.ServeJSON()
		return
	}

	// 4. 返回结果
	res := map[string]string{
		"status": "200",
		"msg":    "完成",
	}
	c.Data["json"] = res
	c.ServeJSON()
}

/*
用户注册
*/
func (c *UserController) RegUser() {
	// 1. 获取参数
	c.Data["json"] = models.InitResponseResultMsg("10000", "服务器异常")

	var user models.User
	reqBody := c.Ctx.Input.RequestBody
	err := json.Unmarshal(reqBody, &user)
	if err != nil {
		return
	}

	// 2. 判断账号、手机号、邮箱是否存在
	if errMsg := dao.CheckReg(user.Account, user.Tel, user.Email, user.NickName); errMsg != "" {
		c.Data["json"] = models.InitResponseResultMsg("10010", errMsg)
		c.ServeJSON()
		return
	}

	// 3. 解密密码并生成新的密码
	salt, newPws, err := dao.CreateNewPwdByFront(user.Pwd)
	if err != nil {
		c.Data["json"] = models.InitResponseResultMsg("10000", "服务器异常")
		c.ServeJSON()
		return
	}
	user.Pwd = newPws
	user.Salt = salt

	// 4. 创建用户插入到数据库当中
	if err := dao.CreateUser(&user); err != nil {
		c.Data["json"] = models.InitResponseResultMsg("10010", "服务器异常")
		c.ServeJSON()
		return
	}

	// 5. 返回结果
	res := map[string]string{
		"status": "200",
		"msg":    "自动跳转到首页",
	}
	c.Data["json"] = res
	c.ServeJSON()
}
