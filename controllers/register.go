package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"pekopekopeko.cloud/pcrcbm/models"
)

func Register(ctx *gin.Context) {
	name := ctx.PostForm("username")
	password := ctx.PostForm("password")
	user := models.User{
		Username: name,
		Password: password,
	}
	err := models.Register(&user)
	if err != nil {
		log.Fatalln("数据库添加失败" + err.Error())
		ctx.String(500, "注册失败")
	} else {
		ctx.String(200, "注册成功")
	}
}
