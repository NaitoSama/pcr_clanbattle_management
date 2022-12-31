package controllers

import (
	"github.com/gin-gonic/gin"
	"pekopekopeko.cloud/pcrcbm/models"
)

func BossValue(ctx *gin.Context) {
	//boss := common.ReadJson()
	//ctx.JSON(200, boss)
	var boss []map[string]interface{}
	models.DB.Model(models.Boss{}).Find(&boss)
	ctx.JSON(200, boss[0])
}
