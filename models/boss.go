package models

type Boss struct {
	BossID      int    `gorm:"column:bossid"`
	BossValue   int64  `gorm:"column:boss_value"`
	BossSyume   int    `gorm:"column:boss_syume"`
	StageOne    int64  `gorm:"column:stage_one"`
	StageTwo    int64  `gorm:"column:stage_two"`
	StageThree  int64  `gorm:"column:stage_three"`
	StageFour   int64  `gorm:"column:stage_four"`
	StageFive   int64  `gorm:"column:stage_five"`
	NowMaxValue int64  `gorm:"column:now_max_value"`
	GuaShu      string `gorm:"column:guashu"`
	JinRu       string `gorm:"column:jinru"`
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
