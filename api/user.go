package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"proj/model"
	"proj/service"
	"proj/util"
)

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	//有效输入
	if username == " " || password == "" {
		util.RespParamErr(c)
		return
	}
	//检索不到用户处理
	u, err := service.SearchUserByUsername(username)
	if err != nil {
		if err == sql.ErrNoRows {
			util.NormErr(c, 300, "用户不存在")
		} else {
			log.Printf("search user error:%v", err)
			util.RespInternalErr(c)
			return
		}
		return
	}
	//密码错误
	if u.Password != password {
		util.NormErr(c, 300, "密码错误")
		return
	}
	util.RespOK(c)
	//设置cookie
	c.SetCookie("name", username, 3600, "/", "localhost", false, true)
}
func AddRepository(c *gin.Context) {
	repository := c.Query("repository")
	name := c.Query("name")
	var m model.Management
	m.Repository = repository
	m.Name = name
	err := service.AddRepository(m)
	if err != nil {
		util.RespInternalErr(c)
		return
	}
	util.RespOK(c)
}
