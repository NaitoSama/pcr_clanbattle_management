package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"pekopekopeko.cloud/pcrcbm/controllers"
)

func InitRouter(r *gin.Engine) {

	r.GET("/login", controllers.LoginC)
	r.GET("/statics/image/:imgid", func(ctx *gin.Context) {
		imgpath := fmt.Sprintf("./statics/image/%s", ctx.Param("imgid"))
		log.Println(ctx.Param("imgid"))
		ctx.FileAttachment(imgpath, ctx.Param("imgid"))
	})
	r.POST("/register", controllers.Register)
	r.POST("/test", controllers.Test)
	r.POST("/login/authentication", controllers.LoginAuth)
}
