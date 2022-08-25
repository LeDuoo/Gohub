package v1

import (
	"Gohub/app/models/article"
	"Gohub/app/requests"
	"Gohub/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ArticlesController struct {
	BaseAPIController
}

//文章列表
func (ctrl *ArticlesController) List(c *gin.Context) {
	//分页参数检测
	request := requests.PaginationRequest{}
	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}

	PerPage := 10
	if request.PerPage != "" {
		PerPage, _ = strconv.Atoi(request.PerPage)
	}

	whereMap := make(map[string]string)
	//条件搜索
	if title := c.Query("title"); title != "" {
		whereMap["title"] = title
	}

	data, pager := article.Paginate(c, PerPage, whereMap)

	response.JSON(c, gin.H{
		"data":  data,
		"paper": pager,
	})
}

//新增文章
func (ctrl *ArticlesController) Create(c *gin.Context) {

	request := requests.ArticleRequest{}
	if ok := requests.Validate(c, &request, requests.ArticleSave); !ok {
		return
	}

	articleModel := article.Article{
		Title:   request.Title,
		Content: request.Content,
	}
	articleModel.Create()
	if articleModel.ID > 0 {
		response.Created(c, articleModel)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

//更新文章
func (ctrl *ArticlesController) Update(c *gin.Context) {
	//查找修改model对象
	articleModel := article.Get(c.Param("id"))

	if articleModel.ID == 0 {
		response.Abort404(c)
		return
	}

	//验证修改参数
	request := requests.ArticleRequest{}
	if ok := requests.Validate(c, &request, requests.ArticleSave); !ok {
		return
	}

	articleModel.Title = request.Title
	articleModel.Content = request.Content

	//修改model 返回变动条数
	rowsAffected := articleModel.Save()

	if rowsAffected > 0 {
		response.Success(c)
	} else {
		response.Abort500(c)
	}
}

//删除文章
func (ctrl *ArticlesController) Delete(c *gin.Context) {
	artcileModel := article.Get(c.Param("id"))
	if artcileModel.ID == 0 {
		response.Abort404(c)
		return
	}

	rowsAffected := artcileModel.Delete()
	if rowsAffected > 0 {
		response.Success(c)
		return
	}
	response.Abort500(c, "删除失败,请稍后再试~")
}
