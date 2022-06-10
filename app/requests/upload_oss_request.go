// Package requests 处理请求数据和表单验证
package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type Base64ImageUploadRequest struct { //omitempty必填字段
	Base64Data []string `json:"base64_data,omitempty" valid:"base64_data"`
}

//检测base64数据是否正确
func Base64ImageUpload(data interface{}, c *gin.Context) map[string][]string {

	// 自定义验证规则
	rules := govalidator.MapData{
		"base64_data": []string{"required"},
	}

	// 自定义验证出错时的提示
	messages := govalidator.MapData{
		"base64_data": []string{
			"required:Base64编码为必填参数, 参数名称:phone",
		},
	}

	return validate(data, rules, messages)
}
