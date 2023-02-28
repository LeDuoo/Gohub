package requests

import (
	"Gohub/pkg/auth"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type UserUpdateProfileRequest struct {
	Name         string `valid:"name" json:"name"`
	City         string `valid:"city" json:"city"`
	Introduction string `valid:"introduction" json:"introduction"`
}

func UserUpdateProfile(data interface{}, c *gin.Context) map[string][]string {

	//查询用户名重复时,过滤掉当前用户ID
	uid := auth.CurrentUID(c)
	rules := govalidator.MapData{
		"name":         []string{"required", "alpha_num", "between:3,20", "not_exists:users,name," + uid},
		"city":         []string{"min_cn:2", "max_cn:20"},
		"introduction": []string{"min_cn:4", "max_cn:240"},
	}
	messages := govalidator.MapData{
		"name": []string{
			"required:名称为必填项",
			"alpha_num:用户名称格式错误,只允许数字和英文",
			"between:名称长度在3~10之间",
			"not_exists:用户名称已存在",
		},
		"city": []string{
			"min_cn:城市名称长度需至少 2 个字",
			"max_cn:城市名称长度不能超过 20 个字",
		},
		"introduction": []string{
			"min_cn:城市名称长度需至少 4 个字",
			"max_cn:城市名称长度不能超过 240 个字",
		},
	}
	return validate(data, rules, messages)
}
