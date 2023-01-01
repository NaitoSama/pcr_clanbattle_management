package controllers

import (
	"github.com/gin-gonic/gin"
	"pekopekopeko.cloud/pcrcbm/common"
)

func BossValue(ctx *gin.Context) {
	//boss := common.ReadJson()
	//ctx.JSON(200, boss)

	//var boss []map[string]interface{}
	//models.DB.Model(models.Boss{}).Find(&boss)
	//ctx.JSON(200, boss[0])
	//boss_value := fmt.Sprintln(boss)
	//ctx.String(200, boss_value)

	boss_value := common.GetBossValue()
	ctx.JSON(200, boss_value)
}
