package controllers

import (
	"github.com/gin-gonic/gin"
	"pekopekopeko.cloud/pcrcbm/common"
	"pekopekopeko.cloud/pcrcbm/models"
)

// JinRu 检测jinru字段是否有值 是则返回以有人在攻击boss 否则添加“用户名”
func JinRu(ctx *gin.Context) {
	// 先获取用户名
	type JinRuJson struct {
		BossID int `json:"bossid"`
	}
	var jinruJson JinRuJson
	_ = ctx.ShouldBindJSON(&jinruJson)
	username := common.CookieToUserName(ctx)
	var jinrustr string
	models.DB.Model(models.Boss{}).Where("bossid = ?", jinruJson.BossID).Pluck("jinru", &jinrustr)
	if jinrustr == " " {
		jinrustr = username
		models.DB.Model(models.Boss{}).Where("bossid = ?", jinruJson.BossID).Update("jinru", jinrustr)
		ctx.Status(200)
	} else {
		ctx.Status(400)
	}
}

// JinRuRevocation 撤销进入，检测值是否与username匹配 是则撤销 否则返回“你不能撤销他人的申请哦” admin无视
func JinRuRevocation(ctx *gin.Context) {
	type JinRuJson struct {
		BossID int `json:"bossid"`
	}
	var jinruJson JinRuJson
	_ = ctx.ShouldBindJSON(&jinruJson)
	username := common.CookieToUserName(ctx)
	var jinrustr string
	models.DB.Model(models.Boss{}).Where("bossid = ?", jinruJson.BossID).Pluck("jinru", &jinrustr)
	if username == "admin" {
		models.DB.Model(models.Boss{}).Where("bossid = ?", jinruJson.BossID).Update("jinru", " ")
		ctx.Status(200)
	} else {
		if jinrustr != username {
			ctx.Status(400)
		} else {
			models.DB.Model(models.Boss{}).Where("bossid = ?", jinruJson.BossID).Update("jinru", " ")
			ctx.Status(200)
		}
	}
}
