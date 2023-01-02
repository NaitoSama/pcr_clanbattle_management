package controllers

import (
	"github.com/gin-gonic/gin"
	"pekopekopeko.cloud/pcrcbm/common"
	"pekopekopeko.cloud/pcrcbm/models"
)

// MyHomePage 个人主页：有用户名 权限等级 昵称
func MyHomePage(ctx *gin.Context) {
	username := common.CookieToUserName(ctx)
	var user models.User
	models.DB.Model(models.User{}).Where("username = ?", username).Take(&user)
	ctx.JSON(200, gin.H{
		"username":  user.Username,
		"nickname":  user.Nickname,
		"authority": user.Authority,
	})
}
