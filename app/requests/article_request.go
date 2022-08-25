package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type ArticleRequest struct {
	Title   string `valid:"title" json:"title"`
	Content string `valid:"content" json:"content,omitempty"`
}

//添加,修改文章验证参数
func ArticleSave(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"title":   []string{"required", "min_cn:2", "not_exists:articles,title"},
		"content": []string{"required"},
	}
	messages := govalidator.MapData{
		"title": []string{
			"required:标题为必填项",
			"min_cn:标题长度至少需要 2 个字",
		},
		"content": []string{
			"required:文章内容为必填项",
		},
	}
	return validate(data, rules, messages)
}
