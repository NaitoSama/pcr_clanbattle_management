package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"pekopekopeko.cloud/pcrcbm/models"
	"time"
)

// Register 注册函数，会验证是否用户名已存在
func Register(ctx *gin.Context) {
	name := ctx.PostForm("username")
	password := ctx.PostForm("password")
	user := models.User{
		Username: name,
		Password: password,
	}
	judg := models.RegisterAuth(&user)
	if judg == false {
		ctx.String(403, "该用户已存在")
		time.Sleep(time.Duration(3) * time.Second)
		ctx.Redirect(304, "http://localhost:8080/register")
	} else {
		err := models.Register(&user)
		if err != nil {
			log.Fatalln("数据库添加失败" + err.Error())
			ctx.String(500, "注册失败")
			time.Sleep(time.Duration(3) * time.Second)
			ctx.Redirect(304, "http://localhost:8080/register")
		} else {
			ctx.String(200, "注册成功")
			time.Sleep(time.Duration(3) * time.Second)
			ctx.Redirect(304, "http://localhost:8080/login")
		}
	}
}

// RegisterC 这是返回注册页面
func RegisterC(ctx *gin.Context) {
	ctx.HTML(200, "register.html", "")

}
