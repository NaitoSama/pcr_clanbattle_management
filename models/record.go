package models

import (
	"fmt"
	"time"
)

type RecordS struct {
	ID          int        `gorm:"column:id"`
	Username    string     `gorm:"column:username"`
	DateTime    *LocalTime `gorm:"column:time"`
	Target      int        `gorm:"column:target"`
	AttackValue int64      `gorm:"column:attackvalue"`
}

func (r *RecordS) TableName() string {
	return "record"
}

type LocalTime time.Time

func (t *LocalTime) MarshalJSON() ([]byte, error) {
	tTime := time.Time(*t)
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format("2006-01-02 15:04:05"))), nil
}

type Record struct {
	ID          int    `gorm:"column:id"`
	Username    string `gorm:"column:username"`
	DateTime    string `gorm:"column:time"`
	Target      int    `gorm:"column:target"`
	AttackValue int64  `gorm:"column:attackvalue"`
}

func (r *Record) TableName() string {
	return "record"
}
