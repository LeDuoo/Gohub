// Package user 存放用户 Model 相关逻辑
package user

import (
	"Gohub/app/models"
	"Gohub/pkg/database"
	"Gohub/pkg/hash"
)

// User 用户模型
type User struct {
    models.BaseModel

    Name     string `json:"name,omitempty"`
    Email    string `json:"-"`
    Phone    string `json:"-"`
    Password string `json:"-"`

    models.CommonTimestampsField
}
//创建用户
func(userModel *User) Create(){
    database.DB.Create(&userModel)
}

// ComparePassword 密码是否正确
func (userModel *User) ComparePassword(_password string) bool {
    return hash.BcryptCheck(_password, userModel.Password)
}

//保存修改数据
func(userModel *User) Save()(RowsAffected int64){
   result := database.DB.Save(&userModel)
    return result.RowsAffected
}