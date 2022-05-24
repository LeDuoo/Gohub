package requests

import(
	"Gohub/app/requests/validators"

    "github.com/gin-gonic/gin"
    "github.com/thedevsaddam/govalidator"
)
//手机号验证码登录参数
type LoginByPhoneRequest struct{
	Phone  string `json:"phone,omitempty" valid:"phone"`
	Verifycode string `json:"verify_code,omitempty" valid:"verify_code"`
}
//密码登录接口参数
type LoginByPasswordRequest struct {
    CaptchaID     string `json:"captcha_id,omitempty" valid:"captcha_id"`
    CaptchaAnswer string `json:"captcha_answer,omitempty" valid:"captcha_answer"`

    LoginID  string `valid:"login_id" json:"login_id"`
    Password string `valid:"password" json:"password,omitempty"`
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

// LoginByPassword 验证表单，返回长度等于零即通过
func LoginByPassword(data interface{}, c *gin.Context) map[string][]string{

	//1 验证规则
	rules := govalidator.MapData{
		"captcha_id" : []string{"required"},
		"captcha_answer" : []string{"required","digits:6"},
		"login_id" : []string{"required","min:3"},
		"password" : []string{"required","min:6"},
	}

	messages := govalidator.MapData{
		"captcha_id" : []string{
			"required: 图片验证码的ID为必填项,参数名称 CaptchaID",
		},
		"captcha_answer" : []string{
			"required: 图片验证码为必填项,参数名称 CaptchaAnswer",
			"digits: 图片验证码长度必须为6位的数字",
		},
		"login_id" : []string{
			"required: 登录ID为必填项,支持手机号/邮箱/姓名,参数名称 LoginID",
			"min: 登录ID长度需大于3",
		},
		"password" : []string{
			"required: 密码为必填项,参数名称 Password",
			"min: 密码长度需大于6",
		},
	}

	//验证参数
	errs := validate(data,rules,messages)

	//图片验证码数据
	_data := data.(*LoginByPasswordRequest)
	errs = validators.ValidateCaptcha(_data.CaptchaID,_data.CaptchaAnswer,errs);

	return errs


}