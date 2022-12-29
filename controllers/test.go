package controllers

import "github.com/gin-gonic/gin"

func Test(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	ctx.String(200, username+password)
}
