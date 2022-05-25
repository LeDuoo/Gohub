package user

import (
	"Gohub/pkg/database"
)

// IsEmailExist 判断email 是否被注册
func IsEmailExist(email string) bool {
	var count int64
	database.DB.Model(User{}).Where("email = ?", email).Count(&count)
	return count > 0
}

//IsPhoneExist 检测手机号码是否已存在
func IsPhoneExist(phone string) bool {
	var count int64
	database.DB.Model(User{}).Where("phone = ?", phone).Count(&count)
	return count > 0
}

//GetUserByPhone 根据手机号码获取用户信息
func GetUserByPhone(phone string) (userModel User) {
	database.DB.Where("phone = ?", phone).First(&userModel)
	return
}

//根据手机号码/email/名称搜索用户
func GetByMutil(loginID string)(userModel User){
	database.DB.Where("phone = ?",loginID).
	Or("email = ?",loginID).
	Or("name = ?",loginID).
	First(&userModel)
	return
}

//Get 根据用户id获取用户
func Get(userId string)(userModel User){
	database.DB.Where("id = ?",userId).First(&userModel)
	return
}