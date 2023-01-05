package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"pekopekopeko.cloud/pcrcbm/common"
	"pekopekopeko.cloud/pcrcbm/routes"
)

func main() {
	r := gin.Default()
	r.Static("/views", "./views")
	r.Use(common.AddCookie())
	r.LoadHTMLGlob("./views/**.html")
	routes.InitRouter(r)
	err := r.Run(":8080")
	if err != nil {
		log.Println(err)
	}
}
