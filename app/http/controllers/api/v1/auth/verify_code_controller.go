package auth

import (
	v1 "Gohub/app/http/controllers/api/v1"
	"Gohub/app/requests"
	"Gohub/pkg/captcha"
	"Gohub/pkg/logger"
	"Gohub/pkg/response"
	"Gohub/pkg/verifycode"

	"github.com/gin-gonic/gin"
)

// VerifyCodeController 用户控制器
type VerifyCodeController struct {
	v1.BaseAPIController
}

// ShowCaptcha 显示图片验证码
func (vc *VerifyCodeController) ShowCaptcha(c *gin.Context) {
	//生成验证码
	id, b64s, err := captcha.NewCaptcha().GenerateCaptcha()
	// 记录错误日志, 应为验证码是用户的入口,出错误时应该记error日志
	logger.LogIf(err)
	//返回给用户
	response.JSON(c, gin.H{
		"captcha_id":    id,
		"captcha_image": b64s,
	})

}

// SendUsingPhone 验证图片验证码是否正确后发送手机验证码
func (vc *VerifyCodeController) SendUsingPhone(c *gin.Context) {

	// 1.验证表单参数是否正确
	request := requests.VerifyCodePhoneRequest{} //所需参数名称
	if ok := requests.Validate(c, &request, requests.VerifyCodePhone); !ok {
		return
	}

	// 2. 发送 SMS  &request取地址后解析json参数,修改了request
	if ok := verifycode.NewVerifyCode().SendSMS(request.Phone); !ok {
		response.Abort500(c, "发送短信失败~")
	} else {
		response.Success(c)
	}

}

// SendUsingEmail 发送 Email 验证码
func (vc *VerifyCodeController) SendUsingEmail(c *gin.Context)  {
	
	//1.验证表单参数是否正确
	request := requests.VerifyCodeEmailRequest{} //所需参数
    if ok := requests.Validate(c, &request, requests.VerifyCodeEmail); !ok {
        return
    }

	//2. 发送 SMS
	err := verifycode.NewVerifyCode().SendEmail(request.Email)

	if err != nil {
		response.Abort500(c,"发送Email 验证码失败")
	}else{
		response.Success(c)
	}
}
