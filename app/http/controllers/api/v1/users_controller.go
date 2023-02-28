package v1

import (
	"Gohub/app/models/user"
	"Gohub/app/requests"
	"Gohub/pkg/auth"
	"Gohub/pkg/response"

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
