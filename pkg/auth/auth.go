package auth

import (
	"Gohub/app/models/user"
	"Gohub/pkg/logger"
	"errors"

	"github.com/gin-gonic/gin"
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

// CurrentUser 从 gin.context 中获取当前登录用户
func CurrentUser(c *gin.Context) user.User {
	userModel, ok := c.MustGet("current_user").(user.User)
	if !ok {
		logger.LogIf(errors.New("无法获取用户信息"))
		return user.User{}
	}
	return userModel
}

// CurrentUID 从 gin.context 中获取当前登录用户 ID
func CurrentUID(c *gin.Context) string {
	return c.GetString("current_user_id")
}