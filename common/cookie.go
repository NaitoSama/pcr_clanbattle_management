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
