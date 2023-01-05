package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"pekopekopeko.cloud/pcrcbm/common"
	"pekopekopeko.cloud/pcrcbm/models"
	"strings"
)

// GuaShu 检测guashu字段是否有值 是则添加“，用户名” 否则添加“用户名”
func GuaShu(ctx *gin.Context) {
	// 先获取用户名
	type GuaShuJson struct {
		BossID int `json:"bossid"`
	}
	var guashuJson GuaShuJson
	_ = ctx.ShouldBindJSON(&guashuJson)
	username := common.CookieToUserName(ctx)
	var guashustr string
	models.DB.Model(models.Boss{}).Where("bossid = ?", guashuJson.BossID).Pluck("guashu", &guashustr)
	result := strings.Contains(guashustr, username)
	if result {
		ctx.Status(400)
	} else {
		if guashustr == " " {
			guashustr = fmt.Sprintln(guashustr + username)
		} else {
			guashustr = fmt.Sprintln(guashustr + "," + username)
		}
		models.DB.Model(models.Boss{}).Where("bossid = ?", guashuJson.BossID).Update("guashu", guashustr)
		ctx.Status(200)
	}
}
