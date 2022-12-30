package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"pekopekopeko.cloud/pcrcbm/models"
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
	} else {
		err := models.Register(&user)
		if err != nil {
			log.Fatalln("数据库添加失败" + err.Error())
			ctx.String(500, "注册失败")
		} else {
			ctx.String(200, "注册成功")
		}
	}
}
