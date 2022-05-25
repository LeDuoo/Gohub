package auth

import (
	v1 "Gohub/app/http/controllers/api/v1"
	"Gohub/app/requests"
	"Gohub/pkg/auth"
	"Gohub/pkg/jwt"
	"Gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

// LoginController 用户控制器
type LoginController struct {
	v1.BaseAPIController
}

type LoginByPasswordRequest struct {
	CaptchaID     string `json:"captcha_id,omitempty" valid:"captcha_id"`
	CaptchaAnswer string `json:"captcha_answer,omitempty" valid:"captcha_answer"`

	LoginID  string `valid:"login_id" json:"login_id"`
	Password string `valid:"password" json:"password,omitempty"`
}

// LoginByPhone 手机号码登录
func (lc *LoginController) LoginByPhone(c *gin.Context) {

	//1.验证表单
	request := requests.LoginByPhoneRequest{}

	if ok := requests.Validate(c, &request, requests.LoginByPhone); !ok {
		return
	}

	//2.登录
	user, err := auth.LoginByPhone(request.Phone)

	if err != nil {
		//失败返回错误信息
		response.Error(c, err, "账号不存在或密码错误!")
	} else {
		//登录成功,生成token返回
		token := jwt.NewJWT().IssueToken(user.GetStringID(), user.Name)

		response.CreatedJSON(c, gin.H{
			"token": token,
		})
	}

}

// LoginByPassword 验证表单，返回长度等于零即通过
func (lc *LoginController) LoginByPassword(c *gin.Context) {

	//获取请求参数
	request := requests.LoginByPasswordRequest{}

	//检测参数
	if ok := requests.Validate(c, &request, requests.LoginByPassword); !ok {
		return

	}

	//登录
	user, err := auth.Attempt(request.LoginID, request.Password)

	if err != nil {
		response.Unauthorized(c, "登录失败")
	} else {
		//生成token
		token := jwt.NewJWT().IssueToken(user.GetStringID(), user.Name)
		response.CreatedJSON(c, gin.H{
			"token": token,
			"data":  user,
		})
	}

}

// RefreshToken 刷新 Access Token
func(lc *LoginController) RefreshToken(c *gin.Context){

	token, err := jwt.NewJWT().RefreshToken(c)

	if err != nil {
		response.Error(c,err,"令牌刷新错误")
	}else{
		response.JSON(c,gin.H{
			"token" : token,
		})
	}
}
