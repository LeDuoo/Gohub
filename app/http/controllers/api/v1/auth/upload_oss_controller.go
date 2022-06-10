package auth

import (
	v1 "Gohub/app/http/controllers/api/v1"
	"Gohub/app/requests"
	"Gohub/pkg/auth"
	"Gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

// UploadOssController 注册控制器
type UploadOssController struct {
	v1.BaseAPIController
}

//检测手机号码是否存在接口
func (uoc *UploadOssController) Base64ImageUpload(c *gin.Context) {
	// 初始化请求对象
	request := requests.Base64ImageUploadRequest{}
	//获取请求参数, 并做表单验证
	if ok := requests.Validate(c, &request, requests.Base64ImageUpload); !ok {
		return
	}

	//实现业务逻辑
	urlData, err := auth.EncodeBase64Upload(c, request.Base64Data)

	if err != nil {
		response.Error(c, err, "解码base64数据并上传至阿里云Oss失败")
	}

	response.CreatedJSON(c, gin.H{
		"code":    200,
		"message": "success",
		"data":    urlData,
	})
}
