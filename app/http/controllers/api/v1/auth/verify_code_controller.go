package auth

import (
	v1 "Gohub/app/http/controllers/api/v1"
	"Gohub/pkg/captcha"
	"Gohub/pkg/logger"
	"Gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

// VerifyCodeController 用户控制器
type VerifyCodeController struct {
    v1.BaseAPIController
}

// ShowCaptcha 显示图片验证码
func (vc *VerifyCodeController) ShowCaptcha(c *gin.Context)  {
	//生成验证码
	id, b64s, err := captcha.NewCaptcha().GenerateCaptcha()
	// 记录错误日志, 应为验证码是用户的入口,出错误时应该记error日志
	logger.LogIf(err)
	//返回给用户
	response.JSON(c,gin.H{
		"captcha_id" : id,
		"captcha_image" : b64s,
	})

}
