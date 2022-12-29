package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"pekopekopeko.cloud/pcrcbm/routes"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("./views/**")
	routes.InitRouter(r)
	err := r.Run(":8080")
	if err != nil {
		log.Println(err)
	}
}
