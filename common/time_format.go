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

// GetNowTimeString 格式化时间datetime返回字符串
func GetNowTimeString() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
