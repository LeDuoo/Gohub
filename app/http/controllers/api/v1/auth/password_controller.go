package auth

import (
	v1 "Gohub/app/http/controllers/api/v1"
	"Gohub/app/models/user"
	"Gohub/app/requests"
	"Gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

// PasswordController 用户控制器
type PasswordController struct {
	v1.BaseAPIController
}

// ResetByPhone 使用手机和验证码重置密码
func (pc *PasswordController) ResetByPhone(c *gin.Context) {

	//请求参数
	request := requests.ResetByPhoneRequest{}
	//验证请求参数
	if ok := requests.Validate(c, &request, requests.ResetByPhone); !ok {
		return
	}

	//根据手机号码查询用户
	userModel := user.GetUserByPhone(request.Phone)

	if userModel.ID == 0 {
		response.Abort404(c)
	} else {
		userModel.Password = request.Password
		userModel.Save()
		response.Success(c)
	}
}

// ResetByEmail 使用邮箱和验证码重置密码
func (pc *PasswordController) ResetByEamil(c *gin.Context) {

	//请求参数
	request := requests.ResetByEamilRequest{}
	//检测参数
	if ok := requests.Validate(c, &request, requests.ResetByEamil); !ok {
		return

	}

	//根据邮箱找到用户 需改用户密码
	userModel := user.GetUserByEmail(request.Email)

	if userModel.ID == 0 {
		response.Abort404(c)
	} else {
		userModel.Password = request.Password
		userModel.Save()
		response.Success(c)
	}
}
