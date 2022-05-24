package requests

import(
	"Gohub/app/requests/validators"

    "github.com/gin-gonic/gin"
    "github.com/thedevsaddam/govalidator"
)

type LoginByPhoneRequest struct{
	Phone  string `json:"phone,omitempty" valid:"phone"`
	Verifycode string `json:"verify_code,omitempty" valid:"verify_code"`
}

//LoginByPhone 验证表单,返回长度等于零为通过验证
func LoginByPhone(data interface{}, c *gin.Context) map[string][]string{
	//验证规则
	rules := govalidator.MapData{
		"phone" : []string{"required","digits:11"},
		"verify_code" : []string{"required","digits:6"},
	}
	//提示文字
	messages := govalidator.MapData{
		"phone" : []string{
			"required:手机号为必填, 参数名称phone",
			"digits:手机号长度必须为11位数字",
		},
		"verify_code":[]string{
			"required:验证码为必填,参数名称verify_code",
			"digits:验证码长度必须为6位数字",
		},
	}

	errs := validate(data,rules,messages)

	//手机验证码
	_data := data.(*LoginByPhoneRequest)
	//检测验证码是否正确
	errs = validators.ValidateVerifyCode(_data.Phone,_data.Verifycode,errs)

	return errs
}