package common

import (
	"github.com/gin-gonic/gin"
	"pekopekopeko.cloud/pcrcbm/models"
)

// todo 退出登录删除cookie和session

// AddCookie cookie中间件 如果没有cookie则给你cookie并添加session
func AddCookie() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, err := ctx.Cookie("peko_cookie")
		if err != nil {
			value := RandAllString(36)
			ctx.SetCookie("peko_cookie", value, 2592000, "/", "", false, true)
			expiration_time := GetAfterTimeString(0, 0, 30)
			session := models.Sessions{
				SessionID:      value,
				Login:          0,
				ExpirationTime: expiration_time,
			}
			models.DB.Create(&session)
		}
	}
}

// CookieToUserName 通过cookie来返回用户名字
func CookieToUserName(ctx *gin.Context) string {
	cookie, _ := ctx.Cookie("peko_cookie")
	var username string
	models.DB.Model(models.Sessions{}).Where("sessionid = ?", cookie).Pluck("account", &username)
	return username
}

// IsLogIned 中间件，验证cookie是否登录过，否则返回登录界面，不能在/login和/register路径下使用！
func IsLogIned() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cookie, _ := ctx.Cookie("peko_cookie")
		var login int
		models.DB.Model(models.Sessions{}).Where("sessionid = ?", cookie).Pluck("login", &login)
		if login == 0 {
			ctx.Redirect(307, "/login")
			panic("未登录")
		}
	}
}
