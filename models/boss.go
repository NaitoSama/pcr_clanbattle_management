package models

type Boss struct {
	BossID     int   `gorm:"column:bossid"`
	BossValue  int64 `gorm:"column:boss_value"`
	BossSyume  int   `gorm:"column:boss_syume"`
	SyumeOne   int64 `gorm:"column:syume_one"`
	SyumeTwo   int64 `gorm:"column:syume_two"`
	SyumeThree int64 `gorm:"column:syume_three"`
	SyumeFour  int64 `gorm:"column:syume_four"`
	SyumeFive  int64 `gorm:"column:syume_five"`
}

func (b *Boss) TableName() string {
	return "boss"
}

type BossValue struct {
	First  Boss
	Second Boss
	Third  Boss
	Fourth Boss
	Fifth  Boss
}
