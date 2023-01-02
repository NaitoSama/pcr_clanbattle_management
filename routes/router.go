package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"pekopekopeko.cloud/pcrcbm/controllers"
)

// todo 将登录和注册的不需要登录状态，其他都需要登录状态
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

	index := r.Group("/index")
	{
		index.GET("", controllers.BossValue)
		index.POST("/attack", controllers.AttackBoss)
		index.POST("/correction", controllers.SetBossSyumeAndValue)
		index.GET("/homepage", controllers.MyHomePage)
	}

}
