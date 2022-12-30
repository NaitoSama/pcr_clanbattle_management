package common

import (
	"fmt"
	"time"
)

// LocalTime 格式化一个时间类型
type LocalTime time.Time

func (t *LocalTime) MarshalJSON() ([]byte, error) {
	tTime := time.Time(*t)
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format("2006-01-02 15:04:05"))), nil
}

// GetMonthAfterTimeString 格式化1个月后的时间，返回字符串
func GetAfterTimeString(year, month, day int) string {
	return time.Now().AddDate(year, month, day).Format("2006-01-02 15:04:05")
}

// GetNowTimeString 格式化现在的时间，返回字符串
func GetNowTimeString() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
