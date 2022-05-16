// Package requests 处理请求数据和表单验证
package requests

import (
	"Gohub/app/requests/validators"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type SignupPhoneExistRequest struct { //omitempty必填字段
	Phone string `json:"phone,omitempty" valid:"phone"`
}

type SignupEmailExistRequest struct { //omitempty必填字段
	Email string `json:"email,omitempty" valid:"email"`
}

// SignupUsingPhoneRequest 通过手机注册的请求信息
type SignupUsingPhoneRequest struct {
	Phone           string `json:"phone,omitempty" valid:"phone"`
	VerifyCode      string `json:"verify_code,omitempty" valid:"verify_code"`
	Name            string `valid:"name" json:"name"`
	Password        string `valid:"password" json:"password,omitempty"`
	PasswordConfirm string `valid:"password_confirm" json:"password_confirm,omitempty"`
}

//手机号码检测是否存在
func SignupPhoneExist(data interface{}, c *gin.Context) map[string][]string {

	// 自定义验证规则
	rules := govalidator.MapData{
		"phone": []string{"required", "digits:11"},
	}

	// 自定义验证出错时的提示
	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号为必填项，参数名称 phone",
			"digits:手机号长度必须为 11 位的数字",
		},
	}

	return validate(data, rules, messages)
}

//邮箱码检测是否存在
func SignupEmailExist(data interface{}, c *gin.Context) map[string][]string {

	// 自定义验证规则
	rules := govalidator.MapData{
		"email": []string{"required", "min:4", "max:30", "email"},
	}

	// 自定义验证出错时的提示
	messages := govalidator.MapData{
		"email": []string{
			"required:Email为必填项",
			"min:Email 长度需大于 4",
			"max:Email 长度需小于 30",
			"email:Email 格式不正确，请提供有效的邮箱地址",
		},
	}

	return validate(data, rules, messages)
}

//用户注册
func SignupUsingPhone(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"phone":            []string{"required", "digits:11", "not_exists:users,phone"},
		"name":             []string{"required", "alpha_num", "between:3,20", "not_exists:users,name"},
		"password":         []string{"required", "min:6"},
		"password_confirm": []string{"required"},
		"verify_code":      []string{"required", "digits:6"},
	}

	message := govalidator.MapData{
        "phone": []string{
            "required:手机号为必填项，参数名称 phone",
            "digits:手机号长度必须为 11 位的数字",
        },
        "name": []string{
            "required:用户名为必填项",
            "alpha_num:用户名格式错误，只允许数字和英文",
            "between:用户名长度需在 3~20 之间",
        },
        "password": []string{
            "required:密码为必填项",
            "min:密码长度需大于 6",
        },
        "password_confirm": []string{
            "required:确认密码框为必填项",
        },
        "verify_code": []string{
            "required:验证码答案必填",
            "digits:验证码长度必须为 6 位的数字",
        },
	}

	errs := validate(data,rules,message)

	_data := data.(*SignupUsingPhoneRequest)

	//调用验证密码方法检测密码是否一致
	errs  = validators.ValidatePasswordConfirm(_data.Password,_data.PasswordConfirm,errs)
	//调用检测验证码是否正确方法
	errs  = validators.ValidateVerifyCode(_data.Phone, _data.VerifyCode,errs)

	return errs
}
