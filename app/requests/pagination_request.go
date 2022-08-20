package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type PaginationRequest struct {
	Sort    string `valid:"sort" form:"sort"`
	Order   string `valid:"order" form:"order"`
	PerPage string `valid:"per_page" form:"per_page"`
}

//验证规则
func Pagination(data interface{}, c *gin.Context) map[string][]string {

	//参数规格
	rule := govalidator.MapData{
		"sort": []string{ //排序白名单
			"in:id,created_at,updated_at",
		},
		"order": []string{ //正序倒序
			"in:asc,desc",
		},
		"per_page": []string{ //页码
			"numeric_between:2,100",
		},
	}

	//参数描述
	message := govalidator.MapData{
		"sort": []string{
			"in:排序字段仅支持 : id ,created_at, updated_at",
		},
		"order": []string{
			"in:排序规则仅支持 : asc(正序), desc(倒序)",
		},
		"per_page": []string{
			"numeric_between : 每页条数的值介于 2-100之间",
		},
	}

	return validate(data, rule, message)
}
