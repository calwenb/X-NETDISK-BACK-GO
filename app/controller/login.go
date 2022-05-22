package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"netdisk-back/app/comm"
	"netdisk-back/app/pojo"
	"netdisk-back/app/service"
	"netdisk-back/app/util"
	"strconv"
)

func LoginRouter(e *gin.Engine) {
	//登录
	e.POST("/login", loginH)

	//注册
	e.POST("/register", registerH)

}

func loginH(c *gin.Context) {
	loginName := c.Query("loginName")
	passWord := c.Query("passWord")
	if util.HasNil(loginName, passWord) {
		util.RespBadReq(c)
		return
	}
	rs, user := service.Login(loginName, passWord)
	fmt.Printf("", rs)
	if rs {
		token, err := service.CreateToken(strconv.Itoa(user.Id), user.PassWord)
		if err != nil {
			util.RespFail(c, "创建令牌失败")
			return
		}
		c.SetCookie(comm.TOKEN, token, 0, "/",
			"127.0.0.1", false, false)
		util.RespSucce(c, token)
	} else {
		util.RespFail(c, "账号或密码错误")
	}

}
func registerH(c *gin.Context) {
	userName := c.Query("userName")
	loginName := c.Query("loginName")
	passWord := c.Query("passWord")

	if util.HasNil(userName, loginName, passWord) {
		util.RespBadReq(c)
		return
	}

	rs, user := service.Register(pojo.User{
		UserName:  userName,
		LoginName: loginName,
		PassWord:  passWord,
	})
	if rs {
		token, err := service.CreateToken(strconv.Itoa(user.Id), user.PassWord)
		if err != nil {
			util.RespFail(c, "创建令牌失败")
			return
		}
		c.SetCookie(comm.TOKEN, token, 0, "/",
			"127.0.0.1", false, false)
		util.RespSucce(c, token)
	} else {
		util.RespFail(c, "账号已经存在")
	}

}
