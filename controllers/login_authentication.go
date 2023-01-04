package controllers

import (
	"github.com/gin-gonic/gin"
	"pekopekopeko.cloud/pcrcbm/models"
)

// LoginAuth 登录验证
func LoginAuth(ctx *gin.Context) {
	type JsonUser struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var jsonData JsonUser
	err := ctx.ShouldBindJSON(&jsonData)
	if err != nil {
		panic(err)
	}
	//username := ctx.PostForm("username")
	//password := ctx.PostForm("password")
	user := models.User{}
	err2 := models.DB.Where("username = ? AND password = ?", jsonData.Username, jsonData.Password).Take(&user)
	if err2.Error != nil {
		ctx.String(401, "账号或密码错误")
	} else {

		cookie, _ := ctx.Cookie("peko_cookie")
		session := models.Sessions{}
		models.DB.Model(&session).Where("sessionid = ?", cookie).Updates(models.Sessions{
			Login:   1,
			Account: jsonData.Username,
		})
		//ctx.Redirect(307, "/index")
		ctx.Status(200)
	}
}
