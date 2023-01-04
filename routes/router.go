package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"pekopekopeko.cloud/pcrcbm/common"
	"pekopekopeko.cloud/pcrcbm/controllers"
)

func InitRouter(r *gin.Engine) {
	login := r.Group("/login")
	{
		login.GET("", controllers.LoginC)
		login.POST("/authentication", controllers.LoginAuth)
	}

	static := r.Group("/statics")
	{
		static.GET("/image/:imgid", func(ctx *gin.Context) {
			imgpath := fmt.Sprintf("./statics/image/%s", ctx.Param("imgid"))
			log.Println(ctx.Param("imgid"))
			ctx.FileAttachment(imgpath, ctx.Param("imgid"))
		})
	}

	register := r.Group("/register")
	{
		register.GET("", controllers.RegisterC)
		register.POST("", controllers.Register)
	}

	r.POST("/test", controllers.Test)

	index := r.Group("/index").Use(common.IsLogIned())
	{
		index.GET("", controllers.BossValue)
		//index.POST("", controllers.BossValue)
		index.POST("/attack", controllers.AttackBoss)
		index.POST("/correction", controllers.SetBossSyumeAndValue)
		index.GET("/homepage", controllers.MyHomePage)
		index.GET("/logout", controllers.LogOut)
	}

}
