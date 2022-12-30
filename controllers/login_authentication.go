package controllers

import (
	"github.com/gin-gonic/gin"
	"pekopekopeko.cloud/pcrcbm/models"
)

// LoginAuth 登录验证
// TODO 登录后将session改为登录状态并跳转到会战页面
func LoginAuth(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	user := models.User{}
	err := models.DB.Where("username = ? AND password = ?", username, password).Take(&user)
	if err.Error != nil {
		ctx.String(401, "账号或密码错误")
	} else {
		ctx.JSON(200, gin.H{
			"username": username,
			"password": password,
		})
		cookie, _ := ctx.Cookie("peko_cookie")
		session := models.Sessions{}
		models.DB.Model(&session).Where("sessionid = ?", cookie).Updates(models.Sessions{
			Login:   1,
			Account: username,
		})
	}
}
