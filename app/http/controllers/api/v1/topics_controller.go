package v1

import (
	"Gohub/app/models/category"
	"Gohub/app/models/topic"
	"Gohub/app/policies"
	"Gohub/app/requests"
	"Gohub/pkg/auth"
	"Gohub/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TopicsController struct {
	BaseAPIController
}

func (ctrl *TopicsController) Create(c *gin.Context) {

	request := requests.TopicRequest{}
	if ok := requests.Validate(c, &request, requests.TopicSave); !ok {
		return
	}

	topicModel := topic.Topic{
		Title:      request.Title,
		Body:       request.Body,
		CategoryID: request.CategoryID,
		UserID:     auth.CurrentUID(c),
	}

	topicModel.Create()
	if topicModel.ID > 0 {
		response.Created(c, topicModel)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

func (ctrl *TopicsController) List(c *gin.Context) {
	//分页参数检测
	request := requests.PaginationRequest{}

	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}

	PerPage := 10
	if request.PerPage != "" {
		PerPage, _ = strconv.Atoi(request.PerPage)
	}

	data, pager := topic.Paginate(c, PerPage)

	response.JSON(c, gin.H{
		"data":  data,
		"pager": pager,
	})
}

func (ctrl *TopicsController) Update(c *gin.Context) {

	//修改 验证URL参数 id是否正确
	id := c.Param("id")
	topicModel := topic.Get(id)

	if topicModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanmodifyTopic(c, topicModel); !ok {
		response.Abort403(c)
		return
	}

	// 验证 表单参数
	request := requests.TopicRequest{}
	if ok := requests.Validate(c, &request, requests.TopicSave); !ok {
		return
	}

	//修改操作检测分类是否存在
	if request.CategoryID != topicModel.CategoryID {
		categoryModel := category.Get(request.CategoryID)
		if categoryModel.ID == 0 {
			response.Abort404(c)
			return
		}
	}

	//修改操作
	topicModel.Title = request.Title
	topicModel.Body = request.Body
	topicModel.CategoryID = request.CategoryID
	rowsAffected := topicModel.Save()

	if rowsAffected > 0 {
		response.Data(c, topicModel)
	} else {
		response.Abort500(c)
	}
}

func (ctrl *TopicsController) Delete(c *gin.Context) {
	//获取URL参数
	id := c.Param("id")
	//删除对象是否存在
	topicModel := topic.Get(id)

	if topicModel.ID == 0 {
		response.Abort404(c)
		return
	}

	rowsAffected := topicModel.Delete()
	if rowsAffected > 0 {
		response.Success(c)
		return
	}
	response.Abort500(c, "删除失败, 请稍后尝试~")
}
