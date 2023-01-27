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
	r.Use(common.AddCookie(), common.TlsHandler())
	r.LoadHTMLGlob("./views/**.html")
	routes.InitRouter(r)
	err := r.RunTLS(":8080", "./pekopekopeko.online_bundle.pem", "pekopekopeko.online.key")
	//err := r.Run(":8080")
	if err != nil {
		log.Println(err)
	}
}
