package routes

import (
	controllers "Gohub/app/http/controllers/api/v1"
	"Gohub/app/http/controllers/api/v1/auth"
	"Gohub/app/http/controllers/api/v1/middlewares"

	"github.com/gin-gonic/gin"
)

// RegisterAPIRoutes 注册路由方法
func RegisterAPIRoutes(r *gin.Engine) {
	// v1路由分组
	v1 := r.Group("/v1")

	v1.Use(middlewares.LimitIP("200-H"))
	{

		authGroup := v1.Group("/auth")
		// 限流中间件：每小时限流，作为参考 Github API 每小时最多 60 个请求（根据 IP）
		// 测试时，可以调高一点
		authGroup.Use(middlewares.LimitIP("1000-H"))
		{
			//验证码控制器
			vcc := new(auth.VerifyCodeController)
			//图片验证码,需要加限流
			authGroup.POST("/verify-codes/captcha", middlewares.LimitPerRoute("50-H"), vcc.ShowCaptcha)
			authGroup.POST("/verify-codes/phone", middlewares.LimitPerRoute("20-H"), vcc.SendUsingPhone)
			authGroup.POST("/verify-codes/email", middlewares.LimitPerRoute("20-H"), vcc.SendUsingEmail)

			//注册控制器
			suc := new(auth.SignupController)
			// 判断手机是否已注册
			authGroup.POST("/signup/phone/exist", middlewares.GuestJWT(), middlewares.LimitPerRoute("60-H"), suc.IsPhoneExist)
			// 判断邮件是否已注册
			authGroup.POST("/signup/email/exist", middlewares.GuestJWT(), middlewares.LimitPerRoute("60-H"), suc.IsEmailExist)
			// 手机号注册用户
			authGroup.POST("/signup/using-phone", middlewares.GuestJWT(), suc.SignupUsingPhone)
			//邮箱注册用户
			authGroup.POST("/signup/using-email", middlewares.GuestJWT(), suc.SignupUsingEmail)

			//登录控制器
			lc := new(auth.LoginController)
			//手机号码登录
			authGroup.POST("/login/login-by-phone", middlewares.GuestJWT(), lc.LoginByPhone)
			//账号密码登录
			authGroup.POST("/login/login-by-password", middlewares.GuestJWT(), lc.LoginByPassword)
			//刷新token
			authGroup.POST("/login/refresh-token", middlewares.AuthJWT(), lc.RefreshToken)

			//重置密码控制器
			pc := new(auth.PasswordController)
			//通过手机号码重置密码
			authGroup.POST("/password-reset/using-phone", middlewares.GuestJWT(), pc.ResetByPhone)
			//通过邮箱重置密码
			authGroup.POST("/password-reset/using-email", middlewares.GuestJWT(), pc.ResetByEamil)

			upc := new(auth.UploadOssController)
			//work code Base64转换图片后上传至阿里云
			authGroup.POST("upload-oss/base64-image-upload", upc.Base64ImageUpload)

		}

	}

	//用户控制器
	uc := new(controllers.UsersController)

	//获取当前用户
	v1.GET("/user", middlewares.AuthJWT(), uc.CurrentUser)

}
