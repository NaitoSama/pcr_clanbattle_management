package common

import (
	"github.com/gin-gonic/gin"
	"pekopekopeko.cloud/pcrcbm/models"
)

// todo 生成cookie 使用mysql存储sessionid和过期时间

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