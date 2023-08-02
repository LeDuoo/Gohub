// Package requests 处理请求数据和表单验证
package requests

import (
	"Gohub/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

// ValidatorFunc 验证函数类型 自定义的 ValidatorFunc 类型，允许我们将验证器方法作为回调函数传参。解析完请求后，调用回调函数验证请求：
type ValidatorFunc func(interface{}, *gin.Context) map[string][]string

// Validate 根据传入函数验证参数是否正确
func Validate(c *gin.Context, obj interface{}, handler ValidatorFunc) bool {
	//1.解析请求,是否满足obj的参数要求,支持JSON数据,表单请求和URL Query
	if err := c.ShouldBind(obj); err != nil {
		response.BadRequest(c, err, "请求解析错误, 请确认请求格式是否正确. 上传文件请使用mutipart 标头, 参数请使用JSON 格式.")
		return false
	}

	//2.表单验证,是否满足验证方法里的参数要求
	errs := handler(obj, c)

	//3. 判断是否通过
	if len(errs) > 0 {
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

//validateFile 校验文件
func validateFile(c *gin.Context, data interface{}, rules govalidator.MapData, messages govalidator.MapData) map[string][]string {
	opts := govalidator.Options{
		Request:       c.Request,
		Rules:         rules,
		Messages:      messages,
		TagIdentifier: "valid",
	}
	// 调用 govalidator 的 Validate 方法来验证文件
	return govalidator.New(opts).Validate()
}
