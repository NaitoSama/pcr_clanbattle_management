package controllers

import (
	"github.com/gin-gonic/gin"
	"pekopekopeko.cloud/pcrcbm/models"
)

// LogOut 退出登录删除cookie和session
func LogOut(ctx *gin.Context) {
	cookie, _ := ctx.Cookie("peko_cookie")
	models.DB.Where("sessionid = ?", cookie).Delete(models.Sessions{})
	ctx.SetCookie("peko_cookie", "", -1, "/", "", false, true)
	//ctx.Redirect(307, "/login")
	ctx.String(200, "退出登录成功")
}
