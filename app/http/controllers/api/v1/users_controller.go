package v1

import (
	"Gohub/app/models/user"
	"Gohub/app/requests"
	"Gohub/pkg/auth"
	"Gohub/pkg/config"
	"Gohub/pkg/file"
	"Gohub/pkg/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

type UsersController struct {
	BaseAPIController
}

// CurrentUser 当前登录用户信息
func (ctrl *UsersController) CurrentUser(c *gin.Context) {
	userModel := auth.CurrentUser(c)
	response.Data(c, userModel)
}

//Index 所有用户
func (ctrl *UsersController) Index(c *gin.Context) {
	request := requests.PaginationRequest{}

	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}
	data, pager := user.Paginate(c, 10)
	response.JSON(c, gin.H{
		"data":  data,
		"pager": pager,
	})
}

//UpdateProfile 编辑个人资料
func (ctrl *UsersController) UpdateProfile(c *gin.Context) {
	//验证请求参数
	request := requests.UserUpdateProfileRequest{}
	if ok := requests.Validate(c, &request, requests.UserUpdateProfile); !ok {
		return
	}
	//验证成功 获取用户模型,修改数据
	userModel := auth.CurrentUser(c)
	userModel.Name = request.Name
	userModel.City = request.City
	userModel.Introduction = request.Introduction
	rowAffected := userModel.Save()
	if rowAffected > 0 {
		response.Data(c, userModel)
	} else {
		response.Abort500(c, "更新失败,请稍后尝试")
	}
}

//UpdateEmail 修改邮箱地址
func (ctrl *UsersController) UpdateEmail(c *gin.Context) {
	//获取参数并验证
	request := requests.UserUpdateEmailRequest{}
	if ok := requests.Validate(c, &request, requests.UserUpdateEmail); !ok {
		return
	}

	userModel := auth.CurrentUser(c)
	userModel.Email = request.Email
	rowAffected := userModel.Save()
	if rowAffected > 0 {
		response.Data(c, userModel)
	} else {
		response.Abort500(c, "修改失败,请稍后尝试")
	}

}

//UpdatePhone 修改用户手机号码
func (ctrl *UsersController) UpdatePhone(c *gin.Context) {
	//获取参数并验证
	request := requests.UserUpdatePhoneRequest{}
	if ok := requests.Validate(c, &request, requests.UserUpdatePhone); !ok {
		return
	}

	//获取用户模型
	userModel := auth.CurrentUser(c)
	userModel.Phone = request.Phone
	rowAffected := userModel.Save()
	fmt.Println("条数", rowAffected)
	if rowAffected > 0 {
		response.Data(c, userModel)
	} else {
		response.Abort500(c, "修改失败稍后尝试")
	}
}

//UpdateUserPassword 修改密码
func (ctrl *UsersController) UpdateUserPassword(c *gin.Context) {
	request := requests.UserUpdatePasswordRequest{}

	if ok := requests.Validate(c, &request, requests.UserUpdatePassword); !ok {
		return
	}

	//入参验证成功,检测密码是否正确
	userModel := auth.CurrentUser(c)
	_, err := auth.Attempt(userModel.Name, request.Password)
	if err != nil {
		response.Unauthorized(c, "原密码不正确")
	} else {
		userModel.Password = request.NewPassword
		rowAffected := userModel.Save()
		if rowAffected > 0 {
			response.Success(c)
		} else {
			response.Abort500(c, "修改失败,稍后尝试")
		}
	}
}

//UpdateAvatar 修改头像
func (ctrl *UsersController) UpdateAvatar(c *gin.Context) {
	request := requests.UserUpdateAvatarRequest{}
	if ok := requests.Validate(c, &request, requests.UserUpdateAvatar); !ok {
		return
	}

	avatar, err := file.SaveUploadAvatar(c, request.Avatar)
	if err != nil {
		response.Abort500(c, "上传头像失败，请稍后尝试~")
		return
	}

	currentUser := auth.CurrentUser(c)
	currentUser.Avatar = config.GetString("app.url") + "/" + avatar
	currentUser.Save()

	response.Data(c, currentUser)
}
