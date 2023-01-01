package common

import "pekopekopeko.cloud/pcrcbm/models"

func GetBossValue() map[string]map[string]interface{} {
	var boss []map[string]interface{}
	models.DB.Model(models.Boss{}).Find(&boss)
	//var boss_value models.Boss
	boss_value := make(map[string]map[string]interface{})

	boss_value["first"] = boss[0]
	boss_value["second"] = boss[1]
	boss_value["third"] = boss[2]
	boss_value["fourth"] = boss[3]
	boss_value["fifth"] = boss[4]
	return boss_value
}
