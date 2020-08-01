package model

import (
	"encoding/base64"
	"ginblog/utils/err_msg"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/scrypt"
	"log"
)

type User struct {
	gorm.Model
	UserName string `gorm:"type:varchar(20);not null" json:"user_name"`
	Password string `gorm:"type:varchar(20);not null" json:"password"`
	Role     int    `gorm:"type:int" json:"role"`
}

// 查询用户是否存在
func CheckUser(username string) (code int) {
	var users User
	db.Select("id").Where("user_name = ?", username).First(&users)
	if users.ID > 0 {
		return err_msg.ErrorUsernameUsed // 1001
	}
	return err_msg.SUCCESS // 200
}

// 新增用户
func CreateUser(data *User) int {
	//data.Password = ScryptPw(data.Password)
	err := db.Create(&data).Error
	if err != nil {
		return err_msg.ERROR // 500
	}
	return err_msg.SUCCESS // 200
}

// 查询用户列表
func GetUsers(pageSize int, pageNum int) []User {
	var users []User
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return users
}

// 编辑用户信息
func EditUser(id int, data *User) int {
	var user User
	var maps = make(map[string]interface{})
	maps["user_name"] = data.UserName
	maps["role"] = data.Role
	err := db.Model(&user).Where("id = ?", id).Update(maps).Error
	if err != nil {
		return err_msg.ERROR
	}
	return err_msg.SUCCESS
}

// 删除用户
func DeleteUser(id int) int {
	var user User
	err := db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return err_msg.ERROR
	}
	return err_msg.SUCCESS
}

// 密码加密
func (u *User) BeforeSave() {
	u.Password = ScryptPw(u.Password)
}
func ScryptPw(password string) string {
	const (
		PwSaltBytes = 8
		PwHashBytes = 10
	)
	salt := make([]byte, PwSaltBytes)
	salt = []byte{23, 56, 11, 74, 34, 75, 27, 93}

	HashPw, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, PwHashBytes)
	if err != nil {
		log.Fatal(err)
	}
	fpw := base64.StdEncoding.EncodeToString(HashPw)
	return fpw
}
