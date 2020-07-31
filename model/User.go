package model

import (
	"ginblog/utils/err_msg"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	UserName string `gorm:"type:varchar(20);not null" json:"user_name"`
	Password string `gorm:"type:varchar(20);not null" json:"password"`
	Role int `gorm:"type:int" json:"role"`
}

// 查询用户是否存在
func CheckUser(username string) (code int) {
	var users User
	db.Select("id").Where("username = ?",username).First(&users)
	if users.ID > 0 {
		return err_msg.ErrorUsernameUsed // 1001
	}
	return err_msg.SUCCESS // 200
}

// 新增用户
func CreateUser(data *User) int {
	err := db.Create(&data).Error
	if err != nil {
		return err_msg.ERROR // 500
	}
	return err_msg.SUCCESS // 200
}