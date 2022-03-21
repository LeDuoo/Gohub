package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterAPIRoutes 注册路由方法
func RegisterAPIRoutes(r *gin.Engine) {
	// v1路由分组
	v1 := r.Group("/v1")
	{

		// 注册一个路由
		v1.GET("/", func(c *gin.Context) {
			// 以 JSON 格式响应
			c.JSON(http.StatusOK, gin.H{
				"Hello": "World!",
			})
		})
	}

}
