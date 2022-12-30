package models

// Sessions 对应sessions表
type Sessions struct {
	SessionID      string `gorm:"column:sessionid"`
	Login          int    `gorm:"column:login"`
	ExpirationTime string `gorm:"column:expiration_time"`
}

func (s Sessions) TableName() string {
	return "sessions"
}
