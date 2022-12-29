package models

type User struct {
	ID       int64
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
}

func (u User) TableName() string {
	return "user"
}

func Register(user *User) error {
	result := DB.Create(&user)
	err := result.Error
	return err

}
