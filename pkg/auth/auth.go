package auth

import (
	"Gohub/app/models/user"
	"errors"
)

// Attempt 账号密码登录
func Attempt(phone string, password string) (user.User, error) {
	userModel := user.GetByMutil(phone)

	if userModel.ID == 0 {
		return user.User{}, errors.New("账号不存在")
	}

	if !userModel.ComparePassword(password) {
		return user.User{}, errors.New("密码错误")
	}
	return userModel, nil
}

//LoginByPhone 根据手机号码查找用户
func LoginByPhone(phone string) (user.User, error) {
	userModel := user.GetUserByPhone(phone)

	if userModel.ID == 0 {
		return user.User{}, errors.New("手机号码未注册")
	}

	return userModel, nil
}
