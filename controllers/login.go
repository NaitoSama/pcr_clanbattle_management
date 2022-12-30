package controllers

import "github.com/gin-gonic/gin"

// LoginC 这是返回登录页面
func LoginC(ctx *gin.Context) {
	ctx.HTML(200, "login.html", "")

}
