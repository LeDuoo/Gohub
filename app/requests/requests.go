// Package requests 处理请求数据和表单验证
package requests

import (
	"Gohub/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

// ValidatorFunc 验证函数类型
type ValidatorFunc func(interface{}, *gin.Context) map[string][]string

//根据传入函数验证参数是否正确
func Validate(c *gin.Context, obj interface{}, handler ValidatorFunc) bool {
	//1.解析请求,是否满足obj的参数要求,支持JSON数据,表单请求和URL Query
	if err := c.ShouldBind(obj); err != nil {
		// c.AbortWithStatusJSON(http.StatusUnprocessableEntity,gin.H{
		// 	"message" : "请求解析错误, 请确认请求格式是否正确. 上传文件请使用mutipart 标头, 参数请使用JSON 格式.",
		// 	"error"	  : err.Error(),
		// })
		// fmt.Println(err.Error())
		response.BadRequest(c, err, "请求解析错误, 请确认请求格式是否正确. 上传文件请使用mutipart 标头, 参数请使用JSON 格式.")
		return false
	}

	//2.表单验证,是否满足验证方法里的参数要求

	errs := handler(obj, c)

	//3. 判断是否通过
	if len(errs) > 0 {
		// c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
		// 	"message" : "请求验证不通过,具体请查看 errors",
		// 	"errors"	  : errs,
		// })
		response.ValidationError(c, errs)

		return false

	}
	return true
}

func validate(data interface{}, rules govalidator.MapData, messages govalidator.MapData) map[string][]string {

	// 配置选项
	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		TagIdentifier: "valid", // 模型中的 Struct 标签标识符
		Messages:      messages,
	}

	// 开始验证
	return govalidator.New(opts).ValidateStruct()
}
