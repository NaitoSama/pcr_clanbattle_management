package models

// User 申明User结构体 对应user表结构
type User struct {
	//ID       int64
	Username  string `gorm:"column:username"`
	Password  string `gorm:"column:password"`
	Nickname  string `gorm:"column:nickname"`
	Authority int    `gorm:"column:authority"`
}

// TableName 将User结构体和user表相关联
func (u User) TableName() string {
	return "user"
}

// Register 注册时对user表添加数据
func Register(user *User) error {
	result := DB.Create(&user)
	err := result.Error
	return err

}

// RegisterAuth 验证user表中是否已存在username
func RegisterAuth(user *User) bool {
	username := user.Username
	err := DB.Where("username = ?", username).Take(&user)
	if err.Error != nil {
		return true
	} else {
		return false
	}
}
