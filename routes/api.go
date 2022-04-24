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
			// 判断邮件是否已注册
            authGroup.POST("/signup/email/exist", suc.IsEmailExist)

			//发送验证码
			vcc := new(auth.VerifyCodeController)
			//图片验证码,需要加限流
			authGroup.POST("/verify-codes/captcha",vcc.ShowCaptcha)
			authGroup.POST("/verify-codes/phone",vcc.SendUsingPhone)
        }

	}

}
