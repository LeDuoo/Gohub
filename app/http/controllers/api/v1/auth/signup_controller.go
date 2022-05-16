package auth

import (
	v1 "Gohub/app/http/controllers/api/v1"
	"Gohub/app/models/user"
	"Gohub/app/requests"
	"Gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

// SignupController 注册控制器
type SignupController struct {
	v1.BaseAPIController
}

func (sc *SignupController) IsPhoneExist(c *gin.Context) {
	// 初始化请求对象
	request := requests.SignupPhoneExistRequest{}
	//获取请求参数, 并做表单验证
	if ok := requests.Validate(c, &request, requests.SignupPhoneExist); !ok {
		return
	}

	// // 解析 JSON 请求
	// if err := c.ShouldBindJSON(&request); err != nil {
	// 	// 解析失败，返回 422 状态码和错误信息
	// 	c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
	// 		"error": err.Error(),
	// 	})
	// 	// 打印错误信息
	// 	fmt.Println(err.Error())
	// 	// 出错了，中断请求
	// 	return
	// }

	// // 表单验证
	// errs := requests.SignupPhoneExist(&request,c)

	// // errs 返回长度不为0为错误
	// if len(errs) > 0 {
	// 	//验证失败返回422错误码
	// 	c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
	//         "errors": errs,
	//     })
	//     return
	// }
	//  检查数据库并返回响应

	// c.JSON(http.StatusOK, gin.H{
	// 	"exist": user.IsPhoneExist(request.Phone),
	// })

	response.JSON(c, gin.H{
		"exist": user.IsPhoneExist(request.Phone),
	})
}

func (sc *SignupController) IsEmailExist(c *gin.Context) {
	// 初始化请求对象
	request := requests.SignupEmailExistRequest{}

	if ok := requests.Validate(c, &request, requests.SignupEmailExist); !ok {
		return
	}
	// // 解析 JSON 请求
	// if err := c.ShouldBindJSON(&request); err != nil {
	// 	// 解析失败，返回 422 状态码和错误信息
	// 	c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
	// 		"error": err.Error(),
	// 	})
	// 	// 打印错误信息
	// 	fmt.Println(err.Error())
	// 	// 出错了，中断请求
	// 	return
	// }

	// // 表单验证
	// errs := requests.SignupEmailExist(&request,c)

	// // errs 返回长度不为0为错误
	// if len(errs) > 0 {
	// 	//验证失败返回422错误码
	// 	c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
	//         "errors": errs,
	//     })
	//     return
	// }
	//  检查数据库并返回响应

	// c.JSON(http.StatusOK, gin.H{
	// 	"exist": user.IsEmailExist(request.Email),
	// })

	response.JSON(c, gin.H{
		"exist": user.IsEmailExist(request.Email),
	})
}

func (sc *SignupController) SignupUsingPhone(c *gin.Context) {

	//验证表单
	request := requests.SignupUsingPhoneRequest{}

	if ok := requests.Validate(c, &request, requests.SignupUsingPhone); !ok {
		return
	}

	//验证成功 创建用户数据
	_user := user.User{
		Name:     request.Name,
		Phone:    request.Phone,
		Password: request.Password,
	}
	//调用创建用户方法
	_user.Create()

	if _user.ID > 0 {
		response.CreatedJSON(c, gin.H{
			"data": _user,
		})
	} else {
		response.Abort500(c, "创建用户失败, 请稍后尝试~")
	}
}
