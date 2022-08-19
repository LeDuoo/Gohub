package user

import (
	"Gohub/pkg/app"
	"Gohub/pkg/database"
	"Gohub/pkg/paginator"

	"github.com/gin-gonic/gin"
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

//GetUserByEmail 根据邮箱获取用户信息
func GetUserByEmail(email string) (userModel User) {
	database.DB.Where("email = ?", email).First(&userModel)
	return
}

//根据手机号码/email/名称搜索用户
func GetByMutil(loginID string) (userModel User) {
	database.DB.Where("phone = ?", loginID).
		Or("email = ?", loginID).
		Or("name = ?", loginID).
		First(&userModel)
	return
}

//Get 根据用户id获取用户
func Get(userId string) (userModel User) {
	database.DB.Where("id = ?", userId).First(&userModel)
	return
}

//Index 获取所有用户
func All() (users []User) {
	database.DB.Find(&users)
	return
}

// Paginate 分页内容
func Paginate(c *gin.Context, perPage int) (users []User, paging paginator.Paging) {
    paging = paginator.Paginate(
        c,
        database.DB.Model(User{}),
        &users,
        app.V1URL(database.TableName(&User{})),
        perPage,
    )
    return
}

