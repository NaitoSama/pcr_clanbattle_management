package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"pekopekopeko.cloud/pcrcbm/common"
	"pekopekopeko.cloud/pcrcbm/routes"
)

func main() {
	r := gin.Default()
	r.Use(common.AddCookie())
	r.LoadHTMLGlob("./views/**")
	routes.InitRouter(r)
	err := r.Run(":8080")
	if err != nil {
		log.Println(err)
	}
}
