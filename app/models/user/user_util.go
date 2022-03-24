package user

import (
	"Gohub/pkg/database"
)

// IsEmailExist 判断email 是否被注册
func IsEmailExist(email string) bool{
	var count int64
	database.DB.Model(User{}).Where("email = ?", email).Ccount(&count)
	return count > 0
}

//IsPhoneExist 检测手机号码是否已存在
func IsPhoneExist(phone string) bool {
	var count int64
	database.DB.Model(User{}).Where("phone = ?",phone).Count(&count)
	return count > 0
}