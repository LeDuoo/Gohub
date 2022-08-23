package v1

import (
	"Gohub/app/models/category"
	"Gohub/app/requests"
	"Gohub/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoriesController struct {
	BaseAPIController
}

func (ctrl *CategoriesController) Store(c *gin.Context) {

	request := requests.CategoryRequest{}
	if ok := requests.Validate(c, &request, requests.CategorySave); !ok {
		return
	}

	categoryModel := category.Category{
		Name:        request.Name,
		Description: request.Description,
	}
	categoryModel.Create()
	if categoryModel.ID > 0 {
		response.Created(c, categoryModel)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

func (ctrl *CategoriesController) Update(c *gin.Context) {

	//根据url传入id获取模型对象
	categoryModel := category.Get(c.Param("id"))

	if categoryModel.ID == 0 {
		response.Abort404(c)
	}

	//验证参数
	request := requests.CategoryRequest{}
	if ok := requests.Validate(c, &request, requests.CategorySave); !ok {
		return
	}

	//修改模型对象数据
	categoryModel.Name = request.Name
	categoryModel.Description = request.Description

	//保存,返回修改条数
	rowsAffected := categoryModel.Save()

	if rowsAffected > 0 {
		response.Data(c, categoryModel)
	} else {
		response.Abort500(c)
	}
}

func (ctrl *CategoriesController) List(c *gin.Context) {
	//验证分页
	request := requests.PaginationRequest{}
	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}
	perPage := 0
	if request.PerPage != "" {
		perPage, _ = strconv.Atoi(request.PerPage)

	} else {
		perPage = 10
	}

	data, pager := category.Paginate(c, perPage)
	response.JSON(c, gin.H{
		"data":  data,
		"pager": pager,
	})
}
