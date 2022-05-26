package requests

import (
    "Gohub/app/requests/validators"

    "github.com/gin-gonic/gin"
    "github.com/thedevsaddam/govalidator"
)

type ResetByPhoneRequest struct {
	Phone      string `json:"phone,omitempty" valid:"phone"`
	VerifyCode string `json:"verify_code,omitempty" valid:"verify_code"`
	Password   string `valid:"password" json:"password,omitempty"`
}

// ResetByPhone 验证表单，返回长度等于零即通过
func ResetByPhone(data interface{},c *gin.Context) map[string][]string{
	rules := govalidator.MapData{
		"phone" : []string{"required","digits:11"},
		"verify_code" : []string{"required","digits:6"},
		"password": []string{"required","min:6"},
	}

	messages := govalidator.MapData{
		"phone" : []string{
			"required : 手机号码为必填项,参数名称:phone",
			"digits   : 手机号长度必须为11位的数字",
		},
		"verify_code"  : []string{
			"required : 验证码为必填项,参数名称:verify_code",
			"digits   : 验证码长度必须为6位的数字",
		},
		"password": []string{
            "required:密码为必填项",
            "min:密码长度需大于 6",
        },

	}

	errs := validate(data,rules,messages)

	//检查验证码
	_data := data.(*ResetByPhoneRequest)
	errs = validators.ValidateVerifyCode(_data.Phone,_data.VerifyCode,errs)

	return errs
}