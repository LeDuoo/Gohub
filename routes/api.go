package routes

import (

	"Gohub/app/http/controllers/api/v1/auth"

	"github.com/gin-gonic/gin"
)

// RegisterAPIRoutes 注册路由方法
func RegisterAPIRoutes(r *gin.Engine) {
	// v1路由分组
	v1 := r.Group("/v1")
	{

		authGroup := v1.Group("/auth")
        {
            suc := new(auth.SignupController)
            // 判断手机是否已注册
            authGroup.POST("/signup/phone/exist", suc.IsPhoneExist)
        }

	}

}
