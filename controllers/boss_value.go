package controllers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"pekopekopeko.cloud/pcrcbm/common"
	"pekopekopeko.cloud/pcrcbm/models"
	"strconv"
)

// todo boss的状态调整

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
func PostToContent(ctx *gin.Context) (int, int64, error) {
	boss_id, err := strconv.Atoi(ctx.PostForm("bossid"))
	if err != nil {
		return 0, 0, errors.New("bossid类型错误，应为整数\n")
	}
	attack_value, err := strconv.ParseInt(ctx.PostForm("atkval"), 0, 64)
	if err != nil {
		return 0, 0, errors.New("atkval类型错误，应为整数\n")
	}
	return boss_id, attack_value, nil
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
	boss_id, attack_value, err := PostToContent(ctx)
	if err != nil {
		ctx.String(415, "表单异常数据")
		panic(err)
	}
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

// SetBossSyumeAndValueDetection 调整boss周目和血量的数值检测
func SetBossSyumeAndValueDetection(boss_id, boss_syume int, boss_value int64) error {
	boss_stage := SyumeToStage(boss_syume)
	var boss_max_value int64
	models.DB.Model(models.Boss{}).Where("bossid = ?", boss_id).Pluck(boss_stage, &boss_max_value)
	if boss_value > boss_max_value || boss_value <= 0 {
		errText := fmt.Sprintf("设置数值不对，应大于0并小于等于%d", boss_max_value)
		return errors.New(errText)
	} else {
		return nil
	}
}

// SetBossSyumeAndValueAuth 调整boss周目血量的权限检测
func SetBossSyumeAndValueAuth(username string) error {
	var authority int
	models.DB.Model(models.User{}).Where("username = ?", username).Pluck("authority", &authority)
	if authority == 1 {
		return nil
	} else {
		return errors.New("用户权限不足")
	}
}

// PostToContentWhenCor 获取调整boss的post
func PostToContentWhenCor(ctx *gin.Context) (int, int, int64, error) {
	boss_id, err := strconv.Atoi(ctx.PostForm("bossid"))
	if err != nil {
		return 0, 0, 0, errors.New("bossid类型错误，应为整数\n")
	}
	boss_syume, err := strconv.Atoi(ctx.PostForm("bosssyume"))
	if err != nil {
		return 0, 0, 0, errors.New("bosssyume类型错误，应为整数\n")
	}
	boss_value, err := strconv.ParseInt(ctx.PostForm("bossvalue"), 0, 64)
	if err != nil {
		return 0, 0, 0, errors.New("atkval类型错误，应为整数\n")
	}
	return boss_id, boss_syume, boss_value, nil
}

// SetBossSyumeAndValue 调整boss周目和血量主函数
func SetBossSyumeAndValue(ctx *gin.Context) {
	// 先验证账户是否有权限调整，再验证post的数据是否有效
	username := common.CookieToUserName(ctx)
	err := SetBossSyumeAndValueAuth(username)
	if err != nil {
		ctx.String(403, err.Error())
		panic(err)
	}
	boss_id, boss_syume, boss_value, err := PostToContentWhenCor(ctx)
	if err != nil {
		ctx.String(415, err.Error())
		panic(err)
	}
	err = SetBossSyumeAndValueDetection(boss_id, boss_syume, boss_value)
	if err != nil {
		ctx.String(415, err.Error())
		panic(err)
	}
	models.DB.Model(models.Boss{}).Where("bossid = ?", boss_id).Updates(models.Boss{
		BossValue: boss_value,
		BossSyume: boss_syume,
	})
	ctx.String(200, "boss状态调整成功")
}
