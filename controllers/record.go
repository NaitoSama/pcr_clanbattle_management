package controllers

import (
	"github.com/gin-gonic/gin"
	"pekopekopeko.cloud/pcrcbm/models"
)

// Record 查表record返回最新20条记录
func Record(ctx *gin.Context) {
	var record []map[string]interface{}
	models.DB.Model(models.Record{}).Order("time desc").Limit(20).Find(&record)
	records := recordToRecords(record)
	ctx.HTML(200, "record.html", records)
}

func recordToRecords(record []map[string]interface{}) map[string]map[string]interface{} {
	records := make(map[string]map[string]interface{})
	records["one"] = record[0]
	records["two"] = record[1]
	records["three"] = record[2]
	records["four"] = record[3]
	records["five"] = record[4]
	records["six"] = record[5]
	records["seven"] = record[6]
	records["eight"] = record[7]
	records["nine"] = record[8]
	records["ten"] = record[9]
	records["eleven"] = record[10]
	records["twelve"] = record[11]
	records["thirteen"] = record[12]
	records["fourteen"] = record[13]
	records["fifteen"] = record[14]
	records["sixteen"] = record[15]
	records["seventeen"] = record[16]
	records["eighteen"] = record[17]
	records["nineteen"] = record[18]
	records["twenty"] = record[19]
	return records
}
