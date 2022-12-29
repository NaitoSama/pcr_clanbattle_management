package controllers

import "github.com/gin-gonic/gin"

func LoginC(ctx *gin.Context) {
	ctx.HTML(200, "login.html", "")

}
