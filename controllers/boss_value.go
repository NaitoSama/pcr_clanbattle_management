package controllers

import (
	"github.com/gin-gonic/gin"
	"pekopekopeko.cloud/pcrcbm/common"
	"pekopekopeko.cloud/pcrcbm/models"
	"strconv"
)

// todo boss的报刀

// BossValue get boss血量信息等
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

// PostToContent IndexPost 获取对boss攻击的post
func PostToContent(ctx *gin.Context) (int, int64) {
	boss_id, err := strconv.Atoi(ctx.PostForm("bossid"))
	if err != nil {
		ctx.String(415, "bossid类型错误，应为整数\n")
	}
	attack_value, err := strconv.ParseInt(ctx.PostForm("atkval"), 0, 64)
	if err != nil {
		ctx.String(415, "atkval类型错误，应为整数\n")
	}
	return boss_id, attack_value
}

// GetBossValueAndSyume 获取目标boss血量，周目数
func GetBossValueAndSyume(boss_id int) (int64, int) {
	var boss_value int64
	var boss_syume int
	models.DB.Model(models.Boss{}).Where("bossid = ?", boss_id).Pluck("boss_value", &boss_value)
	models.DB.Model(models.Boss{}).Where("bossid = ?", boss_id).Pluck("boss_syume", &boss_syume)
	return boss_value, boss_syume
}

// StageToGetBossValue 获取某个阶段的某个boss血量信息
func StageToGetBossValue(boss_id int, boss_stage string) int64 {
	var boss_value int64
	models.DB.Model(models.Boss{}).Where("bossid = ?", boss_id).Pluck(boss_stage, &boss_value)
	return boss_value
}

// SyumeToStage boss周目数对应阶段
func SyumeToStage(boss_syume int) string {
	if boss_syume < 1 {
		panic("数据库boss周目数异常")
	} else if boss_syume >= 1 && boss_syume <= 3 {
		return "stage_one"
	} else if boss_syume >= 4 && boss_syume <= 10 {
		return "stage_two"
	} else if boss_syume >= 11 && boss_syume <= 30 {
		return "stage_three"
	} else if boss_syume >= 31 && boss_syume <= 34 {
		return "stage_four"
	} else {
		return "stage_five"
	}
}

// AttackBoss 攻击boss判定与执行
func AttackBoss(ctx *gin.Context) {
	boss_id, attack_value := PostToContent(ctx)
	boss_value, boss_syume := GetBossValueAndSyume(boss_id)
	if attack_value < boss_value {
		boss_value -= attack_value
		models.DB.Model(models.Boss{}).Where("bossid = ?", boss_id).Update("boss_value", boss_value)
		ctx.String(200, "出刀完成")
	} else {
		boss_syume += 1
		boss_stage := SyumeToStage(boss_syume)
		boss_value = StageToGetBossValue(boss_id, boss_stage)
		models.DB.Model(models.Boss{}).Where("bossid = ?", boss_id).Updates(models.Boss{
			BossValue: boss_value,
			BossSyume: boss_syume,
		})
		ctx.String(200, "出尾刀完成")
	}
}
