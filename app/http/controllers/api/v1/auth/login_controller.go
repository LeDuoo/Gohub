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

// LoginByPhone 手机登录
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
